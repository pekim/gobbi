package gtk

import (
	"github.com/pekim/gobbi/lib/gtk"
	"github.com/pekim/gobbi/lib/gtksource"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCleanBuild(t *testing.T) {
	gtk.Init([]string{})
	view := gtksource.ViewNew()
	assert.NotNil(t, view)
}
