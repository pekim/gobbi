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

## `AxisFlags`

Flags describing the current capabilities of a device/tool.

C - `GdkAxisFlags`

## `DragAction`

Used in #GdkDragContext to indicate what the destination
should do with the dropped data.

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

C - `GdkEventMask`

## `FrameClockPhase`

#GdkFrameClockPhase is used to represent the different paint clock
phases that can be requested. The elements of the enumeration
correspond to the signals of #GdkFrameClock.

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

C - `guint`

## `SeatCapabilities`

Flags describing the seat capabilities.

C - `GdkSeatCapabilities`

## `WMDecoration`

These are hints originally defined by the Motif toolkit.
The window manager can use them when determining how to decorate
the window. The hint must be set before mapping the window.

C - `GdkWMDecoration`

## `WMFunction`

These are hints originally defined by the Motif toolkit. The window manager
can use them when determining the functions to offer for the window. The
hint must be set before mapping the window.

C - `GdkWMFunction`

## `WindowAttributesType`

Used to indicate which fields in the #GdkWindowAttr struct should be honored.
For example, if you filled in the “cursor” and “x” fields of #GdkWindowAttr,
pass “@GDK_WA_X | @GDK_WA_CURSOR” to gdk_window_new(). Fields in
&num;GdkWindowAttr not covered by a bit in this enum are required; for example,
the @width/@height, @wclass, and @window_type fields are required, they have
no corresponding flag in #GdkWindowAttributesType.

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

C - `GdkWindowHints`

## `WindowState`

Specifies the state of a toplevel window.

C - `GdkWindowState`

