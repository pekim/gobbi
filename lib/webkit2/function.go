// Code generated - DO NOT EDIT.

package webkit2

import (
	gi "github.com/pekim/gobbi/internal/gi"
	glib "github.com/pekim/gobbi/lib/glib"
	"sync"
)

var downloadErrorQuarkFunction *gi.Function
var downloadErrorQuarkFunction_Once sync.Once

func downloadErrorQuarkFunction_Set() error {
	var err error
	downloadErrorQuarkFunction_Once.Do(func() {
		downloadErrorQuarkFunction, err = gi.FunctionInvokerNew("WebKit2", "download_error_quark")
	})
	return err
}

// DownloadErrorQuark is a representation of the C type webkit_download_error_quark.
func DownloadErrorQuark() glib.Quark {

	var ret gi.Argument

	err := downloadErrorQuarkFunction_Set()
	if err == nil {
		ret = downloadErrorQuarkFunction.Invoke(nil, nil)
	}

	retGo := glib.Quark(ret.Uint32())

	return retGo
}

var faviconDatabaseErrorQuarkFunction *gi.Function
var faviconDatabaseErrorQuarkFunction_Once sync.Once

func faviconDatabaseErrorQuarkFunction_Set() error {
	var err error
	faviconDatabaseErrorQuarkFunction_Once.Do(func() {
		faviconDatabaseErrorQuarkFunction, err = gi.FunctionInvokerNew("WebKit2", "favicon_database_error_quark")
	})
	return err
}

// FaviconDatabaseErrorQuark is a representation of the C type webkit_favicon_database_error_quark.
func FaviconDatabaseErrorQuark() glib.Quark {

	var ret gi.Argument

	err := faviconDatabaseErrorQuarkFunction_Set()
	if err == nil {
		ret = faviconDatabaseErrorQuarkFunction.Invoke(nil, nil)
	}

	retGo := glib.Quark(ret.Uint32())

	return retGo
}

var getMajorVersionFunction *gi.Function
var getMajorVersionFunction_Once sync.Once

func getMajorVersionFunction_Set() error {
	var err error
	getMajorVersionFunction_Once.Do(func() {
		getMajorVersionFunction, err = gi.FunctionInvokerNew("WebKit2", "get_major_version")
	})
	return err
}

// GetMajorVersion is a representation of the C type webkit_get_major_version.
func GetMajorVersion() uint32 {

	var ret gi.Argument

	err := getMajorVersionFunction_Set()
	if err == nil {
		ret = getMajorVersionFunction.Invoke(nil, nil)
	}

	retGo := ret.Uint32()

	return retGo
}

var getMicroVersionFunction *gi.Function
var getMicroVersionFunction_Once sync.Once

func getMicroVersionFunction_Set() error {
	var err error
	getMicroVersionFunction_Once.Do(func() {
		getMicroVersionFunction, err = gi.FunctionInvokerNew("WebKit2", "get_micro_version")
	})
	return err
}

// GetMicroVersion is a representation of the C type webkit_get_micro_version.
func GetMicroVersion() uint32 {

	var ret gi.Argument

	err := getMicroVersionFunction_Set()
	if err == nil {
		ret = getMicroVersionFunction.Invoke(nil, nil)
	}

	retGo := ret.Uint32()

	return retGo
}

var getMinorVersionFunction *gi.Function
var getMinorVersionFunction_Once sync.Once

func getMinorVersionFunction_Set() error {
	var err error
	getMinorVersionFunction_Once.Do(func() {
		getMinorVersionFunction, err = gi.FunctionInvokerNew("WebKit2", "get_minor_version")
	})
	return err
}

// GetMinorVersion is a representation of the C type webkit_get_minor_version.
func GetMinorVersion() uint32 {

	var ret gi.Argument

	err := getMinorVersionFunction_Set()
	if err == nil {
		ret = getMinorVersionFunction.Invoke(nil, nil)
	}

	retGo := ret.Uint32()

	return retGo
}

var javascriptErrorQuarkFunction *gi.Function
var javascriptErrorQuarkFunction_Once sync.Once

func javascriptErrorQuarkFunction_Set() error {
	var err error
	javascriptErrorQuarkFunction_Once.Do(func() {
		javascriptErrorQuarkFunction, err = gi.FunctionInvokerNew("WebKit2", "javascript_error_quark")
	})
	return err
}

// JavascriptErrorQuark is a representation of the C type webkit_javascript_error_quark.
func JavascriptErrorQuark() glib.Quark {

	var ret gi.Argument

	err := javascriptErrorQuarkFunction_Set()
	if err == nil {
		ret = javascriptErrorQuarkFunction.Invoke(nil, nil)
	}

	retGo := glib.Quark(ret.Uint32())

	return retGo
}

var networkErrorQuarkFunction *gi.Function
var networkErrorQuarkFunction_Once sync.Once

func networkErrorQuarkFunction_Set() error {
	var err error
	networkErrorQuarkFunction_Once.Do(func() {
		networkErrorQuarkFunction, err = gi.FunctionInvokerNew("WebKit2", "network_error_quark")
	})
	return err
}

// NetworkErrorQuark is a representation of the C type webkit_network_error_quark.
func NetworkErrorQuark() glib.Quark {

	var ret gi.Argument

	err := networkErrorQuarkFunction_Set()
	if err == nil {
		ret = networkErrorQuarkFunction.Invoke(nil, nil)
	}

	retGo := glib.Quark(ret.Uint32())

	return retGo
}

var pluginErrorQuarkFunction *gi.Function
var pluginErrorQuarkFunction_Once sync.Once

func pluginErrorQuarkFunction_Set() error {
	var err error
	pluginErrorQuarkFunction_Once.Do(func() {
		pluginErrorQuarkFunction, err = gi.FunctionInvokerNew("WebKit2", "plugin_error_quark")
	})
	return err
}

// PluginErrorQuark is a representation of the C type webkit_plugin_error_quark.
func PluginErrorQuark() glib.Quark {

	var ret gi.Argument

	err := pluginErrorQuarkFunction_Set()
	if err == nil {
		ret = pluginErrorQuarkFunction.Invoke(nil, nil)
	}

	retGo := glib.Quark(ret.Uint32())

	return retGo
}

var policyErrorQuarkFunction *gi.Function
var policyErrorQuarkFunction_Once sync.Once

func policyErrorQuarkFunction_Set() error {
	var err error
	policyErrorQuarkFunction_Once.Do(func() {
		policyErrorQuarkFunction, err = gi.FunctionInvokerNew("WebKit2", "policy_error_quark")
	})
	return err
}

// PolicyErrorQuark is a representation of the C type webkit_policy_error_quark.
func PolicyErrorQuark() glib.Quark {

	var ret gi.Argument

	err := policyErrorQuarkFunction_Set()
	if err == nil {
		ret = policyErrorQuarkFunction.Invoke(nil, nil)
	}

	retGo := glib.Quark(ret.Uint32())

	return retGo
}

var printErrorQuarkFunction *gi.Function
var printErrorQuarkFunction_Once sync.Once

func printErrorQuarkFunction_Set() error {
	var err error
	printErrorQuarkFunction_Once.Do(func() {
		printErrorQuarkFunction, err = gi.FunctionInvokerNew("WebKit2", "print_error_quark")
	})
	return err
}

// PrintErrorQuark is a representation of the C type webkit_print_error_quark.
func PrintErrorQuark() glib.Quark {

	var ret gi.Argument

	err := printErrorQuarkFunction_Set()
	if err == nil {
		ret = printErrorQuarkFunction.Invoke(nil, nil)
	}

	retGo := glib.Quark(ret.Uint32())

	return retGo
}

var snapshotErrorQuarkFunction *gi.Function
var snapshotErrorQuarkFunction_Once sync.Once

func snapshotErrorQuarkFunction_Set() error {
	var err error
	snapshotErrorQuarkFunction_Once.Do(func() {
		snapshotErrorQuarkFunction, err = gi.FunctionInvokerNew("WebKit2", "snapshot_error_quark")
	})
	return err
}

// SnapshotErrorQuark is a representation of the C type webkit_snapshot_error_quark.
func SnapshotErrorQuark() glib.Quark {

	var ret gi.Argument

	err := snapshotErrorQuarkFunction_Set()
	if err == nil {
		ret = snapshotErrorQuarkFunction.Invoke(nil, nil)
	}

	retGo := glib.Quark(ret.Uint32())

	return retGo
}

var uriForDisplayFunction *gi.Function
var uriForDisplayFunction_Once sync.Once

func uriForDisplayFunction_Set() error {
	var err error
	uriForDisplayFunction_Once.Do(func() {
		uriForDisplayFunction, err = gi.FunctionInvokerNew("WebKit2", "uri_for_display")
	})
	return err
}

// UriForDisplay is a representation of the C type webkit_uri_for_display.
func UriForDisplay(uri string) string {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(uri)

	var ret gi.Argument

	err := uriForDisplayFunction_Set()
	if err == nil {
		ret = uriForDisplayFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

var userContentFilterErrorQuarkFunction *gi.Function
var userContentFilterErrorQuarkFunction_Once sync.Once

func userContentFilterErrorQuarkFunction_Set() error {
	var err error
	userContentFilterErrorQuarkFunction_Once.Do(func() {
		userContentFilterErrorQuarkFunction, err = gi.FunctionInvokerNew("WebKit2", "user_content_filter_error_quark")
	})
	return err
}

// UserContentFilterErrorQuark is a representation of the C type webkit_user_content_filter_error_quark.
func UserContentFilterErrorQuark() glib.Quark {

	var ret gi.Argument

	err := userContentFilterErrorQuarkFunction_Set()
	if err == nil {
		ret = userContentFilterErrorQuarkFunction.Invoke(nil, nil)
	}

	retGo := glib.Quark(ret.Uint32())

	return retGo
}

// UNSUPPORTED : C value 'webkit_user_media_permission_is_for_audio_device' : parameter 'request' of type 'UserMediaPermissionRequest' not supported

// UNSUPPORTED : C value 'webkit_user_media_permission_is_for_video_device' : parameter 'request' of type 'UserMediaPermissionRequest' not supported
