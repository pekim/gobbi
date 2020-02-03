// Code generated - DO NOT EDIT.

package soup

import (
	gi "github.com/pekim/gobbi/internal/cgo/gi"
	glib "github.com/pekim/gobbi/lib/glib"
	gobject "github.com/pekim/gobbi/lib/gobject"
	"sync"
	"unsafe"
)

// UNSUPPORTED : C value 'soup_add_completion' : parameter 'function' of type 'GLib.SourceFunc' not supported

// UNSUPPORTED : C value 'soup_add_idle' : parameter 'function' of type 'GLib.SourceFunc' not supported

// UNSUPPORTED : C value 'soup_add_io_watch' : parameter 'condition' of type 'GLib.IOCondition' not supported

// UNSUPPORTED : C value 'soup_add_timeout' : parameter 'function' of type 'GLib.SourceFunc' not supported

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
	inArgs[1].SetPointer(origin.Native())

	var ret gi.Argument

	err := cookieParseFunction_Set()
	if err == nil {
		ret = cookieParseFunction.Invoke(inArgs[:], nil)
	}

	retGo := CookieNewFromNative(ret.Pointer())

	return retGo
}

var cookiesFreeFunction *gi.Function
var cookiesFreeFunction_Once sync.Once

func cookiesFreeFunction_Set() error {
	var err error
	cookiesFreeFunction_Once.Do(func() {
		cookiesFreeFunction, err = gi.FunctionInvokerNew("Soup", "cookies_free")
	})
	return err
}

// CookiesFree is a representation of the C type soup_cookies_free.
func CookiesFree(cookies *glib.SList) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(cookies.Native())

	err := cookiesFreeFunction_Set()
	if err == nil {
		cookiesFreeFunction.Invoke(inArgs[:], nil)
	}

	return
}

var cookiesFromRequestFunction *gi.Function
var cookiesFromRequestFunction_Once sync.Once

func cookiesFromRequestFunction_Set() error {
	var err error
	cookiesFromRequestFunction_Once.Do(func() {
		cookiesFromRequestFunction, err = gi.FunctionInvokerNew("Soup", "cookies_from_request")
	})
	return err
}

// CookiesFromRequest is a representation of the C type soup_cookies_from_request.
func CookiesFromRequest(msg *Message) *glib.SList {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(msg.Native())

	var ret gi.Argument

	err := cookiesFromRequestFunction_Set()
	if err == nil {
		ret = cookiesFromRequestFunction.Invoke(inArgs[:], nil)
	}

	retGo := glib.SListNewFromNative(ret.Pointer())

	return retGo
}

var cookiesFromResponseFunction *gi.Function
var cookiesFromResponseFunction_Once sync.Once

func cookiesFromResponseFunction_Set() error {
	var err error
	cookiesFromResponseFunction_Once.Do(func() {
		cookiesFromResponseFunction, err = gi.FunctionInvokerNew("Soup", "cookies_from_response")
	})
	return err
}

// CookiesFromResponse is a representation of the C type soup_cookies_from_response.
func CookiesFromResponse(msg *Message) *glib.SList {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(msg.Native())

	var ret gi.Argument

	err := cookiesFromResponseFunction_Set()
	if err == nil {
		ret = cookiesFromResponseFunction.Invoke(inArgs[:], nil)
	}

	retGo := glib.SListNewFromNative(ret.Pointer())

	return retGo
}

var cookiesToCookieHeaderFunction *gi.Function
var cookiesToCookieHeaderFunction_Once sync.Once

func cookiesToCookieHeaderFunction_Set() error {
	var err error
	cookiesToCookieHeaderFunction_Once.Do(func() {
		cookiesToCookieHeaderFunction, err = gi.FunctionInvokerNew("Soup", "cookies_to_cookie_header")
	})
	return err
}

// CookiesToCookieHeader is a representation of the C type soup_cookies_to_cookie_header.
func CookiesToCookieHeader(cookies *glib.SList) string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(cookies.Native())

	var ret gi.Argument

	err := cookiesToCookieHeaderFunction_Set()
	if err == nil {
		ret = cookiesToCookieHeaderFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

var cookiesToRequestFunction *gi.Function
var cookiesToRequestFunction_Once sync.Once

func cookiesToRequestFunction_Set() error {
	var err error
	cookiesToRequestFunction_Once.Do(func() {
		cookiesToRequestFunction, err = gi.FunctionInvokerNew("Soup", "cookies_to_request")
	})
	return err
}

// CookiesToRequest is a representation of the C type soup_cookies_to_request.
func CookiesToRequest(cookies *glib.SList, msg *Message) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(cookies.Native())
	inArgs[1].SetPointer(msg.Native())

	err := cookiesToRequestFunction_Set()
	if err == nil {
		cookiesToRequestFunction.Invoke(inArgs[:], nil)
	}

	return
}

var cookiesToResponseFunction *gi.Function
var cookiesToResponseFunction_Once sync.Once

func cookiesToResponseFunction_Set() error {
	var err error
	cookiesToResponseFunction_Once.Do(func() {
		cookiesToResponseFunction, err = gi.FunctionInvokerNew("Soup", "cookies_to_response")
	})
	return err
}

// CookiesToResponse is a representation of the C type soup_cookies_to_response.
func CookiesToResponse(cookies *glib.SList, msg *Message) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(cookies.Native())
	inArgs[1].SetPointer(msg.Native())

	err := cookiesToResponseFunction_Set()
	if err == nil {
		cookiesToResponseFunction.Invoke(inArgs[:], nil)
	}

	return
}

var formDecodeFunction *gi.Function
var formDecodeFunction_Once sync.Once

func formDecodeFunction_Set() error {
	var err error
	formDecodeFunction_Once.Do(func() {
		formDecodeFunction, err = gi.FunctionInvokerNew("Soup", "form_decode")
	})
	return err
}

// FormDecode is a representation of the C type soup_form_decode.
func FormDecode(encodedForm string) *glib.HashTable {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(encodedForm)

	var ret gi.Argument

	err := formDecodeFunction_Set()
	if err == nil {
		ret = formDecodeFunction.Invoke(inArgs[:], nil)
	}

	retGo := glib.HashTableNewFromNative(ret.Pointer())

	return retGo
}

var formDecodeMultipartFunction *gi.Function
var formDecodeMultipartFunction_Once sync.Once

func formDecodeMultipartFunction_Set() error {
	var err error
	formDecodeMultipartFunction_Once.Do(func() {
		formDecodeMultipartFunction, err = gi.FunctionInvokerNew("Soup", "form_decode_multipart")
	})
	return err
}

// FormDecodeMultipart is a representation of the C type soup_form_decode_multipart.
func FormDecodeMultipart(msg *Message, fileControlName string) (*glib.HashTable, string, string, *Buffer) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(msg.Native())
	inArgs[1].SetString(fileControlName)

	var outArgs [3]gi.Argument
	var ret gi.Argument

	err := formDecodeMultipartFunction_Set()
	if err == nil {
		ret = formDecodeMultipartFunction.Invoke(inArgs[:], outArgs[:])
	}

	retGo := glib.HashTableNewFromNative(ret.Pointer())
	out0 := outArgs[0].String(true)
	out1 := outArgs[1].String(true)
	out2 := BufferNewFromNative(outArgs[2].Pointer())

	return retGo, out0, out1, out2
}

// UNSUPPORTED : C value 'soup_form_encode' : parameter '...' of type 'nil' not supported

var formEncodeDatalistFunction *gi.Function
var formEncodeDatalistFunction_Once sync.Once

func formEncodeDatalistFunction_Set() error {
	var err error
	formEncodeDatalistFunction_Once.Do(func() {
		formEncodeDatalistFunction, err = gi.FunctionInvokerNew("Soup", "form_encode_datalist")
	})
	return err
}

// FormEncodeDatalist is a representation of the C type soup_form_encode_datalist.
func FormEncodeDatalist(formDataSet *glib.Data) string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(formDataSet.Native())

	var ret gi.Argument

	err := formEncodeDatalistFunction_Set()
	if err == nil {
		ret = formEncodeDatalistFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

var formEncodeHashFunction *gi.Function
var formEncodeHashFunction_Once sync.Once

func formEncodeHashFunction_Set() error {
	var err error
	formEncodeHashFunction_Once.Do(func() {
		formEncodeHashFunction, err = gi.FunctionInvokerNew("Soup", "form_encode_hash")
	})
	return err
}

// FormEncodeHash is a representation of the C type soup_form_encode_hash.
func FormEncodeHash(formDataSet *glib.HashTable) string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(formDataSet.Native())

	var ret gi.Argument

	err := formEncodeHashFunction_Set()
	if err == nil {
		ret = formEncodeHashFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

// UNSUPPORTED : C value 'soup_form_encode_valist' : parameter 'args' of type 'va_list' not supported

// UNSUPPORTED : C value 'soup_form_request_new' : parameter '...' of type 'nil' not supported

var formRequestNewFromDatalistFunction *gi.Function
var formRequestNewFromDatalistFunction_Once sync.Once

func formRequestNewFromDatalistFunction_Set() error {
	var err error
	formRequestNewFromDatalistFunction_Once.Do(func() {
		formRequestNewFromDatalistFunction, err = gi.FunctionInvokerNew("Soup", "form_request_new_from_datalist")
	})
	return err
}

// FormRequestNewFromDatalist is a representation of the C type soup_form_request_new_from_datalist.
func FormRequestNewFromDatalist(method string, uri string, formDataSet *glib.Data) *Message {
	var inArgs [3]gi.Argument
	inArgs[0].SetString(method)
	inArgs[1].SetString(uri)
	inArgs[2].SetPointer(formDataSet.Native())

	var ret gi.Argument

	err := formRequestNewFromDatalistFunction_Set()
	if err == nil {
		ret = formRequestNewFromDatalistFunction.Invoke(inArgs[:], nil)
	}

	retGo := MessageNewFromNative(ret.Pointer())

	return retGo
}

var formRequestNewFromHashFunction *gi.Function
var formRequestNewFromHashFunction_Once sync.Once

func formRequestNewFromHashFunction_Set() error {
	var err error
	formRequestNewFromHashFunction_Once.Do(func() {
		formRequestNewFromHashFunction, err = gi.FunctionInvokerNew("Soup", "form_request_new_from_hash")
	})
	return err
}

// FormRequestNewFromHash is a representation of the C type soup_form_request_new_from_hash.
func FormRequestNewFromHash(method string, uri string, formDataSet *glib.HashTable) *Message {
	var inArgs [3]gi.Argument
	inArgs[0].SetString(method)
	inArgs[1].SetString(uri)
	inArgs[2].SetPointer(formDataSet.Native())

	var ret gi.Argument

	err := formRequestNewFromHashFunction_Set()
	if err == nil {
		ret = formRequestNewFromHashFunction.Invoke(inArgs[:], nil)
	}

	retGo := MessageNewFromNative(ret.Pointer())

	return retGo
}

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
	inArgs[1].SetPointer(multipart.Native())

	var ret gi.Argument

	err := formRequestNewFromMultipartFunction_Set()
	if err == nil {
		ret = formRequestNewFromMultipartFunction.Invoke(inArgs[:], nil)
	}

	retGo := MessageNewFromNative(ret.Pointer())

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

var headerFreeListFunction *gi.Function
var headerFreeListFunction_Once sync.Once

func headerFreeListFunction_Set() error {
	var err error
	headerFreeListFunction_Once.Do(func() {
		headerFreeListFunction, err = gi.FunctionInvokerNew("Soup", "header_free_list")
	})
	return err
}

// HeaderFreeList is a representation of the C type soup_header_free_list.
func HeaderFreeList(list *glib.SList) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(list.Native())

	err := headerFreeListFunction_Set()
	if err == nil {
		headerFreeListFunction.Invoke(inArgs[:], nil)
	}

	return
}

var headerFreeParamListFunction *gi.Function
var headerFreeParamListFunction_Once sync.Once

func headerFreeParamListFunction_Set() error {
	var err error
	headerFreeParamListFunction_Once.Do(func() {
		headerFreeParamListFunction, err = gi.FunctionInvokerNew("Soup", "header_free_param_list")
	})
	return err
}

// HeaderFreeParamList is a representation of the C type soup_header_free_param_list.
func HeaderFreeParamList(paramList *glib.HashTable) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(paramList.Native())

	err := headerFreeParamListFunction_Set()
	if err == nil {
		headerFreeParamListFunction.Invoke(inArgs[:], nil)
	}

	return
}

var headerGStringAppendParamFunction *gi.Function
var headerGStringAppendParamFunction_Once sync.Once

func headerGStringAppendParamFunction_Set() error {
	var err error
	headerGStringAppendParamFunction_Once.Do(func() {
		headerGStringAppendParamFunction, err = gi.FunctionInvokerNew("Soup", "header_g_string_append_param")
	})
	return err
}

// HeaderGStringAppendParam is a representation of the C type soup_header_g_string_append_param.
func HeaderGStringAppendParam(string_ *glib.String, name string, value string) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(string_.Native())
	inArgs[1].SetString(name)
	inArgs[2].SetString(value)

	err := headerGStringAppendParamFunction_Set()
	if err == nil {
		headerGStringAppendParamFunction.Invoke(inArgs[:], nil)
	}

	return
}

var headerGStringAppendParamQuotedFunction *gi.Function
var headerGStringAppendParamQuotedFunction_Once sync.Once

func headerGStringAppendParamQuotedFunction_Set() error {
	var err error
	headerGStringAppendParamQuotedFunction_Once.Do(func() {
		headerGStringAppendParamQuotedFunction, err = gi.FunctionInvokerNew("Soup", "header_g_string_append_param_quoted")
	})
	return err
}

// HeaderGStringAppendParamQuoted is a representation of the C type soup_header_g_string_append_param_quoted.
func HeaderGStringAppendParamQuoted(string_ *glib.String, name string, value string) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(string_.Native())
	inArgs[1].SetString(name)
	inArgs[2].SetString(value)

	err := headerGStringAppendParamQuotedFunction_Set()
	if err == nil {
		headerGStringAppendParamQuotedFunction.Invoke(inArgs[:], nil)
	}

	return
}

var headerParseListFunction *gi.Function
var headerParseListFunction_Once sync.Once

func headerParseListFunction_Set() error {
	var err error
	headerParseListFunction_Once.Do(func() {
		headerParseListFunction, err = gi.FunctionInvokerNew("Soup", "header_parse_list")
	})
	return err
}

// HeaderParseList is a representation of the C type soup_header_parse_list.
func HeaderParseList(header string) *glib.SList {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(header)

	var ret gi.Argument

	err := headerParseListFunction_Set()
	if err == nil {
		ret = headerParseListFunction.Invoke(inArgs[:], nil)
	}

	retGo := glib.SListNewFromNative(ret.Pointer())

	return retGo
}

var headerParseParamListFunction *gi.Function
var headerParseParamListFunction_Once sync.Once

func headerParseParamListFunction_Set() error {
	var err error
	headerParseParamListFunction_Once.Do(func() {
		headerParseParamListFunction, err = gi.FunctionInvokerNew("Soup", "header_parse_param_list")
	})
	return err
}

// HeaderParseParamList is a representation of the C type soup_header_parse_param_list.
func HeaderParseParamList(header string) *glib.HashTable {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(header)

	var ret gi.Argument

	err := headerParseParamListFunction_Set()
	if err == nil {
		ret = headerParseParamListFunction.Invoke(inArgs[:], nil)
	}

	retGo := glib.HashTableNewFromNative(ret.Pointer())

	return retGo
}

var headerParseParamListStrictFunction *gi.Function
var headerParseParamListStrictFunction_Once sync.Once

func headerParseParamListStrictFunction_Set() error {
	var err error
	headerParseParamListStrictFunction_Once.Do(func() {
		headerParseParamListStrictFunction, err = gi.FunctionInvokerNew("Soup", "header_parse_param_list_strict")
	})
	return err
}

// HeaderParseParamListStrict is a representation of the C type soup_header_parse_param_list_strict.
func HeaderParseParamListStrict(header string) *glib.HashTable {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(header)

	var ret gi.Argument

	err := headerParseParamListStrictFunction_Set()
	if err == nil {
		ret = headerParseParamListStrictFunction.Invoke(inArgs[:], nil)
	}

	retGo := glib.HashTableNewFromNative(ret.Pointer())

	return retGo
}

var headerParseQualityListFunction *gi.Function
var headerParseQualityListFunction_Once sync.Once

func headerParseQualityListFunction_Set() error {
	var err error
	headerParseQualityListFunction_Once.Do(func() {
		headerParseQualityListFunction, err = gi.FunctionInvokerNew("Soup", "header_parse_quality_list")
	})
	return err
}

// HeaderParseQualityList is a representation of the C type soup_header_parse_quality_list.
func HeaderParseQualityList(header string) (*glib.SList, *glib.SList) {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(header)

	var outArgs [1]gi.Argument
	var ret gi.Argument

	err := headerParseQualityListFunction_Set()
	if err == nil {
		ret = headerParseQualityListFunction.Invoke(inArgs[:], outArgs[:])
	}

	retGo := glib.SListNewFromNative(ret.Pointer())
	out0 := glib.SListNewFromNative(outArgs[0].Pointer())

	return retGo, out0
}

var headerParseSemiParamListFunction *gi.Function
var headerParseSemiParamListFunction_Once sync.Once

func headerParseSemiParamListFunction_Set() error {
	var err error
	headerParseSemiParamListFunction_Once.Do(func() {
		headerParseSemiParamListFunction, err = gi.FunctionInvokerNew("Soup", "header_parse_semi_param_list")
	})
	return err
}

// HeaderParseSemiParamList is a representation of the C type soup_header_parse_semi_param_list.
func HeaderParseSemiParamList(header string) *glib.HashTable {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(header)

	var ret gi.Argument

	err := headerParseSemiParamListFunction_Set()
	if err == nil {
		ret = headerParseSemiParamListFunction.Invoke(inArgs[:], nil)
	}

	retGo := glib.HashTableNewFromNative(ret.Pointer())

	return retGo
}

var headerParseSemiParamListStrictFunction *gi.Function
var headerParseSemiParamListStrictFunction_Once sync.Once

func headerParseSemiParamListStrictFunction_Set() error {
	var err error
	headerParseSemiParamListStrictFunction_Once.Do(func() {
		headerParseSemiParamListStrictFunction, err = gi.FunctionInvokerNew("Soup", "header_parse_semi_param_list_strict")
	})
	return err
}

// HeaderParseSemiParamListStrict is a representation of the C type soup_header_parse_semi_param_list_strict.
func HeaderParseSemiParamListStrict(header string) *glib.HashTable {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(header)

	var ret gi.Argument

	err := headerParseSemiParamListStrictFunction_Set()
	if err == nil {
		ret = headerParseSemiParamListStrictFunction.Invoke(inArgs[:], nil)
	}

	retGo := glib.HashTableNewFromNative(ret.Pointer())

	return retGo
}

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
	inArgs[2].SetPointer(dest.Native())

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
	inArgs[2].SetPointer(reqHeaders.Native())

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
	inArgs[2].SetPointer(headers.Native())

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
	inArgs[0].SetPointer(hdrs.Native())

	var outArgs [1]gi.Argument

	err := messageHeadersIterInitFunction_Set()
	if err == nil {
		messageHeadersIterInitFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := MessageHeadersIterNewFromNative(outArgs[0].Pointer())

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

var strCaseEqualFunction *gi.Function
var strCaseEqualFunction_Once sync.Once

func strCaseEqualFunction_Set() error {
	var err error
	strCaseEqualFunction_Once.Do(func() {
		strCaseEqualFunction, err = gi.FunctionInvokerNew("Soup", "str_case_equal")
	})
	return err
}

// StrCaseEqual is a representation of the C type soup_str_case_equal.
func StrCaseEqual(v1 unsafe.Pointer, v2 unsafe.Pointer) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(v1)
	inArgs[1].SetPointer(v2)

	var ret gi.Argument

	err := strCaseEqualFunction_Set()
	if err == nil {
		ret = strCaseEqualFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var strCaseHashFunction *gi.Function
var strCaseHashFunction_Once sync.Once

func strCaseHashFunction_Set() error {
	var err error
	strCaseHashFunction_Once.Do(func() {
		strCaseHashFunction, err = gi.FunctionInvokerNew("Soup", "str_case_hash")
	})
	return err
}

// StrCaseHash is a representation of the C type soup_str_case_hash.
func StrCaseHash(key unsafe.Pointer) uint32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(key)

	var ret gi.Argument

	err := strCaseHashFunction_Set()
	if err == nil {
		ret = strCaseHashFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo
}

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

// UNSUPPORTED : C value 'soup_value_array_append' : parameter '...' of type 'nil' not supported

// UNSUPPORTED : C value 'soup_value_array_append_vals' : parameter '...' of type 'nil' not supported

// UNSUPPORTED : C value 'soup_value_array_from_args' : parameter 'args' of type 'va_list' not supported

// UNSUPPORTED : C value 'soup_value_array_get_nth' : parameter '...' of type 'nil' not supported

// UNSUPPORTED : C value 'soup_value_array_insert' : parameter '...' of type 'nil' not supported

var valueArrayNewFunction *gi.Function
var valueArrayNewFunction_Once sync.Once

func valueArrayNewFunction_Set() error {
	var err error
	valueArrayNewFunction_Once.Do(func() {
		valueArrayNewFunction, err = gi.FunctionInvokerNew("Soup", "value_array_new")
	})
	return err
}

// ValueArrayNew is a representation of the C type soup_value_array_new.
func ValueArrayNew() *gobject.ValueArray {

	var ret gi.Argument

	err := valueArrayNewFunction_Set()
	if err == nil {
		ret = valueArrayNewFunction.Invoke(nil, nil)
	}

	retGo := gobject.ValueArrayNewFromNative(ret.Pointer())

	return retGo
}

// UNSUPPORTED : C value 'soup_value_array_new_with_vals' : parameter '...' of type 'nil' not supported

// UNSUPPORTED : C value 'soup_value_array_to_args' : parameter 'args' of type 'va_list' not supported

// UNSUPPORTED : C value 'soup_value_hash_insert' : parameter '...' of type 'nil' not supported

// UNSUPPORTED : C value 'soup_value_hash_insert_vals' : parameter '...' of type 'nil' not supported

var valueHashInsertValueFunction *gi.Function
var valueHashInsertValueFunction_Once sync.Once

func valueHashInsertValueFunction_Set() error {
	var err error
	valueHashInsertValueFunction_Once.Do(func() {
		valueHashInsertValueFunction, err = gi.FunctionInvokerNew("Soup", "value_hash_insert_value")
	})
	return err
}

// ValueHashInsertValue is a representation of the C type soup_value_hash_insert_value.
func ValueHashInsertValue(hash *glib.HashTable, key string, value *gobject.Value) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(hash.Native())
	inArgs[1].SetString(key)
	inArgs[2].SetPointer(value.Native())

	err := valueHashInsertValueFunction_Set()
	if err == nil {
		valueHashInsertValueFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'soup_value_hash_lookup' : parameter '...' of type 'nil' not supported

// UNSUPPORTED : C value 'soup_value_hash_lookup_vals' : parameter '...' of type 'nil' not supported

var valueHashNewFunction *gi.Function
var valueHashNewFunction_Once sync.Once

func valueHashNewFunction_Set() error {
	var err error
	valueHashNewFunction_Once.Do(func() {
		valueHashNewFunction, err = gi.FunctionInvokerNew("Soup", "value_hash_new")
	})
	return err
}

// ValueHashNew is a representation of the C type soup_value_hash_new.
func ValueHashNew() *glib.HashTable {

	var ret gi.Argument

	err := valueHashNewFunction_Set()
	if err == nil {
		ret = valueHashNewFunction.Invoke(nil, nil)
	}

	retGo := glib.HashTableNewFromNative(ret.Pointer())

	return retGo
}

// UNSUPPORTED : C value 'soup_value_hash_new_with_vals' : parameter '...' of type 'nil' not supported

var websocketClientPrepareHandshakeFunction *gi.Function
var websocketClientPrepareHandshakeFunction_Once sync.Once

func websocketClientPrepareHandshakeFunction_Set() error {
	var err error
	websocketClientPrepareHandshakeFunction_Once.Do(func() {
		websocketClientPrepareHandshakeFunction, err = gi.FunctionInvokerNew("Soup", "websocket_client_prepare_handshake")
	})
	return err
}

// WebsocketClientPrepareHandshake is a representation of the C type soup_websocket_client_prepare_handshake.
func WebsocketClientPrepareHandshake(msg *Message, origin string, protocols []string) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(msg.Native())
	inArgs[1].SetString(origin)
	inArgs[2].SetStringArray(protocols)

	err := websocketClientPrepareHandshakeFunction_Set()
	if err == nil {
		websocketClientPrepareHandshakeFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'soup_websocket_client_prepare_handshake_with_extensions' : parameter 'supported_extensions' of type 'nil' not supported

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
	inArgs[0].SetPointer(msg.Native())

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

var websocketServerCheckHandshakeFunction *gi.Function
var websocketServerCheckHandshakeFunction_Once sync.Once

func websocketServerCheckHandshakeFunction_Set() error {
	var err error
	websocketServerCheckHandshakeFunction_Once.Do(func() {
		websocketServerCheckHandshakeFunction, err = gi.FunctionInvokerNew("Soup", "websocket_server_check_handshake")
	})
	return err
}

// WebsocketServerCheckHandshake is a representation of the C type soup_websocket_server_check_handshake.
func WebsocketServerCheckHandshake(msg *Message, origin string, protocols []string) bool {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(msg.Native())
	inArgs[1].SetString(origin)
	inArgs[2].SetStringArray(protocols)

	var ret gi.Argument

	err := websocketServerCheckHandshakeFunction_Set()
	if err == nil {
		ret = websocketServerCheckHandshakeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'soup_websocket_server_check_handshake_with_extensions' : parameter 'supported_extensions' of type 'nil' not supported

var websocketServerProcessHandshakeFunction *gi.Function
var websocketServerProcessHandshakeFunction_Once sync.Once

func websocketServerProcessHandshakeFunction_Set() error {
	var err error
	websocketServerProcessHandshakeFunction_Once.Do(func() {
		websocketServerProcessHandshakeFunction, err = gi.FunctionInvokerNew("Soup", "websocket_server_process_handshake")
	})
	return err
}

// WebsocketServerProcessHandshake is a representation of the C type soup_websocket_server_process_handshake.
func WebsocketServerProcessHandshake(msg *Message, expectedOrigin string, protocols []string) bool {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(msg.Native())
	inArgs[1].SetString(expectedOrigin)
	inArgs[2].SetStringArray(protocols)

	var ret gi.Argument

	err := websocketServerProcessHandshakeFunction_Set()
	if err == nil {
		ret = websocketServerProcessHandshakeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'soup_websocket_server_process_handshake_with_extensions' : parameter 'supported_extensions' of type 'nil' not supported

// UNSUPPORTED : C value 'soup_xmlrpc_build_fault' : parameter '...' of type 'nil' not supported

// UNSUPPORTED : C value 'soup_xmlrpc_build_method_call' : parameter 'params' of type 'nil' not supported

var xmlrpcBuildMethodResponseFunction *gi.Function
var xmlrpcBuildMethodResponseFunction_Once sync.Once

func xmlrpcBuildMethodResponseFunction_Set() error {
	var err error
	xmlrpcBuildMethodResponseFunction_Once.Do(func() {
		xmlrpcBuildMethodResponseFunction, err = gi.FunctionInvokerNew("Soup", "xmlrpc_build_method_response")
	})
	return err
}

// XmlrpcBuildMethodResponse is a representation of the C type soup_xmlrpc_build_method_response.
func XmlrpcBuildMethodResponse(value *gobject.Value) string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(value.Native())

	var ret gi.Argument

	err := xmlrpcBuildMethodResponseFunction_Set()
	if err == nil {
		ret = xmlrpcBuildMethodResponseFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

var xmlrpcBuildRequestFunction *gi.Function
var xmlrpcBuildRequestFunction_Once sync.Once

func xmlrpcBuildRequestFunction_Set() error {
	var err error
	xmlrpcBuildRequestFunction_Once.Do(func() {
		xmlrpcBuildRequestFunction, err = gi.FunctionInvokerNew("Soup", "xmlrpc_build_request")
	})
	return err
}

// XmlrpcBuildRequest is a representation of the C type soup_xmlrpc_build_request.
func XmlrpcBuildRequest(methodName string, params *glib.Variant) string {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(methodName)
	inArgs[1].SetPointer(params.Native())

	var ret gi.Argument

	err := xmlrpcBuildRequestFunction_Set()
	if err == nil {
		ret = xmlrpcBuildRequestFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

var xmlrpcBuildResponseFunction *gi.Function
var xmlrpcBuildResponseFunction_Once sync.Once

func xmlrpcBuildResponseFunction_Set() error {
	var err error
	xmlrpcBuildResponseFunction_Once.Do(func() {
		xmlrpcBuildResponseFunction, err = gi.FunctionInvokerNew("Soup", "xmlrpc_build_response")
	})
	return err
}

// XmlrpcBuildResponse is a representation of the C type soup_xmlrpc_build_response.
func XmlrpcBuildResponse(value *glib.Variant) string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(value.Native())

	var ret gi.Argument

	err := xmlrpcBuildResponseFunction_Set()
	if err == nil {
		ret = xmlrpcBuildResponseFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

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

// UNSUPPORTED : C value 'soup_xmlrpc_extract_method_response' : parameter '...' of type 'nil' not supported

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

var xmlrpcMessageNewFunction *gi.Function
var xmlrpcMessageNewFunction_Once sync.Once

func xmlrpcMessageNewFunction_Set() error {
	var err error
	xmlrpcMessageNewFunction_Once.Do(func() {
		xmlrpcMessageNewFunction, err = gi.FunctionInvokerNew("Soup", "xmlrpc_message_new")
	})
	return err
}

// XmlrpcMessageNew is a representation of the C type soup_xmlrpc_message_new.
func XmlrpcMessageNew(uri string, methodName string, params *glib.Variant) *Message {
	var inArgs [3]gi.Argument
	inArgs[0].SetString(uri)
	inArgs[1].SetString(methodName)
	inArgs[2].SetPointer(params.Native())

	var ret gi.Argument

	err := xmlrpcMessageNewFunction_Set()
	if err == nil {
		ret = xmlrpcMessageNewFunction.Invoke(inArgs[:], nil)
	}

	retGo := MessageNewFromNative(ret.Pointer())

	return retGo
}

// UNSUPPORTED : C value 'soup_xmlrpc_message_set_fault' : parameter '...' of type 'nil' not supported

var xmlrpcMessageSetResponseFunction *gi.Function
var xmlrpcMessageSetResponseFunction_Once sync.Once

func xmlrpcMessageSetResponseFunction_Set() error {
	var err error
	xmlrpcMessageSetResponseFunction_Once.Do(func() {
		xmlrpcMessageSetResponseFunction, err = gi.FunctionInvokerNew("Soup", "xmlrpc_message_set_response")
	})
	return err
}

// XmlrpcMessageSetResponse is a representation of the C type soup_xmlrpc_message_set_response.
func XmlrpcMessageSetResponse(msg *Message, value *glib.Variant) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(msg.Native())
	inArgs[1].SetPointer(value.Native())

	var ret gi.Argument

	err := xmlrpcMessageSetResponseFunction_Set()
	if err == nil {
		ret = xmlrpcMessageSetResponseFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var xmlrpcParseMethodCallFunction *gi.Function
var xmlrpcParseMethodCallFunction_Once sync.Once

func xmlrpcParseMethodCallFunction_Set() error {
	var err error
	xmlrpcParseMethodCallFunction_Once.Do(func() {
		xmlrpcParseMethodCallFunction, err = gi.FunctionInvokerNew("Soup", "xmlrpc_parse_method_call")
	})
	return err
}

// XmlrpcParseMethodCall is a representation of the C type soup_xmlrpc_parse_method_call.
func XmlrpcParseMethodCall(methodCall string, length int32) (bool, string, *gobject.ValueArray) {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(methodCall)
	inArgs[1].SetInt32(length)

	var outArgs [2]gi.Argument
	var ret gi.Argument

	err := xmlrpcParseMethodCallFunction_Set()
	if err == nil {
		ret = xmlrpcParseMethodCallFunction.Invoke(inArgs[:], outArgs[:])
	}

	retGo := ret.Boolean()
	out0 := outArgs[0].String(true)
	out1 := gobject.ValueArrayNewFromNative(outArgs[1].Pointer())

	return retGo, out0, out1
}

var xmlrpcParseMethodResponseFunction *gi.Function
var xmlrpcParseMethodResponseFunction_Once sync.Once

func xmlrpcParseMethodResponseFunction_Set() error {
	var err error
	xmlrpcParseMethodResponseFunction_Once.Do(func() {
		xmlrpcParseMethodResponseFunction, err = gi.FunctionInvokerNew("Soup", "xmlrpc_parse_method_response")
	})
	return err
}

// XmlrpcParseMethodResponse is a representation of the C type soup_xmlrpc_parse_method_response.
func XmlrpcParseMethodResponse(methodResponse string, length int32) (bool, *gobject.Value) {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(methodResponse)
	inArgs[1].SetInt32(length)

	var outArgs [1]gi.Argument
	var ret gi.Argument

	err := xmlrpcParseMethodResponseFunction_Set()
	if err == nil {
		ret = xmlrpcParseMethodResponseFunction.Invoke(inArgs[:], outArgs[:])
	}

	retGo := ret.Boolean()
	out0 := gobject.ValueNewFromNative(outArgs[0].Pointer())

	return retGo, out0
}

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
	out0 := XMLRPCParamsNewFromNative(outArgs[0].Pointer())

	return retGo, out0
}

var xmlrpcParseResponseFunction *gi.Function
var xmlrpcParseResponseFunction_Once sync.Once

func xmlrpcParseResponseFunction_Set() error {
	var err error
	xmlrpcParseResponseFunction_Once.Do(func() {
		xmlrpcParseResponseFunction, err = gi.FunctionInvokerNew("Soup", "xmlrpc_parse_response")
	})
	return err
}

// XmlrpcParseResponse is a representation of the C type soup_xmlrpc_parse_response.
func XmlrpcParseResponse(methodResponse string, length int32, signature string) *glib.Variant {
	var inArgs [3]gi.Argument
	inArgs[0].SetString(methodResponse)
	inArgs[1].SetInt32(length)
	inArgs[2].SetString(signature)

	var ret gi.Argument

	err := xmlrpcParseResponseFunction_Set()
	if err == nil {
		ret = xmlrpcParseResponseFunction.Invoke(inArgs[:], nil)
	}

	retGo := glib.VariantNewFromNative(ret.Pointer())

	return retGo
}

// UNSUPPORTED : C value 'soup_xmlrpc_request_new' : parameter '...' of type 'nil' not supported

// UNSUPPORTED : C value 'soup_xmlrpc_set_fault' : parameter '...' of type 'nil' not supported

// UNSUPPORTED : C value 'soup_xmlrpc_set_response' : parameter '...' of type 'nil' not supported

var xmlrpcVariantGetDatetimeFunction *gi.Function
var xmlrpcVariantGetDatetimeFunction_Once sync.Once

func xmlrpcVariantGetDatetimeFunction_Set() error {
	var err error
	xmlrpcVariantGetDatetimeFunction_Once.Do(func() {
		xmlrpcVariantGetDatetimeFunction, err = gi.FunctionInvokerNew("Soup", "xmlrpc_variant_get_datetime")
	})
	return err
}

// XmlrpcVariantGetDatetime is a representation of the C type soup_xmlrpc_variant_get_datetime.
func XmlrpcVariantGetDatetime(variant *glib.Variant) *Date {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(variant.Native())

	var ret gi.Argument

	err := xmlrpcVariantGetDatetimeFunction_Set()
	if err == nil {
		ret = xmlrpcVariantGetDatetimeFunction.Invoke(inArgs[:], nil)
	}

	retGo := DateNewFromNative(ret.Pointer())

	return retGo
}

var xmlrpcVariantNewDatetimeFunction *gi.Function
var xmlrpcVariantNewDatetimeFunction_Once sync.Once

func xmlrpcVariantNewDatetimeFunction_Set() error {
	var err error
	xmlrpcVariantNewDatetimeFunction_Once.Do(func() {
		xmlrpcVariantNewDatetimeFunction, err = gi.FunctionInvokerNew("Soup", "xmlrpc_variant_new_datetime")
	})
	return err
}

// XmlrpcVariantNewDatetime is a representation of the C type soup_xmlrpc_variant_new_datetime.
func XmlrpcVariantNewDatetime(date *Date) *glib.Variant {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(date.Native())

	var ret gi.Argument

	err := xmlrpcVariantNewDatetimeFunction_Set()
	if err == nil {
		ret = xmlrpcVariantNewDatetimeFunction.Invoke(inArgs[:], nil)
	}

	retGo := glib.VariantNewFromNative(ret.Pointer())

	return retGo
}
