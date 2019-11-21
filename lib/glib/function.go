// Code generated - DO NOT EDIT.

package glib

import (
	gi "github.com/pekim/gobbi/internal/gi"
	"sync"
)

// UNSUPPORTED : C value 'g_access' : parameter 'filename' of type 'filename' not supported

var asciiDigitValueFunction *gi.Function
var asciiDigitValueFunction_Once sync.Once

func asciiDigitValueFunction_Set() {
	asciiDigitValueFunction_Once.Do(func() {
		asciiDigitValueFunction = gi.FunctionInvokerNew("GLib", "ascii_digit_value")
	})
}

// AsciiDigitValue is a representation of the C type g_ascii_digit_value.
func AsciiDigitValue(c int8) int32 {
	asciiDigitValueFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetInt8(c)

	ret := asciiDigitValueFunction.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

// UNSUPPORTED : C value 'g_ascii_dtostr' : parameter 'd' of type 'gdouble' not supported

// UNSUPPORTED : C value 'g_ascii_formatd' : parameter 'd' of type 'gdouble' not supported

var asciiStrcasecmpFunction *gi.Function
var asciiStrcasecmpFunction_Once sync.Once

func asciiStrcasecmpFunction_Set() {
	asciiStrcasecmpFunction_Once.Do(func() {
		asciiStrcasecmpFunction = gi.FunctionInvokerNew("GLib", "ascii_strcasecmp")
	})
}

// AsciiStrcasecmp is a representation of the C type g_ascii_strcasecmp.
func AsciiStrcasecmp(s1 string, s2 string) int32 {
	asciiStrcasecmpFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetString(s1)
	inArgs[1].SetString(s2)

	ret := asciiStrcasecmpFunction.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var asciiStrdownFunction *gi.Function
var asciiStrdownFunction_Once sync.Once

func asciiStrdownFunction_Set() {
	asciiStrdownFunction_Once.Do(func() {
		asciiStrdownFunction = gi.FunctionInvokerNew("GLib", "ascii_strdown")
	})
}

// AsciiStrdown is a representation of the C type g_ascii_strdown.
func AsciiStrdown(str string, len int32) string {
	asciiStrdownFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetString(str)
	inArgs[1].SetInt32(len)

	ret := asciiStrdownFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

// UNSUPPORTED : C value 'g_ascii_string_to_signed' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_ascii_string_to_unsigned' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_ascii_strncasecmp' : parameter 'n' of type 'gsize' not supported

// UNSUPPORTED : C value 'g_ascii_strtod' : return type 'gdouble' not supported

var asciiStrtollFunction *gi.Function
var asciiStrtollFunction_Once sync.Once

func asciiStrtollFunction_Set() {
	asciiStrtollFunction_Once.Do(func() {
		asciiStrtollFunction = gi.FunctionInvokerNew("GLib", "ascii_strtoll")
	})
}

// AsciiStrtoll is a representation of the C type g_ascii_strtoll.
func AsciiStrtoll(nptr string, base uint32) (int64, string) {
	asciiStrtollFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetString(nptr)
	inArgs[1].SetUint32(base)

	var outArgs [1]gi.Argument

	ret := asciiStrtollFunction.Invoke(inArgs[:], outArgs[:])

	retGo := ret.Int64()
	out0 := outArgs[0].String(false)

	return retGo, out0
}

var asciiStrtoullFunction *gi.Function
var asciiStrtoullFunction_Once sync.Once

func asciiStrtoullFunction_Set() {
	asciiStrtoullFunction_Once.Do(func() {
		asciiStrtoullFunction = gi.FunctionInvokerNew("GLib", "ascii_strtoull")
	})
}

// AsciiStrtoull is a representation of the C type g_ascii_strtoull.
func AsciiStrtoull(nptr string, base uint32) (uint64, string) {
	asciiStrtoullFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetString(nptr)
	inArgs[1].SetUint32(base)

	var outArgs [1]gi.Argument

	ret := asciiStrtoullFunction.Invoke(inArgs[:], outArgs[:])

	retGo := ret.Uint64()
	out0 := outArgs[0].String(false)

	return retGo, out0
}

var asciiStrupFunction *gi.Function
var asciiStrupFunction_Once sync.Once

func asciiStrupFunction_Set() {
	asciiStrupFunction_Once.Do(func() {
		asciiStrupFunction = gi.FunctionInvokerNew("GLib", "ascii_strup")
	})
}

// AsciiStrup is a representation of the C type g_ascii_strup.
func AsciiStrup(str string, len int32) string {
	asciiStrupFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetString(str)
	inArgs[1].SetInt32(len)

	ret := asciiStrupFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

var asciiTolowerFunction *gi.Function
var asciiTolowerFunction_Once sync.Once

func asciiTolowerFunction_Set() {
	asciiTolowerFunction_Once.Do(func() {
		asciiTolowerFunction = gi.FunctionInvokerNew("GLib", "ascii_tolower")
	})
}

// AsciiTolower is a representation of the C type g_ascii_tolower.
func AsciiTolower(c int8) int8 {
	asciiTolowerFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetInt8(c)

	ret := asciiTolowerFunction.Invoke(inArgs[:], nil)

	retGo := ret.Int8()

	return retGo
}

var asciiToupperFunction *gi.Function
var asciiToupperFunction_Once sync.Once

func asciiToupperFunction_Set() {
	asciiToupperFunction_Once.Do(func() {
		asciiToupperFunction = gi.FunctionInvokerNew("GLib", "ascii_toupper")
	})
}

// AsciiToupper is a representation of the C type g_ascii_toupper.
func AsciiToupper(c int8) int8 {
	asciiToupperFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetInt8(c)

	ret := asciiToupperFunction.Invoke(inArgs[:], nil)

	retGo := ret.Int8()

	return retGo
}

var asciiXdigitValueFunction *gi.Function
var asciiXdigitValueFunction_Once sync.Once

func asciiXdigitValueFunction_Set() {
	asciiXdigitValueFunction_Once.Do(func() {
		asciiXdigitValueFunction = gi.FunctionInvokerNew("GLib", "ascii_xdigit_value")
	})
}

// AsciiXdigitValue is a representation of the C type g_ascii_xdigit_value.
func AsciiXdigitValue(c int8) int32 {
	asciiXdigitValueFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetInt8(c)

	ret := asciiXdigitValueFunction.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var assertWarningFunction *gi.Function
var assertWarningFunction_Once sync.Once

func assertWarningFunction_Set() {
	assertWarningFunction_Once.Do(func() {
		assertWarningFunction = gi.FunctionInvokerNew("GLib", "assert_warning")
	})
}

// AssertWarning is a representation of the C type g_assert_warning.
func AssertWarning(logDomain string, file string, line int32, prettyFunction string, expression string) {
	assertWarningFunction_Set()

	var inArgs [5]gi.Argument
	inArgs[0].SetString(logDomain)
	inArgs[1].SetString(file)
	inArgs[2].SetInt32(line)
	inArgs[3].SetString(prettyFunction)
	inArgs[4].SetString(expression)

	assertWarningFunction.Invoke(inArgs[:], nil)

}

var assertionMessageFunction *gi.Function
var assertionMessageFunction_Once sync.Once

func assertionMessageFunction_Set() {
	assertionMessageFunction_Once.Do(func() {
		assertionMessageFunction = gi.FunctionInvokerNew("GLib", "assertion_message")
	})
}

// AssertionMessage is a representation of the C type g_assertion_message.
func AssertionMessage(domain string, file string, line int32, func_ string, message string) {
	assertionMessageFunction_Set()

	var inArgs [5]gi.Argument
	inArgs[0].SetString(domain)
	inArgs[1].SetString(file)
	inArgs[2].SetInt32(line)
	inArgs[3].SetString(func_)
	inArgs[4].SetString(message)

	assertionMessageFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_assertion_message_cmpnum' : parameter 'arg1' of type 'long double' not supported

var assertionMessageCmpstrFunction *gi.Function
var assertionMessageCmpstrFunction_Once sync.Once

func assertionMessageCmpstrFunction_Set() {
	assertionMessageCmpstrFunction_Once.Do(func() {
		assertionMessageCmpstrFunction = gi.FunctionInvokerNew("GLib", "assertion_message_cmpstr")
	})
}

// AssertionMessageCmpstr is a representation of the C type g_assertion_message_cmpstr.
func AssertionMessageCmpstr(domain string, file string, line int32, func_ string, expr string, arg1 string, cmp string, arg2 string) {
	assertionMessageCmpstrFunction_Set()

	var inArgs [8]gi.Argument
	inArgs[0].SetString(domain)
	inArgs[1].SetString(file)
	inArgs[2].SetInt32(line)
	inArgs[3].SetString(func_)
	inArgs[4].SetString(expr)
	inArgs[5].SetString(arg1)
	inArgs[6].SetString(cmp)
	inArgs[7].SetString(arg2)

	assertionMessageCmpstrFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_assertion_message_error' : parameter 'error' of type 'Error' not supported

var assertionMessageExprFunction *gi.Function
var assertionMessageExprFunction_Once sync.Once

func assertionMessageExprFunction_Set() {
	assertionMessageExprFunction_Once.Do(func() {
		assertionMessageExprFunction = gi.FunctionInvokerNew("GLib", "assertion_message_expr")
	})
}

// AssertionMessageExpr is a representation of the C type g_assertion_message_expr.
func AssertionMessageExpr(domain string, file string, line int32, func_ string, expr string) {
	assertionMessageExprFunction_Set()

	var inArgs [5]gi.Argument
	inArgs[0].SetString(domain)
	inArgs[1].SetString(file)
	inArgs[2].SetInt32(line)
	inArgs[3].SetString(func_)
	inArgs[4].SetString(expr)

	assertionMessageExprFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_atexit' : parameter 'func' of type 'VoidFunc' not supported

var atomicIntAddFunction *gi.Function
var atomicIntAddFunction_Once sync.Once

func atomicIntAddFunction_Set() {
	atomicIntAddFunction_Once.Do(func() {
		atomicIntAddFunction = gi.FunctionInvokerNew("GLib", "atomic_int_add")
	})
}

// AtomicIntAdd is a representation of the C type g_atomic_int_add.
func AtomicIntAdd(atomic int32, val int32) int32 {
	atomicIntAddFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetInt32(atomic)
	inArgs[1].SetInt32(val)

	ret := atomicIntAddFunction.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var atomicIntAndFunction *gi.Function
var atomicIntAndFunction_Once sync.Once

func atomicIntAndFunction_Set() {
	atomicIntAndFunction_Once.Do(func() {
		atomicIntAndFunction = gi.FunctionInvokerNew("GLib", "atomic_int_and")
	})
}

// AtomicIntAnd is a representation of the C type g_atomic_int_and.
func AtomicIntAnd(atomic uint32, val uint32) uint32 {
	atomicIntAndFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetUint32(atomic)
	inArgs[1].SetUint32(val)

	ret := atomicIntAndFunction.Invoke(inArgs[:], nil)

	retGo := ret.Uint32()

	return retGo
}

// UNSUPPORTED : C value 'g_atomic_int_compare_and_exchange' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_atomic_int_dec_and_test' : return type 'gboolean' not supported

var atomicIntExchangeAndAddFunction *gi.Function
var atomicIntExchangeAndAddFunction_Once sync.Once

func atomicIntExchangeAndAddFunction_Set() {
	atomicIntExchangeAndAddFunction_Once.Do(func() {
		atomicIntExchangeAndAddFunction = gi.FunctionInvokerNew("GLib", "atomic_int_exchange_and_add")
	})
}

// AtomicIntExchangeAndAdd is a representation of the C type g_atomic_int_exchange_and_add.
func AtomicIntExchangeAndAdd(atomic int32, val int32) int32 {
	atomicIntExchangeAndAddFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetInt32(atomic)
	inArgs[1].SetInt32(val)

	ret := atomicIntExchangeAndAddFunction.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var atomicIntGetFunction *gi.Function
var atomicIntGetFunction_Once sync.Once

func atomicIntGetFunction_Set() {
	atomicIntGetFunction_Once.Do(func() {
		atomicIntGetFunction = gi.FunctionInvokerNew("GLib", "atomic_int_get")
	})
}

// AtomicIntGet is a representation of the C type g_atomic_int_get.
func AtomicIntGet(atomic int32) int32 {
	atomicIntGetFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetInt32(atomic)

	ret := atomicIntGetFunction.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var atomicIntIncFunction *gi.Function
var atomicIntIncFunction_Once sync.Once

func atomicIntIncFunction_Set() {
	atomicIntIncFunction_Once.Do(func() {
		atomicIntIncFunction = gi.FunctionInvokerNew("GLib", "atomic_int_inc")
	})
}

// AtomicIntInc is a representation of the C type g_atomic_int_inc.
func AtomicIntInc(atomic int32) {
	atomicIntIncFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetInt32(atomic)

	atomicIntIncFunction.Invoke(inArgs[:], nil)

}

var atomicIntOrFunction *gi.Function
var atomicIntOrFunction_Once sync.Once

func atomicIntOrFunction_Set() {
	atomicIntOrFunction_Once.Do(func() {
		atomicIntOrFunction = gi.FunctionInvokerNew("GLib", "atomic_int_or")
	})
}

// AtomicIntOr is a representation of the C type g_atomic_int_or.
func AtomicIntOr(atomic uint32, val uint32) uint32 {
	atomicIntOrFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetUint32(atomic)
	inArgs[1].SetUint32(val)

	ret := atomicIntOrFunction.Invoke(inArgs[:], nil)

	retGo := ret.Uint32()

	return retGo
}

var atomicIntSetFunction *gi.Function
var atomicIntSetFunction_Once sync.Once

func atomicIntSetFunction_Set() {
	atomicIntSetFunction_Once.Do(func() {
		atomicIntSetFunction = gi.FunctionInvokerNew("GLib", "atomic_int_set")
	})
}

// AtomicIntSet is a representation of the C type g_atomic_int_set.
func AtomicIntSet(atomic int32, newval int32) {
	atomicIntSetFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetInt32(atomic)
	inArgs[1].SetInt32(newval)

	atomicIntSetFunction.Invoke(inArgs[:], nil)

}

var atomicIntXorFunction *gi.Function
var atomicIntXorFunction_Once sync.Once

func atomicIntXorFunction_Set() {
	atomicIntXorFunction_Once.Do(func() {
		atomicIntXorFunction = gi.FunctionInvokerNew("GLib", "atomic_int_xor")
	})
}

// AtomicIntXor is a representation of the C type g_atomic_int_xor.
func AtomicIntXor(atomic uint32, val uint32) uint32 {
	atomicIntXorFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetUint32(atomic)
	inArgs[1].SetUint32(val)

	ret := atomicIntXorFunction.Invoke(inArgs[:], nil)

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

var atomicRefCountIncFunction *gi.Function
var atomicRefCountIncFunction_Once sync.Once

func atomicRefCountIncFunction_Set() {
	atomicRefCountIncFunction_Once.Do(func() {
		atomicRefCountIncFunction = gi.FunctionInvokerNew("GLib", "atomic_ref_count_inc")
	})
}

// AtomicRefCountInc is a representation of the C type g_atomic_ref_count_inc.
func AtomicRefCountInc(arc int32) {
	atomicRefCountIncFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetInt32(arc)

	atomicRefCountIncFunction.Invoke(inArgs[:], nil)

}

var atomicRefCountInitFunction *gi.Function
var atomicRefCountInitFunction_Once sync.Once

func atomicRefCountInitFunction_Set() {
	atomicRefCountInitFunction_Once.Do(func() {
		atomicRefCountInitFunction = gi.FunctionInvokerNew("GLib", "atomic_ref_count_init")
	})
}

// AtomicRefCountInit is a representation of the C type g_atomic_ref_count_init.
func AtomicRefCountInit(arc int32) {
	atomicRefCountInitFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetInt32(arc)

	atomicRefCountInitFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_base64_decode' : parameter 'out_len' of type 'gsize' not supported

// UNSUPPORTED : C value 'g_base64_decode_inplace' : parameter 'text' has no type

// UNSUPPORTED : C value 'g_base64_decode_step' : parameter 'in' has no type

// UNSUPPORTED : C value 'g_base64_encode' : parameter 'data' has no type

// UNSUPPORTED : C value 'g_base64_encode_close' : parameter 'break_lines' of type 'gboolean' not supported

// UNSUPPORTED : C value 'g_base64_encode_step' : parameter 'in' has no type

// UNSUPPORTED : C value 'g_basename' : parameter 'file_name' of type 'filename' not supported

var bitLockFunction *gi.Function
var bitLockFunction_Once sync.Once

func bitLockFunction_Set() {
	bitLockFunction_Once.Do(func() {
		bitLockFunction = gi.FunctionInvokerNew("GLib", "bit_lock")
	})
}

// BitLock is a representation of the C type g_bit_lock.
func BitLock(address int32, lockBit int32) {
	bitLockFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetInt32(address)
	inArgs[1].SetInt32(lockBit)

	bitLockFunction.Invoke(inArgs[:], nil)

}

var bitNthLsfFunction *gi.Function
var bitNthLsfFunction_Once sync.Once

func bitNthLsfFunction_Set() {
	bitNthLsfFunction_Once.Do(func() {
		bitNthLsfFunction = gi.FunctionInvokerNew("GLib", "bit_nth_lsf")
	})
}

// BitNthLsf is a representation of the C type g_bit_nth_lsf.
func BitNthLsf(mask uint64, nthBit int32) int32 {
	bitNthLsfFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetUint64(mask)
	inArgs[1].SetInt32(nthBit)

	ret := bitNthLsfFunction.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var bitNthMsfFunction *gi.Function
var bitNthMsfFunction_Once sync.Once

func bitNthMsfFunction_Set() {
	bitNthMsfFunction_Once.Do(func() {
		bitNthMsfFunction = gi.FunctionInvokerNew("GLib", "bit_nth_msf")
	})
}

// BitNthMsf is a representation of the C type g_bit_nth_msf.
func BitNthMsf(mask uint64, nthBit int32) int32 {
	bitNthMsfFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetUint64(mask)
	inArgs[1].SetInt32(nthBit)

	ret := bitNthMsfFunction.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var bitStorageFunction *gi.Function
var bitStorageFunction_Once sync.Once

func bitStorageFunction_Set() {
	bitStorageFunction_Once.Do(func() {
		bitStorageFunction = gi.FunctionInvokerNew("GLib", "bit_storage")
	})
}

// BitStorage is a representation of the C type g_bit_storage.
func BitStorage(number uint64) uint32 {
	bitStorageFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetUint64(number)

	ret := bitStorageFunction.Invoke(inArgs[:], nil)

	retGo := ret.Uint32()

	return retGo
}

// UNSUPPORTED : C value 'g_bit_trylock' : return type 'gboolean' not supported

var bitUnlockFunction *gi.Function
var bitUnlockFunction_Once sync.Once

func bitUnlockFunction_Set() {
	bitUnlockFunction_Once.Do(func() {
		bitUnlockFunction = gi.FunctionInvokerNew("GLib", "bit_unlock")
	})
}

// BitUnlock is a representation of the C type g_bit_unlock.
func BitUnlock(address int32, lockBit int32) {
	bitUnlockFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetInt32(address)
	inArgs[1].SetInt32(lockBit)

	bitUnlockFunction.Invoke(inArgs[:], nil)

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

func byteArrayNewFunction_Set() {
	byteArrayNewFunction_Once.Do(func() {
		byteArrayNewFunction = gi.FunctionInvokerNew("GLib", "byte_array_new")
	})
}

// ByteArrayNew is a representation of the C type g_byte_array_new.
func ByteArrayNew() {
	byteArrayNewFunction_Set()

	byteArrayNewFunction.Invoke(nil, nil)

}

// UNSUPPORTED : C value 'g_byte_array_new_take' : parameter 'data' has no type

// UNSUPPORTED : C value 'g_byte_array_unref' : parameter 'array' has no type

// UNSUPPORTED : C value 'g_canonicalize_filename' : parameter 'filename' of type 'filename' not supported

// UNSUPPORTED : C value 'g_chdir' : parameter 'path' of type 'filename' not supported

var checkVersionFunction *gi.Function
var checkVersionFunction_Once sync.Once

func checkVersionFunction_Set() {
	checkVersionFunction_Once.Do(func() {
		checkVersionFunction = gi.FunctionInvokerNew("GLib", "check_version")
	})
}

// CheckVersion is a representation of the C type glib_check_version.
func CheckVersion(requiredMajor uint32, requiredMinor uint32, requiredMicro uint32) string {
	checkVersionFunction_Set()

	var inArgs [3]gi.Argument
	inArgs[0].SetUint32(requiredMajor)
	inArgs[1].SetUint32(requiredMinor)
	inArgs[2].SetUint32(requiredMicro)

	ret := checkVersionFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'g_checksum_type_get_length' : parameter 'checksum_type' of type 'ChecksumType' not supported

// UNSUPPORTED : C value 'g_child_watch_add' : parameter 'pid' of type 'Pid' not supported

// UNSUPPORTED : C value 'g_child_watch_add_full' : parameter 'pid' of type 'Pid' not supported

// UNSUPPORTED : C value 'g_child_watch_source_new' : parameter 'pid' of type 'Pid' not supported

var clearErrorFunction *gi.Function
var clearErrorFunction_Once sync.Once

func clearErrorFunction_Set() {
	clearErrorFunction_Once.Do(func() {
		clearErrorFunction = gi.FunctionInvokerNew("GLib", "clear_error")
	})
}

// ClearError is a representation of the C type g_clear_error.
func ClearError() {
	clearErrorFunction_Set()

	clearErrorFunction.Invoke(nil, nil)

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

var dcgettextFunction *gi.Function
var dcgettextFunction_Once sync.Once

func dcgettextFunction_Set() {
	dcgettextFunction_Once.Do(func() {
		dcgettextFunction = gi.FunctionInvokerNew("GLib", "dcgettext")
	})
}

// Dcgettext is a representation of the C type g_dcgettext.
func Dcgettext(domain string, msgid string, category int32) string {
	dcgettextFunction_Set()

	var inArgs [3]gi.Argument
	inArgs[0].SetString(domain)
	inArgs[1].SetString(msgid)
	inArgs[2].SetInt32(category)

	ret := dcgettextFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

var dgettextFunction *gi.Function
var dgettextFunction_Once sync.Once

func dgettextFunction_Set() {
	dgettextFunction_Once.Do(func() {
		dgettextFunction = gi.FunctionInvokerNew("GLib", "dgettext")
	})
}

// Dgettext is a representation of the C type g_dgettext.
func Dgettext(domain string, msgid string) string {
	dgettextFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetString(domain)
	inArgs[1].SetString(msgid)

	ret := dgettextFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'g_dir_make_tmp' : parameter 'tmpl' of type 'filename' not supported

// UNSUPPORTED : C value 'g_direct_equal' : parameter 'v1' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_direct_hash' : parameter 'v' of type 'gpointer' not supported

var dngettextFunction *gi.Function
var dngettextFunction_Once sync.Once

func dngettextFunction_Set() {
	dngettextFunction_Once.Do(func() {
		dngettextFunction = gi.FunctionInvokerNew("GLib", "dngettext")
	})
}

// Dngettext is a representation of the C type g_dngettext.
func Dngettext(domain string, msgid string, msgidPlural string, n uint64) string {
	dngettextFunction_Set()

	var inArgs [4]gi.Argument
	inArgs[0].SetString(domain)
	inArgs[1].SetString(msgid)
	inArgs[2].SetString(msgidPlural)
	inArgs[3].SetUint64(n)

	ret := dngettextFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'g_double_equal' : parameter 'v1' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_double_hash' : parameter 'v' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_dpgettext' : parameter 'msgidoffset' of type 'gsize' not supported

var dpgettext2Function *gi.Function
var dpgettext2Function_Once sync.Once

func dpgettext2Function_Set() {
	dpgettext2Function_Once.Do(func() {
		dpgettext2Function = gi.FunctionInvokerNew("GLib", "dpgettext2")
	})
}

// Dpgettext2 is a representation of the C type g_dpgettext2.
func Dpgettext2(domain string, context string, msgid string) string {
	dpgettext2Function_Set()

	var inArgs [3]gi.Argument
	inArgs[0].SetString(domain)
	inArgs[1].SetString(context)
	inArgs[2].SetString(msgid)

	ret := dpgettext2Function.Invoke(inArgs[:], nil)

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

var formatSizeFunction *gi.Function
var formatSizeFunction_Once sync.Once

func formatSizeFunction_Set() {
	formatSizeFunction_Once.Do(func() {
		formatSizeFunction = gi.FunctionInvokerNew("GLib", "format_size")
	})
}

// FormatSize is a representation of the C type g_format_size.
func FormatSize(size uint64) string {
	formatSizeFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetUint64(size)

	ret := formatSizeFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

var formatSizeForDisplayFunction *gi.Function
var formatSizeForDisplayFunction_Once sync.Once

func formatSizeForDisplayFunction_Set() {
	formatSizeForDisplayFunction_Once.Do(func() {
		formatSizeForDisplayFunction = gi.FunctionInvokerNew("GLib", "format_size_for_display")
	})
}

// FormatSizeForDisplay is a representation of the C type g_format_size_for_display.
func FormatSizeForDisplay(size int64) string {
	formatSizeForDisplayFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetInt64(size)

	ret := formatSizeForDisplayFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

// UNSUPPORTED : C value 'g_format_size_full' : parameter 'flags' of type 'FormatSizeFlags' not supported

// UNSUPPORTED : C value 'g_fprintf' : parameter 'file' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_free' : parameter 'mem' of type 'gpointer' not supported

var getApplicationNameFunction *gi.Function
var getApplicationNameFunction_Once sync.Once

func getApplicationNameFunction_Set() {
	getApplicationNameFunction_Once.Do(func() {
		getApplicationNameFunction = gi.FunctionInvokerNew("GLib", "get_application_name")
	})
}

// GetApplicationName is a representation of the C type g_get_application_name.
func GetApplicationName() string {
	getApplicationNameFunction_Set()

	ret := getApplicationNameFunction.Invoke(nil, nil)

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'g_get_charset' : return type 'gboolean' not supported

var getCodesetFunction *gi.Function
var getCodesetFunction_Once sync.Once

func getCodesetFunction_Set() {
	getCodesetFunction_Once.Do(func() {
		getCodesetFunction = gi.FunctionInvokerNew("GLib", "get_codeset")
	})
}

// GetCodeset is a representation of the C type g_get_codeset.
func GetCodeset() string {
	getCodesetFunction_Set()

	ret := getCodesetFunction.Invoke(nil, nil)

	retGo := ret.String(true)

	return retGo
}

// UNSUPPORTED : C value 'g_get_console_charset' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_get_current_dir' : return type 'filename' not supported

// UNSUPPORTED : C value 'g_get_current_time' : parameter 'result' of type 'TimeVal' not supported

var getEnvironFunction *gi.Function
var getEnvironFunction_Once sync.Once

func getEnvironFunction_Set() {
	getEnvironFunction_Once.Do(func() {
		getEnvironFunction = gi.FunctionInvokerNew("GLib", "get_environ")
	})
}

// GetEnviron is a representation of the C type g_get_environ.
func GetEnviron() {
	getEnvironFunction_Set()

	getEnvironFunction.Invoke(nil, nil)

}

// UNSUPPORTED : C value 'g_get_filename_charsets' : parameter 'filename_charsets' has no type

// UNSUPPORTED : C value 'g_get_home_dir' : return type 'filename' not supported

var getHostNameFunction *gi.Function
var getHostNameFunction_Once sync.Once

func getHostNameFunction_Set() {
	getHostNameFunction_Once.Do(func() {
		getHostNameFunction = gi.FunctionInvokerNew("GLib", "get_host_name")
	})
}

// GetHostName is a representation of the C type g_get_host_name.
func GetHostName() string {
	getHostNameFunction_Set()

	ret := getHostNameFunction.Invoke(nil, nil)

	retGo := ret.String(false)

	return retGo
}

var getLanguageNamesFunction *gi.Function
var getLanguageNamesFunction_Once sync.Once

func getLanguageNamesFunction_Set() {
	getLanguageNamesFunction_Once.Do(func() {
		getLanguageNamesFunction = gi.FunctionInvokerNew("GLib", "get_language_names")
	})
}

// GetLanguageNames is a representation of the C type g_get_language_names.
func GetLanguageNames() {
	getLanguageNamesFunction_Set()

	getLanguageNamesFunction.Invoke(nil, nil)

}

var getLanguageNamesWithCategoryFunction *gi.Function
var getLanguageNamesWithCategoryFunction_Once sync.Once

func getLanguageNamesWithCategoryFunction_Set() {
	getLanguageNamesWithCategoryFunction_Once.Do(func() {
		getLanguageNamesWithCategoryFunction = gi.FunctionInvokerNew("GLib", "get_language_names_with_category")
	})
}

// GetLanguageNamesWithCategory is a representation of the C type g_get_language_names_with_category.
func GetLanguageNamesWithCategory(categoryName string) {
	getLanguageNamesWithCategoryFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetString(categoryName)

	getLanguageNamesWithCategoryFunction.Invoke(inArgs[:], nil)

}

var getLocaleVariantsFunction *gi.Function
var getLocaleVariantsFunction_Once sync.Once

func getLocaleVariantsFunction_Set() {
	getLocaleVariantsFunction_Once.Do(func() {
		getLocaleVariantsFunction = gi.FunctionInvokerNew("GLib", "get_locale_variants")
	})
}

// GetLocaleVariants is a representation of the C type g_get_locale_variants.
func GetLocaleVariants(locale string) {
	getLocaleVariantsFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetString(locale)

	getLocaleVariantsFunction.Invoke(inArgs[:], nil)

}

var getMonotonicTimeFunction *gi.Function
var getMonotonicTimeFunction_Once sync.Once

func getMonotonicTimeFunction_Set() {
	getMonotonicTimeFunction_Once.Do(func() {
		getMonotonicTimeFunction = gi.FunctionInvokerNew("GLib", "get_monotonic_time")
	})
}

// GetMonotonicTime is a representation of the C type g_get_monotonic_time.
func GetMonotonicTime() int64 {
	getMonotonicTimeFunction_Set()

	ret := getMonotonicTimeFunction.Invoke(nil, nil)

	retGo := ret.Int64()

	return retGo
}

var getNumProcessorsFunction *gi.Function
var getNumProcessorsFunction_Once sync.Once

func getNumProcessorsFunction_Set() {
	getNumProcessorsFunction_Once.Do(func() {
		getNumProcessorsFunction = gi.FunctionInvokerNew("GLib", "get_num_processors")
	})
}

// GetNumProcessors is a representation of the C type g_get_num_processors.
func GetNumProcessors() uint32 {
	getNumProcessorsFunction_Set()

	ret := getNumProcessorsFunction.Invoke(nil, nil)

	retGo := ret.Uint32()

	return retGo
}

var getPrgnameFunction *gi.Function
var getPrgnameFunction_Once sync.Once

func getPrgnameFunction_Set() {
	getPrgnameFunction_Once.Do(func() {
		getPrgnameFunction = gi.FunctionInvokerNew("GLib", "get_prgname")
	})
}

// GetPrgname is a representation of the C type g_get_prgname.
func GetPrgname() string {
	getPrgnameFunction_Set()

	ret := getPrgnameFunction.Invoke(nil, nil)

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'g_get_real_name' : return type 'filename' not supported

var getRealTimeFunction *gi.Function
var getRealTimeFunction_Once sync.Once

func getRealTimeFunction_Set() {
	getRealTimeFunction_Once.Do(func() {
		getRealTimeFunction = gi.FunctionInvokerNew("GLib", "get_real_time")
	})
}

// GetRealTime is a representation of the C type g_get_real_time.
func GetRealTime() int64 {
	getRealTimeFunction_Set()

	ret := getRealTimeFunction.Invoke(nil, nil)

	retGo := ret.Int64()

	return retGo
}

var getSystemConfigDirsFunction *gi.Function
var getSystemConfigDirsFunction_Once sync.Once

func getSystemConfigDirsFunction_Set() {
	getSystemConfigDirsFunction_Once.Do(func() {
		getSystemConfigDirsFunction = gi.FunctionInvokerNew("GLib", "get_system_config_dirs")
	})
}

// GetSystemConfigDirs is a representation of the C type g_get_system_config_dirs.
func GetSystemConfigDirs() {
	getSystemConfigDirsFunction_Set()

	getSystemConfigDirsFunction.Invoke(nil, nil)

}

var getSystemDataDirsFunction *gi.Function
var getSystemDataDirsFunction_Once sync.Once

func getSystemDataDirsFunction_Set() {
	getSystemDataDirsFunction_Once.Do(func() {
		getSystemDataDirsFunction = gi.FunctionInvokerNew("GLib", "get_system_data_dirs")
	})
}

// GetSystemDataDirs is a representation of the C type g_get_system_data_dirs.
func GetSystemDataDirs() {
	getSystemDataDirsFunction_Set()

	getSystemDataDirsFunction.Invoke(nil, nil)

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

var hostnameToAsciiFunction *gi.Function
var hostnameToAsciiFunction_Once sync.Once

func hostnameToAsciiFunction_Set() {
	hostnameToAsciiFunction_Once.Do(func() {
		hostnameToAsciiFunction = gi.FunctionInvokerNew("GLib", "hostname_to_ascii")
	})
}

// HostnameToAscii is a representation of the C type g_hostname_to_ascii.
func HostnameToAscii(hostname string) string {
	hostnameToAsciiFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetString(hostname)

	ret := hostnameToAsciiFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

var hostnameToUnicodeFunction *gi.Function
var hostnameToUnicodeFunction_Once sync.Once

func hostnameToUnicodeFunction_Set() {
	hostnameToUnicodeFunction_Once.Do(func() {
		hostnameToUnicodeFunction = gi.FunctionInvokerNew("GLib", "hostname_to_unicode")
	})
}

// HostnameToUnicode is a representation of the C type g_hostname_to_unicode.
func HostnameToUnicode(hostname string) string {
	hostnameToUnicodeFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetString(hostname)

	ret := hostnameToUnicodeFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

// UNSUPPORTED : C value 'g_iconv' : parameter 'converter' of type 'IConv' not supported

var iconvOpenFunction *gi.Function
var iconvOpenFunction_Once sync.Once

func iconvOpenFunction_Set() {
	iconvOpenFunction_Once.Do(func() {
		iconvOpenFunction = gi.FunctionInvokerNew("GLib", "iconv_open")
	})
}

// IconvOpen is a representation of the C type g_iconv_open.
func IconvOpen(toCodeset string, fromCodeset string) *IConv {
	iconvOpenFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetString(toCodeset)
	inArgs[1].SetString(fromCodeset)

	ret := iconvOpenFunction.Invoke(inArgs[:], nil)

	retGo := &IConv{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_idle_add' : parameter 'function' of type 'SourceFunc' not supported

// UNSUPPORTED : C value 'g_idle_add_full' : parameter 'function' of type 'SourceFunc' not supported

// UNSUPPORTED : C value 'g_idle_remove_by_data' : parameter 'data' of type 'gpointer' not supported

var idleSourceNewFunction *gi.Function
var idleSourceNewFunction_Once sync.Once

func idleSourceNewFunction_Set() {
	idleSourceNewFunction_Once.Do(func() {
		idleSourceNewFunction = gi.FunctionInvokerNew("GLib", "idle_source_new")
	})
}

// IdleSourceNew is a representation of the C type g_idle_source_new.
func IdleSourceNew() *Source {
	idleSourceNewFunction_Set()

	ret := idleSourceNewFunction.Invoke(nil, nil)

	retGo := &Source{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_int64_equal' : parameter 'v1' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_int64_hash' : parameter 'v' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_int_equal' : parameter 'v1' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_int_hash' : parameter 'v' of type 'gpointer' not supported

var internStaticStringFunction *gi.Function
var internStaticStringFunction_Once sync.Once

func internStaticStringFunction_Set() {
	internStaticStringFunction_Once.Do(func() {
		internStaticStringFunction = gi.FunctionInvokerNew("GLib", "intern_static_string")
	})
}

// InternStaticString is a representation of the C type g_intern_static_string.
func InternStaticString(string_ string) string {
	internStaticStringFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetString(string_)

	ret := internStaticStringFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

var internStringFunction *gi.Function
var internStringFunction_Once sync.Once

func internStringFunction_Set() {
	internStringFunction_Once.Do(func() {
		internStringFunction = gi.FunctionInvokerNew("GLib", "intern_string")
	})
}

// InternString is a representation of the C type g_intern_string.
func InternString(string_ string) string {
	internStringFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetString(string_)

	ret := internStringFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'g_io_add_watch' : parameter 'channel' of type 'IOChannel' not supported

// UNSUPPORTED : C value 'g_io_add_watch_full' : parameter 'channel' of type 'IOChannel' not supported

// UNSUPPORTED : C value 'g_io_channel_error_from_errno' : return type 'IOChannelError' not supported

// UNSUPPORTED : C value 'g_io_channel_error_quark' : return type 'Quark' not supported

// UNSUPPORTED : C value 'g_io_create_watch' : parameter 'channel' of type 'IOChannel' not supported

// UNSUPPORTED : C value 'g_key_file_error_quark' : return type 'Quark' not supported

var listenvFunction *gi.Function
var listenvFunction_Once sync.Once

func listenvFunction_Set() {
	listenvFunction_Once.Do(func() {
		listenvFunction = gi.FunctionInvokerNew("GLib", "listenv")
	})
}

// Listenv is a representation of the C type g_listenv.
func Listenv() {
	listenvFunction_Set()

	listenvFunction.Invoke(nil, nil)

}

// UNSUPPORTED : C value 'g_locale_from_utf8' : parameter 'bytes_read' of type 'gsize' not supported

// UNSUPPORTED : C value 'g_locale_to_utf8' : parameter 'opsysstring' has no type

// UNSUPPORTED : C value 'g_log' : parameter 'log_level' of type 'LogLevelFlags' not supported

// UNSUPPORTED : C value 'g_log_default_handler' : parameter 'log_level' of type 'LogLevelFlags' not supported

var logRemoveHandlerFunction *gi.Function
var logRemoveHandlerFunction_Once sync.Once

func logRemoveHandlerFunction_Set() {
	logRemoveHandlerFunction_Once.Do(func() {
		logRemoveHandlerFunction = gi.FunctionInvokerNew("GLib", "log_remove_handler")
	})
}

// LogRemoveHandler is a representation of the C type g_log_remove_handler.
func LogRemoveHandler(logDomain string, handlerId uint32) {
	logRemoveHandlerFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetString(logDomain)
	inArgs[1].SetUint32(handlerId)

	logRemoveHandlerFunction.Invoke(inArgs[:], nil)

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

var mainContextDefaultFunction *gi.Function
var mainContextDefaultFunction_Once sync.Once

func mainContextDefaultFunction_Set() {
	mainContextDefaultFunction_Once.Do(func() {
		mainContextDefaultFunction = gi.FunctionInvokerNew("GLib", "main_context_default")
	})
}

// MainContextDefault is a representation of the C type g_main_context_default.
func MainContextDefault() *MainContext {
	mainContextDefaultFunction_Set()

	ret := mainContextDefaultFunction.Invoke(nil, nil)

	retGo := &MainContext{native: ret.Pointer()}

	return retGo
}

var mainContextGetThreadDefaultFunction *gi.Function
var mainContextGetThreadDefaultFunction_Once sync.Once

func mainContextGetThreadDefaultFunction_Set() {
	mainContextGetThreadDefaultFunction_Once.Do(func() {
		mainContextGetThreadDefaultFunction = gi.FunctionInvokerNew("GLib", "main_context_get_thread_default")
	})
}

// MainContextGetThreadDefault is a representation of the C type g_main_context_get_thread_default.
func MainContextGetThreadDefault() *MainContext {
	mainContextGetThreadDefaultFunction_Set()

	ret := mainContextGetThreadDefaultFunction.Invoke(nil, nil)

	retGo := &MainContext{native: ret.Pointer()}

	return retGo
}

var mainContextRefThreadDefaultFunction *gi.Function
var mainContextRefThreadDefaultFunction_Once sync.Once

func mainContextRefThreadDefaultFunction_Set() {
	mainContextRefThreadDefaultFunction_Once.Do(func() {
		mainContextRefThreadDefaultFunction = gi.FunctionInvokerNew("GLib", "main_context_ref_thread_default")
	})
}

// MainContextRefThreadDefault is a representation of the C type g_main_context_ref_thread_default.
func MainContextRefThreadDefault() *MainContext {
	mainContextRefThreadDefaultFunction_Set()

	ret := mainContextRefThreadDefaultFunction.Invoke(nil, nil)

	retGo := &MainContext{native: ret.Pointer()}

	return retGo
}

var mainCurrentSourceFunction *gi.Function
var mainCurrentSourceFunction_Once sync.Once

func mainCurrentSourceFunction_Set() {
	mainCurrentSourceFunction_Once.Do(func() {
		mainCurrentSourceFunction = gi.FunctionInvokerNew("GLib", "main_current_source")
	})
}

// MainCurrentSource is a representation of the C type g_main_current_source.
func MainCurrentSource() *Source {
	mainCurrentSourceFunction_Set()

	ret := mainCurrentSourceFunction.Invoke(nil, nil)

	retGo := &Source{native: ret.Pointer()}

	return retGo
}

var mainDepthFunction *gi.Function
var mainDepthFunction_Once sync.Once

func mainDepthFunction_Set() {
	mainDepthFunction_Once.Do(func() {
		mainDepthFunction = gi.FunctionInvokerNew("GLib", "main_depth")
	})
}

// MainDepth is a representation of the C type g_main_depth.
func MainDepth() int32 {
	mainDepthFunction_Set()

	ret := mainDepthFunction.Invoke(nil, nil)

	retGo := ret.Int32()

	return retGo
}

// UNSUPPORTED : C value 'g_malloc' : parameter 'n_bytes' of type 'gsize' not supported

// UNSUPPORTED : C value 'g_malloc0' : parameter 'n_bytes' of type 'gsize' not supported

// UNSUPPORTED : C value 'g_malloc0_n' : parameter 'n_blocks' of type 'gsize' not supported

// UNSUPPORTED : C value 'g_malloc_n' : parameter 'n_blocks' of type 'gsize' not supported

// UNSUPPORTED : C value 'g_markup_collect_attributes' : parameter 'error' of type 'Error' not supported

// UNSUPPORTED : C value 'g_markup_error_quark' : return type 'Quark' not supported

var markupEscapeTextFunction *gi.Function
var markupEscapeTextFunction_Once sync.Once

func markupEscapeTextFunction_Set() {
	markupEscapeTextFunction_Once.Do(func() {
		markupEscapeTextFunction = gi.FunctionInvokerNew("GLib", "markup_escape_text")
	})
}

// MarkupEscapeText is a representation of the C type g_markup_escape_text.
func MarkupEscapeText(text string, length int32) string {
	markupEscapeTextFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetString(text)
	inArgs[1].SetInt32(length)

	ret := markupEscapeTextFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

// UNSUPPORTED : C value 'g_markup_printf_escaped' : parameter '...' has no type

// UNSUPPORTED : C value 'g_markup_vprintf_escaped' : parameter 'args' of type 'va_list' not supported

// UNSUPPORTED : C value 'g_mem_is_system_malloc' : return type 'gboolean' not supported

var memProfileFunction *gi.Function
var memProfileFunction_Once sync.Once

func memProfileFunction_Set() {
	memProfileFunction_Once.Do(func() {
		memProfileFunction = gi.FunctionInvokerNew("GLib", "mem_profile")
	})
}

// MemProfile is a representation of the C type g_mem_profile.
func MemProfile() {
	memProfileFunction_Set()

	memProfileFunction.Invoke(nil, nil)

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

func onErrorQueryFunction_Set() {
	onErrorQueryFunction_Once.Do(func() {
		onErrorQueryFunction = gi.FunctionInvokerNew("GLib", "on_error_query")
	})
}

// OnErrorQuery is a representation of the C type g_on_error_query.
func OnErrorQuery(prgName string) {
	onErrorQueryFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetString(prgName)

	onErrorQueryFunction.Invoke(inArgs[:], nil)

}

var onErrorStackTraceFunction *gi.Function
var onErrorStackTraceFunction_Once sync.Once

func onErrorStackTraceFunction_Set() {
	onErrorStackTraceFunction_Once.Do(func() {
		onErrorStackTraceFunction = gi.FunctionInvokerNew("GLib", "on_error_stack_trace")
	})
}

// OnErrorStackTrace is a representation of the C type g_on_error_stack_trace.
func OnErrorStackTrace(prgName string) {
	onErrorStackTraceFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetString(prgName)

	onErrorStackTraceFunction.Invoke(inArgs[:], nil)

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

var randomIntFunction *gi.Function
var randomIntFunction_Once sync.Once

func randomIntFunction_Set() {
	randomIntFunction_Once.Do(func() {
		randomIntFunction = gi.FunctionInvokerNew("GLib", "random_int")
	})
}

// RandomInt is a representation of the C type g_random_int.
func RandomInt() uint32 {
	randomIntFunction_Set()

	ret := randomIntFunction.Invoke(nil, nil)

	retGo := ret.Uint32()

	return retGo
}

var randomIntRangeFunction *gi.Function
var randomIntRangeFunction_Once sync.Once

func randomIntRangeFunction_Set() {
	randomIntRangeFunction_Once.Do(func() {
		randomIntRangeFunction = gi.FunctionInvokerNew("GLib", "random_int_range")
	})
}

// RandomIntRange is a representation of the C type g_random_int_range.
func RandomIntRange(begin int32, end int32) int32 {
	randomIntRangeFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetInt32(begin)
	inArgs[1].SetInt32(end)

	ret := randomIntRangeFunction.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var randomSetSeedFunction *gi.Function
var randomSetSeedFunction_Once sync.Once

func randomSetSeedFunction_Set() {
	randomSetSeedFunction_Once.Do(func() {
		randomSetSeedFunction = gi.FunctionInvokerNew("GLib", "random_set_seed")
	})
}

// RandomSetSeed is a representation of the C type g_random_set_seed.
func RandomSetSeed(seed uint32) {
	randomSetSeedFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetUint32(seed)

	randomSetSeedFunction.Invoke(inArgs[:], nil)

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

var refCountIncFunction *gi.Function
var refCountIncFunction_Once sync.Once

func refCountIncFunction_Set() {
	refCountIncFunction_Once.Do(func() {
		refCountIncFunction = gi.FunctionInvokerNew("GLib", "ref_count_inc")
	})
}

// RefCountInc is a representation of the C type g_ref_count_inc.
func RefCountInc(rc int32) {
	refCountIncFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetInt32(rc)

	refCountIncFunction.Invoke(inArgs[:], nil)

}

var refCountInitFunction *gi.Function
var refCountInitFunction_Once sync.Once

func refCountInitFunction_Set() {
	refCountInitFunction_Once.Do(func() {
		refCountInitFunction = gi.FunctionInvokerNew("GLib", "ref_count_init")
	})
}

// RefCountInit is a representation of the C type g_ref_count_init.
func RefCountInit(rc int32) {
	refCountInitFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetInt32(rc)

	refCountInitFunction.Invoke(inArgs[:], nil)

}

var refStringAcquireFunction *gi.Function
var refStringAcquireFunction_Once sync.Once

func refStringAcquireFunction_Set() {
	refStringAcquireFunction_Once.Do(func() {
		refStringAcquireFunction = gi.FunctionInvokerNew("GLib", "ref_string_acquire")
	})
}

// RefStringAcquire is a representation of the C type g_ref_string_acquire.
func RefStringAcquire(str string) string {
	refStringAcquireFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetString(str)

	ret := refStringAcquireFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

// UNSUPPORTED : C value 'g_ref_string_length' : return type 'gsize' not supported

var refStringNewFunction *gi.Function
var refStringNewFunction_Once sync.Once

func refStringNewFunction_Set() {
	refStringNewFunction_Once.Do(func() {
		refStringNewFunction = gi.FunctionInvokerNew("GLib", "ref_string_new")
	})
}

// RefStringNew is a representation of the C type g_ref_string_new.
func RefStringNew(str string) string {
	refStringNewFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetString(str)

	ret := refStringNewFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

var refStringNewInternFunction *gi.Function
var refStringNewInternFunction_Once sync.Once

func refStringNewInternFunction_Set() {
	refStringNewInternFunction_Once.Do(func() {
		refStringNewInternFunction = gi.FunctionInvokerNew("GLib", "ref_string_new_intern")
	})
}

// RefStringNewIntern is a representation of the C type g_ref_string_new_intern.
func RefStringNewIntern(str string) string {
	refStringNewInternFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetString(str)

	ret := refStringNewInternFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

var refStringNewLenFunction *gi.Function
var refStringNewLenFunction_Once sync.Once

func refStringNewLenFunction_Set() {
	refStringNewLenFunction_Once.Do(func() {
		refStringNewLenFunction = gi.FunctionInvokerNew("GLib", "ref_string_new_len")
	})
}

// RefStringNewLen is a representation of the C type g_ref_string_new_len.
func RefStringNewLen(str string, len int32) string {
	refStringNewLenFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetString(str)
	inArgs[1].SetInt32(len)

	ret := refStringNewLenFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

var refStringReleaseFunction *gi.Function
var refStringReleaseFunction_Once sync.Once

func refStringReleaseFunction_Set() {
	refStringReleaseFunction_Once.Do(func() {
		refStringReleaseFunction = gi.FunctionInvokerNew("GLib", "ref_string_release")
	})
}

// RefStringRelease is a representation of the C type g_ref_string_release.
func RefStringRelease(str string) {
	refStringReleaseFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetString(str)

	refStringReleaseFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_regex_check_replacement' : parameter 'has_references' of type 'gboolean' not supported

// UNSUPPORTED : C value 'g_regex_error_quark' : return type 'Quark' not supported

var regexEscapeNulFunction *gi.Function
var regexEscapeNulFunction_Once sync.Once

func regexEscapeNulFunction_Set() {
	regexEscapeNulFunction_Once.Do(func() {
		regexEscapeNulFunction = gi.FunctionInvokerNew("GLib", "regex_escape_nul")
	})
}

// RegexEscapeNul is a representation of the C type g_regex_escape_nul.
func RegexEscapeNul(string_ string, length int32) string {
	regexEscapeNulFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetString(string_)
	inArgs[1].SetInt32(length)

	ret := regexEscapeNulFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

// UNSUPPORTED : C value 'g_regex_escape_string' : parameter 'string' has no type

// UNSUPPORTED : C value 'g_regex_match_simple' : parameter 'compile_options' of type 'RegexCompileFlags' not supported

// UNSUPPORTED : C value 'g_regex_split_simple' : parameter 'compile_options' of type 'RegexCompileFlags' not supported

var reloadUserSpecialDirsCacheFunction *gi.Function
var reloadUserSpecialDirsCacheFunction_Once sync.Once

func reloadUserSpecialDirsCacheFunction_Set() {
	reloadUserSpecialDirsCacheFunction_Once.Do(func() {
		reloadUserSpecialDirsCacheFunction = gi.FunctionInvokerNew("GLib", "reload_user_special_dirs_cache")
	})
}

// ReloadUserSpecialDirsCache is a representation of the C type g_reload_user_special_dirs_cache.
func ReloadUserSpecialDirsCache() {
	reloadUserSpecialDirsCacheFunction_Set()

	reloadUserSpecialDirsCacheFunction.Invoke(nil, nil)

}

var returnIfFailWarningFunction *gi.Function
var returnIfFailWarningFunction_Once sync.Once

func returnIfFailWarningFunction_Set() {
	returnIfFailWarningFunction_Once.Do(func() {
		returnIfFailWarningFunction = gi.FunctionInvokerNew("GLib", "return_if_fail_warning")
	})
}

// ReturnIfFailWarning is a representation of the C type g_return_if_fail_warning.
func ReturnIfFailWarning(logDomain string, prettyFunction string, expression string) {
	returnIfFailWarningFunction_Set()

	var inArgs [3]gi.Argument
	inArgs[0].SetString(logDomain)
	inArgs[1].SetString(prettyFunction)
	inArgs[2].SetString(expression)

	returnIfFailWarningFunction.Invoke(inArgs[:], nil)

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

func setApplicationNameFunction_Set() {
	setApplicationNameFunction_Once.Do(func() {
		setApplicationNameFunction = gi.FunctionInvokerNew("GLib", "set_application_name")
	})
}

// SetApplicationName is a representation of the C type g_set_application_name.
func SetApplicationName(applicationName string) {
	setApplicationNameFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetString(applicationName)

	setApplicationNameFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_set_error' : parameter 'err' of type 'Error' not supported

// UNSUPPORTED : C value 'g_set_error_literal' : parameter 'err' of type 'Error' not supported

var setPrgnameFunction *gi.Function
var setPrgnameFunction_Once sync.Once

func setPrgnameFunction_Set() {
	setPrgnameFunction_Once.Do(func() {
		setPrgnameFunction = gi.FunctionInvokerNew("GLib", "set_prgname")
	})
}

// SetPrgname is a representation of the C type g_set_prgname.
func SetPrgname(prgname string) {
	setPrgnameFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetString(prgname)

	setPrgnameFunction.Invoke(inArgs[:], nil)

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

var sourceSetNameByIdFunction *gi.Function
var sourceSetNameByIdFunction_Once sync.Once

func sourceSetNameByIdFunction_Set() {
	sourceSetNameByIdFunction_Once.Do(func() {
		sourceSetNameByIdFunction = gi.FunctionInvokerNew("GLib", "source_set_name_by_id")
	})
}

// SourceSetNameById is a representation of the C type g_source_set_name_by_id.
func SourceSetNameById(tag uint32, name string) {
	sourceSetNameByIdFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetUint32(tag)
	inArgs[1].SetString(name)

	sourceSetNameByIdFunction.Invoke(inArgs[:], nil)

}

var spacedPrimesClosestFunction *gi.Function
var spacedPrimesClosestFunction_Once sync.Once

func spacedPrimesClosestFunction_Set() {
	spacedPrimesClosestFunction_Once.Do(func() {
		spacedPrimesClosestFunction = gi.FunctionInvokerNew("GLib", "spaced_primes_closest")
	})
}

// SpacedPrimesClosest is a representation of the C type g_spaced_primes_closest.
func SpacedPrimesClosest(num uint32) uint32 {
	spacedPrimesClosestFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetUint32(num)

	ret := spacedPrimesClosestFunction.Invoke(inArgs[:], nil)

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

var stpcpyFunction *gi.Function
var stpcpyFunction_Once sync.Once

func stpcpyFunction_Set() {
	stpcpyFunction_Once.Do(func() {
		stpcpyFunction = gi.FunctionInvokerNew("GLib", "stpcpy")
	})
}

// Stpcpy is a representation of the C type g_stpcpy.
func Stpcpy(dest string, src string) string {
	stpcpyFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetString(dest)
	inArgs[1].SetString(src)

	ret := stpcpyFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

// UNSUPPORTED : C value 'g_str_equal' : parameter 'v1' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_str_has_prefix' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_str_has_suffix' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_str_hash' : parameter 'v' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_str_is_ascii' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_str_match_string' : parameter 'accept_alternates' of type 'gboolean' not supported

var strToAsciiFunction *gi.Function
var strToAsciiFunction_Once sync.Once

func strToAsciiFunction_Set() {
	strToAsciiFunction_Once.Do(func() {
		strToAsciiFunction = gi.FunctionInvokerNew("GLib", "str_to_ascii")
	})
}

// StrToAscii is a representation of the C type g_str_to_ascii.
func StrToAscii(str string, fromLocale string) string {
	strToAsciiFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetString(str)
	inArgs[1].SetString(fromLocale)

	ret := strToAsciiFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

// UNSUPPORTED : C value 'g_str_tokenize_and_fold' : parameter 'ascii_alternates' has no type

var strcanonFunction *gi.Function
var strcanonFunction_Once sync.Once

func strcanonFunction_Set() {
	strcanonFunction_Once.Do(func() {
		strcanonFunction = gi.FunctionInvokerNew("GLib", "strcanon")
	})
}

// Strcanon is a representation of the C type g_strcanon.
func Strcanon(string_ string, validChars string, substitutor int8) string {
	strcanonFunction_Set()

	var inArgs [3]gi.Argument
	inArgs[0].SetString(string_)
	inArgs[1].SetString(validChars)
	inArgs[2].SetInt8(substitutor)

	ret := strcanonFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

var strcasecmpFunction *gi.Function
var strcasecmpFunction_Once sync.Once

func strcasecmpFunction_Set() {
	strcasecmpFunction_Once.Do(func() {
		strcasecmpFunction = gi.FunctionInvokerNew("GLib", "strcasecmp")
	})
}

// Strcasecmp is a representation of the C type g_strcasecmp.
func Strcasecmp(s1 string, s2 string) int32 {
	strcasecmpFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetString(s1)
	inArgs[1].SetString(s2)

	ret := strcasecmpFunction.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var strchompFunction *gi.Function
var strchompFunction_Once sync.Once

func strchompFunction_Set() {
	strchompFunction_Once.Do(func() {
		strchompFunction = gi.FunctionInvokerNew("GLib", "strchomp")
	})
}

// Strchomp is a representation of the C type g_strchomp.
func Strchomp(string_ string) string {
	strchompFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetString(string_)

	ret := strchompFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

var strchugFunction *gi.Function
var strchugFunction_Once sync.Once

func strchugFunction_Set() {
	strchugFunction_Once.Do(func() {
		strchugFunction = gi.FunctionInvokerNew("GLib", "strchug")
	})
}

// Strchug is a representation of the C type g_strchug.
func Strchug(string_ string) string {
	strchugFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetString(string_)

	ret := strchugFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

var strcmp0Function *gi.Function
var strcmp0Function_Once sync.Once

func strcmp0Function_Set() {
	strcmp0Function_Once.Do(func() {
		strcmp0Function = gi.FunctionInvokerNew("GLib", "strcmp0")
	})
}

// Strcmp0 is a representation of the C type g_strcmp0.
func Strcmp0(str1 string, str2 string) int32 {
	strcmp0Function_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetString(str1)
	inArgs[1].SetString(str2)

	ret := strcmp0Function.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var strcompressFunction *gi.Function
var strcompressFunction_Once sync.Once

func strcompressFunction_Set() {
	strcompressFunction_Once.Do(func() {
		strcompressFunction = gi.FunctionInvokerNew("GLib", "strcompress")
	})
}

// Strcompress is a representation of the C type g_strcompress.
func Strcompress(source string) string {
	strcompressFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetString(source)

	ret := strcompressFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

// UNSUPPORTED : C value 'g_strconcat' : parameter '...' has no type

var strdelimitFunction *gi.Function
var strdelimitFunction_Once sync.Once

func strdelimitFunction_Set() {
	strdelimitFunction_Once.Do(func() {
		strdelimitFunction = gi.FunctionInvokerNew("GLib", "strdelimit")
	})
}

// Strdelimit is a representation of the C type g_strdelimit.
func Strdelimit(string_ string, delimiters string, newDelimiter int8) string {
	strdelimitFunction_Set()

	var inArgs [3]gi.Argument
	inArgs[0].SetString(string_)
	inArgs[1].SetString(delimiters)
	inArgs[2].SetInt8(newDelimiter)

	ret := strdelimitFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

var strdownFunction *gi.Function
var strdownFunction_Once sync.Once

func strdownFunction_Set() {
	strdownFunction_Once.Do(func() {
		strdownFunction = gi.FunctionInvokerNew("GLib", "strdown")
	})
}

// Strdown is a representation of the C type g_strdown.
func Strdown(string_ string) string {
	strdownFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetString(string_)

	ret := strdownFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

var strdupFunction *gi.Function
var strdupFunction_Once sync.Once

func strdupFunction_Set() {
	strdupFunction_Once.Do(func() {
		strdupFunction = gi.FunctionInvokerNew("GLib", "strdup")
	})
}

// Strdup is a representation of the C type g_strdup.
func Strdup(str string) string {
	strdupFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetString(str)

	ret := strdupFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

// UNSUPPORTED : C value 'g_strdup_printf' : parameter '...' has no type

// UNSUPPORTED : C value 'g_strdup_vprintf' : parameter 'args' of type 'va_list' not supported

var strdupvFunction *gi.Function
var strdupvFunction_Once sync.Once

func strdupvFunction_Set() {
	strdupvFunction_Once.Do(func() {
		strdupvFunction = gi.FunctionInvokerNew("GLib", "strdupv")
	})
}

// Strdupv is a representation of the C type g_strdupv.
func Strdupv(strArray string) {
	strdupvFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetString(strArray)

	strdupvFunction.Invoke(inArgs[:], nil)

}

var strerrorFunction *gi.Function
var strerrorFunction_Once sync.Once

func strerrorFunction_Set() {
	strerrorFunction_Once.Do(func() {
		strerrorFunction = gi.FunctionInvokerNew("GLib", "strerror")
	})
}

// Strerror is a representation of the C type g_strerror.
func Strerror(errnum int32) string {
	strerrorFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetInt32(errnum)

	ret := strerrorFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

var strescapeFunction *gi.Function
var strescapeFunction_Once sync.Once

func strescapeFunction_Set() {
	strescapeFunction_Once.Do(func() {
		strescapeFunction = gi.FunctionInvokerNew("GLib", "strescape")
	})
}

// Strescape is a representation of the C type g_strescape.
func Strescape(source string, exceptions string) string {
	strescapeFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetString(source)
	inArgs[1].SetString(exceptions)

	ret := strescapeFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

var strfreevFunction *gi.Function
var strfreevFunction_Once sync.Once

func strfreevFunction_Set() {
	strfreevFunction_Once.Do(func() {
		strfreevFunction = gi.FunctionInvokerNew("GLib", "strfreev")
	})
}

// Strfreev is a representation of the C type g_strfreev.
func Strfreev(strArray string) {
	strfreevFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetString(strArray)

	strfreevFunction.Invoke(inArgs[:], nil)

}

var stringNewFunction *gi.Function
var stringNewFunction_Once sync.Once

func stringNewFunction_Set() {
	stringNewFunction_Once.Do(func() {
		stringNewFunction = gi.FunctionInvokerNew("GLib", "string_new")
	})
}

// StringNew is a representation of the C type g_string_new.
func StringNew(init string) *String {
	stringNewFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetString(init)

	ret := stringNewFunction.Invoke(inArgs[:], nil)

	retGo := &String{native: ret.Pointer()}

	return retGo
}

var stringNewLenFunction *gi.Function
var stringNewLenFunction_Once sync.Once

func stringNewLenFunction_Set() {
	stringNewLenFunction_Once.Do(func() {
		stringNewLenFunction = gi.FunctionInvokerNew("GLib", "string_new_len")
	})
}

// StringNewLen is a representation of the C type g_string_new_len.
func StringNewLen(init string, len int32) *String {
	stringNewLenFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetString(init)
	inArgs[1].SetInt32(len)

	ret := stringNewLenFunction.Invoke(inArgs[:], nil)

	retGo := &String{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_string_sized_new' : parameter 'dfl_size' of type 'gsize' not supported

var stripContextFunction *gi.Function
var stripContextFunction_Once sync.Once

func stripContextFunction_Set() {
	stripContextFunction_Once.Do(func() {
		stripContextFunction = gi.FunctionInvokerNew("GLib", "strip_context")
	})
}

// StripContext is a representation of the C type g_strip_context.
func StripContext(msgid string, msgval string) string {
	stripContextFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetString(msgid)
	inArgs[1].SetString(msgval)

	ret := stripContextFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'g_strjoin' : parameter '...' has no type

var strjoinvFunction *gi.Function
var strjoinvFunction_Once sync.Once

func strjoinvFunction_Set() {
	strjoinvFunction_Once.Do(func() {
		strjoinvFunction = gi.FunctionInvokerNew("GLib", "strjoinv")
	})
}

// Strjoinv is a representation of the C type g_strjoinv.
func Strjoinv(separator string, strArray string) string {
	strjoinvFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetString(separator)
	inArgs[1].SetString(strArray)

	ret := strjoinvFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

// UNSUPPORTED : C value 'g_strlcat' : parameter 'dest_size' of type 'gsize' not supported

// UNSUPPORTED : C value 'g_strlcpy' : parameter 'dest_size' of type 'gsize' not supported

var strncasecmpFunction *gi.Function
var strncasecmpFunction_Once sync.Once

func strncasecmpFunction_Set() {
	strncasecmpFunction_Once.Do(func() {
		strncasecmpFunction = gi.FunctionInvokerNew("GLib", "strncasecmp")
	})
}

// Strncasecmp is a representation of the C type g_strncasecmp.
func Strncasecmp(s1 string, s2 string, n uint32) int32 {
	strncasecmpFunction_Set()

	var inArgs [3]gi.Argument
	inArgs[0].SetString(s1)
	inArgs[1].SetString(s2)
	inArgs[2].SetUint32(n)

	ret := strncasecmpFunction.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

// UNSUPPORTED : C value 'g_strndup' : parameter 'n' of type 'gsize' not supported

// UNSUPPORTED : C value 'g_strnfill' : parameter 'length' of type 'gsize' not supported

var strreverseFunction *gi.Function
var strreverseFunction_Once sync.Once

func strreverseFunction_Set() {
	strreverseFunction_Once.Do(func() {
		strreverseFunction = gi.FunctionInvokerNew("GLib", "strreverse")
	})
}

// Strreverse is a representation of the C type g_strreverse.
func Strreverse(string_ string) string {
	strreverseFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetString(string_)

	ret := strreverseFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

var strrstrFunction *gi.Function
var strrstrFunction_Once sync.Once

func strrstrFunction_Set() {
	strrstrFunction_Once.Do(func() {
		strrstrFunction = gi.FunctionInvokerNew("GLib", "strrstr")
	})
}

// Strrstr is a representation of the C type g_strrstr.
func Strrstr(haystack string, needle string) string {
	strrstrFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetString(haystack)
	inArgs[1].SetString(needle)

	ret := strrstrFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

var strrstrLenFunction *gi.Function
var strrstrLenFunction_Once sync.Once

func strrstrLenFunction_Set() {
	strrstrLenFunction_Once.Do(func() {
		strrstrLenFunction = gi.FunctionInvokerNew("GLib", "strrstr_len")
	})
}

// StrrstrLen is a representation of the C type g_strrstr_len.
func StrrstrLen(haystack string, haystackLen int32, needle string) string {
	strrstrLenFunction_Set()

	var inArgs [3]gi.Argument
	inArgs[0].SetString(haystack)
	inArgs[1].SetInt32(haystackLen)
	inArgs[2].SetString(needle)

	ret := strrstrLenFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

var strsignalFunction *gi.Function
var strsignalFunction_Once sync.Once

func strsignalFunction_Set() {
	strsignalFunction_Once.Do(func() {
		strsignalFunction = gi.FunctionInvokerNew("GLib", "strsignal")
	})
}

// Strsignal is a representation of the C type g_strsignal.
func Strsignal(signum int32) string {
	strsignalFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetInt32(signum)

	ret := strsignalFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

var strsplitFunction *gi.Function
var strsplitFunction_Once sync.Once

func strsplitFunction_Set() {
	strsplitFunction_Once.Do(func() {
		strsplitFunction = gi.FunctionInvokerNew("GLib", "strsplit")
	})
}

// Strsplit is a representation of the C type g_strsplit.
func Strsplit(string_ string, delimiter string, maxTokens int32) {
	strsplitFunction_Set()

	var inArgs [3]gi.Argument
	inArgs[0].SetString(string_)
	inArgs[1].SetString(delimiter)
	inArgs[2].SetInt32(maxTokens)

	strsplitFunction.Invoke(inArgs[:], nil)

}

var strsplitSetFunction *gi.Function
var strsplitSetFunction_Once sync.Once

func strsplitSetFunction_Set() {
	strsplitSetFunction_Once.Do(func() {
		strsplitSetFunction = gi.FunctionInvokerNew("GLib", "strsplit_set")
	})
}

// StrsplitSet is a representation of the C type g_strsplit_set.
func StrsplitSet(string_ string, delimiters string, maxTokens int32) {
	strsplitSetFunction_Set()

	var inArgs [3]gi.Argument
	inArgs[0].SetString(string_)
	inArgs[1].SetString(delimiters)
	inArgs[2].SetInt32(maxTokens)

	strsplitSetFunction.Invoke(inArgs[:], nil)

}

var strstrLenFunction *gi.Function
var strstrLenFunction_Once sync.Once

func strstrLenFunction_Set() {
	strstrLenFunction_Once.Do(func() {
		strstrLenFunction = gi.FunctionInvokerNew("GLib", "strstr_len")
	})
}

// StrstrLen is a representation of the C type g_strstr_len.
func StrstrLen(haystack string, haystackLen int32, needle string) string {
	strstrLenFunction_Set()

	var inArgs [3]gi.Argument
	inArgs[0].SetString(haystack)
	inArgs[1].SetInt32(haystackLen)
	inArgs[2].SetString(needle)

	ret := strstrLenFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

// UNSUPPORTED : C value 'g_strtod' : return type 'gdouble' not supported

var strupFunction *gi.Function
var strupFunction_Once sync.Once

func strupFunction_Set() {
	strupFunction_Once.Do(func() {
		strupFunction = gi.FunctionInvokerNew("GLib", "strup")
	})
}

// Strup is a representation of the C type g_strup.
func Strup(string_ string) string {
	strupFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetString(string_)

	ret := strupFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

// UNSUPPORTED : C value 'g_strv_contains' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_strv_equal' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_strv_get_type' : return type 'GType' not supported

var strvLengthFunction *gi.Function
var strvLengthFunction_Once sync.Once

func strvLengthFunction_Set() {
	strvLengthFunction_Once.Do(func() {
		strvLengthFunction = gi.FunctionInvokerNew("GLib", "strv_length")
	})
}

// StrvLength is a representation of the C type g_strv_length.
func StrvLength(strArray string) uint32 {
	strvLengthFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetString(strArray)

	ret := strvLengthFunction.Invoke(inArgs[:], nil)

	retGo := ret.Uint32()

	return retGo
}

// UNSUPPORTED : C value 'g_test_add_data_func' : parameter 'test_data' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_test_add_data_func_full' : parameter 'test_data' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_test_add_func' : parameter 'test_func' of type 'TestFunc' not supported

// UNSUPPORTED : C value 'g_test_add_vtable' : parameter 'data_size' of type 'gsize' not supported

var testAssertExpectedMessagesInternalFunction *gi.Function
var testAssertExpectedMessagesInternalFunction_Once sync.Once

func testAssertExpectedMessagesInternalFunction_Set() {
	testAssertExpectedMessagesInternalFunction_Once.Do(func() {
		testAssertExpectedMessagesInternalFunction = gi.FunctionInvokerNew("GLib", "test_assert_expected_messages_internal")
	})
}

// TestAssertExpectedMessagesInternal is a representation of the C type g_test_assert_expected_messages_internal.
func TestAssertExpectedMessagesInternal(domain string, file string, line int32, func_ string) {
	testAssertExpectedMessagesInternalFunction_Set()

	var inArgs [4]gi.Argument
	inArgs[0].SetString(domain)
	inArgs[1].SetString(file)
	inArgs[2].SetInt32(line)
	inArgs[3].SetString(func_)

	testAssertExpectedMessagesInternalFunction.Invoke(inArgs[:], nil)

}

var testBugFunction *gi.Function
var testBugFunction_Once sync.Once

func testBugFunction_Set() {
	testBugFunction_Once.Do(func() {
		testBugFunction = gi.FunctionInvokerNew("GLib", "test_bug")
	})
}

// TestBug is a representation of the C type g_test_bug.
func TestBug(bugUriSnippet string) {
	testBugFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetString(bugUriSnippet)

	testBugFunction.Invoke(inArgs[:], nil)

}

var testBugBaseFunction *gi.Function
var testBugBaseFunction_Once sync.Once

func testBugBaseFunction_Set() {
	testBugBaseFunction_Once.Do(func() {
		testBugBaseFunction = gi.FunctionInvokerNew("GLib", "test_bug_base")
	})
}

// TestBugBase is a representation of the C type g_test_bug_base.
func TestBugBase(uriPattern string) {
	testBugBaseFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetString(uriPattern)

	testBugBaseFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_test_build_filename' : parameter 'file_type' of type 'TestFileType' not supported

// UNSUPPORTED : C value 'g_test_create_case' : parameter 'data_size' of type 'gsize' not supported

var testCreateSuiteFunction *gi.Function
var testCreateSuiteFunction_Once sync.Once

func testCreateSuiteFunction_Set() {
	testCreateSuiteFunction_Once.Do(func() {
		testCreateSuiteFunction = gi.FunctionInvokerNew("GLib", "test_create_suite")
	})
}

// TestCreateSuite is a representation of the C type g_test_create_suite.
func TestCreateSuite(suiteName string) *TestSuite {
	testCreateSuiteFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetString(suiteName)

	ret := testCreateSuiteFunction.Invoke(inArgs[:], nil)

	retGo := &TestSuite{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_test_expect_message' : parameter 'log_level' of type 'LogLevelFlags' not supported

var testFailFunction *gi.Function
var testFailFunction_Once sync.Once

func testFailFunction_Set() {
	testFailFunction_Once.Do(func() {
		testFailFunction = gi.FunctionInvokerNew("GLib", "test_fail")
	})
}

// TestFail is a representation of the C type g_test_fail.
func TestFail() {
	testFailFunction_Set()

	testFailFunction.Invoke(nil, nil)

}

// UNSUPPORTED : C value 'g_test_failed' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_test_get_dir' : parameter 'file_type' of type 'TestFileType' not supported

// UNSUPPORTED : C value 'g_test_get_filename' : parameter 'file_type' of type 'TestFileType' not supported

var testGetRootFunction *gi.Function
var testGetRootFunction_Once sync.Once

func testGetRootFunction_Set() {
	testGetRootFunction_Once.Do(func() {
		testGetRootFunction = gi.FunctionInvokerNew("GLib", "test_get_root")
	})
}

// TestGetRoot is a representation of the C type g_test_get_root.
func TestGetRoot() *TestSuite {
	testGetRootFunction_Set()

	ret := testGetRootFunction.Invoke(nil, nil)

	retGo := &TestSuite{native: ret.Pointer()}

	return retGo
}

var testIncompleteFunction *gi.Function
var testIncompleteFunction_Once sync.Once

func testIncompleteFunction_Set() {
	testIncompleteFunction_Once.Do(func() {
		testIncompleteFunction = gi.FunctionInvokerNew("GLib", "test_incomplete")
	})
}

// TestIncomplete is a representation of the C type g_test_incomplete.
func TestIncomplete(msg string) {
	testIncompleteFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetString(msg)

	testIncompleteFunction.Invoke(inArgs[:], nil)

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

var testRandIntFunction *gi.Function
var testRandIntFunction_Once sync.Once

func testRandIntFunction_Set() {
	testRandIntFunction_Once.Do(func() {
		testRandIntFunction = gi.FunctionInvokerNew("GLib", "test_rand_int")
	})
}

// TestRandInt is a representation of the C type g_test_rand_int.
func TestRandInt() int32 {
	testRandIntFunction_Set()

	ret := testRandIntFunction.Invoke(nil, nil)

	retGo := ret.Int32()

	return retGo
}

var testRandIntRangeFunction *gi.Function
var testRandIntRangeFunction_Once sync.Once

func testRandIntRangeFunction_Set() {
	testRandIntRangeFunction_Once.Do(func() {
		testRandIntRangeFunction = gi.FunctionInvokerNew("GLib", "test_rand_int_range")
	})
}

// TestRandIntRange is a representation of the C type g_test_rand_int_range.
func TestRandIntRange(begin int32, end int32) int32 {
	testRandIntRangeFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetInt32(begin)
	inArgs[1].SetInt32(end)

	ret := testRandIntRangeFunction.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var testRunFunction *gi.Function
var testRunFunction_Once sync.Once

func testRunFunction_Set() {
	testRunFunction_Once.Do(func() {
		testRunFunction = gi.FunctionInvokerNew("GLib", "test_run")
	})
}

// TestRun is a representation of the C type g_test_run.
func TestRun() int32 {
	testRunFunction_Set()

	ret := testRunFunction.Invoke(nil, nil)

	retGo := ret.Int32()

	return retGo
}

// UNSUPPORTED : C value 'g_test_run_suite' : parameter 'suite' of type 'TestSuite' not supported

var testSetNonfatalAssertionsFunction *gi.Function
var testSetNonfatalAssertionsFunction_Once sync.Once

func testSetNonfatalAssertionsFunction_Set() {
	testSetNonfatalAssertionsFunction_Once.Do(func() {
		testSetNonfatalAssertionsFunction = gi.FunctionInvokerNew("GLib", "test_set_nonfatal_assertions")
	})
}

// TestSetNonfatalAssertions is a representation of the C type g_test_set_nonfatal_assertions.
func TestSetNonfatalAssertions() {
	testSetNonfatalAssertionsFunction_Set()

	testSetNonfatalAssertionsFunction.Invoke(nil, nil)

}

var testSkipFunction *gi.Function
var testSkipFunction_Once sync.Once

func testSkipFunction_Set() {
	testSkipFunction_Once.Do(func() {
		testSkipFunction = gi.FunctionInvokerNew("GLib", "test_skip")
	})
}

// TestSkip is a representation of the C type g_test_skip.
func TestSkip(msg string) {
	testSkipFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetString(msg)

	testSkipFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_test_subprocess' : return type 'gboolean' not supported

var testSummaryFunction *gi.Function
var testSummaryFunction_Once sync.Once

func testSummaryFunction_Set() {
	testSummaryFunction_Once.Do(func() {
		testSummaryFunction = gi.FunctionInvokerNew("GLib", "test_summary")
	})
}

// TestSummary is a representation of the C type g_test_summary.
func TestSummary(summary string) {
	testSummaryFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetString(summary)

	testSummaryFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_test_timer_elapsed' : return type 'gdouble' not supported

// UNSUPPORTED : C value 'g_test_timer_last' : return type 'gdouble' not supported

var testTimerStartFunction *gi.Function
var testTimerStartFunction_Once sync.Once

func testTimerStartFunction_Set() {
	testTimerStartFunction_Once.Do(func() {
		testTimerStartFunction = gi.FunctionInvokerNew("GLib", "test_timer_start")
	})
}

// TestTimerStart is a representation of the C type g_test_timer_start.
func TestTimerStart() {
	testTimerStartFunction_Set()

	testTimerStartFunction.Invoke(nil, nil)

}

var testTrapAssertionsFunction *gi.Function
var testTrapAssertionsFunction_Once sync.Once

func testTrapAssertionsFunction_Set() {
	testTrapAssertionsFunction_Once.Do(func() {
		testTrapAssertionsFunction = gi.FunctionInvokerNew("GLib", "test_trap_assertions")
	})
}

// TestTrapAssertions is a representation of the C type g_test_trap_assertions.
func TestTrapAssertions(domain string, file string, line int32, func_ string, assertionFlags uint64, pattern string) {
	testTrapAssertionsFunction_Set()

	var inArgs [6]gi.Argument
	inArgs[0].SetString(domain)
	inArgs[1].SetString(file)
	inArgs[2].SetInt32(line)
	inArgs[3].SetString(func_)
	inArgs[4].SetUint64(assertionFlags)
	inArgs[5].SetString(pattern)

	testTrapAssertionsFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_test_trap_fork' : parameter 'test_trap_flags' of type 'TestTrapFlags' not supported

// UNSUPPORTED : C value 'g_test_trap_has_passed' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_test_trap_reached_timeout' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_test_trap_subprocess' : parameter 'test_flags' of type 'TestSubprocessFlags' not supported

// UNSUPPORTED : C value 'g_thread_error_quark' : return type 'Quark' not supported

// UNSUPPORTED : C value 'g_thread_exit' : parameter 'retval' of type 'gpointer' not supported

var threadPoolGetMaxIdleTimeFunction *gi.Function
var threadPoolGetMaxIdleTimeFunction_Once sync.Once

func threadPoolGetMaxIdleTimeFunction_Set() {
	threadPoolGetMaxIdleTimeFunction_Once.Do(func() {
		threadPoolGetMaxIdleTimeFunction = gi.FunctionInvokerNew("GLib", "thread_pool_get_max_idle_time")
	})
}

// ThreadPoolGetMaxIdleTime is a representation of the C type g_thread_pool_get_max_idle_time.
func ThreadPoolGetMaxIdleTime() uint32 {
	threadPoolGetMaxIdleTimeFunction_Set()

	ret := threadPoolGetMaxIdleTimeFunction.Invoke(nil, nil)

	retGo := ret.Uint32()

	return retGo
}

var threadPoolGetMaxUnusedThreadsFunction *gi.Function
var threadPoolGetMaxUnusedThreadsFunction_Once sync.Once

func threadPoolGetMaxUnusedThreadsFunction_Set() {
	threadPoolGetMaxUnusedThreadsFunction_Once.Do(func() {
		threadPoolGetMaxUnusedThreadsFunction = gi.FunctionInvokerNew("GLib", "thread_pool_get_max_unused_threads")
	})
}

// ThreadPoolGetMaxUnusedThreads is a representation of the C type g_thread_pool_get_max_unused_threads.
func ThreadPoolGetMaxUnusedThreads() int32 {
	threadPoolGetMaxUnusedThreadsFunction_Set()

	ret := threadPoolGetMaxUnusedThreadsFunction.Invoke(nil, nil)

	retGo := ret.Int32()

	return retGo
}

var threadPoolGetNumUnusedThreadsFunction *gi.Function
var threadPoolGetNumUnusedThreadsFunction_Once sync.Once

func threadPoolGetNumUnusedThreadsFunction_Set() {
	threadPoolGetNumUnusedThreadsFunction_Once.Do(func() {
		threadPoolGetNumUnusedThreadsFunction = gi.FunctionInvokerNew("GLib", "thread_pool_get_num_unused_threads")
	})
}

// ThreadPoolGetNumUnusedThreads is a representation of the C type g_thread_pool_get_num_unused_threads.
func ThreadPoolGetNumUnusedThreads() uint32 {
	threadPoolGetNumUnusedThreadsFunction_Set()

	ret := threadPoolGetNumUnusedThreadsFunction.Invoke(nil, nil)

	retGo := ret.Uint32()

	return retGo
}

var threadPoolSetMaxIdleTimeFunction *gi.Function
var threadPoolSetMaxIdleTimeFunction_Once sync.Once

func threadPoolSetMaxIdleTimeFunction_Set() {
	threadPoolSetMaxIdleTimeFunction_Once.Do(func() {
		threadPoolSetMaxIdleTimeFunction = gi.FunctionInvokerNew("GLib", "thread_pool_set_max_idle_time")
	})
}

// ThreadPoolSetMaxIdleTime is a representation of the C type g_thread_pool_set_max_idle_time.
func ThreadPoolSetMaxIdleTime(interval uint32) {
	threadPoolSetMaxIdleTimeFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetUint32(interval)

	threadPoolSetMaxIdleTimeFunction.Invoke(inArgs[:], nil)

}

var threadPoolSetMaxUnusedThreadsFunction *gi.Function
var threadPoolSetMaxUnusedThreadsFunction_Once sync.Once

func threadPoolSetMaxUnusedThreadsFunction_Set() {
	threadPoolSetMaxUnusedThreadsFunction_Once.Do(func() {
		threadPoolSetMaxUnusedThreadsFunction = gi.FunctionInvokerNew("GLib", "thread_pool_set_max_unused_threads")
	})
}

// ThreadPoolSetMaxUnusedThreads is a representation of the C type g_thread_pool_set_max_unused_threads.
func ThreadPoolSetMaxUnusedThreads(maxThreads int32) {
	threadPoolSetMaxUnusedThreadsFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetInt32(maxThreads)

	threadPoolSetMaxUnusedThreadsFunction.Invoke(inArgs[:], nil)

}

var threadPoolStopUnusedThreadsFunction *gi.Function
var threadPoolStopUnusedThreadsFunction_Once sync.Once

func threadPoolStopUnusedThreadsFunction_Set() {
	threadPoolStopUnusedThreadsFunction_Once.Do(func() {
		threadPoolStopUnusedThreadsFunction = gi.FunctionInvokerNew("GLib", "thread_pool_stop_unused_threads")
	})
}

// ThreadPoolStopUnusedThreads is a representation of the C type g_thread_pool_stop_unused_threads.
func ThreadPoolStopUnusedThreads() {
	threadPoolStopUnusedThreadsFunction_Set()

	threadPoolStopUnusedThreadsFunction.Invoke(nil, nil)

}

var threadSelfFunction *gi.Function
var threadSelfFunction_Once sync.Once

func threadSelfFunction_Set() {
	threadSelfFunction_Once.Do(func() {
		threadSelfFunction = gi.FunctionInvokerNew("GLib", "thread_self")
	})
}

// ThreadSelf is a representation of the C type g_thread_self.
func ThreadSelf() *Thread {
	threadSelfFunction_Set()

	ret := threadSelfFunction.Invoke(nil, nil)

	retGo := &Thread{native: ret.Pointer()}

	return retGo
}

var threadYieldFunction *gi.Function
var threadYieldFunction_Once sync.Once

func threadYieldFunction_Set() {
	threadYieldFunction_Once.Do(func() {
		threadYieldFunction = gi.FunctionInvokerNew("GLib", "thread_yield")
	})
}

// ThreadYield is a representation of the C type g_thread_yield.
func ThreadYield() {
	threadYieldFunction_Set()

	threadYieldFunction.Invoke(nil, nil)

}

// UNSUPPORTED : C value 'g_time_val_from_iso8601' : parameter 'time_' of type 'TimeVal' not supported

// UNSUPPORTED : C value 'g_timeout_add' : parameter 'function' of type 'SourceFunc' not supported

// UNSUPPORTED : C value 'g_timeout_add_full' : parameter 'function' of type 'SourceFunc' not supported

// UNSUPPORTED : C value 'g_timeout_add_seconds' : parameter 'function' of type 'SourceFunc' not supported

// UNSUPPORTED : C value 'g_timeout_add_seconds_full' : parameter 'function' of type 'SourceFunc' not supported

var timeoutSourceNewFunction *gi.Function
var timeoutSourceNewFunction_Once sync.Once

func timeoutSourceNewFunction_Set() {
	timeoutSourceNewFunction_Once.Do(func() {
		timeoutSourceNewFunction = gi.FunctionInvokerNew("GLib", "timeout_source_new")
	})
}

// TimeoutSourceNew is a representation of the C type g_timeout_source_new.
func TimeoutSourceNew(interval uint32) *Source {
	timeoutSourceNewFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetUint32(interval)

	ret := timeoutSourceNewFunction.Invoke(inArgs[:], nil)

	retGo := &Source{native: ret.Pointer()}

	return retGo
}

var timeoutSourceNewSecondsFunction *gi.Function
var timeoutSourceNewSecondsFunction_Once sync.Once

func timeoutSourceNewSecondsFunction_Set() {
	timeoutSourceNewSecondsFunction_Once.Do(func() {
		timeoutSourceNewSecondsFunction = gi.FunctionInvokerNew("GLib", "timeout_source_new_seconds")
	})
}

// TimeoutSourceNewSeconds is a representation of the C type g_timeout_source_new_seconds.
func TimeoutSourceNewSeconds(interval uint32) *Source {
	timeoutSourceNewSecondsFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetUint32(interval)

	ret := timeoutSourceNewSecondsFunction.Invoke(inArgs[:], nil)

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

var unixSignalSourceNewFunction *gi.Function
var unixSignalSourceNewFunction_Once sync.Once

func unixSignalSourceNewFunction_Set() {
	unixSignalSourceNewFunction_Once.Do(func() {
		unixSignalSourceNewFunction = gi.FunctionInvokerNew("GLib", "unix_signal_source_new")
	})
}

// UnixSignalSourceNew is a representation of the C type g_unix_signal_source_new.
func UnixSignalSourceNew(signum int32) *Source {
	unixSignalSourceNewFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetInt32(signum)

	ret := unixSignalSourceNewFunction.Invoke(inArgs[:], nil)

	retGo := &Source{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_unlink' : parameter 'filename' of type 'filename' not supported

// UNSUPPORTED : C value 'g_unsetenv' : parameter 'variable' of type 'filename' not supported

// UNSUPPORTED : C value 'g_uri_escape_string' : parameter 'allow_utf8' of type 'gboolean' not supported

var uriListExtractUrisFunction *gi.Function
var uriListExtractUrisFunction_Once sync.Once

func uriListExtractUrisFunction_Set() {
	uriListExtractUrisFunction_Once.Do(func() {
		uriListExtractUrisFunction = gi.FunctionInvokerNew("GLib", "uri_list_extract_uris")
	})
}

// UriListExtractUris is a representation of the C type g_uri_list_extract_uris.
func UriListExtractUris(uriList string) {
	uriListExtractUrisFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetString(uriList)

	uriListExtractUrisFunction.Invoke(inArgs[:], nil)

}

var uriParseSchemeFunction *gi.Function
var uriParseSchemeFunction_Once sync.Once

func uriParseSchemeFunction_Set() {
	uriParseSchemeFunction_Once.Do(func() {
		uriParseSchemeFunction = gi.FunctionInvokerNew("GLib", "uri_parse_scheme")
	})
}

// UriParseScheme is a representation of the C type g_uri_parse_scheme.
func UriParseScheme(uri string) string {
	uriParseSchemeFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetString(uri)

	ret := uriParseSchemeFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

var uriUnescapeSegmentFunction *gi.Function
var uriUnescapeSegmentFunction_Once sync.Once

func uriUnescapeSegmentFunction_Set() {
	uriUnescapeSegmentFunction_Once.Do(func() {
		uriUnescapeSegmentFunction = gi.FunctionInvokerNew("GLib", "uri_unescape_segment")
	})
}

// UriUnescapeSegment is a representation of the C type g_uri_unescape_segment.
func UriUnescapeSegment(escapedString string, escapedStringEnd string, illegalCharacters string) string {
	uriUnescapeSegmentFunction_Set()

	var inArgs [3]gi.Argument
	inArgs[0].SetString(escapedString)
	inArgs[1].SetString(escapedStringEnd)
	inArgs[2].SetString(illegalCharacters)

	ret := uriUnescapeSegmentFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

var uriUnescapeStringFunction *gi.Function
var uriUnescapeStringFunction_Once sync.Once

func uriUnescapeStringFunction_Set() {
	uriUnescapeStringFunction_Once.Do(func() {
		uriUnescapeStringFunction = gi.FunctionInvokerNew("GLib", "uri_unescape_string")
	})
}

// UriUnescapeString is a representation of the C type g_uri_unescape_string.
func UriUnescapeString(escapedString string, illegalCharacters string) string {
	uriUnescapeStringFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetString(escapedString)
	inArgs[1].SetString(illegalCharacters)

	ret := uriUnescapeStringFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

var usleepFunction *gi.Function
var usleepFunction_Once sync.Once

func usleepFunction_Set() {
	usleepFunction_Once.Do(func() {
		usleepFunction = gi.FunctionInvokerNew("GLib", "usleep")
	})
}

// Usleep is a representation of the C type g_usleep.
func Usleep(microseconds uint64) {
	usleepFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetUint64(microseconds)

	usleepFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_utf16_to_ucs4' : return type 'gunichar' not supported

var utf16ToUtf8Function *gi.Function
var utf16ToUtf8Function_Once sync.Once

func utf16ToUtf8Function_Set() {
	utf16ToUtf8Function_Once.Do(func() {
		utf16ToUtf8Function = gi.FunctionInvokerNew("GLib", "utf16_to_utf8")
	})
}

// Utf16ToUtf8 is a representation of the C type g_utf16_to_utf8.
func Utf16ToUtf8(str uint16, len int64) (string, int64, int64) {
	utf16ToUtf8Function_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetUint16(str)
	inArgs[1].SetInt64(len)

	var outArgs [2]gi.Argument

	ret := utf16ToUtf8Function.Invoke(inArgs[:], outArgs[:])

	retGo := ret.String(true)
	out0 := outArgs[0].Int64()
	out1 := outArgs[1].Int64()

	return retGo, out0, out1
}

var utf8CasefoldFunction *gi.Function
var utf8CasefoldFunction_Once sync.Once

func utf8CasefoldFunction_Set() {
	utf8CasefoldFunction_Once.Do(func() {
		utf8CasefoldFunction = gi.FunctionInvokerNew("GLib", "utf8_casefold")
	})
}

// Utf8Casefold is a representation of the C type g_utf8_casefold.
func Utf8Casefold(str string, len int32) string {
	utf8CasefoldFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetString(str)
	inArgs[1].SetInt32(len)

	ret := utf8CasefoldFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

var utf8CollateFunction *gi.Function
var utf8CollateFunction_Once sync.Once

func utf8CollateFunction_Set() {
	utf8CollateFunction_Once.Do(func() {
		utf8CollateFunction = gi.FunctionInvokerNew("GLib", "utf8_collate")
	})
}

// Utf8Collate is a representation of the C type g_utf8_collate.
func Utf8Collate(str1 string, str2 string) int32 {
	utf8CollateFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetString(str1)
	inArgs[1].SetString(str2)

	ret := utf8CollateFunction.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var utf8CollateKeyFunction *gi.Function
var utf8CollateKeyFunction_Once sync.Once

func utf8CollateKeyFunction_Set() {
	utf8CollateKeyFunction_Once.Do(func() {
		utf8CollateKeyFunction = gi.FunctionInvokerNew("GLib", "utf8_collate_key")
	})
}

// Utf8CollateKey is a representation of the C type g_utf8_collate_key.
func Utf8CollateKey(str string, len int32) string {
	utf8CollateKeyFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetString(str)
	inArgs[1].SetInt32(len)

	ret := utf8CollateKeyFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

var utf8CollateKeyForFilenameFunction *gi.Function
var utf8CollateKeyForFilenameFunction_Once sync.Once

func utf8CollateKeyForFilenameFunction_Set() {
	utf8CollateKeyForFilenameFunction_Once.Do(func() {
		utf8CollateKeyForFilenameFunction = gi.FunctionInvokerNew("GLib", "utf8_collate_key_for_filename")
	})
}

// Utf8CollateKeyForFilename is a representation of the C type g_utf8_collate_key_for_filename.
func Utf8CollateKeyForFilename(str string, len int32) string {
	utf8CollateKeyForFilenameFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetString(str)
	inArgs[1].SetInt32(len)

	ret := utf8CollateKeyForFilenameFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

var utf8FindNextCharFunction *gi.Function
var utf8FindNextCharFunction_Once sync.Once

func utf8FindNextCharFunction_Set() {
	utf8FindNextCharFunction_Once.Do(func() {
		utf8FindNextCharFunction = gi.FunctionInvokerNew("GLib", "utf8_find_next_char")
	})
}

// Utf8FindNextChar is a representation of the C type g_utf8_find_next_char.
func Utf8FindNextChar(p string, end string) string {
	utf8FindNextCharFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetString(p)
	inArgs[1].SetString(end)

	ret := utf8FindNextCharFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

var utf8FindPrevCharFunction *gi.Function
var utf8FindPrevCharFunction_Once sync.Once

func utf8FindPrevCharFunction_Set() {
	utf8FindPrevCharFunction_Once.Do(func() {
		utf8FindPrevCharFunction = gi.FunctionInvokerNew("GLib", "utf8_find_prev_char")
	})
}

// Utf8FindPrevChar is a representation of the C type g_utf8_find_prev_char.
func Utf8FindPrevChar(str string, p string) string {
	utf8FindPrevCharFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetString(str)
	inArgs[1].SetString(p)

	ret := utf8FindPrevCharFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'g_utf8_get_char' : return type 'gunichar' not supported

// UNSUPPORTED : C value 'g_utf8_get_char_validated' : return type 'gunichar' not supported

var utf8MakeValidFunction *gi.Function
var utf8MakeValidFunction_Once sync.Once

func utf8MakeValidFunction_Set() {
	utf8MakeValidFunction_Once.Do(func() {
		utf8MakeValidFunction = gi.FunctionInvokerNew("GLib", "utf8_make_valid")
	})
}

// Utf8MakeValid is a representation of the C type g_utf8_make_valid.
func Utf8MakeValid(str string, len int32) string {
	utf8MakeValidFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetString(str)
	inArgs[1].SetInt32(len)

	ret := utf8MakeValidFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

// UNSUPPORTED : C value 'g_utf8_normalize' : parameter 'mode' of type 'NormalizeMode' not supported

var utf8OffsetToPointerFunction *gi.Function
var utf8OffsetToPointerFunction_Once sync.Once

func utf8OffsetToPointerFunction_Set() {
	utf8OffsetToPointerFunction_Once.Do(func() {
		utf8OffsetToPointerFunction = gi.FunctionInvokerNew("GLib", "utf8_offset_to_pointer")
	})
}

// Utf8OffsetToPointer is a representation of the C type g_utf8_offset_to_pointer.
func Utf8OffsetToPointer(str string, offset int64) string {
	utf8OffsetToPointerFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetString(str)
	inArgs[1].SetInt64(offset)

	ret := utf8OffsetToPointerFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

var utf8PointerToOffsetFunction *gi.Function
var utf8PointerToOffsetFunction_Once sync.Once

func utf8PointerToOffsetFunction_Set() {
	utf8PointerToOffsetFunction_Once.Do(func() {
		utf8PointerToOffsetFunction = gi.FunctionInvokerNew("GLib", "utf8_pointer_to_offset")
	})
}

// Utf8PointerToOffset is a representation of the C type g_utf8_pointer_to_offset.
func Utf8PointerToOffset(str string, pos string) int64 {
	utf8PointerToOffsetFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetString(str)
	inArgs[1].SetString(pos)

	ret := utf8PointerToOffsetFunction.Invoke(inArgs[:], nil)

	retGo := ret.Int64()

	return retGo
}

var utf8PrevCharFunction *gi.Function
var utf8PrevCharFunction_Once sync.Once

func utf8PrevCharFunction_Set() {
	utf8PrevCharFunction_Once.Do(func() {
		utf8PrevCharFunction = gi.FunctionInvokerNew("GLib", "utf8_prev_char")
	})
}

// Utf8PrevChar is a representation of the C type g_utf8_prev_char.
func Utf8PrevChar(p string) string {
	utf8PrevCharFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetString(p)

	ret := utf8PrevCharFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'g_utf8_strchr' : parameter 'c' of type 'gunichar' not supported

var utf8StrdownFunction *gi.Function
var utf8StrdownFunction_Once sync.Once

func utf8StrdownFunction_Set() {
	utf8StrdownFunction_Once.Do(func() {
		utf8StrdownFunction = gi.FunctionInvokerNew("GLib", "utf8_strdown")
	})
}

// Utf8Strdown is a representation of the C type g_utf8_strdown.
func Utf8Strdown(str string, len int32) string {
	utf8StrdownFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetString(str)
	inArgs[1].SetInt32(len)

	ret := utf8StrdownFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

var utf8StrlenFunction *gi.Function
var utf8StrlenFunction_Once sync.Once

func utf8StrlenFunction_Set() {
	utf8StrlenFunction_Once.Do(func() {
		utf8StrlenFunction = gi.FunctionInvokerNew("GLib", "utf8_strlen")
	})
}

// Utf8Strlen is a representation of the C type g_utf8_strlen.
func Utf8Strlen(p string, max int32) int64 {
	utf8StrlenFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetString(p)
	inArgs[1].SetInt32(max)

	ret := utf8StrlenFunction.Invoke(inArgs[:], nil)

	retGo := ret.Int64()

	return retGo
}

// UNSUPPORTED : C value 'g_utf8_strncpy' : parameter 'n' of type 'gsize' not supported

// UNSUPPORTED : C value 'g_utf8_strrchr' : parameter 'c' of type 'gunichar' not supported

var utf8StrreverseFunction *gi.Function
var utf8StrreverseFunction_Once sync.Once

func utf8StrreverseFunction_Set() {
	utf8StrreverseFunction_Once.Do(func() {
		utf8StrreverseFunction = gi.FunctionInvokerNew("GLib", "utf8_strreverse")
	})
}

// Utf8Strreverse is a representation of the C type g_utf8_strreverse.
func Utf8Strreverse(str string, len int32) string {
	utf8StrreverseFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetString(str)
	inArgs[1].SetInt32(len)

	ret := utf8StrreverseFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

var utf8StrupFunction *gi.Function
var utf8StrupFunction_Once sync.Once

func utf8StrupFunction_Set() {
	utf8StrupFunction_Once.Do(func() {
		utf8StrupFunction = gi.FunctionInvokerNew("GLib", "utf8_strup")
	})
}

// Utf8Strup is a representation of the C type g_utf8_strup.
func Utf8Strup(str string, len int32) string {
	utf8StrupFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetString(str)
	inArgs[1].SetInt32(len)

	ret := utf8StrupFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

var utf8SubstringFunction *gi.Function
var utf8SubstringFunction_Once sync.Once

func utf8SubstringFunction_Set() {
	utf8SubstringFunction_Once.Do(func() {
		utf8SubstringFunction = gi.FunctionInvokerNew("GLib", "utf8_substring")
	})
}

// Utf8Substring is a representation of the C type g_utf8_substring.
func Utf8Substring(str string, startPos int64, endPos int64) string {
	utf8SubstringFunction_Set()

	var inArgs [3]gi.Argument
	inArgs[0].SetString(str)
	inArgs[1].SetInt64(startPos)
	inArgs[2].SetInt64(endPos)

	ret := utf8SubstringFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

// UNSUPPORTED : C value 'g_utf8_to_ucs4' : return type 'gunichar' not supported

// UNSUPPORTED : C value 'g_utf8_to_ucs4_fast' : return type 'gunichar' not supported

var utf8ToUtf16Function *gi.Function
var utf8ToUtf16Function_Once sync.Once

func utf8ToUtf16Function_Set() {
	utf8ToUtf16Function_Once.Do(func() {
		utf8ToUtf16Function = gi.FunctionInvokerNew("GLib", "utf8_to_utf16")
	})
}

// Utf8ToUtf16 is a representation of the C type g_utf8_to_utf16.
func Utf8ToUtf16(str string, len int64) (uint16, int64, int64) {
	utf8ToUtf16Function_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetString(str)
	inArgs[1].SetInt64(len)

	var outArgs [2]gi.Argument

	ret := utf8ToUtf16Function.Invoke(inArgs[:], outArgs[:])

	retGo := ret.Uint16()
	out0 := outArgs[0].Int64()
	out1 := outArgs[1].Int64()

	return retGo, out0, out1
}

// UNSUPPORTED : C value 'g_utf8_validate' : parameter 'str' has no type

// UNSUPPORTED : C value 'g_utf8_validate_len' : parameter 'str' has no type

// UNSUPPORTED : C value 'g_uuid_string_is_valid' : return type 'gboolean' not supported

var uuidStringRandomFunction *gi.Function
var uuidStringRandomFunction_Once sync.Once

func uuidStringRandomFunction_Set() {
	uuidStringRandomFunction_Once.Do(func() {
		uuidStringRandomFunction = gi.FunctionInvokerNew("GLib", "uuid_string_random")
	})
}

// UuidStringRandom is a representation of the C type g_uuid_string_random.
func UuidStringRandom() string {
	uuidStringRandomFunction_Set()

	ret := uuidStringRandomFunction.Invoke(nil, nil)

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

var variantTypeCheckedFunction *gi.Function
var variantTypeCheckedFunction_Once sync.Once

func variantTypeCheckedFunction_Set() {
	variantTypeCheckedFunction_Once.Do(func() {
		variantTypeCheckedFunction = gi.FunctionInvokerNew("GLib", "variant_type_checked_")
	})
}

// VariantTypeChecked is a representation of the C type g_variant_type_checked_.
func VariantTypeChecked(arg0 string) *VariantType {
	variantTypeCheckedFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetString(arg0)

	ret := variantTypeCheckedFunction.Invoke(inArgs[:], nil)

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

var warnMessageFunction *gi.Function
var warnMessageFunction_Once sync.Once

func warnMessageFunction_Set() {
	warnMessageFunction_Once.Do(func() {
		warnMessageFunction = gi.FunctionInvokerNew("GLib", "warn_message")
	})
}

// WarnMessage is a representation of the C type g_warn_message.
func WarnMessage(domain string, file string, line int32, func_ string, warnexpr string) {
	warnMessageFunction_Set()

	var inArgs [5]gi.Argument
	inArgs[0].SetString(domain)
	inArgs[1].SetString(file)
	inArgs[2].SetInt32(line)
	inArgs[3].SetString(func_)
	inArgs[4].SetString(warnexpr)

	warnMessageFunction.Invoke(inArgs[:], nil)

}
