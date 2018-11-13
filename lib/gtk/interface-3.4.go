// This is a generated file - DO NOT EDIT
// +build gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	gdk "github.com/pekim/gobbi/lib/gdk"
	"sync"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
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

// Gets the action name for @actionable.
//
// See gtk_actionable_set_action_name() for more information.
/*

C function : gtk_actionable_get_action_name
*/
func (recv *Actionable) GetActionName() string {
	retC := C.gtk_actionable_get_action_name((*C.GtkActionable)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : gtk_actionable_get_action_target_value : return type : Blacklisted record : GVariant

// Specifies the name of the action with which this widget should be
// associated.  If @action_name is %NULL then the widget will be
// unassociated from any previous action.
//
// Usually this function is used when the widget is located (or will be
// located) within the hierarchy of a #GtkApplicationWindow.
//
// Names are of the form “win.save” or “app.quit” for actions on the
// containing #GtkApplicationWindow or its associated #GtkApplication,
// respectively.  This is the same form used for actions in the #GMenu
// associated with the window.
/*

C function : gtk_actionable_set_action_name
*/
func (recv *Actionable) SetActionName(actionName string) {
	c_action_name := C.CString(actionName)
	defer C.free(unsafe.Pointer(c_action_name))

	C.gtk_actionable_set_action_name((*C.GtkActionable)(recv.native), c_action_name)

	return
}

// Unsupported : gtk_actionable_set_action_target : unsupported parameter ... : varargs

// Unsupported : gtk_actionable_set_action_target_value : unsupported parameter target_value : Blacklisted record : GVariant

// Sets the action-name and associated string target value of an
// actionable widget.
//
// @detailed_action_name is a string in the format accepted by
// g_action_parse_detailed_name().
//
// (Note that prior to version 3.22.25,
// this function is only usable for actions with a simple "s" target, and
// @detailed_action_name must be of the form `"action::target"` where
// `action` is the action name and `target` is the string to use
// as the target.)
/*

C function : gtk_actionable_set_detailed_action_name
*/
func (recv *Actionable) SetDetailedActionName(detailedActionName string) {
	c_detailed_action_name := C.CString(detailedActionName)
	defer C.free(unsafe.Pointer(c_detailed_action_name))

	C.gtk_actionable_set_detailed_action_name((*C.GtkActionable)(recv.native), c_detailed_action_name)

	return
}

type signalColorChooserColorActivatedDetail struct {
	callback  ColorChooserSignalColorActivatedCallback
	handlerID C.gulong
}

var signalColorChooserColorActivatedId int
var signalColorChooserColorActivatedMap = make(map[int]signalColorChooserColorActivatedDetail)
var signalColorChooserColorActivatedLock sync.Mutex

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
	color := gdk.RGBANewFromC(unsafe.Pointer(c_color))

	index := int(uintptr(data))
	callback := signalColorChooserColorActivatedMap[index].callback
	callback(color)
}

// Unsupported : gtk_color_chooser_add_palette : unsupported parameter colors :

// Gets the currently-selected color.
/*

C function : gtk_color_chooser_get_rgba
*/
func (recv *ColorChooser) GetRgba() *gdk.RGBA {
	var c_color C.GdkRGBA

	C.gtk_color_chooser_get_rgba((*C.GtkColorChooser)(recv.native), &c_color)

	color := gdk.RGBANewFromC(unsafe.Pointer(&c_color))

	return color
}

// Returns whether the color chooser shows the alpha channel.
/*

C function : gtk_color_chooser_get_use_alpha
*/
func (recv *ColorChooser) GetUseAlpha() bool {
	retC := C.gtk_color_chooser_get_use_alpha((*C.GtkColorChooser)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Sets the color.
/*

C function : gtk_color_chooser_set_rgba
*/
func (recv *ColorChooser) SetRgba(color *gdk.RGBA) {
	c_color := (*C.GdkRGBA)(C.NULL)
	if color != nil {
		c_color = (*C.GdkRGBA)(color.ToC())
	}

	C.gtk_color_chooser_set_rgba((*C.GtkColorChooser)(recv.native), c_color)

	return
}

// Sets whether or not the color chooser should use the alpha channel.
/*

C function : gtk_color_chooser_set_use_alpha
*/
func (recv *ColorChooser) SetUseAlpha(useAlpha bool) {
	c_use_alpha :=
		boolToGboolean(useAlpha)

	C.gtk_color_chooser_set_use_alpha((*C.GtkColorChooser)(recv.native), c_use_alpha)

	return
}
