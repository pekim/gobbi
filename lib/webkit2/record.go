// Code generated - DO NOT EDIT.

package webkit2

import (
	gi "github.com/pekim/gobbi/internal/gi"
	"sync"
)

var applicationInfoStruct *gi.Struct
var applicationInfoStructOnce sync.Once

func applicationInfoStructSet() {
	applicationInfoStructOnce.Do(func() {
		applicationInfoStruct = gi.StructNew("WebKit2", "ApplicationInfo")
	})
}

type ApplicationInfo struct {
	native uintptr
}

var applicationInfoNewFunction *gi.Function
var applicationInfoNewFunctionOnce sync.Once

func applicationInfoNewFunctionSet() {
	applicationInfoNewFunctionOnce.Do(func() {
		applicationInfoNewFunction = gi.FunctionInvokerNew("WebKit2", "new")
	})
}

var newApplicationInfoInvoker *gi.Function

// ApplicationInfoNew is a representation of the C type webkit_application_info_new.
func ApplicationInfoNew() *ApplicationInfo {
	if newApplicationInfoInvoker == nil {
		newApplicationInfoInvoker = gi.StructFunctionInvokerNew("WebKit2", "ApplicationInfo", "new")
	}

	ret := newApplicationInfoInvoker.Invoke(nil, nil)

	retGo := &ApplicationInfo{native: ret.Pointer()}

	return retGo
}

var applicationInfoGetNameFunction *gi.Function
var applicationInfoGetNameFunctionOnce sync.Once

func applicationInfoGetNameFunctionSet() {
	applicationInfoGetNameFunctionOnce.Do(func() {
		applicationInfoGetNameFunction = gi.FunctionInvokerNew("WebKit2", "get_name")
	})
}

var getNameApplicationInfoInvoker *gi.Function

// GetName is a representation of the C type webkit_application_info_get_name.
func (recv *ApplicationInfo) GetName() string {
	if getNameApplicationInfoInvoker == nil {
		getNameApplicationInfoInvoker = gi.StructFunctionInvokerNew("WebKit2", "ApplicationInfo", "get_name")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getNameApplicationInfoInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

var applicationInfoGetVersionFunction *gi.Function
var applicationInfoGetVersionFunctionOnce sync.Once

func applicationInfoGetVersionFunctionSet() {
	applicationInfoGetVersionFunctionOnce.Do(func() {
		applicationInfoGetVersionFunction = gi.FunctionInvokerNew("WebKit2", "get_version")
	})
}

var getVersionApplicationInfoInvoker *gi.Function

// GetVersion is a representation of the C type webkit_application_info_get_version.
func (recv *ApplicationInfo) GetVersion() (uint64, uint64, uint64) {
	if getVersionApplicationInfoInvoker == nil {
		getVersionApplicationInfoInvoker = gi.StructFunctionInvokerNew("WebKit2", "ApplicationInfo", "get_version")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var outArgs [3]gi.Argument

	getVersionApplicationInfoInvoker.Invoke(inArgs[:], outArgs[:])

	out0 := outArgs[0].Uint64()
	out1 := outArgs[1].Uint64()
	out2 := outArgs[2].Uint64()

	return out0, out1, out2
}

var applicationInfoRefFunction *gi.Function
var applicationInfoRefFunctionOnce sync.Once

func applicationInfoRefFunctionSet() {
	applicationInfoRefFunctionOnce.Do(func() {
		applicationInfoRefFunction = gi.FunctionInvokerNew("WebKit2", "ref")
	})
}

var refApplicationInfoInvoker *gi.Function

// Ref is a representation of the C type webkit_application_info_ref.
func (recv *ApplicationInfo) Ref() *ApplicationInfo {
	if refApplicationInfoInvoker == nil {
		refApplicationInfoInvoker = gi.StructFunctionInvokerNew("WebKit2", "ApplicationInfo", "ref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := refApplicationInfoInvoker.Invoke(inArgs[:], nil)

	retGo := &ApplicationInfo{native: ret.Pointer()}

	return retGo
}

var applicationInfoSetNameFunction *gi.Function
var applicationInfoSetNameFunctionOnce sync.Once

func applicationInfoSetNameFunctionSet() {
	applicationInfoSetNameFunctionOnce.Do(func() {
		applicationInfoSetNameFunction = gi.FunctionInvokerNew("WebKit2", "set_name")
	})
}

var setNameApplicationInfoInvoker *gi.Function

// SetName is a representation of the C type webkit_application_info_set_name.
func (recv *ApplicationInfo) SetName(name string) {
	if setNameApplicationInfoInvoker == nil {
		setNameApplicationInfoInvoker = gi.StructFunctionInvokerNew("WebKit2", "ApplicationInfo", "set_name")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(name)

	setNameApplicationInfoInvoker.Invoke(inArgs[:], nil)

}

var applicationInfoSetVersionFunction *gi.Function
var applicationInfoSetVersionFunctionOnce sync.Once

func applicationInfoSetVersionFunctionSet() {
	applicationInfoSetVersionFunctionOnce.Do(func() {
		applicationInfoSetVersionFunction = gi.FunctionInvokerNew("WebKit2", "set_version")
	})
}

var setVersionApplicationInfoInvoker *gi.Function

// SetVersion is a representation of the C type webkit_application_info_set_version.
func (recv *ApplicationInfo) SetVersion(major uint64, minor uint64, micro uint64) {
	if setVersionApplicationInfoInvoker == nil {
		setVersionApplicationInfoInvoker = gi.StructFunctionInvokerNew("WebKit2", "ApplicationInfo", "set_version")
	}

	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetUint64(major)
	inArgs[2].SetUint64(minor)
	inArgs[3].SetUint64(micro)

	setVersionApplicationInfoInvoker.Invoke(inArgs[:], nil)

}

var applicationInfoUnrefFunction *gi.Function
var applicationInfoUnrefFunctionOnce sync.Once

func applicationInfoUnrefFunctionSet() {
	applicationInfoUnrefFunctionOnce.Do(func() {
		applicationInfoUnrefFunction = gi.FunctionInvokerNew("WebKit2", "unref")
	})
}

var unrefApplicationInfoInvoker *gi.Function

// Unref is a representation of the C type webkit_application_info_unref.
func (recv *ApplicationInfo) Unref() {
	if unrefApplicationInfoInvoker == nil {
		unrefApplicationInfoInvoker = gi.StructFunctionInvokerNew("WebKit2", "ApplicationInfo", "unref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	unrefApplicationInfoInvoker.Invoke(inArgs[:], nil)

}

var authenticationRequestClassStruct *gi.Struct
var authenticationRequestClassStructOnce sync.Once

func authenticationRequestClassStructSet() {
	authenticationRequestClassStructOnce.Do(func() {
		authenticationRequestClassStruct = gi.StructNew("WebKit2", "AuthenticationRequestClass")
	})
}

type AuthenticationRequestClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value '_webkit_reserved0' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved1' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved2' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved3' : missing Type
}

var authenticationRequestPrivateStruct *gi.Struct
var authenticationRequestPrivateStructOnce sync.Once

func authenticationRequestPrivateStructSet() {
	authenticationRequestPrivateStructOnce.Do(func() {
		authenticationRequestPrivateStruct = gi.StructNew("WebKit2", "AuthenticationRequestPrivate")
	})
}

type AuthenticationRequestPrivate struct {
	native uintptr
}

var automationSessionClassStruct *gi.Struct
var automationSessionClassStructOnce sync.Once

func automationSessionClassStructSet() {
	automationSessionClassStructOnce.Do(func() {
		automationSessionClassStruct = gi.StructNew("WebKit2", "AutomationSessionClass")
	})
}

type AutomationSessionClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value '_webkit_reserved0' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved1' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved2' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved3' : missing Type
}

var automationSessionPrivateStruct *gi.Struct
var automationSessionPrivateStructOnce sync.Once

func automationSessionPrivateStructSet() {
	automationSessionPrivateStructOnce.Do(func() {
		automationSessionPrivateStruct = gi.StructNew("WebKit2", "AutomationSessionPrivate")
	})
}

type AutomationSessionPrivate struct {
	native uintptr
}

var backForwardListClassStruct *gi.Struct
var backForwardListClassStructOnce sync.Once

func backForwardListClassStructSet() {
	backForwardListClassStructOnce.Do(func() {
		backForwardListClassStruct = gi.StructNew("WebKit2", "BackForwardListClass")
	})
}

type BackForwardListClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value '_webkit_reserved0' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved1' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved2' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved3' : missing Type
}

var backForwardListItemClassStruct *gi.Struct
var backForwardListItemClassStructOnce sync.Once

func backForwardListItemClassStructSet() {
	backForwardListItemClassStructOnce.Do(func() {
		backForwardListItemClassStruct = gi.StructNew("WebKit2", "BackForwardListItemClass")
	})
}

type BackForwardListItemClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.InitiallyUnownedClass'
	// UNSUPPORTED : C value '_webkit_reserved0' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved1' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved2' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved3' : missing Type
}

var backForwardListItemPrivateStruct *gi.Struct
var backForwardListItemPrivateStructOnce sync.Once

func backForwardListItemPrivateStructSet() {
	backForwardListItemPrivateStructOnce.Do(func() {
		backForwardListItemPrivateStruct = gi.StructNew("WebKit2", "BackForwardListItemPrivate")
	})
}

type BackForwardListItemPrivate struct {
	native uintptr
}

var backForwardListPrivateStruct *gi.Struct
var backForwardListPrivateStructOnce sync.Once

func backForwardListPrivateStructSet() {
	backForwardListPrivateStructOnce.Do(func() {
		backForwardListPrivateStruct = gi.StructNew("WebKit2", "BackForwardListPrivate")
	})
}

type BackForwardListPrivate struct {
	native uintptr
}

var colorChooserRequestClassStruct *gi.Struct
var colorChooserRequestClassStructOnce sync.Once

func colorChooserRequestClassStructSet() {
	colorChooserRequestClassStructOnce.Do(func() {
		colorChooserRequestClassStruct = gi.StructNew("WebKit2", "ColorChooserRequestClass")
	})
}

type ColorChooserRequestClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
}

var colorChooserRequestPrivateStruct *gi.Struct
var colorChooserRequestPrivateStructOnce sync.Once

func colorChooserRequestPrivateStructSet() {
	colorChooserRequestPrivateStructOnce.Do(func() {
		colorChooserRequestPrivateStruct = gi.StructNew("WebKit2", "ColorChooserRequestPrivate")
	})
}

type ColorChooserRequestPrivate struct {
	native uintptr
}

var contextMenuClassStruct *gi.Struct
var contextMenuClassStructOnce sync.Once

func contextMenuClassStructSet() {
	contextMenuClassStructOnce.Do(func() {
		contextMenuClassStruct = gi.StructNew("WebKit2", "ContextMenuClass")
	})
}

type ContextMenuClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value '_webkit_reserved0' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved1' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved2' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved3' : missing Type
}

var contextMenuItemClassStruct *gi.Struct
var contextMenuItemClassStructOnce sync.Once

func contextMenuItemClassStructSet() {
	contextMenuItemClassStructOnce.Do(func() {
		contextMenuItemClassStruct = gi.StructNew("WebKit2", "ContextMenuItemClass")
	})
}

type ContextMenuItemClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.InitiallyUnownedClass'
	// UNSUPPORTED : C value '_webkit_reserved0' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved1' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved2' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved3' : missing Type
}

var contextMenuItemPrivateStruct *gi.Struct
var contextMenuItemPrivateStructOnce sync.Once

func contextMenuItemPrivateStructSet() {
	contextMenuItemPrivateStructOnce.Do(func() {
		contextMenuItemPrivateStruct = gi.StructNew("WebKit2", "ContextMenuItemPrivate")
	})
}

type ContextMenuItemPrivate struct {
	native uintptr
}

var contextMenuPrivateStruct *gi.Struct
var contextMenuPrivateStructOnce sync.Once

func contextMenuPrivateStructSet() {
	contextMenuPrivateStructOnce.Do(func() {
		contextMenuPrivateStruct = gi.StructNew("WebKit2", "ContextMenuPrivate")
	})
}

type ContextMenuPrivate struct {
	native uintptr
}

var cookieManagerClassStruct *gi.Struct
var cookieManagerClassStructOnce sync.Once

func cookieManagerClassStructSet() {
	cookieManagerClassStructOnce.Do(func() {
		cookieManagerClassStruct = gi.StructNew("WebKit2", "CookieManagerClass")
	})
}

type CookieManagerClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value '_webkit_reserved0' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved1' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved2' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved3' : missing Type
}

var cookieManagerPrivateStruct *gi.Struct
var cookieManagerPrivateStructOnce sync.Once

func cookieManagerPrivateStructSet() {
	cookieManagerPrivateStructOnce.Do(func() {
		cookieManagerPrivateStruct = gi.StructNew("WebKit2", "CookieManagerPrivate")
	})
}

type CookieManagerPrivate struct {
	native uintptr
}

var credentialStruct *gi.Struct
var credentialStructOnce sync.Once

func credentialStructSet() {
	credentialStructOnce.Do(func() {
		credentialStruct = gi.StructNew("WebKit2", "Credential")
	})
}

type Credential struct {
	native uintptr
}

// UNSUPPORTED : C value 'webkit_credential_new' : parameter 'persistence' of type 'CredentialPersistence' not supported

var credentialCopyFunction *gi.Function
var credentialCopyFunctionOnce sync.Once

func credentialCopyFunctionSet() {
	credentialCopyFunctionOnce.Do(func() {
		credentialCopyFunction = gi.FunctionInvokerNew("WebKit2", "copy")
	})
}

var copyCredentialInvoker *gi.Function

// Copy is a representation of the C type webkit_credential_copy.
func (recv *Credential) Copy() *Credential {
	if copyCredentialInvoker == nil {
		copyCredentialInvoker = gi.StructFunctionInvokerNew("WebKit2", "Credential", "copy")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := copyCredentialInvoker.Invoke(inArgs[:], nil)

	retGo := &Credential{native: ret.Pointer()}

	return retGo
}

var credentialFreeFunction *gi.Function
var credentialFreeFunctionOnce sync.Once

func credentialFreeFunctionSet() {
	credentialFreeFunctionOnce.Do(func() {
		credentialFreeFunction = gi.FunctionInvokerNew("WebKit2", "free")
	})
}

var freeCredentialInvoker *gi.Function

// Free is a representation of the C type webkit_credential_free.
func (recv *Credential) Free() {
	if freeCredentialInvoker == nil {
		freeCredentialInvoker = gi.StructFunctionInvokerNew("WebKit2", "Credential", "free")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	freeCredentialInvoker.Invoke(inArgs[:], nil)

}

var credentialGetPasswordFunction *gi.Function
var credentialGetPasswordFunctionOnce sync.Once

func credentialGetPasswordFunctionSet() {
	credentialGetPasswordFunctionOnce.Do(func() {
		credentialGetPasswordFunction = gi.FunctionInvokerNew("WebKit2", "get_password")
	})
}

var getPasswordCredentialInvoker *gi.Function

// GetPassword is a representation of the C type webkit_credential_get_password.
func (recv *Credential) GetPassword() string {
	if getPasswordCredentialInvoker == nil {
		getPasswordCredentialInvoker = gi.StructFunctionInvokerNew("WebKit2", "Credential", "get_password")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getPasswordCredentialInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'webkit_credential_get_persistence' : return type 'CredentialPersistence' not supported

var credentialGetUsernameFunction *gi.Function
var credentialGetUsernameFunctionOnce sync.Once

func credentialGetUsernameFunctionSet() {
	credentialGetUsernameFunctionOnce.Do(func() {
		credentialGetUsernameFunction = gi.FunctionInvokerNew("WebKit2", "get_username")
	})
}

var getUsernameCredentialInvoker *gi.Function

// GetUsername is a representation of the C type webkit_credential_get_username.
func (recv *Credential) GetUsername() string {
	if getUsernameCredentialInvoker == nil {
		getUsernameCredentialInvoker = gi.StructFunctionInvokerNew("WebKit2", "Credential", "get_username")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getUsernameCredentialInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'webkit_credential_has_password' : return type 'gboolean' not supported

var deviceInfoPermissionRequestClassStruct *gi.Struct
var deviceInfoPermissionRequestClassStructOnce sync.Once

func deviceInfoPermissionRequestClassStructSet() {
	deviceInfoPermissionRequestClassStructOnce.Do(func() {
		deviceInfoPermissionRequestClassStruct = gi.StructNew("WebKit2", "DeviceInfoPermissionRequestClass")
	})
}

type DeviceInfoPermissionRequestClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value '_webkit_reserved0' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved1' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved2' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved3' : missing Type
}

var deviceInfoPermissionRequestPrivateStruct *gi.Struct
var deviceInfoPermissionRequestPrivateStructOnce sync.Once

func deviceInfoPermissionRequestPrivateStructSet() {
	deviceInfoPermissionRequestPrivateStructOnce.Do(func() {
		deviceInfoPermissionRequestPrivateStruct = gi.StructNew("WebKit2", "DeviceInfoPermissionRequestPrivate")
	})
}

type DeviceInfoPermissionRequestPrivate struct {
	native uintptr
}

var downloadClassStruct *gi.Struct
var downloadClassStructOnce sync.Once

func downloadClassStructSet() {
	downloadClassStructOnce.Do(func() {
		downloadClassStruct = gi.StructNew("WebKit2", "DownloadClass")
	})
}

type DownloadClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value 'decide_destination' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved0' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved1' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved2' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved3' : missing Type
}

var downloadPrivateStruct *gi.Struct
var downloadPrivateStructOnce sync.Once

func downloadPrivateStructSet() {
	downloadPrivateStructOnce.Do(func() {
		downloadPrivateStruct = gi.StructNew("WebKit2", "DownloadPrivate")
	})
}

type DownloadPrivate struct {
	native uintptr
}

var editorStateClassStruct *gi.Struct
var editorStateClassStructOnce sync.Once

func editorStateClassStructSet() {
	editorStateClassStructOnce.Do(func() {
		editorStateClassStruct = gi.StructNew("WebKit2", "EditorStateClass")
	})
}

type EditorStateClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value '_webkit_reserved0' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved1' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved2' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved3' : missing Type
}

var editorStatePrivateStruct *gi.Struct
var editorStatePrivateStructOnce sync.Once

func editorStatePrivateStructSet() {
	editorStatePrivateStructOnce.Do(func() {
		editorStatePrivateStruct = gi.StructNew("WebKit2", "EditorStatePrivate")
	})
}

type EditorStatePrivate struct {
	native uintptr
}

var faviconDatabaseClassStruct *gi.Struct
var faviconDatabaseClassStructOnce sync.Once

func faviconDatabaseClassStructSet() {
	faviconDatabaseClassStructOnce.Do(func() {
		faviconDatabaseClassStruct = gi.StructNew("WebKit2", "FaviconDatabaseClass")
	})
}

type FaviconDatabaseClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value '_webkit_reserved0' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved1' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved2' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved3' : missing Type
}

var faviconDatabasePrivateStruct *gi.Struct
var faviconDatabasePrivateStructOnce sync.Once

func faviconDatabasePrivateStructSet() {
	faviconDatabasePrivateStructOnce.Do(func() {
		faviconDatabasePrivateStruct = gi.StructNew("WebKit2", "FaviconDatabasePrivate")
	})
}

type FaviconDatabasePrivate struct {
	native uintptr
}

var fileChooserRequestClassStruct *gi.Struct
var fileChooserRequestClassStructOnce sync.Once

func fileChooserRequestClassStructSet() {
	fileChooserRequestClassStructOnce.Do(func() {
		fileChooserRequestClassStruct = gi.StructNew("WebKit2", "FileChooserRequestClass")
	})
}

type FileChooserRequestClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value '_webkit_reserved0' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved1' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved2' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved3' : missing Type
}

var fileChooserRequestPrivateStruct *gi.Struct
var fileChooserRequestPrivateStructOnce sync.Once

func fileChooserRequestPrivateStructSet() {
	fileChooserRequestPrivateStructOnce.Do(func() {
		fileChooserRequestPrivateStruct = gi.StructNew("WebKit2", "FileChooserRequestPrivate")
	})
}

type FileChooserRequestPrivate struct {
	native uintptr
}

var findControllerClassStruct *gi.Struct
var findControllerClassStructOnce sync.Once

func findControllerClassStructSet() {
	findControllerClassStructOnce.Do(func() {
		findControllerClassStruct = gi.StructNew("WebKit2", "FindControllerClass")
	})
}

type FindControllerClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value '_webkit_reserved0' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved1' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved2' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved3' : missing Type
}

var findControllerPrivateStruct *gi.Struct
var findControllerPrivateStructOnce sync.Once

func findControllerPrivateStructSet() {
	findControllerPrivateStructOnce.Do(func() {
		findControllerPrivateStruct = gi.StructNew("WebKit2", "FindControllerPrivate")
	})
}

type FindControllerPrivate struct {
	native uintptr
}

var formSubmissionRequestClassStruct *gi.Struct
var formSubmissionRequestClassStructOnce sync.Once

func formSubmissionRequestClassStructSet() {
	formSubmissionRequestClassStructOnce.Do(func() {
		formSubmissionRequestClassStruct = gi.StructNew("WebKit2", "FormSubmissionRequestClass")
	})
}

type FormSubmissionRequestClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value '_webkit_reserved0' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved1' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved2' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved3' : missing Type
}

var formSubmissionRequestPrivateStruct *gi.Struct
var formSubmissionRequestPrivateStructOnce sync.Once

func formSubmissionRequestPrivateStructSet() {
	formSubmissionRequestPrivateStructOnce.Do(func() {
		formSubmissionRequestPrivateStruct = gi.StructNew("WebKit2", "FormSubmissionRequestPrivate")
	})
}

type FormSubmissionRequestPrivate struct {
	native uintptr
}

var geolocationManagerClassStruct *gi.Struct
var geolocationManagerClassStructOnce sync.Once

func geolocationManagerClassStructSet() {
	geolocationManagerClassStructOnce.Do(func() {
		geolocationManagerClassStruct = gi.StructNew("WebKit2", "GeolocationManagerClass")
	})
}

type GeolocationManagerClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value '_webkit_reserved0' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved1' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved2' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved3' : missing Type
}

var geolocationManagerPrivateStruct *gi.Struct
var geolocationManagerPrivateStructOnce sync.Once

func geolocationManagerPrivateStructSet() {
	geolocationManagerPrivateStructOnce.Do(func() {
		geolocationManagerPrivateStruct = gi.StructNew("WebKit2", "GeolocationManagerPrivate")
	})
}

type GeolocationManagerPrivate struct {
	native uintptr
}

var geolocationPermissionRequestClassStruct *gi.Struct
var geolocationPermissionRequestClassStructOnce sync.Once

func geolocationPermissionRequestClassStructSet() {
	geolocationPermissionRequestClassStructOnce.Do(func() {
		geolocationPermissionRequestClassStruct = gi.StructNew("WebKit2", "GeolocationPermissionRequestClass")
	})
}

type GeolocationPermissionRequestClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value '_webkit_reserved0' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved1' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved2' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved3' : missing Type
}

var geolocationPermissionRequestPrivateStruct *gi.Struct
var geolocationPermissionRequestPrivateStructOnce sync.Once

func geolocationPermissionRequestPrivateStructSet() {
	geolocationPermissionRequestPrivateStructOnce.Do(func() {
		geolocationPermissionRequestPrivateStruct = gi.StructNew("WebKit2", "GeolocationPermissionRequestPrivate")
	})
}

type GeolocationPermissionRequestPrivate struct {
	native uintptr
}

var geolocationPositionStruct *gi.Struct
var geolocationPositionStructOnce sync.Once

func geolocationPositionStructSet() {
	geolocationPositionStructOnce.Do(func() {
		geolocationPositionStruct = gi.StructNew("WebKit2", "GeolocationPosition")
	})
}

type GeolocationPosition struct {
	native uintptr
}

// UNSUPPORTED : C value 'webkit_geolocation_position_new' : parameter 'latitude' of type 'gdouble' not supported

var geolocationPositionCopyFunction *gi.Function
var geolocationPositionCopyFunctionOnce sync.Once

func geolocationPositionCopyFunctionSet() {
	geolocationPositionCopyFunctionOnce.Do(func() {
		geolocationPositionCopyFunction = gi.FunctionInvokerNew("WebKit2", "copy")
	})
}

var copyGeolocationPositionInvoker *gi.Function

// Copy is a representation of the C type webkit_geolocation_position_copy.
func (recv *GeolocationPosition) Copy() *GeolocationPosition {
	if copyGeolocationPositionInvoker == nil {
		copyGeolocationPositionInvoker = gi.StructFunctionInvokerNew("WebKit2", "GeolocationPosition", "copy")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := copyGeolocationPositionInvoker.Invoke(inArgs[:], nil)

	retGo := &GeolocationPosition{native: ret.Pointer()}

	return retGo
}

var geolocationPositionFreeFunction *gi.Function
var geolocationPositionFreeFunctionOnce sync.Once

func geolocationPositionFreeFunctionSet() {
	geolocationPositionFreeFunctionOnce.Do(func() {
		geolocationPositionFreeFunction = gi.FunctionInvokerNew("WebKit2", "free")
	})
}

var freeGeolocationPositionInvoker *gi.Function

// Free is a representation of the C type webkit_geolocation_position_free.
func (recv *GeolocationPosition) Free() {
	if freeGeolocationPositionInvoker == nil {
		freeGeolocationPositionInvoker = gi.StructFunctionInvokerNew("WebKit2", "GeolocationPosition", "free")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	freeGeolocationPositionInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'webkit_geolocation_position_set_altitude' : parameter 'altitude' of type 'gdouble' not supported

// UNSUPPORTED : C value 'webkit_geolocation_position_set_altitude_accuracy' : parameter 'altitude_accuracy' of type 'gdouble' not supported

// UNSUPPORTED : C value 'webkit_geolocation_position_set_heading' : parameter 'heading' of type 'gdouble' not supported

// UNSUPPORTED : C value 'webkit_geolocation_position_set_speed' : parameter 'speed' of type 'gdouble' not supported

var geolocationPositionSetTimestampFunction *gi.Function
var geolocationPositionSetTimestampFunctionOnce sync.Once

func geolocationPositionSetTimestampFunctionSet() {
	geolocationPositionSetTimestampFunctionOnce.Do(func() {
		geolocationPositionSetTimestampFunction = gi.FunctionInvokerNew("WebKit2", "set_timestamp")
	})
}

var setTimestampGeolocationPositionInvoker *gi.Function

// SetTimestamp is a representation of the C type webkit_geolocation_position_set_timestamp.
func (recv *GeolocationPosition) SetTimestamp(timestamp uint64) {
	if setTimestampGeolocationPositionInvoker == nil {
		setTimestampGeolocationPositionInvoker = gi.StructFunctionInvokerNew("WebKit2", "GeolocationPosition", "set_timestamp")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetUint64(timestamp)

	setTimestampGeolocationPositionInvoker.Invoke(inArgs[:], nil)

}

var hitTestResultClassStruct *gi.Struct
var hitTestResultClassStructOnce sync.Once

func hitTestResultClassStructSet() {
	hitTestResultClassStructOnce.Do(func() {
		hitTestResultClassStruct = gi.StructNew("WebKit2", "HitTestResultClass")
	})
}

type HitTestResultClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value '_webkit_reserved0' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved1' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved2' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved3' : missing Type
}

var hitTestResultPrivateStruct *gi.Struct
var hitTestResultPrivateStructOnce sync.Once

func hitTestResultPrivateStructSet() {
	hitTestResultPrivateStructOnce.Do(func() {
		hitTestResultPrivateStruct = gi.StructNew("WebKit2", "HitTestResultPrivate")
	})
}

type HitTestResultPrivate struct {
	native uintptr
}

var installMissingMediaPluginsPermissionRequestClassStruct *gi.Struct
var installMissingMediaPluginsPermissionRequestClassStructOnce sync.Once

func installMissingMediaPluginsPermissionRequestClassStructSet() {
	installMissingMediaPluginsPermissionRequestClassStructOnce.Do(func() {
		installMissingMediaPluginsPermissionRequestClassStruct = gi.StructNew("WebKit2", "InstallMissingMediaPluginsPermissionRequestClass")
	})
}

type InstallMissingMediaPluginsPermissionRequestClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value '_webkit_reserved0' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved1' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved2' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved3' : missing Type
}

var installMissingMediaPluginsPermissionRequestPrivateStruct *gi.Struct
var installMissingMediaPluginsPermissionRequestPrivateStructOnce sync.Once

func installMissingMediaPluginsPermissionRequestPrivateStructSet() {
	installMissingMediaPluginsPermissionRequestPrivateStructOnce.Do(func() {
		installMissingMediaPluginsPermissionRequestPrivateStruct = gi.StructNew("WebKit2", "InstallMissingMediaPluginsPermissionRequestPrivate")
	})
}

type InstallMissingMediaPluginsPermissionRequestPrivate struct {
	native uintptr
}

var javascriptResultStruct *gi.Struct
var javascriptResultStructOnce sync.Once

func javascriptResultStructSet() {
	javascriptResultStructOnce.Do(func() {
		javascriptResultStruct = gi.StructNew("WebKit2", "JavascriptResult")
	})
}

type JavascriptResult struct {
	native uintptr
}

// UNSUPPORTED : C value 'webkit_javascript_result_get_global_context' : return type 'JavaScriptCore.GlobalContextRef' not supported

// UNSUPPORTED : C value 'webkit_javascript_result_get_js_value' : return type 'JavaScriptCore.Value' not supported

// UNSUPPORTED : C value 'webkit_javascript_result_get_value' : return type 'JavaScriptCore.ValueRef' not supported

var javascriptResultRefFunction *gi.Function
var javascriptResultRefFunctionOnce sync.Once

func javascriptResultRefFunctionSet() {
	javascriptResultRefFunctionOnce.Do(func() {
		javascriptResultRefFunction = gi.FunctionInvokerNew("WebKit2", "ref")
	})
}

var refJavascriptResultInvoker *gi.Function

// Ref is a representation of the C type webkit_javascript_result_ref.
func (recv *JavascriptResult) Ref() *JavascriptResult {
	if refJavascriptResultInvoker == nil {
		refJavascriptResultInvoker = gi.StructFunctionInvokerNew("WebKit2", "JavascriptResult", "ref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := refJavascriptResultInvoker.Invoke(inArgs[:], nil)

	retGo := &JavascriptResult{native: ret.Pointer()}

	return retGo
}

var javascriptResultUnrefFunction *gi.Function
var javascriptResultUnrefFunctionOnce sync.Once

func javascriptResultUnrefFunctionSet() {
	javascriptResultUnrefFunctionOnce.Do(func() {
		javascriptResultUnrefFunction = gi.FunctionInvokerNew("WebKit2", "unref")
	})
}

var unrefJavascriptResultInvoker *gi.Function

// Unref is a representation of the C type webkit_javascript_result_unref.
func (recv *JavascriptResult) Unref() {
	if unrefJavascriptResultInvoker == nil {
		unrefJavascriptResultInvoker = gi.StructFunctionInvokerNew("WebKit2", "JavascriptResult", "unref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	unrefJavascriptResultInvoker.Invoke(inArgs[:], nil)

}

var mimeInfoStruct *gi.Struct
var mimeInfoStructOnce sync.Once

func mimeInfoStructSet() {
	mimeInfoStructOnce.Do(func() {
		mimeInfoStruct = gi.StructNew("WebKit2", "MimeInfo")
	})
}

type MimeInfo struct {
	native uintptr
}

var mimeInfoGetDescriptionFunction *gi.Function
var mimeInfoGetDescriptionFunctionOnce sync.Once

func mimeInfoGetDescriptionFunctionSet() {
	mimeInfoGetDescriptionFunctionOnce.Do(func() {
		mimeInfoGetDescriptionFunction = gi.FunctionInvokerNew("WebKit2", "get_description")
	})
}

var getDescriptionMimeInfoInvoker *gi.Function

// GetDescription is a representation of the C type webkit_mime_info_get_description.
func (recv *MimeInfo) GetDescription() string {
	if getDescriptionMimeInfoInvoker == nil {
		getDescriptionMimeInfoInvoker = gi.StructFunctionInvokerNew("WebKit2", "MimeInfo", "get_description")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getDescriptionMimeInfoInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

var mimeInfoGetExtensionsFunction *gi.Function
var mimeInfoGetExtensionsFunctionOnce sync.Once

func mimeInfoGetExtensionsFunctionSet() {
	mimeInfoGetExtensionsFunctionOnce.Do(func() {
		mimeInfoGetExtensionsFunction = gi.FunctionInvokerNew("WebKit2", "get_extensions")
	})
}

var getExtensionsMimeInfoInvoker *gi.Function

// GetExtensions is a representation of the C type webkit_mime_info_get_extensions.
func (recv *MimeInfo) GetExtensions() {
	if getExtensionsMimeInfoInvoker == nil {
		getExtensionsMimeInfoInvoker = gi.StructFunctionInvokerNew("WebKit2", "MimeInfo", "get_extensions")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	getExtensionsMimeInfoInvoker.Invoke(inArgs[:], nil)

}

var mimeInfoGetMimeTypeFunction *gi.Function
var mimeInfoGetMimeTypeFunctionOnce sync.Once

func mimeInfoGetMimeTypeFunctionSet() {
	mimeInfoGetMimeTypeFunctionOnce.Do(func() {
		mimeInfoGetMimeTypeFunction = gi.FunctionInvokerNew("WebKit2", "get_mime_type")
	})
}

var getMimeTypeMimeInfoInvoker *gi.Function

// GetMimeType is a representation of the C type webkit_mime_info_get_mime_type.
func (recv *MimeInfo) GetMimeType() string {
	if getMimeTypeMimeInfoInvoker == nil {
		getMimeTypeMimeInfoInvoker = gi.StructFunctionInvokerNew("WebKit2", "MimeInfo", "get_mime_type")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getMimeTypeMimeInfoInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

var mimeInfoRefFunction *gi.Function
var mimeInfoRefFunctionOnce sync.Once

func mimeInfoRefFunctionSet() {
	mimeInfoRefFunctionOnce.Do(func() {
		mimeInfoRefFunction = gi.FunctionInvokerNew("WebKit2", "ref")
	})
}

var refMimeInfoInvoker *gi.Function

// Ref is a representation of the C type webkit_mime_info_ref.
func (recv *MimeInfo) Ref() *MimeInfo {
	if refMimeInfoInvoker == nil {
		refMimeInfoInvoker = gi.StructFunctionInvokerNew("WebKit2", "MimeInfo", "ref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := refMimeInfoInvoker.Invoke(inArgs[:], nil)

	retGo := &MimeInfo{native: ret.Pointer()}

	return retGo
}

var mimeInfoUnrefFunction *gi.Function
var mimeInfoUnrefFunctionOnce sync.Once

func mimeInfoUnrefFunctionSet() {
	mimeInfoUnrefFunctionOnce.Do(func() {
		mimeInfoUnrefFunction = gi.FunctionInvokerNew("WebKit2", "unref")
	})
}

var unrefMimeInfoInvoker *gi.Function

// Unref is a representation of the C type webkit_mime_info_unref.
func (recv *MimeInfo) Unref() {
	if unrefMimeInfoInvoker == nil {
		unrefMimeInfoInvoker = gi.StructFunctionInvokerNew("WebKit2", "MimeInfo", "unref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	unrefMimeInfoInvoker.Invoke(inArgs[:], nil)

}

var navigationActionStruct *gi.Struct
var navigationActionStructOnce sync.Once

func navigationActionStructSet() {
	navigationActionStructOnce.Do(func() {
		navigationActionStruct = gi.StructNew("WebKit2", "NavigationAction")
	})
}

type NavigationAction struct {
	native uintptr
}

var navigationActionCopyFunction *gi.Function
var navigationActionCopyFunctionOnce sync.Once

func navigationActionCopyFunctionSet() {
	navigationActionCopyFunctionOnce.Do(func() {
		navigationActionCopyFunction = gi.FunctionInvokerNew("WebKit2", "copy")
	})
}

var copyNavigationActionInvoker *gi.Function

// Copy is a representation of the C type webkit_navigation_action_copy.
func (recv *NavigationAction) Copy() *NavigationAction {
	if copyNavigationActionInvoker == nil {
		copyNavigationActionInvoker = gi.StructFunctionInvokerNew("WebKit2", "NavigationAction", "copy")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := copyNavigationActionInvoker.Invoke(inArgs[:], nil)

	retGo := &NavigationAction{native: ret.Pointer()}

	return retGo
}

var navigationActionFreeFunction *gi.Function
var navigationActionFreeFunctionOnce sync.Once

func navigationActionFreeFunctionSet() {
	navigationActionFreeFunctionOnce.Do(func() {
		navigationActionFreeFunction = gi.FunctionInvokerNew("WebKit2", "free")
	})
}

var freeNavigationActionInvoker *gi.Function

// Free is a representation of the C type webkit_navigation_action_free.
func (recv *NavigationAction) Free() {
	if freeNavigationActionInvoker == nil {
		freeNavigationActionInvoker = gi.StructFunctionInvokerNew("WebKit2", "NavigationAction", "free")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	freeNavigationActionInvoker.Invoke(inArgs[:], nil)

}

var navigationActionGetModifiersFunction *gi.Function
var navigationActionGetModifiersFunctionOnce sync.Once

func navigationActionGetModifiersFunctionSet() {
	navigationActionGetModifiersFunctionOnce.Do(func() {
		navigationActionGetModifiersFunction = gi.FunctionInvokerNew("WebKit2", "get_modifiers")
	})
}

var getModifiersNavigationActionInvoker *gi.Function

// GetModifiers is a representation of the C type webkit_navigation_action_get_modifiers.
func (recv *NavigationAction) GetModifiers() uint32 {
	if getModifiersNavigationActionInvoker == nil {
		getModifiersNavigationActionInvoker = gi.StructFunctionInvokerNew("WebKit2", "NavigationAction", "get_modifiers")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getModifiersNavigationActionInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Uint32()

	return retGo
}

var navigationActionGetMouseButtonFunction *gi.Function
var navigationActionGetMouseButtonFunctionOnce sync.Once

func navigationActionGetMouseButtonFunctionSet() {
	navigationActionGetMouseButtonFunctionOnce.Do(func() {
		navigationActionGetMouseButtonFunction = gi.FunctionInvokerNew("WebKit2", "get_mouse_button")
	})
}

var getMouseButtonNavigationActionInvoker *gi.Function

// GetMouseButton is a representation of the C type webkit_navigation_action_get_mouse_button.
func (recv *NavigationAction) GetMouseButton() uint32 {
	if getMouseButtonNavigationActionInvoker == nil {
		getMouseButtonNavigationActionInvoker = gi.StructFunctionInvokerNew("WebKit2", "NavigationAction", "get_mouse_button")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getMouseButtonNavigationActionInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Uint32()

	return retGo
}

// UNSUPPORTED : C value 'webkit_navigation_action_get_navigation_type' : return type 'NavigationType' not supported

// UNSUPPORTED : C value 'webkit_navigation_action_get_request' : return type 'URIRequest' not supported

// UNSUPPORTED : C value 'webkit_navigation_action_is_redirect' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'webkit_navigation_action_is_user_gesture' : return type 'gboolean' not supported

var navigationPolicyDecisionClassStruct *gi.Struct
var navigationPolicyDecisionClassStructOnce sync.Once

func navigationPolicyDecisionClassStructSet() {
	navigationPolicyDecisionClassStructOnce.Do(func() {
		navigationPolicyDecisionClassStruct = gi.StructNew("WebKit2", "NavigationPolicyDecisionClass")
	})
}

type NavigationPolicyDecisionClass struct {
	native      uintptr
	ParentClass *PolicyDecisionClass
	// UNSUPPORTED : C value '_webkit_reserved0' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved1' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved2' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved3' : missing Type
}

var navigationPolicyDecisionPrivateStruct *gi.Struct
var navigationPolicyDecisionPrivateStructOnce sync.Once

func navigationPolicyDecisionPrivateStructSet() {
	navigationPolicyDecisionPrivateStructOnce.Do(func() {
		navigationPolicyDecisionPrivateStruct = gi.StructNew("WebKit2", "NavigationPolicyDecisionPrivate")
	})
}

type NavigationPolicyDecisionPrivate struct {
	native uintptr
}

var networkProxySettingsStruct *gi.Struct
var networkProxySettingsStructOnce sync.Once

func networkProxySettingsStructSet() {
	networkProxySettingsStructOnce.Do(func() {
		networkProxySettingsStruct = gi.StructNew("WebKit2", "NetworkProxySettings")
	})
}

type NetworkProxySettings struct {
	native uintptr
}

// UNSUPPORTED : C value 'webkit_network_proxy_settings_new' : parameter 'ignore_hosts' has no type

var networkProxySettingsAddProxyForSchemeFunction *gi.Function
var networkProxySettingsAddProxyForSchemeFunctionOnce sync.Once

func networkProxySettingsAddProxyForSchemeFunctionSet() {
	networkProxySettingsAddProxyForSchemeFunctionOnce.Do(func() {
		networkProxySettingsAddProxyForSchemeFunction = gi.FunctionInvokerNew("WebKit2", "add_proxy_for_scheme")
	})
}

var addProxyForSchemeNetworkProxySettingsInvoker *gi.Function

// AddProxyForScheme is a representation of the C type webkit_network_proxy_settings_add_proxy_for_scheme.
func (recv *NetworkProxySettings) AddProxyForScheme(scheme string, proxyUri string) {
	if addProxyForSchemeNetworkProxySettingsInvoker == nil {
		addProxyForSchemeNetworkProxySettingsInvoker = gi.StructFunctionInvokerNew("WebKit2", "NetworkProxySettings", "add_proxy_for_scheme")
	}

	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(scheme)
	inArgs[2].SetString(proxyUri)

	addProxyForSchemeNetworkProxySettingsInvoker.Invoke(inArgs[:], nil)

}

var networkProxySettingsCopyFunction *gi.Function
var networkProxySettingsCopyFunctionOnce sync.Once

func networkProxySettingsCopyFunctionSet() {
	networkProxySettingsCopyFunctionOnce.Do(func() {
		networkProxySettingsCopyFunction = gi.FunctionInvokerNew("WebKit2", "copy")
	})
}

var copyNetworkProxySettingsInvoker *gi.Function

// Copy is a representation of the C type webkit_network_proxy_settings_copy.
func (recv *NetworkProxySettings) Copy() *NetworkProxySettings {
	if copyNetworkProxySettingsInvoker == nil {
		copyNetworkProxySettingsInvoker = gi.StructFunctionInvokerNew("WebKit2", "NetworkProxySettings", "copy")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := copyNetworkProxySettingsInvoker.Invoke(inArgs[:], nil)

	retGo := &NetworkProxySettings{native: ret.Pointer()}

	return retGo
}

var networkProxySettingsFreeFunction *gi.Function
var networkProxySettingsFreeFunctionOnce sync.Once

func networkProxySettingsFreeFunctionSet() {
	networkProxySettingsFreeFunctionOnce.Do(func() {
		networkProxySettingsFreeFunction = gi.FunctionInvokerNew("WebKit2", "free")
	})
}

var freeNetworkProxySettingsInvoker *gi.Function

// Free is a representation of the C type webkit_network_proxy_settings_free.
func (recv *NetworkProxySettings) Free() {
	if freeNetworkProxySettingsInvoker == nil {
		freeNetworkProxySettingsInvoker = gi.StructFunctionInvokerNew("WebKit2", "NetworkProxySettings", "free")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	freeNetworkProxySettingsInvoker.Invoke(inArgs[:], nil)

}

var notificationClassStruct *gi.Struct
var notificationClassStructOnce sync.Once

func notificationClassStructSet() {
	notificationClassStructOnce.Do(func() {
		notificationClassStruct = gi.StructNew("WebKit2", "NotificationClass")
	})
}

type NotificationClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value '_webkit_reserved0' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved1' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved2' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved3' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved4' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved5' : missing Type
}

var notificationPermissionRequestClassStruct *gi.Struct
var notificationPermissionRequestClassStructOnce sync.Once

func notificationPermissionRequestClassStructSet() {
	notificationPermissionRequestClassStructOnce.Do(func() {
		notificationPermissionRequestClassStruct = gi.StructNew("WebKit2", "NotificationPermissionRequestClass")
	})
}

type NotificationPermissionRequestClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
}

var notificationPermissionRequestPrivateStruct *gi.Struct
var notificationPermissionRequestPrivateStructOnce sync.Once

func notificationPermissionRequestPrivateStructSet() {
	notificationPermissionRequestPrivateStructOnce.Do(func() {
		notificationPermissionRequestPrivateStruct = gi.StructNew("WebKit2", "NotificationPermissionRequestPrivate")
	})
}

type NotificationPermissionRequestPrivate struct {
	native uintptr
}

var notificationPrivateStruct *gi.Struct
var notificationPrivateStructOnce sync.Once

func notificationPrivateStructSet() {
	notificationPrivateStructOnce.Do(func() {
		notificationPrivateStruct = gi.StructNew("WebKit2", "NotificationPrivate")
	})
}

type NotificationPrivate struct {
	native uintptr
}

var optionMenuClassStruct *gi.Struct
var optionMenuClassStructOnce sync.Once

func optionMenuClassStructSet() {
	optionMenuClassStructOnce.Do(func() {
		optionMenuClassStruct = gi.StructNew("WebKit2", "OptionMenuClass")
	})
}

type OptionMenuClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value '_webkit_reserved0' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved1' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved2' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved3' : missing Type
}

var optionMenuItemStruct *gi.Struct
var optionMenuItemStructOnce sync.Once

func optionMenuItemStructSet() {
	optionMenuItemStructOnce.Do(func() {
		optionMenuItemStruct = gi.StructNew("WebKit2", "OptionMenuItem")
	})
}

type OptionMenuItem struct {
	native uintptr
}

var optionMenuItemCopyFunction *gi.Function
var optionMenuItemCopyFunctionOnce sync.Once

func optionMenuItemCopyFunctionSet() {
	optionMenuItemCopyFunctionOnce.Do(func() {
		optionMenuItemCopyFunction = gi.FunctionInvokerNew("WebKit2", "copy")
	})
}

var copyOptionMenuItemInvoker *gi.Function

// Copy is a representation of the C type webkit_option_menu_item_copy.
func (recv *OptionMenuItem) Copy() *OptionMenuItem {
	if copyOptionMenuItemInvoker == nil {
		copyOptionMenuItemInvoker = gi.StructFunctionInvokerNew("WebKit2", "OptionMenuItem", "copy")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := copyOptionMenuItemInvoker.Invoke(inArgs[:], nil)

	retGo := &OptionMenuItem{native: ret.Pointer()}

	return retGo
}

var optionMenuItemFreeFunction *gi.Function
var optionMenuItemFreeFunctionOnce sync.Once

func optionMenuItemFreeFunctionSet() {
	optionMenuItemFreeFunctionOnce.Do(func() {
		optionMenuItemFreeFunction = gi.FunctionInvokerNew("WebKit2", "free")
	})
}

var freeOptionMenuItemInvoker *gi.Function

// Free is a representation of the C type webkit_option_menu_item_free.
func (recv *OptionMenuItem) Free() {
	if freeOptionMenuItemInvoker == nil {
		freeOptionMenuItemInvoker = gi.StructFunctionInvokerNew("WebKit2", "OptionMenuItem", "free")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	freeOptionMenuItemInvoker.Invoke(inArgs[:], nil)

}

var optionMenuItemGetLabelFunction *gi.Function
var optionMenuItemGetLabelFunctionOnce sync.Once

func optionMenuItemGetLabelFunctionSet() {
	optionMenuItemGetLabelFunctionOnce.Do(func() {
		optionMenuItemGetLabelFunction = gi.FunctionInvokerNew("WebKit2", "get_label")
	})
}

var getLabelOptionMenuItemInvoker *gi.Function

// GetLabel is a representation of the C type webkit_option_menu_item_get_label.
func (recv *OptionMenuItem) GetLabel() string {
	if getLabelOptionMenuItemInvoker == nil {
		getLabelOptionMenuItemInvoker = gi.StructFunctionInvokerNew("WebKit2", "OptionMenuItem", "get_label")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getLabelOptionMenuItemInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

var optionMenuItemGetTooltipFunction *gi.Function
var optionMenuItemGetTooltipFunctionOnce sync.Once

func optionMenuItemGetTooltipFunctionSet() {
	optionMenuItemGetTooltipFunctionOnce.Do(func() {
		optionMenuItemGetTooltipFunction = gi.FunctionInvokerNew("WebKit2", "get_tooltip")
	})
}

var getTooltipOptionMenuItemInvoker *gi.Function

// GetTooltip is a representation of the C type webkit_option_menu_item_get_tooltip.
func (recv *OptionMenuItem) GetTooltip() string {
	if getTooltipOptionMenuItemInvoker == nil {
		getTooltipOptionMenuItemInvoker = gi.StructFunctionInvokerNew("WebKit2", "OptionMenuItem", "get_tooltip")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getTooltipOptionMenuItemInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'webkit_option_menu_item_is_enabled' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'webkit_option_menu_item_is_group_child' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'webkit_option_menu_item_is_group_label' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'webkit_option_menu_item_is_selected' : return type 'gboolean' not supported

var optionMenuPrivateStruct *gi.Struct
var optionMenuPrivateStructOnce sync.Once

func optionMenuPrivateStructSet() {
	optionMenuPrivateStructOnce.Do(func() {
		optionMenuPrivateStruct = gi.StructNew("WebKit2", "OptionMenuPrivate")
	})
}

type OptionMenuPrivate struct {
	native uintptr
}

var permissionRequestIfaceStruct *gi.Struct
var permissionRequestIfaceStructOnce sync.Once

func permissionRequestIfaceStructSet() {
	permissionRequestIfaceStructOnce.Do(func() {
		permissionRequestIfaceStruct = gi.StructNew("WebKit2", "PermissionRequestIface")
	})
}

type PermissionRequestIface struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_interface' : no Go type for 'GObject.TypeInterface'
	// UNSUPPORTED : C value 'allow' : missing Type
	// UNSUPPORTED : C value 'deny' : missing Type
}

var pluginClassStruct *gi.Struct
var pluginClassStructOnce sync.Once

func pluginClassStructSet() {
	pluginClassStructOnce.Do(func() {
		pluginClassStruct = gi.StructNew("WebKit2", "PluginClass")
	})
}

type PluginClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value '_webkit_reserved0' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved1' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved2' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved3' : missing Type
}

var pluginPrivateStruct *gi.Struct
var pluginPrivateStructOnce sync.Once

func pluginPrivateStructSet() {
	pluginPrivateStructOnce.Do(func() {
		pluginPrivateStruct = gi.StructNew("WebKit2", "PluginPrivate")
	})
}

type PluginPrivate struct {
	native uintptr
}

var policyDecisionClassStruct *gi.Struct
var policyDecisionClassStructOnce sync.Once

func policyDecisionClassStructSet() {
	policyDecisionClassStructOnce.Do(func() {
		policyDecisionClassStruct = gi.StructNew("WebKit2", "PolicyDecisionClass")
	})
}

type PolicyDecisionClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value '_webkit_reserved0' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved1' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved2' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved3' : missing Type
}

var policyDecisionPrivateStruct *gi.Struct
var policyDecisionPrivateStructOnce sync.Once

func policyDecisionPrivateStructSet() {
	policyDecisionPrivateStructOnce.Do(func() {
		policyDecisionPrivateStruct = gi.StructNew("WebKit2", "PolicyDecisionPrivate")
	})
}

type PolicyDecisionPrivate struct {
	native uintptr
}

var printCustomWidgetClassStruct *gi.Struct
var printCustomWidgetClassStructOnce sync.Once

func printCustomWidgetClassStructSet() {
	printCustomWidgetClassStructOnce.Do(func() {
		printCustomWidgetClassStruct = gi.StructNew("WebKit2", "PrintCustomWidgetClass")
	})
}

type PrintCustomWidgetClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value 'apply' : missing Type
	// UNSUPPORTED : C value 'update' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved0' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved1' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved2' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved3' : missing Type
}

var printCustomWidgetPrivateStruct *gi.Struct
var printCustomWidgetPrivateStructOnce sync.Once

func printCustomWidgetPrivateStructSet() {
	printCustomWidgetPrivateStructOnce.Do(func() {
		printCustomWidgetPrivateStruct = gi.StructNew("WebKit2", "PrintCustomWidgetPrivate")
	})
}

type PrintCustomWidgetPrivate struct {
	native uintptr
}

var printOperationClassStruct *gi.Struct
var printOperationClassStructOnce sync.Once

func printOperationClassStructSet() {
	printOperationClassStructOnce.Do(func() {
		printOperationClassStruct = gi.StructNew("WebKit2", "PrintOperationClass")
	})
}

type PrintOperationClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value '_webkit_reserved0' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved1' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved2' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved3' : missing Type
}

var printOperationPrivateStruct *gi.Struct
var printOperationPrivateStructOnce sync.Once

func printOperationPrivateStructSet() {
	printOperationPrivateStructOnce.Do(func() {
		printOperationPrivateStruct = gi.StructNew("WebKit2", "PrintOperationPrivate")
	})
}

type PrintOperationPrivate struct {
	native uintptr
}

var responsePolicyDecisionClassStruct *gi.Struct
var responsePolicyDecisionClassStructOnce sync.Once

func responsePolicyDecisionClassStructSet() {
	responsePolicyDecisionClassStructOnce.Do(func() {
		responsePolicyDecisionClassStruct = gi.StructNew("WebKit2", "ResponsePolicyDecisionClass")
	})
}

type ResponsePolicyDecisionClass struct {
	native      uintptr
	ParentClass *PolicyDecisionClass
	// UNSUPPORTED : C value '_webkit_reserved0' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved1' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved2' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved3' : missing Type
}

var responsePolicyDecisionPrivateStruct *gi.Struct
var responsePolicyDecisionPrivateStructOnce sync.Once

func responsePolicyDecisionPrivateStructSet() {
	responsePolicyDecisionPrivateStructOnce.Do(func() {
		responsePolicyDecisionPrivateStruct = gi.StructNew("WebKit2", "ResponsePolicyDecisionPrivate")
	})
}

type ResponsePolicyDecisionPrivate struct {
	native uintptr
}

var scriptDialogStruct *gi.Struct
var scriptDialogStructOnce sync.Once

func scriptDialogStructSet() {
	scriptDialogStructOnce.Do(func() {
		scriptDialogStruct = gi.StructNew("WebKit2", "ScriptDialog")
	})
}

type ScriptDialog struct {
	native uintptr
}

var scriptDialogCloseFunction *gi.Function
var scriptDialogCloseFunctionOnce sync.Once

func scriptDialogCloseFunctionSet() {
	scriptDialogCloseFunctionOnce.Do(func() {
		scriptDialogCloseFunction = gi.FunctionInvokerNew("WebKit2", "close")
	})
}

var closeScriptDialogInvoker *gi.Function

// Close is a representation of the C type webkit_script_dialog_close.
func (recv *ScriptDialog) Close() {
	if closeScriptDialogInvoker == nil {
		closeScriptDialogInvoker = gi.StructFunctionInvokerNew("WebKit2", "ScriptDialog", "close")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	closeScriptDialogInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'webkit_script_dialog_confirm_set_confirmed' : parameter 'confirmed' of type 'gboolean' not supported

// UNSUPPORTED : C value 'webkit_script_dialog_get_dialog_type' : return type 'ScriptDialogType' not supported

var scriptDialogGetMessageFunction *gi.Function
var scriptDialogGetMessageFunctionOnce sync.Once

func scriptDialogGetMessageFunctionSet() {
	scriptDialogGetMessageFunctionOnce.Do(func() {
		scriptDialogGetMessageFunction = gi.FunctionInvokerNew("WebKit2", "get_message")
	})
}

var getMessageScriptDialogInvoker *gi.Function

// GetMessage is a representation of the C type webkit_script_dialog_get_message.
func (recv *ScriptDialog) GetMessage() string {
	if getMessageScriptDialogInvoker == nil {
		getMessageScriptDialogInvoker = gi.StructFunctionInvokerNew("WebKit2", "ScriptDialog", "get_message")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getMessageScriptDialogInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

var scriptDialogPromptGetDefaultTextFunction *gi.Function
var scriptDialogPromptGetDefaultTextFunctionOnce sync.Once

func scriptDialogPromptGetDefaultTextFunctionSet() {
	scriptDialogPromptGetDefaultTextFunctionOnce.Do(func() {
		scriptDialogPromptGetDefaultTextFunction = gi.FunctionInvokerNew("WebKit2", "prompt_get_default_text")
	})
}

var promptGetDefaultTextScriptDialogInvoker *gi.Function

// PromptGetDefaultText is a representation of the C type webkit_script_dialog_prompt_get_default_text.
func (recv *ScriptDialog) PromptGetDefaultText() string {
	if promptGetDefaultTextScriptDialogInvoker == nil {
		promptGetDefaultTextScriptDialogInvoker = gi.StructFunctionInvokerNew("WebKit2", "ScriptDialog", "prompt_get_default_text")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := promptGetDefaultTextScriptDialogInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

var scriptDialogPromptSetTextFunction *gi.Function
var scriptDialogPromptSetTextFunctionOnce sync.Once

func scriptDialogPromptSetTextFunctionSet() {
	scriptDialogPromptSetTextFunctionOnce.Do(func() {
		scriptDialogPromptSetTextFunction = gi.FunctionInvokerNew("WebKit2", "prompt_set_text")
	})
}

var promptSetTextScriptDialogInvoker *gi.Function

// PromptSetText is a representation of the C type webkit_script_dialog_prompt_set_text.
func (recv *ScriptDialog) PromptSetText(text string) {
	if promptSetTextScriptDialogInvoker == nil {
		promptSetTextScriptDialogInvoker = gi.StructFunctionInvokerNew("WebKit2", "ScriptDialog", "prompt_set_text")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(text)

	promptSetTextScriptDialogInvoker.Invoke(inArgs[:], nil)

}

var scriptDialogRefFunction *gi.Function
var scriptDialogRefFunctionOnce sync.Once

func scriptDialogRefFunctionSet() {
	scriptDialogRefFunctionOnce.Do(func() {
		scriptDialogRefFunction = gi.FunctionInvokerNew("WebKit2", "ref")
	})
}

var refScriptDialogInvoker *gi.Function

// Ref is a representation of the C type webkit_script_dialog_ref.
func (recv *ScriptDialog) Ref() *ScriptDialog {
	if refScriptDialogInvoker == nil {
		refScriptDialogInvoker = gi.StructFunctionInvokerNew("WebKit2", "ScriptDialog", "ref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := refScriptDialogInvoker.Invoke(inArgs[:], nil)

	retGo := &ScriptDialog{native: ret.Pointer()}

	return retGo
}

var scriptDialogUnrefFunction *gi.Function
var scriptDialogUnrefFunctionOnce sync.Once

func scriptDialogUnrefFunctionSet() {
	scriptDialogUnrefFunctionOnce.Do(func() {
		scriptDialogUnrefFunction = gi.FunctionInvokerNew("WebKit2", "unref")
	})
}

var unrefScriptDialogInvoker *gi.Function

// Unref is a representation of the C type webkit_script_dialog_unref.
func (recv *ScriptDialog) Unref() {
	if unrefScriptDialogInvoker == nil {
		unrefScriptDialogInvoker = gi.StructFunctionInvokerNew("WebKit2", "ScriptDialog", "unref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	unrefScriptDialogInvoker.Invoke(inArgs[:], nil)

}

var securityManagerClassStruct *gi.Struct
var securityManagerClassStructOnce sync.Once

func securityManagerClassStructSet() {
	securityManagerClassStructOnce.Do(func() {
		securityManagerClassStruct = gi.StructNew("WebKit2", "SecurityManagerClass")
	})
}

type SecurityManagerClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value '_webkit_reserved0' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved1' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved2' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved3' : missing Type
}

var securityManagerPrivateStruct *gi.Struct
var securityManagerPrivateStructOnce sync.Once

func securityManagerPrivateStructSet() {
	securityManagerPrivateStructOnce.Do(func() {
		securityManagerPrivateStruct = gi.StructNew("WebKit2", "SecurityManagerPrivate")
	})
}

type SecurityManagerPrivate struct {
	native uintptr
}

var securityOriginStruct *gi.Struct
var securityOriginStructOnce sync.Once

func securityOriginStructSet() {
	securityOriginStructOnce.Do(func() {
		securityOriginStruct = gi.StructNew("WebKit2", "SecurityOrigin")
	})
}

type SecurityOrigin struct {
	native uintptr
}

var securityOriginNewFunction *gi.Function
var securityOriginNewFunctionOnce sync.Once

func securityOriginNewFunctionSet() {
	securityOriginNewFunctionOnce.Do(func() {
		securityOriginNewFunction = gi.FunctionInvokerNew("WebKit2", "new")
	})
}

var newSecurityOriginInvoker *gi.Function

// SecurityOriginNew is a representation of the C type webkit_security_origin_new.
func SecurityOriginNew(protocol string, host string, port uint16) *SecurityOrigin {
	if newSecurityOriginInvoker == nil {
		newSecurityOriginInvoker = gi.StructFunctionInvokerNew("WebKit2", "SecurityOrigin", "new")
	}

	var inArgs [3]gi.Argument
	inArgs[0].SetString(protocol)
	inArgs[1].SetString(host)
	inArgs[2].SetUint16(port)

	ret := newSecurityOriginInvoker.Invoke(inArgs[:], nil)

	retGo := &SecurityOrigin{native: ret.Pointer()}

	return retGo
}

var securityOriginNewForUriFunction *gi.Function
var securityOriginNewForUriFunctionOnce sync.Once

func securityOriginNewForUriFunctionSet() {
	securityOriginNewForUriFunctionOnce.Do(func() {
		securityOriginNewForUriFunction = gi.FunctionInvokerNew("WebKit2", "new_for_uri")
	})
}

var newForUriSecurityOriginInvoker *gi.Function

// SecurityOriginNewForUri is a representation of the C type webkit_security_origin_new_for_uri.
func SecurityOriginNewForUri(uri string) *SecurityOrigin {
	if newForUriSecurityOriginInvoker == nil {
		newForUriSecurityOriginInvoker = gi.StructFunctionInvokerNew("WebKit2", "SecurityOrigin", "new_for_uri")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetString(uri)

	ret := newForUriSecurityOriginInvoker.Invoke(inArgs[:], nil)

	retGo := &SecurityOrigin{native: ret.Pointer()}

	return retGo
}

var securityOriginGetHostFunction *gi.Function
var securityOriginGetHostFunctionOnce sync.Once

func securityOriginGetHostFunctionSet() {
	securityOriginGetHostFunctionOnce.Do(func() {
		securityOriginGetHostFunction = gi.FunctionInvokerNew("WebKit2", "get_host")
	})
}

var getHostSecurityOriginInvoker *gi.Function

// GetHost is a representation of the C type webkit_security_origin_get_host.
func (recv *SecurityOrigin) GetHost() string {
	if getHostSecurityOriginInvoker == nil {
		getHostSecurityOriginInvoker = gi.StructFunctionInvokerNew("WebKit2", "SecurityOrigin", "get_host")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getHostSecurityOriginInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

var securityOriginGetPortFunction *gi.Function
var securityOriginGetPortFunctionOnce sync.Once

func securityOriginGetPortFunctionSet() {
	securityOriginGetPortFunctionOnce.Do(func() {
		securityOriginGetPortFunction = gi.FunctionInvokerNew("WebKit2", "get_port")
	})
}

var getPortSecurityOriginInvoker *gi.Function

// GetPort is a representation of the C type webkit_security_origin_get_port.
func (recv *SecurityOrigin) GetPort() uint16 {
	if getPortSecurityOriginInvoker == nil {
		getPortSecurityOriginInvoker = gi.StructFunctionInvokerNew("WebKit2", "SecurityOrigin", "get_port")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getPortSecurityOriginInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Uint16()

	return retGo
}

var securityOriginGetProtocolFunction *gi.Function
var securityOriginGetProtocolFunctionOnce sync.Once

func securityOriginGetProtocolFunctionSet() {
	securityOriginGetProtocolFunctionOnce.Do(func() {
		securityOriginGetProtocolFunction = gi.FunctionInvokerNew("WebKit2", "get_protocol")
	})
}

var getProtocolSecurityOriginInvoker *gi.Function

// GetProtocol is a representation of the C type webkit_security_origin_get_protocol.
func (recv *SecurityOrigin) GetProtocol() string {
	if getProtocolSecurityOriginInvoker == nil {
		getProtocolSecurityOriginInvoker = gi.StructFunctionInvokerNew("WebKit2", "SecurityOrigin", "get_protocol")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getProtocolSecurityOriginInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'webkit_security_origin_is_opaque' : return type 'gboolean' not supported

var securityOriginRefFunction *gi.Function
var securityOriginRefFunctionOnce sync.Once

func securityOriginRefFunctionSet() {
	securityOriginRefFunctionOnce.Do(func() {
		securityOriginRefFunction = gi.FunctionInvokerNew("WebKit2", "ref")
	})
}

var refSecurityOriginInvoker *gi.Function

// Ref is a representation of the C type webkit_security_origin_ref.
func (recv *SecurityOrigin) Ref() *SecurityOrigin {
	if refSecurityOriginInvoker == nil {
		refSecurityOriginInvoker = gi.StructFunctionInvokerNew("WebKit2", "SecurityOrigin", "ref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := refSecurityOriginInvoker.Invoke(inArgs[:], nil)

	retGo := &SecurityOrigin{native: ret.Pointer()}

	return retGo
}

var securityOriginToStringFunction *gi.Function
var securityOriginToStringFunctionOnce sync.Once

func securityOriginToStringFunctionSet() {
	securityOriginToStringFunctionOnce.Do(func() {
		securityOriginToStringFunction = gi.FunctionInvokerNew("WebKit2", "to_string")
	})
}

var toStringSecurityOriginInvoker *gi.Function

// ToString is a representation of the C type webkit_security_origin_to_string.
func (recv *SecurityOrigin) ToString() string {
	if toStringSecurityOriginInvoker == nil {
		toStringSecurityOriginInvoker = gi.StructFunctionInvokerNew("WebKit2", "SecurityOrigin", "to_string")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := toStringSecurityOriginInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

var securityOriginUnrefFunction *gi.Function
var securityOriginUnrefFunctionOnce sync.Once

func securityOriginUnrefFunctionSet() {
	securityOriginUnrefFunctionOnce.Do(func() {
		securityOriginUnrefFunction = gi.FunctionInvokerNew("WebKit2", "unref")
	})
}

var unrefSecurityOriginInvoker *gi.Function

// Unref is a representation of the C type webkit_security_origin_unref.
func (recv *SecurityOrigin) Unref() {
	if unrefSecurityOriginInvoker == nil {
		unrefSecurityOriginInvoker = gi.StructFunctionInvokerNew("WebKit2", "SecurityOrigin", "unref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	unrefSecurityOriginInvoker.Invoke(inArgs[:], nil)

}

var settingsClassStruct *gi.Struct
var settingsClassStructOnce sync.Once

func settingsClassStructSet() {
	settingsClassStructOnce.Do(func() {
		settingsClassStruct = gi.StructNew("WebKit2", "SettingsClass")
	})
}

type SettingsClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value '_webkit_reserved0' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved1' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved2' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved3' : missing Type
}

var settingsPrivateStruct *gi.Struct
var settingsPrivateStructOnce sync.Once

func settingsPrivateStructSet() {
	settingsPrivateStructOnce.Do(func() {
		settingsPrivateStruct = gi.StructNew("WebKit2", "SettingsPrivate")
	})
}

type SettingsPrivate struct {
	native uintptr
}

var uRIRequestClassStruct *gi.Struct
var uRIRequestClassStructOnce sync.Once

func uRIRequestClassStructSet() {
	uRIRequestClassStructOnce.Do(func() {
		uRIRequestClassStruct = gi.StructNew("WebKit2", "URIRequestClass")
	})
}

type URIRequestClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value '_webkit_reserved0' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved1' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved2' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved3' : missing Type
}

var uRIRequestPrivateStruct *gi.Struct
var uRIRequestPrivateStructOnce sync.Once

func uRIRequestPrivateStructSet() {
	uRIRequestPrivateStructOnce.Do(func() {
		uRIRequestPrivateStruct = gi.StructNew("WebKit2", "URIRequestPrivate")
	})
}

type URIRequestPrivate struct {
	native uintptr
}

var uRIResponseClassStruct *gi.Struct
var uRIResponseClassStructOnce sync.Once

func uRIResponseClassStructSet() {
	uRIResponseClassStructOnce.Do(func() {
		uRIResponseClassStruct = gi.StructNew("WebKit2", "URIResponseClass")
	})
}

type URIResponseClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value '_webkit_reserved0' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved1' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved2' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved3' : missing Type
}

var uRIResponsePrivateStruct *gi.Struct
var uRIResponsePrivateStructOnce sync.Once

func uRIResponsePrivateStructSet() {
	uRIResponsePrivateStructOnce.Do(func() {
		uRIResponsePrivateStruct = gi.StructNew("WebKit2", "URIResponsePrivate")
	})
}

type URIResponsePrivate struct {
	native uintptr
}

var uRISchemeRequestClassStruct *gi.Struct
var uRISchemeRequestClassStructOnce sync.Once

func uRISchemeRequestClassStructSet() {
	uRISchemeRequestClassStructOnce.Do(func() {
		uRISchemeRequestClassStruct = gi.StructNew("WebKit2", "URISchemeRequestClass")
	})
}

type URISchemeRequestClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value '_webkit_reserved0' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved1' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved2' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved3' : missing Type
}

var uRISchemeRequestPrivateStruct *gi.Struct
var uRISchemeRequestPrivateStructOnce sync.Once

func uRISchemeRequestPrivateStructSet() {
	uRISchemeRequestPrivateStructOnce.Do(func() {
		uRISchemeRequestPrivateStruct = gi.StructNew("WebKit2", "URISchemeRequestPrivate")
	})
}

type URISchemeRequestPrivate struct {
	native uintptr
}

var userContentFilterStruct *gi.Struct
var userContentFilterStructOnce sync.Once

func userContentFilterStructSet() {
	userContentFilterStructOnce.Do(func() {
		userContentFilterStruct = gi.StructNew("WebKit2", "UserContentFilter")
	})
}

type UserContentFilter struct {
	native uintptr
}

var userContentFilterGetIdentifierFunction *gi.Function
var userContentFilterGetIdentifierFunctionOnce sync.Once

func userContentFilterGetIdentifierFunctionSet() {
	userContentFilterGetIdentifierFunctionOnce.Do(func() {
		userContentFilterGetIdentifierFunction = gi.FunctionInvokerNew("WebKit2", "get_identifier")
	})
}

var getIdentifierUserContentFilterInvoker *gi.Function

// GetIdentifier is a representation of the C type webkit_user_content_filter_get_identifier.
func (recv *UserContentFilter) GetIdentifier() string {
	if getIdentifierUserContentFilterInvoker == nil {
		getIdentifierUserContentFilterInvoker = gi.StructFunctionInvokerNew("WebKit2", "UserContentFilter", "get_identifier")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getIdentifierUserContentFilterInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

var userContentFilterRefFunction *gi.Function
var userContentFilterRefFunctionOnce sync.Once

func userContentFilterRefFunctionSet() {
	userContentFilterRefFunctionOnce.Do(func() {
		userContentFilterRefFunction = gi.FunctionInvokerNew("WebKit2", "ref")
	})
}

var refUserContentFilterInvoker *gi.Function

// Ref is a representation of the C type webkit_user_content_filter_ref.
func (recv *UserContentFilter) Ref() *UserContentFilter {
	if refUserContentFilterInvoker == nil {
		refUserContentFilterInvoker = gi.StructFunctionInvokerNew("WebKit2", "UserContentFilter", "ref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := refUserContentFilterInvoker.Invoke(inArgs[:], nil)

	retGo := &UserContentFilter{native: ret.Pointer()}

	return retGo
}

var userContentFilterUnrefFunction *gi.Function
var userContentFilterUnrefFunctionOnce sync.Once

func userContentFilterUnrefFunctionSet() {
	userContentFilterUnrefFunctionOnce.Do(func() {
		userContentFilterUnrefFunction = gi.FunctionInvokerNew("WebKit2", "unref")
	})
}

var unrefUserContentFilterInvoker *gi.Function

// Unref is a representation of the C type webkit_user_content_filter_unref.
func (recv *UserContentFilter) Unref() {
	if unrefUserContentFilterInvoker == nil {
		unrefUserContentFilterInvoker = gi.StructFunctionInvokerNew("WebKit2", "UserContentFilter", "unref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	unrefUserContentFilterInvoker.Invoke(inArgs[:], nil)

}

var userContentFilterStoreClassStruct *gi.Struct
var userContentFilterStoreClassStructOnce sync.Once

func userContentFilterStoreClassStructSet() {
	userContentFilterStoreClassStructOnce.Do(func() {
		userContentFilterStoreClassStruct = gi.StructNew("WebKit2", "UserContentFilterStoreClass")
	})
}

type UserContentFilterStoreClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value '_webkit_reserved0' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved1' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved2' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved3' : missing Type
}

var userContentFilterStorePrivateStruct *gi.Struct
var userContentFilterStorePrivateStructOnce sync.Once

func userContentFilterStorePrivateStructSet() {
	userContentFilterStorePrivateStructOnce.Do(func() {
		userContentFilterStorePrivateStruct = gi.StructNew("WebKit2", "UserContentFilterStorePrivate")
	})
}

type UserContentFilterStorePrivate struct {
	native uintptr
}

var userContentManagerClassStruct *gi.Struct
var userContentManagerClassStructOnce sync.Once

func userContentManagerClassStructSet() {
	userContentManagerClassStructOnce.Do(func() {
		userContentManagerClassStruct = gi.StructNew("WebKit2", "UserContentManagerClass")
	})
}

type UserContentManagerClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value '_webkit_reserved0' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved1' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved2' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved3' : missing Type
}

var userContentManagerPrivateStruct *gi.Struct
var userContentManagerPrivateStructOnce sync.Once

func userContentManagerPrivateStructSet() {
	userContentManagerPrivateStructOnce.Do(func() {
		userContentManagerPrivateStruct = gi.StructNew("WebKit2", "UserContentManagerPrivate")
	})
}

type UserContentManagerPrivate struct {
	native uintptr
}

var userMediaPermissionRequestClassStruct *gi.Struct
var userMediaPermissionRequestClassStructOnce sync.Once

func userMediaPermissionRequestClassStructSet() {
	userMediaPermissionRequestClassStructOnce.Do(func() {
		userMediaPermissionRequestClassStruct = gi.StructNew("WebKit2", "UserMediaPermissionRequestClass")
	})
}

type UserMediaPermissionRequestClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value '_webkit_reserved0' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved1' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved2' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved3' : missing Type
}

var userMediaPermissionRequestPrivateStruct *gi.Struct
var userMediaPermissionRequestPrivateStructOnce sync.Once

func userMediaPermissionRequestPrivateStructSet() {
	userMediaPermissionRequestPrivateStructOnce.Do(func() {
		userMediaPermissionRequestPrivateStruct = gi.StructNew("WebKit2", "UserMediaPermissionRequestPrivate")
	})
}

type UserMediaPermissionRequestPrivate struct {
	native uintptr
}

var userScriptStruct *gi.Struct
var userScriptStructOnce sync.Once

func userScriptStructSet() {
	userScriptStructOnce.Do(func() {
		userScriptStruct = gi.StructNew("WebKit2", "UserScript")
	})
}

type UserScript struct {
	native uintptr
}

// UNSUPPORTED : C value 'webkit_user_script_new' : parameter 'injected_frames' of type 'UserContentInjectedFrames' not supported

// UNSUPPORTED : C value 'webkit_user_script_new_for_world' : parameter 'injected_frames' of type 'UserContentInjectedFrames' not supported

var userScriptRefFunction *gi.Function
var userScriptRefFunctionOnce sync.Once

func userScriptRefFunctionSet() {
	userScriptRefFunctionOnce.Do(func() {
		userScriptRefFunction = gi.FunctionInvokerNew("WebKit2", "ref")
	})
}

var refUserScriptInvoker *gi.Function

// Ref is a representation of the C type webkit_user_script_ref.
func (recv *UserScript) Ref() *UserScript {
	if refUserScriptInvoker == nil {
		refUserScriptInvoker = gi.StructFunctionInvokerNew("WebKit2", "UserScript", "ref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := refUserScriptInvoker.Invoke(inArgs[:], nil)

	retGo := &UserScript{native: ret.Pointer()}

	return retGo
}

var userScriptUnrefFunction *gi.Function
var userScriptUnrefFunctionOnce sync.Once

func userScriptUnrefFunctionSet() {
	userScriptUnrefFunctionOnce.Do(func() {
		userScriptUnrefFunction = gi.FunctionInvokerNew("WebKit2", "unref")
	})
}

var unrefUserScriptInvoker *gi.Function

// Unref is a representation of the C type webkit_user_script_unref.
func (recv *UserScript) Unref() {
	if unrefUserScriptInvoker == nil {
		unrefUserScriptInvoker = gi.StructFunctionInvokerNew("WebKit2", "UserScript", "unref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	unrefUserScriptInvoker.Invoke(inArgs[:], nil)

}

var userStyleSheetStruct *gi.Struct
var userStyleSheetStructOnce sync.Once

func userStyleSheetStructSet() {
	userStyleSheetStructOnce.Do(func() {
		userStyleSheetStruct = gi.StructNew("WebKit2", "UserStyleSheet")
	})
}

type UserStyleSheet struct {
	native uintptr
}

// UNSUPPORTED : C value 'webkit_user_style_sheet_new' : parameter 'injected_frames' of type 'UserContentInjectedFrames' not supported

// UNSUPPORTED : C value 'webkit_user_style_sheet_new_for_world' : parameter 'injected_frames' of type 'UserContentInjectedFrames' not supported

var userStyleSheetRefFunction *gi.Function
var userStyleSheetRefFunctionOnce sync.Once

func userStyleSheetRefFunctionSet() {
	userStyleSheetRefFunctionOnce.Do(func() {
		userStyleSheetRefFunction = gi.FunctionInvokerNew("WebKit2", "ref")
	})
}

var refUserStyleSheetInvoker *gi.Function

// Ref is a representation of the C type webkit_user_style_sheet_ref.
func (recv *UserStyleSheet) Ref() *UserStyleSheet {
	if refUserStyleSheetInvoker == nil {
		refUserStyleSheetInvoker = gi.StructFunctionInvokerNew("WebKit2", "UserStyleSheet", "ref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := refUserStyleSheetInvoker.Invoke(inArgs[:], nil)

	retGo := &UserStyleSheet{native: ret.Pointer()}

	return retGo
}

var userStyleSheetUnrefFunction *gi.Function
var userStyleSheetUnrefFunctionOnce sync.Once

func userStyleSheetUnrefFunctionSet() {
	userStyleSheetUnrefFunctionOnce.Do(func() {
		userStyleSheetUnrefFunction = gi.FunctionInvokerNew("WebKit2", "unref")
	})
}

var unrefUserStyleSheetInvoker *gi.Function

// Unref is a representation of the C type webkit_user_style_sheet_unref.
func (recv *UserStyleSheet) Unref() {
	if unrefUserStyleSheetInvoker == nil {
		unrefUserStyleSheetInvoker = gi.StructFunctionInvokerNew("WebKit2", "UserStyleSheet", "unref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	unrefUserStyleSheetInvoker.Invoke(inArgs[:], nil)

}

var webContextClassStruct *gi.Struct
var webContextClassStructOnce sync.Once

func webContextClassStructSet() {
	webContextClassStructOnce.Do(func() {
		webContextClassStruct = gi.StructNew("WebKit2", "WebContextClass")
	})
}

type WebContextClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value 'download_started' : missing Type
	// UNSUPPORTED : C value 'initialize_web_extensions' : missing Type
	// UNSUPPORTED : C value 'initialize_notification_permissions' : missing Type
	// UNSUPPORTED : C value 'automation_started' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved0' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved1' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved2' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved3' : missing Type
}

var webContextPrivateStruct *gi.Struct
var webContextPrivateStructOnce sync.Once

func webContextPrivateStructSet() {
	webContextPrivateStructOnce.Do(func() {
		webContextPrivateStruct = gi.StructNew("WebKit2", "WebContextPrivate")
	})
}

type WebContextPrivate struct {
	native uintptr
}

var webInspectorClassStruct *gi.Struct
var webInspectorClassStructOnce sync.Once

func webInspectorClassStructSet() {
	webInspectorClassStructOnce.Do(func() {
		webInspectorClassStruct = gi.StructNew("WebKit2", "WebInspectorClass")
	})
}

type WebInspectorClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value '_webkit_reserved0' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved1' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved2' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved3' : missing Type
}

var webInspectorPrivateStruct *gi.Struct
var webInspectorPrivateStructOnce sync.Once

func webInspectorPrivateStructSet() {
	webInspectorPrivateStructOnce.Do(func() {
		webInspectorPrivateStruct = gi.StructNew("WebKit2", "WebInspectorPrivate")
	})
}

type WebInspectorPrivate struct {
	native uintptr
}

var webResourceClassStruct *gi.Struct
var webResourceClassStructOnce sync.Once

func webResourceClassStructSet() {
	webResourceClassStructOnce.Do(func() {
		webResourceClassStruct = gi.StructNew("WebKit2", "WebResourceClass")
	})
}

type WebResourceClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value '_webkit_reserved0' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved1' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved2' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved3' : missing Type
}

var webResourcePrivateStruct *gi.Struct
var webResourcePrivateStructOnce sync.Once

func webResourcePrivateStructSet() {
	webResourcePrivateStructOnce.Do(func() {
		webResourcePrivateStruct = gi.StructNew("WebKit2", "WebResourcePrivate")
	})
}

type WebResourcePrivate struct {
	native uintptr
}

var webViewBaseClassStruct *gi.Struct
var webViewBaseClassStructOnce sync.Once

func webViewBaseClassStructSet() {
	webViewBaseClassStructOnce.Do(func() {
		webViewBaseClassStruct = gi.StructNew("WebKit2", "WebViewBaseClass")
	})
}

type WebViewBaseClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parentClass' : no Go type for 'Gtk.ContainerClass'
	// UNSUPPORTED : C value '_webkit_reserved0' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved1' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved2' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved3' : missing Type
}

var webViewBasePrivateStruct *gi.Struct
var webViewBasePrivateStructOnce sync.Once

func webViewBasePrivateStructSet() {
	webViewBasePrivateStructOnce.Do(func() {
		webViewBasePrivateStruct = gi.StructNew("WebKit2", "WebViewBasePrivate")
	})
}

type WebViewBasePrivate struct {
	native uintptr
}

var webViewClassStruct *gi.Struct
var webViewClassStructOnce sync.Once

func webViewClassStructSet() {
	webViewClassStructOnce.Do(func() {
		webViewClassStruct = gi.StructNew("WebKit2", "WebViewClass")
	})
}

type WebViewClass struct {
	native uintptr
	Parent *WebViewBaseClass
	// UNSUPPORTED : C value 'load_changed' : missing Type
	// UNSUPPORTED : C value 'load_failed' : missing Type
	// UNSUPPORTED : C value 'create' : missing Type
	// UNSUPPORTED : C value 'ready_to_show' : missing Type
	// UNSUPPORTED : C value 'run_as_modal' : missing Type
	// UNSUPPORTED : C value 'close' : missing Type
	// UNSUPPORTED : C value 'script_dialog' : missing Type
	// UNSUPPORTED : C value 'decide_policy' : missing Type
	// UNSUPPORTED : C value 'permission_request' : missing Type
	// UNSUPPORTED : C value 'mouse_target_changed' : missing Type
	// UNSUPPORTED : C value 'print' : missing Type
	// UNSUPPORTED : C value 'resource_load_started' : missing Type
	// UNSUPPORTED : C value 'enter_fullscreen' : missing Type
	// UNSUPPORTED : C value 'leave_fullscreen' : missing Type
	// UNSUPPORTED : C value 'run_file_chooser' : missing Type
	// UNSUPPORTED : C value 'context_menu' : missing Type
	// UNSUPPORTED : C value 'context_menu_dismissed' : missing Type
	// UNSUPPORTED : C value 'submit_form' : missing Type
	// UNSUPPORTED : C value 'insecure_content_detected' : missing Type
	// UNSUPPORTED : C value 'web_process_crashed' : missing Type
	// UNSUPPORTED : C value 'authenticate' : missing Type
	// UNSUPPORTED : C value 'load_failed_with_tls_errors' : missing Type
	// UNSUPPORTED : C value 'show_notification' : missing Type
	// UNSUPPORTED : C value 'run_color_chooser' : missing Type
	// UNSUPPORTED : C value 'show_option_menu' : missing Type
	// UNSUPPORTED : C value 'web_process_terminated' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved0' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved1' : missing Type
}

var webViewPrivateStruct *gi.Struct
var webViewPrivateStructOnce sync.Once

func webViewPrivateStructSet() {
	webViewPrivateStructOnce.Do(func() {
		webViewPrivateStruct = gi.StructNew("WebKit2", "WebViewPrivate")
	})
}

type WebViewPrivate struct {
	native uintptr
}

var webViewSessionStateStruct *gi.Struct
var webViewSessionStateStructOnce sync.Once

func webViewSessionStateStructSet() {
	webViewSessionStateStructOnce.Do(func() {
		webViewSessionStateStruct = gi.StructNew("WebKit2", "WebViewSessionState")
	})
}

type WebViewSessionState struct {
	native uintptr
}

// UNSUPPORTED : C value 'webkit_web_view_session_state_new' : parameter 'data' of type 'GLib.Bytes' not supported

var webViewSessionStateRefFunction *gi.Function
var webViewSessionStateRefFunctionOnce sync.Once

func webViewSessionStateRefFunctionSet() {
	webViewSessionStateRefFunctionOnce.Do(func() {
		webViewSessionStateRefFunction = gi.FunctionInvokerNew("WebKit2", "ref")
	})
}

var refWebViewSessionStateInvoker *gi.Function

// Ref is a representation of the C type webkit_web_view_session_state_ref.
func (recv *WebViewSessionState) Ref() *WebViewSessionState {
	if refWebViewSessionStateInvoker == nil {
		refWebViewSessionStateInvoker = gi.StructFunctionInvokerNew("WebKit2", "WebViewSessionState", "ref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := refWebViewSessionStateInvoker.Invoke(inArgs[:], nil)

	retGo := &WebViewSessionState{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'webkit_web_view_session_state_serialize' : return type 'GLib.Bytes' not supported

var webViewSessionStateUnrefFunction *gi.Function
var webViewSessionStateUnrefFunctionOnce sync.Once

func webViewSessionStateUnrefFunctionSet() {
	webViewSessionStateUnrefFunctionOnce.Do(func() {
		webViewSessionStateUnrefFunction = gi.FunctionInvokerNew("WebKit2", "unref")
	})
}

var unrefWebViewSessionStateInvoker *gi.Function

// Unref is a representation of the C type webkit_web_view_session_state_unref.
func (recv *WebViewSessionState) Unref() {
	if unrefWebViewSessionStateInvoker == nil {
		unrefWebViewSessionStateInvoker = gi.StructFunctionInvokerNew("WebKit2", "WebViewSessionState", "unref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	unrefWebViewSessionStateInvoker.Invoke(inArgs[:], nil)

}

var websiteDataStruct *gi.Struct
var websiteDataStructOnce sync.Once

func websiteDataStructSet() {
	websiteDataStructOnce.Do(func() {
		websiteDataStruct = gi.StructNew("WebKit2", "WebsiteData")
	})
}

type WebsiteData struct {
	native uintptr
}

var websiteDataGetNameFunction *gi.Function
var websiteDataGetNameFunctionOnce sync.Once

func websiteDataGetNameFunctionSet() {
	websiteDataGetNameFunctionOnce.Do(func() {
		websiteDataGetNameFunction = gi.FunctionInvokerNew("WebKit2", "get_name")
	})
}

var getNameWebsiteDataInvoker *gi.Function

// GetName is a representation of the C type webkit_website_data_get_name.
func (recv *WebsiteData) GetName() string {
	if getNameWebsiteDataInvoker == nil {
		getNameWebsiteDataInvoker = gi.StructFunctionInvokerNew("WebKit2", "WebsiteData", "get_name")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getNameWebsiteDataInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'webkit_website_data_get_size' : parameter 'types' of type 'WebsiteDataTypes' not supported

// UNSUPPORTED : C value 'webkit_website_data_get_types' : return type 'WebsiteDataTypes' not supported

var websiteDataRefFunction *gi.Function
var websiteDataRefFunctionOnce sync.Once

func websiteDataRefFunctionSet() {
	websiteDataRefFunctionOnce.Do(func() {
		websiteDataRefFunction = gi.FunctionInvokerNew("WebKit2", "ref")
	})
}

var refWebsiteDataInvoker *gi.Function

// Ref is a representation of the C type webkit_website_data_ref.
func (recv *WebsiteData) Ref() *WebsiteData {
	if refWebsiteDataInvoker == nil {
		refWebsiteDataInvoker = gi.StructFunctionInvokerNew("WebKit2", "WebsiteData", "ref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := refWebsiteDataInvoker.Invoke(inArgs[:], nil)

	retGo := &WebsiteData{native: ret.Pointer()}

	return retGo
}

var websiteDataUnrefFunction *gi.Function
var websiteDataUnrefFunctionOnce sync.Once

func websiteDataUnrefFunctionSet() {
	websiteDataUnrefFunctionOnce.Do(func() {
		websiteDataUnrefFunction = gi.FunctionInvokerNew("WebKit2", "unref")
	})
}

var unrefWebsiteDataInvoker *gi.Function

// Unref is a representation of the C type webkit_website_data_unref.
func (recv *WebsiteData) Unref() {
	if unrefWebsiteDataInvoker == nil {
		unrefWebsiteDataInvoker = gi.StructFunctionInvokerNew("WebKit2", "WebsiteData", "unref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	unrefWebsiteDataInvoker.Invoke(inArgs[:], nil)

}

var websiteDataManagerClassStruct *gi.Struct
var websiteDataManagerClassStructOnce sync.Once

func websiteDataManagerClassStructSet() {
	websiteDataManagerClassStructOnce.Do(func() {
		websiteDataManagerClassStruct = gi.StructNew("WebKit2", "WebsiteDataManagerClass")
	})
}

type WebsiteDataManagerClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value '_webkit_reserved0' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved1' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved2' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved3' : missing Type
}

var websiteDataManagerPrivateStruct *gi.Struct
var websiteDataManagerPrivateStructOnce sync.Once

func websiteDataManagerPrivateStructSet() {
	websiteDataManagerPrivateStructOnce.Do(func() {
		websiteDataManagerPrivateStruct = gi.StructNew("WebKit2", "WebsiteDataManagerPrivate")
	})
}

type WebsiteDataManagerPrivate struct {
	native uintptr
}

var windowPropertiesClassStruct *gi.Struct
var windowPropertiesClassStructOnce sync.Once

func windowPropertiesClassStructSet() {
	windowPropertiesClassStructOnce.Do(func() {
		windowPropertiesClassStruct = gi.StructNew("WebKit2", "WindowPropertiesClass")
	})
}

type WindowPropertiesClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value '_webkit_reserved0' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved1' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved2' : missing Type
	// UNSUPPORTED : C value '_webkit_reserved3' : missing Type
}

var windowPropertiesPrivateStruct *gi.Struct
var windowPropertiesPrivateStructOnce sync.Once

func windowPropertiesPrivateStructSet() {
	windowPropertiesPrivateStructOnce.Do(func() {
		windowPropertiesPrivateStruct = gi.StructNew("WebKit2", "WindowPropertiesPrivate")
	})
}

type WindowPropertiesPrivate struct {
	native uintptr
}
