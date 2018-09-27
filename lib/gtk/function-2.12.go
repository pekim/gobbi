// This is a generated file - DO NOT EDIT
// +build gtk_2.12 gtk_2.14 gtk_2.16 gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.20 gtk_3.22

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

// Unsupported : gtk_binding_entry_skip : no return generator

// PaperSizeGetPaperSizes is a wrapper around the C function gtk_paper_size_get_paper_sizes.
func PaperSizeGetPaperSizes(includeCustom bool) *glib.List {
	c_include_custom :=
		boolToGboolean(includeCustom)

	retC := C.gtk_paper_size_get_paper_sizes(c_include_custom)
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// RcParseColorFull is a wrapper around the C function gtk_rc_parse_color_full.
func RcParseColorFull(scanner *glib.Scanner, style *RcStyle) (uint32, *gdk.Color) {
	c_scanner := (*C.GScanner)(scanner.ToC())

	c_style := (*C.GtkRcStyle)(style.ToC())

	var c_color C.GdkColor

	retC := C.gtk_rc_parse_color_full(c_scanner, c_style, &c_color)
	retGo := (uint32)(retC)

	color := gdk.ColorNewFromC(unsafe.Pointer(c_color))

	return retGo, color
}
