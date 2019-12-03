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

	v := GetMajorVersion()
	assert.Equal(t, uint32(3), v)
}

func TestClass(t *testing.T) {
	Init()

	label := LabelNew("test")
	assert.Equal(t, "test", label.GetText())

	label.SetXalign(0.3)
	assert.Equal(t, float32(0.3), label.GetXalign())
}
