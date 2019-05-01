// This is a generated file - DO NOT EDIT
// +build gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29
// +build gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29
// +build gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29
// +build gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29
// +build gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29
// +build gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29
// +build gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29
// +build gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29
// +build gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

// Unsupported : gtk_gesture_drag_new : return type :

// Unsupported : gtk_gesture_long_press_new : return type :

// Unsupported : gtk_gesture_multi_press_new : return type :

// Unsupported : gtk_gesture_pan_new : return type :

// Unsupported : gtk_gesture_rotate_new : return type :

// Unsupported : gtk_gesture_swipe_new : return type :

// Unsupported : gtk_gesture_zoom_new : return type :

const STYLE_CLASS_CSD string = "csd"
const STYLE_CLASS_FLAT string = "flat"
const STYLE_CLASS_MESSAGE_DIALOG string = "message-dialog"
const STYLE_CLASS_OVERSHOOT string = "overshoot"
const STYLE_CLASS_POPOVER string = "popover"
const STYLE_CLASS_POPUP string = "popup"
const STYLE_CLASS_SUBTITLE string = "subtitle"
const STYLE_CLASS_TITLE string = "title"

type EventSequenceState int

const (
	GTK_EVENT_SEQUENCE_NONE    EventSequenceState = 0
	GTK_EVENT_SEQUENCE_CLAIMED EventSequenceState = 1
	GTK_EVENT_SEQUENCE_DENIED  EventSequenceState = 2
)

type PanDirection int

const (
	GTK_PAN_DIRECTION_LEFT  PanDirection = 0
	GTK_PAN_DIRECTION_RIGHT PanDirection = 1
	GTK_PAN_DIRECTION_UP    PanDirection = 2
	GTK_PAN_DIRECTION_DOWN  PanDirection = 3
)

type PropagationPhase int

const (
	GTK_PHASE_NONE    PropagationPhase = 0
	GTK_PHASE_CAPTURE PropagationPhase = 1
	GTK_PHASE_BUBBLE  PropagationPhase = 2
	GTK_PHASE_TARGET  PropagationPhase = 3
)
