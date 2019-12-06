// Code generated - DO NOT EDIT.

package gio

import (
	gi "github.com/pekim/gobbi/internal/gi"
	gobject "github.com/pekim/gobbi/lib/gobject"
	"runtime"
	"sync"
	"unsafe"
)

var actionEntryStruct *gi.Struct
var actionEntryStruct_Once sync.Once

func actionEntryStruct_Set() error {
	var err error
	actionEntryStruct_Once.Do(func() {
		actionEntryStruct, err = gi.StructNew("Gio", "ActionEntry")
	})
	return err
}

type ActionEntry struct {
	native unsafe.Pointer
}

func ActionEntryNewFromNative(native unsafe.Pointer) *ActionEntry {
	instance := &ActionEntry{native: native}

	return instance
}

/*
CastToActionEntry down casts any arbitrary Object to ActionEntry.
Exercise care, as this is a potentially dangerous function
if the Object is not a ActionEntry.
*/
func (recv *ActionEntry) CastToActionEntry(object *gobject.Object) *ActionEntry {
	return ActionEntryNewFromNative(object.Native())
}

func (recv *ActionEntry) Native() unsafe.Pointer {
	return recv.native
}

// FieldName returns the C field 'name'.
func (recv *ActionEntry) FieldName() string {
	argValue := gi.StructFieldGet(actionEntryStruct, recv.native, "name")
	value := argValue.String(false)
	return value
}

// SetFieldName sets the value of the C field 'name'.
func (recv *ActionEntry) SetFieldName(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.StructFieldSet(actionEntryStruct, recv.native, "name", argValue)
}

// UNSUPPORTED : C value 'activate' : for field getter : missing Type

// UNSUPPORTED : C value 'activate' : for field setter : missing Type

// FieldParameterType returns the C field 'parameter_type'.
func (recv *ActionEntry) FieldParameterType() string {
	argValue := gi.StructFieldGet(actionEntryStruct, recv.native, "parameter_type")
	value := argValue.String(false)
	return value
}

// SetFieldParameterType sets the value of the C field 'parameter_type'.
func (recv *ActionEntry) SetFieldParameterType(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.StructFieldSet(actionEntryStruct, recv.native, "parameter_type", argValue)
}

// FieldState returns the C field 'state'.
func (recv *ActionEntry) FieldState() string {
	argValue := gi.StructFieldGet(actionEntryStruct, recv.native, "state")
	value := argValue.String(false)
	return value
}

// SetFieldState sets the value of the C field 'state'.
func (recv *ActionEntry) SetFieldState(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.StructFieldSet(actionEntryStruct, recv.native, "state", argValue)
}

// UNSUPPORTED : C value 'change_state' : for field getter : missing Type

// UNSUPPORTED : C value 'change_state' : for field setter : missing Type

// ActionEntryStruct creates an uninitialised ActionEntry.
func ActionEntryStruct() *ActionEntry {
	err := actionEntryStruct_Set()
	if err != nil {
		return nil
	}

	structGo := ActionEntryNewFromNative(actionEntryStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeActionEntry)
	return structGo
}
func finalizeActionEntry(obj *ActionEntry) {
	actionEntryStruct.Free(obj.native)
}

var actionGroupInterfaceStruct *gi.Struct
var actionGroupInterfaceStruct_Once sync.Once

func actionGroupInterfaceStruct_Set() error {
	var err error
	actionGroupInterfaceStruct_Once.Do(func() {
		actionGroupInterfaceStruct, err = gi.StructNew("Gio", "ActionGroupInterface")
	})
	return err
}

type ActionGroupInterface struct {
	native unsafe.Pointer
}

func ActionGroupInterfaceNewFromNative(native unsafe.Pointer) *ActionGroupInterface {
	instance := &ActionGroupInterface{native: native}

	return instance
}

/*
CastToActionGroupInterface down casts any arbitrary Object to ActionGroupInterface.
Exercise care, as this is a potentially dangerous function
if the Object is not a ActionGroupInterface.
*/
func (recv *ActionGroupInterface) CastToActionGroupInterface(object *gobject.Object) *ActionGroupInterface {
	return ActionGroupInterfaceNewFromNative(object.Native())
}

func (recv *ActionGroupInterface) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'g_iface' : for field getter : no Go type for 'GObject.TypeInterface'

// UNSUPPORTED : C value 'g_iface' : for field setter : no Go type for 'GObject.TypeInterface'

// UNSUPPORTED : C value 'has_action' : for field getter : missing Type

// UNSUPPORTED : C value 'has_action' : for field setter : missing Type

// UNSUPPORTED : C value 'list_actions' : for field getter : missing Type

// UNSUPPORTED : C value 'list_actions' : for field setter : missing Type

// UNSUPPORTED : C value 'get_action_enabled' : for field getter : missing Type

// UNSUPPORTED : C value 'get_action_enabled' : for field setter : missing Type

// UNSUPPORTED : C value 'get_action_parameter_type' : for field getter : missing Type

// UNSUPPORTED : C value 'get_action_parameter_type' : for field setter : missing Type

// UNSUPPORTED : C value 'get_action_state_type' : for field getter : missing Type

// UNSUPPORTED : C value 'get_action_state_type' : for field setter : missing Type

// UNSUPPORTED : C value 'get_action_state_hint' : for field getter : missing Type

// UNSUPPORTED : C value 'get_action_state_hint' : for field setter : missing Type

// UNSUPPORTED : C value 'get_action_state' : for field getter : missing Type

// UNSUPPORTED : C value 'get_action_state' : for field setter : missing Type

// UNSUPPORTED : C value 'change_action_state' : for field getter : missing Type

// UNSUPPORTED : C value 'change_action_state' : for field setter : missing Type

// UNSUPPORTED : C value 'activate_action' : for field getter : missing Type

// UNSUPPORTED : C value 'activate_action' : for field setter : missing Type

// UNSUPPORTED : C value 'action_added' : for field getter : missing Type

// UNSUPPORTED : C value 'action_added' : for field setter : missing Type

// UNSUPPORTED : C value 'action_removed' : for field getter : missing Type

// UNSUPPORTED : C value 'action_removed' : for field setter : missing Type

// UNSUPPORTED : C value 'action_enabled_changed' : for field getter : missing Type

// UNSUPPORTED : C value 'action_enabled_changed' : for field setter : missing Type

// UNSUPPORTED : C value 'action_state_changed' : for field getter : missing Type

// UNSUPPORTED : C value 'action_state_changed' : for field setter : missing Type

// UNSUPPORTED : C value 'query_action' : for field getter : missing Type

// UNSUPPORTED : C value 'query_action' : for field setter : missing Type

// ActionGroupInterfaceStruct creates an uninitialised ActionGroupInterface.
func ActionGroupInterfaceStruct() *ActionGroupInterface {
	err := actionGroupInterfaceStruct_Set()
	if err != nil {
		return nil
	}

	structGo := ActionGroupInterfaceNewFromNative(actionGroupInterfaceStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeActionGroupInterface)
	return structGo
}
func finalizeActionGroupInterface(obj *ActionGroupInterface) {
	actionGroupInterfaceStruct.Free(obj.native)
}

var actionInterfaceStruct *gi.Struct
var actionInterfaceStruct_Once sync.Once

func actionInterfaceStruct_Set() error {
	var err error
	actionInterfaceStruct_Once.Do(func() {
		actionInterfaceStruct, err = gi.StructNew("Gio", "ActionInterface")
	})
	return err
}

type ActionInterface struct {
	native unsafe.Pointer
}

func ActionInterfaceNewFromNative(native unsafe.Pointer) *ActionInterface {
	instance := &ActionInterface{native: native}

	return instance
}

/*
CastToActionInterface down casts any arbitrary Object to ActionInterface.
Exercise care, as this is a potentially dangerous function
if the Object is not a ActionInterface.
*/
func (recv *ActionInterface) CastToActionInterface(object *gobject.Object) *ActionInterface {
	return ActionInterfaceNewFromNative(object.Native())
}

func (recv *ActionInterface) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'g_iface' : for field getter : no Go type for 'GObject.TypeInterface'

// UNSUPPORTED : C value 'g_iface' : for field setter : no Go type for 'GObject.TypeInterface'

// UNSUPPORTED : C value 'get_name' : for field getter : missing Type

// UNSUPPORTED : C value 'get_name' : for field setter : missing Type

// UNSUPPORTED : C value 'get_parameter_type' : for field getter : missing Type

// UNSUPPORTED : C value 'get_parameter_type' : for field setter : missing Type

// UNSUPPORTED : C value 'get_state_type' : for field getter : missing Type

// UNSUPPORTED : C value 'get_state_type' : for field setter : missing Type

// UNSUPPORTED : C value 'get_state_hint' : for field getter : missing Type

// UNSUPPORTED : C value 'get_state_hint' : for field setter : missing Type

// UNSUPPORTED : C value 'get_enabled' : for field getter : missing Type

// UNSUPPORTED : C value 'get_enabled' : for field setter : missing Type

// UNSUPPORTED : C value 'get_state' : for field getter : missing Type

// UNSUPPORTED : C value 'get_state' : for field setter : missing Type

// UNSUPPORTED : C value 'change_state' : for field getter : missing Type

// UNSUPPORTED : C value 'change_state' : for field setter : missing Type

// UNSUPPORTED : C value 'activate' : for field getter : missing Type

// UNSUPPORTED : C value 'activate' : for field setter : missing Type

// ActionInterfaceStruct creates an uninitialised ActionInterface.
func ActionInterfaceStruct() *ActionInterface {
	err := actionInterfaceStruct_Set()
	if err != nil {
		return nil
	}

	structGo := ActionInterfaceNewFromNative(actionInterfaceStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeActionInterface)
	return structGo
}
func finalizeActionInterface(obj *ActionInterface) {
	actionInterfaceStruct.Free(obj.native)
}

var actionMapInterfaceStruct *gi.Struct
var actionMapInterfaceStruct_Once sync.Once

func actionMapInterfaceStruct_Set() error {
	var err error
	actionMapInterfaceStruct_Once.Do(func() {
		actionMapInterfaceStruct, err = gi.StructNew("Gio", "ActionMapInterface")
	})
	return err
}

type ActionMapInterface struct {
	native unsafe.Pointer
}

func ActionMapInterfaceNewFromNative(native unsafe.Pointer) *ActionMapInterface {
	instance := &ActionMapInterface{native: native}

	return instance
}

/*
CastToActionMapInterface down casts any arbitrary Object to ActionMapInterface.
Exercise care, as this is a potentially dangerous function
if the Object is not a ActionMapInterface.
*/
func (recv *ActionMapInterface) CastToActionMapInterface(object *gobject.Object) *ActionMapInterface {
	return ActionMapInterfaceNewFromNative(object.Native())
}

func (recv *ActionMapInterface) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'g_iface' : for field getter : no Go type for 'GObject.TypeInterface'

// UNSUPPORTED : C value 'g_iface' : for field setter : no Go type for 'GObject.TypeInterface'

// UNSUPPORTED : C value 'lookup_action' : for field getter : missing Type

// UNSUPPORTED : C value 'lookup_action' : for field setter : missing Type

// UNSUPPORTED : C value 'add_action' : for field getter : missing Type

// UNSUPPORTED : C value 'add_action' : for field setter : missing Type

// UNSUPPORTED : C value 'remove_action' : for field getter : missing Type

// UNSUPPORTED : C value 'remove_action' : for field setter : missing Type

// ActionMapInterfaceStruct creates an uninitialised ActionMapInterface.
func ActionMapInterfaceStruct() *ActionMapInterface {
	err := actionMapInterfaceStruct_Set()
	if err != nil {
		return nil
	}

	structGo := ActionMapInterfaceNewFromNative(actionMapInterfaceStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeActionMapInterface)
	return structGo
}
func finalizeActionMapInterface(obj *ActionMapInterface) {
	actionMapInterfaceStruct.Free(obj.native)
}

var appInfoIfaceStruct *gi.Struct
var appInfoIfaceStruct_Once sync.Once

func appInfoIfaceStruct_Set() error {
	var err error
	appInfoIfaceStruct_Once.Do(func() {
		appInfoIfaceStruct, err = gi.StructNew("Gio", "AppInfoIface")
	})
	return err
}

type AppInfoIface struct {
	native unsafe.Pointer
}

func AppInfoIfaceNewFromNative(native unsafe.Pointer) *AppInfoIface {
	instance := &AppInfoIface{native: native}

	return instance
}

/*
CastToAppInfoIface down casts any arbitrary Object to AppInfoIface.
Exercise care, as this is a potentially dangerous function
if the Object is not a AppInfoIface.
*/
func (recv *AppInfoIface) CastToAppInfoIface(object *gobject.Object) *AppInfoIface {
	return AppInfoIfaceNewFromNative(object.Native())
}

func (recv *AppInfoIface) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'g_iface' : for field getter : no Go type for 'GObject.TypeInterface'

// UNSUPPORTED : C value 'g_iface' : for field setter : no Go type for 'GObject.TypeInterface'

// UNSUPPORTED : C value 'dup' : for field getter : missing Type

// UNSUPPORTED : C value 'dup' : for field setter : missing Type

// UNSUPPORTED : C value 'equal' : for field getter : missing Type

// UNSUPPORTED : C value 'equal' : for field setter : missing Type

// UNSUPPORTED : C value 'get_id' : for field getter : missing Type

// UNSUPPORTED : C value 'get_id' : for field setter : missing Type

// UNSUPPORTED : C value 'get_name' : for field getter : missing Type

// UNSUPPORTED : C value 'get_name' : for field setter : missing Type

// UNSUPPORTED : C value 'get_description' : for field getter : missing Type

// UNSUPPORTED : C value 'get_description' : for field setter : missing Type

// UNSUPPORTED : C value 'get_executable' : for field getter : missing Type

// UNSUPPORTED : C value 'get_executable' : for field setter : missing Type

// UNSUPPORTED : C value 'get_icon' : for field getter : missing Type

// UNSUPPORTED : C value 'get_icon' : for field setter : missing Type

// UNSUPPORTED : C value 'launch' : for field getter : missing Type

// UNSUPPORTED : C value 'launch' : for field setter : missing Type

// UNSUPPORTED : C value 'supports_uris' : for field getter : missing Type

// UNSUPPORTED : C value 'supports_uris' : for field setter : missing Type

// UNSUPPORTED : C value 'supports_files' : for field getter : missing Type

// UNSUPPORTED : C value 'supports_files' : for field setter : missing Type

// UNSUPPORTED : C value 'launch_uris' : for field getter : missing Type

// UNSUPPORTED : C value 'launch_uris' : for field setter : missing Type

// UNSUPPORTED : C value 'should_show' : for field getter : missing Type

// UNSUPPORTED : C value 'should_show' : for field setter : missing Type

// UNSUPPORTED : C value 'set_as_default_for_type' : for field getter : missing Type

// UNSUPPORTED : C value 'set_as_default_for_type' : for field setter : missing Type

// UNSUPPORTED : C value 'set_as_default_for_extension' : for field getter : missing Type

// UNSUPPORTED : C value 'set_as_default_for_extension' : for field setter : missing Type

// UNSUPPORTED : C value 'add_supports_type' : for field getter : missing Type

// UNSUPPORTED : C value 'add_supports_type' : for field setter : missing Type

// UNSUPPORTED : C value 'can_remove_supports_type' : for field getter : missing Type

// UNSUPPORTED : C value 'can_remove_supports_type' : for field setter : missing Type

// UNSUPPORTED : C value 'remove_supports_type' : for field getter : missing Type

// UNSUPPORTED : C value 'remove_supports_type' : for field setter : missing Type

// UNSUPPORTED : C value 'can_delete' : for field getter : missing Type

// UNSUPPORTED : C value 'can_delete' : for field setter : missing Type

// UNSUPPORTED : C value 'do_delete' : for field getter : missing Type

// UNSUPPORTED : C value 'do_delete' : for field setter : missing Type

// UNSUPPORTED : C value 'get_commandline' : for field getter : missing Type

// UNSUPPORTED : C value 'get_commandline' : for field setter : missing Type

// UNSUPPORTED : C value 'get_display_name' : for field getter : missing Type

// UNSUPPORTED : C value 'get_display_name' : for field setter : missing Type

// UNSUPPORTED : C value 'set_as_last_used_for_type' : for field getter : missing Type

// UNSUPPORTED : C value 'set_as_last_used_for_type' : for field setter : missing Type

// UNSUPPORTED : C value 'get_supported_types' : for field getter : missing Type

// UNSUPPORTED : C value 'get_supported_types' : for field setter : missing Type

// UNSUPPORTED : C value 'launch_uris_async' : for field getter : missing Type

// UNSUPPORTED : C value 'launch_uris_async' : for field setter : missing Type

// UNSUPPORTED : C value 'launch_uris_finish' : for field getter : missing Type

// UNSUPPORTED : C value 'launch_uris_finish' : for field setter : missing Type

// AppInfoIfaceStruct creates an uninitialised AppInfoIface.
func AppInfoIfaceStruct() *AppInfoIface {
	err := appInfoIfaceStruct_Set()
	if err != nil {
		return nil
	}

	structGo := AppInfoIfaceNewFromNative(appInfoIfaceStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeAppInfoIface)
	return structGo
}
func finalizeAppInfoIface(obj *AppInfoIface) {
	appInfoIfaceStruct.Free(obj.native)
}

var appLaunchContextClassStruct *gi.Struct
var appLaunchContextClassStruct_Once sync.Once

func appLaunchContextClassStruct_Set() error {
	var err error
	appLaunchContextClassStruct_Once.Do(func() {
		appLaunchContextClassStruct, err = gi.StructNew("Gio", "AppLaunchContextClass")
	})
	return err
}

type AppLaunchContextClass struct {
	native unsafe.Pointer
}

func AppLaunchContextClassNewFromNative(native unsafe.Pointer) *AppLaunchContextClass {
	instance := &AppLaunchContextClass{native: native}

	return instance
}

/*
CastToAppLaunchContextClass down casts any arbitrary Object to AppLaunchContextClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a AppLaunchContextClass.
*/
func (recv *AppLaunchContextClass) CastToAppLaunchContextClass(object *gobject.Object) *AppLaunchContextClass {
	return AppLaunchContextClassNewFromNative(object.Native())
}

func (recv *AppLaunchContextClass) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'get_display' : for field getter : missing Type

// UNSUPPORTED : C value 'get_display' : for field setter : missing Type

// UNSUPPORTED : C value 'get_startup_notify_id' : for field getter : missing Type

// UNSUPPORTED : C value 'get_startup_notify_id' : for field setter : missing Type

// UNSUPPORTED : C value 'launch_failed' : for field getter : missing Type

// UNSUPPORTED : C value 'launch_failed' : for field setter : missing Type

// UNSUPPORTED : C value 'launched' : for field getter : missing Type

// UNSUPPORTED : C value 'launched' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved4' : for field setter : missing Type

// AppLaunchContextClassStruct creates an uninitialised AppLaunchContextClass.
func AppLaunchContextClassStruct() *AppLaunchContextClass {
	err := appLaunchContextClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := AppLaunchContextClassNewFromNative(appLaunchContextClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeAppLaunchContextClass)
	return structGo
}
func finalizeAppLaunchContextClass(obj *AppLaunchContextClass) {
	appLaunchContextClassStruct.Free(obj.native)
}

var appLaunchContextPrivateStruct *gi.Struct
var appLaunchContextPrivateStruct_Once sync.Once

func appLaunchContextPrivateStruct_Set() error {
	var err error
	appLaunchContextPrivateStruct_Once.Do(func() {
		appLaunchContextPrivateStruct, err = gi.StructNew("Gio", "AppLaunchContextPrivate")
	})
	return err
}

type AppLaunchContextPrivate struct {
	native unsafe.Pointer
}

func AppLaunchContextPrivateNewFromNative(native unsafe.Pointer) *AppLaunchContextPrivate {
	instance := &AppLaunchContextPrivate{native: native}

	return instance
}

/*
CastToAppLaunchContextPrivate down casts any arbitrary Object to AppLaunchContextPrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a AppLaunchContextPrivate.
*/
func (recv *AppLaunchContextPrivate) CastToAppLaunchContextPrivate(object *gobject.Object) *AppLaunchContextPrivate {
	return AppLaunchContextPrivateNewFromNative(object.Native())
}

func (recv *AppLaunchContextPrivate) Native() unsafe.Pointer {
	return recv.native
}

// AppLaunchContextPrivateStruct creates an uninitialised AppLaunchContextPrivate.
func AppLaunchContextPrivateStruct() *AppLaunchContextPrivate {
	err := appLaunchContextPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := AppLaunchContextPrivateNewFromNative(appLaunchContextPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeAppLaunchContextPrivate)
	return structGo
}
func finalizeAppLaunchContextPrivate(obj *AppLaunchContextPrivate) {
	appLaunchContextPrivateStruct.Free(obj.native)
}

var applicationClassStruct *gi.Struct
var applicationClassStruct_Once sync.Once

func applicationClassStruct_Set() error {
	var err error
	applicationClassStruct_Once.Do(func() {
		applicationClassStruct, err = gi.StructNew("Gio", "ApplicationClass")
	})
	return err
}

type ApplicationClass struct {
	native unsafe.Pointer
}

func ApplicationClassNewFromNative(native unsafe.Pointer) *ApplicationClass {
	instance := &ApplicationClass{native: native}

	return instance
}

/*
CastToApplicationClass down casts any arbitrary Object to ApplicationClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a ApplicationClass.
*/
func (recv *ApplicationClass) CastToApplicationClass(object *gobject.Object) *ApplicationClass {
	return ApplicationClassNewFromNative(object.Native())
}

func (recv *ApplicationClass) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'startup' : for field getter : missing Type

// UNSUPPORTED : C value 'startup' : for field setter : missing Type

// UNSUPPORTED : C value 'activate' : for field getter : missing Type

// UNSUPPORTED : C value 'activate' : for field setter : missing Type

// UNSUPPORTED : C value 'open' : for field getter : missing Type

// UNSUPPORTED : C value 'open' : for field setter : missing Type

// UNSUPPORTED : C value 'command_line' : for field getter : missing Type

// UNSUPPORTED : C value 'command_line' : for field setter : missing Type

// UNSUPPORTED : C value 'local_command_line' : for field getter : missing Type

// UNSUPPORTED : C value 'local_command_line' : for field setter : missing Type

// UNSUPPORTED : C value 'before_emit' : for field getter : missing Type

// UNSUPPORTED : C value 'before_emit' : for field setter : missing Type

// UNSUPPORTED : C value 'after_emit' : for field getter : missing Type

// UNSUPPORTED : C value 'after_emit' : for field setter : missing Type

// UNSUPPORTED : C value 'add_platform_data' : for field getter : missing Type

// UNSUPPORTED : C value 'add_platform_data' : for field setter : missing Type

// UNSUPPORTED : C value 'quit_mainloop' : for field getter : missing Type

// UNSUPPORTED : C value 'quit_mainloop' : for field setter : missing Type

// UNSUPPORTED : C value 'run_mainloop' : for field getter : missing Type

// UNSUPPORTED : C value 'run_mainloop' : for field setter : missing Type

// UNSUPPORTED : C value 'shutdown' : for field getter : missing Type

// UNSUPPORTED : C value 'shutdown' : for field setter : missing Type

// UNSUPPORTED : C value 'dbus_register' : for field getter : missing Type

// UNSUPPORTED : C value 'dbus_register' : for field setter : missing Type

// UNSUPPORTED : C value 'dbus_unregister' : for field getter : missing Type

// UNSUPPORTED : C value 'dbus_unregister' : for field setter : missing Type

// UNSUPPORTED : C value 'handle_local_options' : for field getter : missing Type

// UNSUPPORTED : C value 'handle_local_options' : for field setter : missing Type

// UNSUPPORTED : C value 'name_lost' : for field getter : missing Type

// UNSUPPORTED : C value 'name_lost' : for field setter : missing Type

// ApplicationClassStruct creates an uninitialised ApplicationClass.
func ApplicationClassStruct() *ApplicationClass {
	err := applicationClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := ApplicationClassNewFromNative(applicationClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeApplicationClass)
	return structGo
}
func finalizeApplicationClass(obj *ApplicationClass) {
	applicationClassStruct.Free(obj.native)
}

var applicationCommandLineClassStruct *gi.Struct
var applicationCommandLineClassStruct_Once sync.Once

func applicationCommandLineClassStruct_Set() error {
	var err error
	applicationCommandLineClassStruct_Once.Do(func() {
		applicationCommandLineClassStruct, err = gi.StructNew("Gio", "ApplicationCommandLineClass")
	})
	return err
}

type ApplicationCommandLineClass struct {
	native unsafe.Pointer
}

func ApplicationCommandLineClassNewFromNative(native unsafe.Pointer) *ApplicationCommandLineClass {
	instance := &ApplicationCommandLineClass{native: native}

	return instance
}

/*
CastToApplicationCommandLineClass down casts any arbitrary Object to ApplicationCommandLineClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a ApplicationCommandLineClass.
*/
func (recv *ApplicationCommandLineClass) CastToApplicationCommandLineClass(object *gobject.Object) *ApplicationCommandLineClass {
	return ApplicationCommandLineClassNewFromNative(object.Native())
}

func (recv *ApplicationCommandLineClass) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'print_literal' : for field getter : missing Type

// UNSUPPORTED : C value 'print_literal' : for field setter : missing Type

// UNSUPPORTED : C value 'printerr_literal' : for field getter : missing Type

// UNSUPPORTED : C value 'printerr_literal' : for field setter : missing Type

// UNSUPPORTED : C value 'get_stdin' : for field getter : missing Type

// UNSUPPORTED : C value 'get_stdin' : for field setter : missing Type

// ApplicationCommandLineClassStruct creates an uninitialised ApplicationCommandLineClass.
func ApplicationCommandLineClassStruct() *ApplicationCommandLineClass {
	err := applicationCommandLineClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := ApplicationCommandLineClassNewFromNative(applicationCommandLineClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeApplicationCommandLineClass)
	return structGo
}
func finalizeApplicationCommandLineClass(obj *ApplicationCommandLineClass) {
	applicationCommandLineClassStruct.Free(obj.native)
}

var applicationCommandLinePrivateStruct *gi.Struct
var applicationCommandLinePrivateStruct_Once sync.Once

func applicationCommandLinePrivateStruct_Set() error {
	var err error
	applicationCommandLinePrivateStruct_Once.Do(func() {
		applicationCommandLinePrivateStruct, err = gi.StructNew("Gio", "ApplicationCommandLinePrivate")
	})
	return err
}

type ApplicationCommandLinePrivate struct {
	native unsafe.Pointer
}

func ApplicationCommandLinePrivateNewFromNative(native unsafe.Pointer) *ApplicationCommandLinePrivate {
	instance := &ApplicationCommandLinePrivate{native: native}

	return instance
}

/*
CastToApplicationCommandLinePrivate down casts any arbitrary Object to ApplicationCommandLinePrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a ApplicationCommandLinePrivate.
*/
func (recv *ApplicationCommandLinePrivate) CastToApplicationCommandLinePrivate(object *gobject.Object) *ApplicationCommandLinePrivate {
	return ApplicationCommandLinePrivateNewFromNative(object.Native())
}

func (recv *ApplicationCommandLinePrivate) Native() unsafe.Pointer {
	return recv.native
}

// ApplicationCommandLinePrivateStruct creates an uninitialised ApplicationCommandLinePrivate.
func ApplicationCommandLinePrivateStruct() *ApplicationCommandLinePrivate {
	err := applicationCommandLinePrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := ApplicationCommandLinePrivateNewFromNative(applicationCommandLinePrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeApplicationCommandLinePrivate)
	return structGo
}
func finalizeApplicationCommandLinePrivate(obj *ApplicationCommandLinePrivate) {
	applicationCommandLinePrivateStruct.Free(obj.native)
}

var applicationPrivateStruct *gi.Struct
var applicationPrivateStruct_Once sync.Once

func applicationPrivateStruct_Set() error {
	var err error
	applicationPrivateStruct_Once.Do(func() {
		applicationPrivateStruct, err = gi.StructNew("Gio", "ApplicationPrivate")
	})
	return err
}

type ApplicationPrivate struct {
	native unsafe.Pointer
}

func ApplicationPrivateNewFromNative(native unsafe.Pointer) *ApplicationPrivate {
	instance := &ApplicationPrivate{native: native}

	return instance
}

/*
CastToApplicationPrivate down casts any arbitrary Object to ApplicationPrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a ApplicationPrivate.
*/
func (recv *ApplicationPrivate) CastToApplicationPrivate(object *gobject.Object) *ApplicationPrivate {
	return ApplicationPrivateNewFromNative(object.Native())
}

func (recv *ApplicationPrivate) Native() unsafe.Pointer {
	return recv.native
}

// ApplicationPrivateStruct creates an uninitialised ApplicationPrivate.
func ApplicationPrivateStruct() *ApplicationPrivate {
	err := applicationPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := ApplicationPrivateNewFromNative(applicationPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeApplicationPrivate)
	return structGo
}
func finalizeApplicationPrivate(obj *ApplicationPrivate) {
	applicationPrivateStruct.Free(obj.native)
}

var asyncInitableIfaceStruct *gi.Struct
var asyncInitableIfaceStruct_Once sync.Once

func asyncInitableIfaceStruct_Set() error {
	var err error
	asyncInitableIfaceStruct_Once.Do(func() {
		asyncInitableIfaceStruct, err = gi.StructNew("Gio", "AsyncInitableIface")
	})
	return err
}

type AsyncInitableIface struct {
	native unsafe.Pointer
}

func AsyncInitableIfaceNewFromNative(native unsafe.Pointer) *AsyncInitableIface {
	instance := &AsyncInitableIface{native: native}

	return instance
}

/*
CastToAsyncInitableIface down casts any arbitrary Object to AsyncInitableIface.
Exercise care, as this is a potentially dangerous function
if the Object is not a AsyncInitableIface.
*/
func (recv *AsyncInitableIface) CastToAsyncInitableIface(object *gobject.Object) *AsyncInitableIface {
	return AsyncInitableIfaceNewFromNative(object.Native())
}

func (recv *AsyncInitableIface) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'g_iface' : for field getter : no Go type for 'GObject.TypeInterface'

// UNSUPPORTED : C value 'g_iface' : for field setter : no Go type for 'GObject.TypeInterface'

// UNSUPPORTED : C value 'init_async' : for field getter : missing Type

// UNSUPPORTED : C value 'init_async' : for field setter : missing Type

// UNSUPPORTED : C value 'init_finish' : for field getter : missing Type

// UNSUPPORTED : C value 'init_finish' : for field setter : missing Type

// AsyncInitableIfaceStruct creates an uninitialised AsyncInitableIface.
func AsyncInitableIfaceStruct() *AsyncInitableIface {
	err := asyncInitableIfaceStruct_Set()
	if err != nil {
		return nil
	}

	structGo := AsyncInitableIfaceNewFromNative(asyncInitableIfaceStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeAsyncInitableIface)
	return structGo
}
func finalizeAsyncInitableIface(obj *AsyncInitableIface) {
	asyncInitableIfaceStruct.Free(obj.native)
}

var asyncResultIfaceStruct *gi.Struct
var asyncResultIfaceStruct_Once sync.Once

func asyncResultIfaceStruct_Set() error {
	var err error
	asyncResultIfaceStruct_Once.Do(func() {
		asyncResultIfaceStruct, err = gi.StructNew("Gio", "AsyncResultIface")
	})
	return err
}

type AsyncResultIface struct {
	native unsafe.Pointer
}

func AsyncResultIfaceNewFromNative(native unsafe.Pointer) *AsyncResultIface {
	instance := &AsyncResultIface{native: native}

	return instance
}

/*
CastToAsyncResultIface down casts any arbitrary Object to AsyncResultIface.
Exercise care, as this is a potentially dangerous function
if the Object is not a AsyncResultIface.
*/
func (recv *AsyncResultIface) CastToAsyncResultIface(object *gobject.Object) *AsyncResultIface {
	return AsyncResultIfaceNewFromNative(object.Native())
}

func (recv *AsyncResultIface) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'g_iface' : for field getter : no Go type for 'GObject.TypeInterface'

// UNSUPPORTED : C value 'g_iface' : for field setter : no Go type for 'GObject.TypeInterface'

// UNSUPPORTED : C value 'get_user_data' : for field getter : missing Type

// UNSUPPORTED : C value 'get_user_data' : for field setter : missing Type

// UNSUPPORTED : C value 'get_source_object' : for field getter : missing Type

// UNSUPPORTED : C value 'get_source_object' : for field setter : missing Type

// UNSUPPORTED : C value 'is_tagged' : for field getter : missing Type

// UNSUPPORTED : C value 'is_tagged' : for field setter : missing Type

// AsyncResultIfaceStruct creates an uninitialised AsyncResultIface.
func AsyncResultIfaceStruct() *AsyncResultIface {
	err := asyncResultIfaceStruct_Set()
	if err != nil {
		return nil
	}

	structGo := AsyncResultIfaceNewFromNative(asyncResultIfaceStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeAsyncResultIface)
	return structGo
}
func finalizeAsyncResultIface(obj *AsyncResultIface) {
	asyncResultIfaceStruct.Free(obj.native)
}

var bufferedInputStreamClassStruct *gi.Struct
var bufferedInputStreamClassStruct_Once sync.Once

func bufferedInputStreamClassStruct_Set() error {
	var err error
	bufferedInputStreamClassStruct_Once.Do(func() {
		bufferedInputStreamClassStruct, err = gi.StructNew("Gio", "BufferedInputStreamClass")
	})
	return err
}

type BufferedInputStreamClass struct {
	native unsafe.Pointer
}

func BufferedInputStreamClassNewFromNative(native unsafe.Pointer) *BufferedInputStreamClass {
	instance := &BufferedInputStreamClass{native: native}

	return instance
}

/*
CastToBufferedInputStreamClass down casts any arbitrary Object to BufferedInputStreamClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a BufferedInputStreamClass.
*/
func (recv *BufferedInputStreamClass) CastToBufferedInputStreamClass(object *gobject.Object) *BufferedInputStreamClass {
	return BufferedInputStreamClassNewFromNative(object.Native())
}

func (recv *BufferedInputStreamClass) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentClass returns the C field 'parent_class'.
func (recv *BufferedInputStreamClass) FieldParentClass() *FilterInputStreamClass {
	argValue := gi.StructFieldGet(bufferedInputStreamClassStruct, recv.native, "parent_class")
	value := FilterInputStreamClassNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *BufferedInputStreamClass) SetFieldParentClass(value *FilterInputStreamClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.StructFieldSet(bufferedInputStreamClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value 'fill' : for field getter : missing Type

// UNSUPPORTED : C value 'fill' : for field setter : missing Type

// UNSUPPORTED : C value 'fill_async' : for field getter : missing Type

// UNSUPPORTED : C value 'fill_async' : for field setter : missing Type

// UNSUPPORTED : C value 'fill_finish' : for field getter : missing Type

// UNSUPPORTED : C value 'fill_finish' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved4' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved5' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved5' : for field setter : missing Type

// BufferedInputStreamClassStruct creates an uninitialised BufferedInputStreamClass.
func BufferedInputStreamClassStruct() *BufferedInputStreamClass {
	err := bufferedInputStreamClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := BufferedInputStreamClassNewFromNative(bufferedInputStreamClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeBufferedInputStreamClass)
	return structGo
}
func finalizeBufferedInputStreamClass(obj *BufferedInputStreamClass) {
	bufferedInputStreamClassStruct.Free(obj.native)
}

var bufferedInputStreamPrivateStruct *gi.Struct
var bufferedInputStreamPrivateStruct_Once sync.Once

func bufferedInputStreamPrivateStruct_Set() error {
	var err error
	bufferedInputStreamPrivateStruct_Once.Do(func() {
		bufferedInputStreamPrivateStruct, err = gi.StructNew("Gio", "BufferedInputStreamPrivate")
	})
	return err
}

type BufferedInputStreamPrivate struct {
	native unsafe.Pointer
}

func BufferedInputStreamPrivateNewFromNative(native unsafe.Pointer) *BufferedInputStreamPrivate {
	instance := &BufferedInputStreamPrivate{native: native}

	return instance
}

/*
CastToBufferedInputStreamPrivate down casts any arbitrary Object to BufferedInputStreamPrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a BufferedInputStreamPrivate.
*/
func (recv *BufferedInputStreamPrivate) CastToBufferedInputStreamPrivate(object *gobject.Object) *BufferedInputStreamPrivate {
	return BufferedInputStreamPrivateNewFromNative(object.Native())
}

func (recv *BufferedInputStreamPrivate) Native() unsafe.Pointer {
	return recv.native
}

// BufferedInputStreamPrivateStruct creates an uninitialised BufferedInputStreamPrivate.
func BufferedInputStreamPrivateStruct() *BufferedInputStreamPrivate {
	err := bufferedInputStreamPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := BufferedInputStreamPrivateNewFromNative(bufferedInputStreamPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeBufferedInputStreamPrivate)
	return structGo
}
func finalizeBufferedInputStreamPrivate(obj *BufferedInputStreamPrivate) {
	bufferedInputStreamPrivateStruct.Free(obj.native)
}

var bufferedOutputStreamClassStruct *gi.Struct
var bufferedOutputStreamClassStruct_Once sync.Once

func bufferedOutputStreamClassStruct_Set() error {
	var err error
	bufferedOutputStreamClassStruct_Once.Do(func() {
		bufferedOutputStreamClassStruct, err = gi.StructNew("Gio", "BufferedOutputStreamClass")
	})
	return err
}

type BufferedOutputStreamClass struct {
	native unsafe.Pointer
}

func BufferedOutputStreamClassNewFromNative(native unsafe.Pointer) *BufferedOutputStreamClass {
	instance := &BufferedOutputStreamClass{native: native}

	return instance
}

/*
CastToBufferedOutputStreamClass down casts any arbitrary Object to BufferedOutputStreamClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a BufferedOutputStreamClass.
*/
func (recv *BufferedOutputStreamClass) CastToBufferedOutputStreamClass(object *gobject.Object) *BufferedOutputStreamClass {
	return BufferedOutputStreamClassNewFromNative(object.Native())
}

func (recv *BufferedOutputStreamClass) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentClass returns the C field 'parent_class'.
func (recv *BufferedOutputStreamClass) FieldParentClass() *FilterOutputStreamClass {
	argValue := gi.StructFieldGet(bufferedOutputStreamClassStruct, recv.native, "parent_class")
	value := FilterOutputStreamClassNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *BufferedOutputStreamClass) SetFieldParentClass(value *FilterOutputStreamClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.StructFieldSet(bufferedOutputStreamClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value '_g_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved2' : for field setter : missing Type

// BufferedOutputStreamClassStruct creates an uninitialised BufferedOutputStreamClass.
func BufferedOutputStreamClassStruct() *BufferedOutputStreamClass {
	err := bufferedOutputStreamClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := BufferedOutputStreamClassNewFromNative(bufferedOutputStreamClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeBufferedOutputStreamClass)
	return structGo
}
func finalizeBufferedOutputStreamClass(obj *BufferedOutputStreamClass) {
	bufferedOutputStreamClassStruct.Free(obj.native)
}

var bufferedOutputStreamPrivateStruct *gi.Struct
var bufferedOutputStreamPrivateStruct_Once sync.Once

func bufferedOutputStreamPrivateStruct_Set() error {
	var err error
	bufferedOutputStreamPrivateStruct_Once.Do(func() {
		bufferedOutputStreamPrivateStruct, err = gi.StructNew("Gio", "BufferedOutputStreamPrivate")
	})
	return err
}

type BufferedOutputStreamPrivate struct {
	native unsafe.Pointer
}

func BufferedOutputStreamPrivateNewFromNative(native unsafe.Pointer) *BufferedOutputStreamPrivate {
	instance := &BufferedOutputStreamPrivate{native: native}

	return instance
}

/*
CastToBufferedOutputStreamPrivate down casts any arbitrary Object to BufferedOutputStreamPrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a BufferedOutputStreamPrivate.
*/
func (recv *BufferedOutputStreamPrivate) CastToBufferedOutputStreamPrivate(object *gobject.Object) *BufferedOutputStreamPrivate {
	return BufferedOutputStreamPrivateNewFromNative(object.Native())
}

func (recv *BufferedOutputStreamPrivate) Native() unsafe.Pointer {
	return recv.native
}

// BufferedOutputStreamPrivateStruct creates an uninitialised BufferedOutputStreamPrivate.
func BufferedOutputStreamPrivateStruct() *BufferedOutputStreamPrivate {
	err := bufferedOutputStreamPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := BufferedOutputStreamPrivateNewFromNative(bufferedOutputStreamPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeBufferedOutputStreamPrivate)
	return structGo
}
func finalizeBufferedOutputStreamPrivate(obj *BufferedOutputStreamPrivate) {
	bufferedOutputStreamPrivateStruct.Free(obj.native)
}

var cancellableClassStruct *gi.Struct
var cancellableClassStruct_Once sync.Once

func cancellableClassStruct_Set() error {
	var err error
	cancellableClassStruct_Once.Do(func() {
		cancellableClassStruct, err = gi.StructNew("Gio", "CancellableClass")
	})
	return err
}

type CancellableClass struct {
	native unsafe.Pointer
}

func CancellableClassNewFromNative(native unsafe.Pointer) *CancellableClass {
	instance := &CancellableClass{native: native}

	return instance
}

/*
CastToCancellableClass down casts any arbitrary Object to CancellableClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a CancellableClass.
*/
func (recv *CancellableClass) CastToCancellableClass(object *gobject.Object) *CancellableClass {
	return CancellableClassNewFromNative(object.Native())
}

func (recv *CancellableClass) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'cancelled' : for field getter : missing Type

// UNSUPPORTED : C value 'cancelled' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved4' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved5' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved5' : for field setter : missing Type

// CancellableClassStruct creates an uninitialised CancellableClass.
func CancellableClassStruct() *CancellableClass {
	err := cancellableClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := CancellableClassNewFromNative(cancellableClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeCancellableClass)
	return structGo
}
func finalizeCancellableClass(obj *CancellableClass) {
	cancellableClassStruct.Free(obj.native)
}

var cancellablePrivateStruct *gi.Struct
var cancellablePrivateStruct_Once sync.Once

func cancellablePrivateStruct_Set() error {
	var err error
	cancellablePrivateStruct_Once.Do(func() {
		cancellablePrivateStruct, err = gi.StructNew("Gio", "CancellablePrivate")
	})
	return err
}

type CancellablePrivate struct {
	native unsafe.Pointer
}

func CancellablePrivateNewFromNative(native unsafe.Pointer) *CancellablePrivate {
	instance := &CancellablePrivate{native: native}

	return instance
}

/*
CastToCancellablePrivate down casts any arbitrary Object to CancellablePrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a CancellablePrivate.
*/
func (recv *CancellablePrivate) CastToCancellablePrivate(object *gobject.Object) *CancellablePrivate {
	return CancellablePrivateNewFromNative(object.Native())
}

func (recv *CancellablePrivate) Native() unsafe.Pointer {
	return recv.native
}

// CancellablePrivateStruct creates an uninitialised CancellablePrivate.
func CancellablePrivateStruct() *CancellablePrivate {
	err := cancellablePrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := CancellablePrivateNewFromNative(cancellablePrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeCancellablePrivate)
	return structGo
}
func finalizeCancellablePrivate(obj *CancellablePrivate) {
	cancellablePrivateStruct.Free(obj.native)
}

var charsetConverterClassStruct *gi.Struct
var charsetConverterClassStruct_Once sync.Once

func charsetConverterClassStruct_Set() error {
	var err error
	charsetConverterClassStruct_Once.Do(func() {
		charsetConverterClassStruct, err = gi.StructNew("Gio", "CharsetConverterClass")
	})
	return err
}

type CharsetConverterClass struct {
	native unsafe.Pointer
}

func CharsetConverterClassNewFromNative(native unsafe.Pointer) *CharsetConverterClass {
	instance := &CharsetConverterClass{native: native}

	return instance
}

/*
CastToCharsetConverterClass down casts any arbitrary Object to CharsetConverterClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a CharsetConverterClass.
*/
func (recv *CharsetConverterClass) CastToCharsetConverterClass(object *gobject.Object) *CharsetConverterClass {
	return CharsetConverterClassNewFromNative(object.Native())
}

func (recv *CharsetConverterClass) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// CharsetConverterClassStruct creates an uninitialised CharsetConverterClass.
func CharsetConverterClassStruct() *CharsetConverterClass {
	err := charsetConverterClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := CharsetConverterClassNewFromNative(charsetConverterClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeCharsetConverterClass)
	return structGo
}
func finalizeCharsetConverterClass(obj *CharsetConverterClass) {
	charsetConverterClassStruct.Free(obj.native)
}

var converterIfaceStruct *gi.Struct
var converterIfaceStruct_Once sync.Once

func converterIfaceStruct_Set() error {
	var err error
	converterIfaceStruct_Once.Do(func() {
		converterIfaceStruct, err = gi.StructNew("Gio", "ConverterIface")
	})
	return err
}

type ConverterIface struct {
	native unsafe.Pointer
}

func ConverterIfaceNewFromNative(native unsafe.Pointer) *ConverterIface {
	instance := &ConverterIface{native: native}

	return instance
}

/*
CastToConverterIface down casts any arbitrary Object to ConverterIface.
Exercise care, as this is a potentially dangerous function
if the Object is not a ConverterIface.
*/
func (recv *ConverterIface) CastToConverterIface(object *gobject.Object) *ConverterIface {
	return ConverterIfaceNewFromNative(object.Native())
}

func (recv *ConverterIface) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'g_iface' : for field getter : no Go type for 'GObject.TypeInterface'

// UNSUPPORTED : C value 'g_iface' : for field setter : no Go type for 'GObject.TypeInterface'

// UNSUPPORTED : C value 'convert' : for field getter : missing Type

// UNSUPPORTED : C value 'convert' : for field setter : missing Type

// UNSUPPORTED : C value 'reset' : for field getter : missing Type

// UNSUPPORTED : C value 'reset' : for field setter : missing Type

// ConverterIfaceStruct creates an uninitialised ConverterIface.
func ConverterIfaceStruct() *ConverterIface {
	err := converterIfaceStruct_Set()
	if err != nil {
		return nil
	}

	structGo := ConverterIfaceNewFromNative(converterIfaceStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeConverterIface)
	return structGo
}
func finalizeConverterIface(obj *ConverterIface) {
	converterIfaceStruct.Free(obj.native)
}

var converterInputStreamClassStruct *gi.Struct
var converterInputStreamClassStruct_Once sync.Once

func converterInputStreamClassStruct_Set() error {
	var err error
	converterInputStreamClassStruct_Once.Do(func() {
		converterInputStreamClassStruct, err = gi.StructNew("Gio", "ConverterInputStreamClass")
	})
	return err
}

type ConverterInputStreamClass struct {
	native unsafe.Pointer
}

func ConverterInputStreamClassNewFromNative(native unsafe.Pointer) *ConverterInputStreamClass {
	instance := &ConverterInputStreamClass{native: native}

	return instance
}

/*
CastToConverterInputStreamClass down casts any arbitrary Object to ConverterInputStreamClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a ConverterInputStreamClass.
*/
func (recv *ConverterInputStreamClass) CastToConverterInputStreamClass(object *gobject.Object) *ConverterInputStreamClass {
	return ConverterInputStreamClassNewFromNative(object.Native())
}

func (recv *ConverterInputStreamClass) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentClass returns the C field 'parent_class'.
func (recv *ConverterInputStreamClass) FieldParentClass() *FilterInputStreamClass {
	argValue := gi.StructFieldGet(converterInputStreamClassStruct, recv.native, "parent_class")
	value := FilterInputStreamClassNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *ConverterInputStreamClass) SetFieldParentClass(value *FilterInputStreamClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.StructFieldSet(converterInputStreamClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value '_g_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved4' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved5' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved5' : for field setter : missing Type

// ConverterInputStreamClassStruct creates an uninitialised ConverterInputStreamClass.
func ConverterInputStreamClassStruct() *ConverterInputStreamClass {
	err := converterInputStreamClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := ConverterInputStreamClassNewFromNative(converterInputStreamClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeConverterInputStreamClass)
	return structGo
}
func finalizeConverterInputStreamClass(obj *ConverterInputStreamClass) {
	converterInputStreamClassStruct.Free(obj.native)
}

var converterInputStreamPrivateStruct *gi.Struct
var converterInputStreamPrivateStruct_Once sync.Once

func converterInputStreamPrivateStruct_Set() error {
	var err error
	converterInputStreamPrivateStruct_Once.Do(func() {
		converterInputStreamPrivateStruct, err = gi.StructNew("Gio", "ConverterInputStreamPrivate")
	})
	return err
}

type ConverterInputStreamPrivate struct {
	native unsafe.Pointer
}

func ConverterInputStreamPrivateNewFromNative(native unsafe.Pointer) *ConverterInputStreamPrivate {
	instance := &ConverterInputStreamPrivate{native: native}

	return instance
}

/*
CastToConverterInputStreamPrivate down casts any arbitrary Object to ConverterInputStreamPrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a ConverterInputStreamPrivate.
*/
func (recv *ConverterInputStreamPrivate) CastToConverterInputStreamPrivate(object *gobject.Object) *ConverterInputStreamPrivate {
	return ConverterInputStreamPrivateNewFromNative(object.Native())
}

func (recv *ConverterInputStreamPrivate) Native() unsafe.Pointer {
	return recv.native
}

// ConverterInputStreamPrivateStruct creates an uninitialised ConverterInputStreamPrivate.
func ConverterInputStreamPrivateStruct() *ConverterInputStreamPrivate {
	err := converterInputStreamPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := ConverterInputStreamPrivateNewFromNative(converterInputStreamPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeConverterInputStreamPrivate)
	return structGo
}
func finalizeConverterInputStreamPrivate(obj *ConverterInputStreamPrivate) {
	converterInputStreamPrivateStruct.Free(obj.native)
}

var converterOutputStreamClassStruct *gi.Struct
var converterOutputStreamClassStruct_Once sync.Once

func converterOutputStreamClassStruct_Set() error {
	var err error
	converterOutputStreamClassStruct_Once.Do(func() {
		converterOutputStreamClassStruct, err = gi.StructNew("Gio", "ConverterOutputStreamClass")
	})
	return err
}

type ConverterOutputStreamClass struct {
	native unsafe.Pointer
}

func ConverterOutputStreamClassNewFromNative(native unsafe.Pointer) *ConverterOutputStreamClass {
	instance := &ConverterOutputStreamClass{native: native}

	return instance
}

/*
CastToConverterOutputStreamClass down casts any arbitrary Object to ConverterOutputStreamClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a ConverterOutputStreamClass.
*/
func (recv *ConverterOutputStreamClass) CastToConverterOutputStreamClass(object *gobject.Object) *ConverterOutputStreamClass {
	return ConverterOutputStreamClassNewFromNative(object.Native())
}

func (recv *ConverterOutputStreamClass) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentClass returns the C field 'parent_class'.
func (recv *ConverterOutputStreamClass) FieldParentClass() *FilterOutputStreamClass {
	argValue := gi.StructFieldGet(converterOutputStreamClassStruct, recv.native, "parent_class")
	value := FilterOutputStreamClassNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *ConverterOutputStreamClass) SetFieldParentClass(value *FilterOutputStreamClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.StructFieldSet(converterOutputStreamClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value '_g_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved4' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved5' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved5' : for field setter : missing Type

// ConverterOutputStreamClassStruct creates an uninitialised ConverterOutputStreamClass.
func ConverterOutputStreamClassStruct() *ConverterOutputStreamClass {
	err := converterOutputStreamClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := ConverterOutputStreamClassNewFromNative(converterOutputStreamClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeConverterOutputStreamClass)
	return structGo
}
func finalizeConverterOutputStreamClass(obj *ConverterOutputStreamClass) {
	converterOutputStreamClassStruct.Free(obj.native)
}

var converterOutputStreamPrivateStruct *gi.Struct
var converterOutputStreamPrivateStruct_Once sync.Once

func converterOutputStreamPrivateStruct_Set() error {
	var err error
	converterOutputStreamPrivateStruct_Once.Do(func() {
		converterOutputStreamPrivateStruct, err = gi.StructNew("Gio", "ConverterOutputStreamPrivate")
	})
	return err
}

type ConverterOutputStreamPrivate struct {
	native unsafe.Pointer
}

func ConverterOutputStreamPrivateNewFromNative(native unsafe.Pointer) *ConverterOutputStreamPrivate {
	instance := &ConverterOutputStreamPrivate{native: native}

	return instance
}

/*
CastToConverterOutputStreamPrivate down casts any arbitrary Object to ConverterOutputStreamPrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a ConverterOutputStreamPrivate.
*/
func (recv *ConverterOutputStreamPrivate) CastToConverterOutputStreamPrivate(object *gobject.Object) *ConverterOutputStreamPrivate {
	return ConverterOutputStreamPrivateNewFromNative(object.Native())
}

func (recv *ConverterOutputStreamPrivate) Native() unsafe.Pointer {
	return recv.native
}

// ConverterOutputStreamPrivateStruct creates an uninitialised ConverterOutputStreamPrivate.
func ConverterOutputStreamPrivateStruct() *ConverterOutputStreamPrivate {
	err := converterOutputStreamPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := ConverterOutputStreamPrivateNewFromNative(converterOutputStreamPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeConverterOutputStreamPrivate)
	return structGo
}
func finalizeConverterOutputStreamPrivate(obj *ConverterOutputStreamPrivate) {
	converterOutputStreamPrivateStruct.Free(obj.native)
}

var credentialsClassStruct *gi.Struct
var credentialsClassStruct_Once sync.Once

func credentialsClassStruct_Set() error {
	var err error
	credentialsClassStruct_Once.Do(func() {
		credentialsClassStruct, err = gi.StructNew("Gio", "CredentialsClass")
	})
	return err
}

type CredentialsClass struct {
	native unsafe.Pointer
}

func CredentialsClassNewFromNative(native unsafe.Pointer) *CredentialsClass {
	instance := &CredentialsClass{native: native}

	return instance
}

/*
CastToCredentialsClass down casts any arbitrary Object to CredentialsClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a CredentialsClass.
*/
func (recv *CredentialsClass) CastToCredentialsClass(object *gobject.Object) *CredentialsClass {
	return CredentialsClassNewFromNative(object.Native())
}

func (recv *CredentialsClass) Native() unsafe.Pointer {
	return recv.native
}

// CredentialsClassStruct creates an uninitialised CredentialsClass.
func CredentialsClassStruct() *CredentialsClass {
	err := credentialsClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := CredentialsClassNewFromNative(credentialsClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeCredentialsClass)
	return structGo
}
func finalizeCredentialsClass(obj *CredentialsClass) {
	credentialsClassStruct.Free(obj.native)
}

var dBusAnnotationInfoStruct *gi.Struct
var dBusAnnotationInfoStruct_Once sync.Once

func dBusAnnotationInfoStruct_Set() error {
	var err error
	dBusAnnotationInfoStruct_Once.Do(func() {
		dBusAnnotationInfoStruct, err = gi.StructNew("Gio", "DBusAnnotationInfo")
	})
	return err
}

type DBusAnnotationInfo struct {
	native unsafe.Pointer
}

func DBusAnnotationInfoNewFromNative(native unsafe.Pointer) *DBusAnnotationInfo {
	instance := &DBusAnnotationInfo{native: native}

	return instance
}

/*
CastToDBusAnnotationInfo down casts any arbitrary Object to DBusAnnotationInfo.
Exercise care, as this is a potentially dangerous function
if the Object is not a DBusAnnotationInfo.
*/
func (recv *DBusAnnotationInfo) CastToDBusAnnotationInfo(object *gobject.Object) *DBusAnnotationInfo {
	return DBusAnnotationInfoNewFromNative(object.Native())
}

func (recv *DBusAnnotationInfo) Native() unsafe.Pointer {
	return recv.native
}

// FieldRefCount returns the C field 'ref_count'.
func (recv *DBusAnnotationInfo) FieldRefCount() int32 {
	argValue := gi.StructFieldGet(dBusAnnotationInfoStruct, recv.native, "ref_count")
	value := argValue.Int32()
	return value
}

// SetFieldRefCount sets the value of the C field 'ref_count'.
func (recv *DBusAnnotationInfo) SetFieldRefCount(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.StructFieldSet(dBusAnnotationInfoStruct, recv.native, "ref_count", argValue)
}

// FieldKey returns the C field 'key'.
func (recv *DBusAnnotationInfo) FieldKey() string {
	argValue := gi.StructFieldGet(dBusAnnotationInfoStruct, recv.native, "key")
	value := argValue.String(false)
	return value
}

// SetFieldKey sets the value of the C field 'key'.
func (recv *DBusAnnotationInfo) SetFieldKey(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.StructFieldSet(dBusAnnotationInfoStruct, recv.native, "key", argValue)
}

// FieldValue returns the C field 'value'.
func (recv *DBusAnnotationInfo) FieldValue() string {
	argValue := gi.StructFieldGet(dBusAnnotationInfoStruct, recv.native, "value")
	value := argValue.String(false)
	return value
}

// SetFieldValue sets the value of the C field 'value'.
func (recv *DBusAnnotationInfo) SetFieldValue(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.StructFieldSet(dBusAnnotationInfoStruct, recv.native, "value", argValue)
}

// UNSUPPORTED : C value 'annotations' : for field getter : missing Type

// UNSUPPORTED : C value 'annotations' : for field setter : missing Type

var dBusAnnotationInfoRefFunction *gi.Function
var dBusAnnotationInfoRefFunction_Once sync.Once

func dBusAnnotationInfoRefFunction_Set() error {
	var err error
	dBusAnnotationInfoRefFunction_Once.Do(func() {
		err = dBusAnnotationInfoStruct_Set()
		if err != nil {
			return
		}
		dBusAnnotationInfoRefFunction, err = dBusAnnotationInfoStruct.InvokerNew("ref")
	})
	return err
}

// Ref is a representation of the C type g_dbus_annotation_info_ref.
func (recv *DBusAnnotationInfo) Ref() *DBusAnnotationInfo {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := dBusAnnotationInfoRefFunction_Set()
	if err == nil {
		ret = dBusAnnotationInfoRefFunction.Invoke(inArgs[:], nil)
	}

	retGo := DBusAnnotationInfoNewFromNative(ret.Pointer())

	return retGo
}

var dBusAnnotationInfoUnrefFunction *gi.Function
var dBusAnnotationInfoUnrefFunction_Once sync.Once

func dBusAnnotationInfoUnrefFunction_Set() error {
	var err error
	dBusAnnotationInfoUnrefFunction_Once.Do(func() {
		err = dBusAnnotationInfoStruct_Set()
		if err != nil {
			return
		}
		dBusAnnotationInfoUnrefFunction, err = dBusAnnotationInfoStruct.InvokerNew("unref")
	})
	return err
}

// Unref is a representation of the C type g_dbus_annotation_info_unref.
func (recv *DBusAnnotationInfo) Unref() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := dBusAnnotationInfoUnrefFunction_Set()
	if err == nil {
		dBusAnnotationInfoUnrefFunction.Invoke(inArgs[:], nil)
	}

	return
}

// DBusAnnotationInfoStruct creates an uninitialised DBusAnnotationInfo.
func DBusAnnotationInfoStruct() *DBusAnnotationInfo {
	err := dBusAnnotationInfoStruct_Set()
	if err != nil {
		return nil
	}

	structGo := DBusAnnotationInfoNewFromNative(dBusAnnotationInfoStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeDBusAnnotationInfo)
	return structGo
}
func finalizeDBusAnnotationInfo(obj *DBusAnnotationInfo) {
	dBusAnnotationInfoStruct.Free(obj.native)
}

var dBusArgInfoStruct *gi.Struct
var dBusArgInfoStruct_Once sync.Once

func dBusArgInfoStruct_Set() error {
	var err error
	dBusArgInfoStruct_Once.Do(func() {
		dBusArgInfoStruct, err = gi.StructNew("Gio", "DBusArgInfo")
	})
	return err
}

type DBusArgInfo struct {
	native unsafe.Pointer
}

func DBusArgInfoNewFromNative(native unsafe.Pointer) *DBusArgInfo {
	instance := &DBusArgInfo{native: native}

	return instance
}

/*
CastToDBusArgInfo down casts any arbitrary Object to DBusArgInfo.
Exercise care, as this is a potentially dangerous function
if the Object is not a DBusArgInfo.
*/
func (recv *DBusArgInfo) CastToDBusArgInfo(object *gobject.Object) *DBusArgInfo {
	return DBusArgInfoNewFromNative(object.Native())
}

func (recv *DBusArgInfo) Native() unsafe.Pointer {
	return recv.native
}

// FieldRefCount returns the C field 'ref_count'.
func (recv *DBusArgInfo) FieldRefCount() int32 {
	argValue := gi.StructFieldGet(dBusArgInfoStruct, recv.native, "ref_count")
	value := argValue.Int32()
	return value
}

// SetFieldRefCount sets the value of the C field 'ref_count'.
func (recv *DBusArgInfo) SetFieldRefCount(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.StructFieldSet(dBusArgInfoStruct, recv.native, "ref_count", argValue)
}

// FieldName returns the C field 'name'.
func (recv *DBusArgInfo) FieldName() string {
	argValue := gi.StructFieldGet(dBusArgInfoStruct, recv.native, "name")
	value := argValue.String(false)
	return value
}

// SetFieldName sets the value of the C field 'name'.
func (recv *DBusArgInfo) SetFieldName(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.StructFieldSet(dBusArgInfoStruct, recv.native, "name", argValue)
}

// FieldSignature returns the C field 'signature'.
func (recv *DBusArgInfo) FieldSignature() string {
	argValue := gi.StructFieldGet(dBusArgInfoStruct, recv.native, "signature")
	value := argValue.String(false)
	return value
}

// SetFieldSignature sets the value of the C field 'signature'.
func (recv *DBusArgInfo) SetFieldSignature(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.StructFieldSet(dBusArgInfoStruct, recv.native, "signature", argValue)
}

// UNSUPPORTED : C value 'annotations' : for field getter : missing Type

// UNSUPPORTED : C value 'annotations' : for field setter : missing Type

var dBusArgInfoRefFunction *gi.Function
var dBusArgInfoRefFunction_Once sync.Once

func dBusArgInfoRefFunction_Set() error {
	var err error
	dBusArgInfoRefFunction_Once.Do(func() {
		err = dBusArgInfoStruct_Set()
		if err != nil {
			return
		}
		dBusArgInfoRefFunction, err = dBusArgInfoStruct.InvokerNew("ref")
	})
	return err
}

// Ref is a representation of the C type g_dbus_arg_info_ref.
func (recv *DBusArgInfo) Ref() *DBusArgInfo {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := dBusArgInfoRefFunction_Set()
	if err == nil {
		ret = dBusArgInfoRefFunction.Invoke(inArgs[:], nil)
	}

	retGo := DBusArgInfoNewFromNative(ret.Pointer())

	return retGo
}

var dBusArgInfoUnrefFunction *gi.Function
var dBusArgInfoUnrefFunction_Once sync.Once

func dBusArgInfoUnrefFunction_Set() error {
	var err error
	dBusArgInfoUnrefFunction_Once.Do(func() {
		err = dBusArgInfoStruct_Set()
		if err != nil {
			return
		}
		dBusArgInfoUnrefFunction, err = dBusArgInfoStruct.InvokerNew("unref")
	})
	return err
}

// Unref is a representation of the C type g_dbus_arg_info_unref.
func (recv *DBusArgInfo) Unref() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := dBusArgInfoUnrefFunction_Set()
	if err == nil {
		dBusArgInfoUnrefFunction.Invoke(inArgs[:], nil)
	}

	return
}

// DBusArgInfoStruct creates an uninitialised DBusArgInfo.
func DBusArgInfoStruct() *DBusArgInfo {
	err := dBusArgInfoStruct_Set()
	if err != nil {
		return nil
	}

	structGo := DBusArgInfoNewFromNative(dBusArgInfoStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeDBusArgInfo)
	return structGo
}
func finalizeDBusArgInfo(obj *DBusArgInfo) {
	dBusArgInfoStruct.Free(obj.native)
}

var dBusErrorEntryStruct *gi.Struct
var dBusErrorEntryStruct_Once sync.Once

func dBusErrorEntryStruct_Set() error {
	var err error
	dBusErrorEntryStruct_Once.Do(func() {
		dBusErrorEntryStruct, err = gi.StructNew("Gio", "DBusErrorEntry")
	})
	return err
}

type DBusErrorEntry struct {
	native unsafe.Pointer
}

func DBusErrorEntryNewFromNative(native unsafe.Pointer) *DBusErrorEntry {
	instance := &DBusErrorEntry{native: native}

	return instance
}

/*
CastToDBusErrorEntry down casts any arbitrary Object to DBusErrorEntry.
Exercise care, as this is a potentially dangerous function
if the Object is not a DBusErrorEntry.
*/
func (recv *DBusErrorEntry) CastToDBusErrorEntry(object *gobject.Object) *DBusErrorEntry {
	return DBusErrorEntryNewFromNative(object.Native())
}

func (recv *DBusErrorEntry) Native() unsafe.Pointer {
	return recv.native
}

// FieldErrorCode returns the C field 'error_code'.
func (recv *DBusErrorEntry) FieldErrorCode() int32 {
	argValue := gi.StructFieldGet(dBusErrorEntryStruct, recv.native, "error_code")
	value := argValue.Int32()
	return value
}

// SetFieldErrorCode sets the value of the C field 'error_code'.
func (recv *DBusErrorEntry) SetFieldErrorCode(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.StructFieldSet(dBusErrorEntryStruct, recv.native, "error_code", argValue)
}

// FieldDbusErrorName returns the C field 'dbus_error_name'.
func (recv *DBusErrorEntry) FieldDbusErrorName() string {
	argValue := gi.StructFieldGet(dBusErrorEntryStruct, recv.native, "dbus_error_name")
	value := argValue.String(false)
	return value
}

// SetFieldDbusErrorName sets the value of the C field 'dbus_error_name'.
func (recv *DBusErrorEntry) SetFieldDbusErrorName(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.StructFieldSet(dBusErrorEntryStruct, recv.native, "dbus_error_name", argValue)
}

// DBusErrorEntryStruct creates an uninitialised DBusErrorEntry.
func DBusErrorEntryStruct() *DBusErrorEntry {
	err := dBusErrorEntryStruct_Set()
	if err != nil {
		return nil
	}

	structGo := DBusErrorEntryNewFromNative(dBusErrorEntryStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeDBusErrorEntry)
	return structGo
}
func finalizeDBusErrorEntry(obj *DBusErrorEntry) {
	dBusErrorEntryStruct.Free(obj.native)
}

var dBusInterfaceIfaceStruct *gi.Struct
var dBusInterfaceIfaceStruct_Once sync.Once

func dBusInterfaceIfaceStruct_Set() error {
	var err error
	dBusInterfaceIfaceStruct_Once.Do(func() {
		dBusInterfaceIfaceStruct, err = gi.StructNew("Gio", "DBusInterfaceIface")
	})
	return err
}

type DBusInterfaceIface struct {
	native unsafe.Pointer
}

func DBusInterfaceIfaceNewFromNative(native unsafe.Pointer) *DBusInterfaceIface {
	instance := &DBusInterfaceIface{native: native}

	return instance
}

/*
CastToDBusInterfaceIface down casts any arbitrary Object to DBusInterfaceIface.
Exercise care, as this is a potentially dangerous function
if the Object is not a DBusInterfaceIface.
*/
func (recv *DBusInterfaceIface) CastToDBusInterfaceIface(object *gobject.Object) *DBusInterfaceIface {
	return DBusInterfaceIfaceNewFromNative(object.Native())
}

func (recv *DBusInterfaceIface) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'parent_iface' : for field getter : no Go type for 'GObject.TypeInterface'

// UNSUPPORTED : C value 'parent_iface' : for field setter : no Go type for 'GObject.TypeInterface'

// UNSUPPORTED : C value 'get_info' : for field getter : missing Type

// UNSUPPORTED : C value 'get_info' : for field setter : missing Type

// UNSUPPORTED : C value 'get_object' : for field getter : missing Type

// UNSUPPORTED : C value 'get_object' : for field setter : missing Type

// UNSUPPORTED : C value 'set_object' : for field getter : missing Type

// UNSUPPORTED : C value 'set_object' : for field setter : missing Type

// UNSUPPORTED : C value 'dup_object' : for field getter : missing Type

// UNSUPPORTED : C value 'dup_object' : for field setter : missing Type

// DBusInterfaceIfaceStruct creates an uninitialised DBusInterfaceIface.
func DBusInterfaceIfaceStruct() *DBusInterfaceIface {
	err := dBusInterfaceIfaceStruct_Set()
	if err != nil {
		return nil
	}

	structGo := DBusInterfaceIfaceNewFromNative(dBusInterfaceIfaceStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeDBusInterfaceIface)
	return structGo
}
func finalizeDBusInterfaceIface(obj *DBusInterfaceIface) {
	dBusInterfaceIfaceStruct.Free(obj.native)
}

var dBusInterfaceInfoStruct *gi.Struct
var dBusInterfaceInfoStruct_Once sync.Once

func dBusInterfaceInfoStruct_Set() error {
	var err error
	dBusInterfaceInfoStruct_Once.Do(func() {
		dBusInterfaceInfoStruct, err = gi.StructNew("Gio", "DBusInterfaceInfo")
	})
	return err
}

type DBusInterfaceInfo struct {
	native unsafe.Pointer
}

func DBusInterfaceInfoNewFromNative(native unsafe.Pointer) *DBusInterfaceInfo {
	instance := &DBusInterfaceInfo{native: native}

	return instance
}

/*
CastToDBusInterfaceInfo down casts any arbitrary Object to DBusInterfaceInfo.
Exercise care, as this is a potentially dangerous function
if the Object is not a DBusInterfaceInfo.
*/
func (recv *DBusInterfaceInfo) CastToDBusInterfaceInfo(object *gobject.Object) *DBusInterfaceInfo {
	return DBusInterfaceInfoNewFromNative(object.Native())
}

func (recv *DBusInterfaceInfo) Native() unsafe.Pointer {
	return recv.native
}

// FieldRefCount returns the C field 'ref_count'.
func (recv *DBusInterfaceInfo) FieldRefCount() int32 {
	argValue := gi.StructFieldGet(dBusInterfaceInfoStruct, recv.native, "ref_count")
	value := argValue.Int32()
	return value
}

// SetFieldRefCount sets the value of the C field 'ref_count'.
func (recv *DBusInterfaceInfo) SetFieldRefCount(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.StructFieldSet(dBusInterfaceInfoStruct, recv.native, "ref_count", argValue)
}

// FieldName returns the C field 'name'.
func (recv *DBusInterfaceInfo) FieldName() string {
	argValue := gi.StructFieldGet(dBusInterfaceInfoStruct, recv.native, "name")
	value := argValue.String(false)
	return value
}

// SetFieldName sets the value of the C field 'name'.
func (recv *DBusInterfaceInfo) SetFieldName(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.StructFieldSet(dBusInterfaceInfoStruct, recv.native, "name", argValue)
}

// UNSUPPORTED : C value 'methods' : for field getter : missing Type

// UNSUPPORTED : C value 'methods' : for field setter : missing Type

// UNSUPPORTED : C value 'signals' : for field getter : missing Type

// UNSUPPORTED : C value 'signals' : for field setter : missing Type

// UNSUPPORTED : C value 'properties' : for field getter : missing Type

// UNSUPPORTED : C value 'properties' : for field setter : missing Type

// UNSUPPORTED : C value 'annotations' : for field getter : missing Type

// UNSUPPORTED : C value 'annotations' : for field setter : missing Type

var dBusInterfaceInfoCacheBuildFunction *gi.Function
var dBusInterfaceInfoCacheBuildFunction_Once sync.Once

func dBusInterfaceInfoCacheBuildFunction_Set() error {
	var err error
	dBusInterfaceInfoCacheBuildFunction_Once.Do(func() {
		err = dBusInterfaceInfoStruct_Set()
		if err != nil {
			return
		}
		dBusInterfaceInfoCacheBuildFunction, err = dBusInterfaceInfoStruct.InvokerNew("cache_build")
	})
	return err
}

// CacheBuild is a representation of the C type g_dbus_interface_info_cache_build.
func (recv *DBusInterfaceInfo) CacheBuild() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := dBusInterfaceInfoCacheBuildFunction_Set()
	if err == nil {
		dBusInterfaceInfoCacheBuildFunction.Invoke(inArgs[:], nil)
	}

	return
}

var dBusInterfaceInfoCacheReleaseFunction *gi.Function
var dBusInterfaceInfoCacheReleaseFunction_Once sync.Once

func dBusInterfaceInfoCacheReleaseFunction_Set() error {
	var err error
	dBusInterfaceInfoCacheReleaseFunction_Once.Do(func() {
		err = dBusInterfaceInfoStruct_Set()
		if err != nil {
			return
		}
		dBusInterfaceInfoCacheReleaseFunction, err = dBusInterfaceInfoStruct.InvokerNew("cache_release")
	})
	return err
}

// CacheRelease is a representation of the C type g_dbus_interface_info_cache_release.
func (recv *DBusInterfaceInfo) CacheRelease() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := dBusInterfaceInfoCacheReleaseFunction_Set()
	if err == nil {
		dBusInterfaceInfoCacheReleaseFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'g_dbus_interface_info_generate_xml' : parameter 'string_builder' of type 'GLib.String' not supported

var dBusInterfaceInfoLookupMethodFunction *gi.Function
var dBusInterfaceInfoLookupMethodFunction_Once sync.Once

func dBusInterfaceInfoLookupMethodFunction_Set() error {
	var err error
	dBusInterfaceInfoLookupMethodFunction_Once.Do(func() {
		err = dBusInterfaceInfoStruct_Set()
		if err != nil {
			return
		}
		dBusInterfaceInfoLookupMethodFunction, err = dBusInterfaceInfoStruct.InvokerNew("lookup_method")
	})
	return err
}

// LookupMethod is a representation of the C type g_dbus_interface_info_lookup_method.
func (recv *DBusInterfaceInfo) LookupMethod(name string) *DBusMethodInfo {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(name)

	var ret gi.Argument

	err := dBusInterfaceInfoLookupMethodFunction_Set()
	if err == nil {
		ret = dBusInterfaceInfoLookupMethodFunction.Invoke(inArgs[:], nil)
	}

	retGo := DBusMethodInfoNewFromNative(ret.Pointer())

	return retGo
}

var dBusInterfaceInfoLookupPropertyFunction *gi.Function
var dBusInterfaceInfoLookupPropertyFunction_Once sync.Once

func dBusInterfaceInfoLookupPropertyFunction_Set() error {
	var err error
	dBusInterfaceInfoLookupPropertyFunction_Once.Do(func() {
		err = dBusInterfaceInfoStruct_Set()
		if err != nil {
			return
		}
		dBusInterfaceInfoLookupPropertyFunction, err = dBusInterfaceInfoStruct.InvokerNew("lookup_property")
	})
	return err
}

// LookupProperty is a representation of the C type g_dbus_interface_info_lookup_property.
func (recv *DBusInterfaceInfo) LookupProperty(name string) *DBusPropertyInfo {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(name)

	var ret gi.Argument

	err := dBusInterfaceInfoLookupPropertyFunction_Set()
	if err == nil {
		ret = dBusInterfaceInfoLookupPropertyFunction.Invoke(inArgs[:], nil)
	}

	retGo := DBusPropertyInfoNewFromNative(ret.Pointer())

	return retGo
}

var dBusInterfaceInfoLookupSignalFunction *gi.Function
var dBusInterfaceInfoLookupSignalFunction_Once sync.Once

func dBusInterfaceInfoLookupSignalFunction_Set() error {
	var err error
	dBusInterfaceInfoLookupSignalFunction_Once.Do(func() {
		err = dBusInterfaceInfoStruct_Set()
		if err != nil {
			return
		}
		dBusInterfaceInfoLookupSignalFunction, err = dBusInterfaceInfoStruct.InvokerNew("lookup_signal")
	})
	return err
}

// LookupSignal is a representation of the C type g_dbus_interface_info_lookup_signal.
func (recv *DBusInterfaceInfo) LookupSignal(name string) *DBusSignalInfo {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(name)

	var ret gi.Argument

	err := dBusInterfaceInfoLookupSignalFunction_Set()
	if err == nil {
		ret = dBusInterfaceInfoLookupSignalFunction.Invoke(inArgs[:], nil)
	}

	retGo := DBusSignalInfoNewFromNative(ret.Pointer())

	return retGo
}

var dBusInterfaceInfoRefFunction *gi.Function
var dBusInterfaceInfoRefFunction_Once sync.Once

func dBusInterfaceInfoRefFunction_Set() error {
	var err error
	dBusInterfaceInfoRefFunction_Once.Do(func() {
		err = dBusInterfaceInfoStruct_Set()
		if err != nil {
			return
		}
		dBusInterfaceInfoRefFunction, err = dBusInterfaceInfoStruct.InvokerNew("ref")
	})
	return err
}

// Ref is a representation of the C type g_dbus_interface_info_ref.
func (recv *DBusInterfaceInfo) Ref() *DBusInterfaceInfo {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := dBusInterfaceInfoRefFunction_Set()
	if err == nil {
		ret = dBusInterfaceInfoRefFunction.Invoke(inArgs[:], nil)
	}

	retGo := DBusInterfaceInfoNewFromNative(ret.Pointer())

	return retGo
}

var dBusInterfaceInfoUnrefFunction *gi.Function
var dBusInterfaceInfoUnrefFunction_Once sync.Once

func dBusInterfaceInfoUnrefFunction_Set() error {
	var err error
	dBusInterfaceInfoUnrefFunction_Once.Do(func() {
		err = dBusInterfaceInfoStruct_Set()
		if err != nil {
			return
		}
		dBusInterfaceInfoUnrefFunction, err = dBusInterfaceInfoStruct.InvokerNew("unref")
	})
	return err
}

// Unref is a representation of the C type g_dbus_interface_info_unref.
func (recv *DBusInterfaceInfo) Unref() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := dBusInterfaceInfoUnrefFunction_Set()
	if err == nil {
		dBusInterfaceInfoUnrefFunction.Invoke(inArgs[:], nil)
	}

	return
}

// DBusInterfaceInfoStruct creates an uninitialised DBusInterfaceInfo.
func DBusInterfaceInfoStruct() *DBusInterfaceInfo {
	err := dBusInterfaceInfoStruct_Set()
	if err != nil {
		return nil
	}

	structGo := DBusInterfaceInfoNewFromNative(dBusInterfaceInfoStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeDBusInterfaceInfo)
	return structGo
}
func finalizeDBusInterfaceInfo(obj *DBusInterfaceInfo) {
	dBusInterfaceInfoStruct.Free(obj.native)
}

var dBusInterfaceSkeletonClassStruct *gi.Struct
var dBusInterfaceSkeletonClassStruct_Once sync.Once

func dBusInterfaceSkeletonClassStruct_Set() error {
	var err error
	dBusInterfaceSkeletonClassStruct_Once.Do(func() {
		dBusInterfaceSkeletonClassStruct, err = gi.StructNew("Gio", "DBusInterfaceSkeletonClass")
	})
	return err
}

type DBusInterfaceSkeletonClass struct {
	native unsafe.Pointer
}

func DBusInterfaceSkeletonClassNewFromNative(native unsafe.Pointer) *DBusInterfaceSkeletonClass {
	instance := &DBusInterfaceSkeletonClass{native: native}

	return instance
}

/*
CastToDBusInterfaceSkeletonClass down casts any arbitrary Object to DBusInterfaceSkeletonClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a DBusInterfaceSkeletonClass.
*/
func (recv *DBusInterfaceSkeletonClass) CastToDBusInterfaceSkeletonClass(object *gobject.Object) *DBusInterfaceSkeletonClass {
	return DBusInterfaceSkeletonClassNewFromNative(object.Native())
}

func (recv *DBusInterfaceSkeletonClass) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'get_info' : for field getter : missing Type

// UNSUPPORTED : C value 'get_info' : for field setter : missing Type

// UNSUPPORTED : C value 'get_vtable' : for field getter : missing Type

// UNSUPPORTED : C value 'get_vtable' : for field setter : missing Type

// UNSUPPORTED : C value 'get_properties' : for field getter : missing Type

// UNSUPPORTED : C value 'get_properties' : for field setter : missing Type

// UNSUPPORTED : C value 'flush' : for field getter : missing Type

// UNSUPPORTED : C value 'flush' : for field setter : missing Type

// UNSUPPORTED : C value 'g_authorize_method' : for field getter : missing Type

// UNSUPPORTED : C value 'g_authorize_method' : for field setter : missing Type

// DBusInterfaceSkeletonClassStruct creates an uninitialised DBusInterfaceSkeletonClass.
func DBusInterfaceSkeletonClassStruct() *DBusInterfaceSkeletonClass {
	err := dBusInterfaceSkeletonClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := DBusInterfaceSkeletonClassNewFromNative(dBusInterfaceSkeletonClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeDBusInterfaceSkeletonClass)
	return structGo
}
func finalizeDBusInterfaceSkeletonClass(obj *DBusInterfaceSkeletonClass) {
	dBusInterfaceSkeletonClassStruct.Free(obj.native)
}

var dBusInterfaceSkeletonPrivateStruct *gi.Struct
var dBusInterfaceSkeletonPrivateStruct_Once sync.Once

func dBusInterfaceSkeletonPrivateStruct_Set() error {
	var err error
	dBusInterfaceSkeletonPrivateStruct_Once.Do(func() {
		dBusInterfaceSkeletonPrivateStruct, err = gi.StructNew("Gio", "DBusInterfaceSkeletonPrivate")
	})
	return err
}

type DBusInterfaceSkeletonPrivate struct {
	native unsafe.Pointer
}

func DBusInterfaceSkeletonPrivateNewFromNative(native unsafe.Pointer) *DBusInterfaceSkeletonPrivate {
	instance := &DBusInterfaceSkeletonPrivate{native: native}

	return instance
}

/*
CastToDBusInterfaceSkeletonPrivate down casts any arbitrary Object to DBusInterfaceSkeletonPrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a DBusInterfaceSkeletonPrivate.
*/
func (recv *DBusInterfaceSkeletonPrivate) CastToDBusInterfaceSkeletonPrivate(object *gobject.Object) *DBusInterfaceSkeletonPrivate {
	return DBusInterfaceSkeletonPrivateNewFromNative(object.Native())
}

func (recv *DBusInterfaceSkeletonPrivate) Native() unsafe.Pointer {
	return recv.native
}

// DBusInterfaceSkeletonPrivateStruct creates an uninitialised DBusInterfaceSkeletonPrivate.
func DBusInterfaceSkeletonPrivateStruct() *DBusInterfaceSkeletonPrivate {
	err := dBusInterfaceSkeletonPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := DBusInterfaceSkeletonPrivateNewFromNative(dBusInterfaceSkeletonPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeDBusInterfaceSkeletonPrivate)
	return structGo
}
func finalizeDBusInterfaceSkeletonPrivate(obj *DBusInterfaceSkeletonPrivate) {
	dBusInterfaceSkeletonPrivateStruct.Free(obj.native)
}

var dBusInterfaceVTableStruct *gi.Struct
var dBusInterfaceVTableStruct_Once sync.Once

func dBusInterfaceVTableStruct_Set() error {
	var err error
	dBusInterfaceVTableStruct_Once.Do(func() {
		dBusInterfaceVTableStruct, err = gi.StructNew("Gio", "DBusInterfaceVTable")
	})
	return err
}

type DBusInterfaceVTable struct {
	native unsafe.Pointer
}

func DBusInterfaceVTableNewFromNative(native unsafe.Pointer) *DBusInterfaceVTable {
	instance := &DBusInterfaceVTable{native: native}

	return instance
}

/*
CastToDBusInterfaceVTable down casts any arbitrary Object to DBusInterfaceVTable.
Exercise care, as this is a potentially dangerous function
if the Object is not a DBusInterfaceVTable.
*/
func (recv *DBusInterfaceVTable) CastToDBusInterfaceVTable(object *gobject.Object) *DBusInterfaceVTable {
	return DBusInterfaceVTableNewFromNative(object.Native())
}

func (recv *DBusInterfaceVTable) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'method_call' : for field getter : no Go type for 'DBusInterfaceMethodCallFunc'

// UNSUPPORTED : C value 'method_call' : for field setter : no Go type for 'DBusInterfaceMethodCallFunc'

// UNSUPPORTED : C value 'get_property' : for field getter : no Go type for 'DBusInterfaceGetPropertyFunc'

// UNSUPPORTED : C value 'get_property' : for field setter : no Go type for 'DBusInterfaceGetPropertyFunc'

// UNSUPPORTED : C value 'set_property' : for field getter : no Go type for 'DBusInterfaceSetPropertyFunc'

// UNSUPPORTED : C value 'set_property' : for field setter : no Go type for 'DBusInterfaceSetPropertyFunc'

// DBusInterfaceVTableStruct creates an uninitialised DBusInterfaceVTable.
func DBusInterfaceVTableStruct() *DBusInterfaceVTable {
	err := dBusInterfaceVTableStruct_Set()
	if err != nil {
		return nil
	}

	structGo := DBusInterfaceVTableNewFromNative(dBusInterfaceVTableStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeDBusInterfaceVTable)
	return structGo
}
func finalizeDBusInterfaceVTable(obj *DBusInterfaceVTable) {
	dBusInterfaceVTableStruct.Free(obj.native)
}

var dBusMethodInfoStruct *gi.Struct
var dBusMethodInfoStruct_Once sync.Once

func dBusMethodInfoStruct_Set() error {
	var err error
	dBusMethodInfoStruct_Once.Do(func() {
		dBusMethodInfoStruct, err = gi.StructNew("Gio", "DBusMethodInfo")
	})
	return err
}

type DBusMethodInfo struct {
	native unsafe.Pointer
}

func DBusMethodInfoNewFromNative(native unsafe.Pointer) *DBusMethodInfo {
	instance := &DBusMethodInfo{native: native}

	return instance
}

/*
CastToDBusMethodInfo down casts any arbitrary Object to DBusMethodInfo.
Exercise care, as this is a potentially dangerous function
if the Object is not a DBusMethodInfo.
*/
func (recv *DBusMethodInfo) CastToDBusMethodInfo(object *gobject.Object) *DBusMethodInfo {
	return DBusMethodInfoNewFromNative(object.Native())
}

func (recv *DBusMethodInfo) Native() unsafe.Pointer {
	return recv.native
}

// FieldRefCount returns the C field 'ref_count'.
func (recv *DBusMethodInfo) FieldRefCount() int32 {
	argValue := gi.StructFieldGet(dBusMethodInfoStruct, recv.native, "ref_count")
	value := argValue.Int32()
	return value
}

// SetFieldRefCount sets the value of the C field 'ref_count'.
func (recv *DBusMethodInfo) SetFieldRefCount(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.StructFieldSet(dBusMethodInfoStruct, recv.native, "ref_count", argValue)
}

// FieldName returns the C field 'name'.
func (recv *DBusMethodInfo) FieldName() string {
	argValue := gi.StructFieldGet(dBusMethodInfoStruct, recv.native, "name")
	value := argValue.String(false)
	return value
}

// SetFieldName sets the value of the C field 'name'.
func (recv *DBusMethodInfo) SetFieldName(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.StructFieldSet(dBusMethodInfoStruct, recv.native, "name", argValue)
}

// UNSUPPORTED : C value 'in_args' : for field getter : missing Type

// UNSUPPORTED : C value 'in_args' : for field setter : missing Type

// UNSUPPORTED : C value 'out_args' : for field getter : missing Type

// UNSUPPORTED : C value 'out_args' : for field setter : missing Type

// UNSUPPORTED : C value 'annotations' : for field getter : missing Type

// UNSUPPORTED : C value 'annotations' : for field setter : missing Type

var dBusMethodInfoRefFunction *gi.Function
var dBusMethodInfoRefFunction_Once sync.Once

func dBusMethodInfoRefFunction_Set() error {
	var err error
	dBusMethodInfoRefFunction_Once.Do(func() {
		err = dBusMethodInfoStruct_Set()
		if err != nil {
			return
		}
		dBusMethodInfoRefFunction, err = dBusMethodInfoStruct.InvokerNew("ref")
	})
	return err
}

// Ref is a representation of the C type g_dbus_method_info_ref.
func (recv *DBusMethodInfo) Ref() *DBusMethodInfo {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := dBusMethodInfoRefFunction_Set()
	if err == nil {
		ret = dBusMethodInfoRefFunction.Invoke(inArgs[:], nil)
	}

	retGo := DBusMethodInfoNewFromNative(ret.Pointer())

	return retGo
}

var dBusMethodInfoUnrefFunction *gi.Function
var dBusMethodInfoUnrefFunction_Once sync.Once

func dBusMethodInfoUnrefFunction_Set() error {
	var err error
	dBusMethodInfoUnrefFunction_Once.Do(func() {
		err = dBusMethodInfoStruct_Set()
		if err != nil {
			return
		}
		dBusMethodInfoUnrefFunction, err = dBusMethodInfoStruct.InvokerNew("unref")
	})
	return err
}

// Unref is a representation of the C type g_dbus_method_info_unref.
func (recv *DBusMethodInfo) Unref() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := dBusMethodInfoUnrefFunction_Set()
	if err == nil {
		dBusMethodInfoUnrefFunction.Invoke(inArgs[:], nil)
	}

	return
}

// DBusMethodInfoStruct creates an uninitialised DBusMethodInfo.
func DBusMethodInfoStruct() *DBusMethodInfo {
	err := dBusMethodInfoStruct_Set()
	if err != nil {
		return nil
	}

	structGo := DBusMethodInfoNewFromNative(dBusMethodInfoStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeDBusMethodInfo)
	return structGo
}
func finalizeDBusMethodInfo(obj *DBusMethodInfo) {
	dBusMethodInfoStruct.Free(obj.native)
}

var dBusNodeInfoStruct *gi.Struct
var dBusNodeInfoStruct_Once sync.Once

func dBusNodeInfoStruct_Set() error {
	var err error
	dBusNodeInfoStruct_Once.Do(func() {
		dBusNodeInfoStruct, err = gi.StructNew("Gio", "DBusNodeInfo")
	})
	return err
}

type DBusNodeInfo struct {
	native unsafe.Pointer
}

func DBusNodeInfoNewFromNative(native unsafe.Pointer) *DBusNodeInfo {
	instance := &DBusNodeInfo{native: native}

	return instance
}

/*
CastToDBusNodeInfo down casts any arbitrary Object to DBusNodeInfo.
Exercise care, as this is a potentially dangerous function
if the Object is not a DBusNodeInfo.
*/
func (recv *DBusNodeInfo) CastToDBusNodeInfo(object *gobject.Object) *DBusNodeInfo {
	return DBusNodeInfoNewFromNative(object.Native())
}

func (recv *DBusNodeInfo) Native() unsafe.Pointer {
	return recv.native
}

// FieldRefCount returns the C field 'ref_count'.
func (recv *DBusNodeInfo) FieldRefCount() int32 {
	argValue := gi.StructFieldGet(dBusNodeInfoStruct, recv.native, "ref_count")
	value := argValue.Int32()
	return value
}

// SetFieldRefCount sets the value of the C field 'ref_count'.
func (recv *DBusNodeInfo) SetFieldRefCount(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.StructFieldSet(dBusNodeInfoStruct, recv.native, "ref_count", argValue)
}

// FieldPath returns the C field 'path'.
func (recv *DBusNodeInfo) FieldPath() string {
	argValue := gi.StructFieldGet(dBusNodeInfoStruct, recv.native, "path")
	value := argValue.String(false)
	return value
}

// SetFieldPath sets the value of the C field 'path'.
func (recv *DBusNodeInfo) SetFieldPath(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.StructFieldSet(dBusNodeInfoStruct, recv.native, "path", argValue)
}

// UNSUPPORTED : C value 'interfaces' : for field getter : missing Type

// UNSUPPORTED : C value 'interfaces' : for field setter : missing Type

// UNSUPPORTED : C value 'nodes' : for field getter : missing Type

// UNSUPPORTED : C value 'nodes' : for field setter : missing Type

// UNSUPPORTED : C value 'annotations' : for field getter : missing Type

// UNSUPPORTED : C value 'annotations' : for field setter : missing Type

var dBusNodeInfoNewForXmlFunction *gi.Function
var dBusNodeInfoNewForXmlFunction_Once sync.Once

func dBusNodeInfoNewForXmlFunction_Set() error {
	var err error
	dBusNodeInfoNewForXmlFunction_Once.Do(func() {
		err = dBusNodeInfoStruct_Set()
		if err != nil {
			return
		}
		dBusNodeInfoNewForXmlFunction, err = dBusNodeInfoStruct.InvokerNew("new_for_xml")
	})
	return err
}

// DBusNodeInfoNewForXml is a representation of the C type g_dbus_node_info_new_for_xml.
func DBusNodeInfoNewForXml(xmlData string) *DBusNodeInfo {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(xmlData)

	var ret gi.Argument

	err := dBusNodeInfoNewForXmlFunction_Set()
	if err == nil {
		ret = dBusNodeInfoNewForXmlFunction.Invoke(inArgs[:], nil)
	}

	retGo := DBusNodeInfoNewFromNative(ret.Pointer())

	return retGo
}

// UNSUPPORTED : C value 'g_dbus_node_info_generate_xml' : parameter 'string_builder' of type 'GLib.String' not supported

var dBusNodeInfoLookupInterfaceFunction *gi.Function
var dBusNodeInfoLookupInterfaceFunction_Once sync.Once

func dBusNodeInfoLookupInterfaceFunction_Set() error {
	var err error
	dBusNodeInfoLookupInterfaceFunction_Once.Do(func() {
		err = dBusNodeInfoStruct_Set()
		if err != nil {
			return
		}
		dBusNodeInfoLookupInterfaceFunction, err = dBusNodeInfoStruct.InvokerNew("lookup_interface")
	})
	return err
}

// LookupInterface is a representation of the C type g_dbus_node_info_lookup_interface.
func (recv *DBusNodeInfo) LookupInterface(name string) *DBusInterfaceInfo {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(name)

	var ret gi.Argument

	err := dBusNodeInfoLookupInterfaceFunction_Set()
	if err == nil {
		ret = dBusNodeInfoLookupInterfaceFunction.Invoke(inArgs[:], nil)
	}

	retGo := DBusInterfaceInfoNewFromNative(ret.Pointer())

	return retGo
}

var dBusNodeInfoRefFunction *gi.Function
var dBusNodeInfoRefFunction_Once sync.Once

func dBusNodeInfoRefFunction_Set() error {
	var err error
	dBusNodeInfoRefFunction_Once.Do(func() {
		err = dBusNodeInfoStruct_Set()
		if err != nil {
			return
		}
		dBusNodeInfoRefFunction, err = dBusNodeInfoStruct.InvokerNew("ref")
	})
	return err
}

// Ref is a representation of the C type g_dbus_node_info_ref.
func (recv *DBusNodeInfo) Ref() *DBusNodeInfo {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := dBusNodeInfoRefFunction_Set()
	if err == nil {
		ret = dBusNodeInfoRefFunction.Invoke(inArgs[:], nil)
	}

	retGo := DBusNodeInfoNewFromNative(ret.Pointer())

	return retGo
}

var dBusNodeInfoUnrefFunction *gi.Function
var dBusNodeInfoUnrefFunction_Once sync.Once

func dBusNodeInfoUnrefFunction_Set() error {
	var err error
	dBusNodeInfoUnrefFunction_Once.Do(func() {
		err = dBusNodeInfoStruct_Set()
		if err != nil {
			return
		}
		dBusNodeInfoUnrefFunction, err = dBusNodeInfoStruct.InvokerNew("unref")
	})
	return err
}

// Unref is a representation of the C type g_dbus_node_info_unref.
func (recv *DBusNodeInfo) Unref() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := dBusNodeInfoUnrefFunction_Set()
	if err == nil {
		dBusNodeInfoUnrefFunction.Invoke(inArgs[:], nil)
	}

	return
}

var dBusObjectIfaceStruct *gi.Struct
var dBusObjectIfaceStruct_Once sync.Once

func dBusObjectIfaceStruct_Set() error {
	var err error
	dBusObjectIfaceStruct_Once.Do(func() {
		dBusObjectIfaceStruct, err = gi.StructNew("Gio", "DBusObjectIface")
	})
	return err
}

type DBusObjectIface struct {
	native unsafe.Pointer
}

func DBusObjectIfaceNewFromNative(native unsafe.Pointer) *DBusObjectIface {
	instance := &DBusObjectIface{native: native}

	return instance
}

/*
CastToDBusObjectIface down casts any arbitrary Object to DBusObjectIface.
Exercise care, as this is a potentially dangerous function
if the Object is not a DBusObjectIface.
*/
func (recv *DBusObjectIface) CastToDBusObjectIface(object *gobject.Object) *DBusObjectIface {
	return DBusObjectIfaceNewFromNative(object.Native())
}

func (recv *DBusObjectIface) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'parent_iface' : for field getter : no Go type for 'GObject.TypeInterface'

// UNSUPPORTED : C value 'parent_iface' : for field setter : no Go type for 'GObject.TypeInterface'

// UNSUPPORTED : C value 'get_object_path' : for field getter : missing Type

// UNSUPPORTED : C value 'get_object_path' : for field setter : missing Type

// UNSUPPORTED : C value 'get_interfaces' : for field getter : missing Type

// UNSUPPORTED : C value 'get_interfaces' : for field setter : missing Type

// UNSUPPORTED : C value 'get_interface' : for field getter : missing Type

// UNSUPPORTED : C value 'get_interface' : for field setter : missing Type

// UNSUPPORTED : C value 'interface_added' : for field getter : missing Type

// UNSUPPORTED : C value 'interface_added' : for field setter : missing Type

// UNSUPPORTED : C value 'interface_removed' : for field getter : missing Type

// UNSUPPORTED : C value 'interface_removed' : for field setter : missing Type

// DBusObjectIfaceStruct creates an uninitialised DBusObjectIface.
func DBusObjectIfaceStruct() *DBusObjectIface {
	err := dBusObjectIfaceStruct_Set()
	if err != nil {
		return nil
	}

	structGo := DBusObjectIfaceNewFromNative(dBusObjectIfaceStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeDBusObjectIface)
	return structGo
}
func finalizeDBusObjectIface(obj *DBusObjectIface) {
	dBusObjectIfaceStruct.Free(obj.native)
}

var dBusObjectManagerClientClassStruct *gi.Struct
var dBusObjectManagerClientClassStruct_Once sync.Once

func dBusObjectManagerClientClassStruct_Set() error {
	var err error
	dBusObjectManagerClientClassStruct_Once.Do(func() {
		dBusObjectManagerClientClassStruct, err = gi.StructNew("Gio", "DBusObjectManagerClientClass")
	})
	return err
}

type DBusObjectManagerClientClass struct {
	native unsafe.Pointer
}

func DBusObjectManagerClientClassNewFromNative(native unsafe.Pointer) *DBusObjectManagerClientClass {
	instance := &DBusObjectManagerClientClass{native: native}

	return instance
}

/*
CastToDBusObjectManagerClientClass down casts any arbitrary Object to DBusObjectManagerClientClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a DBusObjectManagerClientClass.
*/
func (recv *DBusObjectManagerClientClass) CastToDBusObjectManagerClientClass(object *gobject.Object) *DBusObjectManagerClientClass {
	return DBusObjectManagerClientClassNewFromNative(object.Native())
}

func (recv *DBusObjectManagerClientClass) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'interface_proxy_signal' : for field getter : missing Type

// UNSUPPORTED : C value 'interface_proxy_signal' : for field setter : missing Type

// UNSUPPORTED : C value 'interface_proxy_properties_changed' : for field getter : missing Type

// UNSUPPORTED : C value 'interface_proxy_properties_changed' : for field setter : missing Type

// DBusObjectManagerClientClassStruct creates an uninitialised DBusObjectManagerClientClass.
func DBusObjectManagerClientClassStruct() *DBusObjectManagerClientClass {
	err := dBusObjectManagerClientClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := DBusObjectManagerClientClassNewFromNative(dBusObjectManagerClientClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeDBusObjectManagerClientClass)
	return structGo
}
func finalizeDBusObjectManagerClientClass(obj *DBusObjectManagerClientClass) {
	dBusObjectManagerClientClassStruct.Free(obj.native)
}

var dBusObjectManagerClientPrivateStruct *gi.Struct
var dBusObjectManagerClientPrivateStruct_Once sync.Once

func dBusObjectManagerClientPrivateStruct_Set() error {
	var err error
	dBusObjectManagerClientPrivateStruct_Once.Do(func() {
		dBusObjectManagerClientPrivateStruct, err = gi.StructNew("Gio", "DBusObjectManagerClientPrivate")
	})
	return err
}

type DBusObjectManagerClientPrivate struct {
	native unsafe.Pointer
}

func DBusObjectManagerClientPrivateNewFromNative(native unsafe.Pointer) *DBusObjectManagerClientPrivate {
	instance := &DBusObjectManagerClientPrivate{native: native}

	return instance
}

/*
CastToDBusObjectManagerClientPrivate down casts any arbitrary Object to DBusObjectManagerClientPrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a DBusObjectManagerClientPrivate.
*/
func (recv *DBusObjectManagerClientPrivate) CastToDBusObjectManagerClientPrivate(object *gobject.Object) *DBusObjectManagerClientPrivate {
	return DBusObjectManagerClientPrivateNewFromNative(object.Native())
}

func (recv *DBusObjectManagerClientPrivate) Native() unsafe.Pointer {
	return recv.native
}

// DBusObjectManagerClientPrivateStruct creates an uninitialised DBusObjectManagerClientPrivate.
func DBusObjectManagerClientPrivateStruct() *DBusObjectManagerClientPrivate {
	err := dBusObjectManagerClientPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := DBusObjectManagerClientPrivateNewFromNative(dBusObjectManagerClientPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeDBusObjectManagerClientPrivate)
	return structGo
}
func finalizeDBusObjectManagerClientPrivate(obj *DBusObjectManagerClientPrivate) {
	dBusObjectManagerClientPrivateStruct.Free(obj.native)
}

var dBusObjectManagerIfaceStruct *gi.Struct
var dBusObjectManagerIfaceStruct_Once sync.Once

func dBusObjectManagerIfaceStruct_Set() error {
	var err error
	dBusObjectManagerIfaceStruct_Once.Do(func() {
		dBusObjectManagerIfaceStruct, err = gi.StructNew("Gio", "DBusObjectManagerIface")
	})
	return err
}

type DBusObjectManagerIface struct {
	native unsafe.Pointer
}

func DBusObjectManagerIfaceNewFromNative(native unsafe.Pointer) *DBusObjectManagerIface {
	instance := &DBusObjectManagerIface{native: native}

	return instance
}

/*
CastToDBusObjectManagerIface down casts any arbitrary Object to DBusObjectManagerIface.
Exercise care, as this is a potentially dangerous function
if the Object is not a DBusObjectManagerIface.
*/
func (recv *DBusObjectManagerIface) CastToDBusObjectManagerIface(object *gobject.Object) *DBusObjectManagerIface {
	return DBusObjectManagerIfaceNewFromNative(object.Native())
}

func (recv *DBusObjectManagerIface) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'parent_iface' : for field getter : no Go type for 'GObject.TypeInterface'

// UNSUPPORTED : C value 'parent_iface' : for field setter : no Go type for 'GObject.TypeInterface'

// UNSUPPORTED : C value 'get_object_path' : for field getter : missing Type

// UNSUPPORTED : C value 'get_object_path' : for field setter : missing Type

// UNSUPPORTED : C value 'get_objects' : for field getter : missing Type

// UNSUPPORTED : C value 'get_objects' : for field setter : missing Type

// UNSUPPORTED : C value 'get_object' : for field getter : missing Type

// UNSUPPORTED : C value 'get_object' : for field setter : missing Type

// UNSUPPORTED : C value 'get_interface' : for field getter : missing Type

// UNSUPPORTED : C value 'get_interface' : for field setter : missing Type

// UNSUPPORTED : C value 'object_added' : for field getter : missing Type

// UNSUPPORTED : C value 'object_added' : for field setter : missing Type

// UNSUPPORTED : C value 'object_removed' : for field getter : missing Type

// UNSUPPORTED : C value 'object_removed' : for field setter : missing Type

// UNSUPPORTED : C value 'interface_added' : for field getter : missing Type

// UNSUPPORTED : C value 'interface_added' : for field setter : missing Type

// UNSUPPORTED : C value 'interface_removed' : for field getter : missing Type

// UNSUPPORTED : C value 'interface_removed' : for field setter : missing Type

// DBusObjectManagerIfaceStruct creates an uninitialised DBusObjectManagerIface.
func DBusObjectManagerIfaceStruct() *DBusObjectManagerIface {
	err := dBusObjectManagerIfaceStruct_Set()
	if err != nil {
		return nil
	}

	structGo := DBusObjectManagerIfaceNewFromNative(dBusObjectManagerIfaceStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeDBusObjectManagerIface)
	return structGo
}
func finalizeDBusObjectManagerIface(obj *DBusObjectManagerIface) {
	dBusObjectManagerIfaceStruct.Free(obj.native)
}

var dBusObjectManagerServerClassStruct *gi.Struct
var dBusObjectManagerServerClassStruct_Once sync.Once

func dBusObjectManagerServerClassStruct_Set() error {
	var err error
	dBusObjectManagerServerClassStruct_Once.Do(func() {
		dBusObjectManagerServerClassStruct, err = gi.StructNew("Gio", "DBusObjectManagerServerClass")
	})
	return err
}

type DBusObjectManagerServerClass struct {
	native unsafe.Pointer
}

func DBusObjectManagerServerClassNewFromNative(native unsafe.Pointer) *DBusObjectManagerServerClass {
	instance := &DBusObjectManagerServerClass{native: native}

	return instance
}

/*
CastToDBusObjectManagerServerClass down casts any arbitrary Object to DBusObjectManagerServerClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a DBusObjectManagerServerClass.
*/
func (recv *DBusObjectManagerServerClass) CastToDBusObjectManagerServerClass(object *gobject.Object) *DBusObjectManagerServerClass {
	return DBusObjectManagerServerClassNewFromNative(object.Native())
}

func (recv *DBusObjectManagerServerClass) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// DBusObjectManagerServerClassStruct creates an uninitialised DBusObjectManagerServerClass.
func DBusObjectManagerServerClassStruct() *DBusObjectManagerServerClass {
	err := dBusObjectManagerServerClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := DBusObjectManagerServerClassNewFromNative(dBusObjectManagerServerClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeDBusObjectManagerServerClass)
	return structGo
}
func finalizeDBusObjectManagerServerClass(obj *DBusObjectManagerServerClass) {
	dBusObjectManagerServerClassStruct.Free(obj.native)
}

var dBusObjectManagerServerPrivateStruct *gi.Struct
var dBusObjectManagerServerPrivateStruct_Once sync.Once

func dBusObjectManagerServerPrivateStruct_Set() error {
	var err error
	dBusObjectManagerServerPrivateStruct_Once.Do(func() {
		dBusObjectManagerServerPrivateStruct, err = gi.StructNew("Gio", "DBusObjectManagerServerPrivate")
	})
	return err
}

type DBusObjectManagerServerPrivate struct {
	native unsafe.Pointer
}

func DBusObjectManagerServerPrivateNewFromNative(native unsafe.Pointer) *DBusObjectManagerServerPrivate {
	instance := &DBusObjectManagerServerPrivate{native: native}

	return instance
}

/*
CastToDBusObjectManagerServerPrivate down casts any arbitrary Object to DBusObjectManagerServerPrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a DBusObjectManagerServerPrivate.
*/
func (recv *DBusObjectManagerServerPrivate) CastToDBusObjectManagerServerPrivate(object *gobject.Object) *DBusObjectManagerServerPrivate {
	return DBusObjectManagerServerPrivateNewFromNative(object.Native())
}

func (recv *DBusObjectManagerServerPrivate) Native() unsafe.Pointer {
	return recv.native
}

// DBusObjectManagerServerPrivateStruct creates an uninitialised DBusObjectManagerServerPrivate.
func DBusObjectManagerServerPrivateStruct() *DBusObjectManagerServerPrivate {
	err := dBusObjectManagerServerPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := DBusObjectManagerServerPrivateNewFromNative(dBusObjectManagerServerPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeDBusObjectManagerServerPrivate)
	return structGo
}
func finalizeDBusObjectManagerServerPrivate(obj *DBusObjectManagerServerPrivate) {
	dBusObjectManagerServerPrivateStruct.Free(obj.native)
}

var dBusObjectProxyClassStruct *gi.Struct
var dBusObjectProxyClassStruct_Once sync.Once

func dBusObjectProxyClassStruct_Set() error {
	var err error
	dBusObjectProxyClassStruct_Once.Do(func() {
		dBusObjectProxyClassStruct, err = gi.StructNew("Gio", "DBusObjectProxyClass")
	})
	return err
}

type DBusObjectProxyClass struct {
	native unsafe.Pointer
}

func DBusObjectProxyClassNewFromNative(native unsafe.Pointer) *DBusObjectProxyClass {
	instance := &DBusObjectProxyClass{native: native}

	return instance
}

/*
CastToDBusObjectProxyClass down casts any arbitrary Object to DBusObjectProxyClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a DBusObjectProxyClass.
*/
func (recv *DBusObjectProxyClass) CastToDBusObjectProxyClass(object *gobject.Object) *DBusObjectProxyClass {
	return DBusObjectProxyClassNewFromNative(object.Native())
}

func (recv *DBusObjectProxyClass) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// DBusObjectProxyClassStruct creates an uninitialised DBusObjectProxyClass.
func DBusObjectProxyClassStruct() *DBusObjectProxyClass {
	err := dBusObjectProxyClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := DBusObjectProxyClassNewFromNative(dBusObjectProxyClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeDBusObjectProxyClass)
	return structGo
}
func finalizeDBusObjectProxyClass(obj *DBusObjectProxyClass) {
	dBusObjectProxyClassStruct.Free(obj.native)
}

var dBusObjectProxyPrivateStruct *gi.Struct
var dBusObjectProxyPrivateStruct_Once sync.Once

func dBusObjectProxyPrivateStruct_Set() error {
	var err error
	dBusObjectProxyPrivateStruct_Once.Do(func() {
		dBusObjectProxyPrivateStruct, err = gi.StructNew("Gio", "DBusObjectProxyPrivate")
	})
	return err
}

type DBusObjectProxyPrivate struct {
	native unsafe.Pointer
}

func DBusObjectProxyPrivateNewFromNative(native unsafe.Pointer) *DBusObjectProxyPrivate {
	instance := &DBusObjectProxyPrivate{native: native}

	return instance
}

/*
CastToDBusObjectProxyPrivate down casts any arbitrary Object to DBusObjectProxyPrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a DBusObjectProxyPrivate.
*/
func (recv *DBusObjectProxyPrivate) CastToDBusObjectProxyPrivate(object *gobject.Object) *DBusObjectProxyPrivate {
	return DBusObjectProxyPrivateNewFromNative(object.Native())
}

func (recv *DBusObjectProxyPrivate) Native() unsafe.Pointer {
	return recv.native
}

// DBusObjectProxyPrivateStruct creates an uninitialised DBusObjectProxyPrivate.
func DBusObjectProxyPrivateStruct() *DBusObjectProxyPrivate {
	err := dBusObjectProxyPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := DBusObjectProxyPrivateNewFromNative(dBusObjectProxyPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeDBusObjectProxyPrivate)
	return structGo
}
func finalizeDBusObjectProxyPrivate(obj *DBusObjectProxyPrivate) {
	dBusObjectProxyPrivateStruct.Free(obj.native)
}

var dBusObjectSkeletonClassStruct *gi.Struct
var dBusObjectSkeletonClassStruct_Once sync.Once

func dBusObjectSkeletonClassStruct_Set() error {
	var err error
	dBusObjectSkeletonClassStruct_Once.Do(func() {
		dBusObjectSkeletonClassStruct, err = gi.StructNew("Gio", "DBusObjectSkeletonClass")
	})
	return err
}

type DBusObjectSkeletonClass struct {
	native unsafe.Pointer
}

func DBusObjectSkeletonClassNewFromNative(native unsafe.Pointer) *DBusObjectSkeletonClass {
	instance := &DBusObjectSkeletonClass{native: native}

	return instance
}

/*
CastToDBusObjectSkeletonClass down casts any arbitrary Object to DBusObjectSkeletonClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a DBusObjectSkeletonClass.
*/
func (recv *DBusObjectSkeletonClass) CastToDBusObjectSkeletonClass(object *gobject.Object) *DBusObjectSkeletonClass {
	return DBusObjectSkeletonClassNewFromNative(object.Native())
}

func (recv *DBusObjectSkeletonClass) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'authorize_method' : for field getter : missing Type

// UNSUPPORTED : C value 'authorize_method' : for field setter : missing Type

// DBusObjectSkeletonClassStruct creates an uninitialised DBusObjectSkeletonClass.
func DBusObjectSkeletonClassStruct() *DBusObjectSkeletonClass {
	err := dBusObjectSkeletonClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := DBusObjectSkeletonClassNewFromNative(dBusObjectSkeletonClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeDBusObjectSkeletonClass)
	return structGo
}
func finalizeDBusObjectSkeletonClass(obj *DBusObjectSkeletonClass) {
	dBusObjectSkeletonClassStruct.Free(obj.native)
}

var dBusObjectSkeletonPrivateStruct *gi.Struct
var dBusObjectSkeletonPrivateStruct_Once sync.Once

func dBusObjectSkeletonPrivateStruct_Set() error {
	var err error
	dBusObjectSkeletonPrivateStruct_Once.Do(func() {
		dBusObjectSkeletonPrivateStruct, err = gi.StructNew("Gio", "DBusObjectSkeletonPrivate")
	})
	return err
}

type DBusObjectSkeletonPrivate struct {
	native unsafe.Pointer
}

func DBusObjectSkeletonPrivateNewFromNative(native unsafe.Pointer) *DBusObjectSkeletonPrivate {
	instance := &DBusObjectSkeletonPrivate{native: native}

	return instance
}

/*
CastToDBusObjectSkeletonPrivate down casts any arbitrary Object to DBusObjectSkeletonPrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a DBusObjectSkeletonPrivate.
*/
func (recv *DBusObjectSkeletonPrivate) CastToDBusObjectSkeletonPrivate(object *gobject.Object) *DBusObjectSkeletonPrivate {
	return DBusObjectSkeletonPrivateNewFromNative(object.Native())
}

func (recv *DBusObjectSkeletonPrivate) Native() unsafe.Pointer {
	return recv.native
}

// DBusObjectSkeletonPrivateStruct creates an uninitialised DBusObjectSkeletonPrivate.
func DBusObjectSkeletonPrivateStruct() *DBusObjectSkeletonPrivate {
	err := dBusObjectSkeletonPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := DBusObjectSkeletonPrivateNewFromNative(dBusObjectSkeletonPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeDBusObjectSkeletonPrivate)
	return structGo
}
func finalizeDBusObjectSkeletonPrivate(obj *DBusObjectSkeletonPrivate) {
	dBusObjectSkeletonPrivateStruct.Free(obj.native)
}

var dBusPropertyInfoStruct *gi.Struct
var dBusPropertyInfoStruct_Once sync.Once

func dBusPropertyInfoStruct_Set() error {
	var err error
	dBusPropertyInfoStruct_Once.Do(func() {
		dBusPropertyInfoStruct, err = gi.StructNew("Gio", "DBusPropertyInfo")
	})
	return err
}

type DBusPropertyInfo struct {
	native unsafe.Pointer
}

func DBusPropertyInfoNewFromNative(native unsafe.Pointer) *DBusPropertyInfo {
	instance := &DBusPropertyInfo{native: native}

	return instance
}

/*
CastToDBusPropertyInfo down casts any arbitrary Object to DBusPropertyInfo.
Exercise care, as this is a potentially dangerous function
if the Object is not a DBusPropertyInfo.
*/
func (recv *DBusPropertyInfo) CastToDBusPropertyInfo(object *gobject.Object) *DBusPropertyInfo {
	return DBusPropertyInfoNewFromNative(object.Native())
}

func (recv *DBusPropertyInfo) Native() unsafe.Pointer {
	return recv.native
}

// FieldRefCount returns the C field 'ref_count'.
func (recv *DBusPropertyInfo) FieldRefCount() int32 {
	argValue := gi.StructFieldGet(dBusPropertyInfoStruct, recv.native, "ref_count")
	value := argValue.Int32()
	return value
}

// SetFieldRefCount sets the value of the C field 'ref_count'.
func (recv *DBusPropertyInfo) SetFieldRefCount(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.StructFieldSet(dBusPropertyInfoStruct, recv.native, "ref_count", argValue)
}

// FieldName returns the C field 'name'.
func (recv *DBusPropertyInfo) FieldName() string {
	argValue := gi.StructFieldGet(dBusPropertyInfoStruct, recv.native, "name")
	value := argValue.String(false)
	return value
}

// SetFieldName sets the value of the C field 'name'.
func (recv *DBusPropertyInfo) SetFieldName(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.StructFieldSet(dBusPropertyInfoStruct, recv.native, "name", argValue)
}

// FieldSignature returns the C field 'signature'.
func (recv *DBusPropertyInfo) FieldSignature() string {
	argValue := gi.StructFieldGet(dBusPropertyInfoStruct, recv.native, "signature")
	value := argValue.String(false)
	return value
}

// SetFieldSignature sets the value of the C field 'signature'.
func (recv *DBusPropertyInfo) SetFieldSignature(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.StructFieldSet(dBusPropertyInfoStruct, recv.native, "signature", argValue)
}

// UNSUPPORTED : C value 'flags' : for field getter : no Go type for 'DBusPropertyInfoFlags'

// UNSUPPORTED : C value 'flags' : for field setter : no Go type for 'DBusPropertyInfoFlags'

// UNSUPPORTED : C value 'annotations' : for field getter : missing Type

// UNSUPPORTED : C value 'annotations' : for field setter : missing Type

var dBusPropertyInfoRefFunction *gi.Function
var dBusPropertyInfoRefFunction_Once sync.Once

func dBusPropertyInfoRefFunction_Set() error {
	var err error
	dBusPropertyInfoRefFunction_Once.Do(func() {
		err = dBusPropertyInfoStruct_Set()
		if err != nil {
			return
		}
		dBusPropertyInfoRefFunction, err = dBusPropertyInfoStruct.InvokerNew("ref")
	})
	return err
}

// Ref is a representation of the C type g_dbus_property_info_ref.
func (recv *DBusPropertyInfo) Ref() *DBusPropertyInfo {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := dBusPropertyInfoRefFunction_Set()
	if err == nil {
		ret = dBusPropertyInfoRefFunction.Invoke(inArgs[:], nil)
	}

	retGo := DBusPropertyInfoNewFromNative(ret.Pointer())

	return retGo
}

var dBusPropertyInfoUnrefFunction *gi.Function
var dBusPropertyInfoUnrefFunction_Once sync.Once

func dBusPropertyInfoUnrefFunction_Set() error {
	var err error
	dBusPropertyInfoUnrefFunction_Once.Do(func() {
		err = dBusPropertyInfoStruct_Set()
		if err != nil {
			return
		}
		dBusPropertyInfoUnrefFunction, err = dBusPropertyInfoStruct.InvokerNew("unref")
	})
	return err
}

// Unref is a representation of the C type g_dbus_property_info_unref.
func (recv *DBusPropertyInfo) Unref() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := dBusPropertyInfoUnrefFunction_Set()
	if err == nil {
		dBusPropertyInfoUnrefFunction.Invoke(inArgs[:], nil)
	}

	return
}

// DBusPropertyInfoStruct creates an uninitialised DBusPropertyInfo.
func DBusPropertyInfoStruct() *DBusPropertyInfo {
	err := dBusPropertyInfoStruct_Set()
	if err != nil {
		return nil
	}

	structGo := DBusPropertyInfoNewFromNative(dBusPropertyInfoStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeDBusPropertyInfo)
	return structGo
}
func finalizeDBusPropertyInfo(obj *DBusPropertyInfo) {
	dBusPropertyInfoStruct.Free(obj.native)
}

var dBusProxyClassStruct *gi.Struct
var dBusProxyClassStruct_Once sync.Once

func dBusProxyClassStruct_Set() error {
	var err error
	dBusProxyClassStruct_Once.Do(func() {
		dBusProxyClassStruct, err = gi.StructNew("Gio", "DBusProxyClass")
	})
	return err
}

type DBusProxyClass struct {
	native unsafe.Pointer
}

func DBusProxyClassNewFromNative(native unsafe.Pointer) *DBusProxyClass {
	instance := &DBusProxyClass{native: native}

	return instance
}

/*
CastToDBusProxyClass down casts any arbitrary Object to DBusProxyClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a DBusProxyClass.
*/
func (recv *DBusProxyClass) CastToDBusProxyClass(object *gobject.Object) *DBusProxyClass {
	return DBusProxyClassNewFromNative(object.Native())
}

func (recv *DBusProxyClass) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'g_properties_changed' : for field getter : missing Type

// UNSUPPORTED : C value 'g_properties_changed' : for field setter : missing Type

// UNSUPPORTED : C value 'g_signal' : for field getter : missing Type

// UNSUPPORTED : C value 'g_signal' : for field setter : missing Type

// DBusProxyClassStruct creates an uninitialised DBusProxyClass.
func DBusProxyClassStruct() *DBusProxyClass {
	err := dBusProxyClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := DBusProxyClassNewFromNative(dBusProxyClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeDBusProxyClass)
	return structGo
}
func finalizeDBusProxyClass(obj *DBusProxyClass) {
	dBusProxyClassStruct.Free(obj.native)
}

var dBusProxyPrivateStruct *gi.Struct
var dBusProxyPrivateStruct_Once sync.Once

func dBusProxyPrivateStruct_Set() error {
	var err error
	dBusProxyPrivateStruct_Once.Do(func() {
		dBusProxyPrivateStruct, err = gi.StructNew("Gio", "DBusProxyPrivate")
	})
	return err
}

type DBusProxyPrivate struct {
	native unsafe.Pointer
}

func DBusProxyPrivateNewFromNative(native unsafe.Pointer) *DBusProxyPrivate {
	instance := &DBusProxyPrivate{native: native}

	return instance
}

/*
CastToDBusProxyPrivate down casts any arbitrary Object to DBusProxyPrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a DBusProxyPrivate.
*/
func (recv *DBusProxyPrivate) CastToDBusProxyPrivate(object *gobject.Object) *DBusProxyPrivate {
	return DBusProxyPrivateNewFromNative(object.Native())
}

func (recv *DBusProxyPrivate) Native() unsafe.Pointer {
	return recv.native
}

// DBusProxyPrivateStruct creates an uninitialised DBusProxyPrivate.
func DBusProxyPrivateStruct() *DBusProxyPrivate {
	err := dBusProxyPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := DBusProxyPrivateNewFromNative(dBusProxyPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeDBusProxyPrivate)
	return structGo
}
func finalizeDBusProxyPrivate(obj *DBusProxyPrivate) {
	dBusProxyPrivateStruct.Free(obj.native)
}

var dBusSignalInfoStruct *gi.Struct
var dBusSignalInfoStruct_Once sync.Once

func dBusSignalInfoStruct_Set() error {
	var err error
	dBusSignalInfoStruct_Once.Do(func() {
		dBusSignalInfoStruct, err = gi.StructNew("Gio", "DBusSignalInfo")
	})
	return err
}

type DBusSignalInfo struct {
	native unsafe.Pointer
}

func DBusSignalInfoNewFromNative(native unsafe.Pointer) *DBusSignalInfo {
	instance := &DBusSignalInfo{native: native}

	return instance
}

/*
CastToDBusSignalInfo down casts any arbitrary Object to DBusSignalInfo.
Exercise care, as this is a potentially dangerous function
if the Object is not a DBusSignalInfo.
*/
func (recv *DBusSignalInfo) CastToDBusSignalInfo(object *gobject.Object) *DBusSignalInfo {
	return DBusSignalInfoNewFromNative(object.Native())
}

func (recv *DBusSignalInfo) Native() unsafe.Pointer {
	return recv.native
}

// FieldRefCount returns the C field 'ref_count'.
func (recv *DBusSignalInfo) FieldRefCount() int32 {
	argValue := gi.StructFieldGet(dBusSignalInfoStruct, recv.native, "ref_count")
	value := argValue.Int32()
	return value
}

// SetFieldRefCount sets the value of the C field 'ref_count'.
func (recv *DBusSignalInfo) SetFieldRefCount(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.StructFieldSet(dBusSignalInfoStruct, recv.native, "ref_count", argValue)
}

// FieldName returns the C field 'name'.
func (recv *DBusSignalInfo) FieldName() string {
	argValue := gi.StructFieldGet(dBusSignalInfoStruct, recv.native, "name")
	value := argValue.String(false)
	return value
}

// SetFieldName sets the value of the C field 'name'.
func (recv *DBusSignalInfo) SetFieldName(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.StructFieldSet(dBusSignalInfoStruct, recv.native, "name", argValue)
}

// UNSUPPORTED : C value 'args' : for field getter : missing Type

// UNSUPPORTED : C value 'args' : for field setter : missing Type

// UNSUPPORTED : C value 'annotations' : for field getter : missing Type

// UNSUPPORTED : C value 'annotations' : for field setter : missing Type

var dBusSignalInfoRefFunction *gi.Function
var dBusSignalInfoRefFunction_Once sync.Once

func dBusSignalInfoRefFunction_Set() error {
	var err error
	dBusSignalInfoRefFunction_Once.Do(func() {
		err = dBusSignalInfoStruct_Set()
		if err != nil {
			return
		}
		dBusSignalInfoRefFunction, err = dBusSignalInfoStruct.InvokerNew("ref")
	})
	return err
}

// Ref is a representation of the C type g_dbus_signal_info_ref.
func (recv *DBusSignalInfo) Ref() *DBusSignalInfo {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := dBusSignalInfoRefFunction_Set()
	if err == nil {
		ret = dBusSignalInfoRefFunction.Invoke(inArgs[:], nil)
	}

	retGo := DBusSignalInfoNewFromNative(ret.Pointer())

	return retGo
}

var dBusSignalInfoUnrefFunction *gi.Function
var dBusSignalInfoUnrefFunction_Once sync.Once

func dBusSignalInfoUnrefFunction_Set() error {
	var err error
	dBusSignalInfoUnrefFunction_Once.Do(func() {
		err = dBusSignalInfoStruct_Set()
		if err != nil {
			return
		}
		dBusSignalInfoUnrefFunction, err = dBusSignalInfoStruct.InvokerNew("unref")
	})
	return err
}

// Unref is a representation of the C type g_dbus_signal_info_unref.
func (recv *DBusSignalInfo) Unref() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := dBusSignalInfoUnrefFunction_Set()
	if err == nil {
		dBusSignalInfoUnrefFunction.Invoke(inArgs[:], nil)
	}

	return
}

// DBusSignalInfoStruct creates an uninitialised DBusSignalInfo.
func DBusSignalInfoStruct() *DBusSignalInfo {
	err := dBusSignalInfoStruct_Set()
	if err != nil {
		return nil
	}

	structGo := DBusSignalInfoNewFromNative(dBusSignalInfoStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeDBusSignalInfo)
	return structGo
}
func finalizeDBusSignalInfo(obj *DBusSignalInfo) {
	dBusSignalInfoStruct.Free(obj.native)
}

var dBusSubtreeVTableStruct *gi.Struct
var dBusSubtreeVTableStruct_Once sync.Once

func dBusSubtreeVTableStruct_Set() error {
	var err error
	dBusSubtreeVTableStruct_Once.Do(func() {
		dBusSubtreeVTableStruct, err = gi.StructNew("Gio", "DBusSubtreeVTable")
	})
	return err
}

type DBusSubtreeVTable struct {
	native unsafe.Pointer
}

func DBusSubtreeVTableNewFromNative(native unsafe.Pointer) *DBusSubtreeVTable {
	instance := &DBusSubtreeVTable{native: native}

	return instance
}

/*
CastToDBusSubtreeVTable down casts any arbitrary Object to DBusSubtreeVTable.
Exercise care, as this is a potentially dangerous function
if the Object is not a DBusSubtreeVTable.
*/
func (recv *DBusSubtreeVTable) CastToDBusSubtreeVTable(object *gobject.Object) *DBusSubtreeVTable {
	return DBusSubtreeVTableNewFromNative(object.Native())
}

func (recv *DBusSubtreeVTable) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'enumerate' : for field getter : no Go type for 'DBusSubtreeEnumerateFunc'

// UNSUPPORTED : C value 'enumerate' : for field setter : no Go type for 'DBusSubtreeEnumerateFunc'

// UNSUPPORTED : C value 'introspect' : for field getter : no Go type for 'DBusSubtreeIntrospectFunc'

// UNSUPPORTED : C value 'introspect' : for field setter : no Go type for 'DBusSubtreeIntrospectFunc'

// UNSUPPORTED : C value 'dispatch' : for field getter : no Go type for 'DBusSubtreeDispatchFunc'

// UNSUPPORTED : C value 'dispatch' : for field setter : no Go type for 'DBusSubtreeDispatchFunc'

// DBusSubtreeVTableStruct creates an uninitialised DBusSubtreeVTable.
func DBusSubtreeVTableStruct() *DBusSubtreeVTable {
	err := dBusSubtreeVTableStruct_Set()
	if err != nil {
		return nil
	}

	structGo := DBusSubtreeVTableNewFromNative(dBusSubtreeVTableStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeDBusSubtreeVTable)
	return structGo
}
func finalizeDBusSubtreeVTable(obj *DBusSubtreeVTable) {
	dBusSubtreeVTableStruct.Free(obj.native)
}

var dataInputStreamClassStruct *gi.Struct
var dataInputStreamClassStruct_Once sync.Once

func dataInputStreamClassStruct_Set() error {
	var err error
	dataInputStreamClassStruct_Once.Do(func() {
		dataInputStreamClassStruct, err = gi.StructNew("Gio", "DataInputStreamClass")
	})
	return err
}

type DataInputStreamClass struct {
	native unsafe.Pointer
}

func DataInputStreamClassNewFromNative(native unsafe.Pointer) *DataInputStreamClass {
	instance := &DataInputStreamClass{native: native}

	return instance
}

/*
CastToDataInputStreamClass down casts any arbitrary Object to DataInputStreamClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a DataInputStreamClass.
*/
func (recv *DataInputStreamClass) CastToDataInputStreamClass(object *gobject.Object) *DataInputStreamClass {
	return DataInputStreamClassNewFromNative(object.Native())
}

func (recv *DataInputStreamClass) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentClass returns the C field 'parent_class'.
func (recv *DataInputStreamClass) FieldParentClass() *BufferedInputStreamClass {
	argValue := gi.StructFieldGet(dataInputStreamClassStruct, recv.native, "parent_class")
	value := BufferedInputStreamClassNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *DataInputStreamClass) SetFieldParentClass(value *BufferedInputStreamClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.StructFieldSet(dataInputStreamClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value '_g_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved4' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved5' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved5' : for field setter : missing Type

// DataInputStreamClassStruct creates an uninitialised DataInputStreamClass.
func DataInputStreamClassStruct() *DataInputStreamClass {
	err := dataInputStreamClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := DataInputStreamClassNewFromNative(dataInputStreamClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeDataInputStreamClass)
	return structGo
}
func finalizeDataInputStreamClass(obj *DataInputStreamClass) {
	dataInputStreamClassStruct.Free(obj.native)
}

var dataInputStreamPrivateStruct *gi.Struct
var dataInputStreamPrivateStruct_Once sync.Once

func dataInputStreamPrivateStruct_Set() error {
	var err error
	dataInputStreamPrivateStruct_Once.Do(func() {
		dataInputStreamPrivateStruct, err = gi.StructNew("Gio", "DataInputStreamPrivate")
	})
	return err
}

type DataInputStreamPrivate struct {
	native unsafe.Pointer
}

func DataInputStreamPrivateNewFromNative(native unsafe.Pointer) *DataInputStreamPrivate {
	instance := &DataInputStreamPrivate{native: native}

	return instance
}

/*
CastToDataInputStreamPrivate down casts any arbitrary Object to DataInputStreamPrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a DataInputStreamPrivate.
*/
func (recv *DataInputStreamPrivate) CastToDataInputStreamPrivate(object *gobject.Object) *DataInputStreamPrivate {
	return DataInputStreamPrivateNewFromNative(object.Native())
}

func (recv *DataInputStreamPrivate) Native() unsafe.Pointer {
	return recv.native
}

// DataInputStreamPrivateStruct creates an uninitialised DataInputStreamPrivate.
func DataInputStreamPrivateStruct() *DataInputStreamPrivate {
	err := dataInputStreamPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := DataInputStreamPrivateNewFromNative(dataInputStreamPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeDataInputStreamPrivate)
	return structGo
}
func finalizeDataInputStreamPrivate(obj *DataInputStreamPrivate) {
	dataInputStreamPrivateStruct.Free(obj.native)
}

var dataOutputStreamClassStruct *gi.Struct
var dataOutputStreamClassStruct_Once sync.Once

func dataOutputStreamClassStruct_Set() error {
	var err error
	dataOutputStreamClassStruct_Once.Do(func() {
		dataOutputStreamClassStruct, err = gi.StructNew("Gio", "DataOutputStreamClass")
	})
	return err
}

type DataOutputStreamClass struct {
	native unsafe.Pointer
}

func DataOutputStreamClassNewFromNative(native unsafe.Pointer) *DataOutputStreamClass {
	instance := &DataOutputStreamClass{native: native}

	return instance
}

/*
CastToDataOutputStreamClass down casts any arbitrary Object to DataOutputStreamClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a DataOutputStreamClass.
*/
func (recv *DataOutputStreamClass) CastToDataOutputStreamClass(object *gobject.Object) *DataOutputStreamClass {
	return DataOutputStreamClassNewFromNative(object.Native())
}

func (recv *DataOutputStreamClass) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentClass returns the C field 'parent_class'.
func (recv *DataOutputStreamClass) FieldParentClass() *FilterOutputStreamClass {
	argValue := gi.StructFieldGet(dataOutputStreamClassStruct, recv.native, "parent_class")
	value := FilterOutputStreamClassNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *DataOutputStreamClass) SetFieldParentClass(value *FilterOutputStreamClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.StructFieldSet(dataOutputStreamClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value '_g_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved4' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved5' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved5' : for field setter : missing Type

// DataOutputStreamClassStruct creates an uninitialised DataOutputStreamClass.
func DataOutputStreamClassStruct() *DataOutputStreamClass {
	err := dataOutputStreamClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := DataOutputStreamClassNewFromNative(dataOutputStreamClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeDataOutputStreamClass)
	return structGo
}
func finalizeDataOutputStreamClass(obj *DataOutputStreamClass) {
	dataOutputStreamClassStruct.Free(obj.native)
}

var dataOutputStreamPrivateStruct *gi.Struct
var dataOutputStreamPrivateStruct_Once sync.Once

func dataOutputStreamPrivateStruct_Set() error {
	var err error
	dataOutputStreamPrivateStruct_Once.Do(func() {
		dataOutputStreamPrivateStruct, err = gi.StructNew("Gio", "DataOutputStreamPrivate")
	})
	return err
}

type DataOutputStreamPrivate struct {
	native unsafe.Pointer
}

func DataOutputStreamPrivateNewFromNative(native unsafe.Pointer) *DataOutputStreamPrivate {
	instance := &DataOutputStreamPrivate{native: native}

	return instance
}

/*
CastToDataOutputStreamPrivate down casts any arbitrary Object to DataOutputStreamPrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a DataOutputStreamPrivate.
*/
func (recv *DataOutputStreamPrivate) CastToDataOutputStreamPrivate(object *gobject.Object) *DataOutputStreamPrivate {
	return DataOutputStreamPrivateNewFromNative(object.Native())
}

func (recv *DataOutputStreamPrivate) Native() unsafe.Pointer {
	return recv.native
}

// DataOutputStreamPrivateStruct creates an uninitialised DataOutputStreamPrivate.
func DataOutputStreamPrivateStruct() *DataOutputStreamPrivate {
	err := dataOutputStreamPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := DataOutputStreamPrivateNewFromNative(dataOutputStreamPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeDataOutputStreamPrivate)
	return structGo
}
func finalizeDataOutputStreamPrivate(obj *DataOutputStreamPrivate) {
	dataOutputStreamPrivateStruct.Free(obj.native)
}

var datagramBasedInterfaceStruct *gi.Struct
var datagramBasedInterfaceStruct_Once sync.Once

func datagramBasedInterfaceStruct_Set() error {
	var err error
	datagramBasedInterfaceStruct_Once.Do(func() {
		datagramBasedInterfaceStruct, err = gi.StructNew("Gio", "DatagramBasedInterface")
	})
	return err
}

type DatagramBasedInterface struct {
	native unsafe.Pointer
}

func DatagramBasedInterfaceNewFromNative(native unsafe.Pointer) *DatagramBasedInterface {
	instance := &DatagramBasedInterface{native: native}

	return instance
}

/*
CastToDatagramBasedInterface down casts any arbitrary Object to DatagramBasedInterface.
Exercise care, as this is a potentially dangerous function
if the Object is not a DatagramBasedInterface.
*/
func (recv *DatagramBasedInterface) CastToDatagramBasedInterface(object *gobject.Object) *DatagramBasedInterface {
	return DatagramBasedInterfaceNewFromNative(object.Native())
}

func (recv *DatagramBasedInterface) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'g_iface' : for field getter : no Go type for 'GObject.TypeInterface'

// UNSUPPORTED : C value 'g_iface' : for field setter : no Go type for 'GObject.TypeInterface'

// UNSUPPORTED : C value 'receive_messages' : for field getter : missing Type

// UNSUPPORTED : C value 'receive_messages' : for field setter : missing Type

// UNSUPPORTED : C value 'send_messages' : for field getter : missing Type

// UNSUPPORTED : C value 'send_messages' : for field setter : missing Type

// UNSUPPORTED : C value 'create_source' : for field getter : missing Type

// UNSUPPORTED : C value 'create_source' : for field setter : missing Type

// UNSUPPORTED : C value 'condition_check' : for field getter : missing Type

// UNSUPPORTED : C value 'condition_check' : for field setter : missing Type

// UNSUPPORTED : C value 'condition_wait' : for field getter : missing Type

// UNSUPPORTED : C value 'condition_wait' : for field setter : missing Type

// DatagramBasedInterfaceStruct creates an uninitialised DatagramBasedInterface.
func DatagramBasedInterfaceStruct() *DatagramBasedInterface {
	err := datagramBasedInterfaceStruct_Set()
	if err != nil {
		return nil
	}

	structGo := DatagramBasedInterfaceNewFromNative(datagramBasedInterfaceStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeDatagramBasedInterface)
	return structGo
}
func finalizeDatagramBasedInterface(obj *DatagramBasedInterface) {
	datagramBasedInterfaceStruct.Free(obj.native)
}

var desktopAppInfoClassStruct *gi.Struct
var desktopAppInfoClassStruct_Once sync.Once

func desktopAppInfoClassStruct_Set() error {
	var err error
	desktopAppInfoClassStruct_Once.Do(func() {
		desktopAppInfoClassStruct, err = gi.StructNew("Gio", "DesktopAppInfoClass")
	})
	return err
}

type DesktopAppInfoClass struct {
	native unsafe.Pointer
}

func DesktopAppInfoClassNewFromNative(native unsafe.Pointer) *DesktopAppInfoClass {
	instance := &DesktopAppInfoClass{native: native}

	return instance
}

/*
CastToDesktopAppInfoClass down casts any arbitrary Object to DesktopAppInfoClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a DesktopAppInfoClass.
*/
func (recv *DesktopAppInfoClass) CastToDesktopAppInfoClass(object *gobject.Object) *DesktopAppInfoClass {
	return DesktopAppInfoClassNewFromNative(object.Native())
}

func (recv *DesktopAppInfoClass) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// DesktopAppInfoClassStruct creates an uninitialised DesktopAppInfoClass.
func DesktopAppInfoClassStruct() *DesktopAppInfoClass {
	err := desktopAppInfoClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := DesktopAppInfoClassNewFromNative(desktopAppInfoClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeDesktopAppInfoClass)
	return structGo
}
func finalizeDesktopAppInfoClass(obj *DesktopAppInfoClass) {
	desktopAppInfoClassStruct.Free(obj.native)
}

var desktopAppInfoLookupIfaceStruct *gi.Struct
var desktopAppInfoLookupIfaceStruct_Once sync.Once

func desktopAppInfoLookupIfaceStruct_Set() error {
	var err error
	desktopAppInfoLookupIfaceStruct_Once.Do(func() {
		desktopAppInfoLookupIfaceStruct, err = gi.StructNew("Gio", "DesktopAppInfoLookupIface")
	})
	return err
}

type DesktopAppInfoLookupIface struct {
	native unsafe.Pointer
}

func DesktopAppInfoLookupIfaceNewFromNative(native unsafe.Pointer) *DesktopAppInfoLookupIface {
	instance := &DesktopAppInfoLookupIface{native: native}

	return instance
}

/*
CastToDesktopAppInfoLookupIface down casts any arbitrary Object to DesktopAppInfoLookupIface.
Exercise care, as this is a potentially dangerous function
if the Object is not a DesktopAppInfoLookupIface.
*/
func (recv *DesktopAppInfoLookupIface) CastToDesktopAppInfoLookupIface(object *gobject.Object) *DesktopAppInfoLookupIface {
	return DesktopAppInfoLookupIfaceNewFromNative(object.Native())
}

func (recv *DesktopAppInfoLookupIface) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'g_iface' : for field getter : no Go type for 'GObject.TypeInterface'

// UNSUPPORTED : C value 'g_iface' : for field setter : no Go type for 'GObject.TypeInterface'

// UNSUPPORTED : C value 'get_default_for_uri_scheme' : for field getter : missing Type

// UNSUPPORTED : C value 'get_default_for_uri_scheme' : for field setter : missing Type

// DesktopAppInfoLookupIfaceStruct creates an uninitialised DesktopAppInfoLookupIface.
func DesktopAppInfoLookupIfaceStruct() *DesktopAppInfoLookupIface {
	err := desktopAppInfoLookupIfaceStruct_Set()
	if err != nil {
		return nil
	}

	structGo := DesktopAppInfoLookupIfaceNewFromNative(desktopAppInfoLookupIfaceStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeDesktopAppInfoLookupIface)
	return structGo
}
func finalizeDesktopAppInfoLookupIface(obj *DesktopAppInfoLookupIface) {
	desktopAppInfoLookupIfaceStruct.Free(obj.native)
}

var driveIfaceStruct *gi.Struct
var driveIfaceStruct_Once sync.Once

func driveIfaceStruct_Set() error {
	var err error
	driveIfaceStruct_Once.Do(func() {
		driveIfaceStruct, err = gi.StructNew("Gio", "DriveIface")
	})
	return err
}

type DriveIface struct {
	native unsafe.Pointer
}

func DriveIfaceNewFromNative(native unsafe.Pointer) *DriveIface {
	instance := &DriveIface{native: native}

	return instance
}

/*
CastToDriveIface down casts any arbitrary Object to DriveIface.
Exercise care, as this is a potentially dangerous function
if the Object is not a DriveIface.
*/
func (recv *DriveIface) CastToDriveIface(object *gobject.Object) *DriveIface {
	return DriveIfaceNewFromNative(object.Native())
}

func (recv *DriveIface) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'g_iface' : for field getter : no Go type for 'GObject.TypeInterface'

// UNSUPPORTED : C value 'g_iface' : for field setter : no Go type for 'GObject.TypeInterface'

// UNSUPPORTED : C value 'changed' : for field getter : missing Type

// UNSUPPORTED : C value 'changed' : for field setter : missing Type

// UNSUPPORTED : C value 'disconnected' : for field getter : missing Type

// UNSUPPORTED : C value 'disconnected' : for field setter : missing Type

// UNSUPPORTED : C value 'eject_button' : for field getter : missing Type

// UNSUPPORTED : C value 'eject_button' : for field setter : missing Type

// UNSUPPORTED : C value 'get_name' : for field getter : missing Type

// UNSUPPORTED : C value 'get_name' : for field setter : missing Type

// UNSUPPORTED : C value 'get_icon' : for field getter : missing Type

// UNSUPPORTED : C value 'get_icon' : for field setter : missing Type

// UNSUPPORTED : C value 'has_volumes' : for field getter : missing Type

// UNSUPPORTED : C value 'has_volumes' : for field setter : missing Type

// UNSUPPORTED : C value 'get_volumes' : for field getter : missing Type

// UNSUPPORTED : C value 'get_volumes' : for field setter : missing Type

// UNSUPPORTED : C value 'is_media_removable' : for field getter : missing Type

// UNSUPPORTED : C value 'is_media_removable' : for field setter : missing Type

// UNSUPPORTED : C value 'has_media' : for field getter : missing Type

// UNSUPPORTED : C value 'has_media' : for field setter : missing Type

// UNSUPPORTED : C value 'is_media_check_automatic' : for field getter : missing Type

// UNSUPPORTED : C value 'is_media_check_automatic' : for field setter : missing Type

// UNSUPPORTED : C value 'can_eject' : for field getter : missing Type

// UNSUPPORTED : C value 'can_eject' : for field setter : missing Type

// UNSUPPORTED : C value 'can_poll_for_media' : for field getter : missing Type

// UNSUPPORTED : C value 'can_poll_for_media' : for field setter : missing Type

// UNSUPPORTED : C value 'eject' : for field getter : missing Type

// UNSUPPORTED : C value 'eject' : for field setter : missing Type

// UNSUPPORTED : C value 'eject_finish' : for field getter : missing Type

// UNSUPPORTED : C value 'eject_finish' : for field setter : missing Type

// UNSUPPORTED : C value 'poll_for_media' : for field getter : missing Type

// UNSUPPORTED : C value 'poll_for_media' : for field setter : missing Type

// UNSUPPORTED : C value 'poll_for_media_finish' : for field getter : missing Type

// UNSUPPORTED : C value 'poll_for_media_finish' : for field setter : missing Type

// UNSUPPORTED : C value 'get_identifier' : for field getter : missing Type

// UNSUPPORTED : C value 'get_identifier' : for field setter : missing Type

// UNSUPPORTED : C value 'enumerate_identifiers' : for field getter : missing Type

// UNSUPPORTED : C value 'enumerate_identifiers' : for field setter : missing Type

// UNSUPPORTED : C value 'get_start_stop_type' : for field getter : missing Type

// UNSUPPORTED : C value 'get_start_stop_type' : for field setter : missing Type

// UNSUPPORTED : C value 'can_start' : for field getter : missing Type

// UNSUPPORTED : C value 'can_start' : for field setter : missing Type

// UNSUPPORTED : C value 'can_start_degraded' : for field getter : missing Type

// UNSUPPORTED : C value 'can_start_degraded' : for field setter : missing Type

// UNSUPPORTED : C value 'start' : for field getter : missing Type

// UNSUPPORTED : C value 'start' : for field setter : missing Type

// UNSUPPORTED : C value 'start_finish' : for field getter : missing Type

// UNSUPPORTED : C value 'start_finish' : for field setter : missing Type

// UNSUPPORTED : C value 'can_stop' : for field getter : missing Type

// UNSUPPORTED : C value 'can_stop' : for field setter : missing Type

// UNSUPPORTED : C value 'stop' : for field getter : missing Type

// UNSUPPORTED : C value 'stop' : for field setter : missing Type

// UNSUPPORTED : C value 'stop_finish' : for field getter : missing Type

// UNSUPPORTED : C value 'stop_finish' : for field setter : missing Type

// UNSUPPORTED : C value 'stop_button' : for field getter : missing Type

// UNSUPPORTED : C value 'stop_button' : for field setter : missing Type

// UNSUPPORTED : C value 'eject_with_operation' : for field getter : missing Type

// UNSUPPORTED : C value 'eject_with_operation' : for field setter : missing Type

// UNSUPPORTED : C value 'eject_with_operation_finish' : for field getter : missing Type

// UNSUPPORTED : C value 'eject_with_operation_finish' : for field setter : missing Type

// UNSUPPORTED : C value 'get_sort_key' : for field getter : missing Type

// UNSUPPORTED : C value 'get_sort_key' : for field setter : missing Type

// UNSUPPORTED : C value 'get_symbolic_icon' : for field getter : missing Type

// UNSUPPORTED : C value 'get_symbolic_icon' : for field setter : missing Type

// UNSUPPORTED : C value 'is_removable' : for field getter : missing Type

// UNSUPPORTED : C value 'is_removable' : for field setter : missing Type

// DriveIfaceStruct creates an uninitialised DriveIface.
func DriveIfaceStruct() *DriveIface {
	err := driveIfaceStruct_Set()
	if err != nil {
		return nil
	}

	structGo := DriveIfaceNewFromNative(driveIfaceStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeDriveIface)
	return structGo
}
func finalizeDriveIface(obj *DriveIface) {
	driveIfaceStruct.Free(obj.native)
}

var dtlsClientConnectionInterfaceStruct *gi.Struct
var dtlsClientConnectionInterfaceStruct_Once sync.Once

func dtlsClientConnectionInterfaceStruct_Set() error {
	var err error
	dtlsClientConnectionInterfaceStruct_Once.Do(func() {
		dtlsClientConnectionInterfaceStruct, err = gi.StructNew("Gio", "DtlsClientConnectionInterface")
	})
	return err
}

type DtlsClientConnectionInterface struct {
	native unsafe.Pointer
}

func DtlsClientConnectionInterfaceNewFromNative(native unsafe.Pointer) *DtlsClientConnectionInterface {
	instance := &DtlsClientConnectionInterface{native: native}

	return instance
}

/*
CastToDtlsClientConnectionInterface down casts any arbitrary Object to DtlsClientConnectionInterface.
Exercise care, as this is a potentially dangerous function
if the Object is not a DtlsClientConnectionInterface.
*/
func (recv *DtlsClientConnectionInterface) CastToDtlsClientConnectionInterface(object *gobject.Object) *DtlsClientConnectionInterface {
	return DtlsClientConnectionInterfaceNewFromNative(object.Native())
}

func (recv *DtlsClientConnectionInterface) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'g_iface' : for field getter : no Go type for 'GObject.TypeInterface'

// UNSUPPORTED : C value 'g_iface' : for field setter : no Go type for 'GObject.TypeInterface'

// DtlsClientConnectionInterfaceStruct creates an uninitialised DtlsClientConnectionInterface.
func DtlsClientConnectionInterfaceStruct() *DtlsClientConnectionInterface {
	err := dtlsClientConnectionInterfaceStruct_Set()
	if err != nil {
		return nil
	}

	structGo := DtlsClientConnectionInterfaceNewFromNative(dtlsClientConnectionInterfaceStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeDtlsClientConnectionInterface)
	return structGo
}
func finalizeDtlsClientConnectionInterface(obj *DtlsClientConnectionInterface) {
	dtlsClientConnectionInterfaceStruct.Free(obj.native)
}

var dtlsConnectionInterfaceStruct *gi.Struct
var dtlsConnectionInterfaceStruct_Once sync.Once

func dtlsConnectionInterfaceStruct_Set() error {
	var err error
	dtlsConnectionInterfaceStruct_Once.Do(func() {
		dtlsConnectionInterfaceStruct, err = gi.StructNew("Gio", "DtlsConnectionInterface")
	})
	return err
}

type DtlsConnectionInterface struct {
	native unsafe.Pointer
}

func DtlsConnectionInterfaceNewFromNative(native unsafe.Pointer) *DtlsConnectionInterface {
	instance := &DtlsConnectionInterface{native: native}

	return instance
}

/*
CastToDtlsConnectionInterface down casts any arbitrary Object to DtlsConnectionInterface.
Exercise care, as this is a potentially dangerous function
if the Object is not a DtlsConnectionInterface.
*/
func (recv *DtlsConnectionInterface) CastToDtlsConnectionInterface(object *gobject.Object) *DtlsConnectionInterface {
	return DtlsConnectionInterfaceNewFromNative(object.Native())
}

func (recv *DtlsConnectionInterface) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'g_iface' : for field getter : no Go type for 'GObject.TypeInterface'

// UNSUPPORTED : C value 'g_iface' : for field setter : no Go type for 'GObject.TypeInterface'

// UNSUPPORTED : C value 'accept_certificate' : for field getter : missing Type

// UNSUPPORTED : C value 'accept_certificate' : for field setter : missing Type

// UNSUPPORTED : C value 'handshake' : for field getter : missing Type

// UNSUPPORTED : C value 'handshake' : for field setter : missing Type

// UNSUPPORTED : C value 'handshake_async' : for field getter : missing Type

// UNSUPPORTED : C value 'handshake_async' : for field setter : missing Type

// UNSUPPORTED : C value 'handshake_finish' : for field getter : missing Type

// UNSUPPORTED : C value 'handshake_finish' : for field setter : missing Type

// UNSUPPORTED : C value 'shutdown' : for field getter : missing Type

// UNSUPPORTED : C value 'shutdown' : for field setter : missing Type

// UNSUPPORTED : C value 'shutdown_async' : for field getter : missing Type

// UNSUPPORTED : C value 'shutdown_async' : for field setter : missing Type

// UNSUPPORTED : C value 'shutdown_finish' : for field getter : missing Type

// UNSUPPORTED : C value 'shutdown_finish' : for field setter : missing Type

// UNSUPPORTED : C value 'set_advertised_protocols' : for field getter : missing Type

// UNSUPPORTED : C value 'set_advertised_protocols' : for field setter : missing Type

// UNSUPPORTED : C value 'get_negotiated_protocol' : for field getter : missing Type

// UNSUPPORTED : C value 'get_negotiated_protocol' : for field setter : missing Type

// DtlsConnectionInterfaceStruct creates an uninitialised DtlsConnectionInterface.
func DtlsConnectionInterfaceStruct() *DtlsConnectionInterface {
	err := dtlsConnectionInterfaceStruct_Set()
	if err != nil {
		return nil
	}

	structGo := DtlsConnectionInterfaceNewFromNative(dtlsConnectionInterfaceStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeDtlsConnectionInterface)
	return structGo
}
func finalizeDtlsConnectionInterface(obj *DtlsConnectionInterface) {
	dtlsConnectionInterfaceStruct.Free(obj.native)
}

var dtlsServerConnectionInterfaceStruct *gi.Struct
var dtlsServerConnectionInterfaceStruct_Once sync.Once

func dtlsServerConnectionInterfaceStruct_Set() error {
	var err error
	dtlsServerConnectionInterfaceStruct_Once.Do(func() {
		dtlsServerConnectionInterfaceStruct, err = gi.StructNew("Gio", "DtlsServerConnectionInterface")
	})
	return err
}

type DtlsServerConnectionInterface struct {
	native unsafe.Pointer
}

func DtlsServerConnectionInterfaceNewFromNative(native unsafe.Pointer) *DtlsServerConnectionInterface {
	instance := &DtlsServerConnectionInterface{native: native}

	return instance
}

/*
CastToDtlsServerConnectionInterface down casts any arbitrary Object to DtlsServerConnectionInterface.
Exercise care, as this is a potentially dangerous function
if the Object is not a DtlsServerConnectionInterface.
*/
func (recv *DtlsServerConnectionInterface) CastToDtlsServerConnectionInterface(object *gobject.Object) *DtlsServerConnectionInterface {
	return DtlsServerConnectionInterfaceNewFromNative(object.Native())
}

func (recv *DtlsServerConnectionInterface) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'g_iface' : for field getter : no Go type for 'GObject.TypeInterface'

// UNSUPPORTED : C value 'g_iface' : for field setter : no Go type for 'GObject.TypeInterface'

// DtlsServerConnectionInterfaceStruct creates an uninitialised DtlsServerConnectionInterface.
func DtlsServerConnectionInterfaceStruct() *DtlsServerConnectionInterface {
	err := dtlsServerConnectionInterfaceStruct_Set()
	if err != nil {
		return nil
	}

	structGo := DtlsServerConnectionInterfaceNewFromNative(dtlsServerConnectionInterfaceStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeDtlsServerConnectionInterface)
	return structGo
}
func finalizeDtlsServerConnectionInterface(obj *DtlsServerConnectionInterface) {
	dtlsServerConnectionInterfaceStruct.Free(obj.native)
}

var emblemClassStruct *gi.Struct
var emblemClassStruct_Once sync.Once

func emblemClassStruct_Set() error {
	var err error
	emblemClassStruct_Once.Do(func() {
		emblemClassStruct, err = gi.StructNew("Gio", "EmblemClass")
	})
	return err
}

type EmblemClass struct {
	native unsafe.Pointer
}

func EmblemClassNewFromNative(native unsafe.Pointer) *EmblemClass {
	instance := &EmblemClass{native: native}

	return instance
}

/*
CastToEmblemClass down casts any arbitrary Object to EmblemClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a EmblemClass.
*/
func (recv *EmblemClass) CastToEmblemClass(object *gobject.Object) *EmblemClass {
	return EmblemClassNewFromNative(object.Native())
}

func (recv *EmblemClass) Native() unsafe.Pointer {
	return recv.native
}

// EmblemClassStruct creates an uninitialised EmblemClass.
func EmblemClassStruct() *EmblemClass {
	err := emblemClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := EmblemClassNewFromNative(emblemClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeEmblemClass)
	return structGo
}
func finalizeEmblemClass(obj *EmblemClass) {
	emblemClassStruct.Free(obj.native)
}

var emblemedIconClassStruct *gi.Struct
var emblemedIconClassStruct_Once sync.Once

func emblemedIconClassStruct_Set() error {
	var err error
	emblemedIconClassStruct_Once.Do(func() {
		emblemedIconClassStruct, err = gi.StructNew("Gio", "EmblemedIconClass")
	})
	return err
}

type EmblemedIconClass struct {
	native unsafe.Pointer
}

func EmblemedIconClassNewFromNative(native unsafe.Pointer) *EmblemedIconClass {
	instance := &EmblemedIconClass{native: native}

	return instance
}

/*
CastToEmblemedIconClass down casts any arbitrary Object to EmblemedIconClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a EmblemedIconClass.
*/
func (recv *EmblemedIconClass) CastToEmblemedIconClass(object *gobject.Object) *EmblemedIconClass {
	return EmblemedIconClassNewFromNative(object.Native())
}

func (recv *EmblemedIconClass) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// EmblemedIconClassStruct creates an uninitialised EmblemedIconClass.
func EmblemedIconClassStruct() *EmblemedIconClass {
	err := emblemedIconClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := EmblemedIconClassNewFromNative(emblemedIconClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeEmblemedIconClass)
	return structGo
}
func finalizeEmblemedIconClass(obj *EmblemedIconClass) {
	emblemedIconClassStruct.Free(obj.native)
}

var emblemedIconPrivateStruct *gi.Struct
var emblemedIconPrivateStruct_Once sync.Once

func emblemedIconPrivateStruct_Set() error {
	var err error
	emblemedIconPrivateStruct_Once.Do(func() {
		emblemedIconPrivateStruct, err = gi.StructNew("Gio", "EmblemedIconPrivate")
	})
	return err
}

type EmblemedIconPrivate struct {
	native unsafe.Pointer
}

func EmblemedIconPrivateNewFromNative(native unsafe.Pointer) *EmblemedIconPrivate {
	instance := &EmblemedIconPrivate{native: native}

	return instance
}

/*
CastToEmblemedIconPrivate down casts any arbitrary Object to EmblemedIconPrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a EmblemedIconPrivate.
*/
func (recv *EmblemedIconPrivate) CastToEmblemedIconPrivate(object *gobject.Object) *EmblemedIconPrivate {
	return EmblemedIconPrivateNewFromNative(object.Native())
}

func (recv *EmblemedIconPrivate) Native() unsafe.Pointer {
	return recv.native
}

// EmblemedIconPrivateStruct creates an uninitialised EmblemedIconPrivate.
func EmblemedIconPrivateStruct() *EmblemedIconPrivate {
	err := emblemedIconPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := EmblemedIconPrivateNewFromNative(emblemedIconPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeEmblemedIconPrivate)
	return structGo
}
func finalizeEmblemedIconPrivate(obj *EmblemedIconPrivate) {
	emblemedIconPrivateStruct.Free(obj.native)
}

var fileAttributeInfoStruct *gi.Struct
var fileAttributeInfoStruct_Once sync.Once

func fileAttributeInfoStruct_Set() error {
	var err error
	fileAttributeInfoStruct_Once.Do(func() {
		fileAttributeInfoStruct, err = gi.StructNew("Gio", "FileAttributeInfo")
	})
	return err
}

type FileAttributeInfo struct {
	native unsafe.Pointer
}

func FileAttributeInfoNewFromNative(native unsafe.Pointer) *FileAttributeInfo {
	instance := &FileAttributeInfo{native: native}

	return instance
}

/*
CastToFileAttributeInfo down casts any arbitrary Object to FileAttributeInfo.
Exercise care, as this is a potentially dangerous function
if the Object is not a FileAttributeInfo.
*/
func (recv *FileAttributeInfo) CastToFileAttributeInfo(object *gobject.Object) *FileAttributeInfo {
	return FileAttributeInfoNewFromNative(object.Native())
}

func (recv *FileAttributeInfo) Native() unsafe.Pointer {
	return recv.native
}

// FieldName returns the C field 'name'.
func (recv *FileAttributeInfo) FieldName() string {
	argValue := gi.StructFieldGet(fileAttributeInfoStruct, recv.native, "name")
	value := argValue.String(false)
	return value
}

// SetFieldName sets the value of the C field 'name'.
func (recv *FileAttributeInfo) SetFieldName(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.StructFieldSet(fileAttributeInfoStruct, recv.native, "name", argValue)
}

// FieldType returns the C field 'type'.
func (recv *FileAttributeInfo) FieldType() FileAttributeType {
	argValue := gi.StructFieldGet(fileAttributeInfoStruct, recv.native, "type")
	value := FileAttributeType(argValue.Int32())
	return value
}

// SetFieldType sets the value of the C field 'type'.
func (recv *FileAttributeInfo) SetFieldType(value FileAttributeType) {
	var argValue gi.Argument
	argValue.SetInt32(int32(value))
	gi.StructFieldSet(fileAttributeInfoStruct, recv.native, "type", argValue)
}

// UNSUPPORTED : C value 'flags' : for field getter : no Go type for 'FileAttributeInfoFlags'

// UNSUPPORTED : C value 'flags' : for field setter : no Go type for 'FileAttributeInfoFlags'

// FileAttributeInfoStruct creates an uninitialised FileAttributeInfo.
func FileAttributeInfoStruct() *FileAttributeInfo {
	err := fileAttributeInfoStruct_Set()
	if err != nil {
		return nil
	}

	structGo := FileAttributeInfoNewFromNative(fileAttributeInfoStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeFileAttributeInfo)
	return structGo
}
func finalizeFileAttributeInfo(obj *FileAttributeInfo) {
	fileAttributeInfoStruct.Free(obj.native)
}

var fileAttributeInfoListStruct *gi.Struct
var fileAttributeInfoListStruct_Once sync.Once

func fileAttributeInfoListStruct_Set() error {
	var err error
	fileAttributeInfoListStruct_Once.Do(func() {
		fileAttributeInfoListStruct, err = gi.StructNew("Gio", "FileAttributeInfoList")
	})
	return err
}

type FileAttributeInfoList struct {
	native unsafe.Pointer
}

func FileAttributeInfoListNewFromNative(native unsafe.Pointer) *FileAttributeInfoList {
	instance := &FileAttributeInfoList{native: native}

	return instance
}

/*
CastToFileAttributeInfoList down casts any arbitrary Object to FileAttributeInfoList.
Exercise care, as this is a potentially dangerous function
if the Object is not a FileAttributeInfoList.
*/
func (recv *FileAttributeInfoList) CastToFileAttributeInfoList(object *gobject.Object) *FileAttributeInfoList {
	return FileAttributeInfoListNewFromNative(object.Native())
}

func (recv *FileAttributeInfoList) Native() unsafe.Pointer {
	return recv.native
}

// FieldInfos returns the C field 'infos'.
func (recv *FileAttributeInfoList) FieldInfos() *FileAttributeInfo {
	argValue := gi.StructFieldGet(fileAttributeInfoListStruct, recv.native, "infos")
	value := FileAttributeInfoNewFromNative(argValue.Pointer())
	return value
}

// SetFieldInfos sets the value of the C field 'infos'.
func (recv *FileAttributeInfoList) SetFieldInfos(value *FileAttributeInfo) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.StructFieldSet(fileAttributeInfoListStruct, recv.native, "infos", argValue)
}

// FieldNInfos returns the C field 'n_infos'.
func (recv *FileAttributeInfoList) FieldNInfos() int32 {
	argValue := gi.StructFieldGet(fileAttributeInfoListStruct, recv.native, "n_infos")
	value := argValue.Int32()
	return value
}

// SetFieldNInfos sets the value of the C field 'n_infos'.
func (recv *FileAttributeInfoList) SetFieldNInfos(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.StructFieldSet(fileAttributeInfoListStruct, recv.native, "n_infos", argValue)
}

var fileAttributeInfoListNewFunction *gi.Function
var fileAttributeInfoListNewFunction_Once sync.Once

func fileAttributeInfoListNewFunction_Set() error {
	var err error
	fileAttributeInfoListNewFunction_Once.Do(func() {
		err = fileAttributeInfoListStruct_Set()
		if err != nil {
			return
		}
		fileAttributeInfoListNewFunction, err = fileAttributeInfoListStruct.InvokerNew("new")
	})
	return err
}

// FileAttributeInfoListNew is a representation of the C type g_file_attribute_info_list_new.
func FileAttributeInfoListNew() *FileAttributeInfoList {

	var ret gi.Argument

	err := fileAttributeInfoListNewFunction_Set()
	if err == nil {
		ret = fileAttributeInfoListNewFunction.Invoke(nil, nil)
	}

	retGo := FileAttributeInfoListNewFromNative(ret.Pointer())

	return retGo
}

// UNSUPPORTED : C value 'g_file_attribute_info_list_add' : parameter 'flags' of type 'FileAttributeInfoFlags' not supported

var fileAttributeInfoListDupFunction *gi.Function
var fileAttributeInfoListDupFunction_Once sync.Once

func fileAttributeInfoListDupFunction_Set() error {
	var err error
	fileAttributeInfoListDupFunction_Once.Do(func() {
		err = fileAttributeInfoListStruct_Set()
		if err != nil {
			return
		}
		fileAttributeInfoListDupFunction, err = fileAttributeInfoListStruct.InvokerNew("dup")
	})
	return err
}

// Dup is a representation of the C type g_file_attribute_info_list_dup.
func (recv *FileAttributeInfoList) Dup() *FileAttributeInfoList {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := fileAttributeInfoListDupFunction_Set()
	if err == nil {
		ret = fileAttributeInfoListDupFunction.Invoke(inArgs[:], nil)
	}

	retGo := FileAttributeInfoListNewFromNative(ret.Pointer())

	return retGo
}

var fileAttributeInfoListLookupFunction *gi.Function
var fileAttributeInfoListLookupFunction_Once sync.Once

func fileAttributeInfoListLookupFunction_Set() error {
	var err error
	fileAttributeInfoListLookupFunction_Once.Do(func() {
		err = fileAttributeInfoListStruct_Set()
		if err != nil {
			return
		}
		fileAttributeInfoListLookupFunction, err = fileAttributeInfoListStruct.InvokerNew("lookup")
	})
	return err
}

// Lookup is a representation of the C type g_file_attribute_info_list_lookup.
func (recv *FileAttributeInfoList) Lookup(name string) *FileAttributeInfo {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(name)

	var ret gi.Argument

	err := fileAttributeInfoListLookupFunction_Set()
	if err == nil {
		ret = fileAttributeInfoListLookupFunction.Invoke(inArgs[:], nil)
	}

	retGo := FileAttributeInfoNewFromNative(ret.Pointer())

	return retGo
}

var fileAttributeInfoListRefFunction *gi.Function
var fileAttributeInfoListRefFunction_Once sync.Once

func fileAttributeInfoListRefFunction_Set() error {
	var err error
	fileAttributeInfoListRefFunction_Once.Do(func() {
		err = fileAttributeInfoListStruct_Set()
		if err != nil {
			return
		}
		fileAttributeInfoListRefFunction, err = fileAttributeInfoListStruct.InvokerNew("ref")
	})
	return err
}

// Ref is a representation of the C type g_file_attribute_info_list_ref.
func (recv *FileAttributeInfoList) Ref() *FileAttributeInfoList {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := fileAttributeInfoListRefFunction_Set()
	if err == nil {
		ret = fileAttributeInfoListRefFunction.Invoke(inArgs[:], nil)
	}

	retGo := FileAttributeInfoListNewFromNative(ret.Pointer())

	return retGo
}

var fileAttributeInfoListUnrefFunction *gi.Function
var fileAttributeInfoListUnrefFunction_Once sync.Once

func fileAttributeInfoListUnrefFunction_Set() error {
	var err error
	fileAttributeInfoListUnrefFunction_Once.Do(func() {
		err = fileAttributeInfoListStruct_Set()
		if err != nil {
			return
		}
		fileAttributeInfoListUnrefFunction, err = fileAttributeInfoListStruct.InvokerNew("unref")
	})
	return err
}

// Unref is a representation of the C type g_file_attribute_info_list_unref.
func (recv *FileAttributeInfoList) Unref() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := fileAttributeInfoListUnrefFunction_Set()
	if err == nil {
		fileAttributeInfoListUnrefFunction.Invoke(inArgs[:], nil)
	}

	return
}

var fileAttributeMatcherStruct *gi.Struct
var fileAttributeMatcherStruct_Once sync.Once

func fileAttributeMatcherStruct_Set() error {
	var err error
	fileAttributeMatcherStruct_Once.Do(func() {
		fileAttributeMatcherStruct, err = gi.StructNew("Gio", "FileAttributeMatcher")
	})
	return err
}

type FileAttributeMatcher struct {
	native unsafe.Pointer
}

func FileAttributeMatcherNewFromNative(native unsafe.Pointer) *FileAttributeMatcher {
	instance := &FileAttributeMatcher{native: native}

	return instance
}

/*
CastToFileAttributeMatcher down casts any arbitrary Object to FileAttributeMatcher.
Exercise care, as this is a potentially dangerous function
if the Object is not a FileAttributeMatcher.
*/
func (recv *FileAttributeMatcher) CastToFileAttributeMatcher(object *gobject.Object) *FileAttributeMatcher {
	return FileAttributeMatcherNewFromNative(object.Native())
}

func (recv *FileAttributeMatcher) Native() unsafe.Pointer {
	return recv.native
}

var fileAttributeMatcherNewFunction *gi.Function
var fileAttributeMatcherNewFunction_Once sync.Once

func fileAttributeMatcherNewFunction_Set() error {
	var err error
	fileAttributeMatcherNewFunction_Once.Do(func() {
		err = fileAttributeMatcherStruct_Set()
		if err != nil {
			return
		}
		fileAttributeMatcherNewFunction, err = fileAttributeMatcherStruct.InvokerNew("new")
	})
	return err
}

// FileAttributeMatcherNew is a representation of the C type g_file_attribute_matcher_new.
func FileAttributeMatcherNew(attributes string) *FileAttributeMatcher {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(attributes)

	var ret gi.Argument

	err := fileAttributeMatcherNewFunction_Set()
	if err == nil {
		ret = fileAttributeMatcherNewFunction.Invoke(inArgs[:], nil)
	}

	retGo := FileAttributeMatcherNewFromNative(ret.Pointer())

	return retGo
}

var fileAttributeMatcherEnumerateNamespaceFunction *gi.Function
var fileAttributeMatcherEnumerateNamespaceFunction_Once sync.Once

func fileAttributeMatcherEnumerateNamespaceFunction_Set() error {
	var err error
	fileAttributeMatcherEnumerateNamespaceFunction_Once.Do(func() {
		err = fileAttributeMatcherStruct_Set()
		if err != nil {
			return
		}
		fileAttributeMatcherEnumerateNamespaceFunction, err = fileAttributeMatcherStruct.InvokerNew("enumerate_namespace")
	})
	return err
}

// EnumerateNamespace is a representation of the C type g_file_attribute_matcher_enumerate_namespace.
func (recv *FileAttributeMatcher) EnumerateNamespace(ns string) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(ns)

	var ret gi.Argument

	err := fileAttributeMatcherEnumerateNamespaceFunction_Set()
	if err == nil {
		ret = fileAttributeMatcherEnumerateNamespaceFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var fileAttributeMatcherEnumerateNextFunction *gi.Function
var fileAttributeMatcherEnumerateNextFunction_Once sync.Once

func fileAttributeMatcherEnumerateNextFunction_Set() error {
	var err error
	fileAttributeMatcherEnumerateNextFunction_Once.Do(func() {
		err = fileAttributeMatcherStruct_Set()
		if err != nil {
			return
		}
		fileAttributeMatcherEnumerateNextFunction, err = fileAttributeMatcherStruct.InvokerNew("enumerate_next")
	})
	return err
}

// EnumerateNext is a representation of the C type g_file_attribute_matcher_enumerate_next.
func (recv *FileAttributeMatcher) EnumerateNext() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := fileAttributeMatcherEnumerateNextFunction_Set()
	if err == nil {
		ret = fileAttributeMatcherEnumerateNextFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var fileAttributeMatcherMatchesFunction *gi.Function
var fileAttributeMatcherMatchesFunction_Once sync.Once

func fileAttributeMatcherMatchesFunction_Set() error {
	var err error
	fileAttributeMatcherMatchesFunction_Once.Do(func() {
		err = fileAttributeMatcherStruct_Set()
		if err != nil {
			return
		}
		fileAttributeMatcherMatchesFunction, err = fileAttributeMatcherStruct.InvokerNew("matches")
	})
	return err
}

// Matches is a representation of the C type g_file_attribute_matcher_matches.
func (recv *FileAttributeMatcher) Matches(attribute string) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(attribute)

	var ret gi.Argument

	err := fileAttributeMatcherMatchesFunction_Set()
	if err == nil {
		ret = fileAttributeMatcherMatchesFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var fileAttributeMatcherMatchesOnlyFunction *gi.Function
var fileAttributeMatcherMatchesOnlyFunction_Once sync.Once

func fileAttributeMatcherMatchesOnlyFunction_Set() error {
	var err error
	fileAttributeMatcherMatchesOnlyFunction_Once.Do(func() {
		err = fileAttributeMatcherStruct_Set()
		if err != nil {
			return
		}
		fileAttributeMatcherMatchesOnlyFunction, err = fileAttributeMatcherStruct.InvokerNew("matches_only")
	})
	return err
}

// MatchesOnly is a representation of the C type g_file_attribute_matcher_matches_only.
func (recv *FileAttributeMatcher) MatchesOnly(attribute string) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(attribute)

	var ret gi.Argument

	err := fileAttributeMatcherMatchesOnlyFunction_Set()
	if err == nil {
		ret = fileAttributeMatcherMatchesOnlyFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var fileAttributeMatcherRefFunction *gi.Function
var fileAttributeMatcherRefFunction_Once sync.Once

func fileAttributeMatcherRefFunction_Set() error {
	var err error
	fileAttributeMatcherRefFunction_Once.Do(func() {
		err = fileAttributeMatcherStruct_Set()
		if err != nil {
			return
		}
		fileAttributeMatcherRefFunction, err = fileAttributeMatcherStruct.InvokerNew("ref")
	})
	return err
}

// Ref is a representation of the C type g_file_attribute_matcher_ref.
func (recv *FileAttributeMatcher) Ref() *FileAttributeMatcher {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := fileAttributeMatcherRefFunction_Set()
	if err == nil {
		ret = fileAttributeMatcherRefFunction.Invoke(inArgs[:], nil)
	}

	retGo := FileAttributeMatcherNewFromNative(ret.Pointer())

	return retGo
}

var fileAttributeMatcherSubtractFunction *gi.Function
var fileAttributeMatcherSubtractFunction_Once sync.Once

func fileAttributeMatcherSubtractFunction_Set() error {
	var err error
	fileAttributeMatcherSubtractFunction_Once.Do(func() {
		err = fileAttributeMatcherStruct_Set()
		if err != nil {
			return
		}
		fileAttributeMatcherSubtractFunction, err = fileAttributeMatcherStruct.InvokerNew("subtract")
	})
	return err
}

// Subtract is a representation of the C type g_file_attribute_matcher_subtract.
func (recv *FileAttributeMatcher) Subtract(subtract *FileAttributeMatcher) *FileAttributeMatcher {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(subtract.native)

	var ret gi.Argument

	err := fileAttributeMatcherSubtractFunction_Set()
	if err == nil {
		ret = fileAttributeMatcherSubtractFunction.Invoke(inArgs[:], nil)
	}

	retGo := FileAttributeMatcherNewFromNative(ret.Pointer())

	return retGo
}

var fileAttributeMatcherToStringFunction *gi.Function
var fileAttributeMatcherToStringFunction_Once sync.Once

func fileAttributeMatcherToStringFunction_Set() error {
	var err error
	fileAttributeMatcherToStringFunction_Once.Do(func() {
		err = fileAttributeMatcherStruct_Set()
		if err != nil {
			return
		}
		fileAttributeMatcherToStringFunction, err = fileAttributeMatcherStruct.InvokerNew("to_string")
	})
	return err
}

// ToString is a representation of the C type g_file_attribute_matcher_to_string.
func (recv *FileAttributeMatcher) ToString() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := fileAttributeMatcherToStringFunction_Set()
	if err == nil {
		ret = fileAttributeMatcherToStringFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

var fileAttributeMatcherUnrefFunction *gi.Function
var fileAttributeMatcherUnrefFunction_Once sync.Once

func fileAttributeMatcherUnrefFunction_Set() error {
	var err error
	fileAttributeMatcherUnrefFunction_Once.Do(func() {
		err = fileAttributeMatcherStruct_Set()
		if err != nil {
			return
		}
		fileAttributeMatcherUnrefFunction, err = fileAttributeMatcherStruct.InvokerNew("unref")
	})
	return err
}

// Unref is a representation of the C type g_file_attribute_matcher_unref.
func (recv *FileAttributeMatcher) Unref() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := fileAttributeMatcherUnrefFunction_Set()
	if err == nil {
		fileAttributeMatcherUnrefFunction.Invoke(inArgs[:], nil)
	}

	return
}

var fileDescriptorBasedIfaceStruct *gi.Struct
var fileDescriptorBasedIfaceStruct_Once sync.Once

func fileDescriptorBasedIfaceStruct_Set() error {
	var err error
	fileDescriptorBasedIfaceStruct_Once.Do(func() {
		fileDescriptorBasedIfaceStruct, err = gi.StructNew("Gio", "FileDescriptorBasedIface")
	})
	return err
}

type FileDescriptorBasedIface struct {
	native unsafe.Pointer
}

func FileDescriptorBasedIfaceNewFromNative(native unsafe.Pointer) *FileDescriptorBasedIface {
	instance := &FileDescriptorBasedIface{native: native}

	return instance
}

/*
CastToFileDescriptorBasedIface down casts any arbitrary Object to FileDescriptorBasedIface.
Exercise care, as this is a potentially dangerous function
if the Object is not a FileDescriptorBasedIface.
*/
func (recv *FileDescriptorBasedIface) CastToFileDescriptorBasedIface(object *gobject.Object) *FileDescriptorBasedIface {
	return FileDescriptorBasedIfaceNewFromNative(object.Native())
}

func (recv *FileDescriptorBasedIface) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'g_iface' : for field getter : no Go type for 'GObject.TypeInterface'

// UNSUPPORTED : C value 'g_iface' : for field setter : no Go type for 'GObject.TypeInterface'

// UNSUPPORTED : C value 'get_fd' : for field getter : missing Type

// UNSUPPORTED : C value 'get_fd' : for field setter : missing Type

// FileDescriptorBasedIfaceStruct creates an uninitialised FileDescriptorBasedIface.
func FileDescriptorBasedIfaceStruct() *FileDescriptorBasedIface {
	err := fileDescriptorBasedIfaceStruct_Set()
	if err != nil {
		return nil
	}

	structGo := FileDescriptorBasedIfaceNewFromNative(fileDescriptorBasedIfaceStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeFileDescriptorBasedIface)
	return structGo
}
func finalizeFileDescriptorBasedIface(obj *FileDescriptorBasedIface) {
	fileDescriptorBasedIfaceStruct.Free(obj.native)
}

var fileEnumeratorClassStruct *gi.Struct
var fileEnumeratorClassStruct_Once sync.Once

func fileEnumeratorClassStruct_Set() error {
	var err error
	fileEnumeratorClassStruct_Once.Do(func() {
		fileEnumeratorClassStruct, err = gi.StructNew("Gio", "FileEnumeratorClass")
	})
	return err
}

type FileEnumeratorClass struct {
	native unsafe.Pointer
}

func FileEnumeratorClassNewFromNative(native unsafe.Pointer) *FileEnumeratorClass {
	instance := &FileEnumeratorClass{native: native}

	return instance
}

/*
CastToFileEnumeratorClass down casts any arbitrary Object to FileEnumeratorClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a FileEnumeratorClass.
*/
func (recv *FileEnumeratorClass) CastToFileEnumeratorClass(object *gobject.Object) *FileEnumeratorClass {
	return FileEnumeratorClassNewFromNative(object.Native())
}

func (recv *FileEnumeratorClass) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'next_file' : for field getter : missing Type

// UNSUPPORTED : C value 'next_file' : for field setter : missing Type

// UNSUPPORTED : C value 'close_fn' : for field getter : missing Type

// UNSUPPORTED : C value 'close_fn' : for field setter : missing Type

// UNSUPPORTED : C value 'next_files_async' : for field getter : missing Type

// UNSUPPORTED : C value 'next_files_async' : for field setter : missing Type

// UNSUPPORTED : C value 'next_files_finish' : for field getter : missing Type

// UNSUPPORTED : C value 'next_files_finish' : for field setter : missing Type

// UNSUPPORTED : C value 'close_async' : for field getter : missing Type

// UNSUPPORTED : C value 'close_async' : for field setter : missing Type

// UNSUPPORTED : C value 'close_finish' : for field getter : missing Type

// UNSUPPORTED : C value 'close_finish' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved4' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved5' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved5' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved6' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved6' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved7' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved7' : for field setter : missing Type

// FileEnumeratorClassStruct creates an uninitialised FileEnumeratorClass.
func FileEnumeratorClassStruct() *FileEnumeratorClass {
	err := fileEnumeratorClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := FileEnumeratorClassNewFromNative(fileEnumeratorClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeFileEnumeratorClass)
	return structGo
}
func finalizeFileEnumeratorClass(obj *FileEnumeratorClass) {
	fileEnumeratorClassStruct.Free(obj.native)
}

var fileEnumeratorPrivateStruct *gi.Struct
var fileEnumeratorPrivateStruct_Once sync.Once

func fileEnumeratorPrivateStruct_Set() error {
	var err error
	fileEnumeratorPrivateStruct_Once.Do(func() {
		fileEnumeratorPrivateStruct, err = gi.StructNew("Gio", "FileEnumeratorPrivate")
	})
	return err
}

type FileEnumeratorPrivate struct {
	native unsafe.Pointer
}

func FileEnumeratorPrivateNewFromNative(native unsafe.Pointer) *FileEnumeratorPrivate {
	instance := &FileEnumeratorPrivate{native: native}

	return instance
}

/*
CastToFileEnumeratorPrivate down casts any arbitrary Object to FileEnumeratorPrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a FileEnumeratorPrivate.
*/
func (recv *FileEnumeratorPrivate) CastToFileEnumeratorPrivate(object *gobject.Object) *FileEnumeratorPrivate {
	return FileEnumeratorPrivateNewFromNative(object.Native())
}

func (recv *FileEnumeratorPrivate) Native() unsafe.Pointer {
	return recv.native
}

// FileEnumeratorPrivateStruct creates an uninitialised FileEnumeratorPrivate.
func FileEnumeratorPrivateStruct() *FileEnumeratorPrivate {
	err := fileEnumeratorPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := FileEnumeratorPrivateNewFromNative(fileEnumeratorPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeFileEnumeratorPrivate)
	return structGo
}
func finalizeFileEnumeratorPrivate(obj *FileEnumeratorPrivate) {
	fileEnumeratorPrivateStruct.Free(obj.native)
}

var fileIOStreamClassStruct *gi.Struct
var fileIOStreamClassStruct_Once sync.Once

func fileIOStreamClassStruct_Set() error {
	var err error
	fileIOStreamClassStruct_Once.Do(func() {
		fileIOStreamClassStruct, err = gi.StructNew("Gio", "FileIOStreamClass")
	})
	return err
}

type FileIOStreamClass struct {
	native unsafe.Pointer
}

func FileIOStreamClassNewFromNative(native unsafe.Pointer) *FileIOStreamClass {
	instance := &FileIOStreamClass{native: native}

	return instance
}

/*
CastToFileIOStreamClass down casts any arbitrary Object to FileIOStreamClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a FileIOStreamClass.
*/
func (recv *FileIOStreamClass) CastToFileIOStreamClass(object *gobject.Object) *FileIOStreamClass {
	return FileIOStreamClassNewFromNative(object.Native())
}

func (recv *FileIOStreamClass) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentClass returns the C field 'parent_class'.
func (recv *FileIOStreamClass) FieldParentClass() *IOStreamClass {
	argValue := gi.StructFieldGet(fileIOStreamClassStruct, recv.native, "parent_class")
	value := IOStreamClassNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *FileIOStreamClass) SetFieldParentClass(value *IOStreamClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.StructFieldSet(fileIOStreamClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value 'tell' : for field getter : missing Type

// UNSUPPORTED : C value 'tell' : for field setter : missing Type

// UNSUPPORTED : C value 'can_seek' : for field getter : missing Type

// UNSUPPORTED : C value 'can_seek' : for field setter : missing Type

// UNSUPPORTED : C value 'seek' : for field getter : missing Type

// UNSUPPORTED : C value 'seek' : for field setter : missing Type

// UNSUPPORTED : C value 'can_truncate' : for field getter : missing Type

// UNSUPPORTED : C value 'can_truncate' : for field setter : missing Type

// UNSUPPORTED : C value 'truncate_fn' : for field getter : missing Type

// UNSUPPORTED : C value 'truncate_fn' : for field setter : missing Type

// UNSUPPORTED : C value 'query_info' : for field getter : missing Type

// UNSUPPORTED : C value 'query_info' : for field setter : missing Type

// UNSUPPORTED : C value 'query_info_async' : for field getter : missing Type

// UNSUPPORTED : C value 'query_info_async' : for field setter : missing Type

// UNSUPPORTED : C value 'query_info_finish' : for field getter : missing Type

// UNSUPPORTED : C value 'query_info_finish' : for field setter : missing Type

// UNSUPPORTED : C value 'get_etag' : for field getter : missing Type

// UNSUPPORTED : C value 'get_etag' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved4' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved5' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved5' : for field setter : missing Type

// FileIOStreamClassStruct creates an uninitialised FileIOStreamClass.
func FileIOStreamClassStruct() *FileIOStreamClass {
	err := fileIOStreamClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := FileIOStreamClassNewFromNative(fileIOStreamClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeFileIOStreamClass)
	return structGo
}
func finalizeFileIOStreamClass(obj *FileIOStreamClass) {
	fileIOStreamClassStruct.Free(obj.native)
}

var fileIOStreamPrivateStruct *gi.Struct
var fileIOStreamPrivateStruct_Once sync.Once

func fileIOStreamPrivateStruct_Set() error {
	var err error
	fileIOStreamPrivateStruct_Once.Do(func() {
		fileIOStreamPrivateStruct, err = gi.StructNew("Gio", "FileIOStreamPrivate")
	})
	return err
}

type FileIOStreamPrivate struct {
	native unsafe.Pointer
}

func FileIOStreamPrivateNewFromNative(native unsafe.Pointer) *FileIOStreamPrivate {
	instance := &FileIOStreamPrivate{native: native}

	return instance
}

/*
CastToFileIOStreamPrivate down casts any arbitrary Object to FileIOStreamPrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a FileIOStreamPrivate.
*/
func (recv *FileIOStreamPrivate) CastToFileIOStreamPrivate(object *gobject.Object) *FileIOStreamPrivate {
	return FileIOStreamPrivateNewFromNative(object.Native())
}

func (recv *FileIOStreamPrivate) Native() unsafe.Pointer {
	return recv.native
}

// FileIOStreamPrivateStruct creates an uninitialised FileIOStreamPrivate.
func FileIOStreamPrivateStruct() *FileIOStreamPrivate {
	err := fileIOStreamPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := FileIOStreamPrivateNewFromNative(fileIOStreamPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeFileIOStreamPrivate)
	return structGo
}
func finalizeFileIOStreamPrivate(obj *FileIOStreamPrivate) {
	fileIOStreamPrivateStruct.Free(obj.native)
}

var fileIconClassStruct *gi.Struct
var fileIconClassStruct_Once sync.Once

func fileIconClassStruct_Set() error {
	var err error
	fileIconClassStruct_Once.Do(func() {
		fileIconClassStruct, err = gi.StructNew("Gio", "FileIconClass")
	})
	return err
}

type FileIconClass struct {
	native unsafe.Pointer
}

func FileIconClassNewFromNative(native unsafe.Pointer) *FileIconClass {
	instance := &FileIconClass{native: native}

	return instance
}

/*
CastToFileIconClass down casts any arbitrary Object to FileIconClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a FileIconClass.
*/
func (recv *FileIconClass) CastToFileIconClass(object *gobject.Object) *FileIconClass {
	return FileIconClassNewFromNative(object.Native())
}

func (recv *FileIconClass) Native() unsafe.Pointer {
	return recv.native
}

// FileIconClassStruct creates an uninitialised FileIconClass.
func FileIconClassStruct() *FileIconClass {
	err := fileIconClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := FileIconClassNewFromNative(fileIconClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeFileIconClass)
	return structGo
}
func finalizeFileIconClass(obj *FileIconClass) {
	fileIconClassStruct.Free(obj.native)
}

var fileIfaceStruct *gi.Struct
var fileIfaceStruct_Once sync.Once

func fileIfaceStruct_Set() error {
	var err error
	fileIfaceStruct_Once.Do(func() {
		fileIfaceStruct, err = gi.StructNew("Gio", "FileIface")
	})
	return err
}

type FileIface struct {
	native unsafe.Pointer
}

func FileIfaceNewFromNative(native unsafe.Pointer) *FileIface {
	instance := &FileIface{native: native}

	return instance
}

/*
CastToFileIface down casts any arbitrary Object to FileIface.
Exercise care, as this is a potentially dangerous function
if the Object is not a FileIface.
*/
func (recv *FileIface) CastToFileIface(object *gobject.Object) *FileIface {
	return FileIfaceNewFromNative(object.Native())
}

func (recv *FileIface) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'g_iface' : for field getter : no Go type for 'GObject.TypeInterface'

// UNSUPPORTED : C value 'g_iface' : for field setter : no Go type for 'GObject.TypeInterface'

// UNSUPPORTED : C value 'dup' : for field getter : missing Type

// UNSUPPORTED : C value 'dup' : for field setter : missing Type

// UNSUPPORTED : C value 'hash' : for field getter : missing Type

// UNSUPPORTED : C value 'hash' : for field setter : missing Type

// UNSUPPORTED : C value 'equal' : for field getter : missing Type

// UNSUPPORTED : C value 'equal' : for field setter : missing Type

// UNSUPPORTED : C value 'is_native' : for field getter : missing Type

// UNSUPPORTED : C value 'is_native' : for field setter : missing Type

// UNSUPPORTED : C value 'has_uri_scheme' : for field getter : missing Type

// UNSUPPORTED : C value 'has_uri_scheme' : for field setter : missing Type

// UNSUPPORTED : C value 'get_uri_scheme' : for field getter : missing Type

// UNSUPPORTED : C value 'get_uri_scheme' : for field setter : missing Type

// UNSUPPORTED : C value 'get_basename' : for field getter : missing Type

// UNSUPPORTED : C value 'get_basename' : for field setter : missing Type

// UNSUPPORTED : C value 'get_path' : for field getter : missing Type

// UNSUPPORTED : C value 'get_path' : for field setter : missing Type

// UNSUPPORTED : C value 'get_uri' : for field getter : missing Type

// UNSUPPORTED : C value 'get_uri' : for field setter : missing Type

// UNSUPPORTED : C value 'get_parse_name' : for field getter : missing Type

// UNSUPPORTED : C value 'get_parse_name' : for field setter : missing Type

// UNSUPPORTED : C value 'get_parent' : for field getter : missing Type

// UNSUPPORTED : C value 'get_parent' : for field setter : missing Type

// UNSUPPORTED : C value 'prefix_matches' : for field getter : missing Type

// UNSUPPORTED : C value 'prefix_matches' : for field setter : missing Type

// UNSUPPORTED : C value 'get_relative_path' : for field getter : missing Type

// UNSUPPORTED : C value 'get_relative_path' : for field setter : missing Type

// UNSUPPORTED : C value 'resolve_relative_path' : for field getter : missing Type

// UNSUPPORTED : C value 'resolve_relative_path' : for field setter : missing Type

// UNSUPPORTED : C value 'get_child_for_display_name' : for field getter : missing Type

// UNSUPPORTED : C value 'get_child_for_display_name' : for field setter : missing Type

// UNSUPPORTED : C value 'enumerate_children' : for field getter : missing Type

// UNSUPPORTED : C value 'enumerate_children' : for field setter : missing Type

// UNSUPPORTED : C value 'enumerate_children_async' : for field getter : missing Type

// UNSUPPORTED : C value 'enumerate_children_async' : for field setter : missing Type

// UNSUPPORTED : C value 'enumerate_children_finish' : for field getter : missing Type

// UNSUPPORTED : C value 'enumerate_children_finish' : for field setter : missing Type

// UNSUPPORTED : C value 'query_info' : for field getter : missing Type

// UNSUPPORTED : C value 'query_info' : for field setter : missing Type

// UNSUPPORTED : C value 'query_info_async' : for field getter : missing Type

// UNSUPPORTED : C value 'query_info_async' : for field setter : missing Type

// UNSUPPORTED : C value 'query_info_finish' : for field getter : missing Type

// UNSUPPORTED : C value 'query_info_finish' : for field setter : missing Type

// UNSUPPORTED : C value 'query_filesystem_info' : for field getter : missing Type

// UNSUPPORTED : C value 'query_filesystem_info' : for field setter : missing Type

// UNSUPPORTED : C value 'query_filesystem_info_async' : for field getter : missing Type

// UNSUPPORTED : C value 'query_filesystem_info_async' : for field setter : missing Type

// UNSUPPORTED : C value 'query_filesystem_info_finish' : for field getter : missing Type

// UNSUPPORTED : C value 'query_filesystem_info_finish' : for field setter : missing Type

// UNSUPPORTED : C value 'find_enclosing_mount' : for field getter : missing Type

// UNSUPPORTED : C value 'find_enclosing_mount' : for field setter : missing Type

// UNSUPPORTED : C value 'find_enclosing_mount_async' : for field getter : missing Type

// UNSUPPORTED : C value 'find_enclosing_mount_async' : for field setter : missing Type

// UNSUPPORTED : C value 'find_enclosing_mount_finish' : for field getter : missing Type

// UNSUPPORTED : C value 'find_enclosing_mount_finish' : for field setter : missing Type

// UNSUPPORTED : C value 'set_display_name' : for field getter : missing Type

// UNSUPPORTED : C value 'set_display_name' : for field setter : missing Type

// UNSUPPORTED : C value 'set_display_name_async' : for field getter : missing Type

// UNSUPPORTED : C value 'set_display_name_async' : for field setter : missing Type

// UNSUPPORTED : C value 'set_display_name_finish' : for field getter : missing Type

// UNSUPPORTED : C value 'set_display_name_finish' : for field setter : missing Type

// UNSUPPORTED : C value 'query_settable_attributes' : for field getter : missing Type

// UNSUPPORTED : C value 'query_settable_attributes' : for field setter : missing Type

// UNSUPPORTED : C value '_query_settable_attributes_async' : for field getter : missing Type

// UNSUPPORTED : C value '_query_settable_attributes_async' : for field setter : missing Type

// UNSUPPORTED : C value '_query_settable_attributes_finish' : for field getter : missing Type

// UNSUPPORTED : C value '_query_settable_attributes_finish' : for field setter : missing Type

// UNSUPPORTED : C value 'query_writable_namespaces' : for field getter : missing Type

// UNSUPPORTED : C value 'query_writable_namespaces' : for field setter : missing Type

// UNSUPPORTED : C value '_query_writable_namespaces_async' : for field getter : missing Type

// UNSUPPORTED : C value '_query_writable_namespaces_async' : for field setter : missing Type

// UNSUPPORTED : C value '_query_writable_namespaces_finish' : for field getter : missing Type

// UNSUPPORTED : C value '_query_writable_namespaces_finish' : for field setter : missing Type

// UNSUPPORTED : C value 'set_attribute' : for field getter : missing Type

// UNSUPPORTED : C value 'set_attribute' : for field setter : missing Type

// UNSUPPORTED : C value 'set_attributes_from_info' : for field getter : missing Type

// UNSUPPORTED : C value 'set_attributes_from_info' : for field setter : missing Type

// UNSUPPORTED : C value 'set_attributes_async' : for field getter : missing Type

// UNSUPPORTED : C value 'set_attributes_async' : for field setter : missing Type

// UNSUPPORTED : C value 'set_attributes_finish' : for field getter : missing Type

// UNSUPPORTED : C value 'set_attributes_finish' : for field setter : missing Type

// UNSUPPORTED : C value 'read_fn' : for field getter : missing Type

// UNSUPPORTED : C value 'read_fn' : for field setter : missing Type

// UNSUPPORTED : C value 'read_async' : for field getter : missing Type

// UNSUPPORTED : C value 'read_async' : for field setter : missing Type

// UNSUPPORTED : C value 'read_finish' : for field getter : missing Type

// UNSUPPORTED : C value 'read_finish' : for field setter : missing Type

// UNSUPPORTED : C value 'append_to' : for field getter : missing Type

// UNSUPPORTED : C value 'append_to' : for field setter : missing Type

// UNSUPPORTED : C value 'append_to_async' : for field getter : missing Type

// UNSUPPORTED : C value 'append_to_async' : for field setter : missing Type

// UNSUPPORTED : C value 'append_to_finish' : for field getter : missing Type

// UNSUPPORTED : C value 'append_to_finish' : for field setter : missing Type

// UNSUPPORTED : C value 'create' : for field getter : missing Type

// UNSUPPORTED : C value 'create' : for field setter : missing Type

// UNSUPPORTED : C value 'create_async' : for field getter : missing Type

// UNSUPPORTED : C value 'create_async' : for field setter : missing Type

// UNSUPPORTED : C value 'create_finish' : for field getter : missing Type

// UNSUPPORTED : C value 'create_finish' : for field setter : missing Type

// UNSUPPORTED : C value 'replace' : for field getter : missing Type

// UNSUPPORTED : C value 'replace' : for field setter : missing Type

// UNSUPPORTED : C value 'replace_async' : for field getter : missing Type

// UNSUPPORTED : C value 'replace_async' : for field setter : missing Type

// UNSUPPORTED : C value 'replace_finish' : for field getter : missing Type

// UNSUPPORTED : C value 'replace_finish' : for field setter : missing Type

// UNSUPPORTED : C value 'delete_file' : for field getter : missing Type

// UNSUPPORTED : C value 'delete_file' : for field setter : missing Type

// UNSUPPORTED : C value 'delete_file_async' : for field getter : missing Type

// UNSUPPORTED : C value 'delete_file_async' : for field setter : missing Type

// UNSUPPORTED : C value 'delete_file_finish' : for field getter : missing Type

// UNSUPPORTED : C value 'delete_file_finish' : for field setter : missing Type

// UNSUPPORTED : C value 'trash' : for field getter : missing Type

// UNSUPPORTED : C value 'trash' : for field setter : missing Type

// UNSUPPORTED : C value 'trash_async' : for field getter : missing Type

// UNSUPPORTED : C value 'trash_async' : for field setter : missing Type

// UNSUPPORTED : C value 'trash_finish' : for field getter : missing Type

// UNSUPPORTED : C value 'trash_finish' : for field setter : missing Type

// UNSUPPORTED : C value 'make_directory' : for field getter : missing Type

// UNSUPPORTED : C value 'make_directory' : for field setter : missing Type

// UNSUPPORTED : C value 'make_directory_async' : for field getter : missing Type

// UNSUPPORTED : C value 'make_directory_async' : for field setter : missing Type

// UNSUPPORTED : C value 'make_directory_finish' : for field getter : missing Type

// UNSUPPORTED : C value 'make_directory_finish' : for field setter : missing Type

// UNSUPPORTED : C value 'make_symbolic_link' : for field getter : missing Type

// UNSUPPORTED : C value 'make_symbolic_link' : for field setter : missing Type

// UNSUPPORTED : C value '_make_symbolic_link_async' : for field getter : missing Type

// UNSUPPORTED : C value '_make_symbolic_link_async' : for field setter : missing Type

// UNSUPPORTED : C value '_make_symbolic_link_finish' : for field getter : missing Type

// UNSUPPORTED : C value '_make_symbolic_link_finish' : for field setter : missing Type

// UNSUPPORTED : C value 'copy' : for field getter : missing Type

// UNSUPPORTED : C value 'copy' : for field setter : missing Type

// UNSUPPORTED : C value 'copy_async' : for field getter : missing Type

// UNSUPPORTED : C value 'copy_async' : for field setter : missing Type

// UNSUPPORTED : C value 'copy_finish' : for field getter : missing Type

// UNSUPPORTED : C value 'copy_finish' : for field setter : missing Type

// UNSUPPORTED : C value 'move' : for field getter : missing Type

// UNSUPPORTED : C value 'move' : for field setter : missing Type

// UNSUPPORTED : C value '_move_async' : for field getter : missing Type

// UNSUPPORTED : C value '_move_async' : for field setter : missing Type

// UNSUPPORTED : C value '_move_finish' : for field getter : missing Type

// UNSUPPORTED : C value '_move_finish' : for field setter : missing Type

// UNSUPPORTED : C value 'mount_mountable' : for field getter : missing Type

// UNSUPPORTED : C value 'mount_mountable' : for field setter : missing Type

// UNSUPPORTED : C value 'mount_mountable_finish' : for field getter : missing Type

// UNSUPPORTED : C value 'mount_mountable_finish' : for field setter : missing Type

// UNSUPPORTED : C value 'unmount_mountable' : for field getter : missing Type

// UNSUPPORTED : C value 'unmount_mountable' : for field setter : missing Type

// UNSUPPORTED : C value 'unmount_mountable_finish' : for field getter : missing Type

// UNSUPPORTED : C value 'unmount_mountable_finish' : for field setter : missing Type

// UNSUPPORTED : C value 'eject_mountable' : for field getter : missing Type

// UNSUPPORTED : C value 'eject_mountable' : for field setter : missing Type

// UNSUPPORTED : C value 'eject_mountable_finish' : for field getter : missing Type

// UNSUPPORTED : C value 'eject_mountable_finish' : for field setter : missing Type

// UNSUPPORTED : C value 'mount_enclosing_volume' : for field getter : missing Type

// UNSUPPORTED : C value 'mount_enclosing_volume' : for field setter : missing Type

// UNSUPPORTED : C value 'mount_enclosing_volume_finish' : for field getter : missing Type

// UNSUPPORTED : C value 'mount_enclosing_volume_finish' : for field setter : missing Type

// UNSUPPORTED : C value 'monitor_dir' : for field getter : missing Type

// UNSUPPORTED : C value 'monitor_dir' : for field setter : missing Type

// UNSUPPORTED : C value 'monitor_file' : for field getter : missing Type

// UNSUPPORTED : C value 'monitor_file' : for field setter : missing Type

// UNSUPPORTED : C value 'open_readwrite' : for field getter : missing Type

// UNSUPPORTED : C value 'open_readwrite' : for field setter : missing Type

// UNSUPPORTED : C value 'open_readwrite_async' : for field getter : missing Type

// UNSUPPORTED : C value 'open_readwrite_async' : for field setter : missing Type

// UNSUPPORTED : C value 'open_readwrite_finish' : for field getter : missing Type

// UNSUPPORTED : C value 'open_readwrite_finish' : for field setter : missing Type

// UNSUPPORTED : C value 'create_readwrite' : for field getter : missing Type

// UNSUPPORTED : C value 'create_readwrite' : for field setter : missing Type

// UNSUPPORTED : C value 'create_readwrite_async' : for field getter : missing Type

// UNSUPPORTED : C value 'create_readwrite_async' : for field setter : missing Type

// UNSUPPORTED : C value 'create_readwrite_finish' : for field getter : missing Type

// UNSUPPORTED : C value 'create_readwrite_finish' : for field setter : missing Type

// UNSUPPORTED : C value 'replace_readwrite' : for field getter : missing Type

// UNSUPPORTED : C value 'replace_readwrite' : for field setter : missing Type

// UNSUPPORTED : C value 'replace_readwrite_async' : for field getter : missing Type

// UNSUPPORTED : C value 'replace_readwrite_async' : for field setter : missing Type

// UNSUPPORTED : C value 'replace_readwrite_finish' : for field getter : missing Type

// UNSUPPORTED : C value 'replace_readwrite_finish' : for field setter : missing Type

// UNSUPPORTED : C value 'start_mountable' : for field getter : missing Type

// UNSUPPORTED : C value 'start_mountable' : for field setter : missing Type

// UNSUPPORTED : C value 'start_mountable_finish' : for field getter : missing Type

// UNSUPPORTED : C value 'start_mountable_finish' : for field setter : missing Type

// UNSUPPORTED : C value 'stop_mountable' : for field getter : missing Type

// UNSUPPORTED : C value 'stop_mountable' : for field setter : missing Type

// UNSUPPORTED : C value 'stop_mountable_finish' : for field getter : missing Type

// UNSUPPORTED : C value 'stop_mountable_finish' : for field setter : missing Type

// FieldSupportsThreadContexts returns the C field 'supports_thread_contexts'.
func (recv *FileIface) FieldSupportsThreadContexts() bool {
	argValue := gi.StructFieldGet(fileIfaceStruct, recv.native, "supports_thread_contexts")
	value := argValue.Boolean()
	return value
}

// SetFieldSupportsThreadContexts sets the value of the C field 'supports_thread_contexts'.
func (recv *FileIface) SetFieldSupportsThreadContexts(value bool) {
	var argValue gi.Argument
	argValue.SetBoolean(value)
	gi.StructFieldSet(fileIfaceStruct, recv.native, "supports_thread_contexts", argValue)
}

// UNSUPPORTED : C value 'unmount_mountable_with_operation' : for field getter : missing Type

// UNSUPPORTED : C value 'unmount_mountable_with_operation' : for field setter : missing Type

// UNSUPPORTED : C value 'unmount_mountable_with_operation_finish' : for field getter : missing Type

// UNSUPPORTED : C value 'unmount_mountable_with_operation_finish' : for field setter : missing Type

// UNSUPPORTED : C value 'eject_mountable_with_operation' : for field getter : missing Type

// UNSUPPORTED : C value 'eject_mountable_with_operation' : for field setter : missing Type

// UNSUPPORTED : C value 'eject_mountable_with_operation_finish' : for field getter : missing Type

// UNSUPPORTED : C value 'eject_mountable_with_operation_finish' : for field setter : missing Type

// UNSUPPORTED : C value 'poll_mountable' : for field getter : missing Type

// UNSUPPORTED : C value 'poll_mountable' : for field setter : missing Type

// UNSUPPORTED : C value 'poll_mountable_finish' : for field getter : missing Type

// UNSUPPORTED : C value 'poll_mountable_finish' : for field setter : missing Type

// UNSUPPORTED : C value 'measure_disk_usage' : for field getter : missing Type

// UNSUPPORTED : C value 'measure_disk_usage' : for field setter : missing Type

// UNSUPPORTED : C value 'measure_disk_usage_async' : for field getter : missing Type

// UNSUPPORTED : C value 'measure_disk_usage_async' : for field setter : missing Type

// UNSUPPORTED : C value 'measure_disk_usage_finish' : for field getter : missing Type

// UNSUPPORTED : C value 'measure_disk_usage_finish' : for field setter : missing Type

// FileIfaceStruct creates an uninitialised FileIface.
func FileIfaceStruct() *FileIface {
	err := fileIfaceStruct_Set()
	if err != nil {
		return nil
	}

	structGo := FileIfaceNewFromNative(fileIfaceStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeFileIface)
	return structGo
}
func finalizeFileIface(obj *FileIface) {
	fileIfaceStruct.Free(obj.native)
}

var fileInfoClassStruct *gi.Struct
var fileInfoClassStruct_Once sync.Once

func fileInfoClassStruct_Set() error {
	var err error
	fileInfoClassStruct_Once.Do(func() {
		fileInfoClassStruct, err = gi.StructNew("Gio", "FileInfoClass")
	})
	return err
}

type FileInfoClass struct {
	native unsafe.Pointer
}

func FileInfoClassNewFromNative(native unsafe.Pointer) *FileInfoClass {
	instance := &FileInfoClass{native: native}

	return instance
}

/*
CastToFileInfoClass down casts any arbitrary Object to FileInfoClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a FileInfoClass.
*/
func (recv *FileInfoClass) CastToFileInfoClass(object *gobject.Object) *FileInfoClass {
	return FileInfoClassNewFromNative(object.Native())
}

func (recv *FileInfoClass) Native() unsafe.Pointer {
	return recv.native
}

// FileInfoClassStruct creates an uninitialised FileInfoClass.
func FileInfoClassStruct() *FileInfoClass {
	err := fileInfoClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := FileInfoClassNewFromNative(fileInfoClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeFileInfoClass)
	return structGo
}
func finalizeFileInfoClass(obj *FileInfoClass) {
	fileInfoClassStruct.Free(obj.native)
}

var fileInputStreamClassStruct *gi.Struct
var fileInputStreamClassStruct_Once sync.Once

func fileInputStreamClassStruct_Set() error {
	var err error
	fileInputStreamClassStruct_Once.Do(func() {
		fileInputStreamClassStruct, err = gi.StructNew("Gio", "FileInputStreamClass")
	})
	return err
}

type FileInputStreamClass struct {
	native unsafe.Pointer
}

func FileInputStreamClassNewFromNative(native unsafe.Pointer) *FileInputStreamClass {
	instance := &FileInputStreamClass{native: native}

	return instance
}

/*
CastToFileInputStreamClass down casts any arbitrary Object to FileInputStreamClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a FileInputStreamClass.
*/
func (recv *FileInputStreamClass) CastToFileInputStreamClass(object *gobject.Object) *FileInputStreamClass {
	return FileInputStreamClassNewFromNative(object.Native())
}

func (recv *FileInputStreamClass) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentClass returns the C field 'parent_class'.
func (recv *FileInputStreamClass) FieldParentClass() *InputStreamClass {
	argValue := gi.StructFieldGet(fileInputStreamClassStruct, recv.native, "parent_class")
	value := InputStreamClassNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *FileInputStreamClass) SetFieldParentClass(value *InputStreamClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.StructFieldSet(fileInputStreamClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value 'tell' : for field getter : missing Type

// UNSUPPORTED : C value 'tell' : for field setter : missing Type

// UNSUPPORTED : C value 'can_seek' : for field getter : missing Type

// UNSUPPORTED : C value 'can_seek' : for field setter : missing Type

// UNSUPPORTED : C value 'seek' : for field getter : missing Type

// UNSUPPORTED : C value 'seek' : for field setter : missing Type

// UNSUPPORTED : C value 'query_info' : for field getter : missing Type

// UNSUPPORTED : C value 'query_info' : for field setter : missing Type

// UNSUPPORTED : C value 'query_info_async' : for field getter : missing Type

// UNSUPPORTED : C value 'query_info_async' : for field setter : missing Type

// UNSUPPORTED : C value 'query_info_finish' : for field getter : missing Type

// UNSUPPORTED : C value 'query_info_finish' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved4' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved5' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved5' : for field setter : missing Type

// FileInputStreamClassStruct creates an uninitialised FileInputStreamClass.
func FileInputStreamClassStruct() *FileInputStreamClass {
	err := fileInputStreamClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := FileInputStreamClassNewFromNative(fileInputStreamClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeFileInputStreamClass)
	return structGo
}
func finalizeFileInputStreamClass(obj *FileInputStreamClass) {
	fileInputStreamClassStruct.Free(obj.native)
}

var fileInputStreamPrivateStruct *gi.Struct
var fileInputStreamPrivateStruct_Once sync.Once

func fileInputStreamPrivateStruct_Set() error {
	var err error
	fileInputStreamPrivateStruct_Once.Do(func() {
		fileInputStreamPrivateStruct, err = gi.StructNew("Gio", "FileInputStreamPrivate")
	})
	return err
}

type FileInputStreamPrivate struct {
	native unsafe.Pointer
}

func FileInputStreamPrivateNewFromNative(native unsafe.Pointer) *FileInputStreamPrivate {
	instance := &FileInputStreamPrivate{native: native}

	return instance
}

/*
CastToFileInputStreamPrivate down casts any arbitrary Object to FileInputStreamPrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a FileInputStreamPrivate.
*/
func (recv *FileInputStreamPrivate) CastToFileInputStreamPrivate(object *gobject.Object) *FileInputStreamPrivate {
	return FileInputStreamPrivateNewFromNative(object.Native())
}

func (recv *FileInputStreamPrivate) Native() unsafe.Pointer {
	return recv.native
}

// FileInputStreamPrivateStruct creates an uninitialised FileInputStreamPrivate.
func FileInputStreamPrivateStruct() *FileInputStreamPrivate {
	err := fileInputStreamPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := FileInputStreamPrivateNewFromNative(fileInputStreamPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeFileInputStreamPrivate)
	return structGo
}
func finalizeFileInputStreamPrivate(obj *FileInputStreamPrivate) {
	fileInputStreamPrivateStruct.Free(obj.native)
}

var fileMonitorClassStruct *gi.Struct
var fileMonitorClassStruct_Once sync.Once

func fileMonitorClassStruct_Set() error {
	var err error
	fileMonitorClassStruct_Once.Do(func() {
		fileMonitorClassStruct, err = gi.StructNew("Gio", "FileMonitorClass")
	})
	return err
}

type FileMonitorClass struct {
	native unsafe.Pointer
}

func FileMonitorClassNewFromNative(native unsafe.Pointer) *FileMonitorClass {
	instance := &FileMonitorClass{native: native}

	return instance
}

/*
CastToFileMonitorClass down casts any arbitrary Object to FileMonitorClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a FileMonitorClass.
*/
func (recv *FileMonitorClass) CastToFileMonitorClass(object *gobject.Object) *FileMonitorClass {
	return FileMonitorClassNewFromNative(object.Native())
}

func (recv *FileMonitorClass) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'changed' : for field getter : missing Type

// UNSUPPORTED : C value 'changed' : for field setter : missing Type

// UNSUPPORTED : C value 'cancel' : for field getter : missing Type

// UNSUPPORTED : C value 'cancel' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved4' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved5' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved5' : for field setter : missing Type

// FileMonitorClassStruct creates an uninitialised FileMonitorClass.
func FileMonitorClassStruct() *FileMonitorClass {
	err := fileMonitorClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := FileMonitorClassNewFromNative(fileMonitorClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeFileMonitorClass)
	return structGo
}
func finalizeFileMonitorClass(obj *FileMonitorClass) {
	fileMonitorClassStruct.Free(obj.native)
}

var fileMonitorPrivateStruct *gi.Struct
var fileMonitorPrivateStruct_Once sync.Once

func fileMonitorPrivateStruct_Set() error {
	var err error
	fileMonitorPrivateStruct_Once.Do(func() {
		fileMonitorPrivateStruct, err = gi.StructNew("Gio", "FileMonitorPrivate")
	})
	return err
}

type FileMonitorPrivate struct {
	native unsafe.Pointer
}

func FileMonitorPrivateNewFromNative(native unsafe.Pointer) *FileMonitorPrivate {
	instance := &FileMonitorPrivate{native: native}

	return instance
}

/*
CastToFileMonitorPrivate down casts any arbitrary Object to FileMonitorPrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a FileMonitorPrivate.
*/
func (recv *FileMonitorPrivate) CastToFileMonitorPrivate(object *gobject.Object) *FileMonitorPrivate {
	return FileMonitorPrivateNewFromNative(object.Native())
}

func (recv *FileMonitorPrivate) Native() unsafe.Pointer {
	return recv.native
}

// FileMonitorPrivateStruct creates an uninitialised FileMonitorPrivate.
func FileMonitorPrivateStruct() *FileMonitorPrivate {
	err := fileMonitorPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := FileMonitorPrivateNewFromNative(fileMonitorPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeFileMonitorPrivate)
	return structGo
}
func finalizeFileMonitorPrivate(obj *FileMonitorPrivate) {
	fileMonitorPrivateStruct.Free(obj.native)
}

var fileOutputStreamClassStruct *gi.Struct
var fileOutputStreamClassStruct_Once sync.Once

func fileOutputStreamClassStruct_Set() error {
	var err error
	fileOutputStreamClassStruct_Once.Do(func() {
		fileOutputStreamClassStruct, err = gi.StructNew("Gio", "FileOutputStreamClass")
	})
	return err
}

type FileOutputStreamClass struct {
	native unsafe.Pointer
}

func FileOutputStreamClassNewFromNative(native unsafe.Pointer) *FileOutputStreamClass {
	instance := &FileOutputStreamClass{native: native}

	return instance
}

/*
CastToFileOutputStreamClass down casts any arbitrary Object to FileOutputStreamClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a FileOutputStreamClass.
*/
func (recv *FileOutputStreamClass) CastToFileOutputStreamClass(object *gobject.Object) *FileOutputStreamClass {
	return FileOutputStreamClassNewFromNative(object.Native())
}

func (recv *FileOutputStreamClass) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentClass returns the C field 'parent_class'.
func (recv *FileOutputStreamClass) FieldParentClass() *OutputStreamClass {
	argValue := gi.StructFieldGet(fileOutputStreamClassStruct, recv.native, "parent_class")
	value := OutputStreamClassNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *FileOutputStreamClass) SetFieldParentClass(value *OutputStreamClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.StructFieldSet(fileOutputStreamClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value 'tell' : for field getter : missing Type

// UNSUPPORTED : C value 'tell' : for field setter : missing Type

// UNSUPPORTED : C value 'can_seek' : for field getter : missing Type

// UNSUPPORTED : C value 'can_seek' : for field setter : missing Type

// UNSUPPORTED : C value 'seek' : for field getter : missing Type

// UNSUPPORTED : C value 'seek' : for field setter : missing Type

// UNSUPPORTED : C value 'can_truncate' : for field getter : missing Type

// UNSUPPORTED : C value 'can_truncate' : for field setter : missing Type

// UNSUPPORTED : C value 'truncate_fn' : for field getter : missing Type

// UNSUPPORTED : C value 'truncate_fn' : for field setter : missing Type

// UNSUPPORTED : C value 'query_info' : for field getter : missing Type

// UNSUPPORTED : C value 'query_info' : for field setter : missing Type

// UNSUPPORTED : C value 'query_info_async' : for field getter : missing Type

// UNSUPPORTED : C value 'query_info_async' : for field setter : missing Type

// UNSUPPORTED : C value 'query_info_finish' : for field getter : missing Type

// UNSUPPORTED : C value 'query_info_finish' : for field setter : missing Type

// UNSUPPORTED : C value 'get_etag' : for field getter : missing Type

// UNSUPPORTED : C value 'get_etag' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved4' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved5' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved5' : for field setter : missing Type

// FileOutputStreamClassStruct creates an uninitialised FileOutputStreamClass.
func FileOutputStreamClassStruct() *FileOutputStreamClass {
	err := fileOutputStreamClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := FileOutputStreamClassNewFromNative(fileOutputStreamClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeFileOutputStreamClass)
	return structGo
}
func finalizeFileOutputStreamClass(obj *FileOutputStreamClass) {
	fileOutputStreamClassStruct.Free(obj.native)
}

var fileOutputStreamPrivateStruct *gi.Struct
var fileOutputStreamPrivateStruct_Once sync.Once

func fileOutputStreamPrivateStruct_Set() error {
	var err error
	fileOutputStreamPrivateStruct_Once.Do(func() {
		fileOutputStreamPrivateStruct, err = gi.StructNew("Gio", "FileOutputStreamPrivate")
	})
	return err
}

type FileOutputStreamPrivate struct {
	native unsafe.Pointer
}

func FileOutputStreamPrivateNewFromNative(native unsafe.Pointer) *FileOutputStreamPrivate {
	instance := &FileOutputStreamPrivate{native: native}

	return instance
}

/*
CastToFileOutputStreamPrivate down casts any arbitrary Object to FileOutputStreamPrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a FileOutputStreamPrivate.
*/
func (recv *FileOutputStreamPrivate) CastToFileOutputStreamPrivate(object *gobject.Object) *FileOutputStreamPrivate {
	return FileOutputStreamPrivateNewFromNative(object.Native())
}

func (recv *FileOutputStreamPrivate) Native() unsafe.Pointer {
	return recv.native
}

// FileOutputStreamPrivateStruct creates an uninitialised FileOutputStreamPrivate.
func FileOutputStreamPrivateStruct() *FileOutputStreamPrivate {
	err := fileOutputStreamPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := FileOutputStreamPrivateNewFromNative(fileOutputStreamPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeFileOutputStreamPrivate)
	return structGo
}
func finalizeFileOutputStreamPrivate(obj *FileOutputStreamPrivate) {
	fileOutputStreamPrivateStruct.Free(obj.native)
}

var filenameCompleterClassStruct *gi.Struct
var filenameCompleterClassStruct_Once sync.Once

func filenameCompleterClassStruct_Set() error {
	var err error
	filenameCompleterClassStruct_Once.Do(func() {
		filenameCompleterClassStruct, err = gi.StructNew("Gio", "FilenameCompleterClass")
	})
	return err
}

type FilenameCompleterClass struct {
	native unsafe.Pointer
}

func FilenameCompleterClassNewFromNative(native unsafe.Pointer) *FilenameCompleterClass {
	instance := &FilenameCompleterClass{native: native}

	return instance
}

/*
CastToFilenameCompleterClass down casts any arbitrary Object to FilenameCompleterClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a FilenameCompleterClass.
*/
func (recv *FilenameCompleterClass) CastToFilenameCompleterClass(object *gobject.Object) *FilenameCompleterClass {
	return FilenameCompleterClassNewFromNative(object.Native())
}

func (recv *FilenameCompleterClass) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'got_completion_data' : for field getter : missing Type

// UNSUPPORTED : C value 'got_completion_data' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved3' : for field setter : missing Type

// FilenameCompleterClassStruct creates an uninitialised FilenameCompleterClass.
func FilenameCompleterClassStruct() *FilenameCompleterClass {
	err := filenameCompleterClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := FilenameCompleterClassNewFromNative(filenameCompleterClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeFilenameCompleterClass)
	return structGo
}
func finalizeFilenameCompleterClass(obj *FilenameCompleterClass) {
	filenameCompleterClassStruct.Free(obj.native)
}

var filterInputStreamClassStruct *gi.Struct
var filterInputStreamClassStruct_Once sync.Once

func filterInputStreamClassStruct_Set() error {
	var err error
	filterInputStreamClassStruct_Once.Do(func() {
		filterInputStreamClassStruct, err = gi.StructNew("Gio", "FilterInputStreamClass")
	})
	return err
}

type FilterInputStreamClass struct {
	native unsafe.Pointer
}

func FilterInputStreamClassNewFromNative(native unsafe.Pointer) *FilterInputStreamClass {
	instance := &FilterInputStreamClass{native: native}

	return instance
}

/*
CastToFilterInputStreamClass down casts any arbitrary Object to FilterInputStreamClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a FilterInputStreamClass.
*/
func (recv *FilterInputStreamClass) CastToFilterInputStreamClass(object *gobject.Object) *FilterInputStreamClass {
	return FilterInputStreamClassNewFromNative(object.Native())
}

func (recv *FilterInputStreamClass) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentClass returns the C field 'parent_class'.
func (recv *FilterInputStreamClass) FieldParentClass() *InputStreamClass {
	argValue := gi.StructFieldGet(filterInputStreamClassStruct, recv.native, "parent_class")
	value := InputStreamClassNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *FilterInputStreamClass) SetFieldParentClass(value *InputStreamClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.StructFieldSet(filterInputStreamClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value '_g_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved3' : for field setter : missing Type

// FilterInputStreamClassStruct creates an uninitialised FilterInputStreamClass.
func FilterInputStreamClassStruct() *FilterInputStreamClass {
	err := filterInputStreamClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := FilterInputStreamClassNewFromNative(filterInputStreamClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeFilterInputStreamClass)
	return structGo
}
func finalizeFilterInputStreamClass(obj *FilterInputStreamClass) {
	filterInputStreamClassStruct.Free(obj.native)
}

var filterOutputStreamClassStruct *gi.Struct
var filterOutputStreamClassStruct_Once sync.Once

func filterOutputStreamClassStruct_Set() error {
	var err error
	filterOutputStreamClassStruct_Once.Do(func() {
		filterOutputStreamClassStruct, err = gi.StructNew("Gio", "FilterOutputStreamClass")
	})
	return err
}

type FilterOutputStreamClass struct {
	native unsafe.Pointer
}

func FilterOutputStreamClassNewFromNative(native unsafe.Pointer) *FilterOutputStreamClass {
	instance := &FilterOutputStreamClass{native: native}

	return instance
}

/*
CastToFilterOutputStreamClass down casts any arbitrary Object to FilterOutputStreamClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a FilterOutputStreamClass.
*/
func (recv *FilterOutputStreamClass) CastToFilterOutputStreamClass(object *gobject.Object) *FilterOutputStreamClass {
	return FilterOutputStreamClassNewFromNative(object.Native())
}

func (recv *FilterOutputStreamClass) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentClass returns the C field 'parent_class'.
func (recv *FilterOutputStreamClass) FieldParentClass() *OutputStreamClass {
	argValue := gi.StructFieldGet(filterOutputStreamClassStruct, recv.native, "parent_class")
	value := OutputStreamClassNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *FilterOutputStreamClass) SetFieldParentClass(value *OutputStreamClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.StructFieldSet(filterOutputStreamClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value '_g_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved3' : for field setter : missing Type

// FilterOutputStreamClassStruct creates an uninitialised FilterOutputStreamClass.
func FilterOutputStreamClassStruct() *FilterOutputStreamClass {
	err := filterOutputStreamClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := FilterOutputStreamClassNewFromNative(filterOutputStreamClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeFilterOutputStreamClass)
	return structGo
}
func finalizeFilterOutputStreamClass(obj *FilterOutputStreamClass) {
	filterOutputStreamClassStruct.Free(obj.native)
}

var iOExtensionStruct *gi.Struct
var iOExtensionStruct_Once sync.Once

func iOExtensionStruct_Set() error {
	var err error
	iOExtensionStruct_Once.Do(func() {
		iOExtensionStruct, err = gi.StructNew("Gio", "IOExtension")
	})
	return err
}

type IOExtension struct {
	native unsafe.Pointer
}

func IOExtensionNewFromNative(native unsafe.Pointer) *IOExtension {
	instance := &IOExtension{native: native}

	return instance
}

/*
CastToIOExtension down casts any arbitrary Object to IOExtension.
Exercise care, as this is a potentially dangerous function
if the Object is not a IOExtension.
*/
func (recv *IOExtension) CastToIOExtension(object *gobject.Object) *IOExtension {
	return IOExtensionNewFromNative(object.Native())
}

func (recv *IOExtension) Native() unsafe.Pointer {
	return recv.native
}

var iOExtensionGetNameFunction *gi.Function
var iOExtensionGetNameFunction_Once sync.Once

func iOExtensionGetNameFunction_Set() error {
	var err error
	iOExtensionGetNameFunction_Once.Do(func() {
		err = iOExtensionStruct_Set()
		if err != nil {
			return
		}
		iOExtensionGetNameFunction, err = iOExtensionStruct.InvokerNew("get_name")
	})
	return err
}

// GetName is a representation of the C type g_io_extension_get_name.
func (recv *IOExtension) GetName() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := iOExtensionGetNameFunction_Set()
	if err == nil {
		ret = iOExtensionGetNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var iOExtensionGetPriorityFunction *gi.Function
var iOExtensionGetPriorityFunction_Once sync.Once

func iOExtensionGetPriorityFunction_Set() error {
	var err error
	iOExtensionGetPriorityFunction_Once.Do(func() {
		err = iOExtensionStruct_Set()
		if err != nil {
			return
		}
		iOExtensionGetPriorityFunction, err = iOExtensionStruct.InvokerNew("get_priority")
	})
	return err
}

// GetPriority is a representation of the C type g_io_extension_get_priority.
func (recv *IOExtension) GetPriority() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := iOExtensionGetPriorityFunction_Set()
	if err == nil {
		ret = iOExtensionGetPriorityFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

// UNSUPPORTED : C value 'g_io_extension_get_type' : return type 'GType' not supported

// UNSUPPORTED : C value 'g_io_extension_ref_class' : return type 'GObject.TypeClass' not supported

// IOExtensionStruct creates an uninitialised IOExtension.
func IOExtensionStruct() *IOExtension {
	err := iOExtensionStruct_Set()
	if err != nil {
		return nil
	}

	structGo := IOExtensionNewFromNative(iOExtensionStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeIOExtension)
	return structGo
}
func finalizeIOExtension(obj *IOExtension) {
	iOExtensionStruct.Free(obj.native)
}

var iOExtensionPointStruct *gi.Struct
var iOExtensionPointStruct_Once sync.Once

func iOExtensionPointStruct_Set() error {
	var err error
	iOExtensionPointStruct_Once.Do(func() {
		iOExtensionPointStruct, err = gi.StructNew("Gio", "IOExtensionPoint")
	})
	return err
}

type IOExtensionPoint struct {
	native unsafe.Pointer
}

func IOExtensionPointNewFromNative(native unsafe.Pointer) *IOExtensionPoint {
	instance := &IOExtensionPoint{native: native}

	return instance
}

/*
CastToIOExtensionPoint down casts any arbitrary Object to IOExtensionPoint.
Exercise care, as this is a potentially dangerous function
if the Object is not a IOExtensionPoint.
*/
func (recv *IOExtensionPoint) CastToIOExtensionPoint(object *gobject.Object) *IOExtensionPoint {
	return IOExtensionPointNewFromNative(object.Native())
}

func (recv *IOExtensionPoint) Native() unsafe.Pointer {
	return recv.native
}

var iOExtensionPointGetExtensionByNameFunction *gi.Function
var iOExtensionPointGetExtensionByNameFunction_Once sync.Once

func iOExtensionPointGetExtensionByNameFunction_Set() error {
	var err error
	iOExtensionPointGetExtensionByNameFunction_Once.Do(func() {
		err = iOExtensionPointStruct_Set()
		if err != nil {
			return
		}
		iOExtensionPointGetExtensionByNameFunction, err = iOExtensionPointStruct.InvokerNew("get_extension_by_name")
	})
	return err
}

// GetExtensionByName is a representation of the C type g_io_extension_point_get_extension_by_name.
func (recv *IOExtensionPoint) GetExtensionByName(name string) *IOExtension {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(name)

	var ret gi.Argument

	err := iOExtensionPointGetExtensionByNameFunction_Set()
	if err == nil {
		ret = iOExtensionPointGetExtensionByNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := IOExtensionNewFromNative(ret.Pointer())

	return retGo
}

// UNSUPPORTED : C value 'g_io_extension_point_get_extensions' : return type 'GLib.List' not supported

// UNSUPPORTED : C value 'g_io_extension_point_get_required_type' : return type 'GType' not supported

// UNSUPPORTED : C value 'g_io_extension_point_set_required_type' : parameter 'type' of type 'GType' not supported

// IOExtensionPointStruct creates an uninitialised IOExtensionPoint.
func IOExtensionPointStruct() *IOExtensionPoint {
	err := iOExtensionPointStruct_Set()
	if err != nil {
		return nil
	}

	structGo := IOExtensionPointNewFromNative(iOExtensionPointStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeIOExtensionPoint)
	return structGo
}
func finalizeIOExtensionPoint(obj *IOExtensionPoint) {
	iOExtensionPointStruct.Free(obj.native)
}

var iOModuleClassStruct *gi.Struct
var iOModuleClassStruct_Once sync.Once

func iOModuleClassStruct_Set() error {
	var err error
	iOModuleClassStruct_Once.Do(func() {
		iOModuleClassStruct, err = gi.StructNew("Gio", "IOModuleClass")
	})
	return err
}

type IOModuleClass struct {
	native unsafe.Pointer
}

func IOModuleClassNewFromNative(native unsafe.Pointer) *IOModuleClass {
	instance := &IOModuleClass{native: native}

	return instance
}

/*
CastToIOModuleClass down casts any arbitrary Object to IOModuleClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a IOModuleClass.
*/
func (recv *IOModuleClass) CastToIOModuleClass(object *gobject.Object) *IOModuleClass {
	return IOModuleClassNewFromNative(object.Native())
}

func (recv *IOModuleClass) Native() unsafe.Pointer {
	return recv.native
}

// IOModuleClassStruct creates an uninitialised IOModuleClass.
func IOModuleClassStruct() *IOModuleClass {
	err := iOModuleClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := IOModuleClassNewFromNative(iOModuleClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeIOModuleClass)
	return structGo
}
func finalizeIOModuleClass(obj *IOModuleClass) {
	iOModuleClassStruct.Free(obj.native)
}

var iOModuleScopeStruct *gi.Struct
var iOModuleScopeStruct_Once sync.Once

func iOModuleScopeStruct_Set() error {
	var err error
	iOModuleScopeStruct_Once.Do(func() {
		iOModuleScopeStruct, err = gi.StructNew("Gio", "IOModuleScope")
	})
	return err
}

type IOModuleScope struct {
	native unsafe.Pointer
}

func IOModuleScopeNewFromNative(native unsafe.Pointer) *IOModuleScope {
	instance := &IOModuleScope{native: native}

	return instance
}

/*
CastToIOModuleScope down casts any arbitrary Object to IOModuleScope.
Exercise care, as this is a potentially dangerous function
if the Object is not a IOModuleScope.
*/
func (recv *IOModuleScope) CastToIOModuleScope(object *gobject.Object) *IOModuleScope {
	return IOModuleScopeNewFromNative(object.Native())
}

func (recv *IOModuleScope) Native() unsafe.Pointer {
	return recv.native
}

var iOModuleScopeBlockFunction *gi.Function
var iOModuleScopeBlockFunction_Once sync.Once

func iOModuleScopeBlockFunction_Set() error {
	var err error
	iOModuleScopeBlockFunction_Once.Do(func() {
		err = iOModuleScopeStruct_Set()
		if err != nil {
			return
		}
		iOModuleScopeBlockFunction, err = iOModuleScopeStruct.InvokerNew("block")
	})
	return err
}

// Block is a representation of the C type g_io_module_scope_block.
func (recv *IOModuleScope) Block(basename string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(basename)

	err := iOModuleScopeBlockFunction_Set()
	if err == nil {
		iOModuleScopeBlockFunction.Invoke(inArgs[:], nil)
	}

	return
}

var iOModuleScopeFreeFunction *gi.Function
var iOModuleScopeFreeFunction_Once sync.Once

func iOModuleScopeFreeFunction_Set() error {
	var err error
	iOModuleScopeFreeFunction_Once.Do(func() {
		err = iOModuleScopeStruct_Set()
		if err != nil {
			return
		}
		iOModuleScopeFreeFunction, err = iOModuleScopeStruct.InvokerNew("free")
	})
	return err
}

// Free is a representation of the C type g_io_module_scope_free.
func (recv *IOModuleScope) Free() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := iOModuleScopeFreeFunction_Set()
	if err == nil {
		iOModuleScopeFreeFunction.Invoke(inArgs[:], nil)
	}

	return
}

// IOModuleScopeStruct creates an uninitialised IOModuleScope.
func IOModuleScopeStruct() *IOModuleScope {
	err := iOModuleScopeStruct_Set()
	if err != nil {
		return nil
	}

	structGo := IOModuleScopeNewFromNative(iOModuleScopeStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeIOModuleScope)
	return structGo
}
func finalizeIOModuleScope(obj *IOModuleScope) {
	iOModuleScopeStruct.Free(obj.native)
}

var iOSchedulerJobStruct *gi.Struct
var iOSchedulerJobStruct_Once sync.Once

func iOSchedulerJobStruct_Set() error {
	var err error
	iOSchedulerJobStruct_Once.Do(func() {
		iOSchedulerJobStruct, err = gi.StructNew("Gio", "IOSchedulerJob")
	})
	return err
}

type IOSchedulerJob struct {
	native unsafe.Pointer
}

func IOSchedulerJobNewFromNative(native unsafe.Pointer) *IOSchedulerJob {
	instance := &IOSchedulerJob{native: native}

	return instance
}

/*
CastToIOSchedulerJob down casts any arbitrary Object to IOSchedulerJob.
Exercise care, as this is a potentially dangerous function
if the Object is not a IOSchedulerJob.
*/
func (recv *IOSchedulerJob) CastToIOSchedulerJob(object *gobject.Object) *IOSchedulerJob {
	return IOSchedulerJobNewFromNative(object.Native())
}

func (recv *IOSchedulerJob) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'g_io_scheduler_job_send_to_mainloop' : parameter 'func' of type 'GLib.SourceFunc' not supported

// UNSUPPORTED : C value 'g_io_scheduler_job_send_to_mainloop_async' : parameter 'func' of type 'GLib.SourceFunc' not supported

// IOSchedulerJobStruct creates an uninitialised IOSchedulerJob.
func IOSchedulerJobStruct() *IOSchedulerJob {
	err := iOSchedulerJobStruct_Set()
	if err != nil {
		return nil
	}

	structGo := IOSchedulerJobNewFromNative(iOSchedulerJobStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeIOSchedulerJob)
	return structGo
}
func finalizeIOSchedulerJob(obj *IOSchedulerJob) {
	iOSchedulerJobStruct.Free(obj.native)
}

var iOStreamAdapterStruct *gi.Struct
var iOStreamAdapterStruct_Once sync.Once

func iOStreamAdapterStruct_Set() error {
	var err error
	iOStreamAdapterStruct_Once.Do(func() {
		iOStreamAdapterStruct, err = gi.StructNew("Gio", "IOStreamAdapter")
	})
	return err
}

type IOStreamAdapter struct {
	native unsafe.Pointer
}

func IOStreamAdapterNewFromNative(native unsafe.Pointer) *IOStreamAdapter {
	instance := &IOStreamAdapter{native: native}

	return instance
}

/*
CastToIOStreamAdapter down casts any arbitrary Object to IOStreamAdapter.
Exercise care, as this is a potentially dangerous function
if the Object is not a IOStreamAdapter.
*/
func (recv *IOStreamAdapter) CastToIOStreamAdapter(object *gobject.Object) *IOStreamAdapter {
	return IOStreamAdapterNewFromNative(object.Native())
}

func (recv *IOStreamAdapter) Native() unsafe.Pointer {
	return recv.native
}

// IOStreamAdapterStruct creates an uninitialised IOStreamAdapter.
func IOStreamAdapterStruct() *IOStreamAdapter {
	err := iOStreamAdapterStruct_Set()
	if err != nil {
		return nil
	}

	structGo := IOStreamAdapterNewFromNative(iOStreamAdapterStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeIOStreamAdapter)
	return structGo
}
func finalizeIOStreamAdapter(obj *IOStreamAdapter) {
	iOStreamAdapterStruct.Free(obj.native)
}

var iOStreamClassStruct *gi.Struct
var iOStreamClassStruct_Once sync.Once

func iOStreamClassStruct_Set() error {
	var err error
	iOStreamClassStruct_Once.Do(func() {
		iOStreamClassStruct, err = gi.StructNew("Gio", "IOStreamClass")
	})
	return err
}

type IOStreamClass struct {
	native unsafe.Pointer
}

func IOStreamClassNewFromNative(native unsafe.Pointer) *IOStreamClass {
	instance := &IOStreamClass{native: native}

	return instance
}

/*
CastToIOStreamClass down casts any arbitrary Object to IOStreamClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a IOStreamClass.
*/
func (recv *IOStreamClass) CastToIOStreamClass(object *gobject.Object) *IOStreamClass {
	return IOStreamClassNewFromNative(object.Native())
}

func (recv *IOStreamClass) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'get_input_stream' : for field getter : missing Type

// UNSUPPORTED : C value 'get_input_stream' : for field setter : missing Type

// UNSUPPORTED : C value 'get_output_stream' : for field getter : missing Type

// UNSUPPORTED : C value 'get_output_stream' : for field setter : missing Type

// UNSUPPORTED : C value 'close_fn' : for field getter : missing Type

// UNSUPPORTED : C value 'close_fn' : for field setter : missing Type

// UNSUPPORTED : C value 'close_async' : for field getter : missing Type

// UNSUPPORTED : C value 'close_async' : for field setter : missing Type

// UNSUPPORTED : C value 'close_finish' : for field getter : missing Type

// UNSUPPORTED : C value 'close_finish' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved4' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved5' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved5' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved6' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved6' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved7' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved7' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved8' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved8' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved9' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved9' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved10' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved10' : for field setter : missing Type

// IOStreamClassStruct creates an uninitialised IOStreamClass.
func IOStreamClassStruct() *IOStreamClass {
	err := iOStreamClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := IOStreamClassNewFromNative(iOStreamClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeIOStreamClass)
	return structGo
}
func finalizeIOStreamClass(obj *IOStreamClass) {
	iOStreamClassStruct.Free(obj.native)
}

var iOStreamPrivateStruct *gi.Struct
var iOStreamPrivateStruct_Once sync.Once

func iOStreamPrivateStruct_Set() error {
	var err error
	iOStreamPrivateStruct_Once.Do(func() {
		iOStreamPrivateStruct, err = gi.StructNew("Gio", "IOStreamPrivate")
	})
	return err
}

type IOStreamPrivate struct {
	native unsafe.Pointer
}

func IOStreamPrivateNewFromNative(native unsafe.Pointer) *IOStreamPrivate {
	instance := &IOStreamPrivate{native: native}

	return instance
}

/*
CastToIOStreamPrivate down casts any arbitrary Object to IOStreamPrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a IOStreamPrivate.
*/
func (recv *IOStreamPrivate) CastToIOStreamPrivate(object *gobject.Object) *IOStreamPrivate {
	return IOStreamPrivateNewFromNative(object.Native())
}

func (recv *IOStreamPrivate) Native() unsafe.Pointer {
	return recv.native
}

// IOStreamPrivateStruct creates an uninitialised IOStreamPrivate.
func IOStreamPrivateStruct() *IOStreamPrivate {
	err := iOStreamPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := IOStreamPrivateNewFromNative(iOStreamPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeIOStreamPrivate)
	return structGo
}
func finalizeIOStreamPrivate(obj *IOStreamPrivate) {
	iOStreamPrivateStruct.Free(obj.native)
}

var iconIfaceStruct *gi.Struct
var iconIfaceStruct_Once sync.Once

func iconIfaceStruct_Set() error {
	var err error
	iconIfaceStruct_Once.Do(func() {
		iconIfaceStruct, err = gi.StructNew("Gio", "IconIface")
	})
	return err
}

type IconIface struct {
	native unsafe.Pointer
}

func IconIfaceNewFromNative(native unsafe.Pointer) *IconIface {
	instance := &IconIface{native: native}

	return instance
}

/*
CastToIconIface down casts any arbitrary Object to IconIface.
Exercise care, as this is a potentially dangerous function
if the Object is not a IconIface.
*/
func (recv *IconIface) CastToIconIface(object *gobject.Object) *IconIface {
	return IconIfaceNewFromNative(object.Native())
}

func (recv *IconIface) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'g_iface' : for field getter : no Go type for 'GObject.TypeInterface'

// UNSUPPORTED : C value 'g_iface' : for field setter : no Go type for 'GObject.TypeInterface'

// UNSUPPORTED : C value 'hash' : for field getter : missing Type

// UNSUPPORTED : C value 'hash' : for field setter : missing Type

// UNSUPPORTED : C value 'equal' : for field getter : missing Type

// UNSUPPORTED : C value 'equal' : for field setter : missing Type

// UNSUPPORTED : C value 'to_tokens' : for field getter : missing Type

// UNSUPPORTED : C value 'to_tokens' : for field setter : missing Type

// UNSUPPORTED : C value 'from_tokens' : for field getter : missing Type

// UNSUPPORTED : C value 'from_tokens' : for field setter : missing Type

// UNSUPPORTED : C value 'serialize' : for field getter : missing Type

// UNSUPPORTED : C value 'serialize' : for field setter : missing Type

// IconIfaceStruct creates an uninitialised IconIface.
func IconIfaceStruct() *IconIface {
	err := iconIfaceStruct_Set()
	if err != nil {
		return nil
	}

	structGo := IconIfaceNewFromNative(iconIfaceStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeIconIface)
	return structGo
}
func finalizeIconIface(obj *IconIface) {
	iconIfaceStruct.Free(obj.native)
}

var inetAddressClassStruct *gi.Struct
var inetAddressClassStruct_Once sync.Once

func inetAddressClassStruct_Set() error {
	var err error
	inetAddressClassStruct_Once.Do(func() {
		inetAddressClassStruct, err = gi.StructNew("Gio", "InetAddressClass")
	})
	return err
}

type InetAddressClass struct {
	native unsafe.Pointer
}

func InetAddressClassNewFromNative(native unsafe.Pointer) *InetAddressClass {
	instance := &InetAddressClass{native: native}

	return instance
}

/*
CastToInetAddressClass down casts any arbitrary Object to InetAddressClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a InetAddressClass.
*/
func (recv *InetAddressClass) CastToInetAddressClass(object *gobject.Object) *InetAddressClass {
	return InetAddressClassNewFromNative(object.Native())
}

func (recv *InetAddressClass) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'to_string' : for field getter : missing Type

// UNSUPPORTED : C value 'to_string' : for field setter : missing Type

// UNSUPPORTED : C value 'to_bytes' : for field getter : missing Type

// UNSUPPORTED : C value 'to_bytes' : for field setter : missing Type

// InetAddressClassStruct creates an uninitialised InetAddressClass.
func InetAddressClassStruct() *InetAddressClass {
	err := inetAddressClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := InetAddressClassNewFromNative(inetAddressClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeInetAddressClass)
	return structGo
}
func finalizeInetAddressClass(obj *InetAddressClass) {
	inetAddressClassStruct.Free(obj.native)
}

var inetAddressMaskClassStruct *gi.Struct
var inetAddressMaskClassStruct_Once sync.Once

func inetAddressMaskClassStruct_Set() error {
	var err error
	inetAddressMaskClassStruct_Once.Do(func() {
		inetAddressMaskClassStruct, err = gi.StructNew("Gio", "InetAddressMaskClass")
	})
	return err
}

type InetAddressMaskClass struct {
	native unsafe.Pointer
}

func InetAddressMaskClassNewFromNative(native unsafe.Pointer) *InetAddressMaskClass {
	instance := &InetAddressMaskClass{native: native}

	return instance
}

/*
CastToInetAddressMaskClass down casts any arbitrary Object to InetAddressMaskClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a InetAddressMaskClass.
*/
func (recv *InetAddressMaskClass) CastToInetAddressMaskClass(object *gobject.Object) *InetAddressMaskClass {
	return InetAddressMaskClassNewFromNative(object.Native())
}

func (recv *InetAddressMaskClass) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// InetAddressMaskClassStruct creates an uninitialised InetAddressMaskClass.
func InetAddressMaskClassStruct() *InetAddressMaskClass {
	err := inetAddressMaskClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := InetAddressMaskClassNewFromNative(inetAddressMaskClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeInetAddressMaskClass)
	return structGo
}
func finalizeInetAddressMaskClass(obj *InetAddressMaskClass) {
	inetAddressMaskClassStruct.Free(obj.native)
}

var inetAddressMaskPrivateStruct *gi.Struct
var inetAddressMaskPrivateStruct_Once sync.Once

func inetAddressMaskPrivateStruct_Set() error {
	var err error
	inetAddressMaskPrivateStruct_Once.Do(func() {
		inetAddressMaskPrivateStruct, err = gi.StructNew("Gio", "InetAddressMaskPrivate")
	})
	return err
}

type InetAddressMaskPrivate struct {
	native unsafe.Pointer
}

func InetAddressMaskPrivateNewFromNative(native unsafe.Pointer) *InetAddressMaskPrivate {
	instance := &InetAddressMaskPrivate{native: native}

	return instance
}

/*
CastToInetAddressMaskPrivate down casts any arbitrary Object to InetAddressMaskPrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a InetAddressMaskPrivate.
*/
func (recv *InetAddressMaskPrivate) CastToInetAddressMaskPrivate(object *gobject.Object) *InetAddressMaskPrivate {
	return InetAddressMaskPrivateNewFromNative(object.Native())
}

func (recv *InetAddressMaskPrivate) Native() unsafe.Pointer {
	return recv.native
}

// InetAddressMaskPrivateStruct creates an uninitialised InetAddressMaskPrivate.
func InetAddressMaskPrivateStruct() *InetAddressMaskPrivate {
	err := inetAddressMaskPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := InetAddressMaskPrivateNewFromNative(inetAddressMaskPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeInetAddressMaskPrivate)
	return structGo
}
func finalizeInetAddressMaskPrivate(obj *InetAddressMaskPrivate) {
	inetAddressMaskPrivateStruct.Free(obj.native)
}

var inetAddressPrivateStruct *gi.Struct
var inetAddressPrivateStruct_Once sync.Once

func inetAddressPrivateStruct_Set() error {
	var err error
	inetAddressPrivateStruct_Once.Do(func() {
		inetAddressPrivateStruct, err = gi.StructNew("Gio", "InetAddressPrivate")
	})
	return err
}

type InetAddressPrivate struct {
	native unsafe.Pointer
}

func InetAddressPrivateNewFromNative(native unsafe.Pointer) *InetAddressPrivate {
	instance := &InetAddressPrivate{native: native}

	return instance
}

/*
CastToInetAddressPrivate down casts any arbitrary Object to InetAddressPrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a InetAddressPrivate.
*/
func (recv *InetAddressPrivate) CastToInetAddressPrivate(object *gobject.Object) *InetAddressPrivate {
	return InetAddressPrivateNewFromNative(object.Native())
}

func (recv *InetAddressPrivate) Native() unsafe.Pointer {
	return recv.native
}

// InetAddressPrivateStruct creates an uninitialised InetAddressPrivate.
func InetAddressPrivateStruct() *InetAddressPrivate {
	err := inetAddressPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := InetAddressPrivateNewFromNative(inetAddressPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeInetAddressPrivate)
	return structGo
}
func finalizeInetAddressPrivate(obj *InetAddressPrivate) {
	inetAddressPrivateStruct.Free(obj.native)
}

var inetSocketAddressClassStruct *gi.Struct
var inetSocketAddressClassStruct_Once sync.Once

func inetSocketAddressClassStruct_Set() error {
	var err error
	inetSocketAddressClassStruct_Once.Do(func() {
		inetSocketAddressClassStruct, err = gi.StructNew("Gio", "InetSocketAddressClass")
	})
	return err
}

type InetSocketAddressClass struct {
	native unsafe.Pointer
}

func InetSocketAddressClassNewFromNative(native unsafe.Pointer) *InetSocketAddressClass {
	instance := &InetSocketAddressClass{native: native}

	return instance
}

/*
CastToInetSocketAddressClass down casts any arbitrary Object to InetSocketAddressClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a InetSocketAddressClass.
*/
func (recv *InetSocketAddressClass) CastToInetSocketAddressClass(object *gobject.Object) *InetSocketAddressClass {
	return InetSocketAddressClassNewFromNative(object.Native())
}

func (recv *InetSocketAddressClass) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentClass returns the C field 'parent_class'.
func (recv *InetSocketAddressClass) FieldParentClass() *SocketAddressClass {
	argValue := gi.StructFieldGet(inetSocketAddressClassStruct, recv.native, "parent_class")
	value := SocketAddressClassNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *InetSocketAddressClass) SetFieldParentClass(value *SocketAddressClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.StructFieldSet(inetSocketAddressClassStruct, recv.native, "parent_class", argValue)
}

// InetSocketAddressClassStruct creates an uninitialised InetSocketAddressClass.
func InetSocketAddressClassStruct() *InetSocketAddressClass {
	err := inetSocketAddressClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := InetSocketAddressClassNewFromNative(inetSocketAddressClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeInetSocketAddressClass)
	return structGo
}
func finalizeInetSocketAddressClass(obj *InetSocketAddressClass) {
	inetSocketAddressClassStruct.Free(obj.native)
}

var inetSocketAddressPrivateStruct *gi.Struct
var inetSocketAddressPrivateStruct_Once sync.Once

func inetSocketAddressPrivateStruct_Set() error {
	var err error
	inetSocketAddressPrivateStruct_Once.Do(func() {
		inetSocketAddressPrivateStruct, err = gi.StructNew("Gio", "InetSocketAddressPrivate")
	})
	return err
}

type InetSocketAddressPrivate struct {
	native unsafe.Pointer
}

func InetSocketAddressPrivateNewFromNative(native unsafe.Pointer) *InetSocketAddressPrivate {
	instance := &InetSocketAddressPrivate{native: native}

	return instance
}

/*
CastToInetSocketAddressPrivate down casts any arbitrary Object to InetSocketAddressPrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a InetSocketAddressPrivate.
*/
func (recv *InetSocketAddressPrivate) CastToInetSocketAddressPrivate(object *gobject.Object) *InetSocketAddressPrivate {
	return InetSocketAddressPrivateNewFromNative(object.Native())
}

func (recv *InetSocketAddressPrivate) Native() unsafe.Pointer {
	return recv.native
}

// InetSocketAddressPrivateStruct creates an uninitialised InetSocketAddressPrivate.
func InetSocketAddressPrivateStruct() *InetSocketAddressPrivate {
	err := inetSocketAddressPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := InetSocketAddressPrivateNewFromNative(inetSocketAddressPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeInetSocketAddressPrivate)
	return structGo
}
func finalizeInetSocketAddressPrivate(obj *InetSocketAddressPrivate) {
	inetSocketAddressPrivateStruct.Free(obj.native)
}

var initableIfaceStruct *gi.Struct
var initableIfaceStruct_Once sync.Once

func initableIfaceStruct_Set() error {
	var err error
	initableIfaceStruct_Once.Do(func() {
		initableIfaceStruct, err = gi.StructNew("Gio", "InitableIface")
	})
	return err
}

type InitableIface struct {
	native unsafe.Pointer
}

func InitableIfaceNewFromNative(native unsafe.Pointer) *InitableIface {
	instance := &InitableIface{native: native}

	return instance
}

/*
CastToInitableIface down casts any arbitrary Object to InitableIface.
Exercise care, as this is a potentially dangerous function
if the Object is not a InitableIface.
*/
func (recv *InitableIface) CastToInitableIface(object *gobject.Object) *InitableIface {
	return InitableIfaceNewFromNative(object.Native())
}

func (recv *InitableIface) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'g_iface' : for field getter : no Go type for 'GObject.TypeInterface'

// UNSUPPORTED : C value 'g_iface' : for field setter : no Go type for 'GObject.TypeInterface'

// UNSUPPORTED : C value 'init' : for field getter : missing Type

// UNSUPPORTED : C value 'init' : for field setter : missing Type

// InitableIfaceStruct creates an uninitialised InitableIface.
func InitableIfaceStruct() *InitableIface {
	err := initableIfaceStruct_Set()
	if err != nil {
		return nil
	}

	structGo := InitableIfaceNewFromNative(initableIfaceStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeInitableIface)
	return structGo
}
func finalizeInitableIface(obj *InitableIface) {
	initableIfaceStruct.Free(obj.native)
}

var inputMessageStruct *gi.Struct
var inputMessageStruct_Once sync.Once

func inputMessageStruct_Set() error {
	var err error
	inputMessageStruct_Once.Do(func() {
		inputMessageStruct, err = gi.StructNew("Gio", "InputMessage")
	})
	return err
}

type InputMessage struct {
	native unsafe.Pointer
}

func InputMessageNewFromNative(native unsafe.Pointer) *InputMessage {
	instance := &InputMessage{native: native}

	return instance
}

/*
CastToInputMessage down casts any arbitrary Object to InputMessage.
Exercise care, as this is a potentially dangerous function
if the Object is not a InputMessage.
*/
func (recv *InputMessage) CastToInputMessage(object *gobject.Object) *InputMessage {
	return InputMessageNewFromNative(object.Native())
}

func (recv *InputMessage) Native() unsafe.Pointer {
	return recv.native
}

// FieldAddress returns the C field 'address'.
func (recv *InputMessage) FieldAddress() *SocketAddress {
	argValue := gi.StructFieldGet(inputMessageStruct, recv.native, "address")
	value := SocketAddressNewFromNative(argValue.Pointer())
	return value
}

// SetFieldAddress sets the value of the C field 'address'.
func (recv *InputMessage) SetFieldAddress(value *SocketAddress) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.StructFieldSet(inputMessageStruct, recv.native, "address", argValue)
}

// UNSUPPORTED : C value 'vectors' : for field getter : missing Type

// UNSUPPORTED : C value 'vectors' : for field setter : missing Type

// FieldNumVectors returns the C field 'num_vectors'.
func (recv *InputMessage) FieldNumVectors() uint32 {
	argValue := gi.StructFieldGet(inputMessageStruct, recv.native, "num_vectors")
	value := argValue.Uint32()
	return value
}

// SetFieldNumVectors sets the value of the C field 'num_vectors'.
func (recv *InputMessage) SetFieldNumVectors(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.StructFieldSet(inputMessageStruct, recv.native, "num_vectors", argValue)
}

// FieldBytesReceived returns the C field 'bytes_received'.
func (recv *InputMessage) FieldBytesReceived() uint64 {
	argValue := gi.StructFieldGet(inputMessageStruct, recv.native, "bytes_received")
	value := argValue.Uint64()
	return value
}

// SetFieldBytesReceived sets the value of the C field 'bytes_received'.
func (recv *InputMessage) SetFieldBytesReceived(value uint64) {
	var argValue gi.Argument
	argValue.SetUint64(value)
	gi.StructFieldSet(inputMessageStruct, recv.native, "bytes_received", argValue)
}

// FieldFlags returns the C field 'flags'.
func (recv *InputMessage) FieldFlags() int32 {
	argValue := gi.StructFieldGet(inputMessageStruct, recv.native, "flags")
	value := argValue.Int32()
	return value
}

// SetFieldFlags sets the value of the C field 'flags'.
func (recv *InputMessage) SetFieldFlags(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.StructFieldSet(inputMessageStruct, recv.native, "flags", argValue)
}

// UNSUPPORTED : C value 'control_messages' : for field getter : missing Type

// UNSUPPORTED : C value 'control_messages' : for field setter : missing Type

// FieldNumControlMessages returns the C field 'num_control_messages'.
func (recv *InputMessage) FieldNumControlMessages() uint32 {
	argValue := gi.StructFieldGet(inputMessageStruct, recv.native, "num_control_messages")
	value := argValue.Uint32()
	return value
}

// SetFieldNumControlMessages sets the value of the C field 'num_control_messages'.
func (recv *InputMessage) SetFieldNumControlMessages(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.StructFieldSet(inputMessageStruct, recv.native, "num_control_messages", argValue)
}

// InputMessageStruct creates an uninitialised InputMessage.
func InputMessageStruct() *InputMessage {
	err := inputMessageStruct_Set()
	if err != nil {
		return nil
	}

	structGo := InputMessageNewFromNative(inputMessageStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeInputMessage)
	return structGo
}
func finalizeInputMessage(obj *InputMessage) {
	inputMessageStruct.Free(obj.native)
}

var inputStreamClassStruct *gi.Struct
var inputStreamClassStruct_Once sync.Once

func inputStreamClassStruct_Set() error {
	var err error
	inputStreamClassStruct_Once.Do(func() {
		inputStreamClassStruct, err = gi.StructNew("Gio", "InputStreamClass")
	})
	return err
}

type InputStreamClass struct {
	native unsafe.Pointer
}

func InputStreamClassNewFromNative(native unsafe.Pointer) *InputStreamClass {
	instance := &InputStreamClass{native: native}

	return instance
}

/*
CastToInputStreamClass down casts any arbitrary Object to InputStreamClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a InputStreamClass.
*/
func (recv *InputStreamClass) CastToInputStreamClass(object *gobject.Object) *InputStreamClass {
	return InputStreamClassNewFromNative(object.Native())
}

func (recv *InputStreamClass) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'read_fn' : for field getter : missing Type

// UNSUPPORTED : C value 'read_fn' : for field setter : missing Type

// UNSUPPORTED : C value 'skip' : for field getter : missing Type

// UNSUPPORTED : C value 'skip' : for field setter : missing Type

// UNSUPPORTED : C value 'close_fn' : for field getter : missing Type

// UNSUPPORTED : C value 'close_fn' : for field setter : missing Type

// UNSUPPORTED : C value 'read_async' : for field getter : missing Type

// UNSUPPORTED : C value 'read_async' : for field setter : missing Type

// UNSUPPORTED : C value 'read_finish' : for field getter : missing Type

// UNSUPPORTED : C value 'read_finish' : for field setter : missing Type

// UNSUPPORTED : C value 'skip_async' : for field getter : missing Type

// UNSUPPORTED : C value 'skip_async' : for field setter : missing Type

// UNSUPPORTED : C value 'skip_finish' : for field getter : missing Type

// UNSUPPORTED : C value 'skip_finish' : for field setter : missing Type

// UNSUPPORTED : C value 'close_async' : for field getter : missing Type

// UNSUPPORTED : C value 'close_async' : for field setter : missing Type

// UNSUPPORTED : C value 'close_finish' : for field getter : missing Type

// UNSUPPORTED : C value 'close_finish' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved4' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved5' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved5' : for field setter : missing Type

// InputStreamClassStruct creates an uninitialised InputStreamClass.
func InputStreamClassStruct() *InputStreamClass {
	err := inputStreamClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := InputStreamClassNewFromNative(inputStreamClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeInputStreamClass)
	return structGo
}
func finalizeInputStreamClass(obj *InputStreamClass) {
	inputStreamClassStruct.Free(obj.native)
}

var inputStreamPrivateStruct *gi.Struct
var inputStreamPrivateStruct_Once sync.Once

func inputStreamPrivateStruct_Set() error {
	var err error
	inputStreamPrivateStruct_Once.Do(func() {
		inputStreamPrivateStruct, err = gi.StructNew("Gio", "InputStreamPrivate")
	})
	return err
}

type InputStreamPrivate struct {
	native unsafe.Pointer
}

func InputStreamPrivateNewFromNative(native unsafe.Pointer) *InputStreamPrivate {
	instance := &InputStreamPrivate{native: native}

	return instance
}

/*
CastToInputStreamPrivate down casts any arbitrary Object to InputStreamPrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a InputStreamPrivate.
*/
func (recv *InputStreamPrivate) CastToInputStreamPrivate(object *gobject.Object) *InputStreamPrivate {
	return InputStreamPrivateNewFromNative(object.Native())
}

func (recv *InputStreamPrivate) Native() unsafe.Pointer {
	return recv.native
}

// InputStreamPrivateStruct creates an uninitialised InputStreamPrivate.
func InputStreamPrivateStruct() *InputStreamPrivate {
	err := inputStreamPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := InputStreamPrivateNewFromNative(inputStreamPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeInputStreamPrivate)
	return structGo
}
func finalizeInputStreamPrivate(obj *InputStreamPrivate) {
	inputStreamPrivateStruct.Free(obj.native)
}

var inputVectorStruct *gi.Struct
var inputVectorStruct_Once sync.Once

func inputVectorStruct_Set() error {
	var err error
	inputVectorStruct_Once.Do(func() {
		inputVectorStruct, err = gi.StructNew("Gio", "InputVector")
	})
	return err
}

type InputVector struct {
	native unsafe.Pointer
}

func InputVectorNewFromNative(native unsafe.Pointer) *InputVector {
	instance := &InputVector{native: native}

	return instance
}

/*
CastToInputVector down casts any arbitrary Object to InputVector.
Exercise care, as this is a potentially dangerous function
if the Object is not a InputVector.
*/
func (recv *InputVector) CastToInputVector(object *gobject.Object) *InputVector {
	return InputVectorNewFromNative(object.Native())
}

func (recv *InputVector) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'buffer' : for field getter : no Go type for 'gpointer'

// UNSUPPORTED : C value 'buffer' : for field setter : no Go type for 'gpointer'

// FieldSize returns the C field 'size'.
func (recv *InputVector) FieldSize() uint64 {
	argValue := gi.StructFieldGet(inputVectorStruct, recv.native, "size")
	value := argValue.Uint64()
	return value
}

// SetFieldSize sets the value of the C field 'size'.
func (recv *InputVector) SetFieldSize(value uint64) {
	var argValue gi.Argument
	argValue.SetUint64(value)
	gi.StructFieldSet(inputVectorStruct, recv.native, "size", argValue)
}

// InputVectorStruct creates an uninitialised InputVector.
func InputVectorStruct() *InputVector {
	err := inputVectorStruct_Set()
	if err != nil {
		return nil
	}

	structGo := InputVectorNewFromNative(inputVectorStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeInputVector)
	return structGo
}
func finalizeInputVector(obj *InputVector) {
	inputVectorStruct.Free(obj.native)
}

var listModelInterfaceStruct *gi.Struct
var listModelInterfaceStruct_Once sync.Once

func listModelInterfaceStruct_Set() error {
	var err error
	listModelInterfaceStruct_Once.Do(func() {
		listModelInterfaceStruct, err = gi.StructNew("Gio", "ListModelInterface")
	})
	return err
}

type ListModelInterface struct {
	native unsafe.Pointer
}

func ListModelInterfaceNewFromNative(native unsafe.Pointer) *ListModelInterface {
	instance := &ListModelInterface{native: native}

	return instance
}

/*
CastToListModelInterface down casts any arbitrary Object to ListModelInterface.
Exercise care, as this is a potentially dangerous function
if the Object is not a ListModelInterface.
*/
func (recv *ListModelInterface) CastToListModelInterface(object *gobject.Object) *ListModelInterface {
	return ListModelInterfaceNewFromNative(object.Native())
}

func (recv *ListModelInterface) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'g_iface' : for field getter : no Go type for 'GObject.TypeInterface'

// UNSUPPORTED : C value 'g_iface' : for field setter : no Go type for 'GObject.TypeInterface'

// UNSUPPORTED : C value 'get_item_type' : for field getter : missing Type

// UNSUPPORTED : C value 'get_item_type' : for field setter : missing Type

// UNSUPPORTED : C value 'get_n_items' : for field getter : missing Type

// UNSUPPORTED : C value 'get_n_items' : for field setter : missing Type

// UNSUPPORTED : C value 'get_item' : for field getter : missing Type

// UNSUPPORTED : C value 'get_item' : for field setter : missing Type

// ListModelInterfaceStruct creates an uninitialised ListModelInterface.
func ListModelInterfaceStruct() *ListModelInterface {
	err := listModelInterfaceStruct_Set()
	if err != nil {
		return nil
	}

	structGo := ListModelInterfaceNewFromNative(listModelInterfaceStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeListModelInterface)
	return structGo
}
func finalizeListModelInterface(obj *ListModelInterface) {
	listModelInterfaceStruct.Free(obj.native)
}

var listStoreClassStruct *gi.Struct
var listStoreClassStruct_Once sync.Once

func listStoreClassStruct_Set() error {
	var err error
	listStoreClassStruct_Once.Do(func() {
		listStoreClassStruct, err = gi.StructNew("Gio", "ListStoreClass")
	})
	return err
}

type ListStoreClass struct {
	native unsafe.Pointer
}

func ListStoreClassNewFromNative(native unsafe.Pointer) *ListStoreClass {
	instance := &ListStoreClass{native: native}

	return instance
}

/*
CastToListStoreClass down casts any arbitrary Object to ListStoreClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a ListStoreClass.
*/
func (recv *ListStoreClass) CastToListStoreClass(object *gobject.Object) *ListStoreClass {
	return ListStoreClassNewFromNative(object.Native())
}

func (recv *ListStoreClass) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// ListStoreClassStruct creates an uninitialised ListStoreClass.
func ListStoreClassStruct() *ListStoreClass {
	err := listStoreClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := ListStoreClassNewFromNative(listStoreClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeListStoreClass)
	return structGo
}
func finalizeListStoreClass(obj *ListStoreClass) {
	listStoreClassStruct.Free(obj.native)
}

var loadableIconIfaceStruct *gi.Struct
var loadableIconIfaceStruct_Once sync.Once

func loadableIconIfaceStruct_Set() error {
	var err error
	loadableIconIfaceStruct_Once.Do(func() {
		loadableIconIfaceStruct, err = gi.StructNew("Gio", "LoadableIconIface")
	})
	return err
}

type LoadableIconIface struct {
	native unsafe.Pointer
}

func LoadableIconIfaceNewFromNative(native unsafe.Pointer) *LoadableIconIface {
	instance := &LoadableIconIface{native: native}

	return instance
}

/*
CastToLoadableIconIface down casts any arbitrary Object to LoadableIconIface.
Exercise care, as this is a potentially dangerous function
if the Object is not a LoadableIconIface.
*/
func (recv *LoadableIconIface) CastToLoadableIconIface(object *gobject.Object) *LoadableIconIface {
	return LoadableIconIfaceNewFromNative(object.Native())
}

func (recv *LoadableIconIface) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'g_iface' : for field getter : no Go type for 'GObject.TypeInterface'

// UNSUPPORTED : C value 'g_iface' : for field setter : no Go type for 'GObject.TypeInterface'

// UNSUPPORTED : C value 'load' : for field getter : missing Type

// UNSUPPORTED : C value 'load' : for field setter : missing Type

// UNSUPPORTED : C value 'load_async' : for field getter : missing Type

// UNSUPPORTED : C value 'load_async' : for field setter : missing Type

// UNSUPPORTED : C value 'load_finish' : for field getter : missing Type

// UNSUPPORTED : C value 'load_finish' : for field setter : missing Type

// LoadableIconIfaceStruct creates an uninitialised LoadableIconIface.
func LoadableIconIfaceStruct() *LoadableIconIface {
	err := loadableIconIfaceStruct_Set()
	if err != nil {
		return nil
	}

	structGo := LoadableIconIfaceNewFromNative(loadableIconIfaceStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeLoadableIconIface)
	return structGo
}
func finalizeLoadableIconIface(obj *LoadableIconIface) {
	loadableIconIfaceStruct.Free(obj.native)
}

var memoryInputStreamClassStruct *gi.Struct
var memoryInputStreamClassStruct_Once sync.Once

func memoryInputStreamClassStruct_Set() error {
	var err error
	memoryInputStreamClassStruct_Once.Do(func() {
		memoryInputStreamClassStruct, err = gi.StructNew("Gio", "MemoryInputStreamClass")
	})
	return err
}

type MemoryInputStreamClass struct {
	native unsafe.Pointer
}

func MemoryInputStreamClassNewFromNative(native unsafe.Pointer) *MemoryInputStreamClass {
	instance := &MemoryInputStreamClass{native: native}

	return instance
}

/*
CastToMemoryInputStreamClass down casts any arbitrary Object to MemoryInputStreamClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a MemoryInputStreamClass.
*/
func (recv *MemoryInputStreamClass) CastToMemoryInputStreamClass(object *gobject.Object) *MemoryInputStreamClass {
	return MemoryInputStreamClassNewFromNative(object.Native())
}

func (recv *MemoryInputStreamClass) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentClass returns the C field 'parent_class'.
func (recv *MemoryInputStreamClass) FieldParentClass() *InputStreamClass {
	argValue := gi.StructFieldGet(memoryInputStreamClassStruct, recv.native, "parent_class")
	value := InputStreamClassNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *MemoryInputStreamClass) SetFieldParentClass(value *InputStreamClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.StructFieldSet(memoryInputStreamClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value '_g_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved4' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved5' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved5' : for field setter : missing Type

// MemoryInputStreamClassStruct creates an uninitialised MemoryInputStreamClass.
func MemoryInputStreamClassStruct() *MemoryInputStreamClass {
	err := memoryInputStreamClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := MemoryInputStreamClassNewFromNative(memoryInputStreamClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeMemoryInputStreamClass)
	return structGo
}
func finalizeMemoryInputStreamClass(obj *MemoryInputStreamClass) {
	memoryInputStreamClassStruct.Free(obj.native)
}

var memoryInputStreamPrivateStruct *gi.Struct
var memoryInputStreamPrivateStruct_Once sync.Once

func memoryInputStreamPrivateStruct_Set() error {
	var err error
	memoryInputStreamPrivateStruct_Once.Do(func() {
		memoryInputStreamPrivateStruct, err = gi.StructNew("Gio", "MemoryInputStreamPrivate")
	})
	return err
}

type MemoryInputStreamPrivate struct {
	native unsafe.Pointer
}

func MemoryInputStreamPrivateNewFromNative(native unsafe.Pointer) *MemoryInputStreamPrivate {
	instance := &MemoryInputStreamPrivate{native: native}

	return instance
}

/*
CastToMemoryInputStreamPrivate down casts any arbitrary Object to MemoryInputStreamPrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a MemoryInputStreamPrivate.
*/
func (recv *MemoryInputStreamPrivate) CastToMemoryInputStreamPrivate(object *gobject.Object) *MemoryInputStreamPrivate {
	return MemoryInputStreamPrivateNewFromNative(object.Native())
}

func (recv *MemoryInputStreamPrivate) Native() unsafe.Pointer {
	return recv.native
}

// MemoryInputStreamPrivateStruct creates an uninitialised MemoryInputStreamPrivate.
func MemoryInputStreamPrivateStruct() *MemoryInputStreamPrivate {
	err := memoryInputStreamPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := MemoryInputStreamPrivateNewFromNative(memoryInputStreamPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeMemoryInputStreamPrivate)
	return structGo
}
func finalizeMemoryInputStreamPrivate(obj *MemoryInputStreamPrivate) {
	memoryInputStreamPrivateStruct.Free(obj.native)
}

var memoryOutputStreamClassStruct *gi.Struct
var memoryOutputStreamClassStruct_Once sync.Once

func memoryOutputStreamClassStruct_Set() error {
	var err error
	memoryOutputStreamClassStruct_Once.Do(func() {
		memoryOutputStreamClassStruct, err = gi.StructNew("Gio", "MemoryOutputStreamClass")
	})
	return err
}

type MemoryOutputStreamClass struct {
	native unsafe.Pointer
}

func MemoryOutputStreamClassNewFromNative(native unsafe.Pointer) *MemoryOutputStreamClass {
	instance := &MemoryOutputStreamClass{native: native}

	return instance
}

/*
CastToMemoryOutputStreamClass down casts any arbitrary Object to MemoryOutputStreamClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a MemoryOutputStreamClass.
*/
func (recv *MemoryOutputStreamClass) CastToMemoryOutputStreamClass(object *gobject.Object) *MemoryOutputStreamClass {
	return MemoryOutputStreamClassNewFromNative(object.Native())
}

func (recv *MemoryOutputStreamClass) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentClass returns the C field 'parent_class'.
func (recv *MemoryOutputStreamClass) FieldParentClass() *OutputStreamClass {
	argValue := gi.StructFieldGet(memoryOutputStreamClassStruct, recv.native, "parent_class")
	value := OutputStreamClassNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *MemoryOutputStreamClass) SetFieldParentClass(value *OutputStreamClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.StructFieldSet(memoryOutputStreamClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value '_g_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved4' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved5' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved5' : for field setter : missing Type

// MemoryOutputStreamClassStruct creates an uninitialised MemoryOutputStreamClass.
func MemoryOutputStreamClassStruct() *MemoryOutputStreamClass {
	err := memoryOutputStreamClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := MemoryOutputStreamClassNewFromNative(memoryOutputStreamClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeMemoryOutputStreamClass)
	return structGo
}
func finalizeMemoryOutputStreamClass(obj *MemoryOutputStreamClass) {
	memoryOutputStreamClassStruct.Free(obj.native)
}

var memoryOutputStreamPrivateStruct *gi.Struct
var memoryOutputStreamPrivateStruct_Once sync.Once

func memoryOutputStreamPrivateStruct_Set() error {
	var err error
	memoryOutputStreamPrivateStruct_Once.Do(func() {
		memoryOutputStreamPrivateStruct, err = gi.StructNew("Gio", "MemoryOutputStreamPrivate")
	})
	return err
}

type MemoryOutputStreamPrivate struct {
	native unsafe.Pointer
}

func MemoryOutputStreamPrivateNewFromNative(native unsafe.Pointer) *MemoryOutputStreamPrivate {
	instance := &MemoryOutputStreamPrivate{native: native}

	return instance
}

/*
CastToMemoryOutputStreamPrivate down casts any arbitrary Object to MemoryOutputStreamPrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a MemoryOutputStreamPrivate.
*/
func (recv *MemoryOutputStreamPrivate) CastToMemoryOutputStreamPrivate(object *gobject.Object) *MemoryOutputStreamPrivate {
	return MemoryOutputStreamPrivateNewFromNative(object.Native())
}

func (recv *MemoryOutputStreamPrivate) Native() unsafe.Pointer {
	return recv.native
}

// MemoryOutputStreamPrivateStruct creates an uninitialised MemoryOutputStreamPrivate.
func MemoryOutputStreamPrivateStruct() *MemoryOutputStreamPrivate {
	err := memoryOutputStreamPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := MemoryOutputStreamPrivateNewFromNative(memoryOutputStreamPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeMemoryOutputStreamPrivate)
	return structGo
}
func finalizeMemoryOutputStreamPrivate(obj *MemoryOutputStreamPrivate) {
	memoryOutputStreamPrivateStruct.Free(obj.native)
}

var menuAttributeIterClassStruct *gi.Struct
var menuAttributeIterClassStruct_Once sync.Once

func menuAttributeIterClassStruct_Set() error {
	var err error
	menuAttributeIterClassStruct_Once.Do(func() {
		menuAttributeIterClassStruct, err = gi.StructNew("Gio", "MenuAttributeIterClass")
	})
	return err
}

type MenuAttributeIterClass struct {
	native unsafe.Pointer
}

func MenuAttributeIterClassNewFromNative(native unsafe.Pointer) *MenuAttributeIterClass {
	instance := &MenuAttributeIterClass{native: native}

	return instance
}

/*
CastToMenuAttributeIterClass down casts any arbitrary Object to MenuAttributeIterClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a MenuAttributeIterClass.
*/
func (recv *MenuAttributeIterClass) CastToMenuAttributeIterClass(object *gobject.Object) *MenuAttributeIterClass {
	return MenuAttributeIterClassNewFromNative(object.Native())
}

func (recv *MenuAttributeIterClass) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'get_next' : for field getter : missing Type

// UNSUPPORTED : C value 'get_next' : for field setter : missing Type

// MenuAttributeIterClassStruct creates an uninitialised MenuAttributeIterClass.
func MenuAttributeIterClassStruct() *MenuAttributeIterClass {
	err := menuAttributeIterClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := MenuAttributeIterClassNewFromNative(menuAttributeIterClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeMenuAttributeIterClass)
	return structGo
}
func finalizeMenuAttributeIterClass(obj *MenuAttributeIterClass) {
	menuAttributeIterClassStruct.Free(obj.native)
}

var menuAttributeIterPrivateStruct *gi.Struct
var menuAttributeIterPrivateStruct_Once sync.Once

func menuAttributeIterPrivateStruct_Set() error {
	var err error
	menuAttributeIterPrivateStruct_Once.Do(func() {
		menuAttributeIterPrivateStruct, err = gi.StructNew("Gio", "MenuAttributeIterPrivate")
	})
	return err
}

type MenuAttributeIterPrivate struct {
	native unsafe.Pointer
}

func MenuAttributeIterPrivateNewFromNative(native unsafe.Pointer) *MenuAttributeIterPrivate {
	instance := &MenuAttributeIterPrivate{native: native}

	return instance
}

/*
CastToMenuAttributeIterPrivate down casts any arbitrary Object to MenuAttributeIterPrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a MenuAttributeIterPrivate.
*/
func (recv *MenuAttributeIterPrivate) CastToMenuAttributeIterPrivate(object *gobject.Object) *MenuAttributeIterPrivate {
	return MenuAttributeIterPrivateNewFromNative(object.Native())
}

func (recv *MenuAttributeIterPrivate) Native() unsafe.Pointer {
	return recv.native
}

// MenuAttributeIterPrivateStruct creates an uninitialised MenuAttributeIterPrivate.
func MenuAttributeIterPrivateStruct() *MenuAttributeIterPrivate {
	err := menuAttributeIterPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := MenuAttributeIterPrivateNewFromNative(menuAttributeIterPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeMenuAttributeIterPrivate)
	return structGo
}
func finalizeMenuAttributeIterPrivate(obj *MenuAttributeIterPrivate) {
	menuAttributeIterPrivateStruct.Free(obj.native)
}

var menuLinkIterClassStruct *gi.Struct
var menuLinkIterClassStruct_Once sync.Once

func menuLinkIterClassStruct_Set() error {
	var err error
	menuLinkIterClassStruct_Once.Do(func() {
		menuLinkIterClassStruct, err = gi.StructNew("Gio", "MenuLinkIterClass")
	})
	return err
}

type MenuLinkIterClass struct {
	native unsafe.Pointer
}

func MenuLinkIterClassNewFromNative(native unsafe.Pointer) *MenuLinkIterClass {
	instance := &MenuLinkIterClass{native: native}

	return instance
}

/*
CastToMenuLinkIterClass down casts any arbitrary Object to MenuLinkIterClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a MenuLinkIterClass.
*/
func (recv *MenuLinkIterClass) CastToMenuLinkIterClass(object *gobject.Object) *MenuLinkIterClass {
	return MenuLinkIterClassNewFromNative(object.Native())
}

func (recv *MenuLinkIterClass) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'get_next' : for field getter : missing Type

// UNSUPPORTED : C value 'get_next' : for field setter : missing Type

// MenuLinkIterClassStruct creates an uninitialised MenuLinkIterClass.
func MenuLinkIterClassStruct() *MenuLinkIterClass {
	err := menuLinkIterClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := MenuLinkIterClassNewFromNative(menuLinkIterClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeMenuLinkIterClass)
	return structGo
}
func finalizeMenuLinkIterClass(obj *MenuLinkIterClass) {
	menuLinkIterClassStruct.Free(obj.native)
}

var menuLinkIterPrivateStruct *gi.Struct
var menuLinkIterPrivateStruct_Once sync.Once

func menuLinkIterPrivateStruct_Set() error {
	var err error
	menuLinkIterPrivateStruct_Once.Do(func() {
		menuLinkIterPrivateStruct, err = gi.StructNew("Gio", "MenuLinkIterPrivate")
	})
	return err
}

type MenuLinkIterPrivate struct {
	native unsafe.Pointer
}

func MenuLinkIterPrivateNewFromNative(native unsafe.Pointer) *MenuLinkIterPrivate {
	instance := &MenuLinkIterPrivate{native: native}

	return instance
}

/*
CastToMenuLinkIterPrivate down casts any arbitrary Object to MenuLinkIterPrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a MenuLinkIterPrivate.
*/
func (recv *MenuLinkIterPrivate) CastToMenuLinkIterPrivate(object *gobject.Object) *MenuLinkIterPrivate {
	return MenuLinkIterPrivateNewFromNative(object.Native())
}

func (recv *MenuLinkIterPrivate) Native() unsafe.Pointer {
	return recv.native
}

// MenuLinkIterPrivateStruct creates an uninitialised MenuLinkIterPrivate.
func MenuLinkIterPrivateStruct() *MenuLinkIterPrivate {
	err := menuLinkIterPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := MenuLinkIterPrivateNewFromNative(menuLinkIterPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeMenuLinkIterPrivate)
	return structGo
}
func finalizeMenuLinkIterPrivate(obj *MenuLinkIterPrivate) {
	menuLinkIterPrivateStruct.Free(obj.native)
}

var menuModelClassStruct *gi.Struct
var menuModelClassStruct_Once sync.Once

func menuModelClassStruct_Set() error {
	var err error
	menuModelClassStruct_Once.Do(func() {
		menuModelClassStruct, err = gi.StructNew("Gio", "MenuModelClass")
	})
	return err
}

type MenuModelClass struct {
	native unsafe.Pointer
}

func MenuModelClassNewFromNative(native unsafe.Pointer) *MenuModelClass {
	instance := &MenuModelClass{native: native}

	return instance
}

/*
CastToMenuModelClass down casts any arbitrary Object to MenuModelClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a MenuModelClass.
*/
func (recv *MenuModelClass) CastToMenuModelClass(object *gobject.Object) *MenuModelClass {
	return MenuModelClassNewFromNative(object.Native())
}

func (recv *MenuModelClass) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'is_mutable' : for field getter : missing Type

// UNSUPPORTED : C value 'is_mutable' : for field setter : missing Type

// UNSUPPORTED : C value 'get_n_items' : for field getter : missing Type

// UNSUPPORTED : C value 'get_n_items' : for field setter : missing Type

// UNSUPPORTED : C value 'get_item_attributes' : for field getter : missing Type

// UNSUPPORTED : C value 'get_item_attributes' : for field setter : missing Type

// UNSUPPORTED : C value 'iterate_item_attributes' : for field getter : missing Type

// UNSUPPORTED : C value 'iterate_item_attributes' : for field setter : missing Type

// UNSUPPORTED : C value 'get_item_attribute_value' : for field getter : missing Type

// UNSUPPORTED : C value 'get_item_attribute_value' : for field setter : missing Type

// UNSUPPORTED : C value 'get_item_links' : for field getter : missing Type

// UNSUPPORTED : C value 'get_item_links' : for field setter : missing Type

// UNSUPPORTED : C value 'iterate_item_links' : for field getter : missing Type

// UNSUPPORTED : C value 'iterate_item_links' : for field setter : missing Type

// UNSUPPORTED : C value 'get_item_link' : for field getter : missing Type

// UNSUPPORTED : C value 'get_item_link' : for field setter : missing Type

// MenuModelClassStruct creates an uninitialised MenuModelClass.
func MenuModelClassStruct() *MenuModelClass {
	err := menuModelClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := MenuModelClassNewFromNative(menuModelClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeMenuModelClass)
	return structGo
}
func finalizeMenuModelClass(obj *MenuModelClass) {
	menuModelClassStruct.Free(obj.native)
}

var menuModelPrivateStruct *gi.Struct
var menuModelPrivateStruct_Once sync.Once

func menuModelPrivateStruct_Set() error {
	var err error
	menuModelPrivateStruct_Once.Do(func() {
		menuModelPrivateStruct, err = gi.StructNew("Gio", "MenuModelPrivate")
	})
	return err
}

type MenuModelPrivate struct {
	native unsafe.Pointer
}

func MenuModelPrivateNewFromNative(native unsafe.Pointer) *MenuModelPrivate {
	instance := &MenuModelPrivate{native: native}

	return instance
}

/*
CastToMenuModelPrivate down casts any arbitrary Object to MenuModelPrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a MenuModelPrivate.
*/
func (recv *MenuModelPrivate) CastToMenuModelPrivate(object *gobject.Object) *MenuModelPrivate {
	return MenuModelPrivateNewFromNative(object.Native())
}

func (recv *MenuModelPrivate) Native() unsafe.Pointer {
	return recv.native
}

// MenuModelPrivateStruct creates an uninitialised MenuModelPrivate.
func MenuModelPrivateStruct() *MenuModelPrivate {
	err := menuModelPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := MenuModelPrivateNewFromNative(menuModelPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeMenuModelPrivate)
	return structGo
}
func finalizeMenuModelPrivate(obj *MenuModelPrivate) {
	menuModelPrivateStruct.Free(obj.native)
}

var mountIfaceStruct *gi.Struct
var mountIfaceStruct_Once sync.Once

func mountIfaceStruct_Set() error {
	var err error
	mountIfaceStruct_Once.Do(func() {
		mountIfaceStruct, err = gi.StructNew("Gio", "MountIface")
	})
	return err
}

type MountIface struct {
	native unsafe.Pointer
}

func MountIfaceNewFromNative(native unsafe.Pointer) *MountIface {
	instance := &MountIface{native: native}

	return instance
}

/*
CastToMountIface down casts any arbitrary Object to MountIface.
Exercise care, as this is a potentially dangerous function
if the Object is not a MountIface.
*/
func (recv *MountIface) CastToMountIface(object *gobject.Object) *MountIface {
	return MountIfaceNewFromNative(object.Native())
}

func (recv *MountIface) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'g_iface' : for field getter : no Go type for 'GObject.TypeInterface'

// UNSUPPORTED : C value 'g_iface' : for field setter : no Go type for 'GObject.TypeInterface'

// UNSUPPORTED : C value 'changed' : for field getter : missing Type

// UNSUPPORTED : C value 'changed' : for field setter : missing Type

// UNSUPPORTED : C value 'unmounted' : for field getter : missing Type

// UNSUPPORTED : C value 'unmounted' : for field setter : missing Type

// UNSUPPORTED : C value 'get_root' : for field getter : missing Type

// UNSUPPORTED : C value 'get_root' : for field setter : missing Type

// UNSUPPORTED : C value 'get_name' : for field getter : missing Type

// UNSUPPORTED : C value 'get_name' : for field setter : missing Type

// UNSUPPORTED : C value 'get_icon' : for field getter : missing Type

// UNSUPPORTED : C value 'get_icon' : for field setter : missing Type

// UNSUPPORTED : C value 'get_uuid' : for field getter : missing Type

// UNSUPPORTED : C value 'get_uuid' : for field setter : missing Type

// UNSUPPORTED : C value 'get_volume' : for field getter : missing Type

// UNSUPPORTED : C value 'get_volume' : for field setter : missing Type

// UNSUPPORTED : C value 'get_drive' : for field getter : missing Type

// UNSUPPORTED : C value 'get_drive' : for field setter : missing Type

// UNSUPPORTED : C value 'can_unmount' : for field getter : missing Type

// UNSUPPORTED : C value 'can_unmount' : for field setter : missing Type

// UNSUPPORTED : C value 'can_eject' : for field getter : missing Type

// UNSUPPORTED : C value 'can_eject' : for field setter : missing Type

// UNSUPPORTED : C value 'unmount' : for field getter : missing Type

// UNSUPPORTED : C value 'unmount' : for field setter : missing Type

// UNSUPPORTED : C value 'unmount_finish' : for field getter : missing Type

// UNSUPPORTED : C value 'unmount_finish' : for field setter : missing Type

// UNSUPPORTED : C value 'eject' : for field getter : missing Type

// UNSUPPORTED : C value 'eject' : for field setter : missing Type

// UNSUPPORTED : C value 'eject_finish' : for field getter : missing Type

// UNSUPPORTED : C value 'eject_finish' : for field setter : missing Type

// UNSUPPORTED : C value 'remount' : for field getter : missing Type

// UNSUPPORTED : C value 'remount' : for field setter : missing Type

// UNSUPPORTED : C value 'remount_finish' : for field getter : missing Type

// UNSUPPORTED : C value 'remount_finish' : for field setter : missing Type

// UNSUPPORTED : C value 'guess_content_type' : for field getter : missing Type

// UNSUPPORTED : C value 'guess_content_type' : for field setter : missing Type

// UNSUPPORTED : C value 'guess_content_type_finish' : for field getter : missing Type

// UNSUPPORTED : C value 'guess_content_type_finish' : for field setter : missing Type

// UNSUPPORTED : C value 'guess_content_type_sync' : for field getter : missing Type

// UNSUPPORTED : C value 'guess_content_type_sync' : for field setter : missing Type

// UNSUPPORTED : C value 'pre_unmount' : for field getter : missing Type

// UNSUPPORTED : C value 'pre_unmount' : for field setter : missing Type

// UNSUPPORTED : C value 'unmount_with_operation' : for field getter : missing Type

// UNSUPPORTED : C value 'unmount_with_operation' : for field setter : missing Type

// UNSUPPORTED : C value 'unmount_with_operation_finish' : for field getter : missing Type

// UNSUPPORTED : C value 'unmount_with_operation_finish' : for field setter : missing Type

// UNSUPPORTED : C value 'eject_with_operation' : for field getter : missing Type

// UNSUPPORTED : C value 'eject_with_operation' : for field setter : missing Type

// UNSUPPORTED : C value 'eject_with_operation_finish' : for field getter : missing Type

// UNSUPPORTED : C value 'eject_with_operation_finish' : for field setter : missing Type

// UNSUPPORTED : C value 'get_default_location' : for field getter : missing Type

// UNSUPPORTED : C value 'get_default_location' : for field setter : missing Type

// UNSUPPORTED : C value 'get_sort_key' : for field getter : missing Type

// UNSUPPORTED : C value 'get_sort_key' : for field setter : missing Type

// UNSUPPORTED : C value 'get_symbolic_icon' : for field getter : missing Type

// UNSUPPORTED : C value 'get_symbolic_icon' : for field setter : missing Type

// MountIfaceStruct creates an uninitialised MountIface.
func MountIfaceStruct() *MountIface {
	err := mountIfaceStruct_Set()
	if err != nil {
		return nil
	}

	structGo := MountIfaceNewFromNative(mountIfaceStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeMountIface)
	return structGo
}
func finalizeMountIface(obj *MountIface) {
	mountIfaceStruct.Free(obj.native)
}

var mountOperationClassStruct *gi.Struct
var mountOperationClassStruct_Once sync.Once

func mountOperationClassStruct_Set() error {
	var err error
	mountOperationClassStruct_Once.Do(func() {
		mountOperationClassStruct, err = gi.StructNew("Gio", "MountOperationClass")
	})
	return err
}

type MountOperationClass struct {
	native unsafe.Pointer
}

func MountOperationClassNewFromNative(native unsafe.Pointer) *MountOperationClass {
	instance := &MountOperationClass{native: native}

	return instance
}

/*
CastToMountOperationClass down casts any arbitrary Object to MountOperationClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a MountOperationClass.
*/
func (recv *MountOperationClass) CastToMountOperationClass(object *gobject.Object) *MountOperationClass {
	return MountOperationClassNewFromNative(object.Native())
}

func (recv *MountOperationClass) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'ask_password' : for field getter : missing Type

// UNSUPPORTED : C value 'ask_password' : for field setter : missing Type

// UNSUPPORTED : C value 'ask_question' : for field getter : missing Type

// UNSUPPORTED : C value 'ask_question' : for field setter : missing Type

// UNSUPPORTED : C value 'reply' : for field getter : missing Type

// UNSUPPORTED : C value 'reply' : for field setter : missing Type

// UNSUPPORTED : C value 'aborted' : for field getter : missing Type

// UNSUPPORTED : C value 'aborted' : for field setter : missing Type

// UNSUPPORTED : C value 'show_processes' : for field getter : missing Type

// UNSUPPORTED : C value 'show_processes' : for field setter : missing Type

// UNSUPPORTED : C value 'show_unmount_progress' : for field getter : missing Type

// UNSUPPORTED : C value 'show_unmount_progress' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved4' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved5' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved5' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved6' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved6' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved7' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved7' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved8' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved8' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved9' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved9' : for field setter : missing Type

// MountOperationClassStruct creates an uninitialised MountOperationClass.
func MountOperationClassStruct() *MountOperationClass {
	err := mountOperationClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := MountOperationClassNewFromNative(mountOperationClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeMountOperationClass)
	return structGo
}
func finalizeMountOperationClass(obj *MountOperationClass) {
	mountOperationClassStruct.Free(obj.native)
}

var mountOperationPrivateStruct *gi.Struct
var mountOperationPrivateStruct_Once sync.Once

func mountOperationPrivateStruct_Set() error {
	var err error
	mountOperationPrivateStruct_Once.Do(func() {
		mountOperationPrivateStruct, err = gi.StructNew("Gio", "MountOperationPrivate")
	})
	return err
}

type MountOperationPrivate struct {
	native unsafe.Pointer
}

func MountOperationPrivateNewFromNative(native unsafe.Pointer) *MountOperationPrivate {
	instance := &MountOperationPrivate{native: native}

	return instance
}

/*
CastToMountOperationPrivate down casts any arbitrary Object to MountOperationPrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a MountOperationPrivate.
*/
func (recv *MountOperationPrivate) CastToMountOperationPrivate(object *gobject.Object) *MountOperationPrivate {
	return MountOperationPrivateNewFromNative(object.Native())
}

func (recv *MountOperationPrivate) Native() unsafe.Pointer {
	return recv.native
}

// MountOperationPrivateStruct creates an uninitialised MountOperationPrivate.
func MountOperationPrivateStruct() *MountOperationPrivate {
	err := mountOperationPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := MountOperationPrivateNewFromNative(mountOperationPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeMountOperationPrivate)
	return structGo
}
func finalizeMountOperationPrivate(obj *MountOperationPrivate) {
	mountOperationPrivateStruct.Free(obj.native)
}

var nativeSocketAddressClassStruct *gi.Struct
var nativeSocketAddressClassStruct_Once sync.Once

func nativeSocketAddressClassStruct_Set() error {
	var err error
	nativeSocketAddressClassStruct_Once.Do(func() {
		nativeSocketAddressClassStruct, err = gi.StructNew("Gio", "NativeSocketAddressClass")
	})
	return err
}

type NativeSocketAddressClass struct {
	native unsafe.Pointer
}

func NativeSocketAddressClassNewFromNative(native unsafe.Pointer) *NativeSocketAddressClass {
	instance := &NativeSocketAddressClass{native: native}

	return instance
}

/*
CastToNativeSocketAddressClass down casts any arbitrary Object to NativeSocketAddressClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a NativeSocketAddressClass.
*/
func (recv *NativeSocketAddressClass) CastToNativeSocketAddressClass(object *gobject.Object) *NativeSocketAddressClass {
	return NativeSocketAddressClassNewFromNative(object.Native())
}

func (recv *NativeSocketAddressClass) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentClass returns the C field 'parent_class'.
func (recv *NativeSocketAddressClass) FieldParentClass() *SocketAddressClass {
	argValue := gi.StructFieldGet(nativeSocketAddressClassStruct, recv.native, "parent_class")
	value := SocketAddressClassNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *NativeSocketAddressClass) SetFieldParentClass(value *SocketAddressClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.StructFieldSet(nativeSocketAddressClassStruct, recv.native, "parent_class", argValue)
}

// NativeSocketAddressClassStruct creates an uninitialised NativeSocketAddressClass.
func NativeSocketAddressClassStruct() *NativeSocketAddressClass {
	err := nativeSocketAddressClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := NativeSocketAddressClassNewFromNative(nativeSocketAddressClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeNativeSocketAddressClass)
	return structGo
}
func finalizeNativeSocketAddressClass(obj *NativeSocketAddressClass) {
	nativeSocketAddressClassStruct.Free(obj.native)
}

var nativeSocketAddressPrivateStruct *gi.Struct
var nativeSocketAddressPrivateStruct_Once sync.Once

func nativeSocketAddressPrivateStruct_Set() error {
	var err error
	nativeSocketAddressPrivateStruct_Once.Do(func() {
		nativeSocketAddressPrivateStruct, err = gi.StructNew("Gio", "NativeSocketAddressPrivate")
	})
	return err
}

type NativeSocketAddressPrivate struct {
	native unsafe.Pointer
}

func NativeSocketAddressPrivateNewFromNative(native unsafe.Pointer) *NativeSocketAddressPrivate {
	instance := &NativeSocketAddressPrivate{native: native}

	return instance
}

/*
CastToNativeSocketAddressPrivate down casts any arbitrary Object to NativeSocketAddressPrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a NativeSocketAddressPrivate.
*/
func (recv *NativeSocketAddressPrivate) CastToNativeSocketAddressPrivate(object *gobject.Object) *NativeSocketAddressPrivate {
	return NativeSocketAddressPrivateNewFromNative(object.Native())
}

func (recv *NativeSocketAddressPrivate) Native() unsafe.Pointer {
	return recv.native
}

// NativeSocketAddressPrivateStruct creates an uninitialised NativeSocketAddressPrivate.
func NativeSocketAddressPrivateStruct() *NativeSocketAddressPrivate {
	err := nativeSocketAddressPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := NativeSocketAddressPrivateNewFromNative(nativeSocketAddressPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeNativeSocketAddressPrivate)
	return structGo
}
func finalizeNativeSocketAddressPrivate(obj *NativeSocketAddressPrivate) {
	nativeSocketAddressPrivateStruct.Free(obj.native)
}

var nativeVolumeMonitorClassStruct *gi.Struct
var nativeVolumeMonitorClassStruct_Once sync.Once

func nativeVolumeMonitorClassStruct_Set() error {
	var err error
	nativeVolumeMonitorClassStruct_Once.Do(func() {
		nativeVolumeMonitorClassStruct, err = gi.StructNew("Gio", "NativeVolumeMonitorClass")
	})
	return err
}

type NativeVolumeMonitorClass struct {
	native unsafe.Pointer
}

func NativeVolumeMonitorClassNewFromNative(native unsafe.Pointer) *NativeVolumeMonitorClass {
	instance := &NativeVolumeMonitorClass{native: native}

	return instance
}

/*
CastToNativeVolumeMonitorClass down casts any arbitrary Object to NativeVolumeMonitorClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a NativeVolumeMonitorClass.
*/
func (recv *NativeVolumeMonitorClass) CastToNativeVolumeMonitorClass(object *gobject.Object) *NativeVolumeMonitorClass {
	return NativeVolumeMonitorClassNewFromNative(object.Native())
}

func (recv *NativeVolumeMonitorClass) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentClass returns the C field 'parent_class'.
func (recv *NativeVolumeMonitorClass) FieldParentClass() *VolumeMonitorClass {
	argValue := gi.StructFieldGet(nativeVolumeMonitorClassStruct, recv.native, "parent_class")
	value := VolumeMonitorClassNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *NativeVolumeMonitorClass) SetFieldParentClass(value *VolumeMonitorClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.StructFieldSet(nativeVolumeMonitorClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value 'get_mount_for_mount_path' : for field getter : missing Type

// UNSUPPORTED : C value 'get_mount_for_mount_path' : for field setter : missing Type

// NativeVolumeMonitorClassStruct creates an uninitialised NativeVolumeMonitorClass.
func NativeVolumeMonitorClassStruct() *NativeVolumeMonitorClass {
	err := nativeVolumeMonitorClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := NativeVolumeMonitorClassNewFromNative(nativeVolumeMonitorClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeNativeVolumeMonitorClass)
	return structGo
}
func finalizeNativeVolumeMonitorClass(obj *NativeVolumeMonitorClass) {
	nativeVolumeMonitorClassStruct.Free(obj.native)
}

var networkAddressClassStruct *gi.Struct
var networkAddressClassStruct_Once sync.Once

func networkAddressClassStruct_Set() error {
	var err error
	networkAddressClassStruct_Once.Do(func() {
		networkAddressClassStruct, err = gi.StructNew("Gio", "NetworkAddressClass")
	})
	return err
}

type NetworkAddressClass struct {
	native unsafe.Pointer
}

func NetworkAddressClassNewFromNative(native unsafe.Pointer) *NetworkAddressClass {
	instance := &NetworkAddressClass{native: native}

	return instance
}

/*
CastToNetworkAddressClass down casts any arbitrary Object to NetworkAddressClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a NetworkAddressClass.
*/
func (recv *NetworkAddressClass) CastToNetworkAddressClass(object *gobject.Object) *NetworkAddressClass {
	return NetworkAddressClassNewFromNative(object.Native())
}

func (recv *NetworkAddressClass) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// NetworkAddressClassStruct creates an uninitialised NetworkAddressClass.
func NetworkAddressClassStruct() *NetworkAddressClass {
	err := networkAddressClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := NetworkAddressClassNewFromNative(networkAddressClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeNetworkAddressClass)
	return structGo
}
func finalizeNetworkAddressClass(obj *NetworkAddressClass) {
	networkAddressClassStruct.Free(obj.native)
}

var networkAddressPrivateStruct *gi.Struct
var networkAddressPrivateStruct_Once sync.Once

func networkAddressPrivateStruct_Set() error {
	var err error
	networkAddressPrivateStruct_Once.Do(func() {
		networkAddressPrivateStruct, err = gi.StructNew("Gio", "NetworkAddressPrivate")
	})
	return err
}

type NetworkAddressPrivate struct {
	native unsafe.Pointer
}

func NetworkAddressPrivateNewFromNative(native unsafe.Pointer) *NetworkAddressPrivate {
	instance := &NetworkAddressPrivate{native: native}

	return instance
}

/*
CastToNetworkAddressPrivate down casts any arbitrary Object to NetworkAddressPrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a NetworkAddressPrivate.
*/
func (recv *NetworkAddressPrivate) CastToNetworkAddressPrivate(object *gobject.Object) *NetworkAddressPrivate {
	return NetworkAddressPrivateNewFromNative(object.Native())
}

func (recv *NetworkAddressPrivate) Native() unsafe.Pointer {
	return recv.native
}

// NetworkAddressPrivateStruct creates an uninitialised NetworkAddressPrivate.
func NetworkAddressPrivateStruct() *NetworkAddressPrivate {
	err := networkAddressPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := NetworkAddressPrivateNewFromNative(networkAddressPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeNetworkAddressPrivate)
	return structGo
}
func finalizeNetworkAddressPrivate(obj *NetworkAddressPrivate) {
	networkAddressPrivateStruct.Free(obj.native)
}

var networkMonitorInterfaceStruct *gi.Struct
var networkMonitorInterfaceStruct_Once sync.Once

func networkMonitorInterfaceStruct_Set() error {
	var err error
	networkMonitorInterfaceStruct_Once.Do(func() {
		networkMonitorInterfaceStruct, err = gi.StructNew("Gio", "NetworkMonitorInterface")
	})
	return err
}

type NetworkMonitorInterface struct {
	native unsafe.Pointer
}

func NetworkMonitorInterfaceNewFromNative(native unsafe.Pointer) *NetworkMonitorInterface {
	instance := &NetworkMonitorInterface{native: native}

	return instance
}

/*
CastToNetworkMonitorInterface down casts any arbitrary Object to NetworkMonitorInterface.
Exercise care, as this is a potentially dangerous function
if the Object is not a NetworkMonitorInterface.
*/
func (recv *NetworkMonitorInterface) CastToNetworkMonitorInterface(object *gobject.Object) *NetworkMonitorInterface {
	return NetworkMonitorInterfaceNewFromNative(object.Native())
}

func (recv *NetworkMonitorInterface) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'g_iface' : for field getter : no Go type for 'GObject.TypeInterface'

// UNSUPPORTED : C value 'g_iface' : for field setter : no Go type for 'GObject.TypeInterface'

// UNSUPPORTED : C value 'network_changed' : for field getter : missing Type

// UNSUPPORTED : C value 'network_changed' : for field setter : missing Type

// UNSUPPORTED : C value 'can_reach' : for field getter : missing Type

// UNSUPPORTED : C value 'can_reach' : for field setter : missing Type

// UNSUPPORTED : C value 'can_reach_async' : for field getter : missing Type

// UNSUPPORTED : C value 'can_reach_async' : for field setter : missing Type

// UNSUPPORTED : C value 'can_reach_finish' : for field getter : missing Type

// UNSUPPORTED : C value 'can_reach_finish' : for field setter : missing Type

// NetworkMonitorInterfaceStruct creates an uninitialised NetworkMonitorInterface.
func NetworkMonitorInterfaceStruct() *NetworkMonitorInterface {
	err := networkMonitorInterfaceStruct_Set()
	if err != nil {
		return nil
	}

	structGo := NetworkMonitorInterfaceNewFromNative(networkMonitorInterfaceStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeNetworkMonitorInterface)
	return structGo
}
func finalizeNetworkMonitorInterface(obj *NetworkMonitorInterface) {
	networkMonitorInterfaceStruct.Free(obj.native)
}

var networkServiceClassStruct *gi.Struct
var networkServiceClassStruct_Once sync.Once

func networkServiceClassStruct_Set() error {
	var err error
	networkServiceClassStruct_Once.Do(func() {
		networkServiceClassStruct, err = gi.StructNew("Gio", "NetworkServiceClass")
	})
	return err
}

type NetworkServiceClass struct {
	native unsafe.Pointer
}

func NetworkServiceClassNewFromNative(native unsafe.Pointer) *NetworkServiceClass {
	instance := &NetworkServiceClass{native: native}

	return instance
}

/*
CastToNetworkServiceClass down casts any arbitrary Object to NetworkServiceClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a NetworkServiceClass.
*/
func (recv *NetworkServiceClass) CastToNetworkServiceClass(object *gobject.Object) *NetworkServiceClass {
	return NetworkServiceClassNewFromNative(object.Native())
}

func (recv *NetworkServiceClass) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// NetworkServiceClassStruct creates an uninitialised NetworkServiceClass.
func NetworkServiceClassStruct() *NetworkServiceClass {
	err := networkServiceClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := NetworkServiceClassNewFromNative(networkServiceClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeNetworkServiceClass)
	return structGo
}
func finalizeNetworkServiceClass(obj *NetworkServiceClass) {
	networkServiceClassStruct.Free(obj.native)
}

var networkServicePrivateStruct *gi.Struct
var networkServicePrivateStruct_Once sync.Once

func networkServicePrivateStruct_Set() error {
	var err error
	networkServicePrivateStruct_Once.Do(func() {
		networkServicePrivateStruct, err = gi.StructNew("Gio", "NetworkServicePrivate")
	})
	return err
}

type NetworkServicePrivate struct {
	native unsafe.Pointer
}

func NetworkServicePrivateNewFromNative(native unsafe.Pointer) *NetworkServicePrivate {
	instance := &NetworkServicePrivate{native: native}

	return instance
}

/*
CastToNetworkServicePrivate down casts any arbitrary Object to NetworkServicePrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a NetworkServicePrivate.
*/
func (recv *NetworkServicePrivate) CastToNetworkServicePrivate(object *gobject.Object) *NetworkServicePrivate {
	return NetworkServicePrivateNewFromNative(object.Native())
}

func (recv *NetworkServicePrivate) Native() unsafe.Pointer {
	return recv.native
}

// NetworkServicePrivateStruct creates an uninitialised NetworkServicePrivate.
func NetworkServicePrivateStruct() *NetworkServicePrivate {
	err := networkServicePrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := NetworkServicePrivateNewFromNative(networkServicePrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeNetworkServicePrivate)
	return structGo
}
func finalizeNetworkServicePrivate(obj *NetworkServicePrivate) {
	networkServicePrivateStruct.Free(obj.native)
}

var outputMessageStruct *gi.Struct
var outputMessageStruct_Once sync.Once

func outputMessageStruct_Set() error {
	var err error
	outputMessageStruct_Once.Do(func() {
		outputMessageStruct, err = gi.StructNew("Gio", "OutputMessage")
	})
	return err
}

type OutputMessage struct {
	native unsafe.Pointer
}

func OutputMessageNewFromNative(native unsafe.Pointer) *OutputMessage {
	instance := &OutputMessage{native: native}

	return instance
}

/*
CastToOutputMessage down casts any arbitrary Object to OutputMessage.
Exercise care, as this is a potentially dangerous function
if the Object is not a OutputMessage.
*/
func (recv *OutputMessage) CastToOutputMessage(object *gobject.Object) *OutputMessage {
	return OutputMessageNewFromNative(object.Native())
}

func (recv *OutputMessage) Native() unsafe.Pointer {
	return recv.native
}

// FieldAddress returns the C field 'address'.
func (recv *OutputMessage) FieldAddress() *SocketAddress {
	argValue := gi.StructFieldGet(outputMessageStruct, recv.native, "address")
	value := SocketAddressNewFromNative(argValue.Pointer())
	return value
}

// SetFieldAddress sets the value of the C field 'address'.
func (recv *OutputMessage) SetFieldAddress(value *SocketAddress) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.StructFieldSet(outputMessageStruct, recv.native, "address", argValue)
}

// FieldVectors returns the C field 'vectors'.
func (recv *OutputMessage) FieldVectors() *OutputVector {
	argValue := gi.StructFieldGet(outputMessageStruct, recv.native, "vectors")
	value := OutputVectorNewFromNative(argValue.Pointer())
	return value
}

// SetFieldVectors sets the value of the C field 'vectors'.
func (recv *OutputMessage) SetFieldVectors(value *OutputVector) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.StructFieldSet(outputMessageStruct, recv.native, "vectors", argValue)
}

// FieldNumVectors returns the C field 'num_vectors'.
func (recv *OutputMessage) FieldNumVectors() uint32 {
	argValue := gi.StructFieldGet(outputMessageStruct, recv.native, "num_vectors")
	value := argValue.Uint32()
	return value
}

// SetFieldNumVectors sets the value of the C field 'num_vectors'.
func (recv *OutputMessage) SetFieldNumVectors(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.StructFieldSet(outputMessageStruct, recv.native, "num_vectors", argValue)
}

// FieldBytesSent returns the C field 'bytes_sent'.
func (recv *OutputMessage) FieldBytesSent() uint32 {
	argValue := gi.StructFieldGet(outputMessageStruct, recv.native, "bytes_sent")
	value := argValue.Uint32()
	return value
}

// SetFieldBytesSent sets the value of the C field 'bytes_sent'.
func (recv *OutputMessage) SetFieldBytesSent(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.StructFieldSet(outputMessageStruct, recv.native, "bytes_sent", argValue)
}

// UNSUPPORTED : C value 'control_messages' : for field getter : missing Type

// UNSUPPORTED : C value 'control_messages' : for field setter : missing Type

// FieldNumControlMessages returns the C field 'num_control_messages'.
func (recv *OutputMessage) FieldNumControlMessages() uint32 {
	argValue := gi.StructFieldGet(outputMessageStruct, recv.native, "num_control_messages")
	value := argValue.Uint32()
	return value
}

// SetFieldNumControlMessages sets the value of the C field 'num_control_messages'.
func (recv *OutputMessage) SetFieldNumControlMessages(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.StructFieldSet(outputMessageStruct, recv.native, "num_control_messages", argValue)
}

// OutputMessageStruct creates an uninitialised OutputMessage.
func OutputMessageStruct() *OutputMessage {
	err := outputMessageStruct_Set()
	if err != nil {
		return nil
	}

	structGo := OutputMessageNewFromNative(outputMessageStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeOutputMessage)
	return structGo
}
func finalizeOutputMessage(obj *OutputMessage) {
	outputMessageStruct.Free(obj.native)
}

var outputStreamClassStruct *gi.Struct
var outputStreamClassStruct_Once sync.Once

func outputStreamClassStruct_Set() error {
	var err error
	outputStreamClassStruct_Once.Do(func() {
		outputStreamClassStruct, err = gi.StructNew("Gio", "OutputStreamClass")
	})
	return err
}

type OutputStreamClass struct {
	native unsafe.Pointer
}

func OutputStreamClassNewFromNative(native unsafe.Pointer) *OutputStreamClass {
	instance := &OutputStreamClass{native: native}

	return instance
}

/*
CastToOutputStreamClass down casts any arbitrary Object to OutputStreamClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a OutputStreamClass.
*/
func (recv *OutputStreamClass) CastToOutputStreamClass(object *gobject.Object) *OutputStreamClass {
	return OutputStreamClassNewFromNative(object.Native())
}

func (recv *OutputStreamClass) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'write_fn' : for field getter : missing Type

// UNSUPPORTED : C value 'write_fn' : for field setter : missing Type

// UNSUPPORTED : C value 'splice' : for field getter : missing Type

// UNSUPPORTED : C value 'splice' : for field setter : missing Type

// UNSUPPORTED : C value 'flush' : for field getter : missing Type

// UNSUPPORTED : C value 'flush' : for field setter : missing Type

// UNSUPPORTED : C value 'close_fn' : for field getter : missing Type

// UNSUPPORTED : C value 'close_fn' : for field setter : missing Type

// UNSUPPORTED : C value 'write_async' : for field getter : missing Type

// UNSUPPORTED : C value 'write_async' : for field setter : missing Type

// UNSUPPORTED : C value 'write_finish' : for field getter : missing Type

// UNSUPPORTED : C value 'write_finish' : for field setter : missing Type

// UNSUPPORTED : C value 'splice_async' : for field getter : missing Type

// UNSUPPORTED : C value 'splice_async' : for field setter : missing Type

// UNSUPPORTED : C value 'splice_finish' : for field getter : missing Type

// UNSUPPORTED : C value 'splice_finish' : for field setter : missing Type

// UNSUPPORTED : C value 'flush_async' : for field getter : missing Type

// UNSUPPORTED : C value 'flush_async' : for field setter : missing Type

// UNSUPPORTED : C value 'flush_finish' : for field getter : missing Type

// UNSUPPORTED : C value 'flush_finish' : for field setter : missing Type

// UNSUPPORTED : C value 'close_async' : for field getter : missing Type

// UNSUPPORTED : C value 'close_async' : for field setter : missing Type

// UNSUPPORTED : C value 'close_finish' : for field getter : missing Type

// UNSUPPORTED : C value 'close_finish' : for field setter : missing Type

// UNSUPPORTED : C value 'writev_fn' : for field getter : missing Type

// UNSUPPORTED : C value 'writev_fn' : for field setter : missing Type

// UNSUPPORTED : C value 'writev_async' : for field getter : missing Type

// UNSUPPORTED : C value 'writev_async' : for field setter : missing Type

// UNSUPPORTED : C value 'writev_finish' : for field getter : missing Type

// UNSUPPORTED : C value 'writev_finish' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved4' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved5' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved5' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved6' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved6' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved7' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved7' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved8' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved8' : for field setter : missing Type

// OutputStreamClassStruct creates an uninitialised OutputStreamClass.
func OutputStreamClassStruct() *OutputStreamClass {
	err := outputStreamClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := OutputStreamClassNewFromNative(outputStreamClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeOutputStreamClass)
	return structGo
}
func finalizeOutputStreamClass(obj *OutputStreamClass) {
	outputStreamClassStruct.Free(obj.native)
}

var outputStreamPrivateStruct *gi.Struct
var outputStreamPrivateStruct_Once sync.Once

func outputStreamPrivateStruct_Set() error {
	var err error
	outputStreamPrivateStruct_Once.Do(func() {
		outputStreamPrivateStruct, err = gi.StructNew("Gio", "OutputStreamPrivate")
	})
	return err
}

type OutputStreamPrivate struct {
	native unsafe.Pointer
}

func OutputStreamPrivateNewFromNative(native unsafe.Pointer) *OutputStreamPrivate {
	instance := &OutputStreamPrivate{native: native}

	return instance
}

/*
CastToOutputStreamPrivate down casts any arbitrary Object to OutputStreamPrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a OutputStreamPrivate.
*/
func (recv *OutputStreamPrivate) CastToOutputStreamPrivate(object *gobject.Object) *OutputStreamPrivate {
	return OutputStreamPrivateNewFromNative(object.Native())
}

func (recv *OutputStreamPrivate) Native() unsafe.Pointer {
	return recv.native
}

// OutputStreamPrivateStruct creates an uninitialised OutputStreamPrivate.
func OutputStreamPrivateStruct() *OutputStreamPrivate {
	err := outputStreamPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := OutputStreamPrivateNewFromNative(outputStreamPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeOutputStreamPrivate)
	return structGo
}
func finalizeOutputStreamPrivate(obj *OutputStreamPrivate) {
	outputStreamPrivateStruct.Free(obj.native)
}

var outputVectorStruct *gi.Struct
var outputVectorStruct_Once sync.Once

func outputVectorStruct_Set() error {
	var err error
	outputVectorStruct_Once.Do(func() {
		outputVectorStruct, err = gi.StructNew("Gio", "OutputVector")
	})
	return err
}

type OutputVector struct {
	native unsafe.Pointer
}

func OutputVectorNewFromNative(native unsafe.Pointer) *OutputVector {
	instance := &OutputVector{native: native}

	return instance
}

/*
CastToOutputVector down casts any arbitrary Object to OutputVector.
Exercise care, as this is a potentially dangerous function
if the Object is not a OutputVector.
*/
func (recv *OutputVector) CastToOutputVector(object *gobject.Object) *OutputVector {
	return OutputVectorNewFromNative(object.Native())
}

func (recv *OutputVector) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'buffer' : for field getter : no Go type for 'gpointer'

// UNSUPPORTED : C value 'buffer' : for field setter : no Go type for 'gpointer'

// FieldSize returns the C field 'size'.
func (recv *OutputVector) FieldSize() uint64 {
	argValue := gi.StructFieldGet(outputVectorStruct, recv.native, "size")
	value := argValue.Uint64()
	return value
}

// SetFieldSize sets the value of the C field 'size'.
func (recv *OutputVector) SetFieldSize(value uint64) {
	var argValue gi.Argument
	argValue.SetUint64(value)
	gi.StructFieldSet(outputVectorStruct, recv.native, "size", argValue)
}

// OutputVectorStruct creates an uninitialised OutputVector.
func OutputVectorStruct() *OutputVector {
	err := outputVectorStruct_Set()
	if err != nil {
		return nil
	}

	structGo := OutputVectorNewFromNative(outputVectorStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeOutputVector)
	return structGo
}
func finalizeOutputVector(obj *OutputVector) {
	outputVectorStruct.Free(obj.native)
}

var permissionClassStruct *gi.Struct
var permissionClassStruct_Once sync.Once

func permissionClassStruct_Set() error {
	var err error
	permissionClassStruct_Once.Do(func() {
		permissionClassStruct, err = gi.StructNew("Gio", "PermissionClass")
	})
	return err
}

type PermissionClass struct {
	native unsafe.Pointer
}

func PermissionClassNewFromNative(native unsafe.Pointer) *PermissionClass {
	instance := &PermissionClass{native: native}

	return instance
}

/*
CastToPermissionClass down casts any arbitrary Object to PermissionClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a PermissionClass.
*/
func (recv *PermissionClass) CastToPermissionClass(object *gobject.Object) *PermissionClass {
	return PermissionClassNewFromNative(object.Native())
}

func (recv *PermissionClass) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'acquire' : for field getter : missing Type

// UNSUPPORTED : C value 'acquire' : for field setter : missing Type

// UNSUPPORTED : C value 'acquire_async' : for field getter : missing Type

// UNSUPPORTED : C value 'acquire_async' : for field setter : missing Type

// UNSUPPORTED : C value 'acquire_finish' : for field getter : missing Type

// UNSUPPORTED : C value 'acquire_finish' : for field setter : missing Type

// UNSUPPORTED : C value 'release' : for field getter : missing Type

// UNSUPPORTED : C value 'release' : for field setter : missing Type

// UNSUPPORTED : C value 'release_async' : for field getter : missing Type

// UNSUPPORTED : C value 'release_async' : for field setter : missing Type

// UNSUPPORTED : C value 'release_finish' : for field getter : missing Type

// UNSUPPORTED : C value 'release_finish' : for field setter : missing Type

// UNSUPPORTED : C value 'reserved' : for field getter : missing Type

// UNSUPPORTED : C value 'reserved' : for field setter : missing Type

// PermissionClassStruct creates an uninitialised PermissionClass.
func PermissionClassStruct() *PermissionClass {
	err := permissionClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := PermissionClassNewFromNative(permissionClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizePermissionClass)
	return structGo
}
func finalizePermissionClass(obj *PermissionClass) {
	permissionClassStruct.Free(obj.native)
}

var permissionPrivateStruct *gi.Struct
var permissionPrivateStruct_Once sync.Once

func permissionPrivateStruct_Set() error {
	var err error
	permissionPrivateStruct_Once.Do(func() {
		permissionPrivateStruct, err = gi.StructNew("Gio", "PermissionPrivate")
	})
	return err
}

type PermissionPrivate struct {
	native unsafe.Pointer
}

func PermissionPrivateNewFromNative(native unsafe.Pointer) *PermissionPrivate {
	instance := &PermissionPrivate{native: native}

	return instance
}

/*
CastToPermissionPrivate down casts any arbitrary Object to PermissionPrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a PermissionPrivate.
*/
func (recv *PermissionPrivate) CastToPermissionPrivate(object *gobject.Object) *PermissionPrivate {
	return PermissionPrivateNewFromNative(object.Native())
}

func (recv *PermissionPrivate) Native() unsafe.Pointer {
	return recv.native
}

// PermissionPrivateStruct creates an uninitialised PermissionPrivate.
func PermissionPrivateStruct() *PermissionPrivate {
	err := permissionPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := PermissionPrivateNewFromNative(permissionPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizePermissionPrivate)
	return structGo
}
func finalizePermissionPrivate(obj *PermissionPrivate) {
	permissionPrivateStruct.Free(obj.native)
}

var pollableInputStreamInterfaceStruct *gi.Struct
var pollableInputStreamInterfaceStruct_Once sync.Once

func pollableInputStreamInterfaceStruct_Set() error {
	var err error
	pollableInputStreamInterfaceStruct_Once.Do(func() {
		pollableInputStreamInterfaceStruct, err = gi.StructNew("Gio", "PollableInputStreamInterface")
	})
	return err
}

type PollableInputStreamInterface struct {
	native unsafe.Pointer
}

func PollableInputStreamInterfaceNewFromNative(native unsafe.Pointer) *PollableInputStreamInterface {
	instance := &PollableInputStreamInterface{native: native}

	return instance
}

/*
CastToPollableInputStreamInterface down casts any arbitrary Object to PollableInputStreamInterface.
Exercise care, as this is a potentially dangerous function
if the Object is not a PollableInputStreamInterface.
*/
func (recv *PollableInputStreamInterface) CastToPollableInputStreamInterface(object *gobject.Object) *PollableInputStreamInterface {
	return PollableInputStreamInterfaceNewFromNative(object.Native())
}

func (recv *PollableInputStreamInterface) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'g_iface' : for field getter : no Go type for 'GObject.TypeInterface'

// UNSUPPORTED : C value 'g_iface' : for field setter : no Go type for 'GObject.TypeInterface'

// UNSUPPORTED : C value 'can_poll' : for field getter : missing Type

// UNSUPPORTED : C value 'can_poll' : for field setter : missing Type

// UNSUPPORTED : C value 'is_readable' : for field getter : missing Type

// UNSUPPORTED : C value 'is_readable' : for field setter : missing Type

// UNSUPPORTED : C value 'create_source' : for field getter : missing Type

// UNSUPPORTED : C value 'create_source' : for field setter : missing Type

// UNSUPPORTED : C value 'read_nonblocking' : for field getter : missing Type

// UNSUPPORTED : C value 'read_nonblocking' : for field setter : missing Type

// PollableInputStreamInterfaceStruct creates an uninitialised PollableInputStreamInterface.
func PollableInputStreamInterfaceStruct() *PollableInputStreamInterface {
	err := pollableInputStreamInterfaceStruct_Set()
	if err != nil {
		return nil
	}

	structGo := PollableInputStreamInterfaceNewFromNative(pollableInputStreamInterfaceStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizePollableInputStreamInterface)
	return structGo
}
func finalizePollableInputStreamInterface(obj *PollableInputStreamInterface) {
	pollableInputStreamInterfaceStruct.Free(obj.native)
}

var pollableOutputStreamInterfaceStruct *gi.Struct
var pollableOutputStreamInterfaceStruct_Once sync.Once

func pollableOutputStreamInterfaceStruct_Set() error {
	var err error
	pollableOutputStreamInterfaceStruct_Once.Do(func() {
		pollableOutputStreamInterfaceStruct, err = gi.StructNew("Gio", "PollableOutputStreamInterface")
	})
	return err
}

type PollableOutputStreamInterface struct {
	native unsafe.Pointer
}

func PollableOutputStreamInterfaceNewFromNative(native unsafe.Pointer) *PollableOutputStreamInterface {
	instance := &PollableOutputStreamInterface{native: native}

	return instance
}

/*
CastToPollableOutputStreamInterface down casts any arbitrary Object to PollableOutputStreamInterface.
Exercise care, as this is a potentially dangerous function
if the Object is not a PollableOutputStreamInterface.
*/
func (recv *PollableOutputStreamInterface) CastToPollableOutputStreamInterface(object *gobject.Object) *PollableOutputStreamInterface {
	return PollableOutputStreamInterfaceNewFromNative(object.Native())
}

func (recv *PollableOutputStreamInterface) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'g_iface' : for field getter : no Go type for 'GObject.TypeInterface'

// UNSUPPORTED : C value 'g_iface' : for field setter : no Go type for 'GObject.TypeInterface'

// UNSUPPORTED : C value 'can_poll' : for field getter : missing Type

// UNSUPPORTED : C value 'can_poll' : for field setter : missing Type

// UNSUPPORTED : C value 'is_writable' : for field getter : missing Type

// UNSUPPORTED : C value 'is_writable' : for field setter : missing Type

// UNSUPPORTED : C value 'create_source' : for field getter : missing Type

// UNSUPPORTED : C value 'create_source' : for field setter : missing Type

// UNSUPPORTED : C value 'write_nonblocking' : for field getter : missing Type

// UNSUPPORTED : C value 'write_nonblocking' : for field setter : missing Type

// UNSUPPORTED : C value 'writev_nonblocking' : for field getter : missing Type

// UNSUPPORTED : C value 'writev_nonblocking' : for field setter : missing Type

// PollableOutputStreamInterfaceStruct creates an uninitialised PollableOutputStreamInterface.
func PollableOutputStreamInterfaceStruct() *PollableOutputStreamInterface {
	err := pollableOutputStreamInterfaceStruct_Set()
	if err != nil {
		return nil
	}

	structGo := PollableOutputStreamInterfaceNewFromNative(pollableOutputStreamInterfaceStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizePollableOutputStreamInterface)
	return structGo
}
func finalizePollableOutputStreamInterface(obj *PollableOutputStreamInterface) {
	pollableOutputStreamInterfaceStruct.Free(obj.native)
}

var proxyAddressClassStruct *gi.Struct
var proxyAddressClassStruct_Once sync.Once

func proxyAddressClassStruct_Set() error {
	var err error
	proxyAddressClassStruct_Once.Do(func() {
		proxyAddressClassStruct, err = gi.StructNew("Gio", "ProxyAddressClass")
	})
	return err
}

type ProxyAddressClass struct {
	native unsafe.Pointer
}

func ProxyAddressClassNewFromNative(native unsafe.Pointer) *ProxyAddressClass {
	instance := &ProxyAddressClass{native: native}

	return instance
}

/*
CastToProxyAddressClass down casts any arbitrary Object to ProxyAddressClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a ProxyAddressClass.
*/
func (recv *ProxyAddressClass) CastToProxyAddressClass(object *gobject.Object) *ProxyAddressClass {
	return ProxyAddressClassNewFromNative(object.Native())
}

func (recv *ProxyAddressClass) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentClass returns the C field 'parent_class'.
func (recv *ProxyAddressClass) FieldParentClass() *InetSocketAddressClass {
	argValue := gi.StructFieldGet(proxyAddressClassStruct, recv.native, "parent_class")
	value := InetSocketAddressClassNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *ProxyAddressClass) SetFieldParentClass(value *InetSocketAddressClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.StructFieldSet(proxyAddressClassStruct, recv.native, "parent_class", argValue)
}

// ProxyAddressClassStruct creates an uninitialised ProxyAddressClass.
func ProxyAddressClassStruct() *ProxyAddressClass {
	err := proxyAddressClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := ProxyAddressClassNewFromNative(proxyAddressClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeProxyAddressClass)
	return structGo
}
func finalizeProxyAddressClass(obj *ProxyAddressClass) {
	proxyAddressClassStruct.Free(obj.native)
}

var proxyAddressEnumeratorClassStruct *gi.Struct
var proxyAddressEnumeratorClassStruct_Once sync.Once

func proxyAddressEnumeratorClassStruct_Set() error {
	var err error
	proxyAddressEnumeratorClassStruct_Once.Do(func() {
		proxyAddressEnumeratorClassStruct, err = gi.StructNew("Gio", "ProxyAddressEnumeratorClass")
	})
	return err
}

type ProxyAddressEnumeratorClass struct {
	native unsafe.Pointer
}

func ProxyAddressEnumeratorClassNewFromNative(native unsafe.Pointer) *ProxyAddressEnumeratorClass {
	instance := &ProxyAddressEnumeratorClass{native: native}

	return instance
}

/*
CastToProxyAddressEnumeratorClass down casts any arbitrary Object to ProxyAddressEnumeratorClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a ProxyAddressEnumeratorClass.
*/
func (recv *ProxyAddressEnumeratorClass) CastToProxyAddressEnumeratorClass(object *gobject.Object) *ProxyAddressEnumeratorClass {
	return ProxyAddressEnumeratorClassNewFromNative(object.Native())
}

func (recv *ProxyAddressEnumeratorClass) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value '_g_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved4' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved5' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved5' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved6' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved6' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved7' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved7' : for field setter : missing Type

// ProxyAddressEnumeratorClassStruct creates an uninitialised ProxyAddressEnumeratorClass.
func ProxyAddressEnumeratorClassStruct() *ProxyAddressEnumeratorClass {
	err := proxyAddressEnumeratorClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := ProxyAddressEnumeratorClassNewFromNative(proxyAddressEnumeratorClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeProxyAddressEnumeratorClass)
	return structGo
}
func finalizeProxyAddressEnumeratorClass(obj *ProxyAddressEnumeratorClass) {
	proxyAddressEnumeratorClassStruct.Free(obj.native)
}

var proxyAddressEnumeratorPrivateStruct *gi.Struct
var proxyAddressEnumeratorPrivateStruct_Once sync.Once

func proxyAddressEnumeratorPrivateStruct_Set() error {
	var err error
	proxyAddressEnumeratorPrivateStruct_Once.Do(func() {
		proxyAddressEnumeratorPrivateStruct, err = gi.StructNew("Gio", "ProxyAddressEnumeratorPrivate")
	})
	return err
}

type ProxyAddressEnumeratorPrivate struct {
	native unsafe.Pointer
}

func ProxyAddressEnumeratorPrivateNewFromNative(native unsafe.Pointer) *ProxyAddressEnumeratorPrivate {
	instance := &ProxyAddressEnumeratorPrivate{native: native}

	return instance
}

/*
CastToProxyAddressEnumeratorPrivate down casts any arbitrary Object to ProxyAddressEnumeratorPrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a ProxyAddressEnumeratorPrivate.
*/
func (recv *ProxyAddressEnumeratorPrivate) CastToProxyAddressEnumeratorPrivate(object *gobject.Object) *ProxyAddressEnumeratorPrivate {
	return ProxyAddressEnumeratorPrivateNewFromNative(object.Native())
}

func (recv *ProxyAddressEnumeratorPrivate) Native() unsafe.Pointer {
	return recv.native
}

// ProxyAddressEnumeratorPrivateStruct creates an uninitialised ProxyAddressEnumeratorPrivate.
func ProxyAddressEnumeratorPrivateStruct() *ProxyAddressEnumeratorPrivate {
	err := proxyAddressEnumeratorPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := ProxyAddressEnumeratorPrivateNewFromNative(proxyAddressEnumeratorPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeProxyAddressEnumeratorPrivate)
	return structGo
}
func finalizeProxyAddressEnumeratorPrivate(obj *ProxyAddressEnumeratorPrivate) {
	proxyAddressEnumeratorPrivateStruct.Free(obj.native)
}

var proxyAddressPrivateStruct *gi.Struct
var proxyAddressPrivateStruct_Once sync.Once

func proxyAddressPrivateStruct_Set() error {
	var err error
	proxyAddressPrivateStruct_Once.Do(func() {
		proxyAddressPrivateStruct, err = gi.StructNew("Gio", "ProxyAddressPrivate")
	})
	return err
}

type ProxyAddressPrivate struct {
	native unsafe.Pointer
}

func ProxyAddressPrivateNewFromNative(native unsafe.Pointer) *ProxyAddressPrivate {
	instance := &ProxyAddressPrivate{native: native}

	return instance
}

/*
CastToProxyAddressPrivate down casts any arbitrary Object to ProxyAddressPrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a ProxyAddressPrivate.
*/
func (recv *ProxyAddressPrivate) CastToProxyAddressPrivate(object *gobject.Object) *ProxyAddressPrivate {
	return ProxyAddressPrivateNewFromNative(object.Native())
}

func (recv *ProxyAddressPrivate) Native() unsafe.Pointer {
	return recv.native
}

// ProxyAddressPrivateStruct creates an uninitialised ProxyAddressPrivate.
func ProxyAddressPrivateStruct() *ProxyAddressPrivate {
	err := proxyAddressPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := ProxyAddressPrivateNewFromNative(proxyAddressPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeProxyAddressPrivate)
	return structGo
}
func finalizeProxyAddressPrivate(obj *ProxyAddressPrivate) {
	proxyAddressPrivateStruct.Free(obj.native)
}

var proxyInterfaceStruct *gi.Struct
var proxyInterfaceStruct_Once sync.Once

func proxyInterfaceStruct_Set() error {
	var err error
	proxyInterfaceStruct_Once.Do(func() {
		proxyInterfaceStruct, err = gi.StructNew("Gio", "ProxyInterface")
	})
	return err
}

type ProxyInterface struct {
	native unsafe.Pointer
}

func ProxyInterfaceNewFromNative(native unsafe.Pointer) *ProxyInterface {
	instance := &ProxyInterface{native: native}

	return instance
}

/*
CastToProxyInterface down casts any arbitrary Object to ProxyInterface.
Exercise care, as this is a potentially dangerous function
if the Object is not a ProxyInterface.
*/
func (recv *ProxyInterface) CastToProxyInterface(object *gobject.Object) *ProxyInterface {
	return ProxyInterfaceNewFromNative(object.Native())
}

func (recv *ProxyInterface) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'g_iface' : for field getter : no Go type for 'GObject.TypeInterface'

// UNSUPPORTED : C value 'g_iface' : for field setter : no Go type for 'GObject.TypeInterface'

// UNSUPPORTED : C value 'connect' : for field getter : missing Type

// UNSUPPORTED : C value 'connect' : for field setter : missing Type

// UNSUPPORTED : C value 'connect_async' : for field getter : missing Type

// UNSUPPORTED : C value 'connect_async' : for field setter : missing Type

// UNSUPPORTED : C value 'connect_finish' : for field getter : missing Type

// UNSUPPORTED : C value 'connect_finish' : for field setter : missing Type

// UNSUPPORTED : C value 'supports_hostname' : for field getter : missing Type

// UNSUPPORTED : C value 'supports_hostname' : for field setter : missing Type

// ProxyInterfaceStruct creates an uninitialised ProxyInterface.
func ProxyInterfaceStruct() *ProxyInterface {
	err := proxyInterfaceStruct_Set()
	if err != nil {
		return nil
	}

	structGo := ProxyInterfaceNewFromNative(proxyInterfaceStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeProxyInterface)
	return structGo
}
func finalizeProxyInterface(obj *ProxyInterface) {
	proxyInterfaceStruct.Free(obj.native)
}

var proxyResolverInterfaceStruct *gi.Struct
var proxyResolverInterfaceStruct_Once sync.Once

func proxyResolverInterfaceStruct_Set() error {
	var err error
	proxyResolverInterfaceStruct_Once.Do(func() {
		proxyResolverInterfaceStruct, err = gi.StructNew("Gio", "ProxyResolverInterface")
	})
	return err
}

type ProxyResolverInterface struct {
	native unsafe.Pointer
}

func ProxyResolverInterfaceNewFromNative(native unsafe.Pointer) *ProxyResolverInterface {
	instance := &ProxyResolverInterface{native: native}

	return instance
}

/*
CastToProxyResolverInterface down casts any arbitrary Object to ProxyResolverInterface.
Exercise care, as this is a potentially dangerous function
if the Object is not a ProxyResolverInterface.
*/
func (recv *ProxyResolverInterface) CastToProxyResolverInterface(object *gobject.Object) *ProxyResolverInterface {
	return ProxyResolverInterfaceNewFromNative(object.Native())
}

func (recv *ProxyResolverInterface) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'g_iface' : for field getter : no Go type for 'GObject.TypeInterface'

// UNSUPPORTED : C value 'g_iface' : for field setter : no Go type for 'GObject.TypeInterface'

// UNSUPPORTED : C value 'is_supported' : for field getter : missing Type

// UNSUPPORTED : C value 'is_supported' : for field setter : missing Type

// UNSUPPORTED : C value 'lookup' : for field getter : missing Type

// UNSUPPORTED : C value 'lookup' : for field setter : missing Type

// UNSUPPORTED : C value 'lookup_async' : for field getter : missing Type

// UNSUPPORTED : C value 'lookup_async' : for field setter : missing Type

// UNSUPPORTED : C value 'lookup_finish' : for field getter : missing Type

// UNSUPPORTED : C value 'lookup_finish' : for field setter : missing Type

// ProxyResolverInterfaceStruct creates an uninitialised ProxyResolverInterface.
func ProxyResolverInterfaceStruct() *ProxyResolverInterface {
	err := proxyResolverInterfaceStruct_Set()
	if err != nil {
		return nil
	}

	structGo := ProxyResolverInterfaceNewFromNative(proxyResolverInterfaceStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeProxyResolverInterface)
	return structGo
}
func finalizeProxyResolverInterface(obj *ProxyResolverInterface) {
	proxyResolverInterfaceStruct.Free(obj.native)
}

var remoteActionGroupInterfaceStruct *gi.Struct
var remoteActionGroupInterfaceStruct_Once sync.Once

func remoteActionGroupInterfaceStruct_Set() error {
	var err error
	remoteActionGroupInterfaceStruct_Once.Do(func() {
		remoteActionGroupInterfaceStruct, err = gi.StructNew("Gio", "RemoteActionGroupInterface")
	})
	return err
}

type RemoteActionGroupInterface struct {
	native unsafe.Pointer
}

func RemoteActionGroupInterfaceNewFromNative(native unsafe.Pointer) *RemoteActionGroupInterface {
	instance := &RemoteActionGroupInterface{native: native}

	return instance
}

/*
CastToRemoteActionGroupInterface down casts any arbitrary Object to RemoteActionGroupInterface.
Exercise care, as this is a potentially dangerous function
if the Object is not a RemoteActionGroupInterface.
*/
func (recv *RemoteActionGroupInterface) CastToRemoteActionGroupInterface(object *gobject.Object) *RemoteActionGroupInterface {
	return RemoteActionGroupInterfaceNewFromNative(object.Native())
}

func (recv *RemoteActionGroupInterface) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'g_iface' : for field getter : no Go type for 'GObject.TypeInterface'

// UNSUPPORTED : C value 'g_iface' : for field setter : no Go type for 'GObject.TypeInterface'

// UNSUPPORTED : C value 'activate_action_full' : for field getter : missing Type

// UNSUPPORTED : C value 'activate_action_full' : for field setter : missing Type

// UNSUPPORTED : C value 'change_action_state_full' : for field getter : missing Type

// UNSUPPORTED : C value 'change_action_state_full' : for field setter : missing Type

// RemoteActionGroupInterfaceStruct creates an uninitialised RemoteActionGroupInterface.
func RemoteActionGroupInterfaceStruct() *RemoteActionGroupInterface {
	err := remoteActionGroupInterfaceStruct_Set()
	if err != nil {
		return nil
	}

	structGo := RemoteActionGroupInterfaceNewFromNative(remoteActionGroupInterfaceStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeRemoteActionGroupInterface)
	return structGo
}
func finalizeRemoteActionGroupInterface(obj *RemoteActionGroupInterface) {
	remoteActionGroupInterfaceStruct.Free(obj.native)
}

var resolverClassStruct *gi.Struct
var resolverClassStruct_Once sync.Once

func resolverClassStruct_Set() error {
	var err error
	resolverClassStruct_Once.Do(func() {
		resolverClassStruct, err = gi.StructNew("Gio", "ResolverClass")
	})
	return err
}

type ResolverClass struct {
	native unsafe.Pointer
}

func ResolverClassNewFromNative(native unsafe.Pointer) *ResolverClass {
	instance := &ResolverClass{native: native}

	return instance
}

/*
CastToResolverClass down casts any arbitrary Object to ResolverClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a ResolverClass.
*/
func (recv *ResolverClass) CastToResolverClass(object *gobject.Object) *ResolverClass {
	return ResolverClassNewFromNative(object.Native())
}

func (recv *ResolverClass) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'reload' : for field getter : missing Type

// UNSUPPORTED : C value 'reload' : for field setter : missing Type

// UNSUPPORTED : C value 'lookup_by_name' : for field getter : missing Type

// UNSUPPORTED : C value 'lookup_by_name' : for field setter : missing Type

// UNSUPPORTED : C value 'lookup_by_name_async' : for field getter : missing Type

// UNSUPPORTED : C value 'lookup_by_name_async' : for field setter : missing Type

// UNSUPPORTED : C value 'lookup_by_name_finish' : for field getter : missing Type

// UNSUPPORTED : C value 'lookup_by_name_finish' : for field setter : missing Type

// UNSUPPORTED : C value 'lookup_by_address' : for field getter : missing Type

// UNSUPPORTED : C value 'lookup_by_address' : for field setter : missing Type

// UNSUPPORTED : C value 'lookup_by_address_async' : for field getter : missing Type

// UNSUPPORTED : C value 'lookup_by_address_async' : for field setter : missing Type

// UNSUPPORTED : C value 'lookup_by_address_finish' : for field getter : missing Type

// UNSUPPORTED : C value 'lookup_by_address_finish' : for field setter : missing Type

// UNSUPPORTED : C value 'lookup_service' : for field getter : missing Type

// UNSUPPORTED : C value 'lookup_service' : for field setter : missing Type

// UNSUPPORTED : C value 'lookup_service_async' : for field getter : missing Type

// UNSUPPORTED : C value 'lookup_service_async' : for field setter : missing Type

// UNSUPPORTED : C value 'lookup_service_finish' : for field getter : missing Type

// UNSUPPORTED : C value 'lookup_service_finish' : for field setter : missing Type

// UNSUPPORTED : C value 'lookup_records' : for field getter : missing Type

// UNSUPPORTED : C value 'lookup_records' : for field setter : missing Type

// UNSUPPORTED : C value 'lookup_records_async' : for field getter : missing Type

// UNSUPPORTED : C value 'lookup_records_async' : for field setter : missing Type

// UNSUPPORTED : C value 'lookup_records_finish' : for field getter : missing Type

// UNSUPPORTED : C value 'lookup_records_finish' : for field setter : missing Type

// UNSUPPORTED : C value 'lookup_by_name_with_flags_async' : for field getter : missing Type

// UNSUPPORTED : C value 'lookup_by_name_with_flags_async' : for field setter : missing Type

// UNSUPPORTED : C value 'lookup_by_name_with_flags_finish' : for field getter : missing Type

// UNSUPPORTED : C value 'lookup_by_name_with_flags_finish' : for field setter : missing Type

// UNSUPPORTED : C value 'lookup_by_name_with_flags' : for field getter : missing Type

// UNSUPPORTED : C value 'lookup_by_name_with_flags' : for field setter : missing Type

// ResolverClassStruct creates an uninitialised ResolverClass.
func ResolverClassStruct() *ResolverClass {
	err := resolverClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := ResolverClassNewFromNative(resolverClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeResolverClass)
	return structGo
}
func finalizeResolverClass(obj *ResolverClass) {
	resolverClassStruct.Free(obj.native)
}

var resolverPrivateStruct *gi.Struct
var resolverPrivateStruct_Once sync.Once

func resolverPrivateStruct_Set() error {
	var err error
	resolverPrivateStruct_Once.Do(func() {
		resolverPrivateStruct, err = gi.StructNew("Gio", "ResolverPrivate")
	})
	return err
}

type ResolverPrivate struct {
	native unsafe.Pointer
}

func ResolverPrivateNewFromNative(native unsafe.Pointer) *ResolverPrivate {
	instance := &ResolverPrivate{native: native}

	return instance
}

/*
CastToResolverPrivate down casts any arbitrary Object to ResolverPrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a ResolverPrivate.
*/
func (recv *ResolverPrivate) CastToResolverPrivate(object *gobject.Object) *ResolverPrivate {
	return ResolverPrivateNewFromNative(object.Native())
}

func (recv *ResolverPrivate) Native() unsafe.Pointer {
	return recv.native
}

// ResolverPrivateStruct creates an uninitialised ResolverPrivate.
func ResolverPrivateStruct() *ResolverPrivate {
	err := resolverPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := ResolverPrivateNewFromNative(resolverPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeResolverPrivate)
	return structGo
}
func finalizeResolverPrivate(obj *ResolverPrivate) {
	resolverPrivateStruct.Free(obj.native)
}

var resourceStruct *gi.Struct
var resourceStruct_Once sync.Once

func resourceStruct_Set() error {
	var err error
	resourceStruct_Once.Do(func() {
		resourceStruct, err = gi.StructNew("Gio", "Resource")
	})
	return err
}

type Resource struct {
	native unsafe.Pointer
}

func ResourceNewFromNative(native unsafe.Pointer) *Resource {
	instance := &Resource{native: native}

	return instance
}

/*
CastToResource down casts any arbitrary Object to Resource.
Exercise care, as this is a potentially dangerous function
if the Object is not a Resource.
*/
func (recv *Resource) CastToResource(object *gobject.Object) *Resource {
	return ResourceNewFromNative(object.Native())
}

func (recv *Resource) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'g_resource_new_from_data' : parameter 'data' of type 'GLib.Bytes' not supported

var resourceRegisterFunction *gi.Function
var resourceRegisterFunction_Once sync.Once

func resourceRegisterFunction_Set() error {
	var err error
	resourceRegisterFunction_Once.Do(func() {
		err = resourceStruct_Set()
		if err != nil {
			return
		}
		resourceRegisterFunction, err = resourceStruct.InvokerNew("_register")
	})
	return err
}

// Register is a representation of the C type g_resources_register.
func (recv *Resource) Register() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := resourceRegisterFunction_Set()
	if err == nil {
		resourceRegisterFunction.Invoke(inArgs[:], nil)
	}

	return
}

var resourceUnregisterFunction *gi.Function
var resourceUnregisterFunction_Once sync.Once

func resourceUnregisterFunction_Set() error {
	var err error
	resourceUnregisterFunction_Once.Do(func() {
		err = resourceStruct_Set()
		if err != nil {
			return
		}
		resourceUnregisterFunction, err = resourceStruct.InvokerNew("_unregister")
	})
	return err
}

// Unregister is a representation of the C type g_resources_unregister.
func (recv *Resource) Unregister() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := resourceUnregisterFunction_Set()
	if err == nil {
		resourceUnregisterFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'g_resource_enumerate_children' : parameter 'lookup_flags' of type 'ResourceLookupFlags' not supported

// UNSUPPORTED : C value 'g_resource_get_info' : parameter 'lookup_flags' of type 'ResourceLookupFlags' not supported

// UNSUPPORTED : C value 'g_resource_lookup_data' : parameter 'lookup_flags' of type 'ResourceLookupFlags' not supported

// UNSUPPORTED : C value 'g_resource_open_stream' : parameter 'lookup_flags' of type 'ResourceLookupFlags' not supported

var resourceRefFunction *gi.Function
var resourceRefFunction_Once sync.Once

func resourceRefFunction_Set() error {
	var err error
	resourceRefFunction_Once.Do(func() {
		err = resourceStruct_Set()
		if err != nil {
			return
		}
		resourceRefFunction, err = resourceStruct.InvokerNew("ref")
	})
	return err
}

// Ref is a representation of the C type g_resource_ref.
func (recv *Resource) Ref() *Resource {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := resourceRefFunction_Set()
	if err == nil {
		ret = resourceRefFunction.Invoke(inArgs[:], nil)
	}

	retGo := ResourceNewFromNative(ret.Pointer())

	return retGo
}

var resourceUnrefFunction *gi.Function
var resourceUnrefFunction_Once sync.Once

func resourceUnrefFunction_Set() error {
	var err error
	resourceUnrefFunction_Once.Do(func() {
		err = resourceStruct_Set()
		if err != nil {
			return
		}
		resourceUnrefFunction, err = resourceStruct.InvokerNew("unref")
	})
	return err
}

// Unref is a representation of the C type g_resource_unref.
func (recv *Resource) Unref() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := resourceUnrefFunction_Set()
	if err == nil {
		resourceUnrefFunction.Invoke(inArgs[:], nil)
	}

	return
}

var seekableIfaceStruct *gi.Struct
var seekableIfaceStruct_Once sync.Once

func seekableIfaceStruct_Set() error {
	var err error
	seekableIfaceStruct_Once.Do(func() {
		seekableIfaceStruct, err = gi.StructNew("Gio", "SeekableIface")
	})
	return err
}

type SeekableIface struct {
	native unsafe.Pointer
}

func SeekableIfaceNewFromNative(native unsafe.Pointer) *SeekableIface {
	instance := &SeekableIface{native: native}

	return instance
}

/*
CastToSeekableIface down casts any arbitrary Object to SeekableIface.
Exercise care, as this is a potentially dangerous function
if the Object is not a SeekableIface.
*/
func (recv *SeekableIface) CastToSeekableIface(object *gobject.Object) *SeekableIface {
	return SeekableIfaceNewFromNative(object.Native())
}

func (recv *SeekableIface) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'g_iface' : for field getter : no Go type for 'GObject.TypeInterface'

// UNSUPPORTED : C value 'g_iface' : for field setter : no Go type for 'GObject.TypeInterface'

// UNSUPPORTED : C value 'tell' : for field getter : missing Type

// UNSUPPORTED : C value 'tell' : for field setter : missing Type

// UNSUPPORTED : C value 'can_seek' : for field getter : missing Type

// UNSUPPORTED : C value 'can_seek' : for field setter : missing Type

// UNSUPPORTED : C value 'seek' : for field getter : missing Type

// UNSUPPORTED : C value 'seek' : for field setter : missing Type

// UNSUPPORTED : C value 'can_truncate' : for field getter : missing Type

// UNSUPPORTED : C value 'can_truncate' : for field setter : missing Type

// UNSUPPORTED : C value 'truncate_fn' : for field getter : missing Type

// UNSUPPORTED : C value 'truncate_fn' : for field setter : missing Type

// SeekableIfaceStruct creates an uninitialised SeekableIface.
func SeekableIfaceStruct() *SeekableIface {
	err := seekableIfaceStruct_Set()
	if err != nil {
		return nil
	}

	structGo := SeekableIfaceNewFromNative(seekableIfaceStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeSeekableIface)
	return structGo
}
func finalizeSeekableIface(obj *SeekableIface) {
	seekableIfaceStruct.Free(obj.native)
}

var settingsBackendClassStruct *gi.Struct
var settingsBackendClassStruct_Once sync.Once

func settingsBackendClassStruct_Set() error {
	var err error
	settingsBackendClassStruct_Once.Do(func() {
		settingsBackendClassStruct, err = gi.StructNew("Gio", "SettingsBackendClass")
	})
	return err
}

type SettingsBackendClass struct {
	native unsafe.Pointer
}

func SettingsBackendClassNewFromNative(native unsafe.Pointer) *SettingsBackendClass {
	instance := &SettingsBackendClass{native: native}

	return instance
}

/*
CastToSettingsBackendClass down casts any arbitrary Object to SettingsBackendClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a SettingsBackendClass.
*/
func (recv *SettingsBackendClass) CastToSettingsBackendClass(object *gobject.Object) *SettingsBackendClass {
	return SettingsBackendClassNewFromNative(object.Native())
}

func (recv *SettingsBackendClass) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'read' : for field getter : missing Type

// UNSUPPORTED : C value 'read' : for field setter : missing Type

// UNSUPPORTED : C value 'get_writable' : for field getter : missing Type

// UNSUPPORTED : C value 'get_writable' : for field setter : missing Type

// UNSUPPORTED : C value 'write' : for field getter : missing Type

// UNSUPPORTED : C value 'write' : for field setter : missing Type

// UNSUPPORTED : C value 'write_tree' : for field getter : missing Type

// UNSUPPORTED : C value 'write_tree' : for field setter : missing Type

// UNSUPPORTED : C value 'reset' : for field getter : missing Type

// UNSUPPORTED : C value 'reset' : for field setter : missing Type

// UNSUPPORTED : C value 'subscribe' : for field getter : missing Type

// UNSUPPORTED : C value 'subscribe' : for field setter : missing Type

// UNSUPPORTED : C value 'unsubscribe' : for field getter : missing Type

// UNSUPPORTED : C value 'unsubscribe' : for field setter : missing Type

// UNSUPPORTED : C value 'sync' : for field getter : missing Type

// UNSUPPORTED : C value 'sync' : for field setter : missing Type

// UNSUPPORTED : C value 'get_permission' : for field getter : missing Type

// UNSUPPORTED : C value 'get_permission' : for field setter : missing Type

// UNSUPPORTED : C value 'read_user_value' : for field getter : missing Type

// UNSUPPORTED : C value 'read_user_value' : for field setter : missing Type

// SettingsBackendClassStruct creates an uninitialised SettingsBackendClass.
func SettingsBackendClassStruct() *SettingsBackendClass {
	err := settingsBackendClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := SettingsBackendClassNewFromNative(settingsBackendClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeSettingsBackendClass)
	return structGo
}
func finalizeSettingsBackendClass(obj *SettingsBackendClass) {
	settingsBackendClassStruct.Free(obj.native)
}

var settingsBackendPrivateStruct *gi.Struct
var settingsBackendPrivateStruct_Once sync.Once

func settingsBackendPrivateStruct_Set() error {
	var err error
	settingsBackendPrivateStruct_Once.Do(func() {
		settingsBackendPrivateStruct, err = gi.StructNew("Gio", "SettingsBackendPrivate")
	})
	return err
}

type SettingsBackendPrivate struct {
	native unsafe.Pointer
}

func SettingsBackendPrivateNewFromNative(native unsafe.Pointer) *SettingsBackendPrivate {
	instance := &SettingsBackendPrivate{native: native}

	return instance
}

/*
CastToSettingsBackendPrivate down casts any arbitrary Object to SettingsBackendPrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a SettingsBackendPrivate.
*/
func (recv *SettingsBackendPrivate) CastToSettingsBackendPrivate(object *gobject.Object) *SettingsBackendPrivate {
	return SettingsBackendPrivateNewFromNative(object.Native())
}

func (recv *SettingsBackendPrivate) Native() unsafe.Pointer {
	return recv.native
}

// SettingsBackendPrivateStruct creates an uninitialised SettingsBackendPrivate.
func SettingsBackendPrivateStruct() *SettingsBackendPrivate {
	err := settingsBackendPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := SettingsBackendPrivateNewFromNative(settingsBackendPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeSettingsBackendPrivate)
	return structGo
}
func finalizeSettingsBackendPrivate(obj *SettingsBackendPrivate) {
	settingsBackendPrivateStruct.Free(obj.native)
}

var settingsClassStruct *gi.Struct
var settingsClassStruct_Once sync.Once

func settingsClassStruct_Set() error {
	var err error
	settingsClassStruct_Once.Do(func() {
		settingsClassStruct, err = gi.StructNew("Gio", "SettingsClass")
	})
	return err
}

type SettingsClass struct {
	native unsafe.Pointer
}

func SettingsClassNewFromNative(native unsafe.Pointer) *SettingsClass {
	instance := &SettingsClass{native: native}

	return instance
}

/*
CastToSettingsClass down casts any arbitrary Object to SettingsClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a SettingsClass.
*/
func (recv *SettingsClass) CastToSettingsClass(object *gobject.Object) *SettingsClass {
	return SettingsClassNewFromNative(object.Native())
}

func (recv *SettingsClass) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'writable_changed' : for field getter : missing Type

// UNSUPPORTED : C value 'writable_changed' : for field setter : missing Type

// UNSUPPORTED : C value 'changed' : for field getter : missing Type

// UNSUPPORTED : C value 'changed' : for field setter : missing Type

// UNSUPPORTED : C value 'writable_change_event' : for field getter : missing Type

// UNSUPPORTED : C value 'writable_change_event' : for field setter : missing Type

// UNSUPPORTED : C value 'change_event' : for field getter : missing Type

// UNSUPPORTED : C value 'change_event' : for field setter : missing Type

// UNSUPPORTED : C value 'padding' : for field getter : missing Type

// UNSUPPORTED : C value 'padding' : for field setter : missing Type

// SettingsClassStruct creates an uninitialised SettingsClass.
func SettingsClassStruct() *SettingsClass {
	err := settingsClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := SettingsClassNewFromNative(settingsClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeSettingsClass)
	return structGo
}
func finalizeSettingsClass(obj *SettingsClass) {
	settingsClassStruct.Free(obj.native)
}

var settingsPrivateStruct *gi.Struct
var settingsPrivateStruct_Once sync.Once

func settingsPrivateStruct_Set() error {
	var err error
	settingsPrivateStruct_Once.Do(func() {
		settingsPrivateStruct, err = gi.StructNew("Gio", "SettingsPrivate")
	})
	return err
}

type SettingsPrivate struct {
	native unsafe.Pointer
}

func SettingsPrivateNewFromNative(native unsafe.Pointer) *SettingsPrivate {
	instance := &SettingsPrivate{native: native}

	return instance
}

/*
CastToSettingsPrivate down casts any arbitrary Object to SettingsPrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a SettingsPrivate.
*/
func (recv *SettingsPrivate) CastToSettingsPrivate(object *gobject.Object) *SettingsPrivate {
	return SettingsPrivateNewFromNative(object.Native())
}

func (recv *SettingsPrivate) Native() unsafe.Pointer {
	return recv.native
}

// SettingsPrivateStruct creates an uninitialised SettingsPrivate.
func SettingsPrivateStruct() *SettingsPrivate {
	err := settingsPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := SettingsPrivateNewFromNative(settingsPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeSettingsPrivate)
	return structGo
}
func finalizeSettingsPrivate(obj *SettingsPrivate) {
	settingsPrivateStruct.Free(obj.native)
}

var settingsSchemaStruct *gi.Struct
var settingsSchemaStruct_Once sync.Once

func settingsSchemaStruct_Set() error {
	var err error
	settingsSchemaStruct_Once.Do(func() {
		settingsSchemaStruct, err = gi.StructNew("Gio", "SettingsSchema")
	})
	return err
}

type SettingsSchema struct {
	native unsafe.Pointer
}

func SettingsSchemaNewFromNative(native unsafe.Pointer) *SettingsSchema {
	instance := &SettingsSchema{native: native}

	return instance
}

/*
CastToSettingsSchema down casts any arbitrary Object to SettingsSchema.
Exercise care, as this is a potentially dangerous function
if the Object is not a SettingsSchema.
*/
func (recv *SettingsSchema) CastToSettingsSchema(object *gobject.Object) *SettingsSchema {
	return SettingsSchemaNewFromNative(object.Native())
}

func (recv *SettingsSchema) Native() unsafe.Pointer {
	return recv.native
}

var settingsSchemaGetIdFunction *gi.Function
var settingsSchemaGetIdFunction_Once sync.Once

func settingsSchemaGetIdFunction_Set() error {
	var err error
	settingsSchemaGetIdFunction_Once.Do(func() {
		err = settingsSchemaStruct_Set()
		if err != nil {
			return
		}
		settingsSchemaGetIdFunction, err = settingsSchemaStruct.InvokerNew("get_id")
	})
	return err
}

// GetId is a representation of the C type g_settings_schema_get_id.
func (recv *SettingsSchema) GetId() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := settingsSchemaGetIdFunction_Set()
	if err == nil {
		ret = settingsSchemaGetIdFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var settingsSchemaGetKeyFunction *gi.Function
var settingsSchemaGetKeyFunction_Once sync.Once

func settingsSchemaGetKeyFunction_Set() error {
	var err error
	settingsSchemaGetKeyFunction_Once.Do(func() {
		err = settingsSchemaStruct_Set()
		if err != nil {
			return
		}
		settingsSchemaGetKeyFunction, err = settingsSchemaStruct.InvokerNew("get_key")
	})
	return err
}

// GetKey is a representation of the C type g_settings_schema_get_key.
func (recv *SettingsSchema) GetKey(name string) *SettingsSchemaKey {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(name)

	var ret gi.Argument

	err := settingsSchemaGetKeyFunction_Set()
	if err == nil {
		ret = settingsSchemaGetKeyFunction.Invoke(inArgs[:], nil)
	}

	retGo := SettingsSchemaKeyNewFromNative(ret.Pointer())

	return retGo
}

var settingsSchemaGetPathFunction *gi.Function
var settingsSchemaGetPathFunction_Once sync.Once

func settingsSchemaGetPathFunction_Set() error {
	var err error
	settingsSchemaGetPathFunction_Once.Do(func() {
		err = settingsSchemaStruct_Set()
		if err != nil {
			return
		}
		settingsSchemaGetPathFunction, err = settingsSchemaStruct.InvokerNew("get_path")
	})
	return err
}

// GetPath is a representation of the C type g_settings_schema_get_path.
func (recv *SettingsSchema) GetPath() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := settingsSchemaGetPathFunction_Set()
	if err == nil {
		ret = settingsSchemaGetPathFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var settingsSchemaHasKeyFunction *gi.Function
var settingsSchemaHasKeyFunction_Once sync.Once

func settingsSchemaHasKeyFunction_Set() error {
	var err error
	settingsSchemaHasKeyFunction_Once.Do(func() {
		err = settingsSchemaStruct_Set()
		if err != nil {
			return
		}
		settingsSchemaHasKeyFunction, err = settingsSchemaStruct.InvokerNew("has_key")
	})
	return err
}

// HasKey is a representation of the C type g_settings_schema_has_key.
func (recv *SettingsSchema) HasKey(name string) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(name)

	var ret gi.Argument

	err := settingsSchemaHasKeyFunction_Set()
	if err == nil {
		ret = settingsSchemaHasKeyFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var settingsSchemaListChildrenFunction *gi.Function
var settingsSchemaListChildrenFunction_Once sync.Once

func settingsSchemaListChildrenFunction_Set() error {
	var err error
	settingsSchemaListChildrenFunction_Once.Do(func() {
		err = settingsSchemaStruct_Set()
		if err != nil {
			return
		}
		settingsSchemaListChildrenFunction, err = settingsSchemaStruct.InvokerNew("list_children")
	})
	return err
}

// ListChildren is a representation of the C type g_settings_schema_list_children.
func (recv *SettingsSchema) ListChildren() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := settingsSchemaListChildrenFunction_Set()
	if err == nil {
		settingsSchemaListChildrenFunction.Invoke(inArgs[:], nil)
	}

	return
}

var settingsSchemaListKeysFunction *gi.Function
var settingsSchemaListKeysFunction_Once sync.Once

func settingsSchemaListKeysFunction_Set() error {
	var err error
	settingsSchemaListKeysFunction_Once.Do(func() {
		err = settingsSchemaStruct_Set()
		if err != nil {
			return
		}
		settingsSchemaListKeysFunction, err = settingsSchemaStruct.InvokerNew("list_keys")
	})
	return err
}

// ListKeys is a representation of the C type g_settings_schema_list_keys.
func (recv *SettingsSchema) ListKeys() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := settingsSchemaListKeysFunction_Set()
	if err == nil {
		settingsSchemaListKeysFunction.Invoke(inArgs[:], nil)
	}

	return
}

var settingsSchemaRefFunction *gi.Function
var settingsSchemaRefFunction_Once sync.Once

func settingsSchemaRefFunction_Set() error {
	var err error
	settingsSchemaRefFunction_Once.Do(func() {
		err = settingsSchemaStruct_Set()
		if err != nil {
			return
		}
		settingsSchemaRefFunction, err = settingsSchemaStruct.InvokerNew("ref")
	})
	return err
}

// Ref is a representation of the C type g_settings_schema_ref.
func (recv *SettingsSchema) Ref() *SettingsSchema {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := settingsSchemaRefFunction_Set()
	if err == nil {
		ret = settingsSchemaRefFunction.Invoke(inArgs[:], nil)
	}

	retGo := SettingsSchemaNewFromNative(ret.Pointer())

	return retGo
}

var settingsSchemaUnrefFunction *gi.Function
var settingsSchemaUnrefFunction_Once sync.Once

func settingsSchemaUnrefFunction_Set() error {
	var err error
	settingsSchemaUnrefFunction_Once.Do(func() {
		err = settingsSchemaStruct_Set()
		if err != nil {
			return
		}
		settingsSchemaUnrefFunction, err = settingsSchemaStruct.InvokerNew("unref")
	})
	return err
}

// Unref is a representation of the C type g_settings_schema_unref.
func (recv *SettingsSchema) Unref() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := settingsSchemaUnrefFunction_Set()
	if err == nil {
		settingsSchemaUnrefFunction.Invoke(inArgs[:], nil)
	}

	return
}

// SettingsSchemaStruct creates an uninitialised SettingsSchema.
func SettingsSchemaStruct() *SettingsSchema {
	err := settingsSchemaStruct_Set()
	if err != nil {
		return nil
	}

	structGo := SettingsSchemaNewFromNative(settingsSchemaStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeSettingsSchema)
	return structGo
}
func finalizeSettingsSchema(obj *SettingsSchema) {
	settingsSchemaStruct.Free(obj.native)
}

var settingsSchemaKeyStruct *gi.Struct
var settingsSchemaKeyStruct_Once sync.Once

func settingsSchemaKeyStruct_Set() error {
	var err error
	settingsSchemaKeyStruct_Once.Do(func() {
		settingsSchemaKeyStruct, err = gi.StructNew("Gio", "SettingsSchemaKey")
	})
	return err
}

type SettingsSchemaKey struct {
	native unsafe.Pointer
}

func SettingsSchemaKeyNewFromNative(native unsafe.Pointer) *SettingsSchemaKey {
	instance := &SettingsSchemaKey{native: native}

	return instance
}

/*
CastToSettingsSchemaKey down casts any arbitrary Object to SettingsSchemaKey.
Exercise care, as this is a potentially dangerous function
if the Object is not a SettingsSchemaKey.
*/
func (recv *SettingsSchemaKey) CastToSettingsSchemaKey(object *gobject.Object) *SettingsSchemaKey {
	return SettingsSchemaKeyNewFromNative(object.Native())
}

func (recv *SettingsSchemaKey) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'g_settings_schema_key_get_default_value' : return type 'GLib.Variant' not supported

var settingsSchemaKeyGetDescriptionFunction *gi.Function
var settingsSchemaKeyGetDescriptionFunction_Once sync.Once

func settingsSchemaKeyGetDescriptionFunction_Set() error {
	var err error
	settingsSchemaKeyGetDescriptionFunction_Once.Do(func() {
		err = settingsSchemaKeyStruct_Set()
		if err != nil {
			return
		}
		settingsSchemaKeyGetDescriptionFunction, err = settingsSchemaKeyStruct.InvokerNew("get_description")
	})
	return err
}

// GetDescription is a representation of the C type g_settings_schema_key_get_description.
func (recv *SettingsSchemaKey) GetDescription() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := settingsSchemaKeyGetDescriptionFunction_Set()
	if err == nil {
		ret = settingsSchemaKeyGetDescriptionFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var settingsSchemaKeyGetNameFunction *gi.Function
var settingsSchemaKeyGetNameFunction_Once sync.Once

func settingsSchemaKeyGetNameFunction_Set() error {
	var err error
	settingsSchemaKeyGetNameFunction_Once.Do(func() {
		err = settingsSchemaKeyStruct_Set()
		if err != nil {
			return
		}
		settingsSchemaKeyGetNameFunction, err = settingsSchemaKeyStruct.InvokerNew("get_name")
	})
	return err
}

// GetName is a representation of the C type g_settings_schema_key_get_name.
func (recv *SettingsSchemaKey) GetName() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := settingsSchemaKeyGetNameFunction_Set()
	if err == nil {
		ret = settingsSchemaKeyGetNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'g_settings_schema_key_get_range' : return type 'GLib.Variant' not supported

var settingsSchemaKeyGetSummaryFunction *gi.Function
var settingsSchemaKeyGetSummaryFunction_Once sync.Once

func settingsSchemaKeyGetSummaryFunction_Set() error {
	var err error
	settingsSchemaKeyGetSummaryFunction_Once.Do(func() {
		err = settingsSchemaKeyStruct_Set()
		if err != nil {
			return
		}
		settingsSchemaKeyGetSummaryFunction, err = settingsSchemaKeyStruct.InvokerNew("get_summary")
	})
	return err
}

// GetSummary is a representation of the C type g_settings_schema_key_get_summary.
func (recv *SettingsSchemaKey) GetSummary() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := settingsSchemaKeyGetSummaryFunction_Set()
	if err == nil {
		ret = settingsSchemaKeyGetSummaryFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'g_settings_schema_key_get_value_type' : return type 'GLib.VariantType' not supported

// UNSUPPORTED : C value 'g_settings_schema_key_range_check' : parameter 'value' of type 'GLib.Variant' not supported

var settingsSchemaKeyRefFunction *gi.Function
var settingsSchemaKeyRefFunction_Once sync.Once

func settingsSchemaKeyRefFunction_Set() error {
	var err error
	settingsSchemaKeyRefFunction_Once.Do(func() {
		err = settingsSchemaKeyStruct_Set()
		if err != nil {
			return
		}
		settingsSchemaKeyRefFunction, err = settingsSchemaKeyStruct.InvokerNew("ref")
	})
	return err
}

// Ref is a representation of the C type g_settings_schema_key_ref.
func (recv *SettingsSchemaKey) Ref() *SettingsSchemaKey {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := settingsSchemaKeyRefFunction_Set()
	if err == nil {
		ret = settingsSchemaKeyRefFunction.Invoke(inArgs[:], nil)
	}

	retGo := SettingsSchemaKeyNewFromNative(ret.Pointer())

	return retGo
}

var settingsSchemaKeyUnrefFunction *gi.Function
var settingsSchemaKeyUnrefFunction_Once sync.Once

func settingsSchemaKeyUnrefFunction_Set() error {
	var err error
	settingsSchemaKeyUnrefFunction_Once.Do(func() {
		err = settingsSchemaKeyStruct_Set()
		if err != nil {
			return
		}
		settingsSchemaKeyUnrefFunction, err = settingsSchemaKeyStruct.InvokerNew("unref")
	})
	return err
}

// Unref is a representation of the C type g_settings_schema_key_unref.
func (recv *SettingsSchemaKey) Unref() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := settingsSchemaKeyUnrefFunction_Set()
	if err == nil {
		settingsSchemaKeyUnrefFunction.Invoke(inArgs[:], nil)
	}

	return
}

// SettingsSchemaKeyStruct creates an uninitialised SettingsSchemaKey.
func SettingsSchemaKeyStruct() *SettingsSchemaKey {
	err := settingsSchemaKeyStruct_Set()
	if err != nil {
		return nil
	}

	structGo := SettingsSchemaKeyNewFromNative(settingsSchemaKeyStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeSettingsSchemaKey)
	return structGo
}
func finalizeSettingsSchemaKey(obj *SettingsSchemaKey) {
	settingsSchemaKeyStruct.Free(obj.native)
}

var settingsSchemaSourceStruct *gi.Struct
var settingsSchemaSourceStruct_Once sync.Once

func settingsSchemaSourceStruct_Set() error {
	var err error
	settingsSchemaSourceStruct_Once.Do(func() {
		settingsSchemaSourceStruct, err = gi.StructNew("Gio", "SettingsSchemaSource")
	})
	return err
}

type SettingsSchemaSource struct {
	native unsafe.Pointer
}

func SettingsSchemaSourceNewFromNative(native unsafe.Pointer) *SettingsSchemaSource {
	instance := &SettingsSchemaSource{native: native}

	return instance
}

/*
CastToSettingsSchemaSource down casts any arbitrary Object to SettingsSchemaSource.
Exercise care, as this is a potentially dangerous function
if the Object is not a SettingsSchemaSource.
*/
func (recv *SettingsSchemaSource) CastToSettingsSchemaSource(object *gobject.Object) *SettingsSchemaSource {
	return SettingsSchemaSourceNewFromNative(object.Native())
}

func (recv *SettingsSchemaSource) Native() unsafe.Pointer {
	return recv.native
}

var settingsSchemaSourceNewFromDirectoryFunction *gi.Function
var settingsSchemaSourceNewFromDirectoryFunction_Once sync.Once

func settingsSchemaSourceNewFromDirectoryFunction_Set() error {
	var err error
	settingsSchemaSourceNewFromDirectoryFunction_Once.Do(func() {
		err = settingsSchemaSourceStruct_Set()
		if err != nil {
			return
		}
		settingsSchemaSourceNewFromDirectoryFunction, err = settingsSchemaSourceStruct.InvokerNew("new_from_directory")
	})
	return err
}

// SettingsSchemaSourceNewFromDirectory is a representation of the C type g_settings_schema_source_new_from_directory.
func SettingsSchemaSourceNewFromDirectory(directory string, parent *SettingsSchemaSource, trusted bool) *SettingsSchemaSource {
	var inArgs [3]gi.Argument
	inArgs[0].SetString(directory)
	inArgs[1].SetPointer(parent.native)
	inArgs[2].SetBoolean(trusted)

	var ret gi.Argument

	err := settingsSchemaSourceNewFromDirectoryFunction_Set()
	if err == nil {
		ret = settingsSchemaSourceNewFromDirectoryFunction.Invoke(inArgs[:], nil)
	}

	retGo := SettingsSchemaSourceNewFromNative(ret.Pointer())

	return retGo
}

// UNSUPPORTED : C value 'g_settings_schema_source_list_schemas' : parameter 'non_relocatable' of type 'nil' not supported

var settingsSchemaSourceLookupFunction *gi.Function
var settingsSchemaSourceLookupFunction_Once sync.Once

func settingsSchemaSourceLookupFunction_Set() error {
	var err error
	settingsSchemaSourceLookupFunction_Once.Do(func() {
		err = settingsSchemaSourceStruct_Set()
		if err != nil {
			return
		}
		settingsSchemaSourceLookupFunction, err = settingsSchemaSourceStruct.InvokerNew("lookup")
	})
	return err
}

// Lookup is a representation of the C type g_settings_schema_source_lookup.
func (recv *SettingsSchemaSource) Lookup(schemaId string, recursive bool) *SettingsSchema {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(schemaId)
	inArgs[2].SetBoolean(recursive)

	var ret gi.Argument

	err := settingsSchemaSourceLookupFunction_Set()
	if err == nil {
		ret = settingsSchemaSourceLookupFunction.Invoke(inArgs[:], nil)
	}

	retGo := SettingsSchemaNewFromNative(ret.Pointer())

	return retGo
}

var settingsSchemaSourceRefFunction *gi.Function
var settingsSchemaSourceRefFunction_Once sync.Once

func settingsSchemaSourceRefFunction_Set() error {
	var err error
	settingsSchemaSourceRefFunction_Once.Do(func() {
		err = settingsSchemaSourceStruct_Set()
		if err != nil {
			return
		}
		settingsSchemaSourceRefFunction, err = settingsSchemaSourceStruct.InvokerNew("ref")
	})
	return err
}

// Ref is a representation of the C type g_settings_schema_source_ref.
func (recv *SettingsSchemaSource) Ref() *SettingsSchemaSource {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := settingsSchemaSourceRefFunction_Set()
	if err == nil {
		ret = settingsSchemaSourceRefFunction.Invoke(inArgs[:], nil)
	}

	retGo := SettingsSchemaSourceNewFromNative(ret.Pointer())

	return retGo
}

var settingsSchemaSourceUnrefFunction *gi.Function
var settingsSchemaSourceUnrefFunction_Once sync.Once

func settingsSchemaSourceUnrefFunction_Set() error {
	var err error
	settingsSchemaSourceUnrefFunction_Once.Do(func() {
		err = settingsSchemaSourceStruct_Set()
		if err != nil {
			return
		}
		settingsSchemaSourceUnrefFunction, err = settingsSchemaSourceStruct.InvokerNew("unref")
	})
	return err
}

// Unref is a representation of the C type g_settings_schema_source_unref.
func (recv *SettingsSchemaSource) Unref() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := settingsSchemaSourceUnrefFunction_Set()
	if err == nil {
		settingsSchemaSourceUnrefFunction.Invoke(inArgs[:], nil)
	}

	return
}

var simpleActionGroupClassStruct *gi.Struct
var simpleActionGroupClassStruct_Once sync.Once

func simpleActionGroupClassStruct_Set() error {
	var err error
	simpleActionGroupClassStruct_Once.Do(func() {
		simpleActionGroupClassStruct, err = gi.StructNew("Gio", "SimpleActionGroupClass")
	})
	return err
}

type SimpleActionGroupClass struct {
	native unsafe.Pointer
}

func SimpleActionGroupClassNewFromNative(native unsafe.Pointer) *SimpleActionGroupClass {
	instance := &SimpleActionGroupClass{native: native}

	return instance
}

/*
CastToSimpleActionGroupClass down casts any arbitrary Object to SimpleActionGroupClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a SimpleActionGroupClass.
*/
func (recv *SimpleActionGroupClass) CastToSimpleActionGroupClass(object *gobject.Object) *SimpleActionGroupClass {
	return SimpleActionGroupClassNewFromNative(object.Native())
}

func (recv *SimpleActionGroupClass) Native() unsafe.Pointer {
	return recv.native
}

// SimpleActionGroupClassStruct creates an uninitialised SimpleActionGroupClass.
func SimpleActionGroupClassStruct() *SimpleActionGroupClass {
	err := simpleActionGroupClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := SimpleActionGroupClassNewFromNative(simpleActionGroupClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeSimpleActionGroupClass)
	return structGo
}
func finalizeSimpleActionGroupClass(obj *SimpleActionGroupClass) {
	simpleActionGroupClassStruct.Free(obj.native)
}

var simpleActionGroupPrivateStruct *gi.Struct
var simpleActionGroupPrivateStruct_Once sync.Once

func simpleActionGroupPrivateStruct_Set() error {
	var err error
	simpleActionGroupPrivateStruct_Once.Do(func() {
		simpleActionGroupPrivateStruct, err = gi.StructNew("Gio", "SimpleActionGroupPrivate")
	})
	return err
}

type SimpleActionGroupPrivate struct {
	native unsafe.Pointer
}

func SimpleActionGroupPrivateNewFromNative(native unsafe.Pointer) *SimpleActionGroupPrivate {
	instance := &SimpleActionGroupPrivate{native: native}

	return instance
}

/*
CastToSimpleActionGroupPrivate down casts any arbitrary Object to SimpleActionGroupPrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a SimpleActionGroupPrivate.
*/
func (recv *SimpleActionGroupPrivate) CastToSimpleActionGroupPrivate(object *gobject.Object) *SimpleActionGroupPrivate {
	return SimpleActionGroupPrivateNewFromNative(object.Native())
}

func (recv *SimpleActionGroupPrivate) Native() unsafe.Pointer {
	return recv.native
}

// SimpleActionGroupPrivateStruct creates an uninitialised SimpleActionGroupPrivate.
func SimpleActionGroupPrivateStruct() *SimpleActionGroupPrivate {
	err := simpleActionGroupPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := SimpleActionGroupPrivateNewFromNative(simpleActionGroupPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeSimpleActionGroupPrivate)
	return structGo
}
func finalizeSimpleActionGroupPrivate(obj *SimpleActionGroupPrivate) {
	simpleActionGroupPrivateStruct.Free(obj.native)
}

var simpleAsyncResultClassStruct *gi.Struct
var simpleAsyncResultClassStruct_Once sync.Once

func simpleAsyncResultClassStruct_Set() error {
	var err error
	simpleAsyncResultClassStruct_Once.Do(func() {
		simpleAsyncResultClassStruct, err = gi.StructNew("Gio", "SimpleAsyncResultClass")
	})
	return err
}

type SimpleAsyncResultClass struct {
	native unsafe.Pointer
}

func SimpleAsyncResultClassNewFromNative(native unsafe.Pointer) *SimpleAsyncResultClass {
	instance := &SimpleAsyncResultClass{native: native}

	return instance
}

/*
CastToSimpleAsyncResultClass down casts any arbitrary Object to SimpleAsyncResultClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a SimpleAsyncResultClass.
*/
func (recv *SimpleAsyncResultClass) CastToSimpleAsyncResultClass(object *gobject.Object) *SimpleAsyncResultClass {
	return SimpleAsyncResultClassNewFromNative(object.Native())
}

func (recv *SimpleAsyncResultClass) Native() unsafe.Pointer {
	return recv.native
}

// SimpleAsyncResultClassStruct creates an uninitialised SimpleAsyncResultClass.
func SimpleAsyncResultClassStruct() *SimpleAsyncResultClass {
	err := simpleAsyncResultClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := SimpleAsyncResultClassNewFromNative(simpleAsyncResultClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeSimpleAsyncResultClass)
	return structGo
}
func finalizeSimpleAsyncResultClass(obj *SimpleAsyncResultClass) {
	simpleAsyncResultClassStruct.Free(obj.native)
}

var simpleProxyResolverClassStruct *gi.Struct
var simpleProxyResolverClassStruct_Once sync.Once

func simpleProxyResolverClassStruct_Set() error {
	var err error
	simpleProxyResolverClassStruct_Once.Do(func() {
		simpleProxyResolverClassStruct, err = gi.StructNew("Gio", "SimpleProxyResolverClass")
	})
	return err
}

type SimpleProxyResolverClass struct {
	native unsafe.Pointer
}

func SimpleProxyResolverClassNewFromNative(native unsafe.Pointer) *SimpleProxyResolverClass {
	instance := &SimpleProxyResolverClass{native: native}

	return instance
}

/*
CastToSimpleProxyResolverClass down casts any arbitrary Object to SimpleProxyResolverClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a SimpleProxyResolverClass.
*/
func (recv *SimpleProxyResolverClass) CastToSimpleProxyResolverClass(object *gobject.Object) *SimpleProxyResolverClass {
	return SimpleProxyResolverClassNewFromNative(object.Native())
}

func (recv *SimpleProxyResolverClass) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value '_g_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved4' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved5' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved5' : for field setter : missing Type

// SimpleProxyResolverClassStruct creates an uninitialised SimpleProxyResolverClass.
func SimpleProxyResolverClassStruct() *SimpleProxyResolverClass {
	err := simpleProxyResolverClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := SimpleProxyResolverClassNewFromNative(simpleProxyResolverClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeSimpleProxyResolverClass)
	return structGo
}
func finalizeSimpleProxyResolverClass(obj *SimpleProxyResolverClass) {
	simpleProxyResolverClassStruct.Free(obj.native)
}

var simpleProxyResolverPrivateStruct *gi.Struct
var simpleProxyResolverPrivateStruct_Once sync.Once

func simpleProxyResolverPrivateStruct_Set() error {
	var err error
	simpleProxyResolverPrivateStruct_Once.Do(func() {
		simpleProxyResolverPrivateStruct, err = gi.StructNew("Gio", "SimpleProxyResolverPrivate")
	})
	return err
}

type SimpleProxyResolverPrivate struct {
	native unsafe.Pointer
}

func SimpleProxyResolverPrivateNewFromNative(native unsafe.Pointer) *SimpleProxyResolverPrivate {
	instance := &SimpleProxyResolverPrivate{native: native}

	return instance
}

/*
CastToSimpleProxyResolverPrivate down casts any arbitrary Object to SimpleProxyResolverPrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a SimpleProxyResolverPrivate.
*/
func (recv *SimpleProxyResolverPrivate) CastToSimpleProxyResolverPrivate(object *gobject.Object) *SimpleProxyResolverPrivate {
	return SimpleProxyResolverPrivateNewFromNative(object.Native())
}

func (recv *SimpleProxyResolverPrivate) Native() unsafe.Pointer {
	return recv.native
}

// SimpleProxyResolverPrivateStruct creates an uninitialised SimpleProxyResolverPrivate.
func SimpleProxyResolverPrivateStruct() *SimpleProxyResolverPrivate {
	err := simpleProxyResolverPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := SimpleProxyResolverPrivateNewFromNative(simpleProxyResolverPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeSimpleProxyResolverPrivate)
	return structGo
}
func finalizeSimpleProxyResolverPrivate(obj *SimpleProxyResolverPrivate) {
	simpleProxyResolverPrivateStruct.Free(obj.native)
}

var socketAddressClassStruct *gi.Struct
var socketAddressClassStruct_Once sync.Once

func socketAddressClassStruct_Set() error {
	var err error
	socketAddressClassStruct_Once.Do(func() {
		socketAddressClassStruct, err = gi.StructNew("Gio", "SocketAddressClass")
	})
	return err
}

type SocketAddressClass struct {
	native unsafe.Pointer
}

func SocketAddressClassNewFromNative(native unsafe.Pointer) *SocketAddressClass {
	instance := &SocketAddressClass{native: native}

	return instance
}

/*
CastToSocketAddressClass down casts any arbitrary Object to SocketAddressClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a SocketAddressClass.
*/
func (recv *SocketAddressClass) CastToSocketAddressClass(object *gobject.Object) *SocketAddressClass {
	return SocketAddressClassNewFromNative(object.Native())
}

func (recv *SocketAddressClass) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'get_family' : for field getter : missing Type

// UNSUPPORTED : C value 'get_family' : for field setter : missing Type

// UNSUPPORTED : C value 'get_native_size' : for field getter : missing Type

// UNSUPPORTED : C value 'get_native_size' : for field setter : missing Type

// UNSUPPORTED : C value 'to_native' : for field getter : missing Type

// UNSUPPORTED : C value 'to_native' : for field setter : missing Type

// SocketAddressClassStruct creates an uninitialised SocketAddressClass.
func SocketAddressClassStruct() *SocketAddressClass {
	err := socketAddressClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := SocketAddressClassNewFromNative(socketAddressClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeSocketAddressClass)
	return structGo
}
func finalizeSocketAddressClass(obj *SocketAddressClass) {
	socketAddressClassStruct.Free(obj.native)
}

var socketAddressEnumeratorClassStruct *gi.Struct
var socketAddressEnumeratorClassStruct_Once sync.Once

func socketAddressEnumeratorClassStruct_Set() error {
	var err error
	socketAddressEnumeratorClassStruct_Once.Do(func() {
		socketAddressEnumeratorClassStruct, err = gi.StructNew("Gio", "SocketAddressEnumeratorClass")
	})
	return err
}

type SocketAddressEnumeratorClass struct {
	native unsafe.Pointer
}

func SocketAddressEnumeratorClassNewFromNative(native unsafe.Pointer) *SocketAddressEnumeratorClass {
	instance := &SocketAddressEnumeratorClass{native: native}

	return instance
}

/*
CastToSocketAddressEnumeratorClass down casts any arbitrary Object to SocketAddressEnumeratorClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a SocketAddressEnumeratorClass.
*/
func (recv *SocketAddressEnumeratorClass) CastToSocketAddressEnumeratorClass(object *gobject.Object) *SocketAddressEnumeratorClass {
	return SocketAddressEnumeratorClassNewFromNative(object.Native())
}

func (recv *SocketAddressEnumeratorClass) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'next' : for field getter : missing Type

// UNSUPPORTED : C value 'next' : for field setter : missing Type

// UNSUPPORTED : C value 'next_async' : for field getter : missing Type

// UNSUPPORTED : C value 'next_async' : for field setter : missing Type

// UNSUPPORTED : C value 'next_finish' : for field getter : missing Type

// UNSUPPORTED : C value 'next_finish' : for field setter : missing Type

// SocketAddressEnumeratorClassStruct creates an uninitialised SocketAddressEnumeratorClass.
func SocketAddressEnumeratorClassStruct() *SocketAddressEnumeratorClass {
	err := socketAddressEnumeratorClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := SocketAddressEnumeratorClassNewFromNative(socketAddressEnumeratorClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeSocketAddressEnumeratorClass)
	return structGo
}
func finalizeSocketAddressEnumeratorClass(obj *SocketAddressEnumeratorClass) {
	socketAddressEnumeratorClassStruct.Free(obj.native)
}

var socketClassStruct *gi.Struct
var socketClassStruct_Once sync.Once

func socketClassStruct_Set() error {
	var err error
	socketClassStruct_Once.Do(func() {
		socketClassStruct, err = gi.StructNew("Gio", "SocketClass")
	})
	return err
}

type SocketClass struct {
	native unsafe.Pointer
}

func SocketClassNewFromNative(native unsafe.Pointer) *SocketClass {
	instance := &SocketClass{native: native}

	return instance
}

/*
CastToSocketClass down casts any arbitrary Object to SocketClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a SocketClass.
*/
func (recv *SocketClass) CastToSocketClass(object *gobject.Object) *SocketClass {
	return SocketClassNewFromNative(object.Native())
}

func (recv *SocketClass) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value '_g_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved4' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved5' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved5' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved6' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved6' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved7' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved7' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved8' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved8' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved9' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved9' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved10' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved10' : for field setter : missing Type

// SocketClassStruct creates an uninitialised SocketClass.
func SocketClassStruct() *SocketClass {
	err := socketClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := SocketClassNewFromNative(socketClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeSocketClass)
	return structGo
}
func finalizeSocketClass(obj *SocketClass) {
	socketClassStruct.Free(obj.native)
}

var socketClientClassStruct *gi.Struct
var socketClientClassStruct_Once sync.Once

func socketClientClassStruct_Set() error {
	var err error
	socketClientClassStruct_Once.Do(func() {
		socketClientClassStruct, err = gi.StructNew("Gio", "SocketClientClass")
	})
	return err
}

type SocketClientClass struct {
	native unsafe.Pointer
}

func SocketClientClassNewFromNative(native unsafe.Pointer) *SocketClientClass {
	instance := &SocketClientClass{native: native}

	return instance
}

/*
CastToSocketClientClass down casts any arbitrary Object to SocketClientClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a SocketClientClass.
*/
func (recv *SocketClientClass) CastToSocketClientClass(object *gobject.Object) *SocketClientClass {
	return SocketClientClassNewFromNative(object.Native())
}

func (recv *SocketClientClass) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'event' : for field getter : missing Type

// UNSUPPORTED : C value 'event' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved4' : for field setter : missing Type

// SocketClientClassStruct creates an uninitialised SocketClientClass.
func SocketClientClassStruct() *SocketClientClass {
	err := socketClientClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := SocketClientClassNewFromNative(socketClientClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeSocketClientClass)
	return structGo
}
func finalizeSocketClientClass(obj *SocketClientClass) {
	socketClientClassStruct.Free(obj.native)
}

var socketClientPrivateStruct *gi.Struct
var socketClientPrivateStruct_Once sync.Once

func socketClientPrivateStruct_Set() error {
	var err error
	socketClientPrivateStruct_Once.Do(func() {
		socketClientPrivateStruct, err = gi.StructNew("Gio", "SocketClientPrivate")
	})
	return err
}

type SocketClientPrivate struct {
	native unsafe.Pointer
}

func SocketClientPrivateNewFromNative(native unsafe.Pointer) *SocketClientPrivate {
	instance := &SocketClientPrivate{native: native}

	return instance
}

/*
CastToSocketClientPrivate down casts any arbitrary Object to SocketClientPrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a SocketClientPrivate.
*/
func (recv *SocketClientPrivate) CastToSocketClientPrivate(object *gobject.Object) *SocketClientPrivate {
	return SocketClientPrivateNewFromNative(object.Native())
}

func (recv *SocketClientPrivate) Native() unsafe.Pointer {
	return recv.native
}

// SocketClientPrivateStruct creates an uninitialised SocketClientPrivate.
func SocketClientPrivateStruct() *SocketClientPrivate {
	err := socketClientPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := SocketClientPrivateNewFromNative(socketClientPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeSocketClientPrivate)
	return structGo
}
func finalizeSocketClientPrivate(obj *SocketClientPrivate) {
	socketClientPrivateStruct.Free(obj.native)
}

var socketConnectableIfaceStruct *gi.Struct
var socketConnectableIfaceStruct_Once sync.Once

func socketConnectableIfaceStruct_Set() error {
	var err error
	socketConnectableIfaceStruct_Once.Do(func() {
		socketConnectableIfaceStruct, err = gi.StructNew("Gio", "SocketConnectableIface")
	})
	return err
}

type SocketConnectableIface struct {
	native unsafe.Pointer
}

func SocketConnectableIfaceNewFromNative(native unsafe.Pointer) *SocketConnectableIface {
	instance := &SocketConnectableIface{native: native}

	return instance
}

/*
CastToSocketConnectableIface down casts any arbitrary Object to SocketConnectableIface.
Exercise care, as this is a potentially dangerous function
if the Object is not a SocketConnectableIface.
*/
func (recv *SocketConnectableIface) CastToSocketConnectableIface(object *gobject.Object) *SocketConnectableIface {
	return SocketConnectableIfaceNewFromNative(object.Native())
}

func (recv *SocketConnectableIface) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'g_iface' : for field getter : no Go type for 'GObject.TypeInterface'

// UNSUPPORTED : C value 'g_iface' : for field setter : no Go type for 'GObject.TypeInterface'

// UNSUPPORTED : C value 'enumerate' : for field getter : missing Type

// UNSUPPORTED : C value 'enumerate' : for field setter : missing Type

// UNSUPPORTED : C value 'proxy_enumerate' : for field getter : missing Type

// UNSUPPORTED : C value 'proxy_enumerate' : for field setter : missing Type

// UNSUPPORTED : C value 'to_string' : for field getter : missing Type

// UNSUPPORTED : C value 'to_string' : for field setter : missing Type

// SocketConnectableIfaceStruct creates an uninitialised SocketConnectableIface.
func SocketConnectableIfaceStruct() *SocketConnectableIface {
	err := socketConnectableIfaceStruct_Set()
	if err != nil {
		return nil
	}

	structGo := SocketConnectableIfaceNewFromNative(socketConnectableIfaceStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeSocketConnectableIface)
	return structGo
}
func finalizeSocketConnectableIface(obj *SocketConnectableIface) {
	socketConnectableIfaceStruct.Free(obj.native)
}

var socketConnectionClassStruct *gi.Struct
var socketConnectionClassStruct_Once sync.Once

func socketConnectionClassStruct_Set() error {
	var err error
	socketConnectionClassStruct_Once.Do(func() {
		socketConnectionClassStruct, err = gi.StructNew("Gio", "SocketConnectionClass")
	})
	return err
}

type SocketConnectionClass struct {
	native unsafe.Pointer
}

func SocketConnectionClassNewFromNative(native unsafe.Pointer) *SocketConnectionClass {
	instance := &SocketConnectionClass{native: native}

	return instance
}

/*
CastToSocketConnectionClass down casts any arbitrary Object to SocketConnectionClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a SocketConnectionClass.
*/
func (recv *SocketConnectionClass) CastToSocketConnectionClass(object *gobject.Object) *SocketConnectionClass {
	return SocketConnectionClassNewFromNative(object.Native())
}

func (recv *SocketConnectionClass) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentClass returns the C field 'parent_class'.
func (recv *SocketConnectionClass) FieldParentClass() *IOStreamClass {
	argValue := gi.StructFieldGet(socketConnectionClassStruct, recv.native, "parent_class")
	value := IOStreamClassNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *SocketConnectionClass) SetFieldParentClass(value *IOStreamClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.StructFieldSet(socketConnectionClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value '_g_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved4' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved5' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved5' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved6' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved6' : for field setter : missing Type

// SocketConnectionClassStruct creates an uninitialised SocketConnectionClass.
func SocketConnectionClassStruct() *SocketConnectionClass {
	err := socketConnectionClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := SocketConnectionClassNewFromNative(socketConnectionClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeSocketConnectionClass)
	return structGo
}
func finalizeSocketConnectionClass(obj *SocketConnectionClass) {
	socketConnectionClassStruct.Free(obj.native)
}

var socketConnectionPrivateStruct *gi.Struct
var socketConnectionPrivateStruct_Once sync.Once

func socketConnectionPrivateStruct_Set() error {
	var err error
	socketConnectionPrivateStruct_Once.Do(func() {
		socketConnectionPrivateStruct, err = gi.StructNew("Gio", "SocketConnectionPrivate")
	})
	return err
}

type SocketConnectionPrivate struct {
	native unsafe.Pointer
}

func SocketConnectionPrivateNewFromNative(native unsafe.Pointer) *SocketConnectionPrivate {
	instance := &SocketConnectionPrivate{native: native}

	return instance
}

/*
CastToSocketConnectionPrivate down casts any arbitrary Object to SocketConnectionPrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a SocketConnectionPrivate.
*/
func (recv *SocketConnectionPrivate) CastToSocketConnectionPrivate(object *gobject.Object) *SocketConnectionPrivate {
	return SocketConnectionPrivateNewFromNative(object.Native())
}

func (recv *SocketConnectionPrivate) Native() unsafe.Pointer {
	return recv.native
}

// SocketConnectionPrivateStruct creates an uninitialised SocketConnectionPrivate.
func SocketConnectionPrivateStruct() *SocketConnectionPrivate {
	err := socketConnectionPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := SocketConnectionPrivateNewFromNative(socketConnectionPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeSocketConnectionPrivate)
	return structGo
}
func finalizeSocketConnectionPrivate(obj *SocketConnectionPrivate) {
	socketConnectionPrivateStruct.Free(obj.native)
}

var socketControlMessageClassStruct *gi.Struct
var socketControlMessageClassStruct_Once sync.Once

func socketControlMessageClassStruct_Set() error {
	var err error
	socketControlMessageClassStruct_Once.Do(func() {
		socketControlMessageClassStruct, err = gi.StructNew("Gio", "SocketControlMessageClass")
	})
	return err
}

type SocketControlMessageClass struct {
	native unsafe.Pointer
}

func SocketControlMessageClassNewFromNative(native unsafe.Pointer) *SocketControlMessageClass {
	instance := &SocketControlMessageClass{native: native}

	return instance
}

/*
CastToSocketControlMessageClass down casts any arbitrary Object to SocketControlMessageClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a SocketControlMessageClass.
*/
func (recv *SocketControlMessageClass) CastToSocketControlMessageClass(object *gobject.Object) *SocketControlMessageClass {
	return SocketControlMessageClassNewFromNative(object.Native())
}

func (recv *SocketControlMessageClass) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'get_size' : for field getter : missing Type

// UNSUPPORTED : C value 'get_size' : for field setter : missing Type

// UNSUPPORTED : C value 'get_level' : for field getter : missing Type

// UNSUPPORTED : C value 'get_level' : for field setter : missing Type

// UNSUPPORTED : C value 'get_type' : for field getter : missing Type

// UNSUPPORTED : C value 'get_type' : for field setter : missing Type

// UNSUPPORTED : C value 'serialize' : for field getter : missing Type

// UNSUPPORTED : C value 'serialize' : for field setter : missing Type

// UNSUPPORTED : C value 'deserialize' : for field getter : missing Type

// UNSUPPORTED : C value 'deserialize' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved4' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved5' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved5' : for field setter : missing Type

// SocketControlMessageClassStruct creates an uninitialised SocketControlMessageClass.
func SocketControlMessageClassStruct() *SocketControlMessageClass {
	err := socketControlMessageClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := SocketControlMessageClassNewFromNative(socketControlMessageClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeSocketControlMessageClass)
	return structGo
}
func finalizeSocketControlMessageClass(obj *SocketControlMessageClass) {
	socketControlMessageClassStruct.Free(obj.native)
}

var socketControlMessagePrivateStruct *gi.Struct
var socketControlMessagePrivateStruct_Once sync.Once

func socketControlMessagePrivateStruct_Set() error {
	var err error
	socketControlMessagePrivateStruct_Once.Do(func() {
		socketControlMessagePrivateStruct, err = gi.StructNew("Gio", "SocketControlMessagePrivate")
	})
	return err
}

type SocketControlMessagePrivate struct {
	native unsafe.Pointer
}

func SocketControlMessagePrivateNewFromNative(native unsafe.Pointer) *SocketControlMessagePrivate {
	instance := &SocketControlMessagePrivate{native: native}

	return instance
}

/*
CastToSocketControlMessagePrivate down casts any arbitrary Object to SocketControlMessagePrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a SocketControlMessagePrivate.
*/
func (recv *SocketControlMessagePrivate) CastToSocketControlMessagePrivate(object *gobject.Object) *SocketControlMessagePrivate {
	return SocketControlMessagePrivateNewFromNative(object.Native())
}

func (recv *SocketControlMessagePrivate) Native() unsafe.Pointer {
	return recv.native
}

// SocketControlMessagePrivateStruct creates an uninitialised SocketControlMessagePrivate.
func SocketControlMessagePrivateStruct() *SocketControlMessagePrivate {
	err := socketControlMessagePrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := SocketControlMessagePrivateNewFromNative(socketControlMessagePrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeSocketControlMessagePrivate)
	return structGo
}
func finalizeSocketControlMessagePrivate(obj *SocketControlMessagePrivate) {
	socketControlMessagePrivateStruct.Free(obj.native)
}

var socketListenerClassStruct *gi.Struct
var socketListenerClassStruct_Once sync.Once

func socketListenerClassStruct_Set() error {
	var err error
	socketListenerClassStruct_Once.Do(func() {
		socketListenerClassStruct, err = gi.StructNew("Gio", "SocketListenerClass")
	})
	return err
}

type SocketListenerClass struct {
	native unsafe.Pointer
}

func SocketListenerClassNewFromNative(native unsafe.Pointer) *SocketListenerClass {
	instance := &SocketListenerClass{native: native}

	return instance
}

/*
CastToSocketListenerClass down casts any arbitrary Object to SocketListenerClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a SocketListenerClass.
*/
func (recv *SocketListenerClass) CastToSocketListenerClass(object *gobject.Object) *SocketListenerClass {
	return SocketListenerClassNewFromNative(object.Native())
}

func (recv *SocketListenerClass) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'changed' : for field getter : missing Type

// UNSUPPORTED : C value 'changed' : for field setter : missing Type

// UNSUPPORTED : C value 'event' : for field getter : missing Type

// UNSUPPORTED : C value 'event' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved4' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved5' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved5' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved6' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved6' : for field setter : missing Type

// SocketListenerClassStruct creates an uninitialised SocketListenerClass.
func SocketListenerClassStruct() *SocketListenerClass {
	err := socketListenerClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := SocketListenerClassNewFromNative(socketListenerClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeSocketListenerClass)
	return structGo
}
func finalizeSocketListenerClass(obj *SocketListenerClass) {
	socketListenerClassStruct.Free(obj.native)
}

var socketListenerPrivateStruct *gi.Struct
var socketListenerPrivateStruct_Once sync.Once

func socketListenerPrivateStruct_Set() error {
	var err error
	socketListenerPrivateStruct_Once.Do(func() {
		socketListenerPrivateStruct, err = gi.StructNew("Gio", "SocketListenerPrivate")
	})
	return err
}

type SocketListenerPrivate struct {
	native unsafe.Pointer
}

func SocketListenerPrivateNewFromNative(native unsafe.Pointer) *SocketListenerPrivate {
	instance := &SocketListenerPrivate{native: native}

	return instance
}

/*
CastToSocketListenerPrivate down casts any arbitrary Object to SocketListenerPrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a SocketListenerPrivate.
*/
func (recv *SocketListenerPrivate) CastToSocketListenerPrivate(object *gobject.Object) *SocketListenerPrivate {
	return SocketListenerPrivateNewFromNative(object.Native())
}

func (recv *SocketListenerPrivate) Native() unsafe.Pointer {
	return recv.native
}

// SocketListenerPrivateStruct creates an uninitialised SocketListenerPrivate.
func SocketListenerPrivateStruct() *SocketListenerPrivate {
	err := socketListenerPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := SocketListenerPrivateNewFromNative(socketListenerPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeSocketListenerPrivate)
	return structGo
}
func finalizeSocketListenerPrivate(obj *SocketListenerPrivate) {
	socketListenerPrivateStruct.Free(obj.native)
}

var socketPrivateStruct *gi.Struct
var socketPrivateStruct_Once sync.Once

func socketPrivateStruct_Set() error {
	var err error
	socketPrivateStruct_Once.Do(func() {
		socketPrivateStruct, err = gi.StructNew("Gio", "SocketPrivate")
	})
	return err
}

type SocketPrivate struct {
	native unsafe.Pointer
}

func SocketPrivateNewFromNative(native unsafe.Pointer) *SocketPrivate {
	instance := &SocketPrivate{native: native}

	return instance
}

/*
CastToSocketPrivate down casts any arbitrary Object to SocketPrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a SocketPrivate.
*/
func (recv *SocketPrivate) CastToSocketPrivate(object *gobject.Object) *SocketPrivate {
	return SocketPrivateNewFromNative(object.Native())
}

func (recv *SocketPrivate) Native() unsafe.Pointer {
	return recv.native
}

// SocketPrivateStruct creates an uninitialised SocketPrivate.
func SocketPrivateStruct() *SocketPrivate {
	err := socketPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := SocketPrivateNewFromNative(socketPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeSocketPrivate)
	return structGo
}
func finalizeSocketPrivate(obj *SocketPrivate) {
	socketPrivateStruct.Free(obj.native)
}

var socketServiceClassStruct *gi.Struct
var socketServiceClassStruct_Once sync.Once

func socketServiceClassStruct_Set() error {
	var err error
	socketServiceClassStruct_Once.Do(func() {
		socketServiceClassStruct, err = gi.StructNew("Gio", "SocketServiceClass")
	})
	return err
}

type SocketServiceClass struct {
	native unsafe.Pointer
}

func SocketServiceClassNewFromNative(native unsafe.Pointer) *SocketServiceClass {
	instance := &SocketServiceClass{native: native}

	return instance
}

/*
CastToSocketServiceClass down casts any arbitrary Object to SocketServiceClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a SocketServiceClass.
*/
func (recv *SocketServiceClass) CastToSocketServiceClass(object *gobject.Object) *SocketServiceClass {
	return SocketServiceClassNewFromNative(object.Native())
}

func (recv *SocketServiceClass) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentClass returns the C field 'parent_class'.
func (recv *SocketServiceClass) FieldParentClass() *SocketListenerClass {
	argValue := gi.StructFieldGet(socketServiceClassStruct, recv.native, "parent_class")
	value := SocketListenerClassNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *SocketServiceClass) SetFieldParentClass(value *SocketListenerClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.StructFieldSet(socketServiceClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value 'incoming' : for field getter : missing Type

// UNSUPPORTED : C value 'incoming' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved4' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved5' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved5' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved6' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved6' : for field setter : missing Type

// SocketServiceClassStruct creates an uninitialised SocketServiceClass.
func SocketServiceClassStruct() *SocketServiceClass {
	err := socketServiceClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := SocketServiceClassNewFromNative(socketServiceClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeSocketServiceClass)
	return structGo
}
func finalizeSocketServiceClass(obj *SocketServiceClass) {
	socketServiceClassStruct.Free(obj.native)
}

var socketServicePrivateStruct *gi.Struct
var socketServicePrivateStruct_Once sync.Once

func socketServicePrivateStruct_Set() error {
	var err error
	socketServicePrivateStruct_Once.Do(func() {
		socketServicePrivateStruct, err = gi.StructNew("Gio", "SocketServicePrivate")
	})
	return err
}

type SocketServicePrivate struct {
	native unsafe.Pointer
}

func SocketServicePrivateNewFromNative(native unsafe.Pointer) *SocketServicePrivate {
	instance := &SocketServicePrivate{native: native}

	return instance
}

/*
CastToSocketServicePrivate down casts any arbitrary Object to SocketServicePrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a SocketServicePrivate.
*/
func (recv *SocketServicePrivate) CastToSocketServicePrivate(object *gobject.Object) *SocketServicePrivate {
	return SocketServicePrivateNewFromNative(object.Native())
}

func (recv *SocketServicePrivate) Native() unsafe.Pointer {
	return recv.native
}

// SocketServicePrivateStruct creates an uninitialised SocketServicePrivate.
func SocketServicePrivateStruct() *SocketServicePrivate {
	err := socketServicePrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := SocketServicePrivateNewFromNative(socketServicePrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeSocketServicePrivate)
	return structGo
}
func finalizeSocketServicePrivate(obj *SocketServicePrivate) {
	socketServicePrivateStruct.Free(obj.native)
}

var srvTargetStruct *gi.Struct
var srvTargetStruct_Once sync.Once

func srvTargetStruct_Set() error {
	var err error
	srvTargetStruct_Once.Do(func() {
		srvTargetStruct, err = gi.StructNew("Gio", "SrvTarget")
	})
	return err
}

type SrvTarget struct {
	native unsafe.Pointer
}

func SrvTargetNewFromNative(native unsafe.Pointer) *SrvTarget {
	instance := &SrvTarget{native: native}

	return instance
}

/*
CastToSrvTarget down casts any arbitrary Object to SrvTarget.
Exercise care, as this is a potentially dangerous function
if the Object is not a SrvTarget.
*/
func (recv *SrvTarget) CastToSrvTarget(object *gobject.Object) *SrvTarget {
	return SrvTargetNewFromNative(object.Native())
}

func (recv *SrvTarget) Native() unsafe.Pointer {
	return recv.native
}

var srvTargetNewFunction *gi.Function
var srvTargetNewFunction_Once sync.Once

func srvTargetNewFunction_Set() error {
	var err error
	srvTargetNewFunction_Once.Do(func() {
		err = srvTargetStruct_Set()
		if err != nil {
			return
		}
		srvTargetNewFunction, err = srvTargetStruct.InvokerNew("new")
	})
	return err
}

// SrvTargetNew is a representation of the C type g_srv_target_new.
func SrvTargetNew(hostname string, port uint16, priority uint16, weight uint16) *SrvTarget {
	var inArgs [4]gi.Argument
	inArgs[0].SetString(hostname)
	inArgs[1].SetUint16(port)
	inArgs[2].SetUint16(priority)
	inArgs[3].SetUint16(weight)

	var ret gi.Argument

	err := srvTargetNewFunction_Set()
	if err == nil {
		ret = srvTargetNewFunction.Invoke(inArgs[:], nil)
	}

	retGo := SrvTargetNewFromNative(ret.Pointer())

	return retGo
}

var srvTargetCopyFunction *gi.Function
var srvTargetCopyFunction_Once sync.Once

func srvTargetCopyFunction_Set() error {
	var err error
	srvTargetCopyFunction_Once.Do(func() {
		err = srvTargetStruct_Set()
		if err != nil {
			return
		}
		srvTargetCopyFunction, err = srvTargetStruct.InvokerNew("copy")
	})
	return err
}

// Copy is a representation of the C type g_srv_target_copy.
func (recv *SrvTarget) Copy() *SrvTarget {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := srvTargetCopyFunction_Set()
	if err == nil {
		ret = srvTargetCopyFunction.Invoke(inArgs[:], nil)
	}

	retGo := SrvTargetNewFromNative(ret.Pointer())

	return retGo
}

var srvTargetFreeFunction *gi.Function
var srvTargetFreeFunction_Once sync.Once

func srvTargetFreeFunction_Set() error {
	var err error
	srvTargetFreeFunction_Once.Do(func() {
		err = srvTargetStruct_Set()
		if err != nil {
			return
		}
		srvTargetFreeFunction, err = srvTargetStruct.InvokerNew("free")
	})
	return err
}

// Free is a representation of the C type g_srv_target_free.
func (recv *SrvTarget) Free() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := srvTargetFreeFunction_Set()
	if err == nil {
		srvTargetFreeFunction.Invoke(inArgs[:], nil)
	}

	return
}

var srvTargetGetHostnameFunction *gi.Function
var srvTargetGetHostnameFunction_Once sync.Once

func srvTargetGetHostnameFunction_Set() error {
	var err error
	srvTargetGetHostnameFunction_Once.Do(func() {
		err = srvTargetStruct_Set()
		if err != nil {
			return
		}
		srvTargetGetHostnameFunction, err = srvTargetStruct.InvokerNew("get_hostname")
	})
	return err
}

// GetHostname is a representation of the C type g_srv_target_get_hostname.
func (recv *SrvTarget) GetHostname() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := srvTargetGetHostnameFunction_Set()
	if err == nil {
		ret = srvTargetGetHostnameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var srvTargetGetPortFunction *gi.Function
var srvTargetGetPortFunction_Once sync.Once

func srvTargetGetPortFunction_Set() error {
	var err error
	srvTargetGetPortFunction_Once.Do(func() {
		err = srvTargetStruct_Set()
		if err != nil {
			return
		}
		srvTargetGetPortFunction, err = srvTargetStruct.InvokerNew("get_port")
	})
	return err
}

// GetPort is a representation of the C type g_srv_target_get_port.
func (recv *SrvTarget) GetPort() uint16 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := srvTargetGetPortFunction_Set()
	if err == nil {
		ret = srvTargetGetPortFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint16()

	return retGo
}

var srvTargetGetPriorityFunction *gi.Function
var srvTargetGetPriorityFunction_Once sync.Once

func srvTargetGetPriorityFunction_Set() error {
	var err error
	srvTargetGetPriorityFunction_Once.Do(func() {
		err = srvTargetStruct_Set()
		if err != nil {
			return
		}
		srvTargetGetPriorityFunction, err = srvTargetStruct.InvokerNew("get_priority")
	})
	return err
}

// GetPriority is a representation of the C type g_srv_target_get_priority.
func (recv *SrvTarget) GetPriority() uint16 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := srvTargetGetPriorityFunction_Set()
	if err == nil {
		ret = srvTargetGetPriorityFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint16()

	return retGo
}

var srvTargetGetWeightFunction *gi.Function
var srvTargetGetWeightFunction_Once sync.Once

func srvTargetGetWeightFunction_Set() error {
	var err error
	srvTargetGetWeightFunction_Once.Do(func() {
		err = srvTargetStruct_Set()
		if err != nil {
			return
		}
		srvTargetGetWeightFunction, err = srvTargetStruct.InvokerNew("get_weight")
	})
	return err
}

// GetWeight is a representation of the C type g_srv_target_get_weight.
func (recv *SrvTarget) GetWeight() uint16 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := srvTargetGetWeightFunction_Set()
	if err == nil {
		ret = srvTargetGetWeightFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint16()

	return retGo
}

var staticResourceStruct *gi.Struct
var staticResourceStruct_Once sync.Once

func staticResourceStruct_Set() error {
	var err error
	staticResourceStruct_Once.Do(func() {
		staticResourceStruct, err = gi.StructNew("Gio", "StaticResource")
	})
	return err
}

type StaticResource struct {
	native unsafe.Pointer
}

func StaticResourceNewFromNative(native unsafe.Pointer) *StaticResource {
	instance := &StaticResource{native: native}

	return instance
}

/*
CastToStaticResource down casts any arbitrary Object to StaticResource.
Exercise care, as this is a potentially dangerous function
if the Object is not a StaticResource.
*/
func (recv *StaticResource) CastToStaticResource(object *gobject.Object) *StaticResource {
	return StaticResourceNewFromNative(object.Native())
}

func (recv *StaticResource) Native() unsafe.Pointer {
	return recv.native
}

var staticResourceFiniFunction *gi.Function
var staticResourceFiniFunction_Once sync.Once

func staticResourceFiniFunction_Set() error {
	var err error
	staticResourceFiniFunction_Once.Do(func() {
		err = staticResourceStruct_Set()
		if err != nil {
			return
		}
		staticResourceFiniFunction, err = staticResourceStruct.InvokerNew("fini")
	})
	return err
}

// Fini is a representation of the C type g_static_resource_fini.
func (recv *StaticResource) Fini() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := staticResourceFiniFunction_Set()
	if err == nil {
		staticResourceFiniFunction.Invoke(inArgs[:], nil)
	}

	return
}

var staticResourceGetResourceFunction *gi.Function
var staticResourceGetResourceFunction_Once sync.Once

func staticResourceGetResourceFunction_Set() error {
	var err error
	staticResourceGetResourceFunction_Once.Do(func() {
		err = staticResourceStruct_Set()
		if err != nil {
			return
		}
		staticResourceGetResourceFunction, err = staticResourceStruct.InvokerNew("get_resource")
	})
	return err
}

// GetResource is a representation of the C type g_static_resource_get_resource.
func (recv *StaticResource) GetResource() *Resource {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := staticResourceGetResourceFunction_Set()
	if err == nil {
		ret = staticResourceGetResourceFunction.Invoke(inArgs[:], nil)
	}

	retGo := ResourceNewFromNative(ret.Pointer())

	return retGo
}

var staticResourceInitFunction *gi.Function
var staticResourceInitFunction_Once sync.Once

func staticResourceInitFunction_Set() error {
	var err error
	staticResourceInitFunction_Once.Do(func() {
		err = staticResourceStruct_Set()
		if err != nil {
			return
		}
		staticResourceInitFunction, err = staticResourceStruct.InvokerNew("init")
	})
	return err
}

// Init is a representation of the C type g_static_resource_init.
func (recv *StaticResource) Init() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := staticResourceInitFunction_Set()
	if err == nil {
		staticResourceInitFunction.Invoke(inArgs[:], nil)
	}

	return
}

// StaticResourceStruct creates an uninitialised StaticResource.
func StaticResourceStruct() *StaticResource {
	err := staticResourceStruct_Set()
	if err != nil {
		return nil
	}

	structGo := StaticResourceNewFromNative(staticResourceStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeStaticResource)
	return structGo
}
func finalizeStaticResource(obj *StaticResource) {
	staticResourceStruct.Free(obj.native)
}

var taskClassStruct *gi.Struct
var taskClassStruct_Once sync.Once

func taskClassStruct_Set() error {
	var err error
	taskClassStruct_Once.Do(func() {
		taskClassStruct, err = gi.StructNew("Gio", "TaskClass")
	})
	return err
}

type TaskClass struct {
	native unsafe.Pointer
}

func TaskClassNewFromNative(native unsafe.Pointer) *TaskClass {
	instance := &TaskClass{native: native}

	return instance
}

/*
CastToTaskClass down casts any arbitrary Object to TaskClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a TaskClass.
*/
func (recv *TaskClass) CastToTaskClass(object *gobject.Object) *TaskClass {
	return TaskClassNewFromNative(object.Native())
}

func (recv *TaskClass) Native() unsafe.Pointer {
	return recv.native
}

// TaskClassStruct creates an uninitialised TaskClass.
func TaskClassStruct() *TaskClass {
	err := taskClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := TaskClassNewFromNative(taskClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeTaskClass)
	return structGo
}
func finalizeTaskClass(obj *TaskClass) {
	taskClassStruct.Free(obj.native)
}

var tcpConnectionClassStruct *gi.Struct
var tcpConnectionClassStruct_Once sync.Once

func tcpConnectionClassStruct_Set() error {
	var err error
	tcpConnectionClassStruct_Once.Do(func() {
		tcpConnectionClassStruct, err = gi.StructNew("Gio", "TcpConnectionClass")
	})
	return err
}

type TcpConnectionClass struct {
	native unsafe.Pointer
}

func TcpConnectionClassNewFromNative(native unsafe.Pointer) *TcpConnectionClass {
	instance := &TcpConnectionClass{native: native}

	return instance
}

/*
CastToTcpConnectionClass down casts any arbitrary Object to TcpConnectionClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a TcpConnectionClass.
*/
func (recv *TcpConnectionClass) CastToTcpConnectionClass(object *gobject.Object) *TcpConnectionClass {
	return TcpConnectionClassNewFromNative(object.Native())
}

func (recv *TcpConnectionClass) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentClass returns the C field 'parent_class'.
func (recv *TcpConnectionClass) FieldParentClass() *SocketConnectionClass {
	argValue := gi.StructFieldGet(tcpConnectionClassStruct, recv.native, "parent_class")
	value := SocketConnectionClassNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *TcpConnectionClass) SetFieldParentClass(value *SocketConnectionClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.StructFieldSet(tcpConnectionClassStruct, recv.native, "parent_class", argValue)
}

// TcpConnectionClassStruct creates an uninitialised TcpConnectionClass.
func TcpConnectionClassStruct() *TcpConnectionClass {
	err := tcpConnectionClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := TcpConnectionClassNewFromNative(tcpConnectionClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeTcpConnectionClass)
	return structGo
}
func finalizeTcpConnectionClass(obj *TcpConnectionClass) {
	tcpConnectionClassStruct.Free(obj.native)
}

var tcpConnectionPrivateStruct *gi.Struct
var tcpConnectionPrivateStruct_Once sync.Once

func tcpConnectionPrivateStruct_Set() error {
	var err error
	tcpConnectionPrivateStruct_Once.Do(func() {
		tcpConnectionPrivateStruct, err = gi.StructNew("Gio", "TcpConnectionPrivate")
	})
	return err
}

type TcpConnectionPrivate struct {
	native unsafe.Pointer
}

func TcpConnectionPrivateNewFromNative(native unsafe.Pointer) *TcpConnectionPrivate {
	instance := &TcpConnectionPrivate{native: native}

	return instance
}

/*
CastToTcpConnectionPrivate down casts any arbitrary Object to TcpConnectionPrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a TcpConnectionPrivate.
*/
func (recv *TcpConnectionPrivate) CastToTcpConnectionPrivate(object *gobject.Object) *TcpConnectionPrivate {
	return TcpConnectionPrivateNewFromNative(object.Native())
}

func (recv *TcpConnectionPrivate) Native() unsafe.Pointer {
	return recv.native
}

// TcpConnectionPrivateStruct creates an uninitialised TcpConnectionPrivate.
func TcpConnectionPrivateStruct() *TcpConnectionPrivate {
	err := tcpConnectionPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := TcpConnectionPrivateNewFromNative(tcpConnectionPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeTcpConnectionPrivate)
	return structGo
}
func finalizeTcpConnectionPrivate(obj *TcpConnectionPrivate) {
	tcpConnectionPrivateStruct.Free(obj.native)
}

var tcpWrapperConnectionClassStruct *gi.Struct
var tcpWrapperConnectionClassStruct_Once sync.Once

func tcpWrapperConnectionClassStruct_Set() error {
	var err error
	tcpWrapperConnectionClassStruct_Once.Do(func() {
		tcpWrapperConnectionClassStruct, err = gi.StructNew("Gio", "TcpWrapperConnectionClass")
	})
	return err
}

type TcpWrapperConnectionClass struct {
	native unsafe.Pointer
}

func TcpWrapperConnectionClassNewFromNative(native unsafe.Pointer) *TcpWrapperConnectionClass {
	instance := &TcpWrapperConnectionClass{native: native}

	return instance
}

/*
CastToTcpWrapperConnectionClass down casts any arbitrary Object to TcpWrapperConnectionClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a TcpWrapperConnectionClass.
*/
func (recv *TcpWrapperConnectionClass) CastToTcpWrapperConnectionClass(object *gobject.Object) *TcpWrapperConnectionClass {
	return TcpWrapperConnectionClassNewFromNative(object.Native())
}

func (recv *TcpWrapperConnectionClass) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentClass returns the C field 'parent_class'.
func (recv *TcpWrapperConnectionClass) FieldParentClass() *TcpConnectionClass {
	argValue := gi.StructFieldGet(tcpWrapperConnectionClassStruct, recv.native, "parent_class")
	value := TcpConnectionClassNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *TcpWrapperConnectionClass) SetFieldParentClass(value *TcpConnectionClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.StructFieldSet(tcpWrapperConnectionClassStruct, recv.native, "parent_class", argValue)
}

// TcpWrapperConnectionClassStruct creates an uninitialised TcpWrapperConnectionClass.
func TcpWrapperConnectionClassStruct() *TcpWrapperConnectionClass {
	err := tcpWrapperConnectionClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := TcpWrapperConnectionClassNewFromNative(tcpWrapperConnectionClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeTcpWrapperConnectionClass)
	return structGo
}
func finalizeTcpWrapperConnectionClass(obj *TcpWrapperConnectionClass) {
	tcpWrapperConnectionClassStruct.Free(obj.native)
}

var tcpWrapperConnectionPrivateStruct *gi.Struct
var tcpWrapperConnectionPrivateStruct_Once sync.Once

func tcpWrapperConnectionPrivateStruct_Set() error {
	var err error
	tcpWrapperConnectionPrivateStruct_Once.Do(func() {
		tcpWrapperConnectionPrivateStruct, err = gi.StructNew("Gio", "TcpWrapperConnectionPrivate")
	})
	return err
}

type TcpWrapperConnectionPrivate struct {
	native unsafe.Pointer
}

func TcpWrapperConnectionPrivateNewFromNative(native unsafe.Pointer) *TcpWrapperConnectionPrivate {
	instance := &TcpWrapperConnectionPrivate{native: native}

	return instance
}

/*
CastToTcpWrapperConnectionPrivate down casts any arbitrary Object to TcpWrapperConnectionPrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a TcpWrapperConnectionPrivate.
*/
func (recv *TcpWrapperConnectionPrivate) CastToTcpWrapperConnectionPrivate(object *gobject.Object) *TcpWrapperConnectionPrivate {
	return TcpWrapperConnectionPrivateNewFromNative(object.Native())
}

func (recv *TcpWrapperConnectionPrivate) Native() unsafe.Pointer {
	return recv.native
}

// TcpWrapperConnectionPrivateStruct creates an uninitialised TcpWrapperConnectionPrivate.
func TcpWrapperConnectionPrivateStruct() *TcpWrapperConnectionPrivate {
	err := tcpWrapperConnectionPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := TcpWrapperConnectionPrivateNewFromNative(tcpWrapperConnectionPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeTcpWrapperConnectionPrivate)
	return structGo
}
func finalizeTcpWrapperConnectionPrivate(obj *TcpWrapperConnectionPrivate) {
	tcpWrapperConnectionPrivateStruct.Free(obj.native)
}

var themedIconClassStruct *gi.Struct
var themedIconClassStruct_Once sync.Once

func themedIconClassStruct_Set() error {
	var err error
	themedIconClassStruct_Once.Do(func() {
		themedIconClassStruct, err = gi.StructNew("Gio", "ThemedIconClass")
	})
	return err
}

type ThemedIconClass struct {
	native unsafe.Pointer
}

func ThemedIconClassNewFromNative(native unsafe.Pointer) *ThemedIconClass {
	instance := &ThemedIconClass{native: native}

	return instance
}

/*
CastToThemedIconClass down casts any arbitrary Object to ThemedIconClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a ThemedIconClass.
*/
func (recv *ThemedIconClass) CastToThemedIconClass(object *gobject.Object) *ThemedIconClass {
	return ThemedIconClassNewFromNative(object.Native())
}

func (recv *ThemedIconClass) Native() unsafe.Pointer {
	return recv.native
}

// ThemedIconClassStruct creates an uninitialised ThemedIconClass.
func ThemedIconClassStruct() *ThemedIconClass {
	err := themedIconClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := ThemedIconClassNewFromNative(themedIconClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeThemedIconClass)
	return structGo
}
func finalizeThemedIconClass(obj *ThemedIconClass) {
	themedIconClassStruct.Free(obj.native)
}

var threadedSocketServiceClassStruct *gi.Struct
var threadedSocketServiceClassStruct_Once sync.Once

func threadedSocketServiceClassStruct_Set() error {
	var err error
	threadedSocketServiceClassStruct_Once.Do(func() {
		threadedSocketServiceClassStruct, err = gi.StructNew("Gio", "ThreadedSocketServiceClass")
	})
	return err
}

type ThreadedSocketServiceClass struct {
	native unsafe.Pointer
}

func ThreadedSocketServiceClassNewFromNative(native unsafe.Pointer) *ThreadedSocketServiceClass {
	instance := &ThreadedSocketServiceClass{native: native}

	return instance
}

/*
CastToThreadedSocketServiceClass down casts any arbitrary Object to ThreadedSocketServiceClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a ThreadedSocketServiceClass.
*/
func (recv *ThreadedSocketServiceClass) CastToThreadedSocketServiceClass(object *gobject.Object) *ThreadedSocketServiceClass {
	return ThreadedSocketServiceClassNewFromNative(object.Native())
}

func (recv *ThreadedSocketServiceClass) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentClass returns the C field 'parent_class'.
func (recv *ThreadedSocketServiceClass) FieldParentClass() *SocketServiceClass {
	argValue := gi.StructFieldGet(threadedSocketServiceClassStruct, recv.native, "parent_class")
	value := SocketServiceClassNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *ThreadedSocketServiceClass) SetFieldParentClass(value *SocketServiceClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.StructFieldSet(threadedSocketServiceClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value 'run' : for field getter : missing Type

// UNSUPPORTED : C value 'run' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved4' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved5' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved5' : for field setter : missing Type

// ThreadedSocketServiceClassStruct creates an uninitialised ThreadedSocketServiceClass.
func ThreadedSocketServiceClassStruct() *ThreadedSocketServiceClass {
	err := threadedSocketServiceClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := ThreadedSocketServiceClassNewFromNative(threadedSocketServiceClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeThreadedSocketServiceClass)
	return structGo
}
func finalizeThreadedSocketServiceClass(obj *ThreadedSocketServiceClass) {
	threadedSocketServiceClassStruct.Free(obj.native)
}

var threadedSocketServicePrivateStruct *gi.Struct
var threadedSocketServicePrivateStruct_Once sync.Once

func threadedSocketServicePrivateStruct_Set() error {
	var err error
	threadedSocketServicePrivateStruct_Once.Do(func() {
		threadedSocketServicePrivateStruct, err = gi.StructNew("Gio", "ThreadedSocketServicePrivate")
	})
	return err
}

type ThreadedSocketServicePrivate struct {
	native unsafe.Pointer
}

func ThreadedSocketServicePrivateNewFromNative(native unsafe.Pointer) *ThreadedSocketServicePrivate {
	instance := &ThreadedSocketServicePrivate{native: native}

	return instance
}

/*
CastToThreadedSocketServicePrivate down casts any arbitrary Object to ThreadedSocketServicePrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a ThreadedSocketServicePrivate.
*/
func (recv *ThreadedSocketServicePrivate) CastToThreadedSocketServicePrivate(object *gobject.Object) *ThreadedSocketServicePrivate {
	return ThreadedSocketServicePrivateNewFromNative(object.Native())
}

func (recv *ThreadedSocketServicePrivate) Native() unsafe.Pointer {
	return recv.native
}

// ThreadedSocketServicePrivateStruct creates an uninitialised ThreadedSocketServicePrivate.
func ThreadedSocketServicePrivateStruct() *ThreadedSocketServicePrivate {
	err := threadedSocketServicePrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := ThreadedSocketServicePrivateNewFromNative(threadedSocketServicePrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeThreadedSocketServicePrivate)
	return structGo
}
func finalizeThreadedSocketServicePrivate(obj *ThreadedSocketServicePrivate) {
	threadedSocketServicePrivateStruct.Free(obj.native)
}

var tlsBackendInterfaceStruct *gi.Struct
var tlsBackendInterfaceStruct_Once sync.Once

func tlsBackendInterfaceStruct_Set() error {
	var err error
	tlsBackendInterfaceStruct_Once.Do(func() {
		tlsBackendInterfaceStruct, err = gi.StructNew("Gio", "TlsBackendInterface")
	})
	return err
}

type TlsBackendInterface struct {
	native unsafe.Pointer
}

func TlsBackendInterfaceNewFromNative(native unsafe.Pointer) *TlsBackendInterface {
	instance := &TlsBackendInterface{native: native}

	return instance
}

/*
CastToTlsBackendInterface down casts any arbitrary Object to TlsBackendInterface.
Exercise care, as this is a potentially dangerous function
if the Object is not a TlsBackendInterface.
*/
func (recv *TlsBackendInterface) CastToTlsBackendInterface(object *gobject.Object) *TlsBackendInterface {
	return TlsBackendInterfaceNewFromNative(object.Native())
}

func (recv *TlsBackendInterface) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'g_iface' : for field getter : no Go type for 'GObject.TypeInterface'

// UNSUPPORTED : C value 'g_iface' : for field setter : no Go type for 'GObject.TypeInterface'

// UNSUPPORTED : C value 'supports_tls' : for field getter : missing Type

// UNSUPPORTED : C value 'supports_tls' : for field setter : missing Type

// UNSUPPORTED : C value 'get_certificate_type' : for field getter : missing Type

// UNSUPPORTED : C value 'get_certificate_type' : for field setter : missing Type

// UNSUPPORTED : C value 'get_client_connection_type' : for field getter : missing Type

// UNSUPPORTED : C value 'get_client_connection_type' : for field setter : missing Type

// UNSUPPORTED : C value 'get_server_connection_type' : for field getter : missing Type

// UNSUPPORTED : C value 'get_server_connection_type' : for field setter : missing Type

// UNSUPPORTED : C value 'get_file_database_type' : for field getter : missing Type

// UNSUPPORTED : C value 'get_file_database_type' : for field setter : missing Type

// UNSUPPORTED : C value 'get_default_database' : for field getter : missing Type

// UNSUPPORTED : C value 'get_default_database' : for field setter : missing Type

// UNSUPPORTED : C value 'supports_dtls' : for field getter : missing Type

// UNSUPPORTED : C value 'supports_dtls' : for field setter : missing Type

// UNSUPPORTED : C value 'get_dtls_client_connection_type' : for field getter : missing Type

// UNSUPPORTED : C value 'get_dtls_client_connection_type' : for field setter : missing Type

// UNSUPPORTED : C value 'get_dtls_server_connection_type' : for field getter : missing Type

// UNSUPPORTED : C value 'get_dtls_server_connection_type' : for field setter : missing Type

// TlsBackendInterfaceStruct creates an uninitialised TlsBackendInterface.
func TlsBackendInterfaceStruct() *TlsBackendInterface {
	err := tlsBackendInterfaceStruct_Set()
	if err != nil {
		return nil
	}

	structGo := TlsBackendInterfaceNewFromNative(tlsBackendInterfaceStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeTlsBackendInterface)
	return structGo
}
func finalizeTlsBackendInterface(obj *TlsBackendInterface) {
	tlsBackendInterfaceStruct.Free(obj.native)
}

var tlsCertificateClassStruct *gi.Struct
var tlsCertificateClassStruct_Once sync.Once

func tlsCertificateClassStruct_Set() error {
	var err error
	tlsCertificateClassStruct_Once.Do(func() {
		tlsCertificateClassStruct, err = gi.StructNew("Gio", "TlsCertificateClass")
	})
	return err
}

type TlsCertificateClass struct {
	native unsafe.Pointer
}

func TlsCertificateClassNewFromNative(native unsafe.Pointer) *TlsCertificateClass {
	instance := &TlsCertificateClass{native: native}

	return instance
}

/*
CastToTlsCertificateClass down casts any arbitrary Object to TlsCertificateClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a TlsCertificateClass.
*/
func (recv *TlsCertificateClass) CastToTlsCertificateClass(object *gobject.Object) *TlsCertificateClass {
	return TlsCertificateClassNewFromNative(object.Native())
}

func (recv *TlsCertificateClass) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'verify' : for field getter : missing Type

// UNSUPPORTED : C value 'verify' : for field setter : missing Type

// TlsCertificateClassStruct creates an uninitialised TlsCertificateClass.
func TlsCertificateClassStruct() *TlsCertificateClass {
	err := tlsCertificateClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := TlsCertificateClassNewFromNative(tlsCertificateClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeTlsCertificateClass)
	return structGo
}
func finalizeTlsCertificateClass(obj *TlsCertificateClass) {
	tlsCertificateClassStruct.Free(obj.native)
}

var tlsCertificatePrivateStruct *gi.Struct
var tlsCertificatePrivateStruct_Once sync.Once

func tlsCertificatePrivateStruct_Set() error {
	var err error
	tlsCertificatePrivateStruct_Once.Do(func() {
		tlsCertificatePrivateStruct, err = gi.StructNew("Gio", "TlsCertificatePrivate")
	})
	return err
}

type TlsCertificatePrivate struct {
	native unsafe.Pointer
}

func TlsCertificatePrivateNewFromNative(native unsafe.Pointer) *TlsCertificatePrivate {
	instance := &TlsCertificatePrivate{native: native}

	return instance
}

/*
CastToTlsCertificatePrivate down casts any arbitrary Object to TlsCertificatePrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a TlsCertificatePrivate.
*/
func (recv *TlsCertificatePrivate) CastToTlsCertificatePrivate(object *gobject.Object) *TlsCertificatePrivate {
	return TlsCertificatePrivateNewFromNative(object.Native())
}

func (recv *TlsCertificatePrivate) Native() unsafe.Pointer {
	return recv.native
}

// TlsCertificatePrivateStruct creates an uninitialised TlsCertificatePrivate.
func TlsCertificatePrivateStruct() *TlsCertificatePrivate {
	err := tlsCertificatePrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := TlsCertificatePrivateNewFromNative(tlsCertificatePrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeTlsCertificatePrivate)
	return structGo
}
func finalizeTlsCertificatePrivate(obj *TlsCertificatePrivate) {
	tlsCertificatePrivateStruct.Free(obj.native)
}

var tlsClientConnectionInterfaceStruct *gi.Struct
var tlsClientConnectionInterfaceStruct_Once sync.Once

func tlsClientConnectionInterfaceStruct_Set() error {
	var err error
	tlsClientConnectionInterfaceStruct_Once.Do(func() {
		tlsClientConnectionInterfaceStruct, err = gi.StructNew("Gio", "TlsClientConnectionInterface")
	})
	return err
}

type TlsClientConnectionInterface struct {
	native unsafe.Pointer
}

func TlsClientConnectionInterfaceNewFromNative(native unsafe.Pointer) *TlsClientConnectionInterface {
	instance := &TlsClientConnectionInterface{native: native}

	return instance
}

/*
CastToTlsClientConnectionInterface down casts any arbitrary Object to TlsClientConnectionInterface.
Exercise care, as this is a potentially dangerous function
if the Object is not a TlsClientConnectionInterface.
*/
func (recv *TlsClientConnectionInterface) CastToTlsClientConnectionInterface(object *gobject.Object) *TlsClientConnectionInterface {
	return TlsClientConnectionInterfaceNewFromNative(object.Native())
}

func (recv *TlsClientConnectionInterface) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'g_iface' : for field getter : no Go type for 'GObject.TypeInterface'

// UNSUPPORTED : C value 'g_iface' : for field setter : no Go type for 'GObject.TypeInterface'

// UNSUPPORTED : C value 'copy_session_state' : for field getter : missing Type

// UNSUPPORTED : C value 'copy_session_state' : for field setter : missing Type

// TlsClientConnectionInterfaceStruct creates an uninitialised TlsClientConnectionInterface.
func TlsClientConnectionInterfaceStruct() *TlsClientConnectionInterface {
	err := tlsClientConnectionInterfaceStruct_Set()
	if err != nil {
		return nil
	}

	structGo := TlsClientConnectionInterfaceNewFromNative(tlsClientConnectionInterfaceStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeTlsClientConnectionInterface)
	return structGo
}
func finalizeTlsClientConnectionInterface(obj *TlsClientConnectionInterface) {
	tlsClientConnectionInterfaceStruct.Free(obj.native)
}

var tlsConnectionClassStruct *gi.Struct
var tlsConnectionClassStruct_Once sync.Once

func tlsConnectionClassStruct_Set() error {
	var err error
	tlsConnectionClassStruct_Once.Do(func() {
		tlsConnectionClassStruct, err = gi.StructNew("Gio", "TlsConnectionClass")
	})
	return err
}

type TlsConnectionClass struct {
	native unsafe.Pointer
}

func TlsConnectionClassNewFromNative(native unsafe.Pointer) *TlsConnectionClass {
	instance := &TlsConnectionClass{native: native}

	return instance
}

/*
CastToTlsConnectionClass down casts any arbitrary Object to TlsConnectionClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a TlsConnectionClass.
*/
func (recv *TlsConnectionClass) CastToTlsConnectionClass(object *gobject.Object) *TlsConnectionClass {
	return TlsConnectionClassNewFromNative(object.Native())
}

func (recv *TlsConnectionClass) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentClass returns the C field 'parent_class'.
func (recv *TlsConnectionClass) FieldParentClass() *IOStreamClass {
	argValue := gi.StructFieldGet(tlsConnectionClassStruct, recv.native, "parent_class")
	value := IOStreamClassNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *TlsConnectionClass) SetFieldParentClass(value *IOStreamClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.StructFieldSet(tlsConnectionClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value 'accept_certificate' : for field getter : missing Type

// UNSUPPORTED : C value 'accept_certificate' : for field setter : missing Type

// UNSUPPORTED : C value 'handshake' : for field getter : missing Type

// UNSUPPORTED : C value 'handshake' : for field setter : missing Type

// UNSUPPORTED : C value 'handshake_async' : for field getter : missing Type

// UNSUPPORTED : C value 'handshake_async' : for field setter : missing Type

// UNSUPPORTED : C value 'handshake_finish' : for field getter : missing Type

// UNSUPPORTED : C value 'handshake_finish' : for field setter : missing Type

// TlsConnectionClassStruct creates an uninitialised TlsConnectionClass.
func TlsConnectionClassStruct() *TlsConnectionClass {
	err := tlsConnectionClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := TlsConnectionClassNewFromNative(tlsConnectionClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeTlsConnectionClass)
	return structGo
}
func finalizeTlsConnectionClass(obj *TlsConnectionClass) {
	tlsConnectionClassStruct.Free(obj.native)
}

var tlsConnectionPrivateStruct *gi.Struct
var tlsConnectionPrivateStruct_Once sync.Once

func tlsConnectionPrivateStruct_Set() error {
	var err error
	tlsConnectionPrivateStruct_Once.Do(func() {
		tlsConnectionPrivateStruct, err = gi.StructNew("Gio", "TlsConnectionPrivate")
	})
	return err
}

type TlsConnectionPrivate struct {
	native unsafe.Pointer
}

func TlsConnectionPrivateNewFromNative(native unsafe.Pointer) *TlsConnectionPrivate {
	instance := &TlsConnectionPrivate{native: native}

	return instance
}

/*
CastToTlsConnectionPrivate down casts any arbitrary Object to TlsConnectionPrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a TlsConnectionPrivate.
*/
func (recv *TlsConnectionPrivate) CastToTlsConnectionPrivate(object *gobject.Object) *TlsConnectionPrivate {
	return TlsConnectionPrivateNewFromNative(object.Native())
}

func (recv *TlsConnectionPrivate) Native() unsafe.Pointer {
	return recv.native
}

// TlsConnectionPrivateStruct creates an uninitialised TlsConnectionPrivate.
func TlsConnectionPrivateStruct() *TlsConnectionPrivate {
	err := tlsConnectionPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := TlsConnectionPrivateNewFromNative(tlsConnectionPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeTlsConnectionPrivate)
	return structGo
}
func finalizeTlsConnectionPrivate(obj *TlsConnectionPrivate) {
	tlsConnectionPrivateStruct.Free(obj.native)
}

var tlsDatabaseClassStruct *gi.Struct
var tlsDatabaseClassStruct_Once sync.Once

func tlsDatabaseClassStruct_Set() error {
	var err error
	tlsDatabaseClassStruct_Once.Do(func() {
		tlsDatabaseClassStruct, err = gi.StructNew("Gio", "TlsDatabaseClass")
	})
	return err
}

type TlsDatabaseClass struct {
	native unsafe.Pointer
}

func TlsDatabaseClassNewFromNative(native unsafe.Pointer) *TlsDatabaseClass {
	instance := &TlsDatabaseClass{native: native}

	return instance
}

/*
CastToTlsDatabaseClass down casts any arbitrary Object to TlsDatabaseClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a TlsDatabaseClass.
*/
func (recv *TlsDatabaseClass) CastToTlsDatabaseClass(object *gobject.Object) *TlsDatabaseClass {
	return TlsDatabaseClassNewFromNative(object.Native())
}

func (recv *TlsDatabaseClass) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'verify_chain' : for field getter : missing Type

// UNSUPPORTED : C value 'verify_chain' : for field setter : missing Type

// UNSUPPORTED : C value 'verify_chain_async' : for field getter : missing Type

// UNSUPPORTED : C value 'verify_chain_async' : for field setter : missing Type

// UNSUPPORTED : C value 'verify_chain_finish' : for field getter : missing Type

// UNSUPPORTED : C value 'verify_chain_finish' : for field setter : missing Type

// UNSUPPORTED : C value 'create_certificate_handle' : for field getter : missing Type

// UNSUPPORTED : C value 'create_certificate_handle' : for field setter : missing Type

// UNSUPPORTED : C value 'lookup_certificate_for_handle' : for field getter : missing Type

// UNSUPPORTED : C value 'lookup_certificate_for_handle' : for field setter : missing Type

// UNSUPPORTED : C value 'lookup_certificate_for_handle_async' : for field getter : missing Type

// UNSUPPORTED : C value 'lookup_certificate_for_handle_async' : for field setter : missing Type

// UNSUPPORTED : C value 'lookup_certificate_for_handle_finish' : for field getter : missing Type

// UNSUPPORTED : C value 'lookup_certificate_for_handle_finish' : for field setter : missing Type

// UNSUPPORTED : C value 'lookup_certificate_issuer' : for field getter : missing Type

// UNSUPPORTED : C value 'lookup_certificate_issuer' : for field setter : missing Type

// UNSUPPORTED : C value 'lookup_certificate_issuer_async' : for field getter : missing Type

// UNSUPPORTED : C value 'lookup_certificate_issuer_async' : for field setter : missing Type

// UNSUPPORTED : C value 'lookup_certificate_issuer_finish' : for field getter : missing Type

// UNSUPPORTED : C value 'lookup_certificate_issuer_finish' : for field setter : missing Type

// UNSUPPORTED : C value 'lookup_certificates_issued_by' : for field getter : missing Type

// UNSUPPORTED : C value 'lookup_certificates_issued_by' : for field setter : missing Type

// UNSUPPORTED : C value 'lookup_certificates_issued_by_async' : for field getter : missing Type

// UNSUPPORTED : C value 'lookup_certificates_issued_by_async' : for field setter : missing Type

// UNSUPPORTED : C value 'lookup_certificates_issued_by_finish' : for field getter : missing Type

// UNSUPPORTED : C value 'lookup_certificates_issued_by_finish' : for field setter : missing Type

// TlsDatabaseClassStruct creates an uninitialised TlsDatabaseClass.
func TlsDatabaseClassStruct() *TlsDatabaseClass {
	err := tlsDatabaseClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := TlsDatabaseClassNewFromNative(tlsDatabaseClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeTlsDatabaseClass)
	return structGo
}
func finalizeTlsDatabaseClass(obj *TlsDatabaseClass) {
	tlsDatabaseClassStruct.Free(obj.native)
}

var tlsDatabasePrivateStruct *gi.Struct
var tlsDatabasePrivateStruct_Once sync.Once

func tlsDatabasePrivateStruct_Set() error {
	var err error
	tlsDatabasePrivateStruct_Once.Do(func() {
		tlsDatabasePrivateStruct, err = gi.StructNew("Gio", "TlsDatabasePrivate")
	})
	return err
}

type TlsDatabasePrivate struct {
	native unsafe.Pointer
}

func TlsDatabasePrivateNewFromNative(native unsafe.Pointer) *TlsDatabasePrivate {
	instance := &TlsDatabasePrivate{native: native}

	return instance
}

/*
CastToTlsDatabasePrivate down casts any arbitrary Object to TlsDatabasePrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a TlsDatabasePrivate.
*/
func (recv *TlsDatabasePrivate) CastToTlsDatabasePrivate(object *gobject.Object) *TlsDatabasePrivate {
	return TlsDatabasePrivateNewFromNative(object.Native())
}

func (recv *TlsDatabasePrivate) Native() unsafe.Pointer {
	return recv.native
}

// TlsDatabasePrivateStruct creates an uninitialised TlsDatabasePrivate.
func TlsDatabasePrivateStruct() *TlsDatabasePrivate {
	err := tlsDatabasePrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := TlsDatabasePrivateNewFromNative(tlsDatabasePrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeTlsDatabasePrivate)
	return structGo
}
func finalizeTlsDatabasePrivate(obj *TlsDatabasePrivate) {
	tlsDatabasePrivateStruct.Free(obj.native)
}

var tlsFileDatabaseInterfaceStruct *gi.Struct
var tlsFileDatabaseInterfaceStruct_Once sync.Once

func tlsFileDatabaseInterfaceStruct_Set() error {
	var err error
	tlsFileDatabaseInterfaceStruct_Once.Do(func() {
		tlsFileDatabaseInterfaceStruct, err = gi.StructNew("Gio", "TlsFileDatabaseInterface")
	})
	return err
}

type TlsFileDatabaseInterface struct {
	native unsafe.Pointer
}

func TlsFileDatabaseInterfaceNewFromNative(native unsafe.Pointer) *TlsFileDatabaseInterface {
	instance := &TlsFileDatabaseInterface{native: native}

	return instance
}

/*
CastToTlsFileDatabaseInterface down casts any arbitrary Object to TlsFileDatabaseInterface.
Exercise care, as this is a potentially dangerous function
if the Object is not a TlsFileDatabaseInterface.
*/
func (recv *TlsFileDatabaseInterface) CastToTlsFileDatabaseInterface(object *gobject.Object) *TlsFileDatabaseInterface {
	return TlsFileDatabaseInterfaceNewFromNative(object.Native())
}

func (recv *TlsFileDatabaseInterface) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'g_iface' : for field getter : no Go type for 'GObject.TypeInterface'

// UNSUPPORTED : C value 'g_iface' : for field setter : no Go type for 'GObject.TypeInterface'

// TlsFileDatabaseInterfaceStruct creates an uninitialised TlsFileDatabaseInterface.
func TlsFileDatabaseInterfaceStruct() *TlsFileDatabaseInterface {
	err := tlsFileDatabaseInterfaceStruct_Set()
	if err != nil {
		return nil
	}

	structGo := TlsFileDatabaseInterfaceNewFromNative(tlsFileDatabaseInterfaceStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeTlsFileDatabaseInterface)
	return structGo
}
func finalizeTlsFileDatabaseInterface(obj *TlsFileDatabaseInterface) {
	tlsFileDatabaseInterfaceStruct.Free(obj.native)
}

var tlsInteractionClassStruct *gi.Struct
var tlsInteractionClassStruct_Once sync.Once

func tlsInteractionClassStruct_Set() error {
	var err error
	tlsInteractionClassStruct_Once.Do(func() {
		tlsInteractionClassStruct, err = gi.StructNew("Gio", "TlsInteractionClass")
	})
	return err
}

type TlsInteractionClass struct {
	native unsafe.Pointer
}

func TlsInteractionClassNewFromNative(native unsafe.Pointer) *TlsInteractionClass {
	instance := &TlsInteractionClass{native: native}

	return instance
}

/*
CastToTlsInteractionClass down casts any arbitrary Object to TlsInteractionClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a TlsInteractionClass.
*/
func (recv *TlsInteractionClass) CastToTlsInteractionClass(object *gobject.Object) *TlsInteractionClass {
	return TlsInteractionClassNewFromNative(object.Native())
}

func (recv *TlsInteractionClass) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'ask_password' : for field getter : missing Type

// UNSUPPORTED : C value 'ask_password' : for field setter : missing Type

// UNSUPPORTED : C value 'ask_password_async' : for field getter : missing Type

// UNSUPPORTED : C value 'ask_password_async' : for field setter : missing Type

// UNSUPPORTED : C value 'ask_password_finish' : for field getter : missing Type

// UNSUPPORTED : C value 'ask_password_finish' : for field setter : missing Type

// UNSUPPORTED : C value 'request_certificate' : for field getter : missing Type

// UNSUPPORTED : C value 'request_certificate' : for field setter : missing Type

// UNSUPPORTED : C value 'request_certificate_async' : for field getter : missing Type

// UNSUPPORTED : C value 'request_certificate_async' : for field setter : missing Type

// UNSUPPORTED : C value 'request_certificate_finish' : for field getter : missing Type

// UNSUPPORTED : C value 'request_certificate_finish' : for field setter : missing Type

// TlsInteractionClassStruct creates an uninitialised TlsInteractionClass.
func TlsInteractionClassStruct() *TlsInteractionClass {
	err := tlsInteractionClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := TlsInteractionClassNewFromNative(tlsInteractionClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeTlsInteractionClass)
	return structGo
}
func finalizeTlsInteractionClass(obj *TlsInteractionClass) {
	tlsInteractionClassStruct.Free(obj.native)
}

var tlsInteractionPrivateStruct *gi.Struct
var tlsInteractionPrivateStruct_Once sync.Once

func tlsInteractionPrivateStruct_Set() error {
	var err error
	tlsInteractionPrivateStruct_Once.Do(func() {
		tlsInteractionPrivateStruct, err = gi.StructNew("Gio", "TlsInteractionPrivate")
	})
	return err
}

type TlsInteractionPrivate struct {
	native unsafe.Pointer
}

func TlsInteractionPrivateNewFromNative(native unsafe.Pointer) *TlsInteractionPrivate {
	instance := &TlsInteractionPrivate{native: native}

	return instance
}

/*
CastToTlsInteractionPrivate down casts any arbitrary Object to TlsInteractionPrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a TlsInteractionPrivate.
*/
func (recv *TlsInteractionPrivate) CastToTlsInteractionPrivate(object *gobject.Object) *TlsInteractionPrivate {
	return TlsInteractionPrivateNewFromNative(object.Native())
}

func (recv *TlsInteractionPrivate) Native() unsafe.Pointer {
	return recv.native
}

// TlsInteractionPrivateStruct creates an uninitialised TlsInteractionPrivate.
func TlsInteractionPrivateStruct() *TlsInteractionPrivate {
	err := tlsInteractionPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := TlsInteractionPrivateNewFromNative(tlsInteractionPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeTlsInteractionPrivate)
	return structGo
}
func finalizeTlsInteractionPrivate(obj *TlsInteractionPrivate) {
	tlsInteractionPrivateStruct.Free(obj.native)
}

var tlsPasswordClassStruct *gi.Struct
var tlsPasswordClassStruct_Once sync.Once

func tlsPasswordClassStruct_Set() error {
	var err error
	tlsPasswordClassStruct_Once.Do(func() {
		tlsPasswordClassStruct, err = gi.StructNew("Gio", "TlsPasswordClass")
	})
	return err
}

type TlsPasswordClass struct {
	native unsafe.Pointer
}

func TlsPasswordClassNewFromNative(native unsafe.Pointer) *TlsPasswordClass {
	instance := &TlsPasswordClass{native: native}

	return instance
}

/*
CastToTlsPasswordClass down casts any arbitrary Object to TlsPasswordClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a TlsPasswordClass.
*/
func (recv *TlsPasswordClass) CastToTlsPasswordClass(object *gobject.Object) *TlsPasswordClass {
	return TlsPasswordClassNewFromNative(object.Native())
}

func (recv *TlsPasswordClass) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'get_value' : for field getter : missing Type

// UNSUPPORTED : C value 'get_value' : for field setter : missing Type

// UNSUPPORTED : C value 'set_value' : for field getter : missing Type

// UNSUPPORTED : C value 'set_value' : for field setter : missing Type

// UNSUPPORTED : C value 'get_default_warning' : for field getter : missing Type

// UNSUPPORTED : C value 'get_default_warning' : for field setter : missing Type

// TlsPasswordClassStruct creates an uninitialised TlsPasswordClass.
func TlsPasswordClassStruct() *TlsPasswordClass {
	err := tlsPasswordClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := TlsPasswordClassNewFromNative(tlsPasswordClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeTlsPasswordClass)
	return structGo
}
func finalizeTlsPasswordClass(obj *TlsPasswordClass) {
	tlsPasswordClassStruct.Free(obj.native)
}

var tlsPasswordPrivateStruct *gi.Struct
var tlsPasswordPrivateStruct_Once sync.Once

func tlsPasswordPrivateStruct_Set() error {
	var err error
	tlsPasswordPrivateStruct_Once.Do(func() {
		tlsPasswordPrivateStruct, err = gi.StructNew("Gio", "TlsPasswordPrivate")
	})
	return err
}

type TlsPasswordPrivate struct {
	native unsafe.Pointer
}

func TlsPasswordPrivateNewFromNative(native unsafe.Pointer) *TlsPasswordPrivate {
	instance := &TlsPasswordPrivate{native: native}

	return instance
}

/*
CastToTlsPasswordPrivate down casts any arbitrary Object to TlsPasswordPrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a TlsPasswordPrivate.
*/
func (recv *TlsPasswordPrivate) CastToTlsPasswordPrivate(object *gobject.Object) *TlsPasswordPrivate {
	return TlsPasswordPrivateNewFromNative(object.Native())
}

func (recv *TlsPasswordPrivate) Native() unsafe.Pointer {
	return recv.native
}

// TlsPasswordPrivateStruct creates an uninitialised TlsPasswordPrivate.
func TlsPasswordPrivateStruct() *TlsPasswordPrivate {
	err := tlsPasswordPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := TlsPasswordPrivateNewFromNative(tlsPasswordPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeTlsPasswordPrivate)
	return structGo
}
func finalizeTlsPasswordPrivate(obj *TlsPasswordPrivate) {
	tlsPasswordPrivateStruct.Free(obj.native)
}

var tlsServerConnectionInterfaceStruct *gi.Struct
var tlsServerConnectionInterfaceStruct_Once sync.Once

func tlsServerConnectionInterfaceStruct_Set() error {
	var err error
	tlsServerConnectionInterfaceStruct_Once.Do(func() {
		tlsServerConnectionInterfaceStruct, err = gi.StructNew("Gio", "TlsServerConnectionInterface")
	})
	return err
}

type TlsServerConnectionInterface struct {
	native unsafe.Pointer
}

func TlsServerConnectionInterfaceNewFromNative(native unsafe.Pointer) *TlsServerConnectionInterface {
	instance := &TlsServerConnectionInterface{native: native}

	return instance
}

/*
CastToTlsServerConnectionInterface down casts any arbitrary Object to TlsServerConnectionInterface.
Exercise care, as this is a potentially dangerous function
if the Object is not a TlsServerConnectionInterface.
*/
func (recv *TlsServerConnectionInterface) CastToTlsServerConnectionInterface(object *gobject.Object) *TlsServerConnectionInterface {
	return TlsServerConnectionInterfaceNewFromNative(object.Native())
}

func (recv *TlsServerConnectionInterface) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'g_iface' : for field getter : no Go type for 'GObject.TypeInterface'

// UNSUPPORTED : C value 'g_iface' : for field setter : no Go type for 'GObject.TypeInterface'

// TlsServerConnectionInterfaceStruct creates an uninitialised TlsServerConnectionInterface.
func TlsServerConnectionInterfaceStruct() *TlsServerConnectionInterface {
	err := tlsServerConnectionInterfaceStruct_Set()
	if err != nil {
		return nil
	}

	structGo := TlsServerConnectionInterfaceNewFromNative(tlsServerConnectionInterfaceStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeTlsServerConnectionInterface)
	return structGo
}
func finalizeTlsServerConnectionInterface(obj *TlsServerConnectionInterface) {
	tlsServerConnectionInterfaceStruct.Free(obj.native)
}

var unixConnectionClassStruct *gi.Struct
var unixConnectionClassStruct_Once sync.Once

func unixConnectionClassStruct_Set() error {
	var err error
	unixConnectionClassStruct_Once.Do(func() {
		unixConnectionClassStruct, err = gi.StructNew("Gio", "UnixConnectionClass")
	})
	return err
}

type UnixConnectionClass struct {
	native unsafe.Pointer
}

func UnixConnectionClassNewFromNative(native unsafe.Pointer) *UnixConnectionClass {
	instance := &UnixConnectionClass{native: native}

	return instance
}

/*
CastToUnixConnectionClass down casts any arbitrary Object to UnixConnectionClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a UnixConnectionClass.
*/
func (recv *UnixConnectionClass) CastToUnixConnectionClass(object *gobject.Object) *UnixConnectionClass {
	return UnixConnectionClassNewFromNative(object.Native())
}

func (recv *UnixConnectionClass) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentClass returns the C field 'parent_class'.
func (recv *UnixConnectionClass) FieldParentClass() *SocketConnectionClass {
	argValue := gi.StructFieldGet(unixConnectionClassStruct, recv.native, "parent_class")
	value := SocketConnectionClassNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *UnixConnectionClass) SetFieldParentClass(value *SocketConnectionClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.StructFieldSet(unixConnectionClassStruct, recv.native, "parent_class", argValue)
}

// UnixConnectionClassStruct creates an uninitialised UnixConnectionClass.
func UnixConnectionClassStruct() *UnixConnectionClass {
	err := unixConnectionClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := UnixConnectionClassNewFromNative(unixConnectionClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeUnixConnectionClass)
	return structGo
}
func finalizeUnixConnectionClass(obj *UnixConnectionClass) {
	unixConnectionClassStruct.Free(obj.native)
}

var unixConnectionPrivateStruct *gi.Struct
var unixConnectionPrivateStruct_Once sync.Once

func unixConnectionPrivateStruct_Set() error {
	var err error
	unixConnectionPrivateStruct_Once.Do(func() {
		unixConnectionPrivateStruct, err = gi.StructNew("Gio", "UnixConnectionPrivate")
	})
	return err
}

type UnixConnectionPrivate struct {
	native unsafe.Pointer
}

func UnixConnectionPrivateNewFromNative(native unsafe.Pointer) *UnixConnectionPrivate {
	instance := &UnixConnectionPrivate{native: native}

	return instance
}

/*
CastToUnixConnectionPrivate down casts any arbitrary Object to UnixConnectionPrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a UnixConnectionPrivate.
*/
func (recv *UnixConnectionPrivate) CastToUnixConnectionPrivate(object *gobject.Object) *UnixConnectionPrivate {
	return UnixConnectionPrivateNewFromNative(object.Native())
}

func (recv *UnixConnectionPrivate) Native() unsafe.Pointer {
	return recv.native
}

// UnixConnectionPrivateStruct creates an uninitialised UnixConnectionPrivate.
func UnixConnectionPrivateStruct() *UnixConnectionPrivate {
	err := unixConnectionPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := UnixConnectionPrivateNewFromNative(unixConnectionPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeUnixConnectionPrivate)
	return structGo
}
func finalizeUnixConnectionPrivate(obj *UnixConnectionPrivate) {
	unixConnectionPrivateStruct.Free(obj.native)
}

var unixCredentialsMessageClassStruct *gi.Struct
var unixCredentialsMessageClassStruct_Once sync.Once

func unixCredentialsMessageClassStruct_Set() error {
	var err error
	unixCredentialsMessageClassStruct_Once.Do(func() {
		unixCredentialsMessageClassStruct, err = gi.StructNew("Gio", "UnixCredentialsMessageClass")
	})
	return err
}

type UnixCredentialsMessageClass struct {
	native unsafe.Pointer
}

func UnixCredentialsMessageClassNewFromNative(native unsafe.Pointer) *UnixCredentialsMessageClass {
	instance := &UnixCredentialsMessageClass{native: native}

	return instance
}

/*
CastToUnixCredentialsMessageClass down casts any arbitrary Object to UnixCredentialsMessageClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a UnixCredentialsMessageClass.
*/
func (recv *UnixCredentialsMessageClass) CastToUnixCredentialsMessageClass(object *gobject.Object) *UnixCredentialsMessageClass {
	return UnixCredentialsMessageClassNewFromNative(object.Native())
}

func (recv *UnixCredentialsMessageClass) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentClass returns the C field 'parent_class'.
func (recv *UnixCredentialsMessageClass) FieldParentClass() *SocketControlMessageClass {
	argValue := gi.StructFieldGet(unixCredentialsMessageClassStruct, recv.native, "parent_class")
	value := SocketControlMessageClassNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *UnixCredentialsMessageClass) SetFieldParentClass(value *SocketControlMessageClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.StructFieldSet(unixCredentialsMessageClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value '_g_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved2' : for field setter : missing Type

// UnixCredentialsMessageClassStruct creates an uninitialised UnixCredentialsMessageClass.
func UnixCredentialsMessageClassStruct() *UnixCredentialsMessageClass {
	err := unixCredentialsMessageClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := UnixCredentialsMessageClassNewFromNative(unixCredentialsMessageClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeUnixCredentialsMessageClass)
	return structGo
}
func finalizeUnixCredentialsMessageClass(obj *UnixCredentialsMessageClass) {
	unixCredentialsMessageClassStruct.Free(obj.native)
}

var unixCredentialsMessagePrivateStruct *gi.Struct
var unixCredentialsMessagePrivateStruct_Once sync.Once

func unixCredentialsMessagePrivateStruct_Set() error {
	var err error
	unixCredentialsMessagePrivateStruct_Once.Do(func() {
		unixCredentialsMessagePrivateStruct, err = gi.StructNew("Gio", "UnixCredentialsMessagePrivate")
	})
	return err
}

type UnixCredentialsMessagePrivate struct {
	native unsafe.Pointer
}

func UnixCredentialsMessagePrivateNewFromNative(native unsafe.Pointer) *UnixCredentialsMessagePrivate {
	instance := &UnixCredentialsMessagePrivate{native: native}

	return instance
}

/*
CastToUnixCredentialsMessagePrivate down casts any arbitrary Object to UnixCredentialsMessagePrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a UnixCredentialsMessagePrivate.
*/
func (recv *UnixCredentialsMessagePrivate) CastToUnixCredentialsMessagePrivate(object *gobject.Object) *UnixCredentialsMessagePrivate {
	return UnixCredentialsMessagePrivateNewFromNative(object.Native())
}

func (recv *UnixCredentialsMessagePrivate) Native() unsafe.Pointer {
	return recv.native
}

// UnixCredentialsMessagePrivateStruct creates an uninitialised UnixCredentialsMessagePrivate.
func UnixCredentialsMessagePrivateStruct() *UnixCredentialsMessagePrivate {
	err := unixCredentialsMessagePrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := UnixCredentialsMessagePrivateNewFromNative(unixCredentialsMessagePrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeUnixCredentialsMessagePrivate)
	return structGo
}
func finalizeUnixCredentialsMessagePrivate(obj *UnixCredentialsMessagePrivate) {
	unixCredentialsMessagePrivateStruct.Free(obj.native)
}

var unixFDListClassStruct *gi.Struct
var unixFDListClassStruct_Once sync.Once

func unixFDListClassStruct_Set() error {
	var err error
	unixFDListClassStruct_Once.Do(func() {
		unixFDListClassStruct, err = gi.StructNew("Gio", "UnixFDListClass")
	})
	return err
}

type UnixFDListClass struct {
	native unsafe.Pointer
}

func UnixFDListClassNewFromNative(native unsafe.Pointer) *UnixFDListClass {
	instance := &UnixFDListClass{native: native}

	return instance
}

/*
CastToUnixFDListClass down casts any arbitrary Object to UnixFDListClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a UnixFDListClass.
*/
func (recv *UnixFDListClass) CastToUnixFDListClass(object *gobject.Object) *UnixFDListClass {
	return UnixFDListClassNewFromNative(object.Native())
}

func (recv *UnixFDListClass) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value '_g_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved4' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved5' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved5' : for field setter : missing Type

// UnixFDListClassStruct creates an uninitialised UnixFDListClass.
func UnixFDListClassStruct() *UnixFDListClass {
	err := unixFDListClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := UnixFDListClassNewFromNative(unixFDListClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeUnixFDListClass)
	return structGo
}
func finalizeUnixFDListClass(obj *UnixFDListClass) {
	unixFDListClassStruct.Free(obj.native)
}

var unixFDListPrivateStruct *gi.Struct
var unixFDListPrivateStruct_Once sync.Once

func unixFDListPrivateStruct_Set() error {
	var err error
	unixFDListPrivateStruct_Once.Do(func() {
		unixFDListPrivateStruct, err = gi.StructNew("Gio", "UnixFDListPrivate")
	})
	return err
}

type UnixFDListPrivate struct {
	native unsafe.Pointer
}

func UnixFDListPrivateNewFromNative(native unsafe.Pointer) *UnixFDListPrivate {
	instance := &UnixFDListPrivate{native: native}

	return instance
}

/*
CastToUnixFDListPrivate down casts any arbitrary Object to UnixFDListPrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a UnixFDListPrivate.
*/
func (recv *UnixFDListPrivate) CastToUnixFDListPrivate(object *gobject.Object) *UnixFDListPrivate {
	return UnixFDListPrivateNewFromNative(object.Native())
}

func (recv *UnixFDListPrivate) Native() unsafe.Pointer {
	return recv.native
}

// UnixFDListPrivateStruct creates an uninitialised UnixFDListPrivate.
func UnixFDListPrivateStruct() *UnixFDListPrivate {
	err := unixFDListPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := UnixFDListPrivateNewFromNative(unixFDListPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeUnixFDListPrivate)
	return structGo
}
func finalizeUnixFDListPrivate(obj *UnixFDListPrivate) {
	unixFDListPrivateStruct.Free(obj.native)
}

var unixFDMessageClassStruct *gi.Struct
var unixFDMessageClassStruct_Once sync.Once

func unixFDMessageClassStruct_Set() error {
	var err error
	unixFDMessageClassStruct_Once.Do(func() {
		unixFDMessageClassStruct, err = gi.StructNew("Gio", "UnixFDMessageClass")
	})
	return err
}

type UnixFDMessageClass struct {
	native unsafe.Pointer
}

func UnixFDMessageClassNewFromNative(native unsafe.Pointer) *UnixFDMessageClass {
	instance := &UnixFDMessageClass{native: native}

	return instance
}

/*
CastToUnixFDMessageClass down casts any arbitrary Object to UnixFDMessageClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a UnixFDMessageClass.
*/
func (recv *UnixFDMessageClass) CastToUnixFDMessageClass(object *gobject.Object) *UnixFDMessageClass {
	return UnixFDMessageClassNewFromNative(object.Native())
}

func (recv *UnixFDMessageClass) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentClass returns the C field 'parent_class'.
func (recv *UnixFDMessageClass) FieldParentClass() *SocketControlMessageClass {
	argValue := gi.StructFieldGet(unixFDMessageClassStruct, recv.native, "parent_class")
	value := SocketControlMessageClassNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *UnixFDMessageClass) SetFieldParentClass(value *SocketControlMessageClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.StructFieldSet(unixFDMessageClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value '_g_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved2' : for field setter : missing Type

// UnixFDMessageClassStruct creates an uninitialised UnixFDMessageClass.
func UnixFDMessageClassStruct() *UnixFDMessageClass {
	err := unixFDMessageClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := UnixFDMessageClassNewFromNative(unixFDMessageClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeUnixFDMessageClass)
	return structGo
}
func finalizeUnixFDMessageClass(obj *UnixFDMessageClass) {
	unixFDMessageClassStruct.Free(obj.native)
}

var unixFDMessagePrivateStruct *gi.Struct
var unixFDMessagePrivateStruct_Once sync.Once

func unixFDMessagePrivateStruct_Set() error {
	var err error
	unixFDMessagePrivateStruct_Once.Do(func() {
		unixFDMessagePrivateStruct, err = gi.StructNew("Gio", "UnixFDMessagePrivate")
	})
	return err
}

type UnixFDMessagePrivate struct {
	native unsafe.Pointer
}

func UnixFDMessagePrivateNewFromNative(native unsafe.Pointer) *UnixFDMessagePrivate {
	instance := &UnixFDMessagePrivate{native: native}

	return instance
}

/*
CastToUnixFDMessagePrivate down casts any arbitrary Object to UnixFDMessagePrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a UnixFDMessagePrivate.
*/
func (recv *UnixFDMessagePrivate) CastToUnixFDMessagePrivate(object *gobject.Object) *UnixFDMessagePrivate {
	return UnixFDMessagePrivateNewFromNative(object.Native())
}

func (recv *UnixFDMessagePrivate) Native() unsafe.Pointer {
	return recv.native
}

// UnixFDMessagePrivateStruct creates an uninitialised UnixFDMessagePrivate.
func UnixFDMessagePrivateStruct() *UnixFDMessagePrivate {
	err := unixFDMessagePrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := UnixFDMessagePrivateNewFromNative(unixFDMessagePrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeUnixFDMessagePrivate)
	return structGo
}
func finalizeUnixFDMessagePrivate(obj *UnixFDMessagePrivate) {
	unixFDMessagePrivateStruct.Free(obj.native)
}

var unixInputStreamClassStruct *gi.Struct
var unixInputStreamClassStruct_Once sync.Once

func unixInputStreamClassStruct_Set() error {
	var err error
	unixInputStreamClassStruct_Once.Do(func() {
		unixInputStreamClassStruct, err = gi.StructNew("Gio", "UnixInputStreamClass")
	})
	return err
}

type UnixInputStreamClass struct {
	native unsafe.Pointer
}

func UnixInputStreamClassNewFromNative(native unsafe.Pointer) *UnixInputStreamClass {
	instance := &UnixInputStreamClass{native: native}

	return instance
}

/*
CastToUnixInputStreamClass down casts any arbitrary Object to UnixInputStreamClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a UnixInputStreamClass.
*/
func (recv *UnixInputStreamClass) CastToUnixInputStreamClass(object *gobject.Object) *UnixInputStreamClass {
	return UnixInputStreamClassNewFromNative(object.Native())
}

func (recv *UnixInputStreamClass) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentClass returns the C field 'parent_class'.
func (recv *UnixInputStreamClass) FieldParentClass() *InputStreamClass {
	argValue := gi.StructFieldGet(unixInputStreamClassStruct, recv.native, "parent_class")
	value := InputStreamClassNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *UnixInputStreamClass) SetFieldParentClass(value *InputStreamClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.StructFieldSet(unixInputStreamClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value '_g_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved4' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved5' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved5' : for field setter : missing Type

// UnixInputStreamClassStruct creates an uninitialised UnixInputStreamClass.
func UnixInputStreamClassStruct() *UnixInputStreamClass {
	err := unixInputStreamClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := UnixInputStreamClassNewFromNative(unixInputStreamClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeUnixInputStreamClass)
	return structGo
}
func finalizeUnixInputStreamClass(obj *UnixInputStreamClass) {
	unixInputStreamClassStruct.Free(obj.native)
}

var unixInputStreamPrivateStruct *gi.Struct
var unixInputStreamPrivateStruct_Once sync.Once

func unixInputStreamPrivateStruct_Set() error {
	var err error
	unixInputStreamPrivateStruct_Once.Do(func() {
		unixInputStreamPrivateStruct, err = gi.StructNew("Gio", "UnixInputStreamPrivate")
	})
	return err
}

type UnixInputStreamPrivate struct {
	native unsafe.Pointer
}

func UnixInputStreamPrivateNewFromNative(native unsafe.Pointer) *UnixInputStreamPrivate {
	instance := &UnixInputStreamPrivate{native: native}

	return instance
}

/*
CastToUnixInputStreamPrivate down casts any arbitrary Object to UnixInputStreamPrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a UnixInputStreamPrivate.
*/
func (recv *UnixInputStreamPrivate) CastToUnixInputStreamPrivate(object *gobject.Object) *UnixInputStreamPrivate {
	return UnixInputStreamPrivateNewFromNative(object.Native())
}

func (recv *UnixInputStreamPrivate) Native() unsafe.Pointer {
	return recv.native
}

// UnixInputStreamPrivateStruct creates an uninitialised UnixInputStreamPrivate.
func UnixInputStreamPrivateStruct() *UnixInputStreamPrivate {
	err := unixInputStreamPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := UnixInputStreamPrivateNewFromNative(unixInputStreamPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeUnixInputStreamPrivate)
	return structGo
}
func finalizeUnixInputStreamPrivate(obj *UnixInputStreamPrivate) {
	unixInputStreamPrivateStruct.Free(obj.native)
}

var unixMountEntryStruct *gi.Struct
var unixMountEntryStruct_Once sync.Once

func unixMountEntryStruct_Set() error {
	var err error
	unixMountEntryStruct_Once.Do(func() {
		unixMountEntryStruct, err = gi.StructNew("Gio", "UnixMountEntry")
	})
	return err
}

type UnixMountEntry struct {
	native unsafe.Pointer
}

func UnixMountEntryNewFromNative(native unsafe.Pointer) *UnixMountEntry {
	instance := &UnixMountEntry{native: native}

	return instance
}

/*
CastToUnixMountEntry down casts any arbitrary Object to UnixMountEntry.
Exercise care, as this is a potentially dangerous function
if the Object is not a UnixMountEntry.
*/
func (recv *UnixMountEntry) CastToUnixMountEntry(object *gobject.Object) *UnixMountEntry {
	return UnixMountEntryNewFromNative(object.Native())
}

func (recv *UnixMountEntry) Native() unsafe.Pointer {
	return recv.native
}

// UnixMountEntryStruct creates an uninitialised UnixMountEntry.
func UnixMountEntryStruct() *UnixMountEntry {
	err := unixMountEntryStruct_Set()
	if err != nil {
		return nil
	}

	structGo := UnixMountEntryNewFromNative(unixMountEntryStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeUnixMountEntry)
	return structGo
}
func finalizeUnixMountEntry(obj *UnixMountEntry) {
	unixMountEntryStruct.Free(obj.native)
}

var unixMountMonitorClassStruct *gi.Struct
var unixMountMonitorClassStruct_Once sync.Once

func unixMountMonitorClassStruct_Set() error {
	var err error
	unixMountMonitorClassStruct_Once.Do(func() {
		unixMountMonitorClassStruct, err = gi.StructNew("Gio", "UnixMountMonitorClass")
	})
	return err
}

type UnixMountMonitorClass struct {
	native unsafe.Pointer
}

func UnixMountMonitorClassNewFromNative(native unsafe.Pointer) *UnixMountMonitorClass {
	instance := &UnixMountMonitorClass{native: native}

	return instance
}

/*
CastToUnixMountMonitorClass down casts any arbitrary Object to UnixMountMonitorClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a UnixMountMonitorClass.
*/
func (recv *UnixMountMonitorClass) CastToUnixMountMonitorClass(object *gobject.Object) *UnixMountMonitorClass {
	return UnixMountMonitorClassNewFromNative(object.Native())
}

func (recv *UnixMountMonitorClass) Native() unsafe.Pointer {
	return recv.native
}

// UnixMountMonitorClassStruct creates an uninitialised UnixMountMonitorClass.
func UnixMountMonitorClassStruct() *UnixMountMonitorClass {
	err := unixMountMonitorClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := UnixMountMonitorClassNewFromNative(unixMountMonitorClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeUnixMountMonitorClass)
	return structGo
}
func finalizeUnixMountMonitorClass(obj *UnixMountMonitorClass) {
	unixMountMonitorClassStruct.Free(obj.native)
}

var unixMountPointStruct *gi.Struct
var unixMountPointStruct_Once sync.Once

func unixMountPointStruct_Set() error {
	var err error
	unixMountPointStruct_Once.Do(func() {
		unixMountPointStruct, err = gi.StructNew("Gio", "UnixMountPoint")
	})
	return err
}

type UnixMountPoint struct {
	native unsafe.Pointer
}

func UnixMountPointNewFromNative(native unsafe.Pointer) *UnixMountPoint {
	instance := &UnixMountPoint{native: native}

	return instance
}

/*
CastToUnixMountPoint down casts any arbitrary Object to UnixMountPoint.
Exercise care, as this is a potentially dangerous function
if the Object is not a UnixMountPoint.
*/
func (recv *UnixMountPoint) CastToUnixMountPoint(object *gobject.Object) *UnixMountPoint {
	return UnixMountPointNewFromNative(object.Native())
}

func (recv *UnixMountPoint) Native() unsafe.Pointer {
	return recv.native
}

var unixMountPointCompareFunction *gi.Function
var unixMountPointCompareFunction_Once sync.Once

func unixMountPointCompareFunction_Set() error {
	var err error
	unixMountPointCompareFunction_Once.Do(func() {
		err = unixMountPointStruct_Set()
		if err != nil {
			return
		}
		unixMountPointCompareFunction, err = unixMountPointStruct.InvokerNew("compare")
	})
	return err
}

// Compare is a representation of the C type g_unix_mount_point_compare.
func (recv *UnixMountPoint) Compare(mount2 *UnixMountPoint) int32 {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(mount2.native)

	var ret gi.Argument

	err := unixMountPointCompareFunction_Set()
	if err == nil {
		ret = unixMountPointCompareFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var unixMountPointCopyFunction *gi.Function
var unixMountPointCopyFunction_Once sync.Once

func unixMountPointCopyFunction_Set() error {
	var err error
	unixMountPointCopyFunction_Once.Do(func() {
		err = unixMountPointStruct_Set()
		if err != nil {
			return
		}
		unixMountPointCopyFunction, err = unixMountPointStruct.InvokerNew("copy")
	})
	return err
}

// Copy is a representation of the C type g_unix_mount_point_copy.
func (recv *UnixMountPoint) Copy() *UnixMountPoint {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := unixMountPointCopyFunction_Set()
	if err == nil {
		ret = unixMountPointCopyFunction.Invoke(inArgs[:], nil)
	}

	retGo := UnixMountPointNewFromNative(ret.Pointer())

	return retGo
}

var unixMountPointFreeFunction *gi.Function
var unixMountPointFreeFunction_Once sync.Once

func unixMountPointFreeFunction_Set() error {
	var err error
	unixMountPointFreeFunction_Once.Do(func() {
		err = unixMountPointStruct_Set()
		if err != nil {
			return
		}
		unixMountPointFreeFunction, err = unixMountPointStruct.InvokerNew("free")
	})
	return err
}

// Free is a representation of the C type g_unix_mount_point_free.
func (recv *UnixMountPoint) Free() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := unixMountPointFreeFunction_Set()
	if err == nil {
		unixMountPointFreeFunction.Invoke(inArgs[:], nil)
	}

	return
}

var unixMountPointGetDevicePathFunction *gi.Function
var unixMountPointGetDevicePathFunction_Once sync.Once

func unixMountPointGetDevicePathFunction_Set() error {
	var err error
	unixMountPointGetDevicePathFunction_Once.Do(func() {
		err = unixMountPointStruct_Set()
		if err != nil {
			return
		}
		unixMountPointGetDevicePathFunction, err = unixMountPointStruct.InvokerNew("get_device_path")
	})
	return err
}

// GetDevicePath is a representation of the C type g_unix_mount_point_get_device_path.
func (recv *UnixMountPoint) GetDevicePath() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := unixMountPointGetDevicePathFunction_Set()
	if err == nil {
		ret = unixMountPointGetDevicePathFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var unixMountPointGetFsTypeFunction *gi.Function
var unixMountPointGetFsTypeFunction_Once sync.Once

func unixMountPointGetFsTypeFunction_Set() error {
	var err error
	unixMountPointGetFsTypeFunction_Once.Do(func() {
		err = unixMountPointStruct_Set()
		if err != nil {
			return
		}
		unixMountPointGetFsTypeFunction, err = unixMountPointStruct.InvokerNew("get_fs_type")
	})
	return err
}

// GetFsType is a representation of the C type g_unix_mount_point_get_fs_type.
func (recv *UnixMountPoint) GetFsType() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := unixMountPointGetFsTypeFunction_Set()
	if err == nil {
		ret = unixMountPointGetFsTypeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var unixMountPointGetMountPathFunction *gi.Function
var unixMountPointGetMountPathFunction_Once sync.Once

func unixMountPointGetMountPathFunction_Set() error {
	var err error
	unixMountPointGetMountPathFunction_Once.Do(func() {
		err = unixMountPointStruct_Set()
		if err != nil {
			return
		}
		unixMountPointGetMountPathFunction, err = unixMountPointStruct.InvokerNew("get_mount_path")
	})
	return err
}

// GetMountPath is a representation of the C type g_unix_mount_point_get_mount_path.
func (recv *UnixMountPoint) GetMountPath() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := unixMountPointGetMountPathFunction_Set()
	if err == nil {
		ret = unixMountPointGetMountPathFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var unixMountPointGetOptionsFunction *gi.Function
var unixMountPointGetOptionsFunction_Once sync.Once

func unixMountPointGetOptionsFunction_Set() error {
	var err error
	unixMountPointGetOptionsFunction_Once.Do(func() {
		err = unixMountPointStruct_Set()
		if err != nil {
			return
		}
		unixMountPointGetOptionsFunction, err = unixMountPointStruct.InvokerNew("get_options")
	})
	return err
}

// GetOptions is a representation of the C type g_unix_mount_point_get_options.
func (recv *UnixMountPoint) GetOptions() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := unixMountPointGetOptionsFunction_Set()
	if err == nil {
		ret = unixMountPointGetOptionsFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var unixMountPointGuessCanEjectFunction *gi.Function
var unixMountPointGuessCanEjectFunction_Once sync.Once

func unixMountPointGuessCanEjectFunction_Set() error {
	var err error
	unixMountPointGuessCanEjectFunction_Once.Do(func() {
		err = unixMountPointStruct_Set()
		if err != nil {
			return
		}
		unixMountPointGuessCanEjectFunction, err = unixMountPointStruct.InvokerNew("guess_can_eject")
	})
	return err
}

// GuessCanEject is a representation of the C type g_unix_mount_point_guess_can_eject.
func (recv *UnixMountPoint) GuessCanEject() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := unixMountPointGuessCanEjectFunction_Set()
	if err == nil {
		ret = unixMountPointGuessCanEjectFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'g_unix_mount_point_guess_icon' : return type 'Icon' not supported

var unixMountPointGuessNameFunction *gi.Function
var unixMountPointGuessNameFunction_Once sync.Once

func unixMountPointGuessNameFunction_Set() error {
	var err error
	unixMountPointGuessNameFunction_Once.Do(func() {
		err = unixMountPointStruct_Set()
		if err != nil {
			return
		}
		unixMountPointGuessNameFunction, err = unixMountPointStruct.InvokerNew("guess_name")
	})
	return err
}

// GuessName is a representation of the C type g_unix_mount_point_guess_name.
func (recv *UnixMountPoint) GuessName() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := unixMountPointGuessNameFunction_Set()
	if err == nil {
		ret = unixMountPointGuessNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

// UNSUPPORTED : C value 'g_unix_mount_point_guess_symbolic_icon' : return type 'Icon' not supported

var unixMountPointIsLoopbackFunction *gi.Function
var unixMountPointIsLoopbackFunction_Once sync.Once

func unixMountPointIsLoopbackFunction_Set() error {
	var err error
	unixMountPointIsLoopbackFunction_Once.Do(func() {
		err = unixMountPointStruct_Set()
		if err != nil {
			return
		}
		unixMountPointIsLoopbackFunction, err = unixMountPointStruct.InvokerNew("is_loopback")
	})
	return err
}

// IsLoopback is a representation of the C type g_unix_mount_point_is_loopback.
func (recv *UnixMountPoint) IsLoopback() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := unixMountPointIsLoopbackFunction_Set()
	if err == nil {
		ret = unixMountPointIsLoopbackFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var unixMountPointIsReadonlyFunction *gi.Function
var unixMountPointIsReadonlyFunction_Once sync.Once

func unixMountPointIsReadonlyFunction_Set() error {
	var err error
	unixMountPointIsReadonlyFunction_Once.Do(func() {
		err = unixMountPointStruct_Set()
		if err != nil {
			return
		}
		unixMountPointIsReadonlyFunction, err = unixMountPointStruct.InvokerNew("is_readonly")
	})
	return err
}

// IsReadonly is a representation of the C type g_unix_mount_point_is_readonly.
func (recv *UnixMountPoint) IsReadonly() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := unixMountPointIsReadonlyFunction_Set()
	if err == nil {
		ret = unixMountPointIsReadonlyFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var unixMountPointIsUserMountableFunction *gi.Function
var unixMountPointIsUserMountableFunction_Once sync.Once

func unixMountPointIsUserMountableFunction_Set() error {
	var err error
	unixMountPointIsUserMountableFunction_Once.Do(func() {
		err = unixMountPointStruct_Set()
		if err != nil {
			return
		}
		unixMountPointIsUserMountableFunction, err = unixMountPointStruct.InvokerNew("is_user_mountable")
	})
	return err
}

// IsUserMountable is a representation of the C type g_unix_mount_point_is_user_mountable.
func (recv *UnixMountPoint) IsUserMountable() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := unixMountPointIsUserMountableFunction_Set()
	if err == nil {
		ret = unixMountPointIsUserMountableFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UnixMountPointStruct creates an uninitialised UnixMountPoint.
func UnixMountPointStruct() *UnixMountPoint {
	err := unixMountPointStruct_Set()
	if err != nil {
		return nil
	}

	structGo := UnixMountPointNewFromNative(unixMountPointStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeUnixMountPoint)
	return structGo
}
func finalizeUnixMountPoint(obj *UnixMountPoint) {
	unixMountPointStruct.Free(obj.native)
}

var unixOutputStreamClassStruct *gi.Struct
var unixOutputStreamClassStruct_Once sync.Once

func unixOutputStreamClassStruct_Set() error {
	var err error
	unixOutputStreamClassStruct_Once.Do(func() {
		unixOutputStreamClassStruct, err = gi.StructNew("Gio", "UnixOutputStreamClass")
	})
	return err
}

type UnixOutputStreamClass struct {
	native unsafe.Pointer
}

func UnixOutputStreamClassNewFromNative(native unsafe.Pointer) *UnixOutputStreamClass {
	instance := &UnixOutputStreamClass{native: native}

	return instance
}

/*
CastToUnixOutputStreamClass down casts any arbitrary Object to UnixOutputStreamClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a UnixOutputStreamClass.
*/
func (recv *UnixOutputStreamClass) CastToUnixOutputStreamClass(object *gobject.Object) *UnixOutputStreamClass {
	return UnixOutputStreamClassNewFromNative(object.Native())
}

func (recv *UnixOutputStreamClass) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentClass returns the C field 'parent_class'.
func (recv *UnixOutputStreamClass) FieldParentClass() *OutputStreamClass {
	argValue := gi.StructFieldGet(unixOutputStreamClassStruct, recv.native, "parent_class")
	value := OutputStreamClassNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *UnixOutputStreamClass) SetFieldParentClass(value *OutputStreamClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.StructFieldSet(unixOutputStreamClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value '_g_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved4' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved5' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved5' : for field setter : missing Type

// UnixOutputStreamClassStruct creates an uninitialised UnixOutputStreamClass.
func UnixOutputStreamClassStruct() *UnixOutputStreamClass {
	err := unixOutputStreamClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := UnixOutputStreamClassNewFromNative(unixOutputStreamClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeUnixOutputStreamClass)
	return structGo
}
func finalizeUnixOutputStreamClass(obj *UnixOutputStreamClass) {
	unixOutputStreamClassStruct.Free(obj.native)
}

var unixOutputStreamPrivateStruct *gi.Struct
var unixOutputStreamPrivateStruct_Once sync.Once

func unixOutputStreamPrivateStruct_Set() error {
	var err error
	unixOutputStreamPrivateStruct_Once.Do(func() {
		unixOutputStreamPrivateStruct, err = gi.StructNew("Gio", "UnixOutputStreamPrivate")
	})
	return err
}

type UnixOutputStreamPrivate struct {
	native unsafe.Pointer
}

func UnixOutputStreamPrivateNewFromNative(native unsafe.Pointer) *UnixOutputStreamPrivate {
	instance := &UnixOutputStreamPrivate{native: native}

	return instance
}

/*
CastToUnixOutputStreamPrivate down casts any arbitrary Object to UnixOutputStreamPrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a UnixOutputStreamPrivate.
*/
func (recv *UnixOutputStreamPrivate) CastToUnixOutputStreamPrivate(object *gobject.Object) *UnixOutputStreamPrivate {
	return UnixOutputStreamPrivateNewFromNative(object.Native())
}

func (recv *UnixOutputStreamPrivate) Native() unsafe.Pointer {
	return recv.native
}

// UnixOutputStreamPrivateStruct creates an uninitialised UnixOutputStreamPrivate.
func UnixOutputStreamPrivateStruct() *UnixOutputStreamPrivate {
	err := unixOutputStreamPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := UnixOutputStreamPrivateNewFromNative(unixOutputStreamPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeUnixOutputStreamPrivate)
	return structGo
}
func finalizeUnixOutputStreamPrivate(obj *UnixOutputStreamPrivate) {
	unixOutputStreamPrivateStruct.Free(obj.native)
}

var unixSocketAddressClassStruct *gi.Struct
var unixSocketAddressClassStruct_Once sync.Once

func unixSocketAddressClassStruct_Set() error {
	var err error
	unixSocketAddressClassStruct_Once.Do(func() {
		unixSocketAddressClassStruct, err = gi.StructNew("Gio", "UnixSocketAddressClass")
	})
	return err
}

type UnixSocketAddressClass struct {
	native unsafe.Pointer
}

func UnixSocketAddressClassNewFromNative(native unsafe.Pointer) *UnixSocketAddressClass {
	instance := &UnixSocketAddressClass{native: native}

	return instance
}

/*
CastToUnixSocketAddressClass down casts any arbitrary Object to UnixSocketAddressClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a UnixSocketAddressClass.
*/
func (recv *UnixSocketAddressClass) CastToUnixSocketAddressClass(object *gobject.Object) *UnixSocketAddressClass {
	return UnixSocketAddressClassNewFromNative(object.Native())
}

func (recv *UnixSocketAddressClass) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentClass returns the C field 'parent_class'.
func (recv *UnixSocketAddressClass) FieldParentClass() *SocketAddressClass {
	argValue := gi.StructFieldGet(unixSocketAddressClassStruct, recv.native, "parent_class")
	value := SocketAddressClassNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *UnixSocketAddressClass) SetFieldParentClass(value *SocketAddressClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.StructFieldSet(unixSocketAddressClassStruct, recv.native, "parent_class", argValue)
}

// UnixSocketAddressClassStruct creates an uninitialised UnixSocketAddressClass.
func UnixSocketAddressClassStruct() *UnixSocketAddressClass {
	err := unixSocketAddressClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := UnixSocketAddressClassNewFromNative(unixSocketAddressClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeUnixSocketAddressClass)
	return structGo
}
func finalizeUnixSocketAddressClass(obj *UnixSocketAddressClass) {
	unixSocketAddressClassStruct.Free(obj.native)
}

var unixSocketAddressPrivateStruct *gi.Struct
var unixSocketAddressPrivateStruct_Once sync.Once

func unixSocketAddressPrivateStruct_Set() error {
	var err error
	unixSocketAddressPrivateStruct_Once.Do(func() {
		unixSocketAddressPrivateStruct, err = gi.StructNew("Gio", "UnixSocketAddressPrivate")
	})
	return err
}

type UnixSocketAddressPrivate struct {
	native unsafe.Pointer
}

func UnixSocketAddressPrivateNewFromNative(native unsafe.Pointer) *UnixSocketAddressPrivate {
	instance := &UnixSocketAddressPrivate{native: native}

	return instance
}

/*
CastToUnixSocketAddressPrivate down casts any arbitrary Object to UnixSocketAddressPrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a UnixSocketAddressPrivate.
*/
func (recv *UnixSocketAddressPrivate) CastToUnixSocketAddressPrivate(object *gobject.Object) *UnixSocketAddressPrivate {
	return UnixSocketAddressPrivateNewFromNative(object.Native())
}

func (recv *UnixSocketAddressPrivate) Native() unsafe.Pointer {
	return recv.native
}

// UnixSocketAddressPrivateStruct creates an uninitialised UnixSocketAddressPrivate.
func UnixSocketAddressPrivateStruct() *UnixSocketAddressPrivate {
	err := unixSocketAddressPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := UnixSocketAddressPrivateNewFromNative(unixSocketAddressPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeUnixSocketAddressPrivate)
	return structGo
}
func finalizeUnixSocketAddressPrivate(obj *UnixSocketAddressPrivate) {
	unixSocketAddressPrivateStruct.Free(obj.native)
}

var vfsClassStruct *gi.Struct
var vfsClassStruct_Once sync.Once

func vfsClassStruct_Set() error {
	var err error
	vfsClassStruct_Once.Do(func() {
		vfsClassStruct, err = gi.StructNew("Gio", "VfsClass")
	})
	return err
}

type VfsClass struct {
	native unsafe.Pointer
}

func VfsClassNewFromNative(native unsafe.Pointer) *VfsClass {
	instance := &VfsClass{native: native}

	return instance
}

/*
CastToVfsClass down casts any arbitrary Object to VfsClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a VfsClass.
*/
func (recv *VfsClass) CastToVfsClass(object *gobject.Object) *VfsClass {
	return VfsClassNewFromNative(object.Native())
}

func (recv *VfsClass) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'is_active' : for field getter : missing Type

// UNSUPPORTED : C value 'is_active' : for field setter : missing Type

// UNSUPPORTED : C value 'get_file_for_path' : for field getter : missing Type

// UNSUPPORTED : C value 'get_file_for_path' : for field setter : missing Type

// UNSUPPORTED : C value 'get_file_for_uri' : for field getter : missing Type

// UNSUPPORTED : C value 'get_file_for_uri' : for field setter : missing Type

// UNSUPPORTED : C value 'get_supported_uri_schemes' : for field getter : missing Type

// UNSUPPORTED : C value 'get_supported_uri_schemes' : for field setter : missing Type

// UNSUPPORTED : C value 'parse_name' : for field getter : missing Type

// UNSUPPORTED : C value 'parse_name' : for field setter : missing Type

// UNSUPPORTED : C value 'local_file_add_info' : for field getter : missing Type

// UNSUPPORTED : C value 'local_file_add_info' : for field setter : missing Type

// UNSUPPORTED : C value 'add_writable_namespaces' : for field getter : missing Type

// UNSUPPORTED : C value 'add_writable_namespaces' : for field setter : missing Type

// UNSUPPORTED : C value 'local_file_set_attributes' : for field getter : missing Type

// UNSUPPORTED : C value 'local_file_set_attributes' : for field setter : missing Type

// UNSUPPORTED : C value 'local_file_removed' : for field getter : missing Type

// UNSUPPORTED : C value 'local_file_removed' : for field setter : missing Type

// UNSUPPORTED : C value 'local_file_moved' : for field getter : missing Type

// UNSUPPORTED : C value 'local_file_moved' : for field setter : missing Type

// UNSUPPORTED : C value 'deserialize_icon' : for field getter : missing Type

// UNSUPPORTED : C value 'deserialize_icon' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved4' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved5' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved5' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved6' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved6' : for field setter : missing Type

// VfsClassStruct creates an uninitialised VfsClass.
func VfsClassStruct() *VfsClass {
	err := vfsClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := VfsClassNewFromNative(vfsClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeVfsClass)
	return structGo
}
func finalizeVfsClass(obj *VfsClass) {
	vfsClassStruct.Free(obj.native)
}

var volumeIfaceStruct *gi.Struct
var volumeIfaceStruct_Once sync.Once

func volumeIfaceStruct_Set() error {
	var err error
	volumeIfaceStruct_Once.Do(func() {
		volumeIfaceStruct, err = gi.StructNew("Gio", "VolumeIface")
	})
	return err
}

type VolumeIface struct {
	native unsafe.Pointer
}

func VolumeIfaceNewFromNative(native unsafe.Pointer) *VolumeIface {
	instance := &VolumeIface{native: native}

	return instance
}

/*
CastToVolumeIface down casts any arbitrary Object to VolumeIface.
Exercise care, as this is a potentially dangerous function
if the Object is not a VolumeIface.
*/
func (recv *VolumeIface) CastToVolumeIface(object *gobject.Object) *VolumeIface {
	return VolumeIfaceNewFromNative(object.Native())
}

func (recv *VolumeIface) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'g_iface' : for field getter : no Go type for 'GObject.TypeInterface'

// UNSUPPORTED : C value 'g_iface' : for field setter : no Go type for 'GObject.TypeInterface'

// UNSUPPORTED : C value 'changed' : for field getter : missing Type

// UNSUPPORTED : C value 'changed' : for field setter : missing Type

// UNSUPPORTED : C value 'removed' : for field getter : missing Type

// UNSUPPORTED : C value 'removed' : for field setter : missing Type

// UNSUPPORTED : C value 'get_name' : for field getter : missing Type

// UNSUPPORTED : C value 'get_name' : for field setter : missing Type

// UNSUPPORTED : C value 'get_icon' : for field getter : missing Type

// UNSUPPORTED : C value 'get_icon' : for field setter : missing Type

// UNSUPPORTED : C value 'get_uuid' : for field getter : missing Type

// UNSUPPORTED : C value 'get_uuid' : for field setter : missing Type

// UNSUPPORTED : C value 'get_drive' : for field getter : missing Type

// UNSUPPORTED : C value 'get_drive' : for field setter : missing Type

// UNSUPPORTED : C value 'get_mount' : for field getter : missing Type

// UNSUPPORTED : C value 'get_mount' : for field setter : missing Type

// UNSUPPORTED : C value 'can_mount' : for field getter : missing Type

// UNSUPPORTED : C value 'can_mount' : for field setter : missing Type

// UNSUPPORTED : C value 'can_eject' : for field getter : missing Type

// UNSUPPORTED : C value 'can_eject' : for field setter : missing Type

// UNSUPPORTED : C value 'mount_fn' : for field getter : missing Type

// UNSUPPORTED : C value 'mount_fn' : for field setter : missing Type

// UNSUPPORTED : C value 'mount_finish' : for field getter : missing Type

// UNSUPPORTED : C value 'mount_finish' : for field setter : missing Type

// UNSUPPORTED : C value 'eject' : for field getter : missing Type

// UNSUPPORTED : C value 'eject' : for field setter : missing Type

// UNSUPPORTED : C value 'eject_finish' : for field getter : missing Type

// UNSUPPORTED : C value 'eject_finish' : for field setter : missing Type

// UNSUPPORTED : C value 'get_identifier' : for field getter : missing Type

// UNSUPPORTED : C value 'get_identifier' : for field setter : missing Type

// UNSUPPORTED : C value 'enumerate_identifiers' : for field getter : missing Type

// UNSUPPORTED : C value 'enumerate_identifiers' : for field setter : missing Type

// UNSUPPORTED : C value 'should_automount' : for field getter : missing Type

// UNSUPPORTED : C value 'should_automount' : for field setter : missing Type

// UNSUPPORTED : C value 'get_activation_root' : for field getter : missing Type

// UNSUPPORTED : C value 'get_activation_root' : for field setter : missing Type

// UNSUPPORTED : C value 'eject_with_operation' : for field getter : missing Type

// UNSUPPORTED : C value 'eject_with_operation' : for field setter : missing Type

// UNSUPPORTED : C value 'eject_with_operation_finish' : for field getter : missing Type

// UNSUPPORTED : C value 'eject_with_operation_finish' : for field setter : missing Type

// UNSUPPORTED : C value 'get_sort_key' : for field getter : missing Type

// UNSUPPORTED : C value 'get_sort_key' : for field setter : missing Type

// UNSUPPORTED : C value 'get_symbolic_icon' : for field getter : missing Type

// UNSUPPORTED : C value 'get_symbolic_icon' : for field setter : missing Type

// VolumeIfaceStruct creates an uninitialised VolumeIface.
func VolumeIfaceStruct() *VolumeIface {
	err := volumeIfaceStruct_Set()
	if err != nil {
		return nil
	}

	structGo := VolumeIfaceNewFromNative(volumeIfaceStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeVolumeIface)
	return structGo
}
func finalizeVolumeIface(obj *VolumeIface) {
	volumeIfaceStruct.Free(obj.native)
}

var volumeMonitorClassStruct *gi.Struct
var volumeMonitorClassStruct_Once sync.Once

func volumeMonitorClassStruct_Set() error {
	var err error
	volumeMonitorClassStruct_Once.Do(func() {
		volumeMonitorClassStruct, err = gi.StructNew("Gio", "VolumeMonitorClass")
	})
	return err
}

type VolumeMonitorClass struct {
	native unsafe.Pointer
}

func VolumeMonitorClassNewFromNative(native unsafe.Pointer) *VolumeMonitorClass {
	instance := &VolumeMonitorClass{native: native}

	return instance
}

/*
CastToVolumeMonitorClass down casts any arbitrary Object to VolumeMonitorClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a VolumeMonitorClass.
*/
func (recv *VolumeMonitorClass) CastToVolumeMonitorClass(object *gobject.Object) *VolumeMonitorClass {
	return VolumeMonitorClassNewFromNative(object.Native())
}

func (recv *VolumeMonitorClass) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'volume_added' : for field getter : missing Type

// UNSUPPORTED : C value 'volume_added' : for field setter : missing Type

// UNSUPPORTED : C value 'volume_removed' : for field getter : missing Type

// UNSUPPORTED : C value 'volume_removed' : for field setter : missing Type

// UNSUPPORTED : C value 'volume_changed' : for field getter : missing Type

// UNSUPPORTED : C value 'volume_changed' : for field setter : missing Type

// UNSUPPORTED : C value 'mount_added' : for field getter : missing Type

// UNSUPPORTED : C value 'mount_added' : for field setter : missing Type

// UNSUPPORTED : C value 'mount_removed' : for field getter : missing Type

// UNSUPPORTED : C value 'mount_removed' : for field setter : missing Type

// UNSUPPORTED : C value 'mount_pre_unmount' : for field getter : missing Type

// UNSUPPORTED : C value 'mount_pre_unmount' : for field setter : missing Type

// UNSUPPORTED : C value 'mount_changed' : for field getter : missing Type

// UNSUPPORTED : C value 'mount_changed' : for field setter : missing Type

// UNSUPPORTED : C value 'drive_connected' : for field getter : missing Type

// UNSUPPORTED : C value 'drive_connected' : for field setter : missing Type

// UNSUPPORTED : C value 'drive_disconnected' : for field getter : missing Type

// UNSUPPORTED : C value 'drive_disconnected' : for field setter : missing Type

// UNSUPPORTED : C value 'drive_changed' : for field getter : missing Type

// UNSUPPORTED : C value 'drive_changed' : for field setter : missing Type

// UNSUPPORTED : C value 'is_supported' : for field getter : missing Type

// UNSUPPORTED : C value 'is_supported' : for field setter : missing Type

// UNSUPPORTED : C value 'get_connected_drives' : for field getter : missing Type

// UNSUPPORTED : C value 'get_connected_drives' : for field setter : missing Type

// UNSUPPORTED : C value 'get_volumes' : for field getter : missing Type

// UNSUPPORTED : C value 'get_volumes' : for field setter : missing Type

// UNSUPPORTED : C value 'get_mounts' : for field getter : missing Type

// UNSUPPORTED : C value 'get_mounts' : for field setter : missing Type

// UNSUPPORTED : C value 'get_volume_for_uuid' : for field getter : missing Type

// UNSUPPORTED : C value 'get_volume_for_uuid' : for field setter : missing Type

// UNSUPPORTED : C value 'get_mount_for_uuid' : for field getter : missing Type

// UNSUPPORTED : C value 'get_mount_for_uuid' : for field setter : missing Type

// UNSUPPORTED : C value 'adopt_orphan_mount' : for field getter : missing Type

// UNSUPPORTED : C value 'adopt_orphan_mount' : for field setter : missing Type

// UNSUPPORTED : C value 'drive_eject_button' : for field getter : missing Type

// UNSUPPORTED : C value 'drive_eject_button' : for field setter : missing Type

// UNSUPPORTED : C value 'drive_stop_button' : for field getter : missing Type

// UNSUPPORTED : C value 'drive_stop_button' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved4' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved5' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved5' : for field setter : missing Type

// UNSUPPORTED : C value '_g_reserved6' : for field getter : missing Type

// UNSUPPORTED : C value '_g_reserved6' : for field setter : missing Type

// VolumeMonitorClassStruct creates an uninitialised VolumeMonitorClass.
func VolumeMonitorClassStruct() *VolumeMonitorClass {
	err := volumeMonitorClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := VolumeMonitorClassNewFromNative(volumeMonitorClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeVolumeMonitorClass)
	return structGo
}
func finalizeVolumeMonitorClass(obj *VolumeMonitorClass) {
	volumeMonitorClassStruct.Free(obj.native)
}

var zlibCompressorClassStruct *gi.Struct
var zlibCompressorClassStruct_Once sync.Once

func zlibCompressorClassStruct_Set() error {
	var err error
	zlibCompressorClassStruct_Once.Do(func() {
		zlibCompressorClassStruct, err = gi.StructNew("Gio", "ZlibCompressorClass")
	})
	return err
}

type ZlibCompressorClass struct {
	native unsafe.Pointer
}

func ZlibCompressorClassNewFromNative(native unsafe.Pointer) *ZlibCompressorClass {
	instance := &ZlibCompressorClass{native: native}

	return instance
}

/*
CastToZlibCompressorClass down casts any arbitrary Object to ZlibCompressorClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a ZlibCompressorClass.
*/
func (recv *ZlibCompressorClass) CastToZlibCompressorClass(object *gobject.Object) *ZlibCompressorClass {
	return ZlibCompressorClassNewFromNative(object.Native())
}

func (recv *ZlibCompressorClass) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// ZlibCompressorClassStruct creates an uninitialised ZlibCompressorClass.
func ZlibCompressorClassStruct() *ZlibCompressorClass {
	err := zlibCompressorClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := ZlibCompressorClassNewFromNative(zlibCompressorClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeZlibCompressorClass)
	return structGo
}
func finalizeZlibCompressorClass(obj *ZlibCompressorClass) {
	zlibCompressorClassStruct.Free(obj.native)
}

var zlibDecompressorClassStruct *gi.Struct
var zlibDecompressorClassStruct_Once sync.Once

func zlibDecompressorClassStruct_Set() error {
	var err error
	zlibDecompressorClassStruct_Once.Do(func() {
		zlibDecompressorClassStruct, err = gi.StructNew("Gio", "ZlibDecompressorClass")
	})
	return err
}

type ZlibDecompressorClass struct {
	native unsafe.Pointer
}

func ZlibDecompressorClassNewFromNative(native unsafe.Pointer) *ZlibDecompressorClass {
	instance := &ZlibDecompressorClass{native: native}

	return instance
}

/*
CastToZlibDecompressorClass down casts any arbitrary Object to ZlibDecompressorClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a ZlibDecompressorClass.
*/
func (recv *ZlibDecompressorClass) CastToZlibDecompressorClass(object *gobject.Object) *ZlibDecompressorClass {
	return ZlibDecompressorClassNewFromNative(object.Native())
}

func (recv *ZlibDecompressorClass) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// ZlibDecompressorClassStruct creates an uninitialised ZlibDecompressorClass.
func ZlibDecompressorClassStruct() *ZlibDecompressorClass {
	err := zlibDecompressorClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := ZlibDecompressorClassNewFromNative(zlibDecompressorClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeZlibDecompressorClass)
	return structGo
}
func finalizeZlibDecompressorClass(obj *ZlibDecompressorClass) {
	zlibDecompressorClassStruct.Free(obj.native)
}
