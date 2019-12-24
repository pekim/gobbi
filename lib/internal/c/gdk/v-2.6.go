// Code generated - DO NOT EDIT.
// +build gdk_2.6

package gdk

import "unsafe"

// #include <gdk/gdk.h>
import "C"

func toCBool(b bool) C.gboolean {
	if b {
		return C.TRUE
	}
	return C.TRUE
}

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

func Fn_gdk_add_option_entries_libgtk_only(param0 unsafe.Pointer) {
	cValue0 := (*C.GOptionGroup)(unsafe.Pointer(param0))

	C.gdk_add_option_entries_libgtk_only(cValue0)
}

func Fn_gdk_atom_intern(param0 string, param1 bool) {
	cValue0 := 42
	cValue1 := toCBool(param1)

	C.gdk_atom_intern(cValue0, cValue1)
}

func Fn_gdk_beep() {

	C.gdk_beep()
}

func Fn_gdk_cairo_get_clip_rectangle(param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValue0 := (*C.cairo_t)(unsafe.Pointer(param0))
	cValue1 := (*C.GdkRectangle)(unsafe.Pointer(param1))

	C.gdk_cairo_get_clip_rectangle(cValue0, cValue1)
}

func Fn_gdk_cairo_region_create_from_surface(param0 unsafe.Pointer) {
	cValue0 := (*C.cairo_surface_t)(unsafe.Pointer(param0))

	C.gdk_cairo_region_create_from_surface(cValue0)
}

func Fn_gdk_color_parse(param0 string, param1 unsafe.Pointer) {
	cValue0 := 42
	cValue1 := (*C.GdkColor)(unsafe.Pointer(param1))

	C.gdk_color_parse(cValue0, cValue1)
}

func Fn_gdk_drag_abort(param0 unsafe.Pointer, param1 uint32) {
	cValue0 := (*C.GdkDragContext)(unsafe.Pointer(param0))
	cValue1 := (C.guint32)(param1)

	C.gdk_drag_abort(cValue0, cValue1)
}

func Fn_gdk_drag_begin(param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValue0 := (*C.GdkWindow)(unsafe.Pointer(param0))
	cValue1 := (*C.GList)(unsafe.Pointer(param1))

	C.gdk_drag_begin(cValue0, cValue1)
}

func Fn_gdk_drag_begin_for_device(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer) {
	cValue0 := (*C.GdkWindow)(unsafe.Pointer(param0))
	cValue1 := (*C.GdkDevice)(unsafe.Pointer(param1))
	cValue2 := (*C.GList)(unsafe.Pointer(param2))

	C.gdk_drag_begin_for_device(cValue0, cValue1, cValue2)
}

func Fn_gdk_drag_drop(param0 unsafe.Pointer, param1 uint32) {
	cValue0 := (*C.GdkDragContext)(unsafe.Pointer(param0))
	cValue1 := (C.guint32)(param1)

	C.gdk_drag_drop(cValue0, cValue1)
}

func Fn_gdk_drag_drop_succeeded(param0 unsafe.Pointer) {
	cValue0 := (*C.GdkDragContext)(unsafe.Pointer(param0))

	C.gdk_drag_drop_succeeded(cValue0)
}

func Fn_gdk_drag_find_window_for_screen(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer, param3 int, param4 int, param5 *unsafe.Pointer, param6 *int) {
	cValue0 := (*C.GdkDragContext)(unsafe.Pointer(param0))
	cValue1 := (*C.GdkWindow)(unsafe.Pointer(param1))
	cValue2 := (*C.GdkScreen)(unsafe.Pointer(param2))
	cValue3 := (C.gint)(param3)
	cValue4 := (C.gint)(param4)
	cValue5 := (**C.GdkWindow)(unsafe.Pointer(param5))
	cValue6 := (*C.GdkDragProtocol)(unsafe.Pointer(param6))

	C.gdk_drag_find_window_for_screen(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5, cValue6)
}

func Fn_gdk_drag_get_selection(param0 unsafe.Pointer) {
	cValue0 := (*C.GdkDragContext)(unsafe.Pointer(param0))

	C.gdk_drag_get_selection(cValue0)
}

func Fn_gdk_drag_motion(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int, param3 int, param4 int, param5 int, param6 int, param7 uint32) {
	cValue0 := (*C.GdkDragContext)(unsafe.Pointer(param0))
	cValue1 := (*C.GdkWindow)(unsafe.Pointer(param1))
	cValue2 := (C.GdkDragProtocol)(param2)
	cValue3 := (C.gint)(param3)
	cValue4 := (C.gint)(param4)
	cValue5 := (C.GdkDragAction)(param5)
	cValue6 := (C.GdkDragAction)(param6)
	cValue7 := (C.guint32)(param7)

	C.gdk_drag_motion(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5, cValue6, cValue7)
}

func Fn_gdk_drag_status(param0 unsafe.Pointer, param1 int, param2 uint32) {
	cValue0 := (*C.GdkDragContext)(unsafe.Pointer(param0))
	cValue1 := (C.GdkDragAction)(param1)
	cValue2 := (C.guint32)(param2)

	C.gdk_drag_status(cValue0, cValue1, cValue2)
}

func Fn_gdk_drop_finish(param0 unsafe.Pointer, param1 bool, param2 uint32) {
	cValue0 := (*C.GdkDragContext)(unsafe.Pointer(param0))
	cValue1 := toCBool(param1)
	cValue2 := (C.guint32)(param2)

	C.gdk_drop_finish(cValue0, cValue1, cValue2)
}

func Fn_gdk_drop_reply(param0 unsafe.Pointer, param1 bool, param2 uint32) {
	cValue0 := (*C.GdkDragContext)(unsafe.Pointer(param0))
	cValue1 := toCBool(param1)
	cValue2 := (C.guint32)(param2)

	C.gdk_drop_reply(cValue0, cValue1, cValue2)
}

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

func Fn_gdk_init(param0 *int, param1 *[]string) {
	// has array param
}

func Fn_gdk_init_check(param0 *int, param1 *[]string) {
	// has array param
}

func Fn_gdk_keyboard_grab(param0 unsafe.Pointer, param1 bool, param2 uint32) {
	cValue0 := (*C.GdkWindow)(unsafe.Pointer(param0))
	cValue1 := toCBool(param1)
	cValue2 := (C.guint32)(param2)

	C.gdk_keyboard_grab(cValue0, cValue1, cValue2)
}

func Fn_gdk_keyboard_ungrab(param0 uint32) {
	cValue0 := (C.guint32)(param0)

	C.gdk_keyboard_ungrab(cValue0)
}

func Fn_gdk_keyval_convert_case(param0 uint, param1 *uint, param2 *uint) {
	cValue0 := (C.guint)(param0)
	cValue1 := (*C.guint)(unsafe.Pointer(param1))
	cValue2 := (*C.guint)(unsafe.Pointer(param2))

	C.gdk_keyval_convert_case(cValue0, cValue1, cValue2)
}

func Fn_gdk_keyval_from_name(param0 string) {
	cValue0 := 42

	C.gdk_keyval_from_name(cValue0)
}

func Fn_gdk_keyval_is_lower(param0 uint) {
	cValue0 := (C.guint)(param0)

	C.gdk_keyval_is_lower(cValue0)
}

func Fn_gdk_keyval_is_upper(param0 uint) {
	cValue0 := (C.guint)(param0)

	C.gdk_keyval_is_upper(cValue0)
}

func Fn_gdk_keyval_name(param0 uint) {
	cValue0 := (C.guint)(param0)

	C.gdk_keyval_name(cValue0)
}

func Fn_gdk_keyval_to_lower(param0 uint) {
	cValue0 := (C.guint)(param0)

	C.gdk_keyval_to_lower(cValue0)
}

func Fn_gdk_keyval_to_unicode(param0 uint) {
	cValue0 := (C.guint)(param0)

	C.gdk_keyval_to_unicode(cValue0)
}

func Fn_gdk_keyval_to_upper(param0 uint) {
	cValue0 := (C.guint)(param0)

	C.gdk_keyval_to_upper(cValue0)
}

func Fn_gdk_list_visuals() {

	C.gdk_list_visuals()
}

func Fn_gdk_notify_startup_complete() {

	C.gdk_notify_startup_complete()
}

func Fn_gdk_offscreen_window_get_surface(param0 unsafe.Pointer) {
	cValue0 := (*C.GdkWindow)(unsafe.Pointer(param0))

	C.gdk_offscreen_window_get_surface(cValue0)
}

func Fn_gdk_pango_context_get() {

	C.gdk_pango_context_get()
}

func Fn_gdk_pango_context_get_for_screen(param0 unsafe.Pointer) {
	cValue0 := (*C.GdkScreen)(unsafe.Pointer(param0))

	C.gdk_pango_context_get_for_screen(cValue0)
}

func Fn_gdk_pango_layout_get_clip_region(param0 unsafe.Pointer, param1 int, param2 int, param3 *int, param4 int) {
	cValue0 := (*C.PangoLayout)(unsafe.Pointer(param0))
	cValue1 := (C.gint)(param1)
	cValue2 := (C.gint)(param2)
	cValue3 := (*C.gint)(unsafe.Pointer(param3))
	cValue4 := (C.gint)(param4)

	C.gdk_pango_layout_get_clip_region(cValue0, cValue1, cValue2, cValue3, cValue4)
}

func Fn_gdk_pango_layout_line_get_clip_region(param0 unsafe.Pointer, param1 int, param2 int, param3 []int, param4 int) {
	// has array param
}

func Fn_gdk_parse_args(param0 *int, param1 *[]string) {
	// has array param
}

func Fn_gdk_pixbuf_get_from_surface(param0 unsafe.Pointer, param1 int, param2 int, param3 int, param4 int) {
	cValue0 := (*C.cairo_surface_t)(unsafe.Pointer(param0))
	cValue1 := (C.gint)(param1)
	cValue2 := (C.gint)(param2)
	cValue3 := (C.gint)(param3)
	cValue4 := (C.gint)(param4)

	C.gdk_pixbuf_get_from_surface(cValue0, cValue1, cValue2, cValue3, cValue4)
}

func Fn_gdk_pixbuf_get_from_window(param0 unsafe.Pointer, param1 int, param2 int, param3 int, param4 int) {
	cValue0 := (*C.GdkWindow)(unsafe.Pointer(param0))
	cValue1 := (C.gint)(param1)
	cValue2 := (C.gint)(param2)
	cValue3 := (C.gint)(param3)
	cValue4 := (C.gint)(param4)

	C.gdk_pixbuf_get_from_window(cValue0, cValue1, cValue2, cValue3, cValue4)
}

func Fn_gdk_pointer_grab(param0 unsafe.Pointer, param1 bool, param2 int, param3 unsafe.Pointer, param4 unsafe.Pointer, param5 uint32) {
	cValue0 := (*C.GdkWindow)(unsafe.Pointer(param0))
	cValue1 := toCBool(param1)
	cValue2 := (C.GdkEventMask)(param2)
	cValue3 := (*C.GdkWindow)(unsafe.Pointer(param3))
	cValue4 := (*C.GdkCursor)(unsafe.Pointer(param4))
	cValue5 := (C.guint32)(param5)

	C.gdk_pointer_grab(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5)
}

func Fn_gdk_pointer_is_grabbed() {

	C.gdk_pointer_is_grabbed()
}

func Fn_gdk_pointer_ungrab(param0 uint32) {
	cValue0 := (C.guint32)(param0)

	C.gdk_pointer_ungrab(cValue0)
}

func Fn_gdk_pre_parse_libgtk_only() {

	C.gdk_pre_parse_libgtk_only()
}

func Fn_gdk_property_change(param0 unsafe.Pointer, param1 Atom, param2 Atom, param3 int, param4 int, param5 *uint8, param6 int) {
	cValue0 := (*C.GdkWindow)(unsafe.Pointer(param0))
	cValue1 := (C.GdkAtom)(param1)
	cValue2 := (C.GdkAtom)(param2)
	cValue3 := (C.gint)(param3)
	cValue4 := (C.GdkPropMode)(param4)
	cValue5 := (*C.guchar)(unsafe.Pointer(param5))
	cValue6 := (C.gint)(param6)

	C.gdk_property_change(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5, cValue6)
}

func Fn_gdk_property_delete(param0 unsafe.Pointer, param1 Atom) {
	cValue0 := (*C.GdkWindow)(unsafe.Pointer(param0))
	cValue1 := (C.GdkAtom)(param1)

	C.gdk_property_delete(cValue0, cValue1)
}

func Fn_gdk_property_get(param0 unsafe.Pointer, param1 Atom, param2 Atom, param3 uint64, param4 uint64, param5 int, param6 unsafe.Pointer, param7 *int, param8 *int, param9 []*uint8) {
	// has array param
}

func Fn_gdk_query_depths(param0 []*int, param1 *int) {
	// has array param
}

func Fn_gdk_query_visual_types(param0 []*int, param1 *int) {
	// has array param
}

func Fn_gdk_selection_convert(param0 unsafe.Pointer, param1 Atom, param2 Atom, param3 uint32) {
	cValue0 := (*C.GdkWindow)(unsafe.Pointer(param0))
	cValue1 := (C.GdkAtom)(param1)
	cValue2 := (C.GdkAtom)(param2)
	cValue3 := (C.guint32)(param3)

	C.gdk_selection_convert(cValue0, cValue1, cValue2, cValue3)
}

func Fn_gdk_selection_owner_get(param0 Atom) {
	cValue0 := (C.GdkAtom)(param0)

	C.gdk_selection_owner_get(cValue0)
}

func Fn_gdk_selection_owner_get_for_display(param0 unsafe.Pointer, param1 Atom) {
	cValue0 := (*C.GdkDisplay)(unsafe.Pointer(param0))
	cValue1 := (C.GdkAtom)(param1)

	C.gdk_selection_owner_get_for_display(cValue0, cValue1)
}

func Fn_gdk_selection_owner_set(param0 unsafe.Pointer, param1 Atom, param2 uint32, param3 bool) {
	cValue0 := (*C.GdkWindow)(unsafe.Pointer(param0))
	cValue1 := (C.GdkAtom)(param1)
	cValue2 := (C.guint32)(param2)
	cValue3 := toCBool(param3)

	C.gdk_selection_owner_set(cValue0, cValue1, cValue2, cValue3)
}

func Fn_gdk_selection_owner_set_for_display(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 Atom, param3 uint32, param4 bool) {
	cValue0 := (*C.GdkDisplay)(unsafe.Pointer(param0))
	cValue1 := (*C.GdkWindow)(unsafe.Pointer(param1))
	cValue2 := (C.GdkAtom)(param2)
	cValue3 := (C.guint32)(param3)
	cValue4 := toCBool(param4)

	C.gdk_selection_owner_set_for_display(cValue0, cValue1, cValue2, cValue3, cValue4)
}

func Fn_gdk_selection_property_get(param0 unsafe.Pointer, param1 **uint8, param2 unsafe.Pointer, param3 *int) {
	cValue0 := (*C.GdkWindow)(unsafe.Pointer(param0))
	cValue1 := (**C.guchar)(unsafe.Pointer(param1))
	cValue2 := (*C.GdkAtom)(unsafe.Pointer(param2))
	cValue3 := (*C.gint)(unsafe.Pointer(param3))

	C.gdk_selection_property_get(cValue0, cValue1, cValue2, cValue3)
}

func Fn_gdk_selection_send_notify(param0 unsafe.Pointer, param1 Atom, param2 Atom, param3 Atom, param4 uint32) {
	cValue0 := (*C.GdkWindow)(unsafe.Pointer(param0))
	cValue1 := (C.GdkAtom)(param1)
	cValue2 := (C.GdkAtom)(param2)
	cValue3 := (C.GdkAtom)(param3)
	cValue4 := (C.guint32)(param4)

	C.gdk_selection_send_notify(cValue0, cValue1, cValue2, cValue3, cValue4)
}

func Fn_gdk_selection_send_notify_for_display(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 Atom, param3 Atom, param4 Atom, param5 uint32) {
	cValue0 := (*C.GdkDisplay)(unsafe.Pointer(param0))
	cValue1 := (*C.GdkWindow)(unsafe.Pointer(param1))
	cValue2 := (C.GdkAtom)(param2)
	cValue3 := (C.GdkAtom)(param3)
	cValue4 := (C.GdkAtom)(param4)
	cValue5 := (C.guint32)(param5)

	C.gdk_selection_send_notify_for_display(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5)
}

func Fn_gdk_set_double_click_time(param0 uint) {
	cValue0 := (C.guint)(param0)

	C.gdk_set_double_click_time(cValue0)
}

func Fn_gdk_set_program_class(param0 string) {
	cValue0 := 42

	C.gdk_set_program_class(cValue0)
}

func Fn_gdk_set_show_events(param0 bool) {
	cValue0 := toCBool(param0)

	C.gdk_set_show_events(cValue0)
}

func Fn_gdk_setting_get(param0 string, param1 unsafe.Pointer) {
	cValue0 := 42
	cValue1 := (*C.GValue)(unsafe.Pointer(param1))

	C.gdk_setting_get(cValue0, cValue1)
}

// UNSUPPORTED : synthesize_window_state : blacklisted
func Fn_gdk_text_property_to_utf8_list_for_display(param0 unsafe.Pointer, param1 Atom, param2 int, param3 []uint8, param4 int, param5 *[]string) {
	// has array param
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

func Fn_gdk_unicode_to_keyval(param0 uint32) {
	cValue0 := (C.guint32)(param0)

	C.gdk_unicode_to_keyval(cValue0)
}

func Fn_gdk_utf8_to_string_target(param0 string) {
	cValue0 := 42

	C.gdk_utf8_to_string_target(cValue0)
}

func Fn_gdk_cursor_new(param0 int) {
	cValue0 := (C.GdkCursorType)(param0)

	C.gdk_cursor_new(cValue0)
}

func Fn_gdk_cursor_new_for_display(param0 unsafe.Pointer, param1 int) {
	cValue0 := (*C.GdkDisplay)(unsafe.Pointer(param0))
	cValue1 := (C.GdkCursorType)(param1)

	C.gdk_cursor_new_for_display(cValue0, cValue1)
}

func Fn_gdk_cursor_new_from_pixbuf(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int, param3 int) {
	cValue0 := (*C.GdkDisplay)(unsafe.Pointer(param0))
	cValue1 := (*C.GdkPixbuf)(unsafe.Pointer(param1))
	cValue2 := (C.gint)(param2)
	cValue3 := (C.gint)(param3)

	C.gdk_cursor_new_from_pixbuf(cValue0, cValue1, cValue2, cValue3)
}

func Fn_gdk_cursor_get_display(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkCursor)(unsafe.Pointer(paramInstance))

	C.gdk_cursor_get_display(cValueInstance)
}

func Fn_gdk_cursor_ref(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkCursor)(unsafe.Pointer(paramInstance))

	C.gdk_cursor_ref(cValueInstance)
}

func Fn_gdk_cursor_unref(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkCursor)(unsafe.Pointer(paramInstance))

	C.gdk_cursor_unref(cValueInstance)
}

func Fn_gdk_device_get_axis(paramInstance unsafe.Pointer, param0 []float64, param1 int, param2 *float64) {
	// has array param
}

func Fn_gdk_device_get_history(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 uint32, param2 uint32, param3 []*unsafe.Pointer, param4 *int) {
	// has array param
}

func Fn_gdk_device_get_state(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 []float64, param2 *int) {
	// has array param
}

func Fn_gdk_device_list_slave_devices(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkDevice)(unsafe.Pointer(paramInstance))

	C.gdk_device_list_slave_devices(cValueInstance)
}

func Fn_gdk_device_set_axis_use(paramInstance unsafe.Pointer, param0 uint, param1 int) {
	cValueInstance := (*C.GdkDevice)(unsafe.Pointer(paramInstance))
	cValue0 := (C.guint)(param0)
	cValue1 := (C.GdkAxisUse)(param1)

	C.gdk_device_set_axis_use(cValueInstance, cValue0, cValue1)
}

func Fn_gdk_device_set_key(paramInstance unsafe.Pointer, param0 uint, param1 uint, param2 int) {
	cValueInstance := (*C.GdkDevice)(unsafe.Pointer(paramInstance))
	cValue0 := (C.guint)(param0)
	cValue1 := (C.guint)(param1)
	cValue2 := (C.GdkModifierType)(param2)

	C.gdk_device_set_key(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gdk_device_set_mode(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GdkDevice)(unsafe.Pointer(paramInstance))
	cValue0 := (C.GdkInputMode)(param0)

	C.gdk_device_set_mode(cValueInstance, cValue0)
}

func Fn_gdk_device_free_history(param0 []unsafe.Pointer, param1 int) {
	// has array param
}

func Fn_gdk_device_grab_info_libgtk_only(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 *unsafe.Pointer, param3 *bool) {
	cValue0 := (*C.GdkDisplay)(unsafe.Pointer(param0))
	cValue1 := (*C.GdkDevice)(unsafe.Pointer(param1))
	cValue2 := (**C.GdkWindow)(unsafe.Pointer(param2))
	cValue3 := (*C.gboolean)(unsafe.Pointer(param3))

	C.gdk_device_grab_info_libgtk_only(cValue0, cValue1, cValue2, cValue3)
}

func Fn_gdk_display_beep(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkDisplay)(unsafe.Pointer(paramInstance))

	C.gdk_display_beep(cValueInstance)
}

func Fn_gdk_display_close(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkDisplay)(unsafe.Pointer(paramInstance))

	C.gdk_display_close(cValueInstance)
}

func Fn_gdk_display_device_is_grabbed(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GdkDisplay)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GdkDevice)(unsafe.Pointer(param0))

	C.gdk_display_device_is_grabbed(cValueInstance, cValue0)
}

func Fn_gdk_display_flush(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkDisplay)(unsafe.Pointer(paramInstance))

	C.gdk_display_flush(cValueInstance)
}

func Fn_gdk_display_get_default_cursor_size(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkDisplay)(unsafe.Pointer(paramInstance))

	C.gdk_display_get_default_cursor_size(cValueInstance)
}

func Fn_gdk_display_get_default_group(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkDisplay)(unsafe.Pointer(paramInstance))

	C.gdk_display_get_default_group(cValueInstance)
}

func Fn_gdk_display_get_default_screen(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkDisplay)(unsafe.Pointer(paramInstance))

	C.gdk_display_get_default_screen(cValueInstance)
}

func Fn_gdk_display_get_event(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkDisplay)(unsafe.Pointer(paramInstance))

	C.gdk_display_get_event(cValueInstance)
}

func Fn_gdk_display_get_maximal_cursor_size(paramInstance unsafe.Pointer, param0 *uint, param1 *uint) {
	cValueInstance := (*C.GdkDisplay)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.guint)(unsafe.Pointer(param0))
	cValue1 := (*C.guint)(unsafe.Pointer(param1))

	C.gdk_display_get_maximal_cursor_size(cValueInstance, cValue0, cValue1)
}

func Fn_gdk_display_get_n_screens(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkDisplay)(unsafe.Pointer(paramInstance))

	C.gdk_display_get_n_screens(cValueInstance)
}

func Fn_gdk_display_get_name(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkDisplay)(unsafe.Pointer(paramInstance))

	C.gdk_display_get_name(cValueInstance)
}

func Fn_gdk_display_get_pointer(paramInstance unsafe.Pointer, param0 *unsafe.Pointer, param1 *int, param2 *int, param3 *int) {
	cValueInstance := (*C.GdkDisplay)(unsafe.Pointer(paramInstance))
	cValue0 := (**C.GdkScreen)(unsafe.Pointer(param0))
	cValue1 := (*C.gint)(unsafe.Pointer(param1))
	cValue2 := (*C.gint)(unsafe.Pointer(param2))
	cValue3 := (*C.GdkModifierType)(unsafe.Pointer(param3))

	C.gdk_display_get_pointer(cValueInstance, cValue0, cValue1, cValue2, cValue3)
}

func Fn_gdk_display_get_screen(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GdkDisplay)(unsafe.Pointer(paramInstance))
	cValue0 := (C.gint)(param0)

	C.gdk_display_get_screen(cValueInstance, cValue0)
}

func Fn_gdk_display_get_window_at_pointer(paramInstance unsafe.Pointer, param0 *int, param1 *int) {
	cValueInstance := (*C.GdkDisplay)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.gint)(unsafe.Pointer(param0))
	cValue1 := (*C.gint)(unsafe.Pointer(param1))

	C.gdk_display_get_window_at_pointer(cValueInstance, cValue0, cValue1)
}

func Fn_gdk_display_keyboard_ungrab(paramInstance unsafe.Pointer, param0 uint32) {
	cValueInstance := (*C.GdkDisplay)(unsafe.Pointer(paramInstance))
	cValue0 := (C.guint32)(param0)

	C.gdk_display_keyboard_ungrab(cValueInstance, cValue0)
}

func Fn_gdk_display_list_devices(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkDisplay)(unsafe.Pointer(paramInstance))

	C.gdk_display_list_devices(cValueInstance)
}

func Fn_gdk_display_peek_event(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkDisplay)(unsafe.Pointer(paramInstance))

	C.gdk_display_peek_event(cValueInstance)
}

func Fn_gdk_display_pointer_is_grabbed(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkDisplay)(unsafe.Pointer(paramInstance))

	C.gdk_display_pointer_is_grabbed(cValueInstance)
}

func Fn_gdk_display_pointer_ungrab(paramInstance unsafe.Pointer, param0 uint32) {
	cValueInstance := (*C.GdkDisplay)(unsafe.Pointer(paramInstance))
	cValue0 := (C.guint32)(param0)

	C.gdk_display_pointer_ungrab(cValueInstance, cValue0)
}

func Fn_gdk_display_put_event(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GdkDisplay)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GdkEvent)(unsafe.Pointer(param0))

	C.gdk_display_put_event(cValueInstance, cValue0)
}

func Fn_gdk_display_request_selection_notification(paramInstance unsafe.Pointer, param0 Atom) {
	cValueInstance := (*C.GdkDisplay)(unsafe.Pointer(paramInstance))
	cValue0 := (C.GdkAtom)(param0)

	C.gdk_display_request_selection_notification(cValueInstance, cValue0)
}

func Fn_gdk_display_set_double_click_distance(paramInstance unsafe.Pointer, param0 uint) {
	cValueInstance := (*C.GdkDisplay)(unsafe.Pointer(paramInstance))
	cValue0 := (C.guint)(param0)

	C.gdk_display_set_double_click_distance(cValueInstance, cValue0)
}

func Fn_gdk_display_set_double_click_time(paramInstance unsafe.Pointer, param0 uint) {
	cValueInstance := (*C.GdkDisplay)(unsafe.Pointer(paramInstance))
	cValue0 := (C.guint)(param0)

	C.gdk_display_set_double_click_time(cValueInstance, cValue0)
}

func Fn_gdk_display_store_clipboard(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 uint32, param2 []Atom, param3 int) {
	// has array param
}

func Fn_gdk_display_supports_clipboard_persistence(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkDisplay)(unsafe.Pointer(paramInstance))

	C.gdk_display_supports_clipboard_persistence(cValueInstance)
}

func Fn_gdk_display_supports_cursor_alpha(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkDisplay)(unsafe.Pointer(paramInstance))

	C.gdk_display_supports_cursor_alpha(cValueInstance)
}

func Fn_gdk_display_supports_cursor_color(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkDisplay)(unsafe.Pointer(paramInstance))

	C.gdk_display_supports_cursor_color(cValueInstance)
}

func Fn_gdk_display_supports_selection_notification(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkDisplay)(unsafe.Pointer(paramInstance))

	C.gdk_display_supports_selection_notification(cValueInstance)
}

func Fn_gdk_display_sync(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkDisplay)(unsafe.Pointer(paramInstance))

	C.gdk_display_sync(cValueInstance)
}

func Fn_gdk_display_get_default() {

	C.gdk_display_get_default()
}

func Fn_gdk_display_open(param0 string) {
	cValue0 := 42

	C.gdk_display_open(cValue0)
}

func Fn_gdk_display_open_default_libgtk_only() {

	C.gdk_display_open_default_libgtk_only()
}

func Fn_gdk_display_manager_get_default_display(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkDisplayManager)(unsafe.Pointer(paramInstance))

	C.gdk_display_manager_get_default_display(cValueInstance)
}

func Fn_gdk_display_manager_list_displays(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkDisplayManager)(unsafe.Pointer(paramInstance))

	C.gdk_display_manager_list_displays(cValueInstance)
}

func Fn_gdk_display_manager_set_default_display(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GdkDisplayManager)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GdkDisplay)(unsafe.Pointer(param0))

	C.gdk_display_manager_set_default_display(cValueInstance, cValue0)
}

func Fn_gdk_display_manager_get() {

	C.gdk_display_manager_get()
}

func Fn_gdk_drag_context_get_device(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkDragContext)(unsafe.Pointer(paramInstance))

	C.gdk_drag_context_get_device(cValueInstance)
}

func Fn_gdk_drag_context_set_device(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GdkDragContext)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GdkDevice)(unsafe.Pointer(param0))

	C.gdk_drag_context_set_device(cValueInstance, cValue0)
}

func Fn_gdk_keymap_get_direction(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkKeymap)(unsafe.Pointer(paramInstance))

	C.gdk_keymap_get_direction(cValueInstance)
}

func Fn_gdk_keymap_get_entries_for_keycode(paramInstance unsafe.Pointer, param0 uint, param1 []unsafe.Pointer, param2 []*uint, param3 *int) {
	// has array param
}

func Fn_gdk_keymap_get_entries_for_keyval(paramInstance unsafe.Pointer, param0 uint, param1 []unsafe.Pointer, param2 *int) {
	// has array param
}

func Fn_gdk_keymap_lookup_key(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GdkKeymap)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GdkKeymapKey)(unsafe.Pointer(param0))

	C.gdk_keymap_lookup_key(cValueInstance, cValue0)
}

func Fn_gdk_keymap_translate_keyboard_state(paramInstance unsafe.Pointer, param0 uint, param1 int, param2 int, param3 *uint, param4 *int, param5 *int, param6 *int) {
	cValueInstance := (*C.GdkKeymap)(unsafe.Pointer(paramInstance))
	cValue0 := (C.guint)(param0)
	cValue1 := (C.GdkModifierType)(param1)
	cValue2 := (C.gint)(param2)
	cValue3 := (*C.guint)(unsafe.Pointer(param3))
	cValue4 := (*C.gint)(unsafe.Pointer(param4))
	cValue5 := (*C.gint)(unsafe.Pointer(param5))
	cValue6 := (*C.GdkModifierType)(unsafe.Pointer(param6))

	C.gdk_keymap_translate_keyboard_state(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4, cValue5, cValue6)
}

func Fn_gdk_keymap_get_default() {

	C.gdk_keymap_get_default()
}

func Fn_gdk_keymap_get_for_display(param0 unsafe.Pointer) {
	cValue0 := (*C.GdkDisplay)(unsafe.Pointer(param0))

	C.gdk_keymap_get_for_display(cValue0)
}

func Fn_gdk_monitor_get_manufacturer(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkMonitor)(unsafe.Pointer(paramInstance))

	C.gdk_monitor_get_manufacturer(cValueInstance)
}

func Fn_gdk_monitor_get_model(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkMonitor)(unsafe.Pointer(paramInstance))

	C.gdk_monitor_get_model(cValueInstance)
}

func Fn_gdk_screen_get_display(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkScreen)(unsafe.Pointer(paramInstance))

	C.gdk_screen_get_display(cValueInstance)
}

func Fn_gdk_screen_get_height(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkScreen)(unsafe.Pointer(paramInstance))

	C.gdk_screen_get_height(cValueInstance)
}

func Fn_gdk_screen_get_height_mm(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkScreen)(unsafe.Pointer(paramInstance))

	C.gdk_screen_get_height_mm(cValueInstance)
}

func Fn_gdk_screen_get_monitor_at_point(paramInstance unsafe.Pointer, param0 int, param1 int) {
	cValueInstance := (*C.GdkScreen)(unsafe.Pointer(paramInstance))
	cValue0 := (C.gint)(param0)
	cValue1 := (C.gint)(param1)

	C.gdk_screen_get_monitor_at_point(cValueInstance, cValue0, cValue1)
}

func Fn_gdk_screen_get_monitor_at_window(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GdkScreen)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GdkWindow)(unsafe.Pointer(param0))

	C.gdk_screen_get_monitor_at_window(cValueInstance, cValue0)
}

func Fn_gdk_screen_get_monitor_geometry(paramInstance unsafe.Pointer, param0 int, param1 unsafe.Pointer) {
	cValueInstance := (*C.GdkScreen)(unsafe.Pointer(paramInstance))
	cValue0 := (C.gint)(param0)
	cValue1 := (*C.GdkRectangle)(unsafe.Pointer(param1))

	C.gdk_screen_get_monitor_geometry(cValueInstance, cValue0, cValue1)
}

func Fn_gdk_screen_get_n_monitors(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkScreen)(unsafe.Pointer(paramInstance))

	C.gdk_screen_get_n_monitors(cValueInstance)
}

func Fn_gdk_screen_get_number(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkScreen)(unsafe.Pointer(paramInstance))

	C.gdk_screen_get_number(cValueInstance)
}

func Fn_gdk_screen_get_root_window(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkScreen)(unsafe.Pointer(paramInstance))

	C.gdk_screen_get_root_window(cValueInstance)
}

func Fn_gdk_screen_get_setting(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer) {
	cValueInstance := (*C.GdkScreen)(unsafe.Pointer(paramInstance))
	cValue0 := 42
	cValue1 := (*C.GValue)(unsafe.Pointer(param1))

	C.gdk_screen_get_setting(cValueInstance, cValue0, cValue1)
}

func Fn_gdk_screen_get_system_visual(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkScreen)(unsafe.Pointer(paramInstance))

	C.gdk_screen_get_system_visual(cValueInstance)
}

func Fn_gdk_screen_get_toplevel_windows(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkScreen)(unsafe.Pointer(paramInstance))

	C.gdk_screen_get_toplevel_windows(cValueInstance)
}

func Fn_gdk_screen_get_width(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkScreen)(unsafe.Pointer(paramInstance))

	C.gdk_screen_get_width(cValueInstance)
}

func Fn_gdk_screen_get_width_mm(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkScreen)(unsafe.Pointer(paramInstance))

	C.gdk_screen_get_width_mm(cValueInstance)
}

func Fn_gdk_screen_list_visuals(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkScreen)(unsafe.Pointer(paramInstance))

	C.gdk_screen_list_visuals(cValueInstance)
}

func Fn_gdk_screen_make_display_name(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkScreen)(unsafe.Pointer(paramInstance))

	C.gdk_screen_make_display_name(cValueInstance)
}

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

func Fn_gdk_seat_get_display(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkSeat)(unsafe.Pointer(paramInstance))

	C.gdk_seat_get_display(cValueInstance)
}

// UNSUPPORTED : grab : has callback

func Fn_gdk_visual_get_screen(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkVisual)(unsafe.Pointer(paramInstance))

	C.gdk_visual_get_screen(cValueInstance)
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

func Fn_gdk_visual_get_best_with_both(param0 int, param1 int) {
	cValue0 := (C.gint)(param0)
	cValue1 := (C.GdkVisualType)(param1)

	C.gdk_visual_get_best_with_both(cValue0, cValue1)
}

func Fn_gdk_visual_get_best_with_depth(param0 int) {
	cValue0 := (C.gint)(param0)

	C.gdk_visual_get_best_with_depth(cValue0)
}

func Fn_gdk_visual_get_best_with_type(param0 int) {
	cValue0 := (C.GdkVisualType)(param0)

	C.gdk_visual_get_best_with_type(cValue0)
}

func Fn_gdk_visual_get_system() {

	C.gdk_visual_get_system()
}

func Fn_gdk_window_new(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int) {
	cValue0 := (*C.GdkWindow)(unsafe.Pointer(param0))
	cValue1 := (*C.GdkWindowAttr)(unsafe.Pointer(param1))
	cValue2 := (C.gint)(param2)

	C.gdk_window_new(cValue0, cValue1, cValue2)
}

// UNSUPPORTED : add_filter : has callback

func Fn_gdk_window_begin_move_drag(paramInstance unsafe.Pointer, param0 int, param1 int, param2 int, param3 uint32) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))
	cValue0 := (C.gint)(param0)
	cValue1 := (C.gint)(param1)
	cValue2 := (C.gint)(param2)
	cValue3 := (C.guint32)(param3)

	C.gdk_window_begin_move_drag(cValueInstance, cValue0, cValue1, cValue2, cValue3)
}

func Fn_gdk_window_begin_paint_rect(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GdkRectangle)(unsafe.Pointer(param0))

	C.gdk_window_begin_paint_rect(cValueInstance, cValue0)
}

func Fn_gdk_window_begin_paint_region(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.cairo_region_t)(unsafe.Pointer(param0))

	C.gdk_window_begin_paint_region(cValueInstance, cValue0)
}

func Fn_gdk_window_begin_resize_drag(paramInstance unsafe.Pointer, param0 int, param1 int, param2 int, param3 int, param4 uint32) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))
	cValue0 := (C.GdkWindowEdge)(param0)
	cValue1 := (C.gint)(param1)
	cValue2 := (C.gint)(param2)
	cValue3 := (C.gint)(param3)
	cValue4 := (C.guint32)(param4)

	C.gdk_window_begin_resize_drag(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4)
}

func Fn_gdk_window_configure_finished(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))

	C.gdk_window_configure_finished(cValueInstance)
}

func Fn_gdk_window_deiconify(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))

	C.gdk_window_deiconify(cValueInstance)
}

func Fn_gdk_window_destroy(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))

	C.gdk_window_destroy(cValueInstance)
}

// UNSUPPORTED : destroy_notify : blacklisted
func Fn_gdk_window_enable_synchronized_configure(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))

	C.gdk_window_enable_synchronized_configure(cValueInstance)
}

func Fn_gdk_window_end_paint(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))

	C.gdk_window_end_paint(cValueInstance)
}

func Fn_gdk_window_focus(paramInstance unsafe.Pointer, param0 uint32) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))
	cValue0 := (C.guint32)(param0)

	C.gdk_window_focus(cValueInstance, cValue0)
}

func Fn_gdk_window_freeze_toplevel_updates_libgtk_only(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))

	C.gdk_window_freeze_toplevel_updates_libgtk_only(cValueInstance)
}

func Fn_gdk_window_freeze_updates(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))

	C.gdk_window_freeze_updates(cValueInstance)
}

func Fn_gdk_window_fullscreen(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))

	C.gdk_window_fullscreen(cValueInstance)
}

func Fn_gdk_window_fullscreen_on_monitor(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))
	cValue0 := (C.gint)(param0)

	C.gdk_window_fullscreen_on_monitor(cValueInstance, cValue0)
}

func Fn_gdk_window_get_children(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))

	C.gdk_window_get_children(cValueInstance)
}

func Fn_gdk_window_get_clip_region(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))

	C.gdk_window_get_clip_region(cValueInstance)
}

func Fn_gdk_window_get_decorations(paramInstance unsafe.Pointer, param0 *int) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GdkWMDecoration)(unsafe.Pointer(param0))

	C.gdk_window_get_decorations(cValueInstance, cValue0)
}

func Fn_gdk_window_get_events(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))

	C.gdk_window_get_events(cValueInstance)
}

func Fn_gdk_window_get_frame_extents(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GdkRectangle)(unsafe.Pointer(param0))

	C.gdk_window_get_frame_extents(cValueInstance, cValue0)
}

func Fn_gdk_window_get_geometry(paramInstance unsafe.Pointer, param0 *int, param1 *int, param2 *int, param3 *int) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.gint)(unsafe.Pointer(param0))
	cValue1 := (*C.gint)(unsafe.Pointer(param1))
	cValue2 := (*C.gint)(unsafe.Pointer(param2))
	cValue3 := (*C.gint)(unsafe.Pointer(param3))

	C.gdk_window_get_geometry(cValueInstance, cValue0, cValue1, cValue2, cValue3)
}

func Fn_gdk_window_get_group(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))

	C.gdk_window_get_group(cValueInstance)
}

func Fn_gdk_window_get_origin(paramInstance unsafe.Pointer, param0 *int, param1 *int) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.gint)(unsafe.Pointer(param0))
	cValue1 := (*C.gint)(unsafe.Pointer(param1))

	C.gdk_window_get_origin(cValueInstance, cValue0, cValue1)
}

func Fn_gdk_window_get_parent(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))

	C.gdk_window_get_parent(cValueInstance)
}

func Fn_gdk_window_get_pointer(paramInstance unsafe.Pointer, param0 *int, param1 *int, param2 *int) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.gint)(unsafe.Pointer(param0))
	cValue1 := (*C.gint)(unsafe.Pointer(param1))
	cValue2 := (*C.GdkModifierType)(unsafe.Pointer(param2))

	C.gdk_window_get_pointer(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gdk_window_get_position(paramInstance unsafe.Pointer, param0 *int, param1 *int) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.gint)(unsafe.Pointer(param0))
	cValue1 := (*C.gint)(unsafe.Pointer(param1))

	C.gdk_window_get_position(cValueInstance, cValue0, cValue1)
}

func Fn_gdk_window_get_root_origin(paramInstance unsafe.Pointer, param0 *int, param1 *int) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.gint)(unsafe.Pointer(param0))
	cValue1 := (*C.gint)(unsafe.Pointer(param1))

	C.gdk_window_get_root_origin(cValueInstance, cValue0, cValue1)
}

func Fn_gdk_window_get_source_events(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))
	cValue0 := (C.GdkInputSource)(param0)

	C.gdk_window_get_source_events(cValueInstance, cValue0)
}

func Fn_gdk_window_get_state(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))

	C.gdk_window_get_state(cValueInstance)
}

func Fn_gdk_window_get_toplevel(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))

	C.gdk_window_get_toplevel(cValueInstance)
}

func Fn_gdk_window_get_update_area(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))

	C.gdk_window_get_update_area(cValueInstance)
}

func Fn_gdk_window_get_user_data(paramInstance unsafe.Pointer, param0 *unsafe.Pointer) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.gpointer)(unsafe.Pointer(param0))

	C.gdk_window_get_user_data(cValueInstance, cValue0)
}

func Fn_gdk_window_get_visible_region(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))

	C.gdk_window_get_visible_region(cValueInstance)
}

func Fn_gdk_window_get_window_type(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))

	C.gdk_window_get_window_type(cValueInstance)
}

func Fn_gdk_window_hide(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))

	C.gdk_window_hide(cValueInstance)
}

func Fn_gdk_window_iconify(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))

	C.gdk_window_iconify(cValueInstance)
}

// UNSUPPORTED : invalidate_maybe_recurse : has callback

func Fn_gdk_window_invalidate_rect(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 bool) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GdkRectangle)(unsafe.Pointer(param0))
	cValue1 := toCBool(param1)

	C.gdk_window_invalidate_rect(cValueInstance, cValue0, cValue1)
}

func Fn_gdk_window_invalidate_region(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 bool) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.cairo_region_t)(unsafe.Pointer(param0))
	cValue1 := toCBool(param1)

	C.gdk_window_invalidate_region(cValueInstance, cValue0, cValue1)
}

func Fn_gdk_window_is_viewable(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))

	C.gdk_window_is_viewable(cValueInstance)
}

func Fn_gdk_window_is_visible(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))

	C.gdk_window_is_visible(cValueInstance)
}

func Fn_gdk_window_lower(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))

	C.gdk_window_lower(cValueInstance)
}

func Fn_gdk_window_maximize(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))

	C.gdk_window_maximize(cValueInstance)
}

func Fn_gdk_window_merge_child_shapes(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))

	C.gdk_window_merge_child_shapes(cValueInstance)
}

func Fn_gdk_window_move(paramInstance unsafe.Pointer, param0 int, param1 int) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))
	cValue0 := (C.gint)(param0)
	cValue1 := (C.gint)(param1)

	C.gdk_window_move(cValueInstance, cValue0, cValue1)
}

func Fn_gdk_window_move_resize(paramInstance unsafe.Pointer, param0 int, param1 int, param2 int, param3 int) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))
	cValue0 := (C.gint)(param0)
	cValue1 := (C.gint)(param1)
	cValue2 := (C.gint)(param2)
	cValue3 := (C.gint)(param3)

	C.gdk_window_move_resize(cValueInstance, cValue0, cValue1, cValue2, cValue3)
}

func Fn_gdk_window_peek_children(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))

	C.gdk_window_peek_children(cValueInstance)
}

func Fn_gdk_window_process_updates(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))
	cValue0 := toCBool(param0)

	C.gdk_window_process_updates(cValueInstance, cValue0)
}

func Fn_gdk_window_raise(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))

	C.gdk_window_raise(cValueInstance)
}

func Fn_gdk_window_register_dnd(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))

	C.gdk_window_register_dnd(cValueInstance)
}

// UNSUPPORTED : remove_filter : has callback

func Fn_gdk_window_reparent(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int, param2 int) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GdkWindow)(unsafe.Pointer(param0))
	cValue1 := (C.gint)(param1)
	cValue2 := (C.gint)(param2)

	C.gdk_window_reparent(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gdk_window_resize(paramInstance unsafe.Pointer, param0 int, param1 int) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))
	cValue0 := (C.gint)(param0)
	cValue1 := (C.gint)(param1)

	C.gdk_window_resize(cValueInstance, cValue0, cValue1)
}

func Fn_gdk_window_scroll(paramInstance unsafe.Pointer, param0 int, param1 int) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))
	cValue0 := (C.gint)(param0)
	cValue1 := (C.gint)(param1)

	C.gdk_window_scroll(cValueInstance, cValue0, cValue1)
}

func Fn_gdk_window_set_accept_focus(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))
	cValue0 := toCBool(param0)

	C.gdk_window_set_accept_focus(cValueInstance, cValue0)
}

func Fn_gdk_window_set_background(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GdkColor)(unsafe.Pointer(param0))

	C.gdk_window_set_background(cValueInstance, cValue0)
}

func Fn_gdk_window_set_background_pattern(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.cairo_pattern_t)(unsafe.Pointer(param0))

	C.gdk_window_set_background_pattern(cValueInstance, cValue0)
}

func Fn_gdk_window_set_background_rgba(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GdkRGBA)(unsafe.Pointer(param0))

	C.gdk_window_set_background_rgba(cValueInstance, cValue0)
}

func Fn_gdk_window_set_child_shapes(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))

	C.gdk_window_set_child_shapes(cValueInstance)
}

func Fn_gdk_window_set_cursor(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GdkCursor)(unsafe.Pointer(param0))

	C.gdk_window_set_cursor(cValueInstance, cValue0)
}

func Fn_gdk_window_set_decorations(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))
	cValue0 := (C.GdkWMDecoration)(param0)

	C.gdk_window_set_decorations(cValueInstance, cValue0)
}

func Fn_gdk_window_set_events(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))
	cValue0 := (C.GdkEventMask)(param0)

	C.gdk_window_set_events(cValueInstance, cValue0)
}

func Fn_gdk_window_set_focus_on_map(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))
	cValue0 := toCBool(param0)

	C.gdk_window_set_focus_on_map(cValueInstance, cValue0)
}

func Fn_gdk_window_set_functions(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))
	cValue0 := (C.GdkWMFunction)(param0)

	C.gdk_window_set_functions(cValueInstance, cValue0)
}

func Fn_gdk_window_set_geometry_hints(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GdkGeometry)(unsafe.Pointer(param0))
	cValue1 := (C.GdkWindowHints)(param1)

	C.gdk_window_set_geometry_hints(cValueInstance, cValue0, cValue1)
}

func Fn_gdk_window_set_group(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GdkWindow)(unsafe.Pointer(param0))

	C.gdk_window_set_group(cValueInstance, cValue0)
}

func Fn_gdk_window_set_icon_list(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GList)(unsafe.Pointer(param0))

	C.gdk_window_set_icon_list(cValueInstance, cValue0)
}

func Fn_gdk_window_set_icon_name(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))
	cValue0 := 42

	C.gdk_window_set_icon_name(cValueInstance, cValue0)
}

// UNSUPPORTED : set_invalidate_handler : has callback

func Fn_gdk_window_set_keep_above(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))
	cValue0 := toCBool(param0)

	C.gdk_window_set_keep_above(cValueInstance, cValue0)
}

func Fn_gdk_window_set_keep_below(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))
	cValue0 := toCBool(param0)

	C.gdk_window_set_keep_below(cValueInstance, cValue0)
}

func Fn_gdk_window_set_modal_hint(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))
	cValue0 := toCBool(param0)

	C.gdk_window_set_modal_hint(cValueInstance, cValue0)
}

func Fn_gdk_window_set_override_redirect(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))
	cValue0 := toCBool(param0)

	C.gdk_window_set_override_redirect(cValueInstance, cValue0)
}

func Fn_gdk_window_set_role(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))
	cValue0 := 42

	C.gdk_window_set_role(cValueInstance, cValue0)
}

func Fn_gdk_window_set_skip_pager_hint(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))
	cValue0 := toCBool(param0)

	C.gdk_window_set_skip_pager_hint(cValueInstance, cValue0)
}

func Fn_gdk_window_set_skip_taskbar_hint(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))
	cValue0 := toCBool(param0)

	C.gdk_window_set_skip_taskbar_hint(cValueInstance, cValue0)
}

func Fn_gdk_window_set_static_gravities(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))
	cValue0 := toCBool(param0)

	C.gdk_window_set_static_gravities(cValueInstance, cValue0)
}

func Fn_gdk_window_set_title(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))
	cValue0 := 42

	C.gdk_window_set_title(cValueInstance, cValue0)
}

func Fn_gdk_window_set_transient_for(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GdkWindow)(unsafe.Pointer(param0))

	C.gdk_window_set_transient_for(cValueInstance, cValue0)
}

func Fn_gdk_window_set_type_hint(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))
	cValue0 := (C.GdkWindowTypeHint)(param0)

	C.gdk_window_set_type_hint(cValueInstance, cValue0)
}

func Fn_gdk_window_set_user_data(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))
	cValue0 := (C.gpointer)(param0)

	C.gdk_window_set_user_data(cValueInstance, cValue0)
}

func Fn_gdk_window_shape_combine_region(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int, param2 int) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.cairo_region_t)(unsafe.Pointer(param0))
	cValue1 := (C.gint)(param1)
	cValue2 := (C.gint)(param2)

	C.gdk_window_shape_combine_region(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gdk_window_show(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))

	C.gdk_window_show(cValueInstance)
}

func Fn_gdk_window_show_unraised(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))

	C.gdk_window_show_unraised(cValueInstance)
}

func Fn_gdk_window_stick(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))

	C.gdk_window_stick(cValueInstance)
}

func Fn_gdk_window_thaw_toplevel_updates_libgtk_only(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))

	C.gdk_window_thaw_toplevel_updates_libgtk_only(cValueInstance)
}

func Fn_gdk_window_thaw_updates(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))

	C.gdk_window_thaw_updates(cValueInstance)
}

func Fn_gdk_window_unfullscreen(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))

	C.gdk_window_unfullscreen(cValueInstance)
}

func Fn_gdk_window_unmaximize(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))

	C.gdk_window_unmaximize(cValueInstance)
}

func Fn_gdk_window_unstick(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))

	C.gdk_window_unstick(cValueInstance)
}

func Fn_gdk_window_withdraw(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))

	C.gdk_window_withdraw(cValueInstance)
}

func Fn_gdk_window_at_pointer(param0 *int, param1 *int) {
	cValue0 := (*C.gint)(unsafe.Pointer(param0))
	cValue1 := (*C.gint)(unsafe.Pointer(param1))

	C.gdk_window_at_pointer(cValue0, cValue1)
}

func Fn_gdk_window_constrain_size(param0 unsafe.Pointer, param1 int, param2 int, param3 int, param4 *int, param5 *int) {
	cValue0 := (*C.GdkGeometry)(unsafe.Pointer(param0))
	cValue1 := (C.GdkWindowHints)(param1)
	cValue2 := (C.gint)(param2)
	cValue3 := (C.gint)(param3)
	cValue4 := (*C.gint)(unsafe.Pointer(param4))
	cValue5 := (*C.gint)(unsafe.Pointer(param5))

	C.gdk_window_constrain_size(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5)
}

func Fn_gdk_window_process_all_updates() {

	C.gdk_window_process_all_updates()
}

func Fn_gdk_window_set_debug_updates(param0 bool) {
	cValue0 := toCBool(param0)

	C.gdk_window_set_debug_updates(cValue0)
}
