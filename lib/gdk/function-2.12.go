// This is a generated file - DO NOT EDIT
// +build gdk_2.12 gdk_2.14 gdk_2.16 gdk_2.18 gdk_2.20 gdk_2.22 gdk_2.24 gdk_3.0 gdk_3.4 gdk_3.8 gdk_3.10 gdk_3.12 gdk_3.14 gdk_3.16 gdk_3.18 gdk_3.20 gdk_3.22

package gdk

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <stdlib.h>
import "C"

// Request more motion notifies if @event is a motion notify hint event.
//
// This function should be used instead of gdk_window_get_pointer() to
// request further motion notifies, because it also works for extension
// events where motion notifies are provided for devices other than the
// core pointer. Coordinate extraction, processing and requesting more
// motion events from a %GDK_MOTION_NOTIFY event usually works like this:
//
// |[<!-- language="C" -->
// {
// motion_event handler
// x = motion_event->x;
// y = motion_event->y;
// handle (x,y) motion
// gdk_event_request_motions (motion_event); // handles is_hint events
// }
// ]|
/*

C function : gdk_event_request_motions
*/
func EventRequestMotions(event *EventMotion) {
	c_event := (*C.GdkEventMotion)(C.NULL)
	if event != nil {
		c_event = (*C.GdkEventMotion)(event.ToC())
	}

	C.gdk_event_request_motions(c_event)

	return
}

// Indicates to the GUI environment that the application has
// finished loading, using a given identifier.
//
// GTK+ will call this function automatically for #GtkWindow
// with custom startup-notification identifier unless
// gtk_window_set_auto_startup_notification() is called to
// disable that feature.
/*

C function : gdk_notify_startup_complete_with_id
*/
func NotifyStartupCompleteWithId(startupId string) {
	c_startup_id := C.CString(startupId)
	defer C.free(unsafe.Pointer(c_startup_id))

	C.gdk_notify_startup_complete_with_id(c_startup_id)

	return
}

// Unsupported : gdk_threads_add_idle : unsupported parameter function : no type generator for GLib.SourceFunc (GSourceFunc) for param function

// Unsupported : gdk_threads_add_idle_full : unsupported parameter function : no type generator for GLib.SourceFunc (GSourceFunc) for param function

// Unsupported : gdk_threads_add_timeout : unsupported parameter function : no type generator for GLib.SourceFunc (GSourceFunc) for param function

// Unsupported : gdk_threads_add_timeout_full : unsupported parameter function : no type generator for GLib.SourceFunc (GSourceFunc) for param function
