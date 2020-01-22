// Code generated - DO NOT EDIT.
// +build gdk_3.20

package gdk

import (
	glib "github.com/pekim/gobbi/lib/glib"
	gdk "github.com/pekim/gobbi/lib/internal/c/gdk"
	"unsafe"
)

// UNSUPPORTED : XEvent : blacklisted

// SeatCapabilities is a representation of the C bitfield GdkSeatCapabilities.
type SeatCapabilities int

// SeatCapabilities_none is a representation of the C bitfield member GDK_SEAT_CAPABILITY_NONE.
const SeatCapabilities_none = SeatCapabilities(0)

// SeatCapabilities_pointer is a representation of the C bitfield member GDK_SEAT_CAPABILITY_POINTER.
const SeatCapabilities_pointer = SeatCapabilities(1)

// SeatCapabilities_touch is a representation of the C bitfield member GDK_SEAT_CAPABILITY_TOUCH.
const SeatCapabilities_touch = SeatCapabilities(2)

// SeatCapabilities_tablet_stylus is a representation of the C bitfield member GDK_SEAT_CAPABILITY_TABLET_STYLUS.
const SeatCapabilities_tablet_stylus = SeatCapabilities(4)

// SeatCapabilities_keyboard is a representation of the C bitfield member GDK_SEAT_CAPABILITY_KEYBOARD.
const SeatCapabilities_keyboard = SeatCapabilities(8)

// SeatCapabilities_all_pointing is a representation of the C bitfield member GDK_SEAT_CAPABILITY_ALL_POINTING.
const SeatCapabilities_all_pointing = SeatCapabilities(7)

// SeatCapabilities_all is a representation of the C bitfield member GDK_SEAT_CAPABILITY_ALL.
const SeatCapabilities_all = SeatCapabilities(15)

// DragCancelReason is a representation of the C enumeration GdkDragCancelReason.
type DragCancelReason int

// DragCancelReason_no_target is a representation of the C enumeration member GDK_DRAG_CANCEL_NO_TARGET.
const DragCancelReason_no_target = DragCancelReason(0)

// DragCancelReason_user_cancelled is a representation of the C enumeration member GDK_DRAG_CANCEL_USER_CANCELLED.
const DragCancelReason_user_cancelled = DragCancelReason(1)

// DragCancelReason_error is a representation of the C enumeration member GDK_DRAG_CANCEL_ERROR.
const DragCancelReason_error = DragCancelReason(2)

// UNSUPPORTED : gdk_cairo_get_clip_rectangle : has [in]out param, rect

// UNSUPPORTED : gdk_color_parse : has [in]out param, color

// DragBeginFromPoint wraps the C function gdk_drag_begin_from_point.
//
// since 3.20
func DragBeginFromPoint(window *Window, device *Device, targets *glib.List, xRoot int, yRoot int) *DragContext {
	sys_window := window.ToC()
	sys_device := device.ToC()
	sys_targets := targets.ToC()
	sys_xRoot := xRoot
	sys_yRoot := yRoot
	retSys := gdk.Fn_gdk_drag_begin_from_point(sys_window, sys_device, sys_targets, sys_xRoot, sys_yRoot)
	ret := DragContextNewFromC(retSys)

	return ret
}

// DragDropDone wraps the C function gdk_drag_drop_done.
//
// since 3.20
func DragDropDone(context *DragContext, success bool) {
	sys_context := context.ToC()
	sys_success := success
	gdk.Fn_gdk_drag_drop_done(sys_context, sys_success)
}

// UNSUPPORTED : gdk_drag_find_window_for_screen : has [in]out param, dest_window

// UNSUPPORTED : gdk_event_handler_set : parameter 'func' is callback

// UNSUPPORTED : gdk_events_get_angle : has [in]out param, angle

// UNSUPPORTED : gdk_events_get_center : has [in]out param, x

// UNSUPPORTED : gdk_events_get_distance : has [in]out param, distance

// UNSUPPORTED : gdk_init : has array param, argv

// UNSUPPORTED : gdk_init_check : has array param, argv

// UNSUPPORTED : gdk_keyval_convert_case : has [in]out param, lower

// UNSUPPORTED : gdk_pango_layout_line_get_clip_region : parameter 'index_ranges' is array parameter without length parameter

// UNSUPPORTED : gdk_parse_args : has array param, argv

// UNSUPPORTED : gdk_property_get : has [in]out param, actual_property_type

// UNSUPPORTED : gdk_query_depths : has array param, depths

// UNSUPPORTED : gdk_query_visual_types : has array param, visual_types

// UNSUPPORTED : gdk_synthesize_window_state : blacklisted

// UNSUPPORTED : gdk_text_property_to_utf8_list_for_display : parameter 'list' is array parameter without length parameter

// UNSUPPORTED : gdk_threads_add_idle : parameter 'function' is callback

// UNSUPPORTED : gdk_threads_add_idle_full : parameter 'function' is callback

// UNSUPPORTED : gdk_threads_add_timeout : parameter 'function' is callback

// UNSUPPORTED : gdk_threads_add_timeout_full : parameter 'function' is callback

// UNSUPPORTED : gdk_threads_add_timeout_seconds : parameter 'function' is callback

// UNSUPPORTED : gdk_threads_add_timeout_seconds_full : parameter 'function' is callback

// UNSUPPORTED : gdk_threads_set_lock_functions : parameter 'enter_fn' is callback

// Event is a representation of the C union GdkEvent.
type Event struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GdkEvent that represents the Event.
func (recv *Event) ToC() unsafe.Pointer {
	return recv.native
}

// EventNewFromC creates a new Event from a pointer to the C GdkEvent that represents the Event.
func EventNewFromC(native unsafe.Pointer) *Event {
	return &Event{native: native}
}
