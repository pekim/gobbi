// Code generated - DO NOT EDIT.

package soup

import (
	gi "github.com/pekim/gobbi/internal/gi"
	"runtime"
	"sync"
)

var addressStruct *gi.Struct
var addressStruct_Once sync.Once

func addressStruct_Set() error {
	var err error
	addressStruct_Once.Do(func() {
		addressStruct, err = gi.StructNew("Soup", "Address")
	})
	return err
}

type Address struct {
	native uintptr
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'soup_address_new' : return type 'Address' not supported

// UNSUPPORTED : C value 'soup_address_new_any' : parameter 'family' of type 'AddressFamily' not supported

// UNSUPPORTED : C value 'soup_address_new_from_sockaddr' : parameter 'sa' of type 'gpointer' not supported

// UNSUPPORTED : C value 'soup_address_equal_by_ip' : parameter 'addr2' of type 'Address' not supported

// UNSUPPORTED : C value 'soup_address_equal_by_name' : parameter 'addr2' of type 'Address' not supported

// UNSUPPORTED : C value 'soup_address_get_gsockaddr' : return type 'Gio.SocketAddress' not supported

var addressGetNameFunction *gi.Function
var addressGetNameFunction_Once sync.Once

func addressGetNameFunction_Set() error {
	var err error
	addressGetNameFunction_Once.Do(func() {
		err = addressStruct_Set()
		if err != nil {
			return
		}
		addressGetNameFunction, err = addressStruct.InvokerNew("get_name")
	})
	return err
}

// GetName is a representation of the C type soup_address_get_name.
func (recv *Address) GetName() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

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
		err = addressStruct_Set()
		if err != nil {
			return
		}
		addressGetPhysicalFunction, err = addressStruct.InvokerNew("get_physical")
	})
	return err
}

// GetPhysical is a representation of the C type soup_address_get_physical.
func (recv *Address) GetPhysical() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

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
		err = addressStruct_Set()
		if err != nil {
			return
		}
		addressGetPortFunction, err = addressStruct.InvokerNew("get_port")
	})
	return err
}

// GetPort is a representation of the C type soup_address_get_port.
func (recv *Address) GetPort() uint32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

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
		err = addressStruct_Set()
		if err != nil {
			return
		}
		addressHashByIpFunction, err = addressStruct.InvokerNew("hash_by_ip")
	})
	return err
}

// HashByIp is a representation of the C type soup_address_hash_by_ip.
func (recv *Address) HashByIp() uint32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

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
		err = addressStruct_Set()
		if err != nil {
			return
		}
		addressHashByNameFunction, err = addressStruct.InvokerNew("hash_by_name")
	})
	return err
}

// HashByName is a representation of the C type soup_address_hash_by_name.
func (recv *Address) HashByName() uint32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

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
		err = addressStruct_Set()
		if err != nil {
			return
		}
		addressIsResolvedFunction, err = addressStruct.InvokerNew("is_resolved")
	})
	return err
}

// IsResolved is a representation of the C type soup_address_is_resolved.
func (recv *Address) IsResolved() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

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

var authStruct *gi.Struct
var authStruct_Once sync.Once

func authStruct_Set() error {
	var err error
	authStruct_Once.Do(func() {
		authStruct, err = gi.StructNew("Soup", "Auth")
	})
	return err
}

type Auth struct {
	native uintptr
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'GObject.Object'

// FieldRealm returns the C field 'realm'.
func (recv *Auth) FieldRealm() string {
	argValue := gi.FieldGet(authStruct, recv.native, "realm")
	value := argValue.String(false)
	return value
}

// SetFieldRealm sets the value of the C field 'realm'.
func (recv *Auth) SetFieldRealm(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.FieldSet(authStruct, recv.native, "realm", argValue)
}

// UNSUPPORTED : C value 'soup_auth_new' : parameter 'type' of type 'GType' not supported

var authAuthenticateFunction *gi.Function
var authAuthenticateFunction_Once sync.Once

func authAuthenticateFunction_Set() error {
	var err error
	authAuthenticateFunction_Once.Do(func() {
		err = authStruct_Set()
		if err != nil {
			return
		}
		authAuthenticateFunction, err = authStruct.InvokerNew("authenticate")
	})
	return err
}

// Authenticate is a representation of the C type soup_auth_authenticate.
func (recv *Auth) Authenticate(username string, password string) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
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
		err = authStruct_Set()
		if err != nil {
			return
		}
		authCanAuthenticateFunction, err = authStruct.InvokerNew("can_authenticate")
	})
	return err
}

// CanAuthenticate is a representation of the C type soup_auth_can_authenticate.
func (recv *Auth) CanAuthenticate() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := authCanAuthenticateFunction_Set()
	if err == nil {
		ret = authCanAuthenticateFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'soup_auth_free_protection_space' : parameter 'space' of type 'GLib.SList' not supported

// UNSUPPORTED : C value 'soup_auth_get_authorization' : parameter 'msg' of type 'Message' not supported

var authGetHostFunction *gi.Function
var authGetHostFunction_Once sync.Once

func authGetHostFunction_Set() error {
	var err error
	authGetHostFunction_Once.Do(func() {
		err = authStruct_Set()
		if err != nil {
			return
		}
		authGetHostFunction, err = authStruct.InvokerNew("get_host")
	})
	return err
}

// GetHost is a representation of the C type soup_auth_get_host.
func (recv *Auth) GetHost() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

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
		err = authStruct_Set()
		if err != nil {
			return
		}
		authGetInfoFunction, err = authStruct.InvokerNew("get_info")
	})
	return err
}

// GetInfo is a representation of the C type soup_auth_get_info.
func (recv *Auth) GetInfo() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

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
		err = authStruct_Set()
		if err != nil {
			return
		}
		authGetRealmFunction, err = authStruct.InvokerNew("get_realm")
	})
	return err
}

// GetRealm is a representation of the C type soup_auth_get_realm.
func (recv *Auth) GetRealm() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

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
		err = authStruct_Set()
		if err != nil {
			return
		}
		authGetSavedPasswordFunction, err = authStruct.InvokerNew("get_saved_password")
	})
	return err
}

// GetSavedPassword is a representation of the C type soup_auth_get_saved_password.
func (recv *Auth) GetSavedPassword(user string) string {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
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
		err = authStruct_Set()
		if err != nil {
			return
		}
		authGetSchemeNameFunction, err = authStruct.InvokerNew("get_scheme_name")
	})
	return err
}

// GetSchemeName is a representation of the C type soup_auth_get_scheme_name.
func (recv *Auth) GetSchemeName() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

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
		err = authStruct_Set()
		if err != nil {
			return
		}
		authHasSavedPasswordFunction, err = authStruct.InvokerNew("has_saved_password")
	})
	return err
}

// HasSavedPassword is a representation of the C type soup_auth_has_saved_password.
func (recv *Auth) HasSavedPassword(username string, password string) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
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
		err = authStruct_Set()
		if err != nil {
			return
		}
		authIsAuthenticatedFunction, err = authStruct.InvokerNew("is_authenticated")
	})
	return err
}

// IsAuthenticated is a representation of the C type soup_auth_is_authenticated.
func (recv *Auth) IsAuthenticated() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

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
		err = authStruct_Set()
		if err != nil {
			return
		}
		authIsForProxyFunction, err = authStruct.InvokerNew("is_for_proxy")
	})
	return err
}

// IsForProxy is a representation of the C type soup_auth_is_for_proxy.
func (recv *Auth) IsForProxy() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := authIsForProxyFunction_Set()
	if err == nil {
		ret = authIsForProxyFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'soup_auth_is_ready' : parameter 'msg' of type 'Message' not supported

var authSavePasswordFunction *gi.Function
var authSavePasswordFunction_Once sync.Once

func authSavePasswordFunction_Set() error {
	var err error
	authSavePasswordFunction_Once.Do(func() {
		err = authStruct_Set()
		if err != nil {
			return
		}
		authSavePasswordFunction, err = authStruct.InvokerNew("save_password")
	})
	return err
}

// SavePassword is a representation of the C type soup_auth_save_password.
func (recv *Auth) SavePassword(username string, password string) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(username)
	inArgs[2].SetString(password)

	err := authSavePasswordFunction_Set()
	if err == nil {
		authSavePasswordFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'soup_auth_update' : parameter 'msg' of type 'Message' not supported

var authBasicStruct *gi.Struct
var authBasicStruct_Once sync.Once

func authBasicStruct_Set() error {
	var err error
	authBasicStruct_Once.Do(func() {
		authBasicStruct, err = gi.StructNew("Soup", "AuthBasic")
	})
	return err
}

type AuthBasic struct {
	native uintptr
}

// AuthBasicStruct creates an uninitialised AuthBasic.
func AuthBasicStruct() *AuthBasic {
	err := authBasicStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &AuthBasic{native: authBasicStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeAuthBasic)
	return structGo
}
func finalizeAuthBasic(obj *AuthBasic) {
	authBasicStruct.Free(obj.native)
}

var authDigestStruct *gi.Struct
var authDigestStruct_Once sync.Once

func authDigestStruct_Set() error {
	var err error
	authDigestStruct_Once.Do(func() {
		authDigestStruct, err = gi.StructNew("Soup", "AuthDigest")
	})
	return err
}

type AuthDigest struct {
	native uintptr
}

// AuthDigestStruct creates an uninitialised AuthDigest.
func AuthDigestStruct() *AuthDigest {
	err := authDigestStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &AuthDigest{native: authDigestStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeAuthDigest)
	return structGo
}
func finalizeAuthDigest(obj *AuthDigest) {
	authDigestStruct.Free(obj.native)
}

var authDomainStruct *gi.Struct
var authDomainStruct_Once sync.Once

func authDomainStruct_Set() error {
	var err error
	authDomainStruct_Once.Do(func() {
		authDomainStruct, err = gi.StructNew("Soup", "AuthDomain")
	})
	return err
}

type AuthDomain struct {
	native uintptr
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'soup_auth_domain_accepts' : parameter 'msg' of type 'Message' not supported

var authDomainAddPathFunction *gi.Function
var authDomainAddPathFunction_Once sync.Once

func authDomainAddPathFunction_Set() error {
	var err error
	authDomainAddPathFunction_Once.Do(func() {
		err = authDomainStruct_Set()
		if err != nil {
			return
		}
		authDomainAddPathFunction, err = authDomainStruct.InvokerNew("add_path")
	})
	return err
}

// AddPath is a representation of the C type soup_auth_domain_add_path.
func (recv *AuthDomain) AddPath(path string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(path)

	err := authDomainAddPathFunction_Set()
	if err == nil {
		authDomainAddPathFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'soup_auth_domain_challenge' : parameter 'msg' of type 'Message' not supported

// UNSUPPORTED : C value 'soup_auth_domain_check_password' : parameter 'msg' of type 'Message' not supported

// UNSUPPORTED : C value 'soup_auth_domain_covers' : parameter 'msg' of type 'Message' not supported

var authDomainGetRealmFunction *gi.Function
var authDomainGetRealmFunction_Once sync.Once

func authDomainGetRealmFunction_Set() error {
	var err error
	authDomainGetRealmFunction_Once.Do(func() {
		err = authDomainStruct_Set()
		if err != nil {
			return
		}
		authDomainGetRealmFunction, err = authDomainStruct.InvokerNew("get_realm")
	})
	return err
}

// GetRealm is a representation of the C type soup_auth_domain_get_realm.
func (recv *AuthDomain) GetRealm() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

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
		err = authDomainStruct_Set()
		if err != nil {
			return
		}
		authDomainRemovePathFunction, err = authDomainStruct.InvokerNew("remove_path")
	})
	return err
}

// RemovePath is a representation of the C type soup_auth_domain_remove_path.
func (recv *AuthDomain) RemovePath(path string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(path)

	err := authDomainRemovePathFunction_Set()
	if err == nil {
		authDomainRemovePathFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'soup_auth_domain_set_filter' : parameter 'filter' of type 'AuthDomainFilter' not supported

// UNSUPPORTED : C value 'soup_auth_domain_set_generic_auth_callback' : parameter 'auth_callback' of type 'AuthDomainGenericAuthCallback' not supported

// UNSUPPORTED : C value 'soup_auth_domain_try_generic_auth_callback' : parameter 'msg' of type 'Message' not supported

// AuthDomainStruct creates an uninitialised AuthDomain.
func AuthDomainStruct() *AuthDomain {
	err := authDomainStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &AuthDomain{native: authDomainStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeAuthDomain)
	return structGo
}
func finalizeAuthDomain(obj *AuthDomain) {
	authDomainStruct.Free(obj.native)
}

var authDomainBasicStruct *gi.Struct
var authDomainBasicStruct_Once sync.Once

func authDomainBasicStruct_Set() error {
	var err error
	authDomainBasicStruct_Once.Do(func() {
		authDomainBasicStruct, err = gi.StructNew("Soup", "AuthDomainBasic")
	})
	return err
}

type AuthDomainBasic struct {
	native uintptr
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'AuthDomain'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'AuthDomain'

// UNSUPPORTED : C value 'soup_auth_domain_basic_new' : parameter '...' of type 'nil' not supported

// UNSUPPORTED : C value 'soup_auth_domain_basic_set_auth_callback' : parameter 'callback' of type 'AuthDomainBasicAuthCallback' not supported

var authDomainDigestStruct *gi.Struct
var authDomainDigestStruct_Once sync.Once

func authDomainDigestStruct_Set() error {
	var err error
	authDomainDigestStruct_Once.Do(func() {
		authDomainDigestStruct, err = gi.StructNew("Soup", "AuthDomainDigest")
	})
	return err
}

type AuthDomainDigest struct {
	native uintptr
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'AuthDomain'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'AuthDomain'

// UNSUPPORTED : C value 'soup_auth_domain_digest_new' : parameter '...' of type 'nil' not supported

// UNSUPPORTED : C value 'soup_auth_domain_digest_set_auth_callback' : parameter 'callback' of type 'AuthDomainDigestAuthCallback' not supported

var authManagerStruct *gi.Struct
var authManagerStruct_Once sync.Once

func authManagerStruct_Set() error {
	var err error
	authManagerStruct_Once.Do(func() {
		authManagerStruct, err = gi.StructNew("Soup", "AuthManager")
	})
	return err
}

type AuthManager struct {
	native uintptr
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'GObject.Object'

// FieldPriv returns the C field 'priv'.
func (recv *AuthManager) FieldPriv() *AuthManagerPrivate {
	argValue := gi.FieldGet(authManagerStruct, recv.native, "priv")
	value := &AuthManagerPrivate{native: argValue.Pointer()}
	return value
}

// SetFieldPriv sets the value of the C field 'priv'.
func (recv *AuthManager) SetFieldPriv(value *AuthManagerPrivate) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(authManagerStruct, recv.native, "priv", argValue)
}

var authManagerClearCachedCredentialsFunction *gi.Function
var authManagerClearCachedCredentialsFunction_Once sync.Once

func authManagerClearCachedCredentialsFunction_Set() error {
	var err error
	authManagerClearCachedCredentialsFunction_Once.Do(func() {
		err = authManagerStruct_Set()
		if err != nil {
			return
		}
		authManagerClearCachedCredentialsFunction, err = authManagerStruct.InvokerNew("clear_cached_credentials")
	})
	return err
}

// ClearCachedCredentials is a representation of the C type soup_auth_manager_clear_cached_credentials.
func (recv *AuthManager) ClearCachedCredentials() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := authManagerClearCachedCredentialsFunction_Set()
	if err == nil {
		authManagerClearCachedCredentialsFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'soup_auth_manager_use_auth' : parameter 'auth' of type 'Auth' not supported

// AuthManagerStruct creates an uninitialised AuthManager.
func AuthManagerStruct() *AuthManager {
	err := authManagerStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &AuthManager{native: authManagerStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeAuthManager)
	return structGo
}
func finalizeAuthManager(obj *AuthManager) {
	authManagerStruct.Free(obj.native)
}

var authNTLMStruct *gi.Struct
var authNTLMStruct_Once sync.Once

func authNTLMStruct_Set() error {
	var err error
	authNTLMStruct_Once.Do(func() {
		authNTLMStruct, err = gi.StructNew("Soup", "AuthNTLM")
	})
	return err
}

type AuthNTLM struct {
	native uintptr
}

// AuthNTLMStruct creates an uninitialised AuthNTLM.
func AuthNTLMStruct() *AuthNTLM {
	err := authNTLMStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &AuthNTLM{native: authNTLMStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeAuthNTLM)
	return structGo
}
func finalizeAuthNTLM(obj *AuthNTLM) {
	authNTLMStruct.Free(obj.native)
}

var authNegotiateStruct *gi.Struct
var authNegotiateStruct_Once sync.Once

func authNegotiateStruct_Set() error {
	var err error
	authNegotiateStruct_Once.Do(func() {
		authNegotiateStruct, err = gi.StructNew("Soup", "AuthNegotiate")
	})
	return err
}

type AuthNegotiate struct {
	native uintptr
}

// AuthNegotiateStruct creates an uninitialised AuthNegotiate.
func AuthNegotiateStruct() *AuthNegotiate {
	err := authNegotiateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &AuthNegotiate{native: authNegotiateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeAuthNegotiate)
	return structGo
}
func finalizeAuthNegotiate(obj *AuthNegotiate) {
	authNegotiateStruct.Free(obj.native)
}

var cacheStruct *gi.Struct
var cacheStruct_Once sync.Once

func cacheStruct_Set() error {
	var err error
	cacheStruct_Once.Do(func() {
		cacheStruct, err = gi.StructNew("Soup", "Cache")
	})
	return err
}

type Cache struct {
	native uintptr
}

// UNSUPPORTED : C value 'parent_instance' : for field getter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'parent_instance' : for field setter : no Go type for 'GObject.Object'

// FieldPriv returns the C field 'priv'.
func (recv *Cache) FieldPriv() *CachePrivate {
	argValue := gi.FieldGet(cacheStruct, recv.native, "priv")
	value := &CachePrivate{native: argValue.Pointer()}
	return value
}

// SetFieldPriv sets the value of the C field 'priv'.
func (recv *Cache) SetFieldPriv(value *CachePrivate) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(cacheStruct, recv.native, "priv", argValue)
}

// UNSUPPORTED : C value 'soup_cache_new' : parameter 'cache_type' of type 'CacheType' not supported

var cacheClearFunction *gi.Function
var cacheClearFunction_Once sync.Once

func cacheClearFunction_Set() error {
	var err error
	cacheClearFunction_Once.Do(func() {
		err = cacheStruct_Set()
		if err != nil {
			return
		}
		cacheClearFunction, err = cacheStruct.InvokerNew("clear")
	})
	return err
}

// Clear is a representation of the C type soup_cache_clear.
func (recv *Cache) Clear() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

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
		err = cacheStruct_Set()
		if err != nil {
			return
		}
		cacheDumpFunction, err = cacheStruct.InvokerNew("dump")
	})
	return err
}

// Dump is a representation of the C type soup_cache_dump.
func (recv *Cache) Dump() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

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
		err = cacheStruct_Set()
		if err != nil {
			return
		}
		cacheFlushFunction, err = cacheStruct.InvokerNew("flush")
	})
	return err
}

// Flush is a representation of the C type soup_cache_flush.
func (recv *Cache) Flush() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

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
		err = cacheStruct_Set()
		if err != nil {
			return
		}
		cacheGetMaxSizeFunction, err = cacheStruct.InvokerNew("get_max_size")
	})
	return err
}

// GetMaxSize is a representation of the C type soup_cache_get_max_size.
func (recv *Cache) GetMaxSize() uint32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

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
		err = cacheStruct_Set()
		if err != nil {
			return
		}
		cacheLoadFunction, err = cacheStruct.InvokerNew("load")
	})
	return err
}

// Load is a representation of the C type soup_cache_load.
func (recv *Cache) Load() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

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
		err = cacheStruct_Set()
		if err != nil {
			return
		}
		cacheSetMaxSizeFunction, err = cacheStruct.InvokerNew("set_max_size")
	})
	return err
}

// SetMaxSize is a representation of the C type soup_cache_set_max_size.
func (recv *Cache) SetMaxSize(maxSize uint32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetUint32(maxSize)

	err := cacheSetMaxSizeFunction_Set()
	if err == nil {
		cacheSetMaxSizeFunction.Invoke(inArgs[:], nil)
	}

	return
}

var contentDecoderStruct *gi.Struct
var contentDecoderStruct_Once sync.Once

func contentDecoderStruct_Set() error {
	var err error
	contentDecoderStruct_Once.Do(func() {
		contentDecoderStruct, err = gi.StructNew("Soup", "ContentDecoder")
	})
	return err
}

type ContentDecoder struct {
	native uintptr
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'GObject.Object'

// FieldPriv returns the C field 'priv'.
func (recv *ContentDecoder) FieldPriv() *ContentDecoderPrivate {
	argValue := gi.FieldGet(contentDecoderStruct, recv.native, "priv")
	value := &ContentDecoderPrivate{native: argValue.Pointer()}
	return value
}

// SetFieldPriv sets the value of the C field 'priv'.
func (recv *ContentDecoder) SetFieldPriv(value *ContentDecoderPrivate) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(contentDecoderStruct, recv.native, "priv", argValue)
}

// ContentDecoderStruct creates an uninitialised ContentDecoder.
func ContentDecoderStruct() *ContentDecoder {
	err := contentDecoderStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ContentDecoder{native: contentDecoderStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeContentDecoder)
	return structGo
}
func finalizeContentDecoder(obj *ContentDecoder) {
	contentDecoderStruct.Free(obj.native)
}

var contentSnifferStruct *gi.Struct
var contentSnifferStruct_Once sync.Once

func contentSnifferStruct_Set() error {
	var err error
	contentSnifferStruct_Once.Do(func() {
		contentSnifferStruct, err = gi.StructNew("Soup", "ContentSniffer")
	})
	return err
}

type ContentSniffer struct {
	native uintptr
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'GObject.Object'

// FieldPriv returns the C field 'priv'.
func (recv *ContentSniffer) FieldPriv() *ContentSnifferPrivate {
	argValue := gi.FieldGet(contentSnifferStruct, recv.native, "priv")
	value := &ContentSnifferPrivate{native: argValue.Pointer()}
	return value
}

// SetFieldPriv sets the value of the C field 'priv'.
func (recv *ContentSniffer) SetFieldPriv(value *ContentSnifferPrivate) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(contentSnifferStruct, recv.native, "priv", argValue)
}

// UNSUPPORTED : C value 'soup_content_sniffer_new' : return type 'ContentSniffer' not supported

var contentSnifferGetBufferSizeFunction *gi.Function
var contentSnifferGetBufferSizeFunction_Once sync.Once

func contentSnifferGetBufferSizeFunction_Set() error {
	var err error
	contentSnifferGetBufferSizeFunction_Once.Do(func() {
		err = contentSnifferStruct_Set()
		if err != nil {
			return
		}
		contentSnifferGetBufferSizeFunction, err = contentSnifferStruct.InvokerNew("get_buffer_size")
	})
	return err
}

// GetBufferSize is a representation of the C type soup_content_sniffer_get_buffer_size.
func (recv *ContentSniffer) GetBufferSize() uint64 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := contentSnifferGetBufferSizeFunction_Set()
	if err == nil {
		ret = contentSnifferGetBufferSizeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint64()

	return retGo
}

// UNSUPPORTED : C value 'soup_content_sniffer_sniff' : parameter 'msg' of type 'Message' not supported

var cookieJarStruct *gi.Struct
var cookieJarStruct_Once sync.Once

func cookieJarStruct_Set() error {
	var err error
	cookieJarStruct_Once.Do(func() {
		cookieJarStruct, err = gi.StructNew("Soup", "CookieJar")
	})
	return err
}

type CookieJar struct {
	native uintptr
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'soup_cookie_jar_new' : return type 'CookieJar' not supported

var cookieJarAddCookieFunction *gi.Function
var cookieJarAddCookieFunction_Once sync.Once

func cookieJarAddCookieFunction_Set() error {
	var err error
	cookieJarAddCookieFunction_Once.Do(func() {
		err = cookieJarStruct_Set()
		if err != nil {
			return
		}
		cookieJarAddCookieFunction, err = cookieJarStruct.InvokerNew("add_cookie")
	})
	return err
}

// AddCookie is a representation of the C type soup_cookie_jar_add_cookie.
func (recv *CookieJar) AddCookie(cookie *Cookie) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(cookie.native)

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
		err = cookieJarStruct_Set()
		if err != nil {
			return
		}
		cookieJarAddCookieFullFunction, err = cookieJarStruct.InvokerNew("add_cookie_full")
	})
	return err
}

// AddCookieFull is a representation of the C type soup_cookie_jar_add_cookie_full.
func (recv *CookieJar) AddCookieFull(cookie *Cookie, uri *URI, firstParty *URI) {
	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(cookie.native)
	inArgs[2].SetPointer(uri.native)
	inArgs[3].SetPointer(firstParty.native)

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
		err = cookieJarStruct_Set()
		if err != nil {
			return
		}
		cookieJarAddCookieWithFirstPartyFunction, err = cookieJarStruct.InvokerNew("add_cookie_with_first_party")
	})
	return err
}

// AddCookieWithFirstParty is a representation of the C type soup_cookie_jar_add_cookie_with_first_party.
func (recv *CookieJar) AddCookieWithFirstParty(firstParty *URI, cookie *Cookie) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(firstParty.native)
	inArgs[2].SetPointer(cookie.native)

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
		err = cookieJarStruct_Set()
		if err != nil {
			return
		}
		cookieJarDeleteCookieFunction, err = cookieJarStruct.InvokerNew("delete_cookie")
	})
	return err
}

// DeleteCookie is a representation of the C type soup_cookie_jar_delete_cookie.
func (recv *CookieJar) DeleteCookie(cookie *Cookie) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(cookie.native)

	err := cookieJarDeleteCookieFunction_Set()
	if err == nil {
		cookieJarDeleteCookieFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'soup_cookie_jar_get_accept_policy' : return type 'CookieJarAcceptPolicy' not supported

// UNSUPPORTED : C value 'soup_cookie_jar_get_cookie_list' : return type 'GLib.SList' not supported

var cookieJarGetCookiesFunction *gi.Function
var cookieJarGetCookiesFunction_Once sync.Once

func cookieJarGetCookiesFunction_Set() error {
	var err error
	cookieJarGetCookiesFunction_Once.Do(func() {
		err = cookieJarStruct_Set()
		if err != nil {
			return
		}
		cookieJarGetCookiesFunction, err = cookieJarStruct.InvokerNew("get_cookies")
	})
	return err
}

// GetCookies is a representation of the C type soup_cookie_jar_get_cookies.
func (recv *CookieJar) GetCookies(uri *URI, forHttp bool) string {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(uri.native)
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
		err = cookieJarStruct_Set()
		if err != nil {
			return
		}
		cookieJarIsPersistentFunction, err = cookieJarStruct.InvokerNew("is_persistent")
	})
	return err
}

// IsPersistent is a representation of the C type soup_cookie_jar_is_persistent.
func (recv *CookieJar) IsPersistent() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

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
		err = cookieJarStruct_Set()
		if err != nil {
			return
		}
		cookieJarSaveFunction, err = cookieJarStruct.InvokerNew("save")
	})
	return err
}

// Save is a representation of the C type soup_cookie_jar_save.
func (recv *CookieJar) Save() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := cookieJarSaveFunction_Set()
	if err == nil {
		cookieJarSaveFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'soup_cookie_jar_set_accept_policy' : parameter 'policy' of type 'CookieJarAcceptPolicy' not supported

var cookieJarSetCookieFunction *gi.Function
var cookieJarSetCookieFunction_Once sync.Once

func cookieJarSetCookieFunction_Set() error {
	var err error
	cookieJarSetCookieFunction_Once.Do(func() {
		err = cookieJarStruct_Set()
		if err != nil {
			return
		}
		cookieJarSetCookieFunction, err = cookieJarStruct.InvokerNew("set_cookie")
	})
	return err
}

// SetCookie is a representation of the C type soup_cookie_jar_set_cookie.
func (recv *CookieJar) SetCookie(uri *URI, cookie string) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(uri.native)
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
		err = cookieJarStruct_Set()
		if err != nil {
			return
		}
		cookieJarSetCookieWithFirstPartyFunction, err = cookieJarStruct.InvokerNew("set_cookie_with_first_party")
	})
	return err
}

// SetCookieWithFirstParty is a representation of the C type soup_cookie_jar_set_cookie_with_first_party.
func (recv *CookieJar) SetCookieWithFirstParty(uri *URI, firstParty *URI, cookie string) {
	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(uri.native)
	inArgs[2].SetPointer(firstParty.native)
	inArgs[3].SetString(cookie)

	err := cookieJarSetCookieWithFirstPartyFunction_Set()
	if err == nil {
		cookieJarSetCookieWithFirstPartyFunction.Invoke(inArgs[:], nil)
	}

	return
}

var cookieJarDBStruct *gi.Struct
var cookieJarDBStruct_Once sync.Once

func cookieJarDBStruct_Set() error {
	var err error
	cookieJarDBStruct_Once.Do(func() {
		cookieJarDBStruct, err = gi.StructNew("Soup", "CookieJarDB")
	})
	return err
}

type CookieJarDB struct {
	native uintptr
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'CookieJar'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'CookieJar'

// UNSUPPORTED : C value 'soup_cookie_jar_db_new' : return type 'CookieJarDB' not supported

var cookieJarTextStruct *gi.Struct
var cookieJarTextStruct_Once sync.Once

func cookieJarTextStruct_Set() error {
	var err error
	cookieJarTextStruct_Once.Do(func() {
		cookieJarTextStruct, err = gi.StructNew("Soup", "CookieJarText")
	})
	return err
}

type CookieJarText struct {
	native uintptr
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'CookieJar'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'CookieJar'

// UNSUPPORTED : C value 'soup_cookie_jar_text_new' : return type 'CookieJarText' not supported

var hSTSEnforcerStruct *gi.Struct
var hSTSEnforcerStruct_Once sync.Once

func hSTSEnforcerStruct_Set() error {
	var err error
	hSTSEnforcerStruct_Once.Do(func() {
		hSTSEnforcerStruct, err = gi.StructNew("Soup", "HSTSEnforcer")
	})
	return err
}

type HSTSEnforcer struct {
	native uintptr
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'GObject.Object'

// FieldPriv returns the C field 'priv'.
func (recv *HSTSEnforcer) FieldPriv() *HSTSEnforcerPrivate {
	argValue := gi.FieldGet(hSTSEnforcerStruct, recv.native, "priv")
	value := &HSTSEnforcerPrivate{native: argValue.Pointer()}
	return value
}

// SetFieldPriv sets the value of the C field 'priv'.
func (recv *HSTSEnforcer) SetFieldPriv(value *HSTSEnforcerPrivate) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(hSTSEnforcerStruct, recv.native, "priv", argValue)
}

// UNSUPPORTED : C value 'soup_hsts_enforcer_new' : return type 'HSTSEnforcer' not supported

// UNSUPPORTED : C value 'soup_hsts_enforcer_get_domains' : return type 'GLib.List' not supported

// UNSUPPORTED : C value 'soup_hsts_enforcer_get_policies' : return type 'GLib.List' not supported

var hSTSEnforcerHasValidPolicyFunction *gi.Function
var hSTSEnforcerHasValidPolicyFunction_Once sync.Once

func hSTSEnforcerHasValidPolicyFunction_Set() error {
	var err error
	hSTSEnforcerHasValidPolicyFunction_Once.Do(func() {
		err = hSTSEnforcerStruct_Set()
		if err != nil {
			return
		}
		hSTSEnforcerHasValidPolicyFunction, err = hSTSEnforcerStruct.InvokerNew("has_valid_policy")
	})
	return err
}

// HasValidPolicy is a representation of the C type soup_hsts_enforcer_has_valid_policy.
func (recv *HSTSEnforcer) HasValidPolicy(domain string) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
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
		err = hSTSEnforcerStruct_Set()
		if err != nil {
			return
		}
		hSTSEnforcerIsPersistentFunction, err = hSTSEnforcerStruct.InvokerNew("is_persistent")
	})
	return err
}

// IsPersistent is a representation of the C type soup_hsts_enforcer_is_persistent.
func (recv *HSTSEnforcer) IsPersistent() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

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
		err = hSTSEnforcerStruct_Set()
		if err != nil {
			return
		}
		hSTSEnforcerSetPolicyFunction, err = hSTSEnforcerStruct.InvokerNew("set_policy")
	})
	return err
}

// SetPolicy is a representation of the C type soup_hsts_enforcer_set_policy.
func (recv *HSTSEnforcer) SetPolicy(policy *HSTSPolicy) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(policy.native)

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
		err = hSTSEnforcerStruct_Set()
		if err != nil {
			return
		}
		hSTSEnforcerSetSessionPolicyFunction, err = hSTSEnforcerStruct.InvokerNew("set_session_policy")
	})
	return err
}

// SetSessionPolicy is a representation of the C type soup_hsts_enforcer_set_session_policy.
func (recv *HSTSEnforcer) SetSessionPolicy(domain string, includeSubdomains bool) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(domain)
	inArgs[2].SetBoolean(includeSubdomains)

	err := hSTSEnforcerSetSessionPolicyFunction_Set()
	if err == nil {
		hSTSEnforcerSetSessionPolicyFunction.Invoke(inArgs[:], nil)
	}

	return
}

var hSTSEnforcerDBStruct *gi.Struct
var hSTSEnforcerDBStruct_Once sync.Once

func hSTSEnforcerDBStruct_Set() error {
	var err error
	hSTSEnforcerDBStruct_Once.Do(func() {
		hSTSEnforcerDBStruct, err = gi.StructNew("Soup", "HSTSEnforcerDB")
	})
	return err
}

type HSTSEnforcerDB struct {
	native uintptr
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'HSTSEnforcer'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'HSTSEnforcer'

// FieldPriv returns the C field 'priv'.
func (recv *HSTSEnforcerDB) FieldPriv() *HSTSEnforcerDBPrivate {
	argValue := gi.FieldGet(hSTSEnforcerDBStruct, recv.native, "priv")
	value := &HSTSEnforcerDBPrivate{native: argValue.Pointer()}
	return value
}

// SetFieldPriv sets the value of the C field 'priv'.
func (recv *HSTSEnforcerDB) SetFieldPriv(value *HSTSEnforcerDBPrivate) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(hSTSEnforcerDBStruct, recv.native, "priv", argValue)
}

// UNSUPPORTED : C value 'soup_hsts_enforcer_db_new' : return type 'HSTSEnforcerDB' not supported

var loggerStruct *gi.Struct
var loggerStruct_Once sync.Once

func loggerStruct_Set() error {
	var err error
	loggerStruct_Once.Do(func() {
		loggerStruct, err = gi.StructNew("Soup", "Logger")
	})
	return err
}

type Logger struct {
	native uintptr
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'soup_logger_new' : parameter 'level' of type 'LoggerLogLevel' not supported

// UNSUPPORTED : C value 'soup_logger_attach' : parameter 'session' of type 'Session' not supported

// UNSUPPORTED : C value 'soup_logger_detach' : parameter 'session' of type 'Session' not supported

// UNSUPPORTED : C value 'soup_logger_set_printer' : parameter 'printer' of type 'LoggerPrinter' not supported

// UNSUPPORTED : C value 'soup_logger_set_request_filter' : parameter 'request_filter' of type 'LoggerFilter' not supported

// UNSUPPORTED : C value 'soup_logger_set_response_filter' : parameter 'response_filter' of type 'LoggerFilter' not supported

var messageStruct *gi.Struct
var messageStruct_Once sync.Once

func messageStruct_Set() error {
	var err error
	messageStruct_Once.Do(func() {
		messageStruct, err = gi.StructNew("Soup", "Message")
	})
	return err
}

type Message struct {
	native uintptr
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'GObject.Object'

// FieldMethod returns the C field 'method'.
func (recv *Message) FieldMethod() string {
	argValue := gi.FieldGet(messageStruct, recv.native, "method")
	value := argValue.String(false)
	return value
}

// SetFieldMethod sets the value of the C field 'method'.
func (recv *Message) SetFieldMethod(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.FieldSet(messageStruct, recv.native, "method", argValue)
}

// FieldStatusCode returns the C field 'status_code'.
func (recv *Message) FieldStatusCode() uint32 {
	argValue := gi.FieldGet(messageStruct, recv.native, "status_code")
	value := argValue.Uint32()
	return value
}

// SetFieldStatusCode sets the value of the C field 'status_code'.
func (recv *Message) SetFieldStatusCode(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.FieldSet(messageStruct, recv.native, "status_code", argValue)
}

// FieldReasonPhrase returns the C field 'reason_phrase'.
func (recv *Message) FieldReasonPhrase() string {
	argValue := gi.FieldGet(messageStruct, recv.native, "reason_phrase")
	value := argValue.String(false)
	return value
}

// SetFieldReasonPhrase sets the value of the C field 'reason_phrase'.
func (recv *Message) SetFieldReasonPhrase(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.FieldSet(messageStruct, recv.native, "reason_phrase", argValue)
}

// FieldRequestBody returns the C field 'request_body'.
func (recv *Message) FieldRequestBody() *MessageBody {
	argValue := gi.FieldGet(messageStruct, recv.native, "request_body")
	value := &MessageBody{native: argValue.Pointer()}
	return value
}

// SetFieldRequestBody sets the value of the C field 'request_body'.
func (recv *Message) SetFieldRequestBody(value *MessageBody) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(messageStruct, recv.native, "request_body", argValue)
}

// FieldRequestHeaders returns the C field 'request_headers'.
func (recv *Message) FieldRequestHeaders() *MessageHeaders {
	argValue := gi.FieldGet(messageStruct, recv.native, "request_headers")
	value := &MessageHeaders{native: argValue.Pointer()}
	return value
}

// SetFieldRequestHeaders sets the value of the C field 'request_headers'.
func (recv *Message) SetFieldRequestHeaders(value *MessageHeaders) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(messageStruct, recv.native, "request_headers", argValue)
}

// FieldResponseBody returns the C field 'response_body'.
func (recv *Message) FieldResponseBody() *MessageBody {
	argValue := gi.FieldGet(messageStruct, recv.native, "response_body")
	value := &MessageBody{native: argValue.Pointer()}
	return value
}

// SetFieldResponseBody sets the value of the C field 'response_body'.
func (recv *Message) SetFieldResponseBody(value *MessageBody) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(messageStruct, recv.native, "response_body", argValue)
}

// FieldResponseHeaders returns the C field 'response_headers'.
func (recv *Message) FieldResponseHeaders() *MessageHeaders {
	argValue := gi.FieldGet(messageStruct, recv.native, "response_headers")
	value := &MessageHeaders{native: argValue.Pointer()}
	return value
}

// SetFieldResponseHeaders sets the value of the C field 'response_headers'.
func (recv *Message) SetFieldResponseHeaders(value *MessageHeaders) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(messageStruct, recv.native, "response_headers", argValue)
}

// UNSUPPORTED : C value 'soup_message_new' : return type 'Message' not supported

// UNSUPPORTED : C value 'soup_message_new_from_uri' : return type 'Message' not supported

// UNSUPPORTED : C value 'soup_message_add_header_handler' : parameter 'callback' of type 'GObject.Callback' not supported

// UNSUPPORTED : C value 'soup_message_add_status_code_handler' : parameter 'callback' of type 'GObject.Callback' not supported

// UNSUPPORTED : C value 'soup_message_content_sniffed' : parameter 'params' of type 'GLib.HashTable' not supported

// UNSUPPORTED : C value 'soup_message_disable_feature' : parameter 'feature_type' of type 'GType' not supported

var messageFinishedFunction *gi.Function
var messageFinishedFunction_Once sync.Once

func messageFinishedFunction_Set() error {
	var err error
	messageFinishedFunction_Once.Do(func() {
		err = messageStruct_Set()
		if err != nil {
			return
		}
		messageFinishedFunction, err = messageStruct.InvokerNew("finished")
	})
	return err
}

// Finished is a representation of the C type soup_message_finished.
func (recv *Message) Finished() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := messageFinishedFunction_Set()
	if err == nil {
		messageFinishedFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'soup_message_get_address' : return type 'Address' not supported

var messageGetFirstPartyFunction *gi.Function
var messageGetFirstPartyFunction_Once sync.Once

func messageGetFirstPartyFunction_Set() error {
	var err error
	messageGetFirstPartyFunction_Once.Do(func() {
		err = messageStruct_Set()
		if err != nil {
			return
		}
		messageGetFirstPartyFunction, err = messageStruct.InvokerNew("get_first_party")
	})
	return err
}

// GetFirstParty is a representation of the C type soup_message_get_first_party.
func (recv *Message) GetFirstParty() *URI {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := messageGetFirstPartyFunction_Set()
	if err == nil {
		ret = messageGetFirstPartyFunction.Invoke(inArgs[:], nil)
	}

	retGo := &URI{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'soup_message_get_flags' : return type 'MessageFlags' not supported

// UNSUPPORTED : C value 'soup_message_get_http_version' : return type 'HTTPVersion' not supported

// UNSUPPORTED : C value 'soup_message_get_https_status' : parameter 'certificate' of type 'Gio.TlsCertificate' not supported

// UNSUPPORTED : C value 'soup_message_get_priority' : return type 'MessagePriority' not supported

// UNSUPPORTED : C value 'soup_message_get_soup_request' : return type 'Request' not supported

var messageGetUriFunction *gi.Function
var messageGetUriFunction_Once sync.Once

func messageGetUriFunction_Set() error {
	var err error
	messageGetUriFunction_Once.Do(func() {
		err = messageStruct_Set()
		if err != nil {
			return
		}
		messageGetUriFunction, err = messageStruct.InvokerNew("get_uri")
	})
	return err
}

// GetUri is a representation of the C type soup_message_get_uri.
func (recv *Message) GetUri() *URI {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := messageGetUriFunction_Set()
	if err == nil {
		ret = messageGetUriFunction.Invoke(inArgs[:], nil)
	}

	retGo := &URI{native: ret.Pointer()}

	return retGo
}

var messageGotBodyFunction *gi.Function
var messageGotBodyFunction_Once sync.Once

func messageGotBodyFunction_Set() error {
	var err error
	messageGotBodyFunction_Once.Do(func() {
		err = messageStruct_Set()
		if err != nil {
			return
		}
		messageGotBodyFunction, err = messageStruct.InvokerNew("got_body")
	})
	return err
}

// GotBody is a representation of the C type soup_message_got_body.
func (recv *Message) GotBody() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

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
		err = messageStruct_Set()
		if err != nil {
			return
		}
		messageGotChunkFunction, err = messageStruct.InvokerNew("got_chunk")
	})
	return err
}

// GotChunk is a representation of the C type soup_message_got_chunk.
func (recv *Message) GotChunk(chunk *Buffer) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(chunk.native)

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
		err = messageStruct_Set()
		if err != nil {
			return
		}
		messageGotHeadersFunction, err = messageStruct.InvokerNew("got_headers")
	})
	return err
}

// GotHeaders is a representation of the C type soup_message_got_headers.
func (recv *Message) GotHeaders() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

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
		err = messageStruct_Set()
		if err != nil {
			return
		}
		messageGotInformationalFunction, err = messageStruct.InvokerNew("got_informational")
	})
	return err
}

// GotInformational is a representation of the C type soup_message_got_informational.
func (recv *Message) GotInformational() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

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
		err = messageStruct_Set()
		if err != nil {
			return
		}
		messageIsKeepaliveFunction, err = messageStruct.InvokerNew("is_keepalive")
	})
	return err
}

// IsKeepalive is a representation of the C type soup_message_is_keepalive.
func (recv *Message) IsKeepalive() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

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
		err = messageStruct_Set()
		if err != nil {
			return
		}
		messageRestartedFunction, err = messageStruct.InvokerNew("restarted")
	})
	return err
}

// Restarted is a representation of the C type soup_message_restarted.
func (recv *Message) Restarted() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

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
		err = messageStruct_Set()
		if err != nil {
			return
		}
		messageSetFirstPartyFunction, err = messageStruct.InvokerNew("set_first_party")
	})
	return err
}

// SetFirstParty is a representation of the C type soup_message_set_first_party.
func (recv *Message) SetFirstParty(firstParty *URI) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(firstParty.native)

	err := messageSetFirstPartyFunction_Set()
	if err == nil {
		messageSetFirstPartyFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'soup_message_set_flags' : parameter 'flags' of type 'MessageFlags' not supported

// UNSUPPORTED : C value 'soup_message_set_http_version' : parameter 'version' of type 'HTTPVersion' not supported

// UNSUPPORTED : C value 'soup_message_set_priority' : parameter 'priority' of type 'MessagePriority' not supported

var messageSetRedirectFunction *gi.Function
var messageSetRedirectFunction_Once sync.Once

func messageSetRedirectFunction_Set() error {
	var err error
	messageSetRedirectFunction_Once.Do(func() {
		err = messageStruct_Set()
		if err != nil {
			return
		}
		messageSetRedirectFunction, err = messageStruct.InvokerNew("set_redirect")
	})
	return err
}

// SetRedirect is a representation of the C type soup_message_set_redirect.
func (recv *Message) SetRedirect(statusCode uint32, redirectUri string) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetUint32(statusCode)
	inArgs[2].SetString(redirectUri)

	err := messageSetRedirectFunction_Set()
	if err == nil {
		messageSetRedirectFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'soup_message_set_request' : parameter 'req_use' of type 'MemoryUse' not supported

// UNSUPPORTED : C value 'soup_message_set_response' : parameter 'resp_use' of type 'MemoryUse' not supported

var messageSetStatusFunction *gi.Function
var messageSetStatusFunction_Once sync.Once

func messageSetStatusFunction_Set() error {
	var err error
	messageSetStatusFunction_Once.Do(func() {
		err = messageStruct_Set()
		if err != nil {
			return
		}
		messageSetStatusFunction, err = messageStruct.InvokerNew("set_status")
	})
	return err
}

// SetStatus is a representation of the C type soup_message_set_status.
func (recv *Message) SetStatus(statusCode uint32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
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
		err = messageStruct_Set()
		if err != nil {
			return
		}
		messageSetStatusFullFunction, err = messageStruct.InvokerNew("set_status_full")
	})
	return err
}

// SetStatusFull is a representation of the C type soup_message_set_status_full.
func (recv *Message) SetStatusFull(statusCode uint32, reasonPhrase string) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
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
		err = messageStruct_Set()
		if err != nil {
			return
		}
		messageSetUriFunction, err = messageStruct.InvokerNew("set_uri")
	})
	return err
}

// SetUri is a representation of the C type soup_message_set_uri.
func (recv *Message) SetUri(uri *URI) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(uri.native)

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
		err = messageStruct_Set()
		if err != nil {
			return
		}
		messageStartingFunction, err = messageStruct.InvokerNew("starting")
	})
	return err
}

// Starting is a representation of the C type soup_message_starting.
func (recv *Message) Starting() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

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
		err = messageStruct_Set()
		if err != nil {
			return
		}
		messageWroteBodyFunction, err = messageStruct.InvokerNew("wrote_body")
	})
	return err
}

// WroteBody is a representation of the C type soup_message_wrote_body.
func (recv *Message) WroteBody() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

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
		err = messageStruct_Set()
		if err != nil {
			return
		}
		messageWroteBodyDataFunction, err = messageStruct.InvokerNew("wrote_body_data")
	})
	return err
}

// WroteBodyData is a representation of the C type soup_message_wrote_body_data.
func (recv *Message) WroteBodyData(chunk *Buffer) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(chunk.native)

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
		err = messageStruct_Set()
		if err != nil {
			return
		}
		messageWroteChunkFunction, err = messageStruct.InvokerNew("wrote_chunk")
	})
	return err
}

// WroteChunk is a representation of the C type soup_message_wrote_chunk.
func (recv *Message) WroteChunk() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

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
		err = messageStruct_Set()
		if err != nil {
			return
		}
		messageWroteHeadersFunction, err = messageStruct.InvokerNew("wrote_headers")
	})
	return err
}

// WroteHeaders is a representation of the C type soup_message_wrote_headers.
func (recv *Message) WroteHeaders() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

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
		err = messageStruct_Set()
		if err != nil {
			return
		}
		messageWroteInformationalFunction, err = messageStruct.InvokerNew("wrote_informational")
	})
	return err
}

// WroteInformational is a representation of the C type soup_message_wrote_informational.
func (recv *Message) WroteInformational() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := messageWroteInformationalFunction_Set()
	if err == nil {
		messageWroteInformationalFunction.Invoke(inArgs[:], nil)
	}

	return
}

var multipartInputStreamStruct *gi.Struct
var multipartInputStreamStruct_Once sync.Once

func multipartInputStreamStruct_Set() error {
	var err error
	multipartInputStreamStruct_Once.Do(func() {
		multipartInputStreamStruct, err = gi.StructNew("Soup", "MultipartInputStream")
	})
	return err
}

type MultipartInputStream struct {
	native uintptr
}

// UNSUPPORTED : C value 'parent_instance' : for field getter : no Go type for 'Gio.FilterInputStream'

// UNSUPPORTED : C value 'parent_instance' : for field setter : no Go type for 'Gio.FilterInputStream'

// UNSUPPORTED : C value 'soup_multipart_input_stream_new' : parameter 'msg' of type 'Message' not supported

var multipartInputStreamGetHeadersFunction *gi.Function
var multipartInputStreamGetHeadersFunction_Once sync.Once

func multipartInputStreamGetHeadersFunction_Set() error {
	var err error
	multipartInputStreamGetHeadersFunction_Once.Do(func() {
		err = multipartInputStreamStruct_Set()
		if err != nil {
			return
		}
		multipartInputStreamGetHeadersFunction, err = multipartInputStreamStruct.InvokerNew("get_headers")
	})
	return err
}

// GetHeaders is a representation of the C type soup_multipart_input_stream_get_headers.
func (recv *MultipartInputStream) GetHeaders() *MessageHeaders {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := multipartInputStreamGetHeadersFunction_Set()
	if err == nil {
		ret = multipartInputStreamGetHeadersFunction.Invoke(inArgs[:], nil)
	}

	retGo := &MessageHeaders{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'soup_multipart_input_stream_next_part' : parameter 'cancellable' of type 'Gio.Cancellable' not supported

// UNSUPPORTED : C value 'soup_multipart_input_stream_next_part_async' : parameter 'cancellable' of type 'Gio.Cancellable' not supported

// UNSUPPORTED : C value 'soup_multipart_input_stream_next_part_finish' : parameter 'result' of type 'Gio.AsyncResult' not supported

var proxyResolverDefaultStruct *gi.Struct
var proxyResolverDefaultStruct_Once sync.Once

func proxyResolverDefaultStruct_Set() error {
	var err error
	proxyResolverDefaultStruct_Once.Do(func() {
		proxyResolverDefaultStruct, err = gi.StructNew("Soup", "ProxyResolverDefault")
	})
	return err
}

type ProxyResolverDefault struct {
	native uintptr
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'GObject.Object'

// ProxyResolverDefaultStruct creates an uninitialised ProxyResolverDefault.
func ProxyResolverDefaultStruct() *ProxyResolverDefault {
	err := proxyResolverDefaultStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ProxyResolverDefault{native: proxyResolverDefaultStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeProxyResolverDefault)
	return structGo
}
func finalizeProxyResolverDefault(obj *ProxyResolverDefault) {
	proxyResolverDefaultStruct.Free(obj.native)
}

var requestStruct *gi.Struct
var requestStruct_Once sync.Once

func requestStruct_Set() error {
	var err error
	requestStruct_Once.Do(func() {
		requestStruct, err = gi.StructNew("Soup", "Request")
	})
	return err
}

type Request struct {
	native uintptr
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'GObject.Object'

// FieldPriv returns the C field 'priv'.
func (recv *Request) FieldPriv() *RequestPrivate {
	argValue := gi.FieldGet(requestStruct, recv.native, "priv")
	value := &RequestPrivate{native: argValue.Pointer()}
	return value
}

// SetFieldPriv sets the value of the C field 'priv'.
func (recv *Request) SetFieldPriv(value *RequestPrivate) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(requestStruct, recv.native, "priv", argValue)
}

var requestGetContentLengthFunction *gi.Function
var requestGetContentLengthFunction_Once sync.Once

func requestGetContentLengthFunction_Set() error {
	var err error
	requestGetContentLengthFunction_Once.Do(func() {
		err = requestStruct_Set()
		if err != nil {
			return
		}
		requestGetContentLengthFunction, err = requestStruct.InvokerNew("get_content_length")
	})
	return err
}

// GetContentLength is a representation of the C type soup_request_get_content_length.
func (recv *Request) GetContentLength() int64 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

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
		err = requestStruct_Set()
		if err != nil {
			return
		}
		requestGetContentTypeFunction, err = requestStruct.InvokerNew("get_content_type")
	})
	return err
}

// GetContentType is a representation of the C type soup_request_get_content_type.
func (recv *Request) GetContentType() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := requestGetContentTypeFunction_Set()
	if err == nil {
		ret = requestGetContentTypeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'soup_request_get_session' : return type 'Session' not supported

var requestGetUriFunction *gi.Function
var requestGetUriFunction_Once sync.Once

func requestGetUriFunction_Set() error {
	var err error
	requestGetUriFunction_Once.Do(func() {
		err = requestStruct_Set()
		if err != nil {
			return
		}
		requestGetUriFunction, err = requestStruct.InvokerNew("get_uri")
	})
	return err
}

// GetUri is a representation of the C type soup_request_get_uri.
func (recv *Request) GetUri() *URI {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := requestGetUriFunction_Set()
	if err == nil {
		ret = requestGetUriFunction.Invoke(inArgs[:], nil)
	}

	retGo := &URI{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'soup_request_send' : parameter 'cancellable' of type 'Gio.Cancellable' not supported

// UNSUPPORTED : C value 'soup_request_send_async' : parameter 'cancellable' of type 'Gio.Cancellable' not supported

// UNSUPPORTED : C value 'soup_request_send_finish' : parameter 'result' of type 'Gio.AsyncResult' not supported

// RequestStruct creates an uninitialised Request.
func RequestStruct() *Request {
	err := requestStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &Request{native: requestStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeRequest)
	return structGo
}
func finalizeRequest(obj *Request) {
	requestStruct.Free(obj.native)
}

var requestDataStruct *gi.Struct
var requestDataStruct_Once sync.Once

func requestDataStruct_Set() error {
	var err error
	requestDataStruct_Once.Do(func() {
		requestDataStruct, err = gi.StructNew("Soup", "RequestData")
	})
	return err
}

type RequestData struct {
	native uintptr
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'Request'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'Request'

// FieldPriv returns the C field 'priv'.
func (recv *RequestData) FieldPriv() *RequestDataPrivate {
	argValue := gi.FieldGet(requestDataStruct, recv.native, "priv")
	value := &RequestDataPrivate{native: argValue.Pointer()}
	return value
}

// SetFieldPriv sets the value of the C field 'priv'.
func (recv *RequestData) SetFieldPriv(value *RequestDataPrivate) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(requestDataStruct, recv.native, "priv", argValue)
}

// RequestDataStruct creates an uninitialised RequestData.
func RequestDataStruct() *RequestData {
	err := requestDataStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &RequestData{native: requestDataStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeRequestData)
	return structGo
}
func finalizeRequestData(obj *RequestData) {
	requestDataStruct.Free(obj.native)
}

var requestFileStruct *gi.Struct
var requestFileStruct_Once sync.Once

func requestFileStruct_Set() error {
	var err error
	requestFileStruct_Once.Do(func() {
		requestFileStruct, err = gi.StructNew("Soup", "RequestFile")
	})
	return err
}

type RequestFile struct {
	native uintptr
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'Request'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'Request'

// FieldPriv returns the C field 'priv'.
func (recv *RequestFile) FieldPriv() *RequestFilePrivate {
	argValue := gi.FieldGet(requestFileStruct, recv.native, "priv")
	value := &RequestFilePrivate{native: argValue.Pointer()}
	return value
}

// SetFieldPriv sets the value of the C field 'priv'.
func (recv *RequestFile) SetFieldPriv(value *RequestFilePrivate) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(requestFileStruct, recv.native, "priv", argValue)
}

// UNSUPPORTED : C value 'soup_request_file_get_file' : return type 'Gio.File' not supported

// RequestFileStruct creates an uninitialised RequestFile.
func RequestFileStruct() *RequestFile {
	err := requestFileStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &RequestFile{native: requestFileStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeRequestFile)
	return structGo
}
func finalizeRequestFile(obj *RequestFile) {
	requestFileStruct.Free(obj.native)
}

var requestHTTPStruct *gi.Struct
var requestHTTPStruct_Once sync.Once

func requestHTTPStruct_Set() error {
	var err error
	requestHTTPStruct_Once.Do(func() {
		requestHTTPStruct, err = gi.StructNew("Soup", "RequestHTTP")
	})
	return err
}

type RequestHTTP struct {
	native uintptr
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'Request'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'Request'

// FieldPriv returns the C field 'priv'.
func (recv *RequestHTTP) FieldPriv() *RequestHTTPPrivate {
	argValue := gi.FieldGet(requestHTTPStruct, recv.native, "priv")
	value := &RequestHTTPPrivate{native: argValue.Pointer()}
	return value
}

// SetFieldPriv sets the value of the C field 'priv'.
func (recv *RequestHTTP) SetFieldPriv(value *RequestHTTPPrivate) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(requestHTTPStruct, recv.native, "priv", argValue)
}

// UNSUPPORTED : C value 'soup_request_http_get_message' : return type 'Message' not supported

// RequestHTTPStruct creates an uninitialised RequestHTTP.
func RequestHTTPStruct() *RequestHTTP {
	err := requestHTTPStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &RequestHTTP{native: requestHTTPStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeRequestHTTP)
	return structGo
}
func finalizeRequestHTTP(obj *RequestHTTP) {
	requestHTTPStruct.Free(obj.native)
}

var requesterStruct *gi.Struct
var requesterStruct_Once sync.Once

func requesterStruct_Set() error {
	var err error
	requesterStruct_Once.Do(func() {
		requesterStruct, err = gi.StructNew("Soup", "Requester")
	})
	return err
}

type Requester struct {
	native uintptr
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'GObject.Object'

// FieldPriv returns the C field 'priv'.
func (recv *Requester) FieldPriv() *RequesterPrivate {
	argValue := gi.FieldGet(requesterStruct, recv.native, "priv")
	value := &RequesterPrivate{native: argValue.Pointer()}
	return value
}

// SetFieldPriv sets the value of the C field 'priv'.
func (recv *Requester) SetFieldPriv(value *RequesterPrivate) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(requesterStruct, recv.native, "priv", argValue)
}

// UNSUPPORTED : C value 'soup_requester_new' : return type 'Requester' not supported

// UNSUPPORTED : C value 'soup_requester_request' : return type 'Request' not supported

// UNSUPPORTED : C value 'soup_requester_request_uri' : return type 'Request' not supported

var serverStruct *gi.Struct
var serverStruct_Once sync.Once

func serverStruct_Set() error {
	var err error
	serverStruct_Once.Do(func() {
		serverStruct, err = gi.StructNew("Soup", "Server")
	})
	return err
}

type Server struct {
	native uintptr
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'soup_server_new' : parameter '...' of type 'nil' not supported

// UNSUPPORTED : C value 'soup_server_accept_iostream' : parameter 'stream' of type 'Gio.IOStream' not supported

// UNSUPPORTED : C value 'soup_server_add_auth_domain' : parameter 'auth_domain' of type 'AuthDomain' not supported

// UNSUPPORTED : C value 'soup_server_add_early_handler' : parameter 'callback' of type 'ServerCallback' not supported

// UNSUPPORTED : C value 'soup_server_add_handler' : parameter 'callback' of type 'ServerCallback' not supported

// UNSUPPORTED : C value 'soup_server_add_websocket_extension' : parameter 'extension_type' of type 'GType' not supported

// UNSUPPORTED : C value 'soup_server_add_websocket_handler' : parameter 'protocols' of type 'nil' not supported

var serverDisconnectFunction *gi.Function
var serverDisconnectFunction_Once sync.Once

func serverDisconnectFunction_Set() error {
	var err error
	serverDisconnectFunction_Once.Do(func() {
		err = serverStruct_Set()
		if err != nil {
			return
		}
		serverDisconnectFunction, err = serverStruct.InvokerNew("disconnect")
	})
	return err
}

// Disconnect is a representation of the C type soup_server_disconnect.
func (recv *Server) Disconnect() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := serverDisconnectFunction_Set()
	if err == nil {
		serverDisconnectFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'soup_server_get_async_context' : return type 'GLib.MainContext' not supported

// UNSUPPORTED : C value 'soup_server_get_listener' : return type 'Socket' not supported

// UNSUPPORTED : C value 'soup_server_get_listeners' : return type 'GLib.SList' not supported

var serverGetPortFunction *gi.Function
var serverGetPortFunction_Once sync.Once

func serverGetPortFunction_Set() error {
	var err error
	serverGetPortFunction_Once.Do(func() {
		err = serverStruct_Set()
		if err != nil {
			return
		}
		serverGetPortFunction, err = serverStruct.InvokerNew("get_port")
	})
	return err
}

// GetPort is a representation of the C type soup_server_get_port.
func (recv *Server) GetPort() uint32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

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
		err = serverStruct_Set()
		if err != nil {
			return
		}
		serverIsHttpsFunction, err = serverStruct.InvokerNew("is_https")
	})
	return err
}

// IsHttps is a representation of the C type soup_server_is_https.
func (recv *Server) IsHttps() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

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

// UNSUPPORTED : C value 'soup_server_pause_message' : parameter 'msg' of type 'Message' not supported

var serverQuitFunction *gi.Function
var serverQuitFunction_Once sync.Once

func serverQuitFunction_Set() error {
	var err error
	serverQuitFunction_Once.Do(func() {
		err = serverStruct_Set()
		if err != nil {
			return
		}
		serverQuitFunction, err = serverStruct.InvokerNew("quit")
	})
	return err
}

// Quit is a representation of the C type soup_server_quit.
func (recv *Server) Quit() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := serverQuitFunction_Set()
	if err == nil {
		serverQuitFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'soup_server_remove_auth_domain' : parameter 'auth_domain' of type 'AuthDomain' not supported

var serverRemoveHandlerFunction *gi.Function
var serverRemoveHandlerFunction_Once sync.Once

func serverRemoveHandlerFunction_Set() error {
	var err error
	serverRemoveHandlerFunction_Once.Do(func() {
		err = serverStruct_Set()
		if err != nil {
			return
		}
		serverRemoveHandlerFunction, err = serverStruct.InvokerNew("remove_handler")
	})
	return err
}

// RemoveHandler is a representation of the C type soup_server_remove_handler.
func (recv *Server) RemoveHandler(path string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
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
		err = serverStruct_Set()
		if err != nil {
			return
		}
		serverRunFunction, err = serverStruct.InvokerNew("run")
	})
	return err
}

// Run is a representation of the C type soup_server_run.
func (recv *Server) Run() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

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
		err = serverStruct_Set()
		if err != nil {
			return
		}
		serverRunAsyncFunction, err = serverStruct.InvokerNew("run_async")
	})
	return err
}

// RunAsync is a representation of the C type soup_server_run_async.
func (recv *Server) RunAsync() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

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
		err = serverStruct_Set()
		if err != nil {
			return
		}
		serverSetSslCertFileFunction, err = serverStruct.InvokerNew("set_ssl_cert_file")
	})
	return err
}

// SetSslCertFile is a representation of the C type soup_server_set_ssl_cert_file.
func (recv *Server) SetSslCertFile(sslCertFile string, sslKeyFile string) bool {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
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

// UNSUPPORTED : C value 'soup_server_unpause_message' : parameter 'msg' of type 'Message' not supported

var sessionStruct *gi.Struct
var sessionStruct_Once sync.Once

func sessionStruct_Set() error {
	var err error
	sessionStruct_Once.Do(func() {
		sessionStruct, err = gi.StructNew("Soup", "Session")
	})
	return err
}

type Session struct {
	native uintptr
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'soup_session_new' : return type 'Session' not supported

// UNSUPPORTED : C value 'soup_session_new_with_options' : parameter '...' of type 'nil' not supported

var sessionAbortFunction *gi.Function
var sessionAbortFunction_Once sync.Once

func sessionAbortFunction_Set() error {
	var err error
	sessionAbortFunction_Once.Do(func() {
		err = sessionStruct_Set()
		if err != nil {
			return
		}
		sessionAbortFunction, err = sessionStruct.InvokerNew("abort")
	})
	return err
}

// Abort is a representation of the C type soup_session_abort.
func (recv *Session) Abort() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := sessionAbortFunction_Set()
	if err == nil {
		sessionAbortFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'soup_session_add_feature' : parameter 'feature' of type 'SessionFeature' not supported

// UNSUPPORTED : C value 'soup_session_add_feature_by_type' : parameter 'feature_type' of type 'GType' not supported

// UNSUPPORTED : C value 'soup_session_cancel_message' : parameter 'msg' of type 'Message' not supported

// UNSUPPORTED : C value 'soup_session_connect_async' : parameter 'cancellable' of type 'Gio.Cancellable' not supported

// UNSUPPORTED : C value 'soup_session_connect_finish' : parameter 'result' of type 'Gio.AsyncResult' not supported

// UNSUPPORTED : C value 'soup_session_get_async_context' : return type 'GLib.MainContext' not supported

// UNSUPPORTED : C value 'soup_session_get_feature' : parameter 'feature_type' of type 'GType' not supported

// UNSUPPORTED : C value 'soup_session_get_feature_for_message' : parameter 'feature_type' of type 'GType' not supported

// UNSUPPORTED : C value 'soup_session_get_features' : parameter 'feature_type' of type 'GType' not supported

// UNSUPPORTED : C value 'soup_session_has_feature' : parameter 'feature_type' of type 'GType' not supported

// UNSUPPORTED : C value 'soup_session_pause_message' : parameter 'msg' of type 'Message' not supported

// UNSUPPORTED : C value 'soup_session_prefetch_dns' : parameter 'cancellable' of type 'Gio.Cancellable' not supported

var sessionPrepareForUriFunction *gi.Function
var sessionPrepareForUriFunction_Once sync.Once

func sessionPrepareForUriFunction_Set() error {
	var err error
	sessionPrepareForUriFunction_Once.Do(func() {
		err = sessionStruct_Set()
		if err != nil {
			return
		}
		sessionPrepareForUriFunction, err = sessionStruct.InvokerNew("prepare_for_uri")
	})
	return err
}

// PrepareForUri is a representation of the C type soup_session_prepare_for_uri.
func (recv *Session) PrepareForUri(uri *URI) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(uri.native)

	err := sessionPrepareForUriFunction_Set()
	if err == nil {
		sessionPrepareForUriFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'soup_session_queue_message' : parameter 'msg' of type 'Message' not supported

// UNSUPPORTED : C value 'soup_session_redirect_message' : parameter 'msg' of type 'Message' not supported

// UNSUPPORTED : C value 'soup_session_remove_feature' : parameter 'feature' of type 'SessionFeature' not supported

// UNSUPPORTED : C value 'soup_session_remove_feature_by_type' : parameter 'feature_type' of type 'GType' not supported

// UNSUPPORTED : C value 'soup_session_request' : return type 'Request' not supported

// UNSUPPORTED : C value 'soup_session_request_http' : return type 'RequestHTTP' not supported

// UNSUPPORTED : C value 'soup_session_request_http_uri' : return type 'RequestHTTP' not supported

// UNSUPPORTED : C value 'soup_session_request_uri' : return type 'Request' not supported

// UNSUPPORTED : C value 'soup_session_requeue_message' : parameter 'msg' of type 'Message' not supported

// UNSUPPORTED : C value 'soup_session_send' : parameter 'msg' of type 'Message' not supported

// UNSUPPORTED : C value 'soup_session_send_async' : parameter 'msg' of type 'Message' not supported

// UNSUPPORTED : C value 'soup_session_send_finish' : parameter 'result' of type 'Gio.AsyncResult' not supported

// UNSUPPORTED : C value 'soup_session_send_message' : parameter 'msg' of type 'Message' not supported

// UNSUPPORTED : C value 'soup_session_steal_connection' : parameter 'msg' of type 'Message' not supported

// UNSUPPORTED : C value 'soup_session_unpause_message' : parameter 'msg' of type 'Message' not supported

// UNSUPPORTED : C value 'soup_session_websocket_connect_async' : parameter 'msg' of type 'Message' not supported

// UNSUPPORTED : C value 'soup_session_websocket_connect_finish' : parameter 'result' of type 'Gio.AsyncResult' not supported

// UNSUPPORTED : C value 'soup_session_would_redirect' : parameter 'msg' of type 'Message' not supported

var sessionAsyncStruct *gi.Struct
var sessionAsyncStruct_Once sync.Once

func sessionAsyncStruct_Set() error {
	var err error
	sessionAsyncStruct_Once.Do(func() {
		sessionAsyncStruct, err = gi.StructNew("Soup", "SessionAsync")
	})
	return err
}

type SessionAsync struct {
	native uintptr
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'Session'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'Session'

// UNSUPPORTED : C value 'soup_session_async_new' : return type 'SessionAsync' not supported

// UNSUPPORTED : C value 'soup_session_async_new_with_options' : parameter '...' of type 'nil' not supported

var sessionSyncStruct *gi.Struct
var sessionSyncStruct_Once sync.Once

func sessionSyncStruct_Set() error {
	var err error
	sessionSyncStruct_Once.Do(func() {
		sessionSyncStruct, err = gi.StructNew("Soup", "SessionSync")
	})
	return err
}

type SessionSync struct {
	native uintptr
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'Session'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'Session'

// UNSUPPORTED : C value 'soup_session_sync_new' : return type 'SessionSync' not supported

// UNSUPPORTED : C value 'soup_session_sync_new_with_options' : parameter '...' of type 'nil' not supported

var socketStruct *gi.Struct
var socketStruct_Once sync.Once

func socketStruct_Set() error {
	var err error
	socketStruct_Once.Do(func() {
		socketStruct, err = gi.StructNew("Soup", "Socket")
	})
	return err
}

type Socket struct {
	native uintptr
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
		err = socketStruct_Set()
		if err != nil {
			return
		}
		socketDisconnectFunction, err = socketStruct.InvokerNew("disconnect")
	})
	return err
}

// Disconnect is a representation of the C type soup_socket_disconnect.
func (recv *Socket) Disconnect() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

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
		err = socketStruct_Set()
		if err != nil {
			return
		}
		socketGetFdFunction, err = socketStruct.InvokerNew("get_fd")
	})
	return err
}

// GetFd is a representation of the C type soup_socket_get_fd.
func (recv *Socket) GetFd() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := socketGetFdFunction_Set()
	if err == nil {
		ret = socketGetFdFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

// UNSUPPORTED : C value 'soup_socket_get_local_address' : return type 'Address' not supported

// UNSUPPORTED : C value 'soup_socket_get_remote_address' : return type 'Address' not supported

var socketIsConnectedFunction *gi.Function
var socketIsConnectedFunction_Once sync.Once

func socketIsConnectedFunction_Set() error {
	var err error
	socketIsConnectedFunction_Once.Do(func() {
		err = socketStruct_Set()
		if err != nil {
			return
		}
		socketIsConnectedFunction, err = socketStruct.InvokerNew("is_connected")
	})
	return err
}

// IsConnected is a representation of the C type soup_socket_is_connected.
func (recv *Socket) IsConnected() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

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
		err = socketStruct_Set()
		if err != nil {
			return
		}
		socketIsSslFunction, err = socketStruct.InvokerNew("is_ssl")
	})
	return err
}

// IsSsl is a representation of the C type soup_socket_is_ssl.
func (recv *Socket) IsSsl() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

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
		err = socketStruct_Set()
		if err != nil {
			return
		}
		socketListenFunction, err = socketStruct.InvokerNew("listen")
	})
	return err
}

// Listen is a representation of the C type soup_socket_listen.
func (recv *Socket) Listen() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

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

var websocketConnectionStruct *gi.Struct
var websocketConnectionStruct_Once sync.Once

func websocketConnectionStruct_Set() error {
	var err error
	websocketConnectionStruct_Once.Do(func() {
		websocketConnectionStruct, err = gi.StructNew("Soup", "WebsocketConnection")
	})
	return err
}

type WebsocketConnection struct {
	native uintptr
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
		err = websocketConnectionStruct_Set()
		if err != nil {
			return
		}
		websocketConnectionCloseFunction, err = websocketConnectionStruct.InvokerNew("close")
	})
	return err
}

// Close is a representation of the C type soup_websocket_connection_close.
func (recv *WebsocketConnection) Close(code uint16, data string) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
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
		err = websocketConnectionStruct_Set()
		if err != nil {
			return
		}
		websocketConnectionGetCloseCodeFunction, err = websocketConnectionStruct.InvokerNew("get_close_code")
	})
	return err
}

// GetCloseCode is a representation of the C type soup_websocket_connection_get_close_code.
func (recv *WebsocketConnection) GetCloseCode() uint16 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

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
		err = websocketConnectionStruct_Set()
		if err != nil {
			return
		}
		websocketConnectionGetCloseDataFunction, err = websocketConnectionStruct.InvokerNew("get_close_data")
	})
	return err
}

// GetCloseData is a representation of the C type soup_websocket_connection_get_close_data.
func (recv *WebsocketConnection) GetCloseData() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := websocketConnectionGetCloseDataFunction_Set()
	if err == nil {
		ret = websocketConnectionGetCloseDataFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'soup_websocket_connection_get_connection_type' : return type 'WebsocketConnectionType' not supported

// UNSUPPORTED : C value 'soup_websocket_connection_get_extensions' : return type 'GLib.List' not supported

// UNSUPPORTED : C value 'soup_websocket_connection_get_io_stream' : return type 'Gio.IOStream' not supported

var websocketConnectionGetKeepaliveIntervalFunction *gi.Function
var websocketConnectionGetKeepaliveIntervalFunction_Once sync.Once

func websocketConnectionGetKeepaliveIntervalFunction_Set() error {
	var err error
	websocketConnectionGetKeepaliveIntervalFunction_Once.Do(func() {
		err = websocketConnectionStruct_Set()
		if err != nil {
			return
		}
		websocketConnectionGetKeepaliveIntervalFunction, err = websocketConnectionStruct.InvokerNew("get_keepalive_interval")
	})
	return err
}

// GetKeepaliveInterval is a representation of the C type soup_websocket_connection_get_keepalive_interval.
func (recv *WebsocketConnection) GetKeepaliveInterval() uint32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

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
		err = websocketConnectionStruct_Set()
		if err != nil {
			return
		}
		websocketConnectionGetMaxIncomingPayloadSizeFunction, err = websocketConnectionStruct.InvokerNew("get_max_incoming_payload_size")
	})
	return err
}

// GetMaxIncomingPayloadSize is a representation of the C type soup_websocket_connection_get_max_incoming_payload_size.
func (recv *WebsocketConnection) GetMaxIncomingPayloadSize() uint64 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

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
		err = websocketConnectionStruct_Set()
		if err != nil {
			return
		}
		websocketConnectionGetOriginFunction, err = websocketConnectionStruct.InvokerNew("get_origin")
	})
	return err
}

// GetOrigin is a representation of the C type soup_websocket_connection_get_origin.
func (recv *WebsocketConnection) GetOrigin() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

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
		err = websocketConnectionStruct_Set()
		if err != nil {
			return
		}
		websocketConnectionGetProtocolFunction, err = websocketConnectionStruct.InvokerNew("get_protocol")
	})
	return err
}

// GetProtocol is a representation of the C type soup_websocket_connection_get_protocol.
func (recv *WebsocketConnection) GetProtocol() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := websocketConnectionGetProtocolFunction_Set()
	if err == nil {
		ret = websocketConnectionGetProtocolFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'soup_websocket_connection_get_state' : return type 'WebsocketState' not supported

var websocketConnectionGetUriFunction *gi.Function
var websocketConnectionGetUriFunction_Once sync.Once

func websocketConnectionGetUriFunction_Set() error {
	var err error
	websocketConnectionGetUriFunction_Once.Do(func() {
		err = websocketConnectionStruct_Set()
		if err != nil {
			return
		}
		websocketConnectionGetUriFunction, err = websocketConnectionStruct.InvokerNew("get_uri")
	})
	return err
}

// GetUri is a representation of the C type soup_websocket_connection_get_uri.
func (recv *WebsocketConnection) GetUri() *URI {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := websocketConnectionGetUriFunction_Set()
	if err == nil {
		ret = websocketConnectionGetUriFunction.Invoke(inArgs[:], nil)
	}

	retGo := &URI{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'soup_websocket_connection_send_binary' : parameter 'data' of type 'nil' not supported

// UNSUPPORTED : C value 'soup_websocket_connection_send_message' : parameter 'type' of type 'WebsocketDataType' not supported

var websocketConnectionSendTextFunction *gi.Function
var websocketConnectionSendTextFunction_Once sync.Once

func websocketConnectionSendTextFunction_Set() error {
	var err error
	websocketConnectionSendTextFunction_Once.Do(func() {
		err = websocketConnectionStruct_Set()
		if err != nil {
			return
		}
		websocketConnectionSendTextFunction, err = websocketConnectionStruct.InvokerNew("send_text")
	})
	return err
}

// SendText is a representation of the C type soup_websocket_connection_send_text.
func (recv *WebsocketConnection) SendText(text string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
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
		err = websocketConnectionStruct_Set()
		if err != nil {
			return
		}
		websocketConnectionSetKeepaliveIntervalFunction, err = websocketConnectionStruct.InvokerNew("set_keepalive_interval")
	})
	return err
}

// SetKeepaliveInterval is a representation of the C type soup_websocket_connection_set_keepalive_interval.
func (recv *WebsocketConnection) SetKeepaliveInterval(interval uint32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
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
		err = websocketConnectionStruct_Set()
		if err != nil {
			return
		}
		websocketConnectionSetMaxIncomingPayloadSizeFunction, err = websocketConnectionStruct.InvokerNew("set_max_incoming_payload_size")
	})
	return err
}

// SetMaxIncomingPayloadSize is a representation of the C type soup_websocket_connection_set_max_incoming_payload_size.
func (recv *WebsocketConnection) SetMaxIncomingPayloadSize(maxIncomingPayloadSize uint64) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetUint64(maxIncomingPayloadSize)

	err := websocketConnectionSetMaxIncomingPayloadSizeFunction_Set()
	if err == nil {
		websocketConnectionSetMaxIncomingPayloadSizeFunction.Invoke(inArgs[:], nil)
	}

	return
}

var websocketExtensionStruct *gi.Struct
var websocketExtensionStruct_Once sync.Once

func websocketExtensionStruct_Set() error {
	var err error
	websocketExtensionStruct_Once.Do(func() {
		websocketExtensionStruct, err = gi.StructNew("Soup", "WebsocketExtension")
	})
	return err
}

type WebsocketExtension struct {
	native uintptr
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'soup_websocket_extension_configure' : parameter 'connection_type' of type 'WebsocketConnectionType' not supported

var websocketExtensionGetRequestParamsFunction *gi.Function
var websocketExtensionGetRequestParamsFunction_Once sync.Once

func websocketExtensionGetRequestParamsFunction_Set() error {
	var err error
	websocketExtensionGetRequestParamsFunction_Once.Do(func() {
		err = websocketExtensionStruct_Set()
		if err != nil {
			return
		}
		websocketExtensionGetRequestParamsFunction, err = websocketExtensionStruct.InvokerNew("get_request_params")
	})
	return err
}

// GetRequestParams is a representation of the C type soup_websocket_extension_get_request_params.
func (recv *WebsocketExtension) GetRequestParams() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

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
		err = websocketExtensionStruct_Set()
		if err != nil {
			return
		}
		websocketExtensionGetResponseParamsFunction, err = websocketExtensionStruct.InvokerNew("get_response_params")
	})
	return err
}

// GetResponseParams is a representation of the C type soup_websocket_extension_get_response_params.
func (recv *WebsocketExtension) GetResponseParams() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

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

// WebsocketExtensionStruct creates an uninitialised WebsocketExtension.
func WebsocketExtensionStruct() *WebsocketExtension {
	err := websocketExtensionStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &WebsocketExtension{native: websocketExtensionStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeWebsocketExtension)
	return structGo
}
func finalizeWebsocketExtension(obj *WebsocketExtension) {
	websocketExtensionStruct.Free(obj.native)
}

var websocketExtensionDeflateStruct *gi.Struct
var websocketExtensionDeflateStruct_Once sync.Once

func websocketExtensionDeflateStruct_Set() error {
	var err error
	websocketExtensionDeflateStruct_Once.Do(func() {
		websocketExtensionDeflateStruct, err = gi.StructNew("Soup", "WebsocketExtensionDeflate")
	})
	return err
}

type WebsocketExtensionDeflate struct {
	native uintptr
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'WebsocketExtension'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'WebsocketExtension'

// WebsocketExtensionDeflateStruct creates an uninitialised WebsocketExtensionDeflate.
func WebsocketExtensionDeflateStruct() *WebsocketExtensionDeflate {
	err := websocketExtensionDeflateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &WebsocketExtensionDeflate{native: websocketExtensionDeflateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeWebsocketExtensionDeflate)
	return structGo
}
func finalizeWebsocketExtensionDeflate(obj *WebsocketExtensionDeflate) {
	websocketExtensionDeflateStruct.Free(obj.native)
}

var websocketExtensionManagerStruct *gi.Struct
var websocketExtensionManagerStruct_Once sync.Once

func websocketExtensionManagerStruct_Set() error {
	var err error
	websocketExtensionManagerStruct_Once.Do(func() {
		websocketExtensionManagerStruct, err = gi.StructNew("Soup", "WebsocketExtensionManager")
	})
	return err
}

type WebsocketExtensionManager struct {
	native uintptr
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'GObject.Object'

// WebsocketExtensionManagerStruct creates an uninitialised WebsocketExtensionManager.
func WebsocketExtensionManagerStruct() *WebsocketExtensionManager {
	err := websocketExtensionManagerStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &WebsocketExtensionManager{native: websocketExtensionManagerStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeWebsocketExtensionManager)
	return structGo
}
func finalizeWebsocketExtensionManager(obj *WebsocketExtensionManager) {
	websocketExtensionManagerStruct.Free(obj.native)
}
