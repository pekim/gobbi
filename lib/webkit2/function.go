// Code generated - DO NOT EDIT.

package webkit2

import (
	gi "github.com/pekim/gobbi/internal/gi"
	"sync"
)

// UNSUPPORTED : C value 'webkit_download_error_quark' : return type 'GLib.Quark' not supported

// UNSUPPORTED : C value 'webkit_favicon_database_error_quark' : return type 'GLib.Quark' not supported

var getMajorVersionFunction *gi.Function
var getMajorVersionFunction_Once sync.Once

func getMajorVersionFunction_Set() {
	getMajorVersionFunction_Once.Do(func() {
		getMajorVersionFunction = gi.FunctionInvokerNew("WebKit2", "get_major_version")
	})
}

// GetMajorVersion is a representation of the C type webkit_get_major_version.
func GetMajorVersion() uint32 {

	var ret gi.Argument

	getMajorVersionFunction_Set()

	ret = getMajorVersionFunction.Invoke(nil, nil)

	retGo := ret.Uint32()

	return retGo
}

var getMicroVersionFunction *gi.Function
var getMicroVersionFunction_Once sync.Once

func getMicroVersionFunction_Set() {
	getMicroVersionFunction_Once.Do(func() {
		getMicroVersionFunction = gi.FunctionInvokerNew("WebKit2", "get_micro_version")
	})
}

// GetMicroVersion is a representation of the C type webkit_get_micro_version.
func GetMicroVersion() uint32 {

	var ret gi.Argument

	getMicroVersionFunction_Set()

	ret = getMicroVersionFunction.Invoke(nil, nil)

	retGo := ret.Uint32()

	return retGo
}

var getMinorVersionFunction *gi.Function
var getMinorVersionFunction_Once sync.Once

func getMinorVersionFunction_Set() {
	getMinorVersionFunction_Once.Do(func() {
		getMinorVersionFunction = gi.FunctionInvokerNew("WebKit2", "get_minor_version")
	})
}

// GetMinorVersion is a representation of the C type webkit_get_minor_version.
func GetMinorVersion() uint32 {

	var ret gi.Argument

	getMinorVersionFunction_Set()

	ret = getMinorVersionFunction.Invoke(nil, nil)

	retGo := ret.Uint32()

	return retGo
}

// UNSUPPORTED : C value 'webkit_javascript_error_quark' : return type 'GLib.Quark' not supported

// UNSUPPORTED : C value 'webkit_network_error_quark' : return type 'GLib.Quark' not supported

// UNSUPPORTED : C value 'webkit_plugin_error_quark' : return type 'GLib.Quark' not supported

// UNSUPPORTED : C value 'webkit_policy_error_quark' : return type 'GLib.Quark' not supported

// UNSUPPORTED : C value 'webkit_print_error_quark' : return type 'GLib.Quark' not supported

// UNSUPPORTED : C value 'webkit_snapshot_error_quark' : return type 'GLib.Quark' not supported

var uriForDisplayFunction *gi.Function
var uriForDisplayFunction_Once sync.Once

func uriForDisplayFunction_Set() {
	uriForDisplayFunction_Once.Do(func() {
		uriForDisplayFunction = gi.FunctionInvokerNew("WebKit2", "uri_for_display")
	})
}

// UriForDisplay is a representation of the C type webkit_uri_for_display.
func UriForDisplay(uri string) string {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(uri)

	var ret gi.Argument

	uriForDisplayFunction_Set()

	ret = uriForDisplayFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

// UNSUPPORTED : C value 'webkit_user_content_filter_error_quark' : return type 'GLib.Quark' not supported

// UNSUPPORTED : C value 'webkit_user_media_permission_is_for_audio_device' : parameter 'request' of type 'UserMediaPermissionRequest' not supported

// UNSUPPORTED : C value 'webkit_user_media_permission_is_for_video_device' : parameter 'request' of type 'UserMediaPermissionRequest' not supported
