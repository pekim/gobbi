// This is a generated file - DO NOT EDIT

package gdk

import "C"

type DragAction C.GdkDragAction

const (
	GDK_ACTION_DEFAULT DragAction = 1
	GDK_ACTION_COPY    DragAction = 2
	GDK_ACTION_MOVE    DragAction = 4
	GDK_ACTION_LINK    DragAction = 8
	GDK_ACTION_PRIVATE DragAction = 16
	GDK_ACTION_ASK     DragAction = 32
)

type EventMask C.GdkEventMask

const (
	GDK_EXPOSURE_MASK            EventMask = 2
	GDK_POINTER_MOTION_MASK      EventMask = 4
	GDK_POINTER_MOTION_HINT_MASK EventMask = 8
	GDK_BUTTON_MOTION_MASK       EventMask = 16
	GDK_BUTTON1_MOTION_MASK      EventMask = 32
	GDK_BUTTON2_MOTION_MASK      EventMask = 64
	GDK_BUTTON3_MOTION_MASK      EventMask = 128
	GDK_BUTTON_PRESS_MASK        EventMask = 256
	GDK_BUTTON_RELEASE_MASK      EventMask = 512
	GDK_KEY_PRESS_MASK           EventMask = 1024
	GDK_KEY_RELEASE_MASK         EventMask = 2048
	GDK_ENTER_NOTIFY_MASK        EventMask = 4096
	GDK_LEAVE_NOTIFY_MASK        EventMask = 8192
	GDK_FOCUS_CHANGE_MASK        EventMask = 16384
	GDK_STRUCTURE_MASK           EventMask = 32768
	GDK_PROPERTY_CHANGE_MASK     EventMask = 65536
	GDK_VISIBILITY_NOTIFY_MASK   EventMask = 131072
	GDK_PROXIMITY_IN_MASK        EventMask = 262144
	GDK_PROXIMITY_OUT_MASK       EventMask = 524288
	GDK_SUBSTRUCTURE_MASK        EventMask = 1048576
	GDK_SCROLL_MASK              EventMask = 2097152
	GDK_TOUCH_MASK               EventMask = 4194304
	GDK_SMOOTH_SCROLL_MASK       EventMask = 8388608
	GDK_TOUCHPAD_GESTURE_MASK    EventMask = 16777216
	GDK_TABLET_PAD_MASK          EventMask = 33554432
	GDK_ALL_EVENTS_MASK          EventMask = 67108862
)

type ModifierType C.guint

const (
	GDK_SHIFT_MASK                ModifierType = 1
	GDK_LOCK_MASK                 ModifierType = 2
	GDK_CONTROL_MASK              ModifierType = 4
	GDK_MOD1_MASK                 ModifierType = 8
	GDK_MOD2_MASK                 ModifierType = 16
	GDK_MOD3_MASK                 ModifierType = 32
	GDK_MOD4_MASK                 ModifierType = 64
	GDK_MOD5_MASK                 ModifierType = 128
	GDK_BUTTON1_MASK              ModifierType = 256
	GDK_BUTTON2_MASK              ModifierType = 512
	GDK_BUTTON3_MASK              ModifierType = 1024
	GDK_BUTTON4_MASK              ModifierType = 2048
	GDK_BUTTON5_MASK              ModifierType = 4096
	GDK_MODIFIER_RESERVED_13_MASK ModifierType = 8192
	GDK_MODIFIER_RESERVED_14_MASK ModifierType = 16384
	GDK_MODIFIER_RESERVED_15_MASK ModifierType = 32768
	GDK_MODIFIER_RESERVED_16_MASK ModifierType = 65536
	GDK_MODIFIER_RESERVED_17_MASK ModifierType = 131072
	GDK_MODIFIER_RESERVED_18_MASK ModifierType = 262144
	GDK_MODIFIER_RESERVED_19_MASK ModifierType = 524288
	GDK_MODIFIER_RESERVED_20_MASK ModifierType = 1048576
	GDK_MODIFIER_RESERVED_21_MASK ModifierType = 2097152
	GDK_MODIFIER_RESERVED_22_MASK ModifierType = 4194304
	GDK_MODIFIER_RESERVED_23_MASK ModifierType = 8388608
	GDK_MODIFIER_RESERVED_24_MASK ModifierType = 16777216
	GDK_MODIFIER_RESERVED_25_MASK ModifierType = 33554432
	GDK_SUPER_MASK                ModifierType = 67108864
	GDK_HYPER_MASK                ModifierType = 134217728
	GDK_META_MASK                 ModifierType = 268435456
	GDK_MODIFIER_RESERVED_29_MASK ModifierType = 536870912
	GDK_RELEASE_MASK              ModifierType = 1073741824
	GDK_MODIFIER_MASK             ModifierType = 1543512063
)

type WMDecoration C.GdkWMDecoration

const (
	GDK_DECOR_ALL      WMDecoration = 1
	GDK_DECOR_BORDER   WMDecoration = 2
	GDK_DECOR_RESIZEH  WMDecoration = 4
	GDK_DECOR_TITLE    WMDecoration = 8
	GDK_DECOR_MENU     WMDecoration = 16
	GDK_DECOR_MINIMIZE WMDecoration = 32
	GDK_DECOR_MAXIMIZE WMDecoration = 64
)

type WMFunction C.GdkWMFunction

const (
	GDK_FUNC_ALL      WMFunction = 1
	GDK_FUNC_RESIZE   WMFunction = 2
	GDK_FUNC_MOVE     WMFunction = 4
	GDK_FUNC_MINIMIZE WMFunction = 8
	GDK_FUNC_MAXIMIZE WMFunction = 16
	GDK_FUNC_CLOSE    WMFunction = 32
)

type WindowAttributesType C.GdkWindowAttributesType

const (
	GDK_WA_TITLE     WindowAttributesType = 2
	GDK_WA_X         WindowAttributesType = 4
	GDK_WA_Y         WindowAttributesType = 8
	GDK_WA_CURSOR    WindowAttributesType = 16
	GDK_WA_VISUAL    WindowAttributesType = 32
	GDK_WA_WMCLASS   WindowAttributesType = 64
	GDK_WA_NOREDIR   WindowAttributesType = 128
	GDK_WA_TYPE_HINT WindowAttributesType = 256
)

type WindowHints C.GdkWindowHints

const (
	GDK_HINT_POS         WindowHints = 1
	GDK_HINT_MIN_SIZE    WindowHints = 2
	GDK_HINT_MAX_SIZE    WindowHints = 4
	GDK_HINT_BASE_SIZE   WindowHints = 8
	GDK_HINT_ASPECT      WindowHints = 16
	GDK_HINT_RESIZE_INC  WindowHints = 32
	GDK_HINT_WIN_GRAVITY WindowHints = 64
	GDK_HINT_USER_POS    WindowHints = 128
	GDK_HINT_USER_SIZE   WindowHints = 256
)

type WindowState C.GdkWindowState

const (
	GDK_WINDOW_STATE_WITHDRAWN        WindowState = 1
	GDK_WINDOW_STATE_ICONIFIED        WindowState = 2
	GDK_WINDOW_STATE_MAXIMIZED        WindowState = 4
	GDK_WINDOW_STATE_STICKY           WindowState = 8
	GDK_WINDOW_STATE_FULLSCREEN       WindowState = 16
	GDK_WINDOW_STATE_ABOVE            WindowState = 32
	GDK_WINDOW_STATE_BELOW            WindowState = 64
	GDK_WINDOW_STATE_FOCUSED          WindowState = 128
	GDK_WINDOW_STATE_TILED            WindowState = 256
	GDK_WINDOW_STATE_TOP_TILED        WindowState = 512
	GDK_WINDOW_STATE_TOP_RESIZABLE    WindowState = 1024
	GDK_WINDOW_STATE_RIGHT_TILED      WindowState = 2048
	GDK_WINDOW_STATE_RIGHT_RESIZABLE  WindowState = 4096
	GDK_WINDOW_STATE_BOTTOM_TILED     WindowState = 8192
	GDK_WINDOW_STATE_BOTTOM_RESIZABLE WindowState = 16384
	GDK_WINDOW_STATE_LEFT_TILED       WindowState = 32768
	GDK_WINDOW_STATE_LEFT_RESIZABLE   WindowState = 65536
)
