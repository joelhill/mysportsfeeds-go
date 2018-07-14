package blaster

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"os"
	"sync"
	"time"

	"github.com/opentracing/opentracing-go"
)

// Defaults is a container for setting package level values
type Defaults struct {
	// ServiceName is the name of the calling service
	ServiceName string

	// TracerProviderFunc is a function that provides
	// the opentracing.Tracer for tracing HTTP requests
	TracerProviderFunc func(ctx context.Context, operationName string, r *http.Request) (*http.Request, opentracing.Span)

	// RequestIDProviderFunc is a function that provides the
	// parent Request id used in tracing the caller's Request.
	RequestIDProviderFunc func(ctx context.Context) (string, bool)

	// RequestSourceProviderFunc is a function that provides
	// the Request-Source header
	RequestSourceProviderFunc func(ctx context.Context) (string, bool)

	// UserAgent is a package-level user agent header used for
	// each outgoing request
	UserAgent string

	// RequireHeaders will cancel any request and return an error if any of the following
	// headers are missing:
	// 		Request-ID
	// 		Request-Source
	// 		Calling-Service
	RequireHeaders bool

	// StatsdRate is the statsd reporting rate
	StatsdRate float64
}

var (
	pkgServiceName               string
	pkgUserAgent                 string
	pkgTracerProviderFunc        func(ctx context.Context, operationName string, r *http.Request) (*http.Request, opentracing.Span)
	pkgRequestIDProviderFunc     func(cxt context.Context) (string, bool)
	pkgRequestSourceProviderFunc func(cxt context.Context) (string, bool)
	pkgOnce                      sync.Once
	pkgRequireHeaders            bool
	pkgStatsdRate                float64

	envHTTPMocking = "MOCKING_HTTP"
)

//
// Package Level Functions
// ========================================================
//

// ensurePackageVariables makes sure that the package level
// variables are set.  This function runs once, then no-ops
// on subsequent calls
func ensurePackageVariables() {
	pkgOnce.Do(func() {

		// we need something to be set as a service name
		if pkgServiceName == "" {
			// if caller didnt set it, look in env
			pkgServiceName = os.Getenv("SERVICE_NAME")
			if pkgServiceName == "" {
				// if not in env, just use the hostname
				pkgServiceName = os.Getenv("HOSTNAME")
			}
		}

		// user agent is service name + hostname
		if pkgUserAgent == "" {
			if pkgServiceName == os.Getenv("HOSTNAME") {
				pkgUserAgent = pkgServiceName
			} else {
				pkgUserAgent = fmt.Sprintf("%s-%s", pkgServiceName, os.Getenv("HOSTNAME"))
			}
		}
	})
}

// SetDefaults will apply package-level default values to
// be used on all requests
func SetDefaults(defaults *Defaults) {
	pkgServiceName = defaults.ServiceName
	pkgRequestIDProviderFunc = defaults.RequestIDProviderFunc
	pkgRequestSourceProviderFunc = defaults.RequestSourceProviderFunc
	pkgUserAgent = defaults.UserAgent
	pkgRequireHeaders = defaults.RequireHeaders
	pkgStatsdRate = defaults.StatsdRate
	pkgTracerProviderFunc = defaults.TracerProviderFunc
}

// this creates a http client with sensible defaults
func newHTTPClient() *http.Client {
	// all http mocking libraries can override the default http client,
	// but many cannot override clients that have been tuned with custom
	// transports.  If this env var is set, we return the standard
	// http client.
	if os.Getenv(envHTTPMocking) != "" {
		return http.DefaultClient
	}

	client := &http.Client{
		Timeout: requestTimeout,
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout:   sockTimeout,
				DualStack: true,
				KeepAlive: keepAlive,
			}).DialContext,
			MaxIdleConnsPerHost:   maxIdleConnsPerHost,
			MaxIdleConns:          maxIdleConns,
			IdleConnTimeout:       idleTimeout,
			TLSHandshakeTimeout:   tlsTimeout,
			ExpectContinueTimeout: 1 * time.Second,
		},
	}

	return client
}

// NewClient will initialize and return a new client with a
// request and endpoint.  The client's content type defaults
// to application/json
// Deprecated: Use New(opts ClientOptions) instead
func NewClient(uri string) (*Client, error) {

	ensurePackageVariables()

	ep, err := url.ParseRequestURI(uri)
	if err != nil {
		return nil, err
	}

	c := &Client{
		endpoint: ep,
		method:   http.MethodGet,
		client:   newHTTPClient(),
		headers: map[string]string{
			userAgentHeader:      pkgUserAgent,
			contentTypeHeader:    jsonType,
			callingServiceHeader: pkgServiceName,
			acceptHeader:         jsonType,
		},
	}

	return c, nil
}

func New(opts ClientOptions) (*Client, error) {
	ensurePackageVariables()

	ep, err := url.ParseRequestURI(opts.Endpoint)
	if err != nil {
		return nil, err
	}

	c := &Client{
		endpoint: ep,
		method:   http.MethodGet,
		client:   newHTTPClient(),
		headers: map[string]string{
			userAgentHeader:      pkgUserAgent,
			contentTypeHeader:    jsonType,
			callingServiceHeader: pkgServiceName,
			acceptHeader:         jsonType,
		},
	}

	if opts.Headers != nil {
		for k, v := range opts.Headers {
			c.headers[k] = v
		}
	}
	c.routeMask = opts.RouteMask
	c.calledService = opts.CalledService
	c.prototype = opts.WillSaturate
	c.errorPrototype = opts.WillSaturateOnError
	c.customPrototypes = opts.WillSaturateWithStatusCode
	if opts.TimeoutMS > 0 {
		c.client.Timeout = time.Duration(opts.TimeoutMS) * time.Millisecond
	}
	c.cb = opts.CircuitBreaker
	c.keepRawResponse = opts.KeepRawResponse
	c.logger = opts.Logger

	return c, nil
}

// get a response type string for a specific status code
func responseTypeForStatusCode(statusCode int) string {
	switch {
	case statusCode >= 500:
		return "5xx"
	case statusCode >= 400:
		return "4xx"
	case statusCode >= 300:
		return "3xx"
	case statusCode >= 200:
		return "2xx"
	case statusCode >= 100:
		return "1xx"
	default:
		return "unknown"
	}
}
