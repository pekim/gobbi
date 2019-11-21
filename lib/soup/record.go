// Code generated - DO NOT EDIT.

package soup

import (
	gi "github.com/pekim/gobbi/internal/gi"
	"sync"
)

var addressClassStruct *gi.Struct
var addressClassStruct_Once sync.Once

func addressClassStruct_Set() {
	addressClassStruct_Once.Do(func() {
		addressClassStruct = gi.StructNew("Soup", "AddressClass")
	})
}

type AddressClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value '_libsoup_reserved1' : missing Type
	// UNSUPPORTED : C value '_libsoup_reserved2' : missing Type
	// UNSUPPORTED : C value '_libsoup_reserved3' : missing Type
	// UNSUPPORTED : C value '_libsoup_reserved4' : missing Type
}

var authClassStruct *gi.Struct
var authClassStruct_Once sync.Once

func authClassStruct_Set() {
	authClassStruct_Once.Do(func() {
		authClassStruct = gi.StructNew("Soup", "AuthClass")
	})
}

type AuthClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	SchemeName string
	Strength   uint32
	// UNSUPPORTED : C value 'update' : missing Type
	// UNSUPPORTED : C value 'get_protection_space' : missing Type
	// UNSUPPORTED : C value 'authenticate' : missing Type
	// UNSUPPORTED : C value 'is_authenticated' : missing Type
	// UNSUPPORTED : C value 'get_authorization' : missing Type
	// UNSUPPORTED : C value 'is_ready' : missing Type
	// UNSUPPORTED : C value 'can_authenticate' : missing Type
	// UNSUPPORTED : C value '_libsoup_reserved3' : missing Type
	// UNSUPPORTED : C value '_libsoup_reserved4' : missing Type
}

var authDomainBasicClassStruct *gi.Struct
var authDomainBasicClassStruct_Once sync.Once

func authDomainBasicClassStruct_Set() {
	authDomainBasicClassStruct_Once.Do(func() {
		authDomainBasicClassStruct = gi.StructNew("Soup", "AuthDomainBasicClass")
	})
}

type AuthDomainBasicClass struct {
	native      uintptr
	ParentClass *AuthDomainClass
	// UNSUPPORTED : C value '_libsoup_reserved1' : missing Type
	// UNSUPPORTED : C value '_libsoup_reserved2' : missing Type
	// UNSUPPORTED : C value '_libsoup_reserved3' : missing Type
	// UNSUPPORTED : C value '_libsoup_reserved4' : missing Type
}

var authDomainClassStruct *gi.Struct
var authDomainClassStruct_Once sync.Once

func authDomainClassStruct_Set() {
	authDomainClassStruct_Once.Do(func() {
		authDomainClassStruct = gi.StructNew("Soup", "AuthDomainClass")
	})
}

type AuthDomainClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value 'accepts' : missing Type
	// UNSUPPORTED : C value 'challenge' : missing Type
	// UNSUPPORTED : C value 'check_password' : missing Type
	// UNSUPPORTED : C value '_libsoup_reserved2' : missing Type
	// UNSUPPORTED : C value '_libsoup_reserved3' : missing Type
	// UNSUPPORTED : C value '_libsoup_reserved4' : missing Type
}

var authDomainDigestClassStruct *gi.Struct
var authDomainDigestClassStruct_Once sync.Once

func authDomainDigestClassStruct_Set() {
	authDomainDigestClassStruct_Once.Do(func() {
		authDomainDigestClassStruct = gi.StructNew("Soup", "AuthDomainDigestClass")
	})
}

type AuthDomainDigestClass struct {
	native      uintptr
	ParentClass *AuthDomainClass
	// UNSUPPORTED : C value '_libsoup_reserved1' : missing Type
	// UNSUPPORTED : C value '_libsoup_reserved2' : missing Type
	// UNSUPPORTED : C value '_libsoup_reserved3' : missing Type
	// UNSUPPORTED : C value '_libsoup_reserved4' : missing Type
}

var authManagerClassStruct *gi.Struct
var authManagerClassStruct_Once sync.Once

func authManagerClassStruct_Set() {
	authManagerClassStruct_Once.Do(func() {
		authManagerClassStruct = gi.StructNew("Soup", "AuthManagerClass")
	})
}

type AuthManagerClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value 'authenticate' : missing Type
}

var authManagerPrivateStruct *gi.Struct
var authManagerPrivateStruct_Once sync.Once

func authManagerPrivateStruct_Set() {
	authManagerPrivateStruct_Once.Do(func() {
		authManagerPrivateStruct = gi.StructNew("Soup", "AuthManagerPrivate")
	})
}

type AuthManagerPrivate struct {
	native uintptr
}

var bufferStruct *gi.Struct
var bufferStruct_Once sync.Once

func bufferStruct_Set() {
	bufferStruct_Once.Do(func() {
		bufferStruct = gi.StructNew("Soup", "Buffer")
	})
}

type Buffer struct {
	native uintptr
	// UNSUPPORTED : C value 'data' : no Go type for 'gpointer'
	Length uintptr
}

// UNSUPPORTED : C value 'soup_buffer_new' : parameter 'use' of type 'MemoryUse' not supported

// UNSUPPORTED : C value 'soup_buffer_new_take' : parameter 'data' has no type

// UNSUPPORTED : C value 'soup_buffer_new_with_owner' : parameter 'data' has no type

var bufferCopyFunction *gi.Function
var bufferCopyFunction_Once sync.Once

func bufferCopyFunction_Set() {
	bufferCopyFunction_Once.Do(func() {
		bufferStruct_Set()
		bufferCopyFunction = bufferStruct.InvokerNew("copy")
	})
}

// Copy is a representation of the C type soup_buffer_copy.
func (recv *Buffer) Copy() *Buffer {
	bufferCopyFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := bufferCopyFunction.Invoke(inArgs[:], nil)

	retGo := &Buffer{native: ret.Pointer()}

	return retGo
}

var bufferFreeFunction *gi.Function
var bufferFreeFunction_Once sync.Once

func bufferFreeFunction_Set() {
	bufferFreeFunction_Once.Do(func() {
		bufferStruct_Set()
		bufferFreeFunction = bufferStruct.InvokerNew("free")
	})
}

// Free is a representation of the C type soup_buffer_free.
func (recv *Buffer) Free() {
	bufferFreeFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	bufferFreeFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'soup_buffer_get_as_bytes' : return type 'GLib.Bytes' not supported

// UNSUPPORTED : C value 'soup_buffer_get_data' : parameter 'data' has no type

// UNSUPPORTED : C value 'soup_buffer_get_owner' : return type 'gpointer' not supported

// UNSUPPORTED : C value 'soup_buffer_new_subbuffer' : parameter 'offset' of type 'gsize' not supported

var cacheClassStruct *gi.Struct
var cacheClassStruct_Once sync.Once

func cacheClassStruct_Set() {
	cacheClassStruct_Once.Do(func() {
		cacheClassStruct = gi.StructNew("Soup", "CacheClass")
	})
}

type CacheClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value 'get_cacheability' : missing Type
	// UNSUPPORTED : C value '_libsoup_reserved1' : missing Type
	// UNSUPPORTED : C value '_libsoup_reserved2' : missing Type
	// UNSUPPORTED : C value '_libsoup_reserved3' : missing Type
}

var cachePrivateStruct *gi.Struct
var cachePrivateStruct_Once sync.Once

func cachePrivateStruct_Set() {
	cachePrivateStruct_Once.Do(func() {
		cachePrivateStruct = gi.StructNew("Soup", "CachePrivate")
	})
}

type CachePrivate struct {
	native uintptr
}

var clientContextStruct *gi.Struct
var clientContextStruct_Once sync.Once

func clientContextStruct_Set() {
	clientContextStruct_Once.Do(func() {
		clientContextStruct = gi.StructNew("Soup", "ClientContext")
	})
}

type ClientContext struct {
	native uintptr
}

// UNSUPPORTED : C value 'soup_client_context_get_address' : return type 'Address' not supported

// UNSUPPORTED : C value 'soup_client_context_get_auth_domain' : return type 'AuthDomain' not supported

var clientContextGetAuthUserFunction *gi.Function
var clientContextGetAuthUserFunction_Once sync.Once

func clientContextGetAuthUserFunction_Set() {
	clientContextGetAuthUserFunction_Once.Do(func() {
		clientContextStruct_Set()
		clientContextGetAuthUserFunction = clientContextStruct.InvokerNew("get_auth_user")
	})
}

// GetAuthUser is a representation of the C type soup_client_context_get_auth_user.
func (recv *ClientContext) GetAuthUser() string {
	clientContextGetAuthUserFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := clientContextGetAuthUserFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'soup_client_context_get_gsocket' : return type 'Gio.Socket' not supported

var clientContextGetHostFunction *gi.Function
var clientContextGetHostFunction_Once sync.Once

func clientContextGetHostFunction_Set() {
	clientContextGetHostFunction_Once.Do(func() {
		clientContextStruct_Set()
		clientContextGetHostFunction = clientContextStruct.InvokerNew("get_host")
	})
}

// GetHost is a representation of the C type soup_client_context_get_host.
func (recv *ClientContext) GetHost() string {
	clientContextGetHostFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := clientContextGetHostFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'soup_client_context_get_local_address' : return type 'Gio.SocketAddress' not supported

// UNSUPPORTED : C value 'soup_client_context_get_remote_address' : return type 'Gio.SocketAddress' not supported

// UNSUPPORTED : C value 'soup_client_context_get_socket' : return type 'Socket' not supported

// UNSUPPORTED : C value 'soup_client_context_steal_connection' : return type 'Gio.IOStream' not supported

var connectionStruct *gi.Struct
var connectionStruct_Once sync.Once

func connectionStruct_Set() {
	connectionStruct_Once.Do(func() {
		connectionStruct = gi.StructNew("Soup", "Connection")
	})
}

type Connection struct {
	native uintptr
}

var contentDecoderClassStruct *gi.Struct
var contentDecoderClassStruct_Once sync.Once

func contentDecoderClassStruct_Set() {
	contentDecoderClassStruct_Once.Do(func() {
		contentDecoderClassStruct = gi.StructNew("Soup", "ContentDecoderClass")
	})
}

type ContentDecoderClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value '_libsoup_reserved1' : missing Type
	// UNSUPPORTED : C value '_libsoup_reserved2' : missing Type
	// UNSUPPORTED : C value '_libsoup_reserved3' : missing Type
	// UNSUPPORTED : C value '_libsoup_reserved4' : missing Type
	// UNSUPPORTED : C value '_libsoup_reserved5' : missing Type
}

var contentDecoderPrivateStruct *gi.Struct
var contentDecoderPrivateStruct_Once sync.Once

func contentDecoderPrivateStruct_Set() {
	contentDecoderPrivateStruct_Once.Do(func() {
		contentDecoderPrivateStruct = gi.StructNew("Soup", "ContentDecoderPrivate")
	})
}

type ContentDecoderPrivate struct {
	native uintptr
}

var contentSnifferClassStruct *gi.Struct
var contentSnifferClassStruct_Once sync.Once

func contentSnifferClassStruct_Set() {
	contentSnifferClassStruct_Once.Do(func() {
		contentSnifferClassStruct = gi.StructNew("Soup", "ContentSnifferClass")
	})
}

type ContentSnifferClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value 'sniff' : missing Type
	// UNSUPPORTED : C value 'get_buffer_size' : missing Type
	// UNSUPPORTED : C value '_libsoup_reserved1' : missing Type
	// UNSUPPORTED : C value '_libsoup_reserved2' : missing Type
	// UNSUPPORTED : C value '_libsoup_reserved3' : missing Type
	// UNSUPPORTED : C value '_libsoup_reserved4' : missing Type
	// UNSUPPORTED : C value '_libsoup_reserved5' : missing Type
}

var contentSnifferPrivateStruct *gi.Struct
var contentSnifferPrivateStruct_Once sync.Once

func contentSnifferPrivateStruct_Set() {
	contentSnifferPrivateStruct_Once.Do(func() {
		contentSnifferPrivateStruct = gi.StructNew("Soup", "ContentSnifferPrivate")
	})
}

type ContentSnifferPrivate struct {
	native uintptr
}

var cookieStruct *gi.Struct
var cookieStruct_Once sync.Once

func cookieStruct_Set() {
	cookieStruct_Once.Do(func() {
		cookieStruct = gi.StructNew("Soup", "Cookie")
	})
}

type Cookie struct {
	native  uintptr
	Name    string
	Value   string
	Domain  string
	Path    string
	Expires *Date
	// UNSUPPORTED : C value 'secure' : no Go type for 'gboolean'
	// UNSUPPORTED : C value 'http_only' : no Go type for 'gboolean'
}

var cookieNewFunction *gi.Function
var cookieNewFunction_Once sync.Once

func cookieNewFunction_Set() {
	cookieNewFunction_Once.Do(func() {
		cookieStruct_Set()
		cookieNewFunction = cookieStruct.InvokerNew("new")
	})
}

// CookieNew is a representation of the C type soup_cookie_new.
func CookieNew(name string, value string, domain string, path string, maxAge int32) *Cookie {
	cookieNewFunction_Set()

	var inArgs [5]gi.Argument
	inArgs[0].SetString(name)
	inArgs[1].SetString(value)
	inArgs[2].SetString(domain)
	inArgs[3].SetString(path)
	inArgs[4].SetInt32(maxAge)

	ret := cookieNewFunction.Invoke(inArgs[:], nil)

	retGo := &Cookie{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'soup_cookie_applies_to_uri' : parameter 'uri' of type 'URI' not supported

var cookieCopyFunction *gi.Function
var cookieCopyFunction_Once sync.Once

func cookieCopyFunction_Set() {
	cookieCopyFunction_Once.Do(func() {
		cookieStruct_Set()
		cookieCopyFunction = cookieStruct.InvokerNew("copy")
	})
}

// Copy is a representation of the C type soup_cookie_copy.
func (recv *Cookie) Copy() *Cookie {
	cookieCopyFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := cookieCopyFunction.Invoke(inArgs[:], nil)

	retGo := &Cookie{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'soup_cookie_domain_matches' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'soup_cookie_equal' : parameter 'cookie2' of type 'Cookie' not supported

var cookieFreeFunction *gi.Function
var cookieFreeFunction_Once sync.Once

func cookieFreeFunction_Set() {
	cookieFreeFunction_Once.Do(func() {
		cookieStruct_Set()
		cookieFreeFunction = cookieStruct.InvokerNew("free")
	})
}

// Free is a representation of the C type soup_cookie_free.
func (recv *Cookie) Free() {
	cookieFreeFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	cookieFreeFunction.Invoke(inArgs[:], nil)

}

var cookieGetDomainFunction *gi.Function
var cookieGetDomainFunction_Once sync.Once

func cookieGetDomainFunction_Set() {
	cookieGetDomainFunction_Once.Do(func() {
		cookieStruct_Set()
		cookieGetDomainFunction = cookieStruct.InvokerNew("get_domain")
	})
}

// GetDomain is a representation of the C type soup_cookie_get_domain.
func (recv *Cookie) GetDomain() string {
	cookieGetDomainFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := cookieGetDomainFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

var cookieGetExpiresFunction *gi.Function
var cookieGetExpiresFunction_Once sync.Once

func cookieGetExpiresFunction_Set() {
	cookieGetExpiresFunction_Once.Do(func() {
		cookieStruct_Set()
		cookieGetExpiresFunction = cookieStruct.InvokerNew("get_expires")
	})
}

// GetExpires is a representation of the C type soup_cookie_get_expires.
func (recv *Cookie) GetExpires() *Date {
	cookieGetExpiresFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := cookieGetExpiresFunction.Invoke(inArgs[:], nil)

	retGo := &Date{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'soup_cookie_get_http_only' : return type 'gboolean' not supported

var cookieGetNameFunction *gi.Function
var cookieGetNameFunction_Once sync.Once

func cookieGetNameFunction_Set() {
	cookieGetNameFunction_Once.Do(func() {
		cookieStruct_Set()
		cookieGetNameFunction = cookieStruct.InvokerNew("get_name")
	})
}

// GetName is a representation of the C type soup_cookie_get_name.
func (recv *Cookie) GetName() string {
	cookieGetNameFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := cookieGetNameFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

var cookieGetPathFunction *gi.Function
var cookieGetPathFunction_Once sync.Once

func cookieGetPathFunction_Set() {
	cookieGetPathFunction_Once.Do(func() {
		cookieStruct_Set()
		cookieGetPathFunction = cookieStruct.InvokerNew("get_path")
	})
}

// GetPath is a representation of the C type soup_cookie_get_path.
func (recv *Cookie) GetPath() string {
	cookieGetPathFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := cookieGetPathFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'soup_cookie_get_secure' : return type 'gboolean' not supported

var cookieGetValueFunction *gi.Function
var cookieGetValueFunction_Once sync.Once

func cookieGetValueFunction_Set() {
	cookieGetValueFunction_Once.Do(func() {
		cookieStruct_Set()
		cookieGetValueFunction = cookieStruct.InvokerNew("get_value")
	})
}

// GetValue is a representation of the C type soup_cookie_get_value.
func (recv *Cookie) GetValue() string {
	cookieGetValueFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := cookieGetValueFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

var cookieSetDomainFunction *gi.Function
var cookieSetDomainFunction_Once sync.Once

func cookieSetDomainFunction_Set() {
	cookieSetDomainFunction_Once.Do(func() {
		cookieStruct_Set()
		cookieSetDomainFunction = cookieStruct.InvokerNew("set_domain")
	})
}

// SetDomain is a representation of the C type soup_cookie_set_domain.
func (recv *Cookie) SetDomain(domain string) {
	cookieSetDomainFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(domain)

	cookieSetDomainFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'soup_cookie_set_expires' : parameter 'expires' of type 'Date' not supported

// UNSUPPORTED : C value 'soup_cookie_set_http_only' : parameter 'http_only' of type 'gboolean' not supported

var cookieSetMaxAgeFunction *gi.Function
var cookieSetMaxAgeFunction_Once sync.Once

func cookieSetMaxAgeFunction_Set() {
	cookieSetMaxAgeFunction_Once.Do(func() {
		cookieStruct_Set()
		cookieSetMaxAgeFunction = cookieStruct.InvokerNew("set_max_age")
	})
}

// SetMaxAge is a representation of the C type soup_cookie_set_max_age.
func (recv *Cookie) SetMaxAge(maxAge int32) {
	cookieSetMaxAgeFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(maxAge)

	cookieSetMaxAgeFunction.Invoke(inArgs[:], nil)

}

var cookieSetNameFunction *gi.Function
var cookieSetNameFunction_Once sync.Once

func cookieSetNameFunction_Set() {
	cookieSetNameFunction_Once.Do(func() {
		cookieStruct_Set()
		cookieSetNameFunction = cookieStruct.InvokerNew("set_name")
	})
}

// SetName is a representation of the C type soup_cookie_set_name.
func (recv *Cookie) SetName(name string) {
	cookieSetNameFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(name)

	cookieSetNameFunction.Invoke(inArgs[:], nil)

}

var cookieSetPathFunction *gi.Function
var cookieSetPathFunction_Once sync.Once

func cookieSetPathFunction_Set() {
	cookieSetPathFunction_Once.Do(func() {
		cookieStruct_Set()
		cookieSetPathFunction = cookieStruct.InvokerNew("set_path")
	})
}

// SetPath is a representation of the C type soup_cookie_set_path.
func (recv *Cookie) SetPath(path string) {
	cookieSetPathFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(path)

	cookieSetPathFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'soup_cookie_set_secure' : parameter 'secure' of type 'gboolean' not supported

var cookieSetValueFunction *gi.Function
var cookieSetValueFunction_Once sync.Once

func cookieSetValueFunction_Set() {
	cookieSetValueFunction_Once.Do(func() {
		cookieStruct_Set()
		cookieSetValueFunction = cookieStruct.InvokerNew("set_value")
	})
}

// SetValue is a representation of the C type soup_cookie_set_value.
func (recv *Cookie) SetValue(value string) {
	cookieSetValueFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(value)

	cookieSetValueFunction.Invoke(inArgs[:], nil)

}

var cookieToCookieHeaderFunction *gi.Function
var cookieToCookieHeaderFunction_Once sync.Once

func cookieToCookieHeaderFunction_Set() {
	cookieToCookieHeaderFunction_Once.Do(func() {
		cookieStruct_Set()
		cookieToCookieHeaderFunction = cookieStruct.InvokerNew("to_cookie_header")
	})
}

// ToCookieHeader is a representation of the C type soup_cookie_to_cookie_header.
func (recv *Cookie) ToCookieHeader() string {
	cookieToCookieHeaderFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := cookieToCookieHeaderFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

var cookieToSetCookieHeaderFunction *gi.Function
var cookieToSetCookieHeaderFunction_Once sync.Once

func cookieToSetCookieHeaderFunction_Set() {
	cookieToSetCookieHeaderFunction_Once.Do(func() {
		cookieStruct_Set()
		cookieToSetCookieHeaderFunction = cookieStruct.InvokerNew("to_set_cookie_header")
	})
}

// ToSetCookieHeader is a representation of the C type soup_cookie_to_set_cookie_header.
func (recv *Cookie) ToSetCookieHeader() string {
	cookieToSetCookieHeaderFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := cookieToSetCookieHeaderFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

var cookieJarClassStruct *gi.Struct
var cookieJarClassStruct_Once sync.Once

func cookieJarClassStruct_Set() {
	cookieJarClassStruct_Once.Do(func() {
		cookieJarClassStruct = gi.StructNew("Soup", "CookieJarClass")
	})
}

type CookieJarClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value 'save' : missing Type
	// UNSUPPORTED : C value 'is_persistent' : missing Type
	// UNSUPPORTED : C value 'changed' : missing Type
	// UNSUPPORTED : C value '_libsoup_reserved1' : missing Type
	// UNSUPPORTED : C value '_libsoup_reserved2' : missing Type
}

var cookieJarDBClassStruct *gi.Struct
var cookieJarDBClassStruct_Once sync.Once

func cookieJarDBClassStruct_Set() {
	cookieJarDBClassStruct_Once.Do(func() {
		cookieJarDBClassStruct = gi.StructNew("Soup", "CookieJarDBClass")
	})
}

type CookieJarDBClass struct {
	native      uintptr
	ParentClass *CookieJarClass
	// UNSUPPORTED : C value '_libsoup_reserved1' : missing Type
	// UNSUPPORTED : C value '_libsoup_reserved2' : missing Type
	// UNSUPPORTED : C value '_libsoup_reserved3' : missing Type
	// UNSUPPORTED : C value '_libsoup_reserved4' : missing Type
}

var cookieJarTextClassStruct *gi.Struct
var cookieJarTextClassStruct_Once sync.Once

func cookieJarTextClassStruct_Set() {
	cookieJarTextClassStruct_Once.Do(func() {
		cookieJarTextClassStruct = gi.StructNew("Soup", "CookieJarTextClass")
	})
}

type CookieJarTextClass struct {
	native      uintptr
	ParentClass *CookieJarClass
	// UNSUPPORTED : C value '_libsoup_reserved1' : missing Type
	// UNSUPPORTED : C value '_libsoup_reserved2' : missing Type
	// UNSUPPORTED : C value '_libsoup_reserved3' : missing Type
	// UNSUPPORTED : C value '_libsoup_reserved4' : missing Type
}

var dateStruct *gi.Struct
var dateStruct_Once sync.Once

func dateStruct_Set() {
	dateStruct_Once.Do(func() {
		dateStruct = gi.StructNew("Soup", "Date")
	})
}

type Date struct {
	native uintptr
	Year   int32
	Month  int32
	Day    int32
	Hour   int32
	Minute int32
	Second int32
	// UNSUPPORTED : C value 'utc' : no Go type for 'gboolean'
	Offset int32
}

var dateNewFunction *gi.Function
var dateNewFunction_Once sync.Once

func dateNewFunction_Set() {
	dateNewFunction_Once.Do(func() {
		dateStruct_Set()
		dateNewFunction = dateStruct.InvokerNew("new")
	})
}

// DateNew is a representation of the C type soup_date_new.
func DateNew(year int32, month int32, day int32, hour int32, minute int32, second int32) *Date {
	dateNewFunction_Set()

	var inArgs [6]gi.Argument
	inArgs[0].SetInt32(year)
	inArgs[1].SetInt32(month)
	inArgs[2].SetInt32(day)
	inArgs[3].SetInt32(hour)
	inArgs[4].SetInt32(minute)
	inArgs[5].SetInt32(second)

	ret := dateNewFunction.Invoke(inArgs[:], nil)

	retGo := &Date{native: ret.Pointer()}

	return retGo
}

var dateNewFromNowFunction *gi.Function
var dateNewFromNowFunction_Once sync.Once

func dateNewFromNowFunction_Set() {
	dateNewFromNowFunction_Once.Do(func() {
		dateStruct_Set()
		dateNewFromNowFunction = dateStruct.InvokerNew("new_from_now")
	})
}

// DateNewFromNow is a representation of the C type soup_date_new_from_now.
func DateNewFromNow(offsetSeconds int32) *Date {
	dateNewFromNowFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetInt32(offsetSeconds)

	ret := dateNewFromNowFunction.Invoke(inArgs[:], nil)

	retGo := &Date{native: ret.Pointer()}

	return retGo
}

var dateNewFromStringFunction *gi.Function
var dateNewFromStringFunction_Once sync.Once

func dateNewFromStringFunction_Set() {
	dateNewFromStringFunction_Once.Do(func() {
		dateStruct_Set()
		dateNewFromStringFunction = dateStruct.InvokerNew("new_from_string")
	})
}

// DateNewFromString is a representation of the C type soup_date_new_from_string.
func DateNewFromString(dateString string) *Date {
	dateNewFromStringFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetString(dateString)

	ret := dateNewFromStringFunction.Invoke(inArgs[:], nil)

	retGo := &Date{native: ret.Pointer()}

	return retGo
}

var dateNewFromTimeTFunction *gi.Function
var dateNewFromTimeTFunction_Once sync.Once

func dateNewFromTimeTFunction_Set() {
	dateNewFromTimeTFunction_Once.Do(func() {
		dateStruct_Set()
		dateNewFromTimeTFunction = dateStruct.InvokerNew("new_from_time_t")
	})
}

// DateNewFromTimeT is a representation of the C type soup_date_new_from_time_t.
func DateNewFromTimeT(when int64) *Date {
	dateNewFromTimeTFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetInt64(when)

	ret := dateNewFromTimeTFunction.Invoke(inArgs[:], nil)

	retGo := &Date{native: ret.Pointer()}

	return retGo
}

var dateCopyFunction *gi.Function
var dateCopyFunction_Once sync.Once

func dateCopyFunction_Set() {
	dateCopyFunction_Once.Do(func() {
		dateStruct_Set()
		dateCopyFunction = dateStruct.InvokerNew("copy")
	})
}

// Copy is a representation of the C type soup_date_copy.
func (recv *Date) Copy() *Date {
	dateCopyFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := dateCopyFunction.Invoke(inArgs[:], nil)

	retGo := &Date{native: ret.Pointer()}

	return retGo
}

var dateFreeFunction *gi.Function
var dateFreeFunction_Once sync.Once

func dateFreeFunction_Set() {
	dateFreeFunction_Once.Do(func() {
		dateStruct_Set()
		dateFreeFunction = dateStruct.InvokerNew("free")
	})
}

// Free is a representation of the C type soup_date_free.
func (recv *Date) Free() {
	dateFreeFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	dateFreeFunction.Invoke(inArgs[:], nil)

}

var dateGetDayFunction *gi.Function
var dateGetDayFunction_Once sync.Once

func dateGetDayFunction_Set() {
	dateGetDayFunction_Once.Do(func() {
		dateStruct_Set()
		dateGetDayFunction = dateStruct.InvokerNew("get_day")
	})
}

// GetDay is a representation of the C type soup_date_get_day.
func (recv *Date) GetDay() int32 {
	dateGetDayFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := dateGetDayFunction.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var dateGetHourFunction *gi.Function
var dateGetHourFunction_Once sync.Once

func dateGetHourFunction_Set() {
	dateGetHourFunction_Once.Do(func() {
		dateStruct_Set()
		dateGetHourFunction = dateStruct.InvokerNew("get_hour")
	})
}

// GetHour is a representation of the C type soup_date_get_hour.
func (recv *Date) GetHour() int32 {
	dateGetHourFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := dateGetHourFunction.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var dateGetMinuteFunction *gi.Function
var dateGetMinuteFunction_Once sync.Once

func dateGetMinuteFunction_Set() {
	dateGetMinuteFunction_Once.Do(func() {
		dateStruct_Set()
		dateGetMinuteFunction = dateStruct.InvokerNew("get_minute")
	})
}

// GetMinute is a representation of the C type soup_date_get_minute.
func (recv *Date) GetMinute() int32 {
	dateGetMinuteFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := dateGetMinuteFunction.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var dateGetMonthFunction *gi.Function
var dateGetMonthFunction_Once sync.Once

func dateGetMonthFunction_Set() {
	dateGetMonthFunction_Once.Do(func() {
		dateStruct_Set()
		dateGetMonthFunction = dateStruct.InvokerNew("get_month")
	})
}

// GetMonth is a representation of the C type soup_date_get_month.
func (recv *Date) GetMonth() int32 {
	dateGetMonthFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := dateGetMonthFunction.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var dateGetOffsetFunction *gi.Function
var dateGetOffsetFunction_Once sync.Once

func dateGetOffsetFunction_Set() {
	dateGetOffsetFunction_Once.Do(func() {
		dateStruct_Set()
		dateGetOffsetFunction = dateStruct.InvokerNew("get_offset")
	})
}

// GetOffset is a representation of the C type soup_date_get_offset.
func (recv *Date) GetOffset() int32 {
	dateGetOffsetFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := dateGetOffsetFunction.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var dateGetSecondFunction *gi.Function
var dateGetSecondFunction_Once sync.Once

func dateGetSecondFunction_Set() {
	dateGetSecondFunction_Once.Do(func() {
		dateStruct_Set()
		dateGetSecondFunction = dateStruct.InvokerNew("get_second")
	})
}

// GetSecond is a representation of the C type soup_date_get_second.
func (recv *Date) GetSecond() int32 {
	dateGetSecondFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := dateGetSecondFunction.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var dateGetUtcFunction *gi.Function
var dateGetUtcFunction_Once sync.Once

func dateGetUtcFunction_Set() {
	dateGetUtcFunction_Once.Do(func() {
		dateStruct_Set()
		dateGetUtcFunction = dateStruct.InvokerNew("get_utc")
	})
}

// GetUtc is a representation of the C type soup_date_get_utc.
func (recv *Date) GetUtc() int32 {
	dateGetUtcFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := dateGetUtcFunction.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var dateGetYearFunction *gi.Function
var dateGetYearFunction_Once sync.Once

func dateGetYearFunction_Set() {
	dateGetYearFunction_Once.Do(func() {
		dateStruct_Set()
		dateGetYearFunction = dateStruct.InvokerNew("get_year")
	})
}

// GetYear is a representation of the C type soup_date_get_year.
func (recv *Date) GetYear() int32 {
	dateGetYearFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := dateGetYearFunction.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

// UNSUPPORTED : C value 'soup_date_is_past' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'soup_date_to_string' : parameter 'format' of type 'DateFormat' not supported

var dateToTimeTFunction *gi.Function
var dateToTimeTFunction_Once sync.Once

func dateToTimeTFunction_Set() {
	dateToTimeTFunction_Once.Do(func() {
		dateStruct_Set()
		dateToTimeTFunction = dateStruct.InvokerNew("to_time_t")
	})
}

// ToTimeT is a representation of the C type soup_date_to_time_t.
func (recv *Date) ToTimeT() int64 {
	dateToTimeTFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := dateToTimeTFunction.Invoke(inArgs[:], nil)

	retGo := ret.Int64()

	return retGo
}

// UNSUPPORTED : C value 'soup_date_to_timeval' : parameter 'time' of type 'GLib.TimeVal' not supported

var hSTSEnforcerClassStruct *gi.Struct
var hSTSEnforcerClassStruct_Once sync.Once

func hSTSEnforcerClassStruct_Set() {
	hSTSEnforcerClassStruct_Once.Do(func() {
		hSTSEnforcerClassStruct = gi.StructNew("Soup", "HSTSEnforcerClass")
	})
}

type HSTSEnforcerClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value 'is_persistent' : missing Type
	// UNSUPPORTED : C value 'has_valid_policy' : missing Type
	// UNSUPPORTED : C value 'changed' : missing Type
	// UNSUPPORTED : C value 'hsts_enforced' : missing Type
	// UNSUPPORTED : C value '_libsoup_reserved1' : missing Type
	// UNSUPPORTED : C value '_libsoup_reserved2' : missing Type
	// UNSUPPORTED : C value '_libsoup_reserved3' : missing Type
	// UNSUPPORTED : C value '_libsoup_reserved4' : missing Type
}

var hSTSEnforcerDBClassStruct *gi.Struct
var hSTSEnforcerDBClassStruct_Once sync.Once

func hSTSEnforcerDBClassStruct_Set() {
	hSTSEnforcerDBClassStruct_Once.Do(func() {
		hSTSEnforcerDBClassStruct = gi.StructNew("Soup", "HSTSEnforcerDBClass")
	})
}

type HSTSEnforcerDBClass struct {
	native      uintptr
	ParentClass *HSTSEnforcerClass
	// UNSUPPORTED : C value '_libsoup_reserved1' : missing Type
	// UNSUPPORTED : C value '_libsoup_reserved2' : missing Type
	// UNSUPPORTED : C value '_libsoup_reserved3' : missing Type
	// UNSUPPORTED : C value '_libsoup_reserved4' : missing Type
}

var hSTSEnforcerDBPrivateStruct *gi.Struct
var hSTSEnforcerDBPrivateStruct_Once sync.Once

func hSTSEnforcerDBPrivateStruct_Set() {
	hSTSEnforcerDBPrivateStruct_Once.Do(func() {
		hSTSEnforcerDBPrivateStruct = gi.StructNew("Soup", "HSTSEnforcerDBPrivate")
	})
}

type HSTSEnforcerDBPrivate struct {
	native uintptr
}

var hSTSEnforcerPrivateStruct *gi.Struct
var hSTSEnforcerPrivateStruct_Once sync.Once

func hSTSEnforcerPrivateStruct_Set() {
	hSTSEnforcerPrivateStruct_Once.Do(func() {
		hSTSEnforcerPrivateStruct = gi.StructNew("Soup", "HSTSEnforcerPrivate")
	})
}

type HSTSEnforcerPrivate struct {
	native uintptr
}

var hSTSPolicyStruct *gi.Struct
var hSTSPolicyStruct_Once sync.Once

func hSTSPolicyStruct_Set() {
	hSTSPolicyStruct_Once.Do(func() {
		hSTSPolicyStruct = gi.StructNew("Soup", "HSTSPolicy")
	})
}

type HSTSPolicy struct {
	native  uintptr
	Domain  string
	MaxAge  uint64
	Expires *Date
	// UNSUPPORTED : C value 'include_subdomains' : no Go type for 'gboolean'
}

// UNSUPPORTED : C value 'soup_hsts_policy_new' : parameter 'include_subdomains' of type 'gboolean' not supported

// UNSUPPORTED : C value 'soup_hsts_policy_new_from_response' : parameter 'msg' of type 'Message' not supported

// UNSUPPORTED : C value 'soup_hsts_policy_new_full' : parameter 'expires' of type 'Date' not supported

// UNSUPPORTED : C value 'soup_hsts_policy_new_session_policy' : parameter 'include_subdomains' of type 'gboolean' not supported

var hSTSPolicyCopyFunction *gi.Function
var hSTSPolicyCopyFunction_Once sync.Once

func hSTSPolicyCopyFunction_Set() {
	hSTSPolicyCopyFunction_Once.Do(func() {
		hSTSPolicyStruct_Set()
		hSTSPolicyCopyFunction = hSTSPolicyStruct.InvokerNew("copy")
	})
}

// Copy is a representation of the C type soup_hsts_policy_copy.
func (recv *HSTSPolicy) Copy() *HSTSPolicy {
	hSTSPolicyCopyFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := hSTSPolicyCopyFunction.Invoke(inArgs[:], nil)

	retGo := &HSTSPolicy{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'soup_hsts_policy_equal' : parameter 'policy2' of type 'HSTSPolicy' not supported

var hSTSPolicyFreeFunction *gi.Function
var hSTSPolicyFreeFunction_Once sync.Once

func hSTSPolicyFreeFunction_Set() {
	hSTSPolicyFreeFunction_Once.Do(func() {
		hSTSPolicyStruct_Set()
		hSTSPolicyFreeFunction = hSTSPolicyStruct.InvokerNew("free")
	})
}

// Free is a representation of the C type soup_hsts_policy_free.
func (recv *HSTSPolicy) Free() {
	hSTSPolicyFreeFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	hSTSPolicyFreeFunction.Invoke(inArgs[:], nil)

}

var hSTSPolicyGetDomainFunction *gi.Function
var hSTSPolicyGetDomainFunction_Once sync.Once

func hSTSPolicyGetDomainFunction_Set() {
	hSTSPolicyGetDomainFunction_Once.Do(func() {
		hSTSPolicyStruct_Set()
		hSTSPolicyGetDomainFunction = hSTSPolicyStruct.InvokerNew("get_domain")
	})
}

// GetDomain is a representation of the C type soup_hsts_policy_get_domain.
func (recv *HSTSPolicy) GetDomain() string {
	hSTSPolicyGetDomainFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := hSTSPolicyGetDomainFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'soup_hsts_policy_includes_subdomains' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'soup_hsts_policy_is_expired' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'soup_hsts_policy_is_session_policy' : return type 'gboolean' not supported

var loggerClassStruct *gi.Struct
var loggerClassStruct_Once sync.Once

func loggerClassStruct_Set() {
	loggerClassStruct_Once.Do(func() {
		loggerClassStruct = gi.StructNew("Soup", "LoggerClass")
	})
}

type LoggerClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value '_libsoup_reserved1' : missing Type
	// UNSUPPORTED : C value '_libsoup_reserved2' : missing Type
	// UNSUPPORTED : C value '_libsoup_reserved3' : missing Type
	// UNSUPPORTED : C value '_libsoup_reserved4' : missing Type
}

var messageBodyStruct *gi.Struct
var messageBodyStruct_Once sync.Once

func messageBodyStruct_Set() {
	messageBodyStruct_Once.Do(func() {
		messageBodyStruct = gi.StructNew("Soup", "MessageBody")
	})
}

type MessageBody struct {
	native uintptr
	Data   string
	Length int64
}

var messageBodyNewFunction *gi.Function
var messageBodyNewFunction_Once sync.Once

func messageBodyNewFunction_Set() {
	messageBodyNewFunction_Once.Do(func() {
		messageBodyStruct_Set()
		messageBodyNewFunction = messageBodyStruct.InvokerNew("new")
	})
}

// MessageBodyNew is a representation of the C type soup_message_body_new.
func MessageBodyNew() *MessageBody {
	messageBodyNewFunction_Set()

	ret := messageBodyNewFunction.Invoke(nil, nil)

	retGo := &MessageBody{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'soup_message_body_append' : parameter 'use' of type 'MemoryUse' not supported

// UNSUPPORTED : C value 'soup_message_body_append_buffer' : parameter 'buffer' of type 'Buffer' not supported

// UNSUPPORTED : C value 'soup_message_body_append_take' : parameter 'data' has no type

var messageBodyCompleteFunction *gi.Function
var messageBodyCompleteFunction_Once sync.Once

func messageBodyCompleteFunction_Set() {
	messageBodyCompleteFunction_Once.Do(func() {
		messageBodyStruct_Set()
		messageBodyCompleteFunction = messageBodyStruct.InvokerNew("complete")
	})
}

// Complete is a representation of the C type soup_message_body_complete.
func (recv *MessageBody) Complete() {
	messageBodyCompleteFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	messageBodyCompleteFunction.Invoke(inArgs[:], nil)

}

var messageBodyFlattenFunction *gi.Function
var messageBodyFlattenFunction_Once sync.Once

func messageBodyFlattenFunction_Set() {
	messageBodyFlattenFunction_Once.Do(func() {
		messageBodyStruct_Set()
		messageBodyFlattenFunction = messageBodyStruct.InvokerNew("flatten")
	})
}

// Flatten is a representation of the C type soup_message_body_flatten.
func (recv *MessageBody) Flatten() *Buffer {
	messageBodyFlattenFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := messageBodyFlattenFunction.Invoke(inArgs[:], nil)

	retGo := &Buffer{native: ret.Pointer()}

	return retGo
}

var messageBodyFreeFunction *gi.Function
var messageBodyFreeFunction_Once sync.Once

func messageBodyFreeFunction_Set() {
	messageBodyFreeFunction_Once.Do(func() {
		messageBodyStruct_Set()
		messageBodyFreeFunction = messageBodyStruct.InvokerNew("free")
	})
}

// Free is a representation of the C type soup_message_body_free.
func (recv *MessageBody) Free() {
	messageBodyFreeFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	messageBodyFreeFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'soup_message_body_get_accumulate' : return type 'gboolean' not supported

var messageBodyGetChunkFunction *gi.Function
var messageBodyGetChunkFunction_Once sync.Once

func messageBodyGetChunkFunction_Set() {
	messageBodyGetChunkFunction_Once.Do(func() {
		messageBodyStruct_Set()
		messageBodyGetChunkFunction = messageBodyStruct.InvokerNew("get_chunk")
	})
}

// GetChunk is a representation of the C type soup_message_body_get_chunk.
func (recv *MessageBody) GetChunk(offset int64) *Buffer {
	messageBodyGetChunkFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt64(offset)

	ret := messageBodyGetChunkFunction.Invoke(inArgs[:], nil)

	retGo := &Buffer{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'soup_message_body_got_chunk' : parameter 'chunk' of type 'Buffer' not supported

// UNSUPPORTED : C value 'soup_message_body_set_accumulate' : parameter 'accumulate' of type 'gboolean' not supported

var messageBodyTruncateFunction *gi.Function
var messageBodyTruncateFunction_Once sync.Once

func messageBodyTruncateFunction_Set() {
	messageBodyTruncateFunction_Once.Do(func() {
		messageBodyStruct_Set()
		messageBodyTruncateFunction = messageBodyStruct.InvokerNew("truncate")
	})
}

// Truncate is a representation of the C type soup_message_body_truncate.
func (recv *MessageBody) Truncate() {
	messageBodyTruncateFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	messageBodyTruncateFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'soup_message_body_wrote_chunk' : parameter 'chunk' of type 'Buffer' not supported

var messageClassStruct *gi.Struct
var messageClassStruct_Once sync.Once

func messageClassStruct_Set() {
	messageClassStruct_Once.Do(func() {
		messageClassStruct = gi.StructNew("Soup", "MessageClass")
	})
}

type MessageClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value 'wrote_informational' : missing Type
	// UNSUPPORTED : C value 'wrote_headers' : missing Type
	// UNSUPPORTED : C value 'wrote_chunk' : missing Type
	// UNSUPPORTED : C value 'wrote_body' : missing Type
	// UNSUPPORTED : C value 'got_informational' : missing Type
	// UNSUPPORTED : C value 'got_headers' : missing Type
	// UNSUPPORTED : C value 'got_chunk' : missing Type
	// UNSUPPORTED : C value 'got_body' : missing Type
	// UNSUPPORTED : C value 'restarted' : missing Type
	// UNSUPPORTED : C value 'finished' : missing Type
	// UNSUPPORTED : C value 'starting' : missing Type
	// UNSUPPORTED : C value '_libsoup_reserved1' : missing Type
	// UNSUPPORTED : C value '_libsoup_reserved2' : missing Type
	// UNSUPPORTED : C value '_libsoup_reserved3' : missing Type
}

var messageHeadersStruct *gi.Struct
var messageHeadersStruct_Once sync.Once

func messageHeadersStruct_Set() {
	messageHeadersStruct_Once.Do(func() {
		messageHeadersStruct = gi.StructNew("Soup", "MessageHeaders")
	})
}

type MessageHeaders struct {
	native uintptr
}

// UNSUPPORTED : C value 'soup_message_headers_new' : parameter 'type' of type 'MessageHeadersType' not supported

var messageHeadersAppendFunction *gi.Function
var messageHeadersAppendFunction_Once sync.Once

func messageHeadersAppendFunction_Set() {
	messageHeadersAppendFunction_Once.Do(func() {
		messageHeadersStruct_Set()
		messageHeadersAppendFunction = messageHeadersStruct.InvokerNew("append")
	})
}

// Append is a representation of the C type soup_message_headers_append.
func (recv *MessageHeaders) Append(name string, value string) {
	messageHeadersAppendFunction_Set()

	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(name)
	inArgs[2].SetString(value)

	messageHeadersAppendFunction.Invoke(inArgs[:], nil)

}

var messageHeadersCleanConnectionHeadersFunction *gi.Function
var messageHeadersCleanConnectionHeadersFunction_Once sync.Once

func messageHeadersCleanConnectionHeadersFunction_Set() {
	messageHeadersCleanConnectionHeadersFunction_Once.Do(func() {
		messageHeadersStruct_Set()
		messageHeadersCleanConnectionHeadersFunction = messageHeadersStruct.InvokerNew("clean_connection_headers")
	})
}

// CleanConnectionHeaders is a representation of the C type soup_message_headers_clean_connection_headers.
func (recv *MessageHeaders) CleanConnectionHeaders() {
	messageHeadersCleanConnectionHeadersFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	messageHeadersCleanConnectionHeadersFunction.Invoke(inArgs[:], nil)

}

var messageHeadersClearFunction *gi.Function
var messageHeadersClearFunction_Once sync.Once

func messageHeadersClearFunction_Set() {
	messageHeadersClearFunction_Once.Do(func() {
		messageHeadersStruct_Set()
		messageHeadersClearFunction = messageHeadersStruct.InvokerNew("clear")
	})
}

// Clear is a representation of the C type soup_message_headers_clear.
func (recv *MessageHeaders) Clear() {
	messageHeadersClearFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	messageHeadersClearFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'soup_message_headers_foreach' : parameter 'func' of type 'MessageHeadersForeachFunc' not supported

var messageHeadersFreeFunction *gi.Function
var messageHeadersFreeFunction_Once sync.Once

func messageHeadersFreeFunction_Set() {
	messageHeadersFreeFunction_Once.Do(func() {
		messageHeadersStruct_Set()
		messageHeadersFreeFunction = messageHeadersStruct.InvokerNew("free")
	})
}

// Free is a representation of the C type soup_message_headers_free.
func (recv *MessageHeaders) Free() {
	messageHeadersFreeFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	messageHeadersFreeFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'soup_message_headers_free_ranges' : parameter 'ranges' of type 'Range' not supported

var messageHeadersGetFunction *gi.Function
var messageHeadersGetFunction_Once sync.Once

func messageHeadersGetFunction_Set() {
	messageHeadersGetFunction_Once.Do(func() {
		messageHeadersStruct_Set()
		messageHeadersGetFunction = messageHeadersStruct.InvokerNew("get")
	})
}

// Get is a representation of the C type soup_message_headers_get.
func (recv *MessageHeaders) Get(name string) string {
	messageHeadersGetFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(name)

	ret := messageHeadersGetFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'soup_message_headers_get_content_disposition' : parameter 'params' of type 'GLib.HashTable' not supported

var messageHeadersGetContentLengthFunction *gi.Function
var messageHeadersGetContentLengthFunction_Once sync.Once

func messageHeadersGetContentLengthFunction_Set() {
	messageHeadersGetContentLengthFunction_Once.Do(func() {
		messageHeadersStruct_Set()
		messageHeadersGetContentLengthFunction = messageHeadersStruct.InvokerNew("get_content_length")
	})
}

// GetContentLength is a representation of the C type soup_message_headers_get_content_length.
func (recv *MessageHeaders) GetContentLength() int64 {
	messageHeadersGetContentLengthFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := messageHeadersGetContentLengthFunction.Invoke(inArgs[:], nil)

	retGo := ret.Int64()

	return retGo
}

// UNSUPPORTED : C value 'soup_message_headers_get_content_range' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'soup_message_headers_get_content_type' : parameter 'params' of type 'GLib.HashTable' not supported

// UNSUPPORTED : C value 'soup_message_headers_get_encoding' : return type 'Encoding' not supported

// UNSUPPORTED : C value 'soup_message_headers_get_expectations' : return type 'Expectation' not supported

// UNSUPPORTED : C value 'soup_message_headers_get_headers_type' : return type 'MessageHeadersType' not supported

var messageHeadersGetListFunction *gi.Function
var messageHeadersGetListFunction_Once sync.Once

func messageHeadersGetListFunction_Set() {
	messageHeadersGetListFunction_Once.Do(func() {
		messageHeadersStruct_Set()
		messageHeadersGetListFunction = messageHeadersStruct.InvokerNew("get_list")
	})
}

// GetList is a representation of the C type soup_message_headers_get_list.
func (recv *MessageHeaders) GetList(name string) string {
	messageHeadersGetListFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(name)

	ret := messageHeadersGetListFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

var messageHeadersGetOneFunction *gi.Function
var messageHeadersGetOneFunction_Once sync.Once

func messageHeadersGetOneFunction_Set() {
	messageHeadersGetOneFunction_Once.Do(func() {
		messageHeadersStruct_Set()
		messageHeadersGetOneFunction = messageHeadersStruct.InvokerNew("get_one")
	})
}

// GetOne is a representation of the C type soup_message_headers_get_one.
func (recv *MessageHeaders) GetOne(name string) string {
	messageHeadersGetOneFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(name)

	ret := messageHeadersGetOneFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'soup_message_headers_get_ranges' : parameter 'ranges' has no type

// UNSUPPORTED : C value 'soup_message_headers_header_contains' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'soup_message_headers_header_equals' : return type 'gboolean' not supported

var messageHeadersRemoveFunction *gi.Function
var messageHeadersRemoveFunction_Once sync.Once

func messageHeadersRemoveFunction_Set() {
	messageHeadersRemoveFunction_Once.Do(func() {
		messageHeadersStruct_Set()
		messageHeadersRemoveFunction = messageHeadersStruct.InvokerNew("remove")
	})
}

// Remove is a representation of the C type soup_message_headers_remove.
func (recv *MessageHeaders) Remove(name string) {
	messageHeadersRemoveFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(name)

	messageHeadersRemoveFunction.Invoke(inArgs[:], nil)

}

var messageHeadersReplaceFunction *gi.Function
var messageHeadersReplaceFunction_Once sync.Once

func messageHeadersReplaceFunction_Set() {
	messageHeadersReplaceFunction_Once.Do(func() {
		messageHeadersStruct_Set()
		messageHeadersReplaceFunction = messageHeadersStruct.InvokerNew("replace")
	})
}

// Replace is a representation of the C type soup_message_headers_replace.
func (recv *MessageHeaders) Replace(name string, value string) {
	messageHeadersReplaceFunction_Set()

	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(name)
	inArgs[2].SetString(value)

	messageHeadersReplaceFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'soup_message_headers_set_content_disposition' : parameter 'params' of type 'GLib.HashTable' not supported

var messageHeadersSetContentLengthFunction *gi.Function
var messageHeadersSetContentLengthFunction_Once sync.Once

func messageHeadersSetContentLengthFunction_Set() {
	messageHeadersSetContentLengthFunction_Once.Do(func() {
		messageHeadersStruct_Set()
		messageHeadersSetContentLengthFunction = messageHeadersStruct.InvokerNew("set_content_length")
	})
}

// SetContentLength is a representation of the C type soup_message_headers_set_content_length.
func (recv *MessageHeaders) SetContentLength(contentLength int64) {
	messageHeadersSetContentLengthFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt64(contentLength)

	messageHeadersSetContentLengthFunction.Invoke(inArgs[:], nil)

}

var messageHeadersSetContentRangeFunction *gi.Function
var messageHeadersSetContentRangeFunction_Once sync.Once

func messageHeadersSetContentRangeFunction_Set() {
	messageHeadersSetContentRangeFunction_Once.Do(func() {
		messageHeadersStruct_Set()
		messageHeadersSetContentRangeFunction = messageHeadersStruct.InvokerNew("set_content_range")
	})
}

// SetContentRange is a representation of the C type soup_message_headers_set_content_range.
func (recv *MessageHeaders) SetContentRange(start int64, end int64, totalLength int64) {
	messageHeadersSetContentRangeFunction_Set()

	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt64(start)
	inArgs[2].SetInt64(end)
	inArgs[3].SetInt64(totalLength)

	messageHeadersSetContentRangeFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'soup_message_headers_set_content_type' : parameter 'params' of type 'GLib.HashTable' not supported

// UNSUPPORTED : C value 'soup_message_headers_set_encoding' : parameter 'encoding' of type 'Encoding' not supported

// UNSUPPORTED : C value 'soup_message_headers_set_expectations' : parameter 'expectations' of type 'Expectation' not supported

var messageHeadersSetRangeFunction *gi.Function
var messageHeadersSetRangeFunction_Once sync.Once

func messageHeadersSetRangeFunction_Set() {
	messageHeadersSetRangeFunction_Once.Do(func() {
		messageHeadersStruct_Set()
		messageHeadersSetRangeFunction = messageHeadersStruct.InvokerNew("set_range")
	})
}

// SetRange is a representation of the C type soup_message_headers_set_range.
func (recv *MessageHeaders) SetRange(start int64, end int64) {
	messageHeadersSetRangeFunction_Set()

	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt64(start)
	inArgs[2].SetInt64(end)

	messageHeadersSetRangeFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'soup_message_headers_set_ranges' : parameter 'ranges' of type 'Range' not supported

var messageHeadersIterStruct *gi.Struct
var messageHeadersIterStruct_Once sync.Once

func messageHeadersIterStruct_Set() {
	messageHeadersIterStruct_Once.Do(func() {
		messageHeadersIterStruct = gi.StructNew("Soup", "MessageHeadersIter")
	})
}

type MessageHeadersIter struct {
	native uintptr
}

// UNSUPPORTED : C value 'soup_message_headers_iter_next' : return type 'gboolean' not supported

var messageQueueStruct *gi.Struct
var messageQueueStruct_Once sync.Once

func messageQueueStruct_Set() {
	messageQueueStruct_Once.Do(func() {
		messageQueueStruct = gi.StructNew("Soup", "MessageQueue")
	})
}

type MessageQueue struct {
	native uintptr
}

var messageQueueItemStruct *gi.Struct
var messageQueueItemStruct_Once sync.Once

func messageQueueItemStruct_Set() {
	messageQueueItemStruct_Once.Do(func() {
		messageQueueItemStruct = gi.StructNew("Soup", "MessageQueueItem")
	})
}

type MessageQueueItem struct {
	native uintptr
}

var multipartStruct *gi.Struct
var multipartStruct_Once sync.Once

func multipartStruct_Set() {
	multipartStruct_Once.Do(func() {
		multipartStruct = gi.StructNew("Soup", "Multipart")
	})
}

type Multipart struct {
	native uintptr
}

var multipartNewFunction *gi.Function
var multipartNewFunction_Once sync.Once

func multipartNewFunction_Set() {
	multipartNewFunction_Once.Do(func() {
		multipartStruct_Set()
		multipartNewFunction = multipartStruct.InvokerNew("new")
	})
}

// MultipartNew is a representation of the C type soup_multipart_new.
func MultipartNew(mimeType string) *Multipart {
	multipartNewFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetString(mimeType)

	ret := multipartNewFunction.Invoke(inArgs[:], nil)

	retGo := &Multipart{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'soup_multipart_new_from_message' : parameter 'headers' of type 'MessageHeaders' not supported

// UNSUPPORTED : C value 'soup_multipart_append_form_file' : parameter 'body' of type 'Buffer' not supported

var multipartAppendFormStringFunction *gi.Function
var multipartAppendFormStringFunction_Once sync.Once

func multipartAppendFormStringFunction_Set() {
	multipartAppendFormStringFunction_Once.Do(func() {
		multipartStruct_Set()
		multipartAppendFormStringFunction = multipartStruct.InvokerNew("append_form_string")
	})
}

// AppendFormString is a representation of the C type soup_multipart_append_form_string.
func (recv *Multipart) AppendFormString(controlName string, data string) {
	multipartAppendFormStringFunction_Set()

	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(controlName)
	inArgs[2].SetString(data)

	multipartAppendFormStringFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'soup_multipart_append_part' : parameter 'headers' of type 'MessageHeaders' not supported

var multipartFreeFunction *gi.Function
var multipartFreeFunction_Once sync.Once

func multipartFreeFunction_Set() {
	multipartFreeFunction_Once.Do(func() {
		multipartStruct_Set()
		multipartFreeFunction = multipartStruct.InvokerNew("free")
	})
}

// Free is a representation of the C type soup_multipart_free.
func (recv *Multipart) Free() {
	multipartFreeFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	multipartFreeFunction.Invoke(inArgs[:], nil)

}

var multipartGetLengthFunction *gi.Function
var multipartGetLengthFunction_Once sync.Once

func multipartGetLengthFunction_Set() {
	multipartGetLengthFunction_Once.Do(func() {
		multipartStruct_Set()
		multipartGetLengthFunction = multipartStruct.InvokerNew("get_length")
	})
}

// GetLength is a representation of the C type soup_multipart_get_length.
func (recv *Multipart) GetLength() int32 {
	multipartGetLengthFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := multipartGetLengthFunction.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

// UNSUPPORTED : C value 'soup_multipart_get_part' : parameter 'headers' of type 'MessageHeaders' not supported

// UNSUPPORTED : C value 'soup_multipart_to_message' : parameter 'dest_headers' of type 'MessageHeaders' not supported

var multipartInputStreamClassStruct *gi.Struct
var multipartInputStreamClassStruct_Once sync.Once

func multipartInputStreamClassStruct_Set() {
	multipartInputStreamClassStruct_Once.Do(func() {
		multipartInputStreamClassStruct = gi.StructNew("Soup", "MultipartInputStreamClass")
	})
}

type MultipartInputStreamClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'Gio.FilterInputStreamClass'
}

var multipartInputStreamPrivateStruct *gi.Struct
var multipartInputStreamPrivateStruct_Once sync.Once

func multipartInputStreamPrivateStruct_Set() {
	multipartInputStreamPrivateStruct_Once.Do(func() {
		multipartInputStreamPrivateStruct = gi.StructNew("Soup", "MultipartInputStreamPrivate")
	})
}

type MultipartInputStreamPrivate struct {
	native uintptr
}

var passwordManagerInterfaceStruct *gi.Struct
var passwordManagerInterfaceStruct_Once sync.Once

func passwordManagerInterfaceStruct_Set() {
	passwordManagerInterfaceStruct_Once.Do(func() {
		passwordManagerInterfaceStruct = gi.StructNew("Soup", "PasswordManagerInterface")
	})
}

type PasswordManagerInterface struct {
	native uintptr
	// UNSUPPORTED : C value 'base' : no Go type for 'GObject.TypeInterface'
	// UNSUPPORTED : C value 'get_passwords_async' : missing Type
	// UNSUPPORTED : C value 'get_passwords_sync' : missing Type
}

var proxyResolverDefaultClassStruct *gi.Struct
var proxyResolverDefaultClassStruct_Once sync.Once

func proxyResolverDefaultClassStruct_Set() {
	proxyResolverDefaultClassStruct_Once.Do(func() {
		proxyResolverDefaultClassStruct = gi.StructNew("Soup", "ProxyResolverDefaultClass")
	})
}

type ProxyResolverDefaultClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
}

var proxyResolverInterfaceStruct *gi.Struct
var proxyResolverInterfaceStruct_Once sync.Once

func proxyResolverInterfaceStruct_Set() {
	proxyResolverInterfaceStruct_Once.Do(func() {
		proxyResolverInterfaceStruct = gi.StructNew("Soup", "ProxyResolverInterface")
	})
}

type ProxyResolverInterface struct {
	native uintptr
	// UNSUPPORTED : C value 'base' : no Go type for 'GObject.TypeInterface'
	// UNSUPPORTED : C value 'get_proxy_async' : missing Type
	// UNSUPPORTED : C value 'get_proxy_sync' : missing Type
}

var proxyURIResolverInterfaceStruct *gi.Struct
var proxyURIResolverInterfaceStruct_Once sync.Once

func proxyURIResolverInterfaceStruct_Set() {
	proxyURIResolverInterfaceStruct_Once.Do(func() {
		proxyURIResolverInterfaceStruct = gi.StructNew("Soup", "ProxyURIResolverInterface")
	})
}

type ProxyURIResolverInterface struct {
	native uintptr
	// UNSUPPORTED : C value 'base' : no Go type for 'GObject.TypeInterface'
	// UNSUPPORTED : C value 'get_proxy_uri_async' : missing Type
	// UNSUPPORTED : C value 'get_proxy_uri_sync' : missing Type
	// UNSUPPORTED : C value '_libsoup_reserved1' : missing Type
	// UNSUPPORTED : C value '_libsoup_reserved2' : missing Type
	// UNSUPPORTED : C value '_libsoup_reserved3' : missing Type
	// UNSUPPORTED : C value '_libsoup_reserved4' : missing Type
}

var rangeStruct *gi.Struct
var rangeStruct_Once sync.Once

func rangeStruct_Set() {
	rangeStruct_Once.Do(func() {
		rangeStruct = gi.StructNew("Soup", "Range")
	})
}

type Range struct {
	native uintptr
	Start  int64
	End    int64
}

var requestClassStruct *gi.Struct
var requestClassStruct_Once sync.Once

func requestClassStruct_Set() {
	requestClassStruct_Once.Do(func() {
		requestClassStruct = gi.StructNew("Soup", "RequestClass")
	})
}

type RequestClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent' : no Go type for 'GObject.ObjectClass'
	Schemes string
	// UNSUPPORTED : C value 'check_uri' : missing Type
	// UNSUPPORTED : C value 'send' : missing Type
	// UNSUPPORTED : C value 'send_async' : missing Type
	// UNSUPPORTED : C value 'send_finish' : missing Type
	// UNSUPPORTED : C value 'get_content_length' : missing Type
	// UNSUPPORTED : C value 'get_content_type' : missing Type
}

var requestDataClassStruct *gi.Struct
var requestDataClassStruct_Once sync.Once

func requestDataClassStruct_Set() {
	requestDataClassStruct_Once.Do(func() {
		requestDataClassStruct = gi.StructNew("Soup", "RequestDataClass")
	})
}

type RequestDataClass struct {
	native uintptr
	Parent *RequestClass
}

var requestDataPrivateStruct *gi.Struct
var requestDataPrivateStruct_Once sync.Once

func requestDataPrivateStruct_Set() {
	requestDataPrivateStruct_Once.Do(func() {
		requestDataPrivateStruct = gi.StructNew("Soup", "RequestDataPrivate")
	})
}

type RequestDataPrivate struct {
	native uintptr
}

var requestFileClassStruct *gi.Struct
var requestFileClassStruct_Once sync.Once

func requestFileClassStruct_Set() {
	requestFileClassStruct_Once.Do(func() {
		requestFileClassStruct = gi.StructNew("Soup", "RequestFileClass")
	})
}

type RequestFileClass struct {
	native uintptr
	Parent *RequestClass
}

var requestFilePrivateStruct *gi.Struct
var requestFilePrivateStruct_Once sync.Once

func requestFilePrivateStruct_Set() {
	requestFilePrivateStruct_Once.Do(func() {
		requestFilePrivateStruct = gi.StructNew("Soup", "RequestFilePrivate")
	})
}

type RequestFilePrivate struct {
	native uintptr
}

var requestHTTPClassStruct *gi.Struct
var requestHTTPClassStruct_Once sync.Once

func requestHTTPClassStruct_Set() {
	requestHTTPClassStruct_Once.Do(func() {
		requestHTTPClassStruct = gi.StructNew("Soup", "RequestHTTPClass")
	})
}

type RequestHTTPClass struct {
	native uintptr
	Parent *RequestClass
}

var requestHTTPPrivateStruct *gi.Struct
var requestHTTPPrivateStruct_Once sync.Once

func requestHTTPPrivateStruct_Set() {
	requestHTTPPrivateStruct_Once.Do(func() {
		requestHTTPPrivateStruct = gi.StructNew("Soup", "RequestHTTPPrivate")
	})
}

type RequestHTTPPrivate struct {
	native uintptr
}

var requestPrivateStruct *gi.Struct
var requestPrivateStruct_Once sync.Once

func requestPrivateStruct_Set() {
	requestPrivateStruct_Once.Do(func() {
		requestPrivateStruct = gi.StructNew("Soup", "RequestPrivate")
	})
}

type RequestPrivate struct {
	native uintptr
}

var requesterClassStruct *gi.Struct
var requesterClassStruct_Once sync.Once

func requesterClassStruct_Set() {
	requesterClassStruct_Once.Do(func() {
		requesterClassStruct = gi.StructNew("Soup", "RequesterClass")
	})
}

type RequesterClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
}

var requesterPrivateStruct *gi.Struct
var requesterPrivateStruct_Once sync.Once

func requesterPrivateStruct_Set() {
	requesterPrivateStruct_Once.Do(func() {
		requesterPrivateStruct = gi.StructNew("Soup", "RequesterPrivate")
	})
}

type RequesterPrivate struct {
	native uintptr
}

var serverClassStruct *gi.Struct
var serverClassStruct_Once sync.Once

func serverClassStruct_Set() {
	serverClassStruct_Once.Do(func() {
		serverClassStruct = gi.StructNew("Soup", "ServerClass")
	})
}

type ServerClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value 'request_started' : missing Type
	// UNSUPPORTED : C value 'request_read' : missing Type
	// UNSUPPORTED : C value 'request_finished' : missing Type
	// UNSUPPORTED : C value 'request_aborted' : missing Type
	// UNSUPPORTED : C value '_libsoup_reserved1' : missing Type
	// UNSUPPORTED : C value '_libsoup_reserved2' : missing Type
	// UNSUPPORTED : C value '_libsoup_reserved3' : missing Type
	// UNSUPPORTED : C value '_libsoup_reserved4' : missing Type
}

var sessionAsyncClassStruct *gi.Struct
var sessionAsyncClassStruct_Once sync.Once

func sessionAsyncClassStruct_Set() {
	sessionAsyncClassStruct_Once.Do(func() {
		sessionAsyncClassStruct = gi.StructNew("Soup", "SessionAsyncClass")
	})
}

type SessionAsyncClass struct {
	native      uintptr
	ParentClass *SessionClass
	// UNSUPPORTED : C value '_libsoup_reserved1' : missing Type
	// UNSUPPORTED : C value '_libsoup_reserved2' : missing Type
	// UNSUPPORTED : C value '_libsoup_reserved3' : missing Type
	// UNSUPPORTED : C value '_libsoup_reserved4' : missing Type
}

var sessionClassStruct *gi.Struct
var sessionClassStruct_Once sync.Once

func sessionClassStruct_Set() {
	sessionClassStruct_Once.Do(func() {
		sessionClassStruct = gi.StructNew("Soup", "SessionClass")
	})
}

type SessionClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value 'request_started' : missing Type
	// UNSUPPORTED : C value 'authenticate' : missing Type
	// UNSUPPORTED : C value 'queue_message' : missing Type
	// UNSUPPORTED : C value 'requeue_message' : missing Type
	// UNSUPPORTED : C value 'send_message' : missing Type
	// UNSUPPORTED : C value 'cancel_message' : missing Type
	// UNSUPPORTED : C value 'auth_required' : missing Type
	// UNSUPPORTED : C value 'flush_queue' : missing Type
	// UNSUPPORTED : C value 'kick' : missing Type
	// UNSUPPORTED : C value '_libsoup_reserved4' : missing Type
}

var sessionFeatureInterfaceStruct *gi.Struct
var sessionFeatureInterfaceStruct_Once sync.Once

func sessionFeatureInterfaceStruct_Set() {
	sessionFeatureInterfaceStruct_Once.Do(func() {
		sessionFeatureInterfaceStruct = gi.StructNew("Soup", "SessionFeatureInterface")
	})
}

type SessionFeatureInterface struct {
	native uintptr
	// UNSUPPORTED : C value 'parent' : no Go type for 'GObject.TypeInterface'
	// UNSUPPORTED : C value 'attach' : missing Type
	// UNSUPPORTED : C value 'detach' : missing Type
	// UNSUPPORTED : C value 'request_queued' : missing Type
	// UNSUPPORTED : C value 'request_started' : missing Type
	// UNSUPPORTED : C value 'request_unqueued' : missing Type
	// UNSUPPORTED : C value 'add_feature' : missing Type
	// UNSUPPORTED : C value 'remove_feature' : missing Type
	// UNSUPPORTED : C value 'has_feature' : missing Type
}

var sessionSyncClassStruct *gi.Struct
var sessionSyncClassStruct_Once sync.Once

func sessionSyncClassStruct_Set() {
	sessionSyncClassStruct_Once.Do(func() {
		sessionSyncClassStruct = gi.StructNew("Soup", "SessionSyncClass")
	})
}

type SessionSyncClass struct {
	native      uintptr
	ParentClass *SessionClass
	// UNSUPPORTED : C value '_libsoup_reserved1' : missing Type
	// UNSUPPORTED : C value '_libsoup_reserved2' : missing Type
	// UNSUPPORTED : C value '_libsoup_reserved3' : missing Type
	// UNSUPPORTED : C value '_libsoup_reserved4' : missing Type
}

var socketClassStruct *gi.Struct
var socketClassStruct_Once sync.Once

func socketClassStruct_Set() {
	socketClassStruct_Once.Do(func() {
		socketClassStruct = gi.StructNew("Soup", "SocketClass")
	})
}

type SocketClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value 'readable' : missing Type
	// UNSUPPORTED : C value 'writable' : missing Type
	// UNSUPPORTED : C value 'disconnected' : missing Type
	// UNSUPPORTED : C value 'new_connection' : missing Type
	// UNSUPPORTED : C value '_libsoup_reserved1' : missing Type
	// UNSUPPORTED : C value '_libsoup_reserved2' : missing Type
	// UNSUPPORTED : C value '_libsoup_reserved3' : missing Type
	// UNSUPPORTED : C value '_libsoup_reserved4' : missing Type
}

var uRIStruct *gi.Struct
var uRIStruct_Once sync.Once

func uRIStruct_Set() {
	uRIStruct_Once.Do(func() {
		uRIStruct = gi.StructNew("Soup", "URI")
	})
}

type URI struct {
	native   uintptr
	Scheme   string
	User     string
	Password string
	Host     string
	Port     uint32
	Path     string
	Query    string
	Fragment string
}

var uRINewFunction *gi.Function
var uRINewFunction_Once sync.Once

func uRINewFunction_Set() {
	uRINewFunction_Once.Do(func() {
		uRIStruct_Set()
		uRINewFunction = uRIStruct.InvokerNew("new")
	})
}

// URINew is a representation of the C type soup_uri_new.
func URINew(uriString string) *URI {
	uRINewFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetString(uriString)

	ret := uRINewFunction.Invoke(inArgs[:], nil)

	retGo := &URI{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'soup_uri_new_with_base' : parameter 'base' of type 'URI' not supported

var uRICopyFunction *gi.Function
var uRICopyFunction_Once sync.Once

func uRICopyFunction_Set() {
	uRICopyFunction_Once.Do(func() {
		uRIStruct_Set()
		uRICopyFunction = uRIStruct.InvokerNew("copy")
	})
}

// Copy is a representation of the C type soup_uri_copy.
func (recv *URI) Copy() *URI {
	uRICopyFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := uRICopyFunction.Invoke(inArgs[:], nil)

	retGo := &URI{native: ret.Pointer()}

	return retGo
}

var uRICopyHostFunction *gi.Function
var uRICopyHostFunction_Once sync.Once

func uRICopyHostFunction_Set() {
	uRICopyHostFunction_Once.Do(func() {
		uRIStruct_Set()
		uRICopyHostFunction = uRIStruct.InvokerNew("copy_host")
	})
}

// CopyHost is a representation of the C type soup_uri_copy_host.
func (recv *URI) CopyHost() *URI {
	uRICopyHostFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := uRICopyHostFunction.Invoke(inArgs[:], nil)

	retGo := &URI{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'soup_uri_equal' : parameter 'uri2' of type 'URI' not supported

var uRIFreeFunction *gi.Function
var uRIFreeFunction_Once sync.Once

func uRIFreeFunction_Set() {
	uRIFreeFunction_Once.Do(func() {
		uRIStruct_Set()
		uRIFreeFunction = uRIStruct.InvokerNew("free")
	})
}

// Free is a representation of the C type soup_uri_free.
func (recv *URI) Free() {
	uRIFreeFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	uRIFreeFunction.Invoke(inArgs[:], nil)

}

var uRIGetFragmentFunction *gi.Function
var uRIGetFragmentFunction_Once sync.Once

func uRIGetFragmentFunction_Set() {
	uRIGetFragmentFunction_Once.Do(func() {
		uRIStruct_Set()
		uRIGetFragmentFunction = uRIStruct.InvokerNew("get_fragment")
	})
}

// GetFragment is a representation of the C type soup_uri_get_fragment.
func (recv *URI) GetFragment() string {
	uRIGetFragmentFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := uRIGetFragmentFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

var uRIGetHostFunction *gi.Function
var uRIGetHostFunction_Once sync.Once

func uRIGetHostFunction_Set() {
	uRIGetHostFunction_Once.Do(func() {
		uRIStruct_Set()
		uRIGetHostFunction = uRIStruct.InvokerNew("get_host")
	})
}

// GetHost is a representation of the C type soup_uri_get_host.
func (recv *URI) GetHost() string {
	uRIGetHostFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := uRIGetHostFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

var uRIGetPasswordFunction *gi.Function
var uRIGetPasswordFunction_Once sync.Once

func uRIGetPasswordFunction_Set() {
	uRIGetPasswordFunction_Once.Do(func() {
		uRIStruct_Set()
		uRIGetPasswordFunction = uRIStruct.InvokerNew("get_password")
	})
}

// GetPassword is a representation of the C type soup_uri_get_password.
func (recv *URI) GetPassword() string {
	uRIGetPasswordFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := uRIGetPasswordFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

var uRIGetPathFunction *gi.Function
var uRIGetPathFunction_Once sync.Once

func uRIGetPathFunction_Set() {
	uRIGetPathFunction_Once.Do(func() {
		uRIStruct_Set()
		uRIGetPathFunction = uRIStruct.InvokerNew("get_path")
	})
}

// GetPath is a representation of the C type soup_uri_get_path.
func (recv *URI) GetPath() string {
	uRIGetPathFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := uRIGetPathFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

var uRIGetPortFunction *gi.Function
var uRIGetPortFunction_Once sync.Once

func uRIGetPortFunction_Set() {
	uRIGetPortFunction_Once.Do(func() {
		uRIStruct_Set()
		uRIGetPortFunction = uRIStruct.InvokerNew("get_port")
	})
}

// GetPort is a representation of the C type soup_uri_get_port.
func (recv *URI) GetPort() uint32 {
	uRIGetPortFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := uRIGetPortFunction.Invoke(inArgs[:], nil)

	retGo := ret.Uint32()

	return retGo
}

var uRIGetQueryFunction *gi.Function
var uRIGetQueryFunction_Once sync.Once

func uRIGetQueryFunction_Set() {
	uRIGetQueryFunction_Once.Do(func() {
		uRIStruct_Set()
		uRIGetQueryFunction = uRIStruct.InvokerNew("get_query")
	})
}

// GetQuery is a representation of the C type soup_uri_get_query.
func (recv *URI) GetQuery() string {
	uRIGetQueryFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := uRIGetQueryFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

var uRIGetSchemeFunction *gi.Function
var uRIGetSchemeFunction_Once sync.Once

func uRIGetSchemeFunction_Set() {
	uRIGetSchemeFunction_Once.Do(func() {
		uRIStruct_Set()
		uRIGetSchemeFunction = uRIStruct.InvokerNew("get_scheme")
	})
}

// GetScheme is a representation of the C type soup_uri_get_scheme.
func (recv *URI) GetScheme() string {
	uRIGetSchemeFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := uRIGetSchemeFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

var uRIGetUserFunction *gi.Function
var uRIGetUserFunction_Once sync.Once

func uRIGetUserFunction_Set() {
	uRIGetUserFunction_Once.Do(func() {
		uRIStruct_Set()
		uRIGetUserFunction = uRIStruct.InvokerNew("get_user")
	})
}

// GetUser is a representation of the C type soup_uri_get_user.
func (recv *URI) GetUser() string {
	uRIGetUserFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := uRIGetUserFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'soup_uri_host_equal' : parameter 'v2' of type 'URI' not supported

var uRIHostHashFunction *gi.Function
var uRIHostHashFunction_Once sync.Once

func uRIHostHashFunction_Set() {
	uRIHostHashFunction_Once.Do(func() {
		uRIStruct_Set()
		uRIHostHashFunction = uRIStruct.InvokerNew("host_hash")
	})
}

// HostHash is a representation of the C type soup_uri_host_hash.
func (recv *URI) HostHash() uint32 {
	uRIHostHashFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := uRIHostHashFunction.Invoke(inArgs[:], nil)

	retGo := ret.Uint32()

	return retGo
}

var uRISetFragmentFunction *gi.Function
var uRISetFragmentFunction_Once sync.Once

func uRISetFragmentFunction_Set() {
	uRISetFragmentFunction_Once.Do(func() {
		uRIStruct_Set()
		uRISetFragmentFunction = uRIStruct.InvokerNew("set_fragment")
	})
}

// SetFragment is a representation of the C type soup_uri_set_fragment.
func (recv *URI) SetFragment(fragment string) {
	uRISetFragmentFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(fragment)

	uRISetFragmentFunction.Invoke(inArgs[:], nil)

}

var uRISetHostFunction *gi.Function
var uRISetHostFunction_Once sync.Once

func uRISetHostFunction_Set() {
	uRISetHostFunction_Once.Do(func() {
		uRIStruct_Set()
		uRISetHostFunction = uRIStruct.InvokerNew("set_host")
	})
}

// SetHost is a representation of the C type soup_uri_set_host.
func (recv *URI) SetHost(host string) {
	uRISetHostFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(host)

	uRISetHostFunction.Invoke(inArgs[:], nil)

}

var uRISetPasswordFunction *gi.Function
var uRISetPasswordFunction_Once sync.Once

func uRISetPasswordFunction_Set() {
	uRISetPasswordFunction_Once.Do(func() {
		uRIStruct_Set()
		uRISetPasswordFunction = uRIStruct.InvokerNew("set_password")
	})
}

// SetPassword is a representation of the C type soup_uri_set_password.
func (recv *URI) SetPassword(password string) {
	uRISetPasswordFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(password)

	uRISetPasswordFunction.Invoke(inArgs[:], nil)

}

var uRISetPathFunction *gi.Function
var uRISetPathFunction_Once sync.Once

func uRISetPathFunction_Set() {
	uRISetPathFunction_Once.Do(func() {
		uRIStruct_Set()
		uRISetPathFunction = uRIStruct.InvokerNew("set_path")
	})
}

// SetPath is a representation of the C type soup_uri_set_path.
func (recv *URI) SetPath(path string) {
	uRISetPathFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(path)

	uRISetPathFunction.Invoke(inArgs[:], nil)

}

var uRISetPortFunction *gi.Function
var uRISetPortFunction_Once sync.Once

func uRISetPortFunction_Set() {
	uRISetPortFunction_Once.Do(func() {
		uRIStruct_Set()
		uRISetPortFunction = uRIStruct.InvokerNew("set_port")
	})
}

// SetPort is a representation of the C type soup_uri_set_port.
func (recv *URI) SetPort(port uint32) {
	uRISetPortFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetUint32(port)

	uRISetPortFunction.Invoke(inArgs[:], nil)

}

var uRISetQueryFunction *gi.Function
var uRISetQueryFunction_Once sync.Once

func uRISetQueryFunction_Set() {
	uRISetQueryFunction_Once.Do(func() {
		uRIStruct_Set()
		uRISetQueryFunction = uRIStruct.InvokerNew("set_query")
	})
}

// SetQuery is a representation of the C type soup_uri_set_query.
func (recv *URI) SetQuery(query string) {
	uRISetQueryFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(query)

	uRISetQueryFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'soup_uri_set_query_from_fields' : parameter '...' has no type

// UNSUPPORTED : C value 'soup_uri_set_query_from_form' : parameter 'form' of type 'GLib.HashTable' not supported

var uRISetSchemeFunction *gi.Function
var uRISetSchemeFunction_Once sync.Once

func uRISetSchemeFunction_Set() {
	uRISetSchemeFunction_Once.Do(func() {
		uRIStruct_Set()
		uRISetSchemeFunction = uRIStruct.InvokerNew("set_scheme")
	})
}

// SetScheme is a representation of the C type soup_uri_set_scheme.
func (recv *URI) SetScheme(scheme string) {
	uRISetSchemeFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(scheme)

	uRISetSchemeFunction.Invoke(inArgs[:], nil)

}

var uRISetUserFunction *gi.Function
var uRISetUserFunction_Once sync.Once

func uRISetUserFunction_Set() {
	uRISetUserFunction_Once.Do(func() {
		uRIStruct_Set()
		uRISetUserFunction = uRIStruct.InvokerNew("set_user")
	})
}

// SetUser is a representation of the C type soup_uri_set_user.
func (recv *URI) SetUser(user string) {
	uRISetUserFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(user)

	uRISetUserFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'soup_uri_to_string' : parameter 'just_path_and_query' of type 'gboolean' not supported

// UNSUPPORTED : C value 'soup_uri_uses_default_port' : return type 'gboolean' not supported

var websocketConnectionClassStruct *gi.Struct
var websocketConnectionClassStruct_Once sync.Once

func websocketConnectionClassStruct_Set() {
	websocketConnectionClassStruct_Once.Do(func() {
		websocketConnectionClassStruct = gi.StructNew("Soup", "WebsocketConnectionClass")
	})
}

type WebsocketConnectionClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value 'message' : missing Type
	// UNSUPPORTED : C value 'error' : missing Type
	// UNSUPPORTED : C value 'closing' : missing Type
	// UNSUPPORTED : C value 'closed' : missing Type
	// UNSUPPORTED : C value 'pong' : missing Type
}

var websocketConnectionPrivateStruct *gi.Struct
var websocketConnectionPrivateStruct_Once sync.Once

func websocketConnectionPrivateStruct_Set() {
	websocketConnectionPrivateStruct_Once.Do(func() {
		websocketConnectionPrivateStruct = gi.StructNew("Soup", "WebsocketConnectionPrivate")
	})
}

type WebsocketConnectionPrivate struct {
	native uintptr
}

var websocketExtensionClassStruct *gi.Struct
var websocketExtensionClassStruct_Once sync.Once

func websocketExtensionClassStruct_Set() {
	websocketExtensionClassStruct_Once.Do(func() {
		websocketExtensionClassStruct = gi.StructNew("Soup", "WebsocketExtensionClass")
	})
}

type WebsocketExtensionClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	Name string
	// UNSUPPORTED : C value 'configure' : missing Type
	// UNSUPPORTED : C value 'get_request_params' : missing Type
	// UNSUPPORTED : C value 'get_response_params' : missing Type
	// UNSUPPORTED : C value 'process_outgoing_message' : missing Type
	// UNSUPPORTED : C value 'process_incoming_message' : missing Type
	// UNSUPPORTED : C value '_libsoup_reserved1' : missing Type
	// UNSUPPORTED : C value '_libsoup_reserved2' : missing Type
	// UNSUPPORTED : C value '_libsoup_reserved3' : missing Type
	// UNSUPPORTED : C value '_libsoup_reserved4' : missing Type
}

var websocketExtensionDeflateClassStruct *gi.Struct
var websocketExtensionDeflateClassStruct_Once sync.Once

func websocketExtensionDeflateClassStruct_Set() {
	websocketExtensionDeflateClassStruct_Once.Do(func() {
		websocketExtensionDeflateClassStruct = gi.StructNew("Soup", "WebsocketExtensionDeflateClass")
	})
}

type WebsocketExtensionDeflateClass struct {
	native      uintptr
	ParentClass *WebsocketExtensionClass
}

var websocketExtensionManagerClassStruct *gi.Struct
var websocketExtensionManagerClassStruct_Once sync.Once

func websocketExtensionManagerClassStruct_Set() {
	websocketExtensionManagerClassStruct_Once.Do(func() {
		websocketExtensionManagerClassStruct = gi.StructNew("Soup", "WebsocketExtensionManagerClass")
	})
}

type WebsocketExtensionManagerClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
}

var xMLRPCParamsStruct *gi.Struct
var xMLRPCParamsStruct_Once sync.Once

func xMLRPCParamsStruct_Set() {
	xMLRPCParamsStruct_Once.Do(func() {
		xMLRPCParamsStruct = gi.StructNew("Soup", "XMLRPCParams")
	})
}

type XMLRPCParams struct {
	native uintptr
}

var xMLRPCParamsFreeFunction *gi.Function
var xMLRPCParamsFreeFunction_Once sync.Once

func xMLRPCParamsFreeFunction_Set() {
	xMLRPCParamsFreeFunction_Once.Do(func() {
		xMLRPCParamsStruct_Set()
		xMLRPCParamsFreeFunction = xMLRPCParamsStruct.InvokerNew("free")
	})
}

// Free is a representation of the C type soup_xmlrpc_params_free.
func (recv *XMLRPCParams) Free() {
	xMLRPCParamsFreeFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	xMLRPCParamsFreeFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'soup_xmlrpc_params_parse' : return type 'GLib.Variant' not supported
