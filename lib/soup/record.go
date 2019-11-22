// Code generated - DO NOT EDIT.

package soup

import (
	gi "github.com/pekim/gobbi/internal/gi"
	"sync"
)

var addressClassStruct *gi.Struct
var addressClassStruct_Once sync.Once

func addressClassStruct_Set() error {
	var err error
	addressClassStruct_Once.Do(func() {
		addressClassStruct, err = gi.StructNew("Soup", "AddressClass")
	})
	return err
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

func authClassStruct_Set() error {
	var err error
	authClassStruct_Once.Do(func() {
		authClassStruct, err = gi.StructNew("Soup", "AuthClass")
	})
	return err
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

func authDomainBasicClassStruct_Set() error {
	var err error
	authDomainBasicClassStruct_Once.Do(func() {
		authDomainBasicClassStruct, err = gi.StructNew("Soup", "AuthDomainBasicClass")
	})
	return err
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

func authDomainClassStruct_Set() error {
	var err error
	authDomainClassStruct_Once.Do(func() {
		authDomainClassStruct, err = gi.StructNew("Soup", "AuthDomainClass")
	})
	return err
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

func authDomainDigestClassStruct_Set() error {
	var err error
	authDomainDigestClassStruct_Once.Do(func() {
		authDomainDigestClassStruct, err = gi.StructNew("Soup", "AuthDomainDigestClass")
	})
	return err
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

func authManagerClassStruct_Set() error {
	var err error
	authManagerClassStruct_Once.Do(func() {
		authManagerClassStruct, err = gi.StructNew("Soup", "AuthManagerClass")
	})
	return err
}

type AuthManagerClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value 'authenticate' : missing Type
}

var authManagerPrivateStruct *gi.Struct
var authManagerPrivateStruct_Once sync.Once

func authManagerPrivateStruct_Set() error {
	var err error
	authManagerPrivateStruct_Once.Do(func() {
		authManagerPrivateStruct, err = gi.StructNew("Soup", "AuthManagerPrivate")
	})
	return err
}

type AuthManagerPrivate struct {
	native uintptr
}

var bufferStruct *gi.Struct
var bufferStruct_Once sync.Once

func bufferStruct_Set() error {
	var err error
	bufferStruct_Once.Do(func() {
		bufferStruct, err = gi.StructNew("Soup", "Buffer")
	})
	return err
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

func bufferCopyFunction_Set() error {
	var err error
	bufferCopyFunction_Once.Do(func() {
		err = bufferStruct_Set()
		if err != nil {
			return
		}
		bufferCopyFunction, err = bufferStruct.InvokerNew("copy")
	})
	return err
}

// Copy is a representation of the C type soup_buffer_copy.
func (recv *Buffer) Copy() *Buffer {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := bufferCopyFunction_Set()
	if err == nil {
		ret = bufferCopyFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Buffer{native: ret.Pointer()}

	return retGo
}

var bufferFreeFunction *gi.Function
var bufferFreeFunction_Once sync.Once

func bufferFreeFunction_Set() error {
	var err error
	bufferFreeFunction_Once.Do(func() {
		err = bufferStruct_Set()
		if err != nil {
			return
		}
		bufferFreeFunction, err = bufferStruct.InvokerNew("free")
	})
	return err
}

// Free is a representation of the C type soup_buffer_free.
func (recv *Buffer) Free() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := bufferFreeFunction_Set()
	if err == nil {
		bufferFreeFunction.Invoke(inArgs[:], nil)
	}

}

// UNSUPPORTED : C value 'soup_buffer_get_as_bytes' : return type 'GLib.Bytes' not supported

// UNSUPPORTED : C value 'soup_buffer_get_data' : parameter 'data' has no type

// UNSUPPORTED : C value 'soup_buffer_get_owner' : return type 'gpointer' not supported

// UNSUPPORTED : C value 'soup_buffer_new_subbuffer' : parameter 'offset' of type 'gsize' not supported

var cacheClassStruct *gi.Struct
var cacheClassStruct_Once sync.Once

func cacheClassStruct_Set() error {
	var err error
	cacheClassStruct_Once.Do(func() {
		cacheClassStruct, err = gi.StructNew("Soup", "CacheClass")
	})
	return err
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

func cachePrivateStruct_Set() error {
	var err error
	cachePrivateStruct_Once.Do(func() {
		cachePrivateStruct, err = gi.StructNew("Soup", "CachePrivate")
	})
	return err
}

type CachePrivate struct {
	native uintptr
}

var clientContextStruct *gi.Struct
var clientContextStruct_Once sync.Once

func clientContextStruct_Set() error {
	var err error
	clientContextStruct_Once.Do(func() {
		clientContextStruct, err = gi.StructNew("Soup", "ClientContext")
	})
	return err
}

type ClientContext struct {
	native uintptr
}

// UNSUPPORTED : C value 'soup_client_context_get_address' : return type 'Address' not supported

// UNSUPPORTED : C value 'soup_client_context_get_auth_domain' : return type 'AuthDomain' not supported

var clientContextGetAuthUserFunction *gi.Function
var clientContextGetAuthUserFunction_Once sync.Once

func clientContextGetAuthUserFunction_Set() error {
	var err error
	clientContextGetAuthUserFunction_Once.Do(func() {
		err = clientContextStruct_Set()
		if err != nil {
			return
		}
		clientContextGetAuthUserFunction, err = clientContextStruct.InvokerNew("get_auth_user")
	})
	return err
}

// GetAuthUser is a representation of the C type soup_client_context_get_auth_user.
func (recv *ClientContext) GetAuthUser() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := clientContextGetAuthUserFunction_Set()
	if err == nil {
		ret = clientContextGetAuthUserFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'soup_client_context_get_gsocket' : return type 'Gio.Socket' not supported

var clientContextGetHostFunction *gi.Function
var clientContextGetHostFunction_Once sync.Once

func clientContextGetHostFunction_Set() error {
	var err error
	clientContextGetHostFunction_Once.Do(func() {
		err = clientContextStruct_Set()
		if err != nil {
			return
		}
		clientContextGetHostFunction, err = clientContextStruct.InvokerNew("get_host")
	})
	return err
}

// GetHost is a representation of the C type soup_client_context_get_host.
func (recv *ClientContext) GetHost() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := clientContextGetHostFunction_Set()
	if err == nil {
		ret = clientContextGetHostFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'soup_client_context_get_local_address' : return type 'Gio.SocketAddress' not supported

// UNSUPPORTED : C value 'soup_client_context_get_remote_address' : return type 'Gio.SocketAddress' not supported

// UNSUPPORTED : C value 'soup_client_context_get_socket' : return type 'Socket' not supported

// UNSUPPORTED : C value 'soup_client_context_steal_connection' : return type 'Gio.IOStream' not supported

var connectionStruct *gi.Struct
var connectionStruct_Once sync.Once

func connectionStruct_Set() error {
	var err error
	connectionStruct_Once.Do(func() {
		connectionStruct, err = gi.StructNew("Soup", "Connection")
	})
	return err
}

type Connection struct {
	native uintptr
}

var contentDecoderClassStruct *gi.Struct
var contentDecoderClassStruct_Once sync.Once

func contentDecoderClassStruct_Set() error {
	var err error
	contentDecoderClassStruct_Once.Do(func() {
		contentDecoderClassStruct, err = gi.StructNew("Soup", "ContentDecoderClass")
	})
	return err
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

func contentDecoderPrivateStruct_Set() error {
	var err error
	contentDecoderPrivateStruct_Once.Do(func() {
		contentDecoderPrivateStruct, err = gi.StructNew("Soup", "ContentDecoderPrivate")
	})
	return err
}

type ContentDecoderPrivate struct {
	native uintptr
}

var contentSnifferClassStruct *gi.Struct
var contentSnifferClassStruct_Once sync.Once

func contentSnifferClassStruct_Set() error {
	var err error
	contentSnifferClassStruct_Once.Do(func() {
		contentSnifferClassStruct, err = gi.StructNew("Soup", "ContentSnifferClass")
	})
	return err
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

func contentSnifferPrivateStruct_Set() error {
	var err error
	contentSnifferPrivateStruct_Once.Do(func() {
		contentSnifferPrivateStruct, err = gi.StructNew("Soup", "ContentSnifferPrivate")
	})
	return err
}

type ContentSnifferPrivate struct {
	native uintptr
}

var cookieStruct *gi.Struct
var cookieStruct_Once sync.Once

func cookieStruct_Set() error {
	var err error
	cookieStruct_Once.Do(func() {
		cookieStruct, err = gi.StructNew("Soup", "Cookie")
	})
	return err
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

func cookieNewFunction_Set() error {
	var err error
	cookieNewFunction_Once.Do(func() {
		err = cookieStruct_Set()
		if err != nil {
			return
		}
		cookieNewFunction, err = cookieStruct.InvokerNew("new")
	})
	return err
}

// CookieNew is a representation of the C type soup_cookie_new.
func CookieNew(name string, value string, domain string, path string, maxAge int32) *Cookie {
	var inArgs [5]gi.Argument
	inArgs[0].SetString(name)
	inArgs[1].SetString(value)
	inArgs[2].SetString(domain)
	inArgs[3].SetString(path)
	inArgs[4].SetInt32(maxAge)

	var ret gi.Argument

	err := cookieNewFunction_Set()
	if err == nil {
		ret = cookieNewFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Cookie{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'soup_cookie_applies_to_uri' : parameter 'uri' of type 'URI' not supported

var cookieCopyFunction *gi.Function
var cookieCopyFunction_Once sync.Once

func cookieCopyFunction_Set() error {
	var err error
	cookieCopyFunction_Once.Do(func() {
		err = cookieStruct_Set()
		if err != nil {
			return
		}
		cookieCopyFunction, err = cookieStruct.InvokerNew("copy")
	})
	return err
}

// Copy is a representation of the C type soup_cookie_copy.
func (recv *Cookie) Copy() *Cookie {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := cookieCopyFunction_Set()
	if err == nil {
		ret = cookieCopyFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Cookie{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'soup_cookie_domain_matches' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'soup_cookie_equal' : parameter 'cookie2' of type 'Cookie' not supported

var cookieFreeFunction *gi.Function
var cookieFreeFunction_Once sync.Once

func cookieFreeFunction_Set() error {
	var err error
	cookieFreeFunction_Once.Do(func() {
		err = cookieStruct_Set()
		if err != nil {
			return
		}
		cookieFreeFunction, err = cookieStruct.InvokerNew("free")
	})
	return err
}

// Free is a representation of the C type soup_cookie_free.
func (recv *Cookie) Free() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := cookieFreeFunction_Set()
	if err == nil {
		cookieFreeFunction.Invoke(inArgs[:], nil)
	}

}

var cookieGetDomainFunction *gi.Function
var cookieGetDomainFunction_Once sync.Once

func cookieGetDomainFunction_Set() error {
	var err error
	cookieGetDomainFunction_Once.Do(func() {
		err = cookieStruct_Set()
		if err != nil {
			return
		}
		cookieGetDomainFunction, err = cookieStruct.InvokerNew("get_domain")
	})
	return err
}

// GetDomain is a representation of the C type soup_cookie_get_domain.
func (recv *Cookie) GetDomain() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := cookieGetDomainFunction_Set()
	if err == nil {
		ret = cookieGetDomainFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var cookieGetExpiresFunction *gi.Function
var cookieGetExpiresFunction_Once sync.Once

func cookieGetExpiresFunction_Set() error {
	var err error
	cookieGetExpiresFunction_Once.Do(func() {
		err = cookieStruct_Set()
		if err != nil {
			return
		}
		cookieGetExpiresFunction, err = cookieStruct.InvokerNew("get_expires")
	})
	return err
}

// GetExpires is a representation of the C type soup_cookie_get_expires.
func (recv *Cookie) GetExpires() *Date {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := cookieGetExpiresFunction_Set()
	if err == nil {
		ret = cookieGetExpiresFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Date{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'soup_cookie_get_http_only' : return type 'gboolean' not supported

var cookieGetNameFunction *gi.Function
var cookieGetNameFunction_Once sync.Once

func cookieGetNameFunction_Set() error {
	var err error
	cookieGetNameFunction_Once.Do(func() {
		err = cookieStruct_Set()
		if err != nil {
			return
		}
		cookieGetNameFunction, err = cookieStruct.InvokerNew("get_name")
	})
	return err
}

// GetName is a representation of the C type soup_cookie_get_name.
func (recv *Cookie) GetName() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := cookieGetNameFunction_Set()
	if err == nil {
		ret = cookieGetNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var cookieGetPathFunction *gi.Function
var cookieGetPathFunction_Once sync.Once

func cookieGetPathFunction_Set() error {
	var err error
	cookieGetPathFunction_Once.Do(func() {
		err = cookieStruct_Set()
		if err != nil {
			return
		}
		cookieGetPathFunction, err = cookieStruct.InvokerNew("get_path")
	})
	return err
}

// GetPath is a representation of the C type soup_cookie_get_path.
func (recv *Cookie) GetPath() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := cookieGetPathFunction_Set()
	if err == nil {
		ret = cookieGetPathFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'soup_cookie_get_secure' : return type 'gboolean' not supported

var cookieGetValueFunction *gi.Function
var cookieGetValueFunction_Once sync.Once

func cookieGetValueFunction_Set() error {
	var err error
	cookieGetValueFunction_Once.Do(func() {
		err = cookieStruct_Set()
		if err != nil {
			return
		}
		cookieGetValueFunction, err = cookieStruct.InvokerNew("get_value")
	})
	return err
}

// GetValue is a representation of the C type soup_cookie_get_value.
func (recv *Cookie) GetValue() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := cookieGetValueFunction_Set()
	if err == nil {
		ret = cookieGetValueFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var cookieSetDomainFunction *gi.Function
var cookieSetDomainFunction_Once sync.Once

func cookieSetDomainFunction_Set() error {
	var err error
	cookieSetDomainFunction_Once.Do(func() {
		err = cookieStruct_Set()
		if err != nil {
			return
		}
		cookieSetDomainFunction, err = cookieStruct.InvokerNew("set_domain")
	})
	return err
}

// SetDomain is a representation of the C type soup_cookie_set_domain.
func (recv *Cookie) SetDomain(domain string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(domain)

	err := cookieSetDomainFunction_Set()
	if err == nil {
		cookieSetDomainFunction.Invoke(inArgs[:], nil)
	}

}

// UNSUPPORTED : C value 'soup_cookie_set_expires' : parameter 'expires' of type 'Date' not supported

// UNSUPPORTED : C value 'soup_cookie_set_http_only' : parameter 'http_only' of type 'gboolean' not supported

var cookieSetMaxAgeFunction *gi.Function
var cookieSetMaxAgeFunction_Once sync.Once

func cookieSetMaxAgeFunction_Set() error {
	var err error
	cookieSetMaxAgeFunction_Once.Do(func() {
		err = cookieStruct_Set()
		if err != nil {
			return
		}
		cookieSetMaxAgeFunction, err = cookieStruct.InvokerNew("set_max_age")
	})
	return err
}

// SetMaxAge is a representation of the C type soup_cookie_set_max_age.
func (recv *Cookie) SetMaxAge(maxAge int32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(maxAge)

	err := cookieSetMaxAgeFunction_Set()
	if err == nil {
		cookieSetMaxAgeFunction.Invoke(inArgs[:], nil)
	}

}

var cookieSetNameFunction *gi.Function
var cookieSetNameFunction_Once sync.Once

func cookieSetNameFunction_Set() error {
	var err error
	cookieSetNameFunction_Once.Do(func() {
		err = cookieStruct_Set()
		if err != nil {
			return
		}
		cookieSetNameFunction, err = cookieStruct.InvokerNew("set_name")
	})
	return err
}

// SetName is a representation of the C type soup_cookie_set_name.
func (recv *Cookie) SetName(name string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(name)

	err := cookieSetNameFunction_Set()
	if err == nil {
		cookieSetNameFunction.Invoke(inArgs[:], nil)
	}

}

var cookieSetPathFunction *gi.Function
var cookieSetPathFunction_Once sync.Once

func cookieSetPathFunction_Set() error {
	var err error
	cookieSetPathFunction_Once.Do(func() {
		err = cookieStruct_Set()
		if err != nil {
			return
		}
		cookieSetPathFunction, err = cookieStruct.InvokerNew("set_path")
	})
	return err
}

// SetPath is a representation of the C type soup_cookie_set_path.
func (recv *Cookie) SetPath(path string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(path)

	err := cookieSetPathFunction_Set()
	if err == nil {
		cookieSetPathFunction.Invoke(inArgs[:], nil)
	}

}

// UNSUPPORTED : C value 'soup_cookie_set_secure' : parameter 'secure' of type 'gboolean' not supported

var cookieSetValueFunction *gi.Function
var cookieSetValueFunction_Once sync.Once

func cookieSetValueFunction_Set() error {
	var err error
	cookieSetValueFunction_Once.Do(func() {
		err = cookieStruct_Set()
		if err != nil {
			return
		}
		cookieSetValueFunction, err = cookieStruct.InvokerNew("set_value")
	})
	return err
}

// SetValue is a representation of the C type soup_cookie_set_value.
func (recv *Cookie) SetValue(value string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(value)

	err := cookieSetValueFunction_Set()
	if err == nil {
		cookieSetValueFunction.Invoke(inArgs[:], nil)
	}

}

var cookieToCookieHeaderFunction *gi.Function
var cookieToCookieHeaderFunction_Once sync.Once

func cookieToCookieHeaderFunction_Set() error {
	var err error
	cookieToCookieHeaderFunction_Once.Do(func() {
		err = cookieStruct_Set()
		if err != nil {
			return
		}
		cookieToCookieHeaderFunction, err = cookieStruct.InvokerNew("to_cookie_header")
	})
	return err
}

// ToCookieHeader is a representation of the C type soup_cookie_to_cookie_header.
func (recv *Cookie) ToCookieHeader() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := cookieToCookieHeaderFunction_Set()
	if err == nil {
		ret = cookieToCookieHeaderFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

var cookieToSetCookieHeaderFunction *gi.Function
var cookieToSetCookieHeaderFunction_Once sync.Once

func cookieToSetCookieHeaderFunction_Set() error {
	var err error
	cookieToSetCookieHeaderFunction_Once.Do(func() {
		err = cookieStruct_Set()
		if err != nil {
			return
		}
		cookieToSetCookieHeaderFunction, err = cookieStruct.InvokerNew("to_set_cookie_header")
	})
	return err
}

// ToSetCookieHeader is a representation of the C type soup_cookie_to_set_cookie_header.
func (recv *Cookie) ToSetCookieHeader() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := cookieToSetCookieHeaderFunction_Set()
	if err == nil {
		ret = cookieToSetCookieHeaderFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

var cookieJarClassStruct *gi.Struct
var cookieJarClassStruct_Once sync.Once

func cookieJarClassStruct_Set() error {
	var err error
	cookieJarClassStruct_Once.Do(func() {
		cookieJarClassStruct, err = gi.StructNew("Soup", "CookieJarClass")
	})
	return err
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

func cookieJarDBClassStruct_Set() error {
	var err error
	cookieJarDBClassStruct_Once.Do(func() {
		cookieJarDBClassStruct, err = gi.StructNew("Soup", "CookieJarDBClass")
	})
	return err
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

func cookieJarTextClassStruct_Set() error {
	var err error
	cookieJarTextClassStruct_Once.Do(func() {
		cookieJarTextClassStruct, err = gi.StructNew("Soup", "CookieJarTextClass")
	})
	return err
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

func dateStruct_Set() error {
	var err error
	dateStruct_Once.Do(func() {
		dateStruct, err = gi.StructNew("Soup", "Date")
	})
	return err
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

func dateNewFunction_Set() error {
	var err error
	dateNewFunction_Once.Do(func() {
		err = dateStruct_Set()
		if err != nil {
			return
		}
		dateNewFunction, err = dateStruct.InvokerNew("new")
	})
	return err
}

// DateNew is a representation of the C type soup_date_new.
func DateNew(year int32, month int32, day int32, hour int32, minute int32, second int32) *Date {
	var inArgs [6]gi.Argument
	inArgs[0].SetInt32(year)
	inArgs[1].SetInt32(month)
	inArgs[2].SetInt32(day)
	inArgs[3].SetInt32(hour)
	inArgs[4].SetInt32(minute)
	inArgs[5].SetInt32(second)

	var ret gi.Argument

	err := dateNewFunction_Set()
	if err == nil {
		ret = dateNewFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Date{native: ret.Pointer()}

	return retGo
}

var dateNewFromNowFunction *gi.Function
var dateNewFromNowFunction_Once sync.Once

func dateNewFromNowFunction_Set() error {
	var err error
	dateNewFromNowFunction_Once.Do(func() {
		err = dateStruct_Set()
		if err != nil {
			return
		}
		dateNewFromNowFunction, err = dateStruct.InvokerNew("new_from_now")
	})
	return err
}

// DateNewFromNow is a representation of the C type soup_date_new_from_now.
func DateNewFromNow(offsetSeconds int32) *Date {
	var inArgs [1]gi.Argument
	inArgs[0].SetInt32(offsetSeconds)

	var ret gi.Argument

	err := dateNewFromNowFunction_Set()
	if err == nil {
		ret = dateNewFromNowFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Date{native: ret.Pointer()}

	return retGo
}

var dateNewFromStringFunction *gi.Function
var dateNewFromStringFunction_Once sync.Once

func dateNewFromStringFunction_Set() error {
	var err error
	dateNewFromStringFunction_Once.Do(func() {
		err = dateStruct_Set()
		if err != nil {
			return
		}
		dateNewFromStringFunction, err = dateStruct.InvokerNew("new_from_string")
	})
	return err
}

// DateNewFromString is a representation of the C type soup_date_new_from_string.
func DateNewFromString(dateString string) *Date {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(dateString)

	var ret gi.Argument

	err := dateNewFromStringFunction_Set()
	if err == nil {
		ret = dateNewFromStringFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Date{native: ret.Pointer()}

	return retGo
}

var dateNewFromTimeTFunction *gi.Function
var dateNewFromTimeTFunction_Once sync.Once

func dateNewFromTimeTFunction_Set() error {
	var err error
	dateNewFromTimeTFunction_Once.Do(func() {
		err = dateStruct_Set()
		if err != nil {
			return
		}
		dateNewFromTimeTFunction, err = dateStruct.InvokerNew("new_from_time_t")
	})
	return err
}

// DateNewFromTimeT is a representation of the C type soup_date_new_from_time_t.
func DateNewFromTimeT(when int64) *Date {
	var inArgs [1]gi.Argument
	inArgs[0].SetInt64(when)

	var ret gi.Argument

	err := dateNewFromTimeTFunction_Set()
	if err == nil {
		ret = dateNewFromTimeTFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Date{native: ret.Pointer()}

	return retGo
}

var dateCopyFunction *gi.Function
var dateCopyFunction_Once sync.Once

func dateCopyFunction_Set() error {
	var err error
	dateCopyFunction_Once.Do(func() {
		err = dateStruct_Set()
		if err != nil {
			return
		}
		dateCopyFunction, err = dateStruct.InvokerNew("copy")
	})
	return err
}

// Copy is a representation of the C type soup_date_copy.
func (recv *Date) Copy() *Date {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := dateCopyFunction_Set()
	if err == nil {
		ret = dateCopyFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Date{native: ret.Pointer()}

	return retGo
}

var dateFreeFunction *gi.Function
var dateFreeFunction_Once sync.Once

func dateFreeFunction_Set() error {
	var err error
	dateFreeFunction_Once.Do(func() {
		err = dateStruct_Set()
		if err != nil {
			return
		}
		dateFreeFunction, err = dateStruct.InvokerNew("free")
	})
	return err
}

// Free is a representation of the C type soup_date_free.
func (recv *Date) Free() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := dateFreeFunction_Set()
	if err == nil {
		dateFreeFunction.Invoke(inArgs[:], nil)
	}

}

var dateGetDayFunction *gi.Function
var dateGetDayFunction_Once sync.Once

func dateGetDayFunction_Set() error {
	var err error
	dateGetDayFunction_Once.Do(func() {
		err = dateStruct_Set()
		if err != nil {
			return
		}
		dateGetDayFunction, err = dateStruct.InvokerNew("get_day")
	})
	return err
}

// GetDay is a representation of the C type soup_date_get_day.
func (recv *Date) GetDay() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := dateGetDayFunction_Set()
	if err == nil {
		ret = dateGetDayFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var dateGetHourFunction *gi.Function
var dateGetHourFunction_Once sync.Once

func dateGetHourFunction_Set() error {
	var err error
	dateGetHourFunction_Once.Do(func() {
		err = dateStruct_Set()
		if err != nil {
			return
		}
		dateGetHourFunction, err = dateStruct.InvokerNew("get_hour")
	})
	return err
}

// GetHour is a representation of the C type soup_date_get_hour.
func (recv *Date) GetHour() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := dateGetHourFunction_Set()
	if err == nil {
		ret = dateGetHourFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var dateGetMinuteFunction *gi.Function
var dateGetMinuteFunction_Once sync.Once

func dateGetMinuteFunction_Set() error {
	var err error
	dateGetMinuteFunction_Once.Do(func() {
		err = dateStruct_Set()
		if err != nil {
			return
		}
		dateGetMinuteFunction, err = dateStruct.InvokerNew("get_minute")
	})
	return err
}

// GetMinute is a representation of the C type soup_date_get_minute.
func (recv *Date) GetMinute() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := dateGetMinuteFunction_Set()
	if err == nil {
		ret = dateGetMinuteFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var dateGetMonthFunction *gi.Function
var dateGetMonthFunction_Once sync.Once

func dateGetMonthFunction_Set() error {
	var err error
	dateGetMonthFunction_Once.Do(func() {
		err = dateStruct_Set()
		if err != nil {
			return
		}
		dateGetMonthFunction, err = dateStruct.InvokerNew("get_month")
	})
	return err
}

// GetMonth is a representation of the C type soup_date_get_month.
func (recv *Date) GetMonth() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := dateGetMonthFunction_Set()
	if err == nil {
		ret = dateGetMonthFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var dateGetOffsetFunction *gi.Function
var dateGetOffsetFunction_Once sync.Once

func dateGetOffsetFunction_Set() error {
	var err error
	dateGetOffsetFunction_Once.Do(func() {
		err = dateStruct_Set()
		if err != nil {
			return
		}
		dateGetOffsetFunction, err = dateStruct.InvokerNew("get_offset")
	})
	return err
}

// GetOffset is a representation of the C type soup_date_get_offset.
func (recv *Date) GetOffset() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := dateGetOffsetFunction_Set()
	if err == nil {
		ret = dateGetOffsetFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var dateGetSecondFunction *gi.Function
var dateGetSecondFunction_Once sync.Once

func dateGetSecondFunction_Set() error {
	var err error
	dateGetSecondFunction_Once.Do(func() {
		err = dateStruct_Set()
		if err != nil {
			return
		}
		dateGetSecondFunction, err = dateStruct.InvokerNew("get_second")
	})
	return err
}

// GetSecond is a representation of the C type soup_date_get_second.
func (recv *Date) GetSecond() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := dateGetSecondFunction_Set()
	if err == nil {
		ret = dateGetSecondFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var dateGetUtcFunction *gi.Function
var dateGetUtcFunction_Once sync.Once

func dateGetUtcFunction_Set() error {
	var err error
	dateGetUtcFunction_Once.Do(func() {
		err = dateStruct_Set()
		if err != nil {
			return
		}
		dateGetUtcFunction, err = dateStruct.InvokerNew("get_utc")
	})
	return err
}

// GetUtc is a representation of the C type soup_date_get_utc.
func (recv *Date) GetUtc() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := dateGetUtcFunction_Set()
	if err == nil {
		ret = dateGetUtcFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var dateGetYearFunction *gi.Function
var dateGetYearFunction_Once sync.Once

func dateGetYearFunction_Set() error {
	var err error
	dateGetYearFunction_Once.Do(func() {
		err = dateStruct_Set()
		if err != nil {
			return
		}
		dateGetYearFunction, err = dateStruct.InvokerNew("get_year")
	})
	return err
}

// GetYear is a representation of the C type soup_date_get_year.
func (recv *Date) GetYear() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := dateGetYearFunction_Set()
	if err == nil {
		ret = dateGetYearFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

// UNSUPPORTED : C value 'soup_date_is_past' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'soup_date_to_string' : parameter 'format' of type 'DateFormat' not supported

var dateToTimeTFunction *gi.Function
var dateToTimeTFunction_Once sync.Once

func dateToTimeTFunction_Set() error {
	var err error
	dateToTimeTFunction_Once.Do(func() {
		err = dateStruct_Set()
		if err != nil {
			return
		}
		dateToTimeTFunction, err = dateStruct.InvokerNew("to_time_t")
	})
	return err
}

// ToTimeT is a representation of the C type soup_date_to_time_t.
func (recv *Date) ToTimeT() int64 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := dateToTimeTFunction_Set()
	if err == nil {
		ret = dateToTimeTFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int64()

	return retGo
}

// UNSUPPORTED : C value 'soup_date_to_timeval' : parameter 'time' of type 'GLib.TimeVal' not supported

var hSTSEnforcerClassStruct *gi.Struct
var hSTSEnforcerClassStruct_Once sync.Once

func hSTSEnforcerClassStruct_Set() error {
	var err error
	hSTSEnforcerClassStruct_Once.Do(func() {
		hSTSEnforcerClassStruct, err = gi.StructNew("Soup", "HSTSEnforcerClass")
	})
	return err
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

func hSTSEnforcerDBClassStruct_Set() error {
	var err error
	hSTSEnforcerDBClassStruct_Once.Do(func() {
		hSTSEnforcerDBClassStruct, err = gi.StructNew("Soup", "HSTSEnforcerDBClass")
	})
	return err
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

func hSTSEnforcerDBPrivateStruct_Set() error {
	var err error
	hSTSEnforcerDBPrivateStruct_Once.Do(func() {
		hSTSEnforcerDBPrivateStruct, err = gi.StructNew("Soup", "HSTSEnforcerDBPrivate")
	})
	return err
}

type HSTSEnforcerDBPrivate struct {
	native uintptr
}

var hSTSEnforcerPrivateStruct *gi.Struct
var hSTSEnforcerPrivateStruct_Once sync.Once

func hSTSEnforcerPrivateStruct_Set() error {
	var err error
	hSTSEnforcerPrivateStruct_Once.Do(func() {
		hSTSEnforcerPrivateStruct, err = gi.StructNew("Soup", "HSTSEnforcerPrivate")
	})
	return err
}

type HSTSEnforcerPrivate struct {
	native uintptr
}

var hSTSPolicyStruct *gi.Struct
var hSTSPolicyStruct_Once sync.Once

func hSTSPolicyStruct_Set() error {
	var err error
	hSTSPolicyStruct_Once.Do(func() {
		hSTSPolicyStruct, err = gi.StructNew("Soup", "HSTSPolicy")
	})
	return err
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

func hSTSPolicyCopyFunction_Set() error {
	var err error
	hSTSPolicyCopyFunction_Once.Do(func() {
		err = hSTSPolicyStruct_Set()
		if err != nil {
			return
		}
		hSTSPolicyCopyFunction, err = hSTSPolicyStruct.InvokerNew("copy")
	})
	return err
}

// Copy is a representation of the C type soup_hsts_policy_copy.
func (recv *HSTSPolicy) Copy() *HSTSPolicy {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := hSTSPolicyCopyFunction_Set()
	if err == nil {
		ret = hSTSPolicyCopyFunction.Invoke(inArgs[:], nil)
	}

	retGo := &HSTSPolicy{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'soup_hsts_policy_equal' : parameter 'policy2' of type 'HSTSPolicy' not supported

var hSTSPolicyFreeFunction *gi.Function
var hSTSPolicyFreeFunction_Once sync.Once

func hSTSPolicyFreeFunction_Set() error {
	var err error
	hSTSPolicyFreeFunction_Once.Do(func() {
		err = hSTSPolicyStruct_Set()
		if err != nil {
			return
		}
		hSTSPolicyFreeFunction, err = hSTSPolicyStruct.InvokerNew("free")
	})
	return err
}

// Free is a representation of the C type soup_hsts_policy_free.
func (recv *HSTSPolicy) Free() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := hSTSPolicyFreeFunction_Set()
	if err == nil {
		hSTSPolicyFreeFunction.Invoke(inArgs[:], nil)
	}

}

var hSTSPolicyGetDomainFunction *gi.Function
var hSTSPolicyGetDomainFunction_Once sync.Once

func hSTSPolicyGetDomainFunction_Set() error {
	var err error
	hSTSPolicyGetDomainFunction_Once.Do(func() {
		err = hSTSPolicyStruct_Set()
		if err != nil {
			return
		}
		hSTSPolicyGetDomainFunction, err = hSTSPolicyStruct.InvokerNew("get_domain")
	})
	return err
}

// GetDomain is a representation of the C type soup_hsts_policy_get_domain.
func (recv *HSTSPolicy) GetDomain() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := hSTSPolicyGetDomainFunction_Set()
	if err == nil {
		ret = hSTSPolicyGetDomainFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'soup_hsts_policy_includes_subdomains' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'soup_hsts_policy_is_expired' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'soup_hsts_policy_is_session_policy' : return type 'gboolean' not supported

var loggerClassStruct *gi.Struct
var loggerClassStruct_Once sync.Once

func loggerClassStruct_Set() error {
	var err error
	loggerClassStruct_Once.Do(func() {
		loggerClassStruct, err = gi.StructNew("Soup", "LoggerClass")
	})
	return err
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

func messageBodyStruct_Set() error {
	var err error
	messageBodyStruct_Once.Do(func() {
		messageBodyStruct, err = gi.StructNew("Soup", "MessageBody")
	})
	return err
}

type MessageBody struct {
	native uintptr
	Data   string
	Length int64
}

var messageBodyNewFunction *gi.Function
var messageBodyNewFunction_Once sync.Once

func messageBodyNewFunction_Set() error {
	var err error
	messageBodyNewFunction_Once.Do(func() {
		err = messageBodyStruct_Set()
		if err != nil {
			return
		}
		messageBodyNewFunction, err = messageBodyStruct.InvokerNew("new")
	})
	return err
}

// MessageBodyNew is a representation of the C type soup_message_body_new.
func MessageBodyNew() *MessageBody {

	var ret gi.Argument

	err := messageBodyNewFunction_Set()
	if err == nil {
		ret = messageBodyNewFunction.Invoke(nil, nil)
	}

	retGo := &MessageBody{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'soup_message_body_append' : parameter 'use' of type 'MemoryUse' not supported

// UNSUPPORTED : C value 'soup_message_body_append_buffer' : parameter 'buffer' of type 'Buffer' not supported

// UNSUPPORTED : C value 'soup_message_body_append_take' : parameter 'data' has no type

var messageBodyCompleteFunction *gi.Function
var messageBodyCompleteFunction_Once sync.Once

func messageBodyCompleteFunction_Set() error {
	var err error
	messageBodyCompleteFunction_Once.Do(func() {
		err = messageBodyStruct_Set()
		if err != nil {
			return
		}
		messageBodyCompleteFunction, err = messageBodyStruct.InvokerNew("complete")
	})
	return err
}

// Complete is a representation of the C type soup_message_body_complete.
func (recv *MessageBody) Complete() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := messageBodyCompleteFunction_Set()
	if err == nil {
		messageBodyCompleteFunction.Invoke(inArgs[:], nil)
	}

}

var messageBodyFlattenFunction *gi.Function
var messageBodyFlattenFunction_Once sync.Once

func messageBodyFlattenFunction_Set() error {
	var err error
	messageBodyFlattenFunction_Once.Do(func() {
		err = messageBodyStruct_Set()
		if err != nil {
			return
		}
		messageBodyFlattenFunction, err = messageBodyStruct.InvokerNew("flatten")
	})
	return err
}

// Flatten is a representation of the C type soup_message_body_flatten.
func (recv *MessageBody) Flatten() *Buffer {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := messageBodyFlattenFunction_Set()
	if err == nil {
		ret = messageBodyFlattenFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Buffer{native: ret.Pointer()}

	return retGo
}

var messageBodyFreeFunction *gi.Function
var messageBodyFreeFunction_Once sync.Once

func messageBodyFreeFunction_Set() error {
	var err error
	messageBodyFreeFunction_Once.Do(func() {
		err = messageBodyStruct_Set()
		if err != nil {
			return
		}
		messageBodyFreeFunction, err = messageBodyStruct.InvokerNew("free")
	})
	return err
}

// Free is a representation of the C type soup_message_body_free.
func (recv *MessageBody) Free() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := messageBodyFreeFunction_Set()
	if err == nil {
		messageBodyFreeFunction.Invoke(inArgs[:], nil)
	}

}

// UNSUPPORTED : C value 'soup_message_body_get_accumulate' : return type 'gboolean' not supported

var messageBodyGetChunkFunction *gi.Function
var messageBodyGetChunkFunction_Once sync.Once

func messageBodyGetChunkFunction_Set() error {
	var err error
	messageBodyGetChunkFunction_Once.Do(func() {
		err = messageBodyStruct_Set()
		if err != nil {
			return
		}
		messageBodyGetChunkFunction, err = messageBodyStruct.InvokerNew("get_chunk")
	})
	return err
}

// GetChunk is a representation of the C type soup_message_body_get_chunk.
func (recv *MessageBody) GetChunk(offset int64) *Buffer {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt64(offset)

	var ret gi.Argument

	err := messageBodyGetChunkFunction_Set()
	if err == nil {
		ret = messageBodyGetChunkFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Buffer{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'soup_message_body_got_chunk' : parameter 'chunk' of type 'Buffer' not supported

// UNSUPPORTED : C value 'soup_message_body_set_accumulate' : parameter 'accumulate' of type 'gboolean' not supported

var messageBodyTruncateFunction *gi.Function
var messageBodyTruncateFunction_Once sync.Once

func messageBodyTruncateFunction_Set() error {
	var err error
	messageBodyTruncateFunction_Once.Do(func() {
		err = messageBodyStruct_Set()
		if err != nil {
			return
		}
		messageBodyTruncateFunction, err = messageBodyStruct.InvokerNew("truncate")
	})
	return err
}

// Truncate is a representation of the C type soup_message_body_truncate.
func (recv *MessageBody) Truncate() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := messageBodyTruncateFunction_Set()
	if err == nil {
		messageBodyTruncateFunction.Invoke(inArgs[:], nil)
	}

}

// UNSUPPORTED : C value 'soup_message_body_wrote_chunk' : parameter 'chunk' of type 'Buffer' not supported

var messageClassStruct *gi.Struct
var messageClassStruct_Once sync.Once

func messageClassStruct_Set() error {
	var err error
	messageClassStruct_Once.Do(func() {
		messageClassStruct, err = gi.StructNew("Soup", "MessageClass")
	})
	return err
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

func messageHeadersStruct_Set() error {
	var err error
	messageHeadersStruct_Once.Do(func() {
		messageHeadersStruct, err = gi.StructNew("Soup", "MessageHeaders")
	})
	return err
}

type MessageHeaders struct {
	native uintptr
}

// UNSUPPORTED : C value 'soup_message_headers_new' : parameter 'type' of type 'MessageHeadersType' not supported

var messageHeadersAppendFunction *gi.Function
var messageHeadersAppendFunction_Once sync.Once

func messageHeadersAppendFunction_Set() error {
	var err error
	messageHeadersAppendFunction_Once.Do(func() {
		err = messageHeadersStruct_Set()
		if err != nil {
			return
		}
		messageHeadersAppendFunction, err = messageHeadersStruct.InvokerNew("append")
	})
	return err
}

// Append is a representation of the C type soup_message_headers_append.
func (recv *MessageHeaders) Append(name string, value string) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(name)
	inArgs[2].SetString(value)

	err := messageHeadersAppendFunction_Set()
	if err == nil {
		messageHeadersAppendFunction.Invoke(inArgs[:], nil)
	}

}

var messageHeadersCleanConnectionHeadersFunction *gi.Function
var messageHeadersCleanConnectionHeadersFunction_Once sync.Once

func messageHeadersCleanConnectionHeadersFunction_Set() error {
	var err error
	messageHeadersCleanConnectionHeadersFunction_Once.Do(func() {
		err = messageHeadersStruct_Set()
		if err != nil {
			return
		}
		messageHeadersCleanConnectionHeadersFunction, err = messageHeadersStruct.InvokerNew("clean_connection_headers")
	})
	return err
}

// CleanConnectionHeaders is a representation of the C type soup_message_headers_clean_connection_headers.
func (recv *MessageHeaders) CleanConnectionHeaders() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := messageHeadersCleanConnectionHeadersFunction_Set()
	if err == nil {
		messageHeadersCleanConnectionHeadersFunction.Invoke(inArgs[:], nil)
	}

}

var messageHeadersClearFunction *gi.Function
var messageHeadersClearFunction_Once sync.Once

func messageHeadersClearFunction_Set() error {
	var err error
	messageHeadersClearFunction_Once.Do(func() {
		err = messageHeadersStruct_Set()
		if err != nil {
			return
		}
		messageHeadersClearFunction, err = messageHeadersStruct.InvokerNew("clear")
	})
	return err
}

// Clear is a representation of the C type soup_message_headers_clear.
func (recv *MessageHeaders) Clear() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := messageHeadersClearFunction_Set()
	if err == nil {
		messageHeadersClearFunction.Invoke(inArgs[:], nil)
	}

}

// UNSUPPORTED : C value 'soup_message_headers_foreach' : parameter 'func' of type 'MessageHeadersForeachFunc' not supported

var messageHeadersFreeFunction *gi.Function
var messageHeadersFreeFunction_Once sync.Once

func messageHeadersFreeFunction_Set() error {
	var err error
	messageHeadersFreeFunction_Once.Do(func() {
		err = messageHeadersStruct_Set()
		if err != nil {
			return
		}
		messageHeadersFreeFunction, err = messageHeadersStruct.InvokerNew("free")
	})
	return err
}

// Free is a representation of the C type soup_message_headers_free.
func (recv *MessageHeaders) Free() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := messageHeadersFreeFunction_Set()
	if err == nil {
		messageHeadersFreeFunction.Invoke(inArgs[:], nil)
	}

}

// UNSUPPORTED : C value 'soup_message_headers_free_ranges' : parameter 'ranges' of type 'Range' not supported

var messageHeadersGetFunction *gi.Function
var messageHeadersGetFunction_Once sync.Once

func messageHeadersGetFunction_Set() error {
	var err error
	messageHeadersGetFunction_Once.Do(func() {
		err = messageHeadersStruct_Set()
		if err != nil {
			return
		}
		messageHeadersGetFunction, err = messageHeadersStruct.InvokerNew("get")
	})
	return err
}

// Get is a representation of the C type soup_message_headers_get.
func (recv *MessageHeaders) Get(name string) string {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(name)

	var ret gi.Argument

	err := messageHeadersGetFunction_Set()
	if err == nil {
		ret = messageHeadersGetFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'soup_message_headers_get_content_disposition' : parameter 'params' of type 'GLib.HashTable' not supported

var messageHeadersGetContentLengthFunction *gi.Function
var messageHeadersGetContentLengthFunction_Once sync.Once

func messageHeadersGetContentLengthFunction_Set() error {
	var err error
	messageHeadersGetContentLengthFunction_Once.Do(func() {
		err = messageHeadersStruct_Set()
		if err != nil {
			return
		}
		messageHeadersGetContentLengthFunction, err = messageHeadersStruct.InvokerNew("get_content_length")
	})
	return err
}

// GetContentLength is a representation of the C type soup_message_headers_get_content_length.
func (recv *MessageHeaders) GetContentLength() int64 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := messageHeadersGetContentLengthFunction_Set()
	if err == nil {
		ret = messageHeadersGetContentLengthFunction.Invoke(inArgs[:], nil)
	}

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

func messageHeadersGetListFunction_Set() error {
	var err error
	messageHeadersGetListFunction_Once.Do(func() {
		err = messageHeadersStruct_Set()
		if err != nil {
			return
		}
		messageHeadersGetListFunction, err = messageHeadersStruct.InvokerNew("get_list")
	})
	return err
}

// GetList is a representation of the C type soup_message_headers_get_list.
func (recv *MessageHeaders) GetList(name string) string {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(name)

	var ret gi.Argument

	err := messageHeadersGetListFunction_Set()
	if err == nil {
		ret = messageHeadersGetListFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var messageHeadersGetOneFunction *gi.Function
var messageHeadersGetOneFunction_Once sync.Once

func messageHeadersGetOneFunction_Set() error {
	var err error
	messageHeadersGetOneFunction_Once.Do(func() {
		err = messageHeadersStruct_Set()
		if err != nil {
			return
		}
		messageHeadersGetOneFunction, err = messageHeadersStruct.InvokerNew("get_one")
	})
	return err
}

// GetOne is a representation of the C type soup_message_headers_get_one.
func (recv *MessageHeaders) GetOne(name string) string {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(name)

	var ret gi.Argument

	err := messageHeadersGetOneFunction_Set()
	if err == nil {
		ret = messageHeadersGetOneFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'soup_message_headers_get_ranges' : parameter 'ranges' has no type

// UNSUPPORTED : C value 'soup_message_headers_header_contains' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'soup_message_headers_header_equals' : return type 'gboolean' not supported

var messageHeadersRemoveFunction *gi.Function
var messageHeadersRemoveFunction_Once sync.Once

func messageHeadersRemoveFunction_Set() error {
	var err error
	messageHeadersRemoveFunction_Once.Do(func() {
		err = messageHeadersStruct_Set()
		if err != nil {
			return
		}
		messageHeadersRemoveFunction, err = messageHeadersStruct.InvokerNew("remove")
	})
	return err
}

// Remove is a representation of the C type soup_message_headers_remove.
func (recv *MessageHeaders) Remove(name string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(name)

	err := messageHeadersRemoveFunction_Set()
	if err == nil {
		messageHeadersRemoveFunction.Invoke(inArgs[:], nil)
	}

}

var messageHeadersReplaceFunction *gi.Function
var messageHeadersReplaceFunction_Once sync.Once

func messageHeadersReplaceFunction_Set() error {
	var err error
	messageHeadersReplaceFunction_Once.Do(func() {
		err = messageHeadersStruct_Set()
		if err != nil {
			return
		}
		messageHeadersReplaceFunction, err = messageHeadersStruct.InvokerNew("replace")
	})
	return err
}

// Replace is a representation of the C type soup_message_headers_replace.
func (recv *MessageHeaders) Replace(name string, value string) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(name)
	inArgs[2].SetString(value)

	err := messageHeadersReplaceFunction_Set()
	if err == nil {
		messageHeadersReplaceFunction.Invoke(inArgs[:], nil)
	}

}

// UNSUPPORTED : C value 'soup_message_headers_set_content_disposition' : parameter 'params' of type 'GLib.HashTable' not supported

var messageHeadersSetContentLengthFunction *gi.Function
var messageHeadersSetContentLengthFunction_Once sync.Once

func messageHeadersSetContentLengthFunction_Set() error {
	var err error
	messageHeadersSetContentLengthFunction_Once.Do(func() {
		err = messageHeadersStruct_Set()
		if err != nil {
			return
		}
		messageHeadersSetContentLengthFunction, err = messageHeadersStruct.InvokerNew("set_content_length")
	})
	return err
}

// SetContentLength is a representation of the C type soup_message_headers_set_content_length.
func (recv *MessageHeaders) SetContentLength(contentLength int64) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt64(contentLength)

	err := messageHeadersSetContentLengthFunction_Set()
	if err == nil {
		messageHeadersSetContentLengthFunction.Invoke(inArgs[:], nil)
	}

}

var messageHeadersSetContentRangeFunction *gi.Function
var messageHeadersSetContentRangeFunction_Once sync.Once

func messageHeadersSetContentRangeFunction_Set() error {
	var err error
	messageHeadersSetContentRangeFunction_Once.Do(func() {
		err = messageHeadersStruct_Set()
		if err != nil {
			return
		}
		messageHeadersSetContentRangeFunction, err = messageHeadersStruct.InvokerNew("set_content_range")
	})
	return err
}

// SetContentRange is a representation of the C type soup_message_headers_set_content_range.
func (recv *MessageHeaders) SetContentRange(start int64, end int64, totalLength int64) {
	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt64(start)
	inArgs[2].SetInt64(end)
	inArgs[3].SetInt64(totalLength)

	err := messageHeadersSetContentRangeFunction_Set()
	if err == nil {
		messageHeadersSetContentRangeFunction.Invoke(inArgs[:], nil)
	}

}

// UNSUPPORTED : C value 'soup_message_headers_set_content_type' : parameter 'params' of type 'GLib.HashTable' not supported

// UNSUPPORTED : C value 'soup_message_headers_set_encoding' : parameter 'encoding' of type 'Encoding' not supported

// UNSUPPORTED : C value 'soup_message_headers_set_expectations' : parameter 'expectations' of type 'Expectation' not supported

var messageHeadersSetRangeFunction *gi.Function
var messageHeadersSetRangeFunction_Once sync.Once

func messageHeadersSetRangeFunction_Set() error {
	var err error
	messageHeadersSetRangeFunction_Once.Do(func() {
		err = messageHeadersStruct_Set()
		if err != nil {
			return
		}
		messageHeadersSetRangeFunction, err = messageHeadersStruct.InvokerNew("set_range")
	})
	return err
}

// SetRange is a representation of the C type soup_message_headers_set_range.
func (recv *MessageHeaders) SetRange(start int64, end int64) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt64(start)
	inArgs[2].SetInt64(end)

	err := messageHeadersSetRangeFunction_Set()
	if err == nil {
		messageHeadersSetRangeFunction.Invoke(inArgs[:], nil)
	}

}

// UNSUPPORTED : C value 'soup_message_headers_set_ranges' : parameter 'ranges' of type 'Range' not supported

var messageHeadersIterStruct *gi.Struct
var messageHeadersIterStruct_Once sync.Once

func messageHeadersIterStruct_Set() error {
	var err error
	messageHeadersIterStruct_Once.Do(func() {
		messageHeadersIterStruct, err = gi.StructNew("Soup", "MessageHeadersIter")
	})
	return err
}

type MessageHeadersIter struct {
	native uintptr
}

// UNSUPPORTED : C value 'soup_message_headers_iter_next' : return type 'gboolean' not supported

var messageQueueStruct *gi.Struct
var messageQueueStruct_Once sync.Once

func messageQueueStruct_Set() error {
	var err error
	messageQueueStruct_Once.Do(func() {
		messageQueueStruct, err = gi.StructNew("Soup", "MessageQueue")
	})
	return err
}

type MessageQueue struct {
	native uintptr
}

var messageQueueItemStruct *gi.Struct
var messageQueueItemStruct_Once sync.Once

func messageQueueItemStruct_Set() error {
	var err error
	messageQueueItemStruct_Once.Do(func() {
		messageQueueItemStruct, err = gi.StructNew("Soup", "MessageQueueItem")
	})
	return err
}

type MessageQueueItem struct {
	native uintptr
}

var multipartStruct *gi.Struct
var multipartStruct_Once sync.Once

func multipartStruct_Set() error {
	var err error
	multipartStruct_Once.Do(func() {
		multipartStruct, err = gi.StructNew("Soup", "Multipart")
	})
	return err
}

type Multipart struct {
	native uintptr
}

var multipartNewFunction *gi.Function
var multipartNewFunction_Once sync.Once

func multipartNewFunction_Set() error {
	var err error
	multipartNewFunction_Once.Do(func() {
		err = multipartStruct_Set()
		if err != nil {
			return
		}
		multipartNewFunction, err = multipartStruct.InvokerNew("new")
	})
	return err
}

// MultipartNew is a representation of the C type soup_multipart_new.
func MultipartNew(mimeType string) *Multipart {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(mimeType)

	var ret gi.Argument

	err := multipartNewFunction_Set()
	if err == nil {
		ret = multipartNewFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Multipart{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'soup_multipart_new_from_message' : parameter 'headers' of type 'MessageHeaders' not supported

// UNSUPPORTED : C value 'soup_multipart_append_form_file' : parameter 'body' of type 'Buffer' not supported

var multipartAppendFormStringFunction *gi.Function
var multipartAppendFormStringFunction_Once sync.Once

func multipartAppendFormStringFunction_Set() error {
	var err error
	multipartAppendFormStringFunction_Once.Do(func() {
		err = multipartStruct_Set()
		if err != nil {
			return
		}
		multipartAppendFormStringFunction, err = multipartStruct.InvokerNew("append_form_string")
	})
	return err
}

// AppendFormString is a representation of the C type soup_multipart_append_form_string.
func (recv *Multipart) AppendFormString(controlName string, data string) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(controlName)
	inArgs[2].SetString(data)

	err := multipartAppendFormStringFunction_Set()
	if err == nil {
		multipartAppendFormStringFunction.Invoke(inArgs[:], nil)
	}

}

// UNSUPPORTED : C value 'soup_multipart_append_part' : parameter 'headers' of type 'MessageHeaders' not supported

var multipartFreeFunction *gi.Function
var multipartFreeFunction_Once sync.Once

func multipartFreeFunction_Set() error {
	var err error
	multipartFreeFunction_Once.Do(func() {
		err = multipartStruct_Set()
		if err != nil {
			return
		}
		multipartFreeFunction, err = multipartStruct.InvokerNew("free")
	})
	return err
}

// Free is a representation of the C type soup_multipart_free.
func (recv *Multipart) Free() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := multipartFreeFunction_Set()
	if err == nil {
		multipartFreeFunction.Invoke(inArgs[:], nil)
	}

}

var multipartGetLengthFunction *gi.Function
var multipartGetLengthFunction_Once sync.Once

func multipartGetLengthFunction_Set() error {
	var err error
	multipartGetLengthFunction_Once.Do(func() {
		err = multipartStruct_Set()
		if err != nil {
			return
		}
		multipartGetLengthFunction, err = multipartStruct.InvokerNew("get_length")
	})
	return err
}

// GetLength is a representation of the C type soup_multipart_get_length.
func (recv *Multipart) GetLength() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := multipartGetLengthFunction_Set()
	if err == nil {
		ret = multipartGetLengthFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

// UNSUPPORTED : C value 'soup_multipart_get_part' : parameter 'headers' of type 'MessageHeaders' not supported

// UNSUPPORTED : C value 'soup_multipart_to_message' : parameter 'dest_headers' of type 'MessageHeaders' not supported

var multipartInputStreamClassStruct *gi.Struct
var multipartInputStreamClassStruct_Once sync.Once

func multipartInputStreamClassStruct_Set() error {
	var err error
	multipartInputStreamClassStruct_Once.Do(func() {
		multipartInputStreamClassStruct, err = gi.StructNew("Soup", "MultipartInputStreamClass")
	})
	return err
}

type MultipartInputStreamClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'Gio.FilterInputStreamClass'
}

var multipartInputStreamPrivateStruct *gi.Struct
var multipartInputStreamPrivateStruct_Once sync.Once

func multipartInputStreamPrivateStruct_Set() error {
	var err error
	multipartInputStreamPrivateStruct_Once.Do(func() {
		multipartInputStreamPrivateStruct, err = gi.StructNew("Soup", "MultipartInputStreamPrivate")
	})
	return err
}

type MultipartInputStreamPrivate struct {
	native uintptr
}

var passwordManagerInterfaceStruct *gi.Struct
var passwordManagerInterfaceStruct_Once sync.Once

func passwordManagerInterfaceStruct_Set() error {
	var err error
	passwordManagerInterfaceStruct_Once.Do(func() {
		passwordManagerInterfaceStruct, err = gi.StructNew("Soup", "PasswordManagerInterface")
	})
	return err
}

type PasswordManagerInterface struct {
	native uintptr
	// UNSUPPORTED : C value 'base' : no Go type for 'GObject.TypeInterface'
	// UNSUPPORTED : C value 'get_passwords_async' : missing Type
	// UNSUPPORTED : C value 'get_passwords_sync' : missing Type
}

var proxyResolverDefaultClassStruct *gi.Struct
var proxyResolverDefaultClassStruct_Once sync.Once

func proxyResolverDefaultClassStruct_Set() error {
	var err error
	proxyResolverDefaultClassStruct_Once.Do(func() {
		proxyResolverDefaultClassStruct, err = gi.StructNew("Soup", "ProxyResolverDefaultClass")
	})
	return err
}

type ProxyResolverDefaultClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
}

var proxyResolverInterfaceStruct *gi.Struct
var proxyResolverInterfaceStruct_Once sync.Once

func proxyResolverInterfaceStruct_Set() error {
	var err error
	proxyResolverInterfaceStruct_Once.Do(func() {
		proxyResolverInterfaceStruct, err = gi.StructNew("Soup", "ProxyResolverInterface")
	})
	return err
}

type ProxyResolverInterface struct {
	native uintptr
	// UNSUPPORTED : C value 'base' : no Go type for 'GObject.TypeInterface'
	// UNSUPPORTED : C value 'get_proxy_async' : missing Type
	// UNSUPPORTED : C value 'get_proxy_sync' : missing Type
}

var proxyURIResolverInterfaceStruct *gi.Struct
var proxyURIResolverInterfaceStruct_Once sync.Once

func proxyURIResolverInterfaceStruct_Set() error {
	var err error
	proxyURIResolverInterfaceStruct_Once.Do(func() {
		proxyURIResolverInterfaceStruct, err = gi.StructNew("Soup", "ProxyURIResolverInterface")
	})
	return err
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

func rangeStruct_Set() error {
	var err error
	rangeStruct_Once.Do(func() {
		rangeStruct, err = gi.StructNew("Soup", "Range")
	})
	return err
}

type Range struct {
	native uintptr
	Start  int64
	End    int64
}

var requestClassStruct *gi.Struct
var requestClassStruct_Once sync.Once

func requestClassStruct_Set() error {
	var err error
	requestClassStruct_Once.Do(func() {
		requestClassStruct, err = gi.StructNew("Soup", "RequestClass")
	})
	return err
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

func requestDataClassStruct_Set() error {
	var err error
	requestDataClassStruct_Once.Do(func() {
		requestDataClassStruct, err = gi.StructNew("Soup", "RequestDataClass")
	})
	return err
}

type RequestDataClass struct {
	native uintptr
	Parent *RequestClass
}

var requestDataPrivateStruct *gi.Struct
var requestDataPrivateStruct_Once sync.Once

func requestDataPrivateStruct_Set() error {
	var err error
	requestDataPrivateStruct_Once.Do(func() {
		requestDataPrivateStruct, err = gi.StructNew("Soup", "RequestDataPrivate")
	})
	return err
}

type RequestDataPrivate struct {
	native uintptr
}

var requestFileClassStruct *gi.Struct
var requestFileClassStruct_Once sync.Once

func requestFileClassStruct_Set() error {
	var err error
	requestFileClassStruct_Once.Do(func() {
		requestFileClassStruct, err = gi.StructNew("Soup", "RequestFileClass")
	})
	return err
}

type RequestFileClass struct {
	native uintptr
	Parent *RequestClass
}

var requestFilePrivateStruct *gi.Struct
var requestFilePrivateStruct_Once sync.Once

func requestFilePrivateStruct_Set() error {
	var err error
	requestFilePrivateStruct_Once.Do(func() {
		requestFilePrivateStruct, err = gi.StructNew("Soup", "RequestFilePrivate")
	})
	return err
}

type RequestFilePrivate struct {
	native uintptr
}

var requestHTTPClassStruct *gi.Struct
var requestHTTPClassStruct_Once sync.Once

func requestHTTPClassStruct_Set() error {
	var err error
	requestHTTPClassStruct_Once.Do(func() {
		requestHTTPClassStruct, err = gi.StructNew("Soup", "RequestHTTPClass")
	})
	return err
}

type RequestHTTPClass struct {
	native uintptr
	Parent *RequestClass
}

var requestHTTPPrivateStruct *gi.Struct
var requestHTTPPrivateStruct_Once sync.Once

func requestHTTPPrivateStruct_Set() error {
	var err error
	requestHTTPPrivateStruct_Once.Do(func() {
		requestHTTPPrivateStruct, err = gi.StructNew("Soup", "RequestHTTPPrivate")
	})
	return err
}

type RequestHTTPPrivate struct {
	native uintptr
}

var requestPrivateStruct *gi.Struct
var requestPrivateStruct_Once sync.Once

func requestPrivateStruct_Set() error {
	var err error
	requestPrivateStruct_Once.Do(func() {
		requestPrivateStruct, err = gi.StructNew("Soup", "RequestPrivate")
	})
	return err
}

type RequestPrivate struct {
	native uintptr
}

var requesterClassStruct *gi.Struct
var requesterClassStruct_Once sync.Once

func requesterClassStruct_Set() error {
	var err error
	requesterClassStruct_Once.Do(func() {
		requesterClassStruct, err = gi.StructNew("Soup", "RequesterClass")
	})
	return err
}

type RequesterClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
}

var requesterPrivateStruct *gi.Struct
var requesterPrivateStruct_Once sync.Once

func requesterPrivateStruct_Set() error {
	var err error
	requesterPrivateStruct_Once.Do(func() {
		requesterPrivateStruct, err = gi.StructNew("Soup", "RequesterPrivate")
	})
	return err
}

type RequesterPrivate struct {
	native uintptr
}

var serverClassStruct *gi.Struct
var serverClassStruct_Once sync.Once

func serverClassStruct_Set() error {
	var err error
	serverClassStruct_Once.Do(func() {
		serverClassStruct, err = gi.StructNew("Soup", "ServerClass")
	})
	return err
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

func sessionAsyncClassStruct_Set() error {
	var err error
	sessionAsyncClassStruct_Once.Do(func() {
		sessionAsyncClassStruct, err = gi.StructNew("Soup", "SessionAsyncClass")
	})
	return err
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

func sessionClassStruct_Set() error {
	var err error
	sessionClassStruct_Once.Do(func() {
		sessionClassStruct, err = gi.StructNew("Soup", "SessionClass")
	})
	return err
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

func sessionFeatureInterfaceStruct_Set() error {
	var err error
	sessionFeatureInterfaceStruct_Once.Do(func() {
		sessionFeatureInterfaceStruct, err = gi.StructNew("Soup", "SessionFeatureInterface")
	})
	return err
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

func sessionSyncClassStruct_Set() error {
	var err error
	sessionSyncClassStruct_Once.Do(func() {
		sessionSyncClassStruct, err = gi.StructNew("Soup", "SessionSyncClass")
	})
	return err
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

func socketClassStruct_Set() error {
	var err error
	socketClassStruct_Once.Do(func() {
		socketClassStruct, err = gi.StructNew("Soup", "SocketClass")
	})
	return err
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

func uRIStruct_Set() error {
	var err error
	uRIStruct_Once.Do(func() {
		uRIStruct, err = gi.StructNew("Soup", "URI")
	})
	return err
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

func uRINewFunction_Set() error {
	var err error
	uRINewFunction_Once.Do(func() {
		err = uRIStruct_Set()
		if err != nil {
			return
		}
		uRINewFunction, err = uRIStruct.InvokerNew("new")
	})
	return err
}

// URINew is a representation of the C type soup_uri_new.
func URINew(uriString string) *URI {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(uriString)

	var ret gi.Argument

	err := uRINewFunction_Set()
	if err == nil {
		ret = uRINewFunction.Invoke(inArgs[:], nil)
	}

	retGo := &URI{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'soup_uri_new_with_base' : parameter 'base' of type 'URI' not supported

var uRICopyFunction *gi.Function
var uRICopyFunction_Once sync.Once

func uRICopyFunction_Set() error {
	var err error
	uRICopyFunction_Once.Do(func() {
		err = uRIStruct_Set()
		if err != nil {
			return
		}
		uRICopyFunction, err = uRIStruct.InvokerNew("copy")
	})
	return err
}

// Copy is a representation of the C type soup_uri_copy.
func (recv *URI) Copy() *URI {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := uRICopyFunction_Set()
	if err == nil {
		ret = uRICopyFunction.Invoke(inArgs[:], nil)
	}

	retGo := &URI{native: ret.Pointer()}

	return retGo
}

var uRICopyHostFunction *gi.Function
var uRICopyHostFunction_Once sync.Once

func uRICopyHostFunction_Set() error {
	var err error
	uRICopyHostFunction_Once.Do(func() {
		err = uRIStruct_Set()
		if err != nil {
			return
		}
		uRICopyHostFunction, err = uRIStruct.InvokerNew("copy_host")
	})
	return err
}

// CopyHost is a representation of the C type soup_uri_copy_host.
func (recv *URI) CopyHost() *URI {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := uRICopyHostFunction_Set()
	if err == nil {
		ret = uRICopyHostFunction.Invoke(inArgs[:], nil)
	}

	retGo := &URI{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'soup_uri_equal' : parameter 'uri2' of type 'URI' not supported

var uRIFreeFunction *gi.Function
var uRIFreeFunction_Once sync.Once

func uRIFreeFunction_Set() error {
	var err error
	uRIFreeFunction_Once.Do(func() {
		err = uRIStruct_Set()
		if err != nil {
			return
		}
		uRIFreeFunction, err = uRIStruct.InvokerNew("free")
	})
	return err
}

// Free is a representation of the C type soup_uri_free.
func (recv *URI) Free() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := uRIFreeFunction_Set()
	if err == nil {
		uRIFreeFunction.Invoke(inArgs[:], nil)
	}

}

var uRIGetFragmentFunction *gi.Function
var uRIGetFragmentFunction_Once sync.Once

func uRIGetFragmentFunction_Set() error {
	var err error
	uRIGetFragmentFunction_Once.Do(func() {
		err = uRIStruct_Set()
		if err != nil {
			return
		}
		uRIGetFragmentFunction, err = uRIStruct.InvokerNew("get_fragment")
	})
	return err
}

// GetFragment is a representation of the C type soup_uri_get_fragment.
func (recv *URI) GetFragment() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := uRIGetFragmentFunction_Set()
	if err == nil {
		ret = uRIGetFragmentFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var uRIGetHostFunction *gi.Function
var uRIGetHostFunction_Once sync.Once

func uRIGetHostFunction_Set() error {
	var err error
	uRIGetHostFunction_Once.Do(func() {
		err = uRIStruct_Set()
		if err != nil {
			return
		}
		uRIGetHostFunction, err = uRIStruct.InvokerNew("get_host")
	})
	return err
}

// GetHost is a representation of the C type soup_uri_get_host.
func (recv *URI) GetHost() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := uRIGetHostFunction_Set()
	if err == nil {
		ret = uRIGetHostFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var uRIGetPasswordFunction *gi.Function
var uRIGetPasswordFunction_Once sync.Once

func uRIGetPasswordFunction_Set() error {
	var err error
	uRIGetPasswordFunction_Once.Do(func() {
		err = uRIStruct_Set()
		if err != nil {
			return
		}
		uRIGetPasswordFunction, err = uRIStruct.InvokerNew("get_password")
	})
	return err
}

// GetPassword is a representation of the C type soup_uri_get_password.
func (recv *URI) GetPassword() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := uRIGetPasswordFunction_Set()
	if err == nil {
		ret = uRIGetPasswordFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var uRIGetPathFunction *gi.Function
var uRIGetPathFunction_Once sync.Once

func uRIGetPathFunction_Set() error {
	var err error
	uRIGetPathFunction_Once.Do(func() {
		err = uRIStruct_Set()
		if err != nil {
			return
		}
		uRIGetPathFunction, err = uRIStruct.InvokerNew("get_path")
	})
	return err
}

// GetPath is a representation of the C type soup_uri_get_path.
func (recv *URI) GetPath() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := uRIGetPathFunction_Set()
	if err == nil {
		ret = uRIGetPathFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var uRIGetPortFunction *gi.Function
var uRIGetPortFunction_Once sync.Once

func uRIGetPortFunction_Set() error {
	var err error
	uRIGetPortFunction_Once.Do(func() {
		err = uRIStruct_Set()
		if err != nil {
			return
		}
		uRIGetPortFunction, err = uRIStruct.InvokerNew("get_port")
	})
	return err
}

// GetPort is a representation of the C type soup_uri_get_port.
func (recv *URI) GetPort() uint32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := uRIGetPortFunction_Set()
	if err == nil {
		ret = uRIGetPortFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo
}

var uRIGetQueryFunction *gi.Function
var uRIGetQueryFunction_Once sync.Once

func uRIGetQueryFunction_Set() error {
	var err error
	uRIGetQueryFunction_Once.Do(func() {
		err = uRIStruct_Set()
		if err != nil {
			return
		}
		uRIGetQueryFunction, err = uRIStruct.InvokerNew("get_query")
	})
	return err
}

// GetQuery is a representation of the C type soup_uri_get_query.
func (recv *URI) GetQuery() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := uRIGetQueryFunction_Set()
	if err == nil {
		ret = uRIGetQueryFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var uRIGetSchemeFunction *gi.Function
var uRIGetSchemeFunction_Once sync.Once

func uRIGetSchemeFunction_Set() error {
	var err error
	uRIGetSchemeFunction_Once.Do(func() {
		err = uRIStruct_Set()
		if err != nil {
			return
		}
		uRIGetSchemeFunction, err = uRIStruct.InvokerNew("get_scheme")
	})
	return err
}

// GetScheme is a representation of the C type soup_uri_get_scheme.
func (recv *URI) GetScheme() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := uRIGetSchemeFunction_Set()
	if err == nil {
		ret = uRIGetSchemeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var uRIGetUserFunction *gi.Function
var uRIGetUserFunction_Once sync.Once

func uRIGetUserFunction_Set() error {
	var err error
	uRIGetUserFunction_Once.Do(func() {
		err = uRIStruct_Set()
		if err != nil {
			return
		}
		uRIGetUserFunction, err = uRIStruct.InvokerNew("get_user")
	})
	return err
}

// GetUser is a representation of the C type soup_uri_get_user.
func (recv *URI) GetUser() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := uRIGetUserFunction_Set()
	if err == nil {
		ret = uRIGetUserFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'soup_uri_host_equal' : parameter 'v2' of type 'URI' not supported

var uRIHostHashFunction *gi.Function
var uRIHostHashFunction_Once sync.Once

func uRIHostHashFunction_Set() error {
	var err error
	uRIHostHashFunction_Once.Do(func() {
		err = uRIStruct_Set()
		if err != nil {
			return
		}
		uRIHostHashFunction, err = uRIStruct.InvokerNew("host_hash")
	})
	return err
}

// HostHash is a representation of the C type soup_uri_host_hash.
func (recv *URI) HostHash() uint32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := uRIHostHashFunction_Set()
	if err == nil {
		ret = uRIHostHashFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo
}

var uRISetFragmentFunction *gi.Function
var uRISetFragmentFunction_Once sync.Once

func uRISetFragmentFunction_Set() error {
	var err error
	uRISetFragmentFunction_Once.Do(func() {
		err = uRIStruct_Set()
		if err != nil {
			return
		}
		uRISetFragmentFunction, err = uRIStruct.InvokerNew("set_fragment")
	})
	return err
}

// SetFragment is a representation of the C type soup_uri_set_fragment.
func (recv *URI) SetFragment(fragment string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(fragment)

	err := uRISetFragmentFunction_Set()
	if err == nil {
		uRISetFragmentFunction.Invoke(inArgs[:], nil)
	}

}

var uRISetHostFunction *gi.Function
var uRISetHostFunction_Once sync.Once

func uRISetHostFunction_Set() error {
	var err error
	uRISetHostFunction_Once.Do(func() {
		err = uRIStruct_Set()
		if err != nil {
			return
		}
		uRISetHostFunction, err = uRIStruct.InvokerNew("set_host")
	})
	return err
}

// SetHost is a representation of the C type soup_uri_set_host.
func (recv *URI) SetHost(host string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(host)

	err := uRISetHostFunction_Set()
	if err == nil {
		uRISetHostFunction.Invoke(inArgs[:], nil)
	}

}

var uRISetPasswordFunction *gi.Function
var uRISetPasswordFunction_Once sync.Once

func uRISetPasswordFunction_Set() error {
	var err error
	uRISetPasswordFunction_Once.Do(func() {
		err = uRIStruct_Set()
		if err != nil {
			return
		}
		uRISetPasswordFunction, err = uRIStruct.InvokerNew("set_password")
	})
	return err
}

// SetPassword is a representation of the C type soup_uri_set_password.
func (recv *URI) SetPassword(password string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(password)

	err := uRISetPasswordFunction_Set()
	if err == nil {
		uRISetPasswordFunction.Invoke(inArgs[:], nil)
	}

}

var uRISetPathFunction *gi.Function
var uRISetPathFunction_Once sync.Once

func uRISetPathFunction_Set() error {
	var err error
	uRISetPathFunction_Once.Do(func() {
		err = uRIStruct_Set()
		if err != nil {
			return
		}
		uRISetPathFunction, err = uRIStruct.InvokerNew("set_path")
	})
	return err
}

// SetPath is a representation of the C type soup_uri_set_path.
func (recv *URI) SetPath(path string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(path)

	err := uRISetPathFunction_Set()
	if err == nil {
		uRISetPathFunction.Invoke(inArgs[:], nil)
	}

}

var uRISetPortFunction *gi.Function
var uRISetPortFunction_Once sync.Once

func uRISetPortFunction_Set() error {
	var err error
	uRISetPortFunction_Once.Do(func() {
		err = uRIStruct_Set()
		if err != nil {
			return
		}
		uRISetPortFunction, err = uRIStruct.InvokerNew("set_port")
	})
	return err
}

// SetPort is a representation of the C type soup_uri_set_port.
func (recv *URI) SetPort(port uint32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetUint32(port)

	err := uRISetPortFunction_Set()
	if err == nil {
		uRISetPortFunction.Invoke(inArgs[:], nil)
	}

}

var uRISetQueryFunction *gi.Function
var uRISetQueryFunction_Once sync.Once

func uRISetQueryFunction_Set() error {
	var err error
	uRISetQueryFunction_Once.Do(func() {
		err = uRIStruct_Set()
		if err != nil {
			return
		}
		uRISetQueryFunction, err = uRIStruct.InvokerNew("set_query")
	})
	return err
}

// SetQuery is a representation of the C type soup_uri_set_query.
func (recv *URI) SetQuery(query string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(query)

	err := uRISetQueryFunction_Set()
	if err == nil {
		uRISetQueryFunction.Invoke(inArgs[:], nil)
	}

}

// UNSUPPORTED : C value 'soup_uri_set_query_from_fields' : parameter '...' has no type

// UNSUPPORTED : C value 'soup_uri_set_query_from_form' : parameter 'form' of type 'GLib.HashTable' not supported

var uRISetSchemeFunction *gi.Function
var uRISetSchemeFunction_Once sync.Once

func uRISetSchemeFunction_Set() error {
	var err error
	uRISetSchemeFunction_Once.Do(func() {
		err = uRIStruct_Set()
		if err != nil {
			return
		}
		uRISetSchemeFunction, err = uRIStruct.InvokerNew("set_scheme")
	})
	return err
}

// SetScheme is a representation of the C type soup_uri_set_scheme.
func (recv *URI) SetScheme(scheme string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(scheme)

	err := uRISetSchemeFunction_Set()
	if err == nil {
		uRISetSchemeFunction.Invoke(inArgs[:], nil)
	}

}

var uRISetUserFunction *gi.Function
var uRISetUserFunction_Once sync.Once

func uRISetUserFunction_Set() error {
	var err error
	uRISetUserFunction_Once.Do(func() {
		err = uRIStruct_Set()
		if err != nil {
			return
		}
		uRISetUserFunction, err = uRIStruct.InvokerNew("set_user")
	})
	return err
}

// SetUser is a representation of the C type soup_uri_set_user.
func (recv *URI) SetUser(user string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(user)

	err := uRISetUserFunction_Set()
	if err == nil {
		uRISetUserFunction.Invoke(inArgs[:], nil)
	}

}

// UNSUPPORTED : C value 'soup_uri_to_string' : parameter 'just_path_and_query' of type 'gboolean' not supported

// UNSUPPORTED : C value 'soup_uri_uses_default_port' : return type 'gboolean' not supported

var websocketConnectionClassStruct *gi.Struct
var websocketConnectionClassStruct_Once sync.Once

func websocketConnectionClassStruct_Set() error {
	var err error
	websocketConnectionClassStruct_Once.Do(func() {
		websocketConnectionClassStruct, err = gi.StructNew("Soup", "WebsocketConnectionClass")
	})
	return err
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

func websocketConnectionPrivateStruct_Set() error {
	var err error
	websocketConnectionPrivateStruct_Once.Do(func() {
		websocketConnectionPrivateStruct, err = gi.StructNew("Soup", "WebsocketConnectionPrivate")
	})
	return err
}

type WebsocketConnectionPrivate struct {
	native uintptr
}

var websocketExtensionClassStruct *gi.Struct
var websocketExtensionClassStruct_Once sync.Once

func websocketExtensionClassStruct_Set() error {
	var err error
	websocketExtensionClassStruct_Once.Do(func() {
		websocketExtensionClassStruct, err = gi.StructNew("Soup", "WebsocketExtensionClass")
	})
	return err
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

func websocketExtensionDeflateClassStruct_Set() error {
	var err error
	websocketExtensionDeflateClassStruct_Once.Do(func() {
		websocketExtensionDeflateClassStruct, err = gi.StructNew("Soup", "WebsocketExtensionDeflateClass")
	})
	return err
}

type WebsocketExtensionDeflateClass struct {
	native      uintptr
	ParentClass *WebsocketExtensionClass
}

var websocketExtensionManagerClassStruct *gi.Struct
var websocketExtensionManagerClassStruct_Once sync.Once

func websocketExtensionManagerClassStruct_Set() error {
	var err error
	websocketExtensionManagerClassStruct_Once.Do(func() {
		websocketExtensionManagerClassStruct, err = gi.StructNew("Soup", "WebsocketExtensionManagerClass")
	})
	return err
}

type WebsocketExtensionManagerClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
}

var xMLRPCParamsStruct *gi.Struct
var xMLRPCParamsStruct_Once sync.Once

func xMLRPCParamsStruct_Set() error {
	var err error
	xMLRPCParamsStruct_Once.Do(func() {
		xMLRPCParamsStruct, err = gi.StructNew("Soup", "XMLRPCParams")
	})
	return err
}

type XMLRPCParams struct {
	native uintptr
}

var xMLRPCParamsFreeFunction *gi.Function
var xMLRPCParamsFreeFunction_Once sync.Once

func xMLRPCParamsFreeFunction_Set() error {
	var err error
	xMLRPCParamsFreeFunction_Once.Do(func() {
		err = xMLRPCParamsStruct_Set()
		if err != nil {
			return
		}
		xMLRPCParamsFreeFunction, err = xMLRPCParamsStruct.InvokerNew("free")
	})
	return err
}

// Free is a representation of the C type soup_xmlrpc_params_free.
func (recv *XMLRPCParams) Free() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := xMLRPCParamsFreeFunction_Set()
	if err == nil {
		xMLRPCParamsFreeFunction.Invoke(inArgs[:], nil)
	}

}

// UNSUPPORTED : C value 'soup_xmlrpc_params_parse' : return type 'GLib.Variant' not supported
