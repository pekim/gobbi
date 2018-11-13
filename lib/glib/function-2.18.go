// This is a generated file - DO NOT EDIT
// +build glib_2.18 glib_2.20 glib_2.22 glib_2.24 glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

// This function is a wrapper of dgettext() which does not translate
// the message if the default domain as set with textdomain() has no
// translations for the current locale.
//
// The advantage of using this function over dgettext() proper is that
// libraries using this function (like GTK+) will not use translations
// if the application using the library does not have translations for
// the current locale.  This results in a consistent English-only
// interface instead of one having partial translations.  For this
// feature to work, the call to textdomain() and setlocale() should
// precede any g_dgettext() invocations.  For GTK+, it means calling
// textdomain() before gtk_init or its variants.
//
// This function disables translations if and only if upon its first
// call all the following conditions hold:
//
// - @domain is not %NULL
//
// - textdomain() has been called to set a default text domain
//
// - there is no translations available for the default text domain
// and the current locale
//
// - current locale is not "C" or any English locales (those
// starting with "en_")
//
// Note that this behavior may not be desired for example if an application
// has its untranslated messages in a language other than English. In those
// cases the application should call textdomain() after initializing GTK+.
//
// Applications should normally not use this function directly,
// but use the _() macro for translations.
/*

C function

g_dgettext
*/
func Dgettext(domain string, msgid string) string {
	c_domain := C.CString(domain)
	defer C.free(unsafe.Pointer(c_domain))

	c_msgid := C.CString(msgid)
	defer C.free(unsafe.Pointer(c_msgid))

	retC := C.g_dgettext(c_domain, c_msgid)
	retGo := C.GoString(retC)

	return retGo
}

// This function is a wrapper of dngettext() which does not translate
// the message if the default domain as set with textdomain() has no
// translations for the current locale.
//
// See g_dgettext() for details of how this differs from dngettext()
// proper.
/*

C function

g_dngettext
*/
func Dngettext(domain string, msgid string, msgidPlural string, n uint64) string {
	c_domain := C.CString(domain)
	defer C.free(unsafe.Pointer(c_domain))

	c_msgid := C.CString(msgid)
	defer C.free(unsafe.Pointer(c_msgid))

	c_msgid_plural := C.CString(msgidPlural)
	defer C.free(unsafe.Pointer(c_msgid_plural))

	c_n := (C.gulong)(n)

	retC := C.g_dngettext(c_domain, c_msgid, c_msgid_plural, c_n)
	retGo := C.GoString(retC)

	return retGo
}

// This function is a variant of g_dgettext() which supports
// a disambiguating message context. GNU gettext uses the
// '\004' character to separate the message context and
// message id in @msgctxtid.
//
// This uses g_dgettext() internally. See that functions for differences
// with dgettext() proper.
//
// This function differs from C_() in that it is not a macro and
// thus you may use non-string-literals as context and msgid arguments.
/*

C function

g_dpgettext2
*/
func Dpgettext2(domain string, context string, msgid string) string {
	c_domain := C.CString(domain)
	defer C.free(unsafe.Pointer(c_domain))

	c_context := C.CString(context)
	defer C.free(unsafe.Pointer(c_context))

	c_msgid := C.CString(msgid)
	defer C.free(unsafe.Pointer(c_msgid))

	retC := C.g_dpgettext2(c_domain, c_context, c_msgid)
	retGo := C.GoString(retC)

	return retGo
}

// Does nothing if @err is %NULL; if @err is non-%NULL, then *@err
// must be %NULL. A new #GError is created and assigned to *@err.
// Unlike g_set_error(), @message is not a printf()-style format string.
// Use this function if @message contains text you don't have control over,
// that could include printf() escape sequences.
/*

C function

g_set_error_literal
*/
func SetErrorLiteral(domain Quark, code int32, message string) *Error {
	var c_err *C.GError

	c_domain := (C.GQuark)(domain)

	c_code := (C.gint)(code)

	c_message := C.CString(message)
	defer C.free(unsafe.Pointer(c_message))

	C.g_set_error_literal(&c_err, c_domain, c_code, c_message)

	err := ErrorNewFromC(unsafe.Pointer(c_err))

	return err
}
