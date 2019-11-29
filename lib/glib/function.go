// Code generated - DO NOT EDIT.

package glib

import (
	gi "github.com/pekim/gobbi/internal/gi"
	"sync"
)

var accessFunction *gi.Function
var accessFunction_Once sync.Once

func accessFunction_Set() error {
	var err error
	accessFunction_Once.Do(func() {
		accessFunction, err = gi.FunctionInvokerNew("GLib", "access")
	})
	return err
}

// Access is a representation of the C type g_access.
func Access(filename string, mode int32) int32 {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(filename)
	inArgs[1].SetInt32(mode)

	var ret gi.Argument

	err := accessFunction_Set()
	if err == nil {
		ret = accessFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

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
func AsciiDigitValue(c int8) int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetInt8(c)

	var ret gi.Argument

	err := asciiDigitValueFunction_Set()
	if err == nil {
		ret = asciiDigitValueFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
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
func AsciiDtostr(buffer string, bufLen int32, d float64) string {
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

	return retGo
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
func AsciiFormatd(buffer string, bufLen int32, format string, d float64) string {
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

	return retGo
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
func AsciiStrcasecmp(s1 string, s2 string) int32 {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(s1)
	inArgs[1].SetString(s2)

	var ret gi.Argument

	err := asciiStrcasecmpFunction_Set()
	if err == nil {
		ret = asciiStrcasecmpFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
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
func AsciiStrdown(str string, len int32) string {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(str)
	inArgs[1].SetInt32(len)

	var ret gi.Argument

	err := asciiStrdownFunction_Set()
	if err == nil {
		ret = asciiStrdownFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
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
func AsciiStringToSigned(str string, base uint32, min int64, max int64) (bool, int64) {
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

	return retGo, out0
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
func AsciiStringToUnsigned(str string, base uint32, min uint64, max uint64) (bool, uint64) {
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

	return retGo, out0
}

var asciiStrncasecmpFunction *gi.Function
var asciiStrncasecmpFunction_Once sync.Once

func asciiStrncasecmpFunction_Set() error {
	var err error
	asciiStrncasecmpFunction_Once.Do(func() {
		asciiStrncasecmpFunction, err = gi.FunctionInvokerNew("GLib", "ascii_strncasecmp")
	})
	return err
}

// AsciiStrncasecmp is a representation of the C type g_ascii_strncasecmp.
func AsciiStrncasecmp(s1 string, s2 string, n uint64) int32 {
	var inArgs [3]gi.Argument
	inArgs[0].SetString(s1)
	inArgs[1].SetString(s2)
	inArgs[2].SetUint64(n)

	var ret gi.Argument

	err := asciiStrncasecmpFunction_Set()
	if err == nil {
		ret = asciiStrncasecmpFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

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
func AsciiStrtod(nptr string) (float64, string) {
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

	return retGo, out0
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
func AsciiStrtoll(nptr string, base uint32) (int64, string) {
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

	return retGo, out0
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
func AsciiStrtoull(nptr string, base uint32) (uint64, string) {
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

	return retGo, out0
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
func AsciiStrup(str string, len int32) string {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(str)
	inArgs[1].SetInt32(len)

	var ret gi.Argument

	err := asciiStrupFunction_Set()
	if err == nil {
		ret = asciiStrupFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
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
func AsciiTolower(c int8) int8 {
	var inArgs [1]gi.Argument
	inArgs[0].SetInt8(c)

	var ret gi.Argument

	err := asciiTolowerFunction_Set()
	if err == nil {
		ret = asciiTolowerFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int8()

	return retGo
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
func AsciiToupper(c int8) int8 {
	var inArgs [1]gi.Argument
	inArgs[0].SetInt8(c)

	var ret gi.Argument

	err := asciiToupperFunction_Set()
	if err == nil {
		ret = asciiToupperFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int8()

	return retGo
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
func AsciiXdigitValue(c int8) int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetInt8(c)

	var ret gi.Argument

	err := asciiXdigitValueFunction_Set()
	if err == nil {
		ret = asciiXdigitValueFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
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
func AssertWarning(logDomain string, file string, line int32, prettyFunction string, expression string) {
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

	return
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
func AssertionMessage(domain string, file string, line int32, func_ string, message string) {
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

	return
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
func AssertionMessageCmpstr(domain string, file string, line int32, func_ string, expr string, arg1 string, cmp string, arg2 string) {
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

	return
}

var assertionMessageErrorFunction *gi.Function
var assertionMessageErrorFunction_Once sync.Once

func assertionMessageErrorFunction_Set() error {
	var err error
	assertionMessageErrorFunction_Once.Do(func() {
		assertionMessageErrorFunction, err = gi.FunctionInvokerNew("GLib", "assertion_message_error")
	})
	return err
}

// AssertionMessageError is a representation of the C type g_assertion_message_error.
func AssertionMessageError(domain string, file string, line int32, func_ string, expr string, error *Error, errorDomain Quark, errorCode int32) {
	var inArgs [8]gi.Argument
	inArgs[0].SetString(domain)
	inArgs[1].SetString(file)
	inArgs[2].SetInt32(line)
	inArgs[3].SetString(func_)
	inArgs[4].SetString(expr)
	inArgs[5].SetPointer(error.native)
	inArgs[6].SetUint32(uint32(errorDomain))
	inArgs[7].SetInt32(errorCode)

	err := assertionMessageErrorFunction_Set()
	if err == nil {
		assertionMessageErrorFunction.Invoke(inArgs[:], nil)
	}

	return
}

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
func AssertionMessageExpr(domain string, file string, line int32, func_ string, expr string) {
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

	return
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
func AtomicIntAdd(atomic int32, val int32) int32 {
	var inArgs [2]gi.Argument
	inArgs[0].SetInt32(atomic)
	inArgs[1].SetInt32(val)

	var ret gi.Argument

	err := atomicIntAddFunction_Set()
	if err == nil {
		ret = atomicIntAddFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
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
func AtomicIntAnd(atomic uint32, val uint32) uint32 {
	var inArgs [2]gi.Argument
	inArgs[0].SetUint32(atomic)
	inArgs[1].SetUint32(val)

	var ret gi.Argument

	err := atomicIntAndFunction_Set()
	if err == nil {
		ret = atomicIntAndFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo
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
func AtomicIntCompareAndExchange(atomic int32, oldval int32, newval int32) bool {
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

	return retGo
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
func AtomicIntDecAndTest(atomic int32) bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetInt32(atomic)

	var ret gi.Argument

	err := atomicIntDecAndTestFunction_Set()
	if err == nil {
		ret = atomicIntDecAndTestFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
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
func AtomicIntExchangeAndAdd(atomic int32, val int32) int32 {
	var inArgs [2]gi.Argument
	inArgs[0].SetInt32(atomic)
	inArgs[1].SetInt32(val)

	var ret gi.Argument

	err := atomicIntExchangeAndAddFunction_Set()
	if err == nil {
		ret = atomicIntExchangeAndAddFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
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
func AtomicIntGet(atomic int32) int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetInt32(atomic)

	var ret gi.Argument

	err := atomicIntGetFunction_Set()
	if err == nil {
		ret = atomicIntGetFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
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
func AtomicIntInc(atomic int32) {
	var inArgs [1]gi.Argument
	inArgs[0].SetInt32(atomic)

	err := atomicIntIncFunction_Set()
	if err == nil {
		atomicIntIncFunction.Invoke(inArgs[:], nil)
	}

	return
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
func AtomicIntOr(atomic uint32, val uint32) uint32 {
	var inArgs [2]gi.Argument
	inArgs[0].SetUint32(atomic)
	inArgs[1].SetUint32(val)

	var ret gi.Argument

	err := atomicIntOrFunction_Set()
	if err == nil {
		ret = atomicIntOrFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo
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
func AtomicIntSet(atomic int32, newval int32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetInt32(atomic)
	inArgs[1].SetInt32(newval)

	err := atomicIntSetFunction_Set()
	if err == nil {
		atomicIntSetFunction.Invoke(inArgs[:], nil)
	}

	return
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
func AtomicIntXor(atomic uint32, val uint32) uint32 {
	var inArgs [2]gi.Argument
	inArgs[0].SetUint32(atomic)
	inArgs[1].SetUint32(val)

	var ret gi.Argument

	err := atomicIntXorFunction_Set()
	if err == nil {
		ret = atomicIntXorFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo
}

// UNSUPPORTED : C value 'g_atomic_pointer_add' : parameter 'atomic' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_atomic_pointer_and' : parameter 'atomic' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_atomic_pointer_compare_and_exchange' : parameter 'atomic' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_atomic_pointer_get' : parameter 'atomic' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_atomic_pointer_or' : parameter 'atomic' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_atomic_pointer_set' : parameter 'atomic' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_atomic_pointer_xor' : parameter 'atomic' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_atomic_rc_box_acquire' : parameter 'mem_block' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_atomic_rc_box_alloc' : return type 'gpointer' not supported

// UNSUPPORTED : C value 'g_atomic_rc_box_alloc0' : return type 'gpointer' not supported

// UNSUPPORTED : C value 'g_atomic_rc_box_dup' : parameter 'mem_block' of type 'gpointer' not supported

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
func AtomicRefCountCompare(arc int32, val int32) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetInt32(arc)
	inArgs[1].SetInt32(val)

	var ret gi.Argument

	err := atomicRefCountCompareFunction_Set()
	if err == nil {
		ret = atomicRefCountCompareFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
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
func AtomicRefCountDec(arc int32) bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetInt32(arc)

	var ret gi.Argument

	err := atomicRefCountDecFunction_Set()
	if err == nil {
		ret = atomicRefCountDecFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
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
func AtomicRefCountInc(arc int32) {
	var inArgs [1]gi.Argument
	inArgs[0].SetInt32(arc)

	err := atomicRefCountIncFunction_Set()
	if err == nil {
		atomicRefCountIncFunction.Invoke(inArgs[:], nil)
	}

	return
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
func AtomicRefCountInit(arc int32) {
	var inArgs [1]gi.Argument
	inArgs[0].SetInt32(arc)

	err := atomicRefCountInitFunction_Set()
	if err == nil {
		atomicRefCountInitFunction.Invoke(inArgs[:], nil)
	}

	return
}

var base64DecodeFunction *gi.Function
var base64DecodeFunction_Once sync.Once

func base64DecodeFunction_Set() error {
	var err error
	base64DecodeFunction_Once.Do(func() {
		base64DecodeFunction, err = gi.FunctionInvokerNew("GLib", "base64_decode")
	})
	return err
}

// Base64Decode is a representation of the C type g_base64_decode.
func Base64Decode(text string) uint64 {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(text)

	var outArgs [1]gi.Argument

	err := base64DecodeFunction_Set()
	if err == nil {
		base64DecodeFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := outArgs[0].Uint64()

	return out0
}

// UNSUPPORTED : C value 'g_base64_decode_inplace' : parameter 'text' has no type

// UNSUPPORTED : C value 'g_base64_decode_step' : parameter 'in' has no type

// UNSUPPORTED : C value 'g_base64_encode' : parameter 'data' has no type

// UNSUPPORTED : C value 'g_base64_encode_close' : parameter 'out' has no type

// UNSUPPORTED : C value 'g_base64_encode_step' : parameter 'in' has no type

var basenameFunction *gi.Function
var basenameFunction_Once sync.Once

func basenameFunction_Set() error {
	var err error
	basenameFunction_Once.Do(func() {
		basenameFunction, err = gi.FunctionInvokerNew("GLib", "basename")
	})
	return err
}

// Basename is a representation of the C type g_basename.
func Basename(fileName string) string {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(fileName)

	var ret gi.Argument

	err := basenameFunction_Set()
	if err == nil {
		ret = basenameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

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
func BitLock(address int32, lockBit int32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetInt32(address)
	inArgs[1].SetInt32(lockBit)

	err := bitLockFunction_Set()
	if err == nil {
		bitLockFunction.Invoke(inArgs[:], nil)
	}

	return
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
func BitNthLsf(mask uint64, nthBit int32) int32 {
	var inArgs [2]gi.Argument
	inArgs[0].SetUint64(mask)
	inArgs[1].SetInt32(nthBit)

	var ret gi.Argument

	err := bitNthLsfFunction_Set()
	if err == nil {
		ret = bitNthLsfFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
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
func BitNthMsf(mask uint64, nthBit int32) int32 {
	var inArgs [2]gi.Argument
	inArgs[0].SetUint64(mask)
	inArgs[1].SetInt32(nthBit)

	var ret gi.Argument

	err := bitNthMsfFunction_Set()
	if err == nil {
		ret = bitNthMsfFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
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
func BitStorage(number uint64) uint32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetUint64(number)

	var ret gi.Argument

	err := bitStorageFunction_Set()
	if err == nil {
		ret = bitStorageFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo
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
func BitTrylock(address int32, lockBit int32) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetInt32(address)
	inArgs[1].SetInt32(lockBit)

	var ret gi.Argument

	err := bitTrylockFunction_Set()
	if err == nil {
		ret = bitTrylockFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
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
func BitUnlock(address int32, lockBit int32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetInt32(address)
	inArgs[1].SetInt32(lockBit)

	err := bitUnlockFunction_Set()
	if err == nil {
		bitUnlockFunction.Invoke(inArgs[:], nil)
	}

	return
}

var bookmarkFileErrorQuarkFunction *gi.Function
var bookmarkFileErrorQuarkFunction_Once sync.Once

func bookmarkFileErrorQuarkFunction_Set() error {
	var err error
	bookmarkFileErrorQuarkFunction_Once.Do(func() {
		bookmarkFileErrorQuarkFunction, err = gi.FunctionInvokerNew("GLib", "bookmark_file_error_quark")
	})
	return err
}

// BookmarkFileErrorQuark is a representation of the C type g_bookmark_file_error_quark.
func BookmarkFileErrorQuark() Quark {

	var ret gi.Argument

	err := bookmarkFileErrorQuarkFunction_Set()
	if err == nil {
		ret = bookmarkFileErrorQuarkFunction.Invoke(nil, nil)
	}

	retGo := Quark(ret.Uint32())

	return retGo
}

// UNSUPPORTED : C value 'g_build_filename' : parameter '...' has no type

// UNSUPPORTED : C value 'g_build_filename_valist' : parameter 'args' of type 'va_list' not supported

// UNSUPPORTED : C value 'g_build_filenamev' : parameter 'args' has no type

// UNSUPPORTED : C value 'g_build_path' : parameter '...' has no type

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
func ByteArrayNew() {

	err := byteArrayNewFunction_Set()
	if err == nil {
		byteArrayNewFunction.Invoke(nil, nil)
	}

	return
}

// UNSUPPORTED : C value 'g_byte_array_new_take' : parameter 'data' has no type

// UNSUPPORTED : C value 'g_byte_array_unref' : parameter 'array' has no type

var canonicalizeFilenameFunction *gi.Function
var canonicalizeFilenameFunction_Once sync.Once

func canonicalizeFilenameFunction_Set() error {
	var err error
	canonicalizeFilenameFunction_Once.Do(func() {
		canonicalizeFilenameFunction, err = gi.FunctionInvokerNew("GLib", "canonicalize_filename")
	})
	return err
}

// CanonicalizeFilename is a representation of the C type g_canonicalize_filename.
func CanonicalizeFilename(filename string, relativeTo string) string {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(filename)
	inArgs[1].SetString(relativeTo)

	var ret gi.Argument

	err := canonicalizeFilenameFunction_Set()
	if err == nil {
		ret = canonicalizeFilenameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

var chdirFunction *gi.Function
var chdirFunction_Once sync.Once

func chdirFunction_Set() error {
	var err error
	chdirFunction_Once.Do(func() {
		chdirFunction, err = gi.FunctionInvokerNew("GLib", "chdir")
	})
	return err
}

// Chdir is a representation of the C type g_chdir.
func Chdir(path string) int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(path)

	var ret gi.Argument

	err := chdirFunction_Set()
	if err == nil {
		ret = chdirFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

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
func CheckVersion(requiredMajor uint32, requiredMinor uint32, requiredMicro uint32) string {
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

	return retGo
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
func ChildWatchSourceNew(pid Pid) *Source {
	var inArgs [1]gi.Argument
	inArgs[0].SetInt32(int32(pid))

	var ret gi.Argument

	err := childWatchSourceNewFunction_Set()
	if err == nil {
		ret = childWatchSourceNewFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Source{native: ret.Pointer()}

	return retGo
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
func ClearError() {

	err := clearErrorFunction_Set()
	if err == nil {
		clearErrorFunction.Invoke(nil, nil)
	}

	return
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
func Close(fd int32) bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetInt32(fd)

	var ret gi.Argument

	err := closeFunction_Set()
	if err == nil {
		ret = closeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'g_compute_checksum_for_bytes' : parameter 'checksum_type' of type 'ChecksumType' not supported

// UNSUPPORTED : C value 'g_compute_checksum_for_data' : parameter 'checksum_type' of type 'ChecksumType' not supported

// UNSUPPORTED : C value 'g_compute_checksum_for_string' : parameter 'checksum_type' of type 'ChecksumType' not supported

// UNSUPPORTED : C value 'g_compute_hmac_for_bytes' : parameter 'digest_type' of type 'ChecksumType' not supported

// UNSUPPORTED : C value 'g_compute_hmac_for_data' : parameter 'digest_type' of type 'ChecksumType' not supported

// UNSUPPORTED : C value 'g_compute_hmac_for_string' : parameter 'digest_type' of type 'ChecksumType' not supported

// UNSUPPORTED : C value 'g_convert' : parameter 'str' has no type

var convertErrorQuarkFunction *gi.Function
var convertErrorQuarkFunction_Once sync.Once

func convertErrorQuarkFunction_Set() error {
	var err error
	convertErrorQuarkFunction_Once.Do(func() {
		convertErrorQuarkFunction, err = gi.FunctionInvokerNew("GLib", "convert_error_quark")
	})
	return err
}

// ConvertErrorQuark is a representation of the C type g_convert_error_quark.
func ConvertErrorQuark() Quark {

	var ret gi.Argument

	err := convertErrorQuarkFunction_Set()
	if err == nil {
		ret = convertErrorQuarkFunction.Invoke(nil, nil)
	}

	retGo := Quark(ret.Uint32())

	return retGo
}

// UNSUPPORTED : C value 'g_convert_with_fallback' : parameter 'str' has no type

// UNSUPPORTED : C value 'g_convert_with_iconv' : parameter 'str' has no type

var datalistClearFunction *gi.Function
var datalistClearFunction_Once sync.Once

func datalistClearFunction_Set() error {
	var err error
	datalistClearFunction_Once.Do(func() {
		datalistClearFunction, err = gi.FunctionInvokerNew("GLib", "datalist_clear")
	})
	return err
}

// DatalistClear is a representation of the C type g_datalist_clear.
func DatalistClear(datalist *Data) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(datalist.native)

	err := datalistClearFunction_Set()
	if err == nil {
		datalistClearFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'g_datalist_foreach' : parameter 'func' of type 'DataForeachFunc' not supported

// UNSUPPORTED : C value 'g_datalist_get_data' : return type 'gpointer' not supported

var datalistGetFlagsFunction *gi.Function
var datalistGetFlagsFunction_Once sync.Once

func datalistGetFlagsFunction_Set() error {
	var err error
	datalistGetFlagsFunction_Once.Do(func() {
		datalistGetFlagsFunction, err = gi.FunctionInvokerNew("GLib", "datalist_get_flags")
	})
	return err
}

// DatalistGetFlags is a representation of the C type g_datalist_get_flags.
func DatalistGetFlags(datalist *Data) uint32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(datalist.native)

	var ret gi.Argument

	err := datalistGetFlagsFunction_Set()
	if err == nil {
		ret = datalistGetFlagsFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo
}

// UNSUPPORTED : C value 'g_datalist_id_dup_data' : parameter 'dup_func' of type 'DuplicateFunc' not supported

// UNSUPPORTED : C value 'g_datalist_id_get_data' : return type 'gpointer' not supported

// UNSUPPORTED : C value 'g_datalist_id_remove_no_notify' : return type 'gpointer' not supported

// UNSUPPORTED : C value 'g_datalist_id_replace_data' : parameter 'oldval' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_datalist_id_set_data_full' : parameter 'data' of type 'gpointer' not supported

var datalistInitFunction *gi.Function
var datalistInitFunction_Once sync.Once

func datalistInitFunction_Set() error {
	var err error
	datalistInitFunction_Once.Do(func() {
		datalistInitFunction, err = gi.FunctionInvokerNew("GLib", "datalist_init")
	})
	return err
}

// DatalistInit is a representation of the C type g_datalist_init.
func DatalistInit(datalist *Data) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(datalist.native)

	err := datalistInitFunction_Set()
	if err == nil {
		datalistInitFunction.Invoke(inArgs[:], nil)
	}

	return
}

var datalistSetFlagsFunction *gi.Function
var datalistSetFlagsFunction_Once sync.Once

func datalistSetFlagsFunction_Set() error {
	var err error
	datalistSetFlagsFunction_Once.Do(func() {
		datalistSetFlagsFunction, err = gi.FunctionInvokerNew("GLib", "datalist_set_flags")
	})
	return err
}

// DatalistSetFlags is a representation of the C type g_datalist_set_flags.
func DatalistSetFlags(datalist *Data, flags uint32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(datalist.native)
	inArgs[1].SetUint32(flags)

	err := datalistSetFlagsFunction_Set()
	if err == nil {
		datalistSetFlagsFunction.Invoke(inArgs[:], nil)
	}

	return
}

var datalistUnsetFlagsFunction *gi.Function
var datalistUnsetFlagsFunction_Once sync.Once

func datalistUnsetFlagsFunction_Set() error {
	var err error
	datalistUnsetFlagsFunction_Once.Do(func() {
		datalistUnsetFlagsFunction, err = gi.FunctionInvokerNew("GLib", "datalist_unset_flags")
	})
	return err
}

// DatalistUnsetFlags is a representation of the C type g_datalist_unset_flags.
func DatalistUnsetFlags(datalist *Data, flags uint32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(datalist.native)
	inArgs[1].SetUint32(flags)

	err := datalistUnsetFlagsFunction_Set()
	if err == nil {
		datalistUnsetFlagsFunction.Invoke(inArgs[:], nil)
	}

	return
}

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
func DateGetMondayWeeksInYear(year DateYear) uint8 {
	var inArgs [1]gi.Argument
	inArgs[0].SetUint16(uint16(year))

	var ret gi.Argument

	err := dateGetMondayWeeksInYearFunction_Set()
	if err == nil {
		ret = dateGetMondayWeeksInYearFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint8()

	return retGo
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
func DateGetSundayWeeksInYear(year DateYear) uint8 {
	var inArgs [1]gi.Argument
	inArgs[0].SetUint16(uint16(year))

	var ret gi.Argument

	err := dateGetSundayWeeksInYearFunction_Set()
	if err == nil {
		ret = dateGetSundayWeeksInYearFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint8()

	return retGo
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
func DateIsLeapYear(year DateYear) bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetUint16(uint16(year))

	var ret gi.Argument

	err := dateIsLeapYearFunction_Set()
	if err == nil {
		ret = dateIsLeapYearFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var dateStrftimeFunction *gi.Function
var dateStrftimeFunction_Once sync.Once

func dateStrftimeFunction_Set() error {
	var err error
	dateStrftimeFunction_Once.Do(func() {
		dateStrftimeFunction, err = gi.FunctionInvokerNew("GLib", "date_strftime")
	})
	return err
}

// DateStrftime is a representation of the C type g_date_strftime.
func DateStrftime(s string, slen uint64, format string, date *Date) uint64 {
	var inArgs [4]gi.Argument
	inArgs[0].SetString(s)
	inArgs[1].SetUint64(slen)
	inArgs[2].SetString(format)
	inArgs[3].SetPointer(date.native)

	var ret gi.Argument

	err := dateStrftimeFunction_Set()
	if err == nil {
		ret = dateStrftimeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint64()

	return retGo
}

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
func DateValidDay(day DateDay) bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetUint8(uint8(day))

	var ret gi.Argument

	err := dateValidDayFunction_Set()
	if err == nil {
		ret = dateValidDayFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
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
func DateValidJulian(julianDate uint32) bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetUint32(julianDate)

	var ret gi.Argument

	err := dateValidJulianFunction_Set()
	if err == nil {
		ret = dateValidJulianFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
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
func DateValidYear(year DateYear) bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetUint16(uint16(year))

	var ret gi.Argument

	err := dateValidYearFunction_Set()
	if err == nil {
		ret = dateValidYearFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
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
func Dcgettext(domain string, msgid string, category int32) string {
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

	return retGo
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
func Dgettext(domain string, msgid string) string {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(domain)
	inArgs[1].SetString(msgid)

	var ret gi.Argument

	err := dgettextFunction_Set()
	if err == nil {
		ret = dgettextFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var dirMakeTmpFunction *gi.Function
var dirMakeTmpFunction_Once sync.Once

func dirMakeTmpFunction_Set() error {
	var err error
	dirMakeTmpFunction_Once.Do(func() {
		dirMakeTmpFunction, err = gi.FunctionInvokerNew("GLib", "dir_make_tmp")
	})
	return err
}

// DirMakeTmp is a representation of the C type g_dir_make_tmp.
func DirMakeTmp(tmpl string) string {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(tmpl)

	var ret gi.Argument

	err := dirMakeTmpFunction_Set()
	if err == nil {
		ret = dirMakeTmpFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

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
func Dngettext(domain string, msgid string, msgidPlural string, n uint64) string {
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

	return retGo
}

// UNSUPPORTED : C value 'g_double_equal' : parameter 'v1' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_double_hash' : parameter 'v' of type 'gpointer' not supported

var dpgettextFunction *gi.Function
var dpgettextFunction_Once sync.Once

func dpgettextFunction_Set() error {
	var err error
	dpgettextFunction_Once.Do(func() {
		dpgettextFunction, err = gi.FunctionInvokerNew("GLib", "dpgettext")
	})
	return err
}

// Dpgettext is a representation of the C type g_dpgettext.
func Dpgettext(domain string, msgctxtid string, msgidoffset uint64) string {
	var inArgs [3]gi.Argument
	inArgs[0].SetString(domain)
	inArgs[1].SetString(msgctxtid)
	inArgs[2].SetUint64(msgidoffset)

	var ret gi.Argument

	err := dpgettextFunction_Set()
	if err == nil {
		ret = dpgettextFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

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
func Dpgettext2(domain string, context string, msgid string) string {
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

	return retGo
}

// UNSUPPORTED : C value 'g_environ_getenv' : parameter 'envp' has no type

// UNSUPPORTED : C value 'g_environ_setenv' : parameter 'envp' has no type

// UNSUPPORTED : C value 'g_environ_unsetenv' : parameter 'envp' has no type

// UNSUPPORTED : C value 'g_file_error_from_errno' : return type 'FileError' not supported

var fileErrorQuarkFunction *gi.Function
var fileErrorQuarkFunction_Once sync.Once

func fileErrorQuarkFunction_Set() error {
	var err error
	fileErrorQuarkFunction_Once.Do(func() {
		fileErrorQuarkFunction, err = gi.FunctionInvokerNew("GLib", "file_error_quark")
	})
	return err
}

// FileErrorQuark is a representation of the C type g_file_error_quark.
func FileErrorQuark() Quark {

	var ret gi.Argument

	err := fileErrorQuarkFunction_Set()
	if err == nil {
		ret = fileErrorQuarkFunction.Invoke(nil, nil)
	}

	retGo := Quark(ret.Uint32())

	return retGo
}

// UNSUPPORTED : C value 'g_file_get_contents' : parameter 'contents' has no type

var fileOpenTmpFunction *gi.Function
var fileOpenTmpFunction_Once sync.Once

func fileOpenTmpFunction_Set() error {
	var err error
	fileOpenTmpFunction_Once.Do(func() {
		fileOpenTmpFunction, err = gi.FunctionInvokerNew("GLib", "file_open_tmp")
	})
	return err
}

// FileOpenTmp is a representation of the C type g_file_open_tmp.
func FileOpenTmp(tmpl string) (int32, string) {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(tmpl)

	var outArgs [1]gi.Argument
	var ret gi.Argument

	err := fileOpenTmpFunction_Set()
	if err == nil {
		ret = fileOpenTmpFunction.Invoke(inArgs[:], outArgs[:])
	}

	retGo := ret.Int32()
	out0 := outArgs[0].String(true)

	return retGo, out0
}

var fileReadLinkFunction *gi.Function
var fileReadLinkFunction_Once sync.Once

func fileReadLinkFunction_Set() error {
	var err error
	fileReadLinkFunction_Once.Do(func() {
		fileReadLinkFunction, err = gi.FunctionInvokerNew("GLib", "file_read_link")
	})
	return err
}

// FileReadLink is a representation of the C type g_file_read_link.
func FileReadLink(filename string) string {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(filename)

	var ret gi.Argument

	err := fileReadLinkFunction_Set()
	if err == nil {
		ret = fileReadLinkFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

// UNSUPPORTED : C value 'g_file_set_contents' : parameter 'contents' has no type

// UNSUPPORTED : C value 'g_file_test' : parameter 'test' of type 'FileTest' not supported

var filenameDisplayBasenameFunction *gi.Function
var filenameDisplayBasenameFunction_Once sync.Once

func filenameDisplayBasenameFunction_Set() error {
	var err error
	filenameDisplayBasenameFunction_Once.Do(func() {
		filenameDisplayBasenameFunction, err = gi.FunctionInvokerNew("GLib", "filename_display_basename")
	})
	return err
}

// FilenameDisplayBasename is a representation of the C type g_filename_display_basename.
func FilenameDisplayBasename(filename string) string {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(filename)

	var ret gi.Argument

	err := filenameDisplayBasenameFunction_Set()
	if err == nil {
		ret = filenameDisplayBasenameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

var filenameDisplayNameFunction *gi.Function
var filenameDisplayNameFunction_Once sync.Once

func filenameDisplayNameFunction_Set() error {
	var err error
	filenameDisplayNameFunction_Once.Do(func() {
		filenameDisplayNameFunction, err = gi.FunctionInvokerNew("GLib", "filename_display_name")
	})
	return err
}

// FilenameDisplayName is a representation of the C type g_filename_display_name.
func FilenameDisplayName(filename string) string {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(filename)

	var ret gi.Argument

	err := filenameDisplayNameFunction_Set()
	if err == nil {
		ret = filenameDisplayNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

var filenameFromUriFunction *gi.Function
var filenameFromUriFunction_Once sync.Once

func filenameFromUriFunction_Set() error {
	var err error
	filenameFromUriFunction_Once.Do(func() {
		filenameFromUriFunction, err = gi.FunctionInvokerNew("GLib", "filename_from_uri")
	})
	return err
}

// FilenameFromUri is a representation of the C type g_filename_from_uri.
func FilenameFromUri(uri string) (string, string) {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(uri)

	var outArgs [1]gi.Argument
	var ret gi.Argument

	err := filenameFromUriFunction_Set()
	if err == nil {
		ret = filenameFromUriFunction.Invoke(inArgs[:], outArgs[:])
	}

	retGo := ret.String(true)
	out0 := outArgs[0].String(true)

	return retGo, out0
}

var filenameFromUtf8Function *gi.Function
var filenameFromUtf8Function_Once sync.Once

func filenameFromUtf8Function_Set() error {
	var err error
	filenameFromUtf8Function_Once.Do(func() {
		filenameFromUtf8Function, err = gi.FunctionInvokerNew("GLib", "filename_from_utf8")
	})
	return err
}

// FilenameFromUtf8 is a representation of the C type g_filename_from_utf8.
func FilenameFromUtf8(utf8string string, len int32) (string, uint64, uint64) {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(utf8string)
	inArgs[1].SetInt32(len)

	var outArgs [2]gi.Argument
	var ret gi.Argument

	err := filenameFromUtf8Function_Set()
	if err == nil {
		ret = filenameFromUtf8Function.Invoke(inArgs[:], outArgs[:])
	}

	retGo := ret.String(true)
	out0 := outArgs[0].Uint64()
	out1 := outArgs[1].Uint64()

	return retGo, out0, out1
}

var filenameToUriFunction *gi.Function
var filenameToUriFunction_Once sync.Once

func filenameToUriFunction_Set() error {
	var err error
	filenameToUriFunction_Once.Do(func() {
		filenameToUriFunction, err = gi.FunctionInvokerNew("GLib", "filename_to_uri")
	})
	return err
}

// FilenameToUri is a representation of the C type g_filename_to_uri.
func FilenameToUri(filename string, hostname string) string {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(filename)
	inArgs[1].SetString(hostname)

	var ret gi.Argument

	err := filenameToUriFunction_Set()
	if err == nil {
		ret = filenameToUriFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

var filenameToUtf8Function *gi.Function
var filenameToUtf8Function_Once sync.Once

func filenameToUtf8Function_Set() error {
	var err error
	filenameToUtf8Function_Once.Do(func() {
		filenameToUtf8Function, err = gi.FunctionInvokerNew("GLib", "filename_to_utf8")
	})
	return err
}

// FilenameToUtf8 is a representation of the C type g_filename_to_utf8.
func FilenameToUtf8(opsysstring string, len int32) (string, uint64, uint64) {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(opsysstring)
	inArgs[1].SetInt32(len)

	var outArgs [2]gi.Argument
	var ret gi.Argument

	err := filenameToUtf8Function_Set()
	if err == nil {
		ret = filenameToUtf8Function.Invoke(inArgs[:], outArgs[:])
	}

	retGo := ret.String(true)
	out0 := outArgs[0].Uint64()
	out1 := outArgs[1].Uint64()

	return retGo, out0, out1
}

var findProgramInPathFunction *gi.Function
var findProgramInPathFunction_Once sync.Once

func findProgramInPathFunction_Set() error {
	var err error
	findProgramInPathFunction_Once.Do(func() {
		findProgramInPathFunction, err = gi.FunctionInvokerNew("GLib", "find_program_in_path")
	})
	return err
}

// FindProgramInPath is a representation of the C type g_find_program_in_path.
func FindProgramInPath(program string) string {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(program)

	var ret gi.Argument

	err := findProgramInPathFunction_Set()
	if err == nil {
		ret = findProgramInPathFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

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
func FormatSize(size uint64) string {
	var inArgs [1]gi.Argument
	inArgs[0].SetUint64(size)

	var ret gi.Argument

	err := formatSizeFunction_Set()
	if err == nil {
		ret = formatSizeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
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
func FormatSizeForDisplay(size int64) string {
	var inArgs [1]gi.Argument
	inArgs[0].SetInt64(size)

	var ret gi.Argument

	err := formatSizeForDisplayFunction_Set()
	if err == nil {
		ret = formatSizeForDisplayFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
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
func GetApplicationName() string {

	var ret gi.Argument

	err := getApplicationNameFunction_Set()
	if err == nil {
		ret = getApplicationNameFunction.Invoke(nil, nil)
	}

	retGo := ret.String(false)

	return retGo
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
func GetCharset() (bool, string) {

	var outArgs [1]gi.Argument
	var ret gi.Argument

	err := getCharsetFunction_Set()
	if err == nil {
		ret = getCharsetFunction.Invoke(nil, outArgs[:])
	}

	retGo := ret.Boolean()
	out0 := outArgs[0].String(false)

	return retGo, out0
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
func GetCodeset() string {

	var ret gi.Argument

	err := getCodesetFunction_Set()
	if err == nil {
		ret = getCodesetFunction.Invoke(nil, nil)
	}

	retGo := ret.String(true)

	return retGo
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
func GetConsoleCharset() (bool, string) {

	var outArgs [1]gi.Argument
	var ret gi.Argument

	err := getConsoleCharsetFunction_Set()
	if err == nil {
		ret = getConsoleCharsetFunction.Invoke(nil, outArgs[:])
	}

	retGo := ret.Boolean()
	out0 := outArgs[0].String(false)

	return retGo, out0
}

var getCurrentDirFunction *gi.Function
var getCurrentDirFunction_Once sync.Once

func getCurrentDirFunction_Set() error {
	var err error
	getCurrentDirFunction_Once.Do(func() {
		getCurrentDirFunction, err = gi.FunctionInvokerNew("GLib", "get_current_dir")
	})
	return err
}

// GetCurrentDir is a representation of the C type g_get_current_dir.
func GetCurrentDir() string {

	var ret gi.Argument

	err := getCurrentDirFunction_Set()
	if err == nil {
		ret = getCurrentDirFunction.Invoke(nil, nil)
	}

	retGo := ret.String(true)

	return retGo
}

var getCurrentTimeFunction *gi.Function
var getCurrentTimeFunction_Once sync.Once

func getCurrentTimeFunction_Set() error {
	var err error
	getCurrentTimeFunction_Once.Do(func() {
		getCurrentTimeFunction, err = gi.FunctionInvokerNew("GLib", "get_current_time")
	})
	return err
}

// GetCurrentTime is a representation of the C type g_get_current_time.
func GetCurrentTime(result *TimeVal) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(result.native)

	err := getCurrentTimeFunction_Set()
	if err == nil {
		getCurrentTimeFunction.Invoke(inArgs[:], nil)
	}

	return
}

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
func GetEnviron() {

	err := getEnvironFunction_Set()
	if err == nil {
		getEnvironFunction.Invoke(nil, nil)
	}

	return
}

// UNSUPPORTED : C value 'g_get_filename_charsets' : parameter 'filename_charsets' has no type

var getHomeDirFunction *gi.Function
var getHomeDirFunction_Once sync.Once

func getHomeDirFunction_Set() error {
	var err error
	getHomeDirFunction_Once.Do(func() {
		getHomeDirFunction, err = gi.FunctionInvokerNew("GLib", "get_home_dir")
	})
	return err
}

// GetHomeDir is a representation of the C type g_get_home_dir.
func GetHomeDir() string {

	var ret gi.Argument

	err := getHomeDirFunction_Set()
	if err == nil {
		ret = getHomeDirFunction.Invoke(nil, nil)
	}

	retGo := ret.String(false)

	return retGo
}

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
func GetHostName() string {

	var ret gi.Argument

	err := getHostNameFunction_Set()
	if err == nil {
		ret = getHostNameFunction.Invoke(nil, nil)
	}

	retGo := ret.String(false)

	return retGo
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
func GetLanguageNames() {

	err := getLanguageNamesFunction_Set()
	if err == nil {
		getLanguageNamesFunction.Invoke(nil, nil)
	}

	return
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
func GetLanguageNamesWithCategory(categoryName string) {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(categoryName)

	err := getLanguageNamesWithCategoryFunction_Set()
	if err == nil {
		getLanguageNamesWithCategoryFunction.Invoke(inArgs[:], nil)
	}

	return
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
func GetLocaleVariants(locale string) {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(locale)

	err := getLocaleVariantsFunction_Set()
	if err == nil {
		getLocaleVariantsFunction.Invoke(inArgs[:], nil)
	}

	return
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
func GetMonotonicTime() int64 {

	var ret gi.Argument

	err := getMonotonicTimeFunction_Set()
	if err == nil {
		ret = getMonotonicTimeFunction.Invoke(nil, nil)
	}

	retGo := ret.Int64()

	return retGo
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
func GetNumProcessors() uint32 {

	var ret gi.Argument

	err := getNumProcessorsFunction_Set()
	if err == nil {
		ret = getNumProcessorsFunction.Invoke(nil, nil)
	}

	retGo := ret.Uint32()

	return retGo
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
func GetPrgname() string {

	var ret gi.Argument

	err := getPrgnameFunction_Set()
	if err == nil {
		ret = getPrgnameFunction.Invoke(nil, nil)
	}

	retGo := ret.String(false)

	return retGo
}

var getRealNameFunction *gi.Function
var getRealNameFunction_Once sync.Once

func getRealNameFunction_Set() error {
	var err error
	getRealNameFunction_Once.Do(func() {
		getRealNameFunction, err = gi.FunctionInvokerNew("GLib", "get_real_name")
	})
	return err
}

// GetRealName is a representation of the C type g_get_real_name.
func GetRealName() string {

	var ret gi.Argument

	err := getRealNameFunction_Set()
	if err == nil {
		ret = getRealNameFunction.Invoke(nil, nil)
	}

	retGo := ret.String(false)

	return retGo
}

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
func GetRealTime() int64 {

	var ret gi.Argument

	err := getRealTimeFunction_Set()
	if err == nil {
		ret = getRealTimeFunction.Invoke(nil, nil)
	}

	retGo := ret.Int64()

	return retGo
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
func GetSystemConfigDirs() {

	err := getSystemConfigDirsFunction_Set()
	if err == nil {
		getSystemConfigDirsFunction.Invoke(nil, nil)
	}

	return
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
func GetSystemDataDirs() {

	err := getSystemDataDirsFunction_Set()
	if err == nil {
		getSystemDataDirsFunction.Invoke(nil, nil)
	}

	return
}

var getTmpDirFunction *gi.Function
var getTmpDirFunction_Once sync.Once

func getTmpDirFunction_Set() error {
	var err error
	getTmpDirFunction_Once.Do(func() {
		getTmpDirFunction, err = gi.FunctionInvokerNew("GLib", "get_tmp_dir")
	})
	return err
}

// GetTmpDir is a representation of the C type g_get_tmp_dir.
func GetTmpDir() string {

	var ret gi.Argument

	err := getTmpDirFunction_Set()
	if err == nil {
		ret = getTmpDirFunction.Invoke(nil, nil)
	}

	retGo := ret.String(false)

	return retGo
}

var getUserCacheDirFunction *gi.Function
var getUserCacheDirFunction_Once sync.Once

func getUserCacheDirFunction_Set() error {
	var err error
	getUserCacheDirFunction_Once.Do(func() {
		getUserCacheDirFunction, err = gi.FunctionInvokerNew("GLib", "get_user_cache_dir")
	})
	return err
}

// GetUserCacheDir is a representation of the C type g_get_user_cache_dir.
func GetUserCacheDir() string {

	var ret gi.Argument

	err := getUserCacheDirFunction_Set()
	if err == nil {
		ret = getUserCacheDirFunction.Invoke(nil, nil)
	}

	retGo := ret.String(false)

	return retGo
}

var getUserConfigDirFunction *gi.Function
var getUserConfigDirFunction_Once sync.Once

func getUserConfigDirFunction_Set() error {
	var err error
	getUserConfigDirFunction_Once.Do(func() {
		getUserConfigDirFunction, err = gi.FunctionInvokerNew("GLib", "get_user_config_dir")
	})
	return err
}

// GetUserConfigDir is a representation of the C type g_get_user_config_dir.
func GetUserConfigDir() string {

	var ret gi.Argument

	err := getUserConfigDirFunction_Set()
	if err == nil {
		ret = getUserConfigDirFunction.Invoke(nil, nil)
	}

	retGo := ret.String(false)

	return retGo
}

var getUserDataDirFunction *gi.Function
var getUserDataDirFunction_Once sync.Once

func getUserDataDirFunction_Set() error {
	var err error
	getUserDataDirFunction_Once.Do(func() {
		getUserDataDirFunction, err = gi.FunctionInvokerNew("GLib", "get_user_data_dir")
	})
	return err
}

// GetUserDataDir is a representation of the C type g_get_user_data_dir.
func GetUserDataDir() string {

	var ret gi.Argument

	err := getUserDataDirFunction_Set()
	if err == nil {
		ret = getUserDataDirFunction.Invoke(nil, nil)
	}

	retGo := ret.String(false)

	return retGo
}

var getUserNameFunction *gi.Function
var getUserNameFunction_Once sync.Once

func getUserNameFunction_Set() error {
	var err error
	getUserNameFunction_Once.Do(func() {
		getUserNameFunction, err = gi.FunctionInvokerNew("GLib", "get_user_name")
	})
	return err
}

// GetUserName is a representation of the C type g_get_user_name.
func GetUserName() string {

	var ret gi.Argument

	err := getUserNameFunction_Set()
	if err == nil {
		ret = getUserNameFunction.Invoke(nil, nil)
	}

	retGo := ret.String(false)

	return retGo
}

var getUserRuntimeDirFunction *gi.Function
var getUserRuntimeDirFunction_Once sync.Once

func getUserRuntimeDirFunction_Set() error {
	var err error
	getUserRuntimeDirFunction_Once.Do(func() {
		getUserRuntimeDirFunction, err = gi.FunctionInvokerNew("GLib", "get_user_runtime_dir")
	})
	return err
}

// GetUserRuntimeDir is a representation of the C type g_get_user_runtime_dir.
func GetUserRuntimeDir() string {

	var ret gi.Argument

	err := getUserRuntimeDirFunction_Set()
	if err == nil {
		ret = getUserRuntimeDirFunction.Invoke(nil, nil)
	}

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'g_get_user_special_dir' : parameter 'directory' of type 'UserDirectory' not supported

var getenvFunction *gi.Function
var getenvFunction_Once sync.Once

func getenvFunction_Set() error {
	var err error
	getenvFunction_Once.Do(func() {
		getenvFunction, err = gi.FunctionInvokerNew("GLib", "getenv")
	})
	return err
}

// Getenv is a representation of the C type g_getenv.
func Getenv(variable string) string {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(variable)

	var ret gi.Argument

	err := getenvFunction_Set()
	if err == nil {
		ret = getenvFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

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

var hookDestroyFunction *gi.Function
var hookDestroyFunction_Once sync.Once

func hookDestroyFunction_Set() error {
	var err error
	hookDestroyFunction_Once.Do(func() {
		hookDestroyFunction, err = gi.FunctionInvokerNew("GLib", "hook_destroy")
	})
	return err
}

// HookDestroy is a representation of the C type g_hook_destroy.
func HookDestroy(hookList *HookList, hookId uint64) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(hookList.native)
	inArgs[1].SetUint64(hookId)

	var ret gi.Argument

	err := hookDestroyFunction_Set()
	if err == nil {
		ret = hookDestroyFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var hookDestroyLinkFunction *gi.Function
var hookDestroyLinkFunction_Once sync.Once

func hookDestroyLinkFunction_Set() error {
	var err error
	hookDestroyLinkFunction_Once.Do(func() {
		hookDestroyLinkFunction, err = gi.FunctionInvokerNew("GLib", "hook_destroy_link")
	})
	return err
}

// HookDestroyLink is a representation of the C type g_hook_destroy_link.
func HookDestroyLink(hookList *HookList, hook *Hook) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(hookList.native)
	inArgs[1].SetPointer(hook.native)

	err := hookDestroyLinkFunction_Set()
	if err == nil {
		hookDestroyLinkFunction.Invoke(inArgs[:], nil)
	}

	return
}

var hookFreeFunction *gi.Function
var hookFreeFunction_Once sync.Once

func hookFreeFunction_Set() error {
	var err error
	hookFreeFunction_Once.Do(func() {
		hookFreeFunction, err = gi.FunctionInvokerNew("GLib", "hook_free")
	})
	return err
}

// HookFree is a representation of the C type g_hook_free.
func HookFree(hookList *HookList, hook *Hook) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(hookList.native)
	inArgs[1].SetPointer(hook.native)

	err := hookFreeFunction_Set()
	if err == nil {
		hookFreeFunction.Invoke(inArgs[:], nil)
	}

	return
}

var hookInsertBeforeFunction *gi.Function
var hookInsertBeforeFunction_Once sync.Once

func hookInsertBeforeFunction_Set() error {
	var err error
	hookInsertBeforeFunction_Once.Do(func() {
		hookInsertBeforeFunction, err = gi.FunctionInvokerNew("GLib", "hook_insert_before")
	})
	return err
}

// HookInsertBefore is a representation of the C type g_hook_insert_before.
func HookInsertBefore(hookList *HookList, sibling *Hook, hook *Hook) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(hookList.native)
	inArgs[1].SetPointer(sibling.native)
	inArgs[2].SetPointer(hook.native)

	err := hookInsertBeforeFunction_Set()
	if err == nil {
		hookInsertBeforeFunction.Invoke(inArgs[:], nil)
	}

	return
}

var hookPrependFunction *gi.Function
var hookPrependFunction_Once sync.Once

func hookPrependFunction_Set() error {
	var err error
	hookPrependFunction_Once.Do(func() {
		hookPrependFunction, err = gi.FunctionInvokerNew("GLib", "hook_prepend")
	})
	return err
}

// HookPrepend is a representation of the C type g_hook_prepend.
func HookPrepend(hookList *HookList, hook *Hook) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(hookList.native)
	inArgs[1].SetPointer(hook.native)

	err := hookPrependFunction_Set()
	if err == nil {
		hookPrependFunction.Invoke(inArgs[:], nil)
	}

	return
}

var hookUnrefFunction *gi.Function
var hookUnrefFunction_Once sync.Once

func hookUnrefFunction_Set() error {
	var err error
	hookUnrefFunction_Once.Do(func() {
		hookUnrefFunction, err = gi.FunctionInvokerNew("GLib", "hook_unref")
	})
	return err
}

// HookUnref is a representation of the C type g_hook_unref.
func HookUnref(hookList *HookList, hook *Hook) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(hookList.native)
	inArgs[1].SetPointer(hook.native)

	err := hookUnrefFunction_Set()
	if err == nil {
		hookUnrefFunction.Invoke(inArgs[:], nil)
	}

	return
}

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
func HostnameIsAsciiEncoded(hostname string) bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(hostname)

	var ret gi.Argument

	err := hostnameIsAsciiEncodedFunction_Set()
	if err == nil {
		ret = hostnameIsAsciiEncodedFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
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
func HostnameIsIpAddress(hostname string) bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(hostname)

	var ret gi.Argument

	err := hostnameIsIpAddressFunction_Set()
	if err == nil {
		ret = hostnameIsIpAddressFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
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
func HostnameIsNonAscii(hostname string) bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(hostname)

	var ret gi.Argument

	err := hostnameIsNonAsciiFunction_Set()
	if err == nil {
		ret = hostnameIsNonAsciiFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
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
func HostnameToAscii(hostname string) string {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(hostname)

	var ret gi.Argument

	err := hostnameToAsciiFunction_Set()
	if err == nil {
		ret = hostnameToAsciiFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
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
func HostnameToUnicode(hostname string) string {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(hostname)

	var ret gi.Argument

	err := hostnameToUnicodeFunction_Set()
	if err == nil {
		ret = hostnameToUnicodeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

var iconvFunction *gi.Function
var iconvFunction_Once sync.Once

func iconvFunction_Set() error {
	var err error
	iconvFunction_Once.Do(func() {
		iconvFunction, err = gi.FunctionInvokerNew("GLib", "iconv")
	})
	return err
}

// Iconv is a representation of the C type g_iconv.
func Iconv(converter *IConv, inbuf string, inbytesLeft uint64, outbuf string, outbytesLeft uint64) uint64 {
	var inArgs [5]gi.Argument
	inArgs[0].SetPointer(converter.native)
	inArgs[1].SetString(inbuf)
	inArgs[2].SetUint64(inbytesLeft)
	inArgs[3].SetString(outbuf)
	inArgs[4].SetUint64(outbytesLeft)

	var ret gi.Argument

	err := iconvFunction_Set()
	if err == nil {
		ret = iconvFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint64()

	return retGo
}

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
func IconvOpen(toCodeset string, fromCodeset string) *IConv {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(toCodeset)
	inArgs[1].SetString(fromCodeset)

	var ret gi.Argument

	err := iconvOpenFunction_Set()
	if err == nil {
		ret = iconvOpenFunction.Invoke(inArgs[:], nil)
	}

	retGo := &IConv{native: ret.Pointer()}

	return retGo
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
func IdleSourceNew() *Source {

	var ret gi.Argument

	err := idleSourceNewFunction_Set()
	if err == nil {
		ret = idleSourceNewFunction.Invoke(nil, nil)
	}

	retGo := &Source{native: ret.Pointer()}

	return retGo
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
func InternStaticString(string_ string) string {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(string_)

	var ret gi.Argument

	err := internStaticStringFunction_Set()
	if err == nil {
		ret = internStaticStringFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
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
func InternString(string_ string) string {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(string_)

	var ret gi.Argument

	err := internStringFunction_Set()
	if err == nil {
		ret = internStringFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'g_io_add_watch' : parameter 'condition' of type 'IOCondition' not supported

// UNSUPPORTED : C value 'g_io_add_watch_full' : parameter 'condition' of type 'IOCondition' not supported

// UNSUPPORTED : C value 'g_io_channel_error_from_errno' : return type 'IOChannelError' not supported

var ioChannelErrorQuarkFunction *gi.Function
var ioChannelErrorQuarkFunction_Once sync.Once

func ioChannelErrorQuarkFunction_Set() error {
	var err error
	ioChannelErrorQuarkFunction_Once.Do(func() {
		ioChannelErrorQuarkFunction, err = gi.FunctionInvokerNew("GLib", "io_channel_error_quark")
	})
	return err
}

// IoChannelErrorQuark is a representation of the C type g_io_channel_error_quark.
func IoChannelErrorQuark() Quark {

	var ret gi.Argument

	err := ioChannelErrorQuarkFunction_Set()
	if err == nil {
		ret = ioChannelErrorQuarkFunction.Invoke(nil, nil)
	}

	retGo := Quark(ret.Uint32())

	return retGo
}

// UNSUPPORTED : C value 'g_io_create_watch' : parameter 'condition' of type 'IOCondition' not supported

var keyFileErrorQuarkFunction *gi.Function
var keyFileErrorQuarkFunction_Once sync.Once

func keyFileErrorQuarkFunction_Set() error {
	var err error
	keyFileErrorQuarkFunction_Once.Do(func() {
		keyFileErrorQuarkFunction, err = gi.FunctionInvokerNew("GLib", "key_file_error_quark")
	})
	return err
}

// KeyFileErrorQuark is a representation of the C type g_key_file_error_quark.
func KeyFileErrorQuark() Quark {

	var ret gi.Argument

	err := keyFileErrorQuarkFunction_Set()
	if err == nil {
		ret = keyFileErrorQuarkFunction.Invoke(nil, nil)
	}

	retGo := Quark(ret.Uint32())

	return retGo
}

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
func Listenv() {

	err := listenvFunction_Set()
	if err == nil {
		listenvFunction.Invoke(nil, nil)
	}

	return
}

var localeFromUtf8Function *gi.Function
var localeFromUtf8Function_Once sync.Once

func localeFromUtf8Function_Set() error {
	var err error
	localeFromUtf8Function_Once.Do(func() {
		localeFromUtf8Function, err = gi.FunctionInvokerNew("GLib", "locale_from_utf8")
	})
	return err
}

// LocaleFromUtf8 is a representation of the C type g_locale_from_utf8.
func LocaleFromUtf8(utf8string string, len int32) (uint64, uint64) {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(utf8string)
	inArgs[1].SetInt32(len)

	var outArgs [2]gi.Argument

	err := localeFromUtf8Function_Set()
	if err == nil {
		localeFromUtf8Function.Invoke(inArgs[:], outArgs[:])
	}

	out0 := outArgs[0].Uint64()
	out1 := outArgs[1].Uint64()

	return out0, out1
}

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
func LogRemoveHandler(logDomain string, handlerId uint32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(logDomain)
	inArgs[1].SetUint32(handlerId)

	err := logRemoveHandlerFunction_Set()
	if err == nil {
		logRemoveHandlerFunction.Invoke(inArgs[:], nil)
	}

	return
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
func LogWriterIsJournald(outputFd int32) bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetInt32(outputFd)

	var ret gi.Argument

	err := logWriterIsJournaldFunction_Set()
	if err == nil {
		ret = logWriterIsJournaldFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
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
func LogWriterSupportsColor(outputFd int32) bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetInt32(outputFd)

	var ret gi.Argument

	err := logWriterSupportsColorFunction_Set()
	if err == nil {
		ret = logWriterSupportsColorFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
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
func MainContextDefault() *MainContext {

	var ret gi.Argument

	err := mainContextDefaultFunction_Set()
	if err == nil {
		ret = mainContextDefaultFunction.Invoke(nil, nil)
	}

	retGo := &MainContext{native: ret.Pointer()}

	return retGo
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
func MainContextGetThreadDefault() *MainContext {

	var ret gi.Argument

	err := mainContextGetThreadDefaultFunction_Set()
	if err == nil {
		ret = mainContextGetThreadDefaultFunction.Invoke(nil, nil)
	}

	retGo := &MainContext{native: ret.Pointer()}

	return retGo
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
func MainContextRefThreadDefault() *MainContext {

	var ret gi.Argument

	err := mainContextRefThreadDefaultFunction_Set()
	if err == nil {
		ret = mainContextRefThreadDefaultFunction.Invoke(nil, nil)
	}

	retGo := &MainContext{native: ret.Pointer()}

	return retGo
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
func MainCurrentSource() *Source {

	var ret gi.Argument

	err := mainCurrentSourceFunction_Set()
	if err == nil {
		ret = mainCurrentSourceFunction.Invoke(nil, nil)
	}

	retGo := &Source{native: ret.Pointer()}

	return retGo
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
func MainDepth() int32 {

	var ret gi.Argument

	err := mainDepthFunction_Set()
	if err == nil {
		ret = mainDepthFunction.Invoke(nil, nil)
	}

	retGo := ret.Int32()

	return retGo
}

// UNSUPPORTED : C value 'g_malloc' : return type 'gpointer' not supported

// UNSUPPORTED : C value 'g_malloc0' : return type 'gpointer' not supported

// UNSUPPORTED : C value 'g_malloc0_n' : return type 'gpointer' not supported

// UNSUPPORTED : C value 'g_malloc_n' : return type 'gpointer' not supported

// UNSUPPORTED : C value 'g_markup_collect_attributes' : parameter 'first_type' of type 'MarkupCollectType' not supported

var markupErrorQuarkFunction *gi.Function
var markupErrorQuarkFunction_Once sync.Once

func markupErrorQuarkFunction_Set() error {
	var err error
	markupErrorQuarkFunction_Once.Do(func() {
		markupErrorQuarkFunction, err = gi.FunctionInvokerNew("GLib", "markup_error_quark")
	})
	return err
}

// MarkupErrorQuark is a representation of the C type g_markup_error_quark.
func MarkupErrorQuark() Quark {

	var ret gi.Argument

	err := markupErrorQuarkFunction_Set()
	if err == nil {
		ret = markupErrorQuarkFunction.Invoke(nil, nil)
	}

	retGo := Quark(ret.Uint32())

	return retGo
}

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
func MarkupEscapeText(text string, length int32) string {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(text)
	inArgs[1].SetInt32(length)

	var ret gi.Argument

	err := markupEscapeTextFunction_Set()
	if err == nil {
		ret = markupEscapeTextFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
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
func MemIsSystemMalloc() bool {

	var ret gi.Argument

	err := memIsSystemMallocFunction_Set()
	if err == nil {
		ret = memIsSystemMallocFunction.Invoke(nil, nil)
	}

	retGo := ret.Boolean()

	return retGo
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
func MemProfile() {

	err := memProfileFunction_Set()
	if err == nil {
		memProfileFunction.Invoke(nil, nil)
	}

	return
}

var memSetVtableFunction *gi.Function
var memSetVtableFunction_Once sync.Once

func memSetVtableFunction_Set() error {
	var err error
	memSetVtableFunction_Once.Do(func() {
		memSetVtableFunction, err = gi.FunctionInvokerNew("GLib", "mem_set_vtable")
	})
	return err
}

// MemSetVtable is a representation of the C type g_mem_set_vtable.
func MemSetVtable(vtable *MemVTable) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(vtable.native)

	err := memSetVtableFunction_Set()
	if err == nil {
		memSetVtableFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'g_memdup' : parameter 'mem' of type 'gpointer' not supported

var mkdirWithParentsFunction *gi.Function
var mkdirWithParentsFunction_Once sync.Once

func mkdirWithParentsFunction_Set() error {
	var err error
	mkdirWithParentsFunction_Once.Do(func() {
		mkdirWithParentsFunction, err = gi.FunctionInvokerNew("GLib", "mkdir_with_parents")
	})
	return err
}

// MkdirWithParents is a representation of the C type g_mkdir_with_parents.
func MkdirWithParents(pathname string, mode int32) int32 {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(pathname)
	inArgs[1].SetInt32(mode)

	var ret gi.Argument

	err := mkdirWithParentsFunction_Set()
	if err == nil {
		ret = mkdirWithParentsFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var mkdtempFunction *gi.Function
var mkdtempFunction_Once sync.Once

func mkdtempFunction_Set() error {
	var err error
	mkdtempFunction_Once.Do(func() {
		mkdtempFunction, err = gi.FunctionInvokerNew("GLib", "mkdtemp")
	})
	return err
}

// Mkdtemp is a representation of the C type g_mkdtemp.
func Mkdtemp(tmpl string) string {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(tmpl)

	var ret gi.Argument

	err := mkdtempFunction_Set()
	if err == nil {
		ret = mkdtempFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

var mkdtempFullFunction *gi.Function
var mkdtempFullFunction_Once sync.Once

func mkdtempFullFunction_Set() error {
	var err error
	mkdtempFullFunction_Once.Do(func() {
		mkdtempFullFunction, err = gi.FunctionInvokerNew("GLib", "mkdtemp_full")
	})
	return err
}

// MkdtempFull is a representation of the C type g_mkdtemp_full.
func MkdtempFull(tmpl string, mode int32) string {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(tmpl)
	inArgs[1].SetInt32(mode)

	var ret gi.Argument

	err := mkdtempFullFunction_Set()
	if err == nil {
		ret = mkdtempFullFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

var mkstempFunction *gi.Function
var mkstempFunction_Once sync.Once

func mkstempFunction_Set() error {
	var err error
	mkstempFunction_Once.Do(func() {
		mkstempFunction, err = gi.FunctionInvokerNew("GLib", "mkstemp")
	})
	return err
}

// Mkstemp is a representation of the C type g_mkstemp.
func Mkstemp(tmpl string) int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(tmpl)

	var ret gi.Argument

	err := mkstempFunction_Set()
	if err == nil {
		ret = mkstempFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var mkstempFullFunction *gi.Function
var mkstempFullFunction_Once sync.Once

func mkstempFullFunction_Set() error {
	var err error
	mkstempFullFunction_Once.Do(func() {
		mkstempFullFunction, err = gi.FunctionInvokerNew("GLib", "mkstemp_full")
	})
	return err
}

// MkstempFull is a representation of the C type g_mkstemp_full.
func MkstempFull(tmpl string, flags int32, mode int32) int32 {
	var inArgs [3]gi.Argument
	inArgs[0].SetString(tmpl)
	inArgs[1].SetInt32(flags)
	inArgs[2].SetInt32(mode)

	var ret gi.Argument

	err := mkstempFullFunction_Set()
	if err == nil {
		ret = mkstempFullFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

// UNSUPPORTED : C value 'g_nullify_pointer' : parameter 'nullify_location' of type 'gpointer' not supported

var numberParserErrorQuarkFunction *gi.Function
var numberParserErrorQuarkFunction_Once sync.Once

func numberParserErrorQuarkFunction_Set() error {
	var err error
	numberParserErrorQuarkFunction_Once.Do(func() {
		numberParserErrorQuarkFunction, err = gi.FunctionInvokerNew("GLib", "number_parser_error_quark")
	})
	return err
}

// NumberParserErrorQuark is a representation of the C type g_number_parser_error_quark.
func NumberParserErrorQuark() Quark {

	var ret gi.Argument

	err := numberParserErrorQuarkFunction_Set()
	if err == nil {
		ret = numberParserErrorQuarkFunction.Invoke(nil, nil)
	}

	retGo := Quark(ret.Uint32())

	return retGo
}

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
func OnErrorQuery(prgName string) {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(prgName)

	err := onErrorQueryFunction_Set()
	if err == nil {
		onErrorQueryFunction.Invoke(inArgs[:], nil)
	}

	return
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
func OnErrorStackTrace(prgName string) {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(prgName)

	err := onErrorStackTraceFunction_Set()
	if err == nil {
		onErrorStackTraceFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'g_once_init_enter' : parameter 'location' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_once_init_leave' : parameter 'location' of type 'gpointer' not supported

var optionErrorQuarkFunction *gi.Function
var optionErrorQuarkFunction_Once sync.Once

func optionErrorQuarkFunction_Set() error {
	var err error
	optionErrorQuarkFunction_Once.Do(func() {
		optionErrorQuarkFunction, err = gi.FunctionInvokerNew("GLib", "option_error_quark")
	})
	return err
}

// OptionErrorQuark is a representation of the C type g_option_error_quark.
func OptionErrorQuark() Quark {

	var ret gi.Argument

	err := optionErrorQuarkFunction_Set()
	if err == nil {
		ret = optionErrorQuarkFunction.Invoke(nil, nil)
	}

	retGo := Quark(ret.Uint32())

	return retGo
}

// UNSUPPORTED : C value 'g_parse_debug_string' : parameter 'keys' has no type

var pathGetBasenameFunction *gi.Function
var pathGetBasenameFunction_Once sync.Once

func pathGetBasenameFunction_Set() error {
	var err error
	pathGetBasenameFunction_Once.Do(func() {
		pathGetBasenameFunction, err = gi.FunctionInvokerNew("GLib", "path_get_basename")
	})
	return err
}

// PathGetBasename is a representation of the C type g_path_get_basename.
func PathGetBasename(fileName string) string {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(fileName)

	var ret gi.Argument

	err := pathGetBasenameFunction_Set()
	if err == nil {
		ret = pathGetBasenameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

var pathGetDirnameFunction *gi.Function
var pathGetDirnameFunction_Once sync.Once

func pathGetDirnameFunction_Set() error {
	var err error
	pathGetDirnameFunction_Once.Do(func() {
		pathGetDirnameFunction, err = gi.FunctionInvokerNew("GLib", "path_get_dirname")
	})
	return err
}

// PathGetDirname is a representation of the C type g_path_get_dirname.
func PathGetDirname(fileName string) string {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(fileName)

	var ret gi.Argument

	err := pathGetDirnameFunction_Set()
	if err == nil {
		ret = pathGetDirnameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

var pathIsAbsoluteFunction *gi.Function
var pathIsAbsoluteFunction_Once sync.Once

func pathIsAbsoluteFunction_Set() error {
	var err error
	pathIsAbsoluteFunction_Once.Do(func() {
		pathIsAbsoluteFunction, err = gi.FunctionInvokerNew("GLib", "path_is_absolute")
	})
	return err
}

// PathIsAbsolute is a representation of the C type g_path_is_absolute.
func PathIsAbsolute(fileName string) bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(fileName)

	var ret gi.Argument

	err := pathIsAbsoluteFunction_Set()
	if err == nil {
		ret = pathIsAbsoluteFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var pathSkipRootFunction *gi.Function
var pathSkipRootFunction_Once sync.Once

func pathSkipRootFunction_Set() error {
	var err error
	pathSkipRootFunction_Once.Do(func() {
		pathSkipRootFunction, err = gi.FunctionInvokerNew("GLib", "path_skip_root")
	})
	return err
}

// PathSkipRoot is a representation of the C type g_path_skip_root.
func PathSkipRoot(fileName string) string {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(fileName)

	var ret gi.Argument

	err := pathSkipRootFunction_Set()
	if err == nil {
		ret = pathSkipRootFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var patternMatchFunction *gi.Function
var patternMatchFunction_Once sync.Once

func patternMatchFunction_Set() error {
	var err error
	patternMatchFunction_Once.Do(func() {
		patternMatchFunction, err = gi.FunctionInvokerNew("GLib", "pattern_match")
	})
	return err
}

// PatternMatch is a representation of the C type g_pattern_match.
func PatternMatch(pspec *PatternSpec, stringLength uint32, string_ string, stringReversed string) bool {
	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(pspec.native)
	inArgs[1].SetUint32(stringLength)
	inArgs[2].SetString(string_)
	inArgs[3].SetString(stringReversed)

	var ret gi.Argument

	err := patternMatchFunction_Set()
	if err == nil {
		ret = patternMatchFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

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
func PatternMatchSimple(pattern string, string_ string) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(pattern)
	inArgs[1].SetString(string_)

	var ret gi.Argument

	err := patternMatchSimpleFunction_Set()
	if err == nil {
		ret = patternMatchSimpleFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var patternMatchStringFunction *gi.Function
var patternMatchStringFunction_Once sync.Once

func patternMatchStringFunction_Set() error {
	var err error
	patternMatchStringFunction_Once.Do(func() {
		patternMatchStringFunction, err = gi.FunctionInvokerNew("GLib", "pattern_match_string")
	})
	return err
}

// PatternMatchString is a representation of the C type g_pattern_match_string.
func PatternMatchString(pspec *PatternSpec, string_ string) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(pspec.native)
	inArgs[1].SetString(string_)

	var ret gi.Argument

	err := patternMatchStringFunction_Set()
	if err == nil {
		ret = patternMatchStringFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'g_pointer_bit_lock' : parameter 'address' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_pointer_bit_trylock' : parameter 'address' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_pointer_bit_unlock' : parameter 'address' of type 'gpointer' not supported

var pollFunction *gi.Function
var pollFunction_Once sync.Once

func pollFunction_Set() error {
	var err error
	pollFunction_Once.Do(func() {
		pollFunction, err = gi.FunctionInvokerNew("GLib", "poll")
	})
	return err
}

// Poll is a representation of the C type g_poll.
func Poll(fds *PollFD, nfds uint32, timeout int32) int32 {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(fds.native)
	inArgs[1].SetUint32(nfds)
	inArgs[2].SetInt32(timeout)

	var ret gi.Argument

	err := pollFunction_Set()
	if err == nil {
		ret = pollFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

// UNSUPPORTED : C value 'g_prefix_error' : parameter '...' has no type

// UNSUPPORTED : C value 'g_print' : parameter '...' has no type

// UNSUPPORTED : C value 'g_printerr' : parameter '...' has no type

// UNSUPPORTED : C value 'g_printf' : parameter '...' has no type

// UNSUPPORTED : C value 'g_printf_string_upper_bound' : parameter 'args' of type 'va_list' not supported

var propagateErrorFunction *gi.Function
var propagateErrorFunction_Once sync.Once

func propagateErrorFunction_Set() error {
	var err error
	propagateErrorFunction_Once.Do(func() {
		propagateErrorFunction, err = gi.FunctionInvokerNew("GLib", "propagate_error")
	})
	return err
}

// PropagateError is a representation of the C type g_propagate_error.
func PropagateError(src *Error) *Error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(src.native)

	var outArgs [1]gi.Argument

	err := propagateErrorFunction_Set()
	if err == nil {
		propagateErrorFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := &Error{native: outArgs[0].Pointer()}

	return out0
}

// UNSUPPORTED : C value 'g_propagate_prefixed_error' : parameter '...' has no type

// UNSUPPORTED : C value 'g_ptr_array_find' : parameter 'haystack' has no type

// UNSUPPORTED : C value 'g_ptr_array_find_with_equal_func' : parameter 'haystack' has no type

// UNSUPPORTED : C value 'g_qsort_with_data' : parameter 'pbase' of type 'gpointer' not supported

var quarkFromStaticStringFunction *gi.Function
var quarkFromStaticStringFunction_Once sync.Once

func quarkFromStaticStringFunction_Set() error {
	var err error
	quarkFromStaticStringFunction_Once.Do(func() {
		quarkFromStaticStringFunction, err = gi.FunctionInvokerNew("GLib", "quark_from_static_string")
	})
	return err
}

// QuarkFromStaticString is a representation of the C type g_quark_from_static_string.
func QuarkFromStaticString(string_ string) Quark {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(string_)

	var ret gi.Argument

	err := quarkFromStaticStringFunction_Set()
	if err == nil {
		ret = quarkFromStaticStringFunction.Invoke(inArgs[:], nil)
	}

	retGo := Quark(ret.Uint32())

	return retGo
}

var quarkFromStringFunction *gi.Function
var quarkFromStringFunction_Once sync.Once

func quarkFromStringFunction_Set() error {
	var err error
	quarkFromStringFunction_Once.Do(func() {
		quarkFromStringFunction, err = gi.FunctionInvokerNew("GLib", "quark_from_string")
	})
	return err
}

// QuarkFromString is a representation of the C type g_quark_from_string.
func QuarkFromString(string_ string) Quark {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(string_)

	var ret gi.Argument

	err := quarkFromStringFunction_Set()
	if err == nil {
		ret = quarkFromStringFunction.Invoke(inArgs[:], nil)
	}

	retGo := Quark(ret.Uint32())

	return retGo
}

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
func QuarkToString(quark Quark) string {
	var inArgs [1]gi.Argument
	inArgs[0].SetUint32(uint32(quark))

	var ret gi.Argument

	err := quarkToStringFunction_Set()
	if err == nil {
		ret = quarkToStringFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var quarkTryStringFunction *gi.Function
var quarkTryStringFunction_Once sync.Once

func quarkTryStringFunction_Set() error {
	var err error
	quarkTryStringFunction_Once.Do(func() {
		quarkTryStringFunction, err = gi.FunctionInvokerNew("GLib", "quark_try_string")
	})
	return err
}

// QuarkTryString is a representation of the C type g_quark_try_string.
func QuarkTryString(string_ string) Quark {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(string_)

	var ret gi.Argument

	err := quarkTryStringFunction_Set()
	if err == nil {
		ret = quarkTryStringFunction.Invoke(inArgs[:], nil)
	}

	retGo := Quark(ret.Uint32())

	return retGo
}

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
func RandomDouble() float64 {

	var ret gi.Argument

	err := randomDoubleFunction_Set()
	if err == nil {
		ret = randomDoubleFunction.Invoke(nil, nil)
	}

	retGo := ret.Float64()

	return retGo
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
func RandomDoubleRange(begin float64, end float64) float64 {
	var inArgs [2]gi.Argument
	inArgs[0].SetFloat64(begin)
	inArgs[1].SetFloat64(end)

	var ret gi.Argument

	err := randomDoubleRangeFunction_Set()
	if err == nil {
		ret = randomDoubleRangeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Float64()

	return retGo
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
func RandomInt() uint32 {

	var ret gi.Argument

	err := randomIntFunction_Set()
	if err == nil {
		ret = randomIntFunction.Invoke(nil, nil)
	}

	retGo := ret.Uint32()

	return retGo
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
func RandomIntRange(begin int32, end int32) int32 {
	var inArgs [2]gi.Argument
	inArgs[0].SetInt32(begin)
	inArgs[1].SetInt32(end)

	var ret gi.Argument

	err := randomIntRangeFunction_Set()
	if err == nil {
		ret = randomIntRangeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
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
func RandomSetSeed(seed uint32) {
	var inArgs [1]gi.Argument
	inArgs[0].SetUint32(seed)

	err := randomSetSeedFunction_Set()
	if err == nil {
		randomSetSeedFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'g_rc_box_acquire' : parameter 'mem_block' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_rc_box_alloc' : return type 'gpointer' not supported

// UNSUPPORTED : C value 'g_rc_box_alloc0' : return type 'gpointer' not supported

// UNSUPPORTED : C value 'g_rc_box_dup' : parameter 'mem_block' of type 'gpointer' not supported

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
func RefCountCompare(rc int32, val int32) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetInt32(rc)
	inArgs[1].SetInt32(val)

	var ret gi.Argument

	err := refCountCompareFunction_Set()
	if err == nil {
		ret = refCountCompareFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
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
func RefCountDec(rc int32) bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetInt32(rc)

	var ret gi.Argument

	err := refCountDecFunction_Set()
	if err == nil {
		ret = refCountDecFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
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
func RefCountInc(rc int32) {
	var inArgs [1]gi.Argument
	inArgs[0].SetInt32(rc)

	err := refCountIncFunction_Set()
	if err == nil {
		refCountIncFunction.Invoke(inArgs[:], nil)
	}

	return
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
func RefCountInit(rc int32) {
	var inArgs [1]gi.Argument
	inArgs[0].SetInt32(rc)

	err := refCountInitFunction_Set()
	if err == nil {
		refCountInitFunction.Invoke(inArgs[:], nil)
	}

	return
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
func RefStringAcquire(str string) string {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(str)

	var ret gi.Argument

	err := refStringAcquireFunction_Set()
	if err == nil {
		ret = refStringAcquireFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

var refStringLengthFunction *gi.Function
var refStringLengthFunction_Once sync.Once

func refStringLengthFunction_Set() error {
	var err error
	refStringLengthFunction_Once.Do(func() {
		refStringLengthFunction, err = gi.FunctionInvokerNew("GLib", "ref_string_length")
	})
	return err
}

// RefStringLength is a representation of the C type g_ref_string_length.
func RefStringLength(str string) uint64 {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(str)

	var ret gi.Argument

	err := refStringLengthFunction_Set()
	if err == nil {
		ret = refStringLengthFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint64()

	return retGo
}

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
func RefStringNew(str string) string {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(str)

	var ret gi.Argument

	err := refStringNewFunction_Set()
	if err == nil {
		ret = refStringNewFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
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
func RefStringNewIntern(str string) string {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(str)

	var ret gi.Argument

	err := refStringNewInternFunction_Set()
	if err == nil {
		ret = refStringNewInternFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
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
func RefStringNewLen(str string, len int32) string {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(str)
	inArgs[1].SetInt32(len)

	var ret gi.Argument

	err := refStringNewLenFunction_Set()
	if err == nil {
		ret = refStringNewLenFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
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
func RefStringRelease(str string) {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(str)

	err := refStringReleaseFunction_Set()
	if err == nil {
		refStringReleaseFunction.Invoke(inArgs[:], nil)
	}

	return
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
func RegexCheckReplacement(replacement string) (bool, bool) {
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

	return retGo, out0
}

var regexErrorQuarkFunction *gi.Function
var regexErrorQuarkFunction_Once sync.Once

func regexErrorQuarkFunction_Set() error {
	var err error
	regexErrorQuarkFunction_Once.Do(func() {
		regexErrorQuarkFunction, err = gi.FunctionInvokerNew("GLib", "regex_error_quark")
	})
	return err
}

// RegexErrorQuark is a representation of the C type g_regex_error_quark.
func RegexErrorQuark() Quark {

	var ret gi.Argument

	err := regexErrorQuarkFunction_Set()
	if err == nil {
		ret = regexErrorQuarkFunction.Invoke(nil, nil)
	}

	retGo := Quark(ret.Uint32())

	return retGo
}

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
func RegexEscapeNul(string_ string, length int32) string {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(string_)
	inArgs[1].SetInt32(length)

	var ret gi.Argument

	err := regexEscapeNulFunction_Set()
	if err == nil {
		ret = regexEscapeNulFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
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
func ReloadUserSpecialDirsCache() {

	err := reloadUserSpecialDirsCacheFunction_Set()
	if err == nil {
		reloadUserSpecialDirsCacheFunction.Invoke(nil, nil)
	}

	return
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
func ReturnIfFailWarning(logDomain string, prettyFunction string, expression string) {
	var inArgs [3]gi.Argument
	inArgs[0].SetString(logDomain)
	inArgs[1].SetString(prettyFunction)
	inArgs[2].SetString(expression)

	err := returnIfFailWarningFunction_Set()
	if err == nil {
		returnIfFailWarningFunction.Invoke(inArgs[:], nil)
	}

	return
}

var rmdirFunction *gi.Function
var rmdirFunction_Once sync.Once

func rmdirFunction_Set() error {
	var err error
	rmdirFunction_Once.Do(func() {
		rmdirFunction, err = gi.FunctionInvokerNew("GLib", "rmdir")
	})
	return err
}

// Rmdir is a representation of the C type g_rmdir.
func Rmdir(filename string) int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(filename)

	var ret gi.Argument

	err := rmdirFunction_Set()
	if err == nil {
		ret = rmdirFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

// UNSUPPORTED : C value 'g_sequence_get' : return type 'gpointer' not supported

// UNSUPPORTED : C value 'g_sequence_insert_before' : parameter 'data' of type 'gpointer' not supported

var sequenceMoveFunction *gi.Function
var sequenceMoveFunction_Once sync.Once

func sequenceMoveFunction_Set() error {
	var err error
	sequenceMoveFunction_Once.Do(func() {
		sequenceMoveFunction, err = gi.FunctionInvokerNew("GLib", "sequence_move")
	})
	return err
}

// SequenceMove is a representation of the C type g_sequence_move.
func SequenceMove(src *SequenceIter, dest *SequenceIter) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(src.native)
	inArgs[1].SetPointer(dest.native)

	err := sequenceMoveFunction_Set()
	if err == nil {
		sequenceMoveFunction.Invoke(inArgs[:], nil)
	}

	return
}

var sequenceMoveRangeFunction *gi.Function
var sequenceMoveRangeFunction_Once sync.Once

func sequenceMoveRangeFunction_Set() error {
	var err error
	sequenceMoveRangeFunction_Once.Do(func() {
		sequenceMoveRangeFunction, err = gi.FunctionInvokerNew("GLib", "sequence_move_range")
	})
	return err
}

// SequenceMoveRange is a representation of the C type g_sequence_move_range.
func SequenceMoveRange(dest *SequenceIter, begin *SequenceIter, end *SequenceIter) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(dest.native)
	inArgs[1].SetPointer(begin.native)
	inArgs[2].SetPointer(end.native)

	err := sequenceMoveRangeFunction_Set()
	if err == nil {
		sequenceMoveRangeFunction.Invoke(inArgs[:], nil)
	}

	return
}

var sequenceRangeGetMidpointFunction *gi.Function
var sequenceRangeGetMidpointFunction_Once sync.Once

func sequenceRangeGetMidpointFunction_Set() error {
	var err error
	sequenceRangeGetMidpointFunction_Once.Do(func() {
		sequenceRangeGetMidpointFunction, err = gi.FunctionInvokerNew("GLib", "sequence_range_get_midpoint")
	})
	return err
}

// SequenceRangeGetMidpoint is a representation of the C type g_sequence_range_get_midpoint.
func SequenceRangeGetMidpoint(begin *SequenceIter, end *SequenceIter) *SequenceIter {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(begin.native)
	inArgs[1].SetPointer(end.native)

	var ret gi.Argument

	err := sequenceRangeGetMidpointFunction_Set()
	if err == nil {
		ret = sequenceRangeGetMidpointFunction.Invoke(inArgs[:], nil)
	}

	retGo := &SequenceIter{native: ret.Pointer()}

	return retGo
}

var sequenceRemoveFunction *gi.Function
var sequenceRemoveFunction_Once sync.Once

func sequenceRemoveFunction_Set() error {
	var err error
	sequenceRemoveFunction_Once.Do(func() {
		sequenceRemoveFunction, err = gi.FunctionInvokerNew("GLib", "sequence_remove")
	})
	return err
}

// SequenceRemove is a representation of the C type g_sequence_remove.
func SequenceRemove(iter *SequenceIter) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(iter.native)

	err := sequenceRemoveFunction_Set()
	if err == nil {
		sequenceRemoveFunction.Invoke(inArgs[:], nil)
	}

	return
}

var sequenceRemoveRangeFunction *gi.Function
var sequenceRemoveRangeFunction_Once sync.Once

func sequenceRemoveRangeFunction_Set() error {
	var err error
	sequenceRemoveRangeFunction_Once.Do(func() {
		sequenceRemoveRangeFunction, err = gi.FunctionInvokerNew("GLib", "sequence_remove_range")
	})
	return err
}

// SequenceRemoveRange is a representation of the C type g_sequence_remove_range.
func SequenceRemoveRange(begin *SequenceIter, end *SequenceIter) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(begin.native)
	inArgs[1].SetPointer(end.native)

	err := sequenceRemoveRangeFunction_Set()
	if err == nil {
		sequenceRemoveRangeFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'g_sequence_set' : parameter 'data' of type 'gpointer' not supported

var sequenceSwapFunction *gi.Function
var sequenceSwapFunction_Once sync.Once

func sequenceSwapFunction_Set() error {
	var err error
	sequenceSwapFunction_Once.Do(func() {
		sequenceSwapFunction, err = gi.FunctionInvokerNew("GLib", "sequence_swap")
	})
	return err
}

// SequenceSwap is a representation of the C type g_sequence_swap.
func SequenceSwap(a *SequenceIter, b *SequenceIter) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(a.native)
	inArgs[1].SetPointer(b.native)

	err := sequenceSwapFunction_Set()
	if err == nil {
		sequenceSwapFunction.Invoke(inArgs[:], nil)
	}

	return
}

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
func SetApplicationName(applicationName string) {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(applicationName)

	err := setApplicationNameFunction_Set()
	if err == nil {
		setApplicationNameFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'g_set_error' : parameter '...' has no type

var setErrorLiteralFunction *gi.Function
var setErrorLiteralFunction_Once sync.Once

func setErrorLiteralFunction_Set() error {
	var err error
	setErrorLiteralFunction_Once.Do(func() {
		setErrorLiteralFunction, err = gi.FunctionInvokerNew("GLib", "set_error_literal")
	})
	return err
}

// SetErrorLiteral is a representation of the C type g_set_error_literal.
func SetErrorLiteral(domain Quark, code int32, message string) *Error {
	var inArgs [3]gi.Argument
	inArgs[0].SetUint32(uint32(domain))
	inArgs[1].SetInt32(code)
	inArgs[2].SetString(message)

	var outArgs [1]gi.Argument

	err := setErrorLiteralFunction_Set()
	if err == nil {
		setErrorLiteralFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := &Error{native: outArgs[0].Pointer()}

	return out0
}

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
func SetPrgname(prgname string) {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(prgname)

	err := setPrgnameFunction_Set()
	if err == nil {
		setPrgnameFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'g_set_print_handler' : parameter 'func' of type 'PrintFunc' not supported

// UNSUPPORTED : C value 'g_set_printerr_handler' : parameter 'func' of type 'PrintFunc' not supported

var setenvFunction *gi.Function
var setenvFunction_Once sync.Once

func setenvFunction_Set() error {
	var err error
	setenvFunction_Once.Do(func() {
		setenvFunction, err = gi.FunctionInvokerNew("GLib", "setenv")
	})
	return err
}

// Setenv is a representation of the C type g_setenv.
func Setenv(variable string, value string, overwrite bool) bool {
	var inArgs [3]gi.Argument
	inArgs[0].SetString(variable)
	inArgs[1].SetString(value)
	inArgs[2].SetBoolean(overwrite)

	var ret gi.Argument

	err := setenvFunction_Set()
	if err == nil {
		ret = setenvFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var shellErrorQuarkFunction *gi.Function
var shellErrorQuarkFunction_Once sync.Once

func shellErrorQuarkFunction_Set() error {
	var err error
	shellErrorQuarkFunction_Once.Do(func() {
		shellErrorQuarkFunction, err = gi.FunctionInvokerNew("GLib", "shell_error_quark")
	})
	return err
}

// ShellErrorQuark is a representation of the C type g_shell_error_quark.
func ShellErrorQuark() Quark {

	var ret gi.Argument

	err := shellErrorQuarkFunction_Set()
	if err == nil {
		ret = shellErrorQuarkFunction.Invoke(nil, nil)
	}

	retGo := Quark(ret.Uint32())

	return retGo
}

// UNSUPPORTED : C value 'g_shell_parse_argv' : parameter 'argvp' has no type

var shellQuoteFunction *gi.Function
var shellQuoteFunction_Once sync.Once

func shellQuoteFunction_Set() error {
	var err error
	shellQuoteFunction_Once.Do(func() {
		shellQuoteFunction, err = gi.FunctionInvokerNew("GLib", "shell_quote")
	})
	return err
}

// ShellQuote is a representation of the C type g_shell_quote.
func ShellQuote(unquotedString string) string {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(unquotedString)

	var ret gi.Argument

	err := shellQuoteFunction_Set()
	if err == nil {
		ret = shellQuoteFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

var shellUnquoteFunction *gi.Function
var shellUnquoteFunction_Once sync.Once

func shellUnquoteFunction_Set() error {
	var err error
	shellUnquoteFunction_Once.Do(func() {
		shellUnquoteFunction, err = gi.FunctionInvokerNew("GLib", "shell_unquote")
	})
	return err
}

// ShellUnquote is a representation of the C type g_shell_unquote.
func ShellUnquote(quotedString string) string {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(quotedString)

	var ret gi.Argument

	err := shellUnquoteFunction_Set()
	if err == nil {
		ret = shellUnquoteFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

// UNSUPPORTED : C value 'g_slice_alloc' : return type 'gpointer' not supported

// UNSUPPORTED : C value 'g_slice_alloc0' : return type 'gpointer' not supported

// UNSUPPORTED : C value 'g_slice_copy' : parameter 'mem_block' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_slice_free1' : parameter 'mem_block' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_slice_free_chain_with_offset' : parameter 'mem_chain' of type 'gpointer' not supported

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
func SourceRemove(tag uint32) bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetUint32(tag)

	var ret gi.Argument

	err := sourceRemoveFunction_Set()
	if err == nil {
		ret = sourceRemoveFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'g_source_remove_by_funcs_user_data' : parameter 'user_data' of type 'gpointer' not supported

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
func SourceSetNameById(tag uint32, name string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetUint32(tag)
	inArgs[1].SetString(name)

	err := sourceSetNameByIdFunction_Set()
	if err == nil {
		sourceSetNameByIdFunction.Invoke(inArgs[:], nil)
	}

	return
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
func SpacedPrimesClosest(num uint32) uint32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetUint32(num)

	var ret gi.Argument

	err := spacedPrimesClosestFunction_Set()
	if err == nil {
		ret = spacedPrimesClosestFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo
}

// UNSUPPORTED : C value 'g_spawn_async' : parameter 'argv' has no type

// UNSUPPORTED : C value 'g_spawn_async_with_fds' : parameter 'argv' has no type

// UNSUPPORTED : C value 'g_spawn_async_with_pipes' : parameter 'argv' has no type

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
func SpawnCheckExitStatus(exitStatus int32) bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetInt32(exitStatus)

	var ret gi.Argument

	err := spawnCheckExitStatusFunction_Set()
	if err == nil {
		ret = spawnCheckExitStatusFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
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
func SpawnClosePid(pid Pid) {
	var inArgs [1]gi.Argument
	inArgs[0].SetInt32(int32(pid))

	err := spawnClosePidFunction_Set()
	if err == nil {
		spawnClosePidFunction.Invoke(inArgs[:], nil)
	}

	return
}

var spawnCommandLineAsyncFunction *gi.Function
var spawnCommandLineAsyncFunction_Once sync.Once

func spawnCommandLineAsyncFunction_Set() error {
	var err error
	spawnCommandLineAsyncFunction_Once.Do(func() {
		spawnCommandLineAsyncFunction, err = gi.FunctionInvokerNew("GLib", "spawn_command_line_async")
	})
	return err
}

// SpawnCommandLineAsync is a representation of the C type g_spawn_command_line_async.
func SpawnCommandLineAsync(commandLine string) bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(commandLine)

	var ret gi.Argument

	err := spawnCommandLineAsyncFunction_Set()
	if err == nil {
		ret = spawnCommandLineAsyncFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'g_spawn_command_line_sync' : parameter 'standard_output' has no type

var spawnErrorQuarkFunction *gi.Function
var spawnErrorQuarkFunction_Once sync.Once

func spawnErrorQuarkFunction_Set() error {
	var err error
	spawnErrorQuarkFunction_Once.Do(func() {
		spawnErrorQuarkFunction, err = gi.FunctionInvokerNew("GLib", "spawn_error_quark")
	})
	return err
}

// SpawnErrorQuark is a representation of the C type g_spawn_error_quark.
func SpawnErrorQuark() Quark {

	var ret gi.Argument

	err := spawnErrorQuarkFunction_Set()
	if err == nil {
		ret = spawnErrorQuarkFunction.Invoke(nil, nil)
	}

	retGo := Quark(ret.Uint32())

	return retGo
}

var spawnExitErrorQuarkFunction *gi.Function
var spawnExitErrorQuarkFunction_Once sync.Once

func spawnExitErrorQuarkFunction_Set() error {
	var err error
	spawnExitErrorQuarkFunction_Once.Do(func() {
		spawnExitErrorQuarkFunction, err = gi.FunctionInvokerNew("GLib", "spawn_exit_error_quark")
	})
	return err
}

// SpawnExitErrorQuark is a representation of the C type g_spawn_exit_error_quark.
func SpawnExitErrorQuark() Quark {

	var ret gi.Argument

	err := spawnExitErrorQuarkFunction_Set()
	if err == nil {
		ret = spawnExitErrorQuarkFunction.Invoke(nil, nil)
	}

	retGo := Quark(ret.Uint32())

	return retGo
}

// UNSUPPORTED : C value 'g_spawn_sync' : parameter 'argv' has no type

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
func Stpcpy(dest string, src string) string {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(dest)
	inArgs[1].SetString(src)

	var ret gi.Argument

	err := stpcpyFunction_Set()
	if err == nil {
		ret = stpcpyFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
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
func StrHasPrefix(str string, prefix string) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(str)
	inArgs[1].SetString(prefix)

	var ret gi.Argument

	err := strHasPrefixFunction_Set()
	if err == nil {
		ret = strHasPrefixFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
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
func StrHasSuffix(str string, suffix string) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(str)
	inArgs[1].SetString(suffix)

	var ret gi.Argument

	err := strHasSuffixFunction_Set()
	if err == nil {
		ret = strHasSuffixFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
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
func StrIsAscii(str string) bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(str)

	var ret gi.Argument

	err := strIsAsciiFunction_Set()
	if err == nil {
		ret = strIsAsciiFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
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
func StrMatchString(searchTerm string, potentialHit string, acceptAlternates bool) bool {
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

	return retGo
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
func StrToAscii(str string, fromLocale string) string {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(str)
	inArgs[1].SetString(fromLocale)

	var ret gi.Argument

	err := strToAsciiFunction_Set()
	if err == nil {
		ret = strToAsciiFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
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
func Strcanon(string_ string, validChars string, substitutor int8) string {
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

	return retGo
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
func Strcasecmp(s1 string, s2 string) int32 {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(s1)
	inArgs[1].SetString(s2)

	var ret gi.Argument

	err := strcasecmpFunction_Set()
	if err == nil {
		ret = strcasecmpFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
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
func Strchomp(string_ string) string {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(string_)

	var ret gi.Argument

	err := strchompFunction_Set()
	if err == nil {
		ret = strchompFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
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
func Strchug(string_ string) string {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(string_)

	var ret gi.Argument

	err := strchugFunction_Set()
	if err == nil {
		ret = strchugFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
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
func Strcmp0(str1 string, str2 string) int32 {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(str1)
	inArgs[1].SetString(str2)

	var ret gi.Argument

	err := strcmp0Function_Set()
	if err == nil {
		ret = strcmp0Function.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
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
func Strcompress(source string) string {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(source)

	var ret gi.Argument

	err := strcompressFunction_Set()
	if err == nil {
		ret = strcompressFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
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
func Strdelimit(string_ string, delimiters string, newDelimiter int8) string {
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

	return retGo
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
func Strdown(string_ string) string {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(string_)

	var ret gi.Argument

	err := strdownFunction_Set()
	if err == nil {
		ret = strdownFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
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
func Strdup(str string) string {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(str)

	var ret gi.Argument

	err := strdupFunction_Set()
	if err == nil {
		ret = strdupFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
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
func Strdupv(strArray string) {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(strArray)

	err := strdupvFunction_Set()
	if err == nil {
		strdupvFunction.Invoke(inArgs[:], nil)
	}

	return
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
func Strerror(errnum int32) string {
	var inArgs [1]gi.Argument
	inArgs[0].SetInt32(errnum)

	var ret gi.Argument

	err := strerrorFunction_Set()
	if err == nil {
		ret = strerrorFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
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
func Strescape(source string, exceptions string) string {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(source)
	inArgs[1].SetString(exceptions)

	var ret gi.Argument

	err := strescapeFunction_Set()
	if err == nil {
		ret = strescapeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
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
func Strfreev(strArray string) {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(strArray)

	err := strfreevFunction_Set()
	if err == nil {
		strfreevFunction.Invoke(inArgs[:], nil)
	}

	return
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
func StringNew(init string) *String {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(init)

	var ret gi.Argument

	err := stringNewFunction_Set()
	if err == nil {
		ret = stringNewFunction.Invoke(inArgs[:], nil)
	}

	retGo := &String{native: ret.Pointer()}

	return retGo
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
func StringNewLen(init string, len int32) *String {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(init)
	inArgs[1].SetInt32(len)

	var ret gi.Argument

	err := stringNewLenFunction_Set()
	if err == nil {
		ret = stringNewLenFunction.Invoke(inArgs[:], nil)
	}

	retGo := &String{native: ret.Pointer()}

	return retGo
}

var stringSizedNewFunction *gi.Function
var stringSizedNewFunction_Once sync.Once

func stringSizedNewFunction_Set() error {
	var err error
	stringSizedNewFunction_Once.Do(func() {
		stringSizedNewFunction, err = gi.FunctionInvokerNew("GLib", "string_sized_new")
	})
	return err
}

// StringSizedNew is a representation of the C type g_string_sized_new.
func StringSizedNew(dflSize uint64) *String {
	var inArgs [1]gi.Argument
	inArgs[0].SetUint64(dflSize)

	var ret gi.Argument

	err := stringSizedNewFunction_Set()
	if err == nil {
		ret = stringSizedNewFunction.Invoke(inArgs[:], nil)
	}

	retGo := &String{native: ret.Pointer()}

	return retGo
}

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
func StripContext(msgid string, msgval string) string {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(msgid)
	inArgs[1].SetString(msgval)

	var ret gi.Argument

	err := stripContextFunction_Set()
	if err == nil {
		ret = stripContextFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
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
func Strjoinv(separator string, strArray string) string {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(separator)
	inArgs[1].SetString(strArray)

	var ret gi.Argument

	err := strjoinvFunction_Set()
	if err == nil {
		ret = strjoinvFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

var strlcatFunction *gi.Function
var strlcatFunction_Once sync.Once

func strlcatFunction_Set() error {
	var err error
	strlcatFunction_Once.Do(func() {
		strlcatFunction, err = gi.FunctionInvokerNew("GLib", "strlcat")
	})
	return err
}

// Strlcat is a representation of the C type g_strlcat.
func Strlcat(dest string, src string, destSize uint64) uint64 {
	var inArgs [3]gi.Argument
	inArgs[0].SetString(dest)
	inArgs[1].SetString(src)
	inArgs[2].SetUint64(destSize)

	var ret gi.Argument

	err := strlcatFunction_Set()
	if err == nil {
		ret = strlcatFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint64()

	return retGo
}

var strlcpyFunction *gi.Function
var strlcpyFunction_Once sync.Once

func strlcpyFunction_Set() error {
	var err error
	strlcpyFunction_Once.Do(func() {
		strlcpyFunction, err = gi.FunctionInvokerNew("GLib", "strlcpy")
	})
	return err
}

// Strlcpy is a representation of the C type g_strlcpy.
func Strlcpy(dest string, src string, destSize uint64) uint64 {
	var inArgs [3]gi.Argument
	inArgs[0].SetString(dest)
	inArgs[1].SetString(src)
	inArgs[2].SetUint64(destSize)

	var ret gi.Argument

	err := strlcpyFunction_Set()
	if err == nil {
		ret = strlcpyFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint64()

	return retGo
}

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
func Strncasecmp(s1 string, s2 string, n uint32) int32 {
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

	return retGo
}

var strndupFunction *gi.Function
var strndupFunction_Once sync.Once

func strndupFunction_Set() error {
	var err error
	strndupFunction_Once.Do(func() {
		strndupFunction, err = gi.FunctionInvokerNew("GLib", "strndup")
	})
	return err
}

// Strndup is a representation of the C type g_strndup.
func Strndup(str string, n uint64) string {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(str)
	inArgs[1].SetUint64(n)

	var ret gi.Argument

	err := strndupFunction_Set()
	if err == nil {
		ret = strndupFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

var strnfillFunction *gi.Function
var strnfillFunction_Once sync.Once

func strnfillFunction_Set() error {
	var err error
	strnfillFunction_Once.Do(func() {
		strnfillFunction, err = gi.FunctionInvokerNew("GLib", "strnfill")
	})
	return err
}

// Strnfill is a representation of the C type g_strnfill.
func Strnfill(length uint64, fillChar int8) string {
	var inArgs [2]gi.Argument
	inArgs[0].SetUint64(length)
	inArgs[1].SetInt8(fillChar)

	var ret gi.Argument

	err := strnfillFunction_Set()
	if err == nil {
		ret = strnfillFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

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
func Strreverse(string_ string) string {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(string_)

	var ret gi.Argument

	err := strreverseFunction_Set()
	if err == nil {
		ret = strreverseFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
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
func Strrstr(haystack string, needle string) string {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(haystack)
	inArgs[1].SetString(needle)

	var ret gi.Argument

	err := strrstrFunction_Set()
	if err == nil {
		ret = strrstrFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
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
func StrrstrLen(haystack string, haystackLen int32, needle string) string {
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

	return retGo
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
func Strsignal(signum int32) string {
	var inArgs [1]gi.Argument
	inArgs[0].SetInt32(signum)

	var ret gi.Argument

	err := strsignalFunction_Set()
	if err == nil {
		ret = strsignalFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
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
func Strsplit(string_ string, delimiter string, maxTokens int32) {
	var inArgs [3]gi.Argument
	inArgs[0].SetString(string_)
	inArgs[1].SetString(delimiter)
	inArgs[2].SetInt32(maxTokens)

	err := strsplitFunction_Set()
	if err == nil {
		strsplitFunction.Invoke(inArgs[:], nil)
	}

	return
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
func StrsplitSet(string_ string, delimiters string, maxTokens int32) {
	var inArgs [3]gi.Argument
	inArgs[0].SetString(string_)
	inArgs[1].SetString(delimiters)
	inArgs[2].SetInt32(maxTokens)

	err := strsplitSetFunction_Set()
	if err == nil {
		strsplitSetFunction.Invoke(inArgs[:], nil)
	}

	return
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
func StrstrLen(haystack string, haystackLen int32, needle string) string {
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

	return retGo
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
func Strtod(nptr string) (float64, string) {
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

	return retGo, out0
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
func Strup(string_ string) string {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(string_)

	var ret gi.Argument

	err := strupFunction_Set()
	if err == nil {
		ret = strupFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
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
func StrvContains(strv string, str string) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(strv)
	inArgs[1].SetString(str)

	var ret gi.Argument

	err := strvContainsFunction_Set()
	if err == nil {
		ret = strvContainsFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
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
func StrvEqual(strv1 string, strv2 string) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(strv1)
	inArgs[1].SetString(strv2)

	var ret gi.Argument

	err := strvEqualFunction_Set()
	if err == nil {
		ret = strvEqualFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
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
func StrvLength(strArray string) uint32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(strArray)

	var ret gi.Argument

	err := strvLengthFunction_Set()
	if err == nil {
		ret = strvLengthFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo
}

// UNSUPPORTED : C value 'g_test_add_data_func' : parameter 'test_data' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_test_add_data_func_full' : parameter 'test_data' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_test_add_func' : parameter 'test_func' of type 'TestFunc' not supported

// UNSUPPORTED : C value 'g_test_add_vtable' : parameter 'test_data' of type 'gpointer' not supported

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
func TestAssertExpectedMessagesInternal(domain string, file string, line int32, func_ string) {
	var inArgs [4]gi.Argument
	inArgs[0].SetString(domain)
	inArgs[1].SetString(file)
	inArgs[2].SetInt32(line)
	inArgs[3].SetString(func_)

	err := testAssertExpectedMessagesInternalFunction_Set()
	if err == nil {
		testAssertExpectedMessagesInternalFunction.Invoke(inArgs[:], nil)
	}

	return
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
func TestBug(bugUriSnippet string) {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(bugUriSnippet)

	err := testBugFunction_Set()
	if err == nil {
		testBugFunction.Invoke(inArgs[:], nil)
	}

	return
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
func TestBugBase(uriPattern string) {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(uriPattern)

	err := testBugBaseFunction_Set()
	if err == nil {
		testBugBaseFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'g_test_build_filename' : parameter 'file_type' of type 'TestFileType' not supported

// UNSUPPORTED : C value 'g_test_create_case' : parameter 'test_data' of type 'gpointer' not supported

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
func TestCreateSuite(suiteName string) *TestSuite {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(suiteName)

	var ret gi.Argument

	err := testCreateSuiteFunction_Set()
	if err == nil {
		ret = testCreateSuiteFunction.Invoke(inArgs[:], nil)
	}

	retGo := &TestSuite{native: ret.Pointer()}

	return retGo
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
func TestFail() {

	err := testFailFunction_Set()
	if err == nil {
		testFailFunction.Invoke(nil, nil)
	}

	return
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
func TestFailed() bool {

	var ret gi.Argument

	err := testFailedFunction_Set()
	if err == nil {
		ret = testFailedFunction.Invoke(nil, nil)
	}

	retGo := ret.Boolean()

	return retGo
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
func TestGetRoot() *TestSuite {

	var ret gi.Argument

	err := testGetRootFunction_Set()
	if err == nil {
		ret = testGetRootFunction.Invoke(nil, nil)
	}

	retGo := &TestSuite{native: ret.Pointer()}

	return retGo
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
func TestIncomplete(msg string) {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(msg)

	err := testIncompleteFunction_Set()
	if err == nil {
		testIncompleteFunction.Invoke(inArgs[:], nil)
	}

	return
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
func TestRandDouble() float64 {

	var ret gi.Argument

	err := testRandDoubleFunction_Set()
	if err == nil {
		ret = testRandDoubleFunction.Invoke(nil, nil)
	}

	retGo := ret.Float64()

	return retGo
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
func TestRandDoubleRange(rangeStart float64, rangeEnd float64) float64 {
	var inArgs [2]gi.Argument
	inArgs[0].SetFloat64(rangeStart)
	inArgs[1].SetFloat64(rangeEnd)

	var ret gi.Argument

	err := testRandDoubleRangeFunction_Set()
	if err == nil {
		ret = testRandDoubleRangeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Float64()

	return retGo
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
func TestRandInt() int32 {

	var ret gi.Argument

	err := testRandIntFunction_Set()
	if err == nil {
		ret = testRandIntFunction.Invoke(nil, nil)
	}

	retGo := ret.Int32()

	return retGo
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
func TestRandIntRange(begin int32, end int32) int32 {
	var inArgs [2]gi.Argument
	inArgs[0].SetInt32(begin)
	inArgs[1].SetInt32(end)

	var ret gi.Argument

	err := testRandIntRangeFunction_Set()
	if err == nil {
		ret = testRandIntRangeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
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
func TestRun() int32 {

	var ret gi.Argument

	err := testRunFunction_Set()
	if err == nil {
		ret = testRunFunction.Invoke(nil, nil)
	}

	retGo := ret.Int32()

	return retGo
}

var testRunSuiteFunction *gi.Function
var testRunSuiteFunction_Once sync.Once

func testRunSuiteFunction_Set() error {
	var err error
	testRunSuiteFunction_Once.Do(func() {
		testRunSuiteFunction, err = gi.FunctionInvokerNew("GLib", "test_run_suite")
	})
	return err
}

// TestRunSuite is a representation of the C type g_test_run_suite.
func TestRunSuite(suite *TestSuite) int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(suite.native)

	var ret gi.Argument

	err := testRunSuiteFunction_Set()
	if err == nil {
		ret = testRunSuiteFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

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
func TestSetNonfatalAssertions() {

	err := testSetNonfatalAssertionsFunction_Set()
	if err == nil {
		testSetNonfatalAssertionsFunction.Invoke(nil, nil)
	}

	return
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
func TestSkip(msg string) {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(msg)

	err := testSkipFunction_Set()
	if err == nil {
		testSkipFunction.Invoke(inArgs[:], nil)
	}

	return
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
func TestSubprocess() bool {

	var ret gi.Argument

	err := testSubprocessFunction_Set()
	if err == nil {
		ret = testSubprocessFunction.Invoke(nil, nil)
	}

	retGo := ret.Boolean()

	return retGo
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
func TestSummary(summary string) {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(summary)

	err := testSummaryFunction_Set()
	if err == nil {
		testSummaryFunction.Invoke(inArgs[:], nil)
	}

	return
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
func TestTimerElapsed() float64 {

	var ret gi.Argument

	err := testTimerElapsedFunction_Set()
	if err == nil {
		ret = testTimerElapsedFunction.Invoke(nil, nil)
	}

	retGo := ret.Float64()

	return retGo
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
func TestTimerLast() float64 {

	var ret gi.Argument

	err := testTimerLastFunction_Set()
	if err == nil {
		ret = testTimerLastFunction.Invoke(nil, nil)
	}

	retGo := ret.Float64()

	return retGo
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
func TestTimerStart() {

	err := testTimerStartFunction_Set()
	if err == nil {
		testTimerStartFunction.Invoke(nil, nil)
	}

	return
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
func TestTrapAssertions(domain string, file string, line int32, func_ string, assertionFlags uint64, pattern string) {
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

	return
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
func TestTrapHasPassed() bool {

	var ret gi.Argument

	err := testTrapHasPassedFunction_Set()
	if err == nil {
		ret = testTrapHasPassedFunction.Invoke(nil, nil)
	}

	retGo := ret.Boolean()

	return retGo
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
func TestTrapReachedTimeout() bool {

	var ret gi.Argument

	err := testTrapReachedTimeoutFunction_Set()
	if err == nil {
		ret = testTrapReachedTimeoutFunction.Invoke(nil, nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'g_test_trap_subprocess' : parameter 'test_flags' of type 'TestSubprocessFlags' not supported

var threadErrorQuarkFunction *gi.Function
var threadErrorQuarkFunction_Once sync.Once

func threadErrorQuarkFunction_Set() error {
	var err error
	threadErrorQuarkFunction_Once.Do(func() {
		threadErrorQuarkFunction, err = gi.FunctionInvokerNew("GLib", "thread_error_quark")
	})
	return err
}

// ThreadErrorQuark is a representation of the C type g_thread_error_quark.
func ThreadErrorQuark() Quark {

	var ret gi.Argument

	err := threadErrorQuarkFunction_Set()
	if err == nil {
		ret = threadErrorQuarkFunction.Invoke(nil, nil)
	}

	retGo := Quark(ret.Uint32())

	return retGo
}

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
func ThreadPoolGetMaxIdleTime() uint32 {

	var ret gi.Argument

	err := threadPoolGetMaxIdleTimeFunction_Set()
	if err == nil {
		ret = threadPoolGetMaxIdleTimeFunction.Invoke(nil, nil)
	}

	retGo := ret.Uint32()

	return retGo
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
func ThreadPoolGetMaxUnusedThreads() int32 {

	var ret gi.Argument

	err := threadPoolGetMaxUnusedThreadsFunction_Set()
	if err == nil {
		ret = threadPoolGetMaxUnusedThreadsFunction.Invoke(nil, nil)
	}

	retGo := ret.Int32()

	return retGo
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
func ThreadPoolGetNumUnusedThreads() uint32 {

	var ret gi.Argument

	err := threadPoolGetNumUnusedThreadsFunction_Set()
	if err == nil {
		ret = threadPoolGetNumUnusedThreadsFunction.Invoke(nil, nil)
	}

	retGo := ret.Uint32()

	return retGo
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
func ThreadPoolSetMaxIdleTime(interval uint32) {
	var inArgs [1]gi.Argument
	inArgs[0].SetUint32(interval)

	err := threadPoolSetMaxIdleTimeFunction_Set()
	if err == nil {
		threadPoolSetMaxIdleTimeFunction.Invoke(inArgs[:], nil)
	}

	return
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
func ThreadPoolSetMaxUnusedThreads(maxThreads int32) {
	var inArgs [1]gi.Argument
	inArgs[0].SetInt32(maxThreads)

	err := threadPoolSetMaxUnusedThreadsFunction_Set()
	if err == nil {
		threadPoolSetMaxUnusedThreadsFunction.Invoke(inArgs[:], nil)
	}

	return
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
func ThreadPoolStopUnusedThreads() {

	err := threadPoolStopUnusedThreadsFunction_Set()
	if err == nil {
		threadPoolStopUnusedThreadsFunction.Invoke(nil, nil)
	}

	return
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
func ThreadSelf() *Thread {

	var ret gi.Argument

	err := threadSelfFunction_Set()
	if err == nil {
		ret = threadSelfFunction.Invoke(nil, nil)
	}

	retGo := &Thread{native: ret.Pointer()}

	return retGo
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
func ThreadYield() {

	err := threadYieldFunction_Set()
	if err == nil {
		threadYieldFunction.Invoke(nil, nil)
	}

	return
}

var timeValFromIso8601Function *gi.Function
var timeValFromIso8601Function_Once sync.Once

func timeValFromIso8601Function_Set() error {
	var err error
	timeValFromIso8601Function_Once.Do(func() {
		timeValFromIso8601Function, err = gi.FunctionInvokerNew("GLib", "time_val_from_iso8601")
	})
	return err
}

// TimeValFromIso8601 is a representation of the C type g_time_val_from_iso8601.
func TimeValFromIso8601(isoDate string) (bool, *TimeVal) {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(isoDate)

	var outArgs [1]gi.Argument
	var ret gi.Argument

	err := timeValFromIso8601Function_Set()
	if err == nil {
		ret = timeValFromIso8601Function.Invoke(inArgs[:], outArgs[:])
	}

	retGo := ret.Boolean()
	out0 := &TimeVal{native: outArgs[0].Pointer()}

	return retGo, out0
}

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
func TimeoutSourceNew(interval uint32) *Source {
	var inArgs [1]gi.Argument
	inArgs[0].SetUint32(interval)

	var ret gi.Argument

	err := timeoutSourceNewFunction_Set()
	if err == nil {
		ret = timeoutSourceNewFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Source{native: ret.Pointer()}

	return retGo
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
func TimeoutSourceNewSeconds(interval uint32) *Source {
	var inArgs [1]gi.Argument
	inArgs[0].SetUint32(interval)

	var ret gi.Argument

	err := timeoutSourceNewSecondsFunction_Set()
	if err == nil {
		ret = timeoutSourceNewSecondsFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Source{native: ret.Pointer()}

	return retGo
}

var trashStackHeightFunction *gi.Function
var trashStackHeightFunction_Once sync.Once

func trashStackHeightFunction_Set() error {
	var err error
	trashStackHeightFunction_Once.Do(func() {
		trashStackHeightFunction, err = gi.FunctionInvokerNew("GLib", "trash_stack_height")
	})
	return err
}

// TrashStackHeight is a representation of the C type g_trash_stack_height.
func TrashStackHeight(stackP *TrashStack) uint32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(stackP.native)

	var ret gi.Argument

	err := trashStackHeightFunction_Set()
	if err == nil {
		ret = trashStackHeightFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo
}

// UNSUPPORTED : C value 'g_trash_stack_peek' : return type 'gpointer' not supported

// UNSUPPORTED : C value 'g_trash_stack_pop' : return type 'gpointer' not supported

// UNSUPPORTED : C value 'g_trash_stack_push' : parameter 'data_p' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_try_malloc' : return type 'gpointer' not supported

// UNSUPPORTED : C value 'g_try_malloc0' : return type 'gpointer' not supported

// UNSUPPORTED : C value 'g_try_malloc0_n' : return type 'gpointer' not supported

// UNSUPPORTED : C value 'g_try_malloc_n' : return type 'gpointer' not supported

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

var unixErrorQuarkFunction *gi.Function
var unixErrorQuarkFunction_Once sync.Once

func unixErrorQuarkFunction_Set() error {
	var err error
	unixErrorQuarkFunction_Once.Do(func() {
		unixErrorQuarkFunction, err = gi.FunctionInvokerNew("GLib", "unix_error_quark")
	})
	return err
}

// UnixErrorQuark is a representation of the C type g_unix_error_quark.
func UnixErrorQuark() Quark {

	var ret gi.Argument

	err := unixErrorQuarkFunction_Set()
	if err == nil {
		ret = unixErrorQuarkFunction.Invoke(nil, nil)
	}

	retGo := Quark(ret.Uint32())

	return retGo
}

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
func UnixOpenPipe(fds int32, flags int32) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetInt32(fds)
	inArgs[1].SetInt32(flags)

	var ret gi.Argument

	err := unixOpenPipeFunction_Set()
	if err == nil {
		ret = unixOpenPipeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
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
func UnixSetFdNonblocking(fd int32, nonblock bool) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetInt32(fd)
	inArgs[1].SetBoolean(nonblock)

	var ret gi.Argument

	err := unixSetFdNonblockingFunction_Set()
	if err == nil {
		ret = unixSetFdNonblockingFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
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
func UnixSignalSourceNew(signum int32) *Source {
	var inArgs [1]gi.Argument
	inArgs[0].SetInt32(signum)

	var ret gi.Argument

	err := unixSignalSourceNewFunction_Set()
	if err == nil {
		ret = unixSignalSourceNewFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Source{native: ret.Pointer()}

	return retGo
}

var unlinkFunction *gi.Function
var unlinkFunction_Once sync.Once

func unlinkFunction_Set() error {
	var err error
	unlinkFunction_Once.Do(func() {
		unlinkFunction, err = gi.FunctionInvokerNew("GLib", "unlink")
	})
	return err
}

// Unlink is a representation of the C type g_unlink.
func Unlink(filename string) int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(filename)

	var ret gi.Argument

	err := unlinkFunction_Set()
	if err == nil {
		ret = unlinkFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var unsetenvFunction *gi.Function
var unsetenvFunction_Once sync.Once

func unsetenvFunction_Set() error {
	var err error
	unsetenvFunction_Once.Do(func() {
		unsetenvFunction, err = gi.FunctionInvokerNew("GLib", "unsetenv")
	})
	return err
}

// Unsetenv is a representation of the C type g_unsetenv.
func Unsetenv(variable string) {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(variable)

	err := unsetenvFunction_Set()
	if err == nil {
		unsetenvFunction.Invoke(inArgs[:], nil)
	}

	return
}

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
func UriEscapeString(unescaped string, reservedCharsAllowed string, allowUtf8 bool) string {
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

	return retGo
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
func UriListExtractUris(uriList string) {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(uriList)

	err := uriListExtractUrisFunction_Set()
	if err == nil {
		uriListExtractUrisFunction.Invoke(inArgs[:], nil)
	}

	return
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
func UriParseScheme(uri string) string {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(uri)

	var ret gi.Argument

	err := uriParseSchemeFunction_Set()
	if err == nil {
		ret = uriParseSchemeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
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
func UriUnescapeSegment(escapedString string, escapedStringEnd string, illegalCharacters string) string {
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

	return retGo
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
func UriUnescapeString(escapedString string, illegalCharacters string) string {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(escapedString)
	inArgs[1].SetString(illegalCharacters)

	var ret gi.Argument

	err := uriUnescapeStringFunction_Set()
	if err == nil {
		ret = uriUnescapeStringFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
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
func Usleep(microseconds uint64) {
	var inArgs [1]gi.Argument
	inArgs[0].SetUint64(microseconds)

	err := usleepFunction_Set()
	if err == nil {
		usleepFunction.Invoke(inArgs[:], nil)
	}

	return
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
func Utf16ToUtf8(str uint16, len int64) (string, int64, int64) {
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

	return retGo, out0, out1
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
func Utf8Casefold(str string, len int32) string {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(str)
	inArgs[1].SetInt32(len)

	var ret gi.Argument

	err := utf8CasefoldFunction_Set()
	if err == nil {
		ret = utf8CasefoldFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
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
func Utf8Collate(str1 string, str2 string) int32 {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(str1)
	inArgs[1].SetString(str2)

	var ret gi.Argument

	err := utf8CollateFunction_Set()
	if err == nil {
		ret = utf8CollateFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
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
func Utf8CollateKey(str string, len int32) string {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(str)
	inArgs[1].SetInt32(len)

	var ret gi.Argument

	err := utf8CollateKeyFunction_Set()
	if err == nil {
		ret = utf8CollateKeyFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
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
func Utf8CollateKeyForFilename(str string, len int32) string {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(str)
	inArgs[1].SetInt32(len)

	var ret gi.Argument

	err := utf8CollateKeyForFilenameFunction_Set()
	if err == nil {
		ret = utf8CollateKeyForFilenameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
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
func Utf8FindNextChar(p string, end string) string {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(p)
	inArgs[1].SetString(end)

	var ret gi.Argument

	err := utf8FindNextCharFunction_Set()
	if err == nil {
		ret = utf8FindNextCharFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
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
func Utf8FindPrevChar(str string, p string) string {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(str)
	inArgs[1].SetString(p)

	var ret gi.Argument

	err := utf8FindPrevCharFunction_Set()
	if err == nil {
		ret = utf8FindPrevCharFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
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
func Utf8MakeValid(str string, len int32) string {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(str)
	inArgs[1].SetInt32(len)

	var ret gi.Argument

	err := utf8MakeValidFunction_Set()
	if err == nil {
		ret = utf8MakeValidFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
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
func Utf8OffsetToPointer(str string, offset int64) string {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(str)
	inArgs[1].SetInt64(offset)

	var ret gi.Argument

	err := utf8OffsetToPointerFunction_Set()
	if err == nil {
		ret = utf8OffsetToPointerFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
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
func Utf8PointerToOffset(str string, pos string) int64 {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(str)
	inArgs[1].SetString(pos)

	var ret gi.Argument

	err := utf8PointerToOffsetFunction_Set()
	if err == nil {
		ret = utf8PointerToOffsetFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int64()

	return retGo
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
func Utf8PrevChar(p string) string {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(p)

	var ret gi.Argument

	err := utf8PrevCharFunction_Set()
	if err == nil {
		ret = utf8PrevCharFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
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
func Utf8Strdown(str string, len int32) string {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(str)
	inArgs[1].SetInt32(len)

	var ret gi.Argument

	err := utf8StrdownFunction_Set()
	if err == nil {
		ret = utf8StrdownFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
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
func Utf8Strlen(p string, max int32) int64 {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(p)
	inArgs[1].SetInt32(max)

	var ret gi.Argument

	err := utf8StrlenFunction_Set()
	if err == nil {
		ret = utf8StrlenFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int64()

	return retGo
}

var utf8StrncpyFunction *gi.Function
var utf8StrncpyFunction_Once sync.Once

func utf8StrncpyFunction_Set() error {
	var err error
	utf8StrncpyFunction_Once.Do(func() {
		utf8StrncpyFunction, err = gi.FunctionInvokerNew("GLib", "utf8_strncpy")
	})
	return err
}

// Utf8Strncpy is a representation of the C type g_utf8_strncpy.
func Utf8Strncpy(dest string, src string, n uint64) string {
	var inArgs [3]gi.Argument
	inArgs[0].SetString(dest)
	inArgs[1].SetString(src)
	inArgs[2].SetUint64(n)

	var ret gi.Argument

	err := utf8StrncpyFunction_Set()
	if err == nil {
		ret = utf8StrncpyFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

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
func Utf8Strreverse(str string, len int32) string {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(str)
	inArgs[1].SetInt32(len)

	var ret gi.Argument

	err := utf8StrreverseFunction_Set()
	if err == nil {
		ret = utf8StrreverseFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
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
func Utf8Strup(str string, len int32) string {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(str)
	inArgs[1].SetInt32(len)

	var ret gi.Argument

	err := utf8StrupFunction_Set()
	if err == nil {
		ret = utf8StrupFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
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
func Utf8Substring(str string, startPos int64, endPos int64) string {
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

	return retGo
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
func Utf8ToUtf16(str string, len int64) (uint16, int64, int64) {
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

	return retGo, out0, out1
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
func UuidStringIsValid(str string) bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(str)

	var ret gi.Argument

	err := uuidStringIsValidFunction_Set()
	if err == nil {
		ret = uuidStringIsValidFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
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
func UuidStringRandom() string {

	var ret gi.Argument

	err := uuidStringRandomFunction_Set()
	if err == nil {
		ret = uuidStringRandomFunction.Invoke(nil, nil)
	}

	retGo := ret.String(true)

	return retGo
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
func VariantIsObjectPath(string_ string) bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(string_)

	var ret gi.Argument

	err := variantIsObjectPathFunction_Set()
	if err == nil {
		ret = variantIsObjectPathFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
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
func VariantIsSignature(string_ string) bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(string_)

	var ret gi.Argument

	err := variantIsSignatureFunction_Set()
	if err == nil {
		ret = variantIsSignatureFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var variantParseFunction *gi.Function
var variantParseFunction_Once sync.Once

func variantParseFunction_Set() error {
	var err error
	variantParseFunction_Once.Do(func() {
		variantParseFunction, err = gi.FunctionInvokerNew("GLib", "variant_parse")
	})
	return err
}

// VariantParse is a representation of the C type g_variant_parse.
func VariantParse(type_ *VariantType, text string, limit string, endptr string) *Variant {
	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(type_.native)
	inArgs[1].SetString(text)
	inArgs[2].SetString(limit)
	inArgs[3].SetString(endptr)

	var ret gi.Argument

	err := variantParseFunction_Set()
	if err == nil {
		ret = variantParseFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Variant{native: ret.Pointer()}

	return retGo
}

var variantParseErrorPrintContextFunction *gi.Function
var variantParseErrorPrintContextFunction_Once sync.Once

func variantParseErrorPrintContextFunction_Set() error {
	var err error
	variantParseErrorPrintContextFunction_Once.Do(func() {
		variantParseErrorPrintContextFunction, err = gi.FunctionInvokerNew("GLib", "variant_parse_error_print_context")
	})
	return err
}

// VariantParseErrorPrintContext is a representation of the C type g_variant_parse_error_print_context.
func VariantParseErrorPrintContext(error *Error, sourceStr string) string {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(error.native)
	inArgs[1].SetString(sourceStr)

	var ret gi.Argument

	err := variantParseErrorPrintContextFunction_Set()
	if err == nil {
		ret = variantParseErrorPrintContextFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

var variantParseErrorQuarkFunction *gi.Function
var variantParseErrorQuarkFunction_Once sync.Once

func variantParseErrorQuarkFunction_Set() error {
	var err error
	variantParseErrorQuarkFunction_Once.Do(func() {
		variantParseErrorQuarkFunction, err = gi.FunctionInvokerNew("GLib", "variant_parse_error_quark")
	})
	return err
}

// VariantParseErrorQuark is a representation of the C type g_variant_parse_error_quark.
func VariantParseErrorQuark() Quark {

	var ret gi.Argument

	err := variantParseErrorQuarkFunction_Set()
	if err == nil {
		ret = variantParseErrorQuarkFunction.Invoke(nil, nil)
	}

	retGo := Quark(ret.Uint32())

	return retGo
}

var variantParserGetErrorQuarkFunction *gi.Function
var variantParserGetErrorQuarkFunction_Once sync.Once

func variantParserGetErrorQuarkFunction_Set() error {
	var err error
	variantParserGetErrorQuarkFunction_Once.Do(func() {
		variantParserGetErrorQuarkFunction, err = gi.FunctionInvokerNew("GLib", "variant_parser_get_error_quark")
	})
	return err
}

// VariantParserGetErrorQuark is a representation of the C type g_variant_parser_get_error_quark.
func VariantParserGetErrorQuark() Quark {

	var ret gi.Argument

	err := variantParserGetErrorQuarkFunction_Set()
	if err == nil {
		ret = variantParserGetErrorQuarkFunction.Invoke(nil, nil)
	}

	retGo := Quark(ret.Uint32())

	return retGo
}

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
func VariantTypeChecked(arg0 string) *VariantType {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(arg0)

	var ret gi.Argument

	err := variantTypeCheckedFunction_Set()
	if err == nil {
		ret = variantTypeCheckedFunction.Invoke(inArgs[:], nil)
	}

	retGo := &VariantType{native: ret.Pointer()}

	return retGo
}

var variantTypeStringGetDepthFunction *gi.Function
var variantTypeStringGetDepthFunction_Once sync.Once

func variantTypeStringGetDepthFunction_Set() error {
	var err error
	variantTypeStringGetDepthFunction_Once.Do(func() {
		variantTypeStringGetDepthFunction, err = gi.FunctionInvokerNew("GLib", "variant_type_string_get_depth_")
	})
	return err
}

// VariantTypeStringGetDepth is a representation of the C type g_variant_type_string_get_depth_.
func VariantTypeStringGetDepth(typeString string) uint64 {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(typeString)

	var ret gi.Argument

	err := variantTypeStringGetDepthFunction_Set()
	if err == nil {
		ret = variantTypeStringGetDepthFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint64()

	return retGo
}

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
func VariantTypeStringIsValid(typeString string) bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(typeString)

	var ret gi.Argument

	err := variantTypeStringIsValidFunction_Set()
	if err == nil {
		ret = variantTypeStringIsValidFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
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
func VariantTypeStringScan(string_ string, limit string) (bool, string) {
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

	return retGo, out0
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
func WarnMessage(domain string, file string, line int32, func_ string, warnexpr string) {
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

	return
}
