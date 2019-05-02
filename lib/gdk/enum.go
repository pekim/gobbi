// This is a generated file - DO NOT EDIT

package gdk

import "C"

type AxisUse C.GdkAxisUse

const (
	GDK_AXIS_IGNORE   AxisUse = 0
	GDK_AXIS_X        AxisUse = 1
	GDK_AXIS_Y        AxisUse = 2
	GDK_AXIS_PRESSURE AxisUse = 3
	GDK_AXIS_XTILT    AxisUse = 4
	GDK_AXIS_YTILT    AxisUse = 5
	GDK_AXIS_WHEEL    AxisUse = 6
	GDK_AXIS_DISTANCE AxisUse = 7
	GDK_AXIS_ROTATION AxisUse = 8
	GDK_AXIS_SLIDER   AxisUse = 9
	GDK_AXIS_LAST     AxisUse = 10
)

type ByteOrder C.GdkByteOrder

const (
	GDK_LSB_FIRST ByteOrder = 0
	GDK_MSB_FIRST ByteOrder = 1
)

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

type CursorType C.GdkCursorType

const (
	GDK_X_CURSOR            CursorType = 0
	GDK_ARROW               CursorType = 2
	GDK_BASED_ARROW_DOWN    CursorType = 4
	GDK_BASED_ARROW_UP      CursorType = 6
	GDK_BOAT                CursorType = 8
	GDK_BOGOSITY            CursorType = 10
	GDK_BOTTOM_LEFT_CORNER  CursorType = 12
	GDK_BOTTOM_RIGHT_CORNER CursorType = 14
	GDK_BOTTOM_SIDE         CursorType = 16
	GDK_BOTTOM_TEE          CursorType = 18
	GDK_BOX_SPIRAL          CursorType = 20
	GDK_CENTER_PTR          CursorType = 22
	GDK_CIRCLE              CursorType = 24
	GDK_CLOCK               CursorType = 26
	GDK_COFFEE_MUG          CursorType = 28
	GDK_CROSS               CursorType = 30
	GDK_CROSS_REVERSE       CursorType = 32
	GDK_CROSSHAIR           CursorType = 34
	GDK_DIAMOND_CROSS       CursorType = 36
	GDK_DOT                 CursorType = 38
	GDK_DOTBOX              CursorType = 40
	GDK_DOUBLE_ARROW        CursorType = 42
	GDK_DRAFT_LARGE         CursorType = 44
	GDK_DRAFT_SMALL         CursorType = 46
	GDK_DRAPED_BOX          CursorType = 48
	GDK_EXCHANGE            CursorType = 50
	GDK_FLEUR               CursorType = 52
	GDK_GOBBLER             CursorType = 54
	GDK_GUMBY               CursorType = 56
	GDK_HAND1               CursorType = 58
	GDK_HAND2               CursorType = 60
	GDK_HEART               CursorType = 62
	GDK_ICON                CursorType = 64
	GDK_IRON_CROSS          CursorType = 66
	GDK_LEFT_PTR            CursorType = 68
	GDK_LEFT_SIDE           CursorType = 70
	GDK_LEFT_TEE            CursorType = 72
	GDK_LEFTBUTTON          CursorType = 74
	GDK_LL_ANGLE            CursorType = 76
	GDK_LR_ANGLE            CursorType = 78
	GDK_MAN                 CursorType = 80
	GDK_MIDDLEBUTTON        CursorType = 82
	GDK_MOUSE               CursorType = 84
	GDK_PENCIL              CursorType = 86
	GDK_PIRATE              CursorType = 88
	GDK_PLUS                CursorType = 90
	GDK_QUESTION_ARROW      CursorType = 92
	GDK_RIGHT_PTR           CursorType = 94
	GDK_RIGHT_SIDE          CursorType = 96
	GDK_RIGHT_TEE           CursorType = 98
	GDK_RIGHTBUTTON         CursorType = 100
	GDK_RTL_LOGO            CursorType = 102
	GDK_SAILBOAT            CursorType = 104
	GDK_SB_DOWN_ARROW       CursorType = 106
	GDK_SB_H_DOUBLE_ARROW   CursorType = 108
	GDK_SB_LEFT_ARROW       CursorType = 110
	GDK_SB_RIGHT_ARROW      CursorType = 112
	GDK_SB_UP_ARROW         CursorType = 114
	GDK_SB_V_DOUBLE_ARROW   CursorType = 116
	GDK_SHUTTLE             CursorType = 118
	GDK_SIZING              CursorType = 120
	GDK_SPIDER              CursorType = 122
	GDK_SPRAYCAN            CursorType = 124
	GDK_STAR                CursorType = 126
	GDK_TARGET              CursorType = 128
	GDK_TCROSS              CursorType = 130
	GDK_TOP_LEFT_ARROW      CursorType = 132
	GDK_TOP_LEFT_CORNER     CursorType = 134
	GDK_TOP_RIGHT_CORNER    CursorType = 136
	GDK_TOP_SIDE            CursorType = 138
	GDK_TOP_TEE             CursorType = 140
	GDK_TREK                CursorType = 142
	GDK_UL_ANGLE            CursorType = 144
	GDK_UMBRELLA            CursorType = 146
	GDK_UR_ANGLE            CursorType = 148
	GDK_WATCH               CursorType = 150
	GDK_XTERM               CursorType = 152
	GDK_LAST_CURSOR         CursorType = 153
	GDK_BLANK_CURSOR        CursorType = -2
	GDK_CURSOR_IS_PIXMAP    CursorType = -1
)

type DeviceType C.GdkDeviceType

const (
	GDK_DEVICE_TYPE_MASTER   DeviceType = 0
	GDK_DEVICE_TYPE_SLAVE    DeviceType = 1
	GDK_DEVICE_TYPE_FLOATING DeviceType = 2
)

type DragProtocol C.GdkDragProtocol

const (
	GDK_DRAG_PROTO_NONE            DragProtocol = 0
	GDK_DRAG_PROTO_MOTIF           DragProtocol = 1
	GDK_DRAG_PROTO_XDND            DragProtocol = 2
	GDK_DRAG_PROTO_ROOTWIN         DragProtocol = 3
	GDK_DRAG_PROTO_WIN32_DROPFILES DragProtocol = 4
	GDK_DRAG_PROTO_OLE2            DragProtocol = 5
	GDK_DRAG_PROTO_LOCAL           DragProtocol = 6
	GDK_DRAG_PROTO_WAYLAND         DragProtocol = 7
)

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

type FilterReturn C.GdkFilterReturn

const (
	GDK_FILTER_CONTINUE  FilterReturn = 0
	GDK_FILTER_TRANSLATE FilterReturn = 1
	GDK_FILTER_REMOVE    FilterReturn = 2
)

type GrabOwnership C.GdkGrabOwnership

const (
	GDK_OWNERSHIP_NONE        GrabOwnership = 0
	GDK_OWNERSHIP_WINDOW      GrabOwnership = 1
	GDK_OWNERSHIP_APPLICATION GrabOwnership = 2
)

type GrabStatus C.GdkGrabStatus

const (
	GDK_GRAB_SUCCESS         GrabStatus = 0
	GDK_GRAB_ALREADY_GRABBED GrabStatus = 1
	GDK_GRAB_INVALID_TIME    GrabStatus = 2
	GDK_GRAB_NOT_VIEWABLE    GrabStatus = 3
	GDK_GRAB_FROZEN          GrabStatus = 4
	GDK_GRAB_FAILED          GrabStatus = 5
)

type Gravity C.GdkGravity

const (
	GDK_GRAVITY_NORTH_WEST Gravity = 1
	GDK_GRAVITY_NORTH      Gravity = 2
	GDK_GRAVITY_NORTH_EAST Gravity = 3
	GDK_GRAVITY_WEST       Gravity = 4
	GDK_GRAVITY_CENTER     Gravity = 5
	GDK_GRAVITY_EAST       Gravity = 6
	GDK_GRAVITY_SOUTH_WEST Gravity = 7
	GDK_GRAVITY_SOUTH      Gravity = 8
	GDK_GRAVITY_SOUTH_EAST Gravity = 9
	GDK_GRAVITY_STATIC     Gravity = 10
)

type InputMode C.GdkInputMode

const (
	GDK_MODE_DISABLED InputMode = 0
	GDK_MODE_SCREEN   InputMode = 1
	GDK_MODE_WINDOW   InputMode = 2
)

type InputSource C.GdkInputSource

const (
	GDK_SOURCE_MOUSE       InputSource = 0
	GDK_SOURCE_PEN         InputSource = 1
	GDK_SOURCE_ERASER      InputSource = 2
	GDK_SOURCE_CURSOR      InputSource = 3
	GDK_SOURCE_KEYBOARD    InputSource = 4
	GDK_SOURCE_TOUCHSCREEN InputSource = 5
	GDK_SOURCE_TOUCHPAD    InputSource = 6
	GDK_SOURCE_TRACKPOINT  InputSource = 7
	GDK_SOURCE_TABLET_PAD  InputSource = 8
)

type NotifyType C.GdkNotifyType

const (
	GDK_NOTIFY_ANCESTOR          NotifyType = 0
	GDK_NOTIFY_VIRTUAL           NotifyType = 1
	GDK_NOTIFY_INFERIOR          NotifyType = 2
	GDK_NOTIFY_NONLINEAR         NotifyType = 3
	GDK_NOTIFY_NONLINEAR_VIRTUAL NotifyType = 4
	GDK_NOTIFY_UNKNOWN           NotifyType = 5
)

type OwnerChange C.GdkOwnerChange

const (
	GDK_OWNER_CHANGE_NEW_OWNER OwnerChange = 0
	GDK_OWNER_CHANGE_DESTROY   OwnerChange = 1
	GDK_OWNER_CHANGE_CLOSE     OwnerChange = 2
)

type PropMode C.GdkPropMode

const (
	GDK_PROP_MODE_REPLACE PropMode = 0
	GDK_PROP_MODE_PREPEND PropMode = 1
	GDK_PROP_MODE_APPEND  PropMode = 2
)

type PropertyState C.guint

const (
	GDK_PROPERTY_NEW_VALUE PropertyState = 0
	GDK_PROPERTY_DELETE    PropertyState = 1
)

type ScrollDirection C.GdkScrollDirection

const (
	GDK_SCROLL_UP     ScrollDirection = 0
	GDK_SCROLL_DOWN   ScrollDirection = 1
	GDK_SCROLL_LEFT   ScrollDirection = 2
	GDK_SCROLL_RIGHT  ScrollDirection = 3
	GDK_SCROLL_SMOOTH ScrollDirection = 4
)

type SettingAction C.GdkSettingAction

const (
	GDK_SETTING_ACTION_NEW     SettingAction = 0
	GDK_SETTING_ACTION_CHANGED SettingAction = 1
	GDK_SETTING_ACTION_DELETED SettingAction = 2
)

type Status C.GdkStatus

const (
	GDK_OK          Status = 0
	GDK_ERROR       Status = -1
	GDK_ERROR_PARAM Status = -2
	GDK_ERROR_FILE  Status = -3
	GDK_ERROR_MEM   Status = -4
)

type TouchpadGesturePhase C.GdkTouchpadGesturePhase

const (
	GDK_TOUCHPAD_GESTURE_PHASE_BEGIN  TouchpadGesturePhase = 0
	GDK_TOUCHPAD_GESTURE_PHASE_UPDATE TouchpadGesturePhase = 1
	GDK_TOUCHPAD_GESTURE_PHASE_END    TouchpadGesturePhase = 2
	GDK_TOUCHPAD_GESTURE_PHASE_CANCEL TouchpadGesturePhase = 3
)

type VisibilityState C.GdkVisibilityState

const (
	GDK_VISIBILITY_UNOBSCURED     VisibilityState = 0
	GDK_VISIBILITY_PARTIAL        VisibilityState = 1
	GDK_VISIBILITY_FULLY_OBSCURED VisibilityState = 2
)

type VisualType C.GdkVisualType

const (
	GDK_VISUAL_STATIC_GRAY  VisualType = 0
	GDK_VISUAL_GRAYSCALE    VisualType = 1
	GDK_VISUAL_STATIC_COLOR VisualType = 2
	GDK_VISUAL_PSEUDO_COLOR VisualType = 3
	GDK_VISUAL_TRUE_COLOR   VisualType = 4
	GDK_VISUAL_DIRECT_COLOR VisualType = 5
)

type WindowEdge C.GdkWindowEdge

const (
	GDK_WINDOW_EDGE_NORTH_WEST WindowEdge = 0
	GDK_WINDOW_EDGE_NORTH      WindowEdge = 1
	GDK_WINDOW_EDGE_NORTH_EAST WindowEdge = 2
	GDK_WINDOW_EDGE_WEST       WindowEdge = 3
	GDK_WINDOW_EDGE_EAST       WindowEdge = 4
	GDK_WINDOW_EDGE_SOUTH_WEST WindowEdge = 5
	GDK_WINDOW_EDGE_SOUTH      WindowEdge = 6
	GDK_WINDOW_EDGE_SOUTH_EAST WindowEdge = 7
)

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
