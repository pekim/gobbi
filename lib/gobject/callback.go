// This is a generated file - DO NOT EDIT

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

	void callback_closurenotifyHandler(GObject *, gpointer, GClosure*);

*/
/*

	gboolean callback_signalaccumulatorHandler(GObject *, GSignalInvocationHint*, GValue*, GValue*, gpointer);

*/
/*

	void callback_togglenotifyHandler(GObject *, gpointer, GObject*, gboolean);

*/
/*

	void callback_weaknotifyHandler(GObject *, gpointer, GObject*);

*/
import "C"

// Unsupported : callback BaseFinalizeFunc : no [user_]data param

// Unsupported : callback BaseInitFunc : no [user_]data param

// Unsupported : callback BoxedCopyFunc : no [user_]data param

// Unsupported : callback BoxedFreeFunc : no [user_]data param

// Unsupported : callback Callback : no [user_]data param

// Unsupported : callback ClassFinalizeFunc : no [user_]data param

// Unsupported : callback ClassInitFunc : no [user_]data param

// Unsupported : callback ClosureMarshal : no [user_]data param

var callbackClosurenotifyId int
var callbackClosurenotifyMap = make(map[int]ClosurenotifyCallback)
var callbackClosurenotifyLock sync.RWMutex

// ClosurenotifyCallback is a callback function for a 'ClosureNotify' callback.
type ClosurenotifyCallback func(closure *Closure)

//export callback_closurenotifyHandler
func callback_closurenotifyHandler(_ *C.GObject, c_data C.gpointer, c_closure *C.GClosure) {
	callbackClosurenotifyLock.RLock()
	defer callbackClosurenotifyLock.RUnlock()

	closure := ClosureNewFromC(unsafe.Pointer(c_closure))

	index := int(uintptr(c_data))
	callback := callbackClosurenotifyMap[index]
	callback(closure)
}

// Unsupported : callback InstanceInitFunc : no [user_]data param

// Unsupported : callback InterfaceFinalizeFunc : no [user_]data param

// Unsupported : callback InterfaceInitFunc : no [user_]data param

// Unsupported : callback ObjectFinalizeFunc : no [user_]data param

// Unsupported : callback ObjectGetPropertyFunc : no [user_]data param

// Unsupported : callback ObjectSetPropertyFunc : no [user_]data param

var callbackSignalaccumulatorId int
var callbackSignalaccumulatorMap = make(map[int]SignalaccumulatorCallback)
var callbackSignalaccumulatorLock sync.RWMutex

// SignalaccumulatorCallback is a callback function for a 'SignalAccumulator' callback.
type SignalaccumulatorCallback func(ihint *SignalInvocationHint, returnAccu *Value, handlerReturn *Value) bool

//export callback_signalaccumulatorHandler
func callback_signalaccumulatorHandler(_ *C.GObject, c_ihint *C.GSignalInvocationHint, c_return_accu *C.GValue, c_handler_return *C.GValue, c_data C.gpointer) C.gboolean {
	callbackSignalaccumulatorLock.RLock()
	defer callbackSignalaccumulatorLock.RUnlock()

	ihint := SignalInvocationHintNewFromC(unsafe.Pointer(c_ihint))

	returnAccu := ValueNewFromC(unsafe.Pointer(c_return_accu))

	handlerReturn := ValueNewFromC(unsafe.Pointer(c_handler_return))

	index := int(uintptr(c_data))
	callback := callbackSignalaccumulatorMap[index]
	retGo := callback(ihint, returnAccu, handlerReturn)
	retC :=
		boolToGboolean(retGo)
	return retC
}

// Unsupported : callback SignalEmissionHook : unsupported parameter param_values :

var callbackTogglenotifyId int
var callbackTogglenotifyMap = make(map[int]TogglenotifyCallback)
var callbackTogglenotifyLock sync.RWMutex

// TogglenotifyCallback is a callback function for a 'ToggleNotify' callback.
type TogglenotifyCallback func(object *Object, isLastRef bool)

//export callback_togglenotifyHandler
func callback_togglenotifyHandler(_ *C.GObject, c_data C.gpointer, c_object *C.GObject, c_is_last_ref C.gboolean) {
	callbackTogglenotifyLock.RLock()
	defer callbackTogglenotifyLock.RUnlock()

	object := ObjectNewFromC(unsafe.Pointer(c_object))

	isLastRef := c_is_last_ref == C.TRUE

	index := int(uintptr(c_data))
	callback := callbackTogglenotifyMap[index]
	callback(object, isLastRef)
}

// Unsupported : callback TypeClassCacheFunc : no [user_]data param

// Unsupported : callback TypePluginCompleteInterfaceInfo : no [user_]data param

// Unsupported : callback TypePluginCompleteTypeInfo : no [user_]data param

// Unsupported : callback TypePluginUnuse : no [user_]data param

// Unsupported : callback TypePluginUse : no [user_]data param

// Unsupported : callback VaClosureMarshal : no [user_]data param

// Unsupported : callback ValueTransform : no [user_]data param

var callbackWeaknotifyId int
var callbackWeaknotifyMap = make(map[int]WeaknotifyCallback)
var callbackWeaknotifyLock sync.RWMutex

// WeaknotifyCallback is a callback function for a 'WeakNotify' callback.
type WeaknotifyCallback func(whereTheObjectWas *Object)

//export callback_weaknotifyHandler
func callback_weaknotifyHandler(_ *C.GObject, c_data C.gpointer, c_where_the_object_was *C.GObject) {
	callbackWeaknotifyLock.RLock()
	defer callbackWeaknotifyLock.RUnlock()

	whereTheObjectWas := ObjectNewFromC(unsafe.Pointer(c_where_the_object_was))

	index := int(uintptr(c_data))
	callback := callbackWeaknotifyMap[index]
	callback(whereTheObjectWas)
}
