// Code generated - DO NOT EDIT.

package soup

import (
	gi "github.com/pekim/gobbi/internal/cgo/gi"
	gio "github.com/pekim/gobbi/lib/gio"
	gobject "github.com/pekim/gobbi/lib/gobject"
	"sync"
	"unsafe"
)

var passwordManagerInterface *gi.Interface
var passwordManagerInterface_Once sync.Once

func passwordManagerInterface_Set() error {
	var err error
	passwordManagerInterface_Once.Do(func() {
		passwordManagerInterface, err = gi.InterfaceNew("Soup", "PasswordManager")
	})
	return err
}

type PasswordManager struct {
	native unsafe.Pointer
}

func PasswordManagerNewFromNative(native unsafe.Pointer) *PasswordManager {
	instance := &PasswordManager{native: native}

	return instance
}

/*
CastToPasswordManager down casts any arbitrary Object to PasswordManager.
Exercise care, as this is a potentially dangerous function
if the Object is not a PasswordManager.
*/
func CastToPasswordManager(object *gobject.Object) *PasswordManager {
	return PasswordManagerNewFromNative(object.Native())
}

// Equals compares this PasswordManager with another PasswordManager, and returns true if they represent the same Object.
func (recv *PasswordManager) Equals(other *PasswordManager) bool {
	return other.Native() == recv.Native()
}

func (recv *PasswordManager) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'soup_password_manager_get_passwords_async' : parameter 'callback' of type 'PasswordManagerCallback' not supported

var passwordManagerGetPasswordsSyncFunction *gi.Function
var passwordManagerGetPasswordsSyncFunction_Once sync.Once

func passwordManagerGetPasswordsSyncFunction_Set() error {
	var err error
	passwordManagerGetPasswordsSyncFunction_Once.Do(func() {
		err = passwordManagerInterface_Set()
		if err != nil {
			return
		}
		passwordManagerGetPasswordsSyncFunction, err = passwordManagerInterface.InvokerNew("get_passwords_sync")
	})
	return err
}

// GetPasswordsSync is a representation of the C type soup_password_manager_get_passwords_sync.
func (recv *PasswordManager) GetPasswordsSync(msg *Message, auth *Auth, cancellable *gio.Cancellable) {
	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(msg.Native())
	inArgs[2].SetPointer(auth.Native())
	inArgs[3].SetPointer(cancellable.Native())

	err := passwordManagerGetPasswordsSyncFunction_Set()
	if err == nil {
		passwordManagerGetPasswordsSyncFunction.Invoke(inArgs[:], nil)
	}

	return
}

var proxyResolverInterface *gi.Interface
var proxyResolverInterface_Once sync.Once

func proxyResolverInterface_Set() error {
	var err error
	proxyResolverInterface_Once.Do(func() {
		proxyResolverInterface, err = gi.InterfaceNew("Soup", "ProxyResolver")
	})
	return err
}

type ProxyResolver struct {
	native unsafe.Pointer
}

func ProxyResolverNewFromNative(native unsafe.Pointer) *ProxyResolver {
	instance := &ProxyResolver{native: native}

	return instance
}

/*
CastToProxyResolver down casts any arbitrary Object to ProxyResolver.
Exercise care, as this is a potentially dangerous function
if the Object is not a ProxyResolver.
*/
func CastToProxyResolver(object *gobject.Object) *ProxyResolver {
	return ProxyResolverNewFromNative(object.Native())
}

// Equals compares this ProxyResolver with another ProxyResolver, and returns true if they represent the same Object.
func (recv *ProxyResolver) Equals(other *ProxyResolver) bool {
	return other.Native() == recv.Native()
}

func (recv *ProxyResolver) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'soup_proxy_resolver_get_proxy_async' : parameter 'callback' of type 'ProxyResolverCallback' not supported

var proxyResolverGetProxySyncFunction *gi.Function
var proxyResolverGetProxySyncFunction_Once sync.Once

func proxyResolverGetProxySyncFunction_Set() error {
	var err error
	proxyResolverGetProxySyncFunction_Once.Do(func() {
		err = proxyResolverInterface_Set()
		if err != nil {
			return
		}
		proxyResolverGetProxySyncFunction, err = proxyResolverInterface.InvokerNew("get_proxy_sync")
	})
	return err
}

// GetProxySync is a representation of the C type soup_proxy_resolver_get_proxy_sync.
func (recv *ProxyResolver) GetProxySync(msg *Message, cancellable *gio.Cancellable) (uint32, *Address) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(msg.Native())
	inArgs[2].SetPointer(cancellable.Native())

	var outArgs [1]gi.Argument
	var ret gi.Argument

	err := proxyResolverGetProxySyncFunction_Set()
	if err == nil {
		ret = proxyResolverGetProxySyncFunction.Invoke(inArgs[:], outArgs[:])
	}

	retGo := ret.Uint32()
	out0 := AddressNewFromNative(outArgs[0].Pointer())

	return retGo, out0
}

var proxyURIResolverInterface *gi.Interface
var proxyURIResolverInterface_Once sync.Once

func proxyURIResolverInterface_Set() error {
	var err error
	proxyURIResolverInterface_Once.Do(func() {
		proxyURIResolverInterface, err = gi.InterfaceNew("Soup", "ProxyURIResolver")
	})
	return err
}

type ProxyURIResolver struct {
	native unsafe.Pointer
}

func ProxyURIResolverNewFromNative(native unsafe.Pointer) *ProxyURIResolver {
	instance := &ProxyURIResolver{native: native}

	return instance
}

/*
CastToProxyURIResolver down casts any arbitrary Object to ProxyURIResolver.
Exercise care, as this is a potentially dangerous function
if the Object is not a ProxyURIResolver.
*/
func CastToProxyURIResolver(object *gobject.Object) *ProxyURIResolver {
	return ProxyURIResolverNewFromNative(object.Native())
}

// Equals compares this ProxyURIResolver with another ProxyURIResolver, and returns true if they represent the same Object.
func (recv *ProxyURIResolver) Equals(other *ProxyURIResolver) bool {
	return other.Native() == recv.Native()
}

func (recv *ProxyURIResolver) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'soup_proxy_uri_resolver_get_proxy_uri_async' : parameter 'callback' of type 'ProxyURIResolverCallback' not supported

var proxyURIResolverGetProxyUriSyncFunction *gi.Function
var proxyURIResolverGetProxyUriSyncFunction_Once sync.Once

func proxyURIResolverGetProxyUriSyncFunction_Set() error {
	var err error
	proxyURIResolverGetProxyUriSyncFunction_Once.Do(func() {
		err = proxyURIResolverInterface_Set()
		if err != nil {
			return
		}
		proxyURIResolverGetProxyUriSyncFunction, err = proxyURIResolverInterface.InvokerNew("get_proxy_uri_sync")
	})
	return err
}

// GetProxyUriSync is a representation of the C type soup_proxy_uri_resolver_get_proxy_uri_sync.
func (recv *ProxyURIResolver) GetProxyUriSync(uri *URI, cancellable *gio.Cancellable) (uint32, *URI) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(uri.Native())
	inArgs[2].SetPointer(cancellable.Native())

	var outArgs [1]gi.Argument
	var ret gi.Argument

	err := proxyURIResolverGetProxyUriSyncFunction_Set()
	if err == nil {
		ret = proxyURIResolverGetProxyUriSyncFunction.Invoke(inArgs[:], outArgs[:])
	}

	retGo := ret.Uint32()
	out0 := URINewFromNative(outArgs[0].Pointer())

	return retGo, out0
}

var sessionFeatureInterface *gi.Interface
var sessionFeatureInterface_Once sync.Once

func sessionFeatureInterface_Set() error {
	var err error
	sessionFeatureInterface_Once.Do(func() {
		sessionFeatureInterface, err = gi.InterfaceNew("Soup", "SessionFeature")
	})
	return err
}

type SessionFeature struct {
	native unsafe.Pointer
}

func SessionFeatureNewFromNative(native unsafe.Pointer) *SessionFeature {
	instance := &SessionFeature{native: native}

	return instance
}

/*
CastToSessionFeature down casts any arbitrary Object to SessionFeature.
Exercise care, as this is a potentially dangerous function
if the Object is not a SessionFeature.
*/
func CastToSessionFeature(object *gobject.Object) *SessionFeature {
	return SessionFeatureNewFromNative(object.Native())
}

// Equals compares this SessionFeature with another SessionFeature, and returns true if they represent the same Object.
func (recv *SessionFeature) Equals(other *SessionFeature) bool {
	return other.Native() == recv.Native()
}

func (recv *SessionFeature) Native() unsafe.Pointer {
	return recv.native
}

var sessionFeatureAddFeatureFunction *gi.Function
var sessionFeatureAddFeatureFunction_Once sync.Once

func sessionFeatureAddFeatureFunction_Set() error {
	var err error
	sessionFeatureAddFeatureFunction_Once.Do(func() {
		err = sessionFeatureInterface_Set()
		if err != nil {
			return
		}
		sessionFeatureAddFeatureFunction, err = sessionFeatureInterface.InvokerNew("add_feature")
	})
	return err
}

// AddFeature is a representation of the C type soup_session_feature_add_feature.
func (recv *SessionFeature) AddFeature(type_ int64) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt64(type_)

	var ret gi.Argument

	err := sessionFeatureAddFeatureFunction_Set()
	if err == nil {
		ret = sessionFeatureAddFeatureFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var sessionFeatureAttachFunction *gi.Function
var sessionFeatureAttachFunction_Once sync.Once

func sessionFeatureAttachFunction_Set() error {
	var err error
	sessionFeatureAttachFunction_Once.Do(func() {
		err = sessionFeatureInterface_Set()
		if err != nil {
			return
		}
		sessionFeatureAttachFunction, err = sessionFeatureInterface.InvokerNew("attach")
	})
	return err
}

// Attach is a representation of the C type soup_session_feature_attach.
func (recv *SessionFeature) Attach(session *Session) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(session.Native())

	err := sessionFeatureAttachFunction_Set()
	if err == nil {
		sessionFeatureAttachFunction.Invoke(inArgs[:], nil)
	}

	return
}

var sessionFeatureDetachFunction *gi.Function
var sessionFeatureDetachFunction_Once sync.Once

func sessionFeatureDetachFunction_Set() error {
	var err error
	sessionFeatureDetachFunction_Once.Do(func() {
		err = sessionFeatureInterface_Set()
		if err != nil {
			return
		}
		sessionFeatureDetachFunction, err = sessionFeatureInterface.InvokerNew("detach")
	})
	return err
}

// Detach is a representation of the C type soup_session_feature_detach.
func (recv *SessionFeature) Detach(session *Session) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(session.Native())

	err := sessionFeatureDetachFunction_Set()
	if err == nil {
		sessionFeatureDetachFunction.Invoke(inArgs[:], nil)
	}

	return
}

var sessionFeatureHasFeatureFunction *gi.Function
var sessionFeatureHasFeatureFunction_Once sync.Once

func sessionFeatureHasFeatureFunction_Set() error {
	var err error
	sessionFeatureHasFeatureFunction_Once.Do(func() {
		err = sessionFeatureInterface_Set()
		if err != nil {
			return
		}
		sessionFeatureHasFeatureFunction, err = sessionFeatureInterface.InvokerNew("has_feature")
	})
	return err
}

// HasFeature is a representation of the C type soup_session_feature_has_feature.
func (recv *SessionFeature) HasFeature(type_ int64) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt64(type_)

	var ret gi.Argument

	err := sessionFeatureHasFeatureFunction_Set()
	if err == nil {
		ret = sessionFeatureHasFeatureFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var sessionFeatureRemoveFeatureFunction *gi.Function
var sessionFeatureRemoveFeatureFunction_Once sync.Once

func sessionFeatureRemoveFeatureFunction_Set() error {
	var err error
	sessionFeatureRemoveFeatureFunction_Once.Do(func() {
		err = sessionFeatureInterface_Set()
		if err != nil {
			return
		}
		sessionFeatureRemoveFeatureFunction, err = sessionFeatureInterface.InvokerNew("remove_feature")
	})
	return err
}

// RemoveFeature is a representation of the C type soup_session_feature_remove_feature.
func (recv *SessionFeature) RemoveFeature(type_ int64) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt64(type_)

	var ret gi.Argument

	err := sessionFeatureRemoveFeatureFunction_Set()
	if err == nil {
		ret = sessionFeatureRemoveFeatureFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}
