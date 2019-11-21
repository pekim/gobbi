// Code generated - DO NOT EDIT.

package glib

import (
	gi "github.com/pekim/gobbi/internal/gi"
	"sync"
)

// UNSUPPORTED : C value 'g_access' : parameter 'filename' of type 'filename' not supported

var AsciiDigitValueFunction *gi.Function
var AsciiDigitValueFunctionOnce sync.Once

func AsciiDigitValueFunctionSet() {
	AsciiDigitValueFunctionOnce.Do(func() {
		AsciiDigitValueFunction = gi.FunctionInvokerNew("GLib", "ascii_digit_value")
	})
}

var asciiDigitValueInvoker *gi.Function

// AsciiDigitValue is a representation of the C type g_ascii_digit_value.
func AsciiDigitValue(c int8) int32 {
	if asciiDigitValueInvoker == nil {
		asciiDigitValueInvoker = gi.FunctionInvokerNew("GLib", "ascii_digit_value")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetInt8(c)

	ret := asciiDigitValueInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

// UNSUPPORTED : C value 'g_ascii_dtostr' : parameter 'd' of type 'gdouble' not supported

// UNSUPPORTED : C value 'g_ascii_formatd' : parameter 'd' of type 'gdouble' not supported

var AsciiStrcasecmpFunction *gi.Function
var AsciiStrcasecmpFunctionOnce sync.Once

func AsciiStrcasecmpFunctionSet() {
	AsciiStrcasecmpFunctionOnce.Do(func() {
		AsciiStrcasecmpFunction = gi.FunctionInvokerNew("GLib", "ascii_strcasecmp")
	})
}

var asciiStrcasecmpInvoker *gi.Function

// AsciiStrcasecmp is a representation of the C type g_ascii_strcasecmp.
func AsciiStrcasecmp(s1 string, s2 string) int32 {
	if asciiStrcasecmpInvoker == nil {
		asciiStrcasecmpInvoker = gi.FunctionInvokerNew("GLib", "ascii_strcasecmp")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetString(s1)
	inArgs[1].SetString(s2)

	ret := asciiStrcasecmpInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var AsciiStrdownFunction *gi.Function
var AsciiStrdownFunctionOnce sync.Once

func AsciiStrdownFunctionSet() {
	AsciiStrdownFunctionOnce.Do(func() {
		AsciiStrdownFunction = gi.FunctionInvokerNew("GLib", "ascii_strdown")
	})
}

var asciiStrdownInvoker *gi.Function

// AsciiStrdown is a representation of the C type g_ascii_strdown.
func AsciiStrdown(str string, len int32) string {
	if asciiStrdownInvoker == nil {
		asciiStrdownInvoker = gi.FunctionInvokerNew("GLib", "ascii_strdown")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetString(str)
	inArgs[1].SetInt32(len)

	ret := asciiStrdownInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

// UNSUPPORTED : C value 'g_ascii_string_to_signed' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_ascii_string_to_unsigned' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_ascii_strncasecmp' : parameter 'n' of type 'gsize' not supported

// UNSUPPORTED : C value 'g_ascii_strtod' : return type 'gdouble' not supported

var AsciiStrtollFunction *gi.Function
var AsciiStrtollFunctionOnce sync.Once

func AsciiStrtollFunctionSet() {
	AsciiStrtollFunctionOnce.Do(func() {
		AsciiStrtollFunction = gi.FunctionInvokerNew("GLib", "ascii_strtoll")
	})
}

var asciiStrtollInvoker *gi.Function

// AsciiStrtoll is a representation of the C type g_ascii_strtoll.
func AsciiStrtoll(nptr string, base uint32) (int64, string) {
	if asciiStrtollInvoker == nil {
		asciiStrtollInvoker = gi.FunctionInvokerNew("GLib", "ascii_strtoll")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetString(nptr)
	inArgs[1].SetUint32(base)

	var outArgs [1]gi.Argument

	ret := asciiStrtollInvoker.Invoke(inArgs[:], outArgs[:])

	retGo := ret.Int64()
	out0 := outArgs[0].String(false)

	return retGo, out0
}

var AsciiStrtoullFunction *gi.Function
var AsciiStrtoullFunctionOnce sync.Once

func AsciiStrtoullFunctionSet() {
	AsciiStrtoullFunctionOnce.Do(func() {
		AsciiStrtoullFunction = gi.FunctionInvokerNew("GLib", "ascii_strtoull")
	})
}

var asciiStrtoullInvoker *gi.Function

// AsciiStrtoull is a representation of the C type g_ascii_strtoull.
func AsciiStrtoull(nptr string, base uint32) (uint64, string) {
	if asciiStrtoullInvoker == nil {
		asciiStrtoullInvoker = gi.FunctionInvokerNew("GLib", "ascii_strtoull")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetString(nptr)
	inArgs[1].SetUint32(base)

	var outArgs [1]gi.Argument

	ret := asciiStrtoullInvoker.Invoke(inArgs[:], outArgs[:])

	retGo := ret.Uint64()
	out0 := outArgs[0].String(false)

	return retGo, out0
}

var AsciiStrupFunction *gi.Function
var AsciiStrupFunctionOnce sync.Once

func AsciiStrupFunctionSet() {
	AsciiStrupFunctionOnce.Do(func() {
		AsciiStrupFunction = gi.FunctionInvokerNew("GLib", "ascii_strup")
	})
}

var asciiStrupInvoker *gi.Function

// AsciiStrup is a representation of the C type g_ascii_strup.
func AsciiStrup(str string, len int32) string {
	if asciiStrupInvoker == nil {
		asciiStrupInvoker = gi.FunctionInvokerNew("GLib", "ascii_strup")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetString(str)
	inArgs[1].SetInt32(len)

	ret := asciiStrupInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

var AsciiTolowerFunction *gi.Function
var AsciiTolowerFunctionOnce sync.Once

func AsciiTolowerFunctionSet() {
	AsciiTolowerFunctionOnce.Do(func() {
		AsciiTolowerFunction = gi.FunctionInvokerNew("GLib", "ascii_tolower")
	})
}

var asciiTolowerInvoker *gi.Function

// AsciiTolower is a representation of the C type g_ascii_tolower.
func AsciiTolower(c int8) int8 {
	if asciiTolowerInvoker == nil {
		asciiTolowerInvoker = gi.FunctionInvokerNew("GLib", "ascii_tolower")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetInt8(c)

	ret := asciiTolowerInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int8()

	return retGo
}

var AsciiToupperFunction *gi.Function
var AsciiToupperFunctionOnce sync.Once

func AsciiToupperFunctionSet() {
	AsciiToupperFunctionOnce.Do(func() {
		AsciiToupperFunction = gi.FunctionInvokerNew("GLib", "ascii_toupper")
	})
}

var asciiToupperInvoker *gi.Function

// AsciiToupper is a representation of the C type g_ascii_toupper.
func AsciiToupper(c int8) int8 {
	if asciiToupperInvoker == nil {
		asciiToupperInvoker = gi.FunctionInvokerNew("GLib", "ascii_toupper")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetInt8(c)

	ret := asciiToupperInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int8()

	return retGo
}

var AsciiXdigitValueFunction *gi.Function
var AsciiXdigitValueFunctionOnce sync.Once

func AsciiXdigitValueFunctionSet() {
	AsciiXdigitValueFunctionOnce.Do(func() {
		AsciiXdigitValueFunction = gi.FunctionInvokerNew("GLib", "ascii_xdigit_value")
	})
}

var asciiXdigitValueInvoker *gi.Function

// AsciiXdigitValue is a representation of the C type g_ascii_xdigit_value.
func AsciiXdigitValue(c int8) int32 {
	if asciiXdigitValueInvoker == nil {
		asciiXdigitValueInvoker = gi.FunctionInvokerNew("GLib", "ascii_xdigit_value")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetInt8(c)

	ret := asciiXdigitValueInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var AssertWarningFunction *gi.Function
var AssertWarningFunctionOnce sync.Once

func AssertWarningFunctionSet() {
	AssertWarningFunctionOnce.Do(func() {
		AssertWarningFunction = gi.FunctionInvokerNew("GLib", "assert_warning")
	})
}

var assertWarningInvoker *gi.Function

// AssertWarning is a representation of the C type g_assert_warning.
func AssertWarning(logDomain string, file string, line int32, prettyFunction string, expression string) {
	if assertWarningInvoker == nil {
		assertWarningInvoker = gi.FunctionInvokerNew("GLib", "assert_warning")
	}

	var inArgs [5]gi.Argument
	inArgs[0].SetString(logDomain)
	inArgs[1].SetString(file)
	inArgs[2].SetInt32(line)
	inArgs[3].SetString(prettyFunction)
	inArgs[4].SetString(expression)

	assertWarningInvoker.Invoke(inArgs[:], nil)

}

var AssertionMessageFunction *gi.Function
var AssertionMessageFunctionOnce sync.Once

func AssertionMessageFunctionSet() {
	AssertionMessageFunctionOnce.Do(func() {
		AssertionMessageFunction = gi.FunctionInvokerNew("GLib", "assertion_message")
	})
}

var assertionMessageInvoker *gi.Function

// AssertionMessage is a representation of the C type g_assertion_message.
func AssertionMessage(domain string, file string, line int32, func_ string, message string) {
	if assertionMessageInvoker == nil {
		assertionMessageInvoker = gi.FunctionInvokerNew("GLib", "assertion_message")
	}

	var inArgs [5]gi.Argument
	inArgs[0].SetString(domain)
	inArgs[1].SetString(file)
	inArgs[2].SetInt32(line)
	inArgs[3].SetString(func_)
	inArgs[4].SetString(message)

	assertionMessageInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_assertion_message_cmpnum' : parameter 'arg1' of type 'long double' not supported

var AssertionMessageCmpstrFunction *gi.Function
var AssertionMessageCmpstrFunctionOnce sync.Once

func AssertionMessageCmpstrFunctionSet() {
	AssertionMessageCmpstrFunctionOnce.Do(func() {
		AssertionMessageCmpstrFunction = gi.FunctionInvokerNew("GLib", "assertion_message_cmpstr")
	})
}

var assertionMessageCmpstrInvoker *gi.Function

// AssertionMessageCmpstr is a representation of the C type g_assertion_message_cmpstr.
func AssertionMessageCmpstr(domain string, file string, line int32, func_ string, expr string, arg1 string, cmp string, arg2 string) {
	if assertionMessageCmpstrInvoker == nil {
		assertionMessageCmpstrInvoker = gi.FunctionInvokerNew("GLib", "assertion_message_cmpstr")
	}

	var inArgs [8]gi.Argument
	inArgs[0].SetString(domain)
	inArgs[1].SetString(file)
	inArgs[2].SetInt32(line)
	inArgs[3].SetString(func_)
	inArgs[4].SetString(expr)
	inArgs[5].SetString(arg1)
	inArgs[6].SetString(cmp)
	inArgs[7].SetString(arg2)

	assertionMessageCmpstrInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_assertion_message_error' : parameter 'error' of type 'Error' not supported

var AssertionMessageExprFunction *gi.Function
var AssertionMessageExprFunctionOnce sync.Once

func AssertionMessageExprFunctionSet() {
	AssertionMessageExprFunctionOnce.Do(func() {
		AssertionMessageExprFunction = gi.FunctionInvokerNew("GLib", "assertion_message_expr")
	})
}

var assertionMessageExprInvoker *gi.Function

// AssertionMessageExpr is a representation of the C type g_assertion_message_expr.
func AssertionMessageExpr(domain string, file string, line int32, func_ string, expr string) {
	if assertionMessageExprInvoker == nil {
		assertionMessageExprInvoker = gi.FunctionInvokerNew("GLib", "assertion_message_expr")
	}

	var inArgs [5]gi.Argument
	inArgs[0].SetString(domain)
	inArgs[1].SetString(file)
	inArgs[2].SetInt32(line)
	inArgs[3].SetString(func_)
	inArgs[4].SetString(expr)

	assertionMessageExprInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_atexit' : parameter 'func' of type 'VoidFunc' not supported

var AtomicIntAddFunction *gi.Function
var AtomicIntAddFunctionOnce sync.Once

func AtomicIntAddFunctionSet() {
	AtomicIntAddFunctionOnce.Do(func() {
		AtomicIntAddFunction = gi.FunctionInvokerNew("GLib", "atomic_int_add")
	})
}

var atomicIntAddInvoker *gi.Function

// AtomicIntAdd is a representation of the C type g_atomic_int_add.
func AtomicIntAdd(atomic int32, val int32) int32 {
	if atomicIntAddInvoker == nil {
		atomicIntAddInvoker = gi.FunctionInvokerNew("GLib", "atomic_int_add")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetInt32(atomic)
	inArgs[1].SetInt32(val)

	ret := atomicIntAddInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var AtomicIntAndFunction *gi.Function
var AtomicIntAndFunctionOnce sync.Once

func AtomicIntAndFunctionSet() {
	AtomicIntAndFunctionOnce.Do(func() {
		AtomicIntAndFunction = gi.FunctionInvokerNew("GLib", "atomic_int_and")
	})
}

var atomicIntAndInvoker *gi.Function

// AtomicIntAnd is a representation of the C type g_atomic_int_and.
func AtomicIntAnd(atomic uint32, val uint32) uint32 {
	if atomicIntAndInvoker == nil {
		atomicIntAndInvoker = gi.FunctionInvokerNew("GLib", "atomic_int_and")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetUint32(atomic)
	inArgs[1].SetUint32(val)

	ret := atomicIntAndInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Uint32()

	return retGo
}

// UNSUPPORTED : C value 'g_atomic_int_compare_and_exchange' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_atomic_int_dec_and_test' : return type 'gboolean' not supported

var AtomicIntExchangeAndAddFunction *gi.Function
var AtomicIntExchangeAndAddFunctionOnce sync.Once

func AtomicIntExchangeAndAddFunctionSet() {
	AtomicIntExchangeAndAddFunctionOnce.Do(func() {
		AtomicIntExchangeAndAddFunction = gi.FunctionInvokerNew("GLib", "atomic_int_exchange_and_add")
	})
}

var atomicIntExchangeAndAddInvoker *gi.Function

// AtomicIntExchangeAndAdd is a representation of the C type g_atomic_int_exchange_and_add.
func AtomicIntExchangeAndAdd(atomic int32, val int32) int32 {
	if atomicIntExchangeAndAddInvoker == nil {
		atomicIntExchangeAndAddInvoker = gi.FunctionInvokerNew("GLib", "atomic_int_exchange_and_add")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetInt32(atomic)
	inArgs[1].SetInt32(val)

	ret := atomicIntExchangeAndAddInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var AtomicIntGetFunction *gi.Function
var AtomicIntGetFunctionOnce sync.Once

func AtomicIntGetFunctionSet() {
	AtomicIntGetFunctionOnce.Do(func() {
		AtomicIntGetFunction = gi.FunctionInvokerNew("GLib", "atomic_int_get")
	})
}

var atomicIntGetInvoker *gi.Function

// AtomicIntGet is a representation of the C type g_atomic_int_get.
func AtomicIntGet(atomic int32) int32 {
	if atomicIntGetInvoker == nil {
		atomicIntGetInvoker = gi.FunctionInvokerNew("GLib", "atomic_int_get")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetInt32(atomic)

	ret := atomicIntGetInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var AtomicIntIncFunction *gi.Function
var AtomicIntIncFunctionOnce sync.Once

func AtomicIntIncFunctionSet() {
	AtomicIntIncFunctionOnce.Do(func() {
		AtomicIntIncFunction = gi.FunctionInvokerNew("GLib", "atomic_int_inc")
	})
}

var atomicIntIncInvoker *gi.Function

// AtomicIntInc is a representation of the C type g_atomic_int_inc.
func AtomicIntInc(atomic int32) {
	if atomicIntIncInvoker == nil {
		atomicIntIncInvoker = gi.FunctionInvokerNew("GLib", "atomic_int_inc")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetInt32(atomic)

	atomicIntIncInvoker.Invoke(inArgs[:], nil)

}

var AtomicIntOrFunction *gi.Function
var AtomicIntOrFunctionOnce sync.Once

func AtomicIntOrFunctionSet() {
	AtomicIntOrFunctionOnce.Do(func() {
		AtomicIntOrFunction = gi.FunctionInvokerNew("GLib", "atomic_int_or")
	})
}

var atomicIntOrInvoker *gi.Function

// AtomicIntOr is a representation of the C type g_atomic_int_or.
func AtomicIntOr(atomic uint32, val uint32) uint32 {
	if atomicIntOrInvoker == nil {
		atomicIntOrInvoker = gi.FunctionInvokerNew("GLib", "atomic_int_or")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetUint32(atomic)
	inArgs[1].SetUint32(val)

	ret := atomicIntOrInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Uint32()

	return retGo
}

var AtomicIntSetFunction *gi.Function
var AtomicIntSetFunctionOnce sync.Once

func AtomicIntSetFunctionSet() {
	AtomicIntSetFunctionOnce.Do(func() {
		AtomicIntSetFunction = gi.FunctionInvokerNew("GLib", "atomic_int_set")
	})
}

var atomicIntSetInvoker *gi.Function

// AtomicIntSet is a representation of the C type g_atomic_int_set.
func AtomicIntSet(atomic int32, newval int32) {
	if atomicIntSetInvoker == nil {
		atomicIntSetInvoker = gi.FunctionInvokerNew("GLib", "atomic_int_set")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetInt32(atomic)
	inArgs[1].SetInt32(newval)

	atomicIntSetInvoker.Invoke(inArgs[:], nil)

}

var AtomicIntXorFunction *gi.Function
var AtomicIntXorFunctionOnce sync.Once

func AtomicIntXorFunctionSet() {
	AtomicIntXorFunctionOnce.Do(func() {
		AtomicIntXorFunction = gi.FunctionInvokerNew("GLib", "atomic_int_xor")
	})
}

var atomicIntXorInvoker *gi.Function

// AtomicIntXor is a representation of the C type g_atomic_int_xor.
func AtomicIntXor(atomic uint32, val uint32) uint32 {
	if atomicIntXorInvoker == nil {
		atomicIntXorInvoker = gi.FunctionInvokerNew("GLib", "atomic_int_xor")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetUint32(atomic)
	inArgs[1].SetUint32(val)

	ret := atomicIntXorInvoker.Invoke(inArgs[:], nil)

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

// UNSUPPORTED : C value 'g_atomic_rc_box_alloc' : parameter 'block_size' of type 'gsize' not supported

// UNSUPPORTED : C value 'g_atomic_rc_box_alloc0' : parameter 'block_size' of type 'gsize' not supported

// UNSUPPORTED : C value 'g_atomic_rc_box_dup' : parameter 'block_size' of type 'gsize' not supported

// UNSUPPORTED : C value 'g_atomic_rc_box_get_size' : parameter 'mem_block' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_atomic_rc_box_release' : parameter 'mem_block' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_atomic_rc_box_release_full' : parameter 'mem_block' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_atomic_ref_count_compare' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_atomic_ref_count_dec' : return type 'gboolean' not supported

var AtomicRefCountIncFunction *gi.Function
var AtomicRefCountIncFunctionOnce sync.Once

func AtomicRefCountIncFunctionSet() {
	AtomicRefCountIncFunctionOnce.Do(func() {
		AtomicRefCountIncFunction = gi.FunctionInvokerNew("GLib", "atomic_ref_count_inc")
	})
}

var atomicRefCountIncInvoker *gi.Function

// AtomicRefCountInc is a representation of the C type g_atomic_ref_count_inc.
func AtomicRefCountInc(arc int32) {
	if atomicRefCountIncInvoker == nil {
		atomicRefCountIncInvoker = gi.FunctionInvokerNew("GLib", "atomic_ref_count_inc")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetInt32(arc)

	atomicRefCountIncInvoker.Invoke(inArgs[:], nil)

}

var AtomicRefCountInitFunction *gi.Function
var AtomicRefCountInitFunctionOnce sync.Once

func AtomicRefCountInitFunctionSet() {
	AtomicRefCountInitFunctionOnce.Do(func() {
		AtomicRefCountInitFunction = gi.FunctionInvokerNew("GLib", "atomic_ref_count_init")
	})
}

var atomicRefCountInitInvoker *gi.Function

// AtomicRefCountInit is a representation of the C type g_atomic_ref_count_init.
func AtomicRefCountInit(arc int32) {
	if atomicRefCountInitInvoker == nil {
		atomicRefCountInitInvoker = gi.FunctionInvokerNew("GLib", "atomic_ref_count_init")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetInt32(arc)

	atomicRefCountInitInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_base64_decode' : parameter 'out_len' of type 'gsize' not supported

// UNSUPPORTED : C value 'g_base64_decode_inplace' : parameter 'text' has no type

// UNSUPPORTED : C value 'g_base64_decode_step' : parameter 'in' has no type

// UNSUPPORTED : C value 'g_base64_encode' : parameter 'data' has no type

// UNSUPPORTED : C value 'g_base64_encode_close' : parameter 'break_lines' of type 'gboolean' not supported

// UNSUPPORTED : C value 'g_base64_encode_step' : parameter 'in' has no type

// UNSUPPORTED : C value 'g_basename' : parameter 'file_name' of type 'filename' not supported

var BitLockFunction *gi.Function
var BitLockFunctionOnce sync.Once

func BitLockFunctionSet() {
	BitLockFunctionOnce.Do(func() {
		BitLockFunction = gi.FunctionInvokerNew("GLib", "bit_lock")
	})
}

var bitLockInvoker *gi.Function

// BitLock is a representation of the C type g_bit_lock.
func BitLock(address int32, lockBit int32) {
	if bitLockInvoker == nil {
		bitLockInvoker = gi.FunctionInvokerNew("GLib", "bit_lock")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetInt32(address)
	inArgs[1].SetInt32(lockBit)

	bitLockInvoker.Invoke(inArgs[:], nil)

}

var BitNthLsfFunction *gi.Function
var BitNthLsfFunctionOnce sync.Once

func BitNthLsfFunctionSet() {
	BitNthLsfFunctionOnce.Do(func() {
		BitNthLsfFunction = gi.FunctionInvokerNew("GLib", "bit_nth_lsf")
	})
}

var bitNthLsfInvoker *gi.Function

// BitNthLsf is a representation of the C type g_bit_nth_lsf.
func BitNthLsf(mask uint64, nthBit int32) int32 {
	if bitNthLsfInvoker == nil {
		bitNthLsfInvoker = gi.FunctionInvokerNew("GLib", "bit_nth_lsf")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetUint64(mask)
	inArgs[1].SetInt32(nthBit)

	ret := bitNthLsfInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var BitNthMsfFunction *gi.Function
var BitNthMsfFunctionOnce sync.Once

func BitNthMsfFunctionSet() {
	BitNthMsfFunctionOnce.Do(func() {
		BitNthMsfFunction = gi.FunctionInvokerNew("GLib", "bit_nth_msf")
	})
}

var bitNthMsfInvoker *gi.Function

// BitNthMsf is a representation of the C type g_bit_nth_msf.
func BitNthMsf(mask uint64, nthBit int32) int32 {
	if bitNthMsfInvoker == nil {
		bitNthMsfInvoker = gi.FunctionInvokerNew("GLib", "bit_nth_msf")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetUint64(mask)
	inArgs[1].SetInt32(nthBit)

	ret := bitNthMsfInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var BitStorageFunction *gi.Function
var BitStorageFunctionOnce sync.Once

func BitStorageFunctionSet() {
	BitStorageFunctionOnce.Do(func() {
		BitStorageFunction = gi.FunctionInvokerNew("GLib", "bit_storage")
	})
}

var bitStorageInvoker *gi.Function

// BitStorage is a representation of the C type g_bit_storage.
func BitStorage(number uint64) uint32 {
	if bitStorageInvoker == nil {
		bitStorageInvoker = gi.FunctionInvokerNew("GLib", "bit_storage")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetUint64(number)

	ret := bitStorageInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Uint32()

	return retGo
}

// UNSUPPORTED : C value 'g_bit_trylock' : return type 'gboolean' not supported

var BitUnlockFunction *gi.Function
var BitUnlockFunctionOnce sync.Once

func BitUnlockFunctionSet() {
	BitUnlockFunctionOnce.Do(func() {
		BitUnlockFunction = gi.FunctionInvokerNew("GLib", "bit_unlock")
	})
}

var bitUnlockInvoker *gi.Function

// BitUnlock is a representation of the C type g_bit_unlock.
func BitUnlock(address int32, lockBit int32) {
	if bitUnlockInvoker == nil {
		bitUnlockInvoker = gi.FunctionInvokerNew("GLib", "bit_unlock")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetInt32(address)
	inArgs[1].SetInt32(lockBit)

	bitUnlockInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_bookmark_file_error_quark' : return type 'Quark' not supported

// UNSUPPORTED : C value 'g_build_filename' : parameter 'first_element' of type 'filename' not supported

// UNSUPPORTED : C value 'g_build_filename_valist' : parameter 'first_element' of type 'filename' not supported

// UNSUPPORTED : C value 'g_build_filenamev' : parameter 'args' has no type

// UNSUPPORTED : C value 'g_build_path' : parameter 'separator' of type 'filename' not supported

// UNSUPPORTED : C value 'g_build_pathv' : parameter 'args' has no type

// UNSUPPORTED : C value 'g_byte_array_free' : parameter 'array' has no type

// UNSUPPORTED : C value 'g_byte_array_free_to_bytes' : parameter 'array' has no type

var ByteArrayNewFunction *gi.Function
var ByteArrayNewFunctionOnce sync.Once

func ByteArrayNewFunctionSet() {
	ByteArrayNewFunctionOnce.Do(func() {
		ByteArrayNewFunction = gi.FunctionInvokerNew("GLib", "byte_array_new")
	})
}

var byteArrayNewInvoker *gi.Function

// ByteArrayNew is a representation of the C type g_byte_array_new.
func ByteArrayNew() {
	if byteArrayNewInvoker == nil {
		byteArrayNewInvoker = gi.FunctionInvokerNew("GLib", "byte_array_new")
	}

	byteArrayNewInvoker.Invoke(nil, nil)

}

// UNSUPPORTED : C value 'g_byte_array_new_take' : parameter 'data' has no type

// UNSUPPORTED : C value 'g_byte_array_unref' : parameter 'array' has no type

// UNSUPPORTED : C value 'g_canonicalize_filename' : parameter 'filename' of type 'filename' not supported

// UNSUPPORTED : C value 'g_chdir' : parameter 'path' of type 'filename' not supported

var CheckVersionFunction *gi.Function
var CheckVersionFunctionOnce sync.Once

func CheckVersionFunctionSet() {
	CheckVersionFunctionOnce.Do(func() {
		CheckVersionFunction = gi.FunctionInvokerNew("GLib", "check_version")
	})
}

var checkVersionInvoker *gi.Function

// CheckVersion is a representation of the C type glib_check_version.
func CheckVersion(requiredMajor uint32, requiredMinor uint32, requiredMicro uint32) string {
	if checkVersionInvoker == nil {
		checkVersionInvoker = gi.FunctionInvokerNew("GLib", "check_version")
	}

	var inArgs [3]gi.Argument
	inArgs[0].SetUint32(requiredMajor)
	inArgs[1].SetUint32(requiredMinor)
	inArgs[2].SetUint32(requiredMicro)

	ret := checkVersionInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'g_checksum_type_get_length' : parameter 'checksum_type' of type 'ChecksumType' not supported

// UNSUPPORTED : C value 'g_child_watch_add' : parameter 'pid' of type 'Pid' not supported

// UNSUPPORTED : C value 'g_child_watch_add_full' : parameter 'pid' of type 'Pid' not supported

// UNSUPPORTED : C value 'g_child_watch_source_new' : parameter 'pid' of type 'Pid' not supported

var ClearErrorFunction *gi.Function
var ClearErrorFunctionOnce sync.Once

func ClearErrorFunctionSet() {
	ClearErrorFunctionOnce.Do(func() {
		ClearErrorFunction = gi.FunctionInvokerNew("GLib", "clear_error")
	})
}

var clearErrorInvoker *gi.Function

// ClearError is a representation of the C type g_clear_error.
func ClearError() {
	if clearErrorInvoker == nil {
		clearErrorInvoker = gi.FunctionInvokerNew("GLib", "clear_error")
	}

	clearErrorInvoker.Invoke(nil, nil)

}

// UNSUPPORTED : C value 'g_clear_handle_id' : parameter 'clear_func' of type 'ClearHandleFunc' not supported

// UNSUPPORTED : C value 'g_clear_pointer' : parameter 'pp' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_close' : return type 'gboolean' not supported

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

// UNSUPPORTED : C value 'g_date_get_monday_weeks_in_year' : parameter 'year' of type 'DateYear' not supported

// UNSUPPORTED : C value 'g_date_get_sunday_weeks_in_year' : parameter 'year' of type 'DateYear' not supported

// UNSUPPORTED : C value 'g_date_is_leap_year' : parameter 'year' of type 'DateYear' not supported

// UNSUPPORTED : C value 'g_date_strftime' : parameter 'slen' of type 'gsize' not supported

// UNSUPPORTED : C value 'g_date_time_compare' : parameter 'dt1' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_date_time_equal' : parameter 'dt1' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_date_time_hash' : parameter 'datetime' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_date_valid_day' : parameter 'day' of type 'DateDay' not supported

// UNSUPPORTED : C value 'g_date_valid_dmy' : parameter 'day' of type 'DateDay' not supported

// UNSUPPORTED : C value 'g_date_valid_julian' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_date_valid_month' : parameter 'month' of type 'DateMonth' not supported

// UNSUPPORTED : C value 'g_date_valid_weekday' : parameter 'weekday' of type 'DateWeekday' not supported

// UNSUPPORTED : C value 'g_date_valid_year' : parameter 'year' of type 'DateYear' not supported

var DcgettextFunction *gi.Function
var DcgettextFunctionOnce sync.Once

func DcgettextFunctionSet() {
	DcgettextFunctionOnce.Do(func() {
		DcgettextFunction = gi.FunctionInvokerNew("GLib", "dcgettext")
	})
}

var dcgettextInvoker *gi.Function

// Dcgettext is a representation of the C type g_dcgettext.
func Dcgettext(domain string, msgid string, category int32) string {
	if dcgettextInvoker == nil {
		dcgettextInvoker = gi.FunctionInvokerNew("GLib", "dcgettext")
	}

	var inArgs [3]gi.Argument
	inArgs[0].SetString(domain)
	inArgs[1].SetString(msgid)
	inArgs[2].SetInt32(category)

	ret := dcgettextInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

var DgettextFunction *gi.Function
var DgettextFunctionOnce sync.Once

func DgettextFunctionSet() {
	DgettextFunctionOnce.Do(func() {
		DgettextFunction = gi.FunctionInvokerNew("GLib", "dgettext")
	})
}

var dgettextInvoker *gi.Function

// Dgettext is a representation of the C type g_dgettext.
func Dgettext(domain string, msgid string) string {
	if dgettextInvoker == nil {
		dgettextInvoker = gi.FunctionInvokerNew("GLib", "dgettext")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetString(domain)
	inArgs[1].SetString(msgid)

	ret := dgettextInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'g_dir_make_tmp' : parameter 'tmpl' of type 'filename' not supported

// UNSUPPORTED : C value 'g_direct_equal' : parameter 'v1' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_direct_hash' : parameter 'v' of type 'gpointer' not supported

var DngettextFunction *gi.Function
var DngettextFunctionOnce sync.Once

func DngettextFunctionSet() {
	DngettextFunctionOnce.Do(func() {
		DngettextFunction = gi.FunctionInvokerNew("GLib", "dngettext")
	})
}

var dngettextInvoker *gi.Function

// Dngettext is a representation of the C type g_dngettext.
func Dngettext(domain string, msgid string, msgidPlural string, n uint64) string {
	if dngettextInvoker == nil {
		dngettextInvoker = gi.FunctionInvokerNew("GLib", "dngettext")
	}

	var inArgs [4]gi.Argument
	inArgs[0].SetString(domain)
	inArgs[1].SetString(msgid)
	inArgs[2].SetString(msgidPlural)
	inArgs[3].SetUint64(n)

	ret := dngettextInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'g_double_equal' : parameter 'v1' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_double_hash' : parameter 'v' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_dpgettext' : parameter 'msgidoffset' of type 'gsize' not supported

var Dpgettext2Function *gi.Function
var Dpgettext2FunctionOnce sync.Once

func Dpgettext2FunctionSet() {
	Dpgettext2FunctionOnce.Do(func() {
		Dpgettext2Function = gi.FunctionInvokerNew("GLib", "dpgettext2")
	})
}

var dpgettext2Invoker *gi.Function

// Dpgettext2 is a representation of the C type g_dpgettext2.
func Dpgettext2(domain string, context string, msgid string) string {
	if dpgettext2Invoker == nil {
		dpgettext2Invoker = gi.FunctionInvokerNew("GLib", "dpgettext2")
	}

	var inArgs [3]gi.Argument
	inArgs[0].SetString(domain)
	inArgs[1].SetString(context)
	inArgs[2].SetString(msgid)

	ret := dpgettext2Invoker.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
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

var FormatSizeFunction *gi.Function
var FormatSizeFunctionOnce sync.Once

func FormatSizeFunctionSet() {
	FormatSizeFunctionOnce.Do(func() {
		FormatSizeFunction = gi.FunctionInvokerNew("GLib", "format_size")
	})
}

var formatSizeInvoker *gi.Function

// FormatSize is a representation of the C type g_format_size.
func FormatSize(size uint64) string {
	if formatSizeInvoker == nil {
		formatSizeInvoker = gi.FunctionInvokerNew("GLib", "format_size")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetUint64(size)

	ret := formatSizeInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

var FormatSizeForDisplayFunction *gi.Function
var FormatSizeForDisplayFunctionOnce sync.Once

func FormatSizeForDisplayFunctionSet() {
	FormatSizeForDisplayFunctionOnce.Do(func() {
		FormatSizeForDisplayFunction = gi.FunctionInvokerNew("GLib", "format_size_for_display")
	})
}

var formatSizeForDisplayInvoker *gi.Function

// FormatSizeForDisplay is a representation of the C type g_format_size_for_display.
func FormatSizeForDisplay(size int64) string {
	if formatSizeForDisplayInvoker == nil {
		formatSizeForDisplayInvoker = gi.FunctionInvokerNew("GLib", "format_size_for_display")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetInt64(size)

	ret := formatSizeForDisplayInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

// UNSUPPORTED : C value 'g_format_size_full' : parameter 'flags' of type 'FormatSizeFlags' not supported

// UNSUPPORTED : C value 'g_fprintf' : parameter 'file' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_free' : parameter 'mem' of type 'gpointer' not supported

var GetApplicationNameFunction *gi.Function
var GetApplicationNameFunctionOnce sync.Once

func GetApplicationNameFunctionSet() {
	GetApplicationNameFunctionOnce.Do(func() {
		GetApplicationNameFunction = gi.FunctionInvokerNew("GLib", "get_application_name")
	})
}

var getApplicationNameInvoker *gi.Function

// GetApplicationName is a representation of the C type g_get_application_name.
func GetApplicationName() string {
	if getApplicationNameInvoker == nil {
		getApplicationNameInvoker = gi.FunctionInvokerNew("GLib", "get_application_name")
	}

	ret := getApplicationNameInvoker.Invoke(nil, nil)

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'g_get_charset' : return type 'gboolean' not supported

var GetCodesetFunction *gi.Function
var GetCodesetFunctionOnce sync.Once

func GetCodesetFunctionSet() {
	GetCodesetFunctionOnce.Do(func() {
		GetCodesetFunction = gi.FunctionInvokerNew("GLib", "get_codeset")
	})
}

var getCodesetInvoker *gi.Function

// GetCodeset is a representation of the C type g_get_codeset.
func GetCodeset() string {
	if getCodesetInvoker == nil {
		getCodesetInvoker = gi.FunctionInvokerNew("GLib", "get_codeset")
	}

	ret := getCodesetInvoker.Invoke(nil, nil)

	retGo := ret.String(true)

	return retGo
}

// UNSUPPORTED : C value 'g_get_console_charset' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_get_current_dir' : return type 'filename' not supported

// UNSUPPORTED : C value 'g_get_current_time' : parameter 'result' of type 'TimeVal' not supported

var GetEnvironFunction *gi.Function
var GetEnvironFunctionOnce sync.Once

func GetEnvironFunctionSet() {
	GetEnvironFunctionOnce.Do(func() {
		GetEnvironFunction = gi.FunctionInvokerNew("GLib", "get_environ")
	})
}

var getEnvironInvoker *gi.Function

// GetEnviron is a representation of the C type g_get_environ.
func GetEnviron() {
	if getEnvironInvoker == nil {
		getEnvironInvoker = gi.FunctionInvokerNew("GLib", "get_environ")
	}

	getEnvironInvoker.Invoke(nil, nil)

}

// UNSUPPORTED : C value 'g_get_filename_charsets' : parameter 'filename_charsets' has no type

// UNSUPPORTED : C value 'g_get_home_dir' : return type 'filename' not supported

var GetHostNameFunction *gi.Function
var GetHostNameFunctionOnce sync.Once

func GetHostNameFunctionSet() {
	GetHostNameFunctionOnce.Do(func() {
		GetHostNameFunction = gi.FunctionInvokerNew("GLib", "get_host_name")
	})
}

var getHostNameInvoker *gi.Function

// GetHostName is a representation of the C type g_get_host_name.
func GetHostName() string {
	if getHostNameInvoker == nil {
		getHostNameInvoker = gi.FunctionInvokerNew("GLib", "get_host_name")
	}

	ret := getHostNameInvoker.Invoke(nil, nil)

	retGo := ret.String(false)

	return retGo
}

var GetLanguageNamesFunction *gi.Function
var GetLanguageNamesFunctionOnce sync.Once

func GetLanguageNamesFunctionSet() {
	GetLanguageNamesFunctionOnce.Do(func() {
		GetLanguageNamesFunction = gi.FunctionInvokerNew("GLib", "get_language_names")
	})
}

var getLanguageNamesInvoker *gi.Function

// GetLanguageNames is a representation of the C type g_get_language_names.
func GetLanguageNames() {
	if getLanguageNamesInvoker == nil {
		getLanguageNamesInvoker = gi.FunctionInvokerNew("GLib", "get_language_names")
	}

	getLanguageNamesInvoker.Invoke(nil, nil)

}

var GetLanguageNamesWithCategoryFunction *gi.Function
var GetLanguageNamesWithCategoryFunctionOnce sync.Once

func GetLanguageNamesWithCategoryFunctionSet() {
	GetLanguageNamesWithCategoryFunctionOnce.Do(func() {
		GetLanguageNamesWithCategoryFunction = gi.FunctionInvokerNew("GLib", "get_language_names_with_category")
	})
}

var getLanguageNamesWithCategoryInvoker *gi.Function

// GetLanguageNamesWithCategory is a representation of the C type g_get_language_names_with_category.
func GetLanguageNamesWithCategory(categoryName string) {
	if getLanguageNamesWithCategoryInvoker == nil {
		getLanguageNamesWithCategoryInvoker = gi.FunctionInvokerNew("GLib", "get_language_names_with_category")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetString(categoryName)

	getLanguageNamesWithCategoryInvoker.Invoke(inArgs[:], nil)

}

var GetLocaleVariantsFunction *gi.Function
var GetLocaleVariantsFunctionOnce sync.Once

func GetLocaleVariantsFunctionSet() {
	GetLocaleVariantsFunctionOnce.Do(func() {
		GetLocaleVariantsFunction = gi.FunctionInvokerNew("GLib", "get_locale_variants")
	})
}

var getLocaleVariantsInvoker *gi.Function

// GetLocaleVariants is a representation of the C type g_get_locale_variants.
func GetLocaleVariants(locale string) {
	if getLocaleVariantsInvoker == nil {
		getLocaleVariantsInvoker = gi.FunctionInvokerNew("GLib", "get_locale_variants")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetString(locale)

	getLocaleVariantsInvoker.Invoke(inArgs[:], nil)

}

var GetMonotonicTimeFunction *gi.Function
var GetMonotonicTimeFunctionOnce sync.Once

func GetMonotonicTimeFunctionSet() {
	GetMonotonicTimeFunctionOnce.Do(func() {
		GetMonotonicTimeFunction = gi.FunctionInvokerNew("GLib", "get_monotonic_time")
	})
}

var getMonotonicTimeInvoker *gi.Function

// GetMonotonicTime is a representation of the C type g_get_monotonic_time.
func GetMonotonicTime() int64 {
	if getMonotonicTimeInvoker == nil {
		getMonotonicTimeInvoker = gi.FunctionInvokerNew("GLib", "get_monotonic_time")
	}

	ret := getMonotonicTimeInvoker.Invoke(nil, nil)

	retGo := ret.Int64()

	return retGo
}

var GetNumProcessorsFunction *gi.Function
var GetNumProcessorsFunctionOnce sync.Once

func GetNumProcessorsFunctionSet() {
	GetNumProcessorsFunctionOnce.Do(func() {
		GetNumProcessorsFunction = gi.FunctionInvokerNew("GLib", "get_num_processors")
	})
}

var getNumProcessorsInvoker *gi.Function

// GetNumProcessors is a representation of the C type g_get_num_processors.
func GetNumProcessors() uint32 {
	if getNumProcessorsInvoker == nil {
		getNumProcessorsInvoker = gi.FunctionInvokerNew("GLib", "get_num_processors")
	}

	ret := getNumProcessorsInvoker.Invoke(nil, nil)

	retGo := ret.Uint32()

	return retGo
}

var GetPrgnameFunction *gi.Function
var GetPrgnameFunctionOnce sync.Once

func GetPrgnameFunctionSet() {
	GetPrgnameFunctionOnce.Do(func() {
		GetPrgnameFunction = gi.FunctionInvokerNew("GLib", "get_prgname")
	})
}

var getPrgnameInvoker *gi.Function

// GetPrgname is a representation of the C type g_get_prgname.
func GetPrgname() string {
	if getPrgnameInvoker == nil {
		getPrgnameInvoker = gi.FunctionInvokerNew("GLib", "get_prgname")
	}

	ret := getPrgnameInvoker.Invoke(nil, nil)

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'g_get_real_name' : return type 'filename' not supported

var GetRealTimeFunction *gi.Function
var GetRealTimeFunctionOnce sync.Once

func GetRealTimeFunctionSet() {
	GetRealTimeFunctionOnce.Do(func() {
		GetRealTimeFunction = gi.FunctionInvokerNew("GLib", "get_real_time")
	})
}

var getRealTimeInvoker *gi.Function

// GetRealTime is a representation of the C type g_get_real_time.
func GetRealTime() int64 {
	if getRealTimeInvoker == nil {
		getRealTimeInvoker = gi.FunctionInvokerNew("GLib", "get_real_time")
	}

	ret := getRealTimeInvoker.Invoke(nil, nil)

	retGo := ret.Int64()

	return retGo
}

var GetSystemConfigDirsFunction *gi.Function
var GetSystemConfigDirsFunctionOnce sync.Once

func GetSystemConfigDirsFunctionSet() {
	GetSystemConfigDirsFunctionOnce.Do(func() {
		GetSystemConfigDirsFunction = gi.FunctionInvokerNew("GLib", "get_system_config_dirs")
	})
}

var getSystemConfigDirsInvoker *gi.Function

// GetSystemConfigDirs is a representation of the C type g_get_system_config_dirs.
func GetSystemConfigDirs() {
	if getSystemConfigDirsInvoker == nil {
		getSystemConfigDirsInvoker = gi.FunctionInvokerNew("GLib", "get_system_config_dirs")
	}

	getSystemConfigDirsInvoker.Invoke(nil, nil)

}

var GetSystemDataDirsFunction *gi.Function
var GetSystemDataDirsFunctionOnce sync.Once

func GetSystemDataDirsFunctionSet() {
	GetSystemDataDirsFunctionOnce.Do(func() {
		GetSystemDataDirsFunction = gi.FunctionInvokerNew("GLib", "get_system_data_dirs")
	})
}

var getSystemDataDirsInvoker *gi.Function

// GetSystemDataDirs is a representation of the C type g_get_system_data_dirs.
func GetSystemDataDirs() {
	if getSystemDataDirsInvoker == nil {
		getSystemDataDirsInvoker = gi.FunctionInvokerNew("GLib", "get_system_data_dirs")
	}

	getSystemDataDirsInvoker.Invoke(nil, nil)

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

// UNSUPPORTED : C value 'g_hostname_is_ascii_encoded' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_hostname_is_ip_address' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_hostname_is_non_ascii' : return type 'gboolean' not supported

var HostnameToAsciiFunction *gi.Function
var HostnameToAsciiFunctionOnce sync.Once

func HostnameToAsciiFunctionSet() {
	HostnameToAsciiFunctionOnce.Do(func() {
		HostnameToAsciiFunction = gi.FunctionInvokerNew("GLib", "hostname_to_ascii")
	})
}

var hostnameToAsciiInvoker *gi.Function

// HostnameToAscii is a representation of the C type g_hostname_to_ascii.
func HostnameToAscii(hostname string) string {
	if hostnameToAsciiInvoker == nil {
		hostnameToAsciiInvoker = gi.FunctionInvokerNew("GLib", "hostname_to_ascii")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetString(hostname)

	ret := hostnameToAsciiInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

var HostnameToUnicodeFunction *gi.Function
var HostnameToUnicodeFunctionOnce sync.Once

func HostnameToUnicodeFunctionSet() {
	HostnameToUnicodeFunctionOnce.Do(func() {
		HostnameToUnicodeFunction = gi.FunctionInvokerNew("GLib", "hostname_to_unicode")
	})
}

var hostnameToUnicodeInvoker *gi.Function

// HostnameToUnicode is a representation of the C type g_hostname_to_unicode.
func HostnameToUnicode(hostname string) string {
	if hostnameToUnicodeInvoker == nil {
		hostnameToUnicodeInvoker = gi.FunctionInvokerNew("GLib", "hostname_to_unicode")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetString(hostname)

	ret := hostnameToUnicodeInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

// UNSUPPORTED : C value 'g_iconv' : parameter 'converter' of type 'IConv' not supported

var IconvOpenFunction *gi.Function
var IconvOpenFunctionOnce sync.Once

func IconvOpenFunctionSet() {
	IconvOpenFunctionOnce.Do(func() {
		IconvOpenFunction = gi.FunctionInvokerNew("GLib", "iconv_open")
	})
}

var iconvOpenInvoker *gi.Function

// IconvOpen is a representation of the C type g_iconv_open.
func IconvOpen(toCodeset string, fromCodeset string) *IConv {
	if iconvOpenInvoker == nil {
		iconvOpenInvoker = gi.FunctionInvokerNew("GLib", "iconv_open")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetString(toCodeset)
	inArgs[1].SetString(fromCodeset)

	ret := iconvOpenInvoker.Invoke(inArgs[:], nil)

	retGo := &IConv{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_idle_add' : parameter 'function' of type 'SourceFunc' not supported

// UNSUPPORTED : C value 'g_idle_add_full' : parameter 'function' of type 'SourceFunc' not supported

// UNSUPPORTED : C value 'g_idle_remove_by_data' : parameter 'data' of type 'gpointer' not supported

var IdleSourceNewFunction *gi.Function
var IdleSourceNewFunctionOnce sync.Once

func IdleSourceNewFunctionSet() {
	IdleSourceNewFunctionOnce.Do(func() {
		IdleSourceNewFunction = gi.FunctionInvokerNew("GLib", "idle_source_new")
	})
}

var idleSourceNewInvoker *gi.Function

// IdleSourceNew is a representation of the C type g_idle_source_new.
func IdleSourceNew() *Source {
	if idleSourceNewInvoker == nil {
		idleSourceNewInvoker = gi.FunctionInvokerNew("GLib", "idle_source_new")
	}

	ret := idleSourceNewInvoker.Invoke(nil, nil)

	retGo := &Source{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_int64_equal' : parameter 'v1' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_int64_hash' : parameter 'v' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_int_equal' : parameter 'v1' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_int_hash' : parameter 'v' of type 'gpointer' not supported

var InternStaticStringFunction *gi.Function
var InternStaticStringFunctionOnce sync.Once

func InternStaticStringFunctionSet() {
	InternStaticStringFunctionOnce.Do(func() {
		InternStaticStringFunction = gi.FunctionInvokerNew("GLib", "intern_static_string")
	})
}

var internStaticStringInvoker *gi.Function

// InternStaticString is a representation of the C type g_intern_static_string.
func InternStaticString(string_ string) string {
	if internStaticStringInvoker == nil {
		internStaticStringInvoker = gi.FunctionInvokerNew("GLib", "intern_static_string")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetString(string_)

	ret := internStaticStringInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

var InternStringFunction *gi.Function
var InternStringFunctionOnce sync.Once

func InternStringFunctionSet() {
	InternStringFunctionOnce.Do(func() {
		InternStringFunction = gi.FunctionInvokerNew("GLib", "intern_string")
	})
}

var internStringInvoker *gi.Function

// InternString is a representation of the C type g_intern_string.
func InternString(string_ string) string {
	if internStringInvoker == nil {
		internStringInvoker = gi.FunctionInvokerNew("GLib", "intern_string")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetString(string_)

	ret := internStringInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'g_io_add_watch' : parameter 'channel' of type 'IOChannel' not supported

// UNSUPPORTED : C value 'g_io_add_watch_full' : parameter 'channel' of type 'IOChannel' not supported

// UNSUPPORTED : C value 'g_io_channel_error_from_errno' : return type 'IOChannelError' not supported

// UNSUPPORTED : C value 'g_io_channel_error_quark' : return type 'Quark' not supported

// UNSUPPORTED : C value 'g_io_create_watch' : parameter 'channel' of type 'IOChannel' not supported

// UNSUPPORTED : C value 'g_key_file_error_quark' : return type 'Quark' not supported

var ListenvFunction *gi.Function
var ListenvFunctionOnce sync.Once

func ListenvFunctionSet() {
	ListenvFunctionOnce.Do(func() {
		ListenvFunction = gi.FunctionInvokerNew("GLib", "listenv")
	})
}

var listenvInvoker *gi.Function

// Listenv is a representation of the C type g_listenv.
func Listenv() {
	if listenvInvoker == nil {
		listenvInvoker = gi.FunctionInvokerNew("GLib", "listenv")
	}

	listenvInvoker.Invoke(nil, nil)

}

// UNSUPPORTED : C value 'g_locale_from_utf8' : parameter 'bytes_read' of type 'gsize' not supported

// UNSUPPORTED : C value 'g_locale_to_utf8' : parameter 'opsysstring' has no type

// UNSUPPORTED : C value 'g_log' : parameter 'log_level' of type 'LogLevelFlags' not supported

// UNSUPPORTED : C value 'g_log_default_handler' : parameter 'log_level' of type 'LogLevelFlags' not supported

var LogRemoveHandlerFunction *gi.Function
var LogRemoveHandlerFunctionOnce sync.Once

func LogRemoveHandlerFunctionSet() {
	LogRemoveHandlerFunctionOnce.Do(func() {
		LogRemoveHandlerFunction = gi.FunctionInvokerNew("GLib", "log_remove_handler")
	})
}

var logRemoveHandlerInvoker *gi.Function

// LogRemoveHandler is a representation of the C type g_log_remove_handler.
func LogRemoveHandler(logDomain string, handlerId uint32) {
	if logRemoveHandlerInvoker == nil {
		logRemoveHandlerInvoker = gi.FunctionInvokerNew("GLib", "log_remove_handler")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetString(logDomain)
	inArgs[1].SetUint32(handlerId)

	logRemoveHandlerInvoker.Invoke(inArgs[:], nil)

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

// UNSUPPORTED : C value 'g_log_writer_is_journald' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_log_writer_journald' : parameter 'log_level' of type 'LogLevelFlags' not supported

// UNSUPPORTED : C value 'g_log_writer_standard_streams' : parameter 'log_level' of type 'LogLevelFlags' not supported

// UNSUPPORTED : C value 'g_log_writer_supports_color' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_logv' : parameter 'log_level' of type 'LogLevelFlags' not supported

var MainContextDefaultFunction *gi.Function
var MainContextDefaultFunctionOnce sync.Once

func MainContextDefaultFunctionSet() {
	MainContextDefaultFunctionOnce.Do(func() {
		MainContextDefaultFunction = gi.FunctionInvokerNew("GLib", "main_context_default")
	})
}

var mainContextDefaultInvoker *gi.Function

// MainContextDefault is a representation of the C type g_main_context_default.
func MainContextDefault() *MainContext {
	if mainContextDefaultInvoker == nil {
		mainContextDefaultInvoker = gi.FunctionInvokerNew("GLib", "main_context_default")
	}

	ret := mainContextDefaultInvoker.Invoke(nil, nil)

	retGo := &MainContext{native: ret.Pointer()}

	return retGo
}

var MainContextGetThreadDefaultFunction *gi.Function
var MainContextGetThreadDefaultFunctionOnce sync.Once

func MainContextGetThreadDefaultFunctionSet() {
	MainContextGetThreadDefaultFunctionOnce.Do(func() {
		MainContextGetThreadDefaultFunction = gi.FunctionInvokerNew("GLib", "main_context_get_thread_default")
	})
}

var mainContextGetThreadDefaultInvoker *gi.Function

// MainContextGetThreadDefault is a representation of the C type g_main_context_get_thread_default.
func MainContextGetThreadDefault() *MainContext {
	if mainContextGetThreadDefaultInvoker == nil {
		mainContextGetThreadDefaultInvoker = gi.FunctionInvokerNew("GLib", "main_context_get_thread_default")
	}

	ret := mainContextGetThreadDefaultInvoker.Invoke(nil, nil)

	retGo := &MainContext{native: ret.Pointer()}

	return retGo
}

var MainContextRefThreadDefaultFunction *gi.Function
var MainContextRefThreadDefaultFunctionOnce sync.Once

func MainContextRefThreadDefaultFunctionSet() {
	MainContextRefThreadDefaultFunctionOnce.Do(func() {
		MainContextRefThreadDefaultFunction = gi.FunctionInvokerNew("GLib", "main_context_ref_thread_default")
	})
}

var mainContextRefThreadDefaultInvoker *gi.Function

// MainContextRefThreadDefault is a representation of the C type g_main_context_ref_thread_default.
func MainContextRefThreadDefault() *MainContext {
	if mainContextRefThreadDefaultInvoker == nil {
		mainContextRefThreadDefaultInvoker = gi.FunctionInvokerNew("GLib", "main_context_ref_thread_default")
	}

	ret := mainContextRefThreadDefaultInvoker.Invoke(nil, nil)

	retGo := &MainContext{native: ret.Pointer()}

	return retGo
}

var MainCurrentSourceFunction *gi.Function
var MainCurrentSourceFunctionOnce sync.Once

func MainCurrentSourceFunctionSet() {
	MainCurrentSourceFunctionOnce.Do(func() {
		MainCurrentSourceFunction = gi.FunctionInvokerNew("GLib", "main_current_source")
	})
}

var mainCurrentSourceInvoker *gi.Function

// MainCurrentSource is a representation of the C type g_main_current_source.
func MainCurrentSource() *Source {
	if mainCurrentSourceInvoker == nil {
		mainCurrentSourceInvoker = gi.FunctionInvokerNew("GLib", "main_current_source")
	}

	ret := mainCurrentSourceInvoker.Invoke(nil, nil)

	retGo := &Source{native: ret.Pointer()}

	return retGo
}

var MainDepthFunction *gi.Function
var MainDepthFunctionOnce sync.Once

func MainDepthFunctionSet() {
	MainDepthFunctionOnce.Do(func() {
		MainDepthFunction = gi.FunctionInvokerNew("GLib", "main_depth")
	})
}

var mainDepthInvoker *gi.Function

// MainDepth is a representation of the C type g_main_depth.
func MainDepth() int32 {
	if mainDepthInvoker == nil {
		mainDepthInvoker = gi.FunctionInvokerNew("GLib", "main_depth")
	}

	ret := mainDepthInvoker.Invoke(nil, nil)

	retGo := ret.Int32()

	return retGo
}

// UNSUPPORTED : C value 'g_malloc' : parameter 'n_bytes' of type 'gsize' not supported

// UNSUPPORTED : C value 'g_malloc0' : parameter 'n_bytes' of type 'gsize' not supported

// UNSUPPORTED : C value 'g_malloc0_n' : parameter 'n_blocks' of type 'gsize' not supported

// UNSUPPORTED : C value 'g_malloc_n' : parameter 'n_blocks' of type 'gsize' not supported

// UNSUPPORTED : C value 'g_markup_collect_attributes' : parameter 'error' of type 'Error' not supported

// UNSUPPORTED : C value 'g_markup_error_quark' : return type 'Quark' not supported

var MarkupEscapeTextFunction *gi.Function
var MarkupEscapeTextFunctionOnce sync.Once

func MarkupEscapeTextFunctionSet() {
	MarkupEscapeTextFunctionOnce.Do(func() {
		MarkupEscapeTextFunction = gi.FunctionInvokerNew("GLib", "markup_escape_text")
	})
}

var markupEscapeTextInvoker *gi.Function

// MarkupEscapeText is a representation of the C type g_markup_escape_text.
func MarkupEscapeText(text string, length int32) string {
	if markupEscapeTextInvoker == nil {
		markupEscapeTextInvoker = gi.FunctionInvokerNew("GLib", "markup_escape_text")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetString(text)
	inArgs[1].SetInt32(length)

	ret := markupEscapeTextInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

// UNSUPPORTED : C value 'g_markup_printf_escaped' : parameter '...' has no type

// UNSUPPORTED : C value 'g_markup_vprintf_escaped' : parameter 'args' of type 'va_list' not supported

// UNSUPPORTED : C value 'g_mem_is_system_malloc' : return type 'gboolean' not supported

var MemProfileFunction *gi.Function
var MemProfileFunctionOnce sync.Once

func MemProfileFunctionSet() {
	MemProfileFunctionOnce.Do(func() {
		MemProfileFunction = gi.FunctionInvokerNew("GLib", "mem_profile")
	})
}

var memProfileInvoker *gi.Function

// MemProfile is a representation of the C type g_mem_profile.
func MemProfile() {
	if memProfileInvoker == nil {
		memProfileInvoker = gi.FunctionInvokerNew("GLib", "mem_profile")
	}

	memProfileInvoker.Invoke(nil, nil)

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

var OnErrorQueryFunction *gi.Function
var OnErrorQueryFunctionOnce sync.Once

func OnErrorQueryFunctionSet() {
	OnErrorQueryFunctionOnce.Do(func() {
		OnErrorQueryFunction = gi.FunctionInvokerNew("GLib", "on_error_query")
	})
}

var onErrorQueryInvoker *gi.Function

// OnErrorQuery is a representation of the C type g_on_error_query.
func OnErrorQuery(prgName string) {
	if onErrorQueryInvoker == nil {
		onErrorQueryInvoker = gi.FunctionInvokerNew("GLib", "on_error_query")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetString(prgName)

	onErrorQueryInvoker.Invoke(inArgs[:], nil)

}

var OnErrorStackTraceFunction *gi.Function
var OnErrorStackTraceFunctionOnce sync.Once

func OnErrorStackTraceFunctionSet() {
	OnErrorStackTraceFunctionOnce.Do(func() {
		OnErrorStackTraceFunction = gi.FunctionInvokerNew("GLib", "on_error_stack_trace")
	})
}

var onErrorStackTraceInvoker *gi.Function

// OnErrorStackTrace is a representation of the C type g_on_error_stack_trace.
func OnErrorStackTrace(prgName string) {
	if onErrorStackTraceInvoker == nil {
		onErrorStackTraceInvoker = gi.FunctionInvokerNew("GLib", "on_error_stack_trace")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetString(prgName)

	onErrorStackTraceInvoker.Invoke(inArgs[:], nil)

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

// UNSUPPORTED : C value 'g_pattern_match_simple' : return type 'gboolean' not supported

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

// UNSUPPORTED : C value 'g_quark_to_string' : parameter 'quark' of type 'Quark' not supported

// UNSUPPORTED : C value 'g_quark_try_string' : return type 'Quark' not supported

// UNSUPPORTED : C value 'g_random_double' : return type 'gdouble' not supported

// UNSUPPORTED : C value 'g_random_double_range' : parameter 'begin' of type 'gdouble' not supported

var RandomIntFunction *gi.Function
var RandomIntFunctionOnce sync.Once

func RandomIntFunctionSet() {
	RandomIntFunctionOnce.Do(func() {
		RandomIntFunction = gi.FunctionInvokerNew("GLib", "random_int")
	})
}

var randomIntInvoker *gi.Function

// RandomInt is a representation of the C type g_random_int.
func RandomInt() uint32 {
	if randomIntInvoker == nil {
		randomIntInvoker = gi.FunctionInvokerNew("GLib", "random_int")
	}

	ret := randomIntInvoker.Invoke(nil, nil)

	retGo := ret.Uint32()

	return retGo
}

var RandomIntRangeFunction *gi.Function
var RandomIntRangeFunctionOnce sync.Once

func RandomIntRangeFunctionSet() {
	RandomIntRangeFunctionOnce.Do(func() {
		RandomIntRangeFunction = gi.FunctionInvokerNew("GLib", "random_int_range")
	})
}

var randomIntRangeInvoker *gi.Function

// RandomIntRange is a representation of the C type g_random_int_range.
func RandomIntRange(begin int32, end int32) int32 {
	if randomIntRangeInvoker == nil {
		randomIntRangeInvoker = gi.FunctionInvokerNew("GLib", "random_int_range")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetInt32(begin)
	inArgs[1].SetInt32(end)

	ret := randomIntRangeInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var RandomSetSeedFunction *gi.Function
var RandomSetSeedFunctionOnce sync.Once

func RandomSetSeedFunctionSet() {
	RandomSetSeedFunctionOnce.Do(func() {
		RandomSetSeedFunction = gi.FunctionInvokerNew("GLib", "random_set_seed")
	})
}

var randomSetSeedInvoker *gi.Function

// RandomSetSeed is a representation of the C type g_random_set_seed.
func RandomSetSeed(seed uint32) {
	if randomSetSeedInvoker == nil {
		randomSetSeedInvoker = gi.FunctionInvokerNew("GLib", "random_set_seed")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetUint32(seed)

	randomSetSeedInvoker.Invoke(inArgs[:], nil)

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

// UNSUPPORTED : C value 'g_ref_count_compare' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_ref_count_dec' : return type 'gboolean' not supported

var RefCountIncFunction *gi.Function
var RefCountIncFunctionOnce sync.Once

func RefCountIncFunctionSet() {
	RefCountIncFunctionOnce.Do(func() {
		RefCountIncFunction = gi.FunctionInvokerNew("GLib", "ref_count_inc")
	})
}

var refCountIncInvoker *gi.Function

// RefCountInc is a representation of the C type g_ref_count_inc.
func RefCountInc(rc int32) {
	if refCountIncInvoker == nil {
		refCountIncInvoker = gi.FunctionInvokerNew("GLib", "ref_count_inc")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetInt32(rc)

	refCountIncInvoker.Invoke(inArgs[:], nil)

}

var RefCountInitFunction *gi.Function
var RefCountInitFunctionOnce sync.Once

func RefCountInitFunctionSet() {
	RefCountInitFunctionOnce.Do(func() {
		RefCountInitFunction = gi.FunctionInvokerNew("GLib", "ref_count_init")
	})
}

var refCountInitInvoker *gi.Function

// RefCountInit is a representation of the C type g_ref_count_init.
func RefCountInit(rc int32) {
	if refCountInitInvoker == nil {
		refCountInitInvoker = gi.FunctionInvokerNew("GLib", "ref_count_init")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetInt32(rc)

	refCountInitInvoker.Invoke(inArgs[:], nil)

}

var RefStringAcquireFunction *gi.Function
var RefStringAcquireFunctionOnce sync.Once

func RefStringAcquireFunctionSet() {
	RefStringAcquireFunctionOnce.Do(func() {
		RefStringAcquireFunction = gi.FunctionInvokerNew("GLib", "ref_string_acquire")
	})
}

var refStringAcquireInvoker *gi.Function

// RefStringAcquire is a representation of the C type g_ref_string_acquire.
func RefStringAcquire(str string) string {
	if refStringAcquireInvoker == nil {
		refStringAcquireInvoker = gi.FunctionInvokerNew("GLib", "ref_string_acquire")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetString(str)

	ret := refStringAcquireInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

// UNSUPPORTED : C value 'g_ref_string_length' : return type 'gsize' not supported

var RefStringNewFunction *gi.Function
var RefStringNewFunctionOnce sync.Once

func RefStringNewFunctionSet() {
	RefStringNewFunctionOnce.Do(func() {
		RefStringNewFunction = gi.FunctionInvokerNew("GLib", "ref_string_new")
	})
}

var refStringNewInvoker *gi.Function

// RefStringNew is a representation of the C type g_ref_string_new.
func RefStringNew(str string) string {
	if refStringNewInvoker == nil {
		refStringNewInvoker = gi.FunctionInvokerNew("GLib", "ref_string_new")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetString(str)

	ret := refStringNewInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

var RefStringNewInternFunction *gi.Function
var RefStringNewInternFunctionOnce sync.Once

func RefStringNewInternFunctionSet() {
	RefStringNewInternFunctionOnce.Do(func() {
		RefStringNewInternFunction = gi.FunctionInvokerNew("GLib", "ref_string_new_intern")
	})
}

var refStringNewInternInvoker *gi.Function

// RefStringNewIntern is a representation of the C type g_ref_string_new_intern.
func RefStringNewIntern(str string) string {
	if refStringNewInternInvoker == nil {
		refStringNewInternInvoker = gi.FunctionInvokerNew("GLib", "ref_string_new_intern")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetString(str)

	ret := refStringNewInternInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

var RefStringNewLenFunction *gi.Function
var RefStringNewLenFunctionOnce sync.Once

func RefStringNewLenFunctionSet() {
	RefStringNewLenFunctionOnce.Do(func() {
		RefStringNewLenFunction = gi.FunctionInvokerNew("GLib", "ref_string_new_len")
	})
}

var refStringNewLenInvoker *gi.Function

// RefStringNewLen is a representation of the C type g_ref_string_new_len.
func RefStringNewLen(str string, len int32) string {
	if refStringNewLenInvoker == nil {
		refStringNewLenInvoker = gi.FunctionInvokerNew("GLib", "ref_string_new_len")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetString(str)
	inArgs[1].SetInt32(len)

	ret := refStringNewLenInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

var RefStringReleaseFunction *gi.Function
var RefStringReleaseFunctionOnce sync.Once

func RefStringReleaseFunctionSet() {
	RefStringReleaseFunctionOnce.Do(func() {
		RefStringReleaseFunction = gi.FunctionInvokerNew("GLib", "ref_string_release")
	})
}

var refStringReleaseInvoker *gi.Function

// RefStringRelease is a representation of the C type g_ref_string_release.
func RefStringRelease(str string) {
	if refStringReleaseInvoker == nil {
		refStringReleaseInvoker = gi.FunctionInvokerNew("GLib", "ref_string_release")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetString(str)

	refStringReleaseInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_regex_check_replacement' : parameter 'has_references' of type 'gboolean' not supported

// UNSUPPORTED : C value 'g_regex_error_quark' : return type 'Quark' not supported

var RegexEscapeNulFunction *gi.Function
var RegexEscapeNulFunctionOnce sync.Once

func RegexEscapeNulFunctionSet() {
	RegexEscapeNulFunctionOnce.Do(func() {
		RegexEscapeNulFunction = gi.FunctionInvokerNew("GLib", "regex_escape_nul")
	})
}

var regexEscapeNulInvoker *gi.Function

// RegexEscapeNul is a representation of the C type g_regex_escape_nul.
func RegexEscapeNul(string_ string, length int32) string {
	if regexEscapeNulInvoker == nil {
		regexEscapeNulInvoker = gi.FunctionInvokerNew("GLib", "regex_escape_nul")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetString(string_)
	inArgs[1].SetInt32(length)

	ret := regexEscapeNulInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

// UNSUPPORTED : C value 'g_regex_escape_string' : parameter 'string' has no type

// UNSUPPORTED : C value 'g_regex_match_simple' : parameter 'compile_options' of type 'RegexCompileFlags' not supported

// UNSUPPORTED : C value 'g_regex_split_simple' : parameter 'compile_options' of type 'RegexCompileFlags' not supported

var ReloadUserSpecialDirsCacheFunction *gi.Function
var ReloadUserSpecialDirsCacheFunctionOnce sync.Once

func ReloadUserSpecialDirsCacheFunctionSet() {
	ReloadUserSpecialDirsCacheFunctionOnce.Do(func() {
		ReloadUserSpecialDirsCacheFunction = gi.FunctionInvokerNew("GLib", "reload_user_special_dirs_cache")
	})
}

var reloadUserSpecialDirsCacheInvoker *gi.Function

// ReloadUserSpecialDirsCache is a representation of the C type g_reload_user_special_dirs_cache.
func ReloadUserSpecialDirsCache() {
	if reloadUserSpecialDirsCacheInvoker == nil {
		reloadUserSpecialDirsCacheInvoker = gi.FunctionInvokerNew("GLib", "reload_user_special_dirs_cache")
	}

	reloadUserSpecialDirsCacheInvoker.Invoke(nil, nil)

}

var ReturnIfFailWarningFunction *gi.Function
var ReturnIfFailWarningFunctionOnce sync.Once

func ReturnIfFailWarningFunctionSet() {
	ReturnIfFailWarningFunctionOnce.Do(func() {
		ReturnIfFailWarningFunction = gi.FunctionInvokerNew("GLib", "return_if_fail_warning")
	})
}

var returnIfFailWarningInvoker *gi.Function

// ReturnIfFailWarning is a representation of the C type g_return_if_fail_warning.
func ReturnIfFailWarning(logDomain string, prettyFunction string, expression string) {
	if returnIfFailWarningInvoker == nil {
		returnIfFailWarningInvoker = gi.FunctionInvokerNew("GLib", "return_if_fail_warning")
	}

	var inArgs [3]gi.Argument
	inArgs[0].SetString(logDomain)
	inArgs[1].SetString(prettyFunction)
	inArgs[2].SetString(expression)

	returnIfFailWarningInvoker.Invoke(inArgs[:], nil)

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

var SetApplicationNameFunction *gi.Function
var SetApplicationNameFunctionOnce sync.Once

func SetApplicationNameFunctionSet() {
	SetApplicationNameFunctionOnce.Do(func() {
		SetApplicationNameFunction = gi.FunctionInvokerNew("GLib", "set_application_name")
	})
}

var setApplicationNameInvoker *gi.Function

// SetApplicationName is a representation of the C type g_set_application_name.
func SetApplicationName(applicationName string) {
	if setApplicationNameInvoker == nil {
		setApplicationNameInvoker = gi.FunctionInvokerNew("GLib", "set_application_name")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetString(applicationName)

	setApplicationNameInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_set_error' : parameter 'err' of type 'Error' not supported

// UNSUPPORTED : C value 'g_set_error_literal' : parameter 'err' of type 'Error' not supported

var SetPrgnameFunction *gi.Function
var SetPrgnameFunctionOnce sync.Once

func SetPrgnameFunctionSet() {
	SetPrgnameFunctionOnce.Do(func() {
		SetPrgnameFunction = gi.FunctionInvokerNew("GLib", "set_prgname")
	})
}

var setPrgnameInvoker *gi.Function

// SetPrgname is a representation of the C type g_set_prgname.
func SetPrgname(prgname string) {
	if setPrgnameInvoker == nil {
		setPrgnameInvoker = gi.FunctionInvokerNew("GLib", "set_prgname")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetString(prgname)

	setPrgnameInvoker.Invoke(inArgs[:], nil)

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

// UNSUPPORTED : C value 'g_source_remove' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_source_remove_by_funcs_user_data' : parameter 'funcs' of type 'SourceFuncs' not supported

// UNSUPPORTED : C value 'g_source_remove_by_user_data' : parameter 'user_data' of type 'gpointer' not supported

var SourceSetNameByIdFunction *gi.Function
var SourceSetNameByIdFunctionOnce sync.Once

func SourceSetNameByIdFunctionSet() {
	SourceSetNameByIdFunctionOnce.Do(func() {
		SourceSetNameByIdFunction = gi.FunctionInvokerNew("GLib", "source_set_name_by_id")
	})
}

var sourceSetNameByIdInvoker *gi.Function

// SourceSetNameById is a representation of the C type g_source_set_name_by_id.
func SourceSetNameById(tag uint32, name string) {
	if sourceSetNameByIdInvoker == nil {
		sourceSetNameByIdInvoker = gi.FunctionInvokerNew("GLib", "source_set_name_by_id")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetUint32(tag)
	inArgs[1].SetString(name)

	sourceSetNameByIdInvoker.Invoke(inArgs[:], nil)

}

var SpacedPrimesClosestFunction *gi.Function
var SpacedPrimesClosestFunctionOnce sync.Once

func SpacedPrimesClosestFunctionSet() {
	SpacedPrimesClosestFunctionOnce.Do(func() {
		SpacedPrimesClosestFunction = gi.FunctionInvokerNew("GLib", "spaced_primes_closest")
	})
}

var spacedPrimesClosestInvoker *gi.Function

// SpacedPrimesClosest is a representation of the C type g_spaced_primes_closest.
func SpacedPrimesClosest(num uint32) uint32 {
	if spacedPrimesClosestInvoker == nil {
		spacedPrimesClosestInvoker = gi.FunctionInvokerNew("GLib", "spaced_primes_closest")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetUint32(num)

	ret := spacedPrimesClosestInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Uint32()

	return retGo
}

// UNSUPPORTED : C value 'g_spawn_async' : parameter 'working_directory' of type 'filename' not supported

// UNSUPPORTED : C value 'g_spawn_async_with_fds' : parameter 'working_directory' of type 'filename' not supported

// UNSUPPORTED : C value 'g_spawn_async_with_pipes' : parameter 'working_directory' of type 'filename' not supported

// UNSUPPORTED : C value 'g_spawn_check_exit_status' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_spawn_close_pid' : parameter 'pid' of type 'Pid' not supported

// UNSUPPORTED : C value 'g_spawn_command_line_async' : parameter 'command_line' of type 'filename' not supported

// UNSUPPORTED : C value 'g_spawn_command_line_sync' : parameter 'command_line' of type 'filename' not supported

// UNSUPPORTED : C value 'g_spawn_error_quark' : return type 'Quark' not supported

// UNSUPPORTED : C value 'g_spawn_exit_error_quark' : return type 'Quark' not supported

// UNSUPPORTED : C value 'g_spawn_sync' : parameter 'working_directory' of type 'filename' not supported

// UNSUPPORTED : C value 'g_sprintf' : parameter '...' has no type

var StpcpyFunction *gi.Function
var StpcpyFunctionOnce sync.Once

func StpcpyFunctionSet() {
	StpcpyFunctionOnce.Do(func() {
		StpcpyFunction = gi.FunctionInvokerNew("GLib", "stpcpy")
	})
}

var stpcpyInvoker *gi.Function

// Stpcpy is a representation of the C type g_stpcpy.
func Stpcpy(dest string, src string) string {
	if stpcpyInvoker == nil {
		stpcpyInvoker = gi.FunctionInvokerNew("GLib", "stpcpy")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetString(dest)
	inArgs[1].SetString(src)

	ret := stpcpyInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

// UNSUPPORTED : C value 'g_str_equal' : parameter 'v1' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_str_has_prefix' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_str_has_suffix' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_str_hash' : parameter 'v' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_str_is_ascii' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_str_match_string' : parameter 'accept_alternates' of type 'gboolean' not supported

var StrToAsciiFunction *gi.Function
var StrToAsciiFunctionOnce sync.Once

func StrToAsciiFunctionSet() {
	StrToAsciiFunctionOnce.Do(func() {
		StrToAsciiFunction = gi.FunctionInvokerNew("GLib", "str_to_ascii")
	})
}

var strToAsciiInvoker *gi.Function

// StrToAscii is a representation of the C type g_str_to_ascii.
func StrToAscii(str string, fromLocale string) string {
	if strToAsciiInvoker == nil {
		strToAsciiInvoker = gi.FunctionInvokerNew("GLib", "str_to_ascii")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetString(str)
	inArgs[1].SetString(fromLocale)

	ret := strToAsciiInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

// UNSUPPORTED : C value 'g_str_tokenize_and_fold' : parameter 'ascii_alternates' has no type

var StrcanonFunction *gi.Function
var StrcanonFunctionOnce sync.Once

func StrcanonFunctionSet() {
	StrcanonFunctionOnce.Do(func() {
		StrcanonFunction = gi.FunctionInvokerNew("GLib", "strcanon")
	})
}

var strcanonInvoker *gi.Function

// Strcanon is a representation of the C type g_strcanon.
func Strcanon(string_ string, validChars string, substitutor int8) string {
	if strcanonInvoker == nil {
		strcanonInvoker = gi.FunctionInvokerNew("GLib", "strcanon")
	}

	var inArgs [3]gi.Argument
	inArgs[0].SetString(string_)
	inArgs[1].SetString(validChars)
	inArgs[2].SetInt8(substitutor)

	ret := strcanonInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

var StrcasecmpFunction *gi.Function
var StrcasecmpFunctionOnce sync.Once

func StrcasecmpFunctionSet() {
	StrcasecmpFunctionOnce.Do(func() {
		StrcasecmpFunction = gi.FunctionInvokerNew("GLib", "strcasecmp")
	})
}

var strcasecmpInvoker *gi.Function

// Strcasecmp is a representation of the C type g_strcasecmp.
func Strcasecmp(s1 string, s2 string) int32 {
	if strcasecmpInvoker == nil {
		strcasecmpInvoker = gi.FunctionInvokerNew("GLib", "strcasecmp")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetString(s1)
	inArgs[1].SetString(s2)

	ret := strcasecmpInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var StrchompFunction *gi.Function
var StrchompFunctionOnce sync.Once

func StrchompFunctionSet() {
	StrchompFunctionOnce.Do(func() {
		StrchompFunction = gi.FunctionInvokerNew("GLib", "strchomp")
	})
}

var strchompInvoker *gi.Function

// Strchomp is a representation of the C type g_strchomp.
func Strchomp(string_ string) string {
	if strchompInvoker == nil {
		strchompInvoker = gi.FunctionInvokerNew("GLib", "strchomp")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetString(string_)

	ret := strchompInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

var StrchugFunction *gi.Function
var StrchugFunctionOnce sync.Once

func StrchugFunctionSet() {
	StrchugFunctionOnce.Do(func() {
		StrchugFunction = gi.FunctionInvokerNew("GLib", "strchug")
	})
}

var strchugInvoker *gi.Function

// Strchug is a representation of the C type g_strchug.
func Strchug(string_ string) string {
	if strchugInvoker == nil {
		strchugInvoker = gi.FunctionInvokerNew("GLib", "strchug")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetString(string_)

	ret := strchugInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

var Strcmp0Function *gi.Function
var Strcmp0FunctionOnce sync.Once

func Strcmp0FunctionSet() {
	Strcmp0FunctionOnce.Do(func() {
		Strcmp0Function = gi.FunctionInvokerNew("GLib", "strcmp0")
	})
}

var strcmp0Invoker *gi.Function

// Strcmp0 is a representation of the C type g_strcmp0.
func Strcmp0(str1 string, str2 string) int32 {
	if strcmp0Invoker == nil {
		strcmp0Invoker = gi.FunctionInvokerNew("GLib", "strcmp0")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetString(str1)
	inArgs[1].SetString(str2)

	ret := strcmp0Invoker.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var StrcompressFunction *gi.Function
var StrcompressFunctionOnce sync.Once

func StrcompressFunctionSet() {
	StrcompressFunctionOnce.Do(func() {
		StrcompressFunction = gi.FunctionInvokerNew("GLib", "strcompress")
	})
}

var strcompressInvoker *gi.Function

// Strcompress is a representation of the C type g_strcompress.
func Strcompress(source string) string {
	if strcompressInvoker == nil {
		strcompressInvoker = gi.FunctionInvokerNew("GLib", "strcompress")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetString(source)

	ret := strcompressInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

// UNSUPPORTED : C value 'g_strconcat' : parameter '...' has no type

var StrdelimitFunction *gi.Function
var StrdelimitFunctionOnce sync.Once

func StrdelimitFunctionSet() {
	StrdelimitFunctionOnce.Do(func() {
		StrdelimitFunction = gi.FunctionInvokerNew("GLib", "strdelimit")
	})
}

var strdelimitInvoker *gi.Function

// Strdelimit is a representation of the C type g_strdelimit.
func Strdelimit(string_ string, delimiters string, newDelimiter int8) string {
	if strdelimitInvoker == nil {
		strdelimitInvoker = gi.FunctionInvokerNew("GLib", "strdelimit")
	}

	var inArgs [3]gi.Argument
	inArgs[0].SetString(string_)
	inArgs[1].SetString(delimiters)
	inArgs[2].SetInt8(newDelimiter)

	ret := strdelimitInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

var StrdownFunction *gi.Function
var StrdownFunctionOnce sync.Once

func StrdownFunctionSet() {
	StrdownFunctionOnce.Do(func() {
		StrdownFunction = gi.FunctionInvokerNew("GLib", "strdown")
	})
}

var strdownInvoker *gi.Function

// Strdown is a representation of the C type g_strdown.
func Strdown(string_ string) string {
	if strdownInvoker == nil {
		strdownInvoker = gi.FunctionInvokerNew("GLib", "strdown")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetString(string_)

	ret := strdownInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

var StrdupFunction *gi.Function
var StrdupFunctionOnce sync.Once

func StrdupFunctionSet() {
	StrdupFunctionOnce.Do(func() {
		StrdupFunction = gi.FunctionInvokerNew("GLib", "strdup")
	})
}

var strdupInvoker *gi.Function

// Strdup is a representation of the C type g_strdup.
func Strdup(str string) string {
	if strdupInvoker == nil {
		strdupInvoker = gi.FunctionInvokerNew("GLib", "strdup")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetString(str)

	ret := strdupInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

// UNSUPPORTED : C value 'g_strdup_printf' : parameter '...' has no type

// UNSUPPORTED : C value 'g_strdup_vprintf' : parameter 'args' of type 'va_list' not supported

var StrdupvFunction *gi.Function
var StrdupvFunctionOnce sync.Once

func StrdupvFunctionSet() {
	StrdupvFunctionOnce.Do(func() {
		StrdupvFunction = gi.FunctionInvokerNew("GLib", "strdupv")
	})
}

var strdupvInvoker *gi.Function

// Strdupv is a representation of the C type g_strdupv.
func Strdupv(strArray string) {
	if strdupvInvoker == nil {
		strdupvInvoker = gi.FunctionInvokerNew("GLib", "strdupv")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetString(strArray)

	strdupvInvoker.Invoke(inArgs[:], nil)

}

var StrerrorFunction *gi.Function
var StrerrorFunctionOnce sync.Once

func StrerrorFunctionSet() {
	StrerrorFunctionOnce.Do(func() {
		StrerrorFunction = gi.FunctionInvokerNew("GLib", "strerror")
	})
}

var strerrorInvoker *gi.Function

// Strerror is a representation of the C type g_strerror.
func Strerror(errnum int32) string {
	if strerrorInvoker == nil {
		strerrorInvoker = gi.FunctionInvokerNew("GLib", "strerror")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetInt32(errnum)

	ret := strerrorInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

var StrescapeFunction *gi.Function
var StrescapeFunctionOnce sync.Once

func StrescapeFunctionSet() {
	StrescapeFunctionOnce.Do(func() {
		StrescapeFunction = gi.FunctionInvokerNew("GLib", "strescape")
	})
}

var strescapeInvoker *gi.Function

// Strescape is a representation of the C type g_strescape.
func Strescape(source string, exceptions string) string {
	if strescapeInvoker == nil {
		strescapeInvoker = gi.FunctionInvokerNew("GLib", "strescape")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetString(source)
	inArgs[1].SetString(exceptions)

	ret := strescapeInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

var StrfreevFunction *gi.Function
var StrfreevFunctionOnce sync.Once

func StrfreevFunctionSet() {
	StrfreevFunctionOnce.Do(func() {
		StrfreevFunction = gi.FunctionInvokerNew("GLib", "strfreev")
	})
}

var strfreevInvoker *gi.Function

// Strfreev is a representation of the C type g_strfreev.
func Strfreev(strArray string) {
	if strfreevInvoker == nil {
		strfreevInvoker = gi.FunctionInvokerNew("GLib", "strfreev")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetString(strArray)

	strfreevInvoker.Invoke(inArgs[:], nil)

}

var StringNewFunction *gi.Function
var StringNewFunctionOnce sync.Once

func StringNewFunctionSet() {
	StringNewFunctionOnce.Do(func() {
		StringNewFunction = gi.FunctionInvokerNew("GLib", "string_new")
	})
}

var stringNewInvoker *gi.Function

// StringNew is a representation of the C type g_string_new.
func StringNew(init string) *String {
	if stringNewInvoker == nil {
		stringNewInvoker = gi.FunctionInvokerNew("GLib", "string_new")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetString(init)

	ret := stringNewInvoker.Invoke(inArgs[:], nil)

	retGo := &String{native: ret.Pointer()}

	return retGo
}

var StringNewLenFunction *gi.Function
var StringNewLenFunctionOnce sync.Once

func StringNewLenFunctionSet() {
	StringNewLenFunctionOnce.Do(func() {
		StringNewLenFunction = gi.FunctionInvokerNew("GLib", "string_new_len")
	})
}

var stringNewLenInvoker *gi.Function

// StringNewLen is a representation of the C type g_string_new_len.
func StringNewLen(init string, len int32) *String {
	if stringNewLenInvoker == nil {
		stringNewLenInvoker = gi.FunctionInvokerNew("GLib", "string_new_len")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetString(init)
	inArgs[1].SetInt32(len)

	ret := stringNewLenInvoker.Invoke(inArgs[:], nil)

	retGo := &String{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_string_sized_new' : parameter 'dfl_size' of type 'gsize' not supported

var StripContextFunction *gi.Function
var StripContextFunctionOnce sync.Once

func StripContextFunctionSet() {
	StripContextFunctionOnce.Do(func() {
		StripContextFunction = gi.FunctionInvokerNew("GLib", "strip_context")
	})
}

var stripContextInvoker *gi.Function

// StripContext is a representation of the C type g_strip_context.
func StripContext(msgid string, msgval string) string {
	if stripContextInvoker == nil {
		stripContextInvoker = gi.FunctionInvokerNew("GLib", "strip_context")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetString(msgid)
	inArgs[1].SetString(msgval)

	ret := stripContextInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'g_strjoin' : parameter '...' has no type

var StrjoinvFunction *gi.Function
var StrjoinvFunctionOnce sync.Once

func StrjoinvFunctionSet() {
	StrjoinvFunctionOnce.Do(func() {
		StrjoinvFunction = gi.FunctionInvokerNew("GLib", "strjoinv")
	})
}

var strjoinvInvoker *gi.Function

// Strjoinv is a representation of the C type g_strjoinv.
func Strjoinv(separator string, strArray string) string {
	if strjoinvInvoker == nil {
		strjoinvInvoker = gi.FunctionInvokerNew("GLib", "strjoinv")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetString(separator)
	inArgs[1].SetString(strArray)

	ret := strjoinvInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

// UNSUPPORTED : C value 'g_strlcat' : parameter 'dest_size' of type 'gsize' not supported

// UNSUPPORTED : C value 'g_strlcpy' : parameter 'dest_size' of type 'gsize' not supported

var StrncasecmpFunction *gi.Function
var StrncasecmpFunctionOnce sync.Once

func StrncasecmpFunctionSet() {
	StrncasecmpFunctionOnce.Do(func() {
		StrncasecmpFunction = gi.FunctionInvokerNew("GLib", "strncasecmp")
	})
}

var strncasecmpInvoker *gi.Function

// Strncasecmp is a representation of the C type g_strncasecmp.
func Strncasecmp(s1 string, s2 string, n uint32) int32 {
	if strncasecmpInvoker == nil {
		strncasecmpInvoker = gi.FunctionInvokerNew("GLib", "strncasecmp")
	}

	var inArgs [3]gi.Argument
	inArgs[0].SetString(s1)
	inArgs[1].SetString(s2)
	inArgs[2].SetUint32(n)

	ret := strncasecmpInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

// UNSUPPORTED : C value 'g_strndup' : parameter 'n' of type 'gsize' not supported

// UNSUPPORTED : C value 'g_strnfill' : parameter 'length' of type 'gsize' not supported

var StrreverseFunction *gi.Function
var StrreverseFunctionOnce sync.Once

func StrreverseFunctionSet() {
	StrreverseFunctionOnce.Do(func() {
		StrreverseFunction = gi.FunctionInvokerNew("GLib", "strreverse")
	})
}

var strreverseInvoker *gi.Function

// Strreverse is a representation of the C type g_strreverse.
func Strreverse(string_ string) string {
	if strreverseInvoker == nil {
		strreverseInvoker = gi.FunctionInvokerNew("GLib", "strreverse")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetString(string_)

	ret := strreverseInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

var StrrstrFunction *gi.Function
var StrrstrFunctionOnce sync.Once

func StrrstrFunctionSet() {
	StrrstrFunctionOnce.Do(func() {
		StrrstrFunction = gi.FunctionInvokerNew("GLib", "strrstr")
	})
}

var strrstrInvoker *gi.Function

// Strrstr is a representation of the C type g_strrstr.
func Strrstr(haystack string, needle string) string {
	if strrstrInvoker == nil {
		strrstrInvoker = gi.FunctionInvokerNew("GLib", "strrstr")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetString(haystack)
	inArgs[1].SetString(needle)

	ret := strrstrInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

var StrrstrLenFunction *gi.Function
var StrrstrLenFunctionOnce sync.Once

func StrrstrLenFunctionSet() {
	StrrstrLenFunctionOnce.Do(func() {
		StrrstrLenFunction = gi.FunctionInvokerNew("GLib", "strrstr_len")
	})
}

var strrstrLenInvoker *gi.Function

// StrrstrLen is a representation of the C type g_strrstr_len.
func StrrstrLen(haystack string, haystackLen int32, needle string) string {
	if strrstrLenInvoker == nil {
		strrstrLenInvoker = gi.FunctionInvokerNew("GLib", "strrstr_len")
	}

	var inArgs [3]gi.Argument
	inArgs[0].SetString(haystack)
	inArgs[1].SetInt32(haystackLen)
	inArgs[2].SetString(needle)

	ret := strrstrLenInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

var StrsignalFunction *gi.Function
var StrsignalFunctionOnce sync.Once

func StrsignalFunctionSet() {
	StrsignalFunctionOnce.Do(func() {
		StrsignalFunction = gi.FunctionInvokerNew("GLib", "strsignal")
	})
}

var strsignalInvoker *gi.Function

// Strsignal is a representation of the C type g_strsignal.
func Strsignal(signum int32) string {
	if strsignalInvoker == nil {
		strsignalInvoker = gi.FunctionInvokerNew("GLib", "strsignal")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetInt32(signum)

	ret := strsignalInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

var StrsplitFunction *gi.Function
var StrsplitFunctionOnce sync.Once

func StrsplitFunctionSet() {
	StrsplitFunctionOnce.Do(func() {
		StrsplitFunction = gi.FunctionInvokerNew("GLib", "strsplit")
	})
}

var strsplitInvoker *gi.Function

// Strsplit is a representation of the C type g_strsplit.
func Strsplit(string_ string, delimiter string, maxTokens int32) {
	if strsplitInvoker == nil {
		strsplitInvoker = gi.FunctionInvokerNew("GLib", "strsplit")
	}

	var inArgs [3]gi.Argument
	inArgs[0].SetString(string_)
	inArgs[1].SetString(delimiter)
	inArgs[2].SetInt32(maxTokens)

	strsplitInvoker.Invoke(inArgs[:], nil)

}

var StrsplitSetFunction *gi.Function
var StrsplitSetFunctionOnce sync.Once

func StrsplitSetFunctionSet() {
	StrsplitSetFunctionOnce.Do(func() {
		StrsplitSetFunction = gi.FunctionInvokerNew("GLib", "strsplit_set")
	})
}

var strsplitSetInvoker *gi.Function

// StrsplitSet is a representation of the C type g_strsplit_set.
func StrsplitSet(string_ string, delimiters string, maxTokens int32) {
	if strsplitSetInvoker == nil {
		strsplitSetInvoker = gi.FunctionInvokerNew("GLib", "strsplit_set")
	}

	var inArgs [3]gi.Argument
	inArgs[0].SetString(string_)
	inArgs[1].SetString(delimiters)
	inArgs[2].SetInt32(maxTokens)

	strsplitSetInvoker.Invoke(inArgs[:], nil)

}

var StrstrLenFunction *gi.Function
var StrstrLenFunctionOnce sync.Once

func StrstrLenFunctionSet() {
	StrstrLenFunctionOnce.Do(func() {
		StrstrLenFunction = gi.FunctionInvokerNew("GLib", "strstr_len")
	})
}

var strstrLenInvoker *gi.Function

// StrstrLen is a representation of the C type g_strstr_len.
func StrstrLen(haystack string, haystackLen int32, needle string) string {
	if strstrLenInvoker == nil {
		strstrLenInvoker = gi.FunctionInvokerNew("GLib", "strstr_len")
	}

	var inArgs [3]gi.Argument
	inArgs[0].SetString(haystack)
	inArgs[1].SetInt32(haystackLen)
	inArgs[2].SetString(needle)

	ret := strstrLenInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

// UNSUPPORTED : C value 'g_strtod' : return type 'gdouble' not supported

var StrupFunction *gi.Function
var StrupFunctionOnce sync.Once

func StrupFunctionSet() {
	StrupFunctionOnce.Do(func() {
		StrupFunction = gi.FunctionInvokerNew("GLib", "strup")
	})
}

var strupInvoker *gi.Function

// Strup is a representation of the C type g_strup.
func Strup(string_ string) string {
	if strupInvoker == nil {
		strupInvoker = gi.FunctionInvokerNew("GLib", "strup")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetString(string_)

	ret := strupInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

// UNSUPPORTED : C value 'g_strv_contains' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_strv_equal' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_strv_get_type' : return type 'GType' not supported

var StrvLengthFunction *gi.Function
var StrvLengthFunctionOnce sync.Once

func StrvLengthFunctionSet() {
	StrvLengthFunctionOnce.Do(func() {
		StrvLengthFunction = gi.FunctionInvokerNew("GLib", "strv_length")
	})
}

var strvLengthInvoker *gi.Function

// StrvLength is a representation of the C type g_strv_length.
func StrvLength(strArray string) uint32 {
	if strvLengthInvoker == nil {
		strvLengthInvoker = gi.FunctionInvokerNew("GLib", "strv_length")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetString(strArray)

	ret := strvLengthInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Uint32()

	return retGo
}

// UNSUPPORTED : C value 'g_test_add_data_func' : parameter 'test_data' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_test_add_data_func_full' : parameter 'test_data' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_test_add_func' : parameter 'test_func' of type 'TestFunc' not supported

// UNSUPPORTED : C value 'g_test_add_vtable' : parameter 'data_size' of type 'gsize' not supported

var TestAssertExpectedMessagesInternalFunction *gi.Function
var TestAssertExpectedMessagesInternalFunctionOnce sync.Once

func TestAssertExpectedMessagesInternalFunctionSet() {
	TestAssertExpectedMessagesInternalFunctionOnce.Do(func() {
		TestAssertExpectedMessagesInternalFunction = gi.FunctionInvokerNew("GLib", "test_assert_expected_messages_internal")
	})
}

var testAssertExpectedMessagesInternalInvoker *gi.Function

// TestAssertExpectedMessagesInternal is a representation of the C type g_test_assert_expected_messages_internal.
func TestAssertExpectedMessagesInternal(domain string, file string, line int32, func_ string) {
	if testAssertExpectedMessagesInternalInvoker == nil {
		testAssertExpectedMessagesInternalInvoker = gi.FunctionInvokerNew("GLib", "test_assert_expected_messages_internal")
	}

	var inArgs [4]gi.Argument
	inArgs[0].SetString(domain)
	inArgs[1].SetString(file)
	inArgs[2].SetInt32(line)
	inArgs[3].SetString(func_)

	testAssertExpectedMessagesInternalInvoker.Invoke(inArgs[:], nil)

}

var TestBugFunction *gi.Function
var TestBugFunctionOnce sync.Once

func TestBugFunctionSet() {
	TestBugFunctionOnce.Do(func() {
		TestBugFunction = gi.FunctionInvokerNew("GLib", "test_bug")
	})
}

var testBugInvoker *gi.Function

// TestBug is a representation of the C type g_test_bug.
func TestBug(bugUriSnippet string) {
	if testBugInvoker == nil {
		testBugInvoker = gi.FunctionInvokerNew("GLib", "test_bug")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetString(bugUriSnippet)

	testBugInvoker.Invoke(inArgs[:], nil)

}

var TestBugBaseFunction *gi.Function
var TestBugBaseFunctionOnce sync.Once

func TestBugBaseFunctionSet() {
	TestBugBaseFunctionOnce.Do(func() {
		TestBugBaseFunction = gi.FunctionInvokerNew("GLib", "test_bug_base")
	})
}

var testBugBaseInvoker *gi.Function

// TestBugBase is a representation of the C type g_test_bug_base.
func TestBugBase(uriPattern string) {
	if testBugBaseInvoker == nil {
		testBugBaseInvoker = gi.FunctionInvokerNew("GLib", "test_bug_base")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetString(uriPattern)

	testBugBaseInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_test_build_filename' : parameter 'file_type' of type 'TestFileType' not supported

// UNSUPPORTED : C value 'g_test_create_case' : parameter 'data_size' of type 'gsize' not supported

var TestCreateSuiteFunction *gi.Function
var TestCreateSuiteFunctionOnce sync.Once

func TestCreateSuiteFunctionSet() {
	TestCreateSuiteFunctionOnce.Do(func() {
		TestCreateSuiteFunction = gi.FunctionInvokerNew("GLib", "test_create_suite")
	})
}

var testCreateSuiteInvoker *gi.Function

// TestCreateSuite is a representation of the C type g_test_create_suite.
func TestCreateSuite(suiteName string) *TestSuite {
	if testCreateSuiteInvoker == nil {
		testCreateSuiteInvoker = gi.FunctionInvokerNew("GLib", "test_create_suite")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetString(suiteName)

	ret := testCreateSuiteInvoker.Invoke(inArgs[:], nil)

	retGo := &TestSuite{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_test_expect_message' : parameter 'log_level' of type 'LogLevelFlags' not supported

var TestFailFunction *gi.Function
var TestFailFunctionOnce sync.Once

func TestFailFunctionSet() {
	TestFailFunctionOnce.Do(func() {
		TestFailFunction = gi.FunctionInvokerNew("GLib", "test_fail")
	})
}

var testFailInvoker *gi.Function

// TestFail is a representation of the C type g_test_fail.
func TestFail() {
	if testFailInvoker == nil {
		testFailInvoker = gi.FunctionInvokerNew("GLib", "test_fail")
	}

	testFailInvoker.Invoke(nil, nil)

}

// UNSUPPORTED : C value 'g_test_failed' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_test_get_dir' : parameter 'file_type' of type 'TestFileType' not supported

// UNSUPPORTED : C value 'g_test_get_filename' : parameter 'file_type' of type 'TestFileType' not supported

var TestGetRootFunction *gi.Function
var TestGetRootFunctionOnce sync.Once

func TestGetRootFunctionSet() {
	TestGetRootFunctionOnce.Do(func() {
		TestGetRootFunction = gi.FunctionInvokerNew("GLib", "test_get_root")
	})
}

var testGetRootInvoker *gi.Function

// TestGetRoot is a representation of the C type g_test_get_root.
func TestGetRoot() *TestSuite {
	if testGetRootInvoker == nil {
		testGetRootInvoker = gi.FunctionInvokerNew("GLib", "test_get_root")
	}

	ret := testGetRootInvoker.Invoke(nil, nil)

	retGo := &TestSuite{native: ret.Pointer()}

	return retGo
}

var TestIncompleteFunction *gi.Function
var TestIncompleteFunctionOnce sync.Once

func TestIncompleteFunctionSet() {
	TestIncompleteFunctionOnce.Do(func() {
		TestIncompleteFunction = gi.FunctionInvokerNew("GLib", "test_incomplete")
	})
}

var testIncompleteInvoker *gi.Function

// TestIncomplete is a representation of the C type g_test_incomplete.
func TestIncomplete(msg string) {
	if testIncompleteInvoker == nil {
		testIncompleteInvoker = gi.FunctionInvokerNew("GLib", "test_incomplete")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetString(msg)

	testIncompleteInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_test_init' : parameter '...' has no type

// UNSUPPORTED : C value 'g_test_log_set_fatal_handler' : parameter 'log_func' of type 'TestLogFatalFunc' not supported

// UNSUPPORTED : C value 'g_test_log_type_name' : parameter 'log_type' of type 'TestLogType' not supported

// UNSUPPORTED : C value 'g_test_maximized_result' : parameter 'maximized_quantity' of type 'gdouble' not supported

// UNSUPPORTED : C value 'g_test_message' : parameter '...' has no type

// UNSUPPORTED : C value 'g_test_minimized_result' : parameter 'minimized_quantity' of type 'gdouble' not supported

// UNSUPPORTED : C value 'g_test_queue_destroy' : parameter 'destroy_func' of type 'DestroyNotify' not supported

// UNSUPPORTED : C value 'g_test_queue_free' : parameter 'gfree_pointer' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_test_rand_double' : return type 'gdouble' not supported

// UNSUPPORTED : C value 'g_test_rand_double_range' : parameter 'range_start' of type 'gdouble' not supported

var TestRandIntFunction *gi.Function
var TestRandIntFunctionOnce sync.Once

func TestRandIntFunctionSet() {
	TestRandIntFunctionOnce.Do(func() {
		TestRandIntFunction = gi.FunctionInvokerNew("GLib", "test_rand_int")
	})
}

var testRandIntInvoker *gi.Function

// TestRandInt is a representation of the C type g_test_rand_int.
func TestRandInt() int32 {
	if testRandIntInvoker == nil {
		testRandIntInvoker = gi.FunctionInvokerNew("GLib", "test_rand_int")
	}

	ret := testRandIntInvoker.Invoke(nil, nil)

	retGo := ret.Int32()

	return retGo
}

var TestRandIntRangeFunction *gi.Function
var TestRandIntRangeFunctionOnce sync.Once

func TestRandIntRangeFunctionSet() {
	TestRandIntRangeFunctionOnce.Do(func() {
		TestRandIntRangeFunction = gi.FunctionInvokerNew("GLib", "test_rand_int_range")
	})
}

var testRandIntRangeInvoker *gi.Function

// TestRandIntRange is a representation of the C type g_test_rand_int_range.
func TestRandIntRange(begin int32, end int32) int32 {
	if testRandIntRangeInvoker == nil {
		testRandIntRangeInvoker = gi.FunctionInvokerNew("GLib", "test_rand_int_range")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetInt32(begin)
	inArgs[1].SetInt32(end)

	ret := testRandIntRangeInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var TestRunFunction *gi.Function
var TestRunFunctionOnce sync.Once

func TestRunFunctionSet() {
	TestRunFunctionOnce.Do(func() {
		TestRunFunction = gi.FunctionInvokerNew("GLib", "test_run")
	})
}

var testRunInvoker *gi.Function

// TestRun is a representation of the C type g_test_run.
func TestRun() int32 {
	if testRunInvoker == nil {
		testRunInvoker = gi.FunctionInvokerNew("GLib", "test_run")
	}

	ret := testRunInvoker.Invoke(nil, nil)

	retGo := ret.Int32()

	return retGo
}

// UNSUPPORTED : C value 'g_test_run_suite' : parameter 'suite' of type 'TestSuite' not supported

var TestSetNonfatalAssertionsFunction *gi.Function
var TestSetNonfatalAssertionsFunctionOnce sync.Once

func TestSetNonfatalAssertionsFunctionSet() {
	TestSetNonfatalAssertionsFunctionOnce.Do(func() {
		TestSetNonfatalAssertionsFunction = gi.FunctionInvokerNew("GLib", "test_set_nonfatal_assertions")
	})
}

var testSetNonfatalAssertionsInvoker *gi.Function

// TestSetNonfatalAssertions is a representation of the C type g_test_set_nonfatal_assertions.
func TestSetNonfatalAssertions() {
	if testSetNonfatalAssertionsInvoker == nil {
		testSetNonfatalAssertionsInvoker = gi.FunctionInvokerNew("GLib", "test_set_nonfatal_assertions")
	}

	testSetNonfatalAssertionsInvoker.Invoke(nil, nil)

}

var TestSkipFunction *gi.Function
var TestSkipFunctionOnce sync.Once

func TestSkipFunctionSet() {
	TestSkipFunctionOnce.Do(func() {
		TestSkipFunction = gi.FunctionInvokerNew("GLib", "test_skip")
	})
}

var testSkipInvoker *gi.Function

// TestSkip is a representation of the C type g_test_skip.
func TestSkip(msg string) {
	if testSkipInvoker == nil {
		testSkipInvoker = gi.FunctionInvokerNew("GLib", "test_skip")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetString(msg)

	testSkipInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_test_subprocess' : return type 'gboolean' not supported

var TestSummaryFunction *gi.Function
var TestSummaryFunctionOnce sync.Once

func TestSummaryFunctionSet() {
	TestSummaryFunctionOnce.Do(func() {
		TestSummaryFunction = gi.FunctionInvokerNew("GLib", "test_summary")
	})
}

var testSummaryInvoker *gi.Function

// TestSummary is a representation of the C type g_test_summary.
func TestSummary(summary string) {
	if testSummaryInvoker == nil {
		testSummaryInvoker = gi.FunctionInvokerNew("GLib", "test_summary")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetString(summary)

	testSummaryInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_test_timer_elapsed' : return type 'gdouble' not supported

// UNSUPPORTED : C value 'g_test_timer_last' : return type 'gdouble' not supported

var TestTimerStartFunction *gi.Function
var TestTimerStartFunctionOnce sync.Once

func TestTimerStartFunctionSet() {
	TestTimerStartFunctionOnce.Do(func() {
		TestTimerStartFunction = gi.FunctionInvokerNew("GLib", "test_timer_start")
	})
}

var testTimerStartInvoker *gi.Function

// TestTimerStart is a representation of the C type g_test_timer_start.
func TestTimerStart() {
	if testTimerStartInvoker == nil {
		testTimerStartInvoker = gi.FunctionInvokerNew("GLib", "test_timer_start")
	}

	testTimerStartInvoker.Invoke(nil, nil)

}

var TestTrapAssertionsFunction *gi.Function
var TestTrapAssertionsFunctionOnce sync.Once

func TestTrapAssertionsFunctionSet() {
	TestTrapAssertionsFunctionOnce.Do(func() {
		TestTrapAssertionsFunction = gi.FunctionInvokerNew("GLib", "test_trap_assertions")
	})
}

var testTrapAssertionsInvoker *gi.Function

// TestTrapAssertions is a representation of the C type g_test_trap_assertions.
func TestTrapAssertions(domain string, file string, line int32, func_ string, assertionFlags uint64, pattern string) {
	if testTrapAssertionsInvoker == nil {
		testTrapAssertionsInvoker = gi.FunctionInvokerNew("GLib", "test_trap_assertions")
	}

	var inArgs [6]gi.Argument
	inArgs[0].SetString(domain)
	inArgs[1].SetString(file)
	inArgs[2].SetInt32(line)
	inArgs[3].SetString(func_)
	inArgs[4].SetUint64(assertionFlags)
	inArgs[5].SetString(pattern)

	testTrapAssertionsInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_test_trap_fork' : parameter 'test_trap_flags' of type 'TestTrapFlags' not supported

// UNSUPPORTED : C value 'g_test_trap_has_passed' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_test_trap_reached_timeout' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_test_trap_subprocess' : parameter 'test_flags' of type 'TestSubprocessFlags' not supported

// UNSUPPORTED : C value 'g_thread_error_quark' : return type 'Quark' not supported

// UNSUPPORTED : C value 'g_thread_exit' : parameter 'retval' of type 'gpointer' not supported

var ThreadPoolGetMaxIdleTimeFunction *gi.Function
var ThreadPoolGetMaxIdleTimeFunctionOnce sync.Once

func ThreadPoolGetMaxIdleTimeFunctionSet() {
	ThreadPoolGetMaxIdleTimeFunctionOnce.Do(func() {
		ThreadPoolGetMaxIdleTimeFunction = gi.FunctionInvokerNew("GLib", "thread_pool_get_max_idle_time")
	})
}

var threadPoolGetMaxIdleTimeInvoker *gi.Function

// ThreadPoolGetMaxIdleTime is a representation of the C type g_thread_pool_get_max_idle_time.
func ThreadPoolGetMaxIdleTime() uint32 {
	if threadPoolGetMaxIdleTimeInvoker == nil {
		threadPoolGetMaxIdleTimeInvoker = gi.FunctionInvokerNew("GLib", "thread_pool_get_max_idle_time")
	}

	ret := threadPoolGetMaxIdleTimeInvoker.Invoke(nil, nil)

	retGo := ret.Uint32()

	return retGo
}

var ThreadPoolGetMaxUnusedThreadsFunction *gi.Function
var ThreadPoolGetMaxUnusedThreadsFunctionOnce sync.Once

func ThreadPoolGetMaxUnusedThreadsFunctionSet() {
	ThreadPoolGetMaxUnusedThreadsFunctionOnce.Do(func() {
		ThreadPoolGetMaxUnusedThreadsFunction = gi.FunctionInvokerNew("GLib", "thread_pool_get_max_unused_threads")
	})
}

var threadPoolGetMaxUnusedThreadsInvoker *gi.Function

// ThreadPoolGetMaxUnusedThreads is a representation of the C type g_thread_pool_get_max_unused_threads.
func ThreadPoolGetMaxUnusedThreads() int32 {
	if threadPoolGetMaxUnusedThreadsInvoker == nil {
		threadPoolGetMaxUnusedThreadsInvoker = gi.FunctionInvokerNew("GLib", "thread_pool_get_max_unused_threads")
	}

	ret := threadPoolGetMaxUnusedThreadsInvoker.Invoke(nil, nil)

	retGo := ret.Int32()

	return retGo
}

var ThreadPoolGetNumUnusedThreadsFunction *gi.Function
var ThreadPoolGetNumUnusedThreadsFunctionOnce sync.Once

func ThreadPoolGetNumUnusedThreadsFunctionSet() {
	ThreadPoolGetNumUnusedThreadsFunctionOnce.Do(func() {
		ThreadPoolGetNumUnusedThreadsFunction = gi.FunctionInvokerNew("GLib", "thread_pool_get_num_unused_threads")
	})
}

var threadPoolGetNumUnusedThreadsInvoker *gi.Function

// ThreadPoolGetNumUnusedThreads is a representation of the C type g_thread_pool_get_num_unused_threads.
func ThreadPoolGetNumUnusedThreads() uint32 {
	if threadPoolGetNumUnusedThreadsInvoker == nil {
		threadPoolGetNumUnusedThreadsInvoker = gi.FunctionInvokerNew("GLib", "thread_pool_get_num_unused_threads")
	}

	ret := threadPoolGetNumUnusedThreadsInvoker.Invoke(nil, nil)

	retGo := ret.Uint32()

	return retGo
}

var ThreadPoolSetMaxIdleTimeFunction *gi.Function
var ThreadPoolSetMaxIdleTimeFunctionOnce sync.Once

func ThreadPoolSetMaxIdleTimeFunctionSet() {
	ThreadPoolSetMaxIdleTimeFunctionOnce.Do(func() {
		ThreadPoolSetMaxIdleTimeFunction = gi.FunctionInvokerNew("GLib", "thread_pool_set_max_idle_time")
	})
}

var threadPoolSetMaxIdleTimeInvoker *gi.Function

// ThreadPoolSetMaxIdleTime is a representation of the C type g_thread_pool_set_max_idle_time.
func ThreadPoolSetMaxIdleTime(interval uint32) {
	if threadPoolSetMaxIdleTimeInvoker == nil {
		threadPoolSetMaxIdleTimeInvoker = gi.FunctionInvokerNew("GLib", "thread_pool_set_max_idle_time")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetUint32(interval)

	threadPoolSetMaxIdleTimeInvoker.Invoke(inArgs[:], nil)

}

var ThreadPoolSetMaxUnusedThreadsFunction *gi.Function
var ThreadPoolSetMaxUnusedThreadsFunctionOnce sync.Once

func ThreadPoolSetMaxUnusedThreadsFunctionSet() {
	ThreadPoolSetMaxUnusedThreadsFunctionOnce.Do(func() {
		ThreadPoolSetMaxUnusedThreadsFunction = gi.FunctionInvokerNew("GLib", "thread_pool_set_max_unused_threads")
	})
}

var threadPoolSetMaxUnusedThreadsInvoker *gi.Function

// ThreadPoolSetMaxUnusedThreads is a representation of the C type g_thread_pool_set_max_unused_threads.
func ThreadPoolSetMaxUnusedThreads(maxThreads int32) {
	if threadPoolSetMaxUnusedThreadsInvoker == nil {
		threadPoolSetMaxUnusedThreadsInvoker = gi.FunctionInvokerNew("GLib", "thread_pool_set_max_unused_threads")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetInt32(maxThreads)

	threadPoolSetMaxUnusedThreadsInvoker.Invoke(inArgs[:], nil)

}

var ThreadPoolStopUnusedThreadsFunction *gi.Function
var ThreadPoolStopUnusedThreadsFunctionOnce sync.Once

func ThreadPoolStopUnusedThreadsFunctionSet() {
	ThreadPoolStopUnusedThreadsFunctionOnce.Do(func() {
		ThreadPoolStopUnusedThreadsFunction = gi.FunctionInvokerNew("GLib", "thread_pool_stop_unused_threads")
	})
}

var threadPoolStopUnusedThreadsInvoker *gi.Function

// ThreadPoolStopUnusedThreads is a representation of the C type g_thread_pool_stop_unused_threads.
func ThreadPoolStopUnusedThreads() {
	if threadPoolStopUnusedThreadsInvoker == nil {
		threadPoolStopUnusedThreadsInvoker = gi.FunctionInvokerNew("GLib", "thread_pool_stop_unused_threads")
	}

	threadPoolStopUnusedThreadsInvoker.Invoke(nil, nil)

}

var ThreadSelfFunction *gi.Function
var ThreadSelfFunctionOnce sync.Once

func ThreadSelfFunctionSet() {
	ThreadSelfFunctionOnce.Do(func() {
		ThreadSelfFunction = gi.FunctionInvokerNew("GLib", "thread_self")
	})
}

var threadSelfInvoker *gi.Function

// ThreadSelf is a representation of the C type g_thread_self.
func ThreadSelf() *Thread {
	if threadSelfInvoker == nil {
		threadSelfInvoker = gi.FunctionInvokerNew("GLib", "thread_self")
	}

	ret := threadSelfInvoker.Invoke(nil, nil)

	retGo := &Thread{native: ret.Pointer()}

	return retGo
}

var ThreadYieldFunction *gi.Function
var ThreadYieldFunctionOnce sync.Once

func ThreadYieldFunctionSet() {
	ThreadYieldFunctionOnce.Do(func() {
		ThreadYieldFunction = gi.FunctionInvokerNew("GLib", "thread_yield")
	})
}

var threadYieldInvoker *gi.Function

// ThreadYield is a representation of the C type g_thread_yield.
func ThreadYield() {
	if threadYieldInvoker == nil {
		threadYieldInvoker = gi.FunctionInvokerNew("GLib", "thread_yield")
	}

	threadYieldInvoker.Invoke(nil, nil)

}

// UNSUPPORTED : C value 'g_time_val_from_iso8601' : parameter 'time_' of type 'TimeVal' not supported

// UNSUPPORTED : C value 'g_timeout_add' : parameter 'function' of type 'SourceFunc' not supported

// UNSUPPORTED : C value 'g_timeout_add_full' : parameter 'function' of type 'SourceFunc' not supported

// UNSUPPORTED : C value 'g_timeout_add_seconds' : parameter 'function' of type 'SourceFunc' not supported

// UNSUPPORTED : C value 'g_timeout_add_seconds_full' : parameter 'function' of type 'SourceFunc' not supported

var TimeoutSourceNewFunction *gi.Function
var TimeoutSourceNewFunctionOnce sync.Once

func TimeoutSourceNewFunctionSet() {
	TimeoutSourceNewFunctionOnce.Do(func() {
		TimeoutSourceNewFunction = gi.FunctionInvokerNew("GLib", "timeout_source_new")
	})
}

var timeoutSourceNewInvoker *gi.Function

// TimeoutSourceNew is a representation of the C type g_timeout_source_new.
func TimeoutSourceNew(interval uint32) *Source {
	if timeoutSourceNewInvoker == nil {
		timeoutSourceNewInvoker = gi.FunctionInvokerNew("GLib", "timeout_source_new")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetUint32(interval)

	ret := timeoutSourceNewInvoker.Invoke(inArgs[:], nil)

	retGo := &Source{native: ret.Pointer()}

	return retGo
}

var TimeoutSourceNewSecondsFunction *gi.Function
var TimeoutSourceNewSecondsFunctionOnce sync.Once

func TimeoutSourceNewSecondsFunctionSet() {
	TimeoutSourceNewSecondsFunctionOnce.Do(func() {
		TimeoutSourceNewSecondsFunction = gi.FunctionInvokerNew("GLib", "timeout_source_new_seconds")
	})
}

var timeoutSourceNewSecondsInvoker *gi.Function

// TimeoutSourceNewSeconds is a representation of the C type g_timeout_source_new_seconds.
func TimeoutSourceNewSeconds(interval uint32) *Source {
	if timeoutSourceNewSecondsInvoker == nil {
		timeoutSourceNewSecondsInvoker = gi.FunctionInvokerNew("GLib", "timeout_source_new_seconds")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetUint32(interval)

	ret := timeoutSourceNewSecondsInvoker.Invoke(inArgs[:], nil)

	retGo := &Source{native: ret.Pointer()}

	return retGo
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

// UNSUPPORTED : C value 'g_unix_open_pipe' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_unix_set_fd_nonblocking' : parameter 'nonblock' of type 'gboolean' not supported

// UNSUPPORTED : C value 'g_unix_signal_add' : parameter 'handler' of type 'SourceFunc' not supported

// UNSUPPORTED : C value 'g_unix_signal_add_full' : parameter 'handler' of type 'SourceFunc' not supported

var UnixSignalSourceNewFunction *gi.Function
var UnixSignalSourceNewFunctionOnce sync.Once

func UnixSignalSourceNewFunctionSet() {
	UnixSignalSourceNewFunctionOnce.Do(func() {
		UnixSignalSourceNewFunction = gi.FunctionInvokerNew("GLib", "unix_signal_source_new")
	})
}

var unixSignalSourceNewInvoker *gi.Function

// UnixSignalSourceNew is a representation of the C type g_unix_signal_source_new.
func UnixSignalSourceNew(signum int32) *Source {
	if unixSignalSourceNewInvoker == nil {
		unixSignalSourceNewInvoker = gi.FunctionInvokerNew("GLib", "unix_signal_source_new")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetInt32(signum)

	ret := unixSignalSourceNewInvoker.Invoke(inArgs[:], nil)

	retGo := &Source{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_unlink' : parameter 'filename' of type 'filename' not supported

// UNSUPPORTED : C value 'g_unsetenv' : parameter 'variable' of type 'filename' not supported

// UNSUPPORTED : C value 'g_uri_escape_string' : parameter 'allow_utf8' of type 'gboolean' not supported

var UriListExtractUrisFunction *gi.Function
var UriListExtractUrisFunctionOnce sync.Once

func UriListExtractUrisFunctionSet() {
	UriListExtractUrisFunctionOnce.Do(func() {
		UriListExtractUrisFunction = gi.FunctionInvokerNew("GLib", "uri_list_extract_uris")
	})
}

var uriListExtractUrisInvoker *gi.Function

// UriListExtractUris is a representation of the C type g_uri_list_extract_uris.
func UriListExtractUris(uriList string) {
	if uriListExtractUrisInvoker == nil {
		uriListExtractUrisInvoker = gi.FunctionInvokerNew("GLib", "uri_list_extract_uris")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetString(uriList)

	uriListExtractUrisInvoker.Invoke(inArgs[:], nil)

}

var UriParseSchemeFunction *gi.Function
var UriParseSchemeFunctionOnce sync.Once

func UriParseSchemeFunctionSet() {
	UriParseSchemeFunctionOnce.Do(func() {
		UriParseSchemeFunction = gi.FunctionInvokerNew("GLib", "uri_parse_scheme")
	})
}

var uriParseSchemeInvoker *gi.Function

// UriParseScheme is a representation of the C type g_uri_parse_scheme.
func UriParseScheme(uri string) string {
	if uriParseSchemeInvoker == nil {
		uriParseSchemeInvoker = gi.FunctionInvokerNew("GLib", "uri_parse_scheme")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetString(uri)

	ret := uriParseSchemeInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

var UriUnescapeSegmentFunction *gi.Function
var UriUnescapeSegmentFunctionOnce sync.Once

func UriUnescapeSegmentFunctionSet() {
	UriUnescapeSegmentFunctionOnce.Do(func() {
		UriUnescapeSegmentFunction = gi.FunctionInvokerNew("GLib", "uri_unescape_segment")
	})
}

var uriUnescapeSegmentInvoker *gi.Function

// UriUnescapeSegment is a representation of the C type g_uri_unescape_segment.
func UriUnescapeSegment(escapedString string, escapedStringEnd string, illegalCharacters string) string {
	if uriUnescapeSegmentInvoker == nil {
		uriUnescapeSegmentInvoker = gi.FunctionInvokerNew("GLib", "uri_unescape_segment")
	}

	var inArgs [3]gi.Argument
	inArgs[0].SetString(escapedString)
	inArgs[1].SetString(escapedStringEnd)
	inArgs[2].SetString(illegalCharacters)

	ret := uriUnescapeSegmentInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

var UriUnescapeStringFunction *gi.Function
var UriUnescapeStringFunctionOnce sync.Once

func UriUnescapeStringFunctionSet() {
	UriUnescapeStringFunctionOnce.Do(func() {
		UriUnescapeStringFunction = gi.FunctionInvokerNew("GLib", "uri_unescape_string")
	})
}

var uriUnescapeStringInvoker *gi.Function

// UriUnescapeString is a representation of the C type g_uri_unescape_string.
func UriUnescapeString(escapedString string, illegalCharacters string) string {
	if uriUnescapeStringInvoker == nil {
		uriUnescapeStringInvoker = gi.FunctionInvokerNew("GLib", "uri_unescape_string")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetString(escapedString)
	inArgs[1].SetString(illegalCharacters)

	ret := uriUnescapeStringInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

var UsleepFunction *gi.Function
var UsleepFunctionOnce sync.Once

func UsleepFunctionSet() {
	UsleepFunctionOnce.Do(func() {
		UsleepFunction = gi.FunctionInvokerNew("GLib", "usleep")
	})
}

var usleepInvoker *gi.Function

// Usleep is a representation of the C type g_usleep.
func Usleep(microseconds uint64) {
	if usleepInvoker == nil {
		usleepInvoker = gi.FunctionInvokerNew("GLib", "usleep")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetUint64(microseconds)

	usleepInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_utf16_to_ucs4' : return type 'gunichar' not supported

var Utf16ToUtf8Function *gi.Function
var Utf16ToUtf8FunctionOnce sync.Once

func Utf16ToUtf8FunctionSet() {
	Utf16ToUtf8FunctionOnce.Do(func() {
		Utf16ToUtf8Function = gi.FunctionInvokerNew("GLib", "utf16_to_utf8")
	})
}

var utf16ToUtf8Invoker *gi.Function

// Utf16ToUtf8 is a representation of the C type g_utf16_to_utf8.
func Utf16ToUtf8(str uint16, len int64) (string, int64, int64) {
	if utf16ToUtf8Invoker == nil {
		utf16ToUtf8Invoker = gi.FunctionInvokerNew("GLib", "utf16_to_utf8")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetUint16(str)
	inArgs[1].SetInt64(len)

	var outArgs [2]gi.Argument

	ret := utf16ToUtf8Invoker.Invoke(inArgs[:], outArgs[:])

	retGo := ret.String(true)
	out0 := outArgs[0].Int64()
	out1 := outArgs[1].Int64()

	return retGo, out0, out1
}

var Utf8CasefoldFunction *gi.Function
var Utf8CasefoldFunctionOnce sync.Once

func Utf8CasefoldFunctionSet() {
	Utf8CasefoldFunctionOnce.Do(func() {
		Utf8CasefoldFunction = gi.FunctionInvokerNew("GLib", "utf8_casefold")
	})
}

var utf8CasefoldInvoker *gi.Function

// Utf8Casefold is a representation of the C type g_utf8_casefold.
func Utf8Casefold(str string, len int32) string {
	if utf8CasefoldInvoker == nil {
		utf8CasefoldInvoker = gi.FunctionInvokerNew("GLib", "utf8_casefold")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetString(str)
	inArgs[1].SetInt32(len)

	ret := utf8CasefoldInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

var Utf8CollateFunction *gi.Function
var Utf8CollateFunctionOnce sync.Once

func Utf8CollateFunctionSet() {
	Utf8CollateFunctionOnce.Do(func() {
		Utf8CollateFunction = gi.FunctionInvokerNew("GLib", "utf8_collate")
	})
}

var utf8CollateInvoker *gi.Function

// Utf8Collate is a representation of the C type g_utf8_collate.
func Utf8Collate(str1 string, str2 string) int32 {
	if utf8CollateInvoker == nil {
		utf8CollateInvoker = gi.FunctionInvokerNew("GLib", "utf8_collate")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetString(str1)
	inArgs[1].SetString(str2)

	ret := utf8CollateInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var Utf8CollateKeyFunction *gi.Function
var Utf8CollateKeyFunctionOnce sync.Once

func Utf8CollateKeyFunctionSet() {
	Utf8CollateKeyFunctionOnce.Do(func() {
		Utf8CollateKeyFunction = gi.FunctionInvokerNew("GLib", "utf8_collate_key")
	})
}

var utf8CollateKeyInvoker *gi.Function

// Utf8CollateKey is a representation of the C type g_utf8_collate_key.
func Utf8CollateKey(str string, len int32) string {
	if utf8CollateKeyInvoker == nil {
		utf8CollateKeyInvoker = gi.FunctionInvokerNew("GLib", "utf8_collate_key")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetString(str)
	inArgs[1].SetInt32(len)

	ret := utf8CollateKeyInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

var Utf8CollateKeyForFilenameFunction *gi.Function
var Utf8CollateKeyForFilenameFunctionOnce sync.Once

func Utf8CollateKeyForFilenameFunctionSet() {
	Utf8CollateKeyForFilenameFunctionOnce.Do(func() {
		Utf8CollateKeyForFilenameFunction = gi.FunctionInvokerNew("GLib", "utf8_collate_key_for_filename")
	})
}

var utf8CollateKeyForFilenameInvoker *gi.Function

// Utf8CollateKeyForFilename is a representation of the C type g_utf8_collate_key_for_filename.
func Utf8CollateKeyForFilename(str string, len int32) string {
	if utf8CollateKeyForFilenameInvoker == nil {
		utf8CollateKeyForFilenameInvoker = gi.FunctionInvokerNew("GLib", "utf8_collate_key_for_filename")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetString(str)
	inArgs[1].SetInt32(len)

	ret := utf8CollateKeyForFilenameInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

var Utf8FindNextCharFunction *gi.Function
var Utf8FindNextCharFunctionOnce sync.Once

func Utf8FindNextCharFunctionSet() {
	Utf8FindNextCharFunctionOnce.Do(func() {
		Utf8FindNextCharFunction = gi.FunctionInvokerNew("GLib", "utf8_find_next_char")
	})
}

var utf8FindNextCharInvoker *gi.Function

// Utf8FindNextChar is a representation of the C type g_utf8_find_next_char.
func Utf8FindNextChar(p string, end string) string {
	if utf8FindNextCharInvoker == nil {
		utf8FindNextCharInvoker = gi.FunctionInvokerNew("GLib", "utf8_find_next_char")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetString(p)
	inArgs[1].SetString(end)

	ret := utf8FindNextCharInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

var Utf8FindPrevCharFunction *gi.Function
var Utf8FindPrevCharFunctionOnce sync.Once

func Utf8FindPrevCharFunctionSet() {
	Utf8FindPrevCharFunctionOnce.Do(func() {
		Utf8FindPrevCharFunction = gi.FunctionInvokerNew("GLib", "utf8_find_prev_char")
	})
}

var utf8FindPrevCharInvoker *gi.Function

// Utf8FindPrevChar is a representation of the C type g_utf8_find_prev_char.
func Utf8FindPrevChar(str string, p string) string {
	if utf8FindPrevCharInvoker == nil {
		utf8FindPrevCharInvoker = gi.FunctionInvokerNew("GLib", "utf8_find_prev_char")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetString(str)
	inArgs[1].SetString(p)

	ret := utf8FindPrevCharInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'g_utf8_get_char' : return type 'gunichar' not supported

// UNSUPPORTED : C value 'g_utf8_get_char_validated' : return type 'gunichar' not supported

var Utf8MakeValidFunction *gi.Function
var Utf8MakeValidFunctionOnce sync.Once

func Utf8MakeValidFunctionSet() {
	Utf8MakeValidFunctionOnce.Do(func() {
		Utf8MakeValidFunction = gi.FunctionInvokerNew("GLib", "utf8_make_valid")
	})
}

var utf8MakeValidInvoker *gi.Function

// Utf8MakeValid is a representation of the C type g_utf8_make_valid.
func Utf8MakeValid(str string, len int32) string {
	if utf8MakeValidInvoker == nil {
		utf8MakeValidInvoker = gi.FunctionInvokerNew("GLib", "utf8_make_valid")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetString(str)
	inArgs[1].SetInt32(len)

	ret := utf8MakeValidInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

// UNSUPPORTED : C value 'g_utf8_normalize' : parameter 'mode' of type 'NormalizeMode' not supported

var Utf8OffsetToPointerFunction *gi.Function
var Utf8OffsetToPointerFunctionOnce sync.Once

func Utf8OffsetToPointerFunctionSet() {
	Utf8OffsetToPointerFunctionOnce.Do(func() {
		Utf8OffsetToPointerFunction = gi.FunctionInvokerNew("GLib", "utf8_offset_to_pointer")
	})
}

var utf8OffsetToPointerInvoker *gi.Function

// Utf8OffsetToPointer is a representation of the C type g_utf8_offset_to_pointer.
func Utf8OffsetToPointer(str string, offset int64) string {
	if utf8OffsetToPointerInvoker == nil {
		utf8OffsetToPointerInvoker = gi.FunctionInvokerNew("GLib", "utf8_offset_to_pointer")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetString(str)
	inArgs[1].SetInt64(offset)

	ret := utf8OffsetToPointerInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

var Utf8PointerToOffsetFunction *gi.Function
var Utf8PointerToOffsetFunctionOnce sync.Once

func Utf8PointerToOffsetFunctionSet() {
	Utf8PointerToOffsetFunctionOnce.Do(func() {
		Utf8PointerToOffsetFunction = gi.FunctionInvokerNew("GLib", "utf8_pointer_to_offset")
	})
}

var utf8PointerToOffsetInvoker *gi.Function

// Utf8PointerToOffset is a representation of the C type g_utf8_pointer_to_offset.
func Utf8PointerToOffset(str string, pos string) int64 {
	if utf8PointerToOffsetInvoker == nil {
		utf8PointerToOffsetInvoker = gi.FunctionInvokerNew("GLib", "utf8_pointer_to_offset")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetString(str)
	inArgs[1].SetString(pos)

	ret := utf8PointerToOffsetInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int64()

	return retGo
}

var Utf8PrevCharFunction *gi.Function
var Utf8PrevCharFunctionOnce sync.Once

func Utf8PrevCharFunctionSet() {
	Utf8PrevCharFunctionOnce.Do(func() {
		Utf8PrevCharFunction = gi.FunctionInvokerNew("GLib", "utf8_prev_char")
	})
}

var utf8PrevCharInvoker *gi.Function

// Utf8PrevChar is a representation of the C type g_utf8_prev_char.
func Utf8PrevChar(p string) string {
	if utf8PrevCharInvoker == nil {
		utf8PrevCharInvoker = gi.FunctionInvokerNew("GLib", "utf8_prev_char")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetString(p)

	ret := utf8PrevCharInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'g_utf8_strchr' : parameter 'c' of type 'gunichar' not supported

var Utf8StrdownFunction *gi.Function
var Utf8StrdownFunctionOnce sync.Once

func Utf8StrdownFunctionSet() {
	Utf8StrdownFunctionOnce.Do(func() {
		Utf8StrdownFunction = gi.FunctionInvokerNew("GLib", "utf8_strdown")
	})
}

var utf8StrdownInvoker *gi.Function

// Utf8Strdown is a representation of the C type g_utf8_strdown.
func Utf8Strdown(str string, len int32) string {
	if utf8StrdownInvoker == nil {
		utf8StrdownInvoker = gi.FunctionInvokerNew("GLib", "utf8_strdown")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetString(str)
	inArgs[1].SetInt32(len)

	ret := utf8StrdownInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

var Utf8StrlenFunction *gi.Function
var Utf8StrlenFunctionOnce sync.Once

func Utf8StrlenFunctionSet() {
	Utf8StrlenFunctionOnce.Do(func() {
		Utf8StrlenFunction = gi.FunctionInvokerNew("GLib", "utf8_strlen")
	})
}

var utf8StrlenInvoker *gi.Function

// Utf8Strlen is a representation of the C type g_utf8_strlen.
func Utf8Strlen(p string, max int32) int64 {
	if utf8StrlenInvoker == nil {
		utf8StrlenInvoker = gi.FunctionInvokerNew("GLib", "utf8_strlen")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetString(p)
	inArgs[1].SetInt32(max)

	ret := utf8StrlenInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int64()

	return retGo
}

// UNSUPPORTED : C value 'g_utf8_strncpy' : parameter 'n' of type 'gsize' not supported

// UNSUPPORTED : C value 'g_utf8_strrchr' : parameter 'c' of type 'gunichar' not supported

var Utf8StrreverseFunction *gi.Function
var Utf8StrreverseFunctionOnce sync.Once

func Utf8StrreverseFunctionSet() {
	Utf8StrreverseFunctionOnce.Do(func() {
		Utf8StrreverseFunction = gi.FunctionInvokerNew("GLib", "utf8_strreverse")
	})
}

var utf8StrreverseInvoker *gi.Function

// Utf8Strreverse is a representation of the C type g_utf8_strreverse.
func Utf8Strreverse(str string, len int32) string {
	if utf8StrreverseInvoker == nil {
		utf8StrreverseInvoker = gi.FunctionInvokerNew("GLib", "utf8_strreverse")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetString(str)
	inArgs[1].SetInt32(len)

	ret := utf8StrreverseInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

var Utf8StrupFunction *gi.Function
var Utf8StrupFunctionOnce sync.Once

func Utf8StrupFunctionSet() {
	Utf8StrupFunctionOnce.Do(func() {
		Utf8StrupFunction = gi.FunctionInvokerNew("GLib", "utf8_strup")
	})
}

var utf8StrupInvoker *gi.Function

// Utf8Strup is a representation of the C type g_utf8_strup.
func Utf8Strup(str string, len int32) string {
	if utf8StrupInvoker == nil {
		utf8StrupInvoker = gi.FunctionInvokerNew("GLib", "utf8_strup")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetString(str)
	inArgs[1].SetInt32(len)

	ret := utf8StrupInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

var Utf8SubstringFunction *gi.Function
var Utf8SubstringFunctionOnce sync.Once

func Utf8SubstringFunctionSet() {
	Utf8SubstringFunctionOnce.Do(func() {
		Utf8SubstringFunction = gi.FunctionInvokerNew("GLib", "utf8_substring")
	})
}

var utf8SubstringInvoker *gi.Function

// Utf8Substring is a representation of the C type g_utf8_substring.
func Utf8Substring(str string, startPos int64, endPos int64) string {
	if utf8SubstringInvoker == nil {
		utf8SubstringInvoker = gi.FunctionInvokerNew("GLib", "utf8_substring")
	}

	var inArgs [3]gi.Argument
	inArgs[0].SetString(str)
	inArgs[1].SetInt64(startPos)
	inArgs[2].SetInt64(endPos)

	ret := utf8SubstringInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

// UNSUPPORTED : C value 'g_utf8_to_ucs4' : return type 'gunichar' not supported

// UNSUPPORTED : C value 'g_utf8_to_ucs4_fast' : return type 'gunichar' not supported

var Utf8ToUtf16Function *gi.Function
var Utf8ToUtf16FunctionOnce sync.Once

func Utf8ToUtf16FunctionSet() {
	Utf8ToUtf16FunctionOnce.Do(func() {
		Utf8ToUtf16Function = gi.FunctionInvokerNew("GLib", "utf8_to_utf16")
	})
}

var utf8ToUtf16Invoker *gi.Function

// Utf8ToUtf16 is a representation of the C type g_utf8_to_utf16.
func Utf8ToUtf16(str string, len int64) (uint16, int64, int64) {
	if utf8ToUtf16Invoker == nil {
		utf8ToUtf16Invoker = gi.FunctionInvokerNew("GLib", "utf8_to_utf16")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetString(str)
	inArgs[1].SetInt64(len)

	var outArgs [2]gi.Argument

	ret := utf8ToUtf16Invoker.Invoke(inArgs[:], outArgs[:])

	retGo := ret.Uint16()
	out0 := outArgs[0].Int64()
	out1 := outArgs[1].Int64()

	return retGo, out0, out1
}

// UNSUPPORTED : C value 'g_utf8_validate' : parameter 'str' has no type

// UNSUPPORTED : C value 'g_utf8_validate_len' : parameter 'str' has no type

// UNSUPPORTED : C value 'g_uuid_string_is_valid' : return type 'gboolean' not supported

var UuidStringRandomFunction *gi.Function
var UuidStringRandomFunctionOnce sync.Once

func UuidStringRandomFunctionSet() {
	UuidStringRandomFunctionOnce.Do(func() {
		UuidStringRandomFunction = gi.FunctionInvokerNew("GLib", "uuid_string_random")
	})
}

var uuidStringRandomInvoker *gi.Function

// UuidStringRandom is a representation of the C type g_uuid_string_random.
func UuidStringRandom() string {
	if uuidStringRandomInvoker == nil {
		uuidStringRandomInvoker = gi.FunctionInvokerNew("GLib", "uuid_string_random")
	}

	ret := uuidStringRandomInvoker.Invoke(nil, nil)

	retGo := ret.String(true)

	return retGo
}

// UNSUPPORTED : C value 'g_variant_get_gtype' : return type 'GType' not supported

// UNSUPPORTED : C value 'g_variant_is_object_path' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_variant_is_signature' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_variant_parse' : parameter 'type' of type 'VariantType' not supported

// UNSUPPORTED : C value 'g_variant_parse_error_print_context' : parameter 'error' of type 'Error' not supported

// UNSUPPORTED : C value 'g_variant_parse_error_quark' : return type 'Quark' not supported

// UNSUPPORTED : C value 'g_variant_parser_get_error_quark' : return type 'Quark' not supported

var VariantTypeCheckedFunction *gi.Function
var VariantTypeCheckedFunctionOnce sync.Once

func VariantTypeCheckedFunctionSet() {
	VariantTypeCheckedFunctionOnce.Do(func() {
		VariantTypeCheckedFunction = gi.FunctionInvokerNew("GLib", "variant_type_checked_")
	})
}

var variantTypeCheckedInvoker *gi.Function

// VariantTypeChecked is a representation of the C type g_variant_type_checked_.
func VariantTypeChecked(arg0 string) *VariantType {
	if variantTypeCheckedInvoker == nil {
		variantTypeCheckedInvoker = gi.FunctionInvokerNew("GLib", "variant_type_checked_")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetString(arg0)

	ret := variantTypeCheckedInvoker.Invoke(inArgs[:], nil)

	retGo := &VariantType{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_variant_type_string_get_depth_' : return type 'gsize' not supported

// UNSUPPORTED : C value 'g_variant_type_string_is_valid' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_variant_type_string_scan' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_vasprintf' : parameter 'args' of type 'va_list' not supported

// UNSUPPORTED : C value 'g_vfprintf' : parameter 'file' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_vprintf' : parameter 'args' of type 'va_list' not supported

// UNSUPPORTED : C value 'g_vsnprintf' : parameter 'args' of type 'va_list' not supported

// UNSUPPORTED : C value 'g_vsprintf' : parameter 'args' of type 'va_list' not supported

var WarnMessageFunction *gi.Function
var WarnMessageFunctionOnce sync.Once

func WarnMessageFunctionSet() {
	WarnMessageFunctionOnce.Do(func() {
		WarnMessageFunction = gi.FunctionInvokerNew("GLib", "warn_message")
	})
}

var warnMessageInvoker *gi.Function

// WarnMessage is a representation of the C type g_warn_message.
func WarnMessage(domain string, file string, line int32, func_ string, warnexpr string) {
	if warnMessageInvoker == nil {
		warnMessageInvoker = gi.FunctionInvokerNew("GLib", "warn_message")
	}

	var inArgs [5]gi.Argument
	inArgs[0].SetString(domain)
	inArgs[1].SetString(file)
	inArgs[2].SetInt32(line)
	inArgs[3].SetString(func_)
	inArgs[4].SetString(warnexpr)

	warnMessageInvoker.Invoke(inArgs[:], nil)

}
