// This is a generated file - DO NOT EDIT
// +build gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	gdk "github.com/pekim/gobbi/lib/gdk"
	"sync"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
/*

	void colorchooser_colorActivatedHandler(GObject *, GdkRGBA *, gpointer);

	static gulong ColorChooser_signal_connect_color_activated(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "color-activated", G_CALLBACK(colorchooser_colorActivatedHandler), data);
	}

*/
import "C"

// Blacklisted : gtk_actionable_get_action_name

// Blacklisted : gtk_actionable_get_action_target_value

// Blacklisted : gtk_actionable_set_action_name

// Blacklisted : gtk_actionable_set_action_target

// Blacklisted : gtk_actionable_set_action_target_value

// Blacklisted : gtk_actionable_set_detailed_action_name

type signalColorChooserColorActivatedDetail struct {
	callback  ColorChooserSignalColorActivatedCallback
	handlerID C.gulong
}

var signalColorChooserColorActivatedId int
var signalColorChooserColorActivatedMap = make(map[int]signalColorChooserColorActivatedDetail)
var signalColorChooserColorActivatedLock sync.RWMutex

// ColorChooserSignalColorActivatedCallback is a callback function for a 'color-activated' signal emitted from a ColorChooser.
type ColorChooserSignalColorActivatedCallback func(color *gdk.RGBA)

/*
ConnectColorActivated connects the callback to the 'color-activated' signal for the ColorChooser.

The returned value represents the connection, and may be passed to DisconnectColorActivated to remove it.
*/
func (recv *ColorChooser) ConnectColorActivated(callback ColorChooserSignalColorActivatedCallback) int {
	signalColorChooserColorActivatedLock.Lock()
	defer signalColorChooserColorActivatedLock.Unlock()

	signalColorChooserColorActivatedId++
	instance := C.gpointer(recv.native)
	handlerID := C.ColorChooser_signal_connect_color_activated(instance, C.gpointer(uintptr(signalColorChooserColorActivatedId)))

	detail := signalColorChooserColorActivatedDetail{callback, handlerID}
	signalColorChooserColorActivatedMap[signalColorChooserColorActivatedId] = detail

	return signalColorChooserColorActivatedId
}

/*
DisconnectColorActivated disconnects a callback from the 'color-activated' signal for the ColorChooser.

The connectionID should be a value returned from a call to ConnectColorActivated.
*/
func (recv *ColorChooser) DisconnectColorActivated(connectionID int) {
	signalColorChooserColorActivatedLock.Lock()
	defer signalColorChooserColorActivatedLock.Unlock()

	detail, exists := signalColorChooserColorActivatedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalColorChooserColorActivatedMap, connectionID)
}

//export colorchooser_colorActivatedHandler
func colorchooser_colorActivatedHandler(_ *C.GObject, c_color *C.GdkRGBA, data C.gpointer) {
	signalColorChooserColorActivatedLock.RLock()
	defer signalColorChooserColorActivatedLock.RUnlock()

	color := gdk.RGBANewFromC(unsafe.Pointer(c_color))

	index := int(uintptr(data))
	callback := signalColorChooserColorActivatedMap[index].callback
	callback(color)
}

// Unsupported : gtk_color_chooser_add_palette : unsupported parameter colors :

// Blacklisted : gtk_color_chooser_get_rgba

// Blacklisted : gtk_color_chooser_get_use_alpha

// Blacklisted : gtk_color_chooser_set_rgba

// Blacklisted : gtk_color_chooser_set_use_alpha
