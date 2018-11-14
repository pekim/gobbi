# `gdk` enums

## `AxisUse`

An enumeration describing the way in which a device
axis (valuator) maps onto the predefined valuator
types that GTK+ understands.

Note that the X and Y axes are not really needed; pointer devices
report their location via the x/y members of events regardless. Whether
X and Y are present as axes depends on the GDK backend.

C - `GdkAxisUse`

### GDK_AXIS_IGNORE = 0
the axis is ignored.

### GDK_AXIS_X = 1
the axis is used as the x axis.

### GDK_AXIS_Y = 2
the axis is used as the y axis.

### GDK_AXIS_PRESSURE = 3
the axis is used for pressure information.

### GDK_AXIS_XTILT = 4
the axis is used for x tilt information.

### GDK_AXIS_YTILT = 5
the axis is used for y tilt information.

### GDK_AXIS_WHEEL = 6
the axis is used for wheel information.

### GDK_AXIS_DISTANCE = 7
the axis is used for pen/tablet distance information. (Since: 3.22)

### GDK_AXIS_ROTATION = 8
the axis is used for pen rotation information. (Since: 3.22)

### GDK_AXIS_SLIDER = 9
the axis is used for pen slider information. (Since: 3.22)

### GDK_AXIS_LAST = 10
a constant equal to the numerically highest axis value.


## `ByteOrder`

A set of values describing the possible byte-orders
for storing pixel values in memory.

C - `GdkByteOrder`

### GDK_LSB_FIRST = 0
The values are stored with the least-significant byte
  first. For instance, the 32-bit value 0xffeecc would be stored
  in memory as 0xcc, 0xee, 0xff, 0x00.

### GDK_MSB_FIRST = 1
The values are stored with the most-significant byte
  first. For instance, the 32-bit value 0xffeecc would be stored
  in memory as 0x00, 0xff, 0xee, 0xcc.


## `CrossingMode`

Specifies the crossing mode for #GdkEventCrossing.

C - `GdkCrossingMode`

### GDK_CROSSING_NORMAL = 0
crossing because of pointer motion.

### GDK_CROSSING_GRAB = 1
crossing because a grab is activated.

### GDK_CROSSING_UNGRAB = 2
crossing because a grab is deactivated.

### GDK_CROSSING_GTK_GRAB = 3
crossing because a GTK+ grab is activated.

### GDK_CROSSING_GTK_UNGRAB = 4
crossing because a GTK+ grab is deactivated.

### GDK_CROSSING_STATE_CHANGED = 5
crossing because a GTK+ widget changed
  state (e.g. sensitivity).

### GDK_CROSSING_TOUCH_BEGIN = 6
crossing because a touch sequence has begun,
  this event is synthetic as the pointer might have not left the window.

### GDK_CROSSING_TOUCH_END = 7
crossing because a touch sequence has ended,
  this event is synthetic as the pointer might have not left the window.

### GDK_CROSSING_DEVICE_SWITCH = 8
crossing because of a device switch (i.e.
  a mouse taking control of the pointer after a touch device), this event
  is synthetic as the pointer didn’t leave the window.


## `CursorType`

Predefined cursors.

Note that these IDs are directly taken from the X cursor font, and many
of these cursors are either not useful, or are not available on other platforms.

The recommended way to create cursors is to use gdk_cursor_new_from_name().

C - `GdkCursorType`

### GDK_X_CURSOR = 0
![](X_cursor.png)

### GDK_ARROW = 2
![](arrow.png)

### GDK_BASED_ARROW_DOWN = 4
![](based_arrow_down.png)

### GDK_BASED_ARROW_UP = 6
![](based_arrow_up.png)

### GDK_BOAT = 8
![](boat.png)

### GDK_BOGOSITY = 10
![](bogosity.png)

### GDK_BOTTOM_LEFT_CORNER = 12
![](bottom_left_corner.png)

### GDK_BOTTOM_RIGHT_CORNER = 14
![](bottom_right_corner.png)

### GDK_BOTTOM_SIDE = 16
![](bottom_side.png)

### GDK_BOTTOM_TEE = 18
![](bottom_tee.png)

### GDK_BOX_SPIRAL = 20
![](box_spiral.png)

### GDK_CENTER_PTR = 22
![](center_ptr.png)

### GDK_CIRCLE = 24
![](circle.png)

### GDK_CLOCK = 26
![](clock.png)

### GDK_COFFEE_MUG = 28
![](coffee_mug.png)

### GDK_CROSS = 30
![](cross.png)

### GDK_CROSS_REVERSE = 32
![](cross_reverse.png)

### GDK_CROSSHAIR = 34
![](crosshair.png)

### GDK_DIAMOND_CROSS = 36
![](diamond_cross.png)

### GDK_DOT = 38
![](dot.png)

### GDK_DOTBOX = 40
![](dotbox.png)

### GDK_DOUBLE_ARROW = 42
![](double_arrow.png)

### GDK_DRAFT_LARGE = 44
![](draft_large.png)

### GDK_DRAFT_SMALL = 46
![](draft_small.png)

### GDK_DRAPED_BOX = 48
![](draped_box.png)

### GDK_EXCHANGE = 50
![](exchange.png)

### GDK_FLEUR = 52
![](fleur.png)

### GDK_GOBBLER = 54
![](gobbler.png)

### GDK_GUMBY = 56
![](gumby.png)

### GDK_HAND1 = 58
![](hand1.png)

### GDK_HAND2 = 60
![](hand2.png)

### GDK_HEART = 62
![](heart.png)

### GDK_ICON = 64
![](icon.png)

### GDK_IRON_CROSS = 66
![](iron_cross.png)

### GDK_LEFT_PTR = 68
![](left_ptr.png)

### GDK_LEFT_SIDE = 70
![](left_side.png)

### GDK_LEFT_TEE = 72
![](left_tee.png)

### GDK_LEFTBUTTON = 74
![](leftbutton.png)

### GDK_LL_ANGLE = 76
![](ll_angle.png)

### GDK_LR_ANGLE = 78
![](lr_angle.png)

### GDK_MAN = 80
![](man.png)

### GDK_MIDDLEBUTTON = 82
![](middlebutton.png)

### GDK_MOUSE = 84
![](mouse.png)

### GDK_PENCIL = 86
![](pencil.png)

### GDK_PIRATE = 88
![](pirate.png)

### GDK_PLUS = 90
![](plus.png)

### GDK_QUESTION_ARROW = 92
![](question_arrow.png)

### GDK_RIGHT_PTR = 94
![](right_ptr.png)

### GDK_RIGHT_SIDE = 96
![](right_side.png)

### GDK_RIGHT_TEE = 98
![](right_tee.png)

### GDK_RIGHTBUTTON = 100
![](rightbutton.png)

### GDK_RTL_LOGO = 102
![](rtl_logo.png)

### GDK_SAILBOAT = 104
![](sailboat.png)

### GDK_SB_DOWN_ARROW = 106
![](sb_down_arrow.png)

### GDK_SB_H_DOUBLE_ARROW = 108
![](sb_h_double_arrow.png)

### GDK_SB_LEFT_ARROW = 110
![](sb_left_arrow.png)

### GDK_SB_RIGHT_ARROW = 112
![](sb_right_arrow.png)

### GDK_SB_UP_ARROW = 114
![](sb_up_arrow.png)

### GDK_SB_V_DOUBLE_ARROW = 116
![](sb_v_double_arrow.png)

### GDK_SHUTTLE = 118
![](shuttle.png)

### GDK_SIZING = 120
![](sizing.png)

### GDK_SPIDER = 122
![](spider.png)

### GDK_SPRAYCAN = 124
![](spraycan.png)

### GDK_STAR = 126
![](star.png)

### GDK_TARGET = 128
![](target.png)

### GDK_TCROSS = 130
![](tcross.png)

### GDK_TOP_LEFT_ARROW = 132
![](top_left_arrow.png)

### GDK_TOP_LEFT_CORNER = 134
![](top_left_corner.png)

### GDK_TOP_RIGHT_CORNER = 136
![](top_right_corner.png)

### GDK_TOP_SIDE = 138
![](top_side.png)

### GDK_TOP_TEE = 140
![](top_tee.png)

### GDK_TREK = 142
![](trek.png)

### GDK_UL_ANGLE = 144
![](ul_angle.png)

### GDK_UMBRELLA = 146
![](umbrella.png)

### GDK_UR_ANGLE = 148
![](ur_angle.png)

### GDK_WATCH = 150
![](watch.png)

### GDK_XTERM = 152
![](xterm.png)

### GDK_LAST_CURSOR = 153
last cursor type

### GDK_BLANK_CURSOR = -2
Blank cursor. Since 2.16

### GDK_CURSOR_IS_PIXMAP = -1
type of cursors constructed with
  gdk_cursor_new_from_pixbuf()


## `DevicePadFeature`

A pad feature.

C - `GdkDevicePadFeature`

### GDK_DEVICE_PAD_FEATURE_BUTTON = 0
a button

### GDK_DEVICE_PAD_FEATURE_RING = 1
a ring-shaped interactive area

### GDK_DEVICE_PAD_FEATURE_STRIP = 2
a straight interactive area


## `DeviceToolType`

Indicates the specific type of tool being used being a tablet. Such as an
airbrush, pencil, etc.

C - `GdkDeviceToolType`

### GDK_DEVICE_TOOL_TYPE_UNKNOWN = 0
Tool is of an unknown type.

### GDK_DEVICE_TOOL_TYPE_PEN = 1
Tool is a standard tablet stylus.

### GDK_DEVICE_TOOL_TYPE_ERASER = 2
Tool is standard tablet eraser.

### GDK_DEVICE_TOOL_TYPE_BRUSH = 3
Tool is a brush stylus.

### GDK_DEVICE_TOOL_TYPE_PENCIL = 4
Tool is a pencil stylus.

### GDK_DEVICE_TOOL_TYPE_AIRBRUSH = 5
Tool is an airbrush stylus.

### GDK_DEVICE_TOOL_TYPE_MOUSE = 6
Tool is a mouse.

### GDK_DEVICE_TOOL_TYPE_LENS = 7
Tool is a lens cursor.


## `DeviceType`

Indicates the device type. See [above][GdkDeviceManager.description]
for more information about the meaning of these device types.

C - `GdkDeviceType`

### GDK_DEVICE_TYPE_MASTER = 0
Device is a master (or virtual) device. There will
                         be an associated focus indicator on the screen.

### GDK_DEVICE_TYPE_SLAVE = 1
Device is a slave (or physical) device.

### GDK_DEVICE_TYPE_FLOATING = 2
Device is a physical device, currently not attached to
                           any virtual device.


## `DragCancelReason`

Used in #GdkDragContext to the reason of a cancelled DND operation.

C - `GdkDragCancelReason`

### GDK_DRAG_CANCEL_NO_TARGET = 0
There is no suitable drop target.

### GDK_DRAG_CANCEL_USER_CANCELLED = 1
Drag cancelled by the user

### GDK_DRAG_CANCEL_ERROR = 2
Unspecified error.


## `DragProtocol`

Used in #GdkDragContext to indicate the protocol according to
which DND is done.

C - `GdkDragProtocol`

### GDK_DRAG_PROTO_NONE = 0
no protocol.

### GDK_DRAG_PROTO_MOTIF = 1
The Motif DND protocol. No longer supported

### GDK_DRAG_PROTO_XDND = 2
The Xdnd protocol.

### GDK_DRAG_PROTO_ROOTWIN = 3
An extension to the Xdnd protocol for
 unclaimed root window drops.

### GDK_DRAG_PROTO_WIN32_DROPFILES = 4
The simple WM_DROPFILES protocol.

### GDK_DRAG_PROTO_OLE2 = 5
The complex OLE2 DND protocol (not implemented).

### GDK_DRAG_PROTO_LOCAL = 6
Intra-application DND.

### GDK_DRAG_PROTO_WAYLAND = 7
Wayland DND protocol.


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

### GDK_NOTHING = -1
a special code to indicate a null event.

### GDK_DELETE = 0
the window manager has requested that the toplevel window be
  hidden or destroyed, usually when the user clicks on a special icon in the
  title bar.

### GDK_DESTROY = 1
the window has been destroyed.

### GDK_EXPOSE = 2
all or part of the window has become visible and needs to be
  redrawn.

### GDK_MOTION_NOTIFY = 3
the pointer (usually a mouse) has moved.

### GDK_BUTTON_PRESS = 4
a mouse button has been pressed.

### GDK_2BUTTON_PRESS = 5
a mouse button has been double-clicked (clicked twice
  within a short period of time). Note that each click also generates a
  %GDK_BUTTON_PRESS event.

### GDK_DOUBLE_BUTTON_PRESS = 5
alias for %GDK_2BUTTON_PRESS, added in 3.6.

### GDK_3BUTTON_PRESS = 6
a mouse button has been clicked 3 times in a short period
  of time. Note that each click also generates a %GDK_BUTTON_PRESS event.

### GDK_TRIPLE_BUTTON_PRESS = 6
alias for %GDK_3BUTTON_PRESS, added in 3.6.

### GDK_BUTTON_RELEASE = 7
a mouse button has been released.

### GDK_KEY_PRESS = 8
a key has been pressed.

### GDK_KEY_RELEASE = 9
a key has been released.

### GDK_ENTER_NOTIFY = 10
the pointer has entered the window.

### GDK_LEAVE_NOTIFY = 11
the pointer has left the window.

### GDK_FOCUS_CHANGE = 12
the keyboard focus has entered or left the window.

### GDK_CONFIGURE = 13
the size, position or stacking order of the window has changed.
  Note that GTK+ discards these events for %GDK_WINDOW_CHILD windows.

### GDK_MAP = 14
the window has been mapped.

### GDK_UNMAP = 15
the window has been unmapped.

### GDK_PROPERTY_NOTIFY = 16
a property on the window has been changed or deleted.

### GDK_SELECTION_CLEAR = 17
the application has lost ownership of a selection.

### GDK_SELECTION_REQUEST = 18
another application has requested a selection.

### GDK_SELECTION_NOTIFY = 19
a selection has been received.

### GDK_PROXIMITY_IN = 20
an input device has moved into contact with a sensing
  surface (e.g. a touchscreen or graphics tablet).

### GDK_PROXIMITY_OUT = 21
an input device has moved out of contact with a sensing
  surface.

### GDK_DRAG_ENTER = 22
the mouse has entered the window while a drag is in progress.

### GDK_DRAG_LEAVE = 23
the mouse has left the window while a drag is in progress.

### GDK_DRAG_MOTION = 24
the mouse has moved in the window while a drag is in
  progress.

### GDK_DRAG_STATUS = 25
the status of the drag operation initiated by the window
  has changed.

### GDK_DROP_START = 26
a drop operation onto the window has started.

### GDK_DROP_FINISHED = 27
the drop operation initiated by the window has completed.

### GDK_CLIENT_EVENT = 28
a message has been received from another application.

### GDK_VISIBILITY_NOTIFY = 29
the window visibility status has changed.

### GDK_SCROLL = 31
the scroll wheel was turned

### GDK_WINDOW_STATE = 32
the state of a window has changed. See #GdkWindowState
  for the possible window states

### GDK_SETTING = 33
a setting has been modified.

### GDK_OWNER_CHANGE = 34
the owner of a selection has changed. This event type
  was added in 2.6

### GDK_GRAB_BROKEN = 35
a pointer or keyboard grab was broken. This event type
  was added in 2.8.

### GDK_DAMAGE = 36
the content of the window has been changed. This event type
  was added in 2.14.

### GDK_TOUCH_BEGIN = 37
A new touch event sequence has just started. This event
  type was added in 3.4.

### GDK_TOUCH_UPDATE = 38
A touch event sequence has been updated. This event type
  was added in 3.4.

### GDK_TOUCH_END = 39
A touch event sequence has finished. This event type
  was added in 3.4.

### GDK_TOUCH_CANCEL = 40
A touch event sequence has been canceled. This event type
  was added in 3.4.

### GDK_TOUCHPAD_SWIPE = 41
A touchpad swipe gesture event, the current state
  is determined by its phase field. This event type was added in 3.18.

### GDK_TOUCHPAD_PINCH = 42
A touchpad pinch gesture event, the current state
  is determined by its phase field. This event type was added in 3.18.

### GDK_PAD_BUTTON_PRESS = 43
A tablet pad button press event. This event type
  was added in 3.22.

### GDK_PAD_BUTTON_RELEASE = 44
A tablet pad button release event. This event type
  was added in 3.22.

### GDK_PAD_RING = 45
A tablet pad axis event from a "ring". This event type was
  added in 3.22.

### GDK_PAD_STRIP = 46
A tablet pad axis event from a "strip". This event type was
  added in 3.22.

### GDK_PAD_GROUP_MODE = 47
A tablet pad group mode change. This event type was
  added in 3.22.

### GDK_EVENT_LAST = 48
marks the end of the GdkEventType enumeration. Added in 2.18


## `FilterReturn`

Specifies the result of applying a #GdkFilterFunc to a native event.

C - `GdkFilterReturn`

### GDK_FILTER_CONTINUE = 0
event not handled, continue processing.

### GDK_FILTER_TRANSLATE = 1
native event translated into a GDK event and stored
 in the `event` structure that was passed in.

### GDK_FILTER_REMOVE = 2
event handled, terminate processing.


## `FullscreenMode`

Indicates which monitor (in a multi-head setup) a window should span over
when in fullscreen mode.

C - `GdkFullscreenMode`

### GDK_FULLSCREEN_ON_CURRENT_MONITOR = 0
Fullscreen on current monitor only.

### GDK_FULLSCREEN_ON_ALL_MONITORS = 1
Span across all monitors when fullscreen.


## `GLError`

Error enumeration for #GdkGLContext.

C - `GdkGLError`

### GDK_GL_ERROR_NOT_AVAILABLE = 0
OpenGL support is not available

### GDK_GL_ERROR_UNSUPPORTED_FORMAT = 1
The requested visual format is not supported

### GDK_GL_ERROR_UNSUPPORTED_PROFILE = 2
The requested profile is not supported


## `GrabOwnership`

Defines how device grabs interact with other devices.

C - `GdkGrabOwnership`

### GDK_OWNERSHIP_NONE = 0
All other devices’ events are allowed.

### GDK_OWNERSHIP_WINDOW = 1
Other devices’ events are blocked for the grab window.

### GDK_OWNERSHIP_APPLICATION = 2
Other devices’ events are blocked for the whole application.


## `GrabStatus`

Returned by gdk_device_grab(), gdk_pointer_grab() and gdk_keyboard_grab() to
indicate success or the reason for the failure of the grab attempt.

C - `GdkGrabStatus`

### GDK_GRAB_SUCCESS = 0
the resource was successfully grabbed.

### GDK_GRAB_ALREADY_GRABBED = 1
the resource is actively grabbed by another client.

### GDK_GRAB_INVALID_TIME = 2
the resource was grabbed more recently than the
 specified time.

### GDK_GRAB_NOT_VIEWABLE = 3
the grab window or the @confine_to window are not
 viewable.

### GDK_GRAB_FROZEN = 4
the resource is frozen by an active grab of another client.

### GDK_GRAB_FAILED = 5
the grab failed for some other reason. Since 3.16


## `Gravity`

Defines the reference point of a window and the meaning of coordinates
passed to gtk_window_move(). See gtk_window_move() and the "implementation
notes" section of the
[Extended Window Manager Hints](http://www.freedesktop.org/Standards/wm-spec)
specification for more details.

C - `GdkGravity`

### GDK_GRAVITY_NORTH_WEST = 1
the reference point is at the top left corner.

### GDK_GRAVITY_NORTH = 2
the reference point is in the middle of the top edge.

### GDK_GRAVITY_NORTH_EAST = 3
the reference point is at the top right corner.

### GDK_GRAVITY_WEST = 4
the reference point is at the middle of the left edge.

### GDK_GRAVITY_CENTER = 5
the reference point is at the center of the window.

### GDK_GRAVITY_EAST = 6
the reference point is at the middle of the right edge.

### GDK_GRAVITY_SOUTH_WEST = 7
the reference point is at the lower left corner.

### GDK_GRAVITY_SOUTH = 8
the reference point is at the middle of the lower edge.

### GDK_GRAVITY_SOUTH_EAST = 9
the reference point is at the lower right corner.

### GDK_GRAVITY_STATIC = 10
the reference point is at the top left corner of the
 window itself, ignoring window manager decorations.


## `InputMode`

An enumeration that describes the mode of an input device.

C - `GdkInputMode`

### GDK_MODE_DISABLED = 0
the device is disabled and will not report any events.

### GDK_MODE_SCREEN = 1
the device is enabled. The device’s coordinate space
                  maps to the entire screen.

### GDK_MODE_WINDOW = 2
the device is enabled. The device’s coordinate space
                  is mapped to a single window. The manner in which this window
                  is chosen is undefined, but it will typically be the same
                  way in which the focus window for key events is determined.


## `InputSource`

An enumeration describing the type of an input device in general terms.

C - `GdkInputSource`

### GDK_SOURCE_MOUSE = 0
the device is a mouse. (This will be reported for the core
                   pointer, even if it is something else, such as a trackball.)

### GDK_SOURCE_PEN = 1
the device is a stylus of a graphics tablet or similar device.

### GDK_SOURCE_ERASER = 2
the device is an eraser. Typically, this would be the other end
                    of a stylus on a graphics tablet.

### GDK_SOURCE_CURSOR = 3
the device is a graphics tablet “puck” or similar device.

### GDK_SOURCE_KEYBOARD = 4
the device is a keyboard.

### GDK_SOURCE_TOUCHSCREEN = 5
the device is a direct-input touch device, such
    as a touchscreen or tablet. This device type has been added in 3.4.

### GDK_SOURCE_TOUCHPAD = 6
the device is an indirect touch device, such
    as a touchpad. This device type has been added in 3.4.

### GDK_SOURCE_TRACKPOINT = 7
the device is a trackpoint. This device type has been
    added in 3.22

### GDK_SOURCE_TABLET_PAD = 8
the device is a "pad", a collection of buttons,
    rings and strips found in drawing tablets. This device type has been
    added in 3.22.


## `ModifierIntent`

This enum is used with gdk_keymap_get_modifier_mask()
in order to determine what modifiers the
currently used windowing system backend uses for particular
purposes. For example, on X11/Windows, the Control key is used for
invoking menu shortcuts (accelerators), whereas on Apple computers
it’s the Command key (which correspond to %GDK_CONTROL_MASK and
%GDK_MOD2_MASK, respectively).

C - `GdkModifierIntent`

### GDK_MODIFIER_INTENT_PRIMARY_ACCELERATOR = 0
the primary modifier used to invoke
 menu accelerators.

### GDK_MODIFIER_INTENT_CONTEXT_MENU = 1
the modifier used to invoke context menus.
 Note that mouse button 3 always triggers context menus. When this modifier
 is not 0, it additionally triggers context menus when used with mouse button 1.

### GDK_MODIFIER_INTENT_EXTEND_SELECTION = 2
the modifier used to extend selections
 using `modifier`-click or `modifier`-cursor-key

### GDK_MODIFIER_INTENT_MODIFY_SELECTION = 3
the modifier used to modify selections,
 which in most cases means toggling the clicked item into or out of the selection.

### GDK_MODIFIER_INTENT_NO_TEXT_INPUT = 4
when any of these modifiers is pressed, the
 key event cannot produce a symbol directly. This is meant to be used for
 input methods, and for use cases like typeahead search.

### GDK_MODIFIER_INTENT_SHIFT_GROUP = 5
the modifier that switches between keyboard
 groups (AltGr on X11/Windows and Option/Alt on OS X).

### GDK_MODIFIER_INTENT_DEFAULT_MOD_MASK = 6
The set of modifier masks accepted
as modifiers in accelerators. Needed because Command is mapped to MOD2 on
OSX, which is widely used, but on X11 MOD2 is NumLock and using that for a
mod key is problematic at best.
Ref: https://bugzilla.gnome.org/show_bug.cgi?id=736125.


## `NotifyType`

Specifies the kind of crossing for #GdkEventCrossing.

See the X11 protocol specification of LeaveNotify for
full details of crossing event generation.

C - `GdkNotifyType`

### GDK_NOTIFY_ANCESTOR = 0
the window is entered from an ancestor or
  left towards an ancestor.

### GDK_NOTIFY_VIRTUAL = 1
the pointer moves between an ancestor and an
  inferior of the window.

### GDK_NOTIFY_INFERIOR = 2
the window is entered from an inferior or
  left towards an inferior.

### GDK_NOTIFY_NONLINEAR = 3
the window is entered from or left towards
  a window which is neither an ancestor nor an inferior.

### GDK_NOTIFY_NONLINEAR_VIRTUAL = 4
the pointer moves between two windows
  which are not ancestors of each other and the window is part of
  the ancestor chain between one of these windows and their least
  common ancestor.

### GDK_NOTIFY_UNKNOWN = 5
an unknown type of enter/leave event occurred.


## `OwnerChange`

Specifies why a selection ownership was changed.

C - `GdkOwnerChange`

### GDK_OWNER_CHANGE_NEW_OWNER = 0
some other app claimed the ownership

### GDK_OWNER_CHANGE_DESTROY = 1
the window was destroyed

### GDK_OWNER_CHANGE_CLOSE = 2
the client was closed


## `PropMode`

Describes how existing data is combined with new data when
using gdk_property_change().

C - `GdkPropMode`

### GDK_PROP_MODE_REPLACE = 0
the new data replaces the existing data.

### GDK_PROP_MODE_PREPEND = 1
the new data is prepended to the existing data.

### GDK_PROP_MODE_APPEND = 2
the new data is appended to the existing data.


## `PropertyState`

Specifies the type of a property change for a #GdkEventProperty.

C - `guint`

### GDK_PROPERTY_NEW_VALUE = 0
the property value was changed.

### GDK_PROPERTY_DELETE = 1
the property was deleted.


## `ScrollDirection`

Specifies the direction for #GdkEventScroll.

C - `GdkScrollDirection`

### GDK_SCROLL_UP = 0
the window is scrolled up.

### GDK_SCROLL_DOWN = 1
the window is scrolled down.

### GDK_SCROLL_LEFT = 2
the window is scrolled to the left.

### GDK_SCROLL_RIGHT = 3
the window is scrolled to the right.

### GDK_SCROLL_SMOOTH = 4
the scrolling is determined by the delta values
  in #GdkEventScroll. See gdk_event_get_scroll_deltas(). Since: 3.4


## `SettingAction`

Specifies the kind of modification applied to a setting in a
&num;GdkEventSetting.

C - `GdkSettingAction`

### GDK_SETTING_ACTION_NEW = 0
a setting was added.

### GDK_SETTING_ACTION_CHANGED = 1
a setting was changed.

### GDK_SETTING_ACTION_DELETED = 2
a setting was deleted.


## `Status`



C - `GdkStatus`

### GDK_OK = 0


### GDK_ERROR = -1


### GDK_ERROR_PARAM = -2


### GDK_ERROR_FILE = -3


### GDK_ERROR_MEM = -4



## `SubpixelLayout`

This enumeration describes how the red, green and blue components
of physical pixels on an output device are laid out.

C - `GdkSubpixelLayout`

### GDK_SUBPIXEL_LAYOUT_UNKNOWN = 0
The layout is not known

### GDK_SUBPIXEL_LAYOUT_NONE = 1
Not organized in this way

### GDK_SUBPIXEL_LAYOUT_HORIZONTAL_RGB = 2
The layout is horizontal, the order is RGB

### GDK_SUBPIXEL_LAYOUT_HORIZONTAL_BGR = 3
The layout is horizontal, the order is BGR

### GDK_SUBPIXEL_LAYOUT_VERTICAL_RGB = 4
The layout is vertical, the order is RGB

### GDK_SUBPIXEL_LAYOUT_VERTICAL_BGR = 5
The layout is vertical, the order is BGR


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

### GDK_TOUCHPAD_GESTURE_PHASE_BEGIN = 0
The gesture has begun.

### GDK_TOUCHPAD_GESTURE_PHASE_UPDATE = 1
The gesture has been updated.

### GDK_TOUCHPAD_GESTURE_PHASE_END = 2
The gesture was finished, changes
  should be permanently applied.

### GDK_TOUCHPAD_GESTURE_PHASE_CANCEL = 3
The gesture was cancelled, all
  changes should be undone.


## `VisibilityState`

Specifies the visiblity status of a window for a #GdkEventVisibility.

C - `GdkVisibilityState`

### GDK_VISIBILITY_UNOBSCURED = 0
the window is completely visible.

### GDK_VISIBILITY_PARTIAL = 1
the window is partially visible.

### GDK_VISIBILITY_FULLY_OBSCURED = 2
the window is not visible at all.


## `VisualType`

A set of values that describe the manner in which the pixel values
for a visual are converted into RGB values for display.

C - `GdkVisualType`

### GDK_VISUAL_STATIC_GRAY = 0
Each pixel value indexes a grayscale value
    directly.

### GDK_VISUAL_GRAYSCALE = 1
Each pixel is an index into a color map that
    maps pixel values into grayscale values. The color map can be
    changed by an application.

### GDK_VISUAL_STATIC_COLOR = 2
Each pixel value is an index into a predefined,
    unmodifiable color map that maps pixel values into RGB values.

### GDK_VISUAL_PSEUDO_COLOR = 3
Each pixel is an index into a color map that
    maps pixel values into rgb values. The color map can be changed by
    an application.

### GDK_VISUAL_TRUE_COLOR = 4
Each pixel value directly contains red, green,
    and blue components. Use gdk_visual_get_red_pixel_details(), etc,
    to obtain information about how the components are assembled into
    a pixel value.

### GDK_VISUAL_DIRECT_COLOR = 5
Each pixel value contains red, green, and blue
    components as for %GDK_VISUAL_TRUE_COLOR, but the components are
    mapped via a color table into the final output table instead of
    being converted directly.


## `WindowEdge`

Determines a window edge or corner.

C - `GdkWindowEdge`

### GDK_WINDOW_EDGE_NORTH_WEST = 0
the top left corner.

### GDK_WINDOW_EDGE_NORTH = 1
the top edge.

### GDK_WINDOW_EDGE_NORTH_EAST = 2
the top right corner.

### GDK_WINDOW_EDGE_WEST = 3
the left edge.

### GDK_WINDOW_EDGE_EAST = 4
the right edge.

### GDK_WINDOW_EDGE_SOUTH_WEST = 5
the lower left corner.

### GDK_WINDOW_EDGE_SOUTH = 6
the lower edge.

### GDK_WINDOW_EDGE_SOUTH_EAST = 7
the lower right corner.


## `WindowType`

Describes the kind of window.

C - `GdkWindowType`

### GDK_WINDOW_ROOT = 0
root window; this window has no parent, covers the entire
 screen, and is created by the window system

### GDK_WINDOW_TOPLEVEL = 1
toplevel window (used to implement #GtkWindow)

### GDK_WINDOW_CHILD = 2
child window (used to implement e.g. #GtkEntry)

### GDK_WINDOW_TEMP = 3
override redirect temporary window (used to implement
 #GtkMenu)

### GDK_WINDOW_FOREIGN = 4
foreign window (see gdk_window_foreign_new())

### GDK_WINDOW_OFFSCREEN = 5
offscreen window (see
 [Offscreen Windows][OFFSCREEN-WINDOWS]). Since 2.18

### GDK_WINDOW_SUBSURFACE = 6
subsurface-based window; This window is visually
 tied to a toplevel, and is moved/stacked with it. Currently this window
 type is only implemented in Wayland. Since 3.14


## `WindowTypeHint`

These are hints for the window manager that indicate what type of function
the window has. The window manager can use this when determining decoration
and behaviour of the window. The hint must be set before mapping the window.

See the [Extended Window Manager Hints](http://www.freedesktop.org/Standards/wm-spec)
specification for more details about window types.

C - `GdkWindowTypeHint`

### GDK_WINDOW_TYPE_HINT_NORMAL = 0
Normal toplevel window.

### GDK_WINDOW_TYPE_HINT_DIALOG = 1
Dialog window.

### GDK_WINDOW_TYPE_HINT_MENU = 2
Window used to implement a menu; GTK+ uses
 this hint only for torn-off menus, see #GtkTearoffMenuItem.

### GDK_WINDOW_TYPE_HINT_TOOLBAR = 3
Window used to implement toolbars.

### GDK_WINDOW_TYPE_HINT_SPLASHSCREEN = 4
Window used to display a splash
 screen during application startup.

### GDK_WINDOW_TYPE_HINT_UTILITY = 5
Utility windows which are not detached
 toolbars or dialogs.

### GDK_WINDOW_TYPE_HINT_DOCK = 6
Used for creating dock or panel windows.

### GDK_WINDOW_TYPE_HINT_DESKTOP = 7
Used for creating the desktop background
 window.

### GDK_WINDOW_TYPE_HINT_DROPDOWN_MENU = 8
A menu that belongs to a menubar.

### GDK_WINDOW_TYPE_HINT_POPUP_MENU = 9
A menu that does not belong to a menubar,
 e.g. a context menu.

### GDK_WINDOW_TYPE_HINT_TOOLTIP = 10
A tooltip.

### GDK_WINDOW_TYPE_HINT_NOTIFICATION = 11
A notification - typically a “bubble”
 that belongs to a status icon.

### GDK_WINDOW_TYPE_HINT_COMBO = 12
A popup from a combo box.

### GDK_WINDOW_TYPE_HINT_DND = 13
A window that is used to implement a DND cursor.


## `WindowWindowClass`

@GDK_INPUT_OUTPUT windows are the standard kind of window you might expect.
Such windows receive events and are also displayed on screen.
@GDK_INPUT_ONLY windows are invisible; they are usually placed above other
windows in order to trap or filter the events. You can’t draw on
@GDK_INPUT_ONLY windows.

C - `GdkWindowWindowClass`

### GDK_INPUT_OUTPUT = 0
window for graphics and events

### GDK_INPUT_ONLY = 1
window for events only


