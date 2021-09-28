package lifecycle

import (
	"context"
	"errors"
	"golang.org/x/sync/errgroup"
	"os"
	"os/signal"
	"sync"
)

// App application struct
type App struct {
	sync.Mutex

	opts   options
	ctx    context.Context
	cancel func()

	cleanups []cleanup
	servers  map[string]Server
}

// cleanup ...
type cleanup func()

// ID ...
func (a *App) ID() string {
	return a.opts.id
}

// Name ...
func (a *App) Name() string {
	return a.opts.name
}

// Version ...
func (a *App) Version() string {
	return a.opts.version
}

// Metadata ...
func (a *App) Metadata() map[string]string {
	return a.opts.metadata
}

// NewApp ...
func NewApp(opts ...Option) *App {
	// context with cancel
	ctx, cancel := context.WithCancel(context.Background())

	// override default options
	_opts := defaultOptions
	for _, o := range opts {
		o(&_opts)
	}

	return &App{
		Mutex:    sync.Mutex{},
		opts:     _opts,
		ctx:      ctx,
		cancel:   cancel,
		servers:  make(map[string]Server),
		cleanups: make([]cleanup, 0),
	}
}

// Attach add named server ...
func (a *App) Attach(name string, server Server) {
	a.Lock()
	defer a.Unlock()
	a.servers[name] = server
}

// Cleanup do cleanup functions
func (a *App) Cleanup(f ...cleanup) {
	a.Lock()
	defer a.Unlock()
	a.cleanups = append(a.cleanups, f...)
}

// Run start application ...
func (a *App) Run() (_ error) {
	var ctx context.Context
	var group *errgroup.Group
	var wg = &sync.WaitGroup{}

	appCtx := NewContext(a.ctx, a)
	group, ctx = errgroup.WithContext(appCtx)

	// start servers
	for _, server := range a.servers {
		srv := server

		group.Go(func() error {
			<-ctx.Done()
			return srv.Stop(ctx)
		})

		wg.Add(1)
		group.Go(func() error {
			wg.Done()
			return srv.Run(ctx)
		})
	}

	// waiting for all servers start
	wg.Wait()

	// wait for os signal
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, a.opts.sigs...)

	group.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-sigs:
				return a.Shutdown()
			}
		}
	})

	if err := group.Wait(); err != nil && !errors.Is(err, context.Canceled) {
		return err
	}
	return
}

// Shutdown gracefully stops the application
func (a *App) Shutdown() (_ error) {
	// cancel context
	if a.cancel != nil {
		a.cancel()
	}
	// do cleanups
	for _, _cleanup := range a.cleanups {
		_cleanup()
	}
	return
}
