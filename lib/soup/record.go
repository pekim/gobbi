// Code generated - DO NOT EDIT.

package soup

import (
	gi "github.com/pekim/gobbi/internal/cgo/gi"
	gio "github.com/pekim/gobbi/lib/gio"
	glib "github.com/pekim/gobbi/lib/glib"
	gobject "github.com/pekim/gobbi/lib/gobject"
	"runtime"
	"sync"
	"unsafe"
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
	native unsafe.Pointer
}

func AddressClassNewFromNative(native unsafe.Pointer) *AddressClass {
	instance := &AddressClass{native: native}

	return instance
}

/*
CastToAddressClass down casts any arbitrary Object to AddressClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a AddressClass.
*/
func CastToAddressClass(object *gobject.Object) *AddressClass {
	return AddressClassNewFromNative(object.Native())
}

// Equals compares this AddressClass with another AddressClass, and returns true if they represent the same Object.
func (recv *AddressClass) Equals(other *AddressClass) bool {
	return other.Native() == recv.Native()
}

func (recv *AddressClass) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentClass returns the C field 'parent_class'.
func (recv *AddressClass) FieldParentClass() *gobject.ObjectClass {
	argValue := gi.StructFieldGet(addressClassStruct, recv.Native(), "parent_class")
	value := gobject.ObjectClassNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *AddressClass) SetFieldParentClass(value *gobject.ObjectClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.StructFieldSet(addressClassStruct, recv.Native(), "parent_class", argValue)
}

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

	structGo := AddressClassNewFromNative(addressClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeAddressClass)
	return structGo
}
func finalizeAddressClass(obj *AddressClass) {
	addressClassStruct.Free(obj.Native())
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
	native unsafe.Pointer
}

func AuthClassNewFromNative(native unsafe.Pointer) *AuthClass {
	instance := &AuthClass{native: native}

	return instance
}

/*
CastToAuthClass down casts any arbitrary Object to AuthClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a AuthClass.
*/
func CastToAuthClass(object *gobject.Object) *AuthClass {
	return AuthClassNewFromNative(object.Native())
}

// Equals compares this AuthClass with another AuthClass, and returns true if they represent the same Object.
func (recv *AuthClass) Equals(other *AuthClass) bool {
	return other.Native() == recv.Native()
}

func (recv *AuthClass) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentClass returns the C field 'parent_class'.
func (recv *AuthClass) FieldParentClass() *gobject.ObjectClass {
	argValue := gi.StructFieldGet(authClassStruct, recv.Native(), "parent_class")
	value := gobject.ObjectClassNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *AuthClass) SetFieldParentClass(value *gobject.ObjectClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.StructFieldSet(authClassStruct, recv.Native(), "parent_class", argValue)
}

// FieldSchemeName returns the C field 'scheme_name'.
func (recv *AuthClass) FieldSchemeName() string {
	argValue := gi.StructFieldGet(authClassStruct, recv.Native(), "scheme_name")
	value := argValue.String(false)
	return value
}

// SetFieldSchemeName sets the value of the C field 'scheme_name'.
func (recv *AuthClass) SetFieldSchemeName(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.StructFieldSet(authClassStruct, recv.Native(), "scheme_name", argValue)
}

// FieldStrength returns the C field 'strength'.
func (recv *AuthClass) FieldStrength() uint32 {
	argValue := gi.StructFieldGet(authClassStruct, recv.Native(), "strength")
	value := argValue.Uint32()
	return value
}

// SetFieldStrength sets the value of the C field 'strength'.
func (recv *AuthClass) SetFieldStrength(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.StructFieldSet(authClassStruct, recv.Native(), "strength", argValue)
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

	structGo := AuthClassNewFromNative(authClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeAuthClass)
	return structGo
}
func finalizeAuthClass(obj *AuthClass) {
	authClassStruct.Free(obj.Native())
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
	native unsafe.Pointer
}

func AuthDomainBasicClassNewFromNative(native unsafe.Pointer) *AuthDomainBasicClass {
	instance := &AuthDomainBasicClass{native: native}

	return instance
}

/*
CastToAuthDomainBasicClass down casts any arbitrary Object to AuthDomainBasicClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a AuthDomainBasicClass.
*/
func CastToAuthDomainBasicClass(object *gobject.Object) *AuthDomainBasicClass {
	return AuthDomainBasicClassNewFromNative(object.Native())
}

// Equals compares this AuthDomainBasicClass with another AuthDomainBasicClass, and returns true if they represent the same Object.
func (recv *AuthDomainBasicClass) Equals(other *AuthDomainBasicClass) bool {
	return other.Native() == recv.Native()
}

func (recv *AuthDomainBasicClass) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentClass returns the C field 'parent_class'.
func (recv *AuthDomainBasicClass) FieldParentClass() *AuthDomainClass {
	argValue := gi.StructFieldGet(authDomainBasicClassStruct, recv.Native(), "parent_class")
	value := AuthDomainClassNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *AuthDomainBasicClass) SetFieldParentClass(value *AuthDomainClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.StructFieldSet(authDomainBasicClassStruct, recv.Native(), "parent_class", argValue)
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

	structGo := AuthDomainBasicClassNewFromNative(authDomainBasicClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeAuthDomainBasicClass)
	return structGo
}
func finalizeAuthDomainBasicClass(obj *AuthDomainBasicClass) {
	authDomainBasicClassStruct.Free(obj.Native())
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
	native unsafe.Pointer
}

func AuthDomainClassNewFromNative(native unsafe.Pointer) *AuthDomainClass {
	instance := &AuthDomainClass{native: native}

	return instance
}

/*
CastToAuthDomainClass down casts any arbitrary Object to AuthDomainClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a AuthDomainClass.
*/
func CastToAuthDomainClass(object *gobject.Object) *AuthDomainClass {
	return AuthDomainClassNewFromNative(object.Native())
}

// Equals compares this AuthDomainClass with another AuthDomainClass, and returns true if they represent the same Object.
func (recv *AuthDomainClass) Equals(other *AuthDomainClass) bool {
	return other.Native() == recv.Native()
}

func (recv *AuthDomainClass) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentClass returns the C field 'parent_class'.
func (recv *AuthDomainClass) FieldParentClass() *gobject.ObjectClass {
	argValue := gi.StructFieldGet(authDomainClassStruct, recv.Native(), "parent_class")
	value := gobject.ObjectClassNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *AuthDomainClass) SetFieldParentClass(value *gobject.ObjectClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.StructFieldSet(authDomainClassStruct, recv.Native(), "parent_class", argValue)
}

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

	structGo := AuthDomainClassNewFromNative(authDomainClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeAuthDomainClass)
	return structGo
}
func finalizeAuthDomainClass(obj *AuthDomainClass) {
	authDomainClassStruct.Free(obj.Native())
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
	native unsafe.Pointer
}

func AuthDomainDigestClassNewFromNative(native unsafe.Pointer) *AuthDomainDigestClass {
	instance := &AuthDomainDigestClass{native: native}

	return instance
}

/*
CastToAuthDomainDigestClass down casts any arbitrary Object to AuthDomainDigestClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a AuthDomainDigestClass.
*/
func CastToAuthDomainDigestClass(object *gobject.Object) *AuthDomainDigestClass {
	return AuthDomainDigestClassNewFromNative(object.Native())
}

// Equals compares this AuthDomainDigestClass with another AuthDomainDigestClass, and returns true if they represent the same Object.
func (recv *AuthDomainDigestClass) Equals(other *AuthDomainDigestClass) bool {
	return other.Native() == recv.Native()
}

func (recv *AuthDomainDigestClass) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentClass returns the C field 'parent_class'.
func (recv *AuthDomainDigestClass) FieldParentClass() *AuthDomainClass {
	argValue := gi.StructFieldGet(authDomainDigestClassStruct, recv.Native(), "parent_class")
	value := AuthDomainClassNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *AuthDomainDigestClass) SetFieldParentClass(value *AuthDomainClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.StructFieldSet(authDomainDigestClassStruct, recv.Native(), "parent_class", argValue)
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

	structGo := AuthDomainDigestClassNewFromNative(authDomainDigestClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeAuthDomainDigestClass)
	return structGo
}
func finalizeAuthDomainDigestClass(obj *AuthDomainDigestClass) {
	authDomainDigestClassStruct.Free(obj.Native())
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
	native unsafe.Pointer
}

func AuthManagerClassNewFromNative(native unsafe.Pointer) *AuthManagerClass {
	instance := &AuthManagerClass{native: native}

	return instance
}

/*
CastToAuthManagerClass down casts any arbitrary Object to AuthManagerClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a AuthManagerClass.
*/
func CastToAuthManagerClass(object *gobject.Object) *AuthManagerClass {
	return AuthManagerClassNewFromNative(object.Native())
}

// Equals compares this AuthManagerClass with another AuthManagerClass, and returns true if they represent the same Object.
func (recv *AuthManagerClass) Equals(other *AuthManagerClass) bool {
	return other.Native() == recv.Native()
}

func (recv *AuthManagerClass) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentClass returns the C field 'parent_class'.
func (recv *AuthManagerClass) FieldParentClass() *gobject.ObjectClass {
	argValue := gi.StructFieldGet(authManagerClassStruct, recv.Native(), "parent_class")
	value := gobject.ObjectClassNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *AuthManagerClass) SetFieldParentClass(value *gobject.ObjectClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.StructFieldSet(authManagerClassStruct, recv.Native(), "parent_class", argValue)
}

// UNSUPPORTED : C value 'authenticate' : for field getter : missing Type

// UNSUPPORTED : C value 'authenticate' : for field setter : missing Type

// AuthManagerClassStruct creates an uninitialised AuthManagerClass.
func AuthManagerClassStruct() *AuthManagerClass {
	err := authManagerClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := AuthManagerClassNewFromNative(authManagerClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeAuthManagerClass)
	return structGo
}
func finalizeAuthManagerClass(obj *AuthManagerClass) {
	authManagerClassStruct.Free(obj.Native())
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
	native unsafe.Pointer
}

func AuthManagerPrivateNewFromNative(native unsafe.Pointer) *AuthManagerPrivate {
	instance := &AuthManagerPrivate{native: native}

	return instance
}

/*
CastToAuthManagerPrivate down casts any arbitrary Object to AuthManagerPrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a AuthManagerPrivate.
*/
func CastToAuthManagerPrivate(object *gobject.Object) *AuthManagerPrivate {
	return AuthManagerPrivateNewFromNative(object.Native())
}

// Equals compares this AuthManagerPrivate with another AuthManagerPrivate, and returns true if they represent the same Object.
func (recv *AuthManagerPrivate) Equals(other *AuthManagerPrivate) bool {
	return other.Native() == recv.Native()
}

func (recv *AuthManagerPrivate) Native() unsafe.Pointer {
	return recv.native
}

// AuthManagerPrivateStruct creates an uninitialised AuthManagerPrivate.
func AuthManagerPrivateStruct() *AuthManagerPrivate {
	err := authManagerPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := AuthManagerPrivateNewFromNative(authManagerPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeAuthManagerPrivate)
	return structGo
}
func finalizeAuthManagerPrivate(obj *AuthManagerPrivate) {
	authManagerPrivateStruct.Free(obj.Native())
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
	native unsafe.Pointer
}

func BufferNewFromNative(native unsafe.Pointer) *Buffer {
	instance := &Buffer{native: native}

	return instance
}

/*
CastToBuffer down casts any arbitrary Object to Buffer.
Exercise care, as this is a potentially dangerous function
if the Object is not a Buffer.
*/
func CastToBuffer(object *gobject.Object) *Buffer {
	return BufferNewFromNative(object.Native())
}

// Equals compares this Buffer with another Buffer, and returns true if they represent the same Object.
func (recv *Buffer) Equals(other *Buffer) bool {
	return other.Native() == recv.Native()
}

func (recv *Buffer) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'data' : for field getter : no Go type for 'gpointer'

// UNSUPPORTED : C value 'data' : for field setter : no Go type for 'gpointer'

// FieldLength returns the C field 'length'.
func (recv *Buffer) FieldLength() uint64 {
	argValue := gi.StructFieldGet(bufferStruct, recv.Native(), "length")
	value := argValue.Uint64()
	return value
}

// SetFieldLength sets the value of the C field 'length'.
func (recv *Buffer) SetFieldLength(value uint64) {
	var argValue gi.Argument
	argValue.SetUint64(value)
	gi.StructFieldSet(bufferStruct, recv.Native(), "length", argValue)
}

// UNSUPPORTED : C value 'soup_buffer_new' : parameter 'data' of type 'nil' not supported

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
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := bufferCopyFunction_Set()
	if err == nil {
		ret = bufferCopyFunction.Invoke(inArgs[:], nil)
	}

	retGo := BufferNewFromNative(ret.Pointer())

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
	inArgs[0].SetPointer(recv.Native())

	err := bufferFreeFunction_Set()
	if err == nil {
		bufferFreeFunction.Invoke(inArgs[:], nil)
	}

	return
}

var bufferGetAsBytesFunction *gi.Function
var bufferGetAsBytesFunction_Once sync.Once

func bufferGetAsBytesFunction_Set() error {
	var err error
	bufferGetAsBytesFunction_Once.Do(func() {
		err = bufferStruct_Set()
		if err != nil {
			return
		}
		bufferGetAsBytesFunction, err = bufferStruct.InvokerNew("get_as_bytes")
	})
	return err
}

// GetAsBytes is a representation of the C type soup_buffer_get_as_bytes.
func (recv *Buffer) GetAsBytes() *glib.Bytes {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := bufferGetAsBytesFunction_Set()
	if err == nil {
		ret = bufferGetAsBytesFunction.Invoke(inArgs[:], nil)
	}

	retGo := glib.BytesNewFromNative(ret.Pointer())

	return retGo
}

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
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetUint64(offset)
	inArgs[2].SetUint64(length)

	var ret gi.Argument

	err := bufferNewSubbufferFunction_Set()
	if err == nil {
		ret = bufferNewSubbufferFunction.Invoke(inArgs[:], nil)
	}

	retGo := BufferNewFromNative(ret.Pointer())

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
	native unsafe.Pointer
}

func CacheClassNewFromNative(native unsafe.Pointer) *CacheClass {
	instance := &CacheClass{native: native}

	return instance
}

/*
CastToCacheClass down casts any arbitrary Object to CacheClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a CacheClass.
*/
func CastToCacheClass(object *gobject.Object) *CacheClass {
	return CacheClassNewFromNative(object.Native())
}

// Equals compares this CacheClass with another CacheClass, and returns true if they represent the same Object.
func (recv *CacheClass) Equals(other *CacheClass) bool {
	return other.Native() == recv.Native()
}

func (recv *CacheClass) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentClass returns the C field 'parent_class'.
func (recv *CacheClass) FieldParentClass() *gobject.ObjectClass {
	argValue := gi.StructFieldGet(cacheClassStruct, recv.Native(), "parent_class")
	value := gobject.ObjectClassNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *CacheClass) SetFieldParentClass(value *gobject.ObjectClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.StructFieldSet(cacheClassStruct, recv.Native(), "parent_class", argValue)
}

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

	structGo := CacheClassNewFromNative(cacheClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeCacheClass)
	return structGo
}
func finalizeCacheClass(obj *CacheClass) {
	cacheClassStruct.Free(obj.Native())
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
	native unsafe.Pointer
}

func CachePrivateNewFromNative(native unsafe.Pointer) *CachePrivate {
	instance := &CachePrivate{native: native}

	return instance
}

/*
CastToCachePrivate down casts any arbitrary Object to CachePrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a CachePrivate.
*/
func CastToCachePrivate(object *gobject.Object) *CachePrivate {
	return CachePrivateNewFromNative(object.Native())
}

// Equals compares this CachePrivate with another CachePrivate, and returns true if they represent the same Object.
func (recv *CachePrivate) Equals(other *CachePrivate) bool {
	return other.Native() == recv.Native()
}

func (recv *CachePrivate) Native() unsafe.Pointer {
	return recv.native
}

// CachePrivateStruct creates an uninitialised CachePrivate.
func CachePrivateStruct() *CachePrivate {
	err := cachePrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := CachePrivateNewFromNative(cachePrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeCachePrivate)
	return structGo
}
func finalizeCachePrivate(obj *CachePrivate) {
	cachePrivateStruct.Free(obj.Native())
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
	native unsafe.Pointer
}

func ClientContextNewFromNative(native unsafe.Pointer) *ClientContext {
	instance := &ClientContext{native: native}

	return instance
}

/*
CastToClientContext down casts any arbitrary Object to ClientContext.
Exercise care, as this is a potentially dangerous function
if the Object is not a ClientContext.
*/
func CastToClientContext(object *gobject.Object) *ClientContext {
	return ClientContextNewFromNative(object.Native())
}

// Equals compares this ClientContext with another ClientContext, and returns true if they represent the same Object.
func (recv *ClientContext) Equals(other *ClientContext) bool {
	return other.Native() == recv.Native()
}

func (recv *ClientContext) Native() unsafe.Pointer {
	return recv.native
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
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := clientContextGetAddressFunction_Set()
	if err == nil {
		ret = clientContextGetAddressFunction.Invoke(inArgs[:], nil)
	}

	retGo := AddressNewFromNative(ret.Pointer())

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
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := clientContextGetAuthDomainFunction_Set()
	if err == nil {
		ret = clientContextGetAuthDomainFunction.Invoke(inArgs[:], nil)
	}

	retGo := AuthDomainNewFromNative(ret.Pointer())

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
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := clientContextGetAuthUserFunction_Set()
	if err == nil {
		ret = clientContextGetAuthUserFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var clientContextGetGsocketFunction *gi.Function
var clientContextGetGsocketFunction_Once sync.Once

func clientContextGetGsocketFunction_Set() error {
	var err error
	clientContextGetGsocketFunction_Once.Do(func() {
		err = clientContextStruct_Set()
		if err != nil {
			return
		}
		clientContextGetGsocketFunction, err = clientContextStruct.InvokerNew("get_gsocket")
	})
	return err
}

// GetGsocket is a representation of the C type soup_client_context_get_gsocket.
func (recv *ClientContext) GetGsocket() *gio.Socket {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := clientContextGetGsocketFunction_Set()
	if err == nil {
		ret = clientContextGetGsocketFunction.Invoke(inArgs[:], nil)
	}

	retGo := gio.SocketNewFromNative(ret.Pointer())

	return retGo
}

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
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := clientContextGetHostFunction_Set()
	if err == nil {
		ret = clientContextGetHostFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var clientContextGetLocalAddressFunction *gi.Function
var clientContextGetLocalAddressFunction_Once sync.Once

func clientContextGetLocalAddressFunction_Set() error {
	var err error
	clientContextGetLocalAddressFunction_Once.Do(func() {
		err = clientContextStruct_Set()
		if err != nil {
			return
		}
		clientContextGetLocalAddressFunction, err = clientContextStruct.InvokerNew("get_local_address")
	})
	return err
}

// GetLocalAddress is a representation of the C type soup_client_context_get_local_address.
func (recv *ClientContext) GetLocalAddress() *gio.SocketAddress {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := clientContextGetLocalAddressFunction_Set()
	if err == nil {
		ret = clientContextGetLocalAddressFunction.Invoke(inArgs[:], nil)
	}

	retGo := gio.SocketAddressNewFromNative(ret.Pointer())

	return retGo
}

var clientContextGetRemoteAddressFunction *gi.Function
var clientContextGetRemoteAddressFunction_Once sync.Once

func clientContextGetRemoteAddressFunction_Set() error {
	var err error
	clientContextGetRemoteAddressFunction_Once.Do(func() {
		err = clientContextStruct_Set()
		if err != nil {
			return
		}
		clientContextGetRemoteAddressFunction, err = clientContextStruct.InvokerNew("get_remote_address")
	})
	return err
}

// GetRemoteAddress is a representation of the C type soup_client_context_get_remote_address.
func (recv *ClientContext) GetRemoteAddress() *gio.SocketAddress {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := clientContextGetRemoteAddressFunction_Set()
	if err == nil {
		ret = clientContextGetRemoteAddressFunction.Invoke(inArgs[:], nil)
	}

	retGo := gio.SocketAddressNewFromNative(ret.Pointer())

	return retGo
}

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
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := clientContextGetSocketFunction_Set()
	if err == nil {
		ret = clientContextGetSocketFunction.Invoke(inArgs[:], nil)
	}

	retGo := SocketNewFromNative(ret.Pointer())

	return retGo
}

var clientContextStealConnectionFunction *gi.Function
var clientContextStealConnectionFunction_Once sync.Once

func clientContextStealConnectionFunction_Set() error {
	var err error
	clientContextStealConnectionFunction_Once.Do(func() {
		err = clientContextStruct_Set()
		if err != nil {
			return
		}
		clientContextStealConnectionFunction, err = clientContextStruct.InvokerNew("steal_connection")
	})
	return err
}

// StealConnection is a representation of the C type soup_client_context_steal_connection.
func (recv *ClientContext) StealConnection() *gio.IOStream {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := clientContextStealConnectionFunction_Set()
	if err == nil {
		ret = clientContextStealConnectionFunction.Invoke(inArgs[:], nil)
	}

	retGo := gio.IOStreamNewFromNative(ret.Pointer())

	return retGo
}

// ClientContextStruct creates an uninitialised ClientContext.
func ClientContextStruct() *ClientContext {
	err := clientContextStruct_Set()
	if err != nil {
		return nil
	}

	structGo := ClientContextNewFromNative(clientContextStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeClientContext)
	return structGo
}
func finalizeClientContext(obj *ClientContext) {
	clientContextStruct.Free(obj.Native())
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
	native unsafe.Pointer
}

func ConnectionNewFromNative(native unsafe.Pointer) *Connection {
	instance := &Connection{native: native}

	return instance
}

/*
CastToConnection down casts any arbitrary Object to Connection.
Exercise care, as this is a potentially dangerous function
if the Object is not a Connection.
*/
func CastToConnection(object *gobject.Object) *Connection {
	return ConnectionNewFromNative(object.Native())
}

// Equals compares this Connection with another Connection, and returns true if they represent the same Object.
func (recv *Connection) Equals(other *Connection) bool {
	return other.Native() == recv.Native()
}

func (recv *Connection) Native() unsafe.Pointer {
	return recv.native
}

// ConnectionStruct creates an uninitialised Connection.
func ConnectionStruct() *Connection {
	err := connectionStruct_Set()
	if err != nil {
		return nil
	}

	structGo := ConnectionNewFromNative(connectionStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeConnection)
	return structGo
}
func finalizeConnection(obj *Connection) {
	connectionStruct.Free(obj.Native())
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
	native unsafe.Pointer
}

func ContentDecoderClassNewFromNative(native unsafe.Pointer) *ContentDecoderClass {
	instance := &ContentDecoderClass{native: native}

	return instance
}

/*
CastToContentDecoderClass down casts any arbitrary Object to ContentDecoderClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a ContentDecoderClass.
*/
func CastToContentDecoderClass(object *gobject.Object) *ContentDecoderClass {
	return ContentDecoderClassNewFromNative(object.Native())
}

// Equals compares this ContentDecoderClass with another ContentDecoderClass, and returns true if they represent the same Object.
func (recv *ContentDecoderClass) Equals(other *ContentDecoderClass) bool {
	return other.Native() == recv.Native()
}

func (recv *ContentDecoderClass) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentClass returns the C field 'parent_class'.
func (recv *ContentDecoderClass) FieldParentClass() *gobject.ObjectClass {
	argValue := gi.StructFieldGet(contentDecoderClassStruct, recv.Native(), "parent_class")
	value := gobject.ObjectClassNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *ContentDecoderClass) SetFieldParentClass(value *gobject.ObjectClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.StructFieldSet(contentDecoderClassStruct, recv.Native(), "parent_class", argValue)
}

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

	structGo := ContentDecoderClassNewFromNative(contentDecoderClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeContentDecoderClass)
	return structGo
}
func finalizeContentDecoderClass(obj *ContentDecoderClass) {
	contentDecoderClassStruct.Free(obj.Native())
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
	native unsafe.Pointer
}

func ContentDecoderPrivateNewFromNative(native unsafe.Pointer) *ContentDecoderPrivate {
	instance := &ContentDecoderPrivate{native: native}

	return instance
}

/*
CastToContentDecoderPrivate down casts any arbitrary Object to ContentDecoderPrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a ContentDecoderPrivate.
*/
func CastToContentDecoderPrivate(object *gobject.Object) *ContentDecoderPrivate {
	return ContentDecoderPrivateNewFromNative(object.Native())
}

// Equals compares this ContentDecoderPrivate with another ContentDecoderPrivate, and returns true if they represent the same Object.
func (recv *ContentDecoderPrivate) Equals(other *ContentDecoderPrivate) bool {
	return other.Native() == recv.Native()
}

func (recv *ContentDecoderPrivate) Native() unsafe.Pointer {
	return recv.native
}

// ContentDecoderPrivateStruct creates an uninitialised ContentDecoderPrivate.
func ContentDecoderPrivateStruct() *ContentDecoderPrivate {
	err := contentDecoderPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := ContentDecoderPrivateNewFromNative(contentDecoderPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeContentDecoderPrivate)
	return structGo
}
func finalizeContentDecoderPrivate(obj *ContentDecoderPrivate) {
	contentDecoderPrivateStruct.Free(obj.Native())
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
	native unsafe.Pointer
}

func ContentSnifferClassNewFromNative(native unsafe.Pointer) *ContentSnifferClass {
	instance := &ContentSnifferClass{native: native}

	return instance
}

/*
CastToContentSnifferClass down casts any arbitrary Object to ContentSnifferClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a ContentSnifferClass.
*/
func CastToContentSnifferClass(object *gobject.Object) *ContentSnifferClass {
	return ContentSnifferClassNewFromNative(object.Native())
}

// Equals compares this ContentSnifferClass with another ContentSnifferClass, and returns true if they represent the same Object.
func (recv *ContentSnifferClass) Equals(other *ContentSnifferClass) bool {
	return other.Native() == recv.Native()
}

func (recv *ContentSnifferClass) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentClass returns the C field 'parent_class'.
func (recv *ContentSnifferClass) FieldParentClass() *gobject.ObjectClass {
	argValue := gi.StructFieldGet(contentSnifferClassStruct, recv.Native(), "parent_class")
	value := gobject.ObjectClassNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *ContentSnifferClass) SetFieldParentClass(value *gobject.ObjectClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.StructFieldSet(contentSnifferClassStruct, recv.Native(), "parent_class", argValue)
}

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

	structGo := ContentSnifferClassNewFromNative(contentSnifferClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeContentSnifferClass)
	return structGo
}
func finalizeContentSnifferClass(obj *ContentSnifferClass) {
	contentSnifferClassStruct.Free(obj.Native())
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
	native unsafe.Pointer
}

func ContentSnifferPrivateNewFromNative(native unsafe.Pointer) *ContentSnifferPrivate {
	instance := &ContentSnifferPrivate{native: native}

	return instance
}

/*
CastToContentSnifferPrivate down casts any arbitrary Object to ContentSnifferPrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a ContentSnifferPrivate.
*/
func CastToContentSnifferPrivate(object *gobject.Object) *ContentSnifferPrivate {
	return ContentSnifferPrivateNewFromNative(object.Native())
}

// Equals compares this ContentSnifferPrivate with another ContentSnifferPrivate, and returns true if they represent the same Object.
func (recv *ContentSnifferPrivate) Equals(other *ContentSnifferPrivate) bool {
	return other.Native() == recv.Native()
}

func (recv *ContentSnifferPrivate) Native() unsafe.Pointer {
	return recv.native
}

// ContentSnifferPrivateStruct creates an uninitialised ContentSnifferPrivate.
func ContentSnifferPrivateStruct() *ContentSnifferPrivate {
	err := contentSnifferPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := ContentSnifferPrivateNewFromNative(contentSnifferPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeContentSnifferPrivate)
	return structGo
}
func finalizeContentSnifferPrivate(obj *ContentSnifferPrivate) {
	contentSnifferPrivateStruct.Free(obj.Native())
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
	native unsafe.Pointer
}

func CookieNewFromNative(native unsafe.Pointer) *Cookie {
	instance := &Cookie{native: native}

	return instance
}

/*
CastToCookie down casts any arbitrary Object to Cookie.
Exercise care, as this is a potentially dangerous function
if the Object is not a Cookie.
*/
func CastToCookie(object *gobject.Object) *Cookie {
	return CookieNewFromNative(object.Native())
}

// Equals compares this Cookie with another Cookie, and returns true if they represent the same Object.
func (recv *Cookie) Equals(other *Cookie) bool {
	return other.Native() == recv.Native()
}

func (recv *Cookie) Native() unsafe.Pointer {
	return recv.native
}

// FieldName returns the C field 'name'.
func (recv *Cookie) FieldName() string {
	argValue := gi.StructFieldGet(cookieStruct, recv.Native(), "name")
	value := argValue.String(false)
	return value
}

// SetFieldName sets the value of the C field 'name'.
func (recv *Cookie) SetFieldName(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.StructFieldSet(cookieStruct, recv.Native(), "name", argValue)
}

// FieldValue returns the C field 'value'.
func (recv *Cookie) FieldValue() string {
	argValue := gi.StructFieldGet(cookieStruct, recv.Native(), "value")
	value := argValue.String(false)
	return value
}

// SetFieldValue sets the value of the C field 'value'.
func (recv *Cookie) SetFieldValue(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.StructFieldSet(cookieStruct, recv.Native(), "value", argValue)
}

// FieldDomain returns the C field 'domain'.
func (recv *Cookie) FieldDomain() string {
	argValue := gi.StructFieldGet(cookieStruct, recv.Native(), "domain")
	value := argValue.String(false)
	return value
}

// SetFieldDomain sets the value of the C field 'domain'.
func (recv *Cookie) SetFieldDomain(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.StructFieldSet(cookieStruct, recv.Native(), "domain", argValue)
}

// FieldPath returns the C field 'path'.
func (recv *Cookie) FieldPath() string {
	argValue := gi.StructFieldGet(cookieStruct, recv.Native(), "path")
	value := argValue.String(false)
	return value
}

// SetFieldPath sets the value of the C field 'path'.
func (recv *Cookie) SetFieldPath(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.StructFieldSet(cookieStruct, recv.Native(), "path", argValue)
}

// FieldExpires returns the C field 'expires'.
func (recv *Cookie) FieldExpires() *Date {
	argValue := gi.StructFieldGet(cookieStruct, recv.Native(), "expires")
	value := DateNewFromNative(argValue.Pointer())
	return value
}

// SetFieldExpires sets the value of the C field 'expires'.
func (recv *Cookie) SetFieldExpires(value *Date) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.StructFieldSet(cookieStruct, recv.Native(), "expires", argValue)
}

// FieldSecure returns the C field 'secure'.
func (recv *Cookie) FieldSecure() bool {
	argValue := gi.StructFieldGet(cookieStruct, recv.Native(), "secure")
	value := argValue.Boolean()
	return value
}

// SetFieldSecure sets the value of the C field 'secure'.
func (recv *Cookie) SetFieldSecure(value bool) {
	var argValue gi.Argument
	argValue.SetBoolean(value)
	gi.StructFieldSet(cookieStruct, recv.Native(), "secure", argValue)
}

// FieldHttpOnly returns the C field 'http_only'.
func (recv *Cookie) FieldHttpOnly() bool {
	argValue := gi.StructFieldGet(cookieStruct, recv.Native(), "http_only")
	value := argValue.Boolean()
	return value
}

// SetFieldHttpOnly sets the value of the C field 'http_only'.
func (recv *Cookie) SetFieldHttpOnly(value bool) {
	var argValue gi.Argument
	argValue.SetBoolean(value)
	gi.StructFieldSet(cookieStruct, recv.Native(), "http_only", argValue)
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

	retGo := CookieNewFromNative(ret.Pointer())

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
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(uri.Native())

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
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := cookieCopyFunction_Set()
	if err == nil {
		ret = cookieCopyFunction.Invoke(inArgs[:], nil)
	}

	retGo := CookieNewFromNative(ret.Pointer())

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
	inArgs[0].SetPointer(recv.Native())
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
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(cookie2.Native())

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
	inArgs[0].SetPointer(recv.Native())

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
	inArgs[0].SetPointer(recv.Native())

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
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := cookieGetExpiresFunction_Set()
	if err == nil {
		ret = cookieGetExpiresFunction.Invoke(inArgs[:], nil)
	}

	retGo := DateNewFromNative(ret.Pointer())

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
	inArgs[0].SetPointer(recv.Native())

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
	inArgs[0].SetPointer(recv.Native())

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
	inArgs[0].SetPointer(recv.Native())

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
	inArgs[0].SetPointer(recv.Native())

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
	inArgs[0].SetPointer(recv.Native())

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
	inArgs[0].SetPointer(recv.Native())
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
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(expires.Native())

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
	inArgs[0].SetPointer(recv.Native())
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
	inArgs[0].SetPointer(recv.Native())
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
	inArgs[0].SetPointer(recv.Native())
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
	inArgs[0].SetPointer(recv.Native())
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
	inArgs[0].SetPointer(recv.Native())
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
	inArgs[0].SetPointer(recv.Native())
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
	inArgs[0].SetPointer(recv.Native())

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
	inArgs[0].SetPointer(recv.Native())

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
	native unsafe.Pointer
}

func CookieJarClassNewFromNative(native unsafe.Pointer) *CookieJarClass {
	instance := &CookieJarClass{native: native}

	return instance
}

/*
CastToCookieJarClass down casts any arbitrary Object to CookieJarClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a CookieJarClass.
*/
func CastToCookieJarClass(object *gobject.Object) *CookieJarClass {
	return CookieJarClassNewFromNative(object.Native())
}

// Equals compares this CookieJarClass with another CookieJarClass, and returns true if they represent the same Object.
func (recv *CookieJarClass) Equals(other *CookieJarClass) bool {
	return other.Native() == recv.Native()
}

func (recv *CookieJarClass) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentClass returns the C field 'parent_class'.
func (recv *CookieJarClass) FieldParentClass() *gobject.ObjectClass {
	argValue := gi.StructFieldGet(cookieJarClassStruct, recv.Native(), "parent_class")
	value := gobject.ObjectClassNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *CookieJarClass) SetFieldParentClass(value *gobject.ObjectClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.StructFieldSet(cookieJarClassStruct, recv.Native(), "parent_class", argValue)
}

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

	structGo := CookieJarClassNewFromNative(cookieJarClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeCookieJarClass)
	return structGo
}
func finalizeCookieJarClass(obj *CookieJarClass) {
	cookieJarClassStruct.Free(obj.Native())
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
	native unsafe.Pointer
}

func CookieJarDBClassNewFromNative(native unsafe.Pointer) *CookieJarDBClass {
	instance := &CookieJarDBClass{native: native}

	return instance
}

/*
CastToCookieJarDBClass down casts any arbitrary Object to CookieJarDBClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a CookieJarDBClass.
*/
func CastToCookieJarDBClass(object *gobject.Object) *CookieJarDBClass {
	return CookieJarDBClassNewFromNative(object.Native())
}

// Equals compares this CookieJarDBClass with another CookieJarDBClass, and returns true if they represent the same Object.
func (recv *CookieJarDBClass) Equals(other *CookieJarDBClass) bool {
	return other.Native() == recv.Native()
}

func (recv *CookieJarDBClass) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentClass returns the C field 'parent_class'.
func (recv *CookieJarDBClass) FieldParentClass() *CookieJarClass {
	argValue := gi.StructFieldGet(cookieJarDBClassStruct, recv.Native(), "parent_class")
	value := CookieJarClassNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *CookieJarDBClass) SetFieldParentClass(value *CookieJarClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.StructFieldSet(cookieJarDBClassStruct, recv.Native(), "parent_class", argValue)
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

	structGo := CookieJarDBClassNewFromNative(cookieJarDBClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeCookieJarDBClass)
	return structGo
}
func finalizeCookieJarDBClass(obj *CookieJarDBClass) {
	cookieJarDBClassStruct.Free(obj.Native())
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
	native unsafe.Pointer
}

func CookieJarTextClassNewFromNative(native unsafe.Pointer) *CookieJarTextClass {
	instance := &CookieJarTextClass{native: native}

	return instance
}

/*
CastToCookieJarTextClass down casts any arbitrary Object to CookieJarTextClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a CookieJarTextClass.
*/
func CastToCookieJarTextClass(object *gobject.Object) *CookieJarTextClass {
	return CookieJarTextClassNewFromNative(object.Native())
}

// Equals compares this CookieJarTextClass with another CookieJarTextClass, and returns true if they represent the same Object.
func (recv *CookieJarTextClass) Equals(other *CookieJarTextClass) bool {
	return other.Native() == recv.Native()
}

func (recv *CookieJarTextClass) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentClass returns the C field 'parent_class'.
func (recv *CookieJarTextClass) FieldParentClass() *CookieJarClass {
	argValue := gi.StructFieldGet(cookieJarTextClassStruct, recv.Native(), "parent_class")
	value := CookieJarClassNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *CookieJarTextClass) SetFieldParentClass(value *CookieJarClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.StructFieldSet(cookieJarTextClassStruct, recv.Native(), "parent_class", argValue)
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

	structGo := CookieJarTextClassNewFromNative(cookieJarTextClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeCookieJarTextClass)
	return structGo
}
func finalizeCookieJarTextClass(obj *CookieJarTextClass) {
	cookieJarTextClassStruct.Free(obj.Native())
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
	native unsafe.Pointer
}

func DateNewFromNative(native unsafe.Pointer) *Date {
	instance := &Date{native: native}

	return instance
}

/*
CastToDate down casts any arbitrary Object to Date.
Exercise care, as this is a potentially dangerous function
if the Object is not a Date.
*/
func CastToDate(object *gobject.Object) *Date {
	return DateNewFromNative(object.Native())
}

// Equals compares this Date with another Date, and returns true if they represent the same Object.
func (recv *Date) Equals(other *Date) bool {
	return other.Native() == recv.Native()
}

func (recv *Date) Native() unsafe.Pointer {
	return recv.native
}

// FieldYear returns the C field 'year'.
func (recv *Date) FieldYear() int32 {
	argValue := gi.StructFieldGet(dateStruct, recv.Native(), "year")
	value := argValue.Int32()
	return value
}

// SetFieldYear sets the value of the C field 'year'.
func (recv *Date) SetFieldYear(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.StructFieldSet(dateStruct, recv.Native(), "year", argValue)
}

// FieldMonth returns the C field 'month'.
func (recv *Date) FieldMonth() int32 {
	argValue := gi.StructFieldGet(dateStruct, recv.Native(), "month")
	value := argValue.Int32()
	return value
}

// SetFieldMonth sets the value of the C field 'month'.
func (recv *Date) SetFieldMonth(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.StructFieldSet(dateStruct, recv.Native(), "month", argValue)
}

// FieldDay returns the C field 'day'.
func (recv *Date) FieldDay() int32 {
	argValue := gi.StructFieldGet(dateStruct, recv.Native(), "day")
	value := argValue.Int32()
	return value
}

// SetFieldDay sets the value of the C field 'day'.
func (recv *Date) SetFieldDay(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.StructFieldSet(dateStruct, recv.Native(), "day", argValue)
}

// FieldHour returns the C field 'hour'.
func (recv *Date) FieldHour() int32 {
	argValue := gi.StructFieldGet(dateStruct, recv.Native(), "hour")
	value := argValue.Int32()
	return value
}

// SetFieldHour sets the value of the C field 'hour'.
func (recv *Date) SetFieldHour(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.StructFieldSet(dateStruct, recv.Native(), "hour", argValue)
}

// FieldMinute returns the C field 'minute'.
func (recv *Date) FieldMinute() int32 {
	argValue := gi.StructFieldGet(dateStruct, recv.Native(), "minute")
	value := argValue.Int32()
	return value
}

// SetFieldMinute sets the value of the C field 'minute'.
func (recv *Date) SetFieldMinute(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.StructFieldSet(dateStruct, recv.Native(), "minute", argValue)
}

// FieldSecond returns the C field 'second'.
func (recv *Date) FieldSecond() int32 {
	argValue := gi.StructFieldGet(dateStruct, recv.Native(), "second")
	value := argValue.Int32()
	return value
}

// SetFieldSecond sets the value of the C field 'second'.
func (recv *Date) SetFieldSecond(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.StructFieldSet(dateStruct, recv.Native(), "second", argValue)
}

// FieldUtc returns the C field 'utc'.
func (recv *Date) FieldUtc() bool {
	argValue := gi.StructFieldGet(dateStruct, recv.Native(), "utc")
	value := argValue.Boolean()
	return value
}

// SetFieldUtc sets the value of the C field 'utc'.
func (recv *Date) SetFieldUtc(value bool) {
	var argValue gi.Argument
	argValue.SetBoolean(value)
	gi.StructFieldSet(dateStruct, recv.Native(), "utc", argValue)
}

// FieldOffset returns the C field 'offset'.
func (recv *Date) FieldOffset() int32 {
	argValue := gi.StructFieldGet(dateStruct, recv.Native(), "offset")
	value := argValue.Int32()
	return value
}

// SetFieldOffset sets the value of the C field 'offset'.
func (recv *Date) SetFieldOffset(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.StructFieldSet(dateStruct, recv.Native(), "offset", argValue)
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

	retGo := DateNewFromNative(ret.Pointer())

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

	retGo := DateNewFromNative(ret.Pointer())

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

	retGo := DateNewFromNative(ret.Pointer())

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

	retGo := DateNewFromNative(ret.Pointer())

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
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := dateCopyFunction_Set()
	if err == nil {
		ret = dateCopyFunction.Invoke(inArgs[:], nil)
	}

	retGo := DateNewFromNative(ret.Pointer())

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
	inArgs[0].SetPointer(recv.Native())

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
	inArgs[0].SetPointer(recv.Native())

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
	inArgs[0].SetPointer(recv.Native())

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
	inArgs[0].SetPointer(recv.Native())

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
	inArgs[0].SetPointer(recv.Native())

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
	inArgs[0].SetPointer(recv.Native())

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
	inArgs[0].SetPointer(recv.Native())

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
	inArgs[0].SetPointer(recv.Native())

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
	inArgs[0].SetPointer(recv.Native())

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
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := dateIsPastFunction_Set()
	if err == nil {
		ret = dateIsPastFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var dateToStringFunction *gi.Function
var dateToStringFunction_Once sync.Once

func dateToStringFunction_Set() error {
	var err error
	dateToStringFunction_Once.Do(func() {
		err = dateStruct_Set()
		if err != nil {
			return
		}
		dateToStringFunction, err = dateStruct.InvokerNew("to_string")
	})
	return err
}

// ToString is a representation of the C type soup_date_to_string.
func (recv *Date) ToString(format DateFormat) string {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(int32(format))

	var ret gi.Argument

	err := dateToStringFunction_Set()
	if err == nil {
		ret = dateToStringFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

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
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := dateToTimeTFunction_Set()
	if err == nil {
		ret = dateToTimeTFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int64()

	return retGo
}

var dateToTimevalFunction *gi.Function
var dateToTimevalFunction_Once sync.Once

func dateToTimevalFunction_Set() error {
	var err error
	dateToTimevalFunction_Once.Do(func() {
		err = dateStruct_Set()
		if err != nil {
			return
		}
		dateToTimevalFunction, err = dateStruct.InvokerNew("to_timeval")
	})
	return err
}

// ToTimeval is a representation of the C type soup_date_to_timeval.
func (recv *Date) ToTimeval() *glib.TimeVal {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var outArgs [1]gi.Argument

	err := dateToTimevalFunction_Set()
	if err == nil {
		dateToTimevalFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := glib.TimeValNewFromNative(outArgs[0].Pointer())

	return out0
}

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
	native unsafe.Pointer
}

func HSTSEnforcerClassNewFromNative(native unsafe.Pointer) *HSTSEnforcerClass {
	instance := &HSTSEnforcerClass{native: native}

	return instance
}

/*
CastToHSTSEnforcerClass down casts any arbitrary Object to HSTSEnforcerClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a HSTSEnforcerClass.
*/
func CastToHSTSEnforcerClass(object *gobject.Object) *HSTSEnforcerClass {
	return HSTSEnforcerClassNewFromNative(object.Native())
}

// Equals compares this HSTSEnforcerClass with another HSTSEnforcerClass, and returns true if they represent the same Object.
func (recv *HSTSEnforcerClass) Equals(other *HSTSEnforcerClass) bool {
	return other.Native() == recv.Native()
}

func (recv *HSTSEnforcerClass) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentClass returns the C field 'parent_class'.
func (recv *HSTSEnforcerClass) FieldParentClass() *gobject.ObjectClass {
	argValue := gi.StructFieldGet(hSTSEnforcerClassStruct, recv.Native(), "parent_class")
	value := gobject.ObjectClassNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *HSTSEnforcerClass) SetFieldParentClass(value *gobject.ObjectClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.StructFieldSet(hSTSEnforcerClassStruct, recv.Native(), "parent_class", argValue)
}

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

	structGo := HSTSEnforcerClassNewFromNative(hSTSEnforcerClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeHSTSEnforcerClass)
	return structGo
}
func finalizeHSTSEnforcerClass(obj *HSTSEnforcerClass) {
	hSTSEnforcerClassStruct.Free(obj.Native())
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
	native unsafe.Pointer
}

func HSTSEnforcerDBClassNewFromNative(native unsafe.Pointer) *HSTSEnforcerDBClass {
	instance := &HSTSEnforcerDBClass{native: native}

	return instance
}

/*
CastToHSTSEnforcerDBClass down casts any arbitrary Object to HSTSEnforcerDBClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a HSTSEnforcerDBClass.
*/
func CastToHSTSEnforcerDBClass(object *gobject.Object) *HSTSEnforcerDBClass {
	return HSTSEnforcerDBClassNewFromNative(object.Native())
}

// Equals compares this HSTSEnforcerDBClass with another HSTSEnforcerDBClass, and returns true if they represent the same Object.
func (recv *HSTSEnforcerDBClass) Equals(other *HSTSEnforcerDBClass) bool {
	return other.Native() == recv.Native()
}

func (recv *HSTSEnforcerDBClass) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentClass returns the C field 'parent_class'.
func (recv *HSTSEnforcerDBClass) FieldParentClass() *HSTSEnforcerClass {
	argValue := gi.StructFieldGet(hSTSEnforcerDBClassStruct, recv.Native(), "parent_class")
	value := HSTSEnforcerClassNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *HSTSEnforcerDBClass) SetFieldParentClass(value *HSTSEnforcerClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.StructFieldSet(hSTSEnforcerDBClassStruct, recv.Native(), "parent_class", argValue)
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

	structGo := HSTSEnforcerDBClassNewFromNative(hSTSEnforcerDBClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeHSTSEnforcerDBClass)
	return structGo
}
func finalizeHSTSEnforcerDBClass(obj *HSTSEnforcerDBClass) {
	hSTSEnforcerDBClassStruct.Free(obj.Native())
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
	native unsafe.Pointer
}

func HSTSEnforcerDBPrivateNewFromNative(native unsafe.Pointer) *HSTSEnforcerDBPrivate {
	instance := &HSTSEnforcerDBPrivate{native: native}

	return instance
}

/*
CastToHSTSEnforcerDBPrivate down casts any arbitrary Object to HSTSEnforcerDBPrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a HSTSEnforcerDBPrivate.
*/
func CastToHSTSEnforcerDBPrivate(object *gobject.Object) *HSTSEnforcerDBPrivate {
	return HSTSEnforcerDBPrivateNewFromNative(object.Native())
}

// Equals compares this HSTSEnforcerDBPrivate with another HSTSEnforcerDBPrivate, and returns true if they represent the same Object.
func (recv *HSTSEnforcerDBPrivate) Equals(other *HSTSEnforcerDBPrivate) bool {
	return other.Native() == recv.Native()
}

func (recv *HSTSEnforcerDBPrivate) Native() unsafe.Pointer {
	return recv.native
}

// HSTSEnforcerDBPrivateStruct creates an uninitialised HSTSEnforcerDBPrivate.
func HSTSEnforcerDBPrivateStruct() *HSTSEnforcerDBPrivate {
	err := hSTSEnforcerDBPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := HSTSEnforcerDBPrivateNewFromNative(hSTSEnforcerDBPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeHSTSEnforcerDBPrivate)
	return structGo
}
func finalizeHSTSEnforcerDBPrivate(obj *HSTSEnforcerDBPrivate) {
	hSTSEnforcerDBPrivateStruct.Free(obj.Native())
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
	native unsafe.Pointer
}

func HSTSEnforcerPrivateNewFromNative(native unsafe.Pointer) *HSTSEnforcerPrivate {
	instance := &HSTSEnforcerPrivate{native: native}

	return instance
}

/*
CastToHSTSEnforcerPrivate down casts any arbitrary Object to HSTSEnforcerPrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a HSTSEnforcerPrivate.
*/
func CastToHSTSEnforcerPrivate(object *gobject.Object) *HSTSEnforcerPrivate {
	return HSTSEnforcerPrivateNewFromNative(object.Native())
}

// Equals compares this HSTSEnforcerPrivate with another HSTSEnforcerPrivate, and returns true if they represent the same Object.
func (recv *HSTSEnforcerPrivate) Equals(other *HSTSEnforcerPrivate) bool {
	return other.Native() == recv.Native()
}

func (recv *HSTSEnforcerPrivate) Native() unsafe.Pointer {
	return recv.native
}

// HSTSEnforcerPrivateStruct creates an uninitialised HSTSEnforcerPrivate.
func HSTSEnforcerPrivateStruct() *HSTSEnforcerPrivate {
	err := hSTSEnforcerPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := HSTSEnforcerPrivateNewFromNative(hSTSEnforcerPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeHSTSEnforcerPrivate)
	return structGo
}
func finalizeHSTSEnforcerPrivate(obj *HSTSEnforcerPrivate) {
	hSTSEnforcerPrivateStruct.Free(obj.Native())
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
	native unsafe.Pointer
}

func HSTSPolicyNewFromNative(native unsafe.Pointer) *HSTSPolicy {
	instance := &HSTSPolicy{native: native}

	return instance
}

/*
CastToHSTSPolicy down casts any arbitrary Object to HSTSPolicy.
Exercise care, as this is a potentially dangerous function
if the Object is not a HSTSPolicy.
*/
func CastToHSTSPolicy(object *gobject.Object) *HSTSPolicy {
	return HSTSPolicyNewFromNative(object.Native())
}

// Equals compares this HSTSPolicy with another HSTSPolicy, and returns true if they represent the same Object.
func (recv *HSTSPolicy) Equals(other *HSTSPolicy) bool {
	return other.Native() == recv.Native()
}

func (recv *HSTSPolicy) Native() unsafe.Pointer {
	return recv.native
}

// FieldDomain returns the C field 'domain'.
func (recv *HSTSPolicy) FieldDomain() string {
	argValue := gi.StructFieldGet(hSTSPolicyStruct, recv.Native(), "domain")
	value := argValue.String(false)
	return value
}

// SetFieldDomain sets the value of the C field 'domain'.
func (recv *HSTSPolicy) SetFieldDomain(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.StructFieldSet(hSTSPolicyStruct, recv.Native(), "domain", argValue)
}

// FieldMaxAge returns the C field 'max_age'.
func (recv *HSTSPolicy) FieldMaxAge() uint64 {
	argValue := gi.StructFieldGet(hSTSPolicyStruct, recv.Native(), "max_age")
	value := argValue.Uint64()
	return value
}

// SetFieldMaxAge sets the value of the C field 'max_age'.
func (recv *HSTSPolicy) SetFieldMaxAge(value uint64) {
	var argValue gi.Argument
	argValue.SetUint64(value)
	gi.StructFieldSet(hSTSPolicyStruct, recv.Native(), "max_age", argValue)
}

// FieldExpires returns the C field 'expires'.
func (recv *HSTSPolicy) FieldExpires() *Date {
	argValue := gi.StructFieldGet(hSTSPolicyStruct, recv.Native(), "expires")
	value := DateNewFromNative(argValue.Pointer())
	return value
}

// SetFieldExpires sets the value of the C field 'expires'.
func (recv *HSTSPolicy) SetFieldExpires(value *Date) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.StructFieldSet(hSTSPolicyStruct, recv.Native(), "expires", argValue)
}

// FieldIncludeSubdomains returns the C field 'include_subdomains'.
func (recv *HSTSPolicy) FieldIncludeSubdomains() bool {
	argValue := gi.StructFieldGet(hSTSPolicyStruct, recv.Native(), "include_subdomains")
	value := argValue.Boolean()
	return value
}

// SetFieldIncludeSubdomains sets the value of the C field 'include_subdomains'.
func (recv *HSTSPolicy) SetFieldIncludeSubdomains(value bool) {
	var argValue gi.Argument
	argValue.SetBoolean(value)
	gi.StructFieldSet(hSTSPolicyStruct, recv.Native(), "include_subdomains", argValue)
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

	retGo := HSTSPolicyNewFromNative(ret.Pointer())

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
	inArgs[0].SetPointer(msg.Native())

	var ret gi.Argument

	err := hSTSPolicyNewFromResponseFunction_Set()
	if err == nil {
		ret = hSTSPolicyNewFromResponseFunction.Invoke(inArgs[:], nil)
	}

	retGo := HSTSPolicyNewFromNative(ret.Pointer())

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
	inArgs[2].SetPointer(expires.Native())
	inArgs[3].SetBoolean(includeSubdomains)

	var ret gi.Argument

	err := hSTSPolicyNewFullFunction_Set()
	if err == nil {
		ret = hSTSPolicyNewFullFunction.Invoke(inArgs[:], nil)
	}

	retGo := HSTSPolicyNewFromNative(ret.Pointer())

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

	retGo := HSTSPolicyNewFromNative(ret.Pointer())

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
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := hSTSPolicyCopyFunction_Set()
	if err == nil {
		ret = hSTSPolicyCopyFunction.Invoke(inArgs[:], nil)
	}

	retGo := HSTSPolicyNewFromNative(ret.Pointer())

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
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(policy2.Native())

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
	inArgs[0].SetPointer(recv.Native())

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
	inArgs[0].SetPointer(recv.Native())

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
	inArgs[0].SetPointer(recv.Native())

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
	inArgs[0].SetPointer(recv.Native())

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
	inArgs[0].SetPointer(recv.Native())

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
	native unsafe.Pointer
}

func LoggerClassNewFromNative(native unsafe.Pointer) *LoggerClass {
	instance := &LoggerClass{native: native}

	return instance
}

/*
CastToLoggerClass down casts any arbitrary Object to LoggerClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a LoggerClass.
*/
func CastToLoggerClass(object *gobject.Object) *LoggerClass {
	return LoggerClassNewFromNative(object.Native())
}

// Equals compares this LoggerClass with another LoggerClass, and returns true if they represent the same Object.
func (recv *LoggerClass) Equals(other *LoggerClass) bool {
	return other.Native() == recv.Native()
}

func (recv *LoggerClass) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentClass returns the C field 'parent_class'.
func (recv *LoggerClass) FieldParentClass() *gobject.ObjectClass {
	argValue := gi.StructFieldGet(loggerClassStruct, recv.Native(), "parent_class")
	value := gobject.ObjectClassNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *LoggerClass) SetFieldParentClass(value *gobject.ObjectClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.StructFieldSet(loggerClassStruct, recv.Native(), "parent_class", argValue)
}

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

	structGo := LoggerClassNewFromNative(loggerClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeLoggerClass)
	return structGo
}
func finalizeLoggerClass(obj *LoggerClass) {
	loggerClassStruct.Free(obj.Native())
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
	native unsafe.Pointer
}

func MessageBodyNewFromNative(native unsafe.Pointer) *MessageBody {
	instance := &MessageBody{native: native}

	return instance
}

/*
CastToMessageBody down casts any arbitrary Object to MessageBody.
Exercise care, as this is a potentially dangerous function
if the Object is not a MessageBody.
*/
func CastToMessageBody(object *gobject.Object) *MessageBody {
	return MessageBodyNewFromNative(object.Native())
}

// Equals compares this MessageBody with another MessageBody, and returns true if they represent the same Object.
func (recv *MessageBody) Equals(other *MessageBody) bool {
	return other.Native() == recv.Native()
}

func (recv *MessageBody) Native() unsafe.Pointer {
	return recv.native
}

// FieldData returns the C field 'data'.
func (recv *MessageBody) FieldData() string {
	argValue := gi.StructFieldGet(messageBodyStruct, recv.Native(), "data")
	value := argValue.String(false)
	return value
}

// SetFieldData sets the value of the C field 'data'.
func (recv *MessageBody) SetFieldData(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.StructFieldSet(messageBodyStruct, recv.Native(), "data", argValue)
}

// FieldLength returns the C field 'length'.
func (recv *MessageBody) FieldLength() int64 {
	argValue := gi.StructFieldGet(messageBodyStruct, recv.Native(), "length")
	value := argValue.Int64()
	return value
}

// SetFieldLength sets the value of the C field 'length'.
func (recv *MessageBody) SetFieldLength(value int64) {
	var argValue gi.Argument
	argValue.SetInt64(value)
	gi.StructFieldSet(messageBodyStruct, recv.Native(), "length", argValue)
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

	retGo := MessageBodyNewFromNative(ret.Pointer())

	return retGo
}

// UNSUPPORTED : C value 'soup_message_body_append' : parameter 'data' of type 'nil' not supported

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
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(buffer.Native())

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
	inArgs[0].SetPointer(recv.Native())

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
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := messageBodyFlattenFunction_Set()
	if err == nil {
		ret = messageBodyFlattenFunction.Invoke(inArgs[:], nil)
	}

	retGo := BufferNewFromNative(ret.Pointer())

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
	inArgs[0].SetPointer(recv.Native())

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
	inArgs[0].SetPointer(recv.Native())

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
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt64(offset)

	var ret gi.Argument

	err := messageBodyGetChunkFunction_Set()
	if err == nil {
		ret = messageBodyGetChunkFunction.Invoke(inArgs[:], nil)
	}

	retGo := BufferNewFromNative(ret.Pointer())

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
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(chunk.Native())

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
	inArgs[0].SetPointer(recv.Native())
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
	inArgs[0].SetPointer(recv.Native())

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
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(chunk.Native())

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
	native unsafe.Pointer
}

func MessageClassNewFromNative(native unsafe.Pointer) *MessageClass {
	instance := &MessageClass{native: native}

	return instance
}

/*
CastToMessageClass down casts any arbitrary Object to MessageClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a MessageClass.
*/
func CastToMessageClass(object *gobject.Object) *MessageClass {
	return MessageClassNewFromNative(object.Native())
}

// Equals compares this MessageClass with another MessageClass, and returns true if they represent the same Object.
func (recv *MessageClass) Equals(other *MessageClass) bool {
	return other.Native() == recv.Native()
}

func (recv *MessageClass) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentClass returns the C field 'parent_class'.
func (recv *MessageClass) FieldParentClass() *gobject.ObjectClass {
	argValue := gi.StructFieldGet(messageClassStruct, recv.Native(), "parent_class")
	value := gobject.ObjectClassNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *MessageClass) SetFieldParentClass(value *gobject.ObjectClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.StructFieldSet(messageClassStruct, recv.Native(), "parent_class", argValue)
}

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

	structGo := MessageClassNewFromNative(messageClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeMessageClass)
	return structGo
}
func finalizeMessageClass(obj *MessageClass) {
	messageClassStruct.Free(obj.Native())
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
	native unsafe.Pointer
}

func MessageHeadersNewFromNative(native unsafe.Pointer) *MessageHeaders {
	instance := &MessageHeaders{native: native}

	return instance
}

/*
CastToMessageHeaders down casts any arbitrary Object to MessageHeaders.
Exercise care, as this is a potentially dangerous function
if the Object is not a MessageHeaders.
*/
func CastToMessageHeaders(object *gobject.Object) *MessageHeaders {
	return MessageHeadersNewFromNative(object.Native())
}

// Equals compares this MessageHeaders with another MessageHeaders, and returns true if they represent the same Object.
func (recv *MessageHeaders) Equals(other *MessageHeaders) bool {
	return other.Native() == recv.Native()
}

func (recv *MessageHeaders) Native() unsafe.Pointer {
	return recv.native
}

var messageHeadersNewFunction *gi.Function
var messageHeadersNewFunction_Once sync.Once

func messageHeadersNewFunction_Set() error {
	var err error
	messageHeadersNewFunction_Once.Do(func() {
		err = messageHeadersStruct_Set()
		if err != nil {
			return
		}
		messageHeadersNewFunction, err = messageHeadersStruct.InvokerNew("new")
	})
	return err
}

// MessageHeadersNew is a representation of the C type soup_message_headers_new.
func MessageHeadersNew(type_ MessageHeadersType) *MessageHeaders {
	var inArgs [1]gi.Argument
	inArgs[0].SetInt32(int32(type_))

	var ret gi.Argument

	err := messageHeadersNewFunction_Set()
	if err == nil {
		ret = messageHeadersNewFunction.Invoke(inArgs[:], nil)
	}

	retGo := MessageHeadersNewFromNative(ret.Pointer())

	return retGo
}

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
	inArgs[0].SetPointer(recv.Native())
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
	inArgs[0].SetPointer(recv.Native())

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
	inArgs[0].SetPointer(recv.Native())

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
	inArgs[0].SetPointer(recv.Native())

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
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(ranges.Native())

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
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(name)

	var ret gi.Argument

	err := messageHeadersGetFunction_Set()
	if err == nil {
		ret = messageHeadersGetFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var messageHeadersGetContentDispositionFunction *gi.Function
var messageHeadersGetContentDispositionFunction_Once sync.Once

func messageHeadersGetContentDispositionFunction_Set() error {
	var err error
	messageHeadersGetContentDispositionFunction_Once.Do(func() {
		err = messageHeadersStruct_Set()
		if err != nil {
			return
		}
		messageHeadersGetContentDispositionFunction, err = messageHeadersStruct.InvokerNew("get_content_disposition")
	})
	return err
}

// GetContentDisposition is a representation of the C type soup_message_headers_get_content_disposition.
func (recv *MessageHeaders) GetContentDisposition() (bool, string, *glib.HashTable) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var outArgs [2]gi.Argument
	var ret gi.Argument

	err := messageHeadersGetContentDispositionFunction_Set()
	if err == nil {
		ret = messageHeadersGetContentDispositionFunction.Invoke(inArgs[:], outArgs[:])
	}

	retGo := ret.Boolean()
	out0 := outArgs[0].String(true)
	out1 := glib.HashTableNewFromNative(outArgs[1].Pointer())

	return retGo, out0, out1
}

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
	inArgs[0].SetPointer(recv.Native())

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
	inArgs[0].SetPointer(recv.Native())

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

var messageHeadersGetContentTypeFunction *gi.Function
var messageHeadersGetContentTypeFunction_Once sync.Once

func messageHeadersGetContentTypeFunction_Set() error {
	var err error
	messageHeadersGetContentTypeFunction_Once.Do(func() {
		err = messageHeadersStruct_Set()
		if err != nil {
			return
		}
		messageHeadersGetContentTypeFunction, err = messageHeadersStruct.InvokerNew("get_content_type")
	})
	return err
}

// GetContentType is a representation of the C type soup_message_headers_get_content_type.
func (recv *MessageHeaders) GetContentType() (string, *glib.HashTable) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var outArgs [1]gi.Argument
	var ret gi.Argument

	err := messageHeadersGetContentTypeFunction_Set()
	if err == nil {
		ret = messageHeadersGetContentTypeFunction.Invoke(inArgs[:], outArgs[:])
	}

	retGo := ret.String(false)
	out0 := glib.HashTableNewFromNative(outArgs[0].Pointer())

	return retGo, out0
}

var messageHeadersGetEncodingFunction *gi.Function
var messageHeadersGetEncodingFunction_Once sync.Once

func messageHeadersGetEncodingFunction_Set() error {
	var err error
	messageHeadersGetEncodingFunction_Once.Do(func() {
		err = messageHeadersStruct_Set()
		if err != nil {
			return
		}
		messageHeadersGetEncodingFunction, err = messageHeadersStruct.InvokerNew("get_encoding")
	})
	return err
}

// GetEncoding is a representation of the C type soup_message_headers_get_encoding.
func (recv *MessageHeaders) GetEncoding() Encoding {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := messageHeadersGetEncodingFunction_Set()
	if err == nil {
		ret = messageHeadersGetEncodingFunction.Invoke(inArgs[:], nil)
	}

	retGo := Encoding(ret.Int32())

	return retGo
}

var messageHeadersGetExpectationsFunction *gi.Function
var messageHeadersGetExpectationsFunction_Once sync.Once

func messageHeadersGetExpectationsFunction_Set() error {
	var err error
	messageHeadersGetExpectationsFunction_Once.Do(func() {
		err = messageHeadersStruct_Set()
		if err != nil {
			return
		}
		messageHeadersGetExpectationsFunction, err = messageHeadersStruct.InvokerNew("get_expectations")
	})
	return err
}

// GetExpectations is a representation of the C type soup_message_headers_get_expectations.
func (recv *MessageHeaders) GetExpectations() Expectation {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := messageHeadersGetExpectationsFunction_Set()
	if err == nil {
		ret = messageHeadersGetExpectationsFunction.Invoke(inArgs[:], nil)
	}

	retGo := Expectation(ret.Int32())

	return retGo
}

var messageHeadersGetHeadersTypeFunction *gi.Function
var messageHeadersGetHeadersTypeFunction_Once sync.Once

func messageHeadersGetHeadersTypeFunction_Set() error {
	var err error
	messageHeadersGetHeadersTypeFunction_Once.Do(func() {
		err = messageHeadersStruct_Set()
		if err != nil {
			return
		}
		messageHeadersGetHeadersTypeFunction, err = messageHeadersStruct.InvokerNew("get_headers_type")
	})
	return err
}

// GetHeadersType is a representation of the C type soup_message_headers_get_headers_type.
func (recv *MessageHeaders) GetHeadersType() MessageHeadersType {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := messageHeadersGetHeadersTypeFunction_Set()
	if err == nil {
		ret = messageHeadersGetHeadersTypeFunction.Invoke(inArgs[:], nil)
	}

	retGo := MessageHeadersType(ret.Int32())

	return retGo
}

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
	inArgs[0].SetPointer(recv.Native())
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
	inArgs[0].SetPointer(recv.Native())
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
	inArgs[0].SetPointer(recv.Native())
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
	inArgs[0].SetPointer(recv.Native())
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
	inArgs[0].SetPointer(recv.Native())
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
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(name)
	inArgs[2].SetString(value)

	err := messageHeadersReplaceFunction_Set()
	if err == nil {
		messageHeadersReplaceFunction.Invoke(inArgs[:], nil)
	}

	return
}

var messageHeadersSetContentDispositionFunction *gi.Function
var messageHeadersSetContentDispositionFunction_Once sync.Once

func messageHeadersSetContentDispositionFunction_Set() error {
	var err error
	messageHeadersSetContentDispositionFunction_Once.Do(func() {
		err = messageHeadersStruct_Set()
		if err != nil {
			return
		}
		messageHeadersSetContentDispositionFunction, err = messageHeadersStruct.InvokerNew("set_content_disposition")
	})
	return err
}

// SetContentDisposition is a representation of the C type soup_message_headers_set_content_disposition.
func (recv *MessageHeaders) SetContentDisposition(disposition string, params *glib.HashTable) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(disposition)
	inArgs[2].SetPointer(params.Native())

	err := messageHeadersSetContentDispositionFunction_Set()
	if err == nil {
		messageHeadersSetContentDispositionFunction.Invoke(inArgs[:], nil)
	}

	return
}

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
	inArgs[0].SetPointer(recv.Native())
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
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt64(start)
	inArgs[2].SetInt64(end)
	inArgs[3].SetInt64(totalLength)

	err := messageHeadersSetContentRangeFunction_Set()
	if err == nil {
		messageHeadersSetContentRangeFunction.Invoke(inArgs[:], nil)
	}

	return
}

var messageHeadersSetContentTypeFunction *gi.Function
var messageHeadersSetContentTypeFunction_Once sync.Once

func messageHeadersSetContentTypeFunction_Set() error {
	var err error
	messageHeadersSetContentTypeFunction_Once.Do(func() {
		err = messageHeadersStruct_Set()
		if err != nil {
			return
		}
		messageHeadersSetContentTypeFunction, err = messageHeadersStruct.InvokerNew("set_content_type")
	})
	return err
}

// SetContentType is a representation of the C type soup_message_headers_set_content_type.
func (recv *MessageHeaders) SetContentType(contentType string, params *glib.HashTable) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(contentType)
	inArgs[2].SetPointer(params.Native())

	err := messageHeadersSetContentTypeFunction_Set()
	if err == nil {
		messageHeadersSetContentTypeFunction.Invoke(inArgs[:], nil)
	}

	return
}

var messageHeadersSetEncodingFunction *gi.Function
var messageHeadersSetEncodingFunction_Once sync.Once

func messageHeadersSetEncodingFunction_Set() error {
	var err error
	messageHeadersSetEncodingFunction_Once.Do(func() {
		err = messageHeadersStruct_Set()
		if err != nil {
			return
		}
		messageHeadersSetEncodingFunction, err = messageHeadersStruct.InvokerNew("set_encoding")
	})
	return err
}

// SetEncoding is a representation of the C type soup_message_headers_set_encoding.
func (recv *MessageHeaders) SetEncoding(encoding Encoding) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(int32(encoding))

	err := messageHeadersSetEncodingFunction_Set()
	if err == nil {
		messageHeadersSetEncodingFunction.Invoke(inArgs[:], nil)
	}

	return
}

var messageHeadersSetExpectationsFunction *gi.Function
var messageHeadersSetExpectationsFunction_Once sync.Once

func messageHeadersSetExpectationsFunction_Set() error {
	var err error
	messageHeadersSetExpectationsFunction_Once.Do(func() {
		err = messageHeadersStruct_Set()
		if err != nil {
			return
		}
		messageHeadersSetExpectationsFunction, err = messageHeadersStruct.InvokerNew("set_expectations")
	})
	return err
}

// SetExpectations is a representation of the C type soup_message_headers_set_expectations.
func (recv *MessageHeaders) SetExpectations(expectations Expectation) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(int32(expectations))

	err := messageHeadersSetExpectationsFunction_Set()
	if err == nil {
		messageHeadersSetExpectationsFunction.Invoke(inArgs[:], nil)
	}

	return
}

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
	inArgs[0].SetPointer(recv.Native())
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
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(ranges.Native())
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
	native unsafe.Pointer
}

func MessageHeadersIterNewFromNative(native unsafe.Pointer) *MessageHeadersIter {
	instance := &MessageHeadersIter{native: native}

	return instance
}

/*
CastToMessageHeadersIter down casts any arbitrary Object to MessageHeadersIter.
Exercise care, as this is a potentially dangerous function
if the Object is not a MessageHeadersIter.
*/
func CastToMessageHeadersIter(object *gobject.Object) *MessageHeadersIter {
	return MessageHeadersIterNewFromNative(object.Native())
}

// Equals compares this MessageHeadersIter with another MessageHeadersIter, and returns true if they represent the same Object.
func (recv *MessageHeadersIter) Equals(other *MessageHeadersIter) bool {
	return other.Native() == recv.Native()
}

func (recv *MessageHeadersIter) Native() unsafe.Pointer {
	return recv.native
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
	inArgs[0].SetPointer(recv.Native())

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

	structGo := MessageHeadersIterNewFromNative(messageHeadersIterStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeMessageHeadersIter)
	return structGo
}
func finalizeMessageHeadersIter(obj *MessageHeadersIter) {
	messageHeadersIterStruct.Free(obj.Native())
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
	native unsafe.Pointer
}

func MessageQueueNewFromNative(native unsafe.Pointer) *MessageQueue {
	instance := &MessageQueue{native: native}

	return instance
}

/*
CastToMessageQueue down casts any arbitrary Object to MessageQueue.
Exercise care, as this is a potentially dangerous function
if the Object is not a MessageQueue.
*/
func CastToMessageQueue(object *gobject.Object) *MessageQueue {
	return MessageQueueNewFromNative(object.Native())
}

// Equals compares this MessageQueue with another MessageQueue, and returns true if they represent the same Object.
func (recv *MessageQueue) Equals(other *MessageQueue) bool {
	return other.Native() == recv.Native()
}

func (recv *MessageQueue) Native() unsafe.Pointer {
	return recv.native
}

// MessageQueueStruct creates an uninitialised MessageQueue.
func MessageQueueStruct() *MessageQueue {
	err := messageQueueStruct_Set()
	if err != nil {
		return nil
	}

	structGo := MessageQueueNewFromNative(messageQueueStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeMessageQueue)
	return structGo
}
func finalizeMessageQueue(obj *MessageQueue) {
	messageQueueStruct.Free(obj.Native())
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
	native unsafe.Pointer
}

func MessageQueueItemNewFromNative(native unsafe.Pointer) *MessageQueueItem {
	instance := &MessageQueueItem{native: native}

	return instance
}

/*
CastToMessageQueueItem down casts any arbitrary Object to MessageQueueItem.
Exercise care, as this is a potentially dangerous function
if the Object is not a MessageQueueItem.
*/
func CastToMessageQueueItem(object *gobject.Object) *MessageQueueItem {
	return MessageQueueItemNewFromNative(object.Native())
}

// Equals compares this MessageQueueItem with another MessageQueueItem, and returns true if they represent the same Object.
func (recv *MessageQueueItem) Equals(other *MessageQueueItem) bool {
	return other.Native() == recv.Native()
}

func (recv *MessageQueueItem) Native() unsafe.Pointer {
	return recv.native
}

// MessageQueueItemStruct creates an uninitialised MessageQueueItem.
func MessageQueueItemStruct() *MessageQueueItem {
	err := messageQueueItemStruct_Set()
	if err != nil {
		return nil
	}

	structGo := MessageQueueItemNewFromNative(messageQueueItemStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeMessageQueueItem)
	return structGo
}
func finalizeMessageQueueItem(obj *MessageQueueItem) {
	messageQueueItemStruct.Free(obj.Native())
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
	native unsafe.Pointer
}

func MultipartNewFromNative(native unsafe.Pointer) *Multipart {
	instance := &Multipart{native: native}

	return instance
}

/*
CastToMultipart down casts any arbitrary Object to Multipart.
Exercise care, as this is a potentially dangerous function
if the Object is not a Multipart.
*/
func CastToMultipart(object *gobject.Object) *Multipart {
	return MultipartNewFromNative(object.Native())
}

// Equals compares this Multipart with another Multipart, and returns true if they represent the same Object.
func (recv *Multipart) Equals(other *Multipart) bool {
	return other.Native() == recv.Native()
}

func (recv *Multipart) Native() unsafe.Pointer {
	return recv.native
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

	retGo := MultipartNewFromNative(ret.Pointer())

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
	inArgs[0].SetPointer(headers.Native())
	inArgs[1].SetPointer(body.Native())

	var ret gi.Argument

	err := multipartNewFromMessageFunction_Set()
	if err == nil {
		ret = multipartNewFromMessageFunction.Invoke(inArgs[:], nil)
	}

	retGo := MultipartNewFromNative(ret.Pointer())

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
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(controlName)
	inArgs[2].SetString(filename)
	inArgs[3].SetString(contentType)
	inArgs[4].SetPointer(body.Native())

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
	inArgs[0].SetPointer(recv.Native())
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
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(headers.Native())
	inArgs[2].SetPointer(body.Native())

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
	inArgs[0].SetPointer(recv.Native())

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
	inArgs[0].SetPointer(recv.Native())

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
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(part)

	var outArgs [2]gi.Argument
	var ret gi.Argument

	err := multipartGetPartFunction_Set()
	if err == nil {
		ret = multipartGetPartFunction.Invoke(inArgs[:], outArgs[:])
	}

	retGo := ret.Boolean()
	out0 := MessageHeadersNewFromNative(outArgs[0].Pointer())
	out1 := BufferNewFromNative(outArgs[1].Pointer())

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
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(destHeaders.Native())
	inArgs[2].SetPointer(destBody.Native())

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
	native unsafe.Pointer
}

func MultipartInputStreamClassNewFromNative(native unsafe.Pointer) *MultipartInputStreamClass {
	instance := &MultipartInputStreamClass{native: native}

	return instance
}

/*
CastToMultipartInputStreamClass down casts any arbitrary Object to MultipartInputStreamClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a MultipartInputStreamClass.
*/
func CastToMultipartInputStreamClass(object *gobject.Object) *MultipartInputStreamClass {
	return MultipartInputStreamClassNewFromNative(object.Native())
}

// Equals compares this MultipartInputStreamClass with another MultipartInputStreamClass, and returns true if they represent the same Object.
func (recv *MultipartInputStreamClass) Equals(other *MultipartInputStreamClass) bool {
	return other.Native() == recv.Native()
}

func (recv *MultipartInputStreamClass) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentClass returns the C field 'parent_class'.
func (recv *MultipartInputStreamClass) FieldParentClass() *gio.FilterInputStreamClass {
	argValue := gi.StructFieldGet(multipartInputStreamClassStruct, recv.Native(), "parent_class")
	value := gio.FilterInputStreamClassNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *MultipartInputStreamClass) SetFieldParentClass(value *gio.FilterInputStreamClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.StructFieldSet(multipartInputStreamClassStruct, recv.Native(), "parent_class", argValue)
}

// MultipartInputStreamClassStruct creates an uninitialised MultipartInputStreamClass.
func MultipartInputStreamClassStruct() *MultipartInputStreamClass {
	err := multipartInputStreamClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := MultipartInputStreamClassNewFromNative(multipartInputStreamClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeMultipartInputStreamClass)
	return structGo
}
func finalizeMultipartInputStreamClass(obj *MultipartInputStreamClass) {
	multipartInputStreamClassStruct.Free(obj.Native())
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
	native unsafe.Pointer
}

func MultipartInputStreamPrivateNewFromNative(native unsafe.Pointer) *MultipartInputStreamPrivate {
	instance := &MultipartInputStreamPrivate{native: native}

	return instance
}

/*
CastToMultipartInputStreamPrivate down casts any arbitrary Object to MultipartInputStreamPrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a MultipartInputStreamPrivate.
*/
func CastToMultipartInputStreamPrivate(object *gobject.Object) *MultipartInputStreamPrivate {
	return MultipartInputStreamPrivateNewFromNative(object.Native())
}

// Equals compares this MultipartInputStreamPrivate with another MultipartInputStreamPrivate, and returns true if they represent the same Object.
func (recv *MultipartInputStreamPrivate) Equals(other *MultipartInputStreamPrivate) bool {
	return other.Native() == recv.Native()
}

func (recv *MultipartInputStreamPrivate) Native() unsafe.Pointer {
	return recv.native
}

// MultipartInputStreamPrivateStruct creates an uninitialised MultipartInputStreamPrivate.
func MultipartInputStreamPrivateStruct() *MultipartInputStreamPrivate {
	err := multipartInputStreamPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := MultipartInputStreamPrivateNewFromNative(multipartInputStreamPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeMultipartInputStreamPrivate)
	return structGo
}
func finalizeMultipartInputStreamPrivate(obj *MultipartInputStreamPrivate) {
	multipartInputStreamPrivateStruct.Free(obj.Native())
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
	native unsafe.Pointer
}

func PasswordManagerInterfaceNewFromNative(native unsafe.Pointer) *PasswordManagerInterface {
	instance := &PasswordManagerInterface{native: native}

	return instance
}

/*
CastToPasswordManagerInterface down casts any arbitrary Object to PasswordManagerInterface.
Exercise care, as this is a potentially dangerous function
if the Object is not a PasswordManagerInterface.
*/
func CastToPasswordManagerInterface(object *gobject.Object) *PasswordManagerInterface {
	return PasswordManagerInterfaceNewFromNative(object.Native())
}

// Equals compares this PasswordManagerInterface with another PasswordManagerInterface, and returns true if they represent the same Object.
func (recv *PasswordManagerInterface) Equals(other *PasswordManagerInterface) bool {
	return other.Native() == recv.Native()
}

func (recv *PasswordManagerInterface) Native() unsafe.Pointer {
	return recv.native
}

// FieldBase returns the C field 'base'.
func (recv *PasswordManagerInterface) FieldBase() *gobject.TypeInterface {
	argValue := gi.StructFieldGet(passwordManagerInterfaceStruct, recv.Native(), "base")
	value := gobject.TypeInterfaceNewFromNative(argValue.Pointer())
	return value
}

// SetFieldBase sets the value of the C field 'base'.
func (recv *PasswordManagerInterface) SetFieldBase(value *gobject.TypeInterface) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.StructFieldSet(passwordManagerInterfaceStruct, recv.Native(), "base", argValue)
}

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

	structGo := PasswordManagerInterfaceNewFromNative(passwordManagerInterfaceStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizePasswordManagerInterface)
	return structGo
}
func finalizePasswordManagerInterface(obj *PasswordManagerInterface) {
	passwordManagerInterfaceStruct.Free(obj.Native())
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
	native unsafe.Pointer
}

func ProxyResolverDefaultClassNewFromNative(native unsafe.Pointer) *ProxyResolverDefaultClass {
	instance := &ProxyResolverDefaultClass{native: native}

	return instance
}

/*
CastToProxyResolverDefaultClass down casts any arbitrary Object to ProxyResolverDefaultClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a ProxyResolverDefaultClass.
*/
func CastToProxyResolverDefaultClass(object *gobject.Object) *ProxyResolverDefaultClass {
	return ProxyResolverDefaultClassNewFromNative(object.Native())
}

// Equals compares this ProxyResolverDefaultClass with another ProxyResolverDefaultClass, and returns true if they represent the same Object.
func (recv *ProxyResolverDefaultClass) Equals(other *ProxyResolverDefaultClass) bool {
	return other.Native() == recv.Native()
}

func (recv *ProxyResolverDefaultClass) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentClass returns the C field 'parent_class'.
func (recv *ProxyResolverDefaultClass) FieldParentClass() *gobject.ObjectClass {
	argValue := gi.StructFieldGet(proxyResolverDefaultClassStruct, recv.Native(), "parent_class")
	value := gobject.ObjectClassNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *ProxyResolverDefaultClass) SetFieldParentClass(value *gobject.ObjectClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.StructFieldSet(proxyResolverDefaultClassStruct, recv.Native(), "parent_class", argValue)
}

// ProxyResolverDefaultClassStruct creates an uninitialised ProxyResolverDefaultClass.
func ProxyResolverDefaultClassStruct() *ProxyResolverDefaultClass {
	err := proxyResolverDefaultClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := ProxyResolverDefaultClassNewFromNative(proxyResolverDefaultClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeProxyResolverDefaultClass)
	return structGo
}
func finalizeProxyResolverDefaultClass(obj *ProxyResolverDefaultClass) {
	proxyResolverDefaultClassStruct.Free(obj.Native())
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
	native unsafe.Pointer
}

func ProxyResolverInterfaceNewFromNative(native unsafe.Pointer) *ProxyResolverInterface {
	instance := &ProxyResolverInterface{native: native}

	return instance
}

/*
CastToProxyResolverInterface down casts any arbitrary Object to ProxyResolverInterface.
Exercise care, as this is a potentially dangerous function
if the Object is not a ProxyResolverInterface.
*/
func CastToProxyResolverInterface(object *gobject.Object) *ProxyResolverInterface {
	return ProxyResolverInterfaceNewFromNative(object.Native())
}

// Equals compares this ProxyResolverInterface with another ProxyResolverInterface, and returns true if they represent the same Object.
func (recv *ProxyResolverInterface) Equals(other *ProxyResolverInterface) bool {
	return other.Native() == recv.Native()
}

func (recv *ProxyResolverInterface) Native() unsafe.Pointer {
	return recv.native
}

// FieldBase returns the C field 'base'.
func (recv *ProxyResolverInterface) FieldBase() *gobject.TypeInterface {
	argValue := gi.StructFieldGet(proxyResolverInterfaceStruct, recv.Native(), "base")
	value := gobject.TypeInterfaceNewFromNative(argValue.Pointer())
	return value
}

// SetFieldBase sets the value of the C field 'base'.
func (recv *ProxyResolverInterface) SetFieldBase(value *gobject.TypeInterface) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.StructFieldSet(proxyResolverInterfaceStruct, recv.Native(), "base", argValue)
}

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

	structGo := ProxyResolverInterfaceNewFromNative(proxyResolverInterfaceStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeProxyResolverInterface)
	return structGo
}
func finalizeProxyResolverInterface(obj *ProxyResolverInterface) {
	proxyResolverInterfaceStruct.Free(obj.Native())
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
	native unsafe.Pointer
}

func ProxyURIResolverInterfaceNewFromNative(native unsafe.Pointer) *ProxyURIResolverInterface {
	instance := &ProxyURIResolverInterface{native: native}

	return instance
}

/*
CastToProxyURIResolverInterface down casts any arbitrary Object to ProxyURIResolverInterface.
Exercise care, as this is a potentially dangerous function
if the Object is not a ProxyURIResolverInterface.
*/
func CastToProxyURIResolverInterface(object *gobject.Object) *ProxyURIResolverInterface {
	return ProxyURIResolverInterfaceNewFromNative(object.Native())
}

// Equals compares this ProxyURIResolverInterface with another ProxyURIResolverInterface, and returns true if they represent the same Object.
func (recv *ProxyURIResolverInterface) Equals(other *ProxyURIResolverInterface) bool {
	return other.Native() == recv.Native()
}

func (recv *ProxyURIResolverInterface) Native() unsafe.Pointer {
	return recv.native
}

// FieldBase returns the C field 'base'.
func (recv *ProxyURIResolverInterface) FieldBase() *gobject.TypeInterface {
	argValue := gi.StructFieldGet(proxyURIResolverInterfaceStruct, recv.Native(), "base")
	value := gobject.TypeInterfaceNewFromNative(argValue.Pointer())
	return value
}

// SetFieldBase sets the value of the C field 'base'.
func (recv *ProxyURIResolverInterface) SetFieldBase(value *gobject.TypeInterface) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.StructFieldSet(proxyURIResolverInterfaceStruct, recv.Native(), "base", argValue)
}

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

	structGo := ProxyURIResolverInterfaceNewFromNative(proxyURIResolverInterfaceStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeProxyURIResolverInterface)
	return structGo
}
func finalizeProxyURIResolverInterface(obj *ProxyURIResolverInterface) {
	proxyURIResolverInterfaceStruct.Free(obj.Native())
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
	native unsafe.Pointer
}

func RangeNewFromNative(native unsafe.Pointer) *Range {
	instance := &Range{native: native}

	return instance
}

/*
CastToRange down casts any arbitrary Object to Range.
Exercise care, as this is a potentially dangerous function
if the Object is not a Range.
*/
func CastToRange(object *gobject.Object) *Range {
	return RangeNewFromNative(object.Native())
}

// Equals compares this Range with another Range, and returns true if they represent the same Object.
func (recv *Range) Equals(other *Range) bool {
	return other.Native() == recv.Native()
}

func (recv *Range) Native() unsafe.Pointer {
	return recv.native
}

// FieldStart returns the C field 'start'.
func (recv *Range) FieldStart() int64 {
	argValue := gi.StructFieldGet(rangeStruct, recv.Native(), "start")
	value := argValue.Int64()
	return value
}

// SetFieldStart sets the value of the C field 'start'.
func (recv *Range) SetFieldStart(value int64) {
	var argValue gi.Argument
	argValue.SetInt64(value)
	gi.StructFieldSet(rangeStruct, recv.Native(), "start", argValue)
}

// FieldEnd returns the C field 'end'.
func (recv *Range) FieldEnd() int64 {
	argValue := gi.StructFieldGet(rangeStruct, recv.Native(), "end")
	value := argValue.Int64()
	return value
}

// SetFieldEnd sets the value of the C field 'end'.
func (recv *Range) SetFieldEnd(value int64) {
	var argValue gi.Argument
	argValue.SetInt64(value)
	gi.StructFieldSet(rangeStruct, recv.Native(), "end", argValue)
}

// RangeStruct creates an uninitialised Range.
func RangeStruct() *Range {
	err := rangeStruct_Set()
	if err != nil {
		return nil
	}

	structGo := RangeNewFromNative(rangeStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeRange)
	return structGo
}
func finalizeRange(obj *Range) {
	rangeStruct.Free(obj.Native())
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
	native unsafe.Pointer
}

func RequestClassNewFromNative(native unsafe.Pointer) *RequestClass {
	instance := &RequestClass{native: native}

	return instance
}

/*
CastToRequestClass down casts any arbitrary Object to RequestClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a RequestClass.
*/
func CastToRequestClass(object *gobject.Object) *RequestClass {
	return RequestClassNewFromNative(object.Native())
}

// Equals compares this RequestClass with another RequestClass, and returns true if they represent the same Object.
func (recv *RequestClass) Equals(other *RequestClass) bool {
	return other.Native() == recv.Native()
}

func (recv *RequestClass) Native() unsafe.Pointer {
	return recv.native
}

// FieldParent returns the C field 'parent'.
func (recv *RequestClass) FieldParent() *gobject.ObjectClass {
	argValue := gi.StructFieldGet(requestClassStruct, recv.Native(), "parent")
	value := gobject.ObjectClassNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParent sets the value of the C field 'parent'.
func (recv *RequestClass) SetFieldParent(value *gobject.ObjectClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.StructFieldSet(requestClassStruct, recv.Native(), "parent", argValue)
}

// FieldSchemes returns the C field 'schemes'.
func (recv *RequestClass) FieldSchemes() string {
	argValue := gi.StructFieldGet(requestClassStruct, recv.Native(), "schemes")
	value := argValue.String(false)
	return value
}

// SetFieldSchemes sets the value of the C field 'schemes'.
func (recv *RequestClass) SetFieldSchemes(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.StructFieldSet(requestClassStruct, recv.Native(), "schemes", argValue)
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

	structGo := RequestClassNewFromNative(requestClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeRequestClass)
	return structGo
}
func finalizeRequestClass(obj *RequestClass) {
	requestClassStruct.Free(obj.Native())
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
	native unsafe.Pointer
}

func RequestDataClassNewFromNative(native unsafe.Pointer) *RequestDataClass {
	instance := &RequestDataClass{native: native}

	return instance
}

/*
CastToRequestDataClass down casts any arbitrary Object to RequestDataClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a RequestDataClass.
*/
func CastToRequestDataClass(object *gobject.Object) *RequestDataClass {
	return RequestDataClassNewFromNative(object.Native())
}

// Equals compares this RequestDataClass with another RequestDataClass, and returns true if they represent the same Object.
func (recv *RequestDataClass) Equals(other *RequestDataClass) bool {
	return other.Native() == recv.Native()
}

func (recv *RequestDataClass) Native() unsafe.Pointer {
	return recv.native
}

// FieldParent returns the C field 'parent'.
func (recv *RequestDataClass) FieldParent() *RequestClass {
	argValue := gi.StructFieldGet(requestDataClassStruct, recv.Native(), "parent")
	value := RequestClassNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParent sets the value of the C field 'parent'.
func (recv *RequestDataClass) SetFieldParent(value *RequestClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.StructFieldSet(requestDataClassStruct, recv.Native(), "parent", argValue)
}

// RequestDataClassStruct creates an uninitialised RequestDataClass.
func RequestDataClassStruct() *RequestDataClass {
	err := requestDataClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := RequestDataClassNewFromNative(requestDataClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeRequestDataClass)
	return structGo
}
func finalizeRequestDataClass(obj *RequestDataClass) {
	requestDataClassStruct.Free(obj.Native())
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
	native unsafe.Pointer
}

func RequestDataPrivateNewFromNative(native unsafe.Pointer) *RequestDataPrivate {
	instance := &RequestDataPrivate{native: native}

	return instance
}

/*
CastToRequestDataPrivate down casts any arbitrary Object to RequestDataPrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a RequestDataPrivate.
*/
func CastToRequestDataPrivate(object *gobject.Object) *RequestDataPrivate {
	return RequestDataPrivateNewFromNative(object.Native())
}

// Equals compares this RequestDataPrivate with another RequestDataPrivate, and returns true if they represent the same Object.
func (recv *RequestDataPrivate) Equals(other *RequestDataPrivate) bool {
	return other.Native() == recv.Native()
}

func (recv *RequestDataPrivate) Native() unsafe.Pointer {
	return recv.native
}

// RequestDataPrivateStruct creates an uninitialised RequestDataPrivate.
func RequestDataPrivateStruct() *RequestDataPrivate {
	err := requestDataPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := RequestDataPrivateNewFromNative(requestDataPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeRequestDataPrivate)
	return structGo
}
func finalizeRequestDataPrivate(obj *RequestDataPrivate) {
	requestDataPrivateStruct.Free(obj.Native())
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
	native unsafe.Pointer
}

func RequestFileClassNewFromNative(native unsafe.Pointer) *RequestFileClass {
	instance := &RequestFileClass{native: native}

	return instance
}

/*
CastToRequestFileClass down casts any arbitrary Object to RequestFileClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a RequestFileClass.
*/
func CastToRequestFileClass(object *gobject.Object) *RequestFileClass {
	return RequestFileClassNewFromNative(object.Native())
}

// Equals compares this RequestFileClass with another RequestFileClass, and returns true if they represent the same Object.
func (recv *RequestFileClass) Equals(other *RequestFileClass) bool {
	return other.Native() == recv.Native()
}

func (recv *RequestFileClass) Native() unsafe.Pointer {
	return recv.native
}

// FieldParent returns the C field 'parent'.
func (recv *RequestFileClass) FieldParent() *RequestClass {
	argValue := gi.StructFieldGet(requestFileClassStruct, recv.Native(), "parent")
	value := RequestClassNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParent sets the value of the C field 'parent'.
func (recv *RequestFileClass) SetFieldParent(value *RequestClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.StructFieldSet(requestFileClassStruct, recv.Native(), "parent", argValue)
}

// RequestFileClassStruct creates an uninitialised RequestFileClass.
func RequestFileClassStruct() *RequestFileClass {
	err := requestFileClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := RequestFileClassNewFromNative(requestFileClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeRequestFileClass)
	return structGo
}
func finalizeRequestFileClass(obj *RequestFileClass) {
	requestFileClassStruct.Free(obj.Native())
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
	native unsafe.Pointer
}

func RequestFilePrivateNewFromNative(native unsafe.Pointer) *RequestFilePrivate {
	instance := &RequestFilePrivate{native: native}

	return instance
}

/*
CastToRequestFilePrivate down casts any arbitrary Object to RequestFilePrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a RequestFilePrivate.
*/
func CastToRequestFilePrivate(object *gobject.Object) *RequestFilePrivate {
	return RequestFilePrivateNewFromNative(object.Native())
}

// Equals compares this RequestFilePrivate with another RequestFilePrivate, and returns true if they represent the same Object.
func (recv *RequestFilePrivate) Equals(other *RequestFilePrivate) bool {
	return other.Native() == recv.Native()
}

func (recv *RequestFilePrivate) Native() unsafe.Pointer {
	return recv.native
}

// RequestFilePrivateStruct creates an uninitialised RequestFilePrivate.
func RequestFilePrivateStruct() *RequestFilePrivate {
	err := requestFilePrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := RequestFilePrivateNewFromNative(requestFilePrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeRequestFilePrivate)
	return structGo
}
func finalizeRequestFilePrivate(obj *RequestFilePrivate) {
	requestFilePrivateStruct.Free(obj.Native())
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
	native unsafe.Pointer
}

func RequestHTTPClassNewFromNative(native unsafe.Pointer) *RequestHTTPClass {
	instance := &RequestHTTPClass{native: native}

	return instance
}

/*
CastToRequestHTTPClass down casts any arbitrary Object to RequestHTTPClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a RequestHTTPClass.
*/
func CastToRequestHTTPClass(object *gobject.Object) *RequestHTTPClass {
	return RequestHTTPClassNewFromNative(object.Native())
}

// Equals compares this RequestHTTPClass with another RequestHTTPClass, and returns true if they represent the same Object.
func (recv *RequestHTTPClass) Equals(other *RequestHTTPClass) bool {
	return other.Native() == recv.Native()
}

func (recv *RequestHTTPClass) Native() unsafe.Pointer {
	return recv.native
}

// FieldParent returns the C field 'parent'.
func (recv *RequestHTTPClass) FieldParent() *RequestClass {
	argValue := gi.StructFieldGet(requestHTTPClassStruct, recv.Native(), "parent")
	value := RequestClassNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParent sets the value of the C field 'parent'.
func (recv *RequestHTTPClass) SetFieldParent(value *RequestClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.StructFieldSet(requestHTTPClassStruct, recv.Native(), "parent", argValue)
}

// RequestHTTPClassStruct creates an uninitialised RequestHTTPClass.
func RequestHTTPClassStruct() *RequestHTTPClass {
	err := requestHTTPClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := RequestHTTPClassNewFromNative(requestHTTPClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeRequestHTTPClass)
	return structGo
}
func finalizeRequestHTTPClass(obj *RequestHTTPClass) {
	requestHTTPClassStruct.Free(obj.Native())
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
	native unsafe.Pointer
}

func RequestHTTPPrivateNewFromNative(native unsafe.Pointer) *RequestHTTPPrivate {
	instance := &RequestHTTPPrivate{native: native}

	return instance
}

/*
CastToRequestHTTPPrivate down casts any arbitrary Object to RequestHTTPPrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a RequestHTTPPrivate.
*/
func CastToRequestHTTPPrivate(object *gobject.Object) *RequestHTTPPrivate {
	return RequestHTTPPrivateNewFromNative(object.Native())
}

// Equals compares this RequestHTTPPrivate with another RequestHTTPPrivate, and returns true if they represent the same Object.
func (recv *RequestHTTPPrivate) Equals(other *RequestHTTPPrivate) bool {
	return other.Native() == recv.Native()
}

func (recv *RequestHTTPPrivate) Native() unsafe.Pointer {
	return recv.native
}

// RequestHTTPPrivateStruct creates an uninitialised RequestHTTPPrivate.
func RequestHTTPPrivateStruct() *RequestHTTPPrivate {
	err := requestHTTPPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := RequestHTTPPrivateNewFromNative(requestHTTPPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeRequestHTTPPrivate)
	return structGo
}
func finalizeRequestHTTPPrivate(obj *RequestHTTPPrivate) {
	requestHTTPPrivateStruct.Free(obj.Native())
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
	native unsafe.Pointer
}

func RequestPrivateNewFromNative(native unsafe.Pointer) *RequestPrivate {
	instance := &RequestPrivate{native: native}

	return instance
}

/*
CastToRequestPrivate down casts any arbitrary Object to RequestPrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a RequestPrivate.
*/
func CastToRequestPrivate(object *gobject.Object) *RequestPrivate {
	return RequestPrivateNewFromNative(object.Native())
}

// Equals compares this RequestPrivate with another RequestPrivate, and returns true if they represent the same Object.
func (recv *RequestPrivate) Equals(other *RequestPrivate) bool {
	return other.Native() == recv.Native()
}

func (recv *RequestPrivate) Native() unsafe.Pointer {
	return recv.native
}

// RequestPrivateStruct creates an uninitialised RequestPrivate.
func RequestPrivateStruct() *RequestPrivate {
	err := requestPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := RequestPrivateNewFromNative(requestPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeRequestPrivate)
	return structGo
}
func finalizeRequestPrivate(obj *RequestPrivate) {
	requestPrivateStruct.Free(obj.Native())
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
	native unsafe.Pointer
}

func RequesterClassNewFromNative(native unsafe.Pointer) *RequesterClass {
	instance := &RequesterClass{native: native}

	return instance
}

/*
CastToRequesterClass down casts any arbitrary Object to RequesterClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a RequesterClass.
*/
func CastToRequesterClass(object *gobject.Object) *RequesterClass {
	return RequesterClassNewFromNative(object.Native())
}

// Equals compares this RequesterClass with another RequesterClass, and returns true if they represent the same Object.
func (recv *RequesterClass) Equals(other *RequesterClass) bool {
	return other.Native() == recv.Native()
}

func (recv *RequesterClass) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentClass returns the C field 'parent_class'.
func (recv *RequesterClass) FieldParentClass() *gobject.ObjectClass {
	argValue := gi.StructFieldGet(requesterClassStruct, recv.Native(), "parent_class")
	value := gobject.ObjectClassNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *RequesterClass) SetFieldParentClass(value *gobject.ObjectClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.StructFieldSet(requesterClassStruct, recv.Native(), "parent_class", argValue)
}

// RequesterClassStruct creates an uninitialised RequesterClass.
func RequesterClassStruct() *RequesterClass {
	err := requesterClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := RequesterClassNewFromNative(requesterClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeRequesterClass)
	return structGo
}
func finalizeRequesterClass(obj *RequesterClass) {
	requesterClassStruct.Free(obj.Native())
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
	native unsafe.Pointer
}

func RequesterPrivateNewFromNative(native unsafe.Pointer) *RequesterPrivate {
	instance := &RequesterPrivate{native: native}

	return instance
}

/*
CastToRequesterPrivate down casts any arbitrary Object to RequesterPrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a RequesterPrivate.
*/
func CastToRequesterPrivate(object *gobject.Object) *RequesterPrivate {
	return RequesterPrivateNewFromNative(object.Native())
}

// Equals compares this RequesterPrivate with another RequesterPrivate, and returns true if they represent the same Object.
func (recv *RequesterPrivate) Equals(other *RequesterPrivate) bool {
	return other.Native() == recv.Native()
}

func (recv *RequesterPrivate) Native() unsafe.Pointer {
	return recv.native
}

// RequesterPrivateStruct creates an uninitialised RequesterPrivate.
func RequesterPrivateStruct() *RequesterPrivate {
	err := requesterPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := RequesterPrivateNewFromNative(requesterPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeRequesterPrivate)
	return structGo
}
func finalizeRequesterPrivate(obj *RequesterPrivate) {
	requesterPrivateStruct.Free(obj.Native())
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
	native unsafe.Pointer
}

func ServerClassNewFromNative(native unsafe.Pointer) *ServerClass {
	instance := &ServerClass{native: native}

	return instance
}

/*
CastToServerClass down casts any arbitrary Object to ServerClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a ServerClass.
*/
func CastToServerClass(object *gobject.Object) *ServerClass {
	return ServerClassNewFromNative(object.Native())
}

// Equals compares this ServerClass with another ServerClass, and returns true if they represent the same Object.
func (recv *ServerClass) Equals(other *ServerClass) bool {
	return other.Native() == recv.Native()
}

func (recv *ServerClass) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentClass returns the C field 'parent_class'.
func (recv *ServerClass) FieldParentClass() *gobject.ObjectClass {
	argValue := gi.StructFieldGet(serverClassStruct, recv.Native(), "parent_class")
	value := gobject.ObjectClassNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *ServerClass) SetFieldParentClass(value *gobject.ObjectClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.StructFieldSet(serverClassStruct, recv.Native(), "parent_class", argValue)
}

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

	structGo := ServerClassNewFromNative(serverClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeServerClass)
	return structGo
}
func finalizeServerClass(obj *ServerClass) {
	serverClassStruct.Free(obj.Native())
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
	native unsafe.Pointer
}

func SessionAsyncClassNewFromNative(native unsafe.Pointer) *SessionAsyncClass {
	instance := &SessionAsyncClass{native: native}

	return instance
}

/*
CastToSessionAsyncClass down casts any arbitrary Object to SessionAsyncClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a SessionAsyncClass.
*/
func CastToSessionAsyncClass(object *gobject.Object) *SessionAsyncClass {
	return SessionAsyncClassNewFromNative(object.Native())
}

// Equals compares this SessionAsyncClass with another SessionAsyncClass, and returns true if they represent the same Object.
func (recv *SessionAsyncClass) Equals(other *SessionAsyncClass) bool {
	return other.Native() == recv.Native()
}

func (recv *SessionAsyncClass) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentClass returns the C field 'parent_class'.
func (recv *SessionAsyncClass) FieldParentClass() *SessionClass {
	argValue := gi.StructFieldGet(sessionAsyncClassStruct, recv.Native(), "parent_class")
	value := SessionClassNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *SessionAsyncClass) SetFieldParentClass(value *SessionClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.StructFieldSet(sessionAsyncClassStruct, recv.Native(), "parent_class", argValue)
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

	structGo := SessionAsyncClassNewFromNative(sessionAsyncClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeSessionAsyncClass)
	return structGo
}
func finalizeSessionAsyncClass(obj *SessionAsyncClass) {
	sessionAsyncClassStruct.Free(obj.Native())
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
	native unsafe.Pointer
}

func SessionClassNewFromNative(native unsafe.Pointer) *SessionClass {
	instance := &SessionClass{native: native}

	return instance
}

/*
CastToSessionClass down casts any arbitrary Object to SessionClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a SessionClass.
*/
func CastToSessionClass(object *gobject.Object) *SessionClass {
	return SessionClassNewFromNative(object.Native())
}

// Equals compares this SessionClass with another SessionClass, and returns true if they represent the same Object.
func (recv *SessionClass) Equals(other *SessionClass) bool {
	return other.Native() == recv.Native()
}

func (recv *SessionClass) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentClass returns the C field 'parent_class'.
func (recv *SessionClass) FieldParentClass() *gobject.ObjectClass {
	argValue := gi.StructFieldGet(sessionClassStruct, recv.Native(), "parent_class")
	value := gobject.ObjectClassNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *SessionClass) SetFieldParentClass(value *gobject.ObjectClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.StructFieldSet(sessionClassStruct, recv.Native(), "parent_class", argValue)
}

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

	structGo := SessionClassNewFromNative(sessionClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeSessionClass)
	return structGo
}
func finalizeSessionClass(obj *SessionClass) {
	sessionClassStruct.Free(obj.Native())
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
	native unsafe.Pointer
}

func SessionFeatureInterfaceNewFromNative(native unsafe.Pointer) *SessionFeatureInterface {
	instance := &SessionFeatureInterface{native: native}

	return instance
}

/*
CastToSessionFeatureInterface down casts any arbitrary Object to SessionFeatureInterface.
Exercise care, as this is a potentially dangerous function
if the Object is not a SessionFeatureInterface.
*/
func CastToSessionFeatureInterface(object *gobject.Object) *SessionFeatureInterface {
	return SessionFeatureInterfaceNewFromNative(object.Native())
}

// Equals compares this SessionFeatureInterface with another SessionFeatureInterface, and returns true if they represent the same Object.
func (recv *SessionFeatureInterface) Equals(other *SessionFeatureInterface) bool {
	return other.Native() == recv.Native()
}

func (recv *SessionFeatureInterface) Native() unsafe.Pointer {
	return recv.native
}

// FieldParent returns the C field 'parent'.
func (recv *SessionFeatureInterface) FieldParent() *gobject.TypeInterface {
	argValue := gi.StructFieldGet(sessionFeatureInterfaceStruct, recv.Native(), "parent")
	value := gobject.TypeInterfaceNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParent sets the value of the C field 'parent'.
func (recv *SessionFeatureInterface) SetFieldParent(value *gobject.TypeInterface) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.StructFieldSet(sessionFeatureInterfaceStruct, recv.Native(), "parent", argValue)
}

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

	structGo := SessionFeatureInterfaceNewFromNative(sessionFeatureInterfaceStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeSessionFeatureInterface)
	return structGo
}
func finalizeSessionFeatureInterface(obj *SessionFeatureInterface) {
	sessionFeatureInterfaceStruct.Free(obj.Native())
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
	native unsafe.Pointer
}

func SessionSyncClassNewFromNative(native unsafe.Pointer) *SessionSyncClass {
	instance := &SessionSyncClass{native: native}

	return instance
}

/*
CastToSessionSyncClass down casts any arbitrary Object to SessionSyncClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a SessionSyncClass.
*/
func CastToSessionSyncClass(object *gobject.Object) *SessionSyncClass {
	return SessionSyncClassNewFromNative(object.Native())
}

// Equals compares this SessionSyncClass with another SessionSyncClass, and returns true if they represent the same Object.
func (recv *SessionSyncClass) Equals(other *SessionSyncClass) bool {
	return other.Native() == recv.Native()
}

func (recv *SessionSyncClass) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentClass returns the C field 'parent_class'.
func (recv *SessionSyncClass) FieldParentClass() *SessionClass {
	argValue := gi.StructFieldGet(sessionSyncClassStruct, recv.Native(), "parent_class")
	value := SessionClassNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *SessionSyncClass) SetFieldParentClass(value *SessionClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.StructFieldSet(sessionSyncClassStruct, recv.Native(), "parent_class", argValue)
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

	structGo := SessionSyncClassNewFromNative(sessionSyncClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeSessionSyncClass)
	return structGo
}
func finalizeSessionSyncClass(obj *SessionSyncClass) {
	sessionSyncClassStruct.Free(obj.Native())
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
	native unsafe.Pointer
}

func SocketClassNewFromNative(native unsafe.Pointer) *SocketClass {
	instance := &SocketClass{native: native}

	return instance
}

/*
CastToSocketClass down casts any arbitrary Object to SocketClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a SocketClass.
*/
func CastToSocketClass(object *gobject.Object) *SocketClass {
	return SocketClassNewFromNative(object.Native())
}

// Equals compares this SocketClass with another SocketClass, and returns true if they represent the same Object.
func (recv *SocketClass) Equals(other *SocketClass) bool {
	return other.Native() == recv.Native()
}

func (recv *SocketClass) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentClass returns the C field 'parent_class'.
func (recv *SocketClass) FieldParentClass() *gobject.ObjectClass {
	argValue := gi.StructFieldGet(socketClassStruct, recv.Native(), "parent_class")
	value := gobject.ObjectClassNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *SocketClass) SetFieldParentClass(value *gobject.ObjectClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.StructFieldSet(socketClassStruct, recv.Native(), "parent_class", argValue)
}

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

	structGo := SocketClassNewFromNative(socketClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeSocketClass)
	return structGo
}
func finalizeSocketClass(obj *SocketClass) {
	socketClassStruct.Free(obj.Native())
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
	native unsafe.Pointer
}

func URINewFromNative(native unsafe.Pointer) *URI {
	instance := &URI{native: native}

	return instance
}

/*
CastToURI down casts any arbitrary Object to URI.
Exercise care, as this is a potentially dangerous function
if the Object is not a URI.
*/
func CastToURI(object *gobject.Object) *URI {
	return URINewFromNative(object.Native())
}

// Equals compares this URI with another URI, and returns true if they represent the same Object.
func (recv *URI) Equals(other *URI) bool {
	return other.Native() == recv.Native()
}

func (recv *URI) Native() unsafe.Pointer {
	return recv.native
}

// FieldScheme returns the C field 'scheme'.
func (recv *URI) FieldScheme() string {
	argValue := gi.StructFieldGet(uRIStruct, recv.Native(), "scheme")
	value := argValue.String(false)
	return value
}

// SetFieldScheme sets the value of the C field 'scheme'.
func (recv *URI) SetFieldScheme(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.StructFieldSet(uRIStruct, recv.Native(), "scheme", argValue)
}

// FieldUser returns the C field 'user'.
func (recv *URI) FieldUser() string {
	argValue := gi.StructFieldGet(uRIStruct, recv.Native(), "user")
	value := argValue.String(false)
	return value
}

// SetFieldUser sets the value of the C field 'user'.
func (recv *URI) SetFieldUser(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.StructFieldSet(uRIStruct, recv.Native(), "user", argValue)
}

// FieldPassword returns the C field 'password'.
func (recv *URI) FieldPassword() string {
	argValue := gi.StructFieldGet(uRIStruct, recv.Native(), "password")
	value := argValue.String(false)
	return value
}

// SetFieldPassword sets the value of the C field 'password'.
func (recv *URI) SetFieldPassword(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.StructFieldSet(uRIStruct, recv.Native(), "password", argValue)
}

// FieldHost returns the C field 'host'.
func (recv *URI) FieldHost() string {
	argValue := gi.StructFieldGet(uRIStruct, recv.Native(), "host")
	value := argValue.String(false)
	return value
}

// SetFieldHost sets the value of the C field 'host'.
func (recv *URI) SetFieldHost(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.StructFieldSet(uRIStruct, recv.Native(), "host", argValue)
}

// FieldPort returns the C field 'port'.
func (recv *URI) FieldPort() uint32 {
	argValue := gi.StructFieldGet(uRIStruct, recv.Native(), "port")
	value := argValue.Uint32()
	return value
}

// SetFieldPort sets the value of the C field 'port'.
func (recv *URI) SetFieldPort(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.StructFieldSet(uRIStruct, recv.Native(), "port", argValue)
}

// FieldPath returns the C field 'path'.
func (recv *URI) FieldPath() string {
	argValue := gi.StructFieldGet(uRIStruct, recv.Native(), "path")
	value := argValue.String(false)
	return value
}

// SetFieldPath sets the value of the C field 'path'.
func (recv *URI) SetFieldPath(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.StructFieldSet(uRIStruct, recv.Native(), "path", argValue)
}

// FieldQuery returns the C field 'query'.
func (recv *URI) FieldQuery() string {
	argValue := gi.StructFieldGet(uRIStruct, recv.Native(), "query")
	value := argValue.String(false)
	return value
}

// SetFieldQuery sets the value of the C field 'query'.
func (recv *URI) SetFieldQuery(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.StructFieldSet(uRIStruct, recv.Native(), "query", argValue)
}

// FieldFragment returns the C field 'fragment'.
func (recv *URI) FieldFragment() string {
	argValue := gi.StructFieldGet(uRIStruct, recv.Native(), "fragment")
	value := argValue.String(false)
	return value
}

// SetFieldFragment sets the value of the C field 'fragment'.
func (recv *URI) SetFieldFragment(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.StructFieldSet(uRIStruct, recv.Native(), "fragment", argValue)
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

	retGo := URINewFromNative(ret.Pointer())

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
	inArgs[0].SetPointer(base.Native())
	inArgs[1].SetString(uriString)

	var ret gi.Argument

	err := uRINewWithBaseFunction_Set()
	if err == nil {
		ret = uRINewWithBaseFunction.Invoke(inArgs[:], nil)
	}

	retGo := URINewFromNative(ret.Pointer())

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
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := uRICopyFunction_Set()
	if err == nil {
		ret = uRICopyFunction.Invoke(inArgs[:], nil)
	}

	retGo := URINewFromNative(ret.Pointer())

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
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := uRICopyHostFunction_Set()
	if err == nil {
		ret = uRICopyHostFunction.Invoke(inArgs[:], nil)
	}

	retGo := URINewFromNative(ret.Pointer())

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
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(uri2.Native())

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
	inArgs[0].SetPointer(recv.Native())

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
	inArgs[0].SetPointer(recv.Native())

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
	inArgs[0].SetPointer(recv.Native())

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
	inArgs[0].SetPointer(recv.Native())

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
	inArgs[0].SetPointer(recv.Native())

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
	inArgs[0].SetPointer(recv.Native())

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
	inArgs[0].SetPointer(recv.Native())

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
	inArgs[0].SetPointer(recv.Native())

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
	inArgs[0].SetPointer(recv.Native())

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
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(v2.Native())

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
	inArgs[0].SetPointer(recv.Native())

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
	inArgs[0].SetPointer(recv.Native())
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
	inArgs[0].SetPointer(recv.Native())
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
	inArgs[0].SetPointer(recv.Native())
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
	inArgs[0].SetPointer(recv.Native())
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
	inArgs[0].SetPointer(recv.Native())
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
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(query)

	err := uRISetQueryFunction_Set()
	if err == nil {
		uRISetQueryFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'soup_uri_set_query_from_fields' : parameter '...' of type 'nil' not supported

var uRISetQueryFromFormFunction *gi.Function
var uRISetQueryFromFormFunction_Once sync.Once

func uRISetQueryFromFormFunction_Set() error {
	var err error
	uRISetQueryFromFormFunction_Once.Do(func() {
		err = uRIStruct_Set()
		if err != nil {
			return
		}
		uRISetQueryFromFormFunction, err = uRIStruct.InvokerNew("set_query_from_form")
	})
	return err
}

// SetQueryFromForm is a representation of the C type soup_uri_set_query_from_form.
func (recv *URI) SetQueryFromForm(form *glib.HashTable) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(form.Native())

	err := uRISetQueryFromFormFunction_Set()
	if err == nil {
		uRISetQueryFromFormFunction.Invoke(inArgs[:], nil)
	}

	return
}

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
	inArgs[0].SetPointer(recv.Native())
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
	inArgs[0].SetPointer(recv.Native())
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
	inArgs[0].SetPointer(recv.Native())
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
	inArgs[0].SetPointer(recv.Native())

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
	native unsafe.Pointer
}

func WebsocketConnectionClassNewFromNative(native unsafe.Pointer) *WebsocketConnectionClass {
	instance := &WebsocketConnectionClass{native: native}

	return instance
}

/*
CastToWebsocketConnectionClass down casts any arbitrary Object to WebsocketConnectionClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a WebsocketConnectionClass.
*/
func CastToWebsocketConnectionClass(object *gobject.Object) *WebsocketConnectionClass {
	return WebsocketConnectionClassNewFromNative(object.Native())
}

// Equals compares this WebsocketConnectionClass with another WebsocketConnectionClass, and returns true if they represent the same Object.
func (recv *WebsocketConnectionClass) Equals(other *WebsocketConnectionClass) bool {
	return other.Native() == recv.Native()
}

func (recv *WebsocketConnectionClass) Native() unsafe.Pointer {
	return recv.native
}

// FieldParent returns the C field 'parent'.
func (recv *WebsocketConnectionClass) FieldParent() *gobject.ObjectClass {
	argValue := gi.StructFieldGet(websocketConnectionClassStruct, recv.Native(), "parent")
	value := gobject.ObjectClassNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParent sets the value of the C field 'parent'.
func (recv *WebsocketConnectionClass) SetFieldParent(value *gobject.ObjectClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.StructFieldSet(websocketConnectionClassStruct, recv.Native(), "parent", argValue)
}

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

	structGo := WebsocketConnectionClassNewFromNative(websocketConnectionClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeWebsocketConnectionClass)
	return structGo
}
func finalizeWebsocketConnectionClass(obj *WebsocketConnectionClass) {
	websocketConnectionClassStruct.Free(obj.Native())
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
	native unsafe.Pointer
}

func WebsocketConnectionPrivateNewFromNative(native unsafe.Pointer) *WebsocketConnectionPrivate {
	instance := &WebsocketConnectionPrivate{native: native}

	return instance
}

/*
CastToWebsocketConnectionPrivate down casts any arbitrary Object to WebsocketConnectionPrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a WebsocketConnectionPrivate.
*/
func CastToWebsocketConnectionPrivate(object *gobject.Object) *WebsocketConnectionPrivate {
	return WebsocketConnectionPrivateNewFromNative(object.Native())
}

// Equals compares this WebsocketConnectionPrivate with another WebsocketConnectionPrivate, and returns true if they represent the same Object.
func (recv *WebsocketConnectionPrivate) Equals(other *WebsocketConnectionPrivate) bool {
	return other.Native() == recv.Native()
}

func (recv *WebsocketConnectionPrivate) Native() unsafe.Pointer {
	return recv.native
}

// WebsocketConnectionPrivateStruct creates an uninitialised WebsocketConnectionPrivate.
func WebsocketConnectionPrivateStruct() *WebsocketConnectionPrivate {
	err := websocketConnectionPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := WebsocketConnectionPrivateNewFromNative(websocketConnectionPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeWebsocketConnectionPrivate)
	return structGo
}
func finalizeWebsocketConnectionPrivate(obj *WebsocketConnectionPrivate) {
	websocketConnectionPrivateStruct.Free(obj.Native())
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
	native unsafe.Pointer
}

func WebsocketExtensionClassNewFromNative(native unsafe.Pointer) *WebsocketExtensionClass {
	instance := &WebsocketExtensionClass{native: native}

	return instance
}

/*
CastToWebsocketExtensionClass down casts any arbitrary Object to WebsocketExtensionClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a WebsocketExtensionClass.
*/
func CastToWebsocketExtensionClass(object *gobject.Object) *WebsocketExtensionClass {
	return WebsocketExtensionClassNewFromNative(object.Native())
}

// Equals compares this WebsocketExtensionClass with another WebsocketExtensionClass, and returns true if they represent the same Object.
func (recv *WebsocketExtensionClass) Equals(other *WebsocketExtensionClass) bool {
	return other.Native() == recv.Native()
}

func (recv *WebsocketExtensionClass) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentClass returns the C field 'parent_class'.
func (recv *WebsocketExtensionClass) FieldParentClass() *gobject.ObjectClass {
	argValue := gi.StructFieldGet(websocketExtensionClassStruct, recv.Native(), "parent_class")
	value := gobject.ObjectClassNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *WebsocketExtensionClass) SetFieldParentClass(value *gobject.ObjectClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.StructFieldSet(websocketExtensionClassStruct, recv.Native(), "parent_class", argValue)
}

// FieldName returns the C field 'name'.
func (recv *WebsocketExtensionClass) FieldName() string {
	argValue := gi.StructFieldGet(websocketExtensionClassStruct, recv.Native(), "name")
	value := argValue.String(false)
	return value
}

// SetFieldName sets the value of the C field 'name'.
func (recv *WebsocketExtensionClass) SetFieldName(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.StructFieldSet(websocketExtensionClassStruct, recv.Native(), "name", argValue)
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

	structGo := WebsocketExtensionClassNewFromNative(websocketExtensionClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeWebsocketExtensionClass)
	return structGo
}
func finalizeWebsocketExtensionClass(obj *WebsocketExtensionClass) {
	websocketExtensionClassStruct.Free(obj.Native())
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
	native unsafe.Pointer
}

func WebsocketExtensionDeflateClassNewFromNative(native unsafe.Pointer) *WebsocketExtensionDeflateClass {
	instance := &WebsocketExtensionDeflateClass{native: native}

	return instance
}

/*
CastToWebsocketExtensionDeflateClass down casts any arbitrary Object to WebsocketExtensionDeflateClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a WebsocketExtensionDeflateClass.
*/
func CastToWebsocketExtensionDeflateClass(object *gobject.Object) *WebsocketExtensionDeflateClass {
	return WebsocketExtensionDeflateClassNewFromNative(object.Native())
}

// Equals compares this WebsocketExtensionDeflateClass with another WebsocketExtensionDeflateClass, and returns true if they represent the same Object.
func (recv *WebsocketExtensionDeflateClass) Equals(other *WebsocketExtensionDeflateClass) bool {
	return other.Native() == recv.Native()
}

func (recv *WebsocketExtensionDeflateClass) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentClass returns the C field 'parent_class'.
func (recv *WebsocketExtensionDeflateClass) FieldParentClass() *WebsocketExtensionClass {
	argValue := gi.StructFieldGet(websocketExtensionDeflateClassStruct, recv.Native(), "parent_class")
	value := WebsocketExtensionClassNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *WebsocketExtensionDeflateClass) SetFieldParentClass(value *WebsocketExtensionClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.StructFieldSet(websocketExtensionDeflateClassStruct, recv.Native(), "parent_class", argValue)
}

// WebsocketExtensionDeflateClassStruct creates an uninitialised WebsocketExtensionDeflateClass.
func WebsocketExtensionDeflateClassStruct() *WebsocketExtensionDeflateClass {
	err := websocketExtensionDeflateClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := WebsocketExtensionDeflateClassNewFromNative(websocketExtensionDeflateClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeWebsocketExtensionDeflateClass)
	return structGo
}
func finalizeWebsocketExtensionDeflateClass(obj *WebsocketExtensionDeflateClass) {
	websocketExtensionDeflateClassStruct.Free(obj.Native())
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
	native unsafe.Pointer
}

func WebsocketExtensionManagerClassNewFromNative(native unsafe.Pointer) *WebsocketExtensionManagerClass {
	instance := &WebsocketExtensionManagerClass{native: native}

	return instance
}

/*
CastToWebsocketExtensionManagerClass down casts any arbitrary Object to WebsocketExtensionManagerClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a WebsocketExtensionManagerClass.
*/
func CastToWebsocketExtensionManagerClass(object *gobject.Object) *WebsocketExtensionManagerClass {
	return WebsocketExtensionManagerClassNewFromNative(object.Native())
}

// Equals compares this WebsocketExtensionManagerClass with another WebsocketExtensionManagerClass, and returns true if they represent the same Object.
func (recv *WebsocketExtensionManagerClass) Equals(other *WebsocketExtensionManagerClass) bool {
	return other.Native() == recv.Native()
}

func (recv *WebsocketExtensionManagerClass) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentClass returns the C field 'parent_class'.
func (recv *WebsocketExtensionManagerClass) FieldParentClass() *gobject.ObjectClass {
	argValue := gi.StructFieldGet(websocketExtensionManagerClassStruct, recv.Native(), "parent_class")
	value := gobject.ObjectClassNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *WebsocketExtensionManagerClass) SetFieldParentClass(value *gobject.ObjectClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.StructFieldSet(websocketExtensionManagerClassStruct, recv.Native(), "parent_class", argValue)
}

// WebsocketExtensionManagerClassStruct creates an uninitialised WebsocketExtensionManagerClass.
func WebsocketExtensionManagerClassStruct() *WebsocketExtensionManagerClass {
	err := websocketExtensionManagerClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := WebsocketExtensionManagerClassNewFromNative(websocketExtensionManagerClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeWebsocketExtensionManagerClass)
	return structGo
}
func finalizeWebsocketExtensionManagerClass(obj *WebsocketExtensionManagerClass) {
	websocketExtensionManagerClassStruct.Free(obj.Native())
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
	native unsafe.Pointer
}

func XMLRPCParamsNewFromNative(native unsafe.Pointer) *XMLRPCParams {
	instance := &XMLRPCParams{native: native}

	return instance
}

/*
CastToXMLRPCParams down casts any arbitrary Object to XMLRPCParams.
Exercise care, as this is a potentially dangerous function
if the Object is not a XMLRPCParams.
*/
func CastToXMLRPCParams(object *gobject.Object) *XMLRPCParams {
	return XMLRPCParamsNewFromNative(object.Native())
}

// Equals compares this XMLRPCParams with another XMLRPCParams, and returns true if they represent the same Object.
func (recv *XMLRPCParams) Equals(other *XMLRPCParams) bool {
	return other.Native() == recv.Native()
}

func (recv *XMLRPCParams) Native() unsafe.Pointer {
	return recv.native
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
	inArgs[0].SetPointer(recv.Native())

	err := xMLRPCParamsFreeFunction_Set()
	if err == nil {
		xMLRPCParamsFreeFunction.Invoke(inArgs[:], nil)
	}

	return
}

var xMLRPCParamsParseFunction *gi.Function
var xMLRPCParamsParseFunction_Once sync.Once

func xMLRPCParamsParseFunction_Set() error {
	var err error
	xMLRPCParamsParseFunction_Once.Do(func() {
		err = xMLRPCParamsStruct_Set()
		if err != nil {
			return
		}
		xMLRPCParamsParseFunction, err = xMLRPCParamsStruct.InvokerNew("parse")
	})
	return err
}

// Parse is a representation of the C type soup_xmlrpc_params_parse.
func (recv *XMLRPCParams) Parse(signature string) *glib.Variant {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(signature)

	var ret gi.Argument

	err := xMLRPCParamsParseFunction_Set()
	if err == nil {
		ret = xMLRPCParamsParseFunction.Invoke(inArgs[:], nil)
	}

	retGo := glib.VariantNewFromNative(ret.Pointer())

	return retGo
}

// XMLRPCParamsStruct creates an uninitialised XMLRPCParams.
func XMLRPCParamsStruct() *XMLRPCParams {
	err := xMLRPCParamsStruct_Set()
	if err != nil {
		return nil
	}

	structGo := XMLRPCParamsNewFromNative(xMLRPCParamsStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeXMLRPCParams)
	return structGo
}
func finalizeXMLRPCParams(obj *XMLRPCParams) {
	xMLRPCParamsStruct.Free(obj.Native())
}
