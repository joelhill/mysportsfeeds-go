package blaster

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strings"
	"syscall"
	"time"

	"github.com/InVisionApp/go-logger"
	"github.com/opentracing/opentracing-go"
)

const (
	requestIDHeader      = "Request-ID"
	requestSourceHeader  = "Request-Source"
	callingServiceHeader = "Calling-Service"
	jsonType             = "application/json"
	contentTypeHeader    = "Content-Type"
	userAgentHeader      = "User-Agent"
	contentLengthHeader  = "Content-Length"
	acceptHeader         = "Accept"
	requestTimeout       = 8 * time.Second        // the max amount of time for the entire request before failing
	sockTimeout          = 2 * time.Second        // the max amount of time attempting to make the tcp connection
	tlsTimeout           = 2 * time.Second        // the max amount of time establishing TLS handshake
	idleTimeout          = 10 * time.Second       // the amount of time to keep idle connections available before closing them
	keepAlive            = 750 * time.Millisecond // the keep-alive period for an active network connection
	maxIdleConnsPerHost  = 100                    // the maximum number of idle connections to keep around per host
	maxIdleConns         = 100                    // the maximum number of idle connections to keep around for ALL hosts
)

// NAME is the name of this library
const NAME = "blaster"

// region STRUCT

// ClientOptions encapsulates all available options for a new client
type ClientOptions struct {
	Endpoint                   string
	RouteMask                  string
	CalledService              string
	WillSaturate               interface{}
	WillSaturateOnError        interface{}
	WillSaturateWithStatusCode map[int]interface{}
	TimeoutMS                  int
	CircuitBreaker             CircuitBreakerPrototype
	Headers                    map[string]string
	KeepRawResponse            bool
	Logger                     log.Logger
}

// Client encapsulates the http Request functionality
type Client struct {
	// prototype will be saturated when the Request succeeds.
	prototype interface{}

	// errorPrototype will be saturated when the Request fails.
	// A Request is implicitly considered a failure if the
	// status code of the Response is not in the 2XX range.
	errorPrototype interface{}

	// endpoint is the destination for the http Request
	endpoint *url.URL

	// customPrototypes is a map of interfaces that
	// will be saturated when specific response codes
	// are returned from the endpoint
	customPrototypes map[int]interface{}

	// duration is the length of time the request took to run.
	// Obviously this only has value after the request is run.
	duration time.Duration

	// Internal circuit breaker
	cb CircuitBreakerPrototype

	// if the http response code is < 200 or > 299, this flag
	// gets set true
	responseIsError bool

	// internal http client
	client *http.Client

	// internal headers
	headers map[string]string

	// request method
	method string

	// flag to copy raw response bytes from http response
	keepRawResponse bool

	// raw response bytes
	rawresponse []byte

	// internal statsd client
	statsdClient StatsdClientPrototype

	// statsd stat to record
	statsdStat string

	// statsd tags
	statsdTags []string

	// flag to set this object in an error state
	// this will prevent statsd calls if an error
	// originated within this API
	internalError bool

	// logger that lives throughout request lifecycle, set in Do()
	logger log.Logger

	// openTracingSpan gets attached right before request is made
	openTracingSpan opentracing.Span

	// status code gets tacked on after the request
	statusCode int

	// the service called
	calledService string

	// mask of the route
	routeMask string
}

// endregion

// region UNEXPORTED FUNCS

// applyContextDependentHeaders will apply headers right before
// the request is launched
func (c *Client) applyContextDependentHeaders(ctx context.Context) {
	if pkgRequestIDProviderFunc != nil {
		if requestID, ok := pkgRequestIDProviderFunc(ctx); ok {
			c.headers[requestIDHeader] = requestID
		}
	}

	if pkgRequestSourceProviderFunc != nil {
		if requestSource, ok := pkgRequestSourceProviderFunc(ctx); ok {
			c.headers[requestSourceHeader] = requestSource
		}
	}
}

// reports the duration of the request
func (c *Client) statsdReportDuration() {
	if c.statsdClient != nil {
		tags := []string{
			fmt.Sprintf("response-code:%d", c.statusCode),
			fmt.Sprintf("response-type:%s", responseTypeForStatusCode(c.statusCode)),
			fmt.Sprintf("http-verb:%s", c.method),
			fmt.Sprintf("called-service:%s", c.calledService),
			fmt.Sprintf("route:%s", c.routeMask),
		}
		tags = append(c.statsdTags, tags...)
		c.statsdClient.Timing(c.statsdStat, c.duration, tags, pkgStatsdRate)
	}
}

// make sure the request conforms to invision request tracing policy
func (c *Client) conformsToReq014(request *http.Request) error {
	// add all headers, and also prepare the request
	// tracing headers to be validated
	check := req014HeaderCheck{}
	for k, v := range c.headers {
		request.Header.Set(k, v)
		if v != "" {
			switch k {
			case requestIDHeader:
				check.requestIDOK = true
			case requestSourceHeader:
				check.requestSourceOK = true
			case callingServiceHeader:
				check.callingServiceOK = true
			}
		}
	}

	// if we are strictly enforcing request tracing
	if pkgRequireHeaders && !check.ok() {
		return errors.New("request tracing header requirements check failed")
	}

	return nil
}

// marshal/serialize the outgoing payload if it exists
func (c *Client) processOutgoingPayload(payload interface{}) ([]byte, error) {
	var (
		payloadErr   error
		payloadBytes []byte
	)

	// process the payload if it exists
	if payload != nil {
		// if it's a json Request, marshal the payload.
		// unless changed explicitly, this will be a json
		// request
		if c.headers[contentTypeHeader] == jsonType {
			payloadBytes, payloadErr = json.Marshal(payload)
			if payloadErr != nil {
				return nil, payloadErr
			}
		} else {
			// caller has supplied content-type.  it must be convertible to byte slice
			switch payload.(type) {
			case []byte:
				payloadBytes = payload.([]byte)
			case string:
				payloadBytes = []byte(payload.(string))
			default:
				errBS := errors.New("the payload cannot be converted to a byte slice")
				return nil, errBS
			}
		}

		// if we have a body length, set the content length header
		c.headers[contentLengthHeader] = fmt.Sprintf("%d", len(payloadBytes))
	}

	return payloadBytes, nil
}

// begin tracking request
func (c *Client) immediatePreflight(ctx context.Context, request *http.Request) {
	// if tracing is enabled, wrap the request with the tracing provider
	if pkgTracerProviderFunc != nil {
		var span opentracing.Span
		// The openTracingSpan name needs to be sufficiently generic to avoid a grouping issue in Lightstep (breaking their search).
		// It should not be the full URL, URI or Path, as that often inclues IDs.
		// Note that 'url' is recorded, but as a tag on the openTracingSpan, from https://github.com/InVisionApp/opentracing-go-helpers
		request, span = pkgTracerProviderFunc(ctx, fmt.Sprintf("%s %s", c.method, c.endpoint.Host), request)
		c.openTracingSpan = span
	}
}

// process response
func (c *Client) processResponseData(payload []byte, contentType string) error {
	// if the response has a body, handle it
	if len(payload) > 0 {

		// the thing we are about to potentially unmarshal into
		var unmarshalTo interface{}

		// if there is a custom response for this specific status code
		if c.customPrototypes[c.statusCode] != nil {
			unmarshalTo = c.customPrototypes[c.statusCode]
		} else if c.responseIsError {
			// request returned error code
			unmarshalTo = c.errorPrototype
		} else {
			// request succeeded
			unmarshalTo = c.prototype
		}

		// if there is something that can be unmarshalled into
		if unmarshalTo != nil {
			// a json content-type could be something like `application/json` or `application/json; charset=utf8`
			if strings.Split(contentType, ";")[0] == jsonType {
				decoder := json.NewDecoder(bytes.NewReader(payload))
				if decodeErr := decoder.Decode(unmarshalTo); decodeErr != nil {
					return decodeErr
				}
			} else {
				// This is not the expected result, so it should be logged as a warning.
				// Any non-json responses should be accessed via the raw bytes of the client.
				// Realistically the only thing that should make its way into this block is
				// a non-json error response.
				c.rawresponse = payload
				c.logger.WithFields(map[string]interface{}{
					"type": NAME,
				}).Warn("blaster: received a non-json response where a json type was expected")

				c.statusCode = http.StatusUnprocessableEntity
			}
		}
	}

	return nil
}

// close tracking
func (c *Client) cleanup() {
	if !c.internalError {
		c.statsdReportDuration()
		if c.openTracingSpan != nil {
			c.openTracingSpan.Finish()
		}
	}
}

// the request cannot be launched
func (c *Client) failBeforeRequest(err error) (int, error) {
	c.logger.WithFields(map[string]interface{}{
		"error_message": err.Error(),
		"type":          NAME,
	}).Error("request failed")
	c.statusCode = http.StatusInternalServerError
	c.internalError = true
	return c.statusCode, err
}

// the request happened, but was an error
func (c *Client) failAfterRequest(err error) (int, error) {
	c.logger.WithFields(map[string]interface{}{
		"error_message": err.Error(),
		"type":          NAME,
	}).Error("request failed")
	c.statusCode = http.StatusInternalServerError
	return c.statusCode, err
}

// doInternal will perform the actual request.  This function
// is either called from within a circuit breaker, or directly
// from Do.
func (c *Client) doInternal(ctx context.Context, payload interface{}) (int, error) {

	// set headers that depend on context values
	c.applyContextDependentHeaders(ctx)

	// start the clock and report the duration when this function exits
	defer func(c *Client, begin time.Time) {
		c.duration = time.Now().Sub(begin)
		c.cleanup()
	}(c, time.Now())

	// process outgoing payload
	payloadBytes, payloadErr := c.processOutgoingPayload(payload)
	if payloadErr != nil {
		return c.failBeforeRequest(payloadErr)
	}

	// create the internal HTTP request
	request, createRequestErr := http.NewRequest(c.method, c.endpoint.String(), ioutil.NopCloser(bytes.NewReader(payloadBytes)))
	if createRequestErr != nil {
		return c.failBeforeRequest(createRequestErr)
	}

	// make sure that request conforms to REQ014 if its required
	if req014Err := c.conformsToReq014(request); req014Err != nil {
		return c.failBeforeRequest(req014Err)
	}

	c.logger.WithFields(map[string]interface{}{
		"type": NAME,
	}).Debugf("launching %s request to %s", c.method, c.endpoint.Host)

	// RUN IT
	c.immediatePreflight(ctx, request)
	// --------------------------------------------
	// --------------------------------------------
	response, responseErr := c.client.Do(request)
	// --------------------------------------------
	// --------------------------------------------

	// request error
	if responseErr != nil {
		switch responseErr.(type) {
		case net.Error:
			if responseErr.(net.Error).Timeout() {
				c.statsdTags = append(c.statsdTags, "error:timeout")
			}
		case *net.OpError:
			if responseErr.(*net.OpError).Op == "read" {
				c.statsdTags = append(c.statsdTags, "error:connection_refused")
			} else if responseErr.(*net.OpError).Op == "dial" {
				c.statsdTags = append(c.statsdTags, "error:unknown_host")
			}
		case syscall.Errno:
			if responseErr.(syscall.Errno) == syscall.ECONNREFUSED {
				c.statsdTags = append(c.statsdTags, "error:connection_refused")
			}
		}

		return c.failAfterRequest(responseErr)
	}

	// set status code and error response flag
	c.statusCode = response.StatusCode
	c.responseIsError = c.statusCode < http.StatusOK || c.statusCode >= http.StatusMultipleChoices

	// defer response body reader close
	defer closeResponse(response, c.logger)

	// get response body
	// ReadAll is called previously and would throw an error in http.Client.Do
	body, _ := ioutil.ReadAll(response.Body)

	// process response
	if processResponseErr := c.processResponseData(body, response.Header.Get(contentTypeHeader)); processResponseErr != nil {
		return c.failAfterRequest(processResponseErr)
	}

	// only keep the raw response if explicitly requested
	if c.keepRawResponse {
		c.rawresponse = body
	}

	c.logger.WithFields(map[string]interface{}{
		"type": NAME,
	}).Debugf("%s request to %s returned code %d", c.method, c.endpoint.Host, c.statusCode)

	return c.statusCode, nil
}

// close the http response
func closeResponse(resp *http.Response, logger log.Logger) {
	if closeErr := resp.Body.Close(); closeErr != nil {
		logger.WithFields(map[string]interface{}{
			"error_message": closeErr.Error(),
			"type":          NAME,
		}).Error("unable to close response body")
	}
}

// endregion

// region EXPORTED FUNCS

// Do will prepare the request and either run it directly
// or from within a circuit breaker
func (c *Client) Do(ctx context.Context, method string, payload interface{}) (int, error) {
	if c.logger == nil {
		c.logger = log.NewNoop()
	}

	c.method = method
	if c.endpoint == nil {
		err := errors.New("endpoint for request not set")
		c.logger.WithFields(map[string]interface{}{
			"error_message": err.Error(),
			"type":          NAME,
		}).Error("config error")
		c.internalError = true

		return http.StatusInternalServerError, err
	}

	if c.cb == nil {
		return c.doInternal(ctx, payload)
	}

	sc, err := c.cb.Execute(func() (interface{}, error) {
		return c.doInternal(ctx, payload)
	})

	// although doInternal will always return a status code,
	// the circuit breaker may be open or half open, which
	// could result in a nil value here
	if sc == nil {
		c.logger.WithFields(map[string]interface{}{
			"error_message": "circuit breaker open or half-open",
			"type":          NAME,
		}).Warn("request blocked")
		sc = http.StatusFailedDependency
	}

	return sc.(int), err
}

// KeepRawResponse will cause the raw bytes from the http response
// to be retained
func (c *Client) KeepRawResponse() {
	c.keepRawResponse = true
}

// RawResponse is a shortcut to access the raw bytes returned
// in the http response
func (c *Client) RawResponse() []byte {
	return c.rawresponse
}

// SetTimeoutMS sets the maximum number of milliseconds allowed for
// a request to complete.  The default request timeout is 8 seconds (8000 ms)
func (c *Client) SetTimeoutMS(timeout int) {
	if timeout < 0 {
		timeout = 0
	}

	c.client.Timeout = time.Duration(timeout) * time.Millisecond
}

// SetLogger will set the client's internal logger.
// If no logger is set, a no-op logger will be used
func (c *Client) SetLogger(logger log.Logger) {
	c.logger = logger
}

// StatusCodeIsError is a shortcut to determine if the status code is
// considered an error
func (c *Client) StatusCodeIsError() bool {
	return c.responseIsError
}

// WillSaturate assigns the interface that will be saturated
// when the request succeeds.  It is assumed that the value passed
// into this function can be saturated via the unmarshalling of json.
// If that is not the case, you will need to process the raw bytes
// returned in the response instead
func (c *Client) WillSaturate(proto interface{}) {
	c.prototype = proto
}

// WillSaturateOnError assigns the interface that will be saturated
// when the request fails.  It is assumed that the value passed
// into this function can be saturated via the unmarshalling of json.
// If that is not the case, you will need to process the raw bytes
// returned in the response instead.  This library treats an error
// as any response with a status code not in the 2XX range.
func (c *Client) WillSaturateOnError(proto interface{}) {
	c.errorPrototype = proto
}

// WillSaturateWithStatusCode assigns the interface that will be
// saturated when a specific response code is encountered.
// This overrides the value of WillSaturate or WillSaturateOnError
// for the same code.  For example, if a value is passed into this
// function that should saturate on a 200 response code, that will
// take precedence over anything set in WillSaturate, but will only
// return the saturated value for a 200, and no other 2XX-level code,
// unless specified here.
func (c *Client) WillSaturateWithStatusCode(statusCode int, proto interface{}) {
	if c.customPrototypes == nil {
		c.customPrototypes = make(map[int]interface{}, 1)
	}

	c.customPrototypes[statusCode] = proto
}

// SetCircuitBreaker sets the optional circuit breaker interface that
// wraps the http request.
func (c *Client) SetCircuitBreaker(cb CircuitBreakerPrototype) {
	c.cb = cb
}

// SetStatsdDelegate will set the statsd client, the stat, and tags
func (c *Client) SetStatsdDelegate(sdClient StatsdClientPrototype, stat string, tags []string) {
	c.statsdClient = sdClient
	c.statsdTags = tags
	c.statsdStat = stat
}

// SetContentType will set the request content type.  By default, all
// requests are of type application/json.  If you wish to use a
// different type, here is where you override it.  Also note that if
// you do provide a content type, your payload for POST, PUT, or PATCH
// must be a byte slice or it must be convertible to a byte slice
func (c *Client) SetContentType(ct string) {
	c.headers[contentTypeHeader] = ct

	if ct != jsonType {
		delete(c.headers, acceptHeader)
	} else {
		c.headers[acceptHeader] = jsonType
	}
}

// SetHeader allows for custom http headers
func (c *Client) SetHeader(key string, value string) {
	if key == contentTypeHeader {
		c.SetContentType(value)
		return
	}

	c.headers[key] = value
}

// Duration will return the elapsed time of the request in an
// int64 nanosecond count
func (c *Client) Duration() time.Duration {
	return c.duration
}

//
// Convenience Functions
// ========================================================
//

// Get performs an HTTP GET request
func (c *Client) Get(ctx context.Context) (int, error) {
	return c.Do(ctx, http.MethodGet, nil)
}

// Post performs an HTTP POST request with the specified payload
func (c *Client) Post(ctx context.Context, payload interface{}) (int, error) {
	c.method = http.MethodPost

	return c.Do(ctx, http.MethodPost, payload)
}

// Put performs an HTTP PUT request with the specified payload
func (c *Client) Put(ctx context.Context, payload interface{}) (int, error) {
	c.method = http.MethodPut

	return c.Do(ctx, http.MethodPut, payload)
}

// Patch performs an HTTP PATCH request with the specified payload
func (c *Client) Patch(ctx context.Context, payload interface{}) (int, error) {
	c.method = http.MethodPatch

	return c.Do(ctx, http.MethodPatch, payload)
}

// Delete performs an HTTP DELETE request
func (c *Client) Delete(ctx context.Context, payload interface{}) (int, error) {
	c.method = http.MethodDelete

	return c.Do(ctx, http.MethodDelete, payload)
}

// endregion
