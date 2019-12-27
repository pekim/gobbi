package signal

import (
	"github.com/pekim/gobbi/lib/internal/c/gobject"
	"github.com/pekim/gobbi/lib/internal/c/gtk"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"unsafe"
)

func TestSignalConnect(t *testing.T) {
	argv := os.Args
	var argc int = len(argv)
	gtk.Fn_gtk_init(&argc, &argv)

	callbackCalled := false

	button := gtk.Fn_gtk_button_new()
	gtk.Fn_gtk_button_set_label(button, "qaz")
	ConnectSignal(button, "show", func(return_value *Value, paramValues []Value) {
		assert.Equal(t, button, gobject.Fn_g_value_get_object(unsafe.Pointer(&paramValues[0])))
		callbackCalled = true
	})

	gtk.Fn_gtk_widget_show(button)

	assert.True(t, callbackCalled)
}

func TestSignalDisconnect(t *testing.T) {
	argv := os.Args
	var argc int = len(argv)
	gtk.Fn_gtk_init(&argc, &argv)

	calledCount := 0
	button := gtk.Fn_gtk_button_new()

	var connectShowId int
	connectShowId = ConnectSignal(button, "show", func(return_value *Value, paramValues []Value) {
		calledCount++

		DisconnectSignal(connectShowId)
	})

	gtk.Fn_gtk_widget_show(button)
	gtk.Fn_gtk_widget_hide(button)
	gtk.Fn_gtk_widget_show(button)

	assert.Equal(t, 1, calledCount)
}

func TestSignalMultipleHandlers(t *testing.T) {
	argv := os.Args
	var argc int = len(argv)
	gtk.Fn_gtk_init(&argc, &argv)

	calledCount := 0
	button := gtk.Fn_gtk_button_new()

	ConnectSignal(button, "show", func(return_value *Value, paramValues []Value) {
		calledCount++
	})

	ConnectSignal(button, "show", func(return_value *Value, paramValues []Value) {
		calledCount++
	})

	gtk.Fn_gtk_widget_show(button)

	assert.Equal(t, 2, calledCount)
}
