// This is a generated file - DO NOT EDIT

package gdk

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <stdlib.h>
import "C"

// An enumeration describing the way in which a device
// axis (valuator) maps onto the predefined valuator
// types that GTK+ understands.
//
// Note that the X and Y axes are not really needed; pointer devices
// report their location via the x/y members of events regardless. Whether
// X and Y are present as axes depends on the GDK backend.
type AxisUse C.GdkAxisUse

const (
	// the axis is ignored.
	GDK_AXIS_IGNORE AxisUse = 0
	// the axis is used as the x axis.
	GDK_AXIS_X AxisUse = 1
	// the axis is used as the y axis.
	GDK_AXIS_Y AxisUse = 2
	// the axis is used for pressure information.
	GDK_AXIS_PRESSURE AxisUse = 3
	// the axis is used for x tilt information.
	GDK_AXIS_XTILT AxisUse = 4
	// the axis is used for y tilt information.
	GDK_AXIS_YTILT AxisUse = 5
	// the axis is used for wheel information.
	GDK_AXIS_WHEEL AxisUse = 6
	// the axis is used for pen/tablet distance information. (Since: 3.22)
	GDK_AXIS_DISTANCE AxisUse = 7
	// the axis is used for pen rotation information. (Since: 3.22)
	GDK_AXIS_ROTATION AxisUse = 8
	// the axis is used for pen slider information. (Since: 3.22)
	GDK_AXIS_SLIDER AxisUse = 9
	// a constant equal to the numerically highest axis value.
	GDK_AXIS_LAST AxisUse = 10
)

// A set of values describing the possible byte-orders
// for storing pixel values in memory.
type ByteOrder C.GdkByteOrder

const (
	// The values are stored with the least-significant byte
	// first. For instance, the 32-bit value 0xffeecc would be stored
	// in memory as 0xcc, 0xee, 0xff, 0x00.
	GDK_LSB_FIRST ByteOrder = 0
	// The values are stored with the most-significant byte
	// first. For instance, the 32-bit value 0xffeecc would be stored
	// in memory as 0x00, 0xff, 0xee, 0xcc.
	GDK_MSB_FIRST ByteOrder = 1
)

// Specifies the crossing mode for #GdkEventCrossing.
type CrossingMode C.GdkCrossingMode

const (
	// crossing because of pointer motion.
	GDK_CROSSING_NORMAL CrossingMode = 0
	// crossing because a grab is activated.
	GDK_CROSSING_GRAB CrossingMode = 1
	// crossing because a grab is deactivated.
	GDK_CROSSING_UNGRAB CrossingMode = 2
	// crossing because a GTK+ grab is activated.
	GDK_CROSSING_GTK_GRAB CrossingMode = 3
	// crossing because a GTK+ grab is deactivated.
	GDK_CROSSING_GTK_UNGRAB CrossingMode = 4
	// crossing because a GTK+ widget changed
	// state (e.g. sensitivity).
	GDK_CROSSING_STATE_CHANGED CrossingMode = 5
	// crossing because a touch sequence has begun,
	// this event is synthetic as the pointer might have not left the window.
	GDK_CROSSING_TOUCH_BEGIN CrossingMode = 6
	// crossing because a touch sequence has ended,
	// this event is synthetic as the pointer might have not left the window.
	GDK_CROSSING_TOUCH_END CrossingMode = 7
	// crossing because of a device switch (i.e.
	// a mouse taking control of the pointer after a touch device), this event
	// is synthetic as the pointer didn’t leave the window.
	GDK_CROSSING_DEVICE_SWITCH CrossingMode = 8
)

// Predefined cursors.
//
// Note that these IDs are directly taken from the X cursor font, and many
// of these cursors are either not useful, or are not available on other platforms.
//
// The recommended way to create cursors is to use gdk_cursor_new_from_name().
type CursorType C.GdkCursorType

const (
	// ![](X_cursor.png)
	GDK_X_CURSOR CursorType = 0
	// ![](arrow.png)
	GDK_ARROW CursorType = 2
	// ![](based_arrow_down.png)
	GDK_BASED_ARROW_DOWN CursorType = 4
	// ![](based_arrow_up.png)
	GDK_BASED_ARROW_UP CursorType = 6
	// ![](boat.png)
	GDK_BOAT CursorType = 8
	// ![](bogosity.png)
	GDK_BOGOSITY CursorType = 10
	// ![](bottom_left_corner.png)
	GDK_BOTTOM_LEFT_CORNER CursorType = 12
	// ![](bottom_right_corner.png)
	GDK_BOTTOM_RIGHT_CORNER CursorType = 14
	// ![](bottom_side.png)
	GDK_BOTTOM_SIDE CursorType = 16
	// ![](bottom_tee.png)
	GDK_BOTTOM_TEE CursorType = 18
	// ![](box_spiral.png)
	GDK_BOX_SPIRAL CursorType = 20
	// ![](center_ptr.png)
	GDK_CENTER_PTR CursorType = 22
	// ![](circle.png)
	GDK_CIRCLE CursorType = 24
	// ![](clock.png)
	GDK_CLOCK CursorType = 26
	// ![](coffee_mug.png)
	GDK_COFFEE_MUG CursorType = 28
	// ![](cross.png)
	GDK_CROSS CursorType = 30
	// ![](cross_reverse.png)
	GDK_CROSS_REVERSE CursorType = 32
	// ![](crosshair.png)
	GDK_CROSSHAIR CursorType = 34
	// ![](diamond_cross.png)
	GDK_DIAMOND_CROSS CursorType = 36
	// ![](dot.png)
	GDK_DOT CursorType = 38
	// ![](dotbox.png)
	GDK_DOTBOX CursorType = 40
	// ![](double_arrow.png)
	GDK_DOUBLE_ARROW CursorType = 42
	// ![](draft_large.png)
	GDK_DRAFT_LARGE CursorType = 44
	// ![](draft_small.png)
	GDK_DRAFT_SMALL CursorType = 46
	// ![](draped_box.png)
	GDK_DRAPED_BOX CursorType = 48
	// ![](exchange.png)
	GDK_EXCHANGE CursorType = 50
	// ![](fleur.png)
	GDK_FLEUR CursorType = 52
	// ![](gobbler.png)
	GDK_GOBBLER CursorType = 54
	// ![](gumby.png)
	GDK_GUMBY CursorType = 56
	// ![](hand1.png)
	GDK_HAND1 CursorType = 58
	// ![](hand2.png)
	GDK_HAND2 CursorType = 60
	// ![](heart.png)
	GDK_HEART CursorType = 62
	// ![](icon.png)
	GDK_ICON CursorType = 64
	// ![](iron_cross.png)
	GDK_IRON_CROSS CursorType = 66
	// ![](left_ptr.png)
	GDK_LEFT_PTR CursorType = 68
	// ![](left_side.png)
	GDK_LEFT_SIDE CursorType = 70
	// ![](left_tee.png)
	GDK_LEFT_TEE CursorType = 72
	// ![](leftbutton.png)
	GDK_LEFTBUTTON CursorType = 74
	// ![](ll_angle.png)
	GDK_LL_ANGLE CursorType = 76
	// ![](lr_angle.png)
	GDK_LR_ANGLE CursorType = 78
	// ![](man.png)
	GDK_MAN CursorType = 80
	// ![](middlebutton.png)
	GDK_MIDDLEBUTTON CursorType = 82
	// ![](mouse.png)
	GDK_MOUSE CursorType = 84
	// ![](pencil.png)
	GDK_PENCIL CursorType = 86
	// ![](pirate.png)
	GDK_PIRATE CursorType = 88
	// ![](plus.png)
	GDK_PLUS CursorType = 90
	// ![](question_arrow.png)
	GDK_QUESTION_ARROW CursorType = 92
	// ![](right_ptr.png)
	GDK_RIGHT_PTR CursorType = 94
	// ![](right_side.png)
	GDK_RIGHT_SIDE CursorType = 96
	// ![](right_tee.png)
	GDK_RIGHT_TEE CursorType = 98
	// ![](rightbutton.png)
	GDK_RIGHTBUTTON CursorType = 100
	// ![](rtl_logo.png)
	GDK_RTL_LOGO CursorType = 102
	// ![](sailboat.png)
	GDK_SAILBOAT CursorType = 104
	// ![](sb_down_arrow.png)
	GDK_SB_DOWN_ARROW CursorType = 106
	// ![](sb_h_double_arrow.png)
	GDK_SB_H_DOUBLE_ARROW CursorType = 108
	// ![](sb_left_arrow.png)
	GDK_SB_LEFT_ARROW CursorType = 110
	// ![](sb_right_arrow.png)
	GDK_SB_RIGHT_ARROW CursorType = 112
	// ![](sb_up_arrow.png)
	GDK_SB_UP_ARROW CursorType = 114
	// ![](sb_v_double_arrow.png)
	GDK_SB_V_DOUBLE_ARROW CursorType = 116
	// ![](shuttle.png)
	GDK_SHUTTLE CursorType = 118
	// ![](sizing.png)
	GDK_SIZING CursorType = 120
	// ![](spider.png)
	GDK_SPIDER CursorType = 122
	// ![](spraycan.png)
	GDK_SPRAYCAN CursorType = 124
	// ![](star.png)
	GDK_STAR CursorType = 126
	// ![](target.png)
	GDK_TARGET CursorType = 128
	// ![](tcross.png)
	GDK_TCROSS CursorType = 130
	// ![](top_left_arrow.png)
	GDK_TOP_LEFT_ARROW CursorType = 132
	// ![](top_left_corner.png)
	GDK_TOP_LEFT_CORNER CursorType = 134
	// ![](top_right_corner.png)
	GDK_TOP_RIGHT_CORNER CursorType = 136
	// ![](top_side.png)
	GDK_TOP_SIDE CursorType = 138
	// ![](top_tee.png)
	GDK_TOP_TEE CursorType = 140
	// ![](trek.png)
	GDK_TREK CursorType = 142
	// ![](ul_angle.png)
	GDK_UL_ANGLE CursorType = 144
	// ![](umbrella.png)
	GDK_UMBRELLA CursorType = 146
	// ![](ur_angle.png)
	GDK_UR_ANGLE CursorType = 148
	// ![](watch.png)
	GDK_WATCH CursorType = 150
	// ![](xterm.png)
	GDK_XTERM CursorType = 152
	// last cursor type
	GDK_LAST_CURSOR CursorType = 153
	// Blank cursor. Since 2.16
	GDK_BLANK_CURSOR CursorType = -2
	// type of cursors constructed with
	// gdk_cursor_new_from_pixbuf()
	GDK_CURSOR_IS_PIXMAP CursorType = -1
)

// Indicates the device type. See [above][GdkDeviceManager.description]
// for more information about the meaning of these device types.
type DeviceType C.GdkDeviceType

const (
	// Device is a master (or virtual) device. There will
	// be an associated focus indicator on the screen.
	GDK_DEVICE_TYPE_MASTER DeviceType = 0
	// Device is a slave (or physical) device.
	GDK_DEVICE_TYPE_SLAVE DeviceType = 1
	// Device is a physical device, currently not attached to
	// any virtual device.
	GDK_DEVICE_TYPE_FLOATING DeviceType = 2
)

// Used in #GdkDragContext to indicate the protocol according to
// which DND is done.
type DragProtocol C.GdkDragProtocol

const (
	// no protocol.
	GDK_DRAG_PROTO_NONE DragProtocol = 0
	// The Motif DND protocol. No longer supported
	GDK_DRAG_PROTO_MOTIF DragProtocol = 1
	// The Xdnd protocol.
	GDK_DRAG_PROTO_XDND DragProtocol = 2
	// An extension to the Xdnd protocol for
	// unclaimed root window drops.
	GDK_DRAG_PROTO_ROOTWIN DragProtocol = 3
	// The simple WM_DROPFILES protocol.
	GDK_DRAG_PROTO_WIN32_DROPFILES DragProtocol = 4
	// The complex OLE2 DND protocol (not implemented).
	GDK_DRAG_PROTO_OLE2 DragProtocol = 5
	// Intra-application DND.
	GDK_DRAG_PROTO_LOCAL DragProtocol = 6
	// Wayland DND protocol.
	GDK_DRAG_PROTO_WAYLAND DragProtocol = 7
)

// Specifies the type of the event.
//
// Do not confuse these events with the signals that GTK+ widgets emit.
// Although many of these events result in corresponding signals being emitted,
// the events are often transformed or filtered along the way.
//
// In some language bindings, the values %GDK_2BUTTON_PRESS and
// %GDK_3BUTTON_PRESS would translate into something syntactically
// invalid (eg `Gdk.EventType.2ButtonPress`, where a
// symbol is not allowed to start with a number). In that case, the
// aliases %GDK_DOUBLE_BUTTON_PRESS and %GDK_TRIPLE_BUTTON_PRESS can
// be used instead.
type EventType C.GdkEventType

const (
	// a special code to indicate a null event.
	GDK_NOTHING EventType = -1
	// the window manager has requested that the toplevel window be
	// hidden or destroyed, usually when the user clicks on a special icon in the
	// title bar.
	GDK_DELETE EventType = 0
	// the window has been destroyed.
	GDK_DESTROY EventType = 1
	// all or part of the window has become visible and needs to be
	// redrawn.
	GDK_EXPOSE EventType = 2
	// the pointer (usually a mouse) has moved.
	GDK_MOTION_NOTIFY EventType = 3
	// a mouse button has been pressed.
	GDK_BUTTON_PRESS EventType = 4
	// a mouse button has been double-clicked (clicked twice
	// within a short period of time). Note that each click also generates a
	// %GDK_BUTTON_PRESS event.
	GDK_2BUTTON_PRESS EventType = 5
	// alias for %GDK_2BUTTON_PRESS, added in 3.6.
	GDK_DOUBLE_BUTTON_PRESS EventType = 5
	// a mouse button has been clicked 3 times in a short period
	// of time. Note that each click also generates a %GDK_BUTTON_PRESS event.
	GDK_3BUTTON_PRESS EventType = 6
	// alias for %GDK_3BUTTON_PRESS, added in 3.6.
	GDK_TRIPLE_BUTTON_PRESS EventType = 6
	// a mouse button has been released.
	GDK_BUTTON_RELEASE EventType = 7
	// a key has been pressed.
	GDK_KEY_PRESS EventType = 8
	// a key has been released.
	GDK_KEY_RELEASE EventType = 9
	// the pointer has entered the window.
	GDK_ENTER_NOTIFY EventType = 10
	// the pointer has left the window.
	GDK_LEAVE_NOTIFY EventType = 11
	// the keyboard focus has entered or left the window.
	GDK_FOCUS_CHANGE EventType = 12
	// the size, position or stacking order of the window has changed.
	// Note that GTK+ discards these events for %GDK_WINDOW_CHILD windows.
	GDK_CONFIGURE EventType = 13
	// the window has been mapped.
	GDK_MAP EventType = 14
	// the window has been unmapped.
	GDK_UNMAP EventType = 15
	// a property on the window has been changed or deleted.
	GDK_PROPERTY_NOTIFY EventType = 16
	// the application has lost ownership of a selection.
	GDK_SELECTION_CLEAR EventType = 17
	// another application has requested a selection.
	GDK_SELECTION_REQUEST EventType = 18
	// a selection has been received.
	GDK_SELECTION_NOTIFY EventType = 19
	// an input device has moved into contact with a sensing
	// surface (e.g. a touchscreen or graphics tablet).
	GDK_PROXIMITY_IN EventType = 20
	// an input device has moved out of contact with a sensing
	// surface.
	GDK_PROXIMITY_OUT EventType = 21
	// the mouse has entered the window while a drag is in progress.
	GDK_DRAG_ENTER EventType = 22
	// the mouse has left the window while a drag is in progress.
	GDK_DRAG_LEAVE EventType = 23
	// the mouse has moved in the window while a drag is in
	// progress.
	GDK_DRAG_MOTION EventType = 24
	// the status of the drag operation initiated by the window
	// has changed.
	GDK_DRAG_STATUS EventType = 25
	// a drop operation onto the window has started.
	GDK_DROP_START EventType = 26
	// the drop operation initiated by the window has completed.
	GDK_DROP_FINISHED EventType = 27
	// a message has been received from another application.
	GDK_CLIENT_EVENT EventType = 28
	// the window visibility status has changed.
	GDK_VISIBILITY_NOTIFY EventType = 29
	// the scroll wheel was turned
	GDK_SCROLL EventType = 31
	// the state of a window has changed. See #GdkWindowState
	// for the possible window states
	GDK_WINDOW_STATE EventType = 32
	// a setting has been modified.
	GDK_SETTING EventType = 33
	// the owner of a selection has changed. This event type
	// was added in 2.6
	GDK_OWNER_CHANGE EventType = 34
	// a pointer or keyboard grab was broken. This event type
	// was added in 2.8.
	GDK_GRAB_BROKEN EventType = 35
	// the content of the window has been changed. This event type
	// was added in 2.14.
	GDK_DAMAGE EventType = 36
	// A new touch event sequence has just started. This event
	// type was added in 3.4.
	GDK_TOUCH_BEGIN EventType = 37
	// A touch event sequence has been updated. This event type
	// was added in 3.4.
	GDK_TOUCH_UPDATE EventType = 38
	// A touch event sequence has finished. This event type
	// was added in 3.4.
	GDK_TOUCH_END EventType = 39
	// A touch event sequence has been canceled. This event type
	// was added in 3.4.
	GDK_TOUCH_CANCEL EventType = 40
	// A touchpad swipe gesture event, the current state
	// is determined by its phase field. This event type was added in 3.18.
	GDK_TOUCHPAD_SWIPE EventType = 41
	// A touchpad pinch gesture event, the current state
	// is determined by its phase field. This event type was added in 3.18.
	GDK_TOUCHPAD_PINCH EventType = 42
	// A tablet pad button press event. This event type
	// was added in 3.22.
	GDK_PAD_BUTTON_PRESS EventType = 43
	// A tablet pad button release event. This event type
	// was added in 3.22.
	GDK_PAD_BUTTON_RELEASE EventType = 44
	// A tablet pad axis event from a "ring". This event type was
	// added in 3.22.
	GDK_PAD_RING EventType = 45
	// A tablet pad axis event from a "strip". This event type was
	// added in 3.22.
	GDK_PAD_STRIP EventType = 46
	// A tablet pad group mode change. This event type was
	// added in 3.22.
	GDK_PAD_GROUP_MODE EventType = 47
	// marks the end of the GdkEventType enumeration. Added in 2.18
	GDK_EVENT_LAST EventType = 48
)

// Specifies the result of applying a #GdkFilterFunc to a native event.
type FilterReturn C.GdkFilterReturn

const (
	// event not handled, continue processing.
	GDK_FILTER_CONTINUE FilterReturn = 0
	// native event translated into a GDK event and stored
	// in the `event` structure that was passed in.
	GDK_FILTER_TRANSLATE FilterReturn = 1
	// event handled, terminate processing.
	GDK_FILTER_REMOVE FilterReturn = 2
)

// Defines how device grabs interact with other devices.
type GrabOwnership C.GdkGrabOwnership

const (
	// All other devices’ events are allowed.
	GDK_OWNERSHIP_NONE GrabOwnership = 0
	// Other devices’ events are blocked for the grab window.
	GDK_OWNERSHIP_WINDOW GrabOwnership = 1
	// Other devices’ events are blocked for the whole application.
	GDK_OWNERSHIP_APPLICATION GrabOwnership = 2
)

// Returned by gdk_device_grab(), gdk_pointer_grab() and gdk_keyboard_grab() to
// indicate success or the reason for the failure of the grab attempt.
type GrabStatus C.GdkGrabStatus

const (
	// the resource was successfully grabbed.
	GDK_GRAB_SUCCESS GrabStatus = 0
	// the resource is actively grabbed by another client.
	GDK_GRAB_ALREADY_GRABBED GrabStatus = 1
	// the resource was grabbed more recently than the
	// specified time.
	GDK_GRAB_INVALID_TIME GrabStatus = 2
	// the grab window or the @confine_to window are not
	// viewable.
	GDK_GRAB_NOT_VIEWABLE GrabStatus = 3
	// the resource is frozen by an active grab of another client.
	GDK_GRAB_FROZEN GrabStatus = 4
	// the grab failed for some other reason. Since 3.16
	GDK_GRAB_FAILED GrabStatus = 5
)

// Defines the reference point of a window and the meaning of coordinates
// passed to gtk_window_move(). See gtk_window_move() and the "implementation
// notes" section of the
// [Extended Window Manager Hints](http://www.freedesktop.org/Standards/wm-spec)
// specification for more details.
type Gravity C.GdkGravity

const (
	// the reference point is at the top left corner.
	GDK_GRAVITY_NORTH_WEST Gravity = 1
	// the reference point is in the middle of the top edge.
	GDK_GRAVITY_NORTH Gravity = 2
	// the reference point is at the top right corner.
	GDK_GRAVITY_NORTH_EAST Gravity = 3
	// the reference point is at the middle of the left edge.
	GDK_GRAVITY_WEST Gravity = 4
	// the reference point is at the center of the window.
	GDK_GRAVITY_CENTER Gravity = 5
	// the reference point is at the middle of the right edge.
	GDK_GRAVITY_EAST Gravity = 6
	// the reference point is at the lower left corner.
	GDK_GRAVITY_SOUTH_WEST Gravity = 7
	// the reference point is at the middle of the lower edge.
	GDK_GRAVITY_SOUTH Gravity = 8
	// the reference point is at the lower right corner.
	GDK_GRAVITY_SOUTH_EAST Gravity = 9
	// the reference point is at the top left corner of the
	// window itself, ignoring window manager decorations.
	GDK_GRAVITY_STATIC Gravity = 10
)

// An enumeration that describes the mode of an input device.
type InputMode C.GdkInputMode

const (
	// the device is disabled and will not report any events.
	GDK_MODE_DISABLED InputMode = 0
	// the device is enabled. The device’s coordinate space
	// maps to the entire screen.
	GDK_MODE_SCREEN InputMode = 1
	// the device is enabled. The device’s coordinate space
	// is mapped to a single window. The manner in which this window
	// is chosen is undefined, but it will typically be the same
	// way in which the focus window for key events is determined.
	GDK_MODE_WINDOW InputMode = 2
)

// An enumeration describing the type of an input device in general terms.
type InputSource C.GdkInputSource

const (
	// the device is a mouse. (This will be reported for the core
	// pointer, even if it is something else, such as a trackball.)
	GDK_SOURCE_MOUSE InputSource = 0
	// the device is a stylus of a graphics tablet or similar device.
	GDK_SOURCE_PEN InputSource = 1
	// the device is an eraser. Typically, this would be the other end
	// of a stylus on a graphics tablet.
	GDK_SOURCE_ERASER InputSource = 2
	// the device is a graphics tablet “puck” or similar device.
	GDK_SOURCE_CURSOR InputSource = 3
	// the device is a keyboard.
	GDK_SOURCE_KEYBOARD InputSource = 4
	// the device is a direct-input touch device, such
	// as a touchscreen or tablet. This device type has been added in 3.4.
	GDK_SOURCE_TOUCHSCREEN InputSource = 5
	// the device is an indirect touch device, such
	// as a touchpad. This device type has been added in 3.4.
	GDK_SOURCE_TOUCHPAD InputSource = 6
	// the device is a trackpoint. This device type has been
	// added in 3.22
	GDK_SOURCE_TRACKPOINT InputSource = 7
	// the device is a "pad", a collection of buttons,
	// rings and strips found in drawing tablets. This device type has been
	// added in 3.22.
	GDK_SOURCE_TABLET_PAD InputSource = 8
)

// Specifies the kind of crossing for #GdkEventCrossing.
//
// See the X11 protocol specification of LeaveNotify for
// full details of crossing event generation.
type NotifyType C.GdkNotifyType

const (
	// the window is entered from an ancestor or
	// left towards an ancestor.
	GDK_NOTIFY_ANCESTOR NotifyType = 0
	// the pointer moves between an ancestor and an
	// inferior of the window.
	GDK_NOTIFY_VIRTUAL NotifyType = 1
	// the window is entered from an inferior or
	// left towards an inferior.
	GDK_NOTIFY_INFERIOR NotifyType = 2
	// the window is entered from or left towards
	// a window which is neither an ancestor nor an inferior.
	GDK_NOTIFY_NONLINEAR NotifyType = 3
	// the pointer moves between two windows
	// which are not ancestors of each other and the window is part of
	// the ancestor chain between one of these windows and their least
	// common ancestor.
	GDK_NOTIFY_NONLINEAR_VIRTUAL NotifyType = 4
	// an unknown type of enter/leave event occurred.
	GDK_NOTIFY_UNKNOWN NotifyType = 5
)

// Specifies why a selection ownership was changed.
type OwnerChange C.GdkOwnerChange

const (
	// some other app claimed the ownership
	GDK_OWNER_CHANGE_NEW_OWNER OwnerChange = 0
	// the window was destroyed
	GDK_OWNER_CHANGE_DESTROY OwnerChange = 1
	// the client was closed
	GDK_OWNER_CHANGE_CLOSE OwnerChange = 2
)

// Describes how existing data is combined with new data when
// using gdk_property_change().
type PropMode C.GdkPropMode

const (
	// the new data replaces the existing data.
	GDK_PROP_MODE_REPLACE PropMode = 0
	// the new data is prepended to the existing data.
	GDK_PROP_MODE_PREPEND PropMode = 1
	// the new data is appended to the existing data.
	GDK_PROP_MODE_APPEND PropMode = 2
)

// Specifies the type of a property change for a #GdkEventProperty.
type PropertyState C.guint

const (
	// the property value was changed.
	GDK_PROPERTY_NEW_VALUE PropertyState = 0
	// the property was deleted.
	GDK_PROPERTY_DELETE PropertyState = 1
)

// Specifies the direction for #GdkEventScroll.
type ScrollDirection C.GdkScrollDirection

const (
	// the window is scrolled up.
	GDK_SCROLL_UP ScrollDirection = 0
	// the window is scrolled down.
	GDK_SCROLL_DOWN ScrollDirection = 1
	// the window is scrolled to the left.
	GDK_SCROLL_LEFT ScrollDirection = 2
	// the window is scrolled to the right.
	GDK_SCROLL_RIGHT ScrollDirection = 3
	// the scrolling is determined by the delta values
	// in #GdkEventScroll. See gdk_event_get_scroll_deltas(). Since: 3.4
	GDK_SCROLL_SMOOTH ScrollDirection = 4
)

// Specifies the kind of modification applied to a setting in a
// #GdkEventSetting.
type SettingAction C.GdkSettingAction

const (
	// a setting was added.
	GDK_SETTING_ACTION_NEW SettingAction = 0
	// a setting was changed.
	GDK_SETTING_ACTION_CHANGED SettingAction = 1
	// a setting was deleted.
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

// Specifies the current state of a touchpad gesture. All gestures are
// guaranteed to begin with an event with phase %GDK_TOUCHPAD_GESTURE_PHASE_BEGIN,
// followed by 0 or several events with phase %GDK_TOUCHPAD_GESTURE_PHASE_UPDATE.
//
// A finished gesture may have 2 possible outcomes, an event with phase
// %GDK_TOUCHPAD_GESTURE_PHASE_END will be emitted when the gesture is
// considered successful, this should be used as the hint to perform any
// permanent changes.
//
// Cancelled gestures may be so for a variety of reasons, due to hardware
// or the compositor, or due to the gesture recognition layers hinting the
// gesture did not finish resolutely (eg. a 3rd finger being added during
// a pinch gesture). In these cases, the last event will report the phase
// %GDK_TOUCHPAD_GESTURE_PHASE_CANCEL, this should be used as a hint
// to undo any visible/permanent changes that were done throughout the
// progress of the gesture.
//
// See also #GdkEventTouchpadSwipe and #GdkEventTouchpadPinch.
type TouchpadGesturePhase C.GdkTouchpadGesturePhase

const (
	// The gesture has begun.
	GDK_TOUCHPAD_GESTURE_PHASE_BEGIN TouchpadGesturePhase = 0
	// The gesture has been updated.
	GDK_TOUCHPAD_GESTURE_PHASE_UPDATE TouchpadGesturePhase = 1
	// The gesture was finished, changes
	// should be permanently applied.
	GDK_TOUCHPAD_GESTURE_PHASE_END TouchpadGesturePhase = 2
	// The gesture was cancelled, all
	// changes should be undone.
	GDK_TOUCHPAD_GESTURE_PHASE_CANCEL TouchpadGesturePhase = 3
)

// Specifies the visiblity status of a window for a #GdkEventVisibility.
type VisibilityState C.GdkVisibilityState

const (
	// the window is completely visible.
	GDK_VISIBILITY_UNOBSCURED VisibilityState = 0
	// the window is partially visible.
	GDK_VISIBILITY_PARTIAL VisibilityState = 1
	// the window is not visible at all.
	GDK_VISIBILITY_FULLY_OBSCURED VisibilityState = 2
)

// A set of values that describe the manner in which the pixel values
// for a visual are converted into RGB values for display.
type VisualType C.GdkVisualType

const (
	// Each pixel value indexes a grayscale value
	// directly.
	GDK_VISUAL_STATIC_GRAY VisualType = 0
	// Each pixel is an index into a color map that
	// maps pixel values into grayscale values. The color map can be
	// changed by an application.
	GDK_VISUAL_GRAYSCALE VisualType = 1
	// Each pixel value is an index into a predefined,
	// unmodifiable color map that maps pixel values into RGB values.
	GDK_VISUAL_STATIC_COLOR VisualType = 2
	// Each pixel is an index into a color map that
	// maps pixel values into rgb values. The color map can be changed by
	// an application.
	GDK_VISUAL_PSEUDO_COLOR VisualType = 3
	// Each pixel value directly contains red, green,
	// and blue components. Use gdk_visual_get_red_pixel_details(), etc,
	// to obtain information about how the components are assembled into
	// a pixel value.
	GDK_VISUAL_TRUE_COLOR VisualType = 4
	// Each pixel value contains red, green, and blue
	// components as for %GDK_VISUAL_TRUE_COLOR, but the components are
	// mapped via a color table into the final output table instead of
	// being converted directly.
	GDK_VISUAL_DIRECT_COLOR VisualType = 5
)

// Determines a window edge or corner.
type WindowEdge C.GdkWindowEdge

const (
	// the top left corner.
	GDK_WINDOW_EDGE_NORTH_WEST WindowEdge = 0
	// the top edge.
	GDK_WINDOW_EDGE_NORTH WindowEdge = 1
	// the top right corner.
	GDK_WINDOW_EDGE_NORTH_EAST WindowEdge = 2
	// the left edge.
	GDK_WINDOW_EDGE_WEST WindowEdge = 3
	// the right edge.
	GDK_WINDOW_EDGE_EAST WindowEdge = 4
	// the lower left corner.
	GDK_WINDOW_EDGE_SOUTH_WEST WindowEdge = 5
	// the lower edge.
	GDK_WINDOW_EDGE_SOUTH WindowEdge = 6
	// the lower right corner.
	GDK_WINDOW_EDGE_SOUTH_EAST WindowEdge = 7
)

// Describes the kind of window.
type WindowType C.GdkWindowType

const (
	// root window; this window has no parent, covers the entire
	// screen, and is created by the window system
	GDK_WINDOW_ROOT WindowType = 0
	// toplevel window (used to implement #GtkWindow)
	GDK_WINDOW_TOPLEVEL WindowType = 1
	// child window (used to implement e.g. #GtkEntry)
	GDK_WINDOW_CHILD WindowType = 2
	// override redirect temporary window (used to implement
	// #GtkMenu)
	GDK_WINDOW_TEMP WindowType = 3
	// foreign window (see gdk_window_foreign_new())
	GDK_WINDOW_FOREIGN WindowType = 4
	// offscreen window (see
	// [Offscreen Windows][OFFSCREEN-WINDOWS]). Since 2.18
	GDK_WINDOW_OFFSCREEN WindowType = 5
	// subsurface-based window; This window is visually
	// tied to a toplevel, and is moved/stacked with it. Currently this window
	// type is only implemented in Wayland. Since 3.14
	GDK_WINDOW_SUBSURFACE WindowType = 6
)

// These are hints for the window manager that indicate what type of function
// the window has. The window manager can use this when determining decoration
// and behaviour of the window. The hint must be set before mapping the window.
//
// See the [Extended Window Manager Hints](http://www.freedesktop.org/Standards/wm-spec)
// specification for more details about window types.
type WindowTypeHint C.GdkWindowTypeHint

const (
	// Normal toplevel window.
	GDK_WINDOW_TYPE_HINT_NORMAL WindowTypeHint = 0
	// Dialog window.
	GDK_WINDOW_TYPE_HINT_DIALOG WindowTypeHint = 1
	// Window used to implement a menu; GTK+ uses
	// this hint only for torn-off menus, see #GtkTearoffMenuItem.
	GDK_WINDOW_TYPE_HINT_MENU WindowTypeHint = 2
	// Window used to implement toolbars.
	GDK_WINDOW_TYPE_HINT_TOOLBAR WindowTypeHint = 3
	// Window used to display a splash
	// screen during application startup.
	GDK_WINDOW_TYPE_HINT_SPLASHSCREEN WindowTypeHint = 4
	// Utility windows which are not detached
	// toolbars or dialogs.
	GDK_WINDOW_TYPE_HINT_UTILITY WindowTypeHint = 5
	// Used for creating dock or panel windows.
	GDK_WINDOW_TYPE_HINT_DOCK WindowTypeHint = 6
	// Used for creating the desktop background
	// window.
	GDK_WINDOW_TYPE_HINT_DESKTOP WindowTypeHint = 7
	// A menu that belongs to a menubar.
	GDK_WINDOW_TYPE_HINT_DROPDOWN_MENU WindowTypeHint = 8
	// A menu that does not belong to a menubar,
	// e.g. a context menu.
	GDK_WINDOW_TYPE_HINT_POPUP_MENU WindowTypeHint = 9
	// A tooltip.
	GDK_WINDOW_TYPE_HINT_TOOLTIP WindowTypeHint = 10
	// A notification - typically a “bubble”
	// that belongs to a status icon.
	GDK_WINDOW_TYPE_HINT_NOTIFICATION WindowTypeHint = 11
	// A popup from a combo box.
	GDK_WINDOW_TYPE_HINT_COMBO WindowTypeHint = 12
	// A window that is used to implement a DND cursor.
	GDK_WINDOW_TYPE_HINT_DND WindowTypeHint = 13
)

// @GDK_INPUT_OUTPUT windows are the standard kind of window you might expect.
// Such windows receive events and are also displayed on screen.
// @GDK_INPUT_ONLY windows are invisible; they are usually placed above other
// windows in order to trap or filter the events. You can’t draw on
// @GDK_INPUT_ONLY windows.
type WindowWindowClass C.GdkWindowWindowClass

const (
	// window for graphics and events
	GDK_INPUT_OUTPUT WindowWindowClass = 0
	// window for events only
	GDK_INPUT_ONLY WindowWindowClass = 1
)
