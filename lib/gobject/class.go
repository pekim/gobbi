// Code generated - DO NOT EDIT.

package gobject

import (
	gi "github.com/pekim/gobbi/internal/gi"
	glib "github.com/pekim/gobbi/lib/glib"
	"runtime"
	"sync"
)

var bindingStruct *gi.Struct
var bindingStruct_Once sync.Once

func bindingStruct_Set() error {
	var err error
	bindingStruct_Once.Do(func() {
		bindingStruct, err = gi.StructNew("GObject", "Binding")
	})
	return err
}

type Binding struct {
	native uintptr
}

// UNSUPPORTED : C value 'g_binding_get_flags' : return type 'BindingFlags' not supported

var bindingGetSourceFunction *gi.Function
var bindingGetSourceFunction_Once sync.Once

func bindingGetSourceFunction_Set() error {
	var err error
	bindingGetSourceFunction_Once.Do(func() {
		err = bindingStruct_Set()
		if err != nil {
			return
		}
		bindingGetSourceFunction, err = bindingStruct.InvokerNew("get_source")
	})
	return err
}

// GetSource is a representation of the C type g_binding_get_source.
func (recv *Binding) GetSource() *Object {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := bindingGetSourceFunction_Set()
	if err == nil {
		ret = bindingGetSourceFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Object{native: ret.Pointer()}

	return retGo
}

var bindingGetSourcePropertyFunction *gi.Function
var bindingGetSourcePropertyFunction_Once sync.Once

func bindingGetSourcePropertyFunction_Set() error {
	var err error
	bindingGetSourcePropertyFunction_Once.Do(func() {
		err = bindingStruct_Set()
		if err != nil {
			return
		}
		bindingGetSourcePropertyFunction, err = bindingStruct.InvokerNew("get_source_property")
	})
	return err
}

// GetSourceProperty is a representation of the C type g_binding_get_source_property.
func (recv *Binding) GetSourceProperty() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

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
		err = bindingStruct_Set()
		if err != nil {
			return
		}
		bindingGetTargetFunction, err = bindingStruct.InvokerNew("get_target")
	})
	return err
}

// GetTarget is a representation of the C type g_binding_get_target.
func (recv *Binding) GetTarget() *Object {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := bindingGetTargetFunction_Set()
	if err == nil {
		ret = bindingGetTargetFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Object{native: ret.Pointer()}

	return retGo
}

var bindingGetTargetPropertyFunction *gi.Function
var bindingGetTargetPropertyFunction_Once sync.Once

func bindingGetTargetPropertyFunction_Set() error {
	var err error
	bindingGetTargetPropertyFunction_Once.Do(func() {
		err = bindingStruct_Set()
		if err != nil {
			return
		}
		bindingGetTargetPropertyFunction, err = bindingStruct.InvokerNew("get_target_property")
	})
	return err
}

// GetTargetProperty is a representation of the C type g_binding_get_target_property.
func (recv *Binding) GetTargetProperty() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

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
		err = bindingStruct_Set()
		if err != nil {
			return
		}
		bindingUnbindFunction, err = bindingStruct.InvokerNew("unbind")
	})
	return err
}

// Unbind is a representation of the C type g_binding_unbind.
func (recv *Binding) Unbind() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := bindingUnbindFunction_Set()
	if err == nil {
		bindingUnbindFunction.Invoke(inArgs[:], nil)
	}

	return
}

// BindingStruct creates an uninitialised Binding.
func BindingStruct() *Binding {
	err := bindingStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &Binding{native: bindingStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeBinding)
	return structGo
}
func finalizeBinding(obj *Binding) {
	bindingStruct.Free(obj.native)
}

var initiallyUnownedStruct *gi.Struct
var initiallyUnownedStruct_Once sync.Once

func initiallyUnownedStruct_Set() error {
	var err error
	initiallyUnownedStruct_Once.Do(func() {
		initiallyUnownedStruct, err = gi.StructNew("GObject", "InitiallyUnowned")
	})
	return err
}

type InitiallyUnowned struct {
	native uintptr
}

// FieldGTypeInstance returns the C field 'g_type_instance'.
func (recv *InitiallyUnowned) FieldGTypeInstance() *TypeInstance {
	argValue := gi.FieldGet(initiallyUnownedStruct, recv.native, "g_type_instance")
	value := &TypeInstance{native: argValue.Pointer()}
	return value
}

// SetFieldGTypeInstance sets the value of the C field 'g_type_instance'.
func (recv *InitiallyUnowned) SetFieldGTypeInstance(value *TypeInstance) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(initiallyUnownedStruct, recv.native, "g_type_instance", argValue)
}

// InitiallyUnownedStruct creates an uninitialised InitiallyUnowned.
func InitiallyUnownedStruct() *InitiallyUnowned {
	err := initiallyUnownedStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &InitiallyUnowned{native: initiallyUnownedStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeInitiallyUnowned)
	return structGo
}
func finalizeInitiallyUnowned(obj *InitiallyUnowned) {
	initiallyUnownedStruct.Free(obj.native)
}

var objectStruct *gi.Struct
var objectStruct_Once sync.Once

func objectStruct_Set() error {
	var err error
	objectStruct_Once.Do(func() {
		objectStruct, err = gi.StructNew("GObject", "Object")
	})
	return err
}

type Object struct {
	native uintptr
}

// FieldGTypeInstance returns the C field 'g_type_instance'.
func (recv *Object) FieldGTypeInstance() *TypeInstance {
	argValue := gi.FieldGet(objectStruct, recv.native, "g_type_instance")
	value := &TypeInstance{native: argValue.Pointer()}
	return value
}

// SetFieldGTypeInstance sets the value of the C field 'g_type_instance'.
func (recv *Object) SetFieldGTypeInstance(value *TypeInstance) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(objectStruct, recv.native, "g_type_instance", argValue)
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
		err = objectStruct_Set()
		if err != nil {
			return
		}
		objectForceFloatingFunction, err = objectStruct.InvokerNew("force_floating")
	})
	return err
}

// ForceFloating is a representation of the C type g_object_force_floating.
func (recv *Object) ForceFloating() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

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
		err = objectStruct_Set()
		if err != nil {
			return
		}
		objectFreezeNotifyFunction, err = objectStruct.InvokerNew("freeze_notify")
	})
	return err
}

// FreezeNotify is a representation of the C type g_object_freeze_notify.
func (recv *Object) FreezeNotify() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

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
		err = objectStruct_Set()
		if err != nil {
			return
		}
		objectGetPropertyFunction, err = objectStruct.InvokerNew("get_property")
	})
	return err
}

// GetProperty is a representation of the C type g_object_get_property.
func (recv *Object) GetProperty(propertyName string, value *Value) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(propertyName)
	inArgs[2].SetPointer(value.native)

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
		err = objectStruct_Set()
		if err != nil {
			return
		}
		objectIsFloatingFunction, err = objectStruct.InvokerNew("is_floating")
	})
	return err
}

// IsFloating is a representation of the C type g_object_is_floating.
func (recv *Object) IsFloating() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

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
		err = objectStruct_Set()
		if err != nil {
			return
		}
		objectNotifyFunction, err = objectStruct.InvokerNew("notify")
	})
	return err
}

// Notify is a representation of the C type g_object_notify.
func (recv *Object) Notify(propertyName string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
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
		err = objectStruct_Set()
		if err != nil {
			return
		}
		objectNotifyByPspecFunction, err = objectStruct.InvokerNew("notify_by_pspec")
	})
	return err
}

// NotifyByPspec is a representation of the C type g_object_notify_by_pspec.
func (recv *Object) NotifyByPspec(pspec *ParamSpec) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(pspec.native)

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
		err = objectStruct_Set()
		if err != nil {
			return
		}
		objectRefFunction, err = objectStruct.InvokerNew("ref")
	})
	return err
}

// Ref is a representation of the C type g_object_ref.
func (recv *Object) Ref() *Object {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := objectRefFunction_Set()
	if err == nil {
		ret = objectRefFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Object{native: ret.Pointer()}

	return retGo
}

var objectRefSinkFunction *gi.Function
var objectRefSinkFunction_Once sync.Once

func objectRefSinkFunction_Set() error {
	var err error
	objectRefSinkFunction_Once.Do(func() {
		err = objectStruct_Set()
		if err != nil {
			return
		}
		objectRefSinkFunction, err = objectStruct.InvokerNew("ref_sink")
	})
	return err
}

// RefSink is a representation of the C type g_object_ref_sink.
func (recv *Object) RefSink() *Object {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := objectRefSinkFunction_Set()
	if err == nil {
		ret = objectRefSinkFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Object{native: ret.Pointer()}

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
		err = objectStruct_Set()
		if err != nil {
			return
		}
		objectRunDisposeFunction, err = objectStruct.InvokerNew("run_dispose")
	})
	return err
}

// RunDispose is a representation of the C type g_object_run_dispose.
func (recv *Object) RunDispose() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

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
		err = objectStruct_Set()
		if err != nil {
			return
		}
		objectSetPropertyFunction, err = objectStruct.InvokerNew("set_property")
	})
	return err
}

// SetProperty is a representation of the C type g_object_set_property.
func (recv *Object) SetProperty(propertyName string, value *Value) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(propertyName)
	inArgs[2].SetPointer(value.native)

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
		err = objectStruct_Set()
		if err != nil {
			return
		}
		objectThawNotifyFunction, err = objectStruct.InvokerNew("thaw_notify")
	})
	return err
}

// ThawNotify is a representation of the C type g_object_thaw_notify.
func (recv *Object) ThawNotify() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

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
		err = objectStruct_Set()
		if err != nil {
			return
		}
		objectUnrefFunction, err = objectStruct.InvokerNew("unref")
	})
	return err
}

// Unref is a representation of the C type g_object_unref.
func (recv *Object) Unref() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

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
		err = objectStruct_Set()
		if err != nil {
			return
		}
		objectWatchClosureFunction, err = objectStruct.InvokerNew("watch_closure")
	})
	return err
}

// WatchClosure is a representation of the C type g_object_watch_closure.
func (recv *Object) WatchClosure(closure *Closure) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(closure.native)

	err := objectWatchClosureFunction_Set()
	if err == nil {
		objectWatchClosureFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'g_object_weak_ref' : parameter 'notify' of type 'WeakNotify' not supported

// UNSUPPORTED : C value 'g_object_weak_unref' : parameter 'notify' of type 'WeakNotify' not supported

var paramSpecStruct *gi.Struct
var paramSpecStruct_Once sync.Once

func paramSpecStruct_Set() error {
	var err error
	paramSpecStruct_Once.Do(func() {
		paramSpecStruct, err = gi.StructNew("GObject", "ParamSpec")
	})
	return err
}

type ParamSpec struct {
	native uintptr
}

// FieldGTypeInstance returns the C field 'g_type_instance'.
func (recv *ParamSpec) FieldGTypeInstance() *TypeInstance {
	argValue := gi.FieldGet(paramSpecStruct, recv.native, "g_type_instance")
	value := &TypeInstance{native: argValue.Pointer()}
	return value
}

// SetFieldGTypeInstance sets the value of the C field 'g_type_instance'.
func (recv *ParamSpec) SetFieldGTypeInstance(value *TypeInstance) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(paramSpecStruct, recv.native, "g_type_instance", argValue)
}

// FieldName returns the C field 'name'.
func (recv *ParamSpec) FieldName() string {
	argValue := gi.FieldGet(paramSpecStruct, recv.native, "name")
	value := argValue.String(false)
	return value
}

// SetFieldName sets the value of the C field 'name'.
func (recv *ParamSpec) SetFieldName(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.FieldSet(paramSpecStruct, recv.native, "name", argValue)
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
		err = paramSpecStruct_Set()
		if err != nil {
			return
		}
		paramSpecGetBlurbFunction, err = paramSpecStruct.InvokerNew("get_blurb")
	})
	return err
}

// GetBlurb is a representation of the C type g_param_spec_get_blurb.
func (recv *ParamSpec) GetBlurb() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

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
		err = paramSpecStruct_Set()
		if err != nil {
			return
		}
		paramSpecGetDefaultValueFunction, err = paramSpecStruct.InvokerNew("get_default_value")
	})
	return err
}

// GetDefaultValue is a representation of the C type g_param_spec_get_default_value.
func (recv *ParamSpec) GetDefaultValue() *Value {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := paramSpecGetDefaultValueFunction_Set()
	if err == nil {
		ret = paramSpecGetDefaultValueFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Value{native: ret.Pointer()}

	return retGo
}

var paramSpecGetNameFunction *gi.Function
var paramSpecGetNameFunction_Once sync.Once

func paramSpecGetNameFunction_Set() error {
	var err error
	paramSpecGetNameFunction_Once.Do(func() {
		err = paramSpecStruct_Set()
		if err != nil {
			return
		}
		paramSpecGetNameFunction, err = paramSpecStruct.InvokerNew("get_name")
	})
	return err
}

// GetName is a representation of the C type g_param_spec_get_name.
func (recv *ParamSpec) GetName() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

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
		err = paramSpecStruct_Set()
		if err != nil {
			return
		}
		paramSpecGetNameQuarkFunction, err = paramSpecStruct.InvokerNew("get_name_quark")
	})
	return err
}

// GetNameQuark is a representation of the C type g_param_spec_get_name_quark.
func (recv *ParamSpec) GetNameQuark() glib.Quark {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

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
		err = paramSpecStruct_Set()
		if err != nil {
			return
		}
		paramSpecGetNickFunction, err = paramSpecStruct.InvokerNew("get_nick")
	})
	return err
}

// GetNick is a representation of the C type g_param_spec_get_nick.
func (recv *ParamSpec) GetNick() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

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
		err = paramSpecStruct_Set()
		if err != nil {
			return
		}
		paramSpecGetRedirectTargetFunction, err = paramSpecStruct.InvokerNew("get_redirect_target")
	})
	return err
}

// GetRedirectTarget is a representation of the C type g_param_spec_get_redirect_target.
func (recv *ParamSpec) GetRedirectTarget() *ParamSpec {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := paramSpecGetRedirectTargetFunction_Set()
	if err == nil {
		ret = paramSpecGetRedirectTargetFunction.Invoke(inArgs[:], nil)
	}

	retGo := &ParamSpec{native: ret.Pointer()}

	return retGo
}

var paramSpecRefFunction *gi.Function
var paramSpecRefFunction_Once sync.Once

func paramSpecRefFunction_Set() error {
	var err error
	paramSpecRefFunction_Once.Do(func() {
		err = paramSpecStruct_Set()
		if err != nil {
			return
		}
		paramSpecRefFunction, err = paramSpecStruct.InvokerNew("ref")
	})
	return err
}

// Ref is a representation of the C type g_param_spec_ref.
func (recv *ParamSpec) Ref() *ParamSpec {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := paramSpecRefFunction_Set()
	if err == nil {
		ret = paramSpecRefFunction.Invoke(inArgs[:], nil)
	}

	retGo := &ParamSpec{native: ret.Pointer()}

	return retGo
}

var paramSpecRefSinkFunction *gi.Function
var paramSpecRefSinkFunction_Once sync.Once

func paramSpecRefSinkFunction_Set() error {
	var err error
	paramSpecRefSinkFunction_Once.Do(func() {
		err = paramSpecStruct_Set()
		if err != nil {
			return
		}
		paramSpecRefSinkFunction, err = paramSpecStruct.InvokerNew("ref_sink")
	})
	return err
}

// RefSink is a representation of the C type g_param_spec_ref_sink.
func (recv *ParamSpec) RefSink() *ParamSpec {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := paramSpecRefSinkFunction_Set()
	if err == nil {
		ret = paramSpecRefSinkFunction.Invoke(inArgs[:], nil)
	}

	retGo := &ParamSpec{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_param_spec_set_qdata' : parameter 'data' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_param_spec_set_qdata_full' : parameter 'data' of type 'gpointer' not supported

var paramSpecSinkFunction *gi.Function
var paramSpecSinkFunction_Once sync.Once

func paramSpecSinkFunction_Set() error {
	var err error
	paramSpecSinkFunction_Once.Do(func() {
		err = paramSpecStruct_Set()
		if err != nil {
			return
		}
		paramSpecSinkFunction, err = paramSpecStruct.InvokerNew("sink")
	})
	return err
}

// Sink is a representation of the C type g_param_spec_sink.
func (recv *ParamSpec) Sink() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

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
		err = paramSpecStruct_Set()
		if err != nil {
			return
		}
		paramSpecUnrefFunction, err = paramSpecStruct.InvokerNew("unref")
	})
	return err
}

// Unref is a representation of the C type g_param_spec_unref.
func (recv *ParamSpec) Unref() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := paramSpecUnrefFunction_Set()
	if err == nil {
		paramSpecUnrefFunction.Invoke(inArgs[:], nil)
	}

	return
}

// ParamSpecStruct creates an uninitialised ParamSpec.
func ParamSpecStruct() *ParamSpec {
	err := paramSpecStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ParamSpec{native: paramSpecStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeParamSpec)
	return structGo
}
func finalizeParamSpec(obj *ParamSpec) {
	paramSpecStruct.Free(obj.native)
}

var paramSpecBooleanStruct *gi.Struct
var paramSpecBooleanStruct_Once sync.Once

func paramSpecBooleanStruct_Set() error {
	var err error
	paramSpecBooleanStruct_Once.Do(func() {
		paramSpecBooleanStruct, err = gi.StructNew("GObject", "ParamSpecBoolean")
	})
	return err
}

type ParamSpecBoolean struct {
	native uintptr
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *ParamSpecBoolean) FieldParentInstance() *ParamSpec {
	argValue := gi.FieldGet(paramSpecBooleanStruct, recv.native, "parent_instance")
	value := &ParamSpec{native: argValue.Pointer()}
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *ParamSpecBoolean) SetFieldParentInstance(value *ParamSpec) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(paramSpecBooleanStruct, recv.native, "parent_instance", argValue)
}

// FieldDefaultValue returns the C field 'default_value'.
func (recv *ParamSpecBoolean) FieldDefaultValue() bool {
	argValue := gi.FieldGet(paramSpecBooleanStruct, recv.native, "default_value")
	value := argValue.Boolean()
	return value
}

// SetFieldDefaultValue sets the value of the C field 'default_value'.
func (recv *ParamSpecBoolean) SetFieldDefaultValue(value bool) {
	var argValue gi.Argument
	argValue.SetBoolean(value)
	gi.FieldSet(paramSpecBooleanStruct, recv.native, "default_value", argValue)
}

// ParamSpecBooleanStruct creates an uninitialised ParamSpecBoolean.
func ParamSpecBooleanStruct() *ParamSpecBoolean {
	err := paramSpecBooleanStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ParamSpecBoolean{native: paramSpecBooleanStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeParamSpecBoolean)
	return structGo
}
func finalizeParamSpecBoolean(obj *ParamSpecBoolean) {
	paramSpecBooleanStruct.Free(obj.native)
}

var paramSpecBoxedStruct *gi.Struct
var paramSpecBoxedStruct_Once sync.Once

func paramSpecBoxedStruct_Set() error {
	var err error
	paramSpecBoxedStruct_Once.Do(func() {
		paramSpecBoxedStruct, err = gi.StructNew("GObject", "ParamSpecBoxed")
	})
	return err
}

type ParamSpecBoxed struct {
	native uintptr
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *ParamSpecBoxed) FieldParentInstance() *ParamSpec {
	argValue := gi.FieldGet(paramSpecBoxedStruct, recv.native, "parent_instance")
	value := &ParamSpec{native: argValue.Pointer()}
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *ParamSpecBoxed) SetFieldParentInstance(value *ParamSpec) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(paramSpecBoxedStruct, recv.native, "parent_instance", argValue)
}

// ParamSpecBoxedStruct creates an uninitialised ParamSpecBoxed.
func ParamSpecBoxedStruct() *ParamSpecBoxed {
	err := paramSpecBoxedStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ParamSpecBoxed{native: paramSpecBoxedStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeParamSpecBoxed)
	return structGo
}
func finalizeParamSpecBoxed(obj *ParamSpecBoxed) {
	paramSpecBoxedStruct.Free(obj.native)
}

var paramSpecCharStruct *gi.Struct
var paramSpecCharStruct_Once sync.Once

func paramSpecCharStruct_Set() error {
	var err error
	paramSpecCharStruct_Once.Do(func() {
		paramSpecCharStruct, err = gi.StructNew("GObject", "ParamSpecChar")
	})
	return err
}

type ParamSpecChar struct {
	native uintptr
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *ParamSpecChar) FieldParentInstance() *ParamSpec {
	argValue := gi.FieldGet(paramSpecCharStruct, recv.native, "parent_instance")
	value := &ParamSpec{native: argValue.Pointer()}
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *ParamSpecChar) SetFieldParentInstance(value *ParamSpec) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(paramSpecCharStruct, recv.native, "parent_instance", argValue)
}

// FieldMinimum returns the C field 'minimum'.
func (recv *ParamSpecChar) FieldMinimum() int8 {
	argValue := gi.FieldGet(paramSpecCharStruct, recv.native, "minimum")
	value := argValue.Int8()
	return value
}

// SetFieldMinimum sets the value of the C field 'minimum'.
func (recv *ParamSpecChar) SetFieldMinimum(value int8) {
	var argValue gi.Argument
	argValue.SetInt8(value)
	gi.FieldSet(paramSpecCharStruct, recv.native, "minimum", argValue)
}

// FieldMaximum returns the C field 'maximum'.
func (recv *ParamSpecChar) FieldMaximum() int8 {
	argValue := gi.FieldGet(paramSpecCharStruct, recv.native, "maximum")
	value := argValue.Int8()
	return value
}

// SetFieldMaximum sets the value of the C field 'maximum'.
func (recv *ParamSpecChar) SetFieldMaximum(value int8) {
	var argValue gi.Argument
	argValue.SetInt8(value)
	gi.FieldSet(paramSpecCharStruct, recv.native, "maximum", argValue)
}

// FieldDefaultValue returns the C field 'default_value'.
func (recv *ParamSpecChar) FieldDefaultValue() int8 {
	argValue := gi.FieldGet(paramSpecCharStruct, recv.native, "default_value")
	value := argValue.Int8()
	return value
}

// SetFieldDefaultValue sets the value of the C field 'default_value'.
func (recv *ParamSpecChar) SetFieldDefaultValue(value int8) {
	var argValue gi.Argument
	argValue.SetInt8(value)
	gi.FieldSet(paramSpecCharStruct, recv.native, "default_value", argValue)
}

// ParamSpecCharStruct creates an uninitialised ParamSpecChar.
func ParamSpecCharStruct() *ParamSpecChar {
	err := paramSpecCharStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ParamSpecChar{native: paramSpecCharStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeParamSpecChar)
	return structGo
}
func finalizeParamSpecChar(obj *ParamSpecChar) {
	paramSpecCharStruct.Free(obj.native)
}

var paramSpecDoubleStruct *gi.Struct
var paramSpecDoubleStruct_Once sync.Once

func paramSpecDoubleStruct_Set() error {
	var err error
	paramSpecDoubleStruct_Once.Do(func() {
		paramSpecDoubleStruct, err = gi.StructNew("GObject", "ParamSpecDouble")
	})
	return err
}

type ParamSpecDouble struct {
	native uintptr
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *ParamSpecDouble) FieldParentInstance() *ParamSpec {
	argValue := gi.FieldGet(paramSpecDoubleStruct, recv.native, "parent_instance")
	value := &ParamSpec{native: argValue.Pointer()}
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *ParamSpecDouble) SetFieldParentInstance(value *ParamSpec) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(paramSpecDoubleStruct, recv.native, "parent_instance", argValue)
}

// FieldMinimum returns the C field 'minimum'.
func (recv *ParamSpecDouble) FieldMinimum() float64 {
	argValue := gi.FieldGet(paramSpecDoubleStruct, recv.native, "minimum")
	value := argValue.Float64()
	return value
}

// SetFieldMinimum sets the value of the C field 'minimum'.
func (recv *ParamSpecDouble) SetFieldMinimum(value float64) {
	var argValue gi.Argument
	argValue.SetFloat64(value)
	gi.FieldSet(paramSpecDoubleStruct, recv.native, "minimum", argValue)
}

// FieldMaximum returns the C field 'maximum'.
func (recv *ParamSpecDouble) FieldMaximum() float64 {
	argValue := gi.FieldGet(paramSpecDoubleStruct, recv.native, "maximum")
	value := argValue.Float64()
	return value
}

// SetFieldMaximum sets the value of the C field 'maximum'.
func (recv *ParamSpecDouble) SetFieldMaximum(value float64) {
	var argValue gi.Argument
	argValue.SetFloat64(value)
	gi.FieldSet(paramSpecDoubleStruct, recv.native, "maximum", argValue)
}

// FieldDefaultValue returns the C field 'default_value'.
func (recv *ParamSpecDouble) FieldDefaultValue() float64 {
	argValue := gi.FieldGet(paramSpecDoubleStruct, recv.native, "default_value")
	value := argValue.Float64()
	return value
}

// SetFieldDefaultValue sets the value of the C field 'default_value'.
func (recv *ParamSpecDouble) SetFieldDefaultValue(value float64) {
	var argValue gi.Argument
	argValue.SetFloat64(value)
	gi.FieldSet(paramSpecDoubleStruct, recv.native, "default_value", argValue)
}

// FieldEpsilon returns the C field 'epsilon'.
func (recv *ParamSpecDouble) FieldEpsilon() float64 {
	argValue := gi.FieldGet(paramSpecDoubleStruct, recv.native, "epsilon")
	value := argValue.Float64()
	return value
}

// SetFieldEpsilon sets the value of the C field 'epsilon'.
func (recv *ParamSpecDouble) SetFieldEpsilon(value float64) {
	var argValue gi.Argument
	argValue.SetFloat64(value)
	gi.FieldSet(paramSpecDoubleStruct, recv.native, "epsilon", argValue)
}

// ParamSpecDoubleStruct creates an uninitialised ParamSpecDouble.
func ParamSpecDoubleStruct() *ParamSpecDouble {
	err := paramSpecDoubleStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ParamSpecDouble{native: paramSpecDoubleStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeParamSpecDouble)
	return structGo
}
func finalizeParamSpecDouble(obj *ParamSpecDouble) {
	paramSpecDoubleStruct.Free(obj.native)
}

var paramSpecEnumStruct *gi.Struct
var paramSpecEnumStruct_Once sync.Once

func paramSpecEnumStruct_Set() error {
	var err error
	paramSpecEnumStruct_Once.Do(func() {
		paramSpecEnumStruct, err = gi.StructNew("GObject", "ParamSpecEnum")
	})
	return err
}

type ParamSpecEnum struct {
	native uintptr
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *ParamSpecEnum) FieldParentInstance() *ParamSpec {
	argValue := gi.FieldGet(paramSpecEnumStruct, recv.native, "parent_instance")
	value := &ParamSpec{native: argValue.Pointer()}
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *ParamSpecEnum) SetFieldParentInstance(value *ParamSpec) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(paramSpecEnumStruct, recv.native, "parent_instance", argValue)
}

// FieldEnumClass returns the C field 'enum_class'.
func (recv *ParamSpecEnum) FieldEnumClass() *EnumClass {
	argValue := gi.FieldGet(paramSpecEnumStruct, recv.native, "enum_class")
	value := &EnumClass{native: argValue.Pointer()}
	return value
}

// SetFieldEnumClass sets the value of the C field 'enum_class'.
func (recv *ParamSpecEnum) SetFieldEnumClass(value *EnumClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(paramSpecEnumStruct, recv.native, "enum_class", argValue)
}

// FieldDefaultValue returns the C field 'default_value'.
func (recv *ParamSpecEnum) FieldDefaultValue() int32 {
	argValue := gi.FieldGet(paramSpecEnumStruct, recv.native, "default_value")
	value := argValue.Int32()
	return value
}

// SetFieldDefaultValue sets the value of the C field 'default_value'.
func (recv *ParamSpecEnum) SetFieldDefaultValue(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.FieldSet(paramSpecEnumStruct, recv.native, "default_value", argValue)
}

// ParamSpecEnumStruct creates an uninitialised ParamSpecEnum.
func ParamSpecEnumStruct() *ParamSpecEnum {
	err := paramSpecEnumStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ParamSpecEnum{native: paramSpecEnumStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeParamSpecEnum)
	return structGo
}
func finalizeParamSpecEnum(obj *ParamSpecEnum) {
	paramSpecEnumStruct.Free(obj.native)
}

var paramSpecFlagsStruct *gi.Struct
var paramSpecFlagsStruct_Once sync.Once

func paramSpecFlagsStruct_Set() error {
	var err error
	paramSpecFlagsStruct_Once.Do(func() {
		paramSpecFlagsStruct, err = gi.StructNew("GObject", "ParamSpecFlags")
	})
	return err
}

type ParamSpecFlags struct {
	native uintptr
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *ParamSpecFlags) FieldParentInstance() *ParamSpec {
	argValue := gi.FieldGet(paramSpecFlagsStruct, recv.native, "parent_instance")
	value := &ParamSpec{native: argValue.Pointer()}
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *ParamSpecFlags) SetFieldParentInstance(value *ParamSpec) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(paramSpecFlagsStruct, recv.native, "parent_instance", argValue)
}

// FieldFlagsClass returns the C field 'flags_class'.
func (recv *ParamSpecFlags) FieldFlagsClass() *FlagsClass {
	argValue := gi.FieldGet(paramSpecFlagsStruct, recv.native, "flags_class")
	value := &FlagsClass{native: argValue.Pointer()}
	return value
}

// SetFieldFlagsClass sets the value of the C field 'flags_class'.
func (recv *ParamSpecFlags) SetFieldFlagsClass(value *FlagsClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(paramSpecFlagsStruct, recv.native, "flags_class", argValue)
}

// FieldDefaultValue returns the C field 'default_value'.
func (recv *ParamSpecFlags) FieldDefaultValue() uint32 {
	argValue := gi.FieldGet(paramSpecFlagsStruct, recv.native, "default_value")
	value := argValue.Uint32()
	return value
}

// SetFieldDefaultValue sets the value of the C field 'default_value'.
func (recv *ParamSpecFlags) SetFieldDefaultValue(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.FieldSet(paramSpecFlagsStruct, recv.native, "default_value", argValue)
}

// ParamSpecFlagsStruct creates an uninitialised ParamSpecFlags.
func ParamSpecFlagsStruct() *ParamSpecFlags {
	err := paramSpecFlagsStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ParamSpecFlags{native: paramSpecFlagsStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeParamSpecFlags)
	return structGo
}
func finalizeParamSpecFlags(obj *ParamSpecFlags) {
	paramSpecFlagsStruct.Free(obj.native)
}

var paramSpecFloatStruct *gi.Struct
var paramSpecFloatStruct_Once sync.Once

func paramSpecFloatStruct_Set() error {
	var err error
	paramSpecFloatStruct_Once.Do(func() {
		paramSpecFloatStruct, err = gi.StructNew("GObject", "ParamSpecFloat")
	})
	return err
}

type ParamSpecFloat struct {
	native uintptr
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *ParamSpecFloat) FieldParentInstance() *ParamSpec {
	argValue := gi.FieldGet(paramSpecFloatStruct, recv.native, "parent_instance")
	value := &ParamSpec{native: argValue.Pointer()}
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *ParamSpecFloat) SetFieldParentInstance(value *ParamSpec) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(paramSpecFloatStruct, recv.native, "parent_instance", argValue)
}

// FieldMinimum returns the C field 'minimum'.
func (recv *ParamSpecFloat) FieldMinimum() float32 {
	argValue := gi.FieldGet(paramSpecFloatStruct, recv.native, "minimum")
	value := argValue.Float32()
	return value
}

// SetFieldMinimum sets the value of the C field 'minimum'.
func (recv *ParamSpecFloat) SetFieldMinimum(value float32) {
	var argValue gi.Argument
	argValue.SetFloat32(value)
	gi.FieldSet(paramSpecFloatStruct, recv.native, "minimum", argValue)
}

// FieldMaximum returns the C field 'maximum'.
func (recv *ParamSpecFloat) FieldMaximum() float32 {
	argValue := gi.FieldGet(paramSpecFloatStruct, recv.native, "maximum")
	value := argValue.Float32()
	return value
}

// SetFieldMaximum sets the value of the C field 'maximum'.
func (recv *ParamSpecFloat) SetFieldMaximum(value float32) {
	var argValue gi.Argument
	argValue.SetFloat32(value)
	gi.FieldSet(paramSpecFloatStruct, recv.native, "maximum", argValue)
}

// FieldDefaultValue returns the C field 'default_value'.
func (recv *ParamSpecFloat) FieldDefaultValue() float32 {
	argValue := gi.FieldGet(paramSpecFloatStruct, recv.native, "default_value")
	value := argValue.Float32()
	return value
}

// SetFieldDefaultValue sets the value of the C field 'default_value'.
func (recv *ParamSpecFloat) SetFieldDefaultValue(value float32) {
	var argValue gi.Argument
	argValue.SetFloat32(value)
	gi.FieldSet(paramSpecFloatStruct, recv.native, "default_value", argValue)
}

// FieldEpsilon returns the C field 'epsilon'.
func (recv *ParamSpecFloat) FieldEpsilon() float32 {
	argValue := gi.FieldGet(paramSpecFloatStruct, recv.native, "epsilon")
	value := argValue.Float32()
	return value
}

// SetFieldEpsilon sets the value of the C field 'epsilon'.
func (recv *ParamSpecFloat) SetFieldEpsilon(value float32) {
	var argValue gi.Argument
	argValue.SetFloat32(value)
	gi.FieldSet(paramSpecFloatStruct, recv.native, "epsilon", argValue)
}

// ParamSpecFloatStruct creates an uninitialised ParamSpecFloat.
func ParamSpecFloatStruct() *ParamSpecFloat {
	err := paramSpecFloatStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ParamSpecFloat{native: paramSpecFloatStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeParamSpecFloat)
	return structGo
}
func finalizeParamSpecFloat(obj *ParamSpecFloat) {
	paramSpecFloatStruct.Free(obj.native)
}

var paramSpecGTypeStruct *gi.Struct
var paramSpecGTypeStruct_Once sync.Once

func paramSpecGTypeStruct_Set() error {
	var err error
	paramSpecGTypeStruct_Once.Do(func() {
		paramSpecGTypeStruct, err = gi.StructNew("GObject", "ParamSpecGType")
	})
	return err
}

type ParamSpecGType struct {
	native uintptr
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *ParamSpecGType) FieldParentInstance() *ParamSpec {
	argValue := gi.FieldGet(paramSpecGTypeStruct, recv.native, "parent_instance")
	value := &ParamSpec{native: argValue.Pointer()}
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *ParamSpecGType) SetFieldParentInstance(value *ParamSpec) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(paramSpecGTypeStruct, recv.native, "parent_instance", argValue)
}

// UNSUPPORTED : C value 'is_a_type' : for field getter : no Go type for 'GType'

// UNSUPPORTED : C value 'is_a_type' : for field setter : no Go type for 'GType'

// ParamSpecGTypeStruct creates an uninitialised ParamSpecGType.
func ParamSpecGTypeStruct() *ParamSpecGType {
	err := paramSpecGTypeStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ParamSpecGType{native: paramSpecGTypeStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeParamSpecGType)
	return structGo
}
func finalizeParamSpecGType(obj *ParamSpecGType) {
	paramSpecGTypeStruct.Free(obj.native)
}

var paramSpecIntStruct *gi.Struct
var paramSpecIntStruct_Once sync.Once

func paramSpecIntStruct_Set() error {
	var err error
	paramSpecIntStruct_Once.Do(func() {
		paramSpecIntStruct, err = gi.StructNew("GObject", "ParamSpecInt")
	})
	return err
}

type ParamSpecInt struct {
	native uintptr
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *ParamSpecInt) FieldParentInstance() *ParamSpec {
	argValue := gi.FieldGet(paramSpecIntStruct, recv.native, "parent_instance")
	value := &ParamSpec{native: argValue.Pointer()}
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *ParamSpecInt) SetFieldParentInstance(value *ParamSpec) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(paramSpecIntStruct, recv.native, "parent_instance", argValue)
}

// FieldMinimum returns the C field 'minimum'.
func (recv *ParamSpecInt) FieldMinimum() int32 {
	argValue := gi.FieldGet(paramSpecIntStruct, recv.native, "minimum")
	value := argValue.Int32()
	return value
}

// SetFieldMinimum sets the value of the C field 'minimum'.
func (recv *ParamSpecInt) SetFieldMinimum(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.FieldSet(paramSpecIntStruct, recv.native, "minimum", argValue)
}

// FieldMaximum returns the C field 'maximum'.
func (recv *ParamSpecInt) FieldMaximum() int32 {
	argValue := gi.FieldGet(paramSpecIntStruct, recv.native, "maximum")
	value := argValue.Int32()
	return value
}

// SetFieldMaximum sets the value of the C field 'maximum'.
func (recv *ParamSpecInt) SetFieldMaximum(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.FieldSet(paramSpecIntStruct, recv.native, "maximum", argValue)
}

// FieldDefaultValue returns the C field 'default_value'.
func (recv *ParamSpecInt) FieldDefaultValue() int32 {
	argValue := gi.FieldGet(paramSpecIntStruct, recv.native, "default_value")
	value := argValue.Int32()
	return value
}

// SetFieldDefaultValue sets the value of the C field 'default_value'.
func (recv *ParamSpecInt) SetFieldDefaultValue(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.FieldSet(paramSpecIntStruct, recv.native, "default_value", argValue)
}

// ParamSpecIntStruct creates an uninitialised ParamSpecInt.
func ParamSpecIntStruct() *ParamSpecInt {
	err := paramSpecIntStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ParamSpecInt{native: paramSpecIntStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeParamSpecInt)
	return structGo
}
func finalizeParamSpecInt(obj *ParamSpecInt) {
	paramSpecIntStruct.Free(obj.native)
}

var paramSpecInt64Struct *gi.Struct
var paramSpecInt64Struct_Once sync.Once

func paramSpecInt64Struct_Set() error {
	var err error
	paramSpecInt64Struct_Once.Do(func() {
		paramSpecInt64Struct, err = gi.StructNew("GObject", "ParamSpecInt64")
	})
	return err
}

type ParamSpecInt64 struct {
	native uintptr
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *ParamSpecInt64) FieldParentInstance() *ParamSpec {
	argValue := gi.FieldGet(paramSpecInt64Struct, recv.native, "parent_instance")
	value := &ParamSpec{native: argValue.Pointer()}
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *ParamSpecInt64) SetFieldParentInstance(value *ParamSpec) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(paramSpecInt64Struct, recv.native, "parent_instance", argValue)
}

// FieldMinimum returns the C field 'minimum'.
func (recv *ParamSpecInt64) FieldMinimum() int64 {
	argValue := gi.FieldGet(paramSpecInt64Struct, recv.native, "minimum")
	value := argValue.Int64()
	return value
}

// SetFieldMinimum sets the value of the C field 'minimum'.
func (recv *ParamSpecInt64) SetFieldMinimum(value int64) {
	var argValue gi.Argument
	argValue.SetInt64(value)
	gi.FieldSet(paramSpecInt64Struct, recv.native, "minimum", argValue)
}

// FieldMaximum returns the C field 'maximum'.
func (recv *ParamSpecInt64) FieldMaximum() int64 {
	argValue := gi.FieldGet(paramSpecInt64Struct, recv.native, "maximum")
	value := argValue.Int64()
	return value
}

// SetFieldMaximum sets the value of the C field 'maximum'.
func (recv *ParamSpecInt64) SetFieldMaximum(value int64) {
	var argValue gi.Argument
	argValue.SetInt64(value)
	gi.FieldSet(paramSpecInt64Struct, recv.native, "maximum", argValue)
}

// FieldDefaultValue returns the C field 'default_value'.
func (recv *ParamSpecInt64) FieldDefaultValue() int64 {
	argValue := gi.FieldGet(paramSpecInt64Struct, recv.native, "default_value")
	value := argValue.Int64()
	return value
}

// SetFieldDefaultValue sets the value of the C field 'default_value'.
func (recv *ParamSpecInt64) SetFieldDefaultValue(value int64) {
	var argValue gi.Argument
	argValue.SetInt64(value)
	gi.FieldSet(paramSpecInt64Struct, recv.native, "default_value", argValue)
}

// ParamSpecInt64Struct creates an uninitialised ParamSpecInt64.
func ParamSpecInt64Struct() *ParamSpecInt64 {
	err := paramSpecInt64Struct_Set()
	if err != nil {
		return nil
	}

	structGo := &ParamSpecInt64{native: paramSpecInt64Struct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeParamSpecInt64)
	return structGo
}
func finalizeParamSpecInt64(obj *ParamSpecInt64) {
	paramSpecInt64Struct.Free(obj.native)
}

var paramSpecLongStruct *gi.Struct
var paramSpecLongStruct_Once sync.Once

func paramSpecLongStruct_Set() error {
	var err error
	paramSpecLongStruct_Once.Do(func() {
		paramSpecLongStruct, err = gi.StructNew("GObject", "ParamSpecLong")
	})
	return err
}

type ParamSpecLong struct {
	native uintptr
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *ParamSpecLong) FieldParentInstance() *ParamSpec {
	argValue := gi.FieldGet(paramSpecLongStruct, recv.native, "parent_instance")
	value := &ParamSpec{native: argValue.Pointer()}
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *ParamSpecLong) SetFieldParentInstance(value *ParamSpec) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(paramSpecLongStruct, recv.native, "parent_instance", argValue)
}

// FieldMinimum returns the C field 'minimum'.
func (recv *ParamSpecLong) FieldMinimum() int64 {
	argValue := gi.FieldGet(paramSpecLongStruct, recv.native, "minimum")
	value := argValue.Int64()
	return value
}

// SetFieldMinimum sets the value of the C field 'minimum'.
func (recv *ParamSpecLong) SetFieldMinimum(value int64) {
	var argValue gi.Argument
	argValue.SetInt64(value)
	gi.FieldSet(paramSpecLongStruct, recv.native, "minimum", argValue)
}

// FieldMaximum returns the C field 'maximum'.
func (recv *ParamSpecLong) FieldMaximum() int64 {
	argValue := gi.FieldGet(paramSpecLongStruct, recv.native, "maximum")
	value := argValue.Int64()
	return value
}

// SetFieldMaximum sets the value of the C field 'maximum'.
func (recv *ParamSpecLong) SetFieldMaximum(value int64) {
	var argValue gi.Argument
	argValue.SetInt64(value)
	gi.FieldSet(paramSpecLongStruct, recv.native, "maximum", argValue)
}

// FieldDefaultValue returns the C field 'default_value'.
func (recv *ParamSpecLong) FieldDefaultValue() int64 {
	argValue := gi.FieldGet(paramSpecLongStruct, recv.native, "default_value")
	value := argValue.Int64()
	return value
}

// SetFieldDefaultValue sets the value of the C field 'default_value'.
func (recv *ParamSpecLong) SetFieldDefaultValue(value int64) {
	var argValue gi.Argument
	argValue.SetInt64(value)
	gi.FieldSet(paramSpecLongStruct, recv.native, "default_value", argValue)
}

// ParamSpecLongStruct creates an uninitialised ParamSpecLong.
func ParamSpecLongStruct() *ParamSpecLong {
	err := paramSpecLongStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ParamSpecLong{native: paramSpecLongStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeParamSpecLong)
	return structGo
}
func finalizeParamSpecLong(obj *ParamSpecLong) {
	paramSpecLongStruct.Free(obj.native)
}

var paramSpecObjectStruct *gi.Struct
var paramSpecObjectStruct_Once sync.Once

func paramSpecObjectStruct_Set() error {
	var err error
	paramSpecObjectStruct_Once.Do(func() {
		paramSpecObjectStruct, err = gi.StructNew("GObject", "ParamSpecObject")
	})
	return err
}

type ParamSpecObject struct {
	native uintptr
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *ParamSpecObject) FieldParentInstance() *ParamSpec {
	argValue := gi.FieldGet(paramSpecObjectStruct, recv.native, "parent_instance")
	value := &ParamSpec{native: argValue.Pointer()}
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *ParamSpecObject) SetFieldParentInstance(value *ParamSpec) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(paramSpecObjectStruct, recv.native, "parent_instance", argValue)
}

// ParamSpecObjectStruct creates an uninitialised ParamSpecObject.
func ParamSpecObjectStruct() *ParamSpecObject {
	err := paramSpecObjectStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ParamSpecObject{native: paramSpecObjectStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeParamSpecObject)
	return structGo
}
func finalizeParamSpecObject(obj *ParamSpecObject) {
	paramSpecObjectStruct.Free(obj.native)
}

var paramSpecOverrideStruct *gi.Struct
var paramSpecOverrideStruct_Once sync.Once

func paramSpecOverrideStruct_Set() error {
	var err error
	paramSpecOverrideStruct_Once.Do(func() {
		paramSpecOverrideStruct, err = gi.StructNew("GObject", "ParamSpecOverride")
	})
	return err
}

type ParamSpecOverride struct {
	native uintptr
}

// ParamSpecOverrideStruct creates an uninitialised ParamSpecOverride.
func ParamSpecOverrideStruct() *ParamSpecOverride {
	err := paramSpecOverrideStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ParamSpecOverride{native: paramSpecOverrideStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeParamSpecOverride)
	return structGo
}
func finalizeParamSpecOverride(obj *ParamSpecOverride) {
	paramSpecOverrideStruct.Free(obj.native)
}

var paramSpecParamStruct *gi.Struct
var paramSpecParamStruct_Once sync.Once

func paramSpecParamStruct_Set() error {
	var err error
	paramSpecParamStruct_Once.Do(func() {
		paramSpecParamStruct, err = gi.StructNew("GObject", "ParamSpecParam")
	})
	return err
}

type ParamSpecParam struct {
	native uintptr
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *ParamSpecParam) FieldParentInstance() *ParamSpec {
	argValue := gi.FieldGet(paramSpecParamStruct, recv.native, "parent_instance")
	value := &ParamSpec{native: argValue.Pointer()}
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *ParamSpecParam) SetFieldParentInstance(value *ParamSpec) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(paramSpecParamStruct, recv.native, "parent_instance", argValue)
}

// ParamSpecParamStruct creates an uninitialised ParamSpecParam.
func ParamSpecParamStruct() *ParamSpecParam {
	err := paramSpecParamStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ParamSpecParam{native: paramSpecParamStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeParamSpecParam)
	return structGo
}
func finalizeParamSpecParam(obj *ParamSpecParam) {
	paramSpecParamStruct.Free(obj.native)
}

var paramSpecPointerStruct *gi.Struct
var paramSpecPointerStruct_Once sync.Once

func paramSpecPointerStruct_Set() error {
	var err error
	paramSpecPointerStruct_Once.Do(func() {
		paramSpecPointerStruct, err = gi.StructNew("GObject", "ParamSpecPointer")
	})
	return err
}

type ParamSpecPointer struct {
	native uintptr
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *ParamSpecPointer) FieldParentInstance() *ParamSpec {
	argValue := gi.FieldGet(paramSpecPointerStruct, recv.native, "parent_instance")
	value := &ParamSpec{native: argValue.Pointer()}
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *ParamSpecPointer) SetFieldParentInstance(value *ParamSpec) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(paramSpecPointerStruct, recv.native, "parent_instance", argValue)
}

// ParamSpecPointerStruct creates an uninitialised ParamSpecPointer.
func ParamSpecPointerStruct() *ParamSpecPointer {
	err := paramSpecPointerStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ParamSpecPointer{native: paramSpecPointerStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeParamSpecPointer)
	return structGo
}
func finalizeParamSpecPointer(obj *ParamSpecPointer) {
	paramSpecPointerStruct.Free(obj.native)
}

var paramSpecStringStruct *gi.Struct
var paramSpecStringStruct_Once sync.Once

func paramSpecStringStruct_Set() error {
	var err error
	paramSpecStringStruct_Once.Do(func() {
		paramSpecStringStruct, err = gi.StructNew("GObject", "ParamSpecString")
	})
	return err
}

type ParamSpecString struct {
	native uintptr
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *ParamSpecString) FieldParentInstance() *ParamSpec {
	argValue := gi.FieldGet(paramSpecStringStruct, recv.native, "parent_instance")
	value := &ParamSpec{native: argValue.Pointer()}
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *ParamSpecString) SetFieldParentInstance(value *ParamSpec) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(paramSpecStringStruct, recv.native, "parent_instance", argValue)
}

// FieldDefaultValue returns the C field 'default_value'.
func (recv *ParamSpecString) FieldDefaultValue() string {
	argValue := gi.FieldGet(paramSpecStringStruct, recv.native, "default_value")
	value := argValue.String(false)
	return value
}

// SetFieldDefaultValue sets the value of the C field 'default_value'.
func (recv *ParamSpecString) SetFieldDefaultValue(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.FieldSet(paramSpecStringStruct, recv.native, "default_value", argValue)
}

// FieldCsetFirst returns the C field 'cset_first'.
func (recv *ParamSpecString) FieldCsetFirst() string {
	argValue := gi.FieldGet(paramSpecStringStruct, recv.native, "cset_first")
	value := argValue.String(false)
	return value
}

// SetFieldCsetFirst sets the value of the C field 'cset_first'.
func (recv *ParamSpecString) SetFieldCsetFirst(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.FieldSet(paramSpecStringStruct, recv.native, "cset_first", argValue)
}

// FieldCsetNth returns the C field 'cset_nth'.
func (recv *ParamSpecString) FieldCsetNth() string {
	argValue := gi.FieldGet(paramSpecStringStruct, recv.native, "cset_nth")
	value := argValue.String(false)
	return value
}

// SetFieldCsetNth sets the value of the C field 'cset_nth'.
func (recv *ParamSpecString) SetFieldCsetNth(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.FieldSet(paramSpecStringStruct, recv.native, "cset_nth", argValue)
}

// FieldSubstitutor returns the C field 'substitutor'.
func (recv *ParamSpecString) FieldSubstitutor() int8 {
	argValue := gi.FieldGet(paramSpecStringStruct, recv.native, "substitutor")
	value := argValue.Int8()
	return value
}

// SetFieldSubstitutor sets the value of the C field 'substitutor'.
func (recv *ParamSpecString) SetFieldSubstitutor(value int8) {
	var argValue gi.Argument
	argValue.SetInt8(value)
	gi.FieldSet(paramSpecStringStruct, recv.native, "substitutor", argValue)
}

// FieldNullFoldIfEmpty returns the C field 'null_fold_if_empty'.
func (recv *ParamSpecString) FieldNullFoldIfEmpty() uint32 {
	argValue := gi.FieldGet(paramSpecStringStruct, recv.native, "null_fold_if_empty")
	value := argValue.Uint32()
	return value
}

// SetFieldNullFoldIfEmpty sets the value of the C field 'null_fold_if_empty'.
func (recv *ParamSpecString) SetFieldNullFoldIfEmpty(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.FieldSet(paramSpecStringStruct, recv.native, "null_fold_if_empty", argValue)
}

// FieldEnsureNonNull returns the C field 'ensure_non_null'.
func (recv *ParamSpecString) FieldEnsureNonNull() uint32 {
	argValue := gi.FieldGet(paramSpecStringStruct, recv.native, "ensure_non_null")
	value := argValue.Uint32()
	return value
}

// SetFieldEnsureNonNull sets the value of the C field 'ensure_non_null'.
func (recv *ParamSpecString) SetFieldEnsureNonNull(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.FieldSet(paramSpecStringStruct, recv.native, "ensure_non_null", argValue)
}

// ParamSpecStringStruct creates an uninitialised ParamSpecString.
func ParamSpecStringStruct() *ParamSpecString {
	err := paramSpecStringStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ParamSpecString{native: paramSpecStringStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeParamSpecString)
	return structGo
}
func finalizeParamSpecString(obj *ParamSpecString) {
	paramSpecStringStruct.Free(obj.native)
}

var paramSpecUCharStruct *gi.Struct
var paramSpecUCharStruct_Once sync.Once

func paramSpecUCharStruct_Set() error {
	var err error
	paramSpecUCharStruct_Once.Do(func() {
		paramSpecUCharStruct, err = gi.StructNew("GObject", "ParamSpecUChar")
	})
	return err
}

type ParamSpecUChar struct {
	native uintptr
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *ParamSpecUChar) FieldParentInstance() *ParamSpec {
	argValue := gi.FieldGet(paramSpecUCharStruct, recv.native, "parent_instance")
	value := &ParamSpec{native: argValue.Pointer()}
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *ParamSpecUChar) SetFieldParentInstance(value *ParamSpec) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(paramSpecUCharStruct, recv.native, "parent_instance", argValue)
}

// FieldMinimum returns the C field 'minimum'.
func (recv *ParamSpecUChar) FieldMinimum() uint8 {
	argValue := gi.FieldGet(paramSpecUCharStruct, recv.native, "minimum")
	value := argValue.Uint8()
	return value
}

// SetFieldMinimum sets the value of the C field 'minimum'.
func (recv *ParamSpecUChar) SetFieldMinimum(value uint8) {
	var argValue gi.Argument
	argValue.SetUint8(value)
	gi.FieldSet(paramSpecUCharStruct, recv.native, "minimum", argValue)
}

// FieldMaximum returns the C field 'maximum'.
func (recv *ParamSpecUChar) FieldMaximum() uint8 {
	argValue := gi.FieldGet(paramSpecUCharStruct, recv.native, "maximum")
	value := argValue.Uint8()
	return value
}

// SetFieldMaximum sets the value of the C field 'maximum'.
func (recv *ParamSpecUChar) SetFieldMaximum(value uint8) {
	var argValue gi.Argument
	argValue.SetUint8(value)
	gi.FieldSet(paramSpecUCharStruct, recv.native, "maximum", argValue)
}

// FieldDefaultValue returns the C field 'default_value'.
func (recv *ParamSpecUChar) FieldDefaultValue() uint8 {
	argValue := gi.FieldGet(paramSpecUCharStruct, recv.native, "default_value")
	value := argValue.Uint8()
	return value
}

// SetFieldDefaultValue sets the value of the C field 'default_value'.
func (recv *ParamSpecUChar) SetFieldDefaultValue(value uint8) {
	var argValue gi.Argument
	argValue.SetUint8(value)
	gi.FieldSet(paramSpecUCharStruct, recv.native, "default_value", argValue)
}

// ParamSpecUCharStruct creates an uninitialised ParamSpecUChar.
func ParamSpecUCharStruct() *ParamSpecUChar {
	err := paramSpecUCharStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ParamSpecUChar{native: paramSpecUCharStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeParamSpecUChar)
	return structGo
}
func finalizeParamSpecUChar(obj *ParamSpecUChar) {
	paramSpecUCharStruct.Free(obj.native)
}

var paramSpecUIntStruct *gi.Struct
var paramSpecUIntStruct_Once sync.Once

func paramSpecUIntStruct_Set() error {
	var err error
	paramSpecUIntStruct_Once.Do(func() {
		paramSpecUIntStruct, err = gi.StructNew("GObject", "ParamSpecUInt")
	})
	return err
}

type ParamSpecUInt struct {
	native uintptr
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *ParamSpecUInt) FieldParentInstance() *ParamSpec {
	argValue := gi.FieldGet(paramSpecUIntStruct, recv.native, "parent_instance")
	value := &ParamSpec{native: argValue.Pointer()}
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *ParamSpecUInt) SetFieldParentInstance(value *ParamSpec) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(paramSpecUIntStruct, recv.native, "parent_instance", argValue)
}

// FieldMinimum returns the C field 'minimum'.
func (recv *ParamSpecUInt) FieldMinimum() uint32 {
	argValue := gi.FieldGet(paramSpecUIntStruct, recv.native, "minimum")
	value := argValue.Uint32()
	return value
}

// SetFieldMinimum sets the value of the C field 'minimum'.
func (recv *ParamSpecUInt) SetFieldMinimum(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.FieldSet(paramSpecUIntStruct, recv.native, "minimum", argValue)
}

// FieldMaximum returns the C field 'maximum'.
func (recv *ParamSpecUInt) FieldMaximum() uint32 {
	argValue := gi.FieldGet(paramSpecUIntStruct, recv.native, "maximum")
	value := argValue.Uint32()
	return value
}

// SetFieldMaximum sets the value of the C field 'maximum'.
func (recv *ParamSpecUInt) SetFieldMaximum(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.FieldSet(paramSpecUIntStruct, recv.native, "maximum", argValue)
}

// FieldDefaultValue returns the C field 'default_value'.
func (recv *ParamSpecUInt) FieldDefaultValue() uint32 {
	argValue := gi.FieldGet(paramSpecUIntStruct, recv.native, "default_value")
	value := argValue.Uint32()
	return value
}

// SetFieldDefaultValue sets the value of the C field 'default_value'.
func (recv *ParamSpecUInt) SetFieldDefaultValue(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.FieldSet(paramSpecUIntStruct, recv.native, "default_value", argValue)
}

// ParamSpecUIntStruct creates an uninitialised ParamSpecUInt.
func ParamSpecUIntStruct() *ParamSpecUInt {
	err := paramSpecUIntStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ParamSpecUInt{native: paramSpecUIntStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeParamSpecUInt)
	return structGo
}
func finalizeParamSpecUInt(obj *ParamSpecUInt) {
	paramSpecUIntStruct.Free(obj.native)
}

var paramSpecUInt64Struct *gi.Struct
var paramSpecUInt64Struct_Once sync.Once

func paramSpecUInt64Struct_Set() error {
	var err error
	paramSpecUInt64Struct_Once.Do(func() {
		paramSpecUInt64Struct, err = gi.StructNew("GObject", "ParamSpecUInt64")
	})
	return err
}

type ParamSpecUInt64 struct {
	native uintptr
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *ParamSpecUInt64) FieldParentInstance() *ParamSpec {
	argValue := gi.FieldGet(paramSpecUInt64Struct, recv.native, "parent_instance")
	value := &ParamSpec{native: argValue.Pointer()}
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *ParamSpecUInt64) SetFieldParentInstance(value *ParamSpec) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(paramSpecUInt64Struct, recv.native, "parent_instance", argValue)
}

// FieldMinimum returns the C field 'minimum'.
func (recv *ParamSpecUInt64) FieldMinimum() uint64 {
	argValue := gi.FieldGet(paramSpecUInt64Struct, recv.native, "minimum")
	value := argValue.Uint64()
	return value
}

// SetFieldMinimum sets the value of the C field 'minimum'.
func (recv *ParamSpecUInt64) SetFieldMinimum(value uint64) {
	var argValue gi.Argument
	argValue.SetUint64(value)
	gi.FieldSet(paramSpecUInt64Struct, recv.native, "minimum", argValue)
}

// FieldMaximum returns the C field 'maximum'.
func (recv *ParamSpecUInt64) FieldMaximum() uint64 {
	argValue := gi.FieldGet(paramSpecUInt64Struct, recv.native, "maximum")
	value := argValue.Uint64()
	return value
}

// SetFieldMaximum sets the value of the C field 'maximum'.
func (recv *ParamSpecUInt64) SetFieldMaximum(value uint64) {
	var argValue gi.Argument
	argValue.SetUint64(value)
	gi.FieldSet(paramSpecUInt64Struct, recv.native, "maximum", argValue)
}

// FieldDefaultValue returns the C field 'default_value'.
func (recv *ParamSpecUInt64) FieldDefaultValue() uint64 {
	argValue := gi.FieldGet(paramSpecUInt64Struct, recv.native, "default_value")
	value := argValue.Uint64()
	return value
}

// SetFieldDefaultValue sets the value of the C field 'default_value'.
func (recv *ParamSpecUInt64) SetFieldDefaultValue(value uint64) {
	var argValue gi.Argument
	argValue.SetUint64(value)
	gi.FieldSet(paramSpecUInt64Struct, recv.native, "default_value", argValue)
}

// ParamSpecUInt64Struct creates an uninitialised ParamSpecUInt64.
func ParamSpecUInt64Struct() *ParamSpecUInt64 {
	err := paramSpecUInt64Struct_Set()
	if err != nil {
		return nil
	}

	structGo := &ParamSpecUInt64{native: paramSpecUInt64Struct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeParamSpecUInt64)
	return structGo
}
func finalizeParamSpecUInt64(obj *ParamSpecUInt64) {
	paramSpecUInt64Struct.Free(obj.native)
}

var paramSpecULongStruct *gi.Struct
var paramSpecULongStruct_Once sync.Once

func paramSpecULongStruct_Set() error {
	var err error
	paramSpecULongStruct_Once.Do(func() {
		paramSpecULongStruct, err = gi.StructNew("GObject", "ParamSpecULong")
	})
	return err
}

type ParamSpecULong struct {
	native uintptr
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *ParamSpecULong) FieldParentInstance() *ParamSpec {
	argValue := gi.FieldGet(paramSpecULongStruct, recv.native, "parent_instance")
	value := &ParamSpec{native: argValue.Pointer()}
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *ParamSpecULong) SetFieldParentInstance(value *ParamSpec) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(paramSpecULongStruct, recv.native, "parent_instance", argValue)
}

// FieldMinimum returns the C field 'minimum'.
func (recv *ParamSpecULong) FieldMinimum() uint64 {
	argValue := gi.FieldGet(paramSpecULongStruct, recv.native, "minimum")
	value := argValue.Uint64()
	return value
}

// SetFieldMinimum sets the value of the C field 'minimum'.
func (recv *ParamSpecULong) SetFieldMinimum(value uint64) {
	var argValue gi.Argument
	argValue.SetUint64(value)
	gi.FieldSet(paramSpecULongStruct, recv.native, "minimum", argValue)
}

// FieldMaximum returns the C field 'maximum'.
func (recv *ParamSpecULong) FieldMaximum() uint64 {
	argValue := gi.FieldGet(paramSpecULongStruct, recv.native, "maximum")
	value := argValue.Uint64()
	return value
}

// SetFieldMaximum sets the value of the C field 'maximum'.
func (recv *ParamSpecULong) SetFieldMaximum(value uint64) {
	var argValue gi.Argument
	argValue.SetUint64(value)
	gi.FieldSet(paramSpecULongStruct, recv.native, "maximum", argValue)
}

// FieldDefaultValue returns the C field 'default_value'.
func (recv *ParamSpecULong) FieldDefaultValue() uint64 {
	argValue := gi.FieldGet(paramSpecULongStruct, recv.native, "default_value")
	value := argValue.Uint64()
	return value
}

// SetFieldDefaultValue sets the value of the C field 'default_value'.
func (recv *ParamSpecULong) SetFieldDefaultValue(value uint64) {
	var argValue gi.Argument
	argValue.SetUint64(value)
	gi.FieldSet(paramSpecULongStruct, recv.native, "default_value", argValue)
}

// ParamSpecULongStruct creates an uninitialised ParamSpecULong.
func ParamSpecULongStruct() *ParamSpecULong {
	err := paramSpecULongStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ParamSpecULong{native: paramSpecULongStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeParamSpecULong)
	return structGo
}
func finalizeParamSpecULong(obj *ParamSpecULong) {
	paramSpecULongStruct.Free(obj.native)
}

var paramSpecUnicharStruct *gi.Struct
var paramSpecUnicharStruct_Once sync.Once

func paramSpecUnicharStruct_Set() error {
	var err error
	paramSpecUnicharStruct_Once.Do(func() {
		paramSpecUnicharStruct, err = gi.StructNew("GObject", "ParamSpecUnichar")
	})
	return err
}

type ParamSpecUnichar struct {
	native uintptr
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *ParamSpecUnichar) FieldParentInstance() *ParamSpec {
	argValue := gi.FieldGet(paramSpecUnicharStruct, recv.native, "parent_instance")
	value := &ParamSpec{native: argValue.Pointer()}
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *ParamSpecUnichar) SetFieldParentInstance(value *ParamSpec) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(paramSpecUnicharStruct, recv.native, "parent_instance", argValue)
}

// UNSUPPORTED : C value 'default_value' : for field getter : no Go type for 'gunichar'

// UNSUPPORTED : C value 'default_value' : for field setter : no Go type for 'gunichar'

// ParamSpecUnicharStruct creates an uninitialised ParamSpecUnichar.
func ParamSpecUnicharStruct() *ParamSpecUnichar {
	err := paramSpecUnicharStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ParamSpecUnichar{native: paramSpecUnicharStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeParamSpecUnichar)
	return structGo
}
func finalizeParamSpecUnichar(obj *ParamSpecUnichar) {
	paramSpecUnicharStruct.Free(obj.native)
}

var paramSpecValueArrayStruct *gi.Struct
var paramSpecValueArrayStruct_Once sync.Once

func paramSpecValueArrayStruct_Set() error {
	var err error
	paramSpecValueArrayStruct_Once.Do(func() {
		paramSpecValueArrayStruct, err = gi.StructNew("GObject", "ParamSpecValueArray")
	})
	return err
}

type ParamSpecValueArray struct {
	native uintptr
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *ParamSpecValueArray) FieldParentInstance() *ParamSpec {
	argValue := gi.FieldGet(paramSpecValueArrayStruct, recv.native, "parent_instance")
	value := &ParamSpec{native: argValue.Pointer()}
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *ParamSpecValueArray) SetFieldParentInstance(value *ParamSpec) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(paramSpecValueArrayStruct, recv.native, "parent_instance", argValue)
}

// FieldElementSpec returns the C field 'element_spec'.
func (recv *ParamSpecValueArray) FieldElementSpec() *ParamSpec {
	argValue := gi.FieldGet(paramSpecValueArrayStruct, recv.native, "element_spec")
	value := &ParamSpec{native: argValue.Pointer()}
	return value
}

// SetFieldElementSpec sets the value of the C field 'element_spec'.
func (recv *ParamSpecValueArray) SetFieldElementSpec(value *ParamSpec) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(paramSpecValueArrayStruct, recv.native, "element_spec", argValue)
}

// FieldFixedNElements returns the C field 'fixed_n_elements'.
func (recv *ParamSpecValueArray) FieldFixedNElements() uint32 {
	argValue := gi.FieldGet(paramSpecValueArrayStruct, recv.native, "fixed_n_elements")
	value := argValue.Uint32()
	return value
}

// SetFieldFixedNElements sets the value of the C field 'fixed_n_elements'.
func (recv *ParamSpecValueArray) SetFieldFixedNElements(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.FieldSet(paramSpecValueArrayStruct, recv.native, "fixed_n_elements", argValue)
}

// ParamSpecValueArrayStruct creates an uninitialised ParamSpecValueArray.
func ParamSpecValueArrayStruct() *ParamSpecValueArray {
	err := paramSpecValueArrayStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ParamSpecValueArray{native: paramSpecValueArrayStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeParamSpecValueArray)
	return structGo
}
func finalizeParamSpecValueArray(obj *ParamSpecValueArray) {
	paramSpecValueArrayStruct.Free(obj.native)
}

var paramSpecVariantStruct *gi.Struct
var paramSpecVariantStruct_Once sync.Once

func paramSpecVariantStruct_Set() error {
	var err error
	paramSpecVariantStruct_Once.Do(func() {
		paramSpecVariantStruct, err = gi.StructNew("GObject", "ParamSpecVariant")
	})
	return err
}

type ParamSpecVariant struct {
	native uintptr
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *ParamSpecVariant) FieldParentInstance() *ParamSpec {
	argValue := gi.FieldGet(paramSpecVariantStruct, recv.native, "parent_instance")
	value := &ParamSpec{native: argValue.Pointer()}
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *ParamSpecVariant) SetFieldParentInstance(value *ParamSpec) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(paramSpecVariantStruct, recv.native, "parent_instance", argValue)
}

// UNSUPPORTED : C value 'type' : for field getter : no Go type for 'GLib.VariantType'

// UNSUPPORTED : C value 'type' : for field setter : no Go type for 'GLib.VariantType'

// UNSUPPORTED : C value 'default_value' : for field getter : no Go type for 'GLib.Variant'

// UNSUPPORTED : C value 'default_value' : for field setter : no Go type for 'GLib.Variant'

// ParamSpecVariantStruct creates an uninitialised ParamSpecVariant.
func ParamSpecVariantStruct() *ParamSpecVariant {
	err := paramSpecVariantStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ParamSpecVariant{native: paramSpecVariantStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeParamSpecVariant)
	return structGo
}
func finalizeParamSpecVariant(obj *ParamSpecVariant) {
	paramSpecVariantStruct.Free(obj.native)
}

var typeModuleStruct *gi.Struct
var typeModuleStruct_Once sync.Once

func typeModuleStruct_Set() error {
	var err error
	typeModuleStruct_Once.Do(func() {
		typeModuleStruct, err = gi.StructNew("GObject", "TypeModule")
	})
	return err
}

type TypeModule struct {
	native uintptr
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *TypeModule) FieldParentInstance() *Object {
	argValue := gi.FieldGet(typeModuleStruct, recv.native, "parent_instance")
	value := &Object{native: argValue.Pointer()}
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *TypeModule) SetFieldParentInstance(value *Object) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(typeModuleStruct, recv.native, "parent_instance", argValue)
}

// FieldUseCount returns the C field 'use_count'.
func (recv *TypeModule) FieldUseCount() uint32 {
	argValue := gi.FieldGet(typeModuleStruct, recv.native, "use_count")
	value := argValue.Uint32()
	return value
}

// SetFieldUseCount sets the value of the C field 'use_count'.
func (recv *TypeModule) SetFieldUseCount(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.FieldSet(typeModuleStruct, recv.native, "use_count", argValue)
}

// UNSUPPORTED : C value 'type_infos' : for field getter : no Go type for 'GLib.SList'

// UNSUPPORTED : C value 'type_infos' : for field setter : no Go type for 'GLib.SList'

// UNSUPPORTED : C value 'interface_infos' : for field getter : no Go type for 'GLib.SList'

// UNSUPPORTED : C value 'interface_infos' : for field setter : no Go type for 'GLib.SList'

// FieldName returns the C field 'name'.
func (recv *TypeModule) FieldName() string {
	argValue := gi.FieldGet(typeModuleStruct, recv.native, "name")
	value := argValue.String(false)
	return value
}

// SetFieldName sets the value of the C field 'name'.
func (recv *TypeModule) SetFieldName(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.FieldSet(typeModuleStruct, recv.native, "name", argValue)
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
		err = typeModuleStruct_Set()
		if err != nil {
			return
		}
		typeModuleSetNameFunction, err = typeModuleStruct.InvokerNew("set_name")
	})
	return err
}

// SetName is a representation of the C type g_type_module_set_name.
func (recv *TypeModule) SetName(name string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
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
		err = typeModuleStruct_Set()
		if err != nil {
			return
		}
		typeModuleUnuseFunction, err = typeModuleStruct.InvokerNew("unuse")
	})
	return err
}

// Unuse is a representation of the C type g_type_module_unuse.
func (recv *TypeModule) Unuse() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

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
		err = typeModuleStruct_Set()
		if err != nil {
			return
		}
		typeModuleUseFunction, err = typeModuleStruct.InvokerNew("use")
	})
	return err
}

// Use is a representation of the C type g_type_module_use.
func (recv *TypeModule) Use() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := typeModuleUseFunction_Set()
	if err == nil {
		ret = typeModuleUseFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// TypeModuleStruct creates an uninitialised TypeModule.
func TypeModuleStruct() *TypeModule {
	err := typeModuleStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &TypeModule{native: typeModuleStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeTypeModule)
	return structGo
}
func finalizeTypeModule(obj *TypeModule) {
	typeModuleStruct.Free(obj.native)
}
