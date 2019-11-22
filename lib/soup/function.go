// Code generated - DO NOT EDIT.

package soup

import (
	gi "github.com/pekim/gobbi/internal/gi"
	"sync"
)

// UNSUPPORTED : C value 'soup_add_completion' : parameter 'async_context' of type 'GLib.MainContext' not supported

// UNSUPPORTED : C value 'soup_add_idle' : parameter 'async_context' of type 'GLib.MainContext' not supported

// UNSUPPORTED : C value 'soup_add_io_watch' : parameter 'async_context' of type 'GLib.MainContext' not supported

// UNSUPPORTED : C value 'soup_add_timeout' : parameter 'async_context' of type 'GLib.MainContext' not supported

// UNSUPPORTED : C value 'soup_check_version' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'soup_cookie_parse' : parameter 'origin' of type 'URI' not supported

// UNSUPPORTED : C value 'soup_cookies_free' : parameter 'cookies' of type 'GLib.SList' not supported

// UNSUPPORTED : C value 'soup_cookies_from_request' : parameter 'msg' of type 'Message' not supported

// UNSUPPORTED : C value 'soup_cookies_from_response' : parameter 'msg' of type 'Message' not supported

// UNSUPPORTED : C value 'soup_cookies_to_cookie_header' : parameter 'cookies' of type 'GLib.SList' not supported

// UNSUPPORTED : C value 'soup_cookies_to_request' : parameter 'cookies' of type 'GLib.SList' not supported

// UNSUPPORTED : C value 'soup_cookies_to_response' : parameter 'cookies' of type 'GLib.SList' not supported

// UNSUPPORTED : C value 'soup_form_decode' : return type 'GLib.HashTable' not supported

// UNSUPPORTED : C value 'soup_form_decode_multipart' : parameter 'msg' of type 'Message' not supported

// UNSUPPORTED : C value 'soup_form_encode' : parameter '...' has no type

// UNSUPPORTED : C value 'soup_form_encode_datalist' : parameter 'form_data_set' of type 'GLib.Data' not supported

// UNSUPPORTED : C value 'soup_form_encode_hash' : parameter 'form_data_set' of type 'GLib.HashTable' not supported

// UNSUPPORTED : C value 'soup_form_encode_valist' : parameter 'args' of type 'va_list' not supported

// UNSUPPORTED : C value 'soup_form_request_new' : parameter '...' has no type

// UNSUPPORTED : C value 'soup_form_request_new_from_datalist' : parameter 'form_data_set' of type 'GLib.Data' not supported

// UNSUPPORTED : C value 'soup_form_request_new_from_hash' : parameter 'form_data_set' of type 'GLib.HashTable' not supported

// UNSUPPORTED : C value 'soup_form_request_new_from_multipart' : parameter 'multipart' of type 'Multipart' not supported

var getMajorVersionFunction *gi.Function
var getMajorVersionFunction_Once sync.Once

func getMajorVersionFunction_Set() {
	getMajorVersionFunction_Once.Do(func() {
		getMajorVersionFunction = gi.FunctionInvokerNew("Soup", "get_major_version")
	})
}

// GetMajorVersion is a representation of the C type soup_get_major_version.
func GetMajorVersion() uint32 {

	var ret gi.Argument

	getMajorVersionFunction_Set()

	ret = getMajorVersionFunction.Invoke(nil, nil)

	retGo := ret.Uint32()

	return retGo
}

var getMicroVersionFunction *gi.Function
var getMicroVersionFunction_Once sync.Once

func getMicroVersionFunction_Set() {
	getMicroVersionFunction_Once.Do(func() {
		getMicroVersionFunction = gi.FunctionInvokerNew("Soup", "get_micro_version")
	})
}

// GetMicroVersion is a representation of the C type soup_get_micro_version.
func GetMicroVersion() uint32 {

	var ret gi.Argument

	getMicroVersionFunction_Set()

	ret = getMicroVersionFunction.Invoke(nil, nil)

	retGo := ret.Uint32()

	return retGo
}

var getMinorVersionFunction *gi.Function
var getMinorVersionFunction_Once sync.Once

func getMinorVersionFunction_Set() {
	getMinorVersionFunction_Once.Do(func() {
		getMinorVersionFunction = gi.FunctionInvokerNew("Soup", "get_minor_version")
	})
}

// GetMinorVersion is a representation of the C type soup_get_minor_version.
func GetMinorVersion() uint32 {

	var ret gi.Argument

	getMinorVersionFunction_Set()

	ret = getMinorVersionFunction.Invoke(nil, nil)

	retGo := ret.Uint32()

	return retGo
}

// UNSUPPORTED : C value 'soup_header_contains' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'soup_header_free_list' : parameter 'list' of type 'GLib.SList' not supported

// UNSUPPORTED : C value 'soup_header_free_param_list' : parameter 'param_list' of type 'GLib.HashTable' not supported

// UNSUPPORTED : C value 'soup_header_g_string_append_param' : parameter 'string' of type 'GLib.String' not supported

// UNSUPPORTED : C value 'soup_header_g_string_append_param_quoted' : parameter 'string' of type 'GLib.String' not supported

// UNSUPPORTED : C value 'soup_header_parse_list' : return type 'GLib.SList' not supported

// UNSUPPORTED : C value 'soup_header_parse_param_list' : return type 'GLib.HashTable' not supported

// UNSUPPORTED : C value 'soup_header_parse_param_list_strict' : return type 'GLib.HashTable' not supported

// UNSUPPORTED : C value 'soup_header_parse_quality_list' : parameter 'unacceptable' of type 'GLib.SList' not supported

// UNSUPPORTED : C value 'soup_header_parse_semi_param_list' : return type 'GLib.HashTable' not supported

// UNSUPPORTED : C value 'soup_header_parse_semi_param_list_strict' : return type 'GLib.HashTable' not supported

// UNSUPPORTED : C value 'soup_headers_parse' : parameter 'dest' of type 'MessageHeaders' not supported

// UNSUPPORTED : C value 'soup_headers_parse_request' : parameter 'req_headers' of type 'MessageHeaders' not supported

// UNSUPPORTED : C value 'soup_headers_parse_response' : parameter 'headers' of type 'MessageHeaders' not supported

// UNSUPPORTED : C value 'soup_headers_parse_status_line' : parameter 'ver' of type 'HTTPVersion' not supported

// UNSUPPORTED : C value 'soup_http_error_quark' : return type 'GLib.Quark' not supported

// UNSUPPORTED : C value 'soup_message_headers_iter_init' : parameter 'iter' of type 'MessageHeadersIter' not supported

// UNSUPPORTED : C value 'soup_request_error_quark' : return type 'GLib.Quark' not supported

// UNSUPPORTED : C value 'soup_requester_error_quark' : return type 'GLib.Quark' not supported

var statusGetPhraseFunction *gi.Function
var statusGetPhraseFunction_Once sync.Once

func statusGetPhraseFunction_Set() {
	statusGetPhraseFunction_Once.Do(func() {
		statusGetPhraseFunction = gi.FunctionInvokerNew("Soup", "status_get_phrase")
	})
}

// StatusGetPhrase is a representation of the C type soup_status_get_phrase.
func StatusGetPhrase(statusCode uint32) string {
	var inArgs [1]gi.Argument
	inArgs[0].SetUint32(statusCode)

	var ret gi.Argument

	statusGetPhraseFunction_Set()

	ret = statusGetPhraseFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

var statusProxifyFunction *gi.Function
var statusProxifyFunction_Once sync.Once

func statusProxifyFunction_Set() {
	statusProxifyFunction_Once.Do(func() {
		statusProxifyFunction = gi.FunctionInvokerNew("Soup", "status_proxify")
	})
}

// StatusProxify is a representation of the C type soup_status_proxify.
func StatusProxify(statusCode uint32) uint32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetUint32(statusCode)

	var ret gi.Argument

	statusProxifyFunction_Set()

	ret = statusProxifyFunction.Invoke(inArgs[:], nil)

	retGo := ret.Uint32()

	return retGo
}

// UNSUPPORTED : C value 'soup_str_case_equal' : parameter 'v1' of type 'gpointer' not supported

// UNSUPPORTED : C value 'soup_str_case_hash' : parameter 'key' of type 'gpointer' not supported

// UNSUPPORTED : C value 'soup_tld_domain_is_public_suffix' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'soup_tld_error_quark' : return type 'GLib.Quark' not supported

var tldGetBaseDomainFunction *gi.Function
var tldGetBaseDomainFunction_Once sync.Once

func tldGetBaseDomainFunction_Set() {
	tldGetBaseDomainFunction_Once.Do(func() {
		tldGetBaseDomainFunction = gi.FunctionInvokerNew("Soup", "tld_get_base_domain")
	})
}

// TldGetBaseDomain is a representation of the C type soup_tld_get_base_domain.
func TldGetBaseDomain(hostname string) string {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(hostname)

	var ret gi.Argument

	tldGetBaseDomainFunction_Set()

	ret = tldGetBaseDomainFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

var uriDecodeFunction *gi.Function
var uriDecodeFunction_Once sync.Once

func uriDecodeFunction_Set() {
	uriDecodeFunction_Once.Do(func() {
		uriDecodeFunction = gi.FunctionInvokerNew("Soup", "uri_decode")
	})
}

// UriDecode is a representation of the C type soup_uri_decode.
func UriDecode(part string) string {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(part)

	var ret gi.Argument

	uriDecodeFunction_Set()

	ret = uriDecodeFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

var uriEncodeFunction *gi.Function
var uriEncodeFunction_Once sync.Once

func uriEncodeFunction_Set() {
	uriEncodeFunction_Once.Do(func() {
		uriEncodeFunction = gi.FunctionInvokerNew("Soup", "uri_encode")
	})
}

// UriEncode is a representation of the C type soup_uri_encode.
func UriEncode(part string, escapeExtra string) string {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(part)
	inArgs[1].SetString(escapeExtra)

	var ret gi.Argument

	uriEncodeFunction_Set()

	ret = uriEncodeFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

var uriNormalizeFunction *gi.Function
var uriNormalizeFunction_Once sync.Once

func uriNormalizeFunction_Set() {
	uriNormalizeFunction_Once.Do(func() {
		uriNormalizeFunction = gi.FunctionInvokerNew("Soup", "uri_normalize")
	})
}

// UriNormalize is a representation of the C type soup_uri_normalize.
func UriNormalize(part string, unescapeExtra string) string {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(part)
	inArgs[1].SetString(unescapeExtra)

	var ret gi.Argument

	uriNormalizeFunction_Set()

	ret = uriNormalizeFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

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

// UNSUPPORTED : C value 'soup_value_hash_new_with_vals' : parameter '...' has no type

// UNSUPPORTED : C value 'soup_websocket_client_prepare_handshake' : parameter 'msg' of type 'Message' not supported

// UNSUPPORTED : C value 'soup_websocket_client_prepare_handshake_with_extensions' : parameter 'msg' of type 'Message' not supported

// UNSUPPORTED : C value 'soup_websocket_client_verify_handshake' : parameter 'msg' of type 'Message' not supported

// UNSUPPORTED : C value 'soup_websocket_client_verify_handshake_with_extensions' : parameter 'msg' of type 'Message' not supported

// UNSUPPORTED : C value 'soup_websocket_error_get_quark' : return type 'GLib.Quark' not supported

// UNSUPPORTED : C value 'soup_websocket_server_check_handshake' : parameter 'msg' of type 'Message' not supported

// UNSUPPORTED : C value 'soup_websocket_server_check_handshake_with_extensions' : parameter 'msg' of type 'Message' not supported

// UNSUPPORTED : C value 'soup_websocket_server_process_handshake' : parameter 'msg' of type 'Message' not supported

// UNSUPPORTED : C value 'soup_websocket_server_process_handshake_with_extensions' : parameter 'msg' of type 'Message' not supported

// UNSUPPORTED : C value 'soup_xmlrpc_build_fault' : parameter '...' has no type

// UNSUPPORTED : C value 'soup_xmlrpc_build_method_call' : parameter 'params' has no type

// UNSUPPORTED : C value 'soup_xmlrpc_build_method_response' : parameter 'value' of type 'GObject.Value' not supported

// UNSUPPORTED : C value 'soup_xmlrpc_build_request' : parameter 'params' of type 'GLib.Variant' not supported

// UNSUPPORTED : C value 'soup_xmlrpc_build_response' : parameter 'value' of type 'GLib.Variant' not supported

// UNSUPPORTED : C value 'soup_xmlrpc_error_quark' : return type 'GLib.Quark' not supported

// UNSUPPORTED : C value 'soup_xmlrpc_extract_method_call' : parameter '...' has no type

// UNSUPPORTED : C value 'soup_xmlrpc_extract_method_response' : parameter 'error' of type 'GLib.Error' not supported

// UNSUPPORTED : C value 'soup_xmlrpc_fault_quark' : return type 'GLib.Quark' not supported

// UNSUPPORTED : C value 'soup_xmlrpc_message_new' : parameter 'params' of type 'GLib.Variant' not supported

// UNSUPPORTED : C value 'soup_xmlrpc_message_set_fault' : parameter 'msg' of type 'Message' not supported

// UNSUPPORTED : C value 'soup_xmlrpc_message_set_response' : parameter 'msg' of type 'Message' not supported

// UNSUPPORTED : C value 'soup_xmlrpc_parse_method_call' : parameter 'params' of type 'GObject.ValueArray' not supported

// UNSUPPORTED : C value 'soup_xmlrpc_parse_method_response' : parameter 'value' of type 'GObject.Value' not supported

// UNSUPPORTED : C value 'soup_xmlrpc_parse_request' : parameter 'params' of type 'XMLRPCParams' not supported

// UNSUPPORTED : C value 'soup_xmlrpc_parse_response' : return type 'GLib.Variant' not supported

// UNSUPPORTED : C value 'soup_xmlrpc_request_new' : parameter '...' has no type

// UNSUPPORTED : C value 'soup_xmlrpc_set_fault' : parameter 'msg' of type 'Message' not supported

// UNSUPPORTED : C value 'soup_xmlrpc_set_response' : parameter 'msg' of type 'Message' not supported

// UNSUPPORTED : C value 'soup_xmlrpc_variant_get_datetime' : parameter 'variant' of type 'GLib.Variant' not supported

// UNSUPPORTED : C value 'soup_xmlrpc_variant_new_datetime' : parameter 'date' of type 'Date' not supported
