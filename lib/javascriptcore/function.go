// Code generated - DO NOT EDIT.

package javascriptcore

import gi "github.com/pekim/gobbi/internal/gi"

var getMajorVersionInvoker *gi.Function

// GetMajorVersion is a representation of the C type jsc_get_major_version.
func GetMajorVersion() uint32 {
	if getMajorVersionInvoker == nil {
		getMajorVersionInvoker = gi.FunctionInvokerNew("JavaScriptCore", "get_major_version")
	}

	ret := getMajorVersionInvoker.Invoke(nil)
	return ret.Uint32()
}

var getMicroVersionInvoker *gi.Function

// GetMicroVersion is a representation of the C type jsc_get_micro_version.
func GetMicroVersion() uint32 {
	if getMicroVersionInvoker == nil {
		getMicroVersionInvoker = gi.FunctionInvokerNew("JavaScriptCore", "get_micro_version")
	}

	ret := getMicroVersionInvoker.Invoke(nil)
	return ret.Uint32()
}

var getMinorVersionInvoker *gi.Function

// GetMinorVersion is a representation of the C type jsc_get_minor_version.
func GetMinorVersion() uint32 {
	if getMinorVersionInvoker == nil {
		getMinorVersionInvoker = gi.FunctionInvokerNew("JavaScriptCore", "get_minor_version")
	}

	ret := getMinorVersionInvoker.Invoke(nil)
	return ret.Uint32()
}

// UNSUPPORTED : C value 'jsc_options_foreach' : parameter 'function' of type 'OptionsFunc' not supported

// UNSUPPORTED : C value 'jsc_options_get_boolean' : parameter 'option' of type 'utf8' not supported

// UNSUPPORTED : C value 'jsc_options_get_double' : parameter 'option' of type 'utf8' not supported

// UNSUPPORTED : C value 'jsc_options_get_int' : parameter 'option' of type 'utf8' not supported

// UNSUPPORTED : C value 'jsc_options_get_option_group' : return type 'GLib.OptionGroup' not supported

// UNSUPPORTED : C value 'jsc_options_get_range_string' : parameter 'option' of type 'utf8' not supported

// UNSUPPORTED : C value 'jsc_options_get_size' : parameter 'option' of type 'utf8' not supported

// UNSUPPORTED : C value 'jsc_options_get_string' : parameter 'option' of type 'utf8' not supported

// UNSUPPORTED : C value 'jsc_options_get_uint' : parameter 'option' of type 'utf8' not supported

// UNSUPPORTED : C value 'jsc_options_set_boolean' : parameter 'option' of type 'utf8' not supported

// UNSUPPORTED : C value 'jsc_options_set_double' : parameter 'option' of type 'utf8' not supported

// UNSUPPORTED : C value 'jsc_options_set_int' : parameter 'option' of type 'utf8' not supported

// UNSUPPORTED : C value 'jsc_options_set_range_string' : parameter 'option' of type 'utf8' not supported

// UNSUPPORTED : C value 'jsc_options_set_size' : parameter 'option' of type 'utf8' not supported

// UNSUPPORTED : C value 'jsc_options_set_string' : parameter 'option' of type 'utf8' not supported

// UNSUPPORTED : C value 'jsc_options_set_uint' : parameter 'option' of type 'utf8' not supported
