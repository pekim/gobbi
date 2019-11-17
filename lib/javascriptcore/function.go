// Code generated - DO NOT EDIT.

package javascriptcore

import gi "github.com/pekim/gobbi/internal/gi"

var getMajorVersionInvoker *gi.Function

// GetMajorVersion is a representation of the C type jsc_get_major_version.
func GetMajorVersion() uint32 {
	if getMajorVersionInvoker == nil {
		getMajorVersionInvoker = gi.FunctionInvokerNew("JavaScriptCore", "get_major_version")
	}

	ret := getMajorVersionInvoker.Invoke(nil, nil)

	return ret.Uint32()
}

var getMicroVersionInvoker *gi.Function

// GetMicroVersion is a representation of the C type jsc_get_micro_version.
func GetMicroVersion() uint32 {
	if getMicroVersionInvoker == nil {
		getMicroVersionInvoker = gi.FunctionInvokerNew("JavaScriptCore", "get_micro_version")
	}

	ret := getMicroVersionInvoker.Invoke(nil, nil)

	return ret.Uint32()
}

var getMinorVersionInvoker *gi.Function

// GetMinorVersion is a representation of the C type jsc_get_minor_version.
func GetMinorVersion() uint32 {
	if getMinorVersionInvoker == nil {
		getMinorVersionInvoker = gi.FunctionInvokerNew("JavaScriptCore", "get_minor_version")
	}

	ret := getMinorVersionInvoker.Invoke(nil, nil)

	return ret.Uint32()
}

// UNSUPPORTED : C value 'jsc_options_foreach' : parameter 'function' of type 'OptionsFunc' not supported

// UNSUPPORTED : C value 'jsc_options_get_boolean' : parameter 'value' of type 'gboolean' not supported

// UNSUPPORTED : C value 'jsc_options_get_double' : parameter 'value' of type 'gdouble' not supported

// UNSUPPORTED : C value 'jsc_options_get_int' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'jsc_options_get_option_group' : return type 'GLib.OptionGroup' not supported

// UNSUPPORTED : C value 'jsc_options_get_range_string' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'jsc_options_get_size' : parameter 'value' of type 'gsize' not supported

// UNSUPPORTED : C value 'jsc_options_get_string' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'jsc_options_get_uint' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'jsc_options_set_boolean' : parameter 'value' of type 'gboolean' not supported

// UNSUPPORTED : C value 'jsc_options_set_double' : parameter 'value' of type 'gdouble' not supported

// UNSUPPORTED : C value 'jsc_options_set_int' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'jsc_options_set_range_string' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'jsc_options_set_size' : parameter 'value' of type 'gsize' not supported

// UNSUPPORTED : C value 'jsc_options_set_string' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'jsc_options_set_uint' : return type 'gboolean' not supported
