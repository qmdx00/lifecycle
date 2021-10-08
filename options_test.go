package lifecycle

import (
	"context"
	"github.com/stretchr/testify/assert"
	"os"
	"syscall"
	"testing"
)

func TestWithID(t *testing.T) {
	id := "test"
	opts := options{}
	WithID(id)(&opts)
	assert.Equal(t, id, opts.id)
}

func TestWithName(t *testing.T) {
	name := "test"
	opts := options{}
	WithName(name)(&opts)
	assert.Equal(t, name, opts.name)
}

func TestWithVersion(t *testing.T) {
	version := "v1.0"
	opts := options{}
	WithVersion(version)(&opts)
	assert.Equal(t, version, opts.version)
}

func TestWithMetadata(t *testing.T) {
	meta := map[string]string{
		"test": "test",
	}
	opts := options{}
	WithMetadata(meta)(&opts)
	assert.Equal(t, meta, opts.metadata)
}

func TestWithSignal(t *testing.T) {
	sigs := []os.Signal{syscall.SIGTERM}
	opts := options{}
	WithSignal(sigs...)(&opts)
	assert.Equal(t, sigs, opts.sigs)
}

func TestWithContext(t *testing.T) {
	ctx := context.Background()
	opts := options{}
	WithContext(ctx)(&opts)
	assert.Equal(t, ctx, opts.ctx)
}
