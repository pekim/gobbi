// This is a generated file - DO NOT EDIT

package gdk

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gdk/gdk.h>
// #include <stdlib.h>
import "C"

// Blacklisted : gdk_add_option_entries_libgtk_only

// Blacklisted : gdk_atom_intern

// Blacklisted : gdk_beep

// Blacklisted : gdk_cairo_get_clip_rectangle

// Blacklisted : gdk_cairo_region_create_from_surface

// Blacklisted : gdk_color_parse

// Blacklisted : gdk_drag_abort

// Blacklisted : gdk_drag_begin

// Blacklisted : gdk_drag_begin_for_device

// Blacklisted : gdk_drag_drop

// Blacklisted : gdk_drag_get_selection

// Blacklisted : gdk_drag_motion

// Blacklisted : gdk_drag_status

// Blacklisted : gdk_drop_finish

// Blacklisted : gdk_drop_reply

// Blacklisted : gdk_error_trap_pop

// Blacklisted : gdk_error_trap_push

// Unsupported : gdk_event_get : no return generator

// Unsupported : gdk_event_handler_set : unsupported parameter func : no type generator for EventFunc (GdkEventFunc) for param func

// Unsupported : gdk_event_peek : no return generator

// Blacklisted : gdk_events_pending

// Flush is a wrapper around the C function gdk_flush.
func Flush() {
	C.gdk_flush()

	return
}

// Blacklisted : gdk_get_default_root_window

// Blacklisted : gdk_get_display

// Blacklisted : gdk_get_program_class

// Blacklisted : gdk_get_show_events

// Blacklisted : gdk_gl_error_quark

// Init is a wrapper around the C function gdk_init.
func Init(args []string) []string {
	cArgc, cArgv := argsIn(args)

	C.gdk_init(&cArgc, &cArgv)

	args = argsOut(cArgc, cArgv)

	return args
}

// Blacklisted : gdk_init_check

// Blacklisted : gdk_keyboard_grab

// Blacklisted : gdk_keyboard_ungrab

// Blacklisted : gdk_keyval_convert_case

// Blacklisted : gdk_keyval_from_name

// Blacklisted : gdk_keyval_is_lower

// Blacklisted : gdk_keyval_is_upper

// Blacklisted : gdk_keyval_name

// Blacklisted : gdk_keyval_to_lower

// Blacklisted : gdk_keyval_to_unicode

// Blacklisted : gdk_keyval_to_upper

// Blacklisted : gdk_list_visuals

// Blacklisted : gdk_offscreen_window_get_surface

// Blacklisted : gdk_pango_context_get

// Blacklisted : gdk_pango_layout_get_clip_region

// Blacklisted : gdk_pango_layout_line_get_clip_region

// Blacklisted : gdk_pixbuf_get_from_surface

// Blacklisted : gdk_pixbuf_get_from_window

// Blacklisted : gdk_pointer_grab

// Blacklisted : gdk_pointer_is_grabbed

// Blacklisted : gdk_pointer_ungrab

// Blacklisted : gdk_pre_parse_libgtk_only

// Blacklisted : gdk_property_change

// Blacklisted : gdk_property_delete

// Unsupported : gdk_property_get : unsupported parameter actual_length : array length param actual_length is pointer (gint*)

// Unsupported : gdk_query_depths : unsupported parameter depths : output array param depths

// Unsupported : gdk_query_visual_types : unsupported parameter visual_types : output array param visual_types

// Blacklisted : gdk_selection_convert

// Blacklisted : gdk_selection_owner_get

// Blacklisted : gdk_selection_owner_set

// Unsupported : gdk_selection_property_get : unsupported parameter data : guchar** with indirection level of 2

// Blacklisted : gdk_selection_send_notify

// Blacklisted : gdk_set_double_click_time

// Blacklisted : gdk_set_program_class

// Blacklisted : gdk_set_show_events

// Blacklisted : gdk_setting_get

// Blacklisted : gdk_synthesize_window_state

// Blacklisted : gdk_threads_enter

// Blacklisted : gdk_threads_init

// Blacklisted : gdk_threads_leave

// Blacklisted : gdk_unicode_to_keyval

// Blacklisted : gdk_utf8_to_string_target
