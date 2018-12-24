// This is a generated file - DO NOT EDIT
// +build gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	cairo "github.com/pekim/gobbi/lib/cairo"
	gdk "github.com/pekim/gobbi/lib/gdk"
	pango "github.com/pekim/gobbi/lib/pango"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

// AcceleratorGetLabelWithKeycode is a wrapper around the C function gtk_accelerator_get_label_with_keycode.
func AcceleratorGetLabelWithKeycode(display *gdk.Display, acceleratorKey uint32, keycode uint32, acceleratorMods gdk.ModifierType) string {
	c_display := (*C.GdkDisplay)(C.NULL)
	if display != nil {
		c_display = (*C.GdkDisplay)(display.ToC())
	}

	c_accelerator_key := (C.guint)(acceleratorKey)

	c_keycode := (C.guint)(keycode)

	c_accelerator_mods := (C.GdkModifierType)(acceleratorMods)

	retC := C.gtk_accelerator_get_label_with_keycode(c_display, c_accelerator_key, c_keycode, c_accelerator_mods)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// AcceleratorNameWithKeycode is a wrapper around the C function gtk_accelerator_name_with_keycode.
func AcceleratorNameWithKeycode(display *gdk.Display, acceleratorKey uint32, keycode uint32, acceleratorMods gdk.ModifierType) string {
	c_display := (*C.GdkDisplay)(C.NULL)
	if display != nil {
		c_display = (*C.GdkDisplay)(display.ToC())
	}

	c_accelerator_key := (C.guint)(acceleratorKey)

	c_keycode := (C.guint)(keycode)

	c_accelerator_mods := (C.GdkModifierType)(acceleratorMods)

	retC := C.gtk_accelerator_name_with_keycode(c_display, c_accelerator_key, c_keycode, c_accelerator_mods)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_accelerator_parse_with_keycode : unsupported parameter accelerator_codes : output array param accelerator_codes

// RenderInsertionCursor is a wrapper around the C function gtk_render_insertion_cursor.
func RenderInsertionCursor(context *StyleContext, cr *cairo.Context, x float64, y float64, layout *pango.Layout, index int32, direction pango.Direction) {
	c_context := (*C.GtkStyleContext)(C.NULL)
	if context != nil {
		c_context = (*C.GtkStyleContext)(context.ToC())
	}

	c_cr := (*C.cairo_t)(C.NULL)
	if cr != nil {
		c_cr = (*C.cairo_t)(cr.ToC())
	}

	c_x := (C.gdouble)(x)

	c_y := (C.gdouble)(y)

	c_layout := (*C.PangoLayout)(C.NULL)
	if layout != nil {
		c_layout = (*C.PangoLayout)(layout.ToC())
	}

	c_index := (C.int)(index)

	c_direction := (C.PangoDirection)(direction)

	C.gtk_render_insertion_cursor(c_context, c_cr, c_x, c_y, c_layout, c_index, c_direction)

	return
}
