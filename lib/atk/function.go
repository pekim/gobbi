// Code generated - DO NOT EDIT.

package atk

import gi "github.com/pekim/gobbi/internal/gi"

// UNSUPPORTED : C value 'atk_add_focus_tracker' : has parameters

// UNSUPPORTED : C value 'atk_add_global_event_listener' : has parameters

// UNSUPPORTED : C value 'atk_add_key_event_listener' : has parameters

// UNSUPPORTED : C value 'atk_attribute_set_free' : has parameters

// UNSUPPORTED : C value 'atk_focus_tracker_init' : has parameters

// UNSUPPORTED : C value 'atk_focus_tracker_notify' : has parameters

var getBinaryAgeInvoker *gi.Function

// GetBinaryAge is a representation of the C type atk_get_binary_age.
func GetBinaryAge() uint32 {
	if getBinaryAgeInvoker == nil {
		getBinaryAgeInvoker = gi.FunctionInvokerNew("Atk", "get_binary_age")
	}

	ret := getBinaryAgeInvoker.Invoke()
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

	ret := getInterfaceAgeInvoker.Invoke()
	return ret.Uint32()
}

var getMajorVersionInvoker *gi.Function

// GetMajorVersion is a representation of the C type atk_get_major_version.
func GetMajorVersion() uint32 {
	if getMajorVersionInvoker == nil {
		getMajorVersionInvoker = gi.FunctionInvokerNew("Atk", "get_major_version")
	}

	ret := getMajorVersionInvoker.Invoke()
	return ret.Uint32()
}

var getMicroVersionInvoker *gi.Function

// GetMicroVersion is a representation of the C type atk_get_micro_version.
func GetMicroVersion() uint32 {
	if getMicroVersionInvoker == nil {
		getMicroVersionInvoker = gi.FunctionInvokerNew("Atk", "get_micro_version")
	}

	ret := getMicroVersionInvoker.Invoke()
	return ret.Uint32()
}

var getMinorVersionInvoker *gi.Function

// GetMinorVersion is a representation of the C type atk_get_minor_version.
func GetMinorVersion() uint32 {
	if getMinorVersionInvoker == nil {
		getMinorVersionInvoker = gi.FunctionInvokerNew("Atk", "get_minor_version")
	}

	ret := getMinorVersionInvoker.Invoke()
	return ret.Uint32()
}

// UNSUPPORTED : C value 'atk_get_root' : return type 'Object' not supported

var getToolkitNameInvoker *gi.Function

// GetToolkitName is a representation of the C type atk_get_toolkit_name.
func GetToolkitName() string {
	if getToolkitNameInvoker == nil {
		getToolkitNameInvoker = gi.FunctionInvokerNew("Atk", "get_toolkit_name")
	}

	ret := getToolkitNameInvoker.Invoke()
	return ret.String()
}

var getToolkitVersionInvoker *gi.Function

// GetToolkitVersion is a representation of the C type atk_get_toolkit_version.
func GetToolkitVersion() string {
	if getToolkitVersionInvoker == nil {
		getToolkitVersionInvoker = gi.FunctionInvokerNew("Atk", "get_toolkit_version")
	}

	ret := getToolkitVersionInvoker.Invoke()
	return ret.String()
}

var getVersionInvoker *gi.Function

// GetVersion is a representation of the C type atk_get_version.
func GetVersion() string {
	if getVersionInvoker == nil {
		getVersionInvoker = gi.FunctionInvokerNew("Atk", "get_version")
	}

	ret := getVersionInvoker.Invoke()
	return ret.String()
}

// UNSUPPORTED : C value 'atk_relation_type_for_name' : has parameters

// UNSUPPORTED : C value 'atk_relation_type_get_name' : has parameters

// UNSUPPORTED : C value 'atk_relation_type_register' : has parameters

// UNSUPPORTED : C value 'atk_remove_focus_tracker' : has parameters

// UNSUPPORTED : C value 'atk_remove_global_event_listener' : has parameters

// UNSUPPORTED : C value 'atk_remove_key_event_listener' : has parameters

// UNSUPPORTED : C value 'atk_role_for_name' : has parameters

// UNSUPPORTED : C value 'atk_role_get_localized_name' : has parameters

// UNSUPPORTED : C value 'atk_role_get_name' : has parameters

// UNSUPPORTED : C value 'atk_role_register' : has parameters

// UNSUPPORTED : C value 'atk_state_type_for_name' : has parameters

// UNSUPPORTED : C value 'atk_state_type_get_name' : has parameters

// UNSUPPORTED : C value 'atk_state_type_register' : has parameters

// UNSUPPORTED : C value 'atk_text_attribute_for_name' : has parameters

// UNSUPPORTED : C value 'atk_text_attribute_get_name' : has parameters

// UNSUPPORTED : C value 'atk_text_attribute_get_value' : has parameters

// UNSUPPORTED : C value 'atk_text_attribute_register' : has parameters

// UNSUPPORTED : C value 'atk_text_free_ranges' : has parameters

// UNSUPPORTED : C value 'atk_value_type_get_localized_name' : has parameters

// UNSUPPORTED : C value 'atk_value_type_get_name' : has parameters
