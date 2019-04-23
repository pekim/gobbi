// This is a generated file - DO NOT EDIT
// +build gdk_2.4 gdk_2.6 gdk_2.8 gdk_2.10 gdk_2.12 gdk_2.14 gdk_2.16 gdk_2.18 gdk_2.20 gdk_2.22 gdk_2.24 gdk_3.0 gdk_3.4 gdk_3.8 gdk_3.10 gdk_3.12 gdk_3.14 gdk_3.16 gdk_3.18 gdk_3.20 gdk_3.22

package gdk

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gdk/gdk.h>
// #include <stdlib.h>
import "C"

// Blacklisted : gdk_cursor_new_from_pixbuf

// Flush is a wrapper around the C function gdk_display_flush.
func (recv *Display) Flush() {
	C.gdk_display_flush((*C.GdkDisplay)(recv.native))

	return
}

// Blacklisted : gdk_display_get_default_cursor_size

// Blacklisted : gdk_display_get_default_group

// Blacklisted : gdk_display_get_maximal_cursor_size

// Blacklisted : gdk_display_set_double_click_distance

// Blacklisted : gdk_display_supports_cursor_alpha

// Blacklisted : gdk_display_supports_cursor_color

// Blacklisted : gdk_window_get_group

// Blacklisted : gdk_window_set_accept_focus

// Blacklisted : gdk_window_set_keep_above

// Blacklisted : gdk_window_set_keep_below
