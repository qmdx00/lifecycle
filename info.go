package lifecycle

import "context"

// AppInfo application context value
type AppInfo interface {
	// ID app id
	ID() string
	// Name app name
	Name() string
	// Version app version
	Version() string
	// Metadata app stored metadata
	Metadata() map[string]string
}

// appKey app context key
type appKey struct{}

// NewContext returns a new Context that carries AppInfo
func NewContext(ctx context.Context, s AppInfo) context.Context {
	return context.WithValue(ctx, appKey{}, s)
}

// FromContext returns the AppInfo value stored in ctx, if any
func FromContext(ctx context.Context) (s AppInfo, ok bool) {
	s, ok = ctx.Value(appKey{}).(AppInfo)
	return
}
