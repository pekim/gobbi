// Code generated - DO NOT EDIT.

package javascriptcore

import gi "github.com/pekim/gobbi/internal/gi"

var getMajorVersionInvoker *gi.Function

// GetMajorVersion is a representation of the C type jsc_get_major_version.
func GetMajorVersion() uint32 {
	if getMajorVersionInvoker == nil {
		getMajorVersionInvoker = gi.FunctionInvokerNew("JavaScriptCore", "get_major_version")
	}

	ret := getMajorVersionInvoker.Invoke()
	retValue := ret.Uint32()
	return retValue
}

var getMicroVersionInvoker *gi.Function

// GetMicroVersion is a representation of the C type jsc_get_micro_version.
func GetMicroVersion() uint32 {
	if getMicroVersionInvoker == nil {
		getMicroVersionInvoker = gi.FunctionInvokerNew("JavaScriptCore", "get_micro_version")
	}

	ret := getMicroVersionInvoker.Invoke()
	retValue := ret.Uint32()
	return retValue
}

var getMinorVersionInvoker *gi.Function

// GetMinorVersion is a representation of the C type jsc_get_minor_version.
func GetMinorVersion() uint32 {
	if getMinorVersionInvoker == nil {
		getMinorVersionInvoker = gi.FunctionInvokerNew("JavaScriptCore", "get_minor_version")
	}

	ret := getMinorVersionInvoker.Invoke()
	retValue := ret.Uint32()
	return retValue
}

// UNSUPPORTED : C value 'jsc_options_foreach' : non trivial function

// UNSUPPORTED : C value 'jsc_options_get_boolean' : non trivial function

// UNSUPPORTED : C value 'jsc_options_get_double' : non trivial function

// UNSUPPORTED : C value 'jsc_options_get_int' : non trivial function

var optionsGetOptionGroupInvoker *gi.Function

// OptionsGetOptionGroup is a representation of the C type jsc_options_get_option_group.
func OptionsGetOptionGroup() {
	if optionsGetOptionGroupInvoker == nil {
		optionsGetOptionGroupInvoker = gi.FunctionInvokerNew("JavaScriptCore", "options_get_option_group")
	}

	optionsGetOptionGroupInvoker.Invoke()
}

// UNSUPPORTED : C value 'jsc_options_get_range_string' : non trivial function

// UNSUPPORTED : C value 'jsc_options_get_size' : non trivial function

// UNSUPPORTED : C value 'jsc_options_get_string' : non trivial function

// UNSUPPORTED : C value 'jsc_options_get_uint' : non trivial function

// UNSUPPORTED : C value 'jsc_options_set_boolean' : non trivial function

// UNSUPPORTED : C value 'jsc_options_set_double' : non trivial function

// UNSUPPORTED : C value 'jsc_options_set_int' : non trivial function

// UNSUPPORTED : C value 'jsc_options_set_range_string' : non trivial function

// UNSUPPORTED : C value 'jsc_options_set_size' : non trivial function

// UNSUPPORTED : C value 'jsc_options_set_string' : non trivial function

// UNSUPPORTED : C value 'jsc_options_set_uint' : non trivial function
