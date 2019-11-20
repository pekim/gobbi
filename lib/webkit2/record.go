// Code generated - DO NOT EDIT.

package webkit2

import (
	gi "github.com/pekim/gobbi/internal/gi"
	"sync"
)

var ApplicationInfoStruct *gi.Struct
var ApplicationInfoStructOnce sync.Once

func ApplicationInfoStructSet() {
	ApplicationInfoStructOnce.Do(func() {
		ApplicationInfoStruct = gi.StructNew("WebKit2", "ApplicationInfo")
	})
}

type ApplicationInfo struct {
	native uintptr
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

var AuthenticationRequestClassStruct *gi.Struct
var AuthenticationRequestClassStructOnce sync.Once

func AuthenticationRequestClassStructSet() {
	AuthenticationRequestClassStructOnce.Do(func() {
		AuthenticationRequestClassStruct = gi.StructNew("WebKit2", "AuthenticationRequestClass")
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

var AuthenticationRequestPrivateStruct *gi.Struct
var AuthenticationRequestPrivateStructOnce sync.Once

func AuthenticationRequestPrivateStructSet() {
	AuthenticationRequestPrivateStructOnce.Do(func() {
		AuthenticationRequestPrivateStruct = gi.StructNew("WebKit2", "AuthenticationRequestPrivate")
	})
}

type AuthenticationRequestPrivate struct {
	native uintptr
}

var AutomationSessionClassStruct *gi.Struct
var AutomationSessionClassStructOnce sync.Once

func AutomationSessionClassStructSet() {
	AutomationSessionClassStructOnce.Do(func() {
		AutomationSessionClassStruct = gi.StructNew("WebKit2", "AutomationSessionClass")
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

var AutomationSessionPrivateStruct *gi.Struct
var AutomationSessionPrivateStructOnce sync.Once

func AutomationSessionPrivateStructSet() {
	AutomationSessionPrivateStructOnce.Do(func() {
		AutomationSessionPrivateStruct = gi.StructNew("WebKit2", "AutomationSessionPrivate")
	})
}

type AutomationSessionPrivate struct {
	native uintptr
}

var BackForwardListClassStruct *gi.Struct
var BackForwardListClassStructOnce sync.Once

func BackForwardListClassStructSet() {
	BackForwardListClassStructOnce.Do(func() {
		BackForwardListClassStruct = gi.StructNew("WebKit2", "BackForwardListClass")
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

var BackForwardListItemClassStruct *gi.Struct
var BackForwardListItemClassStructOnce sync.Once

func BackForwardListItemClassStructSet() {
	BackForwardListItemClassStructOnce.Do(func() {
		BackForwardListItemClassStruct = gi.StructNew("WebKit2", "BackForwardListItemClass")
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

var BackForwardListItemPrivateStruct *gi.Struct
var BackForwardListItemPrivateStructOnce sync.Once

func BackForwardListItemPrivateStructSet() {
	BackForwardListItemPrivateStructOnce.Do(func() {
		BackForwardListItemPrivateStruct = gi.StructNew("WebKit2", "BackForwardListItemPrivate")
	})
}

type BackForwardListItemPrivate struct {
	native uintptr
}

var BackForwardListPrivateStruct *gi.Struct
var BackForwardListPrivateStructOnce sync.Once

func BackForwardListPrivateStructSet() {
	BackForwardListPrivateStructOnce.Do(func() {
		BackForwardListPrivateStruct = gi.StructNew("WebKit2", "BackForwardListPrivate")
	})
}

type BackForwardListPrivate struct {
	native uintptr
}

var ColorChooserRequestClassStruct *gi.Struct
var ColorChooserRequestClassStructOnce sync.Once

func ColorChooserRequestClassStructSet() {
	ColorChooserRequestClassStructOnce.Do(func() {
		ColorChooserRequestClassStruct = gi.StructNew("WebKit2", "ColorChooserRequestClass")
	})
}

type ColorChooserRequestClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
}

var ColorChooserRequestPrivateStruct *gi.Struct
var ColorChooserRequestPrivateStructOnce sync.Once

func ColorChooserRequestPrivateStructSet() {
	ColorChooserRequestPrivateStructOnce.Do(func() {
		ColorChooserRequestPrivateStruct = gi.StructNew("WebKit2", "ColorChooserRequestPrivate")
	})
}

type ColorChooserRequestPrivate struct {
	native uintptr
}

var ContextMenuClassStruct *gi.Struct
var ContextMenuClassStructOnce sync.Once

func ContextMenuClassStructSet() {
	ContextMenuClassStructOnce.Do(func() {
		ContextMenuClassStruct = gi.StructNew("WebKit2", "ContextMenuClass")
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

var ContextMenuItemClassStruct *gi.Struct
var ContextMenuItemClassStructOnce sync.Once

func ContextMenuItemClassStructSet() {
	ContextMenuItemClassStructOnce.Do(func() {
		ContextMenuItemClassStruct = gi.StructNew("WebKit2", "ContextMenuItemClass")
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

var ContextMenuItemPrivateStruct *gi.Struct
var ContextMenuItemPrivateStructOnce sync.Once

func ContextMenuItemPrivateStructSet() {
	ContextMenuItemPrivateStructOnce.Do(func() {
		ContextMenuItemPrivateStruct = gi.StructNew("WebKit2", "ContextMenuItemPrivate")
	})
}

type ContextMenuItemPrivate struct {
	native uintptr
}

var ContextMenuPrivateStruct *gi.Struct
var ContextMenuPrivateStructOnce sync.Once

func ContextMenuPrivateStructSet() {
	ContextMenuPrivateStructOnce.Do(func() {
		ContextMenuPrivateStruct = gi.StructNew("WebKit2", "ContextMenuPrivate")
	})
}

type ContextMenuPrivate struct {
	native uintptr
}

var CookieManagerClassStruct *gi.Struct
var CookieManagerClassStructOnce sync.Once

func CookieManagerClassStructSet() {
	CookieManagerClassStructOnce.Do(func() {
		CookieManagerClassStruct = gi.StructNew("WebKit2", "CookieManagerClass")
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

var CookieManagerPrivateStruct *gi.Struct
var CookieManagerPrivateStructOnce sync.Once

func CookieManagerPrivateStructSet() {
	CookieManagerPrivateStructOnce.Do(func() {
		CookieManagerPrivateStruct = gi.StructNew("WebKit2", "CookieManagerPrivate")
	})
}

type CookieManagerPrivate struct {
	native uintptr
}

var CredentialStruct *gi.Struct
var CredentialStructOnce sync.Once

func CredentialStructSet() {
	CredentialStructOnce.Do(func() {
		CredentialStruct = gi.StructNew("WebKit2", "Credential")
	})
}

type Credential struct {
	native uintptr
}

// UNSUPPORTED : C value 'webkit_credential_new' : parameter 'persistence' of type 'CredentialPersistence' not supported

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

var DeviceInfoPermissionRequestClassStruct *gi.Struct
var DeviceInfoPermissionRequestClassStructOnce sync.Once

func DeviceInfoPermissionRequestClassStructSet() {
	DeviceInfoPermissionRequestClassStructOnce.Do(func() {
		DeviceInfoPermissionRequestClassStruct = gi.StructNew("WebKit2", "DeviceInfoPermissionRequestClass")
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

var DeviceInfoPermissionRequestPrivateStruct *gi.Struct
var DeviceInfoPermissionRequestPrivateStructOnce sync.Once

func DeviceInfoPermissionRequestPrivateStructSet() {
	DeviceInfoPermissionRequestPrivateStructOnce.Do(func() {
		DeviceInfoPermissionRequestPrivateStruct = gi.StructNew("WebKit2", "DeviceInfoPermissionRequestPrivate")
	})
}

type DeviceInfoPermissionRequestPrivate struct {
	native uintptr
}

var DownloadClassStruct *gi.Struct
var DownloadClassStructOnce sync.Once

func DownloadClassStructSet() {
	DownloadClassStructOnce.Do(func() {
		DownloadClassStruct = gi.StructNew("WebKit2", "DownloadClass")
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

var DownloadPrivateStruct *gi.Struct
var DownloadPrivateStructOnce sync.Once

func DownloadPrivateStructSet() {
	DownloadPrivateStructOnce.Do(func() {
		DownloadPrivateStruct = gi.StructNew("WebKit2", "DownloadPrivate")
	})
}

type DownloadPrivate struct {
	native uintptr
}

var EditorStateClassStruct *gi.Struct
var EditorStateClassStructOnce sync.Once

func EditorStateClassStructSet() {
	EditorStateClassStructOnce.Do(func() {
		EditorStateClassStruct = gi.StructNew("WebKit2", "EditorStateClass")
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

var EditorStatePrivateStruct *gi.Struct
var EditorStatePrivateStructOnce sync.Once

func EditorStatePrivateStructSet() {
	EditorStatePrivateStructOnce.Do(func() {
		EditorStatePrivateStruct = gi.StructNew("WebKit2", "EditorStatePrivate")
	})
}

type EditorStatePrivate struct {
	native uintptr
}

var FaviconDatabaseClassStruct *gi.Struct
var FaviconDatabaseClassStructOnce sync.Once

func FaviconDatabaseClassStructSet() {
	FaviconDatabaseClassStructOnce.Do(func() {
		FaviconDatabaseClassStruct = gi.StructNew("WebKit2", "FaviconDatabaseClass")
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

var FaviconDatabasePrivateStruct *gi.Struct
var FaviconDatabasePrivateStructOnce sync.Once

func FaviconDatabasePrivateStructSet() {
	FaviconDatabasePrivateStructOnce.Do(func() {
		FaviconDatabasePrivateStruct = gi.StructNew("WebKit2", "FaviconDatabasePrivate")
	})
}

type FaviconDatabasePrivate struct {
	native uintptr
}

var FileChooserRequestClassStruct *gi.Struct
var FileChooserRequestClassStructOnce sync.Once

func FileChooserRequestClassStructSet() {
	FileChooserRequestClassStructOnce.Do(func() {
		FileChooserRequestClassStruct = gi.StructNew("WebKit2", "FileChooserRequestClass")
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

var FileChooserRequestPrivateStruct *gi.Struct
var FileChooserRequestPrivateStructOnce sync.Once

func FileChooserRequestPrivateStructSet() {
	FileChooserRequestPrivateStructOnce.Do(func() {
		FileChooserRequestPrivateStruct = gi.StructNew("WebKit2", "FileChooserRequestPrivate")
	})
}

type FileChooserRequestPrivate struct {
	native uintptr
}

var FindControllerClassStruct *gi.Struct
var FindControllerClassStructOnce sync.Once

func FindControllerClassStructSet() {
	FindControllerClassStructOnce.Do(func() {
		FindControllerClassStruct = gi.StructNew("WebKit2", "FindControllerClass")
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

var FindControllerPrivateStruct *gi.Struct
var FindControllerPrivateStructOnce sync.Once

func FindControllerPrivateStructSet() {
	FindControllerPrivateStructOnce.Do(func() {
		FindControllerPrivateStruct = gi.StructNew("WebKit2", "FindControllerPrivate")
	})
}

type FindControllerPrivate struct {
	native uintptr
}

var FormSubmissionRequestClassStruct *gi.Struct
var FormSubmissionRequestClassStructOnce sync.Once

func FormSubmissionRequestClassStructSet() {
	FormSubmissionRequestClassStructOnce.Do(func() {
		FormSubmissionRequestClassStruct = gi.StructNew("WebKit2", "FormSubmissionRequestClass")
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

var FormSubmissionRequestPrivateStruct *gi.Struct
var FormSubmissionRequestPrivateStructOnce sync.Once

func FormSubmissionRequestPrivateStructSet() {
	FormSubmissionRequestPrivateStructOnce.Do(func() {
		FormSubmissionRequestPrivateStruct = gi.StructNew("WebKit2", "FormSubmissionRequestPrivate")
	})
}

type FormSubmissionRequestPrivate struct {
	native uintptr
}

var GeolocationManagerClassStruct *gi.Struct
var GeolocationManagerClassStructOnce sync.Once

func GeolocationManagerClassStructSet() {
	GeolocationManagerClassStructOnce.Do(func() {
		GeolocationManagerClassStruct = gi.StructNew("WebKit2", "GeolocationManagerClass")
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

var GeolocationManagerPrivateStruct *gi.Struct
var GeolocationManagerPrivateStructOnce sync.Once

func GeolocationManagerPrivateStructSet() {
	GeolocationManagerPrivateStructOnce.Do(func() {
		GeolocationManagerPrivateStruct = gi.StructNew("WebKit2", "GeolocationManagerPrivate")
	})
}

type GeolocationManagerPrivate struct {
	native uintptr
}

var GeolocationPermissionRequestClassStruct *gi.Struct
var GeolocationPermissionRequestClassStructOnce sync.Once

func GeolocationPermissionRequestClassStructSet() {
	GeolocationPermissionRequestClassStructOnce.Do(func() {
		GeolocationPermissionRequestClassStruct = gi.StructNew("WebKit2", "GeolocationPermissionRequestClass")
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

var GeolocationPermissionRequestPrivateStruct *gi.Struct
var GeolocationPermissionRequestPrivateStructOnce sync.Once

func GeolocationPermissionRequestPrivateStructSet() {
	GeolocationPermissionRequestPrivateStructOnce.Do(func() {
		GeolocationPermissionRequestPrivateStruct = gi.StructNew("WebKit2", "GeolocationPermissionRequestPrivate")
	})
}

type GeolocationPermissionRequestPrivate struct {
	native uintptr
}

var GeolocationPositionStruct *gi.Struct
var GeolocationPositionStructOnce sync.Once

func GeolocationPositionStructSet() {
	GeolocationPositionStructOnce.Do(func() {
		GeolocationPositionStruct = gi.StructNew("WebKit2", "GeolocationPosition")
	})
}

type GeolocationPosition struct {
	native uintptr
}

// UNSUPPORTED : C value 'webkit_geolocation_position_new' : parameter 'latitude' of type 'gdouble' not supported

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

var HitTestResultClassStruct *gi.Struct
var HitTestResultClassStructOnce sync.Once

func HitTestResultClassStructSet() {
	HitTestResultClassStructOnce.Do(func() {
		HitTestResultClassStruct = gi.StructNew("WebKit2", "HitTestResultClass")
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

var HitTestResultPrivateStruct *gi.Struct
var HitTestResultPrivateStructOnce sync.Once

func HitTestResultPrivateStructSet() {
	HitTestResultPrivateStructOnce.Do(func() {
		HitTestResultPrivateStruct = gi.StructNew("WebKit2", "HitTestResultPrivate")
	})
}

type HitTestResultPrivate struct {
	native uintptr
}

var InstallMissingMediaPluginsPermissionRequestClassStruct *gi.Struct
var InstallMissingMediaPluginsPermissionRequestClassStructOnce sync.Once

func InstallMissingMediaPluginsPermissionRequestClassStructSet() {
	InstallMissingMediaPluginsPermissionRequestClassStructOnce.Do(func() {
		InstallMissingMediaPluginsPermissionRequestClassStruct = gi.StructNew("WebKit2", "InstallMissingMediaPluginsPermissionRequestClass")
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

var InstallMissingMediaPluginsPermissionRequestPrivateStruct *gi.Struct
var InstallMissingMediaPluginsPermissionRequestPrivateStructOnce sync.Once

func InstallMissingMediaPluginsPermissionRequestPrivateStructSet() {
	InstallMissingMediaPluginsPermissionRequestPrivateStructOnce.Do(func() {
		InstallMissingMediaPluginsPermissionRequestPrivateStruct = gi.StructNew("WebKit2", "InstallMissingMediaPluginsPermissionRequestPrivate")
	})
}

type InstallMissingMediaPluginsPermissionRequestPrivate struct {
	native uintptr
}

var JavascriptResultStruct *gi.Struct
var JavascriptResultStructOnce sync.Once

func JavascriptResultStructSet() {
	JavascriptResultStructOnce.Do(func() {
		JavascriptResultStruct = gi.StructNew("WebKit2", "JavascriptResult")
	})
}

type JavascriptResult struct {
	native uintptr
}

// UNSUPPORTED : C value 'webkit_javascript_result_get_global_context' : return type 'JavaScriptCore.GlobalContextRef' not supported

// UNSUPPORTED : C value 'webkit_javascript_result_get_js_value' : return type 'JavaScriptCore.Value' not supported

// UNSUPPORTED : C value 'webkit_javascript_result_get_value' : return type 'JavaScriptCore.ValueRef' not supported

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

var MimeInfoStruct *gi.Struct
var MimeInfoStructOnce sync.Once

func MimeInfoStructSet() {
	MimeInfoStructOnce.Do(func() {
		MimeInfoStruct = gi.StructNew("WebKit2", "MimeInfo")
	})
}

type MimeInfo struct {
	native uintptr
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

var NavigationActionStruct *gi.Struct
var NavigationActionStructOnce sync.Once

func NavigationActionStructSet() {
	NavigationActionStructOnce.Do(func() {
		NavigationActionStruct = gi.StructNew("WebKit2", "NavigationAction")
	})
}

type NavigationAction struct {
	native uintptr
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

var NavigationPolicyDecisionClassStruct *gi.Struct
var NavigationPolicyDecisionClassStructOnce sync.Once

func NavigationPolicyDecisionClassStructSet() {
	NavigationPolicyDecisionClassStructOnce.Do(func() {
		NavigationPolicyDecisionClassStruct = gi.StructNew("WebKit2", "NavigationPolicyDecisionClass")
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

var NavigationPolicyDecisionPrivateStruct *gi.Struct
var NavigationPolicyDecisionPrivateStructOnce sync.Once

func NavigationPolicyDecisionPrivateStructSet() {
	NavigationPolicyDecisionPrivateStructOnce.Do(func() {
		NavigationPolicyDecisionPrivateStruct = gi.StructNew("WebKit2", "NavigationPolicyDecisionPrivate")
	})
}

type NavigationPolicyDecisionPrivate struct {
	native uintptr
}

var NetworkProxySettingsStruct *gi.Struct
var NetworkProxySettingsStructOnce sync.Once

func NetworkProxySettingsStructSet() {
	NetworkProxySettingsStructOnce.Do(func() {
		NetworkProxySettingsStruct = gi.StructNew("WebKit2", "NetworkProxySettings")
	})
}

type NetworkProxySettings struct {
	native uintptr
}

// UNSUPPORTED : C value 'webkit_network_proxy_settings_new' : parameter 'ignore_hosts' has no type

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

var NotificationClassStruct *gi.Struct
var NotificationClassStructOnce sync.Once

func NotificationClassStructSet() {
	NotificationClassStructOnce.Do(func() {
		NotificationClassStruct = gi.StructNew("WebKit2", "NotificationClass")
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

var NotificationPermissionRequestClassStruct *gi.Struct
var NotificationPermissionRequestClassStructOnce sync.Once

func NotificationPermissionRequestClassStructSet() {
	NotificationPermissionRequestClassStructOnce.Do(func() {
		NotificationPermissionRequestClassStruct = gi.StructNew("WebKit2", "NotificationPermissionRequestClass")
	})
}

type NotificationPermissionRequestClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
}

var NotificationPermissionRequestPrivateStruct *gi.Struct
var NotificationPermissionRequestPrivateStructOnce sync.Once

func NotificationPermissionRequestPrivateStructSet() {
	NotificationPermissionRequestPrivateStructOnce.Do(func() {
		NotificationPermissionRequestPrivateStruct = gi.StructNew("WebKit2", "NotificationPermissionRequestPrivate")
	})
}

type NotificationPermissionRequestPrivate struct {
	native uintptr
}

var NotificationPrivateStruct *gi.Struct
var NotificationPrivateStructOnce sync.Once

func NotificationPrivateStructSet() {
	NotificationPrivateStructOnce.Do(func() {
		NotificationPrivateStruct = gi.StructNew("WebKit2", "NotificationPrivate")
	})
}

type NotificationPrivate struct {
	native uintptr
}

var OptionMenuClassStruct *gi.Struct
var OptionMenuClassStructOnce sync.Once

func OptionMenuClassStructSet() {
	OptionMenuClassStructOnce.Do(func() {
		OptionMenuClassStruct = gi.StructNew("WebKit2", "OptionMenuClass")
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

var OptionMenuItemStruct *gi.Struct
var OptionMenuItemStructOnce sync.Once

func OptionMenuItemStructSet() {
	OptionMenuItemStructOnce.Do(func() {
		OptionMenuItemStruct = gi.StructNew("WebKit2", "OptionMenuItem")
	})
}

type OptionMenuItem struct {
	native uintptr
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

var OptionMenuPrivateStruct *gi.Struct
var OptionMenuPrivateStructOnce sync.Once

func OptionMenuPrivateStructSet() {
	OptionMenuPrivateStructOnce.Do(func() {
		OptionMenuPrivateStruct = gi.StructNew("WebKit2", "OptionMenuPrivate")
	})
}

type OptionMenuPrivate struct {
	native uintptr
}

var PermissionRequestIfaceStruct *gi.Struct
var PermissionRequestIfaceStructOnce sync.Once

func PermissionRequestIfaceStructSet() {
	PermissionRequestIfaceStructOnce.Do(func() {
		PermissionRequestIfaceStruct = gi.StructNew("WebKit2", "PermissionRequestIface")
	})
}

type PermissionRequestIface struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_interface' : no Go type for 'GObject.TypeInterface'
	// UNSUPPORTED : C value 'allow' : missing Type
	// UNSUPPORTED : C value 'deny' : missing Type
}

var PluginClassStruct *gi.Struct
var PluginClassStructOnce sync.Once

func PluginClassStructSet() {
	PluginClassStructOnce.Do(func() {
		PluginClassStruct = gi.StructNew("WebKit2", "PluginClass")
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

var PluginPrivateStruct *gi.Struct
var PluginPrivateStructOnce sync.Once

func PluginPrivateStructSet() {
	PluginPrivateStructOnce.Do(func() {
		PluginPrivateStruct = gi.StructNew("WebKit2", "PluginPrivate")
	})
}

type PluginPrivate struct {
	native uintptr
}

var PolicyDecisionClassStruct *gi.Struct
var PolicyDecisionClassStructOnce sync.Once

func PolicyDecisionClassStructSet() {
	PolicyDecisionClassStructOnce.Do(func() {
		PolicyDecisionClassStruct = gi.StructNew("WebKit2", "PolicyDecisionClass")
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

var PolicyDecisionPrivateStruct *gi.Struct
var PolicyDecisionPrivateStructOnce sync.Once

func PolicyDecisionPrivateStructSet() {
	PolicyDecisionPrivateStructOnce.Do(func() {
		PolicyDecisionPrivateStruct = gi.StructNew("WebKit2", "PolicyDecisionPrivate")
	})
}

type PolicyDecisionPrivate struct {
	native uintptr
}

var PrintCustomWidgetClassStruct *gi.Struct
var PrintCustomWidgetClassStructOnce sync.Once

func PrintCustomWidgetClassStructSet() {
	PrintCustomWidgetClassStructOnce.Do(func() {
		PrintCustomWidgetClassStruct = gi.StructNew("WebKit2", "PrintCustomWidgetClass")
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

var PrintCustomWidgetPrivateStruct *gi.Struct
var PrintCustomWidgetPrivateStructOnce sync.Once

func PrintCustomWidgetPrivateStructSet() {
	PrintCustomWidgetPrivateStructOnce.Do(func() {
		PrintCustomWidgetPrivateStruct = gi.StructNew("WebKit2", "PrintCustomWidgetPrivate")
	})
}

type PrintCustomWidgetPrivate struct {
	native uintptr
}

var PrintOperationClassStruct *gi.Struct
var PrintOperationClassStructOnce sync.Once

func PrintOperationClassStructSet() {
	PrintOperationClassStructOnce.Do(func() {
		PrintOperationClassStruct = gi.StructNew("WebKit2", "PrintOperationClass")
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

var PrintOperationPrivateStruct *gi.Struct
var PrintOperationPrivateStructOnce sync.Once

func PrintOperationPrivateStructSet() {
	PrintOperationPrivateStructOnce.Do(func() {
		PrintOperationPrivateStruct = gi.StructNew("WebKit2", "PrintOperationPrivate")
	})
}

type PrintOperationPrivate struct {
	native uintptr
}

var ResponsePolicyDecisionClassStruct *gi.Struct
var ResponsePolicyDecisionClassStructOnce sync.Once

func ResponsePolicyDecisionClassStructSet() {
	ResponsePolicyDecisionClassStructOnce.Do(func() {
		ResponsePolicyDecisionClassStruct = gi.StructNew("WebKit2", "ResponsePolicyDecisionClass")
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

var ResponsePolicyDecisionPrivateStruct *gi.Struct
var ResponsePolicyDecisionPrivateStructOnce sync.Once

func ResponsePolicyDecisionPrivateStructSet() {
	ResponsePolicyDecisionPrivateStructOnce.Do(func() {
		ResponsePolicyDecisionPrivateStruct = gi.StructNew("WebKit2", "ResponsePolicyDecisionPrivate")
	})
}

type ResponsePolicyDecisionPrivate struct {
	native uintptr
}

var ScriptDialogStruct *gi.Struct
var ScriptDialogStructOnce sync.Once

func ScriptDialogStructSet() {
	ScriptDialogStructOnce.Do(func() {
		ScriptDialogStruct = gi.StructNew("WebKit2", "ScriptDialog")
	})
}

type ScriptDialog struct {
	native uintptr
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

var SecurityManagerClassStruct *gi.Struct
var SecurityManagerClassStructOnce sync.Once

func SecurityManagerClassStructSet() {
	SecurityManagerClassStructOnce.Do(func() {
		SecurityManagerClassStruct = gi.StructNew("WebKit2", "SecurityManagerClass")
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

var SecurityManagerPrivateStruct *gi.Struct
var SecurityManagerPrivateStructOnce sync.Once

func SecurityManagerPrivateStructSet() {
	SecurityManagerPrivateStructOnce.Do(func() {
		SecurityManagerPrivateStruct = gi.StructNew("WebKit2", "SecurityManagerPrivate")
	})
}

type SecurityManagerPrivate struct {
	native uintptr
}

var SecurityOriginStruct *gi.Struct
var SecurityOriginStructOnce sync.Once

func SecurityOriginStructSet() {
	SecurityOriginStructOnce.Do(func() {
		SecurityOriginStruct = gi.StructNew("WebKit2", "SecurityOrigin")
	})
}

type SecurityOrigin struct {
	native uintptr
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

var SettingsClassStruct *gi.Struct
var SettingsClassStructOnce sync.Once

func SettingsClassStructSet() {
	SettingsClassStructOnce.Do(func() {
		SettingsClassStruct = gi.StructNew("WebKit2", "SettingsClass")
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

var SettingsPrivateStruct *gi.Struct
var SettingsPrivateStructOnce sync.Once

func SettingsPrivateStructSet() {
	SettingsPrivateStructOnce.Do(func() {
		SettingsPrivateStruct = gi.StructNew("WebKit2", "SettingsPrivate")
	})
}

type SettingsPrivate struct {
	native uintptr
}

var URIRequestClassStruct *gi.Struct
var URIRequestClassStructOnce sync.Once

func URIRequestClassStructSet() {
	URIRequestClassStructOnce.Do(func() {
		URIRequestClassStruct = gi.StructNew("WebKit2", "URIRequestClass")
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

var URIRequestPrivateStruct *gi.Struct
var URIRequestPrivateStructOnce sync.Once

func URIRequestPrivateStructSet() {
	URIRequestPrivateStructOnce.Do(func() {
		URIRequestPrivateStruct = gi.StructNew("WebKit2", "URIRequestPrivate")
	})
}

type URIRequestPrivate struct {
	native uintptr
}

var URIResponseClassStruct *gi.Struct
var URIResponseClassStructOnce sync.Once

func URIResponseClassStructSet() {
	URIResponseClassStructOnce.Do(func() {
		URIResponseClassStruct = gi.StructNew("WebKit2", "URIResponseClass")
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

var URIResponsePrivateStruct *gi.Struct
var URIResponsePrivateStructOnce sync.Once

func URIResponsePrivateStructSet() {
	URIResponsePrivateStructOnce.Do(func() {
		URIResponsePrivateStruct = gi.StructNew("WebKit2", "URIResponsePrivate")
	})
}

type URIResponsePrivate struct {
	native uintptr
}

var URISchemeRequestClassStruct *gi.Struct
var URISchemeRequestClassStructOnce sync.Once

func URISchemeRequestClassStructSet() {
	URISchemeRequestClassStructOnce.Do(func() {
		URISchemeRequestClassStruct = gi.StructNew("WebKit2", "URISchemeRequestClass")
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

var URISchemeRequestPrivateStruct *gi.Struct
var URISchemeRequestPrivateStructOnce sync.Once

func URISchemeRequestPrivateStructSet() {
	URISchemeRequestPrivateStructOnce.Do(func() {
		URISchemeRequestPrivateStruct = gi.StructNew("WebKit2", "URISchemeRequestPrivate")
	})
}

type URISchemeRequestPrivate struct {
	native uintptr
}

var UserContentFilterStruct *gi.Struct
var UserContentFilterStructOnce sync.Once

func UserContentFilterStructSet() {
	UserContentFilterStructOnce.Do(func() {
		UserContentFilterStruct = gi.StructNew("WebKit2", "UserContentFilter")
	})
}

type UserContentFilter struct {
	native uintptr
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

var UserContentFilterStoreClassStruct *gi.Struct
var UserContentFilterStoreClassStructOnce sync.Once

func UserContentFilterStoreClassStructSet() {
	UserContentFilterStoreClassStructOnce.Do(func() {
		UserContentFilterStoreClassStruct = gi.StructNew("WebKit2", "UserContentFilterStoreClass")
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

var UserContentFilterStorePrivateStruct *gi.Struct
var UserContentFilterStorePrivateStructOnce sync.Once

func UserContentFilterStorePrivateStructSet() {
	UserContentFilterStorePrivateStructOnce.Do(func() {
		UserContentFilterStorePrivateStruct = gi.StructNew("WebKit2", "UserContentFilterStorePrivate")
	})
}

type UserContentFilterStorePrivate struct {
	native uintptr
}

var UserContentManagerClassStruct *gi.Struct
var UserContentManagerClassStructOnce sync.Once

func UserContentManagerClassStructSet() {
	UserContentManagerClassStructOnce.Do(func() {
		UserContentManagerClassStruct = gi.StructNew("WebKit2", "UserContentManagerClass")
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

var UserContentManagerPrivateStruct *gi.Struct
var UserContentManagerPrivateStructOnce sync.Once

func UserContentManagerPrivateStructSet() {
	UserContentManagerPrivateStructOnce.Do(func() {
		UserContentManagerPrivateStruct = gi.StructNew("WebKit2", "UserContentManagerPrivate")
	})
}

type UserContentManagerPrivate struct {
	native uintptr
}

var UserMediaPermissionRequestClassStruct *gi.Struct
var UserMediaPermissionRequestClassStructOnce sync.Once

func UserMediaPermissionRequestClassStructSet() {
	UserMediaPermissionRequestClassStructOnce.Do(func() {
		UserMediaPermissionRequestClassStruct = gi.StructNew("WebKit2", "UserMediaPermissionRequestClass")
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

var UserMediaPermissionRequestPrivateStruct *gi.Struct
var UserMediaPermissionRequestPrivateStructOnce sync.Once

func UserMediaPermissionRequestPrivateStructSet() {
	UserMediaPermissionRequestPrivateStructOnce.Do(func() {
		UserMediaPermissionRequestPrivateStruct = gi.StructNew("WebKit2", "UserMediaPermissionRequestPrivate")
	})
}

type UserMediaPermissionRequestPrivate struct {
	native uintptr
}

var UserScriptStruct *gi.Struct
var UserScriptStructOnce sync.Once

func UserScriptStructSet() {
	UserScriptStructOnce.Do(func() {
		UserScriptStruct = gi.StructNew("WebKit2", "UserScript")
	})
}

type UserScript struct {
	native uintptr
}

// UNSUPPORTED : C value 'webkit_user_script_new' : parameter 'injected_frames' of type 'UserContentInjectedFrames' not supported

// UNSUPPORTED : C value 'webkit_user_script_new_for_world' : parameter 'injected_frames' of type 'UserContentInjectedFrames' not supported

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

var UserStyleSheetStruct *gi.Struct
var UserStyleSheetStructOnce sync.Once

func UserStyleSheetStructSet() {
	UserStyleSheetStructOnce.Do(func() {
		UserStyleSheetStruct = gi.StructNew("WebKit2", "UserStyleSheet")
	})
}

type UserStyleSheet struct {
	native uintptr
}

// UNSUPPORTED : C value 'webkit_user_style_sheet_new' : parameter 'injected_frames' of type 'UserContentInjectedFrames' not supported

// UNSUPPORTED : C value 'webkit_user_style_sheet_new_for_world' : parameter 'injected_frames' of type 'UserContentInjectedFrames' not supported

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

var WebContextClassStruct *gi.Struct
var WebContextClassStructOnce sync.Once

func WebContextClassStructSet() {
	WebContextClassStructOnce.Do(func() {
		WebContextClassStruct = gi.StructNew("WebKit2", "WebContextClass")
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

var WebContextPrivateStruct *gi.Struct
var WebContextPrivateStructOnce sync.Once

func WebContextPrivateStructSet() {
	WebContextPrivateStructOnce.Do(func() {
		WebContextPrivateStruct = gi.StructNew("WebKit2", "WebContextPrivate")
	})
}

type WebContextPrivate struct {
	native uintptr
}

var WebInspectorClassStruct *gi.Struct
var WebInspectorClassStructOnce sync.Once

func WebInspectorClassStructSet() {
	WebInspectorClassStructOnce.Do(func() {
		WebInspectorClassStruct = gi.StructNew("WebKit2", "WebInspectorClass")
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

var WebInspectorPrivateStruct *gi.Struct
var WebInspectorPrivateStructOnce sync.Once

func WebInspectorPrivateStructSet() {
	WebInspectorPrivateStructOnce.Do(func() {
		WebInspectorPrivateStruct = gi.StructNew("WebKit2", "WebInspectorPrivate")
	})
}

type WebInspectorPrivate struct {
	native uintptr
}

var WebResourceClassStruct *gi.Struct
var WebResourceClassStructOnce sync.Once

func WebResourceClassStructSet() {
	WebResourceClassStructOnce.Do(func() {
		WebResourceClassStruct = gi.StructNew("WebKit2", "WebResourceClass")
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

var WebResourcePrivateStruct *gi.Struct
var WebResourcePrivateStructOnce sync.Once

func WebResourcePrivateStructSet() {
	WebResourcePrivateStructOnce.Do(func() {
		WebResourcePrivateStruct = gi.StructNew("WebKit2", "WebResourcePrivate")
	})
}

type WebResourcePrivate struct {
	native uintptr
}

var WebViewBaseClassStruct *gi.Struct
var WebViewBaseClassStructOnce sync.Once

func WebViewBaseClassStructSet() {
	WebViewBaseClassStructOnce.Do(func() {
		WebViewBaseClassStruct = gi.StructNew("WebKit2", "WebViewBaseClass")
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

var WebViewBasePrivateStruct *gi.Struct
var WebViewBasePrivateStructOnce sync.Once

func WebViewBasePrivateStructSet() {
	WebViewBasePrivateStructOnce.Do(func() {
		WebViewBasePrivateStruct = gi.StructNew("WebKit2", "WebViewBasePrivate")
	})
}

type WebViewBasePrivate struct {
	native uintptr
}

var WebViewClassStruct *gi.Struct
var WebViewClassStructOnce sync.Once

func WebViewClassStructSet() {
	WebViewClassStructOnce.Do(func() {
		WebViewClassStruct = gi.StructNew("WebKit2", "WebViewClass")
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

var WebViewPrivateStruct *gi.Struct
var WebViewPrivateStructOnce sync.Once

func WebViewPrivateStructSet() {
	WebViewPrivateStructOnce.Do(func() {
		WebViewPrivateStruct = gi.StructNew("WebKit2", "WebViewPrivate")
	})
}

type WebViewPrivate struct {
	native uintptr
}

var WebViewSessionStateStruct *gi.Struct
var WebViewSessionStateStructOnce sync.Once

func WebViewSessionStateStructSet() {
	WebViewSessionStateStructOnce.Do(func() {
		WebViewSessionStateStruct = gi.StructNew("WebKit2", "WebViewSessionState")
	})
}

type WebViewSessionState struct {
	native uintptr
}

// UNSUPPORTED : C value 'webkit_web_view_session_state_new' : parameter 'data' of type 'GLib.Bytes' not supported

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

var WebsiteDataStruct *gi.Struct
var WebsiteDataStructOnce sync.Once

func WebsiteDataStructSet() {
	WebsiteDataStructOnce.Do(func() {
		WebsiteDataStruct = gi.StructNew("WebKit2", "WebsiteData")
	})
}

type WebsiteData struct {
	native uintptr
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

var WebsiteDataManagerClassStruct *gi.Struct
var WebsiteDataManagerClassStructOnce sync.Once

func WebsiteDataManagerClassStructSet() {
	WebsiteDataManagerClassStructOnce.Do(func() {
		WebsiteDataManagerClassStruct = gi.StructNew("WebKit2", "WebsiteDataManagerClass")
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

var WebsiteDataManagerPrivateStruct *gi.Struct
var WebsiteDataManagerPrivateStructOnce sync.Once

func WebsiteDataManagerPrivateStructSet() {
	WebsiteDataManagerPrivateStructOnce.Do(func() {
		WebsiteDataManagerPrivateStruct = gi.StructNew("WebKit2", "WebsiteDataManagerPrivate")
	})
}

type WebsiteDataManagerPrivate struct {
	native uintptr
}

var WindowPropertiesClassStruct *gi.Struct
var WindowPropertiesClassStructOnce sync.Once

func WindowPropertiesClassStructSet() {
	WindowPropertiesClassStructOnce.Do(func() {
		WindowPropertiesClassStruct = gi.StructNew("WebKit2", "WindowPropertiesClass")
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

var WindowPropertiesPrivateStruct *gi.Struct
var WindowPropertiesPrivateStructOnce sync.Once

func WindowPropertiesPrivateStructSet() {
	WindowPropertiesPrivateStructOnce.Do(func() {
		WindowPropertiesPrivateStruct = gi.StructNew("WebKit2", "WindowPropertiesPrivate")
	})
}

type WindowPropertiesPrivate struct {
	native uintptr
}
