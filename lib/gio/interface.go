// Code generated - DO NOT EDIT.

package gio

import (
	callback "github.com/pekim/gobbi/internal/cgo/callback"
	gi "github.com/pekim/gobbi/internal/cgo/gi"
	gobject "github.com/pekim/gobbi/lib/gobject"
	"sync"
	"unsafe"
)

var actionInterface *gi.Interface
var actionInterface_Once sync.Once

func actionInterface_Set() error {
	var err error
	actionInterface_Once.Do(func() {
		actionInterface, err = gi.InterfaceNew("Gio", "Action")
	})
	return err
}

type Action struct {
	native unsafe.Pointer
}

func ActionNewFromNative(native unsafe.Pointer) *Action {
	instance := &Action{native: native}

	return instance
}

/*
CastToAction down casts any arbitrary Object to Action.
Exercise care, as this is a potentially dangerous function
if the Object is not a Action.
*/
func CastToAction(object *gobject.Object) *Action {
	return ActionNewFromNative(object.Native())
}

// Equals compares this Action with another Action, and returns true if they represent the same Object.
func (recv *Action) Equals(other *Action) bool {
	return other.Native() == recv.Native()
}

func (recv *Action) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'g_action_activate' : parameter 'parameter' of type 'GLib.Variant' not supported

// UNSUPPORTED : C value 'g_action_change_state' : parameter 'value' of type 'GLib.Variant' not supported

var actionGetEnabledFunction *gi.Function
var actionGetEnabledFunction_Once sync.Once

func actionGetEnabledFunction_Set() error {
	var err error
	actionGetEnabledFunction_Once.Do(func() {
		err = actionInterface_Set()
		if err != nil {
			return
		}
		actionGetEnabledFunction, err = actionInterface.InvokerNew("get_enabled")
	})
	return err
}

// GetEnabled is a representation of the C type g_action_get_enabled.
func (recv *Action) GetEnabled() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := actionGetEnabledFunction_Set()
	if err == nil {
		ret = actionGetEnabledFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var actionGetNameFunction *gi.Function
var actionGetNameFunction_Once sync.Once

func actionGetNameFunction_Set() error {
	var err error
	actionGetNameFunction_Once.Do(func() {
		err = actionInterface_Set()
		if err != nil {
			return
		}
		actionGetNameFunction, err = actionInterface.InvokerNew("get_name")
	})
	return err
}

// GetName is a representation of the C type g_action_get_name.
func (recv *Action) GetName() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := actionGetNameFunction_Set()
	if err == nil {
		ret = actionGetNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'g_action_get_parameter_type' : return type 'GLib.VariantType' not supported

// UNSUPPORTED : C value 'g_action_get_state' : return type 'GLib.Variant' not supported

// UNSUPPORTED : C value 'g_action_get_state_hint' : return type 'GLib.Variant' not supported

// UNSUPPORTED : C value 'g_action_get_state_type' : return type 'GLib.VariantType' not supported

var actionGroupInterface *gi.Interface
var actionGroupInterface_Once sync.Once

func actionGroupInterface_Set() error {
	var err error
	actionGroupInterface_Once.Do(func() {
		actionGroupInterface, err = gi.InterfaceNew("Gio", "ActionGroup")
	})
	return err
}

type ActionGroup struct {
	native unsafe.Pointer
}

func ActionGroupNewFromNative(native unsafe.Pointer) *ActionGroup {
	instance := &ActionGroup{native: native}

	return instance
}

/*
CastToActionGroup down casts any arbitrary Object to ActionGroup.
Exercise care, as this is a potentially dangerous function
if the Object is not a ActionGroup.
*/
func CastToActionGroup(object *gobject.Object) *ActionGroup {
	return ActionGroupNewFromNative(object.Native())
}

// Equals compares this ActionGroup with another ActionGroup, and returns true if they represent the same Object.
func (recv *ActionGroup) Equals(other *ActionGroup) bool {
	return other.Native() == recv.Native()
}

func (recv *ActionGroup) Native() unsafe.Pointer {
	return recv.native
}

var actionGroupActionAddedFunction *gi.Function
var actionGroupActionAddedFunction_Once sync.Once

func actionGroupActionAddedFunction_Set() error {
	var err error
	actionGroupActionAddedFunction_Once.Do(func() {
		err = actionGroupInterface_Set()
		if err != nil {
			return
		}
		actionGroupActionAddedFunction, err = actionGroupInterface.InvokerNew("action_added")
	})
	return err
}

// ActionAdded is a representation of the C type g_action_group_action_added.
func (recv *ActionGroup) ActionAdded(actionName string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(actionName)

	err := actionGroupActionAddedFunction_Set()
	if err == nil {
		actionGroupActionAddedFunction.Invoke(inArgs[:], nil)
	}

	return
}

var actionGroupActionEnabledChangedFunction *gi.Function
var actionGroupActionEnabledChangedFunction_Once sync.Once

func actionGroupActionEnabledChangedFunction_Set() error {
	var err error
	actionGroupActionEnabledChangedFunction_Once.Do(func() {
		err = actionGroupInterface_Set()
		if err != nil {
			return
		}
		actionGroupActionEnabledChangedFunction, err = actionGroupInterface.InvokerNew("action_enabled_changed")
	})
	return err
}

// ActionEnabledChanged is a representation of the C type g_action_group_action_enabled_changed.
func (recv *ActionGroup) ActionEnabledChanged(actionName string, enabled bool) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(actionName)
	inArgs[2].SetBoolean(enabled)

	err := actionGroupActionEnabledChangedFunction_Set()
	if err == nil {
		actionGroupActionEnabledChangedFunction.Invoke(inArgs[:], nil)
	}

	return
}

var actionGroupActionRemovedFunction *gi.Function
var actionGroupActionRemovedFunction_Once sync.Once

func actionGroupActionRemovedFunction_Set() error {
	var err error
	actionGroupActionRemovedFunction_Once.Do(func() {
		err = actionGroupInterface_Set()
		if err != nil {
			return
		}
		actionGroupActionRemovedFunction, err = actionGroupInterface.InvokerNew("action_removed")
	})
	return err
}

// ActionRemoved is a representation of the C type g_action_group_action_removed.
func (recv *ActionGroup) ActionRemoved(actionName string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(actionName)

	err := actionGroupActionRemovedFunction_Set()
	if err == nil {
		actionGroupActionRemovedFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'g_action_group_action_state_changed' : parameter 'state' of type 'GLib.Variant' not supported

// UNSUPPORTED : C value 'g_action_group_activate_action' : parameter 'parameter' of type 'GLib.Variant' not supported

// UNSUPPORTED : C value 'g_action_group_change_action_state' : parameter 'value' of type 'GLib.Variant' not supported

var actionGroupGetActionEnabledFunction *gi.Function
var actionGroupGetActionEnabledFunction_Once sync.Once

func actionGroupGetActionEnabledFunction_Set() error {
	var err error
	actionGroupGetActionEnabledFunction_Once.Do(func() {
		err = actionGroupInterface_Set()
		if err != nil {
			return
		}
		actionGroupGetActionEnabledFunction, err = actionGroupInterface.InvokerNew("get_action_enabled")
	})
	return err
}

// GetActionEnabled is a representation of the C type g_action_group_get_action_enabled.
func (recv *ActionGroup) GetActionEnabled(actionName string) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(actionName)

	var ret gi.Argument

	err := actionGroupGetActionEnabledFunction_Set()
	if err == nil {
		ret = actionGroupGetActionEnabledFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'g_action_group_get_action_parameter_type' : return type 'GLib.VariantType' not supported

// UNSUPPORTED : C value 'g_action_group_get_action_state' : return type 'GLib.Variant' not supported

// UNSUPPORTED : C value 'g_action_group_get_action_state_hint' : return type 'GLib.Variant' not supported

// UNSUPPORTED : C value 'g_action_group_get_action_state_type' : return type 'GLib.VariantType' not supported

var actionGroupHasActionFunction *gi.Function
var actionGroupHasActionFunction_Once sync.Once

func actionGroupHasActionFunction_Set() error {
	var err error
	actionGroupHasActionFunction_Once.Do(func() {
		err = actionGroupInterface_Set()
		if err != nil {
			return
		}
		actionGroupHasActionFunction, err = actionGroupInterface.InvokerNew("has_action")
	})
	return err
}

// HasAction is a representation of the C type g_action_group_has_action.
func (recv *ActionGroup) HasAction(actionName string) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(actionName)

	var ret gi.Argument

	err := actionGroupHasActionFunction_Set()
	if err == nil {
		ret = actionGroupHasActionFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var actionGroupListActionsFunction *gi.Function
var actionGroupListActionsFunction_Once sync.Once

func actionGroupListActionsFunction_Set() error {
	var err error
	actionGroupListActionsFunction_Once.Do(func() {
		err = actionGroupInterface_Set()
		if err != nil {
			return
		}
		actionGroupListActionsFunction, err = actionGroupInterface.InvokerNew("list_actions")
	})
	return err
}

// ListActions is a representation of the C type g_action_group_list_actions.
func (recv *ActionGroup) ListActions() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := actionGroupListActionsFunction_Set()
	if err == nil {
		actionGroupListActionsFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'g_action_group_query_action' : parameter 'parameter_type' of type 'GLib.VariantType' not supported

/*
ConnectActionAdded connects a callback to the 'action-added' signal of the ActionGroup.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *ActionGroup) ConnectActionAdded(handler func(instance *ActionGroup, actionName string)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {}

	return callback.ConnectSignal(recv.Native(), "action-added", marshal)
}

/*
ConnectActionEnabledChanged connects a callback to the 'action-enabled-changed' signal of the ActionGroup.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *ActionGroup) ConnectActionEnabledChanged(handler func(instance *ActionGroup, actionName string, enabled bool)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {}

	return callback.ConnectSignal(recv.Native(), "action-enabled-changed", marshal)
}

/*
ConnectActionRemoved connects a callback to the 'action-removed' signal of the ActionGroup.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *ActionGroup) ConnectActionRemoved(handler func(instance *ActionGroup, actionName string)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {}

	return callback.ConnectSignal(recv.Native(), "action-removed", marshal)
}

// UNSUPPORTED : C value 'action-state-changed' : parameter 'value' of type 'GLib.Variant' not supported

/*
Disconnect disconnects a callback previously registered with a Connect...() method.

The connectionID should be a value returned from a call to a Connect...() method.
*/
func (recv *ActionGroup) DisconnectSignal(connectionID int) {
	callback.DisconnectSignal(connectionID)
}

var actionMapInterface *gi.Interface
var actionMapInterface_Once sync.Once

func actionMapInterface_Set() error {
	var err error
	actionMapInterface_Once.Do(func() {
		actionMapInterface, err = gi.InterfaceNew("Gio", "ActionMap")
	})
	return err
}

type ActionMap struct {
	native unsafe.Pointer
}

func ActionMapNewFromNative(native unsafe.Pointer) *ActionMap {
	instance := &ActionMap{native: native}

	return instance
}

/*
CastToActionMap down casts any arbitrary Object to ActionMap.
Exercise care, as this is a potentially dangerous function
if the Object is not a ActionMap.
*/
func CastToActionMap(object *gobject.Object) *ActionMap {
	return ActionMapNewFromNative(object.Native())
}

// Equals compares this ActionMap with another ActionMap, and returns true if they represent the same Object.
func (recv *ActionMap) Equals(other *ActionMap) bool {
	return other.Native() == recv.Native()
}

func (recv *ActionMap) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'g_action_map_add_action' : parameter 'action' of type 'Action' not supported

// UNSUPPORTED : C value 'g_action_map_add_action_entries' : parameter 'entries' of type 'nil' not supported

// UNSUPPORTED : C value 'g_action_map_lookup_action' : return type 'Action' not supported

var actionMapRemoveActionFunction *gi.Function
var actionMapRemoveActionFunction_Once sync.Once

func actionMapRemoveActionFunction_Set() error {
	var err error
	actionMapRemoveActionFunction_Once.Do(func() {
		err = actionMapInterface_Set()
		if err != nil {
			return
		}
		actionMapRemoveActionFunction, err = actionMapInterface.InvokerNew("remove_action")
	})
	return err
}

// RemoveAction is a representation of the C type g_action_map_remove_action.
func (recv *ActionMap) RemoveAction(actionName string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(actionName)

	err := actionMapRemoveActionFunction_Set()
	if err == nil {
		actionMapRemoveActionFunction.Invoke(inArgs[:], nil)
	}

	return
}

var appInfoInterface *gi.Interface
var appInfoInterface_Once sync.Once

func appInfoInterface_Set() error {
	var err error
	appInfoInterface_Once.Do(func() {
		appInfoInterface, err = gi.InterfaceNew("Gio", "AppInfo")
	})
	return err
}

type AppInfo struct {
	native unsafe.Pointer
}

func AppInfoNewFromNative(native unsafe.Pointer) *AppInfo {
	instance := &AppInfo{native: native}

	return instance
}

/*
CastToAppInfo down casts any arbitrary Object to AppInfo.
Exercise care, as this is a potentially dangerous function
if the Object is not a AppInfo.
*/
func CastToAppInfo(object *gobject.Object) *AppInfo {
	return AppInfoNewFromNative(object.Native())
}

// Equals compares this AppInfo with another AppInfo, and returns true if they represent the same Object.
func (recv *AppInfo) Equals(other *AppInfo) bool {
	return other.Native() == recv.Native()
}

func (recv *AppInfo) Native() unsafe.Pointer {
	return recv.native
}

var appInfoAddSupportsTypeFunction *gi.Function
var appInfoAddSupportsTypeFunction_Once sync.Once

func appInfoAddSupportsTypeFunction_Set() error {
	var err error
	appInfoAddSupportsTypeFunction_Once.Do(func() {
		err = appInfoInterface_Set()
		if err != nil {
			return
		}
		appInfoAddSupportsTypeFunction, err = appInfoInterface.InvokerNew("add_supports_type")
	})
	return err
}

// AddSupportsType is a representation of the C type g_app_info_add_supports_type.
func (recv *AppInfo) AddSupportsType(contentType string) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(contentType)

	var ret gi.Argument

	err := appInfoAddSupportsTypeFunction_Set()
	if err == nil {
		ret = appInfoAddSupportsTypeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var appInfoCanDeleteFunction *gi.Function
var appInfoCanDeleteFunction_Once sync.Once

func appInfoCanDeleteFunction_Set() error {
	var err error
	appInfoCanDeleteFunction_Once.Do(func() {
		err = appInfoInterface_Set()
		if err != nil {
			return
		}
		appInfoCanDeleteFunction, err = appInfoInterface.InvokerNew("can_delete")
	})
	return err
}

// CanDelete is a representation of the C type g_app_info_can_delete.
func (recv *AppInfo) CanDelete() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := appInfoCanDeleteFunction_Set()
	if err == nil {
		ret = appInfoCanDeleteFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var appInfoCanRemoveSupportsTypeFunction *gi.Function
var appInfoCanRemoveSupportsTypeFunction_Once sync.Once

func appInfoCanRemoveSupportsTypeFunction_Set() error {
	var err error
	appInfoCanRemoveSupportsTypeFunction_Once.Do(func() {
		err = appInfoInterface_Set()
		if err != nil {
			return
		}
		appInfoCanRemoveSupportsTypeFunction, err = appInfoInterface.InvokerNew("can_remove_supports_type")
	})
	return err
}

// CanRemoveSupportsType is a representation of the C type g_app_info_can_remove_supports_type.
func (recv *AppInfo) CanRemoveSupportsType() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := appInfoCanRemoveSupportsTypeFunction_Set()
	if err == nil {
		ret = appInfoCanRemoveSupportsTypeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var appInfoDeleteFunction *gi.Function
var appInfoDeleteFunction_Once sync.Once

func appInfoDeleteFunction_Set() error {
	var err error
	appInfoDeleteFunction_Once.Do(func() {
		err = appInfoInterface_Set()
		if err != nil {
			return
		}
		appInfoDeleteFunction, err = appInfoInterface.InvokerNew("delete")
	})
	return err
}

// Delete is a representation of the C type g_app_info_delete.
func (recv *AppInfo) Delete() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := appInfoDeleteFunction_Set()
	if err == nil {
		ret = appInfoDeleteFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'g_app_info_dup' : return type 'AppInfo' not supported

// UNSUPPORTED : C value 'g_app_info_equal' : parameter 'appinfo2' of type 'AppInfo' not supported

var appInfoGetCommandlineFunction *gi.Function
var appInfoGetCommandlineFunction_Once sync.Once

func appInfoGetCommandlineFunction_Set() error {
	var err error
	appInfoGetCommandlineFunction_Once.Do(func() {
		err = appInfoInterface_Set()
		if err != nil {
			return
		}
		appInfoGetCommandlineFunction, err = appInfoInterface.InvokerNew("get_commandline")
	})
	return err
}

// GetCommandline is a representation of the C type g_app_info_get_commandline.
func (recv *AppInfo) GetCommandline() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := appInfoGetCommandlineFunction_Set()
	if err == nil {
		ret = appInfoGetCommandlineFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var appInfoGetDescriptionFunction *gi.Function
var appInfoGetDescriptionFunction_Once sync.Once

func appInfoGetDescriptionFunction_Set() error {
	var err error
	appInfoGetDescriptionFunction_Once.Do(func() {
		err = appInfoInterface_Set()
		if err != nil {
			return
		}
		appInfoGetDescriptionFunction, err = appInfoInterface.InvokerNew("get_description")
	})
	return err
}

// GetDescription is a representation of the C type g_app_info_get_description.
func (recv *AppInfo) GetDescription() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := appInfoGetDescriptionFunction_Set()
	if err == nil {
		ret = appInfoGetDescriptionFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var appInfoGetDisplayNameFunction *gi.Function
var appInfoGetDisplayNameFunction_Once sync.Once

func appInfoGetDisplayNameFunction_Set() error {
	var err error
	appInfoGetDisplayNameFunction_Once.Do(func() {
		err = appInfoInterface_Set()
		if err != nil {
			return
		}
		appInfoGetDisplayNameFunction, err = appInfoInterface.InvokerNew("get_display_name")
	})
	return err
}

// GetDisplayName is a representation of the C type g_app_info_get_display_name.
func (recv *AppInfo) GetDisplayName() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := appInfoGetDisplayNameFunction_Set()
	if err == nil {
		ret = appInfoGetDisplayNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var appInfoGetExecutableFunction *gi.Function
var appInfoGetExecutableFunction_Once sync.Once

func appInfoGetExecutableFunction_Set() error {
	var err error
	appInfoGetExecutableFunction_Once.Do(func() {
		err = appInfoInterface_Set()
		if err != nil {
			return
		}
		appInfoGetExecutableFunction, err = appInfoInterface.InvokerNew("get_executable")
	})
	return err
}

// GetExecutable is a representation of the C type g_app_info_get_executable.
func (recv *AppInfo) GetExecutable() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := appInfoGetExecutableFunction_Set()
	if err == nil {
		ret = appInfoGetExecutableFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'g_app_info_get_icon' : return type 'Icon' not supported

var appInfoGetIdFunction *gi.Function
var appInfoGetIdFunction_Once sync.Once

func appInfoGetIdFunction_Set() error {
	var err error
	appInfoGetIdFunction_Once.Do(func() {
		err = appInfoInterface_Set()
		if err != nil {
			return
		}
		appInfoGetIdFunction, err = appInfoInterface.InvokerNew("get_id")
	})
	return err
}

// GetId is a representation of the C type g_app_info_get_id.
func (recv *AppInfo) GetId() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := appInfoGetIdFunction_Set()
	if err == nil {
		ret = appInfoGetIdFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var appInfoGetNameFunction *gi.Function
var appInfoGetNameFunction_Once sync.Once

func appInfoGetNameFunction_Set() error {
	var err error
	appInfoGetNameFunction_Once.Do(func() {
		err = appInfoInterface_Set()
		if err != nil {
			return
		}
		appInfoGetNameFunction, err = appInfoInterface.InvokerNew("get_name")
	})
	return err
}

// GetName is a representation of the C type g_app_info_get_name.
func (recv *AppInfo) GetName() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := appInfoGetNameFunction_Set()
	if err == nil {
		ret = appInfoGetNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var appInfoGetSupportedTypesFunction *gi.Function
var appInfoGetSupportedTypesFunction_Once sync.Once

func appInfoGetSupportedTypesFunction_Set() error {
	var err error
	appInfoGetSupportedTypesFunction_Once.Do(func() {
		err = appInfoInterface_Set()
		if err != nil {
			return
		}
		appInfoGetSupportedTypesFunction, err = appInfoInterface.InvokerNew("get_supported_types")
	})
	return err
}

// GetSupportedTypes is a representation of the C type g_app_info_get_supported_types.
func (recv *AppInfo) GetSupportedTypes() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := appInfoGetSupportedTypesFunction_Set()
	if err == nil {
		appInfoGetSupportedTypesFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'g_app_info_launch' : parameter 'files' of type 'GLib.List' not supported

// UNSUPPORTED : C value 'g_app_info_launch_uris' : parameter 'uris' of type 'GLib.List' not supported

// UNSUPPORTED : C value 'g_app_info_launch_uris_async' : parameter 'uris' of type 'GLib.List' not supported

// UNSUPPORTED : C value 'g_app_info_launch_uris_finish' : parameter 'result' of type 'AsyncResult' not supported

var appInfoRemoveSupportsTypeFunction *gi.Function
var appInfoRemoveSupportsTypeFunction_Once sync.Once

func appInfoRemoveSupportsTypeFunction_Set() error {
	var err error
	appInfoRemoveSupportsTypeFunction_Once.Do(func() {
		err = appInfoInterface_Set()
		if err != nil {
			return
		}
		appInfoRemoveSupportsTypeFunction, err = appInfoInterface.InvokerNew("remove_supports_type")
	})
	return err
}

// RemoveSupportsType is a representation of the C type g_app_info_remove_supports_type.
func (recv *AppInfo) RemoveSupportsType(contentType string) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(contentType)

	var ret gi.Argument

	err := appInfoRemoveSupportsTypeFunction_Set()
	if err == nil {
		ret = appInfoRemoveSupportsTypeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var appInfoSetAsDefaultForExtensionFunction *gi.Function
var appInfoSetAsDefaultForExtensionFunction_Once sync.Once

func appInfoSetAsDefaultForExtensionFunction_Set() error {
	var err error
	appInfoSetAsDefaultForExtensionFunction_Once.Do(func() {
		err = appInfoInterface_Set()
		if err != nil {
			return
		}
		appInfoSetAsDefaultForExtensionFunction, err = appInfoInterface.InvokerNew("set_as_default_for_extension")
	})
	return err
}

// SetAsDefaultForExtension is a representation of the C type g_app_info_set_as_default_for_extension.
func (recv *AppInfo) SetAsDefaultForExtension(extension string) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(extension)

	var ret gi.Argument

	err := appInfoSetAsDefaultForExtensionFunction_Set()
	if err == nil {
		ret = appInfoSetAsDefaultForExtensionFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var appInfoSetAsDefaultForTypeFunction *gi.Function
var appInfoSetAsDefaultForTypeFunction_Once sync.Once

func appInfoSetAsDefaultForTypeFunction_Set() error {
	var err error
	appInfoSetAsDefaultForTypeFunction_Once.Do(func() {
		err = appInfoInterface_Set()
		if err != nil {
			return
		}
		appInfoSetAsDefaultForTypeFunction, err = appInfoInterface.InvokerNew("set_as_default_for_type")
	})
	return err
}

// SetAsDefaultForType is a representation of the C type g_app_info_set_as_default_for_type.
func (recv *AppInfo) SetAsDefaultForType(contentType string) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(contentType)

	var ret gi.Argument

	err := appInfoSetAsDefaultForTypeFunction_Set()
	if err == nil {
		ret = appInfoSetAsDefaultForTypeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var appInfoSetAsLastUsedForTypeFunction *gi.Function
var appInfoSetAsLastUsedForTypeFunction_Once sync.Once

func appInfoSetAsLastUsedForTypeFunction_Set() error {
	var err error
	appInfoSetAsLastUsedForTypeFunction_Once.Do(func() {
		err = appInfoInterface_Set()
		if err != nil {
			return
		}
		appInfoSetAsLastUsedForTypeFunction, err = appInfoInterface.InvokerNew("set_as_last_used_for_type")
	})
	return err
}

// SetAsLastUsedForType is a representation of the C type g_app_info_set_as_last_used_for_type.
func (recv *AppInfo) SetAsLastUsedForType(contentType string) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(contentType)

	var ret gi.Argument

	err := appInfoSetAsLastUsedForTypeFunction_Set()
	if err == nil {
		ret = appInfoSetAsLastUsedForTypeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var appInfoShouldShowFunction *gi.Function
var appInfoShouldShowFunction_Once sync.Once

func appInfoShouldShowFunction_Set() error {
	var err error
	appInfoShouldShowFunction_Once.Do(func() {
		err = appInfoInterface_Set()
		if err != nil {
			return
		}
		appInfoShouldShowFunction, err = appInfoInterface.InvokerNew("should_show")
	})
	return err
}

// ShouldShow is a representation of the C type g_app_info_should_show.
func (recv *AppInfo) ShouldShow() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := appInfoShouldShowFunction_Set()
	if err == nil {
		ret = appInfoShouldShowFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var appInfoSupportsFilesFunction *gi.Function
var appInfoSupportsFilesFunction_Once sync.Once

func appInfoSupportsFilesFunction_Set() error {
	var err error
	appInfoSupportsFilesFunction_Once.Do(func() {
		err = appInfoInterface_Set()
		if err != nil {
			return
		}
		appInfoSupportsFilesFunction, err = appInfoInterface.InvokerNew("supports_files")
	})
	return err
}

// SupportsFiles is a representation of the C type g_app_info_supports_files.
func (recv *AppInfo) SupportsFiles() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := appInfoSupportsFilesFunction_Set()
	if err == nil {
		ret = appInfoSupportsFilesFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var appInfoSupportsUrisFunction *gi.Function
var appInfoSupportsUrisFunction_Once sync.Once

func appInfoSupportsUrisFunction_Set() error {
	var err error
	appInfoSupportsUrisFunction_Once.Do(func() {
		err = appInfoInterface_Set()
		if err != nil {
			return
		}
		appInfoSupportsUrisFunction, err = appInfoInterface.InvokerNew("supports_uris")
	})
	return err
}

// SupportsUris is a representation of the C type g_app_info_supports_uris.
func (recv *AppInfo) SupportsUris() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := appInfoSupportsUrisFunction_Set()
	if err == nil {
		ret = appInfoSupportsUrisFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var asyncInitableInterface *gi.Interface
var asyncInitableInterface_Once sync.Once

func asyncInitableInterface_Set() error {
	var err error
	asyncInitableInterface_Once.Do(func() {
		asyncInitableInterface, err = gi.InterfaceNew("Gio", "AsyncInitable")
	})
	return err
}

type AsyncInitable struct {
	native unsafe.Pointer
}

func AsyncInitableNewFromNative(native unsafe.Pointer) *AsyncInitable {
	instance := &AsyncInitable{native: native}

	return instance
}

/*
CastToAsyncInitable down casts any arbitrary Object to AsyncInitable.
Exercise care, as this is a potentially dangerous function
if the Object is not a AsyncInitable.
*/
func CastToAsyncInitable(object *gobject.Object) *AsyncInitable {
	return AsyncInitableNewFromNative(object.Native())
}

// Equals compares this AsyncInitable with another AsyncInitable, and returns true if they represent the same Object.
func (recv *AsyncInitable) Equals(other *AsyncInitable) bool {
	return other.Native() == recv.Native()
}

func (recv *AsyncInitable) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'g_async_initable_init_async' : parameter 'callback' of type 'AsyncReadyCallback' not supported

// UNSUPPORTED : C value 'g_async_initable_init_finish' : parameter 'res' of type 'AsyncResult' not supported

// UNSUPPORTED : C value 'g_async_initable_new_finish' : parameter 'res' of type 'AsyncResult' not supported

var asyncResultInterface *gi.Interface
var asyncResultInterface_Once sync.Once

func asyncResultInterface_Set() error {
	var err error
	asyncResultInterface_Once.Do(func() {
		asyncResultInterface, err = gi.InterfaceNew("Gio", "AsyncResult")
	})
	return err
}

type AsyncResult struct {
	native unsafe.Pointer
}

func AsyncResultNewFromNative(native unsafe.Pointer) *AsyncResult {
	instance := &AsyncResult{native: native}

	return instance
}

/*
CastToAsyncResult down casts any arbitrary Object to AsyncResult.
Exercise care, as this is a potentially dangerous function
if the Object is not a AsyncResult.
*/
func CastToAsyncResult(object *gobject.Object) *AsyncResult {
	return AsyncResultNewFromNative(object.Native())
}

// Equals compares this AsyncResult with another AsyncResult, and returns true if they represent the same Object.
func (recv *AsyncResult) Equals(other *AsyncResult) bool {
	return other.Native() == recv.Native()
}

func (recv *AsyncResult) Native() unsafe.Pointer {
	return recv.native
}

var asyncResultGetSourceObjectFunction *gi.Function
var asyncResultGetSourceObjectFunction_Once sync.Once

func asyncResultGetSourceObjectFunction_Set() error {
	var err error
	asyncResultGetSourceObjectFunction_Once.Do(func() {
		err = asyncResultInterface_Set()
		if err != nil {
			return
		}
		asyncResultGetSourceObjectFunction, err = asyncResultInterface.InvokerNew("get_source_object")
	})
	return err
}

// GetSourceObject is a representation of the C type g_async_result_get_source_object.
func (recv *AsyncResult) GetSourceObject() *gobject.Object {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := asyncResultGetSourceObjectFunction_Set()
	if err == nil {
		ret = asyncResultGetSourceObjectFunction.Invoke(inArgs[:], nil)
	}

	retGo := gobject.ObjectNewFromNative(ret.Pointer())

	return retGo
}

// UNSUPPORTED : C value 'g_async_result_get_user_data' : return type 'gpointer' not supported

// UNSUPPORTED : C value 'g_async_result_is_tagged' : parameter 'source_tag' of type 'gpointer' not supported

var asyncResultLegacyPropagateErrorFunction *gi.Function
var asyncResultLegacyPropagateErrorFunction_Once sync.Once

func asyncResultLegacyPropagateErrorFunction_Set() error {
	var err error
	asyncResultLegacyPropagateErrorFunction_Once.Do(func() {
		err = asyncResultInterface_Set()
		if err != nil {
			return
		}
		asyncResultLegacyPropagateErrorFunction, err = asyncResultInterface.InvokerNew("legacy_propagate_error")
	})
	return err
}

// LegacyPropagateError is a representation of the C type g_async_result_legacy_propagate_error.
func (recv *AsyncResult) LegacyPropagateError() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := asyncResultLegacyPropagateErrorFunction_Set()
	if err == nil {
		ret = asyncResultLegacyPropagateErrorFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var converterInterface *gi.Interface
var converterInterface_Once sync.Once

func converterInterface_Set() error {
	var err error
	converterInterface_Once.Do(func() {
		converterInterface, err = gi.InterfaceNew("Gio", "Converter")
	})
	return err
}

type Converter struct {
	native unsafe.Pointer
}

func ConverterNewFromNative(native unsafe.Pointer) *Converter {
	instance := &Converter{native: native}

	return instance
}

/*
CastToConverter down casts any arbitrary Object to Converter.
Exercise care, as this is a potentially dangerous function
if the Object is not a Converter.
*/
func CastToConverter(object *gobject.Object) *Converter {
	return ConverterNewFromNative(object.Native())
}

// Equals compares this Converter with another Converter, and returns true if they represent the same Object.
func (recv *Converter) Equals(other *Converter) bool {
	return other.Native() == recv.Native()
}

func (recv *Converter) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'g_converter_convert' : parameter 'inbuf' of type 'nil' not supported

var converterResetFunction *gi.Function
var converterResetFunction_Once sync.Once

func converterResetFunction_Set() error {
	var err error
	converterResetFunction_Once.Do(func() {
		err = converterInterface_Set()
		if err != nil {
			return
		}
		converterResetFunction, err = converterInterface.InvokerNew("reset")
	})
	return err
}

// Reset is a representation of the C type g_converter_reset.
func (recv *Converter) Reset() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := converterResetFunction_Set()
	if err == nil {
		converterResetFunction.Invoke(inArgs[:], nil)
	}

	return
}

var dBusInterfaceInterface *gi.Interface
var dBusInterfaceInterface_Once sync.Once

func dBusInterfaceInterface_Set() error {
	var err error
	dBusInterfaceInterface_Once.Do(func() {
		dBusInterfaceInterface, err = gi.InterfaceNew("Gio", "DBusInterface")
	})
	return err
}

type DBusInterface struct {
	native unsafe.Pointer
}

func DBusInterfaceNewFromNative(native unsafe.Pointer) *DBusInterface {
	instance := &DBusInterface{native: native}

	return instance
}

/*
CastToDBusInterface down casts any arbitrary Object to DBusInterface.
Exercise care, as this is a potentially dangerous function
if the Object is not a DBusInterface.
*/
func CastToDBusInterface(object *gobject.Object) *DBusInterface {
	return DBusInterfaceNewFromNative(object.Native())
}

// Equals compares this DBusInterface with another DBusInterface, and returns true if they represent the same Object.
func (recv *DBusInterface) Equals(other *DBusInterface) bool {
	return other.Native() == recv.Native()
}

func (recv *DBusInterface) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'g_dbus_interface_dup_object' : return type 'DBusObject' not supported

var dBusInterfaceGetInfoFunction *gi.Function
var dBusInterfaceGetInfoFunction_Once sync.Once

func dBusInterfaceGetInfoFunction_Set() error {
	var err error
	dBusInterfaceGetInfoFunction_Once.Do(func() {
		err = dBusInterfaceInterface_Set()
		if err != nil {
			return
		}
		dBusInterfaceGetInfoFunction, err = dBusInterfaceInterface.InvokerNew("get_info")
	})
	return err
}

// GetInfo is a representation of the C type g_dbus_interface_get_info.
func (recv *DBusInterface) GetInfo() *DBusInterfaceInfo {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := dBusInterfaceGetInfoFunction_Set()
	if err == nil {
		ret = dBusInterfaceGetInfoFunction.Invoke(inArgs[:], nil)
	}

	retGo := DBusInterfaceInfoNewFromNative(ret.Pointer())

	return retGo
}

// UNSUPPORTED : C value 'g_dbus_interface_get_object' : return type 'DBusObject' not supported

// UNSUPPORTED : C value 'g_dbus_interface_set_object' : parameter 'object' of type 'DBusObject' not supported

var dBusObjectInterface *gi.Interface
var dBusObjectInterface_Once sync.Once

func dBusObjectInterface_Set() error {
	var err error
	dBusObjectInterface_Once.Do(func() {
		dBusObjectInterface, err = gi.InterfaceNew("Gio", "DBusObject")
	})
	return err
}

type DBusObject struct {
	native unsafe.Pointer
}

func DBusObjectNewFromNative(native unsafe.Pointer) *DBusObject {
	instance := &DBusObject{native: native}

	return instance
}

/*
CastToDBusObject down casts any arbitrary Object to DBusObject.
Exercise care, as this is a potentially dangerous function
if the Object is not a DBusObject.
*/
func CastToDBusObject(object *gobject.Object) *DBusObject {
	return DBusObjectNewFromNative(object.Native())
}

// Equals compares this DBusObject with another DBusObject, and returns true if they represent the same Object.
func (recv *DBusObject) Equals(other *DBusObject) bool {
	return other.Native() == recv.Native()
}

func (recv *DBusObject) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'g_dbus_object_get_interface' : return type 'DBusInterface' not supported

// UNSUPPORTED : C value 'g_dbus_object_get_interfaces' : return type 'GLib.List' not supported

var dBusObjectGetObjectPathFunction *gi.Function
var dBusObjectGetObjectPathFunction_Once sync.Once

func dBusObjectGetObjectPathFunction_Set() error {
	var err error
	dBusObjectGetObjectPathFunction_Once.Do(func() {
		err = dBusObjectInterface_Set()
		if err != nil {
			return
		}
		dBusObjectGetObjectPathFunction, err = dBusObjectInterface.InvokerNew("get_object_path")
	})
	return err
}

// GetObjectPath is a representation of the C type g_dbus_object_get_object_path.
func (recv *DBusObject) GetObjectPath() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := dBusObjectGetObjectPathFunction_Set()
	if err == nil {
		ret = dBusObjectGetObjectPathFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'interface-added' : parameter 'interface' of type 'DBusInterface' not supported

// UNSUPPORTED : C value 'interface-removed' : parameter 'interface' of type 'DBusInterface' not supported

/*
Disconnect disconnects a callback previously registered with a Connect...() method.

The connectionID should be a value returned from a call to a Connect...() method.
*/
func (recv *DBusObject) DisconnectSignal(connectionID int) {
	callback.DisconnectSignal(connectionID)
}

var dBusObjectManagerInterface *gi.Interface
var dBusObjectManagerInterface_Once sync.Once

func dBusObjectManagerInterface_Set() error {
	var err error
	dBusObjectManagerInterface_Once.Do(func() {
		dBusObjectManagerInterface, err = gi.InterfaceNew("Gio", "DBusObjectManager")
	})
	return err
}

type DBusObjectManager struct {
	native unsafe.Pointer
}

func DBusObjectManagerNewFromNative(native unsafe.Pointer) *DBusObjectManager {
	instance := &DBusObjectManager{native: native}

	return instance
}

/*
CastToDBusObjectManager down casts any arbitrary Object to DBusObjectManager.
Exercise care, as this is a potentially dangerous function
if the Object is not a DBusObjectManager.
*/
func CastToDBusObjectManager(object *gobject.Object) *DBusObjectManager {
	return DBusObjectManagerNewFromNative(object.Native())
}

// Equals compares this DBusObjectManager with another DBusObjectManager, and returns true if they represent the same Object.
func (recv *DBusObjectManager) Equals(other *DBusObjectManager) bool {
	return other.Native() == recv.Native()
}

func (recv *DBusObjectManager) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'g_dbus_object_manager_get_interface' : return type 'DBusInterface' not supported

// UNSUPPORTED : C value 'g_dbus_object_manager_get_object' : return type 'DBusObject' not supported

var dBusObjectManagerGetObjectPathFunction *gi.Function
var dBusObjectManagerGetObjectPathFunction_Once sync.Once

func dBusObjectManagerGetObjectPathFunction_Set() error {
	var err error
	dBusObjectManagerGetObjectPathFunction_Once.Do(func() {
		err = dBusObjectManagerInterface_Set()
		if err != nil {
			return
		}
		dBusObjectManagerGetObjectPathFunction, err = dBusObjectManagerInterface.InvokerNew("get_object_path")
	})
	return err
}

// GetObjectPath is a representation of the C type g_dbus_object_manager_get_object_path.
func (recv *DBusObjectManager) GetObjectPath() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := dBusObjectManagerGetObjectPathFunction_Set()
	if err == nil {
		ret = dBusObjectManagerGetObjectPathFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'g_dbus_object_manager_get_objects' : return type 'GLib.List' not supported

// UNSUPPORTED : C value 'interface-added' : parameter 'object' of type 'DBusObject' not supported

// UNSUPPORTED : C value 'interface-removed' : parameter 'object' of type 'DBusObject' not supported

// UNSUPPORTED : C value 'object-added' : parameter 'object' of type 'DBusObject' not supported

// UNSUPPORTED : C value 'object-removed' : parameter 'object' of type 'DBusObject' not supported

/*
Disconnect disconnects a callback previously registered with a Connect...() method.

The connectionID should be a value returned from a call to a Connect...() method.
*/
func (recv *DBusObjectManager) DisconnectSignal(connectionID int) {
	callback.DisconnectSignal(connectionID)
}

var datagramBasedInterface *gi.Interface
var datagramBasedInterface_Once sync.Once

func datagramBasedInterface_Set() error {
	var err error
	datagramBasedInterface_Once.Do(func() {
		datagramBasedInterface, err = gi.InterfaceNew("Gio", "DatagramBased")
	})
	return err
}

type DatagramBased struct {
	native unsafe.Pointer
}

func DatagramBasedNewFromNative(native unsafe.Pointer) *DatagramBased {
	instance := &DatagramBased{native: native}

	return instance
}

/*
CastToDatagramBased down casts any arbitrary Object to DatagramBased.
Exercise care, as this is a potentially dangerous function
if the Object is not a DatagramBased.
*/
func CastToDatagramBased(object *gobject.Object) *DatagramBased {
	return DatagramBasedNewFromNative(object.Native())
}

// Equals compares this DatagramBased with another DatagramBased, and returns true if they represent the same Object.
func (recv *DatagramBased) Equals(other *DatagramBased) bool {
	return other.Native() == recv.Native()
}

func (recv *DatagramBased) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'g_datagram_based_condition_check' : parameter 'condition' of type 'GLib.IOCondition' not supported

// UNSUPPORTED : C value 'g_datagram_based_condition_wait' : parameter 'condition' of type 'GLib.IOCondition' not supported

// UNSUPPORTED : C value 'g_datagram_based_create_source' : parameter 'condition' of type 'GLib.IOCondition' not supported

// UNSUPPORTED : C value 'g_datagram_based_receive_messages' : parameter 'messages' of type 'nil' not supported

// UNSUPPORTED : C value 'g_datagram_based_send_messages' : parameter 'messages' of type 'nil' not supported

var desktopAppInfoLookupInterface *gi.Interface
var desktopAppInfoLookupInterface_Once sync.Once

func desktopAppInfoLookupInterface_Set() error {
	var err error
	desktopAppInfoLookupInterface_Once.Do(func() {
		desktopAppInfoLookupInterface, err = gi.InterfaceNew("Gio", "DesktopAppInfoLookup")
	})
	return err
}

type DesktopAppInfoLookup struct {
	native unsafe.Pointer
}

func DesktopAppInfoLookupNewFromNative(native unsafe.Pointer) *DesktopAppInfoLookup {
	instance := &DesktopAppInfoLookup{native: native}

	return instance
}

/*
CastToDesktopAppInfoLookup down casts any arbitrary Object to DesktopAppInfoLookup.
Exercise care, as this is a potentially dangerous function
if the Object is not a DesktopAppInfoLookup.
*/
func CastToDesktopAppInfoLookup(object *gobject.Object) *DesktopAppInfoLookup {
	return DesktopAppInfoLookupNewFromNative(object.Native())
}

// Equals compares this DesktopAppInfoLookup with another DesktopAppInfoLookup, and returns true if they represent the same Object.
func (recv *DesktopAppInfoLookup) Equals(other *DesktopAppInfoLookup) bool {
	return other.Native() == recv.Native()
}

func (recv *DesktopAppInfoLookup) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'g_desktop_app_info_lookup_get_default_for_uri_scheme' : return type 'AppInfo' not supported

var driveInterface *gi.Interface
var driveInterface_Once sync.Once

func driveInterface_Set() error {
	var err error
	driveInterface_Once.Do(func() {
		driveInterface, err = gi.InterfaceNew("Gio", "Drive")
	})
	return err
}

type Drive struct {
	native unsafe.Pointer
}

func DriveNewFromNative(native unsafe.Pointer) *Drive {
	instance := &Drive{native: native}

	return instance
}

/*
CastToDrive down casts any arbitrary Object to Drive.
Exercise care, as this is a potentially dangerous function
if the Object is not a Drive.
*/
func CastToDrive(object *gobject.Object) *Drive {
	return DriveNewFromNative(object.Native())
}

// Equals compares this Drive with another Drive, and returns true if they represent the same Object.
func (recv *Drive) Equals(other *Drive) bool {
	return other.Native() == recv.Native()
}

func (recv *Drive) Native() unsafe.Pointer {
	return recv.native
}

var driveCanEjectFunction *gi.Function
var driveCanEjectFunction_Once sync.Once

func driveCanEjectFunction_Set() error {
	var err error
	driveCanEjectFunction_Once.Do(func() {
		err = driveInterface_Set()
		if err != nil {
			return
		}
		driveCanEjectFunction, err = driveInterface.InvokerNew("can_eject")
	})
	return err
}

// CanEject is a representation of the C type g_drive_can_eject.
func (recv *Drive) CanEject() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := driveCanEjectFunction_Set()
	if err == nil {
		ret = driveCanEjectFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var driveCanPollForMediaFunction *gi.Function
var driveCanPollForMediaFunction_Once sync.Once

func driveCanPollForMediaFunction_Set() error {
	var err error
	driveCanPollForMediaFunction_Once.Do(func() {
		err = driveInterface_Set()
		if err != nil {
			return
		}
		driveCanPollForMediaFunction, err = driveInterface.InvokerNew("can_poll_for_media")
	})
	return err
}

// CanPollForMedia is a representation of the C type g_drive_can_poll_for_media.
func (recv *Drive) CanPollForMedia() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := driveCanPollForMediaFunction_Set()
	if err == nil {
		ret = driveCanPollForMediaFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var driveCanStartFunction *gi.Function
var driveCanStartFunction_Once sync.Once

func driveCanStartFunction_Set() error {
	var err error
	driveCanStartFunction_Once.Do(func() {
		err = driveInterface_Set()
		if err != nil {
			return
		}
		driveCanStartFunction, err = driveInterface.InvokerNew("can_start")
	})
	return err
}

// CanStart is a representation of the C type g_drive_can_start.
func (recv *Drive) CanStart() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := driveCanStartFunction_Set()
	if err == nil {
		ret = driveCanStartFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var driveCanStartDegradedFunction *gi.Function
var driveCanStartDegradedFunction_Once sync.Once

func driveCanStartDegradedFunction_Set() error {
	var err error
	driveCanStartDegradedFunction_Once.Do(func() {
		err = driveInterface_Set()
		if err != nil {
			return
		}
		driveCanStartDegradedFunction, err = driveInterface.InvokerNew("can_start_degraded")
	})
	return err
}

// CanStartDegraded is a representation of the C type g_drive_can_start_degraded.
func (recv *Drive) CanStartDegraded() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := driveCanStartDegradedFunction_Set()
	if err == nil {
		ret = driveCanStartDegradedFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var driveCanStopFunction *gi.Function
var driveCanStopFunction_Once sync.Once

func driveCanStopFunction_Set() error {
	var err error
	driveCanStopFunction_Once.Do(func() {
		err = driveInterface_Set()
		if err != nil {
			return
		}
		driveCanStopFunction, err = driveInterface.InvokerNew("can_stop")
	})
	return err
}

// CanStop is a representation of the C type g_drive_can_stop.
func (recv *Drive) CanStop() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := driveCanStopFunction_Set()
	if err == nil {
		ret = driveCanStopFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'g_drive_eject' : parameter 'flags' of type 'MountUnmountFlags' not supported

// UNSUPPORTED : C value 'g_drive_eject_finish' : parameter 'result' of type 'AsyncResult' not supported

// UNSUPPORTED : C value 'g_drive_eject_with_operation' : parameter 'flags' of type 'MountUnmountFlags' not supported

// UNSUPPORTED : C value 'g_drive_eject_with_operation_finish' : parameter 'result' of type 'AsyncResult' not supported

var driveEnumerateIdentifiersFunction *gi.Function
var driveEnumerateIdentifiersFunction_Once sync.Once

func driveEnumerateIdentifiersFunction_Set() error {
	var err error
	driveEnumerateIdentifiersFunction_Once.Do(func() {
		err = driveInterface_Set()
		if err != nil {
			return
		}
		driveEnumerateIdentifiersFunction, err = driveInterface.InvokerNew("enumerate_identifiers")
	})
	return err
}

// EnumerateIdentifiers is a representation of the C type g_drive_enumerate_identifiers.
func (recv *Drive) EnumerateIdentifiers() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := driveEnumerateIdentifiersFunction_Set()
	if err == nil {
		driveEnumerateIdentifiersFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'g_drive_get_icon' : return type 'Icon' not supported

var driveGetIdentifierFunction *gi.Function
var driveGetIdentifierFunction_Once sync.Once

func driveGetIdentifierFunction_Set() error {
	var err error
	driveGetIdentifierFunction_Once.Do(func() {
		err = driveInterface_Set()
		if err != nil {
			return
		}
		driveGetIdentifierFunction, err = driveInterface.InvokerNew("get_identifier")
	})
	return err
}

// GetIdentifier is a representation of the C type g_drive_get_identifier.
func (recv *Drive) GetIdentifier(kind string) string {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(kind)

	var ret gi.Argument

	err := driveGetIdentifierFunction_Set()
	if err == nil {
		ret = driveGetIdentifierFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

var driveGetNameFunction *gi.Function
var driveGetNameFunction_Once sync.Once

func driveGetNameFunction_Set() error {
	var err error
	driveGetNameFunction_Once.Do(func() {
		err = driveInterface_Set()
		if err != nil {
			return
		}
		driveGetNameFunction, err = driveInterface.InvokerNew("get_name")
	})
	return err
}

// GetName is a representation of the C type g_drive_get_name.
func (recv *Drive) GetName() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := driveGetNameFunction_Set()
	if err == nil {
		ret = driveGetNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

var driveGetSortKeyFunction *gi.Function
var driveGetSortKeyFunction_Once sync.Once

func driveGetSortKeyFunction_Set() error {
	var err error
	driveGetSortKeyFunction_Once.Do(func() {
		err = driveInterface_Set()
		if err != nil {
			return
		}
		driveGetSortKeyFunction, err = driveInterface.InvokerNew("get_sort_key")
	})
	return err
}

// GetSortKey is a representation of the C type g_drive_get_sort_key.
func (recv *Drive) GetSortKey() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := driveGetSortKeyFunction_Set()
	if err == nil {
		ret = driveGetSortKeyFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var driveGetStartStopTypeFunction *gi.Function
var driveGetStartStopTypeFunction_Once sync.Once

func driveGetStartStopTypeFunction_Set() error {
	var err error
	driveGetStartStopTypeFunction_Once.Do(func() {
		err = driveInterface_Set()
		if err != nil {
			return
		}
		driveGetStartStopTypeFunction, err = driveInterface.InvokerNew("get_start_stop_type")
	})
	return err
}

// GetStartStopType is a representation of the C type g_drive_get_start_stop_type.
func (recv *Drive) GetStartStopType() DriveStartStopType {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := driveGetStartStopTypeFunction_Set()
	if err == nil {
		ret = driveGetStartStopTypeFunction.Invoke(inArgs[:], nil)
	}

	retGo := DriveStartStopType(ret.Int32())

	return retGo
}

// UNSUPPORTED : C value 'g_drive_get_symbolic_icon' : return type 'Icon' not supported

// UNSUPPORTED : C value 'g_drive_get_volumes' : return type 'GLib.List' not supported

var driveHasMediaFunction *gi.Function
var driveHasMediaFunction_Once sync.Once

func driveHasMediaFunction_Set() error {
	var err error
	driveHasMediaFunction_Once.Do(func() {
		err = driveInterface_Set()
		if err != nil {
			return
		}
		driveHasMediaFunction, err = driveInterface.InvokerNew("has_media")
	})
	return err
}

// HasMedia is a representation of the C type g_drive_has_media.
func (recv *Drive) HasMedia() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := driveHasMediaFunction_Set()
	if err == nil {
		ret = driveHasMediaFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var driveHasVolumesFunction *gi.Function
var driveHasVolumesFunction_Once sync.Once

func driveHasVolumesFunction_Set() error {
	var err error
	driveHasVolumesFunction_Once.Do(func() {
		err = driveInterface_Set()
		if err != nil {
			return
		}
		driveHasVolumesFunction, err = driveInterface.InvokerNew("has_volumes")
	})
	return err
}

// HasVolumes is a representation of the C type g_drive_has_volumes.
func (recv *Drive) HasVolumes() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := driveHasVolumesFunction_Set()
	if err == nil {
		ret = driveHasVolumesFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var driveIsMediaCheckAutomaticFunction *gi.Function
var driveIsMediaCheckAutomaticFunction_Once sync.Once

func driveIsMediaCheckAutomaticFunction_Set() error {
	var err error
	driveIsMediaCheckAutomaticFunction_Once.Do(func() {
		err = driveInterface_Set()
		if err != nil {
			return
		}
		driveIsMediaCheckAutomaticFunction, err = driveInterface.InvokerNew("is_media_check_automatic")
	})
	return err
}

// IsMediaCheckAutomatic is a representation of the C type g_drive_is_media_check_automatic.
func (recv *Drive) IsMediaCheckAutomatic() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := driveIsMediaCheckAutomaticFunction_Set()
	if err == nil {
		ret = driveIsMediaCheckAutomaticFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var driveIsMediaRemovableFunction *gi.Function
var driveIsMediaRemovableFunction_Once sync.Once

func driveIsMediaRemovableFunction_Set() error {
	var err error
	driveIsMediaRemovableFunction_Once.Do(func() {
		err = driveInterface_Set()
		if err != nil {
			return
		}
		driveIsMediaRemovableFunction, err = driveInterface.InvokerNew("is_media_removable")
	})
	return err
}

// IsMediaRemovable is a representation of the C type g_drive_is_media_removable.
func (recv *Drive) IsMediaRemovable() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := driveIsMediaRemovableFunction_Set()
	if err == nil {
		ret = driveIsMediaRemovableFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var driveIsRemovableFunction *gi.Function
var driveIsRemovableFunction_Once sync.Once

func driveIsRemovableFunction_Set() error {
	var err error
	driveIsRemovableFunction_Once.Do(func() {
		err = driveInterface_Set()
		if err != nil {
			return
		}
		driveIsRemovableFunction, err = driveInterface.InvokerNew("is_removable")
	})
	return err
}

// IsRemovable is a representation of the C type g_drive_is_removable.
func (recv *Drive) IsRemovable() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := driveIsRemovableFunction_Set()
	if err == nil {
		ret = driveIsRemovableFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'g_drive_poll_for_media' : parameter 'callback' of type 'AsyncReadyCallback' not supported

// UNSUPPORTED : C value 'g_drive_poll_for_media_finish' : parameter 'result' of type 'AsyncResult' not supported

// UNSUPPORTED : C value 'g_drive_start' : parameter 'flags' of type 'DriveStartFlags' not supported

// UNSUPPORTED : C value 'g_drive_start_finish' : parameter 'result' of type 'AsyncResult' not supported

// UNSUPPORTED : C value 'g_drive_stop' : parameter 'flags' of type 'MountUnmountFlags' not supported

// UNSUPPORTED : C value 'g_drive_stop_finish' : parameter 'result' of type 'AsyncResult' not supported

/*
ConnectChanged connects a callback to the 'changed' signal of the Drive.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *Drive) ConnectChanged(handler func(instance *Drive)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {}

	return callback.ConnectSignal(recv.Native(), "changed", marshal)
}

/*
ConnectDisconnected connects a callback to the 'disconnected' signal of the Drive.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *Drive) ConnectDisconnected(handler func(instance *Drive)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {}

	return callback.ConnectSignal(recv.Native(), "disconnected", marshal)
}

/*
ConnectEjectButton connects a callback to the 'eject-button' signal of the Drive.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *Drive) ConnectEjectButton(handler func(instance *Drive)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {}

	return callback.ConnectSignal(recv.Native(), "eject-button", marshal)
}

/*
ConnectStopButton connects a callback to the 'stop-button' signal of the Drive.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *Drive) ConnectStopButton(handler func(instance *Drive)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {}

	return callback.ConnectSignal(recv.Native(), "stop-button", marshal)
}

/*
Disconnect disconnects a callback previously registered with a Connect...() method.

The connectionID should be a value returned from a call to a Connect...() method.
*/
func (recv *Drive) DisconnectSignal(connectionID int) {
	callback.DisconnectSignal(connectionID)
}

var dtlsClientConnectionInterface *gi.Interface
var dtlsClientConnectionInterface_Once sync.Once

func dtlsClientConnectionInterface_Set() error {
	var err error
	dtlsClientConnectionInterface_Once.Do(func() {
		dtlsClientConnectionInterface, err = gi.InterfaceNew("Gio", "DtlsClientConnection")
	})
	return err
}

type DtlsClientConnection struct {
	native unsafe.Pointer
}

func DtlsClientConnectionNewFromNative(native unsafe.Pointer) *DtlsClientConnection {
	instance := &DtlsClientConnection{native: native}

	return instance
}

/*
CastToDtlsClientConnection down casts any arbitrary Object to DtlsClientConnection.
Exercise care, as this is a potentially dangerous function
if the Object is not a DtlsClientConnection.
*/
func CastToDtlsClientConnection(object *gobject.Object) *DtlsClientConnection {
	return DtlsClientConnectionNewFromNative(object.Native())
}

// Equals compares this DtlsClientConnection with another DtlsClientConnection, and returns true if they represent the same Object.
func (recv *DtlsClientConnection) Equals(other *DtlsClientConnection) bool {
	return other.Native() == recv.Native()
}

func (recv *DtlsClientConnection) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'g_dtls_client_connection_get_accepted_cas' : return type 'GLib.List' not supported

// UNSUPPORTED : C value 'g_dtls_client_connection_get_server_identity' : return type 'SocketConnectable' not supported

// UNSUPPORTED : C value 'g_dtls_client_connection_get_validation_flags' : return type 'TlsCertificateFlags' not supported

// UNSUPPORTED : C value 'g_dtls_client_connection_set_server_identity' : parameter 'identity' of type 'SocketConnectable' not supported

// UNSUPPORTED : C value 'g_dtls_client_connection_set_validation_flags' : parameter 'flags' of type 'TlsCertificateFlags' not supported

var dtlsConnectionInterface *gi.Interface
var dtlsConnectionInterface_Once sync.Once

func dtlsConnectionInterface_Set() error {
	var err error
	dtlsConnectionInterface_Once.Do(func() {
		dtlsConnectionInterface, err = gi.InterfaceNew("Gio", "DtlsConnection")
	})
	return err
}

type DtlsConnection struct {
	native unsafe.Pointer
}

func DtlsConnectionNewFromNative(native unsafe.Pointer) *DtlsConnection {
	instance := &DtlsConnection{native: native}

	return instance
}

/*
CastToDtlsConnection down casts any arbitrary Object to DtlsConnection.
Exercise care, as this is a potentially dangerous function
if the Object is not a DtlsConnection.
*/
func CastToDtlsConnection(object *gobject.Object) *DtlsConnection {
	return DtlsConnectionNewFromNative(object.Native())
}

// Equals compares this DtlsConnection with another DtlsConnection, and returns true if they represent the same Object.
func (recv *DtlsConnection) Equals(other *DtlsConnection) bool {
	return other.Native() == recv.Native()
}

func (recv *DtlsConnection) Native() unsafe.Pointer {
	return recv.native
}

var dtlsConnectionCloseFunction *gi.Function
var dtlsConnectionCloseFunction_Once sync.Once

func dtlsConnectionCloseFunction_Set() error {
	var err error
	dtlsConnectionCloseFunction_Once.Do(func() {
		err = dtlsConnectionInterface_Set()
		if err != nil {
			return
		}
		dtlsConnectionCloseFunction, err = dtlsConnectionInterface.InvokerNew("close")
	})
	return err
}

// Close is a representation of the C type g_dtls_connection_close.
func (recv *DtlsConnection) Close(cancellable *Cancellable) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(cancellable.Native())

	var ret gi.Argument

	err := dtlsConnectionCloseFunction_Set()
	if err == nil {
		ret = dtlsConnectionCloseFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'g_dtls_connection_close_async' : parameter 'callback' of type 'AsyncReadyCallback' not supported

// UNSUPPORTED : C value 'g_dtls_connection_close_finish' : parameter 'result' of type 'AsyncResult' not supported

// UNSUPPORTED : C value 'g_dtls_connection_emit_accept_certificate' : parameter 'errors' of type 'TlsCertificateFlags' not supported

var dtlsConnectionGetCertificateFunction *gi.Function
var dtlsConnectionGetCertificateFunction_Once sync.Once

func dtlsConnectionGetCertificateFunction_Set() error {
	var err error
	dtlsConnectionGetCertificateFunction_Once.Do(func() {
		err = dtlsConnectionInterface_Set()
		if err != nil {
			return
		}
		dtlsConnectionGetCertificateFunction, err = dtlsConnectionInterface.InvokerNew("get_certificate")
	})
	return err
}

// GetCertificate is a representation of the C type g_dtls_connection_get_certificate.
func (recv *DtlsConnection) GetCertificate() *TlsCertificate {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := dtlsConnectionGetCertificateFunction_Set()
	if err == nil {
		ret = dtlsConnectionGetCertificateFunction.Invoke(inArgs[:], nil)
	}

	retGo := TlsCertificateNewFromNative(ret.Pointer())

	return retGo
}

var dtlsConnectionGetDatabaseFunction *gi.Function
var dtlsConnectionGetDatabaseFunction_Once sync.Once

func dtlsConnectionGetDatabaseFunction_Set() error {
	var err error
	dtlsConnectionGetDatabaseFunction_Once.Do(func() {
		err = dtlsConnectionInterface_Set()
		if err != nil {
			return
		}
		dtlsConnectionGetDatabaseFunction, err = dtlsConnectionInterface.InvokerNew("get_database")
	})
	return err
}

// GetDatabase is a representation of the C type g_dtls_connection_get_database.
func (recv *DtlsConnection) GetDatabase() *TlsDatabase {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := dtlsConnectionGetDatabaseFunction_Set()
	if err == nil {
		ret = dtlsConnectionGetDatabaseFunction.Invoke(inArgs[:], nil)
	}

	retGo := TlsDatabaseNewFromNative(ret.Pointer())

	return retGo
}

var dtlsConnectionGetInteractionFunction *gi.Function
var dtlsConnectionGetInteractionFunction_Once sync.Once

func dtlsConnectionGetInteractionFunction_Set() error {
	var err error
	dtlsConnectionGetInteractionFunction_Once.Do(func() {
		err = dtlsConnectionInterface_Set()
		if err != nil {
			return
		}
		dtlsConnectionGetInteractionFunction, err = dtlsConnectionInterface.InvokerNew("get_interaction")
	})
	return err
}

// GetInteraction is a representation of the C type g_dtls_connection_get_interaction.
func (recv *DtlsConnection) GetInteraction() *TlsInteraction {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := dtlsConnectionGetInteractionFunction_Set()
	if err == nil {
		ret = dtlsConnectionGetInteractionFunction.Invoke(inArgs[:], nil)
	}

	retGo := TlsInteractionNewFromNative(ret.Pointer())

	return retGo
}

var dtlsConnectionGetNegotiatedProtocolFunction *gi.Function
var dtlsConnectionGetNegotiatedProtocolFunction_Once sync.Once

func dtlsConnectionGetNegotiatedProtocolFunction_Set() error {
	var err error
	dtlsConnectionGetNegotiatedProtocolFunction_Once.Do(func() {
		err = dtlsConnectionInterface_Set()
		if err != nil {
			return
		}
		dtlsConnectionGetNegotiatedProtocolFunction, err = dtlsConnectionInterface.InvokerNew("get_negotiated_protocol")
	})
	return err
}

// GetNegotiatedProtocol is a representation of the C type g_dtls_connection_get_negotiated_protocol.
func (recv *DtlsConnection) GetNegotiatedProtocol() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := dtlsConnectionGetNegotiatedProtocolFunction_Set()
	if err == nil {
		ret = dtlsConnectionGetNegotiatedProtocolFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var dtlsConnectionGetPeerCertificateFunction *gi.Function
var dtlsConnectionGetPeerCertificateFunction_Once sync.Once

func dtlsConnectionGetPeerCertificateFunction_Set() error {
	var err error
	dtlsConnectionGetPeerCertificateFunction_Once.Do(func() {
		err = dtlsConnectionInterface_Set()
		if err != nil {
			return
		}
		dtlsConnectionGetPeerCertificateFunction, err = dtlsConnectionInterface.InvokerNew("get_peer_certificate")
	})
	return err
}

// GetPeerCertificate is a representation of the C type g_dtls_connection_get_peer_certificate.
func (recv *DtlsConnection) GetPeerCertificate() *TlsCertificate {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := dtlsConnectionGetPeerCertificateFunction_Set()
	if err == nil {
		ret = dtlsConnectionGetPeerCertificateFunction.Invoke(inArgs[:], nil)
	}

	retGo := TlsCertificateNewFromNative(ret.Pointer())

	return retGo
}

// UNSUPPORTED : C value 'g_dtls_connection_get_peer_certificate_errors' : return type 'TlsCertificateFlags' not supported

var dtlsConnectionGetRehandshakeModeFunction *gi.Function
var dtlsConnectionGetRehandshakeModeFunction_Once sync.Once

func dtlsConnectionGetRehandshakeModeFunction_Set() error {
	var err error
	dtlsConnectionGetRehandshakeModeFunction_Once.Do(func() {
		err = dtlsConnectionInterface_Set()
		if err != nil {
			return
		}
		dtlsConnectionGetRehandshakeModeFunction, err = dtlsConnectionInterface.InvokerNew("get_rehandshake_mode")
	})
	return err
}

// GetRehandshakeMode is a representation of the C type g_dtls_connection_get_rehandshake_mode.
func (recv *DtlsConnection) GetRehandshakeMode() TlsRehandshakeMode {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := dtlsConnectionGetRehandshakeModeFunction_Set()
	if err == nil {
		ret = dtlsConnectionGetRehandshakeModeFunction.Invoke(inArgs[:], nil)
	}

	retGo := TlsRehandshakeMode(ret.Int32())

	return retGo
}

var dtlsConnectionGetRequireCloseNotifyFunction *gi.Function
var dtlsConnectionGetRequireCloseNotifyFunction_Once sync.Once

func dtlsConnectionGetRequireCloseNotifyFunction_Set() error {
	var err error
	dtlsConnectionGetRequireCloseNotifyFunction_Once.Do(func() {
		err = dtlsConnectionInterface_Set()
		if err != nil {
			return
		}
		dtlsConnectionGetRequireCloseNotifyFunction, err = dtlsConnectionInterface.InvokerNew("get_require_close_notify")
	})
	return err
}

// GetRequireCloseNotify is a representation of the C type g_dtls_connection_get_require_close_notify.
func (recv *DtlsConnection) GetRequireCloseNotify() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := dtlsConnectionGetRequireCloseNotifyFunction_Set()
	if err == nil {
		ret = dtlsConnectionGetRequireCloseNotifyFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var dtlsConnectionHandshakeFunction *gi.Function
var dtlsConnectionHandshakeFunction_Once sync.Once

func dtlsConnectionHandshakeFunction_Set() error {
	var err error
	dtlsConnectionHandshakeFunction_Once.Do(func() {
		err = dtlsConnectionInterface_Set()
		if err != nil {
			return
		}
		dtlsConnectionHandshakeFunction, err = dtlsConnectionInterface.InvokerNew("handshake")
	})
	return err
}

// Handshake is a representation of the C type g_dtls_connection_handshake.
func (recv *DtlsConnection) Handshake(cancellable *Cancellable) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(cancellable.Native())

	var ret gi.Argument

	err := dtlsConnectionHandshakeFunction_Set()
	if err == nil {
		ret = dtlsConnectionHandshakeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'g_dtls_connection_handshake_async' : parameter 'callback' of type 'AsyncReadyCallback' not supported

// UNSUPPORTED : C value 'g_dtls_connection_handshake_finish' : parameter 'result' of type 'AsyncResult' not supported

// UNSUPPORTED : C value 'g_dtls_connection_set_advertised_protocols' : parameter 'protocols' of type 'nil' not supported

var dtlsConnectionSetCertificateFunction *gi.Function
var dtlsConnectionSetCertificateFunction_Once sync.Once

func dtlsConnectionSetCertificateFunction_Set() error {
	var err error
	dtlsConnectionSetCertificateFunction_Once.Do(func() {
		err = dtlsConnectionInterface_Set()
		if err != nil {
			return
		}
		dtlsConnectionSetCertificateFunction, err = dtlsConnectionInterface.InvokerNew("set_certificate")
	})
	return err
}

// SetCertificate is a representation of the C type g_dtls_connection_set_certificate.
func (recv *DtlsConnection) SetCertificate(certificate *TlsCertificate) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(certificate.Native())

	err := dtlsConnectionSetCertificateFunction_Set()
	if err == nil {
		dtlsConnectionSetCertificateFunction.Invoke(inArgs[:], nil)
	}

	return
}

var dtlsConnectionSetDatabaseFunction *gi.Function
var dtlsConnectionSetDatabaseFunction_Once sync.Once

func dtlsConnectionSetDatabaseFunction_Set() error {
	var err error
	dtlsConnectionSetDatabaseFunction_Once.Do(func() {
		err = dtlsConnectionInterface_Set()
		if err != nil {
			return
		}
		dtlsConnectionSetDatabaseFunction, err = dtlsConnectionInterface.InvokerNew("set_database")
	})
	return err
}

// SetDatabase is a representation of the C type g_dtls_connection_set_database.
func (recv *DtlsConnection) SetDatabase(database *TlsDatabase) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(database.Native())

	err := dtlsConnectionSetDatabaseFunction_Set()
	if err == nil {
		dtlsConnectionSetDatabaseFunction.Invoke(inArgs[:], nil)
	}

	return
}

var dtlsConnectionSetInteractionFunction *gi.Function
var dtlsConnectionSetInteractionFunction_Once sync.Once

func dtlsConnectionSetInteractionFunction_Set() error {
	var err error
	dtlsConnectionSetInteractionFunction_Once.Do(func() {
		err = dtlsConnectionInterface_Set()
		if err != nil {
			return
		}
		dtlsConnectionSetInteractionFunction, err = dtlsConnectionInterface.InvokerNew("set_interaction")
	})
	return err
}

// SetInteraction is a representation of the C type g_dtls_connection_set_interaction.
func (recv *DtlsConnection) SetInteraction(interaction *TlsInteraction) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(interaction.Native())

	err := dtlsConnectionSetInteractionFunction_Set()
	if err == nil {
		dtlsConnectionSetInteractionFunction.Invoke(inArgs[:], nil)
	}

	return
}

var dtlsConnectionSetRehandshakeModeFunction *gi.Function
var dtlsConnectionSetRehandshakeModeFunction_Once sync.Once

func dtlsConnectionSetRehandshakeModeFunction_Set() error {
	var err error
	dtlsConnectionSetRehandshakeModeFunction_Once.Do(func() {
		err = dtlsConnectionInterface_Set()
		if err != nil {
			return
		}
		dtlsConnectionSetRehandshakeModeFunction, err = dtlsConnectionInterface.InvokerNew("set_rehandshake_mode")
	})
	return err
}

// SetRehandshakeMode is a representation of the C type g_dtls_connection_set_rehandshake_mode.
func (recv *DtlsConnection) SetRehandshakeMode(mode TlsRehandshakeMode) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(int32(mode))

	err := dtlsConnectionSetRehandshakeModeFunction_Set()
	if err == nil {
		dtlsConnectionSetRehandshakeModeFunction.Invoke(inArgs[:], nil)
	}

	return
}

var dtlsConnectionSetRequireCloseNotifyFunction *gi.Function
var dtlsConnectionSetRequireCloseNotifyFunction_Once sync.Once

func dtlsConnectionSetRequireCloseNotifyFunction_Set() error {
	var err error
	dtlsConnectionSetRequireCloseNotifyFunction_Once.Do(func() {
		err = dtlsConnectionInterface_Set()
		if err != nil {
			return
		}
		dtlsConnectionSetRequireCloseNotifyFunction, err = dtlsConnectionInterface.InvokerNew("set_require_close_notify")
	})
	return err
}

// SetRequireCloseNotify is a representation of the C type g_dtls_connection_set_require_close_notify.
func (recv *DtlsConnection) SetRequireCloseNotify(requireCloseNotify bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetBoolean(requireCloseNotify)

	err := dtlsConnectionSetRequireCloseNotifyFunction_Set()
	if err == nil {
		dtlsConnectionSetRequireCloseNotifyFunction.Invoke(inArgs[:], nil)
	}

	return
}

var dtlsConnectionShutdownFunction *gi.Function
var dtlsConnectionShutdownFunction_Once sync.Once

func dtlsConnectionShutdownFunction_Set() error {
	var err error
	dtlsConnectionShutdownFunction_Once.Do(func() {
		err = dtlsConnectionInterface_Set()
		if err != nil {
			return
		}
		dtlsConnectionShutdownFunction, err = dtlsConnectionInterface.InvokerNew("shutdown")
	})
	return err
}

// Shutdown is a representation of the C type g_dtls_connection_shutdown.
func (recv *DtlsConnection) Shutdown(shutdownRead bool, shutdownWrite bool, cancellable *Cancellable) bool {
	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetBoolean(shutdownRead)
	inArgs[2].SetBoolean(shutdownWrite)
	inArgs[3].SetPointer(cancellable.Native())

	var ret gi.Argument

	err := dtlsConnectionShutdownFunction_Set()
	if err == nil {
		ret = dtlsConnectionShutdownFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'g_dtls_connection_shutdown_async' : parameter 'callback' of type 'AsyncReadyCallback' not supported

// UNSUPPORTED : C value 'g_dtls_connection_shutdown_finish' : parameter 'result' of type 'AsyncResult' not supported

// UNSUPPORTED : C value 'accept-certificate' : parameter 'errors' of type 'TlsCertificateFlags' not supported

/*
Disconnect disconnects a callback previously registered with a Connect...() method.

The connectionID should be a value returned from a call to a Connect...() method.
*/
func (recv *DtlsConnection) DisconnectSignal(connectionID int) {
	callback.DisconnectSignal(connectionID)
}

var dtlsServerConnectionInterface *gi.Interface
var dtlsServerConnectionInterface_Once sync.Once

func dtlsServerConnectionInterface_Set() error {
	var err error
	dtlsServerConnectionInterface_Once.Do(func() {
		dtlsServerConnectionInterface, err = gi.InterfaceNew("Gio", "DtlsServerConnection")
	})
	return err
}

type DtlsServerConnection struct {
	native unsafe.Pointer
}

func DtlsServerConnectionNewFromNative(native unsafe.Pointer) *DtlsServerConnection {
	instance := &DtlsServerConnection{native: native}

	return instance
}

/*
CastToDtlsServerConnection down casts any arbitrary Object to DtlsServerConnection.
Exercise care, as this is a potentially dangerous function
if the Object is not a DtlsServerConnection.
*/
func CastToDtlsServerConnection(object *gobject.Object) *DtlsServerConnection {
	return DtlsServerConnectionNewFromNative(object.Native())
}

// Equals compares this DtlsServerConnection with another DtlsServerConnection, and returns true if they represent the same Object.
func (recv *DtlsServerConnection) Equals(other *DtlsServerConnection) bool {
	return other.Native() == recv.Native()
}

func (recv *DtlsServerConnection) Native() unsafe.Pointer {
	return recv.native
}

var fileInterface *gi.Interface
var fileInterface_Once sync.Once

func fileInterface_Set() error {
	var err error
	fileInterface_Once.Do(func() {
		fileInterface, err = gi.InterfaceNew("Gio", "File")
	})
	return err
}

type File struct {
	native unsafe.Pointer
}

func FileNewFromNative(native unsafe.Pointer) *File {
	instance := &File{native: native}

	return instance
}

/*
CastToFile down casts any arbitrary Object to File.
Exercise care, as this is a potentially dangerous function
if the Object is not a File.
*/
func CastToFile(object *gobject.Object) *File {
	return FileNewFromNative(object.Native())
}

// Equals compares this File with another File, and returns true if they represent the same Object.
func (recv *File) Equals(other *File) bool {
	return other.Native() == recv.Native()
}

func (recv *File) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'g_file_append_to' : parameter 'flags' of type 'FileCreateFlags' not supported

// UNSUPPORTED : C value 'g_file_append_to_async' : parameter 'flags' of type 'FileCreateFlags' not supported

// UNSUPPORTED : C value 'g_file_append_to_finish' : parameter 'res' of type 'AsyncResult' not supported

// UNSUPPORTED : C value 'g_file_copy' : parameter 'destination' of type 'File' not supported

// UNSUPPORTED : C value 'g_file_copy_async' : parameter 'destination' of type 'File' not supported

// UNSUPPORTED : C value 'g_file_copy_attributes' : parameter 'destination' of type 'File' not supported

// UNSUPPORTED : C value 'g_file_copy_finish' : parameter 'res' of type 'AsyncResult' not supported

// UNSUPPORTED : C value 'g_file_create' : parameter 'flags' of type 'FileCreateFlags' not supported

// UNSUPPORTED : C value 'g_file_create_async' : parameter 'flags' of type 'FileCreateFlags' not supported

// UNSUPPORTED : C value 'g_file_create_finish' : parameter 'res' of type 'AsyncResult' not supported

// UNSUPPORTED : C value 'g_file_create_readwrite' : parameter 'flags' of type 'FileCreateFlags' not supported

// UNSUPPORTED : C value 'g_file_create_readwrite_async' : parameter 'flags' of type 'FileCreateFlags' not supported

// UNSUPPORTED : C value 'g_file_create_readwrite_finish' : parameter 'res' of type 'AsyncResult' not supported

var fileDeleteFunction *gi.Function
var fileDeleteFunction_Once sync.Once

func fileDeleteFunction_Set() error {
	var err error
	fileDeleteFunction_Once.Do(func() {
		err = fileInterface_Set()
		if err != nil {
			return
		}
		fileDeleteFunction, err = fileInterface.InvokerNew("delete")
	})
	return err
}

// Delete is a representation of the C type g_file_delete.
func (recv *File) Delete(cancellable *Cancellable) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(cancellable.Native())

	var ret gi.Argument

	err := fileDeleteFunction_Set()
	if err == nil {
		ret = fileDeleteFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'g_file_delete_async' : parameter 'callback' of type 'AsyncReadyCallback' not supported

// UNSUPPORTED : C value 'g_file_delete_finish' : parameter 'result' of type 'AsyncResult' not supported

// UNSUPPORTED : C value 'g_file_dup' : return type 'File' not supported

// UNSUPPORTED : C value 'g_file_eject_mountable' : parameter 'flags' of type 'MountUnmountFlags' not supported

// UNSUPPORTED : C value 'g_file_eject_mountable_finish' : parameter 'result' of type 'AsyncResult' not supported

// UNSUPPORTED : C value 'g_file_eject_mountable_with_operation' : parameter 'flags' of type 'MountUnmountFlags' not supported

// UNSUPPORTED : C value 'g_file_eject_mountable_with_operation_finish' : parameter 'result' of type 'AsyncResult' not supported

// UNSUPPORTED : C value 'g_file_enumerate_children' : parameter 'flags' of type 'FileQueryInfoFlags' not supported

// UNSUPPORTED : C value 'g_file_enumerate_children_async' : parameter 'flags' of type 'FileQueryInfoFlags' not supported

// UNSUPPORTED : C value 'g_file_enumerate_children_finish' : parameter 'res' of type 'AsyncResult' not supported

// UNSUPPORTED : C value 'g_file_equal' : parameter 'file2' of type 'File' not supported

// UNSUPPORTED : C value 'g_file_find_enclosing_mount' : return type 'Mount' not supported

// UNSUPPORTED : C value 'g_file_find_enclosing_mount_async' : parameter 'callback' of type 'AsyncReadyCallback' not supported

// UNSUPPORTED : C value 'g_file_find_enclosing_mount_finish' : parameter 'res' of type 'AsyncResult' not supported

var fileGetBasenameFunction *gi.Function
var fileGetBasenameFunction_Once sync.Once

func fileGetBasenameFunction_Set() error {
	var err error
	fileGetBasenameFunction_Once.Do(func() {
		err = fileInterface_Set()
		if err != nil {
			return
		}
		fileGetBasenameFunction, err = fileInterface.InvokerNew("get_basename")
	})
	return err
}

// GetBasename is a representation of the C type g_file_get_basename.
func (recv *File) GetBasename() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := fileGetBasenameFunction_Set()
	if err == nil {
		ret = fileGetBasenameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

// UNSUPPORTED : C value 'g_file_get_child' : return type 'File' not supported

// UNSUPPORTED : C value 'g_file_get_child_for_display_name' : return type 'File' not supported

// UNSUPPORTED : C value 'g_file_get_parent' : return type 'File' not supported

var fileGetParseNameFunction *gi.Function
var fileGetParseNameFunction_Once sync.Once

func fileGetParseNameFunction_Set() error {
	var err error
	fileGetParseNameFunction_Once.Do(func() {
		err = fileInterface_Set()
		if err != nil {
			return
		}
		fileGetParseNameFunction, err = fileInterface.InvokerNew("get_parse_name")
	})
	return err
}

// GetParseName is a representation of the C type g_file_get_parse_name.
func (recv *File) GetParseName() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := fileGetParseNameFunction_Set()
	if err == nil {
		ret = fileGetParseNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

var fileGetPathFunction *gi.Function
var fileGetPathFunction_Once sync.Once

func fileGetPathFunction_Set() error {
	var err error
	fileGetPathFunction_Once.Do(func() {
		err = fileInterface_Set()
		if err != nil {
			return
		}
		fileGetPathFunction, err = fileInterface.InvokerNew("get_path")
	})
	return err
}

// GetPath is a representation of the C type g_file_get_path.
func (recv *File) GetPath() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := fileGetPathFunction_Set()
	if err == nil {
		ret = fileGetPathFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

// UNSUPPORTED : C value 'g_file_get_relative_path' : parameter 'descendant' of type 'File' not supported

var fileGetUriFunction *gi.Function
var fileGetUriFunction_Once sync.Once

func fileGetUriFunction_Set() error {
	var err error
	fileGetUriFunction_Once.Do(func() {
		err = fileInterface_Set()
		if err != nil {
			return
		}
		fileGetUriFunction, err = fileInterface.InvokerNew("get_uri")
	})
	return err
}

// GetUri is a representation of the C type g_file_get_uri.
func (recv *File) GetUri() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := fileGetUriFunction_Set()
	if err == nil {
		ret = fileGetUriFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

var fileGetUriSchemeFunction *gi.Function
var fileGetUriSchemeFunction_Once sync.Once

func fileGetUriSchemeFunction_Set() error {
	var err error
	fileGetUriSchemeFunction_Once.Do(func() {
		err = fileInterface_Set()
		if err != nil {
			return
		}
		fileGetUriSchemeFunction, err = fileInterface.InvokerNew("get_uri_scheme")
	})
	return err
}

// GetUriScheme is a representation of the C type g_file_get_uri_scheme.
func (recv *File) GetUriScheme() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := fileGetUriSchemeFunction_Set()
	if err == nil {
		ret = fileGetUriSchemeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

// UNSUPPORTED : C value 'g_file_has_parent' : parameter 'parent' of type 'File' not supported

// UNSUPPORTED : C value 'g_file_has_prefix' : parameter 'prefix' of type 'File' not supported

var fileHasUriSchemeFunction *gi.Function
var fileHasUriSchemeFunction_Once sync.Once

func fileHasUriSchemeFunction_Set() error {
	var err error
	fileHasUriSchemeFunction_Once.Do(func() {
		err = fileInterface_Set()
		if err != nil {
			return
		}
		fileHasUriSchemeFunction, err = fileInterface.InvokerNew("has_uri_scheme")
	})
	return err
}

// HasUriScheme is a representation of the C type g_file_has_uri_scheme.
func (recv *File) HasUriScheme(uriScheme string) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(uriScheme)

	var ret gi.Argument

	err := fileHasUriSchemeFunction_Set()
	if err == nil {
		ret = fileHasUriSchemeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var fileHashFunction *gi.Function
var fileHashFunction_Once sync.Once

func fileHashFunction_Set() error {
	var err error
	fileHashFunction_Once.Do(func() {
		err = fileInterface_Set()
		if err != nil {
			return
		}
		fileHashFunction, err = fileInterface.InvokerNew("hash")
	})
	return err
}

// Hash is a representation of the C type g_file_hash.
func (recv *File) Hash() uint32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := fileHashFunction_Set()
	if err == nil {
		ret = fileHashFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo
}

var fileIsNativeFunction *gi.Function
var fileIsNativeFunction_Once sync.Once

func fileIsNativeFunction_Set() error {
	var err error
	fileIsNativeFunction_Once.Do(func() {
		err = fileInterface_Set()
		if err != nil {
			return
		}
		fileIsNativeFunction, err = fileInterface.InvokerNew("is_native")
	})
	return err
}

// IsNative is a representation of the C type g_file_is_native.
func (recv *File) IsNative() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := fileIsNativeFunction_Set()
	if err == nil {
		ret = fileIsNativeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'g_file_load_bytes' : return type 'GLib.Bytes' not supported

// UNSUPPORTED : C value 'g_file_load_bytes_async' : parameter 'callback' of type 'AsyncReadyCallback' not supported

// UNSUPPORTED : C value 'g_file_load_bytes_finish' : parameter 'result' of type 'AsyncResult' not supported

// UNSUPPORTED : C value 'g_file_load_contents' : parameter 'contents' of type 'nil' not supported

// UNSUPPORTED : C value 'g_file_load_contents_async' : parameter 'callback' of type 'AsyncReadyCallback' not supported

// UNSUPPORTED : C value 'g_file_load_contents_finish' : parameter 'res' of type 'AsyncResult' not supported

// UNSUPPORTED : C value 'g_file_load_partial_contents_async' : parameter 'read_more_callback' of type 'FileReadMoreCallback' not supported

// UNSUPPORTED : C value 'g_file_load_partial_contents_finish' : parameter 'res' of type 'AsyncResult' not supported

var fileMakeDirectoryFunction *gi.Function
var fileMakeDirectoryFunction_Once sync.Once

func fileMakeDirectoryFunction_Set() error {
	var err error
	fileMakeDirectoryFunction_Once.Do(func() {
		err = fileInterface_Set()
		if err != nil {
			return
		}
		fileMakeDirectoryFunction, err = fileInterface.InvokerNew("make_directory")
	})
	return err
}

// MakeDirectory is a representation of the C type g_file_make_directory.
func (recv *File) MakeDirectory(cancellable *Cancellable) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(cancellable.Native())

	var ret gi.Argument

	err := fileMakeDirectoryFunction_Set()
	if err == nil {
		ret = fileMakeDirectoryFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'g_file_make_directory_async' : parameter 'callback' of type 'AsyncReadyCallback' not supported

// UNSUPPORTED : C value 'g_file_make_directory_finish' : parameter 'result' of type 'AsyncResult' not supported

var fileMakeDirectoryWithParentsFunction *gi.Function
var fileMakeDirectoryWithParentsFunction_Once sync.Once

func fileMakeDirectoryWithParentsFunction_Set() error {
	var err error
	fileMakeDirectoryWithParentsFunction_Once.Do(func() {
		err = fileInterface_Set()
		if err != nil {
			return
		}
		fileMakeDirectoryWithParentsFunction, err = fileInterface.InvokerNew("make_directory_with_parents")
	})
	return err
}

// MakeDirectoryWithParents is a representation of the C type g_file_make_directory_with_parents.
func (recv *File) MakeDirectoryWithParents(cancellable *Cancellable) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(cancellable.Native())

	var ret gi.Argument

	err := fileMakeDirectoryWithParentsFunction_Set()
	if err == nil {
		ret = fileMakeDirectoryWithParentsFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var fileMakeSymbolicLinkFunction *gi.Function
var fileMakeSymbolicLinkFunction_Once sync.Once

func fileMakeSymbolicLinkFunction_Set() error {
	var err error
	fileMakeSymbolicLinkFunction_Once.Do(func() {
		err = fileInterface_Set()
		if err != nil {
			return
		}
		fileMakeSymbolicLinkFunction, err = fileInterface.InvokerNew("make_symbolic_link")
	})
	return err
}

// MakeSymbolicLink is a representation of the C type g_file_make_symbolic_link.
func (recv *File) MakeSymbolicLink(symlinkValue string, cancellable *Cancellable) bool {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(symlinkValue)
	inArgs[2].SetPointer(cancellable.Native())

	var ret gi.Argument

	err := fileMakeSymbolicLinkFunction_Set()
	if err == nil {
		ret = fileMakeSymbolicLinkFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'g_file_measure_disk_usage' : parameter 'flags' of type 'FileMeasureFlags' not supported

// UNSUPPORTED : C value 'g_file_measure_disk_usage_async' : parameter 'flags' of type 'FileMeasureFlags' not supported

// UNSUPPORTED : C value 'g_file_measure_disk_usage_finish' : parameter 'result' of type 'AsyncResult' not supported

// UNSUPPORTED : C value 'g_file_monitor' : parameter 'flags' of type 'FileMonitorFlags' not supported

// UNSUPPORTED : C value 'g_file_monitor_directory' : parameter 'flags' of type 'FileMonitorFlags' not supported

// UNSUPPORTED : C value 'g_file_monitor_file' : parameter 'flags' of type 'FileMonitorFlags' not supported

// UNSUPPORTED : C value 'g_file_mount_enclosing_volume' : parameter 'flags' of type 'MountMountFlags' not supported

// UNSUPPORTED : C value 'g_file_mount_enclosing_volume_finish' : parameter 'result' of type 'AsyncResult' not supported

// UNSUPPORTED : C value 'g_file_mount_mountable' : parameter 'flags' of type 'MountMountFlags' not supported

// UNSUPPORTED : C value 'g_file_mount_mountable_finish' : parameter 'result' of type 'AsyncResult' not supported

// UNSUPPORTED : C value 'g_file_move' : parameter 'destination' of type 'File' not supported

var fileOpenReadwriteFunction *gi.Function
var fileOpenReadwriteFunction_Once sync.Once

func fileOpenReadwriteFunction_Set() error {
	var err error
	fileOpenReadwriteFunction_Once.Do(func() {
		err = fileInterface_Set()
		if err != nil {
			return
		}
		fileOpenReadwriteFunction, err = fileInterface.InvokerNew("open_readwrite")
	})
	return err
}

// OpenReadwrite is a representation of the C type g_file_open_readwrite.
func (recv *File) OpenReadwrite(cancellable *Cancellable) *FileIOStream {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(cancellable.Native())

	var ret gi.Argument

	err := fileOpenReadwriteFunction_Set()
	if err == nil {
		ret = fileOpenReadwriteFunction.Invoke(inArgs[:], nil)
	}

	retGo := FileIOStreamNewFromNative(ret.Pointer())

	return retGo
}

// UNSUPPORTED : C value 'g_file_open_readwrite_async' : parameter 'callback' of type 'AsyncReadyCallback' not supported

// UNSUPPORTED : C value 'g_file_open_readwrite_finish' : parameter 'res' of type 'AsyncResult' not supported

var filePeekPathFunction *gi.Function
var filePeekPathFunction_Once sync.Once

func filePeekPathFunction_Set() error {
	var err error
	filePeekPathFunction_Once.Do(func() {
		err = fileInterface_Set()
		if err != nil {
			return
		}
		filePeekPathFunction, err = fileInterface.InvokerNew("peek_path")
	})
	return err
}

// PeekPath is a representation of the C type g_file_peek_path.
func (recv *File) PeekPath() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := filePeekPathFunction_Set()
	if err == nil {
		ret = filePeekPathFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'g_file_poll_mountable' : parameter 'callback' of type 'AsyncReadyCallback' not supported

// UNSUPPORTED : C value 'g_file_poll_mountable_finish' : parameter 'result' of type 'AsyncResult' not supported

// UNSUPPORTED : C value 'g_file_query_default_handler' : return type 'AppInfo' not supported

// UNSUPPORTED : C value 'g_file_query_default_handler_async' : parameter 'callback' of type 'AsyncReadyCallback' not supported

// UNSUPPORTED : C value 'g_file_query_default_handler_finish' : parameter 'result' of type 'AsyncResult' not supported

var fileQueryExistsFunction *gi.Function
var fileQueryExistsFunction_Once sync.Once

func fileQueryExistsFunction_Set() error {
	var err error
	fileQueryExistsFunction_Once.Do(func() {
		err = fileInterface_Set()
		if err != nil {
			return
		}
		fileQueryExistsFunction, err = fileInterface.InvokerNew("query_exists")
	})
	return err
}

// QueryExists is a representation of the C type g_file_query_exists.
func (recv *File) QueryExists(cancellable *Cancellable) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(cancellable.Native())

	var ret gi.Argument

	err := fileQueryExistsFunction_Set()
	if err == nil {
		ret = fileQueryExistsFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'g_file_query_file_type' : parameter 'flags' of type 'FileQueryInfoFlags' not supported

var fileQueryFilesystemInfoFunction *gi.Function
var fileQueryFilesystemInfoFunction_Once sync.Once

func fileQueryFilesystemInfoFunction_Set() error {
	var err error
	fileQueryFilesystemInfoFunction_Once.Do(func() {
		err = fileInterface_Set()
		if err != nil {
			return
		}
		fileQueryFilesystemInfoFunction, err = fileInterface.InvokerNew("query_filesystem_info")
	})
	return err
}

// QueryFilesystemInfo is a representation of the C type g_file_query_filesystem_info.
func (recv *File) QueryFilesystemInfo(attributes string, cancellable *Cancellable) *FileInfo {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(attributes)
	inArgs[2].SetPointer(cancellable.Native())

	var ret gi.Argument

	err := fileQueryFilesystemInfoFunction_Set()
	if err == nil {
		ret = fileQueryFilesystemInfoFunction.Invoke(inArgs[:], nil)
	}

	retGo := FileInfoNewFromNative(ret.Pointer())

	return retGo
}

// UNSUPPORTED : C value 'g_file_query_filesystem_info_async' : parameter 'callback' of type 'AsyncReadyCallback' not supported

// UNSUPPORTED : C value 'g_file_query_filesystem_info_finish' : parameter 'res' of type 'AsyncResult' not supported

// UNSUPPORTED : C value 'g_file_query_info' : parameter 'flags' of type 'FileQueryInfoFlags' not supported

// UNSUPPORTED : C value 'g_file_query_info_async' : parameter 'flags' of type 'FileQueryInfoFlags' not supported

// UNSUPPORTED : C value 'g_file_query_info_finish' : parameter 'res' of type 'AsyncResult' not supported

var fileQuerySettableAttributesFunction *gi.Function
var fileQuerySettableAttributesFunction_Once sync.Once

func fileQuerySettableAttributesFunction_Set() error {
	var err error
	fileQuerySettableAttributesFunction_Once.Do(func() {
		err = fileInterface_Set()
		if err != nil {
			return
		}
		fileQuerySettableAttributesFunction, err = fileInterface.InvokerNew("query_settable_attributes")
	})
	return err
}

// QuerySettableAttributes is a representation of the C type g_file_query_settable_attributes.
func (recv *File) QuerySettableAttributes(cancellable *Cancellable) *FileAttributeInfoList {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(cancellable.Native())

	var ret gi.Argument

	err := fileQuerySettableAttributesFunction_Set()
	if err == nil {
		ret = fileQuerySettableAttributesFunction.Invoke(inArgs[:], nil)
	}

	retGo := FileAttributeInfoListNewFromNative(ret.Pointer())

	return retGo
}

var fileQueryWritableNamespacesFunction *gi.Function
var fileQueryWritableNamespacesFunction_Once sync.Once

func fileQueryWritableNamespacesFunction_Set() error {
	var err error
	fileQueryWritableNamespacesFunction_Once.Do(func() {
		err = fileInterface_Set()
		if err != nil {
			return
		}
		fileQueryWritableNamespacesFunction, err = fileInterface.InvokerNew("query_writable_namespaces")
	})
	return err
}

// QueryWritableNamespaces is a representation of the C type g_file_query_writable_namespaces.
func (recv *File) QueryWritableNamespaces(cancellable *Cancellable) *FileAttributeInfoList {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(cancellable.Native())

	var ret gi.Argument

	err := fileQueryWritableNamespacesFunction_Set()
	if err == nil {
		ret = fileQueryWritableNamespacesFunction.Invoke(inArgs[:], nil)
	}

	retGo := FileAttributeInfoListNewFromNative(ret.Pointer())

	return retGo
}

var fileReadFunction *gi.Function
var fileReadFunction_Once sync.Once

func fileReadFunction_Set() error {
	var err error
	fileReadFunction_Once.Do(func() {
		err = fileInterface_Set()
		if err != nil {
			return
		}
		fileReadFunction, err = fileInterface.InvokerNew("read")
	})
	return err
}

// Read is a representation of the C type g_file_read.
func (recv *File) Read(cancellable *Cancellable) *FileInputStream {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(cancellable.Native())

	var ret gi.Argument

	err := fileReadFunction_Set()
	if err == nil {
		ret = fileReadFunction.Invoke(inArgs[:], nil)
	}

	retGo := FileInputStreamNewFromNative(ret.Pointer())

	return retGo
}

// UNSUPPORTED : C value 'g_file_read_async' : parameter 'callback' of type 'AsyncReadyCallback' not supported

// UNSUPPORTED : C value 'g_file_read_finish' : parameter 'res' of type 'AsyncResult' not supported

// UNSUPPORTED : C value 'g_file_replace' : parameter 'flags' of type 'FileCreateFlags' not supported

// UNSUPPORTED : C value 'g_file_replace_async' : parameter 'flags' of type 'FileCreateFlags' not supported

// UNSUPPORTED : C value 'g_file_replace_contents' : parameter 'contents' of type 'nil' not supported

// UNSUPPORTED : C value 'g_file_replace_contents_async' : parameter 'contents' of type 'nil' not supported

// UNSUPPORTED : C value 'g_file_replace_contents_bytes_async' : parameter 'contents' of type 'GLib.Bytes' not supported

// UNSUPPORTED : C value 'g_file_replace_contents_finish' : parameter 'res' of type 'AsyncResult' not supported

// UNSUPPORTED : C value 'g_file_replace_finish' : parameter 'res' of type 'AsyncResult' not supported

// UNSUPPORTED : C value 'g_file_replace_readwrite' : parameter 'flags' of type 'FileCreateFlags' not supported

// UNSUPPORTED : C value 'g_file_replace_readwrite_async' : parameter 'flags' of type 'FileCreateFlags' not supported

// UNSUPPORTED : C value 'g_file_replace_readwrite_finish' : parameter 'res' of type 'AsyncResult' not supported

// UNSUPPORTED : C value 'g_file_resolve_relative_path' : return type 'File' not supported

// UNSUPPORTED : C value 'g_file_set_attribute' : parameter 'value_p' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_file_set_attribute_byte_string' : parameter 'flags' of type 'FileQueryInfoFlags' not supported

// UNSUPPORTED : C value 'g_file_set_attribute_int32' : parameter 'flags' of type 'FileQueryInfoFlags' not supported

// UNSUPPORTED : C value 'g_file_set_attribute_int64' : parameter 'flags' of type 'FileQueryInfoFlags' not supported

// UNSUPPORTED : C value 'g_file_set_attribute_string' : parameter 'flags' of type 'FileQueryInfoFlags' not supported

// UNSUPPORTED : C value 'g_file_set_attribute_uint32' : parameter 'flags' of type 'FileQueryInfoFlags' not supported

// UNSUPPORTED : C value 'g_file_set_attribute_uint64' : parameter 'flags' of type 'FileQueryInfoFlags' not supported

// UNSUPPORTED : C value 'g_file_set_attributes_async' : parameter 'flags' of type 'FileQueryInfoFlags' not supported

// UNSUPPORTED : C value 'g_file_set_attributes_finish' : parameter 'result' of type 'AsyncResult' not supported

// UNSUPPORTED : C value 'g_file_set_attributes_from_info' : parameter 'flags' of type 'FileQueryInfoFlags' not supported

// UNSUPPORTED : C value 'g_file_set_display_name' : return type 'File' not supported

// UNSUPPORTED : C value 'g_file_set_display_name_async' : parameter 'callback' of type 'AsyncReadyCallback' not supported

// UNSUPPORTED : C value 'g_file_set_display_name_finish' : parameter 'res' of type 'AsyncResult' not supported

// UNSUPPORTED : C value 'g_file_start_mountable' : parameter 'flags' of type 'DriveStartFlags' not supported

// UNSUPPORTED : C value 'g_file_start_mountable_finish' : parameter 'result' of type 'AsyncResult' not supported

// UNSUPPORTED : C value 'g_file_stop_mountable' : parameter 'flags' of type 'MountUnmountFlags' not supported

// UNSUPPORTED : C value 'g_file_stop_mountable_finish' : parameter 'result' of type 'AsyncResult' not supported

var fileSupportsThreadContextsFunction *gi.Function
var fileSupportsThreadContextsFunction_Once sync.Once

func fileSupportsThreadContextsFunction_Set() error {
	var err error
	fileSupportsThreadContextsFunction_Once.Do(func() {
		err = fileInterface_Set()
		if err != nil {
			return
		}
		fileSupportsThreadContextsFunction, err = fileInterface.InvokerNew("supports_thread_contexts")
	})
	return err
}

// SupportsThreadContexts is a representation of the C type g_file_supports_thread_contexts.
func (recv *File) SupportsThreadContexts() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := fileSupportsThreadContextsFunction_Set()
	if err == nil {
		ret = fileSupportsThreadContextsFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var fileTrashFunction *gi.Function
var fileTrashFunction_Once sync.Once

func fileTrashFunction_Set() error {
	var err error
	fileTrashFunction_Once.Do(func() {
		err = fileInterface_Set()
		if err != nil {
			return
		}
		fileTrashFunction, err = fileInterface.InvokerNew("trash")
	})
	return err
}

// Trash is a representation of the C type g_file_trash.
func (recv *File) Trash(cancellable *Cancellable) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(cancellable.Native())

	var ret gi.Argument

	err := fileTrashFunction_Set()
	if err == nil {
		ret = fileTrashFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'g_file_trash_async' : parameter 'callback' of type 'AsyncReadyCallback' not supported

// UNSUPPORTED : C value 'g_file_trash_finish' : parameter 'result' of type 'AsyncResult' not supported

// UNSUPPORTED : C value 'g_file_unmount_mountable' : parameter 'flags' of type 'MountUnmountFlags' not supported

// UNSUPPORTED : C value 'g_file_unmount_mountable_finish' : parameter 'result' of type 'AsyncResult' not supported

// UNSUPPORTED : C value 'g_file_unmount_mountable_with_operation' : parameter 'flags' of type 'MountUnmountFlags' not supported

// UNSUPPORTED : C value 'g_file_unmount_mountable_with_operation_finish' : parameter 'result' of type 'AsyncResult' not supported

var fileDescriptorBasedInterface *gi.Interface
var fileDescriptorBasedInterface_Once sync.Once

func fileDescriptorBasedInterface_Set() error {
	var err error
	fileDescriptorBasedInterface_Once.Do(func() {
		fileDescriptorBasedInterface, err = gi.InterfaceNew("Gio", "FileDescriptorBased")
	})
	return err
}

type FileDescriptorBased struct {
	native unsafe.Pointer
}

func FileDescriptorBasedNewFromNative(native unsafe.Pointer) *FileDescriptorBased {
	instance := &FileDescriptorBased{native: native}

	return instance
}

/*
CastToFileDescriptorBased down casts any arbitrary Object to FileDescriptorBased.
Exercise care, as this is a potentially dangerous function
if the Object is not a FileDescriptorBased.
*/
func CastToFileDescriptorBased(object *gobject.Object) *FileDescriptorBased {
	return FileDescriptorBasedNewFromNative(object.Native())
}

// Equals compares this FileDescriptorBased with another FileDescriptorBased, and returns true if they represent the same Object.
func (recv *FileDescriptorBased) Equals(other *FileDescriptorBased) bool {
	return other.Native() == recv.Native()
}

func (recv *FileDescriptorBased) Native() unsafe.Pointer {
	return recv.native
}

var fileDescriptorBasedGetFdFunction *gi.Function
var fileDescriptorBasedGetFdFunction_Once sync.Once

func fileDescriptorBasedGetFdFunction_Set() error {
	var err error
	fileDescriptorBasedGetFdFunction_Once.Do(func() {
		err = fileDescriptorBasedInterface_Set()
		if err != nil {
			return
		}
		fileDescriptorBasedGetFdFunction, err = fileDescriptorBasedInterface.InvokerNew("get_fd")
	})
	return err
}

// GetFd is a representation of the C type g_file_descriptor_based_get_fd.
func (recv *FileDescriptorBased) GetFd() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := fileDescriptorBasedGetFdFunction_Set()
	if err == nil {
		ret = fileDescriptorBasedGetFdFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var iconInterface *gi.Interface
var iconInterface_Once sync.Once

func iconInterface_Set() error {
	var err error
	iconInterface_Once.Do(func() {
		iconInterface, err = gi.InterfaceNew("Gio", "Icon")
	})
	return err
}

type Icon struct {
	native unsafe.Pointer
}

func IconNewFromNative(native unsafe.Pointer) *Icon {
	instance := &Icon{native: native}

	return instance
}

/*
CastToIcon down casts any arbitrary Object to Icon.
Exercise care, as this is a potentially dangerous function
if the Object is not a Icon.
*/
func CastToIcon(object *gobject.Object) *Icon {
	return IconNewFromNative(object.Native())
}

// Equals compares this Icon with another Icon, and returns true if they represent the same Object.
func (recv *Icon) Equals(other *Icon) bool {
	return other.Native() == recv.Native()
}

func (recv *Icon) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'g_icon_equal' : parameter 'icon2' of type 'Icon' not supported

// UNSUPPORTED : C value 'g_icon_serialize' : return type 'GLib.Variant' not supported

var iconToStringFunction *gi.Function
var iconToStringFunction_Once sync.Once

func iconToStringFunction_Set() error {
	var err error
	iconToStringFunction_Once.Do(func() {
		err = iconInterface_Set()
		if err != nil {
			return
		}
		iconToStringFunction, err = iconInterface.InvokerNew("to_string")
	})
	return err
}

// ToString is a representation of the C type g_icon_to_string.
func (recv *Icon) ToString() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := iconToStringFunction_Set()
	if err == nil {
		ret = iconToStringFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

var initableInterface *gi.Interface
var initableInterface_Once sync.Once

func initableInterface_Set() error {
	var err error
	initableInterface_Once.Do(func() {
		initableInterface, err = gi.InterfaceNew("Gio", "Initable")
	})
	return err
}

type Initable struct {
	native unsafe.Pointer
}

func InitableNewFromNative(native unsafe.Pointer) *Initable {
	instance := &Initable{native: native}

	return instance
}

/*
CastToInitable down casts any arbitrary Object to Initable.
Exercise care, as this is a potentially dangerous function
if the Object is not a Initable.
*/
func CastToInitable(object *gobject.Object) *Initable {
	return InitableNewFromNative(object.Native())
}

// Equals compares this Initable with another Initable, and returns true if they represent the same Object.
func (recv *Initable) Equals(other *Initable) bool {
	return other.Native() == recv.Native()
}

func (recv *Initable) Native() unsafe.Pointer {
	return recv.native
}

var initableInitFunction *gi.Function
var initableInitFunction_Once sync.Once

func initableInitFunction_Set() error {
	var err error
	initableInitFunction_Once.Do(func() {
		err = initableInterface_Set()
		if err != nil {
			return
		}
		initableInitFunction, err = initableInterface.InvokerNew("init")
	})
	return err
}

// Init is a representation of the C type g_initable_init.
func (recv *Initable) Init(cancellable *Cancellable) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(cancellable.Native())

	var ret gi.Argument

	err := initableInitFunction_Set()
	if err == nil {
		ret = initableInitFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var listModelInterface *gi.Interface
var listModelInterface_Once sync.Once

func listModelInterface_Set() error {
	var err error
	listModelInterface_Once.Do(func() {
		listModelInterface, err = gi.InterfaceNew("Gio", "ListModel")
	})
	return err
}

type ListModel struct {
	native unsafe.Pointer
}

func ListModelNewFromNative(native unsafe.Pointer) *ListModel {
	instance := &ListModel{native: native}

	return instance
}

/*
CastToListModel down casts any arbitrary Object to ListModel.
Exercise care, as this is a potentially dangerous function
if the Object is not a ListModel.
*/
func CastToListModel(object *gobject.Object) *ListModel {
	return ListModelNewFromNative(object.Native())
}

// Equals compares this ListModel with another ListModel, and returns true if they represent the same Object.
func (recv *ListModel) Equals(other *ListModel) bool {
	return other.Native() == recv.Native()
}

func (recv *ListModel) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'g_list_model_get_item' : return type 'gpointer' not supported

// UNSUPPORTED : C value 'g_list_model_get_item_type' : return type 'GType' not supported

var listModelGetNItemsFunction *gi.Function
var listModelGetNItemsFunction_Once sync.Once

func listModelGetNItemsFunction_Set() error {
	var err error
	listModelGetNItemsFunction_Once.Do(func() {
		err = listModelInterface_Set()
		if err != nil {
			return
		}
		listModelGetNItemsFunction, err = listModelInterface.InvokerNew("get_n_items")
	})
	return err
}

// GetNItems is a representation of the C type g_list_model_get_n_items.
func (recv *ListModel) GetNItems() uint32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := listModelGetNItemsFunction_Set()
	if err == nil {
		ret = listModelGetNItemsFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo
}

var listModelGetObjectFunction *gi.Function
var listModelGetObjectFunction_Once sync.Once

func listModelGetObjectFunction_Set() error {
	var err error
	listModelGetObjectFunction_Once.Do(func() {
		err = listModelInterface_Set()
		if err != nil {
			return
		}
		listModelGetObjectFunction, err = listModelInterface.InvokerNew("get_object")
	})
	return err
}

// GetObject is a representation of the C type g_list_model_get_object.
func (recv *ListModel) GetObject(position uint32) *gobject.Object {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetUint32(position)

	var ret gi.Argument

	err := listModelGetObjectFunction_Set()
	if err == nil {
		ret = listModelGetObjectFunction.Invoke(inArgs[:], nil)
	}

	retGo := gobject.ObjectNewFromNative(ret.Pointer())

	return retGo
}

var listModelItemsChangedFunction *gi.Function
var listModelItemsChangedFunction_Once sync.Once

func listModelItemsChangedFunction_Set() error {
	var err error
	listModelItemsChangedFunction_Once.Do(func() {
		err = listModelInterface_Set()
		if err != nil {
			return
		}
		listModelItemsChangedFunction, err = listModelInterface.InvokerNew("items_changed")
	})
	return err
}

// ItemsChanged is a representation of the C type g_list_model_items_changed.
func (recv *ListModel) ItemsChanged(position uint32, removed uint32, added uint32) {
	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetUint32(position)
	inArgs[2].SetUint32(removed)
	inArgs[3].SetUint32(added)

	err := listModelItemsChangedFunction_Set()
	if err == nil {
		listModelItemsChangedFunction.Invoke(inArgs[:], nil)
	}

	return
}

/*
ConnectItemsChanged connects a callback to the 'items-changed' signal of the ListModel.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *ListModel) ConnectItemsChanged(handler func(instance *ListModel, position uint32, removed uint32, added uint32)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {}

	return callback.ConnectSignal(recv.Native(), "items-changed", marshal)
}

/*
Disconnect disconnects a callback previously registered with a Connect...() method.

The connectionID should be a value returned from a call to a Connect...() method.
*/
func (recv *ListModel) DisconnectSignal(connectionID int) {
	callback.DisconnectSignal(connectionID)
}

var loadableIconInterface *gi.Interface
var loadableIconInterface_Once sync.Once

func loadableIconInterface_Set() error {
	var err error
	loadableIconInterface_Once.Do(func() {
		loadableIconInterface, err = gi.InterfaceNew("Gio", "LoadableIcon")
	})
	return err
}

type LoadableIcon struct {
	native unsafe.Pointer
}

func LoadableIconNewFromNative(native unsafe.Pointer) *LoadableIcon {
	instance := &LoadableIcon{native: native}

	return instance
}

/*
CastToLoadableIcon down casts any arbitrary Object to LoadableIcon.
Exercise care, as this is a potentially dangerous function
if the Object is not a LoadableIcon.
*/
func CastToLoadableIcon(object *gobject.Object) *LoadableIcon {
	return LoadableIconNewFromNative(object.Native())
}

// Equals compares this LoadableIcon with another LoadableIcon, and returns true if they represent the same Object.
func (recv *LoadableIcon) Equals(other *LoadableIcon) bool {
	return other.Native() == recv.Native()
}

func (recv *LoadableIcon) Native() unsafe.Pointer {
	return recv.native
}

var loadableIconLoadFunction *gi.Function
var loadableIconLoadFunction_Once sync.Once

func loadableIconLoadFunction_Set() error {
	var err error
	loadableIconLoadFunction_Once.Do(func() {
		err = loadableIconInterface_Set()
		if err != nil {
			return
		}
		loadableIconLoadFunction, err = loadableIconInterface.InvokerNew("load")
	})
	return err
}

// Load is a representation of the C type g_loadable_icon_load.
func (recv *LoadableIcon) Load(size int32, cancellable *Cancellable) (*InputStream, string) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(size)
	inArgs[2].SetPointer(cancellable.Native())

	var outArgs [1]gi.Argument
	var ret gi.Argument

	err := loadableIconLoadFunction_Set()
	if err == nil {
		ret = loadableIconLoadFunction.Invoke(inArgs[:], outArgs[:])
	}

	retGo := InputStreamNewFromNative(ret.Pointer())
	out0 := outArgs[0].String(true)

	return retGo, out0
}

// UNSUPPORTED : C value 'g_loadable_icon_load_async' : parameter 'callback' of type 'AsyncReadyCallback' not supported

// UNSUPPORTED : C value 'g_loadable_icon_load_finish' : parameter 'res' of type 'AsyncResult' not supported

var mountInterface *gi.Interface
var mountInterface_Once sync.Once

func mountInterface_Set() error {
	var err error
	mountInterface_Once.Do(func() {
		mountInterface, err = gi.InterfaceNew("Gio", "Mount")
	})
	return err
}

type Mount struct {
	native unsafe.Pointer
}

func MountNewFromNative(native unsafe.Pointer) *Mount {
	instance := &Mount{native: native}

	return instance
}

/*
CastToMount down casts any arbitrary Object to Mount.
Exercise care, as this is a potentially dangerous function
if the Object is not a Mount.
*/
func CastToMount(object *gobject.Object) *Mount {
	return MountNewFromNative(object.Native())
}

// Equals compares this Mount with another Mount, and returns true if they represent the same Object.
func (recv *Mount) Equals(other *Mount) bool {
	return other.Native() == recv.Native()
}

func (recv *Mount) Native() unsafe.Pointer {
	return recv.native
}

var mountCanEjectFunction *gi.Function
var mountCanEjectFunction_Once sync.Once

func mountCanEjectFunction_Set() error {
	var err error
	mountCanEjectFunction_Once.Do(func() {
		err = mountInterface_Set()
		if err != nil {
			return
		}
		mountCanEjectFunction, err = mountInterface.InvokerNew("can_eject")
	})
	return err
}

// CanEject is a representation of the C type g_mount_can_eject.
func (recv *Mount) CanEject() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := mountCanEjectFunction_Set()
	if err == nil {
		ret = mountCanEjectFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var mountCanUnmountFunction *gi.Function
var mountCanUnmountFunction_Once sync.Once

func mountCanUnmountFunction_Set() error {
	var err error
	mountCanUnmountFunction_Once.Do(func() {
		err = mountInterface_Set()
		if err != nil {
			return
		}
		mountCanUnmountFunction, err = mountInterface.InvokerNew("can_unmount")
	})
	return err
}

// CanUnmount is a representation of the C type g_mount_can_unmount.
func (recv *Mount) CanUnmount() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := mountCanUnmountFunction_Set()
	if err == nil {
		ret = mountCanUnmountFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'g_mount_eject' : parameter 'flags' of type 'MountUnmountFlags' not supported

// UNSUPPORTED : C value 'g_mount_eject_finish' : parameter 'result' of type 'AsyncResult' not supported

// UNSUPPORTED : C value 'g_mount_eject_with_operation' : parameter 'flags' of type 'MountUnmountFlags' not supported

// UNSUPPORTED : C value 'g_mount_eject_with_operation_finish' : parameter 'result' of type 'AsyncResult' not supported

// UNSUPPORTED : C value 'g_mount_get_default_location' : return type 'File' not supported

// UNSUPPORTED : C value 'g_mount_get_drive' : return type 'Drive' not supported

// UNSUPPORTED : C value 'g_mount_get_icon' : return type 'Icon' not supported

var mountGetNameFunction *gi.Function
var mountGetNameFunction_Once sync.Once

func mountGetNameFunction_Set() error {
	var err error
	mountGetNameFunction_Once.Do(func() {
		err = mountInterface_Set()
		if err != nil {
			return
		}
		mountGetNameFunction, err = mountInterface.InvokerNew("get_name")
	})
	return err
}

// GetName is a representation of the C type g_mount_get_name.
func (recv *Mount) GetName() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := mountGetNameFunction_Set()
	if err == nil {
		ret = mountGetNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

// UNSUPPORTED : C value 'g_mount_get_root' : return type 'File' not supported

var mountGetSortKeyFunction *gi.Function
var mountGetSortKeyFunction_Once sync.Once

func mountGetSortKeyFunction_Set() error {
	var err error
	mountGetSortKeyFunction_Once.Do(func() {
		err = mountInterface_Set()
		if err != nil {
			return
		}
		mountGetSortKeyFunction, err = mountInterface.InvokerNew("get_sort_key")
	})
	return err
}

// GetSortKey is a representation of the C type g_mount_get_sort_key.
func (recv *Mount) GetSortKey() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := mountGetSortKeyFunction_Set()
	if err == nil {
		ret = mountGetSortKeyFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'g_mount_get_symbolic_icon' : return type 'Icon' not supported

var mountGetUuidFunction *gi.Function
var mountGetUuidFunction_Once sync.Once

func mountGetUuidFunction_Set() error {
	var err error
	mountGetUuidFunction_Once.Do(func() {
		err = mountInterface_Set()
		if err != nil {
			return
		}
		mountGetUuidFunction, err = mountInterface.InvokerNew("get_uuid")
	})
	return err
}

// GetUuid is a representation of the C type g_mount_get_uuid.
func (recv *Mount) GetUuid() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := mountGetUuidFunction_Set()
	if err == nil {
		ret = mountGetUuidFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

// UNSUPPORTED : C value 'g_mount_get_volume' : return type 'Volume' not supported

// UNSUPPORTED : C value 'g_mount_guess_content_type' : parameter 'callback' of type 'AsyncReadyCallback' not supported

// UNSUPPORTED : C value 'g_mount_guess_content_type_finish' : parameter 'result' of type 'AsyncResult' not supported

var mountGuessContentTypeSyncFunction *gi.Function
var mountGuessContentTypeSyncFunction_Once sync.Once

func mountGuessContentTypeSyncFunction_Set() error {
	var err error
	mountGuessContentTypeSyncFunction_Once.Do(func() {
		err = mountInterface_Set()
		if err != nil {
			return
		}
		mountGuessContentTypeSyncFunction, err = mountInterface.InvokerNew("guess_content_type_sync")
	})
	return err
}

// GuessContentTypeSync is a representation of the C type g_mount_guess_content_type_sync.
func (recv *Mount) GuessContentTypeSync(forceRescan bool, cancellable *Cancellable) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetBoolean(forceRescan)
	inArgs[2].SetPointer(cancellable.Native())

	err := mountGuessContentTypeSyncFunction_Set()
	if err == nil {
		mountGuessContentTypeSyncFunction.Invoke(inArgs[:], nil)
	}

	return
}

var mountIsShadowedFunction *gi.Function
var mountIsShadowedFunction_Once sync.Once

func mountIsShadowedFunction_Set() error {
	var err error
	mountIsShadowedFunction_Once.Do(func() {
		err = mountInterface_Set()
		if err != nil {
			return
		}
		mountIsShadowedFunction, err = mountInterface.InvokerNew("is_shadowed")
	})
	return err
}

// IsShadowed is a representation of the C type g_mount_is_shadowed.
func (recv *Mount) IsShadowed() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := mountIsShadowedFunction_Set()
	if err == nil {
		ret = mountIsShadowedFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'g_mount_remount' : parameter 'flags' of type 'MountMountFlags' not supported

// UNSUPPORTED : C value 'g_mount_remount_finish' : parameter 'result' of type 'AsyncResult' not supported

var mountShadowFunction *gi.Function
var mountShadowFunction_Once sync.Once

func mountShadowFunction_Set() error {
	var err error
	mountShadowFunction_Once.Do(func() {
		err = mountInterface_Set()
		if err != nil {
			return
		}
		mountShadowFunction, err = mountInterface.InvokerNew("shadow")
	})
	return err
}

// Shadow is a representation of the C type g_mount_shadow.
func (recv *Mount) Shadow() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := mountShadowFunction_Set()
	if err == nil {
		mountShadowFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'g_mount_unmount' : parameter 'flags' of type 'MountUnmountFlags' not supported

// UNSUPPORTED : C value 'g_mount_unmount_finish' : parameter 'result' of type 'AsyncResult' not supported

// UNSUPPORTED : C value 'g_mount_unmount_with_operation' : parameter 'flags' of type 'MountUnmountFlags' not supported

// UNSUPPORTED : C value 'g_mount_unmount_with_operation_finish' : parameter 'result' of type 'AsyncResult' not supported

var mountUnshadowFunction *gi.Function
var mountUnshadowFunction_Once sync.Once

func mountUnshadowFunction_Set() error {
	var err error
	mountUnshadowFunction_Once.Do(func() {
		err = mountInterface_Set()
		if err != nil {
			return
		}
		mountUnshadowFunction, err = mountInterface.InvokerNew("unshadow")
	})
	return err
}

// Unshadow is a representation of the C type g_mount_unshadow.
func (recv *Mount) Unshadow() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := mountUnshadowFunction_Set()
	if err == nil {
		mountUnshadowFunction.Invoke(inArgs[:], nil)
	}

	return
}

/*
ConnectChanged connects a callback to the 'changed' signal of the Mount.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *Mount) ConnectChanged(handler func(instance *Mount)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {}

	return callback.ConnectSignal(recv.Native(), "changed", marshal)
}

/*
ConnectPreUnmount connects a callback to the 'pre-unmount' signal of the Mount.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *Mount) ConnectPreUnmount(handler func(instance *Mount)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {}

	return callback.ConnectSignal(recv.Native(), "pre-unmount", marshal)
}

/*
ConnectUnmounted connects a callback to the 'unmounted' signal of the Mount.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *Mount) ConnectUnmounted(handler func(instance *Mount)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {}

	return callback.ConnectSignal(recv.Native(), "unmounted", marshal)
}

/*
Disconnect disconnects a callback previously registered with a Connect...() method.

The connectionID should be a value returned from a call to a Connect...() method.
*/
func (recv *Mount) DisconnectSignal(connectionID int) {
	callback.DisconnectSignal(connectionID)
}

var networkMonitorInterface *gi.Interface
var networkMonitorInterface_Once sync.Once

func networkMonitorInterface_Set() error {
	var err error
	networkMonitorInterface_Once.Do(func() {
		networkMonitorInterface, err = gi.InterfaceNew("Gio", "NetworkMonitor")
	})
	return err
}

type NetworkMonitor struct {
	native unsafe.Pointer
}

func NetworkMonitorNewFromNative(native unsafe.Pointer) *NetworkMonitor {
	instance := &NetworkMonitor{native: native}

	return instance
}

/*
CastToNetworkMonitor down casts any arbitrary Object to NetworkMonitor.
Exercise care, as this is a potentially dangerous function
if the Object is not a NetworkMonitor.
*/
func CastToNetworkMonitor(object *gobject.Object) *NetworkMonitor {
	return NetworkMonitorNewFromNative(object.Native())
}

// Equals compares this NetworkMonitor with another NetworkMonitor, and returns true if they represent the same Object.
func (recv *NetworkMonitor) Equals(other *NetworkMonitor) bool {
	return other.Native() == recv.Native()
}

func (recv *NetworkMonitor) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'g_network_monitor_can_reach' : parameter 'connectable' of type 'SocketConnectable' not supported

// UNSUPPORTED : C value 'g_network_monitor_can_reach_async' : parameter 'connectable' of type 'SocketConnectable' not supported

// UNSUPPORTED : C value 'g_network_monitor_can_reach_finish' : parameter 'result' of type 'AsyncResult' not supported

var networkMonitorGetConnectivityFunction *gi.Function
var networkMonitorGetConnectivityFunction_Once sync.Once

func networkMonitorGetConnectivityFunction_Set() error {
	var err error
	networkMonitorGetConnectivityFunction_Once.Do(func() {
		err = networkMonitorInterface_Set()
		if err != nil {
			return
		}
		networkMonitorGetConnectivityFunction, err = networkMonitorInterface.InvokerNew("get_connectivity")
	})
	return err
}

// GetConnectivity is a representation of the C type g_network_monitor_get_connectivity.
func (recv *NetworkMonitor) GetConnectivity() NetworkConnectivity {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := networkMonitorGetConnectivityFunction_Set()
	if err == nil {
		ret = networkMonitorGetConnectivityFunction.Invoke(inArgs[:], nil)
	}

	retGo := NetworkConnectivity(ret.Int32())

	return retGo
}

var networkMonitorGetNetworkAvailableFunction *gi.Function
var networkMonitorGetNetworkAvailableFunction_Once sync.Once

func networkMonitorGetNetworkAvailableFunction_Set() error {
	var err error
	networkMonitorGetNetworkAvailableFunction_Once.Do(func() {
		err = networkMonitorInterface_Set()
		if err != nil {
			return
		}
		networkMonitorGetNetworkAvailableFunction, err = networkMonitorInterface.InvokerNew("get_network_available")
	})
	return err
}

// GetNetworkAvailable is a representation of the C type g_network_monitor_get_network_available.
func (recv *NetworkMonitor) GetNetworkAvailable() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := networkMonitorGetNetworkAvailableFunction_Set()
	if err == nil {
		ret = networkMonitorGetNetworkAvailableFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var networkMonitorGetNetworkMeteredFunction *gi.Function
var networkMonitorGetNetworkMeteredFunction_Once sync.Once

func networkMonitorGetNetworkMeteredFunction_Set() error {
	var err error
	networkMonitorGetNetworkMeteredFunction_Once.Do(func() {
		err = networkMonitorInterface_Set()
		if err != nil {
			return
		}
		networkMonitorGetNetworkMeteredFunction, err = networkMonitorInterface.InvokerNew("get_network_metered")
	})
	return err
}

// GetNetworkMetered is a representation of the C type g_network_monitor_get_network_metered.
func (recv *NetworkMonitor) GetNetworkMetered() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := networkMonitorGetNetworkMeteredFunction_Set()
	if err == nil {
		ret = networkMonitorGetNetworkMeteredFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

/*
ConnectNetworkChanged connects a callback to the 'network-changed' signal of the NetworkMonitor.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *NetworkMonitor) ConnectNetworkChanged(handler func(instance *NetworkMonitor, networkAvailable bool)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {}

	return callback.ConnectSignal(recv.Native(), "network-changed", marshal)
}

/*
Disconnect disconnects a callback previously registered with a Connect...() method.

The connectionID should be a value returned from a call to a Connect...() method.
*/
func (recv *NetworkMonitor) DisconnectSignal(connectionID int) {
	callback.DisconnectSignal(connectionID)
}

var pollableInputStreamInterface *gi.Interface
var pollableInputStreamInterface_Once sync.Once

func pollableInputStreamInterface_Set() error {
	var err error
	pollableInputStreamInterface_Once.Do(func() {
		pollableInputStreamInterface, err = gi.InterfaceNew("Gio", "PollableInputStream")
	})
	return err
}

type PollableInputStream struct {
	native unsafe.Pointer
}

func PollableInputStreamNewFromNative(native unsafe.Pointer) *PollableInputStream {
	instance := &PollableInputStream{native: native}

	return instance
}

/*
CastToPollableInputStream down casts any arbitrary Object to PollableInputStream.
Exercise care, as this is a potentially dangerous function
if the Object is not a PollableInputStream.
*/
func CastToPollableInputStream(object *gobject.Object) *PollableInputStream {
	return PollableInputStreamNewFromNative(object.Native())
}

// Equals compares this PollableInputStream with another PollableInputStream, and returns true if they represent the same Object.
func (recv *PollableInputStream) Equals(other *PollableInputStream) bool {
	return other.Native() == recv.Native()
}

func (recv *PollableInputStream) Native() unsafe.Pointer {
	return recv.native
}

var pollableInputStreamCanPollFunction *gi.Function
var pollableInputStreamCanPollFunction_Once sync.Once

func pollableInputStreamCanPollFunction_Set() error {
	var err error
	pollableInputStreamCanPollFunction_Once.Do(func() {
		err = pollableInputStreamInterface_Set()
		if err != nil {
			return
		}
		pollableInputStreamCanPollFunction, err = pollableInputStreamInterface.InvokerNew("can_poll")
	})
	return err
}

// CanPoll is a representation of the C type g_pollable_input_stream_can_poll.
func (recv *PollableInputStream) CanPoll() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := pollableInputStreamCanPollFunction_Set()
	if err == nil {
		ret = pollableInputStreamCanPollFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'g_pollable_input_stream_create_source' : return type 'GLib.Source' not supported

var pollableInputStreamIsReadableFunction *gi.Function
var pollableInputStreamIsReadableFunction_Once sync.Once

func pollableInputStreamIsReadableFunction_Set() error {
	var err error
	pollableInputStreamIsReadableFunction_Once.Do(func() {
		err = pollableInputStreamInterface_Set()
		if err != nil {
			return
		}
		pollableInputStreamIsReadableFunction, err = pollableInputStreamInterface.InvokerNew("is_readable")
	})
	return err
}

// IsReadable is a representation of the C type g_pollable_input_stream_is_readable.
func (recv *PollableInputStream) IsReadable() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := pollableInputStreamIsReadableFunction_Set()
	if err == nil {
		ret = pollableInputStreamIsReadableFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'g_pollable_input_stream_read_nonblocking' : parameter 'buffer' of type 'nil' not supported

var pollableOutputStreamInterface *gi.Interface
var pollableOutputStreamInterface_Once sync.Once

func pollableOutputStreamInterface_Set() error {
	var err error
	pollableOutputStreamInterface_Once.Do(func() {
		pollableOutputStreamInterface, err = gi.InterfaceNew("Gio", "PollableOutputStream")
	})
	return err
}

type PollableOutputStream struct {
	native unsafe.Pointer
}

func PollableOutputStreamNewFromNative(native unsafe.Pointer) *PollableOutputStream {
	instance := &PollableOutputStream{native: native}

	return instance
}

/*
CastToPollableOutputStream down casts any arbitrary Object to PollableOutputStream.
Exercise care, as this is a potentially dangerous function
if the Object is not a PollableOutputStream.
*/
func CastToPollableOutputStream(object *gobject.Object) *PollableOutputStream {
	return PollableOutputStreamNewFromNative(object.Native())
}

// Equals compares this PollableOutputStream with another PollableOutputStream, and returns true if they represent the same Object.
func (recv *PollableOutputStream) Equals(other *PollableOutputStream) bool {
	return other.Native() == recv.Native()
}

func (recv *PollableOutputStream) Native() unsafe.Pointer {
	return recv.native
}

var pollableOutputStreamCanPollFunction *gi.Function
var pollableOutputStreamCanPollFunction_Once sync.Once

func pollableOutputStreamCanPollFunction_Set() error {
	var err error
	pollableOutputStreamCanPollFunction_Once.Do(func() {
		err = pollableOutputStreamInterface_Set()
		if err != nil {
			return
		}
		pollableOutputStreamCanPollFunction, err = pollableOutputStreamInterface.InvokerNew("can_poll")
	})
	return err
}

// CanPoll is a representation of the C type g_pollable_output_stream_can_poll.
func (recv *PollableOutputStream) CanPoll() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := pollableOutputStreamCanPollFunction_Set()
	if err == nil {
		ret = pollableOutputStreamCanPollFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'g_pollable_output_stream_create_source' : return type 'GLib.Source' not supported

var pollableOutputStreamIsWritableFunction *gi.Function
var pollableOutputStreamIsWritableFunction_Once sync.Once

func pollableOutputStreamIsWritableFunction_Set() error {
	var err error
	pollableOutputStreamIsWritableFunction_Once.Do(func() {
		err = pollableOutputStreamInterface_Set()
		if err != nil {
			return
		}
		pollableOutputStreamIsWritableFunction, err = pollableOutputStreamInterface.InvokerNew("is_writable")
	})
	return err
}

// IsWritable is a representation of the C type g_pollable_output_stream_is_writable.
func (recv *PollableOutputStream) IsWritable() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := pollableOutputStreamIsWritableFunction_Set()
	if err == nil {
		ret = pollableOutputStreamIsWritableFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'g_pollable_output_stream_write_nonblocking' : parameter 'buffer' of type 'nil' not supported

// UNSUPPORTED : C value 'g_pollable_output_stream_writev_nonblocking' : parameter 'vectors' of type 'nil' not supported

var proxyInterface *gi.Interface
var proxyInterface_Once sync.Once

func proxyInterface_Set() error {
	var err error
	proxyInterface_Once.Do(func() {
		proxyInterface, err = gi.InterfaceNew("Gio", "Proxy")
	})
	return err
}

type Proxy struct {
	native unsafe.Pointer
}

func ProxyNewFromNative(native unsafe.Pointer) *Proxy {
	instance := &Proxy{native: native}

	return instance
}

/*
CastToProxy down casts any arbitrary Object to Proxy.
Exercise care, as this is a potentially dangerous function
if the Object is not a Proxy.
*/
func CastToProxy(object *gobject.Object) *Proxy {
	return ProxyNewFromNative(object.Native())
}

// Equals compares this Proxy with another Proxy, and returns true if they represent the same Object.
func (recv *Proxy) Equals(other *Proxy) bool {
	return other.Native() == recv.Native()
}

func (recv *Proxy) Native() unsafe.Pointer {
	return recv.native
}

var proxyConnectFunction *gi.Function
var proxyConnectFunction_Once sync.Once

func proxyConnectFunction_Set() error {
	var err error
	proxyConnectFunction_Once.Do(func() {
		err = proxyInterface_Set()
		if err != nil {
			return
		}
		proxyConnectFunction, err = proxyInterface.InvokerNew("connect")
	})
	return err
}

// Connect is a representation of the C type g_proxy_connect.
func (recv *Proxy) Connect(connection *IOStream, proxyAddress *ProxyAddress, cancellable *Cancellable) *IOStream {
	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(connection.Native())
	inArgs[2].SetPointer(proxyAddress.Native())
	inArgs[3].SetPointer(cancellable.Native())

	var ret gi.Argument

	err := proxyConnectFunction_Set()
	if err == nil {
		ret = proxyConnectFunction.Invoke(inArgs[:], nil)
	}

	retGo := IOStreamNewFromNative(ret.Pointer())

	return retGo
}

// UNSUPPORTED : C value 'g_proxy_connect_async' : parameter 'callback' of type 'AsyncReadyCallback' not supported

// UNSUPPORTED : C value 'g_proxy_connect_finish' : parameter 'result' of type 'AsyncResult' not supported

var proxySupportsHostnameFunction *gi.Function
var proxySupportsHostnameFunction_Once sync.Once

func proxySupportsHostnameFunction_Set() error {
	var err error
	proxySupportsHostnameFunction_Once.Do(func() {
		err = proxyInterface_Set()
		if err != nil {
			return
		}
		proxySupportsHostnameFunction, err = proxyInterface.InvokerNew("supports_hostname")
	})
	return err
}

// SupportsHostname is a representation of the C type g_proxy_supports_hostname.
func (recv *Proxy) SupportsHostname() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := proxySupportsHostnameFunction_Set()
	if err == nil {
		ret = proxySupportsHostnameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var proxyResolverInterface *gi.Interface
var proxyResolverInterface_Once sync.Once

func proxyResolverInterface_Set() error {
	var err error
	proxyResolverInterface_Once.Do(func() {
		proxyResolverInterface, err = gi.InterfaceNew("Gio", "ProxyResolver")
	})
	return err
}

type ProxyResolver struct {
	native unsafe.Pointer
}

func ProxyResolverNewFromNative(native unsafe.Pointer) *ProxyResolver {
	instance := &ProxyResolver{native: native}

	return instance
}

/*
CastToProxyResolver down casts any arbitrary Object to ProxyResolver.
Exercise care, as this is a potentially dangerous function
if the Object is not a ProxyResolver.
*/
func CastToProxyResolver(object *gobject.Object) *ProxyResolver {
	return ProxyResolverNewFromNative(object.Native())
}

// Equals compares this ProxyResolver with another ProxyResolver, and returns true if they represent the same Object.
func (recv *ProxyResolver) Equals(other *ProxyResolver) bool {
	return other.Native() == recv.Native()
}

func (recv *ProxyResolver) Native() unsafe.Pointer {
	return recv.native
}

var proxyResolverIsSupportedFunction *gi.Function
var proxyResolverIsSupportedFunction_Once sync.Once

func proxyResolverIsSupportedFunction_Set() error {
	var err error
	proxyResolverIsSupportedFunction_Once.Do(func() {
		err = proxyResolverInterface_Set()
		if err != nil {
			return
		}
		proxyResolverIsSupportedFunction, err = proxyResolverInterface.InvokerNew("is_supported")
	})
	return err
}

// IsSupported is a representation of the C type g_proxy_resolver_is_supported.
func (recv *ProxyResolver) IsSupported() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := proxyResolverIsSupportedFunction_Set()
	if err == nil {
		ret = proxyResolverIsSupportedFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var proxyResolverLookupFunction *gi.Function
var proxyResolverLookupFunction_Once sync.Once

func proxyResolverLookupFunction_Set() error {
	var err error
	proxyResolverLookupFunction_Once.Do(func() {
		err = proxyResolverInterface_Set()
		if err != nil {
			return
		}
		proxyResolverLookupFunction, err = proxyResolverInterface.InvokerNew("lookup")
	})
	return err
}

// Lookup is a representation of the C type g_proxy_resolver_lookup.
func (recv *ProxyResolver) Lookup(uri string, cancellable *Cancellable) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(uri)
	inArgs[2].SetPointer(cancellable.Native())

	err := proxyResolverLookupFunction_Set()
	if err == nil {
		proxyResolverLookupFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'g_proxy_resolver_lookup_async' : parameter 'callback' of type 'AsyncReadyCallback' not supported

// UNSUPPORTED : C value 'g_proxy_resolver_lookup_finish' : parameter 'result' of type 'AsyncResult' not supported

var remoteActionGroupInterface *gi.Interface
var remoteActionGroupInterface_Once sync.Once

func remoteActionGroupInterface_Set() error {
	var err error
	remoteActionGroupInterface_Once.Do(func() {
		remoteActionGroupInterface, err = gi.InterfaceNew("Gio", "RemoteActionGroup")
	})
	return err
}

type RemoteActionGroup struct {
	native unsafe.Pointer
}

func RemoteActionGroupNewFromNative(native unsafe.Pointer) *RemoteActionGroup {
	instance := &RemoteActionGroup{native: native}

	return instance
}

/*
CastToRemoteActionGroup down casts any arbitrary Object to RemoteActionGroup.
Exercise care, as this is a potentially dangerous function
if the Object is not a RemoteActionGroup.
*/
func CastToRemoteActionGroup(object *gobject.Object) *RemoteActionGroup {
	return RemoteActionGroupNewFromNative(object.Native())
}

// Equals compares this RemoteActionGroup with another RemoteActionGroup, and returns true if they represent the same Object.
func (recv *RemoteActionGroup) Equals(other *RemoteActionGroup) bool {
	return other.Native() == recv.Native()
}

func (recv *RemoteActionGroup) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'g_remote_action_group_activate_action_full' : parameter 'parameter' of type 'GLib.Variant' not supported

// UNSUPPORTED : C value 'g_remote_action_group_change_action_state_full' : parameter 'value' of type 'GLib.Variant' not supported

var seekableInterface *gi.Interface
var seekableInterface_Once sync.Once

func seekableInterface_Set() error {
	var err error
	seekableInterface_Once.Do(func() {
		seekableInterface, err = gi.InterfaceNew("Gio", "Seekable")
	})
	return err
}

type Seekable struct {
	native unsafe.Pointer
}

func SeekableNewFromNative(native unsafe.Pointer) *Seekable {
	instance := &Seekable{native: native}

	return instance
}

/*
CastToSeekable down casts any arbitrary Object to Seekable.
Exercise care, as this is a potentially dangerous function
if the Object is not a Seekable.
*/
func CastToSeekable(object *gobject.Object) *Seekable {
	return SeekableNewFromNative(object.Native())
}

// Equals compares this Seekable with another Seekable, and returns true if they represent the same Object.
func (recv *Seekable) Equals(other *Seekable) bool {
	return other.Native() == recv.Native()
}

func (recv *Seekable) Native() unsafe.Pointer {
	return recv.native
}

var seekableCanSeekFunction *gi.Function
var seekableCanSeekFunction_Once sync.Once

func seekableCanSeekFunction_Set() error {
	var err error
	seekableCanSeekFunction_Once.Do(func() {
		err = seekableInterface_Set()
		if err != nil {
			return
		}
		seekableCanSeekFunction, err = seekableInterface.InvokerNew("can_seek")
	})
	return err
}

// CanSeek is a representation of the C type g_seekable_can_seek.
func (recv *Seekable) CanSeek() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := seekableCanSeekFunction_Set()
	if err == nil {
		ret = seekableCanSeekFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var seekableCanTruncateFunction *gi.Function
var seekableCanTruncateFunction_Once sync.Once

func seekableCanTruncateFunction_Set() error {
	var err error
	seekableCanTruncateFunction_Once.Do(func() {
		err = seekableInterface_Set()
		if err != nil {
			return
		}
		seekableCanTruncateFunction, err = seekableInterface.InvokerNew("can_truncate")
	})
	return err
}

// CanTruncate is a representation of the C type g_seekable_can_truncate.
func (recv *Seekable) CanTruncate() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := seekableCanTruncateFunction_Set()
	if err == nil {
		ret = seekableCanTruncateFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'g_seekable_seek' : parameter 'type' of type 'GLib.SeekType' not supported

var seekableTellFunction *gi.Function
var seekableTellFunction_Once sync.Once

func seekableTellFunction_Set() error {
	var err error
	seekableTellFunction_Once.Do(func() {
		err = seekableInterface_Set()
		if err != nil {
			return
		}
		seekableTellFunction, err = seekableInterface.InvokerNew("tell")
	})
	return err
}

// Tell is a representation of the C type g_seekable_tell.
func (recv *Seekable) Tell() int64 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := seekableTellFunction_Set()
	if err == nil {
		ret = seekableTellFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int64()

	return retGo
}

var seekableTruncateFunction *gi.Function
var seekableTruncateFunction_Once sync.Once

func seekableTruncateFunction_Set() error {
	var err error
	seekableTruncateFunction_Once.Do(func() {
		err = seekableInterface_Set()
		if err != nil {
			return
		}
		seekableTruncateFunction, err = seekableInterface.InvokerNew("truncate")
	})
	return err
}

// Truncate is a representation of the C type g_seekable_truncate.
func (recv *Seekable) Truncate(offset int64, cancellable *Cancellable) bool {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt64(offset)
	inArgs[2].SetPointer(cancellable.Native())

	var ret gi.Argument

	err := seekableTruncateFunction_Set()
	if err == nil {
		ret = seekableTruncateFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var socketConnectableInterface *gi.Interface
var socketConnectableInterface_Once sync.Once

func socketConnectableInterface_Set() error {
	var err error
	socketConnectableInterface_Once.Do(func() {
		socketConnectableInterface, err = gi.InterfaceNew("Gio", "SocketConnectable")
	})
	return err
}

type SocketConnectable struct {
	native unsafe.Pointer
}

func SocketConnectableNewFromNative(native unsafe.Pointer) *SocketConnectable {
	instance := &SocketConnectable{native: native}

	return instance
}

/*
CastToSocketConnectable down casts any arbitrary Object to SocketConnectable.
Exercise care, as this is a potentially dangerous function
if the Object is not a SocketConnectable.
*/
func CastToSocketConnectable(object *gobject.Object) *SocketConnectable {
	return SocketConnectableNewFromNative(object.Native())
}

// Equals compares this SocketConnectable with another SocketConnectable, and returns true if they represent the same Object.
func (recv *SocketConnectable) Equals(other *SocketConnectable) bool {
	return other.Native() == recv.Native()
}

func (recv *SocketConnectable) Native() unsafe.Pointer {
	return recv.native
}

var socketConnectableEnumerateFunction *gi.Function
var socketConnectableEnumerateFunction_Once sync.Once

func socketConnectableEnumerateFunction_Set() error {
	var err error
	socketConnectableEnumerateFunction_Once.Do(func() {
		err = socketConnectableInterface_Set()
		if err != nil {
			return
		}
		socketConnectableEnumerateFunction, err = socketConnectableInterface.InvokerNew("enumerate")
	})
	return err
}

// Enumerate is a representation of the C type g_socket_connectable_enumerate.
func (recv *SocketConnectable) Enumerate() *SocketAddressEnumerator {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := socketConnectableEnumerateFunction_Set()
	if err == nil {
		ret = socketConnectableEnumerateFunction.Invoke(inArgs[:], nil)
	}

	retGo := SocketAddressEnumeratorNewFromNative(ret.Pointer())

	return retGo
}

var socketConnectableProxyEnumerateFunction *gi.Function
var socketConnectableProxyEnumerateFunction_Once sync.Once

func socketConnectableProxyEnumerateFunction_Set() error {
	var err error
	socketConnectableProxyEnumerateFunction_Once.Do(func() {
		err = socketConnectableInterface_Set()
		if err != nil {
			return
		}
		socketConnectableProxyEnumerateFunction, err = socketConnectableInterface.InvokerNew("proxy_enumerate")
	})
	return err
}

// ProxyEnumerate is a representation of the C type g_socket_connectable_proxy_enumerate.
func (recv *SocketConnectable) ProxyEnumerate() *SocketAddressEnumerator {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := socketConnectableProxyEnumerateFunction_Set()
	if err == nil {
		ret = socketConnectableProxyEnumerateFunction.Invoke(inArgs[:], nil)
	}

	retGo := SocketAddressEnumeratorNewFromNative(ret.Pointer())

	return retGo
}

var socketConnectableToStringFunction *gi.Function
var socketConnectableToStringFunction_Once sync.Once

func socketConnectableToStringFunction_Set() error {
	var err error
	socketConnectableToStringFunction_Once.Do(func() {
		err = socketConnectableInterface_Set()
		if err != nil {
			return
		}
		socketConnectableToStringFunction, err = socketConnectableInterface.InvokerNew("to_string")
	})
	return err
}

// ToString is a representation of the C type g_socket_connectable_to_string.
func (recv *SocketConnectable) ToString() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := socketConnectableToStringFunction_Set()
	if err == nil {
		ret = socketConnectableToStringFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

var tlsBackendInterface *gi.Interface
var tlsBackendInterface_Once sync.Once

func tlsBackendInterface_Set() error {
	var err error
	tlsBackendInterface_Once.Do(func() {
		tlsBackendInterface, err = gi.InterfaceNew("Gio", "TlsBackend")
	})
	return err
}

type TlsBackend struct {
	native unsafe.Pointer
}

func TlsBackendNewFromNative(native unsafe.Pointer) *TlsBackend {
	instance := &TlsBackend{native: native}

	return instance
}

/*
CastToTlsBackend down casts any arbitrary Object to TlsBackend.
Exercise care, as this is a potentially dangerous function
if the Object is not a TlsBackend.
*/
func CastToTlsBackend(object *gobject.Object) *TlsBackend {
	return TlsBackendNewFromNative(object.Native())
}

// Equals compares this TlsBackend with another TlsBackend, and returns true if they represent the same Object.
func (recv *TlsBackend) Equals(other *TlsBackend) bool {
	return other.Native() == recv.Native()
}

func (recv *TlsBackend) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'g_tls_backend_get_certificate_type' : return type 'GType' not supported

// UNSUPPORTED : C value 'g_tls_backend_get_client_connection_type' : return type 'GType' not supported

var tlsBackendGetDefaultDatabaseFunction *gi.Function
var tlsBackendGetDefaultDatabaseFunction_Once sync.Once

func tlsBackendGetDefaultDatabaseFunction_Set() error {
	var err error
	tlsBackendGetDefaultDatabaseFunction_Once.Do(func() {
		err = tlsBackendInterface_Set()
		if err != nil {
			return
		}
		tlsBackendGetDefaultDatabaseFunction, err = tlsBackendInterface.InvokerNew("get_default_database")
	})
	return err
}

// GetDefaultDatabase is a representation of the C type g_tls_backend_get_default_database.
func (recv *TlsBackend) GetDefaultDatabase() *TlsDatabase {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := tlsBackendGetDefaultDatabaseFunction_Set()
	if err == nil {
		ret = tlsBackendGetDefaultDatabaseFunction.Invoke(inArgs[:], nil)
	}

	retGo := TlsDatabaseNewFromNative(ret.Pointer())

	return retGo
}

// UNSUPPORTED : C value 'g_tls_backend_get_dtls_client_connection_type' : return type 'GType' not supported

// UNSUPPORTED : C value 'g_tls_backend_get_dtls_server_connection_type' : return type 'GType' not supported

// UNSUPPORTED : C value 'g_tls_backend_get_file_database_type' : return type 'GType' not supported

// UNSUPPORTED : C value 'g_tls_backend_get_server_connection_type' : return type 'GType' not supported

var tlsBackendSetDefaultDatabaseFunction *gi.Function
var tlsBackendSetDefaultDatabaseFunction_Once sync.Once

func tlsBackendSetDefaultDatabaseFunction_Set() error {
	var err error
	tlsBackendSetDefaultDatabaseFunction_Once.Do(func() {
		err = tlsBackendInterface_Set()
		if err != nil {
			return
		}
		tlsBackendSetDefaultDatabaseFunction, err = tlsBackendInterface.InvokerNew("set_default_database")
	})
	return err
}

// SetDefaultDatabase is a representation of the C type g_tls_backend_set_default_database.
func (recv *TlsBackend) SetDefaultDatabase(database *TlsDatabase) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(database.Native())

	err := tlsBackendSetDefaultDatabaseFunction_Set()
	if err == nil {
		tlsBackendSetDefaultDatabaseFunction.Invoke(inArgs[:], nil)
	}

	return
}

var tlsBackendSupportsDtlsFunction *gi.Function
var tlsBackendSupportsDtlsFunction_Once sync.Once

func tlsBackendSupportsDtlsFunction_Set() error {
	var err error
	tlsBackendSupportsDtlsFunction_Once.Do(func() {
		err = tlsBackendInterface_Set()
		if err != nil {
			return
		}
		tlsBackendSupportsDtlsFunction, err = tlsBackendInterface.InvokerNew("supports_dtls")
	})
	return err
}

// SupportsDtls is a representation of the C type g_tls_backend_supports_dtls.
func (recv *TlsBackend) SupportsDtls() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := tlsBackendSupportsDtlsFunction_Set()
	if err == nil {
		ret = tlsBackendSupportsDtlsFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var tlsBackendSupportsTlsFunction *gi.Function
var tlsBackendSupportsTlsFunction_Once sync.Once

func tlsBackendSupportsTlsFunction_Set() error {
	var err error
	tlsBackendSupportsTlsFunction_Once.Do(func() {
		err = tlsBackendInterface_Set()
		if err != nil {
			return
		}
		tlsBackendSupportsTlsFunction, err = tlsBackendInterface.InvokerNew("supports_tls")
	})
	return err
}

// SupportsTls is a representation of the C type g_tls_backend_supports_tls.
func (recv *TlsBackend) SupportsTls() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := tlsBackendSupportsTlsFunction_Set()
	if err == nil {
		ret = tlsBackendSupportsTlsFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var tlsClientConnectionInterface *gi.Interface
var tlsClientConnectionInterface_Once sync.Once

func tlsClientConnectionInterface_Set() error {
	var err error
	tlsClientConnectionInterface_Once.Do(func() {
		tlsClientConnectionInterface, err = gi.InterfaceNew("Gio", "TlsClientConnection")
	})
	return err
}

type TlsClientConnection struct {
	native unsafe.Pointer
}

func TlsClientConnectionNewFromNative(native unsafe.Pointer) *TlsClientConnection {
	instance := &TlsClientConnection{native: native}

	return instance
}

/*
CastToTlsClientConnection down casts any arbitrary Object to TlsClientConnection.
Exercise care, as this is a potentially dangerous function
if the Object is not a TlsClientConnection.
*/
func CastToTlsClientConnection(object *gobject.Object) *TlsClientConnection {
	return TlsClientConnectionNewFromNative(object.Native())
}

// Equals compares this TlsClientConnection with another TlsClientConnection, and returns true if they represent the same Object.
func (recv *TlsClientConnection) Equals(other *TlsClientConnection) bool {
	return other.Native() == recv.Native()
}

func (recv *TlsClientConnection) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'g_tls_client_connection_copy_session_state' : parameter 'source' of type 'TlsClientConnection' not supported

// UNSUPPORTED : C value 'g_tls_client_connection_get_accepted_cas' : return type 'GLib.List' not supported

// UNSUPPORTED : C value 'g_tls_client_connection_get_server_identity' : return type 'SocketConnectable' not supported

var tlsClientConnectionGetUseSsl3Function *gi.Function
var tlsClientConnectionGetUseSsl3Function_Once sync.Once

func tlsClientConnectionGetUseSsl3Function_Set() error {
	var err error
	tlsClientConnectionGetUseSsl3Function_Once.Do(func() {
		err = tlsClientConnectionInterface_Set()
		if err != nil {
			return
		}
		tlsClientConnectionGetUseSsl3Function, err = tlsClientConnectionInterface.InvokerNew("get_use_ssl3")
	})
	return err
}

// GetUseSsl3 is a representation of the C type g_tls_client_connection_get_use_ssl3.
func (recv *TlsClientConnection) GetUseSsl3() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := tlsClientConnectionGetUseSsl3Function_Set()
	if err == nil {
		ret = tlsClientConnectionGetUseSsl3Function.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'g_tls_client_connection_get_validation_flags' : return type 'TlsCertificateFlags' not supported

// UNSUPPORTED : C value 'g_tls_client_connection_set_server_identity' : parameter 'identity' of type 'SocketConnectable' not supported

var tlsClientConnectionSetUseSsl3Function *gi.Function
var tlsClientConnectionSetUseSsl3Function_Once sync.Once

func tlsClientConnectionSetUseSsl3Function_Set() error {
	var err error
	tlsClientConnectionSetUseSsl3Function_Once.Do(func() {
		err = tlsClientConnectionInterface_Set()
		if err != nil {
			return
		}
		tlsClientConnectionSetUseSsl3Function, err = tlsClientConnectionInterface.InvokerNew("set_use_ssl3")
	})
	return err
}

// SetUseSsl3 is a representation of the C type g_tls_client_connection_set_use_ssl3.
func (recv *TlsClientConnection) SetUseSsl3(useSsl3 bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetBoolean(useSsl3)

	err := tlsClientConnectionSetUseSsl3Function_Set()
	if err == nil {
		tlsClientConnectionSetUseSsl3Function.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'g_tls_client_connection_set_validation_flags' : parameter 'flags' of type 'TlsCertificateFlags' not supported

var tlsFileDatabaseInterface *gi.Interface
var tlsFileDatabaseInterface_Once sync.Once

func tlsFileDatabaseInterface_Set() error {
	var err error
	tlsFileDatabaseInterface_Once.Do(func() {
		tlsFileDatabaseInterface, err = gi.InterfaceNew("Gio", "TlsFileDatabase")
	})
	return err
}

type TlsFileDatabase struct {
	native unsafe.Pointer
}

func TlsFileDatabaseNewFromNative(native unsafe.Pointer) *TlsFileDatabase {
	instance := &TlsFileDatabase{native: native}

	return instance
}

/*
CastToTlsFileDatabase down casts any arbitrary Object to TlsFileDatabase.
Exercise care, as this is a potentially dangerous function
if the Object is not a TlsFileDatabase.
*/
func CastToTlsFileDatabase(object *gobject.Object) *TlsFileDatabase {
	return TlsFileDatabaseNewFromNative(object.Native())
}

// Equals compares this TlsFileDatabase with another TlsFileDatabase, and returns true if they represent the same Object.
func (recv *TlsFileDatabase) Equals(other *TlsFileDatabase) bool {
	return other.Native() == recv.Native()
}

func (recv *TlsFileDatabase) Native() unsafe.Pointer {
	return recv.native
}

var tlsServerConnectionInterface *gi.Interface
var tlsServerConnectionInterface_Once sync.Once

func tlsServerConnectionInterface_Set() error {
	var err error
	tlsServerConnectionInterface_Once.Do(func() {
		tlsServerConnectionInterface, err = gi.InterfaceNew("Gio", "TlsServerConnection")
	})
	return err
}

type TlsServerConnection struct {
	native unsafe.Pointer
}

func TlsServerConnectionNewFromNative(native unsafe.Pointer) *TlsServerConnection {
	instance := &TlsServerConnection{native: native}

	return instance
}

/*
CastToTlsServerConnection down casts any arbitrary Object to TlsServerConnection.
Exercise care, as this is a potentially dangerous function
if the Object is not a TlsServerConnection.
*/
func CastToTlsServerConnection(object *gobject.Object) *TlsServerConnection {
	return TlsServerConnectionNewFromNative(object.Native())
}

// Equals compares this TlsServerConnection with another TlsServerConnection, and returns true if they represent the same Object.
func (recv *TlsServerConnection) Equals(other *TlsServerConnection) bool {
	return other.Native() == recv.Native()
}

func (recv *TlsServerConnection) Native() unsafe.Pointer {
	return recv.native
}

var volumeInterface *gi.Interface
var volumeInterface_Once sync.Once

func volumeInterface_Set() error {
	var err error
	volumeInterface_Once.Do(func() {
		volumeInterface, err = gi.InterfaceNew("Gio", "Volume")
	})
	return err
}

type Volume struct {
	native unsafe.Pointer
}

func VolumeNewFromNative(native unsafe.Pointer) *Volume {
	instance := &Volume{native: native}

	return instance
}

/*
CastToVolume down casts any arbitrary Object to Volume.
Exercise care, as this is a potentially dangerous function
if the Object is not a Volume.
*/
func CastToVolume(object *gobject.Object) *Volume {
	return VolumeNewFromNative(object.Native())
}

// Equals compares this Volume with another Volume, and returns true if they represent the same Object.
func (recv *Volume) Equals(other *Volume) bool {
	return other.Native() == recv.Native()
}

func (recv *Volume) Native() unsafe.Pointer {
	return recv.native
}

var volumeCanEjectFunction *gi.Function
var volumeCanEjectFunction_Once sync.Once

func volumeCanEjectFunction_Set() error {
	var err error
	volumeCanEjectFunction_Once.Do(func() {
		err = volumeInterface_Set()
		if err != nil {
			return
		}
		volumeCanEjectFunction, err = volumeInterface.InvokerNew("can_eject")
	})
	return err
}

// CanEject is a representation of the C type g_volume_can_eject.
func (recv *Volume) CanEject() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := volumeCanEjectFunction_Set()
	if err == nil {
		ret = volumeCanEjectFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var volumeCanMountFunction *gi.Function
var volumeCanMountFunction_Once sync.Once

func volumeCanMountFunction_Set() error {
	var err error
	volumeCanMountFunction_Once.Do(func() {
		err = volumeInterface_Set()
		if err != nil {
			return
		}
		volumeCanMountFunction, err = volumeInterface.InvokerNew("can_mount")
	})
	return err
}

// CanMount is a representation of the C type g_volume_can_mount.
func (recv *Volume) CanMount() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := volumeCanMountFunction_Set()
	if err == nil {
		ret = volumeCanMountFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'g_volume_eject' : parameter 'flags' of type 'MountUnmountFlags' not supported

// UNSUPPORTED : C value 'g_volume_eject_finish' : parameter 'result' of type 'AsyncResult' not supported

// UNSUPPORTED : C value 'g_volume_eject_with_operation' : parameter 'flags' of type 'MountUnmountFlags' not supported

// UNSUPPORTED : C value 'g_volume_eject_with_operation_finish' : parameter 'result' of type 'AsyncResult' not supported

var volumeEnumerateIdentifiersFunction *gi.Function
var volumeEnumerateIdentifiersFunction_Once sync.Once

func volumeEnumerateIdentifiersFunction_Set() error {
	var err error
	volumeEnumerateIdentifiersFunction_Once.Do(func() {
		err = volumeInterface_Set()
		if err != nil {
			return
		}
		volumeEnumerateIdentifiersFunction, err = volumeInterface.InvokerNew("enumerate_identifiers")
	})
	return err
}

// EnumerateIdentifiers is a representation of the C type g_volume_enumerate_identifiers.
func (recv *Volume) EnumerateIdentifiers() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := volumeEnumerateIdentifiersFunction_Set()
	if err == nil {
		volumeEnumerateIdentifiersFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'g_volume_get_activation_root' : return type 'File' not supported

// UNSUPPORTED : C value 'g_volume_get_drive' : return type 'Drive' not supported

// UNSUPPORTED : C value 'g_volume_get_icon' : return type 'Icon' not supported

var volumeGetIdentifierFunction *gi.Function
var volumeGetIdentifierFunction_Once sync.Once

func volumeGetIdentifierFunction_Set() error {
	var err error
	volumeGetIdentifierFunction_Once.Do(func() {
		err = volumeInterface_Set()
		if err != nil {
			return
		}
		volumeGetIdentifierFunction, err = volumeInterface.InvokerNew("get_identifier")
	})
	return err
}

// GetIdentifier is a representation of the C type g_volume_get_identifier.
func (recv *Volume) GetIdentifier(kind string) string {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(kind)

	var ret gi.Argument

	err := volumeGetIdentifierFunction_Set()
	if err == nil {
		ret = volumeGetIdentifierFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

// UNSUPPORTED : C value 'g_volume_get_mount' : return type 'Mount' not supported

var volumeGetNameFunction *gi.Function
var volumeGetNameFunction_Once sync.Once

func volumeGetNameFunction_Set() error {
	var err error
	volumeGetNameFunction_Once.Do(func() {
		err = volumeInterface_Set()
		if err != nil {
			return
		}
		volumeGetNameFunction, err = volumeInterface.InvokerNew("get_name")
	})
	return err
}

// GetName is a representation of the C type g_volume_get_name.
func (recv *Volume) GetName() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := volumeGetNameFunction_Set()
	if err == nil {
		ret = volumeGetNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

var volumeGetSortKeyFunction *gi.Function
var volumeGetSortKeyFunction_Once sync.Once

func volumeGetSortKeyFunction_Set() error {
	var err error
	volumeGetSortKeyFunction_Once.Do(func() {
		err = volumeInterface_Set()
		if err != nil {
			return
		}
		volumeGetSortKeyFunction, err = volumeInterface.InvokerNew("get_sort_key")
	})
	return err
}

// GetSortKey is a representation of the C type g_volume_get_sort_key.
func (recv *Volume) GetSortKey() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := volumeGetSortKeyFunction_Set()
	if err == nil {
		ret = volumeGetSortKeyFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'g_volume_get_symbolic_icon' : return type 'Icon' not supported

var volumeGetUuidFunction *gi.Function
var volumeGetUuidFunction_Once sync.Once

func volumeGetUuidFunction_Set() error {
	var err error
	volumeGetUuidFunction_Once.Do(func() {
		err = volumeInterface_Set()
		if err != nil {
			return
		}
		volumeGetUuidFunction, err = volumeInterface.InvokerNew("get_uuid")
	})
	return err
}

// GetUuid is a representation of the C type g_volume_get_uuid.
func (recv *Volume) GetUuid() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := volumeGetUuidFunction_Set()
	if err == nil {
		ret = volumeGetUuidFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

// UNSUPPORTED : C value 'g_volume_mount' : parameter 'flags' of type 'MountMountFlags' not supported

// UNSUPPORTED : C value 'g_volume_mount_finish' : parameter 'result' of type 'AsyncResult' not supported

var volumeShouldAutomountFunction *gi.Function
var volumeShouldAutomountFunction_Once sync.Once

func volumeShouldAutomountFunction_Set() error {
	var err error
	volumeShouldAutomountFunction_Once.Do(func() {
		err = volumeInterface_Set()
		if err != nil {
			return
		}
		volumeShouldAutomountFunction, err = volumeInterface.InvokerNew("should_automount")
	})
	return err
}

// ShouldAutomount is a representation of the C type g_volume_should_automount.
func (recv *Volume) ShouldAutomount() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := volumeShouldAutomountFunction_Set()
	if err == nil {
		ret = volumeShouldAutomountFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

/*
ConnectChanged connects a callback to the 'changed' signal of the Volume.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *Volume) ConnectChanged(handler func(instance *Volume)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {}

	return callback.ConnectSignal(recv.Native(), "changed", marshal)
}

/*
ConnectRemoved connects a callback to the 'removed' signal of the Volume.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *Volume) ConnectRemoved(handler func(instance *Volume)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {}

	return callback.ConnectSignal(recv.Native(), "removed", marshal)
}

/*
Disconnect disconnects a callback previously registered with a Connect...() method.

The connectionID should be a value returned from a call to a Connect...() method.
*/
func (recv *Volume) DisconnectSignal(connectionID int) {
	callback.DisconnectSignal(connectionID)
}
