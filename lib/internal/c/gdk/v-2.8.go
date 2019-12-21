// Code generated - DO NOT EDIT.
// +build gdk_2.8

package gdk

import c "github.com/pekim/gobbi/lib/internal/c"

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

func Fn_atom_intern(atomName c.UndefinedParamType, onlyIfExists c.UndefinedParamType) {}

func Fn_beep() {}

func Fn_cairo_create(window c.UndefinedParamType) {}

func Fn_cairo_get_clip_rectangle(cr c.UndefinedParamType, rect c.UndefinedParamType) {}

func Fn_cairo_rectangle(cr c.UndefinedParamType, rectangle c.UndefinedParamType) {}

func Fn_cairo_region(cr c.UndefinedParamType, region c.UndefinedParamType) {}

func Fn_cairo_region_create_from_surface(surface c.UndefinedParamType) {}

func Fn_cairo_set_source_color(cr c.UndefinedParamType, color c.UndefinedParamType) {}

func Fn_cairo_set_source_pixbuf(cr c.UndefinedParamType, pixbuf c.UndefinedParamType, pixbufX c.UndefinedParamType, pixbufY c.UndefinedParamType) {
}

func Fn_color_parse(spec c.UndefinedParamType, color c.UndefinedParamType) {}

func Fn_drag_abort(context c.UndefinedParamType, time c.UndefinedParamType) {}

func Fn_drag_begin(window c.UndefinedParamType, targets c.UndefinedParamType) {}

func Fn_drag_begin_for_device(window c.UndefinedParamType, device c.UndefinedParamType, targets c.UndefinedParamType) {
}

func Fn_drag_drop(context c.UndefinedParamType, time c.UndefinedParamType) {}

func Fn_drag_drop_succeeded(context c.UndefinedParamType) {}

func Fn_drag_find_window_for_screen(context c.UndefinedParamType, dragWindow c.UndefinedParamType, screen c.UndefinedParamType, xRoot c.UndefinedParamType, yRoot c.UndefinedParamType, destWindow c.UndefinedParamType, protocol c.UndefinedParamType) {
}

func Fn_drag_get_selection(context c.UndefinedParamType) {}

func Fn_drag_motion(context c.UndefinedParamType, destWindow c.UndefinedParamType, protocol c.UndefinedParamType, xRoot c.UndefinedParamType, yRoot c.UndefinedParamType, suggestedAction c.UndefinedParamType, possibleActions c.UndefinedParamType, time c.UndefinedParamType) {
}

func Fn_drag_status(context c.UndefinedParamType, action c.UndefinedParamType, time c.UndefinedParamType) {
}

func Fn_drop_finish(context c.UndefinedParamType, success c.UndefinedParamType, time c.UndefinedParamType) {
}

func Fn_drop_reply(context c.UndefinedParamType, accepted c.UndefinedParamType, time c.UndefinedParamType) {
}

func Fn_error_trap_pop() {}

func Fn_error_trap_push() {}

func Fn_event_get() {}

func Fn_event_handler_set(func_ c.UndefinedParamType, data c.UndefinedParamType, notify c.UndefinedParamType) {
}

func Fn_event_peek() {}

func Fn_events_pending() {}

func Fn_flush() {}

func Fn_get_default_root_window() {}

func Fn_get_display() {}

func Fn_get_display_arg_name() {}

func Fn_get_program_class() {}

func Fn_get_show_events() {}

func Fn_gl_error_quark() {}

func Fn_init(argc c.UndefinedParamType, argv c.UndefinedParamType) {}

func Fn_init_check(argc c.UndefinedParamType, argv c.UndefinedParamType) {}

func Fn_keyboard_grab(window c.UndefinedParamType, ownerEvents c.UndefinedParamType, time c.UndefinedParamType) {
}

func Fn_keyboard_ungrab(time c.UndefinedParamType) {}

func Fn_keyval_convert_case(symbol c.UndefinedParamType, lower c.UndefinedParamType, upper c.UndefinedParamType) {
}

func Fn_keyval_from_name(keyvalName c.UndefinedParamType) {}

func Fn_keyval_is_lower(keyval c.UndefinedParamType) {}

func Fn_keyval_is_upper(keyval c.UndefinedParamType) {}

func Fn_keyval_name(keyval c.UndefinedParamType) {}

func Fn_keyval_to_lower(keyval c.UndefinedParamType) {}

func Fn_keyval_to_unicode(keyval c.UndefinedParamType) {}

func Fn_keyval_to_upper(keyval c.UndefinedParamType) {}

func Fn_list_visuals() {}

func Fn_notify_startup_complete() {}

func Fn_offscreen_window_get_surface(window c.UndefinedParamType) {}

func Fn_pango_context_get() {}

func Fn_pango_context_get_for_screen(screen c.UndefinedParamType) {}

func Fn_pango_layout_get_clip_region(layout c.UndefinedParamType, xOrigin c.UndefinedParamType, yOrigin c.UndefinedParamType, indexRanges c.UndefinedParamType, nRanges c.UndefinedParamType) {
}

func Fn_pango_layout_line_get_clip_region(line c.UndefinedParamType, xOrigin c.UndefinedParamType, yOrigin c.UndefinedParamType, indexRanges c.UndefinedParamType, nRanges c.UndefinedParamType) {
}

func Fn_parse_args(argc c.UndefinedParamType, argv c.UndefinedParamType) {}

func Fn_pixbuf_get_from_surface(surface c.UndefinedParamType, srcX c.UndefinedParamType, srcY c.UndefinedParamType, width c.UndefinedParamType, height c.UndefinedParamType) {
}

func Fn_pixbuf_get_from_window(window c.UndefinedParamType, srcX c.UndefinedParamType, srcY c.UndefinedParamType, width c.UndefinedParamType, height c.UndefinedParamType) {
}

func Fn_pointer_grab(window c.UndefinedParamType, ownerEvents c.UndefinedParamType, eventMask c.UndefinedParamType, confineTo c.UndefinedParamType, cursor c.UndefinedParamType, time c.UndefinedParamType) {
}

func Fn_pointer_is_grabbed() {}

func Fn_pointer_ungrab(time c.UndefinedParamType) {}

func Fn_pre_parse_libgtk_only() {}

func Fn_property_change(window c.UndefinedParamType, property c.UndefinedParamType, type_ c.UndefinedParamType, format c.UndefinedParamType, mode c.UndefinedParamType, data c.UndefinedParamType, nelements c.UndefinedParamType) {
}

func Fn_property_delete(window c.UndefinedParamType, property c.UndefinedParamType) {}

func Fn_property_get(window c.UndefinedParamType, property c.UndefinedParamType, type_ c.UndefinedParamType, offset c.UndefinedParamType, length c.UndefinedParamType, pdelete c.UndefinedParamType, actualPropertyType c.UndefinedParamType, actualFormat c.UndefinedParamType, actualLength c.UndefinedParamType, data c.UndefinedParamType) {
}

func Fn_query_depths(depths c.UndefinedParamType, count c.UndefinedParamType) {}

func Fn_query_visual_types(visualTypes c.UndefinedParamType, count c.UndefinedParamType) {}

func Fn_selection_convert(requestor c.UndefinedParamType, selection c.UndefinedParamType, target c.UndefinedParamType, time c.UndefinedParamType) {
}

func Fn_selection_owner_get(selection c.UndefinedParamType) {}

func Fn_selection_owner_get_for_display(display c.UndefinedParamType, selection c.UndefinedParamType) {
}

func Fn_selection_owner_set(owner c.UndefinedParamType, selection c.UndefinedParamType, time c.UndefinedParamType, sendEvent c.UndefinedParamType) {
}

func Fn_selection_owner_set_for_display(display c.UndefinedParamType, owner c.UndefinedParamType, selection c.UndefinedParamType, time c.UndefinedParamType, sendEvent c.UndefinedParamType) {
}

func Fn_selection_property_get(requestor c.UndefinedParamType, data c.UndefinedParamType, propType c.UndefinedParamType, propFormat c.UndefinedParamType) {
}

func Fn_selection_send_notify(requestor c.UndefinedParamType, selection c.UndefinedParamType, target c.UndefinedParamType, property c.UndefinedParamType, time c.UndefinedParamType) {
}

func Fn_selection_send_notify_for_display(display c.UndefinedParamType, requestor c.UndefinedParamType, selection c.UndefinedParamType, target c.UndefinedParamType, property c.UndefinedParamType, time c.UndefinedParamType) {
}

func Fn_set_double_click_time(msec c.UndefinedParamType) {}

func Fn_set_program_class(programClass c.UndefinedParamType) {}

func Fn_set_show_events(showEvents c.UndefinedParamType) {}

func Fn_setting_get(name c.UndefinedParamType, value c.UndefinedParamType) {}

func Fn_synthesize_window_state(window c.UndefinedParamType, unsetFlags c.UndefinedParamType, setFlags c.UndefinedParamType) {
}

func Fn_text_property_to_utf8_list_for_display(display c.UndefinedParamType, encoding c.UndefinedParamType, format c.UndefinedParamType, text c.UndefinedParamType, length c.UndefinedParamType, list c.UndefinedParamType) {
}

func Fn_threads_enter() {}

func Fn_threads_init() {}

func Fn_threads_leave() {}

func Fn_threads_set_lock_functions(enterFn c.UndefinedParamType, leaveFn c.UndefinedParamType) {}

func Fn_unicode_to_keyval(wc c.UndefinedParamType) {}

func Fn_utf8_to_string_target(str c.UndefinedParamType) {}
