// This is a generated file - DO NOT EDIT
// +build glib_2.2 glib_2.4 glib_2.6 glib_2.8 glib_2.10 glib_2.12 glib_2.14 glib_2.16 glib_2.18 glib_2.20 glib_2.22 glib_2.24 glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

// Converts a string to a #guint64 value.
// This function behaves like the standard strtoull() function
// does in the C locale. It does this without actually
// changing the current locale, since that would not be
// thread-safe.
//
// This function is typically used when reading configuration
// files or other non-user input that should be locale independent.
// To handle input from the user you should normally use the
// locale-sensitive system strtoull() function.
//
// If the correct value would cause overflow, %G_MAXUINT64
// is returned, and `ERANGE` is stored in `errno`.
// If the base is outside the valid range, zero is returned, and
// `EINVAL` is stored in `errno`.
// If the string conversion fails, zero is returned, and @endptr returns
// @nptr (if @endptr is non-%NULL).
/*

C function

g_ascii_strtoull
*/
func AsciiStrtoull(nptr string, base uint32) (uint64, string) {
	c_nptr := C.CString(nptr)
	defer C.free(unsafe.Pointer(c_nptr))

	var c_endptr *C.gchar

	c_base := (C.guint)(base)

	retC := C.g_ascii_strtoull(c_nptr, &c_endptr, c_base)
	retGo := (uint64)(retC)

	endptr := C.GoString(c_endptr)

	return retGo, endptr
}

// Unsupported : g_fprintf : unsupported parameter file : no type generator for gpointer (FILE*) for param file

// Gets a human-readable name for the application, as set by
// g_set_application_name(). This name should be localized if
// possible, and is intended for display to the user.  Contrast with
// g_get_prgname(), which gets a non-localized name. If
// g_set_application_name() has not been called, returns the result of
// g_get_prgname() (which may be %NULL if g_set_prgname() has also not
// been called).
/*

C function

g_get_application_name
*/
func GetApplicationName() string {
	retC := C.g_get_application_name()
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : g_printf : unsupported parameter ... : varargs

// Sets a human-readable name for the application. This name should be
// localized if possible, and is intended for display to the user.
// Contrast with g_set_prgname(), which sets a non-localized name.
// g_set_prgname() will be called automatically by gtk_init(),
// but g_set_application_name() will not.
//
// Note that for thread safety reasons, this function can only
// be called once.
//
// The application name will be used in contexts such as error messages,
// or when displaying an application's name in the task list.
/*

C function

g_set_application_name
*/
func SetApplicationName(applicationName string) {
	c_application_name := C.CString(applicationName)
	defer C.free(unsafe.Pointer(c_application_name))

	C.g_set_application_name(c_application_name)

	return
}

// Unsupported : g_sprintf : unsupported parameter ... : varargs

// Looks whether the string @str begins with @prefix.
/*

C function

g_str_has_prefix
*/
func StrHasPrefix(str string, prefix string) bool {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	c_prefix := C.CString(prefix)
	defer C.free(unsafe.Pointer(c_prefix))

	retC := C.g_str_has_prefix(c_str, c_prefix)
	retGo := retC == C.TRUE

	return retGo
}

// Looks whether the string @str ends with @suffix.
/*

C function

g_str_has_suffix
*/
func StrHasSuffix(str string, suffix string) bool {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	c_suffix := C.CString(suffix)
	defer C.free(unsafe.Pointer(c_suffix))

	retC := C.g_str_has_suffix(c_str, c_suffix)
	retGo := retC == C.TRUE

	return retGo
}

// Reverses a UTF-8 string. @str must be valid UTF-8 encoded text.
// (Use g_utf8_validate() on all text before trying to use UTF-8
// utility functions with it.)
//
// This function is intended for programmatic uses of reversed strings.
// It pays no attention to decomposed characters, combining marks, byte
// order marks, directional indicators (LRM, LRO, etc) and similar
// characters which might need special handling when reversing a string
// for display purposes.
//
// Note that unlike g_strreverse(), this function returns
// newly-allocated memory, which should be freed with g_free() when
// no longer needed.
/*

C function

g_utf8_strreverse
*/
func Utf8Strreverse(str string, len int64) string {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	c_len := (C.gssize)(len)

	retC := C.g_utf8_strreverse(c_str, c_len)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_vfprintf : unsupported parameter file : no type generator for gpointer (FILE*) for param file

// Unsupported : g_vprintf : unsupported parameter args : no type generator for va_list (va_list) for param args

// Unsupported : g_vsprintf : unsupported parameter args : no type generator for va_list (va_list) for param args
