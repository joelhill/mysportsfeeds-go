package blaster

import (
	"context"
	"time"
)

//go:generate counterfeiter -o ./fakes/fake_circuitbreaker_prototype.go . CircuitBreakerPrototype
//go:generate counterfeiter -o ./fakes/fake_statsd_client_prototype.go . StatsdClientPrototype
//go:generate counterfeiter -o ./fakes/fake_client.go . IClient

// CircuitBreakerPrototype defines the circuit breaker Execute function signature
type CircuitBreakerPrototype interface {
	Execute(func() (interface{}, error)) (interface{}, error)
}

// StatsdClientPrototype defines the statsd client functions used in this library
type StatsdClientPrototype interface {
	Incr(name string, tags []string, rate float64) error
	Timing(name string, value time.Duration, tags []string, rate float64) error
}

// IClient - interface for the cb api client
type IClient interface {
	Delete(ctx context.Context, payload interface{}) (int, error)
	Duration() time.Duration
	Do(ctx context.Context, method string, payload interface{}) (int, error)
	Get(ctx context.Context) (int, error)
	KeepRawResponse()
	Post(ctx context.Context, payload interface{}) (int, error)
	Put(ctx context.Context, payload interface{}) (int, error)
	Patch(ctx context.Context, payload interface{}) (int, error)
	RawResponse() []byte
	SetCircuitBreaker(cb CircuitBreakerPrototype)
	SetStatsdDelegate(sdClient StatsdClientPrototype, stat string, tags []string)
	SetContentType(ct string)
	SetHeader(key string, value string)
	SetNRTxnName(name string)
	SetTimeoutMS(timeout time.Duration)
	StatusCodeIsError() bool
	WillSaturate(proto interface{})
	WillSaturateOnError(proto interface{})
	WillSaturateWithStatusCode(statusCode int, proto interface{})
}
