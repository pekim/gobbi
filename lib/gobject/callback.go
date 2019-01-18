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

	void callback_basefinalizefuncHandler(GObject *, gpointer, gpointer);

*/
/*

	void callback_baseinitfuncHandler(GObject *, gpointer, gpointer);

*/
/*

	void callback_callbackHandler(GObject *, gpointer);

*/
/*

	void callback_closurenotifyHandler(GObject *, gpointer, GClosure*, gpointer);

*/
/*

	void callback_instanceinitfuncHandler(GObject *, GTypeInstance*, gpointer, gpointer);

*/
/*

	void callback_objectfinalizefuncHandler(GObject *, GObject*, gpointer);

*/
/*

	void callback_objectgetpropertyfuncHandler(GObject *, GObject*, guint, GValue*, GParamSpec*, gpointer);

*/
/*

	void callback_objectsetpropertyfuncHandler(GObject *, GObject*, guint, const GValue*, GParamSpec*, gpointer);

*/
/*

	gboolean callback_signalaccumulatorHandler(GObject *, GSignalInvocationHint*, GValue*, const GValue*, gpointer, gpointer);

*/
/*

	void callback_togglenotifyHandler(GObject *, gpointer, GObject*, gboolean, gpointer);

*/
/*

	void callback_typeplugincompleteinterfaceinfoHandler(GObject *, GTypePlugin*, GType, GType, GInterfaceInfo*, gpointer);

*/
/*

	void callback_typeplugincompletetypeinfoHandler(GObject *, GTypePlugin*, GType, GTypeInfo*, GTypeValueTable*, gpointer);

*/
/*

	void callback_typepluginunuseHandler(GObject *, GTypePlugin*, gpointer);

*/
/*

	void callback_typepluginuseHandler(GObject *, GTypePlugin*, gpointer);

*/
/*

	void callback_valuetransformHandler(GObject *, const GValue*, GValue*, gpointer);

*/
/*

	void callback_weaknotifyHandler(GObject *, gpointer, GObject*, gpointer);

*/
import "C"

var callbackBasefinalizefuncId int
var callbackBasefinalizefuncMap = make(map[int]BasefinalizefuncCallback)
var callbackBasefinalizefuncLock sync.RWMutex

// BasefinalizefuncCallback is a callback function for a 'BaseFinalizeFunc' callback.
type BasefinalizefuncCallback func(gClass *TypeClass)

//export callback_basefinalizefuncHandler
func callback_basefinalizefuncHandler(_ *C.GObject, c_g_class *C.GTypeClass, data C.gpointer) {
	callbackBasefinalizefuncLock.RLock()
	defer callbackBasefinalizefuncLock.RUnlock()

	gClass := TypeClassNewFromC(unsafe.Pointer(c_g_class))

	index := int(uintptr(data))
	callback := callbackBasefinalizefuncMap[index].callback
	callback(gClass)
}

var callbackBaseinitfuncId int
var callbackBaseinitfuncMap = make(map[int]BaseinitfuncCallback)
var callbackBaseinitfuncLock sync.RWMutex

// BaseinitfuncCallback is a callback function for a 'BaseInitFunc' callback.
type BaseinitfuncCallback func(gClass *TypeClass)

//export callback_baseinitfuncHandler
func callback_baseinitfuncHandler(_ *C.GObject, c_g_class *C.GTypeClass, data C.gpointer) {
	callbackBaseinitfuncLock.RLock()
	defer callbackBaseinitfuncLock.RUnlock()

	gClass := TypeClassNewFromC(unsafe.Pointer(c_g_class))

	index := int(uintptr(data))
	callback := callbackBaseinitfuncMap[index].callback
	callback(gClass)
}

// Unsupported : callback BoxedCopyFunc : unsupported parameter boxed : no type generator for gpointer (gpointer) for param boxed

// Unsupported : callback BoxedFreeFunc : unsupported parameter boxed : no type generator for gpointer (gpointer) for param boxed

var callbackCallbackId int
var callbackCallbackMap = make(map[int]CallbackCallback)
var callbackCallbackLock sync.RWMutex

// CallbackCallback is a callback function for a 'Callback' callback.
type CallbackCallback func()

//export callback_callbackHandler
func callback_callbackHandler(_ *C.GObject, data C.gpointer) {
	callbackCallbackLock.RLock()
	defer callbackCallbackLock.RUnlock()

	index := int(uintptr(data))
	callback := callbackCallbackMap[index].callback
	callback()
}

// Unsupported : callback ClassFinalizeFunc : unsupported parameter class_data : no type generator for gpointer (gpointer) for param class_data

// Unsupported : callback ClassInitFunc : unsupported parameter class_data : no type generator for gpointer (gpointer) for param class_data

// Unsupported : callback ClosureMarshal : unsupported parameter param_values :

var callbackClosurenotifyId int
var callbackClosurenotifyMap = make(map[int]ClosurenotifyCallback)
var callbackClosurenotifyLock sync.RWMutex

// ClosurenotifyCallback is a callback function for a 'ClosureNotify' callback.
type ClosurenotifyCallback func(closure *Closure)

//export callback_closurenotifyHandler
func callback_closurenotifyHandler(_ *C.GObject, c_closure *C.GClosure, data C.gpointer) {
	callbackClosurenotifyLock.RLock()
	defer callbackClosurenotifyLock.RUnlock()

	closure := ClosureNewFromC(unsafe.Pointer(c_closure))

	index := int(uintptr(data))
	callback := callbackClosurenotifyMap[index].callback
	callback(data, closure)
}

var callbackInstanceinitfuncId int
var callbackInstanceinitfuncMap = make(map[int]InstanceinitfuncCallback)
var callbackInstanceinitfuncLock sync.RWMutex

// InstanceinitfuncCallback is a callback function for a 'InstanceInitFunc' callback.
type InstanceinitfuncCallback func(instance *TypeInstance, gClass *TypeClass)

//export callback_instanceinitfuncHandler
func callback_instanceinitfuncHandler(_ *C.GObject, c_instance *C.GTypeInstance, c_g_class *C.GTypeClass, data C.gpointer) {
	callbackInstanceinitfuncLock.RLock()
	defer callbackInstanceinitfuncLock.RUnlock()

	instance := TypeInstanceNewFromC(unsafe.Pointer(c_instance))

	gClass := TypeClassNewFromC(unsafe.Pointer(c_g_class))

	index := int(uintptr(data))
	callback := callbackInstanceinitfuncMap[index].callback
	callback(instance, gClass)
}

// Unsupported : callback InterfaceFinalizeFunc : unsupported parameter iface_data : no type generator for gpointer (gpointer) for param iface_data

// Unsupported : callback InterfaceInitFunc : unsupported parameter iface_data : no type generator for gpointer (gpointer) for param iface_data

var callbackObjectfinalizefuncId int
var callbackObjectfinalizefuncMap = make(map[int]ObjectfinalizefuncCallback)
var callbackObjectfinalizefuncLock sync.RWMutex

// ObjectfinalizefuncCallback is a callback function for a 'ObjectFinalizeFunc' callback.
type ObjectfinalizefuncCallback func(object *Object)

//export callback_objectfinalizefuncHandler
func callback_objectfinalizefuncHandler(_ *C.GObject, c_object *C.GObject, data C.gpointer) {
	callbackObjectfinalizefuncLock.RLock()
	defer callbackObjectfinalizefuncLock.RUnlock()

	object := ObjectNewFromC(unsafe.Pointer(c_object))

	index := int(uintptr(data))
	callback := callbackObjectfinalizefuncMap[index].callback
	callback(object)
}

var callbackObjectgetpropertyfuncId int
var callbackObjectgetpropertyfuncMap = make(map[int]ObjectgetpropertyfuncCallback)
var callbackObjectgetpropertyfuncLock sync.RWMutex

// ObjectgetpropertyfuncCallback is a callback function for a 'ObjectGetPropertyFunc' callback.
type ObjectgetpropertyfuncCallback func(object *Object, propertyId uint32, value *Value, pspec *ParamSpec)

//export callback_objectgetpropertyfuncHandler
func callback_objectgetpropertyfuncHandler(_ *C.GObject, c_object *C.GObject, c_property_id C.guint, c_value *C.GValue, c_pspec *C.GParamSpec, data C.gpointer) {
	callbackObjectgetpropertyfuncLock.RLock()
	defer callbackObjectgetpropertyfuncLock.RUnlock()

	object := ObjectNewFromC(unsafe.Pointer(c_object))

	propertyId := uint32(c_property_id)

	value := ValueNewFromC(unsafe.Pointer(c_value))

	pspec := ParamSpecNewFromC(unsafe.Pointer(c_pspec))

	index := int(uintptr(data))
	callback := callbackObjectgetpropertyfuncMap[index].callback
	callback(object, propertyId, value, pspec)
}

var callbackObjectsetpropertyfuncId int
var callbackObjectsetpropertyfuncMap = make(map[int]ObjectsetpropertyfuncCallback)
var callbackObjectsetpropertyfuncLock sync.RWMutex

// ObjectsetpropertyfuncCallback is a callback function for a 'ObjectSetPropertyFunc' callback.
type ObjectsetpropertyfuncCallback func(object *Object, propertyId uint32, value *Value, pspec *ParamSpec)

//export callback_objectsetpropertyfuncHandler
func callback_objectsetpropertyfuncHandler(_ *C.GObject, c_object *C.GObject, c_property_id C.guint, c_value *C.GValue, c_pspec *C.GParamSpec, data C.gpointer) {
	callbackObjectsetpropertyfuncLock.RLock()
	defer callbackObjectsetpropertyfuncLock.RUnlock()

	object := ObjectNewFromC(unsafe.Pointer(c_object))

	propertyId := uint32(c_property_id)

	value := ValueNewFromC(unsafe.Pointer(c_value))

	pspec := ParamSpecNewFromC(unsafe.Pointer(c_pspec))

	index := int(uintptr(data))
	callback := callbackObjectsetpropertyfuncMap[index].callback
	callback(object, propertyId, value, pspec)
}

var callbackSignalaccumulatorId int
var callbackSignalaccumulatorMap = make(map[int]SignalaccumulatorCallback)
var callbackSignalaccumulatorLock sync.RWMutex

// SignalaccumulatorCallback is a callback function for a 'SignalAccumulator' callback.
type SignalaccumulatorCallback func(ihint *SignalInvocationHint, returnAccu *Value, handlerReturn *Value) bool

//export callback_signalaccumulatorHandler
func callback_signalaccumulatorHandler(_ *C.GObject, c_ihint *C.GSignalInvocationHint, c_return_accu *C.GValue, c_handler_return *C.GValue, data C.gpointer) C.gboolean {
	callbackSignalaccumulatorLock.RLock()
	defer callbackSignalaccumulatorLock.RUnlock()

	ihint := SignalInvocationHintNewFromC(unsafe.Pointer(c_ihint))

	returnAccu := ValueNewFromC(unsafe.Pointer(c_return_accu))

	handlerReturn := ValueNewFromC(unsafe.Pointer(c_handler_return))

	index := int(uintptr(data))
	callback := callbackSignalaccumulatorMap[index].callback
	retGo := callback(ihint, returnAccu, handlerReturn, data)
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
func callback_togglenotifyHandler(_ *C.GObject, c_object *C.GObject, c_is_last_ref C.gboolean, data C.gpointer) {
	callbackTogglenotifyLock.RLock()
	defer callbackTogglenotifyLock.RUnlock()

	object := ObjectNewFromC(unsafe.Pointer(c_object))

	isLastRef := c_is_last_ref == C.TRUE

	index := int(uintptr(data))
	callback := callbackTogglenotifyMap[index].callback
	callback(data, object, isLastRef)
}

// Unsupported : callback TypeClassCacheFunc : unsupported parameter cache_data : no type generator for gpointer (gpointer) for param cache_data

var callbackTypeplugincompleteinterfaceinfoId int
var callbackTypeplugincompleteinterfaceinfoMap = make(map[int]TypeplugincompleteinterfaceinfoCallback)
var callbackTypeplugincompleteinterfaceinfoLock sync.RWMutex

// TypeplugincompleteinterfaceinfoCallback is a callback function for a 'TypePluginCompleteInterfaceInfo' callback.
type TypeplugincompleteinterfaceinfoCallback func(plugin *TypePlugin, instanceType Type, interfaceType Type, info *InterfaceInfo)

//export callback_typeplugincompleteinterfaceinfoHandler
func callback_typeplugincompleteinterfaceinfoHandler(_ *C.GObject, c_plugin *C.GTypePlugin, c_instance_type C.GType, c_interface_type C.GType, c_info *C.GInterfaceInfo, data C.gpointer) {
	callbackTypeplugincompleteinterfaceinfoLock.RLock()
	defer callbackTypeplugincompleteinterfaceinfoLock.RUnlock()

	plugin := TypePluginNewFromC(unsafe.Pointer(c_plugin))

	instanceType := (c_instance_type)

	interfaceType := (c_interface_type)

	info := InterfaceInfoNewFromC(unsafe.Pointer(c_info))

	index := int(uintptr(data))
	callback := callbackTypeplugincompleteinterfaceinfoMap[index].callback
	callback(plugin, instanceType, interfaceType, info)
}

var callbackTypeplugincompletetypeinfoId int
var callbackTypeplugincompletetypeinfoMap = make(map[int]TypeplugincompletetypeinfoCallback)
var callbackTypeplugincompletetypeinfoLock sync.RWMutex

// TypeplugincompletetypeinfoCallback is a callback function for a 'TypePluginCompleteTypeInfo' callback.
type TypeplugincompletetypeinfoCallback func(plugin *TypePlugin, gType Type, info *TypeInfo, valueTable *TypeValueTable)

//export callback_typeplugincompletetypeinfoHandler
func callback_typeplugincompletetypeinfoHandler(_ *C.GObject, c_plugin *C.GTypePlugin, c_g_type C.GType, c_info *C.GTypeInfo, c_value_table *C.GTypeValueTable, data C.gpointer) {
	callbackTypeplugincompletetypeinfoLock.RLock()
	defer callbackTypeplugincompletetypeinfoLock.RUnlock()

	plugin := TypePluginNewFromC(unsafe.Pointer(c_plugin))

	gType := (c_g_type)

	info := TypeInfoNewFromC(unsafe.Pointer(c_info))

	valueTable := TypeValueTableNewFromC(unsafe.Pointer(c_value_table))

	index := int(uintptr(data))
	callback := callbackTypeplugincompletetypeinfoMap[index].callback
	callback(plugin, gType, info, valueTable)
}

var callbackTypepluginunuseId int
var callbackTypepluginunuseMap = make(map[int]TypepluginunuseCallback)
var callbackTypepluginunuseLock sync.RWMutex

// TypepluginunuseCallback is a callback function for a 'TypePluginUnuse' callback.
type TypepluginunuseCallback func(plugin *TypePlugin)

//export callback_typepluginunuseHandler
func callback_typepluginunuseHandler(_ *C.GObject, c_plugin *C.GTypePlugin, data C.gpointer) {
	callbackTypepluginunuseLock.RLock()
	defer callbackTypepluginunuseLock.RUnlock()

	plugin := TypePluginNewFromC(unsafe.Pointer(c_plugin))

	index := int(uintptr(data))
	callback := callbackTypepluginunuseMap[index].callback
	callback(plugin)
}

var callbackTypepluginuseId int
var callbackTypepluginuseMap = make(map[int]TypepluginuseCallback)
var callbackTypepluginuseLock sync.RWMutex

// TypepluginuseCallback is a callback function for a 'TypePluginUse' callback.
type TypepluginuseCallback func(plugin *TypePlugin)

//export callback_typepluginuseHandler
func callback_typepluginuseHandler(_ *C.GObject, c_plugin *C.GTypePlugin, data C.gpointer) {
	callbackTypepluginuseLock.RLock()
	defer callbackTypepluginuseLock.RUnlock()

	plugin := TypePluginNewFromC(unsafe.Pointer(c_plugin))

	index := int(uintptr(data))
	callback := callbackTypepluginuseMap[index].callback
	callback(plugin)
}

// Unsupported : callback VaClosureMarshal : unsupported parameter args : no type generator for va_list (va_list) for param args

var callbackValuetransformId int
var callbackValuetransformMap = make(map[int]ValuetransformCallback)
var callbackValuetransformLock sync.RWMutex

// ValuetransformCallback is a callback function for a 'ValueTransform' callback.
type ValuetransformCallback func(srcValue *Value, destValue *Value)

//export callback_valuetransformHandler
func callback_valuetransformHandler(_ *C.GObject, c_src_value *C.GValue, c_dest_value *C.GValue, data C.gpointer) {
	callbackValuetransformLock.RLock()
	defer callbackValuetransformLock.RUnlock()

	srcValue := ValueNewFromC(unsafe.Pointer(c_src_value))

	destValue := ValueNewFromC(unsafe.Pointer(c_dest_value))

	index := int(uintptr(data))
	callback := callbackValuetransformMap[index].callback
	callback(srcValue, destValue)
}

var callbackWeaknotifyId int
var callbackWeaknotifyMap = make(map[int]WeaknotifyCallback)
var callbackWeaknotifyLock sync.RWMutex

// WeaknotifyCallback is a callback function for a 'WeakNotify' callback.
type WeaknotifyCallback func(whereTheObjectWas *Object)

//export callback_weaknotifyHandler
func callback_weaknotifyHandler(_ *C.GObject, c_where_the_object_was *C.GObject, data C.gpointer) {
	callbackWeaknotifyLock.RLock()
	defer callbackWeaknotifyLock.RUnlock()

	whereTheObjectWas := ObjectNewFromC(unsafe.Pointer(c_where_the_object_was))

	index := int(uintptr(data))
	callback := callbackWeaknotifyMap[index].callback
	callback(data, whereTheObjectWas)
}
