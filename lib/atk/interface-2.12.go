// This is a generated file - DO NOT EDIT
// +build atk_2.12

package atk

import "sync"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <atk/atk.h>
// #include <stdlib.h>
/*

	void value_valueChangedHandler(GObject *, gdouble, gchar*, gpointer);

	static gulong Value_signal_connect_value_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "value-changed", G_CALLBACK(value_valueChangedHandler), data);
	}

*/
import "C"

type signalValueValueChangedDetail struct {
	callback  ValueSignalValueChangedCallback
	handlerID C.gulong
}

var signalValueValueChangedId int
var signalValueValueChangedMap = make(map[int]signalValueValueChangedDetail)
var signalValueValueChangedLock sync.RWMutex

// ValueSignalValueChangedCallback is a callback function for a 'value-changed' signal emitted from a Value.
type ValueSignalValueChangedCallback func(value float64, text string)

/*
ConnectValueChanged connects the callback to the 'value-changed' signal for the Value.

The returned value represents the connection, and may be passed to DisconnectValueChanged to remove it.
*/
func (recv *Value) ConnectValueChanged(callback ValueSignalValueChangedCallback) int {
	signalValueValueChangedLock.Lock()
	defer signalValueValueChangedLock.Unlock()

	signalValueValueChangedId++
	instance := C.gpointer(recv.native)
	handlerID := C.Value_signal_connect_value_changed(instance, C.gpointer(uintptr(signalValueValueChangedId)))

	detail := signalValueValueChangedDetail{callback, handlerID}
	signalValueValueChangedMap[signalValueValueChangedId] = detail

	return signalValueValueChangedId
}

/*
DisconnectValueChanged disconnects a callback from the 'value-changed' signal for the Value.

The connectionID should be a value returned from a call to ConnectValueChanged.
*/
func (recv *Value) DisconnectValueChanged(connectionID int) {
	signalValueValueChangedLock.Lock()
	defer signalValueValueChangedLock.Unlock()

	detail, exists := signalValueValueChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalValueValueChangedMap, connectionID)
}

//export value_valueChangedHandler
func value_valueChangedHandler(_ *C.GObject, c_value C.gdouble, c_text *C.gchar, data C.gpointer) {
	signalValueValueChangedLock.RLock()
	defer signalValueValueChangedLock.RUnlock()

	value := float64(c_value)

	text := C.GoString(c_text)

	index := int(uintptr(data))
	callback := signalValueValueChangedMap[index].callback
	callback(value, text)
}

// Blacklisted : atk_value_get_increment

// Blacklisted : atk_value_get_range

// Blacklisted : atk_value_get_sub_ranges

// Blacklisted : atk_value_get_value_and_text

// Blacklisted : atk_value_set_value
