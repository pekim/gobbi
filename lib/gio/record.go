// Code generated - DO NOT EDIT.

package gio

import (
	gi "github.com/pekim/gobbi/internal/gi"
	"runtime"
	"sync"
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
	Native uintptr
}

// FieldName returns the C field 'name'.
func (recv *ActionEntry) FieldName() string {
	argValue := gi.StructFieldGet(actionEntryStruct, recv.Native, "name")
	value := argValue.String(false)
	return value
}

// SetFieldName sets the value of the C field 'name'.
func (recv *ActionEntry) SetFieldName(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.StructFieldSet(actionEntryStruct, recv.Native, "name", argValue)
}

// UNSUPPORTED : C value 'activate' : for field getter : missing Type

// UNSUPPORTED : C value 'activate' : for field setter : missing Type

// FieldParameterType returns the C field 'parameter_type'.
func (recv *ActionEntry) FieldParameterType() string {
	argValue := gi.StructFieldGet(actionEntryStruct, recv.Native, "parameter_type")
	value := argValue.String(false)
	return value
}

// SetFieldParameterType sets the value of the C field 'parameter_type'.
func (recv *ActionEntry) SetFieldParameterType(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.StructFieldSet(actionEntryStruct, recv.Native, "parameter_type", argValue)
}

// FieldState returns the C field 'state'.
func (recv *ActionEntry) FieldState() string {
	argValue := gi.StructFieldGet(actionEntryStruct, recv.Native, "state")
	value := argValue.String(false)
	return value
}

// SetFieldState sets the value of the C field 'state'.
func (recv *ActionEntry) SetFieldState(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.StructFieldSet(actionEntryStruct, recv.Native, "state", argValue)
}

// UNSUPPORTED : C value 'change_state' : for field getter : missing Type

// UNSUPPORTED : C value 'change_state' : for field setter : missing Type

// ActionEntryStruct creates an uninitialised ActionEntry.
func ActionEntryStruct() *ActionEntry {
	err := actionEntryStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ActionEntry{}
	structGo.Native = actionEntryStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeActionEntry)
	return structGo
}
func finalizeActionEntry(obj *ActionEntry) {
	actionEntryStruct.Free(obj.Native)
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
	Native uintptr
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

	structGo := &ActionGroupInterface{}
	structGo.Native = actionGroupInterfaceStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeActionGroupInterface)
	return structGo
}
func finalizeActionGroupInterface(obj *ActionGroupInterface) {
	actionGroupInterfaceStruct.Free(obj.Native)
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
	Native uintptr
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

	structGo := &ActionInterface{}
	structGo.Native = actionInterfaceStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeActionInterface)
	return structGo
}
func finalizeActionInterface(obj *ActionInterface) {
	actionInterfaceStruct.Free(obj.Native)
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
	Native uintptr
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

	structGo := &ActionMapInterface{}
	structGo.Native = actionMapInterfaceStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeActionMapInterface)
	return structGo
}
func finalizeActionMapInterface(obj *ActionMapInterface) {
	actionMapInterfaceStruct.Free(obj.Native)
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
	Native uintptr
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

	structGo := &AppInfoIface{}
	structGo.Native = appInfoIfaceStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeAppInfoIface)
	return structGo
}
func finalizeAppInfoIface(obj *AppInfoIface) {
	appInfoIfaceStruct.Free(obj.Native)
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
	Native uintptr
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

	structGo := &AppLaunchContextClass{}
	structGo.Native = appLaunchContextClassStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeAppLaunchContextClass)
	return structGo
}
func finalizeAppLaunchContextClass(obj *AppLaunchContextClass) {
	appLaunchContextClassStruct.Free(obj.Native)
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
	Native uintptr
}

// AppLaunchContextPrivateStruct creates an uninitialised AppLaunchContextPrivate.
func AppLaunchContextPrivateStruct() *AppLaunchContextPrivate {
	err := appLaunchContextPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &AppLaunchContextPrivate{}
	structGo.Native = appLaunchContextPrivateStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeAppLaunchContextPrivate)
	return structGo
}
func finalizeAppLaunchContextPrivate(obj *AppLaunchContextPrivate) {
	appLaunchContextPrivateStruct.Free(obj.Native)
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
	Native uintptr
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

	structGo := &ApplicationClass{}
	structGo.Native = applicationClassStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeApplicationClass)
	return structGo
}
func finalizeApplicationClass(obj *ApplicationClass) {
	applicationClassStruct.Free(obj.Native)
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
	Native uintptr
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

	structGo := &ApplicationCommandLineClass{}
	structGo.Native = applicationCommandLineClassStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeApplicationCommandLineClass)
	return structGo
}
func finalizeApplicationCommandLineClass(obj *ApplicationCommandLineClass) {
	applicationCommandLineClassStruct.Free(obj.Native)
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
	Native uintptr
}

// ApplicationCommandLinePrivateStruct creates an uninitialised ApplicationCommandLinePrivate.
func ApplicationCommandLinePrivateStruct() *ApplicationCommandLinePrivate {
	err := applicationCommandLinePrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ApplicationCommandLinePrivate{}
	structGo.Native = applicationCommandLinePrivateStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeApplicationCommandLinePrivate)
	return structGo
}
func finalizeApplicationCommandLinePrivate(obj *ApplicationCommandLinePrivate) {
	applicationCommandLinePrivateStruct.Free(obj.Native)
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
	Native uintptr
}

// ApplicationPrivateStruct creates an uninitialised ApplicationPrivate.
func ApplicationPrivateStruct() *ApplicationPrivate {
	err := applicationPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ApplicationPrivate{}
	structGo.Native = applicationPrivateStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeApplicationPrivate)
	return structGo
}
func finalizeApplicationPrivate(obj *ApplicationPrivate) {
	applicationPrivateStruct.Free(obj.Native)
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
	Native uintptr
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

	structGo := &AsyncInitableIface{}
	structGo.Native = asyncInitableIfaceStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeAsyncInitableIface)
	return structGo
}
func finalizeAsyncInitableIface(obj *AsyncInitableIface) {
	asyncInitableIfaceStruct.Free(obj.Native)
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
	Native uintptr
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

	structGo := &AsyncResultIface{}
	structGo.Native = asyncResultIfaceStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeAsyncResultIface)
	return structGo
}
func finalizeAsyncResultIface(obj *AsyncResultIface) {
	asyncResultIfaceStruct.Free(obj.Native)
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
	Native uintptr
}

// FieldParentClass returns the C field 'parent_class'.
func (recv *BufferedInputStreamClass) FieldParentClass() *FilterInputStreamClass {
	argValue := gi.StructFieldGet(bufferedInputStreamClassStruct, recv.Native, "parent_class")
	value := &FilterInputStreamClass{}
	value.Native = argValue.Pointer()
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *BufferedInputStreamClass) SetFieldParentClass(value *FilterInputStreamClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native)
	gi.StructFieldSet(bufferedInputStreamClassStruct, recv.Native, "parent_class", argValue)
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

	structGo := &BufferedInputStreamClass{}
	structGo.Native = bufferedInputStreamClassStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeBufferedInputStreamClass)
	return structGo
}
func finalizeBufferedInputStreamClass(obj *BufferedInputStreamClass) {
	bufferedInputStreamClassStruct.Free(obj.Native)
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
	Native uintptr
}

// BufferedInputStreamPrivateStruct creates an uninitialised BufferedInputStreamPrivate.
func BufferedInputStreamPrivateStruct() *BufferedInputStreamPrivate {
	err := bufferedInputStreamPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &BufferedInputStreamPrivate{}
	structGo.Native = bufferedInputStreamPrivateStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeBufferedInputStreamPrivate)
	return structGo
}
func finalizeBufferedInputStreamPrivate(obj *BufferedInputStreamPrivate) {
	bufferedInputStreamPrivateStruct.Free(obj.Native)
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
	Native uintptr
}

// FieldParentClass returns the C field 'parent_class'.
func (recv *BufferedOutputStreamClass) FieldParentClass() *FilterOutputStreamClass {
	argValue := gi.StructFieldGet(bufferedOutputStreamClassStruct, recv.Native, "parent_class")
	value := &FilterOutputStreamClass{}
	value.Native = argValue.Pointer()
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *BufferedOutputStreamClass) SetFieldParentClass(value *FilterOutputStreamClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native)
	gi.StructFieldSet(bufferedOutputStreamClassStruct, recv.Native, "parent_class", argValue)
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

	structGo := &BufferedOutputStreamClass{}
	structGo.Native = bufferedOutputStreamClassStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeBufferedOutputStreamClass)
	return structGo
}
func finalizeBufferedOutputStreamClass(obj *BufferedOutputStreamClass) {
	bufferedOutputStreamClassStruct.Free(obj.Native)
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
	Native uintptr
}

// BufferedOutputStreamPrivateStruct creates an uninitialised BufferedOutputStreamPrivate.
func BufferedOutputStreamPrivateStruct() *BufferedOutputStreamPrivate {
	err := bufferedOutputStreamPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &BufferedOutputStreamPrivate{}
	structGo.Native = bufferedOutputStreamPrivateStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeBufferedOutputStreamPrivate)
	return structGo
}
func finalizeBufferedOutputStreamPrivate(obj *BufferedOutputStreamPrivate) {
	bufferedOutputStreamPrivateStruct.Free(obj.Native)
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
	Native uintptr
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

	structGo := &CancellableClass{}
	structGo.Native = cancellableClassStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeCancellableClass)
	return structGo
}
func finalizeCancellableClass(obj *CancellableClass) {
	cancellableClassStruct.Free(obj.Native)
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
	Native uintptr
}

// CancellablePrivateStruct creates an uninitialised CancellablePrivate.
func CancellablePrivateStruct() *CancellablePrivate {
	err := cancellablePrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &CancellablePrivate{}
	structGo.Native = cancellablePrivateStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeCancellablePrivate)
	return structGo
}
func finalizeCancellablePrivate(obj *CancellablePrivate) {
	cancellablePrivateStruct.Free(obj.Native)
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
	Native uintptr
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// CharsetConverterClassStruct creates an uninitialised CharsetConverterClass.
func CharsetConverterClassStruct() *CharsetConverterClass {
	err := charsetConverterClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &CharsetConverterClass{}
	structGo.Native = charsetConverterClassStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeCharsetConverterClass)
	return structGo
}
func finalizeCharsetConverterClass(obj *CharsetConverterClass) {
	charsetConverterClassStruct.Free(obj.Native)
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
	Native uintptr
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

	structGo := &ConverterIface{}
	structGo.Native = converterIfaceStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeConverterIface)
	return structGo
}
func finalizeConverterIface(obj *ConverterIface) {
	converterIfaceStruct.Free(obj.Native)
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
	Native uintptr
}

// FieldParentClass returns the C field 'parent_class'.
func (recv *ConverterInputStreamClass) FieldParentClass() *FilterInputStreamClass {
	argValue := gi.StructFieldGet(converterInputStreamClassStruct, recv.Native, "parent_class")
	value := &FilterInputStreamClass{}
	value.Native = argValue.Pointer()
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *ConverterInputStreamClass) SetFieldParentClass(value *FilterInputStreamClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native)
	gi.StructFieldSet(converterInputStreamClassStruct, recv.Native, "parent_class", argValue)
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

	structGo := &ConverterInputStreamClass{}
	structGo.Native = converterInputStreamClassStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeConverterInputStreamClass)
	return structGo
}
func finalizeConverterInputStreamClass(obj *ConverterInputStreamClass) {
	converterInputStreamClassStruct.Free(obj.Native)
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
	Native uintptr
}

// ConverterInputStreamPrivateStruct creates an uninitialised ConverterInputStreamPrivate.
func ConverterInputStreamPrivateStruct() *ConverterInputStreamPrivate {
	err := converterInputStreamPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ConverterInputStreamPrivate{}
	structGo.Native = converterInputStreamPrivateStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeConverterInputStreamPrivate)
	return structGo
}
func finalizeConverterInputStreamPrivate(obj *ConverterInputStreamPrivate) {
	converterInputStreamPrivateStruct.Free(obj.Native)
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
	Native uintptr
}

// FieldParentClass returns the C field 'parent_class'.
func (recv *ConverterOutputStreamClass) FieldParentClass() *FilterOutputStreamClass {
	argValue := gi.StructFieldGet(converterOutputStreamClassStruct, recv.Native, "parent_class")
	value := &FilterOutputStreamClass{}
	value.Native = argValue.Pointer()
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *ConverterOutputStreamClass) SetFieldParentClass(value *FilterOutputStreamClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native)
	gi.StructFieldSet(converterOutputStreamClassStruct, recv.Native, "parent_class", argValue)
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

	structGo := &ConverterOutputStreamClass{}
	structGo.Native = converterOutputStreamClassStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeConverterOutputStreamClass)
	return structGo
}
func finalizeConverterOutputStreamClass(obj *ConverterOutputStreamClass) {
	converterOutputStreamClassStruct.Free(obj.Native)
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
	Native uintptr
}

// ConverterOutputStreamPrivateStruct creates an uninitialised ConverterOutputStreamPrivate.
func ConverterOutputStreamPrivateStruct() *ConverterOutputStreamPrivate {
	err := converterOutputStreamPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ConverterOutputStreamPrivate{}
	structGo.Native = converterOutputStreamPrivateStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeConverterOutputStreamPrivate)
	return structGo
}
func finalizeConverterOutputStreamPrivate(obj *ConverterOutputStreamPrivate) {
	converterOutputStreamPrivateStruct.Free(obj.Native)
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
	Native uintptr
}

// CredentialsClassStruct creates an uninitialised CredentialsClass.
func CredentialsClassStruct() *CredentialsClass {
	err := credentialsClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &CredentialsClass{}
	structGo.Native = credentialsClassStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeCredentialsClass)
	return structGo
}
func finalizeCredentialsClass(obj *CredentialsClass) {
	credentialsClassStruct.Free(obj.Native)
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
	Native uintptr
}

// FieldRefCount returns the C field 'ref_count'.
func (recv *DBusAnnotationInfo) FieldRefCount() int32 {
	argValue := gi.StructFieldGet(dBusAnnotationInfoStruct, recv.Native, "ref_count")
	value := argValue.Int32()
	return value
}

// SetFieldRefCount sets the value of the C field 'ref_count'.
func (recv *DBusAnnotationInfo) SetFieldRefCount(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.StructFieldSet(dBusAnnotationInfoStruct, recv.Native, "ref_count", argValue)
}

// FieldKey returns the C field 'key'.
func (recv *DBusAnnotationInfo) FieldKey() string {
	argValue := gi.StructFieldGet(dBusAnnotationInfoStruct, recv.Native, "key")
	value := argValue.String(false)
	return value
}

// SetFieldKey sets the value of the C field 'key'.
func (recv *DBusAnnotationInfo) SetFieldKey(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.StructFieldSet(dBusAnnotationInfoStruct, recv.Native, "key", argValue)
}

// FieldValue returns the C field 'value'.
func (recv *DBusAnnotationInfo) FieldValue() string {
	argValue := gi.StructFieldGet(dBusAnnotationInfoStruct, recv.Native, "value")
	value := argValue.String(false)
	return value
}

// SetFieldValue sets the value of the C field 'value'.
func (recv *DBusAnnotationInfo) SetFieldValue(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.StructFieldSet(dBusAnnotationInfoStruct, recv.Native, "value", argValue)
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
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := dBusAnnotationInfoRefFunction_Set()
	if err == nil {
		ret = dBusAnnotationInfoRefFunction.Invoke(inArgs[:], nil)
	}

	retGo := &DBusAnnotationInfo{}
	retGo.Native = ret.Pointer()

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
	inArgs[0].SetPointer(recv.Native)

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

	structGo := &DBusAnnotationInfo{}
	structGo.Native = dBusAnnotationInfoStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeDBusAnnotationInfo)
	return structGo
}
func finalizeDBusAnnotationInfo(obj *DBusAnnotationInfo) {
	dBusAnnotationInfoStruct.Free(obj.Native)
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
	Native uintptr
}

// FieldRefCount returns the C field 'ref_count'.
func (recv *DBusArgInfo) FieldRefCount() int32 {
	argValue := gi.StructFieldGet(dBusArgInfoStruct, recv.Native, "ref_count")
	value := argValue.Int32()
	return value
}

// SetFieldRefCount sets the value of the C field 'ref_count'.
func (recv *DBusArgInfo) SetFieldRefCount(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.StructFieldSet(dBusArgInfoStruct, recv.Native, "ref_count", argValue)
}

// FieldName returns the C field 'name'.
func (recv *DBusArgInfo) FieldName() string {
	argValue := gi.StructFieldGet(dBusArgInfoStruct, recv.Native, "name")
	value := argValue.String(false)
	return value
}

// SetFieldName sets the value of the C field 'name'.
func (recv *DBusArgInfo) SetFieldName(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.StructFieldSet(dBusArgInfoStruct, recv.Native, "name", argValue)
}

// FieldSignature returns the C field 'signature'.
func (recv *DBusArgInfo) FieldSignature() string {
	argValue := gi.StructFieldGet(dBusArgInfoStruct, recv.Native, "signature")
	value := argValue.String(false)
	return value
}

// SetFieldSignature sets the value of the C field 'signature'.
func (recv *DBusArgInfo) SetFieldSignature(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.StructFieldSet(dBusArgInfoStruct, recv.Native, "signature", argValue)
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
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := dBusArgInfoRefFunction_Set()
	if err == nil {
		ret = dBusArgInfoRefFunction.Invoke(inArgs[:], nil)
	}

	retGo := &DBusArgInfo{}
	retGo.Native = ret.Pointer()

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
	inArgs[0].SetPointer(recv.Native)

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

	structGo := &DBusArgInfo{}
	structGo.Native = dBusArgInfoStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeDBusArgInfo)
	return structGo
}
func finalizeDBusArgInfo(obj *DBusArgInfo) {
	dBusArgInfoStruct.Free(obj.Native)
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
	Native uintptr
}

// FieldErrorCode returns the C field 'error_code'.
func (recv *DBusErrorEntry) FieldErrorCode() int32 {
	argValue := gi.StructFieldGet(dBusErrorEntryStruct, recv.Native, "error_code")
	value := argValue.Int32()
	return value
}

// SetFieldErrorCode sets the value of the C field 'error_code'.
func (recv *DBusErrorEntry) SetFieldErrorCode(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.StructFieldSet(dBusErrorEntryStruct, recv.Native, "error_code", argValue)
}

// FieldDbusErrorName returns the C field 'dbus_error_name'.
func (recv *DBusErrorEntry) FieldDbusErrorName() string {
	argValue := gi.StructFieldGet(dBusErrorEntryStruct, recv.Native, "dbus_error_name")
	value := argValue.String(false)
	return value
}

// SetFieldDbusErrorName sets the value of the C field 'dbus_error_name'.
func (recv *DBusErrorEntry) SetFieldDbusErrorName(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.StructFieldSet(dBusErrorEntryStruct, recv.Native, "dbus_error_name", argValue)
}

// DBusErrorEntryStruct creates an uninitialised DBusErrorEntry.
func DBusErrorEntryStruct() *DBusErrorEntry {
	err := dBusErrorEntryStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &DBusErrorEntry{}
	structGo.Native = dBusErrorEntryStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeDBusErrorEntry)
	return structGo
}
func finalizeDBusErrorEntry(obj *DBusErrorEntry) {
	dBusErrorEntryStruct.Free(obj.Native)
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
	Native uintptr
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

	structGo := &DBusInterfaceIface{}
	structGo.Native = dBusInterfaceIfaceStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeDBusInterfaceIface)
	return structGo
}
func finalizeDBusInterfaceIface(obj *DBusInterfaceIface) {
	dBusInterfaceIfaceStruct.Free(obj.Native)
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
	Native uintptr
}

// FieldRefCount returns the C field 'ref_count'.
func (recv *DBusInterfaceInfo) FieldRefCount() int32 {
	argValue := gi.StructFieldGet(dBusInterfaceInfoStruct, recv.Native, "ref_count")
	value := argValue.Int32()
	return value
}

// SetFieldRefCount sets the value of the C field 'ref_count'.
func (recv *DBusInterfaceInfo) SetFieldRefCount(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.StructFieldSet(dBusInterfaceInfoStruct, recv.Native, "ref_count", argValue)
}

// FieldName returns the C field 'name'.
func (recv *DBusInterfaceInfo) FieldName() string {
	argValue := gi.StructFieldGet(dBusInterfaceInfoStruct, recv.Native, "name")
	value := argValue.String(false)
	return value
}

// SetFieldName sets the value of the C field 'name'.
func (recv *DBusInterfaceInfo) SetFieldName(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.StructFieldSet(dBusInterfaceInfoStruct, recv.Native, "name", argValue)
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
	inArgs[0].SetPointer(recv.Native)

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
	inArgs[0].SetPointer(recv.Native)

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
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetString(name)

	var ret gi.Argument

	err := dBusInterfaceInfoLookupMethodFunction_Set()
	if err == nil {
		ret = dBusInterfaceInfoLookupMethodFunction.Invoke(inArgs[:], nil)
	}

	retGo := &DBusMethodInfo{}
	retGo.Native = ret.Pointer()

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
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetString(name)

	var ret gi.Argument

	err := dBusInterfaceInfoLookupPropertyFunction_Set()
	if err == nil {
		ret = dBusInterfaceInfoLookupPropertyFunction.Invoke(inArgs[:], nil)
	}

	retGo := &DBusPropertyInfo{}
	retGo.Native = ret.Pointer()

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
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetString(name)

	var ret gi.Argument

	err := dBusInterfaceInfoLookupSignalFunction_Set()
	if err == nil {
		ret = dBusInterfaceInfoLookupSignalFunction.Invoke(inArgs[:], nil)
	}

	retGo := &DBusSignalInfo{}
	retGo.Native = ret.Pointer()

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
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := dBusInterfaceInfoRefFunction_Set()
	if err == nil {
		ret = dBusInterfaceInfoRefFunction.Invoke(inArgs[:], nil)
	}

	retGo := &DBusInterfaceInfo{}
	retGo.Native = ret.Pointer()

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
	inArgs[0].SetPointer(recv.Native)

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

	structGo := &DBusInterfaceInfo{}
	structGo.Native = dBusInterfaceInfoStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeDBusInterfaceInfo)
	return structGo
}
func finalizeDBusInterfaceInfo(obj *DBusInterfaceInfo) {
	dBusInterfaceInfoStruct.Free(obj.Native)
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
	Native uintptr
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

	structGo := &DBusInterfaceSkeletonClass{}
	structGo.Native = dBusInterfaceSkeletonClassStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeDBusInterfaceSkeletonClass)
	return structGo
}
func finalizeDBusInterfaceSkeletonClass(obj *DBusInterfaceSkeletonClass) {
	dBusInterfaceSkeletonClassStruct.Free(obj.Native)
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
	Native uintptr
}

// DBusInterfaceSkeletonPrivateStruct creates an uninitialised DBusInterfaceSkeletonPrivate.
func DBusInterfaceSkeletonPrivateStruct() *DBusInterfaceSkeletonPrivate {
	err := dBusInterfaceSkeletonPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &DBusInterfaceSkeletonPrivate{}
	structGo.Native = dBusInterfaceSkeletonPrivateStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeDBusInterfaceSkeletonPrivate)
	return structGo
}
func finalizeDBusInterfaceSkeletonPrivate(obj *DBusInterfaceSkeletonPrivate) {
	dBusInterfaceSkeletonPrivateStruct.Free(obj.Native)
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
	Native uintptr
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

	structGo := &DBusInterfaceVTable{}
	structGo.Native = dBusInterfaceVTableStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeDBusInterfaceVTable)
	return structGo
}
func finalizeDBusInterfaceVTable(obj *DBusInterfaceVTable) {
	dBusInterfaceVTableStruct.Free(obj.Native)
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
	Native uintptr
}

// FieldRefCount returns the C field 'ref_count'.
func (recv *DBusMethodInfo) FieldRefCount() int32 {
	argValue := gi.StructFieldGet(dBusMethodInfoStruct, recv.Native, "ref_count")
	value := argValue.Int32()
	return value
}

// SetFieldRefCount sets the value of the C field 'ref_count'.
func (recv *DBusMethodInfo) SetFieldRefCount(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.StructFieldSet(dBusMethodInfoStruct, recv.Native, "ref_count", argValue)
}

// FieldName returns the C field 'name'.
func (recv *DBusMethodInfo) FieldName() string {
	argValue := gi.StructFieldGet(dBusMethodInfoStruct, recv.Native, "name")
	value := argValue.String(false)
	return value
}

// SetFieldName sets the value of the C field 'name'.
func (recv *DBusMethodInfo) SetFieldName(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.StructFieldSet(dBusMethodInfoStruct, recv.Native, "name", argValue)
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
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := dBusMethodInfoRefFunction_Set()
	if err == nil {
		ret = dBusMethodInfoRefFunction.Invoke(inArgs[:], nil)
	}

	retGo := &DBusMethodInfo{}
	retGo.Native = ret.Pointer()

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
	inArgs[0].SetPointer(recv.Native)

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

	structGo := &DBusMethodInfo{}
	structGo.Native = dBusMethodInfoStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeDBusMethodInfo)
	return structGo
}
func finalizeDBusMethodInfo(obj *DBusMethodInfo) {
	dBusMethodInfoStruct.Free(obj.Native)
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
	Native uintptr
}

// FieldRefCount returns the C field 'ref_count'.
func (recv *DBusNodeInfo) FieldRefCount() int32 {
	argValue := gi.StructFieldGet(dBusNodeInfoStruct, recv.Native, "ref_count")
	value := argValue.Int32()
	return value
}

// SetFieldRefCount sets the value of the C field 'ref_count'.
func (recv *DBusNodeInfo) SetFieldRefCount(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.StructFieldSet(dBusNodeInfoStruct, recv.Native, "ref_count", argValue)
}

// FieldPath returns the C field 'path'.
func (recv *DBusNodeInfo) FieldPath() string {
	argValue := gi.StructFieldGet(dBusNodeInfoStruct, recv.Native, "path")
	value := argValue.String(false)
	return value
}

// SetFieldPath sets the value of the C field 'path'.
func (recv *DBusNodeInfo) SetFieldPath(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.StructFieldSet(dBusNodeInfoStruct, recv.Native, "path", argValue)
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

	retGo := &DBusNodeInfo{}
	retGo.Native = ret.Pointer()

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
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetString(name)

	var ret gi.Argument

	err := dBusNodeInfoLookupInterfaceFunction_Set()
	if err == nil {
		ret = dBusNodeInfoLookupInterfaceFunction.Invoke(inArgs[:], nil)
	}

	retGo := &DBusInterfaceInfo{}
	retGo.Native = ret.Pointer()

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
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := dBusNodeInfoRefFunction_Set()
	if err == nil {
		ret = dBusNodeInfoRefFunction.Invoke(inArgs[:], nil)
	}

	retGo := &DBusNodeInfo{}
	retGo.Native = ret.Pointer()

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
	inArgs[0].SetPointer(recv.Native)

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
	Native uintptr
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

	structGo := &DBusObjectIface{}
	structGo.Native = dBusObjectIfaceStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeDBusObjectIface)
	return structGo
}
func finalizeDBusObjectIface(obj *DBusObjectIface) {
	dBusObjectIfaceStruct.Free(obj.Native)
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
	Native uintptr
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

	structGo := &DBusObjectManagerClientClass{}
	structGo.Native = dBusObjectManagerClientClassStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeDBusObjectManagerClientClass)
	return structGo
}
func finalizeDBusObjectManagerClientClass(obj *DBusObjectManagerClientClass) {
	dBusObjectManagerClientClassStruct.Free(obj.Native)
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
	Native uintptr
}

// DBusObjectManagerClientPrivateStruct creates an uninitialised DBusObjectManagerClientPrivate.
func DBusObjectManagerClientPrivateStruct() *DBusObjectManagerClientPrivate {
	err := dBusObjectManagerClientPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &DBusObjectManagerClientPrivate{}
	structGo.Native = dBusObjectManagerClientPrivateStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeDBusObjectManagerClientPrivate)
	return structGo
}
func finalizeDBusObjectManagerClientPrivate(obj *DBusObjectManagerClientPrivate) {
	dBusObjectManagerClientPrivateStruct.Free(obj.Native)
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
	Native uintptr
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

	structGo := &DBusObjectManagerIface{}
	structGo.Native = dBusObjectManagerIfaceStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeDBusObjectManagerIface)
	return structGo
}
func finalizeDBusObjectManagerIface(obj *DBusObjectManagerIface) {
	dBusObjectManagerIfaceStruct.Free(obj.Native)
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
	Native uintptr
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// DBusObjectManagerServerClassStruct creates an uninitialised DBusObjectManagerServerClass.
func DBusObjectManagerServerClassStruct() *DBusObjectManagerServerClass {
	err := dBusObjectManagerServerClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &DBusObjectManagerServerClass{}
	structGo.Native = dBusObjectManagerServerClassStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeDBusObjectManagerServerClass)
	return structGo
}
func finalizeDBusObjectManagerServerClass(obj *DBusObjectManagerServerClass) {
	dBusObjectManagerServerClassStruct.Free(obj.Native)
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
	Native uintptr
}

// DBusObjectManagerServerPrivateStruct creates an uninitialised DBusObjectManagerServerPrivate.
func DBusObjectManagerServerPrivateStruct() *DBusObjectManagerServerPrivate {
	err := dBusObjectManagerServerPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &DBusObjectManagerServerPrivate{}
	structGo.Native = dBusObjectManagerServerPrivateStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeDBusObjectManagerServerPrivate)
	return structGo
}
func finalizeDBusObjectManagerServerPrivate(obj *DBusObjectManagerServerPrivate) {
	dBusObjectManagerServerPrivateStruct.Free(obj.Native)
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
	Native uintptr
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// DBusObjectProxyClassStruct creates an uninitialised DBusObjectProxyClass.
func DBusObjectProxyClassStruct() *DBusObjectProxyClass {
	err := dBusObjectProxyClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &DBusObjectProxyClass{}
	structGo.Native = dBusObjectProxyClassStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeDBusObjectProxyClass)
	return structGo
}
func finalizeDBusObjectProxyClass(obj *DBusObjectProxyClass) {
	dBusObjectProxyClassStruct.Free(obj.Native)
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
	Native uintptr
}

// DBusObjectProxyPrivateStruct creates an uninitialised DBusObjectProxyPrivate.
func DBusObjectProxyPrivateStruct() *DBusObjectProxyPrivate {
	err := dBusObjectProxyPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &DBusObjectProxyPrivate{}
	structGo.Native = dBusObjectProxyPrivateStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeDBusObjectProxyPrivate)
	return structGo
}
func finalizeDBusObjectProxyPrivate(obj *DBusObjectProxyPrivate) {
	dBusObjectProxyPrivateStruct.Free(obj.Native)
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
	Native uintptr
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

	structGo := &DBusObjectSkeletonClass{}
	structGo.Native = dBusObjectSkeletonClassStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeDBusObjectSkeletonClass)
	return structGo
}
func finalizeDBusObjectSkeletonClass(obj *DBusObjectSkeletonClass) {
	dBusObjectSkeletonClassStruct.Free(obj.Native)
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
	Native uintptr
}

// DBusObjectSkeletonPrivateStruct creates an uninitialised DBusObjectSkeletonPrivate.
func DBusObjectSkeletonPrivateStruct() *DBusObjectSkeletonPrivate {
	err := dBusObjectSkeletonPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &DBusObjectSkeletonPrivate{}
	structGo.Native = dBusObjectSkeletonPrivateStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeDBusObjectSkeletonPrivate)
	return structGo
}
func finalizeDBusObjectSkeletonPrivate(obj *DBusObjectSkeletonPrivate) {
	dBusObjectSkeletonPrivateStruct.Free(obj.Native)
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
	Native uintptr
}

// FieldRefCount returns the C field 'ref_count'.
func (recv *DBusPropertyInfo) FieldRefCount() int32 {
	argValue := gi.StructFieldGet(dBusPropertyInfoStruct, recv.Native, "ref_count")
	value := argValue.Int32()
	return value
}

// SetFieldRefCount sets the value of the C field 'ref_count'.
func (recv *DBusPropertyInfo) SetFieldRefCount(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.StructFieldSet(dBusPropertyInfoStruct, recv.Native, "ref_count", argValue)
}

// FieldName returns the C field 'name'.
func (recv *DBusPropertyInfo) FieldName() string {
	argValue := gi.StructFieldGet(dBusPropertyInfoStruct, recv.Native, "name")
	value := argValue.String(false)
	return value
}

// SetFieldName sets the value of the C field 'name'.
func (recv *DBusPropertyInfo) SetFieldName(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.StructFieldSet(dBusPropertyInfoStruct, recv.Native, "name", argValue)
}

// FieldSignature returns the C field 'signature'.
func (recv *DBusPropertyInfo) FieldSignature() string {
	argValue := gi.StructFieldGet(dBusPropertyInfoStruct, recv.Native, "signature")
	value := argValue.String(false)
	return value
}

// SetFieldSignature sets the value of the C field 'signature'.
func (recv *DBusPropertyInfo) SetFieldSignature(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.StructFieldSet(dBusPropertyInfoStruct, recv.Native, "signature", argValue)
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
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := dBusPropertyInfoRefFunction_Set()
	if err == nil {
		ret = dBusPropertyInfoRefFunction.Invoke(inArgs[:], nil)
	}

	retGo := &DBusPropertyInfo{}
	retGo.Native = ret.Pointer()

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
	inArgs[0].SetPointer(recv.Native)

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

	structGo := &DBusPropertyInfo{}
	structGo.Native = dBusPropertyInfoStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeDBusPropertyInfo)
	return structGo
}
func finalizeDBusPropertyInfo(obj *DBusPropertyInfo) {
	dBusPropertyInfoStruct.Free(obj.Native)
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
	Native uintptr
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

	structGo := &DBusProxyClass{}
	structGo.Native = dBusProxyClassStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeDBusProxyClass)
	return structGo
}
func finalizeDBusProxyClass(obj *DBusProxyClass) {
	dBusProxyClassStruct.Free(obj.Native)
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
	Native uintptr
}

// DBusProxyPrivateStruct creates an uninitialised DBusProxyPrivate.
func DBusProxyPrivateStruct() *DBusProxyPrivate {
	err := dBusProxyPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &DBusProxyPrivate{}
	structGo.Native = dBusProxyPrivateStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeDBusProxyPrivate)
	return structGo
}
func finalizeDBusProxyPrivate(obj *DBusProxyPrivate) {
	dBusProxyPrivateStruct.Free(obj.Native)
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
	Native uintptr
}

// FieldRefCount returns the C field 'ref_count'.
func (recv *DBusSignalInfo) FieldRefCount() int32 {
	argValue := gi.StructFieldGet(dBusSignalInfoStruct, recv.Native, "ref_count")
	value := argValue.Int32()
	return value
}

// SetFieldRefCount sets the value of the C field 'ref_count'.
func (recv *DBusSignalInfo) SetFieldRefCount(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.StructFieldSet(dBusSignalInfoStruct, recv.Native, "ref_count", argValue)
}

// FieldName returns the C field 'name'.
func (recv *DBusSignalInfo) FieldName() string {
	argValue := gi.StructFieldGet(dBusSignalInfoStruct, recv.Native, "name")
	value := argValue.String(false)
	return value
}

// SetFieldName sets the value of the C field 'name'.
func (recv *DBusSignalInfo) SetFieldName(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.StructFieldSet(dBusSignalInfoStruct, recv.Native, "name", argValue)
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
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := dBusSignalInfoRefFunction_Set()
	if err == nil {
		ret = dBusSignalInfoRefFunction.Invoke(inArgs[:], nil)
	}

	retGo := &DBusSignalInfo{}
	retGo.Native = ret.Pointer()

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
	inArgs[0].SetPointer(recv.Native)

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

	structGo := &DBusSignalInfo{}
	structGo.Native = dBusSignalInfoStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeDBusSignalInfo)
	return structGo
}
func finalizeDBusSignalInfo(obj *DBusSignalInfo) {
	dBusSignalInfoStruct.Free(obj.Native)
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
	Native uintptr
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

	structGo := &DBusSubtreeVTable{}
	structGo.Native = dBusSubtreeVTableStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeDBusSubtreeVTable)
	return structGo
}
func finalizeDBusSubtreeVTable(obj *DBusSubtreeVTable) {
	dBusSubtreeVTableStruct.Free(obj.Native)
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
	Native uintptr
}

// FieldParentClass returns the C field 'parent_class'.
func (recv *DataInputStreamClass) FieldParentClass() *BufferedInputStreamClass {
	argValue := gi.StructFieldGet(dataInputStreamClassStruct, recv.Native, "parent_class")
	value := &BufferedInputStreamClass{}
	value.Native = argValue.Pointer()
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *DataInputStreamClass) SetFieldParentClass(value *BufferedInputStreamClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native)
	gi.StructFieldSet(dataInputStreamClassStruct, recv.Native, "parent_class", argValue)
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

	structGo := &DataInputStreamClass{}
	structGo.Native = dataInputStreamClassStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeDataInputStreamClass)
	return structGo
}
func finalizeDataInputStreamClass(obj *DataInputStreamClass) {
	dataInputStreamClassStruct.Free(obj.Native)
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
	Native uintptr
}

// DataInputStreamPrivateStruct creates an uninitialised DataInputStreamPrivate.
func DataInputStreamPrivateStruct() *DataInputStreamPrivate {
	err := dataInputStreamPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &DataInputStreamPrivate{}
	structGo.Native = dataInputStreamPrivateStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeDataInputStreamPrivate)
	return structGo
}
func finalizeDataInputStreamPrivate(obj *DataInputStreamPrivate) {
	dataInputStreamPrivateStruct.Free(obj.Native)
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
	Native uintptr
}

// FieldParentClass returns the C field 'parent_class'.
func (recv *DataOutputStreamClass) FieldParentClass() *FilterOutputStreamClass {
	argValue := gi.StructFieldGet(dataOutputStreamClassStruct, recv.Native, "parent_class")
	value := &FilterOutputStreamClass{}
	value.Native = argValue.Pointer()
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *DataOutputStreamClass) SetFieldParentClass(value *FilterOutputStreamClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native)
	gi.StructFieldSet(dataOutputStreamClassStruct, recv.Native, "parent_class", argValue)
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

	structGo := &DataOutputStreamClass{}
	structGo.Native = dataOutputStreamClassStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeDataOutputStreamClass)
	return structGo
}
func finalizeDataOutputStreamClass(obj *DataOutputStreamClass) {
	dataOutputStreamClassStruct.Free(obj.Native)
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
	Native uintptr
}

// DataOutputStreamPrivateStruct creates an uninitialised DataOutputStreamPrivate.
func DataOutputStreamPrivateStruct() *DataOutputStreamPrivate {
	err := dataOutputStreamPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &DataOutputStreamPrivate{}
	structGo.Native = dataOutputStreamPrivateStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeDataOutputStreamPrivate)
	return structGo
}
func finalizeDataOutputStreamPrivate(obj *DataOutputStreamPrivate) {
	dataOutputStreamPrivateStruct.Free(obj.Native)
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
	Native uintptr
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

	structGo := &DatagramBasedInterface{}
	structGo.Native = datagramBasedInterfaceStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeDatagramBasedInterface)
	return structGo
}
func finalizeDatagramBasedInterface(obj *DatagramBasedInterface) {
	datagramBasedInterfaceStruct.Free(obj.Native)
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
	Native uintptr
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// DesktopAppInfoClassStruct creates an uninitialised DesktopAppInfoClass.
func DesktopAppInfoClassStruct() *DesktopAppInfoClass {
	err := desktopAppInfoClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &DesktopAppInfoClass{}
	structGo.Native = desktopAppInfoClassStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeDesktopAppInfoClass)
	return structGo
}
func finalizeDesktopAppInfoClass(obj *DesktopAppInfoClass) {
	desktopAppInfoClassStruct.Free(obj.Native)
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
	Native uintptr
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

	structGo := &DesktopAppInfoLookupIface{}
	structGo.Native = desktopAppInfoLookupIfaceStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeDesktopAppInfoLookupIface)
	return structGo
}
func finalizeDesktopAppInfoLookupIface(obj *DesktopAppInfoLookupIface) {
	desktopAppInfoLookupIfaceStruct.Free(obj.Native)
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
	Native uintptr
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

	structGo := &DriveIface{}
	structGo.Native = driveIfaceStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeDriveIface)
	return structGo
}
func finalizeDriveIface(obj *DriveIface) {
	driveIfaceStruct.Free(obj.Native)
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
	Native uintptr
}

// UNSUPPORTED : C value 'g_iface' : for field getter : no Go type for 'GObject.TypeInterface'

// UNSUPPORTED : C value 'g_iface' : for field setter : no Go type for 'GObject.TypeInterface'

// DtlsClientConnectionInterfaceStruct creates an uninitialised DtlsClientConnectionInterface.
func DtlsClientConnectionInterfaceStruct() *DtlsClientConnectionInterface {
	err := dtlsClientConnectionInterfaceStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &DtlsClientConnectionInterface{}
	structGo.Native = dtlsClientConnectionInterfaceStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeDtlsClientConnectionInterface)
	return structGo
}
func finalizeDtlsClientConnectionInterface(obj *DtlsClientConnectionInterface) {
	dtlsClientConnectionInterfaceStruct.Free(obj.Native)
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
	Native uintptr
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

	structGo := &DtlsConnectionInterface{}
	structGo.Native = dtlsConnectionInterfaceStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeDtlsConnectionInterface)
	return structGo
}
func finalizeDtlsConnectionInterface(obj *DtlsConnectionInterface) {
	dtlsConnectionInterfaceStruct.Free(obj.Native)
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
	Native uintptr
}

// UNSUPPORTED : C value 'g_iface' : for field getter : no Go type for 'GObject.TypeInterface'

// UNSUPPORTED : C value 'g_iface' : for field setter : no Go type for 'GObject.TypeInterface'

// DtlsServerConnectionInterfaceStruct creates an uninitialised DtlsServerConnectionInterface.
func DtlsServerConnectionInterfaceStruct() *DtlsServerConnectionInterface {
	err := dtlsServerConnectionInterfaceStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &DtlsServerConnectionInterface{}
	structGo.Native = dtlsServerConnectionInterfaceStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeDtlsServerConnectionInterface)
	return structGo
}
func finalizeDtlsServerConnectionInterface(obj *DtlsServerConnectionInterface) {
	dtlsServerConnectionInterfaceStruct.Free(obj.Native)
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
	Native uintptr
}

// EmblemClassStruct creates an uninitialised EmblemClass.
func EmblemClassStruct() *EmblemClass {
	err := emblemClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &EmblemClass{}
	structGo.Native = emblemClassStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeEmblemClass)
	return structGo
}
func finalizeEmblemClass(obj *EmblemClass) {
	emblemClassStruct.Free(obj.Native)
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
	Native uintptr
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// EmblemedIconClassStruct creates an uninitialised EmblemedIconClass.
func EmblemedIconClassStruct() *EmblemedIconClass {
	err := emblemedIconClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &EmblemedIconClass{}
	structGo.Native = emblemedIconClassStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeEmblemedIconClass)
	return structGo
}
func finalizeEmblemedIconClass(obj *EmblemedIconClass) {
	emblemedIconClassStruct.Free(obj.Native)
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
	Native uintptr
}

// EmblemedIconPrivateStruct creates an uninitialised EmblemedIconPrivate.
func EmblemedIconPrivateStruct() *EmblemedIconPrivate {
	err := emblemedIconPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &EmblemedIconPrivate{}
	structGo.Native = emblemedIconPrivateStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeEmblemedIconPrivate)
	return structGo
}
func finalizeEmblemedIconPrivate(obj *EmblemedIconPrivate) {
	emblemedIconPrivateStruct.Free(obj.Native)
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
	Native uintptr
}

// FieldName returns the C field 'name'.
func (recv *FileAttributeInfo) FieldName() string {
	argValue := gi.StructFieldGet(fileAttributeInfoStruct, recv.Native, "name")
	value := argValue.String(false)
	return value
}

// SetFieldName sets the value of the C field 'name'.
func (recv *FileAttributeInfo) SetFieldName(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.StructFieldSet(fileAttributeInfoStruct, recv.Native, "name", argValue)
}

// FieldType returns the C field 'type'.
func (recv *FileAttributeInfo) FieldType() FileAttributeType {
	argValue := gi.StructFieldGet(fileAttributeInfoStruct, recv.Native, "type")
	value := FileAttributeType(argValue.Int32())
	return value
}

// SetFieldType sets the value of the C field 'type'.
func (recv *FileAttributeInfo) SetFieldType(value FileAttributeType) {
	var argValue gi.Argument
	argValue.SetInt32(int32(value))
	gi.StructFieldSet(fileAttributeInfoStruct, recv.Native, "type", argValue)
}

// UNSUPPORTED : C value 'flags' : for field getter : no Go type for 'FileAttributeInfoFlags'

// UNSUPPORTED : C value 'flags' : for field setter : no Go type for 'FileAttributeInfoFlags'

// FileAttributeInfoStruct creates an uninitialised FileAttributeInfo.
func FileAttributeInfoStruct() *FileAttributeInfo {
	err := fileAttributeInfoStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &FileAttributeInfo{}
	structGo.Native = fileAttributeInfoStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeFileAttributeInfo)
	return structGo
}
func finalizeFileAttributeInfo(obj *FileAttributeInfo) {
	fileAttributeInfoStruct.Free(obj.Native)
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
	Native uintptr
}

// FieldInfos returns the C field 'infos'.
func (recv *FileAttributeInfoList) FieldInfos() *FileAttributeInfo {
	argValue := gi.StructFieldGet(fileAttributeInfoListStruct, recv.Native, "infos")
	value := &FileAttributeInfo{}
	value.Native = argValue.Pointer()
	return value
}

// SetFieldInfos sets the value of the C field 'infos'.
func (recv *FileAttributeInfoList) SetFieldInfos(value *FileAttributeInfo) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native)
	gi.StructFieldSet(fileAttributeInfoListStruct, recv.Native, "infos", argValue)
}

// FieldNInfos returns the C field 'n_infos'.
func (recv *FileAttributeInfoList) FieldNInfos() int32 {
	argValue := gi.StructFieldGet(fileAttributeInfoListStruct, recv.Native, "n_infos")
	value := argValue.Int32()
	return value
}

// SetFieldNInfos sets the value of the C field 'n_infos'.
func (recv *FileAttributeInfoList) SetFieldNInfos(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.StructFieldSet(fileAttributeInfoListStruct, recv.Native, "n_infos", argValue)
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

	retGo := &FileAttributeInfoList{}
	retGo.Native = ret.Pointer()

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
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := fileAttributeInfoListDupFunction_Set()
	if err == nil {
		ret = fileAttributeInfoListDupFunction.Invoke(inArgs[:], nil)
	}

	retGo := &FileAttributeInfoList{}
	retGo.Native = ret.Pointer()

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
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetString(name)

	var ret gi.Argument

	err := fileAttributeInfoListLookupFunction_Set()
	if err == nil {
		ret = fileAttributeInfoListLookupFunction.Invoke(inArgs[:], nil)
	}

	retGo := &FileAttributeInfo{}
	retGo.Native = ret.Pointer()

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
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := fileAttributeInfoListRefFunction_Set()
	if err == nil {
		ret = fileAttributeInfoListRefFunction.Invoke(inArgs[:], nil)
	}

	retGo := &FileAttributeInfoList{}
	retGo.Native = ret.Pointer()

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
	inArgs[0].SetPointer(recv.Native)

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
	Native uintptr
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

	retGo := &FileAttributeMatcher{}
	retGo.Native = ret.Pointer()

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
	inArgs[0].SetPointer(recv.Native)
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
	inArgs[0].SetPointer(recv.Native)

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
	inArgs[0].SetPointer(recv.Native)
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
	inArgs[0].SetPointer(recv.Native)
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
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := fileAttributeMatcherRefFunction_Set()
	if err == nil {
		ret = fileAttributeMatcherRefFunction.Invoke(inArgs[:], nil)
	}

	retGo := &FileAttributeMatcher{}
	retGo.Native = ret.Pointer()

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
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetPointer(subtract.Native)

	var ret gi.Argument

	err := fileAttributeMatcherSubtractFunction_Set()
	if err == nil {
		ret = fileAttributeMatcherSubtractFunction.Invoke(inArgs[:], nil)
	}

	retGo := &FileAttributeMatcher{}
	retGo.Native = ret.Pointer()

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
	inArgs[0].SetPointer(recv.Native)

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
	inArgs[0].SetPointer(recv.Native)

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
	Native uintptr
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

	structGo := &FileDescriptorBasedIface{}
	structGo.Native = fileDescriptorBasedIfaceStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeFileDescriptorBasedIface)
	return structGo
}
func finalizeFileDescriptorBasedIface(obj *FileDescriptorBasedIface) {
	fileDescriptorBasedIfaceStruct.Free(obj.Native)
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
	Native uintptr
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

	structGo := &FileEnumeratorClass{}
	structGo.Native = fileEnumeratorClassStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeFileEnumeratorClass)
	return structGo
}
func finalizeFileEnumeratorClass(obj *FileEnumeratorClass) {
	fileEnumeratorClassStruct.Free(obj.Native)
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
	Native uintptr
}

// FileEnumeratorPrivateStruct creates an uninitialised FileEnumeratorPrivate.
func FileEnumeratorPrivateStruct() *FileEnumeratorPrivate {
	err := fileEnumeratorPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &FileEnumeratorPrivate{}
	structGo.Native = fileEnumeratorPrivateStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeFileEnumeratorPrivate)
	return structGo
}
func finalizeFileEnumeratorPrivate(obj *FileEnumeratorPrivate) {
	fileEnumeratorPrivateStruct.Free(obj.Native)
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
	Native uintptr
}

// FieldParentClass returns the C field 'parent_class'.
func (recv *FileIOStreamClass) FieldParentClass() *IOStreamClass {
	argValue := gi.StructFieldGet(fileIOStreamClassStruct, recv.Native, "parent_class")
	value := &IOStreamClass{}
	value.Native = argValue.Pointer()
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *FileIOStreamClass) SetFieldParentClass(value *IOStreamClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native)
	gi.StructFieldSet(fileIOStreamClassStruct, recv.Native, "parent_class", argValue)
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

	structGo := &FileIOStreamClass{}
	structGo.Native = fileIOStreamClassStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeFileIOStreamClass)
	return structGo
}
func finalizeFileIOStreamClass(obj *FileIOStreamClass) {
	fileIOStreamClassStruct.Free(obj.Native)
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
	Native uintptr
}

// FileIOStreamPrivateStruct creates an uninitialised FileIOStreamPrivate.
func FileIOStreamPrivateStruct() *FileIOStreamPrivate {
	err := fileIOStreamPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &FileIOStreamPrivate{}
	structGo.Native = fileIOStreamPrivateStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeFileIOStreamPrivate)
	return structGo
}
func finalizeFileIOStreamPrivate(obj *FileIOStreamPrivate) {
	fileIOStreamPrivateStruct.Free(obj.Native)
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
	Native uintptr
}

// FileIconClassStruct creates an uninitialised FileIconClass.
func FileIconClassStruct() *FileIconClass {
	err := fileIconClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &FileIconClass{}
	structGo.Native = fileIconClassStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeFileIconClass)
	return structGo
}
func finalizeFileIconClass(obj *FileIconClass) {
	fileIconClassStruct.Free(obj.Native)
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
	Native uintptr
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
	argValue := gi.StructFieldGet(fileIfaceStruct, recv.Native, "supports_thread_contexts")
	value := argValue.Boolean()
	return value
}

// SetFieldSupportsThreadContexts sets the value of the C field 'supports_thread_contexts'.
func (recv *FileIface) SetFieldSupportsThreadContexts(value bool) {
	var argValue gi.Argument
	argValue.SetBoolean(value)
	gi.StructFieldSet(fileIfaceStruct, recv.Native, "supports_thread_contexts", argValue)
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

	structGo := &FileIface{}
	structGo.Native = fileIfaceStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeFileIface)
	return structGo
}
func finalizeFileIface(obj *FileIface) {
	fileIfaceStruct.Free(obj.Native)
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
	Native uintptr
}

// FileInfoClassStruct creates an uninitialised FileInfoClass.
func FileInfoClassStruct() *FileInfoClass {
	err := fileInfoClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &FileInfoClass{}
	structGo.Native = fileInfoClassStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeFileInfoClass)
	return structGo
}
func finalizeFileInfoClass(obj *FileInfoClass) {
	fileInfoClassStruct.Free(obj.Native)
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
	Native uintptr
}

// FieldParentClass returns the C field 'parent_class'.
func (recv *FileInputStreamClass) FieldParentClass() *InputStreamClass {
	argValue := gi.StructFieldGet(fileInputStreamClassStruct, recv.Native, "parent_class")
	value := &InputStreamClass{}
	value.Native = argValue.Pointer()
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *FileInputStreamClass) SetFieldParentClass(value *InputStreamClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native)
	gi.StructFieldSet(fileInputStreamClassStruct, recv.Native, "parent_class", argValue)
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

	structGo := &FileInputStreamClass{}
	structGo.Native = fileInputStreamClassStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeFileInputStreamClass)
	return structGo
}
func finalizeFileInputStreamClass(obj *FileInputStreamClass) {
	fileInputStreamClassStruct.Free(obj.Native)
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
	Native uintptr
}

// FileInputStreamPrivateStruct creates an uninitialised FileInputStreamPrivate.
func FileInputStreamPrivateStruct() *FileInputStreamPrivate {
	err := fileInputStreamPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &FileInputStreamPrivate{}
	structGo.Native = fileInputStreamPrivateStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeFileInputStreamPrivate)
	return structGo
}
func finalizeFileInputStreamPrivate(obj *FileInputStreamPrivate) {
	fileInputStreamPrivateStruct.Free(obj.Native)
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
	Native uintptr
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

	structGo := &FileMonitorClass{}
	structGo.Native = fileMonitorClassStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeFileMonitorClass)
	return structGo
}
func finalizeFileMonitorClass(obj *FileMonitorClass) {
	fileMonitorClassStruct.Free(obj.Native)
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
	Native uintptr
}

// FileMonitorPrivateStruct creates an uninitialised FileMonitorPrivate.
func FileMonitorPrivateStruct() *FileMonitorPrivate {
	err := fileMonitorPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &FileMonitorPrivate{}
	structGo.Native = fileMonitorPrivateStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeFileMonitorPrivate)
	return structGo
}
func finalizeFileMonitorPrivate(obj *FileMonitorPrivate) {
	fileMonitorPrivateStruct.Free(obj.Native)
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
	Native uintptr
}

// FieldParentClass returns the C field 'parent_class'.
func (recv *FileOutputStreamClass) FieldParentClass() *OutputStreamClass {
	argValue := gi.StructFieldGet(fileOutputStreamClassStruct, recv.Native, "parent_class")
	value := &OutputStreamClass{}
	value.Native = argValue.Pointer()
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *FileOutputStreamClass) SetFieldParentClass(value *OutputStreamClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native)
	gi.StructFieldSet(fileOutputStreamClassStruct, recv.Native, "parent_class", argValue)
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

	structGo := &FileOutputStreamClass{}
	structGo.Native = fileOutputStreamClassStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeFileOutputStreamClass)
	return structGo
}
func finalizeFileOutputStreamClass(obj *FileOutputStreamClass) {
	fileOutputStreamClassStruct.Free(obj.Native)
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
	Native uintptr
}

// FileOutputStreamPrivateStruct creates an uninitialised FileOutputStreamPrivate.
func FileOutputStreamPrivateStruct() *FileOutputStreamPrivate {
	err := fileOutputStreamPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &FileOutputStreamPrivate{}
	structGo.Native = fileOutputStreamPrivateStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeFileOutputStreamPrivate)
	return structGo
}
func finalizeFileOutputStreamPrivate(obj *FileOutputStreamPrivate) {
	fileOutputStreamPrivateStruct.Free(obj.Native)
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
	Native uintptr
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

	structGo := &FilenameCompleterClass{}
	structGo.Native = filenameCompleterClassStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeFilenameCompleterClass)
	return structGo
}
func finalizeFilenameCompleterClass(obj *FilenameCompleterClass) {
	filenameCompleterClassStruct.Free(obj.Native)
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
	Native uintptr
}

// FieldParentClass returns the C field 'parent_class'.
func (recv *FilterInputStreamClass) FieldParentClass() *InputStreamClass {
	argValue := gi.StructFieldGet(filterInputStreamClassStruct, recv.Native, "parent_class")
	value := &InputStreamClass{}
	value.Native = argValue.Pointer()
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *FilterInputStreamClass) SetFieldParentClass(value *InputStreamClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native)
	gi.StructFieldSet(filterInputStreamClassStruct, recv.Native, "parent_class", argValue)
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

	structGo := &FilterInputStreamClass{}
	structGo.Native = filterInputStreamClassStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeFilterInputStreamClass)
	return structGo
}
func finalizeFilterInputStreamClass(obj *FilterInputStreamClass) {
	filterInputStreamClassStruct.Free(obj.Native)
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
	Native uintptr
}

// FieldParentClass returns the C field 'parent_class'.
func (recv *FilterOutputStreamClass) FieldParentClass() *OutputStreamClass {
	argValue := gi.StructFieldGet(filterOutputStreamClassStruct, recv.Native, "parent_class")
	value := &OutputStreamClass{}
	value.Native = argValue.Pointer()
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *FilterOutputStreamClass) SetFieldParentClass(value *OutputStreamClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native)
	gi.StructFieldSet(filterOutputStreamClassStruct, recv.Native, "parent_class", argValue)
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

	structGo := &FilterOutputStreamClass{}
	structGo.Native = filterOutputStreamClassStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeFilterOutputStreamClass)
	return structGo
}
func finalizeFilterOutputStreamClass(obj *FilterOutputStreamClass) {
	filterOutputStreamClassStruct.Free(obj.Native)
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
	Native uintptr
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
	inArgs[0].SetPointer(recv.Native)

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
	inArgs[0].SetPointer(recv.Native)

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

	structGo := &IOExtension{}
	structGo.Native = iOExtensionStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeIOExtension)
	return structGo
}
func finalizeIOExtension(obj *IOExtension) {
	iOExtensionStruct.Free(obj.Native)
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
	Native uintptr
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
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetString(name)

	var ret gi.Argument

	err := iOExtensionPointGetExtensionByNameFunction_Set()
	if err == nil {
		ret = iOExtensionPointGetExtensionByNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := &IOExtension{}
	retGo.Native = ret.Pointer()

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

	structGo := &IOExtensionPoint{}
	structGo.Native = iOExtensionPointStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeIOExtensionPoint)
	return structGo
}
func finalizeIOExtensionPoint(obj *IOExtensionPoint) {
	iOExtensionPointStruct.Free(obj.Native)
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
	Native uintptr
}

// IOModuleClassStruct creates an uninitialised IOModuleClass.
func IOModuleClassStruct() *IOModuleClass {
	err := iOModuleClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &IOModuleClass{}
	structGo.Native = iOModuleClassStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeIOModuleClass)
	return structGo
}
func finalizeIOModuleClass(obj *IOModuleClass) {
	iOModuleClassStruct.Free(obj.Native)
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
	Native uintptr
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
	inArgs[0].SetPointer(recv.Native)
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
	inArgs[0].SetPointer(recv.Native)

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

	structGo := &IOModuleScope{}
	structGo.Native = iOModuleScopeStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeIOModuleScope)
	return structGo
}
func finalizeIOModuleScope(obj *IOModuleScope) {
	iOModuleScopeStruct.Free(obj.Native)
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
	Native uintptr
}

// UNSUPPORTED : C value 'g_io_scheduler_job_send_to_mainloop' : parameter 'func' of type 'GLib.SourceFunc' not supported

// UNSUPPORTED : C value 'g_io_scheduler_job_send_to_mainloop_async' : parameter 'func' of type 'GLib.SourceFunc' not supported

// IOSchedulerJobStruct creates an uninitialised IOSchedulerJob.
func IOSchedulerJobStruct() *IOSchedulerJob {
	err := iOSchedulerJobStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &IOSchedulerJob{}
	structGo.Native = iOSchedulerJobStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeIOSchedulerJob)
	return structGo
}
func finalizeIOSchedulerJob(obj *IOSchedulerJob) {
	iOSchedulerJobStruct.Free(obj.Native)
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
	Native uintptr
}

// IOStreamAdapterStruct creates an uninitialised IOStreamAdapter.
func IOStreamAdapterStruct() *IOStreamAdapter {
	err := iOStreamAdapterStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &IOStreamAdapter{}
	structGo.Native = iOStreamAdapterStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeIOStreamAdapter)
	return structGo
}
func finalizeIOStreamAdapter(obj *IOStreamAdapter) {
	iOStreamAdapterStruct.Free(obj.Native)
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
	Native uintptr
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

	structGo := &IOStreamClass{}
	structGo.Native = iOStreamClassStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeIOStreamClass)
	return structGo
}
func finalizeIOStreamClass(obj *IOStreamClass) {
	iOStreamClassStruct.Free(obj.Native)
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
	Native uintptr
}

// IOStreamPrivateStruct creates an uninitialised IOStreamPrivate.
func IOStreamPrivateStruct() *IOStreamPrivate {
	err := iOStreamPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &IOStreamPrivate{}
	structGo.Native = iOStreamPrivateStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeIOStreamPrivate)
	return structGo
}
func finalizeIOStreamPrivate(obj *IOStreamPrivate) {
	iOStreamPrivateStruct.Free(obj.Native)
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
	Native uintptr
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

	structGo := &IconIface{}
	structGo.Native = iconIfaceStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeIconIface)
	return structGo
}
func finalizeIconIface(obj *IconIface) {
	iconIfaceStruct.Free(obj.Native)
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
	Native uintptr
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

	structGo := &InetAddressClass{}
	structGo.Native = inetAddressClassStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeInetAddressClass)
	return structGo
}
func finalizeInetAddressClass(obj *InetAddressClass) {
	inetAddressClassStruct.Free(obj.Native)
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
	Native uintptr
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// InetAddressMaskClassStruct creates an uninitialised InetAddressMaskClass.
func InetAddressMaskClassStruct() *InetAddressMaskClass {
	err := inetAddressMaskClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &InetAddressMaskClass{}
	structGo.Native = inetAddressMaskClassStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeInetAddressMaskClass)
	return structGo
}
func finalizeInetAddressMaskClass(obj *InetAddressMaskClass) {
	inetAddressMaskClassStruct.Free(obj.Native)
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
	Native uintptr
}

// InetAddressMaskPrivateStruct creates an uninitialised InetAddressMaskPrivate.
func InetAddressMaskPrivateStruct() *InetAddressMaskPrivate {
	err := inetAddressMaskPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &InetAddressMaskPrivate{}
	structGo.Native = inetAddressMaskPrivateStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeInetAddressMaskPrivate)
	return structGo
}
func finalizeInetAddressMaskPrivate(obj *InetAddressMaskPrivate) {
	inetAddressMaskPrivateStruct.Free(obj.Native)
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
	Native uintptr
}

// InetAddressPrivateStruct creates an uninitialised InetAddressPrivate.
func InetAddressPrivateStruct() *InetAddressPrivate {
	err := inetAddressPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &InetAddressPrivate{}
	structGo.Native = inetAddressPrivateStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeInetAddressPrivate)
	return structGo
}
func finalizeInetAddressPrivate(obj *InetAddressPrivate) {
	inetAddressPrivateStruct.Free(obj.Native)
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
	Native uintptr
}

// FieldParentClass returns the C field 'parent_class'.
func (recv *InetSocketAddressClass) FieldParentClass() *SocketAddressClass {
	argValue := gi.StructFieldGet(inetSocketAddressClassStruct, recv.Native, "parent_class")
	value := &SocketAddressClass{}
	value.Native = argValue.Pointer()
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *InetSocketAddressClass) SetFieldParentClass(value *SocketAddressClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native)
	gi.StructFieldSet(inetSocketAddressClassStruct, recv.Native, "parent_class", argValue)
}

// InetSocketAddressClassStruct creates an uninitialised InetSocketAddressClass.
func InetSocketAddressClassStruct() *InetSocketAddressClass {
	err := inetSocketAddressClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &InetSocketAddressClass{}
	structGo.Native = inetSocketAddressClassStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeInetSocketAddressClass)
	return structGo
}
func finalizeInetSocketAddressClass(obj *InetSocketAddressClass) {
	inetSocketAddressClassStruct.Free(obj.Native)
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
	Native uintptr
}

// InetSocketAddressPrivateStruct creates an uninitialised InetSocketAddressPrivate.
func InetSocketAddressPrivateStruct() *InetSocketAddressPrivate {
	err := inetSocketAddressPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &InetSocketAddressPrivate{}
	structGo.Native = inetSocketAddressPrivateStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeInetSocketAddressPrivate)
	return structGo
}
func finalizeInetSocketAddressPrivate(obj *InetSocketAddressPrivate) {
	inetSocketAddressPrivateStruct.Free(obj.Native)
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
	Native uintptr
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

	structGo := &InitableIface{}
	structGo.Native = initableIfaceStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeInitableIface)
	return structGo
}
func finalizeInitableIface(obj *InitableIface) {
	initableIfaceStruct.Free(obj.Native)
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
	Native uintptr
}

// FieldAddress returns the C field 'address'.
func (recv *InputMessage) FieldAddress() *SocketAddress {
	argValue := gi.StructFieldGet(inputMessageStruct, recv.Native, "address")
	value := &SocketAddress{}
	value.Native = argValue.Pointer()
	return value
}

// SetFieldAddress sets the value of the C field 'address'.
func (recv *InputMessage) SetFieldAddress(value *SocketAddress) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native)
	gi.StructFieldSet(inputMessageStruct, recv.Native, "address", argValue)
}

// UNSUPPORTED : C value 'vectors' : for field getter : missing Type

// UNSUPPORTED : C value 'vectors' : for field setter : missing Type

// FieldNumVectors returns the C field 'num_vectors'.
func (recv *InputMessage) FieldNumVectors() uint32 {
	argValue := gi.StructFieldGet(inputMessageStruct, recv.Native, "num_vectors")
	value := argValue.Uint32()
	return value
}

// SetFieldNumVectors sets the value of the C field 'num_vectors'.
func (recv *InputMessage) SetFieldNumVectors(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.StructFieldSet(inputMessageStruct, recv.Native, "num_vectors", argValue)
}

// FieldBytesReceived returns the C field 'bytes_received'.
func (recv *InputMessage) FieldBytesReceived() uint64 {
	argValue := gi.StructFieldGet(inputMessageStruct, recv.Native, "bytes_received")
	value := argValue.Uint64()
	return value
}

// SetFieldBytesReceived sets the value of the C field 'bytes_received'.
func (recv *InputMessage) SetFieldBytesReceived(value uint64) {
	var argValue gi.Argument
	argValue.SetUint64(value)
	gi.StructFieldSet(inputMessageStruct, recv.Native, "bytes_received", argValue)
}

// FieldFlags returns the C field 'flags'.
func (recv *InputMessage) FieldFlags() int32 {
	argValue := gi.StructFieldGet(inputMessageStruct, recv.Native, "flags")
	value := argValue.Int32()
	return value
}

// SetFieldFlags sets the value of the C field 'flags'.
func (recv *InputMessage) SetFieldFlags(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.StructFieldSet(inputMessageStruct, recv.Native, "flags", argValue)
}

// UNSUPPORTED : C value 'control_messages' : for field getter : missing Type

// UNSUPPORTED : C value 'control_messages' : for field setter : missing Type

// FieldNumControlMessages returns the C field 'num_control_messages'.
func (recv *InputMessage) FieldNumControlMessages() uint32 {
	argValue := gi.StructFieldGet(inputMessageStruct, recv.Native, "num_control_messages")
	value := argValue.Uint32()
	return value
}

// SetFieldNumControlMessages sets the value of the C field 'num_control_messages'.
func (recv *InputMessage) SetFieldNumControlMessages(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.StructFieldSet(inputMessageStruct, recv.Native, "num_control_messages", argValue)
}

// InputMessageStruct creates an uninitialised InputMessage.
func InputMessageStruct() *InputMessage {
	err := inputMessageStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &InputMessage{}
	structGo.Native = inputMessageStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeInputMessage)
	return structGo
}
func finalizeInputMessage(obj *InputMessage) {
	inputMessageStruct.Free(obj.Native)
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
	Native uintptr
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

	structGo := &InputStreamClass{}
	structGo.Native = inputStreamClassStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeInputStreamClass)
	return structGo
}
func finalizeInputStreamClass(obj *InputStreamClass) {
	inputStreamClassStruct.Free(obj.Native)
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
	Native uintptr
}

// InputStreamPrivateStruct creates an uninitialised InputStreamPrivate.
func InputStreamPrivateStruct() *InputStreamPrivate {
	err := inputStreamPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &InputStreamPrivate{}
	structGo.Native = inputStreamPrivateStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeInputStreamPrivate)
	return structGo
}
func finalizeInputStreamPrivate(obj *InputStreamPrivate) {
	inputStreamPrivateStruct.Free(obj.Native)
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
	Native uintptr
}

// UNSUPPORTED : C value 'buffer' : for field getter : no Go type for 'gpointer'

// UNSUPPORTED : C value 'buffer' : for field setter : no Go type for 'gpointer'

// FieldSize returns the C field 'size'.
func (recv *InputVector) FieldSize() uint64 {
	argValue := gi.StructFieldGet(inputVectorStruct, recv.Native, "size")
	value := argValue.Uint64()
	return value
}

// SetFieldSize sets the value of the C field 'size'.
func (recv *InputVector) SetFieldSize(value uint64) {
	var argValue gi.Argument
	argValue.SetUint64(value)
	gi.StructFieldSet(inputVectorStruct, recv.Native, "size", argValue)
}

// InputVectorStruct creates an uninitialised InputVector.
func InputVectorStruct() *InputVector {
	err := inputVectorStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &InputVector{}
	structGo.Native = inputVectorStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeInputVector)
	return structGo
}
func finalizeInputVector(obj *InputVector) {
	inputVectorStruct.Free(obj.Native)
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
	Native uintptr
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

	structGo := &ListModelInterface{}
	structGo.Native = listModelInterfaceStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeListModelInterface)
	return structGo
}
func finalizeListModelInterface(obj *ListModelInterface) {
	listModelInterfaceStruct.Free(obj.Native)
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
	Native uintptr
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// ListStoreClassStruct creates an uninitialised ListStoreClass.
func ListStoreClassStruct() *ListStoreClass {
	err := listStoreClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ListStoreClass{}
	structGo.Native = listStoreClassStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeListStoreClass)
	return structGo
}
func finalizeListStoreClass(obj *ListStoreClass) {
	listStoreClassStruct.Free(obj.Native)
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
	Native uintptr
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

	structGo := &LoadableIconIface{}
	structGo.Native = loadableIconIfaceStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeLoadableIconIface)
	return structGo
}
func finalizeLoadableIconIface(obj *LoadableIconIface) {
	loadableIconIfaceStruct.Free(obj.Native)
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
	Native uintptr
}

// FieldParentClass returns the C field 'parent_class'.
func (recv *MemoryInputStreamClass) FieldParentClass() *InputStreamClass {
	argValue := gi.StructFieldGet(memoryInputStreamClassStruct, recv.Native, "parent_class")
	value := &InputStreamClass{}
	value.Native = argValue.Pointer()
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *MemoryInputStreamClass) SetFieldParentClass(value *InputStreamClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native)
	gi.StructFieldSet(memoryInputStreamClassStruct, recv.Native, "parent_class", argValue)
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

	structGo := &MemoryInputStreamClass{}
	structGo.Native = memoryInputStreamClassStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeMemoryInputStreamClass)
	return structGo
}
func finalizeMemoryInputStreamClass(obj *MemoryInputStreamClass) {
	memoryInputStreamClassStruct.Free(obj.Native)
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
	Native uintptr
}

// MemoryInputStreamPrivateStruct creates an uninitialised MemoryInputStreamPrivate.
func MemoryInputStreamPrivateStruct() *MemoryInputStreamPrivate {
	err := memoryInputStreamPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &MemoryInputStreamPrivate{}
	structGo.Native = memoryInputStreamPrivateStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeMemoryInputStreamPrivate)
	return structGo
}
func finalizeMemoryInputStreamPrivate(obj *MemoryInputStreamPrivate) {
	memoryInputStreamPrivateStruct.Free(obj.Native)
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
	Native uintptr
}

// FieldParentClass returns the C field 'parent_class'.
func (recv *MemoryOutputStreamClass) FieldParentClass() *OutputStreamClass {
	argValue := gi.StructFieldGet(memoryOutputStreamClassStruct, recv.Native, "parent_class")
	value := &OutputStreamClass{}
	value.Native = argValue.Pointer()
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *MemoryOutputStreamClass) SetFieldParentClass(value *OutputStreamClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native)
	gi.StructFieldSet(memoryOutputStreamClassStruct, recv.Native, "parent_class", argValue)
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

	structGo := &MemoryOutputStreamClass{}
	structGo.Native = memoryOutputStreamClassStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeMemoryOutputStreamClass)
	return structGo
}
func finalizeMemoryOutputStreamClass(obj *MemoryOutputStreamClass) {
	memoryOutputStreamClassStruct.Free(obj.Native)
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
	Native uintptr
}

// MemoryOutputStreamPrivateStruct creates an uninitialised MemoryOutputStreamPrivate.
func MemoryOutputStreamPrivateStruct() *MemoryOutputStreamPrivate {
	err := memoryOutputStreamPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &MemoryOutputStreamPrivate{}
	structGo.Native = memoryOutputStreamPrivateStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeMemoryOutputStreamPrivate)
	return structGo
}
func finalizeMemoryOutputStreamPrivate(obj *MemoryOutputStreamPrivate) {
	memoryOutputStreamPrivateStruct.Free(obj.Native)
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
	Native uintptr
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

	structGo := &MenuAttributeIterClass{}
	structGo.Native = menuAttributeIterClassStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeMenuAttributeIterClass)
	return structGo
}
func finalizeMenuAttributeIterClass(obj *MenuAttributeIterClass) {
	menuAttributeIterClassStruct.Free(obj.Native)
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
	Native uintptr
}

// MenuAttributeIterPrivateStruct creates an uninitialised MenuAttributeIterPrivate.
func MenuAttributeIterPrivateStruct() *MenuAttributeIterPrivate {
	err := menuAttributeIterPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &MenuAttributeIterPrivate{}
	structGo.Native = menuAttributeIterPrivateStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeMenuAttributeIterPrivate)
	return structGo
}
func finalizeMenuAttributeIterPrivate(obj *MenuAttributeIterPrivate) {
	menuAttributeIterPrivateStruct.Free(obj.Native)
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
	Native uintptr
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

	structGo := &MenuLinkIterClass{}
	structGo.Native = menuLinkIterClassStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeMenuLinkIterClass)
	return structGo
}
func finalizeMenuLinkIterClass(obj *MenuLinkIterClass) {
	menuLinkIterClassStruct.Free(obj.Native)
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
	Native uintptr
}

// MenuLinkIterPrivateStruct creates an uninitialised MenuLinkIterPrivate.
func MenuLinkIterPrivateStruct() *MenuLinkIterPrivate {
	err := menuLinkIterPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &MenuLinkIterPrivate{}
	structGo.Native = menuLinkIterPrivateStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeMenuLinkIterPrivate)
	return structGo
}
func finalizeMenuLinkIterPrivate(obj *MenuLinkIterPrivate) {
	menuLinkIterPrivateStruct.Free(obj.Native)
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
	Native uintptr
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

	structGo := &MenuModelClass{}
	structGo.Native = menuModelClassStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeMenuModelClass)
	return structGo
}
func finalizeMenuModelClass(obj *MenuModelClass) {
	menuModelClassStruct.Free(obj.Native)
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
	Native uintptr
}

// MenuModelPrivateStruct creates an uninitialised MenuModelPrivate.
func MenuModelPrivateStruct() *MenuModelPrivate {
	err := menuModelPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &MenuModelPrivate{}
	structGo.Native = menuModelPrivateStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeMenuModelPrivate)
	return structGo
}
func finalizeMenuModelPrivate(obj *MenuModelPrivate) {
	menuModelPrivateStruct.Free(obj.Native)
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
	Native uintptr
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

	structGo := &MountIface{}
	structGo.Native = mountIfaceStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeMountIface)
	return structGo
}
func finalizeMountIface(obj *MountIface) {
	mountIfaceStruct.Free(obj.Native)
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
	Native uintptr
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

	structGo := &MountOperationClass{}
	structGo.Native = mountOperationClassStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeMountOperationClass)
	return structGo
}
func finalizeMountOperationClass(obj *MountOperationClass) {
	mountOperationClassStruct.Free(obj.Native)
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
	Native uintptr
}

// MountOperationPrivateStruct creates an uninitialised MountOperationPrivate.
func MountOperationPrivateStruct() *MountOperationPrivate {
	err := mountOperationPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &MountOperationPrivate{}
	structGo.Native = mountOperationPrivateStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeMountOperationPrivate)
	return structGo
}
func finalizeMountOperationPrivate(obj *MountOperationPrivate) {
	mountOperationPrivateStruct.Free(obj.Native)
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
	Native uintptr
}

// FieldParentClass returns the C field 'parent_class'.
func (recv *NativeSocketAddressClass) FieldParentClass() *SocketAddressClass {
	argValue := gi.StructFieldGet(nativeSocketAddressClassStruct, recv.Native, "parent_class")
	value := &SocketAddressClass{}
	value.Native = argValue.Pointer()
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *NativeSocketAddressClass) SetFieldParentClass(value *SocketAddressClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native)
	gi.StructFieldSet(nativeSocketAddressClassStruct, recv.Native, "parent_class", argValue)
}

// NativeSocketAddressClassStruct creates an uninitialised NativeSocketAddressClass.
func NativeSocketAddressClassStruct() *NativeSocketAddressClass {
	err := nativeSocketAddressClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &NativeSocketAddressClass{}
	structGo.Native = nativeSocketAddressClassStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeNativeSocketAddressClass)
	return structGo
}
func finalizeNativeSocketAddressClass(obj *NativeSocketAddressClass) {
	nativeSocketAddressClassStruct.Free(obj.Native)
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
	Native uintptr
}

// NativeSocketAddressPrivateStruct creates an uninitialised NativeSocketAddressPrivate.
func NativeSocketAddressPrivateStruct() *NativeSocketAddressPrivate {
	err := nativeSocketAddressPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &NativeSocketAddressPrivate{}
	structGo.Native = nativeSocketAddressPrivateStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeNativeSocketAddressPrivate)
	return structGo
}
func finalizeNativeSocketAddressPrivate(obj *NativeSocketAddressPrivate) {
	nativeSocketAddressPrivateStruct.Free(obj.Native)
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
	Native uintptr
}

// FieldParentClass returns the C field 'parent_class'.
func (recv *NativeVolumeMonitorClass) FieldParentClass() *VolumeMonitorClass {
	argValue := gi.StructFieldGet(nativeVolumeMonitorClassStruct, recv.Native, "parent_class")
	value := &VolumeMonitorClass{}
	value.Native = argValue.Pointer()
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *NativeVolumeMonitorClass) SetFieldParentClass(value *VolumeMonitorClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native)
	gi.StructFieldSet(nativeVolumeMonitorClassStruct, recv.Native, "parent_class", argValue)
}

// UNSUPPORTED : C value 'get_mount_for_mount_path' : for field getter : missing Type

// UNSUPPORTED : C value 'get_mount_for_mount_path' : for field setter : missing Type

// NativeVolumeMonitorClassStruct creates an uninitialised NativeVolumeMonitorClass.
func NativeVolumeMonitorClassStruct() *NativeVolumeMonitorClass {
	err := nativeVolumeMonitorClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &NativeVolumeMonitorClass{}
	structGo.Native = nativeVolumeMonitorClassStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeNativeVolumeMonitorClass)
	return structGo
}
func finalizeNativeVolumeMonitorClass(obj *NativeVolumeMonitorClass) {
	nativeVolumeMonitorClassStruct.Free(obj.Native)
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
	Native uintptr
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// NetworkAddressClassStruct creates an uninitialised NetworkAddressClass.
func NetworkAddressClassStruct() *NetworkAddressClass {
	err := networkAddressClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &NetworkAddressClass{}
	structGo.Native = networkAddressClassStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeNetworkAddressClass)
	return structGo
}
func finalizeNetworkAddressClass(obj *NetworkAddressClass) {
	networkAddressClassStruct.Free(obj.Native)
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
	Native uintptr
}

// NetworkAddressPrivateStruct creates an uninitialised NetworkAddressPrivate.
func NetworkAddressPrivateStruct() *NetworkAddressPrivate {
	err := networkAddressPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &NetworkAddressPrivate{}
	structGo.Native = networkAddressPrivateStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeNetworkAddressPrivate)
	return structGo
}
func finalizeNetworkAddressPrivate(obj *NetworkAddressPrivate) {
	networkAddressPrivateStruct.Free(obj.Native)
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
	Native uintptr
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

	structGo := &NetworkMonitorInterface{}
	structGo.Native = networkMonitorInterfaceStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeNetworkMonitorInterface)
	return structGo
}
func finalizeNetworkMonitorInterface(obj *NetworkMonitorInterface) {
	networkMonitorInterfaceStruct.Free(obj.Native)
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
	Native uintptr
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// NetworkServiceClassStruct creates an uninitialised NetworkServiceClass.
func NetworkServiceClassStruct() *NetworkServiceClass {
	err := networkServiceClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &NetworkServiceClass{}
	structGo.Native = networkServiceClassStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeNetworkServiceClass)
	return structGo
}
func finalizeNetworkServiceClass(obj *NetworkServiceClass) {
	networkServiceClassStruct.Free(obj.Native)
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
	Native uintptr
}

// NetworkServicePrivateStruct creates an uninitialised NetworkServicePrivate.
func NetworkServicePrivateStruct() *NetworkServicePrivate {
	err := networkServicePrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &NetworkServicePrivate{}
	structGo.Native = networkServicePrivateStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeNetworkServicePrivate)
	return structGo
}
func finalizeNetworkServicePrivate(obj *NetworkServicePrivate) {
	networkServicePrivateStruct.Free(obj.Native)
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
	Native uintptr
}

// FieldAddress returns the C field 'address'.
func (recv *OutputMessage) FieldAddress() *SocketAddress {
	argValue := gi.StructFieldGet(outputMessageStruct, recv.Native, "address")
	value := &SocketAddress{}
	value.Native = argValue.Pointer()
	return value
}

// SetFieldAddress sets the value of the C field 'address'.
func (recv *OutputMessage) SetFieldAddress(value *SocketAddress) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native)
	gi.StructFieldSet(outputMessageStruct, recv.Native, "address", argValue)
}

// FieldVectors returns the C field 'vectors'.
func (recv *OutputMessage) FieldVectors() *OutputVector {
	argValue := gi.StructFieldGet(outputMessageStruct, recv.Native, "vectors")
	value := &OutputVector{}
	value.Native = argValue.Pointer()
	return value
}

// SetFieldVectors sets the value of the C field 'vectors'.
func (recv *OutputMessage) SetFieldVectors(value *OutputVector) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native)
	gi.StructFieldSet(outputMessageStruct, recv.Native, "vectors", argValue)
}

// FieldNumVectors returns the C field 'num_vectors'.
func (recv *OutputMessage) FieldNumVectors() uint32 {
	argValue := gi.StructFieldGet(outputMessageStruct, recv.Native, "num_vectors")
	value := argValue.Uint32()
	return value
}

// SetFieldNumVectors sets the value of the C field 'num_vectors'.
func (recv *OutputMessage) SetFieldNumVectors(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.StructFieldSet(outputMessageStruct, recv.Native, "num_vectors", argValue)
}

// FieldBytesSent returns the C field 'bytes_sent'.
func (recv *OutputMessage) FieldBytesSent() uint32 {
	argValue := gi.StructFieldGet(outputMessageStruct, recv.Native, "bytes_sent")
	value := argValue.Uint32()
	return value
}

// SetFieldBytesSent sets the value of the C field 'bytes_sent'.
func (recv *OutputMessage) SetFieldBytesSent(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.StructFieldSet(outputMessageStruct, recv.Native, "bytes_sent", argValue)
}

// UNSUPPORTED : C value 'control_messages' : for field getter : missing Type

// UNSUPPORTED : C value 'control_messages' : for field setter : missing Type

// FieldNumControlMessages returns the C field 'num_control_messages'.
func (recv *OutputMessage) FieldNumControlMessages() uint32 {
	argValue := gi.StructFieldGet(outputMessageStruct, recv.Native, "num_control_messages")
	value := argValue.Uint32()
	return value
}

// SetFieldNumControlMessages sets the value of the C field 'num_control_messages'.
func (recv *OutputMessage) SetFieldNumControlMessages(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.StructFieldSet(outputMessageStruct, recv.Native, "num_control_messages", argValue)
}

// OutputMessageStruct creates an uninitialised OutputMessage.
func OutputMessageStruct() *OutputMessage {
	err := outputMessageStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &OutputMessage{}
	structGo.Native = outputMessageStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeOutputMessage)
	return structGo
}
func finalizeOutputMessage(obj *OutputMessage) {
	outputMessageStruct.Free(obj.Native)
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
	Native uintptr
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

	structGo := &OutputStreamClass{}
	structGo.Native = outputStreamClassStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeOutputStreamClass)
	return structGo
}
func finalizeOutputStreamClass(obj *OutputStreamClass) {
	outputStreamClassStruct.Free(obj.Native)
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
	Native uintptr
}

// OutputStreamPrivateStruct creates an uninitialised OutputStreamPrivate.
func OutputStreamPrivateStruct() *OutputStreamPrivate {
	err := outputStreamPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &OutputStreamPrivate{}
	structGo.Native = outputStreamPrivateStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeOutputStreamPrivate)
	return structGo
}
func finalizeOutputStreamPrivate(obj *OutputStreamPrivate) {
	outputStreamPrivateStruct.Free(obj.Native)
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
	Native uintptr
}

// UNSUPPORTED : C value 'buffer' : for field getter : no Go type for 'gpointer'

// UNSUPPORTED : C value 'buffer' : for field setter : no Go type for 'gpointer'

// FieldSize returns the C field 'size'.
func (recv *OutputVector) FieldSize() uint64 {
	argValue := gi.StructFieldGet(outputVectorStruct, recv.Native, "size")
	value := argValue.Uint64()
	return value
}

// SetFieldSize sets the value of the C field 'size'.
func (recv *OutputVector) SetFieldSize(value uint64) {
	var argValue gi.Argument
	argValue.SetUint64(value)
	gi.StructFieldSet(outputVectorStruct, recv.Native, "size", argValue)
}

// OutputVectorStruct creates an uninitialised OutputVector.
func OutputVectorStruct() *OutputVector {
	err := outputVectorStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &OutputVector{}
	structGo.Native = outputVectorStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeOutputVector)
	return structGo
}
func finalizeOutputVector(obj *OutputVector) {
	outputVectorStruct.Free(obj.Native)
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
	Native uintptr
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

	structGo := &PermissionClass{}
	structGo.Native = permissionClassStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizePermissionClass)
	return structGo
}
func finalizePermissionClass(obj *PermissionClass) {
	permissionClassStruct.Free(obj.Native)
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
	Native uintptr
}

// PermissionPrivateStruct creates an uninitialised PermissionPrivate.
func PermissionPrivateStruct() *PermissionPrivate {
	err := permissionPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &PermissionPrivate{}
	structGo.Native = permissionPrivateStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizePermissionPrivate)
	return structGo
}
func finalizePermissionPrivate(obj *PermissionPrivate) {
	permissionPrivateStruct.Free(obj.Native)
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
	Native uintptr
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

	structGo := &PollableInputStreamInterface{}
	structGo.Native = pollableInputStreamInterfaceStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizePollableInputStreamInterface)
	return structGo
}
func finalizePollableInputStreamInterface(obj *PollableInputStreamInterface) {
	pollableInputStreamInterfaceStruct.Free(obj.Native)
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
	Native uintptr
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

	structGo := &PollableOutputStreamInterface{}
	structGo.Native = pollableOutputStreamInterfaceStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizePollableOutputStreamInterface)
	return structGo
}
func finalizePollableOutputStreamInterface(obj *PollableOutputStreamInterface) {
	pollableOutputStreamInterfaceStruct.Free(obj.Native)
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
	Native uintptr
}

// FieldParentClass returns the C field 'parent_class'.
func (recv *ProxyAddressClass) FieldParentClass() *InetSocketAddressClass {
	argValue := gi.StructFieldGet(proxyAddressClassStruct, recv.Native, "parent_class")
	value := &InetSocketAddressClass{}
	value.Native = argValue.Pointer()
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *ProxyAddressClass) SetFieldParentClass(value *InetSocketAddressClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native)
	gi.StructFieldSet(proxyAddressClassStruct, recv.Native, "parent_class", argValue)
}

// ProxyAddressClassStruct creates an uninitialised ProxyAddressClass.
func ProxyAddressClassStruct() *ProxyAddressClass {
	err := proxyAddressClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ProxyAddressClass{}
	structGo.Native = proxyAddressClassStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeProxyAddressClass)
	return structGo
}
func finalizeProxyAddressClass(obj *ProxyAddressClass) {
	proxyAddressClassStruct.Free(obj.Native)
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
	Native uintptr
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

	structGo := &ProxyAddressEnumeratorClass{}
	structGo.Native = proxyAddressEnumeratorClassStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeProxyAddressEnumeratorClass)
	return structGo
}
func finalizeProxyAddressEnumeratorClass(obj *ProxyAddressEnumeratorClass) {
	proxyAddressEnumeratorClassStruct.Free(obj.Native)
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
	Native uintptr
}

// ProxyAddressEnumeratorPrivateStruct creates an uninitialised ProxyAddressEnumeratorPrivate.
func ProxyAddressEnumeratorPrivateStruct() *ProxyAddressEnumeratorPrivate {
	err := proxyAddressEnumeratorPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ProxyAddressEnumeratorPrivate{}
	structGo.Native = proxyAddressEnumeratorPrivateStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeProxyAddressEnumeratorPrivate)
	return structGo
}
func finalizeProxyAddressEnumeratorPrivate(obj *ProxyAddressEnumeratorPrivate) {
	proxyAddressEnumeratorPrivateStruct.Free(obj.Native)
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
	Native uintptr
}

// ProxyAddressPrivateStruct creates an uninitialised ProxyAddressPrivate.
func ProxyAddressPrivateStruct() *ProxyAddressPrivate {
	err := proxyAddressPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ProxyAddressPrivate{}
	structGo.Native = proxyAddressPrivateStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeProxyAddressPrivate)
	return structGo
}
func finalizeProxyAddressPrivate(obj *ProxyAddressPrivate) {
	proxyAddressPrivateStruct.Free(obj.Native)
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
	Native uintptr
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

	structGo := &ProxyInterface{}
	structGo.Native = proxyInterfaceStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeProxyInterface)
	return structGo
}
func finalizeProxyInterface(obj *ProxyInterface) {
	proxyInterfaceStruct.Free(obj.Native)
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
	Native uintptr
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

	structGo := &ProxyResolverInterface{}
	structGo.Native = proxyResolverInterfaceStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeProxyResolverInterface)
	return structGo
}
func finalizeProxyResolverInterface(obj *ProxyResolverInterface) {
	proxyResolverInterfaceStruct.Free(obj.Native)
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
	Native uintptr
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

	structGo := &RemoteActionGroupInterface{}
	structGo.Native = remoteActionGroupInterfaceStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeRemoteActionGroupInterface)
	return structGo
}
func finalizeRemoteActionGroupInterface(obj *RemoteActionGroupInterface) {
	remoteActionGroupInterfaceStruct.Free(obj.Native)
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
	Native uintptr
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

	structGo := &ResolverClass{}
	structGo.Native = resolverClassStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeResolverClass)
	return structGo
}
func finalizeResolverClass(obj *ResolverClass) {
	resolverClassStruct.Free(obj.Native)
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
	Native uintptr
}

// ResolverPrivateStruct creates an uninitialised ResolverPrivate.
func ResolverPrivateStruct() *ResolverPrivate {
	err := resolverPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ResolverPrivate{}
	structGo.Native = resolverPrivateStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeResolverPrivate)
	return structGo
}
func finalizeResolverPrivate(obj *ResolverPrivate) {
	resolverPrivateStruct.Free(obj.Native)
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
	Native uintptr
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
	inArgs[0].SetPointer(recv.Native)

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
	inArgs[0].SetPointer(recv.Native)

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
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := resourceRefFunction_Set()
	if err == nil {
		ret = resourceRefFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Resource{}
	retGo.Native = ret.Pointer()

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
	inArgs[0].SetPointer(recv.Native)

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
	Native uintptr
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

	structGo := &SeekableIface{}
	structGo.Native = seekableIfaceStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeSeekableIface)
	return structGo
}
func finalizeSeekableIface(obj *SeekableIface) {
	seekableIfaceStruct.Free(obj.Native)
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
	Native uintptr
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

	structGo := &SettingsBackendClass{}
	structGo.Native = settingsBackendClassStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeSettingsBackendClass)
	return structGo
}
func finalizeSettingsBackendClass(obj *SettingsBackendClass) {
	settingsBackendClassStruct.Free(obj.Native)
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
	Native uintptr
}

// SettingsBackendPrivateStruct creates an uninitialised SettingsBackendPrivate.
func SettingsBackendPrivateStruct() *SettingsBackendPrivate {
	err := settingsBackendPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &SettingsBackendPrivate{}
	structGo.Native = settingsBackendPrivateStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeSettingsBackendPrivate)
	return structGo
}
func finalizeSettingsBackendPrivate(obj *SettingsBackendPrivate) {
	settingsBackendPrivateStruct.Free(obj.Native)
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
	Native uintptr
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

	structGo := &SettingsClass{}
	structGo.Native = settingsClassStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeSettingsClass)
	return structGo
}
func finalizeSettingsClass(obj *SettingsClass) {
	settingsClassStruct.Free(obj.Native)
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
	Native uintptr
}

// SettingsPrivateStruct creates an uninitialised SettingsPrivate.
func SettingsPrivateStruct() *SettingsPrivate {
	err := settingsPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &SettingsPrivate{}
	structGo.Native = settingsPrivateStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeSettingsPrivate)
	return structGo
}
func finalizeSettingsPrivate(obj *SettingsPrivate) {
	settingsPrivateStruct.Free(obj.Native)
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
	Native uintptr
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
	inArgs[0].SetPointer(recv.Native)

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
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetString(name)

	var ret gi.Argument

	err := settingsSchemaGetKeyFunction_Set()
	if err == nil {
		ret = settingsSchemaGetKeyFunction.Invoke(inArgs[:], nil)
	}

	retGo := &SettingsSchemaKey{}
	retGo.Native = ret.Pointer()

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
	inArgs[0].SetPointer(recv.Native)

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
	inArgs[0].SetPointer(recv.Native)
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
	inArgs[0].SetPointer(recv.Native)

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
	inArgs[0].SetPointer(recv.Native)

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
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := settingsSchemaRefFunction_Set()
	if err == nil {
		ret = settingsSchemaRefFunction.Invoke(inArgs[:], nil)
	}

	retGo := &SettingsSchema{}
	retGo.Native = ret.Pointer()

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
	inArgs[0].SetPointer(recv.Native)

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

	structGo := &SettingsSchema{}
	structGo.Native = settingsSchemaStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeSettingsSchema)
	return structGo
}
func finalizeSettingsSchema(obj *SettingsSchema) {
	settingsSchemaStruct.Free(obj.Native)
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
	Native uintptr
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
	inArgs[0].SetPointer(recv.Native)

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
	inArgs[0].SetPointer(recv.Native)

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
	inArgs[0].SetPointer(recv.Native)

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
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := settingsSchemaKeyRefFunction_Set()
	if err == nil {
		ret = settingsSchemaKeyRefFunction.Invoke(inArgs[:], nil)
	}

	retGo := &SettingsSchemaKey{}
	retGo.Native = ret.Pointer()

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
	inArgs[0].SetPointer(recv.Native)

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

	structGo := &SettingsSchemaKey{}
	structGo.Native = settingsSchemaKeyStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeSettingsSchemaKey)
	return structGo
}
func finalizeSettingsSchemaKey(obj *SettingsSchemaKey) {
	settingsSchemaKeyStruct.Free(obj.Native)
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
	Native uintptr
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
	inArgs[1].SetPointer(parent.Native)
	inArgs[2].SetBoolean(trusted)

	var ret gi.Argument

	err := settingsSchemaSourceNewFromDirectoryFunction_Set()
	if err == nil {
		ret = settingsSchemaSourceNewFromDirectoryFunction.Invoke(inArgs[:], nil)
	}

	retGo := &SettingsSchemaSource{}
	retGo.Native = ret.Pointer()

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
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetString(schemaId)
	inArgs[2].SetBoolean(recursive)

	var ret gi.Argument

	err := settingsSchemaSourceLookupFunction_Set()
	if err == nil {
		ret = settingsSchemaSourceLookupFunction.Invoke(inArgs[:], nil)
	}

	retGo := &SettingsSchema{}
	retGo.Native = ret.Pointer()

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
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := settingsSchemaSourceRefFunction_Set()
	if err == nil {
		ret = settingsSchemaSourceRefFunction.Invoke(inArgs[:], nil)
	}

	retGo := &SettingsSchemaSource{}
	retGo.Native = ret.Pointer()

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
	inArgs[0].SetPointer(recv.Native)

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
	Native uintptr
}

// SimpleActionGroupClassStruct creates an uninitialised SimpleActionGroupClass.
func SimpleActionGroupClassStruct() *SimpleActionGroupClass {
	err := simpleActionGroupClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &SimpleActionGroupClass{}
	structGo.Native = simpleActionGroupClassStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeSimpleActionGroupClass)
	return structGo
}
func finalizeSimpleActionGroupClass(obj *SimpleActionGroupClass) {
	simpleActionGroupClassStruct.Free(obj.Native)
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
	Native uintptr
}

// SimpleActionGroupPrivateStruct creates an uninitialised SimpleActionGroupPrivate.
func SimpleActionGroupPrivateStruct() *SimpleActionGroupPrivate {
	err := simpleActionGroupPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &SimpleActionGroupPrivate{}
	structGo.Native = simpleActionGroupPrivateStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeSimpleActionGroupPrivate)
	return structGo
}
func finalizeSimpleActionGroupPrivate(obj *SimpleActionGroupPrivate) {
	simpleActionGroupPrivateStruct.Free(obj.Native)
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
	Native uintptr
}

// SimpleAsyncResultClassStruct creates an uninitialised SimpleAsyncResultClass.
func SimpleAsyncResultClassStruct() *SimpleAsyncResultClass {
	err := simpleAsyncResultClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &SimpleAsyncResultClass{}
	structGo.Native = simpleAsyncResultClassStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeSimpleAsyncResultClass)
	return structGo
}
func finalizeSimpleAsyncResultClass(obj *SimpleAsyncResultClass) {
	simpleAsyncResultClassStruct.Free(obj.Native)
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
	Native uintptr
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

	structGo := &SimpleProxyResolverClass{}
	structGo.Native = simpleProxyResolverClassStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeSimpleProxyResolverClass)
	return structGo
}
func finalizeSimpleProxyResolverClass(obj *SimpleProxyResolverClass) {
	simpleProxyResolverClassStruct.Free(obj.Native)
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
	Native uintptr
}

// SimpleProxyResolverPrivateStruct creates an uninitialised SimpleProxyResolverPrivate.
func SimpleProxyResolverPrivateStruct() *SimpleProxyResolverPrivate {
	err := simpleProxyResolverPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &SimpleProxyResolverPrivate{}
	structGo.Native = simpleProxyResolverPrivateStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeSimpleProxyResolverPrivate)
	return structGo
}
func finalizeSimpleProxyResolverPrivate(obj *SimpleProxyResolverPrivate) {
	simpleProxyResolverPrivateStruct.Free(obj.Native)
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
	Native uintptr
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

	structGo := &SocketAddressClass{}
	structGo.Native = socketAddressClassStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeSocketAddressClass)
	return structGo
}
func finalizeSocketAddressClass(obj *SocketAddressClass) {
	socketAddressClassStruct.Free(obj.Native)
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
	Native uintptr
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

	structGo := &SocketAddressEnumeratorClass{}
	structGo.Native = socketAddressEnumeratorClassStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeSocketAddressEnumeratorClass)
	return structGo
}
func finalizeSocketAddressEnumeratorClass(obj *SocketAddressEnumeratorClass) {
	socketAddressEnumeratorClassStruct.Free(obj.Native)
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
	Native uintptr
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

	structGo := &SocketClass{}
	structGo.Native = socketClassStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeSocketClass)
	return structGo
}
func finalizeSocketClass(obj *SocketClass) {
	socketClassStruct.Free(obj.Native)
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
	Native uintptr
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

	structGo := &SocketClientClass{}
	structGo.Native = socketClientClassStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeSocketClientClass)
	return structGo
}
func finalizeSocketClientClass(obj *SocketClientClass) {
	socketClientClassStruct.Free(obj.Native)
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
	Native uintptr
}

// SocketClientPrivateStruct creates an uninitialised SocketClientPrivate.
func SocketClientPrivateStruct() *SocketClientPrivate {
	err := socketClientPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &SocketClientPrivate{}
	structGo.Native = socketClientPrivateStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeSocketClientPrivate)
	return structGo
}
func finalizeSocketClientPrivate(obj *SocketClientPrivate) {
	socketClientPrivateStruct.Free(obj.Native)
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
	Native uintptr
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

	structGo := &SocketConnectableIface{}
	structGo.Native = socketConnectableIfaceStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeSocketConnectableIface)
	return structGo
}
func finalizeSocketConnectableIface(obj *SocketConnectableIface) {
	socketConnectableIfaceStruct.Free(obj.Native)
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
	Native uintptr
}

// FieldParentClass returns the C field 'parent_class'.
func (recv *SocketConnectionClass) FieldParentClass() *IOStreamClass {
	argValue := gi.StructFieldGet(socketConnectionClassStruct, recv.Native, "parent_class")
	value := &IOStreamClass{}
	value.Native = argValue.Pointer()
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *SocketConnectionClass) SetFieldParentClass(value *IOStreamClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native)
	gi.StructFieldSet(socketConnectionClassStruct, recv.Native, "parent_class", argValue)
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

	structGo := &SocketConnectionClass{}
	structGo.Native = socketConnectionClassStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeSocketConnectionClass)
	return structGo
}
func finalizeSocketConnectionClass(obj *SocketConnectionClass) {
	socketConnectionClassStruct.Free(obj.Native)
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
	Native uintptr
}

// SocketConnectionPrivateStruct creates an uninitialised SocketConnectionPrivate.
func SocketConnectionPrivateStruct() *SocketConnectionPrivate {
	err := socketConnectionPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &SocketConnectionPrivate{}
	structGo.Native = socketConnectionPrivateStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeSocketConnectionPrivate)
	return structGo
}
func finalizeSocketConnectionPrivate(obj *SocketConnectionPrivate) {
	socketConnectionPrivateStruct.Free(obj.Native)
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
	Native uintptr
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

	structGo := &SocketControlMessageClass{}
	structGo.Native = socketControlMessageClassStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeSocketControlMessageClass)
	return structGo
}
func finalizeSocketControlMessageClass(obj *SocketControlMessageClass) {
	socketControlMessageClassStruct.Free(obj.Native)
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
	Native uintptr
}

// SocketControlMessagePrivateStruct creates an uninitialised SocketControlMessagePrivate.
func SocketControlMessagePrivateStruct() *SocketControlMessagePrivate {
	err := socketControlMessagePrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &SocketControlMessagePrivate{}
	structGo.Native = socketControlMessagePrivateStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeSocketControlMessagePrivate)
	return structGo
}
func finalizeSocketControlMessagePrivate(obj *SocketControlMessagePrivate) {
	socketControlMessagePrivateStruct.Free(obj.Native)
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
	Native uintptr
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

	structGo := &SocketListenerClass{}
	structGo.Native = socketListenerClassStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeSocketListenerClass)
	return structGo
}
func finalizeSocketListenerClass(obj *SocketListenerClass) {
	socketListenerClassStruct.Free(obj.Native)
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
	Native uintptr
}

// SocketListenerPrivateStruct creates an uninitialised SocketListenerPrivate.
func SocketListenerPrivateStruct() *SocketListenerPrivate {
	err := socketListenerPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &SocketListenerPrivate{}
	structGo.Native = socketListenerPrivateStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeSocketListenerPrivate)
	return structGo
}
func finalizeSocketListenerPrivate(obj *SocketListenerPrivate) {
	socketListenerPrivateStruct.Free(obj.Native)
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
	Native uintptr
}

// SocketPrivateStruct creates an uninitialised SocketPrivate.
func SocketPrivateStruct() *SocketPrivate {
	err := socketPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &SocketPrivate{}
	structGo.Native = socketPrivateStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeSocketPrivate)
	return structGo
}
func finalizeSocketPrivate(obj *SocketPrivate) {
	socketPrivateStruct.Free(obj.Native)
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
	Native uintptr
}

// FieldParentClass returns the C field 'parent_class'.
func (recv *SocketServiceClass) FieldParentClass() *SocketListenerClass {
	argValue := gi.StructFieldGet(socketServiceClassStruct, recv.Native, "parent_class")
	value := &SocketListenerClass{}
	value.Native = argValue.Pointer()
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *SocketServiceClass) SetFieldParentClass(value *SocketListenerClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native)
	gi.StructFieldSet(socketServiceClassStruct, recv.Native, "parent_class", argValue)
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

	structGo := &SocketServiceClass{}
	structGo.Native = socketServiceClassStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeSocketServiceClass)
	return structGo
}
func finalizeSocketServiceClass(obj *SocketServiceClass) {
	socketServiceClassStruct.Free(obj.Native)
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
	Native uintptr
}

// SocketServicePrivateStruct creates an uninitialised SocketServicePrivate.
func SocketServicePrivateStruct() *SocketServicePrivate {
	err := socketServicePrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &SocketServicePrivate{}
	structGo.Native = socketServicePrivateStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeSocketServicePrivate)
	return structGo
}
func finalizeSocketServicePrivate(obj *SocketServicePrivate) {
	socketServicePrivateStruct.Free(obj.Native)
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
	Native uintptr
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

	retGo := &SrvTarget{}
	retGo.Native = ret.Pointer()

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
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := srvTargetCopyFunction_Set()
	if err == nil {
		ret = srvTargetCopyFunction.Invoke(inArgs[:], nil)
	}

	retGo := &SrvTarget{}
	retGo.Native = ret.Pointer()

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
	inArgs[0].SetPointer(recv.Native)

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
	inArgs[0].SetPointer(recv.Native)

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
	inArgs[0].SetPointer(recv.Native)

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
	inArgs[0].SetPointer(recv.Native)

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
	inArgs[0].SetPointer(recv.Native)

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
	Native uintptr
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
	inArgs[0].SetPointer(recv.Native)

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
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := staticResourceGetResourceFunction_Set()
	if err == nil {
		ret = staticResourceGetResourceFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Resource{}
	retGo.Native = ret.Pointer()

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
	inArgs[0].SetPointer(recv.Native)

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

	structGo := &StaticResource{}
	structGo.Native = staticResourceStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeStaticResource)
	return structGo
}
func finalizeStaticResource(obj *StaticResource) {
	staticResourceStruct.Free(obj.Native)
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
	Native uintptr
}

// TaskClassStruct creates an uninitialised TaskClass.
func TaskClassStruct() *TaskClass {
	err := taskClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &TaskClass{}
	structGo.Native = taskClassStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeTaskClass)
	return structGo
}
func finalizeTaskClass(obj *TaskClass) {
	taskClassStruct.Free(obj.Native)
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
	Native uintptr
}

// FieldParentClass returns the C field 'parent_class'.
func (recv *TcpConnectionClass) FieldParentClass() *SocketConnectionClass {
	argValue := gi.StructFieldGet(tcpConnectionClassStruct, recv.Native, "parent_class")
	value := &SocketConnectionClass{}
	value.Native = argValue.Pointer()
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *TcpConnectionClass) SetFieldParentClass(value *SocketConnectionClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native)
	gi.StructFieldSet(tcpConnectionClassStruct, recv.Native, "parent_class", argValue)
}

// TcpConnectionClassStruct creates an uninitialised TcpConnectionClass.
func TcpConnectionClassStruct() *TcpConnectionClass {
	err := tcpConnectionClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &TcpConnectionClass{}
	structGo.Native = tcpConnectionClassStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeTcpConnectionClass)
	return structGo
}
func finalizeTcpConnectionClass(obj *TcpConnectionClass) {
	tcpConnectionClassStruct.Free(obj.Native)
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
	Native uintptr
}

// TcpConnectionPrivateStruct creates an uninitialised TcpConnectionPrivate.
func TcpConnectionPrivateStruct() *TcpConnectionPrivate {
	err := tcpConnectionPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &TcpConnectionPrivate{}
	structGo.Native = tcpConnectionPrivateStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeTcpConnectionPrivate)
	return structGo
}
func finalizeTcpConnectionPrivate(obj *TcpConnectionPrivate) {
	tcpConnectionPrivateStruct.Free(obj.Native)
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
	Native uintptr
}

// FieldParentClass returns the C field 'parent_class'.
func (recv *TcpWrapperConnectionClass) FieldParentClass() *TcpConnectionClass {
	argValue := gi.StructFieldGet(tcpWrapperConnectionClassStruct, recv.Native, "parent_class")
	value := &TcpConnectionClass{}
	value.Native = argValue.Pointer()
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *TcpWrapperConnectionClass) SetFieldParentClass(value *TcpConnectionClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native)
	gi.StructFieldSet(tcpWrapperConnectionClassStruct, recv.Native, "parent_class", argValue)
}

// TcpWrapperConnectionClassStruct creates an uninitialised TcpWrapperConnectionClass.
func TcpWrapperConnectionClassStruct() *TcpWrapperConnectionClass {
	err := tcpWrapperConnectionClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &TcpWrapperConnectionClass{}
	structGo.Native = tcpWrapperConnectionClassStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeTcpWrapperConnectionClass)
	return structGo
}
func finalizeTcpWrapperConnectionClass(obj *TcpWrapperConnectionClass) {
	tcpWrapperConnectionClassStruct.Free(obj.Native)
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
	Native uintptr
}

// TcpWrapperConnectionPrivateStruct creates an uninitialised TcpWrapperConnectionPrivate.
func TcpWrapperConnectionPrivateStruct() *TcpWrapperConnectionPrivate {
	err := tcpWrapperConnectionPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &TcpWrapperConnectionPrivate{}
	structGo.Native = tcpWrapperConnectionPrivateStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeTcpWrapperConnectionPrivate)
	return structGo
}
func finalizeTcpWrapperConnectionPrivate(obj *TcpWrapperConnectionPrivate) {
	tcpWrapperConnectionPrivateStruct.Free(obj.Native)
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
	Native uintptr
}

// ThemedIconClassStruct creates an uninitialised ThemedIconClass.
func ThemedIconClassStruct() *ThemedIconClass {
	err := themedIconClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ThemedIconClass{}
	structGo.Native = themedIconClassStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeThemedIconClass)
	return structGo
}
func finalizeThemedIconClass(obj *ThemedIconClass) {
	themedIconClassStruct.Free(obj.Native)
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
	Native uintptr
}

// FieldParentClass returns the C field 'parent_class'.
func (recv *ThreadedSocketServiceClass) FieldParentClass() *SocketServiceClass {
	argValue := gi.StructFieldGet(threadedSocketServiceClassStruct, recv.Native, "parent_class")
	value := &SocketServiceClass{}
	value.Native = argValue.Pointer()
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *ThreadedSocketServiceClass) SetFieldParentClass(value *SocketServiceClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native)
	gi.StructFieldSet(threadedSocketServiceClassStruct, recv.Native, "parent_class", argValue)
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

	structGo := &ThreadedSocketServiceClass{}
	structGo.Native = threadedSocketServiceClassStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeThreadedSocketServiceClass)
	return structGo
}
func finalizeThreadedSocketServiceClass(obj *ThreadedSocketServiceClass) {
	threadedSocketServiceClassStruct.Free(obj.Native)
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
	Native uintptr
}

// ThreadedSocketServicePrivateStruct creates an uninitialised ThreadedSocketServicePrivate.
func ThreadedSocketServicePrivateStruct() *ThreadedSocketServicePrivate {
	err := threadedSocketServicePrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ThreadedSocketServicePrivate{}
	structGo.Native = threadedSocketServicePrivateStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeThreadedSocketServicePrivate)
	return structGo
}
func finalizeThreadedSocketServicePrivate(obj *ThreadedSocketServicePrivate) {
	threadedSocketServicePrivateStruct.Free(obj.Native)
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
	Native uintptr
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

	structGo := &TlsBackendInterface{}
	structGo.Native = tlsBackendInterfaceStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeTlsBackendInterface)
	return structGo
}
func finalizeTlsBackendInterface(obj *TlsBackendInterface) {
	tlsBackendInterfaceStruct.Free(obj.Native)
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
	Native uintptr
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

	structGo := &TlsCertificateClass{}
	structGo.Native = tlsCertificateClassStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeTlsCertificateClass)
	return structGo
}
func finalizeTlsCertificateClass(obj *TlsCertificateClass) {
	tlsCertificateClassStruct.Free(obj.Native)
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
	Native uintptr
}

// TlsCertificatePrivateStruct creates an uninitialised TlsCertificatePrivate.
func TlsCertificatePrivateStruct() *TlsCertificatePrivate {
	err := tlsCertificatePrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &TlsCertificatePrivate{}
	structGo.Native = tlsCertificatePrivateStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeTlsCertificatePrivate)
	return structGo
}
func finalizeTlsCertificatePrivate(obj *TlsCertificatePrivate) {
	tlsCertificatePrivateStruct.Free(obj.Native)
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
	Native uintptr
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

	structGo := &TlsClientConnectionInterface{}
	structGo.Native = tlsClientConnectionInterfaceStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeTlsClientConnectionInterface)
	return structGo
}
func finalizeTlsClientConnectionInterface(obj *TlsClientConnectionInterface) {
	tlsClientConnectionInterfaceStruct.Free(obj.Native)
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
	Native uintptr
}

// FieldParentClass returns the C field 'parent_class'.
func (recv *TlsConnectionClass) FieldParentClass() *IOStreamClass {
	argValue := gi.StructFieldGet(tlsConnectionClassStruct, recv.Native, "parent_class")
	value := &IOStreamClass{}
	value.Native = argValue.Pointer()
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *TlsConnectionClass) SetFieldParentClass(value *IOStreamClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native)
	gi.StructFieldSet(tlsConnectionClassStruct, recv.Native, "parent_class", argValue)
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

	structGo := &TlsConnectionClass{}
	structGo.Native = tlsConnectionClassStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeTlsConnectionClass)
	return structGo
}
func finalizeTlsConnectionClass(obj *TlsConnectionClass) {
	tlsConnectionClassStruct.Free(obj.Native)
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
	Native uintptr
}

// TlsConnectionPrivateStruct creates an uninitialised TlsConnectionPrivate.
func TlsConnectionPrivateStruct() *TlsConnectionPrivate {
	err := tlsConnectionPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &TlsConnectionPrivate{}
	structGo.Native = tlsConnectionPrivateStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeTlsConnectionPrivate)
	return structGo
}
func finalizeTlsConnectionPrivate(obj *TlsConnectionPrivate) {
	tlsConnectionPrivateStruct.Free(obj.Native)
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
	Native uintptr
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

	structGo := &TlsDatabaseClass{}
	structGo.Native = tlsDatabaseClassStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeTlsDatabaseClass)
	return structGo
}
func finalizeTlsDatabaseClass(obj *TlsDatabaseClass) {
	tlsDatabaseClassStruct.Free(obj.Native)
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
	Native uintptr
}

// TlsDatabasePrivateStruct creates an uninitialised TlsDatabasePrivate.
func TlsDatabasePrivateStruct() *TlsDatabasePrivate {
	err := tlsDatabasePrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &TlsDatabasePrivate{}
	structGo.Native = tlsDatabasePrivateStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeTlsDatabasePrivate)
	return structGo
}
func finalizeTlsDatabasePrivate(obj *TlsDatabasePrivate) {
	tlsDatabasePrivateStruct.Free(obj.Native)
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
	Native uintptr
}

// UNSUPPORTED : C value 'g_iface' : for field getter : no Go type for 'GObject.TypeInterface'

// UNSUPPORTED : C value 'g_iface' : for field setter : no Go type for 'GObject.TypeInterface'

// TlsFileDatabaseInterfaceStruct creates an uninitialised TlsFileDatabaseInterface.
func TlsFileDatabaseInterfaceStruct() *TlsFileDatabaseInterface {
	err := tlsFileDatabaseInterfaceStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &TlsFileDatabaseInterface{}
	structGo.Native = tlsFileDatabaseInterfaceStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeTlsFileDatabaseInterface)
	return structGo
}
func finalizeTlsFileDatabaseInterface(obj *TlsFileDatabaseInterface) {
	tlsFileDatabaseInterfaceStruct.Free(obj.Native)
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
	Native uintptr
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

	structGo := &TlsInteractionClass{}
	structGo.Native = tlsInteractionClassStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeTlsInteractionClass)
	return structGo
}
func finalizeTlsInteractionClass(obj *TlsInteractionClass) {
	tlsInteractionClassStruct.Free(obj.Native)
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
	Native uintptr
}

// TlsInteractionPrivateStruct creates an uninitialised TlsInteractionPrivate.
func TlsInteractionPrivateStruct() *TlsInteractionPrivate {
	err := tlsInteractionPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &TlsInteractionPrivate{}
	structGo.Native = tlsInteractionPrivateStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeTlsInteractionPrivate)
	return structGo
}
func finalizeTlsInteractionPrivate(obj *TlsInteractionPrivate) {
	tlsInteractionPrivateStruct.Free(obj.Native)
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
	Native uintptr
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

	structGo := &TlsPasswordClass{}
	structGo.Native = tlsPasswordClassStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeTlsPasswordClass)
	return structGo
}
func finalizeTlsPasswordClass(obj *TlsPasswordClass) {
	tlsPasswordClassStruct.Free(obj.Native)
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
	Native uintptr
}

// TlsPasswordPrivateStruct creates an uninitialised TlsPasswordPrivate.
func TlsPasswordPrivateStruct() *TlsPasswordPrivate {
	err := tlsPasswordPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &TlsPasswordPrivate{}
	structGo.Native = tlsPasswordPrivateStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeTlsPasswordPrivate)
	return structGo
}
func finalizeTlsPasswordPrivate(obj *TlsPasswordPrivate) {
	tlsPasswordPrivateStruct.Free(obj.Native)
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
	Native uintptr
}

// UNSUPPORTED : C value 'g_iface' : for field getter : no Go type for 'GObject.TypeInterface'

// UNSUPPORTED : C value 'g_iface' : for field setter : no Go type for 'GObject.TypeInterface'

// TlsServerConnectionInterfaceStruct creates an uninitialised TlsServerConnectionInterface.
func TlsServerConnectionInterfaceStruct() *TlsServerConnectionInterface {
	err := tlsServerConnectionInterfaceStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &TlsServerConnectionInterface{}
	structGo.Native = tlsServerConnectionInterfaceStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeTlsServerConnectionInterface)
	return structGo
}
func finalizeTlsServerConnectionInterface(obj *TlsServerConnectionInterface) {
	tlsServerConnectionInterfaceStruct.Free(obj.Native)
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
	Native uintptr
}

// FieldParentClass returns the C field 'parent_class'.
func (recv *UnixConnectionClass) FieldParentClass() *SocketConnectionClass {
	argValue := gi.StructFieldGet(unixConnectionClassStruct, recv.Native, "parent_class")
	value := &SocketConnectionClass{}
	value.Native = argValue.Pointer()
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *UnixConnectionClass) SetFieldParentClass(value *SocketConnectionClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native)
	gi.StructFieldSet(unixConnectionClassStruct, recv.Native, "parent_class", argValue)
}

// UnixConnectionClassStruct creates an uninitialised UnixConnectionClass.
func UnixConnectionClassStruct() *UnixConnectionClass {
	err := unixConnectionClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &UnixConnectionClass{}
	structGo.Native = unixConnectionClassStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeUnixConnectionClass)
	return structGo
}
func finalizeUnixConnectionClass(obj *UnixConnectionClass) {
	unixConnectionClassStruct.Free(obj.Native)
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
	Native uintptr
}

// UnixConnectionPrivateStruct creates an uninitialised UnixConnectionPrivate.
func UnixConnectionPrivateStruct() *UnixConnectionPrivate {
	err := unixConnectionPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &UnixConnectionPrivate{}
	structGo.Native = unixConnectionPrivateStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeUnixConnectionPrivate)
	return structGo
}
func finalizeUnixConnectionPrivate(obj *UnixConnectionPrivate) {
	unixConnectionPrivateStruct.Free(obj.Native)
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
	Native uintptr
}

// FieldParentClass returns the C field 'parent_class'.
func (recv *UnixCredentialsMessageClass) FieldParentClass() *SocketControlMessageClass {
	argValue := gi.StructFieldGet(unixCredentialsMessageClassStruct, recv.Native, "parent_class")
	value := &SocketControlMessageClass{}
	value.Native = argValue.Pointer()
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *UnixCredentialsMessageClass) SetFieldParentClass(value *SocketControlMessageClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native)
	gi.StructFieldSet(unixCredentialsMessageClassStruct, recv.Native, "parent_class", argValue)
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

	structGo := &UnixCredentialsMessageClass{}
	structGo.Native = unixCredentialsMessageClassStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeUnixCredentialsMessageClass)
	return structGo
}
func finalizeUnixCredentialsMessageClass(obj *UnixCredentialsMessageClass) {
	unixCredentialsMessageClassStruct.Free(obj.Native)
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
	Native uintptr
}

// UnixCredentialsMessagePrivateStruct creates an uninitialised UnixCredentialsMessagePrivate.
func UnixCredentialsMessagePrivateStruct() *UnixCredentialsMessagePrivate {
	err := unixCredentialsMessagePrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &UnixCredentialsMessagePrivate{}
	structGo.Native = unixCredentialsMessagePrivateStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeUnixCredentialsMessagePrivate)
	return structGo
}
func finalizeUnixCredentialsMessagePrivate(obj *UnixCredentialsMessagePrivate) {
	unixCredentialsMessagePrivateStruct.Free(obj.Native)
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
	Native uintptr
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

	structGo := &UnixFDListClass{}
	structGo.Native = unixFDListClassStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeUnixFDListClass)
	return structGo
}
func finalizeUnixFDListClass(obj *UnixFDListClass) {
	unixFDListClassStruct.Free(obj.Native)
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
	Native uintptr
}

// UnixFDListPrivateStruct creates an uninitialised UnixFDListPrivate.
func UnixFDListPrivateStruct() *UnixFDListPrivate {
	err := unixFDListPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &UnixFDListPrivate{}
	structGo.Native = unixFDListPrivateStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeUnixFDListPrivate)
	return structGo
}
func finalizeUnixFDListPrivate(obj *UnixFDListPrivate) {
	unixFDListPrivateStruct.Free(obj.Native)
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
	Native uintptr
}

// FieldParentClass returns the C field 'parent_class'.
func (recv *UnixFDMessageClass) FieldParentClass() *SocketControlMessageClass {
	argValue := gi.StructFieldGet(unixFDMessageClassStruct, recv.Native, "parent_class")
	value := &SocketControlMessageClass{}
	value.Native = argValue.Pointer()
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *UnixFDMessageClass) SetFieldParentClass(value *SocketControlMessageClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native)
	gi.StructFieldSet(unixFDMessageClassStruct, recv.Native, "parent_class", argValue)
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

	structGo := &UnixFDMessageClass{}
	structGo.Native = unixFDMessageClassStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeUnixFDMessageClass)
	return structGo
}
func finalizeUnixFDMessageClass(obj *UnixFDMessageClass) {
	unixFDMessageClassStruct.Free(obj.Native)
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
	Native uintptr
}

// UnixFDMessagePrivateStruct creates an uninitialised UnixFDMessagePrivate.
func UnixFDMessagePrivateStruct() *UnixFDMessagePrivate {
	err := unixFDMessagePrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &UnixFDMessagePrivate{}
	structGo.Native = unixFDMessagePrivateStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeUnixFDMessagePrivate)
	return structGo
}
func finalizeUnixFDMessagePrivate(obj *UnixFDMessagePrivate) {
	unixFDMessagePrivateStruct.Free(obj.Native)
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
	Native uintptr
}

// FieldParentClass returns the C field 'parent_class'.
func (recv *UnixInputStreamClass) FieldParentClass() *InputStreamClass {
	argValue := gi.StructFieldGet(unixInputStreamClassStruct, recv.Native, "parent_class")
	value := &InputStreamClass{}
	value.Native = argValue.Pointer()
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *UnixInputStreamClass) SetFieldParentClass(value *InputStreamClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native)
	gi.StructFieldSet(unixInputStreamClassStruct, recv.Native, "parent_class", argValue)
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

	structGo := &UnixInputStreamClass{}
	structGo.Native = unixInputStreamClassStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeUnixInputStreamClass)
	return structGo
}
func finalizeUnixInputStreamClass(obj *UnixInputStreamClass) {
	unixInputStreamClassStruct.Free(obj.Native)
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
	Native uintptr
}

// UnixInputStreamPrivateStruct creates an uninitialised UnixInputStreamPrivate.
func UnixInputStreamPrivateStruct() *UnixInputStreamPrivate {
	err := unixInputStreamPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &UnixInputStreamPrivate{}
	structGo.Native = unixInputStreamPrivateStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeUnixInputStreamPrivate)
	return structGo
}
func finalizeUnixInputStreamPrivate(obj *UnixInputStreamPrivate) {
	unixInputStreamPrivateStruct.Free(obj.Native)
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
	Native uintptr
}

// UnixMountEntryStruct creates an uninitialised UnixMountEntry.
func UnixMountEntryStruct() *UnixMountEntry {
	err := unixMountEntryStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &UnixMountEntry{}
	structGo.Native = unixMountEntryStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeUnixMountEntry)
	return structGo
}
func finalizeUnixMountEntry(obj *UnixMountEntry) {
	unixMountEntryStruct.Free(obj.Native)
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
	Native uintptr
}

// UnixMountMonitorClassStruct creates an uninitialised UnixMountMonitorClass.
func UnixMountMonitorClassStruct() *UnixMountMonitorClass {
	err := unixMountMonitorClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &UnixMountMonitorClass{}
	structGo.Native = unixMountMonitorClassStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeUnixMountMonitorClass)
	return structGo
}
func finalizeUnixMountMonitorClass(obj *UnixMountMonitorClass) {
	unixMountMonitorClassStruct.Free(obj.Native)
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
	Native uintptr
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
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetPointer(mount2.Native)

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
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := unixMountPointCopyFunction_Set()
	if err == nil {
		ret = unixMountPointCopyFunction.Invoke(inArgs[:], nil)
	}

	retGo := &UnixMountPoint{}
	retGo.Native = ret.Pointer()

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
	inArgs[0].SetPointer(recv.Native)

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
	inArgs[0].SetPointer(recv.Native)

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
	inArgs[0].SetPointer(recv.Native)

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
	inArgs[0].SetPointer(recv.Native)

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
	inArgs[0].SetPointer(recv.Native)

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
	inArgs[0].SetPointer(recv.Native)

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
	inArgs[0].SetPointer(recv.Native)

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
	inArgs[0].SetPointer(recv.Native)

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
	inArgs[0].SetPointer(recv.Native)

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
	inArgs[0].SetPointer(recv.Native)

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

	structGo := &UnixMountPoint{}
	structGo.Native = unixMountPointStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeUnixMountPoint)
	return structGo
}
func finalizeUnixMountPoint(obj *UnixMountPoint) {
	unixMountPointStruct.Free(obj.Native)
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
	Native uintptr
}

// FieldParentClass returns the C field 'parent_class'.
func (recv *UnixOutputStreamClass) FieldParentClass() *OutputStreamClass {
	argValue := gi.StructFieldGet(unixOutputStreamClassStruct, recv.Native, "parent_class")
	value := &OutputStreamClass{}
	value.Native = argValue.Pointer()
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *UnixOutputStreamClass) SetFieldParentClass(value *OutputStreamClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native)
	gi.StructFieldSet(unixOutputStreamClassStruct, recv.Native, "parent_class", argValue)
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

	structGo := &UnixOutputStreamClass{}
	structGo.Native = unixOutputStreamClassStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeUnixOutputStreamClass)
	return structGo
}
func finalizeUnixOutputStreamClass(obj *UnixOutputStreamClass) {
	unixOutputStreamClassStruct.Free(obj.Native)
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
	Native uintptr
}

// UnixOutputStreamPrivateStruct creates an uninitialised UnixOutputStreamPrivate.
func UnixOutputStreamPrivateStruct() *UnixOutputStreamPrivate {
	err := unixOutputStreamPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &UnixOutputStreamPrivate{}
	structGo.Native = unixOutputStreamPrivateStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeUnixOutputStreamPrivate)
	return structGo
}
func finalizeUnixOutputStreamPrivate(obj *UnixOutputStreamPrivate) {
	unixOutputStreamPrivateStruct.Free(obj.Native)
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
	Native uintptr
}

// FieldParentClass returns the C field 'parent_class'.
func (recv *UnixSocketAddressClass) FieldParentClass() *SocketAddressClass {
	argValue := gi.StructFieldGet(unixSocketAddressClassStruct, recv.Native, "parent_class")
	value := &SocketAddressClass{}
	value.Native = argValue.Pointer()
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *UnixSocketAddressClass) SetFieldParentClass(value *SocketAddressClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native)
	gi.StructFieldSet(unixSocketAddressClassStruct, recv.Native, "parent_class", argValue)
}

// UnixSocketAddressClassStruct creates an uninitialised UnixSocketAddressClass.
func UnixSocketAddressClassStruct() *UnixSocketAddressClass {
	err := unixSocketAddressClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &UnixSocketAddressClass{}
	structGo.Native = unixSocketAddressClassStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeUnixSocketAddressClass)
	return structGo
}
func finalizeUnixSocketAddressClass(obj *UnixSocketAddressClass) {
	unixSocketAddressClassStruct.Free(obj.Native)
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
	Native uintptr
}

// UnixSocketAddressPrivateStruct creates an uninitialised UnixSocketAddressPrivate.
func UnixSocketAddressPrivateStruct() *UnixSocketAddressPrivate {
	err := unixSocketAddressPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &UnixSocketAddressPrivate{}
	structGo.Native = unixSocketAddressPrivateStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeUnixSocketAddressPrivate)
	return structGo
}
func finalizeUnixSocketAddressPrivate(obj *UnixSocketAddressPrivate) {
	unixSocketAddressPrivateStruct.Free(obj.Native)
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
	Native uintptr
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

	structGo := &VfsClass{}
	structGo.Native = vfsClassStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeVfsClass)
	return structGo
}
func finalizeVfsClass(obj *VfsClass) {
	vfsClassStruct.Free(obj.Native)
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
	Native uintptr
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

	structGo := &VolumeIface{}
	structGo.Native = volumeIfaceStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeVolumeIface)
	return structGo
}
func finalizeVolumeIface(obj *VolumeIface) {
	volumeIfaceStruct.Free(obj.Native)
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
	Native uintptr
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

	structGo := &VolumeMonitorClass{}
	structGo.Native = volumeMonitorClassStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeVolumeMonitorClass)
	return structGo
}
func finalizeVolumeMonitorClass(obj *VolumeMonitorClass) {
	volumeMonitorClassStruct.Free(obj.Native)
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
	Native uintptr
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// ZlibCompressorClassStruct creates an uninitialised ZlibCompressorClass.
func ZlibCompressorClassStruct() *ZlibCompressorClass {
	err := zlibCompressorClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ZlibCompressorClass{}
	structGo.Native = zlibCompressorClassStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeZlibCompressorClass)
	return structGo
}
func finalizeZlibCompressorClass(obj *ZlibCompressorClass) {
	zlibCompressorClassStruct.Free(obj.Native)
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
	Native uintptr
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// ZlibDecompressorClassStruct creates an uninitialised ZlibDecompressorClass.
func ZlibDecompressorClassStruct() *ZlibDecompressorClass {
	err := zlibDecompressorClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ZlibDecompressorClass{}
	structGo.Native = zlibDecompressorClassStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeZlibDecompressorClass)
	return structGo
}
func finalizeZlibDecompressorClass(obj *ZlibDecompressorClass) {
	zlibDecompressorClassStruct.Free(obj.Native)
}
