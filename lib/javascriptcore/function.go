// Code generated - DO NOT EDIT.

package javascriptcore

import (
	gi "github.com/pekim/gobbi/internal/gi"
	"sync"
)

var getMajorVersionFunction *gi.Function
var getMajorVersionFunction_Once sync.Once

func getMajorVersionFunction_Set() error {
	var err error
	getMajorVersionFunction_Once.Do(func() {
		getMajorVersionFunction, err = gi.FunctionInvokerNew("JavaScriptCore", "get_major_version")
	})
	return err
}

// GetMajorVersion is a representation of the C type jsc_get_major_version.
func GetMajorVersion() (uint32, error) {

	var ret gi.Argument

	err := getMajorVersionFunction_Set()
	if err == nil {
		ret = getMajorVersionFunction.Invoke(nil, nil)
	}

	retGo := ret.Uint32()

	return retGo, err
}

var getMicroVersionFunction *gi.Function
var getMicroVersionFunction_Once sync.Once

func getMicroVersionFunction_Set() error {
	var err error
	getMicroVersionFunction_Once.Do(func() {
		getMicroVersionFunction, err = gi.FunctionInvokerNew("JavaScriptCore", "get_micro_version")
	})
	return err
}

// GetMicroVersion is a representation of the C type jsc_get_micro_version.
func GetMicroVersion() (uint32, error) {

	var ret gi.Argument

	err := getMicroVersionFunction_Set()
	if err == nil {
		ret = getMicroVersionFunction.Invoke(nil, nil)
	}

	retGo := ret.Uint32()

	return retGo, err
}

var getMinorVersionFunction *gi.Function
var getMinorVersionFunction_Once sync.Once

func getMinorVersionFunction_Set() error {
	var err error
	getMinorVersionFunction_Once.Do(func() {
		getMinorVersionFunction, err = gi.FunctionInvokerNew("JavaScriptCore", "get_minor_version")
	})
	return err
}

// GetMinorVersion is a representation of the C type jsc_get_minor_version.
func GetMinorVersion() (uint32, error) {

	var ret gi.Argument

	err := getMinorVersionFunction_Set()
	if err == nil {
		ret = getMinorVersionFunction.Invoke(nil, nil)
	}

	retGo := ret.Uint32()

	return retGo, err
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
