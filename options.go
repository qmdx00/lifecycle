package lifecycle

import (
	"context"
	"os"
	"syscall"
	"time"
)

// Option application option
type Option func(*options)

// options application option struct
type options struct {
	id       string
	name     string
	version  string
	metadata map[string]string

	// sigs signals for application shutdown
	sigs        []os.Signal
	ctx         context.Context
	stopTimeout time.Duration
}

// defaultOptions default options
var defaultOptions = options{
	metadata:    make(map[string]string),
	sigs:        []os.Signal{syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT},
	ctx:         context.Background(),
	stopTimeout: time.Second,
}

// WithID with application id
func WithID(id string) Option {
	return func(o *options) { o.id = id }
}

// WithName with application name
func WithName(name string) Option {
	return func(o *options) { o.name = name }
}

// WithVersion with application version
func WithVersion(version string) Option {
	return func(o *options) { o.version = version }
}

// WithMetadata with application metadata
func WithMetadata(md map[string]string) Option {
	return func(o *options) { o.metadata = md }
}

// WithSignal with exit signals
func WithSignal(sigs ...os.Signal) Option {
	return func(o *options) { o.sigs = sigs }
}

// WithContext with service context
func WithContext(ctx context.Context) Option {
	return func(o *options) { o.ctx = ctx }
}

// WithStopTimeout with app stop timeout.
func WithStopTimeout(timeout time.Duration) Option {
	return func(o *options) { o.stopTimeout = timeout }
}
