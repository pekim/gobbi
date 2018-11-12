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

C - `GdkAnchorHints`

### GDK_ANCHOR_FLIP_X = 1
allow flipping anchors horizontally

### GDK_ANCHOR_FLIP_Y = 2
allow flipping anchors vertically

### GDK_ANCHOR_SLIDE_X = 4
allow sliding window horizontally

### GDK_ANCHOR_SLIDE_Y = 8
allow sliding window vertically

### GDK_ANCHOR_RESIZE_X = 16
allow resizing window horizontally

### GDK_ANCHOR_RESIZE_Y = 32
allow resizing window vertically

### GDK_ANCHOR_FLIP = 3
allow flipping anchors on both axes

### GDK_ANCHOR_SLIDE = 12
allow sliding window on both axes

### GDK_ANCHOR_RESIZE = 48
allow resizing window on both axes


## `AxisFlags`

Flags describing the current capabilities of a device/tool.

C - `GdkAxisFlags`

### GDK_AXIS_FLAG_X = 2
X axis is present

### GDK_AXIS_FLAG_Y = 4
Y axis is present

### GDK_AXIS_FLAG_PRESSURE = 8
Pressure axis is present

### GDK_AXIS_FLAG_XTILT = 16
X tilt axis is present

### GDK_AXIS_FLAG_YTILT = 32
Y tilt axis is present

### GDK_AXIS_FLAG_WHEEL = 64
Wheel axis is present

### GDK_AXIS_FLAG_DISTANCE = 128
Distance axis is present

### GDK_AXIS_FLAG_ROTATION = 256
Z-axis rotation is present

### GDK_AXIS_FLAG_SLIDER = 512
Slider axis is present


## `DragAction`

Used in #GdkDragContext to indicate what the destination
should do with the dropped data.

C - `GdkDragAction`

### GDK_ACTION_DEFAULT = 1
Means nothing, and should not be used.

### GDK_ACTION_COPY = 2
Copy the data.

### GDK_ACTION_MOVE = 4
Move the data, i.e. first copy it, then delete
 it from the source using the DELETE target of the X selection protocol.

### GDK_ACTION_LINK = 8
Add a link to the data. Note that this is only
 useful if source and destination agree on what it means.

### GDK_ACTION_PRIVATE = 16
Special action which tells the source that the
 destination will do something that the source doesn’t understand.

### GDK_ACTION_ASK = 32
Ask the user what to do with the data.


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

C - `GdkEventMask`

### GDK_EXPOSURE_MASK = 2
receive expose events

### GDK_POINTER_MOTION_MASK = 4
receive all pointer motion events

### GDK_POINTER_MOTION_HINT_MASK = 8
deprecated. see the explanation above

### GDK_BUTTON_MOTION_MASK = 16
receive pointer motion events while any button is pressed

### GDK_BUTTON1_MOTION_MASK = 32
receive pointer motion events while 1 button is pressed

### GDK_BUTTON2_MOTION_MASK = 64
receive pointer motion events while 2 button is pressed

### GDK_BUTTON3_MOTION_MASK = 128
receive pointer motion events while 3 button is pressed

### GDK_BUTTON_PRESS_MASK = 256
receive button press events

### GDK_BUTTON_RELEASE_MASK = 512
receive button release events

### GDK_KEY_PRESS_MASK = 1024
receive key press events

### GDK_KEY_RELEASE_MASK = 2048
receive key release events

### GDK_ENTER_NOTIFY_MASK = 4096
receive window enter events

### GDK_LEAVE_NOTIFY_MASK = 8192
receive window leave events

### GDK_FOCUS_CHANGE_MASK = 16384
receive focus change events

### GDK_STRUCTURE_MASK = 32768
receive events about window configuration change

### GDK_PROPERTY_CHANGE_MASK = 65536
receive property change events

### GDK_VISIBILITY_NOTIFY_MASK = 131072
receive visibility change events

### GDK_PROXIMITY_IN_MASK = 262144
receive proximity in events

### GDK_PROXIMITY_OUT_MASK = 524288
receive proximity out events

### GDK_SUBSTRUCTURE_MASK = 1048576
receive events about window configuration changes of
  child windows

### GDK_SCROLL_MASK = 2097152
receive scroll events

### GDK_TOUCH_MASK = 4194304
receive touch events. Since 3.4

### GDK_SMOOTH_SCROLL_MASK = 8388608
receive smooth scrolling events. Since 3.4

### GDK_TOUCHPAD_GESTURE_MASK = 16777216
receive touchpad gesture events. Since 3.18

### GDK_TABLET_PAD_MASK = 33554432
receive tablet pad events. Since 3.22

### GDK_ALL_EVENTS_MASK = 67108862
the combination of all the above event masks.


## `FrameClockPhase`

#GdkFrameClockPhase is used to represent the different paint clock
phases that can be requested. The elements of the enumeration
correspond to the signals of #GdkFrameClock.

C - `GdkFrameClockPhase`

### GDK_FRAME_CLOCK_PHASE_NONE = 0
no phase

### GDK_FRAME_CLOCK_PHASE_FLUSH_EVENTS = 1
corresponds to GdkFrameClock::flush-events. Should not be handled by applications.

### GDK_FRAME_CLOCK_PHASE_BEFORE_PAINT = 2
corresponds to GdkFrameClock::before-paint. Should not be handled by applications.

### GDK_FRAME_CLOCK_PHASE_UPDATE = 4
corresponds to GdkFrameClock::update.

### GDK_FRAME_CLOCK_PHASE_LAYOUT = 8
corresponds to GdkFrameClock::layout.

### GDK_FRAME_CLOCK_PHASE_PAINT = 16
corresponds to GdkFrameClock::paint.

### GDK_FRAME_CLOCK_PHASE_RESUME_EVENTS = 32
corresponds to GdkFrameClock::resume-events. Should not be handled by applications.

### GDK_FRAME_CLOCK_PHASE_AFTER_PAINT = 64
corresponds to GdkFrameClock::after-paint. Should not be handled by applications.


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

C - `guint`

### GDK_SHIFT_MASK = 1
the Shift key.

### GDK_LOCK_MASK = 2
a Lock key (depending on the modifier mapping of the
 X server this may either be CapsLock or ShiftLock).

### GDK_CONTROL_MASK = 4
the Control key.

### GDK_MOD1_MASK = 8
the fourth modifier key (it depends on the modifier
 mapping of the X server which key is interpreted as this modifier, but
 normally it is the Alt key).

### GDK_MOD2_MASK = 16
the fifth modifier key (it depends on the modifier
 mapping of the X server which key is interpreted as this modifier).

### GDK_MOD3_MASK = 32
the sixth modifier key (it depends on the modifier
 mapping of the X server which key is interpreted as this modifier).

### GDK_MOD4_MASK = 64
the seventh modifier key (it depends on the modifier
 mapping of the X server which key is interpreted as this modifier).

### GDK_MOD5_MASK = 128
the eighth modifier key (it depends on the modifier
 mapping of the X server which key is interpreted as this modifier).

### GDK_BUTTON1_MASK = 256
the first mouse button.

### GDK_BUTTON2_MASK = 512
the second mouse button.

### GDK_BUTTON3_MASK = 1024
the third mouse button.

### GDK_BUTTON4_MASK = 2048
the fourth mouse button.

### GDK_BUTTON5_MASK = 4096
the fifth mouse button.

### GDK_MODIFIER_RESERVED_13_MASK = 8192
A reserved bit flag; do not use in your own code

### GDK_MODIFIER_RESERVED_14_MASK = 16384
A reserved bit flag; do not use in your own code

### GDK_MODIFIER_RESERVED_15_MASK = 32768
A reserved bit flag; do not use in your own code

### GDK_MODIFIER_RESERVED_16_MASK = 65536
A reserved bit flag; do not use in your own code

### GDK_MODIFIER_RESERVED_17_MASK = 131072
A reserved bit flag; do not use in your own code

### GDK_MODIFIER_RESERVED_18_MASK = 262144
A reserved bit flag; do not use in your own code

### GDK_MODIFIER_RESERVED_19_MASK = 524288
A reserved bit flag; do not use in your own code

### GDK_MODIFIER_RESERVED_20_MASK = 1048576
A reserved bit flag; do not use in your own code

### GDK_MODIFIER_RESERVED_21_MASK = 2097152
A reserved bit flag; do not use in your own code

### GDK_MODIFIER_RESERVED_22_MASK = 4194304
A reserved bit flag; do not use in your own code

### GDK_MODIFIER_RESERVED_23_MASK = 8388608
A reserved bit flag; do not use in your own code

### GDK_MODIFIER_RESERVED_24_MASK = 16777216
A reserved bit flag; do not use in your own code

### GDK_MODIFIER_RESERVED_25_MASK = 33554432
A reserved bit flag; do not use in your own code

### GDK_SUPER_MASK = 67108864
the Super modifier. Since 2.10

### GDK_HYPER_MASK = 134217728
the Hyper modifier. Since 2.10

### GDK_META_MASK = 268435456
the Meta modifier. Since 2.10

### GDK_MODIFIER_RESERVED_29_MASK = 536870912
A reserved bit flag; do not use in your own code

### GDK_RELEASE_MASK = 1073741824
not used in GDK itself. GTK+ uses it to differentiate
 between (keyval, modifiers) pairs from key press and release events.

### GDK_MODIFIER_MASK = 1543512063
a mask covering all modifier types.


## `SeatCapabilities`

Flags describing the seat capabilities.

C - `GdkSeatCapabilities`

### GDK_SEAT_CAPABILITY_NONE = 0
No input capabilities

### GDK_SEAT_CAPABILITY_POINTER = 1
The seat has a pointer (e.g. mouse)

### GDK_SEAT_CAPABILITY_TOUCH = 2
The seat has touchscreen(s) attached

### GDK_SEAT_CAPABILITY_TABLET_STYLUS = 4
The seat has drawing tablet(s) attached

### GDK_SEAT_CAPABILITY_KEYBOARD = 8
The seat has keyboard(s) attached

### GDK_SEAT_CAPABILITY_ALL_POINTING = 7
The union of all pointing capabilities

### GDK_SEAT_CAPABILITY_ALL = 15
The union of all capabilities


## `WMDecoration`

These are hints originally defined by the Motif toolkit.
The window manager can use them when determining how to decorate
the window. The hint must be set before mapping the window.

C - `GdkWMDecoration`

### GDK_DECOR_ALL = 1
all decorations should be applied.

### GDK_DECOR_BORDER = 2
a frame should be drawn around the window.

### GDK_DECOR_RESIZEH = 4
the frame should have resize handles.

### GDK_DECOR_TITLE = 8
a titlebar should be placed above the window.

### GDK_DECOR_MENU = 16
a button for opening a menu should be included.

### GDK_DECOR_MINIMIZE = 32
a minimize button should be included.

### GDK_DECOR_MAXIMIZE = 64
a maximize button should be included.


## `WMFunction`

These are hints originally defined by the Motif toolkit. The window manager
can use them when determining the functions to offer for the window. The
hint must be set before mapping the window.

C - `GdkWMFunction`

### GDK_FUNC_ALL = 1
all functions should be offered.

### GDK_FUNC_RESIZE = 2
the window should be resizable.

### GDK_FUNC_MOVE = 4
the window should be movable.

### GDK_FUNC_MINIMIZE = 8
the window should be minimizable.

### GDK_FUNC_MAXIMIZE = 16
the window should be maximizable.

### GDK_FUNC_CLOSE = 32
the window should be closable.


## `WindowAttributesType`

Used to indicate which fields in the #GdkWindowAttr struct should be honored.
For example, if you filled in the “cursor” and “x” fields of #GdkWindowAttr,
pass “@GDK_WA_X | @GDK_WA_CURSOR” to gdk_window_new(). Fields in
&num;GdkWindowAttr not covered by a bit in this enum are required; for example,
the @width/@height, @wclass, and @window_type fields are required, they have
no corresponding flag in #GdkWindowAttributesType.

C - `GdkWindowAttributesType`

### GDK_WA_TITLE = 2
Honor the title field

### GDK_WA_X = 4
Honor the X coordinate field

### GDK_WA_Y = 8
Honor the Y coordinate field

### GDK_WA_CURSOR = 16
Honor the cursor field

### GDK_WA_VISUAL = 32
Honor the visual field

### GDK_WA_WMCLASS = 64
Honor the wmclass_class and wmclass_name fields

### GDK_WA_NOREDIR = 128
Honor the override_redirect field

### GDK_WA_TYPE_HINT = 256
Honor the type_hint field


## `WindowHints`

Used to indicate which fields of a #GdkGeometry struct should be paid
attention to. Also, the presence/absence of @GDK_HINT_POS,
@GDK_HINT_USER_POS, and @GDK_HINT_USER_SIZE is significant, though they don't
directly refer to #GdkGeometry fields. @GDK_HINT_USER_POS will be set
automatically by #GtkWindow if you call gtk_window_move().
@GDK_HINT_USER_POS and @GDK_HINT_USER_SIZE should be set if the user
specified a size/position using a --geometry command-line argument;
gtk_window_parse_geometry() automatically sets these flags.

C - `GdkWindowHints`

### GDK_HINT_POS = 1
indicates that the program has positioned the window

### GDK_HINT_MIN_SIZE = 2
min size fields are set

### GDK_HINT_MAX_SIZE = 4
max size fields are set

### GDK_HINT_BASE_SIZE = 8
base size fields are set

### GDK_HINT_ASPECT = 16
aspect ratio fields are set

### GDK_HINT_RESIZE_INC = 32
resize increment fields are set

### GDK_HINT_WIN_GRAVITY = 64
window gravity field is set

### GDK_HINT_USER_POS = 128
indicates that the window’s position was explicitly set
 by the user

### GDK_HINT_USER_SIZE = 256
indicates that the window’s size was explicitly set by
 the user


## `WindowState`

Specifies the state of a toplevel window.

C - `GdkWindowState`

### GDK_WINDOW_STATE_WITHDRAWN = 1
the window is not shown.

### GDK_WINDOW_STATE_ICONIFIED = 2
the window is minimized.

### GDK_WINDOW_STATE_MAXIMIZED = 4
the window is maximized.

### GDK_WINDOW_STATE_STICKY = 8
the window is sticky.

### GDK_WINDOW_STATE_FULLSCREEN = 16
the window is maximized without
  decorations.

### GDK_WINDOW_STATE_ABOVE = 32
the window is kept above other windows.

### GDK_WINDOW_STATE_BELOW = 64
the window is kept below other windows.

### GDK_WINDOW_STATE_FOCUSED = 128
the window is presented as focused (with active decorations).

### GDK_WINDOW_STATE_TILED = 256
the window is in a tiled state, Since 3.10. Since 3.22.23, this
                         is deprecated in favor of per-edge information.

### GDK_WINDOW_STATE_TOP_TILED = 512
whether the top edge is tiled, Since 3.22.23

### GDK_WINDOW_STATE_TOP_RESIZABLE = 1024
whether the top edge is resizable, Since 3.22.23

### GDK_WINDOW_STATE_RIGHT_TILED = 2048
whether the right edge is tiled, Since 3.22.23

### GDK_WINDOW_STATE_RIGHT_RESIZABLE = 4096
whether the right edge is resizable, Since 3.22.23

### GDK_WINDOW_STATE_BOTTOM_TILED = 8192
whether the bottom edge is tiled, Since 3.22.23

### GDK_WINDOW_STATE_BOTTOM_RESIZABLE = 16384
whether the bottom edge is resizable, Since 3.22.23

### GDK_WINDOW_STATE_LEFT_TILED = 32768
whether the left edge is tiled, Since 3.22.23

### GDK_WINDOW_STATE_LEFT_RESIZABLE = 65536
whether the left edge is resizable, Since 3.22.23


