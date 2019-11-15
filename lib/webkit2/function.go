// Code generated - DO NOT EDIT.

package webkit2

import gi "github.com/pekim/gobbi/internal/gi"

// UNSUPPORTED : C value 'webkit_download_error_quark' : return type not supported

// UNSUPPORTED : C value 'webkit_favicon_database_error_quark' : return type not supported

var getMajorVersionInvoker *gi.Function

// GetMajorVersion is a representation of the C type webkit_get_major_version.
func GetMajorVersion() uint32 {
	if getMajorVersionInvoker == nil {
		getMajorVersionInvoker = gi.FunctionInvokerNew("WebKit2", "get_major_version")
	}

	ret := getMajorVersionInvoker.Invoke()
	return ret.Uint32()
}

var getMicroVersionInvoker *gi.Function

// GetMicroVersion is a representation of the C type webkit_get_micro_version.
func GetMicroVersion() uint32 {
	if getMicroVersionInvoker == nil {
		getMicroVersionInvoker = gi.FunctionInvokerNew("WebKit2", "get_micro_version")
	}

	ret := getMicroVersionInvoker.Invoke()
	return ret.Uint32()
}

var getMinorVersionInvoker *gi.Function

// GetMinorVersion is a representation of the C type webkit_get_minor_version.
func GetMinorVersion() uint32 {
	if getMinorVersionInvoker == nil {
		getMinorVersionInvoker = gi.FunctionInvokerNew("WebKit2", "get_minor_version")
	}

	ret := getMinorVersionInvoker.Invoke()
	return ret.Uint32()
}

// UNSUPPORTED : C value 'webkit_javascript_error_quark' : return type not supported

// UNSUPPORTED : C value 'webkit_network_error_quark' : return type not supported

// UNSUPPORTED : C value 'webkit_plugin_error_quark' : return type not supported

// UNSUPPORTED : C value 'webkit_policy_error_quark' : return type not supported

// UNSUPPORTED : C value 'webkit_print_error_quark' : return type not supported

// UNSUPPORTED : C value 'webkit_snapshot_error_quark' : return type not supported

// UNSUPPORTED : C value 'webkit_uri_for_display' : has parameters

// UNSUPPORTED : C value 'webkit_user_content_filter_error_quark' : return type not supported

// UNSUPPORTED : C value 'webkit_user_media_permission_is_for_audio_device' : has parameters

// UNSUPPORTED : C value 'webkit_user_media_permission_is_for_video_device' : has parameters
