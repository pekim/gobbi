// This is a generated file - DO NOT EDIT
// +build gdk_2.2 gdk_2.4 gdk_2.6 gdk_2.8 gdk_2.10 gdk_2.12 gdk_2.14 gdk_2.16 gdk_2.18 gdk_2.20 gdk_2.22 gdk_2.24 gdk_3.0 gdk_3.4 gdk_3.8 gdk_3.10 gdk_3.12 gdk_3.14 gdk_3.16 gdk_3.18 gdk_3.20 gdk_3.22

package gdk

import (
	pango "github.com/pekim/gobbi/lib/pango"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <stdlib.h>
import "C"

// Unsupported : gdk_drag_find_window_for_screen : unsupported parameter protocol : GdkDragProtocol* with indirection level of 1

// Gets the display name specified in the command line arguments passed
// to gdk_init() or gdk_parse_args(), if any.
/*

C function : gdk_get_display_arg_name
*/
func GetDisplayArgName() string {
	retC := C.gdk_get_display_arg_name()
	retGo := C.GoString(retC)

	return retGo
}

// Indicates to the GUI environment that the application has finished
// loading. If the applications opens windows, this function is
// normally called after opening the application’s initial set of
// windows.
//
// GTK+ will call this function automatically after opening the first
// #GtkWindow unless gtk_window_set_auto_startup_notification() is called
// to disable that feature.
/*

C function : gdk_notify_startup_complete
*/
func NotifyStartupComplete() {
	C.gdk_notify_startup_complete()

	return
}

// Creates a #PangoContext for @screen.
//
// The context must be freed when you’re finished with it.
//
// When using GTK+, normally you should use gtk_widget_get_pango_context()
// instead of this function, to get the appropriate context for
// the widget you intend to render text onto.
//
// The newly created context will have the default font options
// (see #cairo_font_options_t) for the screen; if these options
// change it will not be updated. Using gtk_widget_get_pango_context()
// is more convenient if you want to keep a context around and track
// changes to the screen’s font rendering settings.
/*

C function : gdk_pango_context_get_for_screen
*/
func PangoContextGetForScreen(screen *Screen) *pango.Context {
	c_screen := (*C.GdkScreen)(C.NULL)
	if screen != nil {
		c_screen = (*C.GdkScreen)(screen.ToC())
	}

	retC := C.gdk_pango_context_get_for_screen(c_screen)
	retGo := pango.ContextNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Parse command line arguments, and store for future
// use by calls to gdk_display_open().
//
// Any arguments used by GDK are removed from the array and @argc and @argv are
// updated accordingly.
//
// You shouldn’t call this function explicitly if you are using
// gtk_init(), gtk_init_check(), gdk_init(), or gdk_init_check().
/*

C function : gdk_parse_args
*/
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
