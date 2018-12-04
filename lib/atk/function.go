// This is a generated file - DO NOT EDIT

package atk

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <atk/atk.h>
// #include <stdlib.h>
import "C"

// Unsupported : atk_add_focus_tracker : unsupported parameter focus_tracker : no type generator for EventListener (AtkEventListener) for param focus_tracker

// Unsupported : atk_add_global_event_listener : unsupported parameter listener : no type generator for GObject.SignalEmissionHook (GSignalEmissionHook) for param listener

// Unsupported : atk_add_key_event_listener : unsupported parameter listener : no type generator for KeySnoopFunc (AtkKeySnoopFunc) for param listener

// Blacklisted : atk_attribute_set_free

// Unsupported : atk_focus_tracker_init : unsupported parameter init : no type generator for EventListenerInit (AtkEventListenerInit) for param init

// FocusTrackerNotify is a wrapper around the C function atk_focus_tracker_notify.
func FocusTrackerNotify(object *Object) {
	c_object := (*C.AtkObject)(C.NULL)
	if object != nil {
		c_object = (*C.AtkObject)(object.ToC())
	}

	C.atk_focus_tracker_notify(c_object)

	return
}

// GetDefaultRegistry is a wrapper around the C function atk_get_default_registry.
func GetDefaultRegistry() *Registry {
	retC := C.atk_get_default_registry()
	retGo := RegistryNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetRoot is a wrapper around the C function atk_get_root.
func GetRoot() *Object {
	retC := C.atk_get_root()
	retGo := ObjectNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetToolkitName is a wrapper around the C function atk_get_toolkit_name.
func GetToolkitName() string {
	retC := C.atk_get_toolkit_name()
	retGo := C.GoString(retC)

	return retGo
}

// GetToolkitVersion is a wrapper around the C function atk_get_toolkit_version.
func GetToolkitVersion() string {
	retC := C.atk_get_toolkit_version()
	retGo := C.GoString(retC)

	return retGo
}

// RemoveFocusTracker is a wrapper around the C function atk_remove_focus_tracker.
func RemoveFocusTracker(trackerId uint32) {
	c_tracker_id := (C.guint)(trackerId)

	C.atk_remove_focus_tracker(c_tracker_id)

	return
}

// RemoveGlobalEventListener is a wrapper around the C function atk_remove_global_event_listener.
func RemoveGlobalEventListener(listenerId uint32) {
	c_listener_id := (C.guint)(listenerId)

	C.atk_remove_global_event_listener(c_listener_id)

	return
}

// RemoveKeyEventListener is a wrapper around the C function atk_remove_key_event_listener.
func RemoveKeyEventListener(listenerId uint32) {
	c_listener_id := (C.guint)(listenerId)

	C.atk_remove_key_event_listener(c_listener_id)

	return
}
