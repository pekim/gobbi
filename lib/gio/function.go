// Code generated - DO NOT EDIT.

package gio

import (
	gi "github.com/pekim/gobbi/internal/cgo/gi"
	glib "github.com/pekim/gobbi/lib/glib"
	gobject "github.com/pekim/gobbi/lib/gobject"
	"sync"
	"unsafe"
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

var actionParseDetailedNameFunction *gi.Function
var actionParseDetailedNameFunction_Once sync.Once

func actionParseDetailedNameFunction_Set() error {
	var err error
	actionParseDetailedNameFunction_Once.Do(func() {
		actionParseDetailedNameFunction, err = gi.FunctionInvokerNew("Gio", "action_parse_detailed_name")
	})
	return err
}

// ActionParseDetailedName is a representation of the C type g_action_parse_detailed_name.
func ActionParseDetailedName(detailedName string) (bool, string, *glib.Variant) {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(detailedName)

	var outArgs [2]gi.Argument
	var ret gi.Argument

	err := actionParseDetailedNameFunction_Set()
	if err == nil {
		ret = actionParseDetailedNameFunction.Invoke(inArgs[:], outArgs[:])
	}

	retGo := ret.Boolean()
	out0 := outArgs[0].String(true)
	out1 := glib.VariantNewFromNative(outArgs[1].Pointer())

	return retGo, out0, out1
}

var actionPrintDetailedNameFunction *gi.Function
var actionPrintDetailedNameFunction_Once sync.Once

func actionPrintDetailedNameFunction_Set() error {
	var err error
	actionPrintDetailedNameFunction_Once.Do(func() {
		actionPrintDetailedNameFunction, err = gi.FunctionInvokerNew("Gio", "action_print_detailed_name")
	})
	return err
}

// ActionPrintDetailedName is a representation of the C type g_action_print_detailed_name.
func ActionPrintDetailedName(actionName string, targetValue *glib.Variant) string {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(actionName)
	inArgs[1].SetPointer(targetValue.Native())

	var ret gi.Argument

	err := actionPrintDetailedNameFunction_Set()
	if err == nil {
		ret = actionPrintDetailedNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

var appInfoCreateFromCommandlineFunction *gi.Function
var appInfoCreateFromCommandlineFunction_Once sync.Once

func appInfoCreateFromCommandlineFunction_Set() error {
	var err error
	appInfoCreateFromCommandlineFunction_Once.Do(func() {
		appInfoCreateFromCommandlineFunction, err = gi.FunctionInvokerNew("Gio", "app_info_create_from_commandline")
	})
	return err
}

// AppInfoCreateFromCommandline is a representation of the C type g_app_info_create_from_commandline.
func AppInfoCreateFromCommandline(commandline string, applicationName string, flags AppInfoCreateFlags) *AppInfo {
	var inArgs [3]gi.Argument
	inArgs[0].SetString(commandline)
	inArgs[1].SetString(applicationName)
	inArgs[2].SetInt32(int32(flags))

	var ret gi.Argument

	err := appInfoCreateFromCommandlineFunction_Set()
	if err == nil {
		ret = appInfoCreateFromCommandlineFunction.Invoke(inArgs[:], nil)
	}

	retGo := AppInfoNewFromNative(ret.Pointer())

	return retGo
}

var appInfoGetAllFunction *gi.Function
var appInfoGetAllFunction_Once sync.Once

func appInfoGetAllFunction_Set() error {
	var err error
	appInfoGetAllFunction_Once.Do(func() {
		appInfoGetAllFunction, err = gi.FunctionInvokerNew("Gio", "app_info_get_all")
	})
	return err
}

// AppInfoGetAll is a representation of the C type g_app_info_get_all.
func AppInfoGetAll() *glib.List {

	var ret gi.Argument

	err := appInfoGetAllFunction_Set()
	if err == nil {
		ret = appInfoGetAllFunction.Invoke(nil, nil)
	}

	retGo := glib.ListNewFromNative(ret.Pointer())

	return retGo
}

var appInfoGetAllForTypeFunction *gi.Function
var appInfoGetAllForTypeFunction_Once sync.Once

func appInfoGetAllForTypeFunction_Set() error {
	var err error
	appInfoGetAllForTypeFunction_Once.Do(func() {
		appInfoGetAllForTypeFunction, err = gi.FunctionInvokerNew("Gio", "app_info_get_all_for_type")
	})
	return err
}

// AppInfoGetAllForType is a representation of the C type g_app_info_get_all_for_type.
func AppInfoGetAllForType(contentType string) *glib.List {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(contentType)

	var ret gi.Argument

	err := appInfoGetAllForTypeFunction_Set()
	if err == nil {
		ret = appInfoGetAllForTypeFunction.Invoke(inArgs[:], nil)
	}

	retGo := glib.ListNewFromNative(ret.Pointer())

	return retGo
}

var appInfoGetDefaultForTypeFunction *gi.Function
var appInfoGetDefaultForTypeFunction_Once sync.Once

func appInfoGetDefaultForTypeFunction_Set() error {
	var err error
	appInfoGetDefaultForTypeFunction_Once.Do(func() {
		appInfoGetDefaultForTypeFunction, err = gi.FunctionInvokerNew("Gio", "app_info_get_default_for_type")
	})
	return err
}

// AppInfoGetDefaultForType is a representation of the C type g_app_info_get_default_for_type.
func AppInfoGetDefaultForType(contentType string, mustSupportUris bool) *AppInfo {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(contentType)
	inArgs[1].SetBoolean(mustSupportUris)

	var ret gi.Argument

	err := appInfoGetDefaultForTypeFunction_Set()
	if err == nil {
		ret = appInfoGetDefaultForTypeFunction.Invoke(inArgs[:], nil)
	}

	retGo := AppInfoNewFromNative(ret.Pointer())

	return retGo
}

var appInfoGetDefaultForUriSchemeFunction *gi.Function
var appInfoGetDefaultForUriSchemeFunction_Once sync.Once

func appInfoGetDefaultForUriSchemeFunction_Set() error {
	var err error
	appInfoGetDefaultForUriSchemeFunction_Once.Do(func() {
		appInfoGetDefaultForUriSchemeFunction, err = gi.FunctionInvokerNew("Gio", "app_info_get_default_for_uri_scheme")
	})
	return err
}

// AppInfoGetDefaultForUriScheme is a representation of the C type g_app_info_get_default_for_uri_scheme.
func AppInfoGetDefaultForUriScheme(uriScheme string) *AppInfo {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(uriScheme)

	var ret gi.Argument

	err := appInfoGetDefaultForUriSchemeFunction_Set()
	if err == nil {
		ret = appInfoGetDefaultForUriSchemeFunction.Invoke(inArgs[:], nil)
	}

	retGo := AppInfoNewFromNative(ret.Pointer())

	return retGo
}

var appInfoGetFallbackForTypeFunction *gi.Function
var appInfoGetFallbackForTypeFunction_Once sync.Once

func appInfoGetFallbackForTypeFunction_Set() error {
	var err error
	appInfoGetFallbackForTypeFunction_Once.Do(func() {
		appInfoGetFallbackForTypeFunction, err = gi.FunctionInvokerNew("Gio", "app_info_get_fallback_for_type")
	})
	return err
}

// AppInfoGetFallbackForType is a representation of the C type g_app_info_get_fallback_for_type.
func AppInfoGetFallbackForType(contentType string) *glib.List {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(contentType)

	var ret gi.Argument

	err := appInfoGetFallbackForTypeFunction_Set()
	if err == nil {
		ret = appInfoGetFallbackForTypeFunction.Invoke(inArgs[:], nil)
	}

	retGo := glib.ListNewFromNative(ret.Pointer())

	return retGo
}

var appInfoGetRecommendedForTypeFunction *gi.Function
var appInfoGetRecommendedForTypeFunction_Once sync.Once

func appInfoGetRecommendedForTypeFunction_Set() error {
	var err error
	appInfoGetRecommendedForTypeFunction_Once.Do(func() {
		appInfoGetRecommendedForTypeFunction, err = gi.FunctionInvokerNew("Gio", "app_info_get_recommended_for_type")
	})
	return err
}

// AppInfoGetRecommendedForType is a representation of the C type g_app_info_get_recommended_for_type.
func AppInfoGetRecommendedForType(contentType string) *glib.List {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(contentType)

	var ret gi.Argument

	err := appInfoGetRecommendedForTypeFunction_Set()
	if err == nil {
		ret = appInfoGetRecommendedForTypeFunction.Invoke(inArgs[:], nil)
	}

	retGo := glib.ListNewFromNative(ret.Pointer())

	return retGo
}

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
	inArgs[1].SetPointer(context.Native())

	var ret gi.Argument

	err := appInfoLaunchDefaultForUriFunction_Set()
	if err == nil {
		ret = appInfoLaunchDefaultForUriFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'g_app_info_launch_default_for_uri_async' : parameter 'callback' of type 'AsyncReadyCallback' not supported

var appInfoLaunchDefaultForUriFinishFunction *gi.Function
var appInfoLaunchDefaultForUriFinishFunction_Once sync.Once

func appInfoLaunchDefaultForUriFinishFunction_Set() error {
	var err error
	appInfoLaunchDefaultForUriFinishFunction_Once.Do(func() {
		appInfoLaunchDefaultForUriFinishFunction, err = gi.FunctionInvokerNew("Gio", "app_info_launch_default_for_uri_finish")
	})
	return err
}

// AppInfoLaunchDefaultForUriFinish is a representation of the C type g_app_info_launch_default_for_uri_finish.
func AppInfoLaunchDefaultForUriFinish(result *AsyncResult) bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(result.Native())

	var ret gi.Argument

	err := appInfoLaunchDefaultForUriFinishFunction_Set()
	if err == nil {
		ret = appInfoLaunchDefaultForUriFinishFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

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

// UNSUPPORTED : C value 'g_async_initable_newv_async' : parameter 'callback' of type 'AsyncReadyCallback' not supported

// UNSUPPORTED : C value 'g_bus_get' : parameter 'callback' of type 'AsyncReadyCallback' not supported

var busGetFinishFunction *gi.Function
var busGetFinishFunction_Once sync.Once

func busGetFinishFunction_Set() error {
	var err error
	busGetFinishFunction_Once.Do(func() {
		busGetFinishFunction, err = gi.FunctionInvokerNew("Gio", "bus_get_finish")
	})
	return err
}

// BusGetFinish is a representation of the C type g_bus_get_finish.
func BusGetFinish(res *AsyncResult) *DBusConnection {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(res.Native())

	var ret gi.Argument

	err := busGetFinishFunction_Set()
	if err == nil {
		ret = busGetFinishFunction.Invoke(inArgs[:], nil)
	}

	retGo := DBusConnectionNewFromNative(ret.Pointer())

	return retGo
}

var busGetSyncFunction *gi.Function
var busGetSyncFunction_Once sync.Once

func busGetSyncFunction_Set() error {
	var err error
	busGetSyncFunction_Once.Do(func() {
		busGetSyncFunction, err = gi.FunctionInvokerNew("Gio", "bus_get_sync")
	})
	return err
}

// BusGetSync is a representation of the C type g_bus_get_sync.
func BusGetSync(busType BusType, cancellable *Cancellable) *DBusConnection {
	var inArgs [2]gi.Argument
	inArgs[0].SetInt32(int32(busType))
	inArgs[1].SetPointer(cancellable.Native())

	var ret gi.Argument

	err := busGetSyncFunction_Set()
	if err == nil {
		ret = busGetSyncFunction.Invoke(inArgs[:], nil)
	}

	retGo := DBusConnectionNewFromNative(ret.Pointer())

	return retGo
}

// UNSUPPORTED : C value 'g_bus_own_name' : parameter 'bus_acquired_handler' of type 'BusAcquiredCallback' not supported

// UNSUPPORTED : C value 'g_bus_own_name_on_connection' : parameter 'name_acquired_handler' of type 'BusNameAcquiredCallback' not supported

var busOwnNameOnConnectionWithClosuresFunction *gi.Function
var busOwnNameOnConnectionWithClosuresFunction_Once sync.Once

func busOwnNameOnConnectionWithClosuresFunction_Set() error {
	var err error
	busOwnNameOnConnectionWithClosuresFunction_Once.Do(func() {
		busOwnNameOnConnectionWithClosuresFunction, err = gi.FunctionInvokerNew("Gio", "bus_own_name_on_connection_with_closures")
	})
	return err
}

// BusOwnNameOnConnectionWithClosures is a representation of the C type g_bus_own_name_on_connection_with_closures.
func BusOwnNameOnConnectionWithClosures(connection *DBusConnection, name string, flags BusNameOwnerFlags, nameAcquiredClosure *gobject.Closure, nameLostClosure *gobject.Closure) uint32 {
	var inArgs [5]gi.Argument
	inArgs[0].SetPointer(connection.Native())
	inArgs[1].SetString(name)
	inArgs[2].SetInt32(int32(flags))
	inArgs[3].SetPointer(nameAcquiredClosure.Native())
	inArgs[4].SetPointer(nameLostClosure.Native())

	var ret gi.Argument

	err := busOwnNameOnConnectionWithClosuresFunction_Set()
	if err == nil {
		ret = busOwnNameOnConnectionWithClosuresFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo
}

var busOwnNameWithClosuresFunction *gi.Function
var busOwnNameWithClosuresFunction_Once sync.Once

func busOwnNameWithClosuresFunction_Set() error {
	var err error
	busOwnNameWithClosuresFunction_Once.Do(func() {
		busOwnNameWithClosuresFunction, err = gi.FunctionInvokerNew("Gio", "bus_own_name_with_closures")
	})
	return err
}

// BusOwnNameWithClosures is a representation of the C type g_bus_own_name_with_closures.
func BusOwnNameWithClosures(busType BusType, name string, flags BusNameOwnerFlags, busAcquiredClosure *gobject.Closure, nameAcquiredClosure *gobject.Closure, nameLostClosure *gobject.Closure) uint32 {
	var inArgs [6]gi.Argument
	inArgs[0].SetInt32(int32(busType))
	inArgs[1].SetString(name)
	inArgs[2].SetInt32(int32(flags))
	inArgs[3].SetPointer(busAcquiredClosure.Native())
	inArgs[4].SetPointer(nameAcquiredClosure.Native())
	inArgs[5].SetPointer(nameLostClosure.Native())

	var ret gi.Argument

	err := busOwnNameWithClosuresFunction_Set()
	if err == nil {
		ret = busOwnNameWithClosuresFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo
}

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

// UNSUPPORTED : C value 'g_bus_watch_name' : parameter 'name_appeared_handler' of type 'BusNameAppearedCallback' not supported

// UNSUPPORTED : C value 'g_bus_watch_name_on_connection' : parameter 'name_appeared_handler' of type 'BusNameAppearedCallback' not supported

var busWatchNameOnConnectionWithClosuresFunction *gi.Function
var busWatchNameOnConnectionWithClosuresFunction_Once sync.Once

func busWatchNameOnConnectionWithClosuresFunction_Set() error {
	var err error
	busWatchNameOnConnectionWithClosuresFunction_Once.Do(func() {
		busWatchNameOnConnectionWithClosuresFunction, err = gi.FunctionInvokerNew("Gio", "bus_watch_name_on_connection_with_closures")
	})
	return err
}

// BusWatchNameOnConnectionWithClosures is a representation of the C type g_bus_watch_name_on_connection_with_closures.
func BusWatchNameOnConnectionWithClosures(connection *DBusConnection, name string, flags BusNameWatcherFlags, nameAppearedClosure *gobject.Closure, nameVanishedClosure *gobject.Closure) uint32 {
	var inArgs [5]gi.Argument
	inArgs[0].SetPointer(connection.Native())
	inArgs[1].SetString(name)
	inArgs[2].SetInt32(int32(flags))
	inArgs[3].SetPointer(nameAppearedClosure.Native())
	inArgs[4].SetPointer(nameVanishedClosure.Native())

	var ret gi.Argument

	err := busWatchNameOnConnectionWithClosuresFunction_Set()
	if err == nil {
		ret = busWatchNameOnConnectionWithClosuresFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo
}

var busWatchNameWithClosuresFunction *gi.Function
var busWatchNameWithClosuresFunction_Once sync.Once

func busWatchNameWithClosuresFunction_Set() error {
	var err error
	busWatchNameWithClosuresFunction_Once.Do(func() {
		busWatchNameWithClosuresFunction, err = gi.FunctionInvokerNew("Gio", "bus_watch_name_with_closures")
	})
	return err
}

// BusWatchNameWithClosures is a representation of the C type g_bus_watch_name_with_closures.
func BusWatchNameWithClosures(busType BusType, name string, flags BusNameWatcherFlags, nameAppearedClosure *gobject.Closure, nameVanishedClosure *gobject.Closure) uint32 {
	var inArgs [5]gi.Argument
	inArgs[0].SetInt32(int32(busType))
	inArgs[1].SetString(name)
	inArgs[2].SetInt32(int32(flags))
	inArgs[3].SetPointer(nameAppearedClosure.Native())
	inArgs[4].SetPointer(nameVanishedClosure.Native())

	var ret gi.Argument

	err := busWatchNameWithClosuresFunction_Set()
	if err == nil {
		ret = busWatchNameWithClosuresFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo
}

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

var contentTypeGetIconFunction *gi.Function
var contentTypeGetIconFunction_Once sync.Once

func contentTypeGetIconFunction_Set() error {
	var err error
	contentTypeGetIconFunction_Once.Do(func() {
		contentTypeGetIconFunction, err = gi.FunctionInvokerNew("Gio", "content_type_get_icon")
	})
	return err
}

// ContentTypeGetIcon is a representation of the C type g_content_type_get_icon.
func ContentTypeGetIcon(type_ string) *Icon {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(type_)

	var ret gi.Argument

	err := contentTypeGetIconFunction_Set()
	if err == nil {
		ret = contentTypeGetIconFunction.Invoke(inArgs[:], nil)
	}

	retGo := IconNewFromNative(ret.Pointer())

	return retGo
}

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

var contentTypeGetSymbolicIconFunction *gi.Function
var contentTypeGetSymbolicIconFunction_Once sync.Once

func contentTypeGetSymbolicIconFunction_Set() error {
	var err error
	contentTypeGetSymbolicIconFunction_Once.Do(func() {
		contentTypeGetSymbolicIconFunction, err = gi.FunctionInvokerNew("Gio", "content_type_get_symbolic_icon")
	})
	return err
}

// ContentTypeGetSymbolicIcon is a representation of the C type g_content_type_get_symbolic_icon.
func ContentTypeGetSymbolicIcon(type_ string) *Icon {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(type_)

	var ret gi.Argument

	err := contentTypeGetSymbolicIconFunction_Set()
	if err == nil {
		ret = contentTypeGetSymbolicIconFunction.Invoke(inArgs[:], nil)
	}

	retGo := IconNewFromNative(ret.Pointer())

	return retGo
}

// UNSUPPORTED : C value 'g_content_type_guess' : array parameter 'data'

var contentTypeGuessForTreeFunction *gi.Function
var contentTypeGuessForTreeFunction_Once sync.Once

func contentTypeGuessForTreeFunction_Set() error {
	var err error
	contentTypeGuessForTreeFunction_Once.Do(func() {
		contentTypeGuessForTreeFunction, err = gi.FunctionInvokerNew("Gio", "content_type_guess_for_tree")
	})
	return err
}

// ContentTypeGuessForTree is a representation of the C type g_content_type_guess_for_tree.
func ContentTypeGuessForTree(root *File) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(root.Native())

	err := contentTypeGuessForTreeFunction_Set()
	if err == nil {
		contentTypeGuessForTreeFunction.Invoke(inArgs[:], nil)
	}

	return
}

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

// UNSUPPORTED : C value 'g_content_type_set_mime_dirs' : array parameter 'dirs'

var contentTypesGetRegisteredFunction *gi.Function
var contentTypesGetRegisteredFunction_Once sync.Once

func contentTypesGetRegisteredFunction_Set() error {
	var err error
	contentTypesGetRegisteredFunction_Once.Do(func() {
		contentTypesGetRegisteredFunction, err = gi.FunctionInvokerNew("Gio", "content_types_get_registered")
	})
	return err
}

// ContentTypesGetRegistered is a representation of the C type g_content_types_get_registered.
func ContentTypesGetRegistered() *glib.List {

	var ret gi.Argument

	err := contentTypesGetRegisteredFunction_Set()
	if err == nil {
		ret = contentTypesGetRegisteredFunction.Invoke(nil, nil)
	}

	retGo := glib.ListNewFromNative(ret.Pointer())

	return retGo
}

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

var dbusAddressGetForBusSyncFunction *gi.Function
var dbusAddressGetForBusSyncFunction_Once sync.Once

func dbusAddressGetForBusSyncFunction_Set() error {
	var err error
	dbusAddressGetForBusSyncFunction_Once.Do(func() {
		dbusAddressGetForBusSyncFunction, err = gi.FunctionInvokerNew("Gio", "dbus_address_get_for_bus_sync")
	})
	return err
}

// DbusAddressGetForBusSync is a representation of the C type g_dbus_address_get_for_bus_sync.
func DbusAddressGetForBusSync(busType BusType, cancellable *Cancellable) string {
	var inArgs [2]gi.Argument
	inArgs[0].SetInt32(int32(busType))
	inArgs[1].SetPointer(cancellable.Native())

	var ret gi.Argument

	err := dbusAddressGetForBusSyncFunction_Set()
	if err == nil {
		ret = dbusAddressGetForBusSyncFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

// UNSUPPORTED : C value 'g_dbus_address_get_stream' : parameter 'callback' of type 'AsyncReadyCallback' not supported

var dbusAddressGetStreamFinishFunction *gi.Function
var dbusAddressGetStreamFinishFunction_Once sync.Once

func dbusAddressGetStreamFinishFunction_Set() error {
	var err error
	dbusAddressGetStreamFinishFunction_Once.Do(func() {
		dbusAddressGetStreamFinishFunction, err = gi.FunctionInvokerNew("Gio", "dbus_address_get_stream_finish")
	})
	return err
}

// DbusAddressGetStreamFinish is a representation of the C type g_dbus_address_get_stream_finish.
func DbusAddressGetStreamFinish(res *AsyncResult) (*IOStream, string) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(res.Native())

	var outArgs [1]gi.Argument
	var ret gi.Argument

	err := dbusAddressGetStreamFinishFunction_Set()
	if err == nil {
		ret = dbusAddressGetStreamFinishFunction.Invoke(inArgs[:], outArgs[:])
	}

	retGo := IOStreamNewFromNative(ret.Pointer())
	out0 := outArgs[0].String(true)

	return retGo, out0
}

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
	inArgs[1].SetPointer(cancellable.Native())

	var outArgs [1]gi.Argument
	var ret gi.Argument

	err := dbusAddressGetStreamSyncFunction_Set()
	if err == nil {
		ret = dbusAddressGetStreamSyncFunction.Invoke(inArgs[:], outArgs[:])
	}

	retGo := IOStreamNewFromNative(ret.Pointer())
	out0 := outArgs[0].String(true)

	return retGo, out0
}

// UNSUPPORTED : C value 'g_dbus_annotation_info_lookup' : array parameter 'annotations'

var dbusErrorEncodeGerrorFunction *gi.Function
var dbusErrorEncodeGerrorFunction_Once sync.Once

func dbusErrorEncodeGerrorFunction_Set() error {
	var err error
	dbusErrorEncodeGerrorFunction_Once.Do(func() {
		dbusErrorEncodeGerrorFunction, err = gi.FunctionInvokerNew("Gio", "dbus_error_encode_gerror")
	})
	return err
}

// DbusErrorEncodeGerror is a representation of the C type g_dbus_error_encode_gerror.
func DbusErrorEncodeGerror(error *glib.Error) string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(error.Native())

	var ret gi.Argument

	err := dbusErrorEncodeGerrorFunction_Set()
	if err == nil {
		ret = dbusErrorEncodeGerrorFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

var dbusErrorGetRemoteErrorFunction *gi.Function
var dbusErrorGetRemoteErrorFunction_Once sync.Once

func dbusErrorGetRemoteErrorFunction_Set() error {
	var err error
	dbusErrorGetRemoteErrorFunction_Once.Do(func() {
		dbusErrorGetRemoteErrorFunction, err = gi.FunctionInvokerNew("Gio", "dbus_error_get_remote_error")
	})
	return err
}

// DbusErrorGetRemoteError is a representation of the C type g_dbus_error_get_remote_error.
func DbusErrorGetRemoteError(error *glib.Error) string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(error.Native())

	var ret gi.Argument

	err := dbusErrorGetRemoteErrorFunction_Set()
	if err == nil {
		ret = dbusErrorGetRemoteErrorFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

var dbusErrorIsRemoteErrorFunction *gi.Function
var dbusErrorIsRemoteErrorFunction_Once sync.Once

func dbusErrorIsRemoteErrorFunction_Set() error {
	var err error
	dbusErrorIsRemoteErrorFunction_Once.Do(func() {
		dbusErrorIsRemoteErrorFunction, err = gi.FunctionInvokerNew("Gio", "dbus_error_is_remote_error")
	})
	return err
}

// DbusErrorIsRemoteError is a representation of the C type g_dbus_error_is_remote_error.
func DbusErrorIsRemoteError(error *glib.Error) bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(error.Native())

	var ret gi.Argument

	err := dbusErrorIsRemoteErrorFunction_Set()
	if err == nil {
		ret = dbusErrorIsRemoteErrorFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var dbusErrorNewForDbusErrorFunction *gi.Function
var dbusErrorNewForDbusErrorFunction_Once sync.Once

func dbusErrorNewForDbusErrorFunction_Set() error {
	var err error
	dbusErrorNewForDbusErrorFunction_Once.Do(func() {
		dbusErrorNewForDbusErrorFunction, err = gi.FunctionInvokerNew("Gio", "dbus_error_new_for_dbus_error")
	})
	return err
}

// DbusErrorNewForDbusError is a representation of the C type g_dbus_error_new_for_dbus_error.
func DbusErrorNewForDbusError(dbusErrorName string, dbusErrorMessage string) *glib.Error {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(dbusErrorName)
	inArgs[1].SetString(dbusErrorMessage)

	var ret gi.Argument

	err := dbusErrorNewForDbusErrorFunction_Set()
	if err == nil {
		ret = dbusErrorNewForDbusErrorFunction.Invoke(inArgs[:], nil)
	}

	retGo := glib.ErrorNewFromNative(ret.Pointer())

	return retGo
}

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

// UNSUPPORTED : C value 'g_dbus_error_register_error_domain' : array parameter 'entries'

var dbusErrorStripRemoteErrorFunction *gi.Function
var dbusErrorStripRemoteErrorFunction_Once sync.Once

func dbusErrorStripRemoteErrorFunction_Set() error {
	var err error
	dbusErrorStripRemoteErrorFunction_Once.Do(func() {
		dbusErrorStripRemoteErrorFunction, err = gi.FunctionInvokerNew("Gio", "dbus_error_strip_remote_error")
	})
	return err
}

// DbusErrorStripRemoteError is a representation of the C type g_dbus_error_strip_remote_error.
func DbusErrorStripRemoteError(error *glib.Error) bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(error.Native())

	var ret gi.Argument

	err := dbusErrorStripRemoteErrorFunction_Set()
	if err == nil {
		ret = dbusErrorStripRemoteErrorFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

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

var dbusGvalueToGvariantFunction *gi.Function
var dbusGvalueToGvariantFunction_Once sync.Once

func dbusGvalueToGvariantFunction_Set() error {
	var err error
	dbusGvalueToGvariantFunction_Once.Do(func() {
		dbusGvalueToGvariantFunction, err = gi.FunctionInvokerNew("Gio", "dbus_gvalue_to_gvariant")
	})
	return err
}

// DbusGvalueToGvariant is a representation of the C type g_dbus_gvalue_to_gvariant.
func DbusGvalueToGvariant(gvalue *gobject.Value, type_ *glib.VariantType) *glib.Variant {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(gvalue.Native())
	inArgs[1].SetPointer(type_.Native())

	var ret gi.Argument

	err := dbusGvalueToGvariantFunction_Set()
	if err == nil {
		ret = dbusGvalueToGvariantFunction.Invoke(inArgs[:], nil)
	}

	retGo := glib.VariantNewFromNative(ret.Pointer())

	return retGo
}

var dbusGvariantToGvalueFunction *gi.Function
var dbusGvariantToGvalueFunction_Once sync.Once

func dbusGvariantToGvalueFunction_Set() error {
	var err error
	dbusGvariantToGvalueFunction_Once.Do(func() {
		dbusGvariantToGvalueFunction, err = gi.FunctionInvokerNew("Gio", "dbus_gvariant_to_gvalue")
	})
	return err
}

// DbusGvariantToGvalue is a representation of the C type g_dbus_gvariant_to_gvalue.
func DbusGvariantToGvalue(value *glib.Variant) *gobject.Value {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(value.Native())

	var outArgs [1]gi.Argument

	err := dbusGvariantToGvalueFunction_Set()
	if err == nil {
		dbusGvariantToGvalueFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := gobject.ValueNewFromNative(outArgs[0].Pointer())

	return out0
}

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

var dtlsClientConnectionNewFunction *gi.Function
var dtlsClientConnectionNewFunction_Once sync.Once

func dtlsClientConnectionNewFunction_Set() error {
	var err error
	dtlsClientConnectionNewFunction_Once.Do(func() {
		dtlsClientConnectionNewFunction, err = gi.FunctionInvokerNew("Gio", "dtls_client_connection_new")
	})
	return err
}

// DtlsClientConnectionNew is a representation of the C type g_dtls_client_connection_new.
func DtlsClientConnectionNew(baseSocket *DatagramBased, serverIdentity *SocketConnectable) *DtlsClientConnection {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(baseSocket.Native())
	inArgs[1].SetPointer(serverIdentity.Native())

	var ret gi.Argument

	err := dtlsClientConnectionNewFunction_Set()
	if err == nil {
		ret = dtlsClientConnectionNewFunction.Invoke(inArgs[:], nil)
	}

	retGo := DtlsClientConnectionNewFromNative(ret.Pointer())

	return retGo
}

var dtlsServerConnectionNewFunction *gi.Function
var dtlsServerConnectionNewFunction_Once sync.Once

func dtlsServerConnectionNewFunction_Set() error {
	var err error
	dtlsServerConnectionNewFunction_Once.Do(func() {
		dtlsServerConnectionNewFunction, err = gi.FunctionInvokerNew("Gio", "dtls_server_connection_new")
	})
	return err
}

// DtlsServerConnectionNew is a representation of the C type g_dtls_server_connection_new.
func DtlsServerConnectionNew(baseSocket *DatagramBased, certificate *TlsCertificate) *DtlsServerConnection {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(baseSocket.Native())
	inArgs[1].SetPointer(certificate.Native())

	var ret gi.Argument

	err := dtlsServerConnectionNewFunction_Set()
	if err == nil {
		ret = dtlsServerConnectionNewFunction.Invoke(inArgs[:], nil)
	}

	retGo := DtlsServerConnectionNewFromNative(ret.Pointer())

	return retGo
}

var fileNewForCommandlineArgFunction *gi.Function
var fileNewForCommandlineArgFunction_Once sync.Once

func fileNewForCommandlineArgFunction_Set() error {
	var err error
	fileNewForCommandlineArgFunction_Once.Do(func() {
		fileNewForCommandlineArgFunction, err = gi.FunctionInvokerNew("Gio", "file_new_for_commandline_arg")
	})
	return err
}

// FileNewForCommandlineArg is a representation of the C type g_file_new_for_commandline_arg.
func FileNewForCommandlineArg(arg string) *File {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(arg)

	var ret gi.Argument

	err := fileNewForCommandlineArgFunction_Set()
	if err == nil {
		ret = fileNewForCommandlineArgFunction.Invoke(inArgs[:], nil)
	}

	retGo := FileNewFromNative(ret.Pointer())

	return retGo
}

var fileNewForCommandlineArgAndCwdFunction *gi.Function
var fileNewForCommandlineArgAndCwdFunction_Once sync.Once

func fileNewForCommandlineArgAndCwdFunction_Set() error {
	var err error
	fileNewForCommandlineArgAndCwdFunction_Once.Do(func() {
		fileNewForCommandlineArgAndCwdFunction, err = gi.FunctionInvokerNew("Gio", "file_new_for_commandline_arg_and_cwd")
	})
	return err
}

// FileNewForCommandlineArgAndCwd is a representation of the C type g_file_new_for_commandline_arg_and_cwd.
func FileNewForCommandlineArgAndCwd(arg string, cwd string) *File {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(arg)
	inArgs[1].SetString(cwd)

	var ret gi.Argument

	err := fileNewForCommandlineArgAndCwdFunction_Set()
	if err == nil {
		ret = fileNewForCommandlineArgAndCwdFunction.Invoke(inArgs[:], nil)
	}

	retGo := FileNewFromNative(ret.Pointer())

	return retGo
}

var fileNewForPathFunction *gi.Function
var fileNewForPathFunction_Once sync.Once

func fileNewForPathFunction_Set() error {
	var err error
	fileNewForPathFunction_Once.Do(func() {
		fileNewForPathFunction, err = gi.FunctionInvokerNew("Gio", "file_new_for_path")
	})
	return err
}

// FileNewForPath is a representation of the C type g_file_new_for_path.
func FileNewForPath(path string) *File {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(path)

	var ret gi.Argument

	err := fileNewForPathFunction_Set()
	if err == nil {
		ret = fileNewForPathFunction.Invoke(inArgs[:], nil)
	}

	retGo := FileNewFromNative(ret.Pointer())

	return retGo
}

var fileNewForUriFunction *gi.Function
var fileNewForUriFunction_Once sync.Once

func fileNewForUriFunction_Set() error {
	var err error
	fileNewForUriFunction_Once.Do(func() {
		fileNewForUriFunction, err = gi.FunctionInvokerNew("Gio", "file_new_for_uri")
	})
	return err
}

// FileNewForUri is a representation of the C type g_file_new_for_uri.
func FileNewForUri(uri string) *File {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(uri)

	var ret gi.Argument

	err := fileNewForUriFunction_Set()
	if err == nil {
		ret = fileNewForUriFunction.Invoke(inArgs[:], nil)
	}

	retGo := FileNewFromNative(ret.Pointer())

	return retGo
}

var fileNewTmpFunction *gi.Function
var fileNewTmpFunction_Once sync.Once

func fileNewTmpFunction_Set() error {
	var err error
	fileNewTmpFunction_Once.Do(func() {
		fileNewTmpFunction, err = gi.FunctionInvokerNew("Gio", "file_new_tmp")
	})
	return err
}

// FileNewTmp is a representation of the C type g_file_new_tmp.
func FileNewTmp(tmpl string) (*File, *FileIOStream) {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(tmpl)

	var outArgs [1]gi.Argument
	var ret gi.Argument

	err := fileNewTmpFunction_Set()
	if err == nil {
		ret = fileNewTmpFunction.Invoke(inArgs[:], outArgs[:])
	}

	retGo := FileNewFromNative(ret.Pointer())
	out0 := FileIOStreamNewFromNative(outArgs[0].Pointer())

	return retGo, out0
}

var fileParseNameFunction *gi.Function
var fileParseNameFunction_Once sync.Once

func fileParseNameFunction_Set() error {
	var err error
	fileParseNameFunction_Once.Do(func() {
		fileParseNameFunction, err = gi.FunctionInvokerNew("Gio", "file_parse_name")
	})
	return err
}

// FileParseName is a representation of the C type g_file_parse_name.
func FileParseName(parseName string) *File {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(parseName)

	var ret gi.Argument

	err := fileParseNameFunction_Set()
	if err == nil {
		ret = fileParseNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := FileNewFromNative(ret.Pointer())

	return retGo
}

var iconDeserializeFunction *gi.Function
var iconDeserializeFunction_Once sync.Once

func iconDeserializeFunction_Set() error {
	var err error
	iconDeserializeFunction_Once.Do(func() {
		iconDeserializeFunction, err = gi.FunctionInvokerNew("Gio", "icon_deserialize")
	})
	return err
}

// IconDeserialize is a representation of the C type g_icon_deserialize.
func IconDeserialize(value *glib.Variant) *Icon {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(value.Native())

	var ret gi.Argument

	err := iconDeserializeFunction_Set()
	if err == nil {
		ret = iconDeserializeFunction.Invoke(inArgs[:], nil)
	}

	retGo := IconNewFromNative(ret.Pointer())

	return retGo
}

var iconHashFunction *gi.Function
var iconHashFunction_Once sync.Once

func iconHashFunction_Set() error {
	var err error
	iconHashFunction_Once.Do(func() {
		iconHashFunction, err = gi.FunctionInvokerNew("Gio", "icon_hash")
	})
	return err
}

// IconHash is a representation of the C type g_icon_hash.
func IconHash(icon unsafe.Pointer) uint32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(icon)

	var ret gi.Argument

	err := iconHashFunction_Set()
	if err == nil {
		ret = iconHashFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo
}

var iconNewForStringFunction *gi.Function
var iconNewForStringFunction_Once sync.Once

func iconNewForStringFunction_Set() error {
	var err error
	iconNewForStringFunction_Once.Do(func() {
		iconNewForStringFunction, err = gi.FunctionInvokerNew("Gio", "icon_new_for_string")
	})
	return err
}

// IconNewForString is a representation of the C type g_icon_new_for_string.
func IconNewForString(str string) *Icon {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(str)

	var ret gi.Argument

	err := iconNewForStringFunction_Set()
	if err == nil {
		ret = iconNewForStringFunction.Invoke(inArgs[:], nil)
	}

	retGo := IconNewFromNative(ret.Pointer())

	return retGo
}

// UNSUPPORTED : C value 'g_initable_newv' : array parameter 'parameters'

var ioErrorFromErrnoFunction *gi.Function
var ioErrorFromErrnoFunction_Once sync.Once

func ioErrorFromErrnoFunction_Set() error {
	var err error
	ioErrorFromErrnoFunction_Once.Do(func() {
		ioErrorFromErrnoFunction, err = gi.FunctionInvokerNew("Gio", "io_error_from_errno")
	})
	return err
}

// IoErrorFromErrno is a representation of the C type g_io_error_from_errno.
func IoErrorFromErrno(errNo int32) IOErrorEnum {
	var inArgs [1]gi.Argument
	inArgs[0].SetInt32(errNo)

	var ret gi.Argument

	err := ioErrorFromErrnoFunction_Set()
	if err == nil {
		ret = ioErrorFromErrnoFunction.Invoke(inArgs[:], nil)
	}

	retGo := IOErrorEnum(ret.Int32())

	return retGo
}

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

var ioExtensionPointImplementFunction *gi.Function
var ioExtensionPointImplementFunction_Once sync.Once

func ioExtensionPointImplementFunction_Set() error {
	var err error
	ioExtensionPointImplementFunction_Once.Do(func() {
		ioExtensionPointImplementFunction, err = gi.FunctionInvokerNew("Gio", "io_extension_point_implement")
	})
	return err
}

// IoExtensionPointImplement is a representation of the C type g_io_extension_point_implement.
func IoExtensionPointImplement(extensionPointName string, type_ int64, extensionName string, priority int32) *IOExtension {
	var inArgs [4]gi.Argument
	inArgs[0].SetString(extensionPointName)
	inArgs[1].SetInt64(type_)
	inArgs[2].SetString(extensionName)
	inArgs[3].SetInt32(priority)

	var ret gi.Argument

	err := ioExtensionPointImplementFunction_Set()
	if err == nil {
		ret = ioExtensionPointImplementFunction.Invoke(inArgs[:], nil)
	}

	retGo := IOExtensionNewFromNative(ret.Pointer())

	return retGo
}

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

	retGo := IOExtensionPointNewFromNative(ret.Pointer())

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

	retGo := IOExtensionPointNewFromNative(ret.Pointer())

	return retGo
}

var ioModulesLoadAllInDirectoryFunction *gi.Function
var ioModulesLoadAllInDirectoryFunction_Once sync.Once

func ioModulesLoadAllInDirectoryFunction_Set() error {
	var err error
	ioModulesLoadAllInDirectoryFunction_Once.Do(func() {
		ioModulesLoadAllInDirectoryFunction, err = gi.FunctionInvokerNew("Gio", "io_modules_load_all_in_directory")
	})
	return err
}

// IoModulesLoadAllInDirectory is a representation of the C type g_io_modules_load_all_in_directory.
func IoModulesLoadAllInDirectory(dirname string) *glib.List {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(dirname)

	var ret gi.Argument

	err := ioModulesLoadAllInDirectoryFunction_Set()
	if err == nil {
		ret = ioModulesLoadAllInDirectoryFunction.Invoke(inArgs[:], nil)
	}

	retGo := glib.ListNewFromNative(ret.Pointer())

	return retGo
}

var ioModulesLoadAllInDirectoryWithScopeFunction *gi.Function
var ioModulesLoadAllInDirectoryWithScopeFunction_Once sync.Once

func ioModulesLoadAllInDirectoryWithScopeFunction_Set() error {
	var err error
	ioModulesLoadAllInDirectoryWithScopeFunction_Once.Do(func() {
		ioModulesLoadAllInDirectoryWithScopeFunction, err = gi.FunctionInvokerNew("Gio", "io_modules_load_all_in_directory_with_scope")
	})
	return err
}

// IoModulesLoadAllInDirectoryWithScope is a representation of the C type g_io_modules_load_all_in_directory_with_scope.
func IoModulesLoadAllInDirectoryWithScope(dirname string, scope *IOModuleScope) *glib.List {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(dirname)
	inArgs[1].SetPointer(scope.Native())

	var ret gi.Argument

	err := ioModulesLoadAllInDirectoryWithScopeFunction_Set()
	if err == nil {
		ret = ioModulesLoadAllInDirectoryWithScopeFunction.Invoke(inArgs[:], nil)
	}

	retGo := glib.ListNewFromNative(ret.Pointer())

	return retGo
}

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
	inArgs[1].SetPointer(scope.Native())

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

	retGo := SettingsBackendNewFromNative(ret.Pointer())

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

	retGo := SettingsBackendNewFromNative(ret.Pointer())

	return retGo
}

var networkMonitorGetDefaultFunction *gi.Function
var networkMonitorGetDefaultFunction_Once sync.Once

func networkMonitorGetDefaultFunction_Set() error {
	var err error
	networkMonitorGetDefaultFunction_Once.Do(func() {
		networkMonitorGetDefaultFunction, err = gi.FunctionInvokerNew("Gio", "network_monitor_get_default")
	})
	return err
}

// NetworkMonitorGetDefault is a representation of the C type g_network_monitor_get_default.
func NetworkMonitorGetDefault() *NetworkMonitor {

	var ret gi.Argument

	err := networkMonitorGetDefaultFunction_Set()
	if err == nil {
		ret = networkMonitorGetDefaultFunction.Invoke(nil, nil)
	}

	retGo := NetworkMonitorNewFromNative(ret.Pointer())

	return retGo
}

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

	retGo := SettingsBackendNewFromNative(ret.Pointer())

	return retGo
}

var pollableSourceNewFunction *gi.Function
var pollableSourceNewFunction_Once sync.Once

func pollableSourceNewFunction_Set() error {
	var err error
	pollableSourceNewFunction_Once.Do(func() {
		pollableSourceNewFunction, err = gi.FunctionInvokerNew("Gio", "pollable_source_new")
	})
	return err
}

// PollableSourceNew is a representation of the C type g_pollable_source_new.
func PollableSourceNew(pollableStream *gobject.Object) *glib.Source {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(pollableStream.Native())

	var ret gi.Argument

	err := pollableSourceNewFunction_Set()
	if err == nil {
		ret = pollableSourceNewFunction.Invoke(inArgs[:], nil)
	}

	retGo := glib.SourceNewFromNative(ret.Pointer())

	return retGo
}

var pollableSourceNewFullFunction *gi.Function
var pollableSourceNewFullFunction_Once sync.Once

func pollableSourceNewFullFunction_Set() error {
	var err error
	pollableSourceNewFullFunction_Once.Do(func() {
		pollableSourceNewFullFunction, err = gi.FunctionInvokerNew("Gio", "pollable_source_new_full")
	})
	return err
}

// PollableSourceNewFull is a representation of the C type g_pollable_source_new_full.
func PollableSourceNewFull(pollableStream *gobject.Object, childSource *glib.Source, cancellable *Cancellable) *glib.Source {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(pollableStream.Native())
	inArgs[1].SetPointer(childSource.Native())
	inArgs[2].SetPointer(cancellable.Native())

	var ret gi.Argument

	err := pollableSourceNewFullFunction_Set()
	if err == nil {
		ret = pollableSourceNewFullFunction.Invoke(inArgs[:], nil)
	}

	retGo := glib.SourceNewFromNative(ret.Pointer())

	return retGo
}

// UNSUPPORTED : C value 'g_pollable_stream_read' : array parameter 'buffer'

// UNSUPPORTED : C value 'g_pollable_stream_write' : array parameter 'buffer'

// UNSUPPORTED : C value 'g_pollable_stream_write_all' : array parameter 'buffer'

var proxyGetDefaultForProtocolFunction *gi.Function
var proxyGetDefaultForProtocolFunction_Once sync.Once

func proxyGetDefaultForProtocolFunction_Set() error {
	var err error
	proxyGetDefaultForProtocolFunction_Once.Do(func() {
		proxyGetDefaultForProtocolFunction, err = gi.FunctionInvokerNew("Gio", "proxy_get_default_for_protocol")
	})
	return err
}

// ProxyGetDefaultForProtocol is a representation of the C type g_proxy_get_default_for_protocol.
func ProxyGetDefaultForProtocol(protocol string) *Proxy {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(protocol)

	var ret gi.Argument

	err := proxyGetDefaultForProtocolFunction_Set()
	if err == nil {
		ret = proxyGetDefaultForProtocolFunction.Invoke(inArgs[:], nil)
	}

	retGo := ProxyNewFromNative(ret.Pointer())

	return retGo
}

var proxyResolverGetDefaultFunction *gi.Function
var proxyResolverGetDefaultFunction_Once sync.Once

func proxyResolverGetDefaultFunction_Set() error {
	var err error
	proxyResolverGetDefaultFunction_Once.Do(func() {
		proxyResolverGetDefaultFunction, err = gi.FunctionInvokerNew("Gio", "proxy_resolver_get_default")
	})
	return err
}

// ProxyResolverGetDefault is a representation of the C type g_proxy_resolver_get_default.
func ProxyResolverGetDefault() *ProxyResolver {

	var ret gi.Argument

	err := proxyResolverGetDefaultFunction_Set()
	if err == nil {
		ret = proxyResolverGetDefaultFunction.Invoke(nil, nil)
	}

	retGo := ProxyResolverNewFromNative(ret.Pointer())

	return retGo
}

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

	retGo := ResourceNewFromNative(ret.Pointer())

	return retGo
}

var resourcesEnumerateChildrenFunction *gi.Function
var resourcesEnumerateChildrenFunction_Once sync.Once

func resourcesEnumerateChildrenFunction_Set() error {
	var err error
	resourcesEnumerateChildrenFunction_Once.Do(func() {
		resourcesEnumerateChildrenFunction, err = gi.FunctionInvokerNew("Gio", "resources_enumerate_children")
	})
	return err
}

// ResourcesEnumerateChildren is a representation of the C type g_resources_enumerate_children.
func ResourcesEnumerateChildren(path string, lookupFlags ResourceLookupFlags) {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(path)
	inArgs[1].SetInt32(int32(lookupFlags))

	err := resourcesEnumerateChildrenFunction_Set()
	if err == nil {
		resourcesEnumerateChildrenFunction.Invoke(inArgs[:], nil)
	}

	return
}

var resourcesGetInfoFunction *gi.Function
var resourcesGetInfoFunction_Once sync.Once

func resourcesGetInfoFunction_Set() error {
	var err error
	resourcesGetInfoFunction_Once.Do(func() {
		resourcesGetInfoFunction, err = gi.FunctionInvokerNew("Gio", "resources_get_info")
	})
	return err
}

// ResourcesGetInfo is a representation of the C type g_resources_get_info.
func ResourcesGetInfo(path string, lookupFlags ResourceLookupFlags) (bool, uint64, uint32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(path)
	inArgs[1].SetInt32(int32(lookupFlags))

	var outArgs [2]gi.Argument
	var ret gi.Argument

	err := resourcesGetInfoFunction_Set()
	if err == nil {
		ret = resourcesGetInfoFunction.Invoke(inArgs[:], outArgs[:])
	}

	retGo := ret.Boolean()
	out0 := outArgs[0].Uint64()
	out1 := outArgs[1].Uint32()

	return retGo, out0, out1
}

var resourcesLookupDataFunction *gi.Function
var resourcesLookupDataFunction_Once sync.Once

func resourcesLookupDataFunction_Set() error {
	var err error
	resourcesLookupDataFunction_Once.Do(func() {
		resourcesLookupDataFunction, err = gi.FunctionInvokerNew("Gio", "resources_lookup_data")
	})
	return err
}

// ResourcesLookupData is a representation of the C type g_resources_lookup_data.
func ResourcesLookupData(path string, lookupFlags ResourceLookupFlags) *glib.Bytes {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(path)
	inArgs[1].SetInt32(int32(lookupFlags))

	var ret gi.Argument

	err := resourcesLookupDataFunction_Set()
	if err == nil {
		ret = resourcesLookupDataFunction.Invoke(inArgs[:], nil)
	}

	retGo := glib.BytesNewFromNative(ret.Pointer())

	return retGo
}

var resourcesOpenStreamFunction *gi.Function
var resourcesOpenStreamFunction_Once sync.Once

func resourcesOpenStreamFunction_Set() error {
	var err error
	resourcesOpenStreamFunction_Once.Do(func() {
		resourcesOpenStreamFunction, err = gi.FunctionInvokerNew("Gio", "resources_open_stream")
	})
	return err
}

// ResourcesOpenStream is a representation of the C type g_resources_open_stream.
func ResourcesOpenStream(path string, lookupFlags ResourceLookupFlags) *InputStream {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(path)
	inArgs[1].SetInt32(int32(lookupFlags))

	var ret gi.Argument

	err := resourcesOpenStreamFunction_Set()
	if err == nil {
		ret = resourcesOpenStreamFunction.Invoke(inArgs[:], nil)
	}

	retGo := InputStreamNewFromNative(ret.Pointer())

	return retGo
}

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
	inArgs[0].SetPointer(resource.Native())

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
	inArgs[0].SetPointer(resource.Native())

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

	retGo := SettingsSchemaSourceNewFromNative(ret.Pointer())

	return retGo
}

// UNSUPPORTED : C value 'g_simple_async_report_error_in_idle' : parameter 'callback' of type 'AsyncReadyCallback' not supported

// UNSUPPORTED : C value 'g_simple_async_report_gerror_in_idle' : parameter 'callback' of type 'AsyncReadyCallback' not supported

// UNSUPPORTED : C value 'g_simple_async_report_take_gerror_in_idle' : parameter 'callback' of type 'AsyncReadyCallback' not supported

var srvTargetListSortFunction *gi.Function
var srvTargetListSortFunction_Once sync.Once

func srvTargetListSortFunction_Set() error {
	var err error
	srvTargetListSortFunction_Once.Do(func() {
		srvTargetListSortFunction, err = gi.FunctionInvokerNew("Gio", "srv_target_list_sort")
	})
	return err
}

// SrvTargetListSort is a representation of the C type g_srv_target_list_sort.
func SrvTargetListSort(targets *glib.List) *glib.List {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(targets.Native())

	var ret gi.Argument

	err := srvTargetListSortFunction_Set()
	if err == nil {
		ret = srvTargetListSortFunction.Invoke(inArgs[:], nil)
	}

	retGo := glib.ListNewFromNative(ret.Pointer())

	return retGo
}

var tlsBackendGetDefaultFunction *gi.Function
var tlsBackendGetDefaultFunction_Once sync.Once

func tlsBackendGetDefaultFunction_Set() error {
	var err error
	tlsBackendGetDefaultFunction_Once.Do(func() {
		tlsBackendGetDefaultFunction, err = gi.FunctionInvokerNew("Gio", "tls_backend_get_default")
	})
	return err
}

// TlsBackendGetDefault is a representation of the C type g_tls_backend_get_default.
func TlsBackendGetDefault() *TlsBackend {

	var ret gi.Argument

	err := tlsBackendGetDefaultFunction_Set()
	if err == nil {
		ret = tlsBackendGetDefaultFunction.Invoke(nil, nil)
	}

	retGo := TlsBackendNewFromNative(ret.Pointer())

	return retGo
}

var tlsClientConnectionNewFunction *gi.Function
var tlsClientConnectionNewFunction_Once sync.Once

func tlsClientConnectionNewFunction_Set() error {
	var err error
	tlsClientConnectionNewFunction_Once.Do(func() {
		tlsClientConnectionNewFunction, err = gi.FunctionInvokerNew("Gio", "tls_client_connection_new")
	})
	return err
}

// TlsClientConnectionNew is a representation of the C type g_tls_client_connection_new.
func TlsClientConnectionNew(baseIoStream *IOStream, serverIdentity *SocketConnectable) *TlsClientConnection {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(baseIoStream.Native())
	inArgs[1].SetPointer(serverIdentity.Native())

	var ret gi.Argument

	err := tlsClientConnectionNewFunction_Set()
	if err == nil {
		ret = tlsClientConnectionNewFunction.Invoke(inArgs[:], nil)
	}

	retGo := TlsClientConnectionNewFromNative(ret.Pointer())

	return retGo
}

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

var tlsFileDatabaseNewFunction *gi.Function
var tlsFileDatabaseNewFunction_Once sync.Once

func tlsFileDatabaseNewFunction_Set() error {
	var err error
	tlsFileDatabaseNewFunction_Once.Do(func() {
		tlsFileDatabaseNewFunction, err = gi.FunctionInvokerNew("Gio", "tls_file_database_new")
	})
	return err
}

// TlsFileDatabaseNew is a representation of the C type g_tls_file_database_new.
func TlsFileDatabaseNew(anchors string) *TlsFileDatabase {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(anchors)

	var ret gi.Argument

	err := tlsFileDatabaseNewFunction_Set()
	if err == nil {
		ret = tlsFileDatabaseNewFunction.Invoke(inArgs[:], nil)
	}

	retGo := TlsFileDatabaseNewFromNative(ret.Pointer())

	return retGo
}

var tlsServerConnectionNewFunction *gi.Function
var tlsServerConnectionNewFunction_Once sync.Once

func tlsServerConnectionNewFunction_Set() error {
	var err error
	tlsServerConnectionNewFunction_Once.Do(func() {
		tlsServerConnectionNewFunction, err = gi.FunctionInvokerNew("Gio", "tls_server_connection_new")
	})
	return err
}

// TlsServerConnectionNew is a representation of the C type g_tls_server_connection_new.
func TlsServerConnectionNew(baseIoStream *IOStream, certificate *TlsCertificate) *TlsServerConnection {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(baseIoStream.Native())
	inArgs[1].SetPointer(certificate.Native())

	var ret gi.Argument

	err := tlsServerConnectionNewFunction_Set()
	if err == nil {
		ret = tlsServerConnectionNewFunction.Invoke(inArgs[:], nil)
	}

	retGo := TlsServerConnectionNewFromNative(ret.Pointer())

	return retGo
}

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

	retGo := UnixMountEntryNewFromNative(ret.Pointer())
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
	inArgs[0].SetPointer(mount1.Native())
	inArgs[1].SetPointer(mount2.Native())

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
	inArgs[0].SetPointer(mountEntry.Native())

	var ret gi.Argument

	err := unixMountCopyFunction_Set()
	if err == nil {
		ret = unixMountCopyFunction.Invoke(inArgs[:], nil)
	}

	retGo := UnixMountEntryNewFromNative(ret.Pointer())

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

	retGo := UnixMountEntryNewFromNative(ret.Pointer())
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
	inArgs[0].SetPointer(mountEntry.Native())

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
	inArgs[0].SetPointer(mountEntry.Native())

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
	inArgs[0].SetPointer(mountEntry.Native())

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
	inArgs[0].SetPointer(mountEntry.Native())

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
	inArgs[0].SetPointer(mountEntry.Native())

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
	inArgs[0].SetPointer(mountEntry.Native())

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
	inArgs[0].SetPointer(mountEntry.Native())

	var ret gi.Argument

	err := unixMountGuessCanEjectFunction_Set()
	if err == nil {
		ret = unixMountGuessCanEjectFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var unixMountGuessIconFunction *gi.Function
var unixMountGuessIconFunction_Once sync.Once

func unixMountGuessIconFunction_Set() error {
	var err error
	unixMountGuessIconFunction_Once.Do(func() {
		unixMountGuessIconFunction, err = gi.FunctionInvokerNew("Gio", "unix_mount_guess_icon")
	})
	return err
}

// UnixMountGuessIcon is a representation of the C type g_unix_mount_guess_icon.
func UnixMountGuessIcon(mountEntry *UnixMountEntry) *Icon {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(mountEntry.Native())

	var ret gi.Argument

	err := unixMountGuessIconFunction_Set()
	if err == nil {
		ret = unixMountGuessIconFunction.Invoke(inArgs[:], nil)
	}

	retGo := IconNewFromNative(ret.Pointer())

	return retGo
}

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
	inArgs[0].SetPointer(mountEntry.Native())

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
	inArgs[0].SetPointer(mountEntry.Native())

	var ret gi.Argument

	err := unixMountGuessShouldDisplayFunction_Set()
	if err == nil {
		ret = unixMountGuessShouldDisplayFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var unixMountGuessSymbolicIconFunction *gi.Function
var unixMountGuessSymbolicIconFunction_Once sync.Once

func unixMountGuessSymbolicIconFunction_Set() error {
	var err error
	unixMountGuessSymbolicIconFunction_Once.Do(func() {
		unixMountGuessSymbolicIconFunction, err = gi.FunctionInvokerNew("Gio", "unix_mount_guess_symbolic_icon")
	})
	return err
}

// UnixMountGuessSymbolicIcon is a representation of the C type g_unix_mount_guess_symbolic_icon.
func UnixMountGuessSymbolicIcon(mountEntry *UnixMountEntry) *Icon {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(mountEntry.Native())

	var ret gi.Argument

	err := unixMountGuessSymbolicIconFunction_Set()
	if err == nil {
		ret = unixMountGuessSymbolicIconFunction.Invoke(inArgs[:], nil)
	}

	retGo := IconNewFromNative(ret.Pointer())

	return retGo
}

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
	inArgs[0].SetPointer(mountEntry.Native())

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
	inArgs[0].SetPointer(mountEntry.Native())

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

var unixMountPointsGetFunction *gi.Function
var unixMountPointsGetFunction_Once sync.Once

func unixMountPointsGetFunction_Set() error {
	var err error
	unixMountPointsGetFunction_Once.Do(func() {
		unixMountPointsGetFunction, err = gi.FunctionInvokerNew("Gio", "unix_mount_points_get")
	})
	return err
}

// UnixMountPointsGet is a representation of the C type g_unix_mount_points_get.
func UnixMountPointsGet() (*glib.List, uint64) {

	var outArgs [1]gi.Argument
	var ret gi.Argument

	err := unixMountPointsGetFunction_Set()
	if err == nil {
		ret = unixMountPointsGetFunction.Invoke(nil, outArgs[:])
	}

	retGo := glib.ListNewFromNative(ret.Pointer())
	out0 := outArgs[0].Uint64()

	return retGo, out0
}

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

var unixMountsGetFunction *gi.Function
var unixMountsGetFunction_Once sync.Once

func unixMountsGetFunction_Set() error {
	var err error
	unixMountsGetFunction_Once.Do(func() {
		unixMountsGetFunction, err = gi.FunctionInvokerNew("Gio", "unix_mounts_get")
	})
	return err
}

// UnixMountsGet is a representation of the C type g_unix_mounts_get.
func UnixMountsGet() (*glib.List, uint64) {

	var outArgs [1]gi.Argument
	var ret gi.Argument

	err := unixMountsGetFunction_Set()
	if err == nil {
		ret = unixMountsGetFunction.Invoke(nil, outArgs[:])
	}

	retGo := glib.ListNewFromNative(ret.Pointer())
	out0 := outArgs[0].Uint64()

	return retGo, out0
}
