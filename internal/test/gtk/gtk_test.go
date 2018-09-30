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
