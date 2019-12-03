// Code generated - DO NOT EDIT.

package soup

import (
	gi "github.com/pekim/gobbi/internal/gi"
	gio "github.com/pekim/gobbi/lib/gio"
	gobject "github.com/pekim/gobbi/lib/gobject"
	"sync"
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
	gobject.Object
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'GObject.Object'

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

	retGo := &Address{}
	retGo.Native = ret.Pointer()

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

	retGo := &Address{}
	retGo.Native = ret.Pointer()

	return retGo
}

// UNSUPPORTED : C value 'soup_address_new_from_sockaddr' : parameter 'sa' of type 'gpointer' not supported

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
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetPointer(addr2.Native)

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
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetPointer(addr2.Native)

	var ret gi.Argument

	err := addressEqualByNameFunction_Set()
	if err == nil {
		ret = addressEqualByNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'soup_address_get_gsockaddr' : return type 'Gio.SocketAddress' not supported

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
	inArgs[0].SetPointer(recv.Native)

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
	inArgs[0].SetPointer(recv.Native)

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
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := addressGetPortFunction_Set()
	if err == nil {
		ret = addressGetPortFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo
}

// UNSUPPORTED : C value 'soup_address_get_sockaddr' : return type 'gpointer' not supported

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
	inArgs[0].SetPointer(recv.Native)

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
	inArgs[0].SetPointer(recv.Native)

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
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := addressIsResolvedFunction_Set()
	if err == nil {
		ret = addressIsResolvedFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'soup_address_resolve_async' : parameter 'async_context' of type 'GLib.MainContext' not supported

// UNSUPPORTED : C value 'soup_address_resolve_sync' : parameter 'cancellable' of type 'Gio.Cancellable' not supported

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
	gobject.Object
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'GObject.Object'

// FieldRealm returns the C field 'realm'.
func (recv *Auth) FieldRealm() string {
	argValue := gi.ObjectFieldGet(authObject, recv.Native, "realm")
	value := argValue.String(false)
	return value
}

// SetFieldRealm sets the value of the C field 'realm'.
func (recv *Auth) SetFieldRealm(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.ObjectFieldSet(authObject, recv.Native, "realm", argValue)
}

// UNSUPPORTED : C value 'soup_auth_new' : parameter 'type' of type 'GType' not supported

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
	inArgs[0].SetPointer(recv.Native)
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
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := authCanAuthenticateFunction_Set()
	if err == nil {
		ret = authCanAuthenticateFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'soup_auth_free_protection_space' : parameter 'space' of type 'GLib.SList' not supported

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
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetPointer(msg.Native)

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
	inArgs[0].SetPointer(recv.Native)

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
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := authGetInfoFunction_Set()
	if err == nil {
		ret = authGetInfoFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

// UNSUPPORTED : C value 'soup_auth_get_protection_space' : return type 'GLib.SList' not supported

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
	inArgs[0].SetPointer(recv.Native)

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
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetString(user)

	var ret gi.Argument

	err := authGetSavedPasswordFunction_Set()
	if err == nil {
		ret = authGetSavedPasswordFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'soup_auth_get_saved_users' : return type 'GLib.SList' not supported

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
	inArgs[0].SetPointer(recv.Native)

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
	inArgs[0].SetPointer(recv.Native)
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
	inArgs[0].SetPointer(recv.Native)

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
	inArgs[0].SetPointer(recv.Native)

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
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetPointer(msg.Native)

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
	inArgs[0].SetPointer(recv.Native)
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
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetPointer(msg.Native)
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
	Auth
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
	Auth
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
	gobject.Object
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'GObject.Object'

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
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetPointer(msg.Native)

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
	inArgs[0].SetPointer(recv.Native)
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
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetPointer(msg.Native)

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
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetPointer(msg.Native)
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
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetPointer(msg.Native)

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
	inArgs[0].SetPointer(recv.Native)

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
	inArgs[0].SetPointer(recv.Native)
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
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetPointer(msg.Native)
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
	AuthDomain
}

// FieldParent returns the C field 'parent'.
func (recv *AuthDomainBasic) FieldParent() *AuthDomain {
	argValue := gi.ObjectFieldGet(authDomainBasicObject, recv.Native, "parent")
	value := &AuthDomain{}
	value.Native = argValue.Pointer()
	return value
}

// SetFieldParent sets the value of the C field 'parent'.
func (recv *AuthDomainBasic) SetFieldParent(value *AuthDomain) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native)
	gi.ObjectFieldSet(authDomainBasicObject, recv.Native, "parent", argValue)
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
	AuthDomain
}

// FieldParent returns the C field 'parent'.
func (recv *AuthDomainDigest) FieldParent() *AuthDomain {
	argValue := gi.ObjectFieldGet(authDomainDigestObject, recv.Native, "parent")
	value := &AuthDomain{}
	value.Native = argValue.Pointer()
	return value
}

// SetFieldParent sets the value of the C field 'parent'.
func (recv *AuthDomainDigest) SetFieldParent(value *AuthDomain) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native)
	gi.ObjectFieldSet(authDomainDigestObject, recv.Native, "parent", argValue)
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
	gobject.Object
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'GObject.Object'

// FieldPriv returns the C field 'priv'.
func (recv *AuthManager) FieldPriv() *AuthManagerPrivate {
	argValue := gi.ObjectFieldGet(authManagerObject, recv.Native, "priv")
	value := &AuthManagerPrivate{}
	value.Native = argValue.Pointer()
	return value
}

// SetFieldPriv sets the value of the C field 'priv'.
func (recv *AuthManager) SetFieldPriv(value *AuthManagerPrivate) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native)
	gi.ObjectFieldSet(authManagerObject, recv.Native, "priv", argValue)
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
	inArgs[0].SetPointer(recv.Native)

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
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetPointer(uri.Native)
	inArgs[2].SetPointer(auth.Native)

	err := authManagerUseAuthFunction_Set()
	if err == nil {
		authManagerUseAuthFunction.Invoke(inArgs[:], nil)
	}

	return
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
	Auth
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
	Auth
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
	gobject.Object
}

// UNSUPPORTED : C value 'parent_instance' : for field getter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'parent_instance' : for field setter : no Go type for 'GObject.Object'

// FieldPriv returns the C field 'priv'.
func (recv *Cache) FieldPriv() *CachePrivate {
	argValue := gi.ObjectFieldGet(cacheObject, recv.Native, "priv")
	value := &CachePrivate{}
	value.Native = argValue.Pointer()
	return value
}

// SetFieldPriv sets the value of the C field 'priv'.
func (recv *Cache) SetFieldPriv(value *CachePrivate) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native)
	gi.ObjectFieldSet(cacheObject, recv.Native, "priv", argValue)
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

	retGo := &Cache{}
	retGo.Native = ret.Pointer()

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
	inArgs[0].SetPointer(recv.Native)

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
	inArgs[0].SetPointer(recv.Native)

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
	inArgs[0].SetPointer(recv.Native)

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
	inArgs[0].SetPointer(recv.Native)

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
	inArgs[0].SetPointer(recv.Native)

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
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetUint32(maxSize)

	err := cacheSetMaxSizeFunction_Set()
	if err == nil {
		cacheSetMaxSizeFunction.Invoke(inArgs[:], nil)
	}

	return
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
	gobject.Object
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'GObject.Object'

// FieldPriv returns the C field 'priv'.
func (recv *ContentDecoder) FieldPriv() *ContentDecoderPrivate {
	argValue := gi.ObjectFieldGet(contentDecoderObject, recv.Native, "priv")
	value := &ContentDecoderPrivate{}
	value.Native = argValue.Pointer()
	return value
}

// SetFieldPriv sets the value of the C field 'priv'.
func (recv *ContentDecoder) SetFieldPriv(value *ContentDecoderPrivate) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native)
	gi.ObjectFieldSet(contentDecoderObject, recv.Native, "priv", argValue)
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
	gobject.Object
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'GObject.Object'

// FieldPriv returns the C field 'priv'.
func (recv *ContentSniffer) FieldPriv() *ContentSnifferPrivate {
	argValue := gi.ObjectFieldGet(contentSnifferObject, recv.Native, "priv")
	value := &ContentSnifferPrivate{}
	value.Native = argValue.Pointer()
	return value
}

// SetFieldPriv sets the value of the C field 'priv'.
func (recv *ContentSniffer) SetFieldPriv(value *ContentSnifferPrivate) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native)
	gi.ObjectFieldSet(contentSnifferObject, recv.Native, "priv", argValue)
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

	retGo := &ContentSniffer{}
	retGo.Native = ret.Pointer()

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
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := contentSnifferGetBufferSizeFunction_Set()
	if err == nil {
		ret = contentSnifferGetBufferSizeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint64()

	return retGo
}

// UNSUPPORTED : C value 'soup_content_sniffer_sniff' : parameter 'params' of type 'GLib.HashTable' not supported

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
	gobject.Object
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'GObject.Object'

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

	retGo := &CookieJar{}
	retGo.Native = ret.Pointer()

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
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetPointer(cookie.Native)

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
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetPointer(cookie.Native)
	inArgs[2].SetPointer(uri.Native)
	inArgs[3].SetPointer(firstParty.Native)

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
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetPointer(firstParty.Native)
	inArgs[2].SetPointer(cookie.Native)

	err := cookieJarAddCookieWithFirstPartyFunction_Set()
	if err == nil {
		cookieJarAddCookieWithFirstPartyFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'soup_cookie_jar_all_cookies' : return type 'GLib.SList' not supported

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
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetPointer(cookie.Native)

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
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := cookieJarGetAcceptPolicyFunction_Set()
	if err == nil {
		ret = cookieJarGetAcceptPolicyFunction.Invoke(inArgs[:], nil)
	}

	retGo := CookieJarAcceptPolicy(ret.Int32())

	return retGo
}

// UNSUPPORTED : C value 'soup_cookie_jar_get_cookie_list' : return type 'GLib.SList' not supported

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
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetPointer(uri.Native)
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
	inArgs[0].SetPointer(recv.Native)

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
	inArgs[0].SetPointer(recv.Native)

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
	inArgs[0].SetPointer(recv.Native)
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
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetPointer(uri.Native)
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
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetPointer(uri.Native)
	inArgs[2].SetPointer(firstParty.Native)
	inArgs[3].SetString(cookie)

	err := cookieJarSetCookieWithFirstPartyFunction_Set()
	if err == nil {
		cookieJarSetCookieWithFirstPartyFunction.Invoke(inArgs[:], nil)
	}

	return
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
	CookieJar
}

// FieldParent returns the C field 'parent'.
func (recv *CookieJarDB) FieldParent() *CookieJar {
	argValue := gi.ObjectFieldGet(cookieJarDBObject, recv.Native, "parent")
	value := &CookieJar{}
	value.Native = argValue.Pointer()
	return value
}

// SetFieldParent sets the value of the C field 'parent'.
func (recv *CookieJarDB) SetFieldParent(value *CookieJar) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native)
	gi.ObjectFieldSet(cookieJarDBObject, recv.Native, "parent", argValue)
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

	retGo := &CookieJarDB{}
	retGo.Native = ret.Pointer()

	return retGo
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
	CookieJar
}

// FieldParent returns the C field 'parent'.
func (recv *CookieJarText) FieldParent() *CookieJar {
	argValue := gi.ObjectFieldGet(cookieJarTextObject, recv.Native, "parent")
	value := &CookieJar{}
	value.Native = argValue.Pointer()
	return value
}

// SetFieldParent sets the value of the C field 'parent'.
func (recv *CookieJarText) SetFieldParent(value *CookieJar) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native)
	gi.ObjectFieldSet(cookieJarTextObject, recv.Native, "parent", argValue)
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

	retGo := &CookieJarText{}
	retGo.Native = ret.Pointer()

	return retGo
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
	gobject.Object
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'GObject.Object'

// FieldPriv returns the C field 'priv'.
func (recv *HSTSEnforcer) FieldPriv() *HSTSEnforcerPrivate {
	argValue := gi.ObjectFieldGet(hSTSEnforcerObject, recv.Native, "priv")
	value := &HSTSEnforcerPrivate{}
	value.Native = argValue.Pointer()
	return value
}

// SetFieldPriv sets the value of the C field 'priv'.
func (recv *HSTSEnforcer) SetFieldPriv(value *HSTSEnforcerPrivate) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native)
	gi.ObjectFieldSet(hSTSEnforcerObject, recv.Native, "priv", argValue)
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

	retGo := &HSTSEnforcer{}
	retGo.Native = ret.Pointer()

	return retGo
}

// UNSUPPORTED : C value 'soup_hsts_enforcer_get_domains' : return type 'GLib.List' not supported

// UNSUPPORTED : C value 'soup_hsts_enforcer_get_policies' : return type 'GLib.List' not supported

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
	inArgs[0].SetPointer(recv.Native)
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
	inArgs[0].SetPointer(recv.Native)

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
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetPointer(policy.Native)

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
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetString(domain)
	inArgs[2].SetBoolean(includeSubdomains)

	err := hSTSEnforcerSetSessionPolicyFunction_Set()
	if err == nil {
		hSTSEnforcerSetSessionPolicyFunction.Invoke(inArgs[:], nil)
	}

	return
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
	HSTSEnforcer
}

// FieldParent returns the C field 'parent'.
func (recv *HSTSEnforcerDB) FieldParent() *HSTSEnforcer {
	argValue := gi.ObjectFieldGet(hSTSEnforcerDBObject, recv.Native, "parent")
	value := &HSTSEnforcer{}
	value.Native = argValue.Pointer()
	return value
}

// SetFieldParent sets the value of the C field 'parent'.
func (recv *HSTSEnforcerDB) SetFieldParent(value *HSTSEnforcer) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native)
	gi.ObjectFieldSet(hSTSEnforcerDBObject, recv.Native, "parent", argValue)
}

// FieldPriv returns the C field 'priv'.
func (recv *HSTSEnforcerDB) FieldPriv() *HSTSEnforcerDBPrivate {
	argValue := gi.ObjectFieldGet(hSTSEnforcerDBObject, recv.Native, "priv")
	value := &HSTSEnforcerDBPrivate{}
	value.Native = argValue.Pointer()
	return value
}

// SetFieldPriv sets the value of the C field 'priv'.
func (recv *HSTSEnforcerDB) SetFieldPriv(value *HSTSEnforcerDBPrivate) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native)
	gi.ObjectFieldSet(hSTSEnforcerDBObject, recv.Native, "priv", argValue)
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

	retGo := &HSTSEnforcerDB{}
	retGo.Native = ret.Pointer()

	return retGo
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
	gobject.Object
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'GObject.Object'

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

	retGo := &Logger{}
	retGo.Native = ret.Pointer()

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
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetPointer(session.Native)

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
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetPointer(session.Native)

	err := loggerDetachFunction_Set()
	if err == nil {
		loggerDetachFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'soup_logger_set_printer' : parameter 'printer' of type 'LoggerPrinter' not supported

// UNSUPPORTED : C value 'soup_logger_set_request_filter' : parameter 'request_filter' of type 'LoggerFilter' not supported

// UNSUPPORTED : C value 'soup_logger_set_response_filter' : parameter 'response_filter' of type 'LoggerFilter' not supported

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
	gobject.Object
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'GObject.Object'

// FieldMethod returns the C field 'method'.
func (recv *Message) FieldMethod() string {
	argValue := gi.ObjectFieldGet(messageObject, recv.Native, "method")
	value := argValue.String(false)
	return value
}

// SetFieldMethod sets the value of the C field 'method'.
func (recv *Message) SetFieldMethod(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.ObjectFieldSet(messageObject, recv.Native, "method", argValue)
}

// FieldStatusCode returns the C field 'status_code'.
func (recv *Message) FieldStatusCode() uint32 {
	argValue := gi.ObjectFieldGet(messageObject, recv.Native, "status_code")
	value := argValue.Uint32()
	return value
}

// SetFieldStatusCode sets the value of the C field 'status_code'.
func (recv *Message) SetFieldStatusCode(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.ObjectFieldSet(messageObject, recv.Native, "status_code", argValue)
}

// FieldReasonPhrase returns the C field 'reason_phrase'.
func (recv *Message) FieldReasonPhrase() string {
	argValue := gi.ObjectFieldGet(messageObject, recv.Native, "reason_phrase")
	value := argValue.String(false)
	return value
}

// SetFieldReasonPhrase sets the value of the C field 'reason_phrase'.
func (recv *Message) SetFieldReasonPhrase(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.ObjectFieldSet(messageObject, recv.Native, "reason_phrase", argValue)
}

// FieldRequestBody returns the C field 'request_body'.
func (recv *Message) FieldRequestBody() *MessageBody {
	argValue := gi.ObjectFieldGet(messageObject, recv.Native, "request_body")
	value := &MessageBody{}
	value.Native = argValue.Pointer()
	return value
}

// SetFieldRequestBody sets the value of the C field 'request_body'.
func (recv *Message) SetFieldRequestBody(value *MessageBody) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native)
	gi.ObjectFieldSet(messageObject, recv.Native, "request_body", argValue)
}

// FieldRequestHeaders returns the C field 'request_headers'.
func (recv *Message) FieldRequestHeaders() *MessageHeaders {
	argValue := gi.ObjectFieldGet(messageObject, recv.Native, "request_headers")
	value := &MessageHeaders{}
	value.Native = argValue.Pointer()
	return value
}

// SetFieldRequestHeaders sets the value of the C field 'request_headers'.
func (recv *Message) SetFieldRequestHeaders(value *MessageHeaders) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native)
	gi.ObjectFieldSet(messageObject, recv.Native, "request_headers", argValue)
}

// FieldResponseBody returns the C field 'response_body'.
func (recv *Message) FieldResponseBody() *MessageBody {
	argValue := gi.ObjectFieldGet(messageObject, recv.Native, "response_body")
	value := &MessageBody{}
	value.Native = argValue.Pointer()
	return value
}

// SetFieldResponseBody sets the value of the C field 'response_body'.
func (recv *Message) SetFieldResponseBody(value *MessageBody) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native)
	gi.ObjectFieldSet(messageObject, recv.Native, "response_body", argValue)
}

// FieldResponseHeaders returns the C field 'response_headers'.
func (recv *Message) FieldResponseHeaders() *MessageHeaders {
	argValue := gi.ObjectFieldGet(messageObject, recv.Native, "response_headers")
	value := &MessageHeaders{}
	value.Native = argValue.Pointer()
	return value
}

// SetFieldResponseHeaders sets the value of the C field 'response_headers'.
func (recv *Message) SetFieldResponseHeaders(value *MessageHeaders) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native)
	gi.ObjectFieldSet(messageObject, recv.Native, "response_headers", argValue)
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

	retGo := &Message{}
	retGo.Native = ret.Pointer()

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
	inArgs[1].SetPointer(uri.Native)

	var ret gi.Argument

	err := messageNewFromUriFunction_Set()
	if err == nil {
		ret = messageNewFromUriFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Message{}
	retGo.Native = ret.Pointer()

	return retGo
}

// UNSUPPORTED : C value 'soup_message_add_header_handler' : parameter 'callback' of type 'GObject.Callback' not supported

// UNSUPPORTED : C value 'soup_message_add_status_code_handler' : parameter 'callback' of type 'GObject.Callback' not supported

// UNSUPPORTED : C value 'soup_message_content_sniffed' : parameter 'params' of type 'GLib.HashTable' not supported

// UNSUPPORTED : C value 'soup_message_disable_feature' : parameter 'feature_type' of type 'GType' not supported

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
	inArgs[0].SetPointer(recv.Native)

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
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := messageGetAddressFunction_Set()
	if err == nil {
		ret = messageGetAddressFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Address{}
	retGo.Native = ret.Pointer()

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
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := messageGetFirstPartyFunction_Set()
	if err == nil {
		ret = messageGetFirstPartyFunction.Invoke(inArgs[:], nil)
	}

	retGo := &URI{}
	retGo.Native = ret.Pointer()

	return retGo
}

// UNSUPPORTED : C value 'soup_message_get_flags' : return type 'MessageFlags' not supported

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
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := messageGetHttpVersionFunction_Set()
	if err == nil {
		ret = messageGetHttpVersionFunction.Invoke(inArgs[:], nil)
	}

	retGo := HTTPVersion(ret.Int32())

	return retGo
}

// UNSUPPORTED : C value 'soup_message_get_https_status' : parameter 'certificate' of type 'Gio.TlsCertificate' not supported

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
	inArgs[0].SetPointer(recv.Native)

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
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := messageGetSoupRequestFunction_Set()
	if err == nil {
		ret = messageGetSoupRequestFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Request{}
	retGo.Native = ret.Pointer()

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
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := messageGetUriFunction_Set()
	if err == nil {
		ret = messageGetUriFunction.Invoke(inArgs[:], nil)
	}

	retGo := &URI{}
	retGo.Native = ret.Pointer()

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
	inArgs[0].SetPointer(recv.Native)

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
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetPointer(chunk.Native)

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
	inArgs[0].SetPointer(recv.Native)

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
	inArgs[0].SetPointer(recv.Native)

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
	inArgs[0].SetPointer(recv.Native)

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
	inArgs[0].SetPointer(recv.Native)

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
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetPointer(firstParty.Native)

	err := messageSetFirstPartyFunction_Set()
	if err == nil {
		messageSetFirstPartyFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'soup_message_set_flags' : parameter 'flags' of type 'MessageFlags' not supported

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
	inArgs[0].SetPointer(recv.Native)
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
	inArgs[0].SetPointer(recv.Native)
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
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetUint32(statusCode)
	inArgs[2].SetString(redirectUri)

	err := messageSetRedirectFunction_Set()
	if err == nil {
		messageSetRedirectFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'soup_message_set_request' : parameter 'req_body' of type 'nil' not supported

// UNSUPPORTED : C value 'soup_message_set_response' : parameter 'resp_body' of type 'nil' not supported

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
	inArgs[0].SetPointer(recv.Native)
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
	inArgs[0].SetPointer(recv.Native)
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
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetPointer(uri.Native)

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
	inArgs[0].SetPointer(recv.Native)

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
	inArgs[0].SetPointer(recv.Native)

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
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetPointer(chunk.Native)

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
	inArgs[0].SetPointer(recv.Native)

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
	inArgs[0].SetPointer(recv.Native)

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
	inArgs[0].SetPointer(recv.Native)

	err := messageWroteInformationalFunction_Set()
	if err == nil {
		messageWroteInformationalFunction.Invoke(inArgs[:], nil)
	}

	return
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
	gio.FilterInputStream
}

// UNSUPPORTED : C value 'parent_instance' : for field getter : no Go type for 'Gio.FilterInputStream'

// UNSUPPORTED : C value 'parent_instance' : for field setter : no Go type for 'Gio.FilterInputStream'

// UNSUPPORTED : C value 'soup_multipart_input_stream_new' : parameter 'base_stream' of type 'Gio.InputStream' not supported

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
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := multipartInputStreamGetHeadersFunction_Set()
	if err == nil {
		ret = multipartInputStreamGetHeadersFunction.Invoke(inArgs[:], nil)
	}

	retGo := &MessageHeaders{}
	retGo.Native = ret.Pointer()

	return retGo
}

// UNSUPPORTED : C value 'soup_multipart_input_stream_next_part' : parameter 'cancellable' of type 'Gio.Cancellable' not supported

// UNSUPPORTED : C value 'soup_multipart_input_stream_next_part_async' : parameter 'cancellable' of type 'Gio.Cancellable' not supported

// UNSUPPORTED : C value 'soup_multipart_input_stream_next_part_finish' : parameter 'result' of type 'Gio.AsyncResult' not supported

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
	gobject.Object
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'GObject.Object'

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
	gobject.Object
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'GObject.Object'

// FieldPriv returns the C field 'priv'.
func (recv *Request) FieldPriv() *RequestPrivate {
	argValue := gi.ObjectFieldGet(requestObject, recv.Native, "priv")
	value := &RequestPrivate{}
	value.Native = argValue.Pointer()
	return value
}

// SetFieldPriv sets the value of the C field 'priv'.
func (recv *Request) SetFieldPriv(value *RequestPrivate) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native)
	gi.ObjectFieldSet(requestObject, recv.Native, "priv", argValue)
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
	inArgs[0].SetPointer(recv.Native)

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
	inArgs[0].SetPointer(recv.Native)

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
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := requestGetSessionFunction_Set()
	if err == nil {
		ret = requestGetSessionFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Session{}
	retGo.Native = ret.Pointer()

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
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := requestGetUriFunction_Set()
	if err == nil {
		ret = requestGetUriFunction.Invoke(inArgs[:], nil)
	}

	retGo := &URI{}
	retGo.Native = ret.Pointer()

	return retGo
}

// UNSUPPORTED : C value 'soup_request_send' : parameter 'cancellable' of type 'Gio.Cancellable' not supported

// UNSUPPORTED : C value 'soup_request_send_async' : parameter 'cancellable' of type 'Gio.Cancellable' not supported

// UNSUPPORTED : C value 'soup_request_send_finish' : parameter 'result' of type 'Gio.AsyncResult' not supported

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
	Request
}

// FieldParent returns the C field 'parent'.
func (recv *RequestData) FieldParent() *Request {
	argValue := gi.ObjectFieldGet(requestDataObject, recv.Native, "parent")
	value := &Request{}
	value.Native = argValue.Pointer()
	return value
}

// SetFieldParent sets the value of the C field 'parent'.
func (recv *RequestData) SetFieldParent(value *Request) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native)
	gi.ObjectFieldSet(requestDataObject, recv.Native, "parent", argValue)
}

// FieldPriv returns the C field 'priv'.
func (recv *RequestData) FieldPriv() *RequestDataPrivate {
	argValue := gi.ObjectFieldGet(requestDataObject, recv.Native, "priv")
	value := &RequestDataPrivate{}
	value.Native = argValue.Pointer()
	return value
}

// SetFieldPriv sets the value of the C field 'priv'.
func (recv *RequestData) SetFieldPriv(value *RequestDataPrivate) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native)
	gi.ObjectFieldSet(requestDataObject, recv.Native, "priv", argValue)
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
	Request
}

// FieldParent returns the C field 'parent'.
func (recv *RequestFile) FieldParent() *Request {
	argValue := gi.ObjectFieldGet(requestFileObject, recv.Native, "parent")
	value := &Request{}
	value.Native = argValue.Pointer()
	return value
}

// SetFieldParent sets the value of the C field 'parent'.
func (recv *RequestFile) SetFieldParent(value *Request) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native)
	gi.ObjectFieldSet(requestFileObject, recv.Native, "parent", argValue)
}

// FieldPriv returns the C field 'priv'.
func (recv *RequestFile) FieldPriv() *RequestFilePrivate {
	argValue := gi.ObjectFieldGet(requestFileObject, recv.Native, "priv")
	value := &RequestFilePrivate{}
	value.Native = argValue.Pointer()
	return value
}

// SetFieldPriv sets the value of the C field 'priv'.
func (recv *RequestFile) SetFieldPriv(value *RequestFilePrivate) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native)
	gi.ObjectFieldSet(requestFileObject, recv.Native, "priv", argValue)
}

// UNSUPPORTED : C value 'soup_request_file_get_file' : return type 'Gio.File' not supported

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
	Request
}

// FieldParent returns the C field 'parent'.
func (recv *RequestHTTP) FieldParent() *Request {
	argValue := gi.ObjectFieldGet(requestHTTPObject, recv.Native, "parent")
	value := &Request{}
	value.Native = argValue.Pointer()
	return value
}

// SetFieldParent sets the value of the C field 'parent'.
func (recv *RequestHTTP) SetFieldParent(value *Request) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native)
	gi.ObjectFieldSet(requestHTTPObject, recv.Native, "parent", argValue)
}

// FieldPriv returns the C field 'priv'.
func (recv *RequestHTTP) FieldPriv() *RequestHTTPPrivate {
	argValue := gi.ObjectFieldGet(requestHTTPObject, recv.Native, "priv")
	value := &RequestHTTPPrivate{}
	value.Native = argValue.Pointer()
	return value
}

// SetFieldPriv sets the value of the C field 'priv'.
func (recv *RequestHTTP) SetFieldPriv(value *RequestHTTPPrivate) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native)
	gi.ObjectFieldSet(requestHTTPObject, recv.Native, "priv", argValue)
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
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := requestHTTPGetMessageFunction_Set()
	if err == nil {
		ret = requestHTTPGetMessageFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Message{}
	retGo.Native = ret.Pointer()

	return retGo
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
	gobject.Object
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'GObject.Object'

// FieldPriv returns the C field 'priv'.
func (recv *Requester) FieldPriv() *RequesterPrivate {
	argValue := gi.ObjectFieldGet(requesterObject, recv.Native, "priv")
	value := &RequesterPrivate{}
	value.Native = argValue.Pointer()
	return value
}

// SetFieldPriv sets the value of the C field 'priv'.
func (recv *Requester) SetFieldPriv(value *RequesterPrivate) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native)
	gi.ObjectFieldSet(requesterObject, recv.Native, "priv", argValue)
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

	retGo := &Requester{}
	retGo.Native = ret.Pointer()

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

// Request is a representation of the C type soup_requester_request.
func (recv *Requester) Request(uriString string) *Request {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetString(uriString)

	var ret gi.Argument

	err := requesterRequestFunction_Set()
	if err == nil {
		ret = requesterRequestFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Request{}
	retGo.Native = ret.Pointer()

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
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetPointer(uri.Native)

	var ret gi.Argument

	err := requesterRequestUriFunction_Set()
	if err == nil {
		ret = requesterRequestUriFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Request{}
	retGo.Native = ret.Pointer()

	return retGo
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
	gobject.Object
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'soup_server_new' : parameter '...' of type 'nil' not supported

// UNSUPPORTED : C value 'soup_server_accept_iostream' : parameter 'stream' of type 'Gio.IOStream' not supported

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
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetPointer(authDomain.Native)

	err := serverAddAuthDomainFunction_Set()
	if err == nil {
		serverAddAuthDomainFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'soup_server_add_early_handler' : parameter 'callback' of type 'ServerCallback' not supported

// UNSUPPORTED : C value 'soup_server_add_handler' : parameter 'callback' of type 'ServerCallback' not supported

// UNSUPPORTED : C value 'soup_server_add_websocket_extension' : parameter 'extension_type' of type 'GType' not supported

// UNSUPPORTED : C value 'soup_server_add_websocket_handler' : parameter 'protocols' of type 'nil' not supported

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
	inArgs[0].SetPointer(recv.Native)

	err := serverDisconnectFunction_Set()
	if err == nil {
		serverDisconnectFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'soup_server_get_async_context' : return type 'GLib.MainContext' not supported

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
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := serverGetListenerFunction_Set()
	if err == nil {
		ret = serverGetListenerFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Socket{}
	retGo.Native = ret.Pointer()

	return retGo
}

// UNSUPPORTED : C value 'soup_server_get_listeners' : return type 'GLib.SList' not supported

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
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := serverGetPortFunction_Set()
	if err == nil {
		ret = serverGetPortFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo
}

// UNSUPPORTED : C value 'soup_server_get_uris' : return type 'GLib.SList' not supported

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
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := serverIsHttpsFunction_Set()
	if err == nil {
		ret = serverIsHttpsFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'soup_server_listen' : parameter 'address' of type 'Gio.SocketAddress' not supported

// UNSUPPORTED : C value 'soup_server_listen_all' : parameter 'options' of type 'ServerListenOptions' not supported

// UNSUPPORTED : C value 'soup_server_listen_fd' : parameter 'options' of type 'ServerListenOptions' not supported

// UNSUPPORTED : C value 'soup_server_listen_local' : parameter 'options' of type 'ServerListenOptions' not supported

// UNSUPPORTED : C value 'soup_server_listen_socket' : parameter 'socket' of type 'Gio.Socket' not supported

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
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetPointer(msg.Native)

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
	inArgs[0].SetPointer(recv.Native)

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
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetPointer(authDomain.Native)

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
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetString(path)

	err := serverRemoveHandlerFunction_Set()
	if err == nil {
		serverRemoveHandlerFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'soup_server_remove_websocket_extension' : parameter 'extension_type' of type 'GType' not supported

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
	inArgs[0].SetPointer(recv.Native)

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
	inArgs[0].SetPointer(recv.Native)

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
	inArgs[0].SetPointer(recv.Native)
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
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetPointer(msg.Native)

	err := serverUnpauseMessageFunction_Set()
	if err == nil {
		serverUnpauseMessageFunction.Invoke(inArgs[:], nil)
	}

	return
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
	gobject.Object
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'GObject.Object'

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

	retGo := &Session{}
	retGo.Native = ret.Pointer()

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
	inArgs[0].SetPointer(recv.Native)

	err := sessionAbortFunction_Set()
	if err == nil {
		sessionAbortFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'soup_session_add_feature' : parameter 'feature' of type 'SessionFeature' not supported

// UNSUPPORTED : C value 'soup_session_add_feature_by_type' : parameter 'feature_type' of type 'GType' not supported

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
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetPointer(msg.Native)
	inArgs[2].SetUint32(statusCode)

	err := sessionCancelMessageFunction_Set()
	if err == nil {
		sessionCancelMessageFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'soup_session_connect_async' : parameter 'cancellable' of type 'Gio.Cancellable' not supported

// UNSUPPORTED : C value 'soup_session_connect_finish' : parameter 'result' of type 'Gio.AsyncResult' not supported

// UNSUPPORTED : C value 'soup_session_get_async_context' : return type 'GLib.MainContext' not supported

// UNSUPPORTED : C value 'soup_session_get_feature' : parameter 'feature_type' of type 'GType' not supported

// UNSUPPORTED : C value 'soup_session_get_feature_for_message' : parameter 'feature_type' of type 'GType' not supported

// UNSUPPORTED : C value 'soup_session_get_features' : parameter 'feature_type' of type 'GType' not supported

// UNSUPPORTED : C value 'soup_session_has_feature' : parameter 'feature_type' of type 'GType' not supported

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
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetPointer(msg.Native)

	err := sessionPauseMessageFunction_Set()
	if err == nil {
		sessionPauseMessageFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'soup_session_prefetch_dns' : parameter 'cancellable' of type 'Gio.Cancellable' not supported

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
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetPointer(uri.Native)

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
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetPointer(msg.Native)

	var ret gi.Argument

	err := sessionRedirectMessageFunction_Set()
	if err == nil {
		ret = sessionRedirectMessageFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'soup_session_remove_feature' : parameter 'feature' of type 'SessionFeature' not supported

// UNSUPPORTED : C value 'soup_session_remove_feature_by_type' : parameter 'feature_type' of type 'GType' not supported

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

// Request is a representation of the C type soup_session_request.
func (recv *Session) Request(uriString string) *Request {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetString(uriString)

	var ret gi.Argument

	err := sessionRequestFunction_Set()
	if err == nil {
		ret = sessionRequestFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Request{}
	retGo.Native = ret.Pointer()

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
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetString(method)
	inArgs[2].SetString(uriString)

	var ret gi.Argument

	err := sessionRequestHttpFunction_Set()
	if err == nil {
		ret = sessionRequestHttpFunction.Invoke(inArgs[:], nil)
	}

	retGo := &RequestHTTP{}
	retGo.Native = ret.Pointer()

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
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetString(method)
	inArgs[2].SetPointer(uri.Native)

	var ret gi.Argument

	err := sessionRequestHttpUriFunction_Set()
	if err == nil {
		ret = sessionRequestHttpUriFunction.Invoke(inArgs[:], nil)
	}

	retGo := &RequestHTTP{}
	retGo.Native = ret.Pointer()

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
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetPointer(uri.Native)

	var ret gi.Argument

	err := sessionRequestUriFunction_Set()
	if err == nil {
		ret = sessionRequestUriFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Request{}
	retGo.Native = ret.Pointer()

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
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetPointer(msg.Native)

	err := sessionRequeueMessageFunction_Set()
	if err == nil {
		sessionRequeueMessageFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'soup_session_send' : parameter 'cancellable' of type 'Gio.Cancellable' not supported

// UNSUPPORTED : C value 'soup_session_send_async' : parameter 'cancellable' of type 'Gio.Cancellable' not supported

// UNSUPPORTED : C value 'soup_session_send_finish' : parameter 'result' of type 'Gio.AsyncResult' not supported

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
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetPointer(msg.Native)

	var ret gi.Argument

	err := sessionSendMessageFunction_Set()
	if err == nil {
		ret = sessionSendMessageFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo
}

// UNSUPPORTED : C value 'soup_session_steal_connection' : return type 'Gio.IOStream' not supported

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
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetPointer(msg.Native)

	err := sessionUnpauseMessageFunction_Set()
	if err == nil {
		sessionUnpauseMessageFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'soup_session_websocket_connect_async' : parameter 'protocols' of type 'nil' not supported

// UNSUPPORTED : C value 'soup_session_websocket_connect_finish' : parameter 'result' of type 'Gio.AsyncResult' not supported

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
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetPointer(msg.Native)

	var ret gi.Argument

	err := sessionWouldRedirectFunction_Set()
	if err == nil {
		ret = sessionWouldRedirectFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
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
	Session
}

// FieldParent returns the C field 'parent'.
func (recv *SessionAsync) FieldParent() *Session {
	argValue := gi.ObjectFieldGet(sessionAsyncObject, recv.Native, "parent")
	value := &Session{}
	value.Native = argValue.Pointer()
	return value
}

// SetFieldParent sets the value of the C field 'parent'.
func (recv *SessionAsync) SetFieldParent(value *Session) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native)
	gi.ObjectFieldSet(sessionAsyncObject, recv.Native, "parent", argValue)
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

	retGo := &SessionAsync{}
	retGo.Native = ret.Pointer()

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
	Session
}

// FieldParent returns the C field 'parent'.
func (recv *SessionSync) FieldParent() *Session {
	argValue := gi.ObjectFieldGet(sessionSyncObject, recv.Native, "parent")
	value := &Session{}
	value.Native = argValue.Pointer()
	return value
}

// SetFieldParent sets the value of the C field 'parent'.
func (recv *SessionSync) SetFieldParent(value *Session) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native)
	gi.ObjectFieldSet(sessionSyncObject, recv.Native, "parent", argValue)
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

	retGo := &SessionSync{}
	retGo.Native = ret.Pointer()

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
	gobject.Object
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'soup_socket_new' : parameter '...' of type 'nil' not supported

// UNSUPPORTED : C value 'soup_socket_connect_async' : parameter 'cancellable' of type 'Gio.Cancellable' not supported

// UNSUPPORTED : C value 'soup_socket_connect_sync' : parameter 'cancellable' of type 'Gio.Cancellable' not supported

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
	inArgs[0].SetPointer(recv.Native)

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
	inArgs[0].SetPointer(recv.Native)

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
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := socketGetLocalAddressFunction_Set()
	if err == nil {
		ret = socketGetLocalAddressFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Address{}
	retGo.Native = ret.Pointer()

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
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := socketGetRemoteAddressFunction_Set()
	if err == nil {
		ret = socketGetRemoteAddressFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Address{}
	retGo.Native = ret.Pointer()

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
	inArgs[0].SetPointer(recv.Native)

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
	inArgs[0].SetPointer(recv.Native)

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
	inArgs[0].SetPointer(recv.Native)

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

// UNSUPPORTED : C value 'soup_socket_start_proxy_ssl' : parameter 'cancellable' of type 'Gio.Cancellable' not supported

// UNSUPPORTED : C value 'soup_socket_start_ssl' : parameter 'cancellable' of type 'Gio.Cancellable' not supported

// UNSUPPORTED : C value 'soup_socket_write' : parameter 'buffer' of type 'nil' not supported

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
	gobject.Object
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'soup_websocket_connection_new' : parameter 'stream' of type 'Gio.IOStream' not supported

// UNSUPPORTED : C value 'soup_websocket_connection_new_with_extensions' : parameter 'stream' of type 'Gio.IOStream' not supported

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
	inArgs[0].SetPointer(recv.Native)
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
	inArgs[0].SetPointer(recv.Native)

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
	inArgs[0].SetPointer(recv.Native)

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
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := websocketConnectionGetConnectionTypeFunction_Set()
	if err == nil {
		ret = websocketConnectionGetConnectionTypeFunction.Invoke(inArgs[:], nil)
	}

	retGo := WebsocketConnectionType(ret.Int32())

	return retGo
}

// UNSUPPORTED : C value 'soup_websocket_connection_get_extensions' : return type 'GLib.List' not supported

// UNSUPPORTED : C value 'soup_websocket_connection_get_io_stream' : return type 'Gio.IOStream' not supported

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
	inArgs[0].SetPointer(recv.Native)

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
	inArgs[0].SetPointer(recv.Native)

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
	inArgs[0].SetPointer(recv.Native)

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
	inArgs[0].SetPointer(recv.Native)

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
	inArgs[0].SetPointer(recv.Native)

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
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := websocketConnectionGetUriFunction_Set()
	if err == nil {
		ret = websocketConnectionGetUriFunction.Invoke(inArgs[:], nil)
	}

	retGo := &URI{}
	retGo.Native = ret.Pointer()

	return retGo
}

// UNSUPPORTED : C value 'soup_websocket_connection_send_binary' : parameter 'data' of type 'nil' not supported

// UNSUPPORTED : C value 'soup_websocket_connection_send_message' : parameter 'message' of type 'GLib.Bytes' not supported

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
	inArgs[0].SetPointer(recv.Native)
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
	inArgs[0].SetPointer(recv.Native)
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
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetUint64(maxIncomingPayloadSize)

	err := websocketConnectionSetMaxIncomingPayloadSizeFunction_Set()
	if err == nil {
		websocketConnectionSetMaxIncomingPayloadSizeFunction.Invoke(inArgs[:], nil)
	}

	return
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
	gobject.Object
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'soup_websocket_extension_configure' : parameter 'params' of type 'GLib.HashTable' not supported

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
	inArgs[0].SetPointer(recv.Native)

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
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := websocketExtensionGetResponseParamsFunction_Set()
	if err == nil {
		ret = websocketExtensionGetResponseParamsFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

// UNSUPPORTED : C value 'soup_websocket_extension_process_incoming_message' : parameter 'payload' of type 'GLib.Bytes' not supported

// UNSUPPORTED : C value 'soup_websocket_extension_process_outgoing_message' : parameter 'payload' of type 'GLib.Bytes' not supported

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
	WebsocketExtension
}

// FieldParent returns the C field 'parent'.
func (recv *WebsocketExtensionDeflate) FieldParent() *WebsocketExtension {
	argValue := gi.ObjectFieldGet(websocketExtensionDeflateObject, recv.Native, "parent")
	value := &WebsocketExtension{}
	value.Native = argValue.Pointer()
	return value
}

// SetFieldParent sets the value of the C field 'parent'.
func (recv *WebsocketExtensionDeflate) SetFieldParent(value *WebsocketExtension) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native)
	gi.ObjectFieldSet(websocketExtensionDeflateObject, recv.Native, "parent", argValue)
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
	gobject.Object
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'GObject.Object'
