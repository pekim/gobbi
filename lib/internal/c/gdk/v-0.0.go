// Code generated - DO NOT EDIT.
// +build !gdk_2.6,!gdk_2.8,!gdk_3.4,!gdk_3.8,!gdk_3.16,!gdk_3.20,!gdk_3.22

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
type EventKey C.GdkEventKey
type EventMotion C.GdkEventMotion
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

func Fn_gdk_cairo_get_clip_rectangle(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_gdk_cairo_region_create_from_surface(param0 unsafe.Pointer) {}

func Fn_gdk_color_parse(param0 string, param1 unsafe.Pointer) {}

func Fn_gdk_drag_abort(param0 unsafe.Pointer, param1 uint32) {}

func Fn_gdk_drag_begin(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_gdk_drag_begin_for_device(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer) {
}

func Fn_gdk_drag_drop(param0 unsafe.Pointer, param1 uint32) {}

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

func Fn_gdk_offscreen_window_get_surface(param0 unsafe.Pointer) {}

func Fn_gdk_pango_context_get() {
	C.gdk_pango_context_get()
}

func Fn_gdk_pango_layout_get_clip_region(param0 unsafe.Pointer, param1 int, param2 int, param3 *int, param4 int) {
}

func Fn_gdk_pango_layout_line_get_clip_region(param0 unsafe.Pointer, param1 int, param2 int, param3 []int, param4 int) {
}

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

func Fn_gdk_selection_owner_set(param0 unsafe.Pointer, param1 Atom, param2 uint32, param3 bool) {}

func Fn_gdk_selection_property_get(param0 unsafe.Pointer, param1 **uint8, param2 unsafe.Pointer, param3 *int) {
}

func Fn_gdk_selection_send_notify(param0 unsafe.Pointer, param1 Atom, param2 Atom, param3 Atom, param4 uint32) {
}

func Fn_gdk_set_double_click_time(param0 uint) {}

func Fn_gdk_set_program_class(param0 string) {}

func Fn_gdk_set_show_events(param0 bool) {}

func Fn_gdk_setting_get(param0 string, param1 unsafe.Pointer) {}

func Fn_gdk_synthesize_window_state(param0 unsafe.Pointer, param1 int, param2 int) {}

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

func Fn_gdk_device_free_history(param0 []unsafe.Pointer, param1 int) {}

func Fn_gdk_device_grab_info_libgtk_only(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 *unsafe.Pointer, param3 *bool) {
}

func Fn_gdk_display_open_default_libgtk_only() {
	C.gdk_display_open_default_libgtk_only()
}

func Fn_gdk_keymap_get_default() {
	C.gdk_keymap_get_default()
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

func Fn_gdk_window_at_pointer(param0 *int, param1 *int) {}

func Fn_gdk_window_constrain_size(param0 unsafe.Pointer, param1 int, param2 int, param3 int, param4 *int, param5 *int) {
}

func Fn_gdk_window_process_all_updates() {
	C.gdk_window_process_all_updates()
}

func Fn_gdk_window_set_debug_updates(param0 bool) {}
