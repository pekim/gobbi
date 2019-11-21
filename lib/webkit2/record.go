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
var applicationInfoNewFunction_Once sync.Once

func applicationInfoNewFunction_Set() {
	applicationInfoNewFunction_Once.Do(func() {
		applicationInfoNewFunction = gi.FunctionInvokerNew("WebKit2", "new")
	})
}

var newApplicationInfoInvoker *gi.Function

// ApplicationInfoNew is a representation of the C type webkit_application_info_new.
func ApplicationInfoNew() *ApplicationInfo {
	applicationInfoNewFunction_Set()

	ret := applicationInfoNewFunction.Invoke(nil, nil)

	retGo := &ApplicationInfo{native: ret.Pointer()}

	return retGo
}

var applicationInfoGetNameFunction *gi.Function
var applicationInfoGetNameFunction_Once sync.Once

func applicationInfoGetNameFunction_Set() {
	applicationInfoGetNameFunction_Once.Do(func() {
		applicationInfoGetNameFunction = gi.FunctionInvokerNew("WebKit2", "get_name")
	})
}

var getNameApplicationInfoInvoker *gi.Function

// GetName is a representation of the C type webkit_application_info_get_name.
func (recv *ApplicationInfo) GetName() string {
	applicationInfoGetNameFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := applicationInfoGetNameFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

var applicationInfoGetVersionFunction *gi.Function
var applicationInfoGetVersionFunction_Once sync.Once

func applicationInfoGetVersionFunction_Set() {
	applicationInfoGetVersionFunction_Once.Do(func() {
		applicationInfoGetVersionFunction = gi.FunctionInvokerNew("WebKit2", "get_version")
	})
}

var getVersionApplicationInfoInvoker *gi.Function

// GetVersion is a representation of the C type webkit_application_info_get_version.
func (recv *ApplicationInfo) GetVersion() (uint64, uint64, uint64) {
	applicationInfoGetVersionFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var outArgs [3]gi.Argument

	applicationInfoGetVersionFunction.Invoke(inArgs[:], outArgs[:])

	out0 := outArgs[0].Uint64()
	out1 := outArgs[1].Uint64()
	out2 := outArgs[2].Uint64()

	return out0, out1, out2
}

var applicationInfoRefFunction *gi.Function
var applicationInfoRefFunction_Once sync.Once

func applicationInfoRefFunction_Set() {
	applicationInfoRefFunction_Once.Do(func() {
		applicationInfoRefFunction = gi.FunctionInvokerNew("WebKit2", "ref")
	})
}

var refApplicationInfoInvoker *gi.Function

// Ref is a representation of the C type webkit_application_info_ref.
func (recv *ApplicationInfo) Ref() *ApplicationInfo {
	applicationInfoRefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := applicationInfoRefFunction.Invoke(inArgs[:], nil)

	retGo := &ApplicationInfo{native: ret.Pointer()}

	return retGo
}

var applicationInfoSetNameFunction *gi.Function
var applicationInfoSetNameFunction_Once sync.Once

func applicationInfoSetNameFunction_Set() {
	applicationInfoSetNameFunction_Once.Do(func() {
		applicationInfoSetNameFunction = gi.FunctionInvokerNew("WebKit2", "set_name")
	})
}

var setNameApplicationInfoInvoker *gi.Function

// SetName is a representation of the C type webkit_application_info_set_name.
func (recv *ApplicationInfo) SetName(name string) {
	applicationInfoSetNameFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(name)

	applicationInfoSetNameFunction.Invoke(inArgs[:], nil)

}

var applicationInfoSetVersionFunction *gi.Function
var applicationInfoSetVersionFunction_Once sync.Once

func applicationInfoSetVersionFunction_Set() {
	applicationInfoSetVersionFunction_Once.Do(func() {
		applicationInfoSetVersionFunction = gi.FunctionInvokerNew("WebKit2", "set_version")
	})
}

var setVersionApplicationInfoInvoker *gi.Function

// SetVersion is a representation of the C type webkit_application_info_set_version.
func (recv *ApplicationInfo) SetVersion(major uint64, minor uint64, micro uint64) {
	applicationInfoSetVersionFunction_Set()

	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetUint64(major)
	inArgs[2].SetUint64(minor)
	inArgs[3].SetUint64(micro)

	applicationInfoSetVersionFunction.Invoke(inArgs[:], nil)

}

var applicationInfoUnrefFunction *gi.Function
var applicationInfoUnrefFunction_Once sync.Once

func applicationInfoUnrefFunction_Set() {
	applicationInfoUnrefFunction_Once.Do(func() {
		applicationInfoUnrefFunction = gi.FunctionInvokerNew("WebKit2", "unref")
	})
}

var unrefApplicationInfoInvoker *gi.Function

// Unref is a representation of the C type webkit_application_info_unref.
func (recv *ApplicationInfo) Unref() {
	applicationInfoUnrefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	applicationInfoUnrefFunction.Invoke(inArgs[:], nil)

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
var credentialCopyFunction_Once sync.Once

func credentialCopyFunction_Set() {
	credentialCopyFunction_Once.Do(func() {
		credentialCopyFunction = gi.FunctionInvokerNew("WebKit2", "copy")
	})
}

var copyCredentialInvoker *gi.Function

// Copy is a representation of the C type webkit_credential_copy.
func (recv *Credential) Copy() *Credential {
	credentialCopyFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := credentialCopyFunction.Invoke(inArgs[:], nil)

	retGo := &Credential{native: ret.Pointer()}

	return retGo
}

var credentialFreeFunction *gi.Function
var credentialFreeFunction_Once sync.Once

func credentialFreeFunction_Set() {
	credentialFreeFunction_Once.Do(func() {
		credentialFreeFunction = gi.FunctionInvokerNew("WebKit2", "free")
	})
}

var freeCredentialInvoker *gi.Function

// Free is a representation of the C type webkit_credential_free.
func (recv *Credential) Free() {
	credentialFreeFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	credentialFreeFunction.Invoke(inArgs[:], nil)

}

var credentialGetPasswordFunction *gi.Function
var credentialGetPasswordFunction_Once sync.Once

func credentialGetPasswordFunction_Set() {
	credentialGetPasswordFunction_Once.Do(func() {
		credentialGetPasswordFunction = gi.FunctionInvokerNew("WebKit2", "get_password")
	})
}

var getPasswordCredentialInvoker *gi.Function

// GetPassword is a representation of the C type webkit_credential_get_password.
func (recv *Credential) GetPassword() string {
	credentialGetPasswordFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := credentialGetPasswordFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'webkit_credential_get_persistence' : return type 'CredentialPersistence' not supported

var credentialGetUsernameFunction *gi.Function
var credentialGetUsernameFunction_Once sync.Once

func credentialGetUsernameFunction_Set() {
	credentialGetUsernameFunction_Once.Do(func() {
		credentialGetUsernameFunction = gi.FunctionInvokerNew("WebKit2", "get_username")
	})
}

var getUsernameCredentialInvoker *gi.Function

// GetUsername is a representation of the C type webkit_credential_get_username.
func (recv *Credential) GetUsername() string {
	credentialGetUsernameFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := credentialGetUsernameFunction.Invoke(inArgs[:], nil)

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
var geolocationPositionCopyFunction_Once sync.Once

func geolocationPositionCopyFunction_Set() {
	geolocationPositionCopyFunction_Once.Do(func() {
		geolocationPositionCopyFunction = gi.FunctionInvokerNew("WebKit2", "copy")
	})
}

var copyGeolocationPositionInvoker *gi.Function

// Copy is a representation of the C type webkit_geolocation_position_copy.
func (recv *GeolocationPosition) Copy() *GeolocationPosition {
	geolocationPositionCopyFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := geolocationPositionCopyFunction.Invoke(inArgs[:], nil)

	retGo := &GeolocationPosition{native: ret.Pointer()}

	return retGo
}

var geolocationPositionFreeFunction *gi.Function
var geolocationPositionFreeFunction_Once sync.Once

func geolocationPositionFreeFunction_Set() {
	geolocationPositionFreeFunction_Once.Do(func() {
		geolocationPositionFreeFunction = gi.FunctionInvokerNew("WebKit2", "free")
	})
}

var freeGeolocationPositionInvoker *gi.Function

// Free is a representation of the C type webkit_geolocation_position_free.
func (recv *GeolocationPosition) Free() {
	geolocationPositionFreeFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	geolocationPositionFreeFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'webkit_geolocation_position_set_altitude' : parameter 'altitude' of type 'gdouble' not supported

// UNSUPPORTED : C value 'webkit_geolocation_position_set_altitude_accuracy' : parameter 'altitude_accuracy' of type 'gdouble' not supported

// UNSUPPORTED : C value 'webkit_geolocation_position_set_heading' : parameter 'heading' of type 'gdouble' not supported

// UNSUPPORTED : C value 'webkit_geolocation_position_set_speed' : parameter 'speed' of type 'gdouble' not supported

var geolocationPositionSetTimestampFunction *gi.Function
var geolocationPositionSetTimestampFunction_Once sync.Once

func geolocationPositionSetTimestampFunction_Set() {
	geolocationPositionSetTimestampFunction_Once.Do(func() {
		geolocationPositionSetTimestampFunction = gi.FunctionInvokerNew("WebKit2", "set_timestamp")
	})
}

var setTimestampGeolocationPositionInvoker *gi.Function

// SetTimestamp is a representation of the C type webkit_geolocation_position_set_timestamp.
func (recv *GeolocationPosition) SetTimestamp(timestamp uint64) {
	geolocationPositionSetTimestampFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetUint64(timestamp)

	geolocationPositionSetTimestampFunction.Invoke(inArgs[:], nil)

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
var javascriptResultRefFunction_Once sync.Once

func javascriptResultRefFunction_Set() {
	javascriptResultRefFunction_Once.Do(func() {
		javascriptResultRefFunction = gi.FunctionInvokerNew("WebKit2", "ref")
	})
}

var refJavascriptResultInvoker *gi.Function

// Ref is a representation of the C type webkit_javascript_result_ref.
func (recv *JavascriptResult) Ref() *JavascriptResult {
	javascriptResultRefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := javascriptResultRefFunction.Invoke(inArgs[:], nil)

	retGo := &JavascriptResult{native: ret.Pointer()}

	return retGo
}

var javascriptResultUnrefFunction *gi.Function
var javascriptResultUnrefFunction_Once sync.Once

func javascriptResultUnrefFunction_Set() {
	javascriptResultUnrefFunction_Once.Do(func() {
		javascriptResultUnrefFunction = gi.FunctionInvokerNew("WebKit2", "unref")
	})
}

var unrefJavascriptResultInvoker *gi.Function

// Unref is a representation of the C type webkit_javascript_result_unref.
func (recv *JavascriptResult) Unref() {
	javascriptResultUnrefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	javascriptResultUnrefFunction.Invoke(inArgs[:], nil)

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
var mimeInfoGetDescriptionFunction_Once sync.Once

func mimeInfoGetDescriptionFunction_Set() {
	mimeInfoGetDescriptionFunction_Once.Do(func() {
		mimeInfoGetDescriptionFunction = gi.FunctionInvokerNew("WebKit2", "get_description")
	})
}

var getDescriptionMimeInfoInvoker *gi.Function

// GetDescription is a representation of the C type webkit_mime_info_get_description.
func (recv *MimeInfo) GetDescription() string {
	mimeInfoGetDescriptionFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := mimeInfoGetDescriptionFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

var mimeInfoGetExtensionsFunction *gi.Function
var mimeInfoGetExtensionsFunction_Once sync.Once

func mimeInfoGetExtensionsFunction_Set() {
	mimeInfoGetExtensionsFunction_Once.Do(func() {
		mimeInfoGetExtensionsFunction = gi.FunctionInvokerNew("WebKit2", "get_extensions")
	})
}

var getExtensionsMimeInfoInvoker *gi.Function

// GetExtensions is a representation of the C type webkit_mime_info_get_extensions.
func (recv *MimeInfo) GetExtensions() {
	mimeInfoGetExtensionsFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	mimeInfoGetExtensionsFunction.Invoke(inArgs[:], nil)

}

var mimeInfoGetMimeTypeFunction *gi.Function
var mimeInfoGetMimeTypeFunction_Once sync.Once

func mimeInfoGetMimeTypeFunction_Set() {
	mimeInfoGetMimeTypeFunction_Once.Do(func() {
		mimeInfoGetMimeTypeFunction = gi.FunctionInvokerNew("WebKit2", "get_mime_type")
	})
}

var getMimeTypeMimeInfoInvoker *gi.Function

// GetMimeType is a representation of the C type webkit_mime_info_get_mime_type.
func (recv *MimeInfo) GetMimeType() string {
	mimeInfoGetMimeTypeFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := mimeInfoGetMimeTypeFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

var mimeInfoRefFunction *gi.Function
var mimeInfoRefFunction_Once sync.Once

func mimeInfoRefFunction_Set() {
	mimeInfoRefFunction_Once.Do(func() {
		mimeInfoRefFunction = gi.FunctionInvokerNew("WebKit2", "ref")
	})
}

var refMimeInfoInvoker *gi.Function

// Ref is a representation of the C type webkit_mime_info_ref.
func (recv *MimeInfo) Ref() *MimeInfo {
	mimeInfoRefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := mimeInfoRefFunction.Invoke(inArgs[:], nil)

	retGo := &MimeInfo{native: ret.Pointer()}

	return retGo
}

var mimeInfoUnrefFunction *gi.Function
var mimeInfoUnrefFunction_Once sync.Once

func mimeInfoUnrefFunction_Set() {
	mimeInfoUnrefFunction_Once.Do(func() {
		mimeInfoUnrefFunction = gi.FunctionInvokerNew("WebKit2", "unref")
	})
}

var unrefMimeInfoInvoker *gi.Function

// Unref is a representation of the C type webkit_mime_info_unref.
func (recv *MimeInfo) Unref() {
	mimeInfoUnrefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	mimeInfoUnrefFunction.Invoke(inArgs[:], nil)

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
var navigationActionCopyFunction_Once sync.Once

func navigationActionCopyFunction_Set() {
	navigationActionCopyFunction_Once.Do(func() {
		navigationActionCopyFunction = gi.FunctionInvokerNew("WebKit2", "copy")
	})
}

var copyNavigationActionInvoker *gi.Function

// Copy is a representation of the C type webkit_navigation_action_copy.
func (recv *NavigationAction) Copy() *NavigationAction {
	navigationActionCopyFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := navigationActionCopyFunction.Invoke(inArgs[:], nil)

	retGo := &NavigationAction{native: ret.Pointer()}

	return retGo
}

var navigationActionFreeFunction *gi.Function
var navigationActionFreeFunction_Once sync.Once

func navigationActionFreeFunction_Set() {
	navigationActionFreeFunction_Once.Do(func() {
		navigationActionFreeFunction = gi.FunctionInvokerNew("WebKit2", "free")
	})
}

var freeNavigationActionInvoker *gi.Function

// Free is a representation of the C type webkit_navigation_action_free.
func (recv *NavigationAction) Free() {
	navigationActionFreeFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	navigationActionFreeFunction.Invoke(inArgs[:], nil)

}

var navigationActionGetModifiersFunction *gi.Function
var navigationActionGetModifiersFunction_Once sync.Once

func navigationActionGetModifiersFunction_Set() {
	navigationActionGetModifiersFunction_Once.Do(func() {
		navigationActionGetModifiersFunction = gi.FunctionInvokerNew("WebKit2", "get_modifiers")
	})
}

var getModifiersNavigationActionInvoker *gi.Function

// GetModifiers is a representation of the C type webkit_navigation_action_get_modifiers.
func (recv *NavigationAction) GetModifiers() uint32 {
	navigationActionGetModifiersFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := navigationActionGetModifiersFunction.Invoke(inArgs[:], nil)

	retGo := ret.Uint32()

	return retGo
}

var navigationActionGetMouseButtonFunction *gi.Function
var navigationActionGetMouseButtonFunction_Once sync.Once

func navigationActionGetMouseButtonFunction_Set() {
	navigationActionGetMouseButtonFunction_Once.Do(func() {
		navigationActionGetMouseButtonFunction = gi.FunctionInvokerNew("WebKit2", "get_mouse_button")
	})
}

var getMouseButtonNavigationActionInvoker *gi.Function

// GetMouseButton is a representation of the C type webkit_navigation_action_get_mouse_button.
func (recv *NavigationAction) GetMouseButton() uint32 {
	navigationActionGetMouseButtonFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := navigationActionGetMouseButtonFunction.Invoke(inArgs[:], nil)

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
var networkProxySettingsAddProxyForSchemeFunction_Once sync.Once

func networkProxySettingsAddProxyForSchemeFunction_Set() {
	networkProxySettingsAddProxyForSchemeFunction_Once.Do(func() {
		networkProxySettingsAddProxyForSchemeFunction = gi.FunctionInvokerNew("WebKit2", "add_proxy_for_scheme")
	})
}

var addProxyForSchemeNetworkProxySettingsInvoker *gi.Function

// AddProxyForScheme is a representation of the C type webkit_network_proxy_settings_add_proxy_for_scheme.
func (recv *NetworkProxySettings) AddProxyForScheme(scheme string, proxyUri string) {
	networkProxySettingsAddProxyForSchemeFunction_Set()

	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(scheme)
	inArgs[2].SetString(proxyUri)

	networkProxySettingsAddProxyForSchemeFunction.Invoke(inArgs[:], nil)

}

var networkProxySettingsCopyFunction *gi.Function
var networkProxySettingsCopyFunction_Once sync.Once

func networkProxySettingsCopyFunction_Set() {
	networkProxySettingsCopyFunction_Once.Do(func() {
		networkProxySettingsCopyFunction = gi.FunctionInvokerNew("WebKit2", "copy")
	})
}

var copyNetworkProxySettingsInvoker *gi.Function

// Copy is a representation of the C type webkit_network_proxy_settings_copy.
func (recv *NetworkProxySettings) Copy() *NetworkProxySettings {
	networkProxySettingsCopyFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := networkProxySettingsCopyFunction.Invoke(inArgs[:], nil)

	retGo := &NetworkProxySettings{native: ret.Pointer()}

	return retGo
}

var networkProxySettingsFreeFunction *gi.Function
var networkProxySettingsFreeFunction_Once sync.Once

func networkProxySettingsFreeFunction_Set() {
	networkProxySettingsFreeFunction_Once.Do(func() {
		networkProxySettingsFreeFunction = gi.FunctionInvokerNew("WebKit2", "free")
	})
}

var freeNetworkProxySettingsInvoker *gi.Function

// Free is a representation of the C type webkit_network_proxy_settings_free.
func (recv *NetworkProxySettings) Free() {
	networkProxySettingsFreeFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	networkProxySettingsFreeFunction.Invoke(inArgs[:], nil)

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
var optionMenuItemCopyFunction_Once sync.Once

func optionMenuItemCopyFunction_Set() {
	optionMenuItemCopyFunction_Once.Do(func() {
		optionMenuItemCopyFunction = gi.FunctionInvokerNew("WebKit2", "copy")
	})
}

var copyOptionMenuItemInvoker *gi.Function

// Copy is a representation of the C type webkit_option_menu_item_copy.
func (recv *OptionMenuItem) Copy() *OptionMenuItem {
	optionMenuItemCopyFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := optionMenuItemCopyFunction.Invoke(inArgs[:], nil)

	retGo := &OptionMenuItem{native: ret.Pointer()}

	return retGo
}

var optionMenuItemFreeFunction *gi.Function
var optionMenuItemFreeFunction_Once sync.Once

func optionMenuItemFreeFunction_Set() {
	optionMenuItemFreeFunction_Once.Do(func() {
		optionMenuItemFreeFunction = gi.FunctionInvokerNew("WebKit2", "free")
	})
}

var freeOptionMenuItemInvoker *gi.Function

// Free is a representation of the C type webkit_option_menu_item_free.
func (recv *OptionMenuItem) Free() {
	optionMenuItemFreeFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	optionMenuItemFreeFunction.Invoke(inArgs[:], nil)

}

var optionMenuItemGetLabelFunction *gi.Function
var optionMenuItemGetLabelFunction_Once sync.Once

func optionMenuItemGetLabelFunction_Set() {
	optionMenuItemGetLabelFunction_Once.Do(func() {
		optionMenuItemGetLabelFunction = gi.FunctionInvokerNew("WebKit2", "get_label")
	})
}

var getLabelOptionMenuItemInvoker *gi.Function

// GetLabel is a representation of the C type webkit_option_menu_item_get_label.
func (recv *OptionMenuItem) GetLabel() string {
	optionMenuItemGetLabelFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := optionMenuItemGetLabelFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

var optionMenuItemGetTooltipFunction *gi.Function
var optionMenuItemGetTooltipFunction_Once sync.Once

func optionMenuItemGetTooltipFunction_Set() {
	optionMenuItemGetTooltipFunction_Once.Do(func() {
		optionMenuItemGetTooltipFunction = gi.FunctionInvokerNew("WebKit2", "get_tooltip")
	})
}

var getTooltipOptionMenuItemInvoker *gi.Function

// GetTooltip is a representation of the C type webkit_option_menu_item_get_tooltip.
func (recv *OptionMenuItem) GetTooltip() string {
	optionMenuItemGetTooltipFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := optionMenuItemGetTooltipFunction.Invoke(inArgs[:], nil)

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
var scriptDialogCloseFunction_Once sync.Once

func scriptDialogCloseFunction_Set() {
	scriptDialogCloseFunction_Once.Do(func() {
		scriptDialogCloseFunction = gi.FunctionInvokerNew("WebKit2", "close")
	})
}

var closeScriptDialogInvoker *gi.Function

// Close is a representation of the C type webkit_script_dialog_close.
func (recv *ScriptDialog) Close() {
	scriptDialogCloseFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	scriptDialogCloseFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'webkit_script_dialog_confirm_set_confirmed' : parameter 'confirmed' of type 'gboolean' not supported

// UNSUPPORTED : C value 'webkit_script_dialog_get_dialog_type' : return type 'ScriptDialogType' not supported

var scriptDialogGetMessageFunction *gi.Function
var scriptDialogGetMessageFunction_Once sync.Once

func scriptDialogGetMessageFunction_Set() {
	scriptDialogGetMessageFunction_Once.Do(func() {
		scriptDialogGetMessageFunction = gi.FunctionInvokerNew("WebKit2", "get_message")
	})
}

var getMessageScriptDialogInvoker *gi.Function

// GetMessage is a representation of the C type webkit_script_dialog_get_message.
func (recv *ScriptDialog) GetMessage() string {
	scriptDialogGetMessageFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := scriptDialogGetMessageFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

var scriptDialogPromptGetDefaultTextFunction *gi.Function
var scriptDialogPromptGetDefaultTextFunction_Once sync.Once

func scriptDialogPromptGetDefaultTextFunction_Set() {
	scriptDialogPromptGetDefaultTextFunction_Once.Do(func() {
		scriptDialogPromptGetDefaultTextFunction = gi.FunctionInvokerNew("WebKit2", "prompt_get_default_text")
	})
}

var promptGetDefaultTextScriptDialogInvoker *gi.Function

// PromptGetDefaultText is a representation of the C type webkit_script_dialog_prompt_get_default_text.
func (recv *ScriptDialog) PromptGetDefaultText() string {
	scriptDialogPromptGetDefaultTextFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := scriptDialogPromptGetDefaultTextFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

var scriptDialogPromptSetTextFunction *gi.Function
var scriptDialogPromptSetTextFunction_Once sync.Once

func scriptDialogPromptSetTextFunction_Set() {
	scriptDialogPromptSetTextFunction_Once.Do(func() {
		scriptDialogPromptSetTextFunction = gi.FunctionInvokerNew("WebKit2", "prompt_set_text")
	})
}

var promptSetTextScriptDialogInvoker *gi.Function

// PromptSetText is a representation of the C type webkit_script_dialog_prompt_set_text.
func (recv *ScriptDialog) PromptSetText(text string) {
	scriptDialogPromptSetTextFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(text)

	scriptDialogPromptSetTextFunction.Invoke(inArgs[:], nil)

}

var scriptDialogRefFunction *gi.Function
var scriptDialogRefFunction_Once sync.Once

func scriptDialogRefFunction_Set() {
	scriptDialogRefFunction_Once.Do(func() {
		scriptDialogRefFunction = gi.FunctionInvokerNew("WebKit2", "ref")
	})
}

var refScriptDialogInvoker *gi.Function

// Ref is a representation of the C type webkit_script_dialog_ref.
func (recv *ScriptDialog) Ref() *ScriptDialog {
	scriptDialogRefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := scriptDialogRefFunction.Invoke(inArgs[:], nil)

	retGo := &ScriptDialog{native: ret.Pointer()}

	return retGo
}

var scriptDialogUnrefFunction *gi.Function
var scriptDialogUnrefFunction_Once sync.Once

func scriptDialogUnrefFunction_Set() {
	scriptDialogUnrefFunction_Once.Do(func() {
		scriptDialogUnrefFunction = gi.FunctionInvokerNew("WebKit2", "unref")
	})
}

var unrefScriptDialogInvoker *gi.Function

// Unref is a representation of the C type webkit_script_dialog_unref.
func (recv *ScriptDialog) Unref() {
	scriptDialogUnrefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	scriptDialogUnrefFunction.Invoke(inArgs[:], nil)

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
var securityOriginNewFunction_Once sync.Once

func securityOriginNewFunction_Set() {
	securityOriginNewFunction_Once.Do(func() {
		securityOriginNewFunction = gi.FunctionInvokerNew("WebKit2", "new")
	})
}

var newSecurityOriginInvoker *gi.Function

// SecurityOriginNew is a representation of the C type webkit_security_origin_new.
func SecurityOriginNew(protocol string, host string, port uint16) *SecurityOrigin {
	securityOriginNewFunction_Set()

	var inArgs [3]gi.Argument
	inArgs[0].SetString(protocol)
	inArgs[1].SetString(host)
	inArgs[2].SetUint16(port)

	ret := securityOriginNewFunction.Invoke(inArgs[:], nil)

	retGo := &SecurityOrigin{native: ret.Pointer()}

	return retGo
}

var securityOriginNewForUriFunction *gi.Function
var securityOriginNewForUriFunction_Once sync.Once

func securityOriginNewForUriFunction_Set() {
	securityOriginNewForUriFunction_Once.Do(func() {
		securityOriginNewForUriFunction = gi.FunctionInvokerNew("WebKit2", "new_for_uri")
	})
}

var newForUriSecurityOriginInvoker *gi.Function

// SecurityOriginNewForUri is a representation of the C type webkit_security_origin_new_for_uri.
func SecurityOriginNewForUri(uri string) *SecurityOrigin {
	securityOriginNewForUriFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetString(uri)

	ret := securityOriginNewForUriFunction.Invoke(inArgs[:], nil)

	retGo := &SecurityOrigin{native: ret.Pointer()}

	return retGo
}

var securityOriginGetHostFunction *gi.Function
var securityOriginGetHostFunction_Once sync.Once

func securityOriginGetHostFunction_Set() {
	securityOriginGetHostFunction_Once.Do(func() {
		securityOriginGetHostFunction = gi.FunctionInvokerNew("WebKit2", "get_host")
	})
}

var getHostSecurityOriginInvoker *gi.Function

// GetHost is a representation of the C type webkit_security_origin_get_host.
func (recv *SecurityOrigin) GetHost() string {
	securityOriginGetHostFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := securityOriginGetHostFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

var securityOriginGetPortFunction *gi.Function
var securityOriginGetPortFunction_Once sync.Once

func securityOriginGetPortFunction_Set() {
	securityOriginGetPortFunction_Once.Do(func() {
		securityOriginGetPortFunction = gi.FunctionInvokerNew("WebKit2", "get_port")
	})
}

var getPortSecurityOriginInvoker *gi.Function

// GetPort is a representation of the C type webkit_security_origin_get_port.
func (recv *SecurityOrigin) GetPort() uint16 {
	securityOriginGetPortFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := securityOriginGetPortFunction.Invoke(inArgs[:], nil)

	retGo := ret.Uint16()

	return retGo
}

var securityOriginGetProtocolFunction *gi.Function
var securityOriginGetProtocolFunction_Once sync.Once

func securityOriginGetProtocolFunction_Set() {
	securityOriginGetProtocolFunction_Once.Do(func() {
		securityOriginGetProtocolFunction = gi.FunctionInvokerNew("WebKit2", "get_protocol")
	})
}

var getProtocolSecurityOriginInvoker *gi.Function

// GetProtocol is a representation of the C type webkit_security_origin_get_protocol.
func (recv *SecurityOrigin) GetProtocol() string {
	securityOriginGetProtocolFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := securityOriginGetProtocolFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'webkit_security_origin_is_opaque' : return type 'gboolean' not supported

var securityOriginRefFunction *gi.Function
var securityOriginRefFunction_Once sync.Once

func securityOriginRefFunction_Set() {
	securityOriginRefFunction_Once.Do(func() {
		securityOriginRefFunction = gi.FunctionInvokerNew("WebKit2", "ref")
	})
}

var refSecurityOriginInvoker *gi.Function

// Ref is a representation of the C type webkit_security_origin_ref.
func (recv *SecurityOrigin) Ref() *SecurityOrigin {
	securityOriginRefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := securityOriginRefFunction.Invoke(inArgs[:], nil)

	retGo := &SecurityOrigin{native: ret.Pointer()}

	return retGo
}

var securityOriginToStringFunction *gi.Function
var securityOriginToStringFunction_Once sync.Once

func securityOriginToStringFunction_Set() {
	securityOriginToStringFunction_Once.Do(func() {
		securityOriginToStringFunction = gi.FunctionInvokerNew("WebKit2", "to_string")
	})
}

var toStringSecurityOriginInvoker *gi.Function

// ToString is a representation of the C type webkit_security_origin_to_string.
func (recv *SecurityOrigin) ToString() string {
	securityOriginToStringFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := securityOriginToStringFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

var securityOriginUnrefFunction *gi.Function
var securityOriginUnrefFunction_Once sync.Once

func securityOriginUnrefFunction_Set() {
	securityOriginUnrefFunction_Once.Do(func() {
		securityOriginUnrefFunction = gi.FunctionInvokerNew("WebKit2", "unref")
	})
}

var unrefSecurityOriginInvoker *gi.Function

// Unref is a representation of the C type webkit_security_origin_unref.
func (recv *SecurityOrigin) Unref() {
	securityOriginUnrefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	securityOriginUnrefFunction.Invoke(inArgs[:], nil)

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
var userContentFilterGetIdentifierFunction_Once sync.Once

func userContentFilterGetIdentifierFunction_Set() {
	userContentFilterGetIdentifierFunction_Once.Do(func() {
		userContentFilterGetIdentifierFunction = gi.FunctionInvokerNew("WebKit2", "get_identifier")
	})
}

var getIdentifierUserContentFilterInvoker *gi.Function

// GetIdentifier is a representation of the C type webkit_user_content_filter_get_identifier.
func (recv *UserContentFilter) GetIdentifier() string {
	userContentFilterGetIdentifierFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := userContentFilterGetIdentifierFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

var userContentFilterRefFunction *gi.Function
var userContentFilterRefFunction_Once sync.Once

func userContentFilterRefFunction_Set() {
	userContentFilterRefFunction_Once.Do(func() {
		userContentFilterRefFunction = gi.FunctionInvokerNew("WebKit2", "ref")
	})
}

var refUserContentFilterInvoker *gi.Function

// Ref is a representation of the C type webkit_user_content_filter_ref.
func (recv *UserContentFilter) Ref() *UserContentFilter {
	userContentFilterRefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := userContentFilterRefFunction.Invoke(inArgs[:], nil)

	retGo := &UserContentFilter{native: ret.Pointer()}

	return retGo
}

var userContentFilterUnrefFunction *gi.Function
var userContentFilterUnrefFunction_Once sync.Once

func userContentFilterUnrefFunction_Set() {
	userContentFilterUnrefFunction_Once.Do(func() {
		userContentFilterUnrefFunction = gi.FunctionInvokerNew("WebKit2", "unref")
	})
}

var unrefUserContentFilterInvoker *gi.Function

// Unref is a representation of the C type webkit_user_content_filter_unref.
func (recv *UserContentFilter) Unref() {
	userContentFilterUnrefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	userContentFilterUnrefFunction.Invoke(inArgs[:], nil)

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
var userScriptRefFunction_Once sync.Once

func userScriptRefFunction_Set() {
	userScriptRefFunction_Once.Do(func() {
		userScriptRefFunction = gi.FunctionInvokerNew("WebKit2", "ref")
	})
}

var refUserScriptInvoker *gi.Function

// Ref is a representation of the C type webkit_user_script_ref.
func (recv *UserScript) Ref() *UserScript {
	userScriptRefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := userScriptRefFunction.Invoke(inArgs[:], nil)

	retGo := &UserScript{native: ret.Pointer()}

	return retGo
}

var userScriptUnrefFunction *gi.Function
var userScriptUnrefFunction_Once sync.Once

func userScriptUnrefFunction_Set() {
	userScriptUnrefFunction_Once.Do(func() {
		userScriptUnrefFunction = gi.FunctionInvokerNew("WebKit2", "unref")
	})
}

var unrefUserScriptInvoker *gi.Function

// Unref is a representation of the C type webkit_user_script_unref.
func (recv *UserScript) Unref() {
	userScriptUnrefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	userScriptUnrefFunction.Invoke(inArgs[:], nil)

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
var userStyleSheetRefFunction_Once sync.Once

func userStyleSheetRefFunction_Set() {
	userStyleSheetRefFunction_Once.Do(func() {
		userStyleSheetRefFunction = gi.FunctionInvokerNew("WebKit2", "ref")
	})
}

var refUserStyleSheetInvoker *gi.Function

// Ref is a representation of the C type webkit_user_style_sheet_ref.
func (recv *UserStyleSheet) Ref() *UserStyleSheet {
	userStyleSheetRefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := userStyleSheetRefFunction.Invoke(inArgs[:], nil)

	retGo := &UserStyleSheet{native: ret.Pointer()}

	return retGo
}

var userStyleSheetUnrefFunction *gi.Function
var userStyleSheetUnrefFunction_Once sync.Once

func userStyleSheetUnrefFunction_Set() {
	userStyleSheetUnrefFunction_Once.Do(func() {
		userStyleSheetUnrefFunction = gi.FunctionInvokerNew("WebKit2", "unref")
	})
}

var unrefUserStyleSheetInvoker *gi.Function

// Unref is a representation of the C type webkit_user_style_sheet_unref.
func (recv *UserStyleSheet) Unref() {
	userStyleSheetUnrefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	userStyleSheetUnrefFunction.Invoke(inArgs[:], nil)

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
var webViewSessionStateRefFunction_Once sync.Once

func webViewSessionStateRefFunction_Set() {
	webViewSessionStateRefFunction_Once.Do(func() {
		webViewSessionStateRefFunction = gi.FunctionInvokerNew("WebKit2", "ref")
	})
}

var refWebViewSessionStateInvoker *gi.Function

// Ref is a representation of the C type webkit_web_view_session_state_ref.
func (recv *WebViewSessionState) Ref() *WebViewSessionState {
	webViewSessionStateRefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := webViewSessionStateRefFunction.Invoke(inArgs[:], nil)

	retGo := &WebViewSessionState{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'webkit_web_view_session_state_serialize' : return type 'GLib.Bytes' not supported

var webViewSessionStateUnrefFunction *gi.Function
var webViewSessionStateUnrefFunction_Once sync.Once

func webViewSessionStateUnrefFunction_Set() {
	webViewSessionStateUnrefFunction_Once.Do(func() {
		webViewSessionStateUnrefFunction = gi.FunctionInvokerNew("WebKit2", "unref")
	})
}

var unrefWebViewSessionStateInvoker *gi.Function

// Unref is a representation of the C type webkit_web_view_session_state_unref.
func (recv *WebViewSessionState) Unref() {
	webViewSessionStateUnrefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	webViewSessionStateUnrefFunction.Invoke(inArgs[:], nil)

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
var websiteDataGetNameFunction_Once sync.Once

func websiteDataGetNameFunction_Set() {
	websiteDataGetNameFunction_Once.Do(func() {
		websiteDataGetNameFunction = gi.FunctionInvokerNew("WebKit2", "get_name")
	})
}

var getNameWebsiteDataInvoker *gi.Function

// GetName is a representation of the C type webkit_website_data_get_name.
func (recv *WebsiteData) GetName() string {
	websiteDataGetNameFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := websiteDataGetNameFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'webkit_website_data_get_size' : parameter 'types' of type 'WebsiteDataTypes' not supported

// UNSUPPORTED : C value 'webkit_website_data_get_types' : return type 'WebsiteDataTypes' not supported

var websiteDataRefFunction *gi.Function
var websiteDataRefFunction_Once sync.Once

func websiteDataRefFunction_Set() {
	websiteDataRefFunction_Once.Do(func() {
		websiteDataRefFunction = gi.FunctionInvokerNew("WebKit2", "ref")
	})
}

var refWebsiteDataInvoker *gi.Function

// Ref is a representation of the C type webkit_website_data_ref.
func (recv *WebsiteData) Ref() *WebsiteData {
	websiteDataRefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := websiteDataRefFunction.Invoke(inArgs[:], nil)

	retGo := &WebsiteData{native: ret.Pointer()}

	return retGo
}

var websiteDataUnrefFunction *gi.Function
var websiteDataUnrefFunction_Once sync.Once

func websiteDataUnrefFunction_Set() {
	websiteDataUnrefFunction_Once.Do(func() {
		websiteDataUnrefFunction = gi.FunctionInvokerNew("WebKit2", "unref")
	})
}

var unrefWebsiteDataInvoker *gi.Function

// Unref is a representation of the C type webkit_website_data_unref.
func (recv *WebsiteData) Unref() {
	websiteDataUnrefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	websiteDataUnrefFunction.Invoke(inArgs[:], nil)

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
