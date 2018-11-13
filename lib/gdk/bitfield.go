// This is a generated file - DO NOT EDIT

package gdk

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <stdlib.h>
import "C"

// Used in #GdkDragContext to indicate what the destination
// should do with the dropped data.
type DragAction C.GdkDragAction

const (
	// Means nothing, and should not be used.
	GDK_ACTION_DEFAULT DragAction = 1

	// Copy the data.
	GDK_ACTION_COPY DragAction = 2

	// Move the data, i.e. first copy it, then delete
	// it from the source using the DELETE target of the X selection protocol.
	GDK_ACTION_MOVE DragAction = 4

	// Add a link to the data. Note that this is only
	// useful if source and destination agree on what it means.
	GDK_ACTION_LINK DragAction = 8

	// Special action which tells the source that the
	// destination will do something that the source doesn’t understand.
	GDK_ACTION_PRIVATE DragAction = 16

	// Ask the user what to do with the data.
	GDK_ACTION_ASK DragAction = 32
)

// A set of bit-flags to indicate which events a window is to receive.
// Most of these masks map onto one or more of the #GdkEventType event types
// above.
//
// See the [input handling overview][chap-input-handling] for details of
// [event masks][event-masks] and [event propagation][event-propagation].
//
// %GDK_POINTER_MOTION_HINT_MASK is deprecated. It is a special mask
// to reduce the number of %GDK_MOTION_NOTIFY events received. When using
// %GDK_POINTER_MOTION_HINT_MASK, fewer %GDK_MOTION_NOTIFY events will
// be sent, some of which are marked as a hint (the is_hint member is
// %TRUE). To receive more motion events after a motion hint event,
// the application needs to asks for more, by calling
// gdk_event_request_motions().
//
// Since GTK 3.8, motion events are already compressed by default, independent
// of this mechanism. This compression can be disabled with
// gdk_window_set_event_compression(). See the documentation of that function
// for details.
//
// If %GDK_TOUCH_MASK is enabled, the window will receive touch events
// from touch-enabled devices. Those will come as sequences of #GdkEventTouch
// with type %GDK_TOUCH_UPDATE, enclosed by two events with
// type %GDK_TOUCH_BEGIN and %GDK_TOUCH_END (or %GDK_TOUCH_CANCEL).
// gdk_event_get_event_sequence() returns the event sequence for these
// events, so different sequences may be distinguished.
type EventMask C.GdkEventMask

const (
	// receive expose events
	GDK_EXPOSURE_MASK EventMask = 2

	// receive all pointer motion events
	GDK_POINTER_MOTION_MASK EventMask = 4

	// deprecated. see the explanation above
	GDK_POINTER_MOTION_HINT_MASK EventMask = 8

	// receive pointer motion events while any button is pressed
	GDK_BUTTON_MOTION_MASK EventMask = 16

	// receive pointer motion events while 1 button is pressed
	GDK_BUTTON1_MOTION_MASK EventMask = 32

	// receive pointer motion events while 2 button is pressed
	GDK_BUTTON2_MOTION_MASK EventMask = 64

	// receive pointer motion events while 3 button is pressed
	GDK_BUTTON3_MOTION_MASK EventMask = 128

	// receive button press events
	GDK_BUTTON_PRESS_MASK EventMask = 256

	// receive button release events
	GDK_BUTTON_RELEASE_MASK EventMask = 512

	// receive key press events
	GDK_KEY_PRESS_MASK EventMask = 1024

	// receive key release events
	GDK_KEY_RELEASE_MASK EventMask = 2048

	// receive window enter events
	GDK_ENTER_NOTIFY_MASK EventMask = 4096

	// receive window leave events
	GDK_LEAVE_NOTIFY_MASK EventMask = 8192

	// receive focus change events
	GDK_FOCUS_CHANGE_MASK EventMask = 16384

	// receive events about window configuration change
	GDK_STRUCTURE_MASK EventMask = 32768

	// receive property change events
	GDK_PROPERTY_CHANGE_MASK EventMask = 65536

	// receive visibility change events
	GDK_VISIBILITY_NOTIFY_MASK EventMask = 131072

	// receive proximity in events
	GDK_PROXIMITY_IN_MASK EventMask = 262144

	// receive proximity out events
	GDK_PROXIMITY_OUT_MASK EventMask = 524288

	// receive events about window configuration changes of
	// child windows
	GDK_SUBSTRUCTURE_MASK EventMask = 1048576

	// receive scroll events
	GDK_SCROLL_MASK EventMask = 2097152

	// receive touch events. Since 3.4
	GDK_TOUCH_MASK EventMask = 4194304

	// receive smooth scrolling events. Since 3.4
	GDK_SMOOTH_SCROLL_MASK EventMask = 8388608

	// receive touchpad gesture events. Since 3.18
	GDK_TOUCHPAD_GESTURE_MASK EventMask = 16777216

	// receive tablet pad events. Since 3.22
	GDK_TABLET_PAD_MASK EventMask = 33554432

	// the combination of all the above event masks.
	GDK_ALL_EVENTS_MASK EventMask = 67108862
)

// A set of bit-flags to indicate the state of modifier keys and mouse buttons
// in various event types. Typical modifier keys are Shift, Control, Meta,
// Super, Hyper, Alt, Compose, Apple, CapsLock or ShiftLock.
//
// Like the X Window System, GDK supports 8 modifier keys and 5 mouse buttons.
//
// Since 2.10, GDK recognizes which of the Meta, Super or Hyper keys are mapped
// to Mod2 - Mod5, and indicates this by setting %GDK_SUPER_MASK,
// %GDK_HYPER_MASK or %GDK_META_MASK in the state field of key events.
//
// Note that GDK may add internal values to events which include
// reserved values such as %GDK_MODIFIER_RESERVED_13_MASK.  Your code
// should preserve and ignore them.  You can use %GDK_MODIFIER_MASK to
// remove all reserved values.
//
// Also note that the GDK X backend interprets button press events for button
// 4-7 as scroll events, so %GDK_BUTTON4_MASK and %GDK_BUTTON5_MASK will never
// be set.
type ModifierType C.guint

const (
	// the Shift key.
	GDK_SHIFT_MASK ModifierType = 1

	// a Lock key (depending on the modifier mapping of the
	// X server this may either be CapsLock or ShiftLock).
	GDK_LOCK_MASK ModifierType = 2

	// the Control key.
	GDK_CONTROL_MASK ModifierType = 4

	// the fourth modifier key (it depends on the modifier
	// mapping of the X server which key is interpreted as this modifier, but
	// normally it is the Alt key).
	GDK_MOD1_MASK ModifierType = 8

	// the fifth modifier key (it depends on the modifier
	// mapping of the X server which key is interpreted as this modifier).
	GDK_MOD2_MASK ModifierType = 16

	// the sixth modifier key (it depends on the modifier
	// mapping of the X server which key is interpreted as this modifier).
	GDK_MOD3_MASK ModifierType = 32

	// the seventh modifier key (it depends on the modifier
	// mapping of the X server which key is interpreted as this modifier).
	GDK_MOD4_MASK ModifierType = 64

	// the eighth modifier key (it depends on the modifier
	// mapping of the X server which key is interpreted as this modifier).
	GDK_MOD5_MASK ModifierType = 128

	// the first mouse button.
	GDK_BUTTON1_MASK ModifierType = 256

	// the second mouse button.
	GDK_BUTTON2_MASK ModifierType = 512

	// the third mouse button.
	GDK_BUTTON3_MASK ModifierType = 1024

	// the fourth mouse button.
	GDK_BUTTON4_MASK ModifierType = 2048

	// the fifth mouse button.
	GDK_BUTTON5_MASK ModifierType = 4096

	// A reserved bit flag; do not use in your own code
	GDK_MODIFIER_RESERVED_13_MASK ModifierType = 8192

	// A reserved bit flag; do not use in your own code
	GDK_MODIFIER_RESERVED_14_MASK ModifierType = 16384

	// A reserved bit flag; do not use in your own code
	GDK_MODIFIER_RESERVED_15_MASK ModifierType = 32768

	// A reserved bit flag; do not use in your own code
	GDK_MODIFIER_RESERVED_16_MASK ModifierType = 65536

	// A reserved bit flag; do not use in your own code
	GDK_MODIFIER_RESERVED_17_MASK ModifierType = 131072

	// A reserved bit flag; do not use in your own code
	GDK_MODIFIER_RESERVED_18_MASK ModifierType = 262144

	// A reserved bit flag; do not use in your own code
	GDK_MODIFIER_RESERVED_19_MASK ModifierType = 524288

	// A reserved bit flag; do not use in your own code
	GDK_MODIFIER_RESERVED_20_MASK ModifierType = 1048576

	// A reserved bit flag; do not use in your own code
	GDK_MODIFIER_RESERVED_21_MASK ModifierType = 2097152

	// A reserved bit flag; do not use in your own code
	GDK_MODIFIER_RESERVED_22_MASK ModifierType = 4194304

	// A reserved bit flag; do not use in your own code
	GDK_MODIFIER_RESERVED_23_MASK ModifierType = 8388608

	// A reserved bit flag; do not use in your own code
	GDK_MODIFIER_RESERVED_24_MASK ModifierType = 16777216

	// A reserved bit flag; do not use in your own code
	GDK_MODIFIER_RESERVED_25_MASK ModifierType = 33554432

	// the Super modifier. Since 2.10
	GDK_SUPER_MASK ModifierType = 67108864

	// the Hyper modifier. Since 2.10
	GDK_HYPER_MASK ModifierType = 134217728

	// the Meta modifier. Since 2.10
	GDK_META_MASK ModifierType = 268435456

	// A reserved bit flag; do not use in your own code
	GDK_MODIFIER_RESERVED_29_MASK ModifierType = 536870912

	// not used in GDK itself. GTK+ uses it to differentiate
	// between (keyval, modifiers) pairs from key press and release events.
	GDK_RELEASE_MASK ModifierType = 1073741824

	// a mask covering all modifier types.
	GDK_MODIFIER_MASK ModifierType = 1543512063
)

// These are hints originally defined by the Motif toolkit.
// The window manager can use them when determining how to decorate
// the window. The hint must be set before mapping the window.
type WMDecoration C.GdkWMDecoration

const (
	// all decorations should be applied.
	GDK_DECOR_ALL WMDecoration = 1

	// a frame should be drawn around the window.
	GDK_DECOR_BORDER WMDecoration = 2

	// the frame should have resize handles.
	GDK_DECOR_RESIZEH WMDecoration = 4

	// a titlebar should be placed above the window.
	GDK_DECOR_TITLE WMDecoration = 8

	// a button for opening a menu should be included.
	GDK_DECOR_MENU WMDecoration = 16

	// a minimize button should be included.
	GDK_DECOR_MINIMIZE WMDecoration = 32

	// a maximize button should be included.
	GDK_DECOR_MAXIMIZE WMDecoration = 64
)

// These are hints originally defined by the Motif toolkit. The window manager
// can use them when determining the functions to offer for the window. The
// hint must be set before mapping the window.
type WMFunction C.GdkWMFunction

const (
	// all functions should be offered.
	GDK_FUNC_ALL WMFunction = 1

	// the window should be resizable.
	GDK_FUNC_RESIZE WMFunction = 2

	// the window should be movable.
	GDK_FUNC_MOVE WMFunction = 4

	// the window should be minimizable.
	GDK_FUNC_MINIMIZE WMFunction = 8

	// the window should be maximizable.
	GDK_FUNC_MAXIMIZE WMFunction = 16

	// the window should be closable.
	GDK_FUNC_CLOSE WMFunction = 32
)

// Used to indicate which fields in the #GdkWindowAttr struct should be honored.
// For example, if you filled in the “cursor” and “x” fields of #GdkWindowAttr,
// pass “@GDK_WA_X | @GDK_WA_CURSOR” to gdk_window_new(). Fields in
// #GdkWindowAttr not covered by a bit in this enum are required; for example,
// the @width/@height, @wclass, and @window_type fields are required, they have
// no corresponding flag in #GdkWindowAttributesType.
type WindowAttributesType C.GdkWindowAttributesType

const (
	// Honor the title field
	GDK_WA_TITLE WindowAttributesType = 2

	// Honor the X coordinate field
	GDK_WA_X WindowAttributesType = 4

	// Honor the Y coordinate field
	GDK_WA_Y WindowAttributesType = 8

	// Honor the cursor field
	GDK_WA_CURSOR WindowAttributesType = 16

	// Honor the visual field
	GDK_WA_VISUAL WindowAttributesType = 32

	// Honor the wmclass_class and wmclass_name fields
	GDK_WA_WMCLASS WindowAttributesType = 64

	// Honor the override_redirect field
	GDK_WA_NOREDIR WindowAttributesType = 128

	// Honor the type_hint field
	GDK_WA_TYPE_HINT WindowAttributesType = 256
)

// Used to indicate which fields of a #GdkGeometry struct should be paid
// attention to. Also, the presence/absence of @GDK_HINT_POS,
// @GDK_HINT_USER_POS, and @GDK_HINT_USER_SIZE is significant, though they don't
// directly refer to #GdkGeometry fields. @GDK_HINT_USER_POS will be set
// automatically by #GtkWindow if you call gtk_window_move().
// @GDK_HINT_USER_POS and @GDK_HINT_USER_SIZE should be set if the user
// specified a size/position using a --geometry command-line argument;
// gtk_window_parse_geometry() automatically sets these flags.
type WindowHints C.GdkWindowHints

const (
	// indicates that the program has positioned the window
	GDK_HINT_POS WindowHints = 1

	// min size fields are set
	GDK_HINT_MIN_SIZE WindowHints = 2

	// max size fields are set
	GDK_HINT_MAX_SIZE WindowHints = 4

	// base size fields are set
	GDK_HINT_BASE_SIZE WindowHints = 8

	// aspect ratio fields are set
	GDK_HINT_ASPECT WindowHints = 16

	// resize increment fields are set
	GDK_HINT_RESIZE_INC WindowHints = 32

	// window gravity field is set
	GDK_HINT_WIN_GRAVITY WindowHints = 64

	// indicates that the window’s position was explicitly set
	// by the user
	GDK_HINT_USER_POS WindowHints = 128

	// indicates that the window’s size was explicitly set by
	// the user
	GDK_HINT_USER_SIZE WindowHints = 256
)

// Specifies the state of a toplevel window.
type WindowState C.GdkWindowState

const (
	// the window is not shown.
	GDK_WINDOW_STATE_WITHDRAWN WindowState = 1

	// the window is minimized.
	GDK_WINDOW_STATE_ICONIFIED WindowState = 2

	// the window is maximized.
	GDK_WINDOW_STATE_MAXIMIZED WindowState = 4

	// the window is sticky.
	GDK_WINDOW_STATE_STICKY WindowState = 8

	// the window is maximized without
	// decorations.
	GDK_WINDOW_STATE_FULLSCREEN WindowState = 16

	// the window is kept above other windows.
	GDK_WINDOW_STATE_ABOVE WindowState = 32

	// the window is kept below other windows.
	GDK_WINDOW_STATE_BELOW WindowState = 64

	// the window is presented as focused (with active decorations).
	GDK_WINDOW_STATE_FOCUSED WindowState = 128

	// the window is in a tiled state, Since 3.10. Since 3.22.23, this
	// is deprecated in favor of per-edge information.
	GDK_WINDOW_STATE_TILED WindowState = 256

	// whether the top edge is tiled, Since 3.22.23
	GDK_WINDOW_STATE_TOP_TILED WindowState = 512

	// whether the top edge is resizable, Since 3.22.23
	GDK_WINDOW_STATE_TOP_RESIZABLE WindowState = 1024

	// whether the right edge is tiled, Since 3.22.23
	GDK_WINDOW_STATE_RIGHT_TILED WindowState = 2048

	// whether the right edge is resizable, Since 3.22.23
	GDK_WINDOW_STATE_RIGHT_RESIZABLE WindowState = 4096

	// whether the bottom edge is tiled, Since 3.22.23
	GDK_WINDOW_STATE_BOTTOM_TILED WindowState = 8192

	// whether the bottom edge is resizable, Since 3.22.23
	GDK_WINDOW_STATE_BOTTOM_RESIZABLE WindowState = 16384

	// whether the left edge is tiled, Since 3.22.23
	GDK_WINDOW_STATE_LEFT_TILED WindowState = 32768

	// whether the left edge is resizable, Since 3.22.23
	GDK_WINDOW_STATE_LEFT_RESIZABLE WindowState = 65536
)
