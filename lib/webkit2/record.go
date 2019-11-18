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

	return &ApplicationInfo{native: ret.Pointer()}
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

type MimeInfo struct {
	native uintptr
}

type NavigationAction struct {
	native uintptr
}

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

	return &SecurityOrigin{native: ret.Pointer()}
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

	return &SecurityOrigin{native: ret.Pointer()}
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

type UserStyleSheet struct {
	native uintptr
}

// UNSUPPORTED : C value 'webkit_user_style_sheet_new' : parameter 'injected_frames' of type 'UserContentInjectedFrames' not supported

// UNSUPPORTED : C value 'webkit_user_style_sheet_new_for_world' : parameter 'injected_frames' of type 'UserContentInjectedFrames' not supported

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

type WebsiteData struct {
	native uintptr
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
