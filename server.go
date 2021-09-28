package lifecycle

import "context"

// Server ...
type Server interface {
	// Run func for server start
	Run(ctx context.Context) error
	// Stop func for server shutdown
	Stop(ctx context.Context) error
}
