// This is a generated file - DO NOT EDIT

package glib

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

// AsciiDigitValue is a wrapper around the C function g_ascii_digit_value.
func AsciiDigitValue(c rune) int32 {
	c_c := (C.gchar)(c)

	retC := C.g_ascii_digit_value(c_c)
	retGo := (int32)(retC)

	return retGo
}

// AsciiDtostr is a wrapper around the C function g_ascii_dtostr.
func AsciiDtostr(buffer string, bufLen int32, d float64) string {
	c_buffer := C.CString(buffer)
	defer C.free(unsafe.Pointer(c_buffer))

	c_buf_len := (C.gint)(bufLen)

	c_d := (C.gdouble)(d)

	retC := C.g_ascii_dtostr(c_buffer, c_buf_len, c_d)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// AsciiFormatd is a wrapper around the C function g_ascii_formatd.
func AsciiFormatd(buffer string, bufLen int32, format string, d float64) string {
	c_buffer := C.CString(buffer)
	defer C.free(unsafe.Pointer(c_buffer))

	c_buf_len := (C.gint)(bufLen)

	c_format := C.CString(format)
	defer C.free(unsafe.Pointer(c_format))

	c_d := (C.gdouble)(d)

	retC := C.g_ascii_formatd(c_buffer, c_buf_len, c_format, c_d)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// AsciiStrcasecmp is a wrapper around the C function g_ascii_strcasecmp.
func AsciiStrcasecmp(s1 string, s2 string) int32 {
	c_s1 := C.CString(s1)
	defer C.free(unsafe.Pointer(c_s1))

	c_s2 := C.CString(s2)
	defer C.free(unsafe.Pointer(c_s2))

	retC := C.g_ascii_strcasecmp(c_s1, c_s2)
	retGo := (int32)(retC)

	return retGo
}

// AsciiStrdown is a wrapper around the C function g_ascii_strdown.
func AsciiStrdown(str string, len int64) string {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	c_len := (C.gssize)(len)

	retC := C.g_ascii_strdown(c_str, c_len)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// AsciiStrncasecmp is a wrapper around the C function g_ascii_strncasecmp.
func AsciiStrncasecmp(s1 string, s2 string, n uint64) int32 {
	c_s1 := C.CString(s1)
	defer C.free(unsafe.Pointer(c_s1))

	c_s2 := C.CString(s2)
	defer C.free(unsafe.Pointer(c_s2))

	c_n := (C.gsize)(n)

	retC := C.g_ascii_strncasecmp(c_s1, c_s2, c_n)
	retGo := (int32)(retC)

	return retGo
}

// AsciiStrtod is a wrapper around the C function g_ascii_strtod.
func AsciiStrtod(nptr string) (float64, string) {
	c_nptr := C.CString(nptr)
	defer C.free(unsafe.Pointer(c_nptr))

	var c_endptr *C.gchar

	retC := C.g_ascii_strtod(c_nptr, &c_endptr)
	retGo := (float64)(retC)

	endptr := C.GoString(c_endptr)

	return retGo, endptr
}

// AsciiStrup is a wrapper around the C function g_ascii_strup.
func AsciiStrup(str string, len int64) string {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	c_len := (C.gssize)(len)

	retC := C.g_ascii_strup(c_str, c_len)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// AsciiTolower is a wrapper around the C function g_ascii_tolower.
func AsciiTolower(c rune) rune {
	c_c := (C.gchar)(c)

	retC := C.g_ascii_tolower(c_c)
	retGo := (rune)(retC)

	return retGo
}

// AsciiToupper is a wrapper around the C function g_ascii_toupper.
func AsciiToupper(c rune) rune {
	c_c := (C.gchar)(c)

	retC := C.g_ascii_toupper(c_c)
	retGo := (rune)(retC)

	return retGo
}

// AsciiXdigitValue is a wrapper around the C function g_ascii_xdigit_value.
func AsciiXdigitValue(c rune) int32 {
	c_c := (C.gchar)(c)

	retC := C.g_ascii_xdigit_value(c_c)
	retGo := (int32)(retC)

	return retGo
}

// AssertWarning is a wrapper around the C function g_assert_warning.
func AssertWarning(logDomain string, file string, line int32, prettyFunction string, expression string) {
	c_log_domain := C.CString(logDomain)
	defer C.free(unsafe.Pointer(c_log_domain))

	c_file := C.CString(file)
	defer C.free(unsafe.Pointer(c_file))

	c_line := (C.int)(line)

	c_pretty_function := C.CString(prettyFunction)
	defer C.free(unsafe.Pointer(c_pretty_function))

	c_expression := C.CString(expression)
	defer C.free(unsafe.Pointer(c_expression))

	C.g_assert_warning(c_log_domain, c_file, c_line, c_pretty_function, c_expression)

	return
}

// AssertionMessage is a wrapper around the C function g_assertion_message.
func AssertionMessage(domain string, file string, line int32, func_ string, message string) {
	c_domain := C.CString(domain)
	defer C.free(unsafe.Pointer(c_domain))

	c_file := C.CString(file)
	defer C.free(unsafe.Pointer(c_file))

	c_line := (C.int)(line)

	c_func := C.CString(func_)
	defer C.free(unsafe.Pointer(c_func))

	c_message := C.CString(message)
	defer C.free(unsafe.Pointer(c_message))

	C.g_assertion_message(c_domain, c_file, c_line, c_func, c_message)

	return
}

// Unsupported : g_assertion_message_cmpnum : unsupported parameter numtype : no type generator for gchar (char) for param numtype

// AssertionMessageCmpstr is a wrapper around the C function g_assertion_message_cmpstr.
func AssertionMessageCmpstr(domain string, file string, line int32, func_ string, expr string, arg1 string, cmp string, arg2 string) {
	c_domain := C.CString(domain)
	defer C.free(unsafe.Pointer(c_domain))

	c_file := C.CString(file)
	defer C.free(unsafe.Pointer(c_file))

	c_line := (C.int)(line)

	c_func := C.CString(func_)
	defer C.free(unsafe.Pointer(c_func))

	c_expr := C.CString(expr)
	defer C.free(unsafe.Pointer(c_expr))

	c_arg1 := C.CString(arg1)
	defer C.free(unsafe.Pointer(c_arg1))

	c_cmp := C.CString(cmp)
	defer C.free(unsafe.Pointer(c_cmp))

	c_arg2 := C.CString(arg2)
	defer C.free(unsafe.Pointer(c_arg2))

	C.g_assertion_message_cmpstr(c_domain, c_file, c_line, c_func, c_expr, c_arg1, c_cmp, c_arg2)

	return
}

// AssertionMessageError is a wrapper around the C function g_assertion_message_error.
func AssertionMessageError(domain string, file string, line int32, func_ string, expr string, error *Error, errorDomain Quark, errorCode int32) {
	c_domain := C.CString(domain)
	defer C.free(unsafe.Pointer(c_domain))

	c_file := C.CString(file)
	defer C.free(unsafe.Pointer(c_file))

	c_line := (C.int)(line)

	c_func := C.CString(func_)
	defer C.free(unsafe.Pointer(c_func))

	c_expr := C.CString(expr)
	defer C.free(unsafe.Pointer(c_expr))

	c_error := (*C.GError)(error.ToC())

	c_error_domain := (C.GQuark)(errorDomain)

	c_error_code := (C.int)(errorCode)

	C.g_assertion_message_error(c_domain, c_file, c_line, c_func, c_expr, c_error, c_error_domain, c_error_code)

	return
}

// AssertionMessageExpr is a wrapper around the C function g_assertion_message_expr.
func AssertionMessageExpr(domain string, file string, line int32, func_ string, expr string) {
	c_domain := C.CString(domain)
	defer C.free(unsafe.Pointer(c_domain))

	c_file := C.CString(file)
	defer C.free(unsafe.Pointer(c_file))

	c_line := (C.int)(line)

	c_func := C.CString(func_)
	defer C.free(unsafe.Pointer(c_func))

	c_expr := C.CString(expr)
	defer C.free(unsafe.Pointer(c_expr))

	C.g_assertion_message_expr(c_domain, c_file, c_line, c_func, c_expr)

	return
}

// Unsupported : g_atexit : unsupported parameter func : no type generator for VoidFunc (GVoidFunc) for param func

// Unsupported : g_atomic_pointer_add : unsupported parameter atomic : no type generator for gpointer (void*) for param atomic

// Unsupported : g_atomic_pointer_and : unsupported parameter atomic : no type generator for gpointer (void*) for param atomic

// Unsupported : g_atomic_pointer_compare_and_exchange : unsupported parameter atomic : no type generator for gpointer (void*) for param atomic

// Unsupported : g_atomic_pointer_get : unsupported parameter atomic : no type generator for gpointer (void*) for param atomic

// Unsupported : g_atomic_pointer_or : unsupported parameter atomic : no type generator for gpointer (void*) for param atomic

// Unsupported : g_atomic_pointer_set : unsupported parameter atomic : no type generator for gpointer (void*) for param atomic

// Unsupported : g_atomic_pointer_xor : unsupported parameter atomic : no type generator for gpointer (void*) for param atomic

// Unsupported : g_base64_decode : no return type

// Unsupported : g_base64_decode_inplace : unsupported parameter text : no type generator for guint8 () for array param text

// Unsupported : g_base64_decode_step : unsupported parameter in : no type generator for guint8 () for array param in

// Unsupported : g_base64_encode : unsupported parameter data : no type generator for guint8 () for array param data

// Unsupported : g_base64_encode_close : unsupported parameter out : no type generator for guint8 () for array param out

// Unsupported : g_base64_encode_step : unsupported parameter in : no type generator for guint8 () for array param in

// Basename is a wrapper around the C function g_basename.
func Basename(fileName string) string {
	c_file_name := C.CString(fileName)
	defer C.free(unsafe.Pointer(c_file_name))

	retC := C.g_basename(c_file_name)
	retGo := C.GoString(retC)

	return retGo
}

// BitNthLsf is a wrapper around the C function g_bit_nth_lsf.
func BitNthLsf(mask uint64, nthBit int32) int32 {
	c_mask := (C.gulong)(mask)

	c_nth_bit := (C.gint)(nthBit)

	retC := C.g_bit_nth_lsf(c_mask, c_nth_bit)
	retGo := (int32)(retC)

	return retGo
}

// BitNthMsf is a wrapper around the C function g_bit_nth_msf.
func BitNthMsf(mask uint64, nthBit int32) int32 {
	c_mask := (C.gulong)(mask)

	c_nth_bit := (C.gint)(nthBit)

	retC := C.g_bit_nth_msf(c_mask, c_nth_bit)
	retGo := (int32)(retC)

	return retGo
}

// BitStorage is a wrapper around the C function g_bit_storage.
func BitStorage(number uint64) uint32 {
	c_number := (C.gulong)(number)

	retC := C.g_bit_storage(c_number)
	retGo := (uint32)(retC)

	return retGo
}

// BookmarkFileErrorQuark is a wrapper around the C function g_bookmark_file_error_quark.
func BookmarkFileErrorQuark() Quark {
	retC := C.g_bookmark_file_error_quark()
	retGo := (Quark)(retC)

	return retGo
}

// Unsupported : g_build_filename : unsupported parameter ... : varargs

// Unsupported : g_build_filename_valist : unsupported parameter args : no type generator for va_list (va_list*) for param args

// Unsupported : g_build_filenamev : unsupported parameter args : no type generator for filename () for array param args

// Unsupported : g_build_path : unsupported parameter ... : varargs

// Unsupported : g_build_pathv : unsupported parameter args : no type generator for filename () for array param args

// Unsupported : g_byte_array_free : unsupported parameter array : no type generator for guint8 (guint8) for array param array

// Unsupported : g_byte_array_free_to_bytes : unsupported parameter array : no type generator for guint8 (guint8) for array param array

// Unsupported : g_byte_array_new : no return type

// Unsupported : g_byte_array_new_take : unsupported parameter data : no type generator for guint8 (guint8) for array param data

// Unsupported : g_byte_array_unref : unsupported parameter array : no type generator for guint8 (guint8) for array param array

// Unsupported : g_child_watch_add : unsupported parameter function : no type generator for ChildWatchFunc (GChildWatchFunc) for param function

// Unsupported : g_child_watch_add_full : unsupported parameter function : no type generator for ChildWatchFunc (GChildWatchFunc) for param function

// ClearError is a wrapper around the C function g_clear_error.
func ClearError() error {
	var cThrowableError *C.GError

	C.g_clear_error(&cThrowableError)

	goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return goThrowableError
}

// Unsupported : g_clear_handle_id : unsupported parameter clear_func : no type generator for ClearHandleFunc (GClearHandleFunc) for param clear_func

// Unsupported : g_clear_pointer : unsupported parameter destroy : no type generator for DestroyNotify (GDestroyNotify) for param destroy

// Unsupported : g_compute_checksum_for_data : unsupported parameter data : no type generator for guint8 () for array param data

// Unsupported : g_compute_hmac_for_data : unsupported parameter key : no type generator for guint8 (guchar) for array param key

// Unsupported : g_compute_hmac_for_string : unsupported parameter key : no type generator for guint8 (guchar) for array param key

// Unsupported : g_convert : unsupported parameter str : no type generator for guint8 () for array param str

// ConvertErrorQuark is a wrapper around the C function g_convert_error_quark.
func ConvertErrorQuark() Quark {
	retC := C.g_convert_error_quark()
	retGo := (Quark)(retC)

	return retGo
}

// Unsupported : g_convert_with_fallback : unsupported parameter str : no type generator for guint8 () for array param str

// Unsupported : g_convert_with_iconv : unsupported parameter str : no type generator for guint8 () for array param str

// Unsupported : g_datalist_clear : unsupported parameter datalist : record with indirection level of 2

// Unsupported : g_datalist_foreach : unsupported parameter datalist : record with indirection level of 2

// Unsupported : g_datalist_get_data : unsupported parameter datalist : record with indirection level of 2

// Unsupported : g_datalist_get_flags : unsupported parameter datalist : record with indirection level of 2

// Unsupported : g_datalist_id_dup_data : unsupported parameter datalist : record with indirection level of 2

// Unsupported : g_datalist_id_get_data : unsupported parameter datalist : record with indirection level of 2

// Unsupported : g_datalist_id_remove_no_notify : unsupported parameter datalist : record with indirection level of 2

// Unsupported : g_datalist_id_replace_data : unsupported parameter datalist : record with indirection level of 2

// Unsupported : g_datalist_id_set_data_full : unsupported parameter datalist : record with indirection level of 2

// Unsupported : g_datalist_init : unsupported parameter datalist : record with indirection level of 2

// Unsupported : g_datalist_set_flags : unsupported parameter datalist : record with indirection level of 2

// Unsupported : g_datalist_unset_flags : unsupported parameter datalist : record with indirection level of 2

// DatasetDestroy is a wrapper around the C function g_dataset_destroy.
func DatasetDestroy(datasetLocation uintptr) {
	c_dataset_location := (C.gconstpointer)(datasetLocation)

	C.g_dataset_destroy(c_dataset_location)

	return
}

// Unsupported : g_dataset_foreach : unsupported parameter func : no type generator for DataForeachFunc (GDataForeachFunc) for param func

// DatasetIdGetData is a wrapper around the C function g_dataset_id_get_data.
func DatasetIdGetData(datasetLocation uintptr, keyId Quark) uintptr {
	c_dataset_location := (C.gconstpointer)(datasetLocation)

	c_key_id := (C.GQuark)(keyId)

	retC := C.g_dataset_id_get_data(c_dataset_location, c_key_id)
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// DatasetIdRemoveNoNotify is a wrapper around the C function g_dataset_id_remove_no_notify.
func DatasetIdRemoveNoNotify(datasetLocation uintptr, keyId Quark) uintptr {
	c_dataset_location := (C.gconstpointer)(datasetLocation)

	c_key_id := (C.GQuark)(keyId)

	retC := C.g_dataset_id_remove_no_notify(c_dataset_location, c_key_id)
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_dataset_id_set_data_full : unsupported parameter destroy_func : no type generator for DestroyNotify (GDestroyNotify) for param destroy_func

// DateGetDaysInMonth is a wrapper around the C function g_date_get_days_in_month.
func DateGetDaysInMonth(month DateMonth, year DateYear) uint8 {
	c_month := (C.GDateMonth)(month)

	c_year := (C.GDateYear)(year)

	retC := C.g_date_get_days_in_month(c_month, c_year)
	retGo := (uint8)(retC)

	return retGo
}

// DateGetMondayWeeksInYear is a wrapper around the C function g_date_get_monday_weeks_in_year.
func DateGetMondayWeeksInYear(year DateYear) uint8 {
	c_year := (C.GDateYear)(year)

	retC := C.g_date_get_monday_weeks_in_year(c_year)
	retGo := (uint8)(retC)

	return retGo
}

// DateGetSundayWeeksInYear is a wrapper around the C function g_date_get_sunday_weeks_in_year.
func DateGetSundayWeeksInYear(year DateYear) uint8 {
	c_year := (C.GDateYear)(year)

	retC := C.g_date_get_sunday_weeks_in_year(c_year)
	retGo := (uint8)(retC)

	return retGo
}

// DateIsLeapYear is a wrapper around the C function g_date_is_leap_year.
func DateIsLeapYear(year DateYear) bool {
	c_year := (C.GDateYear)(year)

	retC := C.g_date_is_leap_year(c_year)
	retGo := retC == C.TRUE

	return retGo
}

// DateStrftime is a wrapper around the C function g_date_strftime.
func DateStrftime(s string, slen uint64, format string, date *Date) uint64 {
	c_s := C.CString(s)
	defer C.free(unsafe.Pointer(c_s))

	c_slen := (C.gsize)(slen)

	c_format := C.CString(format)
	defer C.free(unsafe.Pointer(c_format))

	c_date := (*C.GDate)(date.ToC())

	retC := C.g_date_strftime(c_s, c_slen, c_format, c_date)
	retGo := (uint64)(retC)

	return retGo
}

// DateValidDay is a wrapper around the C function g_date_valid_day.
func DateValidDay(day DateDay) bool {
	c_day := (C.GDateDay)(day)

	retC := C.g_date_valid_day(c_day)
	retGo := retC == C.TRUE

	return retGo
}

// DateValidDmy is a wrapper around the C function g_date_valid_dmy.
func DateValidDmy(day DateDay, month DateMonth, year DateYear) bool {
	c_day := (C.GDateDay)(day)

	c_month := (C.GDateMonth)(month)

	c_year := (C.GDateYear)(year)

	retC := C.g_date_valid_dmy(c_day, c_month, c_year)
	retGo := retC == C.TRUE

	return retGo
}

// DateValidJulian is a wrapper around the C function g_date_valid_julian.
func DateValidJulian(julianDate uint32) bool {
	c_julian_date := (C.guint32)(julianDate)

	retC := C.g_date_valid_julian(c_julian_date)
	retGo := retC == C.TRUE

	return retGo
}

// DateValidMonth is a wrapper around the C function g_date_valid_month.
func DateValidMonth(month DateMonth) bool {
	c_month := (C.GDateMonth)(month)

	retC := C.g_date_valid_month(c_month)
	retGo := retC == C.TRUE

	return retGo
}

// DateValidWeekday is a wrapper around the C function g_date_valid_weekday.
func DateValidWeekday(weekday DateWeekday) bool {
	c_weekday := (C.GDateWeekday)(weekday)

	retC := C.g_date_valid_weekday(c_weekday)
	retGo := retC == C.TRUE

	return retGo
}

// DateValidYear is a wrapper around the C function g_date_valid_year.
func DateValidYear(year DateYear) bool {
	c_year := (C.GDateYear)(year)

	retC := C.g_date_valid_year(c_year)
	retGo := retC == C.TRUE

	return retGo
}

// DirectEqual is a wrapper around the C function g_direct_equal.
func DirectEqual(v1 uintptr, v2 uintptr) bool {
	c_v1 := (C.gconstpointer)(v1)

	c_v2 := (C.gconstpointer)(v2)

	retC := C.g_direct_equal(c_v1, c_v2)
	retGo := retC == C.TRUE

	return retGo
}

// DirectHash is a wrapper around the C function g_direct_hash.
func DirectHash(v uintptr) uint32 {
	c_v := (C.gconstpointer)(v)

	retC := C.g_direct_hash(c_v)
	retGo := (uint32)(retC)

	return retGo
}

// Unsupported : g_environ_getenv : unsupported parameter envp : no type generator for filename () for array param envp

// Unsupported : g_environ_setenv : unsupported parameter envp : no type generator for filename () for array param envp

// Unsupported : g_environ_unsetenv : unsupported parameter envp : no type generator for filename () for array param envp

// FileErrorFromErrno is a wrapper around the C function g_file_error_from_errno.
func FileErrorFromErrno(errNo int32) FileError {
	c_err_no := (C.gint)(errNo)

	retC := C.g_file_error_from_errno(c_err_no)
	retGo := (FileError)(retC)

	return retGo
}

// FileErrorQuark is a wrapper around the C function g_file_error_quark.
func FileErrorQuark() Quark {
	retC := C.g_file_error_quark()
	retGo := (Quark)(retC)

	return retGo
}

// Unsupported : g_file_get_contents : unsupported parameter contents : no type generator for guint8 () for array param contents

// FileOpenTmp is a wrapper around the C function g_file_open_tmp.
func FileOpenTmp(tmpl string) (int32, string, error) {
	c_tmpl := C.CString(tmpl)
	defer C.free(unsafe.Pointer(c_tmpl))

	var c_name_used *C.gchar

	var cThrowableError *C.GError

	retC := C.g_file_open_tmp(c_tmpl, &c_name_used, &cThrowableError)
	retGo := (int32)(retC)

	goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	nameUsed := C.GoString(c_name_used)
	defer C.free(unsafe.Pointer(c_name_used))

	return retGo, nameUsed, goThrowableError
}

// Unsupported : g_file_set_contents : unsupported parameter contents : no type generator for guint8 () for array param contents

// FileTest is a wrapper around the C function g_file_test.
func FileTest(filename string, test GFileTest) bool {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))

	c_test := (C.GFileTest)(test)

	retC := C.g_file_test(c_filename, c_test)
	retGo := retC == C.TRUE

	return retGo
}

// FilenameFromUri is a wrapper around the C function g_filename_from_uri.
func FilenameFromUri(uri string) (string, string, error) {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	var c_hostname *C.gchar

	var cThrowableError *C.GError

	retC := C.g_filename_from_uri(c_uri, &c_hostname, &cThrowableError)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	hostname := C.GoString(c_hostname)
	defer C.free(unsafe.Pointer(c_hostname))

	return retGo, hostname, goThrowableError
}

// FilenameFromUtf8 is a wrapper around the C function g_filename_from_utf8.
func FilenameFromUtf8(utf8string string, len int64) (string, uint64, uint64, error) {
	c_utf8string := C.CString(utf8string)
	defer C.free(unsafe.Pointer(c_utf8string))

	c_len := (C.gssize)(len)

	var c_bytes_read C.gsize

	var c_bytes_written C.gsize

	var cThrowableError *C.GError

	retC := C.g_filename_from_utf8(c_utf8string, c_len, &c_bytes_read, &c_bytes_written, &cThrowableError)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	bytesRead := (uint64)(c_bytes_read)

	bytesWritten := (uint64)(c_bytes_written)

	return retGo, bytesRead, bytesWritten, goThrowableError
}

// FilenameToUri is a wrapper around the C function g_filename_to_uri.
func FilenameToUri(filename string, hostname string) (string, error) {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))

	c_hostname := C.CString(hostname)
	defer C.free(unsafe.Pointer(c_hostname))

	var cThrowableError *C.GError

	retC := C.g_filename_to_uri(c_filename, c_hostname, &cThrowableError)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// FilenameToUtf8 is a wrapper around the C function g_filename_to_utf8.
func FilenameToUtf8(opsysstring string, len int64) (string, uint64, uint64, error) {
	c_opsysstring := C.CString(opsysstring)
	defer C.free(unsafe.Pointer(c_opsysstring))

	c_len := (C.gssize)(len)

	var c_bytes_read C.gsize

	var c_bytes_written C.gsize

	var cThrowableError *C.GError

	retC := C.g_filename_to_utf8(c_opsysstring, c_len, &c_bytes_read, &c_bytes_written, &cThrowableError)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	bytesRead := (uint64)(c_bytes_read)

	bytesWritten := (uint64)(c_bytes_written)

	return retGo, bytesRead, bytesWritten, goThrowableError
}

// FindProgramInPath is a wrapper around the C function g_find_program_in_path.
func FindProgramInPath(program string) string {
	c_program := C.CString(program)
	defer C.free(unsafe.Pointer(c_program))

	retC := C.g_find_program_in_path(c_program)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_fprintf : unsupported parameter file : no type generator for gpointer (FILE*) for param file

// Free is a wrapper around the C function g_free.
func Free(mem uintptr) {
	c_mem := (C.gpointer)(mem)

	C.g_free(c_mem)

	return
}

// GetCharset is a wrapper around the C function g_get_charset.
func GetCharset() (bool, string) {
	var c_charset *C.char

	retC := C.g_get_charset(&c_charset)
	retGo := retC == C.TRUE

	charset := C.GoString(c_charset)

	return retGo, charset
}

// GetCodeset is a wrapper around the C function g_get_codeset.
func GetCodeset() string {
	retC := C.g_get_codeset()
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// GetCurrentDir is a wrapper around the C function g_get_current_dir.
func GetCurrentDir() string {
	retC := C.g_get_current_dir()
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// GetCurrentTime is a wrapper around the C function g_get_current_time.
func GetCurrentTime(result *TimeVal) {
	c_result := (*C.GTimeVal)(result.ToC())

	C.g_get_current_time(c_result)

	return
}

// Unsupported : g_get_environ : no return type

// Unsupported : g_get_filename_charsets : unsupported parameter charsets : in string with indirection level of 3

// GetHomeDir is a wrapper around the C function g_get_home_dir.
func GetHomeDir() string {
	retC := C.g_get_home_dir()
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : g_get_language_names : no return type

// Unsupported : g_get_locale_variants : no return type

// GetPrgname is a wrapper around the C function g_get_prgname.
func GetPrgname() string {
	retC := C.g_get_prgname()
	retGo := C.GoString(retC)

	return retGo
}

// GetRealName is a wrapper around the C function g_get_real_name.
func GetRealName() string {
	retC := C.g_get_real_name()
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : g_get_system_config_dirs : no return type

// Unsupported : g_get_system_data_dirs : no return type

// GetTmpDir is a wrapper around the C function g_get_tmp_dir.
func GetTmpDir() string {
	retC := C.g_get_tmp_dir()
	retGo := C.GoString(retC)

	return retGo
}

// GetUserName is a wrapper around the C function g_get_user_name.
func GetUserName() string {
	retC := C.g_get_user_name()
	retGo := C.GoString(retC)

	return retGo
}

// Getenv is a wrapper around the C function g_getenv.
func Getenv(variable string) string {
	c_variable := C.CString(variable)
	defer C.free(unsafe.Pointer(c_variable))

	retC := C.g_getenv(c_variable)
	retGo := C.GoString(retC)

	return retGo
}

// HashTableDestroy is a wrapper around the C function g_hash_table_destroy.
func HashTableDestroy(hashTable *HashTable) {
	c_hash_table := (*C.GHashTable)(hashTable.ToC())

	C.g_hash_table_destroy(c_hash_table)

	return
}

// HashTableInsert is a wrapper around the C function g_hash_table_insert.
func HashTableInsert(hashTable *HashTable, key uintptr, value uintptr) bool {
	c_hash_table := (*C.GHashTable)(hashTable.ToC())

	c_key := (C.gpointer)(key)

	c_value := (C.gpointer)(value)

	retC := C.g_hash_table_insert(c_hash_table, c_key, c_value)
	retGo := retC == C.TRUE

	return retGo
}

// HashTableLookup is a wrapper around the C function g_hash_table_lookup.
func HashTableLookup(hashTable *HashTable, key uintptr) uintptr {
	c_hash_table := (*C.GHashTable)(hashTable.ToC())

	c_key := (C.gconstpointer)(key)

	retC := C.g_hash_table_lookup(c_hash_table, c_key)
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// HashTableLookupExtended is a wrapper around the C function g_hash_table_lookup_extended.
func HashTableLookupExtended(hashTable *HashTable, lookupKey uintptr) (bool, uintptr, uintptr) {
	c_hash_table := (*C.GHashTable)(hashTable.ToC())

	c_lookup_key := (C.gconstpointer)(lookupKey)

	var c_orig_key C.gpointer

	var c_value C.gpointer

	retC := C.g_hash_table_lookup_extended(c_hash_table, c_lookup_key, &c_orig_key, &c_value)
	retGo := retC == C.TRUE

	origKey := (uintptr)(unsafe.Pointer(&c_orig_key))

	value := (uintptr)(unsafe.Pointer(&c_value))

	return retGo, origKey, value
}

// HashTableRemove is a wrapper around the C function g_hash_table_remove.
func HashTableRemove(hashTable *HashTable, key uintptr) bool {
	c_hash_table := (*C.GHashTable)(hashTable.ToC())

	c_key := (C.gconstpointer)(key)

	retC := C.g_hash_table_remove(c_hash_table, c_key)
	retGo := retC == C.TRUE

	return retGo
}

// HashTableReplace is a wrapper around the C function g_hash_table_replace.
func HashTableReplace(hashTable *HashTable, key uintptr, value uintptr) bool {
	c_hash_table := (*C.GHashTable)(hashTable.ToC())

	c_key := (C.gpointer)(key)

	c_value := (C.gpointer)(value)

	retC := C.g_hash_table_replace(c_hash_table, c_key, c_value)
	retGo := retC == C.TRUE

	return retGo
}

// HashTableSize is a wrapper around the C function g_hash_table_size.
func HashTableSize(hashTable *HashTable) uint32 {
	c_hash_table := (*C.GHashTable)(hashTable.ToC())

	retC := C.g_hash_table_size(c_hash_table)
	retGo := (uint32)(retC)

	return retGo
}

// HashTableSteal is a wrapper around the C function g_hash_table_steal.
func HashTableSteal(hashTable *HashTable, key uintptr) bool {
	c_hash_table := (*C.GHashTable)(hashTable.ToC())

	c_key := (C.gconstpointer)(key)

	retC := C.g_hash_table_steal(c_hash_table, c_key)
	retGo := retC == C.TRUE

	return retGo
}

// HookDestroy is a wrapper around the C function g_hook_destroy.
func HookDestroy(hookList *HookList, hookId uint64) bool {
	c_hook_list := (*C.GHookList)(hookList.ToC())

	c_hook_id := (C.gulong)(hookId)

	retC := C.g_hook_destroy(c_hook_list, c_hook_id)
	retGo := retC == C.TRUE

	return retGo
}

// HookDestroyLink is a wrapper around the C function g_hook_destroy_link.
func HookDestroyLink(hookList *HookList, hook *Hook) {
	c_hook_list := (*C.GHookList)(hookList.ToC())

	c_hook := (*C.GHook)(hook.ToC())

	C.g_hook_destroy_link(c_hook_list, c_hook)

	return
}

// HookFree is a wrapper around the C function g_hook_free.
func HookFree(hookList *HookList, hook *Hook) {
	c_hook_list := (*C.GHookList)(hookList.ToC())

	c_hook := (*C.GHook)(hook.ToC())

	C.g_hook_free(c_hook_list, c_hook)

	return
}

// HookInsertBefore is a wrapper around the C function g_hook_insert_before.
func HookInsertBefore(hookList *HookList, sibling *Hook, hook *Hook) {
	c_hook_list := (*C.GHookList)(hookList.ToC())

	c_sibling := (*C.GHook)(sibling.ToC())

	c_hook := (*C.GHook)(hook.ToC())

	C.g_hook_insert_before(c_hook_list, c_sibling, c_hook)

	return
}

// HookPrepend is a wrapper around the C function g_hook_prepend.
func HookPrepend(hookList *HookList, hook *Hook) {
	c_hook_list := (*C.GHookList)(hookList.ToC())

	c_hook := (*C.GHook)(hook.ToC())

	C.g_hook_prepend(c_hook_list, c_hook)

	return
}

// HookUnref is a wrapper around the C function g_hook_unref.
func HookUnref(hookList *HookList, hook *Hook) {
	c_hook_list := (*C.GHookList)(hookList.ToC())

	c_hook := (*C.GHook)(hook.ToC())

	C.g_hook_unref(c_hook_list, c_hook)

	return
}

// Unsupported : g_iconv : unsupported parameter converter : Blacklisted record : GIConv

// Unsupported : g_iconv_open : return type : Blacklisted record : GIConv

// Unsupported : g_idle_add : unsupported parameter function : no type generator for SourceFunc (GSourceFunc) for param function

// Unsupported : g_idle_add_full : unsupported parameter function : no type generator for SourceFunc (GSourceFunc) for param function

// IdleRemoveByData is a wrapper around the C function g_idle_remove_by_data.
func IdleRemoveByData(data uintptr) bool {
	c_data := (C.gpointer)(data)

	retC := C.g_idle_remove_by_data(c_data)
	retGo := retC == C.TRUE

	return retGo
}

// IdleSourceNew is a wrapper around the C function g_idle_source_new.
func IdleSourceNew() *Source {
	retC := C.g_idle_source_new()
	retGo := SourceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// IntEqual is a wrapper around the C function g_int_equal.
func IntEqual(v1 uintptr, v2 uintptr) bool {
	c_v1 := (C.gconstpointer)(v1)

	c_v2 := (C.gconstpointer)(v2)

	retC := C.g_int_equal(c_v1, c_v2)
	retGo := retC == C.TRUE

	return retGo
}

// IntHash is a wrapper around the C function g_int_hash.
func IntHash(v uintptr) uint32 {
	c_v := (C.gconstpointer)(v)

	retC := C.g_int_hash(c_v)
	retGo := (uint32)(retC)

	return retGo
}

// Unsupported : g_io_add_watch : unsupported parameter channel : Blacklisted record : GIOChannel

// Unsupported : g_io_add_watch_full : unsupported parameter channel : Blacklisted record : GIOChannel

// IoChannelErrorFromErrno is a wrapper around the C function g_io_channel_error_from_errno.
func IoChannelErrorFromErrno(en int32) IOChannelError {
	c_en := (C.gint)(en)

	retC := C.g_io_channel_error_from_errno(c_en)
	retGo := (IOChannelError)(retC)

	return retGo
}

// IoChannelErrorQuark is a wrapper around the C function g_io_channel_error_quark.
func IoChannelErrorQuark() Quark {
	retC := C.g_io_channel_error_quark()
	retGo := (Quark)(retC)

	return retGo
}

// Unsupported : g_io_create_watch : unsupported parameter channel : Blacklisted record : GIOChannel

// KeyFileErrorQuark is a wrapper around the C function g_key_file_error_quark.
func KeyFileErrorQuark() Quark {
	retC := C.g_key_file_error_quark()
	retGo := (Quark)(retC)

	return retGo
}

// Unsupported : g_listenv : no return type

// Unsupported : g_locale_from_utf8 : no return type

// Unsupported : g_locale_to_utf8 : unsupported parameter opsysstring : no type generator for guint8 () for array param opsysstring

// Unsupported : g_log : unsupported parameter ... : varargs

// LogDefaultHandler is a wrapper around the C function g_log_default_handler.
func LogDefaultHandler(logDomain string, logLevel LogLevelFlags, message string, unusedData uintptr) {
	c_log_domain := C.CString(logDomain)
	defer C.free(unsafe.Pointer(c_log_domain))

	c_log_level := (C.GLogLevelFlags)(logLevel)

	c_message := C.CString(message)
	defer C.free(unsafe.Pointer(c_message))

	c_unused_data := (C.gpointer)(unusedData)

	C.g_log_default_handler(c_log_domain, c_log_level, c_message, c_unused_data)

	return
}

// LogRemoveHandler is a wrapper around the C function g_log_remove_handler.
func LogRemoveHandler(logDomain string, handlerId uint32) {
	c_log_domain := C.CString(logDomain)
	defer C.free(unsafe.Pointer(c_log_domain))

	c_handler_id := (C.guint)(handlerId)

	C.g_log_remove_handler(c_log_domain, c_handler_id)

	return
}

// LogSetAlwaysFatal is a wrapper around the C function g_log_set_always_fatal.
func LogSetAlwaysFatal(fatalMask LogLevelFlags) LogLevelFlags {
	c_fatal_mask := (C.GLogLevelFlags)(fatalMask)

	retC := C.g_log_set_always_fatal(c_fatal_mask)
	retGo := (LogLevelFlags)(retC)

	return retGo
}

// Unsupported : g_log_set_default_handler : unsupported parameter log_func : no type generator for LogFunc (GLogFunc) for param log_func

// LogSetFatalMask is a wrapper around the C function g_log_set_fatal_mask.
func LogSetFatalMask(logDomain string, fatalMask LogLevelFlags) LogLevelFlags {
	c_log_domain := C.CString(logDomain)
	defer C.free(unsafe.Pointer(c_log_domain))

	c_fatal_mask := (C.GLogLevelFlags)(fatalMask)

	retC := C.g_log_set_fatal_mask(c_log_domain, c_fatal_mask)
	retGo := (LogLevelFlags)(retC)

	return retGo
}

// Unsupported : g_log_set_handler : unsupported parameter log_func : no type generator for LogFunc (GLogFunc) for param log_func

// Unsupported : g_log_set_handler_full : unsupported parameter log_func : no type generator for LogFunc (GLogFunc) for param log_func

// Unsupported : g_log_set_writer_func : unsupported parameter func : no type generator for LogWriterFunc (GLogWriterFunc) for param func

// Unsupported : g_log_structured : unsupported parameter ... : varargs

// Unsupported : g_log_structured_array : unsupported parameter fields : no type generator for LogField (GLogField) for array param fields

// Unsupported : g_log_structured_standard : unsupported parameter ... : varargs

// Unsupported : g_log_variant : unsupported parameter fields : Blacklisted record : GVariant

// Unsupported : g_log_writer_default : unsupported parameter fields : no type generator for LogField (GLogField) for array param fields

// Unsupported : g_log_writer_format_fields : unsupported parameter fields : no type generator for LogField (GLogField) for array param fields

// Unsupported : g_log_writer_journald : unsupported parameter fields : no type generator for LogField (GLogField) for array param fields

// Unsupported : g_log_writer_standard_streams : unsupported parameter fields : no type generator for LogField (GLogField) for array param fields

// Unsupported : g_logv : unsupported parameter args : no type generator for va_list (va_list) for param args

// MainContextDefault is a wrapper around the C function g_main_context_default.
func MainContextDefault() *MainContext {
	retC := C.g_main_context_default()
	retGo := MainContextNewFromC(unsafe.Pointer(retC))

	return retGo
}

// MainDepth is a wrapper around the C function g_main_depth.
func MainDepth() int32 {
	retC := C.g_main_depth()
	retGo := (int32)(retC)

	return retGo
}

// Malloc is a wrapper around the C function g_malloc.
func Malloc(nBytes uint64) uintptr {
	c_n_bytes := (C.gsize)(nBytes)

	retC := C.g_malloc(c_n_bytes)
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// Malloc0 is a wrapper around the C function g_malloc0.
func Malloc0(nBytes uint64) uintptr {
	c_n_bytes := (C.gsize)(nBytes)

	retC := C.g_malloc0(c_n_bytes)
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_markup_collect_attributes : unsupported parameter attribute_names : in string with indirection level of 2

// MarkupErrorQuark is a wrapper around the C function g_markup_error_quark.
func MarkupErrorQuark() Quark {
	retC := C.g_markup_error_quark()
	retGo := (Quark)(retC)

	return retGo
}

// MarkupEscapeText is a wrapper around the C function g_markup_escape_text.
func MarkupEscapeText(text string, length int64) string {
	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))

	c_length := (C.gssize)(length)

	retC := C.g_markup_escape_text(c_text, c_length)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_markup_printf_escaped : unsupported parameter ... : varargs

// Unsupported : g_markup_vprintf_escaped : unsupported parameter args : no type generator for va_list (va_list) for param args

// MemIsSystemMalloc is a wrapper around the C function g_mem_is_system_malloc.
func MemIsSystemMalloc() bool {
	retC := C.g_mem_is_system_malloc()
	retGo := retC == C.TRUE

	return retGo
}

// MemProfile is a wrapper around the C function g_mem_profile.
func MemProfile() {
	C.g_mem_profile()

	return
}

// MemSetVtable is a wrapper around the C function g_mem_set_vtable.
func MemSetVtable(vtable *MemVTable) {
	c_vtable := (*C.GMemVTable)(vtable.ToC())

	C.g_mem_set_vtable(c_vtable)

	return
}

// Memdup is a wrapper around the C function g_memdup.
func Memdup(mem uintptr, byteSize uint32) uintptr {
	c_mem := (C.gconstpointer)(mem)

	c_byte_size := (C.guint)(byteSize)

	retC := C.g_memdup(c_mem, c_byte_size)
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// Mkstemp is a wrapper around the C function g_mkstemp.
func Mkstemp(tmpl string) int32 {
	c_tmpl := C.CString(tmpl)
	defer C.free(unsafe.Pointer(c_tmpl))

	retC := C.g_mkstemp(c_tmpl)
	retGo := (int32)(retC)

	return retGo
}

// NullifyPointer is a wrapper around the C function g_nullify_pointer.
func NullifyPointer(nullifyLocation uintptr) {
	c_nullify_location := (C.gpointer)(nullifyLocation)

	C.g_nullify_pointer(&c_nullify_location)

	return
}

// Blacklisted : g_number_parser_error_quark

// OnErrorQuery is a wrapper around the C function g_on_error_query.
func OnErrorQuery(prgName string) {
	c_prg_name := C.CString(prgName)
	defer C.free(unsafe.Pointer(c_prg_name))

	C.g_on_error_query(c_prg_name)

	return
}

// OnErrorStackTrace is a wrapper around the C function g_on_error_stack_trace.
func OnErrorStackTrace(prgName string) {
	c_prg_name := C.CString(prgName)
	defer C.free(unsafe.Pointer(c_prg_name))

	C.g_on_error_stack_trace(c_prg_name)

	return
}

// Unsupported : g_once_init_enter : unsupported parameter location : no type generator for gpointer (void*) for param location

// Unsupported : g_once_init_leave : unsupported parameter location : no type generator for gpointer (void*) for param location

// OptionErrorQuark is a wrapper around the C function g_option_error_quark.
func OptionErrorQuark() Quark {
	retC := C.g_option_error_quark()
	retGo := (Quark)(retC)

	return retGo
}

// Unsupported : g_parse_debug_string : unsupported parameter keys : no type generator for DebugKey (GDebugKey) for array param keys

// PathGetBasename is a wrapper around the C function g_path_get_basename.
func PathGetBasename(fileName string) string {
	c_file_name := C.CString(fileName)
	defer C.free(unsafe.Pointer(c_file_name))

	retC := C.g_path_get_basename(c_file_name)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// PathGetDirname is a wrapper around the C function g_path_get_dirname.
func PathGetDirname(fileName string) string {
	c_file_name := C.CString(fileName)
	defer C.free(unsafe.Pointer(c_file_name))

	retC := C.g_path_get_dirname(c_file_name)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// PathIsAbsolute is a wrapper around the C function g_path_is_absolute.
func PathIsAbsolute(fileName string) bool {
	c_file_name := C.CString(fileName)
	defer C.free(unsafe.Pointer(c_file_name))

	retC := C.g_path_is_absolute(c_file_name)
	retGo := retC == C.TRUE

	return retGo
}

// PathSkipRoot is a wrapper around the C function g_path_skip_root.
func PathSkipRoot(fileName string) string {
	c_file_name := C.CString(fileName)
	defer C.free(unsafe.Pointer(c_file_name))

	retC := C.g_path_skip_root(c_file_name)
	retGo := C.GoString(retC)

	return retGo
}

// PatternMatch is a wrapper around the C function g_pattern_match.
func PatternMatch(pspec *PatternSpec, stringLength uint32, string string, stringReversed string) bool {
	c_pspec := (*C.GPatternSpec)(pspec.ToC())

	c_string_length := (C.guint)(stringLength)

	c_string := C.CString(string)
	defer C.free(unsafe.Pointer(c_string))

	c_string_reversed := C.CString(stringReversed)
	defer C.free(unsafe.Pointer(c_string_reversed))

	retC := C.g_pattern_match(c_pspec, c_string_length, c_string, c_string_reversed)
	retGo := retC == C.TRUE

	return retGo
}

// PatternMatchSimple is a wrapper around the C function g_pattern_match_simple.
func PatternMatchSimple(pattern string, string string) bool {
	c_pattern := C.CString(pattern)
	defer C.free(unsafe.Pointer(c_pattern))

	c_string := C.CString(string)
	defer C.free(unsafe.Pointer(c_string))

	retC := C.g_pattern_match_simple(c_pattern, c_string)
	retGo := retC == C.TRUE

	return retGo
}

// PatternMatchString is a wrapper around the C function g_pattern_match_string.
func PatternMatchString(pspec *PatternSpec, string string) bool {
	c_pspec := (*C.GPatternSpec)(pspec.ToC())

	c_string := C.CString(string)
	defer C.free(unsafe.Pointer(c_string))

	retC := C.g_pattern_match_string(c_pspec, c_string)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : g_pointer_bit_lock : unsupported parameter address : no type generator for gpointer (void*) for param address

// Unsupported : g_pointer_bit_trylock : unsupported parameter address : no type generator for gpointer (void*) for param address

// Unsupported : g_pointer_bit_unlock : unsupported parameter address : no type generator for gpointer (void*) for param address

// Unsupported : g_prefix_error : unsupported parameter err : record with indirection level of 2

// Unsupported : g_print : unsupported parameter ... : varargs

// Unsupported : g_printerr : unsupported parameter ... : varargs

// Unsupported : g_printf : unsupported parameter ... : varargs

// Unsupported : g_printf_string_upper_bound : unsupported parameter args : no type generator for va_list (va_list) for param args

// Unsupported : g_propagate_error : unsupported parameter dest : record with indirection level of 2

// Unsupported : g_propagate_prefixed_error : unsupported parameter dest : record with indirection level of 2

// Unsupported : g_ptr_array_find : unsupported parameter haystack : no type generator for gpointer (gpointer) for array param haystack

// Unsupported : g_ptr_array_find_with_equal_func : unsupported parameter haystack : no type generator for gpointer (gpointer) for array param haystack

// Unsupported : g_qsort_with_data : unsupported parameter compare_func : no type generator for CompareDataFunc (GCompareDataFunc) for param compare_func

// QuarkFromStaticString is a wrapper around the C function g_quark_from_static_string.
func QuarkFromStaticString(string string) Quark {
	c_string := C.CString(string)
	defer C.free(unsafe.Pointer(c_string))

	retC := C.g_quark_from_static_string(c_string)
	retGo := (Quark)(retC)

	return retGo
}

// QuarkFromString is a wrapper around the C function g_quark_from_string.
func QuarkFromString(string string) Quark {
	c_string := C.CString(string)
	defer C.free(unsafe.Pointer(c_string))

	retC := C.g_quark_from_string(c_string)
	retGo := (Quark)(retC)

	return retGo
}

// QuarkToString is a wrapper around the C function g_quark_to_string.
func QuarkToString(quark Quark) string {
	c_quark := (C.GQuark)(quark)

	retC := C.g_quark_to_string(c_quark)
	retGo := C.GoString(retC)

	return retGo
}

// QuarkTryString is a wrapper around the C function g_quark_try_string.
func QuarkTryString(string string) Quark {
	c_string := C.CString(string)
	defer C.free(unsafe.Pointer(c_string))

	retC := C.g_quark_try_string(c_string)
	retGo := (Quark)(retC)

	return retGo
}

// RandomDouble is a wrapper around the C function g_random_double.
func RandomDouble() float64 {
	retC := C.g_random_double()
	retGo := (float64)(retC)

	return retGo
}

// RandomDoubleRange is a wrapper around the C function g_random_double_range.
func RandomDoubleRange(begin float64, end float64) float64 {
	c_begin := (C.gdouble)(begin)

	c_end := (C.gdouble)(end)

	retC := C.g_random_double_range(c_begin, c_end)
	retGo := (float64)(retC)

	return retGo
}

// RandomInt is a wrapper around the C function g_random_int.
func RandomInt() uint32 {
	retC := C.g_random_int()
	retGo := (uint32)(retC)

	return retGo
}

// RandomIntRange is a wrapper around the C function g_random_int_range.
func RandomIntRange(begin int32, end int32) int32 {
	c_begin := (C.gint32)(begin)

	c_end := (C.gint32)(end)

	retC := C.g_random_int_range(c_begin, c_end)
	retGo := (int32)(retC)

	return retGo
}

// RandomSetSeed is a wrapper around the C function g_random_set_seed.
func RandomSetSeed(seed uint32) {
	c_seed := (C.guint32)(seed)

	C.g_random_set_seed(c_seed)

	return
}

// Realloc is a wrapper around the C function g_realloc.
func Realloc(mem uintptr, nBytes uint64) uintptr {
	c_mem := (C.gpointer)(mem)

	c_n_bytes := (C.gsize)(nBytes)

	retC := C.g_realloc(c_mem, c_n_bytes)
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// RegexErrorQuark is a wrapper around the C function g_regex_error_quark.
func RegexErrorQuark() Quark {
	retC := C.g_regex_error_quark()
	retGo := (Quark)(retC)

	return retGo
}

// Unsupported : g_regex_escape_string : unsupported parameter string : no type generator for utf8 (gchar) for array param string

// Unsupported : g_regex_split_simple : no return type

// ReturnIfFailWarning is a wrapper around the C function g_return_if_fail_warning.
func ReturnIfFailWarning(logDomain string, prettyFunction string, expression string) {
	c_log_domain := C.CString(logDomain)
	defer C.free(unsafe.Pointer(c_log_domain))

	c_pretty_function := C.CString(prettyFunction)
	defer C.free(unsafe.Pointer(c_pretty_function))

	c_expression := C.CString(expression)
	defer C.free(unsafe.Pointer(c_expression))

	C.g_return_if_fail_warning(c_log_domain, c_pretty_function, c_expression)

	return
}

// Unsupported : g_set_error : unsupported parameter err : record with indirection level of 2

// Unsupported : g_set_error_literal : unsupported parameter err : record with indirection level of 2

// SetPrgname is a wrapper around the C function g_set_prgname.
func SetPrgname(prgname string) {
	c_prgname := C.CString(prgname)
	defer C.free(unsafe.Pointer(c_prgname))

	C.g_set_prgname(c_prgname)

	return
}

// Unsupported : g_set_print_handler : unsupported parameter func : no type generator for PrintFunc (GPrintFunc) for param func

// Unsupported : g_set_printerr_handler : unsupported parameter func : no type generator for PrintFunc (GPrintFunc) for param func

// ShellErrorQuark is a wrapper around the C function g_shell_error_quark.
func ShellErrorQuark() Quark {
	retC := C.g_shell_error_quark()
	retGo := (Quark)(retC)

	return retGo
}

// Unsupported : g_shell_parse_argv : unsupported parameter argvp : no type generator for filename () for array param argvp

// ShellQuote is a wrapper around the C function g_shell_quote.
func ShellQuote(unquotedString string) string {
	c_unquoted_string := C.CString(unquotedString)
	defer C.free(unsafe.Pointer(c_unquoted_string))

	retC := C.g_shell_quote(c_unquoted_string)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// ShellUnquote is a wrapper around the C function g_shell_unquote.
func ShellUnquote(quotedString string) (string, error) {
	c_quoted_string := C.CString(quotedString)
	defer C.free(unsafe.Pointer(c_quoted_string))

	var cThrowableError *C.GError

	retC := C.g_shell_unquote(c_quoted_string, &cThrowableError)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// SliceGetConfig is a wrapper around the C function g_slice_get_config.
func SliceGetConfig(ckey SliceConfig) int64 {
	c_ckey := (C.GSliceConfig)(ckey)

	retC := C.g_slice_get_config(c_ckey)
	retGo := (int64)(retC)

	return retGo
}

// Blacklisted : g_slice_get_config_state

// SliceSetConfig is a wrapper around the C function g_slice_set_config.
func SliceSetConfig(ckey SliceConfig, value int64) {
	c_ckey := (C.GSliceConfig)(ckey)

	c_value := (C.gint64)(value)

	C.g_slice_set_config(c_ckey, c_value)

	return
}

// Unsupported : g_snprintf : unsupported parameter ... : varargs

// SourceRemove is a wrapper around the C function g_source_remove.
func SourceRemove(tag uint32) bool {
	c_tag := (C.guint)(tag)

	retC := C.g_source_remove(c_tag)
	retGo := retC == C.TRUE

	return retGo
}

// SourceRemoveByFuncsUserData is a wrapper around the C function g_source_remove_by_funcs_user_data.
func SourceRemoveByFuncsUserData(funcs *SourceFuncs, userData uintptr) bool {
	c_funcs := (*C.GSourceFuncs)(funcs.ToC())

	c_user_data := (C.gpointer)(userData)

	retC := C.g_source_remove_by_funcs_user_data(c_funcs, c_user_data)
	retGo := retC == C.TRUE

	return retGo
}

// SourceRemoveByUserData is a wrapper around the C function g_source_remove_by_user_data.
func SourceRemoveByUserData(userData uintptr) bool {
	c_user_data := (C.gpointer)(userData)

	retC := C.g_source_remove_by_user_data(c_user_data)
	retGo := retC == C.TRUE

	return retGo
}

// SpacedPrimesClosest is a wrapper around the C function g_spaced_primes_closest.
func SpacedPrimesClosest(num uint32) uint32 {
	c_num := (C.guint)(num)

	retC := C.g_spaced_primes_closest(c_num)
	retGo := (uint32)(retC)

	return retGo
}

// Unsupported : g_spawn_async : unsupported parameter argv : no type generator for filename () for array param argv

// Unsupported : g_spawn_async_with_pipes : unsupported parameter argv : no type generator for filename () for array param argv

// SpawnClosePid is a wrapper around the C function g_spawn_close_pid.
func SpawnClosePid(pid Pid) {
	c_pid := (C.GPid)(pid)

	C.g_spawn_close_pid(c_pid)

	return
}

// SpawnCommandLineAsync is a wrapper around the C function g_spawn_command_line_async.
func SpawnCommandLineAsync(commandLine string) (bool, error) {
	c_command_line := C.CString(commandLine)
	defer C.free(unsafe.Pointer(c_command_line))

	var cThrowableError *C.GError

	retC := C.g_spawn_command_line_async(c_command_line, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_spawn_command_line_sync : unsupported parameter standard_output : no type generator for guint8 () for array param standard_output

// SpawnErrorQuark is a wrapper around the C function g_spawn_error_quark.
func SpawnErrorQuark() Quark {
	retC := C.g_spawn_error_quark()
	retGo := (Quark)(retC)

	return retGo
}

// SpawnExitErrorQuark is a wrapper around the C function g_spawn_exit_error_quark.
func SpawnExitErrorQuark() Quark {
	retC := C.g_spawn_exit_error_quark()
	retGo := (Quark)(retC)

	return retGo
}

// Unsupported : g_spawn_sync : unsupported parameter argv : no type generator for filename () for array param argv

// Unsupported : g_sprintf : unsupported parameter ... : varargs

// Stpcpy is a wrapper around the C function g_stpcpy.
func Stpcpy(dest string, src string) string {
	c_dest := C.CString(dest)
	defer C.free(unsafe.Pointer(c_dest))

	c_src := C.CString(src)
	defer C.free(unsafe.Pointer(c_src))

	retC := C.g_stpcpy(c_dest, c_src)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// StrEqual is a wrapper around the C function g_str_equal.
func StrEqual(v1 uintptr, v2 uintptr) bool {
	c_v1 := (C.gconstpointer)(v1)

	c_v2 := (C.gconstpointer)(v2)

	retC := C.g_str_equal(c_v1, c_v2)
	retGo := retC == C.TRUE

	return retGo
}

// StrHash is a wrapper around the C function g_str_hash.
func StrHash(v uintptr) uint32 {
	c_v := (C.gconstpointer)(v)

	retC := C.g_str_hash(c_v)
	retGo := (uint32)(retC)

	return retGo
}

// Unsupported : g_str_tokenize_and_fold : unsupported parameter ascii_alternates : no type generator for utf8 (gchar**) for array param ascii_alternates

// Strcanon is a wrapper around the C function g_strcanon.
func Strcanon(string string, validChars string, substitutor rune) string {
	c_string := C.CString(string)
	defer C.free(unsafe.Pointer(c_string))

	c_valid_chars := C.CString(validChars)
	defer C.free(unsafe.Pointer(c_valid_chars))

	c_substitutor := (C.gchar)(substitutor)

	retC := C.g_strcanon(c_string, c_valid_chars, c_substitutor)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Strcasecmp is a wrapper around the C function g_strcasecmp.
func Strcasecmp(s1 string, s2 string) int32 {
	c_s1 := C.CString(s1)
	defer C.free(unsafe.Pointer(c_s1))

	c_s2 := C.CString(s2)
	defer C.free(unsafe.Pointer(c_s2))

	retC := C.g_strcasecmp(c_s1, c_s2)
	retGo := (int32)(retC)

	return retGo
}

// Strchomp is a wrapper around the C function g_strchomp.
func Strchomp(string string) string {
	c_string := C.CString(string)
	defer C.free(unsafe.Pointer(c_string))

	retC := C.g_strchomp(c_string)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Strchug is a wrapper around the C function g_strchug.
func Strchug(string string) string {
	c_string := C.CString(string)
	defer C.free(unsafe.Pointer(c_string))

	retC := C.g_strchug(c_string)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Strcompress is a wrapper around the C function g_strcompress.
func Strcompress(source string) string {
	c_source := C.CString(source)
	defer C.free(unsafe.Pointer(c_source))

	retC := C.g_strcompress(c_source)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_strconcat : unsupported parameter ... : varargs

// Strdelimit is a wrapper around the C function g_strdelimit.
func Strdelimit(string string, delimiters string, newDelimiter rune) string {
	c_string := C.CString(string)
	defer C.free(unsafe.Pointer(c_string))

	c_delimiters := C.CString(delimiters)
	defer C.free(unsafe.Pointer(c_delimiters))

	c_new_delimiter := (C.gchar)(newDelimiter)

	retC := C.g_strdelimit(c_string, c_delimiters, c_new_delimiter)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Strdown is a wrapper around the C function g_strdown.
func Strdown(string string) string {
	c_string := C.CString(string)
	defer C.free(unsafe.Pointer(c_string))

	retC := C.g_strdown(c_string)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Strdup is a wrapper around the C function g_strdup.
func Strdup(str string) string {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	retC := C.g_strdup(c_str)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_strdup_printf : unsupported parameter ... : varargs

// Unsupported : g_strdup_vprintf : unsupported parameter args : no type generator for va_list (va_list) for param args

// Unsupported : g_strdupv : unsupported parameter str_array : in string with indirection level of 2

// Strerror is a wrapper around the C function g_strerror.
func Strerror(errnum int32) string {
	c_errnum := (C.gint)(errnum)

	retC := C.g_strerror(c_errnum)
	retGo := C.GoString(retC)

	return retGo
}

// Strescape is a wrapper around the C function g_strescape.
func Strescape(source string, exceptions string) string {
	c_source := C.CString(source)
	defer C.free(unsafe.Pointer(c_source))

	c_exceptions := C.CString(exceptions)
	defer C.free(unsafe.Pointer(c_exceptions))

	retC := C.g_strescape(c_source, c_exceptions)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_strfreev : unsupported parameter str_array : in string with indirection level of 2

// StringNew is a wrapper around the C function g_string_new.
func StringNew(init string) *String {
	c_init := C.CString(init)
	defer C.free(unsafe.Pointer(c_init))

	retC := C.g_string_new(c_init)
	retGo := StringNewFromC(unsafe.Pointer(retC))

	return retGo
}

// StringNewLen is a wrapper around the C function g_string_new_len.
func StringNewLen(init string, len int64) *String {
	c_init := C.CString(init)
	defer C.free(unsafe.Pointer(c_init))

	c_len := (C.gssize)(len)

	retC := C.g_string_new_len(c_init, c_len)
	retGo := StringNewFromC(unsafe.Pointer(retC))

	return retGo
}

// StringSizedNew is a wrapper around the C function g_string_sized_new.
func StringSizedNew(dflSize uint64) *String {
	c_dfl_size := (C.gsize)(dflSize)

	retC := C.g_string_sized_new(c_dfl_size)
	retGo := StringNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_strjoin : unsupported parameter ... : varargs

// Unsupported : g_strjoinv : unsupported parameter str_array : in string with indirection level of 2

// Strlcat is a wrapper around the C function g_strlcat.
func Strlcat(dest string, src string, destSize uint64) uint64 {
	c_dest := C.CString(dest)
	defer C.free(unsafe.Pointer(c_dest))

	c_src := C.CString(src)
	defer C.free(unsafe.Pointer(c_src))

	c_dest_size := (C.gsize)(destSize)

	retC := C.g_strlcat(c_dest, c_src, c_dest_size)
	retGo := (uint64)(retC)

	return retGo
}

// Strlcpy is a wrapper around the C function g_strlcpy.
func Strlcpy(dest string, src string, destSize uint64) uint64 {
	c_dest := C.CString(dest)
	defer C.free(unsafe.Pointer(c_dest))

	c_src := C.CString(src)
	defer C.free(unsafe.Pointer(c_src))

	c_dest_size := (C.gsize)(destSize)

	retC := C.g_strlcpy(c_dest, c_src, c_dest_size)
	retGo := (uint64)(retC)

	return retGo
}

// Strncasecmp is a wrapper around the C function g_strncasecmp.
func Strncasecmp(s1 string, s2 string, n uint32) int32 {
	c_s1 := C.CString(s1)
	defer C.free(unsafe.Pointer(c_s1))

	c_s2 := C.CString(s2)
	defer C.free(unsafe.Pointer(c_s2))

	c_n := (C.guint)(n)

	retC := C.g_strncasecmp(c_s1, c_s2, c_n)
	retGo := (int32)(retC)

	return retGo
}

// Strndup is a wrapper around the C function g_strndup.
func Strndup(str string, n uint64) string {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	c_n := (C.gsize)(n)

	retC := C.g_strndup(c_str, c_n)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Strnfill is a wrapper around the C function g_strnfill.
func Strnfill(length uint64, fillChar rune) string {
	c_length := (C.gsize)(length)

	c_fill_char := (C.gchar)(fillChar)

	retC := C.g_strnfill(c_length, c_fill_char)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Strreverse is a wrapper around the C function g_strreverse.
func Strreverse(string string) string {
	c_string := C.CString(string)
	defer C.free(unsafe.Pointer(c_string))

	retC := C.g_strreverse(c_string)
	retGo := C.GoString(retC)

	return retGo
}

// Strrstr is a wrapper around the C function g_strrstr.
func Strrstr(haystack string, needle string) string {
	c_haystack := C.CString(haystack)
	defer C.free(unsafe.Pointer(c_haystack))

	c_needle := C.CString(needle)
	defer C.free(unsafe.Pointer(c_needle))

	retC := C.g_strrstr(c_haystack, c_needle)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// StrrstrLen is a wrapper around the C function g_strrstr_len.
func StrrstrLen(haystack string, haystackLen int64, needle string) string {
	c_haystack := C.CString(haystack)
	defer C.free(unsafe.Pointer(c_haystack))

	c_haystack_len := (C.gssize)(haystackLen)

	c_needle := C.CString(needle)
	defer C.free(unsafe.Pointer(c_needle))

	retC := C.g_strrstr_len(c_haystack, c_haystack_len, c_needle)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Strsignal is a wrapper around the C function g_strsignal.
func Strsignal(signum int32) string {
	c_signum := (C.gint)(signum)

	retC := C.g_strsignal(c_signum)
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : g_strsplit : no return type

// Unsupported : g_strsplit_set : no return type

// StrstrLen is a wrapper around the C function g_strstr_len.
func StrstrLen(haystack string, haystackLen int64, needle string) string {
	c_haystack := C.CString(haystack)
	defer C.free(unsafe.Pointer(c_haystack))

	c_haystack_len := (C.gssize)(haystackLen)

	c_needle := C.CString(needle)
	defer C.free(unsafe.Pointer(c_needle))

	retC := C.g_strstr_len(c_haystack, c_haystack_len, c_needle)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Strtod is a wrapper around the C function g_strtod.
func Strtod(nptr string) (float64, string) {
	c_nptr := C.CString(nptr)
	defer C.free(unsafe.Pointer(c_nptr))

	var c_endptr *C.gchar

	retC := C.g_strtod(c_nptr, &c_endptr)
	retGo := (float64)(retC)

	endptr := C.GoString(c_endptr)

	return retGo, endptr
}

// Strup is a wrapper around the C function g_strup.
func Strup(string string) string {
	c_string := C.CString(string)
	defer C.free(unsafe.Pointer(c_string))

	retC := C.g_strup(c_string)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Blacklisted : g_strv_get_type

// Unsupported : g_strv_length : unsupported parameter str_array : in string with indirection level of 2

// Unsupported : g_test_add_data_func : unsupported parameter test_func : no type generator for TestDataFunc (GTestDataFunc) for param test_func

// Unsupported : g_test_add_data_func_full : unsupported parameter test_func : no type generator for TestDataFunc (GTestDataFunc) for param test_func

// Unsupported : g_test_add_func : unsupported parameter test_func : no type generator for TestFunc (GTestFunc) for param test_func

// Unsupported : g_test_add_vtable : unsupported parameter data_setup : no type generator for TestFixtureFunc (GTestFixtureFunc) for param data_setup

// TestAssertExpectedMessagesInternal is a wrapper around the C function g_test_assert_expected_messages_internal.
func TestAssertExpectedMessagesInternal(domain string, file string, line int32, func_ string) {
	c_domain := C.CString(domain)
	defer C.free(unsafe.Pointer(c_domain))

	c_file := C.CString(file)
	defer C.free(unsafe.Pointer(c_file))

	c_line := (C.int)(line)

	c_func := C.CString(func_)
	defer C.free(unsafe.Pointer(c_func))

	C.g_test_assert_expected_messages_internal(c_domain, c_file, c_line, c_func)

	return
}

// Unsupported : g_test_build_filename : unsupported parameter ... : varargs

// Unsupported : g_test_create_case : unsupported parameter data_setup : no type generator for TestFixtureFunc (GTestFixtureFunc) for param data_setup

// Unsupported : g_test_get_filename : unsupported parameter ... : varargs

// Unsupported : g_test_init : unsupported parameter argv : in string with indirection level of 3

// Unsupported : g_test_log_set_fatal_handler : unsupported parameter log_func : no type generator for TestLogFatalFunc (GTestLogFatalFunc) for param log_func

// TestLogTypeName is a wrapper around the C function g_test_log_type_name.
func TestLogTypeName(logType TestLogType) string {
	c_log_type := (C.GTestLogType)(logType)

	retC := C.g_test_log_type_name(c_log_type)
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : g_test_maximized_result : unsupported parameter ... : varargs

// Unsupported : g_test_message : unsupported parameter ... : varargs

// Unsupported : g_test_minimized_result : unsupported parameter ... : varargs

// Unsupported : g_test_queue_destroy : unsupported parameter destroy_func : no type generator for DestroyNotify (GDestroyNotify) for param destroy_func

// TestTrapAssertions is a wrapper around the C function g_test_trap_assertions.
func TestTrapAssertions(domain string, file string, line int32, func_ string, assertionFlags uint64, pattern string) {
	c_domain := C.CString(domain)
	defer C.free(unsafe.Pointer(c_domain))

	c_file := C.CString(file)
	defer C.free(unsafe.Pointer(c_file))

	c_line := (C.int)(line)

	c_func := C.CString(func_)
	defer C.free(unsafe.Pointer(c_func))

	c_assertion_flags := (C.guint64)(assertionFlags)

	c_pattern := C.CString(pattern)
	defer C.free(unsafe.Pointer(c_pattern))

	C.g_test_trap_assertions(c_domain, c_file, c_line, c_func, c_assertion_flags, c_pattern)

	return
}

// ThreadErrorQuark is a wrapper around the C function g_thread_error_quark.
func ThreadErrorQuark() Quark {
	retC := C.g_thread_error_quark()
	retGo := (Quark)(retC)

	return retGo
}

// ThreadExit is a wrapper around the C function g_thread_exit.
func ThreadExit(retval uintptr) {
	c_retval := (C.gpointer)(retval)

	C.g_thread_exit(c_retval)

	return
}

// ThreadPoolGetMaxUnusedThreads is a wrapper around the C function g_thread_pool_get_max_unused_threads.
func ThreadPoolGetMaxUnusedThreads() int32 {
	retC := C.g_thread_pool_get_max_unused_threads()
	retGo := (int32)(retC)

	return retGo
}

// ThreadPoolGetNumUnusedThreads is a wrapper around the C function g_thread_pool_get_num_unused_threads.
func ThreadPoolGetNumUnusedThreads() uint32 {
	retC := C.g_thread_pool_get_num_unused_threads()
	retGo := (uint32)(retC)

	return retGo
}

// ThreadPoolSetMaxUnusedThreads is a wrapper around the C function g_thread_pool_set_max_unused_threads.
func ThreadPoolSetMaxUnusedThreads(maxThreads int32) {
	c_max_threads := (C.gint)(maxThreads)

	C.g_thread_pool_set_max_unused_threads(c_max_threads)

	return
}

// ThreadPoolStopUnusedThreads is a wrapper around the C function g_thread_pool_stop_unused_threads.
func ThreadPoolStopUnusedThreads() {
	C.g_thread_pool_stop_unused_threads()

	return
}

// ThreadSelf is a wrapper around the C function g_thread_self.
func ThreadSelf() *Thread {
	retC := C.g_thread_self()
	retGo := ThreadNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ThreadYield is a wrapper around the C function g_thread_yield.
func ThreadYield() {
	C.g_thread_yield()

	return
}

// Unsupported : g_timeout_add : unsupported parameter function : no type generator for SourceFunc (GSourceFunc) for param function

// Unsupported : g_timeout_add_full : unsupported parameter function : no type generator for SourceFunc (GSourceFunc) for param function

// Unsupported : g_timeout_add_seconds : unsupported parameter function : no type generator for SourceFunc (GSourceFunc) for param function

// Unsupported : g_timeout_add_seconds_full : unsupported parameter function : no type generator for SourceFunc (GSourceFunc) for param function

// TimeoutSourceNew is a wrapper around the C function g_timeout_source_new.
func TimeoutSourceNew(interval uint32) *Source {
	c_interval := (C.guint)(interval)

	retC := C.g_timeout_source_new(c_interval)
	retGo := SourceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_trash_stack_height : unsupported parameter stack_p : record with indirection level of 2

// Unsupported : g_trash_stack_peek : unsupported parameter stack_p : record with indirection level of 2

// Unsupported : g_trash_stack_pop : unsupported parameter stack_p : record with indirection level of 2

// Unsupported : g_trash_stack_push : unsupported parameter stack_p : record with indirection level of 2

// TryMalloc is a wrapper around the C function g_try_malloc.
func TryMalloc(nBytes uint64) uintptr {
	c_n_bytes := (C.gsize)(nBytes)

	retC := C.g_try_malloc(c_n_bytes)
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// TryRealloc is a wrapper around the C function g_try_realloc.
func TryRealloc(mem uintptr, nBytes uint64) uintptr {
	c_mem := (C.gpointer)(mem)

	c_n_bytes := (C.gsize)(nBytes)

	retC := C.g_try_realloc(c_mem, c_n_bytes)
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_ucs4_to_utf16 : no return generator

// Ucs4ToUtf8 is a wrapper around the C function g_ucs4_to_utf8.
func Ucs4ToUtf8(str rune, len int64) (string, int64, int64, error) {
	c_str := (C.gunichar)(str)

	c_len := (C.glong)(len)

	var c_items_read C.glong

	var c_items_written C.glong

	var cThrowableError *C.GError

	retC := C.g_ucs4_to_utf8(&c_str, c_len, &c_items_read, &c_items_written, &cThrowableError)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	itemsRead := (int64)(c_items_read)

	itemsWritten := (int64)(c_items_written)

	return retGo, itemsRead, itemsWritten, goThrowableError
}

// UnicharBreakType is a wrapper around the C function g_unichar_break_type.
func UnicharBreakType(c rune) UnicodeBreakType {
	c_c := (C.gunichar)(c)

	retC := C.g_unichar_break_type(c_c)
	retGo := (UnicodeBreakType)(retC)

	return retGo
}

// UnicharDigitValue is a wrapper around the C function g_unichar_digit_value.
func UnicharDigitValue(c rune) int32 {
	c_c := (C.gunichar)(c)

	retC := C.g_unichar_digit_value(c_c)
	retGo := (int32)(retC)

	return retGo
}

// UnicharIsalnum is a wrapper around the C function g_unichar_isalnum.
func UnicharIsalnum(c rune) bool {
	c_c := (C.gunichar)(c)

	retC := C.g_unichar_isalnum(c_c)
	retGo := retC == C.TRUE

	return retGo
}

// UnicharIsalpha is a wrapper around the C function g_unichar_isalpha.
func UnicharIsalpha(c rune) bool {
	c_c := (C.gunichar)(c)

	retC := C.g_unichar_isalpha(c_c)
	retGo := retC == C.TRUE

	return retGo
}

// UnicharIscntrl is a wrapper around the C function g_unichar_iscntrl.
func UnicharIscntrl(c rune) bool {
	c_c := (C.gunichar)(c)

	retC := C.g_unichar_iscntrl(c_c)
	retGo := retC == C.TRUE

	return retGo
}

// UnicharIsdefined is a wrapper around the C function g_unichar_isdefined.
func UnicharIsdefined(c rune) bool {
	c_c := (C.gunichar)(c)

	retC := C.g_unichar_isdefined(c_c)
	retGo := retC == C.TRUE

	return retGo
}

// UnicharIsdigit is a wrapper around the C function g_unichar_isdigit.
func UnicharIsdigit(c rune) bool {
	c_c := (C.gunichar)(c)

	retC := C.g_unichar_isdigit(c_c)
	retGo := retC == C.TRUE

	return retGo
}

// UnicharIsgraph is a wrapper around the C function g_unichar_isgraph.
func UnicharIsgraph(c rune) bool {
	c_c := (C.gunichar)(c)

	retC := C.g_unichar_isgraph(c_c)
	retGo := retC == C.TRUE

	return retGo
}

// UnicharIslower is a wrapper around the C function g_unichar_islower.
func UnicharIslower(c rune) bool {
	c_c := (C.gunichar)(c)

	retC := C.g_unichar_islower(c_c)
	retGo := retC == C.TRUE

	return retGo
}

// UnicharIsprint is a wrapper around the C function g_unichar_isprint.
func UnicharIsprint(c rune) bool {
	c_c := (C.gunichar)(c)

	retC := C.g_unichar_isprint(c_c)
	retGo := retC == C.TRUE

	return retGo
}

// UnicharIspunct is a wrapper around the C function g_unichar_ispunct.
func UnicharIspunct(c rune) bool {
	c_c := (C.gunichar)(c)

	retC := C.g_unichar_ispunct(c_c)
	retGo := retC == C.TRUE

	return retGo
}

// UnicharIsspace is a wrapper around the C function g_unichar_isspace.
func UnicharIsspace(c rune) bool {
	c_c := (C.gunichar)(c)

	retC := C.g_unichar_isspace(c_c)
	retGo := retC == C.TRUE

	return retGo
}

// UnicharIstitle is a wrapper around the C function g_unichar_istitle.
func UnicharIstitle(c rune) bool {
	c_c := (C.gunichar)(c)

	retC := C.g_unichar_istitle(c_c)
	retGo := retC == C.TRUE

	return retGo
}

// UnicharIsupper is a wrapper around the C function g_unichar_isupper.
func UnicharIsupper(c rune) bool {
	c_c := (C.gunichar)(c)

	retC := C.g_unichar_isupper(c_c)
	retGo := retC == C.TRUE

	return retGo
}

// UnicharIswide is a wrapper around the C function g_unichar_iswide.
func UnicharIswide(c rune) bool {
	c_c := (C.gunichar)(c)

	retC := C.g_unichar_iswide(c_c)
	retGo := retC == C.TRUE

	return retGo
}

// UnicharIsxdigit is a wrapper around the C function g_unichar_isxdigit.
func UnicharIsxdigit(c rune) bool {
	c_c := (C.gunichar)(c)

	retC := C.g_unichar_isxdigit(c_c)
	retGo := retC == C.TRUE

	return retGo
}

// Blacklisted : g_unichar_to_utf8

// UnicharTolower is a wrapper around the C function g_unichar_tolower.
func UnicharTolower(c rune) rune {
	c_c := (C.gunichar)(c)

	retC := C.g_unichar_tolower(c_c)
	retGo := (rune)(retC)

	return retGo
}

// UnicharTotitle is a wrapper around the C function g_unichar_totitle.
func UnicharTotitle(c rune) rune {
	c_c := (C.gunichar)(c)

	retC := C.g_unichar_totitle(c_c)
	retGo := (rune)(retC)

	return retGo
}

// UnicharToupper is a wrapper around the C function g_unichar_toupper.
func UnicharToupper(c rune) rune {
	c_c := (C.gunichar)(c)

	retC := C.g_unichar_toupper(c_c)
	retGo := (rune)(retC)

	return retGo
}

// UnicharType is a wrapper around the C function g_unichar_type.
func UnicharType(c rune) UnicodeType {
	c_c := (C.gunichar)(c)

	retC := C.g_unichar_type(c_c)
	retGo := (UnicodeType)(retC)

	return retGo
}

// UnicharValidate is a wrapper around the C function g_unichar_validate.
func UnicharValidate(ch rune) bool {
	c_ch := (C.gunichar)(ch)

	retC := C.g_unichar_validate(c_ch)
	retGo := retC == C.TRUE

	return retGo
}

// UnicharXdigitValue is a wrapper around the C function g_unichar_xdigit_value.
func UnicharXdigitValue(c rune) int32 {
	c_c := (C.gunichar)(c)

	retC := C.g_unichar_xdigit_value(c_c)
	retGo := (int32)(retC)

	return retGo
}

// Blacklisted : g_unicode_canonical_decomposition

// UnicodeCanonicalOrdering is a wrapper around the C function g_unicode_canonical_ordering.
func UnicodeCanonicalOrdering(string rune, len uint64) {
	c_string := (C.gunichar)(string)

	c_len := (C.gsize)(len)

	C.g_unicode_canonical_ordering(&c_string, c_len)

	return
}

// Blacklisted : g_unix_error_quark

// Unsupported : g_unix_fd_add : unsupported parameter function : no type generator for UnixFDSourceFunc (GUnixFDSourceFunc) for param function

// Unsupported : g_unix_fd_add_full : unsupported parameter function : no type generator for UnixFDSourceFunc (GUnixFDSourceFunc) for param function

// Unsupported : g_unix_signal_add : unsupported parameter handler : no type generator for SourceFunc (GSourceFunc) for param handler

// Unsupported : g_unix_signal_add_full : unsupported parameter handler : no type generator for SourceFunc (GSourceFunc) for param handler

// Unsupported : g_uri_list_extract_uris : no return type

// Usleep is a wrapper around the C function g_usleep.
func Usleep(microseconds uint64) {
	c_microseconds := (C.gulong)(microseconds)

	C.g_usleep(c_microseconds)

	return
}

// Unsupported : g_utf16_to_ucs4 : unsupported parameter str : no type generator for guint16 (const gunichar2*) for param str

// Unsupported : g_utf16_to_utf8 : unsupported parameter str : no type generator for guint16 (const gunichar2*) for param str

// Utf8Casefold is a wrapper around the C function g_utf8_casefold.
func Utf8Casefold(str string, len int64) string {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	c_len := (C.gssize)(len)

	retC := C.g_utf8_casefold(c_str, c_len)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Utf8Collate is a wrapper around the C function g_utf8_collate.
func Utf8Collate(str1 string, str2 string) int32 {
	c_str1 := C.CString(str1)
	defer C.free(unsafe.Pointer(c_str1))

	c_str2 := C.CString(str2)
	defer C.free(unsafe.Pointer(c_str2))

	retC := C.g_utf8_collate(c_str1, c_str2)
	retGo := (int32)(retC)

	return retGo
}

// Utf8CollateKey is a wrapper around the C function g_utf8_collate_key.
func Utf8CollateKey(str string, len int64) string {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	c_len := (C.gssize)(len)

	retC := C.g_utf8_collate_key(c_str, c_len)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Utf8FindNextChar is a wrapper around the C function g_utf8_find_next_char.
func Utf8FindNextChar(p string, end string) string {
	c_p := C.CString(p)
	defer C.free(unsafe.Pointer(c_p))

	c_end := C.CString(end)
	defer C.free(unsafe.Pointer(c_end))

	retC := C.g_utf8_find_next_char(c_p, c_end)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Utf8FindPrevChar is a wrapper around the C function g_utf8_find_prev_char.
func Utf8FindPrevChar(str string, p string) string {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	c_p := C.CString(p)
	defer C.free(unsafe.Pointer(c_p))

	retC := C.g_utf8_find_prev_char(c_str, c_p)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Utf8GetChar is a wrapper around the C function g_utf8_get_char.
func Utf8GetChar(p string) rune {
	c_p := C.CString(p)
	defer C.free(unsafe.Pointer(c_p))

	retC := C.g_utf8_get_char(c_p)
	retGo := (rune)(retC)

	return retGo
}

// Utf8GetCharValidated is a wrapper around the C function g_utf8_get_char_validated.
func Utf8GetCharValidated(p string, maxLen int64) rune {
	c_p := C.CString(p)
	defer C.free(unsafe.Pointer(c_p))

	c_max_len := (C.gssize)(maxLen)

	retC := C.g_utf8_get_char_validated(c_p, c_max_len)
	retGo := (rune)(retC)

	return retGo
}

// Utf8Normalize is a wrapper around the C function g_utf8_normalize.
func Utf8Normalize(str string, len int64, mode NormalizeMode) string {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	c_len := (C.gssize)(len)

	c_mode := (C.GNormalizeMode)(mode)

	retC := C.g_utf8_normalize(c_str, c_len, c_mode)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Utf8OffsetToPointer is a wrapper around the C function g_utf8_offset_to_pointer.
func Utf8OffsetToPointer(str string, offset int64) string {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	c_offset := (C.glong)(offset)

	retC := C.g_utf8_offset_to_pointer(c_str, c_offset)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Utf8PointerToOffset is a wrapper around the C function g_utf8_pointer_to_offset.
func Utf8PointerToOffset(str string, pos string) int64 {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	c_pos := C.CString(pos)
	defer C.free(unsafe.Pointer(c_pos))

	retC := C.g_utf8_pointer_to_offset(c_str, c_pos)
	retGo := (int64)(retC)

	return retGo
}

// Utf8PrevChar is a wrapper around the C function g_utf8_prev_char.
func Utf8PrevChar(p string) string {
	c_p := C.CString(p)
	defer C.free(unsafe.Pointer(c_p))

	retC := C.g_utf8_prev_char(c_p)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Utf8Strchr is a wrapper around the C function g_utf8_strchr.
func Utf8Strchr(p string, len int64, c rune) string {
	c_p := C.CString(p)
	defer C.free(unsafe.Pointer(c_p))

	c_len := (C.gssize)(len)

	c_c := (C.gunichar)(c)

	retC := C.g_utf8_strchr(c_p, c_len, c_c)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Utf8Strdown is a wrapper around the C function g_utf8_strdown.
func Utf8Strdown(str string, len int64) string {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	c_len := (C.gssize)(len)

	retC := C.g_utf8_strdown(c_str, c_len)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Utf8Strlen is a wrapper around the C function g_utf8_strlen.
func Utf8Strlen(p string, max int64) int64 {
	c_p := C.CString(p)
	defer C.free(unsafe.Pointer(c_p))

	c_max := (C.gssize)(max)

	retC := C.g_utf8_strlen(c_p, c_max)
	retGo := (int64)(retC)

	return retGo
}

// Utf8Strncpy is a wrapper around the C function g_utf8_strncpy.
func Utf8Strncpy(dest string, src string, n uint64) string {
	c_dest := C.CString(dest)
	defer C.free(unsafe.Pointer(c_dest))

	c_src := C.CString(src)
	defer C.free(unsafe.Pointer(c_src))

	c_n := (C.gsize)(n)

	retC := C.g_utf8_strncpy(c_dest, c_src, c_n)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Utf8Strrchr is a wrapper around the C function g_utf8_strrchr.
func Utf8Strrchr(p string, len int64, c rune) string {
	c_p := C.CString(p)
	defer C.free(unsafe.Pointer(c_p))

	c_len := (C.gssize)(len)

	c_c := (C.gunichar)(c)

	retC := C.g_utf8_strrchr(c_p, c_len, c_c)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Utf8Strup is a wrapper around the C function g_utf8_strup.
func Utf8Strup(str string, len int64) string {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	c_len := (C.gssize)(len)

	retC := C.g_utf8_strup(c_str, c_len)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Blacklisted : g_utf8_to_ucs4

// Blacklisted : g_utf8_to_ucs4_fast

// Unsupported : g_utf8_to_utf16 : no return generator

// Unsupported : g_utf8_validate : unsupported parameter str : no type generator for guint8 () for array param str

// Blacklisted : g_variant_get_gtype

// Unsupported : g_variant_parse : unsupported parameter type : Blacklisted record : GVariantType

// VariantParseErrorQuark is a wrapper around the C function g_variant_parse_error_quark.
func VariantParseErrorQuark() Quark {
	retC := C.g_variant_parse_error_quark()
	retGo := (Quark)(retC)

	return retGo
}

// VariantParserGetErrorQuark is a wrapper around the C function g_variant_parser_get_error_quark.
func VariantParserGetErrorQuark() Quark {
	retC := C.g_variant_parser_get_error_quark()
	retGo := (Quark)(retC)

	return retGo
}

// Unsupported : g_variant_type_checked_ : return type : Blacklisted record : GVariantType

// VariantTypeStringIsValid is a wrapper around the C function g_variant_type_string_is_valid.
func VariantTypeStringIsValid(typeString string) bool {
	c_type_string := C.CString(typeString)
	defer C.free(unsafe.Pointer(c_type_string))

	retC := C.g_variant_type_string_is_valid(c_type_string)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : g_vasprintf : unsupported parameter string : in string with indirection level of 2

// Unsupported : g_vfprintf : unsupported parameter file : no type generator for gpointer (FILE*) for param file

// Unsupported : g_vprintf : unsupported parameter args : no type generator for va_list (va_list) for param args

// Unsupported : g_vsnprintf : unsupported parameter args : no type generator for va_list (va_list) for param args

// Unsupported : g_vsprintf : unsupported parameter args : no type generator for va_list (va_list) for param args

// WarnMessage is a wrapper around the C function g_warn_message.
func WarnMessage(domain string, file string, line int32, func_ string, warnexpr string) {
	c_domain := C.CString(domain)
	defer C.free(unsafe.Pointer(c_domain))

	c_file := C.CString(file)
	defer C.free(unsafe.Pointer(c_file))

	c_line := (C.int)(line)

	c_func := C.CString(func_)
	defer C.free(unsafe.Pointer(c_func))

	c_warnexpr := C.CString(warnexpr)
	defer C.free(unsafe.Pointer(c_warnexpr))

	C.g_warn_message(c_domain, c_file, c_line, c_func, c_warnexpr)

	return
}
