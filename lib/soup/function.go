// Code generated - DO NOT EDIT.

package soup

import gi "github.com/pekim/gobbi/internal/gi"

// UNSUPPORTED : C value 'soup_add_completion' : has parameters

// UNSUPPORTED : C value 'soup_add_idle' : has parameters

// UNSUPPORTED : C value 'soup_add_io_watch' : has parameters

// UNSUPPORTED : C value 'soup_add_timeout' : has parameters

// UNSUPPORTED : C value 'soup_check_version' : has parameters

// UNSUPPORTED : C value 'soup_cookie_parse' : has parameters

// UNSUPPORTED : C value 'soup_cookies_free' : has parameters

// UNSUPPORTED : C value 'soup_cookies_from_request' : has parameters

// UNSUPPORTED : C value 'soup_cookies_from_response' : has parameters

// UNSUPPORTED : C value 'soup_cookies_to_cookie_header' : has parameters

// UNSUPPORTED : C value 'soup_cookies_to_request' : has parameters

// UNSUPPORTED : C value 'soup_cookies_to_response' : has parameters

// UNSUPPORTED : C value 'soup_form_decode' : has parameters

// UNSUPPORTED : C value 'soup_form_decode_multipart' : has parameters

// UNSUPPORTED : C value 'soup_form_encode' : has parameters

// UNSUPPORTED : C value 'soup_form_encode_datalist' : has parameters

// UNSUPPORTED : C value 'soup_form_encode_hash' : has parameters

// UNSUPPORTED : C value 'soup_form_encode_valist' : has parameters

// UNSUPPORTED : C value 'soup_form_request_new' : has parameters

// UNSUPPORTED : C value 'soup_form_request_new_from_datalist' : has parameters

// UNSUPPORTED : C value 'soup_form_request_new_from_hash' : has parameters

// UNSUPPORTED : C value 'soup_form_request_new_from_multipart' : has parameters

var getMajorVersionInvoker *gi.Function

// GetMajorVersion is a representation of the C type soup_get_major_version.
func GetMajorVersion() uint32 {
	if getMajorVersionInvoker == nil {
		getMajorVersionInvoker = gi.FunctionInvokerNew("Soup", "get_major_version")
	}

	ret := getMajorVersionInvoker.Invoke()
	return ret.Uint32()
}

var getMicroVersionInvoker *gi.Function

// GetMicroVersion is a representation of the C type soup_get_micro_version.
func GetMicroVersion() uint32 {
	if getMicroVersionInvoker == nil {
		getMicroVersionInvoker = gi.FunctionInvokerNew("Soup", "get_micro_version")
	}

	ret := getMicroVersionInvoker.Invoke()
	return ret.Uint32()
}

var getMinorVersionInvoker *gi.Function

// GetMinorVersion is a representation of the C type soup_get_minor_version.
func GetMinorVersion() uint32 {
	if getMinorVersionInvoker == nil {
		getMinorVersionInvoker = gi.FunctionInvokerNew("Soup", "get_minor_version")
	}

	ret := getMinorVersionInvoker.Invoke()
	return ret.Uint32()
}

// UNSUPPORTED : C value 'soup_header_contains' : has parameters

// UNSUPPORTED : C value 'soup_header_free_list' : has parameters

// UNSUPPORTED : C value 'soup_header_free_param_list' : has parameters

// UNSUPPORTED : C value 'soup_header_g_string_append_param' : has parameters

// UNSUPPORTED : C value 'soup_header_g_string_append_param_quoted' : has parameters

// UNSUPPORTED : C value 'soup_header_parse_list' : has parameters

// UNSUPPORTED : C value 'soup_header_parse_param_list' : has parameters

// UNSUPPORTED : C value 'soup_header_parse_quality_list' : has parameters

// UNSUPPORTED : C value 'soup_header_parse_semi_param_list' : has parameters

// UNSUPPORTED : C value 'soup_headers_parse' : has parameters

// UNSUPPORTED : C value 'soup_headers_parse_request' : has parameters

// UNSUPPORTED : C value 'soup_headers_parse_response' : has parameters

// UNSUPPORTED : C value 'soup_headers_parse_status_line' : has parameters

// UNSUPPORTED : C value 'soup_http_error_quark' : return type 'GLib.Quark' not supported

// UNSUPPORTED : C value 'soup_message_headers_iter_init' : has parameters

// UNSUPPORTED : C value 'soup_request_error_quark' : return type 'GLib.Quark' not supported

// UNSUPPORTED : C value 'soup_requester_error_quark' : return type 'GLib.Quark' not supported

// UNSUPPORTED : C value 'soup_status_get_phrase' : has parameters

// UNSUPPORTED : C value 'soup_status_proxify' : has parameters

// UNSUPPORTED : C value 'soup_str_case_equal' : has parameters

// UNSUPPORTED : C value 'soup_str_case_hash' : has parameters

// UNSUPPORTED : C value 'soup_tld_domain_is_public_suffix' : has parameters

// UNSUPPORTED : C value 'soup_tld_error_quark' : return type 'GLib.Quark' not supported

// UNSUPPORTED : C value 'soup_tld_get_base_domain' : has parameters

// UNSUPPORTED : C value 'soup_uri_decode' : has parameters

// UNSUPPORTED : C value 'soup_uri_encode' : has parameters

// UNSUPPORTED : C value 'soup_uri_normalize' : has parameters

// UNSUPPORTED : C value 'soup_value_array_append' : has parameters

// UNSUPPORTED : C value 'soup_value_array_append_vals' : has parameters

// UNSUPPORTED : C value 'soup_value_array_from_args' : has parameters

// UNSUPPORTED : C value 'soup_value_array_get_nth' : has parameters

// UNSUPPORTED : C value 'soup_value_array_insert' : has parameters

// UNSUPPORTED : C value 'soup_value_array_new' : return type 'GObject.ValueArray' not supported

// UNSUPPORTED : C value 'soup_value_array_new_with_vals' : has parameters

// UNSUPPORTED : C value 'soup_value_array_to_args' : has parameters

// UNSUPPORTED : C value 'soup_value_hash_insert' : has parameters

// UNSUPPORTED : C value 'soup_value_hash_insert_vals' : has parameters

// UNSUPPORTED : C value 'soup_value_hash_insert_value' : has parameters

// UNSUPPORTED : C value 'soup_value_hash_lookup' : has parameters

// UNSUPPORTED : C value 'soup_value_hash_lookup_vals' : has parameters

// UNSUPPORTED : C value 'soup_value_hash_new' : return type 'GLib.HashTable' not supported

// UNSUPPORTED : C value 'soup_value_hash_new_with_vals' : has parameters

// UNSUPPORTED : C value 'soup_websocket_client_prepare_handshake' : has parameters

// UNSUPPORTED : C value 'soup_websocket_client_verify_handshake' : has parameters

// UNSUPPORTED : C value 'soup_websocket_error_get_quark' : return type 'GLib.Quark' not supported

// UNSUPPORTED : C value 'soup_websocket_server_check_handshake' : has parameters

// UNSUPPORTED : C value 'soup_websocket_server_process_handshake' : has parameters

// UNSUPPORTED : C value 'soup_xmlrpc_build_fault' : has parameters

// UNSUPPORTED : C value 'soup_xmlrpc_build_method_call' : has parameters

// UNSUPPORTED : C value 'soup_xmlrpc_build_method_response' : has parameters

// UNSUPPORTED : C value 'soup_xmlrpc_build_request' : has parameters

// UNSUPPORTED : C value 'soup_xmlrpc_build_response' : has parameters

// UNSUPPORTED : C value 'soup_xmlrpc_error_quark' : return type 'GLib.Quark' not supported

// UNSUPPORTED : C value 'soup_xmlrpc_extract_method_call' : has parameters

// UNSUPPORTED : C value 'soup_xmlrpc_extract_method_response' : has parameters

// UNSUPPORTED : C value 'soup_xmlrpc_fault_quark' : return type 'GLib.Quark' not supported

// UNSUPPORTED : C value 'soup_xmlrpc_message_new' : has parameters

// UNSUPPORTED : C value 'soup_xmlrpc_message_set_fault' : has parameters

// UNSUPPORTED : C value 'soup_xmlrpc_message_set_response' : has parameters

// UNSUPPORTED : C value 'soup_xmlrpc_parse_method_call' : has parameters

// UNSUPPORTED : C value 'soup_xmlrpc_parse_method_response' : has parameters

// UNSUPPORTED : C value 'soup_xmlrpc_parse_request' : has parameters

// UNSUPPORTED : C value 'soup_xmlrpc_parse_response' : has parameters

// UNSUPPORTED : C value 'soup_xmlrpc_request_new' : has parameters

// UNSUPPORTED : C value 'soup_xmlrpc_set_fault' : has parameters

// UNSUPPORTED : C value 'soup_xmlrpc_set_response' : has parameters

// UNSUPPORTED : C value 'soup_xmlrpc_variant_get_datetime' : has parameters

// UNSUPPORTED : C value 'soup_xmlrpc_variant_new_datetime' : has parameters
