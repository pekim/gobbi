// Code generated - DO NOT EDIT.

package gtksource

import (
	gi "github.com/pekim/gobbi/internal/gi"
	"sync"
)

// UNSUPPORTED : C value 'gtk_source_completion_error_quark' : return type 'GLib.Quark' not supported

// UNSUPPORTED : C value 'gtk_source_encoding_get_all' : return type 'GLib.SList' not supported

var encodingGetCurrentFunction *gi.Function
var encodingGetCurrentFunction_Once sync.Once

func encodingGetCurrentFunction_Set() {
	encodingGetCurrentFunction_Once.Do(func() {
		encodingGetCurrentFunction = gi.FunctionInvokerNew("GtkSource", "encoding_get_current")
	})
}

var encodingGetCurrentInvoker *gi.Function

// EncodingGetCurrent is a representation of the C type gtk_source_encoding_get_current.
func EncodingGetCurrent() *Encoding {
	encodingGetCurrentFunction_Set()

	ret := encodingGetCurrentFunction.Invoke(nil, nil)

	retGo := &Encoding{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'gtk_source_encoding_get_default_candidates' : return type 'GLib.SList' not supported

var encodingGetFromCharsetFunction *gi.Function
var encodingGetFromCharsetFunction_Once sync.Once

func encodingGetFromCharsetFunction_Set() {
	encodingGetFromCharsetFunction_Once.Do(func() {
		encodingGetFromCharsetFunction = gi.FunctionInvokerNew("GtkSource", "encoding_get_from_charset")
	})
}

var encodingGetFromCharsetInvoker *gi.Function

// EncodingGetFromCharset is a representation of the C type gtk_source_encoding_get_from_charset.
func EncodingGetFromCharset(charset string) *Encoding {
	encodingGetFromCharsetFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetString(charset)

	ret := encodingGetFromCharsetFunction.Invoke(inArgs[:], nil)

	retGo := &Encoding{native: ret.Pointer()}

	return retGo
}

var encodingGetUtf8Function *gi.Function
var encodingGetUtf8Function_Once sync.Once

func encodingGetUtf8Function_Set() {
	encodingGetUtf8Function_Once.Do(func() {
		encodingGetUtf8Function = gi.FunctionInvokerNew("GtkSource", "encoding_get_utf8")
	})
}

var encodingGetUtf8Invoker *gi.Function

// EncodingGetUtf8 is a representation of the C type gtk_source_encoding_get_utf8.
func EncodingGetUtf8() *Encoding {
	encodingGetUtf8Function_Set()

	ret := encodingGetUtf8Function.Invoke(nil, nil)

	retGo := &Encoding{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'gtk_source_file_loader_error_quark' : return type 'GLib.Quark' not supported

// UNSUPPORTED : C value 'gtk_source_file_saver_error_quark' : return type 'GLib.Quark' not supported

var utilsEscapeSearchTextFunction *gi.Function
var utilsEscapeSearchTextFunction_Once sync.Once

func utilsEscapeSearchTextFunction_Set() {
	utilsEscapeSearchTextFunction_Once.Do(func() {
		utilsEscapeSearchTextFunction = gi.FunctionInvokerNew("GtkSource", "utils_escape_search_text")
	})
}

var utilsEscapeSearchTextInvoker *gi.Function

// UtilsEscapeSearchText is a representation of the C type gtk_source_utils_escape_search_text.
func UtilsEscapeSearchText(text string) string {
	utilsEscapeSearchTextFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetString(text)

	ret := utilsEscapeSearchTextFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

var utilsUnescapeSearchTextFunction *gi.Function
var utilsUnescapeSearchTextFunction_Once sync.Once

func utilsUnescapeSearchTextFunction_Set() {
	utilsUnescapeSearchTextFunction_Once.Do(func() {
		utilsUnescapeSearchTextFunction = gi.FunctionInvokerNew("GtkSource", "utils_unescape_search_text")
	})
}

var utilsUnescapeSearchTextInvoker *gi.Function

// UtilsUnescapeSearchText is a representation of the C type gtk_source_utils_unescape_search_text.
func UtilsUnescapeSearchText(text string) string {
	utilsUnescapeSearchTextFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetString(text)

	ret := utilsUnescapeSearchTextFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}
