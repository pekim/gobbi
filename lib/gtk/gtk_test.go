package gtk

import (
	"testing"

	"github.com/stretchr/testify/assert"
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

func TestInit(t *testing.T) {
	argsIn := []string{"one", "--g-fatal-warnings", "two", "--class", "cls", "three"}

	argc, argsOut := Init(argsIn)
	assert.Equal(t, int32(3), argc)
	assert.Equal(t, []string{"one", "two", "three"}, argsOut)
}

func TestClass(t *testing.T) {
	Init([]string{})

	label := LabelNew("test")
	assert.Equal(t, "test", label.GetText())

	label.SetXalign(0.3)
	assert.Equal(t, float32(0.3), label.GetXalign())
}

func TestSignalConnect(t *testing.T) {
	Init([]string{})

	callbackCalled := false

	button := ButtonNew()
	button.SetLabel("qaz")
	button.Widget().ConnectShow(func(widget *Widget) {
		assert.True(t, widget.Object().Equals(button.Object()))
		callbackCalled = true
	})

	button.Widget().Show()

	assert.True(t, callbackCalled)
}

func TestSignalDisconnect(t *testing.T) {
	Init([]string{})

	calledCount := 0

	button := ButtonNew()
	var connectShowId int
	connectShowId = button.Widget().ConnectShow(func(widget *Widget) {
		calledCount++

		button.Widget().DisconnectSignal(connectShowId)
	})

	button.Widget().Show()
	button.Widget().Hide()
	button.Widget().Show()

	assert.Equal(t, 1, calledCount)
}

func TestSignalMultipleHandlers(t *testing.T) {
	Init([]string{})

	calledCount := 0

	button := ButtonNew()

	button.Widget().ConnectShow(func(widget *Widget) {
		calledCount++
	})

	button.Widget().ConnectShow(func(widget *Widget) {
		calledCount++
	})

	button.Widget().Show()

	assert.Equal(t, 2, calledCount)
}
