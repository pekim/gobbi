// This is a generated file - DO NOT EDIT
// +build gtk_3.4 gtk_3.6 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.20 gtk_3.22

package gtk

import (
	gdk "github.com/pekim/gobbi/lib/gdk"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

// AcceleratorGetLabelWithKeycode is a wrapper around the C function gtk_accelerator_get_label_with_keycode.
func AcceleratorGetLabelWithKeycode(display *gdk.Display, acceleratorKey uint32, keycode uint32, acceleratorMods gdk.ModifierType) string {
	c_display := (*C.GdkDisplay)(display.ToC())

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
	c_display := (*C.GdkDisplay)(display.ToC())

	c_accelerator_key := (C.guint)(acceleratorKey)

	c_keycode := (C.guint)(keycode)

	c_accelerator_mods := (C.GdkModifierType)(acceleratorMods)

	retC := C.gtk_accelerator_name_with_keycode(c_display, c_accelerator_key, c_keycode, c_accelerator_mods)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_accelerator_parse_with_keycode : unsupported parameter accelerator_key : no type generator for guint, guint*

// Unsupported : gtk_render_insertion_cursor : no return generator
