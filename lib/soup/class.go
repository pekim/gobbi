// Code generated - DO NOT EDIT.

package soup

import (
	callback "github.com/pekim/gobbi/internal/cgo/callback"
	gi "github.com/pekim/gobbi/internal/cgo/gi"
	gio "github.com/pekim/gobbi/lib/gio"
	glib "github.com/pekim/gobbi/lib/glib"
	gobject "github.com/pekim/gobbi/lib/gobject"
	"runtime"
	"sync"
	"unsafe"
)

var addressObject *gi.Object
var addressObject_Once sync.Once

func addressObject_Set() error {
	var err error
	addressObject_Once.Do(func() {
		addressObject, err = gi.ObjectNew("Soup", "Address")
	})
	return err
}

type Address struct {
	native unsafe.Pointer
}

func AddressNewFromNative(native unsafe.Pointer) *Address {
	err := addressObject_Set()
	if err != nil {
		return nil
	}

	instance := &Address{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *Address) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToAddress down casts any arbitrary Object to Address.
Exercise care, as this is a potentially dangerous function
if the Object is not a Address.
*/
func CastToAddress(object *gobject.Object) *Address {
	return AddressNewFromNative(object.Native())
}

// Equals compares this Address with another Address, and returns true if they represent the same Object.
func (recv *Address) Equals(other *Address) bool {
	return other.Native() == recv.Native()
}

func (recv *Address) Native() unsafe.Pointer {
	return recv.native
}

// FieldParent returns the C field 'parent'.
func (recv *Address) FieldParent() *gobject.Object {
	argValue := gi.ObjectFieldGet(addressObject, recv.Native(), "parent")
	value := gobject.ObjectNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParent sets the value of the C field 'parent'.
func (recv *Address) SetFieldParent(value *gobject.Object) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(addressObject, recv.Native(), "parent", argValue)
}

var addressNewFunction *gi.Function
var addressNewFunction_Once sync.Once

func addressNewFunction_Set() error {
	var err error
	addressNewFunction_Once.Do(func() {
		err = addressObject_Set()
		if err != nil {
			return
		}
		addressNewFunction, err = addressObject.InvokerNew("new")
	})
	return err
}

// AddressNew is a representation of the C type soup_address_new.
func AddressNew(name string, port uint32) *Address {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(name)
	inArgs[1].SetUint32(port)

	var ret gi.Argument

	err := addressNewFunction_Set()
	if err == nil {
		ret = addressNewFunction.Invoke(inArgs[:], nil)
	}

	retGo := AddressNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

var addressNewAnyFunction *gi.Function
var addressNewAnyFunction_Once sync.Once

func addressNewAnyFunction_Set() error {
	var err error
	addressNewAnyFunction_Once.Do(func() {
		err = addressObject_Set()
		if err != nil {
			return
		}
		addressNewAnyFunction, err = addressObject.InvokerNew("new_any")
	})
	return err
}

// AddressNewAny is a representation of the C type soup_address_new_any.
func AddressNewAny(family AddressFamily, port uint32) *Address {
	var inArgs [2]gi.Argument
	inArgs[0].SetInt32(int32(family))
	inArgs[1].SetUint32(port)

	var ret gi.Argument

	err := addressNewAnyFunction_Set()
	if err == nil {
		ret = addressNewAnyFunction.Invoke(inArgs[:], nil)
	}

	retGo := AddressNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

var addressNewFromSockaddrFunction *gi.Function
var addressNewFromSockaddrFunction_Once sync.Once

func addressNewFromSockaddrFunction_Set() error {
	var err error
	addressNewFromSockaddrFunction_Once.Do(func() {
		err = addressObject_Set()
		if err != nil {
			return
		}
		addressNewFromSockaddrFunction, err = addressObject.InvokerNew("new_from_sockaddr")
	})
	return err
}

// AddressNewFromSockaddr is a representation of the C type soup_address_new_from_sockaddr.
func AddressNewFromSockaddr(sa unsafe.Pointer, len int32) *Address {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(sa)
	inArgs[1].SetInt32(len)

	var ret gi.Argument

	err := addressNewFromSockaddrFunction_Set()
	if err == nil {
		ret = addressNewFromSockaddrFunction.Invoke(inArgs[:], nil)
	}

	retGo := AddressNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

var addressEqualByIpFunction *gi.Function
var addressEqualByIpFunction_Once sync.Once

func addressEqualByIpFunction_Set() error {
	var err error
	addressEqualByIpFunction_Once.Do(func() {
		err = addressObject_Set()
		if err != nil {
			return
		}
		addressEqualByIpFunction, err = addressObject.InvokerNew("equal_by_ip")
	})
	return err
}

// EqualByIp is a representation of the C type soup_address_equal_by_ip.
func (recv *Address) EqualByIp(addr2 *Address) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(addr2.Native())

	var ret gi.Argument

	err := addressEqualByIpFunction_Set()
	if err == nil {
		ret = addressEqualByIpFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var addressEqualByNameFunction *gi.Function
var addressEqualByNameFunction_Once sync.Once

func addressEqualByNameFunction_Set() error {
	var err error
	addressEqualByNameFunction_Once.Do(func() {
		err = addressObject_Set()
		if err != nil {
			return
		}
		addressEqualByNameFunction, err = addressObject.InvokerNew("equal_by_name")
	})
	return err
}

// EqualByName is a representation of the C type soup_address_equal_by_name.
func (recv *Address) EqualByName(addr2 *Address) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(addr2.Native())

	var ret gi.Argument

	err := addressEqualByNameFunction_Set()
	if err == nil {
		ret = addressEqualByNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var addressGetGsockaddrFunction *gi.Function
var addressGetGsockaddrFunction_Once sync.Once

func addressGetGsockaddrFunction_Set() error {
	var err error
	addressGetGsockaddrFunction_Once.Do(func() {
		err = addressObject_Set()
		if err != nil {
			return
		}
		addressGetGsockaddrFunction, err = addressObject.InvokerNew("get_gsockaddr")
	})
	return err
}

// GetGsockaddr is a representation of the C type soup_address_get_gsockaddr.
func (recv *Address) GetGsockaddr() *gio.SocketAddress {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := addressGetGsockaddrFunction_Set()
	if err == nil {
		ret = addressGetGsockaddrFunction.Invoke(inArgs[:], nil)
	}

	retGo := gio.SocketAddressNewFromNative(ret.Pointer())

	return retGo
}

var addressGetNameFunction *gi.Function
var addressGetNameFunction_Once sync.Once

func addressGetNameFunction_Set() error {
	var err error
	addressGetNameFunction_Once.Do(func() {
		err = addressObject_Set()
		if err != nil {
			return
		}
		addressGetNameFunction, err = addressObject.InvokerNew("get_name")
	})
	return err
}

// GetName is a representation of the C type soup_address_get_name.
func (recv *Address) GetName() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := addressGetNameFunction_Set()
	if err == nil {
		ret = addressGetNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var addressGetPhysicalFunction *gi.Function
var addressGetPhysicalFunction_Once sync.Once

func addressGetPhysicalFunction_Set() error {
	var err error
	addressGetPhysicalFunction_Once.Do(func() {
		err = addressObject_Set()
		if err != nil {
			return
		}
		addressGetPhysicalFunction, err = addressObject.InvokerNew("get_physical")
	})
	return err
}

// GetPhysical is a representation of the C type soup_address_get_physical.
func (recv *Address) GetPhysical() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := addressGetPhysicalFunction_Set()
	if err == nil {
		ret = addressGetPhysicalFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var addressGetPortFunction *gi.Function
var addressGetPortFunction_Once sync.Once

func addressGetPortFunction_Set() error {
	var err error
	addressGetPortFunction_Once.Do(func() {
		err = addressObject_Set()
		if err != nil {
			return
		}
		addressGetPortFunction, err = addressObject.InvokerNew("get_port")
	})
	return err
}

// GetPort is a representation of the C type soup_address_get_port.
func (recv *Address) GetPort() uint32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := addressGetPortFunction_Set()
	if err == nil {
		ret = addressGetPortFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo
}

var addressGetSockaddrFunction *gi.Function
var addressGetSockaddrFunction_Once sync.Once

func addressGetSockaddrFunction_Set() error {
	var err error
	addressGetSockaddrFunction_Once.Do(func() {
		err = addressObject_Set()
		if err != nil {
			return
		}
		addressGetSockaddrFunction, err = addressObject.InvokerNew("get_sockaddr")
	})
	return err
}

// GetSockaddr is a representation of the C type soup_address_get_sockaddr.
func (recv *Address) GetSockaddr(len int32) unsafe.Pointer {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(len)

	var ret gi.Argument

	err := addressGetSockaddrFunction_Set()
	if err == nil {
		ret = addressGetSockaddrFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Pointer()

	return retGo
}

var addressHashByIpFunction *gi.Function
var addressHashByIpFunction_Once sync.Once

func addressHashByIpFunction_Set() error {
	var err error
	addressHashByIpFunction_Once.Do(func() {
		err = addressObject_Set()
		if err != nil {
			return
		}
		addressHashByIpFunction, err = addressObject.InvokerNew("hash_by_ip")
	})
	return err
}

// HashByIp is a representation of the C type soup_address_hash_by_ip.
func (recv *Address) HashByIp() uint32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := addressHashByIpFunction_Set()
	if err == nil {
		ret = addressHashByIpFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo
}

var addressHashByNameFunction *gi.Function
var addressHashByNameFunction_Once sync.Once

func addressHashByNameFunction_Set() error {
	var err error
	addressHashByNameFunction_Once.Do(func() {
		err = addressObject_Set()
		if err != nil {
			return
		}
		addressHashByNameFunction, err = addressObject.InvokerNew("hash_by_name")
	})
	return err
}

// HashByName is a representation of the C type soup_address_hash_by_name.
func (recv *Address) HashByName() uint32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := addressHashByNameFunction_Set()
	if err == nil {
		ret = addressHashByNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo
}

var addressIsResolvedFunction *gi.Function
var addressIsResolvedFunction_Once sync.Once

func addressIsResolvedFunction_Set() error {
	var err error
	addressIsResolvedFunction_Once.Do(func() {
		err = addressObject_Set()
		if err != nil {
			return
		}
		addressIsResolvedFunction, err = addressObject.InvokerNew("is_resolved")
	})
	return err
}

// IsResolved is a representation of the C type soup_address_is_resolved.
func (recv *Address) IsResolved() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := addressIsResolvedFunction_Set()
	if err == nil {
		ret = addressIsResolvedFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'soup_address_resolve_async' : parameter 'callback' of type 'AddressCallback' not supported

var addressResolveSyncFunction *gi.Function
var addressResolveSyncFunction_Once sync.Once

func addressResolveSyncFunction_Set() error {
	var err error
	addressResolveSyncFunction_Once.Do(func() {
		err = addressObject_Set()
		if err != nil {
			return
		}
		addressResolveSyncFunction, err = addressObject.InvokerNew("resolve_sync")
	})
	return err
}

// ResolveSync is a representation of the C type soup_address_resolve_sync.
func (recv *Address) ResolveSync(cancellable *gio.Cancellable) uint32 {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(cancellable.Native())

	var ret gi.Argument

	err := addressResolveSyncFunction_Set()
	if err == nil {
		ret = addressResolveSyncFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo
}

// GioSocketConnectable returns the Gio.SocketConnectable interface implemented by Address
func (recv *Address) GioSocketConnectable() *gio.SocketConnectable {
	return gio.SocketConnectableNewFromNative(recv.Native())
}

var authObject *gi.Object
var authObject_Once sync.Once

func authObject_Set() error {
	var err error
	authObject_Once.Do(func() {
		authObject, err = gi.ObjectNew("Soup", "Auth")
	})
	return err
}

type Auth struct {
	native unsafe.Pointer
}

func AuthNewFromNative(native unsafe.Pointer) *Auth {
	err := authObject_Set()
	if err != nil {
		return nil
	}

	instance := &Auth{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *Auth) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToAuth down casts any arbitrary Object to Auth.
Exercise care, as this is a potentially dangerous function
if the Object is not a Auth.
*/
func CastToAuth(object *gobject.Object) *Auth {
	return AuthNewFromNative(object.Native())
}

// Equals compares this Auth with another Auth, and returns true if they represent the same Object.
func (recv *Auth) Equals(other *Auth) bool {
	return other.Native() == recv.Native()
}

func (recv *Auth) Native() unsafe.Pointer {
	return recv.native
}

// FieldParent returns the C field 'parent'.
func (recv *Auth) FieldParent() *gobject.Object {
	argValue := gi.ObjectFieldGet(authObject, recv.Native(), "parent")
	value := gobject.ObjectNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParent sets the value of the C field 'parent'.
func (recv *Auth) SetFieldParent(value *gobject.Object) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(authObject, recv.Native(), "parent", argValue)
}

// FieldRealm returns the C field 'realm'.
func (recv *Auth) FieldRealm() string {
	argValue := gi.ObjectFieldGet(authObject, recv.Native(), "realm")
	value := argValue.String(false)
	return value
}

// SetFieldRealm sets the value of the C field 'realm'.
func (recv *Auth) SetFieldRealm(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.ObjectFieldSet(authObject, recv.Native(), "realm", argValue)
}

var authNewFunction *gi.Function
var authNewFunction_Once sync.Once

func authNewFunction_Set() error {
	var err error
	authNewFunction_Once.Do(func() {
		err = authObject_Set()
		if err != nil {
			return
		}
		authNewFunction, err = authObject.InvokerNew("new")
	})
	return err
}

// AuthNew is a representation of the C type soup_auth_new.
func AuthNew(type_ int64, msg *Message, authHeader string) *Auth {
	var inArgs [3]gi.Argument
	inArgs[0].SetInt64(type_)
	inArgs[1].SetPointer(msg.Native())
	inArgs[2].SetString(authHeader)

	var ret gi.Argument

	err := authNewFunction_Set()
	if err == nil {
		ret = authNewFunction.Invoke(inArgs[:], nil)
	}

	retGo := AuthNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

var authAuthenticateFunction *gi.Function
var authAuthenticateFunction_Once sync.Once

func authAuthenticateFunction_Set() error {
	var err error
	authAuthenticateFunction_Once.Do(func() {
		err = authObject_Set()
		if err != nil {
			return
		}
		authAuthenticateFunction, err = authObject.InvokerNew("authenticate")
	})
	return err
}

// Authenticate is a representation of the C type soup_auth_authenticate.
func (recv *Auth) Authenticate(username string, password string) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(username)
	inArgs[2].SetString(password)

	err := authAuthenticateFunction_Set()
	if err == nil {
		authAuthenticateFunction.Invoke(inArgs[:], nil)
	}

	return
}

var authCanAuthenticateFunction *gi.Function
var authCanAuthenticateFunction_Once sync.Once

func authCanAuthenticateFunction_Set() error {
	var err error
	authCanAuthenticateFunction_Once.Do(func() {
		err = authObject_Set()
		if err != nil {
			return
		}
		authCanAuthenticateFunction, err = authObject.InvokerNew("can_authenticate")
	})
	return err
}

// CanAuthenticate is a representation of the C type soup_auth_can_authenticate.
func (recv *Auth) CanAuthenticate() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := authCanAuthenticateFunction_Set()
	if err == nil {
		ret = authCanAuthenticateFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var authFreeProtectionSpaceFunction *gi.Function
var authFreeProtectionSpaceFunction_Once sync.Once

func authFreeProtectionSpaceFunction_Set() error {
	var err error
	authFreeProtectionSpaceFunction_Once.Do(func() {
		err = authObject_Set()
		if err != nil {
			return
		}
		authFreeProtectionSpaceFunction, err = authObject.InvokerNew("free_protection_space")
	})
	return err
}

// FreeProtectionSpace is a representation of the C type soup_auth_free_protection_space.
func (recv *Auth) FreeProtectionSpace(space *glib.SList) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(space.Native())

	err := authFreeProtectionSpaceFunction_Set()
	if err == nil {
		authFreeProtectionSpaceFunction.Invoke(inArgs[:], nil)
	}

	return
}

var authGetAuthorizationFunction *gi.Function
var authGetAuthorizationFunction_Once sync.Once

func authGetAuthorizationFunction_Set() error {
	var err error
	authGetAuthorizationFunction_Once.Do(func() {
		err = authObject_Set()
		if err != nil {
			return
		}
		authGetAuthorizationFunction, err = authObject.InvokerNew("get_authorization")
	})
	return err
}

// GetAuthorization is a representation of the C type soup_auth_get_authorization.
func (recv *Auth) GetAuthorization(msg *Message) string {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(msg.Native())

	var ret gi.Argument

	err := authGetAuthorizationFunction_Set()
	if err == nil {
		ret = authGetAuthorizationFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

var authGetHostFunction *gi.Function
var authGetHostFunction_Once sync.Once

func authGetHostFunction_Set() error {
	var err error
	authGetHostFunction_Once.Do(func() {
		err = authObject_Set()
		if err != nil {
			return
		}
		authGetHostFunction, err = authObject.InvokerNew("get_host")
	})
	return err
}

// GetHost is a representation of the C type soup_auth_get_host.
func (recv *Auth) GetHost() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := authGetHostFunction_Set()
	if err == nil {
		ret = authGetHostFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var authGetInfoFunction *gi.Function
var authGetInfoFunction_Once sync.Once

func authGetInfoFunction_Set() error {
	var err error
	authGetInfoFunction_Once.Do(func() {
		err = authObject_Set()
		if err != nil {
			return
		}
		authGetInfoFunction, err = authObject.InvokerNew("get_info")
	})
	return err
}

// GetInfo is a representation of the C type soup_auth_get_info.
func (recv *Auth) GetInfo() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := authGetInfoFunction_Set()
	if err == nil {
		ret = authGetInfoFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

var authGetProtectionSpaceFunction *gi.Function
var authGetProtectionSpaceFunction_Once sync.Once

func authGetProtectionSpaceFunction_Set() error {
	var err error
	authGetProtectionSpaceFunction_Once.Do(func() {
		err = authObject_Set()
		if err != nil {
			return
		}
		authGetProtectionSpaceFunction, err = authObject.InvokerNew("get_protection_space")
	})
	return err
}

// GetProtectionSpace is a representation of the C type soup_auth_get_protection_space.
func (recv *Auth) GetProtectionSpace(sourceUri *URI) *glib.SList {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(sourceUri.Native())

	var ret gi.Argument

	err := authGetProtectionSpaceFunction_Set()
	if err == nil {
		ret = authGetProtectionSpaceFunction.Invoke(inArgs[:], nil)
	}

	retGo := glib.SListNewFromNative(ret.Pointer())

	return retGo
}

var authGetRealmFunction *gi.Function
var authGetRealmFunction_Once sync.Once

func authGetRealmFunction_Set() error {
	var err error
	authGetRealmFunction_Once.Do(func() {
		err = authObject_Set()
		if err != nil {
			return
		}
		authGetRealmFunction, err = authObject.InvokerNew("get_realm")
	})
	return err
}

// GetRealm is a representation of the C type soup_auth_get_realm.
func (recv *Auth) GetRealm() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := authGetRealmFunction_Set()
	if err == nil {
		ret = authGetRealmFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var authGetSavedPasswordFunction *gi.Function
var authGetSavedPasswordFunction_Once sync.Once

func authGetSavedPasswordFunction_Set() error {
	var err error
	authGetSavedPasswordFunction_Once.Do(func() {
		err = authObject_Set()
		if err != nil {
			return
		}
		authGetSavedPasswordFunction, err = authObject.InvokerNew("get_saved_password")
	})
	return err
}

// GetSavedPassword is a representation of the C type soup_auth_get_saved_password.
func (recv *Auth) GetSavedPassword(user string) string {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(user)

	var ret gi.Argument

	err := authGetSavedPasswordFunction_Set()
	if err == nil {
		ret = authGetSavedPasswordFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var authGetSavedUsersFunction *gi.Function
var authGetSavedUsersFunction_Once sync.Once

func authGetSavedUsersFunction_Set() error {
	var err error
	authGetSavedUsersFunction_Once.Do(func() {
		err = authObject_Set()
		if err != nil {
			return
		}
		authGetSavedUsersFunction, err = authObject.InvokerNew("get_saved_users")
	})
	return err
}

// GetSavedUsers is a representation of the C type soup_auth_get_saved_users.
func (recv *Auth) GetSavedUsers() *glib.SList {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := authGetSavedUsersFunction_Set()
	if err == nil {
		ret = authGetSavedUsersFunction.Invoke(inArgs[:], nil)
	}

	retGo := glib.SListNewFromNative(ret.Pointer())

	return retGo
}

var authGetSchemeNameFunction *gi.Function
var authGetSchemeNameFunction_Once sync.Once

func authGetSchemeNameFunction_Set() error {
	var err error
	authGetSchemeNameFunction_Once.Do(func() {
		err = authObject_Set()
		if err != nil {
			return
		}
		authGetSchemeNameFunction, err = authObject.InvokerNew("get_scheme_name")
	})
	return err
}

// GetSchemeName is a representation of the C type soup_auth_get_scheme_name.
func (recv *Auth) GetSchemeName() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := authGetSchemeNameFunction_Set()
	if err == nil {
		ret = authGetSchemeNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var authHasSavedPasswordFunction *gi.Function
var authHasSavedPasswordFunction_Once sync.Once

func authHasSavedPasswordFunction_Set() error {
	var err error
	authHasSavedPasswordFunction_Once.Do(func() {
		err = authObject_Set()
		if err != nil {
			return
		}
		authHasSavedPasswordFunction, err = authObject.InvokerNew("has_saved_password")
	})
	return err
}

// HasSavedPassword is a representation of the C type soup_auth_has_saved_password.
func (recv *Auth) HasSavedPassword(username string, password string) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(username)
	inArgs[2].SetString(password)

	err := authHasSavedPasswordFunction_Set()
	if err == nil {
		authHasSavedPasswordFunction.Invoke(inArgs[:], nil)
	}

	return
}

var authIsAuthenticatedFunction *gi.Function
var authIsAuthenticatedFunction_Once sync.Once

func authIsAuthenticatedFunction_Set() error {
	var err error
	authIsAuthenticatedFunction_Once.Do(func() {
		err = authObject_Set()
		if err != nil {
			return
		}
		authIsAuthenticatedFunction, err = authObject.InvokerNew("is_authenticated")
	})
	return err
}

// IsAuthenticated is a representation of the C type soup_auth_is_authenticated.
func (recv *Auth) IsAuthenticated() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := authIsAuthenticatedFunction_Set()
	if err == nil {
		ret = authIsAuthenticatedFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var authIsForProxyFunction *gi.Function
var authIsForProxyFunction_Once sync.Once

func authIsForProxyFunction_Set() error {
	var err error
	authIsForProxyFunction_Once.Do(func() {
		err = authObject_Set()
		if err != nil {
			return
		}
		authIsForProxyFunction, err = authObject.InvokerNew("is_for_proxy")
	})
	return err
}

// IsForProxy is a representation of the C type soup_auth_is_for_proxy.
func (recv *Auth) IsForProxy() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := authIsForProxyFunction_Set()
	if err == nil {
		ret = authIsForProxyFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var authIsReadyFunction *gi.Function
var authIsReadyFunction_Once sync.Once

func authIsReadyFunction_Set() error {
	var err error
	authIsReadyFunction_Once.Do(func() {
		err = authObject_Set()
		if err != nil {
			return
		}
		authIsReadyFunction, err = authObject.InvokerNew("is_ready")
	})
	return err
}

// IsReady is a representation of the C type soup_auth_is_ready.
func (recv *Auth) IsReady(msg *Message) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(msg.Native())

	var ret gi.Argument

	err := authIsReadyFunction_Set()
	if err == nil {
		ret = authIsReadyFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var authSavePasswordFunction *gi.Function
var authSavePasswordFunction_Once sync.Once

func authSavePasswordFunction_Set() error {
	var err error
	authSavePasswordFunction_Once.Do(func() {
		err = authObject_Set()
		if err != nil {
			return
		}
		authSavePasswordFunction, err = authObject.InvokerNew("save_password")
	})
	return err
}

// SavePassword is a representation of the C type soup_auth_save_password.
func (recv *Auth) SavePassword(username string, password string) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(username)
	inArgs[2].SetString(password)

	err := authSavePasswordFunction_Set()
	if err == nil {
		authSavePasswordFunction.Invoke(inArgs[:], nil)
	}

	return
}

var authUpdateFunction *gi.Function
var authUpdateFunction_Once sync.Once

func authUpdateFunction_Set() error {
	var err error
	authUpdateFunction_Once.Do(func() {
		err = authObject_Set()
		if err != nil {
			return
		}
		authUpdateFunction, err = authObject.InvokerNew("update")
	})
	return err
}

// Update is a representation of the C type soup_auth_update.
func (recv *Auth) Update(msg *Message, authHeader string) bool {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(msg.Native())
	inArgs[2].SetString(authHeader)

	var ret gi.Argument

	err := authUpdateFunction_Set()
	if err == nil {
		ret = authUpdateFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var authBasicObject *gi.Object
var authBasicObject_Once sync.Once

func authBasicObject_Set() error {
	var err error
	authBasicObject_Once.Do(func() {
		authBasicObject, err = gi.ObjectNew("Soup", "AuthBasic")
	})
	return err
}

type AuthBasic struct {
	native unsafe.Pointer
}

func AuthBasicNewFromNative(native unsafe.Pointer) *AuthBasic {
	err := authBasicObject_Set()
	if err != nil {
		return nil
	}

	instance := &AuthBasic{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Auth upcasts to *Auth
func (recv *AuthBasic) Auth() *Auth {
	return AuthNewFromNative(recv.Native())
}

// Object upcasts to *Object
func (recv *AuthBasic) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToAuthBasic down casts any arbitrary Object to AuthBasic.
Exercise care, as this is a potentially dangerous function
if the Object is not a AuthBasic.
*/
func CastToAuthBasic(object *gobject.Object) *AuthBasic {
	return AuthBasicNewFromNative(object.Native())
}

// Equals compares this AuthBasic with another AuthBasic, and returns true if they represent the same Object.
func (recv *AuthBasic) Equals(other *AuthBasic) bool {
	return other.Native() == recv.Native()
}

func (recv *AuthBasic) Native() unsafe.Pointer {
	return recv.native
}

var authDigestObject *gi.Object
var authDigestObject_Once sync.Once

func authDigestObject_Set() error {
	var err error
	authDigestObject_Once.Do(func() {
		authDigestObject, err = gi.ObjectNew("Soup", "AuthDigest")
	})
	return err
}

type AuthDigest struct {
	native unsafe.Pointer
}

func AuthDigestNewFromNative(native unsafe.Pointer) *AuthDigest {
	err := authDigestObject_Set()
	if err != nil {
		return nil
	}

	instance := &AuthDigest{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Auth upcasts to *Auth
func (recv *AuthDigest) Auth() *Auth {
	return AuthNewFromNative(recv.Native())
}

// Object upcasts to *Object
func (recv *AuthDigest) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToAuthDigest down casts any arbitrary Object to AuthDigest.
Exercise care, as this is a potentially dangerous function
if the Object is not a AuthDigest.
*/
func CastToAuthDigest(object *gobject.Object) *AuthDigest {
	return AuthDigestNewFromNative(object.Native())
}

// Equals compares this AuthDigest with another AuthDigest, and returns true if they represent the same Object.
func (recv *AuthDigest) Equals(other *AuthDigest) bool {
	return other.Native() == recv.Native()
}

func (recv *AuthDigest) Native() unsafe.Pointer {
	return recv.native
}

var authDomainObject *gi.Object
var authDomainObject_Once sync.Once

func authDomainObject_Set() error {
	var err error
	authDomainObject_Once.Do(func() {
		authDomainObject, err = gi.ObjectNew("Soup", "AuthDomain")
	})
	return err
}

type AuthDomain struct {
	native unsafe.Pointer
}

func AuthDomainNewFromNative(native unsafe.Pointer) *AuthDomain {
	err := authDomainObject_Set()
	if err != nil {
		return nil
	}

	instance := &AuthDomain{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *AuthDomain) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToAuthDomain down casts any arbitrary Object to AuthDomain.
Exercise care, as this is a potentially dangerous function
if the Object is not a AuthDomain.
*/
func CastToAuthDomain(object *gobject.Object) *AuthDomain {
	return AuthDomainNewFromNative(object.Native())
}

// Equals compares this AuthDomain with another AuthDomain, and returns true if they represent the same Object.
func (recv *AuthDomain) Equals(other *AuthDomain) bool {
	return other.Native() == recv.Native()
}

func (recv *AuthDomain) Native() unsafe.Pointer {
	return recv.native
}

// FieldParent returns the C field 'parent'.
func (recv *AuthDomain) FieldParent() *gobject.Object {
	argValue := gi.ObjectFieldGet(authDomainObject, recv.Native(), "parent")
	value := gobject.ObjectNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParent sets the value of the C field 'parent'.
func (recv *AuthDomain) SetFieldParent(value *gobject.Object) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(authDomainObject, recv.Native(), "parent", argValue)
}

var authDomainAcceptsFunction *gi.Function
var authDomainAcceptsFunction_Once sync.Once

func authDomainAcceptsFunction_Set() error {
	var err error
	authDomainAcceptsFunction_Once.Do(func() {
		err = authDomainObject_Set()
		if err != nil {
			return
		}
		authDomainAcceptsFunction, err = authDomainObject.InvokerNew("accepts")
	})
	return err
}

// Accepts is a representation of the C type soup_auth_domain_accepts.
func (recv *AuthDomain) Accepts(msg *Message) string {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(msg.Native())

	var ret gi.Argument

	err := authDomainAcceptsFunction_Set()
	if err == nil {
		ret = authDomainAcceptsFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

var authDomainAddPathFunction *gi.Function
var authDomainAddPathFunction_Once sync.Once

func authDomainAddPathFunction_Set() error {
	var err error
	authDomainAddPathFunction_Once.Do(func() {
		err = authDomainObject_Set()
		if err != nil {
			return
		}
		authDomainAddPathFunction, err = authDomainObject.InvokerNew("add_path")
	})
	return err
}

// AddPath is a representation of the C type soup_auth_domain_add_path.
func (recv *AuthDomain) AddPath(path string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(path)

	err := authDomainAddPathFunction_Set()
	if err == nil {
		authDomainAddPathFunction.Invoke(inArgs[:], nil)
	}

	return
}

var authDomainChallengeFunction *gi.Function
var authDomainChallengeFunction_Once sync.Once

func authDomainChallengeFunction_Set() error {
	var err error
	authDomainChallengeFunction_Once.Do(func() {
		err = authDomainObject_Set()
		if err != nil {
			return
		}
		authDomainChallengeFunction, err = authDomainObject.InvokerNew("challenge")
	})
	return err
}

// Challenge is a representation of the C type soup_auth_domain_challenge.
func (recv *AuthDomain) Challenge(msg *Message) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(msg.Native())

	err := authDomainChallengeFunction_Set()
	if err == nil {
		authDomainChallengeFunction.Invoke(inArgs[:], nil)
	}

	return
}

var authDomainCheckPasswordFunction *gi.Function
var authDomainCheckPasswordFunction_Once sync.Once

func authDomainCheckPasswordFunction_Set() error {
	var err error
	authDomainCheckPasswordFunction_Once.Do(func() {
		err = authDomainObject_Set()
		if err != nil {
			return
		}
		authDomainCheckPasswordFunction, err = authDomainObject.InvokerNew("check_password")
	})
	return err
}

// CheckPassword is a representation of the C type soup_auth_domain_check_password.
func (recv *AuthDomain) CheckPassword(msg *Message, username string, password string) bool {
	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(msg.Native())
	inArgs[2].SetString(username)
	inArgs[3].SetString(password)

	var ret gi.Argument

	err := authDomainCheckPasswordFunction_Set()
	if err == nil {
		ret = authDomainCheckPasswordFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var authDomainCoversFunction *gi.Function
var authDomainCoversFunction_Once sync.Once

func authDomainCoversFunction_Set() error {
	var err error
	authDomainCoversFunction_Once.Do(func() {
		err = authDomainObject_Set()
		if err != nil {
			return
		}
		authDomainCoversFunction, err = authDomainObject.InvokerNew("covers")
	})
	return err
}

// Covers is a representation of the C type soup_auth_domain_covers.
func (recv *AuthDomain) Covers(msg *Message) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(msg.Native())

	var ret gi.Argument

	err := authDomainCoversFunction_Set()
	if err == nil {
		ret = authDomainCoversFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var authDomainGetRealmFunction *gi.Function
var authDomainGetRealmFunction_Once sync.Once

func authDomainGetRealmFunction_Set() error {
	var err error
	authDomainGetRealmFunction_Once.Do(func() {
		err = authDomainObject_Set()
		if err != nil {
			return
		}
		authDomainGetRealmFunction, err = authDomainObject.InvokerNew("get_realm")
	})
	return err
}

// GetRealm is a representation of the C type soup_auth_domain_get_realm.
func (recv *AuthDomain) GetRealm() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := authDomainGetRealmFunction_Set()
	if err == nil {
		ret = authDomainGetRealmFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var authDomainRemovePathFunction *gi.Function
var authDomainRemovePathFunction_Once sync.Once

func authDomainRemovePathFunction_Set() error {
	var err error
	authDomainRemovePathFunction_Once.Do(func() {
		err = authDomainObject_Set()
		if err != nil {
			return
		}
		authDomainRemovePathFunction, err = authDomainObject.InvokerNew("remove_path")
	})
	return err
}

// RemovePath is a representation of the C type soup_auth_domain_remove_path.
func (recv *AuthDomain) RemovePath(path string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(path)

	err := authDomainRemovePathFunction_Set()
	if err == nil {
		authDomainRemovePathFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'soup_auth_domain_set_filter' : parameter 'filter' of type 'AuthDomainFilter' not supported

// UNSUPPORTED : C value 'soup_auth_domain_set_generic_auth_callback' : parameter 'auth_callback' of type 'AuthDomainGenericAuthCallback' not supported

var authDomainTryGenericAuthCallbackFunction *gi.Function
var authDomainTryGenericAuthCallbackFunction_Once sync.Once

func authDomainTryGenericAuthCallbackFunction_Set() error {
	var err error
	authDomainTryGenericAuthCallbackFunction_Once.Do(func() {
		err = authDomainObject_Set()
		if err != nil {
			return
		}
		authDomainTryGenericAuthCallbackFunction, err = authDomainObject.InvokerNew("try_generic_auth_callback")
	})
	return err
}

// TryGenericAuthCallback is a representation of the C type soup_auth_domain_try_generic_auth_callback.
func (recv *AuthDomain) TryGenericAuthCallback(msg *Message, username string) bool {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(msg.Native())
	inArgs[2].SetString(username)

	var ret gi.Argument

	err := authDomainTryGenericAuthCallbackFunction_Set()
	if err == nil {
		ret = authDomainTryGenericAuthCallbackFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var authDomainBasicObject *gi.Object
var authDomainBasicObject_Once sync.Once

func authDomainBasicObject_Set() error {
	var err error
	authDomainBasicObject_Once.Do(func() {
		authDomainBasicObject, err = gi.ObjectNew("Soup", "AuthDomainBasic")
	})
	return err
}

type AuthDomainBasic struct {
	native unsafe.Pointer
}

func AuthDomainBasicNewFromNative(native unsafe.Pointer) *AuthDomainBasic {
	err := authDomainBasicObject_Set()
	if err != nil {
		return nil
	}

	instance := &AuthDomainBasic{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// AuthDomain upcasts to *AuthDomain
func (recv *AuthDomainBasic) AuthDomain() *AuthDomain {
	return AuthDomainNewFromNative(recv.Native())
}

// Object upcasts to *Object
func (recv *AuthDomainBasic) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToAuthDomainBasic down casts any arbitrary Object to AuthDomainBasic.
Exercise care, as this is a potentially dangerous function
if the Object is not a AuthDomainBasic.
*/
func CastToAuthDomainBasic(object *gobject.Object) *AuthDomainBasic {
	return AuthDomainBasicNewFromNative(object.Native())
}

// Equals compares this AuthDomainBasic with another AuthDomainBasic, and returns true if they represent the same Object.
func (recv *AuthDomainBasic) Equals(other *AuthDomainBasic) bool {
	return other.Native() == recv.Native()
}

func (recv *AuthDomainBasic) Native() unsafe.Pointer {
	return recv.native
}

// FieldParent returns the C field 'parent'.
func (recv *AuthDomainBasic) FieldParent() *AuthDomain {
	argValue := gi.ObjectFieldGet(authDomainBasicObject, recv.Native(), "parent")
	value := AuthDomainNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParent sets the value of the C field 'parent'.
func (recv *AuthDomainBasic) SetFieldParent(value *AuthDomain) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(authDomainBasicObject, recv.Native(), "parent", argValue)
}

// UNSUPPORTED : C value 'soup_auth_domain_basic_new' : parameter '...' of type 'nil' not supported

// UNSUPPORTED : C value 'soup_auth_domain_basic_set_auth_callback' : parameter 'callback' of type 'AuthDomainBasicAuthCallback' not supported

var authDomainDigestObject *gi.Object
var authDomainDigestObject_Once sync.Once

func authDomainDigestObject_Set() error {
	var err error
	authDomainDigestObject_Once.Do(func() {
		authDomainDigestObject, err = gi.ObjectNew("Soup", "AuthDomainDigest")
	})
	return err
}

type AuthDomainDigest struct {
	native unsafe.Pointer
}

func AuthDomainDigestNewFromNative(native unsafe.Pointer) *AuthDomainDigest {
	err := authDomainDigestObject_Set()
	if err != nil {
		return nil
	}

	instance := &AuthDomainDigest{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// AuthDomain upcasts to *AuthDomain
func (recv *AuthDomainDigest) AuthDomain() *AuthDomain {
	return AuthDomainNewFromNative(recv.Native())
}

// Object upcasts to *Object
func (recv *AuthDomainDigest) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToAuthDomainDigest down casts any arbitrary Object to AuthDomainDigest.
Exercise care, as this is a potentially dangerous function
if the Object is not a AuthDomainDigest.
*/
func CastToAuthDomainDigest(object *gobject.Object) *AuthDomainDigest {
	return AuthDomainDigestNewFromNative(object.Native())
}

// Equals compares this AuthDomainDigest with another AuthDomainDigest, and returns true if they represent the same Object.
func (recv *AuthDomainDigest) Equals(other *AuthDomainDigest) bool {
	return other.Native() == recv.Native()
}

func (recv *AuthDomainDigest) Native() unsafe.Pointer {
	return recv.native
}

// FieldParent returns the C field 'parent'.
func (recv *AuthDomainDigest) FieldParent() *AuthDomain {
	argValue := gi.ObjectFieldGet(authDomainDigestObject, recv.Native(), "parent")
	value := AuthDomainNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParent sets the value of the C field 'parent'.
func (recv *AuthDomainDigest) SetFieldParent(value *AuthDomain) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(authDomainDigestObject, recv.Native(), "parent", argValue)
}

// UNSUPPORTED : C value 'soup_auth_domain_digest_new' : parameter '...' of type 'nil' not supported

// UNSUPPORTED : C value 'soup_auth_domain_digest_set_auth_callback' : parameter 'callback' of type 'AuthDomainDigestAuthCallback' not supported

var authManagerObject *gi.Object
var authManagerObject_Once sync.Once

func authManagerObject_Set() error {
	var err error
	authManagerObject_Once.Do(func() {
		authManagerObject, err = gi.ObjectNew("Soup", "AuthManager")
	})
	return err
}

type AuthManager struct {
	native unsafe.Pointer
}

func AuthManagerNewFromNative(native unsafe.Pointer) *AuthManager {
	err := authManagerObject_Set()
	if err != nil {
		return nil
	}

	instance := &AuthManager{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *AuthManager) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToAuthManager down casts any arbitrary Object to AuthManager.
Exercise care, as this is a potentially dangerous function
if the Object is not a AuthManager.
*/
func CastToAuthManager(object *gobject.Object) *AuthManager {
	return AuthManagerNewFromNative(object.Native())
}

// Equals compares this AuthManager with another AuthManager, and returns true if they represent the same Object.
func (recv *AuthManager) Equals(other *AuthManager) bool {
	return other.Native() == recv.Native()
}

func (recv *AuthManager) Native() unsafe.Pointer {
	return recv.native
}

// FieldParent returns the C field 'parent'.
func (recv *AuthManager) FieldParent() *gobject.Object {
	argValue := gi.ObjectFieldGet(authManagerObject, recv.Native(), "parent")
	value := gobject.ObjectNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParent sets the value of the C field 'parent'.
func (recv *AuthManager) SetFieldParent(value *gobject.Object) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(authManagerObject, recv.Native(), "parent", argValue)
}

// FieldPriv returns the C field 'priv'.
func (recv *AuthManager) FieldPriv() *AuthManagerPrivate {
	argValue := gi.ObjectFieldGet(authManagerObject, recv.Native(), "priv")
	value := AuthManagerPrivateNewFromNative(argValue.Pointer())
	return value
}

// SetFieldPriv sets the value of the C field 'priv'.
func (recv *AuthManager) SetFieldPriv(value *AuthManagerPrivate) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(authManagerObject, recv.Native(), "priv", argValue)
}

var authManagerClearCachedCredentialsFunction *gi.Function
var authManagerClearCachedCredentialsFunction_Once sync.Once

func authManagerClearCachedCredentialsFunction_Set() error {
	var err error
	authManagerClearCachedCredentialsFunction_Once.Do(func() {
		err = authManagerObject_Set()
		if err != nil {
			return
		}
		authManagerClearCachedCredentialsFunction, err = authManagerObject.InvokerNew("clear_cached_credentials")
	})
	return err
}

// ClearCachedCredentials is a representation of the C type soup_auth_manager_clear_cached_credentials.
func (recv *AuthManager) ClearCachedCredentials() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := authManagerClearCachedCredentialsFunction_Set()
	if err == nil {
		authManagerClearCachedCredentialsFunction.Invoke(inArgs[:], nil)
	}

	return
}

var authManagerUseAuthFunction *gi.Function
var authManagerUseAuthFunction_Once sync.Once

func authManagerUseAuthFunction_Set() error {
	var err error
	authManagerUseAuthFunction_Once.Do(func() {
		err = authManagerObject_Set()
		if err != nil {
			return
		}
		authManagerUseAuthFunction, err = authManagerObject.InvokerNew("use_auth")
	})
	return err
}

// UseAuth is a representation of the C type soup_auth_manager_use_auth.
func (recv *AuthManager) UseAuth(uri *URI, auth *Auth) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(uri.Native())
	inArgs[2].SetPointer(auth.Native())

	err := authManagerUseAuthFunction_Set()
	if err == nil {
		authManagerUseAuthFunction.Invoke(inArgs[:], nil)
	}

	return
}

/*
ConnectAuthenticate connects a callback to the 'authenticate' signal of the AuthManager.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *AuthManager) ConnectAuthenticate(handler func(instance *AuthManager, msg *Message, auth *Auth, retrying bool)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {
		objectInstance := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[0]))
		argInstance := AuthManagerNewFromNative(objectInstance.GetObject().Native())

		object1 := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[1]))
		arg1 := MessageNewFromNative(object1.GetObject().Native())

		object2 := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[2]))
		arg2 := AuthNewFromNative(object2.GetObject().Native())

		object3 := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[3]))
		arg3 := object3.GetBoolean()

		handler(argInstance, arg1, arg2, arg3)

	}

	return callback.ConnectSignal(recv.Native(), "authenticate", marshal)
}

/*
Disconnect disconnects a callback previously registered with a Connect...() method.

The connectionID should be a value returned from a call to a Connect...() method.
*/
func (recv *AuthManager) DisconnectSignal(connectionID int) {
	callback.DisconnectSignal(connectionID)
}

// SessionFeature returns the SessionFeature interface implemented by AuthManager
func (recv *AuthManager) SessionFeature() *SessionFeature {
	return SessionFeatureNewFromNative(recv.Native())
}

var authNTLMObject *gi.Object
var authNTLMObject_Once sync.Once

func authNTLMObject_Set() error {
	var err error
	authNTLMObject_Once.Do(func() {
		authNTLMObject, err = gi.ObjectNew("Soup", "AuthNTLM")
	})
	return err
}

type AuthNTLM struct {
	native unsafe.Pointer
}

func AuthNTLMNewFromNative(native unsafe.Pointer) *AuthNTLM {
	err := authNTLMObject_Set()
	if err != nil {
		return nil
	}

	instance := &AuthNTLM{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Auth upcasts to *Auth
func (recv *AuthNTLM) Auth() *Auth {
	return AuthNewFromNative(recv.Native())
}

// Object upcasts to *Object
func (recv *AuthNTLM) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToAuthNTLM down casts any arbitrary Object to AuthNTLM.
Exercise care, as this is a potentially dangerous function
if the Object is not a AuthNTLM.
*/
func CastToAuthNTLM(object *gobject.Object) *AuthNTLM {
	return AuthNTLMNewFromNative(object.Native())
}

// Equals compares this AuthNTLM with another AuthNTLM, and returns true if they represent the same Object.
func (recv *AuthNTLM) Equals(other *AuthNTLM) bool {
	return other.Native() == recv.Native()
}

func (recv *AuthNTLM) Native() unsafe.Pointer {
	return recv.native
}

var authNegotiateObject *gi.Object
var authNegotiateObject_Once sync.Once

func authNegotiateObject_Set() error {
	var err error
	authNegotiateObject_Once.Do(func() {
		authNegotiateObject, err = gi.ObjectNew("Soup", "AuthNegotiate")
	})
	return err
}

type AuthNegotiate struct {
	native unsafe.Pointer
}

func AuthNegotiateNewFromNative(native unsafe.Pointer) *AuthNegotiate {
	err := authNegotiateObject_Set()
	if err != nil {
		return nil
	}

	instance := &AuthNegotiate{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Auth upcasts to *Auth
func (recv *AuthNegotiate) Auth() *Auth {
	return AuthNewFromNative(recv.Native())
}

// Object upcasts to *Object
func (recv *AuthNegotiate) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToAuthNegotiate down casts any arbitrary Object to AuthNegotiate.
Exercise care, as this is a potentially dangerous function
if the Object is not a AuthNegotiate.
*/
func CastToAuthNegotiate(object *gobject.Object) *AuthNegotiate {
	return AuthNegotiateNewFromNative(object.Native())
}

// Equals compares this AuthNegotiate with another AuthNegotiate, and returns true if they represent the same Object.
func (recv *AuthNegotiate) Equals(other *AuthNegotiate) bool {
	return other.Native() == recv.Native()
}

func (recv *AuthNegotiate) Native() unsafe.Pointer {
	return recv.native
}

var cacheObject *gi.Object
var cacheObject_Once sync.Once

func cacheObject_Set() error {
	var err error
	cacheObject_Once.Do(func() {
		cacheObject, err = gi.ObjectNew("Soup", "Cache")
	})
	return err
}

type Cache struct {
	native unsafe.Pointer
}

func CacheNewFromNative(native unsafe.Pointer) *Cache {
	err := cacheObject_Set()
	if err != nil {
		return nil
	}

	instance := &Cache{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *Cache) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToCache down casts any arbitrary Object to Cache.
Exercise care, as this is a potentially dangerous function
if the Object is not a Cache.
*/
func CastToCache(object *gobject.Object) *Cache {
	return CacheNewFromNative(object.Native())
}

// Equals compares this Cache with another Cache, and returns true if they represent the same Object.
func (recv *Cache) Equals(other *Cache) bool {
	return other.Native() == recv.Native()
}

func (recv *Cache) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *Cache) FieldParentInstance() *gobject.Object {
	argValue := gi.ObjectFieldGet(cacheObject, recv.Native(), "parent_instance")
	value := gobject.ObjectNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *Cache) SetFieldParentInstance(value *gobject.Object) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(cacheObject, recv.Native(), "parent_instance", argValue)
}

// FieldPriv returns the C field 'priv'.
func (recv *Cache) FieldPriv() *CachePrivate {
	argValue := gi.ObjectFieldGet(cacheObject, recv.Native(), "priv")
	value := CachePrivateNewFromNative(argValue.Pointer())
	return value
}

// SetFieldPriv sets the value of the C field 'priv'.
func (recv *Cache) SetFieldPriv(value *CachePrivate) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(cacheObject, recv.Native(), "priv", argValue)
}

var cacheNewFunction *gi.Function
var cacheNewFunction_Once sync.Once

func cacheNewFunction_Set() error {
	var err error
	cacheNewFunction_Once.Do(func() {
		err = cacheObject_Set()
		if err != nil {
			return
		}
		cacheNewFunction, err = cacheObject.InvokerNew("new")
	})
	return err
}

// CacheNew is a representation of the C type soup_cache_new.
func CacheNew(cacheDir string, cacheType CacheType) *Cache {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(cacheDir)
	inArgs[1].SetInt32(int32(cacheType))

	var ret gi.Argument

	err := cacheNewFunction_Set()
	if err == nil {
		ret = cacheNewFunction.Invoke(inArgs[:], nil)
	}

	retGo := CacheNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

var cacheClearFunction *gi.Function
var cacheClearFunction_Once sync.Once

func cacheClearFunction_Set() error {
	var err error
	cacheClearFunction_Once.Do(func() {
		err = cacheObject_Set()
		if err != nil {
			return
		}
		cacheClearFunction, err = cacheObject.InvokerNew("clear")
	})
	return err
}

// Clear is a representation of the C type soup_cache_clear.
func (recv *Cache) Clear() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := cacheClearFunction_Set()
	if err == nil {
		cacheClearFunction.Invoke(inArgs[:], nil)
	}

	return
}

var cacheDumpFunction *gi.Function
var cacheDumpFunction_Once sync.Once

func cacheDumpFunction_Set() error {
	var err error
	cacheDumpFunction_Once.Do(func() {
		err = cacheObject_Set()
		if err != nil {
			return
		}
		cacheDumpFunction, err = cacheObject.InvokerNew("dump")
	})
	return err
}

// Dump is a representation of the C type soup_cache_dump.
func (recv *Cache) Dump() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := cacheDumpFunction_Set()
	if err == nil {
		cacheDumpFunction.Invoke(inArgs[:], nil)
	}

	return
}

var cacheFlushFunction *gi.Function
var cacheFlushFunction_Once sync.Once

func cacheFlushFunction_Set() error {
	var err error
	cacheFlushFunction_Once.Do(func() {
		err = cacheObject_Set()
		if err != nil {
			return
		}
		cacheFlushFunction, err = cacheObject.InvokerNew("flush")
	})
	return err
}

// Flush is a representation of the C type soup_cache_flush.
func (recv *Cache) Flush() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := cacheFlushFunction_Set()
	if err == nil {
		cacheFlushFunction.Invoke(inArgs[:], nil)
	}

	return
}

var cacheGetMaxSizeFunction *gi.Function
var cacheGetMaxSizeFunction_Once sync.Once

func cacheGetMaxSizeFunction_Set() error {
	var err error
	cacheGetMaxSizeFunction_Once.Do(func() {
		err = cacheObject_Set()
		if err != nil {
			return
		}
		cacheGetMaxSizeFunction, err = cacheObject.InvokerNew("get_max_size")
	})
	return err
}

// GetMaxSize is a representation of the C type soup_cache_get_max_size.
func (recv *Cache) GetMaxSize() uint32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := cacheGetMaxSizeFunction_Set()
	if err == nil {
		ret = cacheGetMaxSizeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo
}

var cacheLoadFunction *gi.Function
var cacheLoadFunction_Once sync.Once

func cacheLoadFunction_Set() error {
	var err error
	cacheLoadFunction_Once.Do(func() {
		err = cacheObject_Set()
		if err != nil {
			return
		}
		cacheLoadFunction, err = cacheObject.InvokerNew("load")
	})
	return err
}

// Load is a representation of the C type soup_cache_load.
func (recv *Cache) Load() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := cacheLoadFunction_Set()
	if err == nil {
		cacheLoadFunction.Invoke(inArgs[:], nil)
	}

	return
}

var cacheSetMaxSizeFunction *gi.Function
var cacheSetMaxSizeFunction_Once sync.Once

func cacheSetMaxSizeFunction_Set() error {
	var err error
	cacheSetMaxSizeFunction_Once.Do(func() {
		err = cacheObject_Set()
		if err != nil {
			return
		}
		cacheSetMaxSizeFunction, err = cacheObject.InvokerNew("set_max_size")
	})
	return err
}

// SetMaxSize is a representation of the C type soup_cache_set_max_size.
func (recv *Cache) SetMaxSize(maxSize uint32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetUint32(maxSize)

	err := cacheSetMaxSizeFunction_Set()
	if err == nil {
		cacheSetMaxSizeFunction.Invoke(inArgs[:], nil)
	}

	return
}

// SessionFeature returns the SessionFeature interface implemented by Cache
func (recv *Cache) SessionFeature() *SessionFeature {
	return SessionFeatureNewFromNative(recv.Native())
}

var contentDecoderObject *gi.Object
var contentDecoderObject_Once sync.Once

func contentDecoderObject_Set() error {
	var err error
	contentDecoderObject_Once.Do(func() {
		contentDecoderObject, err = gi.ObjectNew("Soup", "ContentDecoder")
	})
	return err
}

type ContentDecoder struct {
	native unsafe.Pointer
}

func ContentDecoderNewFromNative(native unsafe.Pointer) *ContentDecoder {
	err := contentDecoderObject_Set()
	if err != nil {
		return nil
	}

	instance := &ContentDecoder{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *ContentDecoder) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToContentDecoder down casts any arbitrary Object to ContentDecoder.
Exercise care, as this is a potentially dangerous function
if the Object is not a ContentDecoder.
*/
func CastToContentDecoder(object *gobject.Object) *ContentDecoder {
	return ContentDecoderNewFromNative(object.Native())
}

// Equals compares this ContentDecoder with another ContentDecoder, and returns true if they represent the same Object.
func (recv *ContentDecoder) Equals(other *ContentDecoder) bool {
	return other.Native() == recv.Native()
}

func (recv *ContentDecoder) Native() unsafe.Pointer {
	return recv.native
}

// FieldParent returns the C field 'parent'.
func (recv *ContentDecoder) FieldParent() *gobject.Object {
	argValue := gi.ObjectFieldGet(contentDecoderObject, recv.Native(), "parent")
	value := gobject.ObjectNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParent sets the value of the C field 'parent'.
func (recv *ContentDecoder) SetFieldParent(value *gobject.Object) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(contentDecoderObject, recv.Native(), "parent", argValue)
}

// FieldPriv returns the C field 'priv'.
func (recv *ContentDecoder) FieldPriv() *ContentDecoderPrivate {
	argValue := gi.ObjectFieldGet(contentDecoderObject, recv.Native(), "priv")
	value := ContentDecoderPrivateNewFromNative(argValue.Pointer())
	return value
}

// SetFieldPriv sets the value of the C field 'priv'.
func (recv *ContentDecoder) SetFieldPriv(value *ContentDecoderPrivate) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(contentDecoderObject, recv.Native(), "priv", argValue)
}

// SessionFeature returns the SessionFeature interface implemented by ContentDecoder
func (recv *ContentDecoder) SessionFeature() *SessionFeature {
	return SessionFeatureNewFromNative(recv.Native())
}

var contentSnifferObject *gi.Object
var contentSnifferObject_Once sync.Once

func contentSnifferObject_Set() error {
	var err error
	contentSnifferObject_Once.Do(func() {
		contentSnifferObject, err = gi.ObjectNew("Soup", "ContentSniffer")
	})
	return err
}

type ContentSniffer struct {
	native unsafe.Pointer
}

func ContentSnifferNewFromNative(native unsafe.Pointer) *ContentSniffer {
	err := contentSnifferObject_Set()
	if err != nil {
		return nil
	}

	instance := &ContentSniffer{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *ContentSniffer) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToContentSniffer down casts any arbitrary Object to ContentSniffer.
Exercise care, as this is a potentially dangerous function
if the Object is not a ContentSniffer.
*/
func CastToContentSniffer(object *gobject.Object) *ContentSniffer {
	return ContentSnifferNewFromNative(object.Native())
}

// Equals compares this ContentSniffer with another ContentSniffer, and returns true if they represent the same Object.
func (recv *ContentSniffer) Equals(other *ContentSniffer) bool {
	return other.Native() == recv.Native()
}

func (recv *ContentSniffer) Native() unsafe.Pointer {
	return recv.native
}

// FieldParent returns the C field 'parent'.
func (recv *ContentSniffer) FieldParent() *gobject.Object {
	argValue := gi.ObjectFieldGet(contentSnifferObject, recv.Native(), "parent")
	value := gobject.ObjectNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParent sets the value of the C field 'parent'.
func (recv *ContentSniffer) SetFieldParent(value *gobject.Object) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(contentSnifferObject, recv.Native(), "parent", argValue)
}

// FieldPriv returns the C field 'priv'.
func (recv *ContentSniffer) FieldPriv() *ContentSnifferPrivate {
	argValue := gi.ObjectFieldGet(contentSnifferObject, recv.Native(), "priv")
	value := ContentSnifferPrivateNewFromNative(argValue.Pointer())
	return value
}

// SetFieldPriv sets the value of the C field 'priv'.
func (recv *ContentSniffer) SetFieldPriv(value *ContentSnifferPrivate) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(contentSnifferObject, recv.Native(), "priv", argValue)
}

var contentSnifferNewFunction *gi.Function
var contentSnifferNewFunction_Once sync.Once

func contentSnifferNewFunction_Set() error {
	var err error
	contentSnifferNewFunction_Once.Do(func() {
		err = contentSnifferObject_Set()
		if err != nil {
			return
		}
		contentSnifferNewFunction, err = contentSnifferObject.InvokerNew("new")
	})
	return err
}

// ContentSnifferNew is a representation of the C type soup_content_sniffer_new.
func ContentSnifferNew() *ContentSniffer {

	var ret gi.Argument

	err := contentSnifferNewFunction_Set()
	if err == nil {
		ret = contentSnifferNewFunction.Invoke(nil, nil)
	}

	retGo := ContentSnifferNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

var contentSnifferGetBufferSizeFunction *gi.Function
var contentSnifferGetBufferSizeFunction_Once sync.Once

func contentSnifferGetBufferSizeFunction_Set() error {
	var err error
	contentSnifferGetBufferSizeFunction_Once.Do(func() {
		err = contentSnifferObject_Set()
		if err != nil {
			return
		}
		contentSnifferGetBufferSizeFunction, err = contentSnifferObject.InvokerNew("get_buffer_size")
	})
	return err
}

// GetBufferSize is a representation of the C type soup_content_sniffer_get_buffer_size.
func (recv *ContentSniffer) GetBufferSize() uint64 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := contentSnifferGetBufferSizeFunction_Set()
	if err == nil {
		ret = contentSnifferGetBufferSizeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint64()

	return retGo
}

var contentSnifferSniffFunction *gi.Function
var contentSnifferSniffFunction_Once sync.Once

func contentSnifferSniffFunction_Set() error {
	var err error
	contentSnifferSniffFunction_Once.Do(func() {
		err = contentSnifferObject_Set()
		if err != nil {
			return
		}
		contentSnifferSniffFunction, err = contentSnifferObject.InvokerNew("sniff")
	})
	return err
}

// Sniff is a representation of the C type soup_content_sniffer_sniff.
func (recv *ContentSniffer) Sniff(msg *Message, buffer *Buffer) (string, *glib.HashTable) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(msg.Native())
	inArgs[2].SetPointer(buffer.Native())

	var outArgs [1]gi.Argument
	var ret gi.Argument

	err := contentSnifferSniffFunction_Set()
	if err == nil {
		ret = contentSnifferSniffFunction.Invoke(inArgs[:], outArgs[:])
	}

	retGo := ret.String(true)
	out0 := glib.HashTableNewFromNative(outArgs[0].Pointer())

	return retGo, out0
}

// SessionFeature returns the SessionFeature interface implemented by ContentSniffer
func (recv *ContentSniffer) SessionFeature() *SessionFeature {
	return SessionFeatureNewFromNative(recv.Native())
}

var cookieJarObject *gi.Object
var cookieJarObject_Once sync.Once

func cookieJarObject_Set() error {
	var err error
	cookieJarObject_Once.Do(func() {
		cookieJarObject, err = gi.ObjectNew("Soup", "CookieJar")
	})
	return err
}

type CookieJar struct {
	native unsafe.Pointer
}

func CookieJarNewFromNative(native unsafe.Pointer) *CookieJar {
	err := cookieJarObject_Set()
	if err != nil {
		return nil
	}

	instance := &CookieJar{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *CookieJar) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToCookieJar down casts any arbitrary Object to CookieJar.
Exercise care, as this is a potentially dangerous function
if the Object is not a CookieJar.
*/
func CastToCookieJar(object *gobject.Object) *CookieJar {
	return CookieJarNewFromNative(object.Native())
}

// Equals compares this CookieJar with another CookieJar, and returns true if they represent the same Object.
func (recv *CookieJar) Equals(other *CookieJar) bool {
	return other.Native() == recv.Native()
}

func (recv *CookieJar) Native() unsafe.Pointer {
	return recv.native
}

// FieldParent returns the C field 'parent'.
func (recv *CookieJar) FieldParent() *gobject.Object {
	argValue := gi.ObjectFieldGet(cookieJarObject, recv.Native(), "parent")
	value := gobject.ObjectNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParent sets the value of the C field 'parent'.
func (recv *CookieJar) SetFieldParent(value *gobject.Object) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(cookieJarObject, recv.Native(), "parent", argValue)
}

var cookieJarNewFunction *gi.Function
var cookieJarNewFunction_Once sync.Once

func cookieJarNewFunction_Set() error {
	var err error
	cookieJarNewFunction_Once.Do(func() {
		err = cookieJarObject_Set()
		if err != nil {
			return
		}
		cookieJarNewFunction, err = cookieJarObject.InvokerNew("new")
	})
	return err
}

// CookieJarNew is a representation of the C type soup_cookie_jar_new.
func CookieJarNew() *CookieJar {

	var ret gi.Argument

	err := cookieJarNewFunction_Set()
	if err == nil {
		ret = cookieJarNewFunction.Invoke(nil, nil)
	}

	retGo := CookieJarNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

var cookieJarAddCookieFunction *gi.Function
var cookieJarAddCookieFunction_Once sync.Once

func cookieJarAddCookieFunction_Set() error {
	var err error
	cookieJarAddCookieFunction_Once.Do(func() {
		err = cookieJarObject_Set()
		if err != nil {
			return
		}
		cookieJarAddCookieFunction, err = cookieJarObject.InvokerNew("add_cookie")
	})
	return err
}

// AddCookie is a representation of the C type soup_cookie_jar_add_cookie.
func (recv *CookieJar) AddCookie(cookie *Cookie) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(cookie.Native())

	err := cookieJarAddCookieFunction_Set()
	if err == nil {
		cookieJarAddCookieFunction.Invoke(inArgs[:], nil)
	}

	return
}

var cookieJarAddCookieFullFunction *gi.Function
var cookieJarAddCookieFullFunction_Once sync.Once

func cookieJarAddCookieFullFunction_Set() error {
	var err error
	cookieJarAddCookieFullFunction_Once.Do(func() {
		err = cookieJarObject_Set()
		if err != nil {
			return
		}
		cookieJarAddCookieFullFunction, err = cookieJarObject.InvokerNew("add_cookie_full")
	})
	return err
}

// AddCookieFull is a representation of the C type soup_cookie_jar_add_cookie_full.
func (recv *CookieJar) AddCookieFull(cookie *Cookie, uri *URI, firstParty *URI) {
	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(cookie.Native())
	inArgs[2].SetPointer(uri.Native())
	inArgs[3].SetPointer(firstParty.Native())

	err := cookieJarAddCookieFullFunction_Set()
	if err == nil {
		cookieJarAddCookieFullFunction.Invoke(inArgs[:], nil)
	}

	return
}

var cookieJarAddCookieWithFirstPartyFunction *gi.Function
var cookieJarAddCookieWithFirstPartyFunction_Once sync.Once

func cookieJarAddCookieWithFirstPartyFunction_Set() error {
	var err error
	cookieJarAddCookieWithFirstPartyFunction_Once.Do(func() {
		err = cookieJarObject_Set()
		if err != nil {
			return
		}
		cookieJarAddCookieWithFirstPartyFunction, err = cookieJarObject.InvokerNew("add_cookie_with_first_party")
	})
	return err
}

// AddCookieWithFirstParty is a representation of the C type soup_cookie_jar_add_cookie_with_first_party.
func (recv *CookieJar) AddCookieWithFirstParty(firstParty *URI, cookie *Cookie) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(firstParty.Native())
	inArgs[2].SetPointer(cookie.Native())

	err := cookieJarAddCookieWithFirstPartyFunction_Set()
	if err == nil {
		cookieJarAddCookieWithFirstPartyFunction.Invoke(inArgs[:], nil)
	}

	return
}

var cookieJarAllCookiesFunction *gi.Function
var cookieJarAllCookiesFunction_Once sync.Once

func cookieJarAllCookiesFunction_Set() error {
	var err error
	cookieJarAllCookiesFunction_Once.Do(func() {
		err = cookieJarObject_Set()
		if err != nil {
			return
		}
		cookieJarAllCookiesFunction, err = cookieJarObject.InvokerNew("all_cookies")
	})
	return err
}

// AllCookies is a representation of the C type soup_cookie_jar_all_cookies.
func (recv *CookieJar) AllCookies() *glib.SList {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := cookieJarAllCookiesFunction_Set()
	if err == nil {
		ret = cookieJarAllCookiesFunction.Invoke(inArgs[:], nil)
	}

	retGo := glib.SListNewFromNative(ret.Pointer())

	return retGo
}

var cookieJarDeleteCookieFunction *gi.Function
var cookieJarDeleteCookieFunction_Once sync.Once

func cookieJarDeleteCookieFunction_Set() error {
	var err error
	cookieJarDeleteCookieFunction_Once.Do(func() {
		err = cookieJarObject_Set()
		if err != nil {
			return
		}
		cookieJarDeleteCookieFunction, err = cookieJarObject.InvokerNew("delete_cookie")
	})
	return err
}

// DeleteCookie is a representation of the C type soup_cookie_jar_delete_cookie.
func (recv *CookieJar) DeleteCookie(cookie *Cookie) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(cookie.Native())

	err := cookieJarDeleteCookieFunction_Set()
	if err == nil {
		cookieJarDeleteCookieFunction.Invoke(inArgs[:], nil)
	}

	return
}

var cookieJarGetAcceptPolicyFunction *gi.Function
var cookieJarGetAcceptPolicyFunction_Once sync.Once

func cookieJarGetAcceptPolicyFunction_Set() error {
	var err error
	cookieJarGetAcceptPolicyFunction_Once.Do(func() {
		err = cookieJarObject_Set()
		if err != nil {
			return
		}
		cookieJarGetAcceptPolicyFunction, err = cookieJarObject.InvokerNew("get_accept_policy")
	})
	return err
}

// GetAcceptPolicy is a representation of the C type soup_cookie_jar_get_accept_policy.
func (recv *CookieJar) GetAcceptPolicy() CookieJarAcceptPolicy {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := cookieJarGetAcceptPolicyFunction_Set()
	if err == nil {
		ret = cookieJarGetAcceptPolicyFunction.Invoke(inArgs[:], nil)
	}

	retGo := CookieJarAcceptPolicy(ret.Int32())

	return retGo
}

var cookieJarGetCookieListFunction *gi.Function
var cookieJarGetCookieListFunction_Once sync.Once

func cookieJarGetCookieListFunction_Set() error {
	var err error
	cookieJarGetCookieListFunction_Once.Do(func() {
		err = cookieJarObject_Set()
		if err != nil {
			return
		}
		cookieJarGetCookieListFunction, err = cookieJarObject.InvokerNew("get_cookie_list")
	})
	return err
}

// GetCookieList is a representation of the C type soup_cookie_jar_get_cookie_list.
func (recv *CookieJar) GetCookieList(uri *URI, forHttp bool) *glib.SList {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(uri.Native())
	inArgs[2].SetBoolean(forHttp)

	var ret gi.Argument

	err := cookieJarGetCookieListFunction_Set()
	if err == nil {
		ret = cookieJarGetCookieListFunction.Invoke(inArgs[:], nil)
	}

	retGo := glib.SListNewFromNative(ret.Pointer())

	return retGo
}

var cookieJarGetCookiesFunction *gi.Function
var cookieJarGetCookiesFunction_Once sync.Once

func cookieJarGetCookiesFunction_Set() error {
	var err error
	cookieJarGetCookiesFunction_Once.Do(func() {
		err = cookieJarObject_Set()
		if err != nil {
			return
		}
		cookieJarGetCookiesFunction, err = cookieJarObject.InvokerNew("get_cookies")
	})
	return err
}

// GetCookies is a representation of the C type soup_cookie_jar_get_cookies.
func (recv *CookieJar) GetCookies(uri *URI, forHttp bool) string {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(uri.Native())
	inArgs[2].SetBoolean(forHttp)

	var ret gi.Argument

	err := cookieJarGetCookiesFunction_Set()
	if err == nil {
		ret = cookieJarGetCookiesFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

var cookieJarIsPersistentFunction *gi.Function
var cookieJarIsPersistentFunction_Once sync.Once

func cookieJarIsPersistentFunction_Set() error {
	var err error
	cookieJarIsPersistentFunction_Once.Do(func() {
		err = cookieJarObject_Set()
		if err != nil {
			return
		}
		cookieJarIsPersistentFunction, err = cookieJarObject.InvokerNew("is_persistent")
	})
	return err
}

// IsPersistent is a representation of the C type soup_cookie_jar_is_persistent.
func (recv *CookieJar) IsPersistent() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := cookieJarIsPersistentFunction_Set()
	if err == nil {
		ret = cookieJarIsPersistentFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var cookieJarSaveFunction *gi.Function
var cookieJarSaveFunction_Once sync.Once

func cookieJarSaveFunction_Set() error {
	var err error
	cookieJarSaveFunction_Once.Do(func() {
		err = cookieJarObject_Set()
		if err != nil {
			return
		}
		cookieJarSaveFunction, err = cookieJarObject.InvokerNew("save")
	})
	return err
}

// Save is a representation of the C type soup_cookie_jar_save.
func (recv *CookieJar) Save() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := cookieJarSaveFunction_Set()
	if err == nil {
		cookieJarSaveFunction.Invoke(inArgs[:], nil)
	}

	return
}

var cookieJarSetAcceptPolicyFunction *gi.Function
var cookieJarSetAcceptPolicyFunction_Once sync.Once

func cookieJarSetAcceptPolicyFunction_Set() error {
	var err error
	cookieJarSetAcceptPolicyFunction_Once.Do(func() {
		err = cookieJarObject_Set()
		if err != nil {
			return
		}
		cookieJarSetAcceptPolicyFunction, err = cookieJarObject.InvokerNew("set_accept_policy")
	})
	return err
}

// SetAcceptPolicy is a representation of the C type soup_cookie_jar_set_accept_policy.
func (recv *CookieJar) SetAcceptPolicy(policy CookieJarAcceptPolicy) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(int32(policy))

	err := cookieJarSetAcceptPolicyFunction_Set()
	if err == nil {
		cookieJarSetAcceptPolicyFunction.Invoke(inArgs[:], nil)
	}

	return
}

var cookieJarSetCookieFunction *gi.Function
var cookieJarSetCookieFunction_Once sync.Once

func cookieJarSetCookieFunction_Set() error {
	var err error
	cookieJarSetCookieFunction_Once.Do(func() {
		err = cookieJarObject_Set()
		if err != nil {
			return
		}
		cookieJarSetCookieFunction, err = cookieJarObject.InvokerNew("set_cookie")
	})
	return err
}

// SetCookie is a representation of the C type soup_cookie_jar_set_cookie.
func (recv *CookieJar) SetCookie(uri *URI, cookie string) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(uri.Native())
	inArgs[2].SetString(cookie)

	err := cookieJarSetCookieFunction_Set()
	if err == nil {
		cookieJarSetCookieFunction.Invoke(inArgs[:], nil)
	}

	return
}

var cookieJarSetCookieWithFirstPartyFunction *gi.Function
var cookieJarSetCookieWithFirstPartyFunction_Once sync.Once

func cookieJarSetCookieWithFirstPartyFunction_Set() error {
	var err error
	cookieJarSetCookieWithFirstPartyFunction_Once.Do(func() {
		err = cookieJarObject_Set()
		if err != nil {
			return
		}
		cookieJarSetCookieWithFirstPartyFunction, err = cookieJarObject.InvokerNew("set_cookie_with_first_party")
	})
	return err
}

// SetCookieWithFirstParty is a representation of the C type soup_cookie_jar_set_cookie_with_first_party.
func (recv *CookieJar) SetCookieWithFirstParty(uri *URI, firstParty *URI, cookie string) {
	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(uri.Native())
	inArgs[2].SetPointer(firstParty.Native())
	inArgs[3].SetString(cookie)

	err := cookieJarSetCookieWithFirstPartyFunction_Set()
	if err == nil {
		cookieJarSetCookieWithFirstPartyFunction.Invoke(inArgs[:], nil)
	}

	return
}

/*
ConnectChanged connects a callback to the 'changed' signal of the CookieJar.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *CookieJar) ConnectChanged(handler func(instance *CookieJar, oldCookie *Cookie, newCookie *Cookie)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {
		objectInstance := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[0]))
		argInstance := CookieJarNewFromNative(objectInstance.GetObject().Native())

		object1 := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[1]))
		arg1 := CookieNewFromNative(object1.GetBoxed())

		object2 := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[2]))
		arg2 := CookieNewFromNative(object2.GetBoxed())

		handler(argInstance, arg1, arg2)

	}

	return callback.ConnectSignal(recv.Native(), "changed", marshal)
}

/*
Disconnect disconnects a callback previously registered with a Connect...() method.

The connectionID should be a value returned from a call to a Connect...() method.
*/
func (recv *CookieJar) DisconnectSignal(connectionID int) {
	callback.DisconnectSignal(connectionID)
}

// SessionFeature returns the SessionFeature interface implemented by CookieJar
func (recv *CookieJar) SessionFeature() *SessionFeature {
	return SessionFeatureNewFromNative(recv.Native())
}

var cookieJarDBObject *gi.Object
var cookieJarDBObject_Once sync.Once

func cookieJarDBObject_Set() error {
	var err error
	cookieJarDBObject_Once.Do(func() {
		cookieJarDBObject, err = gi.ObjectNew("Soup", "CookieJarDB")
	})
	return err
}

type CookieJarDB struct {
	native unsafe.Pointer
}

func CookieJarDBNewFromNative(native unsafe.Pointer) *CookieJarDB {
	err := cookieJarDBObject_Set()
	if err != nil {
		return nil
	}

	instance := &CookieJarDB{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// CookieJar upcasts to *CookieJar
func (recv *CookieJarDB) CookieJar() *CookieJar {
	return CookieJarNewFromNative(recv.Native())
}

// Object upcasts to *Object
func (recv *CookieJarDB) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToCookieJarDB down casts any arbitrary Object to CookieJarDB.
Exercise care, as this is a potentially dangerous function
if the Object is not a CookieJarDB.
*/
func CastToCookieJarDB(object *gobject.Object) *CookieJarDB {
	return CookieJarDBNewFromNative(object.Native())
}

// Equals compares this CookieJarDB with another CookieJarDB, and returns true if they represent the same Object.
func (recv *CookieJarDB) Equals(other *CookieJarDB) bool {
	return other.Native() == recv.Native()
}

func (recv *CookieJarDB) Native() unsafe.Pointer {
	return recv.native
}

// FieldParent returns the C field 'parent'.
func (recv *CookieJarDB) FieldParent() *CookieJar {
	argValue := gi.ObjectFieldGet(cookieJarDBObject, recv.Native(), "parent")
	value := CookieJarNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParent sets the value of the C field 'parent'.
func (recv *CookieJarDB) SetFieldParent(value *CookieJar) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(cookieJarDBObject, recv.Native(), "parent", argValue)
}

var cookieJarDBNewFunction *gi.Function
var cookieJarDBNewFunction_Once sync.Once

func cookieJarDBNewFunction_Set() error {
	var err error
	cookieJarDBNewFunction_Once.Do(func() {
		err = cookieJarDBObject_Set()
		if err != nil {
			return
		}
		cookieJarDBNewFunction, err = cookieJarDBObject.InvokerNew("new")
	})
	return err
}

// CookieJarDBNew is a representation of the C type soup_cookie_jar_db_new.
func CookieJarDBNew(filename string, readOnly bool) *CookieJarDB {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(filename)
	inArgs[1].SetBoolean(readOnly)

	var ret gi.Argument

	err := cookieJarDBNewFunction_Set()
	if err == nil {
		ret = cookieJarDBNewFunction.Invoke(inArgs[:], nil)
	}

	retGo := CookieJarDBNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

// SessionFeature returns the SessionFeature interface implemented by CookieJarDB
func (recv *CookieJarDB) SessionFeature() *SessionFeature {
	return SessionFeatureNewFromNative(recv.Native())
}

var cookieJarTextObject *gi.Object
var cookieJarTextObject_Once sync.Once

func cookieJarTextObject_Set() error {
	var err error
	cookieJarTextObject_Once.Do(func() {
		cookieJarTextObject, err = gi.ObjectNew("Soup", "CookieJarText")
	})
	return err
}

type CookieJarText struct {
	native unsafe.Pointer
}

func CookieJarTextNewFromNative(native unsafe.Pointer) *CookieJarText {
	err := cookieJarTextObject_Set()
	if err != nil {
		return nil
	}

	instance := &CookieJarText{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// CookieJar upcasts to *CookieJar
func (recv *CookieJarText) CookieJar() *CookieJar {
	return CookieJarNewFromNative(recv.Native())
}

// Object upcasts to *Object
func (recv *CookieJarText) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToCookieJarText down casts any arbitrary Object to CookieJarText.
Exercise care, as this is a potentially dangerous function
if the Object is not a CookieJarText.
*/
func CastToCookieJarText(object *gobject.Object) *CookieJarText {
	return CookieJarTextNewFromNative(object.Native())
}

// Equals compares this CookieJarText with another CookieJarText, and returns true if they represent the same Object.
func (recv *CookieJarText) Equals(other *CookieJarText) bool {
	return other.Native() == recv.Native()
}

func (recv *CookieJarText) Native() unsafe.Pointer {
	return recv.native
}

// FieldParent returns the C field 'parent'.
func (recv *CookieJarText) FieldParent() *CookieJar {
	argValue := gi.ObjectFieldGet(cookieJarTextObject, recv.Native(), "parent")
	value := CookieJarNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParent sets the value of the C field 'parent'.
func (recv *CookieJarText) SetFieldParent(value *CookieJar) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(cookieJarTextObject, recv.Native(), "parent", argValue)
}

var cookieJarTextNewFunction *gi.Function
var cookieJarTextNewFunction_Once sync.Once

func cookieJarTextNewFunction_Set() error {
	var err error
	cookieJarTextNewFunction_Once.Do(func() {
		err = cookieJarTextObject_Set()
		if err != nil {
			return
		}
		cookieJarTextNewFunction, err = cookieJarTextObject.InvokerNew("new")
	})
	return err
}

// CookieJarTextNew is a representation of the C type soup_cookie_jar_text_new.
func CookieJarTextNew(filename string, readOnly bool) *CookieJarText {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(filename)
	inArgs[1].SetBoolean(readOnly)

	var ret gi.Argument

	err := cookieJarTextNewFunction_Set()
	if err == nil {
		ret = cookieJarTextNewFunction.Invoke(inArgs[:], nil)
	}

	retGo := CookieJarTextNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

// SessionFeature returns the SessionFeature interface implemented by CookieJarText
func (recv *CookieJarText) SessionFeature() *SessionFeature {
	return SessionFeatureNewFromNative(recv.Native())
}

var hSTSEnforcerObject *gi.Object
var hSTSEnforcerObject_Once sync.Once

func hSTSEnforcerObject_Set() error {
	var err error
	hSTSEnforcerObject_Once.Do(func() {
		hSTSEnforcerObject, err = gi.ObjectNew("Soup", "HSTSEnforcer")
	})
	return err
}

type HSTSEnforcer struct {
	native unsafe.Pointer
}

func HSTSEnforcerNewFromNative(native unsafe.Pointer) *HSTSEnforcer {
	err := hSTSEnforcerObject_Set()
	if err != nil {
		return nil
	}

	instance := &HSTSEnforcer{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *HSTSEnforcer) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToHSTSEnforcer down casts any arbitrary Object to HSTSEnforcer.
Exercise care, as this is a potentially dangerous function
if the Object is not a HSTSEnforcer.
*/
func CastToHSTSEnforcer(object *gobject.Object) *HSTSEnforcer {
	return HSTSEnforcerNewFromNative(object.Native())
}

// Equals compares this HSTSEnforcer with another HSTSEnforcer, and returns true if they represent the same Object.
func (recv *HSTSEnforcer) Equals(other *HSTSEnforcer) bool {
	return other.Native() == recv.Native()
}

func (recv *HSTSEnforcer) Native() unsafe.Pointer {
	return recv.native
}

// FieldParent returns the C field 'parent'.
func (recv *HSTSEnforcer) FieldParent() *gobject.Object {
	argValue := gi.ObjectFieldGet(hSTSEnforcerObject, recv.Native(), "parent")
	value := gobject.ObjectNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParent sets the value of the C field 'parent'.
func (recv *HSTSEnforcer) SetFieldParent(value *gobject.Object) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(hSTSEnforcerObject, recv.Native(), "parent", argValue)
}

// FieldPriv returns the C field 'priv'.
func (recv *HSTSEnforcer) FieldPriv() *HSTSEnforcerPrivate {
	argValue := gi.ObjectFieldGet(hSTSEnforcerObject, recv.Native(), "priv")
	value := HSTSEnforcerPrivateNewFromNative(argValue.Pointer())
	return value
}

// SetFieldPriv sets the value of the C field 'priv'.
func (recv *HSTSEnforcer) SetFieldPriv(value *HSTSEnforcerPrivate) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(hSTSEnforcerObject, recv.Native(), "priv", argValue)
}

var hSTSEnforcerNewFunction *gi.Function
var hSTSEnforcerNewFunction_Once sync.Once

func hSTSEnforcerNewFunction_Set() error {
	var err error
	hSTSEnforcerNewFunction_Once.Do(func() {
		err = hSTSEnforcerObject_Set()
		if err != nil {
			return
		}
		hSTSEnforcerNewFunction, err = hSTSEnforcerObject.InvokerNew("new")
	})
	return err
}

// HSTSEnforcerNew is a representation of the C type soup_hsts_enforcer_new.
func HSTSEnforcerNew() *HSTSEnforcer {

	var ret gi.Argument

	err := hSTSEnforcerNewFunction_Set()
	if err == nil {
		ret = hSTSEnforcerNewFunction.Invoke(nil, nil)
	}

	retGo := HSTSEnforcerNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

var hSTSEnforcerGetDomainsFunction *gi.Function
var hSTSEnforcerGetDomainsFunction_Once sync.Once

func hSTSEnforcerGetDomainsFunction_Set() error {
	var err error
	hSTSEnforcerGetDomainsFunction_Once.Do(func() {
		err = hSTSEnforcerObject_Set()
		if err != nil {
			return
		}
		hSTSEnforcerGetDomainsFunction, err = hSTSEnforcerObject.InvokerNew("get_domains")
	})
	return err
}

// GetDomains is a representation of the C type soup_hsts_enforcer_get_domains.
func (recv *HSTSEnforcer) GetDomains(sessionPolicies bool) *glib.List {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetBoolean(sessionPolicies)

	var ret gi.Argument

	err := hSTSEnforcerGetDomainsFunction_Set()
	if err == nil {
		ret = hSTSEnforcerGetDomainsFunction.Invoke(inArgs[:], nil)
	}

	retGo := glib.ListNewFromNative(ret.Pointer())

	return retGo
}

var hSTSEnforcerGetPoliciesFunction *gi.Function
var hSTSEnforcerGetPoliciesFunction_Once sync.Once

func hSTSEnforcerGetPoliciesFunction_Set() error {
	var err error
	hSTSEnforcerGetPoliciesFunction_Once.Do(func() {
		err = hSTSEnforcerObject_Set()
		if err != nil {
			return
		}
		hSTSEnforcerGetPoliciesFunction, err = hSTSEnforcerObject.InvokerNew("get_policies")
	})
	return err
}

// GetPolicies is a representation of the C type soup_hsts_enforcer_get_policies.
func (recv *HSTSEnforcer) GetPolicies(sessionPolicies bool) *glib.List {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetBoolean(sessionPolicies)

	var ret gi.Argument

	err := hSTSEnforcerGetPoliciesFunction_Set()
	if err == nil {
		ret = hSTSEnforcerGetPoliciesFunction.Invoke(inArgs[:], nil)
	}

	retGo := glib.ListNewFromNative(ret.Pointer())

	return retGo
}

var hSTSEnforcerHasValidPolicyFunction *gi.Function
var hSTSEnforcerHasValidPolicyFunction_Once sync.Once

func hSTSEnforcerHasValidPolicyFunction_Set() error {
	var err error
	hSTSEnforcerHasValidPolicyFunction_Once.Do(func() {
		err = hSTSEnforcerObject_Set()
		if err != nil {
			return
		}
		hSTSEnforcerHasValidPolicyFunction, err = hSTSEnforcerObject.InvokerNew("has_valid_policy")
	})
	return err
}

// HasValidPolicy is a representation of the C type soup_hsts_enforcer_has_valid_policy.
func (recv *HSTSEnforcer) HasValidPolicy(domain string) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(domain)

	var ret gi.Argument

	err := hSTSEnforcerHasValidPolicyFunction_Set()
	if err == nil {
		ret = hSTSEnforcerHasValidPolicyFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var hSTSEnforcerIsPersistentFunction *gi.Function
var hSTSEnforcerIsPersistentFunction_Once sync.Once

func hSTSEnforcerIsPersistentFunction_Set() error {
	var err error
	hSTSEnforcerIsPersistentFunction_Once.Do(func() {
		err = hSTSEnforcerObject_Set()
		if err != nil {
			return
		}
		hSTSEnforcerIsPersistentFunction, err = hSTSEnforcerObject.InvokerNew("is_persistent")
	})
	return err
}

// IsPersistent is a representation of the C type soup_hsts_enforcer_is_persistent.
func (recv *HSTSEnforcer) IsPersistent() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := hSTSEnforcerIsPersistentFunction_Set()
	if err == nil {
		ret = hSTSEnforcerIsPersistentFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var hSTSEnforcerSetPolicyFunction *gi.Function
var hSTSEnforcerSetPolicyFunction_Once sync.Once

func hSTSEnforcerSetPolicyFunction_Set() error {
	var err error
	hSTSEnforcerSetPolicyFunction_Once.Do(func() {
		err = hSTSEnforcerObject_Set()
		if err != nil {
			return
		}
		hSTSEnforcerSetPolicyFunction, err = hSTSEnforcerObject.InvokerNew("set_policy")
	})
	return err
}

// SetPolicy is a representation of the C type soup_hsts_enforcer_set_policy.
func (recv *HSTSEnforcer) SetPolicy(policy *HSTSPolicy) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(policy.Native())

	err := hSTSEnforcerSetPolicyFunction_Set()
	if err == nil {
		hSTSEnforcerSetPolicyFunction.Invoke(inArgs[:], nil)
	}

	return
}

var hSTSEnforcerSetSessionPolicyFunction *gi.Function
var hSTSEnforcerSetSessionPolicyFunction_Once sync.Once

func hSTSEnforcerSetSessionPolicyFunction_Set() error {
	var err error
	hSTSEnforcerSetSessionPolicyFunction_Once.Do(func() {
		err = hSTSEnforcerObject_Set()
		if err != nil {
			return
		}
		hSTSEnforcerSetSessionPolicyFunction, err = hSTSEnforcerObject.InvokerNew("set_session_policy")
	})
	return err
}

// SetSessionPolicy is a representation of the C type soup_hsts_enforcer_set_session_policy.
func (recv *HSTSEnforcer) SetSessionPolicy(domain string, includeSubdomains bool) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(domain)
	inArgs[2].SetBoolean(includeSubdomains)

	err := hSTSEnforcerSetSessionPolicyFunction_Set()
	if err == nil {
		hSTSEnforcerSetSessionPolicyFunction.Invoke(inArgs[:], nil)
	}

	return
}

/*
ConnectChanged connects a callback to the 'changed' signal of the HSTSEnforcer.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *HSTSEnforcer) ConnectChanged(handler func(instance *HSTSEnforcer, oldPolicy *HSTSPolicy, newPolicy *HSTSPolicy)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {
		objectInstance := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[0]))
		argInstance := HSTSEnforcerNewFromNative(objectInstance.GetObject().Native())

		object1 := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[1]))
		arg1 := HSTSPolicyNewFromNative(object1.GetBoxed())

		object2 := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[2]))
		arg2 := HSTSPolicyNewFromNative(object2.GetBoxed())

		handler(argInstance, arg1, arg2)

	}

	return callback.ConnectSignal(recv.Native(), "changed", marshal)
}

/*
ConnectHstsEnforced connects a callback to the 'hsts-enforced' signal of the HSTSEnforcer.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *HSTSEnforcer) ConnectHstsEnforced(handler func(instance *HSTSEnforcer, message *Message)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {
		objectInstance := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[0]))
		argInstance := HSTSEnforcerNewFromNative(objectInstance.GetObject().Native())

		object1 := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[1]))
		arg1 := MessageNewFromNative(object1.GetObject().Native())

		handler(argInstance, arg1)

	}

	return callback.ConnectSignal(recv.Native(), "hsts-enforced", marshal)
}

/*
Disconnect disconnects a callback previously registered with a Connect...() method.

The connectionID should be a value returned from a call to a Connect...() method.
*/
func (recv *HSTSEnforcer) DisconnectSignal(connectionID int) {
	callback.DisconnectSignal(connectionID)
}

// SessionFeature returns the SessionFeature interface implemented by HSTSEnforcer
func (recv *HSTSEnforcer) SessionFeature() *SessionFeature {
	return SessionFeatureNewFromNative(recv.Native())
}

var hSTSEnforcerDBObject *gi.Object
var hSTSEnforcerDBObject_Once sync.Once

func hSTSEnforcerDBObject_Set() error {
	var err error
	hSTSEnforcerDBObject_Once.Do(func() {
		hSTSEnforcerDBObject, err = gi.ObjectNew("Soup", "HSTSEnforcerDB")
	})
	return err
}

type HSTSEnforcerDB struct {
	native unsafe.Pointer
}

func HSTSEnforcerDBNewFromNative(native unsafe.Pointer) *HSTSEnforcerDB {
	err := hSTSEnforcerDBObject_Set()
	if err != nil {
		return nil
	}

	instance := &HSTSEnforcerDB{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// HSTSEnforcer upcasts to *HSTSEnforcer
func (recv *HSTSEnforcerDB) HSTSEnforcer() *HSTSEnforcer {
	return HSTSEnforcerNewFromNative(recv.Native())
}

// Object upcasts to *Object
func (recv *HSTSEnforcerDB) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToHSTSEnforcerDB down casts any arbitrary Object to HSTSEnforcerDB.
Exercise care, as this is a potentially dangerous function
if the Object is not a HSTSEnforcerDB.
*/
func CastToHSTSEnforcerDB(object *gobject.Object) *HSTSEnforcerDB {
	return HSTSEnforcerDBNewFromNative(object.Native())
}

// Equals compares this HSTSEnforcerDB with another HSTSEnforcerDB, and returns true if they represent the same Object.
func (recv *HSTSEnforcerDB) Equals(other *HSTSEnforcerDB) bool {
	return other.Native() == recv.Native()
}

func (recv *HSTSEnforcerDB) Native() unsafe.Pointer {
	return recv.native
}

// FieldParent returns the C field 'parent'.
func (recv *HSTSEnforcerDB) FieldParent() *HSTSEnforcer {
	argValue := gi.ObjectFieldGet(hSTSEnforcerDBObject, recv.Native(), "parent")
	value := HSTSEnforcerNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParent sets the value of the C field 'parent'.
func (recv *HSTSEnforcerDB) SetFieldParent(value *HSTSEnforcer) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(hSTSEnforcerDBObject, recv.Native(), "parent", argValue)
}

// FieldPriv returns the C field 'priv'.
func (recv *HSTSEnforcerDB) FieldPriv() *HSTSEnforcerDBPrivate {
	argValue := gi.ObjectFieldGet(hSTSEnforcerDBObject, recv.Native(), "priv")
	value := HSTSEnforcerDBPrivateNewFromNative(argValue.Pointer())
	return value
}

// SetFieldPriv sets the value of the C field 'priv'.
func (recv *HSTSEnforcerDB) SetFieldPriv(value *HSTSEnforcerDBPrivate) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(hSTSEnforcerDBObject, recv.Native(), "priv", argValue)
}

var hSTSEnforcerDBNewFunction *gi.Function
var hSTSEnforcerDBNewFunction_Once sync.Once

func hSTSEnforcerDBNewFunction_Set() error {
	var err error
	hSTSEnforcerDBNewFunction_Once.Do(func() {
		err = hSTSEnforcerDBObject_Set()
		if err != nil {
			return
		}
		hSTSEnforcerDBNewFunction, err = hSTSEnforcerDBObject.InvokerNew("new")
	})
	return err
}

// HSTSEnforcerDBNew is a representation of the C type soup_hsts_enforcer_db_new.
func HSTSEnforcerDBNew(filename string) *HSTSEnforcerDB {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(filename)

	var ret gi.Argument

	err := hSTSEnforcerDBNewFunction_Set()
	if err == nil {
		ret = hSTSEnforcerDBNewFunction.Invoke(inArgs[:], nil)
	}

	retGo := HSTSEnforcerDBNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

// SessionFeature returns the SessionFeature interface implemented by HSTSEnforcerDB
func (recv *HSTSEnforcerDB) SessionFeature() *SessionFeature {
	return SessionFeatureNewFromNative(recv.Native())
}

var loggerObject *gi.Object
var loggerObject_Once sync.Once

func loggerObject_Set() error {
	var err error
	loggerObject_Once.Do(func() {
		loggerObject, err = gi.ObjectNew("Soup", "Logger")
	})
	return err
}

type Logger struct {
	native unsafe.Pointer
}

func LoggerNewFromNative(native unsafe.Pointer) *Logger {
	err := loggerObject_Set()
	if err != nil {
		return nil
	}

	instance := &Logger{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *Logger) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToLogger down casts any arbitrary Object to Logger.
Exercise care, as this is a potentially dangerous function
if the Object is not a Logger.
*/
func CastToLogger(object *gobject.Object) *Logger {
	return LoggerNewFromNative(object.Native())
}

// Equals compares this Logger with another Logger, and returns true if they represent the same Object.
func (recv *Logger) Equals(other *Logger) bool {
	return other.Native() == recv.Native()
}

func (recv *Logger) Native() unsafe.Pointer {
	return recv.native
}

// FieldParent returns the C field 'parent'.
func (recv *Logger) FieldParent() *gobject.Object {
	argValue := gi.ObjectFieldGet(loggerObject, recv.Native(), "parent")
	value := gobject.ObjectNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParent sets the value of the C field 'parent'.
func (recv *Logger) SetFieldParent(value *gobject.Object) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(loggerObject, recv.Native(), "parent", argValue)
}

var loggerNewFunction *gi.Function
var loggerNewFunction_Once sync.Once

func loggerNewFunction_Set() error {
	var err error
	loggerNewFunction_Once.Do(func() {
		err = loggerObject_Set()
		if err != nil {
			return
		}
		loggerNewFunction, err = loggerObject.InvokerNew("new")
	})
	return err
}

// LoggerNew is a representation of the C type soup_logger_new.
func LoggerNew(level LoggerLogLevel, maxBodySize int32) *Logger {
	var inArgs [2]gi.Argument
	inArgs[0].SetInt32(int32(level))
	inArgs[1].SetInt32(maxBodySize)

	var ret gi.Argument

	err := loggerNewFunction_Set()
	if err == nil {
		ret = loggerNewFunction.Invoke(inArgs[:], nil)
	}

	retGo := LoggerNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

var loggerAttachFunction *gi.Function
var loggerAttachFunction_Once sync.Once

func loggerAttachFunction_Set() error {
	var err error
	loggerAttachFunction_Once.Do(func() {
		err = loggerObject_Set()
		if err != nil {
			return
		}
		loggerAttachFunction, err = loggerObject.InvokerNew("attach")
	})
	return err
}

// Attach is a representation of the C type soup_logger_attach.
func (recv *Logger) Attach(session *Session) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(session.Native())

	err := loggerAttachFunction_Set()
	if err == nil {
		loggerAttachFunction.Invoke(inArgs[:], nil)
	}

	return
}

var loggerDetachFunction *gi.Function
var loggerDetachFunction_Once sync.Once

func loggerDetachFunction_Set() error {
	var err error
	loggerDetachFunction_Once.Do(func() {
		err = loggerObject_Set()
		if err != nil {
			return
		}
		loggerDetachFunction, err = loggerObject.InvokerNew("detach")
	})
	return err
}

// Detach is a representation of the C type soup_logger_detach.
func (recv *Logger) Detach(session *Session) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(session.Native())

	err := loggerDetachFunction_Set()
	if err == nil {
		loggerDetachFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'soup_logger_set_printer' : parameter 'printer' of type 'LoggerPrinter' not supported

// UNSUPPORTED : C value 'soup_logger_set_request_filter' : parameter 'request_filter' of type 'LoggerFilter' not supported

// UNSUPPORTED : C value 'soup_logger_set_response_filter' : parameter 'response_filter' of type 'LoggerFilter' not supported

// SessionFeature returns the SessionFeature interface implemented by Logger
func (recv *Logger) SessionFeature() *SessionFeature {
	return SessionFeatureNewFromNative(recv.Native())
}

var messageObject *gi.Object
var messageObject_Once sync.Once

func messageObject_Set() error {
	var err error
	messageObject_Once.Do(func() {
		messageObject, err = gi.ObjectNew("Soup", "Message")
	})
	return err
}

type Message struct {
	native unsafe.Pointer
}

func MessageNewFromNative(native unsafe.Pointer) *Message {
	err := messageObject_Set()
	if err != nil {
		return nil
	}

	instance := &Message{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *Message) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToMessage down casts any arbitrary Object to Message.
Exercise care, as this is a potentially dangerous function
if the Object is not a Message.
*/
func CastToMessage(object *gobject.Object) *Message {
	return MessageNewFromNative(object.Native())
}

// Equals compares this Message with another Message, and returns true if they represent the same Object.
func (recv *Message) Equals(other *Message) bool {
	return other.Native() == recv.Native()
}

func (recv *Message) Native() unsafe.Pointer {
	return recv.native
}

// FieldParent returns the C field 'parent'.
func (recv *Message) FieldParent() *gobject.Object {
	argValue := gi.ObjectFieldGet(messageObject, recv.Native(), "parent")
	value := gobject.ObjectNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParent sets the value of the C field 'parent'.
func (recv *Message) SetFieldParent(value *gobject.Object) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(messageObject, recv.Native(), "parent", argValue)
}

// FieldMethod returns the C field 'method'.
func (recv *Message) FieldMethod() string {
	argValue := gi.ObjectFieldGet(messageObject, recv.Native(), "method")
	value := argValue.String(false)
	return value
}

// SetFieldMethod sets the value of the C field 'method'.
func (recv *Message) SetFieldMethod(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.ObjectFieldSet(messageObject, recv.Native(), "method", argValue)
}

// FieldStatusCode returns the C field 'status_code'.
func (recv *Message) FieldStatusCode() uint32 {
	argValue := gi.ObjectFieldGet(messageObject, recv.Native(), "status_code")
	value := argValue.Uint32()
	return value
}

// SetFieldStatusCode sets the value of the C field 'status_code'.
func (recv *Message) SetFieldStatusCode(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.ObjectFieldSet(messageObject, recv.Native(), "status_code", argValue)
}

// FieldReasonPhrase returns the C field 'reason_phrase'.
func (recv *Message) FieldReasonPhrase() string {
	argValue := gi.ObjectFieldGet(messageObject, recv.Native(), "reason_phrase")
	value := argValue.String(false)
	return value
}

// SetFieldReasonPhrase sets the value of the C field 'reason_phrase'.
func (recv *Message) SetFieldReasonPhrase(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.ObjectFieldSet(messageObject, recv.Native(), "reason_phrase", argValue)
}

// FieldRequestBody returns the C field 'request_body'.
func (recv *Message) FieldRequestBody() *MessageBody {
	argValue := gi.ObjectFieldGet(messageObject, recv.Native(), "request_body")
	value := MessageBodyNewFromNative(argValue.Pointer())
	return value
}

// SetFieldRequestBody sets the value of the C field 'request_body'.
func (recv *Message) SetFieldRequestBody(value *MessageBody) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(messageObject, recv.Native(), "request_body", argValue)
}

// FieldRequestHeaders returns the C field 'request_headers'.
func (recv *Message) FieldRequestHeaders() *MessageHeaders {
	argValue := gi.ObjectFieldGet(messageObject, recv.Native(), "request_headers")
	value := MessageHeadersNewFromNative(argValue.Pointer())
	return value
}

// SetFieldRequestHeaders sets the value of the C field 'request_headers'.
func (recv *Message) SetFieldRequestHeaders(value *MessageHeaders) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(messageObject, recv.Native(), "request_headers", argValue)
}

// FieldResponseBody returns the C field 'response_body'.
func (recv *Message) FieldResponseBody() *MessageBody {
	argValue := gi.ObjectFieldGet(messageObject, recv.Native(), "response_body")
	value := MessageBodyNewFromNative(argValue.Pointer())
	return value
}

// SetFieldResponseBody sets the value of the C field 'response_body'.
func (recv *Message) SetFieldResponseBody(value *MessageBody) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(messageObject, recv.Native(), "response_body", argValue)
}

// FieldResponseHeaders returns the C field 'response_headers'.
func (recv *Message) FieldResponseHeaders() *MessageHeaders {
	argValue := gi.ObjectFieldGet(messageObject, recv.Native(), "response_headers")
	value := MessageHeadersNewFromNative(argValue.Pointer())
	return value
}

// SetFieldResponseHeaders sets the value of the C field 'response_headers'.
func (recv *Message) SetFieldResponseHeaders(value *MessageHeaders) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(messageObject, recv.Native(), "response_headers", argValue)
}

var messageNewFunction *gi.Function
var messageNewFunction_Once sync.Once

func messageNewFunction_Set() error {
	var err error
	messageNewFunction_Once.Do(func() {
		err = messageObject_Set()
		if err != nil {
			return
		}
		messageNewFunction, err = messageObject.InvokerNew("new")
	})
	return err
}

// MessageNew is a representation of the C type soup_message_new.
func MessageNew(method string, uriString string) *Message {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(method)
	inArgs[1].SetString(uriString)

	var ret gi.Argument

	err := messageNewFunction_Set()
	if err == nil {
		ret = messageNewFunction.Invoke(inArgs[:], nil)
	}

	retGo := MessageNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

var messageNewFromUriFunction *gi.Function
var messageNewFromUriFunction_Once sync.Once

func messageNewFromUriFunction_Set() error {
	var err error
	messageNewFromUriFunction_Once.Do(func() {
		err = messageObject_Set()
		if err != nil {
			return
		}
		messageNewFromUriFunction, err = messageObject.InvokerNew("new_from_uri")
	})
	return err
}

// MessageNewFromUri is a representation of the C type soup_message_new_from_uri.
func MessageNewFromUri(method string, uri *URI) *Message {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(method)
	inArgs[1].SetPointer(uri.Native())

	var ret gi.Argument

	err := messageNewFromUriFunction_Set()
	if err == nil {
		ret = messageNewFromUriFunction.Invoke(inArgs[:], nil)
	}

	retGo := MessageNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

// UNSUPPORTED : C value 'soup_message_add_header_handler' : parameter 'callback' of type 'GObject.Callback' not supported

// UNSUPPORTED : C value 'soup_message_add_status_code_handler' : parameter 'callback' of type 'GObject.Callback' not supported

var messageContentSniffedFunction *gi.Function
var messageContentSniffedFunction_Once sync.Once

func messageContentSniffedFunction_Set() error {
	var err error
	messageContentSniffedFunction_Once.Do(func() {
		err = messageObject_Set()
		if err != nil {
			return
		}
		messageContentSniffedFunction, err = messageObject.InvokerNew("content_sniffed")
	})
	return err
}

// ContentSniffed is a representation of the C type soup_message_content_sniffed.
func (recv *Message) ContentSniffed(contentType string, params *glib.HashTable) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(contentType)
	inArgs[2].SetPointer(params.Native())

	err := messageContentSniffedFunction_Set()
	if err == nil {
		messageContentSniffedFunction.Invoke(inArgs[:], nil)
	}

	return
}

var messageDisableFeatureFunction *gi.Function
var messageDisableFeatureFunction_Once sync.Once

func messageDisableFeatureFunction_Set() error {
	var err error
	messageDisableFeatureFunction_Once.Do(func() {
		err = messageObject_Set()
		if err != nil {
			return
		}
		messageDisableFeatureFunction, err = messageObject.InvokerNew("disable_feature")
	})
	return err
}

// DisableFeature is a representation of the C type soup_message_disable_feature.
func (recv *Message) DisableFeature(featureType int64) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt64(featureType)

	err := messageDisableFeatureFunction_Set()
	if err == nil {
		messageDisableFeatureFunction.Invoke(inArgs[:], nil)
	}

	return
}

var messageFinishedFunction *gi.Function
var messageFinishedFunction_Once sync.Once

func messageFinishedFunction_Set() error {
	var err error
	messageFinishedFunction_Once.Do(func() {
		err = messageObject_Set()
		if err != nil {
			return
		}
		messageFinishedFunction, err = messageObject.InvokerNew("finished")
	})
	return err
}

// Finished is a representation of the C type soup_message_finished.
func (recv *Message) Finished() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := messageFinishedFunction_Set()
	if err == nil {
		messageFinishedFunction.Invoke(inArgs[:], nil)
	}

	return
}

var messageGetAddressFunction *gi.Function
var messageGetAddressFunction_Once sync.Once

func messageGetAddressFunction_Set() error {
	var err error
	messageGetAddressFunction_Once.Do(func() {
		err = messageObject_Set()
		if err != nil {
			return
		}
		messageGetAddressFunction, err = messageObject.InvokerNew("get_address")
	})
	return err
}

// GetAddress is a representation of the C type soup_message_get_address.
func (recv *Message) GetAddress() *Address {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := messageGetAddressFunction_Set()
	if err == nil {
		ret = messageGetAddressFunction.Invoke(inArgs[:], nil)
	}

	retGo := AddressNewFromNative(ret.Pointer())

	return retGo
}

var messageGetFirstPartyFunction *gi.Function
var messageGetFirstPartyFunction_Once sync.Once

func messageGetFirstPartyFunction_Set() error {
	var err error
	messageGetFirstPartyFunction_Once.Do(func() {
		err = messageObject_Set()
		if err != nil {
			return
		}
		messageGetFirstPartyFunction, err = messageObject.InvokerNew("get_first_party")
	})
	return err
}

// GetFirstParty is a representation of the C type soup_message_get_first_party.
func (recv *Message) GetFirstParty() *URI {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := messageGetFirstPartyFunction_Set()
	if err == nil {
		ret = messageGetFirstPartyFunction.Invoke(inArgs[:], nil)
	}

	retGo := URINewFromNative(ret.Pointer())

	return retGo
}

var messageGetFlagsFunction *gi.Function
var messageGetFlagsFunction_Once sync.Once

func messageGetFlagsFunction_Set() error {
	var err error
	messageGetFlagsFunction_Once.Do(func() {
		err = messageObject_Set()
		if err != nil {
			return
		}
		messageGetFlagsFunction, err = messageObject.InvokerNew("get_flags")
	})
	return err
}

// GetFlags is a representation of the C type soup_message_get_flags.
func (recv *Message) GetFlags() MessageFlags {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := messageGetFlagsFunction_Set()
	if err == nil {
		ret = messageGetFlagsFunction.Invoke(inArgs[:], nil)
	}

	retGo := MessageFlags(ret.Int32())

	return retGo
}

var messageGetHttpVersionFunction *gi.Function
var messageGetHttpVersionFunction_Once sync.Once

func messageGetHttpVersionFunction_Set() error {
	var err error
	messageGetHttpVersionFunction_Once.Do(func() {
		err = messageObject_Set()
		if err != nil {
			return
		}
		messageGetHttpVersionFunction, err = messageObject.InvokerNew("get_http_version")
	})
	return err
}

// GetHttpVersion is a representation of the C type soup_message_get_http_version.
func (recv *Message) GetHttpVersion() HTTPVersion {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := messageGetHttpVersionFunction_Set()
	if err == nil {
		ret = messageGetHttpVersionFunction.Invoke(inArgs[:], nil)
	}

	retGo := HTTPVersion(ret.Int32())

	return retGo
}

// UNSUPPORTED : C value 'soup_message_get_https_status' : parameter 'errors' of type 'Gio.TlsCertificateFlags' not supported

var messageGetPriorityFunction *gi.Function
var messageGetPriorityFunction_Once sync.Once

func messageGetPriorityFunction_Set() error {
	var err error
	messageGetPriorityFunction_Once.Do(func() {
		err = messageObject_Set()
		if err != nil {
			return
		}
		messageGetPriorityFunction, err = messageObject.InvokerNew("get_priority")
	})
	return err
}

// GetPriority is a representation of the C type soup_message_get_priority.
func (recv *Message) GetPriority() MessagePriority {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := messageGetPriorityFunction_Set()
	if err == nil {
		ret = messageGetPriorityFunction.Invoke(inArgs[:], nil)
	}

	retGo := MessagePriority(ret.Int32())

	return retGo
}

var messageGetSoupRequestFunction *gi.Function
var messageGetSoupRequestFunction_Once sync.Once

func messageGetSoupRequestFunction_Set() error {
	var err error
	messageGetSoupRequestFunction_Once.Do(func() {
		err = messageObject_Set()
		if err != nil {
			return
		}
		messageGetSoupRequestFunction, err = messageObject.InvokerNew("get_soup_request")
	})
	return err
}

// GetSoupRequest is a representation of the C type soup_message_get_soup_request.
func (recv *Message) GetSoupRequest() *Request {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := messageGetSoupRequestFunction_Set()
	if err == nil {
		ret = messageGetSoupRequestFunction.Invoke(inArgs[:], nil)
	}

	retGo := RequestNewFromNative(ret.Pointer())

	return retGo
}

var messageGetUriFunction *gi.Function
var messageGetUriFunction_Once sync.Once

func messageGetUriFunction_Set() error {
	var err error
	messageGetUriFunction_Once.Do(func() {
		err = messageObject_Set()
		if err != nil {
			return
		}
		messageGetUriFunction, err = messageObject.InvokerNew("get_uri")
	})
	return err
}

// GetUri is a representation of the C type soup_message_get_uri.
func (recv *Message) GetUri() *URI {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := messageGetUriFunction_Set()
	if err == nil {
		ret = messageGetUriFunction.Invoke(inArgs[:], nil)
	}

	retGo := URINewFromNative(ret.Pointer())

	return retGo
}

var messageGotBodyFunction *gi.Function
var messageGotBodyFunction_Once sync.Once

func messageGotBodyFunction_Set() error {
	var err error
	messageGotBodyFunction_Once.Do(func() {
		err = messageObject_Set()
		if err != nil {
			return
		}
		messageGotBodyFunction, err = messageObject.InvokerNew("got_body")
	})
	return err
}

// GotBody is a representation of the C type soup_message_got_body.
func (recv *Message) GotBody() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := messageGotBodyFunction_Set()
	if err == nil {
		messageGotBodyFunction.Invoke(inArgs[:], nil)
	}

	return
}

var messageGotChunkFunction *gi.Function
var messageGotChunkFunction_Once sync.Once

func messageGotChunkFunction_Set() error {
	var err error
	messageGotChunkFunction_Once.Do(func() {
		err = messageObject_Set()
		if err != nil {
			return
		}
		messageGotChunkFunction, err = messageObject.InvokerNew("got_chunk")
	})
	return err
}

// GotChunk is a representation of the C type soup_message_got_chunk.
func (recv *Message) GotChunk(chunk *Buffer) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(chunk.Native())

	err := messageGotChunkFunction_Set()
	if err == nil {
		messageGotChunkFunction.Invoke(inArgs[:], nil)
	}

	return
}

var messageGotHeadersFunction *gi.Function
var messageGotHeadersFunction_Once sync.Once

func messageGotHeadersFunction_Set() error {
	var err error
	messageGotHeadersFunction_Once.Do(func() {
		err = messageObject_Set()
		if err != nil {
			return
		}
		messageGotHeadersFunction, err = messageObject.InvokerNew("got_headers")
	})
	return err
}

// GotHeaders is a representation of the C type soup_message_got_headers.
func (recv *Message) GotHeaders() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := messageGotHeadersFunction_Set()
	if err == nil {
		messageGotHeadersFunction.Invoke(inArgs[:], nil)
	}

	return
}

var messageGotInformationalFunction *gi.Function
var messageGotInformationalFunction_Once sync.Once

func messageGotInformationalFunction_Set() error {
	var err error
	messageGotInformationalFunction_Once.Do(func() {
		err = messageObject_Set()
		if err != nil {
			return
		}
		messageGotInformationalFunction, err = messageObject.InvokerNew("got_informational")
	})
	return err
}

// GotInformational is a representation of the C type soup_message_got_informational.
func (recv *Message) GotInformational() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := messageGotInformationalFunction_Set()
	if err == nil {
		messageGotInformationalFunction.Invoke(inArgs[:], nil)
	}

	return
}

var messageIsKeepaliveFunction *gi.Function
var messageIsKeepaliveFunction_Once sync.Once

func messageIsKeepaliveFunction_Set() error {
	var err error
	messageIsKeepaliveFunction_Once.Do(func() {
		err = messageObject_Set()
		if err != nil {
			return
		}
		messageIsKeepaliveFunction, err = messageObject.InvokerNew("is_keepalive")
	})
	return err
}

// IsKeepalive is a representation of the C type soup_message_is_keepalive.
func (recv *Message) IsKeepalive() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := messageIsKeepaliveFunction_Set()
	if err == nil {
		ret = messageIsKeepaliveFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var messageRestartedFunction *gi.Function
var messageRestartedFunction_Once sync.Once

func messageRestartedFunction_Set() error {
	var err error
	messageRestartedFunction_Once.Do(func() {
		err = messageObject_Set()
		if err != nil {
			return
		}
		messageRestartedFunction, err = messageObject.InvokerNew("restarted")
	})
	return err
}

// Restarted is a representation of the C type soup_message_restarted.
func (recv *Message) Restarted() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := messageRestartedFunction_Set()
	if err == nil {
		messageRestartedFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'soup_message_set_chunk_allocator' : parameter 'allocator' of type 'ChunkAllocator' not supported

var messageSetFirstPartyFunction *gi.Function
var messageSetFirstPartyFunction_Once sync.Once

func messageSetFirstPartyFunction_Set() error {
	var err error
	messageSetFirstPartyFunction_Once.Do(func() {
		err = messageObject_Set()
		if err != nil {
			return
		}
		messageSetFirstPartyFunction, err = messageObject.InvokerNew("set_first_party")
	})
	return err
}

// SetFirstParty is a representation of the C type soup_message_set_first_party.
func (recv *Message) SetFirstParty(firstParty *URI) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(firstParty.Native())

	err := messageSetFirstPartyFunction_Set()
	if err == nil {
		messageSetFirstPartyFunction.Invoke(inArgs[:], nil)
	}

	return
}

var messageSetFlagsFunction *gi.Function
var messageSetFlagsFunction_Once sync.Once

func messageSetFlagsFunction_Set() error {
	var err error
	messageSetFlagsFunction_Once.Do(func() {
		err = messageObject_Set()
		if err != nil {
			return
		}
		messageSetFlagsFunction, err = messageObject.InvokerNew("set_flags")
	})
	return err
}

// SetFlags is a representation of the C type soup_message_set_flags.
func (recv *Message) SetFlags(flags MessageFlags) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(int32(flags))

	err := messageSetFlagsFunction_Set()
	if err == nil {
		messageSetFlagsFunction.Invoke(inArgs[:], nil)
	}

	return
}

var messageSetHttpVersionFunction *gi.Function
var messageSetHttpVersionFunction_Once sync.Once

func messageSetHttpVersionFunction_Set() error {
	var err error
	messageSetHttpVersionFunction_Once.Do(func() {
		err = messageObject_Set()
		if err != nil {
			return
		}
		messageSetHttpVersionFunction, err = messageObject.InvokerNew("set_http_version")
	})
	return err
}

// SetHttpVersion is a representation of the C type soup_message_set_http_version.
func (recv *Message) SetHttpVersion(version HTTPVersion) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(int32(version))

	err := messageSetHttpVersionFunction_Set()
	if err == nil {
		messageSetHttpVersionFunction.Invoke(inArgs[:], nil)
	}

	return
}

var messageSetPriorityFunction *gi.Function
var messageSetPriorityFunction_Once sync.Once

func messageSetPriorityFunction_Set() error {
	var err error
	messageSetPriorityFunction_Once.Do(func() {
		err = messageObject_Set()
		if err != nil {
			return
		}
		messageSetPriorityFunction, err = messageObject.InvokerNew("set_priority")
	})
	return err
}

// SetPriority is a representation of the C type soup_message_set_priority.
func (recv *Message) SetPriority(priority MessagePriority) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(int32(priority))

	err := messageSetPriorityFunction_Set()
	if err == nil {
		messageSetPriorityFunction.Invoke(inArgs[:], nil)
	}

	return
}

var messageSetRedirectFunction *gi.Function
var messageSetRedirectFunction_Once sync.Once

func messageSetRedirectFunction_Set() error {
	var err error
	messageSetRedirectFunction_Once.Do(func() {
		err = messageObject_Set()
		if err != nil {
			return
		}
		messageSetRedirectFunction, err = messageObject.InvokerNew("set_redirect")
	})
	return err
}

// SetRedirect is a representation of the C type soup_message_set_redirect.
func (recv *Message) SetRedirect(statusCode uint32, redirectUri string) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetUint32(statusCode)
	inArgs[2].SetString(redirectUri)

	err := messageSetRedirectFunction_Set()
	if err == nil {
		messageSetRedirectFunction.Invoke(inArgs[:], nil)
	}

	return
}

var messageSetRequestFunction *gi.Function
var messageSetRequestFunction_Once sync.Once

func messageSetRequestFunction_Set() error {
	var err error
	messageSetRequestFunction_Once.Do(func() {
		err = messageObject_Set()
		if err != nil {
			return
		}
		messageSetRequestFunction, err = messageObject.InvokerNew("set_request")
	})
	return err
}

// SetRequest is a representation of the C type soup_message_set_request.
func (recv *Message) SetRequest(contentType string, reqUse MemoryUse, reqBody string) {
	var inArgs [5]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(contentType)
	inArgs[2].SetInt32(int32(reqUse))
	inArgs[3].SetString(reqBody)
	inArgs[4].SetUint64(uint64(len(reqBody)))

	err := messageSetRequestFunction_Set()
	if err == nil {
		messageSetRequestFunction.Invoke(inArgs[:], nil)
	}

	return
}

var messageSetResponseFunction *gi.Function
var messageSetResponseFunction_Once sync.Once

func messageSetResponseFunction_Set() error {
	var err error
	messageSetResponseFunction_Once.Do(func() {
		err = messageObject_Set()
		if err != nil {
			return
		}
		messageSetResponseFunction, err = messageObject.InvokerNew("set_response")
	})
	return err
}

// SetResponse is a representation of the C type soup_message_set_response.
func (recv *Message) SetResponse(contentType string, respUse MemoryUse, respBody string) {
	var inArgs [5]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(contentType)
	inArgs[2].SetInt32(int32(respUse))
	inArgs[3].SetString(respBody)
	inArgs[4].SetUint64(uint64(len(respBody)))

	err := messageSetResponseFunction_Set()
	if err == nil {
		messageSetResponseFunction.Invoke(inArgs[:], nil)
	}

	return
}

var messageSetStatusFunction *gi.Function
var messageSetStatusFunction_Once sync.Once

func messageSetStatusFunction_Set() error {
	var err error
	messageSetStatusFunction_Once.Do(func() {
		err = messageObject_Set()
		if err != nil {
			return
		}
		messageSetStatusFunction, err = messageObject.InvokerNew("set_status")
	})
	return err
}

// SetStatus is a representation of the C type soup_message_set_status.
func (recv *Message) SetStatus(statusCode uint32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetUint32(statusCode)

	err := messageSetStatusFunction_Set()
	if err == nil {
		messageSetStatusFunction.Invoke(inArgs[:], nil)
	}

	return
}

var messageSetStatusFullFunction *gi.Function
var messageSetStatusFullFunction_Once sync.Once

func messageSetStatusFullFunction_Set() error {
	var err error
	messageSetStatusFullFunction_Once.Do(func() {
		err = messageObject_Set()
		if err != nil {
			return
		}
		messageSetStatusFullFunction, err = messageObject.InvokerNew("set_status_full")
	})
	return err
}

// SetStatusFull is a representation of the C type soup_message_set_status_full.
func (recv *Message) SetStatusFull(statusCode uint32, reasonPhrase string) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetUint32(statusCode)
	inArgs[2].SetString(reasonPhrase)

	err := messageSetStatusFullFunction_Set()
	if err == nil {
		messageSetStatusFullFunction.Invoke(inArgs[:], nil)
	}

	return
}

var messageSetUriFunction *gi.Function
var messageSetUriFunction_Once sync.Once

func messageSetUriFunction_Set() error {
	var err error
	messageSetUriFunction_Once.Do(func() {
		err = messageObject_Set()
		if err != nil {
			return
		}
		messageSetUriFunction, err = messageObject.InvokerNew("set_uri")
	})
	return err
}

// SetUri is a representation of the C type soup_message_set_uri.
func (recv *Message) SetUri(uri *URI) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(uri.Native())

	err := messageSetUriFunction_Set()
	if err == nil {
		messageSetUriFunction.Invoke(inArgs[:], nil)
	}

	return
}

var messageStartingFunction *gi.Function
var messageStartingFunction_Once sync.Once

func messageStartingFunction_Set() error {
	var err error
	messageStartingFunction_Once.Do(func() {
		err = messageObject_Set()
		if err != nil {
			return
		}
		messageStartingFunction, err = messageObject.InvokerNew("starting")
	})
	return err
}

// Starting is a representation of the C type soup_message_starting.
func (recv *Message) Starting() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := messageStartingFunction_Set()
	if err == nil {
		messageStartingFunction.Invoke(inArgs[:], nil)
	}

	return
}

var messageWroteBodyFunction *gi.Function
var messageWroteBodyFunction_Once sync.Once

func messageWroteBodyFunction_Set() error {
	var err error
	messageWroteBodyFunction_Once.Do(func() {
		err = messageObject_Set()
		if err != nil {
			return
		}
		messageWroteBodyFunction, err = messageObject.InvokerNew("wrote_body")
	})
	return err
}

// WroteBody is a representation of the C type soup_message_wrote_body.
func (recv *Message) WroteBody() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := messageWroteBodyFunction_Set()
	if err == nil {
		messageWroteBodyFunction.Invoke(inArgs[:], nil)
	}

	return
}

var messageWroteBodyDataFunction *gi.Function
var messageWroteBodyDataFunction_Once sync.Once

func messageWroteBodyDataFunction_Set() error {
	var err error
	messageWroteBodyDataFunction_Once.Do(func() {
		err = messageObject_Set()
		if err != nil {
			return
		}
		messageWroteBodyDataFunction, err = messageObject.InvokerNew("wrote_body_data")
	})
	return err
}

// WroteBodyData is a representation of the C type soup_message_wrote_body_data.
func (recv *Message) WroteBodyData(chunk *Buffer) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(chunk.Native())

	err := messageWroteBodyDataFunction_Set()
	if err == nil {
		messageWroteBodyDataFunction.Invoke(inArgs[:], nil)
	}

	return
}

var messageWroteChunkFunction *gi.Function
var messageWroteChunkFunction_Once sync.Once

func messageWroteChunkFunction_Set() error {
	var err error
	messageWroteChunkFunction_Once.Do(func() {
		err = messageObject_Set()
		if err != nil {
			return
		}
		messageWroteChunkFunction, err = messageObject.InvokerNew("wrote_chunk")
	})
	return err
}

// WroteChunk is a representation of the C type soup_message_wrote_chunk.
func (recv *Message) WroteChunk() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := messageWroteChunkFunction_Set()
	if err == nil {
		messageWroteChunkFunction.Invoke(inArgs[:], nil)
	}

	return
}

var messageWroteHeadersFunction *gi.Function
var messageWroteHeadersFunction_Once sync.Once

func messageWroteHeadersFunction_Set() error {
	var err error
	messageWroteHeadersFunction_Once.Do(func() {
		err = messageObject_Set()
		if err != nil {
			return
		}
		messageWroteHeadersFunction, err = messageObject.InvokerNew("wrote_headers")
	})
	return err
}

// WroteHeaders is a representation of the C type soup_message_wrote_headers.
func (recv *Message) WroteHeaders() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := messageWroteHeadersFunction_Set()
	if err == nil {
		messageWroteHeadersFunction.Invoke(inArgs[:], nil)
	}

	return
}

var messageWroteInformationalFunction *gi.Function
var messageWroteInformationalFunction_Once sync.Once

func messageWroteInformationalFunction_Set() error {
	var err error
	messageWroteInformationalFunction_Once.Do(func() {
		err = messageObject_Set()
		if err != nil {
			return
		}
		messageWroteInformationalFunction, err = messageObject.InvokerNew("wrote_informational")
	})
	return err
}

// WroteInformational is a representation of the C type soup_message_wrote_informational.
func (recv *Message) WroteInformational() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := messageWroteInformationalFunction_Set()
	if err == nil {
		messageWroteInformationalFunction.Invoke(inArgs[:], nil)
	}

	return
}

/*
ConnectContentSniffed connects a callback to the 'content-sniffed' signal of the Message.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *Message) ConnectContentSniffed(handler func(instance *Message, type_ string, params *glib.HashTable)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {
		objectInstance := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[0]))
		argInstance := MessageNewFromNative(objectInstance.GetObject().Native())

		object1 := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[1]))
		arg1 := object1.GetString()

		object2 := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[2]))
		arg2 := glib.HashTableNewFromNative(object2.GetBoxed())

		handler(argInstance, arg1, arg2)

	}

	return callback.ConnectSignal(recv.Native(), "content-sniffed", marshal)
}

/*
ConnectFinished connects a callback to the 'finished' signal of the Message.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *Message) ConnectFinished(handler func(instance *Message)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {
		objectInstance := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[0]))
		argInstance := MessageNewFromNative(objectInstance.GetObject().Native())

		handler(argInstance)

	}

	return callback.ConnectSignal(recv.Native(), "finished", marshal)
}

/*
ConnectGotBody connects a callback to the 'got-body' signal of the Message.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *Message) ConnectGotBody(handler func(instance *Message)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {
		objectInstance := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[0]))
		argInstance := MessageNewFromNative(objectInstance.GetObject().Native())

		handler(argInstance)

	}

	return callback.ConnectSignal(recv.Native(), "got-body", marshal)
}

/*
ConnectGotChunk connects a callback to the 'got-chunk' signal of the Message.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *Message) ConnectGotChunk(handler func(instance *Message, chunk *Buffer)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {
		objectInstance := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[0]))
		argInstance := MessageNewFromNative(objectInstance.GetObject().Native())

		object1 := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[1]))
		arg1 := BufferNewFromNative(object1.GetBoxed())

		handler(argInstance, arg1)

	}

	return callback.ConnectSignal(recv.Native(), "got-chunk", marshal)
}

/*
ConnectGotHeaders connects a callback to the 'got-headers' signal of the Message.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *Message) ConnectGotHeaders(handler func(instance *Message)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {
		objectInstance := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[0]))
		argInstance := MessageNewFromNative(objectInstance.GetObject().Native())

		handler(argInstance)

	}

	return callback.ConnectSignal(recv.Native(), "got-headers", marshal)
}

/*
ConnectGotInformational connects a callback to the 'got-informational' signal of the Message.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *Message) ConnectGotInformational(handler func(instance *Message)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {
		objectInstance := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[0]))
		argInstance := MessageNewFromNative(objectInstance.GetObject().Native())

		handler(argInstance)

	}

	return callback.ConnectSignal(recv.Native(), "got-informational", marshal)
}

// UNSUPPORTED : C value 'network-event' : parameter 'event' of type 'Gio.SocketClientEvent' not supported

/*
ConnectRestarted connects a callback to the 'restarted' signal of the Message.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *Message) ConnectRestarted(handler func(instance *Message)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {
		objectInstance := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[0]))
		argInstance := MessageNewFromNative(objectInstance.GetObject().Native())

		handler(argInstance)

	}

	return callback.ConnectSignal(recv.Native(), "restarted", marshal)
}

/*
ConnectStarting connects a callback to the 'starting' signal of the Message.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *Message) ConnectStarting(handler func(instance *Message)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {
		objectInstance := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[0]))
		argInstance := MessageNewFromNative(objectInstance.GetObject().Native())

		handler(argInstance)

	}

	return callback.ConnectSignal(recv.Native(), "starting", marshal)
}

/*
ConnectWroteBody connects a callback to the 'wrote-body' signal of the Message.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *Message) ConnectWroteBody(handler func(instance *Message)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {
		objectInstance := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[0]))
		argInstance := MessageNewFromNative(objectInstance.GetObject().Native())

		handler(argInstance)

	}

	return callback.ConnectSignal(recv.Native(), "wrote-body", marshal)
}

/*
ConnectWroteBodyData connects a callback to the 'wrote-body-data' signal of the Message.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *Message) ConnectWroteBodyData(handler func(instance *Message, chunk *Buffer)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {
		objectInstance := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[0]))
		argInstance := MessageNewFromNative(objectInstance.GetObject().Native())

		object1 := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[1]))
		arg1 := BufferNewFromNative(object1.GetBoxed())

		handler(argInstance, arg1)

	}

	return callback.ConnectSignal(recv.Native(), "wrote-body-data", marshal)
}

/*
ConnectWroteChunk connects a callback to the 'wrote-chunk' signal of the Message.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *Message) ConnectWroteChunk(handler func(instance *Message)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {
		objectInstance := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[0]))
		argInstance := MessageNewFromNative(objectInstance.GetObject().Native())

		handler(argInstance)

	}

	return callback.ConnectSignal(recv.Native(), "wrote-chunk", marshal)
}

/*
ConnectWroteHeaders connects a callback to the 'wrote-headers' signal of the Message.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *Message) ConnectWroteHeaders(handler func(instance *Message)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {
		objectInstance := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[0]))
		argInstance := MessageNewFromNative(objectInstance.GetObject().Native())

		handler(argInstance)

	}

	return callback.ConnectSignal(recv.Native(), "wrote-headers", marshal)
}

/*
ConnectWroteInformational connects a callback to the 'wrote-informational' signal of the Message.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *Message) ConnectWroteInformational(handler func(instance *Message)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {
		objectInstance := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[0]))
		argInstance := MessageNewFromNative(objectInstance.GetObject().Native())

		handler(argInstance)

	}

	return callback.ConnectSignal(recv.Native(), "wrote-informational", marshal)
}

/*
Disconnect disconnects a callback previously registered with a Connect...() method.

The connectionID should be a value returned from a call to a Connect...() method.
*/
func (recv *Message) DisconnectSignal(connectionID int) {
	callback.DisconnectSignal(connectionID)
}

var multipartInputStreamObject *gi.Object
var multipartInputStreamObject_Once sync.Once

func multipartInputStreamObject_Set() error {
	var err error
	multipartInputStreamObject_Once.Do(func() {
		multipartInputStreamObject, err = gi.ObjectNew("Soup", "MultipartInputStream")
	})
	return err
}

type MultipartInputStream struct {
	native unsafe.Pointer
}

func MultipartInputStreamNewFromNative(native unsafe.Pointer) *MultipartInputStream {
	err := multipartInputStreamObject_Set()
	if err != nil {
		return nil
	}

	instance := &MultipartInputStream{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// FilterInputStream upcasts to *FilterInputStream
func (recv *MultipartInputStream) FilterInputStream() *gio.FilterInputStream {
	return gio.FilterInputStreamNewFromNative(recv.Native())
}

// InputStream upcasts to *InputStream
func (recv *MultipartInputStream) InputStream() *gio.InputStream {
	return gio.InputStreamNewFromNative(recv.Native())
}

// Object upcasts to *Object
func (recv *MultipartInputStream) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToMultipartInputStream down casts any arbitrary Object to MultipartInputStream.
Exercise care, as this is a potentially dangerous function
if the Object is not a MultipartInputStream.
*/
func CastToMultipartInputStream(object *gobject.Object) *MultipartInputStream {
	return MultipartInputStreamNewFromNative(object.Native())
}

// Equals compares this MultipartInputStream with another MultipartInputStream, and returns true if they represent the same Object.
func (recv *MultipartInputStream) Equals(other *MultipartInputStream) bool {
	return other.Native() == recv.Native()
}

func (recv *MultipartInputStream) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *MultipartInputStream) FieldParentInstance() *gio.FilterInputStream {
	argValue := gi.ObjectFieldGet(multipartInputStreamObject, recv.Native(), "parent_instance")
	value := gio.FilterInputStreamNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *MultipartInputStream) SetFieldParentInstance(value *gio.FilterInputStream) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(multipartInputStreamObject, recv.Native(), "parent_instance", argValue)
}

var multipartInputStreamNewFunction *gi.Function
var multipartInputStreamNewFunction_Once sync.Once

func multipartInputStreamNewFunction_Set() error {
	var err error
	multipartInputStreamNewFunction_Once.Do(func() {
		err = multipartInputStreamObject_Set()
		if err != nil {
			return
		}
		multipartInputStreamNewFunction, err = multipartInputStreamObject.InvokerNew("new")
	})
	return err
}

// MultipartInputStreamNew is a representation of the C type soup_multipart_input_stream_new.
func MultipartInputStreamNew(msg *Message, baseStream *gio.InputStream) *MultipartInputStream {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(msg.Native())
	inArgs[1].SetPointer(baseStream.Native())

	var ret gi.Argument

	err := multipartInputStreamNewFunction_Set()
	if err == nil {
		ret = multipartInputStreamNewFunction.Invoke(inArgs[:], nil)
	}

	retGo := MultipartInputStreamNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

var multipartInputStreamGetHeadersFunction *gi.Function
var multipartInputStreamGetHeadersFunction_Once sync.Once

func multipartInputStreamGetHeadersFunction_Set() error {
	var err error
	multipartInputStreamGetHeadersFunction_Once.Do(func() {
		err = multipartInputStreamObject_Set()
		if err != nil {
			return
		}
		multipartInputStreamGetHeadersFunction, err = multipartInputStreamObject.InvokerNew("get_headers")
	})
	return err
}

// GetHeaders is a representation of the C type soup_multipart_input_stream_get_headers.
func (recv *MultipartInputStream) GetHeaders() *MessageHeaders {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := multipartInputStreamGetHeadersFunction_Set()
	if err == nil {
		ret = multipartInputStreamGetHeadersFunction.Invoke(inArgs[:], nil)
	}

	retGo := MessageHeadersNewFromNative(ret.Pointer())

	return retGo
}

var multipartInputStreamNextPartFunction *gi.Function
var multipartInputStreamNextPartFunction_Once sync.Once

func multipartInputStreamNextPartFunction_Set() error {
	var err error
	multipartInputStreamNextPartFunction_Once.Do(func() {
		err = multipartInputStreamObject_Set()
		if err != nil {
			return
		}
		multipartInputStreamNextPartFunction, err = multipartInputStreamObject.InvokerNew("next_part")
	})
	return err
}

// NextPart is a representation of the C type soup_multipart_input_stream_next_part.
func (recv *MultipartInputStream) NextPart(cancellable *gio.Cancellable) *gio.InputStream {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(cancellable.Native())

	var ret gi.Argument

	err := multipartInputStreamNextPartFunction_Set()
	if err == nil {
		ret = multipartInputStreamNextPartFunction.Invoke(inArgs[:], nil)
	}

	retGo := gio.InputStreamNewFromNative(ret.Pointer())

	return retGo
}

// UNSUPPORTED : C value 'soup_multipart_input_stream_next_part_async' : parameter 'callback' of type 'Gio.AsyncReadyCallback' not supported

var multipartInputStreamNextPartFinishFunction *gi.Function
var multipartInputStreamNextPartFinishFunction_Once sync.Once

func multipartInputStreamNextPartFinishFunction_Set() error {
	var err error
	multipartInputStreamNextPartFinishFunction_Once.Do(func() {
		err = multipartInputStreamObject_Set()
		if err != nil {
			return
		}
		multipartInputStreamNextPartFinishFunction, err = multipartInputStreamObject.InvokerNew("next_part_finish")
	})
	return err
}

// NextPartFinish is a representation of the C type soup_multipart_input_stream_next_part_finish.
func (recv *MultipartInputStream) NextPartFinish(result *gio.AsyncResult) *gio.InputStream {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(result.Native())

	var ret gi.Argument

	err := multipartInputStreamNextPartFinishFunction_Set()
	if err == nil {
		ret = multipartInputStreamNextPartFinishFunction.Invoke(inArgs[:], nil)
	}

	retGo := gio.InputStreamNewFromNative(ret.Pointer())

	return retGo
}

// GioPollableInputStream returns the Gio.PollableInputStream interface implemented by MultipartInputStream
func (recv *MultipartInputStream) GioPollableInputStream() *gio.PollableInputStream {
	return gio.PollableInputStreamNewFromNative(recv.Native())
}

var proxyResolverDefaultObject *gi.Object
var proxyResolverDefaultObject_Once sync.Once

func proxyResolverDefaultObject_Set() error {
	var err error
	proxyResolverDefaultObject_Once.Do(func() {
		proxyResolverDefaultObject, err = gi.ObjectNew("Soup", "ProxyResolverDefault")
	})
	return err
}

type ProxyResolverDefault struct {
	native unsafe.Pointer
}

func ProxyResolverDefaultNewFromNative(native unsafe.Pointer) *ProxyResolverDefault {
	err := proxyResolverDefaultObject_Set()
	if err != nil {
		return nil
	}

	instance := &ProxyResolverDefault{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *ProxyResolverDefault) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToProxyResolverDefault down casts any arbitrary Object to ProxyResolverDefault.
Exercise care, as this is a potentially dangerous function
if the Object is not a ProxyResolverDefault.
*/
func CastToProxyResolverDefault(object *gobject.Object) *ProxyResolverDefault {
	return ProxyResolverDefaultNewFromNative(object.Native())
}

// Equals compares this ProxyResolverDefault with another ProxyResolverDefault, and returns true if they represent the same Object.
func (recv *ProxyResolverDefault) Equals(other *ProxyResolverDefault) bool {
	return other.Native() == recv.Native()
}

func (recv *ProxyResolverDefault) Native() unsafe.Pointer {
	return recv.native
}

// FieldParent returns the C field 'parent'.
func (recv *ProxyResolverDefault) FieldParent() *gobject.Object {
	argValue := gi.ObjectFieldGet(proxyResolverDefaultObject, recv.Native(), "parent")
	value := gobject.ObjectNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParent sets the value of the C field 'parent'.
func (recv *ProxyResolverDefault) SetFieldParent(value *gobject.Object) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(proxyResolverDefaultObject, recv.Native(), "parent", argValue)
}

// ProxyURIResolver returns the ProxyURIResolver interface implemented by ProxyResolverDefault
func (recv *ProxyResolverDefault) ProxyURIResolver() *ProxyURIResolver {
	return ProxyURIResolverNewFromNative(recv.Native())
}

// SessionFeature returns the SessionFeature interface implemented by ProxyResolverDefault
func (recv *ProxyResolverDefault) SessionFeature() *SessionFeature {
	return SessionFeatureNewFromNative(recv.Native())
}

var requestObject *gi.Object
var requestObject_Once sync.Once

func requestObject_Set() error {
	var err error
	requestObject_Once.Do(func() {
		requestObject, err = gi.ObjectNew("Soup", "Request")
	})
	return err
}

type Request struct {
	native unsafe.Pointer
}

func RequestNewFromNative(native unsafe.Pointer) *Request {
	err := requestObject_Set()
	if err != nil {
		return nil
	}

	instance := &Request{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *Request) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToRequest down casts any arbitrary Object to Request.
Exercise care, as this is a potentially dangerous function
if the Object is not a Request.
*/
func CastToRequest(object *gobject.Object) *Request {
	return RequestNewFromNative(object.Native())
}

// Equals compares this Request with another Request, and returns true if they represent the same Object.
func (recv *Request) Equals(other *Request) bool {
	return other.Native() == recv.Native()
}

func (recv *Request) Native() unsafe.Pointer {
	return recv.native
}

// FieldParent returns the C field 'parent'.
func (recv *Request) FieldParent() *gobject.Object {
	argValue := gi.ObjectFieldGet(requestObject, recv.Native(), "parent")
	value := gobject.ObjectNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParent sets the value of the C field 'parent'.
func (recv *Request) SetFieldParent(value *gobject.Object) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(requestObject, recv.Native(), "parent", argValue)
}

// FieldPriv returns the C field 'priv'.
func (recv *Request) FieldPriv() *RequestPrivate {
	argValue := gi.ObjectFieldGet(requestObject, recv.Native(), "priv")
	value := RequestPrivateNewFromNative(argValue.Pointer())
	return value
}

// SetFieldPriv sets the value of the C field 'priv'.
func (recv *Request) SetFieldPriv(value *RequestPrivate) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(requestObject, recv.Native(), "priv", argValue)
}

var requestGetContentLengthFunction *gi.Function
var requestGetContentLengthFunction_Once sync.Once

func requestGetContentLengthFunction_Set() error {
	var err error
	requestGetContentLengthFunction_Once.Do(func() {
		err = requestObject_Set()
		if err != nil {
			return
		}
		requestGetContentLengthFunction, err = requestObject.InvokerNew("get_content_length")
	})
	return err
}

// GetContentLength is a representation of the C type soup_request_get_content_length.
func (recv *Request) GetContentLength() int64 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := requestGetContentLengthFunction_Set()
	if err == nil {
		ret = requestGetContentLengthFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int64()

	return retGo
}

var requestGetContentTypeFunction *gi.Function
var requestGetContentTypeFunction_Once sync.Once

func requestGetContentTypeFunction_Set() error {
	var err error
	requestGetContentTypeFunction_Once.Do(func() {
		err = requestObject_Set()
		if err != nil {
			return
		}
		requestGetContentTypeFunction, err = requestObject.InvokerNew("get_content_type")
	})
	return err
}

// GetContentType is a representation of the C type soup_request_get_content_type.
func (recv *Request) GetContentType() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := requestGetContentTypeFunction_Set()
	if err == nil {
		ret = requestGetContentTypeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var requestGetSessionFunction *gi.Function
var requestGetSessionFunction_Once sync.Once

func requestGetSessionFunction_Set() error {
	var err error
	requestGetSessionFunction_Once.Do(func() {
		err = requestObject_Set()
		if err != nil {
			return
		}
		requestGetSessionFunction, err = requestObject.InvokerNew("get_session")
	})
	return err
}

// GetSession is a representation of the C type soup_request_get_session.
func (recv *Request) GetSession() *Session {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := requestGetSessionFunction_Set()
	if err == nil {
		ret = requestGetSessionFunction.Invoke(inArgs[:], nil)
	}

	retGo := SessionNewFromNative(ret.Pointer())

	return retGo
}

var requestGetUriFunction *gi.Function
var requestGetUriFunction_Once sync.Once

func requestGetUriFunction_Set() error {
	var err error
	requestGetUriFunction_Once.Do(func() {
		err = requestObject_Set()
		if err != nil {
			return
		}
		requestGetUriFunction, err = requestObject.InvokerNew("get_uri")
	})
	return err
}

// GetUri is a representation of the C type soup_request_get_uri.
func (recv *Request) GetUri() *URI {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := requestGetUriFunction_Set()
	if err == nil {
		ret = requestGetUriFunction.Invoke(inArgs[:], nil)
	}

	retGo := URINewFromNative(ret.Pointer())

	return retGo
}

var requestSendFunction *gi.Function
var requestSendFunction_Once sync.Once

func requestSendFunction_Set() error {
	var err error
	requestSendFunction_Once.Do(func() {
		err = requestObject_Set()
		if err != nil {
			return
		}
		requestSendFunction, err = requestObject.InvokerNew("send")
	})
	return err
}

// Send is a representation of the C type soup_request_send.
func (recv *Request) Send(cancellable *gio.Cancellable) *gio.InputStream {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(cancellable.Native())

	var ret gi.Argument

	err := requestSendFunction_Set()
	if err == nil {
		ret = requestSendFunction.Invoke(inArgs[:], nil)
	}

	retGo := gio.InputStreamNewFromNative(ret.Pointer())

	return retGo
}

// UNSUPPORTED : C value 'soup_request_send_async' : parameter 'callback' of type 'Gio.AsyncReadyCallback' not supported

var requestSendFinishFunction *gi.Function
var requestSendFinishFunction_Once sync.Once

func requestSendFinishFunction_Set() error {
	var err error
	requestSendFinishFunction_Once.Do(func() {
		err = requestObject_Set()
		if err != nil {
			return
		}
		requestSendFinishFunction, err = requestObject.InvokerNew("send_finish")
	})
	return err
}

// SendFinish is a representation of the C type soup_request_send_finish.
func (recv *Request) SendFinish(result *gio.AsyncResult) *gio.InputStream {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(result.Native())

	var ret gi.Argument

	err := requestSendFinishFunction_Set()
	if err == nil {
		ret = requestSendFinishFunction.Invoke(inArgs[:], nil)
	}

	retGo := gio.InputStreamNewFromNative(ret.Pointer())

	return retGo
}

// GioInitable returns the Gio.Initable interface implemented by Request
func (recv *Request) GioInitable() *gio.Initable {
	return gio.InitableNewFromNative(recv.Native())
}

var requestDataObject *gi.Object
var requestDataObject_Once sync.Once

func requestDataObject_Set() error {
	var err error
	requestDataObject_Once.Do(func() {
		requestDataObject, err = gi.ObjectNew("Soup", "RequestData")
	})
	return err
}

type RequestData struct {
	native unsafe.Pointer
}

func RequestDataNewFromNative(native unsafe.Pointer) *RequestData {
	err := requestDataObject_Set()
	if err != nil {
		return nil
	}

	instance := &RequestData{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Request upcasts to *Request
func (recv *RequestData) Request() *Request {
	return RequestNewFromNative(recv.Native())
}

// Object upcasts to *Object
func (recv *RequestData) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToRequestData down casts any arbitrary Object to RequestData.
Exercise care, as this is a potentially dangerous function
if the Object is not a RequestData.
*/
func CastToRequestData(object *gobject.Object) *RequestData {
	return RequestDataNewFromNative(object.Native())
}

// Equals compares this RequestData with another RequestData, and returns true if they represent the same Object.
func (recv *RequestData) Equals(other *RequestData) bool {
	return other.Native() == recv.Native()
}

func (recv *RequestData) Native() unsafe.Pointer {
	return recv.native
}

// FieldParent returns the C field 'parent'.
func (recv *RequestData) FieldParent() *Request {
	argValue := gi.ObjectFieldGet(requestDataObject, recv.Native(), "parent")
	value := RequestNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParent sets the value of the C field 'parent'.
func (recv *RequestData) SetFieldParent(value *Request) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(requestDataObject, recv.Native(), "parent", argValue)
}

// FieldPriv returns the C field 'priv'.
func (recv *RequestData) FieldPriv() *RequestDataPrivate {
	argValue := gi.ObjectFieldGet(requestDataObject, recv.Native(), "priv")
	value := RequestDataPrivateNewFromNative(argValue.Pointer())
	return value
}

// SetFieldPriv sets the value of the C field 'priv'.
func (recv *RequestData) SetFieldPriv(value *RequestDataPrivate) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(requestDataObject, recv.Native(), "priv", argValue)
}

// GioInitable returns the Gio.Initable interface implemented by RequestData
func (recv *RequestData) GioInitable() *gio.Initable {
	return gio.InitableNewFromNative(recv.Native())
}

var requestFileObject *gi.Object
var requestFileObject_Once sync.Once

func requestFileObject_Set() error {
	var err error
	requestFileObject_Once.Do(func() {
		requestFileObject, err = gi.ObjectNew("Soup", "RequestFile")
	})
	return err
}

type RequestFile struct {
	native unsafe.Pointer
}

func RequestFileNewFromNative(native unsafe.Pointer) *RequestFile {
	err := requestFileObject_Set()
	if err != nil {
		return nil
	}

	instance := &RequestFile{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Request upcasts to *Request
func (recv *RequestFile) Request() *Request {
	return RequestNewFromNative(recv.Native())
}

// Object upcasts to *Object
func (recv *RequestFile) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToRequestFile down casts any arbitrary Object to RequestFile.
Exercise care, as this is a potentially dangerous function
if the Object is not a RequestFile.
*/
func CastToRequestFile(object *gobject.Object) *RequestFile {
	return RequestFileNewFromNative(object.Native())
}

// Equals compares this RequestFile with another RequestFile, and returns true if they represent the same Object.
func (recv *RequestFile) Equals(other *RequestFile) bool {
	return other.Native() == recv.Native()
}

func (recv *RequestFile) Native() unsafe.Pointer {
	return recv.native
}

// FieldParent returns the C field 'parent'.
func (recv *RequestFile) FieldParent() *Request {
	argValue := gi.ObjectFieldGet(requestFileObject, recv.Native(), "parent")
	value := RequestNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParent sets the value of the C field 'parent'.
func (recv *RequestFile) SetFieldParent(value *Request) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(requestFileObject, recv.Native(), "parent", argValue)
}

// FieldPriv returns the C field 'priv'.
func (recv *RequestFile) FieldPriv() *RequestFilePrivate {
	argValue := gi.ObjectFieldGet(requestFileObject, recv.Native(), "priv")
	value := RequestFilePrivateNewFromNative(argValue.Pointer())
	return value
}

// SetFieldPriv sets the value of the C field 'priv'.
func (recv *RequestFile) SetFieldPriv(value *RequestFilePrivate) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(requestFileObject, recv.Native(), "priv", argValue)
}

var requestFileGetFileFunction *gi.Function
var requestFileGetFileFunction_Once sync.Once

func requestFileGetFileFunction_Set() error {
	var err error
	requestFileGetFileFunction_Once.Do(func() {
		err = requestFileObject_Set()
		if err != nil {
			return
		}
		requestFileGetFileFunction, err = requestFileObject.InvokerNew("get_file")
	})
	return err
}

// GetFile is a representation of the C type soup_request_file_get_file.
func (recv *RequestFile) GetFile() *gio.File {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := requestFileGetFileFunction_Set()
	if err == nil {
		ret = requestFileGetFileFunction.Invoke(inArgs[:], nil)
	}

	retGo := gio.FileNewFromNative(ret.Pointer())

	return retGo
}

// GioInitable returns the Gio.Initable interface implemented by RequestFile
func (recv *RequestFile) GioInitable() *gio.Initable {
	return gio.InitableNewFromNative(recv.Native())
}

var requestHTTPObject *gi.Object
var requestHTTPObject_Once sync.Once

func requestHTTPObject_Set() error {
	var err error
	requestHTTPObject_Once.Do(func() {
		requestHTTPObject, err = gi.ObjectNew("Soup", "RequestHTTP")
	})
	return err
}

type RequestHTTP struct {
	native unsafe.Pointer
}

func RequestHTTPNewFromNative(native unsafe.Pointer) *RequestHTTP {
	err := requestHTTPObject_Set()
	if err != nil {
		return nil
	}

	instance := &RequestHTTP{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Request upcasts to *Request
func (recv *RequestHTTP) Request() *Request {
	return RequestNewFromNative(recv.Native())
}

// Object upcasts to *Object
func (recv *RequestHTTP) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToRequestHTTP down casts any arbitrary Object to RequestHTTP.
Exercise care, as this is a potentially dangerous function
if the Object is not a RequestHTTP.
*/
func CastToRequestHTTP(object *gobject.Object) *RequestHTTP {
	return RequestHTTPNewFromNative(object.Native())
}

// Equals compares this RequestHTTP with another RequestHTTP, and returns true if they represent the same Object.
func (recv *RequestHTTP) Equals(other *RequestHTTP) bool {
	return other.Native() == recv.Native()
}

func (recv *RequestHTTP) Native() unsafe.Pointer {
	return recv.native
}

// FieldParent returns the C field 'parent'.
func (recv *RequestHTTP) FieldParent() *Request {
	argValue := gi.ObjectFieldGet(requestHTTPObject, recv.Native(), "parent")
	value := RequestNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParent sets the value of the C field 'parent'.
func (recv *RequestHTTP) SetFieldParent(value *Request) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(requestHTTPObject, recv.Native(), "parent", argValue)
}

// FieldPriv returns the C field 'priv'.
func (recv *RequestHTTP) FieldPriv() *RequestHTTPPrivate {
	argValue := gi.ObjectFieldGet(requestHTTPObject, recv.Native(), "priv")
	value := RequestHTTPPrivateNewFromNative(argValue.Pointer())
	return value
}

// SetFieldPriv sets the value of the C field 'priv'.
func (recv *RequestHTTP) SetFieldPriv(value *RequestHTTPPrivate) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(requestHTTPObject, recv.Native(), "priv", argValue)
}

var requestHTTPGetMessageFunction *gi.Function
var requestHTTPGetMessageFunction_Once sync.Once

func requestHTTPGetMessageFunction_Set() error {
	var err error
	requestHTTPGetMessageFunction_Once.Do(func() {
		err = requestHTTPObject_Set()
		if err != nil {
			return
		}
		requestHTTPGetMessageFunction, err = requestHTTPObject.InvokerNew("get_message")
	})
	return err
}

// GetMessage is a representation of the C type soup_request_http_get_message.
func (recv *RequestHTTP) GetMessage() *Message {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := requestHTTPGetMessageFunction_Set()
	if err == nil {
		ret = requestHTTPGetMessageFunction.Invoke(inArgs[:], nil)
	}

	retGo := MessageNewFromNative(ret.Pointer())

	return retGo
}

// GioInitable returns the Gio.Initable interface implemented by RequestHTTP
func (recv *RequestHTTP) GioInitable() *gio.Initable {
	return gio.InitableNewFromNative(recv.Native())
}

var requesterObject *gi.Object
var requesterObject_Once sync.Once

func requesterObject_Set() error {
	var err error
	requesterObject_Once.Do(func() {
		requesterObject, err = gi.ObjectNew("Soup", "Requester")
	})
	return err
}

type Requester struct {
	native unsafe.Pointer
}

func RequesterNewFromNative(native unsafe.Pointer) *Requester {
	err := requesterObject_Set()
	if err != nil {
		return nil
	}

	instance := &Requester{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *Requester) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToRequester down casts any arbitrary Object to Requester.
Exercise care, as this is a potentially dangerous function
if the Object is not a Requester.
*/
func CastToRequester(object *gobject.Object) *Requester {
	return RequesterNewFromNative(object.Native())
}

// Equals compares this Requester with another Requester, and returns true if they represent the same Object.
func (recv *Requester) Equals(other *Requester) bool {
	return other.Native() == recv.Native()
}

func (recv *Requester) Native() unsafe.Pointer {
	return recv.native
}

// FieldParent returns the C field 'parent'.
func (recv *Requester) FieldParent() *gobject.Object {
	argValue := gi.ObjectFieldGet(requesterObject, recv.Native(), "parent")
	value := gobject.ObjectNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParent sets the value of the C field 'parent'.
func (recv *Requester) SetFieldParent(value *gobject.Object) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(requesterObject, recv.Native(), "parent", argValue)
}

// FieldPriv returns the C field 'priv'.
func (recv *Requester) FieldPriv() *RequesterPrivate {
	argValue := gi.ObjectFieldGet(requesterObject, recv.Native(), "priv")
	value := RequesterPrivateNewFromNative(argValue.Pointer())
	return value
}

// SetFieldPriv sets the value of the C field 'priv'.
func (recv *Requester) SetFieldPriv(value *RequesterPrivate) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(requesterObject, recv.Native(), "priv", argValue)
}

var requesterNewFunction *gi.Function
var requesterNewFunction_Once sync.Once

func requesterNewFunction_Set() error {
	var err error
	requesterNewFunction_Once.Do(func() {
		err = requesterObject_Set()
		if err != nil {
			return
		}
		requesterNewFunction, err = requesterObject.InvokerNew("new")
	})
	return err
}

// RequesterNew is a representation of the C type soup_requester_new.
func RequesterNew() *Requester {

	var ret gi.Argument

	err := requesterNewFunction_Set()
	if err == nil {
		ret = requesterNewFunction.Invoke(nil, nil)
	}

	retGo := RequesterNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

var requesterRequestFunction *gi.Function
var requesterRequestFunction_Once sync.Once

func requesterRequestFunction_Set() error {
	var err error
	requesterRequestFunction_Once.Do(func() {
		err = requesterObject_Set()
		if err != nil {
			return
		}
		requesterRequestFunction, err = requesterObject.InvokerNew("request")
	})
	return err
}

// Request_ is a representation of the C type soup_requester_request.
func (recv *Requester) Request_(uriString string) *Request {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(uriString)

	var ret gi.Argument

	err := requesterRequestFunction_Set()
	if err == nil {
		ret = requesterRequestFunction.Invoke(inArgs[:], nil)
	}

	retGo := RequestNewFromNative(ret.Pointer())

	return retGo
}

var requesterRequestUriFunction *gi.Function
var requesterRequestUriFunction_Once sync.Once

func requesterRequestUriFunction_Set() error {
	var err error
	requesterRequestUriFunction_Once.Do(func() {
		err = requesterObject_Set()
		if err != nil {
			return
		}
		requesterRequestUriFunction, err = requesterObject.InvokerNew("request_uri")
	})
	return err
}

// RequestUri is a representation of the C type soup_requester_request_uri.
func (recv *Requester) RequestUri(uri *URI) *Request {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(uri.Native())

	var ret gi.Argument

	err := requesterRequestUriFunction_Set()
	if err == nil {
		ret = requesterRequestUriFunction.Invoke(inArgs[:], nil)
	}

	retGo := RequestNewFromNative(ret.Pointer())

	return retGo
}

// SessionFeature returns the SessionFeature interface implemented by Requester
func (recv *Requester) SessionFeature() *SessionFeature {
	return SessionFeatureNewFromNative(recv.Native())
}

var serverObject *gi.Object
var serverObject_Once sync.Once

func serverObject_Set() error {
	var err error
	serverObject_Once.Do(func() {
		serverObject, err = gi.ObjectNew("Soup", "Server")
	})
	return err
}

type Server struct {
	native unsafe.Pointer
}

func ServerNewFromNative(native unsafe.Pointer) *Server {
	err := serverObject_Set()
	if err != nil {
		return nil
	}

	instance := &Server{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *Server) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToServer down casts any arbitrary Object to Server.
Exercise care, as this is a potentially dangerous function
if the Object is not a Server.
*/
func CastToServer(object *gobject.Object) *Server {
	return ServerNewFromNative(object.Native())
}

// Equals compares this Server with another Server, and returns true if they represent the same Object.
func (recv *Server) Equals(other *Server) bool {
	return other.Native() == recv.Native()
}

func (recv *Server) Native() unsafe.Pointer {
	return recv.native
}

// FieldParent returns the C field 'parent'.
func (recv *Server) FieldParent() *gobject.Object {
	argValue := gi.ObjectFieldGet(serverObject, recv.Native(), "parent")
	value := gobject.ObjectNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParent sets the value of the C field 'parent'.
func (recv *Server) SetFieldParent(value *gobject.Object) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(serverObject, recv.Native(), "parent", argValue)
}

// UNSUPPORTED : C value 'soup_server_new' : parameter '...' of type 'nil' not supported

var serverAcceptIostreamFunction *gi.Function
var serverAcceptIostreamFunction_Once sync.Once

func serverAcceptIostreamFunction_Set() error {
	var err error
	serverAcceptIostreamFunction_Once.Do(func() {
		err = serverObject_Set()
		if err != nil {
			return
		}
		serverAcceptIostreamFunction, err = serverObject.InvokerNew("accept_iostream")
	})
	return err
}

// AcceptIostream is a representation of the C type soup_server_accept_iostream.
func (recv *Server) AcceptIostream(stream *gio.IOStream, localAddr *gio.SocketAddress, remoteAddr *gio.SocketAddress) bool {
	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(stream.Native())
	inArgs[2].SetPointer(localAddr.Native())
	inArgs[3].SetPointer(remoteAddr.Native())

	var ret gi.Argument

	err := serverAcceptIostreamFunction_Set()
	if err == nil {
		ret = serverAcceptIostreamFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var serverAddAuthDomainFunction *gi.Function
var serverAddAuthDomainFunction_Once sync.Once

func serverAddAuthDomainFunction_Set() error {
	var err error
	serverAddAuthDomainFunction_Once.Do(func() {
		err = serverObject_Set()
		if err != nil {
			return
		}
		serverAddAuthDomainFunction, err = serverObject.InvokerNew("add_auth_domain")
	})
	return err
}

// AddAuthDomain is a representation of the C type soup_server_add_auth_domain.
func (recv *Server) AddAuthDomain(authDomain *AuthDomain) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(authDomain.Native())

	err := serverAddAuthDomainFunction_Set()
	if err == nil {
		serverAddAuthDomainFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'soup_server_add_early_handler' : parameter 'callback' of type 'ServerCallback' not supported

// UNSUPPORTED : C value 'soup_server_add_handler' : parameter 'callback' of type 'ServerCallback' not supported

var serverAddWebsocketExtensionFunction *gi.Function
var serverAddWebsocketExtensionFunction_Once sync.Once

func serverAddWebsocketExtensionFunction_Set() error {
	var err error
	serverAddWebsocketExtensionFunction_Once.Do(func() {
		err = serverObject_Set()
		if err != nil {
			return
		}
		serverAddWebsocketExtensionFunction, err = serverObject.InvokerNew("add_websocket_extension")
	})
	return err
}

// AddWebsocketExtension is a representation of the C type soup_server_add_websocket_extension.
func (recv *Server) AddWebsocketExtension(extensionType int64) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt64(extensionType)

	err := serverAddWebsocketExtensionFunction_Set()
	if err == nil {
		serverAddWebsocketExtensionFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'soup_server_add_websocket_handler' : parameter 'callback' of type 'ServerWebsocketCallback' not supported

var serverDisconnectFunction *gi.Function
var serverDisconnectFunction_Once sync.Once

func serverDisconnectFunction_Set() error {
	var err error
	serverDisconnectFunction_Once.Do(func() {
		err = serverObject_Set()
		if err != nil {
			return
		}
		serverDisconnectFunction, err = serverObject.InvokerNew("disconnect")
	})
	return err
}

// Disconnect is a representation of the C type soup_server_disconnect.
func (recv *Server) Disconnect() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := serverDisconnectFunction_Set()
	if err == nil {
		serverDisconnectFunction.Invoke(inArgs[:], nil)
	}

	return
}

var serverGetAsyncContextFunction *gi.Function
var serverGetAsyncContextFunction_Once sync.Once

func serverGetAsyncContextFunction_Set() error {
	var err error
	serverGetAsyncContextFunction_Once.Do(func() {
		err = serverObject_Set()
		if err != nil {
			return
		}
		serverGetAsyncContextFunction, err = serverObject.InvokerNew("get_async_context")
	})
	return err
}

// GetAsyncContext is a representation of the C type soup_server_get_async_context.
func (recv *Server) GetAsyncContext() *glib.MainContext {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := serverGetAsyncContextFunction_Set()
	if err == nil {
		ret = serverGetAsyncContextFunction.Invoke(inArgs[:], nil)
	}

	retGo := glib.MainContextNewFromNative(ret.Pointer())

	return retGo
}

var serverGetListenerFunction *gi.Function
var serverGetListenerFunction_Once sync.Once

func serverGetListenerFunction_Set() error {
	var err error
	serverGetListenerFunction_Once.Do(func() {
		err = serverObject_Set()
		if err != nil {
			return
		}
		serverGetListenerFunction, err = serverObject.InvokerNew("get_listener")
	})
	return err
}

// GetListener is a representation of the C type soup_server_get_listener.
func (recv *Server) GetListener() *Socket {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := serverGetListenerFunction_Set()
	if err == nil {
		ret = serverGetListenerFunction.Invoke(inArgs[:], nil)
	}

	retGo := SocketNewFromNative(ret.Pointer())

	return retGo
}

var serverGetListenersFunction *gi.Function
var serverGetListenersFunction_Once sync.Once

func serverGetListenersFunction_Set() error {
	var err error
	serverGetListenersFunction_Once.Do(func() {
		err = serverObject_Set()
		if err != nil {
			return
		}
		serverGetListenersFunction, err = serverObject.InvokerNew("get_listeners")
	})
	return err
}

// GetListeners is a representation of the C type soup_server_get_listeners.
func (recv *Server) GetListeners() *glib.SList {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := serverGetListenersFunction_Set()
	if err == nil {
		ret = serverGetListenersFunction.Invoke(inArgs[:], nil)
	}

	retGo := glib.SListNewFromNative(ret.Pointer())

	return retGo
}

var serverGetPortFunction *gi.Function
var serverGetPortFunction_Once sync.Once

func serverGetPortFunction_Set() error {
	var err error
	serverGetPortFunction_Once.Do(func() {
		err = serverObject_Set()
		if err != nil {
			return
		}
		serverGetPortFunction, err = serverObject.InvokerNew("get_port")
	})
	return err
}

// GetPort is a representation of the C type soup_server_get_port.
func (recv *Server) GetPort() uint32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := serverGetPortFunction_Set()
	if err == nil {
		ret = serverGetPortFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo
}

var serverGetUrisFunction *gi.Function
var serverGetUrisFunction_Once sync.Once

func serverGetUrisFunction_Set() error {
	var err error
	serverGetUrisFunction_Once.Do(func() {
		err = serverObject_Set()
		if err != nil {
			return
		}
		serverGetUrisFunction, err = serverObject.InvokerNew("get_uris")
	})
	return err
}

// GetUris is a representation of the C type soup_server_get_uris.
func (recv *Server) GetUris() *glib.SList {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := serverGetUrisFunction_Set()
	if err == nil {
		ret = serverGetUrisFunction.Invoke(inArgs[:], nil)
	}

	retGo := glib.SListNewFromNative(ret.Pointer())

	return retGo
}

var serverIsHttpsFunction *gi.Function
var serverIsHttpsFunction_Once sync.Once

func serverIsHttpsFunction_Set() error {
	var err error
	serverIsHttpsFunction_Once.Do(func() {
		err = serverObject_Set()
		if err != nil {
			return
		}
		serverIsHttpsFunction, err = serverObject.InvokerNew("is_https")
	})
	return err
}

// IsHttps is a representation of the C type soup_server_is_https.
func (recv *Server) IsHttps() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := serverIsHttpsFunction_Set()
	if err == nil {
		ret = serverIsHttpsFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var serverListenFunction *gi.Function
var serverListenFunction_Once sync.Once

func serverListenFunction_Set() error {
	var err error
	serverListenFunction_Once.Do(func() {
		err = serverObject_Set()
		if err != nil {
			return
		}
		serverListenFunction, err = serverObject.InvokerNew("listen")
	})
	return err
}

// Listen is a representation of the C type soup_server_listen.
func (recv *Server) Listen(address *gio.SocketAddress, options ServerListenOptions) bool {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(address.Native())
	inArgs[2].SetInt32(int32(options))

	var ret gi.Argument

	err := serverListenFunction_Set()
	if err == nil {
		ret = serverListenFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var serverListenAllFunction *gi.Function
var serverListenAllFunction_Once sync.Once

func serverListenAllFunction_Set() error {
	var err error
	serverListenAllFunction_Once.Do(func() {
		err = serverObject_Set()
		if err != nil {
			return
		}
		serverListenAllFunction, err = serverObject.InvokerNew("listen_all")
	})
	return err
}

// ListenAll is a representation of the C type soup_server_listen_all.
func (recv *Server) ListenAll(port uint32, options ServerListenOptions) bool {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetUint32(port)
	inArgs[2].SetInt32(int32(options))

	var ret gi.Argument

	err := serverListenAllFunction_Set()
	if err == nil {
		ret = serverListenAllFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var serverListenFdFunction *gi.Function
var serverListenFdFunction_Once sync.Once

func serverListenFdFunction_Set() error {
	var err error
	serverListenFdFunction_Once.Do(func() {
		err = serverObject_Set()
		if err != nil {
			return
		}
		serverListenFdFunction, err = serverObject.InvokerNew("listen_fd")
	})
	return err
}

// ListenFd is a representation of the C type soup_server_listen_fd.
func (recv *Server) ListenFd(fd int32, options ServerListenOptions) bool {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(fd)
	inArgs[2].SetInt32(int32(options))

	var ret gi.Argument

	err := serverListenFdFunction_Set()
	if err == nil {
		ret = serverListenFdFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var serverListenLocalFunction *gi.Function
var serverListenLocalFunction_Once sync.Once

func serverListenLocalFunction_Set() error {
	var err error
	serverListenLocalFunction_Once.Do(func() {
		err = serverObject_Set()
		if err != nil {
			return
		}
		serverListenLocalFunction, err = serverObject.InvokerNew("listen_local")
	})
	return err
}

// ListenLocal is a representation of the C type soup_server_listen_local.
func (recv *Server) ListenLocal(port uint32, options ServerListenOptions) bool {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetUint32(port)
	inArgs[2].SetInt32(int32(options))

	var ret gi.Argument

	err := serverListenLocalFunction_Set()
	if err == nil {
		ret = serverListenLocalFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var serverListenSocketFunction *gi.Function
var serverListenSocketFunction_Once sync.Once

func serverListenSocketFunction_Set() error {
	var err error
	serverListenSocketFunction_Once.Do(func() {
		err = serverObject_Set()
		if err != nil {
			return
		}
		serverListenSocketFunction, err = serverObject.InvokerNew("listen_socket")
	})
	return err
}

// ListenSocket is a representation of the C type soup_server_listen_socket.
func (recv *Server) ListenSocket(socket *gio.Socket, options ServerListenOptions) bool {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(socket.Native())
	inArgs[2].SetInt32(int32(options))

	var ret gi.Argument

	err := serverListenSocketFunction_Set()
	if err == nil {
		ret = serverListenSocketFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var serverPauseMessageFunction *gi.Function
var serverPauseMessageFunction_Once sync.Once

func serverPauseMessageFunction_Set() error {
	var err error
	serverPauseMessageFunction_Once.Do(func() {
		err = serverObject_Set()
		if err != nil {
			return
		}
		serverPauseMessageFunction, err = serverObject.InvokerNew("pause_message")
	})
	return err
}

// PauseMessage is a representation of the C type soup_server_pause_message.
func (recv *Server) PauseMessage(msg *Message) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(msg.Native())

	err := serverPauseMessageFunction_Set()
	if err == nil {
		serverPauseMessageFunction.Invoke(inArgs[:], nil)
	}

	return
}

var serverQuitFunction *gi.Function
var serverQuitFunction_Once sync.Once

func serverQuitFunction_Set() error {
	var err error
	serverQuitFunction_Once.Do(func() {
		err = serverObject_Set()
		if err != nil {
			return
		}
		serverQuitFunction, err = serverObject.InvokerNew("quit")
	})
	return err
}

// Quit is a representation of the C type soup_server_quit.
func (recv *Server) Quit() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := serverQuitFunction_Set()
	if err == nil {
		serverQuitFunction.Invoke(inArgs[:], nil)
	}

	return
}

var serverRemoveAuthDomainFunction *gi.Function
var serverRemoveAuthDomainFunction_Once sync.Once

func serverRemoveAuthDomainFunction_Set() error {
	var err error
	serverRemoveAuthDomainFunction_Once.Do(func() {
		err = serverObject_Set()
		if err != nil {
			return
		}
		serverRemoveAuthDomainFunction, err = serverObject.InvokerNew("remove_auth_domain")
	})
	return err
}

// RemoveAuthDomain is a representation of the C type soup_server_remove_auth_domain.
func (recv *Server) RemoveAuthDomain(authDomain *AuthDomain) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(authDomain.Native())

	err := serverRemoveAuthDomainFunction_Set()
	if err == nil {
		serverRemoveAuthDomainFunction.Invoke(inArgs[:], nil)
	}

	return
}

var serverRemoveHandlerFunction *gi.Function
var serverRemoveHandlerFunction_Once sync.Once

func serverRemoveHandlerFunction_Set() error {
	var err error
	serverRemoveHandlerFunction_Once.Do(func() {
		err = serverObject_Set()
		if err != nil {
			return
		}
		serverRemoveHandlerFunction, err = serverObject.InvokerNew("remove_handler")
	})
	return err
}

// RemoveHandler is a representation of the C type soup_server_remove_handler.
func (recv *Server) RemoveHandler(path string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(path)

	err := serverRemoveHandlerFunction_Set()
	if err == nil {
		serverRemoveHandlerFunction.Invoke(inArgs[:], nil)
	}

	return
}

var serverRemoveWebsocketExtensionFunction *gi.Function
var serverRemoveWebsocketExtensionFunction_Once sync.Once

func serverRemoveWebsocketExtensionFunction_Set() error {
	var err error
	serverRemoveWebsocketExtensionFunction_Once.Do(func() {
		err = serverObject_Set()
		if err != nil {
			return
		}
		serverRemoveWebsocketExtensionFunction, err = serverObject.InvokerNew("remove_websocket_extension")
	})
	return err
}

// RemoveWebsocketExtension is a representation of the C type soup_server_remove_websocket_extension.
func (recv *Server) RemoveWebsocketExtension(extensionType int64) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt64(extensionType)

	err := serverRemoveWebsocketExtensionFunction_Set()
	if err == nil {
		serverRemoveWebsocketExtensionFunction.Invoke(inArgs[:], nil)
	}

	return
}

var serverRunFunction *gi.Function
var serverRunFunction_Once sync.Once

func serverRunFunction_Set() error {
	var err error
	serverRunFunction_Once.Do(func() {
		err = serverObject_Set()
		if err != nil {
			return
		}
		serverRunFunction, err = serverObject.InvokerNew("run")
	})
	return err
}

// Run is a representation of the C type soup_server_run.
func (recv *Server) Run() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := serverRunFunction_Set()
	if err == nil {
		serverRunFunction.Invoke(inArgs[:], nil)
	}

	return
}

var serverRunAsyncFunction *gi.Function
var serverRunAsyncFunction_Once sync.Once

func serverRunAsyncFunction_Set() error {
	var err error
	serverRunAsyncFunction_Once.Do(func() {
		err = serverObject_Set()
		if err != nil {
			return
		}
		serverRunAsyncFunction, err = serverObject.InvokerNew("run_async")
	})
	return err
}

// RunAsync is a representation of the C type soup_server_run_async.
func (recv *Server) RunAsync() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := serverRunAsyncFunction_Set()
	if err == nil {
		serverRunAsyncFunction.Invoke(inArgs[:], nil)
	}

	return
}

var serverSetSslCertFileFunction *gi.Function
var serverSetSslCertFileFunction_Once sync.Once

func serverSetSslCertFileFunction_Set() error {
	var err error
	serverSetSslCertFileFunction_Once.Do(func() {
		err = serverObject_Set()
		if err != nil {
			return
		}
		serverSetSslCertFileFunction, err = serverObject.InvokerNew("set_ssl_cert_file")
	})
	return err
}

// SetSslCertFile is a representation of the C type soup_server_set_ssl_cert_file.
func (recv *Server) SetSslCertFile(sslCertFile string, sslKeyFile string) bool {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(sslCertFile)
	inArgs[2].SetString(sslKeyFile)

	var ret gi.Argument

	err := serverSetSslCertFileFunction_Set()
	if err == nil {
		ret = serverSetSslCertFileFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var serverUnpauseMessageFunction *gi.Function
var serverUnpauseMessageFunction_Once sync.Once

func serverUnpauseMessageFunction_Set() error {
	var err error
	serverUnpauseMessageFunction_Once.Do(func() {
		err = serverObject_Set()
		if err != nil {
			return
		}
		serverUnpauseMessageFunction, err = serverObject.InvokerNew("unpause_message")
	})
	return err
}

// UnpauseMessage is a representation of the C type soup_server_unpause_message.
func (recv *Server) UnpauseMessage(msg *Message) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(msg.Native())

	err := serverUnpauseMessageFunction_Set()
	if err == nil {
		serverUnpauseMessageFunction.Invoke(inArgs[:], nil)
	}

	return
}

/*
ConnectRequestAborted connects a callback to the 'request-aborted' signal of the Server.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *Server) ConnectRequestAborted(handler func(instance *Server, message *Message, client *ClientContext)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {
		objectInstance := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[0]))
		argInstance := ServerNewFromNative(objectInstance.GetObject().Native())

		object1 := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[1]))
		arg1 := MessageNewFromNative(object1.GetObject().Native())

		object2 := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[2]))
		arg2 := ClientContextNewFromNative(object2.GetBoxed())

		handler(argInstance, arg1, arg2)

	}

	return callback.ConnectSignal(recv.Native(), "request-aborted", marshal)
}

/*
ConnectRequestFinished connects a callback to the 'request-finished' signal of the Server.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *Server) ConnectRequestFinished(handler func(instance *Server, message *Message, client *ClientContext)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {
		objectInstance := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[0]))
		argInstance := ServerNewFromNative(objectInstance.GetObject().Native())

		object1 := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[1]))
		arg1 := MessageNewFromNative(object1.GetObject().Native())

		object2 := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[2]))
		arg2 := ClientContextNewFromNative(object2.GetBoxed())

		handler(argInstance, arg1, arg2)

	}

	return callback.ConnectSignal(recv.Native(), "request-finished", marshal)
}

/*
ConnectRequestRead connects a callback to the 'request-read' signal of the Server.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *Server) ConnectRequestRead(handler func(instance *Server, message *Message, client *ClientContext)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {
		objectInstance := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[0]))
		argInstance := ServerNewFromNative(objectInstance.GetObject().Native())

		object1 := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[1]))
		arg1 := MessageNewFromNative(object1.GetObject().Native())

		object2 := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[2]))
		arg2 := ClientContextNewFromNative(object2.GetBoxed())

		handler(argInstance, arg1, arg2)

	}

	return callback.ConnectSignal(recv.Native(), "request-read", marshal)
}

/*
ConnectRequestStarted connects a callback to the 'request-started' signal of the Server.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *Server) ConnectRequestStarted(handler func(instance *Server, message *Message, client *ClientContext)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {
		objectInstance := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[0]))
		argInstance := ServerNewFromNative(objectInstance.GetObject().Native())

		object1 := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[1]))
		arg1 := MessageNewFromNative(object1.GetObject().Native())

		object2 := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[2]))
		arg2 := ClientContextNewFromNative(object2.GetBoxed())

		handler(argInstance, arg1, arg2)

	}

	return callback.ConnectSignal(recv.Native(), "request-started", marshal)
}

/*
Disconnect disconnects a callback previously registered with a Connect...() method.

The connectionID should be a value returned from a call to a Connect...() method.
*/
func (recv *Server) DisconnectSignal(connectionID int) {
	callback.DisconnectSignal(connectionID)
}

var sessionObject *gi.Object
var sessionObject_Once sync.Once

func sessionObject_Set() error {
	var err error
	sessionObject_Once.Do(func() {
		sessionObject, err = gi.ObjectNew("Soup", "Session")
	})
	return err
}

type Session struct {
	native unsafe.Pointer
}

func SessionNewFromNative(native unsafe.Pointer) *Session {
	err := sessionObject_Set()
	if err != nil {
		return nil
	}

	instance := &Session{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *Session) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToSession down casts any arbitrary Object to Session.
Exercise care, as this is a potentially dangerous function
if the Object is not a Session.
*/
func CastToSession(object *gobject.Object) *Session {
	return SessionNewFromNative(object.Native())
}

// Equals compares this Session with another Session, and returns true if they represent the same Object.
func (recv *Session) Equals(other *Session) bool {
	return other.Native() == recv.Native()
}

func (recv *Session) Native() unsafe.Pointer {
	return recv.native
}

// FieldParent returns the C field 'parent'.
func (recv *Session) FieldParent() *gobject.Object {
	argValue := gi.ObjectFieldGet(sessionObject, recv.Native(), "parent")
	value := gobject.ObjectNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParent sets the value of the C field 'parent'.
func (recv *Session) SetFieldParent(value *gobject.Object) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(sessionObject, recv.Native(), "parent", argValue)
}

var sessionNewFunction *gi.Function
var sessionNewFunction_Once sync.Once

func sessionNewFunction_Set() error {
	var err error
	sessionNewFunction_Once.Do(func() {
		err = sessionObject_Set()
		if err != nil {
			return
		}
		sessionNewFunction, err = sessionObject.InvokerNew("new")
	})
	return err
}

// SessionNew is a representation of the C type soup_session_new.
func SessionNew() *Session {

	var ret gi.Argument

	err := sessionNewFunction_Set()
	if err == nil {
		ret = sessionNewFunction.Invoke(nil, nil)
	}

	retGo := SessionNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

// UNSUPPORTED : C value 'soup_session_new_with_options' : parameter '...' of type 'nil' not supported

var sessionAbortFunction *gi.Function
var sessionAbortFunction_Once sync.Once

func sessionAbortFunction_Set() error {
	var err error
	sessionAbortFunction_Once.Do(func() {
		err = sessionObject_Set()
		if err != nil {
			return
		}
		sessionAbortFunction, err = sessionObject.InvokerNew("abort")
	})
	return err
}

// Abort is a representation of the C type soup_session_abort.
func (recv *Session) Abort() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := sessionAbortFunction_Set()
	if err == nil {
		sessionAbortFunction.Invoke(inArgs[:], nil)
	}

	return
}

var sessionAddFeatureFunction *gi.Function
var sessionAddFeatureFunction_Once sync.Once

func sessionAddFeatureFunction_Set() error {
	var err error
	sessionAddFeatureFunction_Once.Do(func() {
		err = sessionObject_Set()
		if err != nil {
			return
		}
		sessionAddFeatureFunction, err = sessionObject.InvokerNew("add_feature")
	})
	return err
}

// AddFeature is a representation of the C type soup_session_add_feature.
func (recv *Session) AddFeature(feature *SessionFeature) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(feature.Native())

	err := sessionAddFeatureFunction_Set()
	if err == nil {
		sessionAddFeatureFunction.Invoke(inArgs[:], nil)
	}

	return
}

var sessionAddFeatureByTypeFunction *gi.Function
var sessionAddFeatureByTypeFunction_Once sync.Once

func sessionAddFeatureByTypeFunction_Set() error {
	var err error
	sessionAddFeatureByTypeFunction_Once.Do(func() {
		err = sessionObject_Set()
		if err != nil {
			return
		}
		sessionAddFeatureByTypeFunction, err = sessionObject.InvokerNew("add_feature_by_type")
	})
	return err
}

// AddFeatureByType is a representation of the C type soup_session_add_feature_by_type.
func (recv *Session) AddFeatureByType(featureType int64) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt64(featureType)

	err := sessionAddFeatureByTypeFunction_Set()
	if err == nil {
		sessionAddFeatureByTypeFunction.Invoke(inArgs[:], nil)
	}

	return
}

var sessionCancelMessageFunction *gi.Function
var sessionCancelMessageFunction_Once sync.Once

func sessionCancelMessageFunction_Set() error {
	var err error
	sessionCancelMessageFunction_Once.Do(func() {
		err = sessionObject_Set()
		if err != nil {
			return
		}
		sessionCancelMessageFunction, err = sessionObject.InvokerNew("cancel_message")
	})
	return err
}

// CancelMessage is a representation of the C type soup_session_cancel_message.
func (recv *Session) CancelMessage(msg *Message, statusCode uint32) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(msg.Native())
	inArgs[2].SetUint32(statusCode)

	err := sessionCancelMessageFunction_Set()
	if err == nil {
		sessionCancelMessageFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'soup_session_connect_async' : parameter 'progress_callback' of type 'SessionConnectProgressCallback' not supported

var sessionConnectFinishFunction *gi.Function
var sessionConnectFinishFunction_Once sync.Once

func sessionConnectFinishFunction_Set() error {
	var err error
	sessionConnectFinishFunction_Once.Do(func() {
		err = sessionObject_Set()
		if err != nil {
			return
		}
		sessionConnectFinishFunction, err = sessionObject.InvokerNew("connect_finish")
	})
	return err
}

// ConnectFinish is a representation of the C type soup_session_connect_finish.
func (recv *Session) ConnectFinish(result *gio.AsyncResult) *gio.IOStream {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(result.Native())

	var ret gi.Argument

	err := sessionConnectFinishFunction_Set()
	if err == nil {
		ret = sessionConnectFinishFunction.Invoke(inArgs[:], nil)
	}

	retGo := gio.IOStreamNewFromNative(ret.Pointer())

	return retGo
}

var sessionGetAsyncContextFunction *gi.Function
var sessionGetAsyncContextFunction_Once sync.Once

func sessionGetAsyncContextFunction_Set() error {
	var err error
	sessionGetAsyncContextFunction_Once.Do(func() {
		err = sessionObject_Set()
		if err != nil {
			return
		}
		sessionGetAsyncContextFunction, err = sessionObject.InvokerNew("get_async_context")
	})
	return err
}

// GetAsyncContext is a representation of the C type soup_session_get_async_context.
func (recv *Session) GetAsyncContext() *glib.MainContext {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := sessionGetAsyncContextFunction_Set()
	if err == nil {
		ret = sessionGetAsyncContextFunction.Invoke(inArgs[:], nil)
	}

	retGo := glib.MainContextNewFromNative(ret.Pointer())

	return retGo
}

var sessionGetFeatureFunction *gi.Function
var sessionGetFeatureFunction_Once sync.Once

func sessionGetFeatureFunction_Set() error {
	var err error
	sessionGetFeatureFunction_Once.Do(func() {
		err = sessionObject_Set()
		if err != nil {
			return
		}
		sessionGetFeatureFunction, err = sessionObject.InvokerNew("get_feature")
	})
	return err
}

// GetFeature is a representation of the C type soup_session_get_feature.
func (recv *Session) GetFeature(featureType int64) *SessionFeature {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt64(featureType)

	var ret gi.Argument

	err := sessionGetFeatureFunction_Set()
	if err == nil {
		ret = sessionGetFeatureFunction.Invoke(inArgs[:], nil)
	}

	retGo := SessionFeatureNewFromNative(ret.Pointer())

	return retGo
}

var sessionGetFeatureForMessageFunction *gi.Function
var sessionGetFeatureForMessageFunction_Once sync.Once

func sessionGetFeatureForMessageFunction_Set() error {
	var err error
	sessionGetFeatureForMessageFunction_Once.Do(func() {
		err = sessionObject_Set()
		if err != nil {
			return
		}
		sessionGetFeatureForMessageFunction, err = sessionObject.InvokerNew("get_feature_for_message")
	})
	return err
}

// GetFeatureForMessage is a representation of the C type soup_session_get_feature_for_message.
func (recv *Session) GetFeatureForMessage(featureType int64, msg *Message) *SessionFeature {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt64(featureType)
	inArgs[2].SetPointer(msg.Native())

	var ret gi.Argument

	err := sessionGetFeatureForMessageFunction_Set()
	if err == nil {
		ret = sessionGetFeatureForMessageFunction.Invoke(inArgs[:], nil)
	}

	retGo := SessionFeatureNewFromNative(ret.Pointer())

	return retGo
}

var sessionGetFeaturesFunction *gi.Function
var sessionGetFeaturesFunction_Once sync.Once

func sessionGetFeaturesFunction_Set() error {
	var err error
	sessionGetFeaturesFunction_Once.Do(func() {
		err = sessionObject_Set()
		if err != nil {
			return
		}
		sessionGetFeaturesFunction, err = sessionObject.InvokerNew("get_features")
	})
	return err
}

// GetFeatures is a representation of the C type soup_session_get_features.
func (recv *Session) GetFeatures(featureType int64) *glib.SList {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt64(featureType)

	var ret gi.Argument

	err := sessionGetFeaturesFunction_Set()
	if err == nil {
		ret = sessionGetFeaturesFunction.Invoke(inArgs[:], nil)
	}

	retGo := glib.SListNewFromNative(ret.Pointer())

	return retGo
}

var sessionHasFeatureFunction *gi.Function
var sessionHasFeatureFunction_Once sync.Once

func sessionHasFeatureFunction_Set() error {
	var err error
	sessionHasFeatureFunction_Once.Do(func() {
		err = sessionObject_Set()
		if err != nil {
			return
		}
		sessionHasFeatureFunction, err = sessionObject.InvokerNew("has_feature")
	})
	return err
}

// HasFeature is a representation of the C type soup_session_has_feature.
func (recv *Session) HasFeature(featureType int64) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt64(featureType)

	var ret gi.Argument

	err := sessionHasFeatureFunction_Set()
	if err == nil {
		ret = sessionHasFeatureFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var sessionPauseMessageFunction *gi.Function
var sessionPauseMessageFunction_Once sync.Once

func sessionPauseMessageFunction_Set() error {
	var err error
	sessionPauseMessageFunction_Once.Do(func() {
		err = sessionObject_Set()
		if err != nil {
			return
		}
		sessionPauseMessageFunction, err = sessionObject.InvokerNew("pause_message")
	})
	return err
}

// PauseMessage is a representation of the C type soup_session_pause_message.
func (recv *Session) PauseMessage(msg *Message) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(msg.Native())

	err := sessionPauseMessageFunction_Set()
	if err == nil {
		sessionPauseMessageFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'soup_session_prefetch_dns' : parameter 'callback' of type 'AddressCallback' not supported

var sessionPrepareForUriFunction *gi.Function
var sessionPrepareForUriFunction_Once sync.Once

func sessionPrepareForUriFunction_Set() error {
	var err error
	sessionPrepareForUriFunction_Once.Do(func() {
		err = sessionObject_Set()
		if err != nil {
			return
		}
		sessionPrepareForUriFunction, err = sessionObject.InvokerNew("prepare_for_uri")
	})
	return err
}

// PrepareForUri is a representation of the C type soup_session_prepare_for_uri.
func (recv *Session) PrepareForUri(uri *URI) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(uri.Native())

	err := sessionPrepareForUriFunction_Set()
	if err == nil {
		sessionPrepareForUriFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'soup_session_queue_message' : parameter 'callback' of type 'SessionCallback' not supported

var sessionRedirectMessageFunction *gi.Function
var sessionRedirectMessageFunction_Once sync.Once

func sessionRedirectMessageFunction_Set() error {
	var err error
	sessionRedirectMessageFunction_Once.Do(func() {
		err = sessionObject_Set()
		if err != nil {
			return
		}
		sessionRedirectMessageFunction, err = sessionObject.InvokerNew("redirect_message")
	})
	return err
}

// RedirectMessage is a representation of the C type soup_session_redirect_message.
func (recv *Session) RedirectMessage(msg *Message) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(msg.Native())

	var ret gi.Argument

	err := sessionRedirectMessageFunction_Set()
	if err == nil {
		ret = sessionRedirectMessageFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var sessionRemoveFeatureFunction *gi.Function
var sessionRemoveFeatureFunction_Once sync.Once

func sessionRemoveFeatureFunction_Set() error {
	var err error
	sessionRemoveFeatureFunction_Once.Do(func() {
		err = sessionObject_Set()
		if err != nil {
			return
		}
		sessionRemoveFeatureFunction, err = sessionObject.InvokerNew("remove_feature")
	})
	return err
}

// RemoveFeature is a representation of the C type soup_session_remove_feature.
func (recv *Session) RemoveFeature(feature *SessionFeature) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(feature.Native())

	err := sessionRemoveFeatureFunction_Set()
	if err == nil {
		sessionRemoveFeatureFunction.Invoke(inArgs[:], nil)
	}

	return
}

var sessionRemoveFeatureByTypeFunction *gi.Function
var sessionRemoveFeatureByTypeFunction_Once sync.Once

func sessionRemoveFeatureByTypeFunction_Set() error {
	var err error
	sessionRemoveFeatureByTypeFunction_Once.Do(func() {
		err = sessionObject_Set()
		if err != nil {
			return
		}
		sessionRemoveFeatureByTypeFunction, err = sessionObject.InvokerNew("remove_feature_by_type")
	})
	return err
}

// RemoveFeatureByType is a representation of the C type soup_session_remove_feature_by_type.
func (recv *Session) RemoveFeatureByType(featureType int64) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt64(featureType)

	err := sessionRemoveFeatureByTypeFunction_Set()
	if err == nil {
		sessionRemoveFeatureByTypeFunction.Invoke(inArgs[:], nil)
	}

	return
}

var sessionRequestFunction *gi.Function
var sessionRequestFunction_Once sync.Once

func sessionRequestFunction_Set() error {
	var err error
	sessionRequestFunction_Once.Do(func() {
		err = sessionObject_Set()
		if err != nil {
			return
		}
		sessionRequestFunction, err = sessionObject.InvokerNew("request")
	})
	return err
}

// Request_ is a representation of the C type soup_session_request.
func (recv *Session) Request_(uriString string) *Request {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(uriString)

	var ret gi.Argument

	err := sessionRequestFunction_Set()
	if err == nil {
		ret = sessionRequestFunction.Invoke(inArgs[:], nil)
	}

	retGo := RequestNewFromNative(ret.Pointer())

	return retGo
}

var sessionRequestHttpFunction *gi.Function
var sessionRequestHttpFunction_Once sync.Once

func sessionRequestHttpFunction_Set() error {
	var err error
	sessionRequestHttpFunction_Once.Do(func() {
		err = sessionObject_Set()
		if err != nil {
			return
		}
		sessionRequestHttpFunction, err = sessionObject.InvokerNew("request_http")
	})
	return err
}

// RequestHttp is a representation of the C type soup_session_request_http.
func (recv *Session) RequestHttp(method string, uriString string) *RequestHTTP {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(method)
	inArgs[2].SetString(uriString)

	var ret gi.Argument

	err := sessionRequestHttpFunction_Set()
	if err == nil {
		ret = sessionRequestHttpFunction.Invoke(inArgs[:], nil)
	}

	retGo := RequestHTTPNewFromNative(ret.Pointer())

	return retGo
}

var sessionRequestHttpUriFunction *gi.Function
var sessionRequestHttpUriFunction_Once sync.Once

func sessionRequestHttpUriFunction_Set() error {
	var err error
	sessionRequestHttpUriFunction_Once.Do(func() {
		err = sessionObject_Set()
		if err != nil {
			return
		}
		sessionRequestHttpUriFunction, err = sessionObject.InvokerNew("request_http_uri")
	})
	return err
}

// RequestHttpUri is a representation of the C type soup_session_request_http_uri.
func (recv *Session) RequestHttpUri(method string, uri *URI) *RequestHTTP {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(method)
	inArgs[2].SetPointer(uri.Native())

	var ret gi.Argument

	err := sessionRequestHttpUriFunction_Set()
	if err == nil {
		ret = sessionRequestHttpUriFunction.Invoke(inArgs[:], nil)
	}

	retGo := RequestHTTPNewFromNative(ret.Pointer())

	return retGo
}

var sessionRequestUriFunction *gi.Function
var sessionRequestUriFunction_Once sync.Once

func sessionRequestUriFunction_Set() error {
	var err error
	sessionRequestUriFunction_Once.Do(func() {
		err = sessionObject_Set()
		if err != nil {
			return
		}
		sessionRequestUriFunction, err = sessionObject.InvokerNew("request_uri")
	})
	return err
}

// RequestUri is a representation of the C type soup_session_request_uri.
func (recv *Session) RequestUri(uri *URI) *Request {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(uri.Native())

	var ret gi.Argument

	err := sessionRequestUriFunction_Set()
	if err == nil {
		ret = sessionRequestUriFunction.Invoke(inArgs[:], nil)
	}

	retGo := RequestNewFromNative(ret.Pointer())

	return retGo
}

var sessionRequeueMessageFunction *gi.Function
var sessionRequeueMessageFunction_Once sync.Once

func sessionRequeueMessageFunction_Set() error {
	var err error
	sessionRequeueMessageFunction_Once.Do(func() {
		err = sessionObject_Set()
		if err != nil {
			return
		}
		sessionRequeueMessageFunction, err = sessionObject.InvokerNew("requeue_message")
	})
	return err
}

// RequeueMessage is a representation of the C type soup_session_requeue_message.
func (recv *Session) RequeueMessage(msg *Message) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(msg.Native())

	err := sessionRequeueMessageFunction_Set()
	if err == nil {
		sessionRequeueMessageFunction.Invoke(inArgs[:], nil)
	}

	return
}

var sessionSendFunction *gi.Function
var sessionSendFunction_Once sync.Once

func sessionSendFunction_Set() error {
	var err error
	sessionSendFunction_Once.Do(func() {
		err = sessionObject_Set()
		if err != nil {
			return
		}
		sessionSendFunction, err = sessionObject.InvokerNew("send")
	})
	return err
}

// Send is a representation of the C type soup_session_send.
func (recv *Session) Send(msg *Message, cancellable *gio.Cancellable) *gio.InputStream {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(msg.Native())
	inArgs[2].SetPointer(cancellable.Native())

	var ret gi.Argument

	err := sessionSendFunction_Set()
	if err == nil {
		ret = sessionSendFunction.Invoke(inArgs[:], nil)
	}

	retGo := gio.InputStreamNewFromNative(ret.Pointer())

	return retGo
}

// UNSUPPORTED : C value 'soup_session_send_async' : parameter 'callback' of type 'Gio.AsyncReadyCallback' not supported

var sessionSendFinishFunction *gi.Function
var sessionSendFinishFunction_Once sync.Once

func sessionSendFinishFunction_Set() error {
	var err error
	sessionSendFinishFunction_Once.Do(func() {
		err = sessionObject_Set()
		if err != nil {
			return
		}
		sessionSendFinishFunction, err = sessionObject.InvokerNew("send_finish")
	})
	return err
}

// SendFinish is a representation of the C type soup_session_send_finish.
func (recv *Session) SendFinish(result *gio.AsyncResult) *gio.InputStream {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(result.Native())

	var ret gi.Argument

	err := sessionSendFinishFunction_Set()
	if err == nil {
		ret = sessionSendFinishFunction.Invoke(inArgs[:], nil)
	}

	retGo := gio.InputStreamNewFromNative(ret.Pointer())

	return retGo
}

var sessionSendMessageFunction *gi.Function
var sessionSendMessageFunction_Once sync.Once

func sessionSendMessageFunction_Set() error {
	var err error
	sessionSendMessageFunction_Once.Do(func() {
		err = sessionObject_Set()
		if err != nil {
			return
		}
		sessionSendMessageFunction, err = sessionObject.InvokerNew("send_message")
	})
	return err
}

// SendMessage is a representation of the C type soup_session_send_message.
func (recv *Session) SendMessage(msg *Message) uint32 {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(msg.Native())

	var ret gi.Argument

	err := sessionSendMessageFunction_Set()
	if err == nil {
		ret = sessionSendMessageFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo
}

var sessionStealConnectionFunction *gi.Function
var sessionStealConnectionFunction_Once sync.Once

func sessionStealConnectionFunction_Set() error {
	var err error
	sessionStealConnectionFunction_Once.Do(func() {
		err = sessionObject_Set()
		if err != nil {
			return
		}
		sessionStealConnectionFunction, err = sessionObject.InvokerNew("steal_connection")
	})
	return err
}

// StealConnection is a representation of the C type soup_session_steal_connection.
func (recv *Session) StealConnection(msg *Message) *gio.IOStream {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(msg.Native())

	var ret gi.Argument

	err := sessionStealConnectionFunction_Set()
	if err == nil {
		ret = sessionStealConnectionFunction.Invoke(inArgs[:], nil)
	}

	retGo := gio.IOStreamNewFromNative(ret.Pointer())

	return retGo
}

var sessionUnpauseMessageFunction *gi.Function
var sessionUnpauseMessageFunction_Once sync.Once

func sessionUnpauseMessageFunction_Set() error {
	var err error
	sessionUnpauseMessageFunction_Once.Do(func() {
		err = sessionObject_Set()
		if err != nil {
			return
		}
		sessionUnpauseMessageFunction, err = sessionObject.InvokerNew("unpause_message")
	})
	return err
}

// UnpauseMessage is a representation of the C type soup_session_unpause_message.
func (recv *Session) UnpauseMessage(msg *Message) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(msg.Native())

	err := sessionUnpauseMessageFunction_Set()
	if err == nil {
		sessionUnpauseMessageFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'soup_session_websocket_connect_async' : parameter 'callback' of type 'Gio.AsyncReadyCallback' not supported

var sessionWebsocketConnectFinishFunction *gi.Function
var sessionWebsocketConnectFinishFunction_Once sync.Once

func sessionWebsocketConnectFinishFunction_Set() error {
	var err error
	sessionWebsocketConnectFinishFunction_Once.Do(func() {
		err = sessionObject_Set()
		if err != nil {
			return
		}
		sessionWebsocketConnectFinishFunction, err = sessionObject.InvokerNew("websocket_connect_finish")
	})
	return err
}

// WebsocketConnectFinish is a representation of the C type soup_session_websocket_connect_finish.
func (recv *Session) WebsocketConnectFinish(result *gio.AsyncResult) *WebsocketConnection {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(result.Native())

	var ret gi.Argument

	err := sessionWebsocketConnectFinishFunction_Set()
	if err == nil {
		ret = sessionWebsocketConnectFinishFunction.Invoke(inArgs[:], nil)
	}

	retGo := WebsocketConnectionNewFromNative(ret.Pointer())

	return retGo
}

var sessionWouldRedirectFunction *gi.Function
var sessionWouldRedirectFunction_Once sync.Once

func sessionWouldRedirectFunction_Set() error {
	var err error
	sessionWouldRedirectFunction_Once.Do(func() {
		err = sessionObject_Set()
		if err != nil {
			return
		}
		sessionWouldRedirectFunction, err = sessionObject.InvokerNew("would_redirect")
	})
	return err
}

// WouldRedirect is a representation of the C type soup_session_would_redirect.
func (recv *Session) WouldRedirect(msg *Message) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(msg.Native())

	var ret gi.Argument

	err := sessionWouldRedirectFunction_Set()
	if err == nil {
		ret = sessionWouldRedirectFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

/*
ConnectAuthenticate connects a callback to the 'authenticate' signal of the Session.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *Session) ConnectAuthenticate(handler func(instance *Session, msg *Message, auth *Auth, retrying bool)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {
		objectInstance := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[0]))
		argInstance := SessionNewFromNative(objectInstance.GetObject().Native())

		object1 := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[1]))
		arg1 := MessageNewFromNative(object1.GetObject().Native())

		object2 := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[2]))
		arg2 := AuthNewFromNative(object2.GetObject().Native())

		object3 := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[3]))
		arg3 := object3.GetBoolean()

		handler(argInstance, arg1, arg2, arg3)

	}

	return callback.ConnectSignal(recv.Native(), "authenticate", marshal)
}

/*
ConnectConnectionCreated connects a callback to the 'connection-created' signal of the Session.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *Session) ConnectConnectionCreated(handler func(instance *Session, connection *gobject.Object)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {
		objectInstance := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[0]))
		argInstance := SessionNewFromNative(objectInstance.GetObject().Native())

		object1 := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[1]))
		arg1 := gobject.ObjectNewFromNative(object1.GetObject().Native())

		handler(argInstance, arg1)

	}

	return callback.ConnectSignal(recv.Native(), "connection-created", marshal)
}

/*
ConnectRequestQueued connects a callback to the 'request-queued' signal of the Session.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *Session) ConnectRequestQueued(handler func(instance *Session, msg *Message)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {
		objectInstance := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[0]))
		argInstance := SessionNewFromNative(objectInstance.GetObject().Native())

		object1 := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[1]))
		arg1 := MessageNewFromNative(object1.GetObject().Native())

		handler(argInstance, arg1)

	}

	return callback.ConnectSignal(recv.Native(), "request-queued", marshal)
}

/*
ConnectRequestStarted connects a callback to the 'request-started' signal of the Session.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *Session) ConnectRequestStarted(handler func(instance *Session, msg *Message, socket *Socket)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {
		objectInstance := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[0]))
		argInstance := SessionNewFromNative(objectInstance.GetObject().Native())

		object1 := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[1]))
		arg1 := MessageNewFromNative(object1.GetObject().Native())

		object2 := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[2]))
		arg2 := SocketNewFromNative(object2.GetObject().Native())

		handler(argInstance, arg1, arg2)

	}

	return callback.ConnectSignal(recv.Native(), "request-started", marshal)
}

/*
ConnectRequestUnqueued connects a callback to the 'request-unqueued' signal of the Session.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *Session) ConnectRequestUnqueued(handler func(instance *Session, msg *Message)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {
		objectInstance := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[0]))
		argInstance := SessionNewFromNative(objectInstance.GetObject().Native())

		object1 := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[1]))
		arg1 := MessageNewFromNative(object1.GetObject().Native())

		handler(argInstance, arg1)

	}

	return callback.ConnectSignal(recv.Native(), "request-unqueued", marshal)
}

/*
ConnectTunneling connects a callback to the 'tunneling' signal of the Session.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *Session) ConnectTunneling(handler func(instance *Session, connection *gobject.Object)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {
		objectInstance := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[0]))
		argInstance := SessionNewFromNative(objectInstance.GetObject().Native())

		object1 := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[1]))
		arg1 := gobject.ObjectNewFromNative(object1.GetObject().Native())

		handler(argInstance, arg1)

	}

	return callback.ConnectSignal(recv.Native(), "tunneling", marshal)
}

/*
Disconnect disconnects a callback previously registered with a Connect...() method.

The connectionID should be a value returned from a call to a Connect...() method.
*/
func (recv *Session) DisconnectSignal(connectionID int) {
	callback.DisconnectSignal(connectionID)
}

var sessionAsyncObject *gi.Object
var sessionAsyncObject_Once sync.Once

func sessionAsyncObject_Set() error {
	var err error
	sessionAsyncObject_Once.Do(func() {
		sessionAsyncObject, err = gi.ObjectNew("Soup", "SessionAsync")
	})
	return err
}

type SessionAsync struct {
	native unsafe.Pointer
}

func SessionAsyncNewFromNative(native unsafe.Pointer) *SessionAsync {
	err := sessionAsyncObject_Set()
	if err != nil {
		return nil
	}

	instance := &SessionAsync{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Session upcasts to *Session
func (recv *SessionAsync) Session() *Session {
	return SessionNewFromNative(recv.Native())
}

// Object upcasts to *Object
func (recv *SessionAsync) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToSessionAsync down casts any arbitrary Object to SessionAsync.
Exercise care, as this is a potentially dangerous function
if the Object is not a SessionAsync.
*/
func CastToSessionAsync(object *gobject.Object) *SessionAsync {
	return SessionAsyncNewFromNative(object.Native())
}

// Equals compares this SessionAsync with another SessionAsync, and returns true if they represent the same Object.
func (recv *SessionAsync) Equals(other *SessionAsync) bool {
	return other.Native() == recv.Native()
}

func (recv *SessionAsync) Native() unsafe.Pointer {
	return recv.native
}

// FieldParent returns the C field 'parent'.
func (recv *SessionAsync) FieldParent() *Session {
	argValue := gi.ObjectFieldGet(sessionAsyncObject, recv.Native(), "parent")
	value := SessionNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParent sets the value of the C field 'parent'.
func (recv *SessionAsync) SetFieldParent(value *Session) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(sessionAsyncObject, recv.Native(), "parent", argValue)
}

var sessionAsyncNewFunction *gi.Function
var sessionAsyncNewFunction_Once sync.Once

func sessionAsyncNewFunction_Set() error {
	var err error
	sessionAsyncNewFunction_Once.Do(func() {
		err = sessionAsyncObject_Set()
		if err != nil {
			return
		}
		sessionAsyncNewFunction, err = sessionAsyncObject.InvokerNew("new")
	})
	return err
}

// SessionAsyncNew is a representation of the C type soup_session_async_new.
func SessionAsyncNew() *SessionAsync {

	var ret gi.Argument

	err := sessionAsyncNewFunction_Set()
	if err == nil {
		ret = sessionAsyncNewFunction.Invoke(nil, nil)
	}

	retGo := SessionAsyncNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

// UNSUPPORTED : C value 'soup_session_async_new_with_options' : parameter '...' of type 'nil' not supported

var sessionSyncObject *gi.Object
var sessionSyncObject_Once sync.Once

func sessionSyncObject_Set() error {
	var err error
	sessionSyncObject_Once.Do(func() {
		sessionSyncObject, err = gi.ObjectNew("Soup", "SessionSync")
	})
	return err
}

type SessionSync struct {
	native unsafe.Pointer
}

func SessionSyncNewFromNative(native unsafe.Pointer) *SessionSync {
	err := sessionSyncObject_Set()
	if err != nil {
		return nil
	}

	instance := &SessionSync{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Session upcasts to *Session
func (recv *SessionSync) Session() *Session {
	return SessionNewFromNative(recv.Native())
}

// Object upcasts to *Object
func (recv *SessionSync) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToSessionSync down casts any arbitrary Object to SessionSync.
Exercise care, as this is a potentially dangerous function
if the Object is not a SessionSync.
*/
func CastToSessionSync(object *gobject.Object) *SessionSync {
	return SessionSyncNewFromNative(object.Native())
}

// Equals compares this SessionSync with another SessionSync, and returns true if they represent the same Object.
func (recv *SessionSync) Equals(other *SessionSync) bool {
	return other.Native() == recv.Native()
}

func (recv *SessionSync) Native() unsafe.Pointer {
	return recv.native
}

// FieldParent returns the C field 'parent'.
func (recv *SessionSync) FieldParent() *Session {
	argValue := gi.ObjectFieldGet(sessionSyncObject, recv.Native(), "parent")
	value := SessionNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParent sets the value of the C field 'parent'.
func (recv *SessionSync) SetFieldParent(value *Session) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(sessionSyncObject, recv.Native(), "parent", argValue)
}

var sessionSyncNewFunction *gi.Function
var sessionSyncNewFunction_Once sync.Once

func sessionSyncNewFunction_Set() error {
	var err error
	sessionSyncNewFunction_Once.Do(func() {
		err = sessionSyncObject_Set()
		if err != nil {
			return
		}
		sessionSyncNewFunction, err = sessionSyncObject.InvokerNew("new")
	})
	return err
}

// SessionSyncNew is a representation of the C type soup_session_sync_new.
func SessionSyncNew() *SessionSync {

	var ret gi.Argument

	err := sessionSyncNewFunction_Set()
	if err == nil {
		ret = sessionSyncNewFunction.Invoke(nil, nil)
	}

	retGo := SessionSyncNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

// UNSUPPORTED : C value 'soup_session_sync_new_with_options' : parameter '...' of type 'nil' not supported

var socketObject *gi.Object
var socketObject_Once sync.Once

func socketObject_Set() error {
	var err error
	socketObject_Once.Do(func() {
		socketObject, err = gi.ObjectNew("Soup", "Socket")
	})
	return err
}

type Socket struct {
	native unsafe.Pointer
}

func SocketNewFromNative(native unsafe.Pointer) *Socket {
	err := socketObject_Set()
	if err != nil {
		return nil
	}

	instance := &Socket{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *Socket) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToSocket down casts any arbitrary Object to Socket.
Exercise care, as this is a potentially dangerous function
if the Object is not a Socket.
*/
func CastToSocket(object *gobject.Object) *Socket {
	return SocketNewFromNative(object.Native())
}

// Equals compares this Socket with another Socket, and returns true if they represent the same Object.
func (recv *Socket) Equals(other *Socket) bool {
	return other.Native() == recv.Native()
}

func (recv *Socket) Native() unsafe.Pointer {
	return recv.native
}

// FieldParent returns the C field 'parent'.
func (recv *Socket) FieldParent() *gobject.Object {
	argValue := gi.ObjectFieldGet(socketObject, recv.Native(), "parent")
	value := gobject.ObjectNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParent sets the value of the C field 'parent'.
func (recv *Socket) SetFieldParent(value *gobject.Object) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(socketObject, recv.Native(), "parent", argValue)
}

// UNSUPPORTED : C value 'soup_socket_new' : parameter '...' of type 'nil' not supported

// UNSUPPORTED : C value 'soup_socket_connect_async' : parameter 'callback' of type 'SocketCallback' not supported

var socketConnectSyncFunction *gi.Function
var socketConnectSyncFunction_Once sync.Once

func socketConnectSyncFunction_Set() error {
	var err error
	socketConnectSyncFunction_Once.Do(func() {
		err = socketObject_Set()
		if err != nil {
			return
		}
		socketConnectSyncFunction, err = socketObject.InvokerNew("connect_sync")
	})
	return err
}

// ConnectSync is a representation of the C type soup_socket_connect_sync.
func (recv *Socket) ConnectSync(cancellable *gio.Cancellable) uint32 {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(cancellable.Native())

	var ret gi.Argument

	err := socketConnectSyncFunction_Set()
	if err == nil {
		ret = socketConnectSyncFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo
}

var socketDisconnectFunction *gi.Function
var socketDisconnectFunction_Once sync.Once

func socketDisconnectFunction_Set() error {
	var err error
	socketDisconnectFunction_Once.Do(func() {
		err = socketObject_Set()
		if err != nil {
			return
		}
		socketDisconnectFunction, err = socketObject.InvokerNew("disconnect")
	})
	return err
}

// Disconnect is a representation of the C type soup_socket_disconnect.
func (recv *Socket) Disconnect() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := socketDisconnectFunction_Set()
	if err == nil {
		socketDisconnectFunction.Invoke(inArgs[:], nil)
	}

	return
}

var socketGetFdFunction *gi.Function
var socketGetFdFunction_Once sync.Once

func socketGetFdFunction_Set() error {
	var err error
	socketGetFdFunction_Once.Do(func() {
		err = socketObject_Set()
		if err != nil {
			return
		}
		socketGetFdFunction, err = socketObject.InvokerNew("get_fd")
	})
	return err
}

// GetFd is a representation of the C type soup_socket_get_fd.
func (recv *Socket) GetFd() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := socketGetFdFunction_Set()
	if err == nil {
		ret = socketGetFdFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var socketGetLocalAddressFunction *gi.Function
var socketGetLocalAddressFunction_Once sync.Once

func socketGetLocalAddressFunction_Set() error {
	var err error
	socketGetLocalAddressFunction_Once.Do(func() {
		err = socketObject_Set()
		if err != nil {
			return
		}
		socketGetLocalAddressFunction, err = socketObject.InvokerNew("get_local_address")
	})
	return err
}

// GetLocalAddress is a representation of the C type soup_socket_get_local_address.
func (recv *Socket) GetLocalAddress() *Address {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := socketGetLocalAddressFunction_Set()
	if err == nil {
		ret = socketGetLocalAddressFunction.Invoke(inArgs[:], nil)
	}

	retGo := AddressNewFromNative(ret.Pointer())

	return retGo
}

var socketGetRemoteAddressFunction *gi.Function
var socketGetRemoteAddressFunction_Once sync.Once

func socketGetRemoteAddressFunction_Set() error {
	var err error
	socketGetRemoteAddressFunction_Once.Do(func() {
		err = socketObject_Set()
		if err != nil {
			return
		}
		socketGetRemoteAddressFunction, err = socketObject.InvokerNew("get_remote_address")
	})
	return err
}

// GetRemoteAddress is a representation of the C type soup_socket_get_remote_address.
func (recv *Socket) GetRemoteAddress() *Address {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := socketGetRemoteAddressFunction_Set()
	if err == nil {
		ret = socketGetRemoteAddressFunction.Invoke(inArgs[:], nil)
	}

	retGo := AddressNewFromNative(ret.Pointer())

	return retGo
}

var socketIsConnectedFunction *gi.Function
var socketIsConnectedFunction_Once sync.Once

func socketIsConnectedFunction_Set() error {
	var err error
	socketIsConnectedFunction_Once.Do(func() {
		err = socketObject_Set()
		if err != nil {
			return
		}
		socketIsConnectedFunction, err = socketObject.InvokerNew("is_connected")
	})
	return err
}

// IsConnected is a representation of the C type soup_socket_is_connected.
func (recv *Socket) IsConnected() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := socketIsConnectedFunction_Set()
	if err == nil {
		ret = socketIsConnectedFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var socketIsSslFunction *gi.Function
var socketIsSslFunction_Once sync.Once

func socketIsSslFunction_Set() error {
	var err error
	socketIsSslFunction_Once.Do(func() {
		err = socketObject_Set()
		if err != nil {
			return
		}
		socketIsSslFunction, err = socketObject.InvokerNew("is_ssl")
	})
	return err
}

// IsSsl is a representation of the C type soup_socket_is_ssl.
func (recv *Socket) IsSsl() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := socketIsSslFunction_Set()
	if err == nil {
		ret = socketIsSslFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var socketListenFunction *gi.Function
var socketListenFunction_Once sync.Once

func socketListenFunction_Set() error {
	var err error
	socketListenFunction_Once.Do(func() {
		err = socketObject_Set()
		if err != nil {
			return
		}
		socketListenFunction, err = socketObject.InvokerNew("listen")
	})
	return err
}

// Listen is a representation of the C type soup_socket_listen.
func (recv *Socket) Listen() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := socketListenFunction_Set()
	if err == nil {
		ret = socketListenFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'soup_socket_read' : parameter 'buffer' of type 'nil' not supported

// UNSUPPORTED : C value 'soup_socket_read_until' : parameter 'buffer' of type 'nil' not supported

var socketStartProxySslFunction *gi.Function
var socketStartProxySslFunction_Once sync.Once

func socketStartProxySslFunction_Set() error {
	var err error
	socketStartProxySslFunction_Once.Do(func() {
		err = socketObject_Set()
		if err != nil {
			return
		}
		socketStartProxySslFunction, err = socketObject.InvokerNew("start_proxy_ssl")
	})
	return err
}

// StartProxySsl is a representation of the C type soup_socket_start_proxy_ssl.
func (recv *Socket) StartProxySsl(sslHost string, cancellable *gio.Cancellable) bool {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(sslHost)
	inArgs[2].SetPointer(cancellable.Native())

	var ret gi.Argument

	err := socketStartProxySslFunction_Set()
	if err == nil {
		ret = socketStartProxySslFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var socketStartSslFunction *gi.Function
var socketStartSslFunction_Once sync.Once

func socketStartSslFunction_Set() error {
	var err error
	socketStartSslFunction_Once.Do(func() {
		err = socketObject_Set()
		if err != nil {
			return
		}
		socketStartSslFunction, err = socketObject.InvokerNew("start_ssl")
	})
	return err
}

// StartSsl is a representation of the C type soup_socket_start_ssl.
func (recv *Socket) StartSsl(cancellable *gio.Cancellable) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(cancellable.Native())

	var ret gi.Argument

	err := socketStartSslFunction_Set()
	if err == nil {
		ret = socketStartSslFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'soup_socket_write' : parameter 'buffer' of type 'nil' not supported

/*
ConnectDisconnected connects a callback to the 'disconnected' signal of the Socket.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *Socket) ConnectDisconnected(handler func(instance *Socket)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {
		objectInstance := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[0]))
		argInstance := SocketNewFromNative(objectInstance.GetObject().Native())

		handler(argInstance)

	}

	return callback.ConnectSignal(recv.Native(), "disconnected", marshal)
}

// UNSUPPORTED : C value 'event' : parameter 'event' of type 'Gio.SocketClientEvent' not supported

/*
ConnectNewConnection connects a callback to the 'new-connection' signal of the Socket.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *Socket) ConnectNewConnection(handler func(instance *Socket, new *Socket)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {
		objectInstance := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[0]))
		argInstance := SocketNewFromNative(objectInstance.GetObject().Native())

		object1 := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[1]))
		arg1 := SocketNewFromNative(object1.GetObject().Native())

		handler(argInstance, arg1)

	}

	return callback.ConnectSignal(recv.Native(), "new-connection", marshal)
}

/*
ConnectReadable connects a callback to the 'readable' signal of the Socket.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *Socket) ConnectReadable(handler func(instance *Socket)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {
		objectInstance := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[0]))
		argInstance := SocketNewFromNative(objectInstance.GetObject().Native())

		handler(argInstance)

	}

	return callback.ConnectSignal(recv.Native(), "readable", marshal)
}

/*
ConnectWritable connects a callback to the 'writable' signal of the Socket.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *Socket) ConnectWritable(handler func(instance *Socket)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {
		objectInstance := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[0]))
		argInstance := SocketNewFromNative(objectInstance.GetObject().Native())

		handler(argInstance)

	}

	return callback.ConnectSignal(recv.Native(), "writable", marshal)
}

/*
Disconnect disconnects a callback previously registered with a Connect...() method.

The connectionID should be a value returned from a call to a Connect...() method.
*/
func (recv *Socket) DisconnectSignal(connectionID int) {
	callback.DisconnectSignal(connectionID)
}

// GioInitable returns the Gio.Initable interface implemented by Socket
func (recv *Socket) GioInitable() *gio.Initable {
	return gio.InitableNewFromNative(recv.Native())
}

var websocketConnectionObject *gi.Object
var websocketConnectionObject_Once sync.Once

func websocketConnectionObject_Set() error {
	var err error
	websocketConnectionObject_Once.Do(func() {
		websocketConnectionObject, err = gi.ObjectNew("Soup", "WebsocketConnection")
	})
	return err
}

type WebsocketConnection struct {
	native unsafe.Pointer
}

func WebsocketConnectionNewFromNative(native unsafe.Pointer) *WebsocketConnection {
	err := websocketConnectionObject_Set()
	if err != nil {
		return nil
	}

	instance := &WebsocketConnection{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *WebsocketConnection) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToWebsocketConnection down casts any arbitrary Object to WebsocketConnection.
Exercise care, as this is a potentially dangerous function
if the Object is not a WebsocketConnection.
*/
func CastToWebsocketConnection(object *gobject.Object) *WebsocketConnection {
	return WebsocketConnectionNewFromNative(object.Native())
}

// Equals compares this WebsocketConnection with another WebsocketConnection, and returns true if they represent the same Object.
func (recv *WebsocketConnection) Equals(other *WebsocketConnection) bool {
	return other.Native() == recv.Native()
}

func (recv *WebsocketConnection) Native() unsafe.Pointer {
	return recv.native
}

// FieldParent returns the C field 'parent'.
func (recv *WebsocketConnection) FieldParent() *gobject.Object {
	argValue := gi.ObjectFieldGet(websocketConnectionObject, recv.Native(), "parent")
	value := gobject.ObjectNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParent sets the value of the C field 'parent'.
func (recv *WebsocketConnection) SetFieldParent(value *gobject.Object) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(websocketConnectionObject, recv.Native(), "parent", argValue)
}

var websocketConnectionNewFunction *gi.Function
var websocketConnectionNewFunction_Once sync.Once

func websocketConnectionNewFunction_Set() error {
	var err error
	websocketConnectionNewFunction_Once.Do(func() {
		err = websocketConnectionObject_Set()
		if err != nil {
			return
		}
		websocketConnectionNewFunction, err = websocketConnectionObject.InvokerNew("new")
	})
	return err
}

// WebsocketConnectionNew is a representation of the C type soup_websocket_connection_new.
func WebsocketConnectionNew(stream *gio.IOStream, uri *URI, type_ WebsocketConnectionType, origin string, protocol string) *WebsocketConnection {
	var inArgs [5]gi.Argument
	inArgs[0].SetPointer(stream.Native())
	inArgs[1].SetPointer(uri.Native())
	inArgs[2].SetInt32(int32(type_))
	inArgs[3].SetString(origin)
	inArgs[4].SetString(protocol)

	var ret gi.Argument

	err := websocketConnectionNewFunction_Set()
	if err == nil {
		ret = websocketConnectionNewFunction.Invoke(inArgs[:], nil)
	}

	retGo := WebsocketConnectionNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

var websocketConnectionNewWithExtensionsFunction *gi.Function
var websocketConnectionNewWithExtensionsFunction_Once sync.Once

func websocketConnectionNewWithExtensionsFunction_Set() error {
	var err error
	websocketConnectionNewWithExtensionsFunction_Once.Do(func() {
		err = websocketConnectionObject_Set()
		if err != nil {
			return
		}
		websocketConnectionNewWithExtensionsFunction, err = websocketConnectionObject.InvokerNew("new_with_extensions")
	})
	return err
}

// WebsocketConnectionNewWithExtensions is a representation of the C type soup_websocket_connection_new_with_extensions.
func WebsocketConnectionNewWithExtensions(stream *gio.IOStream, uri *URI, type_ WebsocketConnectionType, origin string, protocol string, extensions *glib.List) *WebsocketConnection {
	var inArgs [6]gi.Argument
	inArgs[0].SetPointer(stream.Native())
	inArgs[1].SetPointer(uri.Native())
	inArgs[2].SetInt32(int32(type_))
	inArgs[3].SetString(origin)
	inArgs[4].SetString(protocol)
	inArgs[5].SetPointer(extensions.Native())

	var ret gi.Argument

	err := websocketConnectionNewWithExtensionsFunction_Set()
	if err == nil {
		ret = websocketConnectionNewWithExtensionsFunction.Invoke(inArgs[:], nil)
	}

	retGo := WebsocketConnectionNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

var websocketConnectionCloseFunction *gi.Function
var websocketConnectionCloseFunction_Once sync.Once

func websocketConnectionCloseFunction_Set() error {
	var err error
	websocketConnectionCloseFunction_Once.Do(func() {
		err = websocketConnectionObject_Set()
		if err != nil {
			return
		}
		websocketConnectionCloseFunction, err = websocketConnectionObject.InvokerNew("close")
	})
	return err
}

// Close is a representation of the C type soup_websocket_connection_close.
func (recv *WebsocketConnection) Close(code uint16, data string) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetUint16(code)
	inArgs[2].SetString(data)

	err := websocketConnectionCloseFunction_Set()
	if err == nil {
		websocketConnectionCloseFunction.Invoke(inArgs[:], nil)
	}

	return
}

var websocketConnectionGetCloseCodeFunction *gi.Function
var websocketConnectionGetCloseCodeFunction_Once sync.Once

func websocketConnectionGetCloseCodeFunction_Set() error {
	var err error
	websocketConnectionGetCloseCodeFunction_Once.Do(func() {
		err = websocketConnectionObject_Set()
		if err != nil {
			return
		}
		websocketConnectionGetCloseCodeFunction, err = websocketConnectionObject.InvokerNew("get_close_code")
	})
	return err
}

// GetCloseCode is a representation of the C type soup_websocket_connection_get_close_code.
func (recv *WebsocketConnection) GetCloseCode() uint16 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := websocketConnectionGetCloseCodeFunction_Set()
	if err == nil {
		ret = websocketConnectionGetCloseCodeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint16()

	return retGo
}

var websocketConnectionGetCloseDataFunction *gi.Function
var websocketConnectionGetCloseDataFunction_Once sync.Once

func websocketConnectionGetCloseDataFunction_Set() error {
	var err error
	websocketConnectionGetCloseDataFunction_Once.Do(func() {
		err = websocketConnectionObject_Set()
		if err != nil {
			return
		}
		websocketConnectionGetCloseDataFunction, err = websocketConnectionObject.InvokerNew("get_close_data")
	})
	return err
}

// GetCloseData is a representation of the C type soup_websocket_connection_get_close_data.
func (recv *WebsocketConnection) GetCloseData() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := websocketConnectionGetCloseDataFunction_Set()
	if err == nil {
		ret = websocketConnectionGetCloseDataFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var websocketConnectionGetConnectionTypeFunction *gi.Function
var websocketConnectionGetConnectionTypeFunction_Once sync.Once

func websocketConnectionGetConnectionTypeFunction_Set() error {
	var err error
	websocketConnectionGetConnectionTypeFunction_Once.Do(func() {
		err = websocketConnectionObject_Set()
		if err != nil {
			return
		}
		websocketConnectionGetConnectionTypeFunction, err = websocketConnectionObject.InvokerNew("get_connection_type")
	})
	return err
}

// GetConnectionType is a representation of the C type soup_websocket_connection_get_connection_type.
func (recv *WebsocketConnection) GetConnectionType() WebsocketConnectionType {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := websocketConnectionGetConnectionTypeFunction_Set()
	if err == nil {
		ret = websocketConnectionGetConnectionTypeFunction.Invoke(inArgs[:], nil)
	}

	retGo := WebsocketConnectionType(ret.Int32())

	return retGo
}

var websocketConnectionGetExtensionsFunction *gi.Function
var websocketConnectionGetExtensionsFunction_Once sync.Once

func websocketConnectionGetExtensionsFunction_Set() error {
	var err error
	websocketConnectionGetExtensionsFunction_Once.Do(func() {
		err = websocketConnectionObject_Set()
		if err != nil {
			return
		}
		websocketConnectionGetExtensionsFunction, err = websocketConnectionObject.InvokerNew("get_extensions")
	})
	return err
}

// GetExtensions is a representation of the C type soup_websocket_connection_get_extensions.
func (recv *WebsocketConnection) GetExtensions() *glib.List {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := websocketConnectionGetExtensionsFunction_Set()
	if err == nil {
		ret = websocketConnectionGetExtensionsFunction.Invoke(inArgs[:], nil)
	}

	retGo := glib.ListNewFromNative(ret.Pointer())

	return retGo
}

var websocketConnectionGetIoStreamFunction *gi.Function
var websocketConnectionGetIoStreamFunction_Once sync.Once

func websocketConnectionGetIoStreamFunction_Set() error {
	var err error
	websocketConnectionGetIoStreamFunction_Once.Do(func() {
		err = websocketConnectionObject_Set()
		if err != nil {
			return
		}
		websocketConnectionGetIoStreamFunction, err = websocketConnectionObject.InvokerNew("get_io_stream")
	})
	return err
}

// GetIoStream is a representation of the C type soup_websocket_connection_get_io_stream.
func (recv *WebsocketConnection) GetIoStream() *gio.IOStream {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := websocketConnectionGetIoStreamFunction_Set()
	if err == nil {
		ret = websocketConnectionGetIoStreamFunction.Invoke(inArgs[:], nil)
	}

	retGo := gio.IOStreamNewFromNative(ret.Pointer())

	return retGo
}

var websocketConnectionGetKeepaliveIntervalFunction *gi.Function
var websocketConnectionGetKeepaliveIntervalFunction_Once sync.Once

func websocketConnectionGetKeepaliveIntervalFunction_Set() error {
	var err error
	websocketConnectionGetKeepaliveIntervalFunction_Once.Do(func() {
		err = websocketConnectionObject_Set()
		if err != nil {
			return
		}
		websocketConnectionGetKeepaliveIntervalFunction, err = websocketConnectionObject.InvokerNew("get_keepalive_interval")
	})
	return err
}

// GetKeepaliveInterval is a representation of the C type soup_websocket_connection_get_keepalive_interval.
func (recv *WebsocketConnection) GetKeepaliveInterval() uint32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := websocketConnectionGetKeepaliveIntervalFunction_Set()
	if err == nil {
		ret = websocketConnectionGetKeepaliveIntervalFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo
}

var websocketConnectionGetMaxIncomingPayloadSizeFunction *gi.Function
var websocketConnectionGetMaxIncomingPayloadSizeFunction_Once sync.Once

func websocketConnectionGetMaxIncomingPayloadSizeFunction_Set() error {
	var err error
	websocketConnectionGetMaxIncomingPayloadSizeFunction_Once.Do(func() {
		err = websocketConnectionObject_Set()
		if err != nil {
			return
		}
		websocketConnectionGetMaxIncomingPayloadSizeFunction, err = websocketConnectionObject.InvokerNew("get_max_incoming_payload_size")
	})
	return err
}

// GetMaxIncomingPayloadSize is a representation of the C type soup_websocket_connection_get_max_incoming_payload_size.
func (recv *WebsocketConnection) GetMaxIncomingPayloadSize() uint64 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := websocketConnectionGetMaxIncomingPayloadSizeFunction_Set()
	if err == nil {
		ret = websocketConnectionGetMaxIncomingPayloadSizeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint64()

	return retGo
}

var websocketConnectionGetOriginFunction *gi.Function
var websocketConnectionGetOriginFunction_Once sync.Once

func websocketConnectionGetOriginFunction_Set() error {
	var err error
	websocketConnectionGetOriginFunction_Once.Do(func() {
		err = websocketConnectionObject_Set()
		if err != nil {
			return
		}
		websocketConnectionGetOriginFunction, err = websocketConnectionObject.InvokerNew("get_origin")
	})
	return err
}

// GetOrigin is a representation of the C type soup_websocket_connection_get_origin.
func (recv *WebsocketConnection) GetOrigin() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := websocketConnectionGetOriginFunction_Set()
	if err == nil {
		ret = websocketConnectionGetOriginFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var websocketConnectionGetProtocolFunction *gi.Function
var websocketConnectionGetProtocolFunction_Once sync.Once

func websocketConnectionGetProtocolFunction_Set() error {
	var err error
	websocketConnectionGetProtocolFunction_Once.Do(func() {
		err = websocketConnectionObject_Set()
		if err != nil {
			return
		}
		websocketConnectionGetProtocolFunction, err = websocketConnectionObject.InvokerNew("get_protocol")
	})
	return err
}

// GetProtocol is a representation of the C type soup_websocket_connection_get_protocol.
func (recv *WebsocketConnection) GetProtocol() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := websocketConnectionGetProtocolFunction_Set()
	if err == nil {
		ret = websocketConnectionGetProtocolFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var websocketConnectionGetStateFunction *gi.Function
var websocketConnectionGetStateFunction_Once sync.Once

func websocketConnectionGetStateFunction_Set() error {
	var err error
	websocketConnectionGetStateFunction_Once.Do(func() {
		err = websocketConnectionObject_Set()
		if err != nil {
			return
		}
		websocketConnectionGetStateFunction, err = websocketConnectionObject.InvokerNew("get_state")
	})
	return err
}

// GetState is a representation of the C type soup_websocket_connection_get_state.
func (recv *WebsocketConnection) GetState() WebsocketState {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := websocketConnectionGetStateFunction_Set()
	if err == nil {
		ret = websocketConnectionGetStateFunction.Invoke(inArgs[:], nil)
	}

	retGo := WebsocketState(ret.Int32())

	return retGo
}

var websocketConnectionGetUriFunction *gi.Function
var websocketConnectionGetUriFunction_Once sync.Once

func websocketConnectionGetUriFunction_Set() error {
	var err error
	websocketConnectionGetUriFunction_Once.Do(func() {
		err = websocketConnectionObject_Set()
		if err != nil {
			return
		}
		websocketConnectionGetUriFunction, err = websocketConnectionObject.InvokerNew("get_uri")
	})
	return err
}

// GetUri is a representation of the C type soup_websocket_connection_get_uri.
func (recv *WebsocketConnection) GetUri() *URI {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := websocketConnectionGetUriFunction_Set()
	if err == nil {
		ret = websocketConnectionGetUriFunction.Invoke(inArgs[:], nil)
	}

	retGo := URINewFromNative(ret.Pointer())

	return retGo
}

// UNSUPPORTED : C value 'soup_websocket_connection_send_binary' : parameter 'data' of type 'nil' not supported

var websocketConnectionSendMessageFunction *gi.Function
var websocketConnectionSendMessageFunction_Once sync.Once

func websocketConnectionSendMessageFunction_Set() error {
	var err error
	websocketConnectionSendMessageFunction_Once.Do(func() {
		err = websocketConnectionObject_Set()
		if err != nil {
			return
		}
		websocketConnectionSendMessageFunction, err = websocketConnectionObject.InvokerNew("send_message")
	})
	return err
}

// SendMessage is a representation of the C type soup_websocket_connection_send_message.
func (recv *WebsocketConnection) SendMessage(type_ WebsocketDataType, message *glib.Bytes) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(int32(type_))
	inArgs[2].SetPointer(message.Native())

	err := websocketConnectionSendMessageFunction_Set()
	if err == nil {
		websocketConnectionSendMessageFunction.Invoke(inArgs[:], nil)
	}

	return
}

var websocketConnectionSendTextFunction *gi.Function
var websocketConnectionSendTextFunction_Once sync.Once

func websocketConnectionSendTextFunction_Set() error {
	var err error
	websocketConnectionSendTextFunction_Once.Do(func() {
		err = websocketConnectionObject_Set()
		if err != nil {
			return
		}
		websocketConnectionSendTextFunction, err = websocketConnectionObject.InvokerNew("send_text")
	})
	return err
}

// SendText is a representation of the C type soup_websocket_connection_send_text.
func (recv *WebsocketConnection) SendText(text string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(text)

	err := websocketConnectionSendTextFunction_Set()
	if err == nil {
		websocketConnectionSendTextFunction.Invoke(inArgs[:], nil)
	}

	return
}

var websocketConnectionSetKeepaliveIntervalFunction *gi.Function
var websocketConnectionSetKeepaliveIntervalFunction_Once sync.Once

func websocketConnectionSetKeepaliveIntervalFunction_Set() error {
	var err error
	websocketConnectionSetKeepaliveIntervalFunction_Once.Do(func() {
		err = websocketConnectionObject_Set()
		if err != nil {
			return
		}
		websocketConnectionSetKeepaliveIntervalFunction, err = websocketConnectionObject.InvokerNew("set_keepalive_interval")
	})
	return err
}

// SetKeepaliveInterval is a representation of the C type soup_websocket_connection_set_keepalive_interval.
func (recv *WebsocketConnection) SetKeepaliveInterval(interval uint32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetUint32(interval)

	err := websocketConnectionSetKeepaliveIntervalFunction_Set()
	if err == nil {
		websocketConnectionSetKeepaliveIntervalFunction.Invoke(inArgs[:], nil)
	}

	return
}

var websocketConnectionSetMaxIncomingPayloadSizeFunction *gi.Function
var websocketConnectionSetMaxIncomingPayloadSizeFunction_Once sync.Once

func websocketConnectionSetMaxIncomingPayloadSizeFunction_Set() error {
	var err error
	websocketConnectionSetMaxIncomingPayloadSizeFunction_Once.Do(func() {
		err = websocketConnectionObject_Set()
		if err != nil {
			return
		}
		websocketConnectionSetMaxIncomingPayloadSizeFunction, err = websocketConnectionObject.InvokerNew("set_max_incoming_payload_size")
	})
	return err
}

// SetMaxIncomingPayloadSize is a representation of the C type soup_websocket_connection_set_max_incoming_payload_size.
func (recv *WebsocketConnection) SetMaxIncomingPayloadSize(maxIncomingPayloadSize uint64) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetUint64(maxIncomingPayloadSize)

	err := websocketConnectionSetMaxIncomingPayloadSizeFunction_Set()
	if err == nil {
		websocketConnectionSetMaxIncomingPayloadSizeFunction.Invoke(inArgs[:], nil)
	}

	return
}

/*
ConnectClosed connects a callback to the 'closed' signal of the WebsocketConnection.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *WebsocketConnection) ConnectClosed(handler func(instance *WebsocketConnection)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {
		objectInstance := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[0]))
		argInstance := WebsocketConnectionNewFromNative(objectInstance.GetObject().Native())

		handler(argInstance)

	}

	return callback.ConnectSignal(recv.Native(), "closed", marshal)
}

/*
ConnectClosing connects a callback to the 'closing' signal of the WebsocketConnection.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *WebsocketConnection) ConnectClosing(handler func(instance *WebsocketConnection)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {
		objectInstance := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[0]))
		argInstance := WebsocketConnectionNewFromNative(objectInstance.GetObject().Native())

		handler(argInstance)

	}

	return callback.ConnectSignal(recv.Native(), "closing", marshal)
}

/*
ConnectError connects a callback to the 'error' signal of the WebsocketConnection.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *WebsocketConnection) ConnectError(handler func(instance *WebsocketConnection, error *glib.Error)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {
		objectInstance := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[0]))
		argInstance := WebsocketConnectionNewFromNative(objectInstance.GetObject().Native())

		object1 := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[1]))
		arg1 := glib.ErrorNewFromNative(object1.GetBoxed())

		handler(argInstance, arg1)

	}

	return callback.ConnectSignal(recv.Native(), "error", marshal)
}

/*
ConnectMessage connects a callback to the 'message' signal of the WebsocketConnection.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *WebsocketConnection) ConnectMessage(handler func(instance *WebsocketConnection, type_ int32, message *glib.Bytes)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {
		objectInstance := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[0]))
		argInstance := WebsocketConnectionNewFromNative(objectInstance.GetObject().Native())

		object1 := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[1]))
		arg1 := object1.GetInt()

		object2 := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[2]))
		arg2 := glib.BytesNewFromNative(object2.GetBoxed())

		handler(argInstance, arg1, arg2)

	}

	return callback.ConnectSignal(recv.Native(), "message", marshal)
}

/*
ConnectPong connects a callback to the 'pong' signal of the WebsocketConnection.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *WebsocketConnection) ConnectPong(handler func(instance *WebsocketConnection, message *glib.Bytes)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {
		objectInstance := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[0]))
		argInstance := WebsocketConnectionNewFromNative(objectInstance.GetObject().Native())

		object1 := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[1]))
		arg1 := glib.BytesNewFromNative(object1.GetBoxed())

		handler(argInstance, arg1)

	}

	return callback.ConnectSignal(recv.Native(), "pong", marshal)
}

/*
Disconnect disconnects a callback previously registered with a Connect...() method.

The connectionID should be a value returned from a call to a Connect...() method.
*/
func (recv *WebsocketConnection) DisconnectSignal(connectionID int) {
	callback.DisconnectSignal(connectionID)
}

var websocketExtensionObject *gi.Object
var websocketExtensionObject_Once sync.Once

func websocketExtensionObject_Set() error {
	var err error
	websocketExtensionObject_Once.Do(func() {
		websocketExtensionObject, err = gi.ObjectNew("Soup", "WebsocketExtension")
	})
	return err
}

type WebsocketExtension struct {
	native unsafe.Pointer
}

func WebsocketExtensionNewFromNative(native unsafe.Pointer) *WebsocketExtension {
	err := websocketExtensionObject_Set()
	if err != nil {
		return nil
	}

	instance := &WebsocketExtension{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *WebsocketExtension) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToWebsocketExtension down casts any arbitrary Object to WebsocketExtension.
Exercise care, as this is a potentially dangerous function
if the Object is not a WebsocketExtension.
*/
func CastToWebsocketExtension(object *gobject.Object) *WebsocketExtension {
	return WebsocketExtensionNewFromNative(object.Native())
}

// Equals compares this WebsocketExtension with another WebsocketExtension, and returns true if they represent the same Object.
func (recv *WebsocketExtension) Equals(other *WebsocketExtension) bool {
	return other.Native() == recv.Native()
}

func (recv *WebsocketExtension) Native() unsafe.Pointer {
	return recv.native
}

// FieldParent returns the C field 'parent'.
func (recv *WebsocketExtension) FieldParent() *gobject.Object {
	argValue := gi.ObjectFieldGet(websocketExtensionObject, recv.Native(), "parent")
	value := gobject.ObjectNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParent sets the value of the C field 'parent'.
func (recv *WebsocketExtension) SetFieldParent(value *gobject.Object) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(websocketExtensionObject, recv.Native(), "parent", argValue)
}

var websocketExtensionConfigureFunction *gi.Function
var websocketExtensionConfigureFunction_Once sync.Once

func websocketExtensionConfigureFunction_Set() error {
	var err error
	websocketExtensionConfigureFunction_Once.Do(func() {
		err = websocketExtensionObject_Set()
		if err != nil {
			return
		}
		websocketExtensionConfigureFunction, err = websocketExtensionObject.InvokerNew("configure")
	})
	return err
}

// Configure is a representation of the C type soup_websocket_extension_configure.
func (recv *WebsocketExtension) Configure(connectionType WebsocketConnectionType, params *glib.HashTable) bool {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(int32(connectionType))
	inArgs[2].SetPointer(params.Native())

	var ret gi.Argument

	err := websocketExtensionConfigureFunction_Set()
	if err == nil {
		ret = websocketExtensionConfigureFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var websocketExtensionGetRequestParamsFunction *gi.Function
var websocketExtensionGetRequestParamsFunction_Once sync.Once

func websocketExtensionGetRequestParamsFunction_Set() error {
	var err error
	websocketExtensionGetRequestParamsFunction_Once.Do(func() {
		err = websocketExtensionObject_Set()
		if err != nil {
			return
		}
		websocketExtensionGetRequestParamsFunction, err = websocketExtensionObject.InvokerNew("get_request_params")
	})
	return err
}

// GetRequestParams is a representation of the C type soup_websocket_extension_get_request_params.
func (recv *WebsocketExtension) GetRequestParams() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := websocketExtensionGetRequestParamsFunction_Set()
	if err == nil {
		ret = websocketExtensionGetRequestParamsFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

var websocketExtensionGetResponseParamsFunction *gi.Function
var websocketExtensionGetResponseParamsFunction_Once sync.Once

func websocketExtensionGetResponseParamsFunction_Set() error {
	var err error
	websocketExtensionGetResponseParamsFunction_Once.Do(func() {
		err = websocketExtensionObject_Set()
		if err != nil {
			return
		}
		websocketExtensionGetResponseParamsFunction, err = websocketExtensionObject.InvokerNew("get_response_params")
	})
	return err
}

// GetResponseParams is a representation of the C type soup_websocket_extension_get_response_params.
func (recv *WebsocketExtension) GetResponseParams() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := websocketExtensionGetResponseParamsFunction_Set()
	if err == nil {
		ret = websocketExtensionGetResponseParamsFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

var websocketExtensionProcessIncomingMessageFunction *gi.Function
var websocketExtensionProcessIncomingMessageFunction_Once sync.Once

func websocketExtensionProcessIncomingMessageFunction_Set() error {
	var err error
	websocketExtensionProcessIncomingMessageFunction_Once.Do(func() {
		err = websocketExtensionObject_Set()
		if err != nil {
			return
		}
		websocketExtensionProcessIncomingMessageFunction, err = websocketExtensionObject.InvokerNew("process_incoming_message")
	})
	return err
}

// ProcessIncomingMessage is a representation of the C type soup_websocket_extension_process_incoming_message.
func (recv *WebsocketExtension) ProcessIncomingMessage(header uint8, payload *glib.Bytes) (*glib.Bytes, uint8) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetUint8(header)
	inArgs[2].SetPointer(payload.Native())

	var outArgs [1]gi.Argument
	var ret gi.Argument

	err := websocketExtensionProcessIncomingMessageFunction_Set()
	if err == nil {
		ret = websocketExtensionProcessIncomingMessageFunction.Invoke(inArgs[:], outArgs[:])
	}

	retGo := glib.BytesNewFromNative(ret.Pointer())
	out0 := outArgs[0].Uint8()

	return retGo, out0
}

var websocketExtensionProcessOutgoingMessageFunction *gi.Function
var websocketExtensionProcessOutgoingMessageFunction_Once sync.Once

func websocketExtensionProcessOutgoingMessageFunction_Set() error {
	var err error
	websocketExtensionProcessOutgoingMessageFunction_Once.Do(func() {
		err = websocketExtensionObject_Set()
		if err != nil {
			return
		}
		websocketExtensionProcessOutgoingMessageFunction, err = websocketExtensionObject.InvokerNew("process_outgoing_message")
	})
	return err
}

// ProcessOutgoingMessage is a representation of the C type soup_websocket_extension_process_outgoing_message.
func (recv *WebsocketExtension) ProcessOutgoingMessage(header uint8, payload *glib.Bytes) (*glib.Bytes, uint8) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetUint8(header)
	inArgs[2].SetPointer(payload.Native())

	var outArgs [1]gi.Argument
	var ret gi.Argument

	err := websocketExtensionProcessOutgoingMessageFunction_Set()
	if err == nil {
		ret = websocketExtensionProcessOutgoingMessageFunction.Invoke(inArgs[:], outArgs[:])
	}

	retGo := glib.BytesNewFromNative(ret.Pointer())
	out0 := outArgs[0].Uint8()

	return retGo, out0
}

var websocketExtensionDeflateObject *gi.Object
var websocketExtensionDeflateObject_Once sync.Once

func websocketExtensionDeflateObject_Set() error {
	var err error
	websocketExtensionDeflateObject_Once.Do(func() {
		websocketExtensionDeflateObject, err = gi.ObjectNew("Soup", "WebsocketExtensionDeflate")
	})
	return err
}

type WebsocketExtensionDeflate struct {
	native unsafe.Pointer
}

func WebsocketExtensionDeflateNewFromNative(native unsafe.Pointer) *WebsocketExtensionDeflate {
	err := websocketExtensionDeflateObject_Set()
	if err != nil {
		return nil
	}

	instance := &WebsocketExtensionDeflate{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// WebsocketExtension upcasts to *WebsocketExtension
func (recv *WebsocketExtensionDeflate) WebsocketExtension() *WebsocketExtension {
	return WebsocketExtensionNewFromNative(recv.Native())
}

// Object upcasts to *Object
func (recv *WebsocketExtensionDeflate) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToWebsocketExtensionDeflate down casts any arbitrary Object to WebsocketExtensionDeflate.
Exercise care, as this is a potentially dangerous function
if the Object is not a WebsocketExtensionDeflate.
*/
func CastToWebsocketExtensionDeflate(object *gobject.Object) *WebsocketExtensionDeflate {
	return WebsocketExtensionDeflateNewFromNative(object.Native())
}

// Equals compares this WebsocketExtensionDeflate with another WebsocketExtensionDeflate, and returns true if they represent the same Object.
func (recv *WebsocketExtensionDeflate) Equals(other *WebsocketExtensionDeflate) bool {
	return other.Native() == recv.Native()
}

func (recv *WebsocketExtensionDeflate) Native() unsafe.Pointer {
	return recv.native
}

// FieldParent returns the C field 'parent'.
func (recv *WebsocketExtensionDeflate) FieldParent() *WebsocketExtension {
	argValue := gi.ObjectFieldGet(websocketExtensionDeflateObject, recv.Native(), "parent")
	value := WebsocketExtensionNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParent sets the value of the C field 'parent'.
func (recv *WebsocketExtensionDeflate) SetFieldParent(value *WebsocketExtension) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(websocketExtensionDeflateObject, recv.Native(), "parent", argValue)
}

var websocketExtensionManagerObject *gi.Object
var websocketExtensionManagerObject_Once sync.Once

func websocketExtensionManagerObject_Set() error {
	var err error
	websocketExtensionManagerObject_Once.Do(func() {
		websocketExtensionManagerObject, err = gi.ObjectNew("Soup", "WebsocketExtensionManager")
	})
	return err
}

type WebsocketExtensionManager struct {
	native unsafe.Pointer
}

func WebsocketExtensionManagerNewFromNative(native unsafe.Pointer) *WebsocketExtensionManager {
	err := websocketExtensionManagerObject_Set()
	if err != nil {
		return nil
	}

	instance := &WebsocketExtensionManager{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *WebsocketExtensionManager) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToWebsocketExtensionManager down casts any arbitrary Object to WebsocketExtensionManager.
Exercise care, as this is a potentially dangerous function
if the Object is not a WebsocketExtensionManager.
*/
func CastToWebsocketExtensionManager(object *gobject.Object) *WebsocketExtensionManager {
	return WebsocketExtensionManagerNewFromNative(object.Native())
}

// Equals compares this WebsocketExtensionManager with another WebsocketExtensionManager, and returns true if they represent the same Object.
func (recv *WebsocketExtensionManager) Equals(other *WebsocketExtensionManager) bool {
	return other.Native() == recv.Native()
}

func (recv *WebsocketExtensionManager) Native() unsafe.Pointer {
	return recv.native
}

// FieldParent returns the C field 'parent'.
func (recv *WebsocketExtensionManager) FieldParent() *gobject.Object {
	argValue := gi.ObjectFieldGet(websocketExtensionManagerObject, recv.Native(), "parent")
	value := gobject.ObjectNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParent sets the value of the C field 'parent'.
func (recv *WebsocketExtensionManager) SetFieldParent(value *gobject.Object) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(websocketExtensionManagerObject, recv.Native(), "parent", argValue)
}

// SessionFeature returns the SessionFeature interface implemented by WebsocketExtensionManager
func (recv *WebsocketExtensionManager) SessionFeature() *SessionFeature {
	return SessionFeatureNewFromNative(recv.Native())
}