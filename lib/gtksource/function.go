// Code generated - DO NOT EDIT.

package gtksource

import (
	gi "github.com/pekim/gobbi/internal/gi"
	"sync"
)

// UNSUPPORTED : C value 'gtk_source_completion_error_quark' : return type 'GLib.Quark' not supported

// UNSUPPORTED : C value 'gtk_source_encoding_get_all' : return type 'GLib.SList' not supported

var EncodingGetCurrentFunction *gi.Function
var EncodingGetCurrentFunctionOnce sync.Once

func EncodingGetCurrentFunctionSet() {
	EncodingGetCurrentFunctionOnce.Do(func() {
		EncodingGetCurrentFunction = gi.FunctionInvokerNew("GtkSource", "encoding_get_current")
	})
}

var encodingGetCurrentInvoker *gi.Function

// EncodingGetCurrent is a representation of the C type gtk_source_encoding_get_current.
func EncodingGetCurrent() *Encoding {
	if encodingGetCurrentInvoker == nil {
		encodingGetCurrentInvoker = gi.FunctionInvokerNew("GtkSource", "encoding_get_current")
	}

	ret := encodingGetCurrentInvoker.Invoke(nil, nil)

	retGo := &Encoding{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'gtk_source_encoding_get_default_candidates' : return type 'GLib.SList' not supported

var EncodingGetFromCharsetFunction *gi.Function
var EncodingGetFromCharsetFunctionOnce sync.Once

func EncodingGetFromCharsetFunctionSet() {
	EncodingGetFromCharsetFunctionOnce.Do(func() {
		EncodingGetFromCharsetFunction = gi.FunctionInvokerNew("GtkSource", "encoding_get_from_charset")
	})
}

var encodingGetFromCharsetInvoker *gi.Function

// EncodingGetFromCharset is a representation of the C type gtk_source_encoding_get_from_charset.
func EncodingGetFromCharset(charset string) *Encoding {
	if encodingGetFromCharsetInvoker == nil {
		encodingGetFromCharsetInvoker = gi.FunctionInvokerNew("GtkSource", "encoding_get_from_charset")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetString(charset)

	ret := encodingGetFromCharsetInvoker.Invoke(inArgs[:], nil)

	retGo := &Encoding{native: ret.Pointer()}

	return retGo
}

var EncodingGetUtf8Function *gi.Function
var EncodingGetUtf8FunctionOnce sync.Once

func EncodingGetUtf8FunctionSet() {
	EncodingGetUtf8FunctionOnce.Do(func() {
		EncodingGetUtf8Function = gi.FunctionInvokerNew("GtkSource", "encoding_get_utf8")
	})
}

var encodingGetUtf8Invoker *gi.Function

// EncodingGetUtf8 is a representation of the C type gtk_source_encoding_get_utf8.
func EncodingGetUtf8() *Encoding {
	if encodingGetUtf8Invoker == nil {
		encodingGetUtf8Invoker = gi.FunctionInvokerNew("GtkSource", "encoding_get_utf8")
	}

	ret := encodingGetUtf8Invoker.Invoke(nil, nil)

	retGo := &Encoding{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'gtk_source_file_loader_error_quark' : return type 'GLib.Quark' not supported

// UNSUPPORTED : C value 'gtk_source_file_saver_error_quark' : return type 'GLib.Quark' not supported

var UtilsEscapeSearchTextFunction *gi.Function
var UtilsEscapeSearchTextFunctionOnce sync.Once

func UtilsEscapeSearchTextFunctionSet() {
	UtilsEscapeSearchTextFunctionOnce.Do(func() {
		UtilsEscapeSearchTextFunction = gi.FunctionInvokerNew("GtkSource", "utils_escape_search_text")
	})
}

var utilsEscapeSearchTextInvoker *gi.Function

// UtilsEscapeSearchText is a representation of the C type gtk_source_utils_escape_search_text.
func UtilsEscapeSearchText(text string) string {
	if utilsEscapeSearchTextInvoker == nil {
		utilsEscapeSearchTextInvoker = gi.FunctionInvokerNew("GtkSource", "utils_escape_search_text")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetString(text)

	ret := utilsEscapeSearchTextInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

var UtilsUnescapeSearchTextFunction *gi.Function
var UtilsUnescapeSearchTextFunctionOnce sync.Once

func UtilsUnescapeSearchTextFunctionSet() {
	UtilsUnescapeSearchTextFunctionOnce.Do(func() {
		UtilsUnescapeSearchTextFunction = gi.FunctionInvokerNew("GtkSource", "utils_unescape_search_text")
	})
}

var utilsUnescapeSearchTextInvoker *gi.Function

// UtilsUnescapeSearchText is a representation of the C type gtk_source_utils_unescape_search_text.
func UtilsUnescapeSearchText(text string) string {
	if utilsUnescapeSearchTextInvoker == nil {
		utilsUnescapeSearchTextInvoker = gi.FunctionInvokerNew("GtkSource", "utils_unescape_search_text")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetString(text)

	ret := utilsUnescapeSearchTextInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}
