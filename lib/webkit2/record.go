// Code generated - DO NOT EDIT.

package webkit2

import gi "github.com/pekim/gobbi/internal/gi"

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

type AuthenticationRequestClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value '_webkit_reserved0' : missing Type

	// UNSUPPORTED : C value '_webkit_reserved1' : missing Type

	// UNSUPPORTED : C value '_webkit_reserved2' : missing Type

	// UNSUPPORTED : C value '_webkit_reserved3' : missing Type

}

type AuthenticationRequestPrivate struct {
	native uintptr
}

type AutomationSessionClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value '_webkit_reserved0' : missing Type

	// UNSUPPORTED : C value '_webkit_reserved1' : missing Type

	// UNSUPPORTED : C value '_webkit_reserved2' : missing Type

	// UNSUPPORTED : C value '_webkit_reserved3' : missing Type

}

type AutomationSessionPrivate struct {
	native uintptr
}

type BackForwardListClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value '_webkit_reserved0' : missing Type

	// UNSUPPORTED : C value '_webkit_reserved1' : missing Type

	// UNSUPPORTED : C value '_webkit_reserved2' : missing Type

	// UNSUPPORTED : C value '_webkit_reserved3' : missing Type

}

type BackForwardListItemClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.InitiallyUnownedClass'

	// UNSUPPORTED : C value '_webkit_reserved0' : missing Type

	// UNSUPPORTED : C value '_webkit_reserved1' : missing Type

	// UNSUPPORTED : C value '_webkit_reserved2' : missing Type

	// UNSUPPORTED : C value '_webkit_reserved3' : missing Type

}

type BackForwardListItemPrivate struct {
	native uintptr
}

type BackForwardListPrivate struct {
	native uintptr
}

type ColorChooserRequestClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

}

type ColorChooserRequestPrivate struct {
	native uintptr
}

type ContextMenuClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value '_webkit_reserved0' : missing Type

	// UNSUPPORTED : C value '_webkit_reserved1' : missing Type

	// UNSUPPORTED : C value '_webkit_reserved2' : missing Type

	// UNSUPPORTED : C value '_webkit_reserved3' : missing Type

}

type ContextMenuItemClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.InitiallyUnownedClass'

	// UNSUPPORTED : C value '_webkit_reserved0' : missing Type

	// UNSUPPORTED : C value '_webkit_reserved1' : missing Type

	// UNSUPPORTED : C value '_webkit_reserved2' : missing Type

	// UNSUPPORTED : C value '_webkit_reserved3' : missing Type

}

type ContextMenuItemPrivate struct {
	native uintptr
}

type ContextMenuPrivate struct {
	native uintptr
}

type CookieManagerClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value '_webkit_reserved0' : missing Type

	// UNSUPPORTED : C value '_webkit_reserved1' : missing Type

	// UNSUPPORTED : C value '_webkit_reserved2' : missing Type

	// UNSUPPORTED : C value '_webkit_reserved3' : missing Type

}

type CookieManagerPrivate struct {
	native uintptr
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

type DeviceInfoPermissionRequestClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value '_webkit_reserved0' : missing Type

	// UNSUPPORTED : C value '_webkit_reserved1' : missing Type

	// UNSUPPORTED : C value '_webkit_reserved2' : missing Type

	// UNSUPPORTED : C value '_webkit_reserved3' : missing Type

}

type DeviceInfoPermissionRequestPrivate struct {
	native uintptr
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

type DownloadPrivate struct {
	native uintptr
}

type EditorStateClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value '_webkit_reserved0' : missing Type

	// UNSUPPORTED : C value '_webkit_reserved1' : missing Type

	// UNSUPPORTED : C value '_webkit_reserved2' : missing Type

	// UNSUPPORTED : C value '_webkit_reserved3' : missing Type

}

type EditorStatePrivate struct {
	native uintptr
}

type FaviconDatabaseClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value '_webkit_reserved0' : missing Type

	// UNSUPPORTED : C value '_webkit_reserved1' : missing Type

	// UNSUPPORTED : C value '_webkit_reserved2' : missing Type

	// UNSUPPORTED : C value '_webkit_reserved3' : missing Type

}

type FaviconDatabasePrivate struct {
	native uintptr
}

type FileChooserRequestClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value '_webkit_reserved0' : missing Type

	// UNSUPPORTED : C value '_webkit_reserved1' : missing Type

	// UNSUPPORTED : C value '_webkit_reserved2' : missing Type

	// UNSUPPORTED : C value '_webkit_reserved3' : missing Type

}

type FileChooserRequestPrivate struct {
	native uintptr
}

type FindControllerClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value '_webkit_reserved0' : missing Type

	// UNSUPPORTED : C value '_webkit_reserved1' : missing Type

	// UNSUPPORTED : C value '_webkit_reserved2' : missing Type

	// UNSUPPORTED : C value '_webkit_reserved3' : missing Type

}

type FindControllerPrivate struct {
	native uintptr
}

type FormSubmissionRequestClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value '_webkit_reserved0' : missing Type

	// UNSUPPORTED : C value '_webkit_reserved1' : missing Type

	// UNSUPPORTED : C value '_webkit_reserved2' : missing Type

	// UNSUPPORTED : C value '_webkit_reserved3' : missing Type

}

type FormSubmissionRequestPrivate struct {
	native uintptr
}

type GeolocationManagerClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value '_webkit_reserved0' : missing Type

	// UNSUPPORTED : C value '_webkit_reserved1' : missing Type

	// UNSUPPORTED : C value '_webkit_reserved2' : missing Type

	// UNSUPPORTED : C value '_webkit_reserved3' : missing Type

}

type GeolocationManagerPrivate struct {
	native uintptr
}

type GeolocationPermissionRequestClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value '_webkit_reserved0' : missing Type

	// UNSUPPORTED : C value '_webkit_reserved1' : missing Type

	// UNSUPPORTED : C value '_webkit_reserved2' : missing Type

	// UNSUPPORTED : C value '_webkit_reserved3' : missing Type

}

type GeolocationPermissionRequestPrivate struct {
	native uintptr
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

type HitTestResultClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value '_webkit_reserved0' : missing Type

	// UNSUPPORTED : C value '_webkit_reserved1' : missing Type

	// UNSUPPORTED : C value '_webkit_reserved2' : missing Type

	// UNSUPPORTED : C value '_webkit_reserved3' : missing Type

}

type HitTestResultPrivate struct {
	native uintptr
}

type InstallMissingMediaPluginsPermissionRequestClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value '_webkit_reserved0' : missing Type

	// UNSUPPORTED : C value '_webkit_reserved1' : missing Type

	// UNSUPPORTED : C value '_webkit_reserved2' : missing Type

	// UNSUPPORTED : C value '_webkit_reserved3' : missing Type

}

type InstallMissingMediaPluginsPermissionRequestPrivate struct {
	native uintptr
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

type NavigationPolicyDecisionClass struct {
	native      uintptr
	ParentClass *PolicyDecisionClass
	// UNSUPPORTED : C value '_webkit_reserved0' : missing Type

	// UNSUPPORTED : C value '_webkit_reserved1' : missing Type

	// UNSUPPORTED : C value '_webkit_reserved2' : missing Type

	// UNSUPPORTED : C value '_webkit_reserved3' : missing Type

}

type NavigationPolicyDecisionPrivate struct {
	native uintptr
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

type NotificationPermissionRequestClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

}

type NotificationPermissionRequestPrivate struct {
	native uintptr
}

type NotificationPrivate struct {
	native uintptr
}

type OptionMenuClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value '_webkit_reserved0' : missing Type

	// UNSUPPORTED : C value '_webkit_reserved1' : missing Type

	// UNSUPPORTED : C value '_webkit_reserved2' : missing Type

	// UNSUPPORTED : C value '_webkit_reserved3' : missing Type

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

type OptionMenuPrivate struct {
	native uintptr
}

type PermissionRequestIface struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_interface' : no Go type for 'GObject.TypeInterface'

	// UNSUPPORTED : C value 'allow' : missing Type

	// UNSUPPORTED : C value 'deny' : missing Type

}

type PluginClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value '_webkit_reserved0' : missing Type

	// UNSUPPORTED : C value '_webkit_reserved1' : missing Type

	// UNSUPPORTED : C value '_webkit_reserved2' : missing Type

	// UNSUPPORTED : C value '_webkit_reserved3' : missing Type

}

type PluginPrivate struct {
	native uintptr
}

type PolicyDecisionClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value '_webkit_reserved0' : missing Type

	// UNSUPPORTED : C value '_webkit_reserved1' : missing Type

	// UNSUPPORTED : C value '_webkit_reserved2' : missing Type

	// UNSUPPORTED : C value '_webkit_reserved3' : missing Type

}

type PolicyDecisionPrivate struct {
	native uintptr
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

type PrintCustomWidgetPrivate struct {
	native uintptr
}

type PrintOperationClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value '_webkit_reserved0' : missing Type

	// UNSUPPORTED : C value '_webkit_reserved1' : missing Type

	// UNSUPPORTED : C value '_webkit_reserved2' : missing Type

	// UNSUPPORTED : C value '_webkit_reserved3' : missing Type

}

type PrintOperationPrivate struct {
	native uintptr
}

type ResponsePolicyDecisionClass struct {
	native      uintptr
	ParentClass *PolicyDecisionClass
	// UNSUPPORTED : C value '_webkit_reserved0' : missing Type

	// UNSUPPORTED : C value '_webkit_reserved1' : missing Type

	// UNSUPPORTED : C value '_webkit_reserved2' : missing Type

	// UNSUPPORTED : C value '_webkit_reserved3' : missing Type

}

type ResponsePolicyDecisionPrivate struct {
	native uintptr
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

type SecurityManagerClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value '_webkit_reserved0' : missing Type

	// UNSUPPORTED : C value '_webkit_reserved1' : missing Type

	// UNSUPPORTED : C value '_webkit_reserved2' : missing Type

	// UNSUPPORTED : C value '_webkit_reserved3' : missing Type

}

type SecurityManagerPrivate struct {
	native uintptr
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

type SettingsClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value '_webkit_reserved0' : missing Type

	// UNSUPPORTED : C value '_webkit_reserved1' : missing Type

	// UNSUPPORTED : C value '_webkit_reserved2' : missing Type

	// UNSUPPORTED : C value '_webkit_reserved3' : missing Type

}

type SettingsPrivate struct {
	native uintptr
}

type URIRequestClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value '_webkit_reserved0' : missing Type

	// UNSUPPORTED : C value '_webkit_reserved1' : missing Type

	// UNSUPPORTED : C value '_webkit_reserved2' : missing Type

	// UNSUPPORTED : C value '_webkit_reserved3' : missing Type

}

type URIRequestPrivate struct {
	native uintptr
}

type URIResponseClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value '_webkit_reserved0' : missing Type

	// UNSUPPORTED : C value '_webkit_reserved1' : missing Type

	// UNSUPPORTED : C value '_webkit_reserved2' : missing Type

	// UNSUPPORTED : C value '_webkit_reserved3' : missing Type

}

type URIResponsePrivate struct {
	native uintptr
}

type URISchemeRequestClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value '_webkit_reserved0' : missing Type

	// UNSUPPORTED : C value '_webkit_reserved1' : missing Type

	// UNSUPPORTED : C value '_webkit_reserved2' : missing Type

	// UNSUPPORTED : C value '_webkit_reserved3' : missing Type

}

type URISchemeRequestPrivate struct {
	native uintptr
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

type UserContentFilterStoreClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value '_webkit_reserved0' : missing Type

	// UNSUPPORTED : C value '_webkit_reserved1' : missing Type

	// UNSUPPORTED : C value '_webkit_reserved2' : missing Type

	// UNSUPPORTED : C value '_webkit_reserved3' : missing Type

}

type UserContentFilterStorePrivate struct {
	native uintptr
}

type UserContentManagerClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value '_webkit_reserved0' : missing Type

	// UNSUPPORTED : C value '_webkit_reserved1' : missing Type

	// UNSUPPORTED : C value '_webkit_reserved2' : missing Type

	// UNSUPPORTED : C value '_webkit_reserved3' : missing Type

}

type UserContentManagerPrivate struct {
	native uintptr
}

type UserMediaPermissionRequestClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value '_webkit_reserved0' : missing Type

	// UNSUPPORTED : C value '_webkit_reserved1' : missing Type

	// UNSUPPORTED : C value '_webkit_reserved2' : missing Type

	// UNSUPPORTED : C value '_webkit_reserved3' : missing Type

}

type UserMediaPermissionRequestPrivate struct {
	native uintptr
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

type WebContextPrivate struct {
	native uintptr
}

type WebInspectorClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value '_webkit_reserved0' : missing Type

	// UNSUPPORTED : C value '_webkit_reserved1' : missing Type

	// UNSUPPORTED : C value '_webkit_reserved2' : missing Type

	// UNSUPPORTED : C value '_webkit_reserved3' : missing Type

}

type WebInspectorPrivate struct {
	native uintptr
}

type WebResourceClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value '_webkit_reserved0' : missing Type

	// UNSUPPORTED : C value '_webkit_reserved1' : missing Type

	// UNSUPPORTED : C value '_webkit_reserved2' : missing Type

	// UNSUPPORTED : C value '_webkit_reserved3' : missing Type

}

type WebResourcePrivate struct {
	native uintptr
}

type WebViewBaseClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parentClass' : no Go type for 'Gtk.ContainerClass'

	// UNSUPPORTED : C value '_webkit_reserved0' : missing Type

	// UNSUPPORTED : C value '_webkit_reserved1' : missing Type

	// UNSUPPORTED : C value '_webkit_reserved2' : missing Type

	// UNSUPPORTED : C value '_webkit_reserved3' : missing Type

}

type WebViewBasePrivate struct {
	native uintptr
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

type WebViewPrivate struct {
	native uintptr
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

type WebsiteDataManagerClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value '_webkit_reserved0' : missing Type

	// UNSUPPORTED : C value '_webkit_reserved1' : missing Type

	// UNSUPPORTED : C value '_webkit_reserved2' : missing Type

	// UNSUPPORTED : C value '_webkit_reserved3' : missing Type

}

type WebsiteDataManagerPrivate struct {
	native uintptr
}

type WindowPropertiesClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value '_webkit_reserved0' : missing Type

	// UNSUPPORTED : C value '_webkit_reserved1' : missing Type

	// UNSUPPORTED : C value '_webkit_reserved2' : missing Type

	// UNSUPPORTED : C value '_webkit_reserved3' : missing Type

}

type WindowPropertiesPrivate struct {
	native uintptr
}
