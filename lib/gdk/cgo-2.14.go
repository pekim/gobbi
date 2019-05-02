// This is a generated file - DO NOT EDIT
// +build gdk_2.14 gdk_2.16 gdk_2.18 gdk_2.20 gdk_2.22 gdk_2.24 gdk_3.0 gdk_3.4 gdk_3.8 gdk_3.10 gdk_3.12 gdk_3.14 gdk_3.16 gdk_3.18 gdk_3.20 gdk_3.22

package gdk

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gdk/gdk.h>
// #include <stdlib.h>
import "C"

// AppLaunchContextNew is a wrapper around the C function gdk_app_launch_context_new.
func AppLaunchContextNew() *AppLaunchContext {
	retC := C.gdk_app_launch_context_new()
	retGo := AppLaunchContextNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// Unsupported : gdk_threads_add_timeout_seconds : unsupported parameter function : no type generator for GLib.SourceFunc (GSourceFunc) for param function

// Unsupported : gdk_threads_add_timeout_seconds_full : unsupported parameter function : no type generator for GLib.SourceFunc (GSourceFunc) for param function
