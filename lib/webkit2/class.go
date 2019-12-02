// Code generated - DO NOT EDIT.

package webkit2

import (
	gi "github.com/pekim/gobbi/internal/gi"
	"runtime"
	"sync"
)

var authenticationRequestStruct *gi.Struct
var authenticationRequestStruct_Once sync.Once

func authenticationRequestStruct_Set() error {
	var err error
	authenticationRequestStruct_Once.Do(func() {
		authenticationRequestStruct, err = gi.StructNew("WebKit2", "AuthenticationRequest")
	})
	return err
}

type AuthenticationRequest struct {
	native uintptr
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'GObject.Object'

var authenticationRequestAuthenticateFunction *gi.Function
var authenticationRequestAuthenticateFunction_Once sync.Once

func authenticationRequestAuthenticateFunction_Set() error {
	var err error
	authenticationRequestAuthenticateFunction_Once.Do(func() {
		err = authenticationRequestStruct_Set()
		if err != nil {
			return
		}
		authenticationRequestAuthenticateFunction, err = authenticationRequestStruct.InvokerNew("authenticate")
	})
	return err
}

// Authenticate is a representation of the C type webkit_authentication_request_authenticate.
func (recv *AuthenticationRequest) Authenticate(credential *Credential) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(credential.native)

	err := authenticationRequestAuthenticateFunction_Set()
	if err == nil {
		authenticationRequestAuthenticateFunction.Invoke(inArgs[:], nil)
	}

	return
}

var authenticationRequestCanSaveCredentialsFunction *gi.Function
var authenticationRequestCanSaveCredentialsFunction_Once sync.Once

func authenticationRequestCanSaveCredentialsFunction_Set() error {
	var err error
	authenticationRequestCanSaveCredentialsFunction_Once.Do(func() {
		err = authenticationRequestStruct_Set()
		if err != nil {
			return
		}
		authenticationRequestCanSaveCredentialsFunction, err = authenticationRequestStruct.InvokerNew("can_save_credentials")
	})
	return err
}

// CanSaveCredentials is a representation of the C type webkit_authentication_request_can_save_credentials.
func (recv *AuthenticationRequest) CanSaveCredentials() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := authenticationRequestCanSaveCredentialsFunction_Set()
	if err == nil {
		ret = authenticationRequestCanSaveCredentialsFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var authenticationRequestCancelFunction *gi.Function
var authenticationRequestCancelFunction_Once sync.Once

func authenticationRequestCancelFunction_Set() error {
	var err error
	authenticationRequestCancelFunction_Once.Do(func() {
		err = authenticationRequestStruct_Set()
		if err != nil {
			return
		}
		authenticationRequestCancelFunction, err = authenticationRequestStruct.InvokerNew("cancel")
	})
	return err
}

// Cancel is a representation of the C type webkit_authentication_request_cancel.
func (recv *AuthenticationRequest) Cancel() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := authenticationRequestCancelFunction_Set()
	if err == nil {
		authenticationRequestCancelFunction.Invoke(inArgs[:], nil)
	}

	return
}

var authenticationRequestGetHostFunction *gi.Function
var authenticationRequestGetHostFunction_Once sync.Once

func authenticationRequestGetHostFunction_Set() error {
	var err error
	authenticationRequestGetHostFunction_Once.Do(func() {
		err = authenticationRequestStruct_Set()
		if err != nil {
			return
		}
		authenticationRequestGetHostFunction, err = authenticationRequestStruct.InvokerNew("get_host")
	})
	return err
}

// GetHost is a representation of the C type webkit_authentication_request_get_host.
func (recv *AuthenticationRequest) GetHost() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := authenticationRequestGetHostFunction_Set()
	if err == nil {
		ret = authenticationRequestGetHostFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var authenticationRequestGetPortFunction *gi.Function
var authenticationRequestGetPortFunction_Once sync.Once

func authenticationRequestGetPortFunction_Set() error {
	var err error
	authenticationRequestGetPortFunction_Once.Do(func() {
		err = authenticationRequestStruct_Set()
		if err != nil {
			return
		}
		authenticationRequestGetPortFunction, err = authenticationRequestStruct.InvokerNew("get_port")
	})
	return err
}

// GetPort is a representation of the C type webkit_authentication_request_get_port.
func (recv *AuthenticationRequest) GetPort() uint32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := authenticationRequestGetPortFunction_Set()
	if err == nil {
		ret = authenticationRequestGetPortFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo
}

var authenticationRequestGetProposedCredentialFunction *gi.Function
var authenticationRequestGetProposedCredentialFunction_Once sync.Once

func authenticationRequestGetProposedCredentialFunction_Set() error {
	var err error
	authenticationRequestGetProposedCredentialFunction_Once.Do(func() {
		err = authenticationRequestStruct_Set()
		if err != nil {
			return
		}
		authenticationRequestGetProposedCredentialFunction, err = authenticationRequestStruct.InvokerNew("get_proposed_credential")
	})
	return err
}

// GetProposedCredential is a representation of the C type webkit_authentication_request_get_proposed_credential.
func (recv *AuthenticationRequest) GetProposedCredential() *Credential {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := authenticationRequestGetProposedCredentialFunction_Set()
	if err == nil {
		ret = authenticationRequestGetProposedCredentialFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Credential{native: ret.Pointer()}

	return retGo
}

var authenticationRequestGetRealmFunction *gi.Function
var authenticationRequestGetRealmFunction_Once sync.Once

func authenticationRequestGetRealmFunction_Set() error {
	var err error
	authenticationRequestGetRealmFunction_Once.Do(func() {
		err = authenticationRequestStruct_Set()
		if err != nil {
			return
		}
		authenticationRequestGetRealmFunction, err = authenticationRequestStruct.InvokerNew("get_realm")
	})
	return err
}

// GetRealm is a representation of the C type webkit_authentication_request_get_realm.
func (recv *AuthenticationRequest) GetRealm() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := authenticationRequestGetRealmFunction_Set()
	if err == nil {
		ret = authenticationRequestGetRealmFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'webkit_authentication_request_get_scheme' : return type 'AuthenticationScheme' not supported

var authenticationRequestIsForProxyFunction *gi.Function
var authenticationRequestIsForProxyFunction_Once sync.Once

func authenticationRequestIsForProxyFunction_Set() error {
	var err error
	authenticationRequestIsForProxyFunction_Once.Do(func() {
		err = authenticationRequestStruct_Set()
		if err != nil {
			return
		}
		authenticationRequestIsForProxyFunction, err = authenticationRequestStruct.InvokerNew("is_for_proxy")
	})
	return err
}

// IsForProxy is a representation of the C type webkit_authentication_request_is_for_proxy.
func (recv *AuthenticationRequest) IsForProxy() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := authenticationRequestIsForProxyFunction_Set()
	if err == nil {
		ret = authenticationRequestIsForProxyFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var authenticationRequestIsRetryFunction *gi.Function
var authenticationRequestIsRetryFunction_Once sync.Once

func authenticationRequestIsRetryFunction_Set() error {
	var err error
	authenticationRequestIsRetryFunction_Once.Do(func() {
		err = authenticationRequestStruct_Set()
		if err != nil {
			return
		}
		authenticationRequestIsRetryFunction, err = authenticationRequestStruct.InvokerNew("is_retry")
	})
	return err
}

// IsRetry is a representation of the C type webkit_authentication_request_is_retry.
func (recv *AuthenticationRequest) IsRetry() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := authenticationRequestIsRetryFunction_Set()
	if err == nil {
		ret = authenticationRequestIsRetryFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// AuthenticationRequestStruct creates an uninitialised AuthenticationRequest.
func AuthenticationRequestStruct() *AuthenticationRequest {
	err := authenticationRequestStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &AuthenticationRequest{native: authenticationRequestStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeAuthenticationRequest)
	return structGo
}
func finalizeAuthenticationRequest(obj *AuthenticationRequest) {
	authenticationRequestStruct.Free(obj.native)
}

var automationSessionStruct *gi.Struct
var automationSessionStruct_Once sync.Once

func automationSessionStruct_Set() error {
	var err error
	automationSessionStruct_Once.Do(func() {
		automationSessionStruct, err = gi.StructNew("WebKit2", "AutomationSession")
	})
	return err
}

type AutomationSession struct {
	native uintptr
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'GObject.Object'

// FieldPriv returns the C field 'priv'.
func (recv *AutomationSession) FieldPriv() *AutomationSessionPrivate {
	argValue := gi.FieldGet(automationSessionStruct, recv.native, "priv")
	value := &AutomationSessionPrivate{native: argValue.Pointer()}
	return value
}

// SetFieldPriv sets the value of the C field 'priv'.
func (recv *AutomationSession) SetFieldPriv(value *AutomationSessionPrivate) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(automationSessionStruct, recv.native, "priv", argValue)
}

var automationSessionGetApplicationInfoFunction *gi.Function
var automationSessionGetApplicationInfoFunction_Once sync.Once

func automationSessionGetApplicationInfoFunction_Set() error {
	var err error
	automationSessionGetApplicationInfoFunction_Once.Do(func() {
		err = automationSessionStruct_Set()
		if err != nil {
			return
		}
		automationSessionGetApplicationInfoFunction, err = automationSessionStruct.InvokerNew("get_application_info")
	})
	return err
}

// GetApplicationInfo is a representation of the C type webkit_automation_session_get_application_info.
func (recv *AutomationSession) GetApplicationInfo() *ApplicationInfo {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := automationSessionGetApplicationInfoFunction_Set()
	if err == nil {
		ret = automationSessionGetApplicationInfoFunction.Invoke(inArgs[:], nil)
	}

	retGo := &ApplicationInfo{native: ret.Pointer()}

	return retGo
}

var automationSessionGetIdFunction *gi.Function
var automationSessionGetIdFunction_Once sync.Once

func automationSessionGetIdFunction_Set() error {
	var err error
	automationSessionGetIdFunction_Once.Do(func() {
		err = automationSessionStruct_Set()
		if err != nil {
			return
		}
		automationSessionGetIdFunction, err = automationSessionStruct.InvokerNew("get_id")
	})
	return err
}

// GetId is a representation of the C type webkit_automation_session_get_id.
func (recv *AutomationSession) GetId() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := automationSessionGetIdFunction_Set()
	if err == nil {
		ret = automationSessionGetIdFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var automationSessionSetApplicationInfoFunction *gi.Function
var automationSessionSetApplicationInfoFunction_Once sync.Once

func automationSessionSetApplicationInfoFunction_Set() error {
	var err error
	automationSessionSetApplicationInfoFunction_Once.Do(func() {
		err = automationSessionStruct_Set()
		if err != nil {
			return
		}
		automationSessionSetApplicationInfoFunction, err = automationSessionStruct.InvokerNew("set_application_info")
	})
	return err
}

// SetApplicationInfo is a representation of the C type webkit_automation_session_set_application_info.
func (recv *AutomationSession) SetApplicationInfo(info *ApplicationInfo) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(info.native)

	err := automationSessionSetApplicationInfoFunction_Set()
	if err == nil {
		automationSessionSetApplicationInfoFunction.Invoke(inArgs[:], nil)
	}

	return
}

// AutomationSessionStruct creates an uninitialised AutomationSession.
func AutomationSessionStruct() *AutomationSession {
	err := automationSessionStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &AutomationSession{native: automationSessionStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeAutomationSession)
	return structGo
}
func finalizeAutomationSession(obj *AutomationSession) {
	automationSessionStruct.Free(obj.native)
}

var backForwardListStruct *gi.Struct
var backForwardListStruct_Once sync.Once

func backForwardListStruct_Set() error {
	var err error
	backForwardListStruct_Once.Do(func() {
		backForwardListStruct, err = gi.StructNew("WebKit2", "BackForwardList")
	})
	return err
}

type BackForwardList struct {
	native uintptr
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'GObject.Object'

// FieldPriv returns the C field 'priv'.
func (recv *BackForwardList) FieldPriv() *BackForwardListPrivate {
	argValue := gi.FieldGet(backForwardListStruct, recv.native, "priv")
	value := &BackForwardListPrivate{native: argValue.Pointer()}
	return value
}

// SetFieldPriv sets the value of the C field 'priv'.
func (recv *BackForwardList) SetFieldPriv(value *BackForwardListPrivate) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(backForwardListStruct, recv.native, "priv", argValue)
}

// UNSUPPORTED : C value 'webkit_back_forward_list_get_back_item' : return type 'BackForwardListItem' not supported

// UNSUPPORTED : C value 'webkit_back_forward_list_get_back_list' : return type 'GLib.List' not supported

// UNSUPPORTED : C value 'webkit_back_forward_list_get_back_list_with_limit' : return type 'GLib.List' not supported

// UNSUPPORTED : C value 'webkit_back_forward_list_get_current_item' : return type 'BackForwardListItem' not supported

// UNSUPPORTED : C value 'webkit_back_forward_list_get_forward_item' : return type 'BackForwardListItem' not supported

// UNSUPPORTED : C value 'webkit_back_forward_list_get_forward_list' : return type 'GLib.List' not supported

// UNSUPPORTED : C value 'webkit_back_forward_list_get_forward_list_with_limit' : return type 'GLib.List' not supported

var backForwardListGetLengthFunction *gi.Function
var backForwardListGetLengthFunction_Once sync.Once

func backForwardListGetLengthFunction_Set() error {
	var err error
	backForwardListGetLengthFunction_Once.Do(func() {
		err = backForwardListStruct_Set()
		if err != nil {
			return
		}
		backForwardListGetLengthFunction, err = backForwardListStruct.InvokerNew("get_length")
	})
	return err
}

// GetLength is a representation of the C type webkit_back_forward_list_get_length.
func (recv *BackForwardList) GetLength() uint32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := backForwardListGetLengthFunction_Set()
	if err == nil {
		ret = backForwardListGetLengthFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo
}

// UNSUPPORTED : C value 'webkit_back_forward_list_get_nth_item' : return type 'BackForwardListItem' not supported

// BackForwardListStruct creates an uninitialised BackForwardList.
func BackForwardListStruct() *BackForwardList {
	err := backForwardListStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &BackForwardList{native: backForwardListStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeBackForwardList)
	return structGo
}
func finalizeBackForwardList(obj *BackForwardList) {
	backForwardListStruct.Free(obj.native)
}

var backForwardListItemStruct *gi.Struct
var backForwardListItemStruct_Once sync.Once

func backForwardListItemStruct_Set() error {
	var err error
	backForwardListItemStruct_Once.Do(func() {
		backForwardListItemStruct, err = gi.StructNew("WebKit2", "BackForwardListItem")
	})
	return err
}

type BackForwardListItem struct {
	native uintptr
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.InitiallyUnowned'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'GObject.InitiallyUnowned'

// FieldPriv returns the C field 'priv'.
func (recv *BackForwardListItem) FieldPriv() *BackForwardListItemPrivate {
	argValue := gi.FieldGet(backForwardListItemStruct, recv.native, "priv")
	value := &BackForwardListItemPrivate{native: argValue.Pointer()}
	return value
}

// SetFieldPriv sets the value of the C field 'priv'.
func (recv *BackForwardListItem) SetFieldPriv(value *BackForwardListItemPrivate) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(backForwardListItemStruct, recv.native, "priv", argValue)
}

var backForwardListItemGetOriginalUriFunction *gi.Function
var backForwardListItemGetOriginalUriFunction_Once sync.Once

func backForwardListItemGetOriginalUriFunction_Set() error {
	var err error
	backForwardListItemGetOriginalUriFunction_Once.Do(func() {
		err = backForwardListItemStruct_Set()
		if err != nil {
			return
		}
		backForwardListItemGetOriginalUriFunction, err = backForwardListItemStruct.InvokerNew("get_original_uri")
	})
	return err
}

// GetOriginalUri is a representation of the C type webkit_back_forward_list_item_get_original_uri.
func (recv *BackForwardListItem) GetOriginalUri() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := backForwardListItemGetOriginalUriFunction_Set()
	if err == nil {
		ret = backForwardListItemGetOriginalUriFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var backForwardListItemGetTitleFunction *gi.Function
var backForwardListItemGetTitleFunction_Once sync.Once

func backForwardListItemGetTitleFunction_Set() error {
	var err error
	backForwardListItemGetTitleFunction_Once.Do(func() {
		err = backForwardListItemStruct_Set()
		if err != nil {
			return
		}
		backForwardListItemGetTitleFunction, err = backForwardListItemStruct.InvokerNew("get_title")
	})
	return err
}

// GetTitle is a representation of the C type webkit_back_forward_list_item_get_title.
func (recv *BackForwardListItem) GetTitle() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := backForwardListItemGetTitleFunction_Set()
	if err == nil {
		ret = backForwardListItemGetTitleFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var backForwardListItemGetUriFunction *gi.Function
var backForwardListItemGetUriFunction_Once sync.Once

func backForwardListItemGetUriFunction_Set() error {
	var err error
	backForwardListItemGetUriFunction_Once.Do(func() {
		err = backForwardListItemStruct_Set()
		if err != nil {
			return
		}
		backForwardListItemGetUriFunction, err = backForwardListItemStruct.InvokerNew("get_uri")
	})
	return err
}

// GetUri is a representation of the C type webkit_back_forward_list_item_get_uri.
func (recv *BackForwardListItem) GetUri() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := backForwardListItemGetUriFunction_Set()
	if err == nil {
		ret = backForwardListItemGetUriFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

// BackForwardListItemStruct creates an uninitialised BackForwardListItem.
func BackForwardListItemStruct() *BackForwardListItem {
	err := backForwardListItemStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &BackForwardListItem{native: backForwardListItemStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeBackForwardListItem)
	return structGo
}
func finalizeBackForwardListItem(obj *BackForwardListItem) {
	backForwardListItemStruct.Free(obj.native)
}

var colorChooserRequestStruct *gi.Struct
var colorChooserRequestStruct_Once sync.Once

func colorChooserRequestStruct_Set() error {
	var err error
	colorChooserRequestStruct_Once.Do(func() {
		colorChooserRequestStruct, err = gi.StructNew("WebKit2", "ColorChooserRequest")
	})
	return err
}

type ColorChooserRequest struct {
	native uintptr
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'GObject.Object'

var colorChooserRequestCancelFunction *gi.Function
var colorChooserRequestCancelFunction_Once sync.Once

func colorChooserRequestCancelFunction_Set() error {
	var err error
	colorChooserRequestCancelFunction_Once.Do(func() {
		err = colorChooserRequestStruct_Set()
		if err != nil {
			return
		}
		colorChooserRequestCancelFunction, err = colorChooserRequestStruct.InvokerNew("cancel")
	})
	return err
}

// Cancel is a representation of the C type webkit_color_chooser_request_cancel.
func (recv *ColorChooserRequest) Cancel() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := colorChooserRequestCancelFunction_Set()
	if err == nil {
		colorChooserRequestCancelFunction.Invoke(inArgs[:], nil)
	}

	return
}

var colorChooserRequestFinishFunction *gi.Function
var colorChooserRequestFinishFunction_Once sync.Once

func colorChooserRequestFinishFunction_Set() error {
	var err error
	colorChooserRequestFinishFunction_Once.Do(func() {
		err = colorChooserRequestStruct_Set()
		if err != nil {
			return
		}
		colorChooserRequestFinishFunction, err = colorChooserRequestStruct.InvokerNew("finish")
	})
	return err
}

// Finish is a representation of the C type webkit_color_chooser_request_finish.
func (recv *ColorChooserRequest) Finish() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := colorChooserRequestFinishFunction_Set()
	if err == nil {
		colorChooserRequestFinishFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'webkit_color_chooser_request_get_element_rectangle' : parameter 'rect' of type 'Gdk.Rectangle' not supported

// UNSUPPORTED : C value 'webkit_color_chooser_request_get_rgba' : parameter 'rgba' of type 'Gdk.RGBA' not supported

// UNSUPPORTED : C value 'webkit_color_chooser_request_set_rgba' : parameter 'rgba' of type 'Gdk.RGBA' not supported

// ColorChooserRequestStruct creates an uninitialised ColorChooserRequest.
func ColorChooserRequestStruct() *ColorChooserRequest {
	err := colorChooserRequestStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ColorChooserRequest{native: colorChooserRequestStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeColorChooserRequest)
	return structGo
}
func finalizeColorChooserRequest(obj *ColorChooserRequest) {
	colorChooserRequestStruct.Free(obj.native)
}

var contextMenuStruct *gi.Struct
var contextMenuStruct_Once sync.Once

func contextMenuStruct_Set() error {
	var err error
	contextMenuStruct_Once.Do(func() {
		contextMenuStruct, err = gi.StructNew("WebKit2", "ContextMenu")
	})
	return err
}

type ContextMenu struct {
	native uintptr
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'GObject.Object'

// FieldPriv returns the C field 'priv'.
func (recv *ContextMenu) FieldPriv() *ContextMenuPrivate {
	argValue := gi.FieldGet(contextMenuStruct, recv.native, "priv")
	value := &ContextMenuPrivate{native: argValue.Pointer()}
	return value
}

// SetFieldPriv sets the value of the C field 'priv'.
func (recv *ContextMenu) SetFieldPriv(value *ContextMenuPrivate) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(contextMenuStruct, recv.native, "priv", argValue)
}

// UNSUPPORTED : C value 'webkit_context_menu_new' : return type 'ContextMenu' not supported

// UNSUPPORTED : C value 'webkit_context_menu_new_with_items' : parameter 'items' of type 'GLib.List' not supported

// UNSUPPORTED : C value 'webkit_context_menu_append' : parameter 'item' of type 'ContextMenuItem' not supported

// UNSUPPORTED : C value 'webkit_context_menu_first' : return type 'ContextMenuItem' not supported

// UNSUPPORTED : C value 'webkit_context_menu_get_item_at_position' : return type 'ContextMenuItem' not supported

// UNSUPPORTED : C value 'webkit_context_menu_get_items' : return type 'GLib.List' not supported

var contextMenuGetNItemsFunction *gi.Function
var contextMenuGetNItemsFunction_Once sync.Once

func contextMenuGetNItemsFunction_Set() error {
	var err error
	contextMenuGetNItemsFunction_Once.Do(func() {
		err = contextMenuStruct_Set()
		if err != nil {
			return
		}
		contextMenuGetNItemsFunction, err = contextMenuStruct.InvokerNew("get_n_items")
	})
	return err
}

// GetNItems is a representation of the C type webkit_context_menu_get_n_items.
func (recv *ContextMenu) GetNItems() uint32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := contextMenuGetNItemsFunction_Set()
	if err == nil {
		ret = contextMenuGetNItemsFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo
}

// UNSUPPORTED : C value 'webkit_context_menu_get_user_data' : return type 'GLib.Variant' not supported

// UNSUPPORTED : C value 'webkit_context_menu_insert' : parameter 'item' of type 'ContextMenuItem' not supported

// UNSUPPORTED : C value 'webkit_context_menu_last' : return type 'ContextMenuItem' not supported

// UNSUPPORTED : C value 'webkit_context_menu_move_item' : parameter 'item' of type 'ContextMenuItem' not supported

// UNSUPPORTED : C value 'webkit_context_menu_prepend' : parameter 'item' of type 'ContextMenuItem' not supported

// UNSUPPORTED : C value 'webkit_context_menu_remove' : parameter 'item' of type 'ContextMenuItem' not supported

var contextMenuRemoveAllFunction *gi.Function
var contextMenuRemoveAllFunction_Once sync.Once

func contextMenuRemoveAllFunction_Set() error {
	var err error
	contextMenuRemoveAllFunction_Once.Do(func() {
		err = contextMenuStruct_Set()
		if err != nil {
			return
		}
		contextMenuRemoveAllFunction, err = contextMenuStruct.InvokerNew("remove_all")
	})
	return err
}

// RemoveAll is a representation of the C type webkit_context_menu_remove_all.
func (recv *ContextMenu) RemoveAll() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := contextMenuRemoveAllFunction_Set()
	if err == nil {
		contextMenuRemoveAllFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'webkit_context_menu_set_user_data' : parameter 'user_data' of type 'GLib.Variant' not supported

var contextMenuItemStruct *gi.Struct
var contextMenuItemStruct_Once sync.Once

func contextMenuItemStruct_Set() error {
	var err error
	contextMenuItemStruct_Once.Do(func() {
		contextMenuItemStruct, err = gi.StructNew("WebKit2", "ContextMenuItem")
	})
	return err
}

type ContextMenuItem struct {
	native uintptr
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.InitiallyUnowned'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'GObject.InitiallyUnowned'

// FieldPriv returns the C field 'priv'.
func (recv *ContextMenuItem) FieldPriv() *ContextMenuItemPrivate {
	argValue := gi.FieldGet(contextMenuItemStruct, recv.native, "priv")
	value := &ContextMenuItemPrivate{native: argValue.Pointer()}
	return value
}

// SetFieldPriv sets the value of the C field 'priv'.
func (recv *ContextMenuItem) SetFieldPriv(value *ContextMenuItemPrivate) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(contextMenuItemStruct, recv.native, "priv", argValue)
}

// UNSUPPORTED : C value 'webkit_context_menu_item_new' : parameter 'action' of type 'Gtk.Action' not supported

// UNSUPPORTED : C value 'webkit_context_menu_item_new_from_gaction' : parameter 'action' of type 'Gio.Action' not supported

// UNSUPPORTED : C value 'webkit_context_menu_item_new_from_stock_action' : parameter 'action' of type 'ContextMenuAction' not supported

// UNSUPPORTED : C value 'webkit_context_menu_item_new_from_stock_action_with_label' : parameter 'action' of type 'ContextMenuAction' not supported

// UNSUPPORTED : C value 'webkit_context_menu_item_new_separator' : return type 'ContextMenuItem' not supported

// UNSUPPORTED : C value 'webkit_context_menu_item_new_with_submenu' : parameter 'submenu' of type 'ContextMenu' not supported

// UNSUPPORTED : C value 'webkit_context_menu_item_get_action' : return type 'Gtk.Action' not supported

// UNSUPPORTED : C value 'webkit_context_menu_item_get_gaction' : return type 'Gio.Action' not supported

// UNSUPPORTED : C value 'webkit_context_menu_item_get_stock_action' : return type 'ContextMenuAction' not supported

// UNSUPPORTED : C value 'webkit_context_menu_item_get_submenu' : return type 'ContextMenu' not supported

var contextMenuItemIsSeparatorFunction *gi.Function
var contextMenuItemIsSeparatorFunction_Once sync.Once

func contextMenuItemIsSeparatorFunction_Set() error {
	var err error
	contextMenuItemIsSeparatorFunction_Once.Do(func() {
		err = contextMenuItemStruct_Set()
		if err != nil {
			return
		}
		contextMenuItemIsSeparatorFunction, err = contextMenuItemStruct.InvokerNew("is_separator")
	})
	return err
}

// IsSeparator is a representation of the C type webkit_context_menu_item_is_separator.
func (recv *ContextMenuItem) IsSeparator() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := contextMenuItemIsSeparatorFunction_Set()
	if err == nil {
		ret = contextMenuItemIsSeparatorFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'webkit_context_menu_item_set_submenu' : parameter 'submenu' of type 'ContextMenu' not supported

var cookieManagerStruct *gi.Struct
var cookieManagerStruct_Once sync.Once

func cookieManagerStruct_Set() error {
	var err error
	cookieManagerStruct_Once.Do(func() {
		cookieManagerStruct, err = gi.StructNew("WebKit2", "CookieManager")
	})
	return err
}

type CookieManager struct {
	native uintptr
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'GObject.Object'

// FieldPriv returns the C field 'priv'.
func (recv *CookieManager) FieldPriv() *CookieManagerPrivate {
	argValue := gi.FieldGet(cookieManagerStruct, recv.native, "priv")
	value := &CookieManagerPrivate{native: argValue.Pointer()}
	return value
}

// SetFieldPriv sets the value of the C field 'priv'.
func (recv *CookieManager) SetFieldPriv(value *CookieManagerPrivate) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(cookieManagerStruct, recv.native, "priv", argValue)
}

// UNSUPPORTED : C value 'webkit_cookie_manager_add_cookie' : parameter 'cookie' of type 'Soup.Cookie' not supported

// UNSUPPORTED : C value 'webkit_cookie_manager_add_cookie_finish' : parameter 'result' of type 'Gio.AsyncResult' not supported

var cookieManagerDeleteAllCookiesFunction *gi.Function
var cookieManagerDeleteAllCookiesFunction_Once sync.Once

func cookieManagerDeleteAllCookiesFunction_Set() error {
	var err error
	cookieManagerDeleteAllCookiesFunction_Once.Do(func() {
		err = cookieManagerStruct_Set()
		if err != nil {
			return
		}
		cookieManagerDeleteAllCookiesFunction, err = cookieManagerStruct.InvokerNew("delete_all_cookies")
	})
	return err
}

// DeleteAllCookies is a representation of the C type webkit_cookie_manager_delete_all_cookies.
func (recv *CookieManager) DeleteAllCookies() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := cookieManagerDeleteAllCookiesFunction_Set()
	if err == nil {
		cookieManagerDeleteAllCookiesFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'webkit_cookie_manager_delete_cookie' : parameter 'cookie' of type 'Soup.Cookie' not supported

// UNSUPPORTED : C value 'webkit_cookie_manager_delete_cookie_finish' : parameter 'result' of type 'Gio.AsyncResult' not supported

var cookieManagerDeleteCookiesForDomainFunction *gi.Function
var cookieManagerDeleteCookiesForDomainFunction_Once sync.Once

func cookieManagerDeleteCookiesForDomainFunction_Set() error {
	var err error
	cookieManagerDeleteCookiesForDomainFunction_Once.Do(func() {
		err = cookieManagerStruct_Set()
		if err != nil {
			return
		}
		cookieManagerDeleteCookiesForDomainFunction, err = cookieManagerStruct.InvokerNew("delete_cookies_for_domain")
	})
	return err
}

// DeleteCookiesForDomain is a representation of the C type webkit_cookie_manager_delete_cookies_for_domain.
func (recv *CookieManager) DeleteCookiesForDomain(domain string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(domain)

	err := cookieManagerDeleteCookiesForDomainFunction_Set()
	if err == nil {
		cookieManagerDeleteCookiesForDomainFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'webkit_cookie_manager_get_accept_policy' : parameter 'cancellable' of type 'Gio.Cancellable' not supported

// UNSUPPORTED : C value 'webkit_cookie_manager_get_accept_policy_finish' : parameter 'result' of type 'Gio.AsyncResult' not supported

// UNSUPPORTED : C value 'webkit_cookie_manager_get_cookies' : parameter 'cancellable' of type 'Gio.Cancellable' not supported

// UNSUPPORTED : C value 'webkit_cookie_manager_get_cookies_finish' : parameter 'result' of type 'Gio.AsyncResult' not supported

// UNSUPPORTED : C value 'webkit_cookie_manager_get_domains_with_cookies' : parameter 'cancellable' of type 'Gio.Cancellable' not supported

// UNSUPPORTED : C value 'webkit_cookie_manager_get_domains_with_cookies_finish' : parameter 'result' of type 'Gio.AsyncResult' not supported

// UNSUPPORTED : C value 'webkit_cookie_manager_set_accept_policy' : parameter 'policy' of type 'CookieAcceptPolicy' not supported

// UNSUPPORTED : C value 'webkit_cookie_manager_set_persistent_storage' : parameter 'storage' of type 'CookiePersistentStorage' not supported

// CookieManagerStruct creates an uninitialised CookieManager.
func CookieManagerStruct() *CookieManager {
	err := cookieManagerStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &CookieManager{native: cookieManagerStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeCookieManager)
	return structGo
}
func finalizeCookieManager(obj *CookieManager) {
	cookieManagerStruct.Free(obj.native)
}

var deviceInfoPermissionRequestStruct *gi.Struct
var deviceInfoPermissionRequestStruct_Once sync.Once

func deviceInfoPermissionRequestStruct_Set() error {
	var err error
	deviceInfoPermissionRequestStruct_Once.Do(func() {
		deviceInfoPermissionRequestStruct, err = gi.StructNew("WebKit2", "DeviceInfoPermissionRequest")
	})
	return err
}

type DeviceInfoPermissionRequest struct {
	native uintptr
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'GObject.Object'

// DeviceInfoPermissionRequestStruct creates an uninitialised DeviceInfoPermissionRequest.
func DeviceInfoPermissionRequestStruct() *DeviceInfoPermissionRequest {
	err := deviceInfoPermissionRequestStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &DeviceInfoPermissionRequest{native: deviceInfoPermissionRequestStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeDeviceInfoPermissionRequest)
	return structGo
}
func finalizeDeviceInfoPermissionRequest(obj *DeviceInfoPermissionRequest) {
	deviceInfoPermissionRequestStruct.Free(obj.native)
}

var downloadStruct *gi.Struct
var downloadStruct_Once sync.Once

func downloadStruct_Set() error {
	var err error
	downloadStruct_Once.Do(func() {
		downloadStruct, err = gi.StructNew("WebKit2", "Download")
	})
	return err
}

type Download struct {
	native uintptr
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'GObject.Object'

// FieldPriv returns the C field 'priv'.
func (recv *Download) FieldPriv() *DownloadPrivate {
	argValue := gi.FieldGet(downloadStruct, recv.native, "priv")
	value := &DownloadPrivate{native: argValue.Pointer()}
	return value
}

// SetFieldPriv sets the value of the C field 'priv'.
func (recv *Download) SetFieldPriv(value *DownloadPrivate) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(downloadStruct, recv.native, "priv", argValue)
}

var downloadCancelFunction *gi.Function
var downloadCancelFunction_Once sync.Once

func downloadCancelFunction_Set() error {
	var err error
	downloadCancelFunction_Once.Do(func() {
		err = downloadStruct_Set()
		if err != nil {
			return
		}
		downloadCancelFunction, err = downloadStruct.InvokerNew("cancel")
	})
	return err
}

// Cancel is a representation of the C type webkit_download_cancel.
func (recv *Download) Cancel() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := downloadCancelFunction_Set()
	if err == nil {
		downloadCancelFunction.Invoke(inArgs[:], nil)
	}

	return
}

var downloadGetAllowOverwriteFunction *gi.Function
var downloadGetAllowOverwriteFunction_Once sync.Once

func downloadGetAllowOverwriteFunction_Set() error {
	var err error
	downloadGetAllowOverwriteFunction_Once.Do(func() {
		err = downloadStruct_Set()
		if err != nil {
			return
		}
		downloadGetAllowOverwriteFunction, err = downloadStruct.InvokerNew("get_allow_overwrite")
	})
	return err
}

// GetAllowOverwrite is a representation of the C type webkit_download_get_allow_overwrite.
func (recv *Download) GetAllowOverwrite() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := downloadGetAllowOverwriteFunction_Set()
	if err == nil {
		ret = downloadGetAllowOverwriteFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var downloadGetDestinationFunction *gi.Function
var downloadGetDestinationFunction_Once sync.Once

func downloadGetDestinationFunction_Set() error {
	var err error
	downloadGetDestinationFunction_Once.Do(func() {
		err = downloadStruct_Set()
		if err != nil {
			return
		}
		downloadGetDestinationFunction, err = downloadStruct.InvokerNew("get_destination")
	})
	return err
}

// GetDestination is a representation of the C type webkit_download_get_destination.
func (recv *Download) GetDestination() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := downloadGetDestinationFunction_Set()
	if err == nil {
		ret = downloadGetDestinationFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var downloadGetElapsedTimeFunction *gi.Function
var downloadGetElapsedTimeFunction_Once sync.Once

func downloadGetElapsedTimeFunction_Set() error {
	var err error
	downloadGetElapsedTimeFunction_Once.Do(func() {
		err = downloadStruct_Set()
		if err != nil {
			return
		}
		downloadGetElapsedTimeFunction, err = downloadStruct.InvokerNew("get_elapsed_time")
	})
	return err
}

// GetElapsedTime is a representation of the C type webkit_download_get_elapsed_time.
func (recv *Download) GetElapsedTime() float64 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := downloadGetElapsedTimeFunction_Set()
	if err == nil {
		ret = downloadGetElapsedTimeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Float64()

	return retGo
}

var downloadGetEstimatedProgressFunction *gi.Function
var downloadGetEstimatedProgressFunction_Once sync.Once

func downloadGetEstimatedProgressFunction_Set() error {
	var err error
	downloadGetEstimatedProgressFunction_Once.Do(func() {
		err = downloadStruct_Set()
		if err != nil {
			return
		}
		downloadGetEstimatedProgressFunction, err = downloadStruct.InvokerNew("get_estimated_progress")
	})
	return err
}

// GetEstimatedProgress is a representation of the C type webkit_download_get_estimated_progress.
func (recv *Download) GetEstimatedProgress() float64 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := downloadGetEstimatedProgressFunction_Set()
	if err == nil {
		ret = downloadGetEstimatedProgressFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Float64()

	return retGo
}

var downloadGetReceivedDataLengthFunction *gi.Function
var downloadGetReceivedDataLengthFunction_Once sync.Once

func downloadGetReceivedDataLengthFunction_Set() error {
	var err error
	downloadGetReceivedDataLengthFunction_Once.Do(func() {
		err = downloadStruct_Set()
		if err != nil {
			return
		}
		downloadGetReceivedDataLengthFunction, err = downloadStruct.InvokerNew("get_received_data_length")
	})
	return err
}

// GetReceivedDataLength is a representation of the C type webkit_download_get_received_data_length.
func (recv *Download) GetReceivedDataLength() uint64 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := downloadGetReceivedDataLengthFunction_Set()
	if err == nil {
		ret = downloadGetReceivedDataLengthFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint64()

	return retGo
}

// UNSUPPORTED : C value 'webkit_download_get_request' : return type 'URIRequest' not supported

// UNSUPPORTED : C value 'webkit_download_get_response' : return type 'URIResponse' not supported

// UNSUPPORTED : C value 'webkit_download_get_web_view' : return type 'WebView' not supported

var downloadSetAllowOverwriteFunction *gi.Function
var downloadSetAllowOverwriteFunction_Once sync.Once

func downloadSetAllowOverwriteFunction_Set() error {
	var err error
	downloadSetAllowOverwriteFunction_Once.Do(func() {
		err = downloadStruct_Set()
		if err != nil {
			return
		}
		downloadSetAllowOverwriteFunction, err = downloadStruct.InvokerNew("set_allow_overwrite")
	})
	return err
}

// SetAllowOverwrite is a representation of the C type webkit_download_set_allow_overwrite.
func (recv *Download) SetAllowOverwrite(allowed bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetBoolean(allowed)

	err := downloadSetAllowOverwriteFunction_Set()
	if err == nil {
		downloadSetAllowOverwriteFunction.Invoke(inArgs[:], nil)
	}

	return
}

var downloadSetDestinationFunction *gi.Function
var downloadSetDestinationFunction_Once sync.Once

func downloadSetDestinationFunction_Set() error {
	var err error
	downloadSetDestinationFunction_Once.Do(func() {
		err = downloadStruct_Set()
		if err != nil {
			return
		}
		downloadSetDestinationFunction, err = downloadStruct.InvokerNew("set_destination")
	})
	return err
}

// SetDestination is a representation of the C type webkit_download_set_destination.
func (recv *Download) SetDestination(uri string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(uri)

	err := downloadSetDestinationFunction_Set()
	if err == nil {
		downloadSetDestinationFunction.Invoke(inArgs[:], nil)
	}

	return
}

// DownloadStruct creates an uninitialised Download.
func DownloadStruct() *Download {
	err := downloadStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &Download{native: downloadStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeDownload)
	return structGo
}
func finalizeDownload(obj *Download) {
	downloadStruct.Free(obj.native)
}

var editorStateStruct *gi.Struct
var editorStateStruct_Once sync.Once

func editorStateStruct_Set() error {
	var err error
	editorStateStruct_Once.Do(func() {
		editorStateStruct, err = gi.StructNew("WebKit2", "EditorState")
	})
	return err
}

type EditorState struct {
	native uintptr
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'GObject.Object'

// FieldPriv returns the C field 'priv'.
func (recv *EditorState) FieldPriv() *EditorStatePrivate {
	argValue := gi.FieldGet(editorStateStruct, recv.native, "priv")
	value := &EditorStatePrivate{native: argValue.Pointer()}
	return value
}

// SetFieldPriv sets the value of the C field 'priv'.
func (recv *EditorState) SetFieldPriv(value *EditorStatePrivate) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(editorStateStruct, recv.native, "priv", argValue)
}

var editorStateGetTypingAttributesFunction *gi.Function
var editorStateGetTypingAttributesFunction_Once sync.Once

func editorStateGetTypingAttributesFunction_Set() error {
	var err error
	editorStateGetTypingAttributesFunction_Once.Do(func() {
		err = editorStateStruct_Set()
		if err != nil {
			return
		}
		editorStateGetTypingAttributesFunction, err = editorStateStruct.InvokerNew("get_typing_attributes")
	})
	return err
}

// GetTypingAttributes is a representation of the C type webkit_editor_state_get_typing_attributes.
func (recv *EditorState) GetTypingAttributes() uint32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := editorStateGetTypingAttributesFunction_Set()
	if err == nil {
		ret = editorStateGetTypingAttributesFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo
}

var editorStateIsCopyAvailableFunction *gi.Function
var editorStateIsCopyAvailableFunction_Once sync.Once

func editorStateIsCopyAvailableFunction_Set() error {
	var err error
	editorStateIsCopyAvailableFunction_Once.Do(func() {
		err = editorStateStruct_Set()
		if err != nil {
			return
		}
		editorStateIsCopyAvailableFunction, err = editorStateStruct.InvokerNew("is_copy_available")
	})
	return err
}

// IsCopyAvailable is a representation of the C type webkit_editor_state_is_copy_available.
func (recv *EditorState) IsCopyAvailable() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := editorStateIsCopyAvailableFunction_Set()
	if err == nil {
		ret = editorStateIsCopyAvailableFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var editorStateIsCutAvailableFunction *gi.Function
var editorStateIsCutAvailableFunction_Once sync.Once

func editorStateIsCutAvailableFunction_Set() error {
	var err error
	editorStateIsCutAvailableFunction_Once.Do(func() {
		err = editorStateStruct_Set()
		if err != nil {
			return
		}
		editorStateIsCutAvailableFunction, err = editorStateStruct.InvokerNew("is_cut_available")
	})
	return err
}

// IsCutAvailable is a representation of the C type webkit_editor_state_is_cut_available.
func (recv *EditorState) IsCutAvailable() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := editorStateIsCutAvailableFunction_Set()
	if err == nil {
		ret = editorStateIsCutAvailableFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var editorStateIsPasteAvailableFunction *gi.Function
var editorStateIsPasteAvailableFunction_Once sync.Once

func editorStateIsPasteAvailableFunction_Set() error {
	var err error
	editorStateIsPasteAvailableFunction_Once.Do(func() {
		err = editorStateStruct_Set()
		if err != nil {
			return
		}
		editorStateIsPasteAvailableFunction, err = editorStateStruct.InvokerNew("is_paste_available")
	})
	return err
}

// IsPasteAvailable is a representation of the C type webkit_editor_state_is_paste_available.
func (recv *EditorState) IsPasteAvailable() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := editorStateIsPasteAvailableFunction_Set()
	if err == nil {
		ret = editorStateIsPasteAvailableFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var editorStateIsRedoAvailableFunction *gi.Function
var editorStateIsRedoAvailableFunction_Once sync.Once

func editorStateIsRedoAvailableFunction_Set() error {
	var err error
	editorStateIsRedoAvailableFunction_Once.Do(func() {
		err = editorStateStruct_Set()
		if err != nil {
			return
		}
		editorStateIsRedoAvailableFunction, err = editorStateStruct.InvokerNew("is_redo_available")
	})
	return err
}

// IsRedoAvailable is a representation of the C type webkit_editor_state_is_redo_available.
func (recv *EditorState) IsRedoAvailable() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := editorStateIsRedoAvailableFunction_Set()
	if err == nil {
		ret = editorStateIsRedoAvailableFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var editorStateIsUndoAvailableFunction *gi.Function
var editorStateIsUndoAvailableFunction_Once sync.Once

func editorStateIsUndoAvailableFunction_Set() error {
	var err error
	editorStateIsUndoAvailableFunction_Once.Do(func() {
		err = editorStateStruct_Set()
		if err != nil {
			return
		}
		editorStateIsUndoAvailableFunction, err = editorStateStruct.InvokerNew("is_undo_available")
	})
	return err
}

// IsUndoAvailable is a representation of the C type webkit_editor_state_is_undo_available.
func (recv *EditorState) IsUndoAvailable() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := editorStateIsUndoAvailableFunction_Set()
	if err == nil {
		ret = editorStateIsUndoAvailableFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// EditorStateStruct creates an uninitialised EditorState.
func EditorStateStruct() *EditorState {
	err := editorStateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &EditorState{native: editorStateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeEditorState)
	return structGo
}
func finalizeEditorState(obj *EditorState) {
	editorStateStruct.Free(obj.native)
}

var faviconDatabaseStruct *gi.Struct
var faviconDatabaseStruct_Once sync.Once

func faviconDatabaseStruct_Set() error {
	var err error
	faviconDatabaseStruct_Once.Do(func() {
		faviconDatabaseStruct, err = gi.StructNew("WebKit2", "FaviconDatabase")
	})
	return err
}

type FaviconDatabase struct {
	native uintptr
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'GObject.Object'

// FieldPriv returns the C field 'priv'.
func (recv *FaviconDatabase) FieldPriv() *FaviconDatabasePrivate {
	argValue := gi.FieldGet(faviconDatabaseStruct, recv.native, "priv")
	value := &FaviconDatabasePrivate{native: argValue.Pointer()}
	return value
}

// SetFieldPriv sets the value of the C field 'priv'.
func (recv *FaviconDatabase) SetFieldPriv(value *FaviconDatabasePrivate) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(faviconDatabaseStruct, recv.native, "priv", argValue)
}

var faviconDatabaseClearFunction *gi.Function
var faviconDatabaseClearFunction_Once sync.Once

func faviconDatabaseClearFunction_Set() error {
	var err error
	faviconDatabaseClearFunction_Once.Do(func() {
		err = faviconDatabaseStruct_Set()
		if err != nil {
			return
		}
		faviconDatabaseClearFunction, err = faviconDatabaseStruct.InvokerNew("clear")
	})
	return err
}

// Clear is a representation of the C type webkit_favicon_database_clear.
func (recv *FaviconDatabase) Clear() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := faviconDatabaseClearFunction_Set()
	if err == nil {
		faviconDatabaseClearFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'webkit_favicon_database_get_favicon' : parameter 'cancellable' of type 'Gio.Cancellable' not supported

// UNSUPPORTED : C value 'webkit_favicon_database_get_favicon_finish' : parameter 'result' of type 'Gio.AsyncResult' not supported

var faviconDatabaseGetFaviconUriFunction *gi.Function
var faviconDatabaseGetFaviconUriFunction_Once sync.Once

func faviconDatabaseGetFaviconUriFunction_Set() error {
	var err error
	faviconDatabaseGetFaviconUriFunction_Once.Do(func() {
		err = faviconDatabaseStruct_Set()
		if err != nil {
			return
		}
		faviconDatabaseGetFaviconUriFunction, err = faviconDatabaseStruct.InvokerNew("get_favicon_uri")
	})
	return err
}

// GetFaviconUri is a representation of the C type webkit_favicon_database_get_favicon_uri.
func (recv *FaviconDatabase) GetFaviconUri(pageUri string) string {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(pageUri)

	var ret gi.Argument

	err := faviconDatabaseGetFaviconUriFunction_Set()
	if err == nil {
		ret = faviconDatabaseGetFaviconUriFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

// FaviconDatabaseStruct creates an uninitialised FaviconDatabase.
func FaviconDatabaseStruct() *FaviconDatabase {
	err := faviconDatabaseStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &FaviconDatabase{native: faviconDatabaseStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeFaviconDatabase)
	return structGo
}
func finalizeFaviconDatabase(obj *FaviconDatabase) {
	faviconDatabaseStruct.Free(obj.native)
}

var fileChooserRequestStruct *gi.Struct
var fileChooserRequestStruct_Once sync.Once

func fileChooserRequestStruct_Set() error {
	var err error
	fileChooserRequestStruct_Once.Do(func() {
		fileChooserRequestStruct, err = gi.StructNew("WebKit2", "FileChooserRequest")
	})
	return err
}

type FileChooserRequest struct {
	native uintptr
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'GObject.Object'

var fileChooserRequestCancelFunction *gi.Function
var fileChooserRequestCancelFunction_Once sync.Once

func fileChooserRequestCancelFunction_Set() error {
	var err error
	fileChooserRequestCancelFunction_Once.Do(func() {
		err = fileChooserRequestStruct_Set()
		if err != nil {
			return
		}
		fileChooserRequestCancelFunction, err = fileChooserRequestStruct.InvokerNew("cancel")
	})
	return err
}

// Cancel is a representation of the C type webkit_file_chooser_request_cancel.
func (recv *FileChooserRequest) Cancel() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := fileChooserRequestCancelFunction_Set()
	if err == nil {
		fileChooserRequestCancelFunction.Invoke(inArgs[:], nil)
	}

	return
}

var fileChooserRequestGetMimeTypesFunction *gi.Function
var fileChooserRequestGetMimeTypesFunction_Once sync.Once

func fileChooserRequestGetMimeTypesFunction_Set() error {
	var err error
	fileChooserRequestGetMimeTypesFunction_Once.Do(func() {
		err = fileChooserRequestStruct_Set()
		if err != nil {
			return
		}
		fileChooserRequestGetMimeTypesFunction, err = fileChooserRequestStruct.InvokerNew("get_mime_types")
	})
	return err
}

// GetMimeTypes is a representation of the C type webkit_file_chooser_request_get_mime_types.
func (recv *FileChooserRequest) GetMimeTypes() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := fileChooserRequestGetMimeTypesFunction_Set()
	if err == nil {
		fileChooserRequestGetMimeTypesFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'webkit_file_chooser_request_get_mime_types_filter' : return type 'Gtk.FileFilter' not supported

var fileChooserRequestGetSelectMultipleFunction *gi.Function
var fileChooserRequestGetSelectMultipleFunction_Once sync.Once

func fileChooserRequestGetSelectMultipleFunction_Set() error {
	var err error
	fileChooserRequestGetSelectMultipleFunction_Once.Do(func() {
		err = fileChooserRequestStruct_Set()
		if err != nil {
			return
		}
		fileChooserRequestGetSelectMultipleFunction, err = fileChooserRequestStruct.InvokerNew("get_select_multiple")
	})
	return err
}

// GetSelectMultiple is a representation of the C type webkit_file_chooser_request_get_select_multiple.
func (recv *FileChooserRequest) GetSelectMultiple() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := fileChooserRequestGetSelectMultipleFunction_Set()
	if err == nil {
		ret = fileChooserRequestGetSelectMultipleFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var fileChooserRequestGetSelectedFilesFunction *gi.Function
var fileChooserRequestGetSelectedFilesFunction_Once sync.Once

func fileChooserRequestGetSelectedFilesFunction_Set() error {
	var err error
	fileChooserRequestGetSelectedFilesFunction_Once.Do(func() {
		err = fileChooserRequestStruct_Set()
		if err != nil {
			return
		}
		fileChooserRequestGetSelectedFilesFunction, err = fileChooserRequestStruct.InvokerNew("get_selected_files")
	})
	return err
}

// GetSelectedFiles is a representation of the C type webkit_file_chooser_request_get_selected_files.
func (recv *FileChooserRequest) GetSelectedFiles() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := fileChooserRequestGetSelectedFilesFunction_Set()
	if err == nil {
		fileChooserRequestGetSelectedFilesFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'webkit_file_chooser_request_select_files' : parameter 'files' of type 'nil' not supported

// FileChooserRequestStruct creates an uninitialised FileChooserRequest.
func FileChooserRequestStruct() *FileChooserRequest {
	err := fileChooserRequestStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &FileChooserRequest{native: fileChooserRequestStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeFileChooserRequest)
	return structGo
}
func finalizeFileChooserRequest(obj *FileChooserRequest) {
	fileChooserRequestStruct.Free(obj.native)
}

var findControllerStruct *gi.Struct
var findControllerStruct_Once sync.Once

func findControllerStruct_Set() error {
	var err error
	findControllerStruct_Once.Do(func() {
		findControllerStruct, err = gi.StructNew("WebKit2", "FindController")
	})
	return err
}

type FindController struct {
	native uintptr
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'GObject.Object'

var findControllerCountMatchesFunction *gi.Function
var findControllerCountMatchesFunction_Once sync.Once

func findControllerCountMatchesFunction_Set() error {
	var err error
	findControllerCountMatchesFunction_Once.Do(func() {
		err = findControllerStruct_Set()
		if err != nil {
			return
		}
		findControllerCountMatchesFunction, err = findControllerStruct.InvokerNew("count_matches")
	})
	return err
}

// CountMatches is a representation of the C type webkit_find_controller_count_matches.
func (recv *FindController) CountMatches(searchText string, findOptions uint32, maxMatchCount uint32) {
	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(searchText)
	inArgs[2].SetUint32(findOptions)
	inArgs[3].SetUint32(maxMatchCount)

	err := findControllerCountMatchesFunction_Set()
	if err == nil {
		findControllerCountMatchesFunction.Invoke(inArgs[:], nil)
	}

	return
}

var findControllerGetMaxMatchCountFunction *gi.Function
var findControllerGetMaxMatchCountFunction_Once sync.Once

func findControllerGetMaxMatchCountFunction_Set() error {
	var err error
	findControllerGetMaxMatchCountFunction_Once.Do(func() {
		err = findControllerStruct_Set()
		if err != nil {
			return
		}
		findControllerGetMaxMatchCountFunction, err = findControllerStruct.InvokerNew("get_max_match_count")
	})
	return err
}

// GetMaxMatchCount is a representation of the C type webkit_find_controller_get_max_match_count.
func (recv *FindController) GetMaxMatchCount() uint32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := findControllerGetMaxMatchCountFunction_Set()
	if err == nil {
		ret = findControllerGetMaxMatchCountFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo
}

var findControllerGetOptionsFunction *gi.Function
var findControllerGetOptionsFunction_Once sync.Once

func findControllerGetOptionsFunction_Set() error {
	var err error
	findControllerGetOptionsFunction_Once.Do(func() {
		err = findControllerStruct_Set()
		if err != nil {
			return
		}
		findControllerGetOptionsFunction, err = findControllerStruct.InvokerNew("get_options")
	})
	return err
}

// GetOptions is a representation of the C type webkit_find_controller_get_options.
func (recv *FindController) GetOptions() uint32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := findControllerGetOptionsFunction_Set()
	if err == nil {
		ret = findControllerGetOptionsFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo
}

var findControllerGetSearchTextFunction *gi.Function
var findControllerGetSearchTextFunction_Once sync.Once

func findControllerGetSearchTextFunction_Set() error {
	var err error
	findControllerGetSearchTextFunction_Once.Do(func() {
		err = findControllerStruct_Set()
		if err != nil {
			return
		}
		findControllerGetSearchTextFunction, err = findControllerStruct.InvokerNew("get_search_text")
	})
	return err
}

// GetSearchText is a representation of the C type webkit_find_controller_get_search_text.
func (recv *FindController) GetSearchText() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := findControllerGetSearchTextFunction_Set()
	if err == nil {
		ret = findControllerGetSearchTextFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'webkit_find_controller_get_web_view' : return type 'WebView' not supported

var findControllerSearchFunction *gi.Function
var findControllerSearchFunction_Once sync.Once

func findControllerSearchFunction_Set() error {
	var err error
	findControllerSearchFunction_Once.Do(func() {
		err = findControllerStruct_Set()
		if err != nil {
			return
		}
		findControllerSearchFunction, err = findControllerStruct.InvokerNew("search")
	})
	return err
}

// Search is a representation of the C type webkit_find_controller_search.
func (recv *FindController) Search(searchText string, findOptions uint32, maxMatchCount uint32) {
	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(searchText)
	inArgs[2].SetUint32(findOptions)
	inArgs[3].SetUint32(maxMatchCount)

	err := findControllerSearchFunction_Set()
	if err == nil {
		findControllerSearchFunction.Invoke(inArgs[:], nil)
	}

	return
}

var findControllerSearchFinishFunction *gi.Function
var findControllerSearchFinishFunction_Once sync.Once

func findControllerSearchFinishFunction_Set() error {
	var err error
	findControllerSearchFinishFunction_Once.Do(func() {
		err = findControllerStruct_Set()
		if err != nil {
			return
		}
		findControllerSearchFinishFunction, err = findControllerStruct.InvokerNew("search_finish")
	})
	return err
}

// SearchFinish is a representation of the C type webkit_find_controller_search_finish.
func (recv *FindController) SearchFinish() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := findControllerSearchFinishFunction_Set()
	if err == nil {
		findControllerSearchFinishFunction.Invoke(inArgs[:], nil)
	}

	return
}

var findControllerSearchNextFunction *gi.Function
var findControllerSearchNextFunction_Once sync.Once

func findControllerSearchNextFunction_Set() error {
	var err error
	findControllerSearchNextFunction_Once.Do(func() {
		err = findControllerStruct_Set()
		if err != nil {
			return
		}
		findControllerSearchNextFunction, err = findControllerStruct.InvokerNew("search_next")
	})
	return err
}

// SearchNext is a representation of the C type webkit_find_controller_search_next.
func (recv *FindController) SearchNext() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := findControllerSearchNextFunction_Set()
	if err == nil {
		findControllerSearchNextFunction.Invoke(inArgs[:], nil)
	}

	return
}

var findControllerSearchPreviousFunction *gi.Function
var findControllerSearchPreviousFunction_Once sync.Once

func findControllerSearchPreviousFunction_Set() error {
	var err error
	findControllerSearchPreviousFunction_Once.Do(func() {
		err = findControllerStruct_Set()
		if err != nil {
			return
		}
		findControllerSearchPreviousFunction, err = findControllerStruct.InvokerNew("search_previous")
	})
	return err
}

// SearchPrevious is a representation of the C type webkit_find_controller_search_previous.
func (recv *FindController) SearchPrevious() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := findControllerSearchPreviousFunction_Set()
	if err == nil {
		findControllerSearchPreviousFunction.Invoke(inArgs[:], nil)
	}

	return
}

// FindControllerStruct creates an uninitialised FindController.
func FindControllerStruct() *FindController {
	err := findControllerStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &FindController{native: findControllerStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeFindController)
	return structGo
}
func finalizeFindController(obj *FindController) {
	findControllerStruct.Free(obj.native)
}

var formSubmissionRequestStruct *gi.Struct
var formSubmissionRequestStruct_Once sync.Once

func formSubmissionRequestStruct_Set() error {
	var err error
	formSubmissionRequestStruct_Once.Do(func() {
		formSubmissionRequestStruct, err = gi.StructNew("WebKit2", "FormSubmissionRequest")
	})
	return err
}

type FormSubmissionRequest struct {
	native uintptr
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'webkit_form_submission_request_get_text_fields' : return type 'GLib.HashTable' not supported

// UNSUPPORTED : C value 'webkit_form_submission_request_list_text_fields' : parameter 'field_names' of type 'nil' not supported

var formSubmissionRequestSubmitFunction *gi.Function
var formSubmissionRequestSubmitFunction_Once sync.Once

func formSubmissionRequestSubmitFunction_Set() error {
	var err error
	formSubmissionRequestSubmitFunction_Once.Do(func() {
		err = formSubmissionRequestStruct_Set()
		if err != nil {
			return
		}
		formSubmissionRequestSubmitFunction, err = formSubmissionRequestStruct.InvokerNew("submit")
	})
	return err
}

// Submit is a representation of the C type webkit_form_submission_request_submit.
func (recv *FormSubmissionRequest) Submit() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := formSubmissionRequestSubmitFunction_Set()
	if err == nil {
		formSubmissionRequestSubmitFunction.Invoke(inArgs[:], nil)
	}

	return
}

// FormSubmissionRequestStruct creates an uninitialised FormSubmissionRequest.
func FormSubmissionRequestStruct() *FormSubmissionRequest {
	err := formSubmissionRequestStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &FormSubmissionRequest{native: formSubmissionRequestStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeFormSubmissionRequest)
	return structGo
}
func finalizeFormSubmissionRequest(obj *FormSubmissionRequest) {
	formSubmissionRequestStruct.Free(obj.native)
}

var geolocationManagerStruct *gi.Struct
var geolocationManagerStruct_Once sync.Once

func geolocationManagerStruct_Set() error {
	var err error
	geolocationManagerStruct_Once.Do(func() {
		geolocationManagerStruct, err = gi.StructNew("WebKit2", "GeolocationManager")
	})
	return err
}

type GeolocationManager struct {
	native uintptr
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'GObject.Object'

var geolocationManagerFailedFunction *gi.Function
var geolocationManagerFailedFunction_Once sync.Once

func geolocationManagerFailedFunction_Set() error {
	var err error
	geolocationManagerFailedFunction_Once.Do(func() {
		err = geolocationManagerStruct_Set()
		if err != nil {
			return
		}
		geolocationManagerFailedFunction, err = geolocationManagerStruct.InvokerNew("failed")
	})
	return err
}

// Failed is a representation of the C type webkit_geolocation_manager_failed.
func (recv *GeolocationManager) Failed(errorMessage string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(errorMessage)

	err := geolocationManagerFailedFunction_Set()
	if err == nil {
		geolocationManagerFailedFunction.Invoke(inArgs[:], nil)
	}

	return
}

var geolocationManagerGetEnableHighAccuracyFunction *gi.Function
var geolocationManagerGetEnableHighAccuracyFunction_Once sync.Once

func geolocationManagerGetEnableHighAccuracyFunction_Set() error {
	var err error
	geolocationManagerGetEnableHighAccuracyFunction_Once.Do(func() {
		err = geolocationManagerStruct_Set()
		if err != nil {
			return
		}
		geolocationManagerGetEnableHighAccuracyFunction, err = geolocationManagerStruct.InvokerNew("get_enable_high_accuracy")
	})
	return err
}

// GetEnableHighAccuracy is a representation of the C type webkit_geolocation_manager_get_enable_high_accuracy.
func (recv *GeolocationManager) GetEnableHighAccuracy() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := geolocationManagerGetEnableHighAccuracyFunction_Set()
	if err == nil {
		ret = geolocationManagerGetEnableHighAccuracyFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var geolocationManagerUpdatePositionFunction *gi.Function
var geolocationManagerUpdatePositionFunction_Once sync.Once

func geolocationManagerUpdatePositionFunction_Set() error {
	var err error
	geolocationManagerUpdatePositionFunction_Once.Do(func() {
		err = geolocationManagerStruct_Set()
		if err != nil {
			return
		}
		geolocationManagerUpdatePositionFunction, err = geolocationManagerStruct.InvokerNew("update_position")
	})
	return err
}

// UpdatePosition is a representation of the C type webkit_geolocation_manager_update_position.
func (recv *GeolocationManager) UpdatePosition(position *GeolocationPosition) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(position.native)

	err := geolocationManagerUpdatePositionFunction_Set()
	if err == nil {
		geolocationManagerUpdatePositionFunction.Invoke(inArgs[:], nil)
	}

	return
}

// GeolocationManagerStruct creates an uninitialised GeolocationManager.
func GeolocationManagerStruct() *GeolocationManager {
	err := geolocationManagerStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &GeolocationManager{native: geolocationManagerStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeGeolocationManager)
	return structGo
}
func finalizeGeolocationManager(obj *GeolocationManager) {
	geolocationManagerStruct.Free(obj.native)
}

var geolocationPermissionRequestStruct *gi.Struct
var geolocationPermissionRequestStruct_Once sync.Once

func geolocationPermissionRequestStruct_Set() error {
	var err error
	geolocationPermissionRequestStruct_Once.Do(func() {
		geolocationPermissionRequestStruct, err = gi.StructNew("WebKit2", "GeolocationPermissionRequest")
	})
	return err
}

type GeolocationPermissionRequest struct {
	native uintptr
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'GObject.Object'

// GeolocationPermissionRequestStruct creates an uninitialised GeolocationPermissionRequest.
func GeolocationPermissionRequestStruct() *GeolocationPermissionRequest {
	err := geolocationPermissionRequestStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &GeolocationPermissionRequest{native: geolocationPermissionRequestStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeGeolocationPermissionRequest)
	return structGo
}
func finalizeGeolocationPermissionRequest(obj *GeolocationPermissionRequest) {
	geolocationPermissionRequestStruct.Free(obj.native)
}

var hitTestResultStruct *gi.Struct
var hitTestResultStruct_Once sync.Once

func hitTestResultStruct_Set() error {
	var err error
	hitTestResultStruct_Once.Do(func() {
		hitTestResultStruct, err = gi.StructNew("WebKit2", "HitTestResult")
	})
	return err
}

type HitTestResult struct {
	native uintptr
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'GObject.Object'

// FieldPriv returns the C field 'priv'.
func (recv *HitTestResult) FieldPriv() *HitTestResultPrivate {
	argValue := gi.FieldGet(hitTestResultStruct, recv.native, "priv")
	value := &HitTestResultPrivate{native: argValue.Pointer()}
	return value
}

// SetFieldPriv sets the value of the C field 'priv'.
func (recv *HitTestResult) SetFieldPriv(value *HitTestResultPrivate) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(hitTestResultStruct, recv.native, "priv", argValue)
}

var hitTestResultContextIsEditableFunction *gi.Function
var hitTestResultContextIsEditableFunction_Once sync.Once

func hitTestResultContextIsEditableFunction_Set() error {
	var err error
	hitTestResultContextIsEditableFunction_Once.Do(func() {
		err = hitTestResultStruct_Set()
		if err != nil {
			return
		}
		hitTestResultContextIsEditableFunction, err = hitTestResultStruct.InvokerNew("context_is_editable")
	})
	return err
}

// ContextIsEditable is a representation of the C type webkit_hit_test_result_context_is_editable.
func (recv *HitTestResult) ContextIsEditable() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := hitTestResultContextIsEditableFunction_Set()
	if err == nil {
		ret = hitTestResultContextIsEditableFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var hitTestResultContextIsImageFunction *gi.Function
var hitTestResultContextIsImageFunction_Once sync.Once

func hitTestResultContextIsImageFunction_Set() error {
	var err error
	hitTestResultContextIsImageFunction_Once.Do(func() {
		err = hitTestResultStruct_Set()
		if err != nil {
			return
		}
		hitTestResultContextIsImageFunction, err = hitTestResultStruct.InvokerNew("context_is_image")
	})
	return err
}

// ContextIsImage is a representation of the C type webkit_hit_test_result_context_is_image.
func (recv *HitTestResult) ContextIsImage() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := hitTestResultContextIsImageFunction_Set()
	if err == nil {
		ret = hitTestResultContextIsImageFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var hitTestResultContextIsLinkFunction *gi.Function
var hitTestResultContextIsLinkFunction_Once sync.Once

func hitTestResultContextIsLinkFunction_Set() error {
	var err error
	hitTestResultContextIsLinkFunction_Once.Do(func() {
		err = hitTestResultStruct_Set()
		if err != nil {
			return
		}
		hitTestResultContextIsLinkFunction, err = hitTestResultStruct.InvokerNew("context_is_link")
	})
	return err
}

// ContextIsLink is a representation of the C type webkit_hit_test_result_context_is_link.
func (recv *HitTestResult) ContextIsLink() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := hitTestResultContextIsLinkFunction_Set()
	if err == nil {
		ret = hitTestResultContextIsLinkFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var hitTestResultContextIsMediaFunction *gi.Function
var hitTestResultContextIsMediaFunction_Once sync.Once

func hitTestResultContextIsMediaFunction_Set() error {
	var err error
	hitTestResultContextIsMediaFunction_Once.Do(func() {
		err = hitTestResultStruct_Set()
		if err != nil {
			return
		}
		hitTestResultContextIsMediaFunction, err = hitTestResultStruct.InvokerNew("context_is_media")
	})
	return err
}

// ContextIsMedia is a representation of the C type webkit_hit_test_result_context_is_media.
func (recv *HitTestResult) ContextIsMedia() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := hitTestResultContextIsMediaFunction_Set()
	if err == nil {
		ret = hitTestResultContextIsMediaFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var hitTestResultContextIsScrollbarFunction *gi.Function
var hitTestResultContextIsScrollbarFunction_Once sync.Once

func hitTestResultContextIsScrollbarFunction_Set() error {
	var err error
	hitTestResultContextIsScrollbarFunction_Once.Do(func() {
		err = hitTestResultStruct_Set()
		if err != nil {
			return
		}
		hitTestResultContextIsScrollbarFunction, err = hitTestResultStruct.InvokerNew("context_is_scrollbar")
	})
	return err
}

// ContextIsScrollbar is a representation of the C type webkit_hit_test_result_context_is_scrollbar.
func (recv *HitTestResult) ContextIsScrollbar() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := hitTestResultContextIsScrollbarFunction_Set()
	if err == nil {
		ret = hitTestResultContextIsScrollbarFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var hitTestResultContextIsSelectionFunction *gi.Function
var hitTestResultContextIsSelectionFunction_Once sync.Once

func hitTestResultContextIsSelectionFunction_Set() error {
	var err error
	hitTestResultContextIsSelectionFunction_Once.Do(func() {
		err = hitTestResultStruct_Set()
		if err != nil {
			return
		}
		hitTestResultContextIsSelectionFunction, err = hitTestResultStruct.InvokerNew("context_is_selection")
	})
	return err
}

// ContextIsSelection is a representation of the C type webkit_hit_test_result_context_is_selection.
func (recv *HitTestResult) ContextIsSelection() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := hitTestResultContextIsSelectionFunction_Set()
	if err == nil {
		ret = hitTestResultContextIsSelectionFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var hitTestResultGetContextFunction *gi.Function
var hitTestResultGetContextFunction_Once sync.Once

func hitTestResultGetContextFunction_Set() error {
	var err error
	hitTestResultGetContextFunction_Once.Do(func() {
		err = hitTestResultStruct_Set()
		if err != nil {
			return
		}
		hitTestResultGetContextFunction, err = hitTestResultStruct.InvokerNew("get_context")
	})
	return err
}

// GetContext is a representation of the C type webkit_hit_test_result_get_context.
func (recv *HitTestResult) GetContext() uint32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := hitTestResultGetContextFunction_Set()
	if err == nil {
		ret = hitTestResultGetContextFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo
}

var hitTestResultGetImageUriFunction *gi.Function
var hitTestResultGetImageUriFunction_Once sync.Once

func hitTestResultGetImageUriFunction_Set() error {
	var err error
	hitTestResultGetImageUriFunction_Once.Do(func() {
		err = hitTestResultStruct_Set()
		if err != nil {
			return
		}
		hitTestResultGetImageUriFunction, err = hitTestResultStruct.InvokerNew("get_image_uri")
	})
	return err
}

// GetImageUri is a representation of the C type webkit_hit_test_result_get_image_uri.
func (recv *HitTestResult) GetImageUri() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := hitTestResultGetImageUriFunction_Set()
	if err == nil {
		ret = hitTestResultGetImageUriFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var hitTestResultGetLinkLabelFunction *gi.Function
var hitTestResultGetLinkLabelFunction_Once sync.Once

func hitTestResultGetLinkLabelFunction_Set() error {
	var err error
	hitTestResultGetLinkLabelFunction_Once.Do(func() {
		err = hitTestResultStruct_Set()
		if err != nil {
			return
		}
		hitTestResultGetLinkLabelFunction, err = hitTestResultStruct.InvokerNew("get_link_label")
	})
	return err
}

// GetLinkLabel is a representation of the C type webkit_hit_test_result_get_link_label.
func (recv *HitTestResult) GetLinkLabel() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := hitTestResultGetLinkLabelFunction_Set()
	if err == nil {
		ret = hitTestResultGetLinkLabelFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var hitTestResultGetLinkTitleFunction *gi.Function
var hitTestResultGetLinkTitleFunction_Once sync.Once

func hitTestResultGetLinkTitleFunction_Set() error {
	var err error
	hitTestResultGetLinkTitleFunction_Once.Do(func() {
		err = hitTestResultStruct_Set()
		if err != nil {
			return
		}
		hitTestResultGetLinkTitleFunction, err = hitTestResultStruct.InvokerNew("get_link_title")
	})
	return err
}

// GetLinkTitle is a representation of the C type webkit_hit_test_result_get_link_title.
func (recv *HitTestResult) GetLinkTitle() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := hitTestResultGetLinkTitleFunction_Set()
	if err == nil {
		ret = hitTestResultGetLinkTitleFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var hitTestResultGetLinkUriFunction *gi.Function
var hitTestResultGetLinkUriFunction_Once sync.Once

func hitTestResultGetLinkUriFunction_Set() error {
	var err error
	hitTestResultGetLinkUriFunction_Once.Do(func() {
		err = hitTestResultStruct_Set()
		if err != nil {
			return
		}
		hitTestResultGetLinkUriFunction, err = hitTestResultStruct.InvokerNew("get_link_uri")
	})
	return err
}

// GetLinkUri is a representation of the C type webkit_hit_test_result_get_link_uri.
func (recv *HitTestResult) GetLinkUri() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := hitTestResultGetLinkUriFunction_Set()
	if err == nil {
		ret = hitTestResultGetLinkUriFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var hitTestResultGetMediaUriFunction *gi.Function
var hitTestResultGetMediaUriFunction_Once sync.Once

func hitTestResultGetMediaUriFunction_Set() error {
	var err error
	hitTestResultGetMediaUriFunction_Once.Do(func() {
		err = hitTestResultStruct_Set()
		if err != nil {
			return
		}
		hitTestResultGetMediaUriFunction, err = hitTestResultStruct.InvokerNew("get_media_uri")
	})
	return err
}

// GetMediaUri is a representation of the C type webkit_hit_test_result_get_media_uri.
func (recv *HitTestResult) GetMediaUri() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := hitTestResultGetMediaUriFunction_Set()
	if err == nil {
		ret = hitTestResultGetMediaUriFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

// HitTestResultStruct creates an uninitialised HitTestResult.
func HitTestResultStruct() *HitTestResult {
	err := hitTestResultStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &HitTestResult{native: hitTestResultStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeHitTestResult)
	return structGo
}
func finalizeHitTestResult(obj *HitTestResult) {
	hitTestResultStruct.Free(obj.native)
}

var installMissingMediaPluginsPermissionRequestStruct *gi.Struct
var installMissingMediaPluginsPermissionRequestStruct_Once sync.Once

func installMissingMediaPluginsPermissionRequestStruct_Set() error {
	var err error
	installMissingMediaPluginsPermissionRequestStruct_Once.Do(func() {
		installMissingMediaPluginsPermissionRequestStruct, err = gi.StructNew("WebKit2", "InstallMissingMediaPluginsPermissionRequest")
	})
	return err
}

type InstallMissingMediaPluginsPermissionRequest struct {
	native uintptr
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'GObject.Object'

// FieldPriv returns the C field 'priv'.
func (recv *InstallMissingMediaPluginsPermissionRequest) FieldPriv() *InstallMissingMediaPluginsPermissionRequestPrivate {
	argValue := gi.FieldGet(installMissingMediaPluginsPermissionRequestStruct, recv.native, "priv")
	value := &InstallMissingMediaPluginsPermissionRequestPrivate{native: argValue.Pointer()}
	return value
}

// SetFieldPriv sets the value of the C field 'priv'.
func (recv *InstallMissingMediaPluginsPermissionRequest) SetFieldPriv(value *InstallMissingMediaPluginsPermissionRequestPrivate) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(installMissingMediaPluginsPermissionRequestStruct, recv.native, "priv", argValue)
}

var installMissingMediaPluginsPermissionRequestGetDescriptionFunction *gi.Function
var installMissingMediaPluginsPermissionRequestGetDescriptionFunction_Once sync.Once

func installMissingMediaPluginsPermissionRequestGetDescriptionFunction_Set() error {
	var err error
	installMissingMediaPluginsPermissionRequestGetDescriptionFunction_Once.Do(func() {
		err = installMissingMediaPluginsPermissionRequestStruct_Set()
		if err != nil {
			return
		}
		installMissingMediaPluginsPermissionRequestGetDescriptionFunction, err = installMissingMediaPluginsPermissionRequestStruct.InvokerNew("get_description")
	})
	return err
}

// GetDescription is a representation of the C type webkit_install_missing_media_plugins_permission_request_get_description.
func (recv *InstallMissingMediaPluginsPermissionRequest) GetDescription() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := installMissingMediaPluginsPermissionRequestGetDescriptionFunction_Set()
	if err == nil {
		ret = installMissingMediaPluginsPermissionRequestGetDescriptionFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

// InstallMissingMediaPluginsPermissionRequestStruct creates an uninitialised InstallMissingMediaPluginsPermissionRequest.
func InstallMissingMediaPluginsPermissionRequestStruct() *InstallMissingMediaPluginsPermissionRequest {
	err := installMissingMediaPluginsPermissionRequestStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &InstallMissingMediaPluginsPermissionRequest{native: installMissingMediaPluginsPermissionRequestStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeInstallMissingMediaPluginsPermissionRequest)
	return structGo
}
func finalizeInstallMissingMediaPluginsPermissionRequest(obj *InstallMissingMediaPluginsPermissionRequest) {
	installMissingMediaPluginsPermissionRequestStruct.Free(obj.native)
}

var navigationPolicyDecisionStruct *gi.Struct
var navigationPolicyDecisionStruct_Once sync.Once

func navigationPolicyDecisionStruct_Set() error {
	var err error
	navigationPolicyDecisionStruct_Once.Do(func() {
		navigationPolicyDecisionStruct, err = gi.StructNew("WebKit2", "NavigationPolicyDecision")
	})
	return err
}

type NavigationPolicyDecision struct {
	native uintptr
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'PolicyDecision'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'PolicyDecision'

var navigationPolicyDecisionGetFrameNameFunction *gi.Function
var navigationPolicyDecisionGetFrameNameFunction_Once sync.Once

func navigationPolicyDecisionGetFrameNameFunction_Set() error {
	var err error
	navigationPolicyDecisionGetFrameNameFunction_Once.Do(func() {
		err = navigationPolicyDecisionStruct_Set()
		if err != nil {
			return
		}
		navigationPolicyDecisionGetFrameNameFunction, err = navigationPolicyDecisionStruct.InvokerNew("get_frame_name")
	})
	return err
}

// GetFrameName is a representation of the C type webkit_navigation_policy_decision_get_frame_name.
func (recv *NavigationPolicyDecision) GetFrameName() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := navigationPolicyDecisionGetFrameNameFunction_Set()
	if err == nil {
		ret = navigationPolicyDecisionGetFrameNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var navigationPolicyDecisionGetModifiersFunction *gi.Function
var navigationPolicyDecisionGetModifiersFunction_Once sync.Once

func navigationPolicyDecisionGetModifiersFunction_Set() error {
	var err error
	navigationPolicyDecisionGetModifiersFunction_Once.Do(func() {
		err = navigationPolicyDecisionStruct_Set()
		if err != nil {
			return
		}
		navigationPolicyDecisionGetModifiersFunction, err = navigationPolicyDecisionStruct.InvokerNew("get_modifiers")
	})
	return err
}

// GetModifiers is a representation of the C type webkit_navigation_policy_decision_get_modifiers.
func (recv *NavigationPolicyDecision) GetModifiers() uint32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := navigationPolicyDecisionGetModifiersFunction_Set()
	if err == nil {
		ret = navigationPolicyDecisionGetModifiersFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo
}

var navigationPolicyDecisionGetMouseButtonFunction *gi.Function
var navigationPolicyDecisionGetMouseButtonFunction_Once sync.Once

func navigationPolicyDecisionGetMouseButtonFunction_Set() error {
	var err error
	navigationPolicyDecisionGetMouseButtonFunction_Once.Do(func() {
		err = navigationPolicyDecisionStruct_Set()
		if err != nil {
			return
		}
		navigationPolicyDecisionGetMouseButtonFunction, err = navigationPolicyDecisionStruct.InvokerNew("get_mouse_button")
	})
	return err
}

// GetMouseButton is a representation of the C type webkit_navigation_policy_decision_get_mouse_button.
func (recv *NavigationPolicyDecision) GetMouseButton() uint32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := navigationPolicyDecisionGetMouseButtonFunction_Set()
	if err == nil {
		ret = navigationPolicyDecisionGetMouseButtonFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo
}

var navigationPolicyDecisionGetNavigationActionFunction *gi.Function
var navigationPolicyDecisionGetNavigationActionFunction_Once sync.Once

func navigationPolicyDecisionGetNavigationActionFunction_Set() error {
	var err error
	navigationPolicyDecisionGetNavigationActionFunction_Once.Do(func() {
		err = navigationPolicyDecisionStruct_Set()
		if err != nil {
			return
		}
		navigationPolicyDecisionGetNavigationActionFunction, err = navigationPolicyDecisionStruct.InvokerNew("get_navigation_action")
	})
	return err
}

// GetNavigationAction is a representation of the C type webkit_navigation_policy_decision_get_navigation_action.
func (recv *NavigationPolicyDecision) GetNavigationAction() *NavigationAction {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := navigationPolicyDecisionGetNavigationActionFunction_Set()
	if err == nil {
		ret = navigationPolicyDecisionGetNavigationActionFunction.Invoke(inArgs[:], nil)
	}

	retGo := &NavigationAction{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'webkit_navigation_policy_decision_get_navigation_type' : return type 'NavigationType' not supported

// UNSUPPORTED : C value 'webkit_navigation_policy_decision_get_request' : return type 'URIRequest' not supported

// NavigationPolicyDecisionStruct creates an uninitialised NavigationPolicyDecision.
func NavigationPolicyDecisionStruct() *NavigationPolicyDecision {
	err := navigationPolicyDecisionStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &NavigationPolicyDecision{native: navigationPolicyDecisionStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeNavigationPolicyDecision)
	return structGo
}
func finalizeNavigationPolicyDecision(obj *NavigationPolicyDecision) {
	navigationPolicyDecisionStruct.Free(obj.native)
}

var notificationStruct *gi.Struct
var notificationStruct_Once sync.Once

func notificationStruct_Set() error {
	var err error
	notificationStruct_Once.Do(func() {
		notificationStruct, err = gi.StructNew("WebKit2", "Notification")
	})
	return err
}

type Notification struct {
	native uintptr
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'GObject.Object'

// FieldPriv returns the C field 'priv'.
func (recv *Notification) FieldPriv() *NotificationPrivate {
	argValue := gi.FieldGet(notificationStruct, recv.native, "priv")
	value := &NotificationPrivate{native: argValue.Pointer()}
	return value
}

// SetFieldPriv sets the value of the C field 'priv'.
func (recv *Notification) SetFieldPriv(value *NotificationPrivate) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(notificationStruct, recv.native, "priv", argValue)
}

var notificationClickedFunction *gi.Function
var notificationClickedFunction_Once sync.Once

func notificationClickedFunction_Set() error {
	var err error
	notificationClickedFunction_Once.Do(func() {
		err = notificationStruct_Set()
		if err != nil {
			return
		}
		notificationClickedFunction, err = notificationStruct.InvokerNew("clicked")
	})
	return err
}

// Clicked is a representation of the C type webkit_notification_clicked.
func (recv *Notification) Clicked() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := notificationClickedFunction_Set()
	if err == nil {
		notificationClickedFunction.Invoke(inArgs[:], nil)
	}

	return
}

var notificationCloseFunction *gi.Function
var notificationCloseFunction_Once sync.Once

func notificationCloseFunction_Set() error {
	var err error
	notificationCloseFunction_Once.Do(func() {
		err = notificationStruct_Set()
		if err != nil {
			return
		}
		notificationCloseFunction, err = notificationStruct.InvokerNew("close")
	})
	return err
}

// Close is a representation of the C type webkit_notification_close.
func (recv *Notification) Close() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := notificationCloseFunction_Set()
	if err == nil {
		notificationCloseFunction.Invoke(inArgs[:], nil)
	}

	return
}

var notificationGetBodyFunction *gi.Function
var notificationGetBodyFunction_Once sync.Once

func notificationGetBodyFunction_Set() error {
	var err error
	notificationGetBodyFunction_Once.Do(func() {
		err = notificationStruct_Set()
		if err != nil {
			return
		}
		notificationGetBodyFunction, err = notificationStruct.InvokerNew("get_body")
	})
	return err
}

// GetBody is a representation of the C type webkit_notification_get_body.
func (recv *Notification) GetBody() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := notificationGetBodyFunction_Set()
	if err == nil {
		ret = notificationGetBodyFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var notificationGetIdFunction *gi.Function
var notificationGetIdFunction_Once sync.Once

func notificationGetIdFunction_Set() error {
	var err error
	notificationGetIdFunction_Once.Do(func() {
		err = notificationStruct_Set()
		if err != nil {
			return
		}
		notificationGetIdFunction, err = notificationStruct.InvokerNew("get_id")
	})
	return err
}

// GetId is a representation of the C type webkit_notification_get_id.
func (recv *Notification) GetId() uint64 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := notificationGetIdFunction_Set()
	if err == nil {
		ret = notificationGetIdFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint64()

	return retGo
}

var notificationGetTagFunction *gi.Function
var notificationGetTagFunction_Once sync.Once

func notificationGetTagFunction_Set() error {
	var err error
	notificationGetTagFunction_Once.Do(func() {
		err = notificationStruct_Set()
		if err != nil {
			return
		}
		notificationGetTagFunction, err = notificationStruct.InvokerNew("get_tag")
	})
	return err
}

// GetTag is a representation of the C type webkit_notification_get_tag.
func (recv *Notification) GetTag() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := notificationGetTagFunction_Set()
	if err == nil {
		ret = notificationGetTagFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var notificationGetTitleFunction *gi.Function
var notificationGetTitleFunction_Once sync.Once

func notificationGetTitleFunction_Set() error {
	var err error
	notificationGetTitleFunction_Once.Do(func() {
		err = notificationStruct_Set()
		if err != nil {
			return
		}
		notificationGetTitleFunction, err = notificationStruct.InvokerNew("get_title")
	})
	return err
}

// GetTitle is a representation of the C type webkit_notification_get_title.
func (recv *Notification) GetTitle() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := notificationGetTitleFunction_Set()
	if err == nil {
		ret = notificationGetTitleFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

// NotificationStruct creates an uninitialised Notification.
func NotificationStruct() *Notification {
	err := notificationStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &Notification{native: notificationStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeNotification)
	return structGo
}
func finalizeNotification(obj *Notification) {
	notificationStruct.Free(obj.native)
}

var notificationPermissionRequestStruct *gi.Struct
var notificationPermissionRequestStruct_Once sync.Once

func notificationPermissionRequestStruct_Set() error {
	var err error
	notificationPermissionRequestStruct_Once.Do(func() {
		notificationPermissionRequestStruct, err = gi.StructNew("WebKit2", "NotificationPermissionRequest")
	})
	return err
}

type NotificationPermissionRequest struct {
	native uintptr
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'GObject.Object'

// NotificationPermissionRequestStruct creates an uninitialised NotificationPermissionRequest.
func NotificationPermissionRequestStruct() *NotificationPermissionRequest {
	err := notificationPermissionRequestStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &NotificationPermissionRequest{native: notificationPermissionRequestStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeNotificationPermissionRequest)
	return structGo
}
func finalizeNotificationPermissionRequest(obj *NotificationPermissionRequest) {
	notificationPermissionRequestStruct.Free(obj.native)
}

var optionMenuStruct *gi.Struct
var optionMenuStruct_Once sync.Once

func optionMenuStruct_Set() error {
	var err error
	optionMenuStruct_Once.Do(func() {
		optionMenuStruct, err = gi.StructNew("WebKit2", "OptionMenu")
	})
	return err
}

type OptionMenu struct {
	native uintptr
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'GObject.Object'

// FieldPriv returns the C field 'priv'.
func (recv *OptionMenu) FieldPriv() *OptionMenuPrivate {
	argValue := gi.FieldGet(optionMenuStruct, recv.native, "priv")
	value := &OptionMenuPrivate{native: argValue.Pointer()}
	return value
}

// SetFieldPriv sets the value of the C field 'priv'.
func (recv *OptionMenu) SetFieldPriv(value *OptionMenuPrivate) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(optionMenuStruct, recv.native, "priv", argValue)
}

var optionMenuActivateItemFunction *gi.Function
var optionMenuActivateItemFunction_Once sync.Once

func optionMenuActivateItemFunction_Set() error {
	var err error
	optionMenuActivateItemFunction_Once.Do(func() {
		err = optionMenuStruct_Set()
		if err != nil {
			return
		}
		optionMenuActivateItemFunction, err = optionMenuStruct.InvokerNew("activate_item")
	})
	return err
}

// ActivateItem is a representation of the C type webkit_option_menu_activate_item.
func (recv *OptionMenu) ActivateItem(index uint32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetUint32(index)

	err := optionMenuActivateItemFunction_Set()
	if err == nil {
		optionMenuActivateItemFunction.Invoke(inArgs[:], nil)
	}

	return
}

var optionMenuCloseFunction *gi.Function
var optionMenuCloseFunction_Once sync.Once

func optionMenuCloseFunction_Set() error {
	var err error
	optionMenuCloseFunction_Once.Do(func() {
		err = optionMenuStruct_Set()
		if err != nil {
			return
		}
		optionMenuCloseFunction, err = optionMenuStruct.InvokerNew("close")
	})
	return err
}

// Close is a representation of the C type webkit_option_menu_close.
func (recv *OptionMenu) Close() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := optionMenuCloseFunction_Set()
	if err == nil {
		optionMenuCloseFunction.Invoke(inArgs[:], nil)
	}

	return
}

var optionMenuGetItemFunction *gi.Function
var optionMenuGetItemFunction_Once sync.Once

func optionMenuGetItemFunction_Set() error {
	var err error
	optionMenuGetItemFunction_Once.Do(func() {
		err = optionMenuStruct_Set()
		if err != nil {
			return
		}
		optionMenuGetItemFunction, err = optionMenuStruct.InvokerNew("get_item")
	})
	return err
}

// GetItem is a representation of the C type webkit_option_menu_get_item.
func (recv *OptionMenu) GetItem(index uint32) *OptionMenuItem {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetUint32(index)

	var ret gi.Argument

	err := optionMenuGetItemFunction_Set()
	if err == nil {
		ret = optionMenuGetItemFunction.Invoke(inArgs[:], nil)
	}

	retGo := &OptionMenuItem{native: ret.Pointer()}

	return retGo
}

var optionMenuGetNItemsFunction *gi.Function
var optionMenuGetNItemsFunction_Once sync.Once

func optionMenuGetNItemsFunction_Set() error {
	var err error
	optionMenuGetNItemsFunction_Once.Do(func() {
		err = optionMenuStruct_Set()
		if err != nil {
			return
		}
		optionMenuGetNItemsFunction, err = optionMenuStruct.InvokerNew("get_n_items")
	})
	return err
}

// GetNItems is a representation of the C type webkit_option_menu_get_n_items.
func (recv *OptionMenu) GetNItems() uint32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := optionMenuGetNItemsFunction_Set()
	if err == nil {
		ret = optionMenuGetNItemsFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo
}

var optionMenuSelectItemFunction *gi.Function
var optionMenuSelectItemFunction_Once sync.Once

func optionMenuSelectItemFunction_Set() error {
	var err error
	optionMenuSelectItemFunction_Once.Do(func() {
		err = optionMenuStruct_Set()
		if err != nil {
			return
		}
		optionMenuSelectItemFunction, err = optionMenuStruct.InvokerNew("select_item")
	})
	return err
}

// SelectItem is a representation of the C type webkit_option_menu_select_item.
func (recv *OptionMenu) SelectItem(index uint32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetUint32(index)

	err := optionMenuSelectItemFunction_Set()
	if err == nil {
		optionMenuSelectItemFunction.Invoke(inArgs[:], nil)
	}

	return
}

// OptionMenuStruct creates an uninitialised OptionMenu.
func OptionMenuStruct() *OptionMenu {
	err := optionMenuStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &OptionMenu{native: optionMenuStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeOptionMenu)
	return structGo
}
func finalizeOptionMenu(obj *OptionMenu) {
	optionMenuStruct.Free(obj.native)
}

var pluginStruct *gi.Struct
var pluginStruct_Once sync.Once

func pluginStruct_Set() error {
	var err error
	pluginStruct_Once.Do(func() {
		pluginStruct, err = gi.StructNew("WebKit2", "Plugin")
	})
	return err
}

type Plugin struct {
	native uintptr
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'GObject.Object'

// FieldPriv returns the C field 'priv'.
func (recv *Plugin) FieldPriv() *PluginPrivate {
	argValue := gi.FieldGet(pluginStruct, recv.native, "priv")
	value := &PluginPrivate{native: argValue.Pointer()}
	return value
}

// SetFieldPriv sets the value of the C field 'priv'.
func (recv *Plugin) SetFieldPriv(value *PluginPrivate) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(pluginStruct, recv.native, "priv", argValue)
}

var pluginGetDescriptionFunction *gi.Function
var pluginGetDescriptionFunction_Once sync.Once

func pluginGetDescriptionFunction_Set() error {
	var err error
	pluginGetDescriptionFunction_Once.Do(func() {
		err = pluginStruct_Set()
		if err != nil {
			return
		}
		pluginGetDescriptionFunction, err = pluginStruct.InvokerNew("get_description")
	})
	return err
}

// GetDescription is a representation of the C type webkit_plugin_get_description.
func (recv *Plugin) GetDescription() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := pluginGetDescriptionFunction_Set()
	if err == nil {
		ret = pluginGetDescriptionFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'webkit_plugin_get_mime_info_list' : return type 'GLib.List' not supported

var pluginGetNameFunction *gi.Function
var pluginGetNameFunction_Once sync.Once

func pluginGetNameFunction_Set() error {
	var err error
	pluginGetNameFunction_Once.Do(func() {
		err = pluginStruct_Set()
		if err != nil {
			return
		}
		pluginGetNameFunction, err = pluginStruct.InvokerNew("get_name")
	})
	return err
}

// GetName is a representation of the C type webkit_plugin_get_name.
func (recv *Plugin) GetName() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := pluginGetNameFunction_Set()
	if err == nil {
		ret = pluginGetNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var pluginGetPathFunction *gi.Function
var pluginGetPathFunction_Once sync.Once

func pluginGetPathFunction_Set() error {
	var err error
	pluginGetPathFunction_Once.Do(func() {
		err = pluginStruct_Set()
		if err != nil {
			return
		}
		pluginGetPathFunction, err = pluginStruct.InvokerNew("get_path")
	})
	return err
}

// GetPath is a representation of the C type webkit_plugin_get_path.
func (recv *Plugin) GetPath() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := pluginGetPathFunction_Set()
	if err == nil {
		ret = pluginGetPathFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

// PluginStruct creates an uninitialised Plugin.
func PluginStruct() *Plugin {
	err := pluginStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &Plugin{native: pluginStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizePlugin)
	return structGo
}
func finalizePlugin(obj *Plugin) {
	pluginStruct.Free(obj.native)
}

var policyDecisionStruct *gi.Struct
var policyDecisionStruct_Once sync.Once

func policyDecisionStruct_Set() error {
	var err error
	policyDecisionStruct_Once.Do(func() {
		policyDecisionStruct, err = gi.StructNew("WebKit2", "PolicyDecision")
	})
	return err
}

type PolicyDecision struct {
	native uintptr
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'GObject.Object'

var policyDecisionDownloadFunction *gi.Function
var policyDecisionDownloadFunction_Once sync.Once

func policyDecisionDownloadFunction_Set() error {
	var err error
	policyDecisionDownloadFunction_Once.Do(func() {
		err = policyDecisionStruct_Set()
		if err != nil {
			return
		}
		policyDecisionDownloadFunction, err = policyDecisionStruct.InvokerNew("download")
	})
	return err
}

// Download is a representation of the C type webkit_policy_decision_download.
func (recv *PolicyDecision) Download() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := policyDecisionDownloadFunction_Set()
	if err == nil {
		policyDecisionDownloadFunction.Invoke(inArgs[:], nil)
	}

	return
}

var policyDecisionIgnoreFunction *gi.Function
var policyDecisionIgnoreFunction_Once sync.Once

func policyDecisionIgnoreFunction_Set() error {
	var err error
	policyDecisionIgnoreFunction_Once.Do(func() {
		err = policyDecisionStruct_Set()
		if err != nil {
			return
		}
		policyDecisionIgnoreFunction, err = policyDecisionStruct.InvokerNew("ignore")
	})
	return err
}

// Ignore is a representation of the C type webkit_policy_decision_ignore.
func (recv *PolicyDecision) Ignore() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := policyDecisionIgnoreFunction_Set()
	if err == nil {
		policyDecisionIgnoreFunction.Invoke(inArgs[:], nil)
	}

	return
}

var policyDecisionUseFunction *gi.Function
var policyDecisionUseFunction_Once sync.Once

func policyDecisionUseFunction_Set() error {
	var err error
	policyDecisionUseFunction_Once.Do(func() {
		err = policyDecisionStruct_Set()
		if err != nil {
			return
		}
		policyDecisionUseFunction, err = policyDecisionStruct.InvokerNew("use")
	})
	return err
}

// Use is a representation of the C type webkit_policy_decision_use.
func (recv *PolicyDecision) Use() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := policyDecisionUseFunction_Set()
	if err == nil {
		policyDecisionUseFunction.Invoke(inArgs[:], nil)
	}

	return
}

// PolicyDecisionStruct creates an uninitialised PolicyDecision.
func PolicyDecisionStruct() *PolicyDecision {
	err := policyDecisionStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &PolicyDecision{native: policyDecisionStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizePolicyDecision)
	return structGo
}
func finalizePolicyDecision(obj *PolicyDecision) {
	policyDecisionStruct.Free(obj.native)
}

var printCustomWidgetStruct *gi.Struct
var printCustomWidgetStruct_Once sync.Once

func printCustomWidgetStruct_Set() error {
	var err error
	printCustomWidgetStruct_Once.Do(func() {
		printCustomWidgetStruct, err = gi.StructNew("WebKit2", "PrintCustomWidget")
	})
	return err
}

type PrintCustomWidget struct {
	native uintptr
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'GObject.Object'

// FieldPriv returns the C field 'priv'.
func (recv *PrintCustomWidget) FieldPriv() *PrintCustomWidgetPrivate {
	argValue := gi.FieldGet(printCustomWidgetStruct, recv.native, "priv")
	value := &PrintCustomWidgetPrivate{native: argValue.Pointer()}
	return value
}

// SetFieldPriv sets the value of the C field 'priv'.
func (recv *PrintCustomWidget) SetFieldPriv(value *PrintCustomWidgetPrivate) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(printCustomWidgetStruct, recv.native, "priv", argValue)
}

// UNSUPPORTED : C value 'webkit_print_custom_widget_new' : parameter 'widget' of type 'Gtk.Widget' not supported

var printCustomWidgetGetTitleFunction *gi.Function
var printCustomWidgetGetTitleFunction_Once sync.Once

func printCustomWidgetGetTitleFunction_Set() error {
	var err error
	printCustomWidgetGetTitleFunction_Once.Do(func() {
		err = printCustomWidgetStruct_Set()
		if err != nil {
			return
		}
		printCustomWidgetGetTitleFunction, err = printCustomWidgetStruct.InvokerNew("get_title")
	})
	return err
}

// GetTitle is a representation of the C type webkit_print_custom_widget_get_title.
func (recv *PrintCustomWidget) GetTitle() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := printCustomWidgetGetTitleFunction_Set()
	if err == nil {
		ret = printCustomWidgetGetTitleFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'webkit_print_custom_widget_get_widget' : return type 'Gtk.Widget' not supported

var printOperationStruct *gi.Struct
var printOperationStruct_Once sync.Once

func printOperationStruct_Set() error {
	var err error
	printOperationStruct_Once.Do(func() {
		printOperationStruct, err = gi.StructNew("WebKit2", "PrintOperation")
	})
	return err
}

type PrintOperation struct {
	native uintptr
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'GObject.Object'

// FieldPriv returns the C field 'priv'.
func (recv *PrintOperation) FieldPriv() *PrintOperationPrivate {
	argValue := gi.FieldGet(printOperationStruct, recv.native, "priv")
	value := &PrintOperationPrivate{native: argValue.Pointer()}
	return value
}

// SetFieldPriv sets the value of the C field 'priv'.
func (recv *PrintOperation) SetFieldPriv(value *PrintOperationPrivate) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(printOperationStruct, recv.native, "priv", argValue)
}

// UNSUPPORTED : C value 'webkit_print_operation_new' : parameter 'web_view' of type 'WebView' not supported

// UNSUPPORTED : C value 'webkit_print_operation_get_page_setup' : return type 'Gtk.PageSetup' not supported

// UNSUPPORTED : C value 'webkit_print_operation_get_print_settings' : return type 'Gtk.PrintSettings' not supported

var printOperationPrintFunction *gi.Function
var printOperationPrintFunction_Once sync.Once

func printOperationPrintFunction_Set() error {
	var err error
	printOperationPrintFunction_Once.Do(func() {
		err = printOperationStruct_Set()
		if err != nil {
			return
		}
		printOperationPrintFunction, err = printOperationStruct.InvokerNew("print")
	})
	return err
}

// Print is a representation of the C type webkit_print_operation_print.
func (recv *PrintOperation) Print() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := printOperationPrintFunction_Set()
	if err == nil {
		printOperationPrintFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'webkit_print_operation_run_dialog' : parameter 'parent' of type 'Gtk.Window' not supported

// UNSUPPORTED : C value 'webkit_print_operation_set_page_setup' : parameter 'page_setup' of type 'Gtk.PageSetup' not supported

// UNSUPPORTED : C value 'webkit_print_operation_set_print_settings' : parameter 'print_settings' of type 'Gtk.PrintSettings' not supported

var responsePolicyDecisionStruct *gi.Struct
var responsePolicyDecisionStruct_Once sync.Once

func responsePolicyDecisionStruct_Set() error {
	var err error
	responsePolicyDecisionStruct_Once.Do(func() {
		responsePolicyDecisionStruct, err = gi.StructNew("WebKit2", "ResponsePolicyDecision")
	})
	return err
}

type ResponsePolicyDecision struct {
	native uintptr
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'PolicyDecision'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'PolicyDecision'

// UNSUPPORTED : C value 'webkit_response_policy_decision_get_request' : return type 'URIRequest' not supported

// UNSUPPORTED : C value 'webkit_response_policy_decision_get_response' : return type 'URIResponse' not supported

var responsePolicyDecisionIsMimeTypeSupportedFunction *gi.Function
var responsePolicyDecisionIsMimeTypeSupportedFunction_Once sync.Once

func responsePolicyDecisionIsMimeTypeSupportedFunction_Set() error {
	var err error
	responsePolicyDecisionIsMimeTypeSupportedFunction_Once.Do(func() {
		err = responsePolicyDecisionStruct_Set()
		if err != nil {
			return
		}
		responsePolicyDecisionIsMimeTypeSupportedFunction, err = responsePolicyDecisionStruct.InvokerNew("is_mime_type_supported")
	})
	return err
}

// IsMimeTypeSupported is a representation of the C type webkit_response_policy_decision_is_mime_type_supported.
func (recv *ResponsePolicyDecision) IsMimeTypeSupported() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := responsePolicyDecisionIsMimeTypeSupportedFunction_Set()
	if err == nil {
		ret = responsePolicyDecisionIsMimeTypeSupportedFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// ResponsePolicyDecisionStruct creates an uninitialised ResponsePolicyDecision.
func ResponsePolicyDecisionStruct() *ResponsePolicyDecision {
	err := responsePolicyDecisionStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ResponsePolicyDecision{native: responsePolicyDecisionStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeResponsePolicyDecision)
	return structGo
}
func finalizeResponsePolicyDecision(obj *ResponsePolicyDecision) {
	responsePolicyDecisionStruct.Free(obj.native)
}

var securityManagerStruct *gi.Struct
var securityManagerStruct_Once sync.Once

func securityManagerStruct_Set() error {
	var err error
	securityManagerStruct_Once.Do(func() {
		securityManagerStruct, err = gi.StructNew("WebKit2", "SecurityManager")
	})
	return err
}

type SecurityManager struct {
	native uintptr
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'GObject.Object'

// FieldPriv returns the C field 'priv'.
func (recv *SecurityManager) FieldPriv() *SecurityManagerPrivate {
	argValue := gi.FieldGet(securityManagerStruct, recv.native, "priv")
	value := &SecurityManagerPrivate{native: argValue.Pointer()}
	return value
}

// SetFieldPriv sets the value of the C field 'priv'.
func (recv *SecurityManager) SetFieldPriv(value *SecurityManagerPrivate) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(securityManagerStruct, recv.native, "priv", argValue)
}

var securityManagerRegisterUriSchemeAsCorsEnabledFunction *gi.Function
var securityManagerRegisterUriSchemeAsCorsEnabledFunction_Once sync.Once

func securityManagerRegisterUriSchemeAsCorsEnabledFunction_Set() error {
	var err error
	securityManagerRegisterUriSchemeAsCorsEnabledFunction_Once.Do(func() {
		err = securityManagerStruct_Set()
		if err != nil {
			return
		}
		securityManagerRegisterUriSchemeAsCorsEnabledFunction, err = securityManagerStruct.InvokerNew("register_uri_scheme_as_cors_enabled")
	})
	return err
}

// RegisterUriSchemeAsCorsEnabled is a representation of the C type webkit_security_manager_register_uri_scheme_as_cors_enabled.
func (recv *SecurityManager) RegisterUriSchemeAsCorsEnabled(scheme string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(scheme)

	err := securityManagerRegisterUriSchemeAsCorsEnabledFunction_Set()
	if err == nil {
		securityManagerRegisterUriSchemeAsCorsEnabledFunction.Invoke(inArgs[:], nil)
	}

	return
}

var securityManagerRegisterUriSchemeAsDisplayIsolatedFunction *gi.Function
var securityManagerRegisterUriSchemeAsDisplayIsolatedFunction_Once sync.Once

func securityManagerRegisterUriSchemeAsDisplayIsolatedFunction_Set() error {
	var err error
	securityManagerRegisterUriSchemeAsDisplayIsolatedFunction_Once.Do(func() {
		err = securityManagerStruct_Set()
		if err != nil {
			return
		}
		securityManagerRegisterUriSchemeAsDisplayIsolatedFunction, err = securityManagerStruct.InvokerNew("register_uri_scheme_as_display_isolated")
	})
	return err
}

// RegisterUriSchemeAsDisplayIsolated is a representation of the C type webkit_security_manager_register_uri_scheme_as_display_isolated.
func (recv *SecurityManager) RegisterUriSchemeAsDisplayIsolated(scheme string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(scheme)

	err := securityManagerRegisterUriSchemeAsDisplayIsolatedFunction_Set()
	if err == nil {
		securityManagerRegisterUriSchemeAsDisplayIsolatedFunction.Invoke(inArgs[:], nil)
	}

	return
}

var securityManagerRegisterUriSchemeAsEmptyDocumentFunction *gi.Function
var securityManagerRegisterUriSchemeAsEmptyDocumentFunction_Once sync.Once

func securityManagerRegisterUriSchemeAsEmptyDocumentFunction_Set() error {
	var err error
	securityManagerRegisterUriSchemeAsEmptyDocumentFunction_Once.Do(func() {
		err = securityManagerStruct_Set()
		if err != nil {
			return
		}
		securityManagerRegisterUriSchemeAsEmptyDocumentFunction, err = securityManagerStruct.InvokerNew("register_uri_scheme_as_empty_document")
	})
	return err
}

// RegisterUriSchemeAsEmptyDocument is a representation of the C type webkit_security_manager_register_uri_scheme_as_empty_document.
func (recv *SecurityManager) RegisterUriSchemeAsEmptyDocument(scheme string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(scheme)

	err := securityManagerRegisterUriSchemeAsEmptyDocumentFunction_Set()
	if err == nil {
		securityManagerRegisterUriSchemeAsEmptyDocumentFunction.Invoke(inArgs[:], nil)
	}

	return
}

var securityManagerRegisterUriSchemeAsLocalFunction *gi.Function
var securityManagerRegisterUriSchemeAsLocalFunction_Once sync.Once

func securityManagerRegisterUriSchemeAsLocalFunction_Set() error {
	var err error
	securityManagerRegisterUriSchemeAsLocalFunction_Once.Do(func() {
		err = securityManagerStruct_Set()
		if err != nil {
			return
		}
		securityManagerRegisterUriSchemeAsLocalFunction, err = securityManagerStruct.InvokerNew("register_uri_scheme_as_local")
	})
	return err
}

// RegisterUriSchemeAsLocal is a representation of the C type webkit_security_manager_register_uri_scheme_as_local.
func (recv *SecurityManager) RegisterUriSchemeAsLocal(scheme string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(scheme)

	err := securityManagerRegisterUriSchemeAsLocalFunction_Set()
	if err == nil {
		securityManagerRegisterUriSchemeAsLocalFunction.Invoke(inArgs[:], nil)
	}

	return
}

var securityManagerRegisterUriSchemeAsNoAccessFunction *gi.Function
var securityManagerRegisterUriSchemeAsNoAccessFunction_Once sync.Once

func securityManagerRegisterUriSchemeAsNoAccessFunction_Set() error {
	var err error
	securityManagerRegisterUriSchemeAsNoAccessFunction_Once.Do(func() {
		err = securityManagerStruct_Set()
		if err != nil {
			return
		}
		securityManagerRegisterUriSchemeAsNoAccessFunction, err = securityManagerStruct.InvokerNew("register_uri_scheme_as_no_access")
	})
	return err
}

// RegisterUriSchemeAsNoAccess is a representation of the C type webkit_security_manager_register_uri_scheme_as_no_access.
func (recv *SecurityManager) RegisterUriSchemeAsNoAccess(scheme string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(scheme)

	err := securityManagerRegisterUriSchemeAsNoAccessFunction_Set()
	if err == nil {
		securityManagerRegisterUriSchemeAsNoAccessFunction.Invoke(inArgs[:], nil)
	}

	return
}

var securityManagerRegisterUriSchemeAsSecureFunction *gi.Function
var securityManagerRegisterUriSchemeAsSecureFunction_Once sync.Once

func securityManagerRegisterUriSchemeAsSecureFunction_Set() error {
	var err error
	securityManagerRegisterUriSchemeAsSecureFunction_Once.Do(func() {
		err = securityManagerStruct_Set()
		if err != nil {
			return
		}
		securityManagerRegisterUriSchemeAsSecureFunction, err = securityManagerStruct.InvokerNew("register_uri_scheme_as_secure")
	})
	return err
}

// RegisterUriSchemeAsSecure is a representation of the C type webkit_security_manager_register_uri_scheme_as_secure.
func (recv *SecurityManager) RegisterUriSchemeAsSecure(scheme string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(scheme)

	err := securityManagerRegisterUriSchemeAsSecureFunction_Set()
	if err == nil {
		securityManagerRegisterUriSchemeAsSecureFunction.Invoke(inArgs[:], nil)
	}

	return
}

var securityManagerUriSchemeIsCorsEnabledFunction *gi.Function
var securityManagerUriSchemeIsCorsEnabledFunction_Once sync.Once

func securityManagerUriSchemeIsCorsEnabledFunction_Set() error {
	var err error
	securityManagerUriSchemeIsCorsEnabledFunction_Once.Do(func() {
		err = securityManagerStruct_Set()
		if err != nil {
			return
		}
		securityManagerUriSchemeIsCorsEnabledFunction, err = securityManagerStruct.InvokerNew("uri_scheme_is_cors_enabled")
	})
	return err
}

// UriSchemeIsCorsEnabled is a representation of the C type webkit_security_manager_uri_scheme_is_cors_enabled.
func (recv *SecurityManager) UriSchemeIsCorsEnabled(scheme string) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(scheme)

	var ret gi.Argument

	err := securityManagerUriSchemeIsCorsEnabledFunction_Set()
	if err == nil {
		ret = securityManagerUriSchemeIsCorsEnabledFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var securityManagerUriSchemeIsDisplayIsolatedFunction *gi.Function
var securityManagerUriSchemeIsDisplayIsolatedFunction_Once sync.Once

func securityManagerUriSchemeIsDisplayIsolatedFunction_Set() error {
	var err error
	securityManagerUriSchemeIsDisplayIsolatedFunction_Once.Do(func() {
		err = securityManagerStruct_Set()
		if err != nil {
			return
		}
		securityManagerUriSchemeIsDisplayIsolatedFunction, err = securityManagerStruct.InvokerNew("uri_scheme_is_display_isolated")
	})
	return err
}

// UriSchemeIsDisplayIsolated is a representation of the C type webkit_security_manager_uri_scheme_is_display_isolated.
func (recv *SecurityManager) UriSchemeIsDisplayIsolated(scheme string) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(scheme)

	var ret gi.Argument

	err := securityManagerUriSchemeIsDisplayIsolatedFunction_Set()
	if err == nil {
		ret = securityManagerUriSchemeIsDisplayIsolatedFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var securityManagerUriSchemeIsEmptyDocumentFunction *gi.Function
var securityManagerUriSchemeIsEmptyDocumentFunction_Once sync.Once

func securityManagerUriSchemeIsEmptyDocumentFunction_Set() error {
	var err error
	securityManagerUriSchemeIsEmptyDocumentFunction_Once.Do(func() {
		err = securityManagerStruct_Set()
		if err != nil {
			return
		}
		securityManagerUriSchemeIsEmptyDocumentFunction, err = securityManagerStruct.InvokerNew("uri_scheme_is_empty_document")
	})
	return err
}

// UriSchemeIsEmptyDocument is a representation of the C type webkit_security_manager_uri_scheme_is_empty_document.
func (recv *SecurityManager) UriSchemeIsEmptyDocument(scheme string) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(scheme)

	var ret gi.Argument

	err := securityManagerUriSchemeIsEmptyDocumentFunction_Set()
	if err == nil {
		ret = securityManagerUriSchemeIsEmptyDocumentFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var securityManagerUriSchemeIsLocalFunction *gi.Function
var securityManagerUriSchemeIsLocalFunction_Once sync.Once

func securityManagerUriSchemeIsLocalFunction_Set() error {
	var err error
	securityManagerUriSchemeIsLocalFunction_Once.Do(func() {
		err = securityManagerStruct_Set()
		if err != nil {
			return
		}
		securityManagerUriSchemeIsLocalFunction, err = securityManagerStruct.InvokerNew("uri_scheme_is_local")
	})
	return err
}

// UriSchemeIsLocal is a representation of the C type webkit_security_manager_uri_scheme_is_local.
func (recv *SecurityManager) UriSchemeIsLocal(scheme string) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(scheme)

	var ret gi.Argument

	err := securityManagerUriSchemeIsLocalFunction_Set()
	if err == nil {
		ret = securityManagerUriSchemeIsLocalFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var securityManagerUriSchemeIsNoAccessFunction *gi.Function
var securityManagerUriSchemeIsNoAccessFunction_Once sync.Once

func securityManagerUriSchemeIsNoAccessFunction_Set() error {
	var err error
	securityManagerUriSchemeIsNoAccessFunction_Once.Do(func() {
		err = securityManagerStruct_Set()
		if err != nil {
			return
		}
		securityManagerUriSchemeIsNoAccessFunction, err = securityManagerStruct.InvokerNew("uri_scheme_is_no_access")
	})
	return err
}

// UriSchemeIsNoAccess is a representation of the C type webkit_security_manager_uri_scheme_is_no_access.
func (recv *SecurityManager) UriSchemeIsNoAccess(scheme string) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(scheme)

	var ret gi.Argument

	err := securityManagerUriSchemeIsNoAccessFunction_Set()
	if err == nil {
		ret = securityManagerUriSchemeIsNoAccessFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var securityManagerUriSchemeIsSecureFunction *gi.Function
var securityManagerUriSchemeIsSecureFunction_Once sync.Once

func securityManagerUriSchemeIsSecureFunction_Set() error {
	var err error
	securityManagerUriSchemeIsSecureFunction_Once.Do(func() {
		err = securityManagerStruct_Set()
		if err != nil {
			return
		}
		securityManagerUriSchemeIsSecureFunction, err = securityManagerStruct.InvokerNew("uri_scheme_is_secure")
	})
	return err
}

// UriSchemeIsSecure is a representation of the C type webkit_security_manager_uri_scheme_is_secure.
func (recv *SecurityManager) UriSchemeIsSecure(scheme string) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(scheme)

	var ret gi.Argument

	err := securityManagerUriSchemeIsSecureFunction_Set()
	if err == nil {
		ret = securityManagerUriSchemeIsSecureFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// SecurityManagerStruct creates an uninitialised SecurityManager.
func SecurityManagerStruct() *SecurityManager {
	err := securityManagerStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &SecurityManager{native: securityManagerStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeSecurityManager)
	return structGo
}
func finalizeSecurityManager(obj *SecurityManager) {
	securityManagerStruct.Free(obj.native)
}

var settingsStruct *gi.Struct
var settingsStruct_Once sync.Once

func settingsStruct_Set() error {
	var err error
	settingsStruct_Once.Do(func() {
		settingsStruct, err = gi.StructNew("WebKit2", "Settings")
	})
	return err
}

type Settings struct {
	native uintptr
}

// UNSUPPORTED : C value 'parent_instance' : for field getter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'parent_instance' : for field setter : no Go type for 'GObject.Object'

// FieldPriv returns the C field 'priv'.
func (recv *Settings) FieldPriv() *SettingsPrivate {
	argValue := gi.FieldGet(settingsStruct, recv.native, "priv")
	value := &SettingsPrivate{native: argValue.Pointer()}
	return value
}

// SetFieldPriv sets the value of the C field 'priv'.
func (recv *Settings) SetFieldPriv(value *SettingsPrivate) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(settingsStruct, recv.native, "priv", argValue)
}

// UNSUPPORTED : C value 'webkit_settings_new' : return type 'Settings' not supported

// UNSUPPORTED : C value 'webkit_settings_new_with_settings' : parameter '...' of type 'nil' not supported

var settingsGetAllowFileAccessFromFileUrlsFunction *gi.Function
var settingsGetAllowFileAccessFromFileUrlsFunction_Once sync.Once

func settingsGetAllowFileAccessFromFileUrlsFunction_Set() error {
	var err error
	settingsGetAllowFileAccessFromFileUrlsFunction_Once.Do(func() {
		err = settingsStruct_Set()
		if err != nil {
			return
		}
		settingsGetAllowFileAccessFromFileUrlsFunction, err = settingsStruct.InvokerNew("get_allow_file_access_from_file_urls")
	})
	return err
}

// GetAllowFileAccessFromFileUrls is a representation of the C type webkit_settings_get_allow_file_access_from_file_urls.
func (recv *Settings) GetAllowFileAccessFromFileUrls() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := settingsGetAllowFileAccessFromFileUrlsFunction_Set()
	if err == nil {
		ret = settingsGetAllowFileAccessFromFileUrlsFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var settingsGetAllowModalDialogsFunction *gi.Function
var settingsGetAllowModalDialogsFunction_Once sync.Once

func settingsGetAllowModalDialogsFunction_Set() error {
	var err error
	settingsGetAllowModalDialogsFunction_Once.Do(func() {
		err = settingsStruct_Set()
		if err != nil {
			return
		}
		settingsGetAllowModalDialogsFunction, err = settingsStruct.InvokerNew("get_allow_modal_dialogs")
	})
	return err
}

// GetAllowModalDialogs is a representation of the C type webkit_settings_get_allow_modal_dialogs.
func (recv *Settings) GetAllowModalDialogs() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := settingsGetAllowModalDialogsFunction_Set()
	if err == nil {
		ret = settingsGetAllowModalDialogsFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var settingsGetAllowUniversalAccessFromFileUrlsFunction *gi.Function
var settingsGetAllowUniversalAccessFromFileUrlsFunction_Once sync.Once

func settingsGetAllowUniversalAccessFromFileUrlsFunction_Set() error {
	var err error
	settingsGetAllowUniversalAccessFromFileUrlsFunction_Once.Do(func() {
		err = settingsStruct_Set()
		if err != nil {
			return
		}
		settingsGetAllowUniversalAccessFromFileUrlsFunction, err = settingsStruct.InvokerNew("get_allow_universal_access_from_file_urls")
	})
	return err
}

// GetAllowUniversalAccessFromFileUrls is a representation of the C type webkit_settings_get_allow_universal_access_from_file_urls.
func (recv *Settings) GetAllowUniversalAccessFromFileUrls() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := settingsGetAllowUniversalAccessFromFileUrlsFunction_Set()
	if err == nil {
		ret = settingsGetAllowUniversalAccessFromFileUrlsFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var settingsGetAutoLoadImagesFunction *gi.Function
var settingsGetAutoLoadImagesFunction_Once sync.Once

func settingsGetAutoLoadImagesFunction_Set() error {
	var err error
	settingsGetAutoLoadImagesFunction_Once.Do(func() {
		err = settingsStruct_Set()
		if err != nil {
			return
		}
		settingsGetAutoLoadImagesFunction, err = settingsStruct.InvokerNew("get_auto_load_images")
	})
	return err
}

// GetAutoLoadImages is a representation of the C type webkit_settings_get_auto_load_images.
func (recv *Settings) GetAutoLoadImages() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := settingsGetAutoLoadImagesFunction_Set()
	if err == nil {
		ret = settingsGetAutoLoadImagesFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var settingsGetCursiveFontFamilyFunction *gi.Function
var settingsGetCursiveFontFamilyFunction_Once sync.Once

func settingsGetCursiveFontFamilyFunction_Set() error {
	var err error
	settingsGetCursiveFontFamilyFunction_Once.Do(func() {
		err = settingsStruct_Set()
		if err != nil {
			return
		}
		settingsGetCursiveFontFamilyFunction, err = settingsStruct.InvokerNew("get_cursive_font_family")
	})
	return err
}

// GetCursiveFontFamily is a representation of the C type webkit_settings_get_cursive_font_family.
func (recv *Settings) GetCursiveFontFamily() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := settingsGetCursiveFontFamilyFunction_Set()
	if err == nil {
		ret = settingsGetCursiveFontFamilyFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var settingsGetDefaultCharsetFunction *gi.Function
var settingsGetDefaultCharsetFunction_Once sync.Once

func settingsGetDefaultCharsetFunction_Set() error {
	var err error
	settingsGetDefaultCharsetFunction_Once.Do(func() {
		err = settingsStruct_Set()
		if err != nil {
			return
		}
		settingsGetDefaultCharsetFunction, err = settingsStruct.InvokerNew("get_default_charset")
	})
	return err
}

// GetDefaultCharset is a representation of the C type webkit_settings_get_default_charset.
func (recv *Settings) GetDefaultCharset() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := settingsGetDefaultCharsetFunction_Set()
	if err == nil {
		ret = settingsGetDefaultCharsetFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var settingsGetDefaultFontFamilyFunction *gi.Function
var settingsGetDefaultFontFamilyFunction_Once sync.Once

func settingsGetDefaultFontFamilyFunction_Set() error {
	var err error
	settingsGetDefaultFontFamilyFunction_Once.Do(func() {
		err = settingsStruct_Set()
		if err != nil {
			return
		}
		settingsGetDefaultFontFamilyFunction, err = settingsStruct.InvokerNew("get_default_font_family")
	})
	return err
}

// GetDefaultFontFamily is a representation of the C type webkit_settings_get_default_font_family.
func (recv *Settings) GetDefaultFontFamily() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := settingsGetDefaultFontFamilyFunction_Set()
	if err == nil {
		ret = settingsGetDefaultFontFamilyFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var settingsGetDefaultFontSizeFunction *gi.Function
var settingsGetDefaultFontSizeFunction_Once sync.Once

func settingsGetDefaultFontSizeFunction_Set() error {
	var err error
	settingsGetDefaultFontSizeFunction_Once.Do(func() {
		err = settingsStruct_Set()
		if err != nil {
			return
		}
		settingsGetDefaultFontSizeFunction, err = settingsStruct.InvokerNew("get_default_font_size")
	})
	return err
}

// GetDefaultFontSize is a representation of the C type webkit_settings_get_default_font_size.
func (recv *Settings) GetDefaultFontSize() uint32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := settingsGetDefaultFontSizeFunction_Set()
	if err == nil {
		ret = settingsGetDefaultFontSizeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo
}

var settingsGetDefaultMonospaceFontSizeFunction *gi.Function
var settingsGetDefaultMonospaceFontSizeFunction_Once sync.Once

func settingsGetDefaultMonospaceFontSizeFunction_Set() error {
	var err error
	settingsGetDefaultMonospaceFontSizeFunction_Once.Do(func() {
		err = settingsStruct_Set()
		if err != nil {
			return
		}
		settingsGetDefaultMonospaceFontSizeFunction, err = settingsStruct.InvokerNew("get_default_monospace_font_size")
	})
	return err
}

// GetDefaultMonospaceFontSize is a representation of the C type webkit_settings_get_default_monospace_font_size.
func (recv *Settings) GetDefaultMonospaceFontSize() uint32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := settingsGetDefaultMonospaceFontSizeFunction_Set()
	if err == nil {
		ret = settingsGetDefaultMonospaceFontSizeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo
}

var settingsGetDrawCompositingIndicatorsFunction *gi.Function
var settingsGetDrawCompositingIndicatorsFunction_Once sync.Once

func settingsGetDrawCompositingIndicatorsFunction_Set() error {
	var err error
	settingsGetDrawCompositingIndicatorsFunction_Once.Do(func() {
		err = settingsStruct_Set()
		if err != nil {
			return
		}
		settingsGetDrawCompositingIndicatorsFunction, err = settingsStruct.InvokerNew("get_draw_compositing_indicators")
	})
	return err
}

// GetDrawCompositingIndicators is a representation of the C type webkit_settings_get_draw_compositing_indicators.
func (recv *Settings) GetDrawCompositingIndicators() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := settingsGetDrawCompositingIndicatorsFunction_Set()
	if err == nil {
		ret = settingsGetDrawCompositingIndicatorsFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var settingsGetEnableAccelerated2dCanvasFunction *gi.Function
var settingsGetEnableAccelerated2dCanvasFunction_Once sync.Once

func settingsGetEnableAccelerated2dCanvasFunction_Set() error {
	var err error
	settingsGetEnableAccelerated2dCanvasFunction_Once.Do(func() {
		err = settingsStruct_Set()
		if err != nil {
			return
		}
		settingsGetEnableAccelerated2dCanvasFunction, err = settingsStruct.InvokerNew("get_enable_accelerated_2d_canvas")
	})
	return err
}

// GetEnableAccelerated2dCanvas is a representation of the C type webkit_settings_get_enable_accelerated_2d_canvas.
func (recv *Settings) GetEnableAccelerated2dCanvas() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := settingsGetEnableAccelerated2dCanvasFunction_Set()
	if err == nil {
		ret = settingsGetEnableAccelerated2dCanvasFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var settingsGetEnableBackForwardNavigationGesturesFunction *gi.Function
var settingsGetEnableBackForwardNavigationGesturesFunction_Once sync.Once

func settingsGetEnableBackForwardNavigationGesturesFunction_Set() error {
	var err error
	settingsGetEnableBackForwardNavigationGesturesFunction_Once.Do(func() {
		err = settingsStruct_Set()
		if err != nil {
			return
		}
		settingsGetEnableBackForwardNavigationGesturesFunction, err = settingsStruct.InvokerNew("get_enable_back_forward_navigation_gestures")
	})
	return err
}

// GetEnableBackForwardNavigationGestures is a representation of the C type webkit_settings_get_enable_back_forward_navigation_gestures.
func (recv *Settings) GetEnableBackForwardNavigationGestures() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := settingsGetEnableBackForwardNavigationGesturesFunction_Set()
	if err == nil {
		ret = settingsGetEnableBackForwardNavigationGesturesFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var settingsGetEnableCaretBrowsingFunction *gi.Function
var settingsGetEnableCaretBrowsingFunction_Once sync.Once

func settingsGetEnableCaretBrowsingFunction_Set() error {
	var err error
	settingsGetEnableCaretBrowsingFunction_Once.Do(func() {
		err = settingsStruct_Set()
		if err != nil {
			return
		}
		settingsGetEnableCaretBrowsingFunction, err = settingsStruct.InvokerNew("get_enable_caret_browsing")
	})
	return err
}

// GetEnableCaretBrowsing is a representation of the C type webkit_settings_get_enable_caret_browsing.
func (recv *Settings) GetEnableCaretBrowsing() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := settingsGetEnableCaretBrowsingFunction_Set()
	if err == nil {
		ret = settingsGetEnableCaretBrowsingFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var settingsGetEnableDeveloperExtrasFunction *gi.Function
var settingsGetEnableDeveloperExtrasFunction_Once sync.Once

func settingsGetEnableDeveloperExtrasFunction_Set() error {
	var err error
	settingsGetEnableDeveloperExtrasFunction_Once.Do(func() {
		err = settingsStruct_Set()
		if err != nil {
			return
		}
		settingsGetEnableDeveloperExtrasFunction, err = settingsStruct.InvokerNew("get_enable_developer_extras")
	})
	return err
}

// GetEnableDeveloperExtras is a representation of the C type webkit_settings_get_enable_developer_extras.
func (recv *Settings) GetEnableDeveloperExtras() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := settingsGetEnableDeveloperExtrasFunction_Set()
	if err == nil {
		ret = settingsGetEnableDeveloperExtrasFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var settingsGetEnableDnsPrefetchingFunction *gi.Function
var settingsGetEnableDnsPrefetchingFunction_Once sync.Once

func settingsGetEnableDnsPrefetchingFunction_Set() error {
	var err error
	settingsGetEnableDnsPrefetchingFunction_Once.Do(func() {
		err = settingsStruct_Set()
		if err != nil {
			return
		}
		settingsGetEnableDnsPrefetchingFunction, err = settingsStruct.InvokerNew("get_enable_dns_prefetching")
	})
	return err
}

// GetEnableDnsPrefetching is a representation of the C type webkit_settings_get_enable_dns_prefetching.
func (recv *Settings) GetEnableDnsPrefetching() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := settingsGetEnableDnsPrefetchingFunction_Set()
	if err == nil {
		ret = settingsGetEnableDnsPrefetchingFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var settingsGetEnableEncryptedMediaFunction *gi.Function
var settingsGetEnableEncryptedMediaFunction_Once sync.Once

func settingsGetEnableEncryptedMediaFunction_Set() error {
	var err error
	settingsGetEnableEncryptedMediaFunction_Once.Do(func() {
		err = settingsStruct_Set()
		if err != nil {
			return
		}
		settingsGetEnableEncryptedMediaFunction, err = settingsStruct.InvokerNew("get_enable_encrypted_media")
	})
	return err
}

// GetEnableEncryptedMedia is a representation of the C type webkit_settings_get_enable_encrypted_media.
func (recv *Settings) GetEnableEncryptedMedia() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := settingsGetEnableEncryptedMediaFunction_Set()
	if err == nil {
		ret = settingsGetEnableEncryptedMediaFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var settingsGetEnableFrameFlatteningFunction *gi.Function
var settingsGetEnableFrameFlatteningFunction_Once sync.Once

func settingsGetEnableFrameFlatteningFunction_Set() error {
	var err error
	settingsGetEnableFrameFlatteningFunction_Once.Do(func() {
		err = settingsStruct_Set()
		if err != nil {
			return
		}
		settingsGetEnableFrameFlatteningFunction, err = settingsStruct.InvokerNew("get_enable_frame_flattening")
	})
	return err
}

// GetEnableFrameFlattening is a representation of the C type webkit_settings_get_enable_frame_flattening.
func (recv *Settings) GetEnableFrameFlattening() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := settingsGetEnableFrameFlatteningFunction_Set()
	if err == nil {
		ret = settingsGetEnableFrameFlatteningFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var settingsGetEnableFullscreenFunction *gi.Function
var settingsGetEnableFullscreenFunction_Once sync.Once

func settingsGetEnableFullscreenFunction_Set() error {
	var err error
	settingsGetEnableFullscreenFunction_Once.Do(func() {
		err = settingsStruct_Set()
		if err != nil {
			return
		}
		settingsGetEnableFullscreenFunction, err = settingsStruct.InvokerNew("get_enable_fullscreen")
	})
	return err
}

// GetEnableFullscreen is a representation of the C type webkit_settings_get_enable_fullscreen.
func (recv *Settings) GetEnableFullscreen() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := settingsGetEnableFullscreenFunction_Set()
	if err == nil {
		ret = settingsGetEnableFullscreenFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var settingsGetEnableHtml5DatabaseFunction *gi.Function
var settingsGetEnableHtml5DatabaseFunction_Once sync.Once

func settingsGetEnableHtml5DatabaseFunction_Set() error {
	var err error
	settingsGetEnableHtml5DatabaseFunction_Once.Do(func() {
		err = settingsStruct_Set()
		if err != nil {
			return
		}
		settingsGetEnableHtml5DatabaseFunction, err = settingsStruct.InvokerNew("get_enable_html5_database")
	})
	return err
}

// GetEnableHtml5Database is a representation of the C type webkit_settings_get_enable_html5_database.
func (recv *Settings) GetEnableHtml5Database() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := settingsGetEnableHtml5DatabaseFunction_Set()
	if err == nil {
		ret = settingsGetEnableHtml5DatabaseFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var settingsGetEnableHtml5LocalStorageFunction *gi.Function
var settingsGetEnableHtml5LocalStorageFunction_Once sync.Once

func settingsGetEnableHtml5LocalStorageFunction_Set() error {
	var err error
	settingsGetEnableHtml5LocalStorageFunction_Once.Do(func() {
		err = settingsStruct_Set()
		if err != nil {
			return
		}
		settingsGetEnableHtml5LocalStorageFunction, err = settingsStruct.InvokerNew("get_enable_html5_local_storage")
	})
	return err
}

// GetEnableHtml5LocalStorage is a representation of the C type webkit_settings_get_enable_html5_local_storage.
func (recv *Settings) GetEnableHtml5LocalStorage() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := settingsGetEnableHtml5LocalStorageFunction_Set()
	if err == nil {
		ret = settingsGetEnableHtml5LocalStorageFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var settingsGetEnableHyperlinkAuditingFunction *gi.Function
var settingsGetEnableHyperlinkAuditingFunction_Once sync.Once

func settingsGetEnableHyperlinkAuditingFunction_Set() error {
	var err error
	settingsGetEnableHyperlinkAuditingFunction_Once.Do(func() {
		err = settingsStruct_Set()
		if err != nil {
			return
		}
		settingsGetEnableHyperlinkAuditingFunction, err = settingsStruct.InvokerNew("get_enable_hyperlink_auditing")
	})
	return err
}

// GetEnableHyperlinkAuditing is a representation of the C type webkit_settings_get_enable_hyperlink_auditing.
func (recv *Settings) GetEnableHyperlinkAuditing() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := settingsGetEnableHyperlinkAuditingFunction_Set()
	if err == nil {
		ret = settingsGetEnableHyperlinkAuditingFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var settingsGetEnableJavaFunction *gi.Function
var settingsGetEnableJavaFunction_Once sync.Once

func settingsGetEnableJavaFunction_Set() error {
	var err error
	settingsGetEnableJavaFunction_Once.Do(func() {
		err = settingsStruct_Set()
		if err != nil {
			return
		}
		settingsGetEnableJavaFunction, err = settingsStruct.InvokerNew("get_enable_java")
	})
	return err
}

// GetEnableJava is a representation of the C type webkit_settings_get_enable_java.
func (recv *Settings) GetEnableJava() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := settingsGetEnableJavaFunction_Set()
	if err == nil {
		ret = settingsGetEnableJavaFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var settingsGetEnableJavascriptFunction *gi.Function
var settingsGetEnableJavascriptFunction_Once sync.Once

func settingsGetEnableJavascriptFunction_Set() error {
	var err error
	settingsGetEnableJavascriptFunction_Once.Do(func() {
		err = settingsStruct_Set()
		if err != nil {
			return
		}
		settingsGetEnableJavascriptFunction, err = settingsStruct.InvokerNew("get_enable_javascript")
	})
	return err
}

// GetEnableJavascript is a representation of the C type webkit_settings_get_enable_javascript.
func (recv *Settings) GetEnableJavascript() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := settingsGetEnableJavascriptFunction_Set()
	if err == nil {
		ret = settingsGetEnableJavascriptFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var settingsGetEnableJavascriptMarkupFunction *gi.Function
var settingsGetEnableJavascriptMarkupFunction_Once sync.Once

func settingsGetEnableJavascriptMarkupFunction_Set() error {
	var err error
	settingsGetEnableJavascriptMarkupFunction_Once.Do(func() {
		err = settingsStruct_Set()
		if err != nil {
			return
		}
		settingsGetEnableJavascriptMarkupFunction, err = settingsStruct.InvokerNew("get_enable_javascript_markup")
	})
	return err
}

// GetEnableJavascriptMarkup is a representation of the C type webkit_settings_get_enable_javascript_markup.
func (recv *Settings) GetEnableJavascriptMarkup() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := settingsGetEnableJavascriptMarkupFunction_Set()
	if err == nil {
		ret = settingsGetEnableJavascriptMarkupFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var settingsGetEnableMediaFunction *gi.Function
var settingsGetEnableMediaFunction_Once sync.Once

func settingsGetEnableMediaFunction_Set() error {
	var err error
	settingsGetEnableMediaFunction_Once.Do(func() {
		err = settingsStruct_Set()
		if err != nil {
			return
		}
		settingsGetEnableMediaFunction, err = settingsStruct.InvokerNew("get_enable_media")
	})
	return err
}

// GetEnableMedia is a representation of the C type webkit_settings_get_enable_media.
func (recv *Settings) GetEnableMedia() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := settingsGetEnableMediaFunction_Set()
	if err == nil {
		ret = settingsGetEnableMediaFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var settingsGetEnableMediaCapabilitiesFunction *gi.Function
var settingsGetEnableMediaCapabilitiesFunction_Once sync.Once

func settingsGetEnableMediaCapabilitiesFunction_Set() error {
	var err error
	settingsGetEnableMediaCapabilitiesFunction_Once.Do(func() {
		err = settingsStruct_Set()
		if err != nil {
			return
		}
		settingsGetEnableMediaCapabilitiesFunction, err = settingsStruct.InvokerNew("get_enable_media_capabilities")
	})
	return err
}

// GetEnableMediaCapabilities is a representation of the C type webkit_settings_get_enable_media_capabilities.
func (recv *Settings) GetEnableMediaCapabilities() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := settingsGetEnableMediaCapabilitiesFunction_Set()
	if err == nil {
		ret = settingsGetEnableMediaCapabilitiesFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var settingsGetEnableMediaStreamFunction *gi.Function
var settingsGetEnableMediaStreamFunction_Once sync.Once

func settingsGetEnableMediaStreamFunction_Set() error {
	var err error
	settingsGetEnableMediaStreamFunction_Once.Do(func() {
		err = settingsStruct_Set()
		if err != nil {
			return
		}
		settingsGetEnableMediaStreamFunction, err = settingsStruct.InvokerNew("get_enable_media_stream")
	})
	return err
}

// GetEnableMediaStream is a representation of the C type webkit_settings_get_enable_media_stream.
func (recv *Settings) GetEnableMediaStream() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := settingsGetEnableMediaStreamFunction_Set()
	if err == nil {
		ret = settingsGetEnableMediaStreamFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var settingsGetEnableMediasourceFunction *gi.Function
var settingsGetEnableMediasourceFunction_Once sync.Once

func settingsGetEnableMediasourceFunction_Set() error {
	var err error
	settingsGetEnableMediasourceFunction_Once.Do(func() {
		err = settingsStruct_Set()
		if err != nil {
			return
		}
		settingsGetEnableMediasourceFunction, err = settingsStruct.InvokerNew("get_enable_mediasource")
	})
	return err
}

// GetEnableMediasource is a representation of the C type webkit_settings_get_enable_mediasource.
func (recv *Settings) GetEnableMediasource() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := settingsGetEnableMediasourceFunction_Set()
	if err == nil {
		ret = settingsGetEnableMediasourceFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var settingsGetEnableMockCaptureDevicesFunction *gi.Function
var settingsGetEnableMockCaptureDevicesFunction_Once sync.Once

func settingsGetEnableMockCaptureDevicesFunction_Set() error {
	var err error
	settingsGetEnableMockCaptureDevicesFunction_Once.Do(func() {
		err = settingsStruct_Set()
		if err != nil {
			return
		}
		settingsGetEnableMockCaptureDevicesFunction, err = settingsStruct.InvokerNew("get_enable_mock_capture_devices")
	})
	return err
}

// GetEnableMockCaptureDevices is a representation of the C type webkit_settings_get_enable_mock_capture_devices.
func (recv *Settings) GetEnableMockCaptureDevices() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := settingsGetEnableMockCaptureDevicesFunction_Set()
	if err == nil {
		ret = settingsGetEnableMockCaptureDevicesFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var settingsGetEnableOfflineWebApplicationCacheFunction *gi.Function
var settingsGetEnableOfflineWebApplicationCacheFunction_Once sync.Once

func settingsGetEnableOfflineWebApplicationCacheFunction_Set() error {
	var err error
	settingsGetEnableOfflineWebApplicationCacheFunction_Once.Do(func() {
		err = settingsStruct_Set()
		if err != nil {
			return
		}
		settingsGetEnableOfflineWebApplicationCacheFunction, err = settingsStruct.InvokerNew("get_enable_offline_web_application_cache")
	})
	return err
}

// GetEnableOfflineWebApplicationCache is a representation of the C type webkit_settings_get_enable_offline_web_application_cache.
func (recv *Settings) GetEnableOfflineWebApplicationCache() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := settingsGetEnableOfflineWebApplicationCacheFunction_Set()
	if err == nil {
		ret = settingsGetEnableOfflineWebApplicationCacheFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var settingsGetEnablePageCacheFunction *gi.Function
var settingsGetEnablePageCacheFunction_Once sync.Once

func settingsGetEnablePageCacheFunction_Set() error {
	var err error
	settingsGetEnablePageCacheFunction_Once.Do(func() {
		err = settingsStruct_Set()
		if err != nil {
			return
		}
		settingsGetEnablePageCacheFunction, err = settingsStruct.InvokerNew("get_enable_page_cache")
	})
	return err
}

// GetEnablePageCache is a representation of the C type webkit_settings_get_enable_page_cache.
func (recv *Settings) GetEnablePageCache() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := settingsGetEnablePageCacheFunction_Set()
	if err == nil {
		ret = settingsGetEnablePageCacheFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var settingsGetEnablePluginsFunction *gi.Function
var settingsGetEnablePluginsFunction_Once sync.Once

func settingsGetEnablePluginsFunction_Set() error {
	var err error
	settingsGetEnablePluginsFunction_Once.Do(func() {
		err = settingsStruct_Set()
		if err != nil {
			return
		}
		settingsGetEnablePluginsFunction, err = settingsStruct.InvokerNew("get_enable_plugins")
	})
	return err
}

// GetEnablePlugins is a representation of the C type webkit_settings_get_enable_plugins.
func (recv *Settings) GetEnablePlugins() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := settingsGetEnablePluginsFunction_Set()
	if err == nil {
		ret = settingsGetEnablePluginsFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var settingsGetEnablePrivateBrowsingFunction *gi.Function
var settingsGetEnablePrivateBrowsingFunction_Once sync.Once

func settingsGetEnablePrivateBrowsingFunction_Set() error {
	var err error
	settingsGetEnablePrivateBrowsingFunction_Once.Do(func() {
		err = settingsStruct_Set()
		if err != nil {
			return
		}
		settingsGetEnablePrivateBrowsingFunction, err = settingsStruct.InvokerNew("get_enable_private_browsing")
	})
	return err
}

// GetEnablePrivateBrowsing is a representation of the C type webkit_settings_get_enable_private_browsing.
func (recv *Settings) GetEnablePrivateBrowsing() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := settingsGetEnablePrivateBrowsingFunction_Set()
	if err == nil {
		ret = settingsGetEnablePrivateBrowsingFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var settingsGetEnableResizableTextAreasFunction *gi.Function
var settingsGetEnableResizableTextAreasFunction_Once sync.Once

func settingsGetEnableResizableTextAreasFunction_Set() error {
	var err error
	settingsGetEnableResizableTextAreasFunction_Once.Do(func() {
		err = settingsStruct_Set()
		if err != nil {
			return
		}
		settingsGetEnableResizableTextAreasFunction, err = settingsStruct.InvokerNew("get_enable_resizable_text_areas")
	})
	return err
}

// GetEnableResizableTextAreas is a representation of the C type webkit_settings_get_enable_resizable_text_areas.
func (recv *Settings) GetEnableResizableTextAreas() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := settingsGetEnableResizableTextAreasFunction_Set()
	if err == nil {
		ret = settingsGetEnableResizableTextAreasFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var settingsGetEnableSiteSpecificQuirksFunction *gi.Function
var settingsGetEnableSiteSpecificQuirksFunction_Once sync.Once

func settingsGetEnableSiteSpecificQuirksFunction_Set() error {
	var err error
	settingsGetEnableSiteSpecificQuirksFunction_Once.Do(func() {
		err = settingsStruct_Set()
		if err != nil {
			return
		}
		settingsGetEnableSiteSpecificQuirksFunction, err = settingsStruct.InvokerNew("get_enable_site_specific_quirks")
	})
	return err
}

// GetEnableSiteSpecificQuirks is a representation of the C type webkit_settings_get_enable_site_specific_quirks.
func (recv *Settings) GetEnableSiteSpecificQuirks() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := settingsGetEnableSiteSpecificQuirksFunction_Set()
	if err == nil {
		ret = settingsGetEnableSiteSpecificQuirksFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var settingsGetEnableSmoothScrollingFunction *gi.Function
var settingsGetEnableSmoothScrollingFunction_Once sync.Once

func settingsGetEnableSmoothScrollingFunction_Set() error {
	var err error
	settingsGetEnableSmoothScrollingFunction_Once.Do(func() {
		err = settingsStruct_Set()
		if err != nil {
			return
		}
		settingsGetEnableSmoothScrollingFunction, err = settingsStruct.InvokerNew("get_enable_smooth_scrolling")
	})
	return err
}

// GetEnableSmoothScrolling is a representation of the C type webkit_settings_get_enable_smooth_scrolling.
func (recv *Settings) GetEnableSmoothScrolling() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := settingsGetEnableSmoothScrollingFunction_Set()
	if err == nil {
		ret = settingsGetEnableSmoothScrollingFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var settingsGetEnableSpatialNavigationFunction *gi.Function
var settingsGetEnableSpatialNavigationFunction_Once sync.Once

func settingsGetEnableSpatialNavigationFunction_Set() error {
	var err error
	settingsGetEnableSpatialNavigationFunction_Once.Do(func() {
		err = settingsStruct_Set()
		if err != nil {
			return
		}
		settingsGetEnableSpatialNavigationFunction, err = settingsStruct.InvokerNew("get_enable_spatial_navigation")
	})
	return err
}

// GetEnableSpatialNavigation is a representation of the C type webkit_settings_get_enable_spatial_navigation.
func (recv *Settings) GetEnableSpatialNavigation() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := settingsGetEnableSpatialNavigationFunction_Set()
	if err == nil {
		ret = settingsGetEnableSpatialNavigationFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var settingsGetEnableTabsToLinksFunction *gi.Function
var settingsGetEnableTabsToLinksFunction_Once sync.Once

func settingsGetEnableTabsToLinksFunction_Set() error {
	var err error
	settingsGetEnableTabsToLinksFunction_Once.Do(func() {
		err = settingsStruct_Set()
		if err != nil {
			return
		}
		settingsGetEnableTabsToLinksFunction, err = settingsStruct.InvokerNew("get_enable_tabs_to_links")
	})
	return err
}

// GetEnableTabsToLinks is a representation of the C type webkit_settings_get_enable_tabs_to_links.
func (recv *Settings) GetEnableTabsToLinks() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := settingsGetEnableTabsToLinksFunction_Set()
	if err == nil {
		ret = settingsGetEnableTabsToLinksFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var settingsGetEnableWebaudioFunction *gi.Function
var settingsGetEnableWebaudioFunction_Once sync.Once

func settingsGetEnableWebaudioFunction_Set() error {
	var err error
	settingsGetEnableWebaudioFunction_Once.Do(func() {
		err = settingsStruct_Set()
		if err != nil {
			return
		}
		settingsGetEnableWebaudioFunction, err = settingsStruct.InvokerNew("get_enable_webaudio")
	})
	return err
}

// GetEnableWebaudio is a representation of the C type webkit_settings_get_enable_webaudio.
func (recv *Settings) GetEnableWebaudio() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := settingsGetEnableWebaudioFunction_Set()
	if err == nil {
		ret = settingsGetEnableWebaudioFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var settingsGetEnableWebglFunction *gi.Function
var settingsGetEnableWebglFunction_Once sync.Once

func settingsGetEnableWebglFunction_Set() error {
	var err error
	settingsGetEnableWebglFunction_Once.Do(func() {
		err = settingsStruct_Set()
		if err != nil {
			return
		}
		settingsGetEnableWebglFunction, err = settingsStruct.InvokerNew("get_enable_webgl")
	})
	return err
}

// GetEnableWebgl is a representation of the C type webkit_settings_get_enable_webgl.
func (recv *Settings) GetEnableWebgl() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := settingsGetEnableWebglFunction_Set()
	if err == nil {
		ret = settingsGetEnableWebglFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var settingsGetEnableWriteConsoleMessagesToStdoutFunction *gi.Function
var settingsGetEnableWriteConsoleMessagesToStdoutFunction_Once sync.Once

func settingsGetEnableWriteConsoleMessagesToStdoutFunction_Set() error {
	var err error
	settingsGetEnableWriteConsoleMessagesToStdoutFunction_Once.Do(func() {
		err = settingsStruct_Set()
		if err != nil {
			return
		}
		settingsGetEnableWriteConsoleMessagesToStdoutFunction, err = settingsStruct.InvokerNew("get_enable_write_console_messages_to_stdout")
	})
	return err
}

// GetEnableWriteConsoleMessagesToStdout is a representation of the C type webkit_settings_get_enable_write_console_messages_to_stdout.
func (recv *Settings) GetEnableWriteConsoleMessagesToStdout() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := settingsGetEnableWriteConsoleMessagesToStdoutFunction_Set()
	if err == nil {
		ret = settingsGetEnableWriteConsoleMessagesToStdoutFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var settingsGetEnableXssAuditorFunction *gi.Function
var settingsGetEnableXssAuditorFunction_Once sync.Once

func settingsGetEnableXssAuditorFunction_Set() error {
	var err error
	settingsGetEnableXssAuditorFunction_Once.Do(func() {
		err = settingsStruct_Set()
		if err != nil {
			return
		}
		settingsGetEnableXssAuditorFunction, err = settingsStruct.InvokerNew("get_enable_xss_auditor")
	})
	return err
}

// GetEnableXssAuditor is a representation of the C type webkit_settings_get_enable_xss_auditor.
func (recv *Settings) GetEnableXssAuditor() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := settingsGetEnableXssAuditorFunction_Set()
	if err == nil {
		ret = settingsGetEnableXssAuditorFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var settingsGetFantasyFontFamilyFunction *gi.Function
var settingsGetFantasyFontFamilyFunction_Once sync.Once

func settingsGetFantasyFontFamilyFunction_Set() error {
	var err error
	settingsGetFantasyFontFamilyFunction_Once.Do(func() {
		err = settingsStruct_Set()
		if err != nil {
			return
		}
		settingsGetFantasyFontFamilyFunction, err = settingsStruct.InvokerNew("get_fantasy_font_family")
	})
	return err
}

// GetFantasyFontFamily is a representation of the C type webkit_settings_get_fantasy_font_family.
func (recv *Settings) GetFantasyFontFamily() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := settingsGetFantasyFontFamilyFunction_Set()
	if err == nil {
		ret = settingsGetFantasyFontFamilyFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'webkit_settings_get_hardware_acceleration_policy' : return type 'HardwareAccelerationPolicy' not supported

var settingsGetJavascriptCanAccessClipboardFunction *gi.Function
var settingsGetJavascriptCanAccessClipboardFunction_Once sync.Once

func settingsGetJavascriptCanAccessClipboardFunction_Set() error {
	var err error
	settingsGetJavascriptCanAccessClipboardFunction_Once.Do(func() {
		err = settingsStruct_Set()
		if err != nil {
			return
		}
		settingsGetJavascriptCanAccessClipboardFunction, err = settingsStruct.InvokerNew("get_javascript_can_access_clipboard")
	})
	return err
}

// GetJavascriptCanAccessClipboard is a representation of the C type webkit_settings_get_javascript_can_access_clipboard.
func (recv *Settings) GetJavascriptCanAccessClipboard() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := settingsGetJavascriptCanAccessClipboardFunction_Set()
	if err == nil {
		ret = settingsGetJavascriptCanAccessClipboardFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var settingsGetJavascriptCanOpenWindowsAutomaticallyFunction *gi.Function
var settingsGetJavascriptCanOpenWindowsAutomaticallyFunction_Once sync.Once

func settingsGetJavascriptCanOpenWindowsAutomaticallyFunction_Set() error {
	var err error
	settingsGetJavascriptCanOpenWindowsAutomaticallyFunction_Once.Do(func() {
		err = settingsStruct_Set()
		if err != nil {
			return
		}
		settingsGetJavascriptCanOpenWindowsAutomaticallyFunction, err = settingsStruct.InvokerNew("get_javascript_can_open_windows_automatically")
	})
	return err
}

// GetJavascriptCanOpenWindowsAutomatically is a representation of the C type webkit_settings_get_javascript_can_open_windows_automatically.
func (recv *Settings) GetJavascriptCanOpenWindowsAutomatically() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := settingsGetJavascriptCanOpenWindowsAutomaticallyFunction_Set()
	if err == nil {
		ret = settingsGetJavascriptCanOpenWindowsAutomaticallyFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var settingsGetLoadIconsIgnoringImageLoadSettingFunction *gi.Function
var settingsGetLoadIconsIgnoringImageLoadSettingFunction_Once sync.Once

func settingsGetLoadIconsIgnoringImageLoadSettingFunction_Set() error {
	var err error
	settingsGetLoadIconsIgnoringImageLoadSettingFunction_Once.Do(func() {
		err = settingsStruct_Set()
		if err != nil {
			return
		}
		settingsGetLoadIconsIgnoringImageLoadSettingFunction, err = settingsStruct.InvokerNew("get_load_icons_ignoring_image_load_setting")
	})
	return err
}

// GetLoadIconsIgnoringImageLoadSetting is a representation of the C type webkit_settings_get_load_icons_ignoring_image_load_setting.
func (recv *Settings) GetLoadIconsIgnoringImageLoadSetting() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := settingsGetLoadIconsIgnoringImageLoadSettingFunction_Set()
	if err == nil {
		ret = settingsGetLoadIconsIgnoringImageLoadSettingFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var settingsGetMediaPlaybackAllowsInlineFunction *gi.Function
var settingsGetMediaPlaybackAllowsInlineFunction_Once sync.Once

func settingsGetMediaPlaybackAllowsInlineFunction_Set() error {
	var err error
	settingsGetMediaPlaybackAllowsInlineFunction_Once.Do(func() {
		err = settingsStruct_Set()
		if err != nil {
			return
		}
		settingsGetMediaPlaybackAllowsInlineFunction, err = settingsStruct.InvokerNew("get_media_playback_allows_inline")
	})
	return err
}

// GetMediaPlaybackAllowsInline is a representation of the C type webkit_settings_get_media_playback_allows_inline.
func (recv *Settings) GetMediaPlaybackAllowsInline() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := settingsGetMediaPlaybackAllowsInlineFunction_Set()
	if err == nil {
		ret = settingsGetMediaPlaybackAllowsInlineFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var settingsGetMediaPlaybackRequiresUserGestureFunction *gi.Function
var settingsGetMediaPlaybackRequiresUserGestureFunction_Once sync.Once

func settingsGetMediaPlaybackRequiresUserGestureFunction_Set() error {
	var err error
	settingsGetMediaPlaybackRequiresUserGestureFunction_Once.Do(func() {
		err = settingsStruct_Set()
		if err != nil {
			return
		}
		settingsGetMediaPlaybackRequiresUserGestureFunction, err = settingsStruct.InvokerNew("get_media_playback_requires_user_gesture")
	})
	return err
}

// GetMediaPlaybackRequiresUserGesture is a representation of the C type webkit_settings_get_media_playback_requires_user_gesture.
func (recv *Settings) GetMediaPlaybackRequiresUserGesture() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := settingsGetMediaPlaybackRequiresUserGestureFunction_Set()
	if err == nil {
		ret = settingsGetMediaPlaybackRequiresUserGestureFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var settingsGetMinimumFontSizeFunction *gi.Function
var settingsGetMinimumFontSizeFunction_Once sync.Once

func settingsGetMinimumFontSizeFunction_Set() error {
	var err error
	settingsGetMinimumFontSizeFunction_Once.Do(func() {
		err = settingsStruct_Set()
		if err != nil {
			return
		}
		settingsGetMinimumFontSizeFunction, err = settingsStruct.InvokerNew("get_minimum_font_size")
	})
	return err
}

// GetMinimumFontSize is a representation of the C type webkit_settings_get_minimum_font_size.
func (recv *Settings) GetMinimumFontSize() uint32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := settingsGetMinimumFontSizeFunction_Set()
	if err == nil {
		ret = settingsGetMinimumFontSizeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo
}

var settingsGetMonospaceFontFamilyFunction *gi.Function
var settingsGetMonospaceFontFamilyFunction_Once sync.Once

func settingsGetMonospaceFontFamilyFunction_Set() error {
	var err error
	settingsGetMonospaceFontFamilyFunction_Once.Do(func() {
		err = settingsStruct_Set()
		if err != nil {
			return
		}
		settingsGetMonospaceFontFamilyFunction, err = settingsStruct.InvokerNew("get_monospace_font_family")
	})
	return err
}

// GetMonospaceFontFamily is a representation of the C type webkit_settings_get_monospace_font_family.
func (recv *Settings) GetMonospaceFontFamily() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := settingsGetMonospaceFontFamilyFunction_Set()
	if err == nil {
		ret = settingsGetMonospaceFontFamilyFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var settingsGetPictographFontFamilyFunction *gi.Function
var settingsGetPictographFontFamilyFunction_Once sync.Once

func settingsGetPictographFontFamilyFunction_Set() error {
	var err error
	settingsGetPictographFontFamilyFunction_Once.Do(func() {
		err = settingsStruct_Set()
		if err != nil {
			return
		}
		settingsGetPictographFontFamilyFunction, err = settingsStruct.InvokerNew("get_pictograph_font_family")
	})
	return err
}

// GetPictographFontFamily is a representation of the C type webkit_settings_get_pictograph_font_family.
func (recv *Settings) GetPictographFontFamily() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := settingsGetPictographFontFamilyFunction_Set()
	if err == nil {
		ret = settingsGetPictographFontFamilyFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var settingsGetPrintBackgroundsFunction *gi.Function
var settingsGetPrintBackgroundsFunction_Once sync.Once

func settingsGetPrintBackgroundsFunction_Set() error {
	var err error
	settingsGetPrintBackgroundsFunction_Once.Do(func() {
		err = settingsStruct_Set()
		if err != nil {
			return
		}
		settingsGetPrintBackgroundsFunction, err = settingsStruct.InvokerNew("get_print_backgrounds")
	})
	return err
}

// GetPrintBackgrounds is a representation of the C type webkit_settings_get_print_backgrounds.
func (recv *Settings) GetPrintBackgrounds() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := settingsGetPrintBackgroundsFunction_Set()
	if err == nil {
		ret = settingsGetPrintBackgroundsFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var settingsGetSansSerifFontFamilyFunction *gi.Function
var settingsGetSansSerifFontFamilyFunction_Once sync.Once

func settingsGetSansSerifFontFamilyFunction_Set() error {
	var err error
	settingsGetSansSerifFontFamilyFunction_Once.Do(func() {
		err = settingsStruct_Set()
		if err != nil {
			return
		}
		settingsGetSansSerifFontFamilyFunction, err = settingsStruct.InvokerNew("get_sans_serif_font_family")
	})
	return err
}

// GetSansSerifFontFamily is a representation of the C type webkit_settings_get_sans_serif_font_family.
func (recv *Settings) GetSansSerifFontFamily() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := settingsGetSansSerifFontFamilyFunction_Set()
	if err == nil {
		ret = settingsGetSansSerifFontFamilyFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var settingsGetSerifFontFamilyFunction *gi.Function
var settingsGetSerifFontFamilyFunction_Once sync.Once

func settingsGetSerifFontFamilyFunction_Set() error {
	var err error
	settingsGetSerifFontFamilyFunction_Once.Do(func() {
		err = settingsStruct_Set()
		if err != nil {
			return
		}
		settingsGetSerifFontFamilyFunction, err = settingsStruct.InvokerNew("get_serif_font_family")
	})
	return err
}

// GetSerifFontFamily is a representation of the C type webkit_settings_get_serif_font_family.
func (recv *Settings) GetSerifFontFamily() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := settingsGetSerifFontFamilyFunction_Set()
	if err == nil {
		ret = settingsGetSerifFontFamilyFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var settingsGetUserAgentFunction *gi.Function
var settingsGetUserAgentFunction_Once sync.Once

func settingsGetUserAgentFunction_Set() error {
	var err error
	settingsGetUserAgentFunction_Once.Do(func() {
		err = settingsStruct_Set()
		if err != nil {
			return
		}
		settingsGetUserAgentFunction, err = settingsStruct.InvokerNew("get_user_agent")
	})
	return err
}

// GetUserAgent is a representation of the C type webkit_settings_get_user_agent.
func (recv *Settings) GetUserAgent() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := settingsGetUserAgentFunction_Set()
	if err == nil {
		ret = settingsGetUserAgentFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var settingsGetZoomTextOnlyFunction *gi.Function
var settingsGetZoomTextOnlyFunction_Once sync.Once

func settingsGetZoomTextOnlyFunction_Set() error {
	var err error
	settingsGetZoomTextOnlyFunction_Once.Do(func() {
		err = settingsStruct_Set()
		if err != nil {
			return
		}
		settingsGetZoomTextOnlyFunction, err = settingsStruct.InvokerNew("get_zoom_text_only")
	})
	return err
}

// GetZoomTextOnly is a representation of the C type webkit_settings_get_zoom_text_only.
func (recv *Settings) GetZoomTextOnly() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := settingsGetZoomTextOnlyFunction_Set()
	if err == nil {
		ret = settingsGetZoomTextOnlyFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var settingsSetAllowFileAccessFromFileUrlsFunction *gi.Function
var settingsSetAllowFileAccessFromFileUrlsFunction_Once sync.Once

func settingsSetAllowFileAccessFromFileUrlsFunction_Set() error {
	var err error
	settingsSetAllowFileAccessFromFileUrlsFunction_Once.Do(func() {
		err = settingsStruct_Set()
		if err != nil {
			return
		}
		settingsSetAllowFileAccessFromFileUrlsFunction, err = settingsStruct.InvokerNew("set_allow_file_access_from_file_urls")
	})
	return err
}

// SetAllowFileAccessFromFileUrls is a representation of the C type webkit_settings_set_allow_file_access_from_file_urls.
func (recv *Settings) SetAllowFileAccessFromFileUrls(allowed bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetBoolean(allowed)

	err := settingsSetAllowFileAccessFromFileUrlsFunction_Set()
	if err == nil {
		settingsSetAllowFileAccessFromFileUrlsFunction.Invoke(inArgs[:], nil)
	}

	return
}

var settingsSetAllowModalDialogsFunction *gi.Function
var settingsSetAllowModalDialogsFunction_Once sync.Once

func settingsSetAllowModalDialogsFunction_Set() error {
	var err error
	settingsSetAllowModalDialogsFunction_Once.Do(func() {
		err = settingsStruct_Set()
		if err != nil {
			return
		}
		settingsSetAllowModalDialogsFunction, err = settingsStruct.InvokerNew("set_allow_modal_dialogs")
	})
	return err
}

// SetAllowModalDialogs is a representation of the C type webkit_settings_set_allow_modal_dialogs.
func (recv *Settings) SetAllowModalDialogs(allowed bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetBoolean(allowed)

	err := settingsSetAllowModalDialogsFunction_Set()
	if err == nil {
		settingsSetAllowModalDialogsFunction.Invoke(inArgs[:], nil)
	}

	return
}

var settingsSetAllowUniversalAccessFromFileUrlsFunction *gi.Function
var settingsSetAllowUniversalAccessFromFileUrlsFunction_Once sync.Once

func settingsSetAllowUniversalAccessFromFileUrlsFunction_Set() error {
	var err error
	settingsSetAllowUniversalAccessFromFileUrlsFunction_Once.Do(func() {
		err = settingsStruct_Set()
		if err != nil {
			return
		}
		settingsSetAllowUniversalAccessFromFileUrlsFunction, err = settingsStruct.InvokerNew("set_allow_universal_access_from_file_urls")
	})
	return err
}

// SetAllowUniversalAccessFromFileUrls is a representation of the C type webkit_settings_set_allow_universal_access_from_file_urls.
func (recv *Settings) SetAllowUniversalAccessFromFileUrls(allowed bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetBoolean(allowed)

	err := settingsSetAllowUniversalAccessFromFileUrlsFunction_Set()
	if err == nil {
		settingsSetAllowUniversalAccessFromFileUrlsFunction.Invoke(inArgs[:], nil)
	}

	return
}

var settingsSetAutoLoadImagesFunction *gi.Function
var settingsSetAutoLoadImagesFunction_Once sync.Once

func settingsSetAutoLoadImagesFunction_Set() error {
	var err error
	settingsSetAutoLoadImagesFunction_Once.Do(func() {
		err = settingsStruct_Set()
		if err != nil {
			return
		}
		settingsSetAutoLoadImagesFunction, err = settingsStruct.InvokerNew("set_auto_load_images")
	})
	return err
}

// SetAutoLoadImages is a representation of the C type webkit_settings_set_auto_load_images.
func (recv *Settings) SetAutoLoadImages(enabled bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetBoolean(enabled)

	err := settingsSetAutoLoadImagesFunction_Set()
	if err == nil {
		settingsSetAutoLoadImagesFunction.Invoke(inArgs[:], nil)
	}

	return
}

var settingsSetCursiveFontFamilyFunction *gi.Function
var settingsSetCursiveFontFamilyFunction_Once sync.Once

func settingsSetCursiveFontFamilyFunction_Set() error {
	var err error
	settingsSetCursiveFontFamilyFunction_Once.Do(func() {
		err = settingsStruct_Set()
		if err != nil {
			return
		}
		settingsSetCursiveFontFamilyFunction, err = settingsStruct.InvokerNew("set_cursive_font_family")
	})
	return err
}

// SetCursiveFontFamily is a representation of the C type webkit_settings_set_cursive_font_family.
func (recv *Settings) SetCursiveFontFamily(cursiveFontFamily string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(cursiveFontFamily)

	err := settingsSetCursiveFontFamilyFunction_Set()
	if err == nil {
		settingsSetCursiveFontFamilyFunction.Invoke(inArgs[:], nil)
	}

	return
}

var settingsSetDefaultCharsetFunction *gi.Function
var settingsSetDefaultCharsetFunction_Once sync.Once

func settingsSetDefaultCharsetFunction_Set() error {
	var err error
	settingsSetDefaultCharsetFunction_Once.Do(func() {
		err = settingsStruct_Set()
		if err != nil {
			return
		}
		settingsSetDefaultCharsetFunction, err = settingsStruct.InvokerNew("set_default_charset")
	})
	return err
}

// SetDefaultCharset is a representation of the C type webkit_settings_set_default_charset.
func (recv *Settings) SetDefaultCharset(defaultCharset string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(defaultCharset)

	err := settingsSetDefaultCharsetFunction_Set()
	if err == nil {
		settingsSetDefaultCharsetFunction.Invoke(inArgs[:], nil)
	}

	return
}

var settingsSetDefaultFontFamilyFunction *gi.Function
var settingsSetDefaultFontFamilyFunction_Once sync.Once

func settingsSetDefaultFontFamilyFunction_Set() error {
	var err error
	settingsSetDefaultFontFamilyFunction_Once.Do(func() {
		err = settingsStruct_Set()
		if err != nil {
			return
		}
		settingsSetDefaultFontFamilyFunction, err = settingsStruct.InvokerNew("set_default_font_family")
	})
	return err
}

// SetDefaultFontFamily is a representation of the C type webkit_settings_set_default_font_family.
func (recv *Settings) SetDefaultFontFamily(defaultFontFamily string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(defaultFontFamily)

	err := settingsSetDefaultFontFamilyFunction_Set()
	if err == nil {
		settingsSetDefaultFontFamilyFunction.Invoke(inArgs[:], nil)
	}

	return
}

var settingsSetDefaultFontSizeFunction *gi.Function
var settingsSetDefaultFontSizeFunction_Once sync.Once

func settingsSetDefaultFontSizeFunction_Set() error {
	var err error
	settingsSetDefaultFontSizeFunction_Once.Do(func() {
		err = settingsStruct_Set()
		if err != nil {
			return
		}
		settingsSetDefaultFontSizeFunction, err = settingsStruct.InvokerNew("set_default_font_size")
	})
	return err
}

// SetDefaultFontSize is a representation of the C type webkit_settings_set_default_font_size.
func (recv *Settings) SetDefaultFontSize(fontSize uint32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetUint32(fontSize)

	err := settingsSetDefaultFontSizeFunction_Set()
	if err == nil {
		settingsSetDefaultFontSizeFunction.Invoke(inArgs[:], nil)
	}

	return
}

var settingsSetDefaultMonospaceFontSizeFunction *gi.Function
var settingsSetDefaultMonospaceFontSizeFunction_Once sync.Once

func settingsSetDefaultMonospaceFontSizeFunction_Set() error {
	var err error
	settingsSetDefaultMonospaceFontSizeFunction_Once.Do(func() {
		err = settingsStruct_Set()
		if err != nil {
			return
		}
		settingsSetDefaultMonospaceFontSizeFunction, err = settingsStruct.InvokerNew("set_default_monospace_font_size")
	})
	return err
}

// SetDefaultMonospaceFontSize is a representation of the C type webkit_settings_set_default_monospace_font_size.
func (recv *Settings) SetDefaultMonospaceFontSize(fontSize uint32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetUint32(fontSize)

	err := settingsSetDefaultMonospaceFontSizeFunction_Set()
	if err == nil {
		settingsSetDefaultMonospaceFontSizeFunction.Invoke(inArgs[:], nil)
	}

	return
}

var settingsSetDrawCompositingIndicatorsFunction *gi.Function
var settingsSetDrawCompositingIndicatorsFunction_Once sync.Once

func settingsSetDrawCompositingIndicatorsFunction_Set() error {
	var err error
	settingsSetDrawCompositingIndicatorsFunction_Once.Do(func() {
		err = settingsStruct_Set()
		if err != nil {
			return
		}
		settingsSetDrawCompositingIndicatorsFunction, err = settingsStruct.InvokerNew("set_draw_compositing_indicators")
	})
	return err
}

// SetDrawCompositingIndicators is a representation of the C type webkit_settings_set_draw_compositing_indicators.
func (recv *Settings) SetDrawCompositingIndicators(enabled bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetBoolean(enabled)

	err := settingsSetDrawCompositingIndicatorsFunction_Set()
	if err == nil {
		settingsSetDrawCompositingIndicatorsFunction.Invoke(inArgs[:], nil)
	}

	return
}

var settingsSetEnableAccelerated2dCanvasFunction *gi.Function
var settingsSetEnableAccelerated2dCanvasFunction_Once sync.Once

func settingsSetEnableAccelerated2dCanvasFunction_Set() error {
	var err error
	settingsSetEnableAccelerated2dCanvasFunction_Once.Do(func() {
		err = settingsStruct_Set()
		if err != nil {
			return
		}
		settingsSetEnableAccelerated2dCanvasFunction, err = settingsStruct.InvokerNew("set_enable_accelerated_2d_canvas")
	})
	return err
}

// SetEnableAccelerated2dCanvas is a representation of the C type webkit_settings_set_enable_accelerated_2d_canvas.
func (recv *Settings) SetEnableAccelerated2dCanvas(enabled bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetBoolean(enabled)

	err := settingsSetEnableAccelerated2dCanvasFunction_Set()
	if err == nil {
		settingsSetEnableAccelerated2dCanvasFunction.Invoke(inArgs[:], nil)
	}

	return
}

var settingsSetEnableBackForwardNavigationGesturesFunction *gi.Function
var settingsSetEnableBackForwardNavigationGesturesFunction_Once sync.Once

func settingsSetEnableBackForwardNavigationGesturesFunction_Set() error {
	var err error
	settingsSetEnableBackForwardNavigationGesturesFunction_Once.Do(func() {
		err = settingsStruct_Set()
		if err != nil {
			return
		}
		settingsSetEnableBackForwardNavigationGesturesFunction, err = settingsStruct.InvokerNew("set_enable_back_forward_navigation_gestures")
	})
	return err
}

// SetEnableBackForwardNavigationGestures is a representation of the C type webkit_settings_set_enable_back_forward_navigation_gestures.
func (recv *Settings) SetEnableBackForwardNavigationGestures(enabled bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetBoolean(enabled)

	err := settingsSetEnableBackForwardNavigationGesturesFunction_Set()
	if err == nil {
		settingsSetEnableBackForwardNavigationGesturesFunction.Invoke(inArgs[:], nil)
	}

	return
}

var settingsSetEnableCaretBrowsingFunction *gi.Function
var settingsSetEnableCaretBrowsingFunction_Once sync.Once

func settingsSetEnableCaretBrowsingFunction_Set() error {
	var err error
	settingsSetEnableCaretBrowsingFunction_Once.Do(func() {
		err = settingsStruct_Set()
		if err != nil {
			return
		}
		settingsSetEnableCaretBrowsingFunction, err = settingsStruct.InvokerNew("set_enable_caret_browsing")
	})
	return err
}

// SetEnableCaretBrowsing is a representation of the C type webkit_settings_set_enable_caret_browsing.
func (recv *Settings) SetEnableCaretBrowsing(enabled bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetBoolean(enabled)

	err := settingsSetEnableCaretBrowsingFunction_Set()
	if err == nil {
		settingsSetEnableCaretBrowsingFunction.Invoke(inArgs[:], nil)
	}

	return
}

var settingsSetEnableDeveloperExtrasFunction *gi.Function
var settingsSetEnableDeveloperExtrasFunction_Once sync.Once

func settingsSetEnableDeveloperExtrasFunction_Set() error {
	var err error
	settingsSetEnableDeveloperExtrasFunction_Once.Do(func() {
		err = settingsStruct_Set()
		if err != nil {
			return
		}
		settingsSetEnableDeveloperExtrasFunction, err = settingsStruct.InvokerNew("set_enable_developer_extras")
	})
	return err
}

// SetEnableDeveloperExtras is a representation of the C type webkit_settings_set_enable_developer_extras.
func (recv *Settings) SetEnableDeveloperExtras(enabled bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetBoolean(enabled)

	err := settingsSetEnableDeveloperExtrasFunction_Set()
	if err == nil {
		settingsSetEnableDeveloperExtrasFunction.Invoke(inArgs[:], nil)
	}

	return
}

var settingsSetEnableDnsPrefetchingFunction *gi.Function
var settingsSetEnableDnsPrefetchingFunction_Once sync.Once

func settingsSetEnableDnsPrefetchingFunction_Set() error {
	var err error
	settingsSetEnableDnsPrefetchingFunction_Once.Do(func() {
		err = settingsStruct_Set()
		if err != nil {
			return
		}
		settingsSetEnableDnsPrefetchingFunction, err = settingsStruct.InvokerNew("set_enable_dns_prefetching")
	})
	return err
}

// SetEnableDnsPrefetching is a representation of the C type webkit_settings_set_enable_dns_prefetching.
func (recv *Settings) SetEnableDnsPrefetching(enabled bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetBoolean(enabled)

	err := settingsSetEnableDnsPrefetchingFunction_Set()
	if err == nil {
		settingsSetEnableDnsPrefetchingFunction.Invoke(inArgs[:], nil)
	}

	return
}

var settingsSetEnableEncryptedMediaFunction *gi.Function
var settingsSetEnableEncryptedMediaFunction_Once sync.Once

func settingsSetEnableEncryptedMediaFunction_Set() error {
	var err error
	settingsSetEnableEncryptedMediaFunction_Once.Do(func() {
		err = settingsStruct_Set()
		if err != nil {
			return
		}
		settingsSetEnableEncryptedMediaFunction, err = settingsStruct.InvokerNew("set_enable_encrypted_media")
	})
	return err
}

// SetEnableEncryptedMedia is a representation of the C type webkit_settings_set_enable_encrypted_media.
func (recv *Settings) SetEnableEncryptedMedia(enabled bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetBoolean(enabled)

	err := settingsSetEnableEncryptedMediaFunction_Set()
	if err == nil {
		settingsSetEnableEncryptedMediaFunction.Invoke(inArgs[:], nil)
	}

	return
}

var settingsSetEnableFrameFlatteningFunction *gi.Function
var settingsSetEnableFrameFlatteningFunction_Once sync.Once

func settingsSetEnableFrameFlatteningFunction_Set() error {
	var err error
	settingsSetEnableFrameFlatteningFunction_Once.Do(func() {
		err = settingsStruct_Set()
		if err != nil {
			return
		}
		settingsSetEnableFrameFlatteningFunction, err = settingsStruct.InvokerNew("set_enable_frame_flattening")
	})
	return err
}

// SetEnableFrameFlattening is a representation of the C type webkit_settings_set_enable_frame_flattening.
func (recv *Settings) SetEnableFrameFlattening(enabled bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetBoolean(enabled)

	err := settingsSetEnableFrameFlatteningFunction_Set()
	if err == nil {
		settingsSetEnableFrameFlatteningFunction.Invoke(inArgs[:], nil)
	}

	return
}

var settingsSetEnableFullscreenFunction *gi.Function
var settingsSetEnableFullscreenFunction_Once sync.Once

func settingsSetEnableFullscreenFunction_Set() error {
	var err error
	settingsSetEnableFullscreenFunction_Once.Do(func() {
		err = settingsStruct_Set()
		if err != nil {
			return
		}
		settingsSetEnableFullscreenFunction, err = settingsStruct.InvokerNew("set_enable_fullscreen")
	})
	return err
}

// SetEnableFullscreen is a representation of the C type webkit_settings_set_enable_fullscreen.
func (recv *Settings) SetEnableFullscreen(enabled bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetBoolean(enabled)

	err := settingsSetEnableFullscreenFunction_Set()
	if err == nil {
		settingsSetEnableFullscreenFunction.Invoke(inArgs[:], nil)
	}

	return
}

var settingsSetEnableHtml5DatabaseFunction *gi.Function
var settingsSetEnableHtml5DatabaseFunction_Once sync.Once

func settingsSetEnableHtml5DatabaseFunction_Set() error {
	var err error
	settingsSetEnableHtml5DatabaseFunction_Once.Do(func() {
		err = settingsStruct_Set()
		if err != nil {
			return
		}
		settingsSetEnableHtml5DatabaseFunction, err = settingsStruct.InvokerNew("set_enable_html5_database")
	})
	return err
}

// SetEnableHtml5Database is a representation of the C type webkit_settings_set_enable_html5_database.
func (recv *Settings) SetEnableHtml5Database(enabled bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetBoolean(enabled)

	err := settingsSetEnableHtml5DatabaseFunction_Set()
	if err == nil {
		settingsSetEnableHtml5DatabaseFunction.Invoke(inArgs[:], nil)
	}

	return
}

var settingsSetEnableHtml5LocalStorageFunction *gi.Function
var settingsSetEnableHtml5LocalStorageFunction_Once sync.Once

func settingsSetEnableHtml5LocalStorageFunction_Set() error {
	var err error
	settingsSetEnableHtml5LocalStorageFunction_Once.Do(func() {
		err = settingsStruct_Set()
		if err != nil {
			return
		}
		settingsSetEnableHtml5LocalStorageFunction, err = settingsStruct.InvokerNew("set_enable_html5_local_storage")
	})
	return err
}

// SetEnableHtml5LocalStorage is a representation of the C type webkit_settings_set_enable_html5_local_storage.
func (recv *Settings) SetEnableHtml5LocalStorage(enabled bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetBoolean(enabled)

	err := settingsSetEnableHtml5LocalStorageFunction_Set()
	if err == nil {
		settingsSetEnableHtml5LocalStorageFunction.Invoke(inArgs[:], nil)
	}

	return
}

var settingsSetEnableHyperlinkAuditingFunction *gi.Function
var settingsSetEnableHyperlinkAuditingFunction_Once sync.Once

func settingsSetEnableHyperlinkAuditingFunction_Set() error {
	var err error
	settingsSetEnableHyperlinkAuditingFunction_Once.Do(func() {
		err = settingsStruct_Set()
		if err != nil {
			return
		}
		settingsSetEnableHyperlinkAuditingFunction, err = settingsStruct.InvokerNew("set_enable_hyperlink_auditing")
	})
	return err
}

// SetEnableHyperlinkAuditing is a representation of the C type webkit_settings_set_enable_hyperlink_auditing.
func (recv *Settings) SetEnableHyperlinkAuditing(enabled bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetBoolean(enabled)

	err := settingsSetEnableHyperlinkAuditingFunction_Set()
	if err == nil {
		settingsSetEnableHyperlinkAuditingFunction.Invoke(inArgs[:], nil)
	}

	return
}

var settingsSetEnableJavaFunction *gi.Function
var settingsSetEnableJavaFunction_Once sync.Once

func settingsSetEnableJavaFunction_Set() error {
	var err error
	settingsSetEnableJavaFunction_Once.Do(func() {
		err = settingsStruct_Set()
		if err != nil {
			return
		}
		settingsSetEnableJavaFunction, err = settingsStruct.InvokerNew("set_enable_java")
	})
	return err
}

// SetEnableJava is a representation of the C type webkit_settings_set_enable_java.
func (recv *Settings) SetEnableJava(enabled bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetBoolean(enabled)

	err := settingsSetEnableJavaFunction_Set()
	if err == nil {
		settingsSetEnableJavaFunction.Invoke(inArgs[:], nil)
	}

	return
}

var settingsSetEnableJavascriptFunction *gi.Function
var settingsSetEnableJavascriptFunction_Once sync.Once

func settingsSetEnableJavascriptFunction_Set() error {
	var err error
	settingsSetEnableJavascriptFunction_Once.Do(func() {
		err = settingsStruct_Set()
		if err != nil {
			return
		}
		settingsSetEnableJavascriptFunction, err = settingsStruct.InvokerNew("set_enable_javascript")
	})
	return err
}

// SetEnableJavascript is a representation of the C type webkit_settings_set_enable_javascript.
func (recv *Settings) SetEnableJavascript(enabled bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetBoolean(enabled)

	err := settingsSetEnableJavascriptFunction_Set()
	if err == nil {
		settingsSetEnableJavascriptFunction.Invoke(inArgs[:], nil)
	}

	return
}

var settingsSetEnableJavascriptMarkupFunction *gi.Function
var settingsSetEnableJavascriptMarkupFunction_Once sync.Once

func settingsSetEnableJavascriptMarkupFunction_Set() error {
	var err error
	settingsSetEnableJavascriptMarkupFunction_Once.Do(func() {
		err = settingsStruct_Set()
		if err != nil {
			return
		}
		settingsSetEnableJavascriptMarkupFunction, err = settingsStruct.InvokerNew("set_enable_javascript_markup")
	})
	return err
}

// SetEnableJavascriptMarkup is a representation of the C type webkit_settings_set_enable_javascript_markup.
func (recv *Settings) SetEnableJavascriptMarkup(enabled bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetBoolean(enabled)

	err := settingsSetEnableJavascriptMarkupFunction_Set()
	if err == nil {
		settingsSetEnableJavascriptMarkupFunction.Invoke(inArgs[:], nil)
	}

	return
}

var settingsSetEnableMediaFunction *gi.Function
var settingsSetEnableMediaFunction_Once sync.Once

func settingsSetEnableMediaFunction_Set() error {
	var err error
	settingsSetEnableMediaFunction_Once.Do(func() {
		err = settingsStruct_Set()
		if err != nil {
			return
		}
		settingsSetEnableMediaFunction, err = settingsStruct.InvokerNew("set_enable_media")
	})
	return err
}

// SetEnableMedia is a representation of the C type webkit_settings_set_enable_media.
func (recv *Settings) SetEnableMedia(enabled bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetBoolean(enabled)

	err := settingsSetEnableMediaFunction_Set()
	if err == nil {
		settingsSetEnableMediaFunction.Invoke(inArgs[:], nil)
	}

	return
}

var settingsSetEnableMediaCapabilitiesFunction *gi.Function
var settingsSetEnableMediaCapabilitiesFunction_Once sync.Once

func settingsSetEnableMediaCapabilitiesFunction_Set() error {
	var err error
	settingsSetEnableMediaCapabilitiesFunction_Once.Do(func() {
		err = settingsStruct_Set()
		if err != nil {
			return
		}
		settingsSetEnableMediaCapabilitiesFunction, err = settingsStruct.InvokerNew("set_enable_media_capabilities")
	})
	return err
}

// SetEnableMediaCapabilities is a representation of the C type webkit_settings_set_enable_media_capabilities.
func (recv *Settings) SetEnableMediaCapabilities(enabled bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetBoolean(enabled)

	err := settingsSetEnableMediaCapabilitiesFunction_Set()
	if err == nil {
		settingsSetEnableMediaCapabilitiesFunction.Invoke(inArgs[:], nil)
	}

	return
}

var settingsSetEnableMediaStreamFunction *gi.Function
var settingsSetEnableMediaStreamFunction_Once sync.Once

func settingsSetEnableMediaStreamFunction_Set() error {
	var err error
	settingsSetEnableMediaStreamFunction_Once.Do(func() {
		err = settingsStruct_Set()
		if err != nil {
			return
		}
		settingsSetEnableMediaStreamFunction, err = settingsStruct.InvokerNew("set_enable_media_stream")
	})
	return err
}

// SetEnableMediaStream is a representation of the C type webkit_settings_set_enable_media_stream.
func (recv *Settings) SetEnableMediaStream(enabled bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetBoolean(enabled)

	err := settingsSetEnableMediaStreamFunction_Set()
	if err == nil {
		settingsSetEnableMediaStreamFunction.Invoke(inArgs[:], nil)
	}

	return
}

var settingsSetEnableMediasourceFunction *gi.Function
var settingsSetEnableMediasourceFunction_Once sync.Once

func settingsSetEnableMediasourceFunction_Set() error {
	var err error
	settingsSetEnableMediasourceFunction_Once.Do(func() {
		err = settingsStruct_Set()
		if err != nil {
			return
		}
		settingsSetEnableMediasourceFunction, err = settingsStruct.InvokerNew("set_enable_mediasource")
	})
	return err
}

// SetEnableMediasource is a representation of the C type webkit_settings_set_enable_mediasource.
func (recv *Settings) SetEnableMediasource(enabled bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetBoolean(enabled)

	err := settingsSetEnableMediasourceFunction_Set()
	if err == nil {
		settingsSetEnableMediasourceFunction.Invoke(inArgs[:], nil)
	}

	return
}

var settingsSetEnableMockCaptureDevicesFunction *gi.Function
var settingsSetEnableMockCaptureDevicesFunction_Once sync.Once

func settingsSetEnableMockCaptureDevicesFunction_Set() error {
	var err error
	settingsSetEnableMockCaptureDevicesFunction_Once.Do(func() {
		err = settingsStruct_Set()
		if err != nil {
			return
		}
		settingsSetEnableMockCaptureDevicesFunction, err = settingsStruct.InvokerNew("set_enable_mock_capture_devices")
	})
	return err
}

// SetEnableMockCaptureDevices is a representation of the C type webkit_settings_set_enable_mock_capture_devices.
func (recv *Settings) SetEnableMockCaptureDevices(enabled bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetBoolean(enabled)

	err := settingsSetEnableMockCaptureDevicesFunction_Set()
	if err == nil {
		settingsSetEnableMockCaptureDevicesFunction.Invoke(inArgs[:], nil)
	}

	return
}

var settingsSetEnableOfflineWebApplicationCacheFunction *gi.Function
var settingsSetEnableOfflineWebApplicationCacheFunction_Once sync.Once

func settingsSetEnableOfflineWebApplicationCacheFunction_Set() error {
	var err error
	settingsSetEnableOfflineWebApplicationCacheFunction_Once.Do(func() {
		err = settingsStruct_Set()
		if err != nil {
			return
		}
		settingsSetEnableOfflineWebApplicationCacheFunction, err = settingsStruct.InvokerNew("set_enable_offline_web_application_cache")
	})
	return err
}

// SetEnableOfflineWebApplicationCache is a representation of the C type webkit_settings_set_enable_offline_web_application_cache.
func (recv *Settings) SetEnableOfflineWebApplicationCache(enabled bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetBoolean(enabled)

	err := settingsSetEnableOfflineWebApplicationCacheFunction_Set()
	if err == nil {
		settingsSetEnableOfflineWebApplicationCacheFunction.Invoke(inArgs[:], nil)
	}

	return
}

var settingsSetEnablePageCacheFunction *gi.Function
var settingsSetEnablePageCacheFunction_Once sync.Once

func settingsSetEnablePageCacheFunction_Set() error {
	var err error
	settingsSetEnablePageCacheFunction_Once.Do(func() {
		err = settingsStruct_Set()
		if err != nil {
			return
		}
		settingsSetEnablePageCacheFunction, err = settingsStruct.InvokerNew("set_enable_page_cache")
	})
	return err
}

// SetEnablePageCache is a representation of the C type webkit_settings_set_enable_page_cache.
func (recv *Settings) SetEnablePageCache(enabled bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetBoolean(enabled)

	err := settingsSetEnablePageCacheFunction_Set()
	if err == nil {
		settingsSetEnablePageCacheFunction.Invoke(inArgs[:], nil)
	}

	return
}

var settingsSetEnablePluginsFunction *gi.Function
var settingsSetEnablePluginsFunction_Once sync.Once

func settingsSetEnablePluginsFunction_Set() error {
	var err error
	settingsSetEnablePluginsFunction_Once.Do(func() {
		err = settingsStruct_Set()
		if err != nil {
			return
		}
		settingsSetEnablePluginsFunction, err = settingsStruct.InvokerNew("set_enable_plugins")
	})
	return err
}

// SetEnablePlugins is a representation of the C type webkit_settings_set_enable_plugins.
func (recv *Settings) SetEnablePlugins(enabled bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetBoolean(enabled)

	err := settingsSetEnablePluginsFunction_Set()
	if err == nil {
		settingsSetEnablePluginsFunction.Invoke(inArgs[:], nil)
	}

	return
}

var settingsSetEnablePrivateBrowsingFunction *gi.Function
var settingsSetEnablePrivateBrowsingFunction_Once sync.Once

func settingsSetEnablePrivateBrowsingFunction_Set() error {
	var err error
	settingsSetEnablePrivateBrowsingFunction_Once.Do(func() {
		err = settingsStruct_Set()
		if err != nil {
			return
		}
		settingsSetEnablePrivateBrowsingFunction, err = settingsStruct.InvokerNew("set_enable_private_browsing")
	})
	return err
}

// SetEnablePrivateBrowsing is a representation of the C type webkit_settings_set_enable_private_browsing.
func (recv *Settings) SetEnablePrivateBrowsing(enabled bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetBoolean(enabled)

	err := settingsSetEnablePrivateBrowsingFunction_Set()
	if err == nil {
		settingsSetEnablePrivateBrowsingFunction.Invoke(inArgs[:], nil)
	}

	return
}

var settingsSetEnableResizableTextAreasFunction *gi.Function
var settingsSetEnableResizableTextAreasFunction_Once sync.Once

func settingsSetEnableResizableTextAreasFunction_Set() error {
	var err error
	settingsSetEnableResizableTextAreasFunction_Once.Do(func() {
		err = settingsStruct_Set()
		if err != nil {
			return
		}
		settingsSetEnableResizableTextAreasFunction, err = settingsStruct.InvokerNew("set_enable_resizable_text_areas")
	})
	return err
}

// SetEnableResizableTextAreas is a representation of the C type webkit_settings_set_enable_resizable_text_areas.
func (recv *Settings) SetEnableResizableTextAreas(enabled bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetBoolean(enabled)

	err := settingsSetEnableResizableTextAreasFunction_Set()
	if err == nil {
		settingsSetEnableResizableTextAreasFunction.Invoke(inArgs[:], nil)
	}

	return
}

var settingsSetEnableSiteSpecificQuirksFunction *gi.Function
var settingsSetEnableSiteSpecificQuirksFunction_Once sync.Once

func settingsSetEnableSiteSpecificQuirksFunction_Set() error {
	var err error
	settingsSetEnableSiteSpecificQuirksFunction_Once.Do(func() {
		err = settingsStruct_Set()
		if err != nil {
			return
		}
		settingsSetEnableSiteSpecificQuirksFunction, err = settingsStruct.InvokerNew("set_enable_site_specific_quirks")
	})
	return err
}

// SetEnableSiteSpecificQuirks is a representation of the C type webkit_settings_set_enable_site_specific_quirks.
func (recv *Settings) SetEnableSiteSpecificQuirks(enabled bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetBoolean(enabled)

	err := settingsSetEnableSiteSpecificQuirksFunction_Set()
	if err == nil {
		settingsSetEnableSiteSpecificQuirksFunction.Invoke(inArgs[:], nil)
	}

	return
}

var settingsSetEnableSmoothScrollingFunction *gi.Function
var settingsSetEnableSmoothScrollingFunction_Once sync.Once

func settingsSetEnableSmoothScrollingFunction_Set() error {
	var err error
	settingsSetEnableSmoothScrollingFunction_Once.Do(func() {
		err = settingsStruct_Set()
		if err != nil {
			return
		}
		settingsSetEnableSmoothScrollingFunction, err = settingsStruct.InvokerNew("set_enable_smooth_scrolling")
	})
	return err
}

// SetEnableSmoothScrolling is a representation of the C type webkit_settings_set_enable_smooth_scrolling.
func (recv *Settings) SetEnableSmoothScrolling(enabled bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetBoolean(enabled)

	err := settingsSetEnableSmoothScrollingFunction_Set()
	if err == nil {
		settingsSetEnableSmoothScrollingFunction.Invoke(inArgs[:], nil)
	}

	return
}

var settingsSetEnableSpatialNavigationFunction *gi.Function
var settingsSetEnableSpatialNavigationFunction_Once sync.Once

func settingsSetEnableSpatialNavigationFunction_Set() error {
	var err error
	settingsSetEnableSpatialNavigationFunction_Once.Do(func() {
		err = settingsStruct_Set()
		if err != nil {
			return
		}
		settingsSetEnableSpatialNavigationFunction, err = settingsStruct.InvokerNew("set_enable_spatial_navigation")
	})
	return err
}

// SetEnableSpatialNavigation is a representation of the C type webkit_settings_set_enable_spatial_navigation.
func (recv *Settings) SetEnableSpatialNavigation(enabled bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetBoolean(enabled)

	err := settingsSetEnableSpatialNavigationFunction_Set()
	if err == nil {
		settingsSetEnableSpatialNavigationFunction.Invoke(inArgs[:], nil)
	}

	return
}

var settingsSetEnableTabsToLinksFunction *gi.Function
var settingsSetEnableTabsToLinksFunction_Once sync.Once

func settingsSetEnableTabsToLinksFunction_Set() error {
	var err error
	settingsSetEnableTabsToLinksFunction_Once.Do(func() {
		err = settingsStruct_Set()
		if err != nil {
			return
		}
		settingsSetEnableTabsToLinksFunction, err = settingsStruct.InvokerNew("set_enable_tabs_to_links")
	})
	return err
}

// SetEnableTabsToLinks is a representation of the C type webkit_settings_set_enable_tabs_to_links.
func (recv *Settings) SetEnableTabsToLinks(enabled bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetBoolean(enabled)

	err := settingsSetEnableTabsToLinksFunction_Set()
	if err == nil {
		settingsSetEnableTabsToLinksFunction.Invoke(inArgs[:], nil)
	}

	return
}

var settingsSetEnableWebaudioFunction *gi.Function
var settingsSetEnableWebaudioFunction_Once sync.Once

func settingsSetEnableWebaudioFunction_Set() error {
	var err error
	settingsSetEnableWebaudioFunction_Once.Do(func() {
		err = settingsStruct_Set()
		if err != nil {
			return
		}
		settingsSetEnableWebaudioFunction, err = settingsStruct.InvokerNew("set_enable_webaudio")
	})
	return err
}

// SetEnableWebaudio is a representation of the C type webkit_settings_set_enable_webaudio.
func (recv *Settings) SetEnableWebaudio(enabled bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetBoolean(enabled)

	err := settingsSetEnableWebaudioFunction_Set()
	if err == nil {
		settingsSetEnableWebaudioFunction.Invoke(inArgs[:], nil)
	}

	return
}

var settingsSetEnableWebglFunction *gi.Function
var settingsSetEnableWebglFunction_Once sync.Once

func settingsSetEnableWebglFunction_Set() error {
	var err error
	settingsSetEnableWebglFunction_Once.Do(func() {
		err = settingsStruct_Set()
		if err != nil {
			return
		}
		settingsSetEnableWebglFunction, err = settingsStruct.InvokerNew("set_enable_webgl")
	})
	return err
}

// SetEnableWebgl is a representation of the C type webkit_settings_set_enable_webgl.
func (recv *Settings) SetEnableWebgl(enabled bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetBoolean(enabled)

	err := settingsSetEnableWebglFunction_Set()
	if err == nil {
		settingsSetEnableWebglFunction.Invoke(inArgs[:], nil)
	}

	return
}

var settingsSetEnableWriteConsoleMessagesToStdoutFunction *gi.Function
var settingsSetEnableWriteConsoleMessagesToStdoutFunction_Once sync.Once

func settingsSetEnableWriteConsoleMessagesToStdoutFunction_Set() error {
	var err error
	settingsSetEnableWriteConsoleMessagesToStdoutFunction_Once.Do(func() {
		err = settingsStruct_Set()
		if err != nil {
			return
		}
		settingsSetEnableWriteConsoleMessagesToStdoutFunction, err = settingsStruct.InvokerNew("set_enable_write_console_messages_to_stdout")
	})
	return err
}

// SetEnableWriteConsoleMessagesToStdout is a representation of the C type webkit_settings_set_enable_write_console_messages_to_stdout.
func (recv *Settings) SetEnableWriteConsoleMessagesToStdout(enabled bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetBoolean(enabled)

	err := settingsSetEnableWriteConsoleMessagesToStdoutFunction_Set()
	if err == nil {
		settingsSetEnableWriteConsoleMessagesToStdoutFunction.Invoke(inArgs[:], nil)
	}

	return
}

var settingsSetEnableXssAuditorFunction *gi.Function
var settingsSetEnableXssAuditorFunction_Once sync.Once

func settingsSetEnableXssAuditorFunction_Set() error {
	var err error
	settingsSetEnableXssAuditorFunction_Once.Do(func() {
		err = settingsStruct_Set()
		if err != nil {
			return
		}
		settingsSetEnableXssAuditorFunction, err = settingsStruct.InvokerNew("set_enable_xss_auditor")
	})
	return err
}

// SetEnableXssAuditor is a representation of the C type webkit_settings_set_enable_xss_auditor.
func (recv *Settings) SetEnableXssAuditor(enabled bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetBoolean(enabled)

	err := settingsSetEnableXssAuditorFunction_Set()
	if err == nil {
		settingsSetEnableXssAuditorFunction.Invoke(inArgs[:], nil)
	}

	return
}

var settingsSetFantasyFontFamilyFunction *gi.Function
var settingsSetFantasyFontFamilyFunction_Once sync.Once

func settingsSetFantasyFontFamilyFunction_Set() error {
	var err error
	settingsSetFantasyFontFamilyFunction_Once.Do(func() {
		err = settingsStruct_Set()
		if err != nil {
			return
		}
		settingsSetFantasyFontFamilyFunction, err = settingsStruct.InvokerNew("set_fantasy_font_family")
	})
	return err
}

// SetFantasyFontFamily is a representation of the C type webkit_settings_set_fantasy_font_family.
func (recv *Settings) SetFantasyFontFamily(fantasyFontFamily string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(fantasyFontFamily)

	err := settingsSetFantasyFontFamilyFunction_Set()
	if err == nil {
		settingsSetFantasyFontFamilyFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'webkit_settings_set_hardware_acceleration_policy' : parameter 'policy' of type 'HardwareAccelerationPolicy' not supported

var settingsSetJavascriptCanAccessClipboardFunction *gi.Function
var settingsSetJavascriptCanAccessClipboardFunction_Once sync.Once

func settingsSetJavascriptCanAccessClipboardFunction_Set() error {
	var err error
	settingsSetJavascriptCanAccessClipboardFunction_Once.Do(func() {
		err = settingsStruct_Set()
		if err != nil {
			return
		}
		settingsSetJavascriptCanAccessClipboardFunction, err = settingsStruct.InvokerNew("set_javascript_can_access_clipboard")
	})
	return err
}

// SetJavascriptCanAccessClipboard is a representation of the C type webkit_settings_set_javascript_can_access_clipboard.
func (recv *Settings) SetJavascriptCanAccessClipboard(enabled bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetBoolean(enabled)

	err := settingsSetJavascriptCanAccessClipboardFunction_Set()
	if err == nil {
		settingsSetJavascriptCanAccessClipboardFunction.Invoke(inArgs[:], nil)
	}

	return
}

var settingsSetJavascriptCanOpenWindowsAutomaticallyFunction *gi.Function
var settingsSetJavascriptCanOpenWindowsAutomaticallyFunction_Once sync.Once

func settingsSetJavascriptCanOpenWindowsAutomaticallyFunction_Set() error {
	var err error
	settingsSetJavascriptCanOpenWindowsAutomaticallyFunction_Once.Do(func() {
		err = settingsStruct_Set()
		if err != nil {
			return
		}
		settingsSetJavascriptCanOpenWindowsAutomaticallyFunction, err = settingsStruct.InvokerNew("set_javascript_can_open_windows_automatically")
	})
	return err
}

// SetJavascriptCanOpenWindowsAutomatically is a representation of the C type webkit_settings_set_javascript_can_open_windows_automatically.
func (recv *Settings) SetJavascriptCanOpenWindowsAutomatically(enabled bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetBoolean(enabled)

	err := settingsSetJavascriptCanOpenWindowsAutomaticallyFunction_Set()
	if err == nil {
		settingsSetJavascriptCanOpenWindowsAutomaticallyFunction.Invoke(inArgs[:], nil)
	}

	return
}

var settingsSetLoadIconsIgnoringImageLoadSettingFunction *gi.Function
var settingsSetLoadIconsIgnoringImageLoadSettingFunction_Once sync.Once

func settingsSetLoadIconsIgnoringImageLoadSettingFunction_Set() error {
	var err error
	settingsSetLoadIconsIgnoringImageLoadSettingFunction_Once.Do(func() {
		err = settingsStruct_Set()
		if err != nil {
			return
		}
		settingsSetLoadIconsIgnoringImageLoadSettingFunction, err = settingsStruct.InvokerNew("set_load_icons_ignoring_image_load_setting")
	})
	return err
}

// SetLoadIconsIgnoringImageLoadSetting is a representation of the C type webkit_settings_set_load_icons_ignoring_image_load_setting.
func (recv *Settings) SetLoadIconsIgnoringImageLoadSetting(enabled bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetBoolean(enabled)

	err := settingsSetLoadIconsIgnoringImageLoadSettingFunction_Set()
	if err == nil {
		settingsSetLoadIconsIgnoringImageLoadSettingFunction.Invoke(inArgs[:], nil)
	}

	return
}

var settingsSetMediaPlaybackAllowsInlineFunction *gi.Function
var settingsSetMediaPlaybackAllowsInlineFunction_Once sync.Once

func settingsSetMediaPlaybackAllowsInlineFunction_Set() error {
	var err error
	settingsSetMediaPlaybackAllowsInlineFunction_Once.Do(func() {
		err = settingsStruct_Set()
		if err != nil {
			return
		}
		settingsSetMediaPlaybackAllowsInlineFunction, err = settingsStruct.InvokerNew("set_media_playback_allows_inline")
	})
	return err
}

// SetMediaPlaybackAllowsInline is a representation of the C type webkit_settings_set_media_playback_allows_inline.
func (recv *Settings) SetMediaPlaybackAllowsInline(enabled bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetBoolean(enabled)

	err := settingsSetMediaPlaybackAllowsInlineFunction_Set()
	if err == nil {
		settingsSetMediaPlaybackAllowsInlineFunction.Invoke(inArgs[:], nil)
	}

	return
}

var settingsSetMediaPlaybackRequiresUserGestureFunction *gi.Function
var settingsSetMediaPlaybackRequiresUserGestureFunction_Once sync.Once

func settingsSetMediaPlaybackRequiresUserGestureFunction_Set() error {
	var err error
	settingsSetMediaPlaybackRequiresUserGestureFunction_Once.Do(func() {
		err = settingsStruct_Set()
		if err != nil {
			return
		}
		settingsSetMediaPlaybackRequiresUserGestureFunction, err = settingsStruct.InvokerNew("set_media_playback_requires_user_gesture")
	})
	return err
}

// SetMediaPlaybackRequiresUserGesture is a representation of the C type webkit_settings_set_media_playback_requires_user_gesture.
func (recv *Settings) SetMediaPlaybackRequiresUserGesture(enabled bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetBoolean(enabled)

	err := settingsSetMediaPlaybackRequiresUserGestureFunction_Set()
	if err == nil {
		settingsSetMediaPlaybackRequiresUserGestureFunction.Invoke(inArgs[:], nil)
	}

	return
}

var settingsSetMinimumFontSizeFunction *gi.Function
var settingsSetMinimumFontSizeFunction_Once sync.Once

func settingsSetMinimumFontSizeFunction_Set() error {
	var err error
	settingsSetMinimumFontSizeFunction_Once.Do(func() {
		err = settingsStruct_Set()
		if err != nil {
			return
		}
		settingsSetMinimumFontSizeFunction, err = settingsStruct.InvokerNew("set_minimum_font_size")
	})
	return err
}

// SetMinimumFontSize is a representation of the C type webkit_settings_set_minimum_font_size.
func (recv *Settings) SetMinimumFontSize(fontSize uint32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetUint32(fontSize)

	err := settingsSetMinimumFontSizeFunction_Set()
	if err == nil {
		settingsSetMinimumFontSizeFunction.Invoke(inArgs[:], nil)
	}

	return
}

var settingsSetMonospaceFontFamilyFunction *gi.Function
var settingsSetMonospaceFontFamilyFunction_Once sync.Once

func settingsSetMonospaceFontFamilyFunction_Set() error {
	var err error
	settingsSetMonospaceFontFamilyFunction_Once.Do(func() {
		err = settingsStruct_Set()
		if err != nil {
			return
		}
		settingsSetMonospaceFontFamilyFunction, err = settingsStruct.InvokerNew("set_monospace_font_family")
	})
	return err
}

// SetMonospaceFontFamily is a representation of the C type webkit_settings_set_monospace_font_family.
func (recv *Settings) SetMonospaceFontFamily(monospaceFontFamily string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(monospaceFontFamily)

	err := settingsSetMonospaceFontFamilyFunction_Set()
	if err == nil {
		settingsSetMonospaceFontFamilyFunction.Invoke(inArgs[:], nil)
	}

	return
}

var settingsSetPictographFontFamilyFunction *gi.Function
var settingsSetPictographFontFamilyFunction_Once sync.Once

func settingsSetPictographFontFamilyFunction_Set() error {
	var err error
	settingsSetPictographFontFamilyFunction_Once.Do(func() {
		err = settingsStruct_Set()
		if err != nil {
			return
		}
		settingsSetPictographFontFamilyFunction, err = settingsStruct.InvokerNew("set_pictograph_font_family")
	})
	return err
}

// SetPictographFontFamily is a representation of the C type webkit_settings_set_pictograph_font_family.
func (recv *Settings) SetPictographFontFamily(pictographFontFamily string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(pictographFontFamily)

	err := settingsSetPictographFontFamilyFunction_Set()
	if err == nil {
		settingsSetPictographFontFamilyFunction.Invoke(inArgs[:], nil)
	}

	return
}

var settingsSetPrintBackgroundsFunction *gi.Function
var settingsSetPrintBackgroundsFunction_Once sync.Once

func settingsSetPrintBackgroundsFunction_Set() error {
	var err error
	settingsSetPrintBackgroundsFunction_Once.Do(func() {
		err = settingsStruct_Set()
		if err != nil {
			return
		}
		settingsSetPrintBackgroundsFunction, err = settingsStruct.InvokerNew("set_print_backgrounds")
	})
	return err
}

// SetPrintBackgrounds is a representation of the C type webkit_settings_set_print_backgrounds.
func (recv *Settings) SetPrintBackgrounds(printBackgrounds bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetBoolean(printBackgrounds)

	err := settingsSetPrintBackgroundsFunction_Set()
	if err == nil {
		settingsSetPrintBackgroundsFunction.Invoke(inArgs[:], nil)
	}

	return
}

var settingsSetSansSerifFontFamilyFunction *gi.Function
var settingsSetSansSerifFontFamilyFunction_Once sync.Once

func settingsSetSansSerifFontFamilyFunction_Set() error {
	var err error
	settingsSetSansSerifFontFamilyFunction_Once.Do(func() {
		err = settingsStruct_Set()
		if err != nil {
			return
		}
		settingsSetSansSerifFontFamilyFunction, err = settingsStruct.InvokerNew("set_sans_serif_font_family")
	})
	return err
}

// SetSansSerifFontFamily is a representation of the C type webkit_settings_set_sans_serif_font_family.
func (recv *Settings) SetSansSerifFontFamily(sansSerifFontFamily string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(sansSerifFontFamily)

	err := settingsSetSansSerifFontFamilyFunction_Set()
	if err == nil {
		settingsSetSansSerifFontFamilyFunction.Invoke(inArgs[:], nil)
	}

	return
}

var settingsSetSerifFontFamilyFunction *gi.Function
var settingsSetSerifFontFamilyFunction_Once sync.Once

func settingsSetSerifFontFamilyFunction_Set() error {
	var err error
	settingsSetSerifFontFamilyFunction_Once.Do(func() {
		err = settingsStruct_Set()
		if err != nil {
			return
		}
		settingsSetSerifFontFamilyFunction, err = settingsStruct.InvokerNew("set_serif_font_family")
	})
	return err
}

// SetSerifFontFamily is a representation of the C type webkit_settings_set_serif_font_family.
func (recv *Settings) SetSerifFontFamily(serifFontFamily string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(serifFontFamily)

	err := settingsSetSerifFontFamilyFunction_Set()
	if err == nil {
		settingsSetSerifFontFamilyFunction.Invoke(inArgs[:], nil)
	}

	return
}

var settingsSetUserAgentFunction *gi.Function
var settingsSetUserAgentFunction_Once sync.Once

func settingsSetUserAgentFunction_Set() error {
	var err error
	settingsSetUserAgentFunction_Once.Do(func() {
		err = settingsStruct_Set()
		if err != nil {
			return
		}
		settingsSetUserAgentFunction, err = settingsStruct.InvokerNew("set_user_agent")
	})
	return err
}

// SetUserAgent is a representation of the C type webkit_settings_set_user_agent.
func (recv *Settings) SetUserAgent(userAgent string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(userAgent)

	err := settingsSetUserAgentFunction_Set()
	if err == nil {
		settingsSetUserAgentFunction.Invoke(inArgs[:], nil)
	}

	return
}

var settingsSetUserAgentWithApplicationDetailsFunction *gi.Function
var settingsSetUserAgentWithApplicationDetailsFunction_Once sync.Once

func settingsSetUserAgentWithApplicationDetailsFunction_Set() error {
	var err error
	settingsSetUserAgentWithApplicationDetailsFunction_Once.Do(func() {
		err = settingsStruct_Set()
		if err != nil {
			return
		}
		settingsSetUserAgentWithApplicationDetailsFunction, err = settingsStruct.InvokerNew("set_user_agent_with_application_details")
	})
	return err
}

// SetUserAgentWithApplicationDetails is a representation of the C type webkit_settings_set_user_agent_with_application_details.
func (recv *Settings) SetUserAgentWithApplicationDetails(applicationName string, applicationVersion string) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(applicationName)
	inArgs[2].SetString(applicationVersion)

	err := settingsSetUserAgentWithApplicationDetailsFunction_Set()
	if err == nil {
		settingsSetUserAgentWithApplicationDetailsFunction.Invoke(inArgs[:], nil)
	}

	return
}

var settingsSetZoomTextOnlyFunction *gi.Function
var settingsSetZoomTextOnlyFunction_Once sync.Once

func settingsSetZoomTextOnlyFunction_Set() error {
	var err error
	settingsSetZoomTextOnlyFunction_Once.Do(func() {
		err = settingsStruct_Set()
		if err != nil {
			return
		}
		settingsSetZoomTextOnlyFunction, err = settingsStruct.InvokerNew("set_zoom_text_only")
	})
	return err
}

// SetZoomTextOnly is a representation of the C type webkit_settings_set_zoom_text_only.
func (recv *Settings) SetZoomTextOnly(zoomTextOnly bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetBoolean(zoomTextOnly)

	err := settingsSetZoomTextOnlyFunction_Set()
	if err == nil {
		settingsSetZoomTextOnlyFunction.Invoke(inArgs[:], nil)
	}

	return
}

var uRIRequestStruct *gi.Struct
var uRIRequestStruct_Once sync.Once

func uRIRequestStruct_Set() error {
	var err error
	uRIRequestStruct_Once.Do(func() {
		uRIRequestStruct, err = gi.StructNew("WebKit2", "URIRequest")
	})
	return err
}

type URIRequest struct {
	native uintptr
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'webkit_uri_request_new' : return type 'URIRequest' not supported

// UNSUPPORTED : C value 'webkit_uri_request_get_http_headers' : return type 'Soup.MessageHeaders' not supported

var uRIRequestGetHttpMethodFunction *gi.Function
var uRIRequestGetHttpMethodFunction_Once sync.Once

func uRIRequestGetHttpMethodFunction_Set() error {
	var err error
	uRIRequestGetHttpMethodFunction_Once.Do(func() {
		err = uRIRequestStruct_Set()
		if err != nil {
			return
		}
		uRIRequestGetHttpMethodFunction, err = uRIRequestStruct.InvokerNew("get_http_method")
	})
	return err
}

// GetHttpMethod is a representation of the C type webkit_uri_request_get_http_method.
func (recv *URIRequest) GetHttpMethod() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := uRIRequestGetHttpMethodFunction_Set()
	if err == nil {
		ret = uRIRequestGetHttpMethodFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var uRIRequestGetUriFunction *gi.Function
var uRIRequestGetUriFunction_Once sync.Once

func uRIRequestGetUriFunction_Set() error {
	var err error
	uRIRequestGetUriFunction_Once.Do(func() {
		err = uRIRequestStruct_Set()
		if err != nil {
			return
		}
		uRIRequestGetUriFunction, err = uRIRequestStruct.InvokerNew("get_uri")
	})
	return err
}

// GetUri is a representation of the C type webkit_uri_request_get_uri.
func (recv *URIRequest) GetUri() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := uRIRequestGetUriFunction_Set()
	if err == nil {
		ret = uRIRequestGetUriFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var uRIRequestSetUriFunction *gi.Function
var uRIRequestSetUriFunction_Once sync.Once

func uRIRequestSetUriFunction_Set() error {
	var err error
	uRIRequestSetUriFunction_Once.Do(func() {
		err = uRIRequestStruct_Set()
		if err != nil {
			return
		}
		uRIRequestSetUriFunction, err = uRIRequestStruct.InvokerNew("set_uri")
	})
	return err
}

// SetUri is a representation of the C type webkit_uri_request_set_uri.
func (recv *URIRequest) SetUri(uri string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(uri)

	err := uRIRequestSetUriFunction_Set()
	if err == nil {
		uRIRequestSetUriFunction.Invoke(inArgs[:], nil)
	}

	return
}

var uRIResponseStruct *gi.Struct
var uRIResponseStruct_Once sync.Once

func uRIResponseStruct_Set() error {
	var err error
	uRIResponseStruct_Once.Do(func() {
		uRIResponseStruct, err = gi.StructNew("WebKit2", "URIResponse")
	})
	return err
}

type URIResponse struct {
	native uintptr
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'GObject.Object'

var uRIResponseGetContentLengthFunction *gi.Function
var uRIResponseGetContentLengthFunction_Once sync.Once

func uRIResponseGetContentLengthFunction_Set() error {
	var err error
	uRIResponseGetContentLengthFunction_Once.Do(func() {
		err = uRIResponseStruct_Set()
		if err != nil {
			return
		}
		uRIResponseGetContentLengthFunction, err = uRIResponseStruct.InvokerNew("get_content_length")
	})
	return err
}

// GetContentLength is a representation of the C type webkit_uri_response_get_content_length.
func (recv *URIResponse) GetContentLength() uint64 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := uRIResponseGetContentLengthFunction_Set()
	if err == nil {
		ret = uRIResponseGetContentLengthFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint64()

	return retGo
}

// UNSUPPORTED : C value 'webkit_uri_response_get_http_headers' : return type 'Soup.MessageHeaders' not supported

var uRIResponseGetMimeTypeFunction *gi.Function
var uRIResponseGetMimeTypeFunction_Once sync.Once

func uRIResponseGetMimeTypeFunction_Set() error {
	var err error
	uRIResponseGetMimeTypeFunction_Once.Do(func() {
		err = uRIResponseStruct_Set()
		if err != nil {
			return
		}
		uRIResponseGetMimeTypeFunction, err = uRIResponseStruct.InvokerNew("get_mime_type")
	})
	return err
}

// GetMimeType is a representation of the C type webkit_uri_response_get_mime_type.
func (recv *URIResponse) GetMimeType() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := uRIResponseGetMimeTypeFunction_Set()
	if err == nil {
		ret = uRIResponseGetMimeTypeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var uRIResponseGetStatusCodeFunction *gi.Function
var uRIResponseGetStatusCodeFunction_Once sync.Once

func uRIResponseGetStatusCodeFunction_Set() error {
	var err error
	uRIResponseGetStatusCodeFunction_Once.Do(func() {
		err = uRIResponseStruct_Set()
		if err != nil {
			return
		}
		uRIResponseGetStatusCodeFunction, err = uRIResponseStruct.InvokerNew("get_status_code")
	})
	return err
}

// GetStatusCode is a representation of the C type webkit_uri_response_get_status_code.
func (recv *URIResponse) GetStatusCode() uint32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := uRIResponseGetStatusCodeFunction_Set()
	if err == nil {
		ret = uRIResponseGetStatusCodeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo
}

var uRIResponseGetSuggestedFilenameFunction *gi.Function
var uRIResponseGetSuggestedFilenameFunction_Once sync.Once

func uRIResponseGetSuggestedFilenameFunction_Set() error {
	var err error
	uRIResponseGetSuggestedFilenameFunction_Once.Do(func() {
		err = uRIResponseStruct_Set()
		if err != nil {
			return
		}
		uRIResponseGetSuggestedFilenameFunction, err = uRIResponseStruct.InvokerNew("get_suggested_filename")
	})
	return err
}

// GetSuggestedFilename is a representation of the C type webkit_uri_response_get_suggested_filename.
func (recv *URIResponse) GetSuggestedFilename() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := uRIResponseGetSuggestedFilenameFunction_Set()
	if err == nil {
		ret = uRIResponseGetSuggestedFilenameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var uRIResponseGetUriFunction *gi.Function
var uRIResponseGetUriFunction_Once sync.Once

func uRIResponseGetUriFunction_Set() error {
	var err error
	uRIResponseGetUriFunction_Once.Do(func() {
		err = uRIResponseStruct_Set()
		if err != nil {
			return
		}
		uRIResponseGetUriFunction, err = uRIResponseStruct.InvokerNew("get_uri")
	})
	return err
}

// GetUri is a representation of the C type webkit_uri_response_get_uri.
func (recv *URIResponse) GetUri() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := uRIResponseGetUriFunction_Set()
	if err == nil {
		ret = uRIResponseGetUriFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

// URIResponseStruct creates an uninitialised URIResponse.
func URIResponseStruct() *URIResponse {
	err := uRIResponseStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &URIResponse{native: uRIResponseStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeURIResponse)
	return structGo
}
func finalizeURIResponse(obj *URIResponse) {
	uRIResponseStruct.Free(obj.native)
}

var uRISchemeRequestStruct *gi.Struct
var uRISchemeRequestStruct_Once sync.Once

func uRISchemeRequestStruct_Set() error {
	var err error
	uRISchemeRequestStruct_Once.Do(func() {
		uRISchemeRequestStruct, err = gi.StructNew("WebKit2", "URISchemeRequest")
	})
	return err
}

type URISchemeRequest struct {
	native uintptr
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'GObject.Object'

// FieldPriv returns the C field 'priv'.
func (recv *URISchemeRequest) FieldPriv() *URISchemeRequestPrivate {
	argValue := gi.FieldGet(uRISchemeRequestStruct, recv.native, "priv")
	value := &URISchemeRequestPrivate{native: argValue.Pointer()}
	return value
}

// SetFieldPriv sets the value of the C field 'priv'.
func (recv *URISchemeRequest) SetFieldPriv(value *URISchemeRequestPrivate) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(uRISchemeRequestStruct, recv.native, "priv", argValue)
}

// UNSUPPORTED : C value 'webkit_uri_scheme_request_finish' : parameter 'stream' of type 'Gio.InputStream' not supported

// UNSUPPORTED : C value 'webkit_uri_scheme_request_finish_error' : parameter 'error' of type 'GLib.Error' not supported

var uRISchemeRequestGetPathFunction *gi.Function
var uRISchemeRequestGetPathFunction_Once sync.Once

func uRISchemeRequestGetPathFunction_Set() error {
	var err error
	uRISchemeRequestGetPathFunction_Once.Do(func() {
		err = uRISchemeRequestStruct_Set()
		if err != nil {
			return
		}
		uRISchemeRequestGetPathFunction, err = uRISchemeRequestStruct.InvokerNew("get_path")
	})
	return err
}

// GetPath is a representation of the C type webkit_uri_scheme_request_get_path.
func (recv *URISchemeRequest) GetPath() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := uRISchemeRequestGetPathFunction_Set()
	if err == nil {
		ret = uRISchemeRequestGetPathFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var uRISchemeRequestGetSchemeFunction *gi.Function
var uRISchemeRequestGetSchemeFunction_Once sync.Once

func uRISchemeRequestGetSchemeFunction_Set() error {
	var err error
	uRISchemeRequestGetSchemeFunction_Once.Do(func() {
		err = uRISchemeRequestStruct_Set()
		if err != nil {
			return
		}
		uRISchemeRequestGetSchemeFunction, err = uRISchemeRequestStruct.InvokerNew("get_scheme")
	})
	return err
}

// GetScheme is a representation of the C type webkit_uri_scheme_request_get_scheme.
func (recv *URISchemeRequest) GetScheme() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := uRISchemeRequestGetSchemeFunction_Set()
	if err == nil {
		ret = uRISchemeRequestGetSchemeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var uRISchemeRequestGetUriFunction *gi.Function
var uRISchemeRequestGetUriFunction_Once sync.Once

func uRISchemeRequestGetUriFunction_Set() error {
	var err error
	uRISchemeRequestGetUriFunction_Once.Do(func() {
		err = uRISchemeRequestStruct_Set()
		if err != nil {
			return
		}
		uRISchemeRequestGetUriFunction, err = uRISchemeRequestStruct.InvokerNew("get_uri")
	})
	return err
}

// GetUri is a representation of the C type webkit_uri_scheme_request_get_uri.
func (recv *URISchemeRequest) GetUri() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := uRISchemeRequestGetUriFunction_Set()
	if err == nil {
		ret = uRISchemeRequestGetUriFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'webkit_uri_scheme_request_get_web_view' : return type 'WebView' not supported

// URISchemeRequestStruct creates an uninitialised URISchemeRequest.
func URISchemeRequestStruct() *URISchemeRequest {
	err := uRISchemeRequestStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &URISchemeRequest{native: uRISchemeRequestStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeURISchemeRequest)
	return structGo
}
func finalizeURISchemeRequest(obj *URISchemeRequest) {
	uRISchemeRequestStruct.Free(obj.native)
}

var userContentFilterStoreStruct *gi.Struct
var userContentFilterStoreStruct_Once sync.Once

func userContentFilterStoreStruct_Set() error {
	var err error
	userContentFilterStoreStruct_Once.Do(func() {
		userContentFilterStoreStruct, err = gi.StructNew("WebKit2", "UserContentFilterStore")
	})
	return err
}

type UserContentFilterStore struct {
	native uintptr
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'webkit_user_content_filter_store_new' : return type 'UserContentFilterStore' not supported

// UNSUPPORTED : C value 'webkit_user_content_filter_store_fetch_identifiers' : parameter 'cancellable' of type 'Gio.Cancellable' not supported

// UNSUPPORTED : C value 'webkit_user_content_filter_store_fetch_identifiers_finish' : parameter 'result' of type 'Gio.AsyncResult' not supported

var userContentFilterStoreGetPathFunction *gi.Function
var userContentFilterStoreGetPathFunction_Once sync.Once

func userContentFilterStoreGetPathFunction_Set() error {
	var err error
	userContentFilterStoreGetPathFunction_Once.Do(func() {
		err = userContentFilterStoreStruct_Set()
		if err != nil {
			return
		}
		userContentFilterStoreGetPathFunction, err = userContentFilterStoreStruct.InvokerNew("get_path")
	})
	return err
}

// GetPath is a representation of the C type webkit_user_content_filter_store_get_path.
func (recv *UserContentFilterStore) GetPath() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := userContentFilterStoreGetPathFunction_Set()
	if err == nil {
		ret = userContentFilterStoreGetPathFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'webkit_user_content_filter_store_load' : parameter 'cancellable' of type 'Gio.Cancellable' not supported

// UNSUPPORTED : C value 'webkit_user_content_filter_store_load_finish' : parameter 'result' of type 'Gio.AsyncResult' not supported

// UNSUPPORTED : C value 'webkit_user_content_filter_store_remove' : parameter 'cancellable' of type 'Gio.Cancellable' not supported

// UNSUPPORTED : C value 'webkit_user_content_filter_store_remove_finish' : parameter 'result' of type 'Gio.AsyncResult' not supported

// UNSUPPORTED : C value 'webkit_user_content_filter_store_save' : parameter 'source' of type 'GLib.Bytes' not supported

// UNSUPPORTED : C value 'webkit_user_content_filter_store_save_finish' : parameter 'result' of type 'Gio.AsyncResult' not supported

// UNSUPPORTED : C value 'webkit_user_content_filter_store_save_from_file' : parameter 'file' of type 'Gio.File' not supported

// UNSUPPORTED : C value 'webkit_user_content_filter_store_save_from_file_finish' : parameter 'result' of type 'Gio.AsyncResult' not supported

var userContentManagerStruct *gi.Struct
var userContentManagerStruct_Once sync.Once

func userContentManagerStruct_Set() error {
	var err error
	userContentManagerStruct_Once.Do(func() {
		userContentManagerStruct, err = gi.StructNew("WebKit2", "UserContentManager")
	})
	return err
}

type UserContentManager struct {
	native uintptr
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'webkit_user_content_manager_new' : return type 'UserContentManager' not supported

var userContentManagerAddFilterFunction *gi.Function
var userContentManagerAddFilterFunction_Once sync.Once

func userContentManagerAddFilterFunction_Set() error {
	var err error
	userContentManagerAddFilterFunction_Once.Do(func() {
		err = userContentManagerStruct_Set()
		if err != nil {
			return
		}
		userContentManagerAddFilterFunction, err = userContentManagerStruct.InvokerNew("add_filter")
	})
	return err
}

// AddFilter is a representation of the C type webkit_user_content_manager_add_filter.
func (recv *UserContentManager) AddFilter(filter *UserContentFilter) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(filter.native)

	err := userContentManagerAddFilterFunction_Set()
	if err == nil {
		userContentManagerAddFilterFunction.Invoke(inArgs[:], nil)
	}

	return
}

var userContentManagerAddScriptFunction *gi.Function
var userContentManagerAddScriptFunction_Once sync.Once

func userContentManagerAddScriptFunction_Set() error {
	var err error
	userContentManagerAddScriptFunction_Once.Do(func() {
		err = userContentManagerStruct_Set()
		if err != nil {
			return
		}
		userContentManagerAddScriptFunction, err = userContentManagerStruct.InvokerNew("add_script")
	})
	return err
}

// AddScript is a representation of the C type webkit_user_content_manager_add_script.
func (recv *UserContentManager) AddScript(script *UserScript) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(script.native)

	err := userContentManagerAddScriptFunction_Set()
	if err == nil {
		userContentManagerAddScriptFunction.Invoke(inArgs[:], nil)
	}

	return
}

var userContentManagerAddStyleSheetFunction *gi.Function
var userContentManagerAddStyleSheetFunction_Once sync.Once

func userContentManagerAddStyleSheetFunction_Set() error {
	var err error
	userContentManagerAddStyleSheetFunction_Once.Do(func() {
		err = userContentManagerStruct_Set()
		if err != nil {
			return
		}
		userContentManagerAddStyleSheetFunction, err = userContentManagerStruct.InvokerNew("add_style_sheet")
	})
	return err
}

// AddStyleSheet is a representation of the C type webkit_user_content_manager_add_style_sheet.
func (recv *UserContentManager) AddStyleSheet(stylesheet *UserStyleSheet) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(stylesheet.native)

	err := userContentManagerAddStyleSheetFunction_Set()
	if err == nil {
		userContentManagerAddStyleSheetFunction.Invoke(inArgs[:], nil)
	}

	return
}

var userContentManagerRegisterScriptMessageHandlerFunction *gi.Function
var userContentManagerRegisterScriptMessageHandlerFunction_Once sync.Once

func userContentManagerRegisterScriptMessageHandlerFunction_Set() error {
	var err error
	userContentManagerRegisterScriptMessageHandlerFunction_Once.Do(func() {
		err = userContentManagerStruct_Set()
		if err != nil {
			return
		}
		userContentManagerRegisterScriptMessageHandlerFunction, err = userContentManagerStruct.InvokerNew("register_script_message_handler")
	})
	return err
}

// RegisterScriptMessageHandler is a representation of the C type webkit_user_content_manager_register_script_message_handler.
func (recv *UserContentManager) RegisterScriptMessageHandler(name string) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(name)

	var ret gi.Argument

	err := userContentManagerRegisterScriptMessageHandlerFunction_Set()
	if err == nil {
		ret = userContentManagerRegisterScriptMessageHandlerFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var userContentManagerRegisterScriptMessageHandlerInWorldFunction *gi.Function
var userContentManagerRegisterScriptMessageHandlerInWorldFunction_Once sync.Once

func userContentManagerRegisterScriptMessageHandlerInWorldFunction_Set() error {
	var err error
	userContentManagerRegisterScriptMessageHandlerInWorldFunction_Once.Do(func() {
		err = userContentManagerStruct_Set()
		if err != nil {
			return
		}
		userContentManagerRegisterScriptMessageHandlerInWorldFunction, err = userContentManagerStruct.InvokerNew("register_script_message_handler_in_world")
	})
	return err
}

// RegisterScriptMessageHandlerInWorld is a representation of the C type webkit_user_content_manager_register_script_message_handler_in_world.
func (recv *UserContentManager) RegisterScriptMessageHandlerInWorld(name string, worldName string) bool {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(name)
	inArgs[2].SetString(worldName)

	var ret gi.Argument

	err := userContentManagerRegisterScriptMessageHandlerInWorldFunction_Set()
	if err == nil {
		ret = userContentManagerRegisterScriptMessageHandlerInWorldFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var userContentManagerRemoveAllFiltersFunction *gi.Function
var userContentManagerRemoveAllFiltersFunction_Once sync.Once

func userContentManagerRemoveAllFiltersFunction_Set() error {
	var err error
	userContentManagerRemoveAllFiltersFunction_Once.Do(func() {
		err = userContentManagerStruct_Set()
		if err != nil {
			return
		}
		userContentManagerRemoveAllFiltersFunction, err = userContentManagerStruct.InvokerNew("remove_all_filters")
	})
	return err
}

// RemoveAllFilters is a representation of the C type webkit_user_content_manager_remove_all_filters.
func (recv *UserContentManager) RemoveAllFilters() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := userContentManagerRemoveAllFiltersFunction_Set()
	if err == nil {
		userContentManagerRemoveAllFiltersFunction.Invoke(inArgs[:], nil)
	}

	return
}

var userContentManagerRemoveAllScriptsFunction *gi.Function
var userContentManagerRemoveAllScriptsFunction_Once sync.Once

func userContentManagerRemoveAllScriptsFunction_Set() error {
	var err error
	userContentManagerRemoveAllScriptsFunction_Once.Do(func() {
		err = userContentManagerStruct_Set()
		if err != nil {
			return
		}
		userContentManagerRemoveAllScriptsFunction, err = userContentManagerStruct.InvokerNew("remove_all_scripts")
	})
	return err
}

// RemoveAllScripts is a representation of the C type webkit_user_content_manager_remove_all_scripts.
func (recv *UserContentManager) RemoveAllScripts() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := userContentManagerRemoveAllScriptsFunction_Set()
	if err == nil {
		userContentManagerRemoveAllScriptsFunction.Invoke(inArgs[:], nil)
	}

	return
}

var userContentManagerRemoveAllStyleSheetsFunction *gi.Function
var userContentManagerRemoveAllStyleSheetsFunction_Once sync.Once

func userContentManagerRemoveAllStyleSheetsFunction_Set() error {
	var err error
	userContentManagerRemoveAllStyleSheetsFunction_Once.Do(func() {
		err = userContentManagerStruct_Set()
		if err != nil {
			return
		}
		userContentManagerRemoveAllStyleSheetsFunction, err = userContentManagerStruct.InvokerNew("remove_all_style_sheets")
	})
	return err
}

// RemoveAllStyleSheets is a representation of the C type webkit_user_content_manager_remove_all_style_sheets.
func (recv *UserContentManager) RemoveAllStyleSheets() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := userContentManagerRemoveAllStyleSheetsFunction_Set()
	if err == nil {
		userContentManagerRemoveAllStyleSheetsFunction.Invoke(inArgs[:], nil)
	}

	return
}

var userContentManagerRemoveFilterFunction *gi.Function
var userContentManagerRemoveFilterFunction_Once sync.Once

func userContentManagerRemoveFilterFunction_Set() error {
	var err error
	userContentManagerRemoveFilterFunction_Once.Do(func() {
		err = userContentManagerStruct_Set()
		if err != nil {
			return
		}
		userContentManagerRemoveFilterFunction, err = userContentManagerStruct.InvokerNew("remove_filter")
	})
	return err
}

// RemoveFilter is a representation of the C type webkit_user_content_manager_remove_filter.
func (recv *UserContentManager) RemoveFilter(filter *UserContentFilter) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(filter.native)

	err := userContentManagerRemoveFilterFunction_Set()
	if err == nil {
		userContentManagerRemoveFilterFunction.Invoke(inArgs[:], nil)
	}

	return
}

var userContentManagerRemoveFilterByIdFunction *gi.Function
var userContentManagerRemoveFilterByIdFunction_Once sync.Once

func userContentManagerRemoveFilterByIdFunction_Set() error {
	var err error
	userContentManagerRemoveFilterByIdFunction_Once.Do(func() {
		err = userContentManagerStruct_Set()
		if err != nil {
			return
		}
		userContentManagerRemoveFilterByIdFunction, err = userContentManagerStruct.InvokerNew("remove_filter_by_id")
	})
	return err
}

// RemoveFilterById is a representation of the C type webkit_user_content_manager_remove_filter_by_id.
func (recv *UserContentManager) RemoveFilterById(filterId string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(filterId)

	err := userContentManagerRemoveFilterByIdFunction_Set()
	if err == nil {
		userContentManagerRemoveFilterByIdFunction.Invoke(inArgs[:], nil)
	}

	return
}

var userContentManagerUnregisterScriptMessageHandlerFunction *gi.Function
var userContentManagerUnregisterScriptMessageHandlerFunction_Once sync.Once

func userContentManagerUnregisterScriptMessageHandlerFunction_Set() error {
	var err error
	userContentManagerUnregisterScriptMessageHandlerFunction_Once.Do(func() {
		err = userContentManagerStruct_Set()
		if err != nil {
			return
		}
		userContentManagerUnregisterScriptMessageHandlerFunction, err = userContentManagerStruct.InvokerNew("unregister_script_message_handler")
	})
	return err
}

// UnregisterScriptMessageHandler is a representation of the C type webkit_user_content_manager_unregister_script_message_handler.
func (recv *UserContentManager) UnregisterScriptMessageHandler(name string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(name)

	err := userContentManagerUnregisterScriptMessageHandlerFunction_Set()
	if err == nil {
		userContentManagerUnregisterScriptMessageHandlerFunction.Invoke(inArgs[:], nil)
	}

	return
}

var userContentManagerUnregisterScriptMessageHandlerInWorldFunction *gi.Function
var userContentManagerUnregisterScriptMessageHandlerInWorldFunction_Once sync.Once

func userContentManagerUnregisterScriptMessageHandlerInWorldFunction_Set() error {
	var err error
	userContentManagerUnregisterScriptMessageHandlerInWorldFunction_Once.Do(func() {
		err = userContentManagerStruct_Set()
		if err != nil {
			return
		}
		userContentManagerUnregisterScriptMessageHandlerInWorldFunction, err = userContentManagerStruct.InvokerNew("unregister_script_message_handler_in_world")
	})
	return err
}

// UnregisterScriptMessageHandlerInWorld is a representation of the C type webkit_user_content_manager_unregister_script_message_handler_in_world.
func (recv *UserContentManager) UnregisterScriptMessageHandlerInWorld(name string, worldName string) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(name)
	inArgs[2].SetString(worldName)

	err := userContentManagerUnregisterScriptMessageHandlerInWorldFunction_Set()
	if err == nil {
		userContentManagerUnregisterScriptMessageHandlerInWorldFunction.Invoke(inArgs[:], nil)
	}

	return
}

var userMediaPermissionRequestStruct *gi.Struct
var userMediaPermissionRequestStruct_Once sync.Once

func userMediaPermissionRequestStruct_Set() error {
	var err error
	userMediaPermissionRequestStruct_Once.Do(func() {
		userMediaPermissionRequestStruct, err = gi.StructNew("WebKit2", "UserMediaPermissionRequest")
	})
	return err
}

type UserMediaPermissionRequest struct {
	native uintptr
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'GObject.Object'

// UserMediaPermissionRequestStruct creates an uninitialised UserMediaPermissionRequest.
func UserMediaPermissionRequestStruct() *UserMediaPermissionRequest {
	err := userMediaPermissionRequestStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &UserMediaPermissionRequest{native: userMediaPermissionRequestStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeUserMediaPermissionRequest)
	return structGo
}
func finalizeUserMediaPermissionRequest(obj *UserMediaPermissionRequest) {
	userMediaPermissionRequestStruct.Free(obj.native)
}

var webContextStruct *gi.Struct
var webContextStruct_Once sync.Once

func webContextStruct_Set() error {
	var err error
	webContextStruct_Once.Do(func() {
		webContextStruct, err = gi.StructNew("WebKit2", "WebContext")
	})
	return err
}

type WebContext struct {
	native uintptr
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'webkit_web_context_new' : return type 'WebContext' not supported

// UNSUPPORTED : C value 'webkit_web_context_new_ephemeral' : return type 'WebContext' not supported

// UNSUPPORTED : C value 'webkit_web_context_new_with_website_data_manager' : parameter 'manager' of type 'WebsiteDataManager' not supported

var webContextAddPathToSandboxFunction *gi.Function
var webContextAddPathToSandboxFunction_Once sync.Once

func webContextAddPathToSandboxFunction_Set() error {
	var err error
	webContextAddPathToSandboxFunction_Once.Do(func() {
		err = webContextStruct_Set()
		if err != nil {
			return
		}
		webContextAddPathToSandboxFunction, err = webContextStruct.InvokerNew("add_path_to_sandbox")
	})
	return err
}

// AddPathToSandbox is a representation of the C type webkit_web_context_add_path_to_sandbox.
func (recv *WebContext) AddPathToSandbox(path string, readOnly bool) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(path)
	inArgs[2].SetBoolean(readOnly)

	err := webContextAddPathToSandboxFunction_Set()
	if err == nil {
		webContextAddPathToSandboxFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'webkit_web_context_allow_tls_certificate_for_host' : parameter 'certificate' of type 'Gio.TlsCertificate' not supported

var webContextClearCacheFunction *gi.Function
var webContextClearCacheFunction_Once sync.Once

func webContextClearCacheFunction_Set() error {
	var err error
	webContextClearCacheFunction_Once.Do(func() {
		err = webContextStruct_Set()
		if err != nil {
			return
		}
		webContextClearCacheFunction, err = webContextStruct.InvokerNew("clear_cache")
	})
	return err
}

// ClearCache is a representation of the C type webkit_web_context_clear_cache.
func (recv *WebContext) ClearCache() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := webContextClearCacheFunction_Set()
	if err == nil {
		webContextClearCacheFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'webkit_web_context_download_uri' : return type 'Download' not supported

// UNSUPPORTED : C value 'webkit_web_context_get_cache_model' : return type 'CacheModel' not supported

// UNSUPPORTED : C value 'webkit_web_context_get_cookie_manager' : return type 'CookieManager' not supported

// UNSUPPORTED : C value 'webkit_web_context_get_favicon_database' : return type 'FaviconDatabase' not supported

var webContextGetFaviconDatabaseDirectoryFunction *gi.Function
var webContextGetFaviconDatabaseDirectoryFunction_Once sync.Once

func webContextGetFaviconDatabaseDirectoryFunction_Set() error {
	var err error
	webContextGetFaviconDatabaseDirectoryFunction_Once.Do(func() {
		err = webContextStruct_Set()
		if err != nil {
			return
		}
		webContextGetFaviconDatabaseDirectoryFunction, err = webContextStruct.InvokerNew("get_favicon_database_directory")
	})
	return err
}

// GetFaviconDatabaseDirectory is a representation of the C type webkit_web_context_get_favicon_database_directory.
func (recv *WebContext) GetFaviconDatabaseDirectory() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := webContextGetFaviconDatabaseDirectoryFunction_Set()
	if err == nil {
		ret = webContextGetFaviconDatabaseDirectoryFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'webkit_web_context_get_geolocation_manager' : return type 'GeolocationManager' not supported

// UNSUPPORTED : C value 'webkit_web_context_get_plugins' : parameter 'cancellable' of type 'Gio.Cancellable' not supported

// UNSUPPORTED : C value 'webkit_web_context_get_plugins_finish' : parameter 'result' of type 'Gio.AsyncResult' not supported

// UNSUPPORTED : C value 'webkit_web_context_get_process_model' : return type 'ProcessModel' not supported

var webContextGetSandboxEnabledFunction *gi.Function
var webContextGetSandboxEnabledFunction_Once sync.Once

func webContextGetSandboxEnabledFunction_Set() error {
	var err error
	webContextGetSandboxEnabledFunction_Once.Do(func() {
		err = webContextStruct_Set()
		if err != nil {
			return
		}
		webContextGetSandboxEnabledFunction, err = webContextStruct.InvokerNew("get_sandbox_enabled")
	})
	return err
}

// GetSandboxEnabled is a representation of the C type webkit_web_context_get_sandbox_enabled.
func (recv *WebContext) GetSandboxEnabled() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := webContextGetSandboxEnabledFunction_Set()
	if err == nil {
		ret = webContextGetSandboxEnabledFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'webkit_web_context_get_security_manager' : return type 'SecurityManager' not supported

var webContextGetSpellCheckingEnabledFunction *gi.Function
var webContextGetSpellCheckingEnabledFunction_Once sync.Once

func webContextGetSpellCheckingEnabledFunction_Set() error {
	var err error
	webContextGetSpellCheckingEnabledFunction_Once.Do(func() {
		err = webContextStruct_Set()
		if err != nil {
			return
		}
		webContextGetSpellCheckingEnabledFunction, err = webContextStruct.InvokerNew("get_spell_checking_enabled")
	})
	return err
}

// GetSpellCheckingEnabled is a representation of the C type webkit_web_context_get_spell_checking_enabled.
func (recv *WebContext) GetSpellCheckingEnabled() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := webContextGetSpellCheckingEnabledFunction_Set()
	if err == nil {
		ret = webContextGetSpellCheckingEnabledFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var webContextGetSpellCheckingLanguagesFunction *gi.Function
var webContextGetSpellCheckingLanguagesFunction_Once sync.Once

func webContextGetSpellCheckingLanguagesFunction_Set() error {
	var err error
	webContextGetSpellCheckingLanguagesFunction_Once.Do(func() {
		err = webContextStruct_Set()
		if err != nil {
			return
		}
		webContextGetSpellCheckingLanguagesFunction, err = webContextStruct.InvokerNew("get_spell_checking_languages")
	})
	return err
}

// GetSpellCheckingLanguages is a representation of the C type webkit_web_context_get_spell_checking_languages.
func (recv *WebContext) GetSpellCheckingLanguages() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := webContextGetSpellCheckingLanguagesFunction_Set()
	if err == nil {
		webContextGetSpellCheckingLanguagesFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'webkit_web_context_get_tls_errors_policy' : return type 'TLSErrorsPolicy' not supported

var webContextGetWebProcessCountLimitFunction *gi.Function
var webContextGetWebProcessCountLimitFunction_Once sync.Once

func webContextGetWebProcessCountLimitFunction_Set() error {
	var err error
	webContextGetWebProcessCountLimitFunction_Once.Do(func() {
		err = webContextStruct_Set()
		if err != nil {
			return
		}
		webContextGetWebProcessCountLimitFunction, err = webContextStruct.InvokerNew("get_web_process_count_limit")
	})
	return err
}

// GetWebProcessCountLimit is a representation of the C type webkit_web_context_get_web_process_count_limit.
func (recv *WebContext) GetWebProcessCountLimit() uint32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := webContextGetWebProcessCountLimitFunction_Set()
	if err == nil {
		ret = webContextGetWebProcessCountLimitFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo
}

// UNSUPPORTED : C value 'webkit_web_context_get_website_data_manager' : return type 'WebsiteDataManager' not supported

// UNSUPPORTED : C value 'webkit_web_context_initialize_notification_permissions' : parameter 'allowed_origins' of type 'GLib.List' not supported

var webContextIsAutomationAllowedFunction *gi.Function
var webContextIsAutomationAllowedFunction_Once sync.Once

func webContextIsAutomationAllowedFunction_Set() error {
	var err error
	webContextIsAutomationAllowedFunction_Once.Do(func() {
		err = webContextStruct_Set()
		if err != nil {
			return
		}
		webContextIsAutomationAllowedFunction, err = webContextStruct.InvokerNew("is_automation_allowed")
	})
	return err
}

// IsAutomationAllowed is a representation of the C type webkit_web_context_is_automation_allowed.
func (recv *WebContext) IsAutomationAllowed() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := webContextIsAutomationAllowedFunction_Set()
	if err == nil {
		ret = webContextIsAutomationAllowedFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var webContextIsEphemeralFunction *gi.Function
var webContextIsEphemeralFunction_Once sync.Once

func webContextIsEphemeralFunction_Set() error {
	var err error
	webContextIsEphemeralFunction_Once.Do(func() {
		err = webContextStruct_Set()
		if err != nil {
			return
		}
		webContextIsEphemeralFunction, err = webContextStruct.InvokerNew("is_ephemeral")
	})
	return err
}

// IsEphemeral is a representation of the C type webkit_web_context_is_ephemeral.
func (recv *WebContext) IsEphemeral() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := webContextIsEphemeralFunction_Set()
	if err == nil {
		ret = webContextIsEphemeralFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var webContextPrefetchDnsFunction *gi.Function
var webContextPrefetchDnsFunction_Once sync.Once

func webContextPrefetchDnsFunction_Set() error {
	var err error
	webContextPrefetchDnsFunction_Once.Do(func() {
		err = webContextStruct_Set()
		if err != nil {
			return
		}
		webContextPrefetchDnsFunction, err = webContextStruct.InvokerNew("prefetch_dns")
	})
	return err
}

// PrefetchDns is a representation of the C type webkit_web_context_prefetch_dns.
func (recv *WebContext) PrefetchDns(hostname string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(hostname)

	err := webContextPrefetchDnsFunction_Set()
	if err == nil {
		webContextPrefetchDnsFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'webkit_web_context_register_uri_scheme' : parameter 'callback' of type 'URISchemeRequestCallback' not supported

var webContextSetAdditionalPluginsDirectoryFunction *gi.Function
var webContextSetAdditionalPluginsDirectoryFunction_Once sync.Once

func webContextSetAdditionalPluginsDirectoryFunction_Set() error {
	var err error
	webContextSetAdditionalPluginsDirectoryFunction_Once.Do(func() {
		err = webContextStruct_Set()
		if err != nil {
			return
		}
		webContextSetAdditionalPluginsDirectoryFunction, err = webContextStruct.InvokerNew("set_additional_plugins_directory")
	})
	return err
}

// SetAdditionalPluginsDirectory is a representation of the C type webkit_web_context_set_additional_plugins_directory.
func (recv *WebContext) SetAdditionalPluginsDirectory(directory string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(directory)

	err := webContextSetAdditionalPluginsDirectoryFunction_Set()
	if err == nil {
		webContextSetAdditionalPluginsDirectoryFunction.Invoke(inArgs[:], nil)
	}

	return
}

var webContextSetAutomationAllowedFunction *gi.Function
var webContextSetAutomationAllowedFunction_Once sync.Once

func webContextSetAutomationAllowedFunction_Set() error {
	var err error
	webContextSetAutomationAllowedFunction_Once.Do(func() {
		err = webContextStruct_Set()
		if err != nil {
			return
		}
		webContextSetAutomationAllowedFunction, err = webContextStruct.InvokerNew("set_automation_allowed")
	})
	return err
}

// SetAutomationAllowed is a representation of the C type webkit_web_context_set_automation_allowed.
func (recv *WebContext) SetAutomationAllowed(allowed bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetBoolean(allowed)

	err := webContextSetAutomationAllowedFunction_Set()
	if err == nil {
		webContextSetAutomationAllowedFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'webkit_web_context_set_cache_model' : parameter 'cache_model' of type 'CacheModel' not supported

var webContextSetDiskCacheDirectoryFunction *gi.Function
var webContextSetDiskCacheDirectoryFunction_Once sync.Once

func webContextSetDiskCacheDirectoryFunction_Set() error {
	var err error
	webContextSetDiskCacheDirectoryFunction_Once.Do(func() {
		err = webContextStruct_Set()
		if err != nil {
			return
		}
		webContextSetDiskCacheDirectoryFunction, err = webContextStruct.InvokerNew("set_disk_cache_directory")
	})
	return err
}

// SetDiskCacheDirectory is a representation of the C type webkit_web_context_set_disk_cache_directory.
func (recv *WebContext) SetDiskCacheDirectory(directory string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(directory)

	err := webContextSetDiskCacheDirectoryFunction_Set()
	if err == nil {
		webContextSetDiskCacheDirectoryFunction.Invoke(inArgs[:], nil)
	}

	return
}

var webContextSetFaviconDatabaseDirectoryFunction *gi.Function
var webContextSetFaviconDatabaseDirectoryFunction_Once sync.Once

func webContextSetFaviconDatabaseDirectoryFunction_Set() error {
	var err error
	webContextSetFaviconDatabaseDirectoryFunction_Once.Do(func() {
		err = webContextStruct_Set()
		if err != nil {
			return
		}
		webContextSetFaviconDatabaseDirectoryFunction, err = webContextStruct.InvokerNew("set_favicon_database_directory")
	})
	return err
}

// SetFaviconDatabaseDirectory is a representation of the C type webkit_web_context_set_favicon_database_directory.
func (recv *WebContext) SetFaviconDatabaseDirectory(path string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(path)

	err := webContextSetFaviconDatabaseDirectoryFunction_Set()
	if err == nil {
		webContextSetFaviconDatabaseDirectoryFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'webkit_web_context_set_network_proxy_settings' : parameter 'proxy_mode' of type 'NetworkProxyMode' not supported

// UNSUPPORTED : C value 'webkit_web_context_set_preferred_languages' : parameter 'languages' of type 'nil' not supported

// UNSUPPORTED : C value 'webkit_web_context_set_process_model' : parameter 'process_model' of type 'ProcessModel' not supported

var webContextSetSandboxEnabledFunction *gi.Function
var webContextSetSandboxEnabledFunction_Once sync.Once

func webContextSetSandboxEnabledFunction_Set() error {
	var err error
	webContextSetSandboxEnabledFunction_Once.Do(func() {
		err = webContextStruct_Set()
		if err != nil {
			return
		}
		webContextSetSandboxEnabledFunction, err = webContextStruct.InvokerNew("set_sandbox_enabled")
	})
	return err
}

// SetSandboxEnabled is a representation of the C type webkit_web_context_set_sandbox_enabled.
func (recv *WebContext) SetSandboxEnabled(enabled bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetBoolean(enabled)

	err := webContextSetSandboxEnabledFunction_Set()
	if err == nil {
		webContextSetSandboxEnabledFunction.Invoke(inArgs[:], nil)
	}

	return
}

var webContextSetSpellCheckingEnabledFunction *gi.Function
var webContextSetSpellCheckingEnabledFunction_Once sync.Once

func webContextSetSpellCheckingEnabledFunction_Set() error {
	var err error
	webContextSetSpellCheckingEnabledFunction_Once.Do(func() {
		err = webContextStruct_Set()
		if err != nil {
			return
		}
		webContextSetSpellCheckingEnabledFunction, err = webContextStruct.InvokerNew("set_spell_checking_enabled")
	})
	return err
}

// SetSpellCheckingEnabled is a representation of the C type webkit_web_context_set_spell_checking_enabled.
func (recv *WebContext) SetSpellCheckingEnabled(enabled bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetBoolean(enabled)

	err := webContextSetSpellCheckingEnabledFunction_Set()
	if err == nil {
		webContextSetSpellCheckingEnabledFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'webkit_web_context_set_spell_checking_languages' : parameter 'languages' of type 'nil' not supported

// UNSUPPORTED : C value 'webkit_web_context_set_tls_errors_policy' : parameter 'policy' of type 'TLSErrorsPolicy' not supported

var webContextSetWebExtensionsDirectoryFunction *gi.Function
var webContextSetWebExtensionsDirectoryFunction_Once sync.Once

func webContextSetWebExtensionsDirectoryFunction_Set() error {
	var err error
	webContextSetWebExtensionsDirectoryFunction_Once.Do(func() {
		err = webContextStruct_Set()
		if err != nil {
			return
		}
		webContextSetWebExtensionsDirectoryFunction, err = webContextStruct.InvokerNew("set_web_extensions_directory")
	})
	return err
}

// SetWebExtensionsDirectory is a representation of the C type webkit_web_context_set_web_extensions_directory.
func (recv *WebContext) SetWebExtensionsDirectory(directory string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(directory)

	err := webContextSetWebExtensionsDirectoryFunction_Set()
	if err == nil {
		webContextSetWebExtensionsDirectoryFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'webkit_web_context_set_web_extensions_initialization_user_data' : parameter 'user_data' of type 'GLib.Variant' not supported

var webContextSetWebProcessCountLimitFunction *gi.Function
var webContextSetWebProcessCountLimitFunction_Once sync.Once

func webContextSetWebProcessCountLimitFunction_Set() error {
	var err error
	webContextSetWebProcessCountLimitFunction_Once.Do(func() {
		err = webContextStruct_Set()
		if err != nil {
			return
		}
		webContextSetWebProcessCountLimitFunction, err = webContextStruct.InvokerNew("set_web_process_count_limit")
	})
	return err
}

// SetWebProcessCountLimit is a representation of the C type webkit_web_context_set_web_process_count_limit.
func (recv *WebContext) SetWebProcessCountLimit(limit uint32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetUint32(limit)

	err := webContextSetWebProcessCountLimitFunction_Set()
	if err == nil {
		webContextSetWebProcessCountLimitFunction.Invoke(inArgs[:], nil)
	}

	return
}

var webInspectorStruct *gi.Struct
var webInspectorStruct_Once sync.Once

func webInspectorStruct_Set() error {
	var err error
	webInspectorStruct_Once.Do(func() {
		webInspectorStruct, err = gi.StructNew("WebKit2", "WebInspector")
	})
	return err
}

type WebInspector struct {
	native uintptr
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'GObject.Object'

// FieldPriv returns the C field 'priv'.
func (recv *WebInspector) FieldPriv() *WebInspectorPrivate {
	argValue := gi.FieldGet(webInspectorStruct, recv.native, "priv")
	value := &WebInspectorPrivate{native: argValue.Pointer()}
	return value
}

// SetFieldPriv sets the value of the C field 'priv'.
func (recv *WebInspector) SetFieldPriv(value *WebInspectorPrivate) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(webInspectorStruct, recv.native, "priv", argValue)
}

var webInspectorAttachFunction *gi.Function
var webInspectorAttachFunction_Once sync.Once

func webInspectorAttachFunction_Set() error {
	var err error
	webInspectorAttachFunction_Once.Do(func() {
		err = webInspectorStruct_Set()
		if err != nil {
			return
		}
		webInspectorAttachFunction, err = webInspectorStruct.InvokerNew("attach")
	})
	return err
}

// Attach is a representation of the C type webkit_web_inspector_attach.
func (recv *WebInspector) Attach() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := webInspectorAttachFunction_Set()
	if err == nil {
		webInspectorAttachFunction.Invoke(inArgs[:], nil)
	}

	return
}

var webInspectorCloseFunction *gi.Function
var webInspectorCloseFunction_Once sync.Once

func webInspectorCloseFunction_Set() error {
	var err error
	webInspectorCloseFunction_Once.Do(func() {
		err = webInspectorStruct_Set()
		if err != nil {
			return
		}
		webInspectorCloseFunction, err = webInspectorStruct.InvokerNew("close")
	})
	return err
}

// Close is a representation of the C type webkit_web_inspector_close.
func (recv *WebInspector) Close() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := webInspectorCloseFunction_Set()
	if err == nil {
		webInspectorCloseFunction.Invoke(inArgs[:], nil)
	}

	return
}

var webInspectorDetachFunction *gi.Function
var webInspectorDetachFunction_Once sync.Once

func webInspectorDetachFunction_Set() error {
	var err error
	webInspectorDetachFunction_Once.Do(func() {
		err = webInspectorStruct_Set()
		if err != nil {
			return
		}
		webInspectorDetachFunction, err = webInspectorStruct.InvokerNew("detach")
	})
	return err
}

// Detach is a representation of the C type webkit_web_inspector_detach.
func (recv *WebInspector) Detach() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := webInspectorDetachFunction_Set()
	if err == nil {
		webInspectorDetachFunction.Invoke(inArgs[:], nil)
	}

	return
}

var webInspectorGetAttachedHeightFunction *gi.Function
var webInspectorGetAttachedHeightFunction_Once sync.Once

func webInspectorGetAttachedHeightFunction_Set() error {
	var err error
	webInspectorGetAttachedHeightFunction_Once.Do(func() {
		err = webInspectorStruct_Set()
		if err != nil {
			return
		}
		webInspectorGetAttachedHeightFunction, err = webInspectorStruct.InvokerNew("get_attached_height")
	})
	return err
}

// GetAttachedHeight is a representation of the C type webkit_web_inspector_get_attached_height.
func (recv *WebInspector) GetAttachedHeight() uint32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := webInspectorGetAttachedHeightFunction_Set()
	if err == nil {
		ret = webInspectorGetAttachedHeightFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo
}

var webInspectorGetCanAttachFunction *gi.Function
var webInspectorGetCanAttachFunction_Once sync.Once

func webInspectorGetCanAttachFunction_Set() error {
	var err error
	webInspectorGetCanAttachFunction_Once.Do(func() {
		err = webInspectorStruct_Set()
		if err != nil {
			return
		}
		webInspectorGetCanAttachFunction, err = webInspectorStruct.InvokerNew("get_can_attach")
	})
	return err
}

// GetCanAttach is a representation of the C type webkit_web_inspector_get_can_attach.
func (recv *WebInspector) GetCanAttach() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := webInspectorGetCanAttachFunction_Set()
	if err == nil {
		ret = webInspectorGetCanAttachFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var webInspectorGetInspectedUriFunction *gi.Function
var webInspectorGetInspectedUriFunction_Once sync.Once

func webInspectorGetInspectedUriFunction_Set() error {
	var err error
	webInspectorGetInspectedUriFunction_Once.Do(func() {
		err = webInspectorStruct_Set()
		if err != nil {
			return
		}
		webInspectorGetInspectedUriFunction, err = webInspectorStruct.InvokerNew("get_inspected_uri")
	})
	return err
}

// GetInspectedUri is a representation of the C type webkit_web_inspector_get_inspected_uri.
func (recv *WebInspector) GetInspectedUri() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := webInspectorGetInspectedUriFunction_Set()
	if err == nil {
		ret = webInspectorGetInspectedUriFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'webkit_web_inspector_get_web_view' : return type 'WebViewBase' not supported

var webInspectorIsAttachedFunction *gi.Function
var webInspectorIsAttachedFunction_Once sync.Once

func webInspectorIsAttachedFunction_Set() error {
	var err error
	webInspectorIsAttachedFunction_Once.Do(func() {
		err = webInspectorStruct_Set()
		if err != nil {
			return
		}
		webInspectorIsAttachedFunction, err = webInspectorStruct.InvokerNew("is_attached")
	})
	return err
}

// IsAttached is a representation of the C type webkit_web_inspector_is_attached.
func (recv *WebInspector) IsAttached() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := webInspectorIsAttachedFunction_Set()
	if err == nil {
		ret = webInspectorIsAttachedFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var webInspectorShowFunction *gi.Function
var webInspectorShowFunction_Once sync.Once

func webInspectorShowFunction_Set() error {
	var err error
	webInspectorShowFunction_Once.Do(func() {
		err = webInspectorStruct_Set()
		if err != nil {
			return
		}
		webInspectorShowFunction, err = webInspectorStruct.InvokerNew("show")
	})
	return err
}

// Show is a representation of the C type webkit_web_inspector_show.
func (recv *WebInspector) Show() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := webInspectorShowFunction_Set()
	if err == nil {
		webInspectorShowFunction.Invoke(inArgs[:], nil)
	}

	return
}

// WebInspectorStruct creates an uninitialised WebInspector.
func WebInspectorStruct() *WebInspector {
	err := webInspectorStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &WebInspector{native: webInspectorStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeWebInspector)
	return structGo
}
func finalizeWebInspector(obj *WebInspector) {
	webInspectorStruct.Free(obj.native)
}

var webResourceStruct *gi.Struct
var webResourceStruct_Once sync.Once

func webResourceStruct_Set() error {
	var err error
	webResourceStruct_Once.Do(func() {
		webResourceStruct, err = gi.StructNew("WebKit2", "WebResource")
	})
	return err
}

type WebResource struct {
	native uintptr
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'GObject.Object'

// FieldPriv returns the C field 'priv'.
func (recv *WebResource) FieldPriv() *WebResourcePrivate {
	argValue := gi.FieldGet(webResourceStruct, recv.native, "priv")
	value := &WebResourcePrivate{native: argValue.Pointer()}
	return value
}

// SetFieldPriv sets the value of the C field 'priv'.
func (recv *WebResource) SetFieldPriv(value *WebResourcePrivate) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(webResourceStruct, recv.native, "priv", argValue)
}

// UNSUPPORTED : C value 'webkit_web_resource_get_data' : parameter 'cancellable' of type 'Gio.Cancellable' not supported

// UNSUPPORTED : C value 'webkit_web_resource_get_data_finish' : parameter 'result' of type 'Gio.AsyncResult' not supported

// UNSUPPORTED : C value 'webkit_web_resource_get_response' : return type 'URIResponse' not supported

var webResourceGetUriFunction *gi.Function
var webResourceGetUriFunction_Once sync.Once

func webResourceGetUriFunction_Set() error {
	var err error
	webResourceGetUriFunction_Once.Do(func() {
		err = webResourceStruct_Set()
		if err != nil {
			return
		}
		webResourceGetUriFunction, err = webResourceStruct.InvokerNew("get_uri")
	})
	return err
}

// GetUri is a representation of the C type webkit_web_resource_get_uri.
func (recv *WebResource) GetUri() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := webResourceGetUriFunction_Set()
	if err == nil {
		ret = webResourceGetUriFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

// WebResourceStruct creates an uninitialised WebResource.
func WebResourceStruct() *WebResource {
	err := webResourceStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &WebResource{native: webResourceStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeWebResource)
	return structGo
}
func finalizeWebResource(obj *WebResource) {
	webResourceStruct.Free(obj.native)
}

var webViewStruct *gi.Struct
var webViewStruct_Once sync.Once

func webViewStruct_Set() error {
	var err error
	webViewStruct_Once.Do(func() {
		webViewStruct, err = gi.StructNew("WebKit2", "WebView")
	})
	return err
}

type WebView struct {
	native uintptr
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'WebViewBase'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'WebViewBase'

// UNSUPPORTED : C value 'webkit_web_view_new' : return type 'WebView' not supported

// UNSUPPORTED : C value 'webkit_web_view_new_with_context' : parameter 'context' of type 'WebContext' not supported

// UNSUPPORTED : C value 'webkit_web_view_new_with_related_view' : parameter 'web_view' of type 'WebView' not supported

// UNSUPPORTED : C value 'webkit_web_view_new_with_settings' : parameter 'settings' of type 'Settings' not supported

// UNSUPPORTED : C value 'webkit_web_view_new_with_user_content_manager' : parameter 'user_content_manager' of type 'UserContentManager' not supported

// UNSUPPORTED : C value 'webkit_web_view_can_execute_editing_command' : parameter 'cancellable' of type 'Gio.Cancellable' not supported

// UNSUPPORTED : C value 'webkit_web_view_can_execute_editing_command_finish' : parameter 'result' of type 'Gio.AsyncResult' not supported

var webViewCanGoBackFunction *gi.Function
var webViewCanGoBackFunction_Once sync.Once

func webViewCanGoBackFunction_Set() error {
	var err error
	webViewCanGoBackFunction_Once.Do(func() {
		err = webViewStruct_Set()
		if err != nil {
			return
		}
		webViewCanGoBackFunction, err = webViewStruct.InvokerNew("can_go_back")
	})
	return err
}

// CanGoBack is a representation of the C type webkit_web_view_can_go_back.
func (recv *WebView) CanGoBack() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := webViewCanGoBackFunction_Set()
	if err == nil {
		ret = webViewCanGoBackFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var webViewCanGoForwardFunction *gi.Function
var webViewCanGoForwardFunction_Once sync.Once

func webViewCanGoForwardFunction_Set() error {
	var err error
	webViewCanGoForwardFunction_Once.Do(func() {
		err = webViewStruct_Set()
		if err != nil {
			return
		}
		webViewCanGoForwardFunction, err = webViewStruct.InvokerNew("can_go_forward")
	})
	return err
}

// CanGoForward is a representation of the C type webkit_web_view_can_go_forward.
func (recv *WebView) CanGoForward() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := webViewCanGoForwardFunction_Set()
	if err == nil {
		ret = webViewCanGoForwardFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var webViewCanShowMimeTypeFunction *gi.Function
var webViewCanShowMimeTypeFunction_Once sync.Once

func webViewCanShowMimeTypeFunction_Set() error {
	var err error
	webViewCanShowMimeTypeFunction_Once.Do(func() {
		err = webViewStruct_Set()
		if err != nil {
			return
		}
		webViewCanShowMimeTypeFunction, err = webViewStruct.InvokerNew("can_show_mime_type")
	})
	return err
}

// CanShowMimeType is a representation of the C type webkit_web_view_can_show_mime_type.
func (recv *WebView) CanShowMimeType(mimeType string) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(mimeType)

	var ret gi.Argument

	err := webViewCanShowMimeTypeFunction_Set()
	if err == nil {
		ret = webViewCanShowMimeTypeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'webkit_web_view_download_uri' : return type 'Download' not supported

var webViewExecuteEditingCommandFunction *gi.Function
var webViewExecuteEditingCommandFunction_Once sync.Once

func webViewExecuteEditingCommandFunction_Set() error {
	var err error
	webViewExecuteEditingCommandFunction_Once.Do(func() {
		err = webViewStruct_Set()
		if err != nil {
			return
		}
		webViewExecuteEditingCommandFunction, err = webViewStruct.InvokerNew("execute_editing_command")
	})
	return err
}

// ExecuteEditingCommand is a representation of the C type webkit_web_view_execute_editing_command.
func (recv *WebView) ExecuteEditingCommand(command string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(command)

	err := webViewExecuteEditingCommandFunction_Set()
	if err == nil {
		webViewExecuteEditingCommandFunction.Invoke(inArgs[:], nil)
	}

	return
}

var webViewExecuteEditingCommandWithArgumentFunction *gi.Function
var webViewExecuteEditingCommandWithArgumentFunction_Once sync.Once

func webViewExecuteEditingCommandWithArgumentFunction_Set() error {
	var err error
	webViewExecuteEditingCommandWithArgumentFunction_Once.Do(func() {
		err = webViewStruct_Set()
		if err != nil {
			return
		}
		webViewExecuteEditingCommandWithArgumentFunction, err = webViewStruct.InvokerNew("execute_editing_command_with_argument")
	})
	return err
}

// ExecuteEditingCommandWithArgument is a representation of the C type webkit_web_view_execute_editing_command_with_argument.
func (recv *WebView) ExecuteEditingCommandWithArgument(command string, argument string) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(command)
	inArgs[2].SetString(argument)

	err := webViewExecuteEditingCommandWithArgumentFunction_Set()
	if err == nil {
		webViewExecuteEditingCommandWithArgumentFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'webkit_web_view_get_back_forward_list' : return type 'BackForwardList' not supported

// UNSUPPORTED : C value 'webkit_web_view_get_background_color' : parameter 'rgba' of type 'Gdk.RGBA' not supported

// UNSUPPORTED : C value 'webkit_web_view_get_context' : return type 'WebContext' not supported

var webViewGetCustomCharsetFunction *gi.Function
var webViewGetCustomCharsetFunction_Once sync.Once

func webViewGetCustomCharsetFunction_Set() error {
	var err error
	webViewGetCustomCharsetFunction_Once.Do(func() {
		err = webViewStruct_Set()
		if err != nil {
			return
		}
		webViewGetCustomCharsetFunction, err = webViewStruct.InvokerNew("get_custom_charset")
	})
	return err
}

// GetCustomCharset is a representation of the C type webkit_web_view_get_custom_charset.
func (recv *WebView) GetCustomCharset() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := webViewGetCustomCharsetFunction_Set()
	if err == nil {
		ret = webViewGetCustomCharsetFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'webkit_web_view_get_editor_state' : return type 'EditorState' not supported

var webViewGetEstimatedLoadProgressFunction *gi.Function
var webViewGetEstimatedLoadProgressFunction_Once sync.Once

func webViewGetEstimatedLoadProgressFunction_Set() error {
	var err error
	webViewGetEstimatedLoadProgressFunction_Once.Do(func() {
		err = webViewStruct_Set()
		if err != nil {
			return
		}
		webViewGetEstimatedLoadProgressFunction, err = webViewStruct.InvokerNew("get_estimated_load_progress")
	})
	return err
}

// GetEstimatedLoadProgress is a representation of the C type webkit_web_view_get_estimated_load_progress.
func (recv *WebView) GetEstimatedLoadProgress() float64 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := webViewGetEstimatedLoadProgressFunction_Set()
	if err == nil {
		ret = webViewGetEstimatedLoadProgressFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Float64()

	return retGo
}

// UNSUPPORTED : C value 'webkit_web_view_get_favicon' : return type 'cairo.Surface' not supported

// UNSUPPORTED : C value 'webkit_web_view_get_find_controller' : return type 'FindController' not supported

// UNSUPPORTED : C value 'webkit_web_view_get_inspector' : return type 'WebInspector' not supported

// UNSUPPORTED : C value 'webkit_web_view_get_javascript_global_context' : return type 'JavaScriptCore.GlobalContextRef' not supported

// UNSUPPORTED : C value 'webkit_web_view_get_main_resource' : return type 'WebResource' not supported

var webViewGetPageIdFunction *gi.Function
var webViewGetPageIdFunction_Once sync.Once

func webViewGetPageIdFunction_Set() error {
	var err error
	webViewGetPageIdFunction_Once.Do(func() {
		err = webViewStruct_Set()
		if err != nil {
			return
		}
		webViewGetPageIdFunction, err = webViewStruct.InvokerNew("get_page_id")
	})
	return err
}

// GetPageId is a representation of the C type webkit_web_view_get_page_id.
func (recv *WebView) GetPageId() uint64 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := webViewGetPageIdFunction_Set()
	if err == nil {
		ret = webViewGetPageIdFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint64()

	return retGo
}

var webViewGetSessionStateFunction *gi.Function
var webViewGetSessionStateFunction_Once sync.Once

func webViewGetSessionStateFunction_Set() error {
	var err error
	webViewGetSessionStateFunction_Once.Do(func() {
		err = webViewStruct_Set()
		if err != nil {
			return
		}
		webViewGetSessionStateFunction, err = webViewStruct.InvokerNew("get_session_state")
	})
	return err
}

// GetSessionState is a representation of the C type webkit_web_view_get_session_state.
func (recv *WebView) GetSessionState() *WebViewSessionState {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := webViewGetSessionStateFunction_Set()
	if err == nil {
		ret = webViewGetSessionStateFunction.Invoke(inArgs[:], nil)
	}

	retGo := &WebViewSessionState{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'webkit_web_view_get_settings' : return type 'Settings' not supported

// UNSUPPORTED : C value 'webkit_web_view_get_snapshot' : parameter 'region' of type 'SnapshotRegion' not supported

// UNSUPPORTED : C value 'webkit_web_view_get_snapshot_finish' : parameter 'result' of type 'Gio.AsyncResult' not supported

var webViewGetTitleFunction *gi.Function
var webViewGetTitleFunction_Once sync.Once

func webViewGetTitleFunction_Set() error {
	var err error
	webViewGetTitleFunction_Once.Do(func() {
		err = webViewStruct_Set()
		if err != nil {
			return
		}
		webViewGetTitleFunction, err = webViewStruct.InvokerNew("get_title")
	})
	return err
}

// GetTitle is a representation of the C type webkit_web_view_get_title.
func (recv *WebView) GetTitle() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := webViewGetTitleFunction_Set()
	if err == nil {
		ret = webViewGetTitleFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'webkit_web_view_get_tls_info' : parameter 'certificate' of type 'Gio.TlsCertificate' not supported

var webViewGetUriFunction *gi.Function
var webViewGetUriFunction_Once sync.Once

func webViewGetUriFunction_Set() error {
	var err error
	webViewGetUriFunction_Once.Do(func() {
		err = webViewStruct_Set()
		if err != nil {
			return
		}
		webViewGetUriFunction, err = webViewStruct.InvokerNew("get_uri")
	})
	return err
}

// GetUri is a representation of the C type webkit_web_view_get_uri.
func (recv *WebView) GetUri() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := webViewGetUriFunction_Set()
	if err == nil {
		ret = webViewGetUriFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'webkit_web_view_get_user_content_manager' : return type 'UserContentManager' not supported

// UNSUPPORTED : C value 'webkit_web_view_get_website_data_manager' : return type 'WebsiteDataManager' not supported

// UNSUPPORTED : C value 'webkit_web_view_get_window_properties' : return type 'WindowProperties' not supported

var webViewGetZoomLevelFunction *gi.Function
var webViewGetZoomLevelFunction_Once sync.Once

func webViewGetZoomLevelFunction_Set() error {
	var err error
	webViewGetZoomLevelFunction_Once.Do(func() {
		err = webViewStruct_Set()
		if err != nil {
			return
		}
		webViewGetZoomLevelFunction, err = webViewStruct.InvokerNew("get_zoom_level")
	})
	return err
}

// GetZoomLevel is a representation of the C type webkit_web_view_get_zoom_level.
func (recv *WebView) GetZoomLevel() float64 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := webViewGetZoomLevelFunction_Set()
	if err == nil {
		ret = webViewGetZoomLevelFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Float64()

	return retGo
}

var webViewGoBackFunction *gi.Function
var webViewGoBackFunction_Once sync.Once

func webViewGoBackFunction_Set() error {
	var err error
	webViewGoBackFunction_Once.Do(func() {
		err = webViewStruct_Set()
		if err != nil {
			return
		}
		webViewGoBackFunction, err = webViewStruct.InvokerNew("go_back")
	})
	return err
}

// GoBack is a representation of the C type webkit_web_view_go_back.
func (recv *WebView) GoBack() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := webViewGoBackFunction_Set()
	if err == nil {
		webViewGoBackFunction.Invoke(inArgs[:], nil)
	}

	return
}

var webViewGoForwardFunction *gi.Function
var webViewGoForwardFunction_Once sync.Once

func webViewGoForwardFunction_Set() error {
	var err error
	webViewGoForwardFunction_Once.Do(func() {
		err = webViewStruct_Set()
		if err != nil {
			return
		}
		webViewGoForwardFunction, err = webViewStruct.InvokerNew("go_forward")
	})
	return err
}

// GoForward is a representation of the C type webkit_web_view_go_forward.
func (recv *WebView) GoForward() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := webViewGoForwardFunction_Set()
	if err == nil {
		webViewGoForwardFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'webkit_web_view_go_to_back_forward_list_item' : parameter 'list_item' of type 'BackForwardListItem' not supported

var webViewIsControlledByAutomationFunction *gi.Function
var webViewIsControlledByAutomationFunction_Once sync.Once

func webViewIsControlledByAutomationFunction_Set() error {
	var err error
	webViewIsControlledByAutomationFunction_Once.Do(func() {
		err = webViewStruct_Set()
		if err != nil {
			return
		}
		webViewIsControlledByAutomationFunction, err = webViewStruct.InvokerNew("is_controlled_by_automation")
	})
	return err
}

// IsControlledByAutomation is a representation of the C type webkit_web_view_is_controlled_by_automation.
func (recv *WebView) IsControlledByAutomation() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := webViewIsControlledByAutomationFunction_Set()
	if err == nil {
		ret = webViewIsControlledByAutomationFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var webViewIsEditableFunction *gi.Function
var webViewIsEditableFunction_Once sync.Once

func webViewIsEditableFunction_Set() error {
	var err error
	webViewIsEditableFunction_Once.Do(func() {
		err = webViewStruct_Set()
		if err != nil {
			return
		}
		webViewIsEditableFunction, err = webViewStruct.InvokerNew("is_editable")
	})
	return err
}

// IsEditable is a representation of the C type webkit_web_view_is_editable.
func (recv *WebView) IsEditable() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := webViewIsEditableFunction_Set()
	if err == nil {
		ret = webViewIsEditableFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var webViewIsEphemeralFunction *gi.Function
var webViewIsEphemeralFunction_Once sync.Once

func webViewIsEphemeralFunction_Set() error {
	var err error
	webViewIsEphemeralFunction_Once.Do(func() {
		err = webViewStruct_Set()
		if err != nil {
			return
		}
		webViewIsEphemeralFunction, err = webViewStruct.InvokerNew("is_ephemeral")
	})
	return err
}

// IsEphemeral is a representation of the C type webkit_web_view_is_ephemeral.
func (recv *WebView) IsEphemeral() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := webViewIsEphemeralFunction_Set()
	if err == nil {
		ret = webViewIsEphemeralFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var webViewIsLoadingFunction *gi.Function
var webViewIsLoadingFunction_Once sync.Once

func webViewIsLoadingFunction_Set() error {
	var err error
	webViewIsLoadingFunction_Once.Do(func() {
		err = webViewStruct_Set()
		if err != nil {
			return
		}
		webViewIsLoadingFunction, err = webViewStruct.InvokerNew("is_loading")
	})
	return err
}

// IsLoading is a representation of the C type webkit_web_view_is_loading.
func (recv *WebView) IsLoading() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := webViewIsLoadingFunction_Set()
	if err == nil {
		ret = webViewIsLoadingFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var webViewIsPlayingAudioFunction *gi.Function
var webViewIsPlayingAudioFunction_Once sync.Once

func webViewIsPlayingAudioFunction_Set() error {
	var err error
	webViewIsPlayingAudioFunction_Once.Do(func() {
		err = webViewStruct_Set()
		if err != nil {
			return
		}
		webViewIsPlayingAudioFunction, err = webViewStruct.InvokerNew("is_playing_audio")
	})
	return err
}

// IsPlayingAudio is a representation of the C type webkit_web_view_is_playing_audio.
func (recv *WebView) IsPlayingAudio() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := webViewIsPlayingAudioFunction_Set()
	if err == nil {
		ret = webViewIsPlayingAudioFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var webViewLoadAlternateHtmlFunction *gi.Function
var webViewLoadAlternateHtmlFunction_Once sync.Once

func webViewLoadAlternateHtmlFunction_Set() error {
	var err error
	webViewLoadAlternateHtmlFunction_Once.Do(func() {
		err = webViewStruct_Set()
		if err != nil {
			return
		}
		webViewLoadAlternateHtmlFunction, err = webViewStruct.InvokerNew("load_alternate_html")
	})
	return err
}

// LoadAlternateHtml is a representation of the C type webkit_web_view_load_alternate_html.
func (recv *WebView) LoadAlternateHtml(content string, contentUri string, baseUri string) {
	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(content)
	inArgs[2].SetString(contentUri)
	inArgs[3].SetString(baseUri)

	err := webViewLoadAlternateHtmlFunction_Set()
	if err == nil {
		webViewLoadAlternateHtmlFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'webkit_web_view_load_bytes' : parameter 'bytes' of type 'GLib.Bytes' not supported

var webViewLoadHtmlFunction *gi.Function
var webViewLoadHtmlFunction_Once sync.Once

func webViewLoadHtmlFunction_Set() error {
	var err error
	webViewLoadHtmlFunction_Once.Do(func() {
		err = webViewStruct_Set()
		if err != nil {
			return
		}
		webViewLoadHtmlFunction, err = webViewStruct.InvokerNew("load_html")
	})
	return err
}

// LoadHtml is a representation of the C type webkit_web_view_load_html.
func (recv *WebView) LoadHtml(content string, baseUri string) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(content)
	inArgs[2].SetString(baseUri)

	err := webViewLoadHtmlFunction_Set()
	if err == nil {
		webViewLoadHtmlFunction.Invoke(inArgs[:], nil)
	}

	return
}

var webViewLoadPlainTextFunction *gi.Function
var webViewLoadPlainTextFunction_Once sync.Once

func webViewLoadPlainTextFunction_Set() error {
	var err error
	webViewLoadPlainTextFunction_Once.Do(func() {
		err = webViewStruct_Set()
		if err != nil {
			return
		}
		webViewLoadPlainTextFunction, err = webViewStruct.InvokerNew("load_plain_text")
	})
	return err
}

// LoadPlainText is a representation of the C type webkit_web_view_load_plain_text.
func (recv *WebView) LoadPlainText(plainText string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(plainText)

	err := webViewLoadPlainTextFunction_Set()
	if err == nil {
		webViewLoadPlainTextFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'webkit_web_view_load_request' : parameter 'request' of type 'URIRequest' not supported

var webViewLoadUriFunction *gi.Function
var webViewLoadUriFunction_Once sync.Once

func webViewLoadUriFunction_Set() error {
	var err error
	webViewLoadUriFunction_Once.Do(func() {
		err = webViewStruct_Set()
		if err != nil {
			return
		}
		webViewLoadUriFunction, err = webViewStruct.InvokerNew("load_uri")
	})
	return err
}

// LoadUri is a representation of the C type webkit_web_view_load_uri.
func (recv *WebView) LoadUri(uri string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(uri)

	err := webViewLoadUriFunction_Set()
	if err == nil {
		webViewLoadUriFunction.Invoke(inArgs[:], nil)
	}

	return
}

var webViewReloadFunction *gi.Function
var webViewReloadFunction_Once sync.Once

func webViewReloadFunction_Set() error {
	var err error
	webViewReloadFunction_Once.Do(func() {
		err = webViewStruct_Set()
		if err != nil {
			return
		}
		webViewReloadFunction, err = webViewStruct.InvokerNew("reload")
	})
	return err
}

// Reload is a representation of the C type webkit_web_view_reload.
func (recv *WebView) Reload() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := webViewReloadFunction_Set()
	if err == nil {
		webViewReloadFunction.Invoke(inArgs[:], nil)
	}

	return
}

var webViewReloadBypassCacheFunction *gi.Function
var webViewReloadBypassCacheFunction_Once sync.Once

func webViewReloadBypassCacheFunction_Set() error {
	var err error
	webViewReloadBypassCacheFunction_Once.Do(func() {
		err = webViewStruct_Set()
		if err != nil {
			return
		}
		webViewReloadBypassCacheFunction, err = webViewStruct.InvokerNew("reload_bypass_cache")
	})
	return err
}

// ReloadBypassCache is a representation of the C type webkit_web_view_reload_bypass_cache.
func (recv *WebView) ReloadBypassCache() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := webViewReloadBypassCacheFunction_Set()
	if err == nil {
		webViewReloadBypassCacheFunction.Invoke(inArgs[:], nil)
	}

	return
}

var webViewRestoreSessionStateFunction *gi.Function
var webViewRestoreSessionStateFunction_Once sync.Once

func webViewRestoreSessionStateFunction_Set() error {
	var err error
	webViewRestoreSessionStateFunction_Once.Do(func() {
		err = webViewStruct_Set()
		if err != nil {
			return
		}
		webViewRestoreSessionStateFunction, err = webViewStruct.InvokerNew("restore_session_state")
	})
	return err
}

// RestoreSessionState is a representation of the C type webkit_web_view_restore_session_state.
func (recv *WebView) RestoreSessionState(state *WebViewSessionState) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(state.native)

	err := webViewRestoreSessionStateFunction_Set()
	if err == nil {
		webViewRestoreSessionStateFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'webkit_web_view_run_javascript' : parameter 'cancellable' of type 'Gio.Cancellable' not supported

// UNSUPPORTED : C value 'webkit_web_view_run_javascript_finish' : parameter 'result' of type 'Gio.AsyncResult' not supported

// UNSUPPORTED : C value 'webkit_web_view_run_javascript_from_gresource' : parameter 'cancellable' of type 'Gio.Cancellable' not supported

// UNSUPPORTED : C value 'webkit_web_view_run_javascript_from_gresource_finish' : parameter 'result' of type 'Gio.AsyncResult' not supported

// UNSUPPORTED : C value 'webkit_web_view_run_javascript_in_world' : parameter 'cancellable' of type 'Gio.Cancellable' not supported

// UNSUPPORTED : C value 'webkit_web_view_run_javascript_in_world_finish' : parameter 'result' of type 'Gio.AsyncResult' not supported

// UNSUPPORTED : C value 'webkit_web_view_save' : parameter 'save_mode' of type 'SaveMode' not supported

// UNSUPPORTED : C value 'webkit_web_view_save_finish' : parameter 'result' of type 'Gio.AsyncResult' not supported

// UNSUPPORTED : C value 'webkit_web_view_save_to_file' : parameter 'file' of type 'Gio.File' not supported

// UNSUPPORTED : C value 'webkit_web_view_save_to_file_finish' : parameter 'result' of type 'Gio.AsyncResult' not supported

// UNSUPPORTED : C value 'webkit_web_view_set_background_color' : parameter 'rgba' of type 'Gdk.RGBA' not supported

var webViewSetCustomCharsetFunction *gi.Function
var webViewSetCustomCharsetFunction_Once sync.Once

func webViewSetCustomCharsetFunction_Set() error {
	var err error
	webViewSetCustomCharsetFunction_Once.Do(func() {
		err = webViewStruct_Set()
		if err != nil {
			return
		}
		webViewSetCustomCharsetFunction, err = webViewStruct.InvokerNew("set_custom_charset")
	})
	return err
}

// SetCustomCharset is a representation of the C type webkit_web_view_set_custom_charset.
func (recv *WebView) SetCustomCharset(charset string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(charset)

	err := webViewSetCustomCharsetFunction_Set()
	if err == nil {
		webViewSetCustomCharsetFunction.Invoke(inArgs[:], nil)
	}

	return
}

var webViewSetEditableFunction *gi.Function
var webViewSetEditableFunction_Once sync.Once

func webViewSetEditableFunction_Set() error {
	var err error
	webViewSetEditableFunction_Once.Do(func() {
		err = webViewStruct_Set()
		if err != nil {
			return
		}
		webViewSetEditableFunction, err = webViewStruct.InvokerNew("set_editable")
	})
	return err
}

// SetEditable is a representation of the C type webkit_web_view_set_editable.
func (recv *WebView) SetEditable(editable bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetBoolean(editable)

	err := webViewSetEditableFunction_Set()
	if err == nil {
		webViewSetEditableFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'webkit_web_view_set_settings' : parameter 'settings' of type 'Settings' not supported

var webViewSetZoomLevelFunction *gi.Function
var webViewSetZoomLevelFunction_Once sync.Once

func webViewSetZoomLevelFunction_Set() error {
	var err error
	webViewSetZoomLevelFunction_Once.Do(func() {
		err = webViewStruct_Set()
		if err != nil {
			return
		}
		webViewSetZoomLevelFunction, err = webViewStruct.InvokerNew("set_zoom_level")
	})
	return err
}

// SetZoomLevel is a representation of the C type webkit_web_view_set_zoom_level.
func (recv *WebView) SetZoomLevel(zoomLevel float64) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetFloat64(zoomLevel)

	err := webViewSetZoomLevelFunction_Set()
	if err == nil {
		webViewSetZoomLevelFunction.Invoke(inArgs[:], nil)
	}

	return
}

var webViewStopLoadingFunction *gi.Function
var webViewStopLoadingFunction_Once sync.Once

func webViewStopLoadingFunction_Set() error {
	var err error
	webViewStopLoadingFunction_Once.Do(func() {
		err = webViewStruct_Set()
		if err != nil {
			return
		}
		webViewStopLoadingFunction, err = webViewStruct.InvokerNew("stop_loading")
	})
	return err
}

// StopLoading is a representation of the C type webkit_web_view_stop_loading.
func (recv *WebView) StopLoading() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := webViewStopLoadingFunction_Set()
	if err == nil {
		webViewStopLoadingFunction.Invoke(inArgs[:], nil)
	}

	return
}

var webViewTryCloseFunction *gi.Function
var webViewTryCloseFunction_Once sync.Once

func webViewTryCloseFunction_Set() error {
	var err error
	webViewTryCloseFunction_Once.Do(func() {
		err = webViewStruct_Set()
		if err != nil {
			return
		}
		webViewTryCloseFunction, err = webViewStruct.InvokerNew("try_close")
	})
	return err
}

// TryClose is a representation of the C type webkit_web_view_try_close.
func (recv *WebView) TryClose() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := webViewTryCloseFunction_Set()
	if err == nil {
		webViewTryCloseFunction.Invoke(inArgs[:], nil)
	}

	return
}

var webViewBaseStruct *gi.Struct
var webViewBaseStruct_Once sync.Once

func webViewBaseStruct_Set() error {
	var err error
	webViewBaseStruct_Once.Do(func() {
		webViewBaseStruct, err = gi.StructNew("WebKit2", "WebViewBase")
	})
	return err
}

type WebViewBase struct {
	native uintptr
}

// UNSUPPORTED : C value 'parentInstance' : for field getter : no Go type for 'Gtk.Container'

// UNSUPPORTED : C value 'parentInstance' : for field setter : no Go type for 'Gtk.Container'

// WebViewBaseStruct creates an uninitialised WebViewBase.
func WebViewBaseStruct() *WebViewBase {
	err := webViewBaseStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &WebViewBase{native: webViewBaseStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeWebViewBase)
	return structGo
}
func finalizeWebViewBase(obj *WebViewBase) {
	webViewBaseStruct.Free(obj.native)
}

var websiteDataManagerStruct *gi.Struct
var websiteDataManagerStruct_Once sync.Once

func websiteDataManagerStruct_Set() error {
	var err error
	websiteDataManagerStruct_Once.Do(func() {
		websiteDataManagerStruct, err = gi.StructNew("WebKit2", "WebsiteDataManager")
	})
	return err
}

type WebsiteDataManager struct {
	native uintptr
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'GObject.Object'

// FieldPriv returns the C field 'priv'.
func (recv *WebsiteDataManager) FieldPriv() *WebsiteDataManagerPrivate {
	argValue := gi.FieldGet(websiteDataManagerStruct, recv.native, "priv")
	value := &WebsiteDataManagerPrivate{native: argValue.Pointer()}
	return value
}

// SetFieldPriv sets the value of the C field 'priv'.
func (recv *WebsiteDataManager) SetFieldPriv(value *WebsiteDataManagerPrivate) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(websiteDataManagerStruct, recv.native, "priv", argValue)
}

// UNSUPPORTED : C value 'webkit_website_data_manager_new' : parameter '...' of type 'nil' not supported

// UNSUPPORTED : C value 'webkit_website_data_manager_new_ephemeral' : return type 'WebsiteDataManager' not supported

// UNSUPPORTED : C value 'webkit_website_data_manager_clear' : parameter 'types' of type 'WebsiteDataTypes' not supported

// UNSUPPORTED : C value 'webkit_website_data_manager_clear_finish' : parameter 'result' of type 'Gio.AsyncResult' not supported

// UNSUPPORTED : C value 'webkit_website_data_manager_fetch' : parameter 'types' of type 'WebsiteDataTypes' not supported

// UNSUPPORTED : C value 'webkit_website_data_manager_fetch_finish' : parameter 'result' of type 'Gio.AsyncResult' not supported

var websiteDataManagerGetBaseCacheDirectoryFunction *gi.Function
var websiteDataManagerGetBaseCacheDirectoryFunction_Once sync.Once

func websiteDataManagerGetBaseCacheDirectoryFunction_Set() error {
	var err error
	websiteDataManagerGetBaseCacheDirectoryFunction_Once.Do(func() {
		err = websiteDataManagerStruct_Set()
		if err != nil {
			return
		}
		websiteDataManagerGetBaseCacheDirectoryFunction, err = websiteDataManagerStruct.InvokerNew("get_base_cache_directory")
	})
	return err
}

// GetBaseCacheDirectory is a representation of the C type webkit_website_data_manager_get_base_cache_directory.
func (recv *WebsiteDataManager) GetBaseCacheDirectory() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := websiteDataManagerGetBaseCacheDirectoryFunction_Set()
	if err == nil {
		ret = websiteDataManagerGetBaseCacheDirectoryFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var websiteDataManagerGetBaseDataDirectoryFunction *gi.Function
var websiteDataManagerGetBaseDataDirectoryFunction_Once sync.Once

func websiteDataManagerGetBaseDataDirectoryFunction_Set() error {
	var err error
	websiteDataManagerGetBaseDataDirectoryFunction_Once.Do(func() {
		err = websiteDataManagerStruct_Set()
		if err != nil {
			return
		}
		websiteDataManagerGetBaseDataDirectoryFunction, err = websiteDataManagerStruct.InvokerNew("get_base_data_directory")
	})
	return err
}

// GetBaseDataDirectory is a representation of the C type webkit_website_data_manager_get_base_data_directory.
func (recv *WebsiteDataManager) GetBaseDataDirectory() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := websiteDataManagerGetBaseDataDirectoryFunction_Set()
	if err == nil {
		ret = websiteDataManagerGetBaseDataDirectoryFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'webkit_website_data_manager_get_cookie_manager' : return type 'CookieManager' not supported

var websiteDataManagerGetDiskCacheDirectoryFunction *gi.Function
var websiteDataManagerGetDiskCacheDirectoryFunction_Once sync.Once

func websiteDataManagerGetDiskCacheDirectoryFunction_Set() error {
	var err error
	websiteDataManagerGetDiskCacheDirectoryFunction_Once.Do(func() {
		err = websiteDataManagerStruct_Set()
		if err != nil {
			return
		}
		websiteDataManagerGetDiskCacheDirectoryFunction, err = websiteDataManagerStruct.InvokerNew("get_disk_cache_directory")
	})
	return err
}

// GetDiskCacheDirectory is a representation of the C type webkit_website_data_manager_get_disk_cache_directory.
func (recv *WebsiteDataManager) GetDiskCacheDirectory() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := websiteDataManagerGetDiskCacheDirectoryFunction_Set()
	if err == nil {
		ret = websiteDataManagerGetDiskCacheDirectoryFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var websiteDataManagerGetHstsCacheDirectoryFunction *gi.Function
var websiteDataManagerGetHstsCacheDirectoryFunction_Once sync.Once

func websiteDataManagerGetHstsCacheDirectoryFunction_Set() error {
	var err error
	websiteDataManagerGetHstsCacheDirectoryFunction_Once.Do(func() {
		err = websiteDataManagerStruct_Set()
		if err != nil {
			return
		}
		websiteDataManagerGetHstsCacheDirectoryFunction, err = websiteDataManagerStruct.InvokerNew("get_hsts_cache_directory")
	})
	return err
}

// GetHstsCacheDirectory is a representation of the C type webkit_website_data_manager_get_hsts_cache_directory.
func (recv *WebsiteDataManager) GetHstsCacheDirectory() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := websiteDataManagerGetHstsCacheDirectoryFunction_Set()
	if err == nil {
		ret = websiteDataManagerGetHstsCacheDirectoryFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var websiteDataManagerGetIndexeddbDirectoryFunction *gi.Function
var websiteDataManagerGetIndexeddbDirectoryFunction_Once sync.Once

func websiteDataManagerGetIndexeddbDirectoryFunction_Set() error {
	var err error
	websiteDataManagerGetIndexeddbDirectoryFunction_Once.Do(func() {
		err = websiteDataManagerStruct_Set()
		if err != nil {
			return
		}
		websiteDataManagerGetIndexeddbDirectoryFunction, err = websiteDataManagerStruct.InvokerNew("get_indexeddb_directory")
	})
	return err
}

// GetIndexeddbDirectory is a representation of the C type webkit_website_data_manager_get_indexeddb_directory.
func (recv *WebsiteDataManager) GetIndexeddbDirectory() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := websiteDataManagerGetIndexeddbDirectoryFunction_Set()
	if err == nil {
		ret = websiteDataManagerGetIndexeddbDirectoryFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var websiteDataManagerGetLocalStorageDirectoryFunction *gi.Function
var websiteDataManagerGetLocalStorageDirectoryFunction_Once sync.Once

func websiteDataManagerGetLocalStorageDirectoryFunction_Set() error {
	var err error
	websiteDataManagerGetLocalStorageDirectoryFunction_Once.Do(func() {
		err = websiteDataManagerStruct_Set()
		if err != nil {
			return
		}
		websiteDataManagerGetLocalStorageDirectoryFunction, err = websiteDataManagerStruct.InvokerNew("get_local_storage_directory")
	})
	return err
}

// GetLocalStorageDirectory is a representation of the C type webkit_website_data_manager_get_local_storage_directory.
func (recv *WebsiteDataManager) GetLocalStorageDirectory() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := websiteDataManagerGetLocalStorageDirectoryFunction_Set()
	if err == nil {
		ret = websiteDataManagerGetLocalStorageDirectoryFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var websiteDataManagerGetOfflineApplicationCacheDirectoryFunction *gi.Function
var websiteDataManagerGetOfflineApplicationCacheDirectoryFunction_Once sync.Once

func websiteDataManagerGetOfflineApplicationCacheDirectoryFunction_Set() error {
	var err error
	websiteDataManagerGetOfflineApplicationCacheDirectoryFunction_Once.Do(func() {
		err = websiteDataManagerStruct_Set()
		if err != nil {
			return
		}
		websiteDataManagerGetOfflineApplicationCacheDirectoryFunction, err = websiteDataManagerStruct.InvokerNew("get_offline_application_cache_directory")
	})
	return err
}

// GetOfflineApplicationCacheDirectory is a representation of the C type webkit_website_data_manager_get_offline_application_cache_directory.
func (recv *WebsiteDataManager) GetOfflineApplicationCacheDirectory() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := websiteDataManagerGetOfflineApplicationCacheDirectoryFunction_Set()
	if err == nil {
		ret = websiteDataManagerGetOfflineApplicationCacheDirectoryFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var websiteDataManagerGetWebsqlDirectoryFunction *gi.Function
var websiteDataManagerGetWebsqlDirectoryFunction_Once sync.Once

func websiteDataManagerGetWebsqlDirectoryFunction_Set() error {
	var err error
	websiteDataManagerGetWebsqlDirectoryFunction_Once.Do(func() {
		err = websiteDataManagerStruct_Set()
		if err != nil {
			return
		}
		websiteDataManagerGetWebsqlDirectoryFunction, err = websiteDataManagerStruct.InvokerNew("get_websql_directory")
	})
	return err
}

// GetWebsqlDirectory is a representation of the C type webkit_website_data_manager_get_websql_directory.
func (recv *WebsiteDataManager) GetWebsqlDirectory() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := websiteDataManagerGetWebsqlDirectoryFunction_Set()
	if err == nil {
		ret = websiteDataManagerGetWebsqlDirectoryFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var websiteDataManagerIsEphemeralFunction *gi.Function
var websiteDataManagerIsEphemeralFunction_Once sync.Once

func websiteDataManagerIsEphemeralFunction_Set() error {
	var err error
	websiteDataManagerIsEphemeralFunction_Once.Do(func() {
		err = websiteDataManagerStruct_Set()
		if err != nil {
			return
		}
		websiteDataManagerIsEphemeralFunction, err = websiteDataManagerStruct.InvokerNew("is_ephemeral")
	})
	return err
}

// IsEphemeral is a representation of the C type webkit_website_data_manager_is_ephemeral.
func (recv *WebsiteDataManager) IsEphemeral() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := websiteDataManagerIsEphemeralFunction_Set()
	if err == nil {
		ret = websiteDataManagerIsEphemeralFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'webkit_website_data_manager_remove' : parameter 'types' of type 'WebsiteDataTypes' not supported

// UNSUPPORTED : C value 'webkit_website_data_manager_remove_finish' : parameter 'result' of type 'Gio.AsyncResult' not supported

var windowPropertiesStruct *gi.Struct
var windowPropertiesStruct_Once sync.Once

func windowPropertiesStruct_Set() error {
	var err error
	windowPropertiesStruct_Once.Do(func() {
		windowPropertiesStruct, err = gi.StructNew("WebKit2", "WindowProperties")
	})
	return err
}

type WindowProperties struct {
	native uintptr
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'GObject.Object'

var windowPropertiesGetFullscreenFunction *gi.Function
var windowPropertiesGetFullscreenFunction_Once sync.Once

func windowPropertiesGetFullscreenFunction_Set() error {
	var err error
	windowPropertiesGetFullscreenFunction_Once.Do(func() {
		err = windowPropertiesStruct_Set()
		if err != nil {
			return
		}
		windowPropertiesGetFullscreenFunction, err = windowPropertiesStruct.InvokerNew("get_fullscreen")
	})
	return err
}

// GetFullscreen is a representation of the C type webkit_window_properties_get_fullscreen.
func (recv *WindowProperties) GetFullscreen() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := windowPropertiesGetFullscreenFunction_Set()
	if err == nil {
		ret = windowPropertiesGetFullscreenFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'webkit_window_properties_get_geometry' : parameter 'geometry' of type 'Gdk.Rectangle' not supported

var windowPropertiesGetLocationbarVisibleFunction *gi.Function
var windowPropertiesGetLocationbarVisibleFunction_Once sync.Once

func windowPropertiesGetLocationbarVisibleFunction_Set() error {
	var err error
	windowPropertiesGetLocationbarVisibleFunction_Once.Do(func() {
		err = windowPropertiesStruct_Set()
		if err != nil {
			return
		}
		windowPropertiesGetLocationbarVisibleFunction, err = windowPropertiesStruct.InvokerNew("get_locationbar_visible")
	})
	return err
}

// GetLocationbarVisible is a representation of the C type webkit_window_properties_get_locationbar_visible.
func (recv *WindowProperties) GetLocationbarVisible() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := windowPropertiesGetLocationbarVisibleFunction_Set()
	if err == nil {
		ret = windowPropertiesGetLocationbarVisibleFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var windowPropertiesGetMenubarVisibleFunction *gi.Function
var windowPropertiesGetMenubarVisibleFunction_Once sync.Once

func windowPropertiesGetMenubarVisibleFunction_Set() error {
	var err error
	windowPropertiesGetMenubarVisibleFunction_Once.Do(func() {
		err = windowPropertiesStruct_Set()
		if err != nil {
			return
		}
		windowPropertiesGetMenubarVisibleFunction, err = windowPropertiesStruct.InvokerNew("get_menubar_visible")
	})
	return err
}

// GetMenubarVisible is a representation of the C type webkit_window_properties_get_menubar_visible.
func (recv *WindowProperties) GetMenubarVisible() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := windowPropertiesGetMenubarVisibleFunction_Set()
	if err == nil {
		ret = windowPropertiesGetMenubarVisibleFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var windowPropertiesGetResizableFunction *gi.Function
var windowPropertiesGetResizableFunction_Once sync.Once

func windowPropertiesGetResizableFunction_Set() error {
	var err error
	windowPropertiesGetResizableFunction_Once.Do(func() {
		err = windowPropertiesStruct_Set()
		if err != nil {
			return
		}
		windowPropertiesGetResizableFunction, err = windowPropertiesStruct.InvokerNew("get_resizable")
	})
	return err
}

// GetResizable is a representation of the C type webkit_window_properties_get_resizable.
func (recv *WindowProperties) GetResizable() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := windowPropertiesGetResizableFunction_Set()
	if err == nil {
		ret = windowPropertiesGetResizableFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var windowPropertiesGetScrollbarsVisibleFunction *gi.Function
var windowPropertiesGetScrollbarsVisibleFunction_Once sync.Once

func windowPropertiesGetScrollbarsVisibleFunction_Set() error {
	var err error
	windowPropertiesGetScrollbarsVisibleFunction_Once.Do(func() {
		err = windowPropertiesStruct_Set()
		if err != nil {
			return
		}
		windowPropertiesGetScrollbarsVisibleFunction, err = windowPropertiesStruct.InvokerNew("get_scrollbars_visible")
	})
	return err
}

// GetScrollbarsVisible is a representation of the C type webkit_window_properties_get_scrollbars_visible.
func (recv *WindowProperties) GetScrollbarsVisible() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := windowPropertiesGetScrollbarsVisibleFunction_Set()
	if err == nil {
		ret = windowPropertiesGetScrollbarsVisibleFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var windowPropertiesGetStatusbarVisibleFunction *gi.Function
var windowPropertiesGetStatusbarVisibleFunction_Once sync.Once

func windowPropertiesGetStatusbarVisibleFunction_Set() error {
	var err error
	windowPropertiesGetStatusbarVisibleFunction_Once.Do(func() {
		err = windowPropertiesStruct_Set()
		if err != nil {
			return
		}
		windowPropertiesGetStatusbarVisibleFunction, err = windowPropertiesStruct.InvokerNew("get_statusbar_visible")
	})
	return err
}

// GetStatusbarVisible is a representation of the C type webkit_window_properties_get_statusbar_visible.
func (recv *WindowProperties) GetStatusbarVisible() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := windowPropertiesGetStatusbarVisibleFunction_Set()
	if err == nil {
		ret = windowPropertiesGetStatusbarVisibleFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var windowPropertiesGetToolbarVisibleFunction *gi.Function
var windowPropertiesGetToolbarVisibleFunction_Once sync.Once

func windowPropertiesGetToolbarVisibleFunction_Set() error {
	var err error
	windowPropertiesGetToolbarVisibleFunction_Once.Do(func() {
		err = windowPropertiesStruct_Set()
		if err != nil {
			return
		}
		windowPropertiesGetToolbarVisibleFunction, err = windowPropertiesStruct.InvokerNew("get_toolbar_visible")
	})
	return err
}

// GetToolbarVisible is a representation of the C type webkit_window_properties_get_toolbar_visible.
func (recv *WindowProperties) GetToolbarVisible() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := windowPropertiesGetToolbarVisibleFunction_Set()
	if err == nil {
		ret = windowPropertiesGetToolbarVisibleFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// WindowPropertiesStruct creates an uninitialised WindowProperties.
func WindowPropertiesStruct() *WindowProperties {
	err := windowPropertiesStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &WindowProperties{native: windowPropertiesStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeWindowProperties)
	return structGo
}
func finalizeWindowProperties(obj *WindowProperties) {
	windowPropertiesStruct.Free(obj.native)
}
