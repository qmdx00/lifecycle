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
	opts := options{id: id}
	WithID(id)
	assert.Equal(t, id, opts.id)
}

func TestWithName(t *testing.T) {
	name := "test"
	opts := options{name: name}
	WithName(name)
	assert.Equal(t, name, opts.name)
}

func TestWithVersion(t *testing.T) {
	version := "v1.0"
	opts := options{version: version}
	WithVersion(version)
	assert.Equal(t, version, opts.version)
}

func TestWithMetadata(t *testing.T) {
	meta := map[string]string{
		"test": "test",
	}
	opts := options{metadata: meta}
	WithMetadata(meta)
	assert.Equal(t, meta, opts.metadata)
}

func TestWithSignal(t *testing.T) {
	sigs := []os.Signal{syscall.SIGTERM}
	opts := options{sigs: sigs}
	WithSignal(sigs...)
	assert.Equal(t, sigs, opts.sigs)
}

func TestWithContext(t *testing.T) {
	ctx := context.Background()
	opts := options{ctx: ctx}
	WithContext(ctx)
	assert.Equal(t, ctx, opts.ctx)
}
