// Code generated - DO NOT EDIT.
// +build gdk_2.8

package gdk

import "unsafe"

// #include <gdk/gdk.h>
import "C"

type Atom C.GdkAtom
type Color C.GdkColor
type DevicePadInterface C.GdkDevicePadInterface
type DrawingContextClass C.GdkDrawingContextClass
type EventAny C.GdkEventAny
type EventButton C.GdkEventButton
type EventConfigure C.GdkEventConfigure
type EventCrossing C.GdkEventCrossing
type EventDND C.GdkEventDND
type EventExpose C.GdkEventExpose
type EventFocus C.GdkEventFocus
type EventGrabBroken C.GdkEventGrabBroken
type EventKey C.GdkEventKey
type EventMotion C.GdkEventMotion
type EventOwnerChange C.GdkEventOwnerChange
type EventProperty C.GdkEventProperty
type EventProximity C.GdkEventProximity
type EventScroll C.GdkEventScroll
type EventSelection C.GdkEventSelection
type EventSequence C.GdkEventSequence
type EventSetting C.GdkEventSetting
type EventTouch C.GdkEventTouch
type EventTouchpadPinch C.GdkEventTouchpadPinch
type EventTouchpadSwipe C.GdkEventTouchpadSwipe
type EventVisibility C.GdkEventVisibility
type EventWindowState C.GdkEventWindowState
type FrameClockClass C.GdkFrameClockClass
type FrameClockPrivate C.GdkFrameClockPrivate
type FrameTimings C.GdkFrameTimings
type Geometry C.GdkGeometry
type KeymapKey C.GdkKeymapKey
type MonitorClass C.GdkMonitorClass
type Point C.GdkPoint
type RGBA C.GdkRGBA
type Rectangle C.GdkRectangle
type TimeCoord C.GdkTimeCoord
type WindowAttr C.GdkWindowAttr
type WindowClass C.GdkWindowClass
type WindowRedirect C.GdkWindowRedirect

func Fn_gdk_add_option_entries_libgtk_only(param0 unsafe.Pointer) {}

func Fn_gdk_atom_intern(param0 string, param1 bool) {}

func Fn_gdk_beep() {
	C.gdk_beep()
}

func Fn_gdk_cairo_create(param0 unsafe.Pointer) {}

func Fn_gdk_cairo_get_clip_rectangle(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_gdk_cairo_rectangle(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_gdk_cairo_region(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_gdk_cairo_region_create_from_surface(param0 unsafe.Pointer) {}

func Fn_gdk_cairo_set_source_color(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_gdk_cairo_set_source_pixbuf(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 float64, param3 float64) {
}

func Fn_gdk_color_parse(param0 string, param1 unsafe.Pointer) {}

func Fn_gdk_drag_abort(param0 unsafe.Pointer, param1 uint32) {}

func Fn_gdk_drag_begin(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_gdk_drag_begin_for_device(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer) {
}

func Fn_gdk_drag_drop(param0 unsafe.Pointer, param1 uint32) {}

func Fn_gdk_drag_drop_succeeded(param0 unsafe.Pointer) {}

func Fn_gdk_drag_find_window_for_screen(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer, param3 int, param4 int, param5 *unsafe.Pointer, param6 int) {
}

func Fn_gdk_drag_get_selection(param0 unsafe.Pointer) {}

func Fn_gdk_drag_motion(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int, param3 int, param4 int, param5 int, param6 int, param7 uint32) {
}

func Fn_gdk_drag_status(param0 unsafe.Pointer, param1 int, param2 uint32) {}

func Fn_gdk_drop_finish(param0 unsafe.Pointer, param1 bool, param2 uint32) {}

func Fn_gdk_drop_reply(param0 unsafe.Pointer, param1 bool, param2 uint32) {}

func Fn_gdk_error_trap_pop() {
	C.gdk_error_trap_pop()
}

func Fn_gdk_error_trap_push() {
	C.gdk_error_trap_push()
}

func Fn_gdk_event_get() {
	C.gdk_event_get()
}

// UNSUPPORTED : event_handler_set : has callback

func Fn_gdk_event_peek() {
	C.gdk_event_peek()
}

func Fn_gdk_events_pending() {
	C.gdk_events_pending()
}

func Fn_gdk_flush() {
	C.gdk_flush()
}

func Fn_gdk_get_default_root_window() {
	C.gdk_get_default_root_window()
}

func Fn_gdk_get_display() {
	C.gdk_get_display()
}

func Fn_gdk_get_display_arg_name() {
	C.gdk_get_display_arg_name()
}

func Fn_gdk_get_program_class() {
	C.gdk_get_program_class()
}

func Fn_gdk_get_show_events() {
	C.gdk_get_show_events()
}

func Fn_gdk_gl_error_quark() {
	C.gdk_gl_error_quark()
}

func Fn_gdk_init(param0 *int, param1 *[]string) {}

func Fn_gdk_init_check(param0 *int, param1 *[]string) {}

func Fn_gdk_keyboard_grab(param0 unsafe.Pointer, param1 bool, param2 uint32) {}

func Fn_gdk_keyboard_ungrab(param0 uint32) {}

func Fn_gdk_keyval_convert_case(param0 uint, param1 *uint, param2 *uint) {}

func Fn_gdk_keyval_from_name(param0 string) {}

func Fn_gdk_keyval_is_lower(param0 uint) {}

func Fn_gdk_keyval_is_upper(param0 uint) {}

func Fn_gdk_keyval_name(param0 uint) {}

func Fn_gdk_keyval_to_lower(param0 uint) {}

func Fn_gdk_keyval_to_unicode(param0 uint) {}

func Fn_gdk_keyval_to_upper(param0 uint) {}

func Fn_gdk_list_visuals() {
	C.gdk_list_visuals()
}

func Fn_gdk_notify_startup_complete() {
	C.gdk_notify_startup_complete()
}

func Fn_gdk_offscreen_window_get_surface(param0 unsafe.Pointer) {}

func Fn_gdk_pango_context_get() {
	C.gdk_pango_context_get()
}

func Fn_gdk_pango_context_get_for_screen(param0 unsafe.Pointer) {}

func Fn_gdk_pango_layout_get_clip_region(param0 unsafe.Pointer, param1 int, param2 int, param3 *int, param4 int) {
}

func Fn_gdk_pango_layout_line_get_clip_region(param0 unsafe.Pointer, param1 int, param2 int, param3 []int, param4 int) {
}

func Fn_gdk_parse_args(param0 *int, param1 *[]string) {}

func Fn_gdk_pixbuf_get_from_surface(param0 unsafe.Pointer, param1 int, param2 int, param3 int, param4 int) {
}

func Fn_gdk_pixbuf_get_from_window(param0 unsafe.Pointer, param1 int, param2 int, param3 int, param4 int) {
}

func Fn_gdk_pointer_grab(param0 unsafe.Pointer, param1 bool, param2 int, param3 unsafe.Pointer, param4 unsafe.Pointer, param5 uint32) {
}

func Fn_gdk_pointer_is_grabbed() {
	C.gdk_pointer_is_grabbed()
}

func Fn_gdk_pointer_ungrab(param0 uint32) {}

func Fn_gdk_pre_parse_libgtk_only() {
	C.gdk_pre_parse_libgtk_only()
}

func Fn_gdk_property_change(param0 unsafe.Pointer, param1 Atom, param2 Atom, param3 int, param4 int, param5 *uint8, param6 int) {
}

func Fn_gdk_property_delete(param0 unsafe.Pointer, param1 Atom) {}

func Fn_gdk_property_get(param0 unsafe.Pointer, param1 Atom, param2 Atom, param3 uint64, param4 uint64, param5 int, param6 unsafe.Pointer, param7 *int, param8 *int, param9 []*uint8) {
}

func Fn_gdk_query_depths(param0 []*int, param1 *int) {}

func Fn_gdk_query_visual_types(param0 []int, param1 *int) {}

func Fn_gdk_selection_convert(param0 unsafe.Pointer, param1 Atom, param2 Atom, param3 uint32) {}

func Fn_gdk_selection_owner_get(param0 Atom) {}

func Fn_gdk_selection_owner_get_for_display(param0 unsafe.Pointer, param1 Atom) {}

func Fn_gdk_selection_owner_set(param0 unsafe.Pointer, param1 Atom, param2 uint32, param3 bool) {}

func Fn_gdk_selection_owner_set_for_display(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 Atom, param3 uint32, param4 bool) {
}

func Fn_gdk_selection_property_get(param0 unsafe.Pointer, param1 **uint8, param2 unsafe.Pointer, param3 *int) {
}

func Fn_gdk_selection_send_notify(param0 unsafe.Pointer, param1 Atom, param2 Atom, param3 Atom, param4 uint32) {
}

func Fn_gdk_selection_send_notify_for_display(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 Atom, param3 Atom, param4 Atom, param5 uint32) {
}

func Fn_gdk_set_double_click_time(param0 uint) {}

func Fn_gdk_set_program_class(param0 string) {}

func Fn_gdk_set_show_events(param0 bool) {}

func Fn_gdk_setting_get(param0 string, param1 unsafe.Pointer) {}

func Fn_gdk_synthesize_window_state(param0 unsafe.Pointer, param1 int, param2 int) {}

func Fn_gdk_text_property_to_utf8_list_for_display(param0 unsafe.Pointer, param1 Atom, param2 int, param3 []uint8, param4 int, param5 *[]string) {
}

// UNSUPPORTED : threads_add_idle : has callback

// UNSUPPORTED : threads_add_idle_full : has callback

// UNSUPPORTED : threads_add_timeout : has callback

// UNSUPPORTED : threads_add_timeout_full : has callback

// UNSUPPORTED : threads_add_timeout_seconds : has callback

// UNSUPPORTED : threads_add_timeout_seconds_full : has callback

func Fn_gdk_threads_enter() {
	C.gdk_threads_enter()
}

func Fn_gdk_threads_init() {
	C.gdk_threads_init()
}

func Fn_gdk_threads_leave() {
	C.gdk_threads_leave()
}

// UNSUPPORTED : threads_set_lock_functions : has callback

func Fn_gdk_unicode_to_keyval(param0 uint32) {}

func Fn_gdk_utf8_to_string_target(param0 string) {}

func Fn_gdk_cursor_new(param0 int) {}

func Fn_gdk_cursor_new_for_display(param0 unsafe.Pointer, param1 int) {}

func Fn_gdk_cursor_new_from_name(param0 unsafe.Pointer, param1 string) {}

func Fn_gdk_cursor_new_from_pixbuf(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int, param3 int) {
}

func Fn_gdk_cursor_get_display() {}

func Fn_gdk_cursor_get_image() {}

func Fn_gdk_cursor_ref() {}

func Fn_gdk_cursor_unref() {}

func Fn_gdk_device_get_axis(param0 []float64, param1 int, param2 *float64) {}

func Fn_gdk_device_get_history(param0 unsafe.Pointer, param1 uint32, param2 uint32, param3 []*unsafe.Pointer, param4 *int) {
}

func Fn_gdk_device_get_state(param0 unsafe.Pointer, param1 []float64, param2 int) {}

func Fn_gdk_device_list_slave_devices() {}

func Fn_gdk_device_set_axis_use(param0 uint, param1 int) {}

func Fn_gdk_device_set_key(param0 uint, param1 uint, param2 int) {}

func Fn_gdk_device_set_mode(param0 int) {}

func Fn_gdk_device_free_history(param0 []unsafe.Pointer, param1 int) {}

func Fn_gdk_device_grab_info_libgtk_only(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 *unsafe.Pointer, param3 *bool) {
}

func Fn_gdk_display_beep() {}

func Fn_gdk_display_close() {}

func Fn_gdk_display_device_is_grabbed(param0 unsafe.Pointer) {}

func Fn_gdk_display_flush() {}

func Fn_gdk_display_get_default_cursor_size() {}

func Fn_gdk_display_get_default_group() {}

func Fn_gdk_display_get_default_screen() {}

func Fn_gdk_display_get_event() {}

func Fn_gdk_display_get_maximal_cursor_size(param0 *uint, param1 *uint) {}

func Fn_gdk_display_get_n_screens() {}

func Fn_gdk_display_get_name() {}

func Fn_gdk_display_get_pointer(param0 *unsafe.Pointer, param1 *int, param2 *int, param3 int) {}

func Fn_gdk_display_get_screen(param0 int) {}

func Fn_gdk_display_get_window_at_pointer(param0 *int, param1 *int) {}

func Fn_gdk_display_keyboard_ungrab(param0 uint32) {}

func Fn_gdk_display_list_devices() {}

func Fn_gdk_display_peek_event() {}

func Fn_gdk_display_pointer_is_grabbed() {}

func Fn_gdk_display_pointer_ungrab(param0 uint32) {}

func Fn_gdk_display_put_event(param0 unsafe.Pointer) {}

func Fn_gdk_display_request_selection_notification(param0 Atom) {}

func Fn_gdk_display_set_double_click_distance(param0 uint) {}

func Fn_gdk_display_set_double_click_time(param0 uint) {}

func Fn_gdk_display_store_clipboard(param0 unsafe.Pointer, param1 uint32, param2 []Atom, param3 int) {
}

func Fn_gdk_display_supports_clipboard_persistence() {}

func Fn_gdk_display_supports_cursor_alpha() {}

func Fn_gdk_display_supports_cursor_color() {}

func Fn_gdk_display_supports_selection_notification() {}

func Fn_gdk_display_sync() {}

func Fn_gdk_display_warp_pointer(param0 unsafe.Pointer, param1 int, param2 int) {}

func Fn_gdk_display_get_default() {
	C.gdk_display_get_default()
}

func Fn_gdk_display_open(param0 string) {}

func Fn_gdk_display_open_default_libgtk_only() {
	C.gdk_display_open_default_libgtk_only()
}

func Fn_gdk_display_manager_get_default_display() {}

func Fn_gdk_display_manager_list_displays() {}

func Fn_gdk_display_manager_set_default_display(param0 unsafe.Pointer) {}

func Fn_gdk_display_manager_get() {
	C.gdk_display_manager_get()
}

func Fn_gdk_drag_context_get_device() {}

func Fn_gdk_drag_context_set_device(param0 unsafe.Pointer) {}

func Fn_gdk_keymap_get_direction() {}

func Fn_gdk_keymap_get_entries_for_keycode(param0 uint, param1 []unsafe.Pointer, param2 []*uint, param3 *int) {
}

func Fn_gdk_keymap_get_entries_for_keyval(param0 uint, param1 []unsafe.Pointer, param2 *int) {}

func Fn_gdk_keymap_lookup_key(param0 unsafe.Pointer) {}

func Fn_gdk_keymap_translate_keyboard_state(param0 uint, param1 int, param2 int, param3 *uint, param4 *int, param5 *int, param6 int) {
}

func Fn_gdk_keymap_get_default() {
	C.gdk_keymap_get_default()
}

func Fn_gdk_keymap_get_for_display(param0 unsafe.Pointer) {}

func Fn_gdk_monitor_get_manufacturer() {}

func Fn_gdk_monitor_get_model() {}

func Fn_gdk_screen_get_display() {}

func Fn_gdk_screen_get_height() {}

func Fn_gdk_screen_get_height_mm() {}

func Fn_gdk_screen_get_monitor_at_point(param0 int, param1 int) {}

func Fn_gdk_screen_get_monitor_at_window(param0 unsafe.Pointer) {}

func Fn_gdk_screen_get_monitor_geometry(param0 int, param1 unsafe.Pointer) {}

func Fn_gdk_screen_get_n_monitors() {}

func Fn_gdk_screen_get_number() {}

func Fn_gdk_screen_get_rgba_visual() {}

func Fn_gdk_screen_get_root_window() {}

func Fn_gdk_screen_get_setting(param0 string, param1 unsafe.Pointer) {}

func Fn_gdk_screen_get_system_visual() {}

func Fn_gdk_screen_get_toplevel_windows() {}

func Fn_gdk_screen_get_width() {}

func Fn_gdk_screen_get_width_mm() {}

func Fn_gdk_screen_list_visuals() {}

func Fn_gdk_screen_make_display_name() {}

func Fn_gdk_screen_get_default() {
	C.gdk_screen_get_default()
}

func Fn_gdk_screen_height() {
	C.gdk_screen_height()
}

func Fn_gdk_screen_height_mm() {
	C.gdk_screen_height_mm()
}

func Fn_gdk_screen_width() {
	C.gdk_screen_width()
}

func Fn_gdk_screen_width_mm() {
	C.gdk_screen_width_mm()
}

func Fn_gdk_seat_get_display() {}

// UNSUPPORTED : grab : has callback

func Fn_gdk_visual_get_screen() {}

func Fn_gdk_visual_get_best() {
	C.gdk_visual_get_best()
}

func Fn_gdk_visual_get_best_depth() {
	C.gdk_visual_get_best_depth()
}

func Fn_gdk_visual_get_best_type() {
	C.gdk_visual_get_best_type()
}

func Fn_gdk_visual_get_best_with_both(param0 int, param1 int) {}

func Fn_gdk_visual_get_best_with_depth(param0 int) {}

func Fn_gdk_visual_get_best_with_type(param0 int) {}

func Fn_gdk_visual_get_system() {
	C.gdk_visual_get_system()
}

func Fn_gdk_window_new(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int) {}

// UNSUPPORTED : add_filter : has callback

func Fn_gdk_window_begin_move_drag(param0 int, param1 int, param2 int, param3 uint32) {}

func Fn_gdk_window_begin_paint_rect(param0 unsafe.Pointer) {}

func Fn_gdk_window_begin_paint_region(param0 unsafe.Pointer) {}

func Fn_gdk_window_begin_resize_drag(param0 int, param1 int, param2 int, param3 int, param4 uint32) {}

func Fn_gdk_window_configure_finished() {}

func Fn_gdk_window_deiconify() {}

func Fn_gdk_window_destroy() {}

func Fn_gdk_window_destroy_notify() {}

func Fn_gdk_window_enable_synchronized_configure() {}

func Fn_gdk_window_end_paint() {}

func Fn_gdk_window_focus(param0 uint32) {}

func Fn_gdk_window_freeze_toplevel_updates_libgtk_only() {}

func Fn_gdk_window_freeze_updates() {}

func Fn_gdk_window_fullscreen() {}

func Fn_gdk_window_fullscreen_on_monitor(param0 int) {}

func Fn_gdk_window_get_children() {}

func Fn_gdk_window_get_clip_region() {}

func Fn_gdk_window_get_decorations(param0 int) {}

func Fn_gdk_window_get_events() {}

func Fn_gdk_window_get_frame_extents(param0 unsafe.Pointer) {}

func Fn_gdk_window_get_geometry(param0 *int, param1 *int, param2 *int, param3 *int) {}

func Fn_gdk_window_get_group() {}

func Fn_gdk_window_get_origin(param0 *int, param1 *int) {}

func Fn_gdk_window_get_parent() {}

func Fn_gdk_window_get_pointer(param0 *int, param1 *int, param2 int) {}

func Fn_gdk_window_get_position(param0 *int, param1 *int) {}

func Fn_gdk_window_get_root_origin(param0 *int, param1 *int) {}

func Fn_gdk_window_get_source_events(param0 int) {}

func Fn_gdk_window_get_state() {}

func Fn_gdk_window_get_toplevel() {}

func Fn_gdk_window_get_update_area() {}

func Fn_gdk_window_get_user_data(param0 *unsafe.Pointer) {}

func Fn_gdk_window_get_visible_region() {}

func Fn_gdk_window_get_window_type() {}

func Fn_gdk_window_hide() {}

func Fn_gdk_window_iconify() {}

// UNSUPPORTED : invalidate_maybe_recurse : has callback

func Fn_gdk_window_invalidate_rect(param0 unsafe.Pointer, param1 bool) {}

func Fn_gdk_window_invalidate_region(param0 unsafe.Pointer, param1 bool) {}

func Fn_gdk_window_is_viewable() {}

func Fn_gdk_window_is_visible() {}

func Fn_gdk_window_lower() {}

func Fn_gdk_window_maximize() {}

func Fn_gdk_window_merge_child_shapes() {}

func Fn_gdk_window_move(param0 int, param1 int) {}

func Fn_gdk_window_move_region(param0 unsafe.Pointer, param1 int, param2 int) {}

func Fn_gdk_window_move_resize(param0 int, param1 int, param2 int, param3 int) {}

func Fn_gdk_window_peek_children() {}

func Fn_gdk_window_process_updates(param0 bool) {}

func Fn_gdk_window_raise() {}

func Fn_gdk_window_register_dnd() {}

// UNSUPPORTED : remove_filter : has callback

func Fn_gdk_window_reparent(param0 unsafe.Pointer, param1 int, param2 int) {}

func Fn_gdk_window_resize(param0 int, param1 int) {}

func Fn_gdk_window_scroll(param0 int, param1 int) {}

func Fn_gdk_window_set_accept_focus(param0 bool) {}

func Fn_gdk_window_set_background(param0 unsafe.Pointer) {}

func Fn_gdk_window_set_background_pattern(param0 unsafe.Pointer) {}

func Fn_gdk_window_set_background_rgba(param0 unsafe.Pointer) {}

func Fn_gdk_window_set_child_shapes() {}

func Fn_gdk_window_set_cursor(param0 unsafe.Pointer) {}

func Fn_gdk_window_set_decorations(param0 int) {}

func Fn_gdk_window_set_events(param0 int) {}

func Fn_gdk_window_set_focus_on_map(param0 bool) {}

func Fn_gdk_window_set_functions(param0 int) {}

func Fn_gdk_window_set_geometry_hints(param0 unsafe.Pointer, param1 int) {}

func Fn_gdk_window_set_group(param0 unsafe.Pointer) {}

func Fn_gdk_window_set_icon_list(param0 unsafe.Pointer) {}

func Fn_gdk_window_set_icon_name(param0 string) {}

// UNSUPPORTED : set_invalidate_handler : has callback

func Fn_gdk_window_set_keep_above(param0 bool) {}

func Fn_gdk_window_set_keep_below(param0 bool) {}

func Fn_gdk_window_set_modal_hint(param0 bool) {}

func Fn_gdk_window_set_override_redirect(param0 bool) {}

func Fn_gdk_window_set_role(param0 string) {}

func Fn_gdk_window_set_skip_pager_hint(param0 bool) {}

func Fn_gdk_window_set_skip_taskbar_hint(param0 bool) {}

func Fn_gdk_window_set_static_gravities(param0 bool) {}

func Fn_gdk_window_set_title(param0 string) {}

func Fn_gdk_window_set_transient_for(param0 unsafe.Pointer) {}

func Fn_gdk_window_set_type_hint(param0 int) {}

func Fn_gdk_window_set_urgency_hint(param0 bool) {}

func Fn_gdk_window_set_user_data(param0 *unsafe.Pointer) {}

func Fn_gdk_window_shape_combine_region(param0 unsafe.Pointer, param1 int, param2 int) {}

func Fn_gdk_window_show() {}

func Fn_gdk_window_show_unraised() {}

func Fn_gdk_window_stick() {}

func Fn_gdk_window_thaw_toplevel_updates_libgtk_only() {}

func Fn_gdk_window_thaw_updates() {}

func Fn_gdk_window_unfullscreen() {}

func Fn_gdk_window_unmaximize() {}

func Fn_gdk_window_unstick() {}

func Fn_gdk_window_withdraw() {}

func Fn_gdk_window_at_pointer(param0 *int, param1 *int) {}

func Fn_gdk_window_constrain_size(param0 unsafe.Pointer, param1 int, param2 int, param3 int, param4 *int, param5 *int) {
}

func Fn_gdk_window_process_all_updates() {
	C.gdk_window_process_all_updates()
}

func Fn_gdk_window_set_debug_updates(param0 bool) {}
