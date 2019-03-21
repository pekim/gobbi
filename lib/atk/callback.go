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

	gboolean callback_functionHandler(GObject *, gpointer);

*/
/*

	gint callback_keysnoopfuncHandler(GObject *, AtkKeyEventStruct*, gpointer);

*/
import "C"

// Unsupported : callback EventListener : no [user_]data param

// Unsupported : callback EventListenerInit : no [user_]data param

// Unsupported : callback FocusHandler : no [user_]data param

var callbackFunctionId int
var callbackFunctionMap = make(map[int]FunctionCallback)
var callbackFunctionLock sync.RWMutex

// FunctionCallback is a callback function for a 'Function' callback.
type FunctionCallback func() bool

//export callback_functionHandler
func callback_functionHandler(_ *C.GObject, c_user_data C.gpointer) C.gboolean {
	callbackFunctionLock.RLock()
	defer callbackFunctionLock.RUnlock()

	index := int(uintptr(c_user_data))
	callback := callbackFunctionMap[index]
	retGo := callback()
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
func callback_keysnoopfuncHandler(_ *C.GObject, c_event *C.AtkKeyEventStruct, c_user_data C.gpointer) C.gint {
	callbackKeysnoopfuncLock.RLock()
	defer callbackKeysnoopfuncLock.RUnlock()

	event := KeyEventStructNewFromC(unsafe.Pointer(c_event))

	index := int(uintptr(c_user_data))
	callback := callbackKeysnoopfuncMap[index]
	retGo := callback(event)
	retC :=
		(C.gint)(retGo)
	return retC
}

// Unsupported : callback PropertyChangeHandler : no [user_]data param
