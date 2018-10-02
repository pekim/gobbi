// This is a generated file - DO NOT EDIT
// +build gdk_3.16 gdk_3.18 gdk_3.20 gdk_3.22

package gdk

import cairo "github.com/pekim/gobbi/lib/cairo"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <stdlib.h>
import "C"

// Unsupported : gdk_atom_intern : return type : Blacklisted record : GdkAtom

// Unsupported : gdk_atom_intern_static_string : return type : Blacklisted record : GdkAtom

// CairoDrawFromGl is a wrapper around the C function gdk_cairo_draw_from_gl.
func CairoDrawFromGl(cr *cairo.Context, window *Window, source int32, sourceType int32, bufferScale int32, x int32, y int32, width int32, height int32) {
	c_cr := (*C.cairo_t)(cr.ToC())

	c_window := (*C.GdkWindow)(window.ToC())

	c_source := (C.int)(source)

	c_source_type := (C.int)(sourceType)

	c_buffer_scale := (C.int)(bufferScale)

	c_x := (C.int)(x)

	c_y := (C.int)(y)

	c_width := (C.int)(width)

	c_height := (C.int)(height)

	C.gdk_cairo_draw_from_gl(c_cr, c_window, c_source, c_source_type, c_buffer_scale, c_x, c_y, c_width, c_height)

	return
}

// Unsupported : gdk_cairo_get_clip_rectangle : unsupported parameter rect : Blacklisted record : GdkRectangle

// Unsupported : gdk_cairo_rectangle : unsupported parameter rectangle : Blacklisted record : GdkRectangle

// Unsupported : gdk_drag_find_window_for_screen : unsupported parameter dest_window : record with indirection level of 2

// Unsupported : gdk_drag_get_selection : return type : Blacklisted record : GdkAtom

// Unsupported : gdk_event_get : no return generator

// Unsupported : gdk_event_handler_set : unsupported parameter func : no type generator for EventFunc, GdkEventFunc

// Unsupported : gdk_event_peek : no return generator

// Unsupported : gdk_events_get_angle : unsupported parameter event1 : no type generator for Event, GdkEvent*

// Unsupported : gdk_events_get_center : unsupported parameter event1 : no type generator for Event, GdkEvent*

// Unsupported : gdk_events_get_distance : unsupported parameter event1 : no type generator for Event, GdkEvent*

// Unsupported : gdk_init : unsupported parameter args : no type generator for argcargv,

// Unsupported : gdk_init_check : unsupported parameter args : no type generator for argcargv,

// Unsupported : gdk_pango_layout_line_get_clip_region : unsupported parameter index_ranges : no param type

// Unsupported : gdk_parse_args : unsupported parameter args : no type generator for argcargv,

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

// Unsupported : gdk_selection_property_get : unsupported parameter data : guchar** with indirection level of 2

// Unsupported : gdk_selection_send_notify : unsupported parameter selection : Blacklisted record : GdkAtom

// Unsupported : gdk_selection_send_notify_for_display : unsupported parameter selection : Blacklisted record : GdkAtom

// Unsupported : gdk_text_property_to_utf8_list_for_display : unsupported parameter encoding : Blacklisted record : GdkAtom

// Unsupported : gdk_threads_add_idle : unsupported parameter function : no type generator for GLib.SourceFunc, GSourceFunc

// Unsupported : gdk_threads_add_idle_full : unsupported parameter function : no type generator for GLib.SourceFunc, GSourceFunc

// Unsupported : gdk_threads_add_timeout : unsupported parameter function : no type generator for GLib.SourceFunc, GSourceFunc

// Unsupported : gdk_threads_add_timeout_full : unsupported parameter function : no type generator for GLib.SourceFunc, GSourceFunc

// Unsupported : gdk_threads_add_timeout_seconds : unsupported parameter function : no type generator for GLib.SourceFunc, GSourceFunc

// Unsupported : gdk_threads_add_timeout_seconds_full : unsupported parameter function : no type generator for GLib.SourceFunc, GSourceFunc

// Unsupported : gdk_threads_set_lock_functions : unsupported parameter enter_fn : no type generator for GObject.Callback, GCallback
