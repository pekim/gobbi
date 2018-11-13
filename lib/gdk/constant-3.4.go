// This is a generated file - DO NOT EDIT
// +build gdk_3.4 gdk_3.8 gdk_3.10 gdk_3.12 gdk_3.14 gdk_3.16 gdk_3.18 gdk_3.20 gdk_3.22

package gdk

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <stdlib.h>
import "C"

// The middle button.
const BUTTON_MIDDLE int = C.GDK_BUTTON_MIDDLE

// The primary button. This is typically the left mouse button, or the
// right button in a left-handed setup.
const BUTTON_PRIMARY int = C.GDK_BUTTON_PRIMARY

// The secondary button. This is typically the right mouse button, or the
// left button in a left-handed setup.
const BUTTON_SECONDARY int = C.GDK_BUTTON_SECONDARY

// Use this macro as the return value for continuing the propagation of
// an event handler.
const EVENT_PROPAGATE bool = false // C.GDK_EVENT_PROPAGATE
// Use this macro as the return value for stopping the propagation of
// an event handler.
const EVENT_STOP bool = true // C.GDK_EVENT_STOP
