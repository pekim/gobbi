// This is a generated file - DO NOT EDIT
// +build gtk_2.8 gtk_2.10 gtk_2.12 gtk_2.14 gtk_2.16 gtk_2.18 gtk_2.20 gtk_2.22 gtk_2.24 gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

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

	gboolean widget_grabBrokenEventHandler(GObject *, GdkEventGrabBroken *, gpointer);

	static gulong Widget_signal_connect_grab_broken_event(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "grab-broken-event", G_CALLBACK(widget_grabBrokenEventHandler), data);
	}

*/
import "C"

// Blacklisted : gtk_action_get_accel_closure

type signalWidgetGrabBrokenEventDetail struct {
	callback  WidgetSignalGrabBrokenEventCallback
	handlerID C.gulong
}

var signalWidgetGrabBrokenEventId int
var signalWidgetGrabBrokenEventMap = make(map[int]signalWidgetGrabBrokenEventDetail)
var signalWidgetGrabBrokenEventLock sync.RWMutex

// WidgetSignalGrabBrokenEventCallback is a callback function for a 'grab-broken-event' signal emitted from a Widget.
type WidgetSignalGrabBrokenEventCallback func(event *gdk.EventGrabBroken) bool

/*
ConnectGrabBrokenEvent connects the callback to the 'grab-broken-event' signal for the Widget.

The returned value represents the connection, and may be passed to DisconnectGrabBrokenEvent to remove it.
*/
func (recv *Widget) ConnectGrabBrokenEvent(callback WidgetSignalGrabBrokenEventCallback) int {
	signalWidgetGrabBrokenEventLock.Lock()
	defer signalWidgetGrabBrokenEventLock.Unlock()

	signalWidgetGrabBrokenEventId++
	instance := C.gpointer(recv.native)
	handlerID := C.Widget_signal_connect_grab_broken_event(instance, C.gpointer(uintptr(signalWidgetGrabBrokenEventId)))

	detail := signalWidgetGrabBrokenEventDetail{callback, handlerID}
	signalWidgetGrabBrokenEventMap[signalWidgetGrabBrokenEventId] = detail

	return signalWidgetGrabBrokenEventId
}

/*
DisconnectGrabBrokenEvent disconnects a callback from the 'grab-broken-event' signal for the Widget.

The connectionID should be a value returned from a call to ConnectGrabBrokenEvent.
*/
func (recv *Widget) DisconnectGrabBrokenEvent(connectionID int) {
	signalWidgetGrabBrokenEventLock.Lock()
	defer signalWidgetGrabBrokenEventLock.Unlock()

	detail, exists := signalWidgetGrabBrokenEventMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWidgetGrabBrokenEventMap, connectionID)
}

//export widget_grabBrokenEventHandler
func widget_grabBrokenEventHandler(_ *C.GObject, c_event *C.GdkEventGrabBroken, data C.gpointer) C.gboolean {
	signalWidgetGrabBrokenEventLock.RLock()
	defer signalWidgetGrabBrokenEventLock.RUnlock()

	event := gdk.EventGrabBrokenNewFromC(unsafe.Pointer(c_event))

	index := int(uintptr(data))
	callback := signalWidgetGrabBrokenEventMap[index].callback
	retGo := callback(event)
	retC :=
		boolToGboolean(retGo)
	return retC
}

// Blacklisted : gtk_drag_source_set_icon_name

// Blacklisted : gtk_window_get_urgency_hint

// Blacklisted : gtk_window_present_with_time

// Blacklisted : gtk_window_set_urgency_hint
