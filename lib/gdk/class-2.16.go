// This is a generated file - DO NOT EDIT
// +build gdk_2.16 gdk_2.18 gdk_2.20 gdk_2.22 gdk_2.24 gdk_3.0 gdk_3.4 gdk_3.8 gdk_3.10 gdk_3.12 gdk_3.14 gdk_3.16 gdk_3.18 gdk_3.20 gdk_3.22

package gdk

import "sync"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <stdlib.h>
/*

	void Keymap_stateChangedHandler();

	static gulong Keymap_signal_connect_state_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "state-changed", Keymap_stateChangedHandler, data);
	}

*/
import "C"

var signalStateChangedId int
var signalStateChangedMap = make(map[int]KeymapSignalStateChangedCallback)
var signalStateChangedLock sync.Mutex

// KeymapSignalStateChangedCallback is a callback function for a 'state-changed' signal emitted from a Keymap.
type KeymapSignalStateChangedCallback func()

func Keymap_stateChangedHandler() {}

// GetCapsLockState is a wrapper around the C function gdk_keymap_get_caps_lock_state.
func (recv *Keymap) GetCapsLockState() bool {
	retC := C.gdk_keymap_get_caps_lock_state((*C.GdkKeymap)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}
