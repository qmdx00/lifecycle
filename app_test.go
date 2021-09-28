package lifecycle

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewApp(t *testing.T) {
	app := NewApp()
	time.AfterFunc(time.Second, func() {
		_ = app.Shutdown()
	})
	if err := app.Run(); err != nil {
		t.Fatal(err)
	}
}

func TestApp_ID(t *testing.T) {
	id := "test"
	app := NewApp(WithID(id))
	assert.Equal(t, id, app.ID())
}

func TestApp_Name(t *testing.T) {
	name := "test"
	app := NewApp(WithName(name))
	assert.Equal(t, name, app.Name())
}

func TestApp_Version(t *testing.T) {
	version := "v1.0"
	app := NewApp(WithVersion(version))
	assert.Equal(t, version, app.Version())
}

func TestApp_Metadata(t *testing.T) {
	meta := map[string]string{
		"test": "test",
	}
	app := NewApp(WithMetadata(meta))
	assert.Equal(t, meta, app.Metadata())
}
