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

type signalKeymapStateChangedDetail struct {
	callback  KeymapSignalStateChangedCallback
	handlerID C.gulong
}

var signalKeymapStateChangedId int
var signalKeymapStateChangedMap = make(map[int]signalKeymapStateChangedDetail)
var signalKeymapStateChangedLock sync.Mutex

// KeymapSignalStateChangedCallback is a callback function for a 'state-changed' signal emitted from a Keymap.
type KeymapSignalStateChangedCallback func()

/*
ConnectStateChanged connects the callback to the 'state-changed' signal for the Keymap.

The returned value represents the connection, and may be passed to DisconnectStateChanged to remove it.
*/
func (recv *Keymap) ConnectStateChanged(callback KeymapSignalStateChangedCallback) int {
	signalKeymapStateChangedLock.Lock()
	defer signalKeymapStateChangedLock.Unlock()

	signalKeymapStateChangedId++
	instance := C.gpointer(recv.Object().ToC())
	handlerID := C.Keymap_signal_connect_state_changed(instance, C.gpointer(uintptr(signalKeymapStateChangedId)))

	detail := signalKeymapStateChangedDetail{callback, handlerID}
	signalKeymapStateChangedMap[signalKeymapStateChangedId] = detail

	return signalKeymapStateChangedId
}

/*
DisconnectStateChanged disconnects a callback from the 'state-changed' signal for the Keymap.

The connectionID should be a value returned from a call to ConnectStateChanged.
*/
func (recv *Keymap) DisconnectStateChanged(connectionID int) {
	signalKeymapStateChangedLock.Lock()
	defer signalKeymapStateChangedLock.Unlock()

	detail, exists := signalKeymapStateChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.Object().ToC())
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalKeymapStateChangedMap, connectionID)
}

//export Keymap_stateChangedHandler
func Keymap_stateChangedHandler(_ *C.GObject, data C.gpointer) {
	index := int(uintptr(data))
	callback := signalKeymapStateChangedMap[index].callback
	callback()
}

// GetCapsLockState is a wrapper around the C function gdk_keymap_get_caps_lock_state.
func (recv *Keymap) GetCapsLockState() bool {
	retC := C.gdk_keymap_get_caps_lock_state((*C.GdkKeymap)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}
