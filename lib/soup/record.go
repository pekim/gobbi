// Code generated - DO NOT EDIT.

package soup

import (
	gi "github.com/pekim/gobbi/internal/gi"
	"sync"
)

var addressClassStruct *gi.Struct
var addressClassStructOnce sync.Once

func addressClassStructSet() {
	addressClassStructOnce.Do(func() {
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
var authClassStructOnce sync.Once

func authClassStructSet() {
	authClassStructOnce.Do(func() {
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
var authDomainBasicClassStructOnce sync.Once

func authDomainBasicClassStructSet() {
	authDomainBasicClassStructOnce.Do(func() {
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
var authDomainClassStructOnce sync.Once

func authDomainClassStructSet() {
	authDomainClassStructOnce.Do(func() {
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
var authDomainDigestClassStructOnce sync.Once

func authDomainDigestClassStructSet() {
	authDomainDigestClassStructOnce.Do(func() {
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
var authManagerClassStructOnce sync.Once

func authManagerClassStructSet() {
	authManagerClassStructOnce.Do(func() {
		authManagerClassStruct = gi.StructNew("Soup", "AuthManagerClass")
	})
}

type AuthManagerClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value 'authenticate' : missing Type
}

var authManagerPrivateStruct *gi.Struct
var authManagerPrivateStructOnce sync.Once

func authManagerPrivateStructSet() {
	authManagerPrivateStructOnce.Do(func() {
		authManagerPrivateStruct = gi.StructNew("Soup", "AuthManagerPrivate")
	})
}

type AuthManagerPrivate struct {
	native uintptr
}

var bufferStruct *gi.Struct
var bufferStructOnce sync.Once

func bufferStructSet() {
	bufferStructOnce.Do(func() {
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
var bufferCopyFunctionOnce sync.Once

func bufferCopyFunctionSet() {
	bufferCopyFunctionOnce.Do(func() {
		bufferCopyFunction = gi.FunctionInvokerNew("Soup", "copy")
	})
}

var copyBufferInvoker *gi.Function

// Copy is a representation of the C type soup_buffer_copy.
func (recv *Buffer) Copy() *Buffer {
	if copyBufferInvoker == nil {
		copyBufferInvoker = gi.StructFunctionInvokerNew("Soup", "Buffer", "copy")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := copyBufferInvoker.Invoke(inArgs[:], nil)

	retGo := &Buffer{native: ret.Pointer()}

	return retGo
}

var bufferFreeFunction *gi.Function
var bufferFreeFunctionOnce sync.Once

func bufferFreeFunctionSet() {
	bufferFreeFunctionOnce.Do(func() {
		bufferFreeFunction = gi.FunctionInvokerNew("Soup", "free")
	})
}

var freeBufferInvoker *gi.Function

// Free is a representation of the C type soup_buffer_free.
func (recv *Buffer) Free() {
	if freeBufferInvoker == nil {
		freeBufferInvoker = gi.StructFunctionInvokerNew("Soup", "Buffer", "free")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	freeBufferInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'soup_buffer_get_as_bytes' : return type 'GLib.Bytes' not supported

// UNSUPPORTED : C value 'soup_buffer_get_data' : parameter 'data' has no type

// UNSUPPORTED : C value 'soup_buffer_get_owner' : return type 'gpointer' not supported

// UNSUPPORTED : C value 'soup_buffer_new_subbuffer' : parameter 'offset' of type 'gsize' not supported

var cacheClassStruct *gi.Struct
var cacheClassStructOnce sync.Once

func cacheClassStructSet() {
	cacheClassStructOnce.Do(func() {
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
var cachePrivateStructOnce sync.Once

func cachePrivateStructSet() {
	cachePrivateStructOnce.Do(func() {
		cachePrivateStruct = gi.StructNew("Soup", "CachePrivate")
	})
}

type CachePrivate struct {
	native uintptr
}

var clientContextStruct *gi.Struct
var clientContextStructOnce sync.Once

func clientContextStructSet() {
	clientContextStructOnce.Do(func() {
		clientContextStruct = gi.StructNew("Soup", "ClientContext")
	})
}

type ClientContext struct {
	native uintptr
}

// UNSUPPORTED : C value 'soup_client_context_get_address' : return type 'Address' not supported

// UNSUPPORTED : C value 'soup_client_context_get_auth_domain' : return type 'AuthDomain' not supported

var clientContextGetAuthUserFunction *gi.Function
var clientContextGetAuthUserFunctionOnce sync.Once

func clientContextGetAuthUserFunctionSet() {
	clientContextGetAuthUserFunctionOnce.Do(func() {
		clientContextGetAuthUserFunction = gi.FunctionInvokerNew("Soup", "get_auth_user")
	})
}

var getAuthUserClientContextInvoker *gi.Function

// GetAuthUser is a representation of the C type soup_client_context_get_auth_user.
func (recv *ClientContext) GetAuthUser() string {
	if getAuthUserClientContextInvoker == nil {
		getAuthUserClientContextInvoker = gi.StructFunctionInvokerNew("Soup", "ClientContext", "get_auth_user")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getAuthUserClientContextInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'soup_client_context_get_gsocket' : return type 'Gio.Socket' not supported

var clientContextGetHostFunction *gi.Function
var clientContextGetHostFunctionOnce sync.Once

func clientContextGetHostFunctionSet() {
	clientContextGetHostFunctionOnce.Do(func() {
		clientContextGetHostFunction = gi.FunctionInvokerNew("Soup", "get_host")
	})
}

var getHostClientContextInvoker *gi.Function

// GetHost is a representation of the C type soup_client_context_get_host.
func (recv *ClientContext) GetHost() string {
	if getHostClientContextInvoker == nil {
		getHostClientContextInvoker = gi.StructFunctionInvokerNew("Soup", "ClientContext", "get_host")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getHostClientContextInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'soup_client_context_get_local_address' : return type 'Gio.SocketAddress' not supported

// UNSUPPORTED : C value 'soup_client_context_get_remote_address' : return type 'Gio.SocketAddress' not supported

// UNSUPPORTED : C value 'soup_client_context_get_socket' : return type 'Socket' not supported

// UNSUPPORTED : C value 'soup_client_context_steal_connection' : return type 'Gio.IOStream' not supported

var connectionStruct *gi.Struct
var connectionStructOnce sync.Once

func connectionStructSet() {
	connectionStructOnce.Do(func() {
		connectionStruct = gi.StructNew("Soup", "Connection")
	})
}

type Connection struct {
	native uintptr
}

var contentDecoderClassStruct *gi.Struct
var contentDecoderClassStructOnce sync.Once

func contentDecoderClassStructSet() {
	contentDecoderClassStructOnce.Do(func() {
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
var contentDecoderPrivateStructOnce sync.Once

func contentDecoderPrivateStructSet() {
	contentDecoderPrivateStructOnce.Do(func() {
		contentDecoderPrivateStruct = gi.StructNew("Soup", "ContentDecoderPrivate")
	})
}

type ContentDecoderPrivate struct {
	native uintptr
}

var contentSnifferClassStruct *gi.Struct
var contentSnifferClassStructOnce sync.Once

func contentSnifferClassStructSet() {
	contentSnifferClassStructOnce.Do(func() {
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
var contentSnifferPrivateStructOnce sync.Once

func contentSnifferPrivateStructSet() {
	contentSnifferPrivateStructOnce.Do(func() {
		contentSnifferPrivateStruct = gi.StructNew("Soup", "ContentSnifferPrivate")
	})
}

type ContentSnifferPrivate struct {
	native uintptr
}

var cookieStruct *gi.Struct
var cookieStructOnce sync.Once

func cookieStructSet() {
	cookieStructOnce.Do(func() {
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
var cookieNewFunctionOnce sync.Once

func cookieNewFunctionSet() {
	cookieNewFunctionOnce.Do(func() {
		cookieNewFunction = gi.FunctionInvokerNew("Soup", "new")
	})
}

var newCookieInvoker *gi.Function

// CookieNew is a representation of the C type soup_cookie_new.
func CookieNew(name string, value string, domain string, path string, maxAge int32) *Cookie {
	if newCookieInvoker == nil {
		newCookieInvoker = gi.StructFunctionInvokerNew("Soup", "Cookie", "new")
	}

	var inArgs [5]gi.Argument
	inArgs[0].SetString(name)
	inArgs[1].SetString(value)
	inArgs[2].SetString(domain)
	inArgs[3].SetString(path)
	inArgs[4].SetInt32(maxAge)

	ret := newCookieInvoker.Invoke(inArgs[:], nil)

	retGo := &Cookie{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'soup_cookie_applies_to_uri' : parameter 'uri' of type 'URI' not supported

var cookieCopyFunction *gi.Function
var cookieCopyFunctionOnce sync.Once

func cookieCopyFunctionSet() {
	cookieCopyFunctionOnce.Do(func() {
		cookieCopyFunction = gi.FunctionInvokerNew("Soup", "copy")
	})
}

var copyCookieInvoker *gi.Function

// Copy is a representation of the C type soup_cookie_copy.
func (recv *Cookie) Copy() *Cookie {
	if copyCookieInvoker == nil {
		copyCookieInvoker = gi.StructFunctionInvokerNew("Soup", "Cookie", "copy")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := copyCookieInvoker.Invoke(inArgs[:], nil)

	retGo := &Cookie{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'soup_cookie_domain_matches' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'soup_cookie_equal' : parameter 'cookie2' of type 'Cookie' not supported

var cookieFreeFunction *gi.Function
var cookieFreeFunctionOnce sync.Once

func cookieFreeFunctionSet() {
	cookieFreeFunctionOnce.Do(func() {
		cookieFreeFunction = gi.FunctionInvokerNew("Soup", "free")
	})
}

var freeCookieInvoker *gi.Function

// Free is a representation of the C type soup_cookie_free.
func (recv *Cookie) Free() {
	if freeCookieInvoker == nil {
		freeCookieInvoker = gi.StructFunctionInvokerNew("Soup", "Cookie", "free")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	freeCookieInvoker.Invoke(inArgs[:], nil)

}

var cookieGetDomainFunction *gi.Function
var cookieGetDomainFunctionOnce sync.Once

func cookieGetDomainFunctionSet() {
	cookieGetDomainFunctionOnce.Do(func() {
		cookieGetDomainFunction = gi.FunctionInvokerNew("Soup", "get_domain")
	})
}

var getDomainCookieInvoker *gi.Function

// GetDomain is a representation of the C type soup_cookie_get_domain.
func (recv *Cookie) GetDomain() string {
	if getDomainCookieInvoker == nil {
		getDomainCookieInvoker = gi.StructFunctionInvokerNew("Soup", "Cookie", "get_domain")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getDomainCookieInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

var cookieGetExpiresFunction *gi.Function
var cookieGetExpiresFunctionOnce sync.Once

func cookieGetExpiresFunctionSet() {
	cookieGetExpiresFunctionOnce.Do(func() {
		cookieGetExpiresFunction = gi.FunctionInvokerNew("Soup", "get_expires")
	})
}

var getExpiresCookieInvoker *gi.Function

// GetExpires is a representation of the C type soup_cookie_get_expires.
func (recv *Cookie) GetExpires() *Date {
	if getExpiresCookieInvoker == nil {
		getExpiresCookieInvoker = gi.StructFunctionInvokerNew("Soup", "Cookie", "get_expires")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getExpiresCookieInvoker.Invoke(inArgs[:], nil)

	retGo := &Date{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'soup_cookie_get_http_only' : return type 'gboolean' not supported

var cookieGetNameFunction *gi.Function
var cookieGetNameFunctionOnce sync.Once

func cookieGetNameFunctionSet() {
	cookieGetNameFunctionOnce.Do(func() {
		cookieGetNameFunction = gi.FunctionInvokerNew("Soup", "get_name")
	})
}

var getNameCookieInvoker *gi.Function

// GetName is a representation of the C type soup_cookie_get_name.
func (recv *Cookie) GetName() string {
	if getNameCookieInvoker == nil {
		getNameCookieInvoker = gi.StructFunctionInvokerNew("Soup", "Cookie", "get_name")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getNameCookieInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

var cookieGetPathFunction *gi.Function
var cookieGetPathFunctionOnce sync.Once

func cookieGetPathFunctionSet() {
	cookieGetPathFunctionOnce.Do(func() {
		cookieGetPathFunction = gi.FunctionInvokerNew("Soup", "get_path")
	})
}

var getPathCookieInvoker *gi.Function

// GetPath is a representation of the C type soup_cookie_get_path.
func (recv *Cookie) GetPath() string {
	if getPathCookieInvoker == nil {
		getPathCookieInvoker = gi.StructFunctionInvokerNew("Soup", "Cookie", "get_path")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getPathCookieInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'soup_cookie_get_secure' : return type 'gboolean' not supported

var cookieGetValueFunction *gi.Function
var cookieGetValueFunctionOnce sync.Once

func cookieGetValueFunctionSet() {
	cookieGetValueFunctionOnce.Do(func() {
		cookieGetValueFunction = gi.FunctionInvokerNew("Soup", "get_value")
	})
}

var getValueCookieInvoker *gi.Function

// GetValue is a representation of the C type soup_cookie_get_value.
func (recv *Cookie) GetValue() string {
	if getValueCookieInvoker == nil {
		getValueCookieInvoker = gi.StructFunctionInvokerNew("Soup", "Cookie", "get_value")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getValueCookieInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

var cookieSetDomainFunction *gi.Function
var cookieSetDomainFunctionOnce sync.Once

func cookieSetDomainFunctionSet() {
	cookieSetDomainFunctionOnce.Do(func() {
		cookieSetDomainFunction = gi.FunctionInvokerNew("Soup", "set_domain")
	})
}

var setDomainCookieInvoker *gi.Function

// SetDomain is a representation of the C type soup_cookie_set_domain.
func (recv *Cookie) SetDomain(domain string) {
	if setDomainCookieInvoker == nil {
		setDomainCookieInvoker = gi.StructFunctionInvokerNew("Soup", "Cookie", "set_domain")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(domain)

	setDomainCookieInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'soup_cookie_set_expires' : parameter 'expires' of type 'Date' not supported

// UNSUPPORTED : C value 'soup_cookie_set_http_only' : parameter 'http_only' of type 'gboolean' not supported

var cookieSetMaxAgeFunction *gi.Function
var cookieSetMaxAgeFunctionOnce sync.Once

func cookieSetMaxAgeFunctionSet() {
	cookieSetMaxAgeFunctionOnce.Do(func() {
		cookieSetMaxAgeFunction = gi.FunctionInvokerNew("Soup", "set_max_age")
	})
}

var setMaxAgeCookieInvoker *gi.Function

// SetMaxAge is a representation of the C type soup_cookie_set_max_age.
func (recv *Cookie) SetMaxAge(maxAge int32) {
	if setMaxAgeCookieInvoker == nil {
		setMaxAgeCookieInvoker = gi.StructFunctionInvokerNew("Soup", "Cookie", "set_max_age")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(maxAge)

	setMaxAgeCookieInvoker.Invoke(inArgs[:], nil)

}

var cookieSetNameFunction *gi.Function
var cookieSetNameFunctionOnce sync.Once

func cookieSetNameFunctionSet() {
	cookieSetNameFunctionOnce.Do(func() {
		cookieSetNameFunction = gi.FunctionInvokerNew("Soup", "set_name")
	})
}

var setNameCookieInvoker *gi.Function

// SetName is a representation of the C type soup_cookie_set_name.
func (recv *Cookie) SetName(name string) {
	if setNameCookieInvoker == nil {
		setNameCookieInvoker = gi.StructFunctionInvokerNew("Soup", "Cookie", "set_name")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(name)

	setNameCookieInvoker.Invoke(inArgs[:], nil)

}

var cookieSetPathFunction *gi.Function
var cookieSetPathFunctionOnce sync.Once

func cookieSetPathFunctionSet() {
	cookieSetPathFunctionOnce.Do(func() {
		cookieSetPathFunction = gi.FunctionInvokerNew("Soup", "set_path")
	})
}

var setPathCookieInvoker *gi.Function

// SetPath is a representation of the C type soup_cookie_set_path.
func (recv *Cookie) SetPath(path string) {
	if setPathCookieInvoker == nil {
		setPathCookieInvoker = gi.StructFunctionInvokerNew("Soup", "Cookie", "set_path")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(path)

	setPathCookieInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'soup_cookie_set_secure' : parameter 'secure' of type 'gboolean' not supported

var cookieSetValueFunction *gi.Function
var cookieSetValueFunctionOnce sync.Once

func cookieSetValueFunctionSet() {
	cookieSetValueFunctionOnce.Do(func() {
		cookieSetValueFunction = gi.FunctionInvokerNew("Soup", "set_value")
	})
}

var setValueCookieInvoker *gi.Function

// SetValue is a representation of the C type soup_cookie_set_value.
func (recv *Cookie) SetValue(value string) {
	if setValueCookieInvoker == nil {
		setValueCookieInvoker = gi.StructFunctionInvokerNew("Soup", "Cookie", "set_value")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(value)

	setValueCookieInvoker.Invoke(inArgs[:], nil)

}

var cookieToCookieHeaderFunction *gi.Function
var cookieToCookieHeaderFunctionOnce sync.Once

func cookieToCookieHeaderFunctionSet() {
	cookieToCookieHeaderFunctionOnce.Do(func() {
		cookieToCookieHeaderFunction = gi.FunctionInvokerNew("Soup", "to_cookie_header")
	})
}

var toCookieHeaderCookieInvoker *gi.Function

// ToCookieHeader is a representation of the C type soup_cookie_to_cookie_header.
func (recv *Cookie) ToCookieHeader() string {
	if toCookieHeaderCookieInvoker == nil {
		toCookieHeaderCookieInvoker = gi.StructFunctionInvokerNew("Soup", "Cookie", "to_cookie_header")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := toCookieHeaderCookieInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

var cookieToSetCookieHeaderFunction *gi.Function
var cookieToSetCookieHeaderFunctionOnce sync.Once

func cookieToSetCookieHeaderFunctionSet() {
	cookieToSetCookieHeaderFunctionOnce.Do(func() {
		cookieToSetCookieHeaderFunction = gi.FunctionInvokerNew("Soup", "to_set_cookie_header")
	})
}

var toSetCookieHeaderCookieInvoker *gi.Function

// ToSetCookieHeader is a representation of the C type soup_cookie_to_set_cookie_header.
func (recv *Cookie) ToSetCookieHeader() string {
	if toSetCookieHeaderCookieInvoker == nil {
		toSetCookieHeaderCookieInvoker = gi.StructFunctionInvokerNew("Soup", "Cookie", "to_set_cookie_header")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := toSetCookieHeaderCookieInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

var cookieJarClassStruct *gi.Struct
var cookieJarClassStructOnce sync.Once

func cookieJarClassStructSet() {
	cookieJarClassStructOnce.Do(func() {
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
var cookieJarDBClassStructOnce sync.Once

func cookieJarDBClassStructSet() {
	cookieJarDBClassStructOnce.Do(func() {
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
var cookieJarTextClassStructOnce sync.Once

func cookieJarTextClassStructSet() {
	cookieJarTextClassStructOnce.Do(func() {
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
var dateStructOnce sync.Once

func dateStructSet() {
	dateStructOnce.Do(func() {
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
var dateNewFunctionOnce sync.Once

func dateNewFunctionSet() {
	dateNewFunctionOnce.Do(func() {
		dateNewFunction = gi.FunctionInvokerNew("Soup", "new")
	})
}

var newDateInvoker *gi.Function

// DateNew is a representation of the C type soup_date_new.
func DateNew(year int32, month int32, day int32, hour int32, minute int32, second int32) *Date {
	if newDateInvoker == nil {
		newDateInvoker = gi.StructFunctionInvokerNew("Soup", "Date", "new")
	}

	var inArgs [6]gi.Argument
	inArgs[0].SetInt32(year)
	inArgs[1].SetInt32(month)
	inArgs[2].SetInt32(day)
	inArgs[3].SetInt32(hour)
	inArgs[4].SetInt32(minute)
	inArgs[5].SetInt32(second)

	ret := newDateInvoker.Invoke(inArgs[:], nil)

	retGo := &Date{native: ret.Pointer()}

	return retGo
}

var dateNewFromNowFunction *gi.Function
var dateNewFromNowFunctionOnce sync.Once

func dateNewFromNowFunctionSet() {
	dateNewFromNowFunctionOnce.Do(func() {
		dateNewFromNowFunction = gi.FunctionInvokerNew("Soup", "new_from_now")
	})
}

var newFromNowDateInvoker *gi.Function

// DateNewFromNow is a representation of the C type soup_date_new_from_now.
func DateNewFromNow(offsetSeconds int32) *Date {
	if newFromNowDateInvoker == nil {
		newFromNowDateInvoker = gi.StructFunctionInvokerNew("Soup", "Date", "new_from_now")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetInt32(offsetSeconds)

	ret := newFromNowDateInvoker.Invoke(inArgs[:], nil)

	retGo := &Date{native: ret.Pointer()}

	return retGo
}

var dateNewFromStringFunction *gi.Function
var dateNewFromStringFunctionOnce sync.Once

func dateNewFromStringFunctionSet() {
	dateNewFromStringFunctionOnce.Do(func() {
		dateNewFromStringFunction = gi.FunctionInvokerNew("Soup", "new_from_string")
	})
}

var newFromStringDateInvoker *gi.Function

// DateNewFromString is a representation of the C type soup_date_new_from_string.
func DateNewFromString(dateString string) *Date {
	if newFromStringDateInvoker == nil {
		newFromStringDateInvoker = gi.StructFunctionInvokerNew("Soup", "Date", "new_from_string")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetString(dateString)

	ret := newFromStringDateInvoker.Invoke(inArgs[:], nil)

	retGo := &Date{native: ret.Pointer()}

	return retGo
}

var dateNewFromTimeTFunction *gi.Function
var dateNewFromTimeTFunctionOnce sync.Once

func dateNewFromTimeTFunctionSet() {
	dateNewFromTimeTFunctionOnce.Do(func() {
		dateNewFromTimeTFunction = gi.FunctionInvokerNew("Soup", "new_from_time_t")
	})
}

var newFromTimeTDateInvoker *gi.Function

// DateNewFromTimeT is a representation of the C type soup_date_new_from_time_t.
func DateNewFromTimeT(when int64) *Date {
	if newFromTimeTDateInvoker == nil {
		newFromTimeTDateInvoker = gi.StructFunctionInvokerNew("Soup", "Date", "new_from_time_t")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetInt64(when)

	ret := newFromTimeTDateInvoker.Invoke(inArgs[:], nil)

	retGo := &Date{native: ret.Pointer()}

	return retGo
}

var dateCopyFunction *gi.Function
var dateCopyFunctionOnce sync.Once

func dateCopyFunctionSet() {
	dateCopyFunctionOnce.Do(func() {
		dateCopyFunction = gi.FunctionInvokerNew("Soup", "copy")
	})
}

var copyDateInvoker *gi.Function

// Copy is a representation of the C type soup_date_copy.
func (recv *Date) Copy() *Date {
	if copyDateInvoker == nil {
		copyDateInvoker = gi.StructFunctionInvokerNew("Soup", "Date", "copy")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := copyDateInvoker.Invoke(inArgs[:], nil)

	retGo := &Date{native: ret.Pointer()}

	return retGo
}

var dateFreeFunction *gi.Function
var dateFreeFunctionOnce sync.Once

func dateFreeFunctionSet() {
	dateFreeFunctionOnce.Do(func() {
		dateFreeFunction = gi.FunctionInvokerNew("Soup", "free")
	})
}

var freeDateInvoker *gi.Function

// Free is a representation of the C type soup_date_free.
func (recv *Date) Free() {
	if freeDateInvoker == nil {
		freeDateInvoker = gi.StructFunctionInvokerNew("Soup", "Date", "free")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	freeDateInvoker.Invoke(inArgs[:], nil)

}

var dateGetDayFunction *gi.Function
var dateGetDayFunctionOnce sync.Once

func dateGetDayFunctionSet() {
	dateGetDayFunctionOnce.Do(func() {
		dateGetDayFunction = gi.FunctionInvokerNew("Soup", "get_day")
	})
}

var getDayDateInvoker *gi.Function

// GetDay is a representation of the C type soup_date_get_day.
func (recv *Date) GetDay() int32 {
	if getDayDateInvoker == nil {
		getDayDateInvoker = gi.StructFunctionInvokerNew("Soup", "Date", "get_day")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getDayDateInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var dateGetHourFunction *gi.Function
var dateGetHourFunctionOnce sync.Once

func dateGetHourFunctionSet() {
	dateGetHourFunctionOnce.Do(func() {
		dateGetHourFunction = gi.FunctionInvokerNew("Soup", "get_hour")
	})
}

var getHourDateInvoker *gi.Function

// GetHour is a representation of the C type soup_date_get_hour.
func (recv *Date) GetHour() int32 {
	if getHourDateInvoker == nil {
		getHourDateInvoker = gi.StructFunctionInvokerNew("Soup", "Date", "get_hour")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getHourDateInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var dateGetMinuteFunction *gi.Function
var dateGetMinuteFunctionOnce sync.Once

func dateGetMinuteFunctionSet() {
	dateGetMinuteFunctionOnce.Do(func() {
		dateGetMinuteFunction = gi.FunctionInvokerNew("Soup", "get_minute")
	})
}

var getMinuteDateInvoker *gi.Function

// GetMinute is a representation of the C type soup_date_get_minute.
func (recv *Date) GetMinute() int32 {
	if getMinuteDateInvoker == nil {
		getMinuteDateInvoker = gi.StructFunctionInvokerNew("Soup", "Date", "get_minute")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getMinuteDateInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var dateGetMonthFunction *gi.Function
var dateGetMonthFunctionOnce sync.Once

func dateGetMonthFunctionSet() {
	dateGetMonthFunctionOnce.Do(func() {
		dateGetMonthFunction = gi.FunctionInvokerNew("Soup", "get_month")
	})
}

var getMonthDateInvoker *gi.Function

// GetMonth is a representation of the C type soup_date_get_month.
func (recv *Date) GetMonth() int32 {
	if getMonthDateInvoker == nil {
		getMonthDateInvoker = gi.StructFunctionInvokerNew("Soup", "Date", "get_month")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getMonthDateInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var dateGetOffsetFunction *gi.Function
var dateGetOffsetFunctionOnce sync.Once

func dateGetOffsetFunctionSet() {
	dateGetOffsetFunctionOnce.Do(func() {
		dateGetOffsetFunction = gi.FunctionInvokerNew("Soup", "get_offset")
	})
}

var getOffsetDateInvoker *gi.Function

// GetOffset is a representation of the C type soup_date_get_offset.
func (recv *Date) GetOffset() int32 {
	if getOffsetDateInvoker == nil {
		getOffsetDateInvoker = gi.StructFunctionInvokerNew("Soup", "Date", "get_offset")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getOffsetDateInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var dateGetSecondFunction *gi.Function
var dateGetSecondFunctionOnce sync.Once

func dateGetSecondFunctionSet() {
	dateGetSecondFunctionOnce.Do(func() {
		dateGetSecondFunction = gi.FunctionInvokerNew("Soup", "get_second")
	})
}

var getSecondDateInvoker *gi.Function

// GetSecond is a representation of the C type soup_date_get_second.
func (recv *Date) GetSecond() int32 {
	if getSecondDateInvoker == nil {
		getSecondDateInvoker = gi.StructFunctionInvokerNew("Soup", "Date", "get_second")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getSecondDateInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var dateGetUtcFunction *gi.Function
var dateGetUtcFunctionOnce sync.Once

func dateGetUtcFunctionSet() {
	dateGetUtcFunctionOnce.Do(func() {
		dateGetUtcFunction = gi.FunctionInvokerNew("Soup", "get_utc")
	})
}

var getUtcDateInvoker *gi.Function

// GetUtc is a representation of the C type soup_date_get_utc.
func (recv *Date) GetUtc() int32 {
	if getUtcDateInvoker == nil {
		getUtcDateInvoker = gi.StructFunctionInvokerNew("Soup", "Date", "get_utc")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getUtcDateInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var dateGetYearFunction *gi.Function
var dateGetYearFunctionOnce sync.Once

func dateGetYearFunctionSet() {
	dateGetYearFunctionOnce.Do(func() {
		dateGetYearFunction = gi.FunctionInvokerNew("Soup", "get_year")
	})
}

var getYearDateInvoker *gi.Function

// GetYear is a representation of the C type soup_date_get_year.
func (recv *Date) GetYear() int32 {
	if getYearDateInvoker == nil {
		getYearDateInvoker = gi.StructFunctionInvokerNew("Soup", "Date", "get_year")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getYearDateInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

// UNSUPPORTED : C value 'soup_date_is_past' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'soup_date_to_string' : parameter 'format' of type 'DateFormat' not supported

var dateToTimeTFunction *gi.Function
var dateToTimeTFunctionOnce sync.Once

func dateToTimeTFunctionSet() {
	dateToTimeTFunctionOnce.Do(func() {
		dateToTimeTFunction = gi.FunctionInvokerNew("Soup", "to_time_t")
	})
}

var toTimeTDateInvoker *gi.Function

// ToTimeT is a representation of the C type soup_date_to_time_t.
func (recv *Date) ToTimeT() int64 {
	if toTimeTDateInvoker == nil {
		toTimeTDateInvoker = gi.StructFunctionInvokerNew("Soup", "Date", "to_time_t")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := toTimeTDateInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int64()

	return retGo
}

// UNSUPPORTED : C value 'soup_date_to_timeval' : parameter 'time' of type 'GLib.TimeVal' not supported

var hSTSEnforcerClassStruct *gi.Struct
var hSTSEnforcerClassStructOnce sync.Once

func hSTSEnforcerClassStructSet() {
	hSTSEnforcerClassStructOnce.Do(func() {
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
var hSTSEnforcerDBClassStructOnce sync.Once

func hSTSEnforcerDBClassStructSet() {
	hSTSEnforcerDBClassStructOnce.Do(func() {
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
var hSTSEnforcerDBPrivateStructOnce sync.Once

func hSTSEnforcerDBPrivateStructSet() {
	hSTSEnforcerDBPrivateStructOnce.Do(func() {
		hSTSEnforcerDBPrivateStruct = gi.StructNew("Soup", "HSTSEnforcerDBPrivate")
	})
}

type HSTSEnforcerDBPrivate struct {
	native uintptr
}

var hSTSEnforcerPrivateStruct *gi.Struct
var hSTSEnforcerPrivateStructOnce sync.Once

func hSTSEnforcerPrivateStructSet() {
	hSTSEnforcerPrivateStructOnce.Do(func() {
		hSTSEnforcerPrivateStruct = gi.StructNew("Soup", "HSTSEnforcerPrivate")
	})
}

type HSTSEnforcerPrivate struct {
	native uintptr
}

var hSTSPolicyStruct *gi.Struct
var hSTSPolicyStructOnce sync.Once

func hSTSPolicyStructSet() {
	hSTSPolicyStructOnce.Do(func() {
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
var hSTSPolicyCopyFunctionOnce sync.Once

func hSTSPolicyCopyFunctionSet() {
	hSTSPolicyCopyFunctionOnce.Do(func() {
		hSTSPolicyCopyFunction = gi.FunctionInvokerNew("Soup", "copy")
	})
}

var copyHSTSPolicyInvoker *gi.Function

// Copy is a representation of the C type soup_hsts_policy_copy.
func (recv *HSTSPolicy) Copy() *HSTSPolicy {
	if copyHSTSPolicyInvoker == nil {
		copyHSTSPolicyInvoker = gi.StructFunctionInvokerNew("Soup", "HSTSPolicy", "copy")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := copyHSTSPolicyInvoker.Invoke(inArgs[:], nil)

	retGo := &HSTSPolicy{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'soup_hsts_policy_equal' : parameter 'policy2' of type 'HSTSPolicy' not supported

var hSTSPolicyFreeFunction *gi.Function
var hSTSPolicyFreeFunctionOnce sync.Once

func hSTSPolicyFreeFunctionSet() {
	hSTSPolicyFreeFunctionOnce.Do(func() {
		hSTSPolicyFreeFunction = gi.FunctionInvokerNew("Soup", "free")
	})
}

var freeHSTSPolicyInvoker *gi.Function

// Free is a representation of the C type soup_hsts_policy_free.
func (recv *HSTSPolicy) Free() {
	if freeHSTSPolicyInvoker == nil {
		freeHSTSPolicyInvoker = gi.StructFunctionInvokerNew("Soup", "HSTSPolicy", "free")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	freeHSTSPolicyInvoker.Invoke(inArgs[:], nil)

}

var hSTSPolicyGetDomainFunction *gi.Function
var hSTSPolicyGetDomainFunctionOnce sync.Once

func hSTSPolicyGetDomainFunctionSet() {
	hSTSPolicyGetDomainFunctionOnce.Do(func() {
		hSTSPolicyGetDomainFunction = gi.FunctionInvokerNew("Soup", "get_domain")
	})
}

var getDomainHSTSPolicyInvoker *gi.Function

// GetDomain is a representation of the C type soup_hsts_policy_get_domain.
func (recv *HSTSPolicy) GetDomain() string {
	if getDomainHSTSPolicyInvoker == nil {
		getDomainHSTSPolicyInvoker = gi.StructFunctionInvokerNew("Soup", "HSTSPolicy", "get_domain")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getDomainHSTSPolicyInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'soup_hsts_policy_includes_subdomains' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'soup_hsts_policy_is_expired' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'soup_hsts_policy_is_session_policy' : return type 'gboolean' not supported

var loggerClassStruct *gi.Struct
var loggerClassStructOnce sync.Once

func loggerClassStructSet() {
	loggerClassStructOnce.Do(func() {
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
var messageBodyStructOnce sync.Once

func messageBodyStructSet() {
	messageBodyStructOnce.Do(func() {
		messageBodyStruct = gi.StructNew("Soup", "MessageBody")
	})
}

type MessageBody struct {
	native uintptr
	Data   string
	Length int64
}

var messageBodyNewFunction *gi.Function
var messageBodyNewFunctionOnce sync.Once

func messageBodyNewFunctionSet() {
	messageBodyNewFunctionOnce.Do(func() {
		messageBodyNewFunction = gi.FunctionInvokerNew("Soup", "new")
	})
}

var newMessageBodyInvoker *gi.Function

// MessageBodyNew is a representation of the C type soup_message_body_new.
func MessageBodyNew() *MessageBody {
	if newMessageBodyInvoker == nil {
		newMessageBodyInvoker = gi.StructFunctionInvokerNew("Soup", "MessageBody", "new")
	}

	ret := newMessageBodyInvoker.Invoke(nil, nil)

	retGo := &MessageBody{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'soup_message_body_append' : parameter 'use' of type 'MemoryUse' not supported

// UNSUPPORTED : C value 'soup_message_body_append_buffer' : parameter 'buffer' of type 'Buffer' not supported

// UNSUPPORTED : C value 'soup_message_body_append_take' : parameter 'data' has no type

var messageBodyCompleteFunction *gi.Function
var messageBodyCompleteFunctionOnce sync.Once

func messageBodyCompleteFunctionSet() {
	messageBodyCompleteFunctionOnce.Do(func() {
		messageBodyCompleteFunction = gi.FunctionInvokerNew("Soup", "complete")
	})
}

var completeMessageBodyInvoker *gi.Function

// Complete is a representation of the C type soup_message_body_complete.
func (recv *MessageBody) Complete() {
	if completeMessageBodyInvoker == nil {
		completeMessageBodyInvoker = gi.StructFunctionInvokerNew("Soup", "MessageBody", "complete")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	completeMessageBodyInvoker.Invoke(inArgs[:], nil)

}

var messageBodyFlattenFunction *gi.Function
var messageBodyFlattenFunctionOnce sync.Once

func messageBodyFlattenFunctionSet() {
	messageBodyFlattenFunctionOnce.Do(func() {
		messageBodyFlattenFunction = gi.FunctionInvokerNew("Soup", "flatten")
	})
}

var flattenMessageBodyInvoker *gi.Function

// Flatten is a representation of the C type soup_message_body_flatten.
func (recv *MessageBody) Flatten() *Buffer {
	if flattenMessageBodyInvoker == nil {
		flattenMessageBodyInvoker = gi.StructFunctionInvokerNew("Soup", "MessageBody", "flatten")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := flattenMessageBodyInvoker.Invoke(inArgs[:], nil)

	retGo := &Buffer{native: ret.Pointer()}

	return retGo
}

var messageBodyFreeFunction *gi.Function
var messageBodyFreeFunctionOnce sync.Once

func messageBodyFreeFunctionSet() {
	messageBodyFreeFunctionOnce.Do(func() {
		messageBodyFreeFunction = gi.FunctionInvokerNew("Soup", "free")
	})
}

var freeMessageBodyInvoker *gi.Function

// Free is a representation of the C type soup_message_body_free.
func (recv *MessageBody) Free() {
	if freeMessageBodyInvoker == nil {
		freeMessageBodyInvoker = gi.StructFunctionInvokerNew("Soup", "MessageBody", "free")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	freeMessageBodyInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'soup_message_body_get_accumulate' : return type 'gboolean' not supported

var messageBodyGetChunkFunction *gi.Function
var messageBodyGetChunkFunctionOnce sync.Once

func messageBodyGetChunkFunctionSet() {
	messageBodyGetChunkFunctionOnce.Do(func() {
		messageBodyGetChunkFunction = gi.FunctionInvokerNew("Soup", "get_chunk")
	})
}

var getChunkMessageBodyInvoker *gi.Function

// GetChunk is a representation of the C type soup_message_body_get_chunk.
func (recv *MessageBody) GetChunk(offset int64) *Buffer {
	if getChunkMessageBodyInvoker == nil {
		getChunkMessageBodyInvoker = gi.StructFunctionInvokerNew("Soup", "MessageBody", "get_chunk")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt64(offset)

	ret := getChunkMessageBodyInvoker.Invoke(inArgs[:], nil)

	retGo := &Buffer{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'soup_message_body_got_chunk' : parameter 'chunk' of type 'Buffer' not supported

// UNSUPPORTED : C value 'soup_message_body_set_accumulate' : parameter 'accumulate' of type 'gboolean' not supported

var messageBodyTruncateFunction *gi.Function
var messageBodyTruncateFunctionOnce sync.Once

func messageBodyTruncateFunctionSet() {
	messageBodyTruncateFunctionOnce.Do(func() {
		messageBodyTruncateFunction = gi.FunctionInvokerNew("Soup", "truncate")
	})
}

var truncateMessageBodyInvoker *gi.Function

// Truncate is a representation of the C type soup_message_body_truncate.
func (recv *MessageBody) Truncate() {
	if truncateMessageBodyInvoker == nil {
		truncateMessageBodyInvoker = gi.StructFunctionInvokerNew("Soup", "MessageBody", "truncate")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	truncateMessageBodyInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'soup_message_body_wrote_chunk' : parameter 'chunk' of type 'Buffer' not supported

var messageClassStruct *gi.Struct
var messageClassStructOnce sync.Once

func messageClassStructSet() {
	messageClassStructOnce.Do(func() {
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
var messageHeadersStructOnce sync.Once

func messageHeadersStructSet() {
	messageHeadersStructOnce.Do(func() {
		messageHeadersStruct = gi.StructNew("Soup", "MessageHeaders")
	})
}

type MessageHeaders struct {
	native uintptr
}

// UNSUPPORTED : C value 'soup_message_headers_new' : parameter 'type' of type 'MessageHeadersType' not supported

var messageHeadersAppendFunction *gi.Function
var messageHeadersAppendFunctionOnce sync.Once

func messageHeadersAppendFunctionSet() {
	messageHeadersAppendFunctionOnce.Do(func() {
		messageHeadersAppendFunction = gi.FunctionInvokerNew("Soup", "append")
	})
}

var appendMessageHeadersInvoker *gi.Function

// Append is a representation of the C type soup_message_headers_append.
func (recv *MessageHeaders) Append(name string, value string) {
	if appendMessageHeadersInvoker == nil {
		appendMessageHeadersInvoker = gi.StructFunctionInvokerNew("Soup", "MessageHeaders", "append")
	}

	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(name)
	inArgs[2].SetString(value)

	appendMessageHeadersInvoker.Invoke(inArgs[:], nil)

}

var messageHeadersCleanConnectionHeadersFunction *gi.Function
var messageHeadersCleanConnectionHeadersFunctionOnce sync.Once

func messageHeadersCleanConnectionHeadersFunctionSet() {
	messageHeadersCleanConnectionHeadersFunctionOnce.Do(func() {
		messageHeadersCleanConnectionHeadersFunction = gi.FunctionInvokerNew("Soup", "clean_connection_headers")
	})
}

var cleanConnectionHeadersMessageHeadersInvoker *gi.Function

// CleanConnectionHeaders is a representation of the C type soup_message_headers_clean_connection_headers.
func (recv *MessageHeaders) CleanConnectionHeaders() {
	if cleanConnectionHeadersMessageHeadersInvoker == nil {
		cleanConnectionHeadersMessageHeadersInvoker = gi.StructFunctionInvokerNew("Soup", "MessageHeaders", "clean_connection_headers")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	cleanConnectionHeadersMessageHeadersInvoker.Invoke(inArgs[:], nil)

}

var messageHeadersClearFunction *gi.Function
var messageHeadersClearFunctionOnce sync.Once

func messageHeadersClearFunctionSet() {
	messageHeadersClearFunctionOnce.Do(func() {
		messageHeadersClearFunction = gi.FunctionInvokerNew("Soup", "clear")
	})
}

var clearMessageHeadersInvoker *gi.Function

// Clear is a representation of the C type soup_message_headers_clear.
func (recv *MessageHeaders) Clear() {
	if clearMessageHeadersInvoker == nil {
		clearMessageHeadersInvoker = gi.StructFunctionInvokerNew("Soup", "MessageHeaders", "clear")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	clearMessageHeadersInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'soup_message_headers_foreach' : parameter 'func' of type 'MessageHeadersForeachFunc' not supported

var messageHeadersFreeFunction *gi.Function
var messageHeadersFreeFunctionOnce sync.Once

func messageHeadersFreeFunctionSet() {
	messageHeadersFreeFunctionOnce.Do(func() {
		messageHeadersFreeFunction = gi.FunctionInvokerNew("Soup", "free")
	})
}

var freeMessageHeadersInvoker *gi.Function

// Free is a representation of the C type soup_message_headers_free.
func (recv *MessageHeaders) Free() {
	if freeMessageHeadersInvoker == nil {
		freeMessageHeadersInvoker = gi.StructFunctionInvokerNew("Soup", "MessageHeaders", "free")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	freeMessageHeadersInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'soup_message_headers_free_ranges' : parameter 'ranges' of type 'Range' not supported

var messageHeadersGetFunction *gi.Function
var messageHeadersGetFunctionOnce sync.Once

func messageHeadersGetFunctionSet() {
	messageHeadersGetFunctionOnce.Do(func() {
		messageHeadersGetFunction = gi.FunctionInvokerNew("Soup", "get")
	})
}

var getMessageHeadersInvoker *gi.Function

// Get is a representation of the C type soup_message_headers_get.
func (recv *MessageHeaders) Get(name string) string {
	if getMessageHeadersInvoker == nil {
		getMessageHeadersInvoker = gi.StructFunctionInvokerNew("Soup", "MessageHeaders", "get")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(name)

	ret := getMessageHeadersInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'soup_message_headers_get_content_disposition' : parameter 'params' of type 'GLib.HashTable' not supported

var messageHeadersGetContentLengthFunction *gi.Function
var messageHeadersGetContentLengthFunctionOnce sync.Once

func messageHeadersGetContentLengthFunctionSet() {
	messageHeadersGetContentLengthFunctionOnce.Do(func() {
		messageHeadersGetContentLengthFunction = gi.FunctionInvokerNew("Soup", "get_content_length")
	})
}

var getContentLengthMessageHeadersInvoker *gi.Function

// GetContentLength is a representation of the C type soup_message_headers_get_content_length.
func (recv *MessageHeaders) GetContentLength() int64 {
	if getContentLengthMessageHeadersInvoker == nil {
		getContentLengthMessageHeadersInvoker = gi.StructFunctionInvokerNew("Soup", "MessageHeaders", "get_content_length")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getContentLengthMessageHeadersInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int64()

	return retGo
}

// UNSUPPORTED : C value 'soup_message_headers_get_content_range' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'soup_message_headers_get_content_type' : parameter 'params' of type 'GLib.HashTable' not supported

// UNSUPPORTED : C value 'soup_message_headers_get_encoding' : return type 'Encoding' not supported

// UNSUPPORTED : C value 'soup_message_headers_get_expectations' : return type 'Expectation' not supported

// UNSUPPORTED : C value 'soup_message_headers_get_headers_type' : return type 'MessageHeadersType' not supported

var messageHeadersGetListFunction *gi.Function
var messageHeadersGetListFunctionOnce sync.Once

func messageHeadersGetListFunctionSet() {
	messageHeadersGetListFunctionOnce.Do(func() {
		messageHeadersGetListFunction = gi.FunctionInvokerNew("Soup", "get_list")
	})
}

var getListMessageHeadersInvoker *gi.Function

// GetList is a representation of the C type soup_message_headers_get_list.
func (recv *MessageHeaders) GetList(name string) string {
	if getListMessageHeadersInvoker == nil {
		getListMessageHeadersInvoker = gi.StructFunctionInvokerNew("Soup", "MessageHeaders", "get_list")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(name)

	ret := getListMessageHeadersInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

var messageHeadersGetOneFunction *gi.Function
var messageHeadersGetOneFunctionOnce sync.Once

func messageHeadersGetOneFunctionSet() {
	messageHeadersGetOneFunctionOnce.Do(func() {
		messageHeadersGetOneFunction = gi.FunctionInvokerNew("Soup", "get_one")
	})
}

var getOneMessageHeadersInvoker *gi.Function

// GetOne is a representation of the C type soup_message_headers_get_one.
func (recv *MessageHeaders) GetOne(name string) string {
	if getOneMessageHeadersInvoker == nil {
		getOneMessageHeadersInvoker = gi.StructFunctionInvokerNew("Soup", "MessageHeaders", "get_one")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(name)

	ret := getOneMessageHeadersInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'soup_message_headers_get_ranges' : parameter 'ranges' has no type

// UNSUPPORTED : C value 'soup_message_headers_header_contains' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'soup_message_headers_header_equals' : return type 'gboolean' not supported

var messageHeadersRemoveFunction *gi.Function
var messageHeadersRemoveFunctionOnce sync.Once

func messageHeadersRemoveFunctionSet() {
	messageHeadersRemoveFunctionOnce.Do(func() {
		messageHeadersRemoveFunction = gi.FunctionInvokerNew("Soup", "remove")
	})
}

var removeMessageHeadersInvoker *gi.Function

// Remove is a representation of the C type soup_message_headers_remove.
func (recv *MessageHeaders) Remove(name string) {
	if removeMessageHeadersInvoker == nil {
		removeMessageHeadersInvoker = gi.StructFunctionInvokerNew("Soup", "MessageHeaders", "remove")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(name)

	removeMessageHeadersInvoker.Invoke(inArgs[:], nil)

}

var messageHeadersReplaceFunction *gi.Function
var messageHeadersReplaceFunctionOnce sync.Once

func messageHeadersReplaceFunctionSet() {
	messageHeadersReplaceFunctionOnce.Do(func() {
		messageHeadersReplaceFunction = gi.FunctionInvokerNew("Soup", "replace")
	})
}

var replaceMessageHeadersInvoker *gi.Function

// Replace is a representation of the C type soup_message_headers_replace.
func (recv *MessageHeaders) Replace(name string, value string) {
	if replaceMessageHeadersInvoker == nil {
		replaceMessageHeadersInvoker = gi.StructFunctionInvokerNew("Soup", "MessageHeaders", "replace")
	}

	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(name)
	inArgs[2].SetString(value)

	replaceMessageHeadersInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'soup_message_headers_set_content_disposition' : parameter 'params' of type 'GLib.HashTable' not supported

var messageHeadersSetContentLengthFunction *gi.Function
var messageHeadersSetContentLengthFunctionOnce sync.Once

func messageHeadersSetContentLengthFunctionSet() {
	messageHeadersSetContentLengthFunctionOnce.Do(func() {
		messageHeadersSetContentLengthFunction = gi.FunctionInvokerNew("Soup", "set_content_length")
	})
}

var setContentLengthMessageHeadersInvoker *gi.Function

// SetContentLength is a representation of the C type soup_message_headers_set_content_length.
func (recv *MessageHeaders) SetContentLength(contentLength int64) {
	if setContentLengthMessageHeadersInvoker == nil {
		setContentLengthMessageHeadersInvoker = gi.StructFunctionInvokerNew("Soup", "MessageHeaders", "set_content_length")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt64(contentLength)

	setContentLengthMessageHeadersInvoker.Invoke(inArgs[:], nil)

}

var messageHeadersSetContentRangeFunction *gi.Function
var messageHeadersSetContentRangeFunctionOnce sync.Once

func messageHeadersSetContentRangeFunctionSet() {
	messageHeadersSetContentRangeFunctionOnce.Do(func() {
		messageHeadersSetContentRangeFunction = gi.FunctionInvokerNew("Soup", "set_content_range")
	})
}

var setContentRangeMessageHeadersInvoker *gi.Function

// SetContentRange is a representation of the C type soup_message_headers_set_content_range.
func (recv *MessageHeaders) SetContentRange(start int64, end int64, totalLength int64) {
	if setContentRangeMessageHeadersInvoker == nil {
		setContentRangeMessageHeadersInvoker = gi.StructFunctionInvokerNew("Soup", "MessageHeaders", "set_content_range")
	}

	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt64(start)
	inArgs[2].SetInt64(end)
	inArgs[3].SetInt64(totalLength)

	setContentRangeMessageHeadersInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'soup_message_headers_set_content_type' : parameter 'params' of type 'GLib.HashTable' not supported

// UNSUPPORTED : C value 'soup_message_headers_set_encoding' : parameter 'encoding' of type 'Encoding' not supported

// UNSUPPORTED : C value 'soup_message_headers_set_expectations' : parameter 'expectations' of type 'Expectation' not supported

var messageHeadersSetRangeFunction *gi.Function
var messageHeadersSetRangeFunctionOnce sync.Once

func messageHeadersSetRangeFunctionSet() {
	messageHeadersSetRangeFunctionOnce.Do(func() {
		messageHeadersSetRangeFunction = gi.FunctionInvokerNew("Soup", "set_range")
	})
}

var setRangeMessageHeadersInvoker *gi.Function

// SetRange is a representation of the C type soup_message_headers_set_range.
func (recv *MessageHeaders) SetRange(start int64, end int64) {
	if setRangeMessageHeadersInvoker == nil {
		setRangeMessageHeadersInvoker = gi.StructFunctionInvokerNew("Soup", "MessageHeaders", "set_range")
	}

	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt64(start)
	inArgs[2].SetInt64(end)

	setRangeMessageHeadersInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'soup_message_headers_set_ranges' : parameter 'ranges' of type 'Range' not supported

var messageHeadersIterStruct *gi.Struct
var messageHeadersIterStructOnce sync.Once

func messageHeadersIterStructSet() {
	messageHeadersIterStructOnce.Do(func() {
		messageHeadersIterStruct = gi.StructNew("Soup", "MessageHeadersIter")
	})
}

type MessageHeadersIter struct {
	native uintptr
}

// UNSUPPORTED : C value 'soup_message_headers_iter_next' : return type 'gboolean' not supported

var messageQueueStruct *gi.Struct
var messageQueueStructOnce sync.Once

func messageQueueStructSet() {
	messageQueueStructOnce.Do(func() {
		messageQueueStruct = gi.StructNew("Soup", "MessageQueue")
	})
}

type MessageQueue struct {
	native uintptr
}

var messageQueueItemStruct *gi.Struct
var messageQueueItemStructOnce sync.Once

func messageQueueItemStructSet() {
	messageQueueItemStructOnce.Do(func() {
		messageQueueItemStruct = gi.StructNew("Soup", "MessageQueueItem")
	})
}

type MessageQueueItem struct {
	native uintptr
}

var multipartStruct *gi.Struct
var multipartStructOnce sync.Once

func multipartStructSet() {
	multipartStructOnce.Do(func() {
		multipartStruct = gi.StructNew("Soup", "Multipart")
	})
}

type Multipart struct {
	native uintptr
}

var multipartNewFunction *gi.Function
var multipartNewFunctionOnce sync.Once

func multipartNewFunctionSet() {
	multipartNewFunctionOnce.Do(func() {
		multipartNewFunction = gi.FunctionInvokerNew("Soup", "new")
	})
}

var newMultipartInvoker *gi.Function

// MultipartNew is a representation of the C type soup_multipart_new.
func MultipartNew(mimeType string) *Multipart {
	if newMultipartInvoker == nil {
		newMultipartInvoker = gi.StructFunctionInvokerNew("Soup", "Multipart", "new")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetString(mimeType)

	ret := newMultipartInvoker.Invoke(inArgs[:], nil)

	retGo := &Multipart{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'soup_multipart_new_from_message' : parameter 'headers' of type 'MessageHeaders' not supported

// UNSUPPORTED : C value 'soup_multipart_append_form_file' : parameter 'body' of type 'Buffer' not supported

var multipartAppendFormStringFunction *gi.Function
var multipartAppendFormStringFunctionOnce sync.Once

func multipartAppendFormStringFunctionSet() {
	multipartAppendFormStringFunctionOnce.Do(func() {
		multipartAppendFormStringFunction = gi.FunctionInvokerNew("Soup", "append_form_string")
	})
}

var appendFormStringMultipartInvoker *gi.Function

// AppendFormString is a representation of the C type soup_multipart_append_form_string.
func (recv *Multipart) AppendFormString(controlName string, data string) {
	if appendFormStringMultipartInvoker == nil {
		appendFormStringMultipartInvoker = gi.StructFunctionInvokerNew("Soup", "Multipart", "append_form_string")
	}

	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(controlName)
	inArgs[2].SetString(data)

	appendFormStringMultipartInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'soup_multipart_append_part' : parameter 'headers' of type 'MessageHeaders' not supported

var multipartFreeFunction *gi.Function
var multipartFreeFunctionOnce sync.Once

func multipartFreeFunctionSet() {
	multipartFreeFunctionOnce.Do(func() {
		multipartFreeFunction = gi.FunctionInvokerNew("Soup", "free")
	})
}

var freeMultipartInvoker *gi.Function

// Free is a representation of the C type soup_multipart_free.
func (recv *Multipart) Free() {
	if freeMultipartInvoker == nil {
		freeMultipartInvoker = gi.StructFunctionInvokerNew("Soup", "Multipart", "free")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	freeMultipartInvoker.Invoke(inArgs[:], nil)

}

var multipartGetLengthFunction *gi.Function
var multipartGetLengthFunctionOnce sync.Once

func multipartGetLengthFunctionSet() {
	multipartGetLengthFunctionOnce.Do(func() {
		multipartGetLengthFunction = gi.FunctionInvokerNew("Soup", "get_length")
	})
}

var getLengthMultipartInvoker *gi.Function

// GetLength is a representation of the C type soup_multipart_get_length.
func (recv *Multipart) GetLength() int32 {
	if getLengthMultipartInvoker == nil {
		getLengthMultipartInvoker = gi.StructFunctionInvokerNew("Soup", "Multipart", "get_length")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getLengthMultipartInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

// UNSUPPORTED : C value 'soup_multipart_get_part' : parameter 'headers' of type 'MessageHeaders' not supported

// UNSUPPORTED : C value 'soup_multipart_to_message' : parameter 'dest_headers' of type 'MessageHeaders' not supported

var multipartInputStreamClassStruct *gi.Struct
var multipartInputStreamClassStructOnce sync.Once

func multipartInputStreamClassStructSet() {
	multipartInputStreamClassStructOnce.Do(func() {
		multipartInputStreamClassStruct = gi.StructNew("Soup", "MultipartInputStreamClass")
	})
}

type MultipartInputStreamClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'Gio.FilterInputStreamClass'
}

var multipartInputStreamPrivateStruct *gi.Struct
var multipartInputStreamPrivateStructOnce sync.Once

func multipartInputStreamPrivateStructSet() {
	multipartInputStreamPrivateStructOnce.Do(func() {
		multipartInputStreamPrivateStruct = gi.StructNew("Soup", "MultipartInputStreamPrivate")
	})
}

type MultipartInputStreamPrivate struct {
	native uintptr
}

var passwordManagerInterfaceStruct *gi.Struct
var passwordManagerInterfaceStructOnce sync.Once

func passwordManagerInterfaceStructSet() {
	passwordManagerInterfaceStructOnce.Do(func() {
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
var proxyResolverDefaultClassStructOnce sync.Once

func proxyResolverDefaultClassStructSet() {
	proxyResolverDefaultClassStructOnce.Do(func() {
		proxyResolverDefaultClassStruct = gi.StructNew("Soup", "ProxyResolverDefaultClass")
	})
}

type ProxyResolverDefaultClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
}

var proxyResolverInterfaceStruct *gi.Struct
var proxyResolverInterfaceStructOnce sync.Once

func proxyResolverInterfaceStructSet() {
	proxyResolverInterfaceStructOnce.Do(func() {
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
var proxyURIResolverInterfaceStructOnce sync.Once

func proxyURIResolverInterfaceStructSet() {
	proxyURIResolverInterfaceStructOnce.Do(func() {
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
var rangeStructOnce sync.Once

func rangeStructSet() {
	rangeStructOnce.Do(func() {
		rangeStruct = gi.StructNew("Soup", "Range")
	})
}

type Range struct {
	native uintptr
	Start  int64
	End    int64
}

var requestClassStruct *gi.Struct
var requestClassStructOnce sync.Once

func requestClassStructSet() {
	requestClassStructOnce.Do(func() {
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
var requestDataClassStructOnce sync.Once

func requestDataClassStructSet() {
	requestDataClassStructOnce.Do(func() {
		requestDataClassStruct = gi.StructNew("Soup", "RequestDataClass")
	})
}

type RequestDataClass struct {
	native uintptr
	Parent *RequestClass
}

var requestDataPrivateStruct *gi.Struct
var requestDataPrivateStructOnce sync.Once

func requestDataPrivateStructSet() {
	requestDataPrivateStructOnce.Do(func() {
		requestDataPrivateStruct = gi.StructNew("Soup", "RequestDataPrivate")
	})
}

type RequestDataPrivate struct {
	native uintptr
}

var requestFileClassStruct *gi.Struct
var requestFileClassStructOnce sync.Once

func requestFileClassStructSet() {
	requestFileClassStructOnce.Do(func() {
		requestFileClassStruct = gi.StructNew("Soup", "RequestFileClass")
	})
}

type RequestFileClass struct {
	native uintptr
	Parent *RequestClass
}

var requestFilePrivateStruct *gi.Struct
var requestFilePrivateStructOnce sync.Once

func requestFilePrivateStructSet() {
	requestFilePrivateStructOnce.Do(func() {
		requestFilePrivateStruct = gi.StructNew("Soup", "RequestFilePrivate")
	})
}

type RequestFilePrivate struct {
	native uintptr
}

var requestHTTPClassStruct *gi.Struct
var requestHTTPClassStructOnce sync.Once

func requestHTTPClassStructSet() {
	requestHTTPClassStructOnce.Do(func() {
		requestHTTPClassStruct = gi.StructNew("Soup", "RequestHTTPClass")
	})
}

type RequestHTTPClass struct {
	native uintptr
	Parent *RequestClass
}

var requestHTTPPrivateStruct *gi.Struct
var requestHTTPPrivateStructOnce sync.Once

func requestHTTPPrivateStructSet() {
	requestHTTPPrivateStructOnce.Do(func() {
		requestHTTPPrivateStruct = gi.StructNew("Soup", "RequestHTTPPrivate")
	})
}

type RequestHTTPPrivate struct {
	native uintptr
}

var requestPrivateStruct *gi.Struct
var requestPrivateStructOnce sync.Once

func requestPrivateStructSet() {
	requestPrivateStructOnce.Do(func() {
		requestPrivateStruct = gi.StructNew("Soup", "RequestPrivate")
	})
}

type RequestPrivate struct {
	native uintptr
}

var requesterClassStruct *gi.Struct
var requesterClassStructOnce sync.Once

func requesterClassStructSet() {
	requesterClassStructOnce.Do(func() {
		requesterClassStruct = gi.StructNew("Soup", "RequesterClass")
	})
}

type RequesterClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
}

var requesterPrivateStruct *gi.Struct
var requesterPrivateStructOnce sync.Once

func requesterPrivateStructSet() {
	requesterPrivateStructOnce.Do(func() {
		requesterPrivateStruct = gi.StructNew("Soup", "RequesterPrivate")
	})
}

type RequesterPrivate struct {
	native uintptr
}

var serverClassStruct *gi.Struct
var serverClassStructOnce sync.Once

func serverClassStructSet() {
	serverClassStructOnce.Do(func() {
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
var sessionAsyncClassStructOnce sync.Once

func sessionAsyncClassStructSet() {
	sessionAsyncClassStructOnce.Do(func() {
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
var sessionClassStructOnce sync.Once

func sessionClassStructSet() {
	sessionClassStructOnce.Do(func() {
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
var sessionFeatureInterfaceStructOnce sync.Once

func sessionFeatureInterfaceStructSet() {
	sessionFeatureInterfaceStructOnce.Do(func() {
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
var sessionSyncClassStructOnce sync.Once

func sessionSyncClassStructSet() {
	sessionSyncClassStructOnce.Do(func() {
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
var socketClassStructOnce sync.Once

func socketClassStructSet() {
	socketClassStructOnce.Do(func() {
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
var uRIStructOnce sync.Once

func uRIStructSet() {
	uRIStructOnce.Do(func() {
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
var uRINewFunctionOnce sync.Once

func uRINewFunctionSet() {
	uRINewFunctionOnce.Do(func() {
		uRINewFunction = gi.FunctionInvokerNew("Soup", "new")
	})
}

var newURIInvoker *gi.Function

// URINew is a representation of the C type soup_uri_new.
func URINew(uriString string) *URI {
	if newURIInvoker == nil {
		newURIInvoker = gi.StructFunctionInvokerNew("Soup", "URI", "new")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetString(uriString)

	ret := newURIInvoker.Invoke(inArgs[:], nil)

	retGo := &URI{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'soup_uri_new_with_base' : parameter 'base' of type 'URI' not supported

var uRICopyFunction *gi.Function
var uRICopyFunctionOnce sync.Once

func uRICopyFunctionSet() {
	uRICopyFunctionOnce.Do(func() {
		uRICopyFunction = gi.FunctionInvokerNew("Soup", "copy")
	})
}

var copyURIInvoker *gi.Function

// Copy is a representation of the C type soup_uri_copy.
func (recv *URI) Copy() *URI {
	if copyURIInvoker == nil {
		copyURIInvoker = gi.StructFunctionInvokerNew("Soup", "URI", "copy")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := copyURIInvoker.Invoke(inArgs[:], nil)

	retGo := &URI{native: ret.Pointer()}

	return retGo
}

var uRICopyHostFunction *gi.Function
var uRICopyHostFunctionOnce sync.Once

func uRICopyHostFunctionSet() {
	uRICopyHostFunctionOnce.Do(func() {
		uRICopyHostFunction = gi.FunctionInvokerNew("Soup", "copy_host")
	})
}

var copyHostURIInvoker *gi.Function

// CopyHost is a representation of the C type soup_uri_copy_host.
func (recv *URI) CopyHost() *URI {
	if copyHostURIInvoker == nil {
		copyHostURIInvoker = gi.StructFunctionInvokerNew("Soup", "URI", "copy_host")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := copyHostURIInvoker.Invoke(inArgs[:], nil)

	retGo := &URI{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'soup_uri_equal' : parameter 'uri2' of type 'URI' not supported

var uRIFreeFunction *gi.Function
var uRIFreeFunctionOnce sync.Once

func uRIFreeFunctionSet() {
	uRIFreeFunctionOnce.Do(func() {
		uRIFreeFunction = gi.FunctionInvokerNew("Soup", "free")
	})
}

var freeURIInvoker *gi.Function

// Free is a representation of the C type soup_uri_free.
func (recv *URI) Free() {
	if freeURIInvoker == nil {
		freeURIInvoker = gi.StructFunctionInvokerNew("Soup", "URI", "free")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	freeURIInvoker.Invoke(inArgs[:], nil)

}

var uRIGetFragmentFunction *gi.Function
var uRIGetFragmentFunctionOnce sync.Once

func uRIGetFragmentFunctionSet() {
	uRIGetFragmentFunctionOnce.Do(func() {
		uRIGetFragmentFunction = gi.FunctionInvokerNew("Soup", "get_fragment")
	})
}

var getFragmentURIInvoker *gi.Function

// GetFragment is a representation of the C type soup_uri_get_fragment.
func (recv *URI) GetFragment() string {
	if getFragmentURIInvoker == nil {
		getFragmentURIInvoker = gi.StructFunctionInvokerNew("Soup", "URI", "get_fragment")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getFragmentURIInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

var uRIGetHostFunction *gi.Function
var uRIGetHostFunctionOnce sync.Once

func uRIGetHostFunctionSet() {
	uRIGetHostFunctionOnce.Do(func() {
		uRIGetHostFunction = gi.FunctionInvokerNew("Soup", "get_host")
	})
}

var getHostURIInvoker *gi.Function

// GetHost is a representation of the C type soup_uri_get_host.
func (recv *URI) GetHost() string {
	if getHostURIInvoker == nil {
		getHostURIInvoker = gi.StructFunctionInvokerNew("Soup", "URI", "get_host")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getHostURIInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

var uRIGetPasswordFunction *gi.Function
var uRIGetPasswordFunctionOnce sync.Once

func uRIGetPasswordFunctionSet() {
	uRIGetPasswordFunctionOnce.Do(func() {
		uRIGetPasswordFunction = gi.FunctionInvokerNew("Soup", "get_password")
	})
}

var getPasswordURIInvoker *gi.Function

// GetPassword is a representation of the C type soup_uri_get_password.
func (recv *URI) GetPassword() string {
	if getPasswordURIInvoker == nil {
		getPasswordURIInvoker = gi.StructFunctionInvokerNew("Soup", "URI", "get_password")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getPasswordURIInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

var uRIGetPathFunction *gi.Function
var uRIGetPathFunctionOnce sync.Once

func uRIGetPathFunctionSet() {
	uRIGetPathFunctionOnce.Do(func() {
		uRIGetPathFunction = gi.FunctionInvokerNew("Soup", "get_path")
	})
}

var getPathURIInvoker *gi.Function

// GetPath is a representation of the C type soup_uri_get_path.
func (recv *URI) GetPath() string {
	if getPathURIInvoker == nil {
		getPathURIInvoker = gi.StructFunctionInvokerNew("Soup", "URI", "get_path")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getPathURIInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

var uRIGetPortFunction *gi.Function
var uRIGetPortFunctionOnce sync.Once

func uRIGetPortFunctionSet() {
	uRIGetPortFunctionOnce.Do(func() {
		uRIGetPortFunction = gi.FunctionInvokerNew("Soup", "get_port")
	})
}

var getPortURIInvoker *gi.Function

// GetPort is a representation of the C type soup_uri_get_port.
func (recv *URI) GetPort() uint32 {
	if getPortURIInvoker == nil {
		getPortURIInvoker = gi.StructFunctionInvokerNew("Soup", "URI", "get_port")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getPortURIInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Uint32()

	return retGo
}

var uRIGetQueryFunction *gi.Function
var uRIGetQueryFunctionOnce sync.Once

func uRIGetQueryFunctionSet() {
	uRIGetQueryFunctionOnce.Do(func() {
		uRIGetQueryFunction = gi.FunctionInvokerNew("Soup", "get_query")
	})
}

var getQueryURIInvoker *gi.Function

// GetQuery is a representation of the C type soup_uri_get_query.
func (recv *URI) GetQuery() string {
	if getQueryURIInvoker == nil {
		getQueryURIInvoker = gi.StructFunctionInvokerNew("Soup", "URI", "get_query")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getQueryURIInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

var uRIGetSchemeFunction *gi.Function
var uRIGetSchemeFunctionOnce sync.Once

func uRIGetSchemeFunctionSet() {
	uRIGetSchemeFunctionOnce.Do(func() {
		uRIGetSchemeFunction = gi.FunctionInvokerNew("Soup", "get_scheme")
	})
}

var getSchemeURIInvoker *gi.Function

// GetScheme is a representation of the C type soup_uri_get_scheme.
func (recv *URI) GetScheme() string {
	if getSchemeURIInvoker == nil {
		getSchemeURIInvoker = gi.StructFunctionInvokerNew("Soup", "URI", "get_scheme")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getSchemeURIInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

var uRIGetUserFunction *gi.Function
var uRIGetUserFunctionOnce sync.Once

func uRIGetUserFunctionSet() {
	uRIGetUserFunctionOnce.Do(func() {
		uRIGetUserFunction = gi.FunctionInvokerNew("Soup", "get_user")
	})
}

var getUserURIInvoker *gi.Function

// GetUser is a representation of the C type soup_uri_get_user.
func (recv *URI) GetUser() string {
	if getUserURIInvoker == nil {
		getUserURIInvoker = gi.StructFunctionInvokerNew("Soup", "URI", "get_user")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getUserURIInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'soup_uri_host_equal' : parameter 'v2' of type 'URI' not supported

var uRIHostHashFunction *gi.Function
var uRIHostHashFunctionOnce sync.Once

func uRIHostHashFunctionSet() {
	uRIHostHashFunctionOnce.Do(func() {
		uRIHostHashFunction = gi.FunctionInvokerNew("Soup", "host_hash")
	})
}

var hostHashURIInvoker *gi.Function

// HostHash is a representation of the C type soup_uri_host_hash.
func (recv *URI) HostHash() uint32 {
	if hostHashURIInvoker == nil {
		hostHashURIInvoker = gi.StructFunctionInvokerNew("Soup", "URI", "host_hash")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := hostHashURIInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Uint32()

	return retGo
}

var uRISetFragmentFunction *gi.Function
var uRISetFragmentFunctionOnce sync.Once

func uRISetFragmentFunctionSet() {
	uRISetFragmentFunctionOnce.Do(func() {
		uRISetFragmentFunction = gi.FunctionInvokerNew("Soup", "set_fragment")
	})
}

var setFragmentURIInvoker *gi.Function

// SetFragment is a representation of the C type soup_uri_set_fragment.
func (recv *URI) SetFragment(fragment string) {
	if setFragmentURIInvoker == nil {
		setFragmentURIInvoker = gi.StructFunctionInvokerNew("Soup", "URI", "set_fragment")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(fragment)

	setFragmentURIInvoker.Invoke(inArgs[:], nil)

}

var uRISetHostFunction *gi.Function
var uRISetHostFunctionOnce sync.Once

func uRISetHostFunctionSet() {
	uRISetHostFunctionOnce.Do(func() {
		uRISetHostFunction = gi.FunctionInvokerNew("Soup", "set_host")
	})
}

var setHostURIInvoker *gi.Function

// SetHost is a representation of the C type soup_uri_set_host.
func (recv *URI) SetHost(host string) {
	if setHostURIInvoker == nil {
		setHostURIInvoker = gi.StructFunctionInvokerNew("Soup", "URI", "set_host")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(host)

	setHostURIInvoker.Invoke(inArgs[:], nil)

}

var uRISetPasswordFunction *gi.Function
var uRISetPasswordFunctionOnce sync.Once

func uRISetPasswordFunctionSet() {
	uRISetPasswordFunctionOnce.Do(func() {
		uRISetPasswordFunction = gi.FunctionInvokerNew("Soup", "set_password")
	})
}

var setPasswordURIInvoker *gi.Function

// SetPassword is a representation of the C type soup_uri_set_password.
func (recv *URI) SetPassword(password string) {
	if setPasswordURIInvoker == nil {
		setPasswordURIInvoker = gi.StructFunctionInvokerNew("Soup", "URI", "set_password")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(password)

	setPasswordURIInvoker.Invoke(inArgs[:], nil)

}

var uRISetPathFunction *gi.Function
var uRISetPathFunctionOnce sync.Once

func uRISetPathFunctionSet() {
	uRISetPathFunctionOnce.Do(func() {
		uRISetPathFunction = gi.FunctionInvokerNew("Soup", "set_path")
	})
}

var setPathURIInvoker *gi.Function

// SetPath is a representation of the C type soup_uri_set_path.
func (recv *URI) SetPath(path string) {
	if setPathURIInvoker == nil {
		setPathURIInvoker = gi.StructFunctionInvokerNew("Soup", "URI", "set_path")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(path)

	setPathURIInvoker.Invoke(inArgs[:], nil)

}

var uRISetPortFunction *gi.Function
var uRISetPortFunctionOnce sync.Once

func uRISetPortFunctionSet() {
	uRISetPortFunctionOnce.Do(func() {
		uRISetPortFunction = gi.FunctionInvokerNew("Soup", "set_port")
	})
}

var setPortURIInvoker *gi.Function

// SetPort is a representation of the C type soup_uri_set_port.
func (recv *URI) SetPort(port uint32) {
	if setPortURIInvoker == nil {
		setPortURIInvoker = gi.StructFunctionInvokerNew("Soup", "URI", "set_port")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetUint32(port)

	setPortURIInvoker.Invoke(inArgs[:], nil)

}

var uRISetQueryFunction *gi.Function
var uRISetQueryFunctionOnce sync.Once

func uRISetQueryFunctionSet() {
	uRISetQueryFunctionOnce.Do(func() {
		uRISetQueryFunction = gi.FunctionInvokerNew("Soup", "set_query")
	})
}

var setQueryURIInvoker *gi.Function

// SetQuery is a representation of the C type soup_uri_set_query.
func (recv *URI) SetQuery(query string) {
	if setQueryURIInvoker == nil {
		setQueryURIInvoker = gi.StructFunctionInvokerNew("Soup", "URI", "set_query")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(query)

	setQueryURIInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'soup_uri_set_query_from_fields' : parameter '...' has no type

// UNSUPPORTED : C value 'soup_uri_set_query_from_form' : parameter 'form' of type 'GLib.HashTable' not supported

var uRISetSchemeFunction *gi.Function
var uRISetSchemeFunctionOnce sync.Once

func uRISetSchemeFunctionSet() {
	uRISetSchemeFunctionOnce.Do(func() {
		uRISetSchemeFunction = gi.FunctionInvokerNew("Soup", "set_scheme")
	})
}

var setSchemeURIInvoker *gi.Function

// SetScheme is a representation of the C type soup_uri_set_scheme.
func (recv *URI) SetScheme(scheme string) {
	if setSchemeURIInvoker == nil {
		setSchemeURIInvoker = gi.StructFunctionInvokerNew("Soup", "URI", "set_scheme")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(scheme)

	setSchemeURIInvoker.Invoke(inArgs[:], nil)

}

var uRISetUserFunction *gi.Function
var uRISetUserFunctionOnce sync.Once

func uRISetUserFunctionSet() {
	uRISetUserFunctionOnce.Do(func() {
		uRISetUserFunction = gi.FunctionInvokerNew("Soup", "set_user")
	})
}

var setUserURIInvoker *gi.Function

// SetUser is a representation of the C type soup_uri_set_user.
func (recv *URI) SetUser(user string) {
	if setUserURIInvoker == nil {
		setUserURIInvoker = gi.StructFunctionInvokerNew("Soup", "URI", "set_user")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(user)

	setUserURIInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'soup_uri_to_string' : parameter 'just_path_and_query' of type 'gboolean' not supported

// UNSUPPORTED : C value 'soup_uri_uses_default_port' : return type 'gboolean' not supported

var websocketConnectionClassStruct *gi.Struct
var websocketConnectionClassStructOnce sync.Once

func websocketConnectionClassStructSet() {
	websocketConnectionClassStructOnce.Do(func() {
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
var websocketConnectionPrivateStructOnce sync.Once

func websocketConnectionPrivateStructSet() {
	websocketConnectionPrivateStructOnce.Do(func() {
		websocketConnectionPrivateStruct = gi.StructNew("Soup", "WebsocketConnectionPrivate")
	})
}

type WebsocketConnectionPrivate struct {
	native uintptr
}

var websocketExtensionClassStruct *gi.Struct
var websocketExtensionClassStructOnce sync.Once

func websocketExtensionClassStructSet() {
	websocketExtensionClassStructOnce.Do(func() {
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
var websocketExtensionDeflateClassStructOnce sync.Once

func websocketExtensionDeflateClassStructSet() {
	websocketExtensionDeflateClassStructOnce.Do(func() {
		websocketExtensionDeflateClassStruct = gi.StructNew("Soup", "WebsocketExtensionDeflateClass")
	})
}

type WebsocketExtensionDeflateClass struct {
	native      uintptr
	ParentClass *WebsocketExtensionClass
}

var websocketExtensionManagerClassStruct *gi.Struct
var websocketExtensionManagerClassStructOnce sync.Once

func websocketExtensionManagerClassStructSet() {
	websocketExtensionManagerClassStructOnce.Do(func() {
		websocketExtensionManagerClassStruct = gi.StructNew("Soup", "WebsocketExtensionManagerClass")
	})
}

type WebsocketExtensionManagerClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
}

var xMLRPCParamsStruct *gi.Struct
var xMLRPCParamsStructOnce sync.Once

func xMLRPCParamsStructSet() {
	xMLRPCParamsStructOnce.Do(func() {
		xMLRPCParamsStruct = gi.StructNew("Soup", "XMLRPCParams")
	})
}

type XMLRPCParams struct {
	native uintptr
}

var xMLRPCParamsFreeFunction *gi.Function
var xMLRPCParamsFreeFunctionOnce sync.Once

func xMLRPCParamsFreeFunctionSet() {
	xMLRPCParamsFreeFunctionOnce.Do(func() {
		xMLRPCParamsFreeFunction = gi.FunctionInvokerNew("Soup", "free")
	})
}

var freeXMLRPCParamsInvoker *gi.Function

// Free is a representation of the C type soup_xmlrpc_params_free.
func (recv *XMLRPCParams) Free() {
	if freeXMLRPCParamsInvoker == nil {
		freeXMLRPCParamsInvoker = gi.StructFunctionInvokerNew("Soup", "XMLRPCParams", "free")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	freeXMLRPCParamsInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'soup_xmlrpc_params_parse' : return type 'GLib.Variant' not supported
