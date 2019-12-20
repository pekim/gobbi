// Code generated - DO NOT EDIT.
// +build !gdk_2.6,!gdk_2.8,!gdk_3.4,!gdk_3.8,!gdk_3.16,!gdk_3.20,!gdk_3.22

package gdk

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

func Fn_add_option_entries_libgtk_only(group string) {}

func Fn_atom_intern(atomName string, onlyIfExists string) {}

func Fn_beep() {}

func Fn_cairo_get_clip_rectangle(cr string, rect string) {}

func Fn_cairo_region_create_from_surface(surface string) {}

func Fn_color_parse(spec string, color string) {}

func Fn_drag_abort(context string, time string) {}

func Fn_drag_begin(window string, targets string) {}

func Fn_drag_begin_for_device(window string, device string, targets string) {}

func Fn_drag_drop(context string, time string) {}

func Fn_drag_get_selection(context string) {}

func Fn_drag_motion(context string, destWindow string, protocol string, xRoot string, yRoot string, suggestedAction string, possibleActions string, time string) {
}

func Fn_drag_status(context string, action string, time string) {}

func Fn_drop_finish(context string, success string, time string) {}

func Fn_drop_reply(context string, accepted string, time string) {}

func Fn_error_trap_pop() {}

func Fn_error_trap_push() {}

func Fn_event_get() {}

func Fn_event_handler_set(func_ string, data string, notify string) {}

func Fn_event_peek() {}

func Fn_events_pending() {}

func Fn_flush() {}

func Fn_get_default_root_window() {}

func Fn_get_display() {}

func Fn_get_program_class() {}

func Fn_get_show_events() {}

func Fn_gl_error_quark() {}

func Fn_init(argc string, argv string) {}

func Fn_init_check(argc string, argv string) {}

func Fn_keyboard_grab(window string, ownerEvents string, time string) {}

func Fn_keyboard_ungrab(time string) {}

func Fn_keyval_convert_case(symbol string, lower string, upper string) {}

func Fn_keyval_from_name(keyvalName string) {}

func Fn_keyval_is_lower(keyval string) {}

func Fn_keyval_is_upper(keyval string) {}

func Fn_keyval_name(keyval string) {}

func Fn_keyval_to_lower(keyval string) {}

func Fn_keyval_to_unicode(keyval string) {}

func Fn_keyval_to_upper(keyval string) {}

func Fn_list_visuals() {}

func Fn_offscreen_window_get_surface(window string) {}

func Fn_pango_context_get() {}

func Fn_pango_layout_get_clip_region(layout string, xOrigin string, yOrigin string, indexRanges string, nRanges string) {
}

func Fn_pango_layout_line_get_clip_region(line string, xOrigin string, yOrigin string, indexRanges string, nRanges string) {
}

func Fn_pixbuf_get_from_surface(surface string, srcX string, srcY string, width string, height string) {}

func Fn_pixbuf_get_from_window(window string, srcX string, srcY string, width string, height string) {}

func Fn_pointer_grab(window string, ownerEvents string, eventMask string, confineTo string, cursor string, time string) {
}

func Fn_pointer_is_grabbed() {}

func Fn_pointer_ungrab(time string) {}

func Fn_pre_parse_libgtk_only() {}

func Fn_property_change(window string, property string, type_ string, format string, mode string, data string, nelements string) {
}

func Fn_property_delete(window string, property string) {}

func Fn_property_get(window string, property string, type_ string, offset string, length string, pdelete string, actualPropertyType string, actualFormat string, actualLength string, data string) {
}

func Fn_query_depths(depths string, count string) {}

func Fn_query_visual_types(visualTypes string, count string) {}

func Fn_selection_convert(requestor string, selection string, target string, time string) {}

func Fn_selection_owner_get(selection string) {}

func Fn_selection_owner_set(owner string, selection string, time string, sendEvent string) {}

func Fn_selection_property_get(requestor string, data string, propType string, propFormat string) {}

func Fn_selection_send_notify(requestor string, selection string, target string, property string, time string) {
}

func Fn_set_double_click_time(msec string) {}

func Fn_set_program_class(programClass string) {}

func Fn_set_show_events(showEvents string) {}

func Fn_setting_get(name string, value string) {}

func Fn_synthesize_window_state(window string, unsetFlags string, setFlags string) {}

func Fn_threads_enter() {}

func Fn_threads_init() {}

func Fn_threads_leave() {}

func Fn_unicode_to_keyval(wc string) {}

func Fn_utf8_to_string_target(str string) {}
