// Code generated - DO NOT EDIT.
// +build gdk_2.18

package gdk

import "unsafe"

// #include <gdk/gdk.h>
// #include <stdlib.h>
import "C"

func toCBool(b bool) C.gboolean {
	if b {
		return C.TRUE
	}
	return C.FALSE
}
func toGoBool(b C.gboolean) bool {
	return b == C.TRUE
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

func Fn_gdk_atom_name(atom unsafe.Pointer) string {
	c_atom := (C.GdkAtom)(atom)

	ret := C.gdk_atom_name(c_atom)

	return C.GoString(ret)
}

func Fn_gdk_atom_intern(atomName string, onlyIfExists bool) unsafe.Pointer {
	c_atomName := (*C.gchar)(C.CString(atomName))
	defer C.free(unsafe.Pointer(c_atomName))

	c_onlyIfExists := toCBool(onlyIfExists)

	ret := C.gdk_atom_intern(c_atomName, c_onlyIfExists)

	return unsafe.Pointer(ret)
}

func Fn_gdk_atom_intern_static_string(atomName string) unsafe.Pointer {
	c_atomName := (*C.gchar)(C.CString(atomName))
	defer C.free(unsafe.Pointer(c_atomName))

	ret := C.gdk_atom_intern_static_string(c_atomName)

	return unsafe.Pointer(ret)
}

func Fn_gdk_color_copy(color unsafe.Pointer) unsafe.Pointer {
	c_color := (*C.GdkColor)(unsafe.Pointer(color))

	ret := C.gdk_color_copy(c_color)

	return unsafe.Pointer(ret)
}

func Fn_gdk_color_equal(colora unsafe.Pointer, colorb unsafe.Pointer) bool {
	c_colora := (*C.GdkColor)(unsafe.Pointer(colora))

	c_colorb := (*C.GdkColor)(unsafe.Pointer(colorb))

	ret := C.gdk_color_equal(c_colora, c_colorb)

	return toGoBool(ret)
}

func Fn_gdk_color_free(color unsafe.Pointer) {
	c_color := (*C.GdkColor)(unsafe.Pointer(color))

	C.gdk_color_free(c_color)
}

func Fn_gdk_color_hash(color unsafe.Pointer) uint {
	c_color := (*C.GdkColor)(unsafe.Pointer(color))

	ret := C.gdk_color_hash(c_color)

	return (uint)(ret)
}

func Fn_gdk_color_to_string(color unsafe.Pointer) string {
	c_color := (*C.GdkColor)(unsafe.Pointer(color))

	ret := C.gdk_color_to_string(c_color)

	return C.GoString(ret)
}

func Fn_gdk_color_parse(spec string, color unsafe.Pointer) bool {
	c_spec := (*C.gchar)(C.CString(spec))
	defer C.free(unsafe.Pointer(c_spec))

	c_color := (*C.GdkColor)(unsafe.Pointer(color))

	ret := C.gdk_color_parse(c_spec, c_color)

	return toGoBool(ret)
}

func Fn_gdk_frame_timings_get_frame_time(timings unsafe.Pointer) int64 {
	c_timings := (*C.GdkFrameTimings)(unsafe.Pointer(timings))

	ret := C.gdk_frame_timings_get_frame_time(c_timings)

	return (int64)(ret)
}

func Fn_gdk_rectangle_intersect(src1 unsafe.Pointer, src2 unsafe.Pointer, dest unsafe.Pointer) bool {
	c_src1 := (*C.GdkRectangle)(unsafe.Pointer(src1))

	c_src2 := (*C.GdkRectangle)(unsafe.Pointer(src2))

	c_dest := (*C.GdkRectangle)(unsafe.Pointer(dest))

	ret := C.gdk_rectangle_intersect(c_src1, c_src2, c_dest)

	return toGoBool(ret)
}

func Fn_gdk_rectangle_union(src1 unsafe.Pointer, src2 unsafe.Pointer, dest unsafe.Pointer) {
	c_src1 := (*C.GdkRectangle)(unsafe.Pointer(src1))

	c_src2 := (*C.GdkRectangle)(unsafe.Pointer(src2))

	c_dest := (*C.GdkRectangle)(unsafe.Pointer(dest))

	C.gdk_rectangle_union(c_src1, c_src2, c_dest)
}

func Fn_gdk_add_option_entries_libgtk_only(group unsafe.Pointer) {
	c_group := (*C.GOptionGroup)(unsafe.Pointer(group))

	C.gdk_add_option_entries_libgtk_only(c_group)
}

func Fn_gdk_beep() {
	C.gdk_beep()
}

func Fn_gdk_cairo_create(window unsafe.Pointer) unsafe.Pointer {
	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	ret := C.gdk_cairo_create(c_window)

	return unsafe.Pointer(ret)
}

func Fn_gdk_cairo_get_clip_rectangle(cr unsafe.Pointer, rect unsafe.Pointer) bool {
	c_cr := (*C.cairo_t)(unsafe.Pointer(cr))

	c_rect := (*C.GdkRectangle)(unsafe.Pointer(rect))

	ret := C.gdk_cairo_get_clip_rectangle(c_cr, c_rect)

	return toGoBool(ret)
}

func Fn_gdk_cairo_rectangle(cr unsafe.Pointer, rectangle unsafe.Pointer) {
	c_cr := (*C.cairo_t)(unsafe.Pointer(cr))

	c_rectangle := (*C.GdkRectangle)(unsafe.Pointer(rectangle))

	C.gdk_cairo_rectangle(c_cr, c_rectangle)
}

func Fn_gdk_cairo_region(cr unsafe.Pointer, region unsafe.Pointer) {
	c_cr := (*C.cairo_t)(unsafe.Pointer(cr))

	c_region := (*C.cairo_region_t)(unsafe.Pointer(region))

	C.gdk_cairo_region(c_cr, c_region)
}

func Fn_gdk_cairo_region_create_from_surface(surface unsafe.Pointer) unsafe.Pointer {
	c_surface := (*C.cairo_surface_t)(unsafe.Pointer(surface))

	ret := C.gdk_cairo_region_create_from_surface(c_surface)

	return unsafe.Pointer(ret)
}

func Fn_gdk_cairo_set_source_color(cr unsafe.Pointer, color unsafe.Pointer) {
	c_cr := (*C.cairo_t)(unsafe.Pointer(cr))

	c_color := (*C.GdkColor)(unsafe.Pointer(color))

	C.gdk_cairo_set_source_color(c_cr, c_color)
}

func Fn_gdk_cairo_set_source_pixbuf(cr unsafe.Pointer, pixbuf unsafe.Pointer, pixbufX float64, pixbufY float64) {
	c_cr := (*C.cairo_t)(unsafe.Pointer(cr))

	c_pixbuf := (*C.GdkPixbuf)(unsafe.Pointer(pixbuf))

	c_pixbufX := (C.gdouble)(pixbufX)

	c_pixbufY := (C.gdouble)(pixbufY)

	C.gdk_cairo_set_source_pixbuf(c_cr, c_pixbuf, c_pixbufX, c_pixbufY)
}

func Fn_gdk_drag_abort(context unsafe.Pointer, time uint32) {
	c_context := (*C.GdkDragContext)(unsafe.Pointer(context))

	c_time := (C.guint32)(time)

	C.gdk_drag_abort(c_context, c_time)
}

func Fn_gdk_drag_begin(window unsafe.Pointer, targets unsafe.Pointer) unsafe.Pointer {
	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	c_targets := (*C.GList)(unsafe.Pointer(targets))

	ret := C.gdk_drag_begin(c_window, c_targets)

	return unsafe.Pointer(ret)
}

func Fn_gdk_drag_begin_for_device(window unsafe.Pointer, device unsafe.Pointer, targets unsafe.Pointer) unsafe.Pointer {
	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	c_device := (*C.GdkDevice)(unsafe.Pointer(device))

	c_targets := (*C.GList)(unsafe.Pointer(targets))

	ret := C.gdk_drag_begin_for_device(c_window, c_device, c_targets)

	return unsafe.Pointer(ret)
}

func Fn_gdk_drag_drop(context unsafe.Pointer, time uint32) {
	c_context := (*C.GdkDragContext)(unsafe.Pointer(context))

	c_time := (C.guint32)(time)

	C.gdk_drag_drop(c_context, c_time)
}

func Fn_gdk_drag_drop_succeeded(context unsafe.Pointer) bool {
	c_context := (*C.GdkDragContext)(unsafe.Pointer(context))

	ret := C.gdk_drag_drop_succeeded(c_context)

	return toGoBool(ret)
}

// UNSUPPORTED : gdk_drag_find_window_for_screen : parameter 'dest_window' is non array with indirect count > 1

func Fn_gdk_drag_get_selection(context unsafe.Pointer) unsafe.Pointer {
	c_context := (*C.GdkDragContext)(unsafe.Pointer(context))

	ret := C.gdk_drag_get_selection(c_context)

	return unsafe.Pointer(ret)
}

func Fn_gdk_drag_motion(context unsafe.Pointer, destWindow unsafe.Pointer, protocol int, xRoot int, yRoot int, suggestedAction int, possibleActions int, time uint32) bool {
	c_context := (*C.GdkDragContext)(unsafe.Pointer(context))

	c_destWindow := (*C.GdkWindow)(unsafe.Pointer(destWindow))

	c_protocol := (C.GdkDragProtocol)(protocol)

	c_xRoot := (C.gint)(xRoot)

	c_yRoot := (C.gint)(yRoot)

	c_suggestedAction := (C.GdkDragAction)(suggestedAction)

	c_possibleActions := (C.GdkDragAction)(possibleActions)

	c_time := (C.guint32)(time)

	ret := C.gdk_drag_motion(c_context, c_destWindow, c_protocol, c_xRoot, c_yRoot, c_suggestedAction, c_possibleActions, c_time)

	return toGoBool(ret)
}

func Fn_gdk_drag_status(context unsafe.Pointer, action int, time uint32) {
	c_context := (*C.GdkDragContext)(unsafe.Pointer(context))

	c_action := (C.GdkDragAction)(action)

	c_time := (C.guint32)(time)

	C.gdk_drag_status(c_context, c_action, c_time)
}

func Fn_gdk_drop_finish(context unsafe.Pointer, success bool, time uint32) {
	c_context := (*C.GdkDragContext)(unsafe.Pointer(context))

	c_success := toCBool(success)

	c_time := (C.guint32)(time)

	C.gdk_drop_finish(c_context, c_success, c_time)
}

func Fn_gdk_drop_reply(context unsafe.Pointer, accepted bool, time uint32) {
	c_context := (*C.GdkDragContext)(unsafe.Pointer(context))

	c_accepted := toCBool(accepted)

	c_time := (C.guint32)(time)

	C.gdk_drop_reply(c_context, c_accepted, c_time)
}

func Fn_gdk_error_trap_pop() int {
	ret := C.gdk_error_trap_pop()

	return (int)(ret)
}

func Fn_gdk_error_trap_push() {
	C.gdk_error_trap_push()
}

// UNSUPPORTED : gdk_event_handler_set : parameter 'func' is callback

func Fn_gdk_events_pending() bool {
	ret := C.gdk_events_pending()

	return toGoBool(ret)
}

func Fn_gdk_flush() {
	C.gdk_flush()
}

func Fn_gdk_get_default_root_window() unsafe.Pointer {
	ret := C.gdk_get_default_root_window()

	return unsafe.Pointer(ret)
}

func Fn_gdk_get_display() string {
	ret := C.gdk_get_display()

	return C.GoString(ret)
}

func Fn_gdk_get_display_arg_name() string {
	ret := C.gdk_get_display_arg_name()

	return C.GoString(ret)
}

func Fn_gdk_get_program_class() string {
	ret := C.gdk_get_program_class()

	return C.GoString(ret)
}

func Fn_gdk_get_show_events() bool {
	ret := C.gdk_get_show_events()

	return toGoBool(ret)
}

func Fn_gdk_init(argc *int, argv *[]string) {
	c_argc := (*C.gint)(unsafe.Pointer(argc))

	var c_argvArrayPointer **C.gchar
	c_argv := &c_argvArrayPointer
	argvIndirected := *argv
	argvIndirectedLen := len(argvIndirected)
	c_argvArray := C.malloc((C.ulong)(argvIndirectedLen) * C.sizeof_gpointer)
	defer C.free(unsafe.Pointer(c_argvArray))
	argvIndirectedSlice := (*[1 << 30](*C.gchar))(unsafe.Pointer(c_argvArray))[:argvIndirectedLen:argvIndirectedLen]
	for argvIndirectedi, argvIndirectedString := range argvIndirected {
		argvIndirectedSlice[argvIndirectedi] = (*C.gchar)(C.CString(argvIndirectedString))
		defer C.free(unsafe.Pointer(argvIndirectedSlice[argvIndirectedi]))
	}
	if len(argvIndirectedSlice) > 0 {
		c_argvArrayPointer = &argvIndirectedSlice[0]
	}

	C.gdk_init(c_argc, c_argv)

	argvOutLen := int(*c_argc)
	argvOut := make([]string, argvOutLen, argvOutLen)
	if argvOutLen > 0 {
		argvOutCSlice := (*[1 << 30](*C.gchar))(unsafe.Pointer(c_argvArrayPointer))[:argvOutLen:argvOutLen]
		for argvOuti, argvOutCString := range argvOutCSlice {
			argvOut[argvOuti] = C.GoString(argvOutCString)
		}
	}
	*argv = argvOut
}

func Fn_gdk_init_check(argc *int, argv *[]string) bool {
	c_argc := (*C.gint)(unsafe.Pointer(argc))

	var c_argvArrayPointer **C.gchar
	c_argv := &c_argvArrayPointer
	argvIndirected := *argv
	argvIndirectedLen := len(argvIndirected)
	c_argvArray := C.malloc((C.ulong)(argvIndirectedLen) * C.sizeof_gpointer)
	defer C.free(unsafe.Pointer(c_argvArray))
	argvIndirectedSlice := (*[1 << 30](*C.gchar))(unsafe.Pointer(c_argvArray))[:argvIndirectedLen:argvIndirectedLen]
	for argvIndirectedi, argvIndirectedString := range argvIndirected {
		argvIndirectedSlice[argvIndirectedi] = (*C.gchar)(C.CString(argvIndirectedString))
		defer C.free(unsafe.Pointer(argvIndirectedSlice[argvIndirectedi]))
	}
	if len(argvIndirectedSlice) > 0 {
		c_argvArrayPointer = &argvIndirectedSlice[0]
	}

	ret := C.gdk_init_check(c_argc, c_argv)

	argvOutLen := int(*c_argc)
	argvOut := make([]string, argvOutLen, argvOutLen)
	if argvOutLen > 0 {
		argvOutCSlice := (*[1 << 30](*C.gchar))(unsafe.Pointer(c_argvArrayPointer))[:argvOutLen:argvOutLen]
		for argvOuti, argvOutCString := range argvOutCSlice {
			argvOut[argvOuti] = C.GoString(argvOutCString)
		}
	}
	*argv = argvOut

	return toGoBool(ret)
}

func Fn_gdk_keyboard_grab(window unsafe.Pointer, ownerEvents bool, time uint32) int {
	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	c_ownerEvents := toCBool(ownerEvents)

	c_time := (C.guint32)(time)

	ret := C.gdk_keyboard_grab(c_window, c_ownerEvents, c_time)

	return (int)(ret)
}

func Fn_gdk_keyboard_ungrab(time uint32) {
	c_time := (C.guint32)(time)

	C.gdk_keyboard_ungrab(c_time)
}

func Fn_gdk_keyval_convert_case(symbol uint, lower *uint, upper *uint) {
	c_symbol := (C.guint)(symbol)

	c_lower := (*C.guint)(unsafe.Pointer(lower))

	c_upper := (*C.guint)(unsafe.Pointer(upper))

	C.gdk_keyval_convert_case(c_symbol, c_lower, c_upper)
}

func Fn_gdk_keyval_from_name(keyvalName string) uint {
	c_keyvalName := (*C.gchar)(C.CString(keyvalName))
	defer C.free(unsafe.Pointer(c_keyvalName))

	ret := C.gdk_keyval_from_name(c_keyvalName)

	return (uint)(ret)
}

func Fn_gdk_keyval_is_lower(keyval uint) bool {
	c_keyval := (C.guint)(keyval)

	ret := C.gdk_keyval_is_lower(c_keyval)

	return toGoBool(ret)
}

func Fn_gdk_keyval_is_upper(keyval uint) bool {
	c_keyval := (C.guint)(keyval)

	ret := C.gdk_keyval_is_upper(c_keyval)

	return toGoBool(ret)
}

func Fn_gdk_keyval_name(keyval uint) string {
	c_keyval := (C.guint)(keyval)

	ret := C.gdk_keyval_name(c_keyval)

	return C.GoString(ret)
}

func Fn_gdk_keyval_to_lower(keyval uint) uint {
	c_keyval := (C.guint)(keyval)

	ret := C.gdk_keyval_to_lower(c_keyval)

	return (uint)(ret)
}

func Fn_gdk_keyval_to_unicode(keyval uint) uint32 {
	c_keyval := (C.guint)(keyval)

	ret := C.gdk_keyval_to_unicode(c_keyval)

	return (uint32)(ret)
}

func Fn_gdk_keyval_to_upper(keyval uint) uint {
	c_keyval := (C.guint)(keyval)

	ret := C.gdk_keyval_to_upper(c_keyval)

	return (uint)(ret)
}

func Fn_gdk_list_visuals() unsafe.Pointer {
	ret := C.gdk_list_visuals()

	return unsafe.Pointer(ret)
}

func Fn_gdk_notify_startup_complete() {
	C.gdk_notify_startup_complete()
}

func Fn_gdk_notify_startup_complete_with_id(startupId string) {
	c_startupId := (*C.gchar)(C.CString(startupId))
	defer C.free(unsafe.Pointer(c_startupId))

	C.gdk_notify_startup_complete_with_id(c_startupId)
}

func Fn_gdk_offscreen_window_get_embedder(window unsafe.Pointer) unsafe.Pointer {
	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	ret := C.gdk_offscreen_window_get_embedder(c_window)

	return unsafe.Pointer(ret)
}

func Fn_gdk_offscreen_window_get_surface(window unsafe.Pointer) unsafe.Pointer {
	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	ret := C.gdk_offscreen_window_get_surface(c_window)

	return unsafe.Pointer(ret)
}

func Fn_gdk_offscreen_window_set_embedder(window unsafe.Pointer, embedder unsafe.Pointer) {
	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	c_embedder := (*C.GdkWindow)(unsafe.Pointer(embedder))

	C.gdk_offscreen_window_set_embedder(c_window, c_embedder)
}

func Fn_gdk_pango_context_get() unsafe.Pointer {
	ret := C.gdk_pango_context_get()

	return unsafe.Pointer(ret)
}

func Fn_gdk_pango_context_get_for_screen(screen unsafe.Pointer) unsafe.Pointer {
	c_screen := (*C.GdkScreen)(unsafe.Pointer(screen))

	ret := C.gdk_pango_context_get_for_screen(c_screen)

	return unsafe.Pointer(ret)
}

func Fn_gdk_pango_layout_get_clip_region(layout unsafe.Pointer, xOrigin int, yOrigin int, indexRanges *int, nRanges int) unsafe.Pointer {
	c_layout := (*C.PangoLayout)(unsafe.Pointer(layout))

	c_xOrigin := (C.gint)(xOrigin)

	c_yOrigin := (C.gint)(yOrigin)

	c_indexRanges := (*C.gint)(unsafe.Pointer(indexRanges))

	c_nRanges := (C.gint)(nRanges)

	ret := C.gdk_pango_layout_get_clip_region(c_layout, c_xOrigin, c_yOrigin, c_indexRanges, c_nRanges)

	return unsafe.Pointer(ret)
}

func Fn_gdk_pango_layout_line_get_clip_region(line unsafe.Pointer, xOrigin int, yOrigin int, indexRanges []int, nRanges int) unsafe.Pointer {
	c_line := (*C.PangoLayoutLine)(unsafe.Pointer(line))

	c_xOrigin := (C.gint)(xOrigin)

	c_yOrigin := (C.gint)(yOrigin)

	c_indexRanges := (*C.gint)(unsafe.Pointer(&indexRanges[0]))

	c_nRanges := (C.gint)(nRanges)

	ret := C.gdk_pango_layout_line_get_clip_region(c_line, c_xOrigin, c_yOrigin, c_indexRanges, c_nRanges)

	return unsafe.Pointer(ret)
}

func Fn_gdk_parse_args(argc *int, argv *[]string) {
	c_argc := (*C.gint)(unsafe.Pointer(argc))

	var c_argvArrayPointer **C.gchar
	c_argv := &c_argvArrayPointer
	argvIndirected := *argv
	argvIndirectedLen := len(argvIndirected)
	c_argvArray := C.malloc((C.ulong)(argvIndirectedLen) * C.sizeof_gpointer)
	defer C.free(unsafe.Pointer(c_argvArray))
	argvIndirectedSlice := (*[1 << 30](*C.gchar))(unsafe.Pointer(c_argvArray))[:argvIndirectedLen:argvIndirectedLen]
	for argvIndirectedi, argvIndirectedString := range argvIndirected {
		argvIndirectedSlice[argvIndirectedi] = (*C.gchar)(C.CString(argvIndirectedString))
		defer C.free(unsafe.Pointer(argvIndirectedSlice[argvIndirectedi]))
	}
	if len(argvIndirectedSlice) > 0 {
		c_argvArrayPointer = &argvIndirectedSlice[0]
	}

	C.gdk_parse_args(c_argc, c_argv)

	argvOutLen := int(*c_argc)
	argvOut := make([]string, argvOutLen, argvOutLen)
	if argvOutLen > 0 {
		argvOutCSlice := (*[1 << 30](*C.gchar))(unsafe.Pointer(c_argvArrayPointer))[:argvOutLen:argvOutLen]
		for argvOuti, argvOutCString := range argvOutCSlice {
			argvOut[argvOuti] = C.GoString(argvOutCString)
		}
	}
	*argv = argvOut
}

func Fn_gdk_pixbuf_get_from_surface(surface unsafe.Pointer, srcX int, srcY int, width int, height int) unsafe.Pointer {
	c_surface := (*C.cairo_surface_t)(unsafe.Pointer(surface))

	c_srcX := (C.gint)(srcX)

	c_srcY := (C.gint)(srcY)

	c_width := (C.gint)(width)

	c_height := (C.gint)(height)

	ret := C.gdk_pixbuf_get_from_surface(c_surface, c_srcX, c_srcY, c_width, c_height)

	return unsafe.Pointer(ret)
}

func Fn_gdk_pixbuf_get_from_window(window unsafe.Pointer, srcX int, srcY int, width int, height int) unsafe.Pointer {
	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	c_srcX := (C.gint)(srcX)

	c_srcY := (C.gint)(srcY)

	c_width := (C.gint)(width)

	c_height := (C.gint)(height)

	ret := C.gdk_pixbuf_get_from_window(c_window, c_srcX, c_srcY, c_width, c_height)

	return unsafe.Pointer(ret)
}

func Fn_gdk_pointer_grab(window unsafe.Pointer, ownerEvents bool, eventMask int, confineTo unsafe.Pointer, cursor unsafe.Pointer, time uint32) int {
	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	c_ownerEvents := toCBool(ownerEvents)

	c_eventMask := (C.GdkEventMask)(eventMask)

	c_confineTo := (*C.GdkWindow)(unsafe.Pointer(confineTo))

	c_cursor := (*C.GdkCursor)(unsafe.Pointer(cursor))

	c_time := (C.guint32)(time)

	ret := C.gdk_pointer_grab(c_window, c_ownerEvents, c_eventMask, c_confineTo, c_cursor, c_time)

	return (int)(ret)
}

func Fn_gdk_pointer_is_grabbed() bool {
	ret := C.gdk_pointer_is_grabbed()

	return toGoBool(ret)
}

func Fn_gdk_pointer_ungrab(time uint32) {
	c_time := (C.guint32)(time)

	C.gdk_pointer_ungrab(c_time)
}

func Fn_gdk_pre_parse_libgtk_only() {
	C.gdk_pre_parse_libgtk_only()
}

func Fn_gdk_property_change(window unsafe.Pointer, property unsafe.Pointer, type_ unsafe.Pointer, format int, mode int, data *uint8, nelements int) {
	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	c_property := (C.GdkAtom)(property)

	c_type_ := (C.GdkAtom)(type_)

	c_format := (C.gint)(format)

	c_mode := (C.GdkPropMode)(mode)

	c_data := (*C.guchar)(unsafe.Pointer(data))

	c_nelements := (C.gint)(nelements)

	C.gdk_property_change(c_window, c_property, c_type_, c_format, c_mode, c_data, c_nelements)
}

func Fn_gdk_property_delete(window unsafe.Pointer, property unsafe.Pointer) {
	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	c_property := (C.GdkAtom)(property)

	C.gdk_property_delete(c_window, c_property)
}

func Fn_gdk_property_get(window unsafe.Pointer, property unsafe.Pointer, type_ unsafe.Pointer, offset uint64, length uint64, pdelete int, actualPropertyType unsafe.Pointer, actualFormat *int, actualLength *int, data *[]uint8) bool {
	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	c_property := (C.GdkAtom)(property)

	c_type_ := (C.GdkAtom)(type_)

	c_offset := (C.gulong)(offset)

	c_length := (C.gulong)(length)

	c_pdelete := (C.gint)(pdelete)

	c_actualPropertyType := (*C.GdkAtom)(unsafe.Pointer(actualPropertyType))

	c_actualFormat := (*C.gint)(unsafe.Pointer(actualFormat))

	c_actualLength := (*C.gint)(unsafe.Pointer(actualLength))

	var c_dataArrayPointer (*C.guchar)
	c_data := &c_dataArrayPointer

	ret := C.gdk_property_get(c_window, c_property, c_type_, c_offset, c_length, c_pdelete, c_actualPropertyType, c_actualFormat, c_actualLength, c_data)

	dataOutLen := int(*c_actualLength)
	dataOut := make([]uint8, dataOutLen, dataOutLen)
	if dataOutLen > 0 {
		dataOut = (*[1 << 30](uint8))(unsafe.Pointer(c_dataArrayPointer))[:dataOutLen:dataOutLen]
	}
	*data = dataOut

	return toGoBool(ret)
}

func Fn_gdk_query_depths(depths *[]int, count *int) {
	var c_depthsArrayPointer (*C.gint)
	c_depths := &c_depthsArrayPointer

	c_count := (*C.gint)(unsafe.Pointer(count))

	C.gdk_query_depths(c_depths, c_count)

	depthsOutLen := int(*c_count)
	depthsOut := make([]int, depthsOutLen, depthsOutLen)
	if depthsOutLen > 0 {
		depthsOut = (*[1 << 30](int))(unsafe.Pointer(c_depthsArrayPointer))[:depthsOutLen:depthsOutLen]
	}
	*depths = depthsOut
}

func Fn_gdk_query_visual_types(visualTypes *[]int, count *int) {
	var c_visualTypesArrayPointer (*C.GdkVisualType)
	c_visualTypes := &c_visualTypesArrayPointer

	c_count := (*C.gint)(unsafe.Pointer(count))

	C.gdk_query_visual_types(c_visualTypes, c_count)

	visualTypesOutLen := int(*c_count)
	visualTypesOut := make([]int, visualTypesOutLen, visualTypesOutLen)
	if visualTypesOutLen > 0 {
		visualTypesOut = (*[1 << 30](int))(unsafe.Pointer(c_visualTypesArrayPointer))[:visualTypesOutLen:visualTypesOutLen]
	}
	*visualTypes = visualTypesOut
}

func Fn_gdk_selection_convert(requestor unsafe.Pointer, selection unsafe.Pointer, target unsafe.Pointer, time uint32) {
	c_requestor := (*C.GdkWindow)(unsafe.Pointer(requestor))

	c_selection := (C.GdkAtom)(selection)

	c_target := (C.GdkAtom)(target)

	c_time := (C.guint32)(time)

	C.gdk_selection_convert(c_requestor, c_selection, c_target, c_time)
}

func Fn_gdk_selection_owner_get(selection unsafe.Pointer) unsafe.Pointer {
	c_selection := (C.GdkAtom)(selection)

	ret := C.gdk_selection_owner_get(c_selection)

	return unsafe.Pointer(ret)
}

func Fn_gdk_selection_owner_get_for_display(display unsafe.Pointer, selection unsafe.Pointer) unsafe.Pointer {
	c_display := (*C.GdkDisplay)(unsafe.Pointer(display))

	c_selection := (C.GdkAtom)(selection)

	ret := C.gdk_selection_owner_get_for_display(c_display, c_selection)

	return unsafe.Pointer(ret)
}

func Fn_gdk_selection_owner_set(owner unsafe.Pointer, selection unsafe.Pointer, time uint32, sendEvent bool) bool {
	c_owner := (*C.GdkWindow)(unsafe.Pointer(owner))

	c_selection := (C.GdkAtom)(selection)

	c_time := (C.guint32)(time)

	c_sendEvent := toCBool(sendEvent)

	ret := C.gdk_selection_owner_set(c_owner, c_selection, c_time, c_sendEvent)

	return toGoBool(ret)
}

func Fn_gdk_selection_owner_set_for_display(display unsafe.Pointer, owner unsafe.Pointer, selection unsafe.Pointer, time uint32, sendEvent bool) bool {
	c_display := (*C.GdkDisplay)(unsafe.Pointer(display))

	c_owner := (*C.GdkWindow)(unsafe.Pointer(owner))

	c_selection := (C.GdkAtom)(selection)

	c_time := (C.guint32)(time)

	c_sendEvent := toCBool(sendEvent)

	ret := C.gdk_selection_owner_set_for_display(c_display, c_owner, c_selection, c_time, c_sendEvent)

	return toGoBool(ret)
}

// UNSUPPORTED : gdk_selection_property_get : parameter 'data' is non array with indirect count > 1

func Fn_gdk_selection_send_notify(requestor unsafe.Pointer, selection unsafe.Pointer, target unsafe.Pointer, property unsafe.Pointer, time uint32) {
	c_requestor := (*C.GdkWindow)(unsafe.Pointer(requestor))

	c_selection := (C.GdkAtom)(selection)

	c_target := (C.GdkAtom)(target)

	c_property := (C.GdkAtom)(property)

	c_time := (C.guint32)(time)

	C.gdk_selection_send_notify(c_requestor, c_selection, c_target, c_property, c_time)
}

func Fn_gdk_selection_send_notify_for_display(display unsafe.Pointer, requestor unsafe.Pointer, selection unsafe.Pointer, target unsafe.Pointer, property unsafe.Pointer, time uint32) {
	c_display := (*C.GdkDisplay)(unsafe.Pointer(display))

	c_requestor := (*C.GdkWindow)(unsafe.Pointer(requestor))

	c_selection := (C.GdkAtom)(selection)

	c_target := (C.GdkAtom)(target)

	c_property := (C.GdkAtom)(property)

	c_time := (C.guint32)(time)

	C.gdk_selection_send_notify_for_display(c_display, c_requestor, c_selection, c_target, c_property, c_time)
}

func Fn_gdk_set_double_click_time(msec uint) {
	c_msec := (C.guint)(msec)

	C.gdk_set_double_click_time(c_msec)
}

func Fn_gdk_set_program_class(programClass string) {
	c_programClass := (*C.gchar)(C.CString(programClass))
	defer C.free(unsafe.Pointer(c_programClass))

	C.gdk_set_program_class(c_programClass)
}

func Fn_gdk_set_show_events(showEvents bool) {
	c_showEvents := toCBool(showEvents)

	C.gdk_set_show_events(c_showEvents)
}

func Fn_gdk_setting_get(name string, value unsafe.Pointer) bool {
	c_name := (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(c_name))

	c_value := (*C.GValue)(unsafe.Pointer(value))

	ret := C.gdk_setting_get(c_name, c_value)

	return toGoBool(ret)
}

// UNSUPPORTED : gdk_synthesize_window_state : blacklisted

func Fn_gdk_test_render_sync(window unsafe.Pointer) {
	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	C.gdk_test_render_sync(c_window)
}

func Fn_gdk_test_simulate_button(window unsafe.Pointer, x int, y int, button uint, modifiers int, buttonPressrelease int) bool {
	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	c_button := (C.guint)(button)

	c_modifiers := (C.GdkModifierType)(modifiers)

	c_buttonPressrelease := (C.GdkEventType)(buttonPressrelease)

	ret := C.gdk_test_simulate_button(c_window, c_x, c_y, c_button, c_modifiers, c_buttonPressrelease)

	return toGoBool(ret)
}

func Fn_gdk_test_simulate_key(window unsafe.Pointer, x int, y int, keyval uint, modifiers int, keyPressrelease int) bool {
	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	c_keyval := (C.guint)(keyval)

	c_modifiers := (C.GdkModifierType)(modifiers)

	c_keyPressrelease := (C.GdkEventType)(keyPressrelease)

	ret := C.gdk_test_simulate_key(c_window, c_x, c_y, c_keyval, c_modifiers, c_keyPressrelease)

	return toGoBool(ret)
}

func Fn_gdk_text_property_to_utf8_list_for_display(display unsafe.Pointer, encoding unsafe.Pointer, format int, text []uint8, length int, list *[]string) int {
	c_display := (*C.GdkDisplay)(unsafe.Pointer(display))

	c_encoding := (C.GdkAtom)(encoding)

	c_format := (C.gint)(format)

	c_text := (*C.guchar)(unsafe.Pointer(&text[0]))

	c_length := (C.gint)(length)

	// TODO
	var c_listArrayPointer **C.gchar
	c_list := &c_listArrayPointer

	ret := C.gdk_text_property_to_utf8_list_for_display(c_display, c_encoding, c_format, c_text, c_length, c_list)
	// TODO - 0 terminated array

	return (int)(ret)
}

// UNSUPPORTED : gdk_threads_add_idle : parameter 'function' is callback

// UNSUPPORTED : gdk_threads_add_idle_full : parameter 'function' is callback

// UNSUPPORTED : gdk_threads_add_timeout : parameter 'function' is callback

// UNSUPPORTED : gdk_threads_add_timeout_full : parameter 'function' is callback

// UNSUPPORTED : gdk_threads_add_timeout_seconds : parameter 'function' is callback

// UNSUPPORTED : gdk_threads_add_timeout_seconds_full : parameter 'function' is callback

func Fn_gdk_threads_enter() {
	C.gdk_threads_enter()
}

func Fn_gdk_threads_init() {
	C.gdk_threads_init()
}

func Fn_gdk_threads_leave() {
	C.gdk_threads_leave()
}

// UNSUPPORTED : gdk_threads_set_lock_functions : parameter 'enter_fn' is callback

func Fn_gdk_unicode_to_keyval(wc uint32) uint {
	c_wc := (C.guint32)(wc)

	ret := C.gdk_unicode_to_keyval(c_wc)

	return (uint)(ret)
}

func Fn_gdk_utf8_to_string_target(str string) string {
	c_str := (*C.gchar)(C.CString(str))
	defer C.free(unsafe.Pointer(c_str))

	ret := C.gdk_utf8_to_string_target(c_str)

	return C.GoString(ret)
}

func Fn_gdk_app_launch_context_new() unsafe.Pointer {
	ret := C.gdk_app_launch_context_new()

	return unsafe.Pointer(ret)
}

func Fn_gdk_app_launch_context_set_desktop(context unsafe.Pointer, desktop int) {
	c_context := (*C.GdkAppLaunchContext)(unsafe.Pointer(context))

	c_desktop := (C.gint)(desktop)

	C.gdk_app_launch_context_set_desktop(c_context, c_desktop)
}

func Fn_gdk_app_launch_context_set_display(context unsafe.Pointer, display unsafe.Pointer) {
	c_context := (*C.GdkAppLaunchContext)(unsafe.Pointer(context))

	c_display := (*C.GdkDisplay)(unsafe.Pointer(display))

	C.gdk_app_launch_context_set_display(c_context, c_display)
}

func Fn_gdk_app_launch_context_set_icon(context unsafe.Pointer, icon unsafe.Pointer) {
	c_context := (*C.GdkAppLaunchContext)(unsafe.Pointer(context))

	c_icon := (*C.GIcon)(unsafe.Pointer(icon))

	C.gdk_app_launch_context_set_icon(c_context, c_icon)
}

func Fn_gdk_app_launch_context_set_icon_name(context unsafe.Pointer, iconName *string) {
	c_context := (*C.GdkAppLaunchContext)(unsafe.Pointer(context))

	var c_iconNameValue (*C.char)
	if iconName != nil {
		c_iconNameValue = (*C.char)(C.CString(*iconName))
		defer C.free(unsafe.Pointer(c_iconNameValue))
	}
	c_iconName := c_iconNameValue

	C.gdk_app_launch_context_set_icon_name(c_context, c_iconName)
}

func Fn_gdk_app_launch_context_set_screen(context unsafe.Pointer, screen unsafe.Pointer) {
	c_context := (*C.GdkAppLaunchContext)(unsafe.Pointer(context))

	c_screen := (*C.GdkScreen)(unsafe.Pointer(screen))

	C.gdk_app_launch_context_set_screen(c_context, c_screen)
}

func Fn_gdk_app_launch_context_set_timestamp(context unsafe.Pointer, timestamp uint32) {
	c_context := (*C.GdkAppLaunchContext)(unsafe.Pointer(context))

	c_timestamp := (C.guint32)(timestamp)

	C.gdk_app_launch_context_set_timestamp(c_context, c_timestamp)
}

func Fn_gdk_cursor_new(cursorType int) unsafe.Pointer {
	c_cursorType := (C.GdkCursorType)(cursorType)

	ret := C.gdk_cursor_new(c_cursorType)

	return unsafe.Pointer(ret)
}

func Fn_gdk_cursor_new_for_display(display unsafe.Pointer, cursorType int) unsafe.Pointer {
	c_display := (*C.GdkDisplay)(unsafe.Pointer(display))

	c_cursorType := (C.GdkCursorType)(cursorType)

	ret := C.gdk_cursor_new_for_display(c_display, c_cursorType)

	return unsafe.Pointer(ret)
}

func Fn_gdk_cursor_new_from_name(display unsafe.Pointer, name string) unsafe.Pointer {
	c_display := (*C.GdkDisplay)(unsafe.Pointer(display))

	c_name := (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(c_name))

	ret := C.gdk_cursor_new_from_name(c_display, c_name)

	return unsafe.Pointer(ret)
}

func Fn_gdk_cursor_new_from_pixbuf(display unsafe.Pointer, pixbuf unsafe.Pointer, x int, y int) unsafe.Pointer {
	c_display := (*C.GdkDisplay)(unsafe.Pointer(display))

	c_pixbuf := (*C.GdkPixbuf)(unsafe.Pointer(pixbuf))

	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	ret := C.gdk_cursor_new_from_pixbuf(c_display, c_pixbuf, c_x, c_y)

	return unsafe.Pointer(ret)
}

func Fn_gdk_cursor_get_display(cursor unsafe.Pointer) unsafe.Pointer {
	c_cursor := (*C.GdkCursor)(unsafe.Pointer(cursor))

	ret := C.gdk_cursor_get_display(c_cursor)

	return unsafe.Pointer(ret)
}

func Fn_gdk_cursor_get_image(cursor unsafe.Pointer) unsafe.Pointer {
	c_cursor := (*C.GdkCursor)(unsafe.Pointer(cursor))

	ret := C.gdk_cursor_get_image(c_cursor)

	return unsafe.Pointer(ret)
}

func Fn_gdk_cursor_ref(cursor unsafe.Pointer) unsafe.Pointer {
	c_cursor := (*C.GdkCursor)(unsafe.Pointer(cursor))

	ret := C.gdk_cursor_ref(c_cursor)

	return unsafe.Pointer(ret)
}

func Fn_gdk_cursor_unref(cursor unsafe.Pointer) {
	c_cursor := (*C.GdkCursor)(unsafe.Pointer(cursor))

	C.gdk_cursor_unref(c_cursor)
}

func Fn_gdk_device_get_axis(device unsafe.Pointer, axes []float64, use int, value *float64) bool {
	c_device := (*C.GdkDevice)(unsafe.Pointer(device))

	c_axes := (*C.gdouble)(unsafe.Pointer(&axes[0]))

	c_use := (C.GdkAxisUse)(use)

	c_value := (*C.gdouble)(unsafe.Pointer(value))

	ret := C.gdk_device_get_axis(c_device, c_axes, c_use, c_value)

	return toGoBool(ret)
}

func Fn_gdk_device_get_history(device unsafe.Pointer, window unsafe.Pointer, start uint32, stop uint32, events *[]unsafe.Pointer, nEvents *int) bool {
	c_device := (*C.GdkDevice)(unsafe.Pointer(device))

	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	c_start := (C.guint32)(start)

	c_stop := (C.guint32)(stop)

	var c_eventsArrayPointer (**C.GdkTimeCoord)
	c_events := &c_eventsArrayPointer

	c_nEvents := (*C.gint)(unsafe.Pointer(nEvents))

	ret := C.gdk_device_get_history(c_device, c_window, c_start, c_stop, c_events, c_nEvents)

	eventsOutLen := int(*c_nEvents)
	eventsOut := make([]unsafe.Pointer, eventsOutLen, eventsOutLen)
	if eventsOutLen > 0 {
		eventsOut = (*[1 << 30](unsafe.Pointer))(unsafe.Pointer(c_eventsArrayPointer))[:eventsOutLen:eventsOutLen]
	}
	*events = eventsOut

	return toGoBool(ret)
}

// UNSUPPORTED : gdk_device_get_position : parameter 'screen' is non array with indirect count > 1

// UNSUPPORTED : gdk_device_get_position_double : parameter 'screen' is non array with indirect count > 1

func Fn_gdk_device_get_state(device unsafe.Pointer, window unsafe.Pointer, axes []float64, mask *int) {
	c_device := (*C.GdkDevice)(unsafe.Pointer(device))

	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	c_axes := (*C.gdouble)(unsafe.Pointer(&axes[0]))

	c_mask := (*C.GdkModifierType)(unsafe.Pointer(mask))

	C.gdk_device_get_state(c_device, c_window, c_axes, c_mask)
}

func Fn_gdk_device_list_slave_devices(device unsafe.Pointer) unsafe.Pointer {
	c_device := (*C.GdkDevice)(unsafe.Pointer(device))

	ret := C.gdk_device_list_slave_devices(c_device)

	return unsafe.Pointer(ret)
}

func Fn_gdk_device_set_axis_use(device unsafe.Pointer, index uint, use int) {
	c_device := (*C.GdkDevice)(unsafe.Pointer(device))

	c_index := (C.guint)(index)

	c_use := (C.GdkAxisUse)(use)

	C.gdk_device_set_axis_use(c_device, c_index, c_use)
}

func Fn_gdk_device_set_key(device unsafe.Pointer, index uint, keyval uint, modifiers int) {
	c_device := (*C.GdkDevice)(unsafe.Pointer(device))

	c_index := (C.guint)(index)

	c_keyval := (C.guint)(keyval)

	c_modifiers := (C.GdkModifierType)(modifiers)

	C.gdk_device_set_key(c_device, c_index, c_keyval, c_modifiers)
}

func Fn_gdk_device_set_mode(device unsafe.Pointer, mode int) bool {
	c_device := (*C.GdkDevice)(unsafe.Pointer(device))

	c_mode := (C.GdkInputMode)(mode)

	ret := C.gdk_device_set_mode(c_device, c_mode)

	return toGoBool(ret)
}

func Fn_gdk_device_free_history(events []unsafe.Pointer, nEvents int) {
	c_events := (**C.GdkTimeCoord)(unsafe.Pointer(&events[0]))

	c_nEvents := (C.gint)(nEvents)

	C.gdk_device_free_history(c_events, c_nEvents)
}

// UNSUPPORTED : gdk_device_grab_info_libgtk_only : parameter 'grab_window' is non array with indirect count > 1

func Fn_gdk_display_beep(display unsafe.Pointer) {
	c_display := (*C.GdkDisplay)(unsafe.Pointer(display))

	C.gdk_display_beep(c_display)
}

func Fn_gdk_display_close(display unsafe.Pointer) {
	c_display := (*C.GdkDisplay)(unsafe.Pointer(display))

	C.gdk_display_close(c_display)
}

func Fn_gdk_display_device_is_grabbed(display unsafe.Pointer, device unsafe.Pointer) bool {
	c_display := (*C.GdkDisplay)(unsafe.Pointer(display))

	c_device := (*C.GdkDevice)(unsafe.Pointer(device))

	ret := C.gdk_display_device_is_grabbed(c_display, c_device)

	return toGoBool(ret)
}

func Fn_gdk_display_flush(display unsafe.Pointer) {
	c_display := (*C.GdkDisplay)(unsafe.Pointer(display))

	C.gdk_display_flush(c_display)
}

func Fn_gdk_display_get_default_cursor_size(display unsafe.Pointer) uint {
	c_display := (*C.GdkDisplay)(unsafe.Pointer(display))

	ret := C.gdk_display_get_default_cursor_size(c_display)

	return (uint)(ret)
}

func Fn_gdk_display_get_default_group(display unsafe.Pointer) unsafe.Pointer {
	c_display := (*C.GdkDisplay)(unsafe.Pointer(display))

	ret := C.gdk_display_get_default_group(c_display)

	return unsafe.Pointer(ret)
}

func Fn_gdk_display_get_default_screen(display unsafe.Pointer) unsafe.Pointer {
	c_display := (*C.GdkDisplay)(unsafe.Pointer(display))

	ret := C.gdk_display_get_default_screen(c_display)

	return unsafe.Pointer(ret)
}

func Fn_gdk_display_get_event(display unsafe.Pointer) unsafe.Pointer {
	c_display := (*C.GdkDisplay)(unsafe.Pointer(display))

	ret := C.gdk_display_get_event(c_display)

	return unsafe.Pointer(ret)
}

func Fn_gdk_display_get_maximal_cursor_size(display unsafe.Pointer, width *uint, height *uint) {
	c_display := (*C.GdkDisplay)(unsafe.Pointer(display))

	c_width := (*C.guint)(unsafe.Pointer(width))

	c_height := (*C.guint)(unsafe.Pointer(height))

	C.gdk_display_get_maximal_cursor_size(c_display, c_width, c_height)
}

func Fn_gdk_display_get_n_screens(display unsafe.Pointer) int {
	c_display := (*C.GdkDisplay)(unsafe.Pointer(display))

	ret := C.gdk_display_get_n_screens(c_display)

	return (int)(ret)
}

func Fn_gdk_display_get_name(display unsafe.Pointer) string {
	c_display := (*C.GdkDisplay)(unsafe.Pointer(display))

	ret := C.gdk_display_get_name(c_display)

	return C.GoString(ret)
}

// UNSUPPORTED : gdk_display_get_pointer : parameter 'screen' is non array with indirect count > 1

func Fn_gdk_display_get_screen(display unsafe.Pointer, screenNum int) unsafe.Pointer {
	c_display := (*C.GdkDisplay)(unsafe.Pointer(display))

	c_screenNum := (C.gint)(screenNum)

	ret := C.gdk_display_get_screen(c_display, c_screenNum)

	return unsafe.Pointer(ret)
}

func Fn_gdk_display_get_window_at_pointer(display unsafe.Pointer, winX *int, winY *int) unsafe.Pointer {
	c_display := (*C.GdkDisplay)(unsafe.Pointer(display))

	c_winX := (*C.gint)(unsafe.Pointer(winX))

	c_winY := (*C.gint)(unsafe.Pointer(winY))

	ret := C.gdk_display_get_window_at_pointer(c_display, c_winX, c_winY)

	return unsafe.Pointer(ret)
}

func Fn_gdk_display_keyboard_ungrab(display unsafe.Pointer, time uint32) {
	c_display := (*C.GdkDisplay)(unsafe.Pointer(display))

	c_time := (C.guint32)(time)

	C.gdk_display_keyboard_ungrab(c_display, c_time)
}

func Fn_gdk_display_list_devices(display unsafe.Pointer) unsafe.Pointer {
	c_display := (*C.GdkDisplay)(unsafe.Pointer(display))

	ret := C.gdk_display_list_devices(c_display)

	return unsafe.Pointer(ret)
}

func Fn_gdk_display_peek_event(display unsafe.Pointer) unsafe.Pointer {
	c_display := (*C.GdkDisplay)(unsafe.Pointer(display))

	ret := C.gdk_display_peek_event(c_display)

	return unsafe.Pointer(ret)
}

func Fn_gdk_display_pointer_is_grabbed(display unsafe.Pointer) bool {
	c_display := (*C.GdkDisplay)(unsafe.Pointer(display))

	ret := C.gdk_display_pointer_is_grabbed(c_display)

	return toGoBool(ret)
}

func Fn_gdk_display_pointer_ungrab(display unsafe.Pointer, time uint32) {
	c_display := (*C.GdkDisplay)(unsafe.Pointer(display))

	c_time := (C.guint32)(time)

	C.gdk_display_pointer_ungrab(c_display, c_time)
}

func Fn_gdk_display_put_event(display unsafe.Pointer, event unsafe.Pointer) {
	c_display := (*C.GdkDisplay)(unsafe.Pointer(display))

	c_event := (*C.GdkEvent)(unsafe.Pointer(event))

	C.gdk_display_put_event(c_display, c_event)
}

func Fn_gdk_display_request_selection_notification(display unsafe.Pointer, selection unsafe.Pointer) bool {
	c_display := (*C.GdkDisplay)(unsafe.Pointer(display))

	c_selection := (C.GdkAtom)(selection)

	ret := C.gdk_display_request_selection_notification(c_display, c_selection)

	return toGoBool(ret)
}

func Fn_gdk_display_set_double_click_distance(display unsafe.Pointer, distance uint) {
	c_display := (*C.GdkDisplay)(unsafe.Pointer(display))

	c_distance := (C.guint)(distance)

	C.gdk_display_set_double_click_distance(c_display, c_distance)
}

func Fn_gdk_display_set_double_click_time(display unsafe.Pointer, msec uint) {
	c_display := (*C.GdkDisplay)(unsafe.Pointer(display))

	c_msec := (C.guint)(msec)

	C.gdk_display_set_double_click_time(c_display, c_msec)
}

func Fn_gdk_display_store_clipboard(display unsafe.Pointer, clipboardWindow unsafe.Pointer, time uint32, targets []Atom, nTargets int) {
	c_display := (*C.GdkDisplay)(unsafe.Pointer(display))

	c_clipboardWindow := (*C.GdkWindow)(unsafe.Pointer(clipboardWindow))

	c_time := (C.guint32)(time)

	c_targets := (*C.GdkAtom)(unsafe.Pointer(&targets[0]))

	c_nTargets := (C.gint)(nTargets)

	C.gdk_display_store_clipboard(c_display, c_clipboardWindow, c_time, c_targets, c_nTargets)
}

func Fn_gdk_display_supports_clipboard_persistence(display unsafe.Pointer) bool {
	c_display := (*C.GdkDisplay)(unsafe.Pointer(display))

	ret := C.gdk_display_supports_clipboard_persistence(c_display)

	return toGoBool(ret)
}

func Fn_gdk_display_supports_composite(display unsafe.Pointer) bool {
	c_display := (*C.GdkDisplay)(unsafe.Pointer(display))

	ret := C.gdk_display_supports_composite(c_display)

	return toGoBool(ret)
}

func Fn_gdk_display_supports_cursor_alpha(display unsafe.Pointer) bool {
	c_display := (*C.GdkDisplay)(unsafe.Pointer(display))

	ret := C.gdk_display_supports_cursor_alpha(c_display)

	return toGoBool(ret)
}

func Fn_gdk_display_supports_cursor_color(display unsafe.Pointer) bool {
	c_display := (*C.GdkDisplay)(unsafe.Pointer(display))

	ret := C.gdk_display_supports_cursor_color(c_display)

	return toGoBool(ret)
}

func Fn_gdk_display_supports_input_shapes(display unsafe.Pointer) bool {
	c_display := (*C.GdkDisplay)(unsafe.Pointer(display))

	ret := C.gdk_display_supports_input_shapes(c_display)

	return toGoBool(ret)
}

func Fn_gdk_display_supports_selection_notification(display unsafe.Pointer) bool {
	c_display := (*C.GdkDisplay)(unsafe.Pointer(display))

	ret := C.gdk_display_supports_selection_notification(c_display)

	return toGoBool(ret)
}

func Fn_gdk_display_supports_shapes(display unsafe.Pointer) bool {
	c_display := (*C.GdkDisplay)(unsafe.Pointer(display))

	ret := C.gdk_display_supports_shapes(c_display)

	return toGoBool(ret)
}

func Fn_gdk_display_sync(display unsafe.Pointer) {
	c_display := (*C.GdkDisplay)(unsafe.Pointer(display))

	C.gdk_display_sync(c_display)
}

func Fn_gdk_display_warp_pointer(display unsafe.Pointer, screen unsafe.Pointer, x int, y int) {
	c_display := (*C.GdkDisplay)(unsafe.Pointer(display))

	c_screen := (*C.GdkScreen)(unsafe.Pointer(screen))

	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	C.gdk_display_warp_pointer(c_display, c_screen, c_x, c_y)
}

func Fn_gdk_display_get_default() unsafe.Pointer {
	ret := C.gdk_display_get_default()

	return unsafe.Pointer(ret)
}

func Fn_gdk_display_open(displayName string) unsafe.Pointer {
	c_displayName := (*C.gchar)(C.CString(displayName))
	defer C.free(unsafe.Pointer(c_displayName))

	ret := C.gdk_display_open(c_displayName)

	return unsafe.Pointer(ret)
}

func Fn_gdk_display_open_default_libgtk_only() unsafe.Pointer {
	ret := C.gdk_display_open_default_libgtk_only()

	return unsafe.Pointer(ret)
}

func Fn_gdk_display_manager_get_default_display(manager unsafe.Pointer) unsafe.Pointer {
	c_manager := (*C.GdkDisplayManager)(unsafe.Pointer(manager))

	ret := C.gdk_display_manager_get_default_display(c_manager)

	return unsafe.Pointer(ret)
}

func Fn_gdk_display_manager_list_displays(manager unsafe.Pointer) unsafe.Pointer {
	c_manager := (*C.GdkDisplayManager)(unsafe.Pointer(manager))

	ret := C.gdk_display_manager_list_displays(c_manager)

	return unsafe.Pointer(ret)
}

func Fn_gdk_display_manager_set_default_display(manager unsafe.Pointer, display unsafe.Pointer) {
	c_manager := (*C.GdkDisplayManager)(unsafe.Pointer(manager))

	c_display := (*C.GdkDisplay)(unsafe.Pointer(display))

	C.gdk_display_manager_set_default_display(c_manager, c_display)
}

func Fn_gdk_display_manager_get() unsafe.Pointer {
	ret := C.gdk_display_manager_get()

	return unsafe.Pointer(ret)
}

func Fn_gdk_drag_context_get_device(context unsafe.Pointer) unsafe.Pointer {
	c_context := (*C.GdkDragContext)(unsafe.Pointer(context))

	ret := C.gdk_drag_context_get_device(c_context)

	return unsafe.Pointer(ret)
}

func Fn_gdk_drag_context_set_device(context unsafe.Pointer, device unsafe.Pointer) {
	c_context := (*C.GdkDragContext)(unsafe.Pointer(context))

	c_device := (*C.GdkDevice)(unsafe.Pointer(device))

	C.gdk_drag_context_set_device(c_context, c_device)
}

func Fn_gdk_keymap_get_caps_lock_state(keymap unsafe.Pointer) bool {
	c_keymap := (*C.GdkKeymap)(unsafe.Pointer(keymap))

	ret := C.gdk_keymap_get_caps_lock_state(c_keymap)

	return toGoBool(ret)
}

func Fn_gdk_keymap_get_direction(keymap unsafe.Pointer) int {
	c_keymap := (*C.GdkKeymap)(unsafe.Pointer(keymap))

	ret := C.gdk_keymap_get_direction(c_keymap)

	return (int)(ret)
}

func Fn_gdk_keymap_get_entries_for_keycode(keymap unsafe.Pointer, hardwareKeycode uint, keys *[]unsafe.Pointer, keyvals *[]uint, nEntries *int) bool {
	c_keymap := (*C.GdkKeymap)(unsafe.Pointer(keymap))

	c_hardwareKeycode := (C.guint)(hardwareKeycode)

	var c_keysArrayPointer (*C.GdkKeymapKey)
	c_keys := &c_keysArrayPointer

	var c_keyvalsArrayPointer (*C.guint)
	c_keyvals := &c_keyvalsArrayPointer

	c_nEntries := (*C.gint)(unsafe.Pointer(nEntries))

	ret := C.gdk_keymap_get_entries_for_keycode(c_keymap, c_hardwareKeycode, c_keys, c_keyvals, c_nEntries)

	keysOutLen := int(*c_nEntries)
	keysOut := make([]unsafe.Pointer, keysOutLen, keysOutLen)
	if keysOutLen > 0 {
		keysOut = (*[1 << 30](unsafe.Pointer))(unsafe.Pointer(c_keysArrayPointer))[:keysOutLen:keysOutLen]
	}
	*keys = keysOut

	keyvalsOutLen := int(*c_nEntries)
	keyvalsOut := make([]uint, keyvalsOutLen, keyvalsOutLen)
	if keyvalsOutLen > 0 {
		keyvalsOut = (*[1 << 30](uint))(unsafe.Pointer(c_keyvalsArrayPointer))[:keyvalsOutLen:keyvalsOutLen]
	}
	*keyvals = keyvalsOut

	return toGoBool(ret)
}

func Fn_gdk_keymap_get_entries_for_keyval(keymap unsafe.Pointer, keyval uint, keys *[]unsafe.Pointer, nKeys *int) bool {
	c_keymap := (*C.GdkKeymap)(unsafe.Pointer(keymap))

	c_keyval := (C.guint)(keyval)

	var c_keysArrayPointer (*C.GdkKeymapKey)
	c_keys := &c_keysArrayPointer

	c_nKeys := (*C.gint)(unsafe.Pointer(nKeys))

	ret := C.gdk_keymap_get_entries_for_keyval(c_keymap, c_keyval, c_keys, c_nKeys)

	keysOutLen := int(*c_nKeys)
	keysOut := make([]unsafe.Pointer, keysOutLen, keysOutLen)
	if keysOutLen > 0 {
		keysOut = (*[1 << 30](unsafe.Pointer))(unsafe.Pointer(c_keysArrayPointer))[:keysOutLen:keysOutLen]
	}
	*keys = keysOut

	return toGoBool(ret)
}

func Fn_gdk_keymap_have_bidi_layouts(keymap unsafe.Pointer) bool {
	c_keymap := (*C.GdkKeymap)(unsafe.Pointer(keymap))

	ret := C.gdk_keymap_have_bidi_layouts(c_keymap)

	return toGoBool(ret)
}

func Fn_gdk_keymap_lookup_key(keymap unsafe.Pointer, key unsafe.Pointer) uint {
	c_keymap := (*C.GdkKeymap)(unsafe.Pointer(keymap))

	c_key := (*C.GdkKeymapKey)(unsafe.Pointer(key))

	ret := C.gdk_keymap_lookup_key(c_keymap, c_key)

	return (uint)(ret)
}

func Fn_gdk_keymap_translate_keyboard_state(keymap unsafe.Pointer, hardwareKeycode uint, state int, group int, keyval *uint, effectiveGroup *int, level *int, consumedModifiers *int) bool {
	c_keymap := (*C.GdkKeymap)(unsafe.Pointer(keymap))

	c_hardwareKeycode := (C.guint)(hardwareKeycode)

	c_state := (C.GdkModifierType)(state)

	c_group := (C.gint)(group)

	c_keyval := (*C.guint)(unsafe.Pointer(keyval))

	c_effectiveGroup := (*C.gint)(unsafe.Pointer(effectiveGroup))

	c_level := (*C.gint)(unsafe.Pointer(level))

	c_consumedModifiers := (*C.GdkModifierType)(unsafe.Pointer(consumedModifiers))

	ret := C.gdk_keymap_translate_keyboard_state(c_keymap, c_hardwareKeycode, c_state, c_group, c_keyval, c_effectiveGroup, c_level, c_consumedModifiers)

	return toGoBool(ret)
}

func Fn_gdk_keymap_get_default() unsafe.Pointer {
	ret := C.gdk_keymap_get_default()

	return unsafe.Pointer(ret)
}

func Fn_gdk_keymap_get_for_display(display unsafe.Pointer) unsafe.Pointer {
	c_display := (*C.GdkDisplay)(unsafe.Pointer(display))

	ret := C.gdk_keymap_get_for_display(c_display)

	return unsafe.Pointer(ret)
}

func Fn_gdk_monitor_get_manufacturer(monitor unsafe.Pointer) string {
	c_monitor := (*C.GdkMonitor)(unsafe.Pointer(monitor))

	ret := C.gdk_monitor_get_manufacturer(c_monitor)

	return C.GoString(ret)
}

func Fn_gdk_monitor_get_model(monitor unsafe.Pointer) string {
	c_monitor := (*C.GdkMonitor)(unsafe.Pointer(monitor))

	ret := C.gdk_monitor_get_model(c_monitor)

	return C.GoString(ret)
}

func Fn_gdk_screen_get_active_window(screen unsafe.Pointer) unsafe.Pointer {
	c_screen := (*C.GdkScreen)(unsafe.Pointer(screen))

	ret := C.gdk_screen_get_active_window(c_screen)

	return unsafe.Pointer(ret)
}

func Fn_gdk_screen_get_display(screen unsafe.Pointer) unsafe.Pointer {
	c_screen := (*C.GdkScreen)(unsafe.Pointer(screen))

	ret := C.gdk_screen_get_display(c_screen)

	return unsafe.Pointer(ret)
}

func Fn_gdk_screen_get_font_options(screen unsafe.Pointer) unsafe.Pointer {
	c_screen := (*C.GdkScreen)(unsafe.Pointer(screen))

	ret := C.gdk_screen_get_font_options(c_screen)

	return unsafe.Pointer(ret)
}

func Fn_gdk_screen_get_height(screen unsafe.Pointer) int {
	c_screen := (*C.GdkScreen)(unsafe.Pointer(screen))

	ret := C.gdk_screen_get_height(c_screen)

	return (int)(ret)
}

func Fn_gdk_screen_get_height_mm(screen unsafe.Pointer) int {
	c_screen := (*C.GdkScreen)(unsafe.Pointer(screen))

	ret := C.gdk_screen_get_height_mm(c_screen)

	return (int)(ret)
}

func Fn_gdk_screen_get_monitor_at_point(screen unsafe.Pointer, x int, y int) int {
	c_screen := (*C.GdkScreen)(unsafe.Pointer(screen))

	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	ret := C.gdk_screen_get_monitor_at_point(c_screen, c_x, c_y)

	return (int)(ret)
}

func Fn_gdk_screen_get_monitor_at_window(screen unsafe.Pointer, window unsafe.Pointer) int {
	c_screen := (*C.GdkScreen)(unsafe.Pointer(screen))

	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	ret := C.gdk_screen_get_monitor_at_window(c_screen, c_window)

	return (int)(ret)
}

func Fn_gdk_screen_get_monitor_geometry(screen unsafe.Pointer, monitorNum int, dest unsafe.Pointer) {
	c_screen := (*C.GdkScreen)(unsafe.Pointer(screen))

	c_monitorNum := (C.gint)(monitorNum)

	c_dest := (*C.GdkRectangle)(unsafe.Pointer(dest))

	C.gdk_screen_get_monitor_geometry(c_screen, c_monitorNum, c_dest)
}

func Fn_gdk_screen_get_monitor_height_mm(screen unsafe.Pointer, monitorNum int) int {
	c_screen := (*C.GdkScreen)(unsafe.Pointer(screen))

	c_monitorNum := (C.gint)(monitorNum)

	ret := C.gdk_screen_get_monitor_height_mm(c_screen, c_monitorNum)

	return (int)(ret)
}

func Fn_gdk_screen_get_monitor_plug_name(screen unsafe.Pointer, monitorNum int) string {
	c_screen := (*C.GdkScreen)(unsafe.Pointer(screen))

	c_monitorNum := (C.gint)(monitorNum)

	ret := C.gdk_screen_get_monitor_plug_name(c_screen, c_monitorNum)

	return C.GoString(ret)
}

func Fn_gdk_screen_get_monitor_width_mm(screen unsafe.Pointer, monitorNum int) int {
	c_screen := (*C.GdkScreen)(unsafe.Pointer(screen))

	c_monitorNum := (C.gint)(monitorNum)

	ret := C.gdk_screen_get_monitor_width_mm(c_screen, c_monitorNum)

	return (int)(ret)
}

func Fn_gdk_screen_get_n_monitors(screen unsafe.Pointer) int {
	c_screen := (*C.GdkScreen)(unsafe.Pointer(screen))

	ret := C.gdk_screen_get_n_monitors(c_screen)

	return (int)(ret)
}

func Fn_gdk_screen_get_number(screen unsafe.Pointer) int {
	c_screen := (*C.GdkScreen)(unsafe.Pointer(screen))

	ret := C.gdk_screen_get_number(c_screen)

	return (int)(ret)
}

func Fn_gdk_screen_get_resolution(screen unsafe.Pointer) float64 {
	c_screen := (*C.GdkScreen)(unsafe.Pointer(screen))

	ret := C.gdk_screen_get_resolution(c_screen)

	return (float64)(ret)
}

func Fn_gdk_screen_get_rgba_visual(screen unsafe.Pointer) unsafe.Pointer {
	c_screen := (*C.GdkScreen)(unsafe.Pointer(screen))

	ret := C.gdk_screen_get_rgba_visual(c_screen)

	return unsafe.Pointer(ret)
}

func Fn_gdk_screen_get_root_window(screen unsafe.Pointer) unsafe.Pointer {
	c_screen := (*C.GdkScreen)(unsafe.Pointer(screen))

	ret := C.gdk_screen_get_root_window(c_screen)

	return unsafe.Pointer(ret)
}

func Fn_gdk_screen_get_setting(screen unsafe.Pointer, name string, value unsafe.Pointer) bool {
	c_screen := (*C.GdkScreen)(unsafe.Pointer(screen))

	c_name := (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(c_name))

	c_value := (*C.GValue)(unsafe.Pointer(value))

	ret := C.gdk_screen_get_setting(c_screen, c_name, c_value)

	return toGoBool(ret)
}

func Fn_gdk_screen_get_system_visual(screen unsafe.Pointer) unsafe.Pointer {
	c_screen := (*C.GdkScreen)(unsafe.Pointer(screen))

	ret := C.gdk_screen_get_system_visual(c_screen)

	return unsafe.Pointer(ret)
}

func Fn_gdk_screen_get_toplevel_windows(screen unsafe.Pointer) unsafe.Pointer {
	c_screen := (*C.GdkScreen)(unsafe.Pointer(screen))

	ret := C.gdk_screen_get_toplevel_windows(c_screen)

	return unsafe.Pointer(ret)
}

func Fn_gdk_screen_get_width(screen unsafe.Pointer) int {
	c_screen := (*C.GdkScreen)(unsafe.Pointer(screen))

	ret := C.gdk_screen_get_width(c_screen)

	return (int)(ret)
}

func Fn_gdk_screen_get_width_mm(screen unsafe.Pointer) int {
	c_screen := (*C.GdkScreen)(unsafe.Pointer(screen))

	ret := C.gdk_screen_get_width_mm(c_screen)

	return (int)(ret)
}

func Fn_gdk_screen_get_window_stack(screen unsafe.Pointer) unsafe.Pointer {
	c_screen := (*C.GdkScreen)(unsafe.Pointer(screen))

	ret := C.gdk_screen_get_window_stack(c_screen)

	return unsafe.Pointer(ret)
}

func Fn_gdk_screen_is_composited(screen unsafe.Pointer) bool {
	c_screen := (*C.GdkScreen)(unsafe.Pointer(screen))

	ret := C.gdk_screen_is_composited(c_screen)

	return toGoBool(ret)
}

func Fn_gdk_screen_list_visuals(screen unsafe.Pointer) unsafe.Pointer {
	c_screen := (*C.GdkScreen)(unsafe.Pointer(screen))

	ret := C.gdk_screen_list_visuals(c_screen)

	return unsafe.Pointer(ret)
}

func Fn_gdk_screen_make_display_name(screen unsafe.Pointer) string {
	c_screen := (*C.GdkScreen)(unsafe.Pointer(screen))

	ret := C.gdk_screen_make_display_name(c_screen)

	return C.GoString(ret)
}

func Fn_gdk_screen_set_font_options(screen unsafe.Pointer, options unsafe.Pointer) {
	c_screen := (*C.GdkScreen)(unsafe.Pointer(screen))

	c_options := (*C.cairo_font_options_t)(unsafe.Pointer(options))

	C.gdk_screen_set_font_options(c_screen, c_options)
}

func Fn_gdk_screen_set_resolution(screen unsafe.Pointer, dpi float64) {
	c_screen := (*C.GdkScreen)(unsafe.Pointer(screen))

	c_dpi := (C.gdouble)(dpi)

	C.gdk_screen_set_resolution(c_screen, c_dpi)
}

func Fn_gdk_screen_get_default() unsafe.Pointer {
	ret := C.gdk_screen_get_default()

	return unsafe.Pointer(ret)
}

func Fn_gdk_screen_height() int {
	ret := C.gdk_screen_height()

	return (int)(ret)
}

func Fn_gdk_screen_height_mm() int {
	ret := C.gdk_screen_height_mm()

	return (int)(ret)
}

func Fn_gdk_screen_width() int {
	ret := C.gdk_screen_width()

	return (int)(ret)
}

func Fn_gdk_screen_width_mm() int {
	ret := C.gdk_screen_width_mm()

	return (int)(ret)
}

func Fn_gdk_seat_get_display(seat unsafe.Pointer) unsafe.Pointer {
	c_seat := (*C.GdkSeat)(unsafe.Pointer(seat))

	ret := C.gdk_seat_get_display(c_seat)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gdk_seat_grab : parameter 'prepare_func' is callback

func Fn_gdk_visual_get_screen(visual unsafe.Pointer) unsafe.Pointer {
	c_visual := (*C.GdkVisual)(unsafe.Pointer(visual))

	ret := C.gdk_visual_get_screen(c_visual)

	return unsafe.Pointer(ret)
}

func Fn_gdk_visual_get_best() unsafe.Pointer {
	ret := C.gdk_visual_get_best()

	return unsafe.Pointer(ret)
}

func Fn_gdk_visual_get_best_depth() int {
	ret := C.gdk_visual_get_best_depth()

	return (int)(ret)
}

func Fn_gdk_visual_get_best_type() int {
	ret := C.gdk_visual_get_best_type()

	return (int)(ret)
}

func Fn_gdk_visual_get_best_with_both(depth int, visualType int) unsafe.Pointer {
	c_depth := (C.gint)(depth)

	c_visualType := (C.GdkVisualType)(visualType)

	ret := C.gdk_visual_get_best_with_both(c_depth, c_visualType)

	return unsafe.Pointer(ret)
}

func Fn_gdk_visual_get_best_with_depth(depth int) unsafe.Pointer {
	c_depth := (C.gint)(depth)

	ret := C.gdk_visual_get_best_with_depth(c_depth)

	return unsafe.Pointer(ret)
}

func Fn_gdk_visual_get_best_with_type(visualType int) unsafe.Pointer {
	c_visualType := (C.GdkVisualType)(visualType)

	ret := C.gdk_visual_get_best_with_type(c_visualType)

	return unsafe.Pointer(ret)
}

func Fn_gdk_visual_get_system() unsafe.Pointer {
	ret := C.gdk_visual_get_system()

	return unsafe.Pointer(ret)
}

func Fn_gdk_window_new(parent unsafe.Pointer, attributes unsafe.Pointer, attributesMask int) unsafe.Pointer {
	c_parent := (*C.GdkWindow)(unsafe.Pointer(parent))

	c_attributes := (*C.GdkWindowAttr)(unsafe.Pointer(attributes))

	c_attributesMask := (C.gint)(attributesMask)

	ret := C.gdk_window_new(c_parent, c_attributes, c_attributesMask)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gdk_window_add_filter : parameter 'function' is callback

func Fn_gdk_window_beep(window unsafe.Pointer) {
	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	C.gdk_window_beep(c_window)
}

func Fn_gdk_window_begin_move_drag(window unsafe.Pointer, button int, rootX int, rootY int, timestamp uint32) {
	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	c_button := (C.gint)(button)

	c_rootX := (C.gint)(rootX)

	c_rootY := (C.gint)(rootY)

	c_timestamp := (C.guint32)(timestamp)

	C.gdk_window_begin_move_drag(c_window, c_button, c_rootX, c_rootY, c_timestamp)
}

func Fn_gdk_window_begin_paint_rect(window unsafe.Pointer, rectangle unsafe.Pointer) {
	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	c_rectangle := (*C.GdkRectangle)(unsafe.Pointer(rectangle))

	C.gdk_window_begin_paint_rect(c_window, c_rectangle)
}

func Fn_gdk_window_begin_paint_region(window unsafe.Pointer, region unsafe.Pointer) {
	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	c_region := (*C.cairo_region_t)(unsafe.Pointer(region))

	C.gdk_window_begin_paint_region(c_window, c_region)
}

func Fn_gdk_window_begin_resize_drag(window unsafe.Pointer, edge int, button int, rootX int, rootY int, timestamp uint32) {
	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	c_edge := (C.GdkWindowEdge)(edge)

	c_button := (C.gint)(button)

	c_rootX := (C.gint)(rootX)

	c_rootY := (C.gint)(rootY)

	c_timestamp := (C.guint32)(timestamp)

	C.gdk_window_begin_resize_drag(c_window, c_edge, c_button, c_rootX, c_rootY, c_timestamp)
}

func Fn_gdk_window_configure_finished(window unsafe.Pointer) {
	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	C.gdk_window_configure_finished(c_window)
}

func Fn_gdk_window_deiconify(window unsafe.Pointer) {
	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	C.gdk_window_deiconify(c_window)
}

func Fn_gdk_window_destroy(window unsafe.Pointer) {
	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	C.gdk_window_destroy(c_window)
}

// UNSUPPORTED : gdk_window_destroy_notify : blacklisted

func Fn_gdk_window_enable_synchronized_configure(window unsafe.Pointer) {
	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	C.gdk_window_enable_synchronized_configure(c_window)
}

func Fn_gdk_window_end_paint(window unsafe.Pointer) {
	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	C.gdk_window_end_paint(c_window)
}

func Fn_gdk_window_ensure_native(window unsafe.Pointer) bool {
	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	ret := C.gdk_window_ensure_native(c_window)

	return toGoBool(ret)
}

func Fn_gdk_window_flush(window unsafe.Pointer) {
	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	C.gdk_window_flush(c_window)
}

func Fn_gdk_window_focus(window unsafe.Pointer, timestamp uint32) {
	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	c_timestamp := (C.guint32)(timestamp)

	C.gdk_window_focus(c_window, c_timestamp)
}

func Fn_gdk_window_freeze_toplevel_updates_libgtk_only(window unsafe.Pointer) {
	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	C.gdk_window_freeze_toplevel_updates_libgtk_only(c_window)
}

func Fn_gdk_window_freeze_updates(window unsafe.Pointer) {
	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	C.gdk_window_freeze_updates(c_window)
}

func Fn_gdk_window_fullscreen(window unsafe.Pointer) {
	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	C.gdk_window_fullscreen(c_window)
}

func Fn_gdk_window_fullscreen_on_monitor(window unsafe.Pointer, monitor int) {
	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	c_monitor := (C.gint)(monitor)

	C.gdk_window_fullscreen_on_monitor(c_window, c_monitor)
}

func Fn_gdk_window_geometry_changed(window unsafe.Pointer) {
	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	C.gdk_window_geometry_changed(c_window)
}

func Fn_gdk_window_get_children(window unsafe.Pointer) unsafe.Pointer {
	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	ret := C.gdk_window_get_children(c_window)

	return unsafe.Pointer(ret)
}

func Fn_gdk_window_get_clip_region(window unsafe.Pointer) unsafe.Pointer {
	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	ret := C.gdk_window_get_clip_region(c_window)

	return unsafe.Pointer(ret)
}

func Fn_gdk_window_get_cursor(window unsafe.Pointer) unsafe.Pointer {
	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	ret := C.gdk_window_get_cursor(c_window)

	return unsafe.Pointer(ret)
}

func Fn_gdk_window_get_decorations(window unsafe.Pointer, decorations *int) bool {
	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	c_decorations := (*C.GdkWMDecoration)(unsafe.Pointer(decorations))

	ret := C.gdk_window_get_decorations(c_window, c_decorations)

	return toGoBool(ret)
}

// UNSUPPORTED : gdk_window_get_drag_protocol : parameter 'target' is non array with indirect count > 1

func Fn_gdk_window_get_events(window unsafe.Pointer) int {
	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	ret := C.gdk_window_get_events(c_window)

	return (int)(ret)
}

func Fn_gdk_window_get_frame_extents(window unsafe.Pointer, rect unsafe.Pointer) {
	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	c_rect := (*C.GdkRectangle)(unsafe.Pointer(rect))

	C.gdk_window_get_frame_extents(c_window, c_rect)
}

func Fn_gdk_window_get_geometry(window unsafe.Pointer, x *int, y *int, width *int, height *int) {
	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	c_x := (*C.gint)(unsafe.Pointer(x))

	c_y := (*C.gint)(unsafe.Pointer(y))

	c_width := (*C.gint)(unsafe.Pointer(width))

	c_height := (*C.gint)(unsafe.Pointer(height))

	C.gdk_window_get_geometry(c_window, c_x, c_y, c_width, c_height)
}

func Fn_gdk_window_get_group(window unsafe.Pointer) unsafe.Pointer {
	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	ret := C.gdk_window_get_group(c_window)

	return unsafe.Pointer(ret)
}

func Fn_gdk_window_get_origin(window unsafe.Pointer, x *int, y *int) int {
	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	c_x := (*C.gint)(unsafe.Pointer(x))

	c_y := (*C.gint)(unsafe.Pointer(y))

	ret := C.gdk_window_get_origin(c_window, c_x, c_y)

	return (int)(ret)
}

func Fn_gdk_window_get_parent(window unsafe.Pointer) unsafe.Pointer {
	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	ret := C.gdk_window_get_parent(c_window)

	return unsafe.Pointer(ret)
}

func Fn_gdk_window_get_pointer(window unsafe.Pointer, x *int, y *int, mask *int) unsafe.Pointer {
	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	c_x := (*C.gint)(unsafe.Pointer(x))

	c_y := (*C.gint)(unsafe.Pointer(y))

	c_mask := (*C.GdkModifierType)(unsafe.Pointer(mask))

	ret := C.gdk_window_get_pointer(c_window, c_x, c_y, c_mask)

	return unsafe.Pointer(ret)
}

func Fn_gdk_window_get_position(window unsafe.Pointer, x *int, y *int) {
	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	c_x := (*C.gint)(unsafe.Pointer(x))

	c_y := (*C.gint)(unsafe.Pointer(y))

	C.gdk_window_get_position(c_window, c_x, c_y)
}

func Fn_gdk_window_get_root_coords(window unsafe.Pointer, x int, y int, rootX *int, rootY *int) {
	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	c_rootX := (*C.gint)(unsafe.Pointer(rootX))

	c_rootY := (*C.gint)(unsafe.Pointer(rootY))

	C.gdk_window_get_root_coords(c_window, c_x, c_y, c_rootX, c_rootY)
}

func Fn_gdk_window_get_root_origin(window unsafe.Pointer, x *int, y *int) {
	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	c_x := (*C.gint)(unsafe.Pointer(x))

	c_y := (*C.gint)(unsafe.Pointer(y))

	C.gdk_window_get_root_origin(c_window, c_x, c_y)
}

func Fn_gdk_window_get_source_events(window unsafe.Pointer, source int) int {
	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	c_source := (C.GdkInputSource)(source)

	ret := C.gdk_window_get_source_events(c_window, c_source)

	return (int)(ret)
}

func Fn_gdk_window_get_state(window unsafe.Pointer) int {
	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	ret := C.gdk_window_get_state(c_window)

	return (int)(ret)
}

func Fn_gdk_window_get_toplevel(window unsafe.Pointer) unsafe.Pointer {
	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	ret := C.gdk_window_get_toplevel(c_window)

	return unsafe.Pointer(ret)
}

func Fn_gdk_window_get_type_hint(window unsafe.Pointer) int {
	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	ret := C.gdk_window_get_type_hint(c_window)

	return (int)(ret)
}

func Fn_gdk_window_get_update_area(window unsafe.Pointer) unsafe.Pointer {
	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	ret := C.gdk_window_get_update_area(c_window)

	return unsafe.Pointer(ret)
}

func Fn_gdk_window_get_user_data(window unsafe.Pointer, data *unsafe.Pointer) {
	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	c_data := (*C.gpointer)(unsafe.Pointer(data))

	C.gdk_window_get_user_data(c_window, c_data)
}

func Fn_gdk_window_get_visible_region(window unsafe.Pointer) unsafe.Pointer {
	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	ret := C.gdk_window_get_visible_region(c_window)

	return unsafe.Pointer(ret)
}

func Fn_gdk_window_get_window_type(window unsafe.Pointer) int {
	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	ret := C.gdk_window_get_window_type(c_window)

	return (int)(ret)
}

func Fn_gdk_window_hide(window unsafe.Pointer) {
	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	C.gdk_window_hide(c_window)
}

func Fn_gdk_window_iconify(window unsafe.Pointer) {
	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	C.gdk_window_iconify(c_window)
}

func Fn_gdk_window_input_shape_combine_region(window unsafe.Pointer, shapeRegion unsafe.Pointer, offsetX int, offsetY int) {
	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	c_shapeRegion := (*C.cairo_region_t)(unsafe.Pointer(shapeRegion))

	c_offsetX := (C.gint)(offsetX)

	c_offsetY := (C.gint)(offsetY)

	C.gdk_window_input_shape_combine_region(c_window, c_shapeRegion, c_offsetX, c_offsetY)
}

// UNSUPPORTED : gdk_window_invalidate_maybe_recurse : parameter 'child_func' is callback

func Fn_gdk_window_invalidate_rect(window unsafe.Pointer, rect unsafe.Pointer, invalidateChildren bool) {
	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	c_rect := (*C.GdkRectangle)(unsafe.Pointer(rect))

	c_invalidateChildren := toCBool(invalidateChildren)

	C.gdk_window_invalidate_rect(c_window, c_rect, c_invalidateChildren)
}

func Fn_gdk_window_invalidate_region(window unsafe.Pointer, region unsafe.Pointer, invalidateChildren bool) {
	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	c_region := (*C.cairo_region_t)(unsafe.Pointer(region))

	c_invalidateChildren := toCBool(invalidateChildren)

	C.gdk_window_invalidate_region(c_window, c_region, c_invalidateChildren)
}

func Fn_gdk_window_is_destroyed(window unsafe.Pointer) bool {
	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	ret := C.gdk_window_is_destroyed(c_window)

	return toGoBool(ret)
}

func Fn_gdk_window_is_viewable(window unsafe.Pointer) bool {
	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	ret := C.gdk_window_is_viewable(c_window)

	return toGoBool(ret)
}

func Fn_gdk_window_is_visible(window unsafe.Pointer) bool {
	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	ret := C.gdk_window_is_visible(c_window)

	return toGoBool(ret)
}

func Fn_gdk_window_lower(window unsafe.Pointer) {
	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	C.gdk_window_lower(c_window)
}

func Fn_gdk_window_maximize(window unsafe.Pointer) {
	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	C.gdk_window_maximize(c_window)
}

func Fn_gdk_window_merge_child_input_shapes(window unsafe.Pointer) {
	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	C.gdk_window_merge_child_input_shapes(c_window)
}

func Fn_gdk_window_merge_child_shapes(window unsafe.Pointer) {
	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	C.gdk_window_merge_child_shapes(c_window)
}

func Fn_gdk_window_move(window unsafe.Pointer, x int, y int) {
	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	C.gdk_window_move(c_window, c_x, c_y)
}

func Fn_gdk_window_move_region(window unsafe.Pointer, region unsafe.Pointer, dx int, dy int) {
	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	c_region := (*C.cairo_region_t)(unsafe.Pointer(region))

	c_dx := (C.gint)(dx)

	c_dy := (C.gint)(dy)

	C.gdk_window_move_region(c_window, c_region, c_dx, c_dy)
}

func Fn_gdk_window_move_resize(window unsafe.Pointer, x int, y int, width int, height int) {
	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	c_width := (C.gint)(width)

	c_height := (C.gint)(height)

	C.gdk_window_move_resize(c_window, c_x, c_y, c_width, c_height)
}

func Fn_gdk_window_peek_children(window unsafe.Pointer) unsafe.Pointer {
	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	ret := C.gdk_window_peek_children(c_window)

	return unsafe.Pointer(ret)
}

func Fn_gdk_window_process_updates(window unsafe.Pointer, updateChildren bool) {
	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	c_updateChildren := toCBool(updateChildren)

	C.gdk_window_process_updates(c_window, c_updateChildren)
}

func Fn_gdk_window_raise(window unsafe.Pointer) {
	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	C.gdk_window_raise(c_window)
}

func Fn_gdk_window_register_dnd(window unsafe.Pointer) {
	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	C.gdk_window_register_dnd(c_window)
}

// UNSUPPORTED : gdk_window_remove_filter : parameter 'function' is callback

func Fn_gdk_window_reparent(window unsafe.Pointer, newParent unsafe.Pointer, x int, y int) {
	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	c_newParent := (*C.GdkWindow)(unsafe.Pointer(newParent))

	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	C.gdk_window_reparent(c_window, c_newParent, c_x, c_y)
}

func Fn_gdk_window_resize(window unsafe.Pointer, width int, height int) {
	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	c_width := (C.gint)(width)

	c_height := (C.gint)(height)

	C.gdk_window_resize(c_window, c_width, c_height)
}

func Fn_gdk_window_restack(window unsafe.Pointer, sibling unsafe.Pointer, above bool) {
	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	c_sibling := (*C.GdkWindow)(unsafe.Pointer(sibling))

	c_above := toCBool(above)

	C.gdk_window_restack(c_window, c_sibling, c_above)
}

func Fn_gdk_window_scroll(window unsafe.Pointer, dx int, dy int) {
	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	c_dx := (C.gint)(dx)

	c_dy := (C.gint)(dy)

	C.gdk_window_scroll(c_window, c_dx, c_dy)
}

func Fn_gdk_window_set_accept_focus(window unsafe.Pointer, acceptFocus bool) {
	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	c_acceptFocus := toCBool(acceptFocus)

	C.gdk_window_set_accept_focus(c_window, c_acceptFocus)
}

func Fn_gdk_window_set_background(window unsafe.Pointer, color unsafe.Pointer) {
	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	c_color := (*C.GdkColor)(unsafe.Pointer(color))

	C.gdk_window_set_background(c_window, c_color)
}

func Fn_gdk_window_set_background_pattern(window unsafe.Pointer, pattern unsafe.Pointer) {
	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	c_pattern := (*C.cairo_pattern_t)(unsafe.Pointer(pattern))

	C.gdk_window_set_background_pattern(c_window, c_pattern)
}

func Fn_gdk_window_set_background_rgba(window unsafe.Pointer, rgba unsafe.Pointer) {
	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	c_rgba := (*C.GdkRGBA)(unsafe.Pointer(rgba))

	C.gdk_window_set_background_rgba(c_window, c_rgba)
}

func Fn_gdk_window_set_child_input_shapes(window unsafe.Pointer) {
	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	C.gdk_window_set_child_input_shapes(c_window)
}

func Fn_gdk_window_set_child_shapes(window unsafe.Pointer) {
	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	C.gdk_window_set_child_shapes(c_window)
}

func Fn_gdk_window_set_composited(window unsafe.Pointer, composited bool) {
	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	c_composited := toCBool(composited)

	C.gdk_window_set_composited(c_window, c_composited)
}

func Fn_gdk_window_set_cursor(window unsafe.Pointer, cursor unsafe.Pointer) {
	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	c_cursor := (*C.GdkCursor)(unsafe.Pointer(cursor))

	C.gdk_window_set_cursor(c_window, c_cursor)
}

func Fn_gdk_window_set_decorations(window unsafe.Pointer, decorations int) {
	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	c_decorations := (C.GdkWMDecoration)(decorations)

	C.gdk_window_set_decorations(c_window, c_decorations)
}

func Fn_gdk_window_set_events(window unsafe.Pointer, eventMask int) {
	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	c_eventMask := (C.GdkEventMask)(eventMask)

	C.gdk_window_set_events(c_window, c_eventMask)
}

func Fn_gdk_window_set_focus_on_map(window unsafe.Pointer, focusOnMap bool) {
	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	c_focusOnMap := toCBool(focusOnMap)

	C.gdk_window_set_focus_on_map(c_window, c_focusOnMap)
}

func Fn_gdk_window_set_functions(window unsafe.Pointer, functions int) {
	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	c_functions := (C.GdkWMFunction)(functions)

	C.gdk_window_set_functions(c_window, c_functions)
}

func Fn_gdk_window_set_geometry_hints(window unsafe.Pointer, geometry unsafe.Pointer, geomMask int) {
	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	c_geometry := (*C.GdkGeometry)(unsafe.Pointer(geometry))

	c_geomMask := (C.GdkWindowHints)(geomMask)

	C.gdk_window_set_geometry_hints(c_window, c_geometry, c_geomMask)
}

func Fn_gdk_window_set_group(window unsafe.Pointer, leader unsafe.Pointer) {
	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	c_leader := (*C.GdkWindow)(unsafe.Pointer(leader))

	C.gdk_window_set_group(c_window, c_leader)
}

func Fn_gdk_window_set_icon_list(window unsafe.Pointer, pixbufs unsafe.Pointer) {
	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	c_pixbufs := (*C.GList)(unsafe.Pointer(pixbufs))

	C.gdk_window_set_icon_list(c_window, c_pixbufs)
}

func Fn_gdk_window_set_icon_name(window unsafe.Pointer, name *string) {
	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	var c_nameValue (*C.gchar)
	if name != nil {
		c_nameValue = (*C.gchar)(C.CString(*name))
		defer C.free(unsafe.Pointer(c_nameValue))
	}
	c_name := c_nameValue

	C.gdk_window_set_icon_name(c_window, c_name)
}

// UNSUPPORTED : gdk_window_set_invalidate_handler : parameter 'handler' is callback

func Fn_gdk_window_set_keep_above(window unsafe.Pointer, setting bool) {
	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	c_setting := toCBool(setting)

	C.gdk_window_set_keep_above(c_window, c_setting)
}

func Fn_gdk_window_set_keep_below(window unsafe.Pointer, setting bool) {
	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	c_setting := toCBool(setting)

	C.gdk_window_set_keep_below(c_window, c_setting)
}

func Fn_gdk_window_set_modal_hint(window unsafe.Pointer, modal bool) {
	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	c_modal := toCBool(modal)

	C.gdk_window_set_modal_hint(c_window, c_modal)
}

func Fn_gdk_window_set_opacity(window unsafe.Pointer, opacity float64) {
	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	c_opacity := (C.gdouble)(opacity)

	C.gdk_window_set_opacity(c_window, c_opacity)
}

func Fn_gdk_window_set_override_redirect(window unsafe.Pointer, overrideRedirect bool) {
	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	c_overrideRedirect := toCBool(overrideRedirect)

	C.gdk_window_set_override_redirect(c_window, c_overrideRedirect)
}

func Fn_gdk_window_set_role(window unsafe.Pointer, role string) {
	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	c_role := (*C.gchar)(C.CString(role))
	defer C.free(unsafe.Pointer(c_role))

	C.gdk_window_set_role(c_window, c_role)
}

func Fn_gdk_window_set_skip_pager_hint(window unsafe.Pointer, skipsPager bool) {
	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	c_skipsPager := toCBool(skipsPager)

	C.gdk_window_set_skip_pager_hint(c_window, c_skipsPager)
}

func Fn_gdk_window_set_skip_taskbar_hint(window unsafe.Pointer, skipsTaskbar bool) {
	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	c_skipsTaskbar := toCBool(skipsTaskbar)

	C.gdk_window_set_skip_taskbar_hint(c_window, c_skipsTaskbar)
}

func Fn_gdk_window_set_startup_id(window unsafe.Pointer, startupId string) {
	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	c_startupId := (*C.gchar)(C.CString(startupId))
	defer C.free(unsafe.Pointer(c_startupId))

	C.gdk_window_set_startup_id(c_window, c_startupId)
}

func Fn_gdk_window_set_static_gravities(window unsafe.Pointer, useStatic bool) bool {
	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	c_useStatic := toCBool(useStatic)

	ret := C.gdk_window_set_static_gravities(c_window, c_useStatic)

	return toGoBool(ret)
}

func Fn_gdk_window_set_title(window unsafe.Pointer, title string) {
	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	c_title := (*C.gchar)(C.CString(title))
	defer C.free(unsafe.Pointer(c_title))

	C.gdk_window_set_title(c_window, c_title)
}

func Fn_gdk_window_set_transient_for(window unsafe.Pointer, parent unsafe.Pointer) {
	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	c_parent := (*C.GdkWindow)(unsafe.Pointer(parent))

	C.gdk_window_set_transient_for(c_window, c_parent)
}

func Fn_gdk_window_set_type_hint(window unsafe.Pointer, hint int) {
	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	c_hint := (C.GdkWindowTypeHint)(hint)

	C.gdk_window_set_type_hint(c_window, c_hint)
}

func Fn_gdk_window_set_urgency_hint(window unsafe.Pointer, urgent bool) {
	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	c_urgent := toCBool(urgent)

	C.gdk_window_set_urgency_hint(c_window, c_urgent)
}

func Fn_gdk_window_set_user_data(window unsafe.Pointer, userData unsafe.Pointer) {
	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	c_userData := (C.gpointer)(userData)

	C.gdk_window_set_user_data(c_window, c_userData)
}

func Fn_gdk_window_shape_combine_region(window unsafe.Pointer, shapeRegion unsafe.Pointer, offsetX int, offsetY int) {
	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	c_shapeRegion := (*C.cairo_region_t)(unsafe.Pointer(shapeRegion))

	c_offsetX := (C.gint)(offsetX)

	c_offsetY := (C.gint)(offsetY)

	C.gdk_window_shape_combine_region(c_window, c_shapeRegion, c_offsetX, c_offsetY)
}

func Fn_gdk_window_show(window unsafe.Pointer) {
	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	C.gdk_window_show(c_window)
}

func Fn_gdk_window_show_unraised(window unsafe.Pointer) {
	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	C.gdk_window_show_unraised(c_window)
}

func Fn_gdk_window_stick(window unsafe.Pointer) {
	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	C.gdk_window_stick(c_window)
}

func Fn_gdk_window_thaw_toplevel_updates_libgtk_only(window unsafe.Pointer) {
	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	C.gdk_window_thaw_toplevel_updates_libgtk_only(c_window)
}

func Fn_gdk_window_thaw_updates(window unsafe.Pointer) {
	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	C.gdk_window_thaw_updates(c_window)
}

func Fn_gdk_window_unfullscreen(window unsafe.Pointer) {
	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	C.gdk_window_unfullscreen(c_window)
}

func Fn_gdk_window_unmaximize(window unsafe.Pointer) {
	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	C.gdk_window_unmaximize(c_window)
}

func Fn_gdk_window_unstick(window unsafe.Pointer) {
	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	C.gdk_window_unstick(c_window)
}

func Fn_gdk_window_withdraw(window unsafe.Pointer) {
	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	C.gdk_window_withdraw(c_window)
}

func Fn_gdk_window_at_pointer(winX *int, winY *int) unsafe.Pointer {
	c_winX := (*C.gint)(unsafe.Pointer(winX))

	c_winY := (*C.gint)(unsafe.Pointer(winY))

	ret := C.gdk_window_at_pointer(c_winX, c_winY)

	return unsafe.Pointer(ret)
}

func Fn_gdk_window_constrain_size(geometry unsafe.Pointer, flags int, width int, height int, newWidth *int, newHeight *int) {
	c_geometry := (*C.GdkGeometry)(unsafe.Pointer(geometry))

	c_flags := (C.GdkWindowHints)(flags)

	c_width := (C.gint)(width)

	c_height := (C.gint)(height)

	c_newWidth := (*C.gint)(unsafe.Pointer(newWidth))

	c_newHeight := (*C.gint)(unsafe.Pointer(newHeight))

	C.gdk_window_constrain_size(c_geometry, c_flags, c_width, c_height, c_newWidth, c_newHeight)
}

func Fn_gdk_window_process_all_updates() {
	C.gdk_window_process_all_updates()
}

func Fn_gdk_window_set_debug_updates(setting bool) {
	c_setting := toCBool(setting)

	C.gdk_window_set_debug_updates(c_setting)
}
