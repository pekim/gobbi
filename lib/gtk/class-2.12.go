// This is a generated file - DO NOT EDIT
// +build gtk_2.12 gtk_2.14 gtk_2.16 gtk_2.18 gtk_2.20 gtk_2.22 gtk_2.24 gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

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

	gboolean widget_dragFailedHandler(GObject *, GdkDragContext *, GtkDragResult, gpointer);

	static gulong Widget_signal_connect_drag_failed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "drag-failed", G_CALLBACK(widget_dragFailedHandler), data);
	}

*/
/*

	gboolean widget_keynavFailedHandler(GObject *, GtkDirectionType, gpointer);

	static gulong Widget_signal_connect_keynav_failed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "keynav-failed", G_CALLBACK(widget_keynavFailedHandler), data);
	}

*/
/*

	gboolean widget_queryTooltipHandler(GObject *, gint, gint, gboolean, GtkTooltip *, gpointer);

	static gulong Widget_signal_connect_query_tooltip(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "query-tooltip", G_CALLBACK(widget_queryTooltipHandler), data);
	}

*/
import "C"

// Blacklisted : gtk_action_create_menu

// Unsupported : gtk_list_store_set_valuesv : unsupported parameter values :

// Blacklisted : gtk_tooltip_trigger_tooltip_query

// Blacklisted : gtk_tooltip_set_custom

// Blacklisted : gtk_tooltip_set_icon

// Blacklisted : gtk_tooltip_set_icon_from_stock

// Blacklisted : gtk_tooltip_set_markup

// SetText is a wrapper around the C function gtk_tooltip_set_text.
func (recv *Tooltip) SetText(text string) {
	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))

	C.gtk_tooltip_set_text((*C.GtkTooltip)(recv.native), c_text)

	return
}

// Blacklisted : gtk_tooltip_set_tip_area

type signalWidgetDragFailedDetail struct {
	callback  WidgetSignalDragFailedCallback
	handlerID C.gulong
}

var signalWidgetDragFailedId int
var signalWidgetDragFailedMap = make(map[int]signalWidgetDragFailedDetail)
var signalWidgetDragFailedLock sync.RWMutex

// WidgetSignalDragFailedCallback is a callback function for a 'drag-failed' signal emitted from a Widget.
type WidgetSignalDragFailedCallback func(context *gdk.DragContext, result DragResult) bool

/*
ConnectDragFailed connects the callback to the 'drag-failed' signal for the Widget.

The returned value represents the connection, and may be passed to DisconnectDragFailed to remove it.
*/
func (recv *Widget) ConnectDragFailed(callback WidgetSignalDragFailedCallback) int {
	signalWidgetDragFailedLock.Lock()
	defer signalWidgetDragFailedLock.Unlock()

	signalWidgetDragFailedId++
	instance := C.gpointer(recv.native)
	handlerID := C.Widget_signal_connect_drag_failed(instance, C.gpointer(uintptr(signalWidgetDragFailedId)))

	detail := signalWidgetDragFailedDetail{callback, handlerID}
	signalWidgetDragFailedMap[signalWidgetDragFailedId] = detail

	return signalWidgetDragFailedId
}

/*
DisconnectDragFailed disconnects a callback from the 'drag-failed' signal for the Widget.

The connectionID should be a value returned from a call to ConnectDragFailed.
*/
func (recv *Widget) DisconnectDragFailed(connectionID int) {
	signalWidgetDragFailedLock.Lock()
	defer signalWidgetDragFailedLock.Unlock()

	detail, exists := signalWidgetDragFailedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWidgetDragFailedMap, connectionID)
}

//export widget_dragFailedHandler
func widget_dragFailedHandler(_ *C.GObject, c_context *C.GdkDragContext, c_result C.GtkDragResult, data C.gpointer) C.gboolean {
	signalWidgetDragFailedLock.RLock()
	defer signalWidgetDragFailedLock.RUnlock()

	context := gdk.DragContextNewFromC(unsafe.Pointer(c_context))

	result := DragResult(c_result)

	index := int(uintptr(data))
	callback := signalWidgetDragFailedMap[index].callback
	retGo := callback(context, result)
	retC :=
		boolToGboolean(retGo)
	return retC
}

type signalWidgetKeynavFailedDetail struct {
	callback  WidgetSignalKeynavFailedCallback
	handlerID C.gulong
}

var signalWidgetKeynavFailedId int
var signalWidgetKeynavFailedMap = make(map[int]signalWidgetKeynavFailedDetail)
var signalWidgetKeynavFailedLock sync.RWMutex

// WidgetSignalKeynavFailedCallback is a callback function for a 'keynav-failed' signal emitted from a Widget.
type WidgetSignalKeynavFailedCallback func(direction DirectionType) bool

/*
ConnectKeynavFailed connects the callback to the 'keynav-failed' signal for the Widget.

The returned value represents the connection, and may be passed to DisconnectKeynavFailed to remove it.
*/
func (recv *Widget) ConnectKeynavFailed(callback WidgetSignalKeynavFailedCallback) int {
	signalWidgetKeynavFailedLock.Lock()
	defer signalWidgetKeynavFailedLock.Unlock()

	signalWidgetKeynavFailedId++
	instance := C.gpointer(recv.native)
	handlerID := C.Widget_signal_connect_keynav_failed(instance, C.gpointer(uintptr(signalWidgetKeynavFailedId)))

	detail := signalWidgetKeynavFailedDetail{callback, handlerID}
	signalWidgetKeynavFailedMap[signalWidgetKeynavFailedId] = detail

	return signalWidgetKeynavFailedId
}

/*
DisconnectKeynavFailed disconnects a callback from the 'keynav-failed' signal for the Widget.

The connectionID should be a value returned from a call to ConnectKeynavFailed.
*/
func (recv *Widget) DisconnectKeynavFailed(connectionID int) {
	signalWidgetKeynavFailedLock.Lock()
	defer signalWidgetKeynavFailedLock.Unlock()

	detail, exists := signalWidgetKeynavFailedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWidgetKeynavFailedMap, connectionID)
}

//export widget_keynavFailedHandler
func widget_keynavFailedHandler(_ *C.GObject, c_direction C.GtkDirectionType, data C.gpointer) C.gboolean {
	signalWidgetKeynavFailedLock.RLock()
	defer signalWidgetKeynavFailedLock.RUnlock()

	direction := DirectionType(c_direction)

	index := int(uintptr(data))
	callback := signalWidgetKeynavFailedMap[index].callback
	retGo := callback(direction)
	retC :=
		boolToGboolean(retGo)
	return retC
}

type signalWidgetQueryTooltipDetail struct {
	callback  WidgetSignalQueryTooltipCallback
	handlerID C.gulong
}

var signalWidgetQueryTooltipId int
var signalWidgetQueryTooltipMap = make(map[int]signalWidgetQueryTooltipDetail)
var signalWidgetQueryTooltipLock sync.RWMutex

// WidgetSignalQueryTooltipCallback is a callback function for a 'query-tooltip' signal emitted from a Widget.
type WidgetSignalQueryTooltipCallback func(x int32, y int32, keyboardMode bool, tooltip *Tooltip) bool

/*
ConnectQueryTooltip connects the callback to the 'query-tooltip' signal for the Widget.

The returned value represents the connection, and may be passed to DisconnectQueryTooltip to remove it.
*/
func (recv *Widget) ConnectQueryTooltip(callback WidgetSignalQueryTooltipCallback) int {
	signalWidgetQueryTooltipLock.Lock()
	defer signalWidgetQueryTooltipLock.Unlock()

	signalWidgetQueryTooltipId++
	instance := C.gpointer(recv.native)
	handlerID := C.Widget_signal_connect_query_tooltip(instance, C.gpointer(uintptr(signalWidgetQueryTooltipId)))

	detail := signalWidgetQueryTooltipDetail{callback, handlerID}
	signalWidgetQueryTooltipMap[signalWidgetQueryTooltipId] = detail

	return signalWidgetQueryTooltipId
}

/*
DisconnectQueryTooltip disconnects a callback from the 'query-tooltip' signal for the Widget.

The connectionID should be a value returned from a call to ConnectQueryTooltip.
*/
func (recv *Widget) DisconnectQueryTooltip(connectionID int) {
	signalWidgetQueryTooltipLock.Lock()
	defer signalWidgetQueryTooltipLock.Unlock()

	detail, exists := signalWidgetQueryTooltipMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWidgetQueryTooltipMap, connectionID)
}

//export widget_queryTooltipHandler
func widget_queryTooltipHandler(_ *C.GObject, c_x C.gint, c_y C.gint, c_keyboard_mode C.gboolean, c_tooltip *C.GtkTooltip, data C.gpointer) C.gboolean {
	signalWidgetQueryTooltipLock.RLock()
	defer signalWidgetQueryTooltipLock.RUnlock()

	x := int32(c_x)

	y := int32(c_y)

	keyboardMode := c_keyboard_mode == C.TRUE

	tooltip := TooltipNewFromC(unsafe.Pointer(c_tooltip))

	index := int(uintptr(data))
	callback := signalWidgetQueryTooltipMap[index].callback
	retGo := callback(x, y, keyboardMode, tooltip)
	retC :=
		boolToGboolean(retGo)
	return retC
}

// Blacklisted : gtk_widget_error_bell

// Blacklisted : gtk_widget_get_has_tooltip

// Blacklisted : gtk_widget_get_tooltip_markup

// Blacklisted : gtk_widget_get_tooltip_text

// Blacklisted : gtk_widget_get_tooltip_window

// Blacklisted : gtk_widget_keynav_failed

// Blacklisted : gtk_widget_modify_cursor

// Blacklisted : gtk_widget_set_has_tooltip

// Blacklisted : gtk_widget_set_tooltip_markup

// Blacklisted : gtk_widget_set_tooltip_text

// Blacklisted : gtk_widget_set_tooltip_window

// Blacklisted : gtk_widget_trigger_tooltip_query

// Blacklisted : gtk_window_get_opacity

// Blacklisted : gtk_window_set_opacity

// Blacklisted : gtk_window_set_startup_id
