package gtk

import (
	"github.com/pekim/gobbi/lib/gio"
	"github.com/pekim/gobbi/lib/gtk"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCleanBuild(t *testing.T) {
	app := gtk.ApplicationNew("pekim.gobbi.test", gio.APPLICATION_FLAGS_NONE)
	assert.NotNil(t, app)
}

func TestInit(t *testing.T) {
	argsIn := []string{"one", "--g-fatal-warnings", "two", "--class", "cls", "three"}

	argsOut := gtk.Init(argsIn)
	assert.Equal(t, []string{"one", "two", "three"}, argsOut)
}
