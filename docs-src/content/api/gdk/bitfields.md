+++
title = "bitfields"
+++
<p class="api-heading">AnchorHints</p>
<p class="api-doc">Positioning hints for aligning a window relative to a rectangle.

These hints determine how the window should be positioned in the case that
the window would fall off-screen if placed in its ideal position.

For example, %GDK_ANCHOR_FLIP_X will replace %GDK_GRAVITY_NORTH_WEST with
%GDK_GRAVITY_NORTH_EAST and vice versa if the window extends beyond the left
or right edges of the monitor.

If %GDK_ANCHOR_SLIDE_X is set, the window can be shifted horizontally to fit
on-screen. If %GDK_ANCHOR_RESIZE_X is set, the window can be shrunken
horizontally to fit.

In general, when multiple flags are set, flipping should take precedence over
sliding, which should take precedence over resizing.</p>
<table>
<tr>
<td class="name">GDK_ANCHOR_FLIP_X</td>
<td class="value">1</td>
<td class="doc">allow flipping anchors horizontally</td>
</tr>
<tr>
<td class="name">GDK_ANCHOR_FLIP_Y</td>
<td class="value">2</td>
<td class="doc">allow flipping anchors vertically</td>
</tr>
<tr>
<td class="name">GDK_ANCHOR_SLIDE_X</td>
<td class="value">4</td>
<td class="doc">allow sliding window horizontally</td>
</tr>
<tr>
<td class="name">GDK_ANCHOR_SLIDE_Y</td>
<td class="value">8</td>
<td class="doc">allow sliding window vertically</td>
</tr>
<tr>
<td class="name">GDK_ANCHOR_RESIZE_X</td>
<td class="value">16</td>
<td class="doc">allow resizing window horizontally</td>
</tr>
<tr>
<td class="name">GDK_ANCHOR_RESIZE_Y</td>
<td class="value">32</td>
<td class="doc">allow resizing window vertically</td>
</tr>
<tr>
<td class="name">GDK_ANCHOR_FLIP</td>
<td class="value">3</td>
<td class="doc">allow flipping anchors on both axes</td>
</tr>
<tr>
<td class="name">GDK_ANCHOR_SLIDE</td>
<td class="value">12</td>
<td class="doc">allow sliding window on both axes</td>
</tr>
<tr>
<td class="name">GDK_ANCHOR_RESIZE</td>
<td class="value">48</td>
<td class="doc">allow resizing window on both axes</td>
</tr>
</table>
<div class="api-notes">
  <p class="api-since">since 3.22</p>
  <p class="api-ctype">C type: GdkAnchorHints</p>
</div>
<p class="api-heading">AxisFlags</p>
<p class="api-doc">Flags describing the current capabilities of a device/tool.</p>
<table>
<tr>
<td class="name">GDK_AXIS_FLAG_X</td>
<td class="value">2</td>
<td class="doc">X axis is present</td>
</tr>
<tr>
<td class="name">GDK_AXIS_FLAG_Y</td>
<td class="value">4</td>
<td class="doc">Y axis is present</td>
</tr>
<tr>
<td class="name">GDK_AXIS_FLAG_PRESSURE</td>
<td class="value">8</td>
<td class="doc">Pressure axis is present</td>
</tr>
<tr>
<td class="name">GDK_AXIS_FLAG_XTILT</td>
<td class="value">16</td>
<td class="doc">X tilt axis is present</td>
</tr>
<tr>
<td class="name">GDK_AXIS_FLAG_YTILT</td>
<td class="value">32</td>
<td class="doc">Y tilt axis is present</td>
</tr>
<tr>
<td class="name">GDK_AXIS_FLAG_WHEEL</td>
<td class="value">64</td>
<td class="doc">Wheel axis is present</td>
</tr>
<tr>
<td class="name">GDK_AXIS_FLAG_DISTANCE</td>
<td class="value">128</td>
<td class="doc">Distance axis is present</td>
</tr>
<tr>
<td class="name">GDK_AXIS_FLAG_ROTATION</td>
<td class="value">256</td>
<td class="doc">Z-axis rotation is present</td>
</tr>
<tr>
<td class="name">GDK_AXIS_FLAG_SLIDER</td>
<td class="value">512</td>
<td class="doc">Slider axis is present</td>
</tr>
</table>
<div class="api-notes">
  <p class="api-since">since 3.22</p>
  <p class="api-ctype">C type: GdkAxisFlags</p>
</div>
<p class="api-heading">DragAction</p>
<p class="api-doc">Used in #GdkDragContext to indicate what the destination
should do with the dropped data.</p>
<table>
<tr>
<td class="name">GDK_ACTION_DEFAULT</td>
<td class="value">1</td>
<td class="doc">Means nothing, and should not be used.</td>
</tr>
<tr>
<td class="name">GDK_ACTION_COPY</td>
<td class="value">2</td>
<td class="doc">Copy the data.</td>
</tr>
<tr>
<td class="name">GDK_ACTION_MOVE</td>
<td class="value">4</td>
<td class="doc">Move the data, i.e. first copy it, then delete
 it from the source using the DELETE target of the X selection protocol.</td>
</tr>
<tr>
<td class="name">GDK_ACTION_LINK</td>
<td class="value">8</td>
<td class="doc">Add a link to the data. Note that this is only
 useful if source and destination agree on what it means.</td>
</tr>
<tr>
<td class="name">GDK_ACTION_PRIVATE</td>
<td class="value">16</td>
<td class="doc">Special action which tells the source that the
 destination will do something that the source doesn’t understand.</td>
</tr>
<tr>
<td class="name">GDK_ACTION_ASK</td>
<td class="value">32</td>
<td class="doc">Ask the user what to do with the data.</td>
</tr>
</table>
<div class="api-notes">
  <p class="api-ctype">C type: GdkDragAction</p>
</div>
<p class="api-heading">EventMask</p>
<p class="api-doc">A set of bit-flags to indicate which events a window is to receive.
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
events, so different sequences may be distinguished.</p>
<table>
<tr>
<td class="name">GDK_EXPOSURE_MASK</td>
<td class="value">2</td>
<td class="doc">receive expose events</td>
</tr>
<tr>
<td class="name">GDK_POINTER_MOTION_MASK</td>
<td class="value">4</td>
<td class="doc">receive all pointer motion events</td>
</tr>
<tr>
<td class="name">GDK_POINTER_MOTION_HINT_MASK</td>
<td class="value">8</td>
<td class="doc">deprecated. see the explanation above</td>
</tr>
<tr>
<td class="name">GDK_BUTTON_MOTION_MASK</td>
<td class="value">16</td>
<td class="doc">receive pointer motion events while any button is pressed</td>
</tr>
<tr>
<td class="name">GDK_BUTTON1_MOTION_MASK</td>
<td class="value">32</td>
<td class="doc">receive pointer motion events while 1 button is pressed</td>
</tr>
<tr>
<td class="name">GDK_BUTTON2_MOTION_MASK</td>
<td class="value">64</td>
<td class="doc">receive pointer motion events while 2 button is pressed</td>
</tr>
<tr>
<td class="name">GDK_BUTTON3_MOTION_MASK</td>
<td class="value">128</td>
<td class="doc">receive pointer motion events while 3 button is pressed</td>
</tr>
<tr>
<td class="name">GDK_BUTTON_PRESS_MASK</td>
<td class="value">256</td>
<td class="doc">receive button press events</td>
</tr>
<tr>
<td class="name">GDK_BUTTON_RELEASE_MASK</td>
<td class="value">512</td>
<td class="doc">receive button release events</td>
</tr>
<tr>
<td class="name">GDK_KEY_PRESS_MASK</td>
<td class="value">1024</td>
<td class="doc">receive key press events</td>
</tr>
<tr>
<td class="name">GDK_KEY_RELEASE_MASK</td>
<td class="value">2048</td>
<td class="doc">receive key release events</td>
</tr>
<tr>
<td class="name">GDK_ENTER_NOTIFY_MASK</td>
<td class="value">4096</td>
<td class="doc">receive window enter events</td>
</tr>
<tr>
<td class="name">GDK_LEAVE_NOTIFY_MASK</td>
<td class="value">8192</td>
<td class="doc">receive window leave events</td>
</tr>
<tr>
<td class="name">GDK_FOCUS_CHANGE_MASK</td>
<td class="value">16384</td>
<td class="doc">receive focus change events</td>
</tr>
<tr>
<td class="name">GDK_STRUCTURE_MASK</td>
<td class="value">32768</td>
<td class="doc">receive events about window configuration change</td>
</tr>
<tr>
<td class="name">GDK_PROPERTY_CHANGE_MASK</td>
<td class="value">65536</td>
<td class="doc">receive property change events</td>
</tr>
<tr>
<td class="name">GDK_VISIBILITY_NOTIFY_MASK</td>
<td class="value">131072</td>
<td class="doc">receive visibility change events</td>
</tr>
<tr>
<td class="name">GDK_PROXIMITY_IN_MASK</td>
<td class="value">262144</td>
<td class="doc">receive proximity in events</td>
</tr>
<tr>
<td class="name">GDK_PROXIMITY_OUT_MASK</td>
<td class="value">524288</td>
<td class="doc">receive proximity out events</td>
</tr>
<tr>
<td class="name">GDK_SUBSTRUCTURE_MASK</td>
<td class="value">1048576</td>
<td class="doc">receive events about window configuration changes of
  child windows</td>
</tr>
<tr>
<td class="name">GDK_SCROLL_MASK</td>
<td class="value">2097152</td>
<td class="doc">receive scroll events</td>
</tr>
<tr>
<td class="name">GDK_TOUCH_MASK</td>
<td class="value">4194304</td>
<td class="doc">receive touch events. Since 3.4</td>
</tr>
<tr>
<td class="name">GDK_SMOOTH_SCROLL_MASK</td>
<td class="value">8388608</td>
<td class="doc">receive smooth scrolling events. Since 3.4</td>
</tr>
<tr>
<td class="name">GDK_TOUCHPAD_GESTURE_MASK</td>
<td class="value">16777216</td>
<td class="doc">receive touchpad gesture events. Since 3.18</td>
</tr>
<tr>
<td class="name">GDK_TABLET_PAD_MASK</td>
<td class="value">33554432</td>
<td class="doc">receive tablet pad events. Since 3.22</td>
</tr>
<tr>
<td class="name">GDK_ALL_EVENTS_MASK</td>
<td class="value">67108862</td>
<td class="doc">the combination of all the above event masks.</td>
</tr>
</table>
<div class="api-notes">
  <p class="api-ctype">C type: GdkEventMask</p>
</div>
<p class="api-heading">FrameClockPhase</p>
<p class="api-doc">#GdkFrameClockPhase is used to represent the different paint clock
phases that can be requested. The elements of the enumeration
correspond to the signals of #GdkFrameClock.</p>
<table>
<tr>
<td class="name">GDK_FRAME_CLOCK_PHASE_NONE</td>
<td class="value">0</td>
<td class="doc">no phase</td>
</tr>
<tr>
<td class="name">GDK_FRAME_CLOCK_PHASE_FLUSH_EVENTS</td>
<td class="value">1</td>
<td class="doc">corresponds to GdkFrameClock::flush-events. Should not be handled by applications.</td>
</tr>
<tr>
<td class="name">GDK_FRAME_CLOCK_PHASE_BEFORE_PAINT</td>
<td class="value">2</td>
<td class="doc">corresponds to GdkFrameClock::before-paint. Should not be handled by applications.</td>
</tr>
<tr>
<td class="name">GDK_FRAME_CLOCK_PHASE_UPDATE</td>
<td class="value">4</td>
<td class="doc">corresponds to GdkFrameClock::update.</td>
</tr>
<tr>
<td class="name">GDK_FRAME_CLOCK_PHASE_LAYOUT</td>
<td class="value">8</td>
<td class="doc">corresponds to GdkFrameClock::layout.</td>
</tr>
<tr>
<td class="name">GDK_FRAME_CLOCK_PHASE_PAINT</td>
<td class="value">16</td>
<td class="doc">corresponds to GdkFrameClock::paint.</td>
</tr>
<tr>
<td class="name">GDK_FRAME_CLOCK_PHASE_RESUME_EVENTS</td>
<td class="value">32</td>
<td class="doc">corresponds to GdkFrameClock::resume-events. Should not be handled by applications.</td>
</tr>
<tr>
<td class="name">GDK_FRAME_CLOCK_PHASE_AFTER_PAINT</td>
<td class="value">64</td>
<td class="doc">corresponds to GdkFrameClock::after-paint. Should not be handled by applications.</td>
</tr>
</table>
<div class="api-notes">
  <p class="api-since">since 3.8</p>
  <p class="api-ctype">C type: GdkFrameClockPhase</p>
</div>
<p class="api-heading">ModifierType</p>
<p class="api-doc">A set of bit-flags to indicate the state of modifier keys and mouse buttons
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
be set.</p>
<table>
<tr>
<td class="name">GDK_SHIFT_MASK</td>
<td class="value">1</td>
<td class="doc">the Shift key.</td>
</tr>
<tr>
<td class="name">GDK_LOCK_MASK</td>
<td class="value">2</td>
<td class="doc">a Lock key (depending on the modifier mapping of the
 X server this may either be CapsLock or ShiftLock).</td>
</tr>
<tr>
<td class="name">GDK_CONTROL_MASK</td>
<td class="value">4</td>
<td class="doc">the Control key.</td>
</tr>
<tr>
<td class="name">GDK_MOD1_MASK</td>
<td class="value">8</td>
<td class="doc">the fourth modifier key (it depends on the modifier
 mapping of the X server which key is interpreted as this modifier, but
 normally it is the Alt key).</td>
</tr>
<tr>
<td class="name">GDK_MOD2_MASK</td>
<td class="value">16</td>
<td class="doc">the fifth modifier key (it depends on the modifier
 mapping of the X server which key is interpreted as this modifier).</td>
</tr>
<tr>
<td class="name">GDK_MOD3_MASK</td>
<td class="value">32</td>
<td class="doc">the sixth modifier key (it depends on the modifier
 mapping of the X server which key is interpreted as this modifier).</td>
</tr>
<tr>
<td class="name">GDK_MOD4_MASK</td>
<td class="value">64</td>
<td class="doc">the seventh modifier key (it depends on the modifier
 mapping of the X server which key is interpreted as this modifier).</td>
</tr>
<tr>
<td class="name">GDK_MOD5_MASK</td>
<td class="value">128</td>
<td class="doc">the eighth modifier key (it depends on the modifier
 mapping of the X server which key is interpreted as this modifier).</td>
</tr>
<tr>
<td class="name">GDK_BUTTON1_MASK</td>
<td class="value">256</td>
<td class="doc">the first mouse button.</td>
</tr>
<tr>
<td class="name">GDK_BUTTON2_MASK</td>
<td class="value">512</td>
<td class="doc">the second mouse button.</td>
</tr>
<tr>
<td class="name">GDK_BUTTON3_MASK</td>
<td class="value">1024</td>
<td class="doc">the third mouse button.</td>
</tr>
<tr>
<td class="name">GDK_BUTTON4_MASK</td>
<td class="value">2048</td>
<td class="doc">the fourth mouse button.</td>
</tr>
<tr>
<td class="name">GDK_BUTTON5_MASK</td>
<td class="value">4096</td>
<td class="doc">the fifth mouse button.</td>
</tr>
<tr>
<td class="name">GDK_MODIFIER_RESERVED_13_MASK</td>
<td class="value">8192</td>
<td class="doc">A reserved bit flag; do not use in your own code</td>
</tr>
<tr>
<td class="name">GDK_MODIFIER_RESERVED_14_MASK</td>
<td class="value">16384</td>
<td class="doc">A reserved bit flag; do not use in your own code</td>
</tr>
<tr>
<td class="name">GDK_MODIFIER_RESERVED_15_MASK</td>
<td class="value">32768</td>
<td class="doc">A reserved bit flag; do not use in your own code</td>
</tr>
<tr>
<td class="name">GDK_MODIFIER_RESERVED_16_MASK</td>
<td class="value">65536</td>
<td class="doc">A reserved bit flag; do not use in your own code</td>
</tr>
<tr>
<td class="name">GDK_MODIFIER_RESERVED_17_MASK</td>
<td class="value">131072</td>
<td class="doc">A reserved bit flag; do not use in your own code</td>
</tr>
<tr>
<td class="name">GDK_MODIFIER_RESERVED_18_MASK</td>
<td class="value">262144</td>
<td class="doc">A reserved bit flag; do not use in your own code</td>
</tr>
<tr>
<td class="name">GDK_MODIFIER_RESERVED_19_MASK</td>
<td class="value">524288</td>
<td class="doc">A reserved bit flag; do not use in your own code</td>
</tr>
<tr>
<td class="name">GDK_MODIFIER_RESERVED_20_MASK</td>
<td class="value">1048576</td>
<td class="doc">A reserved bit flag; do not use in your own code</td>
</tr>
<tr>
<td class="name">GDK_MODIFIER_RESERVED_21_MASK</td>
<td class="value">2097152</td>
<td class="doc">A reserved bit flag; do not use in your own code</td>
</tr>
<tr>
<td class="name">GDK_MODIFIER_RESERVED_22_MASK</td>
<td class="value">4194304</td>
<td class="doc">A reserved bit flag; do not use in your own code</td>
</tr>
<tr>
<td class="name">GDK_MODIFIER_RESERVED_23_MASK</td>
<td class="value">8388608</td>
<td class="doc">A reserved bit flag; do not use in your own code</td>
</tr>
<tr>
<td class="name">GDK_MODIFIER_RESERVED_24_MASK</td>
<td class="value">16777216</td>
<td class="doc">A reserved bit flag; do not use in your own code</td>
</tr>
<tr>
<td class="name">GDK_MODIFIER_RESERVED_25_MASK</td>
<td class="value">33554432</td>
<td class="doc">A reserved bit flag; do not use in your own code</td>
</tr>
<tr>
<td class="name">GDK_SUPER_MASK</td>
<td class="value">67108864</td>
<td class="doc">the Super modifier. Since 2.10</td>
</tr>
<tr>
<td class="name">GDK_HYPER_MASK</td>
<td class="value">134217728</td>
<td class="doc">the Hyper modifier. Since 2.10</td>
</tr>
<tr>
<td class="name">GDK_META_MASK</td>
<td class="value">268435456</td>
<td class="doc">the Meta modifier. Since 2.10</td>
</tr>
<tr>
<td class="name">GDK_MODIFIER_RESERVED_29_MASK</td>
<td class="value">536870912</td>
<td class="doc">A reserved bit flag; do not use in your own code</td>
</tr>
<tr>
<td class="name">GDK_RELEASE_MASK</td>
<td class="value">1073741824</td>
<td class="doc">not used in GDK itself. GTK+ uses it to differentiate
 between (keyval, modifiers) pairs from key press and release events.</td>
</tr>
<tr>
<td class="name">GDK_MODIFIER_MASK</td>
<td class="value">1543512063</td>
<td class="doc">a mask covering all modifier types.</td>
</tr>
</table>
<div class="api-notes">
  <p class="api-ctype">C type: guint</p>
</div>
<p class="api-heading">SeatCapabilities</p>
<p class="api-doc">Flags describing the seat capabilities.</p>
<table>
<tr>
<td class="name">GDK_SEAT_CAPABILITY_NONE</td>
<td class="value">0</td>
<td class="doc">No input capabilities</td>
</tr>
<tr>
<td class="name">GDK_SEAT_CAPABILITY_POINTER</td>
<td class="value">1</td>
<td class="doc">The seat has a pointer (e.g. mouse)</td>
</tr>
<tr>
<td class="name">GDK_SEAT_CAPABILITY_TOUCH</td>
<td class="value">2</td>
<td class="doc">The seat has touchscreen(s) attached</td>
</tr>
<tr>
<td class="name">GDK_SEAT_CAPABILITY_TABLET_STYLUS</td>
<td class="value">4</td>
<td class="doc">The seat has drawing tablet(s) attached</td>
</tr>
<tr>
<td class="name">GDK_SEAT_CAPABILITY_KEYBOARD</td>
<td class="value">8</td>
<td class="doc">The seat has keyboard(s) attached</td>
</tr>
<tr>
<td class="name">GDK_SEAT_CAPABILITY_ALL_POINTING</td>
<td class="value">7</td>
<td class="doc">The union of all pointing capabilities</td>
</tr>
<tr>
<td class="name">GDK_SEAT_CAPABILITY_ALL</td>
<td class="value">15</td>
<td class="doc">The union of all capabilities</td>
</tr>
</table>
<div class="api-notes">
  <p class="api-since">since 3.20</p>
  <p class="api-ctype">C type: GdkSeatCapabilities</p>
</div>
<p class="api-heading">WMDecoration</p>
<p class="api-doc">These are hints originally defined by the Motif toolkit.
The window manager can use them when determining how to decorate
the window. The hint must be set before mapping the window.</p>
<table>
<tr>
<td class="name">GDK_DECOR_ALL</td>
<td class="value">1</td>
<td class="doc">all decorations should be applied.</td>
</tr>
<tr>
<td class="name">GDK_DECOR_BORDER</td>
<td class="value">2</td>
<td class="doc">a frame should be drawn around the window.</td>
</tr>
<tr>
<td class="name">GDK_DECOR_RESIZEH</td>
<td class="value">4</td>
<td class="doc">the frame should have resize handles.</td>
</tr>
<tr>
<td class="name">GDK_DECOR_TITLE</td>
<td class="value">8</td>
<td class="doc">a titlebar should be placed above the window.</td>
</tr>
<tr>
<td class="name">GDK_DECOR_MENU</td>
<td class="value">16</td>
<td class="doc">a button for opening a menu should be included.</td>
</tr>
<tr>
<td class="name">GDK_DECOR_MINIMIZE</td>
<td class="value">32</td>
<td class="doc">a minimize button should be included.</td>
</tr>
<tr>
<td class="name">GDK_DECOR_MAXIMIZE</td>
<td class="value">64</td>
<td class="doc">a maximize button should be included.</td>
</tr>
</table>
<div class="api-notes">
  <p class="api-ctype">C type: GdkWMDecoration</p>
</div>
<p class="api-heading">WMFunction</p>
<p class="api-doc">These are hints originally defined by the Motif toolkit. The window manager
can use them when determining the functions to offer for the window. The
hint must be set before mapping the window.</p>
<table>
<tr>
<td class="name">GDK_FUNC_ALL</td>
<td class="value">1</td>
<td class="doc">all functions should be offered.</td>
</tr>
<tr>
<td class="name">GDK_FUNC_RESIZE</td>
<td class="value">2</td>
<td class="doc">the window should be resizable.</td>
</tr>
<tr>
<td class="name">GDK_FUNC_MOVE</td>
<td class="value">4</td>
<td class="doc">the window should be movable.</td>
</tr>
<tr>
<td class="name">GDK_FUNC_MINIMIZE</td>
<td class="value">8</td>
<td class="doc">the window should be minimizable.</td>
</tr>
<tr>
<td class="name">GDK_FUNC_MAXIMIZE</td>
<td class="value">16</td>
<td class="doc">the window should be maximizable.</td>
</tr>
<tr>
<td class="name">GDK_FUNC_CLOSE</td>
<td class="value">32</td>
<td class="doc">the window should be closable.</td>
</tr>
</table>
<div class="api-notes">
  <p class="api-ctype">C type: GdkWMFunction</p>
</div>
<p class="api-heading">WindowAttributesType</p>
<p class="api-doc">Used to indicate which fields in the #GdkWindowAttr struct should be honored.
For example, if you filled in the “cursor” and “x” fields of #GdkWindowAttr,
pass “@GDK_WA_X | @GDK_WA_CURSOR” to gdk_window_new(). Fields in
#GdkWindowAttr not covered by a bit in this enum are required; for example,
the @width/@height, @wclass, and @window_type fields are required, they have
no corresponding flag in #GdkWindowAttributesType.</p>
<table>
<tr>
<td class="name">GDK_WA_TITLE</td>
<td class="value">2</td>
<td class="doc">Honor the title field</td>
</tr>
<tr>
<td class="name">GDK_WA_X</td>
<td class="value">4</td>
<td class="doc">Honor the X coordinate field</td>
</tr>
<tr>
<td class="name">GDK_WA_Y</td>
<td class="value">8</td>
<td class="doc">Honor the Y coordinate field</td>
</tr>
<tr>
<td class="name">GDK_WA_CURSOR</td>
<td class="value">16</td>
<td class="doc">Honor the cursor field</td>
</tr>
<tr>
<td class="name">GDK_WA_VISUAL</td>
<td class="value">32</td>
<td class="doc">Honor the visual field</td>
</tr>
<tr>
<td class="name">GDK_WA_WMCLASS</td>
<td class="value">64</td>
<td class="doc">Honor the wmclass_class and wmclass_name fields</td>
</tr>
<tr>
<td class="name">GDK_WA_NOREDIR</td>
<td class="value">128</td>
<td class="doc">Honor the override_redirect field</td>
</tr>
<tr>
<td class="name">GDK_WA_TYPE_HINT</td>
<td class="value">256</td>
<td class="doc">Honor the type_hint field</td>
</tr>
</table>
<div class="api-notes">
  <p class="api-ctype">C type: GdkWindowAttributesType</p>
</div>
<p class="api-heading">WindowHints</p>
<p class="api-doc">Used to indicate which fields of a #GdkGeometry struct should be paid
attention to. Also, the presence/absence of @GDK_HINT_POS,
@GDK_HINT_USER_POS, and @GDK_HINT_USER_SIZE is significant, though they don't
directly refer to #GdkGeometry fields. @GDK_HINT_USER_POS will be set
automatically by #GtkWindow if you call gtk_window_move().
@GDK_HINT_USER_POS and @GDK_HINT_USER_SIZE should be set if the user
specified a size/position using a --geometry command-line argument;
gtk_window_parse_geometry() automatically sets these flags.</p>
<table>
<tr>
<td class="name">GDK_HINT_POS</td>
<td class="value">1</td>
<td class="doc">indicates that the program has positioned the window</td>
</tr>
<tr>
<td class="name">GDK_HINT_MIN_SIZE</td>
<td class="value">2</td>
<td class="doc">min size fields are set</td>
</tr>
<tr>
<td class="name">GDK_HINT_MAX_SIZE</td>
<td class="value">4</td>
<td class="doc">max size fields are set</td>
</tr>
<tr>
<td class="name">GDK_HINT_BASE_SIZE</td>
<td class="value">8</td>
<td class="doc">base size fields are set</td>
</tr>
<tr>
<td class="name">GDK_HINT_ASPECT</td>
<td class="value">16</td>
<td class="doc">aspect ratio fields are set</td>
</tr>
<tr>
<td class="name">GDK_HINT_RESIZE_INC</td>
<td class="value">32</td>
<td class="doc">resize increment fields are set</td>
</tr>
<tr>
<td class="name">GDK_HINT_WIN_GRAVITY</td>
<td class="value">64</td>
<td class="doc">window gravity field is set</td>
</tr>
<tr>
<td class="name">GDK_HINT_USER_POS</td>
<td class="value">128</td>
<td class="doc">indicates that the window’s position was explicitly set
 by the user</td>
</tr>
<tr>
<td class="name">GDK_HINT_USER_SIZE</td>
<td class="value">256</td>
<td class="doc">indicates that the window’s size was explicitly set by
 the user</td>
</tr>
</table>
<div class="api-notes">
  <p class="api-ctype">C type: GdkWindowHints</p>
</div>
<p class="api-heading">WindowState</p>
<p class="api-doc">Specifies the state of a toplevel window.</p>
<table>
<tr>
<td class="name">GDK_WINDOW_STATE_WITHDRAWN</td>
<td class="value">1</td>
<td class="doc">the window is not shown.</td>
</tr>
<tr>
<td class="name">GDK_WINDOW_STATE_ICONIFIED</td>
<td class="value">2</td>
<td class="doc">the window is minimized.</td>
</tr>
<tr>
<td class="name">GDK_WINDOW_STATE_MAXIMIZED</td>
<td class="value">4</td>
<td class="doc">the window is maximized.</td>
</tr>
<tr>
<td class="name">GDK_WINDOW_STATE_STICKY</td>
<td class="value">8</td>
<td class="doc">the window is sticky.</td>
</tr>
<tr>
<td class="name">GDK_WINDOW_STATE_FULLSCREEN</td>
<td class="value">16</td>
<td class="doc">the window is maximized without
  decorations.</td>
</tr>
<tr>
<td class="name">GDK_WINDOW_STATE_ABOVE</td>
<td class="value">32</td>
<td class="doc">the window is kept above other windows.</td>
</tr>
<tr>
<td class="name">GDK_WINDOW_STATE_BELOW</td>
<td class="value">64</td>
<td class="doc">the window is kept below other windows.</td>
</tr>
<tr>
<td class="name">GDK_WINDOW_STATE_FOCUSED</td>
<td class="value">128</td>
<td class="doc">the window is presented as focused (with active decorations).</td>
</tr>
<tr>
<td class="name">GDK_WINDOW_STATE_TILED</td>
<td class="value">256</td>
<td class="doc">the window is in a tiled state, Since 3.10. Since 3.22.23, this
                         is deprecated in favor of per-edge information.</td>
</tr>
<tr>
<td class="name">GDK_WINDOW_STATE_TOP_TILED</td>
<td class="value">512</td>
<td class="doc">whether the top edge is tiled, Since 3.22.23</td>
</tr>
<tr>
<td class="name">GDK_WINDOW_STATE_TOP_RESIZABLE</td>
<td class="value">1024</td>
<td class="doc">whether the top edge is resizable, Since 3.22.23</td>
</tr>
<tr>
<td class="name">GDK_WINDOW_STATE_RIGHT_TILED</td>
<td class="value">2048</td>
<td class="doc">whether the right edge is tiled, Since 3.22.23</td>
</tr>
<tr>
<td class="name">GDK_WINDOW_STATE_RIGHT_RESIZABLE</td>
<td class="value">4096</td>
<td class="doc">whether the right edge is resizable, Since 3.22.23</td>
</tr>
<tr>
<td class="name">GDK_WINDOW_STATE_BOTTOM_TILED</td>
<td class="value">8192</td>
<td class="doc">whether the bottom edge is tiled, Since 3.22.23</td>
</tr>
<tr>
<td class="name">GDK_WINDOW_STATE_BOTTOM_RESIZABLE</td>
<td class="value">16384</td>
<td class="doc">whether the bottom edge is resizable, Since 3.22.23</td>
</tr>
<tr>
<td class="name">GDK_WINDOW_STATE_LEFT_TILED</td>
<td class="value">32768</td>
<td class="doc">whether the left edge is tiled, Since 3.22.23</td>
</tr>
<tr>
<td class="name">GDK_WINDOW_STATE_LEFT_RESIZABLE</td>
<td class="value">65536</td>
<td class="doc">whether the left edge is resizable, Since 3.22.23</td>
</tr>
</table>
<div class="api-notes">
  <p class="api-ctype">C type: GdkWindowState</p>
</div>
