// Code generated - DO NOT EDIT.

package gio

import (
	gi "github.com/pekim/gobbi/internal/gi"
	glib "github.com/pekim/gobbi/lib/glib"
	"sync"
)

var actionNameIsValidFunction *gi.Function
var actionNameIsValidFunction_Once sync.Once

func actionNameIsValidFunction_Set() error {
	var err error
	actionNameIsValidFunction_Once.Do(func() {
		actionNameIsValidFunction, err = gi.FunctionInvokerNew("Gio", "action_name_is_valid")
	})
	return err
}

// ActionNameIsValid is a representation of the C type g_action_name_is_valid.
func ActionNameIsValid(actionName string) bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(actionName)

	var ret gi.Argument

	err := actionNameIsValidFunction_Set()
	if err == nil {
		ret = actionNameIsValidFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'g_action_parse_detailed_name' : parameter 'target_value' of type 'GLib.Variant' not supported

// UNSUPPORTED : C value 'g_action_print_detailed_name' : parameter 'target_value' of type 'GLib.Variant' not supported

// UNSUPPORTED : C value 'g_app_info_create_from_commandline' : parameter 'flags' of type 'AppInfoCreateFlags' not supported

// UNSUPPORTED : C value 'g_app_info_get_all' : return type 'GLib.List' not supported

// UNSUPPORTED : C value 'g_app_info_get_all_for_type' : return type 'GLib.List' not supported

// UNSUPPORTED : C value 'g_app_info_get_default_for_type' : return type 'AppInfo' not supported

// UNSUPPORTED : C value 'g_app_info_get_default_for_uri_scheme' : return type 'AppInfo' not supported

// UNSUPPORTED : C value 'g_app_info_get_fallback_for_type' : return type 'GLib.List' not supported

// UNSUPPORTED : C value 'g_app_info_get_recommended_for_type' : return type 'GLib.List' not supported

var appInfoLaunchDefaultForUriFunction *gi.Function
var appInfoLaunchDefaultForUriFunction_Once sync.Once

func appInfoLaunchDefaultForUriFunction_Set() error {
	var err error
	appInfoLaunchDefaultForUriFunction_Once.Do(func() {
		appInfoLaunchDefaultForUriFunction, err = gi.FunctionInvokerNew("Gio", "app_info_launch_default_for_uri")
	})
	return err
}

// AppInfoLaunchDefaultForUri is a representation of the C type g_app_info_launch_default_for_uri.
func AppInfoLaunchDefaultForUri(uri string, context *AppLaunchContext) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(uri)
	inArgs[1].SetPointer(context.native)

	var ret gi.Argument

	err := appInfoLaunchDefaultForUriFunction_Set()
	if err == nil {
		ret = appInfoLaunchDefaultForUriFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'g_app_info_launch_default_for_uri_async' : parameter 'callback' of type 'AsyncReadyCallback' not supported

// UNSUPPORTED : C value 'g_app_info_launch_default_for_uri_finish' : parameter 'result' of type 'AsyncResult' not supported

var appInfoResetTypeAssociationsFunction *gi.Function
var appInfoResetTypeAssociationsFunction_Once sync.Once

func appInfoResetTypeAssociationsFunction_Set() error {
	var err error
	appInfoResetTypeAssociationsFunction_Once.Do(func() {
		appInfoResetTypeAssociationsFunction, err = gi.FunctionInvokerNew("Gio", "app_info_reset_type_associations")
	})
	return err
}

// AppInfoResetTypeAssociations is a representation of the C type g_app_info_reset_type_associations.
func AppInfoResetTypeAssociations(contentType string) {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(contentType)

	err := appInfoResetTypeAssociationsFunction_Set()
	if err == nil {
		appInfoResetTypeAssociationsFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'g_async_initable_newv_async' : parameter 'object_type' of type 'GType' not supported

// UNSUPPORTED : C value 'g_bus_get' : parameter 'bus_type' of type 'BusType' not supported

// UNSUPPORTED : C value 'g_bus_get_finish' : parameter 'res' of type 'AsyncResult' not supported

// UNSUPPORTED : C value 'g_bus_get_sync' : parameter 'bus_type' of type 'BusType' not supported

// UNSUPPORTED : C value 'g_bus_own_name' : parameter 'bus_type' of type 'BusType' not supported

// UNSUPPORTED : C value 'g_bus_own_name_on_connection' : parameter 'flags' of type 'BusNameOwnerFlags' not supported

// UNSUPPORTED : C value 'g_bus_own_name_on_connection_with_closures' : parameter 'flags' of type 'BusNameOwnerFlags' not supported

// UNSUPPORTED : C value 'g_bus_own_name_with_closures' : parameter 'bus_type' of type 'BusType' not supported

var busUnownNameFunction *gi.Function
var busUnownNameFunction_Once sync.Once

func busUnownNameFunction_Set() error {
	var err error
	busUnownNameFunction_Once.Do(func() {
		busUnownNameFunction, err = gi.FunctionInvokerNew("Gio", "bus_unown_name")
	})
	return err
}

// BusUnownName is a representation of the C type g_bus_unown_name.
func BusUnownName(ownerId uint32) {
	var inArgs [1]gi.Argument
	inArgs[0].SetUint32(ownerId)

	err := busUnownNameFunction_Set()
	if err == nil {
		busUnownNameFunction.Invoke(inArgs[:], nil)
	}

	return
}

var busUnwatchNameFunction *gi.Function
var busUnwatchNameFunction_Once sync.Once

func busUnwatchNameFunction_Set() error {
	var err error
	busUnwatchNameFunction_Once.Do(func() {
		busUnwatchNameFunction, err = gi.FunctionInvokerNew("Gio", "bus_unwatch_name")
	})
	return err
}

// BusUnwatchName is a representation of the C type g_bus_unwatch_name.
func BusUnwatchName(watcherId uint32) {
	var inArgs [1]gi.Argument
	inArgs[0].SetUint32(watcherId)

	err := busUnwatchNameFunction_Set()
	if err == nil {
		busUnwatchNameFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'g_bus_watch_name' : parameter 'bus_type' of type 'BusType' not supported

// UNSUPPORTED : C value 'g_bus_watch_name_on_connection' : parameter 'flags' of type 'BusNameWatcherFlags' not supported

// UNSUPPORTED : C value 'g_bus_watch_name_on_connection_with_closures' : parameter 'flags' of type 'BusNameWatcherFlags' not supported

// UNSUPPORTED : C value 'g_bus_watch_name_with_closures' : parameter 'bus_type' of type 'BusType' not supported

var contentTypeCanBeExecutableFunction *gi.Function
var contentTypeCanBeExecutableFunction_Once sync.Once

func contentTypeCanBeExecutableFunction_Set() error {
	var err error
	contentTypeCanBeExecutableFunction_Once.Do(func() {
		contentTypeCanBeExecutableFunction, err = gi.FunctionInvokerNew("Gio", "content_type_can_be_executable")
	})
	return err
}

// ContentTypeCanBeExecutable is a representation of the C type g_content_type_can_be_executable.
func ContentTypeCanBeExecutable(type_ string) bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(type_)

	var ret gi.Argument

	err := contentTypeCanBeExecutableFunction_Set()
	if err == nil {
		ret = contentTypeCanBeExecutableFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var contentTypeEqualsFunction *gi.Function
var contentTypeEqualsFunction_Once sync.Once

func contentTypeEqualsFunction_Set() error {
	var err error
	contentTypeEqualsFunction_Once.Do(func() {
		contentTypeEqualsFunction, err = gi.FunctionInvokerNew("Gio", "content_type_equals")
	})
	return err
}

// ContentTypeEquals is a representation of the C type g_content_type_equals.
func ContentTypeEquals(type1 string, type2 string) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(type1)
	inArgs[1].SetString(type2)

	var ret gi.Argument

	err := contentTypeEqualsFunction_Set()
	if err == nil {
		ret = contentTypeEqualsFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var contentTypeFromMimeTypeFunction *gi.Function
var contentTypeFromMimeTypeFunction_Once sync.Once

func contentTypeFromMimeTypeFunction_Set() error {
	var err error
	contentTypeFromMimeTypeFunction_Once.Do(func() {
		contentTypeFromMimeTypeFunction, err = gi.FunctionInvokerNew("Gio", "content_type_from_mime_type")
	})
	return err
}

// ContentTypeFromMimeType is a representation of the C type g_content_type_from_mime_type.
func ContentTypeFromMimeType(mimeType string) string {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(mimeType)

	var ret gi.Argument

	err := contentTypeFromMimeTypeFunction_Set()
	if err == nil {
		ret = contentTypeFromMimeTypeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

var contentTypeGetDescriptionFunction *gi.Function
var contentTypeGetDescriptionFunction_Once sync.Once

func contentTypeGetDescriptionFunction_Set() error {
	var err error
	contentTypeGetDescriptionFunction_Once.Do(func() {
		contentTypeGetDescriptionFunction, err = gi.FunctionInvokerNew("Gio", "content_type_get_description")
	})
	return err
}

// ContentTypeGetDescription is a representation of the C type g_content_type_get_description.
func ContentTypeGetDescription(type_ string) string {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(type_)

	var ret gi.Argument

	err := contentTypeGetDescriptionFunction_Set()
	if err == nil {
		ret = contentTypeGetDescriptionFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

var contentTypeGetGenericIconNameFunction *gi.Function
var contentTypeGetGenericIconNameFunction_Once sync.Once

func contentTypeGetGenericIconNameFunction_Set() error {
	var err error
	contentTypeGetGenericIconNameFunction_Once.Do(func() {
		contentTypeGetGenericIconNameFunction, err = gi.FunctionInvokerNew("Gio", "content_type_get_generic_icon_name")
	})
	return err
}

// ContentTypeGetGenericIconName is a representation of the C type g_content_type_get_generic_icon_name.
func ContentTypeGetGenericIconName(type_ string) string {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(type_)

	var ret gi.Argument

	err := contentTypeGetGenericIconNameFunction_Set()
	if err == nil {
		ret = contentTypeGetGenericIconNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

// UNSUPPORTED : C value 'g_content_type_get_icon' : return type 'Icon' not supported

var contentTypeGetMimeDirsFunction *gi.Function
var contentTypeGetMimeDirsFunction_Once sync.Once

func contentTypeGetMimeDirsFunction_Set() error {
	var err error
	contentTypeGetMimeDirsFunction_Once.Do(func() {
		contentTypeGetMimeDirsFunction, err = gi.FunctionInvokerNew("Gio", "content_type_get_mime_dirs")
	})
	return err
}

// ContentTypeGetMimeDirs is a representation of the C type g_content_type_get_mime_dirs.
func ContentTypeGetMimeDirs() {

	err := contentTypeGetMimeDirsFunction_Set()
	if err == nil {
		contentTypeGetMimeDirsFunction.Invoke(nil, nil)
	}

	return
}

var contentTypeGetMimeTypeFunction *gi.Function
var contentTypeGetMimeTypeFunction_Once sync.Once

func contentTypeGetMimeTypeFunction_Set() error {
	var err error
	contentTypeGetMimeTypeFunction_Once.Do(func() {
		contentTypeGetMimeTypeFunction, err = gi.FunctionInvokerNew("Gio", "content_type_get_mime_type")
	})
	return err
}

// ContentTypeGetMimeType is a representation of the C type g_content_type_get_mime_type.
func ContentTypeGetMimeType(type_ string) string {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(type_)

	var ret gi.Argument

	err := contentTypeGetMimeTypeFunction_Set()
	if err == nil {
		ret = contentTypeGetMimeTypeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

// UNSUPPORTED : C value 'g_content_type_get_symbolic_icon' : return type 'Icon' not supported

// UNSUPPORTED : C value 'g_content_type_guess' : parameter 'data' of type 'nil' not supported

// UNSUPPORTED : C value 'g_content_type_guess_for_tree' : parameter 'root' of type 'File' not supported

var contentTypeIsAFunction *gi.Function
var contentTypeIsAFunction_Once sync.Once

func contentTypeIsAFunction_Set() error {
	var err error
	contentTypeIsAFunction_Once.Do(func() {
		contentTypeIsAFunction, err = gi.FunctionInvokerNew("Gio", "content_type_is_a")
	})
	return err
}

// ContentTypeIsA is a representation of the C type g_content_type_is_a.
func ContentTypeIsA(type_ string, supertype string) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(type_)
	inArgs[1].SetString(supertype)

	var ret gi.Argument

	err := contentTypeIsAFunction_Set()
	if err == nil {
		ret = contentTypeIsAFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var contentTypeIsMimeTypeFunction *gi.Function
var contentTypeIsMimeTypeFunction_Once sync.Once

func contentTypeIsMimeTypeFunction_Set() error {
	var err error
	contentTypeIsMimeTypeFunction_Once.Do(func() {
		contentTypeIsMimeTypeFunction, err = gi.FunctionInvokerNew("Gio", "content_type_is_mime_type")
	})
	return err
}

// ContentTypeIsMimeType is a representation of the C type g_content_type_is_mime_type.
func ContentTypeIsMimeType(type_ string, mimeType string) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(type_)
	inArgs[1].SetString(mimeType)

	var ret gi.Argument

	err := contentTypeIsMimeTypeFunction_Set()
	if err == nil {
		ret = contentTypeIsMimeTypeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var contentTypeIsUnknownFunction *gi.Function
var contentTypeIsUnknownFunction_Once sync.Once

func contentTypeIsUnknownFunction_Set() error {
	var err error
	contentTypeIsUnknownFunction_Once.Do(func() {
		contentTypeIsUnknownFunction, err = gi.FunctionInvokerNew("Gio", "content_type_is_unknown")
	})
	return err
}

// ContentTypeIsUnknown is a representation of the C type g_content_type_is_unknown.
func ContentTypeIsUnknown(type_ string) bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(type_)

	var ret gi.Argument

	err := contentTypeIsUnknownFunction_Set()
	if err == nil {
		ret = contentTypeIsUnknownFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'g_content_type_set_mime_dirs' : parameter 'dirs' of type 'nil' not supported

// UNSUPPORTED : C value 'g_content_types_get_registered' : return type 'GLib.List' not supported

var dbusAddressEscapeValueFunction *gi.Function
var dbusAddressEscapeValueFunction_Once sync.Once

func dbusAddressEscapeValueFunction_Set() error {
	var err error
	dbusAddressEscapeValueFunction_Once.Do(func() {
		dbusAddressEscapeValueFunction, err = gi.FunctionInvokerNew("Gio", "dbus_address_escape_value")
	})
	return err
}

// DbusAddressEscapeValue is a representation of the C type g_dbus_address_escape_value.
func DbusAddressEscapeValue(string_ string) string {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(string_)

	var ret gi.Argument

	err := dbusAddressEscapeValueFunction_Set()
	if err == nil {
		ret = dbusAddressEscapeValueFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

// UNSUPPORTED : C value 'g_dbus_address_get_for_bus_sync' : parameter 'bus_type' of type 'BusType' not supported

// UNSUPPORTED : C value 'g_dbus_address_get_stream' : parameter 'callback' of type 'AsyncReadyCallback' not supported

// UNSUPPORTED : C value 'g_dbus_address_get_stream_finish' : parameter 'res' of type 'AsyncResult' not supported

var dbusAddressGetStreamSyncFunction *gi.Function
var dbusAddressGetStreamSyncFunction_Once sync.Once

func dbusAddressGetStreamSyncFunction_Set() error {
	var err error
	dbusAddressGetStreamSyncFunction_Once.Do(func() {
		dbusAddressGetStreamSyncFunction, err = gi.FunctionInvokerNew("Gio", "dbus_address_get_stream_sync")
	})
	return err
}

// DbusAddressGetStreamSync is a representation of the C type g_dbus_address_get_stream_sync.
func DbusAddressGetStreamSync(address string, cancellable *Cancellable) (*IOStream, string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(address)
	inArgs[1].SetPointer(cancellable.native)

	var outArgs [1]gi.Argument
	var ret gi.Argument

	err := dbusAddressGetStreamSyncFunction_Set()
	if err == nil {
		ret = dbusAddressGetStreamSyncFunction.Invoke(inArgs[:], outArgs[:])
	}

	retGo := &IOStream{native: ret.Pointer()}
	out0 := outArgs[0].String(true)

	return retGo, out0
}

// UNSUPPORTED : C value 'g_dbus_annotation_info_lookup' : parameter 'annotations' of type 'nil' not supported

// UNSUPPORTED : C value 'g_dbus_error_encode_gerror' : parameter 'error' of type 'GLib.Error' not supported

// UNSUPPORTED : C value 'g_dbus_error_get_remote_error' : parameter 'error' of type 'GLib.Error' not supported

// UNSUPPORTED : C value 'g_dbus_error_is_remote_error' : parameter 'error' of type 'GLib.Error' not supported

// UNSUPPORTED : C value 'g_dbus_error_new_for_dbus_error' : return type 'GLib.Error' not supported

var dbusErrorQuarkFunction *gi.Function
var dbusErrorQuarkFunction_Once sync.Once

func dbusErrorQuarkFunction_Set() error {
	var err error
	dbusErrorQuarkFunction_Once.Do(func() {
		dbusErrorQuarkFunction, err = gi.FunctionInvokerNew("Gio", "dbus_error_quark")
	})
	return err
}

// DbusErrorQuark is a representation of the C type g_dbus_error_quark.
func DbusErrorQuark() glib.Quark {

	var ret gi.Argument

	err := dbusErrorQuarkFunction_Set()
	if err == nil {
		ret = dbusErrorQuarkFunction.Invoke(nil, nil)
	}

	retGo := glib.Quark(ret.Uint32())

	return retGo
}

var dbusErrorRegisterErrorFunction *gi.Function
var dbusErrorRegisterErrorFunction_Once sync.Once

func dbusErrorRegisterErrorFunction_Set() error {
	var err error
	dbusErrorRegisterErrorFunction_Once.Do(func() {
		dbusErrorRegisterErrorFunction, err = gi.FunctionInvokerNew("Gio", "dbus_error_register_error")
	})
	return err
}

// DbusErrorRegisterError is a representation of the C type g_dbus_error_register_error.
func DbusErrorRegisterError(errorDomain glib.Quark, errorCode int32, dbusErrorName string) bool {
	var inArgs [3]gi.Argument
	inArgs[0].SetUint32(uint32(errorDomain))
	inArgs[1].SetInt32(errorCode)
	inArgs[2].SetString(dbusErrorName)

	var ret gi.Argument

	err := dbusErrorRegisterErrorFunction_Set()
	if err == nil {
		ret = dbusErrorRegisterErrorFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'g_dbus_error_register_error_domain' : parameter 'entries' of type 'nil' not supported

// UNSUPPORTED : C value 'g_dbus_error_strip_remote_error' : parameter 'error' of type 'GLib.Error' not supported

var dbusErrorUnregisterErrorFunction *gi.Function
var dbusErrorUnregisterErrorFunction_Once sync.Once

func dbusErrorUnregisterErrorFunction_Set() error {
	var err error
	dbusErrorUnregisterErrorFunction_Once.Do(func() {
		dbusErrorUnregisterErrorFunction, err = gi.FunctionInvokerNew("Gio", "dbus_error_unregister_error")
	})
	return err
}

// DbusErrorUnregisterError is a representation of the C type g_dbus_error_unregister_error.
func DbusErrorUnregisterError(errorDomain glib.Quark, errorCode int32, dbusErrorName string) bool {
	var inArgs [3]gi.Argument
	inArgs[0].SetUint32(uint32(errorDomain))
	inArgs[1].SetInt32(errorCode)
	inArgs[2].SetString(dbusErrorName)

	var ret gi.Argument

	err := dbusErrorUnregisterErrorFunction_Set()
	if err == nil {
		ret = dbusErrorUnregisterErrorFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var dbusGenerateGuidFunction *gi.Function
var dbusGenerateGuidFunction_Once sync.Once

func dbusGenerateGuidFunction_Set() error {
	var err error
	dbusGenerateGuidFunction_Once.Do(func() {
		dbusGenerateGuidFunction, err = gi.FunctionInvokerNew("Gio", "dbus_generate_guid")
	})
	return err
}

// DbusGenerateGuid is a representation of the C type g_dbus_generate_guid.
func DbusGenerateGuid() string {

	var ret gi.Argument

	err := dbusGenerateGuidFunction_Set()
	if err == nil {
		ret = dbusGenerateGuidFunction.Invoke(nil, nil)
	}

	retGo := ret.String(true)

	return retGo
}

// UNSUPPORTED : C value 'g_dbus_gvalue_to_gvariant' : parameter 'gvalue' of type 'GObject.Value' not supported

// UNSUPPORTED : C value 'g_dbus_gvariant_to_gvalue' : parameter 'value' of type 'GLib.Variant' not supported

var dbusIsAddressFunction *gi.Function
var dbusIsAddressFunction_Once sync.Once

func dbusIsAddressFunction_Set() error {
	var err error
	dbusIsAddressFunction_Once.Do(func() {
		dbusIsAddressFunction, err = gi.FunctionInvokerNew("Gio", "dbus_is_address")
	})
	return err
}

// DbusIsAddress is a representation of the C type g_dbus_is_address.
func DbusIsAddress(string_ string) bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(string_)

	var ret gi.Argument

	err := dbusIsAddressFunction_Set()
	if err == nil {
		ret = dbusIsAddressFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var dbusIsGuidFunction *gi.Function
var dbusIsGuidFunction_Once sync.Once

func dbusIsGuidFunction_Set() error {
	var err error
	dbusIsGuidFunction_Once.Do(func() {
		dbusIsGuidFunction, err = gi.FunctionInvokerNew("Gio", "dbus_is_guid")
	})
	return err
}

// DbusIsGuid is a representation of the C type g_dbus_is_guid.
func DbusIsGuid(string_ string) bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(string_)

	var ret gi.Argument

	err := dbusIsGuidFunction_Set()
	if err == nil {
		ret = dbusIsGuidFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var dbusIsInterfaceNameFunction *gi.Function
var dbusIsInterfaceNameFunction_Once sync.Once

func dbusIsInterfaceNameFunction_Set() error {
	var err error
	dbusIsInterfaceNameFunction_Once.Do(func() {
		dbusIsInterfaceNameFunction, err = gi.FunctionInvokerNew("Gio", "dbus_is_interface_name")
	})
	return err
}

// DbusIsInterfaceName is a representation of the C type g_dbus_is_interface_name.
func DbusIsInterfaceName(string_ string) bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(string_)

	var ret gi.Argument

	err := dbusIsInterfaceNameFunction_Set()
	if err == nil {
		ret = dbusIsInterfaceNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var dbusIsMemberNameFunction *gi.Function
var dbusIsMemberNameFunction_Once sync.Once

func dbusIsMemberNameFunction_Set() error {
	var err error
	dbusIsMemberNameFunction_Once.Do(func() {
		dbusIsMemberNameFunction, err = gi.FunctionInvokerNew("Gio", "dbus_is_member_name")
	})
	return err
}

// DbusIsMemberName is a representation of the C type g_dbus_is_member_name.
func DbusIsMemberName(string_ string) bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(string_)

	var ret gi.Argument

	err := dbusIsMemberNameFunction_Set()
	if err == nil {
		ret = dbusIsMemberNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var dbusIsNameFunction *gi.Function
var dbusIsNameFunction_Once sync.Once

func dbusIsNameFunction_Set() error {
	var err error
	dbusIsNameFunction_Once.Do(func() {
		dbusIsNameFunction, err = gi.FunctionInvokerNew("Gio", "dbus_is_name")
	})
	return err
}

// DbusIsName is a representation of the C type g_dbus_is_name.
func DbusIsName(string_ string) bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(string_)

	var ret gi.Argument

	err := dbusIsNameFunction_Set()
	if err == nil {
		ret = dbusIsNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var dbusIsSupportedAddressFunction *gi.Function
var dbusIsSupportedAddressFunction_Once sync.Once

func dbusIsSupportedAddressFunction_Set() error {
	var err error
	dbusIsSupportedAddressFunction_Once.Do(func() {
		dbusIsSupportedAddressFunction, err = gi.FunctionInvokerNew("Gio", "dbus_is_supported_address")
	})
	return err
}

// DbusIsSupportedAddress is a representation of the C type g_dbus_is_supported_address.
func DbusIsSupportedAddress(string_ string) bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(string_)

	var ret gi.Argument

	err := dbusIsSupportedAddressFunction_Set()
	if err == nil {
		ret = dbusIsSupportedAddressFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var dbusIsUniqueNameFunction *gi.Function
var dbusIsUniqueNameFunction_Once sync.Once

func dbusIsUniqueNameFunction_Set() error {
	var err error
	dbusIsUniqueNameFunction_Once.Do(func() {
		dbusIsUniqueNameFunction, err = gi.FunctionInvokerNew("Gio", "dbus_is_unique_name")
	})
	return err
}

// DbusIsUniqueName is a representation of the C type g_dbus_is_unique_name.
func DbusIsUniqueName(string_ string) bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(string_)

	var ret gi.Argument

	err := dbusIsUniqueNameFunction_Set()
	if err == nil {
		ret = dbusIsUniqueNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'g_dtls_client_connection_new' : parameter 'base_socket' of type 'DatagramBased' not supported

// UNSUPPORTED : C value 'g_dtls_server_connection_new' : parameter 'base_socket' of type 'DatagramBased' not supported

// UNSUPPORTED : C value 'g_file_new_for_commandline_arg' : return type 'File' not supported

// UNSUPPORTED : C value 'g_file_new_for_commandline_arg_and_cwd' : return type 'File' not supported

// UNSUPPORTED : C value 'g_file_new_for_path' : return type 'File' not supported

// UNSUPPORTED : C value 'g_file_new_for_uri' : return type 'File' not supported

// UNSUPPORTED : C value 'g_file_new_tmp' : return type 'File' not supported

// UNSUPPORTED : C value 'g_file_parse_name' : return type 'File' not supported

// UNSUPPORTED : C value 'g_icon_deserialize' : parameter 'value' of type 'GLib.Variant' not supported

// UNSUPPORTED : C value 'g_icon_hash' : parameter 'icon' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_icon_new_for_string' : return type 'Icon' not supported

// UNSUPPORTED : C value 'g_initable_newv' : parameter 'object_type' of type 'GType' not supported

// UNSUPPORTED : C value 'g_io_error_from_errno' : return type 'IOErrorEnum' not supported

var ioErrorQuarkFunction *gi.Function
var ioErrorQuarkFunction_Once sync.Once

func ioErrorQuarkFunction_Set() error {
	var err error
	ioErrorQuarkFunction_Once.Do(func() {
		ioErrorQuarkFunction, err = gi.FunctionInvokerNew("Gio", "io_error_quark")
	})
	return err
}

// IoErrorQuark is a representation of the C type g_io_error_quark.
func IoErrorQuark() glib.Quark {

	var ret gi.Argument

	err := ioErrorQuarkFunction_Set()
	if err == nil {
		ret = ioErrorQuarkFunction.Invoke(nil, nil)
	}

	retGo := glib.Quark(ret.Uint32())

	return retGo
}

// UNSUPPORTED : C value 'g_io_extension_point_implement' : parameter 'type' of type 'GType' not supported

var ioExtensionPointLookupFunction *gi.Function
var ioExtensionPointLookupFunction_Once sync.Once

func ioExtensionPointLookupFunction_Set() error {
	var err error
	ioExtensionPointLookupFunction_Once.Do(func() {
		ioExtensionPointLookupFunction, err = gi.FunctionInvokerNew("Gio", "io_extension_point_lookup")
	})
	return err
}

// IoExtensionPointLookup is a representation of the C type g_io_extension_point_lookup.
func IoExtensionPointLookup(name string) *IOExtensionPoint {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(name)

	var ret gi.Argument

	err := ioExtensionPointLookupFunction_Set()
	if err == nil {
		ret = ioExtensionPointLookupFunction.Invoke(inArgs[:], nil)
	}

	retGo := &IOExtensionPoint{native: ret.Pointer()}

	return retGo
}

var ioExtensionPointRegisterFunction *gi.Function
var ioExtensionPointRegisterFunction_Once sync.Once

func ioExtensionPointRegisterFunction_Set() error {
	var err error
	ioExtensionPointRegisterFunction_Once.Do(func() {
		ioExtensionPointRegisterFunction, err = gi.FunctionInvokerNew("Gio", "io_extension_point_register")
	})
	return err
}

// IoExtensionPointRegister is a representation of the C type g_io_extension_point_register.
func IoExtensionPointRegister(name string) *IOExtensionPoint {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(name)

	var ret gi.Argument

	err := ioExtensionPointRegisterFunction_Set()
	if err == nil {
		ret = ioExtensionPointRegisterFunction.Invoke(inArgs[:], nil)
	}

	retGo := &IOExtensionPoint{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_io_modules_load_all_in_directory' : return type 'GLib.List' not supported

// UNSUPPORTED : C value 'g_io_modules_load_all_in_directory_with_scope' : return type 'GLib.List' not supported

var ioModulesScanAllInDirectoryFunction *gi.Function
var ioModulesScanAllInDirectoryFunction_Once sync.Once

func ioModulesScanAllInDirectoryFunction_Set() error {
	var err error
	ioModulesScanAllInDirectoryFunction_Once.Do(func() {
		ioModulesScanAllInDirectoryFunction, err = gi.FunctionInvokerNew("Gio", "io_modules_scan_all_in_directory")
	})
	return err
}

// IoModulesScanAllInDirectory is a representation of the C type g_io_modules_scan_all_in_directory.
func IoModulesScanAllInDirectory(dirname string) {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(dirname)

	err := ioModulesScanAllInDirectoryFunction_Set()
	if err == nil {
		ioModulesScanAllInDirectoryFunction.Invoke(inArgs[:], nil)
	}

	return
}

var ioModulesScanAllInDirectoryWithScopeFunction *gi.Function
var ioModulesScanAllInDirectoryWithScopeFunction_Once sync.Once

func ioModulesScanAllInDirectoryWithScopeFunction_Set() error {
	var err error
	ioModulesScanAllInDirectoryWithScopeFunction_Once.Do(func() {
		ioModulesScanAllInDirectoryWithScopeFunction, err = gi.FunctionInvokerNew("Gio", "io_modules_scan_all_in_directory_with_scope")
	})
	return err
}

// IoModulesScanAllInDirectoryWithScope is a representation of the C type g_io_modules_scan_all_in_directory_with_scope.
func IoModulesScanAllInDirectoryWithScope(dirname string, scope *IOModuleScope) {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(dirname)
	inArgs[1].SetPointer(scope.native)

	err := ioModulesScanAllInDirectoryWithScopeFunction_Set()
	if err == nil {
		ioModulesScanAllInDirectoryWithScopeFunction.Invoke(inArgs[:], nil)
	}

	return
}

var ioSchedulerCancelAllJobsFunction *gi.Function
var ioSchedulerCancelAllJobsFunction_Once sync.Once

func ioSchedulerCancelAllJobsFunction_Set() error {
	var err error
	ioSchedulerCancelAllJobsFunction_Once.Do(func() {
		ioSchedulerCancelAllJobsFunction, err = gi.FunctionInvokerNew("Gio", "io_scheduler_cancel_all_jobs")
	})
	return err
}

// IoSchedulerCancelAllJobs is a representation of the C type g_io_scheduler_cancel_all_jobs.
func IoSchedulerCancelAllJobs() {

	err := ioSchedulerCancelAllJobsFunction_Set()
	if err == nil {
		ioSchedulerCancelAllJobsFunction.Invoke(nil, nil)
	}

	return
}

// UNSUPPORTED : C value 'g_io_scheduler_push_job' : parameter 'job_func' of type 'IOSchedulerJobFunc' not supported

var keyfileSettingsBackendNewFunction *gi.Function
var keyfileSettingsBackendNewFunction_Once sync.Once

func keyfileSettingsBackendNewFunction_Set() error {
	var err error
	keyfileSettingsBackendNewFunction_Once.Do(func() {
		keyfileSettingsBackendNewFunction, err = gi.FunctionInvokerNew("Gio", "keyfile_settings_backend_new")
	})
	return err
}

// KeyfileSettingsBackendNew is a representation of the C type g_keyfile_settings_backend_new.
func KeyfileSettingsBackendNew(filename string, rootPath string, rootGroup string) *SettingsBackend {
	var inArgs [3]gi.Argument
	inArgs[0].SetString(filename)
	inArgs[1].SetString(rootPath)
	inArgs[2].SetString(rootGroup)

	var ret gi.Argument

	err := keyfileSettingsBackendNewFunction_Set()
	if err == nil {
		ret = keyfileSettingsBackendNewFunction.Invoke(inArgs[:], nil)
	}

	retGo := &SettingsBackend{native: ret.Pointer()}

	return retGo
}

var memorySettingsBackendNewFunction *gi.Function
var memorySettingsBackendNewFunction_Once sync.Once

func memorySettingsBackendNewFunction_Set() error {
	var err error
	memorySettingsBackendNewFunction_Once.Do(func() {
		memorySettingsBackendNewFunction, err = gi.FunctionInvokerNew("Gio", "memory_settings_backend_new")
	})
	return err
}

// MemorySettingsBackendNew is a representation of the C type g_memory_settings_backend_new.
func MemorySettingsBackendNew() *SettingsBackend {

	var ret gi.Argument

	err := memorySettingsBackendNewFunction_Set()
	if err == nil {
		ret = memorySettingsBackendNewFunction.Invoke(nil, nil)
	}

	retGo := &SettingsBackend{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_network_monitor_get_default' : return type 'NetworkMonitor' not supported

var networkingInitFunction *gi.Function
var networkingInitFunction_Once sync.Once

func networkingInitFunction_Set() error {
	var err error
	networkingInitFunction_Once.Do(func() {
		networkingInitFunction, err = gi.FunctionInvokerNew("Gio", "networking_init")
	})
	return err
}

// NetworkingInit is a representation of the C type g_networking_init.
func NetworkingInit() {

	err := networkingInitFunction_Set()
	if err == nil {
		networkingInitFunction.Invoke(nil, nil)
	}

	return
}

var nullSettingsBackendNewFunction *gi.Function
var nullSettingsBackendNewFunction_Once sync.Once

func nullSettingsBackendNewFunction_Set() error {
	var err error
	nullSettingsBackendNewFunction_Once.Do(func() {
		nullSettingsBackendNewFunction, err = gi.FunctionInvokerNew("Gio", "null_settings_backend_new")
	})
	return err
}

// NullSettingsBackendNew is a representation of the C type g_null_settings_backend_new.
func NullSettingsBackendNew() *SettingsBackend {

	var ret gi.Argument

	err := nullSettingsBackendNewFunction_Set()
	if err == nil {
		ret = nullSettingsBackendNewFunction.Invoke(nil, nil)
	}

	retGo := &SettingsBackend{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_pollable_source_new' : parameter 'pollable_stream' of type 'GObject.Object' not supported

// UNSUPPORTED : C value 'g_pollable_source_new_full' : parameter 'pollable_stream' of type 'GObject.Object' not supported

// UNSUPPORTED : C value 'g_pollable_stream_read' : parameter 'buffer' of type 'nil' not supported

// UNSUPPORTED : C value 'g_pollable_stream_write' : parameter 'buffer' of type 'nil' not supported

// UNSUPPORTED : C value 'g_pollable_stream_write_all' : parameter 'buffer' of type 'nil' not supported

// UNSUPPORTED : C value 'g_proxy_get_default_for_protocol' : return type 'Proxy' not supported

// UNSUPPORTED : C value 'g_proxy_resolver_get_default' : return type 'ProxyResolver' not supported

var resolverErrorQuarkFunction *gi.Function
var resolverErrorQuarkFunction_Once sync.Once

func resolverErrorQuarkFunction_Set() error {
	var err error
	resolverErrorQuarkFunction_Once.Do(func() {
		resolverErrorQuarkFunction, err = gi.FunctionInvokerNew("Gio", "resolver_error_quark")
	})
	return err
}

// ResolverErrorQuark is a representation of the C type g_resolver_error_quark.
func ResolverErrorQuark() glib.Quark {

	var ret gi.Argument

	err := resolverErrorQuarkFunction_Set()
	if err == nil {
		ret = resolverErrorQuarkFunction.Invoke(nil, nil)
	}

	retGo := glib.Quark(ret.Uint32())

	return retGo
}

var resourceErrorQuarkFunction *gi.Function
var resourceErrorQuarkFunction_Once sync.Once

func resourceErrorQuarkFunction_Set() error {
	var err error
	resourceErrorQuarkFunction_Once.Do(func() {
		resourceErrorQuarkFunction, err = gi.FunctionInvokerNew("Gio", "resource_error_quark")
	})
	return err
}

// ResourceErrorQuark is a representation of the C type g_resource_error_quark.
func ResourceErrorQuark() glib.Quark {

	var ret gi.Argument

	err := resourceErrorQuarkFunction_Set()
	if err == nil {
		ret = resourceErrorQuarkFunction.Invoke(nil, nil)
	}

	retGo := glib.Quark(ret.Uint32())

	return retGo
}

var resourceLoadFunction *gi.Function
var resourceLoadFunction_Once sync.Once

func resourceLoadFunction_Set() error {
	var err error
	resourceLoadFunction_Once.Do(func() {
		resourceLoadFunction, err = gi.FunctionInvokerNew("Gio", "resource_load")
	})
	return err
}

// ResourceLoad is a representation of the C type g_resource_load.
func ResourceLoad(filename string) *Resource {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(filename)

	var ret gi.Argument

	err := resourceLoadFunction_Set()
	if err == nil {
		ret = resourceLoadFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Resource{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_resources_enumerate_children' : parameter 'lookup_flags' of type 'ResourceLookupFlags' not supported

// UNSUPPORTED : C value 'g_resources_get_info' : parameter 'lookup_flags' of type 'ResourceLookupFlags' not supported

// UNSUPPORTED : C value 'g_resources_lookup_data' : parameter 'lookup_flags' of type 'ResourceLookupFlags' not supported

// UNSUPPORTED : C value 'g_resources_open_stream' : parameter 'lookup_flags' of type 'ResourceLookupFlags' not supported

var resourcesRegisterFunction *gi.Function
var resourcesRegisterFunction_Once sync.Once

func resourcesRegisterFunction_Set() error {
	var err error
	resourcesRegisterFunction_Once.Do(func() {
		resourcesRegisterFunction, err = gi.FunctionInvokerNew("Gio", "resources_register")
	})
	return err
}

// ResourcesRegister is a representation of the C type g_resources_register.
func ResourcesRegister(resource *Resource) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(resource.native)

	err := resourcesRegisterFunction_Set()
	if err == nil {
		resourcesRegisterFunction.Invoke(inArgs[:], nil)
	}

	return
}

var resourcesUnregisterFunction *gi.Function
var resourcesUnregisterFunction_Once sync.Once

func resourcesUnregisterFunction_Set() error {
	var err error
	resourcesUnregisterFunction_Once.Do(func() {
		resourcesUnregisterFunction, err = gi.FunctionInvokerNew("Gio", "resources_unregister")
	})
	return err
}

// ResourcesUnregister is a representation of the C type g_resources_unregister.
func ResourcesUnregister(resource *Resource) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(resource.native)

	err := resourcesUnregisterFunction_Set()
	if err == nil {
		resourcesUnregisterFunction.Invoke(inArgs[:], nil)
	}

	return
}

var settingsSchemaSourceGetDefaultFunction *gi.Function
var settingsSchemaSourceGetDefaultFunction_Once sync.Once

func settingsSchemaSourceGetDefaultFunction_Set() error {
	var err error
	settingsSchemaSourceGetDefaultFunction_Once.Do(func() {
		settingsSchemaSourceGetDefaultFunction, err = gi.FunctionInvokerNew("Gio", "settings_schema_source_get_default")
	})
	return err
}

// SettingsSchemaSourceGetDefault is a representation of the C type g_settings_schema_source_get_default.
func SettingsSchemaSourceGetDefault() *SettingsSchemaSource {

	var ret gi.Argument

	err := settingsSchemaSourceGetDefaultFunction_Set()
	if err == nil {
		ret = settingsSchemaSourceGetDefaultFunction.Invoke(nil, nil)
	}

	retGo := &SettingsSchemaSource{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_simple_async_report_error_in_idle' : parameter 'object' of type 'GObject.Object' not supported

// UNSUPPORTED : C value 'g_simple_async_report_gerror_in_idle' : parameter 'object' of type 'GObject.Object' not supported

// UNSUPPORTED : C value 'g_simple_async_report_take_gerror_in_idle' : parameter 'object' of type 'GObject.Object' not supported

// UNSUPPORTED : C value 'g_srv_target_list_sort' : parameter 'targets' of type 'GLib.List' not supported

// UNSUPPORTED : C value 'g_tls_backend_get_default' : return type 'TlsBackend' not supported

// UNSUPPORTED : C value 'g_tls_client_connection_new' : parameter 'server_identity' of type 'SocketConnectable' not supported

var tlsErrorQuarkFunction *gi.Function
var tlsErrorQuarkFunction_Once sync.Once

func tlsErrorQuarkFunction_Set() error {
	var err error
	tlsErrorQuarkFunction_Once.Do(func() {
		tlsErrorQuarkFunction, err = gi.FunctionInvokerNew("Gio", "tls_error_quark")
	})
	return err
}

// TlsErrorQuark is a representation of the C type g_tls_error_quark.
func TlsErrorQuark() glib.Quark {

	var ret gi.Argument

	err := tlsErrorQuarkFunction_Set()
	if err == nil {
		ret = tlsErrorQuarkFunction.Invoke(nil, nil)
	}

	retGo := glib.Quark(ret.Uint32())

	return retGo
}

// UNSUPPORTED : C value 'g_tls_file_database_new' : return type 'TlsFileDatabase' not supported

// UNSUPPORTED : C value 'g_tls_server_connection_new' : return type 'TlsServerConnection' not supported

var unixIsMountPathSystemInternalFunction *gi.Function
var unixIsMountPathSystemInternalFunction_Once sync.Once

func unixIsMountPathSystemInternalFunction_Set() error {
	var err error
	unixIsMountPathSystemInternalFunction_Once.Do(func() {
		unixIsMountPathSystemInternalFunction, err = gi.FunctionInvokerNew("Gio", "unix_is_mount_path_system_internal")
	})
	return err
}

// UnixIsMountPathSystemInternal is a representation of the C type g_unix_is_mount_path_system_internal.
func UnixIsMountPathSystemInternal(mountPath string) bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(mountPath)

	var ret gi.Argument

	err := unixIsMountPathSystemInternalFunction_Set()
	if err == nil {
		ret = unixIsMountPathSystemInternalFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var unixIsSystemDevicePathFunction *gi.Function
var unixIsSystemDevicePathFunction_Once sync.Once

func unixIsSystemDevicePathFunction_Set() error {
	var err error
	unixIsSystemDevicePathFunction_Once.Do(func() {
		unixIsSystemDevicePathFunction, err = gi.FunctionInvokerNew("Gio", "unix_is_system_device_path")
	})
	return err
}

// UnixIsSystemDevicePath is a representation of the C type g_unix_is_system_device_path.
func UnixIsSystemDevicePath(devicePath string) bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(devicePath)

	var ret gi.Argument

	err := unixIsSystemDevicePathFunction_Set()
	if err == nil {
		ret = unixIsSystemDevicePathFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var unixIsSystemFsTypeFunction *gi.Function
var unixIsSystemFsTypeFunction_Once sync.Once

func unixIsSystemFsTypeFunction_Set() error {
	var err error
	unixIsSystemFsTypeFunction_Once.Do(func() {
		unixIsSystemFsTypeFunction, err = gi.FunctionInvokerNew("Gio", "unix_is_system_fs_type")
	})
	return err
}

// UnixIsSystemFsType is a representation of the C type g_unix_is_system_fs_type.
func UnixIsSystemFsType(fsType string) bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(fsType)

	var ret gi.Argument

	err := unixIsSystemFsTypeFunction_Set()
	if err == nil {
		ret = unixIsSystemFsTypeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var unixMountAtFunction *gi.Function
var unixMountAtFunction_Once sync.Once

func unixMountAtFunction_Set() error {
	var err error
	unixMountAtFunction_Once.Do(func() {
		unixMountAtFunction, err = gi.FunctionInvokerNew("Gio", "unix_mount_at")
	})
	return err
}

// UnixMountAt is a representation of the C type g_unix_mount_at.
func UnixMountAt(mountPath string) (*UnixMountEntry, uint64) {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(mountPath)

	var outArgs [1]gi.Argument
	var ret gi.Argument

	err := unixMountAtFunction_Set()
	if err == nil {
		ret = unixMountAtFunction.Invoke(inArgs[:], outArgs[:])
	}

	retGo := &UnixMountEntry{native: ret.Pointer()}
	out0 := outArgs[0].Uint64()

	return retGo, out0
}

var unixMountCompareFunction *gi.Function
var unixMountCompareFunction_Once sync.Once

func unixMountCompareFunction_Set() error {
	var err error
	unixMountCompareFunction_Once.Do(func() {
		unixMountCompareFunction, err = gi.FunctionInvokerNew("Gio", "unix_mount_compare")
	})
	return err
}

// UnixMountCompare is a representation of the C type g_unix_mount_compare.
func UnixMountCompare(mount1 *UnixMountEntry, mount2 *UnixMountEntry) int32 {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(mount1.native)
	inArgs[1].SetPointer(mount2.native)

	var ret gi.Argument

	err := unixMountCompareFunction_Set()
	if err == nil {
		ret = unixMountCompareFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var unixMountCopyFunction *gi.Function
var unixMountCopyFunction_Once sync.Once

func unixMountCopyFunction_Set() error {
	var err error
	unixMountCopyFunction_Once.Do(func() {
		unixMountCopyFunction, err = gi.FunctionInvokerNew("Gio", "unix_mount_copy")
	})
	return err
}

// UnixMountCopy is a representation of the C type g_unix_mount_copy.
func UnixMountCopy(mountEntry *UnixMountEntry) *UnixMountEntry {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(mountEntry.native)

	var ret gi.Argument

	err := unixMountCopyFunction_Set()
	if err == nil {
		ret = unixMountCopyFunction.Invoke(inArgs[:], nil)
	}

	retGo := &UnixMountEntry{native: ret.Pointer()}

	return retGo
}

var unixMountForFunction *gi.Function
var unixMountForFunction_Once sync.Once

func unixMountForFunction_Set() error {
	var err error
	unixMountForFunction_Once.Do(func() {
		unixMountForFunction, err = gi.FunctionInvokerNew("Gio", "unix_mount_for")
	})
	return err
}

// UnixMountFor is a representation of the C type g_unix_mount_for.
func UnixMountFor(filePath string) (*UnixMountEntry, uint64) {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(filePath)

	var outArgs [1]gi.Argument
	var ret gi.Argument

	err := unixMountForFunction_Set()
	if err == nil {
		ret = unixMountForFunction.Invoke(inArgs[:], outArgs[:])
	}

	retGo := &UnixMountEntry{native: ret.Pointer()}
	out0 := outArgs[0].Uint64()

	return retGo, out0
}

var unixMountFreeFunction *gi.Function
var unixMountFreeFunction_Once sync.Once

func unixMountFreeFunction_Set() error {
	var err error
	unixMountFreeFunction_Once.Do(func() {
		unixMountFreeFunction, err = gi.FunctionInvokerNew("Gio", "unix_mount_free")
	})
	return err
}

// UnixMountFree is a representation of the C type g_unix_mount_free.
func UnixMountFree(mountEntry *UnixMountEntry) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(mountEntry.native)

	err := unixMountFreeFunction_Set()
	if err == nil {
		unixMountFreeFunction.Invoke(inArgs[:], nil)
	}

	return
}

var unixMountGetDevicePathFunction *gi.Function
var unixMountGetDevicePathFunction_Once sync.Once

func unixMountGetDevicePathFunction_Set() error {
	var err error
	unixMountGetDevicePathFunction_Once.Do(func() {
		unixMountGetDevicePathFunction, err = gi.FunctionInvokerNew("Gio", "unix_mount_get_device_path")
	})
	return err
}

// UnixMountGetDevicePath is a representation of the C type g_unix_mount_get_device_path.
func UnixMountGetDevicePath(mountEntry *UnixMountEntry) string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(mountEntry.native)

	var ret gi.Argument

	err := unixMountGetDevicePathFunction_Set()
	if err == nil {
		ret = unixMountGetDevicePathFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var unixMountGetFsTypeFunction *gi.Function
var unixMountGetFsTypeFunction_Once sync.Once

func unixMountGetFsTypeFunction_Set() error {
	var err error
	unixMountGetFsTypeFunction_Once.Do(func() {
		unixMountGetFsTypeFunction, err = gi.FunctionInvokerNew("Gio", "unix_mount_get_fs_type")
	})
	return err
}

// UnixMountGetFsType is a representation of the C type g_unix_mount_get_fs_type.
func UnixMountGetFsType(mountEntry *UnixMountEntry) string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(mountEntry.native)

	var ret gi.Argument

	err := unixMountGetFsTypeFunction_Set()
	if err == nil {
		ret = unixMountGetFsTypeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var unixMountGetMountPathFunction *gi.Function
var unixMountGetMountPathFunction_Once sync.Once

func unixMountGetMountPathFunction_Set() error {
	var err error
	unixMountGetMountPathFunction_Once.Do(func() {
		unixMountGetMountPathFunction, err = gi.FunctionInvokerNew("Gio", "unix_mount_get_mount_path")
	})
	return err
}

// UnixMountGetMountPath is a representation of the C type g_unix_mount_get_mount_path.
func UnixMountGetMountPath(mountEntry *UnixMountEntry) string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(mountEntry.native)

	var ret gi.Argument

	err := unixMountGetMountPathFunction_Set()
	if err == nil {
		ret = unixMountGetMountPathFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var unixMountGetOptionsFunction *gi.Function
var unixMountGetOptionsFunction_Once sync.Once

func unixMountGetOptionsFunction_Set() error {
	var err error
	unixMountGetOptionsFunction_Once.Do(func() {
		unixMountGetOptionsFunction, err = gi.FunctionInvokerNew("Gio", "unix_mount_get_options")
	})
	return err
}

// UnixMountGetOptions is a representation of the C type g_unix_mount_get_options.
func UnixMountGetOptions(mountEntry *UnixMountEntry) string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(mountEntry.native)

	var ret gi.Argument

	err := unixMountGetOptionsFunction_Set()
	if err == nil {
		ret = unixMountGetOptionsFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var unixMountGetRootPathFunction *gi.Function
var unixMountGetRootPathFunction_Once sync.Once

func unixMountGetRootPathFunction_Set() error {
	var err error
	unixMountGetRootPathFunction_Once.Do(func() {
		unixMountGetRootPathFunction, err = gi.FunctionInvokerNew("Gio", "unix_mount_get_root_path")
	})
	return err
}

// UnixMountGetRootPath is a representation of the C type g_unix_mount_get_root_path.
func UnixMountGetRootPath(mountEntry *UnixMountEntry) string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(mountEntry.native)

	var ret gi.Argument

	err := unixMountGetRootPathFunction_Set()
	if err == nil {
		ret = unixMountGetRootPathFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var unixMountGuessCanEjectFunction *gi.Function
var unixMountGuessCanEjectFunction_Once sync.Once

func unixMountGuessCanEjectFunction_Set() error {
	var err error
	unixMountGuessCanEjectFunction_Once.Do(func() {
		unixMountGuessCanEjectFunction, err = gi.FunctionInvokerNew("Gio", "unix_mount_guess_can_eject")
	})
	return err
}

// UnixMountGuessCanEject is a representation of the C type g_unix_mount_guess_can_eject.
func UnixMountGuessCanEject(mountEntry *UnixMountEntry) bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(mountEntry.native)

	var ret gi.Argument

	err := unixMountGuessCanEjectFunction_Set()
	if err == nil {
		ret = unixMountGuessCanEjectFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'g_unix_mount_guess_icon' : return type 'Icon' not supported

var unixMountGuessNameFunction *gi.Function
var unixMountGuessNameFunction_Once sync.Once

func unixMountGuessNameFunction_Set() error {
	var err error
	unixMountGuessNameFunction_Once.Do(func() {
		unixMountGuessNameFunction, err = gi.FunctionInvokerNew("Gio", "unix_mount_guess_name")
	})
	return err
}

// UnixMountGuessName is a representation of the C type g_unix_mount_guess_name.
func UnixMountGuessName(mountEntry *UnixMountEntry) string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(mountEntry.native)

	var ret gi.Argument

	err := unixMountGuessNameFunction_Set()
	if err == nil {
		ret = unixMountGuessNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

var unixMountGuessShouldDisplayFunction *gi.Function
var unixMountGuessShouldDisplayFunction_Once sync.Once

func unixMountGuessShouldDisplayFunction_Set() error {
	var err error
	unixMountGuessShouldDisplayFunction_Once.Do(func() {
		unixMountGuessShouldDisplayFunction, err = gi.FunctionInvokerNew("Gio", "unix_mount_guess_should_display")
	})
	return err
}

// UnixMountGuessShouldDisplay is a representation of the C type g_unix_mount_guess_should_display.
func UnixMountGuessShouldDisplay(mountEntry *UnixMountEntry) bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(mountEntry.native)

	var ret gi.Argument

	err := unixMountGuessShouldDisplayFunction_Set()
	if err == nil {
		ret = unixMountGuessShouldDisplayFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'g_unix_mount_guess_symbolic_icon' : return type 'Icon' not supported

var unixMountIsReadonlyFunction *gi.Function
var unixMountIsReadonlyFunction_Once sync.Once

func unixMountIsReadonlyFunction_Set() error {
	var err error
	unixMountIsReadonlyFunction_Once.Do(func() {
		unixMountIsReadonlyFunction, err = gi.FunctionInvokerNew("Gio", "unix_mount_is_readonly")
	})
	return err
}

// UnixMountIsReadonly is a representation of the C type g_unix_mount_is_readonly.
func UnixMountIsReadonly(mountEntry *UnixMountEntry) bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(mountEntry.native)

	var ret gi.Argument

	err := unixMountIsReadonlyFunction_Set()
	if err == nil {
		ret = unixMountIsReadonlyFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var unixMountIsSystemInternalFunction *gi.Function
var unixMountIsSystemInternalFunction_Once sync.Once

func unixMountIsSystemInternalFunction_Set() error {
	var err error
	unixMountIsSystemInternalFunction_Once.Do(func() {
		unixMountIsSystemInternalFunction, err = gi.FunctionInvokerNew("Gio", "unix_mount_is_system_internal")
	})
	return err
}

// UnixMountIsSystemInternal is a representation of the C type g_unix_mount_is_system_internal.
func UnixMountIsSystemInternal(mountEntry *UnixMountEntry) bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(mountEntry.native)

	var ret gi.Argument

	err := unixMountIsSystemInternalFunction_Set()
	if err == nil {
		ret = unixMountIsSystemInternalFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var unixMountPointsChangedSinceFunction *gi.Function
var unixMountPointsChangedSinceFunction_Once sync.Once

func unixMountPointsChangedSinceFunction_Set() error {
	var err error
	unixMountPointsChangedSinceFunction_Once.Do(func() {
		unixMountPointsChangedSinceFunction, err = gi.FunctionInvokerNew("Gio", "unix_mount_points_changed_since")
	})
	return err
}

// UnixMountPointsChangedSince is a representation of the C type g_unix_mount_points_changed_since.
func UnixMountPointsChangedSince(time uint64) bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetUint64(time)

	var ret gi.Argument

	err := unixMountPointsChangedSinceFunction_Set()
	if err == nil {
		ret = unixMountPointsChangedSinceFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'g_unix_mount_points_get' : return type 'GLib.List' not supported

var unixMountsChangedSinceFunction *gi.Function
var unixMountsChangedSinceFunction_Once sync.Once

func unixMountsChangedSinceFunction_Set() error {
	var err error
	unixMountsChangedSinceFunction_Once.Do(func() {
		unixMountsChangedSinceFunction, err = gi.FunctionInvokerNew("Gio", "unix_mounts_changed_since")
	})
	return err
}

// UnixMountsChangedSince is a representation of the C type g_unix_mounts_changed_since.
func UnixMountsChangedSince(time uint64) bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetUint64(time)

	var ret gi.Argument

	err := unixMountsChangedSinceFunction_Set()
	if err == nil {
		ret = unixMountsChangedSinceFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'g_unix_mounts_get' : return type 'GLib.List' not supported
