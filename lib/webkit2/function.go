// Code generated - DO NOT EDIT.

package webkit2

import gi "github.com/pekim/gobbi/internal/gi"

// UNSUPPORTED : C value 'webkit_download_error_quark' : return type 'GLib.Quark' not supported

// UNSUPPORTED : C value 'webkit_favicon_database_error_quark' : return type 'GLib.Quark' not supported

var getMajorVersionInvoker *gi.Function

// GetMajorVersion is a representation of the C type webkit_get_major_version.
func GetMajorVersion() uint32 {
	if getMajorVersionInvoker == nil {
		getMajorVersionInvoker = gi.FunctionInvokerNew("WebKit2", "get_major_version")
	}

	ret := getMajorVersionInvoker.Invoke(nil, nil)

	retGo := ret.Uint32()

	return retGo
}

var getMicroVersionInvoker *gi.Function

// GetMicroVersion is a representation of the C type webkit_get_micro_version.
func GetMicroVersion() uint32 {
	if getMicroVersionInvoker == nil {
		getMicroVersionInvoker = gi.FunctionInvokerNew("WebKit2", "get_micro_version")
	}

	ret := getMicroVersionInvoker.Invoke(nil, nil)

	retGo := ret.Uint32()

	return retGo
}

var getMinorVersionInvoker *gi.Function

// GetMinorVersion is a representation of the C type webkit_get_minor_version.
func GetMinorVersion() uint32 {
	if getMinorVersionInvoker == nil {
		getMinorVersionInvoker = gi.FunctionInvokerNew("WebKit2", "get_minor_version")
	}

	ret := getMinorVersionInvoker.Invoke(nil, nil)

	retGo := ret.Uint32()

	return retGo
}

// UNSUPPORTED : C value 'webkit_javascript_error_quark' : return type 'GLib.Quark' not supported

// UNSUPPORTED : C value 'webkit_network_error_quark' : return type 'GLib.Quark' not supported

// UNSUPPORTED : C value 'webkit_plugin_error_quark' : return type 'GLib.Quark' not supported

// UNSUPPORTED : C value 'webkit_policy_error_quark' : return type 'GLib.Quark' not supported

// UNSUPPORTED : C value 'webkit_print_error_quark' : return type 'GLib.Quark' not supported

// UNSUPPORTED : C value 'webkit_snapshot_error_quark' : return type 'GLib.Quark' not supported

var uriForDisplayInvoker *gi.Function

// UriForDisplay is a representation of the C type webkit_uri_for_display.
func UriForDisplay(uri string) string {
	if uriForDisplayInvoker == nil {
		uriForDisplayInvoker = gi.FunctionInvokerNew("WebKit2", "uri_for_display")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetString(uri)

	ret := uriForDisplayInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

// UNSUPPORTED : C value 'webkit_user_content_filter_error_quark' : return type 'GLib.Quark' not supported

// UNSUPPORTED : C value 'webkit_user_media_permission_is_for_audio_device' : parameter 'request' of type 'UserMediaPermissionRequest' not supported

// UNSUPPORTED : C value 'webkit_user_media_permission_is_for_video_device' : parameter 'request' of type 'UserMediaPermissionRequest' not supported
