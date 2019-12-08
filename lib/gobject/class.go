// Code generated - DO NOT EDIT.

package gobject

import (
	callback "github.com/pekim/gobbi/internal/cgo/callback"
	gi "github.com/pekim/gobbi/internal/cgo/gi"
	glib "github.com/pekim/gobbi/lib/glib"
	"sync"
	"unsafe"
)

var bindingObject *gi.Object
var bindingObject_Once sync.Once

func bindingObject_Set() error {
	var err error
	bindingObject_Once.Do(func() {
		bindingObject, err = gi.ObjectNew("GObject", "Binding")
	})
	return err
}

type Binding struct {
	native unsafe.Pointer
}

func BindingNewFromNative(native unsafe.Pointer) *Binding {
	instance := &Binding{native: native}

	return instance
}

// Object upcasts to *Object
func (recv *Binding) Object() *Object {
	return ObjectNewFromNative(recv.Native())
}

// Equals compares this Binding with another Binding, and returns true if they represent the same Object.
func (recv *Binding) Equals(other *Binding) bool {
	return other.Native() == recv.Native()
}

func (recv *Binding) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'g_binding_get_flags' : return type 'BindingFlags' not supported

var bindingGetSourceFunction *gi.Function
var bindingGetSourceFunction_Once sync.Once

func bindingGetSourceFunction_Set() error {
	var err error
	bindingGetSourceFunction_Once.Do(func() {
		err = bindingObject_Set()
		if err != nil {
			return
		}
		bindingGetSourceFunction, err = bindingObject.InvokerNew("get_source")
	})
	return err
}

// GetSource is a representation of the C type g_binding_get_source.
func (recv *Binding) GetSource() *Object {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := bindingGetSourceFunction_Set()
	if err == nil {
		ret = bindingGetSourceFunction.Invoke(inArgs[:], nil)
	}

	retGo := ObjectNewFromNative(ret.Pointer())

	return retGo
}

var bindingGetSourcePropertyFunction *gi.Function
var bindingGetSourcePropertyFunction_Once sync.Once

func bindingGetSourcePropertyFunction_Set() error {
	var err error
	bindingGetSourcePropertyFunction_Once.Do(func() {
		err = bindingObject_Set()
		if err != nil {
			return
		}
		bindingGetSourcePropertyFunction, err = bindingObject.InvokerNew("get_source_property")
	})
	return err
}

// GetSourceProperty is a representation of the C type g_binding_get_source_property.
func (recv *Binding) GetSourceProperty() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := bindingGetSourcePropertyFunction_Set()
	if err == nil {
		ret = bindingGetSourcePropertyFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var bindingGetTargetFunction *gi.Function
var bindingGetTargetFunction_Once sync.Once

func bindingGetTargetFunction_Set() error {
	var err error
	bindingGetTargetFunction_Once.Do(func() {
		err = bindingObject_Set()
		if err != nil {
			return
		}
		bindingGetTargetFunction, err = bindingObject.InvokerNew("get_target")
	})
	return err
}

// GetTarget is a representation of the C type g_binding_get_target.
func (recv *Binding) GetTarget() *Object {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := bindingGetTargetFunction_Set()
	if err == nil {
		ret = bindingGetTargetFunction.Invoke(inArgs[:], nil)
	}

	retGo := ObjectNewFromNative(ret.Pointer())

	return retGo
}

var bindingGetTargetPropertyFunction *gi.Function
var bindingGetTargetPropertyFunction_Once sync.Once

func bindingGetTargetPropertyFunction_Set() error {
	var err error
	bindingGetTargetPropertyFunction_Once.Do(func() {
		err = bindingObject_Set()
		if err != nil {
			return
		}
		bindingGetTargetPropertyFunction, err = bindingObject.InvokerNew("get_target_property")
	})
	return err
}

// GetTargetProperty is a representation of the C type g_binding_get_target_property.
func (recv *Binding) GetTargetProperty() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := bindingGetTargetPropertyFunction_Set()
	if err == nil {
		ret = bindingGetTargetPropertyFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var bindingUnbindFunction *gi.Function
var bindingUnbindFunction_Once sync.Once

func bindingUnbindFunction_Set() error {
	var err error
	bindingUnbindFunction_Once.Do(func() {
		err = bindingObject_Set()
		if err != nil {
			return
		}
		bindingUnbindFunction, err = bindingObject.InvokerNew("unbind")
	})
	return err
}

// Unbind is a representation of the C type g_binding_unbind.
func (recv *Binding) Unbind() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := bindingUnbindFunction_Set()
	if err == nil {
		bindingUnbindFunction.Invoke(inArgs[:], nil)
	}

	return
}

var initiallyUnownedObject *gi.Object
var initiallyUnownedObject_Once sync.Once

func initiallyUnownedObject_Set() error {
	var err error
	initiallyUnownedObject_Once.Do(func() {
		initiallyUnownedObject, err = gi.ObjectNew("GObject", "InitiallyUnowned")
	})
	return err
}

type InitiallyUnowned struct {
	native unsafe.Pointer
}

func InitiallyUnownedNewFromNative(native unsafe.Pointer) *InitiallyUnowned {
	instance := &InitiallyUnowned{native: native}

	return instance
}

// Object upcasts to *Object
func (recv *InitiallyUnowned) Object() *Object {
	return ObjectNewFromNative(recv.Native())
}

// Equals compares this InitiallyUnowned with another InitiallyUnowned, and returns true if they represent the same Object.
func (recv *InitiallyUnowned) Equals(other *InitiallyUnowned) bool {
	return other.Native() == recv.Native()
}

func (recv *InitiallyUnowned) Native() unsafe.Pointer {
	return recv.native
}

// FieldGTypeInstance returns the C field 'g_type_instance'.
func (recv *InitiallyUnowned) FieldGTypeInstance() *TypeInstance {
	argValue := gi.ObjectFieldGet(initiallyUnownedObject, recv.Native(), "g_type_instance")
	value := TypeInstanceNewFromNative(argValue.Pointer())
	return value
}

// SetFieldGTypeInstance sets the value of the C field 'g_type_instance'.
func (recv *InitiallyUnowned) SetFieldGTypeInstance(value *TypeInstance) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(initiallyUnownedObject, recv.Native(), "g_type_instance", argValue)
}

var objectObject *gi.Object
var objectObject_Once sync.Once

func objectObject_Set() error {
	var err error
	objectObject_Once.Do(func() {
		objectObject, err = gi.ObjectNew("GObject", "Object")
	})
	return err
}

type Object struct {
	native unsafe.Pointer
}

func ObjectNewFromNative(native unsafe.Pointer) *Object {
	instance := &Object{native: native}

	return instance
}

// Equals compares this Object with another Object, and returns true if they represent the same Object.
func (recv *Object) Equals(other *Object) bool {
	return other.Native() == recv.Native()
}

func (recv *Object) Native() unsafe.Pointer {
	return recv.native
}

// FieldGTypeInstance returns the C field 'g_type_instance'.
func (recv *Object) FieldGTypeInstance() *TypeInstance {
	argValue := gi.ObjectFieldGet(objectObject, recv.Native(), "g_type_instance")
	value := TypeInstanceNewFromNative(argValue.Pointer())
	return value
}

// SetFieldGTypeInstance sets the value of the C field 'g_type_instance'.
func (recv *Object) SetFieldGTypeInstance(value *TypeInstance) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(objectObject, recv.Native(), "g_type_instance", argValue)
}

// UNSUPPORTED : C value 'g_object_new' : parameter 'object_type' of type 'GType' not supported

// UNSUPPORTED : C value 'g_object_new_valist' : parameter 'object_type' of type 'GType' not supported

// UNSUPPORTED : C value 'g_object_new_with_properties' : parameter 'object_type' of type 'GType' not supported

// UNSUPPORTED : C value 'g_object_newv' : parameter 'object_type' of type 'GType' not supported

// UNSUPPORTED : C value 'g_object_add_toggle_ref' : parameter 'notify' of type 'ToggleNotify' not supported

// UNSUPPORTED : C value 'g_object_add_weak_pointer' : parameter 'weak_pointer_location' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_object_bind_property' : parameter 'flags' of type 'BindingFlags' not supported

// UNSUPPORTED : C value 'g_object_bind_property_full' : parameter 'flags' of type 'BindingFlags' not supported

// UNSUPPORTED : C value 'g_object_bind_property_with_closures' : parameter 'flags' of type 'BindingFlags' not supported

// UNSUPPORTED : C value 'g_object_connect' : parameter '...' of type 'nil' not supported

// UNSUPPORTED : C value 'g_object_disconnect' : parameter '...' of type 'nil' not supported

// UNSUPPORTED : C value 'g_object_dup_data' : parameter 'dup_func' of type 'GLib.DuplicateFunc' not supported

// UNSUPPORTED : C value 'g_object_dup_qdata' : parameter 'dup_func' of type 'GLib.DuplicateFunc' not supported

var objectForceFloatingFunction *gi.Function
var objectForceFloatingFunction_Once sync.Once

func objectForceFloatingFunction_Set() error {
	var err error
	objectForceFloatingFunction_Once.Do(func() {
		err = objectObject_Set()
		if err != nil {
			return
		}
		objectForceFloatingFunction, err = objectObject.InvokerNew("force_floating")
	})
	return err
}

// ForceFloating is a representation of the C type g_object_force_floating.
func (recv *Object) ForceFloating() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := objectForceFloatingFunction_Set()
	if err == nil {
		objectForceFloatingFunction.Invoke(inArgs[:], nil)
	}

	return
}

var objectFreezeNotifyFunction *gi.Function
var objectFreezeNotifyFunction_Once sync.Once

func objectFreezeNotifyFunction_Set() error {
	var err error
	objectFreezeNotifyFunction_Once.Do(func() {
		err = objectObject_Set()
		if err != nil {
			return
		}
		objectFreezeNotifyFunction, err = objectObject.InvokerNew("freeze_notify")
	})
	return err
}

// FreezeNotify is a representation of the C type g_object_freeze_notify.
func (recv *Object) FreezeNotify() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := objectFreezeNotifyFunction_Set()
	if err == nil {
		objectFreezeNotifyFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'g_object_get' : parameter '...' of type 'nil' not supported

// UNSUPPORTED : C value 'g_object_get_data' : return type 'gpointer' not supported

var objectGetPropertyFunction *gi.Function
var objectGetPropertyFunction_Once sync.Once

func objectGetPropertyFunction_Set() error {
	var err error
	objectGetPropertyFunction_Once.Do(func() {
		err = objectObject_Set()
		if err != nil {
			return
		}
		objectGetPropertyFunction, err = objectObject.InvokerNew("get_property")
	})
	return err
}

// GetProperty is a representation of the C type g_object_get_property.
func (recv *Object) GetProperty(propertyName string, value *Value) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(propertyName)
	inArgs[2].SetPointer(value.Native())

	err := objectGetPropertyFunction_Set()
	if err == nil {
		objectGetPropertyFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'g_object_get_qdata' : return type 'gpointer' not supported

// UNSUPPORTED : C value 'g_object_get_valist' : parameter 'var_args' of type 'va_list' not supported

// UNSUPPORTED : C value 'g_object_getv' : parameter 'names' of type 'nil' not supported

var objectIsFloatingFunction *gi.Function
var objectIsFloatingFunction_Once sync.Once

func objectIsFloatingFunction_Set() error {
	var err error
	objectIsFloatingFunction_Once.Do(func() {
		err = objectObject_Set()
		if err != nil {
			return
		}
		objectIsFloatingFunction, err = objectObject.InvokerNew("is_floating")
	})
	return err
}

// IsFloating is a representation of the C type g_object_is_floating.
func (recv *Object) IsFloating() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := objectIsFloatingFunction_Set()
	if err == nil {
		ret = objectIsFloatingFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var objectNotifyFunction *gi.Function
var objectNotifyFunction_Once sync.Once

func objectNotifyFunction_Set() error {
	var err error
	objectNotifyFunction_Once.Do(func() {
		err = objectObject_Set()
		if err != nil {
			return
		}
		objectNotifyFunction, err = objectObject.InvokerNew("notify")
	})
	return err
}

// Notify is a representation of the C type g_object_notify.
func (recv *Object) Notify(propertyName string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(propertyName)

	err := objectNotifyFunction_Set()
	if err == nil {
		objectNotifyFunction.Invoke(inArgs[:], nil)
	}

	return
}

var objectNotifyByPspecFunction *gi.Function
var objectNotifyByPspecFunction_Once sync.Once

func objectNotifyByPspecFunction_Set() error {
	var err error
	objectNotifyByPspecFunction_Once.Do(func() {
		err = objectObject_Set()
		if err != nil {
			return
		}
		objectNotifyByPspecFunction, err = objectObject.InvokerNew("notify_by_pspec")
	})
	return err
}

// NotifyByPspec is a representation of the C type g_object_notify_by_pspec.
func (recv *Object) NotifyByPspec(pspec *ParamSpec) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(pspec.Native())

	err := objectNotifyByPspecFunction_Set()
	if err == nil {
		objectNotifyByPspecFunction.Invoke(inArgs[:], nil)
	}

	return
}

var objectRefFunction *gi.Function
var objectRefFunction_Once sync.Once

func objectRefFunction_Set() error {
	var err error
	objectRefFunction_Once.Do(func() {
		err = objectObject_Set()
		if err != nil {
			return
		}
		objectRefFunction, err = objectObject.InvokerNew("ref")
	})
	return err
}

// Ref is a representation of the C type g_object_ref.
func (recv *Object) Ref() *Object {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := objectRefFunction_Set()
	if err == nil {
		ret = objectRefFunction.Invoke(inArgs[:], nil)
	}

	retGo := ObjectNewFromNative(ret.Pointer())

	return retGo
}

var objectRefSinkFunction *gi.Function
var objectRefSinkFunction_Once sync.Once

func objectRefSinkFunction_Set() error {
	var err error
	objectRefSinkFunction_Once.Do(func() {
		err = objectObject_Set()
		if err != nil {
			return
		}
		objectRefSinkFunction, err = objectObject.InvokerNew("ref_sink")
	})
	return err
}

// RefSink is a representation of the C type g_object_ref_sink.
func (recv *Object) RefSink() *Object {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := objectRefSinkFunction_Set()
	if err == nil {
		ret = objectRefSinkFunction.Invoke(inArgs[:], nil)
	}

	retGo := ObjectNewFromNative(ret.Pointer())

	return retGo
}

// UNSUPPORTED : C value 'g_object_remove_toggle_ref' : parameter 'notify' of type 'ToggleNotify' not supported

// UNSUPPORTED : C value 'g_object_remove_weak_pointer' : parameter 'weak_pointer_location' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_object_replace_data' : parameter 'oldval' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_object_replace_qdata' : parameter 'oldval' of type 'gpointer' not supported

var objectRunDisposeFunction *gi.Function
var objectRunDisposeFunction_Once sync.Once

func objectRunDisposeFunction_Set() error {
	var err error
	objectRunDisposeFunction_Once.Do(func() {
		err = objectObject_Set()
		if err != nil {
			return
		}
		objectRunDisposeFunction, err = objectObject.InvokerNew("run_dispose")
	})
	return err
}

// RunDispose is a representation of the C type g_object_run_dispose.
func (recv *Object) RunDispose() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := objectRunDisposeFunction_Set()
	if err == nil {
		objectRunDisposeFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'g_object_set' : parameter '...' of type 'nil' not supported

// UNSUPPORTED : C value 'g_object_set_data' : parameter 'data' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_object_set_data_full' : parameter 'data' of type 'gpointer' not supported

var objectSetPropertyFunction *gi.Function
var objectSetPropertyFunction_Once sync.Once

func objectSetPropertyFunction_Set() error {
	var err error
	objectSetPropertyFunction_Once.Do(func() {
		err = objectObject_Set()
		if err != nil {
			return
		}
		objectSetPropertyFunction, err = objectObject.InvokerNew("set_property")
	})
	return err
}

// SetProperty is a representation of the C type g_object_set_property.
func (recv *Object) SetProperty(propertyName string, value *Value) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(propertyName)
	inArgs[2].SetPointer(value.Native())

	err := objectSetPropertyFunction_Set()
	if err == nil {
		objectSetPropertyFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'g_object_set_qdata' : parameter 'data' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_object_set_qdata_full' : parameter 'data' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_object_set_valist' : parameter 'var_args' of type 'va_list' not supported

// UNSUPPORTED : C value 'g_object_setv' : parameter 'names' of type 'nil' not supported

// UNSUPPORTED : C value 'g_object_steal_data' : return type 'gpointer' not supported

// UNSUPPORTED : C value 'g_object_steal_qdata' : return type 'gpointer' not supported

var objectThawNotifyFunction *gi.Function
var objectThawNotifyFunction_Once sync.Once

func objectThawNotifyFunction_Set() error {
	var err error
	objectThawNotifyFunction_Once.Do(func() {
		err = objectObject_Set()
		if err != nil {
			return
		}
		objectThawNotifyFunction, err = objectObject.InvokerNew("thaw_notify")
	})
	return err
}

// ThawNotify is a representation of the C type g_object_thaw_notify.
func (recv *Object) ThawNotify() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := objectThawNotifyFunction_Set()
	if err == nil {
		objectThawNotifyFunction.Invoke(inArgs[:], nil)
	}

	return
}

var objectUnrefFunction *gi.Function
var objectUnrefFunction_Once sync.Once

func objectUnrefFunction_Set() error {
	var err error
	objectUnrefFunction_Once.Do(func() {
		err = objectObject_Set()
		if err != nil {
			return
		}
		objectUnrefFunction, err = objectObject.InvokerNew("unref")
	})
	return err
}

// Unref is a representation of the C type g_object_unref.
func (recv *Object) Unref() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := objectUnrefFunction_Set()
	if err == nil {
		objectUnrefFunction.Invoke(inArgs[:], nil)
	}

	return
}

var objectWatchClosureFunction *gi.Function
var objectWatchClosureFunction_Once sync.Once

func objectWatchClosureFunction_Set() error {
	var err error
	objectWatchClosureFunction_Once.Do(func() {
		err = objectObject_Set()
		if err != nil {
			return
		}
		objectWatchClosureFunction, err = objectObject.InvokerNew("watch_closure")
	})
	return err
}

// WatchClosure is a representation of the C type g_object_watch_closure.
func (recv *Object) WatchClosure(closure *Closure) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(closure.Native())

	err := objectWatchClosureFunction_Set()
	if err == nil {
		objectWatchClosureFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'g_object_weak_ref' : parameter 'notify' of type 'WeakNotify' not supported

// UNSUPPORTED : C value 'g_object_weak_unref' : parameter 'notify' of type 'WeakNotify' not supported

/*
Disconnect disconnects a callback previously registered with a Connect...() method.

The connectionID should be a value returned from a call to a Connect...() method.
*/
func (recv *Object) DisconnectSignal(connectionID int) {
	callback.DisconnectSignal(connectionID)
}

var paramSpecObject *gi.Object
var paramSpecObject_Once sync.Once

func paramSpecObject_Set() error {
	var err error
	paramSpecObject_Once.Do(func() {
		paramSpecObject, err = gi.ObjectNew("GObject", "ParamSpec")
	})
	return err
}

type ParamSpec struct {
	native unsafe.Pointer
}

func ParamSpecNewFromNative(native unsafe.Pointer) *ParamSpec {
	instance := &ParamSpec{native: native}

	return instance
}

// Equals compares this ParamSpec with another ParamSpec, and returns true if they represent the same Object.
func (recv *ParamSpec) Equals(other *ParamSpec) bool {
	return other.Native() == recv.Native()
}

func (recv *ParamSpec) Native() unsafe.Pointer {
	return recv.native
}

// FieldGTypeInstance returns the C field 'g_type_instance'.
func (recv *ParamSpec) FieldGTypeInstance() *TypeInstance {
	argValue := gi.ObjectFieldGet(paramSpecObject, recv.Native(), "g_type_instance")
	value := TypeInstanceNewFromNative(argValue.Pointer())
	return value
}

// SetFieldGTypeInstance sets the value of the C field 'g_type_instance'.
func (recv *ParamSpec) SetFieldGTypeInstance(value *TypeInstance) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(paramSpecObject, recv.Native(), "g_type_instance", argValue)
}

// FieldName returns the C field 'name'.
func (recv *ParamSpec) FieldName() string {
	argValue := gi.ObjectFieldGet(paramSpecObject, recv.Native(), "name")
	value := argValue.String(false)
	return value
}

// SetFieldName sets the value of the C field 'name'.
func (recv *ParamSpec) SetFieldName(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.ObjectFieldSet(paramSpecObject, recv.Native(), "name", argValue)
}

// UNSUPPORTED : C value 'flags' : for field getter : no Go type for 'ParamFlags'

// UNSUPPORTED : C value 'flags' : for field setter : no Go type for 'ParamFlags'

// UNSUPPORTED : C value 'value_type' : for field getter : no Go type for 'GType'

// UNSUPPORTED : C value 'value_type' : for field setter : no Go type for 'GType'

// UNSUPPORTED : C value 'owner_type' : for field getter : no Go type for 'GType'

// UNSUPPORTED : C value 'owner_type' : for field setter : no Go type for 'GType'

var paramSpecGetBlurbFunction *gi.Function
var paramSpecGetBlurbFunction_Once sync.Once

func paramSpecGetBlurbFunction_Set() error {
	var err error
	paramSpecGetBlurbFunction_Once.Do(func() {
		err = paramSpecObject_Set()
		if err != nil {
			return
		}
		paramSpecGetBlurbFunction, err = paramSpecObject.InvokerNew("get_blurb")
	})
	return err
}

// GetBlurb is a representation of the C type g_param_spec_get_blurb.
func (recv *ParamSpec) GetBlurb() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := paramSpecGetBlurbFunction_Set()
	if err == nil {
		ret = paramSpecGetBlurbFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var paramSpecGetDefaultValueFunction *gi.Function
var paramSpecGetDefaultValueFunction_Once sync.Once

func paramSpecGetDefaultValueFunction_Set() error {
	var err error
	paramSpecGetDefaultValueFunction_Once.Do(func() {
		err = paramSpecObject_Set()
		if err != nil {
			return
		}
		paramSpecGetDefaultValueFunction, err = paramSpecObject.InvokerNew("get_default_value")
	})
	return err
}

// GetDefaultValue is a representation of the C type g_param_spec_get_default_value.
func (recv *ParamSpec) GetDefaultValue() *Value {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := paramSpecGetDefaultValueFunction_Set()
	if err == nil {
		ret = paramSpecGetDefaultValueFunction.Invoke(inArgs[:], nil)
	}

	retGo := ValueNewFromNative(ret.Pointer())

	return retGo
}

var paramSpecGetNameFunction *gi.Function
var paramSpecGetNameFunction_Once sync.Once

func paramSpecGetNameFunction_Set() error {
	var err error
	paramSpecGetNameFunction_Once.Do(func() {
		err = paramSpecObject_Set()
		if err != nil {
			return
		}
		paramSpecGetNameFunction, err = paramSpecObject.InvokerNew("get_name")
	})
	return err
}

// GetName is a representation of the C type g_param_spec_get_name.
func (recv *ParamSpec) GetName() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := paramSpecGetNameFunction_Set()
	if err == nil {
		ret = paramSpecGetNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var paramSpecGetNameQuarkFunction *gi.Function
var paramSpecGetNameQuarkFunction_Once sync.Once

func paramSpecGetNameQuarkFunction_Set() error {
	var err error
	paramSpecGetNameQuarkFunction_Once.Do(func() {
		err = paramSpecObject_Set()
		if err != nil {
			return
		}
		paramSpecGetNameQuarkFunction, err = paramSpecObject.InvokerNew("get_name_quark")
	})
	return err
}

// GetNameQuark is a representation of the C type g_param_spec_get_name_quark.
func (recv *ParamSpec) GetNameQuark() glib.Quark {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := paramSpecGetNameQuarkFunction_Set()
	if err == nil {
		ret = paramSpecGetNameQuarkFunction.Invoke(inArgs[:], nil)
	}

	retGo := glib.Quark(ret.Uint32())

	return retGo
}

var paramSpecGetNickFunction *gi.Function
var paramSpecGetNickFunction_Once sync.Once

func paramSpecGetNickFunction_Set() error {
	var err error
	paramSpecGetNickFunction_Once.Do(func() {
		err = paramSpecObject_Set()
		if err != nil {
			return
		}
		paramSpecGetNickFunction, err = paramSpecObject.InvokerNew("get_nick")
	})
	return err
}

// GetNick is a representation of the C type g_param_spec_get_nick.
func (recv *ParamSpec) GetNick() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := paramSpecGetNickFunction_Set()
	if err == nil {
		ret = paramSpecGetNickFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'g_param_spec_get_qdata' : return type 'gpointer' not supported

var paramSpecGetRedirectTargetFunction *gi.Function
var paramSpecGetRedirectTargetFunction_Once sync.Once

func paramSpecGetRedirectTargetFunction_Set() error {
	var err error
	paramSpecGetRedirectTargetFunction_Once.Do(func() {
		err = paramSpecObject_Set()
		if err != nil {
			return
		}
		paramSpecGetRedirectTargetFunction, err = paramSpecObject.InvokerNew("get_redirect_target")
	})
	return err
}

// GetRedirectTarget is a representation of the C type g_param_spec_get_redirect_target.
func (recv *ParamSpec) GetRedirectTarget() *ParamSpec {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := paramSpecGetRedirectTargetFunction_Set()
	if err == nil {
		ret = paramSpecGetRedirectTargetFunction.Invoke(inArgs[:], nil)
	}

	retGo := ParamSpecNewFromNative(ret.Pointer())

	return retGo
}

var paramSpecRefFunction *gi.Function
var paramSpecRefFunction_Once sync.Once

func paramSpecRefFunction_Set() error {
	var err error
	paramSpecRefFunction_Once.Do(func() {
		err = paramSpecObject_Set()
		if err != nil {
			return
		}
		paramSpecRefFunction, err = paramSpecObject.InvokerNew("ref")
	})
	return err
}

// Ref is a representation of the C type g_param_spec_ref.
func (recv *ParamSpec) Ref() *ParamSpec {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := paramSpecRefFunction_Set()
	if err == nil {
		ret = paramSpecRefFunction.Invoke(inArgs[:], nil)
	}

	retGo := ParamSpecNewFromNative(ret.Pointer())

	return retGo
}

var paramSpecRefSinkFunction *gi.Function
var paramSpecRefSinkFunction_Once sync.Once

func paramSpecRefSinkFunction_Set() error {
	var err error
	paramSpecRefSinkFunction_Once.Do(func() {
		err = paramSpecObject_Set()
		if err != nil {
			return
		}
		paramSpecRefSinkFunction, err = paramSpecObject.InvokerNew("ref_sink")
	})
	return err
}

// RefSink is a representation of the C type g_param_spec_ref_sink.
func (recv *ParamSpec) RefSink() *ParamSpec {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := paramSpecRefSinkFunction_Set()
	if err == nil {
		ret = paramSpecRefSinkFunction.Invoke(inArgs[:], nil)
	}

	retGo := ParamSpecNewFromNative(ret.Pointer())

	return retGo
}

// UNSUPPORTED : C value 'g_param_spec_set_qdata' : parameter 'data' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_param_spec_set_qdata_full' : parameter 'data' of type 'gpointer' not supported

var paramSpecSinkFunction *gi.Function
var paramSpecSinkFunction_Once sync.Once

func paramSpecSinkFunction_Set() error {
	var err error
	paramSpecSinkFunction_Once.Do(func() {
		err = paramSpecObject_Set()
		if err != nil {
			return
		}
		paramSpecSinkFunction, err = paramSpecObject.InvokerNew("sink")
	})
	return err
}

// Sink is a representation of the C type g_param_spec_sink.
func (recv *ParamSpec) Sink() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := paramSpecSinkFunction_Set()
	if err == nil {
		paramSpecSinkFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'g_param_spec_steal_qdata' : return type 'gpointer' not supported

var paramSpecUnrefFunction *gi.Function
var paramSpecUnrefFunction_Once sync.Once

func paramSpecUnrefFunction_Set() error {
	var err error
	paramSpecUnrefFunction_Once.Do(func() {
		err = paramSpecObject_Set()
		if err != nil {
			return
		}
		paramSpecUnrefFunction, err = paramSpecObject.InvokerNew("unref")
	})
	return err
}

// Unref is a representation of the C type g_param_spec_unref.
func (recv *ParamSpec) Unref() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := paramSpecUnrefFunction_Set()
	if err == nil {
		paramSpecUnrefFunction.Invoke(inArgs[:], nil)
	}

	return
}

var paramSpecBooleanObject *gi.Object
var paramSpecBooleanObject_Once sync.Once

func paramSpecBooleanObject_Set() error {
	var err error
	paramSpecBooleanObject_Once.Do(func() {
		paramSpecBooleanObject, err = gi.ObjectNew("GObject", "ParamSpecBoolean")
	})
	return err
}

type ParamSpecBoolean struct {
	native unsafe.Pointer
}

func ParamSpecBooleanNewFromNative(native unsafe.Pointer) *ParamSpecBoolean {
	instance := &ParamSpecBoolean{native: native}

	return instance
}

// ParamSpec upcasts to *ParamSpec
func (recv *ParamSpecBoolean) ParamSpec() *ParamSpec {
	return ParamSpecNewFromNative(recv.Native())
}

// Equals compares this ParamSpecBoolean with another ParamSpecBoolean, and returns true if they represent the same Object.
func (recv *ParamSpecBoolean) Equals(other *ParamSpecBoolean) bool {
	return other.Native() == recv.Native()
}

func (recv *ParamSpecBoolean) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *ParamSpecBoolean) FieldParentInstance() *ParamSpec {
	argValue := gi.ObjectFieldGet(paramSpecBooleanObject, recv.Native(), "parent_instance")
	value := ParamSpecNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *ParamSpecBoolean) SetFieldParentInstance(value *ParamSpec) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(paramSpecBooleanObject, recv.Native(), "parent_instance", argValue)
}

// FieldDefaultValue returns the C field 'default_value'.
func (recv *ParamSpecBoolean) FieldDefaultValue() bool {
	argValue := gi.ObjectFieldGet(paramSpecBooleanObject, recv.Native(), "default_value")
	value := argValue.Boolean()
	return value
}

// SetFieldDefaultValue sets the value of the C field 'default_value'.
func (recv *ParamSpecBoolean) SetFieldDefaultValue(value bool) {
	var argValue gi.Argument
	argValue.SetBoolean(value)
	gi.ObjectFieldSet(paramSpecBooleanObject, recv.Native(), "default_value", argValue)
}

var paramSpecBoxedObject *gi.Object
var paramSpecBoxedObject_Once sync.Once

func paramSpecBoxedObject_Set() error {
	var err error
	paramSpecBoxedObject_Once.Do(func() {
		paramSpecBoxedObject, err = gi.ObjectNew("GObject", "ParamSpecBoxed")
	})
	return err
}

type ParamSpecBoxed struct {
	native unsafe.Pointer
}

func ParamSpecBoxedNewFromNative(native unsafe.Pointer) *ParamSpecBoxed {
	instance := &ParamSpecBoxed{native: native}

	return instance
}

// ParamSpec upcasts to *ParamSpec
func (recv *ParamSpecBoxed) ParamSpec() *ParamSpec {
	return ParamSpecNewFromNative(recv.Native())
}

// Equals compares this ParamSpecBoxed with another ParamSpecBoxed, and returns true if they represent the same Object.
func (recv *ParamSpecBoxed) Equals(other *ParamSpecBoxed) bool {
	return other.Native() == recv.Native()
}

func (recv *ParamSpecBoxed) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *ParamSpecBoxed) FieldParentInstance() *ParamSpec {
	argValue := gi.ObjectFieldGet(paramSpecBoxedObject, recv.Native(), "parent_instance")
	value := ParamSpecNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *ParamSpecBoxed) SetFieldParentInstance(value *ParamSpec) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(paramSpecBoxedObject, recv.Native(), "parent_instance", argValue)
}

var paramSpecCharObject *gi.Object
var paramSpecCharObject_Once sync.Once

func paramSpecCharObject_Set() error {
	var err error
	paramSpecCharObject_Once.Do(func() {
		paramSpecCharObject, err = gi.ObjectNew("GObject", "ParamSpecChar")
	})
	return err
}

type ParamSpecChar struct {
	native unsafe.Pointer
}

func ParamSpecCharNewFromNative(native unsafe.Pointer) *ParamSpecChar {
	instance := &ParamSpecChar{native: native}

	return instance
}

// ParamSpec upcasts to *ParamSpec
func (recv *ParamSpecChar) ParamSpec() *ParamSpec {
	return ParamSpecNewFromNative(recv.Native())
}

// Equals compares this ParamSpecChar with another ParamSpecChar, and returns true if they represent the same Object.
func (recv *ParamSpecChar) Equals(other *ParamSpecChar) bool {
	return other.Native() == recv.Native()
}

func (recv *ParamSpecChar) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *ParamSpecChar) FieldParentInstance() *ParamSpec {
	argValue := gi.ObjectFieldGet(paramSpecCharObject, recv.Native(), "parent_instance")
	value := ParamSpecNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *ParamSpecChar) SetFieldParentInstance(value *ParamSpec) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(paramSpecCharObject, recv.Native(), "parent_instance", argValue)
}

// FieldMinimum returns the C field 'minimum'.
func (recv *ParamSpecChar) FieldMinimum() int8 {
	argValue := gi.ObjectFieldGet(paramSpecCharObject, recv.Native(), "minimum")
	value := argValue.Int8()
	return value
}

// SetFieldMinimum sets the value of the C field 'minimum'.
func (recv *ParamSpecChar) SetFieldMinimum(value int8) {
	var argValue gi.Argument
	argValue.SetInt8(value)
	gi.ObjectFieldSet(paramSpecCharObject, recv.Native(), "minimum", argValue)
}

// FieldMaximum returns the C field 'maximum'.
func (recv *ParamSpecChar) FieldMaximum() int8 {
	argValue := gi.ObjectFieldGet(paramSpecCharObject, recv.Native(), "maximum")
	value := argValue.Int8()
	return value
}

// SetFieldMaximum sets the value of the C field 'maximum'.
func (recv *ParamSpecChar) SetFieldMaximum(value int8) {
	var argValue gi.Argument
	argValue.SetInt8(value)
	gi.ObjectFieldSet(paramSpecCharObject, recv.Native(), "maximum", argValue)
}

// FieldDefaultValue returns the C field 'default_value'.
func (recv *ParamSpecChar) FieldDefaultValue() int8 {
	argValue := gi.ObjectFieldGet(paramSpecCharObject, recv.Native(), "default_value")
	value := argValue.Int8()
	return value
}

// SetFieldDefaultValue sets the value of the C field 'default_value'.
func (recv *ParamSpecChar) SetFieldDefaultValue(value int8) {
	var argValue gi.Argument
	argValue.SetInt8(value)
	gi.ObjectFieldSet(paramSpecCharObject, recv.Native(), "default_value", argValue)
}

var paramSpecDoubleObject *gi.Object
var paramSpecDoubleObject_Once sync.Once

func paramSpecDoubleObject_Set() error {
	var err error
	paramSpecDoubleObject_Once.Do(func() {
		paramSpecDoubleObject, err = gi.ObjectNew("GObject", "ParamSpecDouble")
	})
	return err
}

type ParamSpecDouble struct {
	native unsafe.Pointer
}

func ParamSpecDoubleNewFromNative(native unsafe.Pointer) *ParamSpecDouble {
	instance := &ParamSpecDouble{native: native}

	return instance
}

// ParamSpec upcasts to *ParamSpec
func (recv *ParamSpecDouble) ParamSpec() *ParamSpec {
	return ParamSpecNewFromNative(recv.Native())
}

// Equals compares this ParamSpecDouble with another ParamSpecDouble, and returns true if they represent the same Object.
func (recv *ParamSpecDouble) Equals(other *ParamSpecDouble) bool {
	return other.Native() == recv.Native()
}

func (recv *ParamSpecDouble) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *ParamSpecDouble) FieldParentInstance() *ParamSpec {
	argValue := gi.ObjectFieldGet(paramSpecDoubleObject, recv.Native(), "parent_instance")
	value := ParamSpecNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *ParamSpecDouble) SetFieldParentInstance(value *ParamSpec) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(paramSpecDoubleObject, recv.Native(), "parent_instance", argValue)
}

// FieldMinimum returns the C field 'minimum'.
func (recv *ParamSpecDouble) FieldMinimum() float64 {
	argValue := gi.ObjectFieldGet(paramSpecDoubleObject, recv.Native(), "minimum")
	value := argValue.Float64()
	return value
}

// SetFieldMinimum sets the value of the C field 'minimum'.
func (recv *ParamSpecDouble) SetFieldMinimum(value float64) {
	var argValue gi.Argument
	argValue.SetFloat64(value)
	gi.ObjectFieldSet(paramSpecDoubleObject, recv.Native(), "minimum", argValue)
}

// FieldMaximum returns the C field 'maximum'.
func (recv *ParamSpecDouble) FieldMaximum() float64 {
	argValue := gi.ObjectFieldGet(paramSpecDoubleObject, recv.Native(), "maximum")
	value := argValue.Float64()
	return value
}

// SetFieldMaximum sets the value of the C field 'maximum'.
func (recv *ParamSpecDouble) SetFieldMaximum(value float64) {
	var argValue gi.Argument
	argValue.SetFloat64(value)
	gi.ObjectFieldSet(paramSpecDoubleObject, recv.Native(), "maximum", argValue)
}

// FieldDefaultValue returns the C field 'default_value'.
func (recv *ParamSpecDouble) FieldDefaultValue() float64 {
	argValue := gi.ObjectFieldGet(paramSpecDoubleObject, recv.Native(), "default_value")
	value := argValue.Float64()
	return value
}

// SetFieldDefaultValue sets the value of the C field 'default_value'.
func (recv *ParamSpecDouble) SetFieldDefaultValue(value float64) {
	var argValue gi.Argument
	argValue.SetFloat64(value)
	gi.ObjectFieldSet(paramSpecDoubleObject, recv.Native(), "default_value", argValue)
}

// FieldEpsilon returns the C field 'epsilon'.
func (recv *ParamSpecDouble) FieldEpsilon() float64 {
	argValue := gi.ObjectFieldGet(paramSpecDoubleObject, recv.Native(), "epsilon")
	value := argValue.Float64()
	return value
}

// SetFieldEpsilon sets the value of the C field 'epsilon'.
func (recv *ParamSpecDouble) SetFieldEpsilon(value float64) {
	var argValue gi.Argument
	argValue.SetFloat64(value)
	gi.ObjectFieldSet(paramSpecDoubleObject, recv.Native(), "epsilon", argValue)
}

var paramSpecEnumObject *gi.Object
var paramSpecEnumObject_Once sync.Once

func paramSpecEnumObject_Set() error {
	var err error
	paramSpecEnumObject_Once.Do(func() {
		paramSpecEnumObject, err = gi.ObjectNew("GObject", "ParamSpecEnum")
	})
	return err
}

type ParamSpecEnum struct {
	native unsafe.Pointer
}

func ParamSpecEnumNewFromNative(native unsafe.Pointer) *ParamSpecEnum {
	instance := &ParamSpecEnum{native: native}

	return instance
}

// ParamSpec upcasts to *ParamSpec
func (recv *ParamSpecEnum) ParamSpec() *ParamSpec {
	return ParamSpecNewFromNative(recv.Native())
}

// Equals compares this ParamSpecEnum with another ParamSpecEnum, and returns true if they represent the same Object.
func (recv *ParamSpecEnum) Equals(other *ParamSpecEnum) bool {
	return other.Native() == recv.Native()
}

func (recv *ParamSpecEnum) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *ParamSpecEnum) FieldParentInstance() *ParamSpec {
	argValue := gi.ObjectFieldGet(paramSpecEnumObject, recv.Native(), "parent_instance")
	value := ParamSpecNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *ParamSpecEnum) SetFieldParentInstance(value *ParamSpec) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(paramSpecEnumObject, recv.Native(), "parent_instance", argValue)
}

// FieldEnumClass returns the C field 'enum_class'.
func (recv *ParamSpecEnum) FieldEnumClass() *EnumClass {
	argValue := gi.ObjectFieldGet(paramSpecEnumObject, recv.Native(), "enum_class")
	value := EnumClassNewFromNative(argValue.Pointer())
	return value
}

// SetFieldEnumClass sets the value of the C field 'enum_class'.
func (recv *ParamSpecEnum) SetFieldEnumClass(value *EnumClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(paramSpecEnumObject, recv.Native(), "enum_class", argValue)
}

// FieldDefaultValue returns the C field 'default_value'.
func (recv *ParamSpecEnum) FieldDefaultValue() int32 {
	argValue := gi.ObjectFieldGet(paramSpecEnumObject, recv.Native(), "default_value")
	value := argValue.Int32()
	return value
}

// SetFieldDefaultValue sets the value of the C field 'default_value'.
func (recv *ParamSpecEnum) SetFieldDefaultValue(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.ObjectFieldSet(paramSpecEnumObject, recv.Native(), "default_value", argValue)
}

var paramSpecFlagsObject *gi.Object
var paramSpecFlagsObject_Once sync.Once

func paramSpecFlagsObject_Set() error {
	var err error
	paramSpecFlagsObject_Once.Do(func() {
		paramSpecFlagsObject, err = gi.ObjectNew("GObject", "ParamSpecFlags")
	})
	return err
}

type ParamSpecFlags struct {
	native unsafe.Pointer
}

func ParamSpecFlagsNewFromNative(native unsafe.Pointer) *ParamSpecFlags {
	instance := &ParamSpecFlags{native: native}

	return instance
}

// ParamSpec upcasts to *ParamSpec
func (recv *ParamSpecFlags) ParamSpec() *ParamSpec {
	return ParamSpecNewFromNative(recv.Native())
}

// Equals compares this ParamSpecFlags with another ParamSpecFlags, and returns true if they represent the same Object.
func (recv *ParamSpecFlags) Equals(other *ParamSpecFlags) bool {
	return other.Native() == recv.Native()
}

func (recv *ParamSpecFlags) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *ParamSpecFlags) FieldParentInstance() *ParamSpec {
	argValue := gi.ObjectFieldGet(paramSpecFlagsObject, recv.Native(), "parent_instance")
	value := ParamSpecNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *ParamSpecFlags) SetFieldParentInstance(value *ParamSpec) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(paramSpecFlagsObject, recv.Native(), "parent_instance", argValue)
}

// FieldFlagsClass returns the C field 'flags_class'.
func (recv *ParamSpecFlags) FieldFlagsClass() *FlagsClass {
	argValue := gi.ObjectFieldGet(paramSpecFlagsObject, recv.Native(), "flags_class")
	value := FlagsClassNewFromNative(argValue.Pointer())
	return value
}

// SetFieldFlagsClass sets the value of the C field 'flags_class'.
func (recv *ParamSpecFlags) SetFieldFlagsClass(value *FlagsClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(paramSpecFlagsObject, recv.Native(), "flags_class", argValue)
}

// FieldDefaultValue returns the C field 'default_value'.
func (recv *ParamSpecFlags) FieldDefaultValue() uint32 {
	argValue := gi.ObjectFieldGet(paramSpecFlagsObject, recv.Native(), "default_value")
	value := argValue.Uint32()
	return value
}

// SetFieldDefaultValue sets the value of the C field 'default_value'.
func (recv *ParamSpecFlags) SetFieldDefaultValue(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.ObjectFieldSet(paramSpecFlagsObject, recv.Native(), "default_value", argValue)
}

var paramSpecFloatObject *gi.Object
var paramSpecFloatObject_Once sync.Once

func paramSpecFloatObject_Set() error {
	var err error
	paramSpecFloatObject_Once.Do(func() {
		paramSpecFloatObject, err = gi.ObjectNew("GObject", "ParamSpecFloat")
	})
	return err
}

type ParamSpecFloat struct {
	native unsafe.Pointer
}

func ParamSpecFloatNewFromNative(native unsafe.Pointer) *ParamSpecFloat {
	instance := &ParamSpecFloat{native: native}

	return instance
}

// ParamSpec upcasts to *ParamSpec
func (recv *ParamSpecFloat) ParamSpec() *ParamSpec {
	return ParamSpecNewFromNative(recv.Native())
}

// Equals compares this ParamSpecFloat with another ParamSpecFloat, and returns true if they represent the same Object.
func (recv *ParamSpecFloat) Equals(other *ParamSpecFloat) bool {
	return other.Native() == recv.Native()
}

func (recv *ParamSpecFloat) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *ParamSpecFloat) FieldParentInstance() *ParamSpec {
	argValue := gi.ObjectFieldGet(paramSpecFloatObject, recv.Native(), "parent_instance")
	value := ParamSpecNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *ParamSpecFloat) SetFieldParentInstance(value *ParamSpec) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(paramSpecFloatObject, recv.Native(), "parent_instance", argValue)
}

// FieldMinimum returns the C field 'minimum'.
func (recv *ParamSpecFloat) FieldMinimum() float32 {
	argValue := gi.ObjectFieldGet(paramSpecFloatObject, recv.Native(), "minimum")
	value := argValue.Float32()
	return value
}

// SetFieldMinimum sets the value of the C field 'minimum'.
func (recv *ParamSpecFloat) SetFieldMinimum(value float32) {
	var argValue gi.Argument
	argValue.SetFloat32(value)
	gi.ObjectFieldSet(paramSpecFloatObject, recv.Native(), "minimum", argValue)
}

// FieldMaximum returns the C field 'maximum'.
func (recv *ParamSpecFloat) FieldMaximum() float32 {
	argValue := gi.ObjectFieldGet(paramSpecFloatObject, recv.Native(), "maximum")
	value := argValue.Float32()
	return value
}

// SetFieldMaximum sets the value of the C field 'maximum'.
func (recv *ParamSpecFloat) SetFieldMaximum(value float32) {
	var argValue gi.Argument
	argValue.SetFloat32(value)
	gi.ObjectFieldSet(paramSpecFloatObject, recv.Native(), "maximum", argValue)
}

// FieldDefaultValue returns the C field 'default_value'.
func (recv *ParamSpecFloat) FieldDefaultValue() float32 {
	argValue := gi.ObjectFieldGet(paramSpecFloatObject, recv.Native(), "default_value")
	value := argValue.Float32()
	return value
}

// SetFieldDefaultValue sets the value of the C field 'default_value'.
func (recv *ParamSpecFloat) SetFieldDefaultValue(value float32) {
	var argValue gi.Argument
	argValue.SetFloat32(value)
	gi.ObjectFieldSet(paramSpecFloatObject, recv.Native(), "default_value", argValue)
}

// FieldEpsilon returns the C field 'epsilon'.
func (recv *ParamSpecFloat) FieldEpsilon() float32 {
	argValue := gi.ObjectFieldGet(paramSpecFloatObject, recv.Native(), "epsilon")
	value := argValue.Float32()
	return value
}

// SetFieldEpsilon sets the value of the C field 'epsilon'.
func (recv *ParamSpecFloat) SetFieldEpsilon(value float32) {
	var argValue gi.Argument
	argValue.SetFloat32(value)
	gi.ObjectFieldSet(paramSpecFloatObject, recv.Native(), "epsilon", argValue)
}

var paramSpecGTypeObject *gi.Object
var paramSpecGTypeObject_Once sync.Once

func paramSpecGTypeObject_Set() error {
	var err error
	paramSpecGTypeObject_Once.Do(func() {
		paramSpecGTypeObject, err = gi.ObjectNew("GObject", "ParamSpecGType")
	})
	return err
}

type ParamSpecGType struct {
	native unsafe.Pointer
}

func ParamSpecGTypeNewFromNative(native unsafe.Pointer) *ParamSpecGType {
	instance := &ParamSpecGType{native: native}

	return instance
}

// ParamSpec upcasts to *ParamSpec
func (recv *ParamSpecGType) ParamSpec() *ParamSpec {
	return ParamSpecNewFromNative(recv.Native())
}

// Equals compares this ParamSpecGType with another ParamSpecGType, and returns true if they represent the same Object.
func (recv *ParamSpecGType) Equals(other *ParamSpecGType) bool {
	return other.Native() == recv.Native()
}

func (recv *ParamSpecGType) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *ParamSpecGType) FieldParentInstance() *ParamSpec {
	argValue := gi.ObjectFieldGet(paramSpecGTypeObject, recv.Native(), "parent_instance")
	value := ParamSpecNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *ParamSpecGType) SetFieldParentInstance(value *ParamSpec) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(paramSpecGTypeObject, recv.Native(), "parent_instance", argValue)
}

// UNSUPPORTED : C value 'is_a_type' : for field getter : no Go type for 'GType'

// UNSUPPORTED : C value 'is_a_type' : for field setter : no Go type for 'GType'

var paramSpecIntObject *gi.Object
var paramSpecIntObject_Once sync.Once

func paramSpecIntObject_Set() error {
	var err error
	paramSpecIntObject_Once.Do(func() {
		paramSpecIntObject, err = gi.ObjectNew("GObject", "ParamSpecInt")
	})
	return err
}

type ParamSpecInt struct {
	native unsafe.Pointer
}

func ParamSpecIntNewFromNative(native unsafe.Pointer) *ParamSpecInt {
	instance := &ParamSpecInt{native: native}

	return instance
}

// ParamSpec upcasts to *ParamSpec
func (recv *ParamSpecInt) ParamSpec() *ParamSpec {
	return ParamSpecNewFromNative(recv.Native())
}

// Equals compares this ParamSpecInt with another ParamSpecInt, and returns true if they represent the same Object.
func (recv *ParamSpecInt) Equals(other *ParamSpecInt) bool {
	return other.Native() == recv.Native()
}

func (recv *ParamSpecInt) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *ParamSpecInt) FieldParentInstance() *ParamSpec {
	argValue := gi.ObjectFieldGet(paramSpecIntObject, recv.Native(), "parent_instance")
	value := ParamSpecNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *ParamSpecInt) SetFieldParentInstance(value *ParamSpec) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(paramSpecIntObject, recv.Native(), "parent_instance", argValue)
}

// FieldMinimum returns the C field 'minimum'.
func (recv *ParamSpecInt) FieldMinimum() int32 {
	argValue := gi.ObjectFieldGet(paramSpecIntObject, recv.Native(), "minimum")
	value := argValue.Int32()
	return value
}

// SetFieldMinimum sets the value of the C field 'minimum'.
func (recv *ParamSpecInt) SetFieldMinimum(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.ObjectFieldSet(paramSpecIntObject, recv.Native(), "minimum", argValue)
}

// FieldMaximum returns the C field 'maximum'.
func (recv *ParamSpecInt) FieldMaximum() int32 {
	argValue := gi.ObjectFieldGet(paramSpecIntObject, recv.Native(), "maximum")
	value := argValue.Int32()
	return value
}

// SetFieldMaximum sets the value of the C field 'maximum'.
func (recv *ParamSpecInt) SetFieldMaximum(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.ObjectFieldSet(paramSpecIntObject, recv.Native(), "maximum", argValue)
}

// FieldDefaultValue returns the C field 'default_value'.
func (recv *ParamSpecInt) FieldDefaultValue() int32 {
	argValue := gi.ObjectFieldGet(paramSpecIntObject, recv.Native(), "default_value")
	value := argValue.Int32()
	return value
}

// SetFieldDefaultValue sets the value of the C field 'default_value'.
func (recv *ParamSpecInt) SetFieldDefaultValue(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.ObjectFieldSet(paramSpecIntObject, recv.Native(), "default_value", argValue)
}

var paramSpecInt64Object *gi.Object
var paramSpecInt64Object_Once sync.Once

func paramSpecInt64Object_Set() error {
	var err error
	paramSpecInt64Object_Once.Do(func() {
		paramSpecInt64Object, err = gi.ObjectNew("GObject", "ParamSpecInt64")
	})
	return err
}

type ParamSpecInt64 struct {
	native unsafe.Pointer
}

func ParamSpecInt64NewFromNative(native unsafe.Pointer) *ParamSpecInt64 {
	instance := &ParamSpecInt64{native: native}

	return instance
}

// ParamSpec upcasts to *ParamSpec
func (recv *ParamSpecInt64) ParamSpec() *ParamSpec {
	return ParamSpecNewFromNative(recv.Native())
}

// Equals compares this ParamSpecInt64 with another ParamSpecInt64, and returns true if they represent the same Object.
func (recv *ParamSpecInt64) Equals(other *ParamSpecInt64) bool {
	return other.Native() == recv.Native()
}

func (recv *ParamSpecInt64) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *ParamSpecInt64) FieldParentInstance() *ParamSpec {
	argValue := gi.ObjectFieldGet(paramSpecInt64Object, recv.Native(), "parent_instance")
	value := ParamSpecNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *ParamSpecInt64) SetFieldParentInstance(value *ParamSpec) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(paramSpecInt64Object, recv.Native(), "parent_instance", argValue)
}

// FieldMinimum returns the C field 'minimum'.
func (recv *ParamSpecInt64) FieldMinimum() int64 {
	argValue := gi.ObjectFieldGet(paramSpecInt64Object, recv.Native(), "minimum")
	value := argValue.Int64()
	return value
}

// SetFieldMinimum sets the value of the C field 'minimum'.
func (recv *ParamSpecInt64) SetFieldMinimum(value int64) {
	var argValue gi.Argument
	argValue.SetInt64(value)
	gi.ObjectFieldSet(paramSpecInt64Object, recv.Native(), "minimum", argValue)
}

// FieldMaximum returns the C field 'maximum'.
func (recv *ParamSpecInt64) FieldMaximum() int64 {
	argValue := gi.ObjectFieldGet(paramSpecInt64Object, recv.Native(), "maximum")
	value := argValue.Int64()
	return value
}

// SetFieldMaximum sets the value of the C field 'maximum'.
func (recv *ParamSpecInt64) SetFieldMaximum(value int64) {
	var argValue gi.Argument
	argValue.SetInt64(value)
	gi.ObjectFieldSet(paramSpecInt64Object, recv.Native(), "maximum", argValue)
}

// FieldDefaultValue returns the C field 'default_value'.
func (recv *ParamSpecInt64) FieldDefaultValue() int64 {
	argValue := gi.ObjectFieldGet(paramSpecInt64Object, recv.Native(), "default_value")
	value := argValue.Int64()
	return value
}

// SetFieldDefaultValue sets the value of the C field 'default_value'.
func (recv *ParamSpecInt64) SetFieldDefaultValue(value int64) {
	var argValue gi.Argument
	argValue.SetInt64(value)
	gi.ObjectFieldSet(paramSpecInt64Object, recv.Native(), "default_value", argValue)
}

var paramSpecLongObject *gi.Object
var paramSpecLongObject_Once sync.Once

func paramSpecLongObject_Set() error {
	var err error
	paramSpecLongObject_Once.Do(func() {
		paramSpecLongObject, err = gi.ObjectNew("GObject", "ParamSpecLong")
	})
	return err
}

type ParamSpecLong struct {
	native unsafe.Pointer
}

func ParamSpecLongNewFromNative(native unsafe.Pointer) *ParamSpecLong {
	instance := &ParamSpecLong{native: native}

	return instance
}

// ParamSpec upcasts to *ParamSpec
func (recv *ParamSpecLong) ParamSpec() *ParamSpec {
	return ParamSpecNewFromNative(recv.Native())
}

// Equals compares this ParamSpecLong with another ParamSpecLong, and returns true if they represent the same Object.
func (recv *ParamSpecLong) Equals(other *ParamSpecLong) bool {
	return other.Native() == recv.Native()
}

func (recv *ParamSpecLong) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *ParamSpecLong) FieldParentInstance() *ParamSpec {
	argValue := gi.ObjectFieldGet(paramSpecLongObject, recv.Native(), "parent_instance")
	value := ParamSpecNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *ParamSpecLong) SetFieldParentInstance(value *ParamSpec) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(paramSpecLongObject, recv.Native(), "parent_instance", argValue)
}

// FieldMinimum returns the C field 'minimum'.
func (recv *ParamSpecLong) FieldMinimum() int64 {
	argValue := gi.ObjectFieldGet(paramSpecLongObject, recv.Native(), "minimum")
	value := argValue.Int64()
	return value
}

// SetFieldMinimum sets the value of the C field 'minimum'.
func (recv *ParamSpecLong) SetFieldMinimum(value int64) {
	var argValue gi.Argument
	argValue.SetInt64(value)
	gi.ObjectFieldSet(paramSpecLongObject, recv.Native(), "minimum", argValue)
}

// FieldMaximum returns the C field 'maximum'.
func (recv *ParamSpecLong) FieldMaximum() int64 {
	argValue := gi.ObjectFieldGet(paramSpecLongObject, recv.Native(), "maximum")
	value := argValue.Int64()
	return value
}

// SetFieldMaximum sets the value of the C field 'maximum'.
func (recv *ParamSpecLong) SetFieldMaximum(value int64) {
	var argValue gi.Argument
	argValue.SetInt64(value)
	gi.ObjectFieldSet(paramSpecLongObject, recv.Native(), "maximum", argValue)
}

// FieldDefaultValue returns the C field 'default_value'.
func (recv *ParamSpecLong) FieldDefaultValue() int64 {
	argValue := gi.ObjectFieldGet(paramSpecLongObject, recv.Native(), "default_value")
	value := argValue.Int64()
	return value
}

// SetFieldDefaultValue sets the value of the C field 'default_value'.
func (recv *ParamSpecLong) SetFieldDefaultValue(value int64) {
	var argValue gi.Argument
	argValue.SetInt64(value)
	gi.ObjectFieldSet(paramSpecLongObject, recv.Native(), "default_value", argValue)
}

var paramSpecObjectObject *gi.Object
var paramSpecObjectObject_Once sync.Once

func paramSpecObjectObject_Set() error {
	var err error
	paramSpecObjectObject_Once.Do(func() {
		paramSpecObjectObject, err = gi.ObjectNew("GObject", "ParamSpecObject")
	})
	return err
}

type ParamSpecObject struct {
	native unsafe.Pointer
}

func ParamSpecObjectNewFromNative(native unsafe.Pointer) *ParamSpecObject {
	instance := &ParamSpecObject{native: native}

	return instance
}

// ParamSpec upcasts to *ParamSpec
func (recv *ParamSpecObject) ParamSpec() *ParamSpec {
	return ParamSpecNewFromNative(recv.Native())
}

// Equals compares this ParamSpecObject with another ParamSpecObject, and returns true if they represent the same Object.
func (recv *ParamSpecObject) Equals(other *ParamSpecObject) bool {
	return other.Native() == recv.Native()
}

func (recv *ParamSpecObject) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *ParamSpecObject) FieldParentInstance() *ParamSpec {
	argValue := gi.ObjectFieldGet(paramSpecObjectObject, recv.Native(), "parent_instance")
	value := ParamSpecNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *ParamSpecObject) SetFieldParentInstance(value *ParamSpec) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(paramSpecObjectObject, recv.Native(), "parent_instance", argValue)
}

var paramSpecOverrideObject *gi.Object
var paramSpecOverrideObject_Once sync.Once

func paramSpecOverrideObject_Set() error {
	var err error
	paramSpecOverrideObject_Once.Do(func() {
		paramSpecOverrideObject, err = gi.ObjectNew("GObject", "ParamSpecOverride")
	})
	return err
}

type ParamSpecOverride struct {
	native unsafe.Pointer
}

func ParamSpecOverrideNewFromNative(native unsafe.Pointer) *ParamSpecOverride {
	instance := &ParamSpecOverride{native: native}

	return instance
}

// ParamSpec upcasts to *ParamSpec
func (recv *ParamSpecOverride) ParamSpec() *ParamSpec {
	return ParamSpecNewFromNative(recv.Native())
}

// Equals compares this ParamSpecOverride with another ParamSpecOverride, and returns true if they represent the same Object.
func (recv *ParamSpecOverride) Equals(other *ParamSpecOverride) bool {
	return other.Native() == recv.Native()
}

func (recv *ParamSpecOverride) Native() unsafe.Pointer {
	return recv.native
}

var paramSpecParamObject *gi.Object
var paramSpecParamObject_Once sync.Once

func paramSpecParamObject_Set() error {
	var err error
	paramSpecParamObject_Once.Do(func() {
		paramSpecParamObject, err = gi.ObjectNew("GObject", "ParamSpecParam")
	})
	return err
}

type ParamSpecParam struct {
	native unsafe.Pointer
}

func ParamSpecParamNewFromNative(native unsafe.Pointer) *ParamSpecParam {
	instance := &ParamSpecParam{native: native}

	return instance
}

// ParamSpec upcasts to *ParamSpec
func (recv *ParamSpecParam) ParamSpec() *ParamSpec {
	return ParamSpecNewFromNative(recv.Native())
}

// Equals compares this ParamSpecParam with another ParamSpecParam, and returns true if they represent the same Object.
func (recv *ParamSpecParam) Equals(other *ParamSpecParam) bool {
	return other.Native() == recv.Native()
}

func (recv *ParamSpecParam) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *ParamSpecParam) FieldParentInstance() *ParamSpec {
	argValue := gi.ObjectFieldGet(paramSpecParamObject, recv.Native(), "parent_instance")
	value := ParamSpecNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *ParamSpecParam) SetFieldParentInstance(value *ParamSpec) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(paramSpecParamObject, recv.Native(), "parent_instance", argValue)
}

var paramSpecPointerObject *gi.Object
var paramSpecPointerObject_Once sync.Once

func paramSpecPointerObject_Set() error {
	var err error
	paramSpecPointerObject_Once.Do(func() {
		paramSpecPointerObject, err = gi.ObjectNew("GObject", "ParamSpecPointer")
	})
	return err
}

type ParamSpecPointer struct {
	native unsafe.Pointer
}

func ParamSpecPointerNewFromNative(native unsafe.Pointer) *ParamSpecPointer {
	instance := &ParamSpecPointer{native: native}

	return instance
}

// ParamSpec upcasts to *ParamSpec
func (recv *ParamSpecPointer) ParamSpec() *ParamSpec {
	return ParamSpecNewFromNative(recv.Native())
}

// Equals compares this ParamSpecPointer with another ParamSpecPointer, and returns true if they represent the same Object.
func (recv *ParamSpecPointer) Equals(other *ParamSpecPointer) bool {
	return other.Native() == recv.Native()
}

func (recv *ParamSpecPointer) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *ParamSpecPointer) FieldParentInstance() *ParamSpec {
	argValue := gi.ObjectFieldGet(paramSpecPointerObject, recv.Native(), "parent_instance")
	value := ParamSpecNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *ParamSpecPointer) SetFieldParentInstance(value *ParamSpec) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(paramSpecPointerObject, recv.Native(), "parent_instance", argValue)
}

var paramSpecStringObject *gi.Object
var paramSpecStringObject_Once sync.Once

func paramSpecStringObject_Set() error {
	var err error
	paramSpecStringObject_Once.Do(func() {
		paramSpecStringObject, err = gi.ObjectNew("GObject", "ParamSpecString")
	})
	return err
}

type ParamSpecString struct {
	native unsafe.Pointer
}

func ParamSpecStringNewFromNative(native unsafe.Pointer) *ParamSpecString {
	instance := &ParamSpecString{native: native}

	return instance
}

// ParamSpec upcasts to *ParamSpec
func (recv *ParamSpecString) ParamSpec() *ParamSpec {
	return ParamSpecNewFromNative(recv.Native())
}

// Equals compares this ParamSpecString with another ParamSpecString, and returns true if they represent the same Object.
func (recv *ParamSpecString) Equals(other *ParamSpecString) bool {
	return other.Native() == recv.Native()
}

func (recv *ParamSpecString) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *ParamSpecString) FieldParentInstance() *ParamSpec {
	argValue := gi.ObjectFieldGet(paramSpecStringObject, recv.Native(), "parent_instance")
	value := ParamSpecNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *ParamSpecString) SetFieldParentInstance(value *ParamSpec) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(paramSpecStringObject, recv.Native(), "parent_instance", argValue)
}

// FieldDefaultValue returns the C field 'default_value'.
func (recv *ParamSpecString) FieldDefaultValue() string {
	argValue := gi.ObjectFieldGet(paramSpecStringObject, recv.Native(), "default_value")
	value := argValue.String(false)
	return value
}

// SetFieldDefaultValue sets the value of the C field 'default_value'.
func (recv *ParamSpecString) SetFieldDefaultValue(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.ObjectFieldSet(paramSpecStringObject, recv.Native(), "default_value", argValue)
}

// FieldCsetFirst returns the C field 'cset_first'.
func (recv *ParamSpecString) FieldCsetFirst() string {
	argValue := gi.ObjectFieldGet(paramSpecStringObject, recv.Native(), "cset_first")
	value := argValue.String(false)
	return value
}

// SetFieldCsetFirst sets the value of the C field 'cset_first'.
func (recv *ParamSpecString) SetFieldCsetFirst(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.ObjectFieldSet(paramSpecStringObject, recv.Native(), "cset_first", argValue)
}

// FieldCsetNth returns the C field 'cset_nth'.
func (recv *ParamSpecString) FieldCsetNth() string {
	argValue := gi.ObjectFieldGet(paramSpecStringObject, recv.Native(), "cset_nth")
	value := argValue.String(false)
	return value
}

// SetFieldCsetNth sets the value of the C field 'cset_nth'.
func (recv *ParamSpecString) SetFieldCsetNth(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.ObjectFieldSet(paramSpecStringObject, recv.Native(), "cset_nth", argValue)
}

// FieldSubstitutor returns the C field 'substitutor'.
func (recv *ParamSpecString) FieldSubstitutor() int8 {
	argValue := gi.ObjectFieldGet(paramSpecStringObject, recv.Native(), "substitutor")
	value := argValue.Int8()
	return value
}

// SetFieldSubstitutor sets the value of the C field 'substitutor'.
func (recv *ParamSpecString) SetFieldSubstitutor(value int8) {
	var argValue gi.Argument
	argValue.SetInt8(value)
	gi.ObjectFieldSet(paramSpecStringObject, recv.Native(), "substitutor", argValue)
}

// FieldNullFoldIfEmpty returns the C field 'null_fold_if_empty'.
func (recv *ParamSpecString) FieldNullFoldIfEmpty() uint32 {
	argValue := gi.ObjectFieldGet(paramSpecStringObject, recv.Native(), "null_fold_if_empty")
	value := argValue.Uint32()
	return value
}

// SetFieldNullFoldIfEmpty sets the value of the C field 'null_fold_if_empty'.
func (recv *ParamSpecString) SetFieldNullFoldIfEmpty(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.ObjectFieldSet(paramSpecStringObject, recv.Native(), "null_fold_if_empty", argValue)
}

// FieldEnsureNonNull returns the C field 'ensure_non_null'.
func (recv *ParamSpecString) FieldEnsureNonNull() uint32 {
	argValue := gi.ObjectFieldGet(paramSpecStringObject, recv.Native(), "ensure_non_null")
	value := argValue.Uint32()
	return value
}

// SetFieldEnsureNonNull sets the value of the C field 'ensure_non_null'.
func (recv *ParamSpecString) SetFieldEnsureNonNull(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.ObjectFieldSet(paramSpecStringObject, recv.Native(), "ensure_non_null", argValue)
}

var paramSpecUCharObject *gi.Object
var paramSpecUCharObject_Once sync.Once

func paramSpecUCharObject_Set() error {
	var err error
	paramSpecUCharObject_Once.Do(func() {
		paramSpecUCharObject, err = gi.ObjectNew("GObject", "ParamSpecUChar")
	})
	return err
}

type ParamSpecUChar struct {
	native unsafe.Pointer
}

func ParamSpecUCharNewFromNative(native unsafe.Pointer) *ParamSpecUChar {
	instance := &ParamSpecUChar{native: native}

	return instance
}

// ParamSpec upcasts to *ParamSpec
func (recv *ParamSpecUChar) ParamSpec() *ParamSpec {
	return ParamSpecNewFromNative(recv.Native())
}

// Equals compares this ParamSpecUChar with another ParamSpecUChar, and returns true if they represent the same Object.
func (recv *ParamSpecUChar) Equals(other *ParamSpecUChar) bool {
	return other.Native() == recv.Native()
}

func (recv *ParamSpecUChar) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *ParamSpecUChar) FieldParentInstance() *ParamSpec {
	argValue := gi.ObjectFieldGet(paramSpecUCharObject, recv.Native(), "parent_instance")
	value := ParamSpecNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *ParamSpecUChar) SetFieldParentInstance(value *ParamSpec) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(paramSpecUCharObject, recv.Native(), "parent_instance", argValue)
}

// FieldMinimum returns the C field 'minimum'.
func (recv *ParamSpecUChar) FieldMinimum() uint8 {
	argValue := gi.ObjectFieldGet(paramSpecUCharObject, recv.Native(), "minimum")
	value := argValue.Uint8()
	return value
}

// SetFieldMinimum sets the value of the C field 'minimum'.
func (recv *ParamSpecUChar) SetFieldMinimum(value uint8) {
	var argValue gi.Argument
	argValue.SetUint8(value)
	gi.ObjectFieldSet(paramSpecUCharObject, recv.Native(), "minimum", argValue)
}

// FieldMaximum returns the C field 'maximum'.
func (recv *ParamSpecUChar) FieldMaximum() uint8 {
	argValue := gi.ObjectFieldGet(paramSpecUCharObject, recv.Native(), "maximum")
	value := argValue.Uint8()
	return value
}

// SetFieldMaximum sets the value of the C field 'maximum'.
func (recv *ParamSpecUChar) SetFieldMaximum(value uint8) {
	var argValue gi.Argument
	argValue.SetUint8(value)
	gi.ObjectFieldSet(paramSpecUCharObject, recv.Native(), "maximum", argValue)
}

// FieldDefaultValue returns the C field 'default_value'.
func (recv *ParamSpecUChar) FieldDefaultValue() uint8 {
	argValue := gi.ObjectFieldGet(paramSpecUCharObject, recv.Native(), "default_value")
	value := argValue.Uint8()
	return value
}

// SetFieldDefaultValue sets the value of the C field 'default_value'.
func (recv *ParamSpecUChar) SetFieldDefaultValue(value uint8) {
	var argValue gi.Argument
	argValue.SetUint8(value)
	gi.ObjectFieldSet(paramSpecUCharObject, recv.Native(), "default_value", argValue)
}

var paramSpecUIntObject *gi.Object
var paramSpecUIntObject_Once sync.Once

func paramSpecUIntObject_Set() error {
	var err error
	paramSpecUIntObject_Once.Do(func() {
		paramSpecUIntObject, err = gi.ObjectNew("GObject", "ParamSpecUInt")
	})
	return err
}

type ParamSpecUInt struct {
	native unsafe.Pointer
}

func ParamSpecUIntNewFromNative(native unsafe.Pointer) *ParamSpecUInt {
	instance := &ParamSpecUInt{native: native}

	return instance
}

// ParamSpec upcasts to *ParamSpec
func (recv *ParamSpecUInt) ParamSpec() *ParamSpec {
	return ParamSpecNewFromNative(recv.Native())
}

// Equals compares this ParamSpecUInt with another ParamSpecUInt, and returns true if they represent the same Object.
func (recv *ParamSpecUInt) Equals(other *ParamSpecUInt) bool {
	return other.Native() == recv.Native()
}

func (recv *ParamSpecUInt) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *ParamSpecUInt) FieldParentInstance() *ParamSpec {
	argValue := gi.ObjectFieldGet(paramSpecUIntObject, recv.Native(), "parent_instance")
	value := ParamSpecNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *ParamSpecUInt) SetFieldParentInstance(value *ParamSpec) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(paramSpecUIntObject, recv.Native(), "parent_instance", argValue)
}

// FieldMinimum returns the C field 'minimum'.
func (recv *ParamSpecUInt) FieldMinimum() uint32 {
	argValue := gi.ObjectFieldGet(paramSpecUIntObject, recv.Native(), "minimum")
	value := argValue.Uint32()
	return value
}

// SetFieldMinimum sets the value of the C field 'minimum'.
func (recv *ParamSpecUInt) SetFieldMinimum(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.ObjectFieldSet(paramSpecUIntObject, recv.Native(), "minimum", argValue)
}

// FieldMaximum returns the C field 'maximum'.
func (recv *ParamSpecUInt) FieldMaximum() uint32 {
	argValue := gi.ObjectFieldGet(paramSpecUIntObject, recv.Native(), "maximum")
	value := argValue.Uint32()
	return value
}

// SetFieldMaximum sets the value of the C field 'maximum'.
func (recv *ParamSpecUInt) SetFieldMaximum(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.ObjectFieldSet(paramSpecUIntObject, recv.Native(), "maximum", argValue)
}

// FieldDefaultValue returns the C field 'default_value'.
func (recv *ParamSpecUInt) FieldDefaultValue() uint32 {
	argValue := gi.ObjectFieldGet(paramSpecUIntObject, recv.Native(), "default_value")
	value := argValue.Uint32()
	return value
}

// SetFieldDefaultValue sets the value of the C field 'default_value'.
func (recv *ParamSpecUInt) SetFieldDefaultValue(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.ObjectFieldSet(paramSpecUIntObject, recv.Native(), "default_value", argValue)
}

var paramSpecUInt64Object *gi.Object
var paramSpecUInt64Object_Once sync.Once

func paramSpecUInt64Object_Set() error {
	var err error
	paramSpecUInt64Object_Once.Do(func() {
		paramSpecUInt64Object, err = gi.ObjectNew("GObject", "ParamSpecUInt64")
	})
	return err
}

type ParamSpecUInt64 struct {
	native unsafe.Pointer
}

func ParamSpecUInt64NewFromNative(native unsafe.Pointer) *ParamSpecUInt64 {
	instance := &ParamSpecUInt64{native: native}

	return instance
}

// ParamSpec upcasts to *ParamSpec
func (recv *ParamSpecUInt64) ParamSpec() *ParamSpec {
	return ParamSpecNewFromNative(recv.Native())
}

// Equals compares this ParamSpecUInt64 with another ParamSpecUInt64, and returns true if they represent the same Object.
func (recv *ParamSpecUInt64) Equals(other *ParamSpecUInt64) bool {
	return other.Native() == recv.Native()
}

func (recv *ParamSpecUInt64) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *ParamSpecUInt64) FieldParentInstance() *ParamSpec {
	argValue := gi.ObjectFieldGet(paramSpecUInt64Object, recv.Native(), "parent_instance")
	value := ParamSpecNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *ParamSpecUInt64) SetFieldParentInstance(value *ParamSpec) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(paramSpecUInt64Object, recv.Native(), "parent_instance", argValue)
}

// FieldMinimum returns the C field 'minimum'.
func (recv *ParamSpecUInt64) FieldMinimum() uint64 {
	argValue := gi.ObjectFieldGet(paramSpecUInt64Object, recv.Native(), "minimum")
	value := argValue.Uint64()
	return value
}

// SetFieldMinimum sets the value of the C field 'minimum'.
func (recv *ParamSpecUInt64) SetFieldMinimum(value uint64) {
	var argValue gi.Argument
	argValue.SetUint64(value)
	gi.ObjectFieldSet(paramSpecUInt64Object, recv.Native(), "minimum", argValue)
}

// FieldMaximum returns the C field 'maximum'.
func (recv *ParamSpecUInt64) FieldMaximum() uint64 {
	argValue := gi.ObjectFieldGet(paramSpecUInt64Object, recv.Native(), "maximum")
	value := argValue.Uint64()
	return value
}

// SetFieldMaximum sets the value of the C field 'maximum'.
func (recv *ParamSpecUInt64) SetFieldMaximum(value uint64) {
	var argValue gi.Argument
	argValue.SetUint64(value)
	gi.ObjectFieldSet(paramSpecUInt64Object, recv.Native(), "maximum", argValue)
}

// FieldDefaultValue returns the C field 'default_value'.
func (recv *ParamSpecUInt64) FieldDefaultValue() uint64 {
	argValue := gi.ObjectFieldGet(paramSpecUInt64Object, recv.Native(), "default_value")
	value := argValue.Uint64()
	return value
}

// SetFieldDefaultValue sets the value of the C field 'default_value'.
func (recv *ParamSpecUInt64) SetFieldDefaultValue(value uint64) {
	var argValue gi.Argument
	argValue.SetUint64(value)
	gi.ObjectFieldSet(paramSpecUInt64Object, recv.Native(), "default_value", argValue)
}

var paramSpecULongObject *gi.Object
var paramSpecULongObject_Once sync.Once

func paramSpecULongObject_Set() error {
	var err error
	paramSpecULongObject_Once.Do(func() {
		paramSpecULongObject, err = gi.ObjectNew("GObject", "ParamSpecULong")
	})
	return err
}

type ParamSpecULong struct {
	native unsafe.Pointer
}

func ParamSpecULongNewFromNative(native unsafe.Pointer) *ParamSpecULong {
	instance := &ParamSpecULong{native: native}

	return instance
}

// ParamSpec upcasts to *ParamSpec
func (recv *ParamSpecULong) ParamSpec() *ParamSpec {
	return ParamSpecNewFromNative(recv.Native())
}

// Equals compares this ParamSpecULong with another ParamSpecULong, and returns true if they represent the same Object.
func (recv *ParamSpecULong) Equals(other *ParamSpecULong) bool {
	return other.Native() == recv.Native()
}

func (recv *ParamSpecULong) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *ParamSpecULong) FieldParentInstance() *ParamSpec {
	argValue := gi.ObjectFieldGet(paramSpecULongObject, recv.Native(), "parent_instance")
	value := ParamSpecNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *ParamSpecULong) SetFieldParentInstance(value *ParamSpec) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(paramSpecULongObject, recv.Native(), "parent_instance", argValue)
}

// FieldMinimum returns the C field 'minimum'.
func (recv *ParamSpecULong) FieldMinimum() uint64 {
	argValue := gi.ObjectFieldGet(paramSpecULongObject, recv.Native(), "minimum")
	value := argValue.Uint64()
	return value
}

// SetFieldMinimum sets the value of the C field 'minimum'.
func (recv *ParamSpecULong) SetFieldMinimum(value uint64) {
	var argValue gi.Argument
	argValue.SetUint64(value)
	gi.ObjectFieldSet(paramSpecULongObject, recv.Native(), "minimum", argValue)
}

// FieldMaximum returns the C field 'maximum'.
func (recv *ParamSpecULong) FieldMaximum() uint64 {
	argValue := gi.ObjectFieldGet(paramSpecULongObject, recv.Native(), "maximum")
	value := argValue.Uint64()
	return value
}

// SetFieldMaximum sets the value of the C field 'maximum'.
func (recv *ParamSpecULong) SetFieldMaximum(value uint64) {
	var argValue gi.Argument
	argValue.SetUint64(value)
	gi.ObjectFieldSet(paramSpecULongObject, recv.Native(), "maximum", argValue)
}

// FieldDefaultValue returns the C field 'default_value'.
func (recv *ParamSpecULong) FieldDefaultValue() uint64 {
	argValue := gi.ObjectFieldGet(paramSpecULongObject, recv.Native(), "default_value")
	value := argValue.Uint64()
	return value
}

// SetFieldDefaultValue sets the value of the C field 'default_value'.
func (recv *ParamSpecULong) SetFieldDefaultValue(value uint64) {
	var argValue gi.Argument
	argValue.SetUint64(value)
	gi.ObjectFieldSet(paramSpecULongObject, recv.Native(), "default_value", argValue)
}

var paramSpecUnicharObject *gi.Object
var paramSpecUnicharObject_Once sync.Once

func paramSpecUnicharObject_Set() error {
	var err error
	paramSpecUnicharObject_Once.Do(func() {
		paramSpecUnicharObject, err = gi.ObjectNew("GObject", "ParamSpecUnichar")
	})
	return err
}

type ParamSpecUnichar struct {
	native unsafe.Pointer
}

func ParamSpecUnicharNewFromNative(native unsafe.Pointer) *ParamSpecUnichar {
	instance := &ParamSpecUnichar{native: native}

	return instance
}

// ParamSpec upcasts to *ParamSpec
func (recv *ParamSpecUnichar) ParamSpec() *ParamSpec {
	return ParamSpecNewFromNative(recv.Native())
}

// Equals compares this ParamSpecUnichar with another ParamSpecUnichar, and returns true if they represent the same Object.
func (recv *ParamSpecUnichar) Equals(other *ParamSpecUnichar) bool {
	return other.Native() == recv.Native()
}

func (recv *ParamSpecUnichar) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *ParamSpecUnichar) FieldParentInstance() *ParamSpec {
	argValue := gi.ObjectFieldGet(paramSpecUnicharObject, recv.Native(), "parent_instance")
	value := ParamSpecNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *ParamSpecUnichar) SetFieldParentInstance(value *ParamSpec) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(paramSpecUnicharObject, recv.Native(), "parent_instance", argValue)
}

// UNSUPPORTED : C value 'default_value' : for field getter : no Go type for 'gunichar'

// UNSUPPORTED : C value 'default_value' : for field setter : no Go type for 'gunichar'

var paramSpecValueArrayObject *gi.Object
var paramSpecValueArrayObject_Once sync.Once

func paramSpecValueArrayObject_Set() error {
	var err error
	paramSpecValueArrayObject_Once.Do(func() {
		paramSpecValueArrayObject, err = gi.ObjectNew("GObject", "ParamSpecValueArray")
	})
	return err
}

type ParamSpecValueArray struct {
	native unsafe.Pointer
}

func ParamSpecValueArrayNewFromNative(native unsafe.Pointer) *ParamSpecValueArray {
	instance := &ParamSpecValueArray{native: native}

	return instance
}

// ParamSpec upcasts to *ParamSpec
func (recv *ParamSpecValueArray) ParamSpec() *ParamSpec {
	return ParamSpecNewFromNative(recv.Native())
}

// Equals compares this ParamSpecValueArray with another ParamSpecValueArray, and returns true if they represent the same Object.
func (recv *ParamSpecValueArray) Equals(other *ParamSpecValueArray) bool {
	return other.Native() == recv.Native()
}

func (recv *ParamSpecValueArray) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *ParamSpecValueArray) FieldParentInstance() *ParamSpec {
	argValue := gi.ObjectFieldGet(paramSpecValueArrayObject, recv.Native(), "parent_instance")
	value := ParamSpecNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *ParamSpecValueArray) SetFieldParentInstance(value *ParamSpec) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(paramSpecValueArrayObject, recv.Native(), "parent_instance", argValue)
}

// FieldElementSpec returns the C field 'element_spec'.
func (recv *ParamSpecValueArray) FieldElementSpec() *ParamSpec {
	argValue := gi.ObjectFieldGet(paramSpecValueArrayObject, recv.Native(), "element_spec")
	value := ParamSpecNewFromNative(argValue.Pointer())
	return value
}

// SetFieldElementSpec sets the value of the C field 'element_spec'.
func (recv *ParamSpecValueArray) SetFieldElementSpec(value *ParamSpec) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(paramSpecValueArrayObject, recv.Native(), "element_spec", argValue)
}

// FieldFixedNElements returns the C field 'fixed_n_elements'.
func (recv *ParamSpecValueArray) FieldFixedNElements() uint32 {
	argValue := gi.ObjectFieldGet(paramSpecValueArrayObject, recv.Native(), "fixed_n_elements")
	value := argValue.Uint32()
	return value
}

// SetFieldFixedNElements sets the value of the C field 'fixed_n_elements'.
func (recv *ParamSpecValueArray) SetFieldFixedNElements(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.ObjectFieldSet(paramSpecValueArrayObject, recv.Native(), "fixed_n_elements", argValue)
}

var paramSpecVariantObject *gi.Object
var paramSpecVariantObject_Once sync.Once

func paramSpecVariantObject_Set() error {
	var err error
	paramSpecVariantObject_Once.Do(func() {
		paramSpecVariantObject, err = gi.ObjectNew("GObject", "ParamSpecVariant")
	})
	return err
}

type ParamSpecVariant struct {
	native unsafe.Pointer
}

func ParamSpecVariantNewFromNative(native unsafe.Pointer) *ParamSpecVariant {
	instance := &ParamSpecVariant{native: native}

	return instance
}

// ParamSpec upcasts to *ParamSpec
func (recv *ParamSpecVariant) ParamSpec() *ParamSpec {
	return ParamSpecNewFromNative(recv.Native())
}

// Equals compares this ParamSpecVariant with another ParamSpecVariant, and returns true if they represent the same Object.
func (recv *ParamSpecVariant) Equals(other *ParamSpecVariant) bool {
	return other.Native() == recv.Native()
}

func (recv *ParamSpecVariant) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *ParamSpecVariant) FieldParentInstance() *ParamSpec {
	argValue := gi.ObjectFieldGet(paramSpecVariantObject, recv.Native(), "parent_instance")
	value := ParamSpecNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *ParamSpecVariant) SetFieldParentInstance(value *ParamSpec) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(paramSpecVariantObject, recv.Native(), "parent_instance", argValue)
}

// UNSUPPORTED : C value 'type' : for field getter : no Go type for 'GLib.VariantType'

// UNSUPPORTED : C value 'type' : for field setter : no Go type for 'GLib.VariantType'

// UNSUPPORTED : C value 'default_value' : for field getter : no Go type for 'GLib.Variant'

// UNSUPPORTED : C value 'default_value' : for field setter : no Go type for 'GLib.Variant'

var typeModuleObject *gi.Object
var typeModuleObject_Once sync.Once

func typeModuleObject_Set() error {
	var err error
	typeModuleObject_Once.Do(func() {
		typeModuleObject, err = gi.ObjectNew("GObject", "TypeModule")
	})
	return err
}

type TypeModule struct {
	native unsafe.Pointer
}

func TypeModuleNewFromNative(native unsafe.Pointer) *TypeModule {
	instance := &TypeModule{native: native}

	return instance
}

// Object upcasts to *Object
func (recv *TypeModule) Object() *Object {
	return ObjectNewFromNative(recv.Native())
}

// Equals compares this TypeModule with another TypeModule, and returns true if they represent the same Object.
func (recv *TypeModule) Equals(other *TypeModule) bool {
	return other.Native() == recv.Native()
}

func (recv *TypeModule) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *TypeModule) FieldParentInstance() *Object {
	argValue := gi.ObjectFieldGet(typeModuleObject, recv.Native(), "parent_instance")
	value := ObjectNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *TypeModule) SetFieldParentInstance(value *Object) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(typeModuleObject, recv.Native(), "parent_instance", argValue)
}

// FieldUseCount returns the C field 'use_count'.
func (recv *TypeModule) FieldUseCount() uint32 {
	argValue := gi.ObjectFieldGet(typeModuleObject, recv.Native(), "use_count")
	value := argValue.Uint32()
	return value
}

// SetFieldUseCount sets the value of the C field 'use_count'.
func (recv *TypeModule) SetFieldUseCount(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.ObjectFieldSet(typeModuleObject, recv.Native(), "use_count", argValue)
}

// UNSUPPORTED : C value 'type_infos' : for field getter : no Go type for 'GLib.SList'

// UNSUPPORTED : C value 'type_infos' : for field setter : no Go type for 'GLib.SList'

// UNSUPPORTED : C value 'interface_infos' : for field getter : no Go type for 'GLib.SList'

// UNSUPPORTED : C value 'interface_infos' : for field setter : no Go type for 'GLib.SList'

// FieldName returns the C field 'name'.
func (recv *TypeModule) FieldName() string {
	argValue := gi.ObjectFieldGet(typeModuleObject, recv.Native(), "name")
	value := argValue.String(false)
	return value
}

// SetFieldName sets the value of the C field 'name'.
func (recv *TypeModule) SetFieldName(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.ObjectFieldSet(typeModuleObject, recv.Native(), "name", argValue)
}

// UNSUPPORTED : C value 'g_type_module_add_interface' : parameter 'instance_type' of type 'GType' not supported

// UNSUPPORTED : C value 'g_type_module_register_enum' : return type 'GType' not supported

// UNSUPPORTED : C value 'g_type_module_register_flags' : return type 'GType' not supported

// UNSUPPORTED : C value 'g_type_module_register_type' : parameter 'parent_type' of type 'GType' not supported

var typeModuleSetNameFunction *gi.Function
var typeModuleSetNameFunction_Once sync.Once

func typeModuleSetNameFunction_Set() error {
	var err error
	typeModuleSetNameFunction_Once.Do(func() {
		err = typeModuleObject_Set()
		if err != nil {
			return
		}
		typeModuleSetNameFunction, err = typeModuleObject.InvokerNew("set_name")
	})
	return err
}

// SetName is a representation of the C type g_type_module_set_name.
func (recv *TypeModule) SetName(name string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(name)

	err := typeModuleSetNameFunction_Set()
	if err == nil {
		typeModuleSetNameFunction.Invoke(inArgs[:], nil)
	}

	return
}

var typeModuleUnuseFunction *gi.Function
var typeModuleUnuseFunction_Once sync.Once

func typeModuleUnuseFunction_Set() error {
	var err error
	typeModuleUnuseFunction_Once.Do(func() {
		err = typeModuleObject_Set()
		if err != nil {
			return
		}
		typeModuleUnuseFunction, err = typeModuleObject.InvokerNew("unuse")
	})
	return err
}

// Unuse is a representation of the C type g_type_module_unuse.
func (recv *TypeModule) Unuse() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := typeModuleUnuseFunction_Set()
	if err == nil {
		typeModuleUnuseFunction.Invoke(inArgs[:], nil)
	}

	return
}

var typeModuleUseFunction *gi.Function
var typeModuleUseFunction_Once sync.Once

func typeModuleUseFunction_Set() error {
	var err error
	typeModuleUseFunction_Once.Do(func() {
		err = typeModuleObject_Set()
		if err != nil {
			return
		}
		typeModuleUseFunction, err = typeModuleObject.InvokerNew("use")
	})
	return err
}

// Use is a representation of the C type g_type_module_use.
func (recv *TypeModule) Use() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := typeModuleUseFunction_Set()
	if err == nil {
		ret = typeModuleUseFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// TypePlugin returns the TypePlugin interface implemented by TypeModule
func (recv *TypeModule) TypePlugin() *TypePlugin {
	return TypePluginNewFromNative(recv.Native())
}
