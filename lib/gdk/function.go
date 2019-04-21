// This is a generated file - DO NOT EDIT

package gdk

import call "github.com/pekim/gobbi/lib/internal/call"

// Unsupported : gdk_atom_intern : return type :

// Beep is a wrapper around the C function gdk_beep.
func Beep() {
	data := call.Data{}
	call.Function(4321, data)
	return
}

// Unsupported : gdk_cairo_get_clip_rectangle : return type :

// Unsupported : gdk_cairo_region_create_from_surface : return type :

// Unsupported : gdk_color_parse : return type :

// Unsupported : gdk_drag_begin : return type :

// Unsupported : gdk_drag_begin_for_device : return type :

// Unsupported : gdk_drag_get_selection : return type :

// Unsupported : gdk_drag_motion : return type :

// ErrorTrapPop is a wrapper around the C function gdk_error_trap_pop.
func ErrorTrapPop() int32 {
	data := call.Data{}
	call.Function(4479, data)
	retGo := int32(3)

	return retGo
}

// ErrorTrapPush is a wrapper around the C function gdk_error_trap_push.
func ErrorTrapPush() {
	data := call.Data{}
	call.Function(4481, data)
	return
}

// Unsupported : gdk_event_get : no return generator

// Unsupported : gdk_event_handler_set : unsupported parameter func : no type generator for EventFunc (GdkEventFunc) for param func

// Unsupported : gdk_event_peek : no return generator

// Unsupported : gdk_events_pending : return type :

// Flush is a wrapper around the C function gdk_flush.
func Flush() {
	data := call.Data{}
	call.Function(4490, data)
	return
}

// Unsupported : gdk_get_default_root_window : return type :

// Unsupported : gdk_get_display : return type :

// Unsupported : gdk_get_program_class : return type :

// Unsupported : gdk_get_show_events : return type :

// Unsupported : gdk_gl_error_quark : return type :

// Unsupported : gdk_init_check : return type :

// Unsupported : gdk_keyboard_grab : return type :

// Unsupported : gdk_keyval_from_name : return type :

// Unsupported : gdk_keyval_is_lower : return type :

// Unsupported : gdk_keyval_is_upper : return type :

// Unsupported : gdk_keyval_name : return type :

// Unsupported : gdk_keyval_to_lower : return type :

// Unsupported : gdk_keyval_to_unicode : return type :

// Unsupported : gdk_keyval_to_upper : return type :

// Unsupported : gdk_list_visuals : return type :

// Unsupported : gdk_offscreen_window_get_surface : return type :

// Unsupported : gdk_pango_context_get : return type :

// Unsupported : gdk_pango_layout_get_clip_region : return type :

// Unsupported : gdk_pango_layout_line_get_clip_region : return type :

// Unsupported : gdk_pixbuf_get_from_surface : return type :

// Unsupported : gdk_pixbuf_get_from_window : return type :

// Unsupported : gdk_pointer_grab : return type :

// Unsupported : gdk_pointer_is_grabbed : return type :

// PreParseLibgtkOnly is a wrapper around the C function gdk_pre_parse_libgtk_only.
func PreParseLibgtkOnly() {
	data := call.Data{}
	call.Function(4700, data)
	return
}

// Unsupported : gdk_property_get : unsupported parameter actual_length : array length param actual_length is pointer (gint*)

// Unsupported : gdk_query_depths : unsupported parameter depths : output array param depths

// Unsupported : gdk_query_visual_types : unsupported parameter visual_types : output array param visual_types

// Unsupported : gdk_selection_owner_get : return type :

// Unsupported : gdk_selection_owner_set : return type :

// Unsupported : gdk_selection_property_get : unsupported parameter data : guchar** with indirection level of 2

// Unsupported : gdk_setting_get : return type :

// Blacklisted : gdk_synthesize_window_state

// ThreadsEnter is a wrapper around the C function gdk_threads_enter.
func ThreadsEnter() {
	data := call.Data{}
	call.Function(4781, data)
	return
}

// ThreadsInit is a wrapper around the C function gdk_threads_init.
func ThreadsInit() {
	data := call.Data{}
	call.Function(4782, data)
	return
}

// ThreadsLeave is a wrapper around the C function gdk_threads_leave.
func ThreadsLeave() {
	data := call.Data{}
	call.Function(4783, data)
	return
}

// Unsupported : gdk_unicode_to_keyval : return type :

// Unsupported : gdk_utf8_to_string_target : return type :
