package gtk

import (
	"github.com/pekim/gobbi/lib/gio"
	"github.com/pekim/gobbi/lib/gobject"
	"github.com/pekim/gobbi/lib/gtk"
	"github.com/stretchr/testify/assert"
	"testing"
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

func TestCreateWindow(t *testing.T) {
	gtk.Init([]string{})

	window := gtk.WindowNew(gtk.GTK_WINDOW_TOPLEVEL)
	assert.NotNil(t, window)
	window.Widget().Destroy()
}

func TestIndirectUpcasting(t *testing.T) {
	gtk.Init([]string{})

	window := gtk.WindowNew(gtk.GTK_WINDOW_TOPLEVEL)
	widget := window.Bin().Container().Widget()

	assert.False(t, widget.IsVisible())
	widget.ShowAll()
	assert.True(t, widget.IsVisible())

	window.Widget().Destroy()
}

func TestDirectUpcasting(t *testing.T) {
	gtk.Init([]string{})

	window := gtk.WindowNew(gtk.GTK_WINDOW_TOPLEVEL)
	widget := window.Widget()

	assert.False(t, widget.IsVisible())
	widget.ShowAll()
	assert.True(t, widget.IsVisible())

	window.Widget().Destroy()
}

func TestDowncasting(t *testing.T) {
	gtk.Init([]string{})

	window := gtk.WindowNew(gtk.GTK_WINDOW_TOPLEVEL)
	window.SetTitle("qaz")

	widget := window.Widget()
	window2 := gtk.CastToWindow(widget.Object())
	assert.Equal(t, "qaz", window2.GetTitle())

	window.Widget().Destroy()
}

func TestArrayParam(t *testing.T) {
	gtk.Init([]string{})

	store := gtk.ListStoreNewv([]gobject.Type{gobject.TYPE_BOOLEAN, gobject.TYPE_INT})
	assert.NotNil(t, store)
}

func TestSignalConnect(t *testing.T) {
	callbackCalled := false

	button := gtk.ButtonNew()
	button.Widget().ConnectShow(func(widget *gtk.Widget) {
		assert.True(t, widget.Object().Equals(button.Object()))
		callbackCalled = true
	})
	button.Widget().Show()

	assert.True(t, callbackCalled)
}

// The ConnectNotify & ConnectNotifyProperty functions are from the gobject package.
// But there's no notify signal in that package that is suitable for this test.
// So it's much easier to test it in this package.
func TestConnectNotify(t *testing.T) {
	notifyCount := 0
	label := gtk.LabelNew("a")

	label.Object().ConnectNotify(func(_ *gobject.Object, pspec *gobject.ParamSpec) {
		notifyCount++
	})

	label.Object().ConnectNotifyProperty("label", func(_ *gobject.Object, pspec *gobject.ParamSpec) {
		notifyCount++
	})

	label.Object().ConnectNotifyProperty("non-such-property", func(_ *gobject.Object, pspec *gobject.ParamSpec) {
		notifyCount++
	})

	label.SetText("b")

	assert.Equal(t, 2, notifyCount)
}

// The BindProperty function is from the gobject package.
// But there's a paucity of easy to use signals in that package.
// So it's much easier to test it in this package.
func TestObjectBindProperty(t *testing.T) {
	label1 := gtk.LabelNew("1")
	label2 := gtk.LabelNew("2")

	label1.Object().BindProperty("label", label2.Object(), "label", gobject.BINDING_DEFAULT)

	assert.Equal(t, "2", label2.GetText())
	label1.SetText("1 changed")
	assert.Equal(t, "1 changed", label2.GetText())
}
