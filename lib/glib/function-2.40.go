// This is a generated file - DO NOT EDIT
// +build glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

// Determines if a string is pure ASCII. A string is pure ASCII if it
// contains no bytes with the high bit set.
/*

C function

g_str_is_ascii
*/
func StrIsAscii(str string) bool {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	retC := C.g_str_is_ascii(c_str)
	retGo := retC == C.TRUE

	return retGo
}

// Checks if a search conducted for @search_term should match
// @potential_hit.
//
// This function calls g_str_tokenize_and_fold() on both
// @search_term and @potential_hit.  ASCII alternates are never taken
// for @search_term but will be taken for @potential_hit according to
// the value of @accept_alternates.
//
// A hit occurs when each folded token in @search_term is a prefix of a
// folded token from @potential_hit.
//
// Depending on how you're performing the search, it will typically be
// faster to call g_str_tokenize_and_fold() on each string in
// your corpus and build an index on the returned folded tokens, then
// call g_str_tokenize_and_fold() on the search term and
// perform lookups into that index.
//
// As some examples, searching for ‘fred’ would match the potential hit
// ‘Smith, Fred’ and also ‘Frédéric’.  Searching for ‘Fréd’ would match
// ‘Frédéric’ but not ‘Frederic’ (due to the one-directional nature of
// accent matching).  Searching ‘fo’ would match ‘Foo’ and ‘Bar Foo
// Baz’, but not ‘SFO’ (because no word has ‘fo’ as a prefix).
/*

C function

g_str_match_string
*/
func StrMatchString(searchTerm string, potentialHit string, acceptAlternates bool) bool {
	c_search_term := C.CString(searchTerm)
	defer C.free(unsafe.Pointer(c_search_term))

	c_potential_hit := C.CString(potentialHit)
	defer C.free(unsafe.Pointer(c_potential_hit))

	c_accept_alternates :=
		boolToGboolean(acceptAlternates)

	retC := C.g_str_match_string(c_search_term, c_potential_hit, c_accept_alternates)
	retGo := retC == C.TRUE

	return retGo
}

// Transliterate @str to plain ASCII.
//
// For best results, @str should be in composed normalised form.
//
// This function performs a reasonably good set of character
// replacements.  The particular set of replacements that is done may
// change by version or even by runtime environment.
//
// If the source language of @str is known, it can used to improve the
// accuracy of the translation by passing it as @from_locale.  It should
// be a valid POSIX locale string (of the form
// "language[_territory][.codeset][@modifier]").
//
// If @from_locale is %NULL then the current locale is used.
//
// If you want to do translation for no specific locale, and you want it
// to be done independently of the currently locale, specify "C" for
// @from_locale.
/*

C function

g_str_to_ascii
*/
func StrToAscii(str string, fromLocale string) string {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	c_from_locale := C.CString(fromLocale)
	defer C.free(unsafe.Pointer(c_from_locale))

	retC := C.g_str_to_ascii(c_str, c_from_locale)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_str_tokenize_and_fold : unsupported parameter ascii_alternates : output array param ascii_alternates

// Pretty-prints a message showing the context of a #GVariant parse
// error within the string for which parsing was attempted.
//
// The resulting string is suitable for output to the console or other
// monospace media where newlines are treated in the usual way.
//
// The message will typically look something like one of the following:
//
// |[
// unterminated string constant:
// (1, 2, 3, 'abc
// ^^^^
// ]|
//
// or
//
// |[
// unable to find a common type:
// [1, 2, 3, 'str']
// ^        ^^^^^
// ]|
//
// The format of the message may change in a future version.
//
// @error must have come from a failed attempt to g_variant_parse() and
// @source_str must be exactly the same string that caused the error.
// If @source_str was not nul-terminated when you passed it to
// g_variant_parse() then you must add nul termination before using this
// function.
/*

C function

g_variant_parse_error_print_context
*/
func VariantParseErrorPrintContext(error *Error, sourceStr string) string {
	c_error := (*C.GError)(C.NULL)
	if error != nil {
		c_error = (*C.GError)(error.ToC())
	}

	c_source_str := C.CString(sourceStr)
	defer C.free(unsafe.Pointer(c_source_str))

	retC := C.g_variant_parse_error_print_context(c_error, c_source_str)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}
