// Code generated - DO NOT EDIT.

package webkit2

import (
	gi "github.com/pekim/gobbi/internal/gi"
	"sync"
)

var applicationInfoStruct *gi.Struct
var applicationInfoStruct_Once sync.Once

func applicationInfoStruct_Set() error {
	var err error
	applicationInfoStruct_Once.Do(func() {
		applicationInfoStruct, err = gi.StructNew("WebKit2", "ApplicationInfo")
	})
	return err
}

type ApplicationInfo struct {
	native uintptr
}

var applicationInfoNewFunction *gi.Function
var applicationInfoNewFunction_Once sync.Once

func applicationInfoNewFunction_Set() error {
	var err error
	applicationInfoNewFunction_Once.Do(func() {
		err = applicationInfoStruct_Set()
		if err != nil {
			return
		}
		applicationInfoNewFunction, err = applicationInfoStruct.InvokerNew("new")
	})
	return err
}

// ApplicationInfoNew is a representation of the C type webkit_application_info_new.
func ApplicationInfoNew() (*ApplicationInfo, error) {

	var ret gi.Argument

	err := applicationInfoNewFunction_Set()
	if err == nil {
		ret = applicationInfoNewFunction.Invoke(nil, nil)
	}

	retGo := &ApplicationInfo{native: ret.Pointer()}

	return retGo, err
}

var applicationInfoGetNameFunction *gi.Function
var applicationInfoGetNameFunction_Once sync.Once

func applicationInfoGetNameFunction_Set() error {
	var err error
	applicationInfoGetNameFunction_Once.Do(func() {
		err = applicationInfoStruct_Set()
		if err != nil {
			return
		}
		applicationInfoGetNameFunction, err = applicationInfoStruct.InvokerNew("get_name")
	})
	return err
}

// GetName is a representation of the C type webkit_application_info_get_name.
func (recv *ApplicationInfo) GetName() (string, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := applicationInfoGetNameFunction_Set()
	if err == nil {
		ret = applicationInfoGetNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo, err
}

var applicationInfoGetVersionFunction *gi.Function
var applicationInfoGetVersionFunction_Once sync.Once

func applicationInfoGetVersionFunction_Set() error {
	var err error
	applicationInfoGetVersionFunction_Once.Do(func() {
		err = applicationInfoStruct_Set()
		if err != nil {
			return
		}
		applicationInfoGetVersionFunction, err = applicationInfoStruct.InvokerNew("get_version")
	})
	return err
}

// GetVersion is a representation of the C type webkit_application_info_get_version.
func (recv *ApplicationInfo) GetVersion() (uint64, uint64, uint64, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var outArgs [3]gi.Argument

	err := applicationInfoGetVersionFunction_Set()
	if err == nil {
		applicationInfoGetVersionFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := outArgs[0].Uint64()
	out1 := outArgs[1].Uint64()
	out2 := outArgs[2].Uint64()

	return out0, out1, out2, err
}

var applicationInfoRefFunction *gi.Function
var applicationInfoRefFunction_Once sync.Once

func applicationInfoRefFunction_Set() error {
	var err error
	applicationInfoRefFunction_Once.Do(func() {
		err = applicationInfoStruct_Set()
		if err != nil {
			return
		}
		applicationInfoRefFunction, err = applicationInfoStruct.InvokerNew("ref")
	})
	return err
}

// Ref is a representation of the C type webkit_application_info_ref.
func (recv *ApplicationInfo) Ref() (*ApplicationInfo, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := applicationInfoRefFunction_Set()
	if err == nil {
		ret = applicationInfoRefFunction.Invoke(inArgs[:], nil)
	}

	retGo := &ApplicationInfo{native: ret.Pointer()}

	return retGo, err
}

var applicationInfoSetNameFunction *gi.Function
var applicationInfoSetNameFunction_Once sync.Once

func applicationInfoSetNameFunction_Set() error {
	var err error
	applicationInfoSetNameFunction_Once.Do(func() {
		err = applicationInfoStruct_Set()
		if err != nil {
			return
		}
		applicationInfoSetNameFunction, err = applicationInfoStruct.InvokerNew("set_name")
	})
	return err
}

// SetName is a representation of the C type webkit_application_info_set_name.
func (recv *ApplicationInfo) SetName(name string) error {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(name)

	err := applicationInfoSetNameFunction_Set()
	if err == nil {
		applicationInfoSetNameFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var applicationInfoSetVersionFunction *gi.Function
var applicationInfoSetVersionFunction_Once sync.Once

func applicationInfoSetVersionFunction_Set() error {
	var err error
	applicationInfoSetVersionFunction_Once.Do(func() {
		err = applicationInfoStruct_Set()
		if err != nil {
			return
		}
		applicationInfoSetVersionFunction, err = applicationInfoStruct.InvokerNew("set_version")
	})
	return err
}

// SetVersion is a representation of the C type webkit_application_info_set_version.
func (recv *ApplicationInfo) SetVersion(major uint64, minor uint64, micro uint64) error {
	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetUint64(major)
	inArgs[2].SetUint64(minor)
	inArgs[3].SetUint64(micro)

	err := applicationInfoSetVersionFunction_Set()
	if err == nil {
		applicationInfoSetVersionFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var applicationInfoUnrefFunction *gi.Function
var applicationInfoUnrefFunction_Once sync.Once

func applicationInfoUnrefFunction_Set() error {
	var err error
	applicationInfoUnrefFunction_Once.Do(func() {
		err = applicationInfoStruct_Set()
		if err != nil {
			return
		}
		applicationInfoUnrefFunction, err = applicationInfoStruct.InvokerNew("unref")
	})
	return err
}

// Unref is a representation of the C type webkit_application_info_unref.
func (recv *ApplicationInfo) Unref() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := applicationInfoUnrefFunction_Set()
	if err == nil {
		applicationInfoUnrefFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var authenticationRequestClassStruct *gi.Struct
var authenticationRequestClassStruct_Once sync.Once

func authenticationRequestClassStruct_Set() error {
	var err error
	authenticationRequestClassStruct_Once.Do(func() {
		authenticationRequestClassStruct, err = gi.StructNew("WebKit2", "AuthenticationRequestClass")
	})
	return err
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
var authenticationRequestPrivateStruct_Once sync.Once

func authenticationRequestPrivateStruct_Set() error {
	var err error
	authenticationRequestPrivateStruct_Once.Do(func() {
		authenticationRequestPrivateStruct, err = gi.StructNew("WebKit2", "AuthenticationRequestPrivate")
	})
	return err
}

type AuthenticationRequestPrivate struct {
	native uintptr
}

var automationSessionClassStruct *gi.Struct
var automationSessionClassStruct_Once sync.Once

func automationSessionClassStruct_Set() error {
	var err error
	automationSessionClassStruct_Once.Do(func() {
		automationSessionClassStruct, err = gi.StructNew("WebKit2", "AutomationSessionClass")
	})
	return err
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
var automationSessionPrivateStruct_Once sync.Once

func automationSessionPrivateStruct_Set() error {
	var err error
	automationSessionPrivateStruct_Once.Do(func() {
		automationSessionPrivateStruct, err = gi.StructNew("WebKit2", "AutomationSessionPrivate")
	})
	return err
}

type AutomationSessionPrivate struct {
	native uintptr
}

var backForwardListClassStruct *gi.Struct
var backForwardListClassStruct_Once sync.Once

func backForwardListClassStruct_Set() error {
	var err error
	backForwardListClassStruct_Once.Do(func() {
		backForwardListClassStruct, err = gi.StructNew("WebKit2", "BackForwardListClass")
	})
	return err
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
var backForwardListItemClassStruct_Once sync.Once

func backForwardListItemClassStruct_Set() error {
	var err error
	backForwardListItemClassStruct_Once.Do(func() {
		backForwardListItemClassStruct, err = gi.StructNew("WebKit2", "BackForwardListItemClass")
	})
	return err
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
var backForwardListItemPrivateStruct_Once sync.Once

func backForwardListItemPrivateStruct_Set() error {
	var err error
	backForwardListItemPrivateStruct_Once.Do(func() {
		backForwardListItemPrivateStruct, err = gi.StructNew("WebKit2", "BackForwardListItemPrivate")
	})
	return err
}

type BackForwardListItemPrivate struct {
	native uintptr
}

var backForwardListPrivateStruct *gi.Struct
var backForwardListPrivateStruct_Once sync.Once

func backForwardListPrivateStruct_Set() error {
	var err error
	backForwardListPrivateStruct_Once.Do(func() {
		backForwardListPrivateStruct, err = gi.StructNew("WebKit2", "BackForwardListPrivate")
	})
	return err
}

type BackForwardListPrivate struct {
	native uintptr
}

var colorChooserRequestClassStruct *gi.Struct
var colorChooserRequestClassStruct_Once sync.Once

func colorChooserRequestClassStruct_Set() error {
	var err error
	colorChooserRequestClassStruct_Once.Do(func() {
		colorChooserRequestClassStruct, err = gi.StructNew("WebKit2", "ColorChooserRequestClass")
	})
	return err
}

type ColorChooserRequestClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
}

var colorChooserRequestPrivateStruct *gi.Struct
var colorChooserRequestPrivateStruct_Once sync.Once

func colorChooserRequestPrivateStruct_Set() error {
	var err error
	colorChooserRequestPrivateStruct_Once.Do(func() {
		colorChooserRequestPrivateStruct, err = gi.StructNew("WebKit2", "ColorChooserRequestPrivate")
	})
	return err
}

type ColorChooserRequestPrivate struct {
	native uintptr
}

var contextMenuClassStruct *gi.Struct
var contextMenuClassStruct_Once sync.Once

func contextMenuClassStruct_Set() error {
	var err error
	contextMenuClassStruct_Once.Do(func() {
		contextMenuClassStruct, err = gi.StructNew("WebKit2", "ContextMenuClass")
	})
	return err
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
var contextMenuItemClassStruct_Once sync.Once

func contextMenuItemClassStruct_Set() error {
	var err error
	contextMenuItemClassStruct_Once.Do(func() {
		contextMenuItemClassStruct, err = gi.StructNew("WebKit2", "ContextMenuItemClass")
	})
	return err
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
var contextMenuItemPrivateStruct_Once sync.Once

func contextMenuItemPrivateStruct_Set() error {
	var err error
	contextMenuItemPrivateStruct_Once.Do(func() {
		contextMenuItemPrivateStruct, err = gi.StructNew("WebKit2", "ContextMenuItemPrivate")
	})
	return err
}

type ContextMenuItemPrivate struct {
	native uintptr
}

var contextMenuPrivateStruct *gi.Struct
var contextMenuPrivateStruct_Once sync.Once

func contextMenuPrivateStruct_Set() error {
	var err error
	contextMenuPrivateStruct_Once.Do(func() {
		contextMenuPrivateStruct, err = gi.StructNew("WebKit2", "ContextMenuPrivate")
	})
	return err
}

type ContextMenuPrivate struct {
	native uintptr
}

var cookieManagerClassStruct *gi.Struct
var cookieManagerClassStruct_Once sync.Once

func cookieManagerClassStruct_Set() error {
	var err error
	cookieManagerClassStruct_Once.Do(func() {
		cookieManagerClassStruct, err = gi.StructNew("WebKit2", "CookieManagerClass")
	})
	return err
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
var cookieManagerPrivateStruct_Once sync.Once

func cookieManagerPrivateStruct_Set() error {
	var err error
	cookieManagerPrivateStruct_Once.Do(func() {
		cookieManagerPrivateStruct, err = gi.StructNew("WebKit2", "CookieManagerPrivate")
	})
	return err
}

type CookieManagerPrivate struct {
	native uintptr
}

var credentialStruct *gi.Struct
var credentialStruct_Once sync.Once

func credentialStruct_Set() error {
	var err error
	credentialStruct_Once.Do(func() {
		credentialStruct, err = gi.StructNew("WebKit2", "Credential")
	})
	return err
}

type Credential struct {
	native uintptr
}

// UNSUPPORTED : C value 'webkit_credential_new' : parameter 'persistence' of type 'CredentialPersistence' not supported

var credentialCopyFunction *gi.Function
var credentialCopyFunction_Once sync.Once

func credentialCopyFunction_Set() error {
	var err error
	credentialCopyFunction_Once.Do(func() {
		err = credentialStruct_Set()
		if err != nil {
			return
		}
		credentialCopyFunction, err = credentialStruct.InvokerNew("copy")
	})
	return err
}

// Copy is a representation of the C type webkit_credential_copy.
func (recv *Credential) Copy() (*Credential, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := credentialCopyFunction_Set()
	if err == nil {
		ret = credentialCopyFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Credential{native: ret.Pointer()}

	return retGo, err
}

var credentialFreeFunction *gi.Function
var credentialFreeFunction_Once sync.Once

func credentialFreeFunction_Set() error {
	var err error
	credentialFreeFunction_Once.Do(func() {
		err = credentialStruct_Set()
		if err != nil {
			return
		}
		credentialFreeFunction, err = credentialStruct.InvokerNew("free")
	})
	return err
}

// Free is a representation of the C type webkit_credential_free.
func (recv *Credential) Free() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := credentialFreeFunction_Set()
	if err == nil {
		credentialFreeFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var credentialGetPasswordFunction *gi.Function
var credentialGetPasswordFunction_Once sync.Once

func credentialGetPasswordFunction_Set() error {
	var err error
	credentialGetPasswordFunction_Once.Do(func() {
		err = credentialStruct_Set()
		if err != nil {
			return
		}
		credentialGetPasswordFunction, err = credentialStruct.InvokerNew("get_password")
	})
	return err
}

// GetPassword is a representation of the C type webkit_credential_get_password.
func (recv *Credential) GetPassword() (string, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := credentialGetPasswordFunction_Set()
	if err == nil {
		ret = credentialGetPasswordFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo, err
}

// UNSUPPORTED : C value 'webkit_credential_get_persistence' : return type 'CredentialPersistence' not supported

var credentialGetUsernameFunction *gi.Function
var credentialGetUsernameFunction_Once sync.Once

func credentialGetUsernameFunction_Set() error {
	var err error
	credentialGetUsernameFunction_Once.Do(func() {
		err = credentialStruct_Set()
		if err != nil {
			return
		}
		credentialGetUsernameFunction, err = credentialStruct.InvokerNew("get_username")
	})
	return err
}

// GetUsername is a representation of the C type webkit_credential_get_username.
func (recv *Credential) GetUsername() (string, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := credentialGetUsernameFunction_Set()
	if err == nil {
		ret = credentialGetUsernameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo, err
}

// UNSUPPORTED : C value 'webkit_credential_has_password' : return type 'gboolean' not supported

var deviceInfoPermissionRequestClassStruct *gi.Struct
var deviceInfoPermissionRequestClassStruct_Once sync.Once

func deviceInfoPermissionRequestClassStruct_Set() error {
	var err error
	deviceInfoPermissionRequestClassStruct_Once.Do(func() {
		deviceInfoPermissionRequestClassStruct, err = gi.StructNew("WebKit2", "DeviceInfoPermissionRequestClass")
	})
	return err
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
var deviceInfoPermissionRequestPrivateStruct_Once sync.Once

func deviceInfoPermissionRequestPrivateStruct_Set() error {
	var err error
	deviceInfoPermissionRequestPrivateStruct_Once.Do(func() {
		deviceInfoPermissionRequestPrivateStruct, err = gi.StructNew("WebKit2", "DeviceInfoPermissionRequestPrivate")
	})
	return err
}

type DeviceInfoPermissionRequestPrivate struct {
	native uintptr
}

var downloadClassStruct *gi.Struct
var downloadClassStruct_Once sync.Once

func downloadClassStruct_Set() error {
	var err error
	downloadClassStruct_Once.Do(func() {
		downloadClassStruct, err = gi.StructNew("WebKit2", "DownloadClass")
	})
	return err
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
var downloadPrivateStruct_Once sync.Once

func downloadPrivateStruct_Set() error {
	var err error
	downloadPrivateStruct_Once.Do(func() {
		downloadPrivateStruct, err = gi.StructNew("WebKit2", "DownloadPrivate")
	})
	return err
}

type DownloadPrivate struct {
	native uintptr
}

var editorStateClassStruct *gi.Struct
var editorStateClassStruct_Once sync.Once

func editorStateClassStruct_Set() error {
	var err error
	editorStateClassStruct_Once.Do(func() {
		editorStateClassStruct, err = gi.StructNew("WebKit2", "EditorStateClass")
	})
	return err
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
var editorStatePrivateStruct_Once sync.Once

func editorStatePrivateStruct_Set() error {
	var err error
	editorStatePrivateStruct_Once.Do(func() {
		editorStatePrivateStruct, err = gi.StructNew("WebKit2", "EditorStatePrivate")
	})
	return err
}

type EditorStatePrivate struct {
	native uintptr
}

var faviconDatabaseClassStruct *gi.Struct
var faviconDatabaseClassStruct_Once sync.Once

func faviconDatabaseClassStruct_Set() error {
	var err error
	faviconDatabaseClassStruct_Once.Do(func() {
		faviconDatabaseClassStruct, err = gi.StructNew("WebKit2", "FaviconDatabaseClass")
	})
	return err
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
var faviconDatabasePrivateStruct_Once sync.Once

func faviconDatabasePrivateStruct_Set() error {
	var err error
	faviconDatabasePrivateStruct_Once.Do(func() {
		faviconDatabasePrivateStruct, err = gi.StructNew("WebKit2", "FaviconDatabasePrivate")
	})
	return err
}

type FaviconDatabasePrivate struct {
	native uintptr
}

var fileChooserRequestClassStruct *gi.Struct
var fileChooserRequestClassStruct_Once sync.Once

func fileChooserRequestClassStruct_Set() error {
	var err error
	fileChooserRequestClassStruct_Once.Do(func() {
		fileChooserRequestClassStruct, err = gi.StructNew("WebKit2", "FileChooserRequestClass")
	})
	return err
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
var fileChooserRequestPrivateStruct_Once sync.Once

func fileChooserRequestPrivateStruct_Set() error {
	var err error
	fileChooserRequestPrivateStruct_Once.Do(func() {
		fileChooserRequestPrivateStruct, err = gi.StructNew("WebKit2", "FileChooserRequestPrivate")
	})
	return err
}

type FileChooserRequestPrivate struct {
	native uintptr
}

var findControllerClassStruct *gi.Struct
var findControllerClassStruct_Once sync.Once

func findControllerClassStruct_Set() error {
	var err error
	findControllerClassStruct_Once.Do(func() {
		findControllerClassStruct, err = gi.StructNew("WebKit2", "FindControllerClass")
	})
	return err
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
var findControllerPrivateStruct_Once sync.Once

func findControllerPrivateStruct_Set() error {
	var err error
	findControllerPrivateStruct_Once.Do(func() {
		findControllerPrivateStruct, err = gi.StructNew("WebKit2", "FindControllerPrivate")
	})
	return err
}

type FindControllerPrivate struct {
	native uintptr
}

var formSubmissionRequestClassStruct *gi.Struct
var formSubmissionRequestClassStruct_Once sync.Once

func formSubmissionRequestClassStruct_Set() error {
	var err error
	formSubmissionRequestClassStruct_Once.Do(func() {
		formSubmissionRequestClassStruct, err = gi.StructNew("WebKit2", "FormSubmissionRequestClass")
	})
	return err
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
var formSubmissionRequestPrivateStruct_Once sync.Once

func formSubmissionRequestPrivateStruct_Set() error {
	var err error
	formSubmissionRequestPrivateStruct_Once.Do(func() {
		formSubmissionRequestPrivateStruct, err = gi.StructNew("WebKit2", "FormSubmissionRequestPrivate")
	})
	return err
}

type FormSubmissionRequestPrivate struct {
	native uintptr
}

var geolocationManagerClassStruct *gi.Struct
var geolocationManagerClassStruct_Once sync.Once

func geolocationManagerClassStruct_Set() error {
	var err error
	geolocationManagerClassStruct_Once.Do(func() {
		geolocationManagerClassStruct, err = gi.StructNew("WebKit2", "GeolocationManagerClass")
	})
	return err
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
var geolocationManagerPrivateStruct_Once sync.Once

func geolocationManagerPrivateStruct_Set() error {
	var err error
	geolocationManagerPrivateStruct_Once.Do(func() {
		geolocationManagerPrivateStruct, err = gi.StructNew("WebKit2", "GeolocationManagerPrivate")
	})
	return err
}

type GeolocationManagerPrivate struct {
	native uintptr
}

var geolocationPermissionRequestClassStruct *gi.Struct
var geolocationPermissionRequestClassStruct_Once sync.Once

func geolocationPermissionRequestClassStruct_Set() error {
	var err error
	geolocationPermissionRequestClassStruct_Once.Do(func() {
		geolocationPermissionRequestClassStruct, err = gi.StructNew("WebKit2", "GeolocationPermissionRequestClass")
	})
	return err
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
var geolocationPermissionRequestPrivateStruct_Once sync.Once

func geolocationPermissionRequestPrivateStruct_Set() error {
	var err error
	geolocationPermissionRequestPrivateStruct_Once.Do(func() {
		geolocationPermissionRequestPrivateStruct, err = gi.StructNew("WebKit2", "GeolocationPermissionRequestPrivate")
	})
	return err
}

type GeolocationPermissionRequestPrivate struct {
	native uintptr
}

var geolocationPositionStruct *gi.Struct
var geolocationPositionStruct_Once sync.Once

func geolocationPositionStruct_Set() error {
	var err error
	geolocationPositionStruct_Once.Do(func() {
		geolocationPositionStruct, err = gi.StructNew("WebKit2", "GeolocationPosition")
	})
	return err
}

type GeolocationPosition struct {
	native uintptr
}

// UNSUPPORTED : C value 'webkit_geolocation_position_new' : parameter 'latitude' of type 'gdouble' not supported

var geolocationPositionCopyFunction *gi.Function
var geolocationPositionCopyFunction_Once sync.Once

func geolocationPositionCopyFunction_Set() error {
	var err error
	geolocationPositionCopyFunction_Once.Do(func() {
		err = geolocationPositionStruct_Set()
		if err != nil {
			return
		}
		geolocationPositionCopyFunction, err = geolocationPositionStruct.InvokerNew("copy")
	})
	return err
}

// Copy is a representation of the C type webkit_geolocation_position_copy.
func (recv *GeolocationPosition) Copy() (*GeolocationPosition, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := geolocationPositionCopyFunction_Set()
	if err == nil {
		ret = geolocationPositionCopyFunction.Invoke(inArgs[:], nil)
	}

	retGo := &GeolocationPosition{native: ret.Pointer()}

	return retGo, err
}

var geolocationPositionFreeFunction *gi.Function
var geolocationPositionFreeFunction_Once sync.Once

func geolocationPositionFreeFunction_Set() error {
	var err error
	geolocationPositionFreeFunction_Once.Do(func() {
		err = geolocationPositionStruct_Set()
		if err != nil {
			return
		}
		geolocationPositionFreeFunction, err = geolocationPositionStruct.InvokerNew("free")
	})
	return err
}

// Free is a representation of the C type webkit_geolocation_position_free.
func (recv *GeolocationPosition) Free() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := geolocationPositionFreeFunction_Set()
	if err == nil {
		geolocationPositionFreeFunction.Invoke(inArgs[:], nil)
	}

	return err
}

// UNSUPPORTED : C value 'webkit_geolocation_position_set_altitude' : parameter 'altitude' of type 'gdouble' not supported

// UNSUPPORTED : C value 'webkit_geolocation_position_set_altitude_accuracy' : parameter 'altitude_accuracy' of type 'gdouble' not supported

// UNSUPPORTED : C value 'webkit_geolocation_position_set_heading' : parameter 'heading' of type 'gdouble' not supported

// UNSUPPORTED : C value 'webkit_geolocation_position_set_speed' : parameter 'speed' of type 'gdouble' not supported

var geolocationPositionSetTimestampFunction *gi.Function
var geolocationPositionSetTimestampFunction_Once sync.Once

func geolocationPositionSetTimestampFunction_Set() error {
	var err error
	geolocationPositionSetTimestampFunction_Once.Do(func() {
		err = geolocationPositionStruct_Set()
		if err != nil {
			return
		}
		geolocationPositionSetTimestampFunction, err = geolocationPositionStruct.InvokerNew("set_timestamp")
	})
	return err
}

// SetTimestamp is a representation of the C type webkit_geolocation_position_set_timestamp.
func (recv *GeolocationPosition) SetTimestamp(timestamp uint64) error {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetUint64(timestamp)

	err := geolocationPositionSetTimestampFunction_Set()
	if err == nil {
		geolocationPositionSetTimestampFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var hitTestResultClassStruct *gi.Struct
var hitTestResultClassStruct_Once sync.Once

func hitTestResultClassStruct_Set() error {
	var err error
	hitTestResultClassStruct_Once.Do(func() {
		hitTestResultClassStruct, err = gi.StructNew("WebKit2", "HitTestResultClass")
	})
	return err
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
var hitTestResultPrivateStruct_Once sync.Once

func hitTestResultPrivateStruct_Set() error {
	var err error
	hitTestResultPrivateStruct_Once.Do(func() {
		hitTestResultPrivateStruct, err = gi.StructNew("WebKit2", "HitTestResultPrivate")
	})
	return err
}

type HitTestResultPrivate struct {
	native uintptr
}

var installMissingMediaPluginsPermissionRequestClassStruct *gi.Struct
var installMissingMediaPluginsPermissionRequestClassStruct_Once sync.Once

func installMissingMediaPluginsPermissionRequestClassStruct_Set() error {
	var err error
	installMissingMediaPluginsPermissionRequestClassStruct_Once.Do(func() {
		installMissingMediaPluginsPermissionRequestClassStruct, err = gi.StructNew("WebKit2", "InstallMissingMediaPluginsPermissionRequestClass")
	})
	return err
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
var installMissingMediaPluginsPermissionRequestPrivateStruct_Once sync.Once

func installMissingMediaPluginsPermissionRequestPrivateStruct_Set() error {
	var err error
	installMissingMediaPluginsPermissionRequestPrivateStruct_Once.Do(func() {
		installMissingMediaPluginsPermissionRequestPrivateStruct, err = gi.StructNew("WebKit2", "InstallMissingMediaPluginsPermissionRequestPrivate")
	})
	return err
}

type InstallMissingMediaPluginsPermissionRequestPrivate struct {
	native uintptr
}

var javascriptResultStruct *gi.Struct
var javascriptResultStruct_Once sync.Once

func javascriptResultStruct_Set() error {
	var err error
	javascriptResultStruct_Once.Do(func() {
		javascriptResultStruct, err = gi.StructNew("WebKit2", "JavascriptResult")
	})
	return err
}

type JavascriptResult struct {
	native uintptr
}

// UNSUPPORTED : C value 'webkit_javascript_result_get_global_context' : return type 'JavaScriptCore.GlobalContextRef' not supported

// UNSUPPORTED : C value 'webkit_javascript_result_get_js_value' : return type 'JavaScriptCore.Value' not supported

// UNSUPPORTED : C value 'webkit_javascript_result_get_value' : return type 'JavaScriptCore.ValueRef' not supported

var javascriptResultRefFunction *gi.Function
var javascriptResultRefFunction_Once sync.Once

func javascriptResultRefFunction_Set() error {
	var err error
	javascriptResultRefFunction_Once.Do(func() {
		err = javascriptResultStruct_Set()
		if err != nil {
			return
		}
		javascriptResultRefFunction, err = javascriptResultStruct.InvokerNew("ref")
	})
	return err
}

// Ref is a representation of the C type webkit_javascript_result_ref.
func (recv *JavascriptResult) Ref() (*JavascriptResult, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := javascriptResultRefFunction_Set()
	if err == nil {
		ret = javascriptResultRefFunction.Invoke(inArgs[:], nil)
	}

	retGo := &JavascriptResult{native: ret.Pointer()}

	return retGo, err
}

var javascriptResultUnrefFunction *gi.Function
var javascriptResultUnrefFunction_Once sync.Once

func javascriptResultUnrefFunction_Set() error {
	var err error
	javascriptResultUnrefFunction_Once.Do(func() {
		err = javascriptResultStruct_Set()
		if err != nil {
			return
		}
		javascriptResultUnrefFunction, err = javascriptResultStruct.InvokerNew("unref")
	})
	return err
}

// Unref is a representation of the C type webkit_javascript_result_unref.
func (recv *JavascriptResult) Unref() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := javascriptResultUnrefFunction_Set()
	if err == nil {
		javascriptResultUnrefFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var mimeInfoStruct *gi.Struct
var mimeInfoStruct_Once sync.Once

func mimeInfoStruct_Set() error {
	var err error
	mimeInfoStruct_Once.Do(func() {
		mimeInfoStruct, err = gi.StructNew("WebKit2", "MimeInfo")
	})
	return err
}

type MimeInfo struct {
	native uintptr
}

var mimeInfoGetDescriptionFunction *gi.Function
var mimeInfoGetDescriptionFunction_Once sync.Once

func mimeInfoGetDescriptionFunction_Set() error {
	var err error
	mimeInfoGetDescriptionFunction_Once.Do(func() {
		err = mimeInfoStruct_Set()
		if err != nil {
			return
		}
		mimeInfoGetDescriptionFunction, err = mimeInfoStruct.InvokerNew("get_description")
	})
	return err
}

// GetDescription is a representation of the C type webkit_mime_info_get_description.
func (recv *MimeInfo) GetDescription() (string, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := mimeInfoGetDescriptionFunction_Set()
	if err == nil {
		ret = mimeInfoGetDescriptionFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo, err
}

var mimeInfoGetExtensionsFunction *gi.Function
var mimeInfoGetExtensionsFunction_Once sync.Once

func mimeInfoGetExtensionsFunction_Set() error {
	var err error
	mimeInfoGetExtensionsFunction_Once.Do(func() {
		err = mimeInfoStruct_Set()
		if err != nil {
			return
		}
		mimeInfoGetExtensionsFunction, err = mimeInfoStruct.InvokerNew("get_extensions")
	})
	return err
}

// GetExtensions is a representation of the C type webkit_mime_info_get_extensions.
func (recv *MimeInfo) GetExtensions() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := mimeInfoGetExtensionsFunction_Set()
	if err == nil {
		mimeInfoGetExtensionsFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var mimeInfoGetMimeTypeFunction *gi.Function
var mimeInfoGetMimeTypeFunction_Once sync.Once

func mimeInfoGetMimeTypeFunction_Set() error {
	var err error
	mimeInfoGetMimeTypeFunction_Once.Do(func() {
		err = mimeInfoStruct_Set()
		if err != nil {
			return
		}
		mimeInfoGetMimeTypeFunction, err = mimeInfoStruct.InvokerNew("get_mime_type")
	})
	return err
}

// GetMimeType is a representation of the C type webkit_mime_info_get_mime_type.
func (recv *MimeInfo) GetMimeType() (string, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := mimeInfoGetMimeTypeFunction_Set()
	if err == nil {
		ret = mimeInfoGetMimeTypeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo, err
}

var mimeInfoRefFunction *gi.Function
var mimeInfoRefFunction_Once sync.Once

func mimeInfoRefFunction_Set() error {
	var err error
	mimeInfoRefFunction_Once.Do(func() {
		err = mimeInfoStruct_Set()
		if err != nil {
			return
		}
		mimeInfoRefFunction, err = mimeInfoStruct.InvokerNew("ref")
	})
	return err
}

// Ref is a representation of the C type webkit_mime_info_ref.
func (recv *MimeInfo) Ref() (*MimeInfo, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := mimeInfoRefFunction_Set()
	if err == nil {
		ret = mimeInfoRefFunction.Invoke(inArgs[:], nil)
	}

	retGo := &MimeInfo{native: ret.Pointer()}

	return retGo, err
}

var mimeInfoUnrefFunction *gi.Function
var mimeInfoUnrefFunction_Once sync.Once

func mimeInfoUnrefFunction_Set() error {
	var err error
	mimeInfoUnrefFunction_Once.Do(func() {
		err = mimeInfoStruct_Set()
		if err != nil {
			return
		}
		mimeInfoUnrefFunction, err = mimeInfoStruct.InvokerNew("unref")
	})
	return err
}

// Unref is a representation of the C type webkit_mime_info_unref.
func (recv *MimeInfo) Unref() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := mimeInfoUnrefFunction_Set()
	if err == nil {
		mimeInfoUnrefFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var navigationActionStruct *gi.Struct
var navigationActionStruct_Once sync.Once

func navigationActionStruct_Set() error {
	var err error
	navigationActionStruct_Once.Do(func() {
		navigationActionStruct, err = gi.StructNew("WebKit2", "NavigationAction")
	})
	return err
}

type NavigationAction struct {
	native uintptr
}

var navigationActionCopyFunction *gi.Function
var navigationActionCopyFunction_Once sync.Once

func navigationActionCopyFunction_Set() error {
	var err error
	navigationActionCopyFunction_Once.Do(func() {
		err = navigationActionStruct_Set()
		if err != nil {
			return
		}
		navigationActionCopyFunction, err = navigationActionStruct.InvokerNew("copy")
	})
	return err
}

// Copy is a representation of the C type webkit_navigation_action_copy.
func (recv *NavigationAction) Copy() (*NavigationAction, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := navigationActionCopyFunction_Set()
	if err == nil {
		ret = navigationActionCopyFunction.Invoke(inArgs[:], nil)
	}

	retGo := &NavigationAction{native: ret.Pointer()}

	return retGo, err
}

var navigationActionFreeFunction *gi.Function
var navigationActionFreeFunction_Once sync.Once

func navigationActionFreeFunction_Set() error {
	var err error
	navigationActionFreeFunction_Once.Do(func() {
		err = navigationActionStruct_Set()
		if err != nil {
			return
		}
		navigationActionFreeFunction, err = navigationActionStruct.InvokerNew("free")
	})
	return err
}

// Free is a representation of the C type webkit_navigation_action_free.
func (recv *NavigationAction) Free() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := navigationActionFreeFunction_Set()
	if err == nil {
		navigationActionFreeFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var navigationActionGetModifiersFunction *gi.Function
var navigationActionGetModifiersFunction_Once sync.Once

func navigationActionGetModifiersFunction_Set() error {
	var err error
	navigationActionGetModifiersFunction_Once.Do(func() {
		err = navigationActionStruct_Set()
		if err != nil {
			return
		}
		navigationActionGetModifiersFunction, err = navigationActionStruct.InvokerNew("get_modifiers")
	})
	return err
}

// GetModifiers is a representation of the C type webkit_navigation_action_get_modifiers.
func (recv *NavigationAction) GetModifiers() (uint32, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := navigationActionGetModifiersFunction_Set()
	if err == nil {
		ret = navigationActionGetModifiersFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo, err
}

var navigationActionGetMouseButtonFunction *gi.Function
var navigationActionGetMouseButtonFunction_Once sync.Once

func navigationActionGetMouseButtonFunction_Set() error {
	var err error
	navigationActionGetMouseButtonFunction_Once.Do(func() {
		err = navigationActionStruct_Set()
		if err != nil {
			return
		}
		navigationActionGetMouseButtonFunction, err = navigationActionStruct.InvokerNew("get_mouse_button")
	})
	return err
}

// GetMouseButton is a representation of the C type webkit_navigation_action_get_mouse_button.
func (recv *NavigationAction) GetMouseButton() (uint32, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := navigationActionGetMouseButtonFunction_Set()
	if err == nil {
		ret = navigationActionGetMouseButtonFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo, err
}

// UNSUPPORTED : C value 'webkit_navigation_action_get_navigation_type' : return type 'NavigationType' not supported

// UNSUPPORTED : C value 'webkit_navigation_action_get_request' : return type 'URIRequest' not supported

// UNSUPPORTED : C value 'webkit_navigation_action_is_redirect' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'webkit_navigation_action_is_user_gesture' : return type 'gboolean' not supported

var navigationPolicyDecisionClassStruct *gi.Struct
var navigationPolicyDecisionClassStruct_Once sync.Once

func navigationPolicyDecisionClassStruct_Set() error {
	var err error
	navigationPolicyDecisionClassStruct_Once.Do(func() {
		navigationPolicyDecisionClassStruct, err = gi.StructNew("WebKit2", "NavigationPolicyDecisionClass")
	})
	return err
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
var navigationPolicyDecisionPrivateStruct_Once sync.Once

func navigationPolicyDecisionPrivateStruct_Set() error {
	var err error
	navigationPolicyDecisionPrivateStruct_Once.Do(func() {
		navigationPolicyDecisionPrivateStruct, err = gi.StructNew("WebKit2", "NavigationPolicyDecisionPrivate")
	})
	return err
}

type NavigationPolicyDecisionPrivate struct {
	native uintptr
}

var networkProxySettingsStruct *gi.Struct
var networkProxySettingsStruct_Once sync.Once

func networkProxySettingsStruct_Set() error {
	var err error
	networkProxySettingsStruct_Once.Do(func() {
		networkProxySettingsStruct, err = gi.StructNew("WebKit2", "NetworkProxySettings")
	})
	return err
}

type NetworkProxySettings struct {
	native uintptr
}

// UNSUPPORTED : C value 'webkit_network_proxy_settings_new' : parameter 'ignore_hosts' has no type

var networkProxySettingsAddProxyForSchemeFunction *gi.Function
var networkProxySettingsAddProxyForSchemeFunction_Once sync.Once

func networkProxySettingsAddProxyForSchemeFunction_Set() error {
	var err error
	networkProxySettingsAddProxyForSchemeFunction_Once.Do(func() {
		err = networkProxySettingsStruct_Set()
		if err != nil {
			return
		}
		networkProxySettingsAddProxyForSchemeFunction, err = networkProxySettingsStruct.InvokerNew("add_proxy_for_scheme")
	})
	return err
}

// AddProxyForScheme is a representation of the C type webkit_network_proxy_settings_add_proxy_for_scheme.
func (recv *NetworkProxySettings) AddProxyForScheme(scheme string, proxyUri string) error {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(scheme)
	inArgs[2].SetString(proxyUri)

	err := networkProxySettingsAddProxyForSchemeFunction_Set()
	if err == nil {
		networkProxySettingsAddProxyForSchemeFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var networkProxySettingsCopyFunction *gi.Function
var networkProxySettingsCopyFunction_Once sync.Once

func networkProxySettingsCopyFunction_Set() error {
	var err error
	networkProxySettingsCopyFunction_Once.Do(func() {
		err = networkProxySettingsStruct_Set()
		if err != nil {
			return
		}
		networkProxySettingsCopyFunction, err = networkProxySettingsStruct.InvokerNew("copy")
	})
	return err
}

// Copy is a representation of the C type webkit_network_proxy_settings_copy.
func (recv *NetworkProxySettings) Copy() (*NetworkProxySettings, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := networkProxySettingsCopyFunction_Set()
	if err == nil {
		ret = networkProxySettingsCopyFunction.Invoke(inArgs[:], nil)
	}

	retGo := &NetworkProxySettings{native: ret.Pointer()}

	return retGo, err
}

var networkProxySettingsFreeFunction *gi.Function
var networkProxySettingsFreeFunction_Once sync.Once

func networkProxySettingsFreeFunction_Set() error {
	var err error
	networkProxySettingsFreeFunction_Once.Do(func() {
		err = networkProxySettingsStruct_Set()
		if err != nil {
			return
		}
		networkProxySettingsFreeFunction, err = networkProxySettingsStruct.InvokerNew("free")
	})
	return err
}

// Free is a representation of the C type webkit_network_proxy_settings_free.
func (recv *NetworkProxySettings) Free() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := networkProxySettingsFreeFunction_Set()
	if err == nil {
		networkProxySettingsFreeFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var notificationClassStruct *gi.Struct
var notificationClassStruct_Once sync.Once

func notificationClassStruct_Set() error {
	var err error
	notificationClassStruct_Once.Do(func() {
		notificationClassStruct, err = gi.StructNew("WebKit2", "NotificationClass")
	})
	return err
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
var notificationPermissionRequestClassStruct_Once sync.Once

func notificationPermissionRequestClassStruct_Set() error {
	var err error
	notificationPermissionRequestClassStruct_Once.Do(func() {
		notificationPermissionRequestClassStruct, err = gi.StructNew("WebKit2", "NotificationPermissionRequestClass")
	})
	return err
}

type NotificationPermissionRequestClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
}

var notificationPermissionRequestPrivateStruct *gi.Struct
var notificationPermissionRequestPrivateStruct_Once sync.Once

func notificationPermissionRequestPrivateStruct_Set() error {
	var err error
	notificationPermissionRequestPrivateStruct_Once.Do(func() {
		notificationPermissionRequestPrivateStruct, err = gi.StructNew("WebKit2", "NotificationPermissionRequestPrivate")
	})
	return err
}

type NotificationPermissionRequestPrivate struct {
	native uintptr
}

var notificationPrivateStruct *gi.Struct
var notificationPrivateStruct_Once sync.Once

func notificationPrivateStruct_Set() error {
	var err error
	notificationPrivateStruct_Once.Do(func() {
		notificationPrivateStruct, err = gi.StructNew("WebKit2", "NotificationPrivate")
	})
	return err
}

type NotificationPrivate struct {
	native uintptr
}

var optionMenuClassStruct *gi.Struct
var optionMenuClassStruct_Once sync.Once

func optionMenuClassStruct_Set() error {
	var err error
	optionMenuClassStruct_Once.Do(func() {
		optionMenuClassStruct, err = gi.StructNew("WebKit2", "OptionMenuClass")
	})
	return err
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
var optionMenuItemStruct_Once sync.Once

func optionMenuItemStruct_Set() error {
	var err error
	optionMenuItemStruct_Once.Do(func() {
		optionMenuItemStruct, err = gi.StructNew("WebKit2", "OptionMenuItem")
	})
	return err
}

type OptionMenuItem struct {
	native uintptr
}

var optionMenuItemCopyFunction *gi.Function
var optionMenuItemCopyFunction_Once sync.Once

func optionMenuItemCopyFunction_Set() error {
	var err error
	optionMenuItemCopyFunction_Once.Do(func() {
		err = optionMenuItemStruct_Set()
		if err != nil {
			return
		}
		optionMenuItemCopyFunction, err = optionMenuItemStruct.InvokerNew("copy")
	})
	return err
}

// Copy is a representation of the C type webkit_option_menu_item_copy.
func (recv *OptionMenuItem) Copy() (*OptionMenuItem, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := optionMenuItemCopyFunction_Set()
	if err == nil {
		ret = optionMenuItemCopyFunction.Invoke(inArgs[:], nil)
	}

	retGo := &OptionMenuItem{native: ret.Pointer()}

	return retGo, err
}

var optionMenuItemFreeFunction *gi.Function
var optionMenuItemFreeFunction_Once sync.Once

func optionMenuItemFreeFunction_Set() error {
	var err error
	optionMenuItemFreeFunction_Once.Do(func() {
		err = optionMenuItemStruct_Set()
		if err != nil {
			return
		}
		optionMenuItemFreeFunction, err = optionMenuItemStruct.InvokerNew("free")
	})
	return err
}

// Free is a representation of the C type webkit_option_menu_item_free.
func (recv *OptionMenuItem) Free() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := optionMenuItemFreeFunction_Set()
	if err == nil {
		optionMenuItemFreeFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var optionMenuItemGetLabelFunction *gi.Function
var optionMenuItemGetLabelFunction_Once sync.Once

func optionMenuItemGetLabelFunction_Set() error {
	var err error
	optionMenuItemGetLabelFunction_Once.Do(func() {
		err = optionMenuItemStruct_Set()
		if err != nil {
			return
		}
		optionMenuItemGetLabelFunction, err = optionMenuItemStruct.InvokerNew("get_label")
	})
	return err
}

// GetLabel is a representation of the C type webkit_option_menu_item_get_label.
func (recv *OptionMenuItem) GetLabel() (string, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := optionMenuItemGetLabelFunction_Set()
	if err == nil {
		ret = optionMenuItemGetLabelFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo, err
}

var optionMenuItemGetTooltipFunction *gi.Function
var optionMenuItemGetTooltipFunction_Once sync.Once

func optionMenuItemGetTooltipFunction_Set() error {
	var err error
	optionMenuItemGetTooltipFunction_Once.Do(func() {
		err = optionMenuItemStruct_Set()
		if err != nil {
			return
		}
		optionMenuItemGetTooltipFunction, err = optionMenuItemStruct.InvokerNew("get_tooltip")
	})
	return err
}

// GetTooltip is a representation of the C type webkit_option_menu_item_get_tooltip.
func (recv *OptionMenuItem) GetTooltip() (string, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := optionMenuItemGetTooltipFunction_Set()
	if err == nil {
		ret = optionMenuItemGetTooltipFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo, err
}

// UNSUPPORTED : C value 'webkit_option_menu_item_is_enabled' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'webkit_option_menu_item_is_group_child' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'webkit_option_menu_item_is_group_label' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'webkit_option_menu_item_is_selected' : return type 'gboolean' not supported

var optionMenuPrivateStruct *gi.Struct
var optionMenuPrivateStruct_Once sync.Once

func optionMenuPrivateStruct_Set() error {
	var err error
	optionMenuPrivateStruct_Once.Do(func() {
		optionMenuPrivateStruct, err = gi.StructNew("WebKit2", "OptionMenuPrivate")
	})
	return err
}

type OptionMenuPrivate struct {
	native uintptr
}

var permissionRequestIfaceStruct *gi.Struct
var permissionRequestIfaceStruct_Once sync.Once

func permissionRequestIfaceStruct_Set() error {
	var err error
	permissionRequestIfaceStruct_Once.Do(func() {
		permissionRequestIfaceStruct, err = gi.StructNew("WebKit2", "PermissionRequestIface")
	})
	return err
}

type PermissionRequestIface struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_interface' : no Go type for 'GObject.TypeInterface'
	// UNSUPPORTED : C value 'allow' : missing Type
	// UNSUPPORTED : C value 'deny' : missing Type
}

var pluginClassStruct *gi.Struct
var pluginClassStruct_Once sync.Once

func pluginClassStruct_Set() error {
	var err error
	pluginClassStruct_Once.Do(func() {
		pluginClassStruct, err = gi.StructNew("WebKit2", "PluginClass")
	})
	return err
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
var pluginPrivateStruct_Once sync.Once

func pluginPrivateStruct_Set() error {
	var err error
	pluginPrivateStruct_Once.Do(func() {
		pluginPrivateStruct, err = gi.StructNew("WebKit2", "PluginPrivate")
	})
	return err
}

type PluginPrivate struct {
	native uintptr
}

var policyDecisionClassStruct *gi.Struct
var policyDecisionClassStruct_Once sync.Once

func policyDecisionClassStruct_Set() error {
	var err error
	policyDecisionClassStruct_Once.Do(func() {
		policyDecisionClassStruct, err = gi.StructNew("WebKit2", "PolicyDecisionClass")
	})
	return err
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
var policyDecisionPrivateStruct_Once sync.Once

func policyDecisionPrivateStruct_Set() error {
	var err error
	policyDecisionPrivateStruct_Once.Do(func() {
		policyDecisionPrivateStruct, err = gi.StructNew("WebKit2", "PolicyDecisionPrivate")
	})
	return err
}

type PolicyDecisionPrivate struct {
	native uintptr
}

var printCustomWidgetClassStruct *gi.Struct
var printCustomWidgetClassStruct_Once sync.Once

func printCustomWidgetClassStruct_Set() error {
	var err error
	printCustomWidgetClassStruct_Once.Do(func() {
		printCustomWidgetClassStruct, err = gi.StructNew("WebKit2", "PrintCustomWidgetClass")
	})
	return err
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
var printCustomWidgetPrivateStruct_Once sync.Once

func printCustomWidgetPrivateStruct_Set() error {
	var err error
	printCustomWidgetPrivateStruct_Once.Do(func() {
		printCustomWidgetPrivateStruct, err = gi.StructNew("WebKit2", "PrintCustomWidgetPrivate")
	})
	return err
}

type PrintCustomWidgetPrivate struct {
	native uintptr
}

var printOperationClassStruct *gi.Struct
var printOperationClassStruct_Once sync.Once

func printOperationClassStruct_Set() error {
	var err error
	printOperationClassStruct_Once.Do(func() {
		printOperationClassStruct, err = gi.StructNew("WebKit2", "PrintOperationClass")
	})
	return err
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
var printOperationPrivateStruct_Once sync.Once

func printOperationPrivateStruct_Set() error {
	var err error
	printOperationPrivateStruct_Once.Do(func() {
		printOperationPrivateStruct, err = gi.StructNew("WebKit2", "PrintOperationPrivate")
	})
	return err
}

type PrintOperationPrivate struct {
	native uintptr
}

var responsePolicyDecisionClassStruct *gi.Struct
var responsePolicyDecisionClassStruct_Once sync.Once

func responsePolicyDecisionClassStruct_Set() error {
	var err error
	responsePolicyDecisionClassStruct_Once.Do(func() {
		responsePolicyDecisionClassStruct, err = gi.StructNew("WebKit2", "ResponsePolicyDecisionClass")
	})
	return err
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
var responsePolicyDecisionPrivateStruct_Once sync.Once

func responsePolicyDecisionPrivateStruct_Set() error {
	var err error
	responsePolicyDecisionPrivateStruct_Once.Do(func() {
		responsePolicyDecisionPrivateStruct, err = gi.StructNew("WebKit2", "ResponsePolicyDecisionPrivate")
	})
	return err
}

type ResponsePolicyDecisionPrivate struct {
	native uintptr
}

var scriptDialogStruct *gi.Struct
var scriptDialogStruct_Once sync.Once

func scriptDialogStruct_Set() error {
	var err error
	scriptDialogStruct_Once.Do(func() {
		scriptDialogStruct, err = gi.StructNew("WebKit2", "ScriptDialog")
	})
	return err
}

type ScriptDialog struct {
	native uintptr
}

var scriptDialogCloseFunction *gi.Function
var scriptDialogCloseFunction_Once sync.Once

func scriptDialogCloseFunction_Set() error {
	var err error
	scriptDialogCloseFunction_Once.Do(func() {
		err = scriptDialogStruct_Set()
		if err != nil {
			return
		}
		scriptDialogCloseFunction, err = scriptDialogStruct.InvokerNew("close")
	})
	return err
}

// Close is a representation of the C type webkit_script_dialog_close.
func (recv *ScriptDialog) Close() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := scriptDialogCloseFunction_Set()
	if err == nil {
		scriptDialogCloseFunction.Invoke(inArgs[:], nil)
	}

	return err
}

// UNSUPPORTED : C value 'webkit_script_dialog_confirm_set_confirmed' : parameter 'confirmed' of type 'gboolean' not supported

// UNSUPPORTED : C value 'webkit_script_dialog_get_dialog_type' : return type 'ScriptDialogType' not supported

var scriptDialogGetMessageFunction *gi.Function
var scriptDialogGetMessageFunction_Once sync.Once

func scriptDialogGetMessageFunction_Set() error {
	var err error
	scriptDialogGetMessageFunction_Once.Do(func() {
		err = scriptDialogStruct_Set()
		if err != nil {
			return
		}
		scriptDialogGetMessageFunction, err = scriptDialogStruct.InvokerNew("get_message")
	})
	return err
}

// GetMessage is a representation of the C type webkit_script_dialog_get_message.
func (recv *ScriptDialog) GetMessage() (string, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := scriptDialogGetMessageFunction_Set()
	if err == nil {
		ret = scriptDialogGetMessageFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo, err
}

var scriptDialogPromptGetDefaultTextFunction *gi.Function
var scriptDialogPromptGetDefaultTextFunction_Once sync.Once

func scriptDialogPromptGetDefaultTextFunction_Set() error {
	var err error
	scriptDialogPromptGetDefaultTextFunction_Once.Do(func() {
		err = scriptDialogStruct_Set()
		if err != nil {
			return
		}
		scriptDialogPromptGetDefaultTextFunction, err = scriptDialogStruct.InvokerNew("prompt_get_default_text")
	})
	return err
}

// PromptGetDefaultText is a representation of the C type webkit_script_dialog_prompt_get_default_text.
func (recv *ScriptDialog) PromptGetDefaultText() (string, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := scriptDialogPromptGetDefaultTextFunction_Set()
	if err == nil {
		ret = scriptDialogPromptGetDefaultTextFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo, err
}

var scriptDialogPromptSetTextFunction *gi.Function
var scriptDialogPromptSetTextFunction_Once sync.Once

func scriptDialogPromptSetTextFunction_Set() error {
	var err error
	scriptDialogPromptSetTextFunction_Once.Do(func() {
		err = scriptDialogStruct_Set()
		if err != nil {
			return
		}
		scriptDialogPromptSetTextFunction, err = scriptDialogStruct.InvokerNew("prompt_set_text")
	})
	return err
}

// PromptSetText is a representation of the C type webkit_script_dialog_prompt_set_text.
func (recv *ScriptDialog) PromptSetText(text string) error {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(text)

	err := scriptDialogPromptSetTextFunction_Set()
	if err == nil {
		scriptDialogPromptSetTextFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var scriptDialogRefFunction *gi.Function
var scriptDialogRefFunction_Once sync.Once

func scriptDialogRefFunction_Set() error {
	var err error
	scriptDialogRefFunction_Once.Do(func() {
		err = scriptDialogStruct_Set()
		if err != nil {
			return
		}
		scriptDialogRefFunction, err = scriptDialogStruct.InvokerNew("ref")
	})
	return err
}

// Ref is a representation of the C type webkit_script_dialog_ref.
func (recv *ScriptDialog) Ref() (*ScriptDialog, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := scriptDialogRefFunction_Set()
	if err == nil {
		ret = scriptDialogRefFunction.Invoke(inArgs[:], nil)
	}

	retGo := &ScriptDialog{native: ret.Pointer()}

	return retGo, err
}

var scriptDialogUnrefFunction *gi.Function
var scriptDialogUnrefFunction_Once sync.Once

func scriptDialogUnrefFunction_Set() error {
	var err error
	scriptDialogUnrefFunction_Once.Do(func() {
		err = scriptDialogStruct_Set()
		if err != nil {
			return
		}
		scriptDialogUnrefFunction, err = scriptDialogStruct.InvokerNew("unref")
	})
	return err
}

// Unref is a representation of the C type webkit_script_dialog_unref.
func (recv *ScriptDialog) Unref() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := scriptDialogUnrefFunction_Set()
	if err == nil {
		scriptDialogUnrefFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var securityManagerClassStruct *gi.Struct
var securityManagerClassStruct_Once sync.Once

func securityManagerClassStruct_Set() error {
	var err error
	securityManagerClassStruct_Once.Do(func() {
		securityManagerClassStruct, err = gi.StructNew("WebKit2", "SecurityManagerClass")
	})
	return err
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
var securityManagerPrivateStruct_Once sync.Once

func securityManagerPrivateStruct_Set() error {
	var err error
	securityManagerPrivateStruct_Once.Do(func() {
		securityManagerPrivateStruct, err = gi.StructNew("WebKit2", "SecurityManagerPrivate")
	})
	return err
}

type SecurityManagerPrivate struct {
	native uintptr
}

var securityOriginStruct *gi.Struct
var securityOriginStruct_Once sync.Once

func securityOriginStruct_Set() error {
	var err error
	securityOriginStruct_Once.Do(func() {
		securityOriginStruct, err = gi.StructNew("WebKit2", "SecurityOrigin")
	})
	return err
}

type SecurityOrigin struct {
	native uintptr
}

var securityOriginNewFunction *gi.Function
var securityOriginNewFunction_Once sync.Once

func securityOriginNewFunction_Set() error {
	var err error
	securityOriginNewFunction_Once.Do(func() {
		err = securityOriginStruct_Set()
		if err != nil {
			return
		}
		securityOriginNewFunction, err = securityOriginStruct.InvokerNew("new")
	})
	return err
}

// SecurityOriginNew is a representation of the C type webkit_security_origin_new.
func SecurityOriginNew(protocol string, host string, port uint16) (*SecurityOrigin, error) {
	var inArgs [3]gi.Argument
	inArgs[0].SetString(protocol)
	inArgs[1].SetString(host)
	inArgs[2].SetUint16(port)

	var ret gi.Argument

	err := securityOriginNewFunction_Set()
	if err == nil {
		ret = securityOriginNewFunction.Invoke(inArgs[:], nil)
	}

	retGo := &SecurityOrigin{native: ret.Pointer()}

	return retGo, err
}

var securityOriginNewForUriFunction *gi.Function
var securityOriginNewForUriFunction_Once sync.Once

func securityOriginNewForUriFunction_Set() error {
	var err error
	securityOriginNewForUriFunction_Once.Do(func() {
		err = securityOriginStruct_Set()
		if err != nil {
			return
		}
		securityOriginNewForUriFunction, err = securityOriginStruct.InvokerNew("new_for_uri")
	})
	return err
}

// SecurityOriginNewForUri is a representation of the C type webkit_security_origin_new_for_uri.
func SecurityOriginNewForUri(uri string) (*SecurityOrigin, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(uri)

	var ret gi.Argument

	err := securityOriginNewForUriFunction_Set()
	if err == nil {
		ret = securityOriginNewForUriFunction.Invoke(inArgs[:], nil)
	}

	retGo := &SecurityOrigin{native: ret.Pointer()}

	return retGo, err
}

var securityOriginGetHostFunction *gi.Function
var securityOriginGetHostFunction_Once sync.Once

func securityOriginGetHostFunction_Set() error {
	var err error
	securityOriginGetHostFunction_Once.Do(func() {
		err = securityOriginStruct_Set()
		if err != nil {
			return
		}
		securityOriginGetHostFunction, err = securityOriginStruct.InvokerNew("get_host")
	})
	return err
}

// GetHost is a representation of the C type webkit_security_origin_get_host.
func (recv *SecurityOrigin) GetHost() (string, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := securityOriginGetHostFunction_Set()
	if err == nil {
		ret = securityOriginGetHostFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo, err
}

var securityOriginGetPortFunction *gi.Function
var securityOriginGetPortFunction_Once sync.Once

func securityOriginGetPortFunction_Set() error {
	var err error
	securityOriginGetPortFunction_Once.Do(func() {
		err = securityOriginStruct_Set()
		if err != nil {
			return
		}
		securityOriginGetPortFunction, err = securityOriginStruct.InvokerNew("get_port")
	})
	return err
}

// GetPort is a representation of the C type webkit_security_origin_get_port.
func (recv *SecurityOrigin) GetPort() (uint16, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := securityOriginGetPortFunction_Set()
	if err == nil {
		ret = securityOriginGetPortFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint16()

	return retGo, err
}

var securityOriginGetProtocolFunction *gi.Function
var securityOriginGetProtocolFunction_Once sync.Once

func securityOriginGetProtocolFunction_Set() error {
	var err error
	securityOriginGetProtocolFunction_Once.Do(func() {
		err = securityOriginStruct_Set()
		if err != nil {
			return
		}
		securityOriginGetProtocolFunction, err = securityOriginStruct.InvokerNew("get_protocol")
	})
	return err
}

// GetProtocol is a representation of the C type webkit_security_origin_get_protocol.
func (recv *SecurityOrigin) GetProtocol() (string, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := securityOriginGetProtocolFunction_Set()
	if err == nil {
		ret = securityOriginGetProtocolFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo, err
}

// UNSUPPORTED : C value 'webkit_security_origin_is_opaque' : return type 'gboolean' not supported

var securityOriginRefFunction *gi.Function
var securityOriginRefFunction_Once sync.Once

func securityOriginRefFunction_Set() error {
	var err error
	securityOriginRefFunction_Once.Do(func() {
		err = securityOriginStruct_Set()
		if err != nil {
			return
		}
		securityOriginRefFunction, err = securityOriginStruct.InvokerNew("ref")
	})
	return err
}

// Ref is a representation of the C type webkit_security_origin_ref.
func (recv *SecurityOrigin) Ref() (*SecurityOrigin, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := securityOriginRefFunction_Set()
	if err == nil {
		ret = securityOriginRefFunction.Invoke(inArgs[:], nil)
	}

	retGo := &SecurityOrigin{native: ret.Pointer()}

	return retGo, err
}

var securityOriginToStringFunction *gi.Function
var securityOriginToStringFunction_Once sync.Once

func securityOriginToStringFunction_Set() error {
	var err error
	securityOriginToStringFunction_Once.Do(func() {
		err = securityOriginStruct_Set()
		if err != nil {
			return
		}
		securityOriginToStringFunction, err = securityOriginStruct.InvokerNew("to_string")
	})
	return err
}

// ToString is a representation of the C type webkit_security_origin_to_string.
func (recv *SecurityOrigin) ToString() (string, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := securityOriginToStringFunction_Set()
	if err == nil {
		ret = securityOriginToStringFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo, err
}

var securityOriginUnrefFunction *gi.Function
var securityOriginUnrefFunction_Once sync.Once

func securityOriginUnrefFunction_Set() error {
	var err error
	securityOriginUnrefFunction_Once.Do(func() {
		err = securityOriginStruct_Set()
		if err != nil {
			return
		}
		securityOriginUnrefFunction, err = securityOriginStruct.InvokerNew("unref")
	})
	return err
}

// Unref is a representation of the C type webkit_security_origin_unref.
func (recv *SecurityOrigin) Unref() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := securityOriginUnrefFunction_Set()
	if err == nil {
		securityOriginUnrefFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var settingsClassStruct *gi.Struct
var settingsClassStruct_Once sync.Once

func settingsClassStruct_Set() error {
	var err error
	settingsClassStruct_Once.Do(func() {
		settingsClassStruct, err = gi.StructNew("WebKit2", "SettingsClass")
	})
	return err
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
var settingsPrivateStruct_Once sync.Once

func settingsPrivateStruct_Set() error {
	var err error
	settingsPrivateStruct_Once.Do(func() {
		settingsPrivateStruct, err = gi.StructNew("WebKit2", "SettingsPrivate")
	})
	return err
}

type SettingsPrivate struct {
	native uintptr
}

var uRIRequestClassStruct *gi.Struct
var uRIRequestClassStruct_Once sync.Once

func uRIRequestClassStruct_Set() error {
	var err error
	uRIRequestClassStruct_Once.Do(func() {
		uRIRequestClassStruct, err = gi.StructNew("WebKit2", "URIRequestClass")
	})
	return err
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
var uRIRequestPrivateStruct_Once sync.Once

func uRIRequestPrivateStruct_Set() error {
	var err error
	uRIRequestPrivateStruct_Once.Do(func() {
		uRIRequestPrivateStruct, err = gi.StructNew("WebKit2", "URIRequestPrivate")
	})
	return err
}

type URIRequestPrivate struct {
	native uintptr
}

var uRIResponseClassStruct *gi.Struct
var uRIResponseClassStruct_Once sync.Once

func uRIResponseClassStruct_Set() error {
	var err error
	uRIResponseClassStruct_Once.Do(func() {
		uRIResponseClassStruct, err = gi.StructNew("WebKit2", "URIResponseClass")
	})
	return err
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
var uRIResponsePrivateStruct_Once sync.Once

func uRIResponsePrivateStruct_Set() error {
	var err error
	uRIResponsePrivateStruct_Once.Do(func() {
		uRIResponsePrivateStruct, err = gi.StructNew("WebKit2", "URIResponsePrivate")
	})
	return err
}

type URIResponsePrivate struct {
	native uintptr
}

var uRISchemeRequestClassStruct *gi.Struct
var uRISchemeRequestClassStruct_Once sync.Once

func uRISchemeRequestClassStruct_Set() error {
	var err error
	uRISchemeRequestClassStruct_Once.Do(func() {
		uRISchemeRequestClassStruct, err = gi.StructNew("WebKit2", "URISchemeRequestClass")
	})
	return err
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
var uRISchemeRequestPrivateStruct_Once sync.Once

func uRISchemeRequestPrivateStruct_Set() error {
	var err error
	uRISchemeRequestPrivateStruct_Once.Do(func() {
		uRISchemeRequestPrivateStruct, err = gi.StructNew("WebKit2", "URISchemeRequestPrivate")
	})
	return err
}

type URISchemeRequestPrivate struct {
	native uintptr
}

var userContentFilterStruct *gi.Struct
var userContentFilterStruct_Once sync.Once

func userContentFilterStruct_Set() error {
	var err error
	userContentFilterStruct_Once.Do(func() {
		userContentFilterStruct, err = gi.StructNew("WebKit2", "UserContentFilter")
	})
	return err
}

type UserContentFilter struct {
	native uintptr
}

var userContentFilterGetIdentifierFunction *gi.Function
var userContentFilterGetIdentifierFunction_Once sync.Once

func userContentFilterGetIdentifierFunction_Set() error {
	var err error
	userContentFilterGetIdentifierFunction_Once.Do(func() {
		err = userContentFilterStruct_Set()
		if err != nil {
			return
		}
		userContentFilterGetIdentifierFunction, err = userContentFilterStruct.InvokerNew("get_identifier")
	})
	return err
}

// GetIdentifier is a representation of the C type webkit_user_content_filter_get_identifier.
func (recv *UserContentFilter) GetIdentifier() (string, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := userContentFilterGetIdentifierFunction_Set()
	if err == nil {
		ret = userContentFilterGetIdentifierFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo, err
}

var userContentFilterRefFunction *gi.Function
var userContentFilterRefFunction_Once sync.Once

func userContentFilterRefFunction_Set() error {
	var err error
	userContentFilterRefFunction_Once.Do(func() {
		err = userContentFilterStruct_Set()
		if err != nil {
			return
		}
		userContentFilterRefFunction, err = userContentFilterStruct.InvokerNew("ref")
	})
	return err
}

// Ref is a representation of the C type webkit_user_content_filter_ref.
func (recv *UserContentFilter) Ref() (*UserContentFilter, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := userContentFilterRefFunction_Set()
	if err == nil {
		ret = userContentFilterRefFunction.Invoke(inArgs[:], nil)
	}

	retGo := &UserContentFilter{native: ret.Pointer()}

	return retGo, err
}

var userContentFilterUnrefFunction *gi.Function
var userContentFilterUnrefFunction_Once sync.Once

func userContentFilterUnrefFunction_Set() error {
	var err error
	userContentFilterUnrefFunction_Once.Do(func() {
		err = userContentFilterStruct_Set()
		if err != nil {
			return
		}
		userContentFilterUnrefFunction, err = userContentFilterStruct.InvokerNew("unref")
	})
	return err
}

// Unref is a representation of the C type webkit_user_content_filter_unref.
func (recv *UserContentFilter) Unref() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := userContentFilterUnrefFunction_Set()
	if err == nil {
		userContentFilterUnrefFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var userContentFilterStoreClassStruct *gi.Struct
var userContentFilterStoreClassStruct_Once sync.Once

func userContentFilterStoreClassStruct_Set() error {
	var err error
	userContentFilterStoreClassStruct_Once.Do(func() {
		userContentFilterStoreClassStruct, err = gi.StructNew("WebKit2", "UserContentFilterStoreClass")
	})
	return err
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
var userContentFilterStorePrivateStruct_Once sync.Once

func userContentFilterStorePrivateStruct_Set() error {
	var err error
	userContentFilterStorePrivateStruct_Once.Do(func() {
		userContentFilterStorePrivateStruct, err = gi.StructNew("WebKit2", "UserContentFilterStorePrivate")
	})
	return err
}

type UserContentFilterStorePrivate struct {
	native uintptr
}

var userContentManagerClassStruct *gi.Struct
var userContentManagerClassStruct_Once sync.Once

func userContentManagerClassStruct_Set() error {
	var err error
	userContentManagerClassStruct_Once.Do(func() {
		userContentManagerClassStruct, err = gi.StructNew("WebKit2", "UserContentManagerClass")
	})
	return err
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
var userContentManagerPrivateStruct_Once sync.Once

func userContentManagerPrivateStruct_Set() error {
	var err error
	userContentManagerPrivateStruct_Once.Do(func() {
		userContentManagerPrivateStruct, err = gi.StructNew("WebKit2", "UserContentManagerPrivate")
	})
	return err
}

type UserContentManagerPrivate struct {
	native uintptr
}

var userMediaPermissionRequestClassStruct *gi.Struct
var userMediaPermissionRequestClassStruct_Once sync.Once

func userMediaPermissionRequestClassStruct_Set() error {
	var err error
	userMediaPermissionRequestClassStruct_Once.Do(func() {
		userMediaPermissionRequestClassStruct, err = gi.StructNew("WebKit2", "UserMediaPermissionRequestClass")
	})
	return err
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
var userMediaPermissionRequestPrivateStruct_Once sync.Once

func userMediaPermissionRequestPrivateStruct_Set() error {
	var err error
	userMediaPermissionRequestPrivateStruct_Once.Do(func() {
		userMediaPermissionRequestPrivateStruct, err = gi.StructNew("WebKit2", "UserMediaPermissionRequestPrivate")
	})
	return err
}

type UserMediaPermissionRequestPrivate struct {
	native uintptr
}

var userScriptStruct *gi.Struct
var userScriptStruct_Once sync.Once

func userScriptStruct_Set() error {
	var err error
	userScriptStruct_Once.Do(func() {
		userScriptStruct, err = gi.StructNew("WebKit2", "UserScript")
	})
	return err
}

type UserScript struct {
	native uintptr
}

// UNSUPPORTED : C value 'webkit_user_script_new' : parameter 'injected_frames' of type 'UserContentInjectedFrames' not supported

// UNSUPPORTED : C value 'webkit_user_script_new_for_world' : parameter 'injected_frames' of type 'UserContentInjectedFrames' not supported

var userScriptRefFunction *gi.Function
var userScriptRefFunction_Once sync.Once

func userScriptRefFunction_Set() error {
	var err error
	userScriptRefFunction_Once.Do(func() {
		err = userScriptStruct_Set()
		if err != nil {
			return
		}
		userScriptRefFunction, err = userScriptStruct.InvokerNew("ref")
	})
	return err
}

// Ref is a representation of the C type webkit_user_script_ref.
func (recv *UserScript) Ref() (*UserScript, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := userScriptRefFunction_Set()
	if err == nil {
		ret = userScriptRefFunction.Invoke(inArgs[:], nil)
	}

	retGo := &UserScript{native: ret.Pointer()}

	return retGo, err
}

var userScriptUnrefFunction *gi.Function
var userScriptUnrefFunction_Once sync.Once

func userScriptUnrefFunction_Set() error {
	var err error
	userScriptUnrefFunction_Once.Do(func() {
		err = userScriptStruct_Set()
		if err != nil {
			return
		}
		userScriptUnrefFunction, err = userScriptStruct.InvokerNew("unref")
	})
	return err
}

// Unref is a representation of the C type webkit_user_script_unref.
func (recv *UserScript) Unref() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := userScriptUnrefFunction_Set()
	if err == nil {
		userScriptUnrefFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var userStyleSheetStruct *gi.Struct
var userStyleSheetStruct_Once sync.Once

func userStyleSheetStruct_Set() error {
	var err error
	userStyleSheetStruct_Once.Do(func() {
		userStyleSheetStruct, err = gi.StructNew("WebKit2", "UserStyleSheet")
	})
	return err
}

type UserStyleSheet struct {
	native uintptr
}

// UNSUPPORTED : C value 'webkit_user_style_sheet_new' : parameter 'injected_frames' of type 'UserContentInjectedFrames' not supported

// UNSUPPORTED : C value 'webkit_user_style_sheet_new_for_world' : parameter 'injected_frames' of type 'UserContentInjectedFrames' not supported

var userStyleSheetRefFunction *gi.Function
var userStyleSheetRefFunction_Once sync.Once

func userStyleSheetRefFunction_Set() error {
	var err error
	userStyleSheetRefFunction_Once.Do(func() {
		err = userStyleSheetStruct_Set()
		if err != nil {
			return
		}
		userStyleSheetRefFunction, err = userStyleSheetStruct.InvokerNew("ref")
	})
	return err
}

// Ref is a representation of the C type webkit_user_style_sheet_ref.
func (recv *UserStyleSheet) Ref() (*UserStyleSheet, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := userStyleSheetRefFunction_Set()
	if err == nil {
		ret = userStyleSheetRefFunction.Invoke(inArgs[:], nil)
	}

	retGo := &UserStyleSheet{native: ret.Pointer()}

	return retGo, err
}

var userStyleSheetUnrefFunction *gi.Function
var userStyleSheetUnrefFunction_Once sync.Once

func userStyleSheetUnrefFunction_Set() error {
	var err error
	userStyleSheetUnrefFunction_Once.Do(func() {
		err = userStyleSheetStruct_Set()
		if err != nil {
			return
		}
		userStyleSheetUnrefFunction, err = userStyleSheetStruct.InvokerNew("unref")
	})
	return err
}

// Unref is a representation of the C type webkit_user_style_sheet_unref.
func (recv *UserStyleSheet) Unref() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := userStyleSheetUnrefFunction_Set()
	if err == nil {
		userStyleSheetUnrefFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var webContextClassStruct *gi.Struct
var webContextClassStruct_Once sync.Once

func webContextClassStruct_Set() error {
	var err error
	webContextClassStruct_Once.Do(func() {
		webContextClassStruct, err = gi.StructNew("WebKit2", "WebContextClass")
	})
	return err
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
var webContextPrivateStruct_Once sync.Once

func webContextPrivateStruct_Set() error {
	var err error
	webContextPrivateStruct_Once.Do(func() {
		webContextPrivateStruct, err = gi.StructNew("WebKit2", "WebContextPrivate")
	})
	return err
}

type WebContextPrivate struct {
	native uintptr
}

var webInspectorClassStruct *gi.Struct
var webInspectorClassStruct_Once sync.Once

func webInspectorClassStruct_Set() error {
	var err error
	webInspectorClassStruct_Once.Do(func() {
		webInspectorClassStruct, err = gi.StructNew("WebKit2", "WebInspectorClass")
	})
	return err
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
var webInspectorPrivateStruct_Once sync.Once

func webInspectorPrivateStruct_Set() error {
	var err error
	webInspectorPrivateStruct_Once.Do(func() {
		webInspectorPrivateStruct, err = gi.StructNew("WebKit2", "WebInspectorPrivate")
	})
	return err
}

type WebInspectorPrivate struct {
	native uintptr
}

var webResourceClassStruct *gi.Struct
var webResourceClassStruct_Once sync.Once

func webResourceClassStruct_Set() error {
	var err error
	webResourceClassStruct_Once.Do(func() {
		webResourceClassStruct, err = gi.StructNew("WebKit2", "WebResourceClass")
	})
	return err
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
var webResourcePrivateStruct_Once sync.Once

func webResourcePrivateStruct_Set() error {
	var err error
	webResourcePrivateStruct_Once.Do(func() {
		webResourcePrivateStruct, err = gi.StructNew("WebKit2", "WebResourcePrivate")
	})
	return err
}

type WebResourcePrivate struct {
	native uintptr
}

var webViewBaseClassStruct *gi.Struct
var webViewBaseClassStruct_Once sync.Once

func webViewBaseClassStruct_Set() error {
	var err error
	webViewBaseClassStruct_Once.Do(func() {
		webViewBaseClassStruct, err = gi.StructNew("WebKit2", "WebViewBaseClass")
	})
	return err
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
var webViewBasePrivateStruct_Once sync.Once

func webViewBasePrivateStruct_Set() error {
	var err error
	webViewBasePrivateStruct_Once.Do(func() {
		webViewBasePrivateStruct, err = gi.StructNew("WebKit2", "WebViewBasePrivate")
	})
	return err
}

type WebViewBasePrivate struct {
	native uintptr
}

var webViewClassStruct *gi.Struct
var webViewClassStruct_Once sync.Once

func webViewClassStruct_Set() error {
	var err error
	webViewClassStruct_Once.Do(func() {
		webViewClassStruct, err = gi.StructNew("WebKit2", "WebViewClass")
	})
	return err
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
var webViewPrivateStruct_Once sync.Once

func webViewPrivateStruct_Set() error {
	var err error
	webViewPrivateStruct_Once.Do(func() {
		webViewPrivateStruct, err = gi.StructNew("WebKit2", "WebViewPrivate")
	})
	return err
}

type WebViewPrivate struct {
	native uintptr
}

var webViewSessionStateStruct *gi.Struct
var webViewSessionStateStruct_Once sync.Once

func webViewSessionStateStruct_Set() error {
	var err error
	webViewSessionStateStruct_Once.Do(func() {
		webViewSessionStateStruct, err = gi.StructNew("WebKit2", "WebViewSessionState")
	})
	return err
}

type WebViewSessionState struct {
	native uintptr
}

// UNSUPPORTED : C value 'webkit_web_view_session_state_new' : parameter 'data' of type 'GLib.Bytes' not supported

var webViewSessionStateRefFunction *gi.Function
var webViewSessionStateRefFunction_Once sync.Once

func webViewSessionStateRefFunction_Set() error {
	var err error
	webViewSessionStateRefFunction_Once.Do(func() {
		err = webViewSessionStateStruct_Set()
		if err != nil {
			return
		}
		webViewSessionStateRefFunction, err = webViewSessionStateStruct.InvokerNew("ref")
	})
	return err
}

// Ref is a representation of the C type webkit_web_view_session_state_ref.
func (recv *WebViewSessionState) Ref() (*WebViewSessionState, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := webViewSessionStateRefFunction_Set()
	if err == nil {
		ret = webViewSessionStateRefFunction.Invoke(inArgs[:], nil)
	}

	retGo := &WebViewSessionState{native: ret.Pointer()}

	return retGo, err
}

// UNSUPPORTED : C value 'webkit_web_view_session_state_serialize' : return type 'GLib.Bytes' not supported

var webViewSessionStateUnrefFunction *gi.Function
var webViewSessionStateUnrefFunction_Once sync.Once

func webViewSessionStateUnrefFunction_Set() error {
	var err error
	webViewSessionStateUnrefFunction_Once.Do(func() {
		err = webViewSessionStateStruct_Set()
		if err != nil {
			return
		}
		webViewSessionStateUnrefFunction, err = webViewSessionStateStruct.InvokerNew("unref")
	})
	return err
}

// Unref is a representation of the C type webkit_web_view_session_state_unref.
func (recv *WebViewSessionState) Unref() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := webViewSessionStateUnrefFunction_Set()
	if err == nil {
		webViewSessionStateUnrefFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var websiteDataStruct *gi.Struct
var websiteDataStruct_Once sync.Once

func websiteDataStruct_Set() error {
	var err error
	websiteDataStruct_Once.Do(func() {
		websiteDataStruct, err = gi.StructNew("WebKit2", "WebsiteData")
	})
	return err
}

type WebsiteData struct {
	native uintptr
}

var websiteDataGetNameFunction *gi.Function
var websiteDataGetNameFunction_Once sync.Once

func websiteDataGetNameFunction_Set() error {
	var err error
	websiteDataGetNameFunction_Once.Do(func() {
		err = websiteDataStruct_Set()
		if err != nil {
			return
		}
		websiteDataGetNameFunction, err = websiteDataStruct.InvokerNew("get_name")
	})
	return err
}

// GetName is a representation of the C type webkit_website_data_get_name.
func (recv *WebsiteData) GetName() (string, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := websiteDataGetNameFunction_Set()
	if err == nil {
		ret = websiteDataGetNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo, err
}

// UNSUPPORTED : C value 'webkit_website_data_get_size' : parameter 'types' of type 'WebsiteDataTypes' not supported

// UNSUPPORTED : C value 'webkit_website_data_get_types' : return type 'WebsiteDataTypes' not supported

var websiteDataRefFunction *gi.Function
var websiteDataRefFunction_Once sync.Once

func websiteDataRefFunction_Set() error {
	var err error
	websiteDataRefFunction_Once.Do(func() {
		err = websiteDataStruct_Set()
		if err != nil {
			return
		}
		websiteDataRefFunction, err = websiteDataStruct.InvokerNew("ref")
	})
	return err
}

// Ref is a representation of the C type webkit_website_data_ref.
func (recv *WebsiteData) Ref() (*WebsiteData, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := websiteDataRefFunction_Set()
	if err == nil {
		ret = websiteDataRefFunction.Invoke(inArgs[:], nil)
	}

	retGo := &WebsiteData{native: ret.Pointer()}

	return retGo, err
}

var websiteDataUnrefFunction *gi.Function
var websiteDataUnrefFunction_Once sync.Once

func websiteDataUnrefFunction_Set() error {
	var err error
	websiteDataUnrefFunction_Once.Do(func() {
		err = websiteDataStruct_Set()
		if err != nil {
			return
		}
		websiteDataUnrefFunction, err = websiteDataStruct.InvokerNew("unref")
	})
	return err
}

// Unref is a representation of the C type webkit_website_data_unref.
func (recv *WebsiteData) Unref() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := websiteDataUnrefFunction_Set()
	if err == nil {
		websiteDataUnrefFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var websiteDataManagerClassStruct *gi.Struct
var websiteDataManagerClassStruct_Once sync.Once

func websiteDataManagerClassStruct_Set() error {
	var err error
	websiteDataManagerClassStruct_Once.Do(func() {
		websiteDataManagerClassStruct, err = gi.StructNew("WebKit2", "WebsiteDataManagerClass")
	})
	return err
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
var websiteDataManagerPrivateStruct_Once sync.Once

func websiteDataManagerPrivateStruct_Set() error {
	var err error
	websiteDataManagerPrivateStruct_Once.Do(func() {
		websiteDataManagerPrivateStruct, err = gi.StructNew("WebKit2", "WebsiteDataManagerPrivate")
	})
	return err
}

type WebsiteDataManagerPrivate struct {
	native uintptr
}

var windowPropertiesClassStruct *gi.Struct
var windowPropertiesClassStruct_Once sync.Once

func windowPropertiesClassStruct_Set() error {
	var err error
	windowPropertiesClassStruct_Once.Do(func() {
		windowPropertiesClassStruct, err = gi.StructNew("WebKit2", "WindowPropertiesClass")
	})
	return err
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
var windowPropertiesPrivateStruct_Once sync.Once

func windowPropertiesPrivateStruct_Set() error {
	var err error
	windowPropertiesPrivateStruct_Once.Do(func() {
		windowPropertiesPrivateStruct, err = gi.StructNew("WebKit2", "WindowPropertiesPrivate")
	})
	return err
}

type WindowPropertiesPrivate struct {
	native uintptr
}
