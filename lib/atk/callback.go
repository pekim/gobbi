// This is a generated file - DO NOT EDIT

package atk

import (
	"sync"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <atk/atk.h>
// #include <stdlib.h>
/*

	void callback_eventlistenerHandler(GObject *, AtkObject*, gpointer);

*/
/*

	void callback_eventlistenerinitHandler(GObject *, gpointer);

*/
/*

	void callback_focushandlerHandler(GObject *, AtkObject*, gboolean, gpointer);

*/
/*

	gboolean callback_functionHandler(GObject *, gpointer, gpointer);

*/
/*

	gint callback_keysnoopfuncHandler(GObject *, AtkKeyEventStruct*, gpointer, gpointer);

*/
/*

	void callback_propertychangehandlerHandler(GObject *, AtkObject*, AtkPropertyValues*, gpointer);

*/
import "C"

var callbackEventlistenerId int
var callbackEventlistenerMap = make(map[int]EventlistenerCallback)
var callbackEventlistenerLock sync.RWMutex

// EventlistenerCallback is a callback function for a 'EventListener' callback.
type EventlistenerCallback func(obj *Object)

//export callback_eventlistenerHandler
func callback_eventlistenerHandler(_ *C.GObject, c_obj *C.AtkObject, data C.gpointer) {
	callbackEventlistenerLock.RLock()
	defer callbackEventlistenerLock.RUnlock()

	obj := ObjectNewFromC(unsafe.Pointer(c_obj))

	index := int(uintptr(data))
	callback := callbackEventlistenerMap[index].callback
	callback(obj)
}

var callbackEventlistenerinitId int
var callbackEventlistenerinitMap = make(map[int]EventlistenerinitCallback)
var callbackEventlistenerinitLock sync.RWMutex

// EventlistenerinitCallback is a callback function for a 'EventListenerInit' callback.
type EventlistenerinitCallback func()

//export callback_eventlistenerinitHandler
func callback_eventlistenerinitHandler(_ *C.GObject, data C.gpointer) {
	callbackEventlistenerinitLock.RLock()
	defer callbackEventlistenerinitLock.RUnlock()

	index := int(uintptr(data))
	callback := callbackEventlistenerinitMap[index].callback
	callback()
}

var callbackFocushandlerId int
var callbackFocushandlerMap = make(map[int]FocushandlerCallback)
var callbackFocushandlerLock sync.RWMutex

// FocushandlerCallback is a callback function for a 'FocusHandler' callback.
type FocushandlerCallback func(object *Object, focusIn bool)

//export callback_focushandlerHandler
func callback_focushandlerHandler(_ *C.GObject, c_object *C.AtkObject, c_focus_in C.gboolean, data C.gpointer) {
	callbackFocushandlerLock.RLock()
	defer callbackFocushandlerLock.RUnlock()

	object := ObjectNewFromC(unsafe.Pointer(c_object))

	focusIn := c_focus_in == C.TRUE

	index := int(uintptr(data))
	callback := callbackFocushandlerMap[index].callback
	callback(object, focusIn)
}

var callbackFunctionId int
var callbackFunctionMap = make(map[int]FunctionCallback)
var callbackFunctionLock sync.RWMutex

// FunctionCallback is a callback function for a 'Function' callback.
type FunctionCallback func() bool

//export callback_functionHandler
func callback_functionHandler(_ *C.GObject, data C.gpointer) C.gboolean {
	callbackFunctionLock.RLock()
	defer callbackFunctionLock.RUnlock()

	index := int(uintptr(data))
	callback := callbackFunctionMap[index].callback
	retGo := callback(userData)
	retC :=
		boolToGboolean(retGo)
	return retC
}

var callbackKeysnoopfuncId int
var callbackKeysnoopfuncMap = make(map[int]KeysnoopfuncCallback)
var callbackKeysnoopfuncLock sync.RWMutex

// KeysnoopfuncCallback is a callback function for a 'KeySnoopFunc' callback.
type KeysnoopfuncCallback func(event *KeyEventStruct) int32

//export callback_keysnoopfuncHandler
func callback_keysnoopfuncHandler(_ *C.GObject, c_event *C.AtkKeyEventStruct, data C.gpointer) C.gint {
	callbackKeysnoopfuncLock.RLock()
	defer callbackKeysnoopfuncLock.RUnlock()

	event := KeyEventStructNewFromC(unsafe.Pointer(c_event))

	index := int(uintptr(data))
	callback := callbackKeysnoopfuncMap[index].callback
	retGo := callback(event, userData)
	retC :=
		(C.gint)(retGo)
	return retC
}

var callbackPropertychangehandlerId int
var callbackPropertychangehandlerMap = make(map[int]PropertychangehandlerCallback)
var callbackPropertychangehandlerLock sync.RWMutex

// PropertychangehandlerCallback is a callback function for a 'PropertyChangeHandler' callback.
type PropertychangehandlerCallback func(obj *Object, vals *PropertyValues)

//export callback_propertychangehandlerHandler
func callback_propertychangehandlerHandler(_ *C.GObject, c_obj *C.AtkObject, c_vals *C.AtkPropertyValues, data C.gpointer) {
	callbackPropertychangehandlerLock.RLock()
	defer callbackPropertychangehandlerLock.RUnlock()

	obj := ObjectNewFromC(unsafe.Pointer(c_obj))

	vals := PropertyValuesNewFromC(unsafe.Pointer(c_vals))

	index := int(uintptr(data))
	callback := callbackPropertychangehandlerMap[index].callback
	callback(obj, vals)
}
