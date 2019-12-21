// Code generated - DO NOT EDIT.
// +build !gdk_2.6,!gdk_2.8,!gdk_3.4,!gdk_3.8,!gdk_3.16,!gdk_3.20,!gdk_3.22

package gdk

import (
	c "github.com/pekim/gobbi/lib/internal/c"
	"unsafe"
)

// #include <gdk/gdk.h>
import "C"

// bitfields
type DragAction C.GdkDragAction
type EventMask C.GdkEventMask
type ModifierType C.GdkModifierType
type WMDecoration C.GdkWMDecoration
type WMFunction C.GdkWMFunction
type WindowAttributesType C.GdkWindowAttributesType
type WindowHints C.GdkWindowHints
type WindowState C.GdkWindowState

// enumerations
type AxisUse C.GdkAxisUse
type ByteOrder C.GdkByteOrder
type CrossingMode C.GdkCrossingMode
type CursorType C.GdkCursorType
type DevicePadFeature C.GdkDevicePadFeature
type DeviceType C.GdkDeviceType
type DragProtocol C.GdkDragProtocol
type EventType C.GdkEventType
type FilterReturn C.GdkFilterReturn
type GrabOwnership C.GdkGrabOwnership
type GrabStatus C.GdkGrabStatus
type Gravity C.GdkGravity
type InputMode C.GdkInputMode
type InputSource C.GdkInputSource
type NotifyType C.GdkNotifyType
type OwnerChange C.GdkOwnerChange
type PropMode C.GdkPropMode
type PropertyState C.GdkPropertyState
type ScrollDirection C.GdkScrollDirection
type SettingAction C.GdkSettingAction
type Status C.GdkStatus
type TouchpadGesturePhase C.GdkTouchpadGesturePhase
type VisibilityState C.GdkVisibilityState
type VisualType C.GdkVisualType
type WindowEdge C.GdkWindowEdge
type WindowType C.GdkWindowType
type WindowTypeHint C.GdkWindowTypeHint
type WindowWindowClass C.GdkWindowWindowClass

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

// classes
type AppLaunchContext C.GdkAppLaunchContext
type Cursor C.GdkCursor
type Device C.GdkDevice
type DeviceManager C.GdkDeviceManager
type DeviceTool C.GdkDeviceTool
type Display C.GdkDisplay
type DisplayManager C.GdkDisplayManager
type DragContext C.GdkDragContext
type DrawingContext C.GdkDrawingContext
type FrameClock C.GdkFrameClock
type GLContext C.GdkGLContext
type Keymap C.GdkKeymap
type Monitor C.GdkMonitor
type Screen C.GdkScreen
type Seat C.GdkSeat
type Visual C.GdkVisual
type Window C.GdkWindow

// interfaces
type DevicePad C.GdkDevicePad

func Fn_add_option_entries_libgtk_only(group c.UndefinedParamType) {}

func Fn_atom_intern(atomName c.UndefinedParamType, onlyIfExists bool) {}

func Fn_beep() {}

func Fn_cairo_get_clip_rectangle(cr c.UndefinedParamType) {}

func Fn_cairo_region_create_from_surface(surface c.UndefinedParamType) {}

func Fn_color_parse(spec c.UndefinedParamType) {}

func Fn_drag_abort(context c.UndefinedParamType, time uint32) {}

func Fn_drag_begin(window c.UndefinedParamType, targets c.UndefinedParamType) {}

func Fn_drag_begin_for_device(window c.UndefinedParamType, device c.UndefinedParamType, targets c.UndefinedParamType) {
}

func Fn_drag_drop(context c.UndefinedParamType, time uint32) {}

func Fn_drag_get_selection(context c.UndefinedParamType) {}

func Fn_drag_motion(context c.UndefinedParamType, destWindow c.UndefinedParamType, protocol c.UndefinedParamType, xRoot int, yRoot int, suggestedAction c.UndefinedParamType, possibleActions c.UndefinedParamType, time uint32) {
}

func Fn_drag_status(context c.UndefinedParamType, action c.UndefinedParamType, time uint32) {}

func Fn_drop_finish(context c.UndefinedParamType, success bool, time uint32) {}

func Fn_drop_reply(context c.UndefinedParamType, accepted bool, time uint32) {}

func Fn_error_trap_pop() {}

func Fn_error_trap_push() {}

func Fn_event_get() {}

func Fn_event_handler_set(func_ c.UndefinedParamType, data unsafe.Pointer, notify c.UndefinedParamType) {
}

func Fn_event_peek() {}

func Fn_events_pending() {}

func Fn_flush() {}

func Fn_get_default_root_window() {}

func Fn_get_display() {}

func Fn_get_program_class() {}

func Fn_get_show_events() {}

func Fn_gl_error_quark() {}

func Fn_init(argc *int, argv c.UndefinedParamType) {}

func Fn_init_check(argc *int, argv c.UndefinedParamType) {}

func Fn_keyboard_grab(window c.UndefinedParamType, ownerEvents bool, time uint32) {}

func Fn_keyboard_ungrab(time uint32) {}

func Fn_keyval_convert_case(symbol uint) {}

func Fn_keyval_from_name(keyvalName c.UndefinedParamType) {}

func Fn_keyval_is_lower(keyval uint) {}

func Fn_keyval_is_upper(keyval uint) {}

func Fn_keyval_name(keyval uint) {}

func Fn_keyval_to_lower(keyval uint) {}

func Fn_keyval_to_unicode(keyval uint) {}

func Fn_keyval_to_upper(keyval uint) {}

func Fn_list_visuals() {}

func Fn_offscreen_window_get_surface(window c.UndefinedParamType) {}

func Fn_pango_context_get() {}

func Fn_pango_layout_get_clip_region(layout c.UndefinedParamType, xOrigin int, yOrigin int, indexRanges *int, nRanges int) {
}

func Fn_pango_layout_line_get_clip_region(line c.UndefinedParamType, xOrigin int, yOrigin int, indexRanges c.UndefinedParamType, nRanges int) {
}

func Fn_pixbuf_get_from_surface(surface c.UndefinedParamType, srcX int, srcY int, width int, height int) {
}

func Fn_pixbuf_get_from_window(window c.UndefinedParamType, srcX int, srcY int, width int, height int) {
}

func Fn_pointer_grab(window c.UndefinedParamType, ownerEvents bool, eventMask c.UndefinedParamType, confineTo c.UndefinedParamType, cursor c.UndefinedParamType, time uint32) {
}

func Fn_pointer_is_grabbed() {}

func Fn_pointer_ungrab(time uint32) {}

func Fn_pre_parse_libgtk_only() {}

func Fn_property_change(window c.UndefinedParamType, property c.UndefinedParamType, type_ c.UndefinedParamType, format int, mode c.UndefinedParamType, data c.UndefinedParamType, nelements int) {
}

func Fn_property_delete(window c.UndefinedParamType, property c.UndefinedParamType) {}

func Fn_property_get(window c.UndefinedParamType, property c.UndefinedParamType, type_ c.UndefinedParamType, offset uint64, length uint64, pdelete int) {
}

func Fn_query_depths() {}

func Fn_query_visual_types() {}

func Fn_selection_convert(requestor c.UndefinedParamType, selection c.UndefinedParamType, target c.UndefinedParamType, time uint32) {
}

func Fn_selection_owner_get(selection c.UndefinedParamType) {}

func Fn_selection_owner_set(owner c.UndefinedParamType, selection c.UndefinedParamType, time uint32, sendEvent bool) {
}

func Fn_selection_property_get(requestor c.UndefinedParamType, data c.UndefinedParamType, propType c.UndefinedParamType, propFormat *int) {
}

func Fn_selection_send_notify(requestor c.UndefinedParamType, selection c.UndefinedParamType, target c.UndefinedParamType, property c.UndefinedParamType, time uint32) {
}

func Fn_set_double_click_time(msec uint) {}

func Fn_set_program_class(programClass c.UndefinedParamType) {}

func Fn_set_show_events(showEvents bool) {}

func Fn_setting_get(name c.UndefinedParamType, value c.UndefinedParamType) {}

func Fn_synthesize_window_state(window c.UndefinedParamType, unsetFlags c.UndefinedParamType, setFlags c.UndefinedParamType) {
}

func Fn_threads_enter() {}

func Fn_threads_init() {}

func Fn_threads_leave() {}

func Fn_unicode_to_keyval(wc uint32) {}

func Fn_utf8_to_string_target(str c.UndefinedParamType) {}
