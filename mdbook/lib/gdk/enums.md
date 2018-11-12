# `gdk` enums

## `AxisUse`

An enumeration describing the way in which a device
axis (valuator) maps onto the predefined valuator
types that GTK+ understands.

Note that the X and Y axes are not really needed; pointer devices
report their location via the x/y members of events regardless. Whether
X and Y are present as axes depends on the GDK backend.

- GDK_AXIS_IGNORE = 0
- GDK_AXIS_X = 1
- GDK_AXIS_Y = 2
- GDK_AXIS_PRESSURE = 3
- GDK_AXIS_XTILT = 4
- GDK_AXIS_YTILT = 5
- GDK_AXIS_WHEEL = 6
- GDK_AXIS_DISTANCE = 7
- GDK_AXIS_ROTATION = 8
- GDK_AXIS_SLIDER = 9
- GDK_AXIS_LAST = 10

C - `GdkAxisUse`

## `ByteOrder`

A set of values describing the possible byte-orders
for storing pixel values in memory.

- GDK_LSB_FIRST = 0
- GDK_MSB_FIRST = 1

C - `GdkByteOrder`

## `CrossingMode`

Specifies the crossing mode for #GdkEventCrossing.

- GDK_CROSSING_NORMAL = 0
- GDK_CROSSING_GRAB = 1
- GDK_CROSSING_UNGRAB = 2
- GDK_CROSSING_GTK_GRAB = 3
- GDK_CROSSING_GTK_UNGRAB = 4
- GDK_CROSSING_STATE_CHANGED = 5
- GDK_CROSSING_TOUCH_BEGIN = 6
- GDK_CROSSING_TOUCH_END = 7
- GDK_CROSSING_DEVICE_SWITCH = 8

C - `GdkCrossingMode`

## `CursorType`

Predefined cursors.

Note that these IDs are directly taken from the X cursor font, and many
of these cursors are either not useful, or are not available on other platforms.

The recommended way to create cursors is to use gdk_cursor_new_from_name().

- GDK_X_CURSOR = 0
- GDK_ARROW = 2
- GDK_BASED_ARROW_DOWN = 4
- GDK_BASED_ARROW_UP = 6
- GDK_BOAT = 8
- GDK_BOGOSITY = 10
- GDK_BOTTOM_LEFT_CORNER = 12
- GDK_BOTTOM_RIGHT_CORNER = 14
- GDK_BOTTOM_SIDE = 16
- GDK_BOTTOM_TEE = 18
- GDK_BOX_SPIRAL = 20
- GDK_CENTER_PTR = 22
- GDK_CIRCLE = 24
- GDK_CLOCK = 26
- GDK_COFFEE_MUG = 28
- GDK_CROSS = 30
- GDK_CROSS_REVERSE = 32
- GDK_CROSSHAIR = 34
- GDK_DIAMOND_CROSS = 36
- GDK_DOT = 38
- GDK_DOTBOX = 40
- GDK_DOUBLE_ARROW = 42
- GDK_DRAFT_LARGE = 44
- GDK_DRAFT_SMALL = 46
- GDK_DRAPED_BOX = 48
- GDK_EXCHANGE = 50
- GDK_FLEUR = 52
- GDK_GOBBLER = 54
- GDK_GUMBY = 56
- GDK_HAND1 = 58
- GDK_HAND2 = 60
- GDK_HEART = 62
- GDK_ICON = 64
- GDK_IRON_CROSS = 66
- GDK_LEFT_PTR = 68
- GDK_LEFT_SIDE = 70
- GDK_LEFT_TEE = 72
- GDK_LEFTBUTTON = 74
- GDK_LL_ANGLE = 76
- GDK_LR_ANGLE = 78
- GDK_MAN = 80
- GDK_MIDDLEBUTTON = 82
- GDK_MOUSE = 84
- GDK_PENCIL = 86
- GDK_PIRATE = 88
- GDK_PLUS = 90
- GDK_QUESTION_ARROW = 92
- GDK_RIGHT_PTR = 94
- GDK_RIGHT_SIDE = 96
- GDK_RIGHT_TEE = 98
- GDK_RIGHTBUTTON = 100
- GDK_RTL_LOGO = 102
- GDK_SAILBOAT = 104
- GDK_SB_DOWN_ARROW = 106
- GDK_SB_H_DOUBLE_ARROW = 108
- GDK_SB_LEFT_ARROW = 110
- GDK_SB_RIGHT_ARROW = 112
- GDK_SB_UP_ARROW = 114
- GDK_SB_V_DOUBLE_ARROW = 116
- GDK_SHUTTLE = 118
- GDK_SIZING = 120
- GDK_SPIDER = 122
- GDK_SPRAYCAN = 124
- GDK_STAR = 126
- GDK_TARGET = 128
- GDK_TCROSS = 130
- GDK_TOP_LEFT_ARROW = 132
- GDK_TOP_LEFT_CORNER = 134
- GDK_TOP_RIGHT_CORNER = 136
- GDK_TOP_SIDE = 138
- GDK_TOP_TEE = 140
- GDK_TREK = 142
- GDK_UL_ANGLE = 144
- GDK_UMBRELLA = 146
- GDK_UR_ANGLE = 148
- GDK_WATCH = 150
- GDK_XTERM = 152
- GDK_LAST_CURSOR = 153
- GDK_BLANK_CURSOR = -2
- GDK_CURSOR_IS_PIXMAP = -1

C - `GdkCursorType`

## `DevicePadFeature`

A pad feature.

- GDK_DEVICE_PAD_FEATURE_BUTTON = 0
- GDK_DEVICE_PAD_FEATURE_RING = 1
- GDK_DEVICE_PAD_FEATURE_STRIP = 2

C - `GdkDevicePadFeature`

## `DeviceToolType`

Indicates the specific type of tool being used being a tablet. Such as an
airbrush, pencil, etc.

- GDK_DEVICE_TOOL_TYPE_UNKNOWN = 0
- GDK_DEVICE_TOOL_TYPE_PEN = 1
- GDK_DEVICE_TOOL_TYPE_ERASER = 2
- GDK_DEVICE_TOOL_TYPE_BRUSH = 3
- GDK_DEVICE_TOOL_TYPE_PENCIL = 4
- GDK_DEVICE_TOOL_TYPE_AIRBRUSH = 5
- GDK_DEVICE_TOOL_TYPE_MOUSE = 6
- GDK_DEVICE_TOOL_TYPE_LENS = 7

C - `GdkDeviceToolType`

## `DeviceType`

Indicates the device type. See [above][GdkDeviceManager.description]
for more information about the meaning of these device types.

- GDK_DEVICE_TYPE_MASTER = 0
- GDK_DEVICE_TYPE_SLAVE = 1
- GDK_DEVICE_TYPE_FLOATING = 2

C - `GdkDeviceType`

## `DragCancelReason`

Used in #GdkDragContext to the reason of a cancelled DND operation.

- GDK_DRAG_CANCEL_NO_TARGET = 0
- GDK_DRAG_CANCEL_USER_CANCELLED = 1
- GDK_DRAG_CANCEL_ERROR = 2

C - `GdkDragCancelReason`

## `DragProtocol`

Used in #GdkDragContext to indicate the protocol according to
which DND is done.

- GDK_DRAG_PROTO_NONE = 0
- GDK_DRAG_PROTO_MOTIF = 1
- GDK_DRAG_PROTO_XDND = 2
- GDK_DRAG_PROTO_ROOTWIN = 3
- GDK_DRAG_PROTO_WIN32_DROPFILES = 4
- GDK_DRAG_PROTO_OLE2 = 5
- GDK_DRAG_PROTO_LOCAL = 6
- GDK_DRAG_PROTO_WAYLAND = 7

C - `GdkDragProtocol`

## `EventType`

Specifies the type of the event.

Do not confuse these events with the signals that GTK+ widgets emit.
Although many of these events result in corresponding signals being emitted,
the events are often transformed or filtered along the way.

In some language bindings, the values %GDK_2BUTTON_PRESS and
%GDK_3BUTTON_PRESS would translate into something syntactically
invalid (eg `Gdk.EventType.2ButtonPress`, where a
symbol is not allowed to start with a number). In that case, the
aliases %GDK_DOUBLE_BUTTON_PRESS and %GDK_TRIPLE_BUTTON_PRESS can
be used instead.

- GDK_NOTHING = -1
- GDK_DELETE = 0
- GDK_DESTROY = 1
- GDK_EXPOSE = 2
- GDK_MOTION_NOTIFY = 3
- GDK_BUTTON_PRESS = 4
- GDK_2BUTTON_PRESS = 5
- GDK_DOUBLE_BUTTON_PRESS = 5
- GDK_3BUTTON_PRESS = 6
- GDK_TRIPLE_BUTTON_PRESS = 6
- GDK_BUTTON_RELEASE = 7
- GDK_KEY_PRESS = 8
- GDK_KEY_RELEASE = 9
- GDK_ENTER_NOTIFY = 10
- GDK_LEAVE_NOTIFY = 11
- GDK_FOCUS_CHANGE = 12
- GDK_CONFIGURE = 13
- GDK_MAP = 14
- GDK_UNMAP = 15
- GDK_PROPERTY_NOTIFY = 16
- GDK_SELECTION_CLEAR = 17
- GDK_SELECTION_REQUEST = 18
- GDK_SELECTION_NOTIFY = 19
- GDK_PROXIMITY_IN = 20
- GDK_PROXIMITY_OUT = 21
- GDK_DRAG_ENTER = 22
- GDK_DRAG_LEAVE = 23
- GDK_DRAG_MOTION = 24
- GDK_DRAG_STATUS = 25
- GDK_DROP_START = 26
- GDK_DROP_FINISHED = 27
- GDK_CLIENT_EVENT = 28
- GDK_VISIBILITY_NOTIFY = 29
- GDK_SCROLL = 31
- GDK_WINDOW_STATE = 32
- GDK_SETTING = 33
- GDK_OWNER_CHANGE = 34
- GDK_GRAB_BROKEN = 35
- GDK_DAMAGE = 36
- GDK_TOUCH_BEGIN = 37
- GDK_TOUCH_UPDATE = 38
- GDK_TOUCH_END = 39
- GDK_TOUCH_CANCEL = 40
- GDK_TOUCHPAD_SWIPE = 41
- GDK_TOUCHPAD_PINCH = 42
- GDK_PAD_BUTTON_PRESS = 43
- GDK_PAD_BUTTON_RELEASE = 44
- GDK_PAD_RING = 45
- GDK_PAD_STRIP = 46
- GDK_PAD_GROUP_MODE = 47
- GDK_EVENT_LAST = 48

C - `GdkEventType`

## `FilterReturn`

Specifies the result of applying a #GdkFilterFunc to a native event.

- GDK_FILTER_CONTINUE = 0
- GDK_FILTER_TRANSLATE = 1
- GDK_FILTER_REMOVE = 2

C - `GdkFilterReturn`

## `FullscreenMode`

Indicates which monitor (in a multi-head setup) a window should span over
when in fullscreen mode.

- GDK_FULLSCREEN_ON_CURRENT_MONITOR = 0
- GDK_FULLSCREEN_ON_ALL_MONITORS = 1

C - `GdkFullscreenMode`

## `GLError`

Error enumeration for #GdkGLContext.

- GDK_GL_ERROR_NOT_AVAILABLE = 0
- GDK_GL_ERROR_UNSUPPORTED_FORMAT = 1
- GDK_GL_ERROR_UNSUPPORTED_PROFILE = 2

C - `GdkGLError`

## `GrabOwnership`

Defines how device grabs interact with other devices.

- GDK_OWNERSHIP_NONE = 0
- GDK_OWNERSHIP_WINDOW = 1
- GDK_OWNERSHIP_APPLICATION = 2

C - `GdkGrabOwnership`

## `GrabStatus`

Returned by gdk_device_grab(), gdk_pointer_grab() and gdk_keyboard_grab() to
indicate success or the reason for the failure of the grab attempt.

- GDK_GRAB_SUCCESS = 0
- GDK_GRAB_ALREADY_GRABBED = 1
- GDK_GRAB_INVALID_TIME = 2
- GDK_GRAB_NOT_VIEWABLE = 3
- GDK_GRAB_FROZEN = 4
- GDK_GRAB_FAILED = 5

C - `GdkGrabStatus`

## `Gravity`

Defines the reference point of a window and the meaning of coordinates
passed to gtk_window_move(). See gtk_window_move() and the "implementation
notes" section of the
[Extended Window Manager Hints](http://www.freedesktop.org/Standards/wm-spec)
specification for more details.

- GDK_GRAVITY_NORTH_WEST = 1
- GDK_GRAVITY_NORTH = 2
- GDK_GRAVITY_NORTH_EAST = 3
- GDK_GRAVITY_WEST = 4
- GDK_GRAVITY_CENTER = 5
- GDK_GRAVITY_EAST = 6
- GDK_GRAVITY_SOUTH_WEST = 7
- GDK_GRAVITY_SOUTH = 8
- GDK_GRAVITY_SOUTH_EAST = 9
- GDK_GRAVITY_STATIC = 10

C - `GdkGravity`

## `InputMode`

An enumeration that describes the mode of an input device.

- GDK_MODE_DISABLED = 0
- GDK_MODE_SCREEN = 1
- GDK_MODE_WINDOW = 2

C - `GdkInputMode`

## `InputSource`

An enumeration describing the type of an input device in general terms.

- GDK_SOURCE_MOUSE = 0
- GDK_SOURCE_PEN = 1
- GDK_SOURCE_ERASER = 2
- GDK_SOURCE_CURSOR = 3
- GDK_SOURCE_KEYBOARD = 4
- GDK_SOURCE_TOUCHSCREEN = 5
- GDK_SOURCE_TOUCHPAD = 6
- GDK_SOURCE_TRACKPOINT = 7
- GDK_SOURCE_TABLET_PAD = 8

C - `GdkInputSource`

## `ModifierIntent`

This enum is used with gdk_keymap_get_modifier_mask()
in order to determine what modifiers the
currently used windowing system backend uses for particular
purposes. For example, on X11/Windows, the Control key is used for
invoking menu shortcuts (accelerators), whereas on Apple computers
it’s the Command key (which correspond to %GDK_CONTROL_MASK and
%GDK_MOD2_MASK, respectively).

- GDK_MODIFIER_INTENT_PRIMARY_ACCELERATOR = 0
- GDK_MODIFIER_INTENT_CONTEXT_MENU = 1
- GDK_MODIFIER_INTENT_EXTEND_SELECTION = 2
- GDK_MODIFIER_INTENT_MODIFY_SELECTION = 3
- GDK_MODIFIER_INTENT_NO_TEXT_INPUT = 4
- GDK_MODIFIER_INTENT_SHIFT_GROUP = 5
- GDK_MODIFIER_INTENT_DEFAULT_MOD_MASK = 6

C - `GdkModifierIntent`

## `NotifyType`

Specifies the kind of crossing for #GdkEventCrossing.

See the X11 protocol specification of LeaveNotify for
full details of crossing event generation.

- GDK_NOTIFY_ANCESTOR = 0
- GDK_NOTIFY_VIRTUAL = 1
- GDK_NOTIFY_INFERIOR = 2
- GDK_NOTIFY_NONLINEAR = 3
- GDK_NOTIFY_NONLINEAR_VIRTUAL = 4
- GDK_NOTIFY_UNKNOWN = 5

C - `GdkNotifyType`

## `OwnerChange`

Specifies why a selection ownership was changed.

- GDK_OWNER_CHANGE_NEW_OWNER = 0
- GDK_OWNER_CHANGE_DESTROY = 1
- GDK_OWNER_CHANGE_CLOSE = 2

C - `GdkOwnerChange`

## `PropMode`

Describes how existing data is combined with new data when
using gdk_property_change().

- GDK_PROP_MODE_REPLACE = 0
- GDK_PROP_MODE_PREPEND = 1
- GDK_PROP_MODE_APPEND = 2

C - `GdkPropMode`

## `PropertyState`

Specifies the type of a property change for a #GdkEventProperty.

- GDK_PROPERTY_NEW_VALUE = 0
- GDK_PROPERTY_DELETE = 1

C - `guint`

## `ScrollDirection`

Specifies the direction for #GdkEventScroll.

- GDK_SCROLL_UP = 0
- GDK_SCROLL_DOWN = 1
- GDK_SCROLL_LEFT = 2
- GDK_SCROLL_RIGHT = 3
- GDK_SCROLL_SMOOTH = 4

C - `GdkScrollDirection`

## `SettingAction`

Specifies the kind of modification applied to a setting in a
&num;GdkEventSetting.

- GDK_SETTING_ACTION_NEW = 0
- GDK_SETTING_ACTION_CHANGED = 1
- GDK_SETTING_ACTION_DELETED = 2

C - `GdkSettingAction`

## `Status`



- GDK_OK = 0
- GDK_ERROR = -1
- GDK_ERROR_PARAM = -2
- GDK_ERROR_FILE = -3
- GDK_ERROR_MEM = -4

C - `GdkStatus`

## `SubpixelLayout`

This enumeration describes how the red, green and blue components
of physical pixels on an output device are laid out.

- GDK_SUBPIXEL_LAYOUT_UNKNOWN = 0
- GDK_SUBPIXEL_LAYOUT_NONE = 1
- GDK_SUBPIXEL_LAYOUT_HORIZONTAL_RGB = 2
- GDK_SUBPIXEL_LAYOUT_HORIZONTAL_BGR = 3
- GDK_SUBPIXEL_LAYOUT_VERTICAL_RGB = 4
- GDK_SUBPIXEL_LAYOUT_VERTICAL_BGR = 5

C - `GdkSubpixelLayout`

## `TouchpadGesturePhase`

Specifies the current state of a touchpad gesture. All gestures are
guaranteed to begin with an event with phase %GDK_TOUCHPAD_GESTURE_PHASE_BEGIN,
followed by 0 or several events with phase %GDK_TOUCHPAD_GESTURE_PHASE_UPDATE.

A finished gesture may have 2 possible outcomes, an event with phase
%GDK_TOUCHPAD_GESTURE_PHASE_END will be emitted when the gesture is
considered successful, this should be used as the hint to perform any
permanent changes.

Cancelled gestures may be so for a variety of reasons, due to hardware
or the compositor, or due to the gesture recognition layers hinting the
gesture did not finish resolutely (eg. a 3rd finger being added during
a pinch gesture). In these cases, the last event will report the phase
%GDK_TOUCHPAD_GESTURE_PHASE_CANCEL, this should be used as a hint
to undo any visible/permanent changes that were done throughout the
progress of the gesture.

See also #GdkEventTouchpadSwipe and #GdkEventTouchpadPinch.

- GDK_TOUCHPAD_GESTURE_PHASE_BEGIN = 0
- GDK_TOUCHPAD_GESTURE_PHASE_UPDATE = 1
- GDK_TOUCHPAD_GESTURE_PHASE_END = 2
- GDK_TOUCHPAD_GESTURE_PHASE_CANCEL = 3

C - `GdkTouchpadGesturePhase`

## `VisibilityState`

Specifies the visiblity status of a window for a #GdkEventVisibility.

- GDK_VISIBILITY_UNOBSCURED = 0
- GDK_VISIBILITY_PARTIAL = 1
- GDK_VISIBILITY_FULLY_OBSCURED = 2

C - `GdkVisibilityState`

## `VisualType`

A set of values that describe the manner in which the pixel values
for a visual are converted into RGB values for display.

- GDK_VISUAL_STATIC_GRAY = 0
- GDK_VISUAL_GRAYSCALE = 1
- GDK_VISUAL_STATIC_COLOR = 2
- GDK_VISUAL_PSEUDO_COLOR = 3
- GDK_VISUAL_TRUE_COLOR = 4
- GDK_VISUAL_DIRECT_COLOR = 5

C - `GdkVisualType`

## `WindowEdge`

Determines a window edge or corner.

- GDK_WINDOW_EDGE_NORTH_WEST = 0
- GDK_WINDOW_EDGE_NORTH = 1
- GDK_WINDOW_EDGE_NORTH_EAST = 2
- GDK_WINDOW_EDGE_WEST = 3
- GDK_WINDOW_EDGE_EAST = 4
- GDK_WINDOW_EDGE_SOUTH_WEST = 5
- GDK_WINDOW_EDGE_SOUTH = 6
- GDK_WINDOW_EDGE_SOUTH_EAST = 7

C - `GdkWindowEdge`

## `WindowType`

Describes the kind of window.

- GDK_WINDOW_ROOT = 0
- GDK_WINDOW_TOPLEVEL = 1
- GDK_WINDOW_CHILD = 2
- GDK_WINDOW_TEMP = 3
- GDK_WINDOW_FOREIGN = 4
- GDK_WINDOW_OFFSCREEN = 5
- GDK_WINDOW_SUBSURFACE = 6

C - `GdkWindowType`

## `WindowTypeHint`

These are hints for the window manager that indicate what type of function
the window has. The window manager can use this when determining decoration
and behaviour of the window. The hint must be set before mapping the window.

See the [Extended Window Manager Hints](http://www.freedesktop.org/Standards/wm-spec)
specification for more details about window types.

- GDK_WINDOW_TYPE_HINT_NORMAL = 0
- GDK_WINDOW_TYPE_HINT_DIALOG = 1
- GDK_WINDOW_TYPE_HINT_MENU = 2
- GDK_WINDOW_TYPE_HINT_TOOLBAR = 3
- GDK_WINDOW_TYPE_HINT_SPLASHSCREEN = 4
- GDK_WINDOW_TYPE_HINT_UTILITY = 5
- GDK_WINDOW_TYPE_HINT_DOCK = 6
- GDK_WINDOW_TYPE_HINT_DESKTOP = 7
- GDK_WINDOW_TYPE_HINT_DROPDOWN_MENU = 8
- GDK_WINDOW_TYPE_HINT_POPUP_MENU = 9
- GDK_WINDOW_TYPE_HINT_TOOLTIP = 10
- GDK_WINDOW_TYPE_HINT_NOTIFICATION = 11
- GDK_WINDOW_TYPE_HINT_COMBO = 12
- GDK_WINDOW_TYPE_HINT_DND = 13

C - `GdkWindowTypeHint`

## `WindowWindowClass`

@GDK_INPUT_OUTPUT windows are the standard kind of window you might expect.
Such windows receive events and are also displayed on screen.
@GDK_INPUT_ONLY windows are invisible; they are usually placed above other
windows in order to trap or filter the events. You can’t draw on
@GDK_INPUT_ONLY windows.

- GDK_INPUT_OUTPUT = 0
- GDK_INPUT_ONLY = 1

C - `GdkWindowWindowClass`

yState`

- GDK_VISIBILITY_UNOBSCURED = %!s(int=0)
- GDK_VISIBILITY_PARTIAL = %!s(int=1)
- GDK_VISIBILITY_FULLY_OBSCURED = %!s(int=2)
## `VisualType`

A set of values that describe the manner in which the pixel values
for a visual are converted into RGB values for display.

C - `GdkVisualType`

- GDK_VISUAL_STATIC_GRAY = %!s(int=0)
- GDK_VISUAL_GRAYSCALE = %!s(int=1)
- GDK_VISUAL_STATIC_COLOR = %!s(int=2)
- GDK_VISUAL_PSEUDO_COLOR = %!s(int=3)
- GDK_VISUAL_TRUE_COLOR = %!s(int=4)
- GDK_VISUAL_DIRECT_COLOR = %!s(int=5)
## `WindowEdge`

Determines a window edge or corner.

C - `GdkWindowEdge`

- GDK_WINDOW_EDGE_NORTH_WEST = %!s(int=0)
- GDK_WINDOW_EDGE_NORTH = %!s(int=1)
- GDK_WINDOW_EDGE_NORTH_EAST = %!s(int=2)
- GDK_WINDOW_EDGE_WEST = %!s(int=3)
- GDK_WINDOW_EDGE_EAST = %!s(int=4)
- GDK_WINDOW_EDGE_SOUTH_WEST = %!s(int=5)
- GDK_WINDOW_EDGE_SOUTH = %!s(int=6)
- GDK_WINDOW_EDGE_SOUTH_EAST = %!s(int=7)
## `WindowType`

Describes the kind of window.

C - `GdkWindowType`

- GDK_WINDOW_ROOT = %!s(int=0)
- GDK_WINDOW_TOPLEVEL = %!s(int=1)
- GDK_WINDOW_CHILD = %!s(int=2)
- GDK_WINDOW_TEMP = %!s(int=3)
- GDK_WINDOW_FOREIGN = %!s(int=4)
- GDK_WINDOW_OFFSCREEN = %!s(int=5)
- GDK_WINDOW_SUBSURFACE = %!s(int=6)
## `WindowTypeHint`

These are hints for the window manager that indicate what type of function
the window has. The window manager can use this when determining decoration
and behaviour of the window. The hint must be set before mapping the window.

See the [Extended Window Manager Hints](http://www.freedesktop.org/Standards/wm-spec)
specification for more details about window types.

C - `GdkWindowTypeHint`

- GDK_WINDOW_TYPE_HINT_NORMAL = %!s(int=0)
- GDK_WINDOW_TYPE_HINT_DIALOG = %!s(int=1)
- GDK_WINDOW_TYPE_HINT_MENU = %!s(int=2)
- GDK_WINDOW_TYPE_HINT_TOOLBAR = %!s(int=3)
- GDK_WINDOW_TYPE_HINT_SPLASHSCREEN = %!s(int=4)
- GDK_WINDOW_TYPE_HINT_UTILITY = %!s(int=5)
- GDK_WINDOW_TYPE_HINT_DOCK = %!s(int=6)
- GDK_WINDOW_TYPE_HINT_DESKTOP = %!s(int=7)
- GDK_WINDOW_TYPE_HINT_DROPDOWN_MENU = %!s(int=8)
- GDK_WINDOW_TYPE_HINT_POPUP_MENU = %!s(int=9)
- GDK_WINDOW_TYPE_HINT_TOOLTIP = %!s(int=10)
- GDK_WINDOW_TYPE_HINT_NOTIFICATION = %!s(int=11)
- GDK_WINDOW_TYPE_HINT_COMBO = %!s(int=12)
- GDK_WINDOW_TYPE_HINT_DND = %!s(int=13)
## `WindowWindowClass`

@GDK_INPUT_OUTPUT windows are the standard kind of window you might expect.
Such windows receive events and are also displayed on screen.
@GDK_INPUT_ONLY windows are invisible; they are usually placed above other
windows in order to trap or filter the events. You can’t draw on
@GDK_INPUT_ONLY windows.

C - `GdkWindowWindowClass`

- GDK_INPUT_OUTPUT = %!s(int=0)
- GDK_INPUT_ONLY = %!s(int=1)
