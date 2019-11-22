// Code generated - DO NOT EDIT.

package javascriptcore

import (
	gi "github.com/pekim/gobbi/internal/gi"
	"sync"
)

var getMajorVersionFunction *gi.Function
var getMajorVersionFunction_Once sync.Once

func getMajorVersionFunction_Set() {
	getMajorVersionFunction_Once.Do(func() {
		getMajorVersionFunction = gi.FunctionInvokerNew("JavaScriptCore", "get_major_version")
	})
}

// GetMajorVersion is a representation of the C type jsc_get_major_version.
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
		getMicroVersionFunction = gi.FunctionInvokerNew("JavaScriptCore", "get_micro_version")
	})
}

// GetMicroVersion is a representation of the C type jsc_get_micro_version.
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
		getMinorVersionFunction = gi.FunctionInvokerNew("JavaScriptCore", "get_minor_version")
	})
}

// GetMinorVersion is a representation of the C type jsc_get_minor_version.
func GetMinorVersion() uint32 {

	var ret gi.Argument

	getMinorVersionFunction_Set()

	ret = getMinorVersionFunction.Invoke(nil, nil)

	retGo := ret.Uint32()

	return retGo
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
