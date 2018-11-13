// This is a generated file - DO NOT EDIT
// +build gtk_2.12 gtk_2.14 gtk_2.16 gtk_2.18 gtk_2.20 gtk_2.22 gtk_2.24 gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

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

// Install a binding on @binding_set which causes key lookups
// to be aborted, to prevent bindings from lower priority sets
// to be activated.
/*

C function : gtk_binding_entry_skip
*/
func BindingEntrySkip(bindingSet *BindingSet, keyval uint32, modifiers gdk.ModifierType) {
	c_binding_set := (*C.GtkBindingSet)(C.NULL)
	if bindingSet != nil {
		c_binding_set = (*C.GtkBindingSet)(bindingSet.ToC())
	}

	c_keyval := (C.guint)(keyval)

	c_modifiers := (C.GdkModifierType)(modifiers)

	C.gtk_binding_entry_skip(c_binding_set, c_keyval, c_modifiers)

	return
}

// Creates a list of known paper sizes.
/*

C function : gtk_paper_size_get_paper_sizes
*/
func PaperSizeGetPaperSizes(includeCustom bool) *glib.List {
	c_include_custom :=
		boolToGboolean(includeCustom)

	retC := C.gtk_paper_size_get_paper_sizes(c_include_custom)
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Parses a color in the format expected
// in a RC file. If @style is not %NULL, it will be consulted to resolve
// references to symbolic colors.
/*

C function : gtk_rc_parse_color_full
*/
func RcParseColorFull(scanner *glib.Scanner, style *RcStyle) (uint32, *gdk.Color) {
	c_scanner := (*C.GScanner)(C.NULL)
	if scanner != nil {
		c_scanner = (*C.GScanner)(scanner.ToC())
	}

	c_style := (*C.GtkRcStyle)(C.NULL)
	if style != nil {
		c_style = (*C.GtkRcStyle)(style.ToC())
	}

	var c_color C.GdkColor

	retC := C.gtk_rc_parse_color_full(c_scanner, c_style, &c_color)
	retGo := (uint32)(retC)

	color := gdk.ColorNewFromC(unsafe.Pointer(&c_color))

	return retGo, color
}
