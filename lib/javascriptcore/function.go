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

var optionsGetBooleanFunction *gi.Function
var optionsGetBooleanFunction_Once sync.Once

func optionsGetBooleanFunction_Set() error {
	var err error
	optionsGetBooleanFunction_Once.Do(func() {
		optionsGetBooleanFunction, err = gi.FunctionInvokerNew("JavaScriptCore", "options_get_boolean")
	})
	return err
}

// OptionsGetBoolean is a representation of the C type jsc_options_get_boolean.
func OptionsGetBoolean(option string) (bool, bool, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(option)

	var outArgs [1]gi.Argument
	var ret gi.Argument

	err := optionsGetBooleanFunction_Set()
	if err == nil {
		ret = optionsGetBooleanFunction.Invoke(inArgs[:], outArgs[:])
	}

	retGo := ret.Boolean()
	out0 := outArgs[0].Boolean()

	return retGo, out0, err
}

var optionsGetDoubleFunction *gi.Function
var optionsGetDoubleFunction_Once sync.Once

func optionsGetDoubleFunction_Set() error {
	var err error
	optionsGetDoubleFunction_Once.Do(func() {
		optionsGetDoubleFunction, err = gi.FunctionInvokerNew("JavaScriptCore", "options_get_double")
	})
	return err
}

// OptionsGetDouble is a representation of the C type jsc_options_get_double.
func OptionsGetDouble(option string) (bool, float64, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(option)

	var outArgs [1]gi.Argument
	var ret gi.Argument

	err := optionsGetDoubleFunction_Set()
	if err == nil {
		ret = optionsGetDoubleFunction.Invoke(inArgs[:], outArgs[:])
	}

	retGo := ret.Boolean()
	out0 := outArgs[0].Float64()

	return retGo, out0, err
}

var optionsGetIntFunction *gi.Function
var optionsGetIntFunction_Once sync.Once

func optionsGetIntFunction_Set() error {
	var err error
	optionsGetIntFunction_Once.Do(func() {
		optionsGetIntFunction, err = gi.FunctionInvokerNew("JavaScriptCore", "options_get_int")
	})
	return err
}

// OptionsGetInt is a representation of the C type jsc_options_get_int.
func OptionsGetInt(option string) (bool, int32, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(option)

	var outArgs [1]gi.Argument
	var ret gi.Argument

	err := optionsGetIntFunction_Set()
	if err == nil {
		ret = optionsGetIntFunction.Invoke(inArgs[:], outArgs[:])
	}

	retGo := ret.Boolean()
	out0 := outArgs[0].Int32()

	return retGo, out0, err
}

// UNSUPPORTED : C value 'jsc_options_get_option_group' : return type 'GLib.OptionGroup' not supported

var optionsGetRangeStringFunction *gi.Function
var optionsGetRangeStringFunction_Once sync.Once

func optionsGetRangeStringFunction_Set() error {
	var err error
	optionsGetRangeStringFunction_Once.Do(func() {
		optionsGetRangeStringFunction, err = gi.FunctionInvokerNew("JavaScriptCore", "options_get_range_string")
	})
	return err
}

// OptionsGetRangeString is a representation of the C type jsc_options_get_range_string.
func OptionsGetRangeString(option string) (bool, string, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(option)

	var outArgs [1]gi.Argument
	var ret gi.Argument

	err := optionsGetRangeStringFunction_Set()
	if err == nil {
		ret = optionsGetRangeStringFunction.Invoke(inArgs[:], outArgs[:])
	}

	retGo := ret.Boolean()
	out0 := outArgs[0].String(true)

	return retGo, out0, err
}

var optionsGetSizeFunction *gi.Function
var optionsGetSizeFunction_Once sync.Once

func optionsGetSizeFunction_Set() error {
	var err error
	optionsGetSizeFunction_Once.Do(func() {
		optionsGetSizeFunction, err = gi.FunctionInvokerNew("JavaScriptCore", "options_get_size")
	})
	return err
}

// OptionsGetSize is a representation of the C type jsc_options_get_size.
func OptionsGetSize(option string) (bool, uint64, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(option)

	var outArgs [1]gi.Argument
	var ret gi.Argument

	err := optionsGetSizeFunction_Set()
	if err == nil {
		ret = optionsGetSizeFunction.Invoke(inArgs[:], outArgs[:])
	}

	retGo := ret.Boolean()
	out0 := outArgs[0].Uint64()

	return retGo, out0, err
}

var optionsGetStringFunction *gi.Function
var optionsGetStringFunction_Once sync.Once

func optionsGetStringFunction_Set() error {
	var err error
	optionsGetStringFunction_Once.Do(func() {
		optionsGetStringFunction, err = gi.FunctionInvokerNew("JavaScriptCore", "options_get_string")
	})
	return err
}

// OptionsGetString is a representation of the C type jsc_options_get_string.
func OptionsGetString(option string) (bool, string, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(option)

	var outArgs [1]gi.Argument
	var ret gi.Argument

	err := optionsGetStringFunction_Set()
	if err == nil {
		ret = optionsGetStringFunction.Invoke(inArgs[:], outArgs[:])
	}

	retGo := ret.Boolean()
	out0 := outArgs[0].String(true)

	return retGo, out0, err
}

var optionsGetUintFunction *gi.Function
var optionsGetUintFunction_Once sync.Once

func optionsGetUintFunction_Set() error {
	var err error
	optionsGetUintFunction_Once.Do(func() {
		optionsGetUintFunction, err = gi.FunctionInvokerNew("JavaScriptCore", "options_get_uint")
	})
	return err
}

// OptionsGetUint is a representation of the C type jsc_options_get_uint.
func OptionsGetUint(option string) (bool, uint32, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(option)

	var outArgs [1]gi.Argument
	var ret gi.Argument

	err := optionsGetUintFunction_Set()
	if err == nil {
		ret = optionsGetUintFunction.Invoke(inArgs[:], outArgs[:])
	}

	retGo := ret.Boolean()
	out0 := outArgs[0].Uint32()

	return retGo, out0, err
}

var optionsSetBooleanFunction *gi.Function
var optionsSetBooleanFunction_Once sync.Once

func optionsSetBooleanFunction_Set() error {
	var err error
	optionsSetBooleanFunction_Once.Do(func() {
		optionsSetBooleanFunction, err = gi.FunctionInvokerNew("JavaScriptCore", "options_set_boolean")
	})
	return err
}

// OptionsSetBoolean is a representation of the C type jsc_options_set_boolean.
func OptionsSetBoolean(option string, value bool) (bool, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(option)
	inArgs[1].SetBoolean(value)

	var ret gi.Argument

	err := optionsSetBooleanFunction_Set()
	if err == nil {
		ret = optionsSetBooleanFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

var optionsSetDoubleFunction *gi.Function
var optionsSetDoubleFunction_Once sync.Once

func optionsSetDoubleFunction_Set() error {
	var err error
	optionsSetDoubleFunction_Once.Do(func() {
		optionsSetDoubleFunction, err = gi.FunctionInvokerNew("JavaScriptCore", "options_set_double")
	})
	return err
}

// OptionsSetDouble is a representation of the C type jsc_options_set_double.
func OptionsSetDouble(option string, value float64) (bool, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(option)
	inArgs[1].SetFloat64(value)

	var ret gi.Argument

	err := optionsSetDoubleFunction_Set()
	if err == nil {
		ret = optionsSetDoubleFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

var optionsSetIntFunction *gi.Function
var optionsSetIntFunction_Once sync.Once

func optionsSetIntFunction_Set() error {
	var err error
	optionsSetIntFunction_Once.Do(func() {
		optionsSetIntFunction, err = gi.FunctionInvokerNew("JavaScriptCore", "options_set_int")
	})
	return err
}

// OptionsSetInt is a representation of the C type jsc_options_set_int.
func OptionsSetInt(option string, value int32) (bool, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(option)
	inArgs[1].SetInt32(value)

	var ret gi.Argument

	err := optionsSetIntFunction_Set()
	if err == nil {
		ret = optionsSetIntFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

var optionsSetRangeStringFunction *gi.Function
var optionsSetRangeStringFunction_Once sync.Once

func optionsSetRangeStringFunction_Set() error {
	var err error
	optionsSetRangeStringFunction_Once.Do(func() {
		optionsSetRangeStringFunction, err = gi.FunctionInvokerNew("JavaScriptCore", "options_set_range_string")
	})
	return err
}

// OptionsSetRangeString is a representation of the C type jsc_options_set_range_string.
func OptionsSetRangeString(option string, value string) (bool, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(option)
	inArgs[1].SetString(value)

	var ret gi.Argument

	err := optionsSetRangeStringFunction_Set()
	if err == nil {
		ret = optionsSetRangeStringFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

var optionsSetSizeFunction *gi.Function
var optionsSetSizeFunction_Once sync.Once

func optionsSetSizeFunction_Set() error {
	var err error
	optionsSetSizeFunction_Once.Do(func() {
		optionsSetSizeFunction, err = gi.FunctionInvokerNew("JavaScriptCore", "options_set_size")
	})
	return err
}

// OptionsSetSize is a representation of the C type jsc_options_set_size.
func OptionsSetSize(option string, value uint64) (bool, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(option)
	inArgs[1].SetUint64(value)

	var ret gi.Argument

	err := optionsSetSizeFunction_Set()
	if err == nil {
		ret = optionsSetSizeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

var optionsSetStringFunction *gi.Function
var optionsSetStringFunction_Once sync.Once

func optionsSetStringFunction_Set() error {
	var err error
	optionsSetStringFunction_Once.Do(func() {
		optionsSetStringFunction, err = gi.FunctionInvokerNew("JavaScriptCore", "options_set_string")
	})
	return err
}

// OptionsSetString is a representation of the C type jsc_options_set_string.
func OptionsSetString(option string, value string) (bool, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(option)
	inArgs[1].SetString(value)

	var ret gi.Argument

	err := optionsSetStringFunction_Set()
	if err == nil {
		ret = optionsSetStringFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

var optionsSetUintFunction *gi.Function
var optionsSetUintFunction_Once sync.Once

func optionsSetUintFunction_Set() error {
	var err error
	optionsSetUintFunction_Once.Do(func() {
		optionsSetUintFunction, err = gi.FunctionInvokerNew("JavaScriptCore", "options_set_uint")
	})
	return err
}

// OptionsSetUint is a representation of the C type jsc_options_set_uint.
func OptionsSetUint(option string, value uint32) (bool, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(option)
	inArgs[1].SetUint32(value)

	var ret gi.Argument

	err := optionsSetUintFunction_Set()
	if err == nil {
		ret = optionsSetUintFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}
