// Code generated - DO NOT EDIT.

package glib

import (
	gi "github.com/pekim/gobbi/internal/gi"
	"sync"
)

// UNSUPPORTED : C value 'g_access' : parameter 'filename' of type 'filename' not supported

var asciiDigitValueFunction *gi.Function
var asciiDigitValueFunction_Once sync.Once

func asciiDigitValueFunction_Set() error {
	var err error
	asciiDigitValueFunction_Once.Do(func() {
		asciiDigitValueFunction, err = gi.FunctionInvokerNew("GLib", "ascii_digit_value")
	})
	return err
}

// AsciiDigitValue is a representation of the C type g_ascii_digit_value.
func AsciiDigitValue(c int8) (int32, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetInt8(c)

	var ret gi.Argument

	err := asciiDigitValueFunction_Set()
	if err == nil {
		ret = asciiDigitValueFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo, err
}

var asciiDtostrFunction *gi.Function
var asciiDtostrFunction_Once sync.Once

func asciiDtostrFunction_Set() error {
	var err error
	asciiDtostrFunction_Once.Do(func() {
		asciiDtostrFunction, err = gi.FunctionInvokerNew("GLib", "ascii_dtostr")
	})
	return err
}

// AsciiDtostr is a representation of the C type g_ascii_dtostr.
func AsciiDtostr(buffer string, bufLen int32, d float64) (string, error) {
	var inArgs [3]gi.Argument
	inArgs[0].SetString(buffer)
	inArgs[1].SetInt32(bufLen)
	inArgs[2].SetFloat64(d)

	var ret gi.Argument

	err := asciiDtostrFunction_Set()
	if err == nil {
		ret = asciiDtostrFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo, err
}

var asciiFormatdFunction *gi.Function
var asciiFormatdFunction_Once sync.Once

func asciiFormatdFunction_Set() error {
	var err error
	asciiFormatdFunction_Once.Do(func() {
		asciiFormatdFunction, err = gi.FunctionInvokerNew("GLib", "ascii_formatd")
	})
	return err
}

// AsciiFormatd is a representation of the C type g_ascii_formatd.
func AsciiFormatd(buffer string, bufLen int32, format string, d float64) (string, error) {
	var inArgs [4]gi.Argument
	inArgs[0].SetString(buffer)
	inArgs[1].SetInt32(bufLen)
	inArgs[2].SetString(format)
	inArgs[3].SetFloat64(d)

	var ret gi.Argument

	err := asciiFormatdFunction_Set()
	if err == nil {
		ret = asciiFormatdFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo, err
}

var asciiStrcasecmpFunction *gi.Function
var asciiStrcasecmpFunction_Once sync.Once

func asciiStrcasecmpFunction_Set() error {
	var err error
	asciiStrcasecmpFunction_Once.Do(func() {
		asciiStrcasecmpFunction, err = gi.FunctionInvokerNew("GLib", "ascii_strcasecmp")
	})
	return err
}

// AsciiStrcasecmp is a representation of the C type g_ascii_strcasecmp.
func AsciiStrcasecmp(s1 string, s2 string) (int32, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(s1)
	inArgs[1].SetString(s2)

	var ret gi.Argument

	err := asciiStrcasecmpFunction_Set()
	if err == nil {
		ret = asciiStrcasecmpFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo, err
}

var asciiStrdownFunction *gi.Function
var asciiStrdownFunction_Once sync.Once

func asciiStrdownFunction_Set() error {
	var err error
	asciiStrdownFunction_Once.Do(func() {
		asciiStrdownFunction, err = gi.FunctionInvokerNew("GLib", "ascii_strdown")
	})
	return err
}

// AsciiStrdown is a representation of the C type g_ascii_strdown.
func AsciiStrdown(str string, len int32) (string, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(str)
	inArgs[1].SetInt32(len)

	var ret gi.Argument

	err := asciiStrdownFunction_Set()
	if err == nil {
		ret = asciiStrdownFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo, err
}

var asciiStringToSignedFunction *gi.Function
var asciiStringToSignedFunction_Once sync.Once

func asciiStringToSignedFunction_Set() error {
	var err error
	asciiStringToSignedFunction_Once.Do(func() {
		asciiStringToSignedFunction, err = gi.FunctionInvokerNew("GLib", "ascii_string_to_signed")
	})
	return err
}

// AsciiStringToSigned is a representation of the C type g_ascii_string_to_signed.
func AsciiStringToSigned(str string, base uint32, min int64, max int64) (bool, int64, error) {
	var inArgs [4]gi.Argument
	inArgs[0].SetString(str)
	inArgs[1].SetUint32(base)
	inArgs[2].SetInt64(min)
	inArgs[3].SetInt64(max)

	var outArgs [1]gi.Argument
	var ret gi.Argument

	err := asciiStringToSignedFunction_Set()
	if err == nil {
		ret = asciiStringToSignedFunction.Invoke(inArgs[:], outArgs[:])
	}

	retGo := ret.Boolean()
	out0 := outArgs[0].Int64()

	return retGo, out0, err
}

var asciiStringToUnsignedFunction *gi.Function
var asciiStringToUnsignedFunction_Once sync.Once

func asciiStringToUnsignedFunction_Set() error {
	var err error
	asciiStringToUnsignedFunction_Once.Do(func() {
		asciiStringToUnsignedFunction, err = gi.FunctionInvokerNew("GLib", "ascii_string_to_unsigned")
	})
	return err
}

// AsciiStringToUnsigned is a representation of the C type g_ascii_string_to_unsigned.
func AsciiStringToUnsigned(str string, base uint32, min uint64, max uint64) (bool, uint64, error) {
	var inArgs [4]gi.Argument
	inArgs[0].SetString(str)
	inArgs[1].SetUint32(base)
	inArgs[2].SetUint64(min)
	inArgs[3].SetUint64(max)

	var outArgs [1]gi.Argument
	var ret gi.Argument

	err := asciiStringToUnsignedFunction_Set()
	if err == nil {
		ret = asciiStringToUnsignedFunction.Invoke(inArgs[:], outArgs[:])
	}

	retGo := ret.Boolean()
	out0 := outArgs[0].Uint64()

	return retGo, out0, err
}

// UNSUPPORTED : C value 'g_ascii_strncasecmp' : parameter 'n' of type 'gsize' not supported

var asciiStrtodFunction *gi.Function
var asciiStrtodFunction_Once sync.Once

func asciiStrtodFunction_Set() error {
	var err error
	asciiStrtodFunction_Once.Do(func() {
		asciiStrtodFunction, err = gi.FunctionInvokerNew("GLib", "ascii_strtod")
	})
	return err
}

// AsciiStrtod is a representation of the C type g_ascii_strtod.
func AsciiStrtod(nptr string) (float64, string, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(nptr)

	var outArgs [1]gi.Argument
	var ret gi.Argument

	err := asciiStrtodFunction_Set()
	if err == nil {
		ret = asciiStrtodFunction.Invoke(inArgs[:], outArgs[:])
	}

	retGo := ret.Float64()
	out0 := outArgs[0].String(false)

	return retGo, out0, err
}

var asciiStrtollFunction *gi.Function
var asciiStrtollFunction_Once sync.Once

func asciiStrtollFunction_Set() error {
	var err error
	asciiStrtollFunction_Once.Do(func() {
		asciiStrtollFunction, err = gi.FunctionInvokerNew("GLib", "ascii_strtoll")
	})
	return err
}

// AsciiStrtoll is a representation of the C type g_ascii_strtoll.
func AsciiStrtoll(nptr string, base uint32) (int64, string, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(nptr)
	inArgs[1].SetUint32(base)

	var outArgs [1]gi.Argument
	var ret gi.Argument

	err := asciiStrtollFunction_Set()
	if err == nil {
		ret = asciiStrtollFunction.Invoke(inArgs[:], outArgs[:])
	}

	retGo := ret.Int64()
	out0 := outArgs[0].String(false)

	return retGo, out0, err
}

var asciiStrtoullFunction *gi.Function
var asciiStrtoullFunction_Once sync.Once

func asciiStrtoullFunction_Set() error {
	var err error
	asciiStrtoullFunction_Once.Do(func() {
		asciiStrtoullFunction, err = gi.FunctionInvokerNew("GLib", "ascii_strtoull")
	})
	return err
}

// AsciiStrtoull is a representation of the C type g_ascii_strtoull.
func AsciiStrtoull(nptr string, base uint32) (uint64, string, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(nptr)
	inArgs[1].SetUint32(base)

	var outArgs [1]gi.Argument
	var ret gi.Argument

	err := asciiStrtoullFunction_Set()
	if err == nil {
		ret = asciiStrtoullFunction.Invoke(inArgs[:], outArgs[:])
	}

	retGo := ret.Uint64()
	out0 := outArgs[0].String(false)

	return retGo, out0, err
}

var asciiStrupFunction *gi.Function
var asciiStrupFunction_Once sync.Once

func asciiStrupFunction_Set() error {
	var err error
	asciiStrupFunction_Once.Do(func() {
		asciiStrupFunction, err = gi.FunctionInvokerNew("GLib", "ascii_strup")
	})
	return err
}

// AsciiStrup is a representation of the C type g_ascii_strup.
func AsciiStrup(str string, len int32) (string, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(str)
	inArgs[1].SetInt32(len)

	var ret gi.Argument

	err := asciiStrupFunction_Set()
	if err == nil {
		ret = asciiStrupFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo, err
}

var asciiTolowerFunction *gi.Function
var asciiTolowerFunction_Once sync.Once

func asciiTolowerFunction_Set() error {
	var err error
	asciiTolowerFunction_Once.Do(func() {
		asciiTolowerFunction, err = gi.FunctionInvokerNew("GLib", "ascii_tolower")
	})
	return err
}

// AsciiTolower is a representation of the C type g_ascii_tolower.
func AsciiTolower(c int8) (int8, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetInt8(c)

	var ret gi.Argument

	err := asciiTolowerFunction_Set()
	if err == nil {
		ret = asciiTolowerFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int8()

	return retGo, err
}

var asciiToupperFunction *gi.Function
var asciiToupperFunction_Once sync.Once

func asciiToupperFunction_Set() error {
	var err error
	asciiToupperFunction_Once.Do(func() {
		asciiToupperFunction, err = gi.FunctionInvokerNew("GLib", "ascii_toupper")
	})
	return err
}

// AsciiToupper is a representation of the C type g_ascii_toupper.
func AsciiToupper(c int8) (int8, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetInt8(c)

	var ret gi.Argument

	err := asciiToupperFunction_Set()
	if err == nil {
		ret = asciiToupperFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int8()

	return retGo, err
}

var asciiXdigitValueFunction *gi.Function
var asciiXdigitValueFunction_Once sync.Once

func asciiXdigitValueFunction_Set() error {
	var err error
	asciiXdigitValueFunction_Once.Do(func() {
		asciiXdigitValueFunction, err = gi.FunctionInvokerNew("GLib", "ascii_xdigit_value")
	})
	return err
}

// AsciiXdigitValue is a representation of the C type g_ascii_xdigit_value.
func AsciiXdigitValue(c int8) (int32, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetInt8(c)

	var ret gi.Argument

	err := asciiXdigitValueFunction_Set()
	if err == nil {
		ret = asciiXdigitValueFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo, err
}

var assertWarningFunction *gi.Function
var assertWarningFunction_Once sync.Once

func assertWarningFunction_Set() error {
	var err error
	assertWarningFunction_Once.Do(func() {
		assertWarningFunction, err = gi.FunctionInvokerNew("GLib", "assert_warning")
	})
	return err
}

// AssertWarning is a representation of the C type g_assert_warning.
func AssertWarning(logDomain string, file string, line int32, prettyFunction string, expression string) error {
	var inArgs [5]gi.Argument
	inArgs[0].SetString(logDomain)
	inArgs[1].SetString(file)
	inArgs[2].SetInt32(line)
	inArgs[3].SetString(prettyFunction)
	inArgs[4].SetString(expression)

	err := assertWarningFunction_Set()
	if err == nil {
		assertWarningFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var assertionMessageFunction *gi.Function
var assertionMessageFunction_Once sync.Once

func assertionMessageFunction_Set() error {
	var err error
	assertionMessageFunction_Once.Do(func() {
		assertionMessageFunction, err = gi.FunctionInvokerNew("GLib", "assertion_message")
	})
	return err
}

// AssertionMessage is a representation of the C type g_assertion_message.
func AssertionMessage(domain string, file string, line int32, func_ string, message string) error {
	var inArgs [5]gi.Argument
	inArgs[0].SetString(domain)
	inArgs[1].SetString(file)
	inArgs[2].SetInt32(line)
	inArgs[3].SetString(func_)
	inArgs[4].SetString(message)

	err := assertionMessageFunction_Set()
	if err == nil {
		assertionMessageFunction.Invoke(inArgs[:], nil)
	}

	return err
}

// UNSUPPORTED : C value 'g_assertion_message_cmpnum' : parameter 'arg1' of type 'long double' not supported

var assertionMessageCmpstrFunction *gi.Function
var assertionMessageCmpstrFunction_Once sync.Once

func assertionMessageCmpstrFunction_Set() error {
	var err error
	assertionMessageCmpstrFunction_Once.Do(func() {
		assertionMessageCmpstrFunction, err = gi.FunctionInvokerNew("GLib", "assertion_message_cmpstr")
	})
	return err
}

// AssertionMessageCmpstr is a representation of the C type g_assertion_message_cmpstr.
func AssertionMessageCmpstr(domain string, file string, line int32, func_ string, expr string, arg1 string, cmp string, arg2 string) error {
	var inArgs [8]gi.Argument
	inArgs[0].SetString(domain)
	inArgs[1].SetString(file)
	inArgs[2].SetInt32(line)
	inArgs[3].SetString(func_)
	inArgs[4].SetString(expr)
	inArgs[5].SetString(arg1)
	inArgs[6].SetString(cmp)
	inArgs[7].SetString(arg2)

	err := assertionMessageCmpstrFunction_Set()
	if err == nil {
		assertionMessageCmpstrFunction.Invoke(inArgs[:], nil)
	}

	return err
}

// UNSUPPORTED : C value 'g_assertion_message_error' : parameter 'error' of type 'Error' not supported

var assertionMessageExprFunction *gi.Function
var assertionMessageExprFunction_Once sync.Once

func assertionMessageExprFunction_Set() error {
	var err error
	assertionMessageExprFunction_Once.Do(func() {
		assertionMessageExprFunction, err = gi.FunctionInvokerNew("GLib", "assertion_message_expr")
	})
	return err
}

// AssertionMessageExpr is a representation of the C type g_assertion_message_expr.
func AssertionMessageExpr(domain string, file string, line int32, func_ string, expr string) error {
	var inArgs [5]gi.Argument
	inArgs[0].SetString(domain)
	inArgs[1].SetString(file)
	inArgs[2].SetInt32(line)
	inArgs[3].SetString(func_)
	inArgs[4].SetString(expr)

	err := assertionMessageExprFunction_Set()
	if err == nil {
		assertionMessageExprFunction.Invoke(inArgs[:], nil)
	}

	return err
}

// UNSUPPORTED : C value 'g_atexit' : parameter 'func' of type 'VoidFunc' not supported

var atomicIntAddFunction *gi.Function
var atomicIntAddFunction_Once sync.Once

func atomicIntAddFunction_Set() error {
	var err error
	atomicIntAddFunction_Once.Do(func() {
		atomicIntAddFunction, err = gi.FunctionInvokerNew("GLib", "atomic_int_add")
	})
	return err
}

// AtomicIntAdd is a representation of the C type g_atomic_int_add.
func AtomicIntAdd(atomic int32, val int32) (int32, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetInt32(atomic)
	inArgs[1].SetInt32(val)

	var ret gi.Argument

	err := atomicIntAddFunction_Set()
	if err == nil {
		ret = atomicIntAddFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo, err
}

var atomicIntAndFunction *gi.Function
var atomicIntAndFunction_Once sync.Once

func atomicIntAndFunction_Set() error {
	var err error
	atomicIntAndFunction_Once.Do(func() {
		atomicIntAndFunction, err = gi.FunctionInvokerNew("GLib", "atomic_int_and")
	})
	return err
}

// AtomicIntAnd is a representation of the C type g_atomic_int_and.
func AtomicIntAnd(atomic uint32, val uint32) (uint32, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetUint32(atomic)
	inArgs[1].SetUint32(val)

	var ret gi.Argument

	err := atomicIntAndFunction_Set()
	if err == nil {
		ret = atomicIntAndFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo, err
}

var atomicIntCompareAndExchangeFunction *gi.Function
var atomicIntCompareAndExchangeFunction_Once sync.Once

func atomicIntCompareAndExchangeFunction_Set() error {
	var err error
	atomicIntCompareAndExchangeFunction_Once.Do(func() {
		atomicIntCompareAndExchangeFunction, err = gi.FunctionInvokerNew("GLib", "atomic_int_compare_and_exchange")
	})
	return err
}

// AtomicIntCompareAndExchange is a representation of the C type g_atomic_int_compare_and_exchange.
func AtomicIntCompareAndExchange(atomic int32, oldval int32, newval int32) (bool, error) {
	var inArgs [3]gi.Argument
	inArgs[0].SetInt32(atomic)
	inArgs[1].SetInt32(oldval)
	inArgs[2].SetInt32(newval)

	var ret gi.Argument

	err := atomicIntCompareAndExchangeFunction_Set()
	if err == nil {
		ret = atomicIntCompareAndExchangeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

var atomicIntDecAndTestFunction *gi.Function
var atomicIntDecAndTestFunction_Once sync.Once

func atomicIntDecAndTestFunction_Set() error {
	var err error
	atomicIntDecAndTestFunction_Once.Do(func() {
		atomicIntDecAndTestFunction, err = gi.FunctionInvokerNew("GLib", "atomic_int_dec_and_test")
	})
	return err
}

// AtomicIntDecAndTest is a representation of the C type g_atomic_int_dec_and_test.
func AtomicIntDecAndTest(atomic int32) (bool, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetInt32(atomic)

	var ret gi.Argument

	err := atomicIntDecAndTestFunction_Set()
	if err == nil {
		ret = atomicIntDecAndTestFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

var atomicIntExchangeAndAddFunction *gi.Function
var atomicIntExchangeAndAddFunction_Once sync.Once

func atomicIntExchangeAndAddFunction_Set() error {
	var err error
	atomicIntExchangeAndAddFunction_Once.Do(func() {
		atomicIntExchangeAndAddFunction, err = gi.FunctionInvokerNew("GLib", "atomic_int_exchange_and_add")
	})
	return err
}

// AtomicIntExchangeAndAdd is a representation of the C type g_atomic_int_exchange_and_add.
func AtomicIntExchangeAndAdd(atomic int32, val int32) (int32, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetInt32(atomic)
	inArgs[1].SetInt32(val)

	var ret gi.Argument

	err := atomicIntExchangeAndAddFunction_Set()
	if err == nil {
		ret = atomicIntExchangeAndAddFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo, err
}

var atomicIntGetFunction *gi.Function
var atomicIntGetFunction_Once sync.Once

func atomicIntGetFunction_Set() error {
	var err error
	atomicIntGetFunction_Once.Do(func() {
		atomicIntGetFunction, err = gi.FunctionInvokerNew("GLib", "atomic_int_get")
	})
	return err
}

// AtomicIntGet is a representation of the C type g_atomic_int_get.
func AtomicIntGet(atomic int32) (int32, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetInt32(atomic)

	var ret gi.Argument

	err := atomicIntGetFunction_Set()
	if err == nil {
		ret = atomicIntGetFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo, err
}

var atomicIntIncFunction *gi.Function
var atomicIntIncFunction_Once sync.Once

func atomicIntIncFunction_Set() error {
	var err error
	atomicIntIncFunction_Once.Do(func() {
		atomicIntIncFunction, err = gi.FunctionInvokerNew("GLib", "atomic_int_inc")
	})
	return err
}

// AtomicIntInc is a representation of the C type g_atomic_int_inc.
func AtomicIntInc(atomic int32) error {
	var inArgs [1]gi.Argument
	inArgs[0].SetInt32(atomic)

	err := atomicIntIncFunction_Set()
	if err == nil {
		atomicIntIncFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var atomicIntOrFunction *gi.Function
var atomicIntOrFunction_Once sync.Once

func atomicIntOrFunction_Set() error {
	var err error
	atomicIntOrFunction_Once.Do(func() {
		atomicIntOrFunction, err = gi.FunctionInvokerNew("GLib", "atomic_int_or")
	})
	return err
}

// AtomicIntOr is a representation of the C type g_atomic_int_or.
func AtomicIntOr(atomic uint32, val uint32) (uint32, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetUint32(atomic)
	inArgs[1].SetUint32(val)

	var ret gi.Argument

	err := atomicIntOrFunction_Set()
	if err == nil {
		ret = atomicIntOrFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo, err
}

var atomicIntSetFunction *gi.Function
var atomicIntSetFunction_Once sync.Once

func atomicIntSetFunction_Set() error {
	var err error
	atomicIntSetFunction_Once.Do(func() {
		atomicIntSetFunction, err = gi.FunctionInvokerNew("GLib", "atomic_int_set")
	})
	return err
}

// AtomicIntSet is a representation of the C type g_atomic_int_set.
func AtomicIntSet(atomic int32, newval int32) error {
	var inArgs [2]gi.Argument
	inArgs[0].SetInt32(atomic)
	inArgs[1].SetInt32(newval)

	err := atomicIntSetFunction_Set()
	if err == nil {
		atomicIntSetFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var atomicIntXorFunction *gi.Function
var atomicIntXorFunction_Once sync.Once

func atomicIntXorFunction_Set() error {
	var err error
	atomicIntXorFunction_Once.Do(func() {
		atomicIntXorFunction, err = gi.FunctionInvokerNew("GLib", "atomic_int_xor")
	})
	return err
}

// AtomicIntXor is a representation of the C type g_atomic_int_xor.
func AtomicIntXor(atomic uint32, val uint32) (uint32, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetUint32(atomic)
	inArgs[1].SetUint32(val)

	var ret gi.Argument

	err := atomicIntXorFunction_Set()
	if err == nil {
		ret = atomicIntXorFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo, err
}

// UNSUPPORTED : C value 'g_atomic_pointer_add' : parameter 'atomic' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_atomic_pointer_and' : parameter 'atomic' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_atomic_pointer_compare_and_exchange' : parameter 'atomic' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_atomic_pointer_get' : parameter 'atomic' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_atomic_pointer_or' : parameter 'atomic' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_atomic_pointer_set' : parameter 'atomic' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_atomic_pointer_xor' : parameter 'atomic' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_atomic_rc_box_acquire' : parameter 'mem_block' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_atomic_rc_box_alloc' : parameter 'block_size' of type 'gsize' not supported

// UNSUPPORTED : C value 'g_atomic_rc_box_alloc0' : parameter 'block_size' of type 'gsize' not supported

// UNSUPPORTED : C value 'g_atomic_rc_box_dup' : parameter 'block_size' of type 'gsize' not supported

// UNSUPPORTED : C value 'g_atomic_rc_box_get_size' : parameter 'mem_block' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_atomic_rc_box_release' : parameter 'mem_block' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_atomic_rc_box_release_full' : parameter 'mem_block' of type 'gpointer' not supported

var atomicRefCountCompareFunction *gi.Function
var atomicRefCountCompareFunction_Once sync.Once

func atomicRefCountCompareFunction_Set() error {
	var err error
	atomicRefCountCompareFunction_Once.Do(func() {
		atomicRefCountCompareFunction, err = gi.FunctionInvokerNew("GLib", "atomic_ref_count_compare")
	})
	return err
}

// AtomicRefCountCompare is a representation of the C type g_atomic_ref_count_compare.
func AtomicRefCountCompare(arc int32, val int32) (bool, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetInt32(arc)
	inArgs[1].SetInt32(val)

	var ret gi.Argument

	err := atomicRefCountCompareFunction_Set()
	if err == nil {
		ret = atomicRefCountCompareFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

var atomicRefCountDecFunction *gi.Function
var atomicRefCountDecFunction_Once sync.Once

func atomicRefCountDecFunction_Set() error {
	var err error
	atomicRefCountDecFunction_Once.Do(func() {
		atomicRefCountDecFunction, err = gi.FunctionInvokerNew("GLib", "atomic_ref_count_dec")
	})
	return err
}

// AtomicRefCountDec is a representation of the C type g_atomic_ref_count_dec.
func AtomicRefCountDec(arc int32) (bool, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetInt32(arc)

	var ret gi.Argument

	err := atomicRefCountDecFunction_Set()
	if err == nil {
		ret = atomicRefCountDecFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

var atomicRefCountIncFunction *gi.Function
var atomicRefCountIncFunction_Once sync.Once

func atomicRefCountIncFunction_Set() error {
	var err error
	atomicRefCountIncFunction_Once.Do(func() {
		atomicRefCountIncFunction, err = gi.FunctionInvokerNew("GLib", "atomic_ref_count_inc")
	})
	return err
}

// AtomicRefCountInc is a representation of the C type g_atomic_ref_count_inc.
func AtomicRefCountInc(arc int32) error {
	var inArgs [1]gi.Argument
	inArgs[0].SetInt32(arc)

	err := atomicRefCountIncFunction_Set()
	if err == nil {
		atomicRefCountIncFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var atomicRefCountInitFunction *gi.Function
var atomicRefCountInitFunction_Once sync.Once

func atomicRefCountInitFunction_Set() error {
	var err error
	atomicRefCountInitFunction_Once.Do(func() {
		atomicRefCountInitFunction, err = gi.FunctionInvokerNew("GLib", "atomic_ref_count_init")
	})
	return err
}

// AtomicRefCountInit is a representation of the C type g_atomic_ref_count_init.
func AtomicRefCountInit(arc int32) error {
	var inArgs [1]gi.Argument
	inArgs[0].SetInt32(arc)

	err := atomicRefCountInitFunction_Set()
	if err == nil {
		atomicRefCountInitFunction.Invoke(inArgs[:], nil)
	}

	return err
}

// UNSUPPORTED : C value 'g_base64_decode' : parameter 'out_len' of type 'gsize' not supported

// UNSUPPORTED : C value 'g_base64_decode_inplace' : parameter 'text' has no type

// UNSUPPORTED : C value 'g_base64_decode_step' : parameter 'in' has no type

// UNSUPPORTED : C value 'g_base64_encode' : parameter 'data' has no type

// UNSUPPORTED : C value 'g_base64_encode_close' : parameter 'out' has no type

// UNSUPPORTED : C value 'g_base64_encode_step' : parameter 'in' has no type

// UNSUPPORTED : C value 'g_basename' : parameter 'file_name' of type 'filename' not supported

var bitLockFunction *gi.Function
var bitLockFunction_Once sync.Once

func bitLockFunction_Set() error {
	var err error
	bitLockFunction_Once.Do(func() {
		bitLockFunction, err = gi.FunctionInvokerNew("GLib", "bit_lock")
	})
	return err
}

// BitLock is a representation of the C type g_bit_lock.
func BitLock(address int32, lockBit int32) error {
	var inArgs [2]gi.Argument
	inArgs[0].SetInt32(address)
	inArgs[1].SetInt32(lockBit)

	err := bitLockFunction_Set()
	if err == nil {
		bitLockFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var bitNthLsfFunction *gi.Function
var bitNthLsfFunction_Once sync.Once

func bitNthLsfFunction_Set() error {
	var err error
	bitNthLsfFunction_Once.Do(func() {
		bitNthLsfFunction, err = gi.FunctionInvokerNew("GLib", "bit_nth_lsf")
	})
	return err
}

// BitNthLsf is a representation of the C type g_bit_nth_lsf.
func BitNthLsf(mask uint64, nthBit int32) (int32, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetUint64(mask)
	inArgs[1].SetInt32(nthBit)

	var ret gi.Argument

	err := bitNthLsfFunction_Set()
	if err == nil {
		ret = bitNthLsfFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo, err
}

var bitNthMsfFunction *gi.Function
var bitNthMsfFunction_Once sync.Once

func bitNthMsfFunction_Set() error {
	var err error
	bitNthMsfFunction_Once.Do(func() {
		bitNthMsfFunction, err = gi.FunctionInvokerNew("GLib", "bit_nth_msf")
	})
	return err
}

// BitNthMsf is a representation of the C type g_bit_nth_msf.
func BitNthMsf(mask uint64, nthBit int32) (int32, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetUint64(mask)
	inArgs[1].SetInt32(nthBit)

	var ret gi.Argument

	err := bitNthMsfFunction_Set()
	if err == nil {
		ret = bitNthMsfFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo, err
}

var bitStorageFunction *gi.Function
var bitStorageFunction_Once sync.Once

func bitStorageFunction_Set() error {
	var err error
	bitStorageFunction_Once.Do(func() {
		bitStorageFunction, err = gi.FunctionInvokerNew("GLib", "bit_storage")
	})
	return err
}

// BitStorage is a representation of the C type g_bit_storage.
func BitStorage(number uint64) (uint32, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetUint64(number)

	var ret gi.Argument

	err := bitStorageFunction_Set()
	if err == nil {
		ret = bitStorageFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo, err
}

var bitTrylockFunction *gi.Function
var bitTrylockFunction_Once sync.Once

func bitTrylockFunction_Set() error {
	var err error
	bitTrylockFunction_Once.Do(func() {
		bitTrylockFunction, err = gi.FunctionInvokerNew("GLib", "bit_trylock")
	})
	return err
}

// BitTrylock is a representation of the C type g_bit_trylock.
func BitTrylock(address int32, lockBit int32) (bool, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetInt32(address)
	inArgs[1].SetInt32(lockBit)

	var ret gi.Argument

	err := bitTrylockFunction_Set()
	if err == nil {
		ret = bitTrylockFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

var bitUnlockFunction *gi.Function
var bitUnlockFunction_Once sync.Once

func bitUnlockFunction_Set() error {
	var err error
	bitUnlockFunction_Once.Do(func() {
		bitUnlockFunction, err = gi.FunctionInvokerNew("GLib", "bit_unlock")
	})
	return err
}

// BitUnlock is a representation of the C type g_bit_unlock.
func BitUnlock(address int32, lockBit int32) error {
	var inArgs [2]gi.Argument
	inArgs[0].SetInt32(address)
	inArgs[1].SetInt32(lockBit)

	err := bitUnlockFunction_Set()
	if err == nil {
		bitUnlockFunction.Invoke(inArgs[:], nil)
	}

	return err
}

// UNSUPPORTED : C value 'g_bookmark_file_error_quark' : return type 'Quark' not supported

// UNSUPPORTED : C value 'g_build_filename' : parameter 'first_element' of type 'filename' not supported

// UNSUPPORTED : C value 'g_build_filename_valist' : parameter 'first_element' of type 'filename' not supported

// UNSUPPORTED : C value 'g_build_filenamev' : parameter 'args' has no type

// UNSUPPORTED : C value 'g_build_path' : parameter 'separator' of type 'filename' not supported

// UNSUPPORTED : C value 'g_build_pathv' : parameter 'args' has no type

// UNSUPPORTED : C value 'g_byte_array_free' : parameter 'array' has no type

// UNSUPPORTED : C value 'g_byte_array_free_to_bytes' : parameter 'array' has no type

var byteArrayNewFunction *gi.Function
var byteArrayNewFunction_Once sync.Once

func byteArrayNewFunction_Set() error {
	var err error
	byteArrayNewFunction_Once.Do(func() {
		byteArrayNewFunction, err = gi.FunctionInvokerNew("GLib", "byte_array_new")
	})
	return err
}

// ByteArrayNew is a representation of the C type g_byte_array_new.
func ByteArrayNew() error {

	err := byteArrayNewFunction_Set()
	if err == nil {
		byteArrayNewFunction.Invoke(nil, nil)
	}

	return err
}

// UNSUPPORTED : C value 'g_byte_array_new_take' : parameter 'data' has no type

// UNSUPPORTED : C value 'g_byte_array_unref' : parameter 'array' has no type

// UNSUPPORTED : C value 'g_canonicalize_filename' : parameter 'filename' of type 'filename' not supported

// UNSUPPORTED : C value 'g_chdir' : parameter 'path' of type 'filename' not supported

var checkVersionFunction *gi.Function
var checkVersionFunction_Once sync.Once

func checkVersionFunction_Set() error {
	var err error
	checkVersionFunction_Once.Do(func() {
		checkVersionFunction, err = gi.FunctionInvokerNew("GLib", "check_version")
	})
	return err
}

// CheckVersion is a representation of the C type glib_check_version.
func CheckVersion(requiredMajor uint32, requiredMinor uint32, requiredMicro uint32) (string, error) {
	var inArgs [3]gi.Argument
	inArgs[0].SetUint32(requiredMajor)
	inArgs[1].SetUint32(requiredMinor)
	inArgs[2].SetUint32(requiredMicro)

	var ret gi.Argument

	err := checkVersionFunction_Set()
	if err == nil {
		ret = checkVersionFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo, err
}

// UNSUPPORTED : C value 'g_checksum_type_get_length' : parameter 'checksum_type' of type 'ChecksumType' not supported

// UNSUPPORTED : C value 'g_child_watch_add' : parameter 'function' of type 'ChildWatchFunc' not supported

// UNSUPPORTED : C value 'g_child_watch_add_full' : parameter 'function' of type 'ChildWatchFunc' not supported

var childWatchSourceNewFunction *gi.Function
var childWatchSourceNewFunction_Once sync.Once

func childWatchSourceNewFunction_Set() error {
	var err error
	childWatchSourceNewFunction_Once.Do(func() {
		childWatchSourceNewFunction, err = gi.FunctionInvokerNew("GLib", "child_watch_source_new")
	})
	return err
}

// ChildWatchSourceNew is a representation of the C type g_child_watch_source_new.
func ChildWatchSourceNew(pid Pid) (*Source, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetInt32(int32(pid))

	var ret gi.Argument

	err := childWatchSourceNewFunction_Set()
	if err == nil {
		ret = childWatchSourceNewFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Source{native: ret.Pointer()}

	return retGo, err
}

var clearErrorFunction *gi.Function
var clearErrorFunction_Once sync.Once

func clearErrorFunction_Set() error {
	var err error
	clearErrorFunction_Once.Do(func() {
		clearErrorFunction, err = gi.FunctionInvokerNew("GLib", "clear_error")
	})
	return err
}

// ClearError is a representation of the C type g_clear_error.
func ClearError() error {

	err := clearErrorFunction_Set()
	if err == nil {
		clearErrorFunction.Invoke(nil, nil)
	}

	return err
}

// UNSUPPORTED : C value 'g_clear_handle_id' : parameter 'clear_func' of type 'ClearHandleFunc' not supported

// UNSUPPORTED : C value 'g_clear_pointer' : parameter 'pp' of type 'gpointer' not supported

var closeFunction *gi.Function
var closeFunction_Once sync.Once

func closeFunction_Set() error {
	var err error
	closeFunction_Once.Do(func() {
		closeFunction, err = gi.FunctionInvokerNew("GLib", "close")
	})
	return err
}

// Close is a representation of the C type g_close.
func Close(fd int32) (bool, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetInt32(fd)

	var ret gi.Argument

	err := closeFunction_Set()
	if err == nil {
		ret = closeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

// UNSUPPORTED : C value 'g_compute_checksum_for_bytes' : parameter 'checksum_type' of type 'ChecksumType' not supported

// UNSUPPORTED : C value 'g_compute_checksum_for_data' : parameter 'checksum_type' of type 'ChecksumType' not supported

// UNSUPPORTED : C value 'g_compute_checksum_for_string' : parameter 'checksum_type' of type 'ChecksumType' not supported

// UNSUPPORTED : C value 'g_compute_hmac_for_bytes' : parameter 'digest_type' of type 'ChecksumType' not supported

// UNSUPPORTED : C value 'g_compute_hmac_for_data' : parameter 'digest_type' of type 'ChecksumType' not supported

// UNSUPPORTED : C value 'g_compute_hmac_for_string' : parameter 'digest_type' of type 'ChecksumType' not supported

// UNSUPPORTED : C value 'g_convert' : parameter 'str' has no type

// UNSUPPORTED : C value 'g_convert_error_quark' : return type 'Quark' not supported

// UNSUPPORTED : C value 'g_convert_with_fallback' : parameter 'str' has no type

// UNSUPPORTED : C value 'g_convert_with_iconv' : parameter 'str' has no type

// UNSUPPORTED : C value 'g_datalist_clear' : parameter 'datalist' of type 'Data' not supported

// UNSUPPORTED : C value 'g_datalist_foreach' : parameter 'datalist' of type 'Data' not supported

// UNSUPPORTED : C value 'g_datalist_get_data' : parameter 'datalist' of type 'Data' not supported

// UNSUPPORTED : C value 'g_datalist_get_flags' : parameter 'datalist' of type 'Data' not supported

// UNSUPPORTED : C value 'g_datalist_id_dup_data' : parameter 'datalist' of type 'Data' not supported

// UNSUPPORTED : C value 'g_datalist_id_get_data' : parameter 'datalist' of type 'Data' not supported

// UNSUPPORTED : C value 'g_datalist_id_remove_no_notify' : parameter 'datalist' of type 'Data' not supported

// UNSUPPORTED : C value 'g_datalist_id_replace_data' : parameter 'datalist' of type 'Data' not supported

// UNSUPPORTED : C value 'g_datalist_id_set_data_full' : parameter 'datalist' of type 'Data' not supported

// UNSUPPORTED : C value 'g_datalist_init' : parameter 'datalist' of type 'Data' not supported

// UNSUPPORTED : C value 'g_datalist_set_flags' : parameter 'datalist' of type 'Data' not supported

// UNSUPPORTED : C value 'g_datalist_unset_flags' : parameter 'datalist' of type 'Data' not supported

// UNSUPPORTED : C value 'g_dataset_destroy' : parameter 'dataset_location' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_dataset_foreach' : parameter 'dataset_location' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_dataset_id_get_data' : parameter 'dataset_location' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_dataset_id_remove_no_notify' : parameter 'dataset_location' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_dataset_id_set_data_full' : parameter 'dataset_location' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_date_get_days_in_month' : parameter 'month' of type 'DateMonth' not supported

var dateGetMondayWeeksInYearFunction *gi.Function
var dateGetMondayWeeksInYearFunction_Once sync.Once

func dateGetMondayWeeksInYearFunction_Set() error {
	var err error
	dateGetMondayWeeksInYearFunction_Once.Do(func() {
		dateGetMondayWeeksInYearFunction, err = gi.FunctionInvokerNew("GLib", "date_get_monday_weeks_in_year")
	})
	return err
}

// DateGetMondayWeeksInYear is a representation of the C type g_date_get_monday_weeks_in_year.
func DateGetMondayWeeksInYear(year DateYear) (uint8, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetUint16(uint16(year))

	var ret gi.Argument

	err := dateGetMondayWeeksInYearFunction_Set()
	if err == nil {
		ret = dateGetMondayWeeksInYearFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint8()

	return retGo, err
}

var dateGetSundayWeeksInYearFunction *gi.Function
var dateGetSundayWeeksInYearFunction_Once sync.Once

func dateGetSundayWeeksInYearFunction_Set() error {
	var err error
	dateGetSundayWeeksInYearFunction_Once.Do(func() {
		dateGetSundayWeeksInYearFunction, err = gi.FunctionInvokerNew("GLib", "date_get_sunday_weeks_in_year")
	})
	return err
}

// DateGetSundayWeeksInYear is a representation of the C type g_date_get_sunday_weeks_in_year.
func DateGetSundayWeeksInYear(year DateYear) (uint8, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetUint16(uint16(year))

	var ret gi.Argument

	err := dateGetSundayWeeksInYearFunction_Set()
	if err == nil {
		ret = dateGetSundayWeeksInYearFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint8()

	return retGo, err
}

var dateIsLeapYearFunction *gi.Function
var dateIsLeapYearFunction_Once sync.Once

func dateIsLeapYearFunction_Set() error {
	var err error
	dateIsLeapYearFunction_Once.Do(func() {
		dateIsLeapYearFunction, err = gi.FunctionInvokerNew("GLib", "date_is_leap_year")
	})
	return err
}

// DateIsLeapYear is a representation of the C type g_date_is_leap_year.
func DateIsLeapYear(year DateYear) (bool, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetUint16(uint16(year))

	var ret gi.Argument

	err := dateIsLeapYearFunction_Set()
	if err == nil {
		ret = dateIsLeapYearFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

// UNSUPPORTED : C value 'g_date_strftime' : parameter 'slen' of type 'gsize' not supported

// UNSUPPORTED : C value 'g_date_time_compare' : parameter 'dt1' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_date_time_equal' : parameter 'dt1' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_date_time_hash' : parameter 'datetime' of type 'gpointer' not supported

var dateValidDayFunction *gi.Function
var dateValidDayFunction_Once sync.Once

func dateValidDayFunction_Set() error {
	var err error
	dateValidDayFunction_Once.Do(func() {
		dateValidDayFunction, err = gi.FunctionInvokerNew("GLib", "date_valid_day")
	})
	return err
}

// DateValidDay is a representation of the C type g_date_valid_day.
func DateValidDay(day DateDay) (bool, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetUint8(uint8(day))

	var ret gi.Argument

	err := dateValidDayFunction_Set()
	if err == nil {
		ret = dateValidDayFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

// UNSUPPORTED : C value 'g_date_valid_dmy' : parameter 'month' of type 'DateMonth' not supported

var dateValidJulianFunction *gi.Function
var dateValidJulianFunction_Once sync.Once

func dateValidJulianFunction_Set() error {
	var err error
	dateValidJulianFunction_Once.Do(func() {
		dateValidJulianFunction, err = gi.FunctionInvokerNew("GLib", "date_valid_julian")
	})
	return err
}

// DateValidJulian is a representation of the C type g_date_valid_julian.
func DateValidJulian(julianDate uint32) (bool, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetUint32(julianDate)

	var ret gi.Argument

	err := dateValidJulianFunction_Set()
	if err == nil {
		ret = dateValidJulianFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

// UNSUPPORTED : C value 'g_date_valid_month' : parameter 'month' of type 'DateMonth' not supported

// UNSUPPORTED : C value 'g_date_valid_weekday' : parameter 'weekday' of type 'DateWeekday' not supported

var dateValidYearFunction *gi.Function
var dateValidYearFunction_Once sync.Once

func dateValidYearFunction_Set() error {
	var err error
	dateValidYearFunction_Once.Do(func() {
		dateValidYearFunction, err = gi.FunctionInvokerNew("GLib", "date_valid_year")
	})
	return err
}

// DateValidYear is a representation of the C type g_date_valid_year.
func DateValidYear(year DateYear) (bool, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetUint16(uint16(year))

	var ret gi.Argument

	err := dateValidYearFunction_Set()
	if err == nil {
		ret = dateValidYearFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

var dcgettextFunction *gi.Function
var dcgettextFunction_Once sync.Once

func dcgettextFunction_Set() error {
	var err error
	dcgettextFunction_Once.Do(func() {
		dcgettextFunction, err = gi.FunctionInvokerNew("GLib", "dcgettext")
	})
	return err
}

// Dcgettext is a representation of the C type g_dcgettext.
func Dcgettext(domain string, msgid string, category int32) (string, error) {
	var inArgs [3]gi.Argument
	inArgs[0].SetString(domain)
	inArgs[1].SetString(msgid)
	inArgs[2].SetInt32(category)

	var ret gi.Argument

	err := dcgettextFunction_Set()
	if err == nil {
		ret = dcgettextFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo, err
}

var dgettextFunction *gi.Function
var dgettextFunction_Once sync.Once

func dgettextFunction_Set() error {
	var err error
	dgettextFunction_Once.Do(func() {
		dgettextFunction, err = gi.FunctionInvokerNew("GLib", "dgettext")
	})
	return err
}

// Dgettext is a representation of the C type g_dgettext.
func Dgettext(domain string, msgid string) (string, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(domain)
	inArgs[1].SetString(msgid)

	var ret gi.Argument

	err := dgettextFunction_Set()
	if err == nil {
		ret = dgettextFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo, err
}

// UNSUPPORTED : C value 'g_dir_make_tmp' : parameter 'tmpl' of type 'filename' not supported

// UNSUPPORTED : C value 'g_direct_equal' : parameter 'v1' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_direct_hash' : parameter 'v' of type 'gpointer' not supported

var dngettextFunction *gi.Function
var dngettextFunction_Once sync.Once

func dngettextFunction_Set() error {
	var err error
	dngettextFunction_Once.Do(func() {
		dngettextFunction, err = gi.FunctionInvokerNew("GLib", "dngettext")
	})
	return err
}

// Dngettext is a representation of the C type g_dngettext.
func Dngettext(domain string, msgid string, msgidPlural string, n uint64) (string, error) {
	var inArgs [4]gi.Argument
	inArgs[0].SetString(domain)
	inArgs[1].SetString(msgid)
	inArgs[2].SetString(msgidPlural)
	inArgs[3].SetUint64(n)

	var ret gi.Argument

	err := dngettextFunction_Set()
	if err == nil {
		ret = dngettextFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo, err
}

// UNSUPPORTED : C value 'g_double_equal' : parameter 'v1' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_double_hash' : parameter 'v' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_dpgettext' : parameter 'msgidoffset' of type 'gsize' not supported

var dpgettext2Function *gi.Function
var dpgettext2Function_Once sync.Once

func dpgettext2Function_Set() error {
	var err error
	dpgettext2Function_Once.Do(func() {
		dpgettext2Function, err = gi.FunctionInvokerNew("GLib", "dpgettext2")
	})
	return err
}

// Dpgettext2 is a representation of the C type g_dpgettext2.
func Dpgettext2(domain string, context string, msgid string) (string, error) {
	var inArgs [3]gi.Argument
	inArgs[0].SetString(domain)
	inArgs[1].SetString(context)
	inArgs[2].SetString(msgid)

	var ret gi.Argument

	err := dpgettext2Function_Set()
	if err == nil {
		ret = dpgettext2Function.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo, err
}

// UNSUPPORTED : C value 'g_environ_getenv' : parameter 'envp' has no type

// UNSUPPORTED : C value 'g_environ_setenv' : parameter 'envp' has no type

// UNSUPPORTED : C value 'g_environ_unsetenv' : parameter 'envp' has no type

// UNSUPPORTED : C value 'g_file_error_from_errno' : return type 'FileError' not supported

// UNSUPPORTED : C value 'g_file_error_quark' : return type 'Quark' not supported

// UNSUPPORTED : C value 'g_file_get_contents' : parameter 'filename' of type 'filename' not supported

// UNSUPPORTED : C value 'g_file_open_tmp' : parameter 'tmpl' of type 'filename' not supported

// UNSUPPORTED : C value 'g_file_read_link' : parameter 'filename' of type 'filename' not supported

// UNSUPPORTED : C value 'g_file_set_contents' : parameter 'filename' of type 'filename' not supported

// UNSUPPORTED : C value 'g_file_test' : parameter 'filename' of type 'filename' not supported

// UNSUPPORTED : C value 'g_filename_display_basename' : parameter 'filename' of type 'filename' not supported

// UNSUPPORTED : C value 'g_filename_display_name' : parameter 'filename' of type 'filename' not supported

// UNSUPPORTED : C value 'g_filename_from_uri' : return type 'filename' not supported

// UNSUPPORTED : C value 'g_filename_from_utf8' : parameter 'bytes_read' of type 'gsize' not supported

// UNSUPPORTED : C value 'g_filename_to_uri' : parameter 'filename' of type 'filename' not supported

// UNSUPPORTED : C value 'g_filename_to_utf8' : parameter 'opsysstring' of type 'filename' not supported

// UNSUPPORTED : C value 'g_find_program_in_path' : parameter 'program' of type 'filename' not supported

var formatSizeFunction *gi.Function
var formatSizeFunction_Once sync.Once

func formatSizeFunction_Set() error {
	var err error
	formatSizeFunction_Once.Do(func() {
		formatSizeFunction, err = gi.FunctionInvokerNew("GLib", "format_size")
	})
	return err
}

// FormatSize is a representation of the C type g_format_size.
func FormatSize(size uint64) (string, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetUint64(size)

	var ret gi.Argument

	err := formatSizeFunction_Set()
	if err == nil {
		ret = formatSizeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo, err
}

var formatSizeForDisplayFunction *gi.Function
var formatSizeForDisplayFunction_Once sync.Once

func formatSizeForDisplayFunction_Set() error {
	var err error
	formatSizeForDisplayFunction_Once.Do(func() {
		formatSizeForDisplayFunction, err = gi.FunctionInvokerNew("GLib", "format_size_for_display")
	})
	return err
}

// FormatSizeForDisplay is a representation of the C type g_format_size_for_display.
func FormatSizeForDisplay(size int64) (string, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetInt64(size)

	var ret gi.Argument

	err := formatSizeForDisplayFunction_Set()
	if err == nil {
		ret = formatSizeForDisplayFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo, err
}

// UNSUPPORTED : C value 'g_format_size_full' : parameter 'flags' of type 'FormatSizeFlags' not supported

// UNSUPPORTED : C value 'g_fprintf' : parameter 'file' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_free' : parameter 'mem' of type 'gpointer' not supported

var getApplicationNameFunction *gi.Function
var getApplicationNameFunction_Once sync.Once

func getApplicationNameFunction_Set() error {
	var err error
	getApplicationNameFunction_Once.Do(func() {
		getApplicationNameFunction, err = gi.FunctionInvokerNew("GLib", "get_application_name")
	})
	return err
}

// GetApplicationName is a representation of the C type g_get_application_name.
func GetApplicationName() (string, error) {

	var ret gi.Argument

	err := getApplicationNameFunction_Set()
	if err == nil {
		ret = getApplicationNameFunction.Invoke(nil, nil)
	}

	retGo := ret.String(false)

	return retGo, err
}

var getCharsetFunction *gi.Function
var getCharsetFunction_Once sync.Once

func getCharsetFunction_Set() error {
	var err error
	getCharsetFunction_Once.Do(func() {
		getCharsetFunction, err = gi.FunctionInvokerNew("GLib", "get_charset")
	})
	return err
}

// GetCharset is a representation of the C type g_get_charset.
func GetCharset() (bool, string, error) {

	var outArgs [1]gi.Argument
	var ret gi.Argument

	err := getCharsetFunction_Set()
	if err == nil {
		ret = getCharsetFunction.Invoke(nil, outArgs[:])
	}

	retGo := ret.Boolean()
	out0 := outArgs[0].String(false)

	return retGo, out0, err
}

var getCodesetFunction *gi.Function
var getCodesetFunction_Once sync.Once

func getCodesetFunction_Set() error {
	var err error
	getCodesetFunction_Once.Do(func() {
		getCodesetFunction, err = gi.FunctionInvokerNew("GLib", "get_codeset")
	})
	return err
}

// GetCodeset is a representation of the C type g_get_codeset.
func GetCodeset() (string, error) {

	var ret gi.Argument

	err := getCodesetFunction_Set()
	if err == nil {
		ret = getCodesetFunction.Invoke(nil, nil)
	}

	retGo := ret.String(true)

	return retGo, err
}

var getConsoleCharsetFunction *gi.Function
var getConsoleCharsetFunction_Once sync.Once

func getConsoleCharsetFunction_Set() error {
	var err error
	getConsoleCharsetFunction_Once.Do(func() {
		getConsoleCharsetFunction, err = gi.FunctionInvokerNew("GLib", "get_console_charset")
	})
	return err
}

// GetConsoleCharset is a representation of the C type g_get_console_charset.
func GetConsoleCharset() (bool, string, error) {

	var outArgs [1]gi.Argument
	var ret gi.Argument

	err := getConsoleCharsetFunction_Set()
	if err == nil {
		ret = getConsoleCharsetFunction.Invoke(nil, outArgs[:])
	}

	retGo := ret.Boolean()
	out0 := outArgs[0].String(false)

	return retGo, out0, err
}

// UNSUPPORTED : C value 'g_get_current_dir' : return type 'filename' not supported

// UNSUPPORTED : C value 'g_get_current_time' : parameter 'result' of type 'TimeVal' not supported

var getEnvironFunction *gi.Function
var getEnvironFunction_Once sync.Once

func getEnvironFunction_Set() error {
	var err error
	getEnvironFunction_Once.Do(func() {
		getEnvironFunction, err = gi.FunctionInvokerNew("GLib", "get_environ")
	})
	return err
}

// GetEnviron is a representation of the C type g_get_environ.
func GetEnviron() error {

	err := getEnvironFunction_Set()
	if err == nil {
		getEnvironFunction.Invoke(nil, nil)
	}

	return err
}

// UNSUPPORTED : C value 'g_get_filename_charsets' : parameter 'filename_charsets' has no type

// UNSUPPORTED : C value 'g_get_home_dir' : return type 'filename' not supported

var getHostNameFunction *gi.Function
var getHostNameFunction_Once sync.Once

func getHostNameFunction_Set() error {
	var err error
	getHostNameFunction_Once.Do(func() {
		getHostNameFunction, err = gi.FunctionInvokerNew("GLib", "get_host_name")
	})
	return err
}

// GetHostName is a representation of the C type g_get_host_name.
func GetHostName() (string, error) {

	var ret gi.Argument

	err := getHostNameFunction_Set()
	if err == nil {
		ret = getHostNameFunction.Invoke(nil, nil)
	}

	retGo := ret.String(false)

	return retGo, err
}

var getLanguageNamesFunction *gi.Function
var getLanguageNamesFunction_Once sync.Once

func getLanguageNamesFunction_Set() error {
	var err error
	getLanguageNamesFunction_Once.Do(func() {
		getLanguageNamesFunction, err = gi.FunctionInvokerNew("GLib", "get_language_names")
	})
	return err
}

// GetLanguageNames is a representation of the C type g_get_language_names.
func GetLanguageNames() error {

	err := getLanguageNamesFunction_Set()
	if err == nil {
		getLanguageNamesFunction.Invoke(nil, nil)
	}

	return err
}

var getLanguageNamesWithCategoryFunction *gi.Function
var getLanguageNamesWithCategoryFunction_Once sync.Once

func getLanguageNamesWithCategoryFunction_Set() error {
	var err error
	getLanguageNamesWithCategoryFunction_Once.Do(func() {
		getLanguageNamesWithCategoryFunction, err = gi.FunctionInvokerNew("GLib", "get_language_names_with_category")
	})
	return err
}

// GetLanguageNamesWithCategory is a representation of the C type g_get_language_names_with_category.
func GetLanguageNamesWithCategory(categoryName string) error {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(categoryName)

	err := getLanguageNamesWithCategoryFunction_Set()
	if err == nil {
		getLanguageNamesWithCategoryFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var getLocaleVariantsFunction *gi.Function
var getLocaleVariantsFunction_Once sync.Once

func getLocaleVariantsFunction_Set() error {
	var err error
	getLocaleVariantsFunction_Once.Do(func() {
		getLocaleVariantsFunction, err = gi.FunctionInvokerNew("GLib", "get_locale_variants")
	})
	return err
}

// GetLocaleVariants is a representation of the C type g_get_locale_variants.
func GetLocaleVariants(locale string) error {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(locale)

	err := getLocaleVariantsFunction_Set()
	if err == nil {
		getLocaleVariantsFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var getMonotonicTimeFunction *gi.Function
var getMonotonicTimeFunction_Once sync.Once

func getMonotonicTimeFunction_Set() error {
	var err error
	getMonotonicTimeFunction_Once.Do(func() {
		getMonotonicTimeFunction, err = gi.FunctionInvokerNew("GLib", "get_monotonic_time")
	})
	return err
}

// GetMonotonicTime is a representation of the C type g_get_monotonic_time.
func GetMonotonicTime() (int64, error) {

	var ret gi.Argument

	err := getMonotonicTimeFunction_Set()
	if err == nil {
		ret = getMonotonicTimeFunction.Invoke(nil, nil)
	}

	retGo := ret.Int64()

	return retGo, err
}

var getNumProcessorsFunction *gi.Function
var getNumProcessorsFunction_Once sync.Once

func getNumProcessorsFunction_Set() error {
	var err error
	getNumProcessorsFunction_Once.Do(func() {
		getNumProcessorsFunction, err = gi.FunctionInvokerNew("GLib", "get_num_processors")
	})
	return err
}

// GetNumProcessors is a representation of the C type g_get_num_processors.
func GetNumProcessors() (uint32, error) {

	var ret gi.Argument

	err := getNumProcessorsFunction_Set()
	if err == nil {
		ret = getNumProcessorsFunction.Invoke(nil, nil)
	}

	retGo := ret.Uint32()

	return retGo, err
}

var getPrgnameFunction *gi.Function
var getPrgnameFunction_Once sync.Once

func getPrgnameFunction_Set() error {
	var err error
	getPrgnameFunction_Once.Do(func() {
		getPrgnameFunction, err = gi.FunctionInvokerNew("GLib", "get_prgname")
	})
	return err
}

// GetPrgname is a representation of the C type g_get_prgname.
func GetPrgname() (string, error) {

	var ret gi.Argument

	err := getPrgnameFunction_Set()
	if err == nil {
		ret = getPrgnameFunction.Invoke(nil, nil)
	}

	retGo := ret.String(false)

	return retGo, err
}

// UNSUPPORTED : C value 'g_get_real_name' : return type 'filename' not supported

var getRealTimeFunction *gi.Function
var getRealTimeFunction_Once sync.Once

func getRealTimeFunction_Set() error {
	var err error
	getRealTimeFunction_Once.Do(func() {
		getRealTimeFunction, err = gi.FunctionInvokerNew("GLib", "get_real_time")
	})
	return err
}

// GetRealTime is a representation of the C type g_get_real_time.
func GetRealTime() (int64, error) {

	var ret gi.Argument

	err := getRealTimeFunction_Set()
	if err == nil {
		ret = getRealTimeFunction.Invoke(nil, nil)
	}

	retGo := ret.Int64()

	return retGo, err
}

var getSystemConfigDirsFunction *gi.Function
var getSystemConfigDirsFunction_Once sync.Once

func getSystemConfigDirsFunction_Set() error {
	var err error
	getSystemConfigDirsFunction_Once.Do(func() {
		getSystemConfigDirsFunction, err = gi.FunctionInvokerNew("GLib", "get_system_config_dirs")
	})
	return err
}

// GetSystemConfigDirs is a representation of the C type g_get_system_config_dirs.
func GetSystemConfigDirs() error {

	err := getSystemConfigDirsFunction_Set()
	if err == nil {
		getSystemConfigDirsFunction.Invoke(nil, nil)
	}

	return err
}

var getSystemDataDirsFunction *gi.Function
var getSystemDataDirsFunction_Once sync.Once

func getSystemDataDirsFunction_Set() error {
	var err error
	getSystemDataDirsFunction_Once.Do(func() {
		getSystemDataDirsFunction, err = gi.FunctionInvokerNew("GLib", "get_system_data_dirs")
	})
	return err
}

// GetSystemDataDirs is a representation of the C type g_get_system_data_dirs.
func GetSystemDataDirs() error {

	err := getSystemDataDirsFunction_Set()
	if err == nil {
		getSystemDataDirsFunction.Invoke(nil, nil)
	}

	return err
}

// UNSUPPORTED : C value 'g_get_tmp_dir' : return type 'filename' not supported

// UNSUPPORTED : C value 'g_get_user_cache_dir' : return type 'filename' not supported

// UNSUPPORTED : C value 'g_get_user_config_dir' : return type 'filename' not supported

// UNSUPPORTED : C value 'g_get_user_data_dir' : return type 'filename' not supported

// UNSUPPORTED : C value 'g_get_user_name' : return type 'filename' not supported

// UNSUPPORTED : C value 'g_get_user_runtime_dir' : return type 'filename' not supported

// UNSUPPORTED : C value 'g_get_user_special_dir' : parameter 'directory' of type 'UserDirectory' not supported

// UNSUPPORTED : C value 'g_getenv' : parameter 'variable' of type 'filename' not supported

// UNSUPPORTED : C value 'g_hash_table_add' : parameter 'hash_table' of type 'GLib.HashTable' not supported

// UNSUPPORTED : C value 'g_hash_table_contains' : parameter 'hash_table' of type 'GLib.HashTable' not supported

// UNSUPPORTED : C value 'g_hash_table_destroy' : parameter 'hash_table' of type 'GLib.HashTable' not supported

// UNSUPPORTED : C value 'g_hash_table_insert' : parameter 'hash_table' of type 'GLib.HashTable' not supported

// UNSUPPORTED : C value 'g_hash_table_lookup' : parameter 'hash_table' of type 'GLib.HashTable' not supported

// UNSUPPORTED : C value 'g_hash_table_lookup_extended' : parameter 'hash_table' of type 'GLib.HashTable' not supported

// UNSUPPORTED : C value 'g_hash_table_remove' : parameter 'hash_table' of type 'GLib.HashTable' not supported

// UNSUPPORTED : C value 'g_hash_table_remove_all' : parameter 'hash_table' of type 'GLib.HashTable' not supported

// UNSUPPORTED : C value 'g_hash_table_replace' : parameter 'hash_table' of type 'GLib.HashTable' not supported

// UNSUPPORTED : C value 'g_hash_table_size' : parameter 'hash_table' of type 'GLib.HashTable' not supported

// UNSUPPORTED : C value 'g_hash_table_steal' : parameter 'hash_table' of type 'GLib.HashTable' not supported

// UNSUPPORTED : C value 'g_hash_table_steal_all' : parameter 'hash_table' of type 'GLib.HashTable' not supported

// UNSUPPORTED : C value 'g_hash_table_steal_extended' : parameter 'hash_table' of type 'GLib.HashTable' not supported

// UNSUPPORTED : C value 'g_hash_table_unref' : parameter 'hash_table' of type 'GLib.HashTable' not supported

// UNSUPPORTED : C value 'g_hook_destroy' : parameter 'hook_list' of type 'HookList' not supported

// UNSUPPORTED : C value 'g_hook_destroy_link' : parameter 'hook_list' of type 'HookList' not supported

// UNSUPPORTED : C value 'g_hook_free' : parameter 'hook_list' of type 'HookList' not supported

// UNSUPPORTED : C value 'g_hook_insert_before' : parameter 'hook_list' of type 'HookList' not supported

// UNSUPPORTED : C value 'g_hook_prepend' : parameter 'hook_list' of type 'HookList' not supported

// UNSUPPORTED : C value 'g_hook_unref' : parameter 'hook_list' of type 'HookList' not supported

var hostnameIsAsciiEncodedFunction *gi.Function
var hostnameIsAsciiEncodedFunction_Once sync.Once

func hostnameIsAsciiEncodedFunction_Set() error {
	var err error
	hostnameIsAsciiEncodedFunction_Once.Do(func() {
		hostnameIsAsciiEncodedFunction, err = gi.FunctionInvokerNew("GLib", "hostname_is_ascii_encoded")
	})
	return err
}

// HostnameIsAsciiEncoded is a representation of the C type g_hostname_is_ascii_encoded.
func HostnameIsAsciiEncoded(hostname string) (bool, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(hostname)

	var ret gi.Argument

	err := hostnameIsAsciiEncodedFunction_Set()
	if err == nil {
		ret = hostnameIsAsciiEncodedFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

var hostnameIsIpAddressFunction *gi.Function
var hostnameIsIpAddressFunction_Once sync.Once

func hostnameIsIpAddressFunction_Set() error {
	var err error
	hostnameIsIpAddressFunction_Once.Do(func() {
		hostnameIsIpAddressFunction, err = gi.FunctionInvokerNew("GLib", "hostname_is_ip_address")
	})
	return err
}

// HostnameIsIpAddress is a representation of the C type g_hostname_is_ip_address.
func HostnameIsIpAddress(hostname string) (bool, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(hostname)

	var ret gi.Argument

	err := hostnameIsIpAddressFunction_Set()
	if err == nil {
		ret = hostnameIsIpAddressFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

var hostnameIsNonAsciiFunction *gi.Function
var hostnameIsNonAsciiFunction_Once sync.Once

func hostnameIsNonAsciiFunction_Set() error {
	var err error
	hostnameIsNonAsciiFunction_Once.Do(func() {
		hostnameIsNonAsciiFunction, err = gi.FunctionInvokerNew("GLib", "hostname_is_non_ascii")
	})
	return err
}

// HostnameIsNonAscii is a representation of the C type g_hostname_is_non_ascii.
func HostnameIsNonAscii(hostname string) (bool, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(hostname)

	var ret gi.Argument

	err := hostnameIsNonAsciiFunction_Set()
	if err == nil {
		ret = hostnameIsNonAsciiFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

var hostnameToAsciiFunction *gi.Function
var hostnameToAsciiFunction_Once sync.Once

func hostnameToAsciiFunction_Set() error {
	var err error
	hostnameToAsciiFunction_Once.Do(func() {
		hostnameToAsciiFunction, err = gi.FunctionInvokerNew("GLib", "hostname_to_ascii")
	})
	return err
}

// HostnameToAscii is a representation of the C type g_hostname_to_ascii.
func HostnameToAscii(hostname string) (string, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(hostname)

	var ret gi.Argument

	err := hostnameToAsciiFunction_Set()
	if err == nil {
		ret = hostnameToAsciiFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo, err
}

var hostnameToUnicodeFunction *gi.Function
var hostnameToUnicodeFunction_Once sync.Once

func hostnameToUnicodeFunction_Set() error {
	var err error
	hostnameToUnicodeFunction_Once.Do(func() {
		hostnameToUnicodeFunction, err = gi.FunctionInvokerNew("GLib", "hostname_to_unicode")
	})
	return err
}

// HostnameToUnicode is a representation of the C type g_hostname_to_unicode.
func HostnameToUnicode(hostname string) (string, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(hostname)

	var ret gi.Argument

	err := hostnameToUnicodeFunction_Set()
	if err == nil {
		ret = hostnameToUnicodeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo, err
}

// UNSUPPORTED : C value 'g_iconv' : parameter 'converter' of type 'IConv' not supported

var iconvOpenFunction *gi.Function
var iconvOpenFunction_Once sync.Once

func iconvOpenFunction_Set() error {
	var err error
	iconvOpenFunction_Once.Do(func() {
		iconvOpenFunction, err = gi.FunctionInvokerNew("GLib", "iconv_open")
	})
	return err
}

// IconvOpen is a representation of the C type g_iconv_open.
func IconvOpen(toCodeset string, fromCodeset string) (*IConv, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(toCodeset)
	inArgs[1].SetString(fromCodeset)

	var ret gi.Argument

	err := iconvOpenFunction_Set()
	if err == nil {
		ret = iconvOpenFunction.Invoke(inArgs[:], nil)
	}

	retGo := &IConv{native: ret.Pointer()}

	return retGo, err
}

// UNSUPPORTED : C value 'g_idle_add' : parameter 'function' of type 'SourceFunc' not supported

// UNSUPPORTED : C value 'g_idle_add_full' : parameter 'function' of type 'SourceFunc' not supported

// UNSUPPORTED : C value 'g_idle_remove_by_data' : parameter 'data' of type 'gpointer' not supported

var idleSourceNewFunction *gi.Function
var idleSourceNewFunction_Once sync.Once

func idleSourceNewFunction_Set() error {
	var err error
	idleSourceNewFunction_Once.Do(func() {
		idleSourceNewFunction, err = gi.FunctionInvokerNew("GLib", "idle_source_new")
	})
	return err
}

// IdleSourceNew is a representation of the C type g_idle_source_new.
func IdleSourceNew() (*Source, error) {

	var ret gi.Argument

	err := idleSourceNewFunction_Set()
	if err == nil {
		ret = idleSourceNewFunction.Invoke(nil, nil)
	}

	retGo := &Source{native: ret.Pointer()}

	return retGo, err
}

// UNSUPPORTED : C value 'g_int64_equal' : parameter 'v1' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_int64_hash' : parameter 'v' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_int_equal' : parameter 'v1' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_int_hash' : parameter 'v' of type 'gpointer' not supported

var internStaticStringFunction *gi.Function
var internStaticStringFunction_Once sync.Once

func internStaticStringFunction_Set() error {
	var err error
	internStaticStringFunction_Once.Do(func() {
		internStaticStringFunction, err = gi.FunctionInvokerNew("GLib", "intern_static_string")
	})
	return err
}

// InternStaticString is a representation of the C type g_intern_static_string.
func InternStaticString(string_ string) (string, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(string_)

	var ret gi.Argument

	err := internStaticStringFunction_Set()
	if err == nil {
		ret = internStaticStringFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo, err
}

var internStringFunction *gi.Function
var internStringFunction_Once sync.Once

func internStringFunction_Set() error {
	var err error
	internStringFunction_Once.Do(func() {
		internStringFunction, err = gi.FunctionInvokerNew("GLib", "intern_string")
	})
	return err
}

// InternString is a representation of the C type g_intern_string.
func InternString(string_ string) (string, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(string_)

	var ret gi.Argument

	err := internStringFunction_Set()
	if err == nil {
		ret = internStringFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo, err
}

// UNSUPPORTED : C value 'g_io_add_watch' : parameter 'channel' of type 'IOChannel' not supported

// UNSUPPORTED : C value 'g_io_add_watch_full' : parameter 'channel' of type 'IOChannel' not supported

// UNSUPPORTED : C value 'g_io_channel_error_from_errno' : return type 'IOChannelError' not supported

// UNSUPPORTED : C value 'g_io_channel_error_quark' : return type 'Quark' not supported

// UNSUPPORTED : C value 'g_io_create_watch' : parameter 'channel' of type 'IOChannel' not supported

// UNSUPPORTED : C value 'g_key_file_error_quark' : return type 'Quark' not supported

var listenvFunction *gi.Function
var listenvFunction_Once sync.Once

func listenvFunction_Set() error {
	var err error
	listenvFunction_Once.Do(func() {
		listenvFunction, err = gi.FunctionInvokerNew("GLib", "listenv")
	})
	return err
}

// Listenv is a representation of the C type g_listenv.
func Listenv() error {

	err := listenvFunction_Set()
	if err == nil {
		listenvFunction.Invoke(nil, nil)
	}

	return err
}

// UNSUPPORTED : C value 'g_locale_from_utf8' : parameter 'bytes_read' of type 'gsize' not supported

// UNSUPPORTED : C value 'g_locale_to_utf8' : parameter 'opsysstring' has no type

// UNSUPPORTED : C value 'g_log' : parameter 'log_level' of type 'LogLevelFlags' not supported

// UNSUPPORTED : C value 'g_log_default_handler' : parameter 'log_level' of type 'LogLevelFlags' not supported

var logRemoveHandlerFunction *gi.Function
var logRemoveHandlerFunction_Once sync.Once

func logRemoveHandlerFunction_Set() error {
	var err error
	logRemoveHandlerFunction_Once.Do(func() {
		logRemoveHandlerFunction, err = gi.FunctionInvokerNew("GLib", "log_remove_handler")
	})
	return err
}

// LogRemoveHandler is a representation of the C type g_log_remove_handler.
func LogRemoveHandler(logDomain string, handlerId uint32) error {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(logDomain)
	inArgs[1].SetUint32(handlerId)

	err := logRemoveHandlerFunction_Set()
	if err == nil {
		logRemoveHandlerFunction.Invoke(inArgs[:], nil)
	}

	return err
}

// UNSUPPORTED : C value 'g_log_set_always_fatal' : parameter 'fatal_mask' of type 'LogLevelFlags' not supported

// UNSUPPORTED : C value 'g_log_set_default_handler' : parameter 'log_func' of type 'LogFunc' not supported

// UNSUPPORTED : C value 'g_log_set_fatal_mask' : parameter 'fatal_mask' of type 'LogLevelFlags' not supported

// UNSUPPORTED : C value 'g_log_set_handler' : parameter 'log_levels' of type 'LogLevelFlags' not supported

// UNSUPPORTED : C value 'g_log_set_handler_full' : parameter 'log_levels' of type 'LogLevelFlags' not supported

// UNSUPPORTED : C value 'g_log_set_writer_func' : parameter 'func' of type 'LogWriterFunc' not supported

// UNSUPPORTED : C value 'g_log_structured' : parameter 'log_level' of type 'LogLevelFlags' not supported

// UNSUPPORTED : C value 'g_log_structured_array' : parameter 'log_level' of type 'LogLevelFlags' not supported

// UNSUPPORTED : C value 'g_log_structured_standard' : parameter 'log_level' of type 'LogLevelFlags' not supported

// UNSUPPORTED : C value 'g_log_variant' : parameter 'log_level' of type 'LogLevelFlags' not supported

// UNSUPPORTED : C value 'g_log_writer_default' : parameter 'log_level' of type 'LogLevelFlags' not supported

// UNSUPPORTED : C value 'g_log_writer_format_fields' : parameter 'log_level' of type 'LogLevelFlags' not supported

var logWriterIsJournaldFunction *gi.Function
var logWriterIsJournaldFunction_Once sync.Once

func logWriterIsJournaldFunction_Set() error {
	var err error
	logWriterIsJournaldFunction_Once.Do(func() {
		logWriterIsJournaldFunction, err = gi.FunctionInvokerNew("GLib", "log_writer_is_journald")
	})
	return err
}

// LogWriterIsJournald is a representation of the C type g_log_writer_is_journald.
func LogWriterIsJournald(outputFd int32) (bool, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetInt32(outputFd)

	var ret gi.Argument

	err := logWriterIsJournaldFunction_Set()
	if err == nil {
		ret = logWriterIsJournaldFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

// UNSUPPORTED : C value 'g_log_writer_journald' : parameter 'log_level' of type 'LogLevelFlags' not supported

// UNSUPPORTED : C value 'g_log_writer_standard_streams' : parameter 'log_level' of type 'LogLevelFlags' not supported

var logWriterSupportsColorFunction *gi.Function
var logWriterSupportsColorFunction_Once sync.Once

func logWriterSupportsColorFunction_Set() error {
	var err error
	logWriterSupportsColorFunction_Once.Do(func() {
		logWriterSupportsColorFunction, err = gi.FunctionInvokerNew("GLib", "log_writer_supports_color")
	})
	return err
}

// LogWriterSupportsColor is a representation of the C type g_log_writer_supports_color.
func LogWriterSupportsColor(outputFd int32) (bool, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetInt32(outputFd)

	var ret gi.Argument

	err := logWriterSupportsColorFunction_Set()
	if err == nil {
		ret = logWriterSupportsColorFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

// UNSUPPORTED : C value 'g_logv' : parameter 'log_level' of type 'LogLevelFlags' not supported

var mainContextDefaultFunction *gi.Function
var mainContextDefaultFunction_Once sync.Once

func mainContextDefaultFunction_Set() error {
	var err error
	mainContextDefaultFunction_Once.Do(func() {
		mainContextDefaultFunction, err = gi.FunctionInvokerNew("GLib", "main_context_default")
	})
	return err
}

// MainContextDefault is a representation of the C type g_main_context_default.
func MainContextDefault() (*MainContext, error) {

	var ret gi.Argument

	err := mainContextDefaultFunction_Set()
	if err == nil {
		ret = mainContextDefaultFunction.Invoke(nil, nil)
	}

	retGo := &MainContext{native: ret.Pointer()}

	return retGo, err
}

var mainContextGetThreadDefaultFunction *gi.Function
var mainContextGetThreadDefaultFunction_Once sync.Once

func mainContextGetThreadDefaultFunction_Set() error {
	var err error
	mainContextGetThreadDefaultFunction_Once.Do(func() {
		mainContextGetThreadDefaultFunction, err = gi.FunctionInvokerNew("GLib", "main_context_get_thread_default")
	})
	return err
}

// MainContextGetThreadDefault is a representation of the C type g_main_context_get_thread_default.
func MainContextGetThreadDefault() (*MainContext, error) {

	var ret gi.Argument

	err := mainContextGetThreadDefaultFunction_Set()
	if err == nil {
		ret = mainContextGetThreadDefaultFunction.Invoke(nil, nil)
	}

	retGo := &MainContext{native: ret.Pointer()}

	return retGo, err
}

var mainContextRefThreadDefaultFunction *gi.Function
var mainContextRefThreadDefaultFunction_Once sync.Once

func mainContextRefThreadDefaultFunction_Set() error {
	var err error
	mainContextRefThreadDefaultFunction_Once.Do(func() {
		mainContextRefThreadDefaultFunction, err = gi.FunctionInvokerNew("GLib", "main_context_ref_thread_default")
	})
	return err
}

// MainContextRefThreadDefault is a representation of the C type g_main_context_ref_thread_default.
func MainContextRefThreadDefault() (*MainContext, error) {

	var ret gi.Argument

	err := mainContextRefThreadDefaultFunction_Set()
	if err == nil {
		ret = mainContextRefThreadDefaultFunction.Invoke(nil, nil)
	}

	retGo := &MainContext{native: ret.Pointer()}

	return retGo, err
}

var mainCurrentSourceFunction *gi.Function
var mainCurrentSourceFunction_Once sync.Once

func mainCurrentSourceFunction_Set() error {
	var err error
	mainCurrentSourceFunction_Once.Do(func() {
		mainCurrentSourceFunction, err = gi.FunctionInvokerNew("GLib", "main_current_source")
	})
	return err
}

// MainCurrentSource is a representation of the C type g_main_current_source.
func MainCurrentSource() (*Source, error) {

	var ret gi.Argument

	err := mainCurrentSourceFunction_Set()
	if err == nil {
		ret = mainCurrentSourceFunction.Invoke(nil, nil)
	}

	retGo := &Source{native: ret.Pointer()}

	return retGo, err
}

var mainDepthFunction *gi.Function
var mainDepthFunction_Once sync.Once

func mainDepthFunction_Set() error {
	var err error
	mainDepthFunction_Once.Do(func() {
		mainDepthFunction, err = gi.FunctionInvokerNew("GLib", "main_depth")
	})
	return err
}

// MainDepth is a representation of the C type g_main_depth.
func MainDepth() (int32, error) {

	var ret gi.Argument

	err := mainDepthFunction_Set()
	if err == nil {
		ret = mainDepthFunction.Invoke(nil, nil)
	}

	retGo := ret.Int32()

	return retGo, err
}

// UNSUPPORTED : C value 'g_malloc' : parameter 'n_bytes' of type 'gsize' not supported

// UNSUPPORTED : C value 'g_malloc0' : parameter 'n_bytes' of type 'gsize' not supported

// UNSUPPORTED : C value 'g_malloc0_n' : parameter 'n_blocks' of type 'gsize' not supported

// UNSUPPORTED : C value 'g_malloc_n' : parameter 'n_blocks' of type 'gsize' not supported

// UNSUPPORTED : C value 'g_markup_collect_attributes' : parameter 'error' of type 'Error' not supported

// UNSUPPORTED : C value 'g_markup_error_quark' : return type 'Quark' not supported

var markupEscapeTextFunction *gi.Function
var markupEscapeTextFunction_Once sync.Once

func markupEscapeTextFunction_Set() error {
	var err error
	markupEscapeTextFunction_Once.Do(func() {
		markupEscapeTextFunction, err = gi.FunctionInvokerNew("GLib", "markup_escape_text")
	})
	return err
}

// MarkupEscapeText is a representation of the C type g_markup_escape_text.
func MarkupEscapeText(text string, length int32) (string, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(text)
	inArgs[1].SetInt32(length)

	var ret gi.Argument

	err := markupEscapeTextFunction_Set()
	if err == nil {
		ret = markupEscapeTextFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo, err
}

// UNSUPPORTED : C value 'g_markup_printf_escaped' : parameter '...' has no type

// UNSUPPORTED : C value 'g_markup_vprintf_escaped' : parameter 'args' of type 'va_list' not supported

var memIsSystemMallocFunction *gi.Function
var memIsSystemMallocFunction_Once sync.Once

func memIsSystemMallocFunction_Set() error {
	var err error
	memIsSystemMallocFunction_Once.Do(func() {
		memIsSystemMallocFunction, err = gi.FunctionInvokerNew("GLib", "mem_is_system_malloc")
	})
	return err
}

// MemIsSystemMalloc is a representation of the C type g_mem_is_system_malloc.
func MemIsSystemMalloc() (bool, error) {

	var ret gi.Argument

	err := memIsSystemMallocFunction_Set()
	if err == nil {
		ret = memIsSystemMallocFunction.Invoke(nil, nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

var memProfileFunction *gi.Function
var memProfileFunction_Once sync.Once

func memProfileFunction_Set() error {
	var err error
	memProfileFunction_Once.Do(func() {
		memProfileFunction, err = gi.FunctionInvokerNew("GLib", "mem_profile")
	})
	return err
}

// MemProfile is a representation of the C type g_mem_profile.
func MemProfile() error {

	err := memProfileFunction_Set()
	if err == nil {
		memProfileFunction.Invoke(nil, nil)
	}

	return err
}

// UNSUPPORTED : C value 'g_mem_set_vtable' : parameter 'vtable' of type 'MemVTable' not supported

// UNSUPPORTED : C value 'g_memdup' : parameter 'mem' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_mkdir_with_parents' : parameter 'pathname' of type 'filename' not supported

// UNSUPPORTED : C value 'g_mkdtemp' : parameter 'tmpl' of type 'filename' not supported

// UNSUPPORTED : C value 'g_mkdtemp_full' : parameter 'tmpl' of type 'filename' not supported

// UNSUPPORTED : C value 'g_mkstemp' : parameter 'tmpl' of type 'filename' not supported

// UNSUPPORTED : C value 'g_mkstemp_full' : parameter 'tmpl' of type 'filename' not supported

// UNSUPPORTED : C value 'g_nullify_pointer' : parameter 'nullify_location' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_number_parser_error_quark' : return type 'Quark' not supported

var onErrorQueryFunction *gi.Function
var onErrorQueryFunction_Once sync.Once

func onErrorQueryFunction_Set() error {
	var err error
	onErrorQueryFunction_Once.Do(func() {
		onErrorQueryFunction, err = gi.FunctionInvokerNew("GLib", "on_error_query")
	})
	return err
}

// OnErrorQuery is a representation of the C type g_on_error_query.
func OnErrorQuery(prgName string) error {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(prgName)

	err := onErrorQueryFunction_Set()
	if err == nil {
		onErrorQueryFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var onErrorStackTraceFunction *gi.Function
var onErrorStackTraceFunction_Once sync.Once

func onErrorStackTraceFunction_Set() error {
	var err error
	onErrorStackTraceFunction_Once.Do(func() {
		onErrorStackTraceFunction, err = gi.FunctionInvokerNew("GLib", "on_error_stack_trace")
	})
	return err
}

// OnErrorStackTrace is a representation of the C type g_on_error_stack_trace.
func OnErrorStackTrace(prgName string) error {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(prgName)

	err := onErrorStackTraceFunction_Set()
	if err == nil {
		onErrorStackTraceFunction.Invoke(inArgs[:], nil)
	}

	return err
}

// UNSUPPORTED : C value 'g_once_init_enter' : parameter 'location' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_once_init_leave' : parameter 'location' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_option_error_quark' : return type 'Quark' not supported

// UNSUPPORTED : C value 'g_parse_debug_string' : parameter 'keys' has no type

// UNSUPPORTED : C value 'g_path_get_basename' : parameter 'file_name' of type 'filename' not supported

// UNSUPPORTED : C value 'g_path_get_dirname' : parameter 'file_name' of type 'filename' not supported

// UNSUPPORTED : C value 'g_path_is_absolute' : parameter 'file_name' of type 'filename' not supported

// UNSUPPORTED : C value 'g_path_skip_root' : parameter 'file_name' of type 'filename' not supported

// UNSUPPORTED : C value 'g_pattern_match' : parameter 'pspec' of type 'PatternSpec' not supported

var patternMatchSimpleFunction *gi.Function
var patternMatchSimpleFunction_Once sync.Once

func patternMatchSimpleFunction_Set() error {
	var err error
	patternMatchSimpleFunction_Once.Do(func() {
		patternMatchSimpleFunction, err = gi.FunctionInvokerNew("GLib", "pattern_match_simple")
	})
	return err
}

// PatternMatchSimple is a representation of the C type g_pattern_match_simple.
func PatternMatchSimple(pattern string, string_ string) (bool, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(pattern)
	inArgs[1].SetString(string_)

	var ret gi.Argument

	err := patternMatchSimpleFunction_Set()
	if err == nil {
		ret = patternMatchSimpleFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

// UNSUPPORTED : C value 'g_pattern_match_string' : parameter 'pspec' of type 'PatternSpec' not supported

// UNSUPPORTED : C value 'g_pointer_bit_lock' : parameter 'address' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_pointer_bit_trylock' : parameter 'address' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_pointer_bit_unlock' : parameter 'address' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_poll' : parameter 'fds' of type 'PollFD' not supported

// UNSUPPORTED : C value 'g_prefix_error' : parameter 'err' of type 'Error' not supported

// UNSUPPORTED : C value 'g_print' : parameter '...' has no type

// UNSUPPORTED : C value 'g_printerr' : parameter '...' has no type

// UNSUPPORTED : C value 'g_printf' : parameter '...' has no type

// UNSUPPORTED : C value 'g_printf_string_upper_bound' : parameter 'args' of type 'va_list' not supported

// UNSUPPORTED : C value 'g_propagate_error' : parameter 'dest' of type 'Error' not supported

// UNSUPPORTED : C value 'g_propagate_prefixed_error' : parameter 'dest' of type 'Error' not supported

// UNSUPPORTED : C value 'g_ptr_array_find' : parameter 'haystack' has no type

// UNSUPPORTED : C value 'g_ptr_array_find_with_equal_func' : parameter 'haystack' has no type

// UNSUPPORTED : C value 'g_qsort_with_data' : parameter 'pbase' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_quark_from_static_string' : return type 'Quark' not supported

// UNSUPPORTED : C value 'g_quark_from_string' : return type 'Quark' not supported

var quarkToStringFunction *gi.Function
var quarkToStringFunction_Once sync.Once

func quarkToStringFunction_Set() error {
	var err error
	quarkToStringFunction_Once.Do(func() {
		quarkToStringFunction, err = gi.FunctionInvokerNew("GLib", "quark_to_string")
	})
	return err
}

// QuarkToString is a representation of the C type g_quark_to_string.
func QuarkToString(quark Quark) (string, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetUint32(uint32(quark))

	var ret gi.Argument

	err := quarkToStringFunction_Set()
	if err == nil {
		ret = quarkToStringFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo, err
}

// UNSUPPORTED : C value 'g_quark_try_string' : return type 'Quark' not supported

var randomDoubleFunction *gi.Function
var randomDoubleFunction_Once sync.Once

func randomDoubleFunction_Set() error {
	var err error
	randomDoubleFunction_Once.Do(func() {
		randomDoubleFunction, err = gi.FunctionInvokerNew("GLib", "random_double")
	})
	return err
}

// RandomDouble is a representation of the C type g_random_double.
func RandomDouble() (float64, error) {

	var ret gi.Argument

	err := randomDoubleFunction_Set()
	if err == nil {
		ret = randomDoubleFunction.Invoke(nil, nil)
	}

	retGo := ret.Float64()

	return retGo, err
}

var randomDoubleRangeFunction *gi.Function
var randomDoubleRangeFunction_Once sync.Once

func randomDoubleRangeFunction_Set() error {
	var err error
	randomDoubleRangeFunction_Once.Do(func() {
		randomDoubleRangeFunction, err = gi.FunctionInvokerNew("GLib", "random_double_range")
	})
	return err
}

// RandomDoubleRange is a representation of the C type g_random_double_range.
func RandomDoubleRange(begin float64, end float64) (float64, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetFloat64(begin)
	inArgs[1].SetFloat64(end)

	var ret gi.Argument

	err := randomDoubleRangeFunction_Set()
	if err == nil {
		ret = randomDoubleRangeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Float64()

	return retGo, err
}

var randomIntFunction *gi.Function
var randomIntFunction_Once sync.Once

func randomIntFunction_Set() error {
	var err error
	randomIntFunction_Once.Do(func() {
		randomIntFunction, err = gi.FunctionInvokerNew("GLib", "random_int")
	})
	return err
}

// RandomInt is a representation of the C type g_random_int.
func RandomInt() (uint32, error) {

	var ret gi.Argument

	err := randomIntFunction_Set()
	if err == nil {
		ret = randomIntFunction.Invoke(nil, nil)
	}

	retGo := ret.Uint32()

	return retGo, err
}

var randomIntRangeFunction *gi.Function
var randomIntRangeFunction_Once sync.Once

func randomIntRangeFunction_Set() error {
	var err error
	randomIntRangeFunction_Once.Do(func() {
		randomIntRangeFunction, err = gi.FunctionInvokerNew("GLib", "random_int_range")
	})
	return err
}

// RandomIntRange is a representation of the C type g_random_int_range.
func RandomIntRange(begin int32, end int32) (int32, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetInt32(begin)
	inArgs[1].SetInt32(end)

	var ret gi.Argument

	err := randomIntRangeFunction_Set()
	if err == nil {
		ret = randomIntRangeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo, err
}

var randomSetSeedFunction *gi.Function
var randomSetSeedFunction_Once sync.Once

func randomSetSeedFunction_Set() error {
	var err error
	randomSetSeedFunction_Once.Do(func() {
		randomSetSeedFunction, err = gi.FunctionInvokerNew("GLib", "random_set_seed")
	})
	return err
}

// RandomSetSeed is a representation of the C type g_random_set_seed.
func RandomSetSeed(seed uint32) error {
	var inArgs [1]gi.Argument
	inArgs[0].SetUint32(seed)

	err := randomSetSeedFunction_Set()
	if err == nil {
		randomSetSeedFunction.Invoke(inArgs[:], nil)
	}

	return err
}

// UNSUPPORTED : C value 'g_rc_box_acquire' : parameter 'mem_block' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_rc_box_alloc' : parameter 'block_size' of type 'gsize' not supported

// UNSUPPORTED : C value 'g_rc_box_alloc0' : parameter 'block_size' of type 'gsize' not supported

// UNSUPPORTED : C value 'g_rc_box_dup' : parameter 'block_size' of type 'gsize' not supported

// UNSUPPORTED : C value 'g_rc_box_get_size' : parameter 'mem_block' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_rc_box_release' : parameter 'mem_block' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_rc_box_release_full' : parameter 'mem_block' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_realloc' : parameter 'mem' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_realloc_n' : parameter 'mem' of type 'gpointer' not supported

var refCountCompareFunction *gi.Function
var refCountCompareFunction_Once sync.Once

func refCountCompareFunction_Set() error {
	var err error
	refCountCompareFunction_Once.Do(func() {
		refCountCompareFunction, err = gi.FunctionInvokerNew("GLib", "ref_count_compare")
	})
	return err
}

// RefCountCompare is a representation of the C type g_ref_count_compare.
func RefCountCompare(rc int32, val int32) (bool, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetInt32(rc)
	inArgs[1].SetInt32(val)

	var ret gi.Argument

	err := refCountCompareFunction_Set()
	if err == nil {
		ret = refCountCompareFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

var refCountDecFunction *gi.Function
var refCountDecFunction_Once sync.Once

func refCountDecFunction_Set() error {
	var err error
	refCountDecFunction_Once.Do(func() {
		refCountDecFunction, err = gi.FunctionInvokerNew("GLib", "ref_count_dec")
	})
	return err
}

// RefCountDec is a representation of the C type g_ref_count_dec.
func RefCountDec(rc int32) (bool, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetInt32(rc)

	var ret gi.Argument

	err := refCountDecFunction_Set()
	if err == nil {
		ret = refCountDecFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

var refCountIncFunction *gi.Function
var refCountIncFunction_Once sync.Once

func refCountIncFunction_Set() error {
	var err error
	refCountIncFunction_Once.Do(func() {
		refCountIncFunction, err = gi.FunctionInvokerNew("GLib", "ref_count_inc")
	})
	return err
}

// RefCountInc is a representation of the C type g_ref_count_inc.
func RefCountInc(rc int32) error {
	var inArgs [1]gi.Argument
	inArgs[0].SetInt32(rc)

	err := refCountIncFunction_Set()
	if err == nil {
		refCountIncFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var refCountInitFunction *gi.Function
var refCountInitFunction_Once sync.Once

func refCountInitFunction_Set() error {
	var err error
	refCountInitFunction_Once.Do(func() {
		refCountInitFunction, err = gi.FunctionInvokerNew("GLib", "ref_count_init")
	})
	return err
}

// RefCountInit is a representation of the C type g_ref_count_init.
func RefCountInit(rc int32) error {
	var inArgs [1]gi.Argument
	inArgs[0].SetInt32(rc)

	err := refCountInitFunction_Set()
	if err == nil {
		refCountInitFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var refStringAcquireFunction *gi.Function
var refStringAcquireFunction_Once sync.Once

func refStringAcquireFunction_Set() error {
	var err error
	refStringAcquireFunction_Once.Do(func() {
		refStringAcquireFunction, err = gi.FunctionInvokerNew("GLib", "ref_string_acquire")
	})
	return err
}

// RefStringAcquire is a representation of the C type g_ref_string_acquire.
func RefStringAcquire(str string) (string, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(str)

	var ret gi.Argument

	err := refStringAcquireFunction_Set()
	if err == nil {
		ret = refStringAcquireFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo, err
}

// UNSUPPORTED : C value 'g_ref_string_length' : return type 'gsize' not supported

var refStringNewFunction *gi.Function
var refStringNewFunction_Once sync.Once

func refStringNewFunction_Set() error {
	var err error
	refStringNewFunction_Once.Do(func() {
		refStringNewFunction, err = gi.FunctionInvokerNew("GLib", "ref_string_new")
	})
	return err
}

// RefStringNew is a representation of the C type g_ref_string_new.
func RefStringNew(str string) (string, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(str)

	var ret gi.Argument

	err := refStringNewFunction_Set()
	if err == nil {
		ret = refStringNewFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo, err
}

var refStringNewInternFunction *gi.Function
var refStringNewInternFunction_Once sync.Once

func refStringNewInternFunction_Set() error {
	var err error
	refStringNewInternFunction_Once.Do(func() {
		refStringNewInternFunction, err = gi.FunctionInvokerNew("GLib", "ref_string_new_intern")
	})
	return err
}

// RefStringNewIntern is a representation of the C type g_ref_string_new_intern.
func RefStringNewIntern(str string) (string, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(str)

	var ret gi.Argument

	err := refStringNewInternFunction_Set()
	if err == nil {
		ret = refStringNewInternFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo, err
}

var refStringNewLenFunction *gi.Function
var refStringNewLenFunction_Once sync.Once

func refStringNewLenFunction_Set() error {
	var err error
	refStringNewLenFunction_Once.Do(func() {
		refStringNewLenFunction, err = gi.FunctionInvokerNew("GLib", "ref_string_new_len")
	})
	return err
}

// RefStringNewLen is a representation of the C type g_ref_string_new_len.
func RefStringNewLen(str string, len int32) (string, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(str)
	inArgs[1].SetInt32(len)

	var ret gi.Argument

	err := refStringNewLenFunction_Set()
	if err == nil {
		ret = refStringNewLenFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo, err
}

var refStringReleaseFunction *gi.Function
var refStringReleaseFunction_Once sync.Once

func refStringReleaseFunction_Set() error {
	var err error
	refStringReleaseFunction_Once.Do(func() {
		refStringReleaseFunction, err = gi.FunctionInvokerNew("GLib", "ref_string_release")
	})
	return err
}

// RefStringRelease is a representation of the C type g_ref_string_release.
func RefStringRelease(str string) error {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(str)

	err := refStringReleaseFunction_Set()
	if err == nil {
		refStringReleaseFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var regexCheckReplacementFunction *gi.Function
var regexCheckReplacementFunction_Once sync.Once

func regexCheckReplacementFunction_Set() error {
	var err error
	regexCheckReplacementFunction_Once.Do(func() {
		regexCheckReplacementFunction, err = gi.FunctionInvokerNew("GLib", "regex_check_replacement")
	})
	return err
}

// RegexCheckReplacement is a representation of the C type g_regex_check_replacement.
func RegexCheckReplacement(replacement string) (bool, bool, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(replacement)

	var outArgs [1]gi.Argument
	var ret gi.Argument

	err := regexCheckReplacementFunction_Set()
	if err == nil {
		ret = regexCheckReplacementFunction.Invoke(inArgs[:], outArgs[:])
	}

	retGo := ret.Boolean()
	out0 := outArgs[0].Boolean()

	return retGo, out0, err
}

// UNSUPPORTED : C value 'g_regex_error_quark' : return type 'Quark' not supported

var regexEscapeNulFunction *gi.Function
var regexEscapeNulFunction_Once sync.Once

func regexEscapeNulFunction_Set() error {
	var err error
	regexEscapeNulFunction_Once.Do(func() {
		regexEscapeNulFunction, err = gi.FunctionInvokerNew("GLib", "regex_escape_nul")
	})
	return err
}

// RegexEscapeNul is a representation of the C type g_regex_escape_nul.
func RegexEscapeNul(string_ string, length int32) (string, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(string_)
	inArgs[1].SetInt32(length)

	var ret gi.Argument

	err := regexEscapeNulFunction_Set()
	if err == nil {
		ret = regexEscapeNulFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo, err
}

// UNSUPPORTED : C value 'g_regex_escape_string' : parameter 'string' has no type

// UNSUPPORTED : C value 'g_regex_match_simple' : parameter 'compile_options' of type 'RegexCompileFlags' not supported

// UNSUPPORTED : C value 'g_regex_split_simple' : parameter 'compile_options' of type 'RegexCompileFlags' not supported

var reloadUserSpecialDirsCacheFunction *gi.Function
var reloadUserSpecialDirsCacheFunction_Once sync.Once

func reloadUserSpecialDirsCacheFunction_Set() error {
	var err error
	reloadUserSpecialDirsCacheFunction_Once.Do(func() {
		reloadUserSpecialDirsCacheFunction, err = gi.FunctionInvokerNew("GLib", "reload_user_special_dirs_cache")
	})
	return err
}

// ReloadUserSpecialDirsCache is a representation of the C type g_reload_user_special_dirs_cache.
func ReloadUserSpecialDirsCache() error {

	err := reloadUserSpecialDirsCacheFunction_Set()
	if err == nil {
		reloadUserSpecialDirsCacheFunction.Invoke(nil, nil)
	}

	return err
}

var returnIfFailWarningFunction *gi.Function
var returnIfFailWarningFunction_Once sync.Once

func returnIfFailWarningFunction_Set() error {
	var err error
	returnIfFailWarningFunction_Once.Do(func() {
		returnIfFailWarningFunction, err = gi.FunctionInvokerNew("GLib", "return_if_fail_warning")
	})
	return err
}

// ReturnIfFailWarning is a representation of the C type g_return_if_fail_warning.
func ReturnIfFailWarning(logDomain string, prettyFunction string, expression string) error {
	var inArgs [3]gi.Argument
	inArgs[0].SetString(logDomain)
	inArgs[1].SetString(prettyFunction)
	inArgs[2].SetString(expression)

	err := returnIfFailWarningFunction_Set()
	if err == nil {
		returnIfFailWarningFunction.Invoke(inArgs[:], nil)
	}

	return err
}

// UNSUPPORTED : C value 'g_rmdir' : parameter 'filename' of type 'filename' not supported

// UNSUPPORTED : C value 'g_sequence_get' : parameter 'iter' of type 'SequenceIter' not supported

// UNSUPPORTED : C value 'g_sequence_insert_before' : parameter 'iter' of type 'SequenceIter' not supported

// UNSUPPORTED : C value 'g_sequence_move' : parameter 'src' of type 'SequenceIter' not supported

// UNSUPPORTED : C value 'g_sequence_move_range' : parameter 'dest' of type 'SequenceIter' not supported

// UNSUPPORTED : C value 'g_sequence_range_get_midpoint' : parameter 'begin' of type 'SequenceIter' not supported

// UNSUPPORTED : C value 'g_sequence_remove' : parameter 'iter' of type 'SequenceIter' not supported

// UNSUPPORTED : C value 'g_sequence_remove_range' : parameter 'begin' of type 'SequenceIter' not supported

// UNSUPPORTED : C value 'g_sequence_set' : parameter 'iter' of type 'SequenceIter' not supported

// UNSUPPORTED : C value 'g_sequence_swap' : parameter 'a' of type 'SequenceIter' not supported

var setApplicationNameFunction *gi.Function
var setApplicationNameFunction_Once sync.Once

func setApplicationNameFunction_Set() error {
	var err error
	setApplicationNameFunction_Once.Do(func() {
		setApplicationNameFunction, err = gi.FunctionInvokerNew("GLib", "set_application_name")
	})
	return err
}

// SetApplicationName is a representation of the C type g_set_application_name.
func SetApplicationName(applicationName string) error {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(applicationName)

	err := setApplicationNameFunction_Set()
	if err == nil {
		setApplicationNameFunction.Invoke(inArgs[:], nil)
	}

	return err
}

// UNSUPPORTED : C value 'g_set_error' : parameter 'err' of type 'Error' not supported

// UNSUPPORTED : C value 'g_set_error_literal' : parameter 'err' of type 'Error' not supported

var setPrgnameFunction *gi.Function
var setPrgnameFunction_Once sync.Once

func setPrgnameFunction_Set() error {
	var err error
	setPrgnameFunction_Once.Do(func() {
		setPrgnameFunction, err = gi.FunctionInvokerNew("GLib", "set_prgname")
	})
	return err
}

// SetPrgname is a representation of the C type g_set_prgname.
func SetPrgname(prgname string) error {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(prgname)

	err := setPrgnameFunction_Set()
	if err == nil {
		setPrgnameFunction.Invoke(inArgs[:], nil)
	}

	return err
}

// UNSUPPORTED : C value 'g_set_print_handler' : parameter 'func' of type 'PrintFunc' not supported

// UNSUPPORTED : C value 'g_set_printerr_handler' : parameter 'func' of type 'PrintFunc' not supported

// UNSUPPORTED : C value 'g_setenv' : parameter 'variable' of type 'filename' not supported

// UNSUPPORTED : C value 'g_shell_error_quark' : return type 'Quark' not supported

// UNSUPPORTED : C value 'g_shell_parse_argv' : parameter 'command_line' of type 'filename' not supported

// UNSUPPORTED : C value 'g_shell_quote' : parameter 'unquoted_string' of type 'filename' not supported

// UNSUPPORTED : C value 'g_shell_unquote' : parameter 'quoted_string' of type 'filename' not supported

// UNSUPPORTED : C value 'g_slice_alloc' : parameter 'block_size' of type 'gsize' not supported

// UNSUPPORTED : C value 'g_slice_alloc0' : parameter 'block_size' of type 'gsize' not supported

// UNSUPPORTED : C value 'g_slice_copy' : parameter 'block_size' of type 'gsize' not supported

// UNSUPPORTED : C value 'g_slice_free1' : parameter 'block_size' of type 'gsize' not supported

// UNSUPPORTED : C value 'g_slice_free_chain_with_offset' : parameter 'block_size' of type 'gsize' not supported

// UNSUPPORTED : C value 'g_slice_get_config' : parameter 'ckey' of type 'SliceConfig' not supported

// UNSUPPORTED : C value 'g_slice_get_config_state' : parameter 'ckey' of type 'SliceConfig' not supported

// UNSUPPORTED : C value 'g_slice_set_config' : parameter 'ckey' of type 'SliceConfig' not supported

// UNSUPPORTED : C value 'g_snprintf' : parameter '...' has no type

var sourceRemoveFunction *gi.Function
var sourceRemoveFunction_Once sync.Once

func sourceRemoveFunction_Set() error {
	var err error
	sourceRemoveFunction_Once.Do(func() {
		sourceRemoveFunction, err = gi.FunctionInvokerNew("GLib", "source_remove")
	})
	return err
}

// SourceRemove is a representation of the C type g_source_remove.
func SourceRemove(tag uint32) (bool, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetUint32(tag)

	var ret gi.Argument

	err := sourceRemoveFunction_Set()
	if err == nil {
		ret = sourceRemoveFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

// UNSUPPORTED : C value 'g_source_remove_by_funcs_user_data' : parameter 'funcs' of type 'SourceFuncs' not supported

// UNSUPPORTED : C value 'g_source_remove_by_user_data' : parameter 'user_data' of type 'gpointer' not supported

var sourceSetNameByIdFunction *gi.Function
var sourceSetNameByIdFunction_Once sync.Once

func sourceSetNameByIdFunction_Set() error {
	var err error
	sourceSetNameByIdFunction_Once.Do(func() {
		sourceSetNameByIdFunction, err = gi.FunctionInvokerNew("GLib", "source_set_name_by_id")
	})
	return err
}

// SourceSetNameById is a representation of the C type g_source_set_name_by_id.
func SourceSetNameById(tag uint32, name string) error {
	var inArgs [2]gi.Argument
	inArgs[0].SetUint32(tag)
	inArgs[1].SetString(name)

	err := sourceSetNameByIdFunction_Set()
	if err == nil {
		sourceSetNameByIdFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var spacedPrimesClosestFunction *gi.Function
var spacedPrimesClosestFunction_Once sync.Once

func spacedPrimesClosestFunction_Set() error {
	var err error
	spacedPrimesClosestFunction_Once.Do(func() {
		spacedPrimesClosestFunction, err = gi.FunctionInvokerNew("GLib", "spaced_primes_closest")
	})
	return err
}

// SpacedPrimesClosest is a representation of the C type g_spaced_primes_closest.
func SpacedPrimesClosest(num uint32) (uint32, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetUint32(num)

	var ret gi.Argument

	err := spacedPrimesClosestFunction_Set()
	if err == nil {
		ret = spacedPrimesClosestFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo, err
}

// UNSUPPORTED : C value 'g_spawn_async' : parameter 'working_directory' of type 'filename' not supported

// UNSUPPORTED : C value 'g_spawn_async_with_fds' : parameter 'working_directory' of type 'filename' not supported

// UNSUPPORTED : C value 'g_spawn_async_with_pipes' : parameter 'working_directory' of type 'filename' not supported

var spawnCheckExitStatusFunction *gi.Function
var spawnCheckExitStatusFunction_Once sync.Once

func spawnCheckExitStatusFunction_Set() error {
	var err error
	spawnCheckExitStatusFunction_Once.Do(func() {
		spawnCheckExitStatusFunction, err = gi.FunctionInvokerNew("GLib", "spawn_check_exit_status")
	})
	return err
}

// SpawnCheckExitStatus is a representation of the C type g_spawn_check_exit_status.
func SpawnCheckExitStatus(exitStatus int32) (bool, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetInt32(exitStatus)

	var ret gi.Argument

	err := spawnCheckExitStatusFunction_Set()
	if err == nil {
		ret = spawnCheckExitStatusFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

var spawnClosePidFunction *gi.Function
var spawnClosePidFunction_Once sync.Once

func spawnClosePidFunction_Set() error {
	var err error
	spawnClosePidFunction_Once.Do(func() {
		spawnClosePidFunction, err = gi.FunctionInvokerNew("GLib", "spawn_close_pid")
	})
	return err
}

// SpawnClosePid is a representation of the C type g_spawn_close_pid.
func SpawnClosePid(pid Pid) error {
	var inArgs [1]gi.Argument
	inArgs[0].SetInt32(int32(pid))

	err := spawnClosePidFunction_Set()
	if err == nil {
		spawnClosePidFunction.Invoke(inArgs[:], nil)
	}

	return err
}

// UNSUPPORTED : C value 'g_spawn_command_line_async' : parameter 'command_line' of type 'filename' not supported

// UNSUPPORTED : C value 'g_spawn_command_line_sync' : parameter 'command_line' of type 'filename' not supported

// UNSUPPORTED : C value 'g_spawn_error_quark' : return type 'Quark' not supported

// UNSUPPORTED : C value 'g_spawn_exit_error_quark' : return type 'Quark' not supported

// UNSUPPORTED : C value 'g_spawn_sync' : parameter 'working_directory' of type 'filename' not supported

// UNSUPPORTED : C value 'g_sprintf' : parameter '...' has no type

var stpcpyFunction *gi.Function
var stpcpyFunction_Once sync.Once

func stpcpyFunction_Set() error {
	var err error
	stpcpyFunction_Once.Do(func() {
		stpcpyFunction, err = gi.FunctionInvokerNew("GLib", "stpcpy")
	})
	return err
}

// Stpcpy is a representation of the C type g_stpcpy.
func Stpcpy(dest string, src string) (string, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(dest)
	inArgs[1].SetString(src)

	var ret gi.Argument

	err := stpcpyFunction_Set()
	if err == nil {
		ret = stpcpyFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo, err
}

// UNSUPPORTED : C value 'g_str_equal' : parameter 'v1' of type 'gpointer' not supported

var strHasPrefixFunction *gi.Function
var strHasPrefixFunction_Once sync.Once

func strHasPrefixFunction_Set() error {
	var err error
	strHasPrefixFunction_Once.Do(func() {
		strHasPrefixFunction, err = gi.FunctionInvokerNew("GLib", "str_has_prefix")
	})
	return err
}

// StrHasPrefix is a representation of the C type g_str_has_prefix.
func StrHasPrefix(str string, prefix string) (bool, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(str)
	inArgs[1].SetString(prefix)

	var ret gi.Argument

	err := strHasPrefixFunction_Set()
	if err == nil {
		ret = strHasPrefixFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

var strHasSuffixFunction *gi.Function
var strHasSuffixFunction_Once sync.Once

func strHasSuffixFunction_Set() error {
	var err error
	strHasSuffixFunction_Once.Do(func() {
		strHasSuffixFunction, err = gi.FunctionInvokerNew("GLib", "str_has_suffix")
	})
	return err
}

// StrHasSuffix is a representation of the C type g_str_has_suffix.
func StrHasSuffix(str string, suffix string) (bool, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(str)
	inArgs[1].SetString(suffix)

	var ret gi.Argument

	err := strHasSuffixFunction_Set()
	if err == nil {
		ret = strHasSuffixFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

// UNSUPPORTED : C value 'g_str_hash' : parameter 'v' of type 'gpointer' not supported

var strIsAsciiFunction *gi.Function
var strIsAsciiFunction_Once sync.Once

func strIsAsciiFunction_Set() error {
	var err error
	strIsAsciiFunction_Once.Do(func() {
		strIsAsciiFunction, err = gi.FunctionInvokerNew("GLib", "str_is_ascii")
	})
	return err
}

// StrIsAscii is a representation of the C type g_str_is_ascii.
func StrIsAscii(str string) (bool, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(str)

	var ret gi.Argument

	err := strIsAsciiFunction_Set()
	if err == nil {
		ret = strIsAsciiFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

var strMatchStringFunction *gi.Function
var strMatchStringFunction_Once sync.Once

func strMatchStringFunction_Set() error {
	var err error
	strMatchStringFunction_Once.Do(func() {
		strMatchStringFunction, err = gi.FunctionInvokerNew("GLib", "str_match_string")
	})
	return err
}

// StrMatchString is a representation of the C type g_str_match_string.
func StrMatchString(searchTerm string, potentialHit string, acceptAlternates bool) (bool, error) {
	var inArgs [3]gi.Argument
	inArgs[0].SetString(searchTerm)
	inArgs[1].SetString(potentialHit)
	inArgs[2].SetBoolean(acceptAlternates)

	var ret gi.Argument

	err := strMatchStringFunction_Set()
	if err == nil {
		ret = strMatchStringFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

var strToAsciiFunction *gi.Function
var strToAsciiFunction_Once sync.Once

func strToAsciiFunction_Set() error {
	var err error
	strToAsciiFunction_Once.Do(func() {
		strToAsciiFunction, err = gi.FunctionInvokerNew("GLib", "str_to_ascii")
	})
	return err
}

// StrToAscii is a representation of the C type g_str_to_ascii.
func StrToAscii(str string, fromLocale string) (string, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(str)
	inArgs[1].SetString(fromLocale)

	var ret gi.Argument

	err := strToAsciiFunction_Set()
	if err == nil {
		ret = strToAsciiFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo, err
}

// UNSUPPORTED : C value 'g_str_tokenize_and_fold' : parameter 'ascii_alternates' has no type

var strcanonFunction *gi.Function
var strcanonFunction_Once sync.Once

func strcanonFunction_Set() error {
	var err error
	strcanonFunction_Once.Do(func() {
		strcanonFunction, err = gi.FunctionInvokerNew("GLib", "strcanon")
	})
	return err
}

// Strcanon is a representation of the C type g_strcanon.
func Strcanon(string_ string, validChars string, substitutor int8) (string, error) {
	var inArgs [3]gi.Argument
	inArgs[0].SetString(string_)
	inArgs[1].SetString(validChars)
	inArgs[2].SetInt8(substitutor)

	var ret gi.Argument

	err := strcanonFunction_Set()
	if err == nil {
		ret = strcanonFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo, err
}

var strcasecmpFunction *gi.Function
var strcasecmpFunction_Once sync.Once

func strcasecmpFunction_Set() error {
	var err error
	strcasecmpFunction_Once.Do(func() {
		strcasecmpFunction, err = gi.FunctionInvokerNew("GLib", "strcasecmp")
	})
	return err
}

// Strcasecmp is a representation of the C type g_strcasecmp.
func Strcasecmp(s1 string, s2 string) (int32, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(s1)
	inArgs[1].SetString(s2)

	var ret gi.Argument

	err := strcasecmpFunction_Set()
	if err == nil {
		ret = strcasecmpFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo, err
}

var strchompFunction *gi.Function
var strchompFunction_Once sync.Once

func strchompFunction_Set() error {
	var err error
	strchompFunction_Once.Do(func() {
		strchompFunction, err = gi.FunctionInvokerNew("GLib", "strchomp")
	})
	return err
}

// Strchomp is a representation of the C type g_strchomp.
func Strchomp(string_ string) (string, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(string_)

	var ret gi.Argument

	err := strchompFunction_Set()
	if err == nil {
		ret = strchompFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo, err
}

var strchugFunction *gi.Function
var strchugFunction_Once sync.Once

func strchugFunction_Set() error {
	var err error
	strchugFunction_Once.Do(func() {
		strchugFunction, err = gi.FunctionInvokerNew("GLib", "strchug")
	})
	return err
}

// Strchug is a representation of the C type g_strchug.
func Strchug(string_ string) (string, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(string_)

	var ret gi.Argument

	err := strchugFunction_Set()
	if err == nil {
		ret = strchugFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo, err
}

var strcmp0Function *gi.Function
var strcmp0Function_Once sync.Once

func strcmp0Function_Set() error {
	var err error
	strcmp0Function_Once.Do(func() {
		strcmp0Function, err = gi.FunctionInvokerNew("GLib", "strcmp0")
	})
	return err
}

// Strcmp0 is a representation of the C type g_strcmp0.
func Strcmp0(str1 string, str2 string) (int32, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(str1)
	inArgs[1].SetString(str2)

	var ret gi.Argument

	err := strcmp0Function_Set()
	if err == nil {
		ret = strcmp0Function.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo, err
}

var strcompressFunction *gi.Function
var strcompressFunction_Once sync.Once

func strcompressFunction_Set() error {
	var err error
	strcompressFunction_Once.Do(func() {
		strcompressFunction, err = gi.FunctionInvokerNew("GLib", "strcompress")
	})
	return err
}

// Strcompress is a representation of the C type g_strcompress.
func Strcompress(source string) (string, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(source)

	var ret gi.Argument

	err := strcompressFunction_Set()
	if err == nil {
		ret = strcompressFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo, err
}

// UNSUPPORTED : C value 'g_strconcat' : parameter '...' has no type

var strdelimitFunction *gi.Function
var strdelimitFunction_Once sync.Once

func strdelimitFunction_Set() error {
	var err error
	strdelimitFunction_Once.Do(func() {
		strdelimitFunction, err = gi.FunctionInvokerNew("GLib", "strdelimit")
	})
	return err
}

// Strdelimit is a representation of the C type g_strdelimit.
func Strdelimit(string_ string, delimiters string, newDelimiter int8) (string, error) {
	var inArgs [3]gi.Argument
	inArgs[0].SetString(string_)
	inArgs[1].SetString(delimiters)
	inArgs[2].SetInt8(newDelimiter)

	var ret gi.Argument

	err := strdelimitFunction_Set()
	if err == nil {
		ret = strdelimitFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo, err
}

var strdownFunction *gi.Function
var strdownFunction_Once sync.Once

func strdownFunction_Set() error {
	var err error
	strdownFunction_Once.Do(func() {
		strdownFunction, err = gi.FunctionInvokerNew("GLib", "strdown")
	})
	return err
}

// Strdown is a representation of the C type g_strdown.
func Strdown(string_ string) (string, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(string_)

	var ret gi.Argument

	err := strdownFunction_Set()
	if err == nil {
		ret = strdownFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo, err
}

var strdupFunction *gi.Function
var strdupFunction_Once sync.Once

func strdupFunction_Set() error {
	var err error
	strdupFunction_Once.Do(func() {
		strdupFunction, err = gi.FunctionInvokerNew("GLib", "strdup")
	})
	return err
}

// Strdup is a representation of the C type g_strdup.
func Strdup(str string) (string, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(str)

	var ret gi.Argument

	err := strdupFunction_Set()
	if err == nil {
		ret = strdupFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo, err
}

// UNSUPPORTED : C value 'g_strdup_printf' : parameter '...' has no type

// UNSUPPORTED : C value 'g_strdup_vprintf' : parameter 'args' of type 'va_list' not supported

var strdupvFunction *gi.Function
var strdupvFunction_Once sync.Once

func strdupvFunction_Set() error {
	var err error
	strdupvFunction_Once.Do(func() {
		strdupvFunction, err = gi.FunctionInvokerNew("GLib", "strdupv")
	})
	return err
}

// Strdupv is a representation of the C type g_strdupv.
func Strdupv(strArray string) error {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(strArray)

	err := strdupvFunction_Set()
	if err == nil {
		strdupvFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var strerrorFunction *gi.Function
var strerrorFunction_Once sync.Once

func strerrorFunction_Set() error {
	var err error
	strerrorFunction_Once.Do(func() {
		strerrorFunction, err = gi.FunctionInvokerNew("GLib", "strerror")
	})
	return err
}

// Strerror is a representation of the C type g_strerror.
func Strerror(errnum int32) (string, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetInt32(errnum)

	var ret gi.Argument

	err := strerrorFunction_Set()
	if err == nil {
		ret = strerrorFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo, err
}

var strescapeFunction *gi.Function
var strescapeFunction_Once sync.Once

func strescapeFunction_Set() error {
	var err error
	strescapeFunction_Once.Do(func() {
		strescapeFunction, err = gi.FunctionInvokerNew("GLib", "strescape")
	})
	return err
}

// Strescape is a representation of the C type g_strescape.
func Strescape(source string, exceptions string) (string, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(source)
	inArgs[1].SetString(exceptions)

	var ret gi.Argument

	err := strescapeFunction_Set()
	if err == nil {
		ret = strescapeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo, err
}

var strfreevFunction *gi.Function
var strfreevFunction_Once sync.Once

func strfreevFunction_Set() error {
	var err error
	strfreevFunction_Once.Do(func() {
		strfreevFunction, err = gi.FunctionInvokerNew("GLib", "strfreev")
	})
	return err
}

// Strfreev is a representation of the C type g_strfreev.
func Strfreev(strArray string) error {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(strArray)

	err := strfreevFunction_Set()
	if err == nil {
		strfreevFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var stringNewFunction *gi.Function
var stringNewFunction_Once sync.Once

func stringNewFunction_Set() error {
	var err error
	stringNewFunction_Once.Do(func() {
		stringNewFunction, err = gi.FunctionInvokerNew("GLib", "string_new")
	})
	return err
}

// StringNew is a representation of the C type g_string_new.
func StringNew(init string) (*String, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(init)

	var ret gi.Argument

	err := stringNewFunction_Set()
	if err == nil {
		ret = stringNewFunction.Invoke(inArgs[:], nil)
	}

	retGo := &String{native: ret.Pointer()}

	return retGo, err
}

var stringNewLenFunction *gi.Function
var stringNewLenFunction_Once sync.Once

func stringNewLenFunction_Set() error {
	var err error
	stringNewLenFunction_Once.Do(func() {
		stringNewLenFunction, err = gi.FunctionInvokerNew("GLib", "string_new_len")
	})
	return err
}

// StringNewLen is a representation of the C type g_string_new_len.
func StringNewLen(init string, len int32) (*String, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(init)
	inArgs[1].SetInt32(len)

	var ret gi.Argument

	err := stringNewLenFunction_Set()
	if err == nil {
		ret = stringNewLenFunction.Invoke(inArgs[:], nil)
	}

	retGo := &String{native: ret.Pointer()}

	return retGo, err
}

// UNSUPPORTED : C value 'g_string_sized_new' : parameter 'dfl_size' of type 'gsize' not supported

var stripContextFunction *gi.Function
var stripContextFunction_Once sync.Once

func stripContextFunction_Set() error {
	var err error
	stripContextFunction_Once.Do(func() {
		stripContextFunction, err = gi.FunctionInvokerNew("GLib", "strip_context")
	})
	return err
}

// StripContext is a representation of the C type g_strip_context.
func StripContext(msgid string, msgval string) (string, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(msgid)
	inArgs[1].SetString(msgval)

	var ret gi.Argument

	err := stripContextFunction_Set()
	if err == nil {
		ret = stripContextFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo, err
}

// UNSUPPORTED : C value 'g_strjoin' : parameter '...' has no type

var strjoinvFunction *gi.Function
var strjoinvFunction_Once sync.Once

func strjoinvFunction_Set() error {
	var err error
	strjoinvFunction_Once.Do(func() {
		strjoinvFunction, err = gi.FunctionInvokerNew("GLib", "strjoinv")
	})
	return err
}

// Strjoinv is a representation of the C type g_strjoinv.
func Strjoinv(separator string, strArray string) (string, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(separator)
	inArgs[1].SetString(strArray)

	var ret gi.Argument

	err := strjoinvFunction_Set()
	if err == nil {
		ret = strjoinvFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo, err
}

// UNSUPPORTED : C value 'g_strlcat' : parameter 'dest_size' of type 'gsize' not supported

// UNSUPPORTED : C value 'g_strlcpy' : parameter 'dest_size' of type 'gsize' not supported

var strncasecmpFunction *gi.Function
var strncasecmpFunction_Once sync.Once

func strncasecmpFunction_Set() error {
	var err error
	strncasecmpFunction_Once.Do(func() {
		strncasecmpFunction, err = gi.FunctionInvokerNew("GLib", "strncasecmp")
	})
	return err
}

// Strncasecmp is a representation of the C type g_strncasecmp.
func Strncasecmp(s1 string, s2 string, n uint32) (int32, error) {
	var inArgs [3]gi.Argument
	inArgs[0].SetString(s1)
	inArgs[1].SetString(s2)
	inArgs[2].SetUint32(n)

	var ret gi.Argument

	err := strncasecmpFunction_Set()
	if err == nil {
		ret = strncasecmpFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo, err
}

// UNSUPPORTED : C value 'g_strndup' : parameter 'n' of type 'gsize' not supported

// UNSUPPORTED : C value 'g_strnfill' : parameter 'length' of type 'gsize' not supported

var strreverseFunction *gi.Function
var strreverseFunction_Once sync.Once

func strreverseFunction_Set() error {
	var err error
	strreverseFunction_Once.Do(func() {
		strreverseFunction, err = gi.FunctionInvokerNew("GLib", "strreverse")
	})
	return err
}

// Strreverse is a representation of the C type g_strreverse.
func Strreverse(string_ string) (string, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(string_)

	var ret gi.Argument

	err := strreverseFunction_Set()
	if err == nil {
		ret = strreverseFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo, err
}

var strrstrFunction *gi.Function
var strrstrFunction_Once sync.Once

func strrstrFunction_Set() error {
	var err error
	strrstrFunction_Once.Do(func() {
		strrstrFunction, err = gi.FunctionInvokerNew("GLib", "strrstr")
	})
	return err
}

// Strrstr is a representation of the C type g_strrstr.
func Strrstr(haystack string, needle string) (string, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(haystack)
	inArgs[1].SetString(needle)

	var ret gi.Argument

	err := strrstrFunction_Set()
	if err == nil {
		ret = strrstrFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo, err
}

var strrstrLenFunction *gi.Function
var strrstrLenFunction_Once sync.Once

func strrstrLenFunction_Set() error {
	var err error
	strrstrLenFunction_Once.Do(func() {
		strrstrLenFunction, err = gi.FunctionInvokerNew("GLib", "strrstr_len")
	})
	return err
}

// StrrstrLen is a representation of the C type g_strrstr_len.
func StrrstrLen(haystack string, haystackLen int32, needle string) (string, error) {
	var inArgs [3]gi.Argument
	inArgs[0].SetString(haystack)
	inArgs[1].SetInt32(haystackLen)
	inArgs[2].SetString(needle)

	var ret gi.Argument

	err := strrstrLenFunction_Set()
	if err == nil {
		ret = strrstrLenFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo, err
}

var strsignalFunction *gi.Function
var strsignalFunction_Once sync.Once

func strsignalFunction_Set() error {
	var err error
	strsignalFunction_Once.Do(func() {
		strsignalFunction, err = gi.FunctionInvokerNew("GLib", "strsignal")
	})
	return err
}

// Strsignal is a representation of the C type g_strsignal.
func Strsignal(signum int32) (string, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetInt32(signum)

	var ret gi.Argument

	err := strsignalFunction_Set()
	if err == nil {
		ret = strsignalFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo, err
}

var strsplitFunction *gi.Function
var strsplitFunction_Once sync.Once

func strsplitFunction_Set() error {
	var err error
	strsplitFunction_Once.Do(func() {
		strsplitFunction, err = gi.FunctionInvokerNew("GLib", "strsplit")
	})
	return err
}

// Strsplit is a representation of the C type g_strsplit.
func Strsplit(string_ string, delimiter string, maxTokens int32) error {
	var inArgs [3]gi.Argument
	inArgs[0].SetString(string_)
	inArgs[1].SetString(delimiter)
	inArgs[2].SetInt32(maxTokens)

	err := strsplitFunction_Set()
	if err == nil {
		strsplitFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var strsplitSetFunction *gi.Function
var strsplitSetFunction_Once sync.Once

func strsplitSetFunction_Set() error {
	var err error
	strsplitSetFunction_Once.Do(func() {
		strsplitSetFunction, err = gi.FunctionInvokerNew("GLib", "strsplit_set")
	})
	return err
}

// StrsplitSet is a representation of the C type g_strsplit_set.
func StrsplitSet(string_ string, delimiters string, maxTokens int32) error {
	var inArgs [3]gi.Argument
	inArgs[0].SetString(string_)
	inArgs[1].SetString(delimiters)
	inArgs[2].SetInt32(maxTokens)

	err := strsplitSetFunction_Set()
	if err == nil {
		strsplitSetFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var strstrLenFunction *gi.Function
var strstrLenFunction_Once sync.Once

func strstrLenFunction_Set() error {
	var err error
	strstrLenFunction_Once.Do(func() {
		strstrLenFunction, err = gi.FunctionInvokerNew("GLib", "strstr_len")
	})
	return err
}

// StrstrLen is a representation of the C type g_strstr_len.
func StrstrLen(haystack string, haystackLen int32, needle string) (string, error) {
	var inArgs [3]gi.Argument
	inArgs[0].SetString(haystack)
	inArgs[1].SetInt32(haystackLen)
	inArgs[2].SetString(needle)

	var ret gi.Argument

	err := strstrLenFunction_Set()
	if err == nil {
		ret = strstrLenFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo, err
}

var strtodFunction *gi.Function
var strtodFunction_Once sync.Once

func strtodFunction_Set() error {
	var err error
	strtodFunction_Once.Do(func() {
		strtodFunction, err = gi.FunctionInvokerNew("GLib", "strtod")
	})
	return err
}

// Strtod is a representation of the C type g_strtod.
func Strtod(nptr string) (float64, string, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(nptr)

	var outArgs [1]gi.Argument
	var ret gi.Argument

	err := strtodFunction_Set()
	if err == nil {
		ret = strtodFunction.Invoke(inArgs[:], outArgs[:])
	}

	retGo := ret.Float64()
	out0 := outArgs[0].String(false)

	return retGo, out0, err
}

var strupFunction *gi.Function
var strupFunction_Once sync.Once

func strupFunction_Set() error {
	var err error
	strupFunction_Once.Do(func() {
		strupFunction, err = gi.FunctionInvokerNew("GLib", "strup")
	})
	return err
}

// Strup is a representation of the C type g_strup.
func Strup(string_ string) (string, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(string_)

	var ret gi.Argument

	err := strupFunction_Set()
	if err == nil {
		ret = strupFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo, err
}

var strvContainsFunction *gi.Function
var strvContainsFunction_Once sync.Once

func strvContainsFunction_Set() error {
	var err error
	strvContainsFunction_Once.Do(func() {
		strvContainsFunction, err = gi.FunctionInvokerNew("GLib", "strv_contains")
	})
	return err
}

// StrvContains is a representation of the C type g_strv_contains.
func StrvContains(strv string, str string) (bool, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(strv)
	inArgs[1].SetString(str)

	var ret gi.Argument

	err := strvContainsFunction_Set()
	if err == nil {
		ret = strvContainsFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

var strvEqualFunction *gi.Function
var strvEqualFunction_Once sync.Once

func strvEqualFunction_Set() error {
	var err error
	strvEqualFunction_Once.Do(func() {
		strvEqualFunction, err = gi.FunctionInvokerNew("GLib", "strv_equal")
	})
	return err
}

// StrvEqual is a representation of the C type g_strv_equal.
func StrvEqual(strv1 string, strv2 string) (bool, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(strv1)
	inArgs[1].SetString(strv2)

	var ret gi.Argument

	err := strvEqualFunction_Set()
	if err == nil {
		ret = strvEqualFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

// UNSUPPORTED : C value 'g_strv_get_type' : return type 'GType' not supported

var strvLengthFunction *gi.Function
var strvLengthFunction_Once sync.Once

func strvLengthFunction_Set() error {
	var err error
	strvLengthFunction_Once.Do(func() {
		strvLengthFunction, err = gi.FunctionInvokerNew("GLib", "strv_length")
	})
	return err
}

// StrvLength is a representation of the C type g_strv_length.
func StrvLength(strArray string) (uint32, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(strArray)

	var ret gi.Argument

	err := strvLengthFunction_Set()
	if err == nil {
		ret = strvLengthFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo, err
}

// UNSUPPORTED : C value 'g_test_add_data_func' : parameter 'test_data' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_test_add_data_func_full' : parameter 'test_data' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_test_add_func' : parameter 'test_func' of type 'TestFunc' not supported

// UNSUPPORTED : C value 'g_test_add_vtable' : parameter 'data_size' of type 'gsize' not supported

var testAssertExpectedMessagesInternalFunction *gi.Function
var testAssertExpectedMessagesInternalFunction_Once sync.Once

func testAssertExpectedMessagesInternalFunction_Set() error {
	var err error
	testAssertExpectedMessagesInternalFunction_Once.Do(func() {
		testAssertExpectedMessagesInternalFunction, err = gi.FunctionInvokerNew("GLib", "test_assert_expected_messages_internal")
	})
	return err
}

// TestAssertExpectedMessagesInternal is a representation of the C type g_test_assert_expected_messages_internal.
func TestAssertExpectedMessagesInternal(domain string, file string, line int32, func_ string) error {
	var inArgs [4]gi.Argument
	inArgs[0].SetString(domain)
	inArgs[1].SetString(file)
	inArgs[2].SetInt32(line)
	inArgs[3].SetString(func_)

	err := testAssertExpectedMessagesInternalFunction_Set()
	if err == nil {
		testAssertExpectedMessagesInternalFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var testBugFunction *gi.Function
var testBugFunction_Once sync.Once

func testBugFunction_Set() error {
	var err error
	testBugFunction_Once.Do(func() {
		testBugFunction, err = gi.FunctionInvokerNew("GLib", "test_bug")
	})
	return err
}

// TestBug is a representation of the C type g_test_bug.
func TestBug(bugUriSnippet string) error {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(bugUriSnippet)

	err := testBugFunction_Set()
	if err == nil {
		testBugFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var testBugBaseFunction *gi.Function
var testBugBaseFunction_Once sync.Once

func testBugBaseFunction_Set() error {
	var err error
	testBugBaseFunction_Once.Do(func() {
		testBugBaseFunction, err = gi.FunctionInvokerNew("GLib", "test_bug_base")
	})
	return err
}

// TestBugBase is a representation of the C type g_test_bug_base.
func TestBugBase(uriPattern string) error {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(uriPattern)

	err := testBugBaseFunction_Set()
	if err == nil {
		testBugBaseFunction.Invoke(inArgs[:], nil)
	}

	return err
}

// UNSUPPORTED : C value 'g_test_build_filename' : parameter 'file_type' of type 'TestFileType' not supported

// UNSUPPORTED : C value 'g_test_create_case' : parameter 'data_size' of type 'gsize' not supported

var testCreateSuiteFunction *gi.Function
var testCreateSuiteFunction_Once sync.Once

func testCreateSuiteFunction_Set() error {
	var err error
	testCreateSuiteFunction_Once.Do(func() {
		testCreateSuiteFunction, err = gi.FunctionInvokerNew("GLib", "test_create_suite")
	})
	return err
}

// TestCreateSuite is a representation of the C type g_test_create_suite.
func TestCreateSuite(suiteName string) (*TestSuite, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(suiteName)

	var ret gi.Argument

	err := testCreateSuiteFunction_Set()
	if err == nil {
		ret = testCreateSuiteFunction.Invoke(inArgs[:], nil)
	}

	retGo := &TestSuite{native: ret.Pointer()}

	return retGo, err
}

// UNSUPPORTED : C value 'g_test_expect_message' : parameter 'log_level' of type 'LogLevelFlags' not supported

var testFailFunction *gi.Function
var testFailFunction_Once sync.Once

func testFailFunction_Set() error {
	var err error
	testFailFunction_Once.Do(func() {
		testFailFunction, err = gi.FunctionInvokerNew("GLib", "test_fail")
	})
	return err
}

// TestFail is a representation of the C type g_test_fail.
func TestFail() error {

	err := testFailFunction_Set()
	if err == nil {
		testFailFunction.Invoke(nil, nil)
	}

	return err
}

var testFailedFunction *gi.Function
var testFailedFunction_Once sync.Once

func testFailedFunction_Set() error {
	var err error
	testFailedFunction_Once.Do(func() {
		testFailedFunction, err = gi.FunctionInvokerNew("GLib", "test_failed")
	})
	return err
}

// TestFailed is a representation of the C type g_test_failed.
func TestFailed() (bool, error) {

	var ret gi.Argument

	err := testFailedFunction_Set()
	if err == nil {
		ret = testFailedFunction.Invoke(nil, nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

// UNSUPPORTED : C value 'g_test_get_dir' : parameter 'file_type' of type 'TestFileType' not supported

// UNSUPPORTED : C value 'g_test_get_filename' : parameter 'file_type' of type 'TestFileType' not supported

var testGetRootFunction *gi.Function
var testGetRootFunction_Once sync.Once

func testGetRootFunction_Set() error {
	var err error
	testGetRootFunction_Once.Do(func() {
		testGetRootFunction, err = gi.FunctionInvokerNew("GLib", "test_get_root")
	})
	return err
}

// TestGetRoot is a representation of the C type g_test_get_root.
func TestGetRoot() (*TestSuite, error) {

	var ret gi.Argument

	err := testGetRootFunction_Set()
	if err == nil {
		ret = testGetRootFunction.Invoke(nil, nil)
	}

	retGo := &TestSuite{native: ret.Pointer()}

	return retGo, err
}

var testIncompleteFunction *gi.Function
var testIncompleteFunction_Once sync.Once

func testIncompleteFunction_Set() error {
	var err error
	testIncompleteFunction_Once.Do(func() {
		testIncompleteFunction, err = gi.FunctionInvokerNew("GLib", "test_incomplete")
	})
	return err
}

// TestIncomplete is a representation of the C type g_test_incomplete.
func TestIncomplete(msg string) error {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(msg)

	err := testIncompleteFunction_Set()
	if err == nil {
		testIncompleteFunction.Invoke(inArgs[:], nil)
	}

	return err
}

// UNSUPPORTED : C value 'g_test_init' : parameter '...' has no type

// UNSUPPORTED : C value 'g_test_log_set_fatal_handler' : parameter 'log_func' of type 'TestLogFatalFunc' not supported

// UNSUPPORTED : C value 'g_test_log_type_name' : parameter 'log_type' of type 'TestLogType' not supported

// UNSUPPORTED : C value 'g_test_maximized_result' : parameter '...' has no type

// UNSUPPORTED : C value 'g_test_message' : parameter '...' has no type

// UNSUPPORTED : C value 'g_test_minimized_result' : parameter '...' has no type

// UNSUPPORTED : C value 'g_test_queue_destroy' : parameter 'destroy_func' of type 'DestroyNotify' not supported

// UNSUPPORTED : C value 'g_test_queue_free' : parameter 'gfree_pointer' of type 'gpointer' not supported

var testRandDoubleFunction *gi.Function
var testRandDoubleFunction_Once sync.Once

func testRandDoubleFunction_Set() error {
	var err error
	testRandDoubleFunction_Once.Do(func() {
		testRandDoubleFunction, err = gi.FunctionInvokerNew("GLib", "test_rand_double")
	})
	return err
}

// TestRandDouble is a representation of the C type g_test_rand_double.
func TestRandDouble() (float64, error) {

	var ret gi.Argument

	err := testRandDoubleFunction_Set()
	if err == nil {
		ret = testRandDoubleFunction.Invoke(nil, nil)
	}

	retGo := ret.Float64()

	return retGo, err
}

var testRandDoubleRangeFunction *gi.Function
var testRandDoubleRangeFunction_Once sync.Once

func testRandDoubleRangeFunction_Set() error {
	var err error
	testRandDoubleRangeFunction_Once.Do(func() {
		testRandDoubleRangeFunction, err = gi.FunctionInvokerNew("GLib", "test_rand_double_range")
	})
	return err
}

// TestRandDoubleRange is a representation of the C type g_test_rand_double_range.
func TestRandDoubleRange(rangeStart float64, rangeEnd float64) (float64, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetFloat64(rangeStart)
	inArgs[1].SetFloat64(rangeEnd)

	var ret gi.Argument

	err := testRandDoubleRangeFunction_Set()
	if err == nil {
		ret = testRandDoubleRangeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Float64()

	return retGo, err
}

var testRandIntFunction *gi.Function
var testRandIntFunction_Once sync.Once

func testRandIntFunction_Set() error {
	var err error
	testRandIntFunction_Once.Do(func() {
		testRandIntFunction, err = gi.FunctionInvokerNew("GLib", "test_rand_int")
	})
	return err
}

// TestRandInt is a representation of the C type g_test_rand_int.
func TestRandInt() (int32, error) {

	var ret gi.Argument

	err := testRandIntFunction_Set()
	if err == nil {
		ret = testRandIntFunction.Invoke(nil, nil)
	}

	retGo := ret.Int32()

	return retGo, err
}

var testRandIntRangeFunction *gi.Function
var testRandIntRangeFunction_Once sync.Once

func testRandIntRangeFunction_Set() error {
	var err error
	testRandIntRangeFunction_Once.Do(func() {
		testRandIntRangeFunction, err = gi.FunctionInvokerNew("GLib", "test_rand_int_range")
	})
	return err
}

// TestRandIntRange is a representation of the C type g_test_rand_int_range.
func TestRandIntRange(begin int32, end int32) (int32, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetInt32(begin)
	inArgs[1].SetInt32(end)

	var ret gi.Argument

	err := testRandIntRangeFunction_Set()
	if err == nil {
		ret = testRandIntRangeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo, err
}

var testRunFunction *gi.Function
var testRunFunction_Once sync.Once

func testRunFunction_Set() error {
	var err error
	testRunFunction_Once.Do(func() {
		testRunFunction, err = gi.FunctionInvokerNew("GLib", "test_run")
	})
	return err
}

// TestRun is a representation of the C type g_test_run.
func TestRun() (int32, error) {

	var ret gi.Argument

	err := testRunFunction_Set()
	if err == nil {
		ret = testRunFunction.Invoke(nil, nil)
	}

	retGo := ret.Int32()

	return retGo, err
}

// UNSUPPORTED : C value 'g_test_run_suite' : parameter 'suite' of type 'TestSuite' not supported

var testSetNonfatalAssertionsFunction *gi.Function
var testSetNonfatalAssertionsFunction_Once sync.Once

func testSetNonfatalAssertionsFunction_Set() error {
	var err error
	testSetNonfatalAssertionsFunction_Once.Do(func() {
		testSetNonfatalAssertionsFunction, err = gi.FunctionInvokerNew("GLib", "test_set_nonfatal_assertions")
	})
	return err
}

// TestSetNonfatalAssertions is a representation of the C type g_test_set_nonfatal_assertions.
func TestSetNonfatalAssertions() error {

	err := testSetNonfatalAssertionsFunction_Set()
	if err == nil {
		testSetNonfatalAssertionsFunction.Invoke(nil, nil)
	}

	return err
}

var testSkipFunction *gi.Function
var testSkipFunction_Once sync.Once

func testSkipFunction_Set() error {
	var err error
	testSkipFunction_Once.Do(func() {
		testSkipFunction, err = gi.FunctionInvokerNew("GLib", "test_skip")
	})
	return err
}

// TestSkip is a representation of the C type g_test_skip.
func TestSkip(msg string) error {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(msg)

	err := testSkipFunction_Set()
	if err == nil {
		testSkipFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var testSubprocessFunction *gi.Function
var testSubprocessFunction_Once sync.Once

func testSubprocessFunction_Set() error {
	var err error
	testSubprocessFunction_Once.Do(func() {
		testSubprocessFunction, err = gi.FunctionInvokerNew("GLib", "test_subprocess")
	})
	return err
}

// TestSubprocess is a representation of the C type g_test_subprocess.
func TestSubprocess() (bool, error) {

	var ret gi.Argument

	err := testSubprocessFunction_Set()
	if err == nil {
		ret = testSubprocessFunction.Invoke(nil, nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

var testSummaryFunction *gi.Function
var testSummaryFunction_Once sync.Once

func testSummaryFunction_Set() error {
	var err error
	testSummaryFunction_Once.Do(func() {
		testSummaryFunction, err = gi.FunctionInvokerNew("GLib", "test_summary")
	})
	return err
}

// TestSummary is a representation of the C type g_test_summary.
func TestSummary(summary string) error {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(summary)

	err := testSummaryFunction_Set()
	if err == nil {
		testSummaryFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var testTimerElapsedFunction *gi.Function
var testTimerElapsedFunction_Once sync.Once

func testTimerElapsedFunction_Set() error {
	var err error
	testTimerElapsedFunction_Once.Do(func() {
		testTimerElapsedFunction, err = gi.FunctionInvokerNew("GLib", "test_timer_elapsed")
	})
	return err
}

// TestTimerElapsed is a representation of the C type g_test_timer_elapsed.
func TestTimerElapsed() (float64, error) {

	var ret gi.Argument

	err := testTimerElapsedFunction_Set()
	if err == nil {
		ret = testTimerElapsedFunction.Invoke(nil, nil)
	}

	retGo := ret.Float64()

	return retGo, err
}

var testTimerLastFunction *gi.Function
var testTimerLastFunction_Once sync.Once

func testTimerLastFunction_Set() error {
	var err error
	testTimerLastFunction_Once.Do(func() {
		testTimerLastFunction, err = gi.FunctionInvokerNew("GLib", "test_timer_last")
	})
	return err
}

// TestTimerLast is a representation of the C type g_test_timer_last.
func TestTimerLast() (float64, error) {

	var ret gi.Argument

	err := testTimerLastFunction_Set()
	if err == nil {
		ret = testTimerLastFunction.Invoke(nil, nil)
	}

	retGo := ret.Float64()

	return retGo, err
}

var testTimerStartFunction *gi.Function
var testTimerStartFunction_Once sync.Once

func testTimerStartFunction_Set() error {
	var err error
	testTimerStartFunction_Once.Do(func() {
		testTimerStartFunction, err = gi.FunctionInvokerNew("GLib", "test_timer_start")
	})
	return err
}

// TestTimerStart is a representation of the C type g_test_timer_start.
func TestTimerStart() error {

	err := testTimerStartFunction_Set()
	if err == nil {
		testTimerStartFunction.Invoke(nil, nil)
	}

	return err
}

var testTrapAssertionsFunction *gi.Function
var testTrapAssertionsFunction_Once sync.Once

func testTrapAssertionsFunction_Set() error {
	var err error
	testTrapAssertionsFunction_Once.Do(func() {
		testTrapAssertionsFunction, err = gi.FunctionInvokerNew("GLib", "test_trap_assertions")
	})
	return err
}

// TestTrapAssertions is a representation of the C type g_test_trap_assertions.
func TestTrapAssertions(domain string, file string, line int32, func_ string, assertionFlags uint64, pattern string) error {
	var inArgs [6]gi.Argument
	inArgs[0].SetString(domain)
	inArgs[1].SetString(file)
	inArgs[2].SetInt32(line)
	inArgs[3].SetString(func_)
	inArgs[4].SetUint64(assertionFlags)
	inArgs[5].SetString(pattern)

	err := testTrapAssertionsFunction_Set()
	if err == nil {
		testTrapAssertionsFunction.Invoke(inArgs[:], nil)
	}

	return err
}

// UNSUPPORTED : C value 'g_test_trap_fork' : parameter 'test_trap_flags' of type 'TestTrapFlags' not supported

var testTrapHasPassedFunction *gi.Function
var testTrapHasPassedFunction_Once sync.Once

func testTrapHasPassedFunction_Set() error {
	var err error
	testTrapHasPassedFunction_Once.Do(func() {
		testTrapHasPassedFunction, err = gi.FunctionInvokerNew("GLib", "test_trap_has_passed")
	})
	return err
}

// TestTrapHasPassed is a representation of the C type g_test_trap_has_passed.
func TestTrapHasPassed() (bool, error) {

	var ret gi.Argument

	err := testTrapHasPassedFunction_Set()
	if err == nil {
		ret = testTrapHasPassedFunction.Invoke(nil, nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

var testTrapReachedTimeoutFunction *gi.Function
var testTrapReachedTimeoutFunction_Once sync.Once

func testTrapReachedTimeoutFunction_Set() error {
	var err error
	testTrapReachedTimeoutFunction_Once.Do(func() {
		testTrapReachedTimeoutFunction, err = gi.FunctionInvokerNew("GLib", "test_trap_reached_timeout")
	})
	return err
}

// TestTrapReachedTimeout is a representation of the C type g_test_trap_reached_timeout.
func TestTrapReachedTimeout() (bool, error) {

	var ret gi.Argument

	err := testTrapReachedTimeoutFunction_Set()
	if err == nil {
		ret = testTrapReachedTimeoutFunction.Invoke(nil, nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

// UNSUPPORTED : C value 'g_test_trap_subprocess' : parameter 'test_flags' of type 'TestSubprocessFlags' not supported

// UNSUPPORTED : C value 'g_thread_error_quark' : return type 'Quark' not supported

// UNSUPPORTED : C value 'g_thread_exit' : parameter 'retval' of type 'gpointer' not supported

var threadPoolGetMaxIdleTimeFunction *gi.Function
var threadPoolGetMaxIdleTimeFunction_Once sync.Once

func threadPoolGetMaxIdleTimeFunction_Set() error {
	var err error
	threadPoolGetMaxIdleTimeFunction_Once.Do(func() {
		threadPoolGetMaxIdleTimeFunction, err = gi.FunctionInvokerNew("GLib", "thread_pool_get_max_idle_time")
	})
	return err
}

// ThreadPoolGetMaxIdleTime is a representation of the C type g_thread_pool_get_max_idle_time.
func ThreadPoolGetMaxIdleTime() (uint32, error) {

	var ret gi.Argument

	err := threadPoolGetMaxIdleTimeFunction_Set()
	if err == nil {
		ret = threadPoolGetMaxIdleTimeFunction.Invoke(nil, nil)
	}

	retGo := ret.Uint32()

	return retGo, err
}

var threadPoolGetMaxUnusedThreadsFunction *gi.Function
var threadPoolGetMaxUnusedThreadsFunction_Once sync.Once

func threadPoolGetMaxUnusedThreadsFunction_Set() error {
	var err error
	threadPoolGetMaxUnusedThreadsFunction_Once.Do(func() {
		threadPoolGetMaxUnusedThreadsFunction, err = gi.FunctionInvokerNew("GLib", "thread_pool_get_max_unused_threads")
	})
	return err
}

// ThreadPoolGetMaxUnusedThreads is a representation of the C type g_thread_pool_get_max_unused_threads.
func ThreadPoolGetMaxUnusedThreads() (int32, error) {

	var ret gi.Argument

	err := threadPoolGetMaxUnusedThreadsFunction_Set()
	if err == nil {
		ret = threadPoolGetMaxUnusedThreadsFunction.Invoke(nil, nil)
	}

	retGo := ret.Int32()

	return retGo, err
}

var threadPoolGetNumUnusedThreadsFunction *gi.Function
var threadPoolGetNumUnusedThreadsFunction_Once sync.Once

func threadPoolGetNumUnusedThreadsFunction_Set() error {
	var err error
	threadPoolGetNumUnusedThreadsFunction_Once.Do(func() {
		threadPoolGetNumUnusedThreadsFunction, err = gi.FunctionInvokerNew("GLib", "thread_pool_get_num_unused_threads")
	})
	return err
}

// ThreadPoolGetNumUnusedThreads is a representation of the C type g_thread_pool_get_num_unused_threads.
func ThreadPoolGetNumUnusedThreads() (uint32, error) {

	var ret gi.Argument

	err := threadPoolGetNumUnusedThreadsFunction_Set()
	if err == nil {
		ret = threadPoolGetNumUnusedThreadsFunction.Invoke(nil, nil)
	}

	retGo := ret.Uint32()

	return retGo, err
}

var threadPoolSetMaxIdleTimeFunction *gi.Function
var threadPoolSetMaxIdleTimeFunction_Once sync.Once

func threadPoolSetMaxIdleTimeFunction_Set() error {
	var err error
	threadPoolSetMaxIdleTimeFunction_Once.Do(func() {
		threadPoolSetMaxIdleTimeFunction, err = gi.FunctionInvokerNew("GLib", "thread_pool_set_max_idle_time")
	})
	return err
}

// ThreadPoolSetMaxIdleTime is a representation of the C type g_thread_pool_set_max_idle_time.
func ThreadPoolSetMaxIdleTime(interval uint32) error {
	var inArgs [1]gi.Argument
	inArgs[0].SetUint32(interval)

	err := threadPoolSetMaxIdleTimeFunction_Set()
	if err == nil {
		threadPoolSetMaxIdleTimeFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var threadPoolSetMaxUnusedThreadsFunction *gi.Function
var threadPoolSetMaxUnusedThreadsFunction_Once sync.Once

func threadPoolSetMaxUnusedThreadsFunction_Set() error {
	var err error
	threadPoolSetMaxUnusedThreadsFunction_Once.Do(func() {
		threadPoolSetMaxUnusedThreadsFunction, err = gi.FunctionInvokerNew("GLib", "thread_pool_set_max_unused_threads")
	})
	return err
}

// ThreadPoolSetMaxUnusedThreads is a representation of the C type g_thread_pool_set_max_unused_threads.
func ThreadPoolSetMaxUnusedThreads(maxThreads int32) error {
	var inArgs [1]gi.Argument
	inArgs[0].SetInt32(maxThreads)

	err := threadPoolSetMaxUnusedThreadsFunction_Set()
	if err == nil {
		threadPoolSetMaxUnusedThreadsFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var threadPoolStopUnusedThreadsFunction *gi.Function
var threadPoolStopUnusedThreadsFunction_Once sync.Once

func threadPoolStopUnusedThreadsFunction_Set() error {
	var err error
	threadPoolStopUnusedThreadsFunction_Once.Do(func() {
		threadPoolStopUnusedThreadsFunction, err = gi.FunctionInvokerNew("GLib", "thread_pool_stop_unused_threads")
	})
	return err
}

// ThreadPoolStopUnusedThreads is a representation of the C type g_thread_pool_stop_unused_threads.
func ThreadPoolStopUnusedThreads() error {

	err := threadPoolStopUnusedThreadsFunction_Set()
	if err == nil {
		threadPoolStopUnusedThreadsFunction.Invoke(nil, nil)
	}

	return err
}

var threadSelfFunction *gi.Function
var threadSelfFunction_Once sync.Once

func threadSelfFunction_Set() error {
	var err error
	threadSelfFunction_Once.Do(func() {
		threadSelfFunction, err = gi.FunctionInvokerNew("GLib", "thread_self")
	})
	return err
}

// ThreadSelf is a representation of the C type g_thread_self.
func ThreadSelf() (*Thread, error) {

	var ret gi.Argument

	err := threadSelfFunction_Set()
	if err == nil {
		ret = threadSelfFunction.Invoke(nil, nil)
	}

	retGo := &Thread{native: ret.Pointer()}

	return retGo, err
}

var threadYieldFunction *gi.Function
var threadYieldFunction_Once sync.Once

func threadYieldFunction_Set() error {
	var err error
	threadYieldFunction_Once.Do(func() {
		threadYieldFunction, err = gi.FunctionInvokerNew("GLib", "thread_yield")
	})
	return err
}

// ThreadYield is a representation of the C type g_thread_yield.
func ThreadYield() error {

	err := threadYieldFunction_Set()
	if err == nil {
		threadYieldFunction.Invoke(nil, nil)
	}

	return err
}

// UNSUPPORTED : C value 'g_time_val_from_iso8601' : parameter 'time_' of type 'TimeVal' not supported

// UNSUPPORTED : C value 'g_timeout_add' : parameter 'function' of type 'SourceFunc' not supported

// UNSUPPORTED : C value 'g_timeout_add_full' : parameter 'function' of type 'SourceFunc' not supported

// UNSUPPORTED : C value 'g_timeout_add_seconds' : parameter 'function' of type 'SourceFunc' not supported

// UNSUPPORTED : C value 'g_timeout_add_seconds_full' : parameter 'function' of type 'SourceFunc' not supported

var timeoutSourceNewFunction *gi.Function
var timeoutSourceNewFunction_Once sync.Once

func timeoutSourceNewFunction_Set() error {
	var err error
	timeoutSourceNewFunction_Once.Do(func() {
		timeoutSourceNewFunction, err = gi.FunctionInvokerNew("GLib", "timeout_source_new")
	})
	return err
}

// TimeoutSourceNew is a representation of the C type g_timeout_source_new.
func TimeoutSourceNew(interval uint32) (*Source, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetUint32(interval)

	var ret gi.Argument

	err := timeoutSourceNewFunction_Set()
	if err == nil {
		ret = timeoutSourceNewFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Source{native: ret.Pointer()}

	return retGo, err
}

var timeoutSourceNewSecondsFunction *gi.Function
var timeoutSourceNewSecondsFunction_Once sync.Once

func timeoutSourceNewSecondsFunction_Set() error {
	var err error
	timeoutSourceNewSecondsFunction_Once.Do(func() {
		timeoutSourceNewSecondsFunction, err = gi.FunctionInvokerNew("GLib", "timeout_source_new_seconds")
	})
	return err
}

// TimeoutSourceNewSeconds is a representation of the C type g_timeout_source_new_seconds.
func TimeoutSourceNewSeconds(interval uint32) (*Source, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetUint32(interval)

	var ret gi.Argument

	err := timeoutSourceNewSecondsFunction_Set()
	if err == nil {
		ret = timeoutSourceNewSecondsFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Source{native: ret.Pointer()}

	return retGo, err
}

// UNSUPPORTED : C value 'g_trash_stack_height' : parameter 'stack_p' of type 'TrashStack' not supported

// UNSUPPORTED : C value 'g_trash_stack_peek' : parameter 'stack_p' of type 'TrashStack' not supported

// UNSUPPORTED : C value 'g_trash_stack_pop' : parameter 'stack_p' of type 'TrashStack' not supported

// UNSUPPORTED : C value 'g_trash_stack_push' : parameter 'stack_p' of type 'TrashStack' not supported

// UNSUPPORTED : C value 'g_try_malloc' : parameter 'n_bytes' of type 'gsize' not supported

// UNSUPPORTED : C value 'g_try_malloc0' : parameter 'n_bytes' of type 'gsize' not supported

// UNSUPPORTED : C value 'g_try_malloc0_n' : parameter 'n_blocks' of type 'gsize' not supported

// UNSUPPORTED : C value 'g_try_malloc_n' : parameter 'n_blocks' of type 'gsize' not supported

// UNSUPPORTED : C value 'g_try_realloc' : parameter 'mem' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_try_realloc_n' : parameter 'mem' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_ucs4_to_utf16' : parameter 'str' of type 'gunichar' not supported

// UNSUPPORTED : C value 'g_ucs4_to_utf8' : parameter 'str' of type 'gunichar' not supported

// UNSUPPORTED : C value 'g_unichar_break_type' : parameter 'c' of type 'gunichar' not supported

// UNSUPPORTED : C value 'g_unichar_combining_class' : parameter 'uc' of type 'gunichar' not supported

// UNSUPPORTED : C value 'g_unichar_compose' : parameter 'a' of type 'gunichar' not supported

// UNSUPPORTED : C value 'g_unichar_decompose' : parameter 'ch' of type 'gunichar' not supported

// UNSUPPORTED : C value 'g_unichar_digit_value' : parameter 'c' of type 'gunichar' not supported

// UNSUPPORTED : C value 'g_unichar_fully_decompose' : parameter 'ch' of type 'gunichar' not supported

// UNSUPPORTED : C value 'g_unichar_get_mirror_char' : parameter 'ch' of type 'gunichar' not supported

// UNSUPPORTED : C value 'g_unichar_get_script' : parameter 'ch' of type 'gunichar' not supported

// UNSUPPORTED : C value 'g_unichar_isalnum' : parameter 'c' of type 'gunichar' not supported

// UNSUPPORTED : C value 'g_unichar_isalpha' : parameter 'c' of type 'gunichar' not supported

// UNSUPPORTED : C value 'g_unichar_iscntrl' : parameter 'c' of type 'gunichar' not supported

// UNSUPPORTED : C value 'g_unichar_isdefined' : parameter 'c' of type 'gunichar' not supported

// UNSUPPORTED : C value 'g_unichar_isdigit' : parameter 'c' of type 'gunichar' not supported

// UNSUPPORTED : C value 'g_unichar_isgraph' : parameter 'c' of type 'gunichar' not supported

// UNSUPPORTED : C value 'g_unichar_islower' : parameter 'c' of type 'gunichar' not supported

// UNSUPPORTED : C value 'g_unichar_ismark' : parameter 'c' of type 'gunichar' not supported

// UNSUPPORTED : C value 'g_unichar_isprint' : parameter 'c' of type 'gunichar' not supported

// UNSUPPORTED : C value 'g_unichar_ispunct' : parameter 'c' of type 'gunichar' not supported

// UNSUPPORTED : C value 'g_unichar_isspace' : parameter 'c' of type 'gunichar' not supported

// UNSUPPORTED : C value 'g_unichar_istitle' : parameter 'c' of type 'gunichar' not supported

// UNSUPPORTED : C value 'g_unichar_isupper' : parameter 'c' of type 'gunichar' not supported

// UNSUPPORTED : C value 'g_unichar_iswide' : parameter 'c' of type 'gunichar' not supported

// UNSUPPORTED : C value 'g_unichar_iswide_cjk' : parameter 'c' of type 'gunichar' not supported

// UNSUPPORTED : C value 'g_unichar_isxdigit' : parameter 'c' of type 'gunichar' not supported

// UNSUPPORTED : C value 'g_unichar_iszerowidth' : parameter 'c' of type 'gunichar' not supported

// UNSUPPORTED : C value 'g_unichar_to_utf8' : parameter 'c' of type 'gunichar' not supported

// UNSUPPORTED : C value 'g_unichar_tolower' : parameter 'c' of type 'gunichar' not supported

// UNSUPPORTED : C value 'g_unichar_totitle' : parameter 'c' of type 'gunichar' not supported

// UNSUPPORTED : C value 'g_unichar_toupper' : parameter 'c' of type 'gunichar' not supported

// UNSUPPORTED : C value 'g_unichar_type' : parameter 'c' of type 'gunichar' not supported

// UNSUPPORTED : C value 'g_unichar_validate' : parameter 'ch' of type 'gunichar' not supported

// UNSUPPORTED : C value 'g_unichar_xdigit_value' : parameter 'c' of type 'gunichar' not supported

// UNSUPPORTED : C value 'g_unicode_canonical_decomposition' : parameter 'ch' of type 'gunichar' not supported

// UNSUPPORTED : C value 'g_unicode_canonical_ordering' : parameter 'string' of type 'gunichar' not supported

// UNSUPPORTED : C value 'g_unicode_script_from_iso15924' : return type 'UnicodeScript' not supported

// UNSUPPORTED : C value 'g_unicode_script_to_iso15924' : parameter 'script' of type 'UnicodeScript' not supported

// UNSUPPORTED : C value 'g_unix_error_quark' : return type 'Quark' not supported

// UNSUPPORTED : C value 'g_unix_fd_add' : parameter 'condition' of type 'IOCondition' not supported

// UNSUPPORTED : C value 'g_unix_fd_add_full' : parameter 'condition' of type 'IOCondition' not supported

// UNSUPPORTED : C value 'g_unix_fd_source_new' : parameter 'condition' of type 'IOCondition' not supported

var unixOpenPipeFunction *gi.Function
var unixOpenPipeFunction_Once sync.Once

func unixOpenPipeFunction_Set() error {
	var err error
	unixOpenPipeFunction_Once.Do(func() {
		unixOpenPipeFunction, err = gi.FunctionInvokerNew("GLib", "unix_open_pipe")
	})
	return err
}

// UnixOpenPipe is a representation of the C type g_unix_open_pipe.
func UnixOpenPipe(fds int32, flags int32) (bool, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetInt32(fds)
	inArgs[1].SetInt32(flags)

	var ret gi.Argument

	err := unixOpenPipeFunction_Set()
	if err == nil {
		ret = unixOpenPipeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

var unixSetFdNonblockingFunction *gi.Function
var unixSetFdNonblockingFunction_Once sync.Once

func unixSetFdNonblockingFunction_Set() error {
	var err error
	unixSetFdNonblockingFunction_Once.Do(func() {
		unixSetFdNonblockingFunction, err = gi.FunctionInvokerNew("GLib", "unix_set_fd_nonblocking")
	})
	return err
}

// UnixSetFdNonblocking is a representation of the C type g_unix_set_fd_nonblocking.
func UnixSetFdNonblocking(fd int32, nonblock bool) (bool, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetInt32(fd)
	inArgs[1].SetBoolean(nonblock)

	var ret gi.Argument

	err := unixSetFdNonblockingFunction_Set()
	if err == nil {
		ret = unixSetFdNonblockingFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

// UNSUPPORTED : C value 'g_unix_signal_add' : parameter 'handler' of type 'SourceFunc' not supported

// UNSUPPORTED : C value 'g_unix_signal_add_full' : parameter 'handler' of type 'SourceFunc' not supported

var unixSignalSourceNewFunction *gi.Function
var unixSignalSourceNewFunction_Once sync.Once

func unixSignalSourceNewFunction_Set() error {
	var err error
	unixSignalSourceNewFunction_Once.Do(func() {
		unixSignalSourceNewFunction, err = gi.FunctionInvokerNew("GLib", "unix_signal_source_new")
	})
	return err
}

// UnixSignalSourceNew is a representation of the C type g_unix_signal_source_new.
func UnixSignalSourceNew(signum int32) (*Source, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetInt32(signum)

	var ret gi.Argument

	err := unixSignalSourceNewFunction_Set()
	if err == nil {
		ret = unixSignalSourceNewFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Source{native: ret.Pointer()}

	return retGo, err
}

// UNSUPPORTED : C value 'g_unlink' : parameter 'filename' of type 'filename' not supported

// UNSUPPORTED : C value 'g_unsetenv' : parameter 'variable' of type 'filename' not supported

var uriEscapeStringFunction *gi.Function
var uriEscapeStringFunction_Once sync.Once

func uriEscapeStringFunction_Set() error {
	var err error
	uriEscapeStringFunction_Once.Do(func() {
		uriEscapeStringFunction, err = gi.FunctionInvokerNew("GLib", "uri_escape_string")
	})
	return err
}

// UriEscapeString is a representation of the C type g_uri_escape_string.
func UriEscapeString(unescaped string, reservedCharsAllowed string, allowUtf8 bool) (string, error) {
	var inArgs [3]gi.Argument
	inArgs[0].SetString(unescaped)
	inArgs[1].SetString(reservedCharsAllowed)
	inArgs[2].SetBoolean(allowUtf8)

	var ret gi.Argument

	err := uriEscapeStringFunction_Set()
	if err == nil {
		ret = uriEscapeStringFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo, err
}

var uriListExtractUrisFunction *gi.Function
var uriListExtractUrisFunction_Once sync.Once

func uriListExtractUrisFunction_Set() error {
	var err error
	uriListExtractUrisFunction_Once.Do(func() {
		uriListExtractUrisFunction, err = gi.FunctionInvokerNew("GLib", "uri_list_extract_uris")
	})
	return err
}

// UriListExtractUris is a representation of the C type g_uri_list_extract_uris.
func UriListExtractUris(uriList string) error {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(uriList)

	err := uriListExtractUrisFunction_Set()
	if err == nil {
		uriListExtractUrisFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var uriParseSchemeFunction *gi.Function
var uriParseSchemeFunction_Once sync.Once

func uriParseSchemeFunction_Set() error {
	var err error
	uriParseSchemeFunction_Once.Do(func() {
		uriParseSchemeFunction, err = gi.FunctionInvokerNew("GLib", "uri_parse_scheme")
	})
	return err
}

// UriParseScheme is a representation of the C type g_uri_parse_scheme.
func UriParseScheme(uri string) (string, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(uri)

	var ret gi.Argument

	err := uriParseSchemeFunction_Set()
	if err == nil {
		ret = uriParseSchemeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo, err
}

var uriUnescapeSegmentFunction *gi.Function
var uriUnescapeSegmentFunction_Once sync.Once

func uriUnescapeSegmentFunction_Set() error {
	var err error
	uriUnescapeSegmentFunction_Once.Do(func() {
		uriUnescapeSegmentFunction, err = gi.FunctionInvokerNew("GLib", "uri_unescape_segment")
	})
	return err
}

// UriUnescapeSegment is a representation of the C type g_uri_unescape_segment.
func UriUnescapeSegment(escapedString string, escapedStringEnd string, illegalCharacters string) (string, error) {
	var inArgs [3]gi.Argument
	inArgs[0].SetString(escapedString)
	inArgs[1].SetString(escapedStringEnd)
	inArgs[2].SetString(illegalCharacters)

	var ret gi.Argument

	err := uriUnescapeSegmentFunction_Set()
	if err == nil {
		ret = uriUnescapeSegmentFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo, err
}

var uriUnescapeStringFunction *gi.Function
var uriUnescapeStringFunction_Once sync.Once

func uriUnescapeStringFunction_Set() error {
	var err error
	uriUnescapeStringFunction_Once.Do(func() {
		uriUnescapeStringFunction, err = gi.FunctionInvokerNew("GLib", "uri_unescape_string")
	})
	return err
}

// UriUnescapeString is a representation of the C type g_uri_unescape_string.
func UriUnescapeString(escapedString string, illegalCharacters string) (string, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(escapedString)
	inArgs[1].SetString(illegalCharacters)

	var ret gi.Argument

	err := uriUnescapeStringFunction_Set()
	if err == nil {
		ret = uriUnescapeStringFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo, err
}

var usleepFunction *gi.Function
var usleepFunction_Once sync.Once

func usleepFunction_Set() error {
	var err error
	usleepFunction_Once.Do(func() {
		usleepFunction, err = gi.FunctionInvokerNew("GLib", "usleep")
	})
	return err
}

// Usleep is a representation of the C type g_usleep.
func Usleep(microseconds uint64) error {
	var inArgs [1]gi.Argument
	inArgs[0].SetUint64(microseconds)

	err := usleepFunction_Set()
	if err == nil {
		usleepFunction.Invoke(inArgs[:], nil)
	}

	return err
}

// UNSUPPORTED : C value 'g_utf16_to_ucs4' : return type 'gunichar' not supported

var utf16ToUtf8Function *gi.Function
var utf16ToUtf8Function_Once sync.Once

func utf16ToUtf8Function_Set() error {
	var err error
	utf16ToUtf8Function_Once.Do(func() {
		utf16ToUtf8Function, err = gi.FunctionInvokerNew("GLib", "utf16_to_utf8")
	})
	return err
}

// Utf16ToUtf8 is a representation of the C type g_utf16_to_utf8.
func Utf16ToUtf8(str uint16, len int64) (string, int64, int64, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetUint16(str)
	inArgs[1].SetInt64(len)

	var outArgs [2]gi.Argument
	var ret gi.Argument

	err := utf16ToUtf8Function_Set()
	if err == nil {
		ret = utf16ToUtf8Function.Invoke(inArgs[:], outArgs[:])
	}

	retGo := ret.String(true)
	out0 := outArgs[0].Int64()
	out1 := outArgs[1].Int64()

	return retGo, out0, out1, err
}

var utf8CasefoldFunction *gi.Function
var utf8CasefoldFunction_Once sync.Once

func utf8CasefoldFunction_Set() error {
	var err error
	utf8CasefoldFunction_Once.Do(func() {
		utf8CasefoldFunction, err = gi.FunctionInvokerNew("GLib", "utf8_casefold")
	})
	return err
}

// Utf8Casefold is a representation of the C type g_utf8_casefold.
func Utf8Casefold(str string, len int32) (string, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(str)
	inArgs[1].SetInt32(len)

	var ret gi.Argument

	err := utf8CasefoldFunction_Set()
	if err == nil {
		ret = utf8CasefoldFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo, err
}

var utf8CollateFunction *gi.Function
var utf8CollateFunction_Once sync.Once

func utf8CollateFunction_Set() error {
	var err error
	utf8CollateFunction_Once.Do(func() {
		utf8CollateFunction, err = gi.FunctionInvokerNew("GLib", "utf8_collate")
	})
	return err
}

// Utf8Collate is a representation of the C type g_utf8_collate.
func Utf8Collate(str1 string, str2 string) (int32, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(str1)
	inArgs[1].SetString(str2)

	var ret gi.Argument

	err := utf8CollateFunction_Set()
	if err == nil {
		ret = utf8CollateFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo, err
}

var utf8CollateKeyFunction *gi.Function
var utf8CollateKeyFunction_Once sync.Once

func utf8CollateKeyFunction_Set() error {
	var err error
	utf8CollateKeyFunction_Once.Do(func() {
		utf8CollateKeyFunction, err = gi.FunctionInvokerNew("GLib", "utf8_collate_key")
	})
	return err
}

// Utf8CollateKey is a representation of the C type g_utf8_collate_key.
func Utf8CollateKey(str string, len int32) (string, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(str)
	inArgs[1].SetInt32(len)

	var ret gi.Argument

	err := utf8CollateKeyFunction_Set()
	if err == nil {
		ret = utf8CollateKeyFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo, err
}

var utf8CollateKeyForFilenameFunction *gi.Function
var utf8CollateKeyForFilenameFunction_Once sync.Once

func utf8CollateKeyForFilenameFunction_Set() error {
	var err error
	utf8CollateKeyForFilenameFunction_Once.Do(func() {
		utf8CollateKeyForFilenameFunction, err = gi.FunctionInvokerNew("GLib", "utf8_collate_key_for_filename")
	})
	return err
}

// Utf8CollateKeyForFilename is a representation of the C type g_utf8_collate_key_for_filename.
func Utf8CollateKeyForFilename(str string, len int32) (string, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(str)
	inArgs[1].SetInt32(len)

	var ret gi.Argument

	err := utf8CollateKeyForFilenameFunction_Set()
	if err == nil {
		ret = utf8CollateKeyForFilenameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo, err
}

var utf8FindNextCharFunction *gi.Function
var utf8FindNextCharFunction_Once sync.Once

func utf8FindNextCharFunction_Set() error {
	var err error
	utf8FindNextCharFunction_Once.Do(func() {
		utf8FindNextCharFunction, err = gi.FunctionInvokerNew("GLib", "utf8_find_next_char")
	})
	return err
}

// Utf8FindNextChar is a representation of the C type g_utf8_find_next_char.
func Utf8FindNextChar(p string, end string) (string, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(p)
	inArgs[1].SetString(end)

	var ret gi.Argument

	err := utf8FindNextCharFunction_Set()
	if err == nil {
		ret = utf8FindNextCharFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo, err
}

var utf8FindPrevCharFunction *gi.Function
var utf8FindPrevCharFunction_Once sync.Once

func utf8FindPrevCharFunction_Set() error {
	var err error
	utf8FindPrevCharFunction_Once.Do(func() {
		utf8FindPrevCharFunction, err = gi.FunctionInvokerNew("GLib", "utf8_find_prev_char")
	})
	return err
}

// Utf8FindPrevChar is a representation of the C type g_utf8_find_prev_char.
func Utf8FindPrevChar(str string, p string) (string, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(str)
	inArgs[1].SetString(p)

	var ret gi.Argument

	err := utf8FindPrevCharFunction_Set()
	if err == nil {
		ret = utf8FindPrevCharFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo, err
}

// UNSUPPORTED : C value 'g_utf8_get_char' : return type 'gunichar' not supported

// UNSUPPORTED : C value 'g_utf8_get_char_validated' : return type 'gunichar' not supported

var utf8MakeValidFunction *gi.Function
var utf8MakeValidFunction_Once sync.Once

func utf8MakeValidFunction_Set() error {
	var err error
	utf8MakeValidFunction_Once.Do(func() {
		utf8MakeValidFunction, err = gi.FunctionInvokerNew("GLib", "utf8_make_valid")
	})
	return err
}

// Utf8MakeValid is a representation of the C type g_utf8_make_valid.
func Utf8MakeValid(str string, len int32) (string, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(str)
	inArgs[1].SetInt32(len)

	var ret gi.Argument

	err := utf8MakeValidFunction_Set()
	if err == nil {
		ret = utf8MakeValidFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo, err
}

// UNSUPPORTED : C value 'g_utf8_normalize' : parameter 'mode' of type 'NormalizeMode' not supported

var utf8OffsetToPointerFunction *gi.Function
var utf8OffsetToPointerFunction_Once sync.Once

func utf8OffsetToPointerFunction_Set() error {
	var err error
	utf8OffsetToPointerFunction_Once.Do(func() {
		utf8OffsetToPointerFunction, err = gi.FunctionInvokerNew("GLib", "utf8_offset_to_pointer")
	})
	return err
}

// Utf8OffsetToPointer is a representation of the C type g_utf8_offset_to_pointer.
func Utf8OffsetToPointer(str string, offset int64) (string, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(str)
	inArgs[1].SetInt64(offset)

	var ret gi.Argument

	err := utf8OffsetToPointerFunction_Set()
	if err == nil {
		ret = utf8OffsetToPointerFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo, err
}

var utf8PointerToOffsetFunction *gi.Function
var utf8PointerToOffsetFunction_Once sync.Once

func utf8PointerToOffsetFunction_Set() error {
	var err error
	utf8PointerToOffsetFunction_Once.Do(func() {
		utf8PointerToOffsetFunction, err = gi.FunctionInvokerNew("GLib", "utf8_pointer_to_offset")
	})
	return err
}

// Utf8PointerToOffset is a representation of the C type g_utf8_pointer_to_offset.
func Utf8PointerToOffset(str string, pos string) (int64, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(str)
	inArgs[1].SetString(pos)

	var ret gi.Argument

	err := utf8PointerToOffsetFunction_Set()
	if err == nil {
		ret = utf8PointerToOffsetFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int64()

	return retGo, err
}

var utf8PrevCharFunction *gi.Function
var utf8PrevCharFunction_Once sync.Once

func utf8PrevCharFunction_Set() error {
	var err error
	utf8PrevCharFunction_Once.Do(func() {
		utf8PrevCharFunction, err = gi.FunctionInvokerNew("GLib", "utf8_prev_char")
	})
	return err
}

// Utf8PrevChar is a representation of the C type g_utf8_prev_char.
func Utf8PrevChar(p string) (string, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(p)

	var ret gi.Argument

	err := utf8PrevCharFunction_Set()
	if err == nil {
		ret = utf8PrevCharFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo, err
}

// UNSUPPORTED : C value 'g_utf8_strchr' : parameter 'c' of type 'gunichar' not supported

var utf8StrdownFunction *gi.Function
var utf8StrdownFunction_Once sync.Once

func utf8StrdownFunction_Set() error {
	var err error
	utf8StrdownFunction_Once.Do(func() {
		utf8StrdownFunction, err = gi.FunctionInvokerNew("GLib", "utf8_strdown")
	})
	return err
}

// Utf8Strdown is a representation of the C type g_utf8_strdown.
func Utf8Strdown(str string, len int32) (string, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(str)
	inArgs[1].SetInt32(len)

	var ret gi.Argument

	err := utf8StrdownFunction_Set()
	if err == nil {
		ret = utf8StrdownFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo, err
}

var utf8StrlenFunction *gi.Function
var utf8StrlenFunction_Once sync.Once

func utf8StrlenFunction_Set() error {
	var err error
	utf8StrlenFunction_Once.Do(func() {
		utf8StrlenFunction, err = gi.FunctionInvokerNew("GLib", "utf8_strlen")
	})
	return err
}

// Utf8Strlen is a representation of the C type g_utf8_strlen.
func Utf8Strlen(p string, max int32) (int64, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(p)
	inArgs[1].SetInt32(max)

	var ret gi.Argument

	err := utf8StrlenFunction_Set()
	if err == nil {
		ret = utf8StrlenFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int64()

	return retGo, err
}

// UNSUPPORTED : C value 'g_utf8_strncpy' : parameter 'n' of type 'gsize' not supported

// UNSUPPORTED : C value 'g_utf8_strrchr' : parameter 'c' of type 'gunichar' not supported

var utf8StrreverseFunction *gi.Function
var utf8StrreverseFunction_Once sync.Once

func utf8StrreverseFunction_Set() error {
	var err error
	utf8StrreverseFunction_Once.Do(func() {
		utf8StrreverseFunction, err = gi.FunctionInvokerNew("GLib", "utf8_strreverse")
	})
	return err
}

// Utf8Strreverse is a representation of the C type g_utf8_strreverse.
func Utf8Strreverse(str string, len int32) (string, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(str)
	inArgs[1].SetInt32(len)

	var ret gi.Argument

	err := utf8StrreverseFunction_Set()
	if err == nil {
		ret = utf8StrreverseFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo, err
}

var utf8StrupFunction *gi.Function
var utf8StrupFunction_Once sync.Once

func utf8StrupFunction_Set() error {
	var err error
	utf8StrupFunction_Once.Do(func() {
		utf8StrupFunction, err = gi.FunctionInvokerNew("GLib", "utf8_strup")
	})
	return err
}

// Utf8Strup is a representation of the C type g_utf8_strup.
func Utf8Strup(str string, len int32) (string, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(str)
	inArgs[1].SetInt32(len)

	var ret gi.Argument

	err := utf8StrupFunction_Set()
	if err == nil {
		ret = utf8StrupFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo, err
}

var utf8SubstringFunction *gi.Function
var utf8SubstringFunction_Once sync.Once

func utf8SubstringFunction_Set() error {
	var err error
	utf8SubstringFunction_Once.Do(func() {
		utf8SubstringFunction, err = gi.FunctionInvokerNew("GLib", "utf8_substring")
	})
	return err
}

// Utf8Substring is a representation of the C type g_utf8_substring.
func Utf8Substring(str string, startPos int64, endPos int64) (string, error) {
	var inArgs [3]gi.Argument
	inArgs[0].SetString(str)
	inArgs[1].SetInt64(startPos)
	inArgs[2].SetInt64(endPos)

	var ret gi.Argument

	err := utf8SubstringFunction_Set()
	if err == nil {
		ret = utf8SubstringFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo, err
}

// UNSUPPORTED : C value 'g_utf8_to_ucs4' : return type 'gunichar' not supported

// UNSUPPORTED : C value 'g_utf8_to_ucs4_fast' : return type 'gunichar' not supported

var utf8ToUtf16Function *gi.Function
var utf8ToUtf16Function_Once sync.Once

func utf8ToUtf16Function_Set() error {
	var err error
	utf8ToUtf16Function_Once.Do(func() {
		utf8ToUtf16Function, err = gi.FunctionInvokerNew("GLib", "utf8_to_utf16")
	})
	return err
}

// Utf8ToUtf16 is a representation of the C type g_utf8_to_utf16.
func Utf8ToUtf16(str string, len int64) (uint16, int64, int64, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(str)
	inArgs[1].SetInt64(len)

	var outArgs [2]gi.Argument
	var ret gi.Argument

	err := utf8ToUtf16Function_Set()
	if err == nil {
		ret = utf8ToUtf16Function.Invoke(inArgs[:], outArgs[:])
	}

	retGo := ret.Uint16()
	out0 := outArgs[0].Int64()
	out1 := outArgs[1].Int64()

	return retGo, out0, out1, err
}

// UNSUPPORTED : C value 'g_utf8_validate' : parameter 'str' has no type

// UNSUPPORTED : C value 'g_utf8_validate_len' : parameter 'str' has no type

var uuidStringIsValidFunction *gi.Function
var uuidStringIsValidFunction_Once sync.Once

func uuidStringIsValidFunction_Set() error {
	var err error
	uuidStringIsValidFunction_Once.Do(func() {
		uuidStringIsValidFunction, err = gi.FunctionInvokerNew("GLib", "uuid_string_is_valid")
	})
	return err
}

// UuidStringIsValid is a representation of the C type g_uuid_string_is_valid.
func UuidStringIsValid(str string) (bool, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(str)

	var ret gi.Argument

	err := uuidStringIsValidFunction_Set()
	if err == nil {
		ret = uuidStringIsValidFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

var uuidStringRandomFunction *gi.Function
var uuidStringRandomFunction_Once sync.Once

func uuidStringRandomFunction_Set() error {
	var err error
	uuidStringRandomFunction_Once.Do(func() {
		uuidStringRandomFunction, err = gi.FunctionInvokerNew("GLib", "uuid_string_random")
	})
	return err
}

// UuidStringRandom is a representation of the C type g_uuid_string_random.
func UuidStringRandom() (string, error) {

	var ret gi.Argument

	err := uuidStringRandomFunction_Set()
	if err == nil {
		ret = uuidStringRandomFunction.Invoke(nil, nil)
	}

	retGo := ret.String(true)

	return retGo, err
}

// UNSUPPORTED : C value 'g_variant_get_gtype' : return type 'GType' not supported

var variantIsObjectPathFunction *gi.Function
var variantIsObjectPathFunction_Once sync.Once

func variantIsObjectPathFunction_Set() error {
	var err error
	variantIsObjectPathFunction_Once.Do(func() {
		variantIsObjectPathFunction, err = gi.FunctionInvokerNew("GLib", "variant_is_object_path")
	})
	return err
}

// VariantIsObjectPath is a representation of the C type g_variant_is_object_path.
func VariantIsObjectPath(string_ string) (bool, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(string_)

	var ret gi.Argument

	err := variantIsObjectPathFunction_Set()
	if err == nil {
		ret = variantIsObjectPathFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

var variantIsSignatureFunction *gi.Function
var variantIsSignatureFunction_Once sync.Once

func variantIsSignatureFunction_Set() error {
	var err error
	variantIsSignatureFunction_Once.Do(func() {
		variantIsSignatureFunction, err = gi.FunctionInvokerNew("GLib", "variant_is_signature")
	})
	return err
}

// VariantIsSignature is a representation of the C type g_variant_is_signature.
func VariantIsSignature(string_ string) (bool, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(string_)

	var ret gi.Argument

	err := variantIsSignatureFunction_Set()
	if err == nil {
		ret = variantIsSignatureFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

// UNSUPPORTED : C value 'g_variant_parse' : parameter 'type' of type 'VariantType' not supported

// UNSUPPORTED : C value 'g_variant_parse_error_print_context' : parameter 'error' of type 'Error' not supported

// UNSUPPORTED : C value 'g_variant_parse_error_quark' : return type 'Quark' not supported

// UNSUPPORTED : C value 'g_variant_parser_get_error_quark' : return type 'Quark' not supported

var variantTypeCheckedFunction *gi.Function
var variantTypeCheckedFunction_Once sync.Once

func variantTypeCheckedFunction_Set() error {
	var err error
	variantTypeCheckedFunction_Once.Do(func() {
		variantTypeCheckedFunction, err = gi.FunctionInvokerNew("GLib", "variant_type_checked_")
	})
	return err
}

// VariantTypeChecked is a representation of the C type g_variant_type_checked_.
func VariantTypeChecked(arg0 string) (*VariantType, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(arg0)

	var ret gi.Argument

	err := variantTypeCheckedFunction_Set()
	if err == nil {
		ret = variantTypeCheckedFunction.Invoke(inArgs[:], nil)
	}

	retGo := &VariantType{native: ret.Pointer()}

	return retGo, err
}

// UNSUPPORTED : C value 'g_variant_type_string_get_depth_' : return type 'gsize' not supported

var variantTypeStringIsValidFunction *gi.Function
var variantTypeStringIsValidFunction_Once sync.Once

func variantTypeStringIsValidFunction_Set() error {
	var err error
	variantTypeStringIsValidFunction_Once.Do(func() {
		variantTypeStringIsValidFunction, err = gi.FunctionInvokerNew("GLib", "variant_type_string_is_valid")
	})
	return err
}

// VariantTypeStringIsValid is a representation of the C type g_variant_type_string_is_valid.
func VariantTypeStringIsValid(typeString string) (bool, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(typeString)

	var ret gi.Argument

	err := variantTypeStringIsValidFunction_Set()
	if err == nil {
		ret = variantTypeStringIsValidFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

var variantTypeStringScanFunction *gi.Function
var variantTypeStringScanFunction_Once sync.Once

func variantTypeStringScanFunction_Set() error {
	var err error
	variantTypeStringScanFunction_Once.Do(func() {
		variantTypeStringScanFunction, err = gi.FunctionInvokerNew("GLib", "variant_type_string_scan")
	})
	return err
}

// VariantTypeStringScan is a representation of the C type g_variant_type_string_scan.
func VariantTypeStringScan(string_ string, limit string) (bool, string, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(string_)
	inArgs[1].SetString(limit)

	var outArgs [1]gi.Argument
	var ret gi.Argument

	err := variantTypeStringScanFunction_Set()
	if err == nil {
		ret = variantTypeStringScanFunction.Invoke(inArgs[:], outArgs[:])
	}

	retGo := ret.Boolean()
	out0 := outArgs[0].String(true)

	return retGo, out0, err
}

// UNSUPPORTED : C value 'g_vasprintf' : parameter 'args' of type 'va_list' not supported

// UNSUPPORTED : C value 'g_vfprintf' : parameter 'file' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_vprintf' : parameter 'args' of type 'va_list' not supported

// UNSUPPORTED : C value 'g_vsnprintf' : parameter 'args' of type 'va_list' not supported

// UNSUPPORTED : C value 'g_vsprintf' : parameter 'args' of type 'va_list' not supported

var warnMessageFunction *gi.Function
var warnMessageFunction_Once sync.Once

func warnMessageFunction_Set() error {
	var err error
	warnMessageFunction_Once.Do(func() {
		warnMessageFunction, err = gi.FunctionInvokerNew("GLib", "warn_message")
	})
	return err
}

// WarnMessage is a representation of the C type g_warn_message.
func WarnMessage(domain string, file string, line int32, func_ string, warnexpr string) error {
	var inArgs [5]gi.Argument
	inArgs[0].SetString(domain)
	inArgs[1].SetString(file)
	inArgs[2].SetInt32(line)
	inArgs[3].SetString(func_)
	inArgs[4].SetString(warnexpr)

	err := warnMessageFunction_Set()
	if err == nil {
		warnMessageFunction.Invoke(inArgs[:], nil)
	}

	return err
}
