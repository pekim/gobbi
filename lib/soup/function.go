// Code generated - DO NOT EDIT.

package soup

import gi "github.com/pekim/gobbi/internal/gi"

// UNSUPPORTED : C value 'soup_add_completion' : parameter 'async_context' of type 'GLib.MainContext' not supported

// UNSUPPORTED : C value 'soup_add_idle' : parameter 'async_context' of type 'GLib.MainContext' not supported

// UNSUPPORTED : C value 'soup_add_io_watch' : parameter 'async_context' of type 'GLib.MainContext' not supported

// UNSUPPORTED : C value 'soup_add_timeout' : parameter 'async_context' of type 'GLib.MainContext' not supported

// UNSUPPORTED : C value 'soup_check_version' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'soup_cookie_parse' : parameter 'header' of type 'utf8' not supported

// UNSUPPORTED : C value 'soup_cookies_free' : parameter 'cookies' of type 'GLib.SList' not supported

// UNSUPPORTED : C value 'soup_cookies_from_request' : parameter 'msg' of type 'Message' not supported

// UNSUPPORTED : C value 'soup_cookies_from_response' : parameter 'msg' of type 'Message' not supported

// UNSUPPORTED : C value 'soup_cookies_to_cookie_header' : parameter 'cookies' of type 'GLib.SList' not supported

// UNSUPPORTED : C value 'soup_cookies_to_request' : parameter 'cookies' of type 'GLib.SList' not supported

// UNSUPPORTED : C value 'soup_cookies_to_response' : parameter 'cookies' of type 'GLib.SList' not supported

// UNSUPPORTED : C value 'soup_form_decode' : parameter 'encoded_form' of type 'utf8' not supported

// UNSUPPORTED : C value 'soup_form_decode_multipart' : parameter 'msg' of type 'Message' not supported

// UNSUPPORTED : C value 'soup_form_encode' : parameter 'first_field' of type 'utf8' not supported

// UNSUPPORTED : C value 'soup_form_encode_datalist' : parameter 'form_data_set' of type 'GLib.Data' not supported

// UNSUPPORTED : C value 'soup_form_encode_hash' : parameter 'form_data_set' of type 'GLib.HashTable' not supported

// UNSUPPORTED : C value 'soup_form_encode_valist' : parameter 'first_field' of type 'utf8' not supported

// UNSUPPORTED : C value 'soup_form_request_new' : parameter 'method' of type 'utf8' not supported

// UNSUPPORTED : C value 'soup_form_request_new_from_datalist' : parameter 'method' of type 'utf8' not supported

// UNSUPPORTED : C value 'soup_form_request_new_from_hash' : parameter 'method' of type 'utf8' not supported

// UNSUPPORTED : C value 'soup_form_request_new_from_multipart' : parameter 'uri' of type 'utf8' not supported

var getMajorVersionInvoker *gi.Function

// GetMajorVersion is a representation of the C type soup_get_major_version.
func GetMajorVersion() uint32 {
	if getMajorVersionInvoker == nil {
		getMajorVersionInvoker = gi.FunctionInvokerNew("Soup", "get_major_version")
	}

	ret := getMajorVersionInvoker.Invoke(nil)
	return ret.Uint32()
}

var getMicroVersionInvoker *gi.Function

// GetMicroVersion is a representation of the C type soup_get_micro_version.
func GetMicroVersion() uint32 {
	if getMicroVersionInvoker == nil {
		getMicroVersionInvoker = gi.FunctionInvokerNew("Soup", "get_micro_version")
	}

	ret := getMicroVersionInvoker.Invoke(nil)
	return ret.Uint32()
}

var getMinorVersionInvoker *gi.Function

// GetMinorVersion is a representation of the C type soup_get_minor_version.
func GetMinorVersion() uint32 {
	if getMinorVersionInvoker == nil {
		getMinorVersionInvoker = gi.FunctionInvokerNew("Soup", "get_minor_version")
	}

	ret := getMinorVersionInvoker.Invoke(nil)
	return ret.Uint32()
}

// UNSUPPORTED : C value 'soup_header_contains' : parameter 'header' of type 'utf8' not supported

// UNSUPPORTED : C value 'soup_header_free_list' : parameter 'list' of type 'GLib.SList' not supported

// UNSUPPORTED : C value 'soup_header_free_param_list' : parameter 'param_list' of type 'GLib.HashTable' not supported

// UNSUPPORTED : C value 'soup_header_g_string_append_param' : parameter 'string' of type 'GLib.String' not supported

// UNSUPPORTED : C value 'soup_header_g_string_append_param_quoted' : parameter 'string' of type 'GLib.String' not supported

// UNSUPPORTED : C value 'soup_header_parse_list' : parameter 'header' of type 'utf8' not supported

// UNSUPPORTED : C value 'soup_header_parse_param_list' : parameter 'header' of type 'utf8' not supported

// UNSUPPORTED : C value 'soup_header_parse_quality_list' : parameter 'header' of type 'utf8' not supported

// UNSUPPORTED : C value 'soup_header_parse_semi_param_list' : parameter 'header' of type 'utf8' not supported

// UNSUPPORTED : C value 'soup_headers_parse' : parameter 'str' of type 'utf8' not supported

// UNSUPPORTED : C value 'soup_headers_parse_request' : parameter 'str' of type 'utf8' not supported

// UNSUPPORTED : C value 'soup_headers_parse_response' : parameter 'str' of type 'utf8' not supported

// UNSUPPORTED : C value 'soup_headers_parse_status_line' : parameter 'status_line' of type 'utf8' not supported

// UNSUPPORTED : C value 'soup_http_error_quark' : return type 'GLib.Quark' not supported

// UNSUPPORTED : C value 'soup_message_headers_iter_init' : parameter 'iter' with direction 'out' not supported

// UNSUPPORTED : C value 'soup_request_error_quark' : return type 'GLib.Quark' not supported

// UNSUPPORTED : C value 'soup_requester_error_quark' : return type 'GLib.Quark' not supported

var statusGetPhraseInvoker *gi.Function

// StatusGetPhrase is a representation of the C type soup_status_get_phrase.
func StatusGetPhrase(statusCode uint32) string {
	if statusGetPhraseInvoker == nil {
		statusGetPhraseInvoker = gi.FunctionInvokerNew("Soup", "status_get_phrase")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetUint32(statusCode)

	ret := statusGetPhraseInvoker.Invoke(inArgs[:])
	return ret.String(false)
}

var statusProxifyInvoker *gi.Function

// StatusProxify is a representation of the C type soup_status_proxify.
func StatusProxify(statusCode uint32) uint32 {
	if statusProxifyInvoker == nil {
		statusProxifyInvoker = gi.FunctionInvokerNew("Soup", "status_proxify")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetUint32(statusCode)

	ret := statusProxifyInvoker.Invoke(inArgs[:])
	return ret.Uint32()
}

// UNSUPPORTED : C value 'soup_str_case_equal' : parameter 'v1' of type 'gpointer' not supported

// UNSUPPORTED : C value 'soup_str_case_hash' : parameter 'key' of type 'gpointer' not supported

// UNSUPPORTED : C value 'soup_tld_domain_is_public_suffix' : parameter 'domain' of type 'utf8' not supported

// UNSUPPORTED : C value 'soup_tld_error_quark' : return type 'GLib.Quark' not supported

// UNSUPPORTED : C value 'soup_tld_get_base_domain' : parameter 'hostname' of type 'utf8' not supported

// UNSUPPORTED : C value 'soup_uri_decode' : parameter 'part' of type 'utf8' not supported

// UNSUPPORTED : C value 'soup_uri_encode' : parameter 'part' of type 'utf8' not supported

// UNSUPPORTED : C value 'soup_uri_normalize' : parameter 'part' of type 'utf8' not supported

// UNSUPPORTED : C value 'soup_value_array_append' : parameter 'array' of type 'GObject.ValueArray' not supported

// UNSUPPORTED : C value 'soup_value_array_append_vals' : parameter 'array' of type 'GObject.ValueArray' not supported

// UNSUPPORTED : C value 'soup_value_array_from_args' : parameter 'args' of type 'va_list' not supported

// UNSUPPORTED : C value 'soup_value_array_get_nth' : parameter 'array' of type 'GObject.ValueArray' not supported

// UNSUPPORTED : C value 'soup_value_array_insert' : parameter 'array' of type 'GObject.ValueArray' not supported

// UNSUPPORTED : C value 'soup_value_array_new' : return type 'GObject.ValueArray' not supported

// UNSUPPORTED : C value 'soup_value_array_new_with_vals' : parameter 'first_type' of type 'GType' not supported

// UNSUPPORTED : C value 'soup_value_array_to_args' : parameter 'array' of type 'GObject.ValueArray' not supported

// UNSUPPORTED : C value 'soup_value_hash_insert' : parameter 'hash' of type 'GLib.HashTable' not supported

// UNSUPPORTED : C value 'soup_value_hash_insert_vals' : parameter 'hash' of type 'GLib.HashTable' not supported

// UNSUPPORTED : C value 'soup_value_hash_insert_value' : parameter 'hash' of type 'GLib.HashTable' not supported

// UNSUPPORTED : C value 'soup_value_hash_lookup' : parameter 'hash' of type 'GLib.HashTable' not supported

// UNSUPPORTED : C value 'soup_value_hash_lookup_vals' : parameter 'hash' of type 'GLib.HashTable' not supported

// UNSUPPORTED : C value 'soup_value_hash_new' : return type 'GLib.HashTable' not supported

// UNSUPPORTED : C value 'soup_value_hash_new_with_vals' : parameter 'first_key' of type 'utf8' not supported

// UNSUPPORTED : C value 'soup_websocket_client_prepare_handshake' : parameter 'msg' of type 'Message' not supported

// UNSUPPORTED : C value 'soup_websocket_client_verify_handshake' : parameter 'msg' of type 'Message' not supported

// UNSUPPORTED : C value 'soup_websocket_error_get_quark' : return type 'GLib.Quark' not supported

// UNSUPPORTED : C value 'soup_websocket_server_check_handshake' : parameter 'msg' of type 'Message' not supported

// UNSUPPORTED : C value 'soup_websocket_server_process_handshake' : parameter 'msg' of type 'Message' not supported

// UNSUPPORTED : C value 'soup_xmlrpc_build_fault' : parameter 'fault_format' of type 'utf8' not supported

// UNSUPPORTED : C value 'soup_xmlrpc_build_method_call' : parameter 'method_name' of type 'utf8' not supported

// UNSUPPORTED : C value 'soup_xmlrpc_build_method_response' : parameter 'value' of type 'GObject.Value' not supported

// UNSUPPORTED : C value 'soup_xmlrpc_build_request' : parameter 'method_name' of type 'utf8' not supported

// UNSUPPORTED : C value 'soup_xmlrpc_build_response' : parameter 'value' of type 'GLib.Variant' not supported

// UNSUPPORTED : C value 'soup_xmlrpc_error_quark' : return type 'GLib.Quark' not supported

// UNSUPPORTED : C value 'soup_xmlrpc_extract_method_call' : parameter 'method_call' of type 'utf8' not supported

// UNSUPPORTED : C value 'soup_xmlrpc_extract_method_response' : parameter 'method_response' of type 'utf8' not supported

// UNSUPPORTED : C value 'soup_xmlrpc_fault_quark' : return type 'GLib.Quark' not supported

// UNSUPPORTED : C value 'soup_xmlrpc_message_new' : parameter 'uri' of type 'utf8' not supported

// UNSUPPORTED : C value 'soup_xmlrpc_message_set_fault' : parameter 'msg' of type 'Message' not supported

// UNSUPPORTED : C value 'soup_xmlrpc_message_set_response' : parameter 'msg' of type 'Message' not supported

// UNSUPPORTED : C value 'soup_xmlrpc_parse_method_call' : parameter 'method_call' of type 'utf8' not supported

// UNSUPPORTED : C value 'soup_xmlrpc_parse_method_response' : parameter 'method_response' of type 'utf8' not supported

// UNSUPPORTED : C value 'soup_xmlrpc_parse_request' : parameter 'method_call' of type 'utf8' not supported

// UNSUPPORTED : C value 'soup_xmlrpc_parse_response' : parameter 'method_response' of type 'utf8' not supported

// UNSUPPORTED : C value 'soup_xmlrpc_request_new' : parameter 'uri' of type 'utf8' not supported

// UNSUPPORTED : C value 'soup_xmlrpc_set_fault' : parameter 'msg' of type 'Message' not supported

// UNSUPPORTED : C value 'soup_xmlrpc_set_response' : parameter 'msg' of type 'Message' not supported

// UNSUPPORTED : C value 'soup_xmlrpc_variant_get_datetime' : parameter 'variant' of type 'GLib.Variant' not supported

// UNSUPPORTED : C value 'soup_xmlrpc_variant_new_datetime' : parameter 'date' of type 'Date' not supported
