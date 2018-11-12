# `gdk` bitfields

## `AnchorHints`

Positioning hints for aligning a window relative to a rectangle.

These hints determine how the window should be positioned in the case that
the window would fall off-screen if placed in its ideal position.

For example, %GDK_ANCHOR_FLIP_X will replace %GDK_GRAVITY_NORTH_WEST with
%GDK_GRAVITY_NORTH_EAST and vice versa if the window extends beyond the left
or right edges of the monitor.

If %GDK_ANCHOR_SLIDE_X is set, the window can be shifted horizontally to fit
on-screen. If %GDK_ANCHOR_RESIZE_X is set, the window can be shrunken
horizontally to fit.

In general, when multiple flags are set, flipping should take precedence over
sliding, which should take precedence over resizing.

- GDK_ANCHOR_FLIP_X = 1
- GDK_ANCHOR_FLIP_Y = 2
- GDK_ANCHOR_SLIDE_X = 4
- GDK_ANCHOR_SLIDE_Y = 8
- GDK_ANCHOR_RESIZE_X = 16
- GDK_ANCHOR_RESIZE_Y = 32
- GDK_ANCHOR_FLIP = 3
- GDK_ANCHOR_SLIDE = 12
- GDK_ANCHOR_RESIZE = 48

C - `GdkAnchorHints`

## `AxisFlags`

Flags describing the current capabilities of a device/tool.

- GDK_AXIS_FLAG_X = 2
- GDK_AXIS_FLAG_Y = 4
- GDK_AXIS_FLAG_PRESSURE = 8
- GDK_AXIS_FLAG_XTILT = 16
- GDK_AXIS_FLAG_YTILT = 32
- GDK_AXIS_FLAG_WHEEL = 64
- GDK_AXIS_FLAG_DISTANCE = 128
- GDK_AXIS_FLAG_ROTATION = 256
- GDK_AXIS_FLAG_SLIDER = 512

C - `GdkAxisFlags`

## `DragAction`

Used in #GdkDragContext to indicate what the destination
should do with the dropped data.

- GDK_ACTION_DEFAULT = 1
- GDK_ACTION_COPY = 2
- GDK_ACTION_MOVE = 4
- GDK_ACTION_LINK = 8
- GDK_ACTION_PRIVATE = 16
- GDK_ACTION_ASK = 32

C - `GdkDragAction`

## `EventMask`

A set of bit-flags to indicate which events a window is to receive.
Most of these masks map onto one or more of the #GdkEventType event types
above.

See the [input handling overview][chap-input-handling] for details of
[event masks][event-masks] and [event propagation][event-propagation].

%GDK_POINTER_MOTION_HINT_MASK is deprecated. It is a special mask
to reduce the number of %GDK_MOTION_NOTIFY events received. When using
%GDK_POINTER_MOTION_HINT_MASK, fewer %GDK_MOTION_NOTIFY events will
be sent, some of which are marked as a hint (the is_hint member is
%TRUE). To receive more motion events after a motion hint event,
the application needs to asks for more, by calling
gdk_event_request_motions().

Since GTK 3.8, motion events are already compressed by default, independent
of this mechanism. This compression can be disabled with
gdk_window_set_event_compression(). See the documentation of that function
for details.

If %GDK_TOUCH_MASK is enabled, the window will receive touch events
from touch-enabled devices. Those will come as sequences of #GdkEventTouch
with type %GDK_TOUCH_UPDATE, enclosed by two events with
type %GDK_TOUCH_BEGIN and %GDK_TOUCH_END (or %GDK_TOUCH_CANCEL).
gdk_event_get_event_sequence() returns the event sequence for these
events, so different sequences may be distinguished.

- GDK_EXPOSURE_MASK = 2
- GDK_POINTER_MOTION_MASK = 4
- GDK_POINTER_MOTION_HINT_MASK = 8
- GDK_BUTTON_MOTION_MASK = 16
- GDK_BUTTON1_MOTION_MASK = 32
- GDK_BUTTON2_MOTION_MASK = 64
- GDK_BUTTON3_MOTION_MASK = 128
- GDK_BUTTON_PRESS_MASK = 256
- GDK_BUTTON_RELEASE_MASK = 512
- GDK_KEY_PRESS_MASK = 1024
- GDK_KEY_RELEASE_MASK = 2048
- GDK_ENTER_NOTIFY_MASK = 4096
- GDK_LEAVE_NOTIFY_MASK = 8192
- GDK_FOCUS_CHANGE_MASK = 16384
- GDK_STRUCTURE_MASK = 32768
- GDK_PROPERTY_CHANGE_MASK = 65536
- GDK_VISIBILITY_NOTIFY_MASK = 131072
- GDK_PROXIMITY_IN_MASK = 262144
- GDK_PROXIMITY_OUT_MASK = 524288
- GDK_SUBSTRUCTURE_MASK = 1048576
- GDK_SCROLL_MASK = 2097152
- GDK_TOUCH_MASK = 4194304
- GDK_SMOOTH_SCROLL_MASK = 8388608
- GDK_TOUCHPAD_GESTURE_MASK = 16777216
- GDK_TABLET_PAD_MASK = 33554432
- GDK_ALL_EVENTS_MASK = 67108862

C - `GdkEventMask`

## `FrameClockPhase`

#GdkFrameClockPhase is used to represent the different paint clock
phases that can be requested. The elements of the enumeration
correspond to the signals of #GdkFrameClock.

- GDK_FRAME_CLOCK_PHASE_NONE = 0
- GDK_FRAME_CLOCK_PHASE_FLUSH_EVENTS = 1
- GDK_FRAME_CLOCK_PHASE_BEFORE_PAINT = 2
- GDK_FRAME_CLOCK_PHASE_UPDATE = 4
- GDK_FRAME_CLOCK_PHASE_LAYOUT = 8
- GDK_FRAME_CLOCK_PHASE_PAINT = 16
- GDK_FRAME_CLOCK_PHASE_RESUME_EVENTS = 32
- GDK_FRAME_CLOCK_PHASE_AFTER_PAINT = 64

C - `GdkFrameClockPhase`

## `ModifierType`

A set of bit-flags to indicate the state of modifier keys and mouse buttons
in various event types. Typical modifier keys are Shift, Control, Meta,
Super, Hyper, Alt, Compose, Apple, CapsLock or ShiftLock.

Like the X Window System, GDK supports 8 modifier keys and 5 mouse buttons.

Since 2.10, GDK recognizes which of the Meta, Super or Hyper keys are mapped
to Mod2 - Mod5, and indicates this by setting %GDK_SUPER_MASK,
%GDK_HYPER_MASK or %GDK_META_MASK in the state field of key events.

Note that GDK may add internal values to events which include
reserved values such as %GDK_MODIFIER_RESERVED_13_MASK.  Your code
should preserve and ignore them.  You can use %GDK_MODIFIER_MASK to
remove all reserved values.

Also note that the GDK X backend interprets button press events for button
4-7 as scroll events, so %GDK_BUTTON4_MASK and %GDK_BUTTON5_MASK will never
be set.

- GDK_SHIFT_MASK = 1
- GDK_LOCK_MASK = 2
- GDK_CONTROL_MASK = 4
- GDK_MOD1_MASK = 8
- GDK_MOD2_MASK = 16
- GDK_MOD3_MASK = 32
- GDK_MOD4_MASK = 64
- GDK_MOD5_MASK = 128
- GDK_BUTTON1_MASK = 256
- GDK_BUTTON2_MASK = 512
- GDK_BUTTON3_MASK = 1024
- GDK_BUTTON4_MASK = 2048
- GDK_BUTTON5_MASK = 4096
- GDK_MODIFIER_RESERVED_13_MASK = 8192
- GDK_MODIFIER_RESERVED_14_MASK = 16384
- GDK_MODIFIER_RESERVED_15_MASK = 32768
- GDK_MODIFIER_RESERVED_16_MASK = 65536
- GDK_MODIFIER_RESERVED_17_MASK = 131072
- GDK_MODIFIER_RESERVED_18_MASK = 262144
- GDK_MODIFIER_RESERVED_19_MASK = 524288
- GDK_MODIFIER_RESERVED_20_MASK = 1048576
- GDK_MODIFIER_RESERVED_21_MASK = 2097152
- GDK_MODIFIER_RESERVED_22_MASK = 4194304
- GDK_MODIFIER_RESERVED_23_MASK = 8388608
- GDK_MODIFIER_RESERVED_24_MASK = 16777216
- GDK_MODIFIER_RESERVED_25_MASK = 33554432
- GDK_SUPER_MASK = 67108864
- GDK_HYPER_MASK = 134217728
- GDK_META_MASK = 268435456
- GDK_MODIFIER_RESERVED_29_MASK = 536870912
- GDK_RELEASE_MASK = 1073741824
- GDK_MODIFIER_MASK = 1543512063

C - `guint`

## `SeatCapabilities`

Flags describing the seat capabilities.

- GDK_SEAT_CAPABILITY_NONE = 0
- GDK_SEAT_CAPABILITY_POINTER = 1
- GDK_SEAT_CAPABILITY_TOUCH = 2
- GDK_SEAT_CAPABILITY_TABLET_STYLUS = 4
- GDK_SEAT_CAPABILITY_KEYBOARD = 8
- GDK_SEAT_CAPABILITY_ALL_POINTING = 7
- GDK_SEAT_CAPABILITY_ALL = 15

C - `GdkSeatCapabilities`

## `WMDecoration`

These are hints originally defined by the Motif toolkit.
The window manager can use them when determining how to decorate
the window. The hint must be set before mapping the window.

- GDK_DECOR_ALL = 1
- GDK_DECOR_BORDER = 2
- GDK_DECOR_RESIZEH = 4
- GDK_DECOR_TITLE = 8
- GDK_DECOR_MENU = 16
- GDK_DECOR_MINIMIZE = 32
- GDK_DECOR_MAXIMIZE = 64

C - `GdkWMDecoration`

## `WMFunction`

These are hints originally defined by the Motif toolkit. The window manager
can use them when determining the functions to offer for the window. The
hint must be set before mapping the window.

- GDK_FUNC_ALL = 1
- GDK_FUNC_RESIZE = 2
- GDK_FUNC_MOVE = 4
- GDK_FUNC_MINIMIZE = 8
- GDK_FUNC_MAXIMIZE = 16
- GDK_FUNC_CLOSE = 32

C - `GdkWMFunction`

## `WindowAttributesType`

Used to indicate which fields in the #GdkWindowAttr struct should be honored.
For example, if you filled in the “cursor” and “x” fields of #GdkWindowAttr,
pass “@GDK_WA_X | @GDK_WA_CURSOR” to gdk_window_new(). Fields in
&num;GdkWindowAttr not covered by a bit in this enum are required; for example,
the @width/@height, @wclass, and @window_type fields are required, they have
no corresponding flag in #GdkWindowAttributesType.

- GDK_WA_TITLE = 2
- GDK_WA_X = 4
- GDK_WA_Y = 8
- GDK_WA_CURSOR = 16
- GDK_WA_VISUAL = 32
- GDK_WA_WMCLASS = 64
- GDK_WA_NOREDIR = 128
- GDK_WA_TYPE_HINT = 256

C - `GdkWindowAttributesType`

## `WindowHints`

Used to indicate which fields of a #GdkGeometry struct should be paid
attention to. Also, the presence/absence of @GDK_HINT_POS,
@GDK_HINT_USER_POS, and @GDK_HINT_USER_SIZE is significant, though they don't
directly refer to #GdkGeometry fields. @GDK_HINT_USER_POS will be set
automatically by #GtkWindow if you call gtk_window_move().
@GDK_HINT_USER_POS and @GDK_HINT_USER_SIZE should be set if the user
specified a size/position using a --geometry command-line argument;
gtk_window_parse_geometry() automatically sets these flags.

- GDK_HINT_POS = 1
- GDK_HINT_MIN_SIZE = 2
- GDK_HINT_MAX_SIZE = 4
- GDK_HINT_BASE_SIZE = 8
- GDK_HINT_ASPECT = 16
- GDK_HINT_RESIZE_INC = 32
- GDK_HINT_WIN_GRAVITY = 64
- GDK_HINT_USER_POS = 128
- GDK_HINT_USER_SIZE = 256

C - `GdkWindowHints`

## `WindowState`

Specifies the state of a toplevel window.

- GDK_WINDOW_STATE_WITHDRAWN = 1
- GDK_WINDOW_STATE_ICONIFIED = 2
- GDK_WINDOW_STATE_MAXIMIZED = 4
- GDK_WINDOW_STATE_STICKY = 8
- GDK_WINDOW_STATE_FULLSCREEN = 16
- GDK_WINDOW_STATE_ABOVE = 32
- GDK_WINDOW_STATE_BELOW = 64
- GDK_WINDOW_STATE_FOCUSED = 128
- GDK_WINDOW_STATE_TILED = 256
- GDK_WINDOW_STATE_TOP_TILED = 512
- GDK_WINDOW_STATE_TOP_RESIZABLE = 1024
- GDK_WINDOW_STATE_RIGHT_TILED = 2048
- GDK_WINDOW_STATE_RIGHT_RESIZABLE = 4096
- GDK_WINDOW_STATE_BOTTOM_TILED = 8192
- GDK_WINDOW_STATE_BOTTOM_RESIZABLE = 16384
- GDK_WINDOW_STATE_LEFT_TILED = 32768
- GDK_WINDOW_STATE_LEFT_RESIZABLE = 65536

C - `GdkWindowState`

ize/position using a --geometry command-line argument;
gtk_window_parse_geometry() automatically sets these flags.

C - `GdkWindowHints`

- GDK_HINT_POS = %!s(int=1)
- GDK_HINT_MIN_SIZE = %!s(int=2)
- GDK_HINT_MAX_SIZE = %!s(int=4)
- GDK_HINT_BASE_SIZE = %!s(int=8)
- GDK_HINT_ASPECT = %!s(int=16)
- GDK_HINT_RESIZE_INC = %!s(int=32)
- GDK_HINT_WIN_GRAVITY = %!s(int=64)
- GDK_HINT_USER_POS = %!s(int=128)
- GDK_HINT_USER_SIZE = %!s(int=256)
## `WindowState`

Specifies the state of a toplevel window.

C - `GdkWindowState`

- GDK_WINDOW_STATE_WITHDRAWN = %!s(int=1)
- GDK_WINDOW_STATE_ICONIFIED = %!s(int=2)
- GDK_WINDOW_STATE_MAXIMIZED = %!s(int=4)
- GDK_WINDOW_STATE_STICKY = %!s(int=8)
- GDK_WINDOW_STATE_FULLSCREEN = %!s(int=16)
- GDK_WINDOW_STATE_ABOVE = %!s(int=32)
- GDK_WINDOW_STATE_BELOW = %!s(int=64)
- GDK_WINDOW_STATE_FOCUSED = %!s(int=128)
- GDK_WINDOW_STATE_TILED = %!s(int=256)
- GDK_WINDOW_STATE_TOP_TILED = %!s(int=512)
- GDK_WINDOW_STATE_TOP_RESIZABLE = %!s(int=1024)
- GDK_WINDOW_STATE_RIGHT_TILED = %!s(int=2048)
- GDK_WINDOW_STATE_RIGHT_RESIZABLE = %!s(int=4096)
- GDK_WINDOW_STATE_BOTTOM_TILED = %!s(int=8192)
- GDK_WINDOW_STATE_BOTTOM_RESIZABLE = %!s(int=16384)
- GDK_WINDOW_STATE_LEFT_TILED = %!s(int=32768)
- GDK_WINDOW_STATE_LEFT_RESIZABLE = %!s(int=65536)
