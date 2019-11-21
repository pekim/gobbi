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

var GetBinaryAgeFunction *gi.Function
var GetBinaryAgeFunctionOnce sync.Once

func GetBinaryAgeFunctionSet() {
	GetBinaryAgeFunctionOnce.Do(func() {
		GetBinaryAgeFunction = gi.FunctionInvokerNew("Atk", "get_binary_age")
	})
}

var getBinaryAgeInvoker *gi.Function

// GetBinaryAge is a representation of the C type atk_get_binary_age.
func GetBinaryAge() uint32 {
	if getBinaryAgeInvoker == nil {
		getBinaryAgeInvoker = gi.FunctionInvokerNew("Atk", "get_binary_age")
	}

	ret := getBinaryAgeInvoker.Invoke(nil, nil)

	retGo := ret.Uint32()

	return retGo
}

// UNSUPPORTED : C value 'atk_get_default_registry' : return type 'Registry' not supported

// UNSUPPORTED : C value 'atk_get_focus_object' : return type 'Object' not supported

var GetInterfaceAgeFunction *gi.Function
var GetInterfaceAgeFunctionOnce sync.Once

func GetInterfaceAgeFunctionSet() {
	GetInterfaceAgeFunctionOnce.Do(func() {
		GetInterfaceAgeFunction = gi.FunctionInvokerNew("Atk", "get_interface_age")
	})
}

var getInterfaceAgeInvoker *gi.Function

// GetInterfaceAge is a representation of the C type atk_get_interface_age.
func GetInterfaceAge() uint32 {
	if getInterfaceAgeInvoker == nil {
		getInterfaceAgeInvoker = gi.FunctionInvokerNew("Atk", "get_interface_age")
	}

	ret := getInterfaceAgeInvoker.Invoke(nil, nil)

	retGo := ret.Uint32()

	return retGo
}

var GetMajorVersionFunction *gi.Function
var GetMajorVersionFunctionOnce sync.Once

func GetMajorVersionFunctionSet() {
	GetMajorVersionFunctionOnce.Do(func() {
		GetMajorVersionFunction = gi.FunctionInvokerNew("Atk", "get_major_version")
	})
}

var getMajorVersionInvoker *gi.Function

// GetMajorVersion is a representation of the C type atk_get_major_version.
func GetMajorVersion() uint32 {
	if getMajorVersionInvoker == nil {
		getMajorVersionInvoker = gi.FunctionInvokerNew("Atk", "get_major_version")
	}

	ret := getMajorVersionInvoker.Invoke(nil, nil)

	retGo := ret.Uint32()

	return retGo
}

var GetMicroVersionFunction *gi.Function
var GetMicroVersionFunctionOnce sync.Once

func GetMicroVersionFunctionSet() {
	GetMicroVersionFunctionOnce.Do(func() {
		GetMicroVersionFunction = gi.FunctionInvokerNew("Atk", "get_micro_version")
	})
}

var getMicroVersionInvoker *gi.Function

// GetMicroVersion is a representation of the C type atk_get_micro_version.
func GetMicroVersion() uint32 {
	if getMicroVersionInvoker == nil {
		getMicroVersionInvoker = gi.FunctionInvokerNew("Atk", "get_micro_version")
	}

	ret := getMicroVersionInvoker.Invoke(nil, nil)

	retGo := ret.Uint32()

	return retGo
}

var GetMinorVersionFunction *gi.Function
var GetMinorVersionFunctionOnce sync.Once

func GetMinorVersionFunctionSet() {
	GetMinorVersionFunctionOnce.Do(func() {
		GetMinorVersionFunction = gi.FunctionInvokerNew("Atk", "get_minor_version")
	})
}

var getMinorVersionInvoker *gi.Function

// GetMinorVersion is a representation of the C type atk_get_minor_version.
func GetMinorVersion() uint32 {
	if getMinorVersionInvoker == nil {
		getMinorVersionInvoker = gi.FunctionInvokerNew("Atk", "get_minor_version")
	}

	ret := getMinorVersionInvoker.Invoke(nil, nil)

	retGo := ret.Uint32()

	return retGo
}

// UNSUPPORTED : C value 'atk_get_root' : return type 'Object' not supported

var GetToolkitNameFunction *gi.Function
var GetToolkitNameFunctionOnce sync.Once

func GetToolkitNameFunctionSet() {
	GetToolkitNameFunctionOnce.Do(func() {
		GetToolkitNameFunction = gi.FunctionInvokerNew("Atk", "get_toolkit_name")
	})
}

var getToolkitNameInvoker *gi.Function

// GetToolkitName is a representation of the C type atk_get_toolkit_name.
func GetToolkitName() string {
	if getToolkitNameInvoker == nil {
		getToolkitNameInvoker = gi.FunctionInvokerNew("Atk", "get_toolkit_name")
	}

	ret := getToolkitNameInvoker.Invoke(nil, nil)

	retGo := ret.String(false)

	return retGo
}

var GetToolkitVersionFunction *gi.Function
var GetToolkitVersionFunctionOnce sync.Once

func GetToolkitVersionFunctionSet() {
	GetToolkitVersionFunctionOnce.Do(func() {
		GetToolkitVersionFunction = gi.FunctionInvokerNew("Atk", "get_toolkit_version")
	})
}

var getToolkitVersionInvoker *gi.Function

// GetToolkitVersion is a representation of the C type atk_get_toolkit_version.
func GetToolkitVersion() string {
	if getToolkitVersionInvoker == nil {
		getToolkitVersionInvoker = gi.FunctionInvokerNew("Atk", "get_toolkit_version")
	}

	ret := getToolkitVersionInvoker.Invoke(nil, nil)

	retGo := ret.String(false)

	return retGo
}

var GetVersionFunction *gi.Function
var GetVersionFunctionOnce sync.Once

func GetVersionFunctionSet() {
	GetVersionFunctionOnce.Do(func() {
		GetVersionFunction = gi.FunctionInvokerNew("Atk", "get_version")
	})
}

var getVersionInvoker *gi.Function

// GetVersion is a representation of the C type atk_get_version.
func GetVersion() string {
	if getVersionInvoker == nil {
		getVersionInvoker = gi.FunctionInvokerNew("Atk", "get_version")
	}

	ret := getVersionInvoker.Invoke(nil, nil)

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'atk_relation_type_for_name' : return type 'RelationType' not supported

// UNSUPPORTED : C value 'atk_relation_type_get_name' : parameter 'type' of type 'RelationType' not supported

// UNSUPPORTED : C value 'atk_relation_type_register' : return type 'RelationType' not supported

var RemoveFocusTrackerFunction *gi.Function
var RemoveFocusTrackerFunctionOnce sync.Once

func RemoveFocusTrackerFunctionSet() {
	RemoveFocusTrackerFunctionOnce.Do(func() {
		RemoveFocusTrackerFunction = gi.FunctionInvokerNew("Atk", "remove_focus_tracker")
	})
}

var removeFocusTrackerInvoker *gi.Function

// RemoveFocusTracker is a representation of the C type atk_remove_focus_tracker.
func RemoveFocusTracker(trackerId uint32) {
	if removeFocusTrackerInvoker == nil {
		removeFocusTrackerInvoker = gi.FunctionInvokerNew("Atk", "remove_focus_tracker")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetUint32(trackerId)

	removeFocusTrackerInvoker.Invoke(inArgs[:], nil)

}

var RemoveGlobalEventListenerFunction *gi.Function
var RemoveGlobalEventListenerFunctionOnce sync.Once

func RemoveGlobalEventListenerFunctionSet() {
	RemoveGlobalEventListenerFunctionOnce.Do(func() {
		RemoveGlobalEventListenerFunction = gi.FunctionInvokerNew("Atk", "remove_global_event_listener")
	})
}

var removeGlobalEventListenerInvoker *gi.Function

// RemoveGlobalEventListener is a representation of the C type atk_remove_global_event_listener.
func RemoveGlobalEventListener(listenerId uint32) {
	if removeGlobalEventListenerInvoker == nil {
		removeGlobalEventListenerInvoker = gi.FunctionInvokerNew("Atk", "remove_global_event_listener")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetUint32(listenerId)

	removeGlobalEventListenerInvoker.Invoke(inArgs[:], nil)

}

var RemoveKeyEventListenerFunction *gi.Function
var RemoveKeyEventListenerFunctionOnce sync.Once

func RemoveKeyEventListenerFunctionSet() {
	RemoveKeyEventListenerFunctionOnce.Do(func() {
		RemoveKeyEventListenerFunction = gi.FunctionInvokerNew("Atk", "remove_key_event_listener")
	})
}

var removeKeyEventListenerInvoker *gi.Function

// RemoveKeyEventListener is a representation of the C type atk_remove_key_event_listener.
func RemoveKeyEventListener(listenerId uint32) {
	if removeKeyEventListenerInvoker == nil {
		removeKeyEventListenerInvoker = gi.FunctionInvokerNew("Atk", "remove_key_event_listener")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetUint32(listenerId)

	removeKeyEventListenerInvoker.Invoke(inArgs[:], nil)

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
