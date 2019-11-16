// Code generated - DO NOT EDIT.

package glib

import gi "github.com/pekim/gobbi/internal/gi"

// UNSUPPORTED : C value 'g_access' : parameter 'filename' of type 'filename' not supported

var asciiDigitValueInvoker *gi.Function

// AsciiDigitValue is a representation of the C type g_ascii_digit_value.
func AsciiDigitValue(c int8) int32 {
	if asciiDigitValueInvoker == nil {
		asciiDigitValueInvoker = gi.FunctionInvokerNew("GLib", "ascii_digit_value")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetInt8(c)

	ret := asciiDigitValueInvoker.Invoke(inArgs[:])
	return ret.Int32()
}

// UNSUPPORTED : C value 'g_ascii_dtostr' : parameter 'd' of type 'gdouble' not supported

// UNSUPPORTED : C value 'g_ascii_formatd' : parameter 'd' of type 'gdouble' not supported

var asciiStrcasecmpInvoker *gi.Function

// AsciiStrcasecmp is a representation of the C type g_ascii_strcasecmp.
func AsciiStrcasecmp(s1 string, s2 string) int32 {
	if asciiStrcasecmpInvoker == nil {
		asciiStrcasecmpInvoker = gi.FunctionInvokerNew("GLib", "ascii_strcasecmp")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetString(s1)
	inArgs[1].SetString(s2)

	ret := asciiStrcasecmpInvoker.Invoke(inArgs[:])
	return ret.Int32()
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

	ret := asciiStrdownInvoker.Invoke(inArgs[:])
	return ret.String(true)
}

// UNSUPPORTED : C value 'g_ascii_string_to_signed' : parameter 'out_num' with direction 'out' not supported

// UNSUPPORTED : C value 'g_ascii_string_to_unsigned' : parameter 'out_num' with direction 'out' not supported

// UNSUPPORTED : C value 'g_ascii_strncasecmp' : parameter 'n' of type 'gsize' not supported

// UNSUPPORTED : C value 'g_ascii_strtod' : parameter 'endptr' with direction 'out' not supported

// UNSUPPORTED : C value 'g_ascii_strtoll' : parameter 'endptr' with direction 'out' not supported

// UNSUPPORTED : C value 'g_ascii_strtoull' : parameter 'endptr' with direction 'out' not supported

var asciiStrupInvoker *gi.Function

// AsciiStrup is a representation of the C type g_ascii_strup.
func AsciiStrup(str string, len int32) string {
	if asciiStrupInvoker == nil {
		asciiStrupInvoker = gi.FunctionInvokerNew("GLib", "ascii_strup")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetString(str)
	inArgs[1].SetInt32(len)

	ret := asciiStrupInvoker.Invoke(inArgs[:])
	return ret.String(true)
}

var asciiTolowerInvoker *gi.Function

// AsciiTolower is a representation of the C type g_ascii_tolower.
func AsciiTolower(asciiStrtollc int8) int8 {
	if asciiTolowerInvoker == nil {
		asciiTolowerInvoker = gi.FunctionInvokerNew("GLib", "ascii_tolower")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetInt8(asciiStrtollc)

	ret := asciiTolowerInvoker.Invoke(inArgs[:])
	return ret.Int8()
}

var asciiToupperInvoker *gi.Function

// AsciiToupper is a representation of the C type g_ascii_toupper.
func AsciiToupper(c int8) int8 {
	if asciiToupperInvoker == nil {
		asciiToupperInvoker = gi.FunctionInvokerNew("GLib", "ascii_toupper")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetInt8(c)

	ret := asciiToupperInvoker.Invoke(inArgs[:])
	return ret.Int8()
}

var asciiXdigitValueInvoker *gi.Function

// AsciiXdigitValue is a representation of the C type g_ascii_xdigit_value.
func AsciiXdigitValue(c int8) int32 {
	if asciiXdigitValueInvoker == nil {
		asciiXdigitValueInvoker = gi.FunctionInvokerNew("GLib", "ascii_xdigit_value")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetInt8(c)

	ret := asciiXdigitValueInvoker.Invoke(inArgs[:])
	return ret.Int32()
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

	assertWarningInvoker.Invoke(inArgs[:])
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

	assertionMessageInvoker.Invoke(inArgs[:])
}

// UNSUPPORTED : C value 'g_assertion_message_cmpnum' : parameter 'arg1' of type 'long double' not supported

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

	assertionMessageCmpstrInvoker.Invoke(inArgs[:])
}

// UNSUPPORTED : C value 'g_assertion_message_error' : parameter 'error' of type 'Error' not supported

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

	assertionMessageExprInvoker.Invoke(inArgs[:])
}

// UNSUPPORTED : C value 'g_atexit' : parameter 'func' of type 'VoidFunc' not supported

var atomicIntAddInvoker *gi.Function

// AtomicIntAdd is a representation of the C type g_atomic_int_add.
func AtomicIntAdd(atomic int32, val int32) int32 {
	if atomicIntAddInvoker == nil {
		atomicIntAddInvoker = gi.FunctionInvokerNew("GLib", "atomic_int_add")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetInt32(atomic)
	inArgs[1].SetInt32(val)

	ret := atomicIntAddInvoker.Invoke(inArgs[:])
	return ret.Int32()
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

	ret := atomicIntAndInvoker.Invoke(inArgs[:])
	return ret.Uint32()
}

// UNSUPPORTED : C value 'g_atomic_int_compare_and_exchange' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_atomic_int_dec_and_test' : return type 'gboolean' not supported

var atomicIntExchangeAndAddInvoker *gi.Function

// AtomicIntExchangeAndAdd is a representation of the C type g_atomic_int_exchange_and_add.
func AtomicIntExchangeAndAdd(atomic int32, val int32) int32 {
	if atomicIntExchangeAndAddInvoker == nil {
		atomicIntExchangeAndAddInvoker = gi.FunctionInvokerNew("GLib", "atomic_int_exchange_and_add")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetInt32(atomic)
	inArgs[1].SetInt32(val)

	ret := atomicIntExchangeAndAddInvoker.Invoke(inArgs[:])
	return ret.Int32()
}

var atomicIntGetInvoker *gi.Function

// AtomicIntGet is a representation of the C type g_atomic_int_get.
func AtomicIntGet(atomic int32) int32 {
	if atomicIntGetInvoker == nil {
		atomicIntGetInvoker = gi.FunctionInvokerNew("GLib", "atomic_int_get")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetInt32(atomic)

	ret := atomicIntGetInvoker.Invoke(inArgs[:])
	return ret.Int32()
}

var atomicIntIncInvoker *gi.Function

// AtomicIntInc is a representation of the C type g_atomic_int_inc.
func AtomicIntInc(atomic int32) {
	if atomicIntIncInvoker == nil {
		atomicIntIncInvoker = gi.FunctionInvokerNew("GLib", "atomic_int_inc")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetInt32(atomic)

	atomicIntIncInvoker.Invoke(inArgs[:])
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

	ret := atomicIntOrInvoker.Invoke(inArgs[:])
	return ret.Uint32()
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

	atomicIntSetInvoker.Invoke(inArgs[:])
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

	ret := atomicIntXorInvoker.Invoke(inArgs[:])
	return ret.Uint32()
}

// UNSUPPORTED : C value 'g_atomic_pointer_add' : parameter 'atomic' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_atomic_pointer_and' : parameter 'atomic' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_atomic_pointer_compare_and_exchange' : parameter 'atomic' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_atomic_pointer_get' : parameter 'atomic' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_atomic_pointer_or' : parameter 'atomic' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_atomic_pointer_set' : parameter 'atomic' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_atomic_pointer_xor' : parameter 'atomic' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_base64_decode' : parameter 'out_len' of type 'gsize' not supported

// UNSUPPORTED : C value 'g_base64_decode_inplace' : parameter 'text' has no type

// UNSUPPORTED : C value 'g_base64_decode_step' : parameter 'in' has no type

// UNSUPPORTED : C value 'g_base64_encode' : parameter 'data' has no type

// UNSUPPORTED : C value 'g_base64_encode_close' : parameter 'break_lines' of type 'gboolean' not supported

// UNSUPPORTED : C value 'g_base64_encode_step' : parameter 'in' has no type

// UNSUPPORTED : C value 'g_basename' : parameter 'file_name' of type 'filename' not supported

var bitLockInvoker *gi.Function

// BitLock is a representation of the C type g_bit_lock.
func BitLock(address int32, lockBit int32) {
	if bitLockInvoker == nil {
		bitLockInvoker = gi.FunctionInvokerNew("GLib", "bit_lock")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetInt32(address)
	inArgs[1].SetInt32(lockBit)

	bitLockInvoker.Invoke(inArgs[:])
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

	ret := bitNthLsfInvoker.Invoke(inArgs[:])
	return ret.Int32()
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

	ret := bitNthMsfInvoker.Invoke(inArgs[:])
	return ret.Int32()
}

var bitStorageInvoker *gi.Function

// BitStorage is a representation of the C type g_bit_storage.
func BitStorage(number uint64) uint32 {
	if bitStorageInvoker == nil {
		bitStorageInvoker = gi.FunctionInvokerNew("GLib", "bit_storage")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetUint64(number)

	ret := bitStorageInvoker.Invoke(inArgs[:])
	return ret.Uint32()
}

// UNSUPPORTED : C value 'g_bit_trylock' : return type 'gboolean' not supported

var bitUnlockInvoker *gi.Function

// BitUnlock is a representation of the C type g_bit_unlock.
func BitUnlock(address int32, lockBit int32) {
	if bitUnlockInvoker == nil {
		bitUnlockInvoker = gi.FunctionInvokerNew("GLib", "bit_unlock")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetInt32(address)
	inArgs[1].SetInt32(lockBit)

	bitUnlockInvoker.Invoke(inArgs[:])
}

// UNSUPPORTED : C value 'g_bookmark_file_error_quark' : return type 'Quark' not supported

// UNSUPPORTED : C value 'g_build_filename' : parameter 'first_element' of type 'filename' not supported

// UNSUPPORTED : C value 'g_build_filename_valist' : parameter 'first_element' of type 'filename' not supported

// UNSUPPORTED : C value 'g_build_filenamev' : parameter 'args' has no type

// UNSUPPORTED : C value 'g_build_path' : parameter 'separator' of type 'filename' not supported

// UNSUPPORTED : C value 'g_build_pathv' : parameter 'args' has no type

// UNSUPPORTED : C value 'g_byte_array_free' : parameter 'array' has no type

// UNSUPPORTED : C value 'g_byte_array_free_to_bytes' : parameter 'array' has no type

var byteArrayNewInvoker *gi.Function

// ByteArrayNew is a representation of the C type g_byte_array_new.
func ByteArrayNew() {
	if byteArrayNewInvoker == nil {
		byteArrayNewInvoker = gi.FunctionInvokerNew("GLib", "byte_array_new")
	}

	byteArrayNewInvoker.Invoke(nil)
}

// UNSUPPORTED : C value 'g_byte_array_new_take' : parameter 'data' has no type

// UNSUPPORTED : C value 'g_byte_array_unref' : parameter 'array' has no type

// UNSUPPORTED : C value 'g_chdir' : parameter 'path' of type 'filename' not supported

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

	ret := checkVersionInvoker.Invoke(inArgs[:])
	return ret.String(false)
}

// UNSUPPORTED : C value 'g_checksum_type_get_length' : parameter 'checksum_type' of type 'ChecksumType' not supported

// UNSUPPORTED : C value 'g_child_watch_add' : parameter 'pid' of type 'Pid' not supported

// UNSUPPORTED : C value 'g_child_watch_add_full' : parameter 'pid' of type 'Pid' not supported

// UNSUPPORTED : C value 'g_child_watch_source_new' : parameter 'pid' of type 'Pid' not supported

var clearErrorInvoker *gi.Function

// ClearError is a representation of the C type g_clear_error.
func ClearError() {
	if clearErrorInvoker == nil {
		clearErrorInvoker = gi.FunctionInvokerNew("GLib", "clear_error")
	}

	clearErrorInvoker.Invoke(nil)
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

	ret := dcgettextInvoker.Invoke(inArgs[:])
	return ret.String(false)
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

	ret := dgettextInvoker.Invoke(inArgs[:])
	return ret.String(false)
}

// UNSUPPORTED : C value 'g_dir_make_tmp' : parameter 'tmpl' of type 'filename' not supported

// UNSUPPORTED : C value 'g_direct_equal' : parameter 'v1' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_direct_hash' : parameter 'v' of type 'gpointer' not supported

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

	ret := dngettextInvoker.Invoke(inArgs[:])
	return ret.String(false)
}

// UNSUPPORTED : C value 'g_double_equal' : parameter 'v1' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_double_hash' : parameter 'v' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_dpgettext' : parameter 'msgidoffset' of type 'gsize' not supported

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

	ret := dpgettext2Invoker.Invoke(inArgs[:])
	return ret.String(false)
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

// UNSUPPORTED : C value 'g_filename_from_uri' : parameter 'hostname' with direction 'out' not supported

// UNSUPPORTED : C value 'g_filename_from_utf8' : parameter 'bytes_read' of type 'gsize' not supported

// UNSUPPORTED : C value 'g_filename_to_uri' : parameter 'filename' of type 'filename' not supported

// UNSUPPORTED : C value 'g_filename_to_utf8' : parameter 'opsysstring' of type 'filename' not supported

// UNSUPPORTED : C value 'g_find_program_in_path' : parameter 'program' of type 'filename' not supported

var formatSizeInvoker *gi.Function

// FormatSize is a representation of the C type g_format_size.
func FormatSize(size uint64) string {
	if formatSizeInvoker == nil {
		formatSizeInvoker = gi.FunctionInvokerNew("GLib", "format_size")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetUint64(size)

	ret := formatSizeInvoker.Invoke(inArgs[:])
	return ret.String(true)
}

var formatSizeForDisplayInvoker *gi.Function

// FormatSizeForDisplay is a representation of the C type g_format_size_for_display.
func FormatSizeForDisplay(size int64) string {
	if formatSizeForDisplayInvoker == nil {
		formatSizeForDisplayInvoker = gi.FunctionInvokerNew("GLib", "format_size_for_display")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetInt64(size)

	ret := formatSizeForDisplayInvoker.Invoke(inArgs[:])
	return ret.String(true)
}

// UNSUPPORTED : C value 'g_format_size_full' : parameter 'flags' of type 'FormatSizeFlags' not supported

// UNSUPPORTED : C value 'g_fprintf' : parameter 'file' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_free' : parameter 'mem' of type 'gpointer' not supported

var getApplicationNameInvoker *gi.Function

// GetApplicationName is a representation of the C type g_get_application_name.
func GetApplicationName() string {
	if getApplicationNameInvoker == nil {
		getApplicationNameInvoker = gi.FunctionInvokerNew("GLib", "get_application_name")
	}

	ret := getApplicationNameInvoker.Invoke(nil)
	return ret.String(false)
}

// UNSUPPORTED : C value 'g_get_charset' : parameter 'charset' with direction 'out' not supported

var getCodesetInvoker *gi.Function

// GetCodeset is a representation of the C type g_get_codeset.
func GetCodeset() string {
	if getCodesetInvoker == nil {
		getCodesetInvoker = gi.FunctionInvokerNew("GLib", "get_codeset")
	}

	ret := getCodesetInvoker.Invoke(nil)
	return ret.String(true)
}

// UNSUPPORTED : C value 'g_get_current_dir' : return type 'filename' not supported

// UNSUPPORTED : C value 'g_get_current_time' : parameter 'result' of type 'TimeVal' not supported

var getEnvironInvoker *gi.Function

// GetEnviron is a representation of the C type g_get_environ.
func GetEnviron() {
	if getEnvironInvoker == nil {
		getEnvironInvoker = gi.FunctionInvokerNew("GLib", "get_environ")
	}

	getEnvironInvoker.Invoke(nil)
}

// UNSUPPORTED : C value 'g_get_filename_charsets' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_get_home_dir' : return type 'filename' not supported

var getHostNameInvoker *gi.Function

// GetHostName is a representation of the C type g_get_host_name.
func GetHostName() string {
	if getHostNameInvoker == nil {
		getHostNameInvoker = gi.FunctionInvokerNew("GLib", "get_host_name")
	}

	ret := getHostNameInvoker.Invoke(nil)
	return ret.String(false)
}

var getLanguageNamesInvoker *gi.Function

// GetLanguageNames is a representation of the C type g_get_language_names.
func GetLanguageNames() {
	if getLanguageNamesInvoker == nil {
		getLanguageNamesInvoker = gi.FunctionInvokerNew("GLib", "get_language_names")
	}

	getLanguageNamesInvoker.Invoke(nil)
}

var getLocaleVariantsInvoker *gi.Function

// GetLocaleVariants is a representation of the C type g_get_locale_variants.
func GetLocaleVariants(locale string) {
	if getLocaleVariantsInvoker == nil {
		getLocaleVariantsInvoker = gi.FunctionInvokerNew("GLib", "get_locale_variants")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetString(locale)

	getLocaleVariantsInvoker.Invoke(inArgs[:])
}

var getMonotonicTimeInvoker *gi.Function

// GetMonotonicTime is a representation of the C type g_get_monotonic_time.
func GetMonotonicTime() int64 {
	if getMonotonicTimeInvoker == nil {
		getMonotonicTimeInvoker = gi.FunctionInvokerNew("GLib", "get_monotonic_time")
	}

	ret := getMonotonicTimeInvoker.Invoke(nil)
	return ret.Int64()
}

var getNumProcessorsInvoker *gi.Function

// GetNumProcessors is a representation of the C type g_get_num_processors.
func GetNumProcessors() uint32 {
	if getNumProcessorsInvoker == nil {
		getNumProcessorsInvoker = gi.FunctionInvokerNew("GLib", "get_num_processors")
	}

	ret := getNumProcessorsInvoker.Invoke(nil)
	return ret.Uint32()
}

var getPrgnameInvoker *gi.Function

// GetPrgname is a representation of the C type g_get_prgname.
func GetPrgname() string {
	if getPrgnameInvoker == nil {
		getPrgnameInvoker = gi.FunctionInvokerNew("GLib", "get_prgname")
	}

	ret := getPrgnameInvoker.Invoke(nil)
	return ret.String(false)
}

// UNSUPPORTED : C value 'g_get_real_name' : return type 'filename' not supported

var getRealTimeInvoker *gi.Function

// GetRealTime is a representation of the C type g_get_real_time.
func GetRealTime() int64 {
	if getRealTimeInvoker == nil {
		getRealTimeInvoker = gi.FunctionInvokerNew("GLib", "get_real_time")
	}

	ret := getRealTimeInvoker.Invoke(nil)
	return ret.Int64()
}

var getSystemConfigDirsInvoker *gi.Function

// GetSystemConfigDirs is a representation of the C type g_get_system_config_dirs.
func GetSystemConfigDirs() {
	if getSystemConfigDirsInvoker == nil {
		getSystemConfigDirsInvoker = gi.FunctionInvokerNew("GLib", "get_system_config_dirs")
	}

	getSystemConfigDirsInvoker.Invoke(nil)
}

var getSystemDataDirsInvoker *gi.Function

// GetSystemDataDirs is a representation of the C type g_get_system_data_dirs.
func GetSystemDataDirs() {
	if getSystemDataDirsInvoker == nil {
		getSystemDataDirsInvoker = gi.FunctionInvokerNew("GLib", "get_system_data_dirs")
	}

	getSystemDataDirsInvoker.Invoke(nil)
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

var hostnameToAsciiInvoker *gi.Function

// HostnameToAscii is a representation of the C type g_hostname_to_ascii.
func HostnameToAscii(hostname string) string {
	if hostnameToAsciiInvoker == nil {
		hostnameToAsciiInvoker = gi.FunctionInvokerNew("GLib", "hostname_to_ascii")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetString(hostname)

	ret := hostnameToAsciiInvoker.Invoke(inArgs[:])
	return ret.String(true)
}

var hostnameToUnicodeInvoker *gi.Function

// HostnameToUnicode is a representation of the C type g_hostname_to_unicode.
func HostnameToUnicode(hostname string) string {
	if hostnameToUnicodeInvoker == nil {
		hostnameToUnicodeInvoker = gi.FunctionInvokerNew("GLib", "hostname_to_unicode")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetString(hostname)

	ret := hostnameToUnicodeInvoker.Invoke(inArgs[:])
	return ret.String(true)
}

// UNSUPPORTED : C value 'g_iconv' : parameter 'converter' of type 'IConv' not supported

// UNSUPPORTED : C value 'g_iconv_open' : return type 'IConv' not supported

// UNSUPPORTED : C value 'g_idle_add' : parameter 'function' of type 'SourceFunc' not supported

// UNSUPPORTED : C value 'g_idle_add_full' : parameter 'function' of type 'SourceFunc' not supported

// UNSUPPORTED : C value 'g_idle_remove_by_data' : parameter 'data' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_idle_source_new' : return type 'Source' not supported

// UNSUPPORTED : C value 'g_int64_equal' : parameter 'v1' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_int64_hash' : parameter 'v' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_int_equal' : parameter 'v1' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_int_hash' : parameter 'v' of type 'gpointer' not supported

var internStaticStringInvoker *gi.Function

// InternStaticString is a representation of the C type g_intern_static_string.
func InternStaticString(string_ string) string {
	if internStaticStringInvoker == nil {
		internStaticStringInvoker = gi.FunctionInvokerNew("GLib", "intern_static_string")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetString(string_)

	ret := internStaticStringInvoker.Invoke(inArgs[:])
	return ret.String(false)
}

var internStringInvoker *gi.Function

// InternString is a representation of the C type g_intern_string.
func InternString(string_ string) string {
	if internStringInvoker == nil {
		internStringInvoker = gi.FunctionInvokerNew("GLib", "intern_string")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetString(string_)

	ret := internStringInvoker.Invoke(inArgs[:])
	return ret.String(false)
}

// UNSUPPORTED : C value 'g_io_add_watch' : parameter 'channel' of type 'IOChannel' not supported

// UNSUPPORTED : C value 'g_io_add_watch_full' : parameter 'channel' of type 'IOChannel' not supported

// UNSUPPORTED : C value 'g_io_channel_error_from_errno' : return type 'IOChannelError' not supported

// UNSUPPORTED : C value 'g_io_channel_error_quark' : return type 'Quark' not supported

// UNSUPPORTED : C value 'g_io_create_watch' : parameter 'channel' of type 'IOChannel' not supported

// UNSUPPORTED : C value 'g_key_file_error_quark' : return type 'Quark' not supported

var listenvInvoker *gi.Function

// Listenv is a representation of the C type g_listenv.
func Listenv() {
	if listenvInvoker == nil {
		listenvInvoker = gi.FunctionInvokerNew("GLib", "listenv")
	}

	listenvInvoker.Invoke(nil)
}

// UNSUPPORTED : C value 'g_locale_from_utf8' : parameter 'bytes_read' of type 'gsize' not supported

// UNSUPPORTED : C value 'g_locale_to_utf8' : parameter 'opsysstring' has no type

// UNSUPPORTED : C value 'g_log' : parameter 'log_level' of type 'LogLevelFlags' not supported

// UNSUPPORTED : C value 'g_log_default_handler' : parameter 'log_level' of type 'LogLevelFlags' not supported

var logRemoveHandlerInvoker *gi.Function

// LogRemoveHandler is a representation of the C type g_log_remove_handler.
func LogRemoveHandler(logDomain string, handlerId uint32) {
	if logRemoveHandlerInvoker == nil {
		logRemoveHandlerInvoker = gi.FunctionInvokerNew("GLib", "log_remove_handler")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetString(logDomain)
	inArgs[1].SetUint32(handlerId)

	logRemoveHandlerInvoker.Invoke(inArgs[:])
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

// UNSUPPORTED : C value 'g_main_context_default' : return type 'MainContext' not supported

// UNSUPPORTED : C value 'g_main_context_get_thread_default' : return type 'MainContext' not supported

// UNSUPPORTED : C value 'g_main_context_ref_thread_default' : return type 'MainContext' not supported

// UNSUPPORTED : C value 'g_main_current_source' : return type 'Source' not supported

var mainDepthInvoker *gi.Function

// MainDepth is a representation of the C type g_main_depth.
func MainDepth() int32 {
	if mainDepthInvoker == nil {
		mainDepthInvoker = gi.FunctionInvokerNew("GLib", "main_depth")
	}

	ret := mainDepthInvoker.Invoke(nil)
	return ret.Int32()
}

// UNSUPPORTED : C value 'g_malloc' : parameter 'n_bytes' of type 'gsize' not supported

// UNSUPPORTED : C value 'g_malloc0' : parameter 'n_bytes' of type 'gsize' not supported

// UNSUPPORTED : C value 'g_malloc0_n' : parameter 'n_blocks' of type 'gsize' not supported

// UNSUPPORTED : C value 'g_malloc_n' : parameter 'n_blocks' of type 'gsize' not supported

// UNSUPPORTED : C value 'g_markup_collect_attributes' : parameter 'error' of type 'Error' not supported

// UNSUPPORTED : C value 'g_markup_error_quark' : return type 'Quark' not supported

var markupEscapeTextInvoker *gi.Function

// MarkupEscapeText is a representation of the C type g_markup_escape_text.
func MarkupEscapeText(text string, length int32) string {
	if markupEscapeTextInvoker == nil {
		markupEscapeTextInvoker = gi.FunctionInvokerNew("GLib", "markup_escape_text")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetString(text)
	inArgs[1].SetInt32(length)

	ret := markupEscapeTextInvoker.Invoke(inArgs[:])
	return ret.String(true)
}

// UNSUPPORTED : C value 'g_markup_printf_escaped' : parameter '...' has no type

// UNSUPPORTED : C value 'g_markup_vprintf_escaped' : parameter 'args' of type 'va_list' not supported

// UNSUPPORTED : C value 'g_mem_is_system_malloc' : return type 'gboolean' not supported

var memProfileInvoker *gi.Function

// MemProfile is a representation of the C type g_mem_profile.
func MemProfile() {
	if memProfileInvoker == nil {
		memProfileInvoker = gi.FunctionInvokerNew("GLib", "mem_profile")
	}

	memProfileInvoker.Invoke(nil)
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

var onErrorQueryInvoker *gi.Function

// OnErrorQuery is a representation of the C type g_on_error_query.
func OnErrorQuery(prgName string) {
	if onErrorQueryInvoker == nil {
		onErrorQueryInvoker = gi.FunctionInvokerNew("GLib", "on_error_query")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetString(prgName)

	onErrorQueryInvoker.Invoke(inArgs[:])
}

var onErrorStackTraceInvoker *gi.Function

// OnErrorStackTrace is a representation of the C type g_on_error_stack_trace.
func OnErrorStackTrace(prgName string) {
	if onErrorStackTraceInvoker == nil {
		onErrorStackTraceInvoker = gi.FunctionInvokerNew("GLib", "on_error_stack_trace")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetString(prgName)

	onErrorStackTraceInvoker.Invoke(inArgs[:])
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

var randomIntInvoker *gi.Function

// RandomInt is a representation of the C type g_random_int.
func RandomInt() uint32 {
	if randomIntInvoker == nil {
		randomIntInvoker = gi.FunctionInvokerNew("GLib", "random_int")
	}

	ret := randomIntInvoker.Invoke(nil)
	return ret.Uint32()
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

	ret := randomIntRangeInvoker.Invoke(inArgs[:])
	return ret.Int32()
}

var randomSetSeedInvoker *gi.Function

// RandomSetSeed is a representation of the C type g_random_set_seed.
func RandomSetSeed(seed uint32) {
	if randomSetSeedInvoker == nil {
		randomSetSeedInvoker = gi.FunctionInvokerNew("GLib", "random_set_seed")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetUint32(seed)

	randomSetSeedInvoker.Invoke(inArgs[:])
}

// UNSUPPORTED : C value 'g_realloc' : parameter 'mem' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_realloc_n' : parameter 'mem' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_regex_check_replacement' : parameter 'has_references' of type 'gboolean' not supported

// UNSUPPORTED : C value 'g_regex_error_quark' : return type 'Quark' not supported

var regexEscapeNulInvoker *gi.Function

// RegexEscapeNul is a representation of the C type g_regex_escape_nul.
func RegexEscapeNul(string_ string, length int32) string {
	if regexEscapeNulInvoker == nil {
		regexEscapeNulInvoker = gi.FunctionInvokerNew("GLib", "regex_escape_nul")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetString(string_)
	inArgs[1].SetInt32(length)

	ret := regexEscapeNulInvoker.Invoke(inArgs[:])
	return ret.String(true)
}

// UNSUPPORTED : C value 'g_regex_escape_string' : parameter 'string' has no type

// UNSUPPORTED : C value 'g_regex_match_simple' : parameter 'compile_options' of type 'RegexCompileFlags' not supported

// UNSUPPORTED : C value 'g_regex_split_simple' : parameter 'compile_options' of type 'RegexCompileFlags' not supported

var reloadUserSpecialDirsCacheInvoker *gi.Function

// ReloadUserSpecialDirsCache is a representation of the C type g_reload_user_special_dirs_cache.
func ReloadUserSpecialDirsCache() {
	if reloadUserSpecialDirsCacheInvoker == nil {
		reloadUserSpecialDirsCacheInvoker = gi.FunctionInvokerNew("GLib", "reload_user_special_dirs_cache")
	}

	reloadUserSpecialDirsCacheInvoker.Invoke(nil)
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

	returnIfFailWarningInvoker.Invoke(inArgs[:])
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

var setApplicationNameInvoker *gi.Function

// SetApplicationName is a representation of the C type g_set_application_name.
func SetApplicationName(applicationName string) {
	if setApplicationNameInvoker == nil {
		setApplicationNameInvoker = gi.FunctionInvokerNew("GLib", "set_application_name")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetString(applicationName)

	setApplicationNameInvoker.Invoke(inArgs[:])
}

// UNSUPPORTED : C value 'g_set_error' : parameter 'err' of type 'Error' not supported

// UNSUPPORTED : C value 'g_set_error_literal' : parameter 'err' of type 'Error' not supported

var setPrgnameInvoker *gi.Function

// SetPrgname is a representation of the C type g_set_prgname.
func SetPrgname(prgname string) {
	if setPrgnameInvoker == nil {
		setPrgnameInvoker = gi.FunctionInvokerNew("GLib", "set_prgname")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetString(prgname)

	setPrgnameInvoker.Invoke(inArgs[:])
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

var sourceSetNameByIdInvoker *gi.Function

// SourceSetNameById is a representation of the C type g_source_set_name_by_id.
func SourceSetNameById(tag uint32, name string) {
	if sourceSetNameByIdInvoker == nil {
		sourceSetNameByIdInvoker = gi.FunctionInvokerNew("GLib", "source_set_name_by_id")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetUint32(tag)
	inArgs[1].SetString(name)

	sourceSetNameByIdInvoker.Invoke(inArgs[:])
}

var spacedPrimesClosestInvoker *gi.Function

// SpacedPrimesClosest is a representation of the C type g_spaced_primes_closest.
func SpacedPrimesClosest(num uint32) uint32 {
	if spacedPrimesClosestInvoker == nil {
		spacedPrimesClosestInvoker = gi.FunctionInvokerNew("GLib", "spaced_primes_closest")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetUint32(num)

	ret := spacedPrimesClosestInvoker.Invoke(inArgs[:])
	return ret.Uint32()
}

// UNSUPPORTED : C value 'g_spawn_async' : parameter 'working_directory' of type 'filename' not supported

// UNSUPPORTED : C value 'g_spawn_async_with_pipes' : parameter 'working_directory' of type 'filename' not supported

// UNSUPPORTED : C value 'g_spawn_check_exit_status' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_spawn_close_pid' : parameter 'pid' of type 'Pid' not supported

// UNSUPPORTED : C value 'g_spawn_command_line_async' : parameter 'command_line' of type 'filename' not supported

// UNSUPPORTED : C value 'g_spawn_command_line_sync' : parameter 'command_line' of type 'filename' not supported

// UNSUPPORTED : C value 'g_spawn_error_quark' : return type 'Quark' not supported

// UNSUPPORTED : C value 'g_spawn_exit_error_quark' : return type 'Quark' not supported

// UNSUPPORTED : C value 'g_spawn_sync' : parameter 'working_directory' of type 'filename' not supported

// UNSUPPORTED : C value 'g_sprintf' : parameter '...' has no type

var stpcpyInvoker *gi.Function

// Stpcpy is a representation of the C type g_stpcpy.
func Stpcpy(dest string, src string) string {
	if stpcpyInvoker == nil {
		stpcpyInvoker = gi.FunctionInvokerNew("GLib", "stpcpy")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetString(dest)
	inArgs[1].SetString(src)

	ret := stpcpyInvoker.Invoke(inArgs[:])
	return ret.String(true)
}

// UNSUPPORTED : C value 'g_str_equal' : parameter 'v1' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_str_has_prefix' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_str_has_suffix' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_str_hash' : parameter 'v' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_str_is_ascii' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_str_match_string' : parameter 'accept_alternates' of type 'gboolean' not supported

var strToAsciiInvoker *gi.Function

// StrToAscii is a representation of the C type g_str_to_ascii.
func StrToAscii(str string, fromLocale string) string {
	if strToAsciiInvoker == nil {
		strToAsciiInvoker = gi.FunctionInvokerNew("GLib", "str_to_ascii")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetString(str)
	inArgs[1].SetString(fromLocale)

	ret := strToAsciiInvoker.Invoke(inArgs[:])
	return ret.String(true)
}

// UNSUPPORTED : C value 'g_str_tokenize_and_fold' : parameter 'ascii_alternates' has no type

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

	ret := strcanonInvoker.Invoke(inArgs[:])
	return ret.String(true)
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

	ret := strcasecmpInvoker.Invoke(inArgs[:])
	return ret.Int32()
}

var strchompInvoker *gi.Function

// Strchomp is a representation of the C type g_strchomp.
func Strchomp(string_ string) string {
	if strchompInvoker == nil {
		strchompInvoker = gi.FunctionInvokerNew("GLib", "strchomp")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetString(string_)

	ret := strchompInvoker.Invoke(inArgs[:])
	return ret.String(true)
}

var strchugInvoker *gi.Function

// Strchug is a representation of the C type g_strchug.
func Strchug(string_ string) string {
	if strchugInvoker == nil {
		strchugInvoker = gi.FunctionInvokerNew("GLib", "strchug")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetString(string_)

	ret := strchugInvoker.Invoke(inArgs[:])
	return ret.String(true)
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

	ret := strcmp0Invoker.Invoke(inArgs[:])
	return ret.Int32()
}

var strcompressInvoker *gi.Function

// Strcompress is a representation of the C type g_strcompress.
func Strcompress(source string) string {
	if strcompressInvoker == nil {
		strcompressInvoker = gi.FunctionInvokerNew("GLib", "strcompress")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetString(source)

	ret := strcompressInvoker.Invoke(inArgs[:])
	return ret.String(true)
}

// UNSUPPORTED : C value 'g_strconcat' : parameter '...' has no type

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

	ret := strdelimitInvoker.Invoke(inArgs[:])
	return ret.String(true)
}

var strdownInvoker *gi.Function

// Strdown is a representation of the C type g_strdown.
func Strdown(string_ string) string {
	if strdownInvoker == nil {
		strdownInvoker = gi.FunctionInvokerNew("GLib", "strdown")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetString(string_)

	ret := strdownInvoker.Invoke(inArgs[:])
	return ret.String(true)
}

var strdupInvoker *gi.Function

// Strdup is a representation of the C type g_strdup.
func Strdup(str string) string {
	if strdupInvoker == nil {
		strdupInvoker = gi.FunctionInvokerNew("GLib", "strdup")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetString(str)

	ret := strdupInvoker.Invoke(inArgs[:])
	return ret.String(true)
}

// UNSUPPORTED : C value 'g_strdup_printf' : parameter '...' has no type

// UNSUPPORTED : C value 'g_strdup_vprintf' : parameter 'args' of type 'va_list' not supported

var strdupvInvoker *gi.Function

// Strdupv is a representation of the C type g_strdupv.
func Strdupv(strArray string) {
	if strdupvInvoker == nil {
		strdupvInvoker = gi.FunctionInvokerNew("GLib", "strdupv")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetString(strArray)

	strdupvInvoker.Invoke(inArgs[:])
}

var strerrorInvoker *gi.Function

// Strerror is a representation of the C type g_strerror.
func Strerror(errnum int32) string {
	if strerrorInvoker == nil {
		strerrorInvoker = gi.FunctionInvokerNew("GLib", "strerror")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetInt32(errnum)

	ret := strerrorInvoker.Invoke(inArgs[:])
	return ret.String(false)
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

	ret := strescapeInvoker.Invoke(inArgs[:])
	return ret.String(true)
}

var strfreevInvoker *gi.Function

// Strfreev is a representation of the C type g_strfreev.
func Strfreev(strArray string) {
	if strfreevInvoker == nil {
		strfreevInvoker = gi.FunctionInvokerNew("GLib", "strfreev")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetString(strArray)

	strfreevInvoker.Invoke(inArgs[:])
}

// UNSUPPORTED : C value 'g_string_new' : return type 'String' not supported

// UNSUPPORTED : C value 'g_string_new_len' : return type 'String' not supported

// UNSUPPORTED : C value 'g_string_sized_new' : parameter 'dfl_size' of type 'gsize' not supported

var stripContextInvoker *gi.Function

// StripContext is a representation of the C type g_strip_context.
func StripContext(msgid string, msgval string) string {
	if stripContextInvoker == nil {
		stripContextInvoker = gi.FunctionInvokerNew("GLib", "strip_context")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetString(msgid)
	inArgs[1].SetString(msgval)

	ret := stripContextInvoker.Invoke(inArgs[:])
	return ret.String(false)
}

// UNSUPPORTED : C value 'g_strjoin' : parameter '...' has no type

var strjoinvInvoker *gi.Function

// Strjoinv is a representation of the C type g_strjoinv.
func Strjoinv(separator string, strArray string) string {
	if strjoinvInvoker == nil {
		strjoinvInvoker = gi.FunctionInvokerNew("GLib", "strjoinv")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetString(separator)
	inArgs[1].SetString(strArray)

	ret := strjoinvInvoker.Invoke(inArgs[:])
	return ret.String(true)
}

// UNSUPPORTED : C value 'g_strlcat' : parameter 'dest_size' of type 'gsize' not supported

// UNSUPPORTED : C value 'g_strlcpy' : parameter 'dest_size' of type 'gsize' not supported

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

	ret := strncasecmpInvoker.Invoke(inArgs[:])
	return ret.Int32()
}

// UNSUPPORTED : C value 'g_strndup' : parameter 'n' of type 'gsize' not supported

// UNSUPPORTED : C value 'g_strnfill' : parameter 'length' of type 'gsize' not supported

var strreverseInvoker *gi.Function

// Strreverse is a representation of the C type g_strreverse.
func Strreverse(string_ string) string {
	if strreverseInvoker == nil {
		strreverseInvoker = gi.FunctionInvokerNew("GLib", "strreverse")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetString(string_)

	ret := strreverseInvoker.Invoke(inArgs[:])
	return ret.String(true)
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

	ret := strrstrInvoker.Invoke(inArgs[:])
	return ret.String(true)
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

	ret := strrstrLenInvoker.Invoke(inArgs[:])
	return ret.String(true)
}

var strsignalInvoker *gi.Function

// Strsignal is a representation of the C type g_strsignal.
func Strsignal(signum int32) string {
	if strsignalInvoker == nil {
		strsignalInvoker = gi.FunctionInvokerNew("GLib", "strsignal")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetInt32(signum)

	ret := strsignalInvoker.Invoke(inArgs[:])
	return ret.String(false)
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

	strsplitInvoker.Invoke(inArgs[:])
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

	strsplitSetInvoker.Invoke(inArgs[:])
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

	ret := strstrLenInvoker.Invoke(inArgs[:])
	return ret.String(true)
}

// UNSUPPORTED : C value 'g_strtod' : parameter 'endptr' with direction 'out' not supported

var strupInvoker *gi.Function

// Strup is a representation of the C type g_strup.
func Strup(string_ string) string {
	if strupInvoker == nil {
		strupInvoker = gi.FunctionInvokerNew("GLib", "strup")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetString(string_)

	ret := strupInvoker.Invoke(inArgs[:])
	return ret.String(true)
}

// UNSUPPORTED : C value 'g_strv_contains' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_strv_get_type' : return type 'GType' not supported

var strvLengthInvoker *gi.Function

// StrvLength is a representation of the C type g_strv_length.
func StrvLength(strArray string) uint32 {
	if strvLengthInvoker == nil {
		strvLengthInvoker = gi.FunctionInvokerNew("GLib", "strv_length")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetString(strArray)

	ret := strvLengthInvoker.Invoke(inArgs[:])
	return ret.Uint32()
}

// UNSUPPORTED : C value 'g_test_add_data_func' : parameter 'test_data' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_test_add_data_func_full' : parameter 'test_data' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_test_add_func' : parameter 'test_func' of type 'TestFunc' not supported

// UNSUPPORTED : C value 'g_test_add_vtable' : parameter 'data_size' of type 'gsize' not supported

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

	testAssertExpectedMessagesInternalInvoker.Invoke(inArgs[:])
}

var testBugInvoker *gi.Function

// TestBug is a representation of the C type g_test_bug.
func TestBug(bugUriSnippet string) {
	if testBugInvoker == nil {
		testBugInvoker = gi.FunctionInvokerNew("GLib", "test_bug")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetString(bugUriSnippet)

	testBugInvoker.Invoke(inArgs[:])
}

var testBugBaseInvoker *gi.Function

// TestBugBase is a representation of the C type g_test_bug_base.
func TestBugBase(uriPattern string) {
	if testBugBaseInvoker == nil {
		testBugBaseInvoker = gi.FunctionInvokerNew("GLib", "test_bug_base")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetString(uriPattern)

	testBugBaseInvoker.Invoke(inArgs[:])
}

// UNSUPPORTED : C value 'g_test_build_filename' : parameter 'file_type' of type 'TestFileType' not supported

// UNSUPPORTED : C value 'g_test_create_case' : parameter 'data_size' of type 'gsize' not supported

// UNSUPPORTED : C value 'g_test_create_suite' : return type 'TestSuite' not supported

// UNSUPPORTED : C value 'g_test_expect_message' : parameter 'log_level' of type 'LogLevelFlags' not supported

var testFailInvoker *gi.Function

// TestFail is a representation of the C type g_test_fail.
func TestFail() {
	if testFailInvoker == nil {
		testFailInvoker = gi.FunctionInvokerNew("GLib", "test_fail")
	}

	testFailInvoker.Invoke(nil)
}

// UNSUPPORTED : C value 'g_test_failed' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_test_get_dir' : parameter 'file_type' of type 'TestFileType' not supported

// UNSUPPORTED : C value 'g_test_get_filename' : parameter 'file_type' of type 'TestFileType' not supported

// UNSUPPORTED : C value 'g_test_get_root' : return type 'TestSuite' not supported

var testIncompleteInvoker *gi.Function

// TestIncomplete is a representation of the C type g_test_incomplete.
func TestIncomplete(msg string) {
	if testIncompleteInvoker == nil {
		testIncompleteInvoker = gi.FunctionInvokerNew("GLib", "test_incomplete")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetString(msg)

	testIncompleteInvoker.Invoke(inArgs[:])
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

var testRandIntInvoker *gi.Function

// TestRandInt is a representation of the C type g_test_rand_int.
func TestRandInt() int32 {
	if testRandIntInvoker == nil {
		testRandIntInvoker = gi.FunctionInvokerNew("GLib", "test_rand_int")
	}

	ret := testRandIntInvoker.Invoke(nil)
	return ret.Int32()
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

	ret := testRandIntRangeInvoker.Invoke(inArgs[:])
	return ret.Int32()
}

var testRunInvoker *gi.Function

// TestRun is a representation of the C type g_test_run.
func TestRun() int32 {
	if testRunInvoker == nil {
		testRunInvoker = gi.FunctionInvokerNew("GLib", "test_run")
	}

	ret := testRunInvoker.Invoke(nil)
	return ret.Int32()
}

// UNSUPPORTED : C value 'g_test_run_suite' : parameter 'suite' of type 'TestSuite' not supported

var testSetNonfatalAssertionsInvoker *gi.Function

// TestSetNonfatalAssertions is a representation of the C type g_test_set_nonfatal_assertions.
func TestSetNonfatalAssertions() {
	if testSetNonfatalAssertionsInvoker == nil {
		testSetNonfatalAssertionsInvoker = gi.FunctionInvokerNew("GLib", "test_set_nonfatal_assertions")
	}

	testSetNonfatalAssertionsInvoker.Invoke(nil)
}

var testSkipInvoker *gi.Function

// TestSkip is a representation of the C type g_test_skip.
func TestSkip(msg string) {
	if testSkipInvoker == nil {
		testSkipInvoker = gi.FunctionInvokerNew("GLib", "test_skip")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetString(msg)

	testSkipInvoker.Invoke(inArgs[:])
}

// UNSUPPORTED : C value 'g_test_subprocess' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_test_timer_elapsed' : return type 'gdouble' not supported

// UNSUPPORTED : C value 'g_test_timer_last' : return type 'gdouble' not supported

var testTimerStartInvoker *gi.Function

// TestTimerStart is a representation of the C type g_test_timer_start.
func TestTimerStart() {
	if testTimerStartInvoker == nil {
		testTimerStartInvoker = gi.FunctionInvokerNew("GLib", "test_timer_start")
	}

	testTimerStartInvoker.Invoke(nil)
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

	testTrapAssertionsInvoker.Invoke(inArgs[:])
}

// UNSUPPORTED : C value 'g_test_trap_fork' : parameter 'test_trap_flags' of type 'TestTrapFlags' not supported

// UNSUPPORTED : C value 'g_test_trap_has_passed' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_test_trap_reached_timeout' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_test_trap_subprocess' : parameter 'test_flags' of type 'TestSubprocessFlags' not supported

// UNSUPPORTED : C value 'g_thread_error_quark' : return type 'Quark' not supported

// UNSUPPORTED : C value 'g_thread_exit' : parameter 'retval' of type 'gpointer' not supported

var threadPoolGetMaxIdleTimeInvoker *gi.Function

// ThreadPoolGetMaxIdleTime is a representation of the C type g_thread_pool_get_max_idle_time.
func ThreadPoolGetMaxIdleTime() uint32 {
	if threadPoolGetMaxIdleTimeInvoker == nil {
		threadPoolGetMaxIdleTimeInvoker = gi.FunctionInvokerNew("GLib", "thread_pool_get_max_idle_time")
	}

	ret := threadPoolGetMaxIdleTimeInvoker.Invoke(nil)
	return ret.Uint32()
}

var threadPoolGetMaxUnusedThreadsInvoker *gi.Function

// ThreadPoolGetMaxUnusedThreads is a representation of the C type g_thread_pool_get_max_unused_threads.
func ThreadPoolGetMaxUnusedThreads() int32 {
	if threadPoolGetMaxUnusedThreadsInvoker == nil {
		threadPoolGetMaxUnusedThreadsInvoker = gi.FunctionInvokerNew("GLib", "thread_pool_get_max_unused_threads")
	}

	ret := threadPoolGetMaxUnusedThreadsInvoker.Invoke(nil)
	return ret.Int32()
}

var threadPoolGetNumUnusedThreadsInvoker *gi.Function

// ThreadPoolGetNumUnusedThreads is a representation of the C type g_thread_pool_get_num_unused_threads.
func ThreadPoolGetNumUnusedThreads() uint32 {
	if threadPoolGetNumUnusedThreadsInvoker == nil {
		threadPoolGetNumUnusedThreadsInvoker = gi.FunctionInvokerNew("GLib", "thread_pool_get_num_unused_threads")
	}

	ret := threadPoolGetNumUnusedThreadsInvoker.Invoke(nil)
	return ret.Uint32()
}

var threadPoolSetMaxIdleTimeInvoker *gi.Function

// ThreadPoolSetMaxIdleTime is a representation of the C type g_thread_pool_set_max_idle_time.
func ThreadPoolSetMaxIdleTime(interval uint32) {
	if threadPoolSetMaxIdleTimeInvoker == nil {
		threadPoolSetMaxIdleTimeInvoker = gi.FunctionInvokerNew("GLib", "thread_pool_set_max_idle_time")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetUint32(interval)

	threadPoolSetMaxIdleTimeInvoker.Invoke(inArgs[:])
}

var threadPoolSetMaxUnusedThreadsInvoker *gi.Function

// ThreadPoolSetMaxUnusedThreads is a representation of the C type g_thread_pool_set_max_unused_threads.
func ThreadPoolSetMaxUnusedThreads(maxThreads int32) {
	if threadPoolSetMaxUnusedThreadsInvoker == nil {
		threadPoolSetMaxUnusedThreadsInvoker = gi.FunctionInvokerNew("GLib", "thread_pool_set_max_unused_threads")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetInt32(maxThreads)

	threadPoolSetMaxUnusedThreadsInvoker.Invoke(inArgs[:])
}

var threadPoolStopUnusedThreadsInvoker *gi.Function

// ThreadPoolStopUnusedThreads is a representation of the C type g_thread_pool_stop_unused_threads.
func ThreadPoolStopUnusedThreads() {
	if threadPoolStopUnusedThreadsInvoker == nil {
		threadPoolStopUnusedThreadsInvoker = gi.FunctionInvokerNew("GLib", "thread_pool_stop_unused_threads")
	}

	threadPoolStopUnusedThreadsInvoker.Invoke(nil)
}

// UNSUPPORTED : C value 'g_thread_self' : return type 'Thread' not supported

var threadYieldInvoker *gi.Function

// ThreadYield is a representation of the C type g_thread_yield.
func ThreadYield() {
	if threadYieldInvoker == nil {
		threadYieldInvoker = gi.FunctionInvokerNew("GLib", "thread_yield")
	}

	threadYieldInvoker.Invoke(nil)
}

// UNSUPPORTED : C value 'g_time_val_from_iso8601' : parameter 'time_' of type 'TimeVal' not supported

// UNSUPPORTED : C value 'g_timeout_add' : parameter 'function' of type 'SourceFunc' not supported

// UNSUPPORTED : C value 'g_timeout_add_full' : parameter 'function' of type 'SourceFunc' not supported

// UNSUPPORTED : C value 'g_timeout_add_seconds' : parameter 'function' of type 'SourceFunc' not supported

// UNSUPPORTED : C value 'g_timeout_add_seconds_full' : parameter 'function' of type 'SourceFunc' not supported

// UNSUPPORTED : C value 'g_timeout_source_new' : return type 'Source' not supported

// UNSUPPORTED : C value 'g_timeout_source_new_seconds' : return type 'Source' not supported

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

// UNSUPPORTED : C value 'g_unix_signal_source_new' : return type 'Source' not supported

// UNSUPPORTED : C value 'g_unlink' : parameter 'filename' of type 'filename' not supported

// UNSUPPORTED : C value 'g_unsetenv' : parameter 'variable' of type 'filename' not supported

// UNSUPPORTED : C value 'g_uri_escape_string' : parameter 'allow_utf8' of type 'gboolean' not supported

var uriListExtractUrisInvoker *gi.Function

// UriListExtractUris is a representation of the C type g_uri_list_extract_uris.
func UriListExtractUris(uriList string) {
	if uriListExtractUrisInvoker == nil {
		uriListExtractUrisInvoker = gi.FunctionInvokerNew("GLib", "uri_list_extract_uris")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetString(uriList)

	uriListExtractUrisInvoker.Invoke(inArgs[:])
}

var uriParseSchemeInvoker *gi.Function

// UriParseScheme is a representation of the C type g_uri_parse_scheme.
func UriParseScheme(uri string) string {
	if uriParseSchemeInvoker == nil {
		uriParseSchemeInvoker = gi.FunctionInvokerNew("GLib", "uri_parse_scheme")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetString(uri)

	ret := uriParseSchemeInvoker.Invoke(inArgs[:])
	return ret.String(true)
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

	ret := uriUnescapeSegmentInvoker.Invoke(inArgs[:])
	return ret.String(true)
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

	ret := uriUnescapeStringInvoker.Invoke(inArgs[:])
	return ret.String(true)
}

var usleepInvoker *gi.Function

// Usleep is a representation of the C type g_usleep.
func Usleep(microseconds uint64) {
	if usleepInvoker == nil {
		usleepInvoker = gi.FunctionInvokerNew("GLib", "usleep")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetUint64(microseconds)

	usleepInvoker.Invoke(inArgs[:])
}

// UNSUPPORTED : C value 'g_utf16_to_ucs4' : parameter 'items_read' with direction 'out' not supported

// UNSUPPORTED : C value 'g_utf16_to_utf8' : parameter 'items_read' with direction 'out' not supported

var utf8CasefoldInvoker *gi.Function

// Utf8Casefold is a representation of the C type g_utf8_casefold.
func Utf8Casefold(str string, len int32) string {
	if utf8CasefoldInvoker == nil {
		utf8CasefoldInvoker = gi.FunctionInvokerNew("GLib", "utf8_casefold")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetString(str)
	inArgs[1].SetInt32(len)

	ret := utf8CasefoldInvoker.Invoke(inArgs[:])
	return ret.String(true)
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

	ret := utf8CollateInvoker.Invoke(inArgs[:])
	return ret.Int32()
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

	ret := utf8CollateKeyInvoker.Invoke(inArgs[:])
	return ret.String(true)
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

	ret := utf8CollateKeyForFilenameInvoker.Invoke(inArgs[:])
	return ret.String(true)
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

	ret := utf8FindNextCharInvoker.Invoke(inArgs[:])
	return ret.String(true)
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

	ret := utf8FindPrevCharInvoker.Invoke(inArgs[:])
	return ret.String(true)
}

// UNSUPPORTED : C value 'g_utf8_get_char' : return type 'gunichar' not supported

// UNSUPPORTED : C value 'g_utf8_get_char_validated' : return type 'gunichar' not supported

var utf8MakeValidInvoker *gi.Function

// Utf8MakeValid is a representation of the C type g_utf8_make_valid.
func Utf8MakeValid(str string, len int32) string {
	if utf8MakeValidInvoker == nil {
		utf8MakeValidInvoker = gi.FunctionInvokerNew("GLib", "utf8_make_valid")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetString(str)
	inArgs[1].SetInt32(len)

	ret := utf8MakeValidInvoker.Invoke(inArgs[:])
	return ret.String(true)
}

// UNSUPPORTED : C value 'g_utf8_normalize' : parameter 'mode' of type 'NormalizeMode' not supported

var utf8OffsetToPointerInvoker *gi.Function

// Utf8OffsetToPointer is a representation of the C type g_utf8_offset_to_pointer.
func Utf8OffsetToPointer(str string, offset int64) string {
	if utf8OffsetToPointerInvoker == nil {
		utf8OffsetToPointerInvoker = gi.FunctionInvokerNew("GLib", "utf8_offset_to_pointer")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetString(str)
	inArgs[1].SetInt64(offset)

	ret := utf8OffsetToPointerInvoker.Invoke(inArgs[:])
	return ret.String(true)
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

	ret := utf8PointerToOffsetInvoker.Invoke(inArgs[:])
	return ret.Int64()
}

var utf8PrevCharInvoker *gi.Function

// Utf8PrevChar is a representation of the C type g_utf8_prev_char.
func Utf8PrevChar(p string) string {
	if utf8PrevCharInvoker == nil {
		utf8PrevCharInvoker = gi.FunctionInvokerNew("GLib", "utf8_prev_char")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetString(p)

	ret := utf8PrevCharInvoker.Invoke(inArgs[:])
	return ret.String(true)
}

// UNSUPPORTED : C value 'g_utf8_strchr' : parameter 'c' of type 'gunichar' not supported

var utf8StrdownInvoker *gi.Function

// Utf8Strdown is a representation of the C type g_utf8_strdown.
func Utf8Strdown(str string, len int32) string {
	if utf8StrdownInvoker == nil {
		utf8StrdownInvoker = gi.FunctionInvokerNew("GLib", "utf8_strdown")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetString(str)
	inArgs[1].SetInt32(len)

	ret := utf8StrdownInvoker.Invoke(inArgs[:])
	return ret.String(true)
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

	ret := utf8StrlenInvoker.Invoke(inArgs[:])
	return ret.Int64()
}

// UNSUPPORTED : C value 'g_utf8_strncpy' : parameter 'n' of type 'gsize' not supported

// UNSUPPORTED : C value 'g_utf8_strrchr' : parameter 'c' of type 'gunichar' not supported

var utf8StrreverseInvoker *gi.Function

// Utf8Strreverse is a representation of the C type g_utf8_strreverse.
func Utf8Strreverse(str string, len int32) string {
	if utf8StrreverseInvoker == nil {
		utf8StrreverseInvoker = gi.FunctionInvokerNew("GLib", "utf8_strreverse")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetString(str)
	inArgs[1].SetInt32(len)

	ret := utf8StrreverseInvoker.Invoke(inArgs[:])
	return ret.String(true)
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

	ret := utf8StrupInvoker.Invoke(inArgs[:])
	return ret.String(true)
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

	ret := utf8SubstringInvoker.Invoke(inArgs[:])
	return ret.String(true)
}

// UNSUPPORTED : C value 'g_utf8_to_ucs4' : parameter 'items_read' with direction 'out' not supported

// UNSUPPORTED : C value 'g_utf8_to_ucs4_fast' : parameter 'items_written' with direction 'out' not supported

// UNSUPPORTED : C value 'g_utf8_to_utf16' : parameter 'items_read' with direction 'out' not supported

// UNSUPPORTED : C value 'g_utf8_validate' : parameter 'str' has no type

// UNSUPPORTED : C value 'g_uuid_string_is_valid' : return type 'gboolean' not supported

var uuidStringRandomInvoker *gi.Function

// UuidStringRandom is a representation of the C type g_uuid_string_random.
func UuidStringRandom() string {
	if uuidStringRandomInvoker == nil {
		uuidStringRandomInvoker = gi.FunctionInvokerNew("GLib", "uuid_string_random")
	}

	ret := uuidStringRandomInvoker.Invoke(nil)
	return ret.String(true)
}

// UNSUPPORTED : C value 'g_variant_get_gtype' : return type 'GType' not supported

// UNSUPPORTED : C value 'g_variant_is_object_path' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_variant_is_signature' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_variant_parse' : parameter 'type' of type 'VariantType' not supported

// UNSUPPORTED : C value 'g_variant_parse_error_print_context' : parameter 'error' of type 'Error' not supported

// UNSUPPORTED : C value 'g_variant_parse_error_quark' : return type 'Quark' not supported

// UNSUPPORTED : C value 'g_variant_parser_get_error_quark' : return type 'Quark' not supported

// UNSUPPORTED : C value 'g_variant_type_checked_' : return type 'VariantType' not supported

// UNSUPPORTED : C value 'g_variant_type_string_is_valid' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_variant_type_string_scan' : parameter 'endptr' with direction 'out' not supported

// UNSUPPORTED : C value 'g_vasprintf' : parameter 'args' of type 'va_list' not supported

// UNSUPPORTED : C value 'g_vfprintf' : parameter 'file' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_vprintf' : parameter 'args' of type 'va_list' not supported

// UNSUPPORTED : C value 'g_vsnprintf' : parameter 'args' of type 'va_list' not supported

// UNSUPPORTED : C value 'g_vsprintf' : parameter 'args' of type 'va_list' not supported

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

	warnMessageInvoker.Invoke(inArgs[:])
}
