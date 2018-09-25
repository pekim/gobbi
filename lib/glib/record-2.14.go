// This is a generated file - DO NOT EDIT
// +build glib_2.14 glib_2.16 glib_2.18 glib_2.20 glib_2.22 glib_2.24 glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

import "unsafe"

// #define GLIB_DISABLE_DEPRECATION_WARNINGS
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

// Regex is a wrapper around the C record GRegex.
type Regex struct {
	native *C.GRegex
}

func RegexNewFromC(u unsafe.Pointer) *Regex {
	c := (*C.GRegex)(u)
	if c == nil {
		return nil
	}

	g := &Regex{native: c}

	return g
}

func (recv *Regex) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// RegexNew is a wrapper around the C function g_regex_new.
func RegexNew(pattern string, compileOptions RegexCompileFlags, matchOptions RegexMatchFlags) (*Regex, error) {
	c_pattern := C.CString(pattern)
	defer C.free(unsafe.Pointer(c_pattern))

	c_compile_options := (C.GRegexCompileFlags)(compileOptions)

	c_match_options := (C.GRegexMatchFlags)(matchOptions)

	var cThrowableError *C.GError

	retC := C.g_regex_new(c_pattern, c_compile_options, c_match_options, &cThrowableError)
	retGo := RegexNewFromC(unsafe.Pointer(retC))

	goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// GetCaptureCount is a wrapper around the C function g_regex_get_capture_count.
func (recv *Regex) GetCaptureCount() int32 {
	retC := C.g_regex_get_capture_count((*C.GRegex)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetMaxBackref is a wrapper around the C function g_regex_get_max_backref.
func (recv *Regex) GetMaxBackref() int32 {
	retC := C.g_regex_get_max_backref((*C.GRegex)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetPattern is a wrapper around the C function g_regex_get_pattern.
func (recv *Regex) GetPattern() string {
	retC := C.g_regex_get_pattern((*C.GRegex)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetStringNumber is a wrapper around the C function g_regex_get_string_number.
func (recv *Regex) GetStringNumber(name string) int32 {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.g_regex_get_string_number((*C.GRegex)(recv.native), c_name)
	retGo := (int32)(retC)

	return retGo
}

// Match is a wrapper around the C function g_regex_match.
func (recv *Regex) Match(string string, matchOptions RegexMatchFlags) (bool, **MatchInfo) {
	c_string := C.CString(string)
	defer C.free(unsafe.Pointer(c_string))

	c_match_options := (C.GRegexMatchFlags)(matchOptions)

	var c_match_info *C.GMatchInfo

	retC := C.g_regex_match((*C.GRegex)(recv.native), c_string, c_match_options, &c_match_info)
	retGo := retC == C.TRUE

	matchInfo := MatchInfoNewFromC(unsafe.Pointer(c_match_info))

	return retGo, matchInfo
}

// MatchAll is a wrapper around the C function g_regex_match_all.
func (recv *Regex) MatchAll(string string, matchOptions RegexMatchFlags) (bool, **MatchInfo) {
	c_string := C.CString(string)
	defer C.free(unsafe.Pointer(c_string))

	c_match_options := (C.GRegexMatchFlags)(matchOptions)

	var c_match_info *C.GMatchInfo

	retC := C.g_regex_match_all((*C.GRegex)(recv.native), c_string, c_match_options, &c_match_info)
	retGo := retC == C.TRUE

	matchInfo := MatchInfoNewFromC(unsafe.Pointer(c_match_info))

	return retGo, matchInfo
}

// Unsupported : g_regex_match_all_full : unsupported parameter string : no param type

// Unsupported : g_regex_match_full : unsupported parameter string : no param type

// Ref is a wrapper around the C function g_regex_ref.
func (recv *Regex) Ref() *Regex {
	retC := C.g_regex_ref((*C.GRegex)(recv.native))
	retGo := RegexNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_regex_replace : unsupported parameter string : no param type

// Unsupported : g_regex_replace_eval : unsupported parameter string : no param type

// Unsupported : g_regex_replace_literal : unsupported parameter string : no param type

// Unsupported : g_regex_split : no return type

// Unsupported : g_regex_split_full : unsupported parameter string : no param type

// Unsupported : g_regex_unref : no return generator
