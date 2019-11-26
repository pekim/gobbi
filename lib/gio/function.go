// Code generated - DO NOT EDIT.

package gio

import (
	gi "github.com/pekim/gobbi/internal/gi"
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
func ActionNameIsValid(actionName string) (bool, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(actionName)

	var ret gi.Argument

	err := actionNameIsValidFunction_Set()
	if err == nil {
		ret = actionNameIsValidFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

// UNSUPPORTED : C value 'g_action_parse_detailed_name' : parameter 'target_value' of type 'GLib.Variant' not supported

// UNSUPPORTED : C value 'g_action_print_detailed_name' : parameter 'target_value' of type 'GLib.Variant' not supported

// UNSUPPORTED : C value 'g_app_info_create_from_commandline' : parameter 'commandline' of type 'filename' not supported

// UNSUPPORTED : C value 'g_app_info_get_all' : return type 'GLib.List' not supported

// UNSUPPORTED : C value 'g_app_info_get_all_for_type' : return type 'GLib.List' not supported

// UNSUPPORTED : C value 'g_app_info_get_default_for_type' : return type 'AppInfo' not supported

// UNSUPPORTED : C value 'g_app_info_get_default_for_uri_scheme' : return type 'AppInfo' not supported

// UNSUPPORTED : C value 'g_app_info_get_fallback_for_type' : return type 'GLib.List' not supported

// UNSUPPORTED : C value 'g_app_info_get_recommended_for_type' : return type 'GLib.List' not supported

// UNSUPPORTED : C value 'g_app_info_launch_default_for_uri' : parameter 'context' of type 'AppLaunchContext' not supported

// UNSUPPORTED : C value 'g_app_info_launch_default_for_uri_async' : parameter 'context' of type 'AppLaunchContext' not supported

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
func AppInfoResetTypeAssociations(contentType string) error {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(contentType)

	err := appInfoResetTypeAssociationsFunction_Set()
	if err == nil {
		appInfoResetTypeAssociationsFunction.Invoke(inArgs[:], nil)
	}

	return err
}

// UNSUPPORTED : C value 'g_async_initable_newv_async' : parameter 'object_type' of type 'GType' not supported

// UNSUPPORTED : C value 'g_bus_get' : parameter 'bus_type' of type 'BusType' not supported

// UNSUPPORTED : C value 'g_bus_get_finish' : parameter 'res' of type 'AsyncResult' not supported

// UNSUPPORTED : C value 'g_bus_get_sync' : parameter 'bus_type' of type 'BusType' not supported

// UNSUPPORTED : C value 'g_bus_own_name' : parameter 'bus_type' of type 'BusType' not supported

// UNSUPPORTED : C value 'g_bus_own_name_on_connection' : parameter 'connection' of type 'DBusConnection' not supported

// UNSUPPORTED : C value 'g_bus_own_name_on_connection_with_closures' : parameter 'connection' of type 'DBusConnection' not supported

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
func BusUnownName(ownerId uint32) error {
	var inArgs [1]gi.Argument
	inArgs[0].SetUint32(ownerId)

	err := busUnownNameFunction_Set()
	if err == nil {
		busUnownNameFunction.Invoke(inArgs[:], nil)
	}

	return err
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
func BusUnwatchName(watcherId uint32) error {
	var inArgs [1]gi.Argument
	inArgs[0].SetUint32(watcherId)

	err := busUnwatchNameFunction_Set()
	if err == nil {
		busUnwatchNameFunction.Invoke(inArgs[:], nil)
	}

	return err
}

// UNSUPPORTED : C value 'g_bus_watch_name' : parameter 'bus_type' of type 'BusType' not supported

// UNSUPPORTED : C value 'g_bus_watch_name_on_connection' : parameter 'connection' of type 'DBusConnection' not supported

// UNSUPPORTED : C value 'g_bus_watch_name_on_connection_with_closures' : parameter 'connection' of type 'DBusConnection' not supported

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
func ContentTypeCanBeExecutable(type_ string) (bool, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(type_)

	var ret gi.Argument

	err := contentTypeCanBeExecutableFunction_Set()
	if err == nil {
		ret = contentTypeCanBeExecutableFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
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
func ContentTypeEquals(type1 string, type2 string) (bool, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(type1)
	inArgs[1].SetString(type2)

	var ret gi.Argument

	err := contentTypeEqualsFunction_Set()
	if err == nil {
		ret = contentTypeEqualsFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
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
func ContentTypeFromMimeType(mimeType string) (string, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(mimeType)

	var ret gi.Argument

	err := contentTypeFromMimeTypeFunction_Set()
	if err == nil {
		ret = contentTypeFromMimeTypeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo, err
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
func ContentTypeGetDescription(type_ string) (string, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(type_)

	var ret gi.Argument

	err := contentTypeGetDescriptionFunction_Set()
	if err == nil {
		ret = contentTypeGetDescriptionFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo, err
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
func ContentTypeGetGenericIconName(type_ string) (string, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(type_)

	var ret gi.Argument

	err := contentTypeGetGenericIconNameFunction_Set()
	if err == nil {
		ret = contentTypeGetGenericIconNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo, err
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
func ContentTypeGetMimeDirs() error {

	err := contentTypeGetMimeDirsFunction_Set()
	if err == nil {
		contentTypeGetMimeDirsFunction.Invoke(nil, nil)
	}

	return err
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
func ContentTypeGetMimeType(type_ string) (string, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(type_)

	var ret gi.Argument

	err := contentTypeGetMimeTypeFunction_Set()
	if err == nil {
		ret = contentTypeGetMimeTypeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo, err
}

// UNSUPPORTED : C value 'g_content_type_get_symbolic_icon' : return type 'Icon' not supported

// UNSUPPORTED : C value 'g_content_type_guess' : parameter 'data' has no type

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
func ContentTypeIsA(type_ string, supertype string) (bool, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(type_)
	inArgs[1].SetString(supertype)

	var ret gi.Argument

	err := contentTypeIsAFunction_Set()
	if err == nil {
		ret = contentTypeIsAFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
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
func ContentTypeIsMimeType(type_ string, mimeType string) (bool, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(type_)
	inArgs[1].SetString(mimeType)

	var ret gi.Argument

	err := contentTypeIsMimeTypeFunction_Set()
	if err == nil {
		ret = contentTypeIsMimeTypeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
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
func ContentTypeIsUnknown(type_ string) (bool, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(type_)

	var ret gi.Argument

	err := contentTypeIsUnknownFunction_Set()
	if err == nil {
		ret = contentTypeIsUnknownFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

// UNSUPPORTED : C value 'g_content_type_set_mime_dirs' : parameter 'dirs' has no type

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
func DbusAddressEscapeValue(string_ string) (string, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(string_)

	var ret gi.Argument

	err := dbusAddressEscapeValueFunction_Set()
	if err == nil {
		ret = dbusAddressEscapeValueFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo, err
}

// UNSUPPORTED : C value 'g_dbus_address_get_for_bus_sync' : parameter 'bus_type' of type 'BusType' not supported

// UNSUPPORTED : C value 'g_dbus_address_get_stream' : parameter 'cancellable' of type 'Cancellable' not supported

// UNSUPPORTED : C value 'g_dbus_address_get_stream_finish' : parameter 'res' of type 'AsyncResult' not supported

// UNSUPPORTED : C value 'g_dbus_address_get_stream_sync' : parameter 'cancellable' of type 'Cancellable' not supported

// UNSUPPORTED : C value 'g_dbus_annotation_info_lookup' : parameter 'annotations' has no type

// UNSUPPORTED : C value 'g_dbus_error_encode_gerror' : parameter 'error' of type 'GLib.Error' not supported

// UNSUPPORTED : C value 'g_dbus_error_get_remote_error' : parameter 'error' of type 'GLib.Error' not supported

// UNSUPPORTED : C value 'g_dbus_error_is_remote_error' : parameter 'error' of type 'GLib.Error' not supported

// UNSUPPORTED : C value 'g_dbus_error_new_for_dbus_error' : return type 'GLib.Error' not supported

// UNSUPPORTED : C value 'g_dbus_error_quark' : return type 'GLib.Quark' not supported

// UNSUPPORTED : C value 'g_dbus_error_register_error' : parameter 'error_domain' of type 'GLib.Quark' not supported

// UNSUPPORTED : C value 'g_dbus_error_register_error_domain' : parameter 'entries' has no type

// UNSUPPORTED : C value 'g_dbus_error_strip_remote_error' : parameter 'error' of type 'GLib.Error' not supported

// UNSUPPORTED : C value 'g_dbus_error_unregister_error' : parameter 'error_domain' of type 'GLib.Quark' not supported

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
func DbusGenerateGuid() (string, error) {

	var ret gi.Argument

	err := dbusGenerateGuidFunction_Set()
	if err == nil {
		ret = dbusGenerateGuidFunction.Invoke(nil, nil)
	}

	retGo := ret.String(true)

	return retGo, err
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
func DbusIsAddress(string_ string) (bool, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(string_)

	var ret gi.Argument

	err := dbusIsAddressFunction_Set()
	if err == nil {
		ret = dbusIsAddressFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
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
func DbusIsGuid(string_ string) (bool, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(string_)

	var ret gi.Argument

	err := dbusIsGuidFunction_Set()
	if err == nil {
		ret = dbusIsGuidFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
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
func DbusIsInterfaceName(string_ string) (bool, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(string_)

	var ret gi.Argument

	err := dbusIsInterfaceNameFunction_Set()
	if err == nil {
		ret = dbusIsInterfaceNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
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
func DbusIsMemberName(string_ string) (bool, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(string_)

	var ret gi.Argument

	err := dbusIsMemberNameFunction_Set()
	if err == nil {
		ret = dbusIsMemberNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
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
func DbusIsName(string_ string) (bool, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(string_)

	var ret gi.Argument

	err := dbusIsNameFunction_Set()
	if err == nil {
		ret = dbusIsNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
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
func DbusIsSupportedAddress(string_ string) (bool, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(string_)

	var ret gi.Argument

	err := dbusIsSupportedAddressFunction_Set()
	if err == nil {
		ret = dbusIsSupportedAddressFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
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
func DbusIsUniqueName(string_ string) (bool, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(string_)

	var ret gi.Argument

	err := dbusIsUniqueNameFunction_Set()
	if err == nil {
		ret = dbusIsUniqueNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

// UNSUPPORTED : C value 'g_dtls_client_connection_new' : parameter 'base_socket' of type 'DatagramBased' not supported

// UNSUPPORTED : C value 'g_dtls_server_connection_new' : parameter 'base_socket' of type 'DatagramBased' not supported

// UNSUPPORTED : C value 'g_file_new_for_commandline_arg' : parameter 'arg' of type 'filename' not supported

// UNSUPPORTED : C value 'g_file_new_for_commandline_arg_and_cwd' : parameter 'arg' of type 'filename' not supported

// UNSUPPORTED : C value 'g_file_new_for_path' : parameter 'path' of type 'filename' not supported

// UNSUPPORTED : C value 'g_file_new_for_uri' : return type 'File' not supported

// UNSUPPORTED : C value 'g_file_new_tmp' : parameter 'tmpl' of type 'filename' not supported

// UNSUPPORTED : C value 'g_file_parse_name' : return type 'File' not supported

// UNSUPPORTED : C value 'g_icon_deserialize' : parameter 'value' of type 'GLib.Variant' not supported

// UNSUPPORTED : C value 'g_icon_hash' : parameter 'icon' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_icon_new_for_string' : return type 'Icon' not supported

// UNSUPPORTED : C value 'g_initable_newv' : parameter 'object_type' of type 'GType' not supported

// UNSUPPORTED : C value 'g_io_error_from_errno' : return type 'IOErrorEnum' not supported

// UNSUPPORTED : C value 'g_io_error_quark' : return type 'GLib.Quark' not supported

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
func IoExtensionPointLookup(name string) (*IOExtensionPoint, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(name)

	var ret gi.Argument

	err := ioExtensionPointLookupFunction_Set()
	if err == nil {
		ret = ioExtensionPointLookupFunction.Invoke(inArgs[:], nil)
	}

	retGo := &IOExtensionPoint{native: ret.Pointer()}

	return retGo, err
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
func IoExtensionPointRegister(name string) (*IOExtensionPoint, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(name)

	var ret gi.Argument

	err := ioExtensionPointRegisterFunction_Set()
	if err == nil {
		ret = ioExtensionPointRegisterFunction.Invoke(inArgs[:], nil)
	}

	retGo := &IOExtensionPoint{native: ret.Pointer()}

	return retGo, err
}

// UNSUPPORTED : C value 'g_io_modules_load_all_in_directory' : parameter 'dirname' of type 'filename' not supported

// UNSUPPORTED : C value 'g_io_modules_load_all_in_directory_with_scope' : parameter 'dirname' of type 'filename' not supported

// UNSUPPORTED : C value 'g_io_modules_scan_all_in_directory' : parameter 'dirname' of type 'filename' not supported

// UNSUPPORTED : C value 'g_io_modules_scan_all_in_directory_with_scope' : parameter 'dirname' of type 'filename' not supported

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
func IoSchedulerCancelAllJobs() error {

	err := ioSchedulerCancelAllJobsFunction_Set()
	if err == nil {
		ioSchedulerCancelAllJobsFunction.Invoke(nil, nil)
	}

	return err
}

// UNSUPPORTED : C value 'g_io_scheduler_push_job' : parameter 'job_func' of type 'IOSchedulerJobFunc' not supported

// UNSUPPORTED : C value 'g_keyfile_settings_backend_new' : return type 'SettingsBackend' not supported

// UNSUPPORTED : C value 'g_memory_settings_backend_new' : return type 'SettingsBackend' not supported

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
func NetworkingInit() error {

	err := networkingInitFunction_Set()
	if err == nil {
		networkingInitFunction.Invoke(nil, nil)
	}

	return err
}

// UNSUPPORTED : C value 'g_null_settings_backend_new' : return type 'SettingsBackend' not supported

// UNSUPPORTED : C value 'g_pollable_source_new' : parameter 'pollable_stream' of type 'GObject.Object' not supported

// UNSUPPORTED : C value 'g_pollable_source_new_full' : parameter 'pollable_stream' of type 'GObject.Object' not supported

// UNSUPPORTED : C value 'g_pollable_stream_read' : parameter 'stream' of type 'InputStream' not supported

// UNSUPPORTED : C value 'g_pollable_stream_write' : parameter 'stream' of type 'OutputStream' not supported

// UNSUPPORTED : C value 'g_pollable_stream_write_all' : parameter 'stream' of type 'OutputStream' not supported

// UNSUPPORTED : C value 'g_proxy_get_default_for_protocol' : return type 'Proxy' not supported

// UNSUPPORTED : C value 'g_proxy_resolver_get_default' : return type 'ProxyResolver' not supported

// UNSUPPORTED : C value 'g_resolver_error_quark' : return type 'GLib.Quark' not supported

// UNSUPPORTED : C value 'g_resource_error_quark' : return type 'GLib.Quark' not supported

// UNSUPPORTED : C value 'g_resource_load' : parameter 'filename' of type 'filename' not supported

// UNSUPPORTED : C value 'g_resources_enumerate_children' : parameter 'lookup_flags' of type 'ResourceLookupFlags' not supported

// UNSUPPORTED : C value 'g_resources_get_info' : parameter 'lookup_flags' of type 'ResourceLookupFlags' not supported

// UNSUPPORTED : C value 'g_resources_lookup_data' : parameter 'lookup_flags' of type 'ResourceLookupFlags' not supported

// UNSUPPORTED : C value 'g_resources_open_stream' : parameter 'lookup_flags' of type 'ResourceLookupFlags' not supported

// UNSUPPORTED : C value 'g_resources_register' : parameter 'resource' of type 'Resource' not supported

// UNSUPPORTED : C value 'g_resources_unregister' : parameter 'resource' of type 'Resource' not supported

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
func SettingsSchemaSourceGetDefault() (*SettingsSchemaSource, error) {

	var ret gi.Argument

	err := settingsSchemaSourceGetDefaultFunction_Set()
	if err == nil {
		ret = settingsSchemaSourceGetDefaultFunction.Invoke(nil, nil)
	}

	retGo := &SettingsSchemaSource{native: ret.Pointer()}

	return retGo, err
}

// UNSUPPORTED : C value 'g_simple_async_report_error_in_idle' : parameter 'object' of type 'GObject.Object' not supported

// UNSUPPORTED : C value 'g_simple_async_report_gerror_in_idle' : parameter 'object' of type 'GObject.Object' not supported

// UNSUPPORTED : C value 'g_simple_async_report_take_gerror_in_idle' : parameter 'object' of type 'GObject.Object' not supported

// UNSUPPORTED : C value 'g_srv_target_list_sort' : parameter 'targets' of type 'GLib.List' not supported

// UNSUPPORTED : C value 'g_tls_backend_get_default' : return type 'TlsBackend' not supported

// UNSUPPORTED : C value 'g_tls_client_connection_new' : parameter 'base_io_stream' of type 'IOStream' not supported

// UNSUPPORTED : C value 'g_tls_error_quark' : return type 'GLib.Quark' not supported

// UNSUPPORTED : C value 'g_tls_file_database_new' : parameter 'anchors' of type 'filename' not supported

// UNSUPPORTED : C value 'g_tls_server_connection_new' : parameter 'base_io_stream' of type 'IOStream' not supported

// UNSUPPORTED : C value 'g_unix_is_mount_path_system_internal' : parameter 'mount_path' of type 'filename' not supported

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
func UnixIsSystemDevicePath(devicePath string) (bool, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(devicePath)

	var ret gi.Argument

	err := unixIsSystemDevicePathFunction_Set()
	if err == nil {
		ret = unixIsSystemDevicePathFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
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
func UnixIsSystemFsType(fsType string) (bool, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(fsType)

	var ret gi.Argument

	err := unixIsSystemFsTypeFunction_Set()
	if err == nil {
		ret = unixIsSystemFsTypeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

// UNSUPPORTED : C value 'g_unix_mount_at' : parameter 'mount_path' of type 'filename' not supported

// UNSUPPORTED : C value 'g_unix_mount_compare' : parameter 'mount1' of type 'UnixMountEntry' not supported

// UNSUPPORTED : C value 'g_unix_mount_copy' : parameter 'mount_entry' of type 'UnixMountEntry' not supported

// UNSUPPORTED : C value 'g_unix_mount_for' : parameter 'file_path' of type 'filename' not supported

// UNSUPPORTED : C value 'g_unix_mount_free' : parameter 'mount_entry' of type 'UnixMountEntry' not supported

// UNSUPPORTED : C value 'g_unix_mount_get_device_path' : parameter 'mount_entry' of type 'UnixMountEntry' not supported

// UNSUPPORTED : C value 'g_unix_mount_get_fs_type' : parameter 'mount_entry' of type 'UnixMountEntry' not supported

// UNSUPPORTED : C value 'g_unix_mount_get_mount_path' : parameter 'mount_entry' of type 'UnixMountEntry' not supported

// UNSUPPORTED : C value 'g_unix_mount_get_options' : parameter 'mount_entry' of type 'UnixMountEntry' not supported

// UNSUPPORTED : C value 'g_unix_mount_get_root_path' : parameter 'mount_entry' of type 'UnixMountEntry' not supported

// UNSUPPORTED : C value 'g_unix_mount_guess_can_eject' : parameter 'mount_entry' of type 'UnixMountEntry' not supported

// UNSUPPORTED : C value 'g_unix_mount_guess_icon' : parameter 'mount_entry' of type 'UnixMountEntry' not supported

// UNSUPPORTED : C value 'g_unix_mount_guess_name' : parameter 'mount_entry' of type 'UnixMountEntry' not supported

// UNSUPPORTED : C value 'g_unix_mount_guess_should_display' : parameter 'mount_entry' of type 'UnixMountEntry' not supported

// UNSUPPORTED : C value 'g_unix_mount_guess_symbolic_icon' : parameter 'mount_entry' of type 'UnixMountEntry' not supported

// UNSUPPORTED : C value 'g_unix_mount_is_readonly' : parameter 'mount_entry' of type 'UnixMountEntry' not supported

// UNSUPPORTED : C value 'g_unix_mount_is_system_internal' : parameter 'mount_entry' of type 'UnixMountEntry' not supported

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
func UnixMountPointsChangedSince(time uint64) (bool, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetUint64(time)

	var ret gi.Argument

	err := unixMountPointsChangedSinceFunction_Set()
	if err == nil {
		ret = unixMountPointsChangedSinceFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
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
func UnixMountsChangedSince(time uint64) (bool, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetUint64(time)

	var ret gi.Argument

	err := unixMountsChangedSinceFunction_Set()
	if err == nil {
		ret = unixMountsChangedSinceFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

// UNSUPPORTED : C value 'g_unix_mounts_get' : return type 'GLib.List' not supported
