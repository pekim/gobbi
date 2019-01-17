// This is a generated file - DO NOT EDIT

package gobject

import "sync"

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

var callbackBaseinitfuncId int
var callbackBaseinitfuncMap = make(map[int]BaseinitfuncCallback)
var callbackBaseinitfuncLock sync.RWMutex

// BaseinitfuncCallback is a callback function for a 'BaseInitFunc' callback.
type BaseinitfuncCallback func(gClass *TypeClass)

// Unsupported : callback BoxedCopyFunc : unsupported parameter boxed : no type generator for gpointer (gpointer) for param boxed

// Unsupported : callback BoxedFreeFunc : unsupported parameter boxed : no type generator for gpointer (gpointer) for param boxed

var callbackCallbackId int
var callbackCallbackMap = make(map[int]CallbackCallback)
var callbackCallbackLock sync.RWMutex

// CallbackCallback is a callback function for a 'Callback' callback.
type CallbackCallback func()

// Unsupported : callback ClassFinalizeFunc : unsupported parameter class_data : no type generator for gpointer (gpointer) for param class_data

// Unsupported : callback ClassInitFunc : unsupported parameter class_data : no type generator for gpointer (gpointer) for param class_data

// Unsupported : callback ClosureMarshal : unsupported parameter param_values :

var callbackClosurenotifyId int
var callbackClosurenotifyMap = make(map[int]ClosurenotifyCallback)
var callbackClosurenotifyLock sync.RWMutex

// ClosurenotifyCallback is a callback function for a 'ClosureNotify' callback.
type ClosurenotifyCallback func(closure *Closure)

var callbackInstanceinitfuncId int
var callbackInstanceinitfuncMap = make(map[int]InstanceinitfuncCallback)
var callbackInstanceinitfuncLock sync.RWMutex

// InstanceinitfuncCallback is a callback function for a 'InstanceInitFunc' callback.
type InstanceinitfuncCallback func(instance *TypeInstance, gClass *TypeClass)

// Unsupported : callback InterfaceFinalizeFunc : unsupported parameter iface_data : no type generator for gpointer (gpointer) for param iface_data

// Unsupported : callback InterfaceInitFunc : unsupported parameter iface_data : no type generator for gpointer (gpointer) for param iface_data

var callbackObjectfinalizefuncId int
var callbackObjectfinalizefuncMap = make(map[int]ObjectfinalizefuncCallback)
var callbackObjectfinalizefuncLock sync.RWMutex

// ObjectfinalizefuncCallback is a callback function for a 'ObjectFinalizeFunc' callback.
type ObjectfinalizefuncCallback func(object *Object)

var callbackObjectgetpropertyfuncId int
var callbackObjectgetpropertyfuncMap = make(map[int]ObjectgetpropertyfuncCallback)
var callbackObjectgetpropertyfuncLock sync.RWMutex

// ObjectgetpropertyfuncCallback is a callback function for a 'ObjectGetPropertyFunc' callback.
type ObjectgetpropertyfuncCallback func(object *Object, propertyId uint32, value *Value, pspec *ParamSpec)

var callbackObjectsetpropertyfuncId int
var callbackObjectsetpropertyfuncMap = make(map[int]ObjectsetpropertyfuncCallback)
var callbackObjectsetpropertyfuncLock sync.RWMutex

// ObjectsetpropertyfuncCallback is a callback function for a 'ObjectSetPropertyFunc' callback.
type ObjectsetpropertyfuncCallback func(object *Object, propertyId uint32, value *Value, pspec *ParamSpec)

var callbackSignalaccumulatorId int
var callbackSignalaccumulatorMap = make(map[int]SignalaccumulatorCallback)
var callbackSignalaccumulatorLock sync.RWMutex

// SignalaccumulatorCallback is a callback function for a 'SignalAccumulator' callback.
type SignalaccumulatorCallback func(ihint *SignalInvocationHint, returnAccu *Value, handlerReturn *Value) bool

// Unsupported : callback SignalEmissionHook : unsupported parameter param_values :

var callbackTogglenotifyId int
var callbackTogglenotifyMap = make(map[int]TogglenotifyCallback)
var callbackTogglenotifyLock sync.RWMutex

// TogglenotifyCallback is a callback function for a 'ToggleNotify' callback.
type TogglenotifyCallback func(object *Object, isLastRef bool)

// Unsupported : callback TypeClassCacheFunc : unsupported parameter cache_data : no type generator for gpointer (gpointer) for param cache_data

var callbackTypeplugincompleteinterfaceinfoId int
var callbackTypeplugincompleteinterfaceinfoMap = make(map[int]TypeplugincompleteinterfaceinfoCallback)
var callbackTypeplugincompleteinterfaceinfoLock sync.RWMutex

// TypeplugincompleteinterfaceinfoCallback is a callback function for a 'TypePluginCompleteInterfaceInfo' callback.
type TypeplugincompleteinterfaceinfoCallback func(plugin *TypePlugin, instanceType Type, interfaceType Type, info *InterfaceInfo)

var callbackTypeplugincompletetypeinfoId int
var callbackTypeplugincompletetypeinfoMap = make(map[int]TypeplugincompletetypeinfoCallback)
var callbackTypeplugincompletetypeinfoLock sync.RWMutex

// TypeplugincompletetypeinfoCallback is a callback function for a 'TypePluginCompleteTypeInfo' callback.
type TypeplugincompletetypeinfoCallback func(plugin *TypePlugin, gType Type, info *TypeInfo, valueTable *TypeValueTable)

var callbackTypepluginunuseId int
var callbackTypepluginunuseMap = make(map[int]TypepluginunuseCallback)
var callbackTypepluginunuseLock sync.RWMutex

// TypepluginunuseCallback is a callback function for a 'TypePluginUnuse' callback.
type TypepluginunuseCallback func(plugin *TypePlugin)

var callbackTypepluginuseId int
var callbackTypepluginuseMap = make(map[int]TypepluginuseCallback)
var callbackTypepluginuseLock sync.RWMutex

// TypepluginuseCallback is a callback function for a 'TypePluginUse' callback.
type TypepluginuseCallback func(plugin *TypePlugin)

// Unsupported : callback VaClosureMarshal : unsupported parameter args : no type generator for va_list (va_list) for param args

var callbackValuetransformId int
var callbackValuetransformMap = make(map[int]ValuetransformCallback)
var callbackValuetransformLock sync.RWMutex

// ValuetransformCallback is a callback function for a 'ValueTransform' callback.
type ValuetransformCallback func(srcValue *Value, destValue *Value)

var callbackWeaknotifyId int
var callbackWeaknotifyMap = make(map[int]WeaknotifyCallback)
var callbackWeaknotifyLock sync.RWMutex

// WeaknotifyCallback is a callback function for a 'WeakNotify' callback.
type WeaknotifyCallback func(whereTheObjectWas *Object)
