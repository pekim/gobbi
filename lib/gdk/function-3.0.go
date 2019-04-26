// This is a generated file - DO NOT EDIT
// +build gdk_3.0 gdk_3.4 gdk_3.8 gdk_3.10 gdk_3.12 gdk_3.14 gdk_3.16 gdk_3.18 gdk_3.20 gdk_3.22

package gdk

import call "github.com/pekim/gobbi/lib/internal/call"

// DisableMultidevice is a wrapper around the C function gdk_disable_multidevice.
func DisableMultidevice() {
	data := call.Data{Return: call.Value{Type: call.TYPE_VOID}}
	call.Function(4147, &data)
	return
}

// ErrorTrapPopIgnored is a wrapper around the C function gdk_error_trap_pop_ignored.
func ErrorTrapPopIgnored() {
	data := call.Data{Return: call.Value{Type: call.TYPE_VOID}}
	call.Function(4231, &data)
	return
}

// Unsupported : gdk_events_get_angle : unsupported parameter event1 : no type generator for Event (GdkEvent*) for param event1

// Unsupported : gdk_events_get_center : unsupported parameter event1 : no type generator for Event (GdkEvent*) for param event1

// Unsupported : gdk_events_get_distance : unsupported parameter event1 : no type generator for Event (GdkEvent*) for param event1
