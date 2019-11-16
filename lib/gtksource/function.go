// Code generated - DO NOT EDIT.

package gtksource

import gi "github.com/pekim/gobbi/internal/gi"

// UNSUPPORTED : C value 'gtk_source_completion_error_quark' : return type 'GLib.Quark' not supported

// UNSUPPORTED : C value 'gtk_source_encoding_get_all' : return type 'GLib.SList' not supported

// UNSUPPORTED : C value 'gtk_source_encoding_get_current' : return type 'Encoding' not supported

// UNSUPPORTED : C value 'gtk_source_encoding_get_default_candidates' : return type 'GLib.SList' not supported

// UNSUPPORTED : C value 'gtk_source_encoding_get_from_charset' : return type 'Encoding' not supported

// UNSUPPORTED : C value 'gtk_source_encoding_get_utf8' : return type 'Encoding' not supported

// UNSUPPORTED : C value 'gtk_source_file_loader_error_quark' : return type 'GLib.Quark' not supported

// UNSUPPORTED : C value 'gtk_source_file_saver_error_quark' : return type 'GLib.Quark' not supported

var utilsEscapeSearchTextInvoker *gi.Function

// UtilsEscapeSearchText is a representation of the C type gtk_source_utils_escape_search_text.
func UtilsEscapeSearchText(text string) string {
	if utilsEscapeSearchTextInvoker == nil {
		utilsEscapeSearchTextInvoker = gi.FunctionInvokerNew("GtkSource", "utils_escape_search_text")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetString(text)

	ret := utilsEscapeSearchTextInvoker.Invoke(inArgs[:])
	return ret.String(true)
}

var utilsUnescapeSearchTextInvoker *gi.Function

// UtilsUnescapeSearchText is a representation of the C type gtk_source_utils_unescape_search_text.
func UtilsUnescapeSearchText(text string) string {
	if utilsUnescapeSearchTextInvoker == nil {
		utilsUnescapeSearchTextInvoker = gi.FunctionInvokerNew("GtkSource", "utils_unescape_search_text")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetString(text)

	ret := utilsUnescapeSearchTextInvoker.Invoke(inArgs[:])
	return ret.String(true)
}
