// This is a generated file - DO NOT EDIT
// +build gobject_2.26 gobject_2.28 gobject_2.30 gobject_2.32 gobject_2.34 gobject_2.36 gobject_2.38 gobject_2.42 gobject_2.44 gobject_2.46 gobject_2.54

package gobject

import (
	"sync"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <glib-object.h>
// #include <stdlib.h>
/*

	gboolean callback_bindingtransformfuncHandler(GObject *, GBinding*, GValue*, GValue*, gpointer);

*/
import "C"

var callbackBindingtransformfuncId int
var callbackBindingtransformfuncMap = make(map[int]BindingtransformfuncCallback)
var callbackBindingtransformfuncLock sync.RWMutex

// BindingtransformfuncCallback is a callback function for a 'BindingTransformFunc' callback.
type BindingtransformfuncCallback func(binding *Binding, fromValue *Value, toValue *Value) bool

//export callback_bindingtransformfuncHandler
func callback_bindingtransformfuncHandler(_ *C.GObject, c_binding *C.GBinding, c_from_value *C.GValue, c_to_value *C.GValue, c_user_data C.gpointer) C.gboolean {
	callbackBindingtransformfuncLock.RLock()
	defer callbackBindingtransformfuncLock.RUnlock()

	binding := BindingNewFromC(unsafe.Pointer(c_binding))

	fromValue := ValueNewFromC(unsafe.Pointer(c_from_value))

	toValue := ValueNewFromC(unsafe.Pointer(c_to_value))

	index := int(uintptr(c_user_data))
	callback := callbackBindingtransformfuncMap[index]
	retGo := callback(binding, fromValue, toValue)
	retC :=
		boolToGboolean(retGo)
	return retC
}
