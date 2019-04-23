// This is a generated file - DO NOT EDIT
// +build gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	cairo "github.com/pekim/gobbi/lib/cairo"
	gio "github.com/pekim/gobbi/lib/gio"
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

	gboolean widget_drawHandler(GObject *, cairo_t *, gpointer);

	static gulong Widget_signal_connect_draw(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "draw", G_CALLBACK(widget_drawHandler), data);
	}

*/
/*

	void widget_stateFlagsChangedHandler(GObject *, GtkStateFlags, gpointer);

	static gulong Widget_signal_connect_state_flags_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "state-flags-changed", G_CALLBACK(widget_stateFlagsChangedHandler), data);
	}

*/
/*

	void widget_styleUpdatedHandler(GObject *, gpointer);

	static gulong Widget_signal_connect_style_updated(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "style-updated", G_CALLBACK(widget_styleUpdatedHandler), data);
	}

*/
import "C"

// ApplicationNew is a wrapper around the C function gtk_application_new.
func ApplicationNew(applicationId string, flags gio.ApplicationFlags) *Application {
	c_application_id := C.CString(applicationId)
	defer C.free(unsafe.Pointer(c_application_id))

	c_flags := (C.GApplicationFlags)(flags)

	retC := C.gtk_application_new(c_application_id, c_flags)
	retGo := ApplicationNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// Blacklisted : gtk_application_add_window

// Blacklisted : gtk_application_get_windows

// Blacklisted : gtk_application_remove_window

// Blacklisted : gtk_style_has_context

type signalWidgetDrawDetail struct {
	callback  WidgetSignalDrawCallback
	handlerID C.gulong
}

var signalWidgetDrawId int
var signalWidgetDrawMap = make(map[int]signalWidgetDrawDetail)
var signalWidgetDrawLock sync.RWMutex

// WidgetSignalDrawCallback is a callback function for a 'draw' signal emitted from a Widget.
type WidgetSignalDrawCallback func(cr *cairo.Context) bool

/*
ConnectDraw connects the callback to the 'draw' signal for the Widget.

The returned value represents the connection, and may be passed to DisconnectDraw to remove it.
*/
func (recv *Widget) ConnectDraw(callback WidgetSignalDrawCallback) int {
	signalWidgetDrawLock.Lock()
	defer signalWidgetDrawLock.Unlock()

	signalWidgetDrawId++
	instance := C.gpointer(recv.native)
	handlerID := C.Widget_signal_connect_draw(instance, C.gpointer(uintptr(signalWidgetDrawId)))

	detail := signalWidgetDrawDetail{callback, handlerID}
	signalWidgetDrawMap[signalWidgetDrawId] = detail

	return signalWidgetDrawId
}

/*
DisconnectDraw disconnects a callback from the 'draw' signal for the Widget.

The connectionID should be a value returned from a call to ConnectDraw.
*/
func (recv *Widget) DisconnectDraw(connectionID int) {
	signalWidgetDrawLock.Lock()
	defer signalWidgetDrawLock.Unlock()

	detail, exists := signalWidgetDrawMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWidgetDrawMap, connectionID)
}

//export widget_drawHandler
func widget_drawHandler(_ *C.GObject, c_cr *C.cairo_t, data C.gpointer) C.gboolean {
	signalWidgetDrawLock.RLock()
	defer signalWidgetDrawLock.RUnlock()

	cr := cairo.ContextNewFromC(unsafe.Pointer(c_cr))

	index := int(uintptr(data))
	callback := signalWidgetDrawMap[index].callback
	retGo := callback(cr)
	retC :=
		boolToGboolean(retGo)
	return retC
}

type signalWidgetStateFlagsChangedDetail struct {
	callback  WidgetSignalStateFlagsChangedCallback
	handlerID C.gulong
}

var signalWidgetStateFlagsChangedId int
var signalWidgetStateFlagsChangedMap = make(map[int]signalWidgetStateFlagsChangedDetail)
var signalWidgetStateFlagsChangedLock sync.RWMutex

// WidgetSignalStateFlagsChangedCallback is a callback function for a 'state-flags-changed' signal emitted from a Widget.
type WidgetSignalStateFlagsChangedCallback func(flags StateFlags)

/*
ConnectStateFlagsChanged connects the callback to the 'state-flags-changed' signal for the Widget.

The returned value represents the connection, and may be passed to DisconnectStateFlagsChanged to remove it.
*/
func (recv *Widget) ConnectStateFlagsChanged(callback WidgetSignalStateFlagsChangedCallback) int {
	signalWidgetStateFlagsChangedLock.Lock()
	defer signalWidgetStateFlagsChangedLock.Unlock()

	signalWidgetStateFlagsChangedId++
	instance := C.gpointer(recv.native)
	handlerID := C.Widget_signal_connect_state_flags_changed(instance, C.gpointer(uintptr(signalWidgetStateFlagsChangedId)))

	detail := signalWidgetStateFlagsChangedDetail{callback, handlerID}
	signalWidgetStateFlagsChangedMap[signalWidgetStateFlagsChangedId] = detail

	return signalWidgetStateFlagsChangedId
}

/*
DisconnectStateFlagsChanged disconnects a callback from the 'state-flags-changed' signal for the Widget.

The connectionID should be a value returned from a call to ConnectStateFlagsChanged.
*/
func (recv *Widget) DisconnectStateFlagsChanged(connectionID int) {
	signalWidgetStateFlagsChangedLock.Lock()
	defer signalWidgetStateFlagsChangedLock.Unlock()

	detail, exists := signalWidgetStateFlagsChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWidgetStateFlagsChangedMap, connectionID)
}

//export widget_stateFlagsChangedHandler
func widget_stateFlagsChangedHandler(_ *C.GObject, c_flags C.GtkStateFlags, data C.gpointer) {
	signalWidgetStateFlagsChangedLock.RLock()
	defer signalWidgetStateFlagsChangedLock.RUnlock()

	flags := StateFlags(c_flags)

	index := int(uintptr(data))
	callback := signalWidgetStateFlagsChangedMap[index].callback
	callback(flags)
}

type signalWidgetStyleUpdatedDetail struct {
	callback  WidgetSignalStyleUpdatedCallback
	handlerID C.gulong
}

var signalWidgetStyleUpdatedId int
var signalWidgetStyleUpdatedMap = make(map[int]signalWidgetStyleUpdatedDetail)
var signalWidgetStyleUpdatedLock sync.RWMutex

// WidgetSignalStyleUpdatedCallback is a callback function for a 'style-updated' signal emitted from a Widget.
type WidgetSignalStyleUpdatedCallback func()

/*
ConnectStyleUpdated connects the callback to the 'style-updated' signal for the Widget.

The returned value represents the connection, and may be passed to DisconnectStyleUpdated to remove it.
*/
func (recv *Widget) ConnectStyleUpdated(callback WidgetSignalStyleUpdatedCallback) int {
	signalWidgetStyleUpdatedLock.Lock()
	defer signalWidgetStyleUpdatedLock.Unlock()

	signalWidgetStyleUpdatedId++
	instance := C.gpointer(recv.native)
	handlerID := C.Widget_signal_connect_style_updated(instance, C.gpointer(uintptr(signalWidgetStyleUpdatedId)))

	detail := signalWidgetStyleUpdatedDetail{callback, handlerID}
	signalWidgetStyleUpdatedMap[signalWidgetStyleUpdatedId] = detail

	return signalWidgetStyleUpdatedId
}

/*
DisconnectStyleUpdated disconnects a callback from the 'style-updated' signal for the Widget.

The connectionID should be a value returned from a call to ConnectStyleUpdated.
*/
func (recv *Widget) DisconnectStyleUpdated(connectionID int) {
	signalWidgetStyleUpdatedLock.Lock()
	defer signalWidgetStyleUpdatedLock.Unlock()

	detail, exists := signalWidgetStyleUpdatedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWidgetStyleUpdatedMap, connectionID)
}

//export widget_styleUpdatedHandler
func widget_styleUpdatedHandler(_ *C.GObject, data C.gpointer) {
	signalWidgetStyleUpdatedLock.RLock()
	defer signalWidgetStyleUpdatedLock.RUnlock()

	index := int(uintptr(data))
	callback := signalWidgetStyleUpdatedMap[index].callback
	callback()
}

// Blacklisted : gtk_widget_add_device_events

// Blacklisted : gtk_widget_device_is_shadowed

// Blacklisted : gtk_widget_draw

// Blacklisted : gtk_widget_get_device_enabled

// Blacklisted : gtk_widget_get_device_events

// Blacklisted : gtk_widget_get_margin_bottom

// Blacklisted : gtk_widget_get_margin_left

// Blacklisted : gtk_widget_get_margin_right

// Blacklisted : gtk_widget_get_margin_top

// Blacklisted : gtk_widget_get_preferred_height

// Blacklisted : gtk_widget_get_preferred_height_for_width

// Blacklisted : gtk_widget_get_preferred_size

// Blacklisted : gtk_widget_get_preferred_width

// Blacklisted : gtk_widget_get_preferred_width_for_height

// Blacklisted : gtk_widget_get_request_mode

// Blacklisted : gtk_widget_get_state_flags

// Blacklisted : gtk_widget_input_shape_combine_region

// Blacklisted : gtk_widget_override_background_color

// Blacklisted : gtk_widget_override_color

// Blacklisted : gtk_widget_override_cursor

// Blacklisted : gtk_widget_override_font

// Blacklisted : gtk_widget_override_symbolic_color

// Blacklisted : gtk_widget_queue_draw_region

// Blacklisted : gtk_widget_render_icon_pixbuf

// Blacklisted : gtk_widget_reset_style

// Blacklisted : gtk_widget_set_device_enabled

// Blacklisted : gtk_widget_set_device_events

// Blacklisted : gtk_widget_set_margin_bottom

// Blacklisted : gtk_widget_set_margin_left

// Blacklisted : gtk_widget_set_margin_right

// Blacklisted : gtk_widget_set_margin_top

// Blacklisted : gtk_widget_set_state_flags

// Blacklisted : gtk_widget_set_support_multidevice

// Blacklisted : gtk_widget_shape_combine_region

// Blacklisted : gtk_widget_unset_state_flags

// Blacklisted : gtk_window_get_application

// Blacklisted : gtk_window_get_has_resize_grip

// Blacklisted : gtk_window_get_resize_grip_area

// Blacklisted : gtk_window_resize_grip_is_visible

// Blacklisted : gtk_window_resize_to_geometry

// Blacklisted : gtk_window_set_application

// Blacklisted : gtk_window_set_default_geometry

// Blacklisted : gtk_window_set_has_resize_grip

// Blacklisted : gtk_window_set_has_user_ref_count
