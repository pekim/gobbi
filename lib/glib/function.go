// This is a generated file - DO NOT EDIT

package glib

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

// Determines the numeric value of a character as a decimal digit.
// Differs from g_unichar_digit_value() because it takes a char, so
// there's no worry about sign extension if characters are signed.
/*

C function

g_ascii_digit_value
*/
func AsciiDigitValue(c rune) int32 {
	c_c := (C.gchar)(c)

	retC := C.g_ascii_digit_value(c_c)
	retGo := (int32)(retC)

	return retGo
}

// Converts a #gdouble to a string, using the '.' as
// decimal point.
//
// This function generates enough precision that converting
// the string back using g_ascii_strtod() gives the same machine-number
// (on machines with IEEE compatible 64bit doubles). It is
// guaranteed that the size of the resulting string will never
// be larger than @G_ASCII_DTOSTR_BUF_SIZE bytes, including the terminating
// nul character, which is always added.
/*

C function

g_ascii_dtostr
*/
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

// Converts a #gdouble to a string, using the '.' as
// decimal point. To format the number you pass in
// a printf()-style format string. Allowed conversion
// specifiers are 'e', 'E', 'f', 'F', 'g' and 'G'.
//
// The returned buffer is guaranteed to be nul-terminated.
//
// If you just want to want to serialize the value into a
// string, use g_ascii_dtostr().
/*

C function

g_ascii_formatd
*/
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

// Compare two strings, ignoring the case of ASCII characters.
//
// Unlike the BSD strcasecmp() function, this only recognizes standard
// ASCII letters and ignores the locale, treating all non-ASCII
// bytes as if they are not letters.
//
// This function should be used only on strings that are known to be
// in encodings where the bytes corresponding to ASCII letters always
// represent themselves. This includes UTF-8 and the ISO-8859-*
// charsets, but not for instance double-byte encodings like the
// Windows Codepage 932, where the trailing bytes of double-byte
// characters include all ASCII letters. If you compare two CP932
// strings using this function, you will get false matches.
//
// Both @s1 and @s2 must be non-%NULL.
/*

C function

g_ascii_strcasecmp
*/
func AsciiStrcasecmp(s1 string, s2 string) int32 {
	c_s1 := C.CString(s1)
	defer C.free(unsafe.Pointer(c_s1))

	c_s2 := C.CString(s2)
	defer C.free(unsafe.Pointer(c_s2))

	retC := C.g_ascii_strcasecmp(c_s1, c_s2)
	retGo := (int32)(retC)

	return retGo
}

// Converts all upper case ASCII letters to lower case ASCII letters.
/*

C function

g_ascii_strdown
*/
func AsciiStrdown(str string, len int64) string {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	c_len := (C.gssize)(len)

	retC := C.g_ascii_strdown(c_str, c_len)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Compare @s1 and @s2, ignoring the case of ASCII characters and any
// characters after the first @n in each string.
//
// Unlike the BSD strcasecmp() function, this only recognizes standard
// ASCII letters and ignores the locale, treating all non-ASCII
// characters as if they are not letters.
//
// The same warning as in g_ascii_strcasecmp() applies: Use this
// function only on strings known to be in encodings where bytes
// corresponding to ASCII letters always represent themselves.
/*

C function

g_ascii_strncasecmp
*/
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

// Converts a string to a #gdouble value.
//
// This function behaves like the standard strtod() function
// does in the C locale. It does this without actually changing
// the current locale, since that would not be thread-safe.
// A limitation of the implementation is that this function
// will still accept localized versions of infinities and NANs.
//
// This function is typically used when reading configuration
// files or other non-user input that should be locale independent.
// To handle input from the user you should normally use the
// locale-sensitive system strtod() function.
//
// To convert from a #gdouble to a string in a locale-insensitive
// way, use g_ascii_dtostr().
//
// If the correct value would cause overflow, plus or minus %HUGE_VAL
// is returned (according to the sign of the value), and %ERANGE is
// stored in %errno. If the correct value would cause underflow,
// zero is returned and %ERANGE is stored in %errno.
//
// This function resets %errno before calling strtod() so that
// you can reliably detect overflow and underflow.
/*

C function

g_ascii_strtod
*/
func AsciiStrtod(nptr string) (float64, string) {
	c_nptr := C.CString(nptr)
	defer C.free(unsafe.Pointer(c_nptr))

	var c_endptr *C.gchar

	retC := C.g_ascii_strtod(c_nptr, &c_endptr)
	retGo := (float64)(retC)

	endptr := C.GoString(c_endptr)

	return retGo, endptr
}

// Converts all lower case ASCII letters to upper case ASCII letters.
/*

C function

g_ascii_strup
*/
func AsciiStrup(str string, len int64) string {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	c_len := (C.gssize)(len)

	retC := C.g_ascii_strup(c_str, c_len)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Convert a character to ASCII lower case.
//
// Unlike the standard C library tolower() function, this only
// recognizes standard ASCII letters and ignores the locale, returning
// all non-ASCII characters unchanged, even if they are lower case
// letters in a particular character set. Also unlike the standard
// library function, this takes and returns a char, not an int, so
// don't call it on %EOF but no need to worry about casting to #guchar
// before passing a possibly non-ASCII character in.
/*

C function

g_ascii_tolower
*/
func AsciiTolower(c rune) rune {
	c_c := (C.gchar)(c)

	retC := C.g_ascii_tolower(c_c)
	retGo := (rune)(retC)

	return retGo
}

// Convert a character to ASCII upper case.
//
// Unlike the standard C library toupper() function, this only
// recognizes standard ASCII letters and ignores the locale, returning
// all non-ASCII characters unchanged, even if they are upper case
// letters in a particular character set. Also unlike the standard
// library function, this takes and returns a char, not an int, so
// don't call it on %EOF but no need to worry about casting to #guchar
// before passing a possibly non-ASCII character in.
/*

C function

g_ascii_toupper
*/
func AsciiToupper(c rune) rune {
	c_c := (C.gchar)(c)

	retC := C.g_ascii_toupper(c_c)
	retGo := (rune)(retC)

	return retGo
}

// Determines the numeric value of a character as a hexidecimal
// digit. Differs from g_unichar_xdigit_value() because it takes
// a char, so there's no worry about sign extension if characters
// are signed.
/*

C function

g_ascii_xdigit_value
*/
func AsciiXdigitValue(c rune) int32 {
	c_c := (C.gchar)(c)

	retC := C.g_ascii_xdigit_value(c_c)
	retGo := (int32)(retC)

	return retGo
}

/*

C function

g_assert_warning
*/
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

/*

C function

g_assertion_message
*/
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

/*

C function

g_assertion_message_cmpstr
*/
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

/*

C function

g_assertion_message_error
*/
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

	c_error := (*C.GError)(C.NULL)
	if error != nil {
		c_error = (*C.GError)(error.ToC())
	}

	c_error_domain := (C.GQuark)(errorDomain)

	c_error_code := (C.int)(errorCode)

	C.g_assertion_message_error(c_domain, c_file, c_line, c_func, c_expr, c_error, c_error_domain, c_error_code)

	return
}

/*

C function

g_assertion_message_expr
*/
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

// Gets the name of the file without any leading directory
// components. It returns a pointer into the given file name
// string.
/*

C function

g_basename
*/
func Basename(fileName string) string {
	c_file_name := C.CString(fileName)
	defer C.free(unsafe.Pointer(c_file_name))

	retC := C.g_basename(c_file_name)
	retGo := C.GoString(retC)

	return retGo
}

// Find the position of the first bit set in @mask, searching
// from (but not including) @nth_bit upwards. Bits are numbered
// from 0 (least significant) to sizeof(#gulong) * 8 - 1 (31 or 63,
// usually). To start searching from the 0th bit, set @nth_bit to -1.
/*

C function

g_bit_nth_lsf
*/
func BitNthLsf(mask uint64, nthBit int32) int32 {
	c_mask := (C.gulong)(mask)

	c_nth_bit := (C.gint)(nthBit)

	retC := C.g_bit_nth_lsf(c_mask, c_nth_bit)
	retGo := (int32)(retC)

	return retGo
}

// Find the position of the first bit set in @mask, searching
// from (but not including) @nth_bit downwards. Bits are numbered
// from 0 (least significant) to sizeof(#gulong) * 8 - 1 (31 or 63,
// usually). To start searching from the last bit, set @nth_bit to
// -1 or GLIB_SIZEOF_LONG * 8.
/*

C function

g_bit_nth_msf
*/
func BitNthMsf(mask uint64, nthBit int32) int32 {
	c_mask := (C.gulong)(mask)

	c_nth_bit := (C.gint)(nthBit)

	retC := C.g_bit_nth_msf(c_mask, c_nth_bit)
	retGo := (int32)(retC)

	return retGo
}

// Gets the number of bits used to hold @number,
// e.g. if @number is 4, 3 bits are needed.
/*

C function

g_bit_storage
*/
func BitStorage(number uint64) uint32 {
	c_number := (C.gulong)(number)

	retC := C.g_bit_storage(c_number)
	retGo := (uint32)(retC)

	return retGo
}

/*

C function

g_bookmark_file_error_quark
*/
func BookmarkFileErrorQuark() Quark {
	retC := C.g_bookmark_file_error_quark()
	retGo := (Quark)(retC)

	return retGo
}

// Unsupported : g_build_filename : unsupported parameter ... : varargs

// Unsupported : g_build_path : unsupported parameter ... : varargs

// Blacklisted : g_byte_array_free

// Unsupported : g_byte_array_new : no return type

// If @err or *@err is %NULL, does nothing. Otherwise,
// calls g_error_free() on *@err and sets *@err to %NULL.
/*

C function

g_clear_error
*/
func ClearError() error {
	var cThrowableError *C.GError

	C.g_clear_error(&cThrowableError)

	goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return goThrowableError
}

// Unsupported : g_convert : no return type

/*

C function

g_convert_error_quark
*/
func ConvertErrorQuark() Quark {
	retC := C.g_convert_error_quark()
	retGo := (Quark)(retC)

	return retGo
}

// Unsupported : g_convert_with_fallback : no return type

// Unsupported : g_convert_with_iconv : unsupported parameter converter : Blacklisted record : GIConv

// Frees all the data elements of the datalist.
// The data elements' destroy functions are called
// if they have been set.
/*

C function

g_datalist_clear
*/
func DatalistClear(datalist *Data) {
	c_datalist := (**C.GData)(C.NULL)
	if datalist != nil {
		c_datalist = (**C.GData)(datalist.ToC())
	}

	C.g_datalist_clear(c_datalist)

	return
}

// Unsupported : g_datalist_foreach : unsupported parameter func : no type generator for DataForeachFunc (GDataForeachFunc) for param func

// Gets a data element, using its string identifier. This is slower than
// g_datalist_id_get_data() because it compares strings.
/*

C function

g_datalist_get_data
*/
func DatalistGetData(datalist *Data, key string) uintptr {
	c_datalist := (**C.GData)(C.NULL)
	if datalist != nil {
		c_datalist = (**C.GData)(datalist.ToC())
	}

	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	retC := C.g_datalist_get_data(c_datalist, c_key)
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// Retrieves the data element corresponding to @key_id.
/*

C function

g_datalist_id_get_data
*/
func DatalistIdGetData(datalist *Data, keyId Quark) uintptr {
	c_datalist := (**C.GData)(C.NULL)
	if datalist != nil {
		c_datalist = (**C.GData)(datalist.ToC())
	}

	c_key_id := (C.GQuark)(keyId)

	retC := C.g_datalist_id_get_data(c_datalist, c_key_id)
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// Removes an element, without calling its destroy notification
// function.
/*

C function

g_datalist_id_remove_no_notify
*/
func DatalistIdRemoveNoNotify(datalist *Data, keyId Quark) uintptr {
	c_datalist := (**C.GData)(C.NULL)
	if datalist != nil {
		c_datalist = (**C.GData)(datalist.ToC())
	}

	c_key_id := (C.GQuark)(keyId)

	retC := C.g_datalist_id_remove_no_notify(c_datalist, c_key_id)
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_datalist_id_set_data_full : unsupported parameter destroy_func : no type generator for DestroyNotify (GDestroyNotify) for param destroy_func

// Resets the datalist to %NULL. It does not free any memory or call
// any destroy functions.
/*

C function

g_datalist_init
*/
func DatalistInit(datalist *Data) {
	c_datalist := (**C.GData)(C.NULL)
	if datalist != nil {
		c_datalist = (**C.GData)(datalist.ToC())
	}

	C.g_datalist_init(c_datalist)

	return
}

// Destroys the dataset, freeing all memory allocated, and calling any
// destroy functions set for data elements.
/*

C function

g_dataset_destroy
*/
func DatasetDestroy(datasetLocation uintptr) {
	c_dataset_location := (C.gconstpointer)(datasetLocation)

	C.g_dataset_destroy(c_dataset_location)

	return
}

// Unsupported : g_dataset_foreach : unsupported parameter func : no type generator for DataForeachFunc (GDataForeachFunc) for param func

// Gets the data element corresponding to a #GQuark.
/*

C function

g_dataset_id_get_data
*/
func DatasetIdGetData(datasetLocation uintptr, keyId Quark) uintptr {
	c_dataset_location := (C.gconstpointer)(datasetLocation)

	c_key_id := (C.GQuark)(keyId)

	retC := C.g_dataset_id_get_data(c_dataset_location, c_key_id)
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// Removes an element, without calling its destroy notification
// function.
/*

C function

g_dataset_id_remove_no_notify
*/
func DatasetIdRemoveNoNotify(datasetLocation uintptr, keyId Quark) uintptr {
	c_dataset_location := (C.gconstpointer)(datasetLocation)

	c_key_id := (C.GQuark)(keyId)

	retC := C.g_dataset_id_remove_no_notify(c_dataset_location, c_key_id)
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_dataset_id_set_data_full : unsupported parameter destroy_func : no type generator for DestroyNotify (GDestroyNotify) for param destroy_func

// Returns the number of days in a month, taking leap
// years into account.
/*

C function

g_date_get_days_in_month
*/
func DateGetDaysInMonth(month DateMonth, year DateYear) uint8 {
	c_month := (C.GDateMonth)(month)

	c_year := (C.GDateYear)(year)

	retC := C.g_date_get_days_in_month(c_month, c_year)
	retGo := (uint8)(retC)

	return retGo
}

// Returns the number of weeks in the year, where weeks
// are taken to start on Monday. Will be 52 or 53. The
// date must be valid. (Years always have 52 7-day periods,
// plus 1 or 2 extra days depending on whether it's a leap
// year. This function is basically telling you how many
// Mondays are in the year, i.e. there are 53 Mondays if
// one of the extra days happens to be a Monday.)
/*

C function

g_date_get_monday_weeks_in_year
*/
func DateGetMondayWeeksInYear(year DateYear) uint8 {
	c_year := (C.GDateYear)(year)

	retC := C.g_date_get_monday_weeks_in_year(c_year)
	retGo := (uint8)(retC)

	return retGo
}

// Returns the number of weeks in the year, where weeks
// are taken to start on Sunday. Will be 52 or 53. The
// date must be valid. (Years always have 52 7-day periods,
// plus 1 or 2 extra days depending on whether it's a leap
// year. This function is basically telling you how many
// Sundays are in the year, i.e. there are 53 Sundays if
// one of the extra days happens to be a Sunday.)
/*

C function

g_date_get_sunday_weeks_in_year
*/
func DateGetSundayWeeksInYear(year DateYear) uint8 {
	c_year := (C.GDateYear)(year)

	retC := C.g_date_get_sunday_weeks_in_year(c_year)
	retGo := (uint8)(retC)

	return retGo
}

// Returns %TRUE if the year is a leap year.
//
// For the purposes of this function, leap year is every year
// divisible by 4 unless that year is divisible by 100. If it
// is divisible by 100 it would be a leap year only if that year
// is also divisible by 400.
/*

C function

g_date_is_leap_year
*/
func DateIsLeapYear(year DateYear) bool {
	c_year := (C.GDateYear)(year)

	retC := C.g_date_is_leap_year(c_year)
	retGo := retC == C.TRUE

	return retGo
}

// Generates a printed representation of the date, in a
// [locale][setlocale]-specific way.
// Works just like the platform's C library strftime() function,
// but only accepts date-related formats; time-related formats
// give undefined results. Date must be valid. Unlike strftime()
// (which uses the locale encoding), works on a UTF-8 format
// string and stores a UTF-8 result.
//
// This function does not provide any conversion specifiers in
// addition to those implemented by the platform's C library.
// For example, don't expect that using g_date_strftime() would
// make the \%F provided by the C99 strftime() work on Windows
// where the C library only complies to C89.
/*

C function

g_date_strftime
*/
func DateStrftime(s string, slen uint64, format string, date *Date) uint64 {
	c_s := C.CString(s)
	defer C.free(unsafe.Pointer(c_s))

	c_slen := (C.gsize)(slen)

	c_format := C.CString(format)
	defer C.free(unsafe.Pointer(c_format))

	c_date := (*C.GDate)(C.NULL)
	if date != nil {
		c_date = (*C.GDate)(date.ToC())
	}

	retC := C.g_date_strftime(c_s, c_slen, c_format, c_date)
	retGo := (uint64)(retC)

	return retGo
}

// Returns %TRUE if the day of the month is valid (a day is valid if it's
// between 1 and 31 inclusive).
/*

C function

g_date_valid_day
*/
func DateValidDay(day DateDay) bool {
	c_day := (C.GDateDay)(day)

	retC := C.g_date_valid_day(c_day)
	retGo := retC == C.TRUE

	return retGo
}

// Returns %TRUE if the day-month-year triplet forms a valid, existing day
// in the range of days #GDate understands (Year 1 or later, no more than
// a few thousand years in the future).
/*

C function

g_date_valid_dmy
*/
func DateValidDmy(day DateDay, month DateMonth, year DateYear) bool {
	c_day := (C.GDateDay)(day)

	c_month := (C.GDateMonth)(month)

	c_year := (C.GDateYear)(year)

	retC := C.g_date_valid_dmy(c_day, c_month, c_year)
	retGo := retC == C.TRUE

	return retGo
}

// Returns %TRUE if the Julian day is valid. Anything greater than zero
// is basically a valid Julian, though there is a 32-bit limit.
/*

C function

g_date_valid_julian
*/
func DateValidJulian(julianDate uint32) bool {
	c_julian_date := (C.guint32)(julianDate)

	retC := C.g_date_valid_julian(c_julian_date)
	retGo := retC == C.TRUE

	return retGo
}

// Returns %TRUE if the month value is valid. The 12 #GDateMonth
// enumeration values are the only valid months.
/*

C function

g_date_valid_month
*/
func DateValidMonth(month DateMonth) bool {
	c_month := (C.GDateMonth)(month)

	retC := C.g_date_valid_month(c_month)
	retGo := retC == C.TRUE

	return retGo
}

// Returns %TRUE if the weekday is valid. The seven #GDateWeekday enumeration
// values are the only valid weekdays.
/*

C function

g_date_valid_weekday
*/
func DateValidWeekday(weekday DateWeekday) bool {
	c_weekday := (C.GDateWeekday)(weekday)

	retC := C.g_date_valid_weekday(c_weekday)
	retGo := retC == C.TRUE

	return retGo
}

// Returns %TRUE if the year is valid. Any year greater than 0 is valid,
// though there is a 16-bit limit to what #GDate will understand.
/*

C function

g_date_valid_year
*/
func DateValidYear(year DateYear) bool {
	c_year := (C.GDateYear)(year)

	retC := C.g_date_valid_year(c_year)
	retGo := retC == C.TRUE

	return retGo
}

// Compares two #gpointer arguments and returns %TRUE if they are equal.
// It can be passed to g_hash_table_new() as the @key_equal_func
// parameter, when using opaque pointers compared by pointer value as
// keys in a #GHashTable.
//
// This equality function is also appropriate for keys that are integers
// stored in pointers, such as `GINT_TO_POINTER (n)`.
/*

C function

g_direct_equal
*/
func DirectEqual(v1 uintptr, v2 uintptr) bool {
	c_v1 := (C.gconstpointer)(v1)

	c_v2 := (C.gconstpointer)(v2)

	retC := C.g_direct_equal(c_v1, c_v2)
	retGo := retC == C.TRUE

	return retGo
}

// Converts a gpointer to a hash value.
// It can be passed to g_hash_table_new() as the @hash_func parameter,
// when using opaque pointers compared by pointer value as keys in a
// #GHashTable.
//
// This hash function is also appropriate for keys that are integers
// stored in pointers, such as `GINT_TO_POINTER (n)`.
/*

C function

g_direct_hash
*/
func DirectHash(v uintptr) uint32 {
	c_v := (C.gconstpointer)(v)

	retC := C.g_direct_hash(c_v)
	retGo := (uint32)(retC)

	return retGo
}

// Gets a #GFileError constant based on the passed-in @err_no.
// For example, if you pass in `EEXIST` this function returns
// #G_FILE_ERROR_EXIST. Unlike `errno` values, you can portably
// assume that all #GFileError values will exist.
//
// Normally a #GFileError value goes into a #GError returned
// from a function that manipulates files. So you would use
// g_file_error_from_errno() when constructing a #GError.
/*

C function

g_file_error_from_errno
*/
func FileErrorFromErrno(errNo int32) FileError {
	c_err_no := (C.gint)(errNo)

	retC := C.g_file_error_from_errno(c_err_no)
	retGo := (FileError)(retC)

	return retGo
}

/*

C function

g_file_error_quark
*/
func FileErrorQuark() Quark {
	retC := C.g_file_error_quark()
	retGo := (Quark)(retC)

	return retGo
}

// Unsupported : g_file_get_contents : unsupported parameter contents : output array param contents

// Opens a file for writing in the preferred directory for temporary
// files (as returned by g_get_tmp_dir()).
//
// @tmpl should be a string in the GLib file name encoding containing
// a sequence of six 'X' characters, as the parameter to g_mkstemp().
// However, unlike these functions, the template should only be a
// basename, no directory components are allowed. If template is
// %NULL, a default template is used.
//
// Note that in contrast to g_mkstemp() (and mkstemp()) @tmpl is not
// modified, and might thus be a read-only literal string.
//
// Upon success, and if @name_used is non-%NULL, the actual name used
// is returned in @name_used. This string should be freed with g_free()
// when not needed any longer. The returned name is in the GLib file
// name encoding.
/*

C function

g_file_open_tmp
*/
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

// Returns %TRUE if any of the tests in the bitfield @test are
// %TRUE. For example, `(G_FILE_TEST_EXISTS | G_FILE_TEST_IS_DIR)`
// will return %TRUE if the file exists; the check whether it's a
// directory doesn't matter since the existence test is %TRUE. With
// the current set of available tests, there's no point passing in
// more than one test at a time.
//
// Apart from %G_FILE_TEST_IS_SYMLINK all tests follow symbolic links,
// so for a symbolic link to a regular file g_file_test() will return
// %TRUE for both %G_FILE_TEST_IS_SYMLINK and %G_FILE_TEST_IS_REGULAR.
//
// Note, that for a dangling symbolic link g_file_test() will return
// %TRUE for %G_FILE_TEST_IS_SYMLINK and %FALSE for all other flags.
//
// You should never use g_file_test() to test whether it is safe
// to perform an operation, because there is always the possibility
// of the condition changing before you actually perform the operation.
// For example, you might think you could use %G_FILE_TEST_IS_SYMLINK
// to know whether it is safe to write to a file without being
// tricked into writing into a different location. It doesn't work!
// |[<!-- language="C" -->
// DON'T DO THIS
// if (!g_file_test (filename, G_FILE_TEST_IS_SYMLINK))
// {
// fd = g_open (filename, O_WRONLY);
// write to fd
// }
// ]|
//
// Another thing to note is that %G_FILE_TEST_EXISTS and
// %G_FILE_TEST_IS_EXECUTABLE are implemented using the access()
// system call. This usually doesn't matter, but if your program
// is setuid or setgid it means that these tests will give you
// the answer for the real user ID and group ID, rather than the
// effective user ID and group ID.
//
// On Windows, there are no symlinks, so testing for
// %G_FILE_TEST_IS_SYMLINK will always return %FALSE. Testing for
// %G_FILE_TEST_IS_EXECUTABLE will just check that the file exists and
// its name indicates that it is executable, checking for well-known
// extensions and those listed in the `PATHEXT` environment variable.
/*

C function

g_file_test
*/
func FileTest(filename string, test GFileTest) bool {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))

	c_test := (C.GFileTest)(test)

	retC := C.g_file_test(c_filename, c_test)
	retGo := retC == C.TRUE

	return retGo
}

// Converts an escaped ASCII-encoded URI to a local filename in the
// encoding used for filenames.
/*

C function

g_filename_from_uri
*/
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

// Converts a string from UTF-8 to the encoding GLib uses for
// filenames. Note that on Windows GLib uses UTF-8 for filenames;
// on other platforms, this function indirectly depends on the
// [current locale][setlocale].
//
// The input string shall not contain nul characters even if the @len
// argument is positive. A nul character found inside the string will result
// in error %G_CONVERT_ERROR_ILLEGAL_SEQUENCE. If the filename encoding is
// not UTF-8 and the conversion output contains a nul character, the error
// %G_CONVERT_ERROR_EMBEDDED_NUL is set and the function returns %NULL.
/*

C function

g_filename_from_utf8
*/
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

// Converts an absolute filename to an escaped ASCII-encoded URI, with the path
// component following Section 3.3. of RFC 2396.
/*

C function

g_filename_to_uri
*/
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

// Converts a string which is in the encoding used by GLib for
// filenames into a UTF-8 string. Note that on Windows GLib uses UTF-8
// for filenames; on other platforms, this function indirectly depends on
// the [current locale][setlocale].
//
// The input string shall not contain nul characters even if the @len
// argument is positive. A nul character found inside the string will result
// in error %G_CONVERT_ERROR_ILLEGAL_SEQUENCE.
// If the source encoding is not UTF-8 and the conversion output contains a
// nul character, the error %G_CONVERT_ERROR_EMBEDDED_NUL is set and the
// function returns %NULL. Use g_convert() to produce output that
// may contain embedded nul characters.
/*

C function

g_filename_to_utf8
*/
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

// Locates the first executable named @program in the user's path, in the
// same way that execvp() would locate it. Returns an allocated string
// with the absolute path name, or %NULL if the program is not found in
// the path. If @program is already an absolute path, returns a copy of
// @program if @program exists and is executable, and %NULL otherwise.
//
// On Windows, if @program does not have a file type suffix, tries
// with the suffixes .exe, .cmd, .bat and .com, and the suffixes in
// the `PATHEXT` environment variable.
//
// On Windows, it looks for the file in the same way as CreateProcess()
// would. This means first in the directory where the executing
// program was loaded from, then in the current directory, then in the
// Windows 32-bit system directory, then in the Windows directory, and
// finally in the directories in the `PATH` environment variable. If
// the program is found, the return value contains the full name
// including the type suffix.
/*

C function

g_find_program_in_path
*/
func FindProgramInPath(program string) string {
	c_program := C.CString(program)
	defer C.free(unsafe.Pointer(c_program))

	retC := C.g_find_program_in_path(c_program)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Frees the memory pointed to by @mem.
//
// If @mem is %NULL it simply returns, so there is no need to check @mem
// against %NULL before calling this function.
/*

C function

g_free
*/
func Free(mem uintptr) {
	c_mem := (C.gpointer)(mem)

	C.g_free(c_mem)

	return
}

// Obtains the character set for the [current locale][setlocale]; you
// might use this character set as an argument to g_convert(), to convert
// from the current locale's encoding to some other encoding. (Frequently
// g_locale_to_utf8() and g_locale_from_utf8() are nice shortcuts, though.)
//
// On Windows the character set returned by this function is the
// so-called system default ANSI code-page. That is the character set
// used by the "narrow" versions of C library and Win32 functions that
// handle file names. It might be different from the character set
// used by the C library's current locale.
//
// On Linux, the character set is found by consulting nl_langinfo() if
// available. If not, the environment variables `LC_ALL`, `LC_CTYPE`, `LANG`
// and `CHARSET` are queried in order.
//
// The return value is %TRUE if the locale's encoding is UTF-8, in that
// case you can perhaps avoid calling g_convert().
//
// The string returned in @charset is not allocated, and should not be
// freed.
/*

C function

g_get_charset
*/
func GetCharset() (bool, string) {
	var c_charset *C.char

	retC := C.g_get_charset(&c_charset)
	retGo := retC == C.TRUE

	charset := C.GoString(c_charset)

	return retGo, charset
}

// Gets the character set for the current locale.
/*

C function

g_get_codeset
*/
func GetCodeset() string {
	retC := C.g_get_codeset()
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Gets the current directory.
//
// The returned string should be freed when no longer needed.
// The encoding of the returned string is system defined.
// On Windows, it is always UTF-8.
//
// Since GLib 2.40, this function will return the value of the "PWD"
// environment variable if it is set and it happens to be the same as
// the current directory.  This can make a difference in the case that
// the current directory is the target of a symbolic link.
/*

C function

g_get_current_dir
*/
func GetCurrentDir() string {
	retC := C.g_get_current_dir()
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Equivalent to the UNIX gettimeofday() function, but portable.
//
// You may find g_get_real_time() to be more convenient.
/*

C function

g_get_current_time
*/
func GetCurrentTime(result *TimeVal) {
	c_result := (*C.GTimeVal)(C.NULL)
	if result != nil {
		c_result = (*C.GTimeVal)(result.ToC())
	}

	C.g_get_current_time(c_result)

	return
}

// Gets the current user's home directory.
//
// As with most UNIX tools, this function will return the value of the
// `HOME` environment variable if it is set to an existing absolute path
// name, falling back to the `passwd` file in the case that it is unset.
//
// If the path given in `HOME` is non-absolute, does not exist, or is
// not a directory, the result is undefined.
//
// Before version 2.36 this function would ignore the `HOME` environment
// variable, taking the value from the `passwd` database instead. This was
// changed to increase the compatibility of GLib with other programs (and
// the XDG basedir specification) and to increase testability of programs
// based on GLib (by making it easier to run them from test frameworks).
//
// If your program has a strong requirement for either the new or the
// old behaviour (and if you don't wish to increase your GLib
// dependency to ensure that the new behaviour is in effect) then you
// should either directly check the `HOME` environment variable yourself
// or unset it before calling any functions in GLib.
/*

C function

g_get_home_dir
*/
func GetHomeDir() string {
	retC := C.g_get_home_dir()
	retGo := C.GoString(retC)

	return retGo
}

// Gets the name of the program. This name should not be localized,
// in contrast to g_get_application_name().
//
// If you are using #GApplication the program name is set in
// g_application_run(). In case of GDK or GTK+ it is set in
// gdk_init(), which is called by gtk_init() and the
// #GtkApplication::startup handler. The program name is found by
// taking the last component of @argv[0].
/*

C function

g_get_prgname
*/
func GetPrgname() string {
	retC := C.g_get_prgname()
	retGo := C.GoString(retC)

	return retGo
}

// Gets the real name of the user. This usually comes from the user's
// entry in the `passwd` file. The encoding of the returned string is
// system-defined. (On Windows, it is, however, always UTF-8.) If the
// real user name cannot be determined, the string "Unknown" is
// returned.
/*

C function

g_get_real_name
*/
func GetRealName() string {
	retC := C.g_get_real_name()
	retGo := C.GoString(retC)

	return retGo
}

// Gets the directory to use for temporary files.
//
// On UNIX, this is taken from the `TMPDIR` environment variable.
// If the variable is not set, `P_tmpdir` is
// used, as defined by the system C library. Failing that, a
// hard-coded default of "/tmp" is returned.
//
// On Windows, the `TEMP` environment variable is used, with the
// root directory of the Windows installation (eg: "C:\") used
// as a default.
//
// The encoding of the returned string is system-defined. On Windows,
// it is always UTF-8. The return value is never %NULL or the empty
// string.
/*

C function

g_get_tmp_dir
*/
func GetTmpDir() string {
	retC := C.g_get_tmp_dir()
	retGo := C.GoString(retC)

	return retGo
}

// Gets the user name of the current user. The encoding of the returned
// string is system-defined. On UNIX, it might be the preferred file name
// encoding, or something else, and there is no guarantee that it is even
// consistent on a machine. On Windows, it is always UTF-8.
/*

C function

g_get_user_name
*/
func GetUserName() string {
	retC := C.g_get_user_name()
	retGo := C.GoString(retC)

	return retGo
}

// Returns the value of an environment variable.
//
// On UNIX, the name and value are byte strings which might or might not
// be in some consistent character set and encoding. On Windows, they are
// in UTF-8.
// On Windows, in case the environment variable's value contains
// references to other environment variables, they are expanded.
/*

C function

g_getenv
*/
func Getenv(variable string) string {
	c_variable := C.CString(variable)
	defer C.free(unsafe.Pointer(c_variable))

	retC := C.g_getenv(c_variable)
	retGo := C.GoString(retC)

	return retGo
}

// Destroys all keys and values in the #GHashTable and decrements its
// reference count by 1. If keys and/or values are dynamically allocated,
// you should either free them first or create the #GHashTable with destroy
// notifiers using g_hash_table_new_full(). In the latter case the destroy
// functions you supplied will be called on all keys and values during the
// destruction phase.
/*

C function

g_hash_table_destroy
*/
func HashTableDestroy(hashTable *HashTable) {
	c_hash_table := (*C.GHashTable)(C.NULL)
	if hashTable != nil {
		c_hash_table = (*C.GHashTable)(hashTable.ToC())
	}

	C.g_hash_table_destroy(c_hash_table)

	return
}

// Inserts a new key and value into a #GHashTable.
//
// If the key already exists in the #GHashTable its current
// value is replaced with the new value. If you supplied a
// @value_destroy_func when creating the #GHashTable, the old
// value is freed using that function. If you supplied a
// @key_destroy_func when creating the #GHashTable, the passed
// key is freed using that function.
//
// Starting from GLib 2.40, this function returns a boolean value to
// indicate whether the newly added value was already in the hash table
// or not.
/*

C function

g_hash_table_insert
*/
func HashTableInsert(hashTable *HashTable, key uintptr, value uintptr) bool {
	c_hash_table := (*C.GHashTable)(C.NULL)
	if hashTable != nil {
		c_hash_table = (*C.GHashTable)(hashTable.ToC())
	}

	c_key := (C.gpointer)(key)

	c_value := (C.gpointer)(value)

	retC := C.g_hash_table_insert(c_hash_table, c_key, c_value)
	retGo := retC == C.TRUE

	return retGo
}

// Looks up a key in a #GHashTable. Note that this function cannot
// distinguish between a key that is not present and one which is present
// and has the value %NULL. If you need this distinction, use
// g_hash_table_lookup_extended().
/*

C function

g_hash_table_lookup
*/
func HashTableLookup(hashTable *HashTable, key uintptr) uintptr {
	c_hash_table := (*C.GHashTable)(C.NULL)
	if hashTable != nil {
		c_hash_table = (*C.GHashTable)(hashTable.ToC())
	}

	c_key := (C.gconstpointer)(key)

	retC := C.g_hash_table_lookup(c_hash_table, c_key)
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// Looks up a key in the #GHashTable, returning the original key and the
// associated value and a #gboolean which is %TRUE if the key was found. This
// is useful if you need to free the memory allocated for the original key,
// for example before calling g_hash_table_remove().
//
// You can actually pass %NULL for @lookup_key to test
// whether the %NULL key exists, provided the hash and equal functions
// of @hash_table are %NULL-safe.
/*

C function

g_hash_table_lookup_extended
*/
func HashTableLookupExtended(hashTable *HashTable, lookupKey uintptr) (bool, uintptr, uintptr) {
	c_hash_table := (*C.GHashTable)(C.NULL)
	if hashTable != nil {
		c_hash_table = (*C.GHashTable)(hashTable.ToC())
	}

	c_lookup_key := (C.gconstpointer)(lookupKey)

	var c_orig_key C.gpointer

	var c_value C.gpointer

	retC := C.g_hash_table_lookup_extended(c_hash_table, c_lookup_key, &c_orig_key, &c_value)
	retGo := retC == C.TRUE

	origKey := (uintptr)(unsafe.Pointer(&c_orig_key))

	value := (uintptr)(unsafe.Pointer(&c_value))

	return retGo, origKey, value
}

// Removes a key and its associated value from a #GHashTable.
//
// If the #GHashTable was created using g_hash_table_new_full(), the
// key and value are freed using the supplied destroy functions, otherwise
// you have to make sure that any dynamically allocated values are freed
// yourself.
/*

C function

g_hash_table_remove
*/
func HashTableRemove(hashTable *HashTable, key uintptr) bool {
	c_hash_table := (*C.GHashTable)(C.NULL)
	if hashTable != nil {
		c_hash_table = (*C.GHashTable)(hashTable.ToC())
	}

	c_key := (C.gconstpointer)(key)

	retC := C.g_hash_table_remove(c_hash_table, c_key)
	retGo := retC == C.TRUE

	return retGo
}

// Inserts a new key and value into a #GHashTable similar to
// g_hash_table_insert(). The difference is that if the key
// already exists in the #GHashTable, it gets replaced by the
// new key. If you supplied a @value_destroy_func when creating
// the #GHashTable, the old value is freed using that function.
// If you supplied a @key_destroy_func when creating the
// #GHashTable, the old key is freed using that function.
//
// Starting from GLib 2.40, this function returns a boolean value to
// indicate whether the newly added value was already in the hash table
// or not.
/*

C function

g_hash_table_replace
*/
func HashTableReplace(hashTable *HashTable, key uintptr, value uintptr) bool {
	c_hash_table := (*C.GHashTable)(C.NULL)
	if hashTable != nil {
		c_hash_table = (*C.GHashTable)(hashTable.ToC())
	}

	c_key := (C.gpointer)(key)

	c_value := (C.gpointer)(value)

	retC := C.g_hash_table_replace(c_hash_table, c_key, c_value)
	retGo := retC == C.TRUE

	return retGo
}

// Returns the number of elements contained in the #GHashTable.
/*

C function

g_hash_table_size
*/
func HashTableSize(hashTable *HashTable) uint32 {
	c_hash_table := (*C.GHashTable)(C.NULL)
	if hashTable != nil {
		c_hash_table = (*C.GHashTable)(hashTable.ToC())
	}

	retC := C.g_hash_table_size(c_hash_table)
	retGo := (uint32)(retC)

	return retGo
}

// Removes a key and its associated value from a #GHashTable without
// calling the key and value destroy functions.
/*

C function

g_hash_table_steal
*/
func HashTableSteal(hashTable *HashTable, key uintptr) bool {
	c_hash_table := (*C.GHashTable)(C.NULL)
	if hashTable != nil {
		c_hash_table = (*C.GHashTable)(hashTable.ToC())
	}

	c_key := (C.gconstpointer)(key)

	retC := C.g_hash_table_steal(c_hash_table, c_key)
	retGo := retC == C.TRUE

	return retGo
}

// Destroys a #GHook, given its ID.
/*

C function

g_hook_destroy
*/
func HookDestroy(hookList *HookList, hookId uint64) bool {
	c_hook_list := (*C.GHookList)(C.NULL)
	if hookList != nil {
		c_hook_list = (*C.GHookList)(hookList.ToC())
	}

	c_hook_id := (C.gulong)(hookId)

	retC := C.g_hook_destroy(c_hook_list, c_hook_id)
	retGo := retC == C.TRUE

	return retGo
}

// Removes one #GHook from a #GHookList, marking it
// inactive and calling g_hook_unref() on it.
/*

C function

g_hook_destroy_link
*/
func HookDestroyLink(hookList *HookList, hook *Hook) {
	c_hook_list := (*C.GHookList)(C.NULL)
	if hookList != nil {
		c_hook_list = (*C.GHookList)(hookList.ToC())
	}

	c_hook := (*C.GHook)(C.NULL)
	if hook != nil {
		c_hook = (*C.GHook)(hook.ToC())
	}

	C.g_hook_destroy_link(c_hook_list, c_hook)

	return
}

// Calls the #GHookList @finalize_hook function if it exists,
// and frees the memory allocated for the #GHook.
/*

C function

g_hook_free
*/
func HookFree(hookList *HookList, hook *Hook) {
	c_hook_list := (*C.GHookList)(C.NULL)
	if hookList != nil {
		c_hook_list = (*C.GHookList)(hookList.ToC())
	}

	c_hook := (*C.GHook)(C.NULL)
	if hook != nil {
		c_hook = (*C.GHook)(hook.ToC())
	}

	C.g_hook_free(c_hook_list, c_hook)

	return
}

// Inserts a #GHook into a #GHookList, before a given #GHook.
/*

C function

g_hook_insert_before
*/
func HookInsertBefore(hookList *HookList, sibling *Hook, hook *Hook) {
	c_hook_list := (*C.GHookList)(C.NULL)
	if hookList != nil {
		c_hook_list = (*C.GHookList)(hookList.ToC())
	}

	c_sibling := (*C.GHook)(C.NULL)
	if sibling != nil {
		c_sibling = (*C.GHook)(sibling.ToC())
	}

	c_hook := (*C.GHook)(C.NULL)
	if hook != nil {
		c_hook = (*C.GHook)(hook.ToC())
	}

	C.g_hook_insert_before(c_hook_list, c_sibling, c_hook)

	return
}

// Prepends a #GHook on the start of a #GHookList.
/*

C function

g_hook_prepend
*/
func HookPrepend(hookList *HookList, hook *Hook) {
	c_hook_list := (*C.GHookList)(C.NULL)
	if hookList != nil {
		c_hook_list = (*C.GHookList)(hookList.ToC())
	}

	c_hook := (*C.GHook)(C.NULL)
	if hook != nil {
		c_hook = (*C.GHook)(hook.ToC())
	}

	C.g_hook_prepend(c_hook_list, c_hook)

	return
}

// Decrements the reference count of a #GHook.
// If the reference count falls to 0, the #GHook is removed
// from the #GHookList and g_hook_free() is called to free it.
/*

C function

g_hook_unref
*/
func HookUnref(hookList *HookList, hook *Hook) {
	c_hook_list := (*C.GHookList)(C.NULL)
	if hookList != nil {
		c_hook_list = (*C.GHookList)(hookList.ToC())
	}

	c_hook := (*C.GHook)(C.NULL)
	if hook != nil {
		c_hook = (*C.GHook)(hook.ToC())
	}

	C.g_hook_unref(c_hook_list, c_hook)

	return
}

// Unsupported : g_iconv : unsupported parameter converter : Blacklisted record : GIConv

// Unsupported : g_iconv_open : return type : Blacklisted record : GIConv

// Unsupported : g_idle_add : unsupported parameter function : no type generator for SourceFunc (GSourceFunc) for param function

// Unsupported : g_idle_add_full : unsupported parameter function : no type generator for SourceFunc (GSourceFunc) for param function

// Removes the idle function with the given data.
/*

C function

g_idle_remove_by_data
*/
func IdleRemoveByData(data uintptr) bool {
	c_data := (C.gpointer)(data)

	retC := C.g_idle_remove_by_data(c_data)
	retGo := retC == C.TRUE

	return retGo
}

// Creates a new idle source.
//
// The source will not initially be associated with any #GMainContext
// and must be added to one with g_source_attach() before it will be
// executed. Note that the default priority for idle sources is
// %G_PRIORITY_DEFAULT_IDLE, as compared to other sources which
// have a default priority of %G_PRIORITY_DEFAULT.
/*

C function

g_idle_source_new
*/
func IdleSourceNew() *Source {
	retC := C.g_idle_source_new()
	retGo := SourceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Compares the two #gint values being pointed to and returns
// %TRUE if they are equal.
// It can be passed to g_hash_table_new() as the @key_equal_func
// parameter, when using non-%NULL pointers to integers as keys in a
// #GHashTable.
//
// Note that this function acts on pointers to #gint, not on #gint
// directly: if your hash table's keys are of the form
// `GINT_TO_POINTER (n)`, use g_direct_equal() instead.
/*

C function

g_int_equal
*/
func IntEqual(v1 uintptr, v2 uintptr) bool {
	c_v1 := (C.gconstpointer)(v1)

	c_v2 := (C.gconstpointer)(v2)

	retC := C.g_int_equal(c_v1, c_v2)
	retGo := retC == C.TRUE

	return retGo
}

// Converts a pointer to a #gint to a hash value.
// It can be passed to g_hash_table_new() as the @hash_func parameter,
// when using non-%NULL pointers to integer values as keys in a #GHashTable.
//
// Note that this function acts on pointers to #gint, not on #gint
// directly: if your hash table's keys are of the form
// `GINT_TO_POINTER (n)`, use g_direct_hash() instead.
/*

C function

g_int_hash
*/
func IntHash(v uintptr) uint32 {
	c_v := (C.gconstpointer)(v)

	retC := C.g_int_hash(c_v)
	retGo := (uint32)(retC)

	return retGo
}

// Unsupported : g_io_add_watch : unsupported parameter channel : Blacklisted record : GIOChannel

// Unsupported : g_io_add_watch_full : unsupported parameter channel : Blacklisted record : GIOChannel

// Converts an `errno` error number to a #GIOChannelError.
/*

C function

g_io_channel_error_from_errno
*/
func IoChannelErrorFromErrno(en int32) IOChannelError {
	c_en := (C.gint)(en)

	retC := C.g_io_channel_error_from_errno(c_en)
	retGo := (IOChannelError)(retC)

	return retGo
}

/*

C function

g_io_channel_error_quark
*/
func IoChannelErrorQuark() Quark {
	retC := C.g_io_channel_error_quark()
	retGo := (Quark)(retC)

	return retGo
}

// Unsupported : g_io_create_watch : unsupported parameter channel : Blacklisted record : GIOChannel

/*

C function

g_key_file_error_quark
*/
func KeyFileErrorQuark() Quark {
	retC := C.g_key_file_error_quark()
	retGo := (Quark)(retC)

	return retGo
}

// Unsupported : g_locale_from_utf8 : no return type

// Converts a string which is in the encoding used for strings by
// the C runtime (usually the same as that used by the operating
// system) in the [current locale][setlocale] into a UTF-8 string.
//
// If the source encoding is not UTF-8 and the conversion output contains a
// nul character, the error %G_CONVERT_ERROR_EMBEDDED_NUL is set and the
// function returns %NULL.
// If the source encoding is UTF-8, an embedded nul character is treated with
// the %G_CONVERT_ERROR_ILLEGAL_SEQUENCE error for backward compatibility with
// earlier versions of this library. Use g_convert() to produce output that
// may contain embedded nul characters.
/*

C function

g_locale_to_utf8
*/
func LocaleToUtf8(opsysstring []uint8) (string, uint64, uint64, error) {
	c_opsysstring := &opsysstring[0]

	c_len := (C.gssize)(len(opsysstring))

	var c_bytes_read C.gsize

	var c_bytes_written C.gsize

	var cThrowableError *C.GError

	retC := C.g_locale_to_utf8((*C.gchar)(unsafe.Pointer(c_opsysstring)), c_len, &c_bytes_read, &c_bytes_written, &cThrowableError)
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

// Unsupported : g_log : unsupported parameter ... : varargs

// The default log handler set up by GLib; g_log_set_default_handler()
// allows to install an alternate default log handler.
// This is used if no log handler has been set for the particular log
// domain and log level combination. It outputs the message to stderr
// or stdout and if the log level is fatal it calls abort(). It automatically
// prints a new-line character after the message, so one does not need to be
// manually included in @message.
//
// The behavior of this log handler can be influenced by a number of
// environment variables:
//
// - `G_MESSAGES_PREFIXED`: A :-separated list of log levels for which
// messages should be prefixed by the program name and PID of the
// aplication.
//
// - `G_MESSAGES_DEBUG`: A space-separated list of log domains for
// which debug and informational messages are printed. By default
// these messages are not printed.
//
// stderr is used for levels %G_LOG_LEVEL_ERROR, %G_LOG_LEVEL_CRITICAL,
// %G_LOG_LEVEL_WARNING and %G_LOG_LEVEL_MESSAGE. stdout is used for
// the rest.
//
// This has no effect if structured logging is enabled; see
// [Using Structured Logging][using-structured-logging].
/*

C function

g_log_default_handler
*/
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

// Removes the log handler.
//
// This has no effect if structured logging is enabled; see
// [Using Structured Logging][using-structured-logging].
/*

C function

g_log_remove_handler
*/
func LogRemoveHandler(logDomain string, handlerId uint32) {
	c_log_domain := C.CString(logDomain)
	defer C.free(unsafe.Pointer(c_log_domain))

	c_handler_id := (C.guint)(handlerId)

	C.g_log_remove_handler(c_log_domain, c_handler_id)

	return
}

// Sets the message levels which are always fatal, in any log domain.
// When a message with any of these levels is logged the program terminates.
// You can only set the levels defined by GLib to be fatal.
// %G_LOG_LEVEL_ERROR is always fatal.
//
// You can also make some message levels fatal at runtime by setting
// the `G_DEBUG` environment variable (see
// [Running GLib Applications](glib-running.html)).
//
// Libraries should not call this function, as it affects all messages logged
// by a process, including those from other libraries.
//
// Structured log messages (using g_log_structured() and
// g_log_structured_array()) are fatal only if the default log writer is used;
// otherwise it is up to the writer function to determine which log messages
// are fatal. See [Using Structured Logging][using-structured-logging].
/*

C function

g_log_set_always_fatal
*/
func LogSetAlwaysFatal(fatalMask LogLevelFlags) LogLevelFlags {
	c_fatal_mask := (C.GLogLevelFlags)(fatalMask)

	retC := C.g_log_set_always_fatal(c_fatal_mask)
	retGo := (LogLevelFlags)(retC)

	return retGo
}

// Sets the log levels which are fatal in the given domain.
// %G_LOG_LEVEL_ERROR is always fatal.
//
// This has no effect on structured log messages (using g_log_structured() or
// g_log_structured_array()). To change the fatal behaviour for specific log
// messages, programs must install a custom log writer function using
// g_log_set_writer_func(). See
// [Using Structured Logging][using-structured-logging].
/*

C function

g_log_set_fatal_mask
*/
func LogSetFatalMask(logDomain string, fatalMask LogLevelFlags) LogLevelFlags {
	c_log_domain := C.CString(logDomain)
	defer C.free(unsafe.Pointer(c_log_domain))

	c_fatal_mask := (C.GLogLevelFlags)(fatalMask)

	retC := C.g_log_set_fatal_mask(c_log_domain, c_fatal_mask)
	retGo := (LogLevelFlags)(retC)

	return retGo
}

// Unsupported : g_log_set_handler : unsupported parameter log_func : no type generator for LogFunc (GLogFunc) for param log_func

// Unsupported : g_log_structured_standard : unsupported parameter ... : varargs

// Unsupported : g_logv : unsupported parameter args : no type generator for va_list (va_list) for param args

// Returns the global default main context. This is the main context
// used for main loop functions when a main loop is not explicitly
// specified, and corresponds to the "main" main loop. See also
// g_main_context_get_thread_default().
/*

C function

g_main_context_default
*/
func MainContextDefault() *MainContext {
	retC := C.g_main_context_default()
	retGo := MainContextNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Returns the depth of the stack of calls to
// g_main_context_dispatch() on any #GMainContext in the current thread.
// That is, when called from the toplevel, it gives 0. When
// called from within a callback from g_main_context_iteration()
// (or g_main_loop_run(), etc.) it returns 1. When called from within
// a callback to a recursive call to g_main_context_iteration(),
// it returns 2. And so forth.
//
// This function is useful in a situation like the following:
// Imagine an extremely simple "garbage collected" system.
//
// |[<!-- language="C" -->
// static GList *free_list;
//
// gpointer
// allocate_memory (gsize size)
// {
// gpointer result = g_malloc (size);
// free_list = g_list_prepend (free_list, result);
// return result;
// }
//
// void
// free_allocated_memory (void)
// {
// GList *l;
// for (l = free_list; l; l = l->next);
// g_free (l->data);
// g_list_free (free_list);
// free_list = NULL;
// }
//
// [...]
//
// while (TRUE);
// {
// g_main_context_iteration (NULL, TRUE);
// free_allocated_memory();
// }
// ]|
//
// This works from an application, however, if you want to do the same
// thing from a library, it gets more difficult, since you no longer
// control the main loop. You might think you can simply use an idle
// function to make the call to free_allocated_memory(), but that
// doesn't work, since the idle function could be called from a
// recursive callback. This can be fixed by using g_main_depth()
//
// |[<!-- language="C" -->
// gpointer
// allocate_memory (gsize size)
// {
// FreeListBlock *block = g_new (FreeListBlock, 1);
// block->mem = g_malloc (size);
// block->depth = g_main_depth ();
// free_list = g_list_prepend (free_list, block);
// return block->mem;
// }
//
// void
// free_allocated_memory (void)
// {
// GList *l;
//
// int depth = g_main_depth ();
// for (l = free_list; l; );
// {
// GList *next = l->next;
// FreeListBlock *block = l->data;
// if (block->depth > depth)
// {
// g_free (block->mem);
// g_free (block);
// free_list = g_list_delete_link (free_list, l);
// }
//
// l = next;
// }
// }
// ]|
//
// There is a temptation to use g_main_depth() to solve
// problems with reentrancy. For instance, while waiting for data
// to be received from the network in response to a menu item,
// the menu item might be selected again. It might seem that
// one could make the menu item's callback return immediately
// and do nothing if g_main_depth() returns a value greater than 1.
// However, this should be avoided since the user then sees selecting
// the menu item do nothing. Furthermore, you'll find yourself adding
// these checks all over your code, since there are doubtless many,
// many things that the user could do. Instead, you can use the
// following techniques:
//
// 1. Use gtk_widget_set_sensitive() or modal dialogs to prevent
// the user from interacting with elements while the main
// loop is recursing.
//
// 2. Avoid main loop recursion in situations where you can't handle
// arbitrary  callbacks. Instead, structure your code so that you
// simply return to the main loop and then get called again when
// there is more work to do.
/*

C function

g_main_depth
*/
func MainDepth() int32 {
	retC := C.g_main_depth()
	retGo := (int32)(retC)

	return retGo
}

// Allocates @n_bytes bytes of memory.
// If @n_bytes is 0 it returns %NULL.
/*

C function

g_malloc
*/
func Malloc(nBytes uint64) uintptr {
	c_n_bytes := (C.gsize)(nBytes)

	retC := C.g_malloc(c_n_bytes)
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// Allocates @n_bytes bytes of memory, initialized to 0's.
// If @n_bytes is 0 it returns %NULL.
/*

C function

g_malloc0
*/
func Malloc0(nBytes uint64) uintptr {
	c_n_bytes := (C.gsize)(nBytes)

	retC := C.g_malloc0(c_n_bytes)
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

/*

C function

g_markup_error_quark
*/
func MarkupErrorQuark() Quark {
	retC := C.g_markup_error_quark()
	retGo := (Quark)(retC)

	return retGo
}

// Escapes text so that the markup parser will parse it verbatim.
// Less than, greater than, ampersand, etc. are replaced with the
// corresponding entities. This function would typically be used
// when writing out a file to be parsed with the markup parser.
//
// Note that this function doesn't protect whitespace and line endings
// from being processed according to the XML rules for normalization
// of line endings and attribute values.
//
// Note also that this function will produce character references in
// the range of &#x1; ... &#x1f; for all control sequences
// except for tabstop, newline and carriage return.  The character
// references in this range are not valid XML 1.0, but they are
// valid XML 1.1 and will be accepted by the GMarkup parser.
/*

C function

g_markup_escape_text
*/
func MarkupEscapeText(text string) string {
	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))

	c_length := (C.gssize)(len(text))

	retC := C.g_markup_escape_text(c_text, c_length)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Checks whether the allocator used by g_malloc() is the system's
// malloc implementation. If it returns %TRUE memory allocated with
// malloc() can be used interchangeable with memory allocated using g_malloc().
// This function is useful for avoiding an extra copy of allocated memory returned
// by a non-GLib-based API.
/*

C function

g_mem_is_system_malloc
*/
func MemIsSystemMalloc() bool {
	retC := C.g_mem_is_system_malloc()
	retGo := retC == C.TRUE

	return retGo
}

// GLib used to support some tools for memory profiling, but this
// no longer works. There are many other useful tools for memory
// profiling these days which can be used instead.
/*

C function

g_mem_profile
*/
func MemProfile() {
	C.g_mem_profile()

	return
}

// This function used to let you override the memory allocation function.
// However, its use was incompatible with the use of global constructors
// in GLib and GIO, because those use the GLib allocators before main is
// reached. Therefore this function is now deprecated and is just a stub.
/*

C function

g_mem_set_vtable
*/
func MemSetVtable(vtable *MemVTable) {
	c_vtable := (*C.GMemVTable)(C.NULL)
	if vtable != nil {
		c_vtable = (*C.GMemVTable)(vtable.ToC())
	}

	C.g_mem_set_vtable(c_vtable)

	return
}

// Allocates @byte_size bytes of memory, and copies @byte_size bytes into it
// from @mem. If @mem is %NULL it returns %NULL.
/*

C function

g_memdup
*/
func Memdup(mem uintptr, byteSize uint32) uintptr {
	c_mem := (C.gconstpointer)(mem)

	c_byte_size := (C.guint)(byteSize)

	retC := C.g_memdup(c_mem, c_byte_size)
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// Opens a temporary file. See the mkstemp() documentation
// on most UNIX-like systems.
//
// The parameter is a string that should follow the rules for
// mkstemp() templates, i.e. contain the string "XXXXXX".
// g_mkstemp() is slightly more flexible than mkstemp() in that the
// sequence does not have to occur at the very end of the template.
// The X string will be modified to form the name of a file that
// didn't exist. The string should be in the GLib file name encoding.
// Most importantly, on Windows it should be in UTF-8.
/*

C function

g_mkstemp
*/
func Mkstemp(tmpl string) int32 {
	c_tmpl := C.CString(tmpl)
	defer C.free(unsafe.Pointer(c_tmpl))

	retC := C.g_mkstemp(c_tmpl)
	retGo := (int32)(retC)

	return retGo
}

// Set the pointer at the specified location to %NULL.
/*

C function

g_nullify_pointer
*/
func NullifyPointer(nullifyLocation uintptr) {
	c_nullify_location := (C.gpointer)(nullifyLocation)

	C.g_nullify_pointer(&c_nullify_location)

	return
}

// Blacklisted : g_number_parser_error_quark

// Prompts the user with
// `[E]xit, [H]alt, show [S]tack trace or [P]roceed`.
// This function is intended to be used for debugging use only.
// The following example shows how it can be used together with
// the g_log() functions.
//
// |[<!-- language="C" -->
// #include <glib.h>
//
// static void
// log_handler (const gchar   *log_domain,
// GLogLevelFlags log_level,
// const gchar   *message,
// gpointer       user_data)
// {
// g_log_default_handler (log_domain, log_level, message, user_data);
//
// g_on_error_query (MY_PROGRAM_NAME);
// }
//
// int
// main (int argc, char *argv[])
// {
// g_log_set_handler (MY_LOG_DOMAIN,
// G_LOG_LEVEL_WARNING |
// G_LOG_LEVEL_ERROR |
// G_LOG_LEVEL_CRITICAL,
// log_handler,
// NULL);
// ...
// ]|
//
// If "[E]xit" is selected, the application terminates with a call
// to _exit(0).
//
// If "[S]tack" trace is selected, g_on_error_stack_trace() is called.
// This invokes gdb, which attaches to the current process and shows
// a stack trace. The prompt is then shown again.
//
// If "[P]roceed" is selected, the function returns.
//
// This function may cause different actions on non-UNIX platforms.
/*

C function

g_on_error_query
*/
func OnErrorQuery(prgName string) {
	c_prg_name := C.CString(prgName)
	defer C.free(unsafe.Pointer(c_prg_name))

	C.g_on_error_query(c_prg_name)

	return
}

// Invokes gdb, which attaches to the current process and shows a
// stack trace. Called by g_on_error_query() when the "[S]tack trace"
// option is selected. You can get the current process's program name
// with g_get_prgname(), assuming that you have called gtk_init() or
// gdk_init().
//
// This function may cause different actions on non-UNIX platforms.
/*

C function

g_on_error_stack_trace
*/
func OnErrorStackTrace(prgName string) {
	c_prg_name := C.CString(prgName)
	defer C.free(unsafe.Pointer(c_prg_name))

	C.g_on_error_stack_trace(c_prg_name)

	return
}

/*

C function

g_option_error_quark
*/
func OptionErrorQuark() Quark {
	retC := C.g_option_error_quark()
	retGo := (Quark)(retC)

	return retGo
}

// Unsupported : g_parse_debug_string : unsupported parameter keys :

// Gets the last component of the filename.
//
// If @file_name ends with a directory separator it gets the component
// before the last slash. If @file_name consists only of directory
// separators (and on Windows, possibly a drive letter), a single
// separator is returned. If @file_name is empty, it gets ".".
/*

C function

g_path_get_basename
*/
func PathGetBasename(fileName string) string {
	c_file_name := C.CString(fileName)
	defer C.free(unsafe.Pointer(c_file_name))

	retC := C.g_path_get_basename(c_file_name)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Gets the directory components of a file name.
//
// If the file name has no directory components "." is returned.
// The returned string should be freed when no longer needed.
/*

C function

g_path_get_dirname
*/
func PathGetDirname(fileName string) string {
	c_file_name := C.CString(fileName)
	defer C.free(unsafe.Pointer(c_file_name))

	retC := C.g_path_get_dirname(c_file_name)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Returns %TRUE if the given @file_name is an absolute file name.
// Note that this is a somewhat vague concept on Windows.
//
// On POSIX systems, an absolute file name is well-defined. It always
// starts from the single root directory. For example "/usr/local".
//
// On Windows, the concepts of current drive and drive-specific
// current directory introduce vagueness. This function interprets as
// an absolute file name one that either begins with a directory
// separator such as "\Users\tml" or begins with the root on a drive,
// for example "C:\Windows". The first case also includes UNC paths
// such as "\\\\myserver\docs\foo". In all cases, either slashes or
// backslashes are accepted.
//
// Note that a file name relative to the current drive root does not
// truly specify a file uniquely over time and across processes, as
// the current drive is a per-process value and can be changed.
//
// File names relative the current directory on some specific drive,
// such as "D:foo/bar", are not interpreted as absolute by this
// function, but they obviously are not relative to the normal current
// directory as returned by getcwd() or g_get_current_dir()
// either. Such paths should be avoided, or need to be handled using
// Windows-specific code.
/*

C function

g_path_is_absolute
*/
func PathIsAbsolute(fileName string) bool {
	c_file_name := C.CString(fileName)
	defer C.free(unsafe.Pointer(c_file_name))

	retC := C.g_path_is_absolute(c_file_name)
	retGo := retC == C.TRUE

	return retGo
}

// Returns a pointer into @file_name after the root component,
// i.e. after the "/" in UNIX or "C:\" under Windows. If @file_name
// is not an absolute path it returns %NULL.
/*

C function

g_path_skip_root
*/
func PathSkipRoot(fileName string) string {
	c_file_name := C.CString(fileName)
	defer C.free(unsafe.Pointer(c_file_name))

	retC := C.g_path_skip_root(c_file_name)
	retGo := C.GoString(retC)

	return retGo
}

// Matches a string against a compiled pattern. Passing the correct
// length of the string given is mandatory. The reversed string can be
// omitted by passing %NULL, this is more efficient if the reversed
// version of the string to be matched is not at hand, as
// g_pattern_match() will only construct it if the compiled pattern
// requires reverse matches.
//
// Note that, if the user code will (possibly) match a string against a
// multitude of patterns containing wildcards, chances are high that
// some patterns will require a reversed string. In this case, it's
// more efficient to provide the reversed string to avoid multiple
// constructions thereof in the various calls to g_pattern_match().
//
// Note also that the reverse of a UTF-8 encoded string can in general
// not be obtained by g_strreverse(). This works only if the string
// does not contain any multibyte characters. GLib offers the
// g_utf8_strreverse() function to reverse UTF-8 encoded strings.
/*

C function

g_pattern_match
*/
func PatternMatch(pspec *PatternSpec, stringLength uint32, string string, stringReversed string) bool {
	c_pspec := (*C.GPatternSpec)(C.NULL)
	if pspec != nil {
		c_pspec = (*C.GPatternSpec)(pspec.ToC())
	}

	c_string_length := (C.guint)(stringLength)

	c_string := C.CString(string)
	defer C.free(unsafe.Pointer(c_string))

	c_string_reversed := C.CString(stringReversed)
	defer C.free(unsafe.Pointer(c_string_reversed))

	retC := C.g_pattern_match(c_pspec, c_string_length, c_string, c_string_reversed)
	retGo := retC == C.TRUE

	return retGo
}

// Matches a string against a pattern given as a string. If this
// function is to be called in a loop, it's more efficient to compile
// the pattern once with g_pattern_spec_new() and call
// g_pattern_match_string() repeatedly.
/*

C function

g_pattern_match_simple
*/
func PatternMatchSimple(pattern string, string string) bool {
	c_pattern := C.CString(pattern)
	defer C.free(unsafe.Pointer(c_pattern))

	c_string := C.CString(string)
	defer C.free(unsafe.Pointer(c_string))

	retC := C.g_pattern_match_simple(c_pattern, c_string)
	retGo := retC == C.TRUE

	return retGo
}

// Matches a string against a compiled pattern. If the string is to be
// matched against more than one pattern, consider using
// g_pattern_match() instead while supplying the reversed string.
/*

C function

g_pattern_match_string
*/
func PatternMatchString(pspec *PatternSpec, string string) bool {
	c_pspec := (*C.GPatternSpec)(C.NULL)
	if pspec != nil {
		c_pspec = (*C.GPatternSpec)(pspec.ToC())
	}

	c_string := C.CString(string)
	defer C.free(unsafe.Pointer(c_string))

	retC := C.g_pattern_match_string(c_pspec, c_string)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : g_print : unsupported parameter ... : varargs

// Unsupported : g_printerr : unsupported parameter ... : varargs

// Unsupported : g_printf_string_upper_bound : unsupported parameter args : no type generator for va_list (va_list) for param args

// If @dest is %NULL, free @src; otherwise, moves @src into *@dest.
// The error variable @dest points to must be %NULL.
//
// @src must be non-%NULL.
//
// Note that @src is no longer valid after this call. If you want
// to keep using the same GError*, you need to set it to %NULL
// after calling this function on it.
/*

C function

g_propagate_error
*/
func PropagateError(src *Error) *Error {
	var c_dest *C.GError

	c_src := (*C.GError)(C.NULL)
	if src != nil {
		c_src = (*C.GError)(src.ToC())
	}

	C.g_propagate_error(&c_dest, c_src)

	dest := ErrorNewFromC(unsafe.Pointer(c_dest))

	return dest
}

// Unsupported : g_qsort_with_data : unsupported parameter compare_func : no type generator for CompareDataFunc (GCompareDataFunc) for param compare_func

// Gets the #GQuark identifying the given (static) string. If the
// string does not currently have an associated #GQuark, a new #GQuark
// is created, linked to the given string.
//
// Note that this function is identical to g_quark_from_string() except
// that if a new #GQuark is created the string itself is used rather
// than a copy. This saves memory, but can only be used if the string
// will continue to exist until the program terminates. It can be used
// with statically allocated strings in the main program, but not with
// statically allocated memory in dynamically loaded modules, if you
// expect to ever unload the module again (e.g. do not use this
// function in GTK+ theme engines).
/*

C function

g_quark_from_static_string
*/
func QuarkFromStaticString(string string) Quark {
	c_string := C.CString(string)
	defer C.free(unsafe.Pointer(c_string))

	retC := C.g_quark_from_static_string(c_string)
	retGo := (Quark)(retC)

	return retGo
}

// Gets the #GQuark identifying the given string. If the string does
// not currently have an associated #GQuark, a new #GQuark is created,
// using a copy of the string.
/*

C function

g_quark_from_string
*/
func QuarkFromString(string string) Quark {
	c_string := C.CString(string)
	defer C.free(unsafe.Pointer(c_string))

	retC := C.g_quark_from_string(c_string)
	retGo := (Quark)(retC)

	return retGo
}

// Gets the string associated with the given #GQuark.
/*

C function

g_quark_to_string
*/
func QuarkToString(quark Quark) string {
	c_quark := (C.GQuark)(quark)

	retC := C.g_quark_to_string(c_quark)
	retGo := C.GoString(retC)

	return retGo
}

// Gets the #GQuark associated with the given string, or 0 if string is
// %NULL or it has no associated #GQuark.
//
// If you want the GQuark to be created if it doesn't already exist,
// use g_quark_from_string() or g_quark_from_static_string().
/*

C function

g_quark_try_string
*/
func QuarkTryString(string string) Quark {
	c_string := C.CString(string)
	defer C.free(unsafe.Pointer(c_string))

	retC := C.g_quark_try_string(c_string)
	retGo := (Quark)(retC)

	return retGo
}

// Returns a random #gdouble equally distributed over the range [0..1).
/*

C function

g_random_double
*/
func RandomDouble() float64 {
	retC := C.g_random_double()
	retGo := (float64)(retC)

	return retGo
}

// Returns a random #gdouble equally distributed over the range
// [@begin..@end).
/*

C function

g_random_double_range
*/
func RandomDoubleRange(begin float64, end float64) float64 {
	c_begin := (C.gdouble)(begin)

	c_end := (C.gdouble)(end)

	retC := C.g_random_double_range(c_begin, c_end)
	retGo := (float64)(retC)

	return retGo
}

// Return a random #guint32 equally distributed over the range
// [0..2^32-1].
/*

C function

g_random_int
*/
func RandomInt() uint32 {
	retC := C.g_random_int()
	retGo := (uint32)(retC)

	return retGo
}

// Returns a random #gint32 equally distributed over the range
// [@begin..@end-1].
/*

C function

g_random_int_range
*/
func RandomIntRange(begin int32, end int32) int32 {
	c_begin := (C.gint32)(begin)

	c_end := (C.gint32)(end)

	retC := C.g_random_int_range(c_begin, c_end)
	retGo := (int32)(retC)

	return retGo
}

// Sets the seed for the global random number generator, which is used
// by the g_random_* functions, to @seed.
/*

C function

g_random_set_seed
*/
func RandomSetSeed(seed uint32) {
	c_seed := (C.guint32)(seed)

	C.g_random_set_seed(c_seed)

	return
}

// Reallocates the memory pointed to by @mem, so that it now has space for
// @n_bytes bytes of memory. It returns the new address of the memory, which may
// have been moved. @mem may be %NULL, in which case it's considered to
// have zero-length. @n_bytes may be 0, in which case %NULL will be returned
// and @mem will be freed unless it is %NULL.
/*

C function

g_realloc
*/
func Realloc(mem uintptr, nBytes uint64) uintptr {
	c_mem := (C.gpointer)(mem)

	c_n_bytes := (C.gsize)(nBytes)

	retC := C.g_realloc(c_mem, c_n_bytes)
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

/*

C function

g_regex_error_quark
*/
func RegexErrorQuark() Quark {
	retC := C.g_regex_error_quark()
	retGo := (Quark)(retC)

	return retGo
}

/*

C function

g_return_if_fail_warning
*/
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

// Unsupported : g_set_error : unsupported parameter ... : varargs

// Sets the name of the program. This name should not be localized,
// in contrast to g_set_application_name().
//
// If you are using #GApplication the program name is set in
// g_application_run(). In case of GDK or GTK+ it is set in
// gdk_init(), which is called by gtk_init() and the
// #GtkApplication::startup handler. The program name is found by
// taking the last component of @argv[0].
//
// Note that for thread-safety reasons this function can only be called once.
/*

C function

g_set_prgname
*/
func SetPrgname(prgname string) {
	c_prgname := C.CString(prgname)
	defer C.free(unsafe.Pointer(c_prgname))

	C.g_set_prgname(c_prgname)

	return
}

// Unsupported : g_set_print_handler : unsupported parameter func : no type generator for PrintFunc (GPrintFunc) for param func

// Unsupported : g_set_printerr_handler : unsupported parameter func : no type generator for PrintFunc (GPrintFunc) for param func

/*

C function

g_shell_error_quark
*/
func ShellErrorQuark() Quark {
	retC := C.g_shell_error_quark()
	retGo := (Quark)(retC)

	return retGo
}

// Unsupported : g_shell_parse_argv : unsupported parameter argcp : array length param argcp is pointer (gint*)

// Quotes a string so that the shell (/bin/sh) will interpret the
// quoted string to mean @unquoted_string. If you pass a filename to
// the shell, for example, you should first quote it with this
// function.  The return value must be freed with g_free(). The
// quoting style used is undefined (single or double quotes may be
// used).
/*

C function

g_shell_quote
*/
func ShellQuote(unquotedString string) string {
	c_unquoted_string := C.CString(unquotedString)
	defer C.free(unsafe.Pointer(c_unquoted_string))

	retC := C.g_shell_quote(c_unquoted_string)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Unquotes a string as the shell (/bin/sh) would. Only handles
// quotes; if a string contains file globs, arithmetic operators,
// variables, backticks, redirections, or other special-to-the-shell
// features, the result will be different from the result a real shell
// would produce (the variables, backticks, etc. will be passed
// through literally instead of being expanded). This function is
// guaranteed to succeed if applied to the result of
// g_shell_quote(). If it fails, it returns %NULL and sets the
// error. The @quoted_string need not actually contain quoted or
// escaped text; g_shell_unquote() simply goes through the string and
// unquotes/unescapes anything that the shell would. Both single and
// double quotes are handled, as are escapes including escaped
// newlines. The return value must be freed with g_free(). Possible
// errors are in the #G_SHELL_ERROR domain.
//
// Shell quoting rules are a bit strange. Single quotes preserve the
// literal string exactly. escape sequences are not allowed; not even
// \' - if you want a ' in the quoted text, you have to do something
// like 'foo'\''bar'.  Double quotes allow $, `, ", \, and newline to
// be escaped with backslash. Otherwise double quotes preserve things
// literally.
/*

C function

g_shell_unquote
*/
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

/*

C function

g_slice_get_config
*/
func SliceGetConfig(ckey SliceConfig) int64 {
	c_ckey := (C.GSliceConfig)(ckey)

	retC := C.g_slice_get_config(c_ckey)
	retGo := (int64)(retC)

	return retGo
}

// Blacklisted : g_slice_get_config_state

/*

C function

g_slice_set_config
*/
func SliceSetConfig(ckey SliceConfig, value int64) {
	c_ckey := (C.GSliceConfig)(ckey)

	c_value := (C.gint64)(value)

	C.g_slice_set_config(c_ckey, c_value)

	return
}

// Unsupported : g_snprintf : unsupported parameter ... : varargs

// Removes the source with the given ID from the default main context. You must
// use g_source_destroy() for sources added to a non-default main context.
//
// The ID of a #GSource is given by g_source_get_id(), or will be
// returned by the functions g_source_attach(), g_idle_add(),
// g_idle_add_full(), g_timeout_add(), g_timeout_add_full(),
// g_child_watch_add(), g_child_watch_add_full(), g_io_add_watch(), and
// g_io_add_watch_full().
//
// It is a programmer error to attempt to remove a non-existent source.
//
// More specifically: source IDs can be reissued after a source has been
// destroyed and therefore it is never valid to use this function with a
// source ID which may have already been removed.  An example is when
// scheduling an idle to run in another thread with g_idle_add(): the
// idle may already have run and been removed by the time this function
// is called on its (now invalid) source ID.  This source ID may have
// been reissued, leading to the operation being performed against the
// wrong source.
/*

C function

g_source_remove
*/
func SourceRemove(tag uint32) bool {
	c_tag := (C.guint)(tag)

	retC := C.g_source_remove(c_tag)
	retGo := retC == C.TRUE

	return retGo
}

// Removes a source from the default main loop context given the
// source functions and user data. If multiple sources exist with the
// same source functions and user data, only one will be destroyed.
/*

C function

g_source_remove_by_funcs_user_data
*/
func SourceRemoveByFuncsUserData(funcs *SourceFuncs, userData uintptr) bool {
	c_funcs := (*C.GSourceFuncs)(C.NULL)
	if funcs != nil {
		c_funcs = (*C.GSourceFuncs)(funcs.ToC())
	}

	c_user_data := (C.gpointer)(userData)

	retC := C.g_source_remove_by_funcs_user_data(c_funcs, c_user_data)
	retGo := retC == C.TRUE

	return retGo
}

// Removes a source from the default main loop context given the user
// data for the callback. If multiple sources exist with the same user
// data, only one will be destroyed.
/*

C function

g_source_remove_by_user_data
*/
func SourceRemoveByUserData(userData uintptr) bool {
	c_user_data := (C.gpointer)(userData)

	retC := C.g_source_remove_by_user_data(c_user_data)
	retGo := retC == C.TRUE

	return retGo
}

// Gets the smallest prime number from a built-in array of primes which
// is larger than @num. This is used within GLib to calculate the optimum
// size of a #GHashTable.
//
// The built-in array of primes ranges from 11 to 13845163 such that
// each prime is approximately 1.5-2 times the previous prime.
/*

C function

g_spaced_primes_closest
*/
func SpacedPrimesClosest(num uint32) uint32 {
	c_num := (C.guint)(num)

	retC := C.g_spaced_primes_closest(c_num)
	retGo := (uint32)(retC)

	return retGo
}

// Unsupported : g_spawn_async : unsupported parameter argv :

// Unsupported : g_spawn_async_with_pipes : unsupported parameter argv :

// On some platforms, notably Windows, the #GPid type represents a resource
// which must be closed to prevent resource leaking. g_spawn_close_pid()
// is provided for this purpose. It should be used on all platforms, even
// though it doesn't do anything under UNIX.
/*

C function

g_spawn_close_pid
*/
func SpawnClosePid(pid Pid) {
	c_pid := (C.GPid)(pid)

	C.g_spawn_close_pid(c_pid)

	return
}

// A simple version of g_spawn_async() that parses a command line with
// g_shell_parse_argv() and passes it to g_spawn_async(). Runs a
// command line in the background. Unlike g_spawn_async(), the
// %G_SPAWN_SEARCH_PATH flag is enabled, other flags are not. Note
// that %G_SPAWN_SEARCH_PATH can have security implications, so
// consider using g_spawn_async() directly if appropriate. Possible
// errors are those from g_shell_parse_argv() and g_spawn_async().
//
// The same concerns on Windows apply as for g_spawn_command_line_sync().
/*

C function

g_spawn_command_line_async
*/
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

// Unsupported : g_spawn_command_line_sync : unsupported parameter standard_output : output array param standard_output

/*

C function

g_spawn_error_quark
*/
func SpawnErrorQuark() Quark {
	retC := C.g_spawn_error_quark()
	retGo := (Quark)(retC)

	return retGo
}

/*

C function

g_spawn_exit_error_quark
*/
func SpawnExitErrorQuark() Quark {
	retC := C.g_spawn_exit_error_quark()
	retGo := (Quark)(retC)

	return retGo
}

// Unsupported : g_spawn_sync : unsupported parameter argv :

// Copies a nul-terminated string into the dest buffer, include the
// trailing nul, and return a pointer to the trailing nul byte.
// This is useful for concatenating multiple strings together
// without having to repeatedly scan for the end.
/*

C function

g_stpcpy
*/
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

// Compares two strings for byte-by-byte equality and returns %TRUE
// if they are equal. It can be passed to g_hash_table_new() as the
// @key_equal_func parameter, when using non-%NULL strings as keys in a
// #GHashTable.
//
// Note that this function is primarily meant as a hash table comparison
// function. For a general-purpose, %NULL-safe string comparison function,
// see g_strcmp0().
/*

C function

g_str_equal
*/
func StrEqual(v1 uintptr, v2 uintptr) bool {
	c_v1 := (C.gconstpointer)(v1)

	c_v2 := (C.gconstpointer)(v2)

	retC := C.g_str_equal(c_v1, c_v2)
	retGo := retC == C.TRUE

	return retGo
}

// Converts a string to a hash value.
//
// This function implements the widely used "djb" hash apparently
// posted by Daniel Bernstein to comp.lang.c some time ago.  The 32
// bit unsigned hash value starts at 5381 and for each byte 'c' in
// the string, is updated: `hash = hash * 33 + c`. This function
// uses the signed value of each byte.
//
// It can be passed to g_hash_table_new() as the @hash_func parameter,
// when using non-%NULL strings as keys in a #GHashTable.
//
// Note that this function may not be a perfect fit for all use cases.
// For example, it produces some hash collisions with strings as short
// as 2.
/*

C function

g_str_hash
*/
func StrHash(v uintptr) uint32 {
	c_v := (C.gconstpointer)(v)

	retC := C.g_str_hash(c_v)
	retGo := (uint32)(retC)

	return retGo
}

// For each character in @string, if the character is not in @valid_chars,
// replaces the character with @substitutor. Modifies @string in place,
// and return @string itself, not a copy. The return value is to allow
// nesting such as
// |[<!-- language="C" -->
// g_ascii_strup (g_strcanon (str, "abc", '?'))
// ]|
/*

C function

g_strcanon
*/
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

// A case-insensitive string comparison, corresponding to the standard
// strcasecmp() function on platforms which support it.
/*

C function

g_strcasecmp
*/
func Strcasecmp(s1 string, s2 string) int32 {
	c_s1 := C.CString(s1)
	defer C.free(unsafe.Pointer(c_s1))

	c_s2 := C.CString(s2)
	defer C.free(unsafe.Pointer(c_s2))

	retC := C.g_strcasecmp(c_s1, c_s2)
	retGo := (int32)(retC)

	return retGo
}

// Removes trailing whitespace from a string.
//
// This function doesn't allocate or reallocate any memory;
// it modifies @string in place. Therefore, it cannot be used
// on statically allocated strings.
//
// The pointer to @string is returned to allow the nesting of functions.
//
// Also see g_strchug() and g_strstrip().
/*

C function

g_strchomp
*/
func Strchomp(string string) string {
	c_string := C.CString(string)
	defer C.free(unsafe.Pointer(c_string))

	retC := C.g_strchomp(c_string)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Removes leading whitespace from a string, by moving the rest
// of the characters forward.
//
// This function doesn't allocate or reallocate any memory;
// it modifies @string in place. Therefore, it cannot be used on
// statically allocated strings.
//
// The pointer to @string is returned to allow the nesting of functions.
//
// Also see g_strchomp() and g_strstrip().
/*

C function

g_strchug
*/
func Strchug(string string) string {
	c_string := C.CString(string)
	defer C.free(unsafe.Pointer(c_string))

	retC := C.g_strchug(c_string)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Replaces all escaped characters with their one byte equivalent.
//
// This function does the reverse conversion of g_strescape().
/*

C function

g_strcompress
*/
func Strcompress(source string) string {
	c_source := C.CString(source)
	defer C.free(unsafe.Pointer(c_source))

	retC := C.g_strcompress(c_source)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_strconcat : unsupported parameter ... : varargs

// Converts any delimiter characters in @string to @new_delimiter.
// Any characters in @string which are found in @delimiters are
// changed to the @new_delimiter character. Modifies @string in place,
// and returns @string itself, not a copy. The return value is to
// allow nesting such as
// |[<!-- language="C" -->
// g_ascii_strup (g_strdelimit (str, "abc", '?'))
// ]|
/*

C function

g_strdelimit
*/
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

// Converts a string to lower case.
/*

C function

g_strdown
*/
func Strdown(string string) string {
	c_string := C.CString(string)
	defer C.free(unsafe.Pointer(c_string))

	retC := C.g_strdown(c_string)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Duplicates a string. If @str is %NULL it returns %NULL.
// The returned string should be freed with g_free()
// when no longer needed.
/*

C function

g_strdup
*/
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

// Returns a string corresponding to the given error code, e.g. "no
// such process". Unlike strerror(), this always returns a string in
// UTF-8 encoding, and the pointer is guaranteed to remain valid for
// the lifetime of the process.
//
// Note that the string may be translated according to the current locale.
//
// The value of %errno will not be changed by this function. However, it may
// be changed by intermediate function calls, so you should save its value
// as soon as the call returns:
// |[
// int saved_errno;
//
// ret = read (blah);
// saved_errno = errno;
//
// g_strerror (saved_errno);
// ]|
/*

C function

g_strerror
*/
func Strerror(errnum int32) string {
	c_errnum := (C.gint)(errnum)

	retC := C.g_strerror(c_errnum)
	retGo := C.GoString(retC)

	return retGo
}

// Escapes the special characters '\b', '\f', '\n', '\r', '\t', '\v', '\'
// and '"' in the string @source by inserting a '\' before
// them. Additionally all characters in the range 0x01-0x1F (everything
// below SPACE) and in the range 0x7F-0xFF (all non-ASCII chars) are
// replaced with a '\' followed by their octal representation.
// Characters supplied in @exceptions are not escaped.
//
// g_strcompress() does the reverse conversion.
/*

C function

g_strescape
*/
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

// Creates a new #GString, initialized with the given string.
/*

C function

g_string_new
*/
func StringNew(init string) *String {
	c_init := C.CString(init)
	defer C.free(unsafe.Pointer(c_init))

	retC := C.g_string_new(c_init)
	retGo := StringNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Creates a new #GString with @len bytes of the @init buffer.
// Because a length is provided, @init need not be nul-terminated,
// and can contain embedded nul bytes.
//
// Since this function does not stop at nul bytes, it is the caller's
// responsibility to ensure that @init has at least @len addressable
// bytes.
/*

C function

g_string_new_len
*/
func StringNewLen(init string, len int64) *String {
	c_init := C.CString(init)
	defer C.free(unsafe.Pointer(c_init))

	c_len := (C.gssize)(len)

	retC := C.g_string_new_len(c_init, c_len)
	retGo := StringNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Creates a new #GString, with enough space for @dfl_size
// bytes. This is useful if you are going to add a lot of
// text to the string and don't want it to be reallocated
// too often.
/*

C function

g_string_sized_new
*/
func StringSizedNew(dflSize uint64) *String {
	c_dfl_size := (C.gsize)(dflSize)

	retC := C.g_string_sized_new(c_dfl_size)
	retGo := StringNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_strjoin : unsupported parameter ... : varargs

// Unsupported : g_strjoinv : unsupported parameter str_array : in string with indirection level of 2

// Portability wrapper that calls strlcat() on systems which have it,
// and emulates it otherwise. Appends nul-terminated @src string to @dest,
// guaranteeing nul-termination for @dest. The total size of @dest won't
// exceed @dest_size.
//
// At most @dest_size - 1 characters will be copied. Unlike strncat(),
// @dest_size is the full size of dest, not the space left over. This
// function does not allocate memory. It always nul-terminates (unless
// @dest_size == 0 or there were no nul characters in the @dest_size
// characters of dest to start with).
//
// Caveat: this is supposedly a more secure alternative to strcat() or
// strncat(), but for real security g_strconcat() is harder to mess up.
/*

C function

g_strlcat
*/
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

// Portability wrapper that calls strlcpy() on systems which have it,
// and emulates strlcpy() otherwise. Copies @src to @dest; @dest is
// guaranteed to be nul-terminated; @src must be nul-terminated;
// @dest_size is the buffer size, not the number of bytes to copy.
//
// At most @dest_size - 1 characters will be copied. Always nul-terminates
// (unless @dest_size is 0). This function does not allocate memory. Unlike
// strncpy(), this function doesn't pad @dest (so it's often faster). It
// returns the size of the attempted result, strlen (src), so if
// @retval >= @dest_size, truncation occurred.
//
// Caveat: strlcpy() is supposedly more secure than strcpy() or strncpy(),
// but if you really want to avoid screwups, g_strdup() is an even better
// idea.
/*

C function

g_strlcpy
*/
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

// A case-insensitive string comparison, corresponding to the standard
// strncasecmp() function on platforms which support it. It is similar
// to g_strcasecmp() except it only compares the first @n characters of
// the strings.
/*

C function

g_strncasecmp
*/
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

// Duplicates the first @n bytes of a string, returning a newly-allocated
// buffer @n + 1 bytes long which will always be nul-terminated. If @str
// is less than @n bytes long the buffer is padded with nuls. If @str is
// %NULL it returns %NULL. The returned value should be freed when no longer
// needed.
//
// To copy a number of characters from a UTF-8 encoded string,
// use g_utf8_strncpy() instead.
/*

C function

g_strndup
*/
func Strndup(str string, n uint64) string {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	c_n := (C.gsize)(n)

	retC := C.g_strndup(c_str, c_n)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Creates a new string @length bytes long filled with @fill_char.
// The returned string should be freed when no longer needed.
/*

C function

g_strnfill
*/
func Strnfill(length uint64, fillChar rune) string {
	c_length := (C.gsize)(length)

	c_fill_char := (C.gchar)(fillChar)

	retC := C.g_strnfill(c_length, c_fill_char)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Reverses all of the bytes in a string. For example,
// `g_strreverse ("abcdef")` will result in "fedcba".
//
// Note that g_strreverse() doesn't work on UTF-8 strings
// containing multibyte characters. For that purpose, use
// g_utf8_strreverse().
/*

C function

g_strreverse
*/
func Strreverse(string string) string {
	c_string := C.CString(string)
	defer C.free(unsafe.Pointer(c_string))

	retC := C.g_strreverse(c_string)
	retGo := C.GoString(retC)

	return retGo
}

// Searches the string @haystack for the last occurrence
// of the string @needle.
/*

C function

g_strrstr
*/
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

// Searches the string @haystack for the last occurrence
// of the string @needle, limiting the length of the search
// to @haystack_len.
/*

C function

g_strrstr_len
*/
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

// Returns a string describing the given signal, e.g. "Segmentation fault".
// You should use this function in preference to strsignal(), because it
// returns a string in UTF-8 encoding, and since not all platforms support
// the strsignal() function.
/*

C function

g_strsignal
*/
func Strsignal(signum int32) string {
	c_signum := (C.gint)(signum)

	retC := C.g_strsignal(c_signum)
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : g_strsplit : no return type

// Searches the string @haystack for the first occurrence
// of the string @needle, limiting the length of the search
// to @haystack_len.
/*

C function

g_strstr_len
*/
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

// Converts a string to a #gdouble value.
// It calls the standard strtod() function to handle the conversion, but
// if the string is not completely converted it attempts the conversion
// again with g_ascii_strtod(), and returns the best match.
//
// This function should seldom be used. The normal situation when reading
// numbers not for human consumption is to use g_ascii_strtod(). Only when
// you know that you must expect both locale formatted and C formatted numbers
// should you use this. Make sure that you don't pass strings such as comma
// separated lists of values, since the commas may be interpreted as a decimal
// point in some locales, causing unexpected results.
/*

C function

g_strtod
*/
func Strtod(nptr string) (float64, string) {
	c_nptr := C.CString(nptr)
	defer C.free(unsafe.Pointer(c_nptr))

	var c_endptr *C.gchar

	retC := C.g_strtod(c_nptr, &c_endptr)
	retGo := (float64)(retC)

	endptr := C.GoString(c_endptr)

	return retGo, endptr
}

// Converts a string to upper case.
/*

C function

g_strup
*/
func Strup(string string) string {
	c_string := C.CString(string)
	defer C.free(unsafe.Pointer(c_string))

	retC := C.g_strup(c_string)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Blacklisted : g_strv_get_type

// Unsupported : g_test_add_vtable : unsupported parameter data_setup : no type generator for TestFixtureFunc (GTestFixtureFunc) for param data_setup

/*

C function

g_test_assert_expected_messages_internal
*/
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

/*

C function

g_test_log_type_name
*/
func TestLogTypeName(logType TestLogType) string {
	c_log_type := (C.GTestLogType)(logType)

	retC := C.g_test_log_type_name(c_log_type)
	retGo := C.GoString(retC)

	return retGo
}

/*

C function

g_test_trap_assertions
*/
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

/*

C function

g_thread_error_quark
*/
func ThreadErrorQuark() Quark {
	retC := C.g_thread_error_quark()
	retGo := (Quark)(retC)

	return retGo
}

// Terminates the current thread.
//
// If another thread is waiting for us using g_thread_join() then the
// waiting thread will be woken up and get @retval as the return value
// of g_thread_join().
//
// Calling g_thread_exit() with a parameter @retval is equivalent to
// returning @retval from the function @func, as given to g_thread_new().
//
// You must only call g_thread_exit() from a thread that you created
// yourself with g_thread_new() or related APIs. You must not call
// this function from a thread created with another threading library
// or or from within a #GThreadPool.
/*

C function

g_thread_exit
*/
func ThreadExit(retval uintptr) {
	c_retval := (C.gpointer)(retval)

	C.g_thread_exit(c_retval)

	return
}

// Returns the maximal allowed number of unused threads.
/*

C function

g_thread_pool_get_max_unused_threads
*/
func ThreadPoolGetMaxUnusedThreads() int32 {
	retC := C.g_thread_pool_get_max_unused_threads()
	retGo := (int32)(retC)

	return retGo
}

// Returns the number of currently unused threads.
/*

C function

g_thread_pool_get_num_unused_threads
*/
func ThreadPoolGetNumUnusedThreads() uint32 {
	retC := C.g_thread_pool_get_num_unused_threads()
	retGo := (uint32)(retC)

	return retGo
}

// Sets the maximal number of unused threads to @max_threads.
// If @max_threads is -1, no limit is imposed on the number
// of unused threads.
//
// The default value is 2.
/*

C function

g_thread_pool_set_max_unused_threads
*/
func ThreadPoolSetMaxUnusedThreads(maxThreads int32) {
	c_max_threads := (C.gint)(maxThreads)

	C.g_thread_pool_set_max_unused_threads(c_max_threads)

	return
}

// Stops all currently unused threads. This does not change the
// maximal number of unused threads. This function can be used to
// regularly stop all unused threads e.g. from g_timeout_add().
/*

C function

g_thread_pool_stop_unused_threads
*/
func ThreadPoolStopUnusedThreads() {
	C.g_thread_pool_stop_unused_threads()

	return
}

// This function returns the #GThread corresponding to the
// current thread. Note that this function does not increase
// the reference count of the returned struct.
//
// This function will return a #GThread even for threads that
// were not created by GLib (i.e. those created by other threading
// APIs). This may be useful for thread identification purposes
// (i.e. comparisons) but you must not use GLib functions (such
// as g_thread_join()) on these threads.
/*

C function

g_thread_self
*/
func ThreadSelf() *Thread {
	retC := C.g_thread_self()
	retGo := ThreadNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Causes the calling thread to voluntarily relinquish the CPU, so
// that other threads can run.
//
// This function is often used as a method to make busy wait less evil.
/*

C function

g_thread_yield
*/
func ThreadYield() {
	C.g_thread_yield()

	return
}

// Unsupported : g_timeout_add : unsupported parameter function : no type generator for SourceFunc (GSourceFunc) for param function

// Unsupported : g_timeout_add_full : unsupported parameter function : no type generator for SourceFunc (GSourceFunc) for param function

// Creates a new timeout source.
//
// The source will not initially be associated with any #GMainContext
// and must be added to one with g_source_attach() before it will be
// executed.
//
// The interval given is in terms of monotonic time, not wall clock
// time.  See g_get_monotonic_time().
/*

C function

g_timeout_source_new
*/
func TimeoutSourceNew(interval uint32) *Source {
	c_interval := (C.guint)(interval)

	retC := C.g_timeout_source_new(c_interval)
	retGo := SourceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Returns the height of a #GTrashStack.
//
// Note that execution of this function is of O(N) complexity
// where N denotes the number of items on the stack.
/*

C function

g_trash_stack_height
*/
func TrashStackHeight(stackP *TrashStack) uint32 {
	c_stack_p := (**C.GTrashStack)(C.NULL)
	if stackP != nil {
		c_stack_p = (**C.GTrashStack)(stackP.ToC())
	}

	retC := C.g_trash_stack_height(c_stack_p)
	retGo := (uint32)(retC)

	return retGo
}

// Returns the element at the top of a #GTrashStack
// which may be %NULL.
/*

C function

g_trash_stack_peek
*/
func TrashStackPeek(stackP *TrashStack) uintptr {
	c_stack_p := (**C.GTrashStack)(C.NULL)
	if stackP != nil {
		c_stack_p = (**C.GTrashStack)(stackP.ToC())
	}

	retC := C.g_trash_stack_peek(c_stack_p)
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// Pops a piece of memory off a #GTrashStack.
/*

C function

g_trash_stack_pop
*/
func TrashStackPop(stackP *TrashStack) uintptr {
	c_stack_p := (**C.GTrashStack)(C.NULL)
	if stackP != nil {
		c_stack_p = (**C.GTrashStack)(stackP.ToC())
	}

	retC := C.g_trash_stack_pop(c_stack_p)
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// Pushes a piece of memory onto a #GTrashStack.
/*

C function

g_trash_stack_push
*/
func TrashStackPush(stackP *TrashStack, dataP uintptr) {
	c_stack_p := (**C.GTrashStack)(C.NULL)
	if stackP != nil {
		c_stack_p = (**C.GTrashStack)(stackP.ToC())
	}

	c_data_p := (C.gpointer)(dataP)

	C.g_trash_stack_push(c_stack_p, c_data_p)

	return
}

// Attempts to allocate @n_bytes, and returns %NULL on failure.
// Contrast with g_malloc(), which aborts the program on failure.
/*

C function

g_try_malloc
*/
func TryMalloc(nBytes uint64) uintptr {
	c_n_bytes := (C.gsize)(nBytes)

	retC := C.g_try_malloc(c_n_bytes)
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// Attempts to realloc @mem to a new size, @n_bytes, and returns %NULL
// on failure. Contrast with g_realloc(), which aborts the program
// on failure.
//
// If @mem is %NULL, behaves the same as g_try_malloc().
/*

C function

g_try_realloc
*/
func TryRealloc(mem uintptr, nBytes uint64) uintptr {
	c_mem := (C.gpointer)(mem)

	c_n_bytes := (C.gsize)(nBytes)

	retC := C.g_try_realloc(c_mem, c_n_bytes)
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_ucs4_to_utf16 : no return generator

// Convert a string from a 32-bit fixed width representation as UCS-4.
// to UTF-8. The result will be terminated with a 0 byte.
/*

C function

g_ucs4_to_utf8
*/
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

// Determines the break type of @c. @c should be a Unicode character
// (to derive a character from UTF-8 encoded text, use
// g_utf8_get_char()). The break type is used to find word and line
// breaks ("text boundaries"), Pango implements the Unicode boundary
// resolution algorithms and normally you would use a function such
// as pango_break() instead of caring about break types yourself.
/*

C function

g_unichar_break_type
*/
func UnicharBreakType(c rune) UnicodeBreakType {
	c_c := (C.gunichar)(c)

	retC := C.g_unichar_break_type(c_c)
	retGo := (UnicodeBreakType)(retC)

	return retGo
}

// Determines the numeric value of a character as a decimal
// digit.
/*

C function

g_unichar_digit_value
*/
func UnicharDigitValue(c rune) int32 {
	c_c := (C.gunichar)(c)

	retC := C.g_unichar_digit_value(c_c)
	retGo := (int32)(retC)

	return retGo
}

// Determines whether a character is alphanumeric.
// Given some UTF-8 text, obtain a character value
// with g_utf8_get_char().
/*

C function

g_unichar_isalnum
*/
func UnicharIsalnum(c rune) bool {
	c_c := (C.gunichar)(c)

	retC := C.g_unichar_isalnum(c_c)
	retGo := retC == C.TRUE

	return retGo
}

// Determines whether a character is alphabetic (i.e. a letter).
// Given some UTF-8 text, obtain a character value with
// g_utf8_get_char().
/*

C function

g_unichar_isalpha
*/
func UnicharIsalpha(c rune) bool {
	c_c := (C.gunichar)(c)

	retC := C.g_unichar_isalpha(c_c)
	retGo := retC == C.TRUE

	return retGo
}

// Determines whether a character is a control character.
// Given some UTF-8 text, obtain a character value with
// g_utf8_get_char().
/*

C function

g_unichar_iscntrl
*/
func UnicharIscntrl(c rune) bool {
	c_c := (C.gunichar)(c)

	retC := C.g_unichar_iscntrl(c_c)
	retGo := retC == C.TRUE

	return retGo
}

// Determines if a given character is assigned in the Unicode
// standard.
/*

C function

g_unichar_isdefined
*/
func UnicharIsdefined(c rune) bool {
	c_c := (C.gunichar)(c)

	retC := C.g_unichar_isdefined(c_c)
	retGo := retC == C.TRUE

	return retGo
}

// Determines whether a character is numeric (i.e. a digit).  This
// covers ASCII 0-9 and also digits in other languages/scripts.  Given
// some UTF-8 text, obtain a character value with g_utf8_get_char().
/*

C function

g_unichar_isdigit
*/
func UnicharIsdigit(c rune) bool {
	c_c := (C.gunichar)(c)

	retC := C.g_unichar_isdigit(c_c)
	retGo := retC == C.TRUE

	return retGo
}

// Determines whether a character is printable and not a space
// (returns %FALSE for control characters, format characters, and
// spaces). g_unichar_isprint() is similar, but returns %TRUE for
// spaces. Given some UTF-8 text, obtain a character value with
// g_utf8_get_char().
/*

C function

g_unichar_isgraph
*/
func UnicharIsgraph(c rune) bool {
	c_c := (C.gunichar)(c)

	retC := C.g_unichar_isgraph(c_c)
	retGo := retC == C.TRUE

	return retGo
}

// Determines whether a character is a lowercase letter.
// Given some UTF-8 text, obtain a character value with
// g_utf8_get_char().
/*

C function

g_unichar_islower
*/
func UnicharIslower(c rune) bool {
	c_c := (C.gunichar)(c)

	retC := C.g_unichar_islower(c_c)
	retGo := retC == C.TRUE

	return retGo
}

// Determines whether a character is printable.
// Unlike g_unichar_isgraph(), returns %TRUE for spaces.
// Given some UTF-8 text, obtain a character value with
// g_utf8_get_char().
/*

C function

g_unichar_isprint
*/
func UnicharIsprint(c rune) bool {
	c_c := (C.gunichar)(c)

	retC := C.g_unichar_isprint(c_c)
	retGo := retC == C.TRUE

	return retGo
}

// Determines whether a character is punctuation or a symbol.
// Given some UTF-8 text, obtain a character value with
// g_utf8_get_char().
/*

C function

g_unichar_ispunct
*/
func UnicharIspunct(c rune) bool {
	c_c := (C.gunichar)(c)

	retC := C.g_unichar_ispunct(c_c)
	retGo := retC == C.TRUE

	return retGo
}

// Determines whether a character is a space, tab, or line separator
// (newline, carriage return, etc.).  Given some UTF-8 text, obtain a
// character value with g_utf8_get_char().
//
// (Note: don't use this to do word breaking; you have to use
// Pango or equivalent to get word breaking right, the algorithm
// is fairly complex.)
/*

C function

g_unichar_isspace
*/
func UnicharIsspace(c rune) bool {
	c_c := (C.gunichar)(c)

	retC := C.g_unichar_isspace(c_c)
	retGo := retC == C.TRUE

	return retGo
}

// Determines if a character is titlecase. Some characters in
// Unicode which are composites, such as the DZ digraph
// have three case variants instead of just two. The titlecase
// form is used at the beginning of a word where only the
// first letter is capitalized. The titlecase form of the DZ
// digraph is U+01F2 LATIN CAPITAL LETTTER D WITH SMALL LETTER Z.
/*

C function

g_unichar_istitle
*/
func UnicharIstitle(c rune) bool {
	c_c := (C.gunichar)(c)

	retC := C.g_unichar_istitle(c_c)
	retGo := retC == C.TRUE

	return retGo
}

// Determines if a character is uppercase.
/*

C function

g_unichar_isupper
*/
func UnicharIsupper(c rune) bool {
	c_c := (C.gunichar)(c)

	retC := C.g_unichar_isupper(c_c)
	retGo := retC == C.TRUE

	return retGo
}

// Determines if a character is typically rendered in a double-width
// cell.
/*

C function

g_unichar_iswide
*/
func UnicharIswide(c rune) bool {
	c_c := (C.gunichar)(c)

	retC := C.g_unichar_iswide(c_c)
	retGo := retC == C.TRUE

	return retGo
}

// Determines if a character is a hexidecimal digit.
/*

C function

g_unichar_isxdigit
*/
func UnicharIsxdigit(c rune) bool {
	c_c := (C.gunichar)(c)

	retC := C.g_unichar_isxdigit(c_c)
	retGo := retC == C.TRUE

	return retGo
}

// Blacklisted : g_unichar_to_utf8

// Converts a character to lower case.
/*

C function

g_unichar_tolower
*/
func UnicharTolower(c rune) rune {
	c_c := (C.gunichar)(c)

	retC := C.g_unichar_tolower(c_c)
	retGo := (rune)(retC)

	return retGo
}

// Converts a character to the titlecase.
/*

C function

g_unichar_totitle
*/
func UnicharTotitle(c rune) rune {
	c_c := (C.gunichar)(c)

	retC := C.g_unichar_totitle(c_c)
	retGo := (rune)(retC)

	return retGo
}

// Converts a character to uppercase.
/*

C function

g_unichar_toupper
*/
func UnicharToupper(c rune) rune {
	c_c := (C.gunichar)(c)

	retC := C.g_unichar_toupper(c_c)
	retGo := (rune)(retC)

	return retGo
}

// Classifies a Unicode character by type.
/*

C function

g_unichar_type
*/
func UnicharType(c rune) UnicodeType {
	c_c := (C.gunichar)(c)

	retC := C.g_unichar_type(c_c)
	retGo := (UnicodeType)(retC)

	return retGo
}

// Checks whether @ch is a valid Unicode character. Some possible
// integer values of @ch will not be valid. 0 is considered a valid
// character, though it's normally a string terminator.
/*

C function

g_unichar_validate
*/
func UnicharValidate(ch rune) bool {
	c_ch := (C.gunichar)(ch)

	retC := C.g_unichar_validate(c_ch)
	retGo := retC == C.TRUE

	return retGo
}

// Determines the numeric value of a character as a hexidecimal
// digit.
/*

C function

g_unichar_xdigit_value
*/
func UnicharXdigitValue(c rune) int32 {
	c_c := (C.gunichar)(c)

	retC := C.g_unichar_xdigit_value(c_c)
	retGo := (int32)(retC)

	return retGo
}

// Blacklisted : g_unicode_canonical_decomposition

// Computes the canonical ordering of a string in-place.
// This rearranges decomposed characters in the string
// according to their combining classes.  See the Unicode
// manual for more information.
/*

C function

g_unicode_canonical_ordering
*/
func UnicodeCanonicalOrdering(string rune, len uint64) {
	c_string := (C.gunichar)(string)

	c_len := (C.gsize)(len)

	C.g_unicode_canonical_ordering(&c_string, c_len)

	return
}

// Blacklisted : g_unix_error_quark

// Pauses the current thread for the given number of microseconds.
//
// There are 1 million microseconds per second (represented by the
// #G_USEC_PER_SEC macro). g_usleep() may have limited precision,
// depending on hardware and operating system; don't rely on the exact
// length of the sleep.
/*

C function

g_usleep
*/
func Usleep(microseconds uint64) {
	c_microseconds := (C.gulong)(microseconds)

	C.g_usleep(c_microseconds)

	return
}

// Unsupported : g_utf16_to_ucs4 : unsupported parameter str : no type generator for guint16 (const gunichar2*) for param str

// Unsupported : g_utf16_to_utf8 : unsupported parameter str : no type generator for guint16 (const gunichar2*) for param str

// Converts a string into a form that is independent of case. The
// result will not correspond to any particular case, but can be
// compared for equality or ordered with the results of calling
// g_utf8_casefold() on other strings.
//
// Note that calling g_utf8_casefold() followed by g_utf8_collate() is
// only an approximation to the correct linguistic case insensitive
// ordering, though it is a fairly good one. Getting this exactly
// right would require a more sophisticated collation function that
// takes case sensitivity into account. GLib does not currently
// provide such a function.
/*

C function

g_utf8_casefold
*/
func Utf8Casefold(str string, len int64) string {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	c_len := (C.gssize)(len)

	retC := C.g_utf8_casefold(c_str, c_len)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Compares two strings for ordering using the linguistically
// correct rules for the [current locale][setlocale].
// When sorting a large number of strings, it will be significantly
// faster to obtain collation keys with g_utf8_collate_key() and
// compare the keys with strcmp() when sorting instead of sorting
// the original strings.
/*

C function

g_utf8_collate
*/
func Utf8Collate(str1 string, str2 string) int32 {
	c_str1 := C.CString(str1)
	defer C.free(unsafe.Pointer(c_str1))

	c_str2 := C.CString(str2)
	defer C.free(unsafe.Pointer(c_str2))

	retC := C.g_utf8_collate(c_str1, c_str2)
	retGo := (int32)(retC)

	return retGo
}

// Converts a string into a collation key that can be compared
// with other collation keys produced by the same function using
// strcmp().
//
// The results of comparing the collation keys of two strings
// with strcmp() will always be the same as comparing the two
// original keys with g_utf8_collate().
//
// Note that this function depends on the [current locale][setlocale].
/*

C function

g_utf8_collate_key
*/
func Utf8CollateKey(str string, len int64) string {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	c_len := (C.gssize)(len)

	retC := C.g_utf8_collate_key(c_str, c_len)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Finds the start of the next UTF-8 character in the string after @p.
//
// @p does not have to be at the beginning of a UTF-8 character. No check
// is made to see if the character found is actually valid other than
// it starts with an appropriate byte.
//
// If @end is %NULL, the return value will never be %NULL: if the end of the
// string is reached, a pointer to the terminating nul byte is returned. If
// @end is non-%NULL, the return value will be %NULL if the end of the string
// is reached.
/*

C function

g_utf8_find_next_char
*/
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

// Given a position @p with a UTF-8 encoded string @str, find the start
// of the previous UTF-8 character starting before @p. Returns %NULL if no
// UTF-8 characters are present in @str before @p.
//
// @p does not have to be at the beginning of a UTF-8 character. No check
// is made to see if the character found is actually valid other than
// it starts with an appropriate byte.
/*

C function

g_utf8_find_prev_char
*/
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

// Converts a sequence of bytes encoded as UTF-8 to a Unicode character.
//
// If @p does not point to a valid UTF-8 encoded character, results
// are undefined. If you are not sure that the bytes are complete
// valid Unicode characters, you should use g_utf8_get_char_validated()
// instead.
/*

C function

g_utf8_get_char
*/
func Utf8GetChar(p string) rune {
	c_p := C.CString(p)
	defer C.free(unsafe.Pointer(c_p))

	retC := C.g_utf8_get_char(c_p)
	retGo := (rune)(retC)

	return retGo
}

// Convert a sequence of bytes encoded as UTF-8 to a Unicode character.
// This function checks for incomplete characters, for invalid characters
// such as characters that are out of the range of Unicode, and for
// overlong encodings of valid characters.
//
// Note that g_utf8_get_char_validated() returns (gunichar)-2 if
// @max_len is positive and any of the bytes in the first UTF-8 character
// sequence are nul.
/*

C function

g_utf8_get_char_validated
*/
func Utf8GetCharValidated(p string, maxLen int64) rune {
	c_p := C.CString(p)
	defer C.free(unsafe.Pointer(c_p))

	c_max_len := (C.gssize)(maxLen)

	retC := C.g_utf8_get_char_validated(c_p, c_max_len)
	retGo := (rune)(retC)

	return retGo
}

// Converts a string into canonical form, standardizing
// such issues as whether a character with an accent
// is represented as a base character and combining
// accent or as a single precomposed character. The
// string has to be valid UTF-8, otherwise %NULL is
// returned. You should generally call g_utf8_normalize()
// before comparing two Unicode strings.
//
// The normalization mode %G_NORMALIZE_DEFAULT only
// standardizes differences that do not affect the
// text content, such as the above-mentioned accent
// representation. %G_NORMALIZE_ALL also standardizes
// the "compatibility" characters in Unicode, such
// as SUPERSCRIPT THREE to the standard forms
// (in this case DIGIT THREE). Formatting information
// may be lost but for most text operations such
// characters should be considered the same.
//
// %G_NORMALIZE_DEFAULT_COMPOSE and %G_NORMALIZE_ALL_COMPOSE
// are like %G_NORMALIZE_DEFAULT and %G_NORMALIZE_ALL,
// but returned a result with composed forms rather
// than a maximally decomposed form. This is often
// useful if you intend to convert the string to
// a legacy encoding or pass it to a system with
// less capable Unicode handling.
/*

C function

g_utf8_normalize
*/
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

// Converts from an integer character offset to a pointer to a position
// within the string.
//
// Since 2.10, this function allows to pass a negative @offset to
// step backwards. It is usually worth stepping backwards from the end
// instead of forwards if @offset is in the last fourth of the string,
// since moving forward is about 3 times faster than moving backward.
//
// Note that this function doesn't abort when reaching the end of @str.
// Therefore you should be sure that @offset is within string boundaries
// before calling that function. Call g_utf8_strlen() when unsure.
// This limitation exists as this function is called frequently during
// text rendering and therefore has to be as fast as possible.
/*

C function

g_utf8_offset_to_pointer
*/
func Utf8OffsetToPointer(str string, offset int64) string {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	c_offset := (C.glong)(offset)

	retC := C.g_utf8_offset_to_pointer(c_str, c_offset)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Converts from a pointer to position within a string to a integer
// character offset.
//
// Since 2.10, this function allows @pos to be before @str, and returns
// a negative offset in this case.
/*

C function

g_utf8_pointer_to_offset
*/
func Utf8PointerToOffset(str string, pos string) int64 {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	c_pos := C.CString(pos)
	defer C.free(unsafe.Pointer(c_pos))

	retC := C.g_utf8_pointer_to_offset(c_str, c_pos)
	retGo := (int64)(retC)

	return retGo
}

// Finds the previous UTF-8 character in the string before @p.
//
// @p does not have to be at the beginning of a UTF-8 character. No check
// is made to see if the character found is actually valid other than
// it starts with an appropriate byte. If @p might be the first
// character of the string, you must use g_utf8_find_prev_char() instead.
/*

C function

g_utf8_prev_char
*/
func Utf8PrevChar(p string) string {
	c_p := C.CString(p)
	defer C.free(unsafe.Pointer(c_p))

	retC := C.g_utf8_prev_char(c_p)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Finds the leftmost occurrence of the given Unicode character
// in a UTF-8 encoded string, while limiting the search to @len bytes.
// If @len is -1, allow unbounded search.
/*

C function

g_utf8_strchr
*/
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

// Converts all Unicode characters in the string that have a case
// to lowercase. The exact manner that this is done depends
// on the current locale, and may result in the number of
// characters in the string changing.
/*

C function

g_utf8_strdown
*/
func Utf8Strdown(str string, len int64) string {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	c_len := (C.gssize)(len)

	retC := C.g_utf8_strdown(c_str, c_len)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Computes the length of the string in characters, not including
// the terminating nul character. If the @max'th byte falls in the
// middle of a character, the last (partial) character is not counted.
/*

C function

g_utf8_strlen
*/
func Utf8Strlen(p string, max int64) int64 {
	c_p := C.CString(p)
	defer C.free(unsafe.Pointer(c_p))

	c_max := (C.gssize)(max)

	retC := C.g_utf8_strlen(c_p, c_max)
	retGo := (int64)(retC)

	return retGo
}

// Like the standard C strncpy() function, but copies a given number
// of characters instead of a given number of bytes. The @src string
// must be valid UTF-8 encoded text. (Use g_utf8_validate() on all
// text before trying to use UTF-8 utility functions with it.)
//
// Note you must ensure @dest is at least 4 * @n to fit the
// largest possible UTF-8 characters
/*

C function

g_utf8_strncpy
*/
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

// Find the rightmost occurrence of the given Unicode character
// in a UTF-8 encoded string, while limiting the search to @len bytes.
// If @len is -1, allow unbounded search.
/*

C function

g_utf8_strrchr
*/
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

// Converts all Unicode characters in the string that have a case
// to uppercase. The exact manner that this is done depends
// on the current locale, and may result in the number of
// characters in the string increasing. (For instance, the
// German ess-zet will be changed to SS.)
/*

C function

g_utf8_strup
*/
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

// Validates UTF-8 encoded text. @str is the text to validate;
// if @str is nul-terminated, then @max_len can be -1, otherwise
// @max_len should be the number of bytes to validate.
// If @end is non-%NULL, then the end of the valid range
// will be stored there (i.e. the start of the first invalid
// character if some bytes were invalid, or the end of the text
// being validated otherwise).
//
// Note that g_utf8_validate() returns %FALSE if @max_len is
// positive and any of the @max_len bytes are nul.
//
// Returns %TRUE if all of @str was valid. Many GLib and GTK+
// routines require valid UTF-8 as input; so data read from a file
// or the network should be checked with g_utf8_validate() before
// doing anything else with it.
/*

C function

g_utf8_validate
*/
func Utf8Validate(str []uint8) (bool, string) {
	c_str := &str[0]

	c_max_len := (C.gssize)(len(str))

	var c_end *C.gchar

	retC := C.g_utf8_validate((*C.gchar)(unsafe.Pointer(c_str)), c_max_len, &c_end)
	retGo := retC == C.TRUE

	end := C.GoString(c_end)

	return retGo, end
}

// Blacklisted : g_variant_get_gtype

// Unsupported : g_variant_parse : unsupported parameter type : Blacklisted record : GVariantType

/*

C function

g_variant_parse_error_quark
*/
func VariantParseErrorQuark() Quark {
	retC := C.g_variant_parse_error_quark()
	retGo := (Quark)(retC)

	return retGo
}

// Same as g_variant_error_quark().
/*

C function

g_variant_parser_get_error_quark
*/
func VariantParserGetErrorQuark() Quark {
	retC := C.g_variant_parser_get_error_quark()
	retGo := (Quark)(retC)

	return retGo
}

// Unsupported : g_variant_type_checked_ : return type : Blacklisted record : GVariantType

// Checks if @type_string is a valid GVariant type string.  This call is
// equivalent to calling g_variant_type_string_scan() and confirming
// that the following character is a nul terminator.
/*

C function

g_variant_type_string_is_valid
*/
func VariantTypeStringIsValid(typeString string) bool {
	c_type_string := C.CString(typeString)
	defer C.free(unsafe.Pointer(c_type_string))

	retC := C.g_variant_type_string_is_valid(c_type_string)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : g_vsnprintf : unsupported parameter args : no type generator for va_list (va_list) for param args

/*

C function

g_warn_message
*/
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
