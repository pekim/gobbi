// Code generated - DO NOT EDIT.

package webkit2

import gi "github.com/pekim/gobbi/internal/gi"

var downloadErrorQuarkInvoker *gi.Function

// DownloadErrorQuark is a representation of the C type webkit_download_error_quark.
func DownloadErrorQuark() {
	if downloadErrorQuarkInvoker == nil {
		downloadErrorQuarkInvoker = gi.FunctionInvokerNew("WebKit2", "download_error_quark")
	}

	downloadErrorQuarkInvoker.Invoke()
}

var faviconDatabaseErrorQuarkInvoker *gi.Function

// FaviconDatabaseErrorQuark is a representation of the C type webkit_favicon_database_error_quark.
func FaviconDatabaseErrorQuark() {
	if faviconDatabaseErrorQuarkInvoker == nil {
		faviconDatabaseErrorQuarkInvoker = gi.FunctionInvokerNew("WebKit2", "favicon_database_error_quark")
	}

	faviconDatabaseErrorQuarkInvoker.Invoke()
}

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

var javascriptErrorQuarkInvoker *gi.Function

// JavascriptErrorQuark is a representation of the C type webkit_javascript_error_quark.
func JavascriptErrorQuark() {
	if javascriptErrorQuarkInvoker == nil {
		javascriptErrorQuarkInvoker = gi.FunctionInvokerNew("WebKit2", "javascript_error_quark")
	}

	javascriptErrorQuarkInvoker.Invoke()
}

var networkErrorQuarkInvoker *gi.Function

// NetworkErrorQuark is a representation of the C type webkit_network_error_quark.
func NetworkErrorQuark() {
	if networkErrorQuarkInvoker == nil {
		networkErrorQuarkInvoker = gi.FunctionInvokerNew("WebKit2", "network_error_quark")
	}

	networkErrorQuarkInvoker.Invoke()
}

var pluginErrorQuarkInvoker *gi.Function

// PluginErrorQuark is a representation of the C type webkit_plugin_error_quark.
func PluginErrorQuark() {
	if pluginErrorQuarkInvoker == nil {
		pluginErrorQuarkInvoker = gi.FunctionInvokerNew("WebKit2", "plugin_error_quark")
	}

	pluginErrorQuarkInvoker.Invoke()
}

var policyErrorQuarkInvoker *gi.Function

// PolicyErrorQuark is a representation of the C type webkit_policy_error_quark.
func PolicyErrorQuark() {
	if policyErrorQuarkInvoker == nil {
		policyErrorQuarkInvoker = gi.FunctionInvokerNew("WebKit2", "policy_error_quark")
	}

	policyErrorQuarkInvoker.Invoke()
}

var printErrorQuarkInvoker *gi.Function

// PrintErrorQuark is a representation of the C type webkit_print_error_quark.
func PrintErrorQuark() {
	if printErrorQuarkInvoker == nil {
		printErrorQuarkInvoker = gi.FunctionInvokerNew("WebKit2", "print_error_quark")
	}

	printErrorQuarkInvoker.Invoke()
}

var snapshotErrorQuarkInvoker *gi.Function

// SnapshotErrorQuark is a representation of the C type webkit_snapshot_error_quark.
func SnapshotErrorQuark() {
	if snapshotErrorQuarkInvoker == nil {
		snapshotErrorQuarkInvoker = gi.FunctionInvokerNew("WebKit2", "snapshot_error_quark")
	}

	snapshotErrorQuarkInvoker.Invoke()
}

// UNSUPPORTED : C value 'webkit_uri_for_display' : non trivial function

var userContentFilterErrorQuarkInvoker *gi.Function

// UserContentFilterErrorQuark is a representation of the C type webkit_user_content_filter_error_quark.
func UserContentFilterErrorQuark() {
	if userContentFilterErrorQuarkInvoker == nil {
		userContentFilterErrorQuarkInvoker = gi.FunctionInvokerNew("WebKit2", "user_content_filter_error_quark")
	}

	userContentFilterErrorQuarkInvoker.Invoke()
}

// UNSUPPORTED : C value 'webkit_user_media_permission_is_for_audio_device' : non trivial function

// UNSUPPORTED : C value 'webkit_user_media_permission_is_for_video_device' : non trivial function
