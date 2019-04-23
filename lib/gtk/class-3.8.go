// This is a generated file - DO NOT EDIT
// +build gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

// Blacklisted : gtk_builder_expose_object

// Unsupported : gtk_icon_info_load_icon_async : unsupported parameter callback : no type generator for Gio.AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Blacklisted : gtk_icon_info_load_icon_finish

// Unsupported : gtk_icon_info_load_symbolic_async : unsupported parameter callback : no type generator for Gio.AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Blacklisted : gtk_icon_info_load_symbolic_finish

// Unsupported : gtk_icon_info_load_symbolic_for_context_async : unsupported parameter callback : no type generator for Gio.AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Blacklisted : gtk_icon_info_load_symbolic_for_context_finish

// Blacklisted : gtk_icon_view_get_activate_on_single_click

// Blacklisted : gtk_icon_view_set_activate_on_single_click

// Blacklisted : gtk_level_bar_get_inverted

// Blacklisted : gtk_level_bar_set_inverted

// Blacklisted : gtk_style_context_get_frame_clock

// Blacklisted : gtk_style_context_set_frame_clock

// Blacklisted : gtk_tree_view_get_activate_on_single_click

// Blacklisted : gtk_tree_view_set_activate_on_single_click

// Unsupported : gtk_widget_add_tick_callback : unsupported parameter callback : no type generator for TickCallback (GtkTickCallback) for param callback

// Blacklisted : gtk_widget_get_frame_clock

// Blacklisted : gtk_widget_get_opacity

// IsVisible is a wrapper around the C function gtk_widget_is_visible.
func (recv *Widget) IsVisible() bool {
	retC := C.gtk_widget_is_visible((*C.GtkWidget)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Blacklisted : gtk_widget_register_window

// Blacklisted : gtk_widget_remove_tick_callback

// Blacklisted : gtk_widget_set_opacity

// Blacklisted : gtk_widget_unregister_window
