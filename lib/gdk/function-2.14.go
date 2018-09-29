// This is a generated file - DO NOT EDIT
// +build gdk_2.14 gdk_2.16 gdk_2.18 gdk_2.20 gdk_2.22 gdk_2.24 gdk_3.0 gdk_3.4 gdk_3.8 gdk_3.10 gdk_3.12 gdk_3.14 gdk_3.16 gdk_3.18 gdk_3.20 gdk_3.22

package gdk

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <stdlib.h>
import "C"

// Unsupported : gdk_atom_intern : return type : Blacklisted record : GdkAtom

// Unsupported : gdk_atom_intern_static_string : return type : Blacklisted record : GdkAtom

// Unsupported : gdk_cairo_get_clip_rectangle : unsupported parameter rect : Blacklisted record : GdkRectangle

// Unsupported : gdk_cairo_rectangle : unsupported parameter rectangle : Blacklisted record : GdkRectangle

// Blacklisted : gdk_color_parse

// Unsupported : gdk_drag_find_window_for_screen : unsupported parameter dest_window : record with indirection level of 2

// Unsupported : gdk_drag_get_selection : return type : Blacklisted record : GdkAtom

// Unsupported : gdk_event_get : no return generator

// Unsupported : gdk_event_handler_set : unsupported parameter func : no type generator for EventFunc, GdkEventFunc

// Unsupported : gdk_event_peek : no return generator

// Unsupported : gdk_events_get_angle : unsupported parameter event1 : no type generator for Event, GdkEvent*

// Unsupported : gdk_events_get_center : unsupported parameter event1 : no type generator for Event, GdkEvent*

// Unsupported : gdk_events_get_distance : unsupported parameter event1 : no type generator for Event, GdkEvent*

// Unsupported : gdk_init : unsupported parameter argc : no type generator for gint, gint*

// Unsupported : gdk_init_check : unsupported parameter argc : no type generator for gint, gint*

// Unsupported : gdk_keyval_convert_case : unsupported parameter lower : no type generator for guint, guint*

// Unsupported : gdk_pango_layout_get_clip_region : unsupported parameter index_ranges : no type generator for gint, const gint*

// Unsupported : gdk_pango_layout_line_get_clip_region : unsupported parameter index_ranges : no param type

// Unsupported : gdk_parse_args : unsupported parameter argc : no type generator for gint, gint*

// Unsupported : gdk_property_change : unsupported parameter property : Blacklisted record : GdkAtom

// Unsupported : gdk_property_delete : unsupported parameter property : Blacklisted record : GdkAtom

// Unsupported : gdk_property_get : unsupported parameter property : Blacklisted record : GdkAtom

// Unsupported : gdk_query_depths : unsupported parameter depths : no param type

// Unsupported : gdk_query_visual_types : unsupported parameter visual_types : no param type

// Unsupported : gdk_selection_convert : unsupported parameter selection : Blacklisted record : GdkAtom

// Unsupported : gdk_selection_owner_get : unsupported parameter selection : Blacklisted record : GdkAtom

// Unsupported : gdk_selection_owner_get_for_display : unsupported parameter selection : Blacklisted record : GdkAtom

// Unsupported : gdk_selection_owner_set : unsupported parameter selection : Blacklisted record : GdkAtom

// Unsupported : gdk_selection_owner_set_for_display : unsupported parameter selection : Blacklisted record : GdkAtom

// Unsupported : gdk_selection_property_get : unsupported parameter data : no type generator for guint8, guchar**

// Unsupported : gdk_selection_send_notify : unsupported parameter selection : Blacklisted record : GdkAtom

// Unsupported : gdk_selection_send_notify_for_display : unsupported parameter selection : Blacklisted record : GdkAtom

// Blacklisted : gdk_synthesize_window_state

// TestRenderSync is a wrapper around the C function gdk_test_render_sync.
func TestRenderSync(window *Window) {
	c_window := (*C.GdkWindow)(window.ToC())

	C.gdk_test_render_sync(c_window)

	return
}

// TestSimulateButton is a wrapper around the C function gdk_test_simulate_button.
func TestSimulateButton(window *Window, x int32, y int32, button uint32, modifiers ModifierType, buttonPressrelease EventType) bool {
	c_window := (*C.GdkWindow)(window.ToC())

	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	c_button := (C.guint)(button)

	c_modifiers := (C.GdkModifierType)(modifiers)

	c_button_pressrelease := (C.GdkEventType)(buttonPressrelease)

	retC := C.gdk_test_simulate_button(c_window, c_x, c_y, c_button, c_modifiers, c_button_pressrelease)
	retGo := retC == C.TRUE

	return retGo
}

// TestSimulateKey is a wrapper around the C function gdk_test_simulate_key.
func TestSimulateKey(window *Window, x int32, y int32, keyval uint32, modifiers ModifierType, keyPressrelease EventType) bool {
	c_window := (*C.GdkWindow)(window.ToC())

	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	c_keyval := (C.guint)(keyval)

	c_modifiers := (C.GdkModifierType)(modifiers)

	c_key_pressrelease := (C.GdkEventType)(keyPressrelease)

	retC := C.gdk_test_simulate_key(c_window, c_x, c_y, c_keyval, c_modifiers, c_key_pressrelease)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gdk_text_property_to_utf8_list_for_display : unsupported parameter encoding : Blacklisted record : GdkAtom

// Unsupported : gdk_threads_add_idle : unsupported parameter function : no type generator for GLib.SourceFunc, GSourceFunc

// Unsupported : gdk_threads_add_idle_full : unsupported parameter function : no type generator for GLib.SourceFunc, GSourceFunc

// Unsupported : gdk_threads_add_timeout : unsupported parameter function : no type generator for GLib.SourceFunc, GSourceFunc

// Unsupported : gdk_threads_add_timeout_full : unsupported parameter function : no type generator for GLib.SourceFunc, GSourceFunc

// Unsupported : gdk_threads_add_timeout_seconds : unsupported parameter function : no type generator for GLib.SourceFunc, GSourceFunc

// Unsupported : gdk_threads_add_timeout_seconds_full : unsupported parameter function : no type generator for GLib.SourceFunc, GSourceFunc

// Unsupported : gdk_threads_set_lock_functions : unsupported parameter enter_fn : no type generator for GObject.Callback, GCallback
