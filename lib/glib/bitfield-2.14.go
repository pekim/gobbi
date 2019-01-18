// This is a generated file - DO NOT EDIT
// +build glib_2.14 glib_2.16 glib_2.18 glib_2.20 glib_2.22 glib_2.24 glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <glib-object.h>
// #include <stdlib.h>
import "C"

type RegexCompileFlags C.GRegexCompileFlags

const (
	REGEX_CASELESS          RegexCompileFlags = 1
	REGEX_MULTILINE         RegexCompileFlags = 2
	REGEX_DOTALL            RegexCompileFlags = 4
	REGEX_EXTENDED          RegexCompileFlags = 8
	REGEX_ANCHORED          RegexCompileFlags = 16
	REGEX_DOLLAR_ENDONLY    RegexCompileFlags = 32
	REGEX_UNGREEDY          RegexCompileFlags = 512
	REGEX_RAW               RegexCompileFlags = 2048
	REGEX_NO_AUTO_CAPTURE   RegexCompileFlags = 4096
	REGEX_OPTIMIZE          RegexCompileFlags = 8192
	REGEX_FIRSTLINE         RegexCompileFlags = 262144
	REGEX_DUPNAMES          RegexCompileFlags = 524288
	REGEX_NEWLINE_CR        RegexCompileFlags = 1048576
	REGEX_NEWLINE_LF        RegexCompileFlags = 2097152
	REGEX_NEWLINE_CRLF      RegexCompileFlags = 3145728
	REGEX_NEWLINE_ANYCRLF   RegexCompileFlags = 5242880
	REGEX_BSR_ANYCRLF       RegexCompileFlags = 8388608
	REGEX_JAVASCRIPT_COMPAT RegexCompileFlags = 33554432
)

type RegexMatchFlags C.GRegexMatchFlags

const (
	REGEX_MATCH_ANCHORED         RegexMatchFlags = 16
	REGEX_MATCH_NOTBOL           RegexMatchFlags = 128
	REGEX_MATCH_NOTEOL           RegexMatchFlags = 256
	REGEX_MATCH_NOTEMPTY         RegexMatchFlags = 1024
	REGEX_MATCH_PARTIAL          RegexMatchFlags = 32768
	REGEX_MATCH_NEWLINE_CR       RegexMatchFlags = 1048576
	REGEX_MATCH_NEWLINE_LF       RegexMatchFlags = 2097152
	REGEX_MATCH_NEWLINE_CRLF     RegexMatchFlags = 3145728
	REGEX_MATCH_NEWLINE_ANY      RegexMatchFlags = 4194304
	REGEX_MATCH_NEWLINE_ANYCRLF  RegexMatchFlags = 5242880
	REGEX_MATCH_BSR_ANYCRLF      RegexMatchFlags = 8388608
	REGEX_MATCH_BSR_ANY          RegexMatchFlags = 16777216
	REGEX_MATCH_PARTIAL_SOFT     RegexMatchFlags = 32768
	REGEX_MATCH_PARTIAL_HARD     RegexMatchFlags = 134217728
	REGEX_MATCH_NOTEMPTY_ATSTART RegexMatchFlags = 268435456
)
