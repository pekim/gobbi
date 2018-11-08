// This is a generated file - DO NOT EDIT
// +build atk_1.20 atk_1.30 atk_2.8 atk_2.12

package atk

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <atk/atk.h>
// #include <stdlib.h>
import "C"

// Unsupported : atk_add_focus_tracker : unsupported parameter focus_tracker : no type generator for EventListener (AtkEventListener) for param focus_tracker

// Unsupported : atk_add_global_event_listener : unsupported parameter listener : no type generator for GObject.SignalEmissionHook (GSignalEmissionHook) for param listener

// Unsupported : atk_add_key_event_listener : unsupported parameter listener : no type generator for KeySnoopFunc (AtkKeySnoopFunc) for param listener

// Unsupported : atk_focus_tracker_init : unsupported parameter init : no type generator for EventListenerInit (AtkEventListenerInit) for param init

// GetVersion is a wrapper around the C function atk_get_version.
func GetVersion() string {
	retC := C.atk_get_version()
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : atk_text_free_ranges : unsupported parameter ranges :
