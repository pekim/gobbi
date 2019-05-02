// This is a generated file - DO NOT EDIT
// +build gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

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

// LevelBarNew is a wrapper around the C function gtk_level_bar_new.
func LevelBarNew() *LevelBar {
	retC := C.gtk_level_bar_new()
	retGo := LevelBarNewFromC(unsafe.Pointer(retC))

	return retGo
}

// LevelBarNewForInterval is a wrapper around the C function gtk_level_bar_new_for_interval.
func LevelBarNewForInterval(minValue float64, maxValue float64) *LevelBar {
	c_min_value := (C.gdouble)(minValue)

	c_max_value := (C.gdouble)(maxValue)

	retC := C.gtk_level_bar_new_for_interval(c_min_value, c_max_value)
	retGo := LevelBarNewFromC(unsafe.Pointer(retC))

	return retGo
}

// MenuButtonNew is a wrapper around the C function gtk_menu_button_new.
func MenuButtonNew() *MenuButton {
	retC := C.gtk_menu_button_new()
	retGo := MenuButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SearchEntryNew is a wrapper around the C function gtk_search_entry_new.
func SearchEntryNew() *SearchEntry {
	retC := C.gtk_search_entry_new()
	retGo := SearchEntryNewFromC(unsafe.Pointer(retC))

	return retGo
}
