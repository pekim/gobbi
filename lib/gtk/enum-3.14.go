// This is a generated file - DO NOT EDIT
// +build gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

type EventSequenceState C.GtkEventSequenceState

const (
	GTK_EVENT_SEQUENCE_NONE    EventSequenceState = 0
	GTK_EVENT_SEQUENCE_CLAIMED EventSequenceState = 1
	GTK_EVENT_SEQUENCE_DENIED  EventSequenceState = 2
)

type PanDirection C.GtkPanDirection

const (
	GTK_PAN_DIRECTION_LEFT  PanDirection = 0
	GTK_PAN_DIRECTION_RIGHT PanDirection = 1
	GTK_PAN_DIRECTION_UP    PanDirection = 2
	GTK_PAN_DIRECTION_DOWN  PanDirection = 3
)

type PropagationPhase C.GtkPropagationPhase

const (
	GTK_PHASE_NONE    PropagationPhase = 0
	GTK_PHASE_CAPTURE PropagationPhase = 1
	GTK_PHASE_BUBBLE  PropagationPhase = 2
	GTK_PHASE_TARGET  PropagationPhase = 3
)
