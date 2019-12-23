// Code generated - DO NOT EDIT.
// +build gdk_3.8

package gdk

import "unsafe"

// #include <gdk/gdk.h>
import "C"

// records
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

func Fn_add_option_entries_libgtk_only(param0 unsafe.Pointer) {}

func Fn_atom_intern(param0 string, param1 bool) {}

func Fn_atom_intern_static_string(param0 string) {}

func Fn_beep() {
	C.gdk_beep()
}

func Fn_cairo_create(param0 unsafe.Pointer) {}

func Fn_cairo_get_clip_rectangle(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_cairo_rectangle(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_cairo_region(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_cairo_region_create_from_surface(param0 unsafe.Pointer) {}

func Fn_cairo_set_source_color(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_cairo_set_source_pixbuf(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 float64, param3 float64) {
}

func Fn_cairo_set_source_rgba(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_cairo_set_source_window(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 float64, param3 float64) {
}

func Fn_color_parse(param0 string, param1 unsafe.Pointer) {}

func Fn_disable_multidevice() {
	C.gdk_disable_multidevice()
}

func Fn_drag_abort(param0 unsafe.Pointer, param1 uint32) {}

func Fn_drag_begin(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_drag_begin_for_device(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer) {}

func Fn_drag_drop(param0 unsafe.Pointer, param1 uint32) {}

func Fn_drag_drop_succeeded(param0 unsafe.Pointer) {}

func Fn_drag_find_window_for_screen(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer, param3 int, param4 int, param5 *unsafe.Pointer, param6 int) {
}

func Fn_drag_get_selection(param0 unsafe.Pointer) {}

func Fn_drag_motion(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int, param3 int, param4 int, param5 int, param6 int, param7 uint32) {
}

func Fn_drag_status(param0 unsafe.Pointer, param1 int, param2 uint32) {}

func Fn_drop_finish(param0 unsafe.Pointer, param1 bool, param2 uint32) {}

func Fn_drop_reply(param0 unsafe.Pointer, param1 bool, param2 uint32) {}

func Fn_error_trap_pop() {
	C.gdk_error_trap_pop()
}

func Fn_error_trap_pop_ignored() {
	C.gdk_error_trap_pop_ignored()
}

func Fn_error_trap_push() {
	C.gdk_error_trap_push()
}

func Fn_event_get() {
	C.gdk_event_get()
}

// UNSUPPORTED : event_handler_set : has callback

func Fn_event_peek() {
	C.gdk_event_peek()
}

func Fn_event_request_motions(param0 unsafe.Pointer) {}

func Fn_events_get_angle(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 *float64) {}

func Fn_events_get_center(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 *float64, param3 *float64) {
}

func Fn_events_get_distance(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 *float64) {}

func Fn_events_pending() {
	C.gdk_events_pending()
}

func Fn_flush() {
	C.gdk_flush()
}

func Fn_get_default_root_window() {
	C.gdk_get_default_root_window()
}

func Fn_get_display() {
	C.gdk_get_display()
}

func Fn_get_display_arg_name() {
	C.gdk_get_display_arg_name()
}

func Fn_get_program_class() {
	C.gdk_get_program_class()
}

func Fn_get_show_events() {
	C.gdk_get_show_events()
}

func Fn_gl_error_quark() {
	C.gdk_gl_error_quark()
}

func Fn_init(param0 *int, param1 *[]string) {}

func Fn_init_check(param0 *int, param1 *[]string) {}

func Fn_keyboard_grab(param0 unsafe.Pointer, param1 bool, param2 uint32) {}

func Fn_keyboard_ungrab(param0 uint32) {}

func Fn_keyval_convert_case(param0 uint, param1 *uint, param2 *uint) {}

func Fn_keyval_from_name(param0 string) {}

func Fn_keyval_is_lower(param0 uint) {}

func Fn_keyval_is_upper(param0 uint) {}

func Fn_keyval_name(param0 uint) {}

func Fn_keyval_to_lower(param0 uint) {}

func Fn_keyval_to_unicode(param0 uint) {}

func Fn_keyval_to_upper(param0 uint) {}

func Fn_list_visuals() {
	C.gdk_list_visuals()
}

func Fn_notify_startup_complete() {
	C.gdk_notify_startup_complete()
}

func Fn_notify_startup_complete_with_id(param0 string) {}

func Fn_offscreen_window_get_embedder(param0 unsafe.Pointer) {}

func Fn_offscreen_window_get_surface(param0 unsafe.Pointer) {}

func Fn_offscreen_window_set_embedder(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_pango_context_get() {
	C.gdk_pango_context_get()
}

func Fn_pango_context_get_for_screen(param0 unsafe.Pointer) {}

func Fn_pango_layout_get_clip_region(param0 unsafe.Pointer, param1 int, param2 int, param3 *int, param4 int) {
}

func Fn_pango_layout_line_get_clip_region(param0 unsafe.Pointer, param1 int, param2 int, param3 []int, param4 int) {
}

func Fn_parse_args(param0 *int, param1 *[]string) {}

func Fn_pixbuf_get_from_surface(param0 unsafe.Pointer, param1 int, param2 int, param3 int, param4 int) {
}

func Fn_pixbuf_get_from_window(param0 unsafe.Pointer, param1 int, param2 int, param3 int, param4 int) {
}

func Fn_pointer_grab(param0 unsafe.Pointer, param1 bool, param2 int, param3 unsafe.Pointer, param4 unsafe.Pointer, param5 uint32) {
}

func Fn_pointer_is_grabbed() {
	C.gdk_pointer_is_grabbed()
}

func Fn_pointer_ungrab(param0 uint32) {}

func Fn_pre_parse_libgtk_only() {
	C.gdk_pre_parse_libgtk_only()
}

func Fn_property_change(param0 unsafe.Pointer, param1 Atom, param2 Atom, param3 int, param4 int, param5 *uint8, param6 int) {
}

func Fn_property_delete(param0 unsafe.Pointer, param1 Atom) {}

func Fn_property_get(param0 unsafe.Pointer, param1 Atom, param2 Atom, param3 uint64, param4 uint64, param5 int, param6 unsafe.Pointer, param7 *int, param8 *int, param9 []*uint8) {
}

func Fn_query_depths(param0 []*int, param1 *int) {}

func Fn_query_visual_types(param0 []int, param1 *int) {}

func Fn_selection_convert(param0 unsafe.Pointer, param1 Atom, param2 Atom, param3 uint32) {}

func Fn_selection_owner_get(param0 Atom) {}

func Fn_selection_owner_get_for_display(param0 unsafe.Pointer, param1 Atom) {}

func Fn_selection_owner_set(param0 unsafe.Pointer, param1 Atom, param2 uint32, param3 bool) {}

func Fn_selection_owner_set_for_display(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 Atom, param3 uint32, param4 bool) {
}

func Fn_selection_property_get(param0 unsafe.Pointer, param1 **uint8, param2 unsafe.Pointer, param3 *int) {
}

func Fn_selection_send_notify(param0 unsafe.Pointer, param1 Atom, param2 Atom, param3 Atom, param4 uint32) {
}

func Fn_selection_send_notify_for_display(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 Atom, param3 Atom, param4 Atom, param5 uint32) {
}

func Fn_set_double_click_time(param0 uint) {}

func Fn_set_program_class(param0 string) {}

func Fn_set_show_events(param0 bool) {}

func Fn_setting_get(param0 string, param1 unsafe.Pointer) {}

func Fn_synthesize_window_state(param0 unsafe.Pointer, param1 int, param2 int) {}

func Fn_test_render_sync(param0 unsafe.Pointer) {}

func Fn_test_simulate_button(param0 unsafe.Pointer, param1 int, param2 int, param3 uint, param4 int, param5 int) {
}

func Fn_test_simulate_key(param0 unsafe.Pointer, param1 int, param2 int, param3 uint, param4 int, param5 int) {
}

func Fn_text_property_to_utf8_list_for_display(param0 unsafe.Pointer, param1 Atom, param2 int, param3 []uint8, param4 int, param5 *[]string) {
}

// UNSUPPORTED : threads_add_idle : has callback

// UNSUPPORTED : threads_add_idle_full : has callback

// UNSUPPORTED : threads_add_timeout : has callback

// UNSUPPORTED : threads_add_timeout_full : has callback

// UNSUPPORTED : threads_add_timeout_seconds : has callback

// UNSUPPORTED : threads_add_timeout_seconds_full : has callback

func Fn_threads_enter() {
	C.gdk_threads_enter()
}

func Fn_threads_init() {
	C.gdk_threads_init()
}

func Fn_threads_leave() {
	C.gdk_threads_leave()
}

// UNSUPPORTED : threads_set_lock_functions : has callback

func Fn_unicode_to_keyval(param0 uint32) {}

func Fn_utf8_to_string_target(param0 string) {}
