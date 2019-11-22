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

// UNSUPPORTED : C value 'atk_focus_tracker_notify' : parameter 'object' of type 'Object' not supported

var getBinaryAgeFunction *gi.Function
var getBinaryAgeFunction_Once sync.Once

func getBinaryAgeFunction_Set() {
	getBinaryAgeFunction_Once.Do(func() {
		getBinaryAgeFunction = gi.FunctionInvokerNew("Atk", "get_binary_age")
	})
}

// GetBinaryAge is a representation of the C type atk_get_binary_age.
func GetBinaryAge() uint32 {

	var ret gi.Argument

	getBinaryAgeFunction_Set()

	ret = getBinaryAgeFunction.Invoke(nil, nil)

	retGo := ret.Uint32()

	return retGo
}

// UNSUPPORTED : C value 'atk_get_default_registry' : return type 'Registry' not supported

// UNSUPPORTED : C value 'atk_get_focus_object' : return type 'Object' not supported

var getInterfaceAgeFunction *gi.Function
var getInterfaceAgeFunction_Once sync.Once

func getInterfaceAgeFunction_Set() {
	getInterfaceAgeFunction_Once.Do(func() {
		getInterfaceAgeFunction = gi.FunctionInvokerNew("Atk", "get_interface_age")
	})
}

// GetInterfaceAge is a representation of the C type atk_get_interface_age.
func GetInterfaceAge() uint32 {

	var ret gi.Argument

	getInterfaceAgeFunction_Set()

	ret = getInterfaceAgeFunction.Invoke(nil, nil)

	retGo := ret.Uint32()

	return retGo
}

var getMajorVersionFunction *gi.Function
var getMajorVersionFunction_Once sync.Once

func getMajorVersionFunction_Set() {
	getMajorVersionFunction_Once.Do(func() {
		getMajorVersionFunction = gi.FunctionInvokerNew("Atk", "get_major_version")
	})
}

// GetMajorVersion is a representation of the C type atk_get_major_version.
func GetMajorVersion() uint32 {

	var ret gi.Argument

	getMajorVersionFunction_Set()

	ret = getMajorVersionFunction.Invoke(nil, nil)

	retGo := ret.Uint32()

	return retGo
}

var getMicroVersionFunction *gi.Function
var getMicroVersionFunction_Once sync.Once

func getMicroVersionFunction_Set() {
	getMicroVersionFunction_Once.Do(func() {
		getMicroVersionFunction = gi.FunctionInvokerNew("Atk", "get_micro_version")
	})
}

// GetMicroVersion is a representation of the C type atk_get_micro_version.
func GetMicroVersion() uint32 {

	var ret gi.Argument

	getMicroVersionFunction_Set()

	ret = getMicroVersionFunction.Invoke(nil, nil)

	retGo := ret.Uint32()

	return retGo
}

var getMinorVersionFunction *gi.Function
var getMinorVersionFunction_Once sync.Once

func getMinorVersionFunction_Set() {
	getMinorVersionFunction_Once.Do(func() {
		getMinorVersionFunction = gi.FunctionInvokerNew("Atk", "get_minor_version")
	})
}

// GetMinorVersion is a representation of the C type atk_get_minor_version.
func GetMinorVersion() uint32 {

	var ret gi.Argument

	getMinorVersionFunction_Set()

	ret = getMinorVersionFunction.Invoke(nil, nil)

	retGo := ret.Uint32()

	return retGo
}

// UNSUPPORTED : C value 'atk_get_root' : return type 'Object' not supported

var getToolkitNameFunction *gi.Function
var getToolkitNameFunction_Once sync.Once

func getToolkitNameFunction_Set() {
	getToolkitNameFunction_Once.Do(func() {
		getToolkitNameFunction = gi.FunctionInvokerNew("Atk", "get_toolkit_name")
	})
}

// GetToolkitName is a representation of the C type atk_get_toolkit_name.
func GetToolkitName() string {

	var ret gi.Argument

	getToolkitNameFunction_Set()

	ret = getToolkitNameFunction.Invoke(nil, nil)

	retGo := ret.String(false)

	return retGo
}

var getToolkitVersionFunction *gi.Function
var getToolkitVersionFunction_Once sync.Once

func getToolkitVersionFunction_Set() {
	getToolkitVersionFunction_Once.Do(func() {
		getToolkitVersionFunction = gi.FunctionInvokerNew("Atk", "get_toolkit_version")
	})
}

// GetToolkitVersion is a representation of the C type atk_get_toolkit_version.
func GetToolkitVersion() string {

	var ret gi.Argument

	getToolkitVersionFunction_Set()

	ret = getToolkitVersionFunction.Invoke(nil, nil)

	retGo := ret.String(false)

	return retGo
}

var getVersionFunction *gi.Function
var getVersionFunction_Once sync.Once

func getVersionFunction_Set() {
	getVersionFunction_Once.Do(func() {
		getVersionFunction = gi.FunctionInvokerNew("Atk", "get_version")
	})
}

// GetVersion is a representation of the C type atk_get_version.
func GetVersion() string {

	var ret gi.Argument

	getVersionFunction_Set()

	ret = getVersionFunction.Invoke(nil, nil)

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'atk_relation_type_for_name' : return type 'RelationType' not supported

// UNSUPPORTED : C value 'atk_relation_type_get_name' : parameter 'type' of type 'RelationType' not supported

// UNSUPPORTED : C value 'atk_relation_type_register' : return type 'RelationType' not supported

var removeFocusTrackerFunction *gi.Function
var removeFocusTrackerFunction_Once sync.Once

func removeFocusTrackerFunction_Set() {
	removeFocusTrackerFunction_Once.Do(func() {
		removeFocusTrackerFunction = gi.FunctionInvokerNew("Atk", "remove_focus_tracker")
	})
}

// RemoveFocusTracker is a representation of the C type atk_remove_focus_tracker.
func RemoveFocusTracker(trackerId uint32) {
	var inArgs [1]gi.Argument
	inArgs[0].SetUint32(trackerId)

	removeFocusTrackerFunction_Set()

	removeFocusTrackerFunction.Invoke(inArgs[:], nil)

}

var removeGlobalEventListenerFunction *gi.Function
var removeGlobalEventListenerFunction_Once sync.Once

func removeGlobalEventListenerFunction_Set() {
	removeGlobalEventListenerFunction_Once.Do(func() {
		removeGlobalEventListenerFunction = gi.FunctionInvokerNew("Atk", "remove_global_event_listener")
	})
}

// RemoveGlobalEventListener is a representation of the C type atk_remove_global_event_listener.
func RemoveGlobalEventListener(listenerId uint32) {
	var inArgs [1]gi.Argument
	inArgs[0].SetUint32(listenerId)

	removeGlobalEventListenerFunction_Set()

	removeGlobalEventListenerFunction.Invoke(inArgs[:], nil)

}

var removeKeyEventListenerFunction *gi.Function
var removeKeyEventListenerFunction_Once sync.Once

func removeKeyEventListenerFunction_Set() {
	removeKeyEventListenerFunction_Once.Do(func() {
		removeKeyEventListenerFunction = gi.FunctionInvokerNew("Atk", "remove_key_event_listener")
	})
}

// RemoveKeyEventListener is a representation of the C type atk_remove_key_event_listener.
func RemoveKeyEventListener(listenerId uint32) {
	var inArgs [1]gi.Argument
	inArgs[0].SetUint32(listenerId)

	removeKeyEventListenerFunction_Set()

	removeKeyEventListenerFunction.Invoke(inArgs[:], nil)

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

// UNSUPPORTED : C value 'atk_text_free_ranges' : parameter 'ranges' has no type

// UNSUPPORTED : C value 'atk_value_type_get_localized_name' : parameter 'value_type' of type 'ValueType' not supported

// UNSUPPORTED : C value 'atk_value_type_get_name' : parameter 'value_type' of type 'ValueType' not supported
