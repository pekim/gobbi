// Code generated - DO NOT EDIT.

package webkit2

import (
	gi "github.com/pekim/gobbi/internal/cgo/gi"
	"sync"
)

// UNSUPPORTED : C value 'webkit_download_error_quark' : return type 'GLib.Quark' not supported

// UNSUPPORTED : C value 'webkit_favicon_database_error_quark' : return type 'GLib.Quark' not supported

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

// UNSUPPORTED : C value 'webkit_javascript_error_quark' : return type 'GLib.Quark' not supported

// UNSUPPORTED : C value 'webkit_network_error_quark' : return type 'GLib.Quark' not supported

// UNSUPPORTED : C value 'webkit_plugin_error_quark' : return type 'GLib.Quark' not supported

// UNSUPPORTED : C value 'webkit_policy_error_quark' : return type 'GLib.Quark' not supported

// UNSUPPORTED : C value 'webkit_print_error_quark' : return type 'GLib.Quark' not supported

// UNSUPPORTED : C value 'webkit_snapshot_error_quark' : return type 'GLib.Quark' not supported

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

// UNSUPPORTED : C value 'webkit_user_content_filter_error_quark' : return type 'GLib.Quark' not supported

var userMediaPermissionIsForAudioDeviceFunction *gi.Function
var userMediaPermissionIsForAudioDeviceFunction_Once sync.Once

func userMediaPermissionIsForAudioDeviceFunction_Set() error {
	var err error
	userMediaPermissionIsForAudioDeviceFunction_Once.Do(func() {
		userMediaPermissionIsForAudioDeviceFunction, err = gi.FunctionInvokerNew("WebKit2", "user_media_permission_is_for_audio_device")
	})
	return err
}

// UserMediaPermissionIsForAudioDevice is a representation of the C type webkit_user_media_permission_is_for_audio_device.
func UserMediaPermissionIsForAudioDevice(request *UserMediaPermissionRequest) bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(request.Native())

	var ret gi.Argument

	err := userMediaPermissionIsForAudioDeviceFunction_Set()
	if err == nil {
		ret = userMediaPermissionIsForAudioDeviceFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var userMediaPermissionIsForVideoDeviceFunction *gi.Function
var userMediaPermissionIsForVideoDeviceFunction_Once sync.Once

func userMediaPermissionIsForVideoDeviceFunction_Set() error {
	var err error
	userMediaPermissionIsForVideoDeviceFunction_Once.Do(func() {
		userMediaPermissionIsForVideoDeviceFunction, err = gi.FunctionInvokerNew("WebKit2", "user_media_permission_is_for_video_device")
	})
	return err
}

// UserMediaPermissionIsForVideoDevice is a representation of the C type webkit_user_media_permission_is_for_video_device.
func UserMediaPermissionIsForVideoDevice(request *UserMediaPermissionRequest) bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(request.Native())

	var ret gi.Argument

	err := userMediaPermissionIsForVideoDeviceFunction_Set()
	if err == nil {
		ret = userMediaPermissionIsForVideoDeviceFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}
