// Code generated - DO NOT EDIT.

package soup

import gi "github.com/pekim/gobbi/internal/gi"

// UNSUPPORTED : C value 'soup_add_completion' : non trivial function

// UNSUPPORTED : C value 'soup_add_idle' : non trivial function

// UNSUPPORTED : C value 'soup_add_io_watch' : non trivial function

// UNSUPPORTED : C value 'soup_add_timeout' : non trivial function

// UNSUPPORTED : C value 'soup_check_version' : non trivial function

// UNSUPPORTED : C value 'soup_cookie_parse' : non trivial function

// UNSUPPORTED : C value 'soup_cookies_free' : non trivial function

// UNSUPPORTED : C value 'soup_cookies_from_request' : non trivial function

// UNSUPPORTED : C value 'soup_cookies_from_response' : non trivial function

// UNSUPPORTED : C value 'soup_cookies_to_cookie_header' : non trivial function

// UNSUPPORTED : C value 'soup_cookies_to_request' : non trivial function

// UNSUPPORTED : C value 'soup_cookies_to_response' : non trivial function

// UNSUPPORTED : C value 'soup_form_decode' : non trivial function

// UNSUPPORTED : C value 'soup_form_decode_multipart' : non trivial function

// UNSUPPORTED : C value 'soup_form_encode' : non trivial function

// UNSUPPORTED : C value 'soup_form_encode_datalist' : non trivial function

// UNSUPPORTED : C value 'soup_form_encode_hash' : non trivial function

// UNSUPPORTED : C value 'soup_form_encode_valist' : non trivial function

// UNSUPPORTED : C value 'soup_form_request_new' : non trivial function

// UNSUPPORTED : C value 'soup_form_request_new_from_datalist' : non trivial function

// UNSUPPORTED : C value 'soup_form_request_new_from_hash' : non trivial function

// UNSUPPORTED : C value 'soup_form_request_new_from_multipart' : non trivial function

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

// UNSUPPORTED : C value 'soup_header_contains' : non trivial function

// UNSUPPORTED : C value 'soup_header_free_list' : non trivial function

// UNSUPPORTED : C value 'soup_header_free_param_list' : non trivial function

// UNSUPPORTED : C value 'soup_header_g_string_append_param' : non trivial function

// UNSUPPORTED : C value 'soup_header_g_string_append_param_quoted' : non trivial function

// UNSUPPORTED : C value 'soup_header_parse_list' : non trivial function

// UNSUPPORTED : C value 'soup_header_parse_param_list' : non trivial function

// UNSUPPORTED : C value 'soup_header_parse_quality_list' : non trivial function

// UNSUPPORTED : C value 'soup_header_parse_semi_param_list' : non trivial function

// UNSUPPORTED : C value 'soup_headers_parse' : non trivial function

// UNSUPPORTED : C value 'soup_headers_parse_request' : non trivial function

// UNSUPPORTED : C value 'soup_headers_parse_response' : non trivial function

// UNSUPPORTED : C value 'soup_headers_parse_status_line' : non trivial function

var httpErrorQuarkInvoker *gi.Function

// HttpErrorQuark is a representation of the C type soup_http_error_quark.
func HttpErrorQuark() {
	if httpErrorQuarkInvoker == nil {
		httpErrorQuarkInvoker = gi.FunctionInvokerNew("Soup", "http_error_quark")
	}

	httpErrorQuarkInvoker.Invoke()
}

// UNSUPPORTED : C value 'soup_message_headers_iter_init' : non trivial function

var requestErrorQuarkInvoker *gi.Function

// RequestErrorQuark is a representation of the C type soup_request_error_quark.
func RequestErrorQuark() {
	if requestErrorQuarkInvoker == nil {
		requestErrorQuarkInvoker = gi.FunctionInvokerNew("Soup", "request_error_quark")
	}

	requestErrorQuarkInvoker.Invoke()
}

var requesterErrorQuarkInvoker *gi.Function

// RequesterErrorQuark is a representation of the C type soup_requester_error_quark.
func RequesterErrorQuark() {
	if requesterErrorQuarkInvoker == nil {
		requesterErrorQuarkInvoker = gi.FunctionInvokerNew("Soup", "requester_error_quark")
	}

	requesterErrorQuarkInvoker.Invoke()
}

// UNSUPPORTED : C value 'soup_status_get_phrase' : non trivial function

// UNSUPPORTED : C value 'soup_status_proxify' : non trivial function

// UNSUPPORTED : C value 'soup_str_case_equal' : non trivial function

// UNSUPPORTED : C value 'soup_str_case_hash' : non trivial function

// UNSUPPORTED : C value 'soup_tld_domain_is_public_suffix' : non trivial function

var tldErrorQuarkInvoker *gi.Function

// TldErrorQuark is a representation of the C type soup_tld_error_quark.
func TldErrorQuark() {
	if tldErrorQuarkInvoker == nil {
		tldErrorQuarkInvoker = gi.FunctionInvokerNew("Soup", "tld_error_quark")
	}

	tldErrorQuarkInvoker.Invoke()
}

// UNSUPPORTED : C value 'soup_tld_get_base_domain' : non trivial function

// UNSUPPORTED : C value 'soup_uri_decode' : non trivial function

// UNSUPPORTED : C value 'soup_uri_encode' : non trivial function

// UNSUPPORTED : C value 'soup_uri_normalize' : non trivial function

// UNSUPPORTED : C value 'soup_value_array_append' : non trivial function

// UNSUPPORTED : C value 'soup_value_array_append_vals' : non trivial function

// UNSUPPORTED : C value 'soup_value_array_from_args' : non trivial function

// UNSUPPORTED : C value 'soup_value_array_get_nth' : non trivial function

// UNSUPPORTED : C value 'soup_value_array_insert' : non trivial function

var valueArrayNewInvoker *gi.Function

// ValueArrayNew is a representation of the C type soup_value_array_new.
func ValueArrayNew() {
	if valueArrayNewInvoker == nil {
		valueArrayNewInvoker = gi.FunctionInvokerNew("Soup", "value_array_new")
	}

	valueArrayNewInvoker.Invoke()
}

// UNSUPPORTED : C value 'soup_value_array_new_with_vals' : non trivial function

// UNSUPPORTED : C value 'soup_value_array_to_args' : non trivial function

// UNSUPPORTED : C value 'soup_value_hash_insert' : non trivial function

// UNSUPPORTED : C value 'soup_value_hash_insert_vals' : non trivial function

// UNSUPPORTED : C value 'soup_value_hash_insert_value' : non trivial function

// UNSUPPORTED : C value 'soup_value_hash_lookup' : non trivial function

// UNSUPPORTED : C value 'soup_value_hash_lookup_vals' : non trivial function

var valueHashNewInvoker *gi.Function

// ValueHashNew is a representation of the C type soup_value_hash_new.
func ValueHashNew() {
	if valueHashNewInvoker == nil {
		valueHashNewInvoker = gi.FunctionInvokerNew("Soup", "value_hash_new")
	}

	valueHashNewInvoker.Invoke()
}

// UNSUPPORTED : C value 'soup_value_hash_new_with_vals' : non trivial function

// UNSUPPORTED : C value 'soup_websocket_client_prepare_handshake' : non trivial function

// UNSUPPORTED : C value 'soup_websocket_client_verify_handshake' : non trivial function

var websocketErrorGetQuarkInvoker *gi.Function

// WebsocketErrorGetQuark is a representation of the C type soup_websocket_error_get_quark.
func WebsocketErrorGetQuark() {
	if websocketErrorGetQuarkInvoker == nil {
		websocketErrorGetQuarkInvoker = gi.FunctionInvokerNew("Soup", "websocket_error_get_quark")
	}

	websocketErrorGetQuarkInvoker.Invoke()
}

// UNSUPPORTED : C value 'soup_websocket_server_check_handshake' : non trivial function

// UNSUPPORTED : C value 'soup_websocket_server_process_handshake' : non trivial function

// UNSUPPORTED : C value 'soup_xmlrpc_build_fault' : non trivial function

// UNSUPPORTED : C value 'soup_xmlrpc_build_method_call' : non trivial function

// UNSUPPORTED : C value 'soup_xmlrpc_build_method_response' : non trivial function

// UNSUPPORTED : C value 'soup_xmlrpc_build_request' : non trivial function

// UNSUPPORTED : C value 'soup_xmlrpc_build_response' : non trivial function

var xmlrpcErrorQuarkInvoker *gi.Function

// XmlrpcErrorQuark is a representation of the C type soup_xmlrpc_error_quark.
func XmlrpcErrorQuark() {
	if xmlrpcErrorQuarkInvoker == nil {
		xmlrpcErrorQuarkInvoker = gi.FunctionInvokerNew("Soup", "xmlrpc_error_quark")
	}

	xmlrpcErrorQuarkInvoker.Invoke()
}

// UNSUPPORTED : C value 'soup_xmlrpc_extract_method_call' : non trivial function

// UNSUPPORTED : C value 'soup_xmlrpc_extract_method_response' : non trivial function

var xmlrpcFaultQuarkInvoker *gi.Function

// XmlrpcFaultQuark is a representation of the C type soup_xmlrpc_fault_quark.
func XmlrpcFaultQuark() {
	if xmlrpcFaultQuarkInvoker == nil {
		xmlrpcFaultQuarkInvoker = gi.FunctionInvokerNew("Soup", "xmlrpc_fault_quark")
	}

	xmlrpcFaultQuarkInvoker.Invoke()
}

// UNSUPPORTED : C value 'soup_xmlrpc_message_new' : non trivial function

// UNSUPPORTED : C value 'soup_xmlrpc_message_set_fault' : non trivial function

// UNSUPPORTED : C value 'soup_xmlrpc_message_set_response' : non trivial function

// UNSUPPORTED : C value 'soup_xmlrpc_parse_method_call' : non trivial function

// UNSUPPORTED : C value 'soup_xmlrpc_parse_method_response' : non trivial function

// UNSUPPORTED : C value 'soup_xmlrpc_parse_request' : non trivial function

// UNSUPPORTED : C value 'soup_xmlrpc_parse_response' : non trivial function

// UNSUPPORTED : C value 'soup_xmlrpc_request_new' : non trivial function

// UNSUPPORTED : C value 'soup_xmlrpc_set_fault' : non trivial function

// UNSUPPORTED : C value 'soup_xmlrpc_set_response' : non trivial function

// UNSUPPORTED : C value 'soup_xmlrpc_variant_get_datetime' : non trivial function

// UNSUPPORTED : C value 'soup_xmlrpc_variant_new_datetime' : non trivial function
