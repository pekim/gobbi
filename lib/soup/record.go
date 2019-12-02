// Code generated - DO NOT EDIT.

package soup

import (
	gi "github.com/pekim/gobbi/internal/gi"
	"runtime"
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
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value '_libsoup_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved4' : for field setter : missing Type

// AddressClassStruct creates an uninitialised AddressClass.
func AddressClassStruct() *AddressClass {
	err := addressClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &AddressClass{native: addressClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeAddressClass)
	return structGo
}
func finalizeAddressClass(obj *AddressClass) {
	addressClassStruct.Free(obj.native)
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
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// FieldSchemeName returns the C field 'scheme_name'.
func (recv *AuthClass) FieldSchemeName() string {
	argValue := gi.FieldGet(authClassStruct, recv.native, "scheme_name")
	value := argValue.String(false)
	return value
}

// SetFieldSchemeName sets the value of the C field 'scheme_name'.
func (recv *AuthClass) SetFieldSchemeName(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.FieldSet(authClassStruct, recv.native, "scheme_name", argValue)
}

// FieldStrength returns the C field 'strength'.
func (recv *AuthClass) FieldStrength() uint32 {
	argValue := gi.FieldGet(authClassStruct, recv.native, "strength")
	value := argValue.Uint32()
	return value
}

// SetFieldStrength sets the value of the C field 'strength'.
func (recv *AuthClass) SetFieldStrength(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.FieldSet(authClassStruct, recv.native, "strength", argValue)
}

// UNSUPPORTED : C value 'update' : for field getter : missing Type

// UNSUPPORTED : C value 'update' : for field setter : missing Type

// UNSUPPORTED : C value 'get_protection_space' : for field getter : missing Type

// UNSUPPORTED : C value 'get_protection_space' : for field setter : missing Type

// UNSUPPORTED : C value 'authenticate' : for field getter : missing Type

// UNSUPPORTED : C value 'authenticate' : for field setter : missing Type

// UNSUPPORTED : C value 'is_authenticated' : for field getter : missing Type

// UNSUPPORTED : C value 'is_authenticated' : for field setter : missing Type

// UNSUPPORTED : C value 'get_authorization' : for field getter : missing Type

// UNSUPPORTED : C value 'get_authorization' : for field setter : missing Type

// UNSUPPORTED : C value 'is_ready' : for field getter : missing Type

// UNSUPPORTED : C value 'is_ready' : for field setter : missing Type

// UNSUPPORTED : C value 'can_authenticate' : for field getter : missing Type

// UNSUPPORTED : C value 'can_authenticate' : for field setter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved4' : for field setter : missing Type

// AuthClassStruct creates an uninitialised AuthClass.
func AuthClassStruct() *AuthClass {
	err := authClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &AuthClass{native: authClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeAuthClass)
	return structGo
}
func finalizeAuthClass(obj *AuthClass) {
	authClassStruct.Free(obj.native)
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
	native uintptr
}

// FieldParentClass returns the C field 'parent_class'.
func (recv *AuthDomainBasicClass) FieldParentClass() *AuthDomainClass {
	argValue := gi.FieldGet(authDomainBasicClassStruct, recv.native, "parent_class")
	value := &AuthDomainClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *AuthDomainBasicClass) SetFieldParentClass(value *AuthDomainClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(authDomainBasicClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value '_libsoup_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved4' : for field setter : missing Type

// AuthDomainBasicClassStruct creates an uninitialised AuthDomainBasicClass.
func AuthDomainBasicClassStruct() *AuthDomainBasicClass {
	err := authDomainBasicClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &AuthDomainBasicClass{native: authDomainBasicClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeAuthDomainBasicClass)
	return structGo
}
func finalizeAuthDomainBasicClass(obj *AuthDomainBasicClass) {
	authDomainBasicClassStruct.Free(obj.native)
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
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'accepts' : for field getter : missing Type

// UNSUPPORTED : C value 'accepts' : for field setter : missing Type

// UNSUPPORTED : C value 'challenge' : for field getter : missing Type

// UNSUPPORTED : C value 'challenge' : for field setter : missing Type

// UNSUPPORTED : C value 'check_password' : for field getter : missing Type

// UNSUPPORTED : C value 'check_password' : for field setter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved4' : for field setter : missing Type

// AuthDomainClassStruct creates an uninitialised AuthDomainClass.
func AuthDomainClassStruct() *AuthDomainClass {
	err := authDomainClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &AuthDomainClass{native: authDomainClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeAuthDomainClass)
	return structGo
}
func finalizeAuthDomainClass(obj *AuthDomainClass) {
	authDomainClassStruct.Free(obj.native)
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
	native uintptr
}

// FieldParentClass returns the C field 'parent_class'.
func (recv *AuthDomainDigestClass) FieldParentClass() *AuthDomainClass {
	argValue := gi.FieldGet(authDomainDigestClassStruct, recv.native, "parent_class")
	value := &AuthDomainClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *AuthDomainDigestClass) SetFieldParentClass(value *AuthDomainClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(authDomainDigestClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value '_libsoup_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved4' : for field setter : missing Type

// AuthDomainDigestClassStruct creates an uninitialised AuthDomainDigestClass.
func AuthDomainDigestClassStruct() *AuthDomainDigestClass {
	err := authDomainDigestClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &AuthDomainDigestClass{native: authDomainDigestClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeAuthDomainDigestClass)
	return structGo
}
func finalizeAuthDomainDigestClass(obj *AuthDomainDigestClass) {
	authDomainDigestClassStruct.Free(obj.native)
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
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'authenticate' : for field getter : missing Type

// UNSUPPORTED : C value 'authenticate' : for field setter : missing Type

// AuthManagerClassStruct creates an uninitialised AuthManagerClass.
func AuthManagerClassStruct() *AuthManagerClass {
	err := authManagerClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &AuthManagerClass{native: authManagerClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeAuthManagerClass)
	return structGo
}
func finalizeAuthManagerClass(obj *AuthManagerClass) {
	authManagerClassStruct.Free(obj.native)
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

// AuthManagerPrivateStruct creates an uninitialised AuthManagerPrivate.
func AuthManagerPrivateStruct() *AuthManagerPrivate {
	err := authManagerPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &AuthManagerPrivate{native: authManagerPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeAuthManagerPrivate)
	return structGo
}
func finalizeAuthManagerPrivate(obj *AuthManagerPrivate) {
	authManagerPrivateStruct.Free(obj.native)
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
}

// UNSUPPORTED : C value 'data' : for field getter : no Go type for 'gpointer'

// UNSUPPORTED : C value 'data' : for field setter : no Go type for 'gpointer'

// FieldLength returns the C field 'length'.
func (recv *Buffer) FieldLength() uint64 {
	argValue := gi.FieldGet(bufferStruct, recv.native, "length")
	value := argValue.Uint64()
	return value
}

// SetFieldLength sets the value of the C field 'length'.
func (recv *Buffer) SetFieldLength(value uint64) {
	var argValue gi.Argument
	argValue.SetUint64(value)
	gi.FieldSet(bufferStruct, recv.native, "length", argValue)
}

// UNSUPPORTED : C value 'soup_buffer_new' : parameter 'use' of type 'MemoryUse' not supported

// UNSUPPORTED : C value 'soup_buffer_new_take' : parameter 'data' of type 'nil' not supported

// UNSUPPORTED : C value 'soup_buffer_new_with_owner' : parameter 'data' of type 'nil' not supported

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

	return
}

// UNSUPPORTED : C value 'soup_buffer_get_as_bytes' : return type 'GLib.Bytes' not supported

// UNSUPPORTED : C value 'soup_buffer_get_data' : parameter 'data' of type 'nil' not supported

// UNSUPPORTED : C value 'soup_buffer_get_owner' : return type 'gpointer' not supported

var bufferNewSubbufferFunction *gi.Function
var bufferNewSubbufferFunction_Once sync.Once

func bufferNewSubbufferFunction_Set() error {
	var err error
	bufferNewSubbufferFunction_Once.Do(func() {
		err = bufferStruct_Set()
		if err != nil {
			return
		}
		bufferNewSubbufferFunction, err = bufferStruct.InvokerNew("new_subbuffer")
	})
	return err
}

// NewSubbuffer is a representation of the C type soup_buffer_new_subbuffer.
func (recv *Buffer) NewSubbuffer(offset uint64, length uint64) *Buffer {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetUint64(offset)
	inArgs[2].SetUint64(length)

	var ret gi.Argument

	err := bufferNewSubbufferFunction_Set()
	if err == nil {
		ret = bufferNewSubbufferFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Buffer{native: ret.Pointer()}

	return retGo
}

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
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'get_cacheability' : for field getter : missing Type

// UNSUPPORTED : C value 'get_cacheability' : for field setter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved3' : for field setter : missing Type

// CacheClassStruct creates an uninitialised CacheClass.
func CacheClassStruct() *CacheClass {
	err := cacheClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &CacheClass{native: cacheClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeCacheClass)
	return structGo
}
func finalizeCacheClass(obj *CacheClass) {
	cacheClassStruct.Free(obj.native)
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

// CachePrivateStruct creates an uninitialised CachePrivate.
func CachePrivateStruct() *CachePrivate {
	err := cachePrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &CachePrivate{native: cachePrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeCachePrivate)
	return structGo
}
func finalizeCachePrivate(obj *CachePrivate) {
	cachePrivateStruct.Free(obj.native)
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

var clientContextGetAddressFunction *gi.Function
var clientContextGetAddressFunction_Once sync.Once

func clientContextGetAddressFunction_Set() error {
	var err error
	clientContextGetAddressFunction_Once.Do(func() {
		err = clientContextStruct_Set()
		if err != nil {
			return
		}
		clientContextGetAddressFunction, err = clientContextStruct.InvokerNew("get_address")
	})
	return err
}

// GetAddress is a representation of the C type soup_client_context_get_address.
func (recv *ClientContext) GetAddress() *Address {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := clientContextGetAddressFunction_Set()
	if err == nil {
		ret = clientContextGetAddressFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Address{native: ret.Pointer()}

	return retGo
}

var clientContextGetAuthDomainFunction *gi.Function
var clientContextGetAuthDomainFunction_Once sync.Once

func clientContextGetAuthDomainFunction_Set() error {
	var err error
	clientContextGetAuthDomainFunction_Once.Do(func() {
		err = clientContextStruct_Set()
		if err != nil {
			return
		}
		clientContextGetAuthDomainFunction, err = clientContextStruct.InvokerNew("get_auth_domain")
	})
	return err
}

// GetAuthDomain is a representation of the C type soup_client_context_get_auth_domain.
func (recv *ClientContext) GetAuthDomain() *AuthDomain {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := clientContextGetAuthDomainFunction_Set()
	if err == nil {
		ret = clientContextGetAuthDomainFunction.Invoke(inArgs[:], nil)
	}

	retGo := &AuthDomain{native: ret.Pointer()}

	return retGo
}

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

var clientContextGetSocketFunction *gi.Function
var clientContextGetSocketFunction_Once sync.Once

func clientContextGetSocketFunction_Set() error {
	var err error
	clientContextGetSocketFunction_Once.Do(func() {
		err = clientContextStruct_Set()
		if err != nil {
			return
		}
		clientContextGetSocketFunction, err = clientContextStruct.InvokerNew("get_socket")
	})
	return err
}

// GetSocket is a representation of the C type soup_client_context_get_socket.
func (recv *ClientContext) GetSocket() *Socket {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := clientContextGetSocketFunction_Set()
	if err == nil {
		ret = clientContextGetSocketFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Socket{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'soup_client_context_steal_connection' : return type 'Gio.IOStream' not supported

// ClientContextStruct creates an uninitialised ClientContext.
func ClientContextStruct() *ClientContext {
	err := clientContextStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ClientContext{native: clientContextStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeClientContext)
	return structGo
}
func finalizeClientContext(obj *ClientContext) {
	clientContextStruct.Free(obj.native)
}

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

// ConnectionStruct creates an uninitialised Connection.
func ConnectionStruct() *Connection {
	err := connectionStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &Connection{native: connectionStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeConnection)
	return structGo
}
func finalizeConnection(obj *Connection) {
	connectionStruct.Free(obj.native)
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
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value '_libsoup_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved4' : for field setter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved5' : for field getter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved5' : for field setter : missing Type

// ContentDecoderClassStruct creates an uninitialised ContentDecoderClass.
func ContentDecoderClassStruct() *ContentDecoderClass {
	err := contentDecoderClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ContentDecoderClass{native: contentDecoderClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeContentDecoderClass)
	return structGo
}
func finalizeContentDecoderClass(obj *ContentDecoderClass) {
	contentDecoderClassStruct.Free(obj.native)
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

// ContentDecoderPrivateStruct creates an uninitialised ContentDecoderPrivate.
func ContentDecoderPrivateStruct() *ContentDecoderPrivate {
	err := contentDecoderPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ContentDecoderPrivate{native: contentDecoderPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeContentDecoderPrivate)
	return structGo
}
func finalizeContentDecoderPrivate(obj *ContentDecoderPrivate) {
	contentDecoderPrivateStruct.Free(obj.native)
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
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'sniff' : for field getter : missing Type

// UNSUPPORTED : C value 'sniff' : for field setter : missing Type

// UNSUPPORTED : C value 'get_buffer_size' : for field getter : missing Type

// UNSUPPORTED : C value 'get_buffer_size' : for field setter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved4' : for field setter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved5' : for field getter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved5' : for field setter : missing Type

// ContentSnifferClassStruct creates an uninitialised ContentSnifferClass.
func ContentSnifferClassStruct() *ContentSnifferClass {
	err := contentSnifferClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ContentSnifferClass{native: contentSnifferClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeContentSnifferClass)
	return structGo
}
func finalizeContentSnifferClass(obj *ContentSnifferClass) {
	contentSnifferClassStruct.Free(obj.native)
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

// ContentSnifferPrivateStruct creates an uninitialised ContentSnifferPrivate.
func ContentSnifferPrivateStruct() *ContentSnifferPrivate {
	err := contentSnifferPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ContentSnifferPrivate{native: contentSnifferPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeContentSnifferPrivate)
	return structGo
}
func finalizeContentSnifferPrivate(obj *ContentSnifferPrivate) {
	contentSnifferPrivateStruct.Free(obj.native)
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
	native uintptr
}

// FieldName returns the C field 'name'.
func (recv *Cookie) FieldName() string {
	argValue := gi.FieldGet(cookieStruct, recv.native, "name")
	value := argValue.String(false)
	return value
}

// SetFieldName sets the value of the C field 'name'.
func (recv *Cookie) SetFieldName(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.FieldSet(cookieStruct, recv.native, "name", argValue)
}

// FieldValue returns the C field 'value'.
func (recv *Cookie) FieldValue() string {
	argValue := gi.FieldGet(cookieStruct, recv.native, "value")
	value := argValue.String(false)
	return value
}

// SetFieldValue sets the value of the C field 'value'.
func (recv *Cookie) SetFieldValue(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.FieldSet(cookieStruct, recv.native, "value", argValue)
}

// FieldDomain returns the C field 'domain'.
func (recv *Cookie) FieldDomain() string {
	argValue := gi.FieldGet(cookieStruct, recv.native, "domain")
	value := argValue.String(false)
	return value
}

// SetFieldDomain sets the value of the C field 'domain'.
func (recv *Cookie) SetFieldDomain(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.FieldSet(cookieStruct, recv.native, "domain", argValue)
}

// FieldPath returns the C field 'path'.
func (recv *Cookie) FieldPath() string {
	argValue := gi.FieldGet(cookieStruct, recv.native, "path")
	value := argValue.String(false)
	return value
}

// SetFieldPath sets the value of the C field 'path'.
func (recv *Cookie) SetFieldPath(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.FieldSet(cookieStruct, recv.native, "path", argValue)
}

// FieldExpires returns the C field 'expires'.
func (recv *Cookie) FieldExpires() *Date {
	argValue := gi.FieldGet(cookieStruct, recv.native, "expires")
	value := &Date{native: argValue.Pointer()}
	return value
}

// SetFieldExpires sets the value of the C field 'expires'.
func (recv *Cookie) SetFieldExpires(value *Date) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(cookieStruct, recv.native, "expires", argValue)
}

// FieldSecure returns the C field 'secure'.
func (recv *Cookie) FieldSecure() bool {
	argValue := gi.FieldGet(cookieStruct, recv.native, "secure")
	value := argValue.Boolean()
	return value
}

// SetFieldSecure sets the value of the C field 'secure'.
func (recv *Cookie) SetFieldSecure(value bool) {
	var argValue gi.Argument
	argValue.SetBoolean(value)
	gi.FieldSet(cookieStruct, recv.native, "secure", argValue)
}

// FieldHttpOnly returns the C field 'http_only'.
func (recv *Cookie) FieldHttpOnly() bool {
	argValue := gi.FieldGet(cookieStruct, recv.native, "http_only")
	value := argValue.Boolean()
	return value
}

// SetFieldHttpOnly sets the value of the C field 'http_only'.
func (recv *Cookie) SetFieldHttpOnly(value bool) {
	var argValue gi.Argument
	argValue.SetBoolean(value)
	gi.FieldSet(cookieStruct, recv.native, "http_only", argValue)
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

var cookieAppliesToUriFunction *gi.Function
var cookieAppliesToUriFunction_Once sync.Once

func cookieAppliesToUriFunction_Set() error {
	var err error
	cookieAppliesToUriFunction_Once.Do(func() {
		err = cookieStruct_Set()
		if err != nil {
			return
		}
		cookieAppliesToUriFunction, err = cookieStruct.InvokerNew("applies_to_uri")
	})
	return err
}

// AppliesToUri is a representation of the C type soup_cookie_applies_to_uri.
func (recv *Cookie) AppliesToUri(uri *URI) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(uri.native)

	var ret gi.Argument

	err := cookieAppliesToUriFunction_Set()
	if err == nil {
		ret = cookieAppliesToUriFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

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

var cookieDomainMatchesFunction *gi.Function
var cookieDomainMatchesFunction_Once sync.Once

func cookieDomainMatchesFunction_Set() error {
	var err error
	cookieDomainMatchesFunction_Once.Do(func() {
		err = cookieStruct_Set()
		if err != nil {
			return
		}
		cookieDomainMatchesFunction, err = cookieStruct.InvokerNew("domain_matches")
	})
	return err
}

// DomainMatches is a representation of the C type soup_cookie_domain_matches.
func (recv *Cookie) DomainMatches(host string) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(host)

	var ret gi.Argument

	err := cookieDomainMatchesFunction_Set()
	if err == nil {
		ret = cookieDomainMatchesFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var cookieEqualFunction *gi.Function
var cookieEqualFunction_Once sync.Once

func cookieEqualFunction_Set() error {
	var err error
	cookieEqualFunction_Once.Do(func() {
		err = cookieStruct_Set()
		if err != nil {
			return
		}
		cookieEqualFunction, err = cookieStruct.InvokerNew("equal")
	})
	return err
}

// Equal is a representation of the C type soup_cookie_equal.
func (recv *Cookie) Equal(cookie2 *Cookie) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(cookie2.native)

	var ret gi.Argument

	err := cookieEqualFunction_Set()
	if err == nil {
		ret = cookieEqualFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

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

	return
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

var cookieGetHttpOnlyFunction *gi.Function
var cookieGetHttpOnlyFunction_Once sync.Once

func cookieGetHttpOnlyFunction_Set() error {
	var err error
	cookieGetHttpOnlyFunction_Once.Do(func() {
		err = cookieStruct_Set()
		if err != nil {
			return
		}
		cookieGetHttpOnlyFunction, err = cookieStruct.InvokerNew("get_http_only")
	})
	return err
}

// GetHttpOnly is a representation of the C type soup_cookie_get_http_only.
func (recv *Cookie) GetHttpOnly() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := cookieGetHttpOnlyFunction_Set()
	if err == nil {
		ret = cookieGetHttpOnlyFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

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

var cookieGetSecureFunction *gi.Function
var cookieGetSecureFunction_Once sync.Once

func cookieGetSecureFunction_Set() error {
	var err error
	cookieGetSecureFunction_Once.Do(func() {
		err = cookieStruct_Set()
		if err != nil {
			return
		}
		cookieGetSecureFunction, err = cookieStruct.InvokerNew("get_secure")
	})
	return err
}

// GetSecure is a representation of the C type soup_cookie_get_secure.
func (recv *Cookie) GetSecure() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := cookieGetSecureFunction_Set()
	if err == nil {
		ret = cookieGetSecureFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

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

	return
}

var cookieSetExpiresFunction *gi.Function
var cookieSetExpiresFunction_Once sync.Once

func cookieSetExpiresFunction_Set() error {
	var err error
	cookieSetExpiresFunction_Once.Do(func() {
		err = cookieStruct_Set()
		if err != nil {
			return
		}
		cookieSetExpiresFunction, err = cookieStruct.InvokerNew("set_expires")
	})
	return err
}

// SetExpires is a representation of the C type soup_cookie_set_expires.
func (recv *Cookie) SetExpires(expires *Date) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(expires.native)

	err := cookieSetExpiresFunction_Set()
	if err == nil {
		cookieSetExpiresFunction.Invoke(inArgs[:], nil)
	}

	return
}

var cookieSetHttpOnlyFunction *gi.Function
var cookieSetHttpOnlyFunction_Once sync.Once

func cookieSetHttpOnlyFunction_Set() error {
	var err error
	cookieSetHttpOnlyFunction_Once.Do(func() {
		err = cookieStruct_Set()
		if err != nil {
			return
		}
		cookieSetHttpOnlyFunction, err = cookieStruct.InvokerNew("set_http_only")
	})
	return err
}

// SetHttpOnly is a representation of the C type soup_cookie_set_http_only.
func (recv *Cookie) SetHttpOnly(httpOnly bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetBoolean(httpOnly)

	err := cookieSetHttpOnlyFunction_Set()
	if err == nil {
		cookieSetHttpOnlyFunction.Invoke(inArgs[:], nil)
	}

	return
}

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

	return
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

	return
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

	return
}

var cookieSetSecureFunction *gi.Function
var cookieSetSecureFunction_Once sync.Once

func cookieSetSecureFunction_Set() error {
	var err error
	cookieSetSecureFunction_Once.Do(func() {
		err = cookieStruct_Set()
		if err != nil {
			return
		}
		cookieSetSecureFunction, err = cookieStruct.InvokerNew("set_secure")
	})
	return err
}

// SetSecure is a representation of the C type soup_cookie_set_secure.
func (recv *Cookie) SetSecure(secure bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetBoolean(secure)

	err := cookieSetSecureFunction_Set()
	if err == nil {
		cookieSetSecureFunction.Invoke(inArgs[:], nil)
	}

	return
}

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

	return
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
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'save' : for field getter : missing Type

// UNSUPPORTED : C value 'save' : for field setter : missing Type

// UNSUPPORTED : C value 'is_persistent' : for field getter : missing Type

// UNSUPPORTED : C value 'is_persistent' : for field setter : missing Type

// UNSUPPORTED : C value 'changed' : for field getter : missing Type

// UNSUPPORTED : C value 'changed' : for field setter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved2' : for field setter : missing Type

// CookieJarClassStruct creates an uninitialised CookieJarClass.
func CookieJarClassStruct() *CookieJarClass {
	err := cookieJarClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &CookieJarClass{native: cookieJarClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeCookieJarClass)
	return structGo
}
func finalizeCookieJarClass(obj *CookieJarClass) {
	cookieJarClassStruct.Free(obj.native)
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
	native uintptr
}

// FieldParentClass returns the C field 'parent_class'.
func (recv *CookieJarDBClass) FieldParentClass() *CookieJarClass {
	argValue := gi.FieldGet(cookieJarDBClassStruct, recv.native, "parent_class")
	value := &CookieJarClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *CookieJarDBClass) SetFieldParentClass(value *CookieJarClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(cookieJarDBClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value '_libsoup_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved4' : for field setter : missing Type

// CookieJarDBClassStruct creates an uninitialised CookieJarDBClass.
func CookieJarDBClassStruct() *CookieJarDBClass {
	err := cookieJarDBClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &CookieJarDBClass{native: cookieJarDBClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeCookieJarDBClass)
	return structGo
}
func finalizeCookieJarDBClass(obj *CookieJarDBClass) {
	cookieJarDBClassStruct.Free(obj.native)
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
	native uintptr
}

// FieldParentClass returns the C field 'parent_class'.
func (recv *CookieJarTextClass) FieldParentClass() *CookieJarClass {
	argValue := gi.FieldGet(cookieJarTextClassStruct, recv.native, "parent_class")
	value := &CookieJarClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *CookieJarTextClass) SetFieldParentClass(value *CookieJarClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(cookieJarTextClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value '_libsoup_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved4' : for field setter : missing Type

// CookieJarTextClassStruct creates an uninitialised CookieJarTextClass.
func CookieJarTextClassStruct() *CookieJarTextClass {
	err := cookieJarTextClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &CookieJarTextClass{native: cookieJarTextClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeCookieJarTextClass)
	return structGo
}
func finalizeCookieJarTextClass(obj *CookieJarTextClass) {
	cookieJarTextClassStruct.Free(obj.native)
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
}

// FieldYear returns the C field 'year'.
func (recv *Date) FieldYear() int32 {
	argValue := gi.FieldGet(dateStruct, recv.native, "year")
	value := argValue.Int32()
	return value
}

// SetFieldYear sets the value of the C field 'year'.
func (recv *Date) SetFieldYear(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.FieldSet(dateStruct, recv.native, "year", argValue)
}

// FieldMonth returns the C field 'month'.
func (recv *Date) FieldMonth() int32 {
	argValue := gi.FieldGet(dateStruct, recv.native, "month")
	value := argValue.Int32()
	return value
}

// SetFieldMonth sets the value of the C field 'month'.
func (recv *Date) SetFieldMonth(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.FieldSet(dateStruct, recv.native, "month", argValue)
}

// FieldDay returns the C field 'day'.
func (recv *Date) FieldDay() int32 {
	argValue := gi.FieldGet(dateStruct, recv.native, "day")
	value := argValue.Int32()
	return value
}

// SetFieldDay sets the value of the C field 'day'.
func (recv *Date) SetFieldDay(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.FieldSet(dateStruct, recv.native, "day", argValue)
}

// FieldHour returns the C field 'hour'.
func (recv *Date) FieldHour() int32 {
	argValue := gi.FieldGet(dateStruct, recv.native, "hour")
	value := argValue.Int32()
	return value
}

// SetFieldHour sets the value of the C field 'hour'.
func (recv *Date) SetFieldHour(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.FieldSet(dateStruct, recv.native, "hour", argValue)
}

// FieldMinute returns the C field 'minute'.
func (recv *Date) FieldMinute() int32 {
	argValue := gi.FieldGet(dateStruct, recv.native, "minute")
	value := argValue.Int32()
	return value
}

// SetFieldMinute sets the value of the C field 'minute'.
func (recv *Date) SetFieldMinute(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.FieldSet(dateStruct, recv.native, "minute", argValue)
}

// FieldSecond returns the C field 'second'.
func (recv *Date) FieldSecond() int32 {
	argValue := gi.FieldGet(dateStruct, recv.native, "second")
	value := argValue.Int32()
	return value
}

// SetFieldSecond sets the value of the C field 'second'.
func (recv *Date) SetFieldSecond(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.FieldSet(dateStruct, recv.native, "second", argValue)
}

// FieldUtc returns the C field 'utc'.
func (recv *Date) FieldUtc() bool {
	argValue := gi.FieldGet(dateStruct, recv.native, "utc")
	value := argValue.Boolean()
	return value
}

// SetFieldUtc sets the value of the C field 'utc'.
func (recv *Date) SetFieldUtc(value bool) {
	var argValue gi.Argument
	argValue.SetBoolean(value)
	gi.FieldSet(dateStruct, recv.native, "utc", argValue)
}

// FieldOffset returns the C field 'offset'.
func (recv *Date) FieldOffset() int32 {
	argValue := gi.FieldGet(dateStruct, recv.native, "offset")
	value := argValue.Int32()
	return value
}

// SetFieldOffset sets the value of the C field 'offset'.
func (recv *Date) SetFieldOffset(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.FieldSet(dateStruct, recv.native, "offset", argValue)
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

	return
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

var dateIsPastFunction *gi.Function
var dateIsPastFunction_Once sync.Once

func dateIsPastFunction_Set() error {
	var err error
	dateIsPastFunction_Once.Do(func() {
		err = dateStruct_Set()
		if err != nil {
			return
		}
		dateIsPastFunction, err = dateStruct.InvokerNew("is_past")
	})
	return err
}

// IsPast is a representation of the C type soup_date_is_past.
func (recv *Date) IsPast() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := dateIsPastFunction_Set()
	if err == nil {
		ret = dateIsPastFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

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
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'is_persistent' : for field getter : missing Type

// UNSUPPORTED : C value 'is_persistent' : for field setter : missing Type

// UNSUPPORTED : C value 'has_valid_policy' : for field getter : missing Type

// UNSUPPORTED : C value 'has_valid_policy' : for field setter : missing Type

// UNSUPPORTED : C value 'changed' : for field getter : missing Type

// UNSUPPORTED : C value 'changed' : for field setter : missing Type

// UNSUPPORTED : C value 'hsts_enforced' : for field getter : missing Type

// UNSUPPORTED : C value 'hsts_enforced' : for field setter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved4' : for field setter : missing Type

// HSTSEnforcerClassStruct creates an uninitialised HSTSEnforcerClass.
func HSTSEnforcerClassStruct() *HSTSEnforcerClass {
	err := hSTSEnforcerClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &HSTSEnforcerClass{native: hSTSEnforcerClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeHSTSEnforcerClass)
	return structGo
}
func finalizeHSTSEnforcerClass(obj *HSTSEnforcerClass) {
	hSTSEnforcerClassStruct.Free(obj.native)
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
	native uintptr
}

// FieldParentClass returns the C field 'parent_class'.
func (recv *HSTSEnforcerDBClass) FieldParentClass() *HSTSEnforcerClass {
	argValue := gi.FieldGet(hSTSEnforcerDBClassStruct, recv.native, "parent_class")
	value := &HSTSEnforcerClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *HSTSEnforcerDBClass) SetFieldParentClass(value *HSTSEnforcerClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(hSTSEnforcerDBClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value '_libsoup_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved4' : for field setter : missing Type

// HSTSEnforcerDBClassStruct creates an uninitialised HSTSEnforcerDBClass.
func HSTSEnforcerDBClassStruct() *HSTSEnforcerDBClass {
	err := hSTSEnforcerDBClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &HSTSEnforcerDBClass{native: hSTSEnforcerDBClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeHSTSEnforcerDBClass)
	return structGo
}
func finalizeHSTSEnforcerDBClass(obj *HSTSEnforcerDBClass) {
	hSTSEnforcerDBClassStruct.Free(obj.native)
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

// HSTSEnforcerDBPrivateStruct creates an uninitialised HSTSEnforcerDBPrivate.
func HSTSEnforcerDBPrivateStruct() *HSTSEnforcerDBPrivate {
	err := hSTSEnforcerDBPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &HSTSEnforcerDBPrivate{native: hSTSEnforcerDBPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeHSTSEnforcerDBPrivate)
	return structGo
}
func finalizeHSTSEnforcerDBPrivate(obj *HSTSEnforcerDBPrivate) {
	hSTSEnforcerDBPrivateStruct.Free(obj.native)
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

// HSTSEnforcerPrivateStruct creates an uninitialised HSTSEnforcerPrivate.
func HSTSEnforcerPrivateStruct() *HSTSEnforcerPrivate {
	err := hSTSEnforcerPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &HSTSEnforcerPrivate{native: hSTSEnforcerPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeHSTSEnforcerPrivate)
	return structGo
}
func finalizeHSTSEnforcerPrivate(obj *HSTSEnforcerPrivate) {
	hSTSEnforcerPrivateStruct.Free(obj.native)
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
	native uintptr
}

// FieldDomain returns the C field 'domain'.
func (recv *HSTSPolicy) FieldDomain() string {
	argValue := gi.FieldGet(hSTSPolicyStruct, recv.native, "domain")
	value := argValue.String(false)
	return value
}

// SetFieldDomain sets the value of the C field 'domain'.
func (recv *HSTSPolicy) SetFieldDomain(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.FieldSet(hSTSPolicyStruct, recv.native, "domain", argValue)
}

// FieldMaxAge returns the C field 'max_age'.
func (recv *HSTSPolicy) FieldMaxAge() uint64 {
	argValue := gi.FieldGet(hSTSPolicyStruct, recv.native, "max_age")
	value := argValue.Uint64()
	return value
}

// SetFieldMaxAge sets the value of the C field 'max_age'.
func (recv *HSTSPolicy) SetFieldMaxAge(value uint64) {
	var argValue gi.Argument
	argValue.SetUint64(value)
	gi.FieldSet(hSTSPolicyStruct, recv.native, "max_age", argValue)
}

// FieldExpires returns the C field 'expires'.
func (recv *HSTSPolicy) FieldExpires() *Date {
	argValue := gi.FieldGet(hSTSPolicyStruct, recv.native, "expires")
	value := &Date{native: argValue.Pointer()}
	return value
}

// SetFieldExpires sets the value of the C field 'expires'.
func (recv *HSTSPolicy) SetFieldExpires(value *Date) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(hSTSPolicyStruct, recv.native, "expires", argValue)
}

// FieldIncludeSubdomains returns the C field 'include_subdomains'.
func (recv *HSTSPolicy) FieldIncludeSubdomains() bool {
	argValue := gi.FieldGet(hSTSPolicyStruct, recv.native, "include_subdomains")
	value := argValue.Boolean()
	return value
}

// SetFieldIncludeSubdomains sets the value of the C field 'include_subdomains'.
func (recv *HSTSPolicy) SetFieldIncludeSubdomains(value bool) {
	var argValue gi.Argument
	argValue.SetBoolean(value)
	gi.FieldSet(hSTSPolicyStruct, recv.native, "include_subdomains", argValue)
}

var hSTSPolicyNewFunction *gi.Function
var hSTSPolicyNewFunction_Once sync.Once

func hSTSPolicyNewFunction_Set() error {
	var err error
	hSTSPolicyNewFunction_Once.Do(func() {
		err = hSTSPolicyStruct_Set()
		if err != nil {
			return
		}
		hSTSPolicyNewFunction, err = hSTSPolicyStruct.InvokerNew("new")
	})
	return err
}

// HSTSPolicyNew is a representation of the C type soup_hsts_policy_new.
func HSTSPolicyNew(domain string, maxAge uint64, includeSubdomains bool) *HSTSPolicy {
	var inArgs [3]gi.Argument
	inArgs[0].SetString(domain)
	inArgs[1].SetUint64(maxAge)
	inArgs[2].SetBoolean(includeSubdomains)

	var ret gi.Argument

	err := hSTSPolicyNewFunction_Set()
	if err == nil {
		ret = hSTSPolicyNewFunction.Invoke(inArgs[:], nil)
	}

	retGo := &HSTSPolicy{native: ret.Pointer()}

	return retGo
}

var hSTSPolicyNewFromResponseFunction *gi.Function
var hSTSPolicyNewFromResponseFunction_Once sync.Once

func hSTSPolicyNewFromResponseFunction_Set() error {
	var err error
	hSTSPolicyNewFromResponseFunction_Once.Do(func() {
		err = hSTSPolicyStruct_Set()
		if err != nil {
			return
		}
		hSTSPolicyNewFromResponseFunction, err = hSTSPolicyStruct.InvokerNew("new_from_response")
	})
	return err
}

// HSTSPolicyNewFromResponse is a representation of the C type soup_hsts_policy_new_from_response.
func HSTSPolicyNewFromResponse(msg *Message) *HSTSPolicy {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(msg.native)

	var ret gi.Argument

	err := hSTSPolicyNewFromResponseFunction_Set()
	if err == nil {
		ret = hSTSPolicyNewFromResponseFunction.Invoke(inArgs[:], nil)
	}

	retGo := &HSTSPolicy{native: ret.Pointer()}

	return retGo
}

var hSTSPolicyNewFullFunction *gi.Function
var hSTSPolicyNewFullFunction_Once sync.Once

func hSTSPolicyNewFullFunction_Set() error {
	var err error
	hSTSPolicyNewFullFunction_Once.Do(func() {
		err = hSTSPolicyStruct_Set()
		if err != nil {
			return
		}
		hSTSPolicyNewFullFunction, err = hSTSPolicyStruct.InvokerNew("new_full")
	})
	return err
}

// HSTSPolicyNewFull is a representation of the C type soup_hsts_policy_new_full.
func HSTSPolicyNewFull(domain string, maxAge uint64, expires *Date, includeSubdomains bool) *HSTSPolicy {
	var inArgs [4]gi.Argument
	inArgs[0].SetString(domain)
	inArgs[1].SetUint64(maxAge)
	inArgs[2].SetPointer(expires.native)
	inArgs[3].SetBoolean(includeSubdomains)

	var ret gi.Argument

	err := hSTSPolicyNewFullFunction_Set()
	if err == nil {
		ret = hSTSPolicyNewFullFunction.Invoke(inArgs[:], nil)
	}

	retGo := &HSTSPolicy{native: ret.Pointer()}

	return retGo
}

var hSTSPolicyNewSessionPolicyFunction *gi.Function
var hSTSPolicyNewSessionPolicyFunction_Once sync.Once

func hSTSPolicyNewSessionPolicyFunction_Set() error {
	var err error
	hSTSPolicyNewSessionPolicyFunction_Once.Do(func() {
		err = hSTSPolicyStruct_Set()
		if err != nil {
			return
		}
		hSTSPolicyNewSessionPolicyFunction, err = hSTSPolicyStruct.InvokerNew("new_session_policy")
	})
	return err
}

// HSTSPolicyNewSessionPolicy is a representation of the C type soup_hsts_policy_new_session_policy.
func HSTSPolicyNewSessionPolicy(domain string, includeSubdomains bool) *HSTSPolicy {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(domain)
	inArgs[1].SetBoolean(includeSubdomains)

	var ret gi.Argument

	err := hSTSPolicyNewSessionPolicyFunction_Set()
	if err == nil {
		ret = hSTSPolicyNewSessionPolicyFunction.Invoke(inArgs[:], nil)
	}

	retGo := &HSTSPolicy{native: ret.Pointer()}

	return retGo
}

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

var hSTSPolicyEqualFunction *gi.Function
var hSTSPolicyEqualFunction_Once sync.Once

func hSTSPolicyEqualFunction_Set() error {
	var err error
	hSTSPolicyEqualFunction_Once.Do(func() {
		err = hSTSPolicyStruct_Set()
		if err != nil {
			return
		}
		hSTSPolicyEqualFunction, err = hSTSPolicyStruct.InvokerNew("equal")
	})
	return err
}

// Equal is a representation of the C type soup_hsts_policy_equal.
func (recv *HSTSPolicy) Equal(policy2 *HSTSPolicy) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(policy2.native)

	var ret gi.Argument

	err := hSTSPolicyEqualFunction_Set()
	if err == nil {
		ret = hSTSPolicyEqualFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

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

	return
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

var hSTSPolicyIncludesSubdomainsFunction *gi.Function
var hSTSPolicyIncludesSubdomainsFunction_Once sync.Once

func hSTSPolicyIncludesSubdomainsFunction_Set() error {
	var err error
	hSTSPolicyIncludesSubdomainsFunction_Once.Do(func() {
		err = hSTSPolicyStruct_Set()
		if err != nil {
			return
		}
		hSTSPolicyIncludesSubdomainsFunction, err = hSTSPolicyStruct.InvokerNew("includes_subdomains")
	})
	return err
}

// IncludesSubdomains is a representation of the C type soup_hsts_policy_includes_subdomains.
func (recv *HSTSPolicy) IncludesSubdomains() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := hSTSPolicyIncludesSubdomainsFunction_Set()
	if err == nil {
		ret = hSTSPolicyIncludesSubdomainsFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var hSTSPolicyIsExpiredFunction *gi.Function
var hSTSPolicyIsExpiredFunction_Once sync.Once

func hSTSPolicyIsExpiredFunction_Set() error {
	var err error
	hSTSPolicyIsExpiredFunction_Once.Do(func() {
		err = hSTSPolicyStruct_Set()
		if err != nil {
			return
		}
		hSTSPolicyIsExpiredFunction, err = hSTSPolicyStruct.InvokerNew("is_expired")
	})
	return err
}

// IsExpired is a representation of the C type soup_hsts_policy_is_expired.
func (recv *HSTSPolicy) IsExpired() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := hSTSPolicyIsExpiredFunction_Set()
	if err == nil {
		ret = hSTSPolicyIsExpiredFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var hSTSPolicyIsSessionPolicyFunction *gi.Function
var hSTSPolicyIsSessionPolicyFunction_Once sync.Once

func hSTSPolicyIsSessionPolicyFunction_Set() error {
	var err error
	hSTSPolicyIsSessionPolicyFunction_Once.Do(func() {
		err = hSTSPolicyStruct_Set()
		if err != nil {
			return
		}
		hSTSPolicyIsSessionPolicyFunction, err = hSTSPolicyStruct.InvokerNew("is_session_policy")
	})
	return err
}

// IsSessionPolicy is a representation of the C type soup_hsts_policy_is_session_policy.
func (recv *HSTSPolicy) IsSessionPolicy() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := hSTSPolicyIsSessionPolicyFunction_Set()
	if err == nil {
		ret = hSTSPolicyIsSessionPolicyFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

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
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value '_libsoup_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved4' : for field setter : missing Type

// LoggerClassStruct creates an uninitialised LoggerClass.
func LoggerClassStruct() *LoggerClass {
	err := loggerClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &LoggerClass{native: loggerClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeLoggerClass)
	return structGo
}
func finalizeLoggerClass(obj *LoggerClass) {
	loggerClassStruct.Free(obj.native)
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
}

// FieldData returns the C field 'data'.
func (recv *MessageBody) FieldData() string {
	argValue := gi.FieldGet(messageBodyStruct, recv.native, "data")
	value := argValue.String(false)
	return value
}

// SetFieldData sets the value of the C field 'data'.
func (recv *MessageBody) SetFieldData(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.FieldSet(messageBodyStruct, recv.native, "data", argValue)
}

// FieldLength returns the C field 'length'.
func (recv *MessageBody) FieldLength() int64 {
	argValue := gi.FieldGet(messageBodyStruct, recv.native, "length")
	value := argValue.Int64()
	return value
}

// SetFieldLength sets the value of the C field 'length'.
func (recv *MessageBody) SetFieldLength(value int64) {
	var argValue gi.Argument
	argValue.SetInt64(value)
	gi.FieldSet(messageBodyStruct, recv.native, "length", argValue)
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

var messageBodyAppendBufferFunction *gi.Function
var messageBodyAppendBufferFunction_Once sync.Once

func messageBodyAppendBufferFunction_Set() error {
	var err error
	messageBodyAppendBufferFunction_Once.Do(func() {
		err = messageBodyStruct_Set()
		if err != nil {
			return
		}
		messageBodyAppendBufferFunction, err = messageBodyStruct.InvokerNew("append_buffer")
	})
	return err
}

// AppendBuffer is a representation of the C type soup_message_body_append_buffer.
func (recv *MessageBody) AppendBuffer(buffer *Buffer) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(buffer.native)

	err := messageBodyAppendBufferFunction_Set()
	if err == nil {
		messageBodyAppendBufferFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'soup_message_body_append_take' : parameter 'data' of type 'nil' not supported

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

	return
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

	return
}

var messageBodyGetAccumulateFunction *gi.Function
var messageBodyGetAccumulateFunction_Once sync.Once

func messageBodyGetAccumulateFunction_Set() error {
	var err error
	messageBodyGetAccumulateFunction_Once.Do(func() {
		err = messageBodyStruct_Set()
		if err != nil {
			return
		}
		messageBodyGetAccumulateFunction, err = messageBodyStruct.InvokerNew("get_accumulate")
	})
	return err
}

// GetAccumulate is a representation of the C type soup_message_body_get_accumulate.
func (recv *MessageBody) GetAccumulate() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := messageBodyGetAccumulateFunction_Set()
	if err == nil {
		ret = messageBodyGetAccumulateFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

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

var messageBodyGotChunkFunction *gi.Function
var messageBodyGotChunkFunction_Once sync.Once

func messageBodyGotChunkFunction_Set() error {
	var err error
	messageBodyGotChunkFunction_Once.Do(func() {
		err = messageBodyStruct_Set()
		if err != nil {
			return
		}
		messageBodyGotChunkFunction, err = messageBodyStruct.InvokerNew("got_chunk")
	})
	return err
}

// GotChunk is a representation of the C type soup_message_body_got_chunk.
func (recv *MessageBody) GotChunk(chunk *Buffer) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(chunk.native)

	err := messageBodyGotChunkFunction_Set()
	if err == nil {
		messageBodyGotChunkFunction.Invoke(inArgs[:], nil)
	}

	return
}

var messageBodySetAccumulateFunction *gi.Function
var messageBodySetAccumulateFunction_Once sync.Once

func messageBodySetAccumulateFunction_Set() error {
	var err error
	messageBodySetAccumulateFunction_Once.Do(func() {
		err = messageBodyStruct_Set()
		if err != nil {
			return
		}
		messageBodySetAccumulateFunction, err = messageBodyStruct.InvokerNew("set_accumulate")
	})
	return err
}

// SetAccumulate is a representation of the C type soup_message_body_set_accumulate.
func (recv *MessageBody) SetAccumulate(accumulate bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetBoolean(accumulate)

	err := messageBodySetAccumulateFunction_Set()
	if err == nil {
		messageBodySetAccumulateFunction.Invoke(inArgs[:], nil)
	}

	return
}

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

	return
}

var messageBodyWroteChunkFunction *gi.Function
var messageBodyWroteChunkFunction_Once sync.Once

func messageBodyWroteChunkFunction_Set() error {
	var err error
	messageBodyWroteChunkFunction_Once.Do(func() {
		err = messageBodyStruct_Set()
		if err != nil {
			return
		}
		messageBodyWroteChunkFunction, err = messageBodyStruct.InvokerNew("wrote_chunk")
	})
	return err
}

// WroteChunk is a representation of the C type soup_message_body_wrote_chunk.
func (recv *MessageBody) WroteChunk(chunk *Buffer) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(chunk.native)

	err := messageBodyWroteChunkFunction_Set()
	if err == nil {
		messageBodyWroteChunkFunction.Invoke(inArgs[:], nil)
	}

	return
}

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
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'wrote_informational' : for field getter : missing Type

// UNSUPPORTED : C value 'wrote_informational' : for field setter : missing Type

// UNSUPPORTED : C value 'wrote_headers' : for field getter : missing Type

// UNSUPPORTED : C value 'wrote_headers' : for field setter : missing Type

// UNSUPPORTED : C value 'wrote_chunk' : for field getter : missing Type

// UNSUPPORTED : C value 'wrote_chunk' : for field setter : missing Type

// UNSUPPORTED : C value 'wrote_body' : for field getter : missing Type

// UNSUPPORTED : C value 'wrote_body' : for field setter : missing Type

// UNSUPPORTED : C value 'got_informational' : for field getter : missing Type

// UNSUPPORTED : C value 'got_informational' : for field setter : missing Type

// UNSUPPORTED : C value 'got_headers' : for field getter : missing Type

// UNSUPPORTED : C value 'got_headers' : for field setter : missing Type

// UNSUPPORTED : C value 'got_chunk' : for field getter : missing Type

// UNSUPPORTED : C value 'got_chunk' : for field setter : missing Type

// UNSUPPORTED : C value 'got_body' : for field getter : missing Type

// UNSUPPORTED : C value 'got_body' : for field setter : missing Type

// UNSUPPORTED : C value 'restarted' : for field getter : missing Type

// UNSUPPORTED : C value 'restarted' : for field setter : missing Type

// UNSUPPORTED : C value 'finished' : for field getter : missing Type

// UNSUPPORTED : C value 'finished' : for field setter : missing Type

// UNSUPPORTED : C value 'starting' : for field getter : missing Type

// UNSUPPORTED : C value 'starting' : for field setter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved3' : for field setter : missing Type

// MessageClassStruct creates an uninitialised MessageClass.
func MessageClassStruct() *MessageClass {
	err := messageClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &MessageClass{native: messageClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeMessageClass)
	return structGo
}
func finalizeMessageClass(obj *MessageClass) {
	messageClassStruct.Free(obj.native)
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

	return
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

	return
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

	return
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

	return
}

var messageHeadersFreeRangesFunction *gi.Function
var messageHeadersFreeRangesFunction_Once sync.Once

func messageHeadersFreeRangesFunction_Set() error {
	var err error
	messageHeadersFreeRangesFunction_Once.Do(func() {
		err = messageHeadersStruct_Set()
		if err != nil {
			return
		}
		messageHeadersFreeRangesFunction, err = messageHeadersStruct.InvokerNew("free_ranges")
	})
	return err
}

// FreeRanges is a representation of the C type soup_message_headers_free_ranges.
func (recv *MessageHeaders) FreeRanges(ranges *Range) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(ranges.native)

	err := messageHeadersFreeRangesFunction_Set()
	if err == nil {
		messageHeadersFreeRangesFunction.Invoke(inArgs[:], nil)
	}

	return
}

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

var messageHeadersGetContentRangeFunction *gi.Function
var messageHeadersGetContentRangeFunction_Once sync.Once

func messageHeadersGetContentRangeFunction_Set() error {
	var err error
	messageHeadersGetContentRangeFunction_Once.Do(func() {
		err = messageHeadersStruct_Set()
		if err != nil {
			return
		}
		messageHeadersGetContentRangeFunction, err = messageHeadersStruct.InvokerNew("get_content_range")
	})
	return err
}

// GetContentRange is a representation of the C type soup_message_headers_get_content_range.
func (recv *MessageHeaders) GetContentRange() (bool, int64, int64, int64) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var outArgs [3]gi.Argument
	var ret gi.Argument

	err := messageHeadersGetContentRangeFunction_Set()
	if err == nil {
		ret = messageHeadersGetContentRangeFunction.Invoke(inArgs[:], outArgs[:])
	}

	retGo := ret.Boolean()
	out0 := outArgs[0].Int64()
	out1 := outArgs[1].Int64()
	out2 := outArgs[2].Int64()

	return retGo, out0, out1, out2
}

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

// UNSUPPORTED : C value 'soup_message_headers_get_ranges' : parameter 'ranges' of type 'nil' not supported

var messageHeadersHeaderContainsFunction *gi.Function
var messageHeadersHeaderContainsFunction_Once sync.Once

func messageHeadersHeaderContainsFunction_Set() error {
	var err error
	messageHeadersHeaderContainsFunction_Once.Do(func() {
		err = messageHeadersStruct_Set()
		if err != nil {
			return
		}
		messageHeadersHeaderContainsFunction, err = messageHeadersStruct.InvokerNew("header_contains")
	})
	return err
}

// HeaderContains is a representation of the C type soup_message_headers_header_contains.
func (recv *MessageHeaders) HeaderContains(name string, token string) bool {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(name)
	inArgs[2].SetString(token)

	var ret gi.Argument

	err := messageHeadersHeaderContainsFunction_Set()
	if err == nil {
		ret = messageHeadersHeaderContainsFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var messageHeadersHeaderEqualsFunction *gi.Function
var messageHeadersHeaderEqualsFunction_Once sync.Once

func messageHeadersHeaderEqualsFunction_Set() error {
	var err error
	messageHeadersHeaderEqualsFunction_Once.Do(func() {
		err = messageHeadersStruct_Set()
		if err != nil {
			return
		}
		messageHeadersHeaderEqualsFunction, err = messageHeadersStruct.InvokerNew("header_equals")
	})
	return err
}

// HeaderEquals is a representation of the C type soup_message_headers_header_equals.
func (recv *MessageHeaders) HeaderEquals(name string, value string) bool {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(name)
	inArgs[2].SetString(value)

	var ret gi.Argument

	err := messageHeadersHeaderEqualsFunction_Set()
	if err == nil {
		ret = messageHeadersHeaderEqualsFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

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

	return
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

	return
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

	return
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

	return
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

	return
}

var messageHeadersSetRangesFunction *gi.Function
var messageHeadersSetRangesFunction_Once sync.Once

func messageHeadersSetRangesFunction_Set() error {
	var err error
	messageHeadersSetRangesFunction_Once.Do(func() {
		err = messageHeadersStruct_Set()
		if err != nil {
			return
		}
		messageHeadersSetRangesFunction, err = messageHeadersStruct.InvokerNew("set_ranges")
	})
	return err
}

// SetRanges is a representation of the C type soup_message_headers_set_ranges.
func (recv *MessageHeaders) SetRanges(ranges *Range, length int32) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(ranges.native)
	inArgs[2].SetInt32(length)

	err := messageHeadersSetRangesFunction_Set()
	if err == nil {
		messageHeadersSetRangesFunction.Invoke(inArgs[:], nil)
	}

	return
}

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

var messageHeadersIterNextFunction *gi.Function
var messageHeadersIterNextFunction_Once sync.Once

func messageHeadersIterNextFunction_Set() error {
	var err error
	messageHeadersIterNextFunction_Once.Do(func() {
		err = messageHeadersIterStruct_Set()
		if err != nil {
			return
		}
		messageHeadersIterNextFunction, err = messageHeadersIterStruct.InvokerNew("next")
	})
	return err
}

// Next is a representation of the C type soup_message_headers_iter_next.
func (recv *MessageHeadersIter) Next() (bool, string, string) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var outArgs [2]gi.Argument
	var ret gi.Argument

	err := messageHeadersIterNextFunction_Set()
	if err == nil {
		ret = messageHeadersIterNextFunction.Invoke(inArgs[:], outArgs[:])
	}

	retGo := ret.Boolean()
	out0 := outArgs[0].String(false)
	out1 := outArgs[1].String(false)

	return retGo, out0, out1
}

// MessageHeadersIterStruct creates an uninitialised MessageHeadersIter.
func MessageHeadersIterStruct() *MessageHeadersIter {
	err := messageHeadersIterStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &MessageHeadersIter{native: messageHeadersIterStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeMessageHeadersIter)
	return structGo
}
func finalizeMessageHeadersIter(obj *MessageHeadersIter) {
	messageHeadersIterStruct.Free(obj.native)
}

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

// MessageQueueStruct creates an uninitialised MessageQueue.
func MessageQueueStruct() *MessageQueue {
	err := messageQueueStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &MessageQueue{native: messageQueueStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeMessageQueue)
	return structGo
}
func finalizeMessageQueue(obj *MessageQueue) {
	messageQueueStruct.Free(obj.native)
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

// MessageQueueItemStruct creates an uninitialised MessageQueueItem.
func MessageQueueItemStruct() *MessageQueueItem {
	err := messageQueueItemStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &MessageQueueItem{native: messageQueueItemStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeMessageQueueItem)
	return structGo
}
func finalizeMessageQueueItem(obj *MessageQueueItem) {
	messageQueueItemStruct.Free(obj.native)
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

var multipartNewFromMessageFunction *gi.Function
var multipartNewFromMessageFunction_Once sync.Once

func multipartNewFromMessageFunction_Set() error {
	var err error
	multipartNewFromMessageFunction_Once.Do(func() {
		err = multipartStruct_Set()
		if err != nil {
			return
		}
		multipartNewFromMessageFunction, err = multipartStruct.InvokerNew("new_from_message")
	})
	return err
}

// MultipartNewFromMessage is a representation of the C type soup_multipart_new_from_message.
func MultipartNewFromMessage(headers *MessageHeaders, body *MessageBody) *Multipart {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(headers.native)
	inArgs[1].SetPointer(body.native)

	var ret gi.Argument

	err := multipartNewFromMessageFunction_Set()
	if err == nil {
		ret = multipartNewFromMessageFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Multipart{native: ret.Pointer()}

	return retGo
}

var multipartAppendFormFileFunction *gi.Function
var multipartAppendFormFileFunction_Once sync.Once

func multipartAppendFormFileFunction_Set() error {
	var err error
	multipartAppendFormFileFunction_Once.Do(func() {
		err = multipartStruct_Set()
		if err != nil {
			return
		}
		multipartAppendFormFileFunction, err = multipartStruct.InvokerNew("append_form_file")
	})
	return err
}

// AppendFormFile is a representation of the C type soup_multipart_append_form_file.
func (recv *Multipart) AppendFormFile(controlName string, filename string, contentType string, body *Buffer) {
	var inArgs [5]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(controlName)
	inArgs[2].SetString(filename)
	inArgs[3].SetString(contentType)
	inArgs[4].SetPointer(body.native)

	err := multipartAppendFormFileFunction_Set()
	if err == nil {
		multipartAppendFormFileFunction.Invoke(inArgs[:], nil)
	}

	return
}

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

	return
}

var multipartAppendPartFunction *gi.Function
var multipartAppendPartFunction_Once sync.Once

func multipartAppendPartFunction_Set() error {
	var err error
	multipartAppendPartFunction_Once.Do(func() {
		err = multipartStruct_Set()
		if err != nil {
			return
		}
		multipartAppendPartFunction, err = multipartStruct.InvokerNew("append_part")
	})
	return err
}

// AppendPart is a representation of the C type soup_multipart_append_part.
func (recv *Multipart) AppendPart(headers *MessageHeaders, body *Buffer) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(headers.native)
	inArgs[2].SetPointer(body.native)

	err := multipartAppendPartFunction_Set()
	if err == nil {
		multipartAppendPartFunction.Invoke(inArgs[:], nil)
	}

	return
}

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

	return
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

var multipartGetPartFunction *gi.Function
var multipartGetPartFunction_Once sync.Once

func multipartGetPartFunction_Set() error {
	var err error
	multipartGetPartFunction_Once.Do(func() {
		err = multipartStruct_Set()
		if err != nil {
			return
		}
		multipartGetPartFunction, err = multipartStruct.InvokerNew("get_part")
	})
	return err
}

// GetPart is a representation of the C type soup_multipart_get_part.
func (recv *Multipart) GetPart(part int32) (bool, *MessageHeaders, *Buffer) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(part)

	var outArgs [2]gi.Argument
	var ret gi.Argument

	err := multipartGetPartFunction_Set()
	if err == nil {
		ret = multipartGetPartFunction.Invoke(inArgs[:], outArgs[:])
	}

	retGo := ret.Boolean()
	out0 := &MessageHeaders{native: outArgs[0].Pointer()}
	out1 := &Buffer{native: outArgs[1].Pointer()}

	return retGo, out0, out1
}

var multipartToMessageFunction *gi.Function
var multipartToMessageFunction_Once sync.Once

func multipartToMessageFunction_Set() error {
	var err error
	multipartToMessageFunction_Once.Do(func() {
		err = multipartStruct_Set()
		if err != nil {
			return
		}
		multipartToMessageFunction, err = multipartStruct.InvokerNew("to_message")
	})
	return err
}

// ToMessage is a representation of the C type soup_multipart_to_message.
func (recv *Multipart) ToMessage(destHeaders *MessageHeaders, destBody *MessageBody) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(destHeaders.native)
	inArgs[2].SetPointer(destBody.native)

	err := multipartToMessageFunction_Set()
	if err == nil {
		multipartToMessageFunction.Invoke(inArgs[:], nil)
	}

	return
}

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
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'Gio.FilterInputStreamClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'Gio.FilterInputStreamClass'

// MultipartInputStreamClassStruct creates an uninitialised MultipartInputStreamClass.
func MultipartInputStreamClassStruct() *MultipartInputStreamClass {
	err := multipartInputStreamClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &MultipartInputStreamClass{native: multipartInputStreamClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeMultipartInputStreamClass)
	return structGo
}
func finalizeMultipartInputStreamClass(obj *MultipartInputStreamClass) {
	multipartInputStreamClassStruct.Free(obj.native)
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

// MultipartInputStreamPrivateStruct creates an uninitialised MultipartInputStreamPrivate.
func MultipartInputStreamPrivateStruct() *MultipartInputStreamPrivate {
	err := multipartInputStreamPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &MultipartInputStreamPrivate{native: multipartInputStreamPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeMultipartInputStreamPrivate)
	return structGo
}
func finalizeMultipartInputStreamPrivate(obj *MultipartInputStreamPrivate) {
	multipartInputStreamPrivateStruct.Free(obj.native)
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
}

// UNSUPPORTED : C value 'base' : for field getter : no Go type for 'GObject.TypeInterface'

// UNSUPPORTED : C value 'base' : for field setter : no Go type for 'GObject.TypeInterface'

// UNSUPPORTED : C value 'get_passwords_async' : for field getter : missing Type

// UNSUPPORTED : C value 'get_passwords_async' : for field setter : missing Type

// UNSUPPORTED : C value 'get_passwords_sync' : for field getter : missing Type

// UNSUPPORTED : C value 'get_passwords_sync' : for field setter : missing Type

// PasswordManagerInterfaceStruct creates an uninitialised PasswordManagerInterface.
func PasswordManagerInterfaceStruct() *PasswordManagerInterface {
	err := passwordManagerInterfaceStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &PasswordManagerInterface{native: passwordManagerInterfaceStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizePasswordManagerInterface)
	return structGo
}
func finalizePasswordManagerInterface(obj *PasswordManagerInterface) {
	passwordManagerInterfaceStruct.Free(obj.native)
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
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// ProxyResolverDefaultClassStruct creates an uninitialised ProxyResolverDefaultClass.
func ProxyResolverDefaultClassStruct() *ProxyResolverDefaultClass {
	err := proxyResolverDefaultClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ProxyResolverDefaultClass{native: proxyResolverDefaultClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeProxyResolverDefaultClass)
	return structGo
}
func finalizeProxyResolverDefaultClass(obj *ProxyResolverDefaultClass) {
	proxyResolverDefaultClassStruct.Free(obj.native)
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
}

// UNSUPPORTED : C value 'base' : for field getter : no Go type for 'GObject.TypeInterface'

// UNSUPPORTED : C value 'base' : for field setter : no Go type for 'GObject.TypeInterface'

// UNSUPPORTED : C value 'get_proxy_async' : for field getter : missing Type

// UNSUPPORTED : C value 'get_proxy_async' : for field setter : missing Type

// UNSUPPORTED : C value 'get_proxy_sync' : for field getter : missing Type

// UNSUPPORTED : C value 'get_proxy_sync' : for field setter : missing Type

// ProxyResolverInterfaceStruct creates an uninitialised ProxyResolverInterface.
func ProxyResolverInterfaceStruct() *ProxyResolverInterface {
	err := proxyResolverInterfaceStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ProxyResolverInterface{native: proxyResolverInterfaceStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeProxyResolverInterface)
	return structGo
}
func finalizeProxyResolverInterface(obj *ProxyResolverInterface) {
	proxyResolverInterfaceStruct.Free(obj.native)
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
}

// UNSUPPORTED : C value 'base' : for field getter : no Go type for 'GObject.TypeInterface'

// UNSUPPORTED : C value 'base' : for field setter : no Go type for 'GObject.TypeInterface'

// UNSUPPORTED : C value 'get_proxy_uri_async' : for field getter : missing Type

// UNSUPPORTED : C value 'get_proxy_uri_async' : for field setter : missing Type

// UNSUPPORTED : C value 'get_proxy_uri_sync' : for field getter : missing Type

// UNSUPPORTED : C value 'get_proxy_uri_sync' : for field setter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved4' : for field setter : missing Type

// ProxyURIResolverInterfaceStruct creates an uninitialised ProxyURIResolverInterface.
func ProxyURIResolverInterfaceStruct() *ProxyURIResolverInterface {
	err := proxyURIResolverInterfaceStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ProxyURIResolverInterface{native: proxyURIResolverInterfaceStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeProxyURIResolverInterface)
	return structGo
}
func finalizeProxyURIResolverInterface(obj *ProxyURIResolverInterface) {
	proxyURIResolverInterfaceStruct.Free(obj.native)
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
}

// FieldStart returns the C field 'start'.
func (recv *Range) FieldStart() int64 {
	argValue := gi.FieldGet(rangeStruct, recv.native, "start")
	value := argValue.Int64()
	return value
}

// SetFieldStart sets the value of the C field 'start'.
func (recv *Range) SetFieldStart(value int64) {
	var argValue gi.Argument
	argValue.SetInt64(value)
	gi.FieldSet(rangeStruct, recv.native, "start", argValue)
}

// FieldEnd returns the C field 'end'.
func (recv *Range) FieldEnd() int64 {
	argValue := gi.FieldGet(rangeStruct, recv.native, "end")
	value := argValue.Int64()
	return value
}

// SetFieldEnd sets the value of the C field 'end'.
func (recv *Range) SetFieldEnd(value int64) {
	var argValue gi.Argument
	argValue.SetInt64(value)
	gi.FieldSet(rangeStruct, recv.native, "end", argValue)
}

// RangeStruct creates an uninitialised Range.
func RangeStruct() *Range {
	err := rangeStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &Range{native: rangeStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeRange)
	return structGo
}
func finalizeRange(obj *Range) {
	rangeStruct.Free(obj.native)
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
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'GObject.ObjectClass'

// FieldSchemes returns the C field 'schemes'.
func (recv *RequestClass) FieldSchemes() string {
	argValue := gi.FieldGet(requestClassStruct, recv.native, "schemes")
	value := argValue.String(false)
	return value
}

// SetFieldSchemes sets the value of the C field 'schemes'.
func (recv *RequestClass) SetFieldSchemes(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.FieldSet(requestClassStruct, recv.native, "schemes", argValue)
}

// UNSUPPORTED : C value 'check_uri' : for field getter : missing Type

// UNSUPPORTED : C value 'check_uri' : for field setter : missing Type

// UNSUPPORTED : C value 'send' : for field getter : missing Type

// UNSUPPORTED : C value 'send' : for field setter : missing Type

// UNSUPPORTED : C value 'send_async' : for field getter : missing Type

// UNSUPPORTED : C value 'send_async' : for field setter : missing Type

// UNSUPPORTED : C value 'send_finish' : for field getter : missing Type

// UNSUPPORTED : C value 'send_finish' : for field setter : missing Type

// UNSUPPORTED : C value 'get_content_length' : for field getter : missing Type

// UNSUPPORTED : C value 'get_content_length' : for field setter : missing Type

// UNSUPPORTED : C value 'get_content_type' : for field getter : missing Type

// UNSUPPORTED : C value 'get_content_type' : for field setter : missing Type

// RequestClassStruct creates an uninitialised RequestClass.
func RequestClassStruct() *RequestClass {
	err := requestClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &RequestClass{native: requestClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeRequestClass)
	return structGo
}
func finalizeRequestClass(obj *RequestClass) {
	requestClassStruct.Free(obj.native)
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
}

// FieldParent returns the C field 'parent'.
func (recv *RequestDataClass) FieldParent() *RequestClass {
	argValue := gi.FieldGet(requestDataClassStruct, recv.native, "parent")
	value := &RequestClass{native: argValue.Pointer()}
	return value
}

// SetFieldParent sets the value of the C field 'parent'.
func (recv *RequestDataClass) SetFieldParent(value *RequestClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(requestDataClassStruct, recv.native, "parent", argValue)
}

// RequestDataClassStruct creates an uninitialised RequestDataClass.
func RequestDataClassStruct() *RequestDataClass {
	err := requestDataClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &RequestDataClass{native: requestDataClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeRequestDataClass)
	return structGo
}
func finalizeRequestDataClass(obj *RequestDataClass) {
	requestDataClassStruct.Free(obj.native)
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

// RequestDataPrivateStruct creates an uninitialised RequestDataPrivate.
func RequestDataPrivateStruct() *RequestDataPrivate {
	err := requestDataPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &RequestDataPrivate{native: requestDataPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeRequestDataPrivate)
	return structGo
}
func finalizeRequestDataPrivate(obj *RequestDataPrivate) {
	requestDataPrivateStruct.Free(obj.native)
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
}

// FieldParent returns the C field 'parent'.
func (recv *RequestFileClass) FieldParent() *RequestClass {
	argValue := gi.FieldGet(requestFileClassStruct, recv.native, "parent")
	value := &RequestClass{native: argValue.Pointer()}
	return value
}

// SetFieldParent sets the value of the C field 'parent'.
func (recv *RequestFileClass) SetFieldParent(value *RequestClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(requestFileClassStruct, recv.native, "parent", argValue)
}

// RequestFileClassStruct creates an uninitialised RequestFileClass.
func RequestFileClassStruct() *RequestFileClass {
	err := requestFileClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &RequestFileClass{native: requestFileClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeRequestFileClass)
	return structGo
}
func finalizeRequestFileClass(obj *RequestFileClass) {
	requestFileClassStruct.Free(obj.native)
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

// RequestFilePrivateStruct creates an uninitialised RequestFilePrivate.
func RequestFilePrivateStruct() *RequestFilePrivate {
	err := requestFilePrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &RequestFilePrivate{native: requestFilePrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeRequestFilePrivate)
	return structGo
}
func finalizeRequestFilePrivate(obj *RequestFilePrivate) {
	requestFilePrivateStruct.Free(obj.native)
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
}

// FieldParent returns the C field 'parent'.
func (recv *RequestHTTPClass) FieldParent() *RequestClass {
	argValue := gi.FieldGet(requestHTTPClassStruct, recv.native, "parent")
	value := &RequestClass{native: argValue.Pointer()}
	return value
}

// SetFieldParent sets the value of the C field 'parent'.
func (recv *RequestHTTPClass) SetFieldParent(value *RequestClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(requestHTTPClassStruct, recv.native, "parent", argValue)
}

// RequestHTTPClassStruct creates an uninitialised RequestHTTPClass.
func RequestHTTPClassStruct() *RequestHTTPClass {
	err := requestHTTPClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &RequestHTTPClass{native: requestHTTPClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeRequestHTTPClass)
	return structGo
}
func finalizeRequestHTTPClass(obj *RequestHTTPClass) {
	requestHTTPClassStruct.Free(obj.native)
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

// RequestHTTPPrivateStruct creates an uninitialised RequestHTTPPrivate.
func RequestHTTPPrivateStruct() *RequestHTTPPrivate {
	err := requestHTTPPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &RequestHTTPPrivate{native: requestHTTPPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeRequestHTTPPrivate)
	return structGo
}
func finalizeRequestHTTPPrivate(obj *RequestHTTPPrivate) {
	requestHTTPPrivateStruct.Free(obj.native)
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

// RequestPrivateStruct creates an uninitialised RequestPrivate.
func RequestPrivateStruct() *RequestPrivate {
	err := requestPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &RequestPrivate{native: requestPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeRequestPrivate)
	return structGo
}
func finalizeRequestPrivate(obj *RequestPrivate) {
	requestPrivateStruct.Free(obj.native)
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
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// RequesterClassStruct creates an uninitialised RequesterClass.
func RequesterClassStruct() *RequesterClass {
	err := requesterClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &RequesterClass{native: requesterClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeRequesterClass)
	return structGo
}
func finalizeRequesterClass(obj *RequesterClass) {
	requesterClassStruct.Free(obj.native)
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

// RequesterPrivateStruct creates an uninitialised RequesterPrivate.
func RequesterPrivateStruct() *RequesterPrivate {
	err := requesterPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &RequesterPrivate{native: requesterPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeRequesterPrivate)
	return structGo
}
func finalizeRequesterPrivate(obj *RequesterPrivate) {
	requesterPrivateStruct.Free(obj.native)
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
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'request_started' : for field getter : missing Type

// UNSUPPORTED : C value 'request_started' : for field setter : missing Type

// UNSUPPORTED : C value 'request_read' : for field getter : missing Type

// UNSUPPORTED : C value 'request_read' : for field setter : missing Type

// UNSUPPORTED : C value 'request_finished' : for field getter : missing Type

// UNSUPPORTED : C value 'request_finished' : for field setter : missing Type

// UNSUPPORTED : C value 'request_aborted' : for field getter : missing Type

// UNSUPPORTED : C value 'request_aborted' : for field setter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved4' : for field setter : missing Type

// ServerClassStruct creates an uninitialised ServerClass.
func ServerClassStruct() *ServerClass {
	err := serverClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ServerClass{native: serverClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeServerClass)
	return structGo
}
func finalizeServerClass(obj *ServerClass) {
	serverClassStruct.Free(obj.native)
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
	native uintptr
}

// FieldParentClass returns the C field 'parent_class'.
func (recv *SessionAsyncClass) FieldParentClass() *SessionClass {
	argValue := gi.FieldGet(sessionAsyncClassStruct, recv.native, "parent_class")
	value := &SessionClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *SessionAsyncClass) SetFieldParentClass(value *SessionClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(sessionAsyncClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value '_libsoup_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved4' : for field setter : missing Type

// SessionAsyncClassStruct creates an uninitialised SessionAsyncClass.
func SessionAsyncClassStruct() *SessionAsyncClass {
	err := sessionAsyncClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &SessionAsyncClass{native: sessionAsyncClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeSessionAsyncClass)
	return structGo
}
func finalizeSessionAsyncClass(obj *SessionAsyncClass) {
	sessionAsyncClassStruct.Free(obj.native)
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
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'request_started' : for field getter : missing Type

// UNSUPPORTED : C value 'request_started' : for field setter : missing Type

// UNSUPPORTED : C value 'authenticate' : for field getter : missing Type

// UNSUPPORTED : C value 'authenticate' : for field setter : missing Type

// UNSUPPORTED : C value 'queue_message' : for field getter : missing Type

// UNSUPPORTED : C value 'queue_message' : for field setter : missing Type

// UNSUPPORTED : C value 'requeue_message' : for field getter : missing Type

// UNSUPPORTED : C value 'requeue_message' : for field setter : missing Type

// UNSUPPORTED : C value 'send_message' : for field getter : missing Type

// UNSUPPORTED : C value 'send_message' : for field setter : missing Type

// UNSUPPORTED : C value 'cancel_message' : for field getter : missing Type

// UNSUPPORTED : C value 'cancel_message' : for field setter : missing Type

// UNSUPPORTED : C value 'auth_required' : for field getter : missing Type

// UNSUPPORTED : C value 'auth_required' : for field setter : missing Type

// UNSUPPORTED : C value 'flush_queue' : for field getter : missing Type

// UNSUPPORTED : C value 'flush_queue' : for field setter : missing Type

// UNSUPPORTED : C value 'kick' : for field getter : missing Type

// UNSUPPORTED : C value 'kick' : for field setter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved4' : for field setter : missing Type

// SessionClassStruct creates an uninitialised SessionClass.
func SessionClassStruct() *SessionClass {
	err := sessionClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &SessionClass{native: sessionClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeSessionClass)
	return structGo
}
func finalizeSessionClass(obj *SessionClass) {
	sessionClassStruct.Free(obj.native)
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
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.TypeInterface'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'GObject.TypeInterface'

// UNSUPPORTED : C value 'attach' : for field getter : missing Type

// UNSUPPORTED : C value 'attach' : for field setter : missing Type

// UNSUPPORTED : C value 'detach' : for field getter : missing Type

// UNSUPPORTED : C value 'detach' : for field setter : missing Type

// UNSUPPORTED : C value 'request_queued' : for field getter : missing Type

// UNSUPPORTED : C value 'request_queued' : for field setter : missing Type

// UNSUPPORTED : C value 'request_started' : for field getter : missing Type

// UNSUPPORTED : C value 'request_started' : for field setter : missing Type

// UNSUPPORTED : C value 'request_unqueued' : for field getter : missing Type

// UNSUPPORTED : C value 'request_unqueued' : for field setter : missing Type

// UNSUPPORTED : C value 'add_feature' : for field getter : missing Type

// UNSUPPORTED : C value 'add_feature' : for field setter : missing Type

// UNSUPPORTED : C value 'remove_feature' : for field getter : missing Type

// UNSUPPORTED : C value 'remove_feature' : for field setter : missing Type

// UNSUPPORTED : C value 'has_feature' : for field getter : missing Type

// UNSUPPORTED : C value 'has_feature' : for field setter : missing Type

// SessionFeatureInterfaceStruct creates an uninitialised SessionFeatureInterface.
func SessionFeatureInterfaceStruct() *SessionFeatureInterface {
	err := sessionFeatureInterfaceStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &SessionFeatureInterface{native: sessionFeatureInterfaceStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeSessionFeatureInterface)
	return structGo
}
func finalizeSessionFeatureInterface(obj *SessionFeatureInterface) {
	sessionFeatureInterfaceStruct.Free(obj.native)
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
	native uintptr
}

// FieldParentClass returns the C field 'parent_class'.
func (recv *SessionSyncClass) FieldParentClass() *SessionClass {
	argValue := gi.FieldGet(sessionSyncClassStruct, recv.native, "parent_class")
	value := &SessionClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *SessionSyncClass) SetFieldParentClass(value *SessionClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(sessionSyncClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value '_libsoup_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved4' : for field setter : missing Type

// SessionSyncClassStruct creates an uninitialised SessionSyncClass.
func SessionSyncClassStruct() *SessionSyncClass {
	err := sessionSyncClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &SessionSyncClass{native: sessionSyncClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeSessionSyncClass)
	return structGo
}
func finalizeSessionSyncClass(obj *SessionSyncClass) {
	sessionSyncClassStruct.Free(obj.native)
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
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'readable' : for field getter : missing Type

// UNSUPPORTED : C value 'readable' : for field setter : missing Type

// UNSUPPORTED : C value 'writable' : for field getter : missing Type

// UNSUPPORTED : C value 'writable' : for field setter : missing Type

// UNSUPPORTED : C value 'disconnected' : for field getter : missing Type

// UNSUPPORTED : C value 'disconnected' : for field setter : missing Type

// UNSUPPORTED : C value 'new_connection' : for field getter : missing Type

// UNSUPPORTED : C value 'new_connection' : for field setter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved4' : for field setter : missing Type

// SocketClassStruct creates an uninitialised SocketClass.
func SocketClassStruct() *SocketClass {
	err := socketClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &SocketClass{native: socketClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeSocketClass)
	return structGo
}
func finalizeSocketClass(obj *SocketClass) {
	socketClassStruct.Free(obj.native)
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
	native uintptr
}

// FieldScheme returns the C field 'scheme'.
func (recv *URI) FieldScheme() string {
	argValue := gi.FieldGet(uRIStruct, recv.native, "scheme")
	value := argValue.String(false)
	return value
}

// SetFieldScheme sets the value of the C field 'scheme'.
func (recv *URI) SetFieldScheme(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.FieldSet(uRIStruct, recv.native, "scheme", argValue)
}

// FieldUser returns the C field 'user'.
func (recv *URI) FieldUser() string {
	argValue := gi.FieldGet(uRIStruct, recv.native, "user")
	value := argValue.String(false)
	return value
}

// SetFieldUser sets the value of the C field 'user'.
func (recv *URI) SetFieldUser(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.FieldSet(uRIStruct, recv.native, "user", argValue)
}

// FieldPassword returns the C field 'password'.
func (recv *URI) FieldPassword() string {
	argValue := gi.FieldGet(uRIStruct, recv.native, "password")
	value := argValue.String(false)
	return value
}

// SetFieldPassword sets the value of the C field 'password'.
func (recv *URI) SetFieldPassword(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.FieldSet(uRIStruct, recv.native, "password", argValue)
}

// FieldHost returns the C field 'host'.
func (recv *URI) FieldHost() string {
	argValue := gi.FieldGet(uRIStruct, recv.native, "host")
	value := argValue.String(false)
	return value
}

// SetFieldHost sets the value of the C field 'host'.
func (recv *URI) SetFieldHost(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.FieldSet(uRIStruct, recv.native, "host", argValue)
}

// FieldPort returns the C field 'port'.
func (recv *URI) FieldPort() uint32 {
	argValue := gi.FieldGet(uRIStruct, recv.native, "port")
	value := argValue.Uint32()
	return value
}

// SetFieldPort sets the value of the C field 'port'.
func (recv *URI) SetFieldPort(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.FieldSet(uRIStruct, recv.native, "port", argValue)
}

// FieldPath returns the C field 'path'.
func (recv *URI) FieldPath() string {
	argValue := gi.FieldGet(uRIStruct, recv.native, "path")
	value := argValue.String(false)
	return value
}

// SetFieldPath sets the value of the C field 'path'.
func (recv *URI) SetFieldPath(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.FieldSet(uRIStruct, recv.native, "path", argValue)
}

// FieldQuery returns the C field 'query'.
func (recv *URI) FieldQuery() string {
	argValue := gi.FieldGet(uRIStruct, recv.native, "query")
	value := argValue.String(false)
	return value
}

// SetFieldQuery sets the value of the C field 'query'.
func (recv *URI) SetFieldQuery(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.FieldSet(uRIStruct, recv.native, "query", argValue)
}

// FieldFragment returns the C field 'fragment'.
func (recv *URI) FieldFragment() string {
	argValue := gi.FieldGet(uRIStruct, recv.native, "fragment")
	value := argValue.String(false)
	return value
}

// SetFieldFragment sets the value of the C field 'fragment'.
func (recv *URI) SetFieldFragment(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.FieldSet(uRIStruct, recv.native, "fragment", argValue)
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

var uRINewWithBaseFunction *gi.Function
var uRINewWithBaseFunction_Once sync.Once

func uRINewWithBaseFunction_Set() error {
	var err error
	uRINewWithBaseFunction_Once.Do(func() {
		err = uRIStruct_Set()
		if err != nil {
			return
		}
		uRINewWithBaseFunction, err = uRIStruct.InvokerNew("new_with_base")
	})
	return err
}

// URINewWithBase is a representation of the C type soup_uri_new_with_base.
func URINewWithBase(base *URI, uriString string) *URI {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(base.native)
	inArgs[1].SetString(uriString)

	var ret gi.Argument

	err := uRINewWithBaseFunction_Set()
	if err == nil {
		ret = uRINewWithBaseFunction.Invoke(inArgs[:], nil)
	}

	retGo := &URI{native: ret.Pointer()}

	return retGo
}

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

var uRIEqualFunction *gi.Function
var uRIEqualFunction_Once sync.Once

func uRIEqualFunction_Set() error {
	var err error
	uRIEqualFunction_Once.Do(func() {
		err = uRIStruct_Set()
		if err != nil {
			return
		}
		uRIEqualFunction, err = uRIStruct.InvokerNew("equal")
	})
	return err
}

// Equal is a representation of the C type soup_uri_equal.
func (recv *URI) Equal(uri2 *URI) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(uri2.native)

	var ret gi.Argument

	err := uRIEqualFunction_Set()
	if err == nil {
		ret = uRIEqualFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

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

	return
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

var uRIHostEqualFunction *gi.Function
var uRIHostEqualFunction_Once sync.Once

func uRIHostEqualFunction_Set() error {
	var err error
	uRIHostEqualFunction_Once.Do(func() {
		err = uRIStruct_Set()
		if err != nil {
			return
		}
		uRIHostEqualFunction, err = uRIStruct.InvokerNew("host_equal")
	})
	return err
}

// HostEqual is a representation of the C type soup_uri_host_equal.
func (recv *URI) HostEqual(v2 *URI) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(v2.native)

	var ret gi.Argument

	err := uRIHostEqualFunction_Set()
	if err == nil {
		ret = uRIHostEqualFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

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

	return
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

	return
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

	return
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

	return
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

	return
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

	return
}

// UNSUPPORTED : C value 'soup_uri_set_query_from_fields' : parameter '...' of type 'nil' not supported

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

	return
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

	return
}

var uRIToStringFunction *gi.Function
var uRIToStringFunction_Once sync.Once

func uRIToStringFunction_Set() error {
	var err error
	uRIToStringFunction_Once.Do(func() {
		err = uRIStruct_Set()
		if err != nil {
			return
		}
		uRIToStringFunction, err = uRIStruct.InvokerNew("to_string")
	})
	return err
}

// ToString is a representation of the C type soup_uri_to_string.
func (recv *URI) ToString(justPathAndQuery bool) string {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetBoolean(justPathAndQuery)

	var ret gi.Argument

	err := uRIToStringFunction_Set()
	if err == nil {
		ret = uRIToStringFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

var uRIUsesDefaultPortFunction *gi.Function
var uRIUsesDefaultPortFunction_Once sync.Once

func uRIUsesDefaultPortFunction_Set() error {
	var err error
	uRIUsesDefaultPortFunction_Once.Do(func() {
		err = uRIStruct_Set()
		if err != nil {
			return
		}
		uRIUsesDefaultPortFunction, err = uRIStruct.InvokerNew("uses_default_port")
	})
	return err
}

// UsesDefaultPort is a representation of the C type soup_uri_uses_default_port.
func (recv *URI) UsesDefaultPort() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := uRIUsesDefaultPortFunction_Set()
	if err == nil {
		ret = uRIUsesDefaultPortFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

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
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'message' : for field getter : missing Type

// UNSUPPORTED : C value 'message' : for field setter : missing Type

// UNSUPPORTED : C value 'error' : for field getter : missing Type

// UNSUPPORTED : C value 'error' : for field setter : missing Type

// UNSUPPORTED : C value 'closing' : for field getter : missing Type

// UNSUPPORTED : C value 'closing' : for field setter : missing Type

// UNSUPPORTED : C value 'closed' : for field getter : missing Type

// UNSUPPORTED : C value 'closed' : for field setter : missing Type

// UNSUPPORTED : C value 'pong' : for field getter : missing Type

// UNSUPPORTED : C value 'pong' : for field setter : missing Type

// WebsocketConnectionClassStruct creates an uninitialised WebsocketConnectionClass.
func WebsocketConnectionClassStruct() *WebsocketConnectionClass {
	err := websocketConnectionClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &WebsocketConnectionClass{native: websocketConnectionClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeWebsocketConnectionClass)
	return structGo
}
func finalizeWebsocketConnectionClass(obj *WebsocketConnectionClass) {
	websocketConnectionClassStruct.Free(obj.native)
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

// WebsocketConnectionPrivateStruct creates an uninitialised WebsocketConnectionPrivate.
func WebsocketConnectionPrivateStruct() *WebsocketConnectionPrivate {
	err := websocketConnectionPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &WebsocketConnectionPrivate{native: websocketConnectionPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeWebsocketConnectionPrivate)
	return structGo
}
func finalizeWebsocketConnectionPrivate(obj *WebsocketConnectionPrivate) {
	websocketConnectionPrivateStruct.Free(obj.native)
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
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// FieldName returns the C field 'name'.
func (recv *WebsocketExtensionClass) FieldName() string {
	argValue := gi.FieldGet(websocketExtensionClassStruct, recv.native, "name")
	value := argValue.String(false)
	return value
}

// SetFieldName sets the value of the C field 'name'.
func (recv *WebsocketExtensionClass) SetFieldName(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.FieldSet(websocketExtensionClassStruct, recv.native, "name", argValue)
}

// UNSUPPORTED : C value 'configure' : for field getter : missing Type

// UNSUPPORTED : C value 'configure' : for field setter : missing Type

// UNSUPPORTED : C value 'get_request_params' : for field getter : missing Type

// UNSUPPORTED : C value 'get_request_params' : for field setter : missing Type

// UNSUPPORTED : C value 'get_response_params' : for field getter : missing Type

// UNSUPPORTED : C value 'get_response_params' : for field setter : missing Type

// UNSUPPORTED : C value 'process_outgoing_message' : for field getter : missing Type

// UNSUPPORTED : C value 'process_outgoing_message' : for field setter : missing Type

// UNSUPPORTED : C value 'process_incoming_message' : for field getter : missing Type

// UNSUPPORTED : C value 'process_incoming_message' : for field setter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_libsoup_reserved4' : for field setter : missing Type

// WebsocketExtensionClassStruct creates an uninitialised WebsocketExtensionClass.
func WebsocketExtensionClassStruct() *WebsocketExtensionClass {
	err := websocketExtensionClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &WebsocketExtensionClass{native: websocketExtensionClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeWebsocketExtensionClass)
	return structGo
}
func finalizeWebsocketExtensionClass(obj *WebsocketExtensionClass) {
	websocketExtensionClassStruct.Free(obj.native)
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
	native uintptr
}

// FieldParentClass returns the C field 'parent_class'.
func (recv *WebsocketExtensionDeflateClass) FieldParentClass() *WebsocketExtensionClass {
	argValue := gi.FieldGet(websocketExtensionDeflateClassStruct, recv.native, "parent_class")
	value := &WebsocketExtensionClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *WebsocketExtensionDeflateClass) SetFieldParentClass(value *WebsocketExtensionClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(websocketExtensionDeflateClassStruct, recv.native, "parent_class", argValue)
}

// WebsocketExtensionDeflateClassStruct creates an uninitialised WebsocketExtensionDeflateClass.
func WebsocketExtensionDeflateClassStruct() *WebsocketExtensionDeflateClass {
	err := websocketExtensionDeflateClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &WebsocketExtensionDeflateClass{native: websocketExtensionDeflateClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeWebsocketExtensionDeflateClass)
	return structGo
}
func finalizeWebsocketExtensionDeflateClass(obj *WebsocketExtensionDeflateClass) {
	websocketExtensionDeflateClassStruct.Free(obj.native)
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
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// WebsocketExtensionManagerClassStruct creates an uninitialised WebsocketExtensionManagerClass.
func WebsocketExtensionManagerClassStruct() *WebsocketExtensionManagerClass {
	err := websocketExtensionManagerClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &WebsocketExtensionManagerClass{native: websocketExtensionManagerClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeWebsocketExtensionManagerClass)
	return structGo
}
func finalizeWebsocketExtensionManagerClass(obj *WebsocketExtensionManagerClass) {
	websocketExtensionManagerClassStruct.Free(obj.native)
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

	return
}

// UNSUPPORTED : C value 'soup_xmlrpc_params_parse' : return type 'GLib.Variant' not supported

// XMLRPCParamsStruct creates an uninitialised XMLRPCParams.
func XMLRPCParamsStruct() *XMLRPCParams {
	err := xMLRPCParamsStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &XMLRPCParams{native: xMLRPCParamsStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeXMLRPCParams)
	return structGo
}
func finalizeXMLRPCParams(obj *XMLRPCParams) {
	xMLRPCParamsStruct.Free(obj.native)
}
