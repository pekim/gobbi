// Code generated - DO NOT EDIT.

package atk

import (
	gi "github.com/pekim/gobbi/internal/gi"
	"sync"
)

// UNSUPPORTED : C value 'atk_add_focus_tracker' : parameter 'focus_tracker' of type 'EventListener' not supported

// UNSUPPORTED : C value 'atk_add_global_event_listener' : parameter 'listener' of type 'GObject.SignalEmissionHook' not supported

// UNSUPPORTED : C value 'atk_add_key_event_listener' : parameter 'listener' of type 'KeySnoopFunc' not supported

// UNSUPPORTED : C value 'atk_attribute_set_free' : parameter 'attrib_set' of type 'AttributeSet' not supported

// UNSUPPORTED : C value 'atk_focus_tracker_init' : parameter 'init' of type 'EventListenerInit' not supported

var focusTrackerNotifyFunction *gi.Function
var focusTrackerNotifyFunction_Once sync.Once

func focusTrackerNotifyFunction_Set() error {
	var err error
	focusTrackerNotifyFunction_Once.Do(func() {
		focusTrackerNotifyFunction, err = gi.FunctionInvokerNew("Atk", "focus_tracker_notify")
	})
	return err
}

// FocusTrackerNotify is a representation of the C type atk_focus_tracker_notify.
func FocusTrackerNotify(object *Object) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(object.native)

	err := focusTrackerNotifyFunction_Set()
	if err == nil {
		focusTrackerNotifyFunction.Invoke(inArgs[:], nil)
	}

	return
}

var getBinaryAgeFunction *gi.Function
var getBinaryAgeFunction_Once sync.Once

func getBinaryAgeFunction_Set() error {
	var err error
	getBinaryAgeFunction_Once.Do(func() {
		getBinaryAgeFunction, err = gi.FunctionInvokerNew("Atk", "get_binary_age")
	})
	return err
}

// GetBinaryAge is a representation of the C type atk_get_binary_age.
func GetBinaryAge() uint32 {

	var ret gi.Argument

	err := getBinaryAgeFunction_Set()
	if err == nil {
		ret = getBinaryAgeFunction.Invoke(nil, nil)
	}

	retGo := ret.Uint32()

	return retGo
}

var getDefaultRegistryFunction *gi.Function
var getDefaultRegistryFunction_Once sync.Once

func getDefaultRegistryFunction_Set() error {
	var err error
	getDefaultRegistryFunction_Once.Do(func() {
		getDefaultRegistryFunction, err = gi.FunctionInvokerNew("Atk", "get_default_registry")
	})
	return err
}

// GetDefaultRegistry is a representation of the C type atk_get_default_registry.
func GetDefaultRegistry() *Registry {

	var ret gi.Argument

	err := getDefaultRegistryFunction_Set()
	if err == nil {
		ret = getDefaultRegistryFunction.Invoke(nil, nil)
	}

	retGo := &Registry{native: ret.Pointer()}

	return retGo
}

var getFocusObjectFunction *gi.Function
var getFocusObjectFunction_Once sync.Once

func getFocusObjectFunction_Set() error {
	var err error
	getFocusObjectFunction_Once.Do(func() {
		getFocusObjectFunction, err = gi.FunctionInvokerNew("Atk", "get_focus_object")
	})
	return err
}

// GetFocusObject is a representation of the C type atk_get_focus_object.
func GetFocusObject() *Object {

	var ret gi.Argument

	err := getFocusObjectFunction_Set()
	if err == nil {
		ret = getFocusObjectFunction.Invoke(nil, nil)
	}

	retGo := &Object{native: ret.Pointer()}

	return retGo
}

var getInterfaceAgeFunction *gi.Function
var getInterfaceAgeFunction_Once sync.Once

func getInterfaceAgeFunction_Set() error {
	var err error
	getInterfaceAgeFunction_Once.Do(func() {
		getInterfaceAgeFunction, err = gi.FunctionInvokerNew("Atk", "get_interface_age")
	})
	return err
}

// GetInterfaceAge is a representation of the C type atk_get_interface_age.
func GetInterfaceAge() uint32 {

	var ret gi.Argument

	err := getInterfaceAgeFunction_Set()
	if err == nil {
		ret = getInterfaceAgeFunction.Invoke(nil, nil)
	}

	retGo := ret.Uint32()

	return retGo
}

var getMajorVersionFunction *gi.Function
var getMajorVersionFunction_Once sync.Once

func getMajorVersionFunction_Set() error {
	var err error
	getMajorVersionFunction_Once.Do(func() {
		getMajorVersionFunction, err = gi.FunctionInvokerNew("Atk", "get_major_version")
	})
	return err
}

// GetMajorVersion is a representation of the C type atk_get_major_version.
func GetMajorVersion() uint32 {

	var ret gi.Argument

	err := getMajorVersionFunction_Set()
	if err == nil {
		ret = getMajorVersionFunction.Invoke(nil, nil)
	}

	retGo := ret.Uint32()

	return retGo
}

var getMicroVersionFunction *gi.Function
var getMicroVersionFunction_Once sync.Once

func getMicroVersionFunction_Set() error {
	var err error
	getMicroVersionFunction_Once.Do(func() {
		getMicroVersionFunction, err = gi.FunctionInvokerNew("Atk", "get_micro_version")
	})
	return err
}

// GetMicroVersion is a representation of the C type atk_get_micro_version.
func GetMicroVersion() uint32 {

	var ret gi.Argument

	err := getMicroVersionFunction_Set()
	if err == nil {
		ret = getMicroVersionFunction.Invoke(nil, nil)
	}

	retGo := ret.Uint32()

	return retGo
}

var getMinorVersionFunction *gi.Function
var getMinorVersionFunction_Once sync.Once

func getMinorVersionFunction_Set() error {
	var err error
	getMinorVersionFunction_Once.Do(func() {
		getMinorVersionFunction, err = gi.FunctionInvokerNew("Atk", "get_minor_version")
	})
	return err
}

// GetMinorVersion is a representation of the C type atk_get_minor_version.
func GetMinorVersion() uint32 {

	var ret gi.Argument

	err := getMinorVersionFunction_Set()
	if err == nil {
		ret = getMinorVersionFunction.Invoke(nil, nil)
	}

	retGo := ret.Uint32()

	return retGo
}

var getRootFunction *gi.Function
var getRootFunction_Once sync.Once

func getRootFunction_Set() error {
	var err error
	getRootFunction_Once.Do(func() {
		getRootFunction, err = gi.FunctionInvokerNew("Atk", "get_root")
	})
	return err
}

// GetRoot is a representation of the C type atk_get_root.
func GetRoot() *Object {

	var ret gi.Argument

	err := getRootFunction_Set()
	if err == nil {
		ret = getRootFunction.Invoke(nil, nil)
	}

	retGo := &Object{native: ret.Pointer()}

	return retGo
}

var getToolkitNameFunction *gi.Function
var getToolkitNameFunction_Once sync.Once

func getToolkitNameFunction_Set() error {
	var err error
	getToolkitNameFunction_Once.Do(func() {
		getToolkitNameFunction, err = gi.FunctionInvokerNew("Atk", "get_toolkit_name")
	})
	return err
}

// GetToolkitName is a representation of the C type atk_get_toolkit_name.
func GetToolkitName() string {

	var ret gi.Argument

	err := getToolkitNameFunction_Set()
	if err == nil {
		ret = getToolkitNameFunction.Invoke(nil, nil)
	}

	retGo := ret.String(false)

	return retGo
}

var getToolkitVersionFunction *gi.Function
var getToolkitVersionFunction_Once sync.Once

func getToolkitVersionFunction_Set() error {
	var err error
	getToolkitVersionFunction_Once.Do(func() {
		getToolkitVersionFunction, err = gi.FunctionInvokerNew("Atk", "get_toolkit_version")
	})
	return err
}

// GetToolkitVersion is a representation of the C type atk_get_toolkit_version.
func GetToolkitVersion() string {

	var ret gi.Argument

	err := getToolkitVersionFunction_Set()
	if err == nil {
		ret = getToolkitVersionFunction.Invoke(nil, nil)
	}

	retGo := ret.String(false)

	return retGo
}

var getVersionFunction *gi.Function
var getVersionFunction_Once sync.Once

func getVersionFunction_Set() error {
	var err error
	getVersionFunction_Once.Do(func() {
		getVersionFunction, err = gi.FunctionInvokerNew("Atk", "get_version")
	})
	return err
}

// GetVersion is a representation of the C type atk_get_version.
func GetVersion() string {

	var ret gi.Argument

	err := getVersionFunction_Set()
	if err == nil {
		ret = getVersionFunction.Invoke(nil, nil)
	}

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'atk_relation_type_for_name' : return type 'RelationType' not supported

// UNSUPPORTED : C value 'atk_relation_type_get_name' : parameter 'type' of type 'RelationType' not supported

// UNSUPPORTED : C value 'atk_relation_type_register' : return type 'RelationType' not supported

var removeFocusTrackerFunction *gi.Function
var removeFocusTrackerFunction_Once sync.Once

func removeFocusTrackerFunction_Set() error {
	var err error
	removeFocusTrackerFunction_Once.Do(func() {
		removeFocusTrackerFunction, err = gi.FunctionInvokerNew("Atk", "remove_focus_tracker")
	})
	return err
}

// RemoveFocusTracker is a representation of the C type atk_remove_focus_tracker.
func RemoveFocusTracker(trackerId uint32) {
	var inArgs [1]gi.Argument
	inArgs[0].SetUint32(trackerId)

	err := removeFocusTrackerFunction_Set()
	if err == nil {
		removeFocusTrackerFunction.Invoke(inArgs[:], nil)
	}

	return
}

var removeGlobalEventListenerFunction *gi.Function
var removeGlobalEventListenerFunction_Once sync.Once

func removeGlobalEventListenerFunction_Set() error {
	var err error
	removeGlobalEventListenerFunction_Once.Do(func() {
		removeGlobalEventListenerFunction, err = gi.FunctionInvokerNew("Atk", "remove_global_event_listener")
	})
	return err
}

// RemoveGlobalEventListener is a representation of the C type atk_remove_global_event_listener.
func RemoveGlobalEventListener(listenerId uint32) {
	var inArgs [1]gi.Argument
	inArgs[0].SetUint32(listenerId)

	err := removeGlobalEventListenerFunction_Set()
	if err == nil {
		removeGlobalEventListenerFunction.Invoke(inArgs[:], nil)
	}

	return
}

var removeKeyEventListenerFunction *gi.Function
var removeKeyEventListenerFunction_Once sync.Once

func removeKeyEventListenerFunction_Set() error {
	var err error
	removeKeyEventListenerFunction_Once.Do(func() {
		removeKeyEventListenerFunction, err = gi.FunctionInvokerNew("Atk", "remove_key_event_listener")
	})
	return err
}

// RemoveKeyEventListener is a representation of the C type atk_remove_key_event_listener.
func RemoveKeyEventListener(listenerId uint32) {
	var inArgs [1]gi.Argument
	inArgs[0].SetUint32(listenerId)

	err := removeKeyEventListenerFunction_Set()
	if err == nil {
		removeKeyEventListenerFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'atk_role_for_name' : return type 'Role' not supported

// UNSUPPORTED : C value 'atk_role_get_localized_name' : parameter 'role' of type 'Role' not supported

// UNSUPPORTED : C value 'atk_role_get_name' : parameter 'role' of type 'Role' not supported

// UNSUPPORTED : C value 'atk_role_register' : return type 'Role' not supported

// UNSUPPORTED : C value 'atk_state_type_for_name' : return type 'StateType' not supported

// UNSUPPORTED : C value 'atk_state_type_get_name' : parameter 'type' of type 'StateType' not supported

// UNSUPPORTED : C value 'atk_state_type_register' : return type 'StateType' not supported

// UNSUPPORTED : C value 'atk_text_attribute_for_name' : return type 'TextAttribute' not supported

// UNSUPPORTED : C value 'atk_text_attribute_get_name' : parameter 'attr' of type 'TextAttribute' not supported

// UNSUPPORTED : C value 'atk_text_attribute_get_value' : parameter 'attr' of type 'TextAttribute' not supported

// UNSUPPORTED : C value 'atk_text_attribute_register' : return type 'TextAttribute' not supported

// UNSUPPORTED : C value 'atk_text_free_ranges' : parameter 'ranges' of type 'nil' not supported

// UNSUPPORTED : C value 'atk_value_type_get_localized_name' : parameter 'value_type' of type 'ValueType' not supported

// UNSUPPORTED : C value 'atk_value_type_get_name' : parameter 'value_type' of type 'ValueType' not supported
