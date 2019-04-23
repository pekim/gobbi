// This is a generated file - DO NOT EDIT

package gdk

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gdk/gdk.h>
// #include <stdlib.h>
import "C"

// Blacklisted : GdkAxisUse

// Blacklisted : GdkByteOrder

type CrossingMode C.GdkCrossingMode

const (
	GDK_CROSSING_NORMAL        CrossingMode = 0
	GDK_CROSSING_GRAB          CrossingMode = 1
	GDK_CROSSING_UNGRAB        CrossingMode = 2
	GDK_CROSSING_GTK_GRAB      CrossingMode = 3
	GDK_CROSSING_GTK_UNGRAB    CrossingMode = 4
	GDK_CROSSING_STATE_CHANGED CrossingMode = 5
	GDK_CROSSING_TOUCH_BEGIN   CrossingMode = 6
	GDK_CROSSING_TOUCH_END     CrossingMode = 7
	GDK_CROSSING_DEVICE_SWITCH CrossingMode = 8
)

// Blacklisted : GdkCursorType

// Blacklisted : GdkDeviceType

// Blacklisted : GdkDragProtocol

type EventType C.GdkEventType

const (
	GDK_NOTHING             EventType = -1
	GDK_DELETE              EventType = 0
	GDK_DESTROY             EventType = 1
	GDK_EXPOSE              EventType = 2
	GDK_MOTION_NOTIFY       EventType = 3
	GDK_BUTTON_PRESS        EventType = 4
	GDK_2BUTTON_PRESS       EventType = 5
	GDK_DOUBLE_BUTTON_PRESS EventType = 5
	GDK_3BUTTON_PRESS       EventType = 6
	GDK_TRIPLE_BUTTON_PRESS EventType = 6
	GDK_BUTTON_RELEASE      EventType = 7
	GDK_KEY_PRESS           EventType = 8
	GDK_KEY_RELEASE         EventType = 9
	GDK_ENTER_NOTIFY        EventType = 10
	GDK_LEAVE_NOTIFY        EventType = 11
	GDK_FOCUS_CHANGE        EventType = 12
	GDK_CONFIGURE           EventType = 13
	GDK_MAP                 EventType = 14
	GDK_UNMAP               EventType = 15
	GDK_PROPERTY_NOTIFY     EventType = 16
	GDK_SELECTION_CLEAR     EventType = 17
	GDK_SELECTION_REQUEST   EventType = 18
	GDK_SELECTION_NOTIFY    EventType = 19
	GDK_PROXIMITY_IN        EventType = 20
	GDK_PROXIMITY_OUT       EventType = 21
	GDK_DRAG_ENTER          EventType = 22
	GDK_DRAG_LEAVE          EventType = 23
	GDK_DRAG_MOTION         EventType = 24
	GDK_DRAG_STATUS         EventType = 25
	GDK_DROP_START          EventType = 26
	GDK_DROP_FINISHED       EventType = 27
	GDK_CLIENT_EVENT        EventType = 28
	GDK_VISIBILITY_NOTIFY   EventType = 29
	GDK_SCROLL              EventType = 31
	GDK_WINDOW_STATE        EventType = 32
	GDK_SETTING             EventType = 33
	GDK_OWNER_CHANGE        EventType = 34
	GDK_GRAB_BROKEN         EventType = 35
	GDK_DAMAGE              EventType = 36
	GDK_TOUCH_BEGIN         EventType = 37
	GDK_TOUCH_UPDATE        EventType = 38
	GDK_TOUCH_END           EventType = 39
	GDK_TOUCH_CANCEL        EventType = 40
	GDK_TOUCHPAD_SWIPE      EventType = 41
	GDK_TOUCHPAD_PINCH      EventType = 42
	GDK_PAD_BUTTON_PRESS    EventType = 43
	GDK_PAD_BUTTON_RELEASE  EventType = 44
	GDK_PAD_RING            EventType = 45
	GDK_PAD_STRIP           EventType = 46
	GDK_PAD_GROUP_MODE      EventType = 47
	GDK_EVENT_LAST          EventType = 48
)

// Blacklisted : GdkFilterReturn

// Blacklisted : GdkGrabOwnership

// Blacklisted : GdkGrabStatus

// Blacklisted : GdkGravity

// Blacklisted : GdkInputMode

// Blacklisted : GdkInputSource

type NotifyType C.GdkNotifyType

const (
	GDK_NOTIFY_ANCESTOR          NotifyType = 0
	GDK_NOTIFY_VIRTUAL           NotifyType = 1
	GDK_NOTIFY_INFERIOR          NotifyType = 2
	GDK_NOTIFY_NONLINEAR         NotifyType = 3
	GDK_NOTIFY_NONLINEAR_VIRTUAL NotifyType = 4
	GDK_NOTIFY_UNKNOWN           NotifyType = 5
)

// Blacklisted : GdkOwnerChange

// Blacklisted : GdkPropMode

// Blacklisted : guint

// Blacklisted : GdkScrollDirection

// Blacklisted : GdkSettingAction

type Status C.GdkStatus

const (
	GDK_OK          Status = 0
	GDK_ERROR       Status = -1
	GDK_ERROR_PARAM Status = -2
	GDK_ERROR_FILE  Status = -3
	GDK_ERROR_MEM   Status = -4
)

// Blacklisted : GdkTouchpadGesturePhase

// Blacklisted : GdkVisibilityState

// Blacklisted : GdkVisualType

// Blacklisted : GdkWindowEdge

type WindowType C.GdkWindowType

const (
	GDK_WINDOW_ROOT       WindowType = 0
	GDK_WINDOW_TOPLEVEL   WindowType = 1
	GDK_WINDOW_CHILD      WindowType = 2
	GDK_WINDOW_TEMP       WindowType = 3
	GDK_WINDOW_FOREIGN    WindowType = 4
	GDK_WINDOW_OFFSCREEN  WindowType = 5
	GDK_WINDOW_SUBSURFACE WindowType = 6
)

type WindowTypeHint C.GdkWindowTypeHint

const (
	GDK_WINDOW_TYPE_HINT_NORMAL        WindowTypeHint = 0
	GDK_WINDOW_TYPE_HINT_DIALOG        WindowTypeHint = 1
	GDK_WINDOW_TYPE_HINT_MENU          WindowTypeHint = 2
	GDK_WINDOW_TYPE_HINT_TOOLBAR       WindowTypeHint = 3
	GDK_WINDOW_TYPE_HINT_SPLASHSCREEN  WindowTypeHint = 4
	GDK_WINDOW_TYPE_HINT_UTILITY       WindowTypeHint = 5
	GDK_WINDOW_TYPE_HINT_DOCK          WindowTypeHint = 6
	GDK_WINDOW_TYPE_HINT_DESKTOP       WindowTypeHint = 7
	GDK_WINDOW_TYPE_HINT_DROPDOWN_MENU WindowTypeHint = 8
	GDK_WINDOW_TYPE_HINT_POPUP_MENU    WindowTypeHint = 9
	GDK_WINDOW_TYPE_HINT_TOOLTIP       WindowTypeHint = 10
	GDK_WINDOW_TYPE_HINT_NOTIFICATION  WindowTypeHint = 11
	GDK_WINDOW_TYPE_HINT_COMBO         WindowTypeHint = 12
	GDK_WINDOW_TYPE_HINT_DND           WindowTypeHint = 13
)

type WindowWindowClass C.GdkWindowWindowClass

const (
	GDK_INPUT_OUTPUT WindowWindowClass = 0
	GDK_INPUT_ONLY   WindowWindowClass = 1
)
