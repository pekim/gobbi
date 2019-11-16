// Code generated - DO NOT EDIT.

package atk

import gi "github.com/pekim/gobbi/internal/gi"

// UNSUPPORTED : C value 'atk_add_focus_tracker' : parameter 'focus_tracker' of type 'EventListener' not supported

// UNSUPPORTED : C value 'atk_add_global_event_listener' : parameter 'listener' of type 'GObject.SignalEmissionHook' not supported

// UNSUPPORTED : C value 'atk_add_key_event_listener' : parameter 'listener' of type 'KeySnoopFunc' not supported

// UNSUPPORTED : C value 'atk_attribute_set_free' : parameter 'attrib_set' of type 'AttributeSet' not supported

// UNSUPPORTED : C value 'atk_focus_tracker_init' : parameter 'init' of type 'EventListenerInit' not supported

// UNSUPPORTED : C value 'atk_focus_tracker_notify' : parameter 'object' of type 'Object' not supported

var getBinaryAgeInvoker *gi.Function

// GetBinaryAge is a representation of the C type atk_get_binary_age.
func GetBinaryAge() uint32 {
	if getBinaryAgeInvoker == nil {
		getBinaryAgeInvoker = gi.FunctionInvokerNew("Atk", "get_binary_age")
	}

	ret := getBinaryAgeInvoker.Invoke(nil)
	return ret.Uint32()
}

// UNSUPPORTED : C value 'atk_get_default_registry' : return type 'Registry' not supported

// UNSUPPORTED : C value 'atk_get_focus_object' : return type 'Object' not supported

var getInterfaceAgeInvoker *gi.Function

// GetInterfaceAge is a representation of the C type atk_get_interface_age.
func GetInterfaceAge() uint32 {
	if getInterfaceAgeInvoker == nil {
		getInterfaceAgeInvoker = gi.FunctionInvokerNew("Atk", "get_interface_age")
	}

	ret := getInterfaceAgeInvoker.Invoke(nil)
	return ret.Uint32()
}

var getMajorVersionInvoker *gi.Function

// GetMajorVersion is a representation of the C type atk_get_major_version.
func GetMajorVersion() uint32 {
	if getMajorVersionInvoker == nil {
		getMajorVersionInvoker = gi.FunctionInvokerNew("Atk", "get_major_version")
	}

	ret := getMajorVersionInvoker.Invoke(nil)
	return ret.Uint32()
}

var getMicroVersionInvoker *gi.Function

// GetMicroVersion is a representation of the C type atk_get_micro_version.
func GetMicroVersion() uint32 {
	if getMicroVersionInvoker == nil {
		getMicroVersionInvoker = gi.FunctionInvokerNew("Atk", "get_micro_version")
	}

	ret := getMicroVersionInvoker.Invoke(nil)
	return ret.Uint32()
}

var getMinorVersionInvoker *gi.Function

// GetMinorVersion is a representation of the C type atk_get_minor_version.
func GetMinorVersion() uint32 {
	if getMinorVersionInvoker == nil {
		getMinorVersionInvoker = gi.FunctionInvokerNew("Atk", "get_minor_version")
	}

	ret := getMinorVersionInvoker.Invoke(nil)
	return ret.Uint32()
}

// UNSUPPORTED : C value 'atk_get_root' : return type 'Object' not supported

var getToolkitNameInvoker *gi.Function

// GetToolkitName is a representation of the C type atk_get_toolkit_name.
func GetToolkitName() string {
	if getToolkitNameInvoker == nil {
		getToolkitNameInvoker = gi.FunctionInvokerNew("Atk", "get_toolkit_name")
	}

	ret := getToolkitNameInvoker.Invoke(nil)
	return ret.String(false)
}

var getToolkitVersionInvoker *gi.Function

// GetToolkitVersion is a representation of the C type atk_get_toolkit_version.
func GetToolkitVersion() string {
	if getToolkitVersionInvoker == nil {
		getToolkitVersionInvoker = gi.FunctionInvokerNew("Atk", "get_toolkit_version")
	}

	ret := getToolkitVersionInvoker.Invoke(nil)
	return ret.String(false)
}

var getVersionInvoker *gi.Function

// GetVersion is a representation of the C type atk_get_version.
func GetVersion() string {
	if getVersionInvoker == nil {
		getVersionInvoker = gi.FunctionInvokerNew("Atk", "get_version")
	}

	ret := getVersionInvoker.Invoke(nil)
	return ret.String(false)
}

// UNSUPPORTED : C value 'atk_relation_type_for_name' : parameter 'name' of type 'utf8' not supported

// UNSUPPORTED : C value 'atk_relation_type_get_name' : parameter 'type' of type 'RelationType' not supported

// UNSUPPORTED : C value 'atk_relation_type_register' : parameter 'name' of type 'utf8' not supported

var removeFocusTrackerInvoker *gi.Function

// RemoveFocusTracker is a representation of the C type atk_remove_focus_tracker.
func RemoveFocusTracker(trackerId uint32) {
	if removeFocusTrackerInvoker == nil {
		removeFocusTrackerInvoker = gi.FunctionInvokerNew("Atk", "remove_focus_tracker")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetUint32(trackerId)

	removeFocusTrackerInvoker.Invoke(inArgs[:])
}

var removeGlobalEventListenerInvoker *gi.Function

// RemoveGlobalEventListener is a representation of the C type atk_remove_global_event_listener.
func RemoveGlobalEventListener(listenerId uint32) {
	if removeGlobalEventListenerInvoker == nil {
		removeGlobalEventListenerInvoker = gi.FunctionInvokerNew("Atk", "remove_global_event_listener")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetUint32(listenerId)

	removeGlobalEventListenerInvoker.Invoke(inArgs[:])
}

var removeKeyEventListenerInvoker *gi.Function

// RemoveKeyEventListener is a representation of the C type atk_remove_key_event_listener.
func RemoveKeyEventListener(listenerId uint32) {
	if removeKeyEventListenerInvoker == nil {
		removeKeyEventListenerInvoker = gi.FunctionInvokerNew("Atk", "remove_key_event_listener")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetUint32(listenerId)

	removeKeyEventListenerInvoker.Invoke(inArgs[:])
}

// UNSUPPORTED : C value 'atk_role_for_name' : parameter 'name' of type 'utf8' not supported

// UNSUPPORTED : C value 'atk_role_get_localized_name' : parameter 'role' of type 'Role' not supported

// UNSUPPORTED : C value 'atk_role_get_name' : parameter 'role' of type 'Role' not supported

// UNSUPPORTED : C value 'atk_role_register' : parameter 'name' of type 'utf8' not supported

// UNSUPPORTED : C value 'atk_state_type_for_name' : parameter 'name' of type 'utf8' not supported

// UNSUPPORTED : C value 'atk_state_type_get_name' : parameter 'type' of type 'StateType' not supported

// UNSUPPORTED : C value 'atk_state_type_register' : parameter 'name' of type 'utf8' not supported

// UNSUPPORTED : C value 'atk_text_attribute_for_name' : parameter 'name' of type 'utf8' not supported

// UNSUPPORTED : C value 'atk_text_attribute_get_name' : parameter 'attr' of type 'TextAttribute' not supported

// UNSUPPORTED : C value 'atk_text_attribute_get_value' : parameter 'attr' of type 'TextAttribute' not supported

// UNSUPPORTED : C value 'atk_text_attribute_register' : parameter 'name' of type 'utf8' not supported

// UNSUPPORTED : C value 'atk_text_free_ranges' : parameter 'ranges' has no type

// UNSUPPORTED : C value 'atk_value_type_get_localized_name' : parameter 'value_type' of type 'ValueType' not supported

// UNSUPPORTED : C value 'atk_value_type_get_name' : parameter 'value_type' of type 'ValueType' not supported
