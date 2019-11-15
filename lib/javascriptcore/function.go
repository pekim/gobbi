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
	return ret.Uint32()
}

var getMicroVersionInvoker *gi.Function

// GetMicroVersion is a representation of the C type jsc_get_micro_version.
func GetMicroVersion() uint32 {
	if getMicroVersionInvoker == nil {
		getMicroVersionInvoker = gi.FunctionInvokerNew("JavaScriptCore", "get_micro_version")
	}

	ret := getMicroVersionInvoker.Invoke()
	return ret.Uint32()
}

var getMinorVersionInvoker *gi.Function

// GetMinorVersion is a representation of the C type jsc_get_minor_version.
func GetMinorVersion() uint32 {
	if getMinorVersionInvoker == nil {
		getMinorVersionInvoker = gi.FunctionInvokerNew("JavaScriptCore", "get_minor_version")
	}

	ret := getMinorVersionInvoker.Invoke()
	return ret.Uint32()
}

// UNSUPPORTED : C value 'jsc_options_foreach' : has parameters

// UNSUPPORTED : C value 'jsc_options_get_boolean' : has parameters

// UNSUPPORTED : C value 'jsc_options_get_double' : has parameters

// UNSUPPORTED : C value 'jsc_options_get_int' : has parameters

// UNSUPPORTED : C value 'jsc_options_get_option_group' : return type not supported

// UNSUPPORTED : C value 'jsc_options_get_range_string' : has parameters

// UNSUPPORTED : C value 'jsc_options_get_size' : has parameters

// UNSUPPORTED : C value 'jsc_options_get_string' : has parameters

// UNSUPPORTED : C value 'jsc_options_get_uint' : has parameters

// UNSUPPORTED : C value 'jsc_options_set_boolean' : has parameters

// UNSUPPORTED : C value 'jsc_options_set_double' : has parameters

// UNSUPPORTED : C value 'jsc_options_set_int' : has parameters

// UNSUPPORTED : C value 'jsc_options_set_range_string' : has parameters

// UNSUPPORTED : C value 'jsc_options_set_size' : has parameters

// UNSUPPORTED : C value 'jsc_options_set_string' : has parameters

// UNSUPPORTED : C value 'jsc_options_set_uint' : has parameters
