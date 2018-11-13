// This is a generated file - DO NOT EDIT
// +build gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

// Describes the state of a #GdkEventSequence in a #GtkGesture.
type EventSequenceState C.GtkEventSequenceState

const (
	// The sequence is handled, but not grabbed.
	GTK_EVENT_SEQUENCE_NONE EventSequenceState = 0

	// The sequence is handled and grabbed.
	GTK_EVENT_SEQUENCE_CLAIMED EventSequenceState = 1

	// The sequence is denied.
	GTK_EVENT_SEQUENCE_DENIED EventSequenceState = 2
)

// Describes the panning direction of a #GtkGesturePan
type PanDirection C.GtkPanDirection

const (
	// panned towards the left
	GTK_PAN_DIRECTION_LEFT PanDirection = 0

	// panned towards the right
	GTK_PAN_DIRECTION_RIGHT PanDirection = 1

	// panned upwards
	GTK_PAN_DIRECTION_UP PanDirection = 2

	// panned downwards
	GTK_PAN_DIRECTION_DOWN PanDirection = 3
)

// Describes the stage at which events are fed into a #GtkEventController.
type PropagationPhase C.GtkPropagationPhase

const (
	// Events are not delivered automatically. Those can be
	// manually fed through gtk_event_controller_handle_event(). This should
	// only be used when full control about when, or whether the controller
	// handles the event is needed.
	GTK_PHASE_NONE PropagationPhase = 0

	// Events are delivered in the capture phase. The
	// capture phase happens before the bubble phase, runs from the toplevel down
	// to the event widget. This option should only be used on containers that
	// might possibly handle events before their children do.
	GTK_PHASE_CAPTURE PropagationPhase = 1

	// Events are delivered in the bubble phase. The bubble
	// phase happens after the capture phase, and before the default handlers
	// are run. This phase runs from the event widget, up to the toplevel.
	GTK_PHASE_BUBBLE PropagationPhase = 2

	// Events are delivered in the default widget event handlers,
	// note that widget implementations must chain up on button, motion, touch and
	// grab broken handlers for controllers in this phase to be run.
	GTK_PHASE_TARGET PropagationPhase = 3
)
