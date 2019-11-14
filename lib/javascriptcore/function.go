// Code generated - DO NOT EDIT.

package javascriptcore

import gi "github.com/pekim/gobbi/internal/gi"

var getMajorVersionInvoker *gi.FunctionInvoker

// GetMajorVersion is a representation of the C type jsc_get_major_version.
func GetMajorVersion() {
	if getMajorVersionInvoker == nil {
		getMajorVersionInvoker = gi.FunctionInvokerNew("JavaScriptCore", "get_major_version")
	}

	getMajorVersionInvoker.Call()
}

var getMicroVersionInvoker *gi.FunctionInvoker

// GetMicroVersion is a representation of the C type jsc_get_micro_version.
func GetMicroVersion() {
	if getMicroVersionInvoker == nil {
		getMicroVersionInvoker = gi.FunctionInvokerNew("JavaScriptCore", "get_micro_version")
	}

	getMicroVersionInvoker.Call()
}

var getMinorVersionInvoker *gi.FunctionInvoker

// GetMinorVersion is a representation of the C type jsc_get_minor_version.
func GetMinorVersion() {
	if getMinorVersionInvoker == nil {
		getMinorVersionInvoker = gi.FunctionInvokerNew("JavaScriptCore", "get_minor_version")
	}

	getMinorVersionInvoker.Call()
}

// UNSUPPORTED : C value 'jsc_options_foreach' : non trivial function

// UNSUPPORTED : C value 'jsc_options_get_boolean' : non trivial function

// UNSUPPORTED : C value 'jsc_options_get_double' : non trivial function

// UNSUPPORTED : C value 'jsc_options_get_int' : non trivial function

var optionsGetOptionGroupInvoker *gi.FunctionInvoker

// OptionsGetOptionGroup is a representation of the C type jsc_options_get_option_group.
func OptionsGetOptionGroup() {
	if optionsGetOptionGroupInvoker == nil {
		optionsGetOptionGroupInvoker = gi.FunctionInvokerNew("JavaScriptCore", "options_get_option_group")
	}

	optionsGetOptionGroupInvoker.Call()
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
