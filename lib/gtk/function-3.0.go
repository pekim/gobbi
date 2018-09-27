// This is a generated file - DO NOT EDIT
// +build gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.20 gtk_3.22

package gtk

import (
	cairo "github.com/pekim/gobbi/lib/cairo"
	gdk "github.com/pekim/gobbi/lib/gdk"
	glib "github.com/pekim/gobbi/lib/glib"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

// BindingEntryAddSignalFromString is a wrapper around the C function gtk_binding_entry_add_signal_from_string.
func BindingEntryAddSignalFromString(bindingSet *BindingSet, signalDesc string) glib.TokenType {
	c_binding_set := (*C.GtkBindingSet)(bindingSet.ToC())

	c_signal_desc := C.CString(signalDesc)
	defer C.free(unsafe.Pointer(c_signal_desc))

	retC := C.gtk_binding_entry_add_signal_from_string(c_binding_set, c_signal_desc)
	retGo := (glib.TokenType)(retC)

	return retGo
}

// CairoShouldDrawWindow is a wrapper around the C function gtk_cairo_should_draw_window.
func CairoShouldDrawWindow(cr *cairo.Context, window *gdk.Window) bool {
	c_cr := (*C.cairo_t)(cr.ToC())

	c_window := (*C.GdkWindow)(window.ToC())

	retC := C.gtk_cairo_should_draw_window(c_cr, c_window)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_cairo_transform_to_window : no return generator

// Unsupported : gtk_device_grab_add : no return generator

// Unsupported : gtk_device_grab_remove : no return generator

// Unsupported : gtk_draw_insertion_cursor : unsupported parameter location : Blacklisted record : GdkRectangle

// GetBinaryAge is a wrapper around the C function gtk_get_binary_age.
func GetBinaryAge() uint32 {
	retC := C.gtk_get_binary_age()
	retGo := (uint32)(retC)

	return retGo
}

// GetInterfaceAge is a wrapper around the C function gtk_get_interface_age.
func GetInterfaceAge() uint32 {
	retC := C.gtk_get_interface_age()
	retGo := (uint32)(retC)

	return retGo
}

// GetMajorVersion is a wrapper around the C function gtk_get_major_version.
func GetMajorVersion() uint32 {
	retC := C.gtk_get_major_version()
	retGo := (uint32)(retC)

	return retGo
}

// GetMicroVersion is a wrapper around the C function gtk_get_micro_version.
func GetMicroVersion() uint32 {
	retC := C.gtk_get_micro_version()
	retGo := (uint32)(retC)

	return retGo
}

// GetMinorVersion is a wrapper around the C function gtk_get_minor_version.
func GetMinorVersion() uint32 {
	retC := C.gtk_get_minor_version()
	retGo := (uint32)(retC)

	return retGo
}

// Unsupported : gtk_render_activity : no return generator

// Unsupported : gtk_render_arrow : no return generator

// Unsupported : gtk_render_check : no return generator

// Unsupported : gtk_render_expander : no return generator

// Unsupported : gtk_render_extension : no return generator

// Unsupported : gtk_render_focus : no return generator

// Unsupported : gtk_render_frame : no return generator

// Unsupported : gtk_render_frame_gap : no return generator

// Unsupported : gtk_render_handle : no return generator

// Unsupported : gtk_render_icon_pixbuf : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_render_layout : no return generator

// Unsupported : gtk_render_line : no return generator

// Unsupported : gtk_render_option : no return generator

// Unsupported : gtk_render_slider : no return generator
