// Code generated - DO NOT EDIT.

package soup

import (
	gi "github.com/pekim/gobbi/internal/gi"
	glib "github.com/pekim/gobbi/lib/glib"
	"sync"
)

// UNSUPPORTED : C value 'soup_add_completion' : parameter 'async_context' of type 'GLib.MainContext' not supported

// UNSUPPORTED : C value 'soup_add_idle' : parameter 'async_context' of type 'GLib.MainContext' not supported

// UNSUPPORTED : C value 'soup_add_io_watch' : parameter 'async_context' of type 'GLib.MainContext' not supported

// UNSUPPORTED : C value 'soup_add_timeout' : parameter 'async_context' of type 'GLib.MainContext' not supported

var checkVersionFunction *gi.Function
var checkVersionFunction_Once sync.Once

func checkVersionFunction_Set() error {
	var err error
	checkVersionFunction_Once.Do(func() {
		checkVersionFunction, err = gi.FunctionInvokerNew("Soup", "check_version")
	})
	return err
}

// CheckVersion is a representation of the C type soup_check_version.
func CheckVersion(major uint32, minor uint32, micro uint32) bool {
	var inArgs [3]gi.Argument
	inArgs[0].SetUint32(major)
	inArgs[1].SetUint32(minor)
	inArgs[2].SetUint32(micro)

	var ret gi.Argument

	err := checkVersionFunction_Set()
	if err == nil {
		ret = checkVersionFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var cookieParseFunction *gi.Function
var cookieParseFunction_Once sync.Once

func cookieParseFunction_Set() error {
	var err error
	cookieParseFunction_Once.Do(func() {
		cookieParseFunction, err = gi.FunctionInvokerNew("Soup", "cookie_parse")
	})
	return err
}

// CookieParse is a representation of the C type soup_cookie_parse.
func CookieParse(header string, origin *URI) *Cookie {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(header)
	inArgs[1].SetPointer(origin.Native)

	var ret gi.Argument

	err := cookieParseFunction_Set()
	if err == nil {
		ret = cookieParseFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Cookie{}
	retGo.Native = ret.Pointer()

	return retGo
}

// UNSUPPORTED : C value 'soup_cookies_free' : parameter 'cookies' of type 'GLib.SList' not supported

// UNSUPPORTED : C value 'soup_cookies_from_request' : return type 'GLib.SList' not supported

// UNSUPPORTED : C value 'soup_cookies_from_response' : return type 'GLib.SList' not supported

// UNSUPPORTED : C value 'soup_cookies_to_cookie_header' : parameter 'cookies' of type 'GLib.SList' not supported

// UNSUPPORTED : C value 'soup_cookies_to_request' : parameter 'cookies' of type 'GLib.SList' not supported

// UNSUPPORTED : C value 'soup_cookies_to_response' : parameter 'cookies' of type 'GLib.SList' not supported

// UNSUPPORTED : C value 'soup_form_decode' : return type 'GLib.HashTable' not supported

// UNSUPPORTED : C value 'soup_form_decode_multipart' : return type 'GLib.HashTable' not supported

// UNSUPPORTED : C value 'soup_form_encode' : parameter '...' of type 'nil' not supported

// UNSUPPORTED : C value 'soup_form_encode_datalist' : parameter 'form_data_set' of type 'GLib.Data' not supported

// UNSUPPORTED : C value 'soup_form_encode_hash' : parameter 'form_data_set' of type 'GLib.HashTable' not supported

// UNSUPPORTED : C value 'soup_form_encode_valist' : parameter 'args' of type 'va_list' not supported

// UNSUPPORTED : C value 'soup_form_request_new' : parameter '...' of type 'nil' not supported

// UNSUPPORTED : C value 'soup_form_request_new_from_datalist' : parameter 'form_data_set' of type 'GLib.Data' not supported

// UNSUPPORTED : C value 'soup_form_request_new_from_hash' : parameter 'form_data_set' of type 'GLib.HashTable' not supported

var formRequestNewFromMultipartFunction *gi.Function
var formRequestNewFromMultipartFunction_Once sync.Once

func formRequestNewFromMultipartFunction_Set() error {
	var err error
	formRequestNewFromMultipartFunction_Once.Do(func() {
		formRequestNewFromMultipartFunction, err = gi.FunctionInvokerNew("Soup", "form_request_new_from_multipart")
	})
	return err
}

// FormRequestNewFromMultipart is a representation of the C type soup_form_request_new_from_multipart.
func FormRequestNewFromMultipart(uri string, multipart *Multipart) *Message {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(uri)
	inArgs[1].SetPointer(multipart.Native)

	var ret gi.Argument

	err := formRequestNewFromMultipartFunction_Set()
	if err == nil {
		ret = formRequestNewFromMultipartFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Message{}
	retGo.Native = ret.Pointer()

	return retGo
}

var getMajorVersionFunction *gi.Function
var getMajorVersionFunction_Once sync.Once

func getMajorVersionFunction_Set() error {
	var err error
	getMajorVersionFunction_Once.Do(func() {
		getMajorVersionFunction, err = gi.FunctionInvokerNew("Soup", "get_major_version")
	})
	return err
}

// GetMajorVersion is a representation of the C type soup_get_major_version.
func GetMajorVersion() uint32 {

	var ret gi.Argument

	err := getMajorVersionFunction_Set()
	if err == nil {
		ret = getMajorVersionFunction.Invoke(nil, nil)
	}

	retGo := ret.Uint32()

	return retGo
}

var getMicroVersionFunction *gi.Function
var getMicroVersionFunction_Once sync.Once

func getMicroVersionFunction_Set() error {
	var err error
	getMicroVersionFunction_Once.Do(func() {
		getMicroVersionFunction, err = gi.FunctionInvokerNew("Soup", "get_micro_version")
	})
	return err
}

// GetMicroVersion is a representation of the C type soup_get_micro_version.
func GetMicroVersion() uint32 {

	var ret gi.Argument

	err := getMicroVersionFunction_Set()
	if err == nil {
		ret = getMicroVersionFunction.Invoke(nil, nil)
	}

	retGo := ret.Uint32()

	return retGo
}

var getMinorVersionFunction *gi.Function
var getMinorVersionFunction_Once sync.Once

func getMinorVersionFunction_Set() error {
	var err error
	getMinorVersionFunction_Once.Do(func() {
		getMinorVersionFunction, err = gi.FunctionInvokerNew("Soup", "get_minor_version")
	})
	return err
}

// GetMinorVersion is a representation of the C type soup_get_minor_version.
func GetMinorVersion() uint32 {

	var ret gi.Argument

	err := getMinorVersionFunction_Set()
	if err == nil {
		ret = getMinorVersionFunction.Invoke(nil, nil)
	}

	retGo := ret.Uint32()

	return retGo
}

var headerContainsFunction *gi.Function
var headerContainsFunction_Once sync.Once

func headerContainsFunction_Set() error {
	var err error
	headerContainsFunction_Once.Do(func() {
		headerContainsFunction, err = gi.FunctionInvokerNew("Soup", "header_contains")
	})
	return err
}

// HeaderContains is a representation of the C type soup_header_contains.
func HeaderContains(header string, token string) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(header)
	inArgs[1].SetString(token)

	var ret gi.Argument

	err := headerContainsFunction_Set()
	if err == nil {
		ret = headerContainsFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

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

var headersParseFunction *gi.Function
var headersParseFunction_Once sync.Once

func headersParseFunction_Set() error {
	var err error
	headersParseFunction_Once.Do(func() {
		headersParseFunction, err = gi.FunctionInvokerNew("Soup", "headers_parse")
	})
	return err
}

// HeadersParse is a representation of the C type soup_headers_parse.
func HeadersParse(str string, len int32, dest *MessageHeaders) bool {
	var inArgs [3]gi.Argument
	inArgs[0].SetString(str)
	inArgs[1].SetInt32(len)
	inArgs[2].SetPointer(dest.Native)

	var ret gi.Argument

	err := headersParseFunction_Set()
	if err == nil {
		ret = headersParseFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var headersParseRequestFunction *gi.Function
var headersParseRequestFunction_Once sync.Once

func headersParseRequestFunction_Set() error {
	var err error
	headersParseRequestFunction_Once.Do(func() {
		headersParseRequestFunction, err = gi.FunctionInvokerNew("Soup", "headers_parse_request")
	})
	return err
}

// HeadersParseRequest is a representation of the C type soup_headers_parse_request.
func HeadersParseRequest(str string, len int32, reqHeaders *MessageHeaders) (uint32, string, string, HTTPVersion) {
	var inArgs [3]gi.Argument
	inArgs[0].SetString(str)
	inArgs[1].SetInt32(len)
	inArgs[2].SetPointer(reqHeaders.Native)

	var outArgs [3]gi.Argument
	var ret gi.Argument

	err := headersParseRequestFunction_Set()
	if err == nil {
		ret = headersParseRequestFunction.Invoke(inArgs[:], outArgs[:])
	}

	retGo := ret.Uint32()
	out0 := outArgs[0].String(true)
	out1 := outArgs[1].String(true)
	out2 := HTTPVersion(outArgs[2].Int32())

	return retGo, out0, out1, out2
}

var headersParseResponseFunction *gi.Function
var headersParseResponseFunction_Once sync.Once

func headersParseResponseFunction_Set() error {
	var err error
	headersParseResponseFunction_Once.Do(func() {
		headersParseResponseFunction, err = gi.FunctionInvokerNew("Soup", "headers_parse_response")
	})
	return err
}

// HeadersParseResponse is a representation of the C type soup_headers_parse_response.
func HeadersParseResponse(str string, len int32, headers *MessageHeaders) (bool, HTTPVersion, uint32, string) {
	var inArgs [3]gi.Argument
	inArgs[0].SetString(str)
	inArgs[1].SetInt32(len)
	inArgs[2].SetPointer(headers.Native)

	var outArgs [3]gi.Argument
	var ret gi.Argument

	err := headersParseResponseFunction_Set()
	if err == nil {
		ret = headersParseResponseFunction.Invoke(inArgs[:], outArgs[:])
	}

	retGo := ret.Boolean()
	out0 := HTTPVersion(outArgs[0].Int32())
	out1 := outArgs[1].Uint32()
	out2 := outArgs[2].String(true)

	return retGo, out0, out1, out2
}

var headersParseStatusLineFunction *gi.Function
var headersParseStatusLineFunction_Once sync.Once

func headersParseStatusLineFunction_Set() error {
	var err error
	headersParseStatusLineFunction_Once.Do(func() {
		headersParseStatusLineFunction, err = gi.FunctionInvokerNew("Soup", "headers_parse_status_line")
	})
	return err
}

// HeadersParseStatusLine is a representation of the C type soup_headers_parse_status_line.
func HeadersParseStatusLine(statusLine string) (bool, HTTPVersion, uint32, string) {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(statusLine)

	var outArgs [3]gi.Argument
	var ret gi.Argument

	err := headersParseStatusLineFunction_Set()
	if err == nil {
		ret = headersParseStatusLineFunction.Invoke(inArgs[:], outArgs[:])
	}

	retGo := ret.Boolean()
	out0 := HTTPVersion(outArgs[0].Int32())
	out1 := outArgs[1].Uint32()
	out2 := outArgs[2].String(true)

	return retGo, out0, out1, out2
}

var httpErrorQuarkFunction *gi.Function
var httpErrorQuarkFunction_Once sync.Once

func httpErrorQuarkFunction_Set() error {
	var err error
	httpErrorQuarkFunction_Once.Do(func() {
		httpErrorQuarkFunction, err = gi.FunctionInvokerNew("Soup", "http_error_quark")
	})
	return err
}

// HttpErrorQuark is a representation of the C type soup_http_error_quark.
func HttpErrorQuark() glib.Quark {

	var ret gi.Argument

	err := httpErrorQuarkFunction_Set()
	if err == nil {
		ret = httpErrorQuarkFunction.Invoke(nil, nil)
	}

	retGo := glib.Quark(ret.Uint32())

	return retGo
}

var messageHeadersIterInitFunction *gi.Function
var messageHeadersIterInitFunction_Once sync.Once

func messageHeadersIterInitFunction_Set() error {
	var err error
	messageHeadersIterInitFunction_Once.Do(func() {
		messageHeadersIterInitFunction, err = gi.FunctionInvokerNew("Soup", "message_headers_iter_init")
	})
	return err
}

// MessageHeadersIterInit is a representation of the C type soup_message_headers_iter_init.
func MessageHeadersIterInit(hdrs *MessageHeaders) *MessageHeadersIter {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(hdrs.Native)

	var outArgs [1]gi.Argument

	err := messageHeadersIterInitFunction_Set()
	if err == nil {
		messageHeadersIterInitFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := &MessageHeadersIter{}
	out0.Native = outArgs[0].Pointer()

	return out0
}

var requestErrorQuarkFunction *gi.Function
var requestErrorQuarkFunction_Once sync.Once

func requestErrorQuarkFunction_Set() error {
	var err error
	requestErrorQuarkFunction_Once.Do(func() {
		requestErrorQuarkFunction, err = gi.FunctionInvokerNew("Soup", "request_error_quark")
	})
	return err
}

// RequestErrorQuark is a representation of the C type soup_request_error_quark.
func RequestErrorQuark() glib.Quark {

	var ret gi.Argument

	err := requestErrorQuarkFunction_Set()
	if err == nil {
		ret = requestErrorQuarkFunction.Invoke(nil, nil)
	}

	retGo := glib.Quark(ret.Uint32())

	return retGo
}

var requesterErrorQuarkFunction *gi.Function
var requesterErrorQuarkFunction_Once sync.Once

func requesterErrorQuarkFunction_Set() error {
	var err error
	requesterErrorQuarkFunction_Once.Do(func() {
		requesterErrorQuarkFunction, err = gi.FunctionInvokerNew("Soup", "requester_error_quark")
	})
	return err
}

// RequesterErrorQuark is a representation of the C type soup_requester_error_quark.
func RequesterErrorQuark() glib.Quark {

	var ret gi.Argument

	err := requesterErrorQuarkFunction_Set()
	if err == nil {
		ret = requesterErrorQuarkFunction.Invoke(nil, nil)
	}

	retGo := glib.Quark(ret.Uint32())

	return retGo
}

var statusGetPhraseFunction *gi.Function
var statusGetPhraseFunction_Once sync.Once

func statusGetPhraseFunction_Set() error {
	var err error
	statusGetPhraseFunction_Once.Do(func() {
		statusGetPhraseFunction, err = gi.FunctionInvokerNew("Soup", "status_get_phrase")
	})
	return err
}

// StatusGetPhrase is a representation of the C type soup_status_get_phrase.
func StatusGetPhrase(statusCode uint32) string {
	var inArgs [1]gi.Argument
	inArgs[0].SetUint32(statusCode)

	var ret gi.Argument

	err := statusGetPhraseFunction_Set()
	if err == nil {
		ret = statusGetPhraseFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var statusProxifyFunction *gi.Function
var statusProxifyFunction_Once sync.Once

func statusProxifyFunction_Set() error {
	var err error
	statusProxifyFunction_Once.Do(func() {
		statusProxifyFunction, err = gi.FunctionInvokerNew("Soup", "status_proxify")
	})
	return err
}

// StatusProxify is a representation of the C type soup_status_proxify.
func StatusProxify(statusCode uint32) uint32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetUint32(statusCode)

	var ret gi.Argument

	err := statusProxifyFunction_Set()
	if err == nil {
		ret = statusProxifyFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo
}

// UNSUPPORTED : C value 'soup_str_case_equal' : parameter 'v1' of type 'gpointer' not supported

// UNSUPPORTED : C value 'soup_str_case_hash' : parameter 'key' of type 'gpointer' not supported

var tldDomainIsPublicSuffixFunction *gi.Function
var tldDomainIsPublicSuffixFunction_Once sync.Once

func tldDomainIsPublicSuffixFunction_Set() error {
	var err error
	tldDomainIsPublicSuffixFunction_Once.Do(func() {
		tldDomainIsPublicSuffixFunction, err = gi.FunctionInvokerNew("Soup", "tld_domain_is_public_suffix")
	})
	return err
}

// TldDomainIsPublicSuffix is a representation of the C type soup_tld_domain_is_public_suffix.
func TldDomainIsPublicSuffix(domain string) bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(domain)

	var ret gi.Argument

	err := tldDomainIsPublicSuffixFunction_Set()
	if err == nil {
		ret = tldDomainIsPublicSuffixFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var tldErrorQuarkFunction *gi.Function
var tldErrorQuarkFunction_Once sync.Once

func tldErrorQuarkFunction_Set() error {
	var err error
	tldErrorQuarkFunction_Once.Do(func() {
		tldErrorQuarkFunction, err = gi.FunctionInvokerNew("Soup", "tld_error_quark")
	})
	return err
}

// TldErrorQuark is a representation of the C type soup_tld_error_quark.
func TldErrorQuark() glib.Quark {

	var ret gi.Argument

	err := tldErrorQuarkFunction_Set()
	if err == nil {
		ret = tldErrorQuarkFunction.Invoke(nil, nil)
	}

	retGo := glib.Quark(ret.Uint32())

	return retGo
}

var tldGetBaseDomainFunction *gi.Function
var tldGetBaseDomainFunction_Once sync.Once

func tldGetBaseDomainFunction_Set() error {
	var err error
	tldGetBaseDomainFunction_Once.Do(func() {
		tldGetBaseDomainFunction, err = gi.FunctionInvokerNew("Soup", "tld_get_base_domain")
	})
	return err
}

// TldGetBaseDomain is a representation of the C type soup_tld_get_base_domain.
func TldGetBaseDomain(hostname string) string {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(hostname)

	var ret gi.Argument

	err := tldGetBaseDomainFunction_Set()
	if err == nil {
		ret = tldGetBaseDomainFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var uriDecodeFunction *gi.Function
var uriDecodeFunction_Once sync.Once

func uriDecodeFunction_Set() error {
	var err error
	uriDecodeFunction_Once.Do(func() {
		uriDecodeFunction, err = gi.FunctionInvokerNew("Soup", "uri_decode")
	})
	return err
}

// UriDecode is a representation of the C type soup_uri_decode.
func UriDecode(part string) string {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(part)

	var ret gi.Argument

	err := uriDecodeFunction_Set()
	if err == nil {
		ret = uriDecodeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

var uriEncodeFunction *gi.Function
var uriEncodeFunction_Once sync.Once

func uriEncodeFunction_Set() error {
	var err error
	uriEncodeFunction_Once.Do(func() {
		uriEncodeFunction, err = gi.FunctionInvokerNew("Soup", "uri_encode")
	})
	return err
}

// UriEncode is a representation of the C type soup_uri_encode.
func UriEncode(part string, escapeExtra string) string {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(part)
	inArgs[1].SetString(escapeExtra)

	var ret gi.Argument

	err := uriEncodeFunction_Set()
	if err == nil {
		ret = uriEncodeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

var uriNormalizeFunction *gi.Function
var uriNormalizeFunction_Once sync.Once

func uriNormalizeFunction_Set() error {
	var err error
	uriNormalizeFunction_Once.Do(func() {
		uriNormalizeFunction, err = gi.FunctionInvokerNew("Soup", "uri_normalize")
	})
	return err
}

// UriNormalize is a representation of the C type soup_uri_normalize.
func UriNormalize(part string, unescapeExtra string) string {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(part)
	inArgs[1].SetString(unescapeExtra)

	var ret gi.Argument

	err := uriNormalizeFunction_Set()
	if err == nil {
		ret = uriNormalizeFunction.Invoke(inArgs[:], nil)
	}

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

// UNSUPPORTED : C value 'soup_value_hash_new_with_vals' : parameter '...' of type 'nil' not supported

// UNSUPPORTED : C value 'soup_websocket_client_prepare_handshake' : parameter 'protocols' of type 'nil' not supported

// UNSUPPORTED : C value 'soup_websocket_client_prepare_handshake_with_extensions' : parameter 'protocols' of type 'nil' not supported

var websocketClientVerifyHandshakeFunction *gi.Function
var websocketClientVerifyHandshakeFunction_Once sync.Once

func websocketClientVerifyHandshakeFunction_Set() error {
	var err error
	websocketClientVerifyHandshakeFunction_Once.Do(func() {
		websocketClientVerifyHandshakeFunction, err = gi.FunctionInvokerNew("Soup", "websocket_client_verify_handshake")
	})
	return err
}

// WebsocketClientVerifyHandshake is a representation of the C type soup_websocket_client_verify_handshake.
func WebsocketClientVerifyHandshake(msg *Message) bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(msg.Native)

	var ret gi.Argument

	err := websocketClientVerifyHandshakeFunction_Set()
	if err == nil {
		ret = websocketClientVerifyHandshakeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'soup_websocket_client_verify_handshake_with_extensions' : parameter 'supported_extensions' of type 'nil' not supported

var websocketErrorGetQuarkFunction *gi.Function
var websocketErrorGetQuarkFunction_Once sync.Once

func websocketErrorGetQuarkFunction_Set() error {
	var err error
	websocketErrorGetQuarkFunction_Once.Do(func() {
		websocketErrorGetQuarkFunction, err = gi.FunctionInvokerNew("Soup", "websocket_error_get_quark")
	})
	return err
}

// WebsocketErrorGetQuark is a representation of the C type soup_websocket_error_get_quark.
func WebsocketErrorGetQuark() glib.Quark {

	var ret gi.Argument

	err := websocketErrorGetQuarkFunction_Set()
	if err == nil {
		ret = websocketErrorGetQuarkFunction.Invoke(nil, nil)
	}

	retGo := glib.Quark(ret.Uint32())

	return retGo
}

// UNSUPPORTED : C value 'soup_websocket_server_check_handshake' : parameter 'protocols' of type 'nil' not supported

// UNSUPPORTED : C value 'soup_websocket_server_check_handshake_with_extensions' : parameter 'protocols' of type 'nil' not supported

// UNSUPPORTED : C value 'soup_websocket_server_process_handshake' : parameter 'protocols' of type 'nil' not supported

// UNSUPPORTED : C value 'soup_websocket_server_process_handshake_with_extensions' : parameter 'protocols' of type 'nil' not supported

// UNSUPPORTED : C value 'soup_xmlrpc_build_fault' : parameter '...' of type 'nil' not supported

// UNSUPPORTED : C value 'soup_xmlrpc_build_method_call' : parameter 'params' of type 'nil' not supported

// UNSUPPORTED : C value 'soup_xmlrpc_build_method_response' : parameter 'value' of type 'GObject.Value' not supported

// UNSUPPORTED : C value 'soup_xmlrpc_build_request' : parameter 'params' of type 'GLib.Variant' not supported

// UNSUPPORTED : C value 'soup_xmlrpc_build_response' : parameter 'value' of type 'GLib.Variant' not supported

var xmlrpcErrorQuarkFunction *gi.Function
var xmlrpcErrorQuarkFunction_Once sync.Once

func xmlrpcErrorQuarkFunction_Set() error {
	var err error
	xmlrpcErrorQuarkFunction_Once.Do(func() {
		xmlrpcErrorQuarkFunction, err = gi.FunctionInvokerNew("Soup", "xmlrpc_error_quark")
	})
	return err
}

// XmlrpcErrorQuark is a representation of the C type soup_xmlrpc_error_quark.
func XmlrpcErrorQuark() glib.Quark {

	var ret gi.Argument

	err := xmlrpcErrorQuarkFunction_Set()
	if err == nil {
		ret = xmlrpcErrorQuarkFunction.Invoke(nil, nil)
	}

	retGo := glib.Quark(ret.Uint32())

	return retGo
}

// UNSUPPORTED : C value 'soup_xmlrpc_extract_method_call' : parameter '...' of type 'nil' not supported

// UNSUPPORTED : C value 'soup_xmlrpc_extract_method_response' : parameter 'error' of type 'GLib.Error' not supported

var xmlrpcFaultQuarkFunction *gi.Function
var xmlrpcFaultQuarkFunction_Once sync.Once

func xmlrpcFaultQuarkFunction_Set() error {
	var err error
	xmlrpcFaultQuarkFunction_Once.Do(func() {
		xmlrpcFaultQuarkFunction, err = gi.FunctionInvokerNew("Soup", "xmlrpc_fault_quark")
	})
	return err
}

// XmlrpcFaultQuark is a representation of the C type soup_xmlrpc_fault_quark.
func XmlrpcFaultQuark() glib.Quark {

	var ret gi.Argument

	err := xmlrpcFaultQuarkFunction_Set()
	if err == nil {
		ret = xmlrpcFaultQuarkFunction.Invoke(nil, nil)
	}

	retGo := glib.Quark(ret.Uint32())

	return retGo
}

// UNSUPPORTED : C value 'soup_xmlrpc_message_new' : parameter 'params' of type 'GLib.Variant' not supported

// UNSUPPORTED : C value 'soup_xmlrpc_message_set_fault' : parameter '...' of type 'nil' not supported

// UNSUPPORTED : C value 'soup_xmlrpc_message_set_response' : parameter 'value' of type 'GLib.Variant' not supported

// UNSUPPORTED : C value 'soup_xmlrpc_parse_method_call' : parameter 'params' of type 'GObject.ValueArray' not supported

// UNSUPPORTED : C value 'soup_xmlrpc_parse_method_response' : parameter 'value' of type 'GObject.Value' not supported

var xmlrpcParseRequestFunction *gi.Function
var xmlrpcParseRequestFunction_Once sync.Once

func xmlrpcParseRequestFunction_Set() error {
	var err error
	xmlrpcParseRequestFunction_Once.Do(func() {
		xmlrpcParseRequestFunction, err = gi.FunctionInvokerNew("Soup", "xmlrpc_parse_request")
	})
	return err
}

// XmlrpcParseRequest is a representation of the C type soup_xmlrpc_parse_request.
func XmlrpcParseRequest(methodCall string, length int32) (string, *XMLRPCParams) {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(methodCall)
	inArgs[1].SetInt32(length)

	var outArgs [1]gi.Argument
	var ret gi.Argument

	err := xmlrpcParseRequestFunction_Set()
	if err == nil {
		ret = xmlrpcParseRequestFunction.Invoke(inArgs[:], outArgs[:])
	}

	retGo := ret.String(true)
	out0 := &XMLRPCParams{}
	out0.Native = outArgs[0].Pointer()

	return retGo, out0
}

// UNSUPPORTED : C value 'soup_xmlrpc_parse_response' : return type 'GLib.Variant' not supported

// UNSUPPORTED : C value 'soup_xmlrpc_request_new' : parameter '...' of type 'nil' not supported

// UNSUPPORTED : C value 'soup_xmlrpc_set_fault' : parameter '...' of type 'nil' not supported

// UNSUPPORTED : C value 'soup_xmlrpc_set_response' : parameter 'type' of type 'GType' not supported

// UNSUPPORTED : C value 'soup_xmlrpc_variant_get_datetime' : parameter 'variant' of type 'GLib.Variant' not supported

// UNSUPPORTED : C value 'soup_xmlrpc_variant_new_datetime' : return type 'GLib.Variant' not supported
