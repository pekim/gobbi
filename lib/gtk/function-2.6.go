// This is a generated file - DO NOT EDIT
// +build gtk_2.6 gtk_2.8 gtk_2.10 gtk_2.12 gtk_2.14 gtk_2.16 gtk_2.18 gtk_2.20 gtk_2.22 gtk_2.24 gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
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

// AcceleratorGetLabel is a wrapper around the C function gtk_accelerator_get_label.
func AcceleratorGetLabel(acceleratorKey uint32, acceleratorMods gdk.ModifierType) string {
	c_accelerator_key := (C.guint)(acceleratorKey)

	c_accelerator_mods := (C.GdkModifierType)(acceleratorMods)

	retC := C.gtk_accelerator_get_label(c_accelerator_key, c_accelerator_mods)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// AlternativeDialogButtonOrder is a wrapper around the C function gtk_alternative_dialog_button_order.
func AlternativeDialogButtonOrder(screen *gdk.Screen) bool {
	c_screen := (*C.GdkScreen)(screen.ToC())

	retC := C.gtk_alternative_dialog_button_order(c_screen)
	retGo := retC == C.TRUE

	return retGo
}

// GetOptionGroup is a wrapper around the C function gtk_get_option_group.
func GetOptionGroup(openDefaultDisplay bool) *glib.OptionGroup {
	c_open_default_display :=
		boolToGboolean(openDefaultDisplay)

	retC := C.gtk_get_option_group(c_open_default_display)
	retGo := glib.OptionGroupNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_init_with_args : unsupported parameter entries :

// Unsupported : gtk_show_about_dialog : unsupported parameter ... : varargs
