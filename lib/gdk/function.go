// Code generated - DO NOT EDIT.

package gdk

import gi "github.com/pekim/gobbi/internal/gi"

// UNSUPPORTED : C value 'gdk_add_option_entries_libgtk_only' : parameter 'group' of type 'GLib.OptionGroup' not supported

// UNSUPPORTED : C value 'gdk_atom_intern' : parameter 'atom_name' of type 'utf8' not supported

// UNSUPPORTED : C value 'gdk_atom_intern_static_string' : parameter 'atom_name' of type 'utf8' not supported

var beepInvoker *gi.Function

// Beep is a representation of the C type gdk_beep.
func Beep() {
	if beepInvoker == nil {
		beepInvoker = gi.FunctionInvokerNew("Gdk", "beep")
	}

	beepInvoker.Invoke(nil)
}

// UNSUPPORTED : C value 'gdk_cairo_create' : parameter 'window' of type 'Window' not supported

// UNSUPPORTED : C value 'gdk_cairo_draw_from_gl' : parameter 'cr' of type 'cairo.Context' not supported

// UNSUPPORTED : C value 'gdk_cairo_get_clip_rectangle' : parameter 'cr' of type 'cairo.Context' not supported

// UNSUPPORTED : C value 'gdk_cairo_get_drawing_context' : parameter 'cr' of type 'cairo.Context' not supported

// UNSUPPORTED : C value 'gdk_cairo_rectangle' : parameter 'cr' of type 'cairo.Context' not supported

// UNSUPPORTED : C value 'gdk_cairo_region' : parameter 'cr' of type 'cairo.Context' not supported

// UNSUPPORTED : C value 'gdk_cairo_region_create_from_surface' : parameter 'surface' of type 'cairo.Surface' not supported

// UNSUPPORTED : C value 'gdk_cairo_set_source_color' : parameter 'cr' of type 'cairo.Context' not supported

// UNSUPPORTED : C value 'gdk_cairo_set_source_pixbuf' : parameter 'cr' of type 'cairo.Context' not supported

// UNSUPPORTED : C value 'gdk_cairo_set_source_rgba' : parameter 'cr' of type 'cairo.Context' not supported

// UNSUPPORTED : C value 'gdk_cairo_set_source_window' : parameter 'cr' of type 'cairo.Context' not supported

// UNSUPPORTED : C value 'gdk_cairo_surface_create_from_pixbuf' : parameter 'pixbuf' of type 'GdkPixbuf.Pixbuf' not supported

// UNSUPPORTED : C value 'gdk_color_parse' : parameter 'spec' of type 'utf8' not supported

var disableMultideviceInvoker *gi.Function

// DisableMultidevice is a representation of the C type gdk_disable_multidevice.
func DisableMultidevice() {
	if disableMultideviceInvoker == nil {
		disableMultideviceInvoker = gi.FunctionInvokerNew("Gdk", "disable_multidevice")
	}

	disableMultideviceInvoker.Invoke(nil)
}

// UNSUPPORTED : C value 'gdk_drag_abort' : parameter 'context' of type 'DragContext' not supported

// UNSUPPORTED : C value 'gdk_drag_begin' : parameter 'window' of type 'Window' not supported

// UNSUPPORTED : C value 'gdk_drag_begin_for_device' : parameter 'window' of type 'Window' not supported

// UNSUPPORTED : C value 'gdk_drag_begin_from_point' : parameter 'window' of type 'Window' not supported

// UNSUPPORTED : C value 'gdk_drag_drop' : parameter 'context' of type 'DragContext' not supported

// UNSUPPORTED : C value 'gdk_drag_drop_done' : parameter 'context' of type 'DragContext' not supported

// UNSUPPORTED : C value 'gdk_drag_drop_succeeded' : parameter 'context' of type 'DragContext' not supported

// UNSUPPORTED : C value 'gdk_drag_find_window_for_screen' : parameter 'context' of type 'DragContext' not supported

// UNSUPPORTED : C value 'gdk_drag_get_selection' : parameter 'context' of type 'DragContext' not supported

// UNSUPPORTED : C value 'gdk_drag_motion' : parameter 'context' of type 'DragContext' not supported

// UNSUPPORTED : C value 'gdk_drag_status' : parameter 'context' of type 'DragContext' not supported

// UNSUPPORTED : C value 'gdk_drop_finish' : parameter 'context' of type 'DragContext' not supported

// UNSUPPORTED : C value 'gdk_drop_reply' : parameter 'context' of type 'DragContext' not supported

var errorTrapPopInvoker *gi.Function

// ErrorTrapPop is a representation of the C type gdk_error_trap_pop.
func ErrorTrapPop() int32 {
	if errorTrapPopInvoker == nil {
		errorTrapPopInvoker = gi.FunctionInvokerNew("Gdk", "error_trap_pop")
	}

	ret := errorTrapPopInvoker.Invoke(nil)
	return ret.Int32()
}

var errorTrapPopIgnoredInvoker *gi.Function

// ErrorTrapPopIgnored is a representation of the C type gdk_error_trap_pop_ignored.
func ErrorTrapPopIgnored() {
	if errorTrapPopIgnoredInvoker == nil {
		errorTrapPopIgnoredInvoker = gi.FunctionInvokerNew("Gdk", "error_trap_pop_ignored")
	}

	errorTrapPopIgnoredInvoker.Invoke(nil)
}

var errorTrapPushInvoker *gi.Function

// ErrorTrapPush is a representation of the C type gdk_error_trap_push.
func ErrorTrapPush() {
	if errorTrapPushInvoker == nil {
		errorTrapPushInvoker = gi.FunctionInvokerNew("Gdk", "error_trap_push")
	}

	errorTrapPushInvoker.Invoke(nil)
}

// UNSUPPORTED : C value 'gdk_event_get' : return type 'Event' not supported

// UNSUPPORTED : C value 'gdk_event_handler_set' : parameter 'func' of type 'EventFunc' not supported

// UNSUPPORTED : C value 'gdk_event_peek' : return type 'Event' not supported

// UNSUPPORTED : C value 'gdk_event_request_motions' : parameter 'event' of type 'EventMotion' not supported

// UNSUPPORTED : C value 'gdk_events_get_angle' : parameter 'event1' of type 'Event' not supported

// UNSUPPORTED : C value 'gdk_events_get_center' : parameter 'event1' of type 'Event' not supported

// UNSUPPORTED : C value 'gdk_events_get_distance' : parameter 'event1' of type 'Event' not supported

// UNSUPPORTED : C value 'gdk_events_pending' : return type 'gboolean' not supported

var flushInvoker *gi.Function

// Flush is a representation of the C type gdk_flush.
func Flush() {
	if flushInvoker == nil {
		flushInvoker = gi.FunctionInvokerNew("Gdk", "flush")
	}

	flushInvoker.Invoke(nil)
}

// UNSUPPORTED : C value 'gdk_get_default_root_window' : return type 'Window' not supported

var getDisplayInvoker *gi.Function

// GetDisplay is a representation of the C type gdk_get_display.
func GetDisplay() string {
	if getDisplayInvoker == nil {
		getDisplayInvoker = gi.FunctionInvokerNew("Gdk", "get_display")
	}

	ret := getDisplayInvoker.Invoke(nil)
	return ret.String(true)
}

var getDisplayArgNameInvoker *gi.Function

// GetDisplayArgName is a representation of the C type gdk_get_display_arg_name.
func GetDisplayArgName() string {
	if getDisplayArgNameInvoker == nil {
		getDisplayArgNameInvoker = gi.FunctionInvokerNew("Gdk", "get_display_arg_name")
	}

	ret := getDisplayArgNameInvoker.Invoke(nil)
	return ret.String(false)
}

var getProgramClassInvoker *gi.Function

// GetProgramClass is a representation of the C type gdk_get_program_class.
func GetProgramClass() string {
	if getProgramClassInvoker == nil {
		getProgramClassInvoker = gi.FunctionInvokerNew("Gdk", "get_program_class")
	}

	ret := getProgramClassInvoker.Invoke(nil)
	return ret.String(false)
}

// UNSUPPORTED : C value 'gdk_get_show_events' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'gdk_gl_error_quark' : return type 'GLib.Quark' not supported

// UNSUPPORTED : C value 'gdk_init' : parameter 'argc' with direction 'inout' not supported

// UNSUPPORTED : C value 'gdk_init_check' : parameter 'argc' with direction 'inout' not supported

// UNSUPPORTED : C value 'gdk_keyboard_grab' : parameter 'window' of type 'Window' not supported

// UNSUPPORTED : C value 'gdk_keyboard_ungrab' : parameter 'time_' of type 'guint32' not supported

// UNSUPPORTED : C value 'gdk_keyval_convert_case' : parameter 'symbol' of type 'guint' not supported

// UNSUPPORTED : C value 'gdk_keyval_from_name' : parameter 'keyval_name' of type 'utf8' not supported

// UNSUPPORTED : C value 'gdk_keyval_is_lower' : parameter 'keyval' of type 'guint' not supported

// UNSUPPORTED : C value 'gdk_keyval_is_upper' : parameter 'keyval' of type 'guint' not supported

// UNSUPPORTED : C value 'gdk_keyval_name' : parameter 'keyval' of type 'guint' not supported

// UNSUPPORTED : C value 'gdk_keyval_to_lower' : parameter 'keyval' of type 'guint' not supported

// UNSUPPORTED : C value 'gdk_keyval_to_unicode' : parameter 'keyval' of type 'guint' not supported

// UNSUPPORTED : C value 'gdk_keyval_to_upper' : parameter 'keyval' of type 'guint' not supported

// UNSUPPORTED : C value 'gdk_list_visuals' : return type 'GLib.List' not supported

var notifyStartupCompleteInvoker *gi.Function

// NotifyStartupComplete is a representation of the C type gdk_notify_startup_complete.
func NotifyStartupComplete() {
	if notifyStartupCompleteInvoker == nil {
		notifyStartupCompleteInvoker = gi.FunctionInvokerNew("Gdk", "notify_startup_complete")
	}

	notifyStartupCompleteInvoker.Invoke(nil)
}

// UNSUPPORTED : C value 'gdk_notify_startup_complete_with_id' : parameter 'startup_id' of type 'utf8' not supported

// UNSUPPORTED : C value 'gdk_offscreen_window_get_embedder' : parameter 'window' of type 'Window' not supported

// UNSUPPORTED : C value 'gdk_offscreen_window_get_surface' : parameter 'window' of type 'Window' not supported

// UNSUPPORTED : C value 'gdk_offscreen_window_set_embedder' : parameter 'window' of type 'Window' not supported

// UNSUPPORTED : C value 'gdk_pango_context_get' : return type 'Pango.Context' not supported

// UNSUPPORTED : C value 'gdk_pango_context_get_for_display' : parameter 'display' of type 'Display' not supported

// UNSUPPORTED : C value 'gdk_pango_context_get_for_screen' : parameter 'screen' of type 'Screen' not supported

// UNSUPPORTED : C value 'gdk_pango_layout_get_clip_region' : parameter 'layout' of type 'Pango.Layout' not supported

// UNSUPPORTED : C value 'gdk_pango_layout_line_get_clip_region' : parameter 'line' of type 'Pango.LayoutLine' not supported

// UNSUPPORTED : C value 'gdk_parse_args' : parameter 'argc' with direction 'inout' not supported

// UNSUPPORTED : C value 'gdk_pixbuf_get_from_surface' : parameter 'surface' of type 'cairo.Surface' not supported

// UNSUPPORTED : C value 'gdk_pixbuf_get_from_window' : parameter 'window' of type 'Window' not supported

// UNSUPPORTED : C value 'gdk_pointer_grab' : parameter 'window' of type 'Window' not supported

// UNSUPPORTED : C value 'gdk_pointer_is_grabbed' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'gdk_pointer_ungrab' : parameter 'time_' of type 'guint32' not supported

var preParseLibgtkOnlyInvoker *gi.Function

// PreParseLibgtkOnly is a representation of the C type gdk_pre_parse_libgtk_only.
func PreParseLibgtkOnly() {
	if preParseLibgtkOnlyInvoker == nil {
		preParseLibgtkOnlyInvoker = gi.FunctionInvokerNew("Gdk", "pre_parse_libgtk_only")
	}

	preParseLibgtkOnlyInvoker.Invoke(nil)
}

// UNSUPPORTED : C value 'gdk_property_change' : parameter 'window' of type 'Window' not supported

// UNSUPPORTED : C value 'gdk_property_delete' : parameter 'window' of type 'Window' not supported

// UNSUPPORTED : C value 'gdk_property_get' : parameter 'window' of type 'Window' not supported

// UNSUPPORTED : C value 'gdk_query_depths' : parameter 'depths' has no type

// UNSUPPORTED : C value 'gdk_query_visual_types' : parameter 'visual_types' has no type

// UNSUPPORTED : C value 'gdk_selection_convert' : parameter 'requestor' of type 'Window' not supported

// UNSUPPORTED : C value 'gdk_selection_owner_get' : parameter 'selection' of type 'Atom' not supported

// UNSUPPORTED : C value 'gdk_selection_owner_get_for_display' : parameter 'display' of type 'Display' not supported

// UNSUPPORTED : C value 'gdk_selection_owner_set' : parameter 'owner' of type 'Window' not supported

// UNSUPPORTED : C value 'gdk_selection_owner_set_for_display' : parameter 'display' of type 'Display' not supported

// UNSUPPORTED : C value 'gdk_selection_property_get' : parameter 'requestor' of type 'Window' not supported

// UNSUPPORTED : C value 'gdk_selection_send_notify' : parameter 'requestor' of type 'Window' not supported

// UNSUPPORTED : C value 'gdk_selection_send_notify_for_display' : parameter 'display' of type 'Display' not supported

// UNSUPPORTED : C value 'gdk_set_allowed_backends' : parameter 'backends' of type 'utf8' not supported

// UNSUPPORTED : C value 'gdk_set_double_click_time' : parameter 'msec' of type 'guint' not supported

// UNSUPPORTED : C value 'gdk_set_program_class' : parameter 'program_class' of type 'utf8' not supported

// UNSUPPORTED : C value 'gdk_set_show_events' : parameter 'show_events' of type 'gboolean' not supported

// UNSUPPORTED : C value 'gdk_setting_get' : parameter 'name' of type 'utf8' not supported

// UNSUPPORTED : C value 'gdk_synthesize_window_state' : parameter 'window' of type 'Window' not supported

// UNSUPPORTED : C value 'gdk_test_render_sync' : parameter 'window' of type 'Window' not supported

// UNSUPPORTED : C value 'gdk_test_simulate_button' : parameter 'window' of type 'Window' not supported

// UNSUPPORTED : C value 'gdk_test_simulate_key' : parameter 'window' of type 'Window' not supported

// UNSUPPORTED : C value 'gdk_text_property_to_utf8_list_for_display' : parameter 'display' of type 'Display' not supported

// UNSUPPORTED : C value 'gdk_threads_add_idle' : parameter 'function' of type 'GLib.SourceFunc' not supported

// UNSUPPORTED : C value 'gdk_threads_add_idle_full' : parameter 'priority' of type 'gint' not supported

// UNSUPPORTED : C value 'gdk_threads_add_timeout' : parameter 'interval' of type 'guint' not supported

// UNSUPPORTED : C value 'gdk_threads_add_timeout_full' : parameter 'priority' of type 'gint' not supported

// UNSUPPORTED : C value 'gdk_threads_add_timeout_seconds' : parameter 'interval' of type 'guint' not supported

// UNSUPPORTED : C value 'gdk_threads_add_timeout_seconds_full' : parameter 'priority' of type 'gint' not supported

var threadsEnterInvoker *gi.Function

// ThreadsEnter is a representation of the C type gdk_threads_enter.
func ThreadsEnter() {
	if threadsEnterInvoker == nil {
		threadsEnterInvoker = gi.FunctionInvokerNew("Gdk", "threads_enter")
	}

	threadsEnterInvoker.Invoke(nil)
}

var threadsInitInvoker *gi.Function

// ThreadsInit is a representation of the C type gdk_threads_init.
func ThreadsInit() {
	if threadsInitInvoker == nil {
		threadsInitInvoker = gi.FunctionInvokerNew("Gdk", "threads_init")
	}

	threadsInitInvoker.Invoke(nil)
}

var threadsLeaveInvoker *gi.Function

// ThreadsLeave is a representation of the C type gdk_threads_leave.
func ThreadsLeave() {
	if threadsLeaveInvoker == nil {
		threadsLeaveInvoker = gi.FunctionInvokerNew("Gdk", "threads_leave")
	}

	threadsLeaveInvoker.Invoke(nil)
}

// UNSUPPORTED : C value 'gdk_threads_set_lock_functions' : parameter 'enter_fn' of type 'GObject.Callback' not supported

// UNSUPPORTED : C value 'gdk_unicode_to_keyval' : parameter 'wc' of type 'guint32' not supported

// UNSUPPORTED : C value 'gdk_utf8_to_string_target' : parameter 'str' of type 'utf8' not supported
