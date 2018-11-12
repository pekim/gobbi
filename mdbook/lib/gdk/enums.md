# `gdk` enums

## `AxisUse`

An enumeration describing the way in which a device
axis (valuator) maps onto the predefined valuator
types that GTK+ understands.

Note that the X and Y axes are not really needed; pointer devices
report their location via the x/y members of events regardless. Whether
X and Y are present as axes depends on the GDK backend.

C - `GdkAxisUse`

## `ByteOrder`

A set of values describing the possible byte-orders
for storing pixel values in memory.

C - `GdkByteOrder`

## `CrossingMode`

Specifies the crossing mode for #GdkEventCrossing.

C - `GdkCrossingMode`

## `CursorType`

Predefined cursors.

Note that these IDs are directly taken from the X cursor font, and many
of these cursors are either not useful, or are not available on other platforms.

The recommended way to create cursors is to use gdk_cursor_new_from_name().

C - `GdkCursorType`

## `DevicePadFeature`

A pad feature.

C - `GdkDevicePadFeature`

## `DeviceToolType`

Indicates the specific type of tool being used being a tablet. Such as an
airbrush, pencil, etc.

C - `GdkDeviceToolType`

## `DeviceType`

Indicates the device type. See [above][GdkDeviceManager.description]
for more information about the meaning of these device types.

C - `GdkDeviceType`

## `DragCancelReason`

Used in #GdkDragContext to the reason of a cancelled DND operation.

C - `GdkDragCancelReason`

## `DragProtocol`

Used in #GdkDragContext to indicate the protocol according to
which DND is done.

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

C - `GdkEventType`

## `FilterReturn`

Specifies the result of applying a #GdkFilterFunc to a native event.

C - `GdkFilterReturn`

## `FullscreenMode`

Indicates which monitor (in a multi-head setup) a window should span over
when in fullscreen mode.

C - `GdkFullscreenMode`

## `GLError`

Error enumeration for #GdkGLContext.

C - `GdkGLError`

## `GrabOwnership`

Defines how device grabs interact with other devices.

C - `GdkGrabOwnership`

## `GrabStatus`

Returned by gdk_device_grab(), gdk_pointer_grab() and gdk_keyboard_grab() to
indicate success or the reason for the failure of the grab attempt.

C - `GdkGrabStatus`

## `Gravity`

Defines the reference point of a window and the meaning of coordinates
passed to gtk_window_move(). See gtk_window_move() and the "implementation
notes" section of the
[Extended Window Manager Hints](http://www.freedesktop.org/Standards/wm-spec)
specification for more details.

C - `GdkGravity`

## `InputMode`

An enumeration that describes the mode of an input device.

C - `GdkInputMode`

## `InputSource`

An enumeration describing the type of an input device in general terms.

C - `GdkInputSource`

## `ModifierIntent`

This enum is used with gdk_keymap_get_modifier_mask()
in order to determine what modifiers the
currently used windowing system backend uses for particular
purposes. For example, on X11/Windows, the Control key is used for
invoking menu shortcuts (accelerators), whereas on Apple computers
it’s the Command key (which correspond to %GDK_CONTROL_MASK and
%GDK_MOD2_MASK, respectively).

C - `GdkModifierIntent`

## `NotifyType`

Specifies the kind of crossing for #GdkEventCrossing.

See the X11 protocol specification of LeaveNotify for
full details of crossing event generation.

C - `GdkNotifyType`

## `OwnerChange`

Specifies why a selection ownership was changed.

C - `GdkOwnerChange`

## `PropMode`

Describes how existing data is combined with new data when
using gdk_property_change().

C - `GdkPropMode`

## `PropertyState`

Specifies the type of a property change for a #GdkEventProperty.

C - `guint`

## `ScrollDirection`

Specifies the direction for #GdkEventScroll.

C - `GdkScrollDirection`

## `SettingAction`

Specifies the kind of modification applied to a setting in a
&num;GdkEventSetting.

C - `GdkSettingAction`

## `Status`



C - `GdkStatus`

## `SubpixelLayout`

This enumeration describes how the red, green and blue components
of physical pixels on an output device are laid out.

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

C - `GdkTouchpadGesturePhase`

## `VisibilityState`

Specifies the visiblity status of a window for a #GdkEventVisibility.

C - `GdkVisibilityState`

## `VisualType`

A set of values that describe the manner in which the pixel values
for a visual are converted into RGB values for display.

C - `GdkVisualType`

## `WindowEdge`

Determines a window edge or corner.

C - `GdkWindowEdge`

## `WindowType`

Describes the kind of window.

C - `GdkWindowType`

## `WindowTypeHint`

These are hints for the window manager that indicate what type of function
the window has. The window manager can use this when determining decoration
and behaviour of the window. The hint must be set before mapping the window.

See the [Extended Window Manager Hints](http://www.freedesktop.org/Standards/wm-spec)
specification for more details about window types.

C - `GdkWindowTypeHint`

## `WindowWindowClass`

@GDK_INPUT_OUTPUT windows are the standard kind of window you might expect.
Such windows receive events and are also displayed on screen.
@GDK_INPUT_ONLY windows are invisible; they are usually placed above other
windows in order to trap or filter the events. You can’t draw on
@GDK_INPUT_ONLY windows.

C - `GdkWindowWindowClass`

