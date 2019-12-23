// Code generated - DO NOT EDIT.
// +build gdk_3.16

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

func Fn_add_option_entries_libgtk_only(group unsafe.Pointer) {}

func Fn_atom_intern(atomName string, onlyIfExists bool) {}

func Fn_atom_intern_static_string(atomName string) {}

func Fn_beep() {
	C.gdk_beep()
}

func Fn_cairo_create(window unsafe.Pointer) {}

func Fn_cairo_draw_from_gl(cr unsafe.Pointer, window unsafe.Pointer, source int, sourceType int, bufferScale int, x int, y int, width int, height int) {
}

func Fn_cairo_get_clip_rectangle(cr unsafe.Pointer) {}

func Fn_cairo_rectangle(cr unsafe.Pointer, rectangle unsafe.Pointer) {}

func Fn_cairo_region(cr unsafe.Pointer, region unsafe.Pointer) {}

func Fn_cairo_region_create_from_surface(surface unsafe.Pointer) {}

func Fn_cairo_set_source_color(cr unsafe.Pointer, color unsafe.Pointer) {}

func Fn_cairo_set_source_pixbuf(cr unsafe.Pointer, pixbuf unsafe.Pointer, pixbufX float64, pixbufY float64) {
}

func Fn_cairo_set_source_rgba(cr unsafe.Pointer, rgba unsafe.Pointer) {}

func Fn_cairo_set_source_window(cr unsafe.Pointer, window unsafe.Pointer, x float64, y float64) {}

func Fn_cairo_surface_create_from_pixbuf(pixbuf unsafe.Pointer, scale int, forWindow unsafe.Pointer) {
}

func Fn_color_parse(spec string) {}

func Fn_disable_multidevice() {
	C.gdk_disable_multidevice()
}

func Fn_drag_abort(context unsafe.Pointer, time uint32) {}

func Fn_drag_begin(window unsafe.Pointer, targets unsafe.Pointer) {}

func Fn_drag_begin_for_device(window unsafe.Pointer, device unsafe.Pointer, targets unsafe.Pointer) {
}

func Fn_drag_drop(context unsafe.Pointer, time uint32) {}

func Fn_drag_drop_succeeded(context unsafe.Pointer) {}

func Fn_drag_find_window_for_screen(context unsafe.Pointer, dragWindow unsafe.Pointer, screen unsafe.Pointer, xRoot int, yRoot int) {
}

func Fn_drag_get_selection(context unsafe.Pointer) {}

func Fn_drag_motion(context unsafe.Pointer, destWindow unsafe.Pointer, protocol int, xRoot int, yRoot int, suggestedAction int, possibleActions int, time uint32) {
}

func Fn_drag_status(context unsafe.Pointer, action int, time uint32) {}

func Fn_drop_finish(context unsafe.Pointer, success bool, time uint32) {}

func Fn_drop_reply(context unsafe.Pointer, accepted bool, time uint32) {}

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

func Fn_event_request_motions(event unsafe.Pointer) {}

func Fn_events_get_angle(event1 unsafe.Pointer, event2 unsafe.Pointer) {}

func Fn_events_get_center(event1 unsafe.Pointer, event2 unsafe.Pointer) {}

func Fn_events_get_distance(event1 unsafe.Pointer, event2 unsafe.Pointer) {}

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

func Fn_init(argc *int, argv *string) {}

func Fn_init_check(argc *int, argv *string) {}

func Fn_keyboard_grab(window unsafe.Pointer, ownerEvents bool, time uint32) {}

func Fn_keyboard_ungrab(time uint32) {}

func Fn_keyval_convert_case(symbol uint) {}

func Fn_keyval_from_name(keyvalName string) {}

func Fn_keyval_is_lower(keyval uint) {}

func Fn_keyval_is_upper(keyval uint) {}

func Fn_keyval_name(keyval uint) {}

func Fn_keyval_to_lower(keyval uint) {}

func Fn_keyval_to_unicode(keyval uint) {}

func Fn_keyval_to_upper(keyval uint) {}

func Fn_list_visuals() {
	C.gdk_list_visuals()
}

func Fn_notify_startup_complete() {
	C.gdk_notify_startup_complete()
}

func Fn_notify_startup_complete_with_id(startupId string) {}

func Fn_offscreen_window_get_embedder(window unsafe.Pointer) {}

func Fn_offscreen_window_get_surface(window unsafe.Pointer) {}

func Fn_offscreen_window_set_embedder(window unsafe.Pointer, embedder unsafe.Pointer) {}

func Fn_pango_context_get() {
	C.gdk_pango_context_get()
}

func Fn_pango_context_get_for_screen(screen unsafe.Pointer) {}

func Fn_pango_layout_get_clip_region(layout unsafe.Pointer, xOrigin int, yOrigin int, indexRanges *int, nRanges int) {
}

func Fn_pango_layout_line_get_clip_region(line unsafe.Pointer, xOrigin int, yOrigin int, indexRanges *int, nRanges int) {
}

func Fn_parse_args(argc *int, argv *string) {}

func Fn_pixbuf_get_from_surface(surface unsafe.Pointer, srcX int, srcY int, width int, height int) {}

func Fn_pixbuf_get_from_window(window unsafe.Pointer, srcX int, srcY int, width int, height int) {}

func Fn_pointer_grab(window unsafe.Pointer, ownerEvents bool, eventMask int, confineTo unsafe.Pointer, cursor unsafe.Pointer, time uint32) {
}

func Fn_pointer_is_grabbed() {
	C.gdk_pointer_is_grabbed()
}

func Fn_pointer_ungrab(time uint32) {}

func Fn_pre_parse_libgtk_only() {
	C.gdk_pre_parse_libgtk_only()
}

func Fn_property_change(window unsafe.Pointer, property Atom, type_ Atom, format int, mode int, data *uint8, nelements int) {
}

func Fn_property_delete(window unsafe.Pointer, property Atom) {}

func Fn_property_get(window unsafe.Pointer, property Atom, type_ Atom, offset uint64, length uint64, pdelete int) {
}

func Fn_query_depths() {}

func Fn_query_visual_types() {}

func Fn_selection_convert(requestor unsafe.Pointer, selection Atom, target Atom, time uint32) {}

func Fn_selection_owner_get(selection Atom) {}

func Fn_selection_owner_get_for_display(display unsafe.Pointer, selection Atom) {}

func Fn_selection_owner_set(owner unsafe.Pointer, selection Atom, time uint32, sendEvent bool) {}

func Fn_selection_owner_set_for_display(display unsafe.Pointer, owner unsafe.Pointer, selection Atom, time uint32, sendEvent bool) {
}

func Fn_selection_property_get(requestor unsafe.Pointer, data **uint8, propType unsafe.Pointer, propFormat *int) {
}

func Fn_selection_send_notify(requestor unsafe.Pointer, selection Atom, target Atom, property Atom, time uint32) {
}

func Fn_selection_send_notify_for_display(display unsafe.Pointer, requestor unsafe.Pointer, selection Atom, target Atom, property Atom, time uint32) {
}

func Fn_set_allowed_backends(backends string) {}

func Fn_set_double_click_time(msec uint) {}

func Fn_set_program_class(programClass string) {}

func Fn_set_show_events(showEvents bool) {}

func Fn_setting_get(name string, value unsafe.Pointer) {}

func Fn_synthesize_window_state(window unsafe.Pointer, unsetFlags int, setFlags int) {}

func Fn_test_render_sync(window unsafe.Pointer) {}

func Fn_test_simulate_button(window unsafe.Pointer, x int, y int, button uint, modifiers int, buttonPressrelease int) {
}

func Fn_test_simulate_key(window unsafe.Pointer, x int, y int, keyval uint, modifiers int, keyPressrelease int) {
}

func Fn_text_property_to_utf8_list_for_display(display unsafe.Pointer, encoding Atom, format int, text *uint8, length int) {
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

func Fn_unicode_to_keyval(wc uint32) {}

func Fn_utf8_to_string_target(str string) {}
