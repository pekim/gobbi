// Code generated - DO NOT EDIT.

package atk

import gi "github.com/pekim/gobbi/internal/gi"

// UNSUPPORTED : C value 'atk_add_focus_tracker' : non trivial function

// UNSUPPORTED : C value 'atk_add_global_event_listener' : non trivial function

// UNSUPPORTED : C value 'atk_add_key_event_listener' : non trivial function

// UNSUPPORTED : C value 'atk_attribute_set_free' : non trivial function

// UNSUPPORTED : C value 'atk_focus_tracker_init' : non trivial function

// UNSUPPORTED : C value 'atk_focus_tracker_notify' : non trivial function

var getBinaryAgeInvoker *gi.Function

// GetBinaryAge is a representation of the C type atk_get_binary_age.
func GetBinaryAge() uint32 {
	if getBinaryAgeInvoker == nil {
		getBinaryAgeInvoker = gi.FunctionInvokerNew("Atk", "get_binary_age")
	}

	ret := getBinaryAgeInvoker.Invoke()
	retValue := ret.Uint32()
	return retValue
}

var getDefaultRegistryInvoker *gi.Function

// GetDefaultRegistry is a representation of the C type atk_get_default_registry.
func GetDefaultRegistry() {
	if getDefaultRegistryInvoker == nil {
		getDefaultRegistryInvoker = gi.FunctionInvokerNew("Atk", "get_default_registry")
	}

	getDefaultRegistryInvoker.Invoke()
}

var getFocusObjectInvoker *gi.Function

// GetFocusObject is a representation of the C type atk_get_focus_object.
func GetFocusObject() {
	if getFocusObjectInvoker == nil {
		getFocusObjectInvoker = gi.FunctionInvokerNew("Atk", "get_focus_object")
	}

	getFocusObjectInvoker.Invoke()
}

var getInterfaceAgeInvoker *gi.Function

// GetInterfaceAge is a representation of the C type atk_get_interface_age.
func GetInterfaceAge() uint32 {
	if getInterfaceAgeInvoker == nil {
		getInterfaceAgeInvoker = gi.FunctionInvokerNew("Atk", "get_interface_age")
	}

	ret := getInterfaceAgeInvoker.Invoke()
	retValue := ret.Uint32()
	return retValue
}

var getMajorVersionInvoker *gi.Function

// GetMajorVersion is a representation of the C type atk_get_major_version.
func GetMajorVersion() uint32 {
	if getMajorVersionInvoker == nil {
		getMajorVersionInvoker = gi.FunctionInvokerNew("Atk", "get_major_version")
	}

	ret := getMajorVersionInvoker.Invoke()
	retValue := ret.Uint32()
	return retValue
}

var getMicroVersionInvoker *gi.Function

// GetMicroVersion is a representation of the C type atk_get_micro_version.
func GetMicroVersion() uint32 {
	if getMicroVersionInvoker == nil {
		getMicroVersionInvoker = gi.FunctionInvokerNew("Atk", "get_micro_version")
	}

	ret := getMicroVersionInvoker.Invoke()
	retValue := ret.Uint32()
	return retValue
}

var getMinorVersionInvoker *gi.Function

// GetMinorVersion is a representation of the C type atk_get_minor_version.
func GetMinorVersion() uint32 {
	if getMinorVersionInvoker == nil {
		getMinorVersionInvoker = gi.FunctionInvokerNew("Atk", "get_minor_version")
	}

	ret := getMinorVersionInvoker.Invoke()
	retValue := ret.Uint32()
	return retValue
}

var getRootInvoker *gi.Function

// GetRoot is a representation of the C type atk_get_root.
func GetRoot() {
	if getRootInvoker == nil {
		getRootInvoker = gi.FunctionInvokerNew("Atk", "get_root")
	}

	getRootInvoker.Invoke()
}

var getToolkitNameInvoker *gi.Function

// GetToolkitName is a representation of the C type atk_get_toolkit_name.
func GetToolkitName() {
	if getToolkitNameInvoker == nil {
		getToolkitNameInvoker = gi.FunctionInvokerNew("Atk", "get_toolkit_name")
	}

	getToolkitNameInvoker.Invoke()
}

var getToolkitVersionInvoker *gi.Function

// GetToolkitVersion is a representation of the C type atk_get_toolkit_version.
func GetToolkitVersion() {
	if getToolkitVersionInvoker == nil {
		getToolkitVersionInvoker = gi.FunctionInvokerNew("Atk", "get_toolkit_version")
	}

	getToolkitVersionInvoker.Invoke()
}

var getVersionInvoker *gi.Function

// GetVersion is a representation of the C type atk_get_version.
func GetVersion() {
	if getVersionInvoker == nil {
		getVersionInvoker = gi.FunctionInvokerNew("Atk", "get_version")
	}

	getVersionInvoker.Invoke()
}

// UNSUPPORTED : C value 'atk_relation_type_for_name' : non trivial function

// UNSUPPORTED : C value 'atk_relation_type_get_name' : non trivial function

// UNSUPPORTED : C value 'atk_relation_type_register' : non trivial function

// UNSUPPORTED : C value 'atk_remove_focus_tracker' : non trivial function

// UNSUPPORTED : C value 'atk_remove_global_event_listener' : non trivial function

// UNSUPPORTED : C value 'atk_remove_key_event_listener' : non trivial function

// UNSUPPORTED : C value 'atk_role_for_name' : non trivial function

// UNSUPPORTED : C value 'atk_role_get_localized_name' : non trivial function

// UNSUPPORTED : C value 'atk_role_get_name' : non trivial function

// UNSUPPORTED : C value 'atk_role_register' : non trivial function

// UNSUPPORTED : C value 'atk_state_type_for_name' : non trivial function

// UNSUPPORTED : C value 'atk_state_type_get_name' : non trivial function

// UNSUPPORTED : C value 'atk_state_type_register' : non trivial function

// UNSUPPORTED : C value 'atk_text_attribute_for_name' : non trivial function

// UNSUPPORTED : C value 'atk_text_attribute_get_name' : non trivial function

// UNSUPPORTED : C value 'atk_text_attribute_get_value' : non trivial function

// UNSUPPORTED : C value 'atk_text_attribute_register' : non trivial function

// UNSUPPORTED : C value 'atk_text_free_ranges' : non trivial function

// UNSUPPORTED : C value 'atk_value_type_get_localized_name' : non trivial function

// UNSUPPORTED : C value 'atk_value_type_get_name' : non trivial function
