package lifecycle

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

type info struct {
	id       string
	name     string
	version  string
	metadata map[string]string
}

func (i *info) ID() string {
	return i.id
}

func (i *info) Name() string {
	return i.name
}

func (i *info) Version() string {
	return i.version
}

func (i *info) Metadata() map[string]string {
	return i.metadata
}

func TestNewContext(t *testing.T) {
	ctx := context.Background()
	assert.NotEmpty(t, NewContext(ctx, &info{}))
}

func TestFromContext(t *testing.T) {
	foo := &info{name: "test"}
	ctx := NewContext(context.Background(), foo)

	bar, ok := FromContext(ctx)
	if !ok {
		t.Fatal()
	}

	assert.Equal(t, foo, bar)
}
