package atk

import (
	"github.com/pekim/gobbi/lib/gtk"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCleanBuild(t *testing.T) {
	assert.Equal(t, int32(3), gtk.MAJOR_VERSION)
	assert.Equal(t, gtk.Align(2), gtk.Align_End)
}

func TestFunctionCall(t *testing.T) {
	//gtk.Init()
	//gtk.MainQuit()
}
