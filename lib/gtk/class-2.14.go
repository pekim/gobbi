// This is a generated file - DO NOT EDIT
// +build gtk_2.14 gtk_2.16 gtk_2.18 gtk_2.20 gtk_2.22 gtk_2.24 gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

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

	gboolean widget_damageEventHandler(GObject *, GdkEventExpose *, gpointer);

	static gulong Widget_signal_connect_damage_event(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "damage-event", G_CALLBACK(widget_damageEventHandler), data);
	}

*/
import "C"

// Blacklisted : gtk_container_get_focus_child

// Blacklisted : gtk_tooltip_set_icon_from_icon_name

type signalWidgetDamageEventDetail struct {
	callback  WidgetSignalDamageEventCallback
	handlerID C.gulong
}

var signalWidgetDamageEventId int
var signalWidgetDamageEventMap = make(map[int]signalWidgetDamageEventDetail)
var signalWidgetDamageEventLock sync.RWMutex

// WidgetSignalDamageEventCallback is a callback function for a 'damage-event' signal emitted from a Widget.
type WidgetSignalDamageEventCallback func(event *gdk.EventExpose) bool

/*
ConnectDamageEvent connects the callback to the 'damage-event' signal for the Widget.

The returned value represents the connection, and may be passed to DisconnectDamageEvent to remove it.
*/
func (recv *Widget) ConnectDamageEvent(callback WidgetSignalDamageEventCallback) int {
	signalWidgetDamageEventLock.Lock()
	defer signalWidgetDamageEventLock.Unlock()

	signalWidgetDamageEventId++
	instance := C.gpointer(recv.native)
	handlerID := C.Widget_signal_connect_damage_event(instance, C.gpointer(uintptr(signalWidgetDamageEventId)))

	detail := signalWidgetDamageEventDetail{callback, handlerID}
	signalWidgetDamageEventMap[signalWidgetDamageEventId] = detail

	return signalWidgetDamageEventId
}

/*
DisconnectDamageEvent disconnects a callback from the 'damage-event' signal for the Widget.

The connectionID should be a value returned from a call to ConnectDamageEvent.
*/
func (recv *Widget) DisconnectDamageEvent(connectionID int) {
	signalWidgetDamageEventLock.Lock()
	defer signalWidgetDamageEventLock.Unlock()

	detail, exists := signalWidgetDamageEventMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWidgetDamageEventMap, connectionID)
}

//export widget_damageEventHandler
func widget_damageEventHandler(_ *C.GObject, c_event *C.GdkEventExpose, data C.gpointer) C.gboolean {
	signalWidgetDamageEventLock.RLock()
	defer signalWidgetDamageEventLock.RUnlock()

	event := gdk.EventExposeNewFromC(unsafe.Pointer(c_event))

	index := int(uintptr(data))
	callback := signalWidgetDamageEventMap[index].callback
	retGo := callback(event)
	retC :=
		boolToGboolean(retGo)
	return retC
}

// Blacklisted : gtk_widget_get_window

// Blacklisted : gtk_window_get_default_widget
