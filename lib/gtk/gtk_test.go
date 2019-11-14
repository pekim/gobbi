package gtk

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCleanBuild(t *testing.T) {
	assert.Equal(t, int32(3), MAJOR_VERSION)
	assert.Equal(t, Align(2), Align_End)
}

func TestFunctionCall(t *testing.T) {
	//gtk.Init()
	//gtk.MainQuit()
}
