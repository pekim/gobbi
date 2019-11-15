// Code generated - DO NOT EDIT.

package gdk

import gi "github.com/pekim/gobbi/internal/gi"

// UNSUPPORTED : C value 'gdk_add_option_entries_libgtk_only' : has parameters

// UNSUPPORTED : C value 'gdk_atom_intern' : has parameters

// UNSUPPORTED : C value 'gdk_atom_intern_static_string' : has parameters

var beepInvoker *gi.Function

// Beep is a representation of the C type gdk_beep.
func Beep() {
	if beepInvoker == nil {
		beepInvoker = gi.FunctionInvokerNew("Gdk", "beep")
	}

	beepInvoker.Invoke()
}

// UNSUPPORTED : C value 'gdk_cairo_create' : has parameters

// UNSUPPORTED : C value 'gdk_cairo_draw_from_gl' : has parameters

// UNSUPPORTED : C value 'gdk_cairo_get_clip_rectangle' : has parameters

// UNSUPPORTED : C value 'gdk_cairo_get_drawing_context' : has parameters

// UNSUPPORTED : C value 'gdk_cairo_rectangle' : has parameters

// UNSUPPORTED : C value 'gdk_cairo_region' : has parameters

// UNSUPPORTED : C value 'gdk_cairo_region_create_from_surface' : has parameters

// UNSUPPORTED : C value 'gdk_cairo_set_source_color' : has parameters

// UNSUPPORTED : C value 'gdk_cairo_set_source_pixbuf' : has parameters

// UNSUPPORTED : C value 'gdk_cairo_set_source_rgba' : has parameters

// UNSUPPORTED : C value 'gdk_cairo_set_source_window' : has parameters

// UNSUPPORTED : C value 'gdk_cairo_surface_create_from_pixbuf' : has parameters

// UNSUPPORTED : C value 'gdk_color_parse' : has parameters

var disableMultideviceInvoker *gi.Function

// DisableMultidevice is a representation of the C type gdk_disable_multidevice.
func DisableMultidevice() {
	if disableMultideviceInvoker == nil {
		disableMultideviceInvoker = gi.FunctionInvokerNew("Gdk", "disable_multidevice")
	}

	disableMultideviceInvoker.Invoke()
}

// UNSUPPORTED : C value 'gdk_drag_abort' : has parameters

// UNSUPPORTED : C value 'gdk_drag_begin' : has parameters

// UNSUPPORTED : C value 'gdk_drag_begin_for_device' : has parameters

// UNSUPPORTED : C value 'gdk_drag_begin_from_point' : has parameters

// UNSUPPORTED : C value 'gdk_drag_drop' : has parameters

// UNSUPPORTED : C value 'gdk_drag_drop_done' : has parameters

// UNSUPPORTED : C value 'gdk_drag_drop_succeeded' : has parameters

// UNSUPPORTED : C value 'gdk_drag_find_window_for_screen' : has parameters

// UNSUPPORTED : C value 'gdk_drag_get_selection' : has parameters

// UNSUPPORTED : C value 'gdk_drag_motion' : has parameters

// UNSUPPORTED : C value 'gdk_drag_status' : has parameters

// UNSUPPORTED : C value 'gdk_drop_finish' : has parameters

// UNSUPPORTED : C value 'gdk_drop_reply' : has parameters

var errorTrapPopInvoker *gi.Function

// ErrorTrapPop is a representation of the C type gdk_error_trap_pop.
func ErrorTrapPop() int32 {
	if errorTrapPopInvoker == nil {
		errorTrapPopInvoker = gi.FunctionInvokerNew("Gdk", "error_trap_pop")
	}

	ret := errorTrapPopInvoker.Invoke()
	return ret.Int32()
}

var errorTrapPopIgnoredInvoker *gi.Function

// ErrorTrapPopIgnored is a representation of the C type gdk_error_trap_pop_ignored.
func ErrorTrapPopIgnored() {
	if errorTrapPopIgnoredInvoker == nil {
		errorTrapPopIgnoredInvoker = gi.FunctionInvokerNew("Gdk", "error_trap_pop_ignored")
	}

	errorTrapPopIgnoredInvoker.Invoke()
}

var errorTrapPushInvoker *gi.Function

// ErrorTrapPush is a representation of the C type gdk_error_trap_push.
func ErrorTrapPush() {
	if errorTrapPushInvoker == nil {
		errorTrapPushInvoker = gi.FunctionInvokerNew("Gdk", "error_trap_push")
	}

	errorTrapPushInvoker.Invoke()
}

// UNSUPPORTED : C value 'gdk_event_get' : return type 'Event' not supported

// UNSUPPORTED : C value 'gdk_event_handler_set' : has parameters

// UNSUPPORTED : C value 'gdk_event_peek' : return type 'Event' not supported

// UNSUPPORTED : C value 'gdk_event_request_motions' : has parameters

// UNSUPPORTED : C value 'gdk_events_get_angle' : has parameters

// UNSUPPORTED : C value 'gdk_events_get_center' : has parameters

// UNSUPPORTED : C value 'gdk_events_get_distance' : has parameters

// UNSUPPORTED : C value 'gdk_events_pending' : return type 'gboolean' not supported

var flushInvoker *gi.Function

// Flush is a representation of the C type gdk_flush.
func Flush() {
	if flushInvoker == nil {
		flushInvoker = gi.FunctionInvokerNew("Gdk", "flush")
	}

	flushInvoker.Invoke()
}

// UNSUPPORTED : C value 'gdk_get_default_root_window' : return type 'Window' not supported

var getDisplayInvoker *gi.Function

// GetDisplay is a representation of the C type gdk_get_display.
func GetDisplay() string {
	if getDisplayInvoker == nil {
		getDisplayInvoker = gi.FunctionInvokerNew("Gdk", "get_display")
	}

	ret := getDisplayInvoker.Invoke()
	return ret.String()
}

var getDisplayArgNameInvoker *gi.Function

// GetDisplayArgName is a representation of the C type gdk_get_display_arg_name.
func GetDisplayArgName() string {
	if getDisplayArgNameInvoker == nil {
		getDisplayArgNameInvoker = gi.FunctionInvokerNew("Gdk", "get_display_arg_name")
	}

	ret := getDisplayArgNameInvoker.Invoke()
	return ret.String()
}

var getProgramClassInvoker *gi.Function

// GetProgramClass is a representation of the C type gdk_get_program_class.
func GetProgramClass() string {
	if getProgramClassInvoker == nil {
		getProgramClassInvoker = gi.FunctionInvokerNew("Gdk", "get_program_class")
	}

	ret := getProgramClassInvoker.Invoke()
	return ret.String()
}

// UNSUPPORTED : C value 'gdk_get_show_events' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'gdk_gl_error_quark' : return type 'GLib.Quark' not supported

// UNSUPPORTED : C value 'gdk_init' : has parameters

// UNSUPPORTED : C value 'gdk_init_check' : has parameters

// UNSUPPORTED : C value 'gdk_keyboard_grab' : has parameters

// UNSUPPORTED : C value 'gdk_keyboard_ungrab' : has parameters

// UNSUPPORTED : C value 'gdk_keyval_convert_case' : has parameters

// UNSUPPORTED : C value 'gdk_keyval_from_name' : has parameters

// UNSUPPORTED : C value 'gdk_keyval_is_lower' : has parameters

// UNSUPPORTED : C value 'gdk_keyval_is_upper' : has parameters

// UNSUPPORTED : C value 'gdk_keyval_name' : has parameters

// UNSUPPORTED : C value 'gdk_keyval_to_lower' : has parameters

// UNSUPPORTED : C value 'gdk_keyval_to_unicode' : has parameters

// UNSUPPORTED : C value 'gdk_keyval_to_upper' : has parameters

// UNSUPPORTED : C value 'gdk_list_visuals' : return type 'GLib.List' not supported

var notifyStartupCompleteInvoker *gi.Function

// NotifyStartupComplete is a representation of the C type gdk_notify_startup_complete.
func NotifyStartupComplete() {
	if notifyStartupCompleteInvoker == nil {
		notifyStartupCompleteInvoker = gi.FunctionInvokerNew("Gdk", "notify_startup_complete")
	}

	notifyStartupCompleteInvoker.Invoke()
}

// UNSUPPORTED : C value 'gdk_notify_startup_complete_with_id' : has parameters

// UNSUPPORTED : C value 'gdk_offscreen_window_get_embedder' : has parameters

// UNSUPPORTED : C value 'gdk_offscreen_window_get_surface' : has parameters

// UNSUPPORTED : C value 'gdk_offscreen_window_set_embedder' : has parameters

// UNSUPPORTED : C value 'gdk_pango_context_get' : return type 'Pango.Context' not supported

// UNSUPPORTED : C value 'gdk_pango_context_get_for_display' : has parameters

// UNSUPPORTED : C value 'gdk_pango_context_get_for_screen' : has parameters

// UNSUPPORTED : C value 'gdk_pango_layout_get_clip_region' : has parameters

// UNSUPPORTED : C value 'gdk_pango_layout_line_get_clip_region' : has parameters

// UNSUPPORTED : C value 'gdk_parse_args' : has parameters

// UNSUPPORTED : C value 'gdk_pixbuf_get_from_surface' : has parameters

// UNSUPPORTED : C value 'gdk_pixbuf_get_from_window' : has parameters

// UNSUPPORTED : C value 'gdk_pointer_grab' : has parameters

// UNSUPPORTED : C value 'gdk_pointer_is_grabbed' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'gdk_pointer_ungrab' : has parameters

var preParseLibgtkOnlyInvoker *gi.Function

// PreParseLibgtkOnly is a representation of the C type gdk_pre_parse_libgtk_only.
func PreParseLibgtkOnly() {
	if preParseLibgtkOnlyInvoker == nil {
		preParseLibgtkOnlyInvoker = gi.FunctionInvokerNew("Gdk", "pre_parse_libgtk_only")
	}

	preParseLibgtkOnlyInvoker.Invoke()
}

// UNSUPPORTED : C value 'gdk_property_change' : has parameters

// UNSUPPORTED : C value 'gdk_property_delete' : has parameters

// UNSUPPORTED : C value 'gdk_property_get' : has parameters

// UNSUPPORTED : C value 'gdk_query_depths' : has parameters

// UNSUPPORTED : C value 'gdk_query_visual_types' : has parameters

// UNSUPPORTED : C value 'gdk_selection_convert' : has parameters

// UNSUPPORTED : C value 'gdk_selection_owner_get' : has parameters

// UNSUPPORTED : C value 'gdk_selection_owner_get_for_display' : has parameters

// UNSUPPORTED : C value 'gdk_selection_owner_set' : has parameters

// UNSUPPORTED : C value 'gdk_selection_owner_set_for_display' : has parameters

// UNSUPPORTED : C value 'gdk_selection_property_get' : has parameters

// UNSUPPORTED : C value 'gdk_selection_send_notify' : has parameters

// UNSUPPORTED : C value 'gdk_selection_send_notify_for_display' : has parameters

// UNSUPPORTED : C value 'gdk_set_allowed_backends' : has parameters

// UNSUPPORTED : C value 'gdk_set_double_click_time' : has parameters

// UNSUPPORTED : C value 'gdk_set_program_class' : has parameters

// UNSUPPORTED : C value 'gdk_set_show_events' : has parameters

// UNSUPPORTED : C value 'gdk_setting_get' : has parameters

// UNSUPPORTED : C value 'gdk_synthesize_window_state' : has parameters

// UNSUPPORTED : C value 'gdk_test_render_sync' : has parameters

// UNSUPPORTED : C value 'gdk_test_simulate_button' : has parameters

// UNSUPPORTED : C value 'gdk_test_simulate_key' : has parameters

// UNSUPPORTED : C value 'gdk_text_property_to_utf8_list_for_display' : has parameters

// UNSUPPORTED : C value 'gdk_threads_add_idle' : has parameters

// UNSUPPORTED : C value 'gdk_threads_add_idle_full' : has parameters

// UNSUPPORTED : C value 'gdk_threads_add_timeout' : has parameters

// UNSUPPORTED : C value 'gdk_threads_add_timeout_full' : has parameters

// UNSUPPORTED : C value 'gdk_threads_add_timeout_seconds' : has parameters

// UNSUPPORTED : C value 'gdk_threads_add_timeout_seconds_full' : has parameters

var threadsEnterInvoker *gi.Function

// ThreadsEnter is a representation of the C type gdk_threads_enter.
func ThreadsEnter() {
	if threadsEnterInvoker == nil {
		threadsEnterInvoker = gi.FunctionInvokerNew("Gdk", "threads_enter")
	}

	threadsEnterInvoker.Invoke()
}

var threadsInitInvoker *gi.Function

// ThreadsInit is a representation of the C type gdk_threads_init.
func ThreadsInit() {
	if threadsInitInvoker == nil {
		threadsInitInvoker = gi.FunctionInvokerNew("Gdk", "threads_init")
	}

	threadsInitInvoker.Invoke()
}

var threadsLeaveInvoker *gi.Function

// ThreadsLeave is a representation of the C type gdk_threads_leave.
func ThreadsLeave() {
	if threadsLeaveInvoker == nil {
		threadsLeaveInvoker = gi.FunctionInvokerNew("Gdk", "threads_leave")
	}

	threadsLeaveInvoker.Invoke()
}

// UNSUPPORTED : C value 'gdk_threads_set_lock_functions' : has parameters

// UNSUPPORTED : C value 'gdk_unicode_to_keyval' : has parameters

// UNSUPPORTED : C value 'gdk_utf8_to_string_target' : has parameters
