// This is a generated file - DO NOT EDIT
// +build gdk_2.2 gdk_2.4 gdk_2.6 gdk_2.8 gdk_2.10 gdk_2.12 gdk_2.14 gdk_2.16 gdk_2.18 gdk_2.20 gdk_2.22 gdk_2.24 gdk_3.0 gdk_3.4 gdk_3.8 gdk_3.10 gdk_3.12 gdk_3.14 gdk_3.16 gdk_3.18 gdk_3.20 gdk_3.22

package gdk

import (
	pango "github.com/pekim/gobbi/lib/pango"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gdk/gdk.h>
// #include <stdlib.h>
import "C"

// Unsupported : gdk_drag_find_window_for_screen : unsupported parameter protocol : GdkDragProtocol* with indirection level of 1

// GetDisplayArgName is a wrapper around the C function gdk_get_display_arg_name.
func GetDisplayArgName() string {
	retC := C.gdk_get_display_arg_name()
	retGo := C.GoString(retC)

	return retGo
}

// NotifyStartupComplete is a wrapper around the C function gdk_notify_startup_complete.
func NotifyStartupComplete() {
	C.gdk_notify_startup_complete()

	return
}

// PangoContextGetForScreen is a wrapper around the C function gdk_pango_context_get_for_screen.
func PangoContextGetForScreen(screen *Screen) *pango.Context {
	c_screen := (*C.GdkScreen)(C.NULL)
	if screen != nil {
		c_screen = (*C.GdkScreen)(screen.ToC())
	}

	retC := C.gdk_pango_context_get_for_screen(c_screen)
	retGo := pango.ContextNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ParseArgs is a wrapper around the C function gdk_parse_args.
func ParseArgs(args []string) []string {
	cArgc, cArgv := argsIn(args)

	C.gdk_parse_args(&cArgc, &cArgv)

	args = argsOut(cArgc, cArgv)

	return args
}

// Unsupported : gdk_selection_owner_get_for_display : unsupported parameter selection : Blacklisted record : GdkAtom

// Unsupported : gdk_selection_owner_set_for_display : unsupported parameter selection : Blacklisted record : GdkAtom

// Unsupported : gdk_selection_send_notify_for_display : unsupported parameter selection : Blacklisted record : GdkAtom

// Unsupported : gdk_text_property_to_utf8_list_for_display : unsupported parameter encoding : Blacklisted record : GdkAtom
