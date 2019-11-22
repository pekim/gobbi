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

func encodingGetCurrentFunction_Set() error {
	var err error
	encodingGetCurrentFunction_Once.Do(func() {
		encodingGetCurrentFunction, err = gi.FunctionInvokerNew("GtkSource", "encoding_get_current")
	})
	return err
}

// EncodingGetCurrent is a representation of the C type gtk_source_encoding_get_current.
func EncodingGetCurrent() (*Encoding, error) {

	var ret gi.Argument

	err := encodingGetCurrentFunction_Set()
	if err == nil {
		ret = encodingGetCurrentFunction.Invoke(nil, nil)
	}

	retGo := &Encoding{native: ret.Pointer()}

	return retGo, err
}

// UNSUPPORTED : C value 'gtk_source_encoding_get_default_candidates' : return type 'GLib.SList' not supported

var encodingGetFromCharsetFunction *gi.Function
var encodingGetFromCharsetFunction_Once sync.Once

func encodingGetFromCharsetFunction_Set() error {
	var err error
	encodingGetFromCharsetFunction_Once.Do(func() {
		encodingGetFromCharsetFunction, err = gi.FunctionInvokerNew("GtkSource", "encoding_get_from_charset")
	})
	return err
}

// EncodingGetFromCharset is a representation of the C type gtk_source_encoding_get_from_charset.
func EncodingGetFromCharset(charset string) (*Encoding, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(charset)

	var ret gi.Argument

	err := encodingGetFromCharsetFunction_Set()
	if err == nil {
		ret = encodingGetFromCharsetFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Encoding{native: ret.Pointer()}

	return retGo, err
}

var encodingGetUtf8Function *gi.Function
var encodingGetUtf8Function_Once sync.Once

func encodingGetUtf8Function_Set() error {
	var err error
	encodingGetUtf8Function_Once.Do(func() {
		encodingGetUtf8Function, err = gi.FunctionInvokerNew("GtkSource", "encoding_get_utf8")
	})
	return err
}

// EncodingGetUtf8 is a representation of the C type gtk_source_encoding_get_utf8.
func EncodingGetUtf8() (*Encoding, error) {

	var ret gi.Argument

	err := encodingGetUtf8Function_Set()
	if err == nil {
		ret = encodingGetUtf8Function.Invoke(nil, nil)
	}

	retGo := &Encoding{native: ret.Pointer()}

	return retGo, err
}

// UNSUPPORTED : C value 'gtk_source_file_loader_error_quark' : return type 'GLib.Quark' not supported

// UNSUPPORTED : C value 'gtk_source_file_saver_error_quark' : return type 'GLib.Quark' not supported

var utilsEscapeSearchTextFunction *gi.Function
var utilsEscapeSearchTextFunction_Once sync.Once

func utilsEscapeSearchTextFunction_Set() error {
	var err error
	utilsEscapeSearchTextFunction_Once.Do(func() {
		utilsEscapeSearchTextFunction, err = gi.FunctionInvokerNew("GtkSource", "utils_escape_search_text")
	})
	return err
}

// UtilsEscapeSearchText is a representation of the C type gtk_source_utils_escape_search_text.
func UtilsEscapeSearchText(text string) (string, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(text)

	var ret gi.Argument

	err := utilsEscapeSearchTextFunction_Set()
	if err == nil {
		ret = utilsEscapeSearchTextFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo, err
}

var utilsUnescapeSearchTextFunction *gi.Function
var utilsUnescapeSearchTextFunction_Once sync.Once

func utilsUnescapeSearchTextFunction_Set() error {
	var err error
	utilsUnescapeSearchTextFunction_Once.Do(func() {
		utilsUnescapeSearchTextFunction, err = gi.FunctionInvokerNew("GtkSource", "utils_unescape_search_text")
	})
	return err
}

// UtilsUnescapeSearchText is a representation of the C type gtk_source_utils_unescape_search_text.
func UtilsUnescapeSearchText(text string) (string, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(text)

	var ret gi.Argument

	err := utilsUnescapeSearchTextFunction_Set()
	if err == nil {
		ret = utilsUnescapeSearchTextFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo, err
}
