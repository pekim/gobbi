// This is a generated file - DO NOT EDIT
// +build gtk_2.20 gtk_2.22 gtk_2.24 gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

// CellRendererSpinnerNew is a wrapper around the C function gtk_cell_renderer_spinner_new.
func CellRendererSpinnerNew() *CellRendererSpinner {
	retC := C.gtk_cell_renderer_spinner_new()
	retGo := CellRendererSpinnerNewFromC(unsafe.Pointer(retC))

	return retGo
}

// OffscreenWindowNew is a wrapper around the C function gtk_offscreen_window_new.
func OffscreenWindowNew() *OffscreenWindow {
	retC := C.gtk_offscreen_window_new()
	retGo := OffscreenWindowNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SpinnerNew is a wrapper around the C function gtk_spinner_new.
func SpinnerNew() *Spinner {
	retC := C.gtk_spinner_new()
	retGo := SpinnerNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ToolItemGroupNew is a wrapper around the C function gtk_tool_item_group_new.
func ToolItemGroupNew(label string) *ToolItemGroup {
	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	retC := C.gtk_tool_item_group_new(c_label)
	retGo := ToolItemGroupNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ToolPaletteNew is a wrapper around the C function gtk_tool_palette_new.
func ToolPaletteNew() *ToolPalette {
	retC := C.gtk_tool_palette_new()
	retGo := ToolPaletteNewFromC(unsafe.Pointer(retC))

	return retGo
}
