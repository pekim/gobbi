// Code generated - DO NOT EDIT.
// +build gdk_2.6

package gdk

import (
	cairo "github.com/pekim/gobbi/lib/internal/c/cairo"
	glib "github.com/pekim/gobbi/lib/internal/c/glib"
	gobject "github.com/pekim/gobbi/lib/internal/c/gobject"
	pango "github.com/pekim/gobbi/lib/internal/c/pango"
)

// #include <gdk/gdk.h>
import "C"

// aliases
type XEvent C.GdkXEvent

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

// unions
type Event C.GdkEvent

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

func Fn_add_option_entries_libgtk_only(group *glib.OptionGroup) {}

func Fn_atom_intern(atomName string, onlyIfExists bool) {}

func Fn_beep() {}

func Fn_cairo_get_clip_rectangle(cr cairo.Context) {}

func Fn_cairo_region_create_from_surface(surface cairo.Surface) {}

func Fn_color_parse(spec string) {}

func Fn_drag_abort(context *DragContext, time uint32) {}

func Fn_drag_begin(window *Window, targets *glib.List) {}

func Fn_drag_begin_for_device(window *Window, device *Device, targets *glib.List) {}

func Fn_drag_drop(context *DragContext, time uint32) {}

func Fn_drag_drop_succeeded(context *DragContext) {}

func Fn_drag_find_window_for_screen(context *DragContext, dragWindow *Window, screen *Screen, xRoot int, yRoot int) {
}

func Fn_drag_get_selection(context *DragContext) {}

func Fn_drag_motion(context *DragContext, destWindow *Window, protocol DragProtocol, xRoot int, yRoot int, suggestedAction DragAction, possibleActions DragAction, time uint32) {
}

func Fn_drag_status(context *DragContext, action DragAction, time uint32) {}

func Fn_drop_finish(context *DragContext, success bool, time uint32) {}

func Fn_drop_reply(context *DragContext, accepted bool, time uint32) {}

func Fn_error_trap_pop() {}

func Fn_error_trap_push() {}

func Fn_event_get() {}

// UNSUPPORTED : event_handler_set : has callback

func Fn_event_peek() {}

func Fn_events_pending() {}

func Fn_flush() {}

func Fn_get_default_root_window() {}

func Fn_get_display() {}

func Fn_get_display_arg_name() {}

func Fn_get_program_class() {}

func Fn_get_show_events() {}

func Fn_gl_error_quark() {}

func Fn_init(argc *int, argv *string) {}

func Fn_init_check(argc *int, argv *string) {}

func Fn_keyboard_grab(window *Window, ownerEvents bool, time uint32) {}

func Fn_keyboard_ungrab(time uint32) {}

func Fn_keyval_convert_case(symbol uint) {}

func Fn_keyval_from_name(keyvalName string) {}

func Fn_keyval_is_lower(keyval uint) {}

func Fn_keyval_is_upper(keyval uint) {}

func Fn_keyval_name(keyval uint) {}

func Fn_keyval_to_lower(keyval uint) {}

func Fn_keyval_to_unicode(keyval uint) {}

func Fn_keyval_to_upper(keyval uint) {}

func Fn_list_visuals() {}

func Fn_notify_startup_complete() {}

func Fn_offscreen_window_get_surface(window *Window) {}

func Fn_pango_context_get() {}

func Fn_pango_context_get_for_screen(screen *Screen) {}

func Fn_pango_layout_get_clip_region(layout *pango.Layout, xOrigin int, yOrigin int, indexRanges *int, nRanges int) {
}

func Fn_pango_layout_line_get_clip_region(line *pango.LayoutLine, xOrigin int, yOrigin int, indexRanges *int, nRanges int) {
}

func Fn_parse_args(argc *int, argv *string) {}

func Fn_pixbuf_get_from_surface(surface cairo.Surface, srcX int, srcY int, width int, height int) {}

func Fn_pixbuf_get_from_window(window *Window, srcX int, srcY int, width int, height int) {}

func Fn_pointer_grab(window *Window, ownerEvents bool, eventMask EventMask, confineTo *Window, cursor *Cursor, time uint32) {
}

func Fn_pointer_is_grabbed() {}

func Fn_pointer_ungrab(time uint32) {}

func Fn_pre_parse_libgtk_only() {}

func Fn_property_change(window *Window, property Atom, type_ Atom, format int, mode PropMode, data *uint8, nelements int) {
}

func Fn_property_delete(window *Window, property Atom) {}

func Fn_property_get(window *Window, property Atom, type_ Atom, offset uint64, length uint64, pdelete int) {
}

func Fn_query_depths() {}

func Fn_query_visual_types() {}

func Fn_selection_convert(requestor *Window, selection Atom, target Atom, time uint32) {}

func Fn_selection_owner_get(selection Atom) {}

func Fn_selection_owner_get_for_display(display *Display, selection Atom) {}

func Fn_selection_owner_set(owner *Window, selection Atom, time uint32, sendEvent bool) {}

func Fn_selection_owner_set_for_display(display *Display, owner *Window, selection Atom, time uint32, sendEvent bool) {
}

func Fn_selection_property_get(requestor *Window, data **uint8, propType *Atom, propFormat *int) {}

func Fn_selection_send_notify(requestor *Window, selection Atom, target Atom, property Atom, time uint32) {
}

func Fn_selection_send_notify_for_display(display *Display, requestor *Window, selection Atom, target Atom, property Atom, time uint32) {
}

func Fn_set_double_click_time(msec uint) {}

func Fn_set_program_class(programClass string) {}

func Fn_set_show_events(showEvents bool) {}

func Fn_setting_get(name string, value *gobject.Value) {}

func Fn_synthesize_window_state(window *Window, unsetFlags WindowState, setFlags WindowState) {}

func Fn_text_property_to_utf8_list_for_display(display *Display, encoding Atom, format int, text *uint8, length int) {
}

// UNSUPPORTED : threads_add_idle : has callback

// UNSUPPORTED : threads_add_idle_full : has callback

// UNSUPPORTED : threads_add_timeout : has callback

// UNSUPPORTED : threads_add_timeout_full : has callback

// UNSUPPORTED : threads_add_timeout_seconds : has callback

// UNSUPPORTED : threads_add_timeout_seconds_full : has callback

func Fn_threads_enter() {}

func Fn_threads_init() {}

func Fn_threads_leave() {}

// UNSUPPORTED : threads_set_lock_functions : has callback

func Fn_unicode_to_keyval(wc uint32) {}

func Fn_utf8_to_string_target(str string) {}
