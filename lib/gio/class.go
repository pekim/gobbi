// Code generated - DO NOT EDIT.

package gio

import (
	gi "github.com/pekim/gobbi/internal/cgo/gi"
	glib "github.com/pekim/gobbi/lib/glib"
	gobject "github.com/pekim/gobbi/lib/gobject"
	"runtime"
	"sync"
	"unsafe"
)

var appInfoMonitorObject *gi.Object
var appInfoMonitorObject_Once sync.Once

func appInfoMonitorObject_Set() error {
	var err error
	appInfoMonitorObject_Once.Do(func() {
		appInfoMonitorObject, err = gi.ObjectNew("Gio", "AppInfoMonitor")
	})
	return err
}

type AppInfoMonitor struct {
	native unsafe.Pointer
}

func AppInfoMonitorNewFromNative(native unsafe.Pointer) *AppInfoMonitor {
	instance := &AppInfoMonitor{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *AppInfoMonitor) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToAppInfoMonitor down casts any arbitrary Object to AppInfoMonitor.
Exercise care, as this is a potentially dangerous function
if the Object is not a AppInfoMonitor.
*/
func CastToAppInfoMonitor(object *gobject.Object) *AppInfoMonitor {
	return AppInfoMonitorNewFromNative(object.Native())
}

// Equals compares this AppInfoMonitor with another AppInfoMonitor, and returns true if they represent the same GObject.
func (recv *AppInfoMonitor) Equals(other *AppInfoMonitor) bool {
	return other.Native() == recv.Native()
}

func (recv *AppInfoMonitor) Native() unsafe.Pointer {
	return recv.native
}

var appLaunchContextObject *gi.Object
var appLaunchContextObject_Once sync.Once

func appLaunchContextObject_Set() error {
	var err error
	appLaunchContextObject_Once.Do(func() {
		appLaunchContextObject, err = gi.ObjectNew("Gio", "AppLaunchContext")
	})
	return err
}

type AppLaunchContext struct {
	native unsafe.Pointer
}

func AppLaunchContextNewFromNative(native unsafe.Pointer) *AppLaunchContext {
	instance := &AppLaunchContext{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *AppLaunchContext) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToAppLaunchContext down casts any arbitrary Object to AppLaunchContext.
Exercise care, as this is a potentially dangerous function
if the Object is not a AppLaunchContext.
*/
func CastToAppLaunchContext(object *gobject.Object) *AppLaunchContext {
	return AppLaunchContextNewFromNative(object.Native())
}

// Equals compares this AppLaunchContext with another AppLaunchContext, and returns true if they represent the same GObject.
func (recv *AppLaunchContext) Equals(other *AppLaunchContext) bool {
	return other.Native() == recv.Native()
}

func (recv *AppLaunchContext) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *AppLaunchContext) FieldParentInstance() *gobject.Object {
	argValue := gi.ObjectFieldGet(appLaunchContextObject, recv.Native(), "parent_instance")
	value := gobject.ObjectNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *AppLaunchContext) SetFieldParentInstance(value *gobject.Object) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(appLaunchContextObject, recv.Native(), "parent_instance", argValue)
}

var appLaunchContextNewFunction *gi.Function
var appLaunchContextNewFunction_Once sync.Once

func appLaunchContextNewFunction_Set() error {
	var err error
	appLaunchContextNewFunction_Once.Do(func() {
		err = appLaunchContextObject_Set()
		if err != nil {
			return
		}
		appLaunchContextNewFunction, err = appLaunchContextObject.InvokerNew("new")
	})
	return err
}

// AppLaunchContextNew is a representation of the C type g_app_launch_context_new.
func AppLaunchContextNew() *AppLaunchContext {

	var ret gi.Argument

	err := appLaunchContextNewFunction_Set()
	if err == nil {
		ret = appLaunchContextNewFunction.Invoke(nil, nil)
	}

	retGo := AppLaunchContextNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

// UNSUPPORTED : C value 'g_app_launch_context_get_display' : parameter 'info' of type 'AppInfo' not supported

var appLaunchContextGetEnvironmentFunction *gi.Function
var appLaunchContextGetEnvironmentFunction_Once sync.Once

func appLaunchContextGetEnvironmentFunction_Set() error {
	var err error
	appLaunchContextGetEnvironmentFunction_Once.Do(func() {
		err = appLaunchContextObject_Set()
		if err != nil {
			return
		}
		appLaunchContextGetEnvironmentFunction, err = appLaunchContextObject.InvokerNew("get_environment")
	})
	return err
}

// GetEnvironment is a representation of the C type g_app_launch_context_get_environment.
func (recv *AppLaunchContext) GetEnvironment() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := appLaunchContextGetEnvironmentFunction_Set()
	if err == nil {
		appLaunchContextGetEnvironmentFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'g_app_launch_context_get_startup_notify_id' : parameter 'info' of type 'AppInfo' not supported

var appLaunchContextLaunchFailedFunction *gi.Function
var appLaunchContextLaunchFailedFunction_Once sync.Once

func appLaunchContextLaunchFailedFunction_Set() error {
	var err error
	appLaunchContextLaunchFailedFunction_Once.Do(func() {
		err = appLaunchContextObject_Set()
		if err != nil {
			return
		}
		appLaunchContextLaunchFailedFunction, err = appLaunchContextObject.InvokerNew("launch_failed")
	})
	return err
}

// LaunchFailed is a representation of the C type g_app_launch_context_launch_failed.
func (recv *AppLaunchContext) LaunchFailed(startupNotifyId string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(startupNotifyId)

	err := appLaunchContextLaunchFailedFunction_Set()
	if err == nil {
		appLaunchContextLaunchFailedFunction.Invoke(inArgs[:], nil)
	}

	return
}

var appLaunchContextSetenvFunction *gi.Function
var appLaunchContextSetenvFunction_Once sync.Once

func appLaunchContextSetenvFunction_Set() error {
	var err error
	appLaunchContextSetenvFunction_Once.Do(func() {
		err = appLaunchContextObject_Set()
		if err != nil {
			return
		}
		appLaunchContextSetenvFunction, err = appLaunchContextObject.InvokerNew("setenv")
	})
	return err
}

// Setenv is a representation of the C type g_app_launch_context_setenv.
func (recv *AppLaunchContext) Setenv(variable string, value string) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(variable)
	inArgs[2].SetString(value)

	err := appLaunchContextSetenvFunction_Set()
	if err == nil {
		appLaunchContextSetenvFunction.Invoke(inArgs[:], nil)
	}

	return
}

var appLaunchContextUnsetenvFunction *gi.Function
var appLaunchContextUnsetenvFunction_Once sync.Once

func appLaunchContextUnsetenvFunction_Set() error {
	var err error
	appLaunchContextUnsetenvFunction_Once.Do(func() {
		err = appLaunchContextObject_Set()
		if err != nil {
			return
		}
		appLaunchContextUnsetenvFunction, err = appLaunchContextObject.InvokerNew("unsetenv")
	})
	return err
}

// Unsetenv is a representation of the C type g_app_launch_context_unsetenv.
func (recv *AppLaunchContext) Unsetenv(variable string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(variable)

	err := appLaunchContextUnsetenvFunction_Set()
	if err == nil {
		appLaunchContextUnsetenvFunction.Invoke(inArgs[:], nil)
	}

	return
}

var applicationObject *gi.Object
var applicationObject_Once sync.Once

func applicationObject_Set() error {
	var err error
	applicationObject_Once.Do(func() {
		applicationObject, err = gi.ObjectNew("Gio", "Application")
	})
	return err
}

type Application struct {
	native unsafe.Pointer
}

func ApplicationNewFromNative(native unsafe.Pointer) *Application {
	instance := &Application{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *Application) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToApplication down casts any arbitrary Object to Application.
Exercise care, as this is a potentially dangerous function
if the Object is not a Application.
*/
func CastToApplication(object *gobject.Object) *Application {
	return ApplicationNewFromNative(object.Native())
}

// Equals compares this Application with another Application, and returns true if they represent the same GObject.
func (recv *Application) Equals(other *Application) bool {
	return other.Native() == recv.Native()
}

func (recv *Application) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'g_application_new' : parameter 'flags' of type 'ApplicationFlags' not supported

var applicationActivateFunction *gi.Function
var applicationActivateFunction_Once sync.Once

func applicationActivateFunction_Set() error {
	var err error
	applicationActivateFunction_Once.Do(func() {
		err = applicationObject_Set()
		if err != nil {
			return
		}
		applicationActivateFunction, err = applicationObject.InvokerNew("activate")
	})
	return err
}

// Activate is a representation of the C type g_application_activate.
func (recv *Application) Activate() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := applicationActivateFunction_Set()
	if err == nil {
		applicationActivateFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'g_application_add_main_option' : parameter 'flags' of type 'GLib.OptionFlags' not supported

// UNSUPPORTED : C value 'g_application_add_main_option_entries' : parameter 'entries' of type 'nil' not supported

// UNSUPPORTED : C value 'g_application_add_option_group' : parameter 'group' of type 'GLib.OptionGroup' not supported

var applicationBindBusyPropertyFunction *gi.Function
var applicationBindBusyPropertyFunction_Once sync.Once

func applicationBindBusyPropertyFunction_Set() error {
	var err error
	applicationBindBusyPropertyFunction_Once.Do(func() {
		err = applicationObject_Set()
		if err != nil {
			return
		}
		applicationBindBusyPropertyFunction, err = applicationObject.InvokerNew("bind_busy_property")
	})
	return err
}

// BindBusyProperty is a representation of the C type g_application_bind_busy_property.
func (recv *Application) BindBusyProperty(object_ *gobject.Object, property string) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(object_.Native())
	inArgs[2].SetString(property)

	err := applicationBindBusyPropertyFunction_Set()
	if err == nil {
		applicationBindBusyPropertyFunction.Invoke(inArgs[:], nil)
	}

	return
}

var applicationGetApplicationIdFunction *gi.Function
var applicationGetApplicationIdFunction_Once sync.Once

func applicationGetApplicationIdFunction_Set() error {
	var err error
	applicationGetApplicationIdFunction_Once.Do(func() {
		err = applicationObject_Set()
		if err != nil {
			return
		}
		applicationGetApplicationIdFunction, err = applicationObject.InvokerNew("get_application_id")
	})
	return err
}

// GetApplicationId is a representation of the C type g_application_get_application_id.
func (recv *Application) GetApplicationId() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := applicationGetApplicationIdFunction_Set()
	if err == nil {
		ret = applicationGetApplicationIdFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var applicationGetDbusConnectionFunction *gi.Function
var applicationGetDbusConnectionFunction_Once sync.Once

func applicationGetDbusConnectionFunction_Set() error {
	var err error
	applicationGetDbusConnectionFunction_Once.Do(func() {
		err = applicationObject_Set()
		if err != nil {
			return
		}
		applicationGetDbusConnectionFunction, err = applicationObject.InvokerNew("get_dbus_connection")
	})
	return err
}

// GetDbusConnection is a representation of the C type g_application_get_dbus_connection.
func (recv *Application) GetDbusConnection() *DBusConnection {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := applicationGetDbusConnectionFunction_Set()
	if err == nil {
		ret = applicationGetDbusConnectionFunction.Invoke(inArgs[:], nil)
	}

	retGo := DBusConnectionNewFromNative(ret.Pointer())

	return retGo
}

var applicationGetDbusObjectPathFunction *gi.Function
var applicationGetDbusObjectPathFunction_Once sync.Once

func applicationGetDbusObjectPathFunction_Set() error {
	var err error
	applicationGetDbusObjectPathFunction_Once.Do(func() {
		err = applicationObject_Set()
		if err != nil {
			return
		}
		applicationGetDbusObjectPathFunction, err = applicationObject.InvokerNew("get_dbus_object_path")
	})
	return err
}

// GetDbusObjectPath is a representation of the C type g_application_get_dbus_object_path.
func (recv *Application) GetDbusObjectPath() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := applicationGetDbusObjectPathFunction_Set()
	if err == nil {
		ret = applicationGetDbusObjectPathFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'g_application_get_flags' : return type 'ApplicationFlags' not supported

var applicationGetInactivityTimeoutFunction *gi.Function
var applicationGetInactivityTimeoutFunction_Once sync.Once

func applicationGetInactivityTimeoutFunction_Set() error {
	var err error
	applicationGetInactivityTimeoutFunction_Once.Do(func() {
		err = applicationObject_Set()
		if err != nil {
			return
		}
		applicationGetInactivityTimeoutFunction, err = applicationObject.InvokerNew("get_inactivity_timeout")
	})
	return err
}

// GetInactivityTimeout is a representation of the C type g_application_get_inactivity_timeout.
func (recv *Application) GetInactivityTimeout() uint32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := applicationGetInactivityTimeoutFunction_Set()
	if err == nil {
		ret = applicationGetInactivityTimeoutFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo
}

var applicationGetIsBusyFunction *gi.Function
var applicationGetIsBusyFunction_Once sync.Once

func applicationGetIsBusyFunction_Set() error {
	var err error
	applicationGetIsBusyFunction_Once.Do(func() {
		err = applicationObject_Set()
		if err != nil {
			return
		}
		applicationGetIsBusyFunction, err = applicationObject.InvokerNew("get_is_busy")
	})
	return err
}

// GetIsBusy is a representation of the C type g_application_get_is_busy.
func (recv *Application) GetIsBusy() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := applicationGetIsBusyFunction_Set()
	if err == nil {
		ret = applicationGetIsBusyFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var applicationGetIsRegisteredFunction *gi.Function
var applicationGetIsRegisteredFunction_Once sync.Once

func applicationGetIsRegisteredFunction_Set() error {
	var err error
	applicationGetIsRegisteredFunction_Once.Do(func() {
		err = applicationObject_Set()
		if err != nil {
			return
		}
		applicationGetIsRegisteredFunction, err = applicationObject.InvokerNew("get_is_registered")
	})
	return err
}

// GetIsRegistered is a representation of the C type g_application_get_is_registered.
func (recv *Application) GetIsRegistered() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := applicationGetIsRegisteredFunction_Set()
	if err == nil {
		ret = applicationGetIsRegisteredFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var applicationGetIsRemoteFunction *gi.Function
var applicationGetIsRemoteFunction_Once sync.Once

func applicationGetIsRemoteFunction_Set() error {
	var err error
	applicationGetIsRemoteFunction_Once.Do(func() {
		err = applicationObject_Set()
		if err != nil {
			return
		}
		applicationGetIsRemoteFunction, err = applicationObject.InvokerNew("get_is_remote")
	})
	return err
}

// GetIsRemote is a representation of the C type g_application_get_is_remote.
func (recv *Application) GetIsRemote() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := applicationGetIsRemoteFunction_Set()
	if err == nil {
		ret = applicationGetIsRemoteFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var applicationGetResourceBasePathFunction *gi.Function
var applicationGetResourceBasePathFunction_Once sync.Once

func applicationGetResourceBasePathFunction_Set() error {
	var err error
	applicationGetResourceBasePathFunction_Once.Do(func() {
		err = applicationObject_Set()
		if err != nil {
			return
		}
		applicationGetResourceBasePathFunction, err = applicationObject.InvokerNew("get_resource_base_path")
	})
	return err
}

// GetResourceBasePath is a representation of the C type g_application_get_resource_base_path.
func (recv *Application) GetResourceBasePath() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := applicationGetResourceBasePathFunction_Set()
	if err == nil {
		ret = applicationGetResourceBasePathFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var applicationHoldFunction *gi.Function
var applicationHoldFunction_Once sync.Once

func applicationHoldFunction_Set() error {
	var err error
	applicationHoldFunction_Once.Do(func() {
		err = applicationObject_Set()
		if err != nil {
			return
		}
		applicationHoldFunction, err = applicationObject.InvokerNew("hold")
	})
	return err
}

// Hold is a representation of the C type g_application_hold.
func (recv *Application) Hold() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := applicationHoldFunction_Set()
	if err == nil {
		applicationHoldFunction.Invoke(inArgs[:], nil)
	}

	return
}

var applicationMarkBusyFunction *gi.Function
var applicationMarkBusyFunction_Once sync.Once

func applicationMarkBusyFunction_Set() error {
	var err error
	applicationMarkBusyFunction_Once.Do(func() {
		err = applicationObject_Set()
		if err != nil {
			return
		}
		applicationMarkBusyFunction, err = applicationObject.InvokerNew("mark_busy")
	})
	return err
}

// MarkBusy is a representation of the C type g_application_mark_busy.
func (recv *Application) MarkBusy() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := applicationMarkBusyFunction_Set()
	if err == nil {
		applicationMarkBusyFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'g_application_open' : parameter 'files' of type 'nil' not supported

var applicationQuitFunction *gi.Function
var applicationQuitFunction_Once sync.Once

func applicationQuitFunction_Set() error {
	var err error
	applicationQuitFunction_Once.Do(func() {
		err = applicationObject_Set()
		if err != nil {
			return
		}
		applicationQuitFunction, err = applicationObject.InvokerNew("quit")
	})
	return err
}

// Quit is a representation of the C type g_application_quit.
func (recv *Application) Quit() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := applicationQuitFunction_Set()
	if err == nil {
		applicationQuitFunction.Invoke(inArgs[:], nil)
	}

	return
}

var applicationRegisterFunction *gi.Function
var applicationRegisterFunction_Once sync.Once

func applicationRegisterFunction_Set() error {
	var err error
	applicationRegisterFunction_Once.Do(func() {
		err = applicationObject_Set()
		if err != nil {
			return
		}
		applicationRegisterFunction, err = applicationObject.InvokerNew("register")
	})
	return err
}

// Register is a representation of the C type g_application_register.
func (recv *Application) Register(cancellable *Cancellable) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(cancellable.Native())

	var ret gi.Argument

	err := applicationRegisterFunction_Set()
	if err == nil {
		ret = applicationRegisterFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var applicationReleaseFunction *gi.Function
var applicationReleaseFunction_Once sync.Once

func applicationReleaseFunction_Set() error {
	var err error
	applicationReleaseFunction_Once.Do(func() {
		err = applicationObject_Set()
		if err != nil {
			return
		}
		applicationReleaseFunction, err = applicationObject.InvokerNew("release")
	})
	return err
}

// Release is a representation of the C type g_application_release.
func (recv *Application) Release() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := applicationReleaseFunction_Set()
	if err == nil {
		applicationReleaseFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'g_application_run' : parameter 'argv' of type 'nil' not supported

var applicationSendNotificationFunction *gi.Function
var applicationSendNotificationFunction_Once sync.Once

func applicationSendNotificationFunction_Set() error {
	var err error
	applicationSendNotificationFunction_Once.Do(func() {
		err = applicationObject_Set()
		if err != nil {
			return
		}
		applicationSendNotificationFunction, err = applicationObject.InvokerNew("send_notification")
	})
	return err
}

// SendNotification is a representation of the C type g_application_send_notification.
func (recv *Application) SendNotification(id string, notification *Notification) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(id)
	inArgs[2].SetPointer(notification.Native())

	err := applicationSendNotificationFunction_Set()
	if err == nil {
		applicationSendNotificationFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'g_application_set_action_group' : parameter 'action_group' of type 'ActionGroup' not supported

var applicationSetApplicationIdFunction *gi.Function
var applicationSetApplicationIdFunction_Once sync.Once

func applicationSetApplicationIdFunction_Set() error {
	var err error
	applicationSetApplicationIdFunction_Once.Do(func() {
		err = applicationObject_Set()
		if err != nil {
			return
		}
		applicationSetApplicationIdFunction, err = applicationObject.InvokerNew("set_application_id")
	})
	return err
}

// SetApplicationId is a representation of the C type g_application_set_application_id.
func (recv *Application) SetApplicationId(applicationId string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(applicationId)

	err := applicationSetApplicationIdFunction_Set()
	if err == nil {
		applicationSetApplicationIdFunction.Invoke(inArgs[:], nil)
	}

	return
}

var applicationSetDefaultFunction *gi.Function
var applicationSetDefaultFunction_Once sync.Once

func applicationSetDefaultFunction_Set() error {
	var err error
	applicationSetDefaultFunction_Once.Do(func() {
		err = applicationObject_Set()
		if err != nil {
			return
		}
		applicationSetDefaultFunction, err = applicationObject.InvokerNew("set_default")
	})
	return err
}

// SetDefault is a representation of the C type g_application_set_default.
func (recv *Application) SetDefault() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := applicationSetDefaultFunction_Set()
	if err == nil {
		applicationSetDefaultFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'g_application_set_flags' : parameter 'flags' of type 'ApplicationFlags' not supported

var applicationSetInactivityTimeoutFunction *gi.Function
var applicationSetInactivityTimeoutFunction_Once sync.Once

func applicationSetInactivityTimeoutFunction_Set() error {
	var err error
	applicationSetInactivityTimeoutFunction_Once.Do(func() {
		err = applicationObject_Set()
		if err != nil {
			return
		}
		applicationSetInactivityTimeoutFunction, err = applicationObject.InvokerNew("set_inactivity_timeout")
	})
	return err
}

// SetInactivityTimeout is a representation of the C type g_application_set_inactivity_timeout.
func (recv *Application) SetInactivityTimeout(inactivityTimeout uint32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetUint32(inactivityTimeout)

	err := applicationSetInactivityTimeoutFunction_Set()
	if err == nil {
		applicationSetInactivityTimeoutFunction.Invoke(inArgs[:], nil)
	}

	return
}

var applicationSetOptionContextDescriptionFunction *gi.Function
var applicationSetOptionContextDescriptionFunction_Once sync.Once

func applicationSetOptionContextDescriptionFunction_Set() error {
	var err error
	applicationSetOptionContextDescriptionFunction_Once.Do(func() {
		err = applicationObject_Set()
		if err != nil {
			return
		}
		applicationSetOptionContextDescriptionFunction, err = applicationObject.InvokerNew("set_option_context_description")
	})
	return err
}

// SetOptionContextDescription is a representation of the C type g_application_set_option_context_description.
func (recv *Application) SetOptionContextDescription(description string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(description)

	err := applicationSetOptionContextDescriptionFunction_Set()
	if err == nil {
		applicationSetOptionContextDescriptionFunction.Invoke(inArgs[:], nil)
	}

	return
}

var applicationSetOptionContextParameterStringFunction *gi.Function
var applicationSetOptionContextParameterStringFunction_Once sync.Once

func applicationSetOptionContextParameterStringFunction_Set() error {
	var err error
	applicationSetOptionContextParameterStringFunction_Once.Do(func() {
		err = applicationObject_Set()
		if err != nil {
			return
		}
		applicationSetOptionContextParameterStringFunction, err = applicationObject.InvokerNew("set_option_context_parameter_string")
	})
	return err
}

// SetOptionContextParameterString is a representation of the C type g_application_set_option_context_parameter_string.
func (recv *Application) SetOptionContextParameterString(parameterString string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(parameterString)

	err := applicationSetOptionContextParameterStringFunction_Set()
	if err == nil {
		applicationSetOptionContextParameterStringFunction.Invoke(inArgs[:], nil)
	}

	return
}

var applicationSetOptionContextSummaryFunction *gi.Function
var applicationSetOptionContextSummaryFunction_Once sync.Once

func applicationSetOptionContextSummaryFunction_Set() error {
	var err error
	applicationSetOptionContextSummaryFunction_Once.Do(func() {
		err = applicationObject_Set()
		if err != nil {
			return
		}
		applicationSetOptionContextSummaryFunction, err = applicationObject.InvokerNew("set_option_context_summary")
	})
	return err
}

// SetOptionContextSummary is a representation of the C type g_application_set_option_context_summary.
func (recv *Application) SetOptionContextSummary(summary string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(summary)

	err := applicationSetOptionContextSummaryFunction_Set()
	if err == nil {
		applicationSetOptionContextSummaryFunction.Invoke(inArgs[:], nil)
	}

	return
}

var applicationSetResourceBasePathFunction *gi.Function
var applicationSetResourceBasePathFunction_Once sync.Once

func applicationSetResourceBasePathFunction_Set() error {
	var err error
	applicationSetResourceBasePathFunction_Once.Do(func() {
		err = applicationObject_Set()
		if err != nil {
			return
		}
		applicationSetResourceBasePathFunction, err = applicationObject.InvokerNew("set_resource_base_path")
	})
	return err
}

// SetResourceBasePath is a representation of the C type g_application_set_resource_base_path.
func (recv *Application) SetResourceBasePath(resourcePath string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(resourcePath)

	err := applicationSetResourceBasePathFunction_Set()
	if err == nil {
		applicationSetResourceBasePathFunction.Invoke(inArgs[:], nil)
	}

	return
}

var applicationUnbindBusyPropertyFunction *gi.Function
var applicationUnbindBusyPropertyFunction_Once sync.Once

func applicationUnbindBusyPropertyFunction_Set() error {
	var err error
	applicationUnbindBusyPropertyFunction_Once.Do(func() {
		err = applicationObject_Set()
		if err != nil {
			return
		}
		applicationUnbindBusyPropertyFunction, err = applicationObject.InvokerNew("unbind_busy_property")
	})
	return err
}

// UnbindBusyProperty is a representation of the C type g_application_unbind_busy_property.
func (recv *Application) UnbindBusyProperty(object_ *gobject.Object, property string) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(object_.Native())
	inArgs[2].SetString(property)

	err := applicationUnbindBusyPropertyFunction_Set()
	if err == nil {
		applicationUnbindBusyPropertyFunction.Invoke(inArgs[:], nil)
	}

	return
}

var applicationUnmarkBusyFunction *gi.Function
var applicationUnmarkBusyFunction_Once sync.Once

func applicationUnmarkBusyFunction_Set() error {
	var err error
	applicationUnmarkBusyFunction_Once.Do(func() {
		err = applicationObject_Set()
		if err != nil {
			return
		}
		applicationUnmarkBusyFunction, err = applicationObject.InvokerNew("unmark_busy")
	})
	return err
}

// UnmarkBusy is a representation of the C type g_application_unmark_busy.
func (recv *Application) UnmarkBusy() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := applicationUnmarkBusyFunction_Set()
	if err == nil {
		applicationUnmarkBusyFunction.Invoke(inArgs[:], nil)
	}

	return
}

var applicationWithdrawNotificationFunction *gi.Function
var applicationWithdrawNotificationFunction_Once sync.Once

func applicationWithdrawNotificationFunction_Set() error {
	var err error
	applicationWithdrawNotificationFunction_Once.Do(func() {
		err = applicationObject_Set()
		if err != nil {
			return
		}
		applicationWithdrawNotificationFunction, err = applicationObject.InvokerNew("withdraw_notification")
	})
	return err
}

// WithdrawNotification is a representation of the C type g_application_withdraw_notification.
func (recv *Application) WithdrawNotification(id string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(id)

	err := applicationWithdrawNotificationFunction_Set()
	if err == nil {
		applicationWithdrawNotificationFunction.Invoke(inArgs[:], nil)
	}

	return
}

var applicationCommandLineObject *gi.Object
var applicationCommandLineObject_Once sync.Once

func applicationCommandLineObject_Set() error {
	var err error
	applicationCommandLineObject_Once.Do(func() {
		applicationCommandLineObject, err = gi.ObjectNew("Gio", "ApplicationCommandLine")
	})
	return err
}

type ApplicationCommandLine struct {
	native unsafe.Pointer
}

func ApplicationCommandLineNewFromNative(native unsafe.Pointer) *ApplicationCommandLine {
	instance := &ApplicationCommandLine{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *ApplicationCommandLine) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToApplicationCommandLine down casts any arbitrary Object to ApplicationCommandLine.
Exercise care, as this is a potentially dangerous function
if the Object is not a ApplicationCommandLine.
*/
func CastToApplicationCommandLine(object *gobject.Object) *ApplicationCommandLine {
	return ApplicationCommandLineNewFromNative(object.Native())
}

// Equals compares this ApplicationCommandLine with another ApplicationCommandLine, and returns true if they represent the same GObject.
func (recv *ApplicationCommandLine) Equals(other *ApplicationCommandLine) bool {
	return other.Native() == recv.Native()
}

func (recv *ApplicationCommandLine) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'g_application_command_line_create_file_for_arg' : return type 'File' not supported

var applicationCommandLineGetArgumentsFunction *gi.Function
var applicationCommandLineGetArgumentsFunction_Once sync.Once

func applicationCommandLineGetArgumentsFunction_Set() error {
	var err error
	applicationCommandLineGetArgumentsFunction_Once.Do(func() {
		err = applicationCommandLineObject_Set()
		if err != nil {
			return
		}
		applicationCommandLineGetArgumentsFunction, err = applicationCommandLineObject.InvokerNew("get_arguments")
	})
	return err
}

// GetArguments is a representation of the C type g_application_command_line_get_arguments.
func (recv *ApplicationCommandLine) GetArguments() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var outArgs [1]gi.Argument

	err := applicationCommandLineGetArgumentsFunction_Set()
	if err == nil {
		applicationCommandLineGetArgumentsFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := outArgs[0].Int32()

	return out0
}

var applicationCommandLineGetCwdFunction *gi.Function
var applicationCommandLineGetCwdFunction_Once sync.Once

func applicationCommandLineGetCwdFunction_Set() error {
	var err error
	applicationCommandLineGetCwdFunction_Once.Do(func() {
		err = applicationCommandLineObject_Set()
		if err != nil {
			return
		}
		applicationCommandLineGetCwdFunction, err = applicationCommandLineObject.InvokerNew("get_cwd")
	})
	return err
}

// GetCwd is a representation of the C type g_application_command_line_get_cwd.
func (recv *ApplicationCommandLine) GetCwd() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := applicationCommandLineGetCwdFunction_Set()
	if err == nil {
		ret = applicationCommandLineGetCwdFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var applicationCommandLineGetEnvironFunction *gi.Function
var applicationCommandLineGetEnvironFunction_Once sync.Once

func applicationCommandLineGetEnvironFunction_Set() error {
	var err error
	applicationCommandLineGetEnvironFunction_Once.Do(func() {
		err = applicationCommandLineObject_Set()
		if err != nil {
			return
		}
		applicationCommandLineGetEnvironFunction, err = applicationCommandLineObject.InvokerNew("get_environ")
	})
	return err
}

// GetEnviron is a representation of the C type g_application_command_line_get_environ.
func (recv *ApplicationCommandLine) GetEnviron() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := applicationCommandLineGetEnvironFunction_Set()
	if err == nil {
		applicationCommandLineGetEnvironFunction.Invoke(inArgs[:], nil)
	}

	return
}

var applicationCommandLineGetExitStatusFunction *gi.Function
var applicationCommandLineGetExitStatusFunction_Once sync.Once

func applicationCommandLineGetExitStatusFunction_Set() error {
	var err error
	applicationCommandLineGetExitStatusFunction_Once.Do(func() {
		err = applicationCommandLineObject_Set()
		if err != nil {
			return
		}
		applicationCommandLineGetExitStatusFunction, err = applicationCommandLineObject.InvokerNew("get_exit_status")
	})
	return err
}

// GetExitStatus is a representation of the C type g_application_command_line_get_exit_status.
func (recv *ApplicationCommandLine) GetExitStatus() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := applicationCommandLineGetExitStatusFunction_Set()
	if err == nil {
		ret = applicationCommandLineGetExitStatusFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var applicationCommandLineGetIsRemoteFunction *gi.Function
var applicationCommandLineGetIsRemoteFunction_Once sync.Once

func applicationCommandLineGetIsRemoteFunction_Set() error {
	var err error
	applicationCommandLineGetIsRemoteFunction_Once.Do(func() {
		err = applicationCommandLineObject_Set()
		if err != nil {
			return
		}
		applicationCommandLineGetIsRemoteFunction, err = applicationCommandLineObject.InvokerNew("get_is_remote")
	})
	return err
}

// GetIsRemote is a representation of the C type g_application_command_line_get_is_remote.
func (recv *ApplicationCommandLine) GetIsRemote() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := applicationCommandLineGetIsRemoteFunction_Set()
	if err == nil {
		ret = applicationCommandLineGetIsRemoteFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'g_application_command_line_get_options_dict' : return type 'GLib.VariantDict' not supported

// UNSUPPORTED : C value 'g_application_command_line_get_platform_data' : return type 'GLib.Variant' not supported

var applicationCommandLineGetStdinFunction *gi.Function
var applicationCommandLineGetStdinFunction_Once sync.Once

func applicationCommandLineGetStdinFunction_Set() error {
	var err error
	applicationCommandLineGetStdinFunction_Once.Do(func() {
		err = applicationCommandLineObject_Set()
		if err != nil {
			return
		}
		applicationCommandLineGetStdinFunction, err = applicationCommandLineObject.InvokerNew("get_stdin")
	})
	return err
}

// GetStdin is a representation of the C type g_application_command_line_get_stdin.
func (recv *ApplicationCommandLine) GetStdin() *InputStream {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := applicationCommandLineGetStdinFunction_Set()
	if err == nil {
		ret = applicationCommandLineGetStdinFunction.Invoke(inArgs[:], nil)
	}

	retGo := InputStreamNewFromNative(ret.Pointer())

	return retGo
}

var applicationCommandLineGetenvFunction *gi.Function
var applicationCommandLineGetenvFunction_Once sync.Once

func applicationCommandLineGetenvFunction_Set() error {
	var err error
	applicationCommandLineGetenvFunction_Once.Do(func() {
		err = applicationCommandLineObject_Set()
		if err != nil {
			return
		}
		applicationCommandLineGetenvFunction, err = applicationCommandLineObject.InvokerNew("getenv")
	})
	return err
}

// Getenv is a representation of the C type g_application_command_line_getenv.
func (recv *ApplicationCommandLine) Getenv(name string) string {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(name)

	var ret gi.Argument

	err := applicationCommandLineGetenvFunction_Set()
	if err == nil {
		ret = applicationCommandLineGetenvFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'g_application_command_line_print' : parameter '...' of type 'nil' not supported

// UNSUPPORTED : C value 'g_application_command_line_printerr' : parameter '...' of type 'nil' not supported

var applicationCommandLineSetExitStatusFunction *gi.Function
var applicationCommandLineSetExitStatusFunction_Once sync.Once

func applicationCommandLineSetExitStatusFunction_Set() error {
	var err error
	applicationCommandLineSetExitStatusFunction_Once.Do(func() {
		err = applicationCommandLineObject_Set()
		if err != nil {
			return
		}
		applicationCommandLineSetExitStatusFunction, err = applicationCommandLineObject.InvokerNew("set_exit_status")
	})
	return err
}

// SetExitStatus is a representation of the C type g_application_command_line_set_exit_status.
func (recv *ApplicationCommandLine) SetExitStatus(exitStatus int32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(exitStatus)

	err := applicationCommandLineSetExitStatusFunction_Set()
	if err == nil {
		applicationCommandLineSetExitStatusFunction.Invoke(inArgs[:], nil)
	}

	return
}

var bufferedInputStreamObject *gi.Object
var bufferedInputStreamObject_Once sync.Once

func bufferedInputStreamObject_Set() error {
	var err error
	bufferedInputStreamObject_Once.Do(func() {
		bufferedInputStreamObject, err = gi.ObjectNew("Gio", "BufferedInputStream")
	})
	return err
}

type BufferedInputStream struct {
	native unsafe.Pointer
}

func BufferedInputStreamNewFromNative(native unsafe.Pointer) *BufferedInputStream {
	instance := &BufferedInputStream{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// FilterInputStream upcasts to *FilterInputStream
func (recv *BufferedInputStream) FilterInputStream() *FilterInputStream {
	return FilterInputStreamNewFromNative(recv.Native())
}

// InputStream upcasts to *InputStream
func (recv *BufferedInputStream) InputStream() *InputStream {
	return InputStreamNewFromNative(recv.Native())
}

// Object upcasts to *Object
func (recv *BufferedInputStream) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToBufferedInputStream down casts any arbitrary Object to BufferedInputStream.
Exercise care, as this is a potentially dangerous function
if the Object is not a BufferedInputStream.
*/
func CastToBufferedInputStream(object *gobject.Object) *BufferedInputStream {
	return BufferedInputStreamNewFromNative(object.Native())
}

// Equals compares this BufferedInputStream with another BufferedInputStream, and returns true if they represent the same GObject.
func (recv *BufferedInputStream) Equals(other *BufferedInputStream) bool {
	return other.Native() == recv.Native()
}

func (recv *BufferedInputStream) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *BufferedInputStream) FieldParentInstance() *FilterInputStream {
	argValue := gi.ObjectFieldGet(bufferedInputStreamObject, recv.Native(), "parent_instance")
	value := FilterInputStreamNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *BufferedInputStream) SetFieldParentInstance(value *FilterInputStream) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(bufferedInputStreamObject, recv.Native(), "parent_instance", argValue)
}

var bufferedInputStreamNewFunction *gi.Function
var bufferedInputStreamNewFunction_Once sync.Once

func bufferedInputStreamNewFunction_Set() error {
	var err error
	bufferedInputStreamNewFunction_Once.Do(func() {
		err = bufferedInputStreamObject_Set()
		if err != nil {
			return
		}
		bufferedInputStreamNewFunction, err = bufferedInputStreamObject.InvokerNew("new")
	})
	return err
}

// BufferedInputStreamNew is a representation of the C type g_buffered_input_stream_new.
func BufferedInputStreamNew(baseStream *InputStream) *BufferedInputStream {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(baseStream.Native())

	var ret gi.Argument

	err := bufferedInputStreamNewFunction_Set()
	if err == nil {
		ret = bufferedInputStreamNewFunction.Invoke(inArgs[:], nil)
	}

	retGo := BufferedInputStreamNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

var bufferedInputStreamNewSizedFunction *gi.Function
var bufferedInputStreamNewSizedFunction_Once sync.Once

func bufferedInputStreamNewSizedFunction_Set() error {
	var err error
	bufferedInputStreamNewSizedFunction_Once.Do(func() {
		err = bufferedInputStreamObject_Set()
		if err != nil {
			return
		}
		bufferedInputStreamNewSizedFunction, err = bufferedInputStreamObject.InvokerNew("new_sized")
	})
	return err
}

// BufferedInputStreamNewSized is a representation of the C type g_buffered_input_stream_new_sized.
func BufferedInputStreamNewSized(baseStream *InputStream, size uint64) *BufferedInputStream {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(baseStream.Native())
	inArgs[1].SetUint64(size)

	var ret gi.Argument

	err := bufferedInputStreamNewSizedFunction_Set()
	if err == nil {
		ret = bufferedInputStreamNewSizedFunction.Invoke(inArgs[:], nil)
	}

	retGo := BufferedInputStreamNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

var bufferedInputStreamFillFunction *gi.Function
var bufferedInputStreamFillFunction_Once sync.Once

func bufferedInputStreamFillFunction_Set() error {
	var err error
	bufferedInputStreamFillFunction_Once.Do(func() {
		err = bufferedInputStreamObject_Set()
		if err != nil {
			return
		}
		bufferedInputStreamFillFunction, err = bufferedInputStreamObject.InvokerNew("fill")
	})
	return err
}

// Fill is a representation of the C type g_buffered_input_stream_fill.
func (recv *BufferedInputStream) Fill(count int32, cancellable *Cancellable) int32 {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(count)
	inArgs[2].SetPointer(cancellable.Native())

	var ret gi.Argument

	err := bufferedInputStreamFillFunction_Set()
	if err == nil {
		ret = bufferedInputStreamFillFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

// UNSUPPORTED : C value 'g_buffered_input_stream_fill_async' : parameter 'callback' of type 'AsyncReadyCallback' not supported

// UNSUPPORTED : C value 'g_buffered_input_stream_fill_finish' : parameter 'result' of type 'AsyncResult' not supported

var bufferedInputStreamGetAvailableFunction *gi.Function
var bufferedInputStreamGetAvailableFunction_Once sync.Once

func bufferedInputStreamGetAvailableFunction_Set() error {
	var err error
	bufferedInputStreamGetAvailableFunction_Once.Do(func() {
		err = bufferedInputStreamObject_Set()
		if err != nil {
			return
		}
		bufferedInputStreamGetAvailableFunction, err = bufferedInputStreamObject.InvokerNew("get_available")
	})
	return err
}

// GetAvailable is a representation of the C type g_buffered_input_stream_get_available.
func (recv *BufferedInputStream) GetAvailable() uint64 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := bufferedInputStreamGetAvailableFunction_Set()
	if err == nil {
		ret = bufferedInputStreamGetAvailableFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint64()

	return retGo
}

var bufferedInputStreamGetBufferSizeFunction *gi.Function
var bufferedInputStreamGetBufferSizeFunction_Once sync.Once

func bufferedInputStreamGetBufferSizeFunction_Set() error {
	var err error
	bufferedInputStreamGetBufferSizeFunction_Once.Do(func() {
		err = bufferedInputStreamObject_Set()
		if err != nil {
			return
		}
		bufferedInputStreamGetBufferSizeFunction, err = bufferedInputStreamObject.InvokerNew("get_buffer_size")
	})
	return err
}

// GetBufferSize is a representation of the C type g_buffered_input_stream_get_buffer_size.
func (recv *BufferedInputStream) GetBufferSize() uint64 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := bufferedInputStreamGetBufferSizeFunction_Set()
	if err == nil {
		ret = bufferedInputStreamGetBufferSizeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint64()

	return retGo
}

// UNSUPPORTED : C value 'g_buffered_input_stream_peek' : parameter 'buffer' of type 'nil' not supported

var bufferedInputStreamPeekBufferFunction *gi.Function
var bufferedInputStreamPeekBufferFunction_Once sync.Once

func bufferedInputStreamPeekBufferFunction_Set() error {
	var err error
	bufferedInputStreamPeekBufferFunction_Once.Do(func() {
		err = bufferedInputStreamObject_Set()
		if err != nil {
			return
		}
		bufferedInputStreamPeekBufferFunction, err = bufferedInputStreamObject.InvokerNew("peek_buffer")
	})
	return err
}

// PeekBuffer is a representation of the C type g_buffered_input_stream_peek_buffer.
func (recv *BufferedInputStream) PeekBuffer() uint64 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var outArgs [1]gi.Argument

	err := bufferedInputStreamPeekBufferFunction_Set()
	if err == nil {
		bufferedInputStreamPeekBufferFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := outArgs[0].Uint64()

	return out0
}

var bufferedInputStreamReadByteFunction *gi.Function
var bufferedInputStreamReadByteFunction_Once sync.Once

func bufferedInputStreamReadByteFunction_Set() error {
	var err error
	bufferedInputStreamReadByteFunction_Once.Do(func() {
		err = bufferedInputStreamObject_Set()
		if err != nil {
			return
		}
		bufferedInputStreamReadByteFunction, err = bufferedInputStreamObject.InvokerNew("read_byte")
	})
	return err
}

// ReadByte is a representation of the C type g_buffered_input_stream_read_byte.
func (recv *BufferedInputStream) ReadByte(cancellable *Cancellable) int32 {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(cancellable.Native())

	var ret gi.Argument

	err := bufferedInputStreamReadByteFunction_Set()
	if err == nil {
		ret = bufferedInputStreamReadByteFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var bufferedInputStreamSetBufferSizeFunction *gi.Function
var bufferedInputStreamSetBufferSizeFunction_Once sync.Once

func bufferedInputStreamSetBufferSizeFunction_Set() error {
	var err error
	bufferedInputStreamSetBufferSizeFunction_Once.Do(func() {
		err = bufferedInputStreamObject_Set()
		if err != nil {
			return
		}
		bufferedInputStreamSetBufferSizeFunction, err = bufferedInputStreamObject.InvokerNew("set_buffer_size")
	})
	return err
}

// SetBufferSize is a representation of the C type g_buffered_input_stream_set_buffer_size.
func (recv *BufferedInputStream) SetBufferSize(size uint64) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetUint64(size)

	err := bufferedInputStreamSetBufferSizeFunction_Set()
	if err == nil {
		bufferedInputStreamSetBufferSizeFunction.Invoke(inArgs[:], nil)
	}

	return
}

var bufferedOutputStreamObject *gi.Object
var bufferedOutputStreamObject_Once sync.Once

func bufferedOutputStreamObject_Set() error {
	var err error
	bufferedOutputStreamObject_Once.Do(func() {
		bufferedOutputStreamObject, err = gi.ObjectNew("Gio", "BufferedOutputStream")
	})
	return err
}

type BufferedOutputStream struct {
	native unsafe.Pointer
}

func BufferedOutputStreamNewFromNative(native unsafe.Pointer) *BufferedOutputStream {
	instance := &BufferedOutputStream{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// FilterOutputStream upcasts to *FilterOutputStream
func (recv *BufferedOutputStream) FilterOutputStream() *FilterOutputStream {
	return FilterOutputStreamNewFromNative(recv.Native())
}

// OutputStream upcasts to *OutputStream
func (recv *BufferedOutputStream) OutputStream() *OutputStream {
	return OutputStreamNewFromNative(recv.Native())
}

// Object upcasts to *Object
func (recv *BufferedOutputStream) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToBufferedOutputStream down casts any arbitrary Object to BufferedOutputStream.
Exercise care, as this is a potentially dangerous function
if the Object is not a BufferedOutputStream.
*/
func CastToBufferedOutputStream(object *gobject.Object) *BufferedOutputStream {
	return BufferedOutputStreamNewFromNative(object.Native())
}

// Equals compares this BufferedOutputStream with another BufferedOutputStream, and returns true if they represent the same GObject.
func (recv *BufferedOutputStream) Equals(other *BufferedOutputStream) bool {
	return other.Native() == recv.Native()
}

func (recv *BufferedOutputStream) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *BufferedOutputStream) FieldParentInstance() *FilterOutputStream {
	argValue := gi.ObjectFieldGet(bufferedOutputStreamObject, recv.Native(), "parent_instance")
	value := FilterOutputStreamNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *BufferedOutputStream) SetFieldParentInstance(value *FilterOutputStream) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(bufferedOutputStreamObject, recv.Native(), "parent_instance", argValue)
}

// FieldPriv returns the C field 'priv'.
func (recv *BufferedOutputStream) FieldPriv() *BufferedOutputStreamPrivate {
	argValue := gi.ObjectFieldGet(bufferedOutputStreamObject, recv.Native(), "priv")
	value := BufferedOutputStreamPrivateNewFromNative(argValue.Pointer())
	return value
}

// SetFieldPriv sets the value of the C field 'priv'.
func (recv *BufferedOutputStream) SetFieldPriv(value *BufferedOutputStreamPrivate) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(bufferedOutputStreamObject, recv.Native(), "priv", argValue)
}

var bufferedOutputStreamNewFunction *gi.Function
var bufferedOutputStreamNewFunction_Once sync.Once

func bufferedOutputStreamNewFunction_Set() error {
	var err error
	bufferedOutputStreamNewFunction_Once.Do(func() {
		err = bufferedOutputStreamObject_Set()
		if err != nil {
			return
		}
		bufferedOutputStreamNewFunction, err = bufferedOutputStreamObject.InvokerNew("new")
	})
	return err
}

// BufferedOutputStreamNew is a representation of the C type g_buffered_output_stream_new.
func BufferedOutputStreamNew(baseStream *OutputStream) *BufferedOutputStream {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(baseStream.Native())

	var ret gi.Argument

	err := bufferedOutputStreamNewFunction_Set()
	if err == nil {
		ret = bufferedOutputStreamNewFunction.Invoke(inArgs[:], nil)
	}

	retGo := BufferedOutputStreamNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

var bufferedOutputStreamNewSizedFunction *gi.Function
var bufferedOutputStreamNewSizedFunction_Once sync.Once

func bufferedOutputStreamNewSizedFunction_Set() error {
	var err error
	bufferedOutputStreamNewSizedFunction_Once.Do(func() {
		err = bufferedOutputStreamObject_Set()
		if err != nil {
			return
		}
		bufferedOutputStreamNewSizedFunction, err = bufferedOutputStreamObject.InvokerNew("new_sized")
	})
	return err
}

// BufferedOutputStreamNewSized is a representation of the C type g_buffered_output_stream_new_sized.
func BufferedOutputStreamNewSized(baseStream *OutputStream, size uint64) *BufferedOutputStream {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(baseStream.Native())
	inArgs[1].SetUint64(size)

	var ret gi.Argument

	err := bufferedOutputStreamNewSizedFunction_Set()
	if err == nil {
		ret = bufferedOutputStreamNewSizedFunction.Invoke(inArgs[:], nil)
	}

	retGo := BufferedOutputStreamNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

var bufferedOutputStreamGetAutoGrowFunction *gi.Function
var bufferedOutputStreamGetAutoGrowFunction_Once sync.Once

func bufferedOutputStreamGetAutoGrowFunction_Set() error {
	var err error
	bufferedOutputStreamGetAutoGrowFunction_Once.Do(func() {
		err = bufferedOutputStreamObject_Set()
		if err != nil {
			return
		}
		bufferedOutputStreamGetAutoGrowFunction, err = bufferedOutputStreamObject.InvokerNew("get_auto_grow")
	})
	return err
}

// GetAutoGrow is a representation of the C type g_buffered_output_stream_get_auto_grow.
func (recv *BufferedOutputStream) GetAutoGrow() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := bufferedOutputStreamGetAutoGrowFunction_Set()
	if err == nil {
		ret = bufferedOutputStreamGetAutoGrowFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var bufferedOutputStreamGetBufferSizeFunction *gi.Function
var bufferedOutputStreamGetBufferSizeFunction_Once sync.Once

func bufferedOutputStreamGetBufferSizeFunction_Set() error {
	var err error
	bufferedOutputStreamGetBufferSizeFunction_Once.Do(func() {
		err = bufferedOutputStreamObject_Set()
		if err != nil {
			return
		}
		bufferedOutputStreamGetBufferSizeFunction, err = bufferedOutputStreamObject.InvokerNew("get_buffer_size")
	})
	return err
}

// GetBufferSize is a representation of the C type g_buffered_output_stream_get_buffer_size.
func (recv *BufferedOutputStream) GetBufferSize() uint64 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := bufferedOutputStreamGetBufferSizeFunction_Set()
	if err == nil {
		ret = bufferedOutputStreamGetBufferSizeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint64()

	return retGo
}

var bufferedOutputStreamSetAutoGrowFunction *gi.Function
var bufferedOutputStreamSetAutoGrowFunction_Once sync.Once

func bufferedOutputStreamSetAutoGrowFunction_Set() error {
	var err error
	bufferedOutputStreamSetAutoGrowFunction_Once.Do(func() {
		err = bufferedOutputStreamObject_Set()
		if err != nil {
			return
		}
		bufferedOutputStreamSetAutoGrowFunction, err = bufferedOutputStreamObject.InvokerNew("set_auto_grow")
	})
	return err
}

// SetAutoGrow is a representation of the C type g_buffered_output_stream_set_auto_grow.
func (recv *BufferedOutputStream) SetAutoGrow(autoGrow bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetBoolean(autoGrow)

	err := bufferedOutputStreamSetAutoGrowFunction_Set()
	if err == nil {
		bufferedOutputStreamSetAutoGrowFunction.Invoke(inArgs[:], nil)
	}

	return
}

var bufferedOutputStreamSetBufferSizeFunction *gi.Function
var bufferedOutputStreamSetBufferSizeFunction_Once sync.Once

func bufferedOutputStreamSetBufferSizeFunction_Set() error {
	var err error
	bufferedOutputStreamSetBufferSizeFunction_Once.Do(func() {
		err = bufferedOutputStreamObject_Set()
		if err != nil {
			return
		}
		bufferedOutputStreamSetBufferSizeFunction, err = bufferedOutputStreamObject.InvokerNew("set_buffer_size")
	})
	return err
}

// SetBufferSize is a representation of the C type g_buffered_output_stream_set_buffer_size.
func (recv *BufferedOutputStream) SetBufferSize(size uint64) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetUint64(size)

	err := bufferedOutputStreamSetBufferSizeFunction_Set()
	if err == nil {
		bufferedOutputStreamSetBufferSizeFunction.Invoke(inArgs[:], nil)
	}

	return
}

var bytesIconObject *gi.Object
var bytesIconObject_Once sync.Once

func bytesIconObject_Set() error {
	var err error
	bytesIconObject_Once.Do(func() {
		bytesIconObject, err = gi.ObjectNew("Gio", "BytesIcon")
	})
	return err
}

type BytesIcon struct {
	native unsafe.Pointer
}

func BytesIconNewFromNative(native unsafe.Pointer) *BytesIcon {
	instance := &BytesIcon{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *BytesIcon) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToBytesIcon down casts any arbitrary Object to BytesIcon.
Exercise care, as this is a potentially dangerous function
if the Object is not a BytesIcon.
*/
func CastToBytesIcon(object *gobject.Object) *BytesIcon {
	return BytesIconNewFromNative(object.Native())
}

// Equals compares this BytesIcon with another BytesIcon, and returns true if they represent the same GObject.
func (recv *BytesIcon) Equals(other *BytesIcon) bool {
	return other.Native() == recv.Native()
}

func (recv *BytesIcon) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'g_bytes_icon_new' : parameter 'bytes' of type 'GLib.Bytes' not supported

// UNSUPPORTED : C value 'g_bytes_icon_get_bytes' : return type 'GLib.Bytes' not supported

var cancellableObject *gi.Object
var cancellableObject_Once sync.Once

func cancellableObject_Set() error {
	var err error
	cancellableObject_Once.Do(func() {
		cancellableObject, err = gi.ObjectNew("Gio", "Cancellable")
	})
	return err
}

type Cancellable struct {
	native unsafe.Pointer
}

func CancellableNewFromNative(native unsafe.Pointer) *Cancellable {
	instance := &Cancellable{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *Cancellable) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToCancellable down casts any arbitrary Object to Cancellable.
Exercise care, as this is a potentially dangerous function
if the Object is not a Cancellable.
*/
func CastToCancellable(object *gobject.Object) *Cancellable {
	return CancellableNewFromNative(object.Native())
}

// Equals compares this Cancellable with another Cancellable, and returns true if they represent the same GObject.
func (recv *Cancellable) Equals(other *Cancellable) bool {
	return other.Native() == recv.Native()
}

func (recv *Cancellable) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *Cancellable) FieldParentInstance() *gobject.Object {
	argValue := gi.ObjectFieldGet(cancellableObject, recv.Native(), "parent_instance")
	value := gobject.ObjectNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *Cancellable) SetFieldParentInstance(value *gobject.Object) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(cancellableObject, recv.Native(), "parent_instance", argValue)
}

var cancellableNewFunction *gi.Function
var cancellableNewFunction_Once sync.Once

func cancellableNewFunction_Set() error {
	var err error
	cancellableNewFunction_Once.Do(func() {
		err = cancellableObject_Set()
		if err != nil {
			return
		}
		cancellableNewFunction, err = cancellableObject.InvokerNew("new")
	})
	return err
}

// CancellableNew is a representation of the C type g_cancellable_new.
func CancellableNew() *Cancellable {

	var ret gi.Argument

	err := cancellableNewFunction_Set()
	if err == nil {
		ret = cancellableNewFunction.Invoke(nil, nil)
	}

	retGo := CancellableNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

var cancellableCancelFunction *gi.Function
var cancellableCancelFunction_Once sync.Once

func cancellableCancelFunction_Set() error {
	var err error
	cancellableCancelFunction_Once.Do(func() {
		err = cancellableObject_Set()
		if err != nil {
			return
		}
		cancellableCancelFunction, err = cancellableObject.InvokerNew("cancel")
	})
	return err
}

// Cancel is a representation of the C type g_cancellable_cancel.
func (recv *Cancellable) Cancel() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := cancellableCancelFunction_Set()
	if err == nil {
		cancellableCancelFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'g_cancellable_connect' : parameter 'callback' of type 'GObject.Callback' not supported

var cancellableDisconnectFunction *gi.Function
var cancellableDisconnectFunction_Once sync.Once

func cancellableDisconnectFunction_Set() error {
	var err error
	cancellableDisconnectFunction_Once.Do(func() {
		err = cancellableObject_Set()
		if err != nil {
			return
		}
		cancellableDisconnectFunction, err = cancellableObject.InvokerNew("disconnect")
	})
	return err
}

// Disconnect is a representation of the C type g_cancellable_disconnect.
func (recv *Cancellable) Disconnect(handlerId uint64) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetUint64(handlerId)

	err := cancellableDisconnectFunction_Set()
	if err == nil {
		cancellableDisconnectFunction.Invoke(inArgs[:], nil)
	}

	return
}

var cancellableGetFdFunction *gi.Function
var cancellableGetFdFunction_Once sync.Once

func cancellableGetFdFunction_Set() error {
	var err error
	cancellableGetFdFunction_Once.Do(func() {
		err = cancellableObject_Set()
		if err != nil {
			return
		}
		cancellableGetFdFunction, err = cancellableObject.InvokerNew("get_fd")
	})
	return err
}

// GetFd is a representation of the C type g_cancellable_get_fd.
func (recv *Cancellable) GetFd() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := cancellableGetFdFunction_Set()
	if err == nil {
		ret = cancellableGetFdFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var cancellableIsCancelledFunction *gi.Function
var cancellableIsCancelledFunction_Once sync.Once

func cancellableIsCancelledFunction_Set() error {
	var err error
	cancellableIsCancelledFunction_Once.Do(func() {
		err = cancellableObject_Set()
		if err != nil {
			return
		}
		cancellableIsCancelledFunction, err = cancellableObject.InvokerNew("is_cancelled")
	})
	return err
}

// IsCancelled is a representation of the C type g_cancellable_is_cancelled.
func (recv *Cancellable) IsCancelled() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := cancellableIsCancelledFunction_Set()
	if err == nil {
		ret = cancellableIsCancelledFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'g_cancellable_make_pollfd' : parameter 'pollfd' of type 'GLib.PollFD' not supported

var cancellablePopCurrentFunction *gi.Function
var cancellablePopCurrentFunction_Once sync.Once

func cancellablePopCurrentFunction_Set() error {
	var err error
	cancellablePopCurrentFunction_Once.Do(func() {
		err = cancellableObject_Set()
		if err != nil {
			return
		}
		cancellablePopCurrentFunction, err = cancellableObject.InvokerNew("pop_current")
	})
	return err
}

// PopCurrent is a representation of the C type g_cancellable_pop_current.
func (recv *Cancellable) PopCurrent() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := cancellablePopCurrentFunction_Set()
	if err == nil {
		cancellablePopCurrentFunction.Invoke(inArgs[:], nil)
	}

	return
}

var cancellablePushCurrentFunction *gi.Function
var cancellablePushCurrentFunction_Once sync.Once

func cancellablePushCurrentFunction_Set() error {
	var err error
	cancellablePushCurrentFunction_Once.Do(func() {
		err = cancellableObject_Set()
		if err != nil {
			return
		}
		cancellablePushCurrentFunction, err = cancellableObject.InvokerNew("push_current")
	})
	return err
}

// PushCurrent is a representation of the C type g_cancellable_push_current.
func (recv *Cancellable) PushCurrent() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := cancellablePushCurrentFunction_Set()
	if err == nil {
		cancellablePushCurrentFunction.Invoke(inArgs[:], nil)
	}

	return
}

var cancellableReleaseFdFunction *gi.Function
var cancellableReleaseFdFunction_Once sync.Once

func cancellableReleaseFdFunction_Set() error {
	var err error
	cancellableReleaseFdFunction_Once.Do(func() {
		err = cancellableObject_Set()
		if err != nil {
			return
		}
		cancellableReleaseFdFunction, err = cancellableObject.InvokerNew("release_fd")
	})
	return err
}

// ReleaseFd is a representation of the C type g_cancellable_release_fd.
func (recv *Cancellable) ReleaseFd() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := cancellableReleaseFdFunction_Set()
	if err == nil {
		cancellableReleaseFdFunction.Invoke(inArgs[:], nil)
	}

	return
}

var cancellableResetFunction *gi.Function
var cancellableResetFunction_Once sync.Once

func cancellableResetFunction_Set() error {
	var err error
	cancellableResetFunction_Once.Do(func() {
		err = cancellableObject_Set()
		if err != nil {
			return
		}
		cancellableResetFunction, err = cancellableObject.InvokerNew("reset")
	})
	return err
}

// Reset is a representation of the C type g_cancellable_reset.
func (recv *Cancellable) Reset() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := cancellableResetFunction_Set()
	if err == nil {
		cancellableResetFunction.Invoke(inArgs[:], nil)
	}

	return
}

var cancellableSetErrorIfCancelledFunction *gi.Function
var cancellableSetErrorIfCancelledFunction_Once sync.Once

func cancellableSetErrorIfCancelledFunction_Set() error {
	var err error
	cancellableSetErrorIfCancelledFunction_Once.Do(func() {
		err = cancellableObject_Set()
		if err != nil {
			return
		}
		cancellableSetErrorIfCancelledFunction, err = cancellableObject.InvokerNew("set_error_if_cancelled")
	})
	return err
}

// SetErrorIfCancelled is a representation of the C type g_cancellable_set_error_if_cancelled.
func (recv *Cancellable) SetErrorIfCancelled() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := cancellableSetErrorIfCancelledFunction_Set()
	if err == nil {
		ret = cancellableSetErrorIfCancelledFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'g_cancellable_source_new' : return type 'GLib.Source' not supported

var charsetConverterObject *gi.Object
var charsetConverterObject_Once sync.Once

func charsetConverterObject_Set() error {
	var err error
	charsetConverterObject_Once.Do(func() {
		charsetConverterObject, err = gi.ObjectNew("Gio", "CharsetConverter")
	})
	return err
}

type CharsetConverter struct {
	native unsafe.Pointer
}

func CharsetConverterNewFromNative(native unsafe.Pointer) *CharsetConverter {
	instance := &CharsetConverter{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *CharsetConverter) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToCharsetConverter down casts any arbitrary Object to CharsetConverter.
Exercise care, as this is a potentially dangerous function
if the Object is not a CharsetConverter.
*/
func CastToCharsetConverter(object *gobject.Object) *CharsetConverter {
	return CharsetConverterNewFromNative(object.Native())
}

// Equals compares this CharsetConverter with another CharsetConverter, and returns true if they represent the same GObject.
func (recv *CharsetConverter) Equals(other *CharsetConverter) bool {
	return other.Native() == recv.Native()
}

func (recv *CharsetConverter) Native() unsafe.Pointer {
	return recv.native
}

var charsetConverterNewFunction *gi.Function
var charsetConverterNewFunction_Once sync.Once

func charsetConverterNewFunction_Set() error {
	var err error
	charsetConverterNewFunction_Once.Do(func() {
		err = charsetConverterObject_Set()
		if err != nil {
			return
		}
		charsetConverterNewFunction, err = charsetConverterObject.InvokerNew("new")
	})
	return err
}

// CharsetConverterNew is a representation of the C type g_charset_converter_new.
func CharsetConverterNew(toCharset string, fromCharset string) *CharsetConverter {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(toCharset)
	inArgs[1].SetString(fromCharset)

	var ret gi.Argument

	err := charsetConverterNewFunction_Set()
	if err == nil {
		ret = charsetConverterNewFunction.Invoke(inArgs[:], nil)
	}

	retGo := CharsetConverterNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

var charsetConverterGetNumFallbacksFunction *gi.Function
var charsetConverterGetNumFallbacksFunction_Once sync.Once

func charsetConverterGetNumFallbacksFunction_Set() error {
	var err error
	charsetConverterGetNumFallbacksFunction_Once.Do(func() {
		err = charsetConverterObject_Set()
		if err != nil {
			return
		}
		charsetConverterGetNumFallbacksFunction, err = charsetConverterObject.InvokerNew("get_num_fallbacks")
	})
	return err
}

// GetNumFallbacks is a representation of the C type g_charset_converter_get_num_fallbacks.
func (recv *CharsetConverter) GetNumFallbacks() uint32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := charsetConverterGetNumFallbacksFunction_Set()
	if err == nil {
		ret = charsetConverterGetNumFallbacksFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo
}

var charsetConverterGetUseFallbackFunction *gi.Function
var charsetConverterGetUseFallbackFunction_Once sync.Once

func charsetConverterGetUseFallbackFunction_Set() error {
	var err error
	charsetConverterGetUseFallbackFunction_Once.Do(func() {
		err = charsetConverterObject_Set()
		if err != nil {
			return
		}
		charsetConverterGetUseFallbackFunction, err = charsetConverterObject.InvokerNew("get_use_fallback")
	})
	return err
}

// GetUseFallback is a representation of the C type g_charset_converter_get_use_fallback.
func (recv *CharsetConverter) GetUseFallback() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := charsetConverterGetUseFallbackFunction_Set()
	if err == nil {
		ret = charsetConverterGetUseFallbackFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var charsetConverterSetUseFallbackFunction *gi.Function
var charsetConverterSetUseFallbackFunction_Once sync.Once

func charsetConverterSetUseFallbackFunction_Set() error {
	var err error
	charsetConverterSetUseFallbackFunction_Once.Do(func() {
		err = charsetConverterObject_Set()
		if err != nil {
			return
		}
		charsetConverterSetUseFallbackFunction, err = charsetConverterObject.InvokerNew("set_use_fallback")
	})
	return err
}

// SetUseFallback is a representation of the C type g_charset_converter_set_use_fallback.
func (recv *CharsetConverter) SetUseFallback(useFallback bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetBoolean(useFallback)

	err := charsetConverterSetUseFallbackFunction_Set()
	if err == nil {
		charsetConverterSetUseFallbackFunction.Invoke(inArgs[:], nil)
	}

	return
}

var converterInputStreamObject *gi.Object
var converterInputStreamObject_Once sync.Once

func converterInputStreamObject_Set() error {
	var err error
	converterInputStreamObject_Once.Do(func() {
		converterInputStreamObject, err = gi.ObjectNew("Gio", "ConverterInputStream")
	})
	return err
}

type ConverterInputStream struct {
	native unsafe.Pointer
}

func ConverterInputStreamNewFromNative(native unsafe.Pointer) *ConverterInputStream {
	instance := &ConverterInputStream{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// FilterInputStream upcasts to *FilterInputStream
func (recv *ConverterInputStream) FilterInputStream() *FilterInputStream {
	return FilterInputStreamNewFromNative(recv.Native())
}

// InputStream upcasts to *InputStream
func (recv *ConverterInputStream) InputStream() *InputStream {
	return InputStreamNewFromNative(recv.Native())
}

// Object upcasts to *Object
func (recv *ConverterInputStream) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToConverterInputStream down casts any arbitrary Object to ConverterInputStream.
Exercise care, as this is a potentially dangerous function
if the Object is not a ConverterInputStream.
*/
func CastToConverterInputStream(object *gobject.Object) *ConverterInputStream {
	return ConverterInputStreamNewFromNative(object.Native())
}

// Equals compares this ConverterInputStream with another ConverterInputStream, and returns true if they represent the same GObject.
func (recv *ConverterInputStream) Equals(other *ConverterInputStream) bool {
	return other.Native() == recv.Native()
}

func (recv *ConverterInputStream) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *ConverterInputStream) FieldParentInstance() *FilterInputStream {
	argValue := gi.ObjectFieldGet(converterInputStreamObject, recv.Native(), "parent_instance")
	value := FilterInputStreamNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *ConverterInputStream) SetFieldParentInstance(value *FilterInputStream) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(converterInputStreamObject, recv.Native(), "parent_instance", argValue)
}

// UNSUPPORTED : C value 'g_converter_input_stream_new' : parameter 'converter' of type 'Converter' not supported

// UNSUPPORTED : C value 'g_converter_input_stream_get_converter' : return type 'Converter' not supported

var converterOutputStreamObject *gi.Object
var converterOutputStreamObject_Once sync.Once

func converterOutputStreamObject_Set() error {
	var err error
	converterOutputStreamObject_Once.Do(func() {
		converterOutputStreamObject, err = gi.ObjectNew("Gio", "ConverterOutputStream")
	})
	return err
}

type ConverterOutputStream struct {
	native unsafe.Pointer
}

func ConverterOutputStreamNewFromNative(native unsafe.Pointer) *ConverterOutputStream {
	instance := &ConverterOutputStream{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// FilterOutputStream upcasts to *FilterOutputStream
func (recv *ConverterOutputStream) FilterOutputStream() *FilterOutputStream {
	return FilterOutputStreamNewFromNative(recv.Native())
}

// OutputStream upcasts to *OutputStream
func (recv *ConverterOutputStream) OutputStream() *OutputStream {
	return OutputStreamNewFromNative(recv.Native())
}

// Object upcasts to *Object
func (recv *ConverterOutputStream) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToConverterOutputStream down casts any arbitrary Object to ConverterOutputStream.
Exercise care, as this is a potentially dangerous function
if the Object is not a ConverterOutputStream.
*/
func CastToConverterOutputStream(object *gobject.Object) *ConverterOutputStream {
	return ConverterOutputStreamNewFromNative(object.Native())
}

// Equals compares this ConverterOutputStream with another ConverterOutputStream, and returns true if they represent the same GObject.
func (recv *ConverterOutputStream) Equals(other *ConverterOutputStream) bool {
	return other.Native() == recv.Native()
}

func (recv *ConverterOutputStream) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *ConverterOutputStream) FieldParentInstance() *FilterOutputStream {
	argValue := gi.ObjectFieldGet(converterOutputStreamObject, recv.Native(), "parent_instance")
	value := FilterOutputStreamNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *ConverterOutputStream) SetFieldParentInstance(value *FilterOutputStream) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(converterOutputStreamObject, recv.Native(), "parent_instance", argValue)
}

// UNSUPPORTED : C value 'g_converter_output_stream_new' : parameter 'converter' of type 'Converter' not supported

// UNSUPPORTED : C value 'g_converter_output_stream_get_converter' : return type 'Converter' not supported

var credentialsObject *gi.Object
var credentialsObject_Once sync.Once

func credentialsObject_Set() error {
	var err error
	credentialsObject_Once.Do(func() {
		credentialsObject, err = gi.ObjectNew("Gio", "Credentials")
	})
	return err
}

type Credentials struct {
	native unsafe.Pointer
}

func CredentialsNewFromNative(native unsafe.Pointer) *Credentials {
	instance := &Credentials{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *Credentials) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToCredentials down casts any arbitrary Object to Credentials.
Exercise care, as this is a potentially dangerous function
if the Object is not a Credentials.
*/
func CastToCredentials(object *gobject.Object) *Credentials {
	return CredentialsNewFromNative(object.Native())
}

// Equals compares this Credentials with another Credentials, and returns true if they represent the same GObject.
func (recv *Credentials) Equals(other *Credentials) bool {
	return other.Native() == recv.Native()
}

func (recv *Credentials) Native() unsafe.Pointer {
	return recv.native
}

var credentialsNewFunction *gi.Function
var credentialsNewFunction_Once sync.Once

func credentialsNewFunction_Set() error {
	var err error
	credentialsNewFunction_Once.Do(func() {
		err = credentialsObject_Set()
		if err != nil {
			return
		}
		credentialsNewFunction, err = credentialsObject.InvokerNew("new")
	})
	return err
}

// CredentialsNew is a representation of the C type g_credentials_new.
func CredentialsNew() *Credentials {

	var ret gi.Argument

	err := credentialsNewFunction_Set()
	if err == nil {
		ret = credentialsNewFunction.Invoke(nil, nil)
	}

	retGo := CredentialsNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

// UNSUPPORTED : C value 'g_credentials_get_native' : return type 'gpointer' not supported

var credentialsGetUnixPidFunction *gi.Function
var credentialsGetUnixPidFunction_Once sync.Once

func credentialsGetUnixPidFunction_Set() error {
	var err error
	credentialsGetUnixPidFunction_Once.Do(func() {
		err = credentialsObject_Set()
		if err != nil {
			return
		}
		credentialsGetUnixPidFunction, err = credentialsObject.InvokerNew("get_unix_pid")
	})
	return err
}

// GetUnixPid is a representation of the C type g_credentials_get_unix_pid.
func (recv *Credentials) GetUnixPid() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := credentialsGetUnixPidFunction_Set()
	if err == nil {
		ret = credentialsGetUnixPidFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var credentialsGetUnixUserFunction *gi.Function
var credentialsGetUnixUserFunction_Once sync.Once

func credentialsGetUnixUserFunction_Set() error {
	var err error
	credentialsGetUnixUserFunction_Once.Do(func() {
		err = credentialsObject_Set()
		if err != nil {
			return
		}
		credentialsGetUnixUserFunction, err = credentialsObject.InvokerNew("get_unix_user")
	})
	return err
}

// GetUnixUser is a representation of the C type g_credentials_get_unix_user.
func (recv *Credentials) GetUnixUser() uint32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := credentialsGetUnixUserFunction_Set()
	if err == nil {
		ret = credentialsGetUnixUserFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo
}

var credentialsIsSameUserFunction *gi.Function
var credentialsIsSameUserFunction_Once sync.Once

func credentialsIsSameUserFunction_Set() error {
	var err error
	credentialsIsSameUserFunction_Once.Do(func() {
		err = credentialsObject_Set()
		if err != nil {
			return
		}
		credentialsIsSameUserFunction, err = credentialsObject.InvokerNew("is_same_user")
	})
	return err
}

// IsSameUser is a representation of the C type g_credentials_is_same_user.
func (recv *Credentials) IsSameUser(otherCredentials *Credentials) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(otherCredentials.Native())

	var ret gi.Argument

	err := credentialsIsSameUserFunction_Set()
	if err == nil {
		ret = credentialsIsSameUserFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'g_credentials_set_native' : parameter 'native' of type 'gpointer' not supported

var credentialsSetUnixUserFunction *gi.Function
var credentialsSetUnixUserFunction_Once sync.Once

func credentialsSetUnixUserFunction_Set() error {
	var err error
	credentialsSetUnixUserFunction_Once.Do(func() {
		err = credentialsObject_Set()
		if err != nil {
			return
		}
		credentialsSetUnixUserFunction, err = credentialsObject.InvokerNew("set_unix_user")
	})
	return err
}

// SetUnixUser is a representation of the C type g_credentials_set_unix_user.
func (recv *Credentials) SetUnixUser(uid uint32) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetUint32(uid)

	var ret gi.Argument

	err := credentialsSetUnixUserFunction_Set()
	if err == nil {
		ret = credentialsSetUnixUserFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var credentialsToStringFunction *gi.Function
var credentialsToStringFunction_Once sync.Once

func credentialsToStringFunction_Set() error {
	var err error
	credentialsToStringFunction_Once.Do(func() {
		err = credentialsObject_Set()
		if err != nil {
			return
		}
		credentialsToStringFunction, err = credentialsObject.InvokerNew("to_string")
	})
	return err
}

// ToString is a representation of the C type g_credentials_to_string.
func (recv *Credentials) ToString() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := credentialsToStringFunction_Set()
	if err == nil {
		ret = credentialsToStringFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

var dBusActionGroupObject *gi.Object
var dBusActionGroupObject_Once sync.Once

func dBusActionGroupObject_Set() error {
	var err error
	dBusActionGroupObject_Once.Do(func() {
		dBusActionGroupObject, err = gi.ObjectNew("Gio", "DBusActionGroup")
	})
	return err
}

type DBusActionGroup struct {
	native unsafe.Pointer
}

func DBusActionGroupNewFromNative(native unsafe.Pointer) *DBusActionGroup {
	instance := &DBusActionGroup{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *DBusActionGroup) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToDBusActionGroup down casts any arbitrary Object to DBusActionGroup.
Exercise care, as this is a potentially dangerous function
if the Object is not a DBusActionGroup.
*/
func CastToDBusActionGroup(object *gobject.Object) *DBusActionGroup {
	return DBusActionGroupNewFromNative(object.Native())
}

// Equals compares this DBusActionGroup with another DBusActionGroup, and returns true if they represent the same GObject.
func (recv *DBusActionGroup) Equals(other *DBusActionGroup) bool {
	return other.Native() == recv.Native()
}

func (recv *DBusActionGroup) Native() unsafe.Pointer {
	return recv.native
}

var dBusAuthObserverObject *gi.Object
var dBusAuthObserverObject_Once sync.Once

func dBusAuthObserverObject_Set() error {
	var err error
	dBusAuthObserverObject_Once.Do(func() {
		dBusAuthObserverObject, err = gi.ObjectNew("Gio", "DBusAuthObserver")
	})
	return err
}

type DBusAuthObserver struct {
	native unsafe.Pointer
}

func DBusAuthObserverNewFromNative(native unsafe.Pointer) *DBusAuthObserver {
	instance := &DBusAuthObserver{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *DBusAuthObserver) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToDBusAuthObserver down casts any arbitrary Object to DBusAuthObserver.
Exercise care, as this is a potentially dangerous function
if the Object is not a DBusAuthObserver.
*/
func CastToDBusAuthObserver(object *gobject.Object) *DBusAuthObserver {
	return DBusAuthObserverNewFromNative(object.Native())
}

// Equals compares this DBusAuthObserver with another DBusAuthObserver, and returns true if they represent the same GObject.
func (recv *DBusAuthObserver) Equals(other *DBusAuthObserver) bool {
	return other.Native() == recv.Native()
}

func (recv *DBusAuthObserver) Native() unsafe.Pointer {
	return recv.native
}

var dBusAuthObserverNewFunction *gi.Function
var dBusAuthObserverNewFunction_Once sync.Once

func dBusAuthObserverNewFunction_Set() error {
	var err error
	dBusAuthObserverNewFunction_Once.Do(func() {
		err = dBusAuthObserverObject_Set()
		if err != nil {
			return
		}
		dBusAuthObserverNewFunction, err = dBusAuthObserverObject.InvokerNew("new")
	})
	return err
}

// DBusAuthObserverNew is a representation of the C type g_dbus_auth_observer_new.
func DBusAuthObserverNew() *DBusAuthObserver {

	var ret gi.Argument

	err := dBusAuthObserverNewFunction_Set()
	if err == nil {
		ret = dBusAuthObserverNewFunction.Invoke(nil, nil)
	}

	retGo := DBusAuthObserverNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

var dBusAuthObserverAllowMechanismFunction *gi.Function
var dBusAuthObserverAllowMechanismFunction_Once sync.Once

func dBusAuthObserverAllowMechanismFunction_Set() error {
	var err error
	dBusAuthObserverAllowMechanismFunction_Once.Do(func() {
		err = dBusAuthObserverObject_Set()
		if err != nil {
			return
		}
		dBusAuthObserverAllowMechanismFunction, err = dBusAuthObserverObject.InvokerNew("allow_mechanism")
	})
	return err
}

// AllowMechanism is a representation of the C type g_dbus_auth_observer_allow_mechanism.
func (recv *DBusAuthObserver) AllowMechanism(mechanism string) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(mechanism)

	var ret gi.Argument

	err := dBusAuthObserverAllowMechanismFunction_Set()
	if err == nil {
		ret = dBusAuthObserverAllowMechanismFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var dBusAuthObserverAuthorizeAuthenticatedPeerFunction *gi.Function
var dBusAuthObserverAuthorizeAuthenticatedPeerFunction_Once sync.Once

func dBusAuthObserverAuthorizeAuthenticatedPeerFunction_Set() error {
	var err error
	dBusAuthObserverAuthorizeAuthenticatedPeerFunction_Once.Do(func() {
		err = dBusAuthObserverObject_Set()
		if err != nil {
			return
		}
		dBusAuthObserverAuthorizeAuthenticatedPeerFunction, err = dBusAuthObserverObject.InvokerNew("authorize_authenticated_peer")
	})
	return err
}

// AuthorizeAuthenticatedPeer is a representation of the C type g_dbus_auth_observer_authorize_authenticated_peer.
func (recv *DBusAuthObserver) AuthorizeAuthenticatedPeer(stream *IOStream, credentials *Credentials) bool {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(stream.Native())
	inArgs[2].SetPointer(credentials.Native())

	var ret gi.Argument

	err := dBusAuthObserverAuthorizeAuthenticatedPeerFunction_Set()
	if err == nil {
		ret = dBusAuthObserverAuthorizeAuthenticatedPeerFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var dBusConnectionObject *gi.Object
var dBusConnectionObject_Once sync.Once

func dBusConnectionObject_Set() error {
	var err error
	dBusConnectionObject_Once.Do(func() {
		dBusConnectionObject, err = gi.ObjectNew("Gio", "DBusConnection")
	})
	return err
}

type DBusConnection struct {
	native unsafe.Pointer
}

func DBusConnectionNewFromNative(native unsafe.Pointer) *DBusConnection {
	instance := &DBusConnection{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *DBusConnection) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToDBusConnection down casts any arbitrary Object to DBusConnection.
Exercise care, as this is a potentially dangerous function
if the Object is not a DBusConnection.
*/
func CastToDBusConnection(object *gobject.Object) *DBusConnection {
	return DBusConnectionNewFromNative(object.Native())
}

// Equals compares this DBusConnection with another DBusConnection, and returns true if they represent the same GObject.
func (recv *DBusConnection) Equals(other *DBusConnection) bool {
	return other.Native() == recv.Native()
}

func (recv *DBusConnection) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'g_dbus_connection_new_finish' : parameter 'res' of type 'AsyncResult' not supported

// UNSUPPORTED : C value 'g_dbus_connection_new_for_address_finish' : parameter 'res' of type 'AsyncResult' not supported

// UNSUPPORTED : C value 'g_dbus_connection_new_for_address_sync' : parameter 'flags' of type 'DBusConnectionFlags' not supported

// UNSUPPORTED : C value 'g_dbus_connection_new_sync' : parameter 'flags' of type 'DBusConnectionFlags' not supported

// UNSUPPORTED : C value 'g_dbus_connection_add_filter' : parameter 'filter_function' of type 'DBusMessageFilterFunction' not supported

// UNSUPPORTED : C value 'g_dbus_connection_call' : parameter 'parameters' of type 'GLib.Variant' not supported

// UNSUPPORTED : C value 'g_dbus_connection_call_finish' : parameter 'res' of type 'AsyncResult' not supported

// UNSUPPORTED : C value 'g_dbus_connection_call_sync' : parameter 'parameters' of type 'GLib.Variant' not supported

// UNSUPPORTED : C value 'g_dbus_connection_call_with_unix_fd_list' : parameter 'parameters' of type 'GLib.Variant' not supported

// UNSUPPORTED : C value 'g_dbus_connection_call_with_unix_fd_list_finish' : parameter 'res' of type 'AsyncResult' not supported

// UNSUPPORTED : C value 'g_dbus_connection_call_with_unix_fd_list_sync' : parameter 'parameters' of type 'GLib.Variant' not supported

// UNSUPPORTED : C value 'g_dbus_connection_close' : parameter 'callback' of type 'AsyncReadyCallback' not supported

// UNSUPPORTED : C value 'g_dbus_connection_close_finish' : parameter 'res' of type 'AsyncResult' not supported

var dBusConnectionCloseSyncFunction *gi.Function
var dBusConnectionCloseSyncFunction_Once sync.Once

func dBusConnectionCloseSyncFunction_Set() error {
	var err error
	dBusConnectionCloseSyncFunction_Once.Do(func() {
		err = dBusConnectionObject_Set()
		if err != nil {
			return
		}
		dBusConnectionCloseSyncFunction, err = dBusConnectionObject.InvokerNew("close_sync")
	})
	return err
}

// CloseSync is a representation of the C type g_dbus_connection_close_sync.
func (recv *DBusConnection) CloseSync(cancellable *Cancellable) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(cancellable.Native())

	var ret gi.Argument

	err := dBusConnectionCloseSyncFunction_Set()
	if err == nil {
		ret = dBusConnectionCloseSyncFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'g_dbus_connection_emit_signal' : parameter 'parameters' of type 'GLib.Variant' not supported

// UNSUPPORTED : C value 'g_dbus_connection_export_action_group' : parameter 'action_group' of type 'ActionGroup' not supported

var dBusConnectionExportMenuModelFunction *gi.Function
var dBusConnectionExportMenuModelFunction_Once sync.Once

func dBusConnectionExportMenuModelFunction_Set() error {
	var err error
	dBusConnectionExportMenuModelFunction_Once.Do(func() {
		err = dBusConnectionObject_Set()
		if err != nil {
			return
		}
		dBusConnectionExportMenuModelFunction, err = dBusConnectionObject.InvokerNew("export_menu_model")
	})
	return err
}

// ExportMenuModel is a representation of the C type g_dbus_connection_export_menu_model.
func (recv *DBusConnection) ExportMenuModel(objectPath string, menu *MenuModel) uint32 {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(objectPath)
	inArgs[2].SetPointer(menu.Native())

	var ret gi.Argument

	err := dBusConnectionExportMenuModelFunction_Set()
	if err == nil {
		ret = dBusConnectionExportMenuModelFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo
}

// UNSUPPORTED : C value 'g_dbus_connection_flush' : parameter 'callback' of type 'AsyncReadyCallback' not supported

// UNSUPPORTED : C value 'g_dbus_connection_flush_finish' : parameter 'res' of type 'AsyncResult' not supported

var dBusConnectionFlushSyncFunction *gi.Function
var dBusConnectionFlushSyncFunction_Once sync.Once

func dBusConnectionFlushSyncFunction_Set() error {
	var err error
	dBusConnectionFlushSyncFunction_Once.Do(func() {
		err = dBusConnectionObject_Set()
		if err != nil {
			return
		}
		dBusConnectionFlushSyncFunction, err = dBusConnectionObject.InvokerNew("flush_sync")
	})
	return err
}

// FlushSync is a representation of the C type g_dbus_connection_flush_sync.
func (recv *DBusConnection) FlushSync(cancellable *Cancellable) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(cancellable.Native())

	var ret gi.Argument

	err := dBusConnectionFlushSyncFunction_Set()
	if err == nil {
		ret = dBusConnectionFlushSyncFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'g_dbus_connection_get_capabilities' : return type 'DBusCapabilityFlags' not supported

var dBusConnectionGetExitOnCloseFunction *gi.Function
var dBusConnectionGetExitOnCloseFunction_Once sync.Once

func dBusConnectionGetExitOnCloseFunction_Set() error {
	var err error
	dBusConnectionGetExitOnCloseFunction_Once.Do(func() {
		err = dBusConnectionObject_Set()
		if err != nil {
			return
		}
		dBusConnectionGetExitOnCloseFunction, err = dBusConnectionObject.InvokerNew("get_exit_on_close")
	})
	return err
}

// GetExitOnClose is a representation of the C type g_dbus_connection_get_exit_on_close.
func (recv *DBusConnection) GetExitOnClose() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := dBusConnectionGetExitOnCloseFunction_Set()
	if err == nil {
		ret = dBusConnectionGetExitOnCloseFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'g_dbus_connection_get_flags' : return type 'DBusConnectionFlags' not supported

var dBusConnectionGetGuidFunction *gi.Function
var dBusConnectionGetGuidFunction_Once sync.Once

func dBusConnectionGetGuidFunction_Set() error {
	var err error
	dBusConnectionGetGuidFunction_Once.Do(func() {
		err = dBusConnectionObject_Set()
		if err != nil {
			return
		}
		dBusConnectionGetGuidFunction, err = dBusConnectionObject.InvokerNew("get_guid")
	})
	return err
}

// GetGuid is a representation of the C type g_dbus_connection_get_guid.
func (recv *DBusConnection) GetGuid() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := dBusConnectionGetGuidFunction_Set()
	if err == nil {
		ret = dBusConnectionGetGuidFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var dBusConnectionGetLastSerialFunction *gi.Function
var dBusConnectionGetLastSerialFunction_Once sync.Once

func dBusConnectionGetLastSerialFunction_Set() error {
	var err error
	dBusConnectionGetLastSerialFunction_Once.Do(func() {
		err = dBusConnectionObject_Set()
		if err != nil {
			return
		}
		dBusConnectionGetLastSerialFunction, err = dBusConnectionObject.InvokerNew("get_last_serial")
	})
	return err
}

// GetLastSerial is a representation of the C type g_dbus_connection_get_last_serial.
func (recv *DBusConnection) GetLastSerial() uint32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := dBusConnectionGetLastSerialFunction_Set()
	if err == nil {
		ret = dBusConnectionGetLastSerialFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo
}

var dBusConnectionGetPeerCredentialsFunction *gi.Function
var dBusConnectionGetPeerCredentialsFunction_Once sync.Once

func dBusConnectionGetPeerCredentialsFunction_Set() error {
	var err error
	dBusConnectionGetPeerCredentialsFunction_Once.Do(func() {
		err = dBusConnectionObject_Set()
		if err != nil {
			return
		}
		dBusConnectionGetPeerCredentialsFunction, err = dBusConnectionObject.InvokerNew("get_peer_credentials")
	})
	return err
}

// GetPeerCredentials is a representation of the C type g_dbus_connection_get_peer_credentials.
func (recv *DBusConnection) GetPeerCredentials() *Credentials {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := dBusConnectionGetPeerCredentialsFunction_Set()
	if err == nil {
		ret = dBusConnectionGetPeerCredentialsFunction.Invoke(inArgs[:], nil)
	}

	retGo := CredentialsNewFromNative(ret.Pointer())

	return retGo
}

var dBusConnectionGetStreamFunction *gi.Function
var dBusConnectionGetStreamFunction_Once sync.Once

func dBusConnectionGetStreamFunction_Set() error {
	var err error
	dBusConnectionGetStreamFunction_Once.Do(func() {
		err = dBusConnectionObject_Set()
		if err != nil {
			return
		}
		dBusConnectionGetStreamFunction, err = dBusConnectionObject.InvokerNew("get_stream")
	})
	return err
}

// GetStream is a representation of the C type g_dbus_connection_get_stream.
func (recv *DBusConnection) GetStream() *IOStream {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := dBusConnectionGetStreamFunction_Set()
	if err == nil {
		ret = dBusConnectionGetStreamFunction.Invoke(inArgs[:], nil)
	}

	retGo := IOStreamNewFromNative(ret.Pointer())

	return retGo
}

var dBusConnectionGetUniqueNameFunction *gi.Function
var dBusConnectionGetUniqueNameFunction_Once sync.Once

func dBusConnectionGetUniqueNameFunction_Set() error {
	var err error
	dBusConnectionGetUniqueNameFunction_Once.Do(func() {
		err = dBusConnectionObject_Set()
		if err != nil {
			return
		}
		dBusConnectionGetUniqueNameFunction, err = dBusConnectionObject.InvokerNew("get_unique_name")
	})
	return err
}

// GetUniqueName is a representation of the C type g_dbus_connection_get_unique_name.
func (recv *DBusConnection) GetUniqueName() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := dBusConnectionGetUniqueNameFunction_Set()
	if err == nil {
		ret = dBusConnectionGetUniqueNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var dBusConnectionIsClosedFunction *gi.Function
var dBusConnectionIsClosedFunction_Once sync.Once

func dBusConnectionIsClosedFunction_Set() error {
	var err error
	dBusConnectionIsClosedFunction_Once.Do(func() {
		err = dBusConnectionObject_Set()
		if err != nil {
			return
		}
		dBusConnectionIsClosedFunction, err = dBusConnectionObject.InvokerNew("is_closed")
	})
	return err
}

// IsClosed is a representation of the C type g_dbus_connection_is_closed.
func (recv *DBusConnection) IsClosed() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := dBusConnectionIsClosedFunction_Set()
	if err == nil {
		ret = dBusConnectionIsClosedFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'g_dbus_connection_register_object' : parameter 'user_data' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_dbus_connection_register_object_with_closures' : parameter 'method_call_closure' of type 'GObject.Closure' not supported

// UNSUPPORTED : C value 'g_dbus_connection_register_subtree' : parameter 'flags' of type 'DBusSubtreeFlags' not supported

var dBusConnectionRemoveFilterFunction *gi.Function
var dBusConnectionRemoveFilterFunction_Once sync.Once

func dBusConnectionRemoveFilterFunction_Set() error {
	var err error
	dBusConnectionRemoveFilterFunction_Once.Do(func() {
		err = dBusConnectionObject_Set()
		if err != nil {
			return
		}
		dBusConnectionRemoveFilterFunction, err = dBusConnectionObject.InvokerNew("remove_filter")
	})
	return err
}

// RemoveFilter is a representation of the C type g_dbus_connection_remove_filter.
func (recv *DBusConnection) RemoveFilter(filterId uint32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetUint32(filterId)

	err := dBusConnectionRemoveFilterFunction_Set()
	if err == nil {
		dBusConnectionRemoveFilterFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'g_dbus_connection_send_message' : parameter 'flags' of type 'DBusSendMessageFlags' not supported

// UNSUPPORTED : C value 'g_dbus_connection_send_message_with_reply' : parameter 'flags' of type 'DBusSendMessageFlags' not supported

// UNSUPPORTED : C value 'g_dbus_connection_send_message_with_reply_finish' : parameter 'res' of type 'AsyncResult' not supported

// UNSUPPORTED : C value 'g_dbus_connection_send_message_with_reply_sync' : parameter 'flags' of type 'DBusSendMessageFlags' not supported

var dBusConnectionSetExitOnCloseFunction *gi.Function
var dBusConnectionSetExitOnCloseFunction_Once sync.Once

func dBusConnectionSetExitOnCloseFunction_Set() error {
	var err error
	dBusConnectionSetExitOnCloseFunction_Once.Do(func() {
		err = dBusConnectionObject_Set()
		if err != nil {
			return
		}
		dBusConnectionSetExitOnCloseFunction, err = dBusConnectionObject.InvokerNew("set_exit_on_close")
	})
	return err
}

// SetExitOnClose is a representation of the C type g_dbus_connection_set_exit_on_close.
func (recv *DBusConnection) SetExitOnClose(exitOnClose bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetBoolean(exitOnClose)

	err := dBusConnectionSetExitOnCloseFunction_Set()
	if err == nil {
		dBusConnectionSetExitOnCloseFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'g_dbus_connection_signal_subscribe' : parameter 'flags' of type 'DBusSignalFlags' not supported

var dBusConnectionSignalUnsubscribeFunction *gi.Function
var dBusConnectionSignalUnsubscribeFunction_Once sync.Once

func dBusConnectionSignalUnsubscribeFunction_Set() error {
	var err error
	dBusConnectionSignalUnsubscribeFunction_Once.Do(func() {
		err = dBusConnectionObject_Set()
		if err != nil {
			return
		}
		dBusConnectionSignalUnsubscribeFunction, err = dBusConnectionObject.InvokerNew("signal_unsubscribe")
	})
	return err
}

// SignalUnsubscribe is a representation of the C type g_dbus_connection_signal_unsubscribe.
func (recv *DBusConnection) SignalUnsubscribe(subscriptionId uint32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetUint32(subscriptionId)

	err := dBusConnectionSignalUnsubscribeFunction_Set()
	if err == nil {
		dBusConnectionSignalUnsubscribeFunction.Invoke(inArgs[:], nil)
	}

	return
}

var dBusConnectionStartMessageProcessingFunction *gi.Function
var dBusConnectionStartMessageProcessingFunction_Once sync.Once

func dBusConnectionStartMessageProcessingFunction_Set() error {
	var err error
	dBusConnectionStartMessageProcessingFunction_Once.Do(func() {
		err = dBusConnectionObject_Set()
		if err != nil {
			return
		}
		dBusConnectionStartMessageProcessingFunction, err = dBusConnectionObject.InvokerNew("start_message_processing")
	})
	return err
}

// StartMessageProcessing is a representation of the C type g_dbus_connection_start_message_processing.
func (recv *DBusConnection) StartMessageProcessing() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := dBusConnectionStartMessageProcessingFunction_Set()
	if err == nil {
		dBusConnectionStartMessageProcessingFunction.Invoke(inArgs[:], nil)
	}

	return
}

var dBusConnectionUnexportActionGroupFunction *gi.Function
var dBusConnectionUnexportActionGroupFunction_Once sync.Once

func dBusConnectionUnexportActionGroupFunction_Set() error {
	var err error
	dBusConnectionUnexportActionGroupFunction_Once.Do(func() {
		err = dBusConnectionObject_Set()
		if err != nil {
			return
		}
		dBusConnectionUnexportActionGroupFunction, err = dBusConnectionObject.InvokerNew("unexport_action_group")
	})
	return err
}

// UnexportActionGroup is a representation of the C type g_dbus_connection_unexport_action_group.
func (recv *DBusConnection) UnexportActionGroup(exportId uint32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetUint32(exportId)

	err := dBusConnectionUnexportActionGroupFunction_Set()
	if err == nil {
		dBusConnectionUnexportActionGroupFunction.Invoke(inArgs[:], nil)
	}

	return
}

var dBusConnectionUnexportMenuModelFunction *gi.Function
var dBusConnectionUnexportMenuModelFunction_Once sync.Once

func dBusConnectionUnexportMenuModelFunction_Set() error {
	var err error
	dBusConnectionUnexportMenuModelFunction_Once.Do(func() {
		err = dBusConnectionObject_Set()
		if err != nil {
			return
		}
		dBusConnectionUnexportMenuModelFunction, err = dBusConnectionObject.InvokerNew("unexport_menu_model")
	})
	return err
}

// UnexportMenuModel is a representation of the C type g_dbus_connection_unexport_menu_model.
func (recv *DBusConnection) UnexportMenuModel(exportId uint32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetUint32(exportId)

	err := dBusConnectionUnexportMenuModelFunction_Set()
	if err == nil {
		dBusConnectionUnexportMenuModelFunction.Invoke(inArgs[:], nil)
	}

	return
}

var dBusConnectionUnregisterObjectFunction *gi.Function
var dBusConnectionUnregisterObjectFunction_Once sync.Once

func dBusConnectionUnregisterObjectFunction_Set() error {
	var err error
	dBusConnectionUnregisterObjectFunction_Once.Do(func() {
		err = dBusConnectionObject_Set()
		if err != nil {
			return
		}
		dBusConnectionUnregisterObjectFunction, err = dBusConnectionObject.InvokerNew("unregister_object")
	})
	return err
}

// UnregisterObject is a representation of the C type g_dbus_connection_unregister_object.
func (recv *DBusConnection) UnregisterObject(registrationId uint32) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetUint32(registrationId)

	var ret gi.Argument

	err := dBusConnectionUnregisterObjectFunction_Set()
	if err == nil {
		ret = dBusConnectionUnregisterObjectFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var dBusConnectionUnregisterSubtreeFunction *gi.Function
var dBusConnectionUnregisterSubtreeFunction_Once sync.Once

func dBusConnectionUnregisterSubtreeFunction_Set() error {
	var err error
	dBusConnectionUnregisterSubtreeFunction_Once.Do(func() {
		err = dBusConnectionObject_Set()
		if err != nil {
			return
		}
		dBusConnectionUnregisterSubtreeFunction, err = dBusConnectionObject.InvokerNew("unregister_subtree")
	})
	return err
}

// UnregisterSubtree is a representation of the C type g_dbus_connection_unregister_subtree.
func (recv *DBusConnection) UnregisterSubtree(registrationId uint32) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetUint32(registrationId)

	var ret gi.Argument

	err := dBusConnectionUnregisterSubtreeFunction_Set()
	if err == nil {
		ret = dBusConnectionUnregisterSubtreeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var dBusInterfaceSkeletonObject *gi.Object
var dBusInterfaceSkeletonObject_Once sync.Once

func dBusInterfaceSkeletonObject_Set() error {
	var err error
	dBusInterfaceSkeletonObject_Once.Do(func() {
		dBusInterfaceSkeletonObject, err = gi.ObjectNew("Gio", "DBusInterfaceSkeleton")
	})
	return err
}

type DBusInterfaceSkeleton struct {
	native unsafe.Pointer
}

func DBusInterfaceSkeletonNewFromNative(native unsafe.Pointer) *DBusInterfaceSkeleton {
	instance := &DBusInterfaceSkeleton{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *DBusInterfaceSkeleton) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToDBusInterfaceSkeleton down casts any arbitrary Object to DBusInterfaceSkeleton.
Exercise care, as this is a potentially dangerous function
if the Object is not a DBusInterfaceSkeleton.
*/
func CastToDBusInterfaceSkeleton(object *gobject.Object) *DBusInterfaceSkeleton {
	return DBusInterfaceSkeletonNewFromNative(object.Native())
}

// Equals compares this DBusInterfaceSkeleton with another DBusInterfaceSkeleton, and returns true if they represent the same GObject.
func (recv *DBusInterfaceSkeleton) Equals(other *DBusInterfaceSkeleton) bool {
	return other.Native() == recv.Native()
}

func (recv *DBusInterfaceSkeleton) Native() unsafe.Pointer {
	return recv.native
}

var dBusInterfaceSkeletonExportFunction *gi.Function
var dBusInterfaceSkeletonExportFunction_Once sync.Once

func dBusInterfaceSkeletonExportFunction_Set() error {
	var err error
	dBusInterfaceSkeletonExportFunction_Once.Do(func() {
		err = dBusInterfaceSkeletonObject_Set()
		if err != nil {
			return
		}
		dBusInterfaceSkeletonExportFunction, err = dBusInterfaceSkeletonObject.InvokerNew("export")
	})
	return err
}

// Export is a representation of the C type g_dbus_interface_skeleton_export.
func (recv *DBusInterfaceSkeleton) Export(connection *DBusConnection, objectPath string) bool {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(connection.Native())
	inArgs[2].SetString(objectPath)

	var ret gi.Argument

	err := dBusInterfaceSkeletonExportFunction_Set()
	if err == nil {
		ret = dBusInterfaceSkeletonExportFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var dBusInterfaceSkeletonFlushFunction *gi.Function
var dBusInterfaceSkeletonFlushFunction_Once sync.Once

func dBusInterfaceSkeletonFlushFunction_Set() error {
	var err error
	dBusInterfaceSkeletonFlushFunction_Once.Do(func() {
		err = dBusInterfaceSkeletonObject_Set()
		if err != nil {
			return
		}
		dBusInterfaceSkeletonFlushFunction, err = dBusInterfaceSkeletonObject.InvokerNew("flush")
	})
	return err
}

// Flush is a representation of the C type g_dbus_interface_skeleton_flush.
func (recv *DBusInterfaceSkeleton) Flush() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := dBusInterfaceSkeletonFlushFunction_Set()
	if err == nil {
		dBusInterfaceSkeletonFlushFunction.Invoke(inArgs[:], nil)
	}

	return
}

var dBusInterfaceSkeletonGetConnectionFunction *gi.Function
var dBusInterfaceSkeletonGetConnectionFunction_Once sync.Once

func dBusInterfaceSkeletonGetConnectionFunction_Set() error {
	var err error
	dBusInterfaceSkeletonGetConnectionFunction_Once.Do(func() {
		err = dBusInterfaceSkeletonObject_Set()
		if err != nil {
			return
		}
		dBusInterfaceSkeletonGetConnectionFunction, err = dBusInterfaceSkeletonObject.InvokerNew("get_connection")
	})
	return err
}

// GetConnection is a representation of the C type g_dbus_interface_skeleton_get_connection.
func (recv *DBusInterfaceSkeleton) GetConnection() *DBusConnection {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := dBusInterfaceSkeletonGetConnectionFunction_Set()
	if err == nil {
		ret = dBusInterfaceSkeletonGetConnectionFunction.Invoke(inArgs[:], nil)
	}

	retGo := DBusConnectionNewFromNative(ret.Pointer())

	return retGo
}

// UNSUPPORTED : C value 'g_dbus_interface_skeleton_get_connections' : return type 'GLib.List' not supported

// UNSUPPORTED : C value 'g_dbus_interface_skeleton_get_flags' : return type 'DBusInterfaceSkeletonFlags' not supported

var dBusInterfaceSkeletonGetInfoFunction *gi.Function
var dBusInterfaceSkeletonGetInfoFunction_Once sync.Once

func dBusInterfaceSkeletonGetInfoFunction_Set() error {
	var err error
	dBusInterfaceSkeletonGetInfoFunction_Once.Do(func() {
		err = dBusInterfaceSkeletonObject_Set()
		if err != nil {
			return
		}
		dBusInterfaceSkeletonGetInfoFunction, err = dBusInterfaceSkeletonObject.InvokerNew("get_info")
	})
	return err
}

// GetInfo is a representation of the C type g_dbus_interface_skeleton_get_info.
func (recv *DBusInterfaceSkeleton) GetInfo() *DBusInterfaceInfo {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := dBusInterfaceSkeletonGetInfoFunction_Set()
	if err == nil {
		ret = dBusInterfaceSkeletonGetInfoFunction.Invoke(inArgs[:], nil)
	}

	retGo := DBusInterfaceInfoNewFromNative(ret.Pointer())

	return retGo
}

var dBusInterfaceSkeletonGetObjectPathFunction *gi.Function
var dBusInterfaceSkeletonGetObjectPathFunction_Once sync.Once

func dBusInterfaceSkeletonGetObjectPathFunction_Set() error {
	var err error
	dBusInterfaceSkeletonGetObjectPathFunction_Once.Do(func() {
		err = dBusInterfaceSkeletonObject_Set()
		if err != nil {
			return
		}
		dBusInterfaceSkeletonGetObjectPathFunction, err = dBusInterfaceSkeletonObject.InvokerNew("get_object_path")
	})
	return err
}

// GetObjectPath is a representation of the C type g_dbus_interface_skeleton_get_object_path.
func (recv *DBusInterfaceSkeleton) GetObjectPath() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := dBusInterfaceSkeletonGetObjectPathFunction_Set()
	if err == nil {
		ret = dBusInterfaceSkeletonGetObjectPathFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'g_dbus_interface_skeleton_get_properties' : return type 'GLib.Variant' not supported

var dBusInterfaceSkeletonGetVtableFunction *gi.Function
var dBusInterfaceSkeletonGetVtableFunction_Once sync.Once

func dBusInterfaceSkeletonGetVtableFunction_Set() error {
	var err error
	dBusInterfaceSkeletonGetVtableFunction_Once.Do(func() {
		err = dBusInterfaceSkeletonObject_Set()
		if err != nil {
			return
		}
		dBusInterfaceSkeletonGetVtableFunction, err = dBusInterfaceSkeletonObject.InvokerNew("get_vtable")
	})
	return err
}

// GetVtable is a representation of the C type g_dbus_interface_skeleton_get_vtable.
func (recv *DBusInterfaceSkeleton) GetVtable() *DBusInterfaceVTable {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := dBusInterfaceSkeletonGetVtableFunction_Set()
	if err == nil {
		ret = dBusInterfaceSkeletonGetVtableFunction.Invoke(inArgs[:], nil)
	}

	retGo := DBusInterfaceVTableNewFromNative(ret.Pointer())

	return retGo
}

var dBusInterfaceSkeletonHasConnectionFunction *gi.Function
var dBusInterfaceSkeletonHasConnectionFunction_Once sync.Once

func dBusInterfaceSkeletonHasConnectionFunction_Set() error {
	var err error
	dBusInterfaceSkeletonHasConnectionFunction_Once.Do(func() {
		err = dBusInterfaceSkeletonObject_Set()
		if err != nil {
			return
		}
		dBusInterfaceSkeletonHasConnectionFunction, err = dBusInterfaceSkeletonObject.InvokerNew("has_connection")
	})
	return err
}

// HasConnection is a representation of the C type g_dbus_interface_skeleton_has_connection.
func (recv *DBusInterfaceSkeleton) HasConnection(connection *DBusConnection) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(connection.Native())

	var ret gi.Argument

	err := dBusInterfaceSkeletonHasConnectionFunction_Set()
	if err == nil {
		ret = dBusInterfaceSkeletonHasConnectionFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'g_dbus_interface_skeleton_set_flags' : parameter 'flags' of type 'DBusInterfaceSkeletonFlags' not supported

var dBusInterfaceSkeletonUnexportFunction *gi.Function
var dBusInterfaceSkeletonUnexportFunction_Once sync.Once

func dBusInterfaceSkeletonUnexportFunction_Set() error {
	var err error
	dBusInterfaceSkeletonUnexportFunction_Once.Do(func() {
		err = dBusInterfaceSkeletonObject_Set()
		if err != nil {
			return
		}
		dBusInterfaceSkeletonUnexportFunction, err = dBusInterfaceSkeletonObject.InvokerNew("unexport")
	})
	return err
}

// Unexport is a representation of the C type g_dbus_interface_skeleton_unexport.
func (recv *DBusInterfaceSkeleton) Unexport() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := dBusInterfaceSkeletonUnexportFunction_Set()
	if err == nil {
		dBusInterfaceSkeletonUnexportFunction.Invoke(inArgs[:], nil)
	}

	return
}

var dBusInterfaceSkeletonUnexportFromConnectionFunction *gi.Function
var dBusInterfaceSkeletonUnexportFromConnectionFunction_Once sync.Once

func dBusInterfaceSkeletonUnexportFromConnectionFunction_Set() error {
	var err error
	dBusInterfaceSkeletonUnexportFromConnectionFunction_Once.Do(func() {
		err = dBusInterfaceSkeletonObject_Set()
		if err != nil {
			return
		}
		dBusInterfaceSkeletonUnexportFromConnectionFunction, err = dBusInterfaceSkeletonObject.InvokerNew("unexport_from_connection")
	})
	return err
}

// UnexportFromConnection is a representation of the C type g_dbus_interface_skeleton_unexport_from_connection.
func (recv *DBusInterfaceSkeleton) UnexportFromConnection(connection *DBusConnection) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(connection.Native())

	err := dBusInterfaceSkeletonUnexportFromConnectionFunction_Set()
	if err == nil {
		dBusInterfaceSkeletonUnexportFromConnectionFunction.Invoke(inArgs[:], nil)
	}

	return
}

var dBusMenuModelObject *gi.Object
var dBusMenuModelObject_Once sync.Once

func dBusMenuModelObject_Set() error {
	var err error
	dBusMenuModelObject_Once.Do(func() {
		dBusMenuModelObject, err = gi.ObjectNew("Gio", "DBusMenuModel")
	})
	return err
}

type DBusMenuModel struct {
	native unsafe.Pointer
}

func DBusMenuModelNewFromNative(native unsafe.Pointer) *DBusMenuModel {
	instance := &DBusMenuModel{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// MenuModel upcasts to *MenuModel
func (recv *DBusMenuModel) MenuModel() *MenuModel {
	return MenuModelNewFromNative(recv.Native())
}

// Object upcasts to *Object
func (recv *DBusMenuModel) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToDBusMenuModel down casts any arbitrary Object to DBusMenuModel.
Exercise care, as this is a potentially dangerous function
if the Object is not a DBusMenuModel.
*/
func CastToDBusMenuModel(object *gobject.Object) *DBusMenuModel {
	return DBusMenuModelNewFromNative(object.Native())
}

// Equals compares this DBusMenuModel with another DBusMenuModel, and returns true if they represent the same GObject.
func (recv *DBusMenuModel) Equals(other *DBusMenuModel) bool {
	return other.Native() == recv.Native()
}

func (recv *DBusMenuModel) Native() unsafe.Pointer {
	return recv.native
}

var dBusMessageObject *gi.Object
var dBusMessageObject_Once sync.Once

func dBusMessageObject_Set() error {
	var err error
	dBusMessageObject_Once.Do(func() {
		dBusMessageObject, err = gi.ObjectNew("Gio", "DBusMessage")
	})
	return err
}

type DBusMessage struct {
	native unsafe.Pointer
}

func DBusMessageNewFromNative(native unsafe.Pointer) *DBusMessage {
	instance := &DBusMessage{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *DBusMessage) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToDBusMessage down casts any arbitrary Object to DBusMessage.
Exercise care, as this is a potentially dangerous function
if the Object is not a DBusMessage.
*/
func CastToDBusMessage(object *gobject.Object) *DBusMessage {
	return DBusMessageNewFromNative(object.Native())
}

// Equals compares this DBusMessage with another DBusMessage, and returns true if they represent the same GObject.
func (recv *DBusMessage) Equals(other *DBusMessage) bool {
	return other.Native() == recv.Native()
}

func (recv *DBusMessage) Native() unsafe.Pointer {
	return recv.native
}

var dBusMessageNewFunction *gi.Function
var dBusMessageNewFunction_Once sync.Once

func dBusMessageNewFunction_Set() error {
	var err error
	dBusMessageNewFunction_Once.Do(func() {
		err = dBusMessageObject_Set()
		if err != nil {
			return
		}
		dBusMessageNewFunction, err = dBusMessageObject.InvokerNew("new")
	})
	return err
}

// DBusMessageNew is a representation of the C type g_dbus_message_new.
func DBusMessageNew() *DBusMessage {

	var ret gi.Argument

	err := dBusMessageNewFunction_Set()
	if err == nil {
		ret = dBusMessageNewFunction.Invoke(nil, nil)
	}

	retGo := DBusMessageNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

// UNSUPPORTED : C value 'g_dbus_message_new_from_blob' : parameter 'blob' of type 'nil' not supported

var dBusMessageNewMethodCallFunction *gi.Function
var dBusMessageNewMethodCallFunction_Once sync.Once

func dBusMessageNewMethodCallFunction_Set() error {
	var err error
	dBusMessageNewMethodCallFunction_Once.Do(func() {
		err = dBusMessageObject_Set()
		if err != nil {
			return
		}
		dBusMessageNewMethodCallFunction, err = dBusMessageObject.InvokerNew("new_method_call")
	})
	return err
}

// DBusMessageNewMethodCall is a representation of the C type g_dbus_message_new_method_call.
func DBusMessageNewMethodCall(name string, path string, interface_ string, method string) *DBusMessage {
	var inArgs [4]gi.Argument
	inArgs[0].SetString(name)
	inArgs[1].SetString(path)
	inArgs[2].SetString(interface_)
	inArgs[3].SetString(method)

	var ret gi.Argument

	err := dBusMessageNewMethodCallFunction_Set()
	if err == nil {
		ret = dBusMessageNewMethodCallFunction.Invoke(inArgs[:], nil)
	}

	retGo := DBusMessageNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

var dBusMessageNewSignalFunction *gi.Function
var dBusMessageNewSignalFunction_Once sync.Once

func dBusMessageNewSignalFunction_Set() error {
	var err error
	dBusMessageNewSignalFunction_Once.Do(func() {
		err = dBusMessageObject_Set()
		if err != nil {
			return
		}
		dBusMessageNewSignalFunction, err = dBusMessageObject.InvokerNew("new_signal")
	})
	return err
}

// DBusMessageNewSignal is a representation of the C type g_dbus_message_new_signal.
func DBusMessageNewSignal(path string, interface_ string, signal string) *DBusMessage {
	var inArgs [3]gi.Argument
	inArgs[0].SetString(path)
	inArgs[1].SetString(interface_)
	inArgs[2].SetString(signal)

	var ret gi.Argument

	err := dBusMessageNewSignalFunction_Set()
	if err == nil {
		ret = dBusMessageNewSignalFunction.Invoke(inArgs[:], nil)
	}

	retGo := DBusMessageNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

var dBusMessageCopyFunction *gi.Function
var dBusMessageCopyFunction_Once sync.Once

func dBusMessageCopyFunction_Set() error {
	var err error
	dBusMessageCopyFunction_Once.Do(func() {
		err = dBusMessageObject_Set()
		if err != nil {
			return
		}
		dBusMessageCopyFunction, err = dBusMessageObject.InvokerNew("copy")
	})
	return err
}

// Copy is a representation of the C type g_dbus_message_copy.
func (recv *DBusMessage) Copy() *DBusMessage {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := dBusMessageCopyFunction_Set()
	if err == nil {
		ret = dBusMessageCopyFunction.Invoke(inArgs[:], nil)
	}

	retGo := DBusMessageNewFromNative(ret.Pointer())

	return retGo
}

var dBusMessageGetArg0Function *gi.Function
var dBusMessageGetArg0Function_Once sync.Once

func dBusMessageGetArg0Function_Set() error {
	var err error
	dBusMessageGetArg0Function_Once.Do(func() {
		err = dBusMessageObject_Set()
		if err != nil {
			return
		}
		dBusMessageGetArg0Function, err = dBusMessageObject.InvokerNew("get_arg0")
	})
	return err
}

// GetArg0 is a representation of the C type g_dbus_message_get_arg0.
func (recv *DBusMessage) GetArg0() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := dBusMessageGetArg0Function_Set()
	if err == nil {
		ret = dBusMessageGetArg0Function.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'g_dbus_message_get_body' : return type 'GLib.Variant' not supported

var dBusMessageGetByteOrderFunction *gi.Function
var dBusMessageGetByteOrderFunction_Once sync.Once

func dBusMessageGetByteOrderFunction_Set() error {
	var err error
	dBusMessageGetByteOrderFunction_Once.Do(func() {
		err = dBusMessageObject_Set()
		if err != nil {
			return
		}
		dBusMessageGetByteOrderFunction, err = dBusMessageObject.InvokerNew("get_byte_order")
	})
	return err
}

// GetByteOrder is a representation of the C type g_dbus_message_get_byte_order.
func (recv *DBusMessage) GetByteOrder() DBusMessageByteOrder {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := dBusMessageGetByteOrderFunction_Set()
	if err == nil {
		ret = dBusMessageGetByteOrderFunction.Invoke(inArgs[:], nil)
	}

	retGo := DBusMessageByteOrder(ret.Int32())

	return retGo
}

var dBusMessageGetDestinationFunction *gi.Function
var dBusMessageGetDestinationFunction_Once sync.Once

func dBusMessageGetDestinationFunction_Set() error {
	var err error
	dBusMessageGetDestinationFunction_Once.Do(func() {
		err = dBusMessageObject_Set()
		if err != nil {
			return
		}
		dBusMessageGetDestinationFunction, err = dBusMessageObject.InvokerNew("get_destination")
	})
	return err
}

// GetDestination is a representation of the C type g_dbus_message_get_destination.
func (recv *DBusMessage) GetDestination() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := dBusMessageGetDestinationFunction_Set()
	if err == nil {
		ret = dBusMessageGetDestinationFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var dBusMessageGetErrorNameFunction *gi.Function
var dBusMessageGetErrorNameFunction_Once sync.Once

func dBusMessageGetErrorNameFunction_Set() error {
	var err error
	dBusMessageGetErrorNameFunction_Once.Do(func() {
		err = dBusMessageObject_Set()
		if err != nil {
			return
		}
		dBusMessageGetErrorNameFunction, err = dBusMessageObject.InvokerNew("get_error_name")
	})
	return err
}

// GetErrorName is a representation of the C type g_dbus_message_get_error_name.
func (recv *DBusMessage) GetErrorName() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := dBusMessageGetErrorNameFunction_Set()
	if err == nil {
		ret = dBusMessageGetErrorNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'g_dbus_message_get_flags' : return type 'DBusMessageFlags' not supported

// UNSUPPORTED : C value 'g_dbus_message_get_header' : return type 'GLib.Variant' not supported

var dBusMessageGetHeaderFieldsFunction *gi.Function
var dBusMessageGetHeaderFieldsFunction_Once sync.Once

func dBusMessageGetHeaderFieldsFunction_Set() error {
	var err error
	dBusMessageGetHeaderFieldsFunction_Once.Do(func() {
		err = dBusMessageObject_Set()
		if err != nil {
			return
		}
		dBusMessageGetHeaderFieldsFunction, err = dBusMessageObject.InvokerNew("get_header_fields")
	})
	return err
}

// GetHeaderFields is a representation of the C type g_dbus_message_get_header_fields.
func (recv *DBusMessage) GetHeaderFields() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := dBusMessageGetHeaderFieldsFunction_Set()
	if err == nil {
		dBusMessageGetHeaderFieldsFunction.Invoke(inArgs[:], nil)
	}

	return
}

var dBusMessageGetInterfaceFunction *gi.Function
var dBusMessageGetInterfaceFunction_Once sync.Once

func dBusMessageGetInterfaceFunction_Set() error {
	var err error
	dBusMessageGetInterfaceFunction_Once.Do(func() {
		err = dBusMessageObject_Set()
		if err != nil {
			return
		}
		dBusMessageGetInterfaceFunction, err = dBusMessageObject.InvokerNew("get_interface")
	})
	return err
}

// GetInterface is a representation of the C type g_dbus_message_get_interface.
func (recv *DBusMessage) GetInterface() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := dBusMessageGetInterfaceFunction_Set()
	if err == nil {
		ret = dBusMessageGetInterfaceFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var dBusMessageGetLockedFunction *gi.Function
var dBusMessageGetLockedFunction_Once sync.Once

func dBusMessageGetLockedFunction_Set() error {
	var err error
	dBusMessageGetLockedFunction_Once.Do(func() {
		err = dBusMessageObject_Set()
		if err != nil {
			return
		}
		dBusMessageGetLockedFunction, err = dBusMessageObject.InvokerNew("get_locked")
	})
	return err
}

// GetLocked is a representation of the C type g_dbus_message_get_locked.
func (recv *DBusMessage) GetLocked() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := dBusMessageGetLockedFunction_Set()
	if err == nil {
		ret = dBusMessageGetLockedFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var dBusMessageGetMemberFunction *gi.Function
var dBusMessageGetMemberFunction_Once sync.Once

func dBusMessageGetMemberFunction_Set() error {
	var err error
	dBusMessageGetMemberFunction_Once.Do(func() {
		err = dBusMessageObject_Set()
		if err != nil {
			return
		}
		dBusMessageGetMemberFunction, err = dBusMessageObject.InvokerNew("get_member")
	})
	return err
}

// GetMember is a representation of the C type g_dbus_message_get_member.
func (recv *DBusMessage) GetMember() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := dBusMessageGetMemberFunction_Set()
	if err == nil {
		ret = dBusMessageGetMemberFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var dBusMessageGetMessageTypeFunction *gi.Function
var dBusMessageGetMessageTypeFunction_Once sync.Once

func dBusMessageGetMessageTypeFunction_Set() error {
	var err error
	dBusMessageGetMessageTypeFunction_Once.Do(func() {
		err = dBusMessageObject_Set()
		if err != nil {
			return
		}
		dBusMessageGetMessageTypeFunction, err = dBusMessageObject.InvokerNew("get_message_type")
	})
	return err
}

// GetMessageType is a representation of the C type g_dbus_message_get_message_type.
func (recv *DBusMessage) GetMessageType() DBusMessageType {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := dBusMessageGetMessageTypeFunction_Set()
	if err == nil {
		ret = dBusMessageGetMessageTypeFunction.Invoke(inArgs[:], nil)
	}

	retGo := DBusMessageType(ret.Int32())

	return retGo
}

var dBusMessageGetNumUnixFdsFunction *gi.Function
var dBusMessageGetNumUnixFdsFunction_Once sync.Once

func dBusMessageGetNumUnixFdsFunction_Set() error {
	var err error
	dBusMessageGetNumUnixFdsFunction_Once.Do(func() {
		err = dBusMessageObject_Set()
		if err != nil {
			return
		}
		dBusMessageGetNumUnixFdsFunction, err = dBusMessageObject.InvokerNew("get_num_unix_fds")
	})
	return err
}

// GetNumUnixFds is a representation of the C type g_dbus_message_get_num_unix_fds.
func (recv *DBusMessage) GetNumUnixFds() uint32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := dBusMessageGetNumUnixFdsFunction_Set()
	if err == nil {
		ret = dBusMessageGetNumUnixFdsFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo
}

var dBusMessageGetPathFunction *gi.Function
var dBusMessageGetPathFunction_Once sync.Once

func dBusMessageGetPathFunction_Set() error {
	var err error
	dBusMessageGetPathFunction_Once.Do(func() {
		err = dBusMessageObject_Set()
		if err != nil {
			return
		}
		dBusMessageGetPathFunction, err = dBusMessageObject.InvokerNew("get_path")
	})
	return err
}

// GetPath is a representation of the C type g_dbus_message_get_path.
func (recv *DBusMessage) GetPath() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := dBusMessageGetPathFunction_Set()
	if err == nil {
		ret = dBusMessageGetPathFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var dBusMessageGetReplySerialFunction *gi.Function
var dBusMessageGetReplySerialFunction_Once sync.Once

func dBusMessageGetReplySerialFunction_Set() error {
	var err error
	dBusMessageGetReplySerialFunction_Once.Do(func() {
		err = dBusMessageObject_Set()
		if err != nil {
			return
		}
		dBusMessageGetReplySerialFunction, err = dBusMessageObject.InvokerNew("get_reply_serial")
	})
	return err
}

// GetReplySerial is a representation of the C type g_dbus_message_get_reply_serial.
func (recv *DBusMessage) GetReplySerial() uint32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := dBusMessageGetReplySerialFunction_Set()
	if err == nil {
		ret = dBusMessageGetReplySerialFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo
}

var dBusMessageGetSenderFunction *gi.Function
var dBusMessageGetSenderFunction_Once sync.Once

func dBusMessageGetSenderFunction_Set() error {
	var err error
	dBusMessageGetSenderFunction_Once.Do(func() {
		err = dBusMessageObject_Set()
		if err != nil {
			return
		}
		dBusMessageGetSenderFunction, err = dBusMessageObject.InvokerNew("get_sender")
	})
	return err
}

// GetSender is a representation of the C type g_dbus_message_get_sender.
func (recv *DBusMessage) GetSender() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := dBusMessageGetSenderFunction_Set()
	if err == nil {
		ret = dBusMessageGetSenderFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var dBusMessageGetSerialFunction *gi.Function
var dBusMessageGetSerialFunction_Once sync.Once

func dBusMessageGetSerialFunction_Set() error {
	var err error
	dBusMessageGetSerialFunction_Once.Do(func() {
		err = dBusMessageObject_Set()
		if err != nil {
			return
		}
		dBusMessageGetSerialFunction, err = dBusMessageObject.InvokerNew("get_serial")
	})
	return err
}

// GetSerial is a representation of the C type g_dbus_message_get_serial.
func (recv *DBusMessage) GetSerial() uint32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := dBusMessageGetSerialFunction_Set()
	if err == nil {
		ret = dBusMessageGetSerialFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo
}

var dBusMessageGetSignatureFunction *gi.Function
var dBusMessageGetSignatureFunction_Once sync.Once

func dBusMessageGetSignatureFunction_Set() error {
	var err error
	dBusMessageGetSignatureFunction_Once.Do(func() {
		err = dBusMessageObject_Set()
		if err != nil {
			return
		}
		dBusMessageGetSignatureFunction, err = dBusMessageObject.InvokerNew("get_signature")
	})
	return err
}

// GetSignature is a representation of the C type g_dbus_message_get_signature.
func (recv *DBusMessage) GetSignature() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := dBusMessageGetSignatureFunction_Set()
	if err == nil {
		ret = dBusMessageGetSignatureFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var dBusMessageGetUnixFdListFunction *gi.Function
var dBusMessageGetUnixFdListFunction_Once sync.Once

func dBusMessageGetUnixFdListFunction_Set() error {
	var err error
	dBusMessageGetUnixFdListFunction_Once.Do(func() {
		err = dBusMessageObject_Set()
		if err != nil {
			return
		}
		dBusMessageGetUnixFdListFunction, err = dBusMessageObject.InvokerNew("get_unix_fd_list")
	})
	return err
}

// GetUnixFdList is a representation of the C type g_dbus_message_get_unix_fd_list.
func (recv *DBusMessage) GetUnixFdList() *UnixFDList {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := dBusMessageGetUnixFdListFunction_Set()
	if err == nil {
		ret = dBusMessageGetUnixFdListFunction.Invoke(inArgs[:], nil)
	}

	retGo := UnixFDListNewFromNative(ret.Pointer())

	return retGo
}

var dBusMessageLockFunction *gi.Function
var dBusMessageLockFunction_Once sync.Once

func dBusMessageLockFunction_Set() error {
	var err error
	dBusMessageLockFunction_Once.Do(func() {
		err = dBusMessageObject_Set()
		if err != nil {
			return
		}
		dBusMessageLockFunction, err = dBusMessageObject.InvokerNew("lock")
	})
	return err
}

// Lock is a representation of the C type g_dbus_message_lock.
func (recv *DBusMessage) Lock() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := dBusMessageLockFunction_Set()
	if err == nil {
		dBusMessageLockFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'g_dbus_message_new_method_error' : parameter '...' of type 'nil' not supported

var dBusMessageNewMethodErrorLiteralFunction *gi.Function
var dBusMessageNewMethodErrorLiteralFunction_Once sync.Once

func dBusMessageNewMethodErrorLiteralFunction_Set() error {
	var err error
	dBusMessageNewMethodErrorLiteralFunction_Once.Do(func() {
		err = dBusMessageObject_Set()
		if err != nil {
			return
		}
		dBusMessageNewMethodErrorLiteralFunction, err = dBusMessageObject.InvokerNew("new_method_error_literal")
	})
	return err
}

// NewMethodErrorLiteral is a representation of the C type g_dbus_message_new_method_error_literal.
func (recv *DBusMessage) NewMethodErrorLiteral(errorName string, errorMessage string) *DBusMessage {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(errorName)
	inArgs[2].SetString(errorMessage)

	var ret gi.Argument

	err := dBusMessageNewMethodErrorLiteralFunction_Set()
	if err == nil {
		ret = dBusMessageNewMethodErrorLiteralFunction.Invoke(inArgs[:], nil)
	}

	retGo := DBusMessageNewFromNative(ret.Pointer())

	return retGo
}

// UNSUPPORTED : C value 'g_dbus_message_new_method_error_valist' : parameter 'var_args' of type 'va_list' not supported

var dBusMessageNewMethodReplyFunction *gi.Function
var dBusMessageNewMethodReplyFunction_Once sync.Once

func dBusMessageNewMethodReplyFunction_Set() error {
	var err error
	dBusMessageNewMethodReplyFunction_Once.Do(func() {
		err = dBusMessageObject_Set()
		if err != nil {
			return
		}
		dBusMessageNewMethodReplyFunction, err = dBusMessageObject.InvokerNew("new_method_reply")
	})
	return err
}

// NewMethodReply is a representation of the C type g_dbus_message_new_method_reply.
func (recv *DBusMessage) NewMethodReply() *DBusMessage {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := dBusMessageNewMethodReplyFunction_Set()
	if err == nil {
		ret = dBusMessageNewMethodReplyFunction.Invoke(inArgs[:], nil)
	}

	retGo := DBusMessageNewFromNative(ret.Pointer())

	return retGo
}

var dBusMessagePrintFunction *gi.Function
var dBusMessagePrintFunction_Once sync.Once

func dBusMessagePrintFunction_Set() error {
	var err error
	dBusMessagePrintFunction_Once.Do(func() {
		err = dBusMessageObject_Set()
		if err != nil {
			return
		}
		dBusMessagePrintFunction, err = dBusMessageObject.InvokerNew("print")
	})
	return err
}

// Print is a representation of the C type g_dbus_message_print.
func (recv *DBusMessage) Print(indent uint32) string {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetUint32(indent)

	var ret gi.Argument

	err := dBusMessagePrintFunction_Set()
	if err == nil {
		ret = dBusMessagePrintFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

// UNSUPPORTED : C value 'g_dbus_message_set_body' : parameter 'body' of type 'GLib.Variant' not supported

var dBusMessageSetByteOrderFunction *gi.Function
var dBusMessageSetByteOrderFunction_Once sync.Once

func dBusMessageSetByteOrderFunction_Set() error {
	var err error
	dBusMessageSetByteOrderFunction_Once.Do(func() {
		err = dBusMessageObject_Set()
		if err != nil {
			return
		}
		dBusMessageSetByteOrderFunction, err = dBusMessageObject.InvokerNew("set_byte_order")
	})
	return err
}

// SetByteOrder is a representation of the C type g_dbus_message_set_byte_order.
func (recv *DBusMessage) SetByteOrder(byteOrder DBusMessageByteOrder) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(int32(byteOrder))

	err := dBusMessageSetByteOrderFunction_Set()
	if err == nil {
		dBusMessageSetByteOrderFunction.Invoke(inArgs[:], nil)
	}

	return
}

var dBusMessageSetDestinationFunction *gi.Function
var dBusMessageSetDestinationFunction_Once sync.Once

func dBusMessageSetDestinationFunction_Set() error {
	var err error
	dBusMessageSetDestinationFunction_Once.Do(func() {
		err = dBusMessageObject_Set()
		if err != nil {
			return
		}
		dBusMessageSetDestinationFunction, err = dBusMessageObject.InvokerNew("set_destination")
	})
	return err
}

// SetDestination is a representation of the C type g_dbus_message_set_destination.
func (recv *DBusMessage) SetDestination(value string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(value)

	err := dBusMessageSetDestinationFunction_Set()
	if err == nil {
		dBusMessageSetDestinationFunction.Invoke(inArgs[:], nil)
	}

	return
}

var dBusMessageSetErrorNameFunction *gi.Function
var dBusMessageSetErrorNameFunction_Once sync.Once

func dBusMessageSetErrorNameFunction_Set() error {
	var err error
	dBusMessageSetErrorNameFunction_Once.Do(func() {
		err = dBusMessageObject_Set()
		if err != nil {
			return
		}
		dBusMessageSetErrorNameFunction, err = dBusMessageObject.InvokerNew("set_error_name")
	})
	return err
}

// SetErrorName is a representation of the C type g_dbus_message_set_error_name.
func (recv *DBusMessage) SetErrorName(value string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(value)

	err := dBusMessageSetErrorNameFunction_Set()
	if err == nil {
		dBusMessageSetErrorNameFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'g_dbus_message_set_flags' : parameter 'flags' of type 'DBusMessageFlags' not supported

// UNSUPPORTED : C value 'g_dbus_message_set_header' : parameter 'value' of type 'GLib.Variant' not supported

var dBusMessageSetInterfaceFunction *gi.Function
var dBusMessageSetInterfaceFunction_Once sync.Once

func dBusMessageSetInterfaceFunction_Set() error {
	var err error
	dBusMessageSetInterfaceFunction_Once.Do(func() {
		err = dBusMessageObject_Set()
		if err != nil {
			return
		}
		dBusMessageSetInterfaceFunction, err = dBusMessageObject.InvokerNew("set_interface")
	})
	return err
}

// SetInterface is a representation of the C type g_dbus_message_set_interface.
func (recv *DBusMessage) SetInterface(value string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(value)

	err := dBusMessageSetInterfaceFunction_Set()
	if err == nil {
		dBusMessageSetInterfaceFunction.Invoke(inArgs[:], nil)
	}

	return
}

var dBusMessageSetMemberFunction *gi.Function
var dBusMessageSetMemberFunction_Once sync.Once

func dBusMessageSetMemberFunction_Set() error {
	var err error
	dBusMessageSetMemberFunction_Once.Do(func() {
		err = dBusMessageObject_Set()
		if err != nil {
			return
		}
		dBusMessageSetMemberFunction, err = dBusMessageObject.InvokerNew("set_member")
	})
	return err
}

// SetMember is a representation of the C type g_dbus_message_set_member.
func (recv *DBusMessage) SetMember(value string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(value)

	err := dBusMessageSetMemberFunction_Set()
	if err == nil {
		dBusMessageSetMemberFunction.Invoke(inArgs[:], nil)
	}

	return
}

var dBusMessageSetMessageTypeFunction *gi.Function
var dBusMessageSetMessageTypeFunction_Once sync.Once

func dBusMessageSetMessageTypeFunction_Set() error {
	var err error
	dBusMessageSetMessageTypeFunction_Once.Do(func() {
		err = dBusMessageObject_Set()
		if err != nil {
			return
		}
		dBusMessageSetMessageTypeFunction, err = dBusMessageObject.InvokerNew("set_message_type")
	})
	return err
}

// SetMessageType is a representation of the C type g_dbus_message_set_message_type.
func (recv *DBusMessage) SetMessageType(type_ DBusMessageType) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(int32(type_))

	err := dBusMessageSetMessageTypeFunction_Set()
	if err == nil {
		dBusMessageSetMessageTypeFunction.Invoke(inArgs[:], nil)
	}

	return
}

var dBusMessageSetNumUnixFdsFunction *gi.Function
var dBusMessageSetNumUnixFdsFunction_Once sync.Once

func dBusMessageSetNumUnixFdsFunction_Set() error {
	var err error
	dBusMessageSetNumUnixFdsFunction_Once.Do(func() {
		err = dBusMessageObject_Set()
		if err != nil {
			return
		}
		dBusMessageSetNumUnixFdsFunction, err = dBusMessageObject.InvokerNew("set_num_unix_fds")
	})
	return err
}

// SetNumUnixFds is a representation of the C type g_dbus_message_set_num_unix_fds.
func (recv *DBusMessage) SetNumUnixFds(value uint32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetUint32(value)

	err := dBusMessageSetNumUnixFdsFunction_Set()
	if err == nil {
		dBusMessageSetNumUnixFdsFunction.Invoke(inArgs[:], nil)
	}

	return
}

var dBusMessageSetPathFunction *gi.Function
var dBusMessageSetPathFunction_Once sync.Once

func dBusMessageSetPathFunction_Set() error {
	var err error
	dBusMessageSetPathFunction_Once.Do(func() {
		err = dBusMessageObject_Set()
		if err != nil {
			return
		}
		dBusMessageSetPathFunction, err = dBusMessageObject.InvokerNew("set_path")
	})
	return err
}

// SetPath is a representation of the C type g_dbus_message_set_path.
func (recv *DBusMessage) SetPath(value string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(value)

	err := dBusMessageSetPathFunction_Set()
	if err == nil {
		dBusMessageSetPathFunction.Invoke(inArgs[:], nil)
	}

	return
}

var dBusMessageSetReplySerialFunction *gi.Function
var dBusMessageSetReplySerialFunction_Once sync.Once

func dBusMessageSetReplySerialFunction_Set() error {
	var err error
	dBusMessageSetReplySerialFunction_Once.Do(func() {
		err = dBusMessageObject_Set()
		if err != nil {
			return
		}
		dBusMessageSetReplySerialFunction, err = dBusMessageObject.InvokerNew("set_reply_serial")
	})
	return err
}

// SetReplySerial is a representation of the C type g_dbus_message_set_reply_serial.
func (recv *DBusMessage) SetReplySerial(value uint32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetUint32(value)

	err := dBusMessageSetReplySerialFunction_Set()
	if err == nil {
		dBusMessageSetReplySerialFunction.Invoke(inArgs[:], nil)
	}

	return
}

var dBusMessageSetSenderFunction *gi.Function
var dBusMessageSetSenderFunction_Once sync.Once

func dBusMessageSetSenderFunction_Set() error {
	var err error
	dBusMessageSetSenderFunction_Once.Do(func() {
		err = dBusMessageObject_Set()
		if err != nil {
			return
		}
		dBusMessageSetSenderFunction, err = dBusMessageObject.InvokerNew("set_sender")
	})
	return err
}

// SetSender is a representation of the C type g_dbus_message_set_sender.
func (recv *DBusMessage) SetSender(value string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(value)

	err := dBusMessageSetSenderFunction_Set()
	if err == nil {
		dBusMessageSetSenderFunction.Invoke(inArgs[:], nil)
	}

	return
}

var dBusMessageSetSerialFunction *gi.Function
var dBusMessageSetSerialFunction_Once sync.Once

func dBusMessageSetSerialFunction_Set() error {
	var err error
	dBusMessageSetSerialFunction_Once.Do(func() {
		err = dBusMessageObject_Set()
		if err != nil {
			return
		}
		dBusMessageSetSerialFunction, err = dBusMessageObject.InvokerNew("set_serial")
	})
	return err
}

// SetSerial is a representation of the C type g_dbus_message_set_serial.
func (recv *DBusMessage) SetSerial(serial uint32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetUint32(serial)

	err := dBusMessageSetSerialFunction_Set()
	if err == nil {
		dBusMessageSetSerialFunction.Invoke(inArgs[:], nil)
	}

	return
}

var dBusMessageSetSignatureFunction *gi.Function
var dBusMessageSetSignatureFunction_Once sync.Once

func dBusMessageSetSignatureFunction_Set() error {
	var err error
	dBusMessageSetSignatureFunction_Once.Do(func() {
		err = dBusMessageObject_Set()
		if err != nil {
			return
		}
		dBusMessageSetSignatureFunction, err = dBusMessageObject.InvokerNew("set_signature")
	})
	return err
}

// SetSignature is a representation of the C type g_dbus_message_set_signature.
func (recv *DBusMessage) SetSignature(value string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(value)

	err := dBusMessageSetSignatureFunction_Set()
	if err == nil {
		dBusMessageSetSignatureFunction.Invoke(inArgs[:], nil)
	}

	return
}

var dBusMessageSetUnixFdListFunction *gi.Function
var dBusMessageSetUnixFdListFunction_Once sync.Once

func dBusMessageSetUnixFdListFunction_Set() error {
	var err error
	dBusMessageSetUnixFdListFunction_Once.Do(func() {
		err = dBusMessageObject_Set()
		if err != nil {
			return
		}
		dBusMessageSetUnixFdListFunction, err = dBusMessageObject.InvokerNew("set_unix_fd_list")
	})
	return err
}

// SetUnixFdList is a representation of the C type g_dbus_message_set_unix_fd_list.
func (recv *DBusMessage) SetUnixFdList(fdList *UnixFDList) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(fdList.Native())

	err := dBusMessageSetUnixFdListFunction_Set()
	if err == nil {
		dBusMessageSetUnixFdListFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'g_dbus_message_to_blob' : parameter 'capabilities' of type 'DBusCapabilityFlags' not supported

var dBusMessageToGerrorFunction *gi.Function
var dBusMessageToGerrorFunction_Once sync.Once

func dBusMessageToGerrorFunction_Set() error {
	var err error
	dBusMessageToGerrorFunction_Once.Do(func() {
		err = dBusMessageObject_Set()
		if err != nil {
			return
		}
		dBusMessageToGerrorFunction, err = dBusMessageObject.InvokerNew("to_gerror")
	})
	return err
}

// ToGerror is a representation of the C type g_dbus_message_to_gerror.
func (recv *DBusMessage) ToGerror() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := dBusMessageToGerrorFunction_Set()
	if err == nil {
		ret = dBusMessageToGerrorFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var dBusMethodInvocationObject *gi.Object
var dBusMethodInvocationObject_Once sync.Once

func dBusMethodInvocationObject_Set() error {
	var err error
	dBusMethodInvocationObject_Once.Do(func() {
		dBusMethodInvocationObject, err = gi.ObjectNew("Gio", "DBusMethodInvocation")
	})
	return err
}

type DBusMethodInvocation struct {
	native unsafe.Pointer
}

func DBusMethodInvocationNewFromNative(native unsafe.Pointer) *DBusMethodInvocation {
	instance := &DBusMethodInvocation{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *DBusMethodInvocation) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToDBusMethodInvocation down casts any arbitrary Object to DBusMethodInvocation.
Exercise care, as this is a potentially dangerous function
if the Object is not a DBusMethodInvocation.
*/
func CastToDBusMethodInvocation(object *gobject.Object) *DBusMethodInvocation {
	return DBusMethodInvocationNewFromNative(object.Native())
}

// Equals compares this DBusMethodInvocation with another DBusMethodInvocation, and returns true if they represent the same GObject.
func (recv *DBusMethodInvocation) Equals(other *DBusMethodInvocation) bool {
	return other.Native() == recv.Native()
}

func (recv *DBusMethodInvocation) Native() unsafe.Pointer {
	return recv.native
}

var dBusMethodInvocationGetConnectionFunction *gi.Function
var dBusMethodInvocationGetConnectionFunction_Once sync.Once

func dBusMethodInvocationGetConnectionFunction_Set() error {
	var err error
	dBusMethodInvocationGetConnectionFunction_Once.Do(func() {
		err = dBusMethodInvocationObject_Set()
		if err != nil {
			return
		}
		dBusMethodInvocationGetConnectionFunction, err = dBusMethodInvocationObject.InvokerNew("get_connection")
	})
	return err
}

// GetConnection is a representation of the C type g_dbus_method_invocation_get_connection.
func (recv *DBusMethodInvocation) GetConnection() *DBusConnection {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := dBusMethodInvocationGetConnectionFunction_Set()
	if err == nil {
		ret = dBusMethodInvocationGetConnectionFunction.Invoke(inArgs[:], nil)
	}

	retGo := DBusConnectionNewFromNative(ret.Pointer())

	return retGo
}

var dBusMethodInvocationGetInterfaceNameFunction *gi.Function
var dBusMethodInvocationGetInterfaceNameFunction_Once sync.Once

func dBusMethodInvocationGetInterfaceNameFunction_Set() error {
	var err error
	dBusMethodInvocationGetInterfaceNameFunction_Once.Do(func() {
		err = dBusMethodInvocationObject_Set()
		if err != nil {
			return
		}
		dBusMethodInvocationGetInterfaceNameFunction, err = dBusMethodInvocationObject.InvokerNew("get_interface_name")
	})
	return err
}

// GetInterfaceName is a representation of the C type g_dbus_method_invocation_get_interface_name.
func (recv *DBusMethodInvocation) GetInterfaceName() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := dBusMethodInvocationGetInterfaceNameFunction_Set()
	if err == nil {
		ret = dBusMethodInvocationGetInterfaceNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var dBusMethodInvocationGetMessageFunction *gi.Function
var dBusMethodInvocationGetMessageFunction_Once sync.Once

func dBusMethodInvocationGetMessageFunction_Set() error {
	var err error
	dBusMethodInvocationGetMessageFunction_Once.Do(func() {
		err = dBusMethodInvocationObject_Set()
		if err != nil {
			return
		}
		dBusMethodInvocationGetMessageFunction, err = dBusMethodInvocationObject.InvokerNew("get_message")
	})
	return err
}

// GetMessage is a representation of the C type g_dbus_method_invocation_get_message.
func (recv *DBusMethodInvocation) GetMessage() *DBusMessage {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := dBusMethodInvocationGetMessageFunction_Set()
	if err == nil {
		ret = dBusMethodInvocationGetMessageFunction.Invoke(inArgs[:], nil)
	}

	retGo := DBusMessageNewFromNative(ret.Pointer())

	return retGo
}

var dBusMethodInvocationGetMethodInfoFunction *gi.Function
var dBusMethodInvocationGetMethodInfoFunction_Once sync.Once

func dBusMethodInvocationGetMethodInfoFunction_Set() error {
	var err error
	dBusMethodInvocationGetMethodInfoFunction_Once.Do(func() {
		err = dBusMethodInvocationObject_Set()
		if err != nil {
			return
		}
		dBusMethodInvocationGetMethodInfoFunction, err = dBusMethodInvocationObject.InvokerNew("get_method_info")
	})
	return err
}

// GetMethodInfo is a representation of the C type g_dbus_method_invocation_get_method_info.
func (recv *DBusMethodInvocation) GetMethodInfo() *DBusMethodInfo {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := dBusMethodInvocationGetMethodInfoFunction_Set()
	if err == nil {
		ret = dBusMethodInvocationGetMethodInfoFunction.Invoke(inArgs[:], nil)
	}

	retGo := DBusMethodInfoNewFromNative(ret.Pointer())

	return retGo
}

var dBusMethodInvocationGetMethodNameFunction *gi.Function
var dBusMethodInvocationGetMethodNameFunction_Once sync.Once

func dBusMethodInvocationGetMethodNameFunction_Set() error {
	var err error
	dBusMethodInvocationGetMethodNameFunction_Once.Do(func() {
		err = dBusMethodInvocationObject_Set()
		if err != nil {
			return
		}
		dBusMethodInvocationGetMethodNameFunction, err = dBusMethodInvocationObject.InvokerNew("get_method_name")
	})
	return err
}

// GetMethodName is a representation of the C type g_dbus_method_invocation_get_method_name.
func (recv *DBusMethodInvocation) GetMethodName() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := dBusMethodInvocationGetMethodNameFunction_Set()
	if err == nil {
		ret = dBusMethodInvocationGetMethodNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var dBusMethodInvocationGetObjectPathFunction *gi.Function
var dBusMethodInvocationGetObjectPathFunction_Once sync.Once

func dBusMethodInvocationGetObjectPathFunction_Set() error {
	var err error
	dBusMethodInvocationGetObjectPathFunction_Once.Do(func() {
		err = dBusMethodInvocationObject_Set()
		if err != nil {
			return
		}
		dBusMethodInvocationGetObjectPathFunction, err = dBusMethodInvocationObject.InvokerNew("get_object_path")
	})
	return err
}

// GetObjectPath is a representation of the C type g_dbus_method_invocation_get_object_path.
func (recv *DBusMethodInvocation) GetObjectPath() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := dBusMethodInvocationGetObjectPathFunction_Set()
	if err == nil {
		ret = dBusMethodInvocationGetObjectPathFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'g_dbus_method_invocation_get_parameters' : return type 'GLib.Variant' not supported

var dBusMethodInvocationGetPropertyInfoFunction *gi.Function
var dBusMethodInvocationGetPropertyInfoFunction_Once sync.Once

func dBusMethodInvocationGetPropertyInfoFunction_Set() error {
	var err error
	dBusMethodInvocationGetPropertyInfoFunction_Once.Do(func() {
		err = dBusMethodInvocationObject_Set()
		if err != nil {
			return
		}
		dBusMethodInvocationGetPropertyInfoFunction, err = dBusMethodInvocationObject.InvokerNew("get_property_info")
	})
	return err
}

// GetPropertyInfo is a representation of the C type g_dbus_method_invocation_get_property_info.
func (recv *DBusMethodInvocation) GetPropertyInfo() *DBusPropertyInfo {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := dBusMethodInvocationGetPropertyInfoFunction_Set()
	if err == nil {
		ret = dBusMethodInvocationGetPropertyInfoFunction.Invoke(inArgs[:], nil)
	}

	retGo := DBusPropertyInfoNewFromNative(ret.Pointer())

	return retGo
}

var dBusMethodInvocationGetSenderFunction *gi.Function
var dBusMethodInvocationGetSenderFunction_Once sync.Once

func dBusMethodInvocationGetSenderFunction_Set() error {
	var err error
	dBusMethodInvocationGetSenderFunction_Once.Do(func() {
		err = dBusMethodInvocationObject_Set()
		if err != nil {
			return
		}
		dBusMethodInvocationGetSenderFunction, err = dBusMethodInvocationObject.InvokerNew("get_sender")
	})
	return err
}

// GetSender is a representation of the C type g_dbus_method_invocation_get_sender.
func (recv *DBusMethodInvocation) GetSender() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := dBusMethodInvocationGetSenderFunction_Set()
	if err == nil {
		ret = dBusMethodInvocationGetSenderFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'g_dbus_method_invocation_get_user_data' : return type 'gpointer' not supported

var dBusMethodInvocationReturnDbusErrorFunction *gi.Function
var dBusMethodInvocationReturnDbusErrorFunction_Once sync.Once

func dBusMethodInvocationReturnDbusErrorFunction_Set() error {
	var err error
	dBusMethodInvocationReturnDbusErrorFunction_Once.Do(func() {
		err = dBusMethodInvocationObject_Set()
		if err != nil {
			return
		}
		dBusMethodInvocationReturnDbusErrorFunction, err = dBusMethodInvocationObject.InvokerNew("return_dbus_error")
	})
	return err
}

// ReturnDbusError is a representation of the C type g_dbus_method_invocation_return_dbus_error.
func (recv *DBusMethodInvocation) ReturnDbusError(errorName string, errorMessage string) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(errorName)
	inArgs[2].SetString(errorMessage)

	err := dBusMethodInvocationReturnDbusErrorFunction_Set()
	if err == nil {
		dBusMethodInvocationReturnDbusErrorFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'g_dbus_method_invocation_return_error' : parameter '...' of type 'nil' not supported

var dBusMethodInvocationReturnErrorLiteralFunction *gi.Function
var dBusMethodInvocationReturnErrorLiteralFunction_Once sync.Once

func dBusMethodInvocationReturnErrorLiteralFunction_Set() error {
	var err error
	dBusMethodInvocationReturnErrorLiteralFunction_Once.Do(func() {
		err = dBusMethodInvocationObject_Set()
		if err != nil {
			return
		}
		dBusMethodInvocationReturnErrorLiteralFunction, err = dBusMethodInvocationObject.InvokerNew("return_error_literal")
	})
	return err
}

// ReturnErrorLiteral is a representation of the C type g_dbus_method_invocation_return_error_literal.
func (recv *DBusMethodInvocation) ReturnErrorLiteral(domain glib.Quark, code int32, message string) {
	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetUint32(uint32(domain))
	inArgs[2].SetInt32(code)
	inArgs[3].SetString(message)

	err := dBusMethodInvocationReturnErrorLiteralFunction_Set()
	if err == nil {
		dBusMethodInvocationReturnErrorLiteralFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'g_dbus_method_invocation_return_error_valist' : parameter 'var_args' of type 'va_list' not supported

// UNSUPPORTED : C value 'g_dbus_method_invocation_return_gerror' : parameter 'error' of type 'GLib.Error' not supported

// UNSUPPORTED : C value 'g_dbus_method_invocation_return_value' : parameter 'parameters' of type 'GLib.Variant' not supported

// UNSUPPORTED : C value 'g_dbus_method_invocation_return_value_with_unix_fd_list' : parameter 'parameters' of type 'GLib.Variant' not supported

// UNSUPPORTED : C value 'g_dbus_method_invocation_take_error' : parameter 'error' of type 'GLib.Error' not supported

var dBusObjectManagerClientObject *gi.Object
var dBusObjectManagerClientObject_Once sync.Once

func dBusObjectManagerClientObject_Set() error {
	var err error
	dBusObjectManagerClientObject_Once.Do(func() {
		dBusObjectManagerClientObject, err = gi.ObjectNew("Gio", "DBusObjectManagerClient")
	})
	return err
}

type DBusObjectManagerClient struct {
	native unsafe.Pointer
}

func DBusObjectManagerClientNewFromNative(native unsafe.Pointer) *DBusObjectManagerClient {
	instance := &DBusObjectManagerClient{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *DBusObjectManagerClient) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToDBusObjectManagerClient down casts any arbitrary Object to DBusObjectManagerClient.
Exercise care, as this is a potentially dangerous function
if the Object is not a DBusObjectManagerClient.
*/
func CastToDBusObjectManagerClient(object *gobject.Object) *DBusObjectManagerClient {
	return DBusObjectManagerClientNewFromNative(object.Native())
}

// Equals compares this DBusObjectManagerClient with another DBusObjectManagerClient, and returns true if they represent the same GObject.
func (recv *DBusObjectManagerClient) Equals(other *DBusObjectManagerClient) bool {
	return other.Native() == recv.Native()
}

func (recv *DBusObjectManagerClient) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'g_dbus_object_manager_client_new_finish' : parameter 'res' of type 'AsyncResult' not supported

// UNSUPPORTED : C value 'g_dbus_object_manager_client_new_for_bus_finish' : parameter 'res' of type 'AsyncResult' not supported

// UNSUPPORTED : C value 'g_dbus_object_manager_client_new_for_bus_sync' : parameter 'flags' of type 'DBusObjectManagerClientFlags' not supported

// UNSUPPORTED : C value 'g_dbus_object_manager_client_new_sync' : parameter 'flags' of type 'DBusObjectManagerClientFlags' not supported

var dBusObjectManagerClientGetConnectionFunction *gi.Function
var dBusObjectManagerClientGetConnectionFunction_Once sync.Once

func dBusObjectManagerClientGetConnectionFunction_Set() error {
	var err error
	dBusObjectManagerClientGetConnectionFunction_Once.Do(func() {
		err = dBusObjectManagerClientObject_Set()
		if err != nil {
			return
		}
		dBusObjectManagerClientGetConnectionFunction, err = dBusObjectManagerClientObject.InvokerNew("get_connection")
	})
	return err
}

// GetConnection is a representation of the C type g_dbus_object_manager_client_get_connection.
func (recv *DBusObjectManagerClient) GetConnection() *DBusConnection {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := dBusObjectManagerClientGetConnectionFunction_Set()
	if err == nil {
		ret = dBusObjectManagerClientGetConnectionFunction.Invoke(inArgs[:], nil)
	}

	retGo := DBusConnectionNewFromNative(ret.Pointer())

	return retGo
}

// UNSUPPORTED : C value 'g_dbus_object_manager_client_get_flags' : return type 'DBusObjectManagerClientFlags' not supported

var dBusObjectManagerClientGetNameFunction *gi.Function
var dBusObjectManagerClientGetNameFunction_Once sync.Once

func dBusObjectManagerClientGetNameFunction_Set() error {
	var err error
	dBusObjectManagerClientGetNameFunction_Once.Do(func() {
		err = dBusObjectManagerClientObject_Set()
		if err != nil {
			return
		}
		dBusObjectManagerClientGetNameFunction, err = dBusObjectManagerClientObject.InvokerNew("get_name")
	})
	return err
}

// GetName is a representation of the C type g_dbus_object_manager_client_get_name.
func (recv *DBusObjectManagerClient) GetName() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := dBusObjectManagerClientGetNameFunction_Set()
	if err == nil {
		ret = dBusObjectManagerClientGetNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var dBusObjectManagerClientGetNameOwnerFunction *gi.Function
var dBusObjectManagerClientGetNameOwnerFunction_Once sync.Once

func dBusObjectManagerClientGetNameOwnerFunction_Set() error {
	var err error
	dBusObjectManagerClientGetNameOwnerFunction_Once.Do(func() {
		err = dBusObjectManagerClientObject_Set()
		if err != nil {
			return
		}
		dBusObjectManagerClientGetNameOwnerFunction, err = dBusObjectManagerClientObject.InvokerNew("get_name_owner")
	})
	return err
}

// GetNameOwner is a representation of the C type g_dbus_object_manager_client_get_name_owner.
func (recv *DBusObjectManagerClient) GetNameOwner() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := dBusObjectManagerClientGetNameOwnerFunction_Set()
	if err == nil {
		ret = dBusObjectManagerClientGetNameOwnerFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

var dBusObjectManagerServerObject *gi.Object
var dBusObjectManagerServerObject_Once sync.Once

func dBusObjectManagerServerObject_Set() error {
	var err error
	dBusObjectManagerServerObject_Once.Do(func() {
		dBusObjectManagerServerObject, err = gi.ObjectNew("Gio", "DBusObjectManagerServer")
	})
	return err
}

type DBusObjectManagerServer struct {
	native unsafe.Pointer
}

func DBusObjectManagerServerNewFromNative(native unsafe.Pointer) *DBusObjectManagerServer {
	instance := &DBusObjectManagerServer{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *DBusObjectManagerServer) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToDBusObjectManagerServer down casts any arbitrary Object to DBusObjectManagerServer.
Exercise care, as this is a potentially dangerous function
if the Object is not a DBusObjectManagerServer.
*/
func CastToDBusObjectManagerServer(object *gobject.Object) *DBusObjectManagerServer {
	return DBusObjectManagerServerNewFromNative(object.Native())
}

// Equals compares this DBusObjectManagerServer with another DBusObjectManagerServer, and returns true if they represent the same GObject.
func (recv *DBusObjectManagerServer) Equals(other *DBusObjectManagerServer) bool {
	return other.Native() == recv.Native()
}

func (recv *DBusObjectManagerServer) Native() unsafe.Pointer {
	return recv.native
}

var dBusObjectManagerServerNewFunction *gi.Function
var dBusObjectManagerServerNewFunction_Once sync.Once

func dBusObjectManagerServerNewFunction_Set() error {
	var err error
	dBusObjectManagerServerNewFunction_Once.Do(func() {
		err = dBusObjectManagerServerObject_Set()
		if err != nil {
			return
		}
		dBusObjectManagerServerNewFunction, err = dBusObjectManagerServerObject.InvokerNew("new")
	})
	return err
}

// DBusObjectManagerServerNew is a representation of the C type g_dbus_object_manager_server_new.
func DBusObjectManagerServerNew(objectPath string) *DBusObjectManagerServer {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(objectPath)

	var ret gi.Argument

	err := dBusObjectManagerServerNewFunction_Set()
	if err == nil {
		ret = dBusObjectManagerServerNewFunction.Invoke(inArgs[:], nil)
	}

	retGo := DBusObjectManagerServerNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

var dBusObjectManagerServerExportFunction *gi.Function
var dBusObjectManagerServerExportFunction_Once sync.Once

func dBusObjectManagerServerExportFunction_Set() error {
	var err error
	dBusObjectManagerServerExportFunction_Once.Do(func() {
		err = dBusObjectManagerServerObject_Set()
		if err != nil {
			return
		}
		dBusObjectManagerServerExportFunction, err = dBusObjectManagerServerObject.InvokerNew("export")
	})
	return err
}

// Export is a representation of the C type g_dbus_object_manager_server_export.
func (recv *DBusObjectManagerServer) Export(object_ *DBusObjectSkeleton) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(object_.Native())

	err := dBusObjectManagerServerExportFunction_Set()
	if err == nil {
		dBusObjectManagerServerExportFunction.Invoke(inArgs[:], nil)
	}

	return
}

var dBusObjectManagerServerExportUniquelyFunction *gi.Function
var dBusObjectManagerServerExportUniquelyFunction_Once sync.Once

func dBusObjectManagerServerExportUniquelyFunction_Set() error {
	var err error
	dBusObjectManagerServerExportUniquelyFunction_Once.Do(func() {
		err = dBusObjectManagerServerObject_Set()
		if err != nil {
			return
		}
		dBusObjectManagerServerExportUniquelyFunction, err = dBusObjectManagerServerObject.InvokerNew("export_uniquely")
	})
	return err
}

// ExportUniquely is a representation of the C type g_dbus_object_manager_server_export_uniquely.
func (recv *DBusObjectManagerServer) ExportUniquely(object_ *DBusObjectSkeleton) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(object_.Native())

	err := dBusObjectManagerServerExportUniquelyFunction_Set()
	if err == nil {
		dBusObjectManagerServerExportUniquelyFunction.Invoke(inArgs[:], nil)
	}

	return
}

var dBusObjectManagerServerGetConnectionFunction *gi.Function
var dBusObjectManagerServerGetConnectionFunction_Once sync.Once

func dBusObjectManagerServerGetConnectionFunction_Set() error {
	var err error
	dBusObjectManagerServerGetConnectionFunction_Once.Do(func() {
		err = dBusObjectManagerServerObject_Set()
		if err != nil {
			return
		}
		dBusObjectManagerServerGetConnectionFunction, err = dBusObjectManagerServerObject.InvokerNew("get_connection")
	})
	return err
}

// GetConnection is a representation of the C type g_dbus_object_manager_server_get_connection.
func (recv *DBusObjectManagerServer) GetConnection() *DBusConnection {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := dBusObjectManagerServerGetConnectionFunction_Set()
	if err == nil {
		ret = dBusObjectManagerServerGetConnectionFunction.Invoke(inArgs[:], nil)
	}

	retGo := DBusConnectionNewFromNative(ret.Pointer())

	return retGo
}

var dBusObjectManagerServerIsExportedFunction *gi.Function
var dBusObjectManagerServerIsExportedFunction_Once sync.Once

func dBusObjectManagerServerIsExportedFunction_Set() error {
	var err error
	dBusObjectManagerServerIsExportedFunction_Once.Do(func() {
		err = dBusObjectManagerServerObject_Set()
		if err != nil {
			return
		}
		dBusObjectManagerServerIsExportedFunction, err = dBusObjectManagerServerObject.InvokerNew("is_exported")
	})
	return err
}

// IsExported is a representation of the C type g_dbus_object_manager_server_is_exported.
func (recv *DBusObjectManagerServer) IsExported(object_ *DBusObjectSkeleton) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(object_.Native())

	var ret gi.Argument

	err := dBusObjectManagerServerIsExportedFunction_Set()
	if err == nil {
		ret = dBusObjectManagerServerIsExportedFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var dBusObjectManagerServerSetConnectionFunction *gi.Function
var dBusObjectManagerServerSetConnectionFunction_Once sync.Once

func dBusObjectManagerServerSetConnectionFunction_Set() error {
	var err error
	dBusObjectManagerServerSetConnectionFunction_Once.Do(func() {
		err = dBusObjectManagerServerObject_Set()
		if err != nil {
			return
		}
		dBusObjectManagerServerSetConnectionFunction, err = dBusObjectManagerServerObject.InvokerNew("set_connection")
	})
	return err
}

// SetConnection is a representation of the C type g_dbus_object_manager_server_set_connection.
func (recv *DBusObjectManagerServer) SetConnection(connection *DBusConnection) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(connection.Native())

	err := dBusObjectManagerServerSetConnectionFunction_Set()
	if err == nil {
		dBusObjectManagerServerSetConnectionFunction.Invoke(inArgs[:], nil)
	}

	return
}

var dBusObjectManagerServerUnexportFunction *gi.Function
var dBusObjectManagerServerUnexportFunction_Once sync.Once

func dBusObjectManagerServerUnexportFunction_Set() error {
	var err error
	dBusObjectManagerServerUnexportFunction_Once.Do(func() {
		err = dBusObjectManagerServerObject_Set()
		if err != nil {
			return
		}
		dBusObjectManagerServerUnexportFunction, err = dBusObjectManagerServerObject.InvokerNew("unexport")
	})
	return err
}

// Unexport is a representation of the C type g_dbus_object_manager_server_unexport.
func (recv *DBusObjectManagerServer) Unexport(objectPath string) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(objectPath)

	var ret gi.Argument

	err := dBusObjectManagerServerUnexportFunction_Set()
	if err == nil {
		ret = dBusObjectManagerServerUnexportFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var dBusObjectProxyObject *gi.Object
var dBusObjectProxyObject_Once sync.Once

func dBusObjectProxyObject_Set() error {
	var err error
	dBusObjectProxyObject_Once.Do(func() {
		dBusObjectProxyObject, err = gi.ObjectNew("Gio", "DBusObjectProxy")
	})
	return err
}

type DBusObjectProxy struct {
	native unsafe.Pointer
}

func DBusObjectProxyNewFromNative(native unsafe.Pointer) *DBusObjectProxy {
	instance := &DBusObjectProxy{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *DBusObjectProxy) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToDBusObjectProxy down casts any arbitrary Object to DBusObjectProxy.
Exercise care, as this is a potentially dangerous function
if the Object is not a DBusObjectProxy.
*/
func CastToDBusObjectProxy(object *gobject.Object) *DBusObjectProxy {
	return DBusObjectProxyNewFromNative(object.Native())
}

// Equals compares this DBusObjectProxy with another DBusObjectProxy, and returns true if they represent the same GObject.
func (recv *DBusObjectProxy) Equals(other *DBusObjectProxy) bool {
	return other.Native() == recv.Native()
}

func (recv *DBusObjectProxy) Native() unsafe.Pointer {
	return recv.native
}

var dBusObjectProxyNewFunction *gi.Function
var dBusObjectProxyNewFunction_Once sync.Once

func dBusObjectProxyNewFunction_Set() error {
	var err error
	dBusObjectProxyNewFunction_Once.Do(func() {
		err = dBusObjectProxyObject_Set()
		if err != nil {
			return
		}
		dBusObjectProxyNewFunction, err = dBusObjectProxyObject.InvokerNew("new")
	})
	return err
}

// DBusObjectProxyNew is a representation of the C type g_dbus_object_proxy_new.
func DBusObjectProxyNew(connection *DBusConnection, objectPath string) *DBusObjectProxy {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(connection.Native())
	inArgs[1].SetString(objectPath)

	var ret gi.Argument

	err := dBusObjectProxyNewFunction_Set()
	if err == nil {
		ret = dBusObjectProxyNewFunction.Invoke(inArgs[:], nil)
	}

	retGo := DBusObjectProxyNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

var dBusObjectProxyGetConnectionFunction *gi.Function
var dBusObjectProxyGetConnectionFunction_Once sync.Once

func dBusObjectProxyGetConnectionFunction_Set() error {
	var err error
	dBusObjectProxyGetConnectionFunction_Once.Do(func() {
		err = dBusObjectProxyObject_Set()
		if err != nil {
			return
		}
		dBusObjectProxyGetConnectionFunction, err = dBusObjectProxyObject.InvokerNew("get_connection")
	})
	return err
}

// GetConnection is a representation of the C type g_dbus_object_proxy_get_connection.
func (recv *DBusObjectProxy) GetConnection() *DBusConnection {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := dBusObjectProxyGetConnectionFunction_Set()
	if err == nil {
		ret = dBusObjectProxyGetConnectionFunction.Invoke(inArgs[:], nil)
	}

	retGo := DBusConnectionNewFromNative(ret.Pointer())

	return retGo
}

var dBusObjectSkeletonObject *gi.Object
var dBusObjectSkeletonObject_Once sync.Once

func dBusObjectSkeletonObject_Set() error {
	var err error
	dBusObjectSkeletonObject_Once.Do(func() {
		dBusObjectSkeletonObject, err = gi.ObjectNew("Gio", "DBusObjectSkeleton")
	})
	return err
}

type DBusObjectSkeleton struct {
	native unsafe.Pointer
}

func DBusObjectSkeletonNewFromNative(native unsafe.Pointer) *DBusObjectSkeleton {
	instance := &DBusObjectSkeleton{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *DBusObjectSkeleton) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToDBusObjectSkeleton down casts any arbitrary Object to DBusObjectSkeleton.
Exercise care, as this is a potentially dangerous function
if the Object is not a DBusObjectSkeleton.
*/
func CastToDBusObjectSkeleton(object *gobject.Object) *DBusObjectSkeleton {
	return DBusObjectSkeletonNewFromNative(object.Native())
}

// Equals compares this DBusObjectSkeleton with another DBusObjectSkeleton, and returns true if they represent the same GObject.
func (recv *DBusObjectSkeleton) Equals(other *DBusObjectSkeleton) bool {
	return other.Native() == recv.Native()
}

func (recv *DBusObjectSkeleton) Native() unsafe.Pointer {
	return recv.native
}

var dBusObjectSkeletonNewFunction *gi.Function
var dBusObjectSkeletonNewFunction_Once sync.Once

func dBusObjectSkeletonNewFunction_Set() error {
	var err error
	dBusObjectSkeletonNewFunction_Once.Do(func() {
		err = dBusObjectSkeletonObject_Set()
		if err != nil {
			return
		}
		dBusObjectSkeletonNewFunction, err = dBusObjectSkeletonObject.InvokerNew("new")
	})
	return err
}

// DBusObjectSkeletonNew is a representation of the C type g_dbus_object_skeleton_new.
func DBusObjectSkeletonNew(objectPath string) *DBusObjectSkeleton {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(objectPath)

	var ret gi.Argument

	err := dBusObjectSkeletonNewFunction_Set()
	if err == nil {
		ret = dBusObjectSkeletonNewFunction.Invoke(inArgs[:], nil)
	}

	retGo := DBusObjectSkeletonNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

var dBusObjectSkeletonAddInterfaceFunction *gi.Function
var dBusObjectSkeletonAddInterfaceFunction_Once sync.Once

func dBusObjectSkeletonAddInterfaceFunction_Set() error {
	var err error
	dBusObjectSkeletonAddInterfaceFunction_Once.Do(func() {
		err = dBusObjectSkeletonObject_Set()
		if err != nil {
			return
		}
		dBusObjectSkeletonAddInterfaceFunction, err = dBusObjectSkeletonObject.InvokerNew("add_interface")
	})
	return err
}

// AddInterface is a representation of the C type g_dbus_object_skeleton_add_interface.
func (recv *DBusObjectSkeleton) AddInterface(interface_ *DBusInterfaceSkeleton) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(interface_.Native())

	err := dBusObjectSkeletonAddInterfaceFunction_Set()
	if err == nil {
		dBusObjectSkeletonAddInterfaceFunction.Invoke(inArgs[:], nil)
	}

	return
}

var dBusObjectSkeletonFlushFunction *gi.Function
var dBusObjectSkeletonFlushFunction_Once sync.Once

func dBusObjectSkeletonFlushFunction_Set() error {
	var err error
	dBusObjectSkeletonFlushFunction_Once.Do(func() {
		err = dBusObjectSkeletonObject_Set()
		if err != nil {
			return
		}
		dBusObjectSkeletonFlushFunction, err = dBusObjectSkeletonObject.InvokerNew("flush")
	})
	return err
}

// Flush is a representation of the C type g_dbus_object_skeleton_flush.
func (recv *DBusObjectSkeleton) Flush() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := dBusObjectSkeletonFlushFunction_Set()
	if err == nil {
		dBusObjectSkeletonFlushFunction.Invoke(inArgs[:], nil)
	}

	return
}

var dBusObjectSkeletonRemoveInterfaceFunction *gi.Function
var dBusObjectSkeletonRemoveInterfaceFunction_Once sync.Once

func dBusObjectSkeletonRemoveInterfaceFunction_Set() error {
	var err error
	dBusObjectSkeletonRemoveInterfaceFunction_Once.Do(func() {
		err = dBusObjectSkeletonObject_Set()
		if err != nil {
			return
		}
		dBusObjectSkeletonRemoveInterfaceFunction, err = dBusObjectSkeletonObject.InvokerNew("remove_interface")
	})
	return err
}

// RemoveInterface is a representation of the C type g_dbus_object_skeleton_remove_interface.
func (recv *DBusObjectSkeleton) RemoveInterface(interface_ *DBusInterfaceSkeleton) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(interface_.Native())

	err := dBusObjectSkeletonRemoveInterfaceFunction_Set()
	if err == nil {
		dBusObjectSkeletonRemoveInterfaceFunction.Invoke(inArgs[:], nil)
	}

	return
}

var dBusObjectSkeletonRemoveInterfaceByNameFunction *gi.Function
var dBusObjectSkeletonRemoveInterfaceByNameFunction_Once sync.Once

func dBusObjectSkeletonRemoveInterfaceByNameFunction_Set() error {
	var err error
	dBusObjectSkeletonRemoveInterfaceByNameFunction_Once.Do(func() {
		err = dBusObjectSkeletonObject_Set()
		if err != nil {
			return
		}
		dBusObjectSkeletonRemoveInterfaceByNameFunction, err = dBusObjectSkeletonObject.InvokerNew("remove_interface_by_name")
	})
	return err
}

// RemoveInterfaceByName is a representation of the C type g_dbus_object_skeleton_remove_interface_by_name.
func (recv *DBusObjectSkeleton) RemoveInterfaceByName(interfaceName string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(interfaceName)

	err := dBusObjectSkeletonRemoveInterfaceByNameFunction_Set()
	if err == nil {
		dBusObjectSkeletonRemoveInterfaceByNameFunction.Invoke(inArgs[:], nil)
	}

	return
}

var dBusObjectSkeletonSetObjectPathFunction *gi.Function
var dBusObjectSkeletonSetObjectPathFunction_Once sync.Once

func dBusObjectSkeletonSetObjectPathFunction_Set() error {
	var err error
	dBusObjectSkeletonSetObjectPathFunction_Once.Do(func() {
		err = dBusObjectSkeletonObject_Set()
		if err != nil {
			return
		}
		dBusObjectSkeletonSetObjectPathFunction, err = dBusObjectSkeletonObject.InvokerNew("set_object_path")
	})
	return err
}

// SetObjectPath is a representation of the C type g_dbus_object_skeleton_set_object_path.
func (recv *DBusObjectSkeleton) SetObjectPath(objectPath string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(objectPath)

	err := dBusObjectSkeletonSetObjectPathFunction_Set()
	if err == nil {
		dBusObjectSkeletonSetObjectPathFunction.Invoke(inArgs[:], nil)
	}

	return
}

var dBusProxyObject *gi.Object
var dBusProxyObject_Once sync.Once

func dBusProxyObject_Set() error {
	var err error
	dBusProxyObject_Once.Do(func() {
		dBusProxyObject, err = gi.ObjectNew("Gio", "DBusProxy")
	})
	return err
}

type DBusProxy struct {
	native unsafe.Pointer
}

func DBusProxyNewFromNative(native unsafe.Pointer) *DBusProxy {
	instance := &DBusProxy{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *DBusProxy) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToDBusProxy down casts any arbitrary Object to DBusProxy.
Exercise care, as this is a potentially dangerous function
if the Object is not a DBusProxy.
*/
func CastToDBusProxy(object *gobject.Object) *DBusProxy {
	return DBusProxyNewFromNative(object.Native())
}

// Equals compares this DBusProxy with another DBusProxy, and returns true if they represent the same GObject.
func (recv *DBusProxy) Equals(other *DBusProxy) bool {
	return other.Native() == recv.Native()
}

func (recv *DBusProxy) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'g_dbus_proxy_new_finish' : parameter 'res' of type 'AsyncResult' not supported

// UNSUPPORTED : C value 'g_dbus_proxy_new_for_bus_finish' : parameter 'res' of type 'AsyncResult' not supported

// UNSUPPORTED : C value 'g_dbus_proxy_new_for_bus_sync' : parameter 'flags' of type 'DBusProxyFlags' not supported

// UNSUPPORTED : C value 'g_dbus_proxy_new_sync' : parameter 'flags' of type 'DBusProxyFlags' not supported

// UNSUPPORTED : C value 'g_dbus_proxy_call' : parameter 'parameters' of type 'GLib.Variant' not supported

// UNSUPPORTED : C value 'g_dbus_proxy_call_finish' : parameter 'res' of type 'AsyncResult' not supported

// UNSUPPORTED : C value 'g_dbus_proxy_call_sync' : parameter 'parameters' of type 'GLib.Variant' not supported

// UNSUPPORTED : C value 'g_dbus_proxy_call_with_unix_fd_list' : parameter 'parameters' of type 'GLib.Variant' not supported

// UNSUPPORTED : C value 'g_dbus_proxy_call_with_unix_fd_list_finish' : parameter 'res' of type 'AsyncResult' not supported

// UNSUPPORTED : C value 'g_dbus_proxy_call_with_unix_fd_list_sync' : parameter 'parameters' of type 'GLib.Variant' not supported

// UNSUPPORTED : C value 'g_dbus_proxy_get_cached_property' : return type 'GLib.Variant' not supported

var dBusProxyGetCachedPropertyNamesFunction *gi.Function
var dBusProxyGetCachedPropertyNamesFunction_Once sync.Once

func dBusProxyGetCachedPropertyNamesFunction_Set() error {
	var err error
	dBusProxyGetCachedPropertyNamesFunction_Once.Do(func() {
		err = dBusProxyObject_Set()
		if err != nil {
			return
		}
		dBusProxyGetCachedPropertyNamesFunction, err = dBusProxyObject.InvokerNew("get_cached_property_names")
	})
	return err
}

// GetCachedPropertyNames is a representation of the C type g_dbus_proxy_get_cached_property_names.
func (recv *DBusProxy) GetCachedPropertyNames() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := dBusProxyGetCachedPropertyNamesFunction_Set()
	if err == nil {
		dBusProxyGetCachedPropertyNamesFunction.Invoke(inArgs[:], nil)
	}

	return
}

var dBusProxyGetConnectionFunction *gi.Function
var dBusProxyGetConnectionFunction_Once sync.Once

func dBusProxyGetConnectionFunction_Set() error {
	var err error
	dBusProxyGetConnectionFunction_Once.Do(func() {
		err = dBusProxyObject_Set()
		if err != nil {
			return
		}
		dBusProxyGetConnectionFunction, err = dBusProxyObject.InvokerNew("get_connection")
	})
	return err
}

// GetConnection is a representation of the C type g_dbus_proxy_get_connection.
func (recv *DBusProxy) GetConnection() *DBusConnection {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := dBusProxyGetConnectionFunction_Set()
	if err == nil {
		ret = dBusProxyGetConnectionFunction.Invoke(inArgs[:], nil)
	}

	retGo := DBusConnectionNewFromNative(ret.Pointer())

	return retGo
}

var dBusProxyGetDefaultTimeoutFunction *gi.Function
var dBusProxyGetDefaultTimeoutFunction_Once sync.Once

func dBusProxyGetDefaultTimeoutFunction_Set() error {
	var err error
	dBusProxyGetDefaultTimeoutFunction_Once.Do(func() {
		err = dBusProxyObject_Set()
		if err != nil {
			return
		}
		dBusProxyGetDefaultTimeoutFunction, err = dBusProxyObject.InvokerNew("get_default_timeout")
	})
	return err
}

// GetDefaultTimeout is a representation of the C type g_dbus_proxy_get_default_timeout.
func (recv *DBusProxy) GetDefaultTimeout() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := dBusProxyGetDefaultTimeoutFunction_Set()
	if err == nil {
		ret = dBusProxyGetDefaultTimeoutFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

// UNSUPPORTED : C value 'g_dbus_proxy_get_flags' : return type 'DBusProxyFlags' not supported

var dBusProxyGetInterfaceInfoFunction *gi.Function
var dBusProxyGetInterfaceInfoFunction_Once sync.Once

func dBusProxyGetInterfaceInfoFunction_Set() error {
	var err error
	dBusProxyGetInterfaceInfoFunction_Once.Do(func() {
		err = dBusProxyObject_Set()
		if err != nil {
			return
		}
		dBusProxyGetInterfaceInfoFunction, err = dBusProxyObject.InvokerNew("get_interface_info")
	})
	return err
}

// GetInterfaceInfo is a representation of the C type g_dbus_proxy_get_interface_info.
func (recv *DBusProxy) GetInterfaceInfo() *DBusInterfaceInfo {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := dBusProxyGetInterfaceInfoFunction_Set()
	if err == nil {
		ret = dBusProxyGetInterfaceInfoFunction.Invoke(inArgs[:], nil)
	}

	retGo := DBusInterfaceInfoNewFromNative(ret.Pointer())

	return retGo
}

var dBusProxyGetInterfaceNameFunction *gi.Function
var dBusProxyGetInterfaceNameFunction_Once sync.Once

func dBusProxyGetInterfaceNameFunction_Set() error {
	var err error
	dBusProxyGetInterfaceNameFunction_Once.Do(func() {
		err = dBusProxyObject_Set()
		if err != nil {
			return
		}
		dBusProxyGetInterfaceNameFunction, err = dBusProxyObject.InvokerNew("get_interface_name")
	})
	return err
}

// GetInterfaceName is a representation of the C type g_dbus_proxy_get_interface_name.
func (recv *DBusProxy) GetInterfaceName() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := dBusProxyGetInterfaceNameFunction_Set()
	if err == nil {
		ret = dBusProxyGetInterfaceNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var dBusProxyGetNameFunction *gi.Function
var dBusProxyGetNameFunction_Once sync.Once

func dBusProxyGetNameFunction_Set() error {
	var err error
	dBusProxyGetNameFunction_Once.Do(func() {
		err = dBusProxyObject_Set()
		if err != nil {
			return
		}
		dBusProxyGetNameFunction, err = dBusProxyObject.InvokerNew("get_name")
	})
	return err
}

// GetName is a representation of the C type g_dbus_proxy_get_name.
func (recv *DBusProxy) GetName() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := dBusProxyGetNameFunction_Set()
	if err == nil {
		ret = dBusProxyGetNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var dBusProxyGetNameOwnerFunction *gi.Function
var dBusProxyGetNameOwnerFunction_Once sync.Once

func dBusProxyGetNameOwnerFunction_Set() error {
	var err error
	dBusProxyGetNameOwnerFunction_Once.Do(func() {
		err = dBusProxyObject_Set()
		if err != nil {
			return
		}
		dBusProxyGetNameOwnerFunction, err = dBusProxyObject.InvokerNew("get_name_owner")
	})
	return err
}

// GetNameOwner is a representation of the C type g_dbus_proxy_get_name_owner.
func (recv *DBusProxy) GetNameOwner() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := dBusProxyGetNameOwnerFunction_Set()
	if err == nil {
		ret = dBusProxyGetNameOwnerFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

var dBusProxyGetObjectPathFunction *gi.Function
var dBusProxyGetObjectPathFunction_Once sync.Once

func dBusProxyGetObjectPathFunction_Set() error {
	var err error
	dBusProxyGetObjectPathFunction_Once.Do(func() {
		err = dBusProxyObject_Set()
		if err != nil {
			return
		}
		dBusProxyGetObjectPathFunction, err = dBusProxyObject.InvokerNew("get_object_path")
	})
	return err
}

// GetObjectPath is a representation of the C type g_dbus_proxy_get_object_path.
func (recv *DBusProxy) GetObjectPath() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := dBusProxyGetObjectPathFunction_Set()
	if err == nil {
		ret = dBusProxyGetObjectPathFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'g_dbus_proxy_set_cached_property' : parameter 'value' of type 'GLib.Variant' not supported

var dBusProxySetDefaultTimeoutFunction *gi.Function
var dBusProxySetDefaultTimeoutFunction_Once sync.Once

func dBusProxySetDefaultTimeoutFunction_Set() error {
	var err error
	dBusProxySetDefaultTimeoutFunction_Once.Do(func() {
		err = dBusProxyObject_Set()
		if err != nil {
			return
		}
		dBusProxySetDefaultTimeoutFunction, err = dBusProxyObject.InvokerNew("set_default_timeout")
	})
	return err
}

// SetDefaultTimeout is a representation of the C type g_dbus_proxy_set_default_timeout.
func (recv *DBusProxy) SetDefaultTimeout(timeoutMsec int32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(timeoutMsec)

	err := dBusProxySetDefaultTimeoutFunction_Set()
	if err == nil {
		dBusProxySetDefaultTimeoutFunction.Invoke(inArgs[:], nil)
	}

	return
}

var dBusProxySetInterfaceInfoFunction *gi.Function
var dBusProxySetInterfaceInfoFunction_Once sync.Once

func dBusProxySetInterfaceInfoFunction_Set() error {
	var err error
	dBusProxySetInterfaceInfoFunction_Once.Do(func() {
		err = dBusProxyObject_Set()
		if err != nil {
			return
		}
		dBusProxySetInterfaceInfoFunction, err = dBusProxyObject.InvokerNew("set_interface_info")
	})
	return err
}

// SetInterfaceInfo is a representation of the C type g_dbus_proxy_set_interface_info.
func (recv *DBusProxy) SetInterfaceInfo(info *DBusInterfaceInfo) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(info.Native())

	err := dBusProxySetInterfaceInfoFunction_Set()
	if err == nil {
		dBusProxySetInterfaceInfoFunction.Invoke(inArgs[:], nil)
	}

	return
}

var dBusServerObject *gi.Object
var dBusServerObject_Once sync.Once

func dBusServerObject_Set() error {
	var err error
	dBusServerObject_Once.Do(func() {
		dBusServerObject, err = gi.ObjectNew("Gio", "DBusServer")
	})
	return err
}

type DBusServer struct {
	native unsafe.Pointer
}

func DBusServerNewFromNative(native unsafe.Pointer) *DBusServer {
	instance := &DBusServer{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *DBusServer) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToDBusServer down casts any arbitrary Object to DBusServer.
Exercise care, as this is a potentially dangerous function
if the Object is not a DBusServer.
*/
func CastToDBusServer(object *gobject.Object) *DBusServer {
	return DBusServerNewFromNative(object.Native())
}

// Equals compares this DBusServer with another DBusServer, and returns true if they represent the same GObject.
func (recv *DBusServer) Equals(other *DBusServer) bool {
	return other.Native() == recv.Native()
}

func (recv *DBusServer) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'g_dbus_server_new_sync' : parameter 'flags' of type 'DBusServerFlags' not supported

var dBusServerGetClientAddressFunction *gi.Function
var dBusServerGetClientAddressFunction_Once sync.Once

func dBusServerGetClientAddressFunction_Set() error {
	var err error
	dBusServerGetClientAddressFunction_Once.Do(func() {
		err = dBusServerObject_Set()
		if err != nil {
			return
		}
		dBusServerGetClientAddressFunction, err = dBusServerObject.InvokerNew("get_client_address")
	})
	return err
}

// GetClientAddress is a representation of the C type g_dbus_server_get_client_address.
func (recv *DBusServer) GetClientAddress() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := dBusServerGetClientAddressFunction_Set()
	if err == nil {
		ret = dBusServerGetClientAddressFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'g_dbus_server_get_flags' : return type 'DBusServerFlags' not supported

var dBusServerGetGuidFunction *gi.Function
var dBusServerGetGuidFunction_Once sync.Once

func dBusServerGetGuidFunction_Set() error {
	var err error
	dBusServerGetGuidFunction_Once.Do(func() {
		err = dBusServerObject_Set()
		if err != nil {
			return
		}
		dBusServerGetGuidFunction, err = dBusServerObject.InvokerNew("get_guid")
	})
	return err
}

// GetGuid is a representation of the C type g_dbus_server_get_guid.
func (recv *DBusServer) GetGuid() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := dBusServerGetGuidFunction_Set()
	if err == nil {
		ret = dBusServerGetGuidFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var dBusServerIsActiveFunction *gi.Function
var dBusServerIsActiveFunction_Once sync.Once

func dBusServerIsActiveFunction_Set() error {
	var err error
	dBusServerIsActiveFunction_Once.Do(func() {
		err = dBusServerObject_Set()
		if err != nil {
			return
		}
		dBusServerIsActiveFunction, err = dBusServerObject.InvokerNew("is_active")
	})
	return err
}

// IsActive is a representation of the C type g_dbus_server_is_active.
func (recv *DBusServer) IsActive() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := dBusServerIsActiveFunction_Set()
	if err == nil {
		ret = dBusServerIsActiveFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var dBusServerStartFunction *gi.Function
var dBusServerStartFunction_Once sync.Once

func dBusServerStartFunction_Set() error {
	var err error
	dBusServerStartFunction_Once.Do(func() {
		err = dBusServerObject_Set()
		if err != nil {
			return
		}
		dBusServerStartFunction, err = dBusServerObject.InvokerNew("start")
	})
	return err
}

// Start is a representation of the C type g_dbus_server_start.
func (recv *DBusServer) Start() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := dBusServerStartFunction_Set()
	if err == nil {
		dBusServerStartFunction.Invoke(inArgs[:], nil)
	}

	return
}

var dBusServerStopFunction *gi.Function
var dBusServerStopFunction_Once sync.Once

func dBusServerStopFunction_Set() error {
	var err error
	dBusServerStopFunction_Once.Do(func() {
		err = dBusServerObject_Set()
		if err != nil {
			return
		}
		dBusServerStopFunction, err = dBusServerObject.InvokerNew("stop")
	})
	return err
}

// Stop is a representation of the C type g_dbus_server_stop.
func (recv *DBusServer) Stop() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := dBusServerStopFunction_Set()
	if err == nil {
		dBusServerStopFunction.Invoke(inArgs[:], nil)
	}

	return
}

var dataInputStreamObject *gi.Object
var dataInputStreamObject_Once sync.Once

func dataInputStreamObject_Set() error {
	var err error
	dataInputStreamObject_Once.Do(func() {
		dataInputStreamObject, err = gi.ObjectNew("Gio", "DataInputStream")
	})
	return err
}

type DataInputStream struct {
	native unsafe.Pointer
}

func DataInputStreamNewFromNative(native unsafe.Pointer) *DataInputStream {
	instance := &DataInputStream{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// BufferedInputStream upcasts to *BufferedInputStream
func (recv *DataInputStream) BufferedInputStream() *BufferedInputStream {
	return BufferedInputStreamNewFromNative(recv.Native())
}

// FilterInputStream upcasts to *FilterInputStream
func (recv *DataInputStream) FilterInputStream() *FilterInputStream {
	return FilterInputStreamNewFromNative(recv.Native())
}

// InputStream upcasts to *InputStream
func (recv *DataInputStream) InputStream() *InputStream {
	return InputStreamNewFromNative(recv.Native())
}

// Object upcasts to *Object
func (recv *DataInputStream) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToDataInputStream down casts any arbitrary Object to DataInputStream.
Exercise care, as this is a potentially dangerous function
if the Object is not a DataInputStream.
*/
func CastToDataInputStream(object *gobject.Object) *DataInputStream {
	return DataInputStreamNewFromNative(object.Native())
}

// Equals compares this DataInputStream with another DataInputStream, and returns true if they represent the same GObject.
func (recv *DataInputStream) Equals(other *DataInputStream) bool {
	return other.Native() == recv.Native()
}

func (recv *DataInputStream) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *DataInputStream) FieldParentInstance() *BufferedInputStream {
	argValue := gi.ObjectFieldGet(dataInputStreamObject, recv.Native(), "parent_instance")
	value := BufferedInputStreamNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *DataInputStream) SetFieldParentInstance(value *BufferedInputStream) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(dataInputStreamObject, recv.Native(), "parent_instance", argValue)
}

var dataInputStreamNewFunction *gi.Function
var dataInputStreamNewFunction_Once sync.Once

func dataInputStreamNewFunction_Set() error {
	var err error
	dataInputStreamNewFunction_Once.Do(func() {
		err = dataInputStreamObject_Set()
		if err != nil {
			return
		}
		dataInputStreamNewFunction, err = dataInputStreamObject.InvokerNew("new")
	})
	return err
}

// DataInputStreamNew is a representation of the C type g_data_input_stream_new.
func DataInputStreamNew(baseStream *InputStream) *DataInputStream {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(baseStream.Native())

	var ret gi.Argument

	err := dataInputStreamNewFunction_Set()
	if err == nil {
		ret = dataInputStreamNewFunction.Invoke(inArgs[:], nil)
	}

	retGo := DataInputStreamNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

var dataInputStreamGetByteOrderFunction *gi.Function
var dataInputStreamGetByteOrderFunction_Once sync.Once

func dataInputStreamGetByteOrderFunction_Set() error {
	var err error
	dataInputStreamGetByteOrderFunction_Once.Do(func() {
		err = dataInputStreamObject_Set()
		if err != nil {
			return
		}
		dataInputStreamGetByteOrderFunction, err = dataInputStreamObject.InvokerNew("get_byte_order")
	})
	return err
}

// GetByteOrder is a representation of the C type g_data_input_stream_get_byte_order.
func (recv *DataInputStream) GetByteOrder() DataStreamByteOrder {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := dataInputStreamGetByteOrderFunction_Set()
	if err == nil {
		ret = dataInputStreamGetByteOrderFunction.Invoke(inArgs[:], nil)
	}

	retGo := DataStreamByteOrder(ret.Int32())

	return retGo
}

var dataInputStreamGetNewlineTypeFunction *gi.Function
var dataInputStreamGetNewlineTypeFunction_Once sync.Once

func dataInputStreamGetNewlineTypeFunction_Set() error {
	var err error
	dataInputStreamGetNewlineTypeFunction_Once.Do(func() {
		err = dataInputStreamObject_Set()
		if err != nil {
			return
		}
		dataInputStreamGetNewlineTypeFunction, err = dataInputStreamObject.InvokerNew("get_newline_type")
	})
	return err
}

// GetNewlineType is a representation of the C type g_data_input_stream_get_newline_type.
func (recv *DataInputStream) GetNewlineType() DataStreamNewlineType {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := dataInputStreamGetNewlineTypeFunction_Set()
	if err == nil {
		ret = dataInputStreamGetNewlineTypeFunction.Invoke(inArgs[:], nil)
	}

	retGo := DataStreamNewlineType(ret.Int32())

	return retGo
}

var dataInputStreamReadByteFunction *gi.Function
var dataInputStreamReadByteFunction_Once sync.Once

func dataInputStreamReadByteFunction_Set() error {
	var err error
	dataInputStreamReadByteFunction_Once.Do(func() {
		err = dataInputStreamObject_Set()
		if err != nil {
			return
		}
		dataInputStreamReadByteFunction, err = dataInputStreamObject.InvokerNew("read_byte")
	})
	return err
}

// ReadByte is a representation of the C type g_data_input_stream_read_byte.
func (recv *DataInputStream) ReadByte(cancellable *Cancellable) uint8 {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(cancellable.Native())

	var ret gi.Argument

	err := dataInputStreamReadByteFunction_Set()
	if err == nil {
		ret = dataInputStreamReadByteFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint8()

	return retGo
}

var dataInputStreamReadInt16Function *gi.Function
var dataInputStreamReadInt16Function_Once sync.Once

func dataInputStreamReadInt16Function_Set() error {
	var err error
	dataInputStreamReadInt16Function_Once.Do(func() {
		err = dataInputStreamObject_Set()
		if err != nil {
			return
		}
		dataInputStreamReadInt16Function, err = dataInputStreamObject.InvokerNew("read_int16")
	})
	return err
}

// ReadInt16 is a representation of the C type g_data_input_stream_read_int16.
func (recv *DataInputStream) ReadInt16(cancellable *Cancellable) int16 {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(cancellable.Native())

	var ret gi.Argument

	err := dataInputStreamReadInt16Function_Set()
	if err == nil {
		ret = dataInputStreamReadInt16Function.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int16()

	return retGo
}

var dataInputStreamReadInt32Function *gi.Function
var dataInputStreamReadInt32Function_Once sync.Once

func dataInputStreamReadInt32Function_Set() error {
	var err error
	dataInputStreamReadInt32Function_Once.Do(func() {
		err = dataInputStreamObject_Set()
		if err != nil {
			return
		}
		dataInputStreamReadInt32Function, err = dataInputStreamObject.InvokerNew("read_int32")
	})
	return err
}

// ReadInt32 is a representation of the C type g_data_input_stream_read_int32.
func (recv *DataInputStream) ReadInt32(cancellable *Cancellable) int32 {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(cancellable.Native())

	var ret gi.Argument

	err := dataInputStreamReadInt32Function_Set()
	if err == nil {
		ret = dataInputStreamReadInt32Function.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var dataInputStreamReadInt64Function *gi.Function
var dataInputStreamReadInt64Function_Once sync.Once

func dataInputStreamReadInt64Function_Set() error {
	var err error
	dataInputStreamReadInt64Function_Once.Do(func() {
		err = dataInputStreamObject_Set()
		if err != nil {
			return
		}
		dataInputStreamReadInt64Function, err = dataInputStreamObject.InvokerNew("read_int64")
	})
	return err
}

// ReadInt64 is a representation of the C type g_data_input_stream_read_int64.
func (recv *DataInputStream) ReadInt64(cancellable *Cancellable) int64 {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(cancellable.Native())

	var ret gi.Argument

	err := dataInputStreamReadInt64Function_Set()
	if err == nil {
		ret = dataInputStreamReadInt64Function.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int64()

	return retGo
}

var dataInputStreamReadLineFunction *gi.Function
var dataInputStreamReadLineFunction_Once sync.Once

func dataInputStreamReadLineFunction_Set() error {
	var err error
	dataInputStreamReadLineFunction_Once.Do(func() {
		err = dataInputStreamObject_Set()
		if err != nil {
			return
		}
		dataInputStreamReadLineFunction, err = dataInputStreamObject.InvokerNew("read_line")
	})
	return err
}

// ReadLine is a representation of the C type g_data_input_stream_read_line.
func (recv *DataInputStream) ReadLine(cancellable *Cancellable) uint64 {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(cancellable.Native())

	var outArgs [1]gi.Argument

	err := dataInputStreamReadLineFunction_Set()
	if err == nil {
		dataInputStreamReadLineFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := outArgs[0].Uint64()

	return out0
}

// UNSUPPORTED : C value 'g_data_input_stream_read_line_async' : parameter 'callback' of type 'AsyncReadyCallback' not supported

// UNSUPPORTED : C value 'g_data_input_stream_read_line_finish' : parameter 'result' of type 'AsyncResult' not supported

// UNSUPPORTED : C value 'g_data_input_stream_read_line_finish_utf8' : parameter 'result' of type 'AsyncResult' not supported

var dataInputStreamReadLineUtf8Function *gi.Function
var dataInputStreamReadLineUtf8Function_Once sync.Once

func dataInputStreamReadLineUtf8Function_Set() error {
	var err error
	dataInputStreamReadLineUtf8Function_Once.Do(func() {
		err = dataInputStreamObject_Set()
		if err != nil {
			return
		}
		dataInputStreamReadLineUtf8Function, err = dataInputStreamObject.InvokerNew("read_line_utf8")
	})
	return err
}

// ReadLineUtf8 is a representation of the C type g_data_input_stream_read_line_utf8.
func (recv *DataInputStream) ReadLineUtf8(cancellable *Cancellable) (string, uint64) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(cancellable.Native())

	var outArgs [1]gi.Argument
	var ret gi.Argument

	err := dataInputStreamReadLineUtf8Function_Set()
	if err == nil {
		ret = dataInputStreamReadLineUtf8Function.Invoke(inArgs[:], outArgs[:])
	}

	retGo := ret.String(true)
	out0 := outArgs[0].Uint64()

	return retGo, out0
}

var dataInputStreamReadUint16Function *gi.Function
var dataInputStreamReadUint16Function_Once sync.Once

func dataInputStreamReadUint16Function_Set() error {
	var err error
	dataInputStreamReadUint16Function_Once.Do(func() {
		err = dataInputStreamObject_Set()
		if err != nil {
			return
		}
		dataInputStreamReadUint16Function, err = dataInputStreamObject.InvokerNew("read_uint16")
	})
	return err
}

// ReadUint16 is a representation of the C type g_data_input_stream_read_uint16.
func (recv *DataInputStream) ReadUint16(cancellable *Cancellable) uint16 {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(cancellable.Native())

	var ret gi.Argument

	err := dataInputStreamReadUint16Function_Set()
	if err == nil {
		ret = dataInputStreamReadUint16Function.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint16()

	return retGo
}

var dataInputStreamReadUint32Function *gi.Function
var dataInputStreamReadUint32Function_Once sync.Once

func dataInputStreamReadUint32Function_Set() error {
	var err error
	dataInputStreamReadUint32Function_Once.Do(func() {
		err = dataInputStreamObject_Set()
		if err != nil {
			return
		}
		dataInputStreamReadUint32Function, err = dataInputStreamObject.InvokerNew("read_uint32")
	})
	return err
}

// ReadUint32 is a representation of the C type g_data_input_stream_read_uint32.
func (recv *DataInputStream) ReadUint32(cancellable *Cancellable) uint32 {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(cancellable.Native())

	var ret gi.Argument

	err := dataInputStreamReadUint32Function_Set()
	if err == nil {
		ret = dataInputStreamReadUint32Function.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo
}

var dataInputStreamReadUint64Function *gi.Function
var dataInputStreamReadUint64Function_Once sync.Once

func dataInputStreamReadUint64Function_Set() error {
	var err error
	dataInputStreamReadUint64Function_Once.Do(func() {
		err = dataInputStreamObject_Set()
		if err != nil {
			return
		}
		dataInputStreamReadUint64Function, err = dataInputStreamObject.InvokerNew("read_uint64")
	})
	return err
}

// ReadUint64 is a representation of the C type g_data_input_stream_read_uint64.
func (recv *DataInputStream) ReadUint64(cancellable *Cancellable) uint64 {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(cancellable.Native())

	var ret gi.Argument

	err := dataInputStreamReadUint64Function_Set()
	if err == nil {
		ret = dataInputStreamReadUint64Function.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint64()

	return retGo
}

var dataInputStreamReadUntilFunction *gi.Function
var dataInputStreamReadUntilFunction_Once sync.Once

func dataInputStreamReadUntilFunction_Set() error {
	var err error
	dataInputStreamReadUntilFunction_Once.Do(func() {
		err = dataInputStreamObject_Set()
		if err != nil {
			return
		}
		dataInputStreamReadUntilFunction, err = dataInputStreamObject.InvokerNew("read_until")
	})
	return err
}

// ReadUntil is a representation of the C type g_data_input_stream_read_until.
func (recv *DataInputStream) ReadUntil(stopChars string, cancellable *Cancellable) (string, uint64) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(stopChars)
	inArgs[2].SetPointer(cancellable.Native())

	var outArgs [1]gi.Argument
	var ret gi.Argument

	err := dataInputStreamReadUntilFunction_Set()
	if err == nil {
		ret = dataInputStreamReadUntilFunction.Invoke(inArgs[:], outArgs[:])
	}

	retGo := ret.String(true)
	out0 := outArgs[0].Uint64()

	return retGo, out0
}

// UNSUPPORTED : C value 'g_data_input_stream_read_until_async' : parameter 'callback' of type 'AsyncReadyCallback' not supported

// UNSUPPORTED : C value 'g_data_input_stream_read_until_finish' : parameter 'result' of type 'AsyncResult' not supported

var dataInputStreamReadUptoFunction *gi.Function
var dataInputStreamReadUptoFunction_Once sync.Once

func dataInputStreamReadUptoFunction_Set() error {
	var err error
	dataInputStreamReadUptoFunction_Once.Do(func() {
		err = dataInputStreamObject_Set()
		if err != nil {
			return
		}
		dataInputStreamReadUptoFunction, err = dataInputStreamObject.InvokerNew("read_upto")
	})
	return err
}

// ReadUpto is a representation of the C type g_data_input_stream_read_upto.
func (recv *DataInputStream) ReadUpto(stopChars string, stopCharsLen int32, cancellable *Cancellable) (string, uint64) {
	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(stopChars)
	inArgs[2].SetInt32(stopCharsLen)
	inArgs[3].SetPointer(cancellable.Native())

	var outArgs [1]gi.Argument
	var ret gi.Argument

	err := dataInputStreamReadUptoFunction_Set()
	if err == nil {
		ret = dataInputStreamReadUptoFunction.Invoke(inArgs[:], outArgs[:])
	}

	retGo := ret.String(true)
	out0 := outArgs[0].Uint64()

	return retGo, out0
}

// UNSUPPORTED : C value 'g_data_input_stream_read_upto_async' : parameter 'callback' of type 'AsyncReadyCallback' not supported

// UNSUPPORTED : C value 'g_data_input_stream_read_upto_finish' : parameter 'result' of type 'AsyncResult' not supported

var dataInputStreamSetByteOrderFunction *gi.Function
var dataInputStreamSetByteOrderFunction_Once sync.Once

func dataInputStreamSetByteOrderFunction_Set() error {
	var err error
	dataInputStreamSetByteOrderFunction_Once.Do(func() {
		err = dataInputStreamObject_Set()
		if err != nil {
			return
		}
		dataInputStreamSetByteOrderFunction, err = dataInputStreamObject.InvokerNew("set_byte_order")
	})
	return err
}

// SetByteOrder is a representation of the C type g_data_input_stream_set_byte_order.
func (recv *DataInputStream) SetByteOrder(order DataStreamByteOrder) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(int32(order))

	err := dataInputStreamSetByteOrderFunction_Set()
	if err == nil {
		dataInputStreamSetByteOrderFunction.Invoke(inArgs[:], nil)
	}

	return
}

var dataInputStreamSetNewlineTypeFunction *gi.Function
var dataInputStreamSetNewlineTypeFunction_Once sync.Once

func dataInputStreamSetNewlineTypeFunction_Set() error {
	var err error
	dataInputStreamSetNewlineTypeFunction_Once.Do(func() {
		err = dataInputStreamObject_Set()
		if err != nil {
			return
		}
		dataInputStreamSetNewlineTypeFunction, err = dataInputStreamObject.InvokerNew("set_newline_type")
	})
	return err
}

// SetNewlineType is a representation of the C type g_data_input_stream_set_newline_type.
func (recv *DataInputStream) SetNewlineType(type_ DataStreamNewlineType) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(int32(type_))

	err := dataInputStreamSetNewlineTypeFunction_Set()
	if err == nil {
		dataInputStreamSetNewlineTypeFunction.Invoke(inArgs[:], nil)
	}

	return
}

var dataOutputStreamObject *gi.Object
var dataOutputStreamObject_Once sync.Once

func dataOutputStreamObject_Set() error {
	var err error
	dataOutputStreamObject_Once.Do(func() {
		dataOutputStreamObject, err = gi.ObjectNew("Gio", "DataOutputStream")
	})
	return err
}

type DataOutputStream struct {
	native unsafe.Pointer
}

func DataOutputStreamNewFromNative(native unsafe.Pointer) *DataOutputStream {
	instance := &DataOutputStream{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// FilterOutputStream upcasts to *FilterOutputStream
func (recv *DataOutputStream) FilterOutputStream() *FilterOutputStream {
	return FilterOutputStreamNewFromNative(recv.Native())
}

// OutputStream upcasts to *OutputStream
func (recv *DataOutputStream) OutputStream() *OutputStream {
	return OutputStreamNewFromNative(recv.Native())
}

// Object upcasts to *Object
func (recv *DataOutputStream) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToDataOutputStream down casts any arbitrary Object to DataOutputStream.
Exercise care, as this is a potentially dangerous function
if the Object is not a DataOutputStream.
*/
func CastToDataOutputStream(object *gobject.Object) *DataOutputStream {
	return DataOutputStreamNewFromNative(object.Native())
}

// Equals compares this DataOutputStream with another DataOutputStream, and returns true if they represent the same GObject.
func (recv *DataOutputStream) Equals(other *DataOutputStream) bool {
	return other.Native() == recv.Native()
}

func (recv *DataOutputStream) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *DataOutputStream) FieldParentInstance() *FilterOutputStream {
	argValue := gi.ObjectFieldGet(dataOutputStreamObject, recv.Native(), "parent_instance")
	value := FilterOutputStreamNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *DataOutputStream) SetFieldParentInstance(value *FilterOutputStream) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(dataOutputStreamObject, recv.Native(), "parent_instance", argValue)
}

var dataOutputStreamNewFunction *gi.Function
var dataOutputStreamNewFunction_Once sync.Once

func dataOutputStreamNewFunction_Set() error {
	var err error
	dataOutputStreamNewFunction_Once.Do(func() {
		err = dataOutputStreamObject_Set()
		if err != nil {
			return
		}
		dataOutputStreamNewFunction, err = dataOutputStreamObject.InvokerNew("new")
	})
	return err
}

// DataOutputStreamNew is a representation of the C type g_data_output_stream_new.
func DataOutputStreamNew(baseStream *OutputStream) *DataOutputStream {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(baseStream.Native())

	var ret gi.Argument

	err := dataOutputStreamNewFunction_Set()
	if err == nil {
		ret = dataOutputStreamNewFunction.Invoke(inArgs[:], nil)
	}

	retGo := DataOutputStreamNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

var dataOutputStreamGetByteOrderFunction *gi.Function
var dataOutputStreamGetByteOrderFunction_Once sync.Once

func dataOutputStreamGetByteOrderFunction_Set() error {
	var err error
	dataOutputStreamGetByteOrderFunction_Once.Do(func() {
		err = dataOutputStreamObject_Set()
		if err != nil {
			return
		}
		dataOutputStreamGetByteOrderFunction, err = dataOutputStreamObject.InvokerNew("get_byte_order")
	})
	return err
}

// GetByteOrder is a representation of the C type g_data_output_stream_get_byte_order.
func (recv *DataOutputStream) GetByteOrder() DataStreamByteOrder {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := dataOutputStreamGetByteOrderFunction_Set()
	if err == nil {
		ret = dataOutputStreamGetByteOrderFunction.Invoke(inArgs[:], nil)
	}

	retGo := DataStreamByteOrder(ret.Int32())

	return retGo
}

var dataOutputStreamPutByteFunction *gi.Function
var dataOutputStreamPutByteFunction_Once sync.Once

func dataOutputStreamPutByteFunction_Set() error {
	var err error
	dataOutputStreamPutByteFunction_Once.Do(func() {
		err = dataOutputStreamObject_Set()
		if err != nil {
			return
		}
		dataOutputStreamPutByteFunction, err = dataOutputStreamObject.InvokerNew("put_byte")
	})
	return err
}

// PutByte is a representation of the C type g_data_output_stream_put_byte.
func (recv *DataOutputStream) PutByte(data uint8, cancellable *Cancellable) bool {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetUint8(data)
	inArgs[2].SetPointer(cancellable.Native())

	var ret gi.Argument

	err := dataOutputStreamPutByteFunction_Set()
	if err == nil {
		ret = dataOutputStreamPutByteFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var dataOutputStreamPutInt16Function *gi.Function
var dataOutputStreamPutInt16Function_Once sync.Once

func dataOutputStreamPutInt16Function_Set() error {
	var err error
	dataOutputStreamPutInt16Function_Once.Do(func() {
		err = dataOutputStreamObject_Set()
		if err != nil {
			return
		}
		dataOutputStreamPutInt16Function, err = dataOutputStreamObject.InvokerNew("put_int16")
	})
	return err
}

// PutInt16 is a representation of the C type g_data_output_stream_put_int16.
func (recv *DataOutputStream) PutInt16(data int16, cancellable *Cancellable) bool {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt16(data)
	inArgs[2].SetPointer(cancellable.Native())

	var ret gi.Argument

	err := dataOutputStreamPutInt16Function_Set()
	if err == nil {
		ret = dataOutputStreamPutInt16Function.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var dataOutputStreamPutInt32Function *gi.Function
var dataOutputStreamPutInt32Function_Once sync.Once

func dataOutputStreamPutInt32Function_Set() error {
	var err error
	dataOutputStreamPutInt32Function_Once.Do(func() {
		err = dataOutputStreamObject_Set()
		if err != nil {
			return
		}
		dataOutputStreamPutInt32Function, err = dataOutputStreamObject.InvokerNew("put_int32")
	})
	return err
}

// PutInt32 is a representation of the C type g_data_output_stream_put_int32.
func (recv *DataOutputStream) PutInt32(data int32, cancellable *Cancellable) bool {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(data)
	inArgs[2].SetPointer(cancellable.Native())

	var ret gi.Argument

	err := dataOutputStreamPutInt32Function_Set()
	if err == nil {
		ret = dataOutputStreamPutInt32Function.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var dataOutputStreamPutInt64Function *gi.Function
var dataOutputStreamPutInt64Function_Once sync.Once

func dataOutputStreamPutInt64Function_Set() error {
	var err error
	dataOutputStreamPutInt64Function_Once.Do(func() {
		err = dataOutputStreamObject_Set()
		if err != nil {
			return
		}
		dataOutputStreamPutInt64Function, err = dataOutputStreamObject.InvokerNew("put_int64")
	})
	return err
}

// PutInt64 is a representation of the C type g_data_output_stream_put_int64.
func (recv *DataOutputStream) PutInt64(data int64, cancellable *Cancellable) bool {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt64(data)
	inArgs[2].SetPointer(cancellable.Native())

	var ret gi.Argument

	err := dataOutputStreamPutInt64Function_Set()
	if err == nil {
		ret = dataOutputStreamPutInt64Function.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var dataOutputStreamPutStringFunction *gi.Function
var dataOutputStreamPutStringFunction_Once sync.Once

func dataOutputStreamPutStringFunction_Set() error {
	var err error
	dataOutputStreamPutStringFunction_Once.Do(func() {
		err = dataOutputStreamObject_Set()
		if err != nil {
			return
		}
		dataOutputStreamPutStringFunction, err = dataOutputStreamObject.InvokerNew("put_string")
	})
	return err
}

// PutString is a representation of the C type g_data_output_stream_put_string.
func (recv *DataOutputStream) PutString(str string, cancellable *Cancellable) bool {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(str)
	inArgs[2].SetPointer(cancellable.Native())

	var ret gi.Argument

	err := dataOutputStreamPutStringFunction_Set()
	if err == nil {
		ret = dataOutputStreamPutStringFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var dataOutputStreamPutUint16Function *gi.Function
var dataOutputStreamPutUint16Function_Once sync.Once

func dataOutputStreamPutUint16Function_Set() error {
	var err error
	dataOutputStreamPutUint16Function_Once.Do(func() {
		err = dataOutputStreamObject_Set()
		if err != nil {
			return
		}
		dataOutputStreamPutUint16Function, err = dataOutputStreamObject.InvokerNew("put_uint16")
	})
	return err
}

// PutUint16 is a representation of the C type g_data_output_stream_put_uint16.
func (recv *DataOutputStream) PutUint16(data uint16, cancellable *Cancellable) bool {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetUint16(data)
	inArgs[2].SetPointer(cancellable.Native())

	var ret gi.Argument

	err := dataOutputStreamPutUint16Function_Set()
	if err == nil {
		ret = dataOutputStreamPutUint16Function.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var dataOutputStreamPutUint32Function *gi.Function
var dataOutputStreamPutUint32Function_Once sync.Once

func dataOutputStreamPutUint32Function_Set() error {
	var err error
	dataOutputStreamPutUint32Function_Once.Do(func() {
		err = dataOutputStreamObject_Set()
		if err != nil {
			return
		}
		dataOutputStreamPutUint32Function, err = dataOutputStreamObject.InvokerNew("put_uint32")
	})
	return err
}

// PutUint32 is a representation of the C type g_data_output_stream_put_uint32.
func (recv *DataOutputStream) PutUint32(data uint32, cancellable *Cancellable) bool {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetUint32(data)
	inArgs[2].SetPointer(cancellable.Native())

	var ret gi.Argument

	err := dataOutputStreamPutUint32Function_Set()
	if err == nil {
		ret = dataOutputStreamPutUint32Function.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var dataOutputStreamPutUint64Function *gi.Function
var dataOutputStreamPutUint64Function_Once sync.Once

func dataOutputStreamPutUint64Function_Set() error {
	var err error
	dataOutputStreamPutUint64Function_Once.Do(func() {
		err = dataOutputStreamObject_Set()
		if err != nil {
			return
		}
		dataOutputStreamPutUint64Function, err = dataOutputStreamObject.InvokerNew("put_uint64")
	})
	return err
}

// PutUint64 is a representation of the C type g_data_output_stream_put_uint64.
func (recv *DataOutputStream) PutUint64(data uint64, cancellable *Cancellable) bool {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetUint64(data)
	inArgs[2].SetPointer(cancellable.Native())

	var ret gi.Argument

	err := dataOutputStreamPutUint64Function_Set()
	if err == nil {
		ret = dataOutputStreamPutUint64Function.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var dataOutputStreamSetByteOrderFunction *gi.Function
var dataOutputStreamSetByteOrderFunction_Once sync.Once

func dataOutputStreamSetByteOrderFunction_Set() error {
	var err error
	dataOutputStreamSetByteOrderFunction_Once.Do(func() {
		err = dataOutputStreamObject_Set()
		if err != nil {
			return
		}
		dataOutputStreamSetByteOrderFunction, err = dataOutputStreamObject.InvokerNew("set_byte_order")
	})
	return err
}

// SetByteOrder is a representation of the C type g_data_output_stream_set_byte_order.
func (recv *DataOutputStream) SetByteOrder(order DataStreamByteOrder) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(int32(order))

	err := dataOutputStreamSetByteOrderFunction_Set()
	if err == nil {
		dataOutputStreamSetByteOrderFunction.Invoke(inArgs[:], nil)
	}

	return
}

var desktopAppInfoObject *gi.Object
var desktopAppInfoObject_Once sync.Once

func desktopAppInfoObject_Set() error {
	var err error
	desktopAppInfoObject_Once.Do(func() {
		desktopAppInfoObject, err = gi.ObjectNew("Gio", "DesktopAppInfo")
	})
	return err
}

type DesktopAppInfo struct {
	native unsafe.Pointer
}

func DesktopAppInfoNewFromNative(native unsafe.Pointer) *DesktopAppInfo {
	instance := &DesktopAppInfo{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *DesktopAppInfo) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToDesktopAppInfo down casts any arbitrary Object to DesktopAppInfo.
Exercise care, as this is a potentially dangerous function
if the Object is not a DesktopAppInfo.
*/
func CastToDesktopAppInfo(object *gobject.Object) *DesktopAppInfo {
	return DesktopAppInfoNewFromNative(object.Native())
}

// Equals compares this DesktopAppInfo with another DesktopAppInfo, and returns true if they represent the same GObject.
func (recv *DesktopAppInfo) Equals(other *DesktopAppInfo) bool {
	return other.Native() == recv.Native()
}

func (recv *DesktopAppInfo) Native() unsafe.Pointer {
	return recv.native
}

var desktopAppInfoNewFunction *gi.Function
var desktopAppInfoNewFunction_Once sync.Once

func desktopAppInfoNewFunction_Set() error {
	var err error
	desktopAppInfoNewFunction_Once.Do(func() {
		err = desktopAppInfoObject_Set()
		if err != nil {
			return
		}
		desktopAppInfoNewFunction, err = desktopAppInfoObject.InvokerNew("new")
	})
	return err
}

// DesktopAppInfoNew is a representation of the C type g_desktop_app_info_new.
func DesktopAppInfoNew(desktopId string) *DesktopAppInfo {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(desktopId)

	var ret gi.Argument

	err := desktopAppInfoNewFunction_Set()
	if err == nil {
		ret = desktopAppInfoNewFunction.Invoke(inArgs[:], nil)
	}

	retGo := DesktopAppInfoNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

var desktopAppInfoNewFromFilenameFunction *gi.Function
var desktopAppInfoNewFromFilenameFunction_Once sync.Once

func desktopAppInfoNewFromFilenameFunction_Set() error {
	var err error
	desktopAppInfoNewFromFilenameFunction_Once.Do(func() {
		err = desktopAppInfoObject_Set()
		if err != nil {
			return
		}
		desktopAppInfoNewFromFilenameFunction, err = desktopAppInfoObject.InvokerNew("new_from_filename")
	})
	return err
}

// DesktopAppInfoNewFromFilename is a representation of the C type g_desktop_app_info_new_from_filename.
func DesktopAppInfoNewFromFilename(filename string) *DesktopAppInfo {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(filename)

	var ret gi.Argument

	err := desktopAppInfoNewFromFilenameFunction_Set()
	if err == nil {
		ret = desktopAppInfoNewFromFilenameFunction.Invoke(inArgs[:], nil)
	}

	retGo := DesktopAppInfoNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

// UNSUPPORTED : C value 'g_desktop_app_info_new_from_keyfile' : parameter 'key_file' of type 'GLib.KeyFile' not supported

var desktopAppInfoGetActionNameFunction *gi.Function
var desktopAppInfoGetActionNameFunction_Once sync.Once

func desktopAppInfoGetActionNameFunction_Set() error {
	var err error
	desktopAppInfoGetActionNameFunction_Once.Do(func() {
		err = desktopAppInfoObject_Set()
		if err != nil {
			return
		}
		desktopAppInfoGetActionNameFunction, err = desktopAppInfoObject.InvokerNew("get_action_name")
	})
	return err
}

// GetActionName is a representation of the C type g_desktop_app_info_get_action_name.
func (recv *DesktopAppInfo) GetActionName(actionName string) string {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(actionName)

	var ret gi.Argument

	err := desktopAppInfoGetActionNameFunction_Set()
	if err == nil {
		ret = desktopAppInfoGetActionNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

var desktopAppInfoGetBooleanFunction *gi.Function
var desktopAppInfoGetBooleanFunction_Once sync.Once

func desktopAppInfoGetBooleanFunction_Set() error {
	var err error
	desktopAppInfoGetBooleanFunction_Once.Do(func() {
		err = desktopAppInfoObject_Set()
		if err != nil {
			return
		}
		desktopAppInfoGetBooleanFunction, err = desktopAppInfoObject.InvokerNew("get_boolean")
	})
	return err
}

// GetBoolean is a representation of the C type g_desktop_app_info_get_boolean.
func (recv *DesktopAppInfo) GetBoolean(key string) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(key)

	var ret gi.Argument

	err := desktopAppInfoGetBooleanFunction_Set()
	if err == nil {
		ret = desktopAppInfoGetBooleanFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var desktopAppInfoGetCategoriesFunction *gi.Function
var desktopAppInfoGetCategoriesFunction_Once sync.Once

func desktopAppInfoGetCategoriesFunction_Set() error {
	var err error
	desktopAppInfoGetCategoriesFunction_Once.Do(func() {
		err = desktopAppInfoObject_Set()
		if err != nil {
			return
		}
		desktopAppInfoGetCategoriesFunction, err = desktopAppInfoObject.InvokerNew("get_categories")
	})
	return err
}

// GetCategories is a representation of the C type g_desktop_app_info_get_categories.
func (recv *DesktopAppInfo) GetCategories() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := desktopAppInfoGetCategoriesFunction_Set()
	if err == nil {
		ret = desktopAppInfoGetCategoriesFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var desktopAppInfoGetFilenameFunction *gi.Function
var desktopAppInfoGetFilenameFunction_Once sync.Once

func desktopAppInfoGetFilenameFunction_Set() error {
	var err error
	desktopAppInfoGetFilenameFunction_Once.Do(func() {
		err = desktopAppInfoObject_Set()
		if err != nil {
			return
		}
		desktopAppInfoGetFilenameFunction, err = desktopAppInfoObject.InvokerNew("get_filename")
	})
	return err
}

// GetFilename is a representation of the C type g_desktop_app_info_get_filename.
func (recv *DesktopAppInfo) GetFilename() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := desktopAppInfoGetFilenameFunction_Set()
	if err == nil {
		ret = desktopAppInfoGetFilenameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var desktopAppInfoGetGenericNameFunction *gi.Function
var desktopAppInfoGetGenericNameFunction_Once sync.Once

func desktopAppInfoGetGenericNameFunction_Set() error {
	var err error
	desktopAppInfoGetGenericNameFunction_Once.Do(func() {
		err = desktopAppInfoObject_Set()
		if err != nil {
			return
		}
		desktopAppInfoGetGenericNameFunction, err = desktopAppInfoObject.InvokerNew("get_generic_name")
	})
	return err
}

// GetGenericName is a representation of the C type g_desktop_app_info_get_generic_name.
func (recv *DesktopAppInfo) GetGenericName() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := desktopAppInfoGetGenericNameFunction_Set()
	if err == nil {
		ret = desktopAppInfoGetGenericNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var desktopAppInfoGetIsHiddenFunction *gi.Function
var desktopAppInfoGetIsHiddenFunction_Once sync.Once

func desktopAppInfoGetIsHiddenFunction_Set() error {
	var err error
	desktopAppInfoGetIsHiddenFunction_Once.Do(func() {
		err = desktopAppInfoObject_Set()
		if err != nil {
			return
		}
		desktopAppInfoGetIsHiddenFunction, err = desktopAppInfoObject.InvokerNew("get_is_hidden")
	})
	return err
}

// GetIsHidden is a representation of the C type g_desktop_app_info_get_is_hidden.
func (recv *DesktopAppInfo) GetIsHidden() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := desktopAppInfoGetIsHiddenFunction_Set()
	if err == nil {
		ret = desktopAppInfoGetIsHiddenFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var desktopAppInfoGetKeywordsFunction *gi.Function
var desktopAppInfoGetKeywordsFunction_Once sync.Once

func desktopAppInfoGetKeywordsFunction_Set() error {
	var err error
	desktopAppInfoGetKeywordsFunction_Once.Do(func() {
		err = desktopAppInfoObject_Set()
		if err != nil {
			return
		}
		desktopAppInfoGetKeywordsFunction, err = desktopAppInfoObject.InvokerNew("get_keywords")
	})
	return err
}

// GetKeywords is a representation of the C type g_desktop_app_info_get_keywords.
func (recv *DesktopAppInfo) GetKeywords() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := desktopAppInfoGetKeywordsFunction_Set()
	if err == nil {
		desktopAppInfoGetKeywordsFunction.Invoke(inArgs[:], nil)
	}

	return
}

var desktopAppInfoGetLocaleStringFunction *gi.Function
var desktopAppInfoGetLocaleStringFunction_Once sync.Once

func desktopAppInfoGetLocaleStringFunction_Set() error {
	var err error
	desktopAppInfoGetLocaleStringFunction_Once.Do(func() {
		err = desktopAppInfoObject_Set()
		if err != nil {
			return
		}
		desktopAppInfoGetLocaleStringFunction, err = desktopAppInfoObject.InvokerNew("get_locale_string")
	})
	return err
}

// GetLocaleString is a representation of the C type g_desktop_app_info_get_locale_string.
func (recv *DesktopAppInfo) GetLocaleString(key string) string {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(key)

	var ret gi.Argument

	err := desktopAppInfoGetLocaleStringFunction_Set()
	if err == nil {
		ret = desktopAppInfoGetLocaleStringFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

var desktopAppInfoGetNodisplayFunction *gi.Function
var desktopAppInfoGetNodisplayFunction_Once sync.Once

func desktopAppInfoGetNodisplayFunction_Set() error {
	var err error
	desktopAppInfoGetNodisplayFunction_Once.Do(func() {
		err = desktopAppInfoObject_Set()
		if err != nil {
			return
		}
		desktopAppInfoGetNodisplayFunction, err = desktopAppInfoObject.InvokerNew("get_nodisplay")
	})
	return err
}

// GetNodisplay is a representation of the C type g_desktop_app_info_get_nodisplay.
func (recv *DesktopAppInfo) GetNodisplay() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := desktopAppInfoGetNodisplayFunction_Set()
	if err == nil {
		ret = desktopAppInfoGetNodisplayFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var desktopAppInfoGetShowInFunction *gi.Function
var desktopAppInfoGetShowInFunction_Once sync.Once

func desktopAppInfoGetShowInFunction_Set() error {
	var err error
	desktopAppInfoGetShowInFunction_Once.Do(func() {
		err = desktopAppInfoObject_Set()
		if err != nil {
			return
		}
		desktopAppInfoGetShowInFunction, err = desktopAppInfoObject.InvokerNew("get_show_in")
	})
	return err
}

// GetShowIn is a representation of the C type g_desktop_app_info_get_show_in.
func (recv *DesktopAppInfo) GetShowIn(desktopEnv string) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(desktopEnv)

	var ret gi.Argument

	err := desktopAppInfoGetShowInFunction_Set()
	if err == nil {
		ret = desktopAppInfoGetShowInFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var desktopAppInfoGetStartupWmClassFunction *gi.Function
var desktopAppInfoGetStartupWmClassFunction_Once sync.Once

func desktopAppInfoGetStartupWmClassFunction_Set() error {
	var err error
	desktopAppInfoGetStartupWmClassFunction_Once.Do(func() {
		err = desktopAppInfoObject_Set()
		if err != nil {
			return
		}
		desktopAppInfoGetStartupWmClassFunction, err = desktopAppInfoObject.InvokerNew("get_startup_wm_class")
	})
	return err
}

// GetStartupWmClass is a representation of the C type g_desktop_app_info_get_startup_wm_class.
func (recv *DesktopAppInfo) GetStartupWmClass() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := desktopAppInfoGetStartupWmClassFunction_Set()
	if err == nil {
		ret = desktopAppInfoGetStartupWmClassFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var desktopAppInfoGetStringFunction *gi.Function
var desktopAppInfoGetStringFunction_Once sync.Once

func desktopAppInfoGetStringFunction_Set() error {
	var err error
	desktopAppInfoGetStringFunction_Once.Do(func() {
		err = desktopAppInfoObject_Set()
		if err != nil {
			return
		}
		desktopAppInfoGetStringFunction, err = desktopAppInfoObject.InvokerNew("get_string")
	})
	return err
}

// GetString is a representation of the C type g_desktop_app_info_get_string.
func (recv *DesktopAppInfo) GetString(key string) string {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(key)

	var ret gi.Argument

	err := desktopAppInfoGetStringFunction_Set()
	if err == nil {
		ret = desktopAppInfoGetStringFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

var desktopAppInfoGetStringListFunction *gi.Function
var desktopAppInfoGetStringListFunction_Once sync.Once

func desktopAppInfoGetStringListFunction_Set() error {
	var err error
	desktopAppInfoGetStringListFunction_Once.Do(func() {
		err = desktopAppInfoObject_Set()
		if err != nil {
			return
		}
		desktopAppInfoGetStringListFunction, err = desktopAppInfoObject.InvokerNew("get_string_list")
	})
	return err
}

// GetStringList is a representation of the C type g_desktop_app_info_get_string_list.
func (recv *DesktopAppInfo) GetStringList(key string) uint64 {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(key)

	var outArgs [1]gi.Argument

	err := desktopAppInfoGetStringListFunction_Set()
	if err == nil {
		desktopAppInfoGetStringListFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := outArgs[0].Uint64()

	return out0
}

var desktopAppInfoHasKeyFunction *gi.Function
var desktopAppInfoHasKeyFunction_Once sync.Once

func desktopAppInfoHasKeyFunction_Set() error {
	var err error
	desktopAppInfoHasKeyFunction_Once.Do(func() {
		err = desktopAppInfoObject_Set()
		if err != nil {
			return
		}
		desktopAppInfoHasKeyFunction, err = desktopAppInfoObject.InvokerNew("has_key")
	})
	return err
}

// HasKey is a representation of the C type g_desktop_app_info_has_key.
func (recv *DesktopAppInfo) HasKey(key string) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(key)

	var ret gi.Argument

	err := desktopAppInfoHasKeyFunction_Set()
	if err == nil {
		ret = desktopAppInfoHasKeyFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var desktopAppInfoLaunchActionFunction *gi.Function
var desktopAppInfoLaunchActionFunction_Once sync.Once

func desktopAppInfoLaunchActionFunction_Set() error {
	var err error
	desktopAppInfoLaunchActionFunction_Once.Do(func() {
		err = desktopAppInfoObject_Set()
		if err != nil {
			return
		}
		desktopAppInfoLaunchActionFunction, err = desktopAppInfoObject.InvokerNew("launch_action")
	})
	return err
}

// LaunchAction is a representation of the C type g_desktop_app_info_launch_action.
func (recv *DesktopAppInfo) LaunchAction(actionName string, launchContext *AppLaunchContext) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(actionName)
	inArgs[2].SetPointer(launchContext.Native())

	err := desktopAppInfoLaunchActionFunction_Set()
	if err == nil {
		desktopAppInfoLaunchActionFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'g_desktop_app_info_launch_uris_as_manager' : parameter 'uris' of type 'GLib.List' not supported

// UNSUPPORTED : C value 'g_desktop_app_info_launch_uris_as_manager_with_fds' : parameter 'uris' of type 'GLib.List' not supported

var desktopAppInfoListActionsFunction *gi.Function
var desktopAppInfoListActionsFunction_Once sync.Once

func desktopAppInfoListActionsFunction_Set() error {
	var err error
	desktopAppInfoListActionsFunction_Once.Do(func() {
		err = desktopAppInfoObject_Set()
		if err != nil {
			return
		}
		desktopAppInfoListActionsFunction, err = desktopAppInfoObject.InvokerNew("list_actions")
	})
	return err
}

// ListActions is a representation of the C type g_desktop_app_info_list_actions.
func (recv *DesktopAppInfo) ListActions() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := desktopAppInfoListActionsFunction_Set()
	if err == nil {
		desktopAppInfoListActionsFunction.Invoke(inArgs[:], nil)
	}

	return
}

var emblemObject *gi.Object
var emblemObject_Once sync.Once

func emblemObject_Set() error {
	var err error
	emblemObject_Once.Do(func() {
		emblemObject, err = gi.ObjectNew("Gio", "Emblem")
	})
	return err
}

type Emblem struct {
	native unsafe.Pointer
}

func EmblemNewFromNative(native unsafe.Pointer) *Emblem {
	instance := &Emblem{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *Emblem) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToEmblem down casts any arbitrary Object to Emblem.
Exercise care, as this is a potentially dangerous function
if the Object is not a Emblem.
*/
func CastToEmblem(object *gobject.Object) *Emblem {
	return EmblemNewFromNative(object.Native())
}

// Equals compares this Emblem with another Emblem, and returns true if they represent the same GObject.
func (recv *Emblem) Equals(other *Emblem) bool {
	return other.Native() == recv.Native()
}

func (recv *Emblem) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'g_emblem_new' : parameter 'icon' of type 'Icon' not supported

// UNSUPPORTED : C value 'g_emblem_new_with_origin' : parameter 'icon' of type 'Icon' not supported

// UNSUPPORTED : C value 'g_emblem_get_icon' : return type 'Icon' not supported

var emblemGetOriginFunction *gi.Function
var emblemGetOriginFunction_Once sync.Once

func emblemGetOriginFunction_Set() error {
	var err error
	emblemGetOriginFunction_Once.Do(func() {
		err = emblemObject_Set()
		if err != nil {
			return
		}
		emblemGetOriginFunction, err = emblemObject.InvokerNew("get_origin")
	})
	return err
}

// GetOrigin is a representation of the C type g_emblem_get_origin.
func (recv *Emblem) GetOrigin() EmblemOrigin {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := emblemGetOriginFunction_Set()
	if err == nil {
		ret = emblemGetOriginFunction.Invoke(inArgs[:], nil)
	}

	retGo := EmblemOrigin(ret.Int32())

	return retGo
}

var emblemedIconObject *gi.Object
var emblemedIconObject_Once sync.Once

func emblemedIconObject_Set() error {
	var err error
	emblemedIconObject_Once.Do(func() {
		emblemedIconObject, err = gi.ObjectNew("Gio", "EmblemedIcon")
	})
	return err
}

type EmblemedIcon struct {
	native unsafe.Pointer
}

func EmblemedIconNewFromNative(native unsafe.Pointer) *EmblemedIcon {
	instance := &EmblemedIcon{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *EmblemedIcon) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToEmblemedIcon down casts any arbitrary Object to EmblemedIcon.
Exercise care, as this is a potentially dangerous function
if the Object is not a EmblemedIcon.
*/
func CastToEmblemedIcon(object *gobject.Object) *EmblemedIcon {
	return EmblemedIconNewFromNative(object.Native())
}

// Equals compares this EmblemedIcon with another EmblemedIcon, and returns true if they represent the same GObject.
func (recv *EmblemedIcon) Equals(other *EmblemedIcon) bool {
	return other.Native() == recv.Native()
}

func (recv *EmblemedIcon) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *EmblemedIcon) FieldParentInstance() *gobject.Object {
	argValue := gi.ObjectFieldGet(emblemedIconObject, recv.Native(), "parent_instance")
	value := gobject.ObjectNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *EmblemedIcon) SetFieldParentInstance(value *gobject.Object) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(emblemedIconObject, recv.Native(), "parent_instance", argValue)
}

// UNSUPPORTED : C value 'g_emblemed_icon_new' : parameter 'icon' of type 'Icon' not supported

var emblemedIconAddEmblemFunction *gi.Function
var emblemedIconAddEmblemFunction_Once sync.Once

func emblemedIconAddEmblemFunction_Set() error {
	var err error
	emblemedIconAddEmblemFunction_Once.Do(func() {
		err = emblemedIconObject_Set()
		if err != nil {
			return
		}
		emblemedIconAddEmblemFunction, err = emblemedIconObject.InvokerNew("add_emblem")
	})
	return err
}

// AddEmblem is a representation of the C type g_emblemed_icon_add_emblem.
func (recv *EmblemedIcon) AddEmblem(emblem *Emblem) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(emblem.Native())

	err := emblemedIconAddEmblemFunction_Set()
	if err == nil {
		emblemedIconAddEmblemFunction.Invoke(inArgs[:], nil)
	}

	return
}

var emblemedIconClearEmblemsFunction *gi.Function
var emblemedIconClearEmblemsFunction_Once sync.Once

func emblemedIconClearEmblemsFunction_Set() error {
	var err error
	emblemedIconClearEmblemsFunction_Once.Do(func() {
		err = emblemedIconObject_Set()
		if err != nil {
			return
		}
		emblemedIconClearEmblemsFunction, err = emblemedIconObject.InvokerNew("clear_emblems")
	})
	return err
}

// ClearEmblems is a representation of the C type g_emblemed_icon_clear_emblems.
func (recv *EmblemedIcon) ClearEmblems() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := emblemedIconClearEmblemsFunction_Set()
	if err == nil {
		emblemedIconClearEmblemsFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'g_emblemed_icon_get_emblems' : return type 'GLib.List' not supported

// UNSUPPORTED : C value 'g_emblemed_icon_get_icon' : return type 'Icon' not supported

var fileEnumeratorObject *gi.Object
var fileEnumeratorObject_Once sync.Once

func fileEnumeratorObject_Set() error {
	var err error
	fileEnumeratorObject_Once.Do(func() {
		fileEnumeratorObject, err = gi.ObjectNew("Gio", "FileEnumerator")
	})
	return err
}

type FileEnumerator struct {
	native unsafe.Pointer
}

func FileEnumeratorNewFromNative(native unsafe.Pointer) *FileEnumerator {
	instance := &FileEnumerator{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *FileEnumerator) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToFileEnumerator down casts any arbitrary Object to FileEnumerator.
Exercise care, as this is a potentially dangerous function
if the Object is not a FileEnumerator.
*/
func CastToFileEnumerator(object *gobject.Object) *FileEnumerator {
	return FileEnumeratorNewFromNative(object.Native())
}

// Equals compares this FileEnumerator with another FileEnumerator, and returns true if they represent the same GObject.
func (recv *FileEnumerator) Equals(other *FileEnumerator) bool {
	return other.Native() == recv.Native()
}

func (recv *FileEnumerator) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *FileEnumerator) FieldParentInstance() *gobject.Object {
	argValue := gi.ObjectFieldGet(fileEnumeratorObject, recv.Native(), "parent_instance")
	value := gobject.ObjectNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *FileEnumerator) SetFieldParentInstance(value *gobject.Object) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(fileEnumeratorObject, recv.Native(), "parent_instance", argValue)
}

var fileEnumeratorCloseFunction *gi.Function
var fileEnumeratorCloseFunction_Once sync.Once

func fileEnumeratorCloseFunction_Set() error {
	var err error
	fileEnumeratorCloseFunction_Once.Do(func() {
		err = fileEnumeratorObject_Set()
		if err != nil {
			return
		}
		fileEnumeratorCloseFunction, err = fileEnumeratorObject.InvokerNew("close")
	})
	return err
}

// Close is a representation of the C type g_file_enumerator_close.
func (recv *FileEnumerator) Close(cancellable *Cancellable) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(cancellable.Native())

	var ret gi.Argument

	err := fileEnumeratorCloseFunction_Set()
	if err == nil {
		ret = fileEnumeratorCloseFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'g_file_enumerator_close_async' : parameter 'callback' of type 'AsyncReadyCallback' not supported

// UNSUPPORTED : C value 'g_file_enumerator_close_finish' : parameter 'result' of type 'AsyncResult' not supported

// UNSUPPORTED : C value 'g_file_enumerator_get_child' : return type 'File' not supported

// UNSUPPORTED : C value 'g_file_enumerator_get_container' : return type 'File' not supported

var fileEnumeratorHasPendingFunction *gi.Function
var fileEnumeratorHasPendingFunction_Once sync.Once

func fileEnumeratorHasPendingFunction_Set() error {
	var err error
	fileEnumeratorHasPendingFunction_Once.Do(func() {
		err = fileEnumeratorObject_Set()
		if err != nil {
			return
		}
		fileEnumeratorHasPendingFunction, err = fileEnumeratorObject.InvokerNew("has_pending")
	})
	return err
}

// HasPending is a representation of the C type g_file_enumerator_has_pending.
func (recv *FileEnumerator) HasPending() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := fileEnumeratorHasPendingFunction_Set()
	if err == nil {
		ret = fileEnumeratorHasPendingFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var fileEnumeratorIsClosedFunction *gi.Function
var fileEnumeratorIsClosedFunction_Once sync.Once

func fileEnumeratorIsClosedFunction_Set() error {
	var err error
	fileEnumeratorIsClosedFunction_Once.Do(func() {
		err = fileEnumeratorObject_Set()
		if err != nil {
			return
		}
		fileEnumeratorIsClosedFunction, err = fileEnumeratorObject.InvokerNew("is_closed")
	})
	return err
}

// IsClosed is a representation of the C type g_file_enumerator_is_closed.
func (recv *FileEnumerator) IsClosed() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := fileEnumeratorIsClosedFunction_Set()
	if err == nil {
		ret = fileEnumeratorIsClosedFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'g_file_enumerator_iterate' : parameter 'out_child' of type 'File' not supported

var fileEnumeratorNextFileFunction *gi.Function
var fileEnumeratorNextFileFunction_Once sync.Once

func fileEnumeratorNextFileFunction_Set() error {
	var err error
	fileEnumeratorNextFileFunction_Once.Do(func() {
		err = fileEnumeratorObject_Set()
		if err != nil {
			return
		}
		fileEnumeratorNextFileFunction, err = fileEnumeratorObject.InvokerNew("next_file")
	})
	return err
}

// NextFile is a representation of the C type g_file_enumerator_next_file.
func (recv *FileEnumerator) NextFile(cancellable *Cancellable) *FileInfo {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(cancellable.Native())

	var ret gi.Argument

	err := fileEnumeratorNextFileFunction_Set()
	if err == nil {
		ret = fileEnumeratorNextFileFunction.Invoke(inArgs[:], nil)
	}

	retGo := FileInfoNewFromNative(ret.Pointer())

	return retGo
}

// UNSUPPORTED : C value 'g_file_enumerator_next_files_async' : parameter 'callback' of type 'AsyncReadyCallback' not supported

// UNSUPPORTED : C value 'g_file_enumerator_next_files_finish' : parameter 'result' of type 'AsyncResult' not supported

var fileEnumeratorSetPendingFunction *gi.Function
var fileEnumeratorSetPendingFunction_Once sync.Once

func fileEnumeratorSetPendingFunction_Set() error {
	var err error
	fileEnumeratorSetPendingFunction_Once.Do(func() {
		err = fileEnumeratorObject_Set()
		if err != nil {
			return
		}
		fileEnumeratorSetPendingFunction, err = fileEnumeratorObject.InvokerNew("set_pending")
	})
	return err
}

// SetPending is a representation of the C type g_file_enumerator_set_pending.
func (recv *FileEnumerator) SetPending(pending bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetBoolean(pending)

	err := fileEnumeratorSetPendingFunction_Set()
	if err == nil {
		fileEnumeratorSetPendingFunction.Invoke(inArgs[:], nil)
	}

	return
}

var fileIOStreamObject *gi.Object
var fileIOStreamObject_Once sync.Once

func fileIOStreamObject_Set() error {
	var err error
	fileIOStreamObject_Once.Do(func() {
		fileIOStreamObject, err = gi.ObjectNew("Gio", "FileIOStream")
	})
	return err
}

type FileIOStream struct {
	native unsafe.Pointer
}

func FileIOStreamNewFromNative(native unsafe.Pointer) *FileIOStream {
	instance := &FileIOStream{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// IOStream upcasts to *IOStream
func (recv *FileIOStream) IOStream() *IOStream {
	return IOStreamNewFromNative(recv.Native())
}

// Object upcasts to *Object
func (recv *FileIOStream) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToFileIOStream down casts any arbitrary Object to FileIOStream.
Exercise care, as this is a potentially dangerous function
if the Object is not a FileIOStream.
*/
func CastToFileIOStream(object *gobject.Object) *FileIOStream {
	return FileIOStreamNewFromNative(object.Native())
}

// Equals compares this FileIOStream with another FileIOStream, and returns true if they represent the same GObject.
func (recv *FileIOStream) Equals(other *FileIOStream) bool {
	return other.Native() == recv.Native()
}

func (recv *FileIOStream) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *FileIOStream) FieldParentInstance() *IOStream {
	argValue := gi.ObjectFieldGet(fileIOStreamObject, recv.Native(), "parent_instance")
	value := IOStreamNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *FileIOStream) SetFieldParentInstance(value *IOStream) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(fileIOStreamObject, recv.Native(), "parent_instance", argValue)
}

var fileIOStreamGetEtagFunction *gi.Function
var fileIOStreamGetEtagFunction_Once sync.Once

func fileIOStreamGetEtagFunction_Set() error {
	var err error
	fileIOStreamGetEtagFunction_Once.Do(func() {
		err = fileIOStreamObject_Set()
		if err != nil {
			return
		}
		fileIOStreamGetEtagFunction, err = fileIOStreamObject.InvokerNew("get_etag")
	})
	return err
}

// GetEtag is a representation of the C type g_file_io_stream_get_etag.
func (recv *FileIOStream) GetEtag() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := fileIOStreamGetEtagFunction_Set()
	if err == nil {
		ret = fileIOStreamGetEtagFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

var fileIOStreamQueryInfoFunction *gi.Function
var fileIOStreamQueryInfoFunction_Once sync.Once

func fileIOStreamQueryInfoFunction_Set() error {
	var err error
	fileIOStreamQueryInfoFunction_Once.Do(func() {
		err = fileIOStreamObject_Set()
		if err != nil {
			return
		}
		fileIOStreamQueryInfoFunction, err = fileIOStreamObject.InvokerNew("query_info")
	})
	return err
}

// QueryInfo is a representation of the C type g_file_io_stream_query_info.
func (recv *FileIOStream) QueryInfo(attributes string, cancellable *Cancellable) *FileInfo {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(attributes)
	inArgs[2].SetPointer(cancellable.Native())

	var ret gi.Argument

	err := fileIOStreamQueryInfoFunction_Set()
	if err == nil {
		ret = fileIOStreamQueryInfoFunction.Invoke(inArgs[:], nil)
	}

	retGo := FileInfoNewFromNative(ret.Pointer())

	return retGo
}

// UNSUPPORTED : C value 'g_file_io_stream_query_info_async' : parameter 'callback' of type 'AsyncReadyCallback' not supported

// UNSUPPORTED : C value 'g_file_io_stream_query_info_finish' : parameter 'result' of type 'AsyncResult' not supported

var fileIconObject *gi.Object
var fileIconObject_Once sync.Once

func fileIconObject_Set() error {
	var err error
	fileIconObject_Once.Do(func() {
		fileIconObject, err = gi.ObjectNew("Gio", "FileIcon")
	})
	return err
}

type FileIcon struct {
	native unsafe.Pointer
}

func FileIconNewFromNative(native unsafe.Pointer) *FileIcon {
	instance := &FileIcon{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *FileIcon) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToFileIcon down casts any arbitrary Object to FileIcon.
Exercise care, as this is a potentially dangerous function
if the Object is not a FileIcon.
*/
func CastToFileIcon(object *gobject.Object) *FileIcon {
	return FileIconNewFromNative(object.Native())
}

// Equals compares this FileIcon with another FileIcon, and returns true if they represent the same GObject.
func (recv *FileIcon) Equals(other *FileIcon) bool {
	return other.Native() == recv.Native()
}

func (recv *FileIcon) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'g_file_icon_new' : parameter 'file' of type 'File' not supported

// UNSUPPORTED : C value 'g_file_icon_get_file' : return type 'File' not supported

var fileInfoObject *gi.Object
var fileInfoObject_Once sync.Once

func fileInfoObject_Set() error {
	var err error
	fileInfoObject_Once.Do(func() {
		fileInfoObject, err = gi.ObjectNew("Gio", "FileInfo")
	})
	return err
}

type FileInfo struct {
	native unsafe.Pointer
}

func FileInfoNewFromNative(native unsafe.Pointer) *FileInfo {
	instance := &FileInfo{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *FileInfo) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToFileInfo down casts any arbitrary Object to FileInfo.
Exercise care, as this is a potentially dangerous function
if the Object is not a FileInfo.
*/
func CastToFileInfo(object *gobject.Object) *FileInfo {
	return FileInfoNewFromNative(object.Native())
}

// Equals compares this FileInfo with another FileInfo, and returns true if they represent the same GObject.
func (recv *FileInfo) Equals(other *FileInfo) bool {
	return other.Native() == recv.Native()
}

func (recv *FileInfo) Native() unsafe.Pointer {
	return recv.native
}

var fileInfoNewFunction *gi.Function
var fileInfoNewFunction_Once sync.Once

func fileInfoNewFunction_Set() error {
	var err error
	fileInfoNewFunction_Once.Do(func() {
		err = fileInfoObject_Set()
		if err != nil {
			return
		}
		fileInfoNewFunction, err = fileInfoObject.InvokerNew("new")
	})
	return err
}

// FileInfoNew is a representation of the C type g_file_info_new.
func FileInfoNew() *FileInfo {

	var ret gi.Argument

	err := fileInfoNewFunction_Set()
	if err == nil {
		ret = fileInfoNewFunction.Invoke(nil, nil)
	}

	retGo := FileInfoNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

var fileInfoClearStatusFunction *gi.Function
var fileInfoClearStatusFunction_Once sync.Once

func fileInfoClearStatusFunction_Set() error {
	var err error
	fileInfoClearStatusFunction_Once.Do(func() {
		err = fileInfoObject_Set()
		if err != nil {
			return
		}
		fileInfoClearStatusFunction, err = fileInfoObject.InvokerNew("clear_status")
	})
	return err
}

// ClearStatus is a representation of the C type g_file_info_clear_status.
func (recv *FileInfo) ClearStatus() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := fileInfoClearStatusFunction_Set()
	if err == nil {
		fileInfoClearStatusFunction.Invoke(inArgs[:], nil)
	}

	return
}

var fileInfoCopyIntoFunction *gi.Function
var fileInfoCopyIntoFunction_Once sync.Once

func fileInfoCopyIntoFunction_Set() error {
	var err error
	fileInfoCopyIntoFunction_Once.Do(func() {
		err = fileInfoObject_Set()
		if err != nil {
			return
		}
		fileInfoCopyIntoFunction, err = fileInfoObject.InvokerNew("copy_into")
	})
	return err
}

// CopyInto is a representation of the C type g_file_info_copy_into.
func (recv *FileInfo) CopyInto(destInfo *FileInfo) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(destInfo.Native())

	err := fileInfoCopyIntoFunction_Set()
	if err == nil {
		fileInfoCopyIntoFunction.Invoke(inArgs[:], nil)
	}

	return
}

var fileInfoDupFunction *gi.Function
var fileInfoDupFunction_Once sync.Once

func fileInfoDupFunction_Set() error {
	var err error
	fileInfoDupFunction_Once.Do(func() {
		err = fileInfoObject_Set()
		if err != nil {
			return
		}
		fileInfoDupFunction, err = fileInfoObject.InvokerNew("dup")
	})
	return err
}

// Dup is a representation of the C type g_file_info_dup.
func (recv *FileInfo) Dup() *FileInfo {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := fileInfoDupFunction_Set()
	if err == nil {
		ret = fileInfoDupFunction.Invoke(inArgs[:], nil)
	}

	retGo := FileInfoNewFromNative(ret.Pointer())

	return retGo
}

var fileInfoGetAttributeAsStringFunction *gi.Function
var fileInfoGetAttributeAsStringFunction_Once sync.Once

func fileInfoGetAttributeAsStringFunction_Set() error {
	var err error
	fileInfoGetAttributeAsStringFunction_Once.Do(func() {
		err = fileInfoObject_Set()
		if err != nil {
			return
		}
		fileInfoGetAttributeAsStringFunction, err = fileInfoObject.InvokerNew("get_attribute_as_string")
	})
	return err
}

// GetAttributeAsString is a representation of the C type g_file_info_get_attribute_as_string.
func (recv *FileInfo) GetAttributeAsString(attribute string) string {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(attribute)

	var ret gi.Argument

	err := fileInfoGetAttributeAsStringFunction_Set()
	if err == nil {
		ret = fileInfoGetAttributeAsStringFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

var fileInfoGetAttributeBooleanFunction *gi.Function
var fileInfoGetAttributeBooleanFunction_Once sync.Once

func fileInfoGetAttributeBooleanFunction_Set() error {
	var err error
	fileInfoGetAttributeBooleanFunction_Once.Do(func() {
		err = fileInfoObject_Set()
		if err != nil {
			return
		}
		fileInfoGetAttributeBooleanFunction, err = fileInfoObject.InvokerNew("get_attribute_boolean")
	})
	return err
}

// GetAttributeBoolean is a representation of the C type g_file_info_get_attribute_boolean.
func (recv *FileInfo) GetAttributeBoolean(attribute string) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(attribute)

	var ret gi.Argument

	err := fileInfoGetAttributeBooleanFunction_Set()
	if err == nil {
		ret = fileInfoGetAttributeBooleanFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var fileInfoGetAttributeByteStringFunction *gi.Function
var fileInfoGetAttributeByteStringFunction_Once sync.Once

func fileInfoGetAttributeByteStringFunction_Set() error {
	var err error
	fileInfoGetAttributeByteStringFunction_Once.Do(func() {
		err = fileInfoObject_Set()
		if err != nil {
			return
		}
		fileInfoGetAttributeByteStringFunction, err = fileInfoObject.InvokerNew("get_attribute_byte_string")
	})
	return err
}

// GetAttributeByteString is a representation of the C type g_file_info_get_attribute_byte_string.
func (recv *FileInfo) GetAttributeByteString(attribute string) string {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(attribute)

	var ret gi.Argument

	err := fileInfoGetAttributeByteStringFunction_Set()
	if err == nil {
		ret = fileInfoGetAttributeByteStringFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'g_file_info_get_attribute_data' : parameter 'value_pp' of type 'gpointer' not supported

var fileInfoGetAttributeInt32Function *gi.Function
var fileInfoGetAttributeInt32Function_Once sync.Once

func fileInfoGetAttributeInt32Function_Set() error {
	var err error
	fileInfoGetAttributeInt32Function_Once.Do(func() {
		err = fileInfoObject_Set()
		if err != nil {
			return
		}
		fileInfoGetAttributeInt32Function, err = fileInfoObject.InvokerNew("get_attribute_int32")
	})
	return err
}

// GetAttributeInt32 is a representation of the C type g_file_info_get_attribute_int32.
func (recv *FileInfo) GetAttributeInt32(attribute string) int32 {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(attribute)

	var ret gi.Argument

	err := fileInfoGetAttributeInt32Function_Set()
	if err == nil {
		ret = fileInfoGetAttributeInt32Function.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var fileInfoGetAttributeInt64Function *gi.Function
var fileInfoGetAttributeInt64Function_Once sync.Once

func fileInfoGetAttributeInt64Function_Set() error {
	var err error
	fileInfoGetAttributeInt64Function_Once.Do(func() {
		err = fileInfoObject_Set()
		if err != nil {
			return
		}
		fileInfoGetAttributeInt64Function, err = fileInfoObject.InvokerNew("get_attribute_int64")
	})
	return err
}

// GetAttributeInt64 is a representation of the C type g_file_info_get_attribute_int64.
func (recv *FileInfo) GetAttributeInt64(attribute string) int64 {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(attribute)

	var ret gi.Argument

	err := fileInfoGetAttributeInt64Function_Set()
	if err == nil {
		ret = fileInfoGetAttributeInt64Function.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int64()

	return retGo
}

var fileInfoGetAttributeObjectFunction *gi.Function
var fileInfoGetAttributeObjectFunction_Once sync.Once

func fileInfoGetAttributeObjectFunction_Set() error {
	var err error
	fileInfoGetAttributeObjectFunction_Once.Do(func() {
		err = fileInfoObject_Set()
		if err != nil {
			return
		}
		fileInfoGetAttributeObjectFunction, err = fileInfoObject.InvokerNew("get_attribute_object")
	})
	return err
}

// GetAttributeObject is a representation of the C type g_file_info_get_attribute_object.
func (recv *FileInfo) GetAttributeObject(attribute string) *gobject.Object {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(attribute)

	var ret gi.Argument

	err := fileInfoGetAttributeObjectFunction_Set()
	if err == nil {
		ret = fileInfoGetAttributeObjectFunction.Invoke(inArgs[:], nil)
	}

	retGo := gobject.ObjectNewFromNative(ret.Pointer())

	return retGo
}

var fileInfoGetAttributeStatusFunction *gi.Function
var fileInfoGetAttributeStatusFunction_Once sync.Once

func fileInfoGetAttributeStatusFunction_Set() error {
	var err error
	fileInfoGetAttributeStatusFunction_Once.Do(func() {
		err = fileInfoObject_Set()
		if err != nil {
			return
		}
		fileInfoGetAttributeStatusFunction, err = fileInfoObject.InvokerNew("get_attribute_status")
	})
	return err
}

// GetAttributeStatus is a representation of the C type g_file_info_get_attribute_status.
func (recv *FileInfo) GetAttributeStatus(attribute string) FileAttributeStatus {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(attribute)

	var ret gi.Argument

	err := fileInfoGetAttributeStatusFunction_Set()
	if err == nil {
		ret = fileInfoGetAttributeStatusFunction.Invoke(inArgs[:], nil)
	}

	retGo := FileAttributeStatus(ret.Int32())

	return retGo
}

var fileInfoGetAttributeStringFunction *gi.Function
var fileInfoGetAttributeStringFunction_Once sync.Once

func fileInfoGetAttributeStringFunction_Set() error {
	var err error
	fileInfoGetAttributeStringFunction_Once.Do(func() {
		err = fileInfoObject_Set()
		if err != nil {
			return
		}
		fileInfoGetAttributeStringFunction, err = fileInfoObject.InvokerNew("get_attribute_string")
	})
	return err
}

// GetAttributeString is a representation of the C type g_file_info_get_attribute_string.
func (recv *FileInfo) GetAttributeString(attribute string) string {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(attribute)

	var ret gi.Argument

	err := fileInfoGetAttributeStringFunction_Set()
	if err == nil {
		ret = fileInfoGetAttributeStringFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var fileInfoGetAttributeStringvFunction *gi.Function
var fileInfoGetAttributeStringvFunction_Once sync.Once

func fileInfoGetAttributeStringvFunction_Set() error {
	var err error
	fileInfoGetAttributeStringvFunction_Once.Do(func() {
		err = fileInfoObject_Set()
		if err != nil {
			return
		}
		fileInfoGetAttributeStringvFunction, err = fileInfoObject.InvokerNew("get_attribute_stringv")
	})
	return err
}

// GetAttributeStringv is a representation of the C type g_file_info_get_attribute_stringv.
func (recv *FileInfo) GetAttributeStringv(attribute string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(attribute)

	err := fileInfoGetAttributeStringvFunction_Set()
	if err == nil {
		fileInfoGetAttributeStringvFunction.Invoke(inArgs[:], nil)
	}

	return
}

var fileInfoGetAttributeTypeFunction *gi.Function
var fileInfoGetAttributeTypeFunction_Once sync.Once

func fileInfoGetAttributeTypeFunction_Set() error {
	var err error
	fileInfoGetAttributeTypeFunction_Once.Do(func() {
		err = fileInfoObject_Set()
		if err != nil {
			return
		}
		fileInfoGetAttributeTypeFunction, err = fileInfoObject.InvokerNew("get_attribute_type")
	})
	return err
}

// GetAttributeType is a representation of the C type g_file_info_get_attribute_type.
func (recv *FileInfo) GetAttributeType(attribute string) FileAttributeType {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(attribute)

	var ret gi.Argument

	err := fileInfoGetAttributeTypeFunction_Set()
	if err == nil {
		ret = fileInfoGetAttributeTypeFunction.Invoke(inArgs[:], nil)
	}

	retGo := FileAttributeType(ret.Int32())

	return retGo
}

var fileInfoGetAttributeUint32Function *gi.Function
var fileInfoGetAttributeUint32Function_Once sync.Once

func fileInfoGetAttributeUint32Function_Set() error {
	var err error
	fileInfoGetAttributeUint32Function_Once.Do(func() {
		err = fileInfoObject_Set()
		if err != nil {
			return
		}
		fileInfoGetAttributeUint32Function, err = fileInfoObject.InvokerNew("get_attribute_uint32")
	})
	return err
}

// GetAttributeUint32 is a representation of the C type g_file_info_get_attribute_uint32.
func (recv *FileInfo) GetAttributeUint32(attribute string) uint32 {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(attribute)

	var ret gi.Argument

	err := fileInfoGetAttributeUint32Function_Set()
	if err == nil {
		ret = fileInfoGetAttributeUint32Function.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo
}

var fileInfoGetAttributeUint64Function *gi.Function
var fileInfoGetAttributeUint64Function_Once sync.Once

func fileInfoGetAttributeUint64Function_Set() error {
	var err error
	fileInfoGetAttributeUint64Function_Once.Do(func() {
		err = fileInfoObject_Set()
		if err != nil {
			return
		}
		fileInfoGetAttributeUint64Function, err = fileInfoObject.InvokerNew("get_attribute_uint64")
	})
	return err
}

// GetAttributeUint64 is a representation of the C type g_file_info_get_attribute_uint64.
func (recv *FileInfo) GetAttributeUint64(attribute string) uint64 {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(attribute)

	var ret gi.Argument

	err := fileInfoGetAttributeUint64Function_Set()
	if err == nil {
		ret = fileInfoGetAttributeUint64Function.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint64()

	return retGo
}

var fileInfoGetContentTypeFunction *gi.Function
var fileInfoGetContentTypeFunction_Once sync.Once

func fileInfoGetContentTypeFunction_Set() error {
	var err error
	fileInfoGetContentTypeFunction_Once.Do(func() {
		err = fileInfoObject_Set()
		if err != nil {
			return
		}
		fileInfoGetContentTypeFunction, err = fileInfoObject.InvokerNew("get_content_type")
	})
	return err
}

// GetContentType is a representation of the C type g_file_info_get_content_type.
func (recv *FileInfo) GetContentType() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := fileInfoGetContentTypeFunction_Set()
	if err == nil {
		ret = fileInfoGetContentTypeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'g_file_info_get_deletion_date' : return type 'GLib.DateTime' not supported

var fileInfoGetDisplayNameFunction *gi.Function
var fileInfoGetDisplayNameFunction_Once sync.Once

func fileInfoGetDisplayNameFunction_Set() error {
	var err error
	fileInfoGetDisplayNameFunction_Once.Do(func() {
		err = fileInfoObject_Set()
		if err != nil {
			return
		}
		fileInfoGetDisplayNameFunction, err = fileInfoObject.InvokerNew("get_display_name")
	})
	return err
}

// GetDisplayName is a representation of the C type g_file_info_get_display_name.
func (recv *FileInfo) GetDisplayName() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := fileInfoGetDisplayNameFunction_Set()
	if err == nil {
		ret = fileInfoGetDisplayNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var fileInfoGetEditNameFunction *gi.Function
var fileInfoGetEditNameFunction_Once sync.Once

func fileInfoGetEditNameFunction_Set() error {
	var err error
	fileInfoGetEditNameFunction_Once.Do(func() {
		err = fileInfoObject_Set()
		if err != nil {
			return
		}
		fileInfoGetEditNameFunction, err = fileInfoObject.InvokerNew("get_edit_name")
	})
	return err
}

// GetEditName is a representation of the C type g_file_info_get_edit_name.
func (recv *FileInfo) GetEditName() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := fileInfoGetEditNameFunction_Set()
	if err == nil {
		ret = fileInfoGetEditNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var fileInfoGetEtagFunction *gi.Function
var fileInfoGetEtagFunction_Once sync.Once

func fileInfoGetEtagFunction_Set() error {
	var err error
	fileInfoGetEtagFunction_Once.Do(func() {
		err = fileInfoObject_Set()
		if err != nil {
			return
		}
		fileInfoGetEtagFunction, err = fileInfoObject.InvokerNew("get_etag")
	})
	return err
}

// GetEtag is a representation of the C type g_file_info_get_etag.
func (recv *FileInfo) GetEtag() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := fileInfoGetEtagFunction_Set()
	if err == nil {
		ret = fileInfoGetEtagFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var fileInfoGetFileTypeFunction *gi.Function
var fileInfoGetFileTypeFunction_Once sync.Once

func fileInfoGetFileTypeFunction_Set() error {
	var err error
	fileInfoGetFileTypeFunction_Once.Do(func() {
		err = fileInfoObject_Set()
		if err != nil {
			return
		}
		fileInfoGetFileTypeFunction, err = fileInfoObject.InvokerNew("get_file_type")
	})
	return err
}

// GetFileType is a representation of the C type g_file_info_get_file_type.
func (recv *FileInfo) GetFileType() FileType {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := fileInfoGetFileTypeFunction_Set()
	if err == nil {
		ret = fileInfoGetFileTypeFunction.Invoke(inArgs[:], nil)
	}

	retGo := FileType(ret.Int32())

	return retGo
}

// UNSUPPORTED : C value 'g_file_info_get_icon' : return type 'Icon' not supported

var fileInfoGetIsBackupFunction *gi.Function
var fileInfoGetIsBackupFunction_Once sync.Once

func fileInfoGetIsBackupFunction_Set() error {
	var err error
	fileInfoGetIsBackupFunction_Once.Do(func() {
		err = fileInfoObject_Set()
		if err != nil {
			return
		}
		fileInfoGetIsBackupFunction, err = fileInfoObject.InvokerNew("get_is_backup")
	})
	return err
}

// GetIsBackup is a representation of the C type g_file_info_get_is_backup.
func (recv *FileInfo) GetIsBackup() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := fileInfoGetIsBackupFunction_Set()
	if err == nil {
		ret = fileInfoGetIsBackupFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var fileInfoGetIsHiddenFunction *gi.Function
var fileInfoGetIsHiddenFunction_Once sync.Once

func fileInfoGetIsHiddenFunction_Set() error {
	var err error
	fileInfoGetIsHiddenFunction_Once.Do(func() {
		err = fileInfoObject_Set()
		if err != nil {
			return
		}
		fileInfoGetIsHiddenFunction, err = fileInfoObject.InvokerNew("get_is_hidden")
	})
	return err
}

// GetIsHidden is a representation of the C type g_file_info_get_is_hidden.
func (recv *FileInfo) GetIsHidden() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := fileInfoGetIsHiddenFunction_Set()
	if err == nil {
		ret = fileInfoGetIsHiddenFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var fileInfoGetIsSymlinkFunction *gi.Function
var fileInfoGetIsSymlinkFunction_Once sync.Once

func fileInfoGetIsSymlinkFunction_Set() error {
	var err error
	fileInfoGetIsSymlinkFunction_Once.Do(func() {
		err = fileInfoObject_Set()
		if err != nil {
			return
		}
		fileInfoGetIsSymlinkFunction, err = fileInfoObject.InvokerNew("get_is_symlink")
	})
	return err
}

// GetIsSymlink is a representation of the C type g_file_info_get_is_symlink.
func (recv *FileInfo) GetIsSymlink() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := fileInfoGetIsSymlinkFunction_Set()
	if err == nil {
		ret = fileInfoGetIsSymlinkFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'g_file_info_get_modification_date_time' : return type 'GLib.DateTime' not supported

// UNSUPPORTED : C value 'g_file_info_get_modification_time' : parameter 'result' of type 'GLib.TimeVal' not supported

var fileInfoGetNameFunction *gi.Function
var fileInfoGetNameFunction_Once sync.Once

func fileInfoGetNameFunction_Set() error {
	var err error
	fileInfoGetNameFunction_Once.Do(func() {
		err = fileInfoObject_Set()
		if err != nil {
			return
		}
		fileInfoGetNameFunction, err = fileInfoObject.InvokerNew("get_name")
	})
	return err
}

// GetName is a representation of the C type g_file_info_get_name.
func (recv *FileInfo) GetName() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := fileInfoGetNameFunction_Set()
	if err == nil {
		ret = fileInfoGetNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var fileInfoGetSizeFunction *gi.Function
var fileInfoGetSizeFunction_Once sync.Once

func fileInfoGetSizeFunction_Set() error {
	var err error
	fileInfoGetSizeFunction_Once.Do(func() {
		err = fileInfoObject_Set()
		if err != nil {
			return
		}
		fileInfoGetSizeFunction, err = fileInfoObject.InvokerNew("get_size")
	})
	return err
}

// GetSize is a representation of the C type g_file_info_get_size.
func (recv *FileInfo) GetSize() int64 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := fileInfoGetSizeFunction_Set()
	if err == nil {
		ret = fileInfoGetSizeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int64()

	return retGo
}

var fileInfoGetSortOrderFunction *gi.Function
var fileInfoGetSortOrderFunction_Once sync.Once

func fileInfoGetSortOrderFunction_Set() error {
	var err error
	fileInfoGetSortOrderFunction_Once.Do(func() {
		err = fileInfoObject_Set()
		if err != nil {
			return
		}
		fileInfoGetSortOrderFunction, err = fileInfoObject.InvokerNew("get_sort_order")
	})
	return err
}

// GetSortOrder is a representation of the C type g_file_info_get_sort_order.
func (recv *FileInfo) GetSortOrder() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := fileInfoGetSortOrderFunction_Set()
	if err == nil {
		ret = fileInfoGetSortOrderFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

// UNSUPPORTED : C value 'g_file_info_get_symbolic_icon' : return type 'Icon' not supported

var fileInfoGetSymlinkTargetFunction *gi.Function
var fileInfoGetSymlinkTargetFunction_Once sync.Once

func fileInfoGetSymlinkTargetFunction_Set() error {
	var err error
	fileInfoGetSymlinkTargetFunction_Once.Do(func() {
		err = fileInfoObject_Set()
		if err != nil {
			return
		}
		fileInfoGetSymlinkTargetFunction, err = fileInfoObject.InvokerNew("get_symlink_target")
	})
	return err
}

// GetSymlinkTarget is a representation of the C type g_file_info_get_symlink_target.
func (recv *FileInfo) GetSymlinkTarget() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := fileInfoGetSymlinkTargetFunction_Set()
	if err == nil {
		ret = fileInfoGetSymlinkTargetFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var fileInfoHasAttributeFunction *gi.Function
var fileInfoHasAttributeFunction_Once sync.Once

func fileInfoHasAttributeFunction_Set() error {
	var err error
	fileInfoHasAttributeFunction_Once.Do(func() {
		err = fileInfoObject_Set()
		if err != nil {
			return
		}
		fileInfoHasAttributeFunction, err = fileInfoObject.InvokerNew("has_attribute")
	})
	return err
}

// HasAttribute is a representation of the C type g_file_info_has_attribute.
func (recv *FileInfo) HasAttribute(attribute string) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(attribute)

	var ret gi.Argument

	err := fileInfoHasAttributeFunction_Set()
	if err == nil {
		ret = fileInfoHasAttributeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var fileInfoHasNamespaceFunction *gi.Function
var fileInfoHasNamespaceFunction_Once sync.Once

func fileInfoHasNamespaceFunction_Set() error {
	var err error
	fileInfoHasNamespaceFunction_Once.Do(func() {
		err = fileInfoObject_Set()
		if err != nil {
			return
		}
		fileInfoHasNamespaceFunction, err = fileInfoObject.InvokerNew("has_namespace")
	})
	return err
}

// HasNamespace is a representation of the C type g_file_info_has_namespace.
func (recv *FileInfo) HasNamespace(nameSpace string) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(nameSpace)

	var ret gi.Argument

	err := fileInfoHasNamespaceFunction_Set()
	if err == nil {
		ret = fileInfoHasNamespaceFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var fileInfoListAttributesFunction *gi.Function
var fileInfoListAttributesFunction_Once sync.Once

func fileInfoListAttributesFunction_Set() error {
	var err error
	fileInfoListAttributesFunction_Once.Do(func() {
		err = fileInfoObject_Set()
		if err != nil {
			return
		}
		fileInfoListAttributesFunction, err = fileInfoObject.InvokerNew("list_attributes")
	})
	return err
}

// ListAttributes is a representation of the C type g_file_info_list_attributes.
func (recv *FileInfo) ListAttributes(nameSpace string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(nameSpace)

	err := fileInfoListAttributesFunction_Set()
	if err == nil {
		fileInfoListAttributesFunction.Invoke(inArgs[:], nil)
	}

	return
}

var fileInfoRemoveAttributeFunction *gi.Function
var fileInfoRemoveAttributeFunction_Once sync.Once

func fileInfoRemoveAttributeFunction_Set() error {
	var err error
	fileInfoRemoveAttributeFunction_Once.Do(func() {
		err = fileInfoObject_Set()
		if err != nil {
			return
		}
		fileInfoRemoveAttributeFunction, err = fileInfoObject.InvokerNew("remove_attribute")
	})
	return err
}

// RemoveAttribute is a representation of the C type g_file_info_remove_attribute.
func (recv *FileInfo) RemoveAttribute(attribute string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(attribute)

	err := fileInfoRemoveAttributeFunction_Set()
	if err == nil {
		fileInfoRemoveAttributeFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'g_file_info_set_attribute' : parameter 'value_p' of type 'gpointer' not supported

var fileInfoSetAttributeBooleanFunction *gi.Function
var fileInfoSetAttributeBooleanFunction_Once sync.Once

func fileInfoSetAttributeBooleanFunction_Set() error {
	var err error
	fileInfoSetAttributeBooleanFunction_Once.Do(func() {
		err = fileInfoObject_Set()
		if err != nil {
			return
		}
		fileInfoSetAttributeBooleanFunction, err = fileInfoObject.InvokerNew("set_attribute_boolean")
	})
	return err
}

// SetAttributeBoolean is a representation of the C type g_file_info_set_attribute_boolean.
func (recv *FileInfo) SetAttributeBoolean(attribute string, attrValue bool) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(attribute)
	inArgs[2].SetBoolean(attrValue)

	err := fileInfoSetAttributeBooleanFunction_Set()
	if err == nil {
		fileInfoSetAttributeBooleanFunction.Invoke(inArgs[:], nil)
	}

	return
}

var fileInfoSetAttributeByteStringFunction *gi.Function
var fileInfoSetAttributeByteStringFunction_Once sync.Once

func fileInfoSetAttributeByteStringFunction_Set() error {
	var err error
	fileInfoSetAttributeByteStringFunction_Once.Do(func() {
		err = fileInfoObject_Set()
		if err != nil {
			return
		}
		fileInfoSetAttributeByteStringFunction, err = fileInfoObject.InvokerNew("set_attribute_byte_string")
	})
	return err
}

// SetAttributeByteString is a representation of the C type g_file_info_set_attribute_byte_string.
func (recv *FileInfo) SetAttributeByteString(attribute string, attrValue string) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(attribute)
	inArgs[2].SetString(attrValue)

	err := fileInfoSetAttributeByteStringFunction_Set()
	if err == nil {
		fileInfoSetAttributeByteStringFunction.Invoke(inArgs[:], nil)
	}

	return
}

var fileInfoSetAttributeInt32Function *gi.Function
var fileInfoSetAttributeInt32Function_Once sync.Once

func fileInfoSetAttributeInt32Function_Set() error {
	var err error
	fileInfoSetAttributeInt32Function_Once.Do(func() {
		err = fileInfoObject_Set()
		if err != nil {
			return
		}
		fileInfoSetAttributeInt32Function, err = fileInfoObject.InvokerNew("set_attribute_int32")
	})
	return err
}

// SetAttributeInt32 is a representation of the C type g_file_info_set_attribute_int32.
func (recv *FileInfo) SetAttributeInt32(attribute string, attrValue int32) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(attribute)
	inArgs[2].SetInt32(attrValue)

	err := fileInfoSetAttributeInt32Function_Set()
	if err == nil {
		fileInfoSetAttributeInt32Function.Invoke(inArgs[:], nil)
	}

	return
}

var fileInfoSetAttributeInt64Function *gi.Function
var fileInfoSetAttributeInt64Function_Once sync.Once

func fileInfoSetAttributeInt64Function_Set() error {
	var err error
	fileInfoSetAttributeInt64Function_Once.Do(func() {
		err = fileInfoObject_Set()
		if err != nil {
			return
		}
		fileInfoSetAttributeInt64Function, err = fileInfoObject.InvokerNew("set_attribute_int64")
	})
	return err
}

// SetAttributeInt64 is a representation of the C type g_file_info_set_attribute_int64.
func (recv *FileInfo) SetAttributeInt64(attribute string, attrValue int64) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(attribute)
	inArgs[2].SetInt64(attrValue)

	err := fileInfoSetAttributeInt64Function_Set()
	if err == nil {
		fileInfoSetAttributeInt64Function.Invoke(inArgs[:], nil)
	}

	return
}

var fileInfoSetAttributeMaskFunction *gi.Function
var fileInfoSetAttributeMaskFunction_Once sync.Once

func fileInfoSetAttributeMaskFunction_Set() error {
	var err error
	fileInfoSetAttributeMaskFunction_Once.Do(func() {
		err = fileInfoObject_Set()
		if err != nil {
			return
		}
		fileInfoSetAttributeMaskFunction, err = fileInfoObject.InvokerNew("set_attribute_mask")
	})
	return err
}

// SetAttributeMask is a representation of the C type g_file_info_set_attribute_mask.
func (recv *FileInfo) SetAttributeMask(mask *FileAttributeMatcher) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(mask.Native())

	err := fileInfoSetAttributeMaskFunction_Set()
	if err == nil {
		fileInfoSetAttributeMaskFunction.Invoke(inArgs[:], nil)
	}

	return
}

var fileInfoSetAttributeObjectFunction *gi.Function
var fileInfoSetAttributeObjectFunction_Once sync.Once

func fileInfoSetAttributeObjectFunction_Set() error {
	var err error
	fileInfoSetAttributeObjectFunction_Once.Do(func() {
		err = fileInfoObject_Set()
		if err != nil {
			return
		}
		fileInfoSetAttributeObjectFunction, err = fileInfoObject.InvokerNew("set_attribute_object")
	})
	return err
}

// SetAttributeObject is a representation of the C type g_file_info_set_attribute_object.
func (recv *FileInfo) SetAttributeObject(attribute string, attrValue *gobject.Object) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(attribute)
	inArgs[2].SetPointer(attrValue.Native())

	err := fileInfoSetAttributeObjectFunction_Set()
	if err == nil {
		fileInfoSetAttributeObjectFunction.Invoke(inArgs[:], nil)
	}

	return
}

var fileInfoSetAttributeStatusFunction *gi.Function
var fileInfoSetAttributeStatusFunction_Once sync.Once

func fileInfoSetAttributeStatusFunction_Set() error {
	var err error
	fileInfoSetAttributeStatusFunction_Once.Do(func() {
		err = fileInfoObject_Set()
		if err != nil {
			return
		}
		fileInfoSetAttributeStatusFunction, err = fileInfoObject.InvokerNew("set_attribute_status")
	})
	return err
}

// SetAttributeStatus is a representation of the C type g_file_info_set_attribute_status.
func (recv *FileInfo) SetAttributeStatus(attribute string, status FileAttributeStatus) bool {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(attribute)
	inArgs[2].SetInt32(int32(status))

	var ret gi.Argument

	err := fileInfoSetAttributeStatusFunction_Set()
	if err == nil {
		ret = fileInfoSetAttributeStatusFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var fileInfoSetAttributeStringFunction *gi.Function
var fileInfoSetAttributeStringFunction_Once sync.Once

func fileInfoSetAttributeStringFunction_Set() error {
	var err error
	fileInfoSetAttributeStringFunction_Once.Do(func() {
		err = fileInfoObject_Set()
		if err != nil {
			return
		}
		fileInfoSetAttributeStringFunction, err = fileInfoObject.InvokerNew("set_attribute_string")
	})
	return err
}

// SetAttributeString is a representation of the C type g_file_info_set_attribute_string.
func (recv *FileInfo) SetAttributeString(attribute string, attrValue string) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(attribute)
	inArgs[2].SetString(attrValue)

	err := fileInfoSetAttributeStringFunction_Set()
	if err == nil {
		fileInfoSetAttributeStringFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'g_file_info_set_attribute_stringv' : parameter 'attr_value' of type 'nil' not supported

var fileInfoSetAttributeUint32Function *gi.Function
var fileInfoSetAttributeUint32Function_Once sync.Once

func fileInfoSetAttributeUint32Function_Set() error {
	var err error
	fileInfoSetAttributeUint32Function_Once.Do(func() {
		err = fileInfoObject_Set()
		if err != nil {
			return
		}
		fileInfoSetAttributeUint32Function, err = fileInfoObject.InvokerNew("set_attribute_uint32")
	})
	return err
}

// SetAttributeUint32 is a representation of the C type g_file_info_set_attribute_uint32.
func (recv *FileInfo) SetAttributeUint32(attribute string, attrValue uint32) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(attribute)
	inArgs[2].SetUint32(attrValue)

	err := fileInfoSetAttributeUint32Function_Set()
	if err == nil {
		fileInfoSetAttributeUint32Function.Invoke(inArgs[:], nil)
	}

	return
}

var fileInfoSetAttributeUint64Function *gi.Function
var fileInfoSetAttributeUint64Function_Once sync.Once

func fileInfoSetAttributeUint64Function_Set() error {
	var err error
	fileInfoSetAttributeUint64Function_Once.Do(func() {
		err = fileInfoObject_Set()
		if err != nil {
			return
		}
		fileInfoSetAttributeUint64Function, err = fileInfoObject.InvokerNew("set_attribute_uint64")
	})
	return err
}

// SetAttributeUint64 is a representation of the C type g_file_info_set_attribute_uint64.
func (recv *FileInfo) SetAttributeUint64(attribute string, attrValue uint64) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(attribute)
	inArgs[2].SetUint64(attrValue)

	err := fileInfoSetAttributeUint64Function_Set()
	if err == nil {
		fileInfoSetAttributeUint64Function.Invoke(inArgs[:], nil)
	}

	return
}

var fileInfoSetContentTypeFunction *gi.Function
var fileInfoSetContentTypeFunction_Once sync.Once

func fileInfoSetContentTypeFunction_Set() error {
	var err error
	fileInfoSetContentTypeFunction_Once.Do(func() {
		err = fileInfoObject_Set()
		if err != nil {
			return
		}
		fileInfoSetContentTypeFunction, err = fileInfoObject.InvokerNew("set_content_type")
	})
	return err
}

// SetContentType is a representation of the C type g_file_info_set_content_type.
func (recv *FileInfo) SetContentType(contentType string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(contentType)

	err := fileInfoSetContentTypeFunction_Set()
	if err == nil {
		fileInfoSetContentTypeFunction.Invoke(inArgs[:], nil)
	}

	return
}

var fileInfoSetDisplayNameFunction *gi.Function
var fileInfoSetDisplayNameFunction_Once sync.Once

func fileInfoSetDisplayNameFunction_Set() error {
	var err error
	fileInfoSetDisplayNameFunction_Once.Do(func() {
		err = fileInfoObject_Set()
		if err != nil {
			return
		}
		fileInfoSetDisplayNameFunction, err = fileInfoObject.InvokerNew("set_display_name")
	})
	return err
}

// SetDisplayName is a representation of the C type g_file_info_set_display_name.
func (recv *FileInfo) SetDisplayName(displayName string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(displayName)

	err := fileInfoSetDisplayNameFunction_Set()
	if err == nil {
		fileInfoSetDisplayNameFunction.Invoke(inArgs[:], nil)
	}

	return
}

var fileInfoSetEditNameFunction *gi.Function
var fileInfoSetEditNameFunction_Once sync.Once

func fileInfoSetEditNameFunction_Set() error {
	var err error
	fileInfoSetEditNameFunction_Once.Do(func() {
		err = fileInfoObject_Set()
		if err != nil {
			return
		}
		fileInfoSetEditNameFunction, err = fileInfoObject.InvokerNew("set_edit_name")
	})
	return err
}

// SetEditName is a representation of the C type g_file_info_set_edit_name.
func (recv *FileInfo) SetEditName(editName string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(editName)

	err := fileInfoSetEditNameFunction_Set()
	if err == nil {
		fileInfoSetEditNameFunction.Invoke(inArgs[:], nil)
	}

	return
}

var fileInfoSetFileTypeFunction *gi.Function
var fileInfoSetFileTypeFunction_Once sync.Once

func fileInfoSetFileTypeFunction_Set() error {
	var err error
	fileInfoSetFileTypeFunction_Once.Do(func() {
		err = fileInfoObject_Set()
		if err != nil {
			return
		}
		fileInfoSetFileTypeFunction, err = fileInfoObject.InvokerNew("set_file_type")
	})
	return err
}

// SetFileType is a representation of the C type g_file_info_set_file_type.
func (recv *FileInfo) SetFileType(type_ FileType) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(int32(type_))

	err := fileInfoSetFileTypeFunction_Set()
	if err == nil {
		fileInfoSetFileTypeFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'g_file_info_set_icon' : parameter 'icon' of type 'Icon' not supported

var fileInfoSetIsHiddenFunction *gi.Function
var fileInfoSetIsHiddenFunction_Once sync.Once

func fileInfoSetIsHiddenFunction_Set() error {
	var err error
	fileInfoSetIsHiddenFunction_Once.Do(func() {
		err = fileInfoObject_Set()
		if err != nil {
			return
		}
		fileInfoSetIsHiddenFunction, err = fileInfoObject.InvokerNew("set_is_hidden")
	})
	return err
}

// SetIsHidden is a representation of the C type g_file_info_set_is_hidden.
func (recv *FileInfo) SetIsHidden(isHidden bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetBoolean(isHidden)

	err := fileInfoSetIsHiddenFunction_Set()
	if err == nil {
		fileInfoSetIsHiddenFunction.Invoke(inArgs[:], nil)
	}

	return
}

var fileInfoSetIsSymlinkFunction *gi.Function
var fileInfoSetIsSymlinkFunction_Once sync.Once

func fileInfoSetIsSymlinkFunction_Set() error {
	var err error
	fileInfoSetIsSymlinkFunction_Once.Do(func() {
		err = fileInfoObject_Set()
		if err != nil {
			return
		}
		fileInfoSetIsSymlinkFunction, err = fileInfoObject.InvokerNew("set_is_symlink")
	})
	return err
}

// SetIsSymlink is a representation of the C type g_file_info_set_is_symlink.
func (recv *FileInfo) SetIsSymlink(isSymlink bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetBoolean(isSymlink)

	err := fileInfoSetIsSymlinkFunction_Set()
	if err == nil {
		fileInfoSetIsSymlinkFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'g_file_info_set_modification_date_time' : parameter 'mtime' of type 'GLib.DateTime' not supported

// UNSUPPORTED : C value 'g_file_info_set_modification_time' : parameter 'mtime' of type 'GLib.TimeVal' not supported

var fileInfoSetNameFunction *gi.Function
var fileInfoSetNameFunction_Once sync.Once

func fileInfoSetNameFunction_Set() error {
	var err error
	fileInfoSetNameFunction_Once.Do(func() {
		err = fileInfoObject_Set()
		if err != nil {
			return
		}
		fileInfoSetNameFunction, err = fileInfoObject.InvokerNew("set_name")
	})
	return err
}

// SetName is a representation of the C type g_file_info_set_name.
func (recv *FileInfo) SetName(name string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(name)

	err := fileInfoSetNameFunction_Set()
	if err == nil {
		fileInfoSetNameFunction.Invoke(inArgs[:], nil)
	}

	return
}

var fileInfoSetSizeFunction *gi.Function
var fileInfoSetSizeFunction_Once sync.Once

func fileInfoSetSizeFunction_Set() error {
	var err error
	fileInfoSetSizeFunction_Once.Do(func() {
		err = fileInfoObject_Set()
		if err != nil {
			return
		}
		fileInfoSetSizeFunction, err = fileInfoObject.InvokerNew("set_size")
	})
	return err
}

// SetSize is a representation of the C type g_file_info_set_size.
func (recv *FileInfo) SetSize(size int64) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt64(size)

	err := fileInfoSetSizeFunction_Set()
	if err == nil {
		fileInfoSetSizeFunction.Invoke(inArgs[:], nil)
	}

	return
}

var fileInfoSetSortOrderFunction *gi.Function
var fileInfoSetSortOrderFunction_Once sync.Once

func fileInfoSetSortOrderFunction_Set() error {
	var err error
	fileInfoSetSortOrderFunction_Once.Do(func() {
		err = fileInfoObject_Set()
		if err != nil {
			return
		}
		fileInfoSetSortOrderFunction, err = fileInfoObject.InvokerNew("set_sort_order")
	})
	return err
}

// SetSortOrder is a representation of the C type g_file_info_set_sort_order.
func (recv *FileInfo) SetSortOrder(sortOrder int32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(sortOrder)

	err := fileInfoSetSortOrderFunction_Set()
	if err == nil {
		fileInfoSetSortOrderFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'g_file_info_set_symbolic_icon' : parameter 'icon' of type 'Icon' not supported

var fileInfoSetSymlinkTargetFunction *gi.Function
var fileInfoSetSymlinkTargetFunction_Once sync.Once

func fileInfoSetSymlinkTargetFunction_Set() error {
	var err error
	fileInfoSetSymlinkTargetFunction_Once.Do(func() {
		err = fileInfoObject_Set()
		if err != nil {
			return
		}
		fileInfoSetSymlinkTargetFunction, err = fileInfoObject.InvokerNew("set_symlink_target")
	})
	return err
}

// SetSymlinkTarget is a representation of the C type g_file_info_set_symlink_target.
func (recv *FileInfo) SetSymlinkTarget(symlinkTarget string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(symlinkTarget)

	err := fileInfoSetSymlinkTargetFunction_Set()
	if err == nil {
		fileInfoSetSymlinkTargetFunction.Invoke(inArgs[:], nil)
	}

	return
}

var fileInfoUnsetAttributeMaskFunction *gi.Function
var fileInfoUnsetAttributeMaskFunction_Once sync.Once

func fileInfoUnsetAttributeMaskFunction_Set() error {
	var err error
	fileInfoUnsetAttributeMaskFunction_Once.Do(func() {
		err = fileInfoObject_Set()
		if err != nil {
			return
		}
		fileInfoUnsetAttributeMaskFunction, err = fileInfoObject.InvokerNew("unset_attribute_mask")
	})
	return err
}

// UnsetAttributeMask is a representation of the C type g_file_info_unset_attribute_mask.
func (recv *FileInfo) UnsetAttributeMask() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := fileInfoUnsetAttributeMaskFunction_Set()
	if err == nil {
		fileInfoUnsetAttributeMaskFunction.Invoke(inArgs[:], nil)
	}

	return
}

var fileInputStreamObject *gi.Object
var fileInputStreamObject_Once sync.Once

func fileInputStreamObject_Set() error {
	var err error
	fileInputStreamObject_Once.Do(func() {
		fileInputStreamObject, err = gi.ObjectNew("Gio", "FileInputStream")
	})
	return err
}

type FileInputStream struct {
	native unsafe.Pointer
}

func FileInputStreamNewFromNative(native unsafe.Pointer) *FileInputStream {
	instance := &FileInputStream{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// InputStream upcasts to *InputStream
func (recv *FileInputStream) InputStream() *InputStream {
	return InputStreamNewFromNative(recv.Native())
}

// Object upcasts to *Object
func (recv *FileInputStream) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToFileInputStream down casts any arbitrary Object to FileInputStream.
Exercise care, as this is a potentially dangerous function
if the Object is not a FileInputStream.
*/
func CastToFileInputStream(object *gobject.Object) *FileInputStream {
	return FileInputStreamNewFromNative(object.Native())
}

// Equals compares this FileInputStream with another FileInputStream, and returns true if they represent the same GObject.
func (recv *FileInputStream) Equals(other *FileInputStream) bool {
	return other.Native() == recv.Native()
}

func (recv *FileInputStream) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *FileInputStream) FieldParentInstance() *InputStream {
	argValue := gi.ObjectFieldGet(fileInputStreamObject, recv.Native(), "parent_instance")
	value := InputStreamNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *FileInputStream) SetFieldParentInstance(value *InputStream) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(fileInputStreamObject, recv.Native(), "parent_instance", argValue)
}

var fileInputStreamQueryInfoFunction *gi.Function
var fileInputStreamQueryInfoFunction_Once sync.Once

func fileInputStreamQueryInfoFunction_Set() error {
	var err error
	fileInputStreamQueryInfoFunction_Once.Do(func() {
		err = fileInputStreamObject_Set()
		if err != nil {
			return
		}
		fileInputStreamQueryInfoFunction, err = fileInputStreamObject.InvokerNew("query_info")
	})
	return err
}

// QueryInfo is a representation of the C type g_file_input_stream_query_info.
func (recv *FileInputStream) QueryInfo(attributes string, cancellable *Cancellable) *FileInfo {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(attributes)
	inArgs[2].SetPointer(cancellable.Native())

	var ret gi.Argument

	err := fileInputStreamQueryInfoFunction_Set()
	if err == nil {
		ret = fileInputStreamQueryInfoFunction.Invoke(inArgs[:], nil)
	}

	retGo := FileInfoNewFromNative(ret.Pointer())

	return retGo
}

// UNSUPPORTED : C value 'g_file_input_stream_query_info_async' : parameter 'callback' of type 'AsyncReadyCallback' not supported

// UNSUPPORTED : C value 'g_file_input_stream_query_info_finish' : parameter 'result' of type 'AsyncResult' not supported

var fileMonitorObject *gi.Object
var fileMonitorObject_Once sync.Once

func fileMonitorObject_Set() error {
	var err error
	fileMonitorObject_Once.Do(func() {
		fileMonitorObject, err = gi.ObjectNew("Gio", "FileMonitor")
	})
	return err
}

type FileMonitor struct {
	native unsafe.Pointer
}

func FileMonitorNewFromNative(native unsafe.Pointer) *FileMonitor {
	instance := &FileMonitor{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *FileMonitor) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToFileMonitor down casts any arbitrary Object to FileMonitor.
Exercise care, as this is a potentially dangerous function
if the Object is not a FileMonitor.
*/
func CastToFileMonitor(object *gobject.Object) *FileMonitor {
	return FileMonitorNewFromNative(object.Native())
}

// Equals compares this FileMonitor with another FileMonitor, and returns true if they represent the same GObject.
func (recv *FileMonitor) Equals(other *FileMonitor) bool {
	return other.Native() == recv.Native()
}

func (recv *FileMonitor) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *FileMonitor) FieldParentInstance() *gobject.Object {
	argValue := gi.ObjectFieldGet(fileMonitorObject, recv.Native(), "parent_instance")
	value := gobject.ObjectNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *FileMonitor) SetFieldParentInstance(value *gobject.Object) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(fileMonitorObject, recv.Native(), "parent_instance", argValue)
}

var fileMonitorCancelFunction *gi.Function
var fileMonitorCancelFunction_Once sync.Once

func fileMonitorCancelFunction_Set() error {
	var err error
	fileMonitorCancelFunction_Once.Do(func() {
		err = fileMonitorObject_Set()
		if err != nil {
			return
		}
		fileMonitorCancelFunction, err = fileMonitorObject.InvokerNew("cancel")
	})
	return err
}

// Cancel is a representation of the C type g_file_monitor_cancel.
func (recv *FileMonitor) Cancel() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := fileMonitorCancelFunction_Set()
	if err == nil {
		ret = fileMonitorCancelFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'g_file_monitor_emit_event' : parameter 'child' of type 'File' not supported

var fileMonitorIsCancelledFunction *gi.Function
var fileMonitorIsCancelledFunction_Once sync.Once

func fileMonitorIsCancelledFunction_Set() error {
	var err error
	fileMonitorIsCancelledFunction_Once.Do(func() {
		err = fileMonitorObject_Set()
		if err != nil {
			return
		}
		fileMonitorIsCancelledFunction, err = fileMonitorObject.InvokerNew("is_cancelled")
	})
	return err
}

// IsCancelled is a representation of the C type g_file_monitor_is_cancelled.
func (recv *FileMonitor) IsCancelled() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := fileMonitorIsCancelledFunction_Set()
	if err == nil {
		ret = fileMonitorIsCancelledFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var fileMonitorSetRateLimitFunction *gi.Function
var fileMonitorSetRateLimitFunction_Once sync.Once

func fileMonitorSetRateLimitFunction_Set() error {
	var err error
	fileMonitorSetRateLimitFunction_Once.Do(func() {
		err = fileMonitorObject_Set()
		if err != nil {
			return
		}
		fileMonitorSetRateLimitFunction, err = fileMonitorObject.InvokerNew("set_rate_limit")
	})
	return err
}

// SetRateLimit is a representation of the C type g_file_monitor_set_rate_limit.
func (recv *FileMonitor) SetRateLimit(limitMsecs int32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(limitMsecs)

	err := fileMonitorSetRateLimitFunction_Set()
	if err == nil {
		fileMonitorSetRateLimitFunction.Invoke(inArgs[:], nil)
	}

	return
}

var fileOutputStreamObject *gi.Object
var fileOutputStreamObject_Once sync.Once

func fileOutputStreamObject_Set() error {
	var err error
	fileOutputStreamObject_Once.Do(func() {
		fileOutputStreamObject, err = gi.ObjectNew("Gio", "FileOutputStream")
	})
	return err
}

type FileOutputStream struct {
	native unsafe.Pointer
}

func FileOutputStreamNewFromNative(native unsafe.Pointer) *FileOutputStream {
	instance := &FileOutputStream{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// OutputStream upcasts to *OutputStream
func (recv *FileOutputStream) OutputStream() *OutputStream {
	return OutputStreamNewFromNative(recv.Native())
}

// Object upcasts to *Object
func (recv *FileOutputStream) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToFileOutputStream down casts any arbitrary Object to FileOutputStream.
Exercise care, as this is a potentially dangerous function
if the Object is not a FileOutputStream.
*/
func CastToFileOutputStream(object *gobject.Object) *FileOutputStream {
	return FileOutputStreamNewFromNative(object.Native())
}

// Equals compares this FileOutputStream with another FileOutputStream, and returns true if they represent the same GObject.
func (recv *FileOutputStream) Equals(other *FileOutputStream) bool {
	return other.Native() == recv.Native()
}

func (recv *FileOutputStream) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *FileOutputStream) FieldParentInstance() *OutputStream {
	argValue := gi.ObjectFieldGet(fileOutputStreamObject, recv.Native(), "parent_instance")
	value := OutputStreamNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *FileOutputStream) SetFieldParentInstance(value *OutputStream) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(fileOutputStreamObject, recv.Native(), "parent_instance", argValue)
}

var fileOutputStreamGetEtagFunction *gi.Function
var fileOutputStreamGetEtagFunction_Once sync.Once

func fileOutputStreamGetEtagFunction_Set() error {
	var err error
	fileOutputStreamGetEtagFunction_Once.Do(func() {
		err = fileOutputStreamObject_Set()
		if err != nil {
			return
		}
		fileOutputStreamGetEtagFunction, err = fileOutputStreamObject.InvokerNew("get_etag")
	})
	return err
}

// GetEtag is a representation of the C type g_file_output_stream_get_etag.
func (recv *FileOutputStream) GetEtag() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := fileOutputStreamGetEtagFunction_Set()
	if err == nil {
		ret = fileOutputStreamGetEtagFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

var fileOutputStreamQueryInfoFunction *gi.Function
var fileOutputStreamQueryInfoFunction_Once sync.Once

func fileOutputStreamQueryInfoFunction_Set() error {
	var err error
	fileOutputStreamQueryInfoFunction_Once.Do(func() {
		err = fileOutputStreamObject_Set()
		if err != nil {
			return
		}
		fileOutputStreamQueryInfoFunction, err = fileOutputStreamObject.InvokerNew("query_info")
	})
	return err
}

// QueryInfo is a representation of the C type g_file_output_stream_query_info.
func (recv *FileOutputStream) QueryInfo(attributes string, cancellable *Cancellable) *FileInfo {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(attributes)
	inArgs[2].SetPointer(cancellable.Native())

	var ret gi.Argument

	err := fileOutputStreamQueryInfoFunction_Set()
	if err == nil {
		ret = fileOutputStreamQueryInfoFunction.Invoke(inArgs[:], nil)
	}

	retGo := FileInfoNewFromNative(ret.Pointer())

	return retGo
}

// UNSUPPORTED : C value 'g_file_output_stream_query_info_async' : parameter 'callback' of type 'AsyncReadyCallback' not supported

// UNSUPPORTED : C value 'g_file_output_stream_query_info_finish' : parameter 'result' of type 'AsyncResult' not supported

var filenameCompleterObject *gi.Object
var filenameCompleterObject_Once sync.Once

func filenameCompleterObject_Set() error {
	var err error
	filenameCompleterObject_Once.Do(func() {
		filenameCompleterObject, err = gi.ObjectNew("Gio", "FilenameCompleter")
	})
	return err
}

type FilenameCompleter struct {
	native unsafe.Pointer
}

func FilenameCompleterNewFromNative(native unsafe.Pointer) *FilenameCompleter {
	instance := &FilenameCompleter{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *FilenameCompleter) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToFilenameCompleter down casts any arbitrary Object to FilenameCompleter.
Exercise care, as this is a potentially dangerous function
if the Object is not a FilenameCompleter.
*/
func CastToFilenameCompleter(object *gobject.Object) *FilenameCompleter {
	return FilenameCompleterNewFromNative(object.Native())
}

// Equals compares this FilenameCompleter with another FilenameCompleter, and returns true if they represent the same GObject.
func (recv *FilenameCompleter) Equals(other *FilenameCompleter) bool {
	return other.Native() == recv.Native()
}

func (recv *FilenameCompleter) Native() unsafe.Pointer {
	return recv.native
}

var filenameCompleterNewFunction *gi.Function
var filenameCompleterNewFunction_Once sync.Once

func filenameCompleterNewFunction_Set() error {
	var err error
	filenameCompleterNewFunction_Once.Do(func() {
		err = filenameCompleterObject_Set()
		if err != nil {
			return
		}
		filenameCompleterNewFunction, err = filenameCompleterObject.InvokerNew("new")
	})
	return err
}

// FilenameCompleterNew is a representation of the C type g_filename_completer_new.
func FilenameCompleterNew() *FilenameCompleter {

	var ret gi.Argument

	err := filenameCompleterNewFunction_Set()
	if err == nil {
		ret = filenameCompleterNewFunction.Invoke(nil, nil)
	}

	retGo := FilenameCompleterNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

var filenameCompleterGetCompletionSuffixFunction *gi.Function
var filenameCompleterGetCompletionSuffixFunction_Once sync.Once

func filenameCompleterGetCompletionSuffixFunction_Set() error {
	var err error
	filenameCompleterGetCompletionSuffixFunction_Once.Do(func() {
		err = filenameCompleterObject_Set()
		if err != nil {
			return
		}
		filenameCompleterGetCompletionSuffixFunction, err = filenameCompleterObject.InvokerNew("get_completion_suffix")
	})
	return err
}

// GetCompletionSuffix is a representation of the C type g_filename_completer_get_completion_suffix.
func (recv *FilenameCompleter) GetCompletionSuffix(initialText string) string {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(initialText)

	var ret gi.Argument

	err := filenameCompleterGetCompletionSuffixFunction_Set()
	if err == nil {
		ret = filenameCompleterGetCompletionSuffixFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

var filenameCompleterGetCompletionsFunction *gi.Function
var filenameCompleterGetCompletionsFunction_Once sync.Once

func filenameCompleterGetCompletionsFunction_Set() error {
	var err error
	filenameCompleterGetCompletionsFunction_Once.Do(func() {
		err = filenameCompleterObject_Set()
		if err != nil {
			return
		}
		filenameCompleterGetCompletionsFunction, err = filenameCompleterObject.InvokerNew("get_completions")
	})
	return err
}

// GetCompletions is a representation of the C type g_filename_completer_get_completions.
func (recv *FilenameCompleter) GetCompletions(initialText string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(initialText)

	err := filenameCompleterGetCompletionsFunction_Set()
	if err == nil {
		filenameCompleterGetCompletionsFunction.Invoke(inArgs[:], nil)
	}

	return
}

var filenameCompleterSetDirsOnlyFunction *gi.Function
var filenameCompleterSetDirsOnlyFunction_Once sync.Once

func filenameCompleterSetDirsOnlyFunction_Set() error {
	var err error
	filenameCompleterSetDirsOnlyFunction_Once.Do(func() {
		err = filenameCompleterObject_Set()
		if err != nil {
			return
		}
		filenameCompleterSetDirsOnlyFunction, err = filenameCompleterObject.InvokerNew("set_dirs_only")
	})
	return err
}

// SetDirsOnly is a representation of the C type g_filename_completer_set_dirs_only.
func (recv *FilenameCompleter) SetDirsOnly(dirsOnly bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetBoolean(dirsOnly)

	err := filenameCompleterSetDirsOnlyFunction_Set()
	if err == nil {
		filenameCompleterSetDirsOnlyFunction.Invoke(inArgs[:], nil)
	}

	return
}

var filterInputStreamObject *gi.Object
var filterInputStreamObject_Once sync.Once

func filterInputStreamObject_Set() error {
	var err error
	filterInputStreamObject_Once.Do(func() {
		filterInputStreamObject, err = gi.ObjectNew("Gio", "FilterInputStream")
	})
	return err
}

type FilterInputStream struct {
	native unsafe.Pointer
}

func FilterInputStreamNewFromNative(native unsafe.Pointer) *FilterInputStream {
	instance := &FilterInputStream{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// InputStream upcasts to *InputStream
func (recv *FilterInputStream) InputStream() *InputStream {
	return InputStreamNewFromNative(recv.Native())
}

// Object upcasts to *Object
func (recv *FilterInputStream) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToFilterInputStream down casts any arbitrary Object to FilterInputStream.
Exercise care, as this is a potentially dangerous function
if the Object is not a FilterInputStream.
*/
func CastToFilterInputStream(object *gobject.Object) *FilterInputStream {
	return FilterInputStreamNewFromNative(object.Native())
}

// Equals compares this FilterInputStream with another FilterInputStream, and returns true if they represent the same GObject.
func (recv *FilterInputStream) Equals(other *FilterInputStream) bool {
	return other.Native() == recv.Native()
}

func (recv *FilterInputStream) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *FilterInputStream) FieldParentInstance() *InputStream {
	argValue := gi.ObjectFieldGet(filterInputStreamObject, recv.Native(), "parent_instance")
	value := InputStreamNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *FilterInputStream) SetFieldParentInstance(value *InputStream) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(filterInputStreamObject, recv.Native(), "parent_instance", argValue)
}

// FieldBaseStream returns the C field 'base_stream'.
func (recv *FilterInputStream) FieldBaseStream() *InputStream {
	argValue := gi.ObjectFieldGet(filterInputStreamObject, recv.Native(), "base_stream")
	value := InputStreamNewFromNative(argValue.Pointer())
	return value
}

// SetFieldBaseStream sets the value of the C field 'base_stream'.
func (recv *FilterInputStream) SetFieldBaseStream(value *InputStream) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(filterInputStreamObject, recv.Native(), "base_stream", argValue)
}

var filterInputStreamGetBaseStreamFunction *gi.Function
var filterInputStreamGetBaseStreamFunction_Once sync.Once

func filterInputStreamGetBaseStreamFunction_Set() error {
	var err error
	filterInputStreamGetBaseStreamFunction_Once.Do(func() {
		err = filterInputStreamObject_Set()
		if err != nil {
			return
		}
		filterInputStreamGetBaseStreamFunction, err = filterInputStreamObject.InvokerNew("get_base_stream")
	})
	return err
}

// GetBaseStream is a representation of the C type g_filter_input_stream_get_base_stream.
func (recv *FilterInputStream) GetBaseStream() *InputStream {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := filterInputStreamGetBaseStreamFunction_Set()
	if err == nil {
		ret = filterInputStreamGetBaseStreamFunction.Invoke(inArgs[:], nil)
	}

	retGo := InputStreamNewFromNative(ret.Pointer())

	return retGo
}

var filterInputStreamGetCloseBaseStreamFunction *gi.Function
var filterInputStreamGetCloseBaseStreamFunction_Once sync.Once

func filterInputStreamGetCloseBaseStreamFunction_Set() error {
	var err error
	filterInputStreamGetCloseBaseStreamFunction_Once.Do(func() {
		err = filterInputStreamObject_Set()
		if err != nil {
			return
		}
		filterInputStreamGetCloseBaseStreamFunction, err = filterInputStreamObject.InvokerNew("get_close_base_stream")
	})
	return err
}

// GetCloseBaseStream is a representation of the C type g_filter_input_stream_get_close_base_stream.
func (recv *FilterInputStream) GetCloseBaseStream() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := filterInputStreamGetCloseBaseStreamFunction_Set()
	if err == nil {
		ret = filterInputStreamGetCloseBaseStreamFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var filterInputStreamSetCloseBaseStreamFunction *gi.Function
var filterInputStreamSetCloseBaseStreamFunction_Once sync.Once

func filterInputStreamSetCloseBaseStreamFunction_Set() error {
	var err error
	filterInputStreamSetCloseBaseStreamFunction_Once.Do(func() {
		err = filterInputStreamObject_Set()
		if err != nil {
			return
		}
		filterInputStreamSetCloseBaseStreamFunction, err = filterInputStreamObject.InvokerNew("set_close_base_stream")
	})
	return err
}

// SetCloseBaseStream is a representation of the C type g_filter_input_stream_set_close_base_stream.
func (recv *FilterInputStream) SetCloseBaseStream(closeBase bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetBoolean(closeBase)

	err := filterInputStreamSetCloseBaseStreamFunction_Set()
	if err == nil {
		filterInputStreamSetCloseBaseStreamFunction.Invoke(inArgs[:], nil)
	}

	return
}

var filterOutputStreamObject *gi.Object
var filterOutputStreamObject_Once sync.Once

func filterOutputStreamObject_Set() error {
	var err error
	filterOutputStreamObject_Once.Do(func() {
		filterOutputStreamObject, err = gi.ObjectNew("Gio", "FilterOutputStream")
	})
	return err
}

type FilterOutputStream struct {
	native unsafe.Pointer
}

func FilterOutputStreamNewFromNative(native unsafe.Pointer) *FilterOutputStream {
	instance := &FilterOutputStream{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// OutputStream upcasts to *OutputStream
func (recv *FilterOutputStream) OutputStream() *OutputStream {
	return OutputStreamNewFromNative(recv.Native())
}

// Object upcasts to *Object
func (recv *FilterOutputStream) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToFilterOutputStream down casts any arbitrary Object to FilterOutputStream.
Exercise care, as this is a potentially dangerous function
if the Object is not a FilterOutputStream.
*/
func CastToFilterOutputStream(object *gobject.Object) *FilterOutputStream {
	return FilterOutputStreamNewFromNative(object.Native())
}

// Equals compares this FilterOutputStream with another FilterOutputStream, and returns true if they represent the same GObject.
func (recv *FilterOutputStream) Equals(other *FilterOutputStream) bool {
	return other.Native() == recv.Native()
}

func (recv *FilterOutputStream) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *FilterOutputStream) FieldParentInstance() *OutputStream {
	argValue := gi.ObjectFieldGet(filterOutputStreamObject, recv.Native(), "parent_instance")
	value := OutputStreamNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *FilterOutputStream) SetFieldParentInstance(value *OutputStream) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(filterOutputStreamObject, recv.Native(), "parent_instance", argValue)
}

// FieldBaseStream returns the C field 'base_stream'.
func (recv *FilterOutputStream) FieldBaseStream() *OutputStream {
	argValue := gi.ObjectFieldGet(filterOutputStreamObject, recv.Native(), "base_stream")
	value := OutputStreamNewFromNative(argValue.Pointer())
	return value
}

// SetFieldBaseStream sets the value of the C field 'base_stream'.
func (recv *FilterOutputStream) SetFieldBaseStream(value *OutputStream) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(filterOutputStreamObject, recv.Native(), "base_stream", argValue)
}

var filterOutputStreamGetBaseStreamFunction *gi.Function
var filterOutputStreamGetBaseStreamFunction_Once sync.Once

func filterOutputStreamGetBaseStreamFunction_Set() error {
	var err error
	filterOutputStreamGetBaseStreamFunction_Once.Do(func() {
		err = filterOutputStreamObject_Set()
		if err != nil {
			return
		}
		filterOutputStreamGetBaseStreamFunction, err = filterOutputStreamObject.InvokerNew("get_base_stream")
	})
	return err
}

// GetBaseStream is a representation of the C type g_filter_output_stream_get_base_stream.
func (recv *FilterOutputStream) GetBaseStream() *OutputStream {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := filterOutputStreamGetBaseStreamFunction_Set()
	if err == nil {
		ret = filterOutputStreamGetBaseStreamFunction.Invoke(inArgs[:], nil)
	}

	retGo := OutputStreamNewFromNative(ret.Pointer())

	return retGo
}

var filterOutputStreamGetCloseBaseStreamFunction *gi.Function
var filterOutputStreamGetCloseBaseStreamFunction_Once sync.Once

func filterOutputStreamGetCloseBaseStreamFunction_Set() error {
	var err error
	filterOutputStreamGetCloseBaseStreamFunction_Once.Do(func() {
		err = filterOutputStreamObject_Set()
		if err != nil {
			return
		}
		filterOutputStreamGetCloseBaseStreamFunction, err = filterOutputStreamObject.InvokerNew("get_close_base_stream")
	})
	return err
}

// GetCloseBaseStream is a representation of the C type g_filter_output_stream_get_close_base_stream.
func (recv *FilterOutputStream) GetCloseBaseStream() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := filterOutputStreamGetCloseBaseStreamFunction_Set()
	if err == nil {
		ret = filterOutputStreamGetCloseBaseStreamFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var filterOutputStreamSetCloseBaseStreamFunction *gi.Function
var filterOutputStreamSetCloseBaseStreamFunction_Once sync.Once

func filterOutputStreamSetCloseBaseStreamFunction_Set() error {
	var err error
	filterOutputStreamSetCloseBaseStreamFunction_Once.Do(func() {
		err = filterOutputStreamObject_Set()
		if err != nil {
			return
		}
		filterOutputStreamSetCloseBaseStreamFunction, err = filterOutputStreamObject.InvokerNew("set_close_base_stream")
	})
	return err
}

// SetCloseBaseStream is a representation of the C type g_filter_output_stream_set_close_base_stream.
func (recv *FilterOutputStream) SetCloseBaseStream(closeBase bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetBoolean(closeBase)

	err := filterOutputStreamSetCloseBaseStreamFunction_Set()
	if err == nil {
		filterOutputStreamSetCloseBaseStreamFunction.Invoke(inArgs[:], nil)
	}

	return
}

var iOModuleObject *gi.Object
var iOModuleObject_Once sync.Once

func iOModuleObject_Set() error {
	var err error
	iOModuleObject_Once.Do(func() {
		iOModuleObject, err = gi.ObjectNew("Gio", "IOModule")
	})
	return err
}

type IOModule struct {
	native unsafe.Pointer
}

func IOModuleNewFromNative(native unsafe.Pointer) *IOModule {
	instance := &IOModule{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// TypeModule upcasts to *TypeModule
func (recv *IOModule) TypeModule() *gobject.TypeModule {
	return gobject.TypeModuleNewFromNative(recv.Native())
}

// Object upcasts to *Object
func (recv *IOModule) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToIOModule down casts any arbitrary Object to IOModule.
Exercise care, as this is a potentially dangerous function
if the Object is not a IOModule.
*/
func CastToIOModule(object *gobject.Object) *IOModule {
	return IOModuleNewFromNative(object.Native())
}

// Equals compares this IOModule with another IOModule, and returns true if they represent the same GObject.
func (recv *IOModule) Equals(other *IOModule) bool {
	return other.Native() == recv.Native()
}

func (recv *IOModule) Native() unsafe.Pointer {
	return recv.native
}

var iOModuleNewFunction *gi.Function
var iOModuleNewFunction_Once sync.Once

func iOModuleNewFunction_Set() error {
	var err error
	iOModuleNewFunction_Once.Do(func() {
		err = iOModuleObject_Set()
		if err != nil {
			return
		}
		iOModuleNewFunction, err = iOModuleObject.InvokerNew("new")
	})
	return err
}

// IOModuleNew is a representation of the C type g_io_module_new.
func IOModuleNew(filename string) *IOModule {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(filename)

	var ret gi.Argument

	err := iOModuleNewFunction_Set()
	if err == nil {
		ret = iOModuleNewFunction.Invoke(inArgs[:], nil)
	}

	retGo := IOModuleNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

var iOModuleLoadFunction *gi.Function
var iOModuleLoadFunction_Once sync.Once

func iOModuleLoadFunction_Set() error {
	var err error
	iOModuleLoadFunction_Once.Do(func() {
		err = iOModuleObject_Set()
		if err != nil {
			return
		}
		iOModuleLoadFunction, err = iOModuleObject.InvokerNew("load")
	})
	return err
}

// Load is a representation of the C type g_io_module_load.
func (recv *IOModule) Load() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := iOModuleLoadFunction_Set()
	if err == nil {
		iOModuleLoadFunction.Invoke(inArgs[:], nil)
	}

	return
}

var iOModuleUnloadFunction *gi.Function
var iOModuleUnloadFunction_Once sync.Once

func iOModuleUnloadFunction_Set() error {
	var err error
	iOModuleUnloadFunction_Once.Do(func() {
		err = iOModuleObject_Set()
		if err != nil {
			return
		}
		iOModuleUnloadFunction, err = iOModuleObject.InvokerNew("unload")
	})
	return err
}

// Unload is a representation of the C type g_io_module_unload.
func (recv *IOModule) Unload() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := iOModuleUnloadFunction_Set()
	if err == nil {
		iOModuleUnloadFunction.Invoke(inArgs[:], nil)
	}

	return
}

var iOStreamObject *gi.Object
var iOStreamObject_Once sync.Once

func iOStreamObject_Set() error {
	var err error
	iOStreamObject_Once.Do(func() {
		iOStreamObject, err = gi.ObjectNew("Gio", "IOStream")
	})
	return err
}

type IOStream struct {
	native unsafe.Pointer
}

func IOStreamNewFromNative(native unsafe.Pointer) *IOStream {
	instance := &IOStream{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *IOStream) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToIOStream down casts any arbitrary Object to IOStream.
Exercise care, as this is a potentially dangerous function
if the Object is not a IOStream.
*/
func CastToIOStream(object *gobject.Object) *IOStream {
	return IOStreamNewFromNative(object.Native())
}

// Equals compares this IOStream with another IOStream, and returns true if they represent the same GObject.
func (recv *IOStream) Equals(other *IOStream) bool {
	return other.Native() == recv.Native()
}

func (recv *IOStream) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *IOStream) FieldParentInstance() *gobject.Object {
	argValue := gi.ObjectFieldGet(iOStreamObject, recv.Native(), "parent_instance")
	value := gobject.ObjectNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *IOStream) SetFieldParentInstance(value *gobject.Object) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(iOStreamObject, recv.Native(), "parent_instance", argValue)
}

var iOStreamClearPendingFunction *gi.Function
var iOStreamClearPendingFunction_Once sync.Once

func iOStreamClearPendingFunction_Set() error {
	var err error
	iOStreamClearPendingFunction_Once.Do(func() {
		err = iOStreamObject_Set()
		if err != nil {
			return
		}
		iOStreamClearPendingFunction, err = iOStreamObject.InvokerNew("clear_pending")
	})
	return err
}

// ClearPending is a representation of the C type g_io_stream_clear_pending.
func (recv *IOStream) ClearPending() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := iOStreamClearPendingFunction_Set()
	if err == nil {
		iOStreamClearPendingFunction.Invoke(inArgs[:], nil)
	}

	return
}

var iOStreamCloseFunction *gi.Function
var iOStreamCloseFunction_Once sync.Once

func iOStreamCloseFunction_Set() error {
	var err error
	iOStreamCloseFunction_Once.Do(func() {
		err = iOStreamObject_Set()
		if err != nil {
			return
		}
		iOStreamCloseFunction, err = iOStreamObject.InvokerNew("close")
	})
	return err
}

// Close is a representation of the C type g_io_stream_close.
func (recv *IOStream) Close(cancellable *Cancellable) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(cancellable.Native())

	var ret gi.Argument

	err := iOStreamCloseFunction_Set()
	if err == nil {
		ret = iOStreamCloseFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'g_io_stream_close_async' : parameter 'callback' of type 'AsyncReadyCallback' not supported

// UNSUPPORTED : C value 'g_io_stream_close_finish' : parameter 'result' of type 'AsyncResult' not supported

var iOStreamGetInputStreamFunction *gi.Function
var iOStreamGetInputStreamFunction_Once sync.Once

func iOStreamGetInputStreamFunction_Set() error {
	var err error
	iOStreamGetInputStreamFunction_Once.Do(func() {
		err = iOStreamObject_Set()
		if err != nil {
			return
		}
		iOStreamGetInputStreamFunction, err = iOStreamObject.InvokerNew("get_input_stream")
	})
	return err
}

// GetInputStream is a representation of the C type g_io_stream_get_input_stream.
func (recv *IOStream) GetInputStream() *InputStream {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := iOStreamGetInputStreamFunction_Set()
	if err == nil {
		ret = iOStreamGetInputStreamFunction.Invoke(inArgs[:], nil)
	}

	retGo := InputStreamNewFromNative(ret.Pointer())

	return retGo
}

var iOStreamGetOutputStreamFunction *gi.Function
var iOStreamGetOutputStreamFunction_Once sync.Once

func iOStreamGetOutputStreamFunction_Set() error {
	var err error
	iOStreamGetOutputStreamFunction_Once.Do(func() {
		err = iOStreamObject_Set()
		if err != nil {
			return
		}
		iOStreamGetOutputStreamFunction, err = iOStreamObject.InvokerNew("get_output_stream")
	})
	return err
}

// GetOutputStream is a representation of the C type g_io_stream_get_output_stream.
func (recv *IOStream) GetOutputStream() *OutputStream {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := iOStreamGetOutputStreamFunction_Set()
	if err == nil {
		ret = iOStreamGetOutputStreamFunction.Invoke(inArgs[:], nil)
	}

	retGo := OutputStreamNewFromNative(ret.Pointer())

	return retGo
}

var iOStreamHasPendingFunction *gi.Function
var iOStreamHasPendingFunction_Once sync.Once

func iOStreamHasPendingFunction_Set() error {
	var err error
	iOStreamHasPendingFunction_Once.Do(func() {
		err = iOStreamObject_Set()
		if err != nil {
			return
		}
		iOStreamHasPendingFunction, err = iOStreamObject.InvokerNew("has_pending")
	})
	return err
}

// HasPending is a representation of the C type g_io_stream_has_pending.
func (recv *IOStream) HasPending() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := iOStreamHasPendingFunction_Set()
	if err == nil {
		ret = iOStreamHasPendingFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var iOStreamIsClosedFunction *gi.Function
var iOStreamIsClosedFunction_Once sync.Once

func iOStreamIsClosedFunction_Set() error {
	var err error
	iOStreamIsClosedFunction_Once.Do(func() {
		err = iOStreamObject_Set()
		if err != nil {
			return
		}
		iOStreamIsClosedFunction, err = iOStreamObject.InvokerNew("is_closed")
	})
	return err
}

// IsClosed is a representation of the C type g_io_stream_is_closed.
func (recv *IOStream) IsClosed() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := iOStreamIsClosedFunction_Set()
	if err == nil {
		ret = iOStreamIsClosedFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var iOStreamSetPendingFunction *gi.Function
var iOStreamSetPendingFunction_Once sync.Once

func iOStreamSetPendingFunction_Set() error {
	var err error
	iOStreamSetPendingFunction_Once.Do(func() {
		err = iOStreamObject_Set()
		if err != nil {
			return
		}
		iOStreamSetPendingFunction, err = iOStreamObject.InvokerNew("set_pending")
	})
	return err
}

// SetPending is a representation of the C type g_io_stream_set_pending.
func (recv *IOStream) SetPending() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := iOStreamSetPendingFunction_Set()
	if err == nil {
		ret = iOStreamSetPendingFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'g_io_stream_splice_async' : parameter 'flags' of type 'IOStreamSpliceFlags' not supported

var inetAddressObject *gi.Object
var inetAddressObject_Once sync.Once

func inetAddressObject_Set() error {
	var err error
	inetAddressObject_Once.Do(func() {
		inetAddressObject, err = gi.ObjectNew("Gio", "InetAddress")
	})
	return err
}

type InetAddress struct {
	native unsafe.Pointer
}

func InetAddressNewFromNative(native unsafe.Pointer) *InetAddress {
	instance := &InetAddress{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *InetAddress) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToInetAddress down casts any arbitrary Object to InetAddress.
Exercise care, as this is a potentially dangerous function
if the Object is not a InetAddress.
*/
func CastToInetAddress(object *gobject.Object) *InetAddress {
	return InetAddressNewFromNative(object.Native())
}

// Equals compares this InetAddress with another InetAddress, and returns true if they represent the same GObject.
func (recv *InetAddress) Equals(other *InetAddress) bool {
	return other.Native() == recv.Native()
}

func (recv *InetAddress) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *InetAddress) FieldParentInstance() *gobject.Object {
	argValue := gi.ObjectFieldGet(inetAddressObject, recv.Native(), "parent_instance")
	value := gobject.ObjectNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *InetAddress) SetFieldParentInstance(value *gobject.Object) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(inetAddressObject, recv.Native(), "parent_instance", argValue)
}

var inetAddressNewAnyFunction *gi.Function
var inetAddressNewAnyFunction_Once sync.Once

func inetAddressNewAnyFunction_Set() error {
	var err error
	inetAddressNewAnyFunction_Once.Do(func() {
		err = inetAddressObject_Set()
		if err != nil {
			return
		}
		inetAddressNewAnyFunction, err = inetAddressObject.InvokerNew("new_any")
	})
	return err
}

// InetAddressNewAny is a representation of the C type g_inet_address_new_any.
func InetAddressNewAny(family SocketFamily) *InetAddress {
	var inArgs [1]gi.Argument
	inArgs[0].SetInt32(int32(family))

	var ret gi.Argument

	err := inetAddressNewAnyFunction_Set()
	if err == nil {
		ret = inetAddressNewAnyFunction.Invoke(inArgs[:], nil)
	}

	retGo := InetAddressNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

// UNSUPPORTED : C value 'g_inet_address_new_from_bytes' : parameter 'bytes' of type 'nil' not supported

var inetAddressNewFromStringFunction *gi.Function
var inetAddressNewFromStringFunction_Once sync.Once

func inetAddressNewFromStringFunction_Set() error {
	var err error
	inetAddressNewFromStringFunction_Once.Do(func() {
		err = inetAddressObject_Set()
		if err != nil {
			return
		}
		inetAddressNewFromStringFunction, err = inetAddressObject.InvokerNew("new_from_string")
	})
	return err
}

// InetAddressNewFromString is a representation of the C type g_inet_address_new_from_string.
func InetAddressNewFromString(string_ string) *InetAddress {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(string_)

	var ret gi.Argument

	err := inetAddressNewFromStringFunction_Set()
	if err == nil {
		ret = inetAddressNewFromStringFunction.Invoke(inArgs[:], nil)
	}

	retGo := InetAddressNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

var inetAddressNewLoopbackFunction *gi.Function
var inetAddressNewLoopbackFunction_Once sync.Once

func inetAddressNewLoopbackFunction_Set() error {
	var err error
	inetAddressNewLoopbackFunction_Once.Do(func() {
		err = inetAddressObject_Set()
		if err != nil {
			return
		}
		inetAddressNewLoopbackFunction, err = inetAddressObject.InvokerNew("new_loopback")
	})
	return err
}

// InetAddressNewLoopback is a representation of the C type g_inet_address_new_loopback.
func InetAddressNewLoopback(family SocketFamily) *InetAddress {
	var inArgs [1]gi.Argument
	inArgs[0].SetInt32(int32(family))

	var ret gi.Argument

	err := inetAddressNewLoopbackFunction_Set()
	if err == nil {
		ret = inetAddressNewLoopbackFunction.Invoke(inArgs[:], nil)
	}

	retGo := InetAddressNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

var inetAddressEqualFunction *gi.Function
var inetAddressEqualFunction_Once sync.Once

func inetAddressEqualFunction_Set() error {
	var err error
	inetAddressEqualFunction_Once.Do(func() {
		err = inetAddressObject_Set()
		if err != nil {
			return
		}
		inetAddressEqualFunction, err = inetAddressObject.InvokerNew("equal")
	})
	return err
}

// Equal is a representation of the C type g_inet_address_equal.
func (recv *InetAddress) Equal(otherAddress *InetAddress) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(otherAddress.Native())

	var ret gi.Argument

	err := inetAddressEqualFunction_Set()
	if err == nil {
		ret = inetAddressEqualFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var inetAddressGetFamilyFunction *gi.Function
var inetAddressGetFamilyFunction_Once sync.Once

func inetAddressGetFamilyFunction_Set() error {
	var err error
	inetAddressGetFamilyFunction_Once.Do(func() {
		err = inetAddressObject_Set()
		if err != nil {
			return
		}
		inetAddressGetFamilyFunction, err = inetAddressObject.InvokerNew("get_family")
	})
	return err
}

// GetFamily is a representation of the C type g_inet_address_get_family.
func (recv *InetAddress) GetFamily() SocketFamily {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := inetAddressGetFamilyFunction_Set()
	if err == nil {
		ret = inetAddressGetFamilyFunction.Invoke(inArgs[:], nil)
	}

	retGo := SocketFamily(ret.Int32())

	return retGo
}

var inetAddressGetIsAnyFunction *gi.Function
var inetAddressGetIsAnyFunction_Once sync.Once

func inetAddressGetIsAnyFunction_Set() error {
	var err error
	inetAddressGetIsAnyFunction_Once.Do(func() {
		err = inetAddressObject_Set()
		if err != nil {
			return
		}
		inetAddressGetIsAnyFunction, err = inetAddressObject.InvokerNew("get_is_any")
	})
	return err
}

// GetIsAny is a representation of the C type g_inet_address_get_is_any.
func (recv *InetAddress) GetIsAny() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := inetAddressGetIsAnyFunction_Set()
	if err == nil {
		ret = inetAddressGetIsAnyFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var inetAddressGetIsLinkLocalFunction *gi.Function
var inetAddressGetIsLinkLocalFunction_Once sync.Once

func inetAddressGetIsLinkLocalFunction_Set() error {
	var err error
	inetAddressGetIsLinkLocalFunction_Once.Do(func() {
		err = inetAddressObject_Set()
		if err != nil {
			return
		}
		inetAddressGetIsLinkLocalFunction, err = inetAddressObject.InvokerNew("get_is_link_local")
	})
	return err
}

// GetIsLinkLocal is a representation of the C type g_inet_address_get_is_link_local.
func (recv *InetAddress) GetIsLinkLocal() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := inetAddressGetIsLinkLocalFunction_Set()
	if err == nil {
		ret = inetAddressGetIsLinkLocalFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var inetAddressGetIsLoopbackFunction *gi.Function
var inetAddressGetIsLoopbackFunction_Once sync.Once

func inetAddressGetIsLoopbackFunction_Set() error {
	var err error
	inetAddressGetIsLoopbackFunction_Once.Do(func() {
		err = inetAddressObject_Set()
		if err != nil {
			return
		}
		inetAddressGetIsLoopbackFunction, err = inetAddressObject.InvokerNew("get_is_loopback")
	})
	return err
}

// GetIsLoopback is a representation of the C type g_inet_address_get_is_loopback.
func (recv *InetAddress) GetIsLoopback() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := inetAddressGetIsLoopbackFunction_Set()
	if err == nil {
		ret = inetAddressGetIsLoopbackFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var inetAddressGetIsMcGlobalFunction *gi.Function
var inetAddressGetIsMcGlobalFunction_Once sync.Once

func inetAddressGetIsMcGlobalFunction_Set() error {
	var err error
	inetAddressGetIsMcGlobalFunction_Once.Do(func() {
		err = inetAddressObject_Set()
		if err != nil {
			return
		}
		inetAddressGetIsMcGlobalFunction, err = inetAddressObject.InvokerNew("get_is_mc_global")
	})
	return err
}

// GetIsMcGlobal is a representation of the C type g_inet_address_get_is_mc_global.
func (recv *InetAddress) GetIsMcGlobal() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := inetAddressGetIsMcGlobalFunction_Set()
	if err == nil {
		ret = inetAddressGetIsMcGlobalFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var inetAddressGetIsMcLinkLocalFunction *gi.Function
var inetAddressGetIsMcLinkLocalFunction_Once sync.Once

func inetAddressGetIsMcLinkLocalFunction_Set() error {
	var err error
	inetAddressGetIsMcLinkLocalFunction_Once.Do(func() {
		err = inetAddressObject_Set()
		if err != nil {
			return
		}
		inetAddressGetIsMcLinkLocalFunction, err = inetAddressObject.InvokerNew("get_is_mc_link_local")
	})
	return err
}

// GetIsMcLinkLocal is a representation of the C type g_inet_address_get_is_mc_link_local.
func (recv *InetAddress) GetIsMcLinkLocal() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := inetAddressGetIsMcLinkLocalFunction_Set()
	if err == nil {
		ret = inetAddressGetIsMcLinkLocalFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var inetAddressGetIsMcNodeLocalFunction *gi.Function
var inetAddressGetIsMcNodeLocalFunction_Once sync.Once

func inetAddressGetIsMcNodeLocalFunction_Set() error {
	var err error
	inetAddressGetIsMcNodeLocalFunction_Once.Do(func() {
		err = inetAddressObject_Set()
		if err != nil {
			return
		}
		inetAddressGetIsMcNodeLocalFunction, err = inetAddressObject.InvokerNew("get_is_mc_node_local")
	})
	return err
}

// GetIsMcNodeLocal is a representation of the C type g_inet_address_get_is_mc_node_local.
func (recv *InetAddress) GetIsMcNodeLocal() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := inetAddressGetIsMcNodeLocalFunction_Set()
	if err == nil {
		ret = inetAddressGetIsMcNodeLocalFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var inetAddressGetIsMcOrgLocalFunction *gi.Function
var inetAddressGetIsMcOrgLocalFunction_Once sync.Once

func inetAddressGetIsMcOrgLocalFunction_Set() error {
	var err error
	inetAddressGetIsMcOrgLocalFunction_Once.Do(func() {
		err = inetAddressObject_Set()
		if err != nil {
			return
		}
		inetAddressGetIsMcOrgLocalFunction, err = inetAddressObject.InvokerNew("get_is_mc_org_local")
	})
	return err
}

// GetIsMcOrgLocal is a representation of the C type g_inet_address_get_is_mc_org_local.
func (recv *InetAddress) GetIsMcOrgLocal() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := inetAddressGetIsMcOrgLocalFunction_Set()
	if err == nil {
		ret = inetAddressGetIsMcOrgLocalFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var inetAddressGetIsMcSiteLocalFunction *gi.Function
var inetAddressGetIsMcSiteLocalFunction_Once sync.Once

func inetAddressGetIsMcSiteLocalFunction_Set() error {
	var err error
	inetAddressGetIsMcSiteLocalFunction_Once.Do(func() {
		err = inetAddressObject_Set()
		if err != nil {
			return
		}
		inetAddressGetIsMcSiteLocalFunction, err = inetAddressObject.InvokerNew("get_is_mc_site_local")
	})
	return err
}

// GetIsMcSiteLocal is a representation of the C type g_inet_address_get_is_mc_site_local.
func (recv *InetAddress) GetIsMcSiteLocal() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := inetAddressGetIsMcSiteLocalFunction_Set()
	if err == nil {
		ret = inetAddressGetIsMcSiteLocalFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var inetAddressGetIsMulticastFunction *gi.Function
var inetAddressGetIsMulticastFunction_Once sync.Once

func inetAddressGetIsMulticastFunction_Set() error {
	var err error
	inetAddressGetIsMulticastFunction_Once.Do(func() {
		err = inetAddressObject_Set()
		if err != nil {
			return
		}
		inetAddressGetIsMulticastFunction, err = inetAddressObject.InvokerNew("get_is_multicast")
	})
	return err
}

// GetIsMulticast is a representation of the C type g_inet_address_get_is_multicast.
func (recv *InetAddress) GetIsMulticast() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := inetAddressGetIsMulticastFunction_Set()
	if err == nil {
		ret = inetAddressGetIsMulticastFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var inetAddressGetIsSiteLocalFunction *gi.Function
var inetAddressGetIsSiteLocalFunction_Once sync.Once

func inetAddressGetIsSiteLocalFunction_Set() error {
	var err error
	inetAddressGetIsSiteLocalFunction_Once.Do(func() {
		err = inetAddressObject_Set()
		if err != nil {
			return
		}
		inetAddressGetIsSiteLocalFunction, err = inetAddressObject.InvokerNew("get_is_site_local")
	})
	return err
}

// GetIsSiteLocal is a representation of the C type g_inet_address_get_is_site_local.
func (recv *InetAddress) GetIsSiteLocal() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := inetAddressGetIsSiteLocalFunction_Set()
	if err == nil {
		ret = inetAddressGetIsSiteLocalFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var inetAddressGetNativeSizeFunction *gi.Function
var inetAddressGetNativeSizeFunction_Once sync.Once

func inetAddressGetNativeSizeFunction_Set() error {
	var err error
	inetAddressGetNativeSizeFunction_Once.Do(func() {
		err = inetAddressObject_Set()
		if err != nil {
			return
		}
		inetAddressGetNativeSizeFunction, err = inetAddressObject.InvokerNew("get_native_size")
	})
	return err
}

// GetNativeSize is a representation of the C type g_inet_address_get_native_size.
func (recv *InetAddress) GetNativeSize() uint64 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := inetAddressGetNativeSizeFunction_Set()
	if err == nil {
		ret = inetAddressGetNativeSizeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint64()

	return retGo
}

var inetAddressToBytesFunction *gi.Function
var inetAddressToBytesFunction_Once sync.Once

func inetAddressToBytesFunction_Set() error {
	var err error
	inetAddressToBytesFunction_Once.Do(func() {
		err = inetAddressObject_Set()
		if err != nil {
			return
		}
		inetAddressToBytesFunction, err = inetAddressObject.InvokerNew("to_bytes")
	})
	return err
}

// ToBytes is a representation of the C type g_inet_address_to_bytes.
func (recv *InetAddress) ToBytes() uint8 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := inetAddressToBytesFunction_Set()
	if err == nil {
		ret = inetAddressToBytesFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint8()

	return retGo
}

var inetAddressToStringFunction *gi.Function
var inetAddressToStringFunction_Once sync.Once

func inetAddressToStringFunction_Set() error {
	var err error
	inetAddressToStringFunction_Once.Do(func() {
		err = inetAddressObject_Set()
		if err != nil {
			return
		}
		inetAddressToStringFunction, err = inetAddressObject.InvokerNew("to_string")
	})
	return err
}

// ToString is a representation of the C type g_inet_address_to_string.
func (recv *InetAddress) ToString() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := inetAddressToStringFunction_Set()
	if err == nil {
		ret = inetAddressToStringFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

var inetAddressMaskObject *gi.Object
var inetAddressMaskObject_Once sync.Once

func inetAddressMaskObject_Set() error {
	var err error
	inetAddressMaskObject_Once.Do(func() {
		inetAddressMaskObject, err = gi.ObjectNew("Gio", "InetAddressMask")
	})
	return err
}

type InetAddressMask struct {
	native unsafe.Pointer
}

func InetAddressMaskNewFromNative(native unsafe.Pointer) *InetAddressMask {
	instance := &InetAddressMask{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *InetAddressMask) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToInetAddressMask down casts any arbitrary Object to InetAddressMask.
Exercise care, as this is a potentially dangerous function
if the Object is not a InetAddressMask.
*/
func CastToInetAddressMask(object *gobject.Object) *InetAddressMask {
	return InetAddressMaskNewFromNative(object.Native())
}

// Equals compares this InetAddressMask with another InetAddressMask, and returns true if they represent the same GObject.
func (recv *InetAddressMask) Equals(other *InetAddressMask) bool {
	return other.Native() == recv.Native()
}

func (recv *InetAddressMask) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *InetAddressMask) FieldParentInstance() *gobject.Object {
	argValue := gi.ObjectFieldGet(inetAddressMaskObject, recv.Native(), "parent_instance")
	value := gobject.ObjectNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *InetAddressMask) SetFieldParentInstance(value *gobject.Object) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(inetAddressMaskObject, recv.Native(), "parent_instance", argValue)
}

var inetAddressMaskNewFunction *gi.Function
var inetAddressMaskNewFunction_Once sync.Once

func inetAddressMaskNewFunction_Set() error {
	var err error
	inetAddressMaskNewFunction_Once.Do(func() {
		err = inetAddressMaskObject_Set()
		if err != nil {
			return
		}
		inetAddressMaskNewFunction, err = inetAddressMaskObject.InvokerNew("new")
	})
	return err
}

// InetAddressMaskNew is a representation of the C type g_inet_address_mask_new.
func InetAddressMaskNew(addr *InetAddress, length uint32) *InetAddressMask {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(addr.Native())
	inArgs[1].SetUint32(length)

	var ret gi.Argument

	err := inetAddressMaskNewFunction_Set()
	if err == nil {
		ret = inetAddressMaskNewFunction.Invoke(inArgs[:], nil)
	}

	retGo := InetAddressMaskNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

var inetAddressMaskNewFromStringFunction *gi.Function
var inetAddressMaskNewFromStringFunction_Once sync.Once

func inetAddressMaskNewFromStringFunction_Set() error {
	var err error
	inetAddressMaskNewFromStringFunction_Once.Do(func() {
		err = inetAddressMaskObject_Set()
		if err != nil {
			return
		}
		inetAddressMaskNewFromStringFunction, err = inetAddressMaskObject.InvokerNew("new_from_string")
	})
	return err
}

// InetAddressMaskNewFromString is a representation of the C type g_inet_address_mask_new_from_string.
func InetAddressMaskNewFromString(maskString string) *InetAddressMask {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(maskString)

	var ret gi.Argument

	err := inetAddressMaskNewFromStringFunction_Set()
	if err == nil {
		ret = inetAddressMaskNewFromStringFunction.Invoke(inArgs[:], nil)
	}

	retGo := InetAddressMaskNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

var inetAddressMaskEqualFunction *gi.Function
var inetAddressMaskEqualFunction_Once sync.Once

func inetAddressMaskEqualFunction_Set() error {
	var err error
	inetAddressMaskEqualFunction_Once.Do(func() {
		err = inetAddressMaskObject_Set()
		if err != nil {
			return
		}
		inetAddressMaskEqualFunction, err = inetAddressMaskObject.InvokerNew("equal")
	})
	return err
}

// Equal is a representation of the C type g_inet_address_mask_equal.
func (recv *InetAddressMask) Equal(mask2 *InetAddressMask) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(mask2.Native())

	var ret gi.Argument

	err := inetAddressMaskEqualFunction_Set()
	if err == nil {
		ret = inetAddressMaskEqualFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var inetAddressMaskGetAddressFunction *gi.Function
var inetAddressMaskGetAddressFunction_Once sync.Once

func inetAddressMaskGetAddressFunction_Set() error {
	var err error
	inetAddressMaskGetAddressFunction_Once.Do(func() {
		err = inetAddressMaskObject_Set()
		if err != nil {
			return
		}
		inetAddressMaskGetAddressFunction, err = inetAddressMaskObject.InvokerNew("get_address")
	})
	return err
}

// GetAddress is a representation of the C type g_inet_address_mask_get_address.
func (recv *InetAddressMask) GetAddress() *InetAddress {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := inetAddressMaskGetAddressFunction_Set()
	if err == nil {
		ret = inetAddressMaskGetAddressFunction.Invoke(inArgs[:], nil)
	}

	retGo := InetAddressNewFromNative(ret.Pointer())

	return retGo
}

var inetAddressMaskGetFamilyFunction *gi.Function
var inetAddressMaskGetFamilyFunction_Once sync.Once

func inetAddressMaskGetFamilyFunction_Set() error {
	var err error
	inetAddressMaskGetFamilyFunction_Once.Do(func() {
		err = inetAddressMaskObject_Set()
		if err != nil {
			return
		}
		inetAddressMaskGetFamilyFunction, err = inetAddressMaskObject.InvokerNew("get_family")
	})
	return err
}

// GetFamily is a representation of the C type g_inet_address_mask_get_family.
func (recv *InetAddressMask) GetFamily() SocketFamily {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := inetAddressMaskGetFamilyFunction_Set()
	if err == nil {
		ret = inetAddressMaskGetFamilyFunction.Invoke(inArgs[:], nil)
	}

	retGo := SocketFamily(ret.Int32())

	return retGo
}

var inetAddressMaskGetLengthFunction *gi.Function
var inetAddressMaskGetLengthFunction_Once sync.Once

func inetAddressMaskGetLengthFunction_Set() error {
	var err error
	inetAddressMaskGetLengthFunction_Once.Do(func() {
		err = inetAddressMaskObject_Set()
		if err != nil {
			return
		}
		inetAddressMaskGetLengthFunction, err = inetAddressMaskObject.InvokerNew("get_length")
	})
	return err
}

// GetLength is a representation of the C type g_inet_address_mask_get_length.
func (recv *InetAddressMask) GetLength() uint32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := inetAddressMaskGetLengthFunction_Set()
	if err == nil {
		ret = inetAddressMaskGetLengthFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo
}

var inetAddressMaskMatchesFunction *gi.Function
var inetAddressMaskMatchesFunction_Once sync.Once

func inetAddressMaskMatchesFunction_Set() error {
	var err error
	inetAddressMaskMatchesFunction_Once.Do(func() {
		err = inetAddressMaskObject_Set()
		if err != nil {
			return
		}
		inetAddressMaskMatchesFunction, err = inetAddressMaskObject.InvokerNew("matches")
	})
	return err
}

// Matches is a representation of the C type g_inet_address_mask_matches.
func (recv *InetAddressMask) Matches(address *InetAddress) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(address.Native())

	var ret gi.Argument

	err := inetAddressMaskMatchesFunction_Set()
	if err == nil {
		ret = inetAddressMaskMatchesFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var inetAddressMaskToStringFunction *gi.Function
var inetAddressMaskToStringFunction_Once sync.Once

func inetAddressMaskToStringFunction_Set() error {
	var err error
	inetAddressMaskToStringFunction_Once.Do(func() {
		err = inetAddressMaskObject_Set()
		if err != nil {
			return
		}
		inetAddressMaskToStringFunction, err = inetAddressMaskObject.InvokerNew("to_string")
	})
	return err
}

// ToString is a representation of the C type g_inet_address_mask_to_string.
func (recv *InetAddressMask) ToString() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := inetAddressMaskToStringFunction_Set()
	if err == nil {
		ret = inetAddressMaskToStringFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

var inetSocketAddressObject *gi.Object
var inetSocketAddressObject_Once sync.Once

func inetSocketAddressObject_Set() error {
	var err error
	inetSocketAddressObject_Once.Do(func() {
		inetSocketAddressObject, err = gi.ObjectNew("Gio", "InetSocketAddress")
	})
	return err
}

type InetSocketAddress struct {
	native unsafe.Pointer
}

func InetSocketAddressNewFromNative(native unsafe.Pointer) *InetSocketAddress {
	instance := &InetSocketAddress{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// SocketAddress upcasts to *SocketAddress
func (recv *InetSocketAddress) SocketAddress() *SocketAddress {
	return SocketAddressNewFromNative(recv.Native())
}

// Object upcasts to *Object
func (recv *InetSocketAddress) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToInetSocketAddress down casts any arbitrary Object to InetSocketAddress.
Exercise care, as this is a potentially dangerous function
if the Object is not a InetSocketAddress.
*/
func CastToInetSocketAddress(object *gobject.Object) *InetSocketAddress {
	return InetSocketAddressNewFromNative(object.Native())
}

// Equals compares this InetSocketAddress with another InetSocketAddress, and returns true if they represent the same GObject.
func (recv *InetSocketAddress) Equals(other *InetSocketAddress) bool {
	return other.Native() == recv.Native()
}

func (recv *InetSocketAddress) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *InetSocketAddress) FieldParentInstance() *SocketAddress {
	argValue := gi.ObjectFieldGet(inetSocketAddressObject, recv.Native(), "parent_instance")
	value := SocketAddressNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *InetSocketAddress) SetFieldParentInstance(value *SocketAddress) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(inetSocketAddressObject, recv.Native(), "parent_instance", argValue)
}

var inetSocketAddressNewFunction *gi.Function
var inetSocketAddressNewFunction_Once sync.Once

func inetSocketAddressNewFunction_Set() error {
	var err error
	inetSocketAddressNewFunction_Once.Do(func() {
		err = inetSocketAddressObject_Set()
		if err != nil {
			return
		}
		inetSocketAddressNewFunction, err = inetSocketAddressObject.InvokerNew("new")
	})
	return err
}

// InetSocketAddressNew is a representation of the C type g_inet_socket_address_new.
func InetSocketAddressNew(address *InetAddress, port uint16) *InetSocketAddress {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(address.Native())
	inArgs[1].SetUint16(port)

	var ret gi.Argument

	err := inetSocketAddressNewFunction_Set()
	if err == nil {
		ret = inetSocketAddressNewFunction.Invoke(inArgs[:], nil)
	}

	retGo := InetSocketAddressNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

var inetSocketAddressNewFromStringFunction *gi.Function
var inetSocketAddressNewFromStringFunction_Once sync.Once

func inetSocketAddressNewFromStringFunction_Set() error {
	var err error
	inetSocketAddressNewFromStringFunction_Once.Do(func() {
		err = inetSocketAddressObject_Set()
		if err != nil {
			return
		}
		inetSocketAddressNewFromStringFunction, err = inetSocketAddressObject.InvokerNew("new_from_string")
	})
	return err
}

// InetSocketAddressNewFromString is a representation of the C type g_inet_socket_address_new_from_string.
func InetSocketAddressNewFromString(address string, port uint32) *InetSocketAddress {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(address)
	inArgs[1].SetUint32(port)

	var ret gi.Argument

	err := inetSocketAddressNewFromStringFunction_Set()
	if err == nil {
		ret = inetSocketAddressNewFromStringFunction.Invoke(inArgs[:], nil)
	}

	retGo := InetSocketAddressNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

var inetSocketAddressGetAddressFunction *gi.Function
var inetSocketAddressGetAddressFunction_Once sync.Once

func inetSocketAddressGetAddressFunction_Set() error {
	var err error
	inetSocketAddressGetAddressFunction_Once.Do(func() {
		err = inetSocketAddressObject_Set()
		if err != nil {
			return
		}
		inetSocketAddressGetAddressFunction, err = inetSocketAddressObject.InvokerNew("get_address")
	})
	return err
}

// GetAddress is a representation of the C type g_inet_socket_address_get_address.
func (recv *InetSocketAddress) GetAddress() *InetAddress {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := inetSocketAddressGetAddressFunction_Set()
	if err == nil {
		ret = inetSocketAddressGetAddressFunction.Invoke(inArgs[:], nil)
	}

	retGo := InetAddressNewFromNative(ret.Pointer())

	return retGo
}

var inetSocketAddressGetFlowinfoFunction *gi.Function
var inetSocketAddressGetFlowinfoFunction_Once sync.Once

func inetSocketAddressGetFlowinfoFunction_Set() error {
	var err error
	inetSocketAddressGetFlowinfoFunction_Once.Do(func() {
		err = inetSocketAddressObject_Set()
		if err != nil {
			return
		}
		inetSocketAddressGetFlowinfoFunction, err = inetSocketAddressObject.InvokerNew("get_flowinfo")
	})
	return err
}

// GetFlowinfo is a representation of the C type g_inet_socket_address_get_flowinfo.
func (recv *InetSocketAddress) GetFlowinfo() uint32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := inetSocketAddressGetFlowinfoFunction_Set()
	if err == nil {
		ret = inetSocketAddressGetFlowinfoFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo
}

var inetSocketAddressGetPortFunction *gi.Function
var inetSocketAddressGetPortFunction_Once sync.Once

func inetSocketAddressGetPortFunction_Set() error {
	var err error
	inetSocketAddressGetPortFunction_Once.Do(func() {
		err = inetSocketAddressObject_Set()
		if err != nil {
			return
		}
		inetSocketAddressGetPortFunction, err = inetSocketAddressObject.InvokerNew("get_port")
	})
	return err
}

// GetPort is a representation of the C type g_inet_socket_address_get_port.
func (recv *InetSocketAddress) GetPort() uint16 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := inetSocketAddressGetPortFunction_Set()
	if err == nil {
		ret = inetSocketAddressGetPortFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint16()

	return retGo
}

var inetSocketAddressGetScopeIdFunction *gi.Function
var inetSocketAddressGetScopeIdFunction_Once sync.Once

func inetSocketAddressGetScopeIdFunction_Set() error {
	var err error
	inetSocketAddressGetScopeIdFunction_Once.Do(func() {
		err = inetSocketAddressObject_Set()
		if err != nil {
			return
		}
		inetSocketAddressGetScopeIdFunction, err = inetSocketAddressObject.InvokerNew("get_scope_id")
	})
	return err
}

// GetScopeId is a representation of the C type g_inet_socket_address_get_scope_id.
func (recv *InetSocketAddress) GetScopeId() uint32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := inetSocketAddressGetScopeIdFunction_Set()
	if err == nil {
		ret = inetSocketAddressGetScopeIdFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo
}

var inputStreamObject *gi.Object
var inputStreamObject_Once sync.Once

func inputStreamObject_Set() error {
	var err error
	inputStreamObject_Once.Do(func() {
		inputStreamObject, err = gi.ObjectNew("Gio", "InputStream")
	})
	return err
}

type InputStream struct {
	native unsafe.Pointer
}

func InputStreamNewFromNative(native unsafe.Pointer) *InputStream {
	instance := &InputStream{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *InputStream) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToInputStream down casts any arbitrary Object to InputStream.
Exercise care, as this is a potentially dangerous function
if the Object is not a InputStream.
*/
func CastToInputStream(object *gobject.Object) *InputStream {
	return InputStreamNewFromNative(object.Native())
}

// Equals compares this InputStream with another InputStream, and returns true if they represent the same GObject.
func (recv *InputStream) Equals(other *InputStream) bool {
	return other.Native() == recv.Native()
}

func (recv *InputStream) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *InputStream) FieldParentInstance() *gobject.Object {
	argValue := gi.ObjectFieldGet(inputStreamObject, recv.Native(), "parent_instance")
	value := gobject.ObjectNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *InputStream) SetFieldParentInstance(value *gobject.Object) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(inputStreamObject, recv.Native(), "parent_instance", argValue)
}

var inputStreamClearPendingFunction *gi.Function
var inputStreamClearPendingFunction_Once sync.Once

func inputStreamClearPendingFunction_Set() error {
	var err error
	inputStreamClearPendingFunction_Once.Do(func() {
		err = inputStreamObject_Set()
		if err != nil {
			return
		}
		inputStreamClearPendingFunction, err = inputStreamObject.InvokerNew("clear_pending")
	})
	return err
}

// ClearPending is a representation of the C type g_input_stream_clear_pending.
func (recv *InputStream) ClearPending() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := inputStreamClearPendingFunction_Set()
	if err == nil {
		inputStreamClearPendingFunction.Invoke(inArgs[:], nil)
	}

	return
}

var inputStreamCloseFunction *gi.Function
var inputStreamCloseFunction_Once sync.Once

func inputStreamCloseFunction_Set() error {
	var err error
	inputStreamCloseFunction_Once.Do(func() {
		err = inputStreamObject_Set()
		if err != nil {
			return
		}
		inputStreamCloseFunction, err = inputStreamObject.InvokerNew("close")
	})
	return err
}

// Close is a representation of the C type g_input_stream_close.
func (recv *InputStream) Close(cancellable *Cancellable) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(cancellable.Native())

	var ret gi.Argument

	err := inputStreamCloseFunction_Set()
	if err == nil {
		ret = inputStreamCloseFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'g_input_stream_close_async' : parameter 'callback' of type 'AsyncReadyCallback' not supported

// UNSUPPORTED : C value 'g_input_stream_close_finish' : parameter 'result' of type 'AsyncResult' not supported

var inputStreamHasPendingFunction *gi.Function
var inputStreamHasPendingFunction_Once sync.Once

func inputStreamHasPendingFunction_Set() error {
	var err error
	inputStreamHasPendingFunction_Once.Do(func() {
		err = inputStreamObject_Set()
		if err != nil {
			return
		}
		inputStreamHasPendingFunction, err = inputStreamObject.InvokerNew("has_pending")
	})
	return err
}

// HasPending is a representation of the C type g_input_stream_has_pending.
func (recv *InputStream) HasPending() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := inputStreamHasPendingFunction_Set()
	if err == nil {
		ret = inputStreamHasPendingFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var inputStreamIsClosedFunction *gi.Function
var inputStreamIsClosedFunction_Once sync.Once

func inputStreamIsClosedFunction_Set() error {
	var err error
	inputStreamIsClosedFunction_Once.Do(func() {
		err = inputStreamObject_Set()
		if err != nil {
			return
		}
		inputStreamIsClosedFunction, err = inputStreamObject.InvokerNew("is_closed")
	})
	return err
}

// IsClosed is a representation of the C type g_input_stream_is_closed.
func (recv *InputStream) IsClosed() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := inputStreamIsClosedFunction_Set()
	if err == nil {
		ret = inputStreamIsClosedFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'g_input_stream_read' : parameter 'buffer' of type 'nil' not supported

// UNSUPPORTED : C value 'g_input_stream_read_all' : parameter 'buffer' of type 'nil' not supported

// UNSUPPORTED : C value 'g_input_stream_read_all_async' : parameter 'buffer' of type 'nil' not supported

// UNSUPPORTED : C value 'g_input_stream_read_all_finish' : parameter 'result' of type 'AsyncResult' not supported

// UNSUPPORTED : C value 'g_input_stream_read_async' : parameter 'buffer' of type 'nil' not supported

// UNSUPPORTED : C value 'g_input_stream_read_bytes' : return type 'GLib.Bytes' not supported

// UNSUPPORTED : C value 'g_input_stream_read_bytes_async' : parameter 'callback' of type 'AsyncReadyCallback' not supported

// UNSUPPORTED : C value 'g_input_stream_read_bytes_finish' : parameter 'result' of type 'AsyncResult' not supported

// UNSUPPORTED : C value 'g_input_stream_read_finish' : parameter 'result' of type 'AsyncResult' not supported

var inputStreamSetPendingFunction *gi.Function
var inputStreamSetPendingFunction_Once sync.Once

func inputStreamSetPendingFunction_Set() error {
	var err error
	inputStreamSetPendingFunction_Once.Do(func() {
		err = inputStreamObject_Set()
		if err != nil {
			return
		}
		inputStreamSetPendingFunction, err = inputStreamObject.InvokerNew("set_pending")
	})
	return err
}

// SetPending is a representation of the C type g_input_stream_set_pending.
func (recv *InputStream) SetPending() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := inputStreamSetPendingFunction_Set()
	if err == nil {
		ret = inputStreamSetPendingFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var inputStreamSkipFunction *gi.Function
var inputStreamSkipFunction_Once sync.Once

func inputStreamSkipFunction_Set() error {
	var err error
	inputStreamSkipFunction_Once.Do(func() {
		err = inputStreamObject_Set()
		if err != nil {
			return
		}
		inputStreamSkipFunction, err = inputStreamObject.InvokerNew("skip")
	})
	return err
}

// Skip is a representation of the C type g_input_stream_skip.
func (recv *InputStream) Skip(count uint64, cancellable *Cancellable) int32 {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetUint64(count)
	inArgs[2].SetPointer(cancellable.Native())

	var ret gi.Argument

	err := inputStreamSkipFunction_Set()
	if err == nil {
		ret = inputStreamSkipFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

// UNSUPPORTED : C value 'g_input_stream_skip_async' : parameter 'callback' of type 'AsyncReadyCallback' not supported

// UNSUPPORTED : C value 'g_input_stream_skip_finish' : parameter 'result' of type 'AsyncResult' not supported

var listStoreObject *gi.Object
var listStoreObject_Once sync.Once

func listStoreObject_Set() error {
	var err error
	listStoreObject_Once.Do(func() {
		listStoreObject, err = gi.ObjectNew("Gio", "ListStore")
	})
	return err
}

type ListStore struct {
	native unsafe.Pointer
}

func ListStoreNewFromNative(native unsafe.Pointer) *ListStore {
	instance := &ListStore{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *ListStore) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToListStore down casts any arbitrary Object to ListStore.
Exercise care, as this is a potentially dangerous function
if the Object is not a ListStore.
*/
func CastToListStore(object *gobject.Object) *ListStore {
	return ListStoreNewFromNative(object.Native())
}

// Equals compares this ListStore with another ListStore, and returns true if they represent the same GObject.
func (recv *ListStore) Equals(other *ListStore) bool {
	return other.Native() == recv.Native()
}

func (recv *ListStore) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'g_list_store_new' : parameter 'item_type' of type 'GType' not supported

var listStoreAppendFunction *gi.Function
var listStoreAppendFunction_Once sync.Once

func listStoreAppendFunction_Set() error {
	var err error
	listStoreAppendFunction_Once.Do(func() {
		err = listStoreObject_Set()
		if err != nil {
			return
		}
		listStoreAppendFunction, err = listStoreObject.InvokerNew("append")
	})
	return err
}

// Append is a representation of the C type g_list_store_append.
func (recv *ListStore) Append(item *gobject.Object) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(item.Native())

	err := listStoreAppendFunction_Set()
	if err == nil {
		listStoreAppendFunction.Invoke(inArgs[:], nil)
	}

	return
}

var listStoreInsertFunction *gi.Function
var listStoreInsertFunction_Once sync.Once

func listStoreInsertFunction_Set() error {
	var err error
	listStoreInsertFunction_Once.Do(func() {
		err = listStoreObject_Set()
		if err != nil {
			return
		}
		listStoreInsertFunction, err = listStoreObject.InvokerNew("insert")
	})
	return err
}

// Insert is a representation of the C type g_list_store_insert.
func (recv *ListStore) Insert(position uint32, item *gobject.Object) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetUint32(position)
	inArgs[2].SetPointer(item.Native())

	err := listStoreInsertFunction_Set()
	if err == nil {
		listStoreInsertFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'g_list_store_insert_sorted' : parameter 'compare_func' of type 'GLib.CompareDataFunc' not supported

var listStoreRemoveFunction *gi.Function
var listStoreRemoveFunction_Once sync.Once

func listStoreRemoveFunction_Set() error {
	var err error
	listStoreRemoveFunction_Once.Do(func() {
		err = listStoreObject_Set()
		if err != nil {
			return
		}
		listStoreRemoveFunction, err = listStoreObject.InvokerNew("remove")
	})
	return err
}

// Remove is a representation of the C type g_list_store_remove.
func (recv *ListStore) Remove(position uint32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetUint32(position)

	err := listStoreRemoveFunction_Set()
	if err == nil {
		listStoreRemoveFunction.Invoke(inArgs[:], nil)
	}

	return
}

var listStoreRemoveAllFunction *gi.Function
var listStoreRemoveAllFunction_Once sync.Once

func listStoreRemoveAllFunction_Set() error {
	var err error
	listStoreRemoveAllFunction_Once.Do(func() {
		err = listStoreObject_Set()
		if err != nil {
			return
		}
		listStoreRemoveAllFunction, err = listStoreObject.InvokerNew("remove_all")
	})
	return err
}

// RemoveAll is a representation of the C type g_list_store_remove_all.
func (recv *ListStore) RemoveAll() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := listStoreRemoveAllFunction_Set()
	if err == nil {
		listStoreRemoveAllFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'g_list_store_sort' : parameter 'compare_func' of type 'GLib.CompareDataFunc' not supported

// UNSUPPORTED : C value 'g_list_store_splice' : parameter 'additions' of type 'nil' not supported

var memoryInputStreamObject *gi.Object
var memoryInputStreamObject_Once sync.Once

func memoryInputStreamObject_Set() error {
	var err error
	memoryInputStreamObject_Once.Do(func() {
		memoryInputStreamObject, err = gi.ObjectNew("Gio", "MemoryInputStream")
	})
	return err
}

type MemoryInputStream struct {
	native unsafe.Pointer
}

func MemoryInputStreamNewFromNative(native unsafe.Pointer) *MemoryInputStream {
	instance := &MemoryInputStream{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// InputStream upcasts to *InputStream
func (recv *MemoryInputStream) InputStream() *InputStream {
	return InputStreamNewFromNative(recv.Native())
}

// Object upcasts to *Object
func (recv *MemoryInputStream) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToMemoryInputStream down casts any arbitrary Object to MemoryInputStream.
Exercise care, as this is a potentially dangerous function
if the Object is not a MemoryInputStream.
*/
func CastToMemoryInputStream(object *gobject.Object) *MemoryInputStream {
	return MemoryInputStreamNewFromNative(object.Native())
}

// Equals compares this MemoryInputStream with another MemoryInputStream, and returns true if they represent the same GObject.
func (recv *MemoryInputStream) Equals(other *MemoryInputStream) bool {
	return other.Native() == recv.Native()
}

func (recv *MemoryInputStream) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *MemoryInputStream) FieldParentInstance() *InputStream {
	argValue := gi.ObjectFieldGet(memoryInputStreamObject, recv.Native(), "parent_instance")
	value := InputStreamNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *MemoryInputStream) SetFieldParentInstance(value *InputStream) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(memoryInputStreamObject, recv.Native(), "parent_instance", argValue)
}

var memoryInputStreamNewFunction *gi.Function
var memoryInputStreamNewFunction_Once sync.Once

func memoryInputStreamNewFunction_Set() error {
	var err error
	memoryInputStreamNewFunction_Once.Do(func() {
		err = memoryInputStreamObject_Set()
		if err != nil {
			return
		}
		memoryInputStreamNewFunction, err = memoryInputStreamObject.InvokerNew("new")
	})
	return err
}

// MemoryInputStreamNew is a representation of the C type g_memory_input_stream_new.
func MemoryInputStreamNew() *MemoryInputStream {

	var ret gi.Argument

	err := memoryInputStreamNewFunction_Set()
	if err == nil {
		ret = memoryInputStreamNewFunction.Invoke(nil, nil)
	}

	retGo := MemoryInputStreamNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

// UNSUPPORTED : C value 'g_memory_input_stream_new_from_bytes' : parameter 'bytes' of type 'GLib.Bytes' not supported

// UNSUPPORTED : C value 'g_memory_input_stream_new_from_data' : parameter 'data' of type 'nil' not supported

// UNSUPPORTED : C value 'g_memory_input_stream_add_bytes' : parameter 'bytes' of type 'GLib.Bytes' not supported

// UNSUPPORTED : C value 'g_memory_input_stream_add_data' : parameter 'data' of type 'nil' not supported

var memoryOutputStreamObject *gi.Object
var memoryOutputStreamObject_Once sync.Once

func memoryOutputStreamObject_Set() error {
	var err error
	memoryOutputStreamObject_Once.Do(func() {
		memoryOutputStreamObject, err = gi.ObjectNew("Gio", "MemoryOutputStream")
	})
	return err
}

type MemoryOutputStream struct {
	native unsafe.Pointer
}

func MemoryOutputStreamNewFromNative(native unsafe.Pointer) *MemoryOutputStream {
	instance := &MemoryOutputStream{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// OutputStream upcasts to *OutputStream
func (recv *MemoryOutputStream) OutputStream() *OutputStream {
	return OutputStreamNewFromNative(recv.Native())
}

// Object upcasts to *Object
func (recv *MemoryOutputStream) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToMemoryOutputStream down casts any arbitrary Object to MemoryOutputStream.
Exercise care, as this is a potentially dangerous function
if the Object is not a MemoryOutputStream.
*/
func CastToMemoryOutputStream(object *gobject.Object) *MemoryOutputStream {
	return MemoryOutputStreamNewFromNative(object.Native())
}

// Equals compares this MemoryOutputStream with another MemoryOutputStream, and returns true if they represent the same GObject.
func (recv *MemoryOutputStream) Equals(other *MemoryOutputStream) bool {
	return other.Native() == recv.Native()
}

func (recv *MemoryOutputStream) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *MemoryOutputStream) FieldParentInstance() *OutputStream {
	argValue := gi.ObjectFieldGet(memoryOutputStreamObject, recv.Native(), "parent_instance")
	value := OutputStreamNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *MemoryOutputStream) SetFieldParentInstance(value *OutputStream) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(memoryOutputStreamObject, recv.Native(), "parent_instance", argValue)
}

// UNSUPPORTED : C value 'g_memory_output_stream_new' : parameter 'data' of type 'gpointer' not supported

var memoryOutputStreamNewResizableFunction *gi.Function
var memoryOutputStreamNewResizableFunction_Once sync.Once

func memoryOutputStreamNewResizableFunction_Set() error {
	var err error
	memoryOutputStreamNewResizableFunction_Once.Do(func() {
		err = memoryOutputStreamObject_Set()
		if err != nil {
			return
		}
		memoryOutputStreamNewResizableFunction, err = memoryOutputStreamObject.InvokerNew("new_resizable")
	})
	return err
}

// MemoryOutputStreamNewResizable is a representation of the C type g_memory_output_stream_new_resizable.
func MemoryOutputStreamNewResizable() *MemoryOutputStream {

	var ret gi.Argument

	err := memoryOutputStreamNewResizableFunction_Set()
	if err == nil {
		ret = memoryOutputStreamNewResizableFunction.Invoke(nil, nil)
	}

	retGo := MemoryOutputStreamNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

// UNSUPPORTED : C value 'g_memory_output_stream_get_data' : return type 'gpointer' not supported

var memoryOutputStreamGetDataSizeFunction *gi.Function
var memoryOutputStreamGetDataSizeFunction_Once sync.Once

func memoryOutputStreamGetDataSizeFunction_Set() error {
	var err error
	memoryOutputStreamGetDataSizeFunction_Once.Do(func() {
		err = memoryOutputStreamObject_Set()
		if err != nil {
			return
		}
		memoryOutputStreamGetDataSizeFunction, err = memoryOutputStreamObject.InvokerNew("get_data_size")
	})
	return err
}

// GetDataSize is a representation of the C type g_memory_output_stream_get_data_size.
func (recv *MemoryOutputStream) GetDataSize() uint64 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := memoryOutputStreamGetDataSizeFunction_Set()
	if err == nil {
		ret = memoryOutputStreamGetDataSizeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint64()

	return retGo
}

var memoryOutputStreamGetSizeFunction *gi.Function
var memoryOutputStreamGetSizeFunction_Once sync.Once

func memoryOutputStreamGetSizeFunction_Set() error {
	var err error
	memoryOutputStreamGetSizeFunction_Once.Do(func() {
		err = memoryOutputStreamObject_Set()
		if err != nil {
			return
		}
		memoryOutputStreamGetSizeFunction, err = memoryOutputStreamObject.InvokerNew("get_size")
	})
	return err
}

// GetSize is a representation of the C type g_memory_output_stream_get_size.
func (recv *MemoryOutputStream) GetSize() uint64 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := memoryOutputStreamGetSizeFunction_Set()
	if err == nil {
		ret = memoryOutputStreamGetSizeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint64()

	return retGo
}

// UNSUPPORTED : C value 'g_memory_output_stream_steal_as_bytes' : return type 'GLib.Bytes' not supported

// UNSUPPORTED : C value 'g_memory_output_stream_steal_data' : return type 'gpointer' not supported

var menuObject *gi.Object
var menuObject_Once sync.Once

func menuObject_Set() error {
	var err error
	menuObject_Once.Do(func() {
		menuObject, err = gi.ObjectNew("Gio", "Menu")
	})
	return err
}

type Menu struct {
	native unsafe.Pointer
}

func MenuNewFromNative(native unsafe.Pointer) *Menu {
	instance := &Menu{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// MenuModel upcasts to *MenuModel
func (recv *Menu) MenuModel() *MenuModel {
	return MenuModelNewFromNative(recv.Native())
}

// Object upcasts to *Object
func (recv *Menu) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToMenu down casts any arbitrary Object to Menu.
Exercise care, as this is a potentially dangerous function
if the Object is not a Menu.
*/
func CastToMenu(object *gobject.Object) *Menu {
	return MenuNewFromNative(object.Native())
}

// Equals compares this Menu with another Menu, and returns true if they represent the same GObject.
func (recv *Menu) Equals(other *Menu) bool {
	return other.Native() == recv.Native()
}

func (recv *Menu) Native() unsafe.Pointer {
	return recv.native
}

var menuNewFunction *gi.Function
var menuNewFunction_Once sync.Once

func menuNewFunction_Set() error {
	var err error
	menuNewFunction_Once.Do(func() {
		err = menuObject_Set()
		if err != nil {
			return
		}
		menuNewFunction, err = menuObject.InvokerNew("new")
	})
	return err
}

// MenuNew is a representation of the C type g_menu_new.
func MenuNew() *Menu {

	var ret gi.Argument

	err := menuNewFunction_Set()
	if err == nil {
		ret = menuNewFunction.Invoke(nil, nil)
	}

	retGo := MenuNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

var menuAppendFunction *gi.Function
var menuAppendFunction_Once sync.Once

func menuAppendFunction_Set() error {
	var err error
	menuAppendFunction_Once.Do(func() {
		err = menuObject_Set()
		if err != nil {
			return
		}
		menuAppendFunction, err = menuObject.InvokerNew("append")
	})
	return err
}

// Append is a representation of the C type g_menu_append.
func (recv *Menu) Append(label string, detailedAction string) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(label)
	inArgs[2].SetString(detailedAction)

	err := menuAppendFunction_Set()
	if err == nil {
		menuAppendFunction.Invoke(inArgs[:], nil)
	}

	return
}

var menuAppendItemFunction *gi.Function
var menuAppendItemFunction_Once sync.Once

func menuAppendItemFunction_Set() error {
	var err error
	menuAppendItemFunction_Once.Do(func() {
		err = menuObject_Set()
		if err != nil {
			return
		}
		menuAppendItemFunction, err = menuObject.InvokerNew("append_item")
	})
	return err
}

// AppendItem is a representation of the C type g_menu_append_item.
func (recv *Menu) AppendItem(item *MenuItem) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(item.Native())

	err := menuAppendItemFunction_Set()
	if err == nil {
		menuAppendItemFunction.Invoke(inArgs[:], nil)
	}

	return
}

var menuAppendSectionFunction *gi.Function
var menuAppendSectionFunction_Once sync.Once

func menuAppendSectionFunction_Set() error {
	var err error
	menuAppendSectionFunction_Once.Do(func() {
		err = menuObject_Set()
		if err != nil {
			return
		}
		menuAppendSectionFunction, err = menuObject.InvokerNew("append_section")
	})
	return err
}

// AppendSection is a representation of the C type g_menu_append_section.
func (recv *Menu) AppendSection(label string, section *MenuModel) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(label)
	inArgs[2].SetPointer(section.Native())

	err := menuAppendSectionFunction_Set()
	if err == nil {
		menuAppendSectionFunction.Invoke(inArgs[:], nil)
	}

	return
}

var menuAppendSubmenuFunction *gi.Function
var menuAppendSubmenuFunction_Once sync.Once

func menuAppendSubmenuFunction_Set() error {
	var err error
	menuAppendSubmenuFunction_Once.Do(func() {
		err = menuObject_Set()
		if err != nil {
			return
		}
		menuAppendSubmenuFunction, err = menuObject.InvokerNew("append_submenu")
	})
	return err
}

// AppendSubmenu is a representation of the C type g_menu_append_submenu.
func (recv *Menu) AppendSubmenu(label string, submenu *MenuModel) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(label)
	inArgs[2].SetPointer(submenu.Native())

	err := menuAppendSubmenuFunction_Set()
	if err == nil {
		menuAppendSubmenuFunction.Invoke(inArgs[:], nil)
	}

	return
}

var menuFreezeFunction *gi.Function
var menuFreezeFunction_Once sync.Once

func menuFreezeFunction_Set() error {
	var err error
	menuFreezeFunction_Once.Do(func() {
		err = menuObject_Set()
		if err != nil {
			return
		}
		menuFreezeFunction, err = menuObject.InvokerNew("freeze")
	})
	return err
}

// Freeze is a representation of the C type g_menu_freeze.
func (recv *Menu) Freeze() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := menuFreezeFunction_Set()
	if err == nil {
		menuFreezeFunction.Invoke(inArgs[:], nil)
	}

	return
}

var menuInsertFunction *gi.Function
var menuInsertFunction_Once sync.Once

func menuInsertFunction_Set() error {
	var err error
	menuInsertFunction_Once.Do(func() {
		err = menuObject_Set()
		if err != nil {
			return
		}
		menuInsertFunction, err = menuObject.InvokerNew("insert")
	})
	return err
}

// Insert is a representation of the C type g_menu_insert.
func (recv *Menu) Insert(position int32, label string, detailedAction string) {
	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(position)
	inArgs[2].SetString(label)
	inArgs[3].SetString(detailedAction)

	err := menuInsertFunction_Set()
	if err == nil {
		menuInsertFunction.Invoke(inArgs[:], nil)
	}

	return
}

var menuInsertItemFunction *gi.Function
var menuInsertItemFunction_Once sync.Once

func menuInsertItemFunction_Set() error {
	var err error
	menuInsertItemFunction_Once.Do(func() {
		err = menuObject_Set()
		if err != nil {
			return
		}
		menuInsertItemFunction, err = menuObject.InvokerNew("insert_item")
	})
	return err
}

// InsertItem is a representation of the C type g_menu_insert_item.
func (recv *Menu) InsertItem(position int32, item *MenuItem) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(position)
	inArgs[2].SetPointer(item.Native())

	err := menuInsertItemFunction_Set()
	if err == nil {
		menuInsertItemFunction.Invoke(inArgs[:], nil)
	}

	return
}

var menuInsertSectionFunction *gi.Function
var menuInsertSectionFunction_Once sync.Once

func menuInsertSectionFunction_Set() error {
	var err error
	menuInsertSectionFunction_Once.Do(func() {
		err = menuObject_Set()
		if err != nil {
			return
		}
		menuInsertSectionFunction, err = menuObject.InvokerNew("insert_section")
	})
	return err
}

// InsertSection is a representation of the C type g_menu_insert_section.
func (recv *Menu) InsertSection(position int32, label string, section *MenuModel) {
	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(position)
	inArgs[2].SetString(label)
	inArgs[3].SetPointer(section.Native())

	err := menuInsertSectionFunction_Set()
	if err == nil {
		menuInsertSectionFunction.Invoke(inArgs[:], nil)
	}

	return
}

var menuInsertSubmenuFunction *gi.Function
var menuInsertSubmenuFunction_Once sync.Once

func menuInsertSubmenuFunction_Set() error {
	var err error
	menuInsertSubmenuFunction_Once.Do(func() {
		err = menuObject_Set()
		if err != nil {
			return
		}
		menuInsertSubmenuFunction, err = menuObject.InvokerNew("insert_submenu")
	})
	return err
}

// InsertSubmenu is a representation of the C type g_menu_insert_submenu.
func (recv *Menu) InsertSubmenu(position int32, label string, submenu *MenuModel) {
	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(position)
	inArgs[2].SetString(label)
	inArgs[3].SetPointer(submenu.Native())

	err := menuInsertSubmenuFunction_Set()
	if err == nil {
		menuInsertSubmenuFunction.Invoke(inArgs[:], nil)
	}

	return
}

var menuPrependFunction *gi.Function
var menuPrependFunction_Once sync.Once

func menuPrependFunction_Set() error {
	var err error
	menuPrependFunction_Once.Do(func() {
		err = menuObject_Set()
		if err != nil {
			return
		}
		menuPrependFunction, err = menuObject.InvokerNew("prepend")
	})
	return err
}

// Prepend is a representation of the C type g_menu_prepend.
func (recv *Menu) Prepend(label string, detailedAction string) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(label)
	inArgs[2].SetString(detailedAction)

	err := menuPrependFunction_Set()
	if err == nil {
		menuPrependFunction.Invoke(inArgs[:], nil)
	}

	return
}

var menuPrependItemFunction *gi.Function
var menuPrependItemFunction_Once sync.Once

func menuPrependItemFunction_Set() error {
	var err error
	menuPrependItemFunction_Once.Do(func() {
		err = menuObject_Set()
		if err != nil {
			return
		}
		menuPrependItemFunction, err = menuObject.InvokerNew("prepend_item")
	})
	return err
}

// PrependItem is a representation of the C type g_menu_prepend_item.
func (recv *Menu) PrependItem(item *MenuItem) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(item.Native())

	err := menuPrependItemFunction_Set()
	if err == nil {
		menuPrependItemFunction.Invoke(inArgs[:], nil)
	}

	return
}

var menuPrependSectionFunction *gi.Function
var menuPrependSectionFunction_Once sync.Once

func menuPrependSectionFunction_Set() error {
	var err error
	menuPrependSectionFunction_Once.Do(func() {
		err = menuObject_Set()
		if err != nil {
			return
		}
		menuPrependSectionFunction, err = menuObject.InvokerNew("prepend_section")
	})
	return err
}

// PrependSection is a representation of the C type g_menu_prepend_section.
func (recv *Menu) PrependSection(label string, section *MenuModel) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(label)
	inArgs[2].SetPointer(section.Native())

	err := menuPrependSectionFunction_Set()
	if err == nil {
		menuPrependSectionFunction.Invoke(inArgs[:], nil)
	}

	return
}

var menuPrependSubmenuFunction *gi.Function
var menuPrependSubmenuFunction_Once sync.Once

func menuPrependSubmenuFunction_Set() error {
	var err error
	menuPrependSubmenuFunction_Once.Do(func() {
		err = menuObject_Set()
		if err != nil {
			return
		}
		menuPrependSubmenuFunction, err = menuObject.InvokerNew("prepend_submenu")
	})
	return err
}

// PrependSubmenu is a representation of the C type g_menu_prepend_submenu.
func (recv *Menu) PrependSubmenu(label string, submenu *MenuModel) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(label)
	inArgs[2].SetPointer(submenu.Native())

	err := menuPrependSubmenuFunction_Set()
	if err == nil {
		menuPrependSubmenuFunction.Invoke(inArgs[:], nil)
	}

	return
}

var menuRemoveFunction *gi.Function
var menuRemoveFunction_Once sync.Once

func menuRemoveFunction_Set() error {
	var err error
	menuRemoveFunction_Once.Do(func() {
		err = menuObject_Set()
		if err != nil {
			return
		}
		menuRemoveFunction, err = menuObject.InvokerNew("remove")
	})
	return err
}

// Remove is a representation of the C type g_menu_remove.
func (recv *Menu) Remove(position int32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(position)

	err := menuRemoveFunction_Set()
	if err == nil {
		menuRemoveFunction.Invoke(inArgs[:], nil)
	}

	return
}

var menuRemoveAllFunction *gi.Function
var menuRemoveAllFunction_Once sync.Once

func menuRemoveAllFunction_Set() error {
	var err error
	menuRemoveAllFunction_Once.Do(func() {
		err = menuObject_Set()
		if err != nil {
			return
		}
		menuRemoveAllFunction, err = menuObject.InvokerNew("remove_all")
	})
	return err
}

// RemoveAll is a representation of the C type g_menu_remove_all.
func (recv *Menu) RemoveAll() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := menuRemoveAllFunction_Set()
	if err == nil {
		menuRemoveAllFunction.Invoke(inArgs[:], nil)
	}

	return
}

var menuAttributeIterObject *gi.Object
var menuAttributeIterObject_Once sync.Once

func menuAttributeIterObject_Set() error {
	var err error
	menuAttributeIterObject_Once.Do(func() {
		menuAttributeIterObject, err = gi.ObjectNew("Gio", "MenuAttributeIter")
	})
	return err
}

type MenuAttributeIter struct {
	native unsafe.Pointer
}

func MenuAttributeIterNewFromNative(native unsafe.Pointer) *MenuAttributeIter {
	instance := &MenuAttributeIter{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *MenuAttributeIter) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToMenuAttributeIter down casts any arbitrary Object to MenuAttributeIter.
Exercise care, as this is a potentially dangerous function
if the Object is not a MenuAttributeIter.
*/
func CastToMenuAttributeIter(object *gobject.Object) *MenuAttributeIter {
	return MenuAttributeIterNewFromNative(object.Native())
}

// Equals compares this MenuAttributeIter with another MenuAttributeIter, and returns true if they represent the same GObject.
func (recv *MenuAttributeIter) Equals(other *MenuAttributeIter) bool {
	return other.Native() == recv.Native()
}

func (recv *MenuAttributeIter) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *MenuAttributeIter) FieldParentInstance() *gobject.Object {
	argValue := gi.ObjectFieldGet(menuAttributeIterObject, recv.Native(), "parent_instance")
	value := gobject.ObjectNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *MenuAttributeIter) SetFieldParentInstance(value *gobject.Object) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(menuAttributeIterObject, recv.Native(), "parent_instance", argValue)
}

// FieldPriv returns the C field 'priv'.
func (recv *MenuAttributeIter) FieldPriv() *MenuAttributeIterPrivate {
	argValue := gi.ObjectFieldGet(menuAttributeIterObject, recv.Native(), "priv")
	value := MenuAttributeIterPrivateNewFromNative(argValue.Pointer())
	return value
}

// SetFieldPriv sets the value of the C field 'priv'.
func (recv *MenuAttributeIter) SetFieldPriv(value *MenuAttributeIterPrivate) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(menuAttributeIterObject, recv.Native(), "priv", argValue)
}

var menuAttributeIterGetNameFunction *gi.Function
var menuAttributeIterGetNameFunction_Once sync.Once

func menuAttributeIterGetNameFunction_Set() error {
	var err error
	menuAttributeIterGetNameFunction_Once.Do(func() {
		err = menuAttributeIterObject_Set()
		if err != nil {
			return
		}
		menuAttributeIterGetNameFunction, err = menuAttributeIterObject.InvokerNew("get_name")
	})
	return err
}

// GetName is a representation of the C type g_menu_attribute_iter_get_name.
func (recv *MenuAttributeIter) GetName() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := menuAttributeIterGetNameFunction_Set()
	if err == nil {
		ret = menuAttributeIterGetNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'g_menu_attribute_iter_get_next' : parameter 'value' of type 'GLib.Variant' not supported

// UNSUPPORTED : C value 'g_menu_attribute_iter_get_value' : return type 'GLib.Variant' not supported

var menuAttributeIterNextFunction *gi.Function
var menuAttributeIterNextFunction_Once sync.Once

func menuAttributeIterNextFunction_Set() error {
	var err error
	menuAttributeIterNextFunction_Once.Do(func() {
		err = menuAttributeIterObject_Set()
		if err != nil {
			return
		}
		menuAttributeIterNextFunction, err = menuAttributeIterObject.InvokerNew("next")
	})
	return err
}

// Next is a representation of the C type g_menu_attribute_iter_next.
func (recv *MenuAttributeIter) Next() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := menuAttributeIterNextFunction_Set()
	if err == nil {
		ret = menuAttributeIterNextFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var menuItemObject *gi.Object
var menuItemObject_Once sync.Once

func menuItemObject_Set() error {
	var err error
	menuItemObject_Once.Do(func() {
		menuItemObject, err = gi.ObjectNew("Gio", "MenuItem")
	})
	return err
}

type MenuItem struct {
	native unsafe.Pointer
}

func MenuItemNewFromNative(native unsafe.Pointer) *MenuItem {
	instance := &MenuItem{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *MenuItem) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToMenuItem down casts any arbitrary Object to MenuItem.
Exercise care, as this is a potentially dangerous function
if the Object is not a MenuItem.
*/
func CastToMenuItem(object *gobject.Object) *MenuItem {
	return MenuItemNewFromNative(object.Native())
}

// Equals compares this MenuItem with another MenuItem, and returns true if they represent the same GObject.
func (recv *MenuItem) Equals(other *MenuItem) bool {
	return other.Native() == recv.Native()
}

func (recv *MenuItem) Native() unsafe.Pointer {
	return recv.native
}

var menuItemNewFunction *gi.Function
var menuItemNewFunction_Once sync.Once

func menuItemNewFunction_Set() error {
	var err error
	menuItemNewFunction_Once.Do(func() {
		err = menuItemObject_Set()
		if err != nil {
			return
		}
		menuItemNewFunction, err = menuItemObject.InvokerNew("new")
	})
	return err
}

// MenuItemNew is a representation of the C type g_menu_item_new.
func MenuItemNew(label string, detailedAction string) *MenuItem {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(label)
	inArgs[1].SetString(detailedAction)

	var ret gi.Argument

	err := menuItemNewFunction_Set()
	if err == nil {
		ret = menuItemNewFunction.Invoke(inArgs[:], nil)
	}

	retGo := MenuItemNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

var menuItemNewFromModelFunction *gi.Function
var menuItemNewFromModelFunction_Once sync.Once

func menuItemNewFromModelFunction_Set() error {
	var err error
	menuItemNewFromModelFunction_Once.Do(func() {
		err = menuItemObject_Set()
		if err != nil {
			return
		}
		menuItemNewFromModelFunction, err = menuItemObject.InvokerNew("new_from_model")
	})
	return err
}

// MenuItemNewFromModel is a representation of the C type g_menu_item_new_from_model.
func MenuItemNewFromModel(model *MenuModel, itemIndex int32) *MenuItem {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(model.Native())
	inArgs[1].SetInt32(itemIndex)

	var ret gi.Argument

	err := menuItemNewFromModelFunction_Set()
	if err == nil {
		ret = menuItemNewFromModelFunction.Invoke(inArgs[:], nil)
	}

	retGo := MenuItemNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

var menuItemNewSectionFunction *gi.Function
var menuItemNewSectionFunction_Once sync.Once

func menuItemNewSectionFunction_Set() error {
	var err error
	menuItemNewSectionFunction_Once.Do(func() {
		err = menuItemObject_Set()
		if err != nil {
			return
		}
		menuItemNewSectionFunction, err = menuItemObject.InvokerNew("new_section")
	})
	return err
}

// MenuItemNewSection is a representation of the C type g_menu_item_new_section.
func MenuItemNewSection(label string, section *MenuModel) *MenuItem {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(label)
	inArgs[1].SetPointer(section.Native())

	var ret gi.Argument

	err := menuItemNewSectionFunction_Set()
	if err == nil {
		ret = menuItemNewSectionFunction.Invoke(inArgs[:], nil)
	}

	retGo := MenuItemNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

var menuItemNewSubmenuFunction *gi.Function
var menuItemNewSubmenuFunction_Once sync.Once

func menuItemNewSubmenuFunction_Set() error {
	var err error
	menuItemNewSubmenuFunction_Once.Do(func() {
		err = menuItemObject_Set()
		if err != nil {
			return
		}
		menuItemNewSubmenuFunction, err = menuItemObject.InvokerNew("new_submenu")
	})
	return err
}

// MenuItemNewSubmenu is a representation of the C type g_menu_item_new_submenu.
func MenuItemNewSubmenu(label string, submenu *MenuModel) *MenuItem {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(label)
	inArgs[1].SetPointer(submenu.Native())

	var ret gi.Argument

	err := menuItemNewSubmenuFunction_Set()
	if err == nil {
		ret = menuItemNewSubmenuFunction.Invoke(inArgs[:], nil)
	}

	retGo := MenuItemNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

// UNSUPPORTED : C value 'g_menu_item_get_attribute' : parameter '...' of type 'nil' not supported

// UNSUPPORTED : C value 'g_menu_item_get_attribute_value' : parameter 'expected_type' of type 'GLib.VariantType' not supported

var menuItemGetLinkFunction *gi.Function
var menuItemGetLinkFunction_Once sync.Once

func menuItemGetLinkFunction_Set() error {
	var err error
	menuItemGetLinkFunction_Once.Do(func() {
		err = menuItemObject_Set()
		if err != nil {
			return
		}
		menuItemGetLinkFunction, err = menuItemObject.InvokerNew("get_link")
	})
	return err
}

// GetLink is a representation of the C type g_menu_item_get_link.
func (recv *MenuItem) GetLink(link string) *MenuModel {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(link)

	var ret gi.Argument

	err := menuItemGetLinkFunction_Set()
	if err == nil {
		ret = menuItemGetLinkFunction.Invoke(inArgs[:], nil)
	}

	retGo := MenuModelNewFromNative(ret.Pointer())

	return retGo
}

// UNSUPPORTED : C value 'g_menu_item_set_action_and_target' : parameter '...' of type 'nil' not supported

// UNSUPPORTED : C value 'g_menu_item_set_action_and_target_value' : parameter 'target_value' of type 'GLib.Variant' not supported

// UNSUPPORTED : C value 'g_menu_item_set_attribute' : parameter '...' of type 'nil' not supported

// UNSUPPORTED : C value 'g_menu_item_set_attribute_value' : parameter 'value' of type 'GLib.Variant' not supported

var menuItemSetDetailedActionFunction *gi.Function
var menuItemSetDetailedActionFunction_Once sync.Once

func menuItemSetDetailedActionFunction_Set() error {
	var err error
	menuItemSetDetailedActionFunction_Once.Do(func() {
		err = menuItemObject_Set()
		if err != nil {
			return
		}
		menuItemSetDetailedActionFunction, err = menuItemObject.InvokerNew("set_detailed_action")
	})
	return err
}

// SetDetailedAction is a representation of the C type g_menu_item_set_detailed_action.
func (recv *MenuItem) SetDetailedAction(detailedAction string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(detailedAction)

	err := menuItemSetDetailedActionFunction_Set()
	if err == nil {
		menuItemSetDetailedActionFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'g_menu_item_set_icon' : parameter 'icon' of type 'Icon' not supported

var menuItemSetLabelFunction *gi.Function
var menuItemSetLabelFunction_Once sync.Once

func menuItemSetLabelFunction_Set() error {
	var err error
	menuItemSetLabelFunction_Once.Do(func() {
		err = menuItemObject_Set()
		if err != nil {
			return
		}
		menuItemSetLabelFunction, err = menuItemObject.InvokerNew("set_label")
	})
	return err
}

// SetLabel is a representation of the C type g_menu_item_set_label.
func (recv *MenuItem) SetLabel(label string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(label)

	err := menuItemSetLabelFunction_Set()
	if err == nil {
		menuItemSetLabelFunction.Invoke(inArgs[:], nil)
	}

	return
}

var menuItemSetLinkFunction *gi.Function
var menuItemSetLinkFunction_Once sync.Once

func menuItemSetLinkFunction_Set() error {
	var err error
	menuItemSetLinkFunction_Once.Do(func() {
		err = menuItemObject_Set()
		if err != nil {
			return
		}
		menuItemSetLinkFunction, err = menuItemObject.InvokerNew("set_link")
	})
	return err
}

// SetLink is a representation of the C type g_menu_item_set_link.
func (recv *MenuItem) SetLink(link string, model *MenuModel) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(link)
	inArgs[2].SetPointer(model.Native())

	err := menuItemSetLinkFunction_Set()
	if err == nil {
		menuItemSetLinkFunction.Invoke(inArgs[:], nil)
	}

	return
}

var menuItemSetSectionFunction *gi.Function
var menuItemSetSectionFunction_Once sync.Once

func menuItemSetSectionFunction_Set() error {
	var err error
	menuItemSetSectionFunction_Once.Do(func() {
		err = menuItemObject_Set()
		if err != nil {
			return
		}
		menuItemSetSectionFunction, err = menuItemObject.InvokerNew("set_section")
	})
	return err
}

// SetSection is a representation of the C type g_menu_item_set_section.
func (recv *MenuItem) SetSection(section *MenuModel) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(section.Native())

	err := menuItemSetSectionFunction_Set()
	if err == nil {
		menuItemSetSectionFunction.Invoke(inArgs[:], nil)
	}

	return
}

var menuItemSetSubmenuFunction *gi.Function
var menuItemSetSubmenuFunction_Once sync.Once

func menuItemSetSubmenuFunction_Set() error {
	var err error
	menuItemSetSubmenuFunction_Once.Do(func() {
		err = menuItemObject_Set()
		if err != nil {
			return
		}
		menuItemSetSubmenuFunction, err = menuItemObject.InvokerNew("set_submenu")
	})
	return err
}

// SetSubmenu is a representation of the C type g_menu_item_set_submenu.
func (recv *MenuItem) SetSubmenu(submenu *MenuModel) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(submenu.Native())

	err := menuItemSetSubmenuFunction_Set()
	if err == nil {
		menuItemSetSubmenuFunction.Invoke(inArgs[:], nil)
	}

	return
}

var menuLinkIterObject *gi.Object
var menuLinkIterObject_Once sync.Once

func menuLinkIterObject_Set() error {
	var err error
	menuLinkIterObject_Once.Do(func() {
		menuLinkIterObject, err = gi.ObjectNew("Gio", "MenuLinkIter")
	})
	return err
}

type MenuLinkIter struct {
	native unsafe.Pointer
}

func MenuLinkIterNewFromNative(native unsafe.Pointer) *MenuLinkIter {
	instance := &MenuLinkIter{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *MenuLinkIter) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToMenuLinkIter down casts any arbitrary Object to MenuLinkIter.
Exercise care, as this is a potentially dangerous function
if the Object is not a MenuLinkIter.
*/
func CastToMenuLinkIter(object *gobject.Object) *MenuLinkIter {
	return MenuLinkIterNewFromNative(object.Native())
}

// Equals compares this MenuLinkIter with another MenuLinkIter, and returns true if they represent the same GObject.
func (recv *MenuLinkIter) Equals(other *MenuLinkIter) bool {
	return other.Native() == recv.Native()
}

func (recv *MenuLinkIter) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *MenuLinkIter) FieldParentInstance() *gobject.Object {
	argValue := gi.ObjectFieldGet(menuLinkIterObject, recv.Native(), "parent_instance")
	value := gobject.ObjectNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *MenuLinkIter) SetFieldParentInstance(value *gobject.Object) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(menuLinkIterObject, recv.Native(), "parent_instance", argValue)
}

// FieldPriv returns the C field 'priv'.
func (recv *MenuLinkIter) FieldPriv() *MenuLinkIterPrivate {
	argValue := gi.ObjectFieldGet(menuLinkIterObject, recv.Native(), "priv")
	value := MenuLinkIterPrivateNewFromNative(argValue.Pointer())
	return value
}

// SetFieldPriv sets the value of the C field 'priv'.
func (recv *MenuLinkIter) SetFieldPriv(value *MenuLinkIterPrivate) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(menuLinkIterObject, recv.Native(), "priv", argValue)
}

var menuLinkIterGetNameFunction *gi.Function
var menuLinkIterGetNameFunction_Once sync.Once

func menuLinkIterGetNameFunction_Set() error {
	var err error
	menuLinkIterGetNameFunction_Once.Do(func() {
		err = menuLinkIterObject_Set()
		if err != nil {
			return
		}
		menuLinkIterGetNameFunction, err = menuLinkIterObject.InvokerNew("get_name")
	})
	return err
}

// GetName is a representation of the C type g_menu_link_iter_get_name.
func (recv *MenuLinkIter) GetName() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := menuLinkIterGetNameFunction_Set()
	if err == nil {
		ret = menuLinkIterGetNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var menuLinkIterGetNextFunction *gi.Function
var menuLinkIterGetNextFunction_Once sync.Once

func menuLinkIterGetNextFunction_Set() error {
	var err error
	menuLinkIterGetNextFunction_Once.Do(func() {
		err = menuLinkIterObject_Set()
		if err != nil {
			return
		}
		menuLinkIterGetNextFunction, err = menuLinkIterObject.InvokerNew("get_next")
	})
	return err
}

// GetNext is a representation of the C type g_menu_link_iter_get_next.
func (recv *MenuLinkIter) GetNext() (bool, string, *MenuModel) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var outArgs [2]gi.Argument
	var ret gi.Argument

	err := menuLinkIterGetNextFunction_Set()
	if err == nil {
		ret = menuLinkIterGetNextFunction.Invoke(inArgs[:], outArgs[:])
	}

	retGo := ret.Boolean()
	out0 := outArgs[0].String(false)
	out1 := MenuModelNewFromNative(outArgs[1].Pointer())

	return retGo, out0, out1
}

var menuLinkIterGetValueFunction *gi.Function
var menuLinkIterGetValueFunction_Once sync.Once

func menuLinkIterGetValueFunction_Set() error {
	var err error
	menuLinkIterGetValueFunction_Once.Do(func() {
		err = menuLinkIterObject_Set()
		if err != nil {
			return
		}
		menuLinkIterGetValueFunction, err = menuLinkIterObject.InvokerNew("get_value")
	})
	return err
}

// GetValue is a representation of the C type g_menu_link_iter_get_value.
func (recv *MenuLinkIter) GetValue() *MenuModel {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := menuLinkIterGetValueFunction_Set()
	if err == nil {
		ret = menuLinkIterGetValueFunction.Invoke(inArgs[:], nil)
	}

	retGo := MenuModelNewFromNative(ret.Pointer())

	return retGo
}

var menuLinkIterNextFunction *gi.Function
var menuLinkIterNextFunction_Once sync.Once

func menuLinkIterNextFunction_Set() error {
	var err error
	menuLinkIterNextFunction_Once.Do(func() {
		err = menuLinkIterObject_Set()
		if err != nil {
			return
		}
		menuLinkIterNextFunction, err = menuLinkIterObject.InvokerNew("next")
	})
	return err
}

// Next is a representation of the C type g_menu_link_iter_next.
func (recv *MenuLinkIter) Next() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := menuLinkIterNextFunction_Set()
	if err == nil {
		ret = menuLinkIterNextFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var menuModelObject *gi.Object
var menuModelObject_Once sync.Once

func menuModelObject_Set() error {
	var err error
	menuModelObject_Once.Do(func() {
		menuModelObject, err = gi.ObjectNew("Gio", "MenuModel")
	})
	return err
}

type MenuModel struct {
	native unsafe.Pointer
}

func MenuModelNewFromNative(native unsafe.Pointer) *MenuModel {
	instance := &MenuModel{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *MenuModel) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToMenuModel down casts any arbitrary Object to MenuModel.
Exercise care, as this is a potentially dangerous function
if the Object is not a MenuModel.
*/
func CastToMenuModel(object *gobject.Object) *MenuModel {
	return MenuModelNewFromNative(object.Native())
}

// Equals compares this MenuModel with another MenuModel, and returns true if they represent the same GObject.
func (recv *MenuModel) Equals(other *MenuModel) bool {
	return other.Native() == recv.Native()
}

func (recv *MenuModel) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *MenuModel) FieldParentInstance() *gobject.Object {
	argValue := gi.ObjectFieldGet(menuModelObject, recv.Native(), "parent_instance")
	value := gobject.ObjectNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *MenuModel) SetFieldParentInstance(value *gobject.Object) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(menuModelObject, recv.Native(), "parent_instance", argValue)
}

// FieldPriv returns the C field 'priv'.
func (recv *MenuModel) FieldPriv() *MenuModelPrivate {
	argValue := gi.ObjectFieldGet(menuModelObject, recv.Native(), "priv")
	value := MenuModelPrivateNewFromNative(argValue.Pointer())
	return value
}

// SetFieldPriv sets the value of the C field 'priv'.
func (recv *MenuModel) SetFieldPriv(value *MenuModelPrivate) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(menuModelObject, recv.Native(), "priv", argValue)
}

// UNSUPPORTED : C value 'g_menu_model_get_item_attribute' : parameter '...' of type 'nil' not supported

// UNSUPPORTED : C value 'g_menu_model_get_item_attribute_value' : parameter 'expected_type' of type 'GLib.VariantType' not supported

var menuModelGetItemLinkFunction *gi.Function
var menuModelGetItemLinkFunction_Once sync.Once

func menuModelGetItemLinkFunction_Set() error {
	var err error
	menuModelGetItemLinkFunction_Once.Do(func() {
		err = menuModelObject_Set()
		if err != nil {
			return
		}
		menuModelGetItemLinkFunction, err = menuModelObject.InvokerNew("get_item_link")
	})
	return err
}

// GetItemLink is a representation of the C type g_menu_model_get_item_link.
func (recv *MenuModel) GetItemLink(itemIndex int32, link string) *MenuModel {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(itemIndex)
	inArgs[2].SetString(link)

	var ret gi.Argument

	err := menuModelGetItemLinkFunction_Set()
	if err == nil {
		ret = menuModelGetItemLinkFunction.Invoke(inArgs[:], nil)
	}

	retGo := MenuModelNewFromNative(ret.Pointer())

	return retGo
}

var menuModelGetNItemsFunction *gi.Function
var menuModelGetNItemsFunction_Once sync.Once

func menuModelGetNItemsFunction_Set() error {
	var err error
	menuModelGetNItemsFunction_Once.Do(func() {
		err = menuModelObject_Set()
		if err != nil {
			return
		}
		menuModelGetNItemsFunction, err = menuModelObject.InvokerNew("get_n_items")
	})
	return err
}

// GetNItems is a representation of the C type g_menu_model_get_n_items.
func (recv *MenuModel) GetNItems() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := menuModelGetNItemsFunction_Set()
	if err == nil {
		ret = menuModelGetNItemsFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var menuModelIsMutableFunction *gi.Function
var menuModelIsMutableFunction_Once sync.Once

func menuModelIsMutableFunction_Set() error {
	var err error
	menuModelIsMutableFunction_Once.Do(func() {
		err = menuModelObject_Set()
		if err != nil {
			return
		}
		menuModelIsMutableFunction, err = menuModelObject.InvokerNew("is_mutable")
	})
	return err
}

// IsMutable is a representation of the C type g_menu_model_is_mutable.
func (recv *MenuModel) IsMutable() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := menuModelIsMutableFunction_Set()
	if err == nil {
		ret = menuModelIsMutableFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var menuModelItemsChangedFunction *gi.Function
var menuModelItemsChangedFunction_Once sync.Once

func menuModelItemsChangedFunction_Set() error {
	var err error
	menuModelItemsChangedFunction_Once.Do(func() {
		err = menuModelObject_Set()
		if err != nil {
			return
		}
		menuModelItemsChangedFunction, err = menuModelObject.InvokerNew("items_changed")
	})
	return err
}

// ItemsChanged is a representation of the C type g_menu_model_items_changed.
func (recv *MenuModel) ItemsChanged(position int32, removed int32, added int32) {
	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(position)
	inArgs[2].SetInt32(removed)
	inArgs[3].SetInt32(added)

	err := menuModelItemsChangedFunction_Set()
	if err == nil {
		menuModelItemsChangedFunction.Invoke(inArgs[:], nil)
	}

	return
}

var menuModelIterateItemAttributesFunction *gi.Function
var menuModelIterateItemAttributesFunction_Once sync.Once

func menuModelIterateItemAttributesFunction_Set() error {
	var err error
	menuModelIterateItemAttributesFunction_Once.Do(func() {
		err = menuModelObject_Set()
		if err != nil {
			return
		}
		menuModelIterateItemAttributesFunction, err = menuModelObject.InvokerNew("iterate_item_attributes")
	})
	return err
}

// IterateItemAttributes is a representation of the C type g_menu_model_iterate_item_attributes.
func (recv *MenuModel) IterateItemAttributes(itemIndex int32) *MenuAttributeIter {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(itemIndex)

	var ret gi.Argument

	err := menuModelIterateItemAttributesFunction_Set()
	if err == nil {
		ret = menuModelIterateItemAttributesFunction.Invoke(inArgs[:], nil)
	}

	retGo := MenuAttributeIterNewFromNative(ret.Pointer())

	return retGo
}

var menuModelIterateItemLinksFunction *gi.Function
var menuModelIterateItemLinksFunction_Once sync.Once

func menuModelIterateItemLinksFunction_Set() error {
	var err error
	menuModelIterateItemLinksFunction_Once.Do(func() {
		err = menuModelObject_Set()
		if err != nil {
			return
		}
		menuModelIterateItemLinksFunction, err = menuModelObject.InvokerNew("iterate_item_links")
	})
	return err
}

// IterateItemLinks is a representation of the C type g_menu_model_iterate_item_links.
func (recv *MenuModel) IterateItemLinks(itemIndex int32) *MenuLinkIter {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(itemIndex)

	var ret gi.Argument

	err := menuModelIterateItemLinksFunction_Set()
	if err == nil {
		ret = menuModelIterateItemLinksFunction.Invoke(inArgs[:], nil)
	}

	retGo := MenuLinkIterNewFromNative(ret.Pointer())

	return retGo
}

var mountOperationObject *gi.Object
var mountOperationObject_Once sync.Once

func mountOperationObject_Set() error {
	var err error
	mountOperationObject_Once.Do(func() {
		mountOperationObject, err = gi.ObjectNew("Gio", "MountOperation")
	})
	return err
}

type MountOperation struct {
	native unsafe.Pointer
}

func MountOperationNewFromNative(native unsafe.Pointer) *MountOperation {
	instance := &MountOperation{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *MountOperation) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToMountOperation down casts any arbitrary Object to MountOperation.
Exercise care, as this is a potentially dangerous function
if the Object is not a MountOperation.
*/
func CastToMountOperation(object *gobject.Object) *MountOperation {
	return MountOperationNewFromNative(object.Native())
}

// Equals compares this MountOperation with another MountOperation, and returns true if they represent the same GObject.
func (recv *MountOperation) Equals(other *MountOperation) bool {
	return other.Native() == recv.Native()
}

func (recv *MountOperation) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *MountOperation) FieldParentInstance() *gobject.Object {
	argValue := gi.ObjectFieldGet(mountOperationObject, recv.Native(), "parent_instance")
	value := gobject.ObjectNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *MountOperation) SetFieldParentInstance(value *gobject.Object) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(mountOperationObject, recv.Native(), "parent_instance", argValue)
}

// FieldPriv returns the C field 'priv'.
func (recv *MountOperation) FieldPriv() *MountOperationPrivate {
	argValue := gi.ObjectFieldGet(mountOperationObject, recv.Native(), "priv")
	value := MountOperationPrivateNewFromNative(argValue.Pointer())
	return value
}

// SetFieldPriv sets the value of the C field 'priv'.
func (recv *MountOperation) SetFieldPriv(value *MountOperationPrivate) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(mountOperationObject, recv.Native(), "priv", argValue)
}

var mountOperationNewFunction *gi.Function
var mountOperationNewFunction_Once sync.Once

func mountOperationNewFunction_Set() error {
	var err error
	mountOperationNewFunction_Once.Do(func() {
		err = mountOperationObject_Set()
		if err != nil {
			return
		}
		mountOperationNewFunction, err = mountOperationObject.InvokerNew("new")
	})
	return err
}

// MountOperationNew is a representation of the C type g_mount_operation_new.
func MountOperationNew() *MountOperation {

	var ret gi.Argument

	err := mountOperationNewFunction_Set()
	if err == nil {
		ret = mountOperationNewFunction.Invoke(nil, nil)
	}

	retGo := MountOperationNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

var mountOperationGetAnonymousFunction *gi.Function
var mountOperationGetAnonymousFunction_Once sync.Once

func mountOperationGetAnonymousFunction_Set() error {
	var err error
	mountOperationGetAnonymousFunction_Once.Do(func() {
		err = mountOperationObject_Set()
		if err != nil {
			return
		}
		mountOperationGetAnonymousFunction, err = mountOperationObject.InvokerNew("get_anonymous")
	})
	return err
}

// GetAnonymous is a representation of the C type g_mount_operation_get_anonymous.
func (recv *MountOperation) GetAnonymous() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := mountOperationGetAnonymousFunction_Set()
	if err == nil {
		ret = mountOperationGetAnonymousFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var mountOperationGetChoiceFunction *gi.Function
var mountOperationGetChoiceFunction_Once sync.Once

func mountOperationGetChoiceFunction_Set() error {
	var err error
	mountOperationGetChoiceFunction_Once.Do(func() {
		err = mountOperationObject_Set()
		if err != nil {
			return
		}
		mountOperationGetChoiceFunction, err = mountOperationObject.InvokerNew("get_choice")
	})
	return err
}

// GetChoice is a representation of the C type g_mount_operation_get_choice.
func (recv *MountOperation) GetChoice() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := mountOperationGetChoiceFunction_Set()
	if err == nil {
		ret = mountOperationGetChoiceFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var mountOperationGetDomainFunction *gi.Function
var mountOperationGetDomainFunction_Once sync.Once

func mountOperationGetDomainFunction_Set() error {
	var err error
	mountOperationGetDomainFunction_Once.Do(func() {
		err = mountOperationObject_Set()
		if err != nil {
			return
		}
		mountOperationGetDomainFunction, err = mountOperationObject.InvokerNew("get_domain")
	})
	return err
}

// GetDomain is a representation of the C type g_mount_operation_get_domain.
func (recv *MountOperation) GetDomain() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := mountOperationGetDomainFunction_Set()
	if err == nil {
		ret = mountOperationGetDomainFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var mountOperationGetIsTcryptHiddenVolumeFunction *gi.Function
var mountOperationGetIsTcryptHiddenVolumeFunction_Once sync.Once

func mountOperationGetIsTcryptHiddenVolumeFunction_Set() error {
	var err error
	mountOperationGetIsTcryptHiddenVolumeFunction_Once.Do(func() {
		err = mountOperationObject_Set()
		if err != nil {
			return
		}
		mountOperationGetIsTcryptHiddenVolumeFunction, err = mountOperationObject.InvokerNew("get_is_tcrypt_hidden_volume")
	})
	return err
}

// GetIsTcryptHiddenVolume is a representation of the C type g_mount_operation_get_is_tcrypt_hidden_volume.
func (recv *MountOperation) GetIsTcryptHiddenVolume() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := mountOperationGetIsTcryptHiddenVolumeFunction_Set()
	if err == nil {
		ret = mountOperationGetIsTcryptHiddenVolumeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var mountOperationGetIsTcryptSystemVolumeFunction *gi.Function
var mountOperationGetIsTcryptSystemVolumeFunction_Once sync.Once

func mountOperationGetIsTcryptSystemVolumeFunction_Set() error {
	var err error
	mountOperationGetIsTcryptSystemVolumeFunction_Once.Do(func() {
		err = mountOperationObject_Set()
		if err != nil {
			return
		}
		mountOperationGetIsTcryptSystemVolumeFunction, err = mountOperationObject.InvokerNew("get_is_tcrypt_system_volume")
	})
	return err
}

// GetIsTcryptSystemVolume is a representation of the C type g_mount_operation_get_is_tcrypt_system_volume.
func (recv *MountOperation) GetIsTcryptSystemVolume() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := mountOperationGetIsTcryptSystemVolumeFunction_Set()
	if err == nil {
		ret = mountOperationGetIsTcryptSystemVolumeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var mountOperationGetPasswordFunction *gi.Function
var mountOperationGetPasswordFunction_Once sync.Once

func mountOperationGetPasswordFunction_Set() error {
	var err error
	mountOperationGetPasswordFunction_Once.Do(func() {
		err = mountOperationObject_Set()
		if err != nil {
			return
		}
		mountOperationGetPasswordFunction, err = mountOperationObject.InvokerNew("get_password")
	})
	return err
}

// GetPassword is a representation of the C type g_mount_operation_get_password.
func (recv *MountOperation) GetPassword() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := mountOperationGetPasswordFunction_Set()
	if err == nil {
		ret = mountOperationGetPasswordFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var mountOperationGetPasswordSaveFunction *gi.Function
var mountOperationGetPasswordSaveFunction_Once sync.Once

func mountOperationGetPasswordSaveFunction_Set() error {
	var err error
	mountOperationGetPasswordSaveFunction_Once.Do(func() {
		err = mountOperationObject_Set()
		if err != nil {
			return
		}
		mountOperationGetPasswordSaveFunction, err = mountOperationObject.InvokerNew("get_password_save")
	})
	return err
}

// GetPasswordSave is a representation of the C type g_mount_operation_get_password_save.
func (recv *MountOperation) GetPasswordSave() PasswordSave {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := mountOperationGetPasswordSaveFunction_Set()
	if err == nil {
		ret = mountOperationGetPasswordSaveFunction.Invoke(inArgs[:], nil)
	}

	retGo := PasswordSave(ret.Int32())

	return retGo
}

var mountOperationGetPimFunction *gi.Function
var mountOperationGetPimFunction_Once sync.Once

func mountOperationGetPimFunction_Set() error {
	var err error
	mountOperationGetPimFunction_Once.Do(func() {
		err = mountOperationObject_Set()
		if err != nil {
			return
		}
		mountOperationGetPimFunction, err = mountOperationObject.InvokerNew("get_pim")
	})
	return err
}

// GetPim is a representation of the C type g_mount_operation_get_pim.
func (recv *MountOperation) GetPim() uint32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := mountOperationGetPimFunction_Set()
	if err == nil {
		ret = mountOperationGetPimFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo
}

var mountOperationGetUsernameFunction *gi.Function
var mountOperationGetUsernameFunction_Once sync.Once

func mountOperationGetUsernameFunction_Set() error {
	var err error
	mountOperationGetUsernameFunction_Once.Do(func() {
		err = mountOperationObject_Set()
		if err != nil {
			return
		}
		mountOperationGetUsernameFunction, err = mountOperationObject.InvokerNew("get_username")
	})
	return err
}

// GetUsername is a representation of the C type g_mount_operation_get_username.
func (recv *MountOperation) GetUsername() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := mountOperationGetUsernameFunction_Set()
	if err == nil {
		ret = mountOperationGetUsernameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var mountOperationReplyFunction *gi.Function
var mountOperationReplyFunction_Once sync.Once

func mountOperationReplyFunction_Set() error {
	var err error
	mountOperationReplyFunction_Once.Do(func() {
		err = mountOperationObject_Set()
		if err != nil {
			return
		}
		mountOperationReplyFunction, err = mountOperationObject.InvokerNew("reply")
	})
	return err
}

// Reply is a representation of the C type g_mount_operation_reply.
func (recv *MountOperation) Reply(result MountOperationResult) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(int32(result))

	err := mountOperationReplyFunction_Set()
	if err == nil {
		mountOperationReplyFunction.Invoke(inArgs[:], nil)
	}

	return
}

var mountOperationSetAnonymousFunction *gi.Function
var mountOperationSetAnonymousFunction_Once sync.Once

func mountOperationSetAnonymousFunction_Set() error {
	var err error
	mountOperationSetAnonymousFunction_Once.Do(func() {
		err = mountOperationObject_Set()
		if err != nil {
			return
		}
		mountOperationSetAnonymousFunction, err = mountOperationObject.InvokerNew("set_anonymous")
	})
	return err
}

// SetAnonymous is a representation of the C type g_mount_operation_set_anonymous.
func (recv *MountOperation) SetAnonymous(anonymous bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetBoolean(anonymous)

	err := mountOperationSetAnonymousFunction_Set()
	if err == nil {
		mountOperationSetAnonymousFunction.Invoke(inArgs[:], nil)
	}

	return
}

var mountOperationSetChoiceFunction *gi.Function
var mountOperationSetChoiceFunction_Once sync.Once

func mountOperationSetChoiceFunction_Set() error {
	var err error
	mountOperationSetChoiceFunction_Once.Do(func() {
		err = mountOperationObject_Set()
		if err != nil {
			return
		}
		mountOperationSetChoiceFunction, err = mountOperationObject.InvokerNew("set_choice")
	})
	return err
}

// SetChoice is a representation of the C type g_mount_operation_set_choice.
func (recv *MountOperation) SetChoice(choice int32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(choice)

	err := mountOperationSetChoiceFunction_Set()
	if err == nil {
		mountOperationSetChoiceFunction.Invoke(inArgs[:], nil)
	}

	return
}

var mountOperationSetDomainFunction *gi.Function
var mountOperationSetDomainFunction_Once sync.Once

func mountOperationSetDomainFunction_Set() error {
	var err error
	mountOperationSetDomainFunction_Once.Do(func() {
		err = mountOperationObject_Set()
		if err != nil {
			return
		}
		mountOperationSetDomainFunction, err = mountOperationObject.InvokerNew("set_domain")
	})
	return err
}

// SetDomain is a representation of the C type g_mount_operation_set_domain.
func (recv *MountOperation) SetDomain(domain string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(domain)

	err := mountOperationSetDomainFunction_Set()
	if err == nil {
		mountOperationSetDomainFunction.Invoke(inArgs[:], nil)
	}

	return
}

var mountOperationSetIsTcryptHiddenVolumeFunction *gi.Function
var mountOperationSetIsTcryptHiddenVolumeFunction_Once sync.Once

func mountOperationSetIsTcryptHiddenVolumeFunction_Set() error {
	var err error
	mountOperationSetIsTcryptHiddenVolumeFunction_Once.Do(func() {
		err = mountOperationObject_Set()
		if err != nil {
			return
		}
		mountOperationSetIsTcryptHiddenVolumeFunction, err = mountOperationObject.InvokerNew("set_is_tcrypt_hidden_volume")
	})
	return err
}

// SetIsTcryptHiddenVolume is a representation of the C type g_mount_operation_set_is_tcrypt_hidden_volume.
func (recv *MountOperation) SetIsTcryptHiddenVolume(hiddenVolume bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetBoolean(hiddenVolume)

	err := mountOperationSetIsTcryptHiddenVolumeFunction_Set()
	if err == nil {
		mountOperationSetIsTcryptHiddenVolumeFunction.Invoke(inArgs[:], nil)
	}

	return
}

var mountOperationSetIsTcryptSystemVolumeFunction *gi.Function
var mountOperationSetIsTcryptSystemVolumeFunction_Once sync.Once

func mountOperationSetIsTcryptSystemVolumeFunction_Set() error {
	var err error
	mountOperationSetIsTcryptSystemVolumeFunction_Once.Do(func() {
		err = mountOperationObject_Set()
		if err != nil {
			return
		}
		mountOperationSetIsTcryptSystemVolumeFunction, err = mountOperationObject.InvokerNew("set_is_tcrypt_system_volume")
	})
	return err
}

// SetIsTcryptSystemVolume is a representation of the C type g_mount_operation_set_is_tcrypt_system_volume.
func (recv *MountOperation) SetIsTcryptSystemVolume(systemVolume bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetBoolean(systemVolume)

	err := mountOperationSetIsTcryptSystemVolumeFunction_Set()
	if err == nil {
		mountOperationSetIsTcryptSystemVolumeFunction.Invoke(inArgs[:], nil)
	}

	return
}

var mountOperationSetPasswordFunction *gi.Function
var mountOperationSetPasswordFunction_Once sync.Once

func mountOperationSetPasswordFunction_Set() error {
	var err error
	mountOperationSetPasswordFunction_Once.Do(func() {
		err = mountOperationObject_Set()
		if err != nil {
			return
		}
		mountOperationSetPasswordFunction, err = mountOperationObject.InvokerNew("set_password")
	})
	return err
}

// SetPassword is a representation of the C type g_mount_operation_set_password.
func (recv *MountOperation) SetPassword(password string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(password)

	err := mountOperationSetPasswordFunction_Set()
	if err == nil {
		mountOperationSetPasswordFunction.Invoke(inArgs[:], nil)
	}

	return
}

var mountOperationSetPasswordSaveFunction *gi.Function
var mountOperationSetPasswordSaveFunction_Once sync.Once

func mountOperationSetPasswordSaveFunction_Set() error {
	var err error
	mountOperationSetPasswordSaveFunction_Once.Do(func() {
		err = mountOperationObject_Set()
		if err != nil {
			return
		}
		mountOperationSetPasswordSaveFunction, err = mountOperationObject.InvokerNew("set_password_save")
	})
	return err
}

// SetPasswordSave is a representation of the C type g_mount_operation_set_password_save.
func (recv *MountOperation) SetPasswordSave(save PasswordSave) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(int32(save))

	err := mountOperationSetPasswordSaveFunction_Set()
	if err == nil {
		mountOperationSetPasswordSaveFunction.Invoke(inArgs[:], nil)
	}

	return
}

var mountOperationSetPimFunction *gi.Function
var mountOperationSetPimFunction_Once sync.Once

func mountOperationSetPimFunction_Set() error {
	var err error
	mountOperationSetPimFunction_Once.Do(func() {
		err = mountOperationObject_Set()
		if err != nil {
			return
		}
		mountOperationSetPimFunction, err = mountOperationObject.InvokerNew("set_pim")
	})
	return err
}

// SetPim is a representation of the C type g_mount_operation_set_pim.
func (recv *MountOperation) SetPim(pim uint32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetUint32(pim)

	err := mountOperationSetPimFunction_Set()
	if err == nil {
		mountOperationSetPimFunction.Invoke(inArgs[:], nil)
	}

	return
}

var mountOperationSetUsernameFunction *gi.Function
var mountOperationSetUsernameFunction_Once sync.Once

func mountOperationSetUsernameFunction_Set() error {
	var err error
	mountOperationSetUsernameFunction_Once.Do(func() {
		err = mountOperationObject_Set()
		if err != nil {
			return
		}
		mountOperationSetUsernameFunction, err = mountOperationObject.InvokerNew("set_username")
	})
	return err
}

// SetUsername is a representation of the C type g_mount_operation_set_username.
func (recv *MountOperation) SetUsername(username string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(username)

	err := mountOperationSetUsernameFunction_Set()
	if err == nil {
		mountOperationSetUsernameFunction.Invoke(inArgs[:], nil)
	}

	return
}

var nativeSocketAddressObject *gi.Object
var nativeSocketAddressObject_Once sync.Once

func nativeSocketAddressObject_Set() error {
	var err error
	nativeSocketAddressObject_Once.Do(func() {
		nativeSocketAddressObject, err = gi.ObjectNew("Gio", "NativeSocketAddress")
	})
	return err
}

type NativeSocketAddress struct {
	native unsafe.Pointer
}

func NativeSocketAddressNewFromNative(native unsafe.Pointer) *NativeSocketAddress {
	instance := &NativeSocketAddress{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// SocketAddress upcasts to *SocketAddress
func (recv *NativeSocketAddress) SocketAddress() *SocketAddress {
	return SocketAddressNewFromNative(recv.Native())
}

// Object upcasts to *Object
func (recv *NativeSocketAddress) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToNativeSocketAddress down casts any arbitrary Object to NativeSocketAddress.
Exercise care, as this is a potentially dangerous function
if the Object is not a NativeSocketAddress.
*/
func CastToNativeSocketAddress(object *gobject.Object) *NativeSocketAddress {
	return NativeSocketAddressNewFromNative(object.Native())
}

// Equals compares this NativeSocketAddress with another NativeSocketAddress, and returns true if they represent the same GObject.
func (recv *NativeSocketAddress) Equals(other *NativeSocketAddress) bool {
	return other.Native() == recv.Native()
}

func (recv *NativeSocketAddress) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *NativeSocketAddress) FieldParentInstance() *SocketAddress {
	argValue := gi.ObjectFieldGet(nativeSocketAddressObject, recv.Native(), "parent_instance")
	value := SocketAddressNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *NativeSocketAddress) SetFieldParentInstance(value *SocketAddress) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(nativeSocketAddressObject, recv.Native(), "parent_instance", argValue)
}

// UNSUPPORTED : C value 'g_native_socket_address_new' : parameter 'native' of type 'gpointer' not supported

var nativeVolumeMonitorObject *gi.Object
var nativeVolumeMonitorObject_Once sync.Once

func nativeVolumeMonitorObject_Set() error {
	var err error
	nativeVolumeMonitorObject_Once.Do(func() {
		nativeVolumeMonitorObject, err = gi.ObjectNew("Gio", "NativeVolumeMonitor")
	})
	return err
}

type NativeVolumeMonitor struct {
	native unsafe.Pointer
}

func NativeVolumeMonitorNewFromNative(native unsafe.Pointer) *NativeVolumeMonitor {
	instance := &NativeVolumeMonitor{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// VolumeMonitor upcasts to *VolumeMonitor
func (recv *NativeVolumeMonitor) VolumeMonitor() *VolumeMonitor {
	return VolumeMonitorNewFromNative(recv.Native())
}

// Object upcasts to *Object
func (recv *NativeVolumeMonitor) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToNativeVolumeMonitor down casts any arbitrary Object to NativeVolumeMonitor.
Exercise care, as this is a potentially dangerous function
if the Object is not a NativeVolumeMonitor.
*/
func CastToNativeVolumeMonitor(object *gobject.Object) *NativeVolumeMonitor {
	return NativeVolumeMonitorNewFromNative(object.Native())
}

// Equals compares this NativeVolumeMonitor with another NativeVolumeMonitor, and returns true if they represent the same GObject.
func (recv *NativeVolumeMonitor) Equals(other *NativeVolumeMonitor) bool {
	return other.Native() == recv.Native()
}

func (recv *NativeVolumeMonitor) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *NativeVolumeMonitor) FieldParentInstance() *VolumeMonitor {
	argValue := gi.ObjectFieldGet(nativeVolumeMonitorObject, recv.Native(), "parent_instance")
	value := VolumeMonitorNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *NativeVolumeMonitor) SetFieldParentInstance(value *VolumeMonitor) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(nativeVolumeMonitorObject, recv.Native(), "parent_instance", argValue)
}

var networkAddressObject *gi.Object
var networkAddressObject_Once sync.Once

func networkAddressObject_Set() error {
	var err error
	networkAddressObject_Once.Do(func() {
		networkAddressObject, err = gi.ObjectNew("Gio", "NetworkAddress")
	})
	return err
}

type NetworkAddress struct {
	native unsafe.Pointer
}

func NetworkAddressNewFromNative(native unsafe.Pointer) *NetworkAddress {
	instance := &NetworkAddress{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *NetworkAddress) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToNetworkAddress down casts any arbitrary Object to NetworkAddress.
Exercise care, as this is a potentially dangerous function
if the Object is not a NetworkAddress.
*/
func CastToNetworkAddress(object *gobject.Object) *NetworkAddress {
	return NetworkAddressNewFromNative(object.Native())
}

// Equals compares this NetworkAddress with another NetworkAddress, and returns true if they represent the same GObject.
func (recv *NetworkAddress) Equals(other *NetworkAddress) bool {
	return other.Native() == recv.Native()
}

func (recv *NetworkAddress) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *NetworkAddress) FieldParentInstance() *gobject.Object {
	argValue := gi.ObjectFieldGet(networkAddressObject, recv.Native(), "parent_instance")
	value := gobject.ObjectNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *NetworkAddress) SetFieldParentInstance(value *gobject.Object) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(networkAddressObject, recv.Native(), "parent_instance", argValue)
}

var networkAddressNewFunction *gi.Function
var networkAddressNewFunction_Once sync.Once

func networkAddressNewFunction_Set() error {
	var err error
	networkAddressNewFunction_Once.Do(func() {
		err = networkAddressObject_Set()
		if err != nil {
			return
		}
		networkAddressNewFunction, err = networkAddressObject.InvokerNew("new")
	})
	return err
}

// NetworkAddressNew is a representation of the C type g_network_address_new.
func NetworkAddressNew(hostname string, port uint16) *NetworkAddress {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(hostname)
	inArgs[1].SetUint16(port)

	var ret gi.Argument

	err := networkAddressNewFunction_Set()
	if err == nil {
		ret = networkAddressNewFunction.Invoke(inArgs[:], nil)
	}

	retGo := NetworkAddressNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

var networkAddressNewLoopbackFunction *gi.Function
var networkAddressNewLoopbackFunction_Once sync.Once

func networkAddressNewLoopbackFunction_Set() error {
	var err error
	networkAddressNewLoopbackFunction_Once.Do(func() {
		err = networkAddressObject_Set()
		if err != nil {
			return
		}
		networkAddressNewLoopbackFunction, err = networkAddressObject.InvokerNew("new_loopback")
	})
	return err
}

// NetworkAddressNewLoopback is a representation of the C type g_network_address_new_loopback.
func NetworkAddressNewLoopback(port uint16) *NetworkAddress {
	var inArgs [1]gi.Argument
	inArgs[0].SetUint16(port)

	var ret gi.Argument

	err := networkAddressNewLoopbackFunction_Set()
	if err == nil {
		ret = networkAddressNewLoopbackFunction.Invoke(inArgs[:], nil)
	}

	retGo := NetworkAddressNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

var networkAddressGetHostnameFunction *gi.Function
var networkAddressGetHostnameFunction_Once sync.Once

func networkAddressGetHostnameFunction_Set() error {
	var err error
	networkAddressGetHostnameFunction_Once.Do(func() {
		err = networkAddressObject_Set()
		if err != nil {
			return
		}
		networkAddressGetHostnameFunction, err = networkAddressObject.InvokerNew("get_hostname")
	})
	return err
}

// GetHostname is a representation of the C type g_network_address_get_hostname.
func (recv *NetworkAddress) GetHostname() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := networkAddressGetHostnameFunction_Set()
	if err == nil {
		ret = networkAddressGetHostnameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var networkAddressGetPortFunction *gi.Function
var networkAddressGetPortFunction_Once sync.Once

func networkAddressGetPortFunction_Set() error {
	var err error
	networkAddressGetPortFunction_Once.Do(func() {
		err = networkAddressObject_Set()
		if err != nil {
			return
		}
		networkAddressGetPortFunction, err = networkAddressObject.InvokerNew("get_port")
	})
	return err
}

// GetPort is a representation of the C type g_network_address_get_port.
func (recv *NetworkAddress) GetPort() uint16 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := networkAddressGetPortFunction_Set()
	if err == nil {
		ret = networkAddressGetPortFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint16()

	return retGo
}

var networkAddressGetSchemeFunction *gi.Function
var networkAddressGetSchemeFunction_Once sync.Once

func networkAddressGetSchemeFunction_Set() error {
	var err error
	networkAddressGetSchemeFunction_Once.Do(func() {
		err = networkAddressObject_Set()
		if err != nil {
			return
		}
		networkAddressGetSchemeFunction, err = networkAddressObject.InvokerNew("get_scheme")
	})
	return err
}

// GetScheme is a representation of the C type g_network_address_get_scheme.
func (recv *NetworkAddress) GetScheme() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := networkAddressGetSchemeFunction_Set()
	if err == nil {
		ret = networkAddressGetSchemeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var networkServiceObject *gi.Object
var networkServiceObject_Once sync.Once

func networkServiceObject_Set() error {
	var err error
	networkServiceObject_Once.Do(func() {
		networkServiceObject, err = gi.ObjectNew("Gio", "NetworkService")
	})
	return err
}

type NetworkService struct {
	native unsafe.Pointer
}

func NetworkServiceNewFromNative(native unsafe.Pointer) *NetworkService {
	instance := &NetworkService{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *NetworkService) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToNetworkService down casts any arbitrary Object to NetworkService.
Exercise care, as this is a potentially dangerous function
if the Object is not a NetworkService.
*/
func CastToNetworkService(object *gobject.Object) *NetworkService {
	return NetworkServiceNewFromNative(object.Native())
}

// Equals compares this NetworkService with another NetworkService, and returns true if they represent the same GObject.
func (recv *NetworkService) Equals(other *NetworkService) bool {
	return other.Native() == recv.Native()
}

func (recv *NetworkService) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *NetworkService) FieldParentInstance() *gobject.Object {
	argValue := gi.ObjectFieldGet(networkServiceObject, recv.Native(), "parent_instance")
	value := gobject.ObjectNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *NetworkService) SetFieldParentInstance(value *gobject.Object) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(networkServiceObject, recv.Native(), "parent_instance", argValue)
}

var networkServiceNewFunction *gi.Function
var networkServiceNewFunction_Once sync.Once

func networkServiceNewFunction_Set() error {
	var err error
	networkServiceNewFunction_Once.Do(func() {
		err = networkServiceObject_Set()
		if err != nil {
			return
		}
		networkServiceNewFunction, err = networkServiceObject.InvokerNew("new")
	})
	return err
}

// NetworkServiceNew is a representation of the C type g_network_service_new.
func NetworkServiceNew(service string, protocol string, domain string) *NetworkService {
	var inArgs [3]gi.Argument
	inArgs[0].SetString(service)
	inArgs[1].SetString(protocol)
	inArgs[2].SetString(domain)

	var ret gi.Argument

	err := networkServiceNewFunction_Set()
	if err == nil {
		ret = networkServiceNewFunction.Invoke(inArgs[:], nil)
	}

	retGo := NetworkServiceNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

var networkServiceGetDomainFunction *gi.Function
var networkServiceGetDomainFunction_Once sync.Once

func networkServiceGetDomainFunction_Set() error {
	var err error
	networkServiceGetDomainFunction_Once.Do(func() {
		err = networkServiceObject_Set()
		if err != nil {
			return
		}
		networkServiceGetDomainFunction, err = networkServiceObject.InvokerNew("get_domain")
	})
	return err
}

// GetDomain is a representation of the C type g_network_service_get_domain.
func (recv *NetworkService) GetDomain() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := networkServiceGetDomainFunction_Set()
	if err == nil {
		ret = networkServiceGetDomainFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var networkServiceGetProtocolFunction *gi.Function
var networkServiceGetProtocolFunction_Once sync.Once

func networkServiceGetProtocolFunction_Set() error {
	var err error
	networkServiceGetProtocolFunction_Once.Do(func() {
		err = networkServiceObject_Set()
		if err != nil {
			return
		}
		networkServiceGetProtocolFunction, err = networkServiceObject.InvokerNew("get_protocol")
	})
	return err
}

// GetProtocol is a representation of the C type g_network_service_get_protocol.
func (recv *NetworkService) GetProtocol() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := networkServiceGetProtocolFunction_Set()
	if err == nil {
		ret = networkServiceGetProtocolFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var networkServiceGetSchemeFunction *gi.Function
var networkServiceGetSchemeFunction_Once sync.Once

func networkServiceGetSchemeFunction_Set() error {
	var err error
	networkServiceGetSchemeFunction_Once.Do(func() {
		err = networkServiceObject_Set()
		if err != nil {
			return
		}
		networkServiceGetSchemeFunction, err = networkServiceObject.InvokerNew("get_scheme")
	})
	return err
}

// GetScheme is a representation of the C type g_network_service_get_scheme.
func (recv *NetworkService) GetScheme() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := networkServiceGetSchemeFunction_Set()
	if err == nil {
		ret = networkServiceGetSchemeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var networkServiceGetServiceFunction *gi.Function
var networkServiceGetServiceFunction_Once sync.Once

func networkServiceGetServiceFunction_Set() error {
	var err error
	networkServiceGetServiceFunction_Once.Do(func() {
		err = networkServiceObject_Set()
		if err != nil {
			return
		}
		networkServiceGetServiceFunction, err = networkServiceObject.InvokerNew("get_service")
	})
	return err
}

// GetService is a representation of the C type g_network_service_get_service.
func (recv *NetworkService) GetService() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := networkServiceGetServiceFunction_Set()
	if err == nil {
		ret = networkServiceGetServiceFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var networkServiceSetSchemeFunction *gi.Function
var networkServiceSetSchemeFunction_Once sync.Once

func networkServiceSetSchemeFunction_Set() error {
	var err error
	networkServiceSetSchemeFunction_Once.Do(func() {
		err = networkServiceObject_Set()
		if err != nil {
			return
		}
		networkServiceSetSchemeFunction, err = networkServiceObject.InvokerNew("set_scheme")
	})
	return err
}

// SetScheme is a representation of the C type g_network_service_set_scheme.
func (recv *NetworkService) SetScheme(scheme string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(scheme)

	err := networkServiceSetSchemeFunction_Set()
	if err == nil {
		networkServiceSetSchemeFunction.Invoke(inArgs[:], nil)
	}

	return
}

var notificationObject *gi.Object
var notificationObject_Once sync.Once

func notificationObject_Set() error {
	var err error
	notificationObject_Once.Do(func() {
		notificationObject, err = gi.ObjectNew("Gio", "Notification")
	})
	return err
}

type Notification struct {
	native unsafe.Pointer
}

func NotificationNewFromNative(native unsafe.Pointer) *Notification {
	instance := &Notification{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *Notification) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToNotification down casts any arbitrary Object to Notification.
Exercise care, as this is a potentially dangerous function
if the Object is not a Notification.
*/
func CastToNotification(object *gobject.Object) *Notification {
	return NotificationNewFromNative(object.Native())
}

// Equals compares this Notification with another Notification, and returns true if they represent the same GObject.
func (recv *Notification) Equals(other *Notification) bool {
	return other.Native() == recv.Native()
}

func (recv *Notification) Native() unsafe.Pointer {
	return recv.native
}

var notificationNewFunction *gi.Function
var notificationNewFunction_Once sync.Once

func notificationNewFunction_Set() error {
	var err error
	notificationNewFunction_Once.Do(func() {
		err = notificationObject_Set()
		if err != nil {
			return
		}
		notificationNewFunction, err = notificationObject.InvokerNew("new")
	})
	return err
}

// NotificationNew is a representation of the C type g_notification_new.
func NotificationNew(title string) *Notification {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(title)

	var ret gi.Argument

	err := notificationNewFunction_Set()
	if err == nil {
		ret = notificationNewFunction.Invoke(inArgs[:], nil)
	}

	retGo := NotificationNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

var notificationAddButtonFunction *gi.Function
var notificationAddButtonFunction_Once sync.Once

func notificationAddButtonFunction_Set() error {
	var err error
	notificationAddButtonFunction_Once.Do(func() {
		err = notificationObject_Set()
		if err != nil {
			return
		}
		notificationAddButtonFunction, err = notificationObject.InvokerNew("add_button")
	})
	return err
}

// AddButton is a representation of the C type g_notification_add_button.
func (recv *Notification) AddButton(label string, detailedAction string) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(label)
	inArgs[2].SetString(detailedAction)

	err := notificationAddButtonFunction_Set()
	if err == nil {
		notificationAddButtonFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'g_notification_add_button_with_target' : parameter '...' of type 'nil' not supported

// UNSUPPORTED : C value 'g_notification_add_button_with_target_value' : parameter 'target' of type 'GLib.Variant' not supported

var notificationSetBodyFunction *gi.Function
var notificationSetBodyFunction_Once sync.Once

func notificationSetBodyFunction_Set() error {
	var err error
	notificationSetBodyFunction_Once.Do(func() {
		err = notificationObject_Set()
		if err != nil {
			return
		}
		notificationSetBodyFunction, err = notificationObject.InvokerNew("set_body")
	})
	return err
}

// SetBody is a representation of the C type g_notification_set_body.
func (recv *Notification) SetBody(body string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(body)

	err := notificationSetBodyFunction_Set()
	if err == nil {
		notificationSetBodyFunction.Invoke(inArgs[:], nil)
	}

	return
}

var notificationSetDefaultActionFunction *gi.Function
var notificationSetDefaultActionFunction_Once sync.Once

func notificationSetDefaultActionFunction_Set() error {
	var err error
	notificationSetDefaultActionFunction_Once.Do(func() {
		err = notificationObject_Set()
		if err != nil {
			return
		}
		notificationSetDefaultActionFunction, err = notificationObject.InvokerNew("set_default_action")
	})
	return err
}

// SetDefaultAction is a representation of the C type g_notification_set_default_action.
func (recv *Notification) SetDefaultAction(detailedAction string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(detailedAction)

	err := notificationSetDefaultActionFunction_Set()
	if err == nil {
		notificationSetDefaultActionFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'g_notification_set_default_action_and_target' : parameter '...' of type 'nil' not supported

// UNSUPPORTED : C value 'g_notification_set_default_action_and_target_value' : parameter 'target' of type 'GLib.Variant' not supported

// UNSUPPORTED : C value 'g_notification_set_icon' : parameter 'icon' of type 'Icon' not supported

var notificationSetPriorityFunction *gi.Function
var notificationSetPriorityFunction_Once sync.Once

func notificationSetPriorityFunction_Set() error {
	var err error
	notificationSetPriorityFunction_Once.Do(func() {
		err = notificationObject_Set()
		if err != nil {
			return
		}
		notificationSetPriorityFunction, err = notificationObject.InvokerNew("set_priority")
	})
	return err
}

// SetPriority is a representation of the C type g_notification_set_priority.
func (recv *Notification) SetPriority(priority NotificationPriority) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(int32(priority))

	err := notificationSetPriorityFunction_Set()
	if err == nil {
		notificationSetPriorityFunction.Invoke(inArgs[:], nil)
	}

	return
}

var notificationSetTitleFunction *gi.Function
var notificationSetTitleFunction_Once sync.Once

func notificationSetTitleFunction_Set() error {
	var err error
	notificationSetTitleFunction_Once.Do(func() {
		err = notificationObject_Set()
		if err != nil {
			return
		}
		notificationSetTitleFunction, err = notificationObject.InvokerNew("set_title")
	})
	return err
}

// SetTitle is a representation of the C type g_notification_set_title.
func (recv *Notification) SetTitle(title string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(title)

	err := notificationSetTitleFunction_Set()
	if err == nil {
		notificationSetTitleFunction.Invoke(inArgs[:], nil)
	}

	return
}

var notificationSetUrgentFunction *gi.Function
var notificationSetUrgentFunction_Once sync.Once

func notificationSetUrgentFunction_Set() error {
	var err error
	notificationSetUrgentFunction_Once.Do(func() {
		err = notificationObject_Set()
		if err != nil {
			return
		}
		notificationSetUrgentFunction, err = notificationObject.InvokerNew("set_urgent")
	})
	return err
}

// SetUrgent is a representation of the C type g_notification_set_urgent.
func (recv *Notification) SetUrgent(urgent bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetBoolean(urgent)

	err := notificationSetUrgentFunction_Set()
	if err == nil {
		notificationSetUrgentFunction.Invoke(inArgs[:], nil)
	}

	return
}

var outputStreamObject *gi.Object
var outputStreamObject_Once sync.Once

func outputStreamObject_Set() error {
	var err error
	outputStreamObject_Once.Do(func() {
		outputStreamObject, err = gi.ObjectNew("Gio", "OutputStream")
	})
	return err
}

type OutputStream struct {
	native unsafe.Pointer
}

func OutputStreamNewFromNative(native unsafe.Pointer) *OutputStream {
	instance := &OutputStream{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *OutputStream) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToOutputStream down casts any arbitrary Object to OutputStream.
Exercise care, as this is a potentially dangerous function
if the Object is not a OutputStream.
*/
func CastToOutputStream(object *gobject.Object) *OutputStream {
	return OutputStreamNewFromNative(object.Native())
}

// Equals compares this OutputStream with another OutputStream, and returns true if they represent the same GObject.
func (recv *OutputStream) Equals(other *OutputStream) bool {
	return other.Native() == recv.Native()
}

func (recv *OutputStream) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *OutputStream) FieldParentInstance() *gobject.Object {
	argValue := gi.ObjectFieldGet(outputStreamObject, recv.Native(), "parent_instance")
	value := gobject.ObjectNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *OutputStream) SetFieldParentInstance(value *gobject.Object) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(outputStreamObject, recv.Native(), "parent_instance", argValue)
}

var outputStreamClearPendingFunction *gi.Function
var outputStreamClearPendingFunction_Once sync.Once

func outputStreamClearPendingFunction_Set() error {
	var err error
	outputStreamClearPendingFunction_Once.Do(func() {
		err = outputStreamObject_Set()
		if err != nil {
			return
		}
		outputStreamClearPendingFunction, err = outputStreamObject.InvokerNew("clear_pending")
	})
	return err
}

// ClearPending is a representation of the C type g_output_stream_clear_pending.
func (recv *OutputStream) ClearPending() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := outputStreamClearPendingFunction_Set()
	if err == nil {
		outputStreamClearPendingFunction.Invoke(inArgs[:], nil)
	}

	return
}

var outputStreamCloseFunction *gi.Function
var outputStreamCloseFunction_Once sync.Once

func outputStreamCloseFunction_Set() error {
	var err error
	outputStreamCloseFunction_Once.Do(func() {
		err = outputStreamObject_Set()
		if err != nil {
			return
		}
		outputStreamCloseFunction, err = outputStreamObject.InvokerNew("close")
	})
	return err
}

// Close is a representation of the C type g_output_stream_close.
func (recv *OutputStream) Close(cancellable *Cancellable) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(cancellable.Native())

	var ret gi.Argument

	err := outputStreamCloseFunction_Set()
	if err == nil {
		ret = outputStreamCloseFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'g_output_stream_close_async' : parameter 'callback' of type 'AsyncReadyCallback' not supported

// UNSUPPORTED : C value 'g_output_stream_close_finish' : parameter 'result' of type 'AsyncResult' not supported

var outputStreamFlushFunction *gi.Function
var outputStreamFlushFunction_Once sync.Once

func outputStreamFlushFunction_Set() error {
	var err error
	outputStreamFlushFunction_Once.Do(func() {
		err = outputStreamObject_Set()
		if err != nil {
			return
		}
		outputStreamFlushFunction, err = outputStreamObject.InvokerNew("flush")
	})
	return err
}

// Flush is a representation of the C type g_output_stream_flush.
func (recv *OutputStream) Flush(cancellable *Cancellable) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(cancellable.Native())

	var ret gi.Argument

	err := outputStreamFlushFunction_Set()
	if err == nil {
		ret = outputStreamFlushFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'g_output_stream_flush_async' : parameter 'callback' of type 'AsyncReadyCallback' not supported

// UNSUPPORTED : C value 'g_output_stream_flush_finish' : parameter 'result' of type 'AsyncResult' not supported

var outputStreamHasPendingFunction *gi.Function
var outputStreamHasPendingFunction_Once sync.Once

func outputStreamHasPendingFunction_Set() error {
	var err error
	outputStreamHasPendingFunction_Once.Do(func() {
		err = outputStreamObject_Set()
		if err != nil {
			return
		}
		outputStreamHasPendingFunction, err = outputStreamObject.InvokerNew("has_pending")
	})
	return err
}

// HasPending is a representation of the C type g_output_stream_has_pending.
func (recv *OutputStream) HasPending() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := outputStreamHasPendingFunction_Set()
	if err == nil {
		ret = outputStreamHasPendingFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var outputStreamIsClosedFunction *gi.Function
var outputStreamIsClosedFunction_Once sync.Once

func outputStreamIsClosedFunction_Set() error {
	var err error
	outputStreamIsClosedFunction_Once.Do(func() {
		err = outputStreamObject_Set()
		if err != nil {
			return
		}
		outputStreamIsClosedFunction, err = outputStreamObject.InvokerNew("is_closed")
	})
	return err
}

// IsClosed is a representation of the C type g_output_stream_is_closed.
func (recv *OutputStream) IsClosed() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := outputStreamIsClosedFunction_Set()
	if err == nil {
		ret = outputStreamIsClosedFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var outputStreamIsClosingFunction *gi.Function
var outputStreamIsClosingFunction_Once sync.Once

func outputStreamIsClosingFunction_Set() error {
	var err error
	outputStreamIsClosingFunction_Once.Do(func() {
		err = outputStreamObject_Set()
		if err != nil {
			return
		}
		outputStreamIsClosingFunction, err = outputStreamObject.InvokerNew("is_closing")
	})
	return err
}

// IsClosing is a representation of the C type g_output_stream_is_closing.
func (recv *OutputStream) IsClosing() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := outputStreamIsClosingFunction_Set()
	if err == nil {
		ret = outputStreamIsClosingFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'g_output_stream_printf' : parameter 'error' of type 'GLib.Error' not supported

var outputStreamSetPendingFunction *gi.Function
var outputStreamSetPendingFunction_Once sync.Once

func outputStreamSetPendingFunction_Set() error {
	var err error
	outputStreamSetPendingFunction_Once.Do(func() {
		err = outputStreamObject_Set()
		if err != nil {
			return
		}
		outputStreamSetPendingFunction, err = outputStreamObject.InvokerNew("set_pending")
	})
	return err
}

// SetPending is a representation of the C type g_output_stream_set_pending.
func (recv *OutputStream) SetPending() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := outputStreamSetPendingFunction_Set()
	if err == nil {
		ret = outputStreamSetPendingFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'g_output_stream_splice' : parameter 'flags' of type 'OutputStreamSpliceFlags' not supported

// UNSUPPORTED : C value 'g_output_stream_splice_async' : parameter 'flags' of type 'OutputStreamSpliceFlags' not supported

// UNSUPPORTED : C value 'g_output_stream_splice_finish' : parameter 'result' of type 'AsyncResult' not supported

// UNSUPPORTED : C value 'g_output_stream_vprintf' : parameter 'error' of type 'GLib.Error' not supported

// UNSUPPORTED : C value 'g_output_stream_write' : parameter 'buffer' of type 'nil' not supported

// UNSUPPORTED : C value 'g_output_stream_write_all' : parameter 'buffer' of type 'nil' not supported

// UNSUPPORTED : C value 'g_output_stream_write_all_async' : parameter 'buffer' of type 'nil' not supported

// UNSUPPORTED : C value 'g_output_stream_write_all_finish' : parameter 'result' of type 'AsyncResult' not supported

// UNSUPPORTED : C value 'g_output_stream_write_async' : parameter 'buffer' of type 'nil' not supported

// UNSUPPORTED : C value 'g_output_stream_write_bytes' : parameter 'bytes' of type 'GLib.Bytes' not supported

// UNSUPPORTED : C value 'g_output_stream_write_bytes_async' : parameter 'bytes' of type 'GLib.Bytes' not supported

// UNSUPPORTED : C value 'g_output_stream_write_bytes_finish' : parameter 'result' of type 'AsyncResult' not supported

// UNSUPPORTED : C value 'g_output_stream_write_finish' : parameter 'result' of type 'AsyncResult' not supported

// UNSUPPORTED : C value 'g_output_stream_writev' : parameter 'vectors' of type 'nil' not supported

// UNSUPPORTED : C value 'g_output_stream_writev_all' : parameter 'vectors' of type 'nil' not supported

// UNSUPPORTED : C value 'g_output_stream_writev_all_async' : parameter 'vectors' of type 'nil' not supported

// UNSUPPORTED : C value 'g_output_stream_writev_all_finish' : parameter 'result' of type 'AsyncResult' not supported

// UNSUPPORTED : C value 'g_output_stream_writev_async' : parameter 'vectors' of type 'nil' not supported

// UNSUPPORTED : C value 'g_output_stream_writev_finish' : parameter 'result' of type 'AsyncResult' not supported

var permissionObject *gi.Object
var permissionObject_Once sync.Once

func permissionObject_Set() error {
	var err error
	permissionObject_Once.Do(func() {
		permissionObject, err = gi.ObjectNew("Gio", "Permission")
	})
	return err
}

type Permission struct {
	native unsafe.Pointer
}

func PermissionNewFromNative(native unsafe.Pointer) *Permission {
	instance := &Permission{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *Permission) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToPermission down casts any arbitrary Object to Permission.
Exercise care, as this is a potentially dangerous function
if the Object is not a Permission.
*/
func CastToPermission(object *gobject.Object) *Permission {
	return PermissionNewFromNative(object.Native())
}

// Equals compares this Permission with another Permission, and returns true if they represent the same GObject.
func (recv *Permission) Equals(other *Permission) bool {
	return other.Native() == recv.Native()
}

func (recv *Permission) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *Permission) FieldParentInstance() *gobject.Object {
	argValue := gi.ObjectFieldGet(permissionObject, recv.Native(), "parent_instance")
	value := gobject.ObjectNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *Permission) SetFieldParentInstance(value *gobject.Object) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(permissionObject, recv.Native(), "parent_instance", argValue)
}

var permissionAcquireFunction *gi.Function
var permissionAcquireFunction_Once sync.Once

func permissionAcquireFunction_Set() error {
	var err error
	permissionAcquireFunction_Once.Do(func() {
		err = permissionObject_Set()
		if err != nil {
			return
		}
		permissionAcquireFunction, err = permissionObject.InvokerNew("acquire")
	})
	return err
}

// Acquire is a representation of the C type g_permission_acquire.
func (recv *Permission) Acquire(cancellable *Cancellable) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(cancellable.Native())

	var ret gi.Argument

	err := permissionAcquireFunction_Set()
	if err == nil {
		ret = permissionAcquireFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'g_permission_acquire_async' : parameter 'callback' of type 'AsyncReadyCallback' not supported

// UNSUPPORTED : C value 'g_permission_acquire_finish' : parameter 'result' of type 'AsyncResult' not supported

var permissionGetAllowedFunction *gi.Function
var permissionGetAllowedFunction_Once sync.Once

func permissionGetAllowedFunction_Set() error {
	var err error
	permissionGetAllowedFunction_Once.Do(func() {
		err = permissionObject_Set()
		if err != nil {
			return
		}
		permissionGetAllowedFunction, err = permissionObject.InvokerNew("get_allowed")
	})
	return err
}

// GetAllowed is a representation of the C type g_permission_get_allowed.
func (recv *Permission) GetAllowed() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := permissionGetAllowedFunction_Set()
	if err == nil {
		ret = permissionGetAllowedFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var permissionGetCanAcquireFunction *gi.Function
var permissionGetCanAcquireFunction_Once sync.Once

func permissionGetCanAcquireFunction_Set() error {
	var err error
	permissionGetCanAcquireFunction_Once.Do(func() {
		err = permissionObject_Set()
		if err != nil {
			return
		}
		permissionGetCanAcquireFunction, err = permissionObject.InvokerNew("get_can_acquire")
	})
	return err
}

// GetCanAcquire is a representation of the C type g_permission_get_can_acquire.
func (recv *Permission) GetCanAcquire() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := permissionGetCanAcquireFunction_Set()
	if err == nil {
		ret = permissionGetCanAcquireFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var permissionGetCanReleaseFunction *gi.Function
var permissionGetCanReleaseFunction_Once sync.Once

func permissionGetCanReleaseFunction_Set() error {
	var err error
	permissionGetCanReleaseFunction_Once.Do(func() {
		err = permissionObject_Set()
		if err != nil {
			return
		}
		permissionGetCanReleaseFunction, err = permissionObject.InvokerNew("get_can_release")
	})
	return err
}

// GetCanRelease is a representation of the C type g_permission_get_can_release.
func (recv *Permission) GetCanRelease() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := permissionGetCanReleaseFunction_Set()
	if err == nil {
		ret = permissionGetCanReleaseFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var permissionImplUpdateFunction *gi.Function
var permissionImplUpdateFunction_Once sync.Once

func permissionImplUpdateFunction_Set() error {
	var err error
	permissionImplUpdateFunction_Once.Do(func() {
		err = permissionObject_Set()
		if err != nil {
			return
		}
		permissionImplUpdateFunction, err = permissionObject.InvokerNew("impl_update")
	})
	return err
}

// ImplUpdate is a representation of the C type g_permission_impl_update.
func (recv *Permission) ImplUpdate(allowed bool, canAcquire bool, canRelease bool) {
	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetBoolean(allowed)
	inArgs[2].SetBoolean(canAcquire)
	inArgs[3].SetBoolean(canRelease)

	err := permissionImplUpdateFunction_Set()
	if err == nil {
		permissionImplUpdateFunction.Invoke(inArgs[:], nil)
	}

	return
}

var permissionReleaseFunction *gi.Function
var permissionReleaseFunction_Once sync.Once

func permissionReleaseFunction_Set() error {
	var err error
	permissionReleaseFunction_Once.Do(func() {
		err = permissionObject_Set()
		if err != nil {
			return
		}
		permissionReleaseFunction, err = permissionObject.InvokerNew("release")
	})
	return err
}

// Release is a representation of the C type g_permission_release.
func (recv *Permission) Release(cancellable *Cancellable) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(cancellable.Native())

	var ret gi.Argument

	err := permissionReleaseFunction_Set()
	if err == nil {
		ret = permissionReleaseFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'g_permission_release_async' : parameter 'callback' of type 'AsyncReadyCallback' not supported

// UNSUPPORTED : C value 'g_permission_release_finish' : parameter 'result' of type 'AsyncResult' not supported

var propertyActionObject *gi.Object
var propertyActionObject_Once sync.Once

func propertyActionObject_Set() error {
	var err error
	propertyActionObject_Once.Do(func() {
		propertyActionObject, err = gi.ObjectNew("Gio", "PropertyAction")
	})
	return err
}

type PropertyAction struct {
	native unsafe.Pointer
}

func PropertyActionNewFromNative(native unsafe.Pointer) *PropertyAction {
	instance := &PropertyAction{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *PropertyAction) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToPropertyAction down casts any arbitrary Object to PropertyAction.
Exercise care, as this is a potentially dangerous function
if the Object is not a PropertyAction.
*/
func CastToPropertyAction(object *gobject.Object) *PropertyAction {
	return PropertyActionNewFromNative(object.Native())
}

// Equals compares this PropertyAction with another PropertyAction, and returns true if they represent the same GObject.
func (recv *PropertyAction) Equals(other *PropertyAction) bool {
	return other.Native() == recv.Native()
}

func (recv *PropertyAction) Native() unsafe.Pointer {
	return recv.native
}

var propertyActionNewFunction *gi.Function
var propertyActionNewFunction_Once sync.Once

func propertyActionNewFunction_Set() error {
	var err error
	propertyActionNewFunction_Once.Do(func() {
		err = propertyActionObject_Set()
		if err != nil {
			return
		}
		propertyActionNewFunction, err = propertyActionObject.InvokerNew("new")
	})
	return err
}

// PropertyActionNew is a representation of the C type g_property_action_new.
func PropertyActionNew(name string, object_ *gobject.Object, propertyName string) *PropertyAction {
	var inArgs [3]gi.Argument
	inArgs[0].SetString(name)
	inArgs[1].SetPointer(object_.Native())
	inArgs[2].SetString(propertyName)

	var ret gi.Argument

	err := propertyActionNewFunction_Set()
	if err == nil {
		ret = propertyActionNewFunction.Invoke(inArgs[:], nil)
	}

	retGo := PropertyActionNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

var proxyAddressObject *gi.Object
var proxyAddressObject_Once sync.Once

func proxyAddressObject_Set() error {
	var err error
	proxyAddressObject_Once.Do(func() {
		proxyAddressObject, err = gi.ObjectNew("Gio", "ProxyAddress")
	})
	return err
}

type ProxyAddress struct {
	native unsafe.Pointer
}

func ProxyAddressNewFromNative(native unsafe.Pointer) *ProxyAddress {
	instance := &ProxyAddress{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// InetSocketAddress upcasts to *InetSocketAddress
func (recv *ProxyAddress) InetSocketAddress() *InetSocketAddress {
	return InetSocketAddressNewFromNative(recv.Native())
}

// SocketAddress upcasts to *SocketAddress
func (recv *ProxyAddress) SocketAddress() *SocketAddress {
	return SocketAddressNewFromNative(recv.Native())
}

// Object upcasts to *Object
func (recv *ProxyAddress) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToProxyAddress down casts any arbitrary Object to ProxyAddress.
Exercise care, as this is a potentially dangerous function
if the Object is not a ProxyAddress.
*/
func CastToProxyAddress(object *gobject.Object) *ProxyAddress {
	return ProxyAddressNewFromNative(object.Native())
}

// Equals compares this ProxyAddress with another ProxyAddress, and returns true if they represent the same GObject.
func (recv *ProxyAddress) Equals(other *ProxyAddress) bool {
	return other.Native() == recv.Native()
}

func (recv *ProxyAddress) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *ProxyAddress) FieldParentInstance() *InetSocketAddress {
	argValue := gi.ObjectFieldGet(proxyAddressObject, recv.Native(), "parent_instance")
	value := InetSocketAddressNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *ProxyAddress) SetFieldParentInstance(value *InetSocketAddress) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(proxyAddressObject, recv.Native(), "parent_instance", argValue)
}

var proxyAddressNewFunction *gi.Function
var proxyAddressNewFunction_Once sync.Once

func proxyAddressNewFunction_Set() error {
	var err error
	proxyAddressNewFunction_Once.Do(func() {
		err = proxyAddressObject_Set()
		if err != nil {
			return
		}
		proxyAddressNewFunction, err = proxyAddressObject.InvokerNew("new")
	})
	return err
}

// ProxyAddressNew is a representation of the C type g_proxy_address_new.
func ProxyAddressNew(inetaddr *InetAddress, port uint16, protocol string, destHostname string, destPort uint16, username string, password string) *ProxyAddress {
	var inArgs [7]gi.Argument
	inArgs[0].SetPointer(inetaddr.Native())
	inArgs[1].SetUint16(port)
	inArgs[2].SetString(protocol)
	inArgs[3].SetString(destHostname)
	inArgs[4].SetUint16(destPort)
	inArgs[5].SetString(username)
	inArgs[6].SetString(password)

	var ret gi.Argument

	err := proxyAddressNewFunction_Set()
	if err == nil {
		ret = proxyAddressNewFunction.Invoke(inArgs[:], nil)
	}

	retGo := ProxyAddressNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

var proxyAddressGetDestinationHostnameFunction *gi.Function
var proxyAddressGetDestinationHostnameFunction_Once sync.Once

func proxyAddressGetDestinationHostnameFunction_Set() error {
	var err error
	proxyAddressGetDestinationHostnameFunction_Once.Do(func() {
		err = proxyAddressObject_Set()
		if err != nil {
			return
		}
		proxyAddressGetDestinationHostnameFunction, err = proxyAddressObject.InvokerNew("get_destination_hostname")
	})
	return err
}

// GetDestinationHostname is a representation of the C type g_proxy_address_get_destination_hostname.
func (recv *ProxyAddress) GetDestinationHostname() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := proxyAddressGetDestinationHostnameFunction_Set()
	if err == nil {
		ret = proxyAddressGetDestinationHostnameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var proxyAddressGetDestinationPortFunction *gi.Function
var proxyAddressGetDestinationPortFunction_Once sync.Once

func proxyAddressGetDestinationPortFunction_Set() error {
	var err error
	proxyAddressGetDestinationPortFunction_Once.Do(func() {
		err = proxyAddressObject_Set()
		if err != nil {
			return
		}
		proxyAddressGetDestinationPortFunction, err = proxyAddressObject.InvokerNew("get_destination_port")
	})
	return err
}

// GetDestinationPort is a representation of the C type g_proxy_address_get_destination_port.
func (recv *ProxyAddress) GetDestinationPort() uint16 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := proxyAddressGetDestinationPortFunction_Set()
	if err == nil {
		ret = proxyAddressGetDestinationPortFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint16()

	return retGo
}

var proxyAddressGetDestinationProtocolFunction *gi.Function
var proxyAddressGetDestinationProtocolFunction_Once sync.Once

func proxyAddressGetDestinationProtocolFunction_Set() error {
	var err error
	proxyAddressGetDestinationProtocolFunction_Once.Do(func() {
		err = proxyAddressObject_Set()
		if err != nil {
			return
		}
		proxyAddressGetDestinationProtocolFunction, err = proxyAddressObject.InvokerNew("get_destination_protocol")
	})
	return err
}

// GetDestinationProtocol is a representation of the C type g_proxy_address_get_destination_protocol.
func (recv *ProxyAddress) GetDestinationProtocol() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := proxyAddressGetDestinationProtocolFunction_Set()
	if err == nil {
		ret = proxyAddressGetDestinationProtocolFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var proxyAddressGetPasswordFunction *gi.Function
var proxyAddressGetPasswordFunction_Once sync.Once

func proxyAddressGetPasswordFunction_Set() error {
	var err error
	proxyAddressGetPasswordFunction_Once.Do(func() {
		err = proxyAddressObject_Set()
		if err != nil {
			return
		}
		proxyAddressGetPasswordFunction, err = proxyAddressObject.InvokerNew("get_password")
	})
	return err
}

// GetPassword is a representation of the C type g_proxy_address_get_password.
func (recv *ProxyAddress) GetPassword() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := proxyAddressGetPasswordFunction_Set()
	if err == nil {
		ret = proxyAddressGetPasswordFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var proxyAddressGetProtocolFunction *gi.Function
var proxyAddressGetProtocolFunction_Once sync.Once

func proxyAddressGetProtocolFunction_Set() error {
	var err error
	proxyAddressGetProtocolFunction_Once.Do(func() {
		err = proxyAddressObject_Set()
		if err != nil {
			return
		}
		proxyAddressGetProtocolFunction, err = proxyAddressObject.InvokerNew("get_protocol")
	})
	return err
}

// GetProtocol is a representation of the C type g_proxy_address_get_protocol.
func (recv *ProxyAddress) GetProtocol() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := proxyAddressGetProtocolFunction_Set()
	if err == nil {
		ret = proxyAddressGetProtocolFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var proxyAddressGetUriFunction *gi.Function
var proxyAddressGetUriFunction_Once sync.Once

func proxyAddressGetUriFunction_Set() error {
	var err error
	proxyAddressGetUriFunction_Once.Do(func() {
		err = proxyAddressObject_Set()
		if err != nil {
			return
		}
		proxyAddressGetUriFunction, err = proxyAddressObject.InvokerNew("get_uri")
	})
	return err
}

// GetUri is a representation of the C type g_proxy_address_get_uri.
func (recv *ProxyAddress) GetUri() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := proxyAddressGetUriFunction_Set()
	if err == nil {
		ret = proxyAddressGetUriFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var proxyAddressGetUsernameFunction *gi.Function
var proxyAddressGetUsernameFunction_Once sync.Once

func proxyAddressGetUsernameFunction_Set() error {
	var err error
	proxyAddressGetUsernameFunction_Once.Do(func() {
		err = proxyAddressObject_Set()
		if err != nil {
			return
		}
		proxyAddressGetUsernameFunction, err = proxyAddressObject.InvokerNew("get_username")
	})
	return err
}

// GetUsername is a representation of the C type g_proxy_address_get_username.
func (recv *ProxyAddress) GetUsername() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := proxyAddressGetUsernameFunction_Set()
	if err == nil {
		ret = proxyAddressGetUsernameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var proxyAddressEnumeratorObject *gi.Object
var proxyAddressEnumeratorObject_Once sync.Once

func proxyAddressEnumeratorObject_Set() error {
	var err error
	proxyAddressEnumeratorObject_Once.Do(func() {
		proxyAddressEnumeratorObject, err = gi.ObjectNew("Gio", "ProxyAddressEnumerator")
	})
	return err
}

type ProxyAddressEnumerator struct {
	native unsafe.Pointer
}

func ProxyAddressEnumeratorNewFromNative(native unsafe.Pointer) *ProxyAddressEnumerator {
	instance := &ProxyAddressEnumerator{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// SocketAddressEnumerator upcasts to *SocketAddressEnumerator
func (recv *ProxyAddressEnumerator) SocketAddressEnumerator() *SocketAddressEnumerator {
	return SocketAddressEnumeratorNewFromNative(recv.Native())
}

// Object upcasts to *Object
func (recv *ProxyAddressEnumerator) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToProxyAddressEnumerator down casts any arbitrary Object to ProxyAddressEnumerator.
Exercise care, as this is a potentially dangerous function
if the Object is not a ProxyAddressEnumerator.
*/
func CastToProxyAddressEnumerator(object *gobject.Object) *ProxyAddressEnumerator {
	return ProxyAddressEnumeratorNewFromNative(object.Native())
}

// Equals compares this ProxyAddressEnumerator with another ProxyAddressEnumerator, and returns true if they represent the same GObject.
func (recv *ProxyAddressEnumerator) Equals(other *ProxyAddressEnumerator) bool {
	return other.Native() == recv.Native()
}

func (recv *ProxyAddressEnumerator) Native() unsafe.Pointer {
	return recv.native
}

var resolverObject *gi.Object
var resolverObject_Once sync.Once

func resolverObject_Set() error {
	var err error
	resolverObject_Once.Do(func() {
		resolverObject, err = gi.ObjectNew("Gio", "Resolver")
	})
	return err
}

type Resolver struct {
	native unsafe.Pointer
}

func ResolverNewFromNative(native unsafe.Pointer) *Resolver {
	instance := &Resolver{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *Resolver) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToResolver down casts any arbitrary Object to Resolver.
Exercise care, as this is a potentially dangerous function
if the Object is not a Resolver.
*/
func CastToResolver(object *gobject.Object) *Resolver {
	return ResolverNewFromNative(object.Native())
}

// Equals compares this Resolver with another Resolver, and returns true if they represent the same GObject.
func (recv *Resolver) Equals(other *Resolver) bool {
	return other.Native() == recv.Native()
}

func (recv *Resolver) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *Resolver) FieldParentInstance() *gobject.Object {
	argValue := gi.ObjectFieldGet(resolverObject, recv.Native(), "parent_instance")
	value := gobject.ObjectNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *Resolver) SetFieldParentInstance(value *gobject.Object) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(resolverObject, recv.Native(), "parent_instance", argValue)
}

// FieldPriv returns the C field 'priv'.
func (recv *Resolver) FieldPriv() *ResolverPrivate {
	argValue := gi.ObjectFieldGet(resolverObject, recv.Native(), "priv")
	value := ResolverPrivateNewFromNative(argValue.Pointer())
	return value
}

// SetFieldPriv sets the value of the C field 'priv'.
func (recv *Resolver) SetFieldPriv(value *ResolverPrivate) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(resolverObject, recv.Native(), "priv", argValue)
}

var resolverLookupByAddressFunction *gi.Function
var resolverLookupByAddressFunction_Once sync.Once

func resolverLookupByAddressFunction_Set() error {
	var err error
	resolverLookupByAddressFunction_Once.Do(func() {
		err = resolverObject_Set()
		if err != nil {
			return
		}
		resolverLookupByAddressFunction, err = resolverObject.InvokerNew("lookup_by_address")
	})
	return err
}

// LookupByAddress is a representation of the C type g_resolver_lookup_by_address.
func (recv *Resolver) LookupByAddress(address *InetAddress, cancellable *Cancellable) string {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(address.Native())
	inArgs[2].SetPointer(cancellable.Native())

	var ret gi.Argument

	err := resolverLookupByAddressFunction_Set()
	if err == nil {
		ret = resolverLookupByAddressFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

// UNSUPPORTED : C value 'g_resolver_lookup_by_address_async' : parameter 'callback' of type 'AsyncReadyCallback' not supported

// UNSUPPORTED : C value 'g_resolver_lookup_by_address_finish' : parameter 'result' of type 'AsyncResult' not supported

// UNSUPPORTED : C value 'g_resolver_lookup_by_name' : return type 'GLib.List' not supported

// UNSUPPORTED : C value 'g_resolver_lookup_by_name_async' : parameter 'callback' of type 'AsyncReadyCallback' not supported

// UNSUPPORTED : C value 'g_resolver_lookup_by_name_finish' : parameter 'result' of type 'AsyncResult' not supported

// UNSUPPORTED : C value 'g_resolver_lookup_by_name_with_flags' : parameter 'flags' of type 'ResolverNameLookupFlags' not supported

// UNSUPPORTED : C value 'g_resolver_lookup_by_name_with_flags_async' : parameter 'flags' of type 'ResolverNameLookupFlags' not supported

// UNSUPPORTED : C value 'g_resolver_lookup_by_name_with_flags_finish' : parameter 'result' of type 'AsyncResult' not supported

// UNSUPPORTED : C value 'g_resolver_lookup_records' : return type 'GLib.List' not supported

// UNSUPPORTED : C value 'g_resolver_lookup_records_async' : parameter 'callback' of type 'AsyncReadyCallback' not supported

// UNSUPPORTED : C value 'g_resolver_lookup_records_finish' : parameter 'result' of type 'AsyncResult' not supported

// UNSUPPORTED : C value 'g_resolver_lookup_service' : return type 'GLib.List' not supported

// UNSUPPORTED : C value 'g_resolver_lookup_service_async' : parameter 'callback' of type 'AsyncReadyCallback' not supported

// UNSUPPORTED : C value 'g_resolver_lookup_service_finish' : parameter 'result' of type 'AsyncResult' not supported

var resolverSetDefaultFunction *gi.Function
var resolverSetDefaultFunction_Once sync.Once

func resolverSetDefaultFunction_Set() error {
	var err error
	resolverSetDefaultFunction_Once.Do(func() {
		err = resolverObject_Set()
		if err != nil {
			return
		}
		resolverSetDefaultFunction, err = resolverObject.InvokerNew("set_default")
	})
	return err
}

// SetDefault is a representation of the C type g_resolver_set_default.
func (recv *Resolver) SetDefault() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := resolverSetDefaultFunction_Set()
	if err == nil {
		resolverSetDefaultFunction.Invoke(inArgs[:], nil)
	}

	return
}

var settingsObject *gi.Object
var settingsObject_Once sync.Once

func settingsObject_Set() error {
	var err error
	settingsObject_Once.Do(func() {
		settingsObject, err = gi.ObjectNew("Gio", "Settings")
	})
	return err
}

type Settings struct {
	native unsafe.Pointer
}

func SettingsNewFromNative(native unsafe.Pointer) *Settings {
	instance := &Settings{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *Settings) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToSettings down casts any arbitrary Object to Settings.
Exercise care, as this is a potentially dangerous function
if the Object is not a Settings.
*/
func CastToSettings(object *gobject.Object) *Settings {
	return SettingsNewFromNative(object.Native())
}

// Equals compares this Settings with another Settings, and returns true if they represent the same GObject.
func (recv *Settings) Equals(other *Settings) bool {
	return other.Native() == recv.Native()
}

func (recv *Settings) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *Settings) FieldParentInstance() *gobject.Object {
	argValue := gi.ObjectFieldGet(settingsObject, recv.Native(), "parent_instance")
	value := gobject.ObjectNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *Settings) SetFieldParentInstance(value *gobject.Object) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(settingsObject, recv.Native(), "parent_instance", argValue)
}

// FieldPriv returns the C field 'priv'.
func (recv *Settings) FieldPriv() *SettingsPrivate {
	argValue := gi.ObjectFieldGet(settingsObject, recv.Native(), "priv")
	value := SettingsPrivateNewFromNative(argValue.Pointer())
	return value
}

// SetFieldPriv sets the value of the C field 'priv'.
func (recv *Settings) SetFieldPriv(value *SettingsPrivate) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(settingsObject, recv.Native(), "priv", argValue)
}

var settingsNewFunction *gi.Function
var settingsNewFunction_Once sync.Once

func settingsNewFunction_Set() error {
	var err error
	settingsNewFunction_Once.Do(func() {
		err = settingsObject_Set()
		if err != nil {
			return
		}
		settingsNewFunction, err = settingsObject.InvokerNew("new")
	})
	return err
}

// SettingsNew is a representation of the C type g_settings_new.
func SettingsNew(schemaId string) *Settings {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(schemaId)

	var ret gi.Argument

	err := settingsNewFunction_Set()
	if err == nil {
		ret = settingsNewFunction.Invoke(inArgs[:], nil)
	}

	retGo := SettingsNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

var settingsNewFullFunction *gi.Function
var settingsNewFullFunction_Once sync.Once

func settingsNewFullFunction_Set() error {
	var err error
	settingsNewFullFunction_Once.Do(func() {
		err = settingsObject_Set()
		if err != nil {
			return
		}
		settingsNewFullFunction, err = settingsObject.InvokerNew("new_full")
	})
	return err
}

// SettingsNewFull is a representation of the C type g_settings_new_full.
func SettingsNewFull(schema *SettingsSchema, backend *SettingsBackend, path string) *Settings {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(schema.Native())
	inArgs[1].SetPointer(backend.Native())
	inArgs[2].SetString(path)

	var ret gi.Argument

	err := settingsNewFullFunction_Set()
	if err == nil {
		ret = settingsNewFullFunction.Invoke(inArgs[:], nil)
	}

	retGo := SettingsNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

var settingsNewWithBackendFunction *gi.Function
var settingsNewWithBackendFunction_Once sync.Once

func settingsNewWithBackendFunction_Set() error {
	var err error
	settingsNewWithBackendFunction_Once.Do(func() {
		err = settingsObject_Set()
		if err != nil {
			return
		}
		settingsNewWithBackendFunction, err = settingsObject.InvokerNew("new_with_backend")
	})
	return err
}

// SettingsNewWithBackend is a representation of the C type g_settings_new_with_backend.
func SettingsNewWithBackend(schemaId string, backend *SettingsBackend) *Settings {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(schemaId)
	inArgs[1].SetPointer(backend.Native())

	var ret gi.Argument

	err := settingsNewWithBackendFunction_Set()
	if err == nil {
		ret = settingsNewWithBackendFunction.Invoke(inArgs[:], nil)
	}

	retGo := SettingsNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

var settingsNewWithBackendAndPathFunction *gi.Function
var settingsNewWithBackendAndPathFunction_Once sync.Once

func settingsNewWithBackendAndPathFunction_Set() error {
	var err error
	settingsNewWithBackendAndPathFunction_Once.Do(func() {
		err = settingsObject_Set()
		if err != nil {
			return
		}
		settingsNewWithBackendAndPathFunction, err = settingsObject.InvokerNew("new_with_backend_and_path")
	})
	return err
}

// SettingsNewWithBackendAndPath is a representation of the C type g_settings_new_with_backend_and_path.
func SettingsNewWithBackendAndPath(schemaId string, backend *SettingsBackend, path string) *Settings {
	var inArgs [3]gi.Argument
	inArgs[0].SetString(schemaId)
	inArgs[1].SetPointer(backend.Native())
	inArgs[2].SetString(path)

	var ret gi.Argument

	err := settingsNewWithBackendAndPathFunction_Set()
	if err == nil {
		ret = settingsNewWithBackendAndPathFunction.Invoke(inArgs[:], nil)
	}

	retGo := SettingsNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

var settingsNewWithPathFunction *gi.Function
var settingsNewWithPathFunction_Once sync.Once

func settingsNewWithPathFunction_Set() error {
	var err error
	settingsNewWithPathFunction_Once.Do(func() {
		err = settingsObject_Set()
		if err != nil {
			return
		}
		settingsNewWithPathFunction, err = settingsObject.InvokerNew("new_with_path")
	})
	return err
}

// SettingsNewWithPath is a representation of the C type g_settings_new_with_path.
func SettingsNewWithPath(schemaId string, path string) *Settings {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(schemaId)
	inArgs[1].SetString(path)

	var ret gi.Argument

	err := settingsNewWithPathFunction_Set()
	if err == nil {
		ret = settingsNewWithPathFunction.Invoke(inArgs[:], nil)
	}

	retGo := SettingsNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

var settingsApplyFunction *gi.Function
var settingsApplyFunction_Once sync.Once

func settingsApplyFunction_Set() error {
	var err error
	settingsApplyFunction_Once.Do(func() {
		err = settingsObject_Set()
		if err != nil {
			return
		}
		settingsApplyFunction, err = settingsObject.InvokerNew("apply")
	})
	return err
}

// Apply is a representation of the C type g_settings_apply.
func (recv *Settings) Apply() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := settingsApplyFunction_Set()
	if err == nil {
		settingsApplyFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'g_settings_bind' : parameter 'flags' of type 'SettingsBindFlags' not supported

// UNSUPPORTED : C value 'g_settings_bind_with_mapping' : parameter 'flags' of type 'SettingsBindFlags' not supported

var settingsBindWritableFunction *gi.Function
var settingsBindWritableFunction_Once sync.Once

func settingsBindWritableFunction_Set() error {
	var err error
	settingsBindWritableFunction_Once.Do(func() {
		err = settingsObject_Set()
		if err != nil {
			return
		}
		settingsBindWritableFunction, err = settingsObject.InvokerNew("bind_writable")
	})
	return err
}

// BindWritable is a representation of the C type g_settings_bind_writable.
func (recv *Settings) BindWritable(key string, object_ *gobject.Object, property string, inverted bool) {
	var inArgs [5]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(key)
	inArgs[2].SetPointer(object_.Native())
	inArgs[3].SetString(property)
	inArgs[4].SetBoolean(inverted)

	err := settingsBindWritableFunction_Set()
	if err == nil {
		settingsBindWritableFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'g_settings_create_action' : return type 'Action' not supported

var settingsDelayFunction *gi.Function
var settingsDelayFunction_Once sync.Once

func settingsDelayFunction_Set() error {
	var err error
	settingsDelayFunction_Once.Do(func() {
		err = settingsObject_Set()
		if err != nil {
			return
		}
		settingsDelayFunction, err = settingsObject.InvokerNew("delay")
	})
	return err
}

// Delay is a representation of the C type g_settings_delay.
func (recv *Settings) Delay() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := settingsDelayFunction_Set()
	if err == nil {
		settingsDelayFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'g_settings_get' : parameter '...' of type 'nil' not supported

var settingsGetBooleanFunction *gi.Function
var settingsGetBooleanFunction_Once sync.Once

func settingsGetBooleanFunction_Set() error {
	var err error
	settingsGetBooleanFunction_Once.Do(func() {
		err = settingsObject_Set()
		if err != nil {
			return
		}
		settingsGetBooleanFunction, err = settingsObject.InvokerNew("get_boolean")
	})
	return err
}

// GetBoolean is a representation of the C type g_settings_get_boolean.
func (recv *Settings) GetBoolean(key string) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(key)

	var ret gi.Argument

	err := settingsGetBooleanFunction_Set()
	if err == nil {
		ret = settingsGetBooleanFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var settingsGetChildFunction *gi.Function
var settingsGetChildFunction_Once sync.Once

func settingsGetChildFunction_Set() error {
	var err error
	settingsGetChildFunction_Once.Do(func() {
		err = settingsObject_Set()
		if err != nil {
			return
		}
		settingsGetChildFunction, err = settingsObject.InvokerNew("get_child")
	})
	return err
}

// GetChild is a representation of the C type g_settings_get_child.
func (recv *Settings) GetChild(name string) *Settings {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(name)

	var ret gi.Argument

	err := settingsGetChildFunction_Set()
	if err == nil {
		ret = settingsGetChildFunction.Invoke(inArgs[:], nil)
	}

	retGo := SettingsNewFromNative(ret.Pointer())

	return retGo
}

// UNSUPPORTED : C value 'g_settings_get_default_value' : return type 'GLib.Variant' not supported

var settingsGetDoubleFunction *gi.Function
var settingsGetDoubleFunction_Once sync.Once

func settingsGetDoubleFunction_Set() error {
	var err error
	settingsGetDoubleFunction_Once.Do(func() {
		err = settingsObject_Set()
		if err != nil {
			return
		}
		settingsGetDoubleFunction, err = settingsObject.InvokerNew("get_double")
	})
	return err
}

// GetDouble is a representation of the C type g_settings_get_double.
func (recv *Settings) GetDouble(key string) float64 {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(key)

	var ret gi.Argument

	err := settingsGetDoubleFunction_Set()
	if err == nil {
		ret = settingsGetDoubleFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Float64()

	return retGo
}

var settingsGetEnumFunction *gi.Function
var settingsGetEnumFunction_Once sync.Once

func settingsGetEnumFunction_Set() error {
	var err error
	settingsGetEnumFunction_Once.Do(func() {
		err = settingsObject_Set()
		if err != nil {
			return
		}
		settingsGetEnumFunction, err = settingsObject.InvokerNew("get_enum")
	})
	return err
}

// GetEnum is a representation of the C type g_settings_get_enum.
func (recv *Settings) GetEnum(key string) int32 {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(key)

	var ret gi.Argument

	err := settingsGetEnumFunction_Set()
	if err == nil {
		ret = settingsGetEnumFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var settingsGetFlagsFunction *gi.Function
var settingsGetFlagsFunction_Once sync.Once

func settingsGetFlagsFunction_Set() error {
	var err error
	settingsGetFlagsFunction_Once.Do(func() {
		err = settingsObject_Set()
		if err != nil {
			return
		}
		settingsGetFlagsFunction, err = settingsObject.InvokerNew("get_flags")
	})
	return err
}

// GetFlags is a representation of the C type g_settings_get_flags.
func (recv *Settings) GetFlags(key string) uint32 {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(key)

	var ret gi.Argument

	err := settingsGetFlagsFunction_Set()
	if err == nil {
		ret = settingsGetFlagsFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo
}

var settingsGetHasUnappliedFunction *gi.Function
var settingsGetHasUnappliedFunction_Once sync.Once

func settingsGetHasUnappliedFunction_Set() error {
	var err error
	settingsGetHasUnappliedFunction_Once.Do(func() {
		err = settingsObject_Set()
		if err != nil {
			return
		}
		settingsGetHasUnappliedFunction, err = settingsObject.InvokerNew("get_has_unapplied")
	})
	return err
}

// GetHasUnapplied is a representation of the C type g_settings_get_has_unapplied.
func (recv *Settings) GetHasUnapplied() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := settingsGetHasUnappliedFunction_Set()
	if err == nil {
		ret = settingsGetHasUnappliedFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var settingsGetIntFunction *gi.Function
var settingsGetIntFunction_Once sync.Once

func settingsGetIntFunction_Set() error {
	var err error
	settingsGetIntFunction_Once.Do(func() {
		err = settingsObject_Set()
		if err != nil {
			return
		}
		settingsGetIntFunction, err = settingsObject.InvokerNew("get_int")
	})
	return err
}

// GetInt is a representation of the C type g_settings_get_int.
func (recv *Settings) GetInt(key string) int32 {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(key)

	var ret gi.Argument

	err := settingsGetIntFunction_Set()
	if err == nil {
		ret = settingsGetIntFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var settingsGetInt64Function *gi.Function
var settingsGetInt64Function_Once sync.Once

func settingsGetInt64Function_Set() error {
	var err error
	settingsGetInt64Function_Once.Do(func() {
		err = settingsObject_Set()
		if err != nil {
			return
		}
		settingsGetInt64Function, err = settingsObject.InvokerNew("get_int64")
	})
	return err
}

// GetInt64 is a representation of the C type g_settings_get_int64.
func (recv *Settings) GetInt64(key string) int64 {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(key)

	var ret gi.Argument

	err := settingsGetInt64Function_Set()
	if err == nil {
		ret = settingsGetInt64Function.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int64()

	return retGo
}

// UNSUPPORTED : C value 'g_settings_get_mapped' : parameter 'mapping' of type 'SettingsGetMapping' not supported

// UNSUPPORTED : C value 'g_settings_get_range' : return type 'GLib.Variant' not supported

var settingsGetStringFunction *gi.Function
var settingsGetStringFunction_Once sync.Once

func settingsGetStringFunction_Set() error {
	var err error
	settingsGetStringFunction_Once.Do(func() {
		err = settingsObject_Set()
		if err != nil {
			return
		}
		settingsGetStringFunction, err = settingsObject.InvokerNew("get_string")
	})
	return err
}

// GetString is a representation of the C type g_settings_get_string.
func (recv *Settings) GetString(key string) string {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(key)

	var ret gi.Argument

	err := settingsGetStringFunction_Set()
	if err == nil {
		ret = settingsGetStringFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

var settingsGetStrvFunction *gi.Function
var settingsGetStrvFunction_Once sync.Once

func settingsGetStrvFunction_Set() error {
	var err error
	settingsGetStrvFunction_Once.Do(func() {
		err = settingsObject_Set()
		if err != nil {
			return
		}
		settingsGetStrvFunction, err = settingsObject.InvokerNew("get_strv")
	})
	return err
}

// GetStrv is a representation of the C type g_settings_get_strv.
func (recv *Settings) GetStrv(key string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(key)

	err := settingsGetStrvFunction_Set()
	if err == nil {
		settingsGetStrvFunction.Invoke(inArgs[:], nil)
	}

	return
}

var settingsGetUintFunction *gi.Function
var settingsGetUintFunction_Once sync.Once

func settingsGetUintFunction_Set() error {
	var err error
	settingsGetUintFunction_Once.Do(func() {
		err = settingsObject_Set()
		if err != nil {
			return
		}
		settingsGetUintFunction, err = settingsObject.InvokerNew("get_uint")
	})
	return err
}

// GetUint is a representation of the C type g_settings_get_uint.
func (recv *Settings) GetUint(key string) uint32 {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(key)

	var ret gi.Argument

	err := settingsGetUintFunction_Set()
	if err == nil {
		ret = settingsGetUintFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo
}

var settingsGetUint64Function *gi.Function
var settingsGetUint64Function_Once sync.Once

func settingsGetUint64Function_Set() error {
	var err error
	settingsGetUint64Function_Once.Do(func() {
		err = settingsObject_Set()
		if err != nil {
			return
		}
		settingsGetUint64Function, err = settingsObject.InvokerNew("get_uint64")
	})
	return err
}

// GetUint64 is a representation of the C type g_settings_get_uint64.
func (recv *Settings) GetUint64(key string) uint64 {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(key)

	var ret gi.Argument

	err := settingsGetUint64Function_Set()
	if err == nil {
		ret = settingsGetUint64Function.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint64()

	return retGo
}

// UNSUPPORTED : C value 'g_settings_get_user_value' : return type 'GLib.Variant' not supported

// UNSUPPORTED : C value 'g_settings_get_value' : return type 'GLib.Variant' not supported

var settingsIsWritableFunction *gi.Function
var settingsIsWritableFunction_Once sync.Once

func settingsIsWritableFunction_Set() error {
	var err error
	settingsIsWritableFunction_Once.Do(func() {
		err = settingsObject_Set()
		if err != nil {
			return
		}
		settingsIsWritableFunction, err = settingsObject.InvokerNew("is_writable")
	})
	return err
}

// IsWritable is a representation of the C type g_settings_is_writable.
func (recv *Settings) IsWritable(name string) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(name)

	var ret gi.Argument

	err := settingsIsWritableFunction_Set()
	if err == nil {
		ret = settingsIsWritableFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var settingsListChildrenFunction *gi.Function
var settingsListChildrenFunction_Once sync.Once

func settingsListChildrenFunction_Set() error {
	var err error
	settingsListChildrenFunction_Once.Do(func() {
		err = settingsObject_Set()
		if err != nil {
			return
		}
		settingsListChildrenFunction, err = settingsObject.InvokerNew("list_children")
	})
	return err
}

// ListChildren is a representation of the C type g_settings_list_children.
func (recv *Settings) ListChildren() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := settingsListChildrenFunction_Set()
	if err == nil {
		settingsListChildrenFunction.Invoke(inArgs[:], nil)
	}

	return
}

var settingsListKeysFunction *gi.Function
var settingsListKeysFunction_Once sync.Once

func settingsListKeysFunction_Set() error {
	var err error
	settingsListKeysFunction_Once.Do(func() {
		err = settingsObject_Set()
		if err != nil {
			return
		}
		settingsListKeysFunction, err = settingsObject.InvokerNew("list_keys")
	})
	return err
}

// ListKeys is a representation of the C type g_settings_list_keys.
func (recv *Settings) ListKeys() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := settingsListKeysFunction_Set()
	if err == nil {
		settingsListKeysFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'g_settings_range_check' : parameter 'value' of type 'GLib.Variant' not supported

var settingsResetFunction *gi.Function
var settingsResetFunction_Once sync.Once

func settingsResetFunction_Set() error {
	var err error
	settingsResetFunction_Once.Do(func() {
		err = settingsObject_Set()
		if err != nil {
			return
		}
		settingsResetFunction, err = settingsObject.InvokerNew("reset")
	})
	return err
}

// Reset is a representation of the C type g_settings_reset.
func (recv *Settings) Reset(key string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(key)

	err := settingsResetFunction_Set()
	if err == nil {
		settingsResetFunction.Invoke(inArgs[:], nil)
	}

	return
}

var settingsRevertFunction *gi.Function
var settingsRevertFunction_Once sync.Once

func settingsRevertFunction_Set() error {
	var err error
	settingsRevertFunction_Once.Do(func() {
		err = settingsObject_Set()
		if err != nil {
			return
		}
		settingsRevertFunction, err = settingsObject.InvokerNew("revert")
	})
	return err
}

// Revert is a representation of the C type g_settings_revert.
func (recv *Settings) Revert() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := settingsRevertFunction_Set()
	if err == nil {
		settingsRevertFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'g_settings_set' : parameter '...' of type 'nil' not supported

var settingsSetBooleanFunction *gi.Function
var settingsSetBooleanFunction_Once sync.Once

func settingsSetBooleanFunction_Set() error {
	var err error
	settingsSetBooleanFunction_Once.Do(func() {
		err = settingsObject_Set()
		if err != nil {
			return
		}
		settingsSetBooleanFunction, err = settingsObject.InvokerNew("set_boolean")
	})
	return err
}

// SetBoolean is a representation of the C type g_settings_set_boolean.
func (recv *Settings) SetBoolean(key string, value bool) bool {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(key)
	inArgs[2].SetBoolean(value)

	var ret gi.Argument

	err := settingsSetBooleanFunction_Set()
	if err == nil {
		ret = settingsSetBooleanFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var settingsSetDoubleFunction *gi.Function
var settingsSetDoubleFunction_Once sync.Once

func settingsSetDoubleFunction_Set() error {
	var err error
	settingsSetDoubleFunction_Once.Do(func() {
		err = settingsObject_Set()
		if err != nil {
			return
		}
		settingsSetDoubleFunction, err = settingsObject.InvokerNew("set_double")
	})
	return err
}

// SetDouble is a representation of the C type g_settings_set_double.
func (recv *Settings) SetDouble(key string, value float64) bool {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(key)
	inArgs[2].SetFloat64(value)

	var ret gi.Argument

	err := settingsSetDoubleFunction_Set()
	if err == nil {
		ret = settingsSetDoubleFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var settingsSetEnumFunction *gi.Function
var settingsSetEnumFunction_Once sync.Once

func settingsSetEnumFunction_Set() error {
	var err error
	settingsSetEnumFunction_Once.Do(func() {
		err = settingsObject_Set()
		if err != nil {
			return
		}
		settingsSetEnumFunction, err = settingsObject.InvokerNew("set_enum")
	})
	return err
}

// SetEnum is a representation of the C type g_settings_set_enum.
func (recv *Settings) SetEnum(key string, value int32) bool {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(key)
	inArgs[2].SetInt32(value)

	var ret gi.Argument

	err := settingsSetEnumFunction_Set()
	if err == nil {
		ret = settingsSetEnumFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var settingsSetFlagsFunction *gi.Function
var settingsSetFlagsFunction_Once sync.Once

func settingsSetFlagsFunction_Set() error {
	var err error
	settingsSetFlagsFunction_Once.Do(func() {
		err = settingsObject_Set()
		if err != nil {
			return
		}
		settingsSetFlagsFunction, err = settingsObject.InvokerNew("set_flags")
	})
	return err
}

// SetFlags is a representation of the C type g_settings_set_flags.
func (recv *Settings) SetFlags(key string, value uint32) bool {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(key)
	inArgs[2].SetUint32(value)

	var ret gi.Argument

	err := settingsSetFlagsFunction_Set()
	if err == nil {
		ret = settingsSetFlagsFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var settingsSetIntFunction *gi.Function
var settingsSetIntFunction_Once sync.Once

func settingsSetIntFunction_Set() error {
	var err error
	settingsSetIntFunction_Once.Do(func() {
		err = settingsObject_Set()
		if err != nil {
			return
		}
		settingsSetIntFunction, err = settingsObject.InvokerNew("set_int")
	})
	return err
}

// SetInt is a representation of the C type g_settings_set_int.
func (recv *Settings) SetInt(key string, value int32) bool {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(key)
	inArgs[2].SetInt32(value)

	var ret gi.Argument

	err := settingsSetIntFunction_Set()
	if err == nil {
		ret = settingsSetIntFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var settingsSetInt64Function *gi.Function
var settingsSetInt64Function_Once sync.Once

func settingsSetInt64Function_Set() error {
	var err error
	settingsSetInt64Function_Once.Do(func() {
		err = settingsObject_Set()
		if err != nil {
			return
		}
		settingsSetInt64Function, err = settingsObject.InvokerNew("set_int64")
	})
	return err
}

// SetInt64 is a representation of the C type g_settings_set_int64.
func (recv *Settings) SetInt64(key string, value int64) bool {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(key)
	inArgs[2].SetInt64(value)

	var ret gi.Argument

	err := settingsSetInt64Function_Set()
	if err == nil {
		ret = settingsSetInt64Function.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var settingsSetStringFunction *gi.Function
var settingsSetStringFunction_Once sync.Once

func settingsSetStringFunction_Set() error {
	var err error
	settingsSetStringFunction_Once.Do(func() {
		err = settingsObject_Set()
		if err != nil {
			return
		}
		settingsSetStringFunction, err = settingsObject.InvokerNew("set_string")
	})
	return err
}

// SetString is a representation of the C type g_settings_set_string.
func (recv *Settings) SetString(key string, value string) bool {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(key)
	inArgs[2].SetString(value)

	var ret gi.Argument

	err := settingsSetStringFunction_Set()
	if err == nil {
		ret = settingsSetStringFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'g_settings_set_strv' : parameter 'value' of type 'nil' not supported

var settingsSetUintFunction *gi.Function
var settingsSetUintFunction_Once sync.Once

func settingsSetUintFunction_Set() error {
	var err error
	settingsSetUintFunction_Once.Do(func() {
		err = settingsObject_Set()
		if err != nil {
			return
		}
		settingsSetUintFunction, err = settingsObject.InvokerNew("set_uint")
	})
	return err
}

// SetUint is a representation of the C type g_settings_set_uint.
func (recv *Settings) SetUint(key string, value uint32) bool {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(key)
	inArgs[2].SetUint32(value)

	var ret gi.Argument

	err := settingsSetUintFunction_Set()
	if err == nil {
		ret = settingsSetUintFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var settingsSetUint64Function *gi.Function
var settingsSetUint64Function_Once sync.Once

func settingsSetUint64Function_Set() error {
	var err error
	settingsSetUint64Function_Once.Do(func() {
		err = settingsObject_Set()
		if err != nil {
			return
		}
		settingsSetUint64Function, err = settingsObject.InvokerNew("set_uint64")
	})
	return err
}

// SetUint64 is a representation of the C type g_settings_set_uint64.
func (recv *Settings) SetUint64(key string, value uint64) bool {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(key)
	inArgs[2].SetUint64(value)

	var ret gi.Argument

	err := settingsSetUint64Function_Set()
	if err == nil {
		ret = settingsSetUint64Function.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'g_settings_set_value' : parameter 'value' of type 'GLib.Variant' not supported

var settingsBackendObject *gi.Object
var settingsBackendObject_Once sync.Once

func settingsBackendObject_Set() error {
	var err error
	settingsBackendObject_Once.Do(func() {
		settingsBackendObject, err = gi.ObjectNew("Gio", "SettingsBackend")
	})
	return err
}

type SettingsBackend struct {
	native unsafe.Pointer
}

func SettingsBackendNewFromNative(native unsafe.Pointer) *SettingsBackend {
	instance := &SettingsBackend{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *SettingsBackend) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToSettingsBackend down casts any arbitrary Object to SettingsBackend.
Exercise care, as this is a potentially dangerous function
if the Object is not a SettingsBackend.
*/
func CastToSettingsBackend(object *gobject.Object) *SettingsBackend {
	return SettingsBackendNewFromNative(object.Native())
}

// Equals compares this SettingsBackend with another SettingsBackend, and returns true if they represent the same GObject.
func (recv *SettingsBackend) Equals(other *SettingsBackend) bool {
	return other.Native() == recv.Native()
}

func (recv *SettingsBackend) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *SettingsBackend) FieldParentInstance() *gobject.Object {
	argValue := gi.ObjectFieldGet(settingsBackendObject, recv.Native(), "parent_instance")
	value := gobject.ObjectNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *SettingsBackend) SetFieldParentInstance(value *gobject.Object) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(settingsBackendObject, recv.Native(), "parent_instance", argValue)
}

// UNSUPPORTED : C value 'g_settings_backend_changed' : parameter 'origin_tag' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_settings_backend_changed_tree' : parameter 'tree' of type 'GLib.Tree' not supported

// UNSUPPORTED : C value 'g_settings_backend_keys_changed' : parameter 'items' of type 'nil' not supported

// UNSUPPORTED : C value 'g_settings_backend_path_changed' : parameter 'origin_tag' of type 'gpointer' not supported

var settingsBackendPathWritableChangedFunction *gi.Function
var settingsBackendPathWritableChangedFunction_Once sync.Once

func settingsBackendPathWritableChangedFunction_Set() error {
	var err error
	settingsBackendPathWritableChangedFunction_Once.Do(func() {
		err = settingsBackendObject_Set()
		if err != nil {
			return
		}
		settingsBackendPathWritableChangedFunction, err = settingsBackendObject.InvokerNew("path_writable_changed")
	})
	return err
}

// PathWritableChanged is a representation of the C type g_settings_backend_path_writable_changed.
func (recv *SettingsBackend) PathWritableChanged(path string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(path)

	err := settingsBackendPathWritableChangedFunction_Set()
	if err == nil {
		settingsBackendPathWritableChangedFunction.Invoke(inArgs[:], nil)
	}

	return
}

var settingsBackendWritableChangedFunction *gi.Function
var settingsBackendWritableChangedFunction_Once sync.Once

func settingsBackendWritableChangedFunction_Set() error {
	var err error
	settingsBackendWritableChangedFunction_Once.Do(func() {
		err = settingsBackendObject_Set()
		if err != nil {
			return
		}
		settingsBackendWritableChangedFunction, err = settingsBackendObject.InvokerNew("writable_changed")
	})
	return err
}

// WritableChanged is a representation of the C type g_settings_backend_writable_changed.
func (recv *SettingsBackend) WritableChanged(key string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(key)

	err := settingsBackendWritableChangedFunction_Set()
	if err == nil {
		settingsBackendWritableChangedFunction.Invoke(inArgs[:], nil)
	}

	return
}

var simpleActionObject *gi.Object
var simpleActionObject_Once sync.Once

func simpleActionObject_Set() error {
	var err error
	simpleActionObject_Once.Do(func() {
		simpleActionObject, err = gi.ObjectNew("Gio", "SimpleAction")
	})
	return err
}

type SimpleAction struct {
	native unsafe.Pointer
}

func SimpleActionNewFromNative(native unsafe.Pointer) *SimpleAction {
	instance := &SimpleAction{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *SimpleAction) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToSimpleAction down casts any arbitrary Object to SimpleAction.
Exercise care, as this is a potentially dangerous function
if the Object is not a SimpleAction.
*/
func CastToSimpleAction(object *gobject.Object) *SimpleAction {
	return SimpleActionNewFromNative(object.Native())
}

// Equals compares this SimpleAction with another SimpleAction, and returns true if they represent the same GObject.
func (recv *SimpleAction) Equals(other *SimpleAction) bool {
	return other.Native() == recv.Native()
}

func (recv *SimpleAction) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'g_simple_action_new' : parameter 'parameter_type' of type 'GLib.VariantType' not supported

// UNSUPPORTED : C value 'g_simple_action_new_stateful' : parameter 'parameter_type' of type 'GLib.VariantType' not supported

var simpleActionSetEnabledFunction *gi.Function
var simpleActionSetEnabledFunction_Once sync.Once

func simpleActionSetEnabledFunction_Set() error {
	var err error
	simpleActionSetEnabledFunction_Once.Do(func() {
		err = simpleActionObject_Set()
		if err != nil {
			return
		}
		simpleActionSetEnabledFunction, err = simpleActionObject.InvokerNew("set_enabled")
	})
	return err
}

// SetEnabled is a representation of the C type g_simple_action_set_enabled.
func (recv *SimpleAction) SetEnabled(enabled bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetBoolean(enabled)

	err := simpleActionSetEnabledFunction_Set()
	if err == nil {
		simpleActionSetEnabledFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'g_simple_action_set_state' : parameter 'value' of type 'GLib.Variant' not supported

// UNSUPPORTED : C value 'g_simple_action_set_state_hint' : parameter 'state_hint' of type 'GLib.Variant' not supported

var simpleActionGroupObject *gi.Object
var simpleActionGroupObject_Once sync.Once

func simpleActionGroupObject_Set() error {
	var err error
	simpleActionGroupObject_Once.Do(func() {
		simpleActionGroupObject, err = gi.ObjectNew("Gio", "SimpleActionGroup")
	})
	return err
}

type SimpleActionGroup struct {
	native unsafe.Pointer
}

func SimpleActionGroupNewFromNative(native unsafe.Pointer) *SimpleActionGroup {
	instance := &SimpleActionGroup{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *SimpleActionGroup) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToSimpleActionGroup down casts any arbitrary Object to SimpleActionGroup.
Exercise care, as this is a potentially dangerous function
if the Object is not a SimpleActionGroup.
*/
func CastToSimpleActionGroup(object *gobject.Object) *SimpleActionGroup {
	return SimpleActionGroupNewFromNative(object.Native())
}

// Equals compares this SimpleActionGroup with another SimpleActionGroup, and returns true if they represent the same GObject.
func (recv *SimpleActionGroup) Equals(other *SimpleActionGroup) bool {
	return other.Native() == recv.Native()
}

func (recv *SimpleActionGroup) Native() unsafe.Pointer {
	return recv.native
}

var simpleActionGroupNewFunction *gi.Function
var simpleActionGroupNewFunction_Once sync.Once

func simpleActionGroupNewFunction_Set() error {
	var err error
	simpleActionGroupNewFunction_Once.Do(func() {
		err = simpleActionGroupObject_Set()
		if err != nil {
			return
		}
		simpleActionGroupNewFunction, err = simpleActionGroupObject.InvokerNew("new")
	})
	return err
}

// SimpleActionGroupNew is a representation of the C type g_simple_action_group_new.
func SimpleActionGroupNew() *SimpleActionGroup {

	var ret gi.Argument

	err := simpleActionGroupNewFunction_Set()
	if err == nil {
		ret = simpleActionGroupNewFunction.Invoke(nil, nil)
	}

	retGo := SimpleActionGroupNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

// UNSUPPORTED : C value 'g_simple_action_group_add_entries' : parameter 'entries' of type 'nil' not supported

// UNSUPPORTED : C value 'g_simple_action_group_insert' : parameter 'action' of type 'Action' not supported

// UNSUPPORTED : C value 'g_simple_action_group_lookup' : return type 'Action' not supported

var simpleActionGroupRemoveFunction *gi.Function
var simpleActionGroupRemoveFunction_Once sync.Once

func simpleActionGroupRemoveFunction_Set() error {
	var err error
	simpleActionGroupRemoveFunction_Once.Do(func() {
		err = simpleActionGroupObject_Set()
		if err != nil {
			return
		}
		simpleActionGroupRemoveFunction, err = simpleActionGroupObject.InvokerNew("remove")
	})
	return err
}

// Remove is a representation of the C type g_simple_action_group_remove.
func (recv *SimpleActionGroup) Remove(actionName string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(actionName)

	err := simpleActionGroupRemoveFunction_Set()
	if err == nil {
		simpleActionGroupRemoveFunction.Invoke(inArgs[:], nil)
	}

	return
}

var simpleAsyncResultObject *gi.Object
var simpleAsyncResultObject_Once sync.Once

func simpleAsyncResultObject_Set() error {
	var err error
	simpleAsyncResultObject_Once.Do(func() {
		simpleAsyncResultObject, err = gi.ObjectNew("Gio", "SimpleAsyncResult")
	})
	return err
}

type SimpleAsyncResult struct {
	native unsafe.Pointer
}

func SimpleAsyncResultNewFromNative(native unsafe.Pointer) *SimpleAsyncResult {
	instance := &SimpleAsyncResult{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *SimpleAsyncResult) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToSimpleAsyncResult down casts any arbitrary Object to SimpleAsyncResult.
Exercise care, as this is a potentially dangerous function
if the Object is not a SimpleAsyncResult.
*/
func CastToSimpleAsyncResult(object *gobject.Object) *SimpleAsyncResult {
	return SimpleAsyncResultNewFromNative(object.Native())
}

// Equals compares this SimpleAsyncResult with another SimpleAsyncResult, and returns true if they represent the same GObject.
func (recv *SimpleAsyncResult) Equals(other *SimpleAsyncResult) bool {
	return other.Native() == recv.Native()
}

func (recv *SimpleAsyncResult) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'g_simple_async_result_new' : parameter 'callback' of type 'AsyncReadyCallback' not supported

// UNSUPPORTED : C value 'g_simple_async_result_new_error' : parameter 'callback' of type 'AsyncReadyCallback' not supported

// UNSUPPORTED : C value 'g_simple_async_result_new_from_error' : parameter 'callback' of type 'AsyncReadyCallback' not supported

// UNSUPPORTED : C value 'g_simple_async_result_new_take_error' : parameter 'callback' of type 'AsyncReadyCallback' not supported

var simpleAsyncResultCompleteFunction *gi.Function
var simpleAsyncResultCompleteFunction_Once sync.Once

func simpleAsyncResultCompleteFunction_Set() error {
	var err error
	simpleAsyncResultCompleteFunction_Once.Do(func() {
		err = simpleAsyncResultObject_Set()
		if err != nil {
			return
		}
		simpleAsyncResultCompleteFunction, err = simpleAsyncResultObject.InvokerNew("complete")
	})
	return err
}

// Complete is a representation of the C type g_simple_async_result_complete.
func (recv *SimpleAsyncResult) Complete() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := simpleAsyncResultCompleteFunction_Set()
	if err == nil {
		simpleAsyncResultCompleteFunction.Invoke(inArgs[:], nil)
	}

	return
}

var simpleAsyncResultCompleteInIdleFunction *gi.Function
var simpleAsyncResultCompleteInIdleFunction_Once sync.Once

func simpleAsyncResultCompleteInIdleFunction_Set() error {
	var err error
	simpleAsyncResultCompleteInIdleFunction_Once.Do(func() {
		err = simpleAsyncResultObject_Set()
		if err != nil {
			return
		}
		simpleAsyncResultCompleteInIdleFunction, err = simpleAsyncResultObject.InvokerNew("complete_in_idle")
	})
	return err
}

// CompleteInIdle is a representation of the C type g_simple_async_result_complete_in_idle.
func (recv *SimpleAsyncResult) CompleteInIdle() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := simpleAsyncResultCompleteInIdleFunction_Set()
	if err == nil {
		simpleAsyncResultCompleteInIdleFunction.Invoke(inArgs[:], nil)
	}

	return
}

var simpleAsyncResultGetOpResGbooleanFunction *gi.Function
var simpleAsyncResultGetOpResGbooleanFunction_Once sync.Once

func simpleAsyncResultGetOpResGbooleanFunction_Set() error {
	var err error
	simpleAsyncResultGetOpResGbooleanFunction_Once.Do(func() {
		err = simpleAsyncResultObject_Set()
		if err != nil {
			return
		}
		simpleAsyncResultGetOpResGbooleanFunction, err = simpleAsyncResultObject.InvokerNew("get_op_res_gboolean")
	})
	return err
}

// GetOpResGboolean is a representation of the C type g_simple_async_result_get_op_res_gboolean.
func (recv *SimpleAsyncResult) GetOpResGboolean() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := simpleAsyncResultGetOpResGbooleanFunction_Set()
	if err == nil {
		ret = simpleAsyncResultGetOpResGbooleanFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'g_simple_async_result_get_op_res_gpointer' : return type 'gpointer' not supported

var simpleAsyncResultGetOpResGssizeFunction *gi.Function
var simpleAsyncResultGetOpResGssizeFunction_Once sync.Once

func simpleAsyncResultGetOpResGssizeFunction_Set() error {
	var err error
	simpleAsyncResultGetOpResGssizeFunction_Once.Do(func() {
		err = simpleAsyncResultObject_Set()
		if err != nil {
			return
		}
		simpleAsyncResultGetOpResGssizeFunction, err = simpleAsyncResultObject.InvokerNew("get_op_res_gssize")
	})
	return err
}

// GetOpResGssize is a representation of the C type g_simple_async_result_get_op_res_gssize.
func (recv *SimpleAsyncResult) GetOpResGssize() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := simpleAsyncResultGetOpResGssizeFunction_Set()
	if err == nil {
		ret = simpleAsyncResultGetOpResGssizeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

// UNSUPPORTED : C value 'g_simple_async_result_get_source_tag' : return type 'gpointer' not supported

var simpleAsyncResultPropagateErrorFunction *gi.Function
var simpleAsyncResultPropagateErrorFunction_Once sync.Once

func simpleAsyncResultPropagateErrorFunction_Set() error {
	var err error
	simpleAsyncResultPropagateErrorFunction_Once.Do(func() {
		err = simpleAsyncResultObject_Set()
		if err != nil {
			return
		}
		simpleAsyncResultPropagateErrorFunction, err = simpleAsyncResultObject.InvokerNew("propagate_error")
	})
	return err
}

// PropagateError is a representation of the C type g_simple_async_result_propagate_error.
func (recv *SimpleAsyncResult) PropagateError() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := simpleAsyncResultPropagateErrorFunction_Set()
	if err == nil {
		ret = simpleAsyncResultPropagateErrorFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'g_simple_async_result_run_in_thread' : parameter 'func' of type 'SimpleAsyncThreadFunc' not supported

var simpleAsyncResultSetCheckCancellableFunction *gi.Function
var simpleAsyncResultSetCheckCancellableFunction_Once sync.Once

func simpleAsyncResultSetCheckCancellableFunction_Set() error {
	var err error
	simpleAsyncResultSetCheckCancellableFunction_Once.Do(func() {
		err = simpleAsyncResultObject_Set()
		if err != nil {
			return
		}
		simpleAsyncResultSetCheckCancellableFunction, err = simpleAsyncResultObject.InvokerNew("set_check_cancellable")
	})
	return err
}

// SetCheckCancellable is a representation of the C type g_simple_async_result_set_check_cancellable.
func (recv *SimpleAsyncResult) SetCheckCancellable(checkCancellable *Cancellable) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(checkCancellable.Native())

	err := simpleAsyncResultSetCheckCancellableFunction_Set()
	if err == nil {
		simpleAsyncResultSetCheckCancellableFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'g_simple_async_result_set_error' : parameter '...' of type 'nil' not supported

// UNSUPPORTED : C value 'g_simple_async_result_set_error_va' : parameter 'args' of type 'va_list' not supported

// UNSUPPORTED : C value 'g_simple_async_result_set_from_error' : parameter 'error' of type 'GLib.Error' not supported

var simpleAsyncResultSetHandleCancellationFunction *gi.Function
var simpleAsyncResultSetHandleCancellationFunction_Once sync.Once

func simpleAsyncResultSetHandleCancellationFunction_Set() error {
	var err error
	simpleAsyncResultSetHandleCancellationFunction_Once.Do(func() {
		err = simpleAsyncResultObject_Set()
		if err != nil {
			return
		}
		simpleAsyncResultSetHandleCancellationFunction, err = simpleAsyncResultObject.InvokerNew("set_handle_cancellation")
	})
	return err
}

// SetHandleCancellation is a representation of the C type g_simple_async_result_set_handle_cancellation.
func (recv *SimpleAsyncResult) SetHandleCancellation(handleCancellation bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetBoolean(handleCancellation)

	err := simpleAsyncResultSetHandleCancellationFunction_Set()
	if err == nil {
		simpleAsyncResultSetHandleCancellationFunction.Invoke(inArgs[:], nil)
	}

	return
}

var simpleAsyncResultSetOpResGbooleanFunction *gi.Function
var simpleAsyncResultSetOpResGbooleanFunction_Once sync.Once

func simpleAsyncResultSetOpResGbooleanFunction_Set() error {
	var err error
	simpleAsyncResultSetOpResGbooleanFunction_Once.Do(func() {
		err = simpleAsyncResultObject_Set()
		if err != nil {
			return
		}
		simpleAsyncResultSetOpResGbooleanFunction, err = simpleAsyncResultObject.InvokerNew("set_op_res_gboolean")
	})
	return err
}

// SetOpResGboolean is a representation of the C type g_simple_async_result_set_op_res_gboolean.
func (recv *SimpleAsyncResult) SetOpResGboolean(opRes bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetBoolean(opRes)

	err := simpleAsyncResultSetOpResGbooleanFunction_Set()
	if err == nil {
		simpleAsyncResultSetOpResGbooleanFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'g_simple_async_result_set_op_res_gpointer' : parameter 'op_res' of type 'gpointer' not supported

var simpleAsyncResultSetOpResGssizeFunction *gi.Function
var simpleAsyncResultSetOpResGssizeFunction_Once sync.Once

func simpleAsyncResultSetOpResGssizeFunction_Set() error {
	var err error
	simpleAsyncResultSetOpResGssizeFunction_Once.Do(func() {
		err = simpleAsyncResultObject_Set()
		if err != nil {
			return
		}
		simpleAsyncResultSetOpResGssizeFunction, err = simpleAsyncResultObject.InvokerNew("set_op_res_gssize")
	})
	return err
}

// SetOpResGssize is a representation of the C type g_simple_async_result_set_op_res_gssize.
func (recv *SimpleAsyncResult) SetOpResGssize(opRes int32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(opRes)

	err := simpleAsyncResultSetOpResGssizeFunction_Set()
	if err == nil {
		simpleAsyncResultSetOpResGssizeFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'g_simple_async_result_take_error' : parameter 'error' of type 'GLib.Error' not supported

var simpleIOStreamObject *gi.Object
var simpleIOStreamObject_Once sync.Once

func simpleIOStreamObject_Set() error {
	var err error
	simpleIOStreamObject_Once.Do(func() {
		simpleIOStreamObject, err = gi.ObjectNew("Gio", "SimpleIOStream")
	})
	return err
}

type SimpleIOStream struct {
	native unsafe.Pointer
}

func SimpleIOStreamNewFromNative(native unsafe.Pointer) *SimpleIOStream {
	instance := &SimpleIOStream{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// IOStream upcasts to *IOStream
func (recv *SimpleIOStream) IOStream() *IOStream {
	return IOStreamNewFromNative(recv.Native())
}

// Object upcasts to *Object
func (recv *SimpleIOStream) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToSimpleIOStream down casts any arbitrary Object to SimpleIOStream.
Exercise care, as this is a potentially dangerous function
if the Object is not a SimpleIOStream.
*/
func CastToSimpleIOStream(object *gobject.Object) *SimpleIOStream {
	return SimpleIOStreamNewFromNative(object.Native())
}

// Equals compares this SimpleIOStream with another SimpleIOStream, and returns true if they represent the same GObject.
func (recv *SimpleIOStream) Equals(other *SimpleIOStream) bool {
	return other.Native() == recv.Native()
}

func (recv *SimpleIOStream) Native() unsafe.Pointer {
	return recv.native
}

var simpleIOStreamNewFunction *gi.Function
var simpleIOStreamNewFunction_Once sync.Once

func simpleIOStreamNewFunction_Set() error {
	var err error
	simpleIOStreamNewFunction_Once.Do(func() {
		err = simpleIOStreamObject_Set()
		if err != nil {
			return
		}
		simpleIOStreamNewFunction, err = simpleIOStreamObject.InvokerNew("new")
	})
	return err
}

// SimpleIOStreamNew is a representation of the C type g_simple_io_stream_new.
func SimpleIOStreamNew(inputStream *InputStream, outputStream *OutputStream) *SimpleIOStream {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(inputStream.Native())
	inArgs[1].SetPointer(outputStream.Native())

	var ret gi.Argument

	err := simpleIOStreamNewFunction_Set()
	if err == nil {
		ret = simpleIOStreamNewFunction.Invoke(inArgs[:], nil)
	}

	retGo := SimpleIOStreamNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

var simplePermissionObject *gi.Object
var simplePermissionObject_Once sync.Once

func simplePermissionObject_Set() error {
	var err error
	simplePermissionObject_Once.Do(func() {
		simplePermissionObject, err = gi.ObjectNew("Gio", "SimplePermission")
	})
	return err
}

type SimplePermission struct {
	native unsafe.Pointer
}

func SimplePermissionNewFromNative(native unsafe.Pointer) *SimplePermission {
	instance := &SimplePermission{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Permission upcasts to *Permission
func (recv *SimplePermission) Permission() *Permission {
	return PermissionNewFromNative(recv.Native())
}

// Object upcasts to *Object
func (recv *SimplePermission) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToSimplePermission down casts any arbitrary Object to SimplePermission.
Exercise care, as this is a potentially dangerous function
if the Object is not a SimplePermission.
*/
func CastToSimplePermission(object *gobject.Object) *SimplePermission {
	return SimplePermissionNewFromNative(object.Native())
}

// Equals compares this SimplePermission with another SimplePermission, and returns true if they represent the same GObject.
func (recv *SimplePermission) Equals(other *SimplePermission) bool {
	return other.Native() == recv.Native()
}

func (recv *SimplePermission) Native() unsafe.Pointer {
	return recv.native
}

var simplePermissionNewFunction *gi.Function
var simplePermissionNewFunction_Once sync.Once

func simplePermissionNewFunction_Set() error {
	var err error
	simplePermissionNewFunction_Once.Do(func() {
		err = simplePermissionObject_Set()
		if err != nil {
			return
		}
		simplePermissionNewFunction, err = simplePermissionObject.InvokerNew("new")
	})
	return err
}

// SimplePermissionNew is a representation of the C type g_simple_permission_new.
func SimplePermissionNew(allowed bool) *SimplePermission {
	var inArgs [1]gi.Argument
	inArgs[0].SetBoolean(allowed)

	var ret gi.Argument

	err := simplePermissionNewFunction_Set()
	if err == nil {
		ret = simplePermissionNewFunction.Invoke(inArgs[:], nil)
	}

	retGo := SimplePermissionNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

var simpleProxyResolverObject *gi.Object
var simpleProxyResolverObject_Once sync.Once

func simpleProxyResolverObject_Set() error {
	var err error
	simpleProxyResolverObject_Once.Do(func() {
		simpleProxyResolverObject, err = gi.ObjectNew("Gio", "SimpleProxyResolver")
	})
	return err
}

type SimpleProxyResolver struct {
	native unsafe.Pointer
}

func SimpleProxyResolverNewFromNative(native unsafe.Pointer) *SimpleProxyResolver {
	instance := &SimpleProxyResolver{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *SimpleProxyResolver) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToSimpleProxyResolver down casts any arbitrary Object to SimpleProxyResolver.
Exercise care, as this is a potentially dangerous function
if the Object is not a SimpleProxyResolver.
*/
func CastToSimpleProxyResolver(object *gobject.Object) *SimpleProxyResolver {
	return SimpleProxyResolverNewFromNative(object.Native())
}

// Equals compares this SimpleProxyResolver with another SimpleProxyResolver, and returns true if they represent the same GObject.
func (recv *SimpleProxyResolver) Equals(other *SimpleProxyResolver) bool {
	return other.Native() == recv.Native()
}

func (recv *SimpleProxyResolver) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *SimpleProxyResolver) FieldParentInstance() *gobject.Object {
	argValue := gi.ObjectFieldGet(simpleProxyResolverObject, recv.Native(), "parent_instance")
	value := gobject.ObjectNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *SimpleProxyResolver) SetFieldParentInstance(value *gobject.Object) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(simpleProxyResolverObject, recv.Native(), "parent_instance", argValue)
}

var simpleProxyResolverSetDefaultProxyFunction *gi.Function
var simpleProxyResolverSetDefaultProxyFunction_Once sync.Once

func simpleProxyResolverSetDefaultProxyFunction_Set() error {
	var err error
	simpleProxyResolverSetDefaultProxyFunction_Once.Do(func() {
		err = simpleProxyResolverObject_Set()
		if err != nil {
			return
		}
		simpleProxyResolverSetDefaultProxyFunction, err = simpleProxyResolverObject.InvokerNew("set_default_proxy")
	})
	return err
}

// SetDefaultProxy is a representation of the C type g_simple_proxy_resolver_set_default_proxy.
func (recv *SimpleProxyResolver) SetDefaultProxy(defaultProxy string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(defaultProxy)

	err := simpleProxyResolverSetDefaultProxyFunction_Set()
	if err == nil {
		simpleProxyResolverSetDefaultProxyFunction.Invoke(inArgs[:], nil)
	}

	return
}

var simpleProxyResolverSetIgnoreHostsFunction *gi.Function
var simpleProxyResolverSetIgnoreHostsFunction_Once sync.Once

func simpleProxyResolverSetIgnoreHostsFunction_Set() error {
	var err error
	simpleProxyResolverSetIgnoreHostsFunction_Once.Do(func() {
		err = simpleProxyResolverObject_Set()
		if err != nil {
			return
		}
		simpleProxyResolverSetIgnoreHostsFunction, err = simpleProxyResolverObject.InvokerNew("set_ignore_hosts")
	})
	return err
}

// SetIgnoreHosts is a representation of the C type g_simple_proxy_resolver_set_ignore_hosts.
func (recv *SimpleProxyResolver) SetIgnoreHosts(ignoreHosts string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(ignoreHosts)

	err := simpleProxyResolverSetIgnoreHostsFunction_Set()
	if err == nil {
		simpleProxyResolverSetIgnoreHostsFunction.Invoke(inArgs[:], nil)
	}

	return
}

var simpleProxyResolverSetUriProxyFunction *gi.Function
var simpleProxyResolverSetUriProxyFunction_Once sync.Once

func simpleProxyResolverSetUriProxyFunction_Set() error {
	var err error
	simpleProxyResolverSetUriProxyFunction_Once.Do(func() {
		err = simpleProxyResolverObject_Set()
		if err != nil {
			return
		}
		simpleProxyResolverSetUriProxyFunction, err = simpleProxyResolverObject.InvokerNew("set_uri_proxy")
	})
	return err
}

// SetUriProxy is a representation of the C type g_simple_proxy_resolver_set_uri_proxy.
func (recv *SimpleProxyResolver) SetUriProxy(uriScheme string, proxy string) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(uriScheme)
	inArgs[2].SetString(proxy)

	err := simpleProxyResolverSetUriProxyFunction_Set()
	if err == nil {
		simpleProxyResolverSetUriProxyFunction.Invoke(inArgs[:], nil)
	}

	return
}

var socketObject *gi.Object
var socketObject_Once sync.Once

func socketObject_Set() error {
	var err error
	socketObject_Once.Do(func() {
		socketObject, err = gi.ObjectNew("Gio", "Socket")
	})
	return err
}

type Socket struct {
	native unsafe.Pointer
}

func SocketNewFromNative(native unsafe.Pointer) *Socket {
	instance := &Socket{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *Socket) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToSocket down casts any arbitrary Object to Socket.
Exercise care, as this is a potentially dangerous function
if the Object is not a Socket.
*/
func CastToSocket(object *gobject.Object) *Socket {
	return SocketNewFromNative(object.Native())
}

// Equals compares this Socket with another Socket, and returns true if they represent the same GObject.
func (recv *Socket) Equals(other *Socket) bool {
	return other.Native() == recv.Native()
}

func (recv *Socket) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *Socket) FieldParentInstance() *gobject.Object {
	argValue := gi.ObjectFieldGet(socketObject, recv.Native(), "parent_instance")
	value := gobject.ObjectNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *Socket) SetFieldParentInstance(value *gobject.Object) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(socketObject, recv.Native(), "parent_instance", argValue)
}

// FieldPriv returns the C field 'priv'.
func (recv *Socket) FieldPriv() *SocketPrivate {
	argValue := gi.ObjectFieldGet(socketObject, recv.Native(), "priv")
	value := SocketPrivateNewFromNative(argValue.Pointer())
	return value
}

// SetFieldPriv sets the value of the C field 'priv'.
func (recv *Socket) SetFieldPriv(value *SocketPrivate) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(socketObject, recv.Native(), "priv", argValue)
}

var socketNewFunction *gi.Function
var socketNewFunction_Once sync.Once

func socketNewFunction_Set() error {
	var err error
	socketNewFunction_Once.Do(func() {
		err = socketObject_Set()
		if err != nil {
			return
		}
		socketNewFunction, err = socketObject.InvokerNew("new")
	})
	return err
}

// SocketNew is a representation of the C type g_socket_new.
func SocketNew(family SocketFamily, type_ SocketType, protocol SocketProtocol) *Socket {
	var inArgs [3]gi.Argument
	inArgs[0].SetInt32(int32(family))
	inArgs[1].SetInt32(int32(type_))
	inArgs[2].SetInt32(int32(protocol))

	var ret gi.Argument

	err := socketNewFunction_Set()
	if err == nil {
		ret = socketNewFunction.Invoke(inArgs[:], nil)
	}

	retGo := SocketNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

var socketNewFromFdFunction *gi.Function
var socketNewFromFdFunction_Once sync.Once

func socketNewFromFdFunction_Set() error {
	var err error
	socketNewFromFdFunction_Once.Do(func() {
		err = socketObject_Set()
		if err != nil {
			return
		}
		socketNewFromFdFunction, err = socketObject.InvokerNew("new_from_fd")
	})
	return err
}

// SocketNewFromFd is a representation of the C type g_socket_new_from_fd.
func SocketNewFromFd(fd int32) *Socket {
	var inArgs [1]gi.Argument
	inArgs[0].SetInt32(fd)

	var ret gi.Argument

	err := socketNewFromFdFunction_Set()
	if err == nil {
		ret = socketNewFromFdFunction.Invoke(inArgs[:], nil)
	}

	retGo := SocketNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

var socketAcceptFunction *gi.Function
var socketAcceptFunction_Once sync.Once

func socketAcceptFunction_Set() error {
	var err error
	socketAcceptFunction_Once.Do(func() {
		err = socketObject_Set()
		if err != nil {
			return
		}
		socketAcceptFunction, err = socketObject.InvokerNew("accept")
	})
	return err
}

// Accept is a representation of the C type g_socket_accept.
func (recv *Socket) Accept(cancellable *Cancellable) *Socket {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(cancellable.Native())

	var ret gi.Argument

	err := socketAcceptFunction_Set()
	if err == nil {
		ret = socketAcceptFunction.Invoke(inArgs[:], nil)
	}

	retGo := SocketNewFromNative(ret.Pointer())

	return retGo
}

var socketBindFunction *gi.Function
var socketBindFunction_Once sync.Once

func socketBindFunction_Set() error {
	var err error
	socketBindFunction_Once.Do(func() {
		err = socketObject_Set()
		if err != nil {
			return
		}
		socketBindFunction, err = socketObject.InvokerNew("bind")
	})
	return err
}

// Bind is a representation of the C type g_socket_bind.
func (recv *Socket) Bind(address *SocketAddress, allowReuse bool) bool {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(address.Native())
	inArgs[2].SetBoolean(allowReuse)

	var ret gi.Argument

	err := socketBindFunction_Set()
	if err == nil {
		ret = socketBindFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var socketCheckConnectResultFunction *gi.Function
var socketCheckConnectResultFunction_Once sync.Once

func socketCheckConnectResultFunction_Set() error {
	var err error
	socketCheckConnectResultFunction_Once.Do(func() {
		err = socketObject_Set()
		if err != nil {
			return
		}
		socketCheckConnectResultFunction, err = socketObject.InvokerNew("check_connect_result")
	})
	return err
}

// CheckConnectResult is a representation of the C type g_socket_check_connect_result.
func (recv *Socket) CheckConnectResult() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := socketCheckConnectResultFunction_Set()
	if err == nil {
		ret = socketCheckConnectResultFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var socketCloseFunction *gi.Function
var socketCloseFunction_Once sync.Once

func socketCloseFunction_Set() error {
	var err error
	socketCloseFunction_Once.Do(func() {
		err = socketObject_Set()
		if err != nil {
			return
		}
		socketCloseFunction, err = socketObject.InvokerNew("close")
	})
	return err
}

// Close is a representation of the C type g_socket_close.
func (recv *Socket) Close() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := socketCloseFunction_Set()
	if err == nil {
		ret = socketCloseFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'g_socket_condition_check' : parameter 'condition' of type 'GLib.IOCondition' not supported

// UNSUPPORTED : C value 'g_socket_condition_timed_wait' : parameter 'condition' of type 'GLib.IOCondition' not supported

// UNSUPPORTED : C value 'g_socket_condition_wait' : parameter 'condition' of type 'GLib.IOCondition' not supported

var socketConnectFunction *gi.Function
var socketConnectFunction_Once sync.Once

func socketConnectFunction_Set() error {
	var err error
	socketConnectFunction_Once.Do(func() {
		err = socketObject_Set()
		if err != nil {
			return
		}
		socketConnectFunction, err = socketObject.InvokerNew("connect")
	})
	return err
}

// Connect is a representation of the C type g_socket_connect.
func (recv *Socket) Connect(address *SocketAddress, cancellable *Cancellable) bool {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(address.Native())
	inArgs[2].SetPointer(cancellable.Native())

	var ret gi.Argument

	err := socketConnectFunction_Set()
	if err == nil {
		ret = socketConnectFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var socketConnectionFactoryCreateConnectionFunction *gi.Function
var socketConnectionFactoryCreateConnectionFunction_Once sync.Once

func socketConnectionFactoryCreateConnectionFunction_Set() error {
	var err error
	socketConnectionFactoryCreateConnectionFunction_Once.Do(func() {
		err = socketObject_Set()
		if err != nil {
			return
		}
		socketConnectionFactoryCreateConnectionFunction, err = socketObject.InvokerNew("connection_factory_create_connection")
	})
	return err
}

// ConnectionFactoryCreateConnection is a representation of the C type g_socket_connection_factory_create_connection.
func (recv *Socket) ConnectionFactoryCreateConnection() *SocketConnection {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := socketConnectionFactoryCreateConnectionFunction_Set()
	if err == nil {
		ret = socketConnectionFactoryCreateConnectionFunction.Invoke(inArgs[:], nil)
	}

	retGo := SocketConnectionNewFromNative(ret.Pointer())

	return retGo
}

// UNSUPPORTED : C value 'g_socket_create_source' : parameter 'condition' of type 'GLib.IOCondition' not supported

var socketGetAvailableBytesFunction *gi.Function
var socketGetAvailableBytesFunction_Once sync.Once

func socketGetAvailableBytesFunction_Set() error {
	var err error
	socketGetAvailableBytesFunction_Once.Do(func() {
		err = socketObject_Set()
		if err != nil {
			return
		}
		socketGetAvailableBytesFunction, err = socketObject.InvokerNew("get_available_bytes")
	})
	return err
}

// GetAvailableBytes is a representation of the C type g_socket_get_available_bytes.
func (recv *Socket) GetAvailableBytes() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := socketGetAvailableBytesFunction_Set()
	if err == nil {
		ret = socketGetAvailableBytesFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var socketGetBlockingFunction *gi.Function
var socketGetBlockingFunction_Once sync.Once

func socketGetBlockingFunction_Set() error {
	var err error
	socketGetBlockingFunction_Once.Do(func() {
		err = socketObject_Set()
		if err != nil {
			return
		}
		socketGetBlockingFunction, err = socketObject.InvokerNew("get_blocking")
	})
	return err
}

// GetBlocking is a representation of the C type g_socket_get_blocking.
func (recv *Socket) GetBlocking() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := socketGetBlockingFunction_Set()
	if err == nil {
		ret = socketGetBlockingFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var socketGetBroadcastFunction *gi.Function
var socketGetBroadcastFunction_Once sync.Once

func socketGetBroadcastFunction_Set() error {
	var err error
	socketGetBroadcastFunction_Once.Do(func() {
		err = socketObject_Set()
		if err != nil {
			return
		}
		socketGetBroadcastFunction, err = socketObject.InvokerNew("get_broadcast")
	})
	return err
}

// GetBroadcast is a representation of the C type g_socket_get_broadcast.
func (recv *Socket) GetBroadcast() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := socketGetBroadcastFunction_Set()
	if err == nil {
		ret = socketGetBroadcastFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var socketGetCredentialsFunction *gi.Function
var socketGetCredentialsFunction_Once sync.Once

func socketGetCredentialsFunction_Set() error {
	var err error
	socketGetCredentialsFunction_Once.Do(func() {
		err = socketObject_Set()
		if err != nil {
			return
		}
		socketGetCredentialsFunction, err = socketObject.InvokerNew("get_credentials")
	})
	return err
}

// GetCredentials is a representation of the C type g_socket_get_credentials.
func (recv *Socket) GetCredentials() *Credentials {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := socketGetCredentialsFunction_Set()
	if err == nil {
		ret = socketGetCredentialsFunction.Invoke(inArgs[:], nil)
	}

	retGo := CredentialsNewFromNative(ret.Pointer())

	return retGo
}

var socketGetFamilyFunction *gi.Function
var socketGetFamilyFunction_Once sync.Once

func socketGetFamilyFunction_Set() error {
	var err error
	socketGetFamilyFunction_Once.Do(func() {
		err = socketObject_Set()
		if err != nil {
			return
		}
		socketGetFamilyFunction, err = socketObject.InvokerNew("get_family")
	})
	return err
}

// GetFamily is a representation of the C type g_socket_get_family.
func (recv *Socket) GetFamily() SocketFamily {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := socketGetFamilyFunction_Set()
	if err == nil {
		ret = socketGetFamilyFunction.Invoke(inArgs[:], nil)
	}

	retGo := SocketFamily(ret.Int32())

	return retGo
}

var socketGetFdFunction *gi.Function
var socketGetFdFunction_Once sync.Once

func socketGetFdFunction_Set() error {
	var err error
	socketGetFdFunction_Once.Do(func() {
		err = socketObject_Set()
		if err != nil {
			return
		}
		socketGetFdFunction, err = socketObject.InvokerNew("get_fd")
	})
	return err
}

// GetFd is a representation of the C type g_socket_get_fd.
func (recv *Socket) GetFd() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := socketGetFdFunction_Set()
	if err == nil {
		ret = socketGetFdFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var socketGetKeepaliveFunction *gi.Function
var socketGetKeepaliveFunction_Once sync.Once

func socketGetKeepaliveFunction_Set() error {
	var err error
	socketGetKeepaliveFunction_Once.Do(func() {
		err = socketObject_Set()
		if err != nil {
			return
		}
		socketGetKeepaliveFunction, err = socketObject.InvokerNew("get_keepalive")
	})
	return err
}

// GetKeepalive is a representation of the C type g_socket_get_keepalive.
func (recv *Socket) GetKeepalive() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := socketGetKeepaliveFunction_Set()
	if err == nil {
		ret = socketGetKeepaliveFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var socketGetListenBacklogFunction *gi.Function
var socketGetListenBacklogFunction_Once sync.Once

func socketGetListenBacklogFunction_Set() error {
	var err error
	socketGetListenBacklogFunction_Once.Do(func() {
		err = socketObject_Set()
		if err != nil {
			return
		}
		socketGetListenBacklogFunction, err = socketObject.InvokerNew("get_listen_backlog")
	})
	return err
}

// GetListenBacklog is a representation of the C type g_socket_get_listen_backlog.
func (recv *Socket) GetListenBacklog() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := socketGetListenBacklogFunction_Set()
	if err == nil {
		ret = socketGetListenBacklogFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var socketGetLocalAddressFunction *gi.Function
var socketGetLocalAddressFunction_Once sync.Once

func socketGetLocalAddressFunction_Set() error {
	var err error
	socketGetLocalAddressFunction_Once.Do(func() {
		err = socketObject_Set()
		if err != nil {
			return
		}
		socketGetLocalAddressFunction, err = socketObject.InvokerNew("get_local_address")
	})
	return err
}

// GetLocalAddress is a representation of the C type g_socket_get_local_address.
func (recv *Socket) GetLocalAddress() *SocketAddress {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := socketGetLocalAddressFunction_Set()
	if err == nil {
		ret = socketGetLocalAddressFunction.Invoke(inArgs[:], nil)
	}

	retGo := SocketAddressNewFromNative(ret.Pointer())

	return retGo
}

var socketGetMulticastLoopbackFunction *gi.Function
var socketGetMulticastLoopbackFunction_Once sync.Once

func socketGetMulticastLoopbackFunction_Set() error {
	var err error
	socketGetMulticastLoopbackFunction_Once.Do(func() {
		err = socketObject_Set()
		if err != nil {
			return
		}
		socketGetMulticastLoopbackFunction, err = socketObject.InvokerNew("get_multicast_loopback")
	})
	return err
}

// GetMulticastLoopback is a representation of the C type g_socket_get_multicast_loopback.
func (recv *Socket) GetMulticastLoopback() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := socketGetMulticastLoopbackFunction_Set()
	if err == nil {
		ret = socketGetMulticastLoopbackFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var socketGetMulticastTtlFunction *gi.Function
var socketGetMulticastTtlFunction_Once sync.Once

func socketGetMulticastTtlFunction_Set() error {
	var err error
	socketGetMulticastTtlFunction_Once.Do(func() {
		err = socketObject_Set()
		if err != nil {
			return
		}
		socketGetMulticastTtlFunction, err = socketObject.InvokerNew("get_multicast_ttl")
	})
	return err
}

// GetMulticastTtl is a representation of the C type g_socket_get_multicast_ttl.
func (recv *Socket) GetMulticastTtl() uint32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := socketGetMulticastTtlFunction_Set()
	if err == nil {
		ret = socketGetMulticastTtlFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo
}

var socketGetOptionFunction *gi.Function
var socketGetOptionFunction_Once sync.Once

func socketGetOptionFunction_Set() error {
	var err error
	socketGetOptionFunction_Once.Do(func() {
		err = socketObject_Set()
		if err != nil {
			return
		}
		socketGetOptionFunction, err = socketObject.InvokerNew("get_option")
	})
	return err
}

// GetOption is a representation of the C type g_socket_get_option.
func (recv *Socket) GetOption(level int32, optname int32) (bool, int32) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(level)
	inArgs[2].SetInt32(optname)

	var outArgs [1]gi.Argument
	var ret gi.Argument

	err := socketGetOptionFunction_Set()
	if err == nil {
		ret = socketGetOptionFunction.Invoke(inArgs[:], outArgs[:])
	}

	retGo := ret.Boolean()
	out0 := outArgs[0].Int32()

	return retGo, out0
}

var socketGetProtocolFunction *gi.Function
var socketGetProtocolFunction_Once sync.Once

func socketGetProtocolFunction_Set() error {
	var err error
	socketGetProtocolFunction_Once.Do(func() {
		err = socketObject_Set()
		if err != nil {
			return
		}
		socketGetProtocolFunction, err = socketObject.InvokerNew("get_protocol")
	})
	return err
}

// GetProtocol is a representation of the C type g_socket_get_protocol.
func (recv *Socket) GetProtocol() SocketProtocol {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := socketGetProtocolFunction_Set()
	if err == nil {
		ret = socketGetProtocolFunction.Invoke(inArgs[:], nil)
	}

	retGo := SocketProtocol(ret.Int32())

	return retGo
}

var socketGetRemoteAddressFunction *gi.Function
var socketGetRemoteAddressFunction_Once sync.Once

func socketGetRemoteAddressFunction_Set() error {
	var err error
	socketGetRemoteAddressFunction_Once.Do(func() {
		err = socketObject_Set()
		if err != nil {
			return
		}
		socketGetRemoteAddressFunction, err = socketObject.InvokerNew("get_remote_address")
	})
	return err
}

// GetRemoteAddress is a representation of the C type g_socket_get_remote_address.
func (recv *Socket) GetRemoteAddress() *SocketAddress {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := socketGetRemoteAddressFunction_Set()
	if err == nil {
		ret = socketGetRemoteAddressFunction.Invoke(inArgs[:], nil)
	}

	retGo := SocketAddressNewFromNative(ret.Pointer())

	return retGo
}

var socketGetSocketTypeFunction *gi.Function
var socketGetSocketTypeFunction_Once sync.Once

func socketGetSocketTypeFunction_Set() error {
	var err error
	socketGetSocketTypeFunction_Once.Do(func() {
		err = socketObject_Set()
		if err != nil {
			return
		}
		socketGetSocketTypeFunction, err = socketObject.InvokerNew("get_socket_type")
	})
	return err
}

// GetSocketType is a representation of the C type g_socket_get_socket_type.
func (recv *Socket) GetSocketType() SocketType {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := socketGetSocketTypeFunction_Set()
	if err == nil {
		ret = socketGetSocketTypeFunction.Invoke(inArgs[:], nil)
	}

	retGo := SocketType(ret.Int32())

	return retGo
}

var socketGetTimeoutFunction *gi.Function
var socketGetTimeoutFunction_Once sync.Once

func socketGetTimeoutFunction_Set() error {
	var err error
	socketGetTimeoutFunction_Once.Do(func() {
		err = socketObject_Set()
		if err != nil {
			return
		}
		socketGetTimeoutFunction, err = socketObject.InvokerNew("get_timeout")
	})
	return err
}

// GetTimeout is a representation of the C type g_socket_get_timeout.
func (recv *Socket) GetTimeout() uint32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := socketGetTimeoutFunction_Set()
	if err == nil {
		ret = socketGetTimeoutFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo
}

var socketGetTtlFunction *gi.Function
var socketGetTtlFunction_Once sync.Once

func socketGetTtlFunction_Set() error {
	var err error
	socketGetTtlFunction_Once.Do(func() {
		err = socketObject_Set()
		if err != nil {
			return
		}
		socketGetTtlFunction, err = socketObject.InvokerNew("get_ttl")
	})
	return err
}

// GetTtl is a representation of the C type g_socket_get_ttl.
func (recv *Socket) GetTtl() uint32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := socketGetTtlFunction_Set()
	if err == nil {
		ret = socketGetTtlFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo
}

var socketIsClosedFunction *gi.Function
var socketIsClosedFunction_Once sync.Once

func socketIsClosedFunction_Set() error {
	var err error
	socketIsClosedFunction_Once.Do(func() {
		err = socketObject_Set()
		if err != nil {
			return
		}
		socketIsClosedFunction, err = socketObject.InvokerNew("is_closed")
	})
	return err
}

// IsClosed is a representation of the C type g_socket_is_closed.
func (recv *Socket) IsClosed() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := socketIsClosedFunction_Set()
	if err == nil {
		ret = socketIsClosedFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var socketIsConnectedFunction *gi.Function
var socketIsConnectedFunction_Once sync.Once

func socketIsConnectedFunction_Set() error {
	var err error
	socketIsConnectedFunction_Once.Do(func() {
		err = socketObject_Set()
		if err != nil {
			return
		}
		socketIsConnectedFunction, err = socketObject.InvokerNew("is_connected")
	})
	return err
}

// IsConnected is a representation of the C type g_socket_is_connected.
func (recv *Socket) IsConnected() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := socketIsConnectedFunction_Set()
	if err == nil {
		ret = socketIsConnectedFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var socketJoinMulticastGroupFunction *gi.Function
var socketJoinMulticastGroupFunction_Once sync.Once

func socketJoinMulticastGroupFunction_Set() error {
	var err error
	socketJoinMulticastGroupFunction_Once.Do(func() {
		err = socketObject_Set()
		if err != nil {
			return
		}
		socketJoinMulticastGroupFunction, err = socketObject.InvokerNew("join_multicast_group")
	})
	return err
}

// JoinMulticastGroup is a representation of the C type g_socket_join_multicast_group.
func (recv *Socket) JoinMulticastGroup(group *InetAddress, sourceSpecific bool, iface string) bool {
	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(group.Native())
	inArgs[2].SetBoolean(sourceSpecific)
	inArgs[3].SetString(iface)

	var ret gi.Argument

	err := socketJoinMulticastGroupFunction_Set()
	if err == nil {
		ret = socketJoinMulticastGroupFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var socketJoinMulticastGroupSsmFunction *gi.Function
var socketJoinMulticastGroupSsmFunction_Once sync.Once

func socketJoinMulticastGroupSsmFunction_Set() error {
	var err error
	socketJoinMulticastGroupSsmFunction_Once.Do(func() {
		err = socketObject_Set()
		if err != nil {
			return
		}
		socketJoinMulticastGroupSsmFunction, err = socketObject.InvokerNew("join_multicast_group_ssm")
	})
	return err
}

// JoinMulticastGroupSsm is a representation of the C type g_socket_join_multicast_group_ssm.
func (recv *Socket) JoinMulticastGroupSsm(group *InetAddress, sourceSpecific *InetAddress, iface string) bool {
	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(group.Native())
	inArgs[2].SetPointer(sourceSpecific.Native())
	inArgs[3].SetString(iface)

	var ret gi.Argument

	err := socketJoinMulticastGroupSsmFunction_Set()
	if err == nil {
		ret = socketJoinMulticastGroupSsmFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var socketLeaveMulticastGroupFunction *gi.Function
var socketLeaveMulticastGroupFunction_Once sync.Once

func socketLeaveMulticastGroupFunction_Set() error {
	var err error
	socketLeaveMulticastGroupFunction_Once.Do(func() {
		err = socketObject_Set()
		if err != nil {
			return
		}
		socketLeaveMulticastGroupFunction, err = socketObject.InvokerNew("leave_multicast_group")
	})
	return err
}

// LeaveMulticastGroup is a representation of the C type g_socket_leave_multicast_group.
func (recv *Socket) LeaveMulticastGroup(group *InetAddress, sourceSpecific bool, iface string) bool {
	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(group.Native())
	inArgs[2].SetBoolean(sourceSpecific)
	inArgs[3].SetString(iface)

	var ret gi.Argument

	err := socketLeaveMulticastGroupFunction_Set()
	if err == nil {
		ret = socketLeaveMulticastGroupFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var socketLeaveMulticastGroupSsmFunction *gi.Function
var socketLeaveMulticastGroupSsmFunction_Once sync.Once

func socketLeaveMulticastGroupSsmFunction_Set() error {
	var err error
	socketLeaveMulticastGroupSsmFunction_Once.Do(func() {
		err = socketObject_Set()
		if err != nil {
			return
		}
		socketLeaveMulticastGroupSsmFunction, err = socketObject.InvokerNew("leave_multicast_group_ssm")
	})
	return err
}

// LeaveMulticastGroupSsm is a representation of the C type g_socket_leave_multicast_group_ssm.
func (recv *Socket) LeaveMulticastGroupSsm(group *InetAddress, sourceSpecific *InetAddress, iface string) bool {
	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(group.Native())
	inArgs[2].SetPointer(sourceSpecific.Native())
	inArgs[3].SetString(iface)

	var ret gi.Argument

	err := socketLeaveMulticastGroupSsmFunction_Set()
	if err == nil {
		ret = socketLeaveMulticastGroupSsmFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var socketListenFunction *gi.Function
var socketListenFunction_Once sync.Once

func socketListenFunction_Set() error {
	var err error
	socketListenFunction_Once.Do(func() {
		err = socketObject_Set()
		if err != nil {
			return
		}
		socketListenFunction, err = socketObject.InvokerNew("listen")
	})
	return err
}

// Listen is a representation of the C type g_socket_listen.
func (recv *Socket) Listen() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := socketListenFunction_Set()
	if err == nil {
		ret = socketListenFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'g_socket_receive' : parameter 'buffer' of type 'nil' not supported

// UNSUPPORTED : C value 'g_socket_receive_from' : parameter 'buffer' of type 'nil' not supported

// UNSUPPORTED : C value 'g_socket_receive_message' : parameter 'vectors' of type 'nil' not supported

// UNSUPPORTED : C value 'g_socket_receive_messages' : parameter 'messages' of type 'nil' not supported

// UNSUPPORTED : C value 'g_socket_receive_with_blocking' : parameter 'buffer' of type 'nil' not supported

// UNSUPPORTED : C value 'g_socket_send' : parameter 'buffer' of type 'nil' not supported

// UNSUPPORTED : C value 'g_socket_send_message' : parameter 'vectors' of type 'nil' not supported

// UNSUPPORTED : C value 'g_socket_send_message_with_timeout' : parameter 'vectors' of type 'nil' not supported

// UNSUPPORTED : C value 'g_socket_send_messages' : parameter 'messages' of type 'nil' not supported

// UNSUPPORTED : C value 'g_socket_send_to' : parameter 'buffer' of type 'nil' not supported

// UNSUPPORTED : C value 'g_socket_send_with_blocking' : parameter 'buffer' of type 'nil' not supported

var socketSetBlockingFunction *gi.Function
var socketSetBlockingFunction_Once sync.Once

func socketSetBlockingFunction_Set() error {
	var err error
	socketSetBlockingFunction_Once.Do(func() {
		err = socketObject_Set()
		if err != nil {
			return
		}
		socketSetBlockingFunction, err = socketObject.InvokerNew("set_blocking")
	})
	return err
}

// SetBlocking is a representation of the C type g_socket_set_blocking.
func (recv *Socket) SetBlocking(blocking bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetBoolean(blocking)

	err := socketSetBlockingFunction_Set()
	if err == nil {
		socketSetBlockingFunction.Invoke(inArgs[:], nil)
	}

	return
}

var socketSetBroadcastFunction *gi.Function
var socketSetBroadcastFunction_Once sync.Once

func socketSetBroadcastFunction_Set() error {
	var err error
	socketSetBroadcastFunction_Once.Do(func() {
		err = socketObject_Set()
		if err != nil {
			return
		}
		socketSetBroadcastFunction, err = socketObject.InvokerNew("set_broadcast")
	})
	return err
}

// SetBroadcast is a representation of the C type g_socket_set_broadcast.
func (recv *Socket) SetBroadcast(broadcast bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetBoolean(broadcast)

	err := socketSetBroadcastFunction_Set()
	if err == nil {
		socketSetBroadcastFunction.Invoke(inArgs[:], nil)
	}

	return
}

var socketSetKeepaliveFunction *gi.Function
var socketSetKeepaliveFunction_Once sync.Once

func socketSetKeepaliveFunction_Set() error {
	var err error
	socketSetKeepaliveFunction_Once.Do(func() {
		err = socketObject_Set()
		if err != nil {
			return
		}
		socketSetKeepaliveFunction, err = socketObject.InvokerNew("set_keepalive")
	})
	return err
}

// SetKeepalive is a representation of the C type g_socket_set_keepalive.
func (recv *Socket) SetKeepalive(keepalive bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetBoolean(keepalive)

	err := socketSetKeepaliveFunction_Set()
	if err == nil {
		socketSetKeepaliveFunction.Invoke(inArgs[:], nil)
	}

	return
}

var socketSetListenBacklogFunction *gi.Function
var socketSetListenBacklogFunction_Once sync.Once

func socketSetListenBacklogFunction_Set() error {
	var err error
	socketSetListenBacklogFunction_Once.Do(func() {
		err = socketObject_Set()
		if err != nil {
			return
		}
		socketSetListenBacklogFunction, err = socketObject.InvokerNew("set_listen_backlog")
	})
	return err
}

// SetListenBacklog is a representation of the C type g_socket_set_listen_backlog.
func (recv *Socket) SetListenBacklog(backlog int32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(backlog)

	err := socketSetListenBacklogFunction_Set()
	if err == nil {
		socketSetListenBacklogFunction.Invoke(inArgs[:], nil)
	}

	return
}

var socketSetMulticastLoopbackFunction *gi.Function
var socketSetMulticastLoopbackFunction_Once sync.Once

func socketSetMulticastLoopbackFunction_Set() error {
	var err error
	socketSetMulticastLoopbackFunction_Once.Do(func() {
		err = socketObject_Set()
		if err != nil {
			return
		}
		socketSetMulticastLoopbackFunction, err = socketObject.InvokerNew("set_multicast_loopback")
	})
	return err
}

// SetMulticastLoopback is a representation of the C type g_socket_set_multicast_loopback.
func (recv *Socket) SetMulticastLoopback(loopback bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetBoolean(loopback)

	err := socketSetMulticastLoopbackFunction_Set()
	if err == nil {
		socketSetMulticastLoopbackFunction.Invoke(inArgs[:], nil)
	}

	return
}

var socketSetMulticastTtlFunction *gi.Function
var socketSetMulticastTtlFunction_Once sync.Once

func socketSetMulticastTtlFunction_Set() error {
	var err error
	socketSetMulticastTtlFunction_Once.Do(func() {
		err = socketObject_Set()
		if err != nil {
			return
		}
		socketSetMulticastTtlFunction, err = socketObject.InvokerNew("set_multicast_ttl")
	})
	return err
}

// SetMulticastTtl is a representation of the C type g_socket_set_multicast_ttl.
func (recv *Socket) SetMulticastTtl(ttl uint32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetUint32(ttl)

	err := socketSetMulticastTtlFunction_Set()
	if err == nil {
		socketSetMulticastTtlFunction.Invoke(inArgs[:], nil)
	}

	return
}

var socketSetOptionFunction *gi.Function
var socketSetOptionFunction_Once sync.Once

func socketSetOptionFunction_Set() error {
	var err error
	socketSetOptionFunction_Once.Do(func() {
		err = socketObject_Set()
		if err != nil {
			return
		}
		socketSetOptionFunction, err = socketObject.InvokerNew("set_option")
	})
	return err
}

// SetOption is a representation of the C type g_socket_set_option.
func (recv *Socket) SetOption(level int32, optname int32, value int32) bool {
	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(level)
	inArgs[2].SetInt32(optname)
	inArgs[3].SetInt32(value)

	var ret gi.Argument

	err := socketSetOptionFunction_Set()
	if err == nil {
		ret = socketSetOptionFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var socketSetTimeoutFunction *gi.Function
var socketSetTimeoutFunction_Once sync.Once

func socketSetTimeoutFunction_Set() error {
	var err error
	socketSetTimeoutFunction_Once.Do(func() {
		err = socketObject_Set()
		if err != nil {
			return
		}
		socketSetTimeoutFunction, err = socketObject.InvokerNew("set_timeout")
	})
	return err
}

// SetTimeout is a representation of the C type g_socket_set_timeout.
func (recv *Socket) SetTimeout(timeout uint32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetUint32(timeout)

	err := socketSetTimeoutFunction_Set()
	if err == nil {
		socketSetTimeoutFunction.Invoke(inArgs[:], nil)
	}

	return
}

var socketSetTtlFunction *gi.Function
var socketSetTtlFunction_Once sync.Once

func socketSetTtlFunction_Set() error {
	var err error
	socketSetTtlFunction_Once.Do(func() {
		err = socketObject_Set()
		if err != nil {
			return
		}
		socketSetTtlFunction, err = socketObject.InvokerNew("set_ttl")
	})
	return err
}

// SetTtl is a representation of the C type g_socket_set_ttl.
func (recv *Socket) SetTtl(ttl uint32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetUint32(ttl)

	err := socketSetTtlFunction_Set()
	if err == nil {
		socketSetTtlFunction.Invoke(inArgs[:], nil)
	}

	return
}

var socketShutdownFunction *gi.Function
var socketShutdownFunction_Once sync.Once

func socketShutdownFunction_Set() error {
	var err error
	socketShutdownFunction_Once.Do(func() {
		err = socketObject_Set()
		if err != nil {
			return
		}
		socketShutdownFunction, err = socketObject.InvokerNew("shutdown")
	})
	return err
}

// Shutdown is a representation of the C type g_socket_shutdown.
func (recv *Socket) Shutdown(shutdownRead bool, shutdownWrite bool) bool {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetBoolean(shutdownRead)
	inArgs[2].SetBoolean(shutdownWrite)

	var ret gi.Argument

	err := socketShutdownFunction_Set()
	if err == nil {
		ret = socketShutdownFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var socketSpeaksIpv4Function *gi.Function
var socketSpeaksIpv4Function_Once sync.Once

func socketSpeaksIpv4Function_Set() error {
	var err error
	socketSpeaksIpv4Function_Once.Do(func() {
		err = socketObject_Set()
		if err != nil {
			return
		}
		socketSpeaksIpv4Function, err = socketObject.InvokerNew("speaks_ipv4")
	})
	return err
}

// SpeaksIpv4 is a representation of the C type g_socket_speaks_ipv4.
func (recv *Socket) SpeaksIpv4() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := socketSpeaksIpv4Function_Set()
	if err == nil {
		ret = socketSpeaksIpv4Function.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var socketAddressObject *gi.Object
var socketAddressObject_Once sync.Once

func socketAddressObject_Set() error {
	var err error
	socketAddressObject_Once.Do(func() {
		socketAddressObject, err = gi.ObjectNew("Gio", "SocketAddress")
	})
	return err
}

type SocketAddress struct {
	native unsafe.Pointer
}

func SocketAddressNewFromNative(native unsafe.Pointer) *SocketAddress {
	instance := &SocketAddress{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *SocketAddress) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToSocketAddress down casts any arbitrary Object to SocketAddress.
Exercise care, as this is a potentially dangerous function
if the Object is not a SocketAddress.
*/
func CastToSocketAddress(object *gobject.Object) *SocketAddress {
	return SocketAddressNewFromNative(object.Native())
}

// Equals compares this SocketAddress with another SocketAddress, and returns true if they represent the same GObject.
func (recv *SocketAddress) Equals(other *SocketAddress) bool {
	return other.Native() == recv.Native()
}

func (recv *SocketAddress) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *SocketAddress) FieldParentInstance() *gobject.Object {
	argValue := gi.ObjectFieldGet(socketAddressObject, recv.Native(), "parent_instance")
	value := gobject.ObjectNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *SocketAddress) SetFieldParentInstance(value *gobject.Object) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(socketAddressObject, recv.Native(), "parent_instance", argValue)
}

// UNSUPPORTED : C value 'g_socket_address_new_from_native' : parameter 'native' of type 'gpointer' not supported

var socketAddressGetFamilyFunction *gi.Function
var socketAddressGetFamilyFunction_Once sync.Once

func socketAddressGetFamilyFunction_Set() error {
	var err error
	socketAddressGetFamilyFunction_Once.Do(func() {
		err = socketAddressObject_Set()
		if err != nil {
			return
		}
		socketAddressGetFamilyFunction, err = socketAddressObject.InvokerNew("get_family")
	})
	return err
}

// GetFamily is a representation of the C type g_socket_address_get_family.
func (recv *SocketAddress) GetFamily() SocketFamily {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := socketAddressGetFamilyFunction_Set()
	if err == nil {
		ret = socketAddressGetFamilyFunction.Invoke(inArgs[:], nil)
	}

	retGo := SocketFamily(ret.Int32())

	return retGo
}

var socketAddressGetNativeSizeFunction *gi.Function
var socketAddressGetNativeSizeFunction_Once sync.Once

func socketAddressGetNativeSizeFunction_Set() error {
	var err error
	socketAddressGetNativeSizeFunction_Once.Do(func() {
		err = socketAddressObject_Set()
		if err != nil {
			return
		}
		socketAddressGetNativeSizeFunction, err = socketAddressObject.InvokerNew("get_native_size")
	})
	return err
}

// GetNativeSize is a representation of the C type g_socket_address_get_native_size.
func (recv *SocketAddress) GetNativeSize() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := socketAddressGetNativeSizeFunction_Set()
	if err == nil {
		ret = socketAddressGetNativeSizeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

// UNSUPPORTED : C value 'g_socket_address_to_native' : parameter 'dest' of type 'gpointer' not supported

var socketAddressEnumeratorObject *gi.Object
var socketAddressEnumeratorObject_Once sync.Once

func socketAddressEnumeratorObject_Set() error {
	var err error
	socketAddressEnumeratorObject_Once.Do(func() {
		socketAddressEnumeratorObject, err = gi.ObjectNew("Gio", "SocketAddressEnumerator")
	})
	return err
}

type SocketAddressEnumerator struct {
	native unsafe.Pointer
}

func SocketAddressEnumeratorNewFromNative(native unsafe.Pointer) *SocketAddressEnumerator {
	instance := &SocketAddressEnumerator{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *SocketAddressEnumerator) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToSocketAddressEnumerator down casts any arbitrary Object to SocketAddressEnumerator.
Exercise care, as this is a potentially dangerous function
if the Object is not a SocketAddressEnumerator.
*/
func CastToSocketAddressEnumerator(object *gobject.Object) *SocketAddressEnumerator {
	return SocketAddressEnumeratorNewFromNative(object.Native())
}

// Equals compares this SocketAddressEnumerator with another SocketAddressEnumerator, and returns true if they represent the same GObject.
func (recv *SocketAddressEnumerator) Equals(other *SocketAddressEnumerator) bool {
	return other.Native() == recv.Native()
}

func (recv *SocketAddressEnumerator) Native() unsafe.Pointer {
	return recv.native
}

var socketAddressEnumeratorNextFunction *gi.Function
var socketAddressEnumeratorNextFunction_Once sync.Once

func socketAddressEnumeratorNextFunction_Set() error {
	var err error
	socketAddressEnumeratorNextFunction_Once.Do(func() {
		err = socketAddressEnumeratorObject_Set()
		if err != nil {
			return
		}
		socketAddressEnumeratorNextFunction, err = socketAddressEnumeratorObject.InvokerNew("next")
	})
	return err
}

// Next is a representation of the C type g_socket_address_enumerator_next.
func (recv *SocketAddressEnumerator) Next(cancellable *Cancellable) *SocketAddress {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(cancellable.Native())

	var ret gi.Argument

	err := socketAddressEnumeratorNextFunction_Set()
	if err == nil {
		ret = socketAddressEnumeratorNextFunction.Invoke(inArgs[:], nil)
	}

	retGo := SocketAddressNewFromNative(ret.Pointer())

	return retGo
}

// UNSUPPORTED : C value 'g_socket_address_enumerator_next_async' : parameter 'callback' of type 'AsyncReadyCallback' not supported

// UNSUPPORTED : C value 'g_socket_address_enumerator_next_finish' : parameter 'result' of type 'AsyncResult' not supported

var socketClientObject *gi.Object
var socketClientObject_Once sync.Once

func socketClientObject_Set() error {
	var err error
	socketClientObject_Once.Do(func() {
		socketClientObject, err = gi.ObjectNew("Gio", "SocketClient")
	})
	return err
}

type SocketClient struct {
	native unsafe.Pointer
}

func SocketClientNewFromNative(native unsafe.Pointer) *SocketClient {
	instance := &SocketClient{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *SocketClient) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToSocketClient down casts any arbitrary Object to SocketClient.
Exercise care, as this is a potentially dangerous function
if the Object is not a SocketClient.
*/
func CastToSocketClient(object *gobject.Object) *SocketClient {
	return SocketClientNewFromNative(object.Native())
}

// Equals compares this SocketClient with another SocketClient, and returns true if they represent the same GObject.
func (recv *SocketClient) Equals(other *SocketClient) bool {
	return other.Native() == recv.Native()
}

func (recv *SocketClient) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *SocketClient) FieldParentInstance() *gobject.Object {
	argValue := gi.ObjectFieldGet(socketClientObject, recv.Native(), "parent_instance")
	value := gobject.ObjectNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *SocketClient) SetFieldParentInstance(value *gobject.Object) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(socketClientObject, recv.Native(), "parent_instance", argValue)
}

// FieldPriv returns the C field 'priv'.
func (recv *SocketClient) FieldPriv() *SocketClientPrivate {
	argValue := gi.ObjectFieldGet(socketClientObject, recv.Native(), "priv")
	value := SocketClientPrivateNewFromNative(argValue.Pointer())
	return value
}

// SetFieldPriv sets the value of the C field 'priv'.
func (recv *SocketClient) SetFieldPriv(value *SocketClientPrivate) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(socketClientObject, recv.Native(), "priv", argValue)
}

var socketClientNewFunction *gi.Function
var socketClientNewFunction_Once sync.Once

func socketClientNewFunction_Set() error {
	var err error
	socketClientNewFunction_Once.Do(func() {
		err = socketClientObject_Set()
		if err != nil {
			return
		}
		socketClientNewFunction, err = socketClientObject.InvokerNew("new")
	})
	return err
}

// SocketClientNew is a representation of the C type g_socket_client_new.
func SocketClientNew() *SocketClient {

	var ret gi.Argument

	err := socketClientNewFunction_Set()
	if err == nil {
		ret = socketClientNewFunction.Invoke(nil, nil)
	}

	retGo := SocketClientNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

var socketClientAddApplicationProxyFunction *gi.Function
var socketClientAddApplicationProxyFunction_Once sync.Once

func socketClientAddApplicationProxyFunction_Set() error {
	var err error
	socketClientAddApplicationProxyFunction_Once.Do(func() {
		err = socketClientObject_Set()
		if err != nil {
			return
		}
		socketClientAddApplicationProxyFunction, err = socketClientObject.InvokerNew("add_application_proxy")
	})
	return err
}

// AddApplicationProxy is a representation of the C type g_socket_client_add_application_proxy.
func (recv *SocketClient) AddApplicationProxy(protocol string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(protocol)

	err := socketClientAddApplicationProxyFunction_Set()
	if err == nil {
		socketClientAddApplicationProxyFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'g_socket_client_connect' : parameter 'connectable' of type 'SocketConnectable' not supported

// UNSUPPORTED : C value 'g_socket_client_connect_async' : parameter 'connectable' of type 'SocketConnectable' not supported

// UNSUPPORTED : C value 'g_socket_client_connect_finish' : parameter 'result' of type 'AsyncResult' not supported

var socketClientConnectToHostFunction *gi.Function
var socketClientConnectToHostFunction_Once sync.Once

func socketClientConnectToHostFunction_Set() error {
	var err error
	socketClientConnectToHostFunction_Once.Do(func() {
		err = socketClientObject_Set()
		if err != nil {
			return
		}
		socketClientConnectToHostFunction, err = socketClientObject.InvokerNew("connect_to_host")
	})
	return err
}

// ConnectToHost is a representation of the C type g_socket_client_connect_to_host.
func (recv *SocketClient) ConnectToHost(hostAndPort string, defaultPort uint16, cancellable *Cancellable) *SocketConnection {
	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(hostAndPort)
	inArgs[2].SetUint16(defaultPort)
	inArgs[3].SetPointer(cancellable.Native())

	var ret gi.Argument

	err := socketClientConnectToHostFunction_Set()
	if err == nil {
		ret = socketClientConnectToHostFunction.Invoke(inArgs[:], nil)
	}

	retGo := SocketConnectionNewFromNative(ret.Pointer())

	return retGo
}

// UNSUPPORTED : C value 'g_socket_client_connect_to_host_async' : parameter 'callback' of type 'AsyncReadyCallback' not supported

// UNSUPPORTED : C value 'g_socket_client_connect_to_host_finish' : parameter 'result' of type 'AsyncResult' not supported

var socketClientConnectToServiceFunction *gi.Function
var socketClientConnectToServiceFunction_Once sync.Once

func socketClientConnectToServiceFunction_Set() error {
	var err error
	socketClientConnectToServiceFunction_Once.Do(func() {
		err = socketClientObject_Set()
		if err != nil {
			return
		}
		socketClientConnectToServiceFunction, err = socketClientObject.InvokerNew("connect_to_service")
	})
	return err
}

// ConnectToService is a representation of the C type g_socket_client_connect_to_service.
func (recv *SocketClient) ConnectToService(domain string, service string, cancellable *Cancellable) *SocketConnection {
	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(domain)
	inArgs[2].SetString(service)
	inArgs[3].SetPointer(cancellable.Native())

	var ret gi.Argument

	err := socketClientConnectToServiceFunction_Set()
	if err == nil {
		ret = socketClientConnectToServiceFunction.Invoke(inArgs[:], nil)
	}

	retGo := SocketConnectionNewFromNative(ret.Pointer())

	return retGo
}

// UNSUPPORTED : C value 'g_socket_client_connect_to_service_async' : parameter 'callback' of type 'AsyncReadyCallback' not supported

// UNSUPPORTED : C value 'g_socket_client_connect_to_service_finish' : parameter 'result' of type 'AsyncResult' not supported

var socketClientConnectToUriFunction *gi.Function
var socketClientConnectToUriFunction_Once sync.Once

func socketClientConnectToUriFunction_Set() error {
	var err error
	socketClientConnectToUriFunction_Once.Do(func() {
		err = socketClientObject_Set()
		if err != nil {
			return
		}
		socketClientConnectToUriFunction, err = socketClientObject.InvokerNew("connect_to_uri")
	})
	return err
}

// ConnectToUri is a representation of the C type g_socket_client_connect_to_uri.
func (recv *SocketClient) ConnectToUri(uri string, defaultPort uint16, cancellable *Cancellable) *SocketConnection {
	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(uri)
	inArgs[2].SetUint16(defaultPort)
	inArgs[3].SetPointer(cancellable.Native())

	var ret gi.Argument

	err := socketClientConnectToUriFunction_Set()
	if err == nil {
		ret = socketClientConnectToUriFunction.Invoke(inArgs[:], nil)
	}

	retGo := SocketConnectionNewFromNative(ret.Pointer())

	return retGo
}

// UNSUPPORTED : C value 'g_socket_client_connect_to_uri_async' : parameter 'callback' of type 'AsyncReadyCallback' not supported

// UNSUPPORTED : C value 'g_socket_client_connect_to_uri_finish' : parameter 'result' of type 'AsyncResult' not supported

var socketClientGetEnableProxyFunction *gi.Function
var socketClientGetEnableProxyFunction_Once sync.Once

func socketClientGetEnableProxyFunction_Set() error {
	var err error
	socketClientGetEnableProxyFunction_Once.Do(func() {
		err = socketClientObject_Set()
		if err != nil {
			return
		}
		socketClientGetEnableProxyFunction, err = socketClientObject.InvokerNew("get_enable_proxy")
	})
	return err
}

// GetEnableProxy is a representation of the C type g_socket_client_get_enable_proxy.
func (recv *SocketClient) GetEnableProxy() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := socketClientGetEnableProxyFunction_Set()
	if err == nil {
		ret = socketClientGetEnableProxyFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var socketClientGetFamilyFunction *gi.Function
var socketClientGetFamilyFunction_Once sync.Once

func socketClientGetFamilyFunction_Set() error {
	var err error
	socketClientGetFamilyFunction_Once.Do(func() {
		err = socketClientObject_Set()
		if err != nil {
			return
		}
		socketClientGetFamilyFunction, err = socketClientObject.InvokerNew("get_family")
	})
	return err
}

// GetFamily is a representation of the C type g_socket_client_get_family.
func (recv *SocketClient) GetFamily() SocketFamily {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := socketClientGetFamilyFunction_Set()
	if err == nil {
		ret = socketClientGetFamilyFunction.Invoke(inArgs[:], nil)
	}

	retGo := SocketFamily(ret.Int32())

	return retGo
}

var socketClientGetLocalAddressFunction *gi.Function
var socketClientGetLocalAddressFunction_Once sync.Once

func socketClientGetLocalAddressFunction_Set() error {
	var err error
	socketClientGetLocalAddressFunction_Once.Do(func() {
		err = socketClientObject_Set()
		if err != nil {
			return
		}
		socketClientGetLocalAddressFunction, err = socketClientObject.InvokerNew("get_local_address")
	})
	return err
}

// GetLocalAddress is a representation of the C type g_socket_client_get_local_address.
func (recv *SocketClient) GetLocalAddress() *SocketAddress {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := socketClientGetLocalAddressFunction_Set()
	if err == nil {
		ret = socketClientGetLocalAddressFunction.Invoke(inArgs[:], nil)
	}

	retGo := SocketAddressNewFromNative(ret.Pointer())

	return retGo
}

var socketClientGetProtocolFunction *gi.Function
var socketClientGetProtocolFunction_Once sync.Once

func socketClientGetProtocolFunction_Set() error {
	var err error
	socketClientGetProtocolFunction_Once.Do(func() {
		err = socketClientObject_Set()
		if err != nil {
			return
		}
		socketClientGetProtocolFunction, err = socketClientObject.InvokerNew("get_protocol")
	})
	return err
}

// GetProtocol is a representation of the C type g_socket_client_get_protocol.
func (recv *SocketClient) GetProtocol() SocketProtocol {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := socketClientGetProtocolFunction_Set()
	if err == nil {
		ret = socketClientGetProtocolFunction.Invoke(inArgs[:], nil)
	}

	retGo := SocketProtocol(ret.Int32())

	return retGo
}

// UNSUPPORTED : C value 'g_socket_client_get_proxy_resolver' : return type 'ProxyResolver' not supported

var socketClientGetSocketTypeFunction *gi.Function
var socketClientGetSocketTypeFunction_Once sync.Once

func socketClientGetSocketTypeFunction_Set() error {
	var err error
	socketClientGetSocketTypeFunction_Once.Do(func() {
		err = socketClientObject_Set()
		if err != nil {
			return
		}
		socketClientGetSocketTypeFunction, err = socketClientObject.InvokerNew("get_socket_type")
	})
	return err
}

// GetSocketType is a representation of the C type g_socket_client_get_socket_type.
func (recv *SocketClient) GetSocketType() SocketType {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := socketClientGetSocketTypeFunction_Set()
	if err == nil {
		ret = socketClientGetSocketTypeFunction.Invoke(inArgs[:], nil)
	}

	retGo := SocketType(ret.Int32())

	return retGo
}

var socketClientGetTimeoutFunction *gi.Function
var socketClientGetTimeoutFunction_Once sync.Once

func socketClientGetTimeoutFunction_Set() error {
	var err error
	socketClientGetTimeoutFunction_Once.Do(func() {
		err = socketClientObject_Set()
		if err != nil {
			return
		}
		socketClientGetTimeoutFunction, err = socketClientObject.InvokerNew("get_timeout")
	})
	return err
}

// GetTimeout is a representation of the C type g_socket_client_get_timeout.
func (recv *SocketClient) GetTimeout() uint32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := socketClientGetTimeoutFunction_Set()
	if err == nil {
		ret = socketClientGetTimeoutFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo
}

var socketClientGetTlsFunction *gi.Function
var socketClientGetTlsFunction_Once sync.Once

func socketClientGetTlsFunction_Set() error {
	var err error
	socketClientGetTlsFunction_Once.Do(func() {
		err = socketClientObject_Set()
		if err != nil {
			return
		}
		socketClientGetTlsFunction, err = socketClientObject.InvokerNew("get_tls")
	})
	return err
}

// GetTls is a representation of the C type g_socket_client_get_tls.
func (recv *SocketClient) GetTls() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := socketClientGetTlsFunction_Set()
	if err == nil {
		ret = socketClientGetTlsFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'g_socket_client_get_tls_validation_flags' : return type 'TlsCertificateFlags' not supported

var socketClientSetEnableProxyFunction *gi.Function
var socketClientSetEnableProxyFunction_Once sync.Once

func socketClientSetEnableProxyFunction_Set() error {
	var err error
	socketClientSetEnableProxyFunction_Once.Do(func() {
		err = socketClientObject_Set()
		if err != nil {
			return
		}
		socketClientSetEnableProxyFunction, err = socketClientObject.InvokerNew("set_enable_proxy")
	})
	return err
}

// SetEnableProxy is a representation of the C type g_socket_client_set_enable_proxy.
func (recv *SocketClient) SetEnableProxy(enable bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetBoolean(enable)

	err := socketClientSetEnableProxyFunction_Set()
	if err == nil {
		socketClientSetEnableProxyFunction.Invoke(inArgs[:], nil)
	}

	return
}

var socketClientSetFamilyFunction *gi.Function
var socketClientSetFamilyFunction_Once sync.Once

func socketClientSetFamilyFunction_Set() error {
	var err error
	socketClientSetFamilyFunction_Once.Do(func() {
		err = socketClientObject_Set()
		if err != nil {
			return
		}
		socketClientSetFamilyFunction, err = socketClientObject.InvokerNew("set_family")
	})
	return err
}

// SetFamily is a representation of the C type g_socket_client_set_family.
func (recv *SocketClient) SetFamily(family SocketFamily) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(int32(family))

	err := socketClientSetFamilyFunction_Set()
	if err == nil {
		socketClientSetFamilyFunction.Invoke(inArgs[:], nil)
	}

	return
}

var socketClientSetLocalAddressFunction *gi.Function
var socketClientSetLocalAddressFunction_Once sync.Once

func socketClientSetLocalAddressFunction_Set() error {
	var err error
	socketClientSetLocalAddressFunction_Once.Do(func() {
		err = socketClientObject_Set()
		if err != nil {
			return
		}
		socketClientSetLocalAddressFunction, err = socketClientObject.InvokerNew("set_local_address")
	})
	return err
}

// SetLocalAddress is a representation of the C type g_socket_client_set_local_address.
func (recv *SocketClient) SetLocalAddress(address *SocketAddress) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(address.Native())

	err := socketClientSetLocalAddressFunction_Set()
	if err == nil {
		socketClientSetLocalAddressFunction.Invoke(inArgs[:], nil)
	}

	return
}

var socketClientSetProtocolFunction *gi.Function
var socketClientSetProtocolFunction_Once sync.Once

func socketClientSetProtocolFunction_Set() error {
	var err error
	socketClientSetProtocolFunction_Once.Do(func() {
		err = socketClientObject_Set()
		if err != nil {
			return
		}
		socketClientSetProtocolFunction, err = socketClientObject.InvokerNew("set_protocol")
	})
	return err
}

// SetProtocol is a representation of the C type g_socket_client_set_protocol.
func (recv *SocketClient) SetProtocol(protocol SocketProtocol) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(int32(protocol))

	err := socketClientSetProtocolFunction_Set()
	if err == nil {
		socketClientSetProtocolFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'g_socket_client_set_proxy_resolver' : parameter 'proxy_resolver' of type 'ProxyResolver' not supported

var socketClientSetSocketTypeFunction *gi.Function
var socketClientSetSocketTypeFunction_Once sync.Once

func socketClientSetSocketTypeFunction_Set() error {
	var err error
	socketClientSetSocketTypeFunction_Once.Do(func() {
		err = socketClientObject_Set()
		if err != nil {
			return
		}
		socketClientSetSocketTypeFunction, err = socketClientObject.InvokerNew("set_socket_type")
	})
	return err
}

// SetSocketType is a representation of the C type g_socket_client_set_socket_type.
func (recv *SocketClient) SetSocketType(type_ SocketType) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(int32(type_))

	err := socketClientSetSocketTypeFunction_Set()
	if err == nil {
		socketClientSetSocketTypeFunction.Invoke(inArgs[:], nil)
	}

	return
}

var socketClientSetTimeoutFunction *gi.Function
var socketClientSetTimeoutFunction_Once sync.Once

func socketClientSetTimeoutFunction_Set() error {
	var err error
	socketClientSetTimeoutFunction_Once.Do(func() {
		err = socketClientObject_Set()
		if err != nil {
			return
		}
		socketClientSetTimeoutFunction, err = socketClientObject.InvokerNew("set_timeout")
	})
	return err
}

// SetTimeout is a representation of the C type g_socket_client_set_timeout.
func (recv *SocketClient) SetTimeout(timeout uint32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetUint32(timeout)

	err := socketClientSetTimeoutFunction_Set()
	if err == nil {
		socketClientSetTimeoutFunction.Invoke(inArgs[:], nil)
	}

	return
}

var socketClientSetTlsFunction *gi.Function
var socketClientSetTlsFunction_Once sync.Once

func socketClientSetTlsFunction_Set() error {
	var err error
	socketClientSetTlsFunction_Once.Do(func() {
		err = socketClientObject_Set()
		if err != nil {
			return
		}
		socketClientSetTlsFunction, err = socketClientObject.InvokerNew("set_tls")
	})
	return err
}

// SetTls is a representation of the C type g_socket_client_set_tls.
func (recv *SocketClient) SetTls(tls bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetBoolean(tls)

	err := socketClientSetTlsFunction_Set()
	if err == nil {
		socketClientSetTlsFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'g_socket_client_set_tls_validation_flags' : parameter 'flags' of type 'TlsCertificateFlags' not supported

var socketConnectionObject *gi.Object
var socketConnectionObject_Once sync.Once

func socketConnectionObject_Set() error {
	var err error
	socketConnectionObject_Once.Do(func() {
		socketConnectionObject, err = gi.ObjectNew("Gio", "SocketConnection")
	})
	return err
}

type SocketConnection struct {
	native unsafe.Pointer
}

func SocketConnectionNewFromNative(native unsafe.Pointer) *SocketConnection {
	instance := &SocketConnection{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// IOStream upcasts to *IOStream
func (recv *SocketConnection) IOStream() *IOStream {
	return IOStreamNewFromNative(recv.Native())
}

// Object upcasts to *Object
func (recv *SocketConnection) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToSocketConnection down casts any arbitrary Object to SocketConnection.
Exercise care, as this is a potentially dangerous function
if the Object is not a SocketConnection.
*/
func CastToSocketConnection(object *gobject.Object) *SocketConnection {
	return SocketConnectionNewFromNative(object.Native())
}

// Equals compares this SocketConnection with another SocketConnection, and returns true if they represent the same GObject.
func (recv *SocketConnection) Equals(other *SocketConnection) bool {
	return other.Native() == recv.Native()
}

func (recv *SocketConnection) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *SocketConnection) FieldParentInstance() *IOStream {
	argValue := gi.ObjectFieldGet(socketConnectionObject, recv.Native(), "parent_instance")
	value := IOStreamNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *SocketConnection) SetFieldParentInstance(value *IOStream) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(socketConnectionObject, recv.Native(), "parent_instance", argValue)
}

// FieldPriv returns the C field 'priv'.
func (recv *SocketConnection) FieldPriv() *SocketConnectionPrivate {
	argValue := gi.ObjectFieldGet(socketConnectionObject, recv.Native(), "priv")
	value := SocketConnectionPrivateNewFromNative(argValue.Pointer())
	return value
}

// SetFieldPriv sets the value of the C field 'priv'.
func (recv *SocketConnection) SetFieldPriv(value *SocketConnectionPrivate) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(socketConnectionObject, recv.Native(), "priv", argValue)
}

var socketConnectionConnectFunction *gi.Function
var socketConnectionConnectFunction_Once sync.Once

func socketConnectionConnectFunction_Set() error {
	var err error
	socketConnectionConnectFunction_Once.Do(func() {
		err = socketConnectionObject_Set()
		if err != nil {
			return
		}
		socketConnectionConnectFunction, err = socketConnectionObject.InvokerNew("connect")
	})
	return err
}

// Connect is a representation of the C type g_socket_connection_connect.
func (recv *SocketConnection) Connect(address *SocketAddress, cancellable *Cancellable) bool {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(address.Native())
	inArgs[2].SetPointer(cancellable.Native())

	var ret gi.Argument

	err := socketConnectionConnectFunction_Set()
	if err == nil {
		ret = socketConnectionConnectFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'g_socket_connection_connect_async' : parameter 'callback' of type 'AsyncReadyCallback' not supported

// UNSUPPORTED : C value 'g_socket_connection_connect_finish' : parameter 'result' of type 'AsyncResult' not supported

var socketConnectionGetLocalAddressFunction *gi.Function
var socketConnectionGetLocalAddressFunction_Once sync.Once

func socketConnectionGetLocalAddressFunction_Set() error {
	var err error
	socketConnectionGetLocalAddressFunction_Once.Do(func() {
		err = socketConnectionObject_Set()
		if err != nil {
			return
		}
		socketConnectionGetLocalAddressFunction, err = socketConnectionObject.InvokerNew("get_local_address")
	})
	return err
}

// GetLocalAddress is a representation of the C type g_socket_connection_get_local_address.
func (recv *SocketConnection) GetLocalAddress() *SocketAddress {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := socketConnectionGetLocalAddressFunction_Set()
	if err == nil {
		ret = socketConnectionGetLocalAddressFunction.Invoke(inArgs[:], nil)
	}

	retGo := SocketAddressNewFromNative(ret.Pointer())

	return retGo
}

var socketConnectionGetRemoteAddressFunction *gi.Function
var socketConnectionGetRemoteAddressFunction_Once sync.Once

func socketConnectionGetRemoteAddressFunction_Set() error {
	var err error
	socketConnectionGetRemoteAddressFunction_Once.Do(func() {
		err = socketConnectionObject_Set()
		if err != nil {
			return
		}
		socketConnectionGetRemoteAddressFunction, err = socketConnectionObject.InvokerNew("get_remote_address")
	})
	return err
}

// GetRemoteAddress is a representation of the C type g_socket_connection_get_remote_address.
func (recv *SocketConnection) GetRemoteAddress() *SocketAddress {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := socketConnectionGetRemoteAddressFunction_Set()
	if err == nil {
		ret = socketConnectionGetRemoteAddressFunction.Invoke(inArgs[:], nil)
	}

	retGo := SocketAddressNewFromNative(ret.Pointer())

	return retGo
}

var socketConnectionGetSocketFunction *gi.Function
var socketConnectionGetSocketFunction_Once sync.Once

func socketConnectionGetSocketFunction_Set() error {
	var err error
	socketConnectionGetSocketFunction_Once.Do(func() {
		err = socketConnectionObject_Set()
		if err != nil {
			return
		}
		socketConnectionGetSocketFunction, err = socketConnectionObject.InvokerNew("get_socket")
	})
	return err
}

// GetSocket is a representation of the C type g_socket_connection_get_socket.
func (recv *SocketConnection) GetSocket() *Socket {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := socketConnectionGetSocketFunction_Set()
	if err == nil {
		ret = socketConnectionGetSocketFunction.Invoke(inArgs[:], nil)
	}

	retGo := SocketNewFromNative(ret.Pointer())

	return retGo
}

var socketConnectionIsConnectedFunction *gi.Function
var socketConnectionIsConnectedFunction_Once sync.Once

func socketConnectionIsConnectedFunction_Set() error {
	var err error
	socketConnectionIsConnectedFunction_Once.Do(func() {
		err = socketConnectionObject_Set()
		if err != nil {
			return
		}
		socketConnectionIsConnectedFunction, err = socketConnectionObject.InvokerNew("is_connected")
	})
	return err
}

// IsConnected is a representation of the C type g_socket_connection_is_connected.
func (recv *SocketConnection) IsConnected() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := socketConnectionIsConnectedFunction_Set()
	if err == nil {
		ret = socketConnectionIsConnectedFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var socketControlMessageObject *gi.Object
var socketControlMessageObject_Once sync.Once

func socketControlMessageObject_Set() error {
	var err error
	socketControlMessageObject_Once.Do(func() {
		socketControlMessageObject, err = gi.ObjectNew("Gio", "SocketControlMessage")
	})
	return err
}

type SocketControlMessage struct {
	native unsafe.Pointer
}

func SocketControlMessageNewFromNative(native unsafe.Pointer) *SocketControlMessage {
	instance := &SocketControlMessage{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *SocketControlMessage) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToSocketControlMessage down casts any arbitrary Object to SocketControlMessage.
Exercise care, as this is a potentially dangerous function
if the Object is not a SocketControlMessage.
*/
func CastToSocketControlMessage(object *gobject.Object) *SocketControlMessage {
	return SocketControlMessageNewFromNative(object.Native())
}

// Equals compares this SocketControlMessage with another SocketControlMessage, and returns true if they represent the same GObject.
func (recv *SocketControlMessage) Equals(other *SocketControlMessage) bool {
	return other.Native() == recv.Native()
}

func (recv *SocketControlMessage) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *SocketControlMessage) FieldParentInstance() *gobject.Object {
	argValue := gi.ObjectFieldGet(socketControlMessageObject, recv.Native(), "parent_instance")
	value := gobject.ObjectNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *SocketControlMessage) SetFieldParentInstance(value *gobject.Object) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(socketControlMessageObject, recv.Native(), "parent_instance", argValue)
}

// FieldPriv returns the C field 'priv'.
func (recv *SocketControlMessage) FieldPriv() *SocketControlMessagePrivate {
	argValue := gi.ObjectFieldGet(socketControlMessageObject, recv.Native(), "priv")
	value := SocketControlMessagePrivateNewFromNative(argValue.Pointer())
	return value
}

// SetFieldPriv sets the value of the C field 'priv'.
func (recv *SocketControlMessage) SetFieldPriv(value *SocketControlMessagePrivate) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(socketControlMessageObject, recv.Native(), "priv", argValue)
}

var socketControlMessageGetLevelFunction *gi.Function
var socketControlMessageGetLevelFunction_Once sync.Once

func socketControlMessageGetLevelFunction_Set() error {
	var err error
	socketControlMessageGetLevelFunction_Once.Do(func() {
		err = socketControlMessageObject_Set()
		if err != nil {
			return
		}
		socketControlMessageGetLevelFunction, err = socketControlMessageObject.InvokerNew("get_level")
	})
	return err
}

// GetLevel is a representation of the C type g_socket_control_message_get_level.
func (recv *SocketControlMessage) GetLevel() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := socketControlMessageGetLevelFunction_Set()
	if err == nil {
		ret = socketControlMessageGetLevelFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var socketControlMessageGetMsgTypeFunction *gi.Function
var socketControlMessageGetMsgTypeFunction_Once sync.Once

func socketControlMessageGetMsgTypeFunction_Set() error {
	var err error
	socketControlMessageGetMsgTypeFunction_Once.Do(func() {
		err = socketControlMessageObject_Set()
		if err != nil {
			return
		}
		socketControlMessageGetMsgTypeFunction, err = socketControlMessageObject.InvokerNew("get_msg_type")
	})
	return err
}

// GetMsgType is a representation of the C type g_socket_control_message_get_msg_type.
func (recv *SocketControlMessage) GetMsgType() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := socketControlMessageGetMsgTypeFunction_Set()
	if err == nil {
		ret = socketControlMessageGetMsgTypeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var socketControlMessageGetSizeFunction *gi.Function
var socketControlMessageGetSizeFunction_Once sync.Once

func socketControlMessageGetSizeFunction_Set() error {
	var err error
	socketControlMessageGetSizeFunction_Once.Do(func() {
		err = socketControlMessageObject_Set()
		if err != nil {
			return
		}
		socketControlMessageGetSizeFunction, err = socketControlMessageObject.InvokerNew("get_size")
	})
	return err
}

// GetSize is a representation of the C type g_socket_control_message_get_size.
func (recv *SocketControlMessage) GetSize() uint64 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := socketControlMessageGetSizeFunction_Set()
	if err == nil {
		ret = socketControlMessageGetSizeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint64()

	return retGo
}

// UNSUPPORTED : C value 'g_socket_control_message_serialize' : parameter 'data' of type 'gpointer' not supported

var socketListenerObject *gi.Object
var socketListenerObject_Once sync.Once

func socketListenerObject_Set() error {
	var err error
	socketListenerObject_Once.Do(func() {
		socketListenerObject, err = gi.ObjectNew("Gio", "SocketListener")
	})
	return err
}

type SocketListener struct {
	native unsafe.Pointer
}

func SocketListenerNewFromNative(native unsafe.Pointer) *SocketListener {
	instance := &SocketListener{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *SocketListener) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToSocketListener down casts any arbitrary Object to SocketListener.
Exercise care, as this is a potentially dangerous function
if the Object is not a SocketListener.
*/
func CastToSocketListener(object *gobject.Object) *SocketListener {
	return SocketListenerNewFromNative(object.Native())
}

// Equals compares this SocketListener with another SocketListener, and returns true if they represent the same GObject.
func (recv *SocketListener) Equals(other *SocketListener) bool {
	return other.Native() == recv.Native()
}

func (recv *SocketListener) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *SocketListener) FieldParentInstance() *gobject.Object {
	argValue := gi.ObjectFieldGet(socketListenerObject, recv.Native(), "parent_instance")
	value := gobject.ObjectNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *SocketListener) SetFieldParentInstance(value *gobject.Object) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(socketListenerObject, recv.Native(), "parent_instance", argValue)
}

// FieldPriv returns the C field 'priv'.
func (recv *SocketListener) FieldPriv() *SocketListenerPrivate {
	argValue := gi.ObjectFieldGet(socketListenerObject, recv.Native(), "priv")
	value := SocketListenerPrivateNewFromNative(argValue.Pointer())
	return value
}

// SetFieldPriv sets the value of the C field 'priv'.
func (recv *SocketListener) SetFieldPriv(value *SocketListenerPrivate) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(socketListenerObject, recv.Native(), "priv", argValue)
}

var socketListenerNewFunction *gi.Function
var socketListenerNewFunction_Once sync.Once

func socketListenerNewFunction_Set() error {
	var err error
	socketListenerNewFunction_Once.Do(func() {
		err = socketListenerObject_Set()
		if err != nil {
			return
		}
		socketListenerNewFunction, err = socketListenerObject.InvokerNew("new")
	})
	return err
}

// SocketListenerNew is a representation of the C type g_socket_listener_new.
func SocketListenerNew() *SocketListener {

	var ret gi.Argument

	err := socketListenerNewFunction_Set()
	if err == nil {
		ret = socketListenerNewFunction.Invoke(nil, nil)
	}

	retGo := SocketListenerNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

var socketListenerAcceptFunction *gi.Function
var socketListenerAcceptFunction_Once sync.Once

func socketListenerAcceptFunction_Set() error {
	var err error
	socketListenerAcceptFunction_Once.Do(func() {
		err = socketListenerObject_Set()
		if err != nil {
			return
		}
		socketListenerAcceptFunction, err = socketListenerObject.InvokerNew("accept")
	})
	return err
}

// Accept is a representation of the C type g_socket_listener_accept.
func (recv *SocketListener) Accept(cancellable *Cancellable) (*SocketConnection, *gobject.Object) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(cancellable.Native())

	var outArgs [1]gi.Argument
	var ret gi.Argument

	err := socketListenerAcceptFunction_Set()
	if err == nil {
		ret = socketListenerAcceptFunction.Invoke(inArgs[:], outArgs[:])
	}

	retGo := SocketConnectionNewFromNative(ret.Pointer())
	out0 := gobject.ObjectNewFromNative(outArgs[0].Pointer())

	return retGo, out0
}

// UNSUPPORTED : C value 'g_socket_listener_accept_async' : parameter 'callback' of type 'AsyncReadyCallback' not supported

// UNSUPPORTED : C value 'g_socket_listener_accept_finish' : parameter 'result' of type 'AsyncResult' not supported

var socketListenerAcceptSocketFunction *gi.Function
var socketListenerAcceptSocketFunction_Once sync.Once

func socketListenerAcceptSocketFunction_Set() error {
	var err error
	socketListenerAcceptSocketFunction_Once.Do(func() {
		err = socketListenerObject_Set()
		if err != nil {
			return
		}
		socketListenerAcceptSocketFunction, err = socketListenerObject.InvokerNew("accept_socket")
	})
	return err
}

// AcceptSocket is a representation of the C type g_socket_listener_accept_socket.
func (recv *SocketListener) AcceptSocket(cancellable *Cancellable) (*Socket, *gobject.Object) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(cancellable.Native())

	var outArgs [1]gi.Argument
	var ret gi.Argument

	err := socketListenerAcceptSocketFunction_Set()
	if err == nil {
		ret = socketListenerAcceptSocketFunction.Invoke(inArgs[:], outArgs[:])
	}

	retGo := SocketNewFromNative(ret.Pointer())
	out0 := gobject.ObjectNewFromNative(outArgs[0].Pointer())

	return retGo, out0
}

// UNSUPPORTED : C value 'g_socket_listener_accept_socket_async' : parameter 'callback' of type 'AsyncReadyCallback' not supported

// UNSUPPORTED : C value 'g_socket_listener_accept_socket_finish' : parameter 'result' of type 'AsyncResult' not supported

var socketListenerAddAddressFunction *gi.Function
var socketListenerAddAddressFunction_Once sync.Once

func socketListenerAddAddressFunction_Set() error {
	var err error
	socketListenerAddAddressFunction_Once.Do(func() {
		err = socketListenerObject_Set()
		if err != nil {
			return
		}
		socketListenerAddAddressFunction, err = socketListenerObject.InvokerNew("add_address")
	})
	return err
}

// AddAddress is a representation of the C type g_socket_listener_add_address.
func (recv *SocketListener) AddAddress(address *SocketAddress, type_ SocketType, protocol SocketProtocol, sourceObject *gobject.Object) (bool, *SocketAddress) {
	var inArgs [5]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(address.Native())
	inArgs[2].SetInt32(int32(type_))
	inArgs[3].SetInt32(int32(protocol))
	inArgs[4].SetPointer(sourceObject.Native())

	var outArgs [1]gi.Argument
	var ret gi.Argument

	err := socketListenerAddAddressFunction_Set()
	if err == nil {
		ret = socketListenerAddAddressFunction.Invoke(inArgs[:], outArgs[:])
	}

	retGo := ret.Boolean()
	out0 := SocketAddressNewFromNative(outArgs[0].Pointer())

	return retGo, out0
}

var socketListenerAddAnyInetPortFunction *gi.Function
var socketListenerAddAnyInetPortFunction_Once sync.Once

func socketListenerAddAnyInetPortFunction_Set() error {
	var err error
	socketListenerAddAnyInetPortFunction_Once.Do(func() {
		err = socketListenerObject_Set()
		if err != nil {
			return
		}
		socketListenerAddAnyInetPortFunction, err = socketListenerObject.InvokerNew("add_any_inet_port")
	})
	return err
}

// AddAnyInetPort is a representation of the C type g_socket_listener_add_any_inet_port.
func (recv *SocketListener) AddAnyInetPort(sourceObject *gobject.Object) uint16 {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(sourceObject.Native())

	var ret gi.Argument

	err := socketListenerAddAnyInetPortFunction_Set()
	if err == nil {
		ret = socketListenerAddAnyInetPortFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint16()

	return retGo
}

var socketListenerAddInetPortFunction *gi.Function
var socketListenerAddInetPortFunction_Once sync.Once

func socketListenerAddInetPortFunction_Set() error {
	var err error
	socketListenerAddInetPortFunction_Once.Do(func() {
		err = socketListenerObject_Set()
		if err != nil {
			return
		}
		socketListenerAddInetPortFunction, err = socketListenerObject.InvokerNew("add_inet_port")
	})
	return err
}

// AddInetPort is a representation of the C type g_socket_listener_add_inet_port.
func (recv *SocketListener) AddInetPort(port uint16, sourceObject *gobject.Object) bool {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetUint16(port)
	inArgs[2].SetPointer(sourceObject.Native())

	var ret gi.Argument

	err := socketListenerAddInetPortFunction_Set()
	if err == nil {
		ret = socketListenerAddInetPortFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var socketListenerAddSocketFunction *gi.Function
var socketListenerAddSocketFunction_Once sync.Once

func socketListenerAddSocketFunction_Set() error {
	var err error
	socketListenerAddSocketFunction_Once.Do(func() {
		err = socketListenerObject_Set()
		if err != nil {
			return
		}
		socketListenerAddSocketFunction, err = socketListenerObject.InvokerNew("add_socket")
	})
	return err
}

// AddSocket is a representation of the C type g_socket_listener_add_socket.
func (recv *SocketListener) AddSocket(socket *Socket, sourceObject *gobject.Object) bool {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(socket.Native())
	inArgs[2].SetPointer(sourceObject.Native())

	var ret gi.Argument

	err := socketListenerAddSocketFunction_Set()
	if err == nil {
		ret = socketListenerAddSocketFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var socketListenerCloseFunction *gi.Function
var socketListenerCloseFunction_Once sync.Once

func socketListenerCloseFunction_Set() error {
	var err error
	socketListenerCloseFunction_Once.Do(func() {
		err = socketListenerObject_Set()
		if err != nil {
			return
		}
		socketListenerCloseFunction, err = socketListenerObject.InvokerNew("close")
	})
	return err
}

// Close is a representation of the C type g_socket_listener_close.
func (recv *SocketListener) Close() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := socketListenerCloseFunction_Set()
	if err == nil {
		socketListenerCloseFunction.Invoke(inArgs[:], nil)
	}

	return
}

var socketListenerSetBacklogFunction *gi.Function
var socketListenerSetBacklogFunction_Once sync.Once

func socketListenerSetBacklogFunction_Set() error {
	var err error
	socketListenerSetBacklogFunction_Once.Do(func() {
		err = socketListenerObject_Set()
		if err != nil {
			return
		}
		socketListenerSetBacklogFunction, err = socketListenerObject.InvokerNew("set_backlog")
	})
	return err
}

// SetBacklog is a representation of the C type g_socket_listener_set_backlog.
func (recv *SocketListener) SetBacklog(listenBacklog int32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(listenBacklog)

	err := socketListenerSetBacklogFunction_Set()
	if err == nil {
		socketListenerSetBacklogFunction.Invoke(inArgs[:], nil)
	}

	return
}

var socketServiceObject *gi.Object
var socketServiceObject_Once sync.Once

func socketServiceObject_Set() error {
	var err error
	socketServiceObject_Once.Do(func() {
		socketServiceObject, err = gi.ObjectNew("Gio", "SocketService")
	})
	return err
}

type SocketService struct {
	native unsafe.Pointer
}

func SocketServiceNewFromNative(native unsafe.Pointer) *SocketService {
	instance := &SocketService{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// SocketListener upcasts to *SocketListener
func (recv *SocketService) SocketListener() *SocketListener {
	return SocketListenerNewFromNative(recv.Native())
}

// Object upcasts to *Object
func (recv *SocketService) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToSocketService down casts any arbitrary Object to SocketService.
Exercise care, as this is a potentially dangerous function
if the Object is not a SocketService.
*/
func CastToSocketService(object *gobject.Object) *SocketService {
	return SocketServiceNewFromNative(object.Native())
}

// Equals compares this SocketService with another SocketService, and returns true if they represent the same GObject.
func (recv *SocketService) Equals(other *SocketService) bool {
	return other.Native() == recv.Native()
}

func (recv *SocketService) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *SocketService) FieldParentInstance() *SocketListener {
	argValue := gi.ObjectFieldGet(socketServiceObject, recv.Native(), "parent_instance")
	value := SocketListenerNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *SocketService) SetFieldParentInstance(value *SocketListener) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(socketServiceObject, recv.Native(), "parent_instance", argValue)
}

// FieldPriv returns the C field 'priv'.
func (recv *SocketService) FieldPriv() *SocketServicePrivate {
	argValue := gi.ObjectFieldGet(socketServiceObject, recv.Native(), "priv")
	value := SocketServicePrivateNewFromNative(argValue.Pointer())
	return value
}

// SetFieldPriv sets the value of the C field 'priv'.
func (recv *SocketService) SetFieldPriv(value *SocketServicePrivate) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(socketServiceObject, recv.Native(), "priv", argValue)
}

var socketServiceNewFunction *gi.Function
var socketServiceNewFunction_Once sync.Once

func socketServiceNewFunction_Set() error {
	var err error
	socketServiceNewFunction_Once.Do(func() {
		err = socketServiceObject_Set()
		if err != nil {
			return
		}
		socketServiceNewFunction, err = socketServiceObject.InvokerNew("new")
	})
	return err
}

// SocketServiceNew is a representation of the C type g_socket_service_new.
func SocketServiceNew() *SocketService {

	var ret gi.Argument

	err := socketServiceNewFunction_Set()
	if err == nil {
		ret = socketServiceNewFunction.Invoke(nil, nil)
	}

	retGo := SocketServiceNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

var socketServiceIsActiveFunction *gi.Function
var socketServiceIsActiveFunction_Once sync.Once

func socketServiceIsActiveFunction_Set() error {
	var err error
	socketServiceIsActiveFunction_Once.Do(func() {
		err = socketServiceObject_Set()
		if err != nil {
			return
		}
		socketServiceIsActiveFunction, err = socketServiceObject.InvokerNew("is_active")
	})
	return err
}

// IsActive is a representation of the C type g_socket_service_is_active.
func (recv *SocketService) IsActive() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := socketServiceIsActiveFunction_Set()
	if err == nil {
		ret = socketServiceIsActiveFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var socketServiceStartFunction *gi.Function
var socketServiceStartFunction_Once sync.Once

func socketServiceStartFunction_Set() error {
	var err error
	socketServiceStartFunction_Once.Do(func() {
		err = socketServiceObject_Set()
		if err != nil {
			return
		}
		socketServiceStartFunction, err = socketServiceObject.InvokerNew("start")
	})
	return err
}

// Start is a representation of the C type g_socket_service_start.
func (recv *SocketService) Start() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := socketServiceStartFunction_Set()
	if err == nil {
		socketServiceStartFunction.Invoke(inArgs[:], nil)
	}

	return
}

var socketServiceStopFunction *gi.Function
var socketServiceStopFunction_Once sync.Once

func socketServiceStopFunction_Set() error {
	var err error
	socketServiceStopFunction_Once.Do(func() {
		err = socketServiceObject_Set()
		if err != nil {
			return
		}
		socketServiceStopFunction, err = socketServiceObject.InvokerNew("stop")
	})
	return err
}

// Stop is a representation of the C type g_socket_service_stop.
func (recv *SocketService) Stop() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := socketServiceStopFunction_Set()
	if err == nil {
		socketServiceStopFunction.Invoke(inArgs[:], nil)
	}

	return
}

var subprocessObject *gi.Object
var subprocessObject_Once sync.Once

func subprocessObject_Set() error {
	var err error
	subprocessObject_Once.Do(func() {
		subprocessObject, err = gi.ObjectNew("Gio", "Subprocess")
	})
	return err
}

type Subprocess struct {
	native unsafe.Pointer
}

func SubprocessNewFromNative(native unsafe.Pointer) *Subprocess {
	instance := &Subprocess{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *Subprocess) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToSubprocess down casts any arbitrary Object to Subprocess.
Exercise care, as this is a potentially dangerous function
if the Object is not a Subprocess.
*/
func CastToSubprocess(object *gobject.Object) *Subprocess {
	return SubprocessNewFromNative(object.Native())
}

// Equals compares this Subprocess with another Subprocess, and returns true if they represent the same GObject.
func (recv *Subprocess) Equals(other *Subprocess) bool {
	return other.Native() == recv.Native()
}

func (recv *Subprocess) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'g_subprocess_new' : parameter 'flags' of type 'SubprocessFlags' not supported

// UNSUPPORTED : C value 'g_subprocess_newv' : parameter 'argv' of type 'nil' not supported

// UNSUPPORTED : C value 'g_subprocess_communicate' : parameter 'stdin_buf' of type 'GLib.Bytes' not supported

// UNSUPPORTED : C value 'g_subprocess_communicate_async' : parameter 'stdin_buf' of type 'GLib.Bytes' not supported

// UNSUPPORTED : C value 'g_subprocess_communicate_finish' : parameter 'result' of type 'AsyncResult' not supported

var subprocessCommunicateUtf8Function *gi.Function
var subprocessCommunicateUtf8Function_Once sync.Once

func subprocessCommunicateUtf8Function_Set() error {
	var err error
	subprocessCommunicateUtf8Function_Once.Do(func() {
		err = subprocessObject_Set()
		if err != nil {
			return
		}
		subprocessCommunicateUtf8Function, err = subprocessObject.InvokerNew("communicate_utf8")
	})
	return err
}

// CommunicateUtf8 is a representation of the C type g_subprocess_communicate_utf8.
func (recv *Subprocess) CommunicateUtf8(stdinBuf string, cancellable *Cancellable) (bool, string, string) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(stdinBuf)
	inArgs[2].SetPointer(cancellable.Native())

	var outArgs [2]gi.Argument
	var ret gi.Argument

	err := subprocessCommunicateUtf8Function_Set()
	if err == nil {
		ret = subprocessCommunicateUtf8Function.Invoke(inArgs[:], outArgs[:])
	}

	retGo := ret.Boolean()
	out0 := outArgs[0].String(true)
	out1 := outArgs[1].String(true)

	return retGo, out0, out1
}

// UNSUPPORTED : C value 'g_subprocess_communicate_utf8_async' : parameter 'callback' of type 'AsyncReadyCallback' not supported

// UNSUPPORTED : C value 'g_subprocess_communicate_utf8_finish' : parameter 'result' of type 'AsyncResult' not supported

var subprocessForceExitFunction *gi.Function
var subprocessForceExitFunction_Once sync.Once

func subprocessForceExitFunction_Set() error {
	var err error
	subprocessForceExitFunction_Once.Do(func() {
		err = subprocessObject_Set()
		if err != nil {
			return
		}
		subprocessForceExitFunction, err = subprocessObject.InvokerNew("force_exit")
	})
	return err
}

// ForceExit is a representation of the C type g_subprocess_force_exit.
func (recv *Subprocess) ForceExit() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := subprocessForceExitFunction_Set()
	if err == nil {
		subprocessForceExitFunction.Invoke(inArgs[:], nil)
	}

	return
}

var subprocessGetExitStatusFunction *gi.Function
var subprocessGetExitStatusFunction_Once sync.Once

func subprocessGetExitStatusFunction_Set() error {
	var err error
	subprocessGetExitStatusFunction_Once.Do(func() {
		err = subprocessObject_Set()
		if err != nil {
			return
		}
		subprocessGetExitStatusFunction, err = subprocessObject.InvokerNew("get_exit_status")
	})
	return err
}

// GetExitStatus is a representation of the C type g_subprocess_get_exit_status.
func (recv *Subprocess) GetExitStatus() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := subprocessGetExitStatusFunction_Set()
	if err == nil {
		ret = subprocessGetExitStatusFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var subprocessGetIdentifierFunction *gi.Function
var subprocessGetIdentifierFunction_Once sync.Once

func subprocessGetIdentifierFunction_Set() error {
	var err error
	subprocessGetIdentifierFunction_Once.Do(func() {
		err = subprocessObject_Set()
		if err != nil {
			return
		}
		subprocessGetIdentifierFunction, err = subprocessObject.InvokerNew("get_identifier")
	})
	return err
}

// GetIdentifier is a representation of the C type g_subprocess_get_identifier.
func (recv *Subprocess) GetIdentifier() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := subprocessGetIdentifierFunction_Set()
	if err == nil {
		ret = subprocessGetIdentifierFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var subprocessGetIfExitedFunction *gi.Function
var subprocessGetIfExitedFunction_Once sync.Once

func subprocessGetIfExitedFunction_Set() error {
	var err error
	subprocessGetIfExitedFunction_Once.Do(func() {
		err = subprocessObject_Set()
		if err != nil {
			return
		}
		subprocessGetIfExitedFunction, err = subprocessObject.InvokerNew("get_if_exited")
	})
	return err
}

// GetIfExited is a representation of the C type g_subprocess_get_if_exited.
func (recv *Subprocess) GetIfExited() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := subprocessGetIfExitedFunction_Set()
	if err == nil {
		ret = subprocessGetIfExitedFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var subprocessGetIfSignaledFunction *gi.Function
var subprocessGetIfSignaledFunction_Once sync.Once

func subprocessGetIfSignaledFunction_Set() error {
	var err error
	subprocessGetIfSignaledFunction_Once.Do(func() {
		err = subprocessObject_Set()
		if err != nil {
			return
		}
		subprocessGetIfSignaledFunction, err = subprocessObject.InvokerNew("get_if_signaled")
	})
	return err
}

// GetIfSignaled is a representation of the C type g_subprocess_get_if_signaled.
func (recv *Subprocess) GetIfSignaled() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := subprocessGetIfSignaledFunction_Set()
	if err == nil {
		ret = subprocessGetIfSignaledFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var subprocessGetStatusFunction *gi.Function
var subprocessGetStatusFunction_Once sync.Once

func subprocessGetStatusFunction_Set() error {
	var err error
	subprocessGetStatusFunction_Once.Do(func() {
		err = subprocessObject_Set()
		if err != nil {
			return
		}
		subprocessGetStatusFunction, err = subprocessObject.InvokerNew("get_status")
	})
	return err
}

// GetStatus is a representation of the C type g_subprocess_get_status.
func (recv *Subprocess) GetStatus() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := subprocessGetStatusFunction_Set()
	if err == nil {
		ret = subprocessGetStatusFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var subprocessGetStderrPipeFunction *gi.Function
var subprocessGetStderrPipeFunction_Once sync.Once

func subprocessGetStderrPipeFunction_Set() error {
	var err error
	subprocessGetStderrPipeFunction_Once.Do(func() {
		err = subprocessObject_Set()
		if err != nil {
			return
		}
		subprocessGetStderrPipeFunction, err = subprocessObject.InvokerNew("get_stderr_pipe")
	})
	return err
}

// GetStderrPipe is a representation of the C type g_subprocess_get_stderr_pipe.
func (recv *Subprocess) GetStderrPipe() *InputStream {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := subprocessGetStderrPipeFunction_Set()
	if err == nil {
		ret = subprocessGetStderrPipeFunction.Invoke(inArgs[:], nil)
	}

	retGo := InputStreamNewFromNative(ret.Pointer())

	return retGo
}

var subprocessGetStdinPipeFunction *gi.Function
var subprocessGetStdinPipeFunction_Once sync.Once

func subprocessGetStdinPipeFunction_Set() error {
	var err error
	subprocessGetStdinPipeFunction_Once.Do(func() {
		err = subprocessObject_Set()
		if err != nil {
			return
		}
		subprocessGetStdinPipeFunction, err = subprocessObject.InvokerNew("get_stdin_pipe")
	})
	return err
}

// GetStdinPipe is a representation of the C type g_subprocess_get_stdin_pipe.
func (recv *Subprocess) GetStdinPipe() *OutputStream {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := subprocessGetStdinPipeFunction_Set()
	if err == nil {
		ret = subprocessGetStdinPipeFunction.Invoke(inArgs[:], nil)
	}

	retGo := OutputStreamNewFromNative(ret.Pointer())

	return retGo
}

var subprocessGetStdoutPipeFunction *gi.Function
var subprocessGetStdoutPipeFunction_Once sync.Once

func subprocessGetStdoutPipeFunction_Set() error {
	var err error
	subprocessGetStdoutPipeFunction_Once.Do(func() {
		err = subprocessObject_Set()
		if err != nil {
			return
		}
		subprocessGetStdoutPipeFunction, err = subprocessObject.InvokerNew("get_stdout_pipe")
	})
	return err
}

// GetStdoutPipe is a representation of the C type g_subprocess_get_stdout_pipe.
func (recv *Subprocess) GetStdoutPipe() *InputStream {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := subprocessGetStdoutPipeFunction_Set()
	if err == nil {
		ret = subprocessGetStdoutPipeFunction.Invoke(inArgs[:], nil)
	}

	retGo := InputStreamNewFromNative(ret.Pointer())

	return retGo
}

var subprocessGetSuccessfulFunction *gi.Function
var subprocessGetSuccessfulFunction_Once sync.Once

func subprocessGetSuccessfulFunction_Set() error {
	var err error
	subprocessGetSuccessfulFunction_Once.Do(func() {
		err = subprocessObject_Set()
		if err != nil {
			return
		}
		subprocessGetSuccessfulFunction, err = subprocessObject.InvokerNew("get_successful")
	})
	return err
}

// GetSuccessful is a representation of the C type g_subprocess_get_successful.
func (recv *Subprocess) GetSuccessful() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := subprocessGetSuccessfulFunction_Set()
	if err == nil {
		ret = subprocessGetSuccessfulFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var subprocessGetTermSigFunction *gi.Function
var subprocessGetTermSigFunction_Once sync.Once

func subprocessGetTermSigFunction_Set() error {
	var err error
	subprocessGetTermSigFunction_Once.Do(func() {
		err = subprocessObject_Set()
		if err != nil {
			return
		}
		subprocessGetTermSigFunction, err = subprocessObject.InvokerNew("get_term_sig")
	})
	return err
}

// GetTermSig is a representation of the C type g_subprocess_get_term_sig.
func (recv *Subprocess) GetTermSig() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := subprocessGetTermSigFunction_Set()
	if err == nil {
		ret = subprocessGetTermSigFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var subprocessSendSignalFunction *gi.Function
var subprocessSendSignalFunction_Once sync.Once

func subprocessSendSignalFunction_Set() error {
	var err error
	subprocessSendSignalFunction_Once.Do(func() {
		err = subprocessObject_Set()
		if err != nil {
			return
		}
		subprocessSendSignalFunction, err = subprocessObject.InvokerNew("send_signal")
	})
	return err
}

// SendSignal is a representation of the C type g_subprocess_send_signal.
func (recv *Subprocess) SendSignal(signalNum int32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(signalNum)

	err := subprocessSendSignalFunction_Set()
	if err == nil {
		subprocessSendSignalFunction.Invoke(inArgs[:], nil)
	}

	return
}

var subprocessWaitFunction *gi.Function
var subprocessWaitFunction_Once sync.Once

func subprocessWaitFunction_Set() error {
	var err error
	subprocessWaitFunction_Once.Do(func() {
		err = subprocessObject_Set()
		if err != nil {
			return
		}
		subprocessWaitFunction, err = subprocessObject.InvokerNew("wait")
	})
	return err
}

// Wait is a representation of the C type g_subprocess_wait.
func (recv *Subprocess) Wait(cancellable *Cancellable) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(cancellable.Native())

	var ret gi.Argument

	err := subprocessWaitFunction_Set()
	if err == nil {
		ret = subprocessWaitFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'g_subprocess_wait_async' : parameter 'callback' of type 'AsyncReadyCallback' not supported

var subprocessWaitCheckFunction *gi.Function
var subprocessWaitCheckFunction_Once sync.Once

func subprocessWaitCheckFunction_Set() error {
	var err error
	subprocessWaitCheckFunction_Once.Do(func() {
		err = subprocessObject_Set()
		if err != nil {
			return
		}
		subprocessWaitCheckFunction, err = subprocessObject.InvokerNew("wait_check")
	})
	return err
}

// WaitCheck is a representation of the C type g_subprocess_wait_check.
func (recv *Subprocess) WaitCheck(cancellable *Cancellable) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(cancellable.Native())

	var ret gi.Argument

	err := subprocessWaitCheckFunction_Set()
	if err == nil {
		ret = subprocessWaitCheckFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'g_subprocess_wait_check_async' : parameter 'callback' of type 'AsyncReadyCallback' not supported

// UNSUPPORTED : C value 'g_subprocess_wait_check_finish' : parameter 'result' of type 'AsyncResult' not supported

// UNSUPPORTED : C value 'g_subprocess_wait_finish' : parameter 'result' of type 'AsyncResult' not supported

var subprocessLauncherObject *gi.Object
var subprocessLauncherObject_Once sync.Once

func subprocessLauncherObject_Set() error {
	var err error
	subprocessLauncherObject_Once.Do(func() {
		subprocessLauncherObject, err = gi.ObjectNew("Gio", "SubprocessLauncher")
	})
	return err
}

type SubprocessLauncher struct {
	native unsafe.Pointer
}

func SubprocessLauncherNewFromNative(native unsafe.Pointer) *SubprocessLauncher {
	instance := &SubprocessLauncher{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *SubprocessLauncher) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToSubprocessLauncher down casts any arbitrary Object to SubprocessLauncher.
Exercise care, as this is a potentially dangerous function
if the Object is not a SubprocessLauncher.
*/
func CastToSubprocessLauncher(object *gobject.Object) *SubprocessLauncher {
	return SubprocessLauncherNewFromNative(object.Native())
}

// Equals compares this SubprocessLauncher with another SubprocessLauncher, and returns true if they represent the same GObject.
func (recv *SubprocessLauncher) Equals(other *SubprocessLauncher) bool {
	return other.Native() == recv.Native()
}

func (recv *SubprocessLauncher) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'g_subprocess_launcher_new' : parameter 'flags' of type 'SubprocessFlags' not supported

var subprocessLauncherGetenvFunction *gi.Function
var subprocessLauncherGetenvFunction_Once sync.Once

func subprocessLauncherGetenvFunction_Set() error {
	var err error
	subprocessLauncherGetenvFunction_Once.Do(func() {
		err = subprocessLauncherObject_Set()
		if err != nil {
			return
		}
		subprocessLauncherGetenvFunction, err = subprocessLauncherObject.InvokerNew("getenv")
	})
	return err
}

// Getenv is a representation of the C type g_subprocess_launcher_getenv.
func (recv *SubprocessLauncher) Getenv(variable string) string {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(variable)

	var ret gi.Argument

	err := subprocessLauncherGetenvFunction_Set()
	if err == nil {
		ret = subprocessLauncherGetenvFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'g_subprocess_launcher_set_child_setup' : parameter 'child_setup' of type 'GLib.SpawnChildSetupFunc' not supported

var subprocessLauncherSetCwdFunction *gi.Function
var subprocessLauncherSetCwdFunction_Once sync.Once

func subprocessLauncherSetCwdFunction_Set() error {
	var err error
	subprocessLauncherSetCwdFunction_Once.Do(func() {
		err = subprocessLauncherObject_Set()
		if err != nil {
			return
		}
		subprocessLauncherSetCwdFunction, err = subprocessLauncherObject.InvokerNew("set_cwd")
	})
	return err
}

// SetCwd is a representation of the C type g_subprocess_launcher_set_cwd.
func (recv *SubprocessLauncher) SetCwd(cwd string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(cwd)

	err := subprocessLauncherSetCwdFunction_Set()
	if err == nil {
		subprocessLauncherSetCwdFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'g_subprocess_launcher_set_environ' : parameter 'env' of type 'nil' not supported

// UNSUPPORTED : C value 'g_subprocess_launcher_set_flags' : parameter 'flags' of type 'SubprocessFlags' not supported

var subprocessLauncherSetStderrFilePathFunction *gi.Function
var subprocessLauncherSetStderrFilePathFunction_Once sync.Once

func subprocessLauncherSetStderrFilePathFunction_Set() error {
	var err error
	subprocessLauncherSetStderrFilePathFunction_Once.Do(func() {
		err = subprocessLauncherObject_Set()
		if err != nil {
			return
		}
		subprocessLauncherSetStderrFilePathFunction, err = subprocessLauncherObject.InvokerNew("set_stderr_file_path")
	})
	return err
}

// SetStderrFilePath is a representation of the C type g_subprocess_launcher_set_stderr_file_path.
func (recv *SubprocessLauncher) SetStderrFilePath(path string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(path)

	err := subprocessLauncherSetStderrFilePathFunction_Set()
	if err == nil {
		subprocessLauncherSetStderrFilePathFunction.Invoke(inArgs[:], nil)
	}

	return
}

var subprocessLauncherSetStdinFilePathFunction *gi.Function
var subprocessLauncherSetStdinFilePathFunction_Once sync.Once

func subprocessLauncherSetStdinFilePathFunction_Set() error {
	var err error
	subprocessLauncherSetStdinFilePathFunction_Once.Do(func() {
		err = subprocessLauncherObject_Set()
		if err != nil {
			return
		}
		subprocessLauncherSetStdinFilePathFunction, err = subprocessLauncherObject.InvokerNew("set_stdin_file_path")
	})
	return err
}

// SetStdinFilePath is a representation of the C type g_subprocess_launcher_set_stdin_file_path.
func (recv *SubprocessLauncher) SetStdinFilePath(path string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(path)

	err := subprocessLauncherSetStdinFilePathFunction_Set()
	if err == nil {
		subprocessLauncherSetStdinFilePathFunction.Invoke(inArgs[:], nil)
	}

	return
}

var subprocessLauncherSetStdoutFilePathFunction *gi.Function
var subprocessLauncherSetStdoutFilePathFunction_Once sync.Once

func subprocessLauncherSetStdoutFilePathFunction_Set() error {
	var err error
	subprocessLauncherSetStdoutFilePathFunction_Once.Do(func() {
		err = subprocessLauncherObject_Set()
		if err != nil {
			return
		}
		subprocessLauncherSetStdoutFilePathFunction, err = subprocessLauncherObject.InvokerNew("set_stdout_file_path")
	})
	return err
}

// SetStdoutFilePath is a representation of the C type g_subprocess_launcher_set_stdout_file_path.
func (recv *SubprocessLauncher) SetStdoutFilePath(path string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(path)

	err := subprocessLauncherSetStdoutFilePathFunction_Set()
	if err == nil {
		subprocessLauncherSetStdoutFilePathFunction.Invoke(inArgs[:], nil)
	}

	return
}

var subprocessLauncherSetenvFunction *gi.Function
var subprocessLauncherSetenvFunction_Once sync.Once

func subprocessLauncherSetenvFunction_Set() error {
	var err error
	subprocessLauncherSetenvFunction_Once.Do(func() {
		err = subprocessLauncherObject_Set()
		if err != nil {
			return
		}
		subprocessLauncherSetenvFunction, err = subprocessLauncherObject.InvokerNew("setenv")
	})
	return err
}

// Setenv is a representation of the C type g_subprocess_launcher_setenv.
func (recv *SubprocessLauncher) Setenv(variable string, value string, overwrite bool) {
	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(variable)
	inArgs[2].SetString(value)
	inArgs[3].SetBoolean(overwrite)

	err := subprocessLauncherSetenvFunction_Set()
	if err == nil {
		subprocessLauncherSetenvFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'g_subprocess_launcher_spawn' : parameter 'error' of type 'GLib.Error' not supported

// UNSUPPORTED : C value 'g_subprocess_launcher_spawnv' : parameter 'argv' of type 'nil' not supported

var subprocessLauncherTakeFdFunction *gi.Function
var subprocessLauncherTakeFdFunction_Once sync.Once

func subprocessLauncherTakeFdFunction_Set() error {
	var err error
	subprocessLauncherTakeFdFunction_Once.Do(func() {
		err = subprocessLauncherObject_Set()
		if err != nil {
			return
		}
		subprocessLauncherTakeFdFunction, err = subprocessLauncherObject.InvokerNew("take_fd")
	})
	return err
}

// TakeFd is a representation of the C type g_subprocess_launcher_take_fd.
func (recv *SubprocessLauncher) TakeFd(sourceFd int32, targetFd int32) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(sourceFd)
	inArgs[2].SetInt32(targetFd)

	err := subprocessLauncherTakeFdFunction_Set()
	if err == nil {
		subprocessLauncherTakeFdFunction.Invoke(inArgs[:], nil)
	}

	return
}

var subprocessLauncherTakeStderrFdFunction *gi.Function
var subprocessLauncherTakeStderrFdFunction_Once sync.Once

func subprocessLauncherTakeStderrFdFunction_Set() error {
	var err error
	subprocessLauncherTakeStderrFdFunction_Once.Do(func() {
		err = subprocessLauncherObject_Set()
		if err != nil {
			return
		}
		subprocessLauncherTakeStderrFdFunction, err = subprocessLauncherObject.InvokerNew("take_stderr_fd")
	})
	return err
}

// TakeStderrFd is a representation of the C type g_subprocess_launcher_take_stderr_fd.
func (recv *SubprocessLauncher) TakeStderrFd(fd int32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(fd)

	err := subprocessLauncherTakeStderrFdFunction_Set()
	if err == nil {
		subprocessLauncherTakeStderrFdFunction.Invoke(inArgs[:], nil)
	}

	return
}

var subprocessLauncherTakeStdinFdFunction *gi.Function
var subprocessLauncherTakeStdinFdFunction_Once sync.Once

func subprocessLauncherTakeStdinFdFunction_Set() error {
	var err error
	subprocessLauncherTakeStdinFdFunction_Once.Do(func() {
		err = subprocessLauncherObject_Set()
		if err != nil {
			return
		}
		subprocessLauncherTakeStdinFdFunction, err = subprocessLauncherObject.InvokerNew("take_stdin_fd")
	})
	return err
}

// TakeStdinFd is a representation of the C type g_subprocess_launcher_take_stdin_fd.
func (recv *SubprocessLauncher) TakeStdinFd(fd int32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(fd)

	err := subprocessLauncherTakeStdinFdFunction_Set()
	if err == nil {
		subprocessLauncherTakeStdinFdFunction.Invoke(inArgs[:], nil)
	}

	return
}

var subprocessLauncherTakeStdoutFdFunction *gi.Function
var subprocessLauncherTakeStdoutFdFunction_Once sync.Once

func subprocessLauncherTakeStdoutFdFunction_Set() error {
	var err error
	subprocessLauncherTakeStdoutFdFunction_Once.Do(func() {
		err = subprocessLauncherObject_Set()
		if err != nil {
			return
		}
		subprocessLauncherTakeStdoutFdFunction, err = subprocessLauncherObject.InvokerNew("take_stdout_fd")
	})
	return err
}

// TakeStdoutFd is a representation of the C type g_subprocess_launcher_take_stdout_fd.
func (recv *SubprocessLauncher) TakeStdoutFd(fd int32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(fd)

	err := subprocessLauncherTakeStdoutFdFunction_Set()
	if err == nil {
		subprocessLauncherTakeStdoutFdFunction.Invoke(inArgs[:], nil)
	}

	return
}

var subprocessLauncherUnsetenvFunction *gi.Function
var subprocessLauncherUnsetenvFunction_Once sync.Once

func subprocessLauncherUnsetenvFunction_Set() error {
	var err error
	subprocessLauncherUnsetenvFunction_Once.Do(func() {
		err = subprocessLauncherObject_Set()
		if err != nil {
			return
		}
		subprocessLauncherUnsetenvFunction, err = subprocessLauncherObject.InvokerNew("unsetenv")
	})
	return err
}

// Unsetenv is a representation of the C type g_subprocess_launcher_unsetenv.
func (recv *SubprocessLauncher) Unsetenv(variable string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(variable)

	err := subprocessLauncherUnsetenvFunction_Set()
	if err == nil {
		subprocessLauncherUnsetenvFunction.Invoke(inArgs[:], nil)
	}

	return
}

var taskObject *gi.Object
var taskObject_Once sync.Once

func taskObject_Set() error {
	var err error
	taskObject_Once.Do(func() {
		taskObject, err = gi.ObjectNew("Gio", "Task")
	})
	return err
}

type Task struct {
	native unsafe.Pointer
}

func TaskNewFromNative(native unsafe.Pointer) *Task {
	instance := &Task{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *Task) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToTask down casts any arbitrary Object to Task.
Exercise care, as this is a potentially dangerous function
if the Object is not a Task.
*/
func CastToTask(object *gobject.Object) *Task {
	return TaskNewFromNative(object.Native())
}

// Equals compares this Task with another Task, and returns true if they represent the same GObject.
func (recv *Task) Equals(other *Task) bool {
	return other.Native() == recv.Native()
}

func (recv *Task) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'g_task_new' : parameter 'callback' of type 'AsyncReadyCallback' not supported

// UNSUPPORTED : C value 'g_task_attach_source' : parameter 'source' of type 'GLib.Source' not supported

var taskGetCancellableFunction *gi.Function
var taskGetCancellableFunction_Once sync.Once

func taskGetCancellableFunction_Set() error {
	var err error
	taskGetCancellableFunction_Once.Do(func() {
		err = taskObject_Set()
		if err != nil {
			return
		}
		taskGetCancellableFunction, err = taskObject.InvokerNew("get_cancellable")
	})
	return err
}

// GetCancellable is a representation of the C type g_task_get_cancellable.
func (recv *Task) GetCancellable() *Cancellable {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := taskGetCancellableFunction_Set()
	if err == nil {
		ret = taskGetCancellableFunction.Invoke(inArgs[:], nil)
	}

	retGo := CancellableNewFromNative(ret.Pointer())

	return retGo
}

var taskGetCheckCancellableFunction *gi.Function
var taskGetCheckCancellableFunction_Once sync.Once

func taskGetCheckCancellableFunction_Set() error {
	var err error
	taskGetCheckCancellableFunction_Once.Do(func() {
		err = taskObject_Set()
		if err != nil {
			return
		}
		taskGetCheckCancellableFunction, err = taskObject.InvokerNew("get_check_cancellable")
	})
	return err
}

// GetCheckCancellable is a representation of the C type g_task_get_check_cancellable.
func (recv *Task) GetCheckCancellable() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := taskGetCheckCancellableFunction_Set()
	if err == nil {
		ret = taskGetCheckCancellableFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var taskGetCompletedFunction *gi.Function
var taskGetCompletedFunction_Once sync.Once

func taskGetCompletedFunction_Set() error {
	var err error
	taskGetCompletedFunction_Once.Do(func() {
		err = taskObject_Set()
		if err != nil {
			return
		}
		taskGetCompletedFunction, err = taskObject.InvokerNew("get_completed")
	})
	return err
}

// GetCompleted is a representation of the C type g_task_get_completed.
func (recv *Task) GetCompleted() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := taskGetCompletedFunction_Set()
	if err == nil {
		ret = taskGetCompletedFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'g_task_get_context' : return type 'GLib.MainContext' not supported

var taskGetNameFunction *gi.Function
var taskGetNameFunction_Once sync.Once

func taskGetNameFunction_Set() error {
	var err error
	taskGetNameFunction_Once.Do(func() {
		err = taskObject_Set()
		if err != nil {
			return
		}
		taskGetNameFunction, err = taskObject.InvokerNew("get_name")
	})
	return err
}

// GetName is a representation of the C type g_task_get_name.
func (recv *Task) GetName() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := taskGetNameFunction_Set()
	if err == nil {
		ret = taskGetNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var taskGetPriorityFunction *gi.Function
var taskGetPriorityFunction_Once sync.Once

func taskGetPriorityFunction_Set() error {
	var err error
	taskGetPriorityFunction_Once.Do(func() {
		err = taskObject_Set()
		if err != nil {
			return
		}
		taskGetPriorityFunction, err = taskObject.InvokerNew("get_priority")
	})
	return err
}

// GetPriority is a representation of the C type g_task_get_priority.
func (recv *Task) GetPriority() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := taskGetPriorityFunction_Set()
	if err == nil {
		ret = taskGetPriorityFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var taskGetReturnOnCancelFunction *gi.Function
var taskGetReturnOnCancelFunction_Once sync.Once

func taskGetReturnOnCancelFunction_Set() error {
	var err error
	taskGetReturnOnCancelFunction_Once.Do(func() {
		err = taskObject_Set()
		if err != nil {
			return
		}
		taskGetReturnOnCancelFunction, err = taskObject.InvokerNew("get_return_on_cancel")
	})
	return err
}

// GetReturnOnCancel is a representation of the C type g_task_get_return_on_cancel.
func (recv *Task) GetReturnOnCancel() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := taskGetReturnOnCancelFunction_Set()
	if err == nil {
		ret = taskGetReturnOnCancelFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var taskGetSourceObjectFunction *gi.Function
var taskGetSourceObjectFunction_Once sync.Once

func taskGetSourceObjectFunction_Set() error {
	var err error
	taskGetSourceObjectFunction_Once.Do(func() {
		err = taskObject_Set()
		if err != nil {
			return
		}
		taskGetSourceObjectFunction, err = taskObject.InvokerNew("get_source_object")
	})
	return err
}

// GetSourceObject is a representation of the C type g_task_get_source_object.
func (recv *Task) GetSourceObject() *gobject.Object {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := taskGetSourceObjectFunction_Set()
	if err == nil {
		ret = taskGetSourceObjectFunction.Invoke(inArgs[:], nil)
	}

	retGo := gobject.ObjectNewFromNative(ret.Pointer())

	return retGo
}

// UNSUPPORTED : C value 'g_task_get_source_tag' : return type 'gpointer' not supported

// UNSUPPORTED : C value 'g_task_get_task_data' : return type 'gpointer' not supported

var taskHadErrorFunction *gi.Function
var taskHadErrorFunction_Once sync.Once

func taskHadErrorFunction_Set() error {
	var err error
	taskHadErrorFunction_Once.Do(func() {
		err = taskObject_Set()
		if err != nil {
			return
		}
		taskHadErrorFunction, err = taskObject.InvokerNew("had_error")
	})
	return err
}

// HadError is a representation of the C type g_task_had_error.
func (recv *Task) HadError() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := taskHadErrorFunction_Set()
	if err == nil {
		ret = taskHadErrorFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var taskPropagateBooleanFunction *gi.Function
var taskPropagateBooleanFunction_Once sync.Once

func taskPropagateBooleanFunction_Set() error {
	var err error
	taskPropagateBooleanFunction_Once.Do(func() {
		err = taskObject_Set()
		if err != nil {
			return
		}
		taskPropagateBooleanFunction, err = taskObject.InvokerNew("propagate_boolean")
	})
	return err
}

// PropagateBoolean is a representation of the C type g_task_propagate_boolean.
func (recv *Task) PropagateBoolean() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := taskPropagateBooleanFunction_Set()
	if err == nil {
		ret = taskPropagateBooleanFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var taskPropagateIntFunction *gi.Function
var taskPropagateIntFunction_Once sync.Once

func taskPropagateIntFunction_Set() error {
	var err error
	taskPropagateIntFunction_Once.Do(func() {
		err = taskObject_Set()
		if err != nil {
			return
		}
		taskPropagateIntFunction, err = taskObject.InvokerNew("propagate_int")
	})
	return err
}

// PropagateInt is a representation of the C type g_task_propagate_int.
func (recv *Task) PropagateInt() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := taskPropagateIntFunction_Set()
	if err == nil {
		ret = taskPropagateIntFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

// UNSUPPORTED : C value 'g_task_propagate_pointer' : return type 'gpointer' not supported

var taskReturnBooleanFunction *gi.Function
var taskReturnBooleanFunction_Once sync.Once

func taskReturnBooleanFunction_Set() error {
	var err error
	taskReturnBooleanFunction_Once.Do(func() {
		err = taskObject_Set()
		if err != nil {
			return
		}
		taskReturnBooleanFunction, err = taskObject.InvokerNew("return_boolean")
	})
	return err
}

// ReturnBoolean is a representation of the C type g_task_return_boolean.
func (recv *Task) ReturnBoolean(result bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetBoolean(result)

	err := taskReturnBooleanFunction_Set()
	if err == nil {
		taskReturnBooleanFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'g_task_return_error' : parameter 'error' of type 'GLib.Error' not supported

var taskReturnErrorIfCancelledFunction *gi.Function
var taskReturnErrorIfCancelledFunction_Once sync.Once

func taskReturnErrorIfCancelledFunction_Set() error {
	var err error
	taskReturnErrorIfCancelledFunction_Once.Do(func() {
		err = taskObject_Set()
		if err != nil {
			return
		}
		taskReturnErrorIfCancelledFunction, err = taskObject.InvokerNew("return_error_if_cancelled")
	})
	return err
}

// ReturnErrorIfCancelled is a representation of the C type g_task_return_error_if_cancelled.
func (recv *Task) ReturnErrorIfCancelled() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := taskReturnErrorIfCancelledFunction_Set()
	if err == nil {
		ret = taskReturnErrorIfCancelledFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var taskReturnIntFunction *gi.Function
var taskReturnIntFunction_Once sync.Once

func taskReturnIntFunction_Set() error {
	var err error
	taskReturnIntFunction_Once.Do(func() {
		err = taskObject_Set()
		if err != nil {
			return
		}
		taskReturnIntFunction, err = taskObject.InvokerNew("return_int")
	})
	return err
}

// ReturnInt is a representation of the C type g_task_return_int.
func (recv *Task) ReturnInt(result int32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(result)

	err := taskReturnIntFunction_Set()
	if err == nil {
		taskReturnIntFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'g_task_return_new_error' : parameter '...' of type 'nil' not supported

// UNSUPPORTED : C value 'g_task_return_pointer' : parameter 'result' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_task_run_in_thread' : parameter 'task_func' of type 'TaskThreadFunc' not supported

// UNSUPPORTED : C value 'g_task_run_in_thread_sync' : parameter 'task_func' of type 'TaskThreadFunc' not supported

var taskSetCheckCancellableFunction *gi.Function
var taskSetCheckCancellableFunction_Once sync.Once

func taskSetCheckCancellableFunction_Set() error {
	var err error
	taskSetCheckCancellableFunction_Once.Do(func() {
		err = taskObject_Set()
		if err != nil {
			return
		}
		taskSetCheckCancellableFunction, err = taskObject.InvokerNew("set_check_cancellable")
	})
	return err
}

// SetCheckCancellable is a representation of the C type g_task_set_check_cancellable.
func (recv *Task) SetCheckCancellable(checkCancellable bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetBoolean(checkCancellable)

	err := taskSetCheckCancellableFunction_Set()
	if err == nil {
		taskSetCheckCancellableFunction.Invoke(inArgs[:], nil)
	}

	return
}

var taskSetNameFunction *gi.Function
var taskSetNameFunction_Once sync.Once

func taskSetNameFunction_Set() error {
	var err error
	taskSetNameFunction_Once.Do(func() {
		err = taskObject_Set()
		if err != nil {
			return
		}
		taskSetNameFunction, err = taskObject.InvokerNew("set_name")
	})
	return err
}

// SetName is a representation of the C type g_task_set_name.
func (recv *Task) SetName(name string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(name)

	err := taskSetNameFunction_Set()
	if err == nil {
		taskSetNameFunction.Invoke(inArgs[:], nil)
	}

	return
}

var taskSetPriorityFunction *gi.Function
var taskSetPriorityFunction_Once sync.Once

func taskSetPriorityFunction_Set() error {
	var err error
	taskSetPriorityFunction_Once.Do(func() {
		err = taskObject_Set()
		if err != nil {
			return
		}
		taskSetPriorityFunction, err = taskObject.InvokerNew("set_priority")
	})
	return err
}

// SetPriority is a representation of the C type g_task_set_priority.
func (recv *Task) SetPriority(priority int32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(priority)

	err := taskSetPriorityFunction_Set()
	if err == nil {
		taskSetPriorityFunction.Invoke(inArgs[:], nil)
	}

	return
}

var taskSetReturnOnCancelFunction *gi.Function
var taskSetReturnOnCancelFunction_Once sync.Once

func taskSetReturnOnCancelFunction_Set() error {
	var err error
	taskSetReturnOnCancelFunction_Once.Do(func() {
		err = taskObject_Set()
		if err != nil {
			return
		}
		taskSetReturnOnCancelFunction, err = taskObject.InvokerNew("set_return_on_cancel")
	})
	return err
}

// SetReturnOnCancel is a representation of the C type g_task_set_return_on_cancel.
func (recv *Task) SetReturnOnCancel(returnOnCancel bool) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetBoolean(returnOnCancel)

	var ret gi.Argument

	err := taskSetReturnOnCancelFunction_Set()
	if err == nil {
		ret = taskSetReturnOnCancelFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'g_task_set_source_tag' : parameter 'source_tag' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_task_set_task_data' : parameter 'task_data' of type 'gpointer' not supported

var tcpConnectionObject *gi.Object
var tcpConnectionObject_Once sync.Once

func tcpConnectionObject_Set() error {
	var err error
	tcpConnectionObject_Once.Do(func() {
		tcpConnectionObject, err = gi.ObjectNew("Gio", "TcpConnection")
	})
	return err
}

type TcpConnection struct {
	native unsafe.Pointer
}

func TcpConnectionNewFromNative(native unsafe.Pointer) *TcpConnection {
	instance := &TcpConnection{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// SocketConnection upcasts to *SocketConnection
func (recv *TcpConnection) SocketConnection() *SocketConnection {
	return SocketConnectionNewFromNative(recv.Native())
}

// IOStream upcasts to *IOStream
func (recv *TcpConnection) IOStream() *IOStream {
	return IOStreamNewFromNative(recv.Native())
}

// Object upcasts to *Object
func (recv *TcpConnection) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToTcpConnection down casts any arbitrary Object to TcpConnection.
Exercise care, as this is a potentially dangerous function
if the Object is not a TcpConnection.
*/
func CastToTcpConnection(object *gobject.Object) *TcpConnection {
	return TcpConnectionNewFromNative(object.Native())
}

// Equals compares this TcpConnection with another TcpConnection, and returns true if they represent the same GObject.
func (recv *TcpConnection) Equals(other *TcpConnection) bool {
	return other.Native() == recv.Native()
}

func (recv *TcpConnection) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *TcpConnection) FieldParentInstance() *SocketConnection {
	argValue := gi.ObjectFieldGet(tcpConnectionObject, recv.Native(), "parent_instance")
	value := SocketConnectionNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *TcpConnection) SetFieldParentInstance(value *SocketConnection) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(tcpConnectionObject, recv.Native(), "parent_instance", argValue)
}

// FieldPriv returns the C field 'priv'.
func (recv *TcpConnection) FieldPriv() *TcpConnectionPrivate {
	argValue := gi.ObjectFieldGet(tcpConnectionObject, recv.Native(), "priv")
	value := TcpConnectionPrivateNewFromNative(argValue.Pointer())
	return value
}

// SetFieldPriv sets the value of the C field 'priv'.
func (recv *TcpConnection) SetFieldPriv(value *TcpConnectionPrivate) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(tcpConnectionObject, recv.Native(), "priv", argValue)
}

var tcpConnectionGetGracefulDisconnectFunction *gi.Function
var tcpConnectionGetGracefulDisconnectFunction_Once sync.Once

func tcpConnectionGetGracefulDisconnectFunction_Set() error {
	var err error
	tcpConnectionGetGracefulDisconnectFunction_Once.Do(func() {
		err = tcpConnectionObject_Set()
		if err != nil {
			return
		}
		tcpConnectionGetGracefulDisconnectFunction, err = tcpConnectionObject.InvokerNew("get_graceful_disconnect")
	})
	return err
}

// GetGracefulDisconnect is a representation of the C type g_tcp_connection_get_graceful_disconnect.
func (recv *TcpConnection) GetGracefulDisconnect() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := tcpConnectionGetGracefulDisconnectFunction_Set()
	if err == nil {
		ret = tcpConnectionGetGracefulDisconnectFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var tcpConnectionSetGracefulDisconnectFunction *gi.Function
var tcpConnectionSetGracefulDisconnectFunction_Once sync.Once

func tcpConnectionSetGracefulDisconnectFunction_Set() error {
	var err error
	tcpConnectionSetGracefulDisconnectFunction_Once.Do(func() {
		err = tcpConnectionObject_Set()
		if err != nil {
			return
		}
		tcpConnectionSetGracefulDisconnectFunction, err = tcpConnectionObject.InvokerNew("set_graceful_disconnect")
	})
	return err
}

// SetGracefulDisconnect is a representation of the C type g_tcp_connection_set_graceful_disconnect.
func (recv *TcpConnection) SetGracefulDisconnect(gracefulDisconnect bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetBoolean(gracefulDisconnect)

	err := tcpConnectionSetGracefulDisconnectFunction_Set()
	if err == nil {
		tcpConnectionSetGracefulDisconnectFunction.Invoke(inArgs[:], nil)
	}

	return
}

var tcpWrapperConnectionObject *gi.Object
var tcpWrapperConnectionObject_Once sync.Once

func tcpWrapperConnectionObject_Set() error {
	var err error
	tcpWrapperConnectionObject_Once.Do(func() {
		tcpWrapperConnectionObject, err = gi.ObjectNew("Gio", "TcpWrapperConnection")
	})
	return err
}

type TcpWrapperConnection struct {
	native unsafe.Pointer
}

func TcpWrapperConnectionNewFromNative(native unsafe.Pointer) *TcpWrapperConnection {
	instance := &TcpWrapperConnection{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// TcpConnection upcasts to *TcpConnection
func (recv *TcpWrapperConnection) TcpConnection() *TcpConnection {
	return TcpConnectionNewFromNative(recv.Native())
}

// SocketConnection upcasts to *SocketConnection
func (recv *TcpWrapperConnection) SocketConnection() *SocketConnection {
	return SocketConnectionNewFromNative(recv.Native())
}

// IOStream upcasts to *IOStream
func (recv *TcpWrapperConnection) IOStream() *IOStream {
	return IOStreamNewFromNative(recv.Native())
}

// Object upcasts to *Object
func (recv *TcpWrapperConnection) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToTcpWrapperConnection down casts any arbitrary Object to TcpWrapperConnection.
Exercise care, as this is a potentially dangerous function
if the Object is not a TcpWrapperConnection.
*/
func CastToTcpWrapperConnection(object *gobject.Object) *TcpWrapperConnection {
	return TcpWrapperConnectionNewFromNative(object.Native())
}

// Equals compares this TcpWrapperConnection with another TcpWrapperConnection, and returns true if they represent the same GObject.
func (recv *TcpWrapperConnection) Equals(other *TcpWrapperConnection) bool {
	return other.Native() == recv.Native()
}

func (recv *TcpWrapperConnection) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *TcpWrapperConnection) FieldParentInstance() *TcpConnection {
	argValue := gi.ObjectFieldGet(tcpWrapperConnectionObject, recv.Native(), "parent_instance")
	value := TcpConnectionNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *TcpWrapperConnection) SetFieldParentInstance(value *TcpConnection) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(tcpWrapperConnectionObject, recv.Native(), "parent_instance", argValue)
}

// FieldPriv returns the C field 'priv'.
func (recv *TcpWrapperConnection) FieldPriv() *TcpWrapperConnectionPrivate {
	argValue := gi.ObjectFieldGet(tcpWrapperConnectionObject, recv.Native(), "priv")
	value := TcpWrapperConnectionPrivateNewFromNative(argValue.Pointer())
	return value
}

// SetFieldPriv sets the value of the C field 'priv'.
func (recv *TcpWrapperConnection) SetFieldPriv(value *TcpWrapperConnectionPrivate) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(tcpWrapperConnectionObject, recv.Native(), "priv", argValue)
}

var tcpWrapperConnectionNewFunction *gi.Function
var tcpWrapperConnectionNewFunction_Once sync.Once

func tcpWrapperConnectionNewFunction_Set() error {
	var err error
	tcpWrapperConnectionNewFunction_Once.Do(func() {
		err = tcpWrapperConnectionObject_Set()
		if err != nil {
			return
		}
		tcpWrapperConnectionNewFunction, err = tcpWrapperConnectionObject.InvokerNew("new")
	})
	return err
}

// TcpWrapperConnectionNew is a representation of the C type g_tcp_wrapper_connection_new.
func TcpWrapperConnectionNew(baseIoStream *IOStream, socket *Socket) *TcpWrapperConnection {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(baseIoStream.Native())
	inArgs[1].SetPointer(socket.Native())

	var ret gi.Argument

	err := tcpWrapperConnectionNewFunction_Set()
	if err == nil {
		ret = tcpWrapperConnectionNewFunction.Invoke(inArgs[:], nil)
	}

	retGo := TcpWrapperConnectionNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

var tcpWrapperConnectionGetBaseIoStreamFunction *gi.Function
var tcpWrapperConnectionGetBaseIoStreamFunction_Once sync.Once

func tcpWrapperConnectionGetBaseIoStreamFunction_Set() error {
	var err error
	tcpWrapperConnectionGetBaseIoStreamFunction_Once.Do(func() {
		err = tcpWrapperConnectionObject_Set()
		if err != nil {
			return
		}
		tcpWrapperConnectionGetBaseIoStreamFunction, err = tcpWrapperConnectionObject.InvokerNew("get_base_io_stream")
	})
	return err
}

// GetBaseIoStream is a representation of the C type g_tcp_wrapper_connection_get_base_io_stream.
func (recv *TcpWrapperConnection) GetBaseIoStream() *IOStream {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := tcpWrapperConnectionGetBaseIoStreamFunction_Set()
	if err == nil {
		ret = tcpWrapperConnectionGetBaseIoStreamFunction.Invoke(inArgs[:], nil)
	}

	retGo := IOStreamNewFromNative(ret.Pointer())

	return retGo
}

var testDBusObject *gi.Object
var testDBusObject_Once sync.Once

func testDBusObject_Set() error {
	var err error
	testDBusObject_Once.Do(func() {
		testDBusObject, err = gi.ObjectNew("Gio", "TestDBus")
	})
	return err
}

type TestDBus struct {
	native unsafe.Pointer
}

func TestDBusNewFromNative(native unsafe.Pointer) *TestDBus {
	instance := &TestDBus{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *TestDBus) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToTestDBus down casts any arbitrary Object to TestDBus.
Exercise care, as this is a potentially dangerous function
if the Object is not a TestDBus.
*/
func CastToTestDBus(object *gobject.Object) *TestDBus {
	return TestDBusNewFromNative(object.Native())
}

// Equals compares this TestDBus with another TestDBus, and returns true if they represent the same GObject.
func (recv *TestDBus) Equals(other *TestDBus) bool {
	return other.Native() == recv.Native()
}

func (recv *TestDBus) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'g_test_dbus_new' : parameter 'flags' of type 'TestDBusFlags' not supported

var testDBusAddServiceDirFunction *gi.Function
var testDBusAddServiceDirFunction_Once sync.Once

func testDBusAddServiceDirFunction_Set() error {
	var err error
	testDBusAddServiceDirFunction_Once.Do(func() {
		err = testDBusObject_Set()
		if err != nil {
			return
		}
		testDBusAddServiceDirFunction, err = testDBusObject.InvokerNew("add_service_dir")
	})
	return err
}

// AddServiceDir is a representation of the C type g_test_dbus_add_service_dir.
func (recv *TestDBus) AddServiceDir(path string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(path)

	err := testDBusAddServiceDirFunction_Set()
	if err == nil {
		testDBusAddServiceDirFunction.Invoke(inArgs[:], nil)
	}

	return
}

var testDBusDownFunction *gi.Function
var testDBusDownFunction_Once sync.Once

func testDBusDownFunction_Set() error {
	var err error
	testDBusDownFunction_Once.Do(func() {
		err = testDBusObject_Set()
		if err != nil {
			return
		}
		testDBusDownFunction, err = testDBusObject.InvokerNew("down")
	})
	return err
}

// Down is a representation of the C type g_test_dbus_down.
func (recv *TestDBus) Down() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := testDBusDownFunction_Set()
	if err == nil {
		testDBusDownFunction.Invoke(inArgs[:], nil)
	}

	return
}

var testDBusGetBusAddressFunction *gi.Function
var testDBusGetBusAddressFunction_Once sync.Once

func testDBusGetBusAddressFunction_Set() error {
	var err error
	testDBusGetBusAddressFunction_Once.Do(func() {
		err = testDBusObject_Set()
		if err != nil {
			return
		}
		testDBusGetBusAddressFunction, err = testDBusObject.InvokerNew("get_bus_address")
	})
	return err
}

// GetBusAddress is a representation of the C type g_test_dbus_get_bus_address.
func (recv *TestDBus) GetBusAddress() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := testDBusGetBusAddressFunction_Set()
	if err == nil {
		ret = testDBusGetBusAddressFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'g_test_dbus_get_flags' : return type 'TestDBusFlags' not supported

var testDBusStopFunction *gi.Function
var testDBusStopFunction_Once sync.Once

func testDBusStopFunction_Set() error {
	var err error
	testDBusStopFunction_Once.Do(func() {
		err = testDBusObject_Set()
		if err != nil {
			return
		}
		testDBusStopFunction, err = testDBusObject.InvokerNew("stop")
	})
	return err
}

// Stop is a representation of the C type g_test_dbus_stop.
func (recv *TestDBus) Stop() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := testDBusStopFunction_Set()
	if err == nil {
		testDBusStopFunction.Invoke(inArgs[:], nil)
	}

	return
}

var testDBusUpFunction *gi.Function
var testDBusUpFunction_Once sync.Once

func testDBusUpFunction_Set() error {
	var err error
	testDBusUpFunction_Once.Do(func() {
		err = testDBusObject_Set()
		if err != nil {
			return
		}
		testDBusUpFunction, err = testDBusObject.InvokerNew("up")
	})
	return err
}

// Up is a representation of the C type g_test_dbus_up.
func (recv *TestDBus) Up() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := testDBusUpFunction_Set()
	if err == nil {
		testDBusUpFunction.Invoke(inArgs[:], nil)
	}

	return
}

var themedIconObject *gi.Object
var themedIconObject_Once sync.Once

func themedIconObject_Set() error {
	var err error
	themedIconObject_Once.Do(func() {
		themedIconObject, err = gi.ObjectNew("Gio", "ThemedIcon")
	})
	return err
}

type ThemedIcon struct {
	native unsafe.Pointer
}

func ThemedIconNewFromNative(native unsafe.Pointer) *ThemedIcon {
	instance := &ThemedIcon{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *ThemedIcon) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToThemedIcon down casts any arbitrary Object to ThemedIcon.
Exercise care, as this is a potentially dangerous function
if the Object is not a ThemedIcon.
*/
func CastToThemedIcon(object *gobject.Object) *ThemedIcon {
	return ThemedIconNewFromNative(object.Native())
}

// Equals compares this ThemedIcon with another ThemedIcon, and returns true if they represent the same GObject.
func (recv *ThemedIcon) Equals(other *ThemedIcon) bool {
	return other.Native() == recv.Native()
}

func (recv *ThemedIcon) Native() unsafe.Pointer {
	return recv.native
}

var themedIconNewFunction *gi.Function
var themedIconNewFunction_Once sync.Once

func themedIconNewFunction_Set() error {
	var err error
	themedIconNewFunction_Once.Do(func() {
		err = themedIconObject_Set()
		if err != nil {
			return
		}
		themedIconNewFunction, err = themedIconObject.InvokerNew("new")
	})
	return err
}

// ThemedIconNew is a representation of the C type g_themed_icon_new.
func ThemedIconNew(iconname string) *ThemedIcon {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(iconname)

	var ret gi.Argument

	err := themedIconNewFunction_Set()
	if err == nil {
		ret = themedIconNewFunction.Invoke(inArgs[:], nil)
	}

	retGo := ThemedIconNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

// UNSUPPORTED : C value 'g_themed_icon_new_from_names' : parameter 'iconnames' of type 'nil' not supported

var themedIconNewWithDefaultFallbacksFunction *gi.Function
var themedIconNewWithDefaultFallbacksFunction_Once sync.Once

func themedIconNewWithDefaultFallbacksFunction_Set() error {
	var err error
	themedIconNewWithDefaultFallbacksFunction_Once.Do(func() {
		err = themedIconObject_Set()
		if err != nil {
			return
		}
		themedIconNewWithDefaultFallbacksFunction, err = themedIconObject.InvokerNew("new_with_default_fallbacks")
	})
	return err
}

// ThemedIconNewWithDefaultFallbacks is a representation of the C type g_themed_icon_new_with_default_fallbacks.
func ThemedIconNewWithDefaultFallbacks(iconname string) *ThemedIcon {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(iconname)

	var ret gi.Argument

	err := themedIconNewWithDefaultFallbacksFunction_Set()
	if err == nil {
		ret = themedIconNewWithDefaultFallbacksFunction.Invoke(inArgs[:], nil)
	}

	retGo := ThemedIconNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

var themedIconAppendNameFunction *gi.Function
var themedIconAppendNameFunction_Once sync.Once

func themedIconAppendNameFunction_Set() error {
	var err error
	themedIconAppendNameFunction_Once.Do(func() {
		err = themedIconObject_Set()
		if err != nil {
			return
		}
		themedIconAppendNameFunction, err = themedIconObject.InvokerNew("append_name")
	})
	return err
}

// AppendName is a representation of the C type g_themed_icon_append_name.
func (recv *ThemedIcon) AppendName(iconname string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(iconname)

	err := themedIconAppendNameFunction_Set()
	if err == nil {
		themedIconAppendNameFunction.Invoke(inArgs[:], nil)
	}

	return
}

var themedIconGetNamesFunction *gi.Function
var themedIconGetNamesFunction_Once sync.Once

func themedIconGetNamesFunction_Set() error {
	var err error
	themedIconGetNamesFunction_Once.Do(func() {
		err = themedIconObject_Set()
		if err != nil {
			return
		}
		themedIconGetNamesFunction, err = themedIconObject.InvokerNew("get_names")
	})
	return err
}

// GetNames is a representation of the C type g_themed_icon_get_names.
func (recv *ThemedIcon) GetNames() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := themedIconGetNamesFunction_Set()
	if err == nil {
		themedIconGetNamesFunction.Invoke(inArgs[:], nil)
	}

	return
}

var themedIconPrependNameFunction *gi.Function
var themedIconPrependNameFunction_Once sync.Once

func themedIconPrependNameFunction_Set() error {
	var err error
	themedIconPrependNameFunction_Once.Do(func() {
		err = themedIconObject_Set()
		if err != nil {
			return
		}
		themedIconPrependNameFunction, err = themedIconObject.InvokerNew("prepend_name")
	})
	return err
}

// PrependName is a representation of the C type g_themed_icon_prepend_name.
func (recv *ThemedIcon) PrependName(iconname string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(iconname)

	err := themedIconPrependNameFunction_Set()
	if err == nil {
		themedIconPrependNameFunction.Invoke(inArgs[:], nil)
	}

	return
}

var threadedSocketServiceObject *gi.Object
var threadedSocketServiceObject_Once sync.Once

func threadedSocketServiceObject_Set() error {
	var err error
	threadedSocketServiceObject_Once.Do(func() {
		threadedSocketServiceObject, err = gi.ObjectNew("Gio", "ThreadedSocketService")
	})
	return err
}

type ThreadedSocketService struct {
	native unsafe.Pointer
}

func ThreadedSocketServiceNewFromNative(native unsafe.Pointer) *ThreadedSocketService {
	instance := &ThreadedSocketService{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// SocketService upcasts to *SocketService
func (recv *ThreadedSocketService) SocketService() *SocketService {
	return SocketServiceNewFromNative(recv.Native())
}

// SocketListener upcasts to *SocketListener
func (recv *ThreadedSocketService) SocketListener() *SocketListener {
	return SocketListenerNewFromNative(recv.Native())
}

// Object upcasts to *Object
func (recv *ThreadedSocketService) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToThreadedSocketService down casts any arbitrary Object to ThreadedSocketService.
Exercise care, as this is a potentially dangerous function
if the Object is not a ThreadedSocketService.
*/
func CastToThreadedSocketService(object *gobject.Object) *ThreadedSocketService {
	return ThreadedSocketServiceNewFromNative(object.Native())
}

// Equals compares this ThreadedSocketService with another ThreadedSocketService, and returns true if they represent the same GObject.
func (recv *ThreadedSocketService) Equals(other *ThreadedSocketService) bool {
	return other.Native() == recv.Native()
}

func (recv *ThreadedSocketService) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *ThreadedSocketService) FieldParentInstance() *SocketService {
	argValue := gi.ObjectFieldGet(threadedSocketServiceObject, recv.Native(), "parent_instance")
	value := SocketServiceNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *ThreadedSocketService) SetFieldParentInstance(value *SocketService) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(threadedSocketServiceObject, recv.Native(), "parent_instance", argValue)
}

// FieldPriv returns the C field 'priv'.
func (recv *ThreadedSocketService) FieldPriv() *ThreadedSocketServicePrivate {
	argValue := gi.ObjectFieldGet(threadedSocketServiceObject, recv.Native(), "priv")
	value := ThreadedSocketServicePrivateNewFromNative(argValue.Pointer())
	return value
}

// SetFieldPriv sets the value of the C field 'priv'.
func (recv *ThreadedSocketService) SetFieldPriv(value *ThreadedSocketServicePrivate) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(threadedSocketServiceObject, recv.Native(), "priv", argValue)
}

var threadedSocketServiceNewFunction *gi.Function
var threadedSocketServiceNewFunction_Once sync.Once

func threadedSocketServiceNewFunction_Set() error {
	var err error
	threadedSocketServiceNewFunction_Once.Do(func() {
		err = threadedSocketServiceObject_Set()
		if err != nil {
			return
		}
		threadedSocketServiceNewFunction, err = threadedSocketServiceObject.InvokerNew("new")
	})
	return err
}

// ThreadedSocketServiceNew is a representation of the C type g_threaded_socket_service_new.
func ThreadedSocketServiceNew(maxThreads int32) *ThreadedSocketService {
	var inArgs [1]gi.Argument
	inArgs[0].SetInt32(maxThreads)

	var ret gi.Argument

	err := threadedSocketServiceNewFunction_Set()
	if err == nil {
		ret = threadedSocketServiceNewFunction.Invoke(inArgs[:], nil)
	}

	retGo := ThreadedSocketServiceNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

var tlsCertificateObject *gi.Object
var tlsCertificateObject_Once sync.Once

func tlsCertificateObject_Set() error {
	var err error
	tlsCertificateObject_Once.Do(func() {
		tlsCertificateObject, err = gi.ObjectNew("Gio", "TlsCertificate")
	})
	return err
}

type TlsCertificate struct {
	native unsafe.Pointer
}

func TlsCertificateNewFromNative(native unsafe.Pointer) *TlsCertificate {
	instance := &TlsCertificate{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *TlsCertificate) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToTlsCertificate down casts any arbitrary Object to TlsCertificate.
Exercise care, as this is a potentially dangerous function
if the Object is not a TlsCertificate.
*/
func CastToTlsCertificate(object *gobject.Object) *TlsCertificate {
	return TlsCertificateNewFromNative(object.Native())
}

// Equals compares this TlsCertificate with another TlsCertificate, and returns true if they represent the same GObject.
func (recv *TlsCertificate) Equals(other *TlsCertificate) bool {
	return other.Native() == recv.Native()
}

func (recv *TlsCertificate) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *TlsCertificate) FieldParentInstance() *gobject.Object {
	argValue := gi.ObjectFieldGet(tlsCertificateObject, recv.Native(), "parent_instance")
	value := gobject.ObjectNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *TlsCertificate) SetFieldParentInstance(value *gobject.Object) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(tlsCertificateObject, recv.Native(), "parent_instance", argValue)
}

// FieldPriv returns the C field 'priv'.
func (recv *TlsCertificate) FieldPriv() *TlsCertificatePrivate {
	argValue := gi.ObjectFieldGet(tlsCertificateObject, recv.Native(), "priv")
	value := TlsCertificatePrivateNewFromNative(argValue.Pointer())
	return value
}

// SetFieldPriv sets the value of the C field 'priv'.
func (recv *TlsCertificate) SetFieldPriv(value *TlsCertificatePrivate) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(tlsCertificateObject, recv.Native(), "priv", argValue)
}

var tlsCertificateNewFromFileFunction *gi.Function
var tlsCertificateNewFromFileFunction_Once sync.Once

func tlsCertificateNewFromFileFunction_Set() error {
	var err error
	tlsCertificateNewFromFileFunction_Once.Do(func() {
		err = tlsCertificateObject_Set()
		if err != nil {
			return
		}
		tlsCertificateNewFromFileFunction, err = tlsCertificateObject.InvokerNew("new_from_file")
	})
	return err
}

// TlsCertificateNewFromFile is a representation of the C type g_tls_certificate_new_from_file.
func TlsCertificateNewFromFile(file string) *TlsCertificate {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(file)

	var ret gi.Argument

	err := tlsCertificateNewFromFileFunction_Set()
	if err == nil {
		ret = tlsCertificateNewFromFileFunction.Invoke(inArgs[:], nil)
	}

	retGo := TlsCertificateNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

var tlsCertificateNewFromFilesFunction *gi.Function
var tlsCertificateNewFromFilesFunction_Once sync.Once

func tlsCertificateNewFromFilesFunction_Set() error {
	var err error
	tlsCertificateNewFromFilesFunction_Once.Do(func() {
		err = tlsCertificateObject_Set()
		if err != nil {
			return
		}
		tlsCertificateNewFromFilesFunction, err = tlsCertificateObject.InvokerNew("new_from_files")
	})
	return err
}

// TlsCertificateNewFromFiles is a representation of the C type g_tls_certificate_new_from_files.
func TlsCertificateNewFromFiles(certFile string, keyFile string) *TlsCertificate {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(certFile)
	inArgs[1].SetString(keyFile)

	var ret gi.Argument

	err := tlsCertificateNewFromFilesFunction_Set()
	if err == nil {
		ret = tlsCertificateNewFromFilesFunction.Invoke(inArgs[:], nil)
	}

	retGo := TlsCertificateNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

var tlsCertificateNewFromPemFunction *gi.Function
var tlsCertificateNewFromPemFunction_Once sync.Once

func tlsCertificateNewFromPemFunction_Set() error {
	var err error
	tlsCertificateNewFromPemFunction_Once.Do(func() {
		err = tlsCertificateObject_Set()
		if err != nil {
			return
		}
		tlsCertificateNewFromPemFunction, err = tlsCertificateObject.InvokerNew("new_from_pem")
	})
	return err
}

// TlsCertificateNewFromPem is a representation of the C type g_tls_certificate_new_from_pem.
func TlsCertificateNewFromPem(data string, length int32) *TlsCertificate {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(data)
	inArgs[1].SetInt32(length)

	var ret gi.Argument

	err := tlsCertificateNewFromPemFunction_Set()
	if err == nil {
		ret = tlsCertificateNewFromPemFunction.Invoke(inArgs[:], nil)
	}

	retGo := TlsCertificateNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

var tlsCertificateGetIssuerFunction *gi.Function
var tlsCertificateGetIssuerFunction_Once sync.Once

func tlsCertificateGetIssuerFunction_Set() error {
	var err error
	tlsCertificateGetIssuerFunction_Once.Do(func() {
		err = tlsCertificateObject_Set()
		if err != nil {
			return
		}
		tlsCertificateGetIssuerFunction, err = tlsCertificateObject.InvokerNew("get_issuer")
	})
	return err
}

// GetIssuer is a representation of the C type g_tls_certificate_get_issuer.
func (recv *TlsCertificate) GetIssuer() *TlsCertificate {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := tlsCertificateGetIssuerFunction_Set()
	if err == nil {
		ret = tlsCertificateGetIssuerFunction.Invoke(inArgs[:], nil)
	}

	retGo := TlsCertificateNewFromNative(ret.Pointer())

	return retGo
}

var tlsCertificateIsSameFunction *gi.Function
var tlsCertificateIsSameFunction_Once sync.Once

func tlsCertificateIsSameFunction_Set() error {
	var err error
	tlsCertificateIsSameFunction_Once.Do(func() {
		err = tlsCertificateObject_Set()
		if err != nil {
			return
		}
		tlsCertificateIsSameFunction, err = tlsCertificateObject.InvokerNew("is_same")
	})
	return err
}

// IsSame is a representation of the C type g_tls_certificate_is_same.
func (recv *TlsCertificate) IsSame(certTwo *TlsCertificate) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(certTwo.Native())

	var ret gi.Argument

	err := tlsCertificateIsSameFunction_Set()
	if err == nil {
		ret = tlsCertificateIsSameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'g_tls_certificate_verify' : parameter 'identity' of type 'SocketConnectable' not supported

var tlsConnectionObject *gi.Object
var tlsConnectionObject_Once sync.Once

func tlsConnectionObject_Set() error {
	var err error
	tlsConnectionObject_Once.Do(func() {
		tlsConnectionObject, err = gi.ObjectNew("Gio", "TlsConnection")
	})
	return err
}

type TlsConnection struct {
	native unsafe.Pointer
}

func TlsConnectionNewFromNative(native unsafe.Pointer) *TlsConnection {
	instance := &TlsConnection{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// IOStream upcasts to *IOStream
func (recv *TlsConnection) IOStream() *IOStream {
	return IOStreamNewFromNative(recv.Native())
}

// Object upcasts to *Object
func (recv *TlsConnection) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToTlsConnection down casts any arbitrary Object to TlsConnection.
Exercise care, as this is a potentially dangerous function
if the Object is not a TlsConnection.
*/
func CastToTlsConnection(object *gobject.Object) *TlsConnection {
	return TlsConnectionNewFromNative(object.Native())
}

// Equals compares this TlsConnection with another TlsConnection, and returns true if they represent the same GObject.
func (recv *TlsConnection) Equals(other *TlsConnection) bool {
	return other.Native() == recv.Native()
}

func (recv *TlsConnection) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *TlsConnection) FieldParentInstance() *IOStream {
	argValue := gi.ObjectFieldGet(tlsConnectionObject, recv.Native(), "parent_instance")
	value := IOStreamNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *TlsConnection) SetFieldParentInstance(value *IOStream) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(tlsConnectionObject, recv.Native(), "parent_instance", argValue)
}

// FieldPriv returns the C field 'priv'.
func (recv *TlsConnection) FieldPriv() *TlsConnectionPrivate {
	argValue := gi.ObjectFieldGet(tlsConnectionObject, recv.Native(), "priv")
	value := TlsConnectionPrivateNewFromNative(argValue.Pointer())
	return value
}

// SetFieldPriv sets the value of the C field 'priv'.
func (recv *TlsConnection) SetFieldPriv(value *TlsConnectionPrivate) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(tlsConnectionObject, recv.Native(), "priv", argValue)
}

// UNSUPPORTED : C value 'g_tls_connection_emit_accept_certificate' : parameter 'errors' of type 'TlsCertificateFlags' not supported

var tlsConnectionGetCertificateFunction *gi.Function
var tlsConnectionGetCertificateFunction_Once sync.Once

func tlsConnectionGetCertificateFunction_Set() error {
	var err error
	tlsConnectionGetCertificateFunction_Once.Do(func() {
		err = tlsConnectionObject_Set()
		if err != nil {
			return
		}
		tlsConnectionGetCertificateFunction, err = tlsConnectionObject.InvokerNew("get_certificate")
	})
	return err
}

// GetCertificate is a representation of the C type g_tls_connection_get_certificate.
func (recv *TlsConnection) GetCertificate() *TlsCertificate {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := tlsConnectionGetCertificateFunction_Set()
	if err == nil {
		ret = tlsConnectionGetCertificateFunction.Invoke(inArgs[:], nil)
	}

	retGo := TlsCertificateNewFromNative(ret.Pointer())

	return retGo
}

var tlsConnectionGetDatabaseFunction *gi.Function
var tlsConnectionGetDatabaseFunction_Once sync.Once

func tlsConnectionGetDatabaseFunction_Set() error {
	var err error
	tlsConnectionGetDatabaseFunction_Once.Do(func() {
		err = tlsConnectionObject_Set()
		if err != nil {
			return
		}
		tlsConnectionGetDatabaseFunction, err = tlsConnectionObject.InvokerNew("get_database")
	})
	return err
}

// GetDatabase is a representation of the C type g_tls_connection_get_database.
func (recv *TlsConnection) GetDatabase() *TlsDatabase {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := tlsConnectionGetDatabaseFunction_Set()
	if err == nil {
		ret = tlsConnectionGetDatabaseFunction.Invoke(inArgs[:], nil)
	}

	retGo := TlsDatabaseNewFromNative(ret.Pointer())

	return retGo
}

var tlsConnectionGetInteractionFunction *gi.Function
var tlsConnectionGetInteractionFunction_Once sync.Once

func tlsConnectionGetInteractionFunction_Set() error {
	var err error
	tlsConnectionGetInteractionFunction_Once.Do(func() {
		err = tlsConnectionObject_Set()
		if err != nil {
			return
		}
		tlsConnectionGetInteractionFunction, err = tlsConnectionObject.InvokerNew("get_interaction")
	})
	return err
}

// GetInteraction is a representation of the C type g_tls_connection_get_interaction.
func (recv *TlsConnection) GetInteraction() *TlsInteraction {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := tlsConnectionGetInteractionFunction_Set()
	if err == nil {
		ret = tlsConnectionGetInteractionFunction.Invoke(inArgs[:], nil)
	}

	retGo := TlsInteractionNewFromNative(ret.Pointer())

	return retGo
}

var tlsConnectionGetNegotiatedProtocolFunction *gi.Function
var tlsConnectionGetNegotiatedProtocolFunction_Once sync.Once

func tlsConnectionGetNegotiatedProtocolFunction_Set() error {
	var err error
	tlsConnectionGetNegotiatedProtocolFunction_Once.Do(func() {
		err = tlsConnectionObject_Set()
		if err != nil {
			return
		}
		tlsConnectionGetNegotiatedProtocolFunction, err = tlsConnectionObject.InvokerNew("get_negotiated_protocol")
	})
	return err
}

// GetNegotiatedProtocol is a representation of the C type g_tls_connection_get_negotiated_protocol.
func (recv *TlsConnection) GetNegotiatedProtocol() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := tlsConnectionGetNegotiatedProtocolFunction_Set()
	if err == nil {
		ret = tlsConnectionGetNegotiatedProtocolFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var tlsConnectionGetPeerCertificateFunction *gi.Function
var tlsConnectionGetPeerCertificateFunction_Once sync.Once

func tlsConnectionGetPeerCertificateFunction_Set() error {
	var err error
	tlsConnectionGetPeerCertificateFunction_Once.Do(func() {
		err = tlsConnectionObject_Set()
		if err != nil {
			return
		}
		tlsConnectionGetPeerCertificateFunction, err = tlsConnectionObject.InvokerNew("get_peer_certificate")
	})
	return err
}

// GetPeerCertificate is a representation of the C type g_tls_connection_get_peer_certificate.
func (recv *TlsConnection) GetPeerCertificate() *TlsCertificate {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := tlsConnectionGetPeerCertificateFunction_Set()
	if err == nil {
		ret = tlsConnectionGetPeerCertificateFunction.Invoke(inArgs[:], nil)
	}

	retGo := TlsCertificateNewFromNative(ret.Pointer())

	return retGo
}

// UNSUPPORTED : C value 'g_tls_connection_get_peer_certificate_errors' : return type 'TlsCertificateFlags' not supported

var tlsConnectionGetRehandshakeModeFunction *gi.Function
var tlsConnectionGetRehandshakeModeFunction_Once sync.Once

func tlsConnectionGetRehandshakeModeFunction_Set() error {
	var err error
	tlsConnectionGetRehandshakeModeFunction_Once.Do(func() {
		err = tlsConnectionObject_Set()
		if err != nil {
			return
		}
		tlsConnectionGetRehandshakeModeFunction, err = tlsConnectionObject.InvokerNew("get_rehandshake_mode")
	})
	return err
}

// GetRehandshakeMode is a representation of the C type g_tls_connection_get_rehandshake_mode.
func (recv *TlsConnection) GetRehandshakeMode() TlsRehandshakeMode {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := tlsConnectionGetRehandshakeModeFunction_Set()
	if err == nil {
		ret = tlsConnectionGetRehandshakeModeFunction.Invoke(inArgs[:], nil)
	}

	retGo := TlsRehandshakeMode(ret.Int32())

	return retGo
}

var tlsConnectionGetRequireCloseNotifyFunction *gi.Function
var tlsConnectionGetRequireCloseNotifyFunction_Once sync.Once

func tlsConnectionGetRequireCloseNotifyFunction_Set() error {
	var err error
	tlsConnectionGetRequireCloseNotifyFunction_Once.Do(func() {
		err = tlsConnectionObject_Set()
		if err != nil {
			return
		}
		tlsConnectionGetRequireCloseNotifyFunction, err = tlsConnectionObject.InvokerNew("get_require_close_notify")
	})
	return err
}

// GetRequireCloseNotify is a representation of the C type g_tls_connection_get_require_close_notify.
func (recv *TlsConnection) GetRequireCloseNotify() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := tlsConnectionGetRequireCloseNotifyFunction_Set()
	if err == nil {
		ret = tlsConnectionGetRequireCloseNotifyFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var tlsConnectionGetUseSystemCertdbFunction *gi.Function
var tlsConnectionGetUseSystemCertdbFunction_Once sync.Once

func tlsConnectionGetUseSystemCertdbFunction_Set() error {
	var err error
	tlsConnectionGetUseSystemCertdbFunction_Once.Do(func() {
		err = tlsConnectionObject_Set()
		if err != nil {
			return
		}
		tlsConnectionGetUseSystemCertdbFunction, err = tlsConnectionObject.InvokerNew("get_use_system_certdb")
	})
	return err
}

// GetUseSystemCertdb is a representation of the C type g_tls_connection_get_use_system_certdb.
func (recv *TlsConnection) GetUseSystemCertdb() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := tlsConnectionGetUseSystemCertdbFunction_Set()
	if err == nil {
		ret = tlsConnectionGetUseSystemCertdbFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var tlsConnectionHandshakeFunction *gi.Function
var tlsConnectionHandshakeFunction_Once sync.Once

func tlsConnectionHandshakeFunction_Set() error {
	var err error
	tlsConnectionHandshakeFunction_Once.Do(func() {
		err = tlsConnectionObject_Set()
		if err != nil {
			return
		}
		tlsConnectionHandshakeFunction, err = tlsConnectionObject.InvokerNew("handshake")
	})
	return err
}

// Handshake is a representation of the C type g_tls_connection_handshake.
func (recv *TlsConnection) Handshake(cancellable *Cancellable) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(cancellable.Native())

	var ret gi.Argument

	err := tlsConnectionHandshakeFunction_Set()
	if err == nil {
		ret = tlsConnectionHandshakeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'g_tls_connection_handshake_async' : parameter 'callback' of type 'AsyncReadyCallback' not supported

// UNSUPPORTED : C value 'g_tls_connection_handshake_finish' : parameter 'result' of type 'AsyncResult' not supported

// UNSUPPORTED : C value 'g_tls_connection_set_advertised_protocols' : parameter 'protocols' of type 'nil' not supported

var tlsConnectionSetCertificateFunction *gi.Function
var tlsConnectionSetCertificateFunction_Once sync.Once

func tlsConnectionSetCertificateFunction_Set() error {
	var err error
	tlsConnectionSetCertificateFunction_Once.Do(func() {
		err = tlsConnectionObject_Set()
		if err != nil {
			return
		}
		tlsConnectionSetCertificateFunction, err = tlsConnectionObject.InvokerNew("set_certificate")
	})
	return err
}

// SetCertificate is a representation of the C type g_tls_connection_set_certificate.
func (recv *TlsConnection) SetCertificate(certificate *TlsCertificate) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(certificate.Native())

	err := tlsConnectionSetCertificateFunction_Set()
	if err == nil {
		tlsConnectionSetCertificateFunction.Invoke(inArgs[:], nil)
	}

	return
}

var tlsConnectionSetDatabaseFunction *gi.Function
var tlsConnectionSetDatabaseFunction_Once sync.Once

func tlsConnectionSetDatabaseFunction_Set() error {
	var err error
	tlsConnectionSetDatabaseFunction_Once.Do(func() {
		err = tlsConnectionObject_Set()
		if err != nil {
			return
		}
		tlsConnectionSetDatabaseFunction, err = tlsConnectionObject.InvokerNew("set_database")
	})
	return err
}

// SetDatabase is a representation of the C type g_tls_connection_set_database.
func (recv *TlsConnection) SetDatabase(database *TlsDatabase) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(database.Native())

	err := tlsConnectionSetDatabaseFunction_Set()
	if err == nil {
		tlsConnectionSetDatabaseFunction.Invoke(inArgs[:], nil)
	}

	return
}

var tlsConnectionSetInteractionFunction *gi.Function
var tlsConnectionSetInteractionFunction_Once sync.Once

func tlsConnectionSetInteractionFunction_Set() error {
	var err error
	tlsConnectionSetInteractionFunction_Once.Do(func() {
		err = tlsConnectionObject_Set()
		if err != nil {
			return
		}
		tlsConnectionSetInteractionFunction, err = tlsConnectionObject.InvokerNew("set_interaction")
	})
	return err
}

// SetInteraction is a representation of the C type g_tls_connection_set_interaction.
func (recv *TlsConnection) SetInteraction(interaction *TlsInteraction) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(interaction.Native())

	err := tlsConnectionSetInteractionFunction_Set()
	if err == nil {
		tlsConnectionSetInteractionFunction.Invoke(inArgs[:], nil)
	}

	return
}

var tlsConnectionSetRehandshakeModeFunction *gi.Function
var tlsConnectionSetRehandshakeModeFunction_Once sync.Once

func tlsConnectionSetRehandshakeModeFunction_Set() error {
	var err error
	tlsConnectionSetRehandshakeModeFunction_Once.Do(func() {
		err = tlsConnectionObject_Set()
		if err != nil {
			return
		}
		tlsConnectionSetRehandshakeModeFunction, err = tlsConnectionObject.InvokerNew("set_rehandshake_mode")
	})
	return err
}

// SetRehandshakeMode is a representation of the C type g_tls_connection_set_rehandshake_mode.
func (recv *TlsConnection) SetRehandshakeMode(mode TlsRehandshakeMode) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(int32(mode))

	err := tlsConnectionSetRehandshakeModeFunction_Set()
	if err == nil {
		tlsConnectionSetRehandshakeModeFunction.Invoke(inArgs[:], nil)
	}

	return
}

var tlsConnectionSetRequireCloseNotifyFunction *gi.Function
var tlsConnectionSetRequireCloseNotifyFunction_Once sync.Once

func tlsConnectionSetRequireCloseNotifyFunction_Set() error {
	var err error
	tlsConnectionSetRequireCloseNotifyFunction_Once.Do(func() {
		err = tlsConnectionObject_Set()
		if err != nil {
			return
		}
		tlsConnectionSetRequireCloseNotifyFunction, err = tlsConnectionObject.InvokerNew("set_require_close_notify")
	})
	return err
}

// SetRequireCloseNotify is a representation of the C type g_tls_connection_set_require_close_notify.
func (recv *TlsConnection) SetRequireCloseNotify(requireCloseNotify bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetBoolean(requireCloseNotify)

	err := tlsConnectionSetRequireCloseNotifyFunction_Set()
	if err == nil {
		tlsConnectionSetRequireCloseNotifyFunction.Invoke(inArgs[:], nil)
	}

	return
}

var tlsConnectionSetUseSystemCertdbFunction *gi.Function
var tlsConnectionSetUseSystemCertdbFunction_Once sync.Once

func tlsConnectionSetUseSystemCertdbFunction_Set() error {
	var err error
	tlsConnectionSetUseSystemCertdbFunction_Once.Do(func() {
		err = tlsConnectionObject_Set()
		if err != nil {
			return
		}
		tlsConnectionSetUseSystemCertdbFunction, err = tlsConnectionObject.InvokerNew("set_use_system_certdb")
	})
	return err
}

// SetUseSystemCertdb is a representation of the C type g_tls_connection_set_use_system_certdb.
func (recv *TlsConnection) SetUseSystemCertdb(useSystemCertdb bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetBoolean(useSystemCertdb)

	err := tlsConnectionSetUseSystemCertdbFunction_Set()
	if err == nil {
		tlsConnectionSetUseSystemCertdbFunction.Invoke(inArgs[:], nil)
	}

	return
}

var tlsDatabaseObject *gi.Object
var tlsDatabaseObject_Once sync.Once

func tlsDatabaseObject_Set() error {
	var err error
	tlsDatabaseObject_Once.Do(func() {
		tlsDatabaseObject, err = gi.ObjectNew("Gio", "TlsDatabase")
	})
	return err
}

type TlsDatabase struct {
	native unsafe.Pointer
}

func TlsDatabaseNewFromNative(native unsafe.Pointer) *TlsDatabase {
	instance := &TlsDatabase{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *TlsDatabase) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToTlsDatabase down casts any arbitrary Object to TlsDatabase.
Exercise care, as this is a potentially dangerous function
if the Object is not a TlsDatabase.
*/
func CastToTlsDatabase(object *gobject.Object) *TlsDatabase {
	return TlsDatabaseNewFromNative(object.Native())
}

// Equals compares this TlsDatabase with another TlsDatabase, and returns true if they represent the same GObject.
func (recv *TlsDatabase) Equals(other *TlsDatabase) bool {
	return other.Native() == recv.Native()
}

func (recv *TlsDatabase) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *TlsDatabase) FieldParentInstance() *gobject.Object {
	argValue := gi.ObjectFieldGet(tlsDatabaseObject, recv.Native(), "parent_instance")
	value := gobject.ObjectNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *TlsDatabase) SetFieldParentInstance(value *gobject.Object) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(tlsDatabaseObject, recv.Native(), "parent_instance", argValue)
}

// FieldPriv returns the C field 'priv'.
func (recv *TlsDatabase) FieldPriv() *TlsDatabasePrivate {
	argValue := gi.ObjectFieldGet(tlsDatabaseObject, recv.Native(), "priv")
	value := TlsDatabasePrivateNewFromNative(argValue.Pointer())
	return value
}

// SetFieldPriv sets the value of the C field 'priv'.
func (recv *TlsDatabase) SetFieldPriv(value *TlsDatabasePrivate) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(tlsDatabaseObject, recv.Native(), "priv", argValue)
}

var tlsDatabaseCreateCertificateHandleFunction *gi.Function
var tlsDatabaseCreateCertificateHandleFunction_Once sync.Once

func tlsDatabaseCreateCertificateHandleFunction_Set() error {
	var err error
	tlsDatabaseCreateCertificateHandleFunction_Once.Do(func() {
		err = tlsDatabaseObject_Set()
		if err != nil {
			return
		}
		tlsDatabaseCreateCertificateHandleFunction, err = tlsDatabaseObject.InvokerNew("create_certificate_handle")
	})
	return err
}

// CreateCertificateHandle is a representation of the C type g_tls_database_create_certificate_handle.
func (recv *TlsDatabase) CreateCertificateHandle(certificate *TlsCertificate) string {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(certificate.Native())

	var ret gi.Argument

	err := tlsDatabaseCreateCertificateHandleFunction_Set()
	if err == nil {
		ret = tlsDatabaseCreateCertificateHandleFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

var tlsDatabaseLookupCertificateForHandleFunction *gi.Function
var tlsDatabaseLookupCertificateForHandleFunction_Once sync.Once

func tlsDatabaseLookupCertificateForHandleFunction_Set() error {
	var err error
	tlsDatabaseLookupCertificateForHandleFunction_Once.Do(func() {
		err = tlsDatabaseObject_Set()
		if err != nil {
			return
		}
		tlsDatabaseLookupCertificateForHandleFunction, err = tlsDatabaseObject.InvokerNew("lookup_certificate_for_handle")
	})
	return err
}

// LookupCertificateForHandle is a representation of the C type g_tls_database_lookup_certificate_for_handle.
func (recv *TlsDatabase) LookupCertificateForHandle(handle string, interaction *TlsInteraction, flags TlsDatabaseLookupFlags, cancellable *Cancellable) *TlsCertificate {
	var inArgs [5]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(handle)
	inArgs[2].SetPointer(interaction.Native())
	inArgs[3].SetInt32(int32(flags))
	inArgs[4].SetPointer(cancellable.Native())

	var ret gi.Argument

	err := tlsDatabaseLookupCertificateForHandleFunction_Set()
	if err == nil {
		ret = tlsDatabaseLookupCertificateForHandleFunction.Invoke(inArgs[:], nil)
	}

	retGo := TlsCertificateNewFromNative(ret.Pointer())

	return retGo
}

// UNSUPPORTED : C value 'g_tls_database_lookup_certificate_for_handle_async' : parameter 'callback' of type 'AsyncReadyCallback' not supported

// UNSUPPORTED : C value 'g_tls_database_lookup_certificate_for_handle_finish' : parameter 'result' of type 'AsyncResult' not supported

var tlsDatabaseLookupCertificateIssuerFunction *gi.Function
var tlsDatabaseLookupCertificateIssuerFunction_Once sync.Once

func tlsDatabaseLookupCertificateIssuerFunction_Set() error {
	var err error
	tlsDatabaseLookupCertificateIssuerFunction_Once.Do(func() {
		err = tlsDatabaseObject_Set()
		if err != nil {
			return
		}
		tlsDatabaseLookupCertificateIssuerFunction, err = tlsDatabaseObject.InvokerNew("lookup_certificate_issuer")
	})
	return err
}

// LookupCertificateIssuer is a representation of the C type g_tls_database_lookup_certificate_issuer.
func (recv *TlsDatabase) LookupCertificateIssuer(certificate *TlsCertificate, interaction *TlsInteraction, flags TlsDatabaseLookupFlags, cancellable *Cancellable) *TlsCertificate {
	var inArgs [5]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(certificate.Native())
	inArgs[2].SetPointer(interaction.Native())
	inArgs[3].SetInt32(int32(flags))
	inArgs[4].SetPointer(cancellable.Native())

	var ret gi.Argument

	err := tlsDatabaseLookupCertificateIssuerFunction_Set()
	if err == nil {
		ret = tlsDatabaseLookupCertificateIssuerFunction.Invoke(inArgs[:], nil)
	}

	retGo := TlsCertificateNewFromNative(ret.Pointer())

	return retGo
}

// UNSUPPORTED : C value 'g_tls_database_lookup_certificate_issuer_async' : parameter 'callback' of type 'AsyncReadyCallback' not supported

// UNSUPPORTED : C value 'g_tls_database_lookup_certificate_issuer_finish' : parameter 'result' of type 'AsyncResult' not supported

// UNSUPPORTED : C value 'g_tls_database_lookup_certificates_issued_by' : parameter 'issuer_raw_dn' of type 'nil' not supported

// UNSUPPORTED : C value 'g_tls_database_lookup_certificates_issued_by_async' : parameter 'issuer_raw_dn' of type 'nil' not supported

// UNSUPPORTED : C value 'g_tls_database_lookup_certificates_issued_by_finish' : parameter 'result' of type 'AsyncResult' not supported

// UNSUPPORTED : C value 'g_tls_database_verify_chain' : parameter 'identity' of type 'SocketConnectable' not supported

// UNSUPPORTED : C value 'g_tls_database_verify_chain_async' : parameter 'identity' of type 'SocketConnectable' not supported

// UNSUPPORTED : C value 'g_tls_database_verify_chain_finish' : parameter 'result' of type 'AsyncResult' not supported

var tlsInteractionObject *gi.Object
var tlsInteractionObject_Once sync.Once

func tlsInteractionObject_Set() error {
	var err error
	tlsInteractionObject_Once.Do(func() {
		tlsInteractionObject, err = gi.ObjectNew("Gio", "TlsInteraction")
	})
	return err
}

type TlsInteraction struct {
	native unsafe.Pointer
}

func TlsInteractionNewFromNative(native unsafe.Pointer) *TlsInteraction {
	instance := &TlsInteraction{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *TlsInteraction) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToTlsInteraction down casts any arbitrary Object to TlsInteraction.
Exercise care, as this is a potentially dangerous function
if the Object is not a TlsInteraction.
*/
func CastToTlsInteraction(object *gobject.Object) *TlsInteraction {
	return TlsInteractionNewFromNative(object.Native())
}

// Equals compares this TlsInteraction with another TlsInteraction, and returns true if they represent the same GObject.
func (recv *TlsInteraction) Equals(other *TlsInteraction) bool {
	return other.Native() == recv.Native()
}

func (recv *TlsInteraction) Native() unsafe.Pointer {
	return recv.native
}

var tlsInteractionAskPasswordFunction *gi.Function
var tlsInteractionAskPasswordFunction_Once sync.Once

func tlsInteractionAskPasswordFunction_Set() error {
	var err error
	tlsInteractionAskPasswordFunction_Once.Do(func() {
		err = tlsInteractionObject_Set()
		if err != nil {
			return
		}
		tlsInteractionAskPasswordFunction, err = tlsInteractionObject.InvokerNew("ask_password")
	})
	return err
}

// AskPassword is a representation of the C type g_tls_interaction_ask_password.
func (recv *TlsInteraction) AskPassword(password *TlsPassword, cancellable *Cancellable) TlsInteractionResult {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(password.Native())
	inArgs[2].SetPointer(cancellable.Native())

	var ret gi.Argument

	err := tlsInteractionAskPasswordFunction_Set()
	if err == nil {
		ret = tlsInteractionAskPasswordFunction.Invoke(inArgs[:], nil)
	}

	retGo := TlsInteractionResult(ret.Int32())

	return retGo
}

// UNSUPPORTED : C value 'g_tls_interaction_ask_password_async' : parameter 'callback' of type 'AsyncReadyCallback' not supported

// UNSUPPORTED : C value 'g_tls_interaction_ask_password_finish' : parameter 'result' of type 'AsyncResult' not supported

var tlsInteractionInvokeAskPasswordFunction *gi.Function
var tlsInteractionInvokeAskPasswordFunction_Once sync.Once

func tlsInteractionInvokeAskPasswordFunction_Set() error {
	var err error
	tlsInteractionInvokeAskPasswordFunction_Once.Do(func() {
		err = tlsInteractionObject_Set()
		if err != nil {
			return
		}
		tlsInteractionInvokeAskPasswordFunction, err = tlsInteractionObject.InvokerNew("invoke_ask_password")
	})
	return err
}

// InvokeAskPassword is a representation of the C type g_tls_interaction_invoke_ask_password.
func (recv *TlsInteraction) InvokeAskPassword(password *TlsPassword, cancellable *Cancellable) TlsInteractionResult {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(password.Native())
	inArgs[2].SetPointer(cancellable.Native())

	var ret gi.Argument

	err := tlsInteractionInvokeAskPasswordFunction_Set()
	if err == nil {
		ret = tlsInteractionInvokeAskPasswordFunction.Invoke(inArgs[:], nil)
	}

	retGo := TlsInteractionResult(ret.Int32())

	return retGo
}

var tlsInteractionInvokeRequestCertificateFunction *gi.Function
var tlsInteractionInvokeRequestCertificateFunction_Once sync.Once

func tlsInteractionInvokeRequestCertificateFunction_Set() error {
	var err error
	tlsInteractionInvokeRequestCertificateFunction_Once.Do(func() {
		err = tlsInteractionObject_Set()
		if err != nil {
			return
		}
		tlsInteractionInvokeRequestCertificateFunction, err = tlsInteractionObject.InvokerNew("invoke_request_certificate")
	})
	return err
}

// InvokeRequestCertificate is a representation of the C type g_tls_interaction_invoke_request_certificate.
func (recv *TlsInteraction) InvokeRequestCertificate(connection *TlsConnection, flags TlsCertificateRequestFlags, cancellable *Cancellable) TlsInteractionResult {
	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(connection.Native())
	inArgs[2].SetInt32(int32(flags))
	inArgs[3].SetPointer(cancellable.Native())

	var ret gi.Argument

	err := tlsInteractionInvokeRequestCertificateFunction_Set()
	if err == nil {
		ret = tlsInteractionInvokeRequestCertificateFunction.Invoke(inArgs[:], nil)
	}

	retGo := TlsInteractionResult(ret.Int32())

	return retGo
}

var tlsInteractionRequestCertificateFunction *gi.Function
var tlsInteractionRequestCertificateFunction_Once sync.Once

func tlsInteractionRequestCertificateFunction_Set() error {
	var err error
	tlsInteractionRequestCertificateFunction_Once.Do(func() {
		err = tlsInteractionObject_Set()
		if err != nil {
			return
		}
		tlsInteractionRequestCertificateFunction, err = tlsInteractionObject.InvokerNew("request_certificate")
	})
	return err
}

// RequestCertificate is a representation of the C type g_tls_interaction_request_certificate.
func (recv *TlsInteraction) RequestCertificate(connection *TlsConnection, flags TlsCertificateRequestFlags, cancellable *Cancellable) TlsInteractionResult {
	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(connection.Native())
	inArgs[2].SetInt32(int32(flags))
	inArgs[3].SetPointer(cancellable.Native())

	var ret gi.Argument

	err := tlsInteractionRequestCertificateFunction_Set()
	if err == nil {
		ret = tlsInteractionRequestCertificateFunction.Invoke(inArgs[:], nil)
	}

	retGo := TlsInteractionResult(ret.Int32())

	return retGo
}

// UNSUPPORTED : C value 'g_tls_interaction_request_certificate_async' : parameter 'callback' of type 'AsyncReadyCallback' not supported

// UNSUPPORTED : C value 'g_tls_interaction_request_certificate_finish' : parameter 'result' of type 'AsyncResult' not supported

var tlsPasswordObject *gi.Object
var tlsPasswordObject_Once sync.Once

func tlsPasswordObject_Set() error {
	var err error
	tlsPasswordObject_Once.Do(func() {
		tlsPasswordObject, err = gi.ObjectNew("Gio", "TlsPassword")
	})
	return err
}

type TlsPassword struct {
	native unsafe.Pointer
}

func TlsPasswordNewFromNative(native unsafe.Pointer) *TlsPassword {
	instance := &TlsPassword{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *TlsPassword) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToTlsPassword down casts any arbitrary Object to TlsPassword.
Exercise care, as this is a potentially dangerous function
if the Object is not a TlsPassword.
*/
func CastToTlsPassword(object *gobject.Object) *TlsPassword {
	return TlsPasswordNewFromNative(object.Native())
}

// Equals compares this TlsPassword with another TlsPassword, and returns true if they represent the same GObject.
func (recv *TlsPassword) Equals(other *TlsPassword) bool {
	return other.Native() == recv.Native()
}

func (recv *TlsPassword) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *TlsPassword) FieldParentInstance() *gobject.Object {
	argValue := gi.ObjectFieldGet(tlsPasswordObject, recv.Native(), "parent_instance")
	value := gobject.ObjectNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *TlsPassword) SetFieldParentInstance(value *gobject.Object) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(tlsPasswordObject, recv.Native(), "parent_instance", argValue)
}

// FieldPriv returns the C field 'priv'.
func (recv *TlsPassword) FieldPriv() *TlsPasswordPrivate {
	argValue := gi.ObjectFieldGet(tlsPasswordObject, recv.Native(), "priv")
	value := TlsPasswordPrivateNewFromNative(argValue.Pointer())
	return value
}

// SetFieldPriv sets the value of the C field 'priv'.
func (recv *TlsPassword) SetFieldPriv(value *TlsPasswordPrivate) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(tlsPasswordObject, recv.Native(), "priv", argValue)
}

// UNSUPPORTED : C value 'g_tls_password_new' : parameter 'flags' of type 'TlsPasswordFlags' not supported

var tlsPasswordGetDescriptionFunction *gi.Function
var tlsPasswordGetDescriptionFunction_Once sync.Once

func tlsPasswordGetDescriptionFunction_Set() error {
	var err error
	tlsPasswordGetDescriptionFunction_Once.Do(func() {
		err = tlsPasswordObject_Set()
		if err != nil {
			return
		}
		tlsPasswordGetDescriptionFunction, err = tlsPasswordObject.InvokerNew("get_description")
	})
	return err
}

// GetDescription is a representation of the C type g_tls_password_get_description.
func (recv *TlsPassword) GetDescription() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := tlsPasswordGetDescriptionFunction_Set()
	if err == nil {
		ret = tlsPasswordGetDescriptionFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'g_tls_password_get_flags' : return type 'TlsPasswordFlags' not supported

var tlsPasswordGetValueFunction *gi.Function
var tlsPasswordGetValueFunction_Once sync.Once

func tlsPasswordGetValueFunction_Set() error {
	var err error
	tlsPasswordGetValueFunction_Once.Do(func() {
		err = tlsPasswordObject_Set()
		if err != nil {
			return
		}
		tlsPasswordGetValueFunction, err = tlsPasswordObject.InvokerNew("get_value")
	})
	return err
}

// GetValue is a representation of the C type g_tls_password_get_value.
func (recv *TlsPassword) GetValue(length uint64) uint8 {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetUint64(length)

	var ret gi.Argument

	err := tlsPasswordGetValueFunction_Set()
	if err == nil {
		ret = tlsPasswordGetValueFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint8()

	return retGo
}

var tlsPasswordGetWarningFunction *gi.Function
var tlsPasswordGetWarningFunction_Once sync.Once

func tlsPasswordGetWarningFunction_Set() error {
	var err error
	tlsPasswordGetWarningFunction_Once.Do(func() {
		err = tlsPasswordObject_Set()
		if err != nil {
			return
		}
		tlsPasswordGetWarningFunction, err = tlsPasswordObject.InvokerNew("get_warning")
	})
	return err
}

// GetWarning is a representation of the C type g_tls_password_get_warning.
func (recv *TlsPassword) GetWarning() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := tlsPasswordGetWarningFunction_Set()
	if err == nil {
		ret = tlsPasswordGetWarningFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var tlsPasswordSetDescriptionFunction *gi.Function
var tlsPasswordSetDescriptionFunction_Once sync.Once

func tlsPasswordSetDescriptionFunction_Set() error {
	var err error
	tlsPasswordSetDescriptionFunction_Once.Do(func() {
		err = tlsPasswordObject_Set()
		if err != nil {
			return
		}
		tlsPasswordSetDescriptionFunction, err = tlsPasswordObject.InvokerNew("set_description")
	})
	return err
}

// SetDescription is a representation of the C type g_tls_password_set_description.
func (recv *TlsPassword) SetDescription(description string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(description)

	err := tlsPasswordSetDescriptionFunction_Set()
	if err == nil {
		tlsPasswordSetDescriptionFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'g_tls_password_set_flags' : parameter 'flags' of type 'TlsPasswordFlags' not supported

// UNSUPPORTED : C value 'g_tls_password_set_value' : parameter 'value' of type 'nil' not supported

// UNSUPPORTED : C value 'g_tls_password_set_value_full' : parameter 'value' of type 'nil' not supported

var tlsPasswordSetWarningFunction *gi.Function
var tlsPasswordSetWarningFunction_Once sync.Once

func tlsPasswordSetWarningFunction_Set() error {
	var err error
	tlsPasswordSetWarningFunction_Once.Do(func() {
		err = tlsPasswordObject_Set()
		if err != nil {
			return
		}
		tlsPasswordSetWarningFunction, err = tlsPasswordObject.InvokerNew("set_warning")
	})
	return err
}

// SetWarning is a representation of the C type g_tls_password_set_warning.
func (recv *TlsPassword) SetWarning(warning string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(warning)

	err := tlsPasswordSetWarningFunction_Set()
	if err == nil {
		tlsPasswordSetWarningFunction.Invoke(inArgs[:], nil)
	}

	return
}

var unixConnectionObject *gi.Object
var unixConnectionObject_Once sync.Once

func unixConnectionObject_Set() error {
	var err error
	unixConnectionObject_Once.Do(func() {
		unixConnectionObject, err = gi.ObjectNew("Gio", "UnixConnection")
	})
	return err
}

type UnixConnection struct {
	native unsafe.Pointer
}

func UnixConnectionNewFromNative(native unsafe.Pointer) *UnixConnection {
	instance := &UnixConnection{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// SocketConnection upcasts to *SocketConnection
func (recv *UnixConnection) SocketConnection() *SocketConnection {
	return SocketConnectionNewFromNative(recv.Native())
}

// IOStream upcasts to *IOStream
func (recv *UnixConnection) IOStream() *IOStream {
	return IOStreamNewFromNative(recv.Native())
}

// Object upcasts to *Object
func (recv *UnixConnection) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToUnixConnection down casts any arbitrary Object to UnixConnection.
Exercise care, as this is a potentially dangerous function
if the Object is not a UnixConnection.
*/
func CastToUnixConnection(object *gobject.Object) *UnixConnection {
	return UnixConnectionNewFromNative(object.Native())
}

// Equals compares this UnixConnection with another UnixConnection, and returns true if they represent the same GObject.
func (recv *UnixConnection) Equals(other *UnixConnection) bool {
	return other.Native() == recv.Native()
}

func (recv *UnixConnection) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *UnixConnection) FieldParentInstance() *SocketConnection {
	argValue := gi.ObjectFieldGet(unixConnectionObject, recv.Native(), "parent_instance")
	value := SocketConnectionNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *UnixConnection) SetFieldParentInstance(value *SocketConnection) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(unixConnectionObject, recv.Native(), "parent_instance", argValue)
}

// FieldPriv returns the C field 'priv'.
func (recv *UnixConnection) FieldPriv() *UnixConnectionPrivate {
	argValue := gi.ObjectFieldGet(unixConnectionObject, recv.Native(), "priv")
	value := UnixConnectionPrivateNewFromNative(argValue.Pointer())
	return value
}

// SetFieldPriv sets the value of the C field 'priv'.
func (recv *UnixConnection) SetFieldPriv(value *UnixConnectionPrivate) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(unixConnectionObject, recv.Native(), "priv", argValue)
}

var unixConnectionReceiveCredentialsFunction *gi.Function
var unixConnectionReceiveCredentialsFunction_Once sync.Once

func unixConnectionReceiveCredentialsFunction_Set() error {
	var err error
	unixConnectionReceiveCredentialsFunction_Once.Do(func() {
		err = unixConnectionObject_Set()
		if err != nil {
			return
		}
		unixConnectionReceiveCredentialsFunction, err = unixConnectionObject.InvokerNew("receive_credentials")
	})
	return err
}

// ReceiveCredentials is a representation of the C type g_unix_connection_receive_credentials.
func (recv *UnixConnection) ReceiveCredentials(cancellable *Cancellable) *Credentials {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(cancellable.Native())

	var ret gi.Argument

	err := unixConnectionReceiveCredentialsFunction_Set()
	if err == nil {
		ret = unixConnectionReceiveCredentialsFunction.Invoke(inArgs[:], nil)
	}

	retGo := CredentialsNewFromNative(ret.Pointer())

	return retGo
}

// UNSUPPORTED : C value 'g_unix_connection_receive_credentials_async' : parameter 'callback' of type 'AsyncReadyCallback' not supported

// UNSUPPORTED : C value 'g_unix_connection_receive_credentials_finish' : parameter 'result' of type 'AsyncResult' not supported

var unixConnectionReceiveFdFunction *gi.Function
var unixConnectionReceiveFdFunction_Once sync.Once

func unixConnectionReceiveFdFunction_Set() error {
	var err error
	unixConnectionReceiveFdFunction_Once.Do(func() {
		err = unixConnectionObject_Set()
		if err != nil {
			return
		}
		unixConnectionReceiveFdFunction, err = unixConnectionObject.InvokerNew("receive_fd")
	})
	return err
}

// ReceiveFd is a representation of the C type g_unix_connection_receive_fd.
func (recv *UnixConnection) ReceiveFd(cancellable *Cancellable) int32 {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(cancellable.Native())

	var ret gi.Argument

	err := unixConnectionReceiveFdFunction_Set()
	if err == nil {
		ret = unixConnectionReceiveFdFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var unixConnectionSendCredentialsFunction *gi.Function
var unixConnectionSendCredentialsFunction_Once sync.Once

func unixConnectionSendCredentialsFunction_Set() error {
	var err error
	unixConnectionSendCredentialsFunction_Once.Do(func() {
		err = unixConnectionObject_Set()
		if err != nil {
			return
		}
		unixConnectionSendCredentialsFunction, err = unixConnectionObject.InvokerNew("send_credentials")
	})
	return err
}

// SendCredentials is a representation of the C type g_unix_connection_send_credentials.
func (recv *UnixConnection) SendCredentials(cancellable *Cancellable) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(cancellable.Native())

	var ret gi.Argument

	err := unixConnectionSendCredentialsFunction_Set()
	if err == nil {
		ret = unixConnectionSendCredentialsFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'g_unix_connection_send_credentials_async' : parameter 'callback' of type 'AsyncReadyCallback' not supported

// UNSUPPORTED : C value 'g_unix_connection_send_credentials_finish' : parameter 'result' of type 'AsyncResult' not supported

var unixConnectionSendFdFunction *gi.Function
var unixConnectionSendFdFunction_Once sync.Once

func unixConnectionSendFdFunction_Set() error {
	var err error
	unixConnectionSendFdFunction_Once.Do(func() {
		err = unixConnectionObject_Set()
		if err != nil {
			return
		}
		unixConnectionSendFdFunction, err = unixConnectionObject.InvokerNew("send_fd")
	})
	return err
}

// SendFd is a representation of the C type g_unix_connection_send_fd.
func (recv *UnixConnection) SendFd(fd int32, cancellable *Cancellable) bool {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(fd)
	inArgs[2].SetPointer(cancellable.Native())

	var ret gi.Argument

	err := unixConnectionSendFdFunction_Set()
	if err == nil {
		ret = unixConnectionSendFdFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var unixCredentialsMessageObject *gi.Object
var unixCredentialsMessageObject_Once sync.Once

func unixCredentialsMessageObject_Set() error {
	var err error
	unixCredentialsMessageObject_Once.Do(func() {
		unixCredentialsMessageObject, err = gi.ObjectNew("Gio", "UnixCredentialsMessage")
	})
	return err
}

type UnixCredentialsMessage struct {
	native unsafe.Pointer
}

func UnixCredentialsMessageNewFromNative(native unsafe.Pointer) *UnixCredentialsMessage {
	instance := &UnixCredentialsMessage{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// SocketControlMessage upcasts to *SocketControlMessage
func (recv *UnixCredentialsMessage) SocketControlMessage() *SocketControlMessage {
	return SocketControlMessageNewFromNative(recv.Native())
}

// Object upcasts to *Object
func (recv *UnixCredentialsMessage) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToUnixCredentialsMessage down casts any arbitrary Object to UnixCredentialsMessage.
Exercise care, as this is a potentially dangerous function
if the Object is not a UnixCredentialsMessage.
*/
func CastToUnixCredentialsMessage(object *gobject.Object) *UnixCredentialsMessage {
	return UnixCredentialsMessageNewFromNative(object.Native())
}

// Equals compares this UnixCredentialsMessage with another UnixCredentialsMessage, and returns true if they represent the same GObject.
func (recv *UnixCredentialsMessage) Equals(other *UnixCredentialsMessage) bool {
	return other.Native() == recv.Native()
}

func (recv *UnixCredentialsMessage) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *UnixCredentialsMessage) FieldParentInstance() *SocketControlMessage {
	argValue := gi.ObjectFieldGet(unixCredentialsMessageObject, recv.Native(), "parent_instance")
	value := SocketControlMessageNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *UnixCredentialsMessage) SetFieldParentInstance(value *SocketControlMessage) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(unixCredentialsMessageObject, recv.Native(), "parent_instance", argValue)
}

// FieldPriv returns the C field 'priv'.
func (recv *UnixCredentialsMessage) FieldPriv() *UnixCredentialsMessagePrivate {
	argValue := gi.ObjectFieldGet(unixCredentialsMessageObject, recv.Native(), "priv")
	value := UnixCredentialsMessagePrivateNewFromNative(argValue.Pointer())
	return value
}

// SetFieldPriv sets the value of the C field 'priv'.
func (recv *UnixCredentialsMessage) SetFieldPriv(value *UnixCredentialsMessagePrivate) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(unixCredentialsMessageObject, recv.Native(), "priv", argValue)
}

var unixCredentialsMessageNewFunction *gi.Function
var unixCredentialsMessageNewFunction_Once sync.Once

func unixCredentialsMessageNewFunction_Set() error {
	var err error
	unixCredentialsMessageNewFunction_Once.Do(func() {
		err = unixCredentialsMessageObject_Set()
		if err != nil {
			return
		}
		unixCredentialsMessageNewFunction, err = unixCredentialsMessageObject.InvokerNew("new")
	})
	return err
}

// UnixCredentialsMessageNew is a representation of the C type g_unix_credentials_message_new.
func UnixCredentialsMessageNew() *UnixCredentialsMessage {

	var ret gi.Argument

	err := unixCredentialsMessageNewFunction_Set()
	if err == nil {
		ret = unixCredentialsMessageNewFunction.Invoke(nil, nil)
	}

	retGo := UnixCredentialsMessageNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

var unixCredentialsMessageNewWithCredentialsFunction *gi.Function
var unixCredentialsMessageNewWithCredentialsFunction_Once sync.Once

func unixCredentialsMessageNewWithCredentialsFunction_Set() error {
	var err error
	unixCredentialsMessageNewWithCredentialsFunction_Once.Do(func() {
		err = unixCredentialsMessageObject_Set()
		if err != nil {
			return
		}
		unixCredentialsMessageNewWithCredentialsFunction, err = unixCredentialsMessageObject.InvokerNew("new_with_credentials")
	})
	return err
}

// UnixCredentialsMessageNewWithCredentials is a representation of the C type g_unix_credentials_message_new_with_credentials.
func UnixCredentialsMessageNewWithCredentials(credentials *Credentials) *UnixCredentialsMessage {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(credentials.Native())

	var ret gi.Argument

	err := unixCredentialsMessageNewWithCredentialsFunction_Set()
	if err == nil {
		ret = unixCredentialsMessageNewWithCredentialsFunction.Invoke(inArgs[:], nil)
	}

	retGo := UnixCredentialsMessageNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

var unixCredentialsMessageGetCredentialsFunction *gi.Function
var unixCredentialsMessageGetCredentialsFunction_Once sync.Once

func unixCredentialsMessageGetCredentialsFunction_Set() error {
	var err error
	unixCredentialsMessageGetCredentialsFunction_Once.Do(func() {
		err = unixCredentialsMessageObject_Set()
		if err != nil {
			return
		}
		unixCredentialsMessageGetCredentialsFunction, err = unixCredentialsMessageObject.InvokerNew("get_credentials")
	})
	return err
}

// GetCredentials is a representation of the C type g_unix_credentials_message_get_credentials.
func (recv *UnixCredentialsMessage) GetCredentials() *Credentials {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := unixCredentialsMessageGetCredentialsFunction_Set()
	if err == nil {
		ret = unixCredentialsMessageGetCredentialsFunction.Invoke(inArgs[:], nil)
	}

	retGo := CredentialsNewFromNative(ret.Pointer())

	return retGo
}

var unixFDListObject *gi.Object
var unixFDListObject_Once sync.Once

func unixFDListObject_Set() error {
	var err error
	unixFDListObject_Once.Do(func() {
		unixFDListObject, err = gi.ObjectNew("Gio", "UnixFDList")
	})
	return err
}

type UnixFDList struct {
	native unsafe.Pointer
}

func UnixFDListNewFromNative(native unsafe.Pointer) *UnixFDList {
	instance := &UnixFDList{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *UnixFDList) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToUnixFDList down casts any arbitrary Object to UnixFDList.
Exercise care, as this is a potentially dangerous function
if the Object is not a UnixFDList.
*/
func CastToUnixFDList(object *gobject.Object) *UnixFDList {
	return UnixFDListNewFromNative(object.Native())
}

// Equals compares this UnixFDList with another UnixFDList, and returns true if they represent the same GObject.
func (recv *UnixFDList) Equals(other *UnixFDList) bool {
	return other.Native() == recv.Native()
}

func (recv *UnixFDList) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *UnixFDList) FieldParentInstance() *gobject.Object {
	argValue := gi.ObjectFieldGet(unixFDListObject, recv.Native(), "parent_instance")
	value := gobject.ObjectNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *UnixFDList) SetFieldParentInstance(value *gobject.Object) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(unixFDListObject, recv.Native(), "parent_instance", argValue)
}

// FieldPriv returns the C field 'priv'.
func (recv *UnixFDList) FieldPriv() *UnixFDListPrivate {
	argValue := gi.ObjectFieldGet(unixFDListObject, recv.Native(), "priv")
	value := UnixFDListPrivateNewFromNative(argValue.Pointer())
	return value
}

// SetFieldPriv sets the value of the C field 'priv'.
func (recv *UnixFDList) SetFieldPriv(value *UnixFDListPrivate) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(unixFDListObject, recv.Native(), "priv", argValue)
}

var unixFDListNewFunction *gi.Function
var unixFDListNewFunction_Once sync.Once

func unixFDListNewFunction_Set() error {
	var err error
	unixFDListNewFunction_Once.Do(func() {
		err = unixFDListObject_Set()
		if err != nil {
			return
		}
		unixFDListNewFunction, err = unixFDListObject.InvokerNew("new")
	})
	return err
}

// UnixFDListNew is a representation of the C type g_unix_fd_list_new.
func UnixFDListNew() *UnixFDList {

	var ret gi.Argument

	err := unixFDListNewFunction_Set()
	if err == nil {
		ret = unixFDListNewFunction.Invoke(nil, nil)
	}

	retGo := UnixFDListNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

// UNSUPPORTED : C value 'g_unix_fd_list_new_from_array' : parameter 'fds' of type 'nil' not supported

var unixFDListAppendFunction *gi.Function
var unixFDListAppendFunction_Once sync.Once

func unixFDListAppendFunction_Set() error {
	var err error
	unixFDListAppendFunction_Once.Do(func() {
		err = unixFDListObject_Set()
		if err != nil {
			return
		}
		unixFDListAppendFunction, err = unixFDListObject.InvokerNew("append")
	})
	return err
}

// Append is a representation of the C type g_unix_fd_list_append.
func (recv *UnixFDList) Append(fd int32) int32 {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(fd)

	var ret gi.Argument

	err := unixFDListAppendFunction_Set()
	if err == nil {
		ret = unixFDListAppendFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var unixFDListGetFunction *gi.Function
var unixFDListGetFunction_Once sync.Once

func unixFDListGetFunction_Set() error {
	var err error
	unixFDListGetFunction_Once.Do(func() {
		err = unixFDListObject_Set()
		if err != nil {
			return
		}
		unixFDListGetFunction, err = unixFDListObject.InvokerNew("get")
	})
	return err
}

// Get is a representation of the C type g_unix_fd_list_get.
func (recv *UnixFDList) Get(index int32) int32 {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(index)

	var ret gi.Argument

	err := unixFDListGetFunction_Set()
	if err == nil {
		ret = unixFDListGetFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var unixFDListGetLengthFunction *gi.Function
var unixFDListGetLengthFunction_Once sync.Once

func unixFDListGetLengthFunction_Set() error {
	var err error
	unixFDListGetLengthFunction_Once.Do(func() {
		err = unixFDListObject_Set()
		if err != nil {
			return
		}
		unixFDListGetLengthFunction, err = unixFDListObject.InvokerNew("get_length")
	})
	return err
}

// GetLength is a representation of the C type g_unix_fd_list_get_length.
func (recv *UnixFDList) GetLength() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := unixFDListGetLengthFunction_Set()
	if err == nil {
		ret = unixFDListGetLengthFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var unixFDListPeekFdsFunction *gi.Function
var unixFDListPeekFdsFunction_Once sync.Once

func unixFDListPeekFdsFunction_Set() error {
	var err error
	unixFDListPeekFdsFunction_Once.Do(func() {
		err = unixFDListObject_Set()
		if err != nil {
			return
		}
		unixFDListPeekFdsFunction, err = unixFDListObject.InvokerNew("peek_fds")
	})
	return err
}

// PeekFds is a representation of the C type g_unix_fd_list_peek_fds.
func (recv *UnixFDList) PeekFds() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var outArgs [1]gi.Argument

	err := unixFDListPeekFdsFunction_Set()
	if err == nil {
		unixFDListPeekFdsFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := outArgs[0].Int32()

	return out0
}

var unixFDListStealFdsFunction *gi.Function
var unixFDListStealFdsFunction_Once sync.Once

func unixFDListStealFdsFunction_Set() error {
	var err error
	unixFDListStealFdsFunction_Once.Do(func() {
		err = unixFDListObject_Set()
		if err != nil {
			return
		}
		unixFDListStealFdsFunction, err = unixFDListObject.InvokerNew("steal_fds")
	})
	return err
}

// StealFds is a representation of the C type g_unix_fd_list_steal_fds.
func (recv *UnixFDList) StealFds() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var outArgs [1]gi.Argument

	err := unixFDListStealFdsFunction_Set()
	if err == nil {
		unixFDListStealFdsFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := outArgs[0].Int32()

	return out0
}

var unixFDMessageObject *gi.Object
var unixFDMessageObject_Once sync.Once

func unixFDMessageObject_Set() error {
	var err error
	unixFDMessageObject_Once.Do(func() {
		unixFDMessageObject, err = gi.ObjectNew("Gio", "UnixFDMessage")
	})
	return err
}

type UnixFDMessage struct {
	native unsafe.Pointer
}

func UnixFDMessageNewFromNative(native unsafe.Pointer) *UnixFDMessage {
	instance := &UnixFDMessage{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// SocketControlMessage upcasts to *SocketControlMessage
func (recv *UnixFDMessage) SocketControlMessage() *SocketControlMessage {
	return SocketControlMessageNewFromNative(recv.Native())
}

// Object upcasts to *Object
func (recv *UnixFDMessage) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToUnixFDMessage down casts any arbitrary Object to UnixFDMessage.
Exercise care, as this is a potentially dangerous function
if the Object is not a UnixFDMessage.
*/
func CastToUnixFDMessage(object *gobject.Object) *UnixFDMessage {
	return UnixFDMessageNewFromNative(object.Native())
}

// Equals compares this UnixFDMessage with another UnixFDMessage, and returns true if they represent the same GObject.
func (recv *UnixFDMessage) Equals(other *UnixFDMessage) bool {
	return other.Native() == recv.Native()
}

func (recv *UnixFDMessage) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *UnixFDMessage) FieldParentInstance() *SocketControlMessage {
	argValue := gi.ObjectFieldGet(unixFDMessageObject, recv.Native(), "parent_instance")
	value := SocketControlMessageNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *UnixFDMessage) SetFieldParentInstance(value *SocketControlMessage) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(unixFDMessageObject, recv.Native(), "parent_instance", argValue)
}

// FieldPriv returns the C field 'priv'.
func (recv *UnixFDMessage) FieldPriv() *UnixFDMessagePrivate {
	argValue := gi.ObjectFieldGet(unixFDMessageObject, recv.Native(), "priv")
	value := UnixFDMessagePrivateNewFromNative(argValue.Pointer())
	return value
}

// SetFieldPriv sets the value of the C field 'priv'.
func (recv *UnixFDMessage) SetFieldPriv(value *UnixFDMessagePrivate) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(unixFDMessageObject, recv.Native(), "priv", argValue)
}

var unixFDMessageNewFunction *gi.Function
var unixFDMessageNewFunction_Once sync.Once

func unixFDMessageNewFunction_Set() error {
	var err error
	unixFDMessageNewFunction_Once.Do(func() {
		err = unixFDMessageObject_Set()
		if err != nil {
			return
		}
		unixFDMessageNewFunction, err = unixFDMessageObject.InvokerNew("new")
	})
	return err
}

// UnixFDMessageNew is a representation of the C type g_unix_fd_message_new.
func UnixFDMessageNew() *UnixFDMessage {

	var ret gi.Argument

	err := unixFDMessageNewFunction_Set()
	if err == nil {
		ret = unixFDMessageNewFunction.Invoke(nil, nil)
	}

	retGo := UnixFDMessageNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

var unixFDMessageNewWithFdListFunction *gi.Function
var unixFDMessageNewWithFdListFunction_Once sync.Once

func unixFDMessageNewWithFdListFunction_Set() error {
	var err error
	unixFDMessageNewWithFdListFunction_Once.Do(func() {
		err = unixFDMessageObject_Set()
		if err != nil {
			return
		}
		unixFDMessageNewWithFdListFunction, err = unixFDMessageObject.InvokerNew("new_with_fd_list")
	})
	return err
}

// UnixFDMessageNewWithFdList is a representation of the C type g_unix_fd_message_new_with_fd_list.
func UnixFDMessageNewWithFdList(fdList *UnixFDList) *UnixFDMessage {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(fdList.Native())

	var ret gi.Argument

	err := unixFDMessageNewWithFdListFunction_Set()
	if err == nil {
		ret = unixFDMessageNewWithFdListFunction.Invoke(inArgs[:], nil)
	}

	retGo := UnixFDMessageNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

var unixFDMessageAppendFdFunction *gi.Function
var unixFDMessageAppendFdFunction_Once sync.Once

func unixFDMessageAppendFdFunction_Set() error {
	var err error
	unixFDMessageAppendFdFunction_Once.Do(func() {
		err = unixFDMessageObject_Set()
		if err != nil {
			return
		}
		unixFDMessageAppendFdFunction, err = unixFDMessageObject.InvokerNew("append_fd")
	})
	return err
}

// AppendFd is a representation of the C type g_unix_fd_message_append_fd.
func (recv *UnixFDMessage) AppendFd(fd int32) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(fd)

	var ret gi.Argument

	err := unixFDMessageAppendFdFunction_Set()
	if err == nil {
		ret = unixFDMessageAppendFdFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var unixFDMessageGetFdListFunction *gi.Function
var unixFDMessageGetFdListFunction_Once sync.Once

func unixFDMessageGetFdListFunction_Set() error {
	var err error
	unixFDMessageGetFdListFunction_Once.Do(func() {
		err = unixFDMessageObject_Set()
		if err != nil {
			return
		}
		unixFDMessageGetFdListFunction, err = unixFDMessageObject.InvokerNew("get_fd_list")
	})
	return err
}

// GetFdList is a representation of the C type g_unix_fd_message_get_fd_list.
func (recv *UnixFDMessage) GetFdList() *UnixFDList {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := unixFDMessageGetFdListFunction_Set()
	if err == nil {
		ret = unixFDMessageGetFdListFunction.Invoke(inArgs[:], nil)
	}

	retGo := UnixFDListNewFromNative(ret.Pointer())

	return retGo
}

var unixFDMessageStealFdsFunction *gi.Function
var unixFDMessageStealFdsFunction_Once sync.Once

func unixFDMessageStealFdsFunction_Set() error {
	var err error
	unixFDMessageStealFdsFunction_Once.Do(func() {
		err = unixFDMessageObject_Set()
		if err != nil {
			return
		}
		unixFDMessageStealFdsFunction, err = unixFDMessageObject.InvokerNew("steal_fds")
	})
	return err
}

// StealFds is a representation of the C type g_unix_fd_message_steal_fds.
func (recv *UnixFDMessage) StealFds() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var outArgs [1]gi.Argument

	err := unixFDMessageStealFdsFunction_Set()
	if err == nil {
		unixFDMessageStealFdsFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := outArgs[0].Int32()

	return out0
}

var unixInputStreamObject *gi.Object
var unixInputStreamObject_Once sync.Once

func unixInputStreamObject_Set() error {
	var err error
	unixInputStreamObject_Once.Do(func() {
		unixInputStreamObject, err = gi.ObjectNew("Gio", "UnixInputStream")
	})
	return err
}

type UnixInputStream struct {
	native unsafe.Pointer
}

func UnixInputStreamNewFromNative(native unsafe.Pointer) *UnixInputStream {
	instance := &UnixInputStream{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// InputStream upcasts to *InputStream
func (recv *UnixInputStream) InputStream() *InputStream {
	return InputStreamNewFromNative(recv.Native())
}

// Object upcasts to *Object
func (recv *UnixInputStream) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToUnixInputStream down casts any arbitrary Object to UnixInputStream.
Exercise care, as this is a potentially dangerous function
if the Object is not a UnixInputStream.
*/
func CastToUnixInputStream(object *gobject.Object) *UnixInputStream {
	return UnixInputStreamNewFromNative(object.Native())
}

// Equals compares this UnixInputStream with another UnixInputStream, and returns true if they represent the same GObject.
func (recv *UnixInputStream) Equals(other *UnixInputStream) bool {
	return other.Native() == recv.Native()
}

func (recv *UnixInputStream) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *UnixInputStream) FieldParentInstance() *InputStream {
	argValue := gi.ObjectFieldGet(unixInputStreamObject, recv.Native(), "parent_instance")
	value := InputStreamNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *UnixInputStream) SetFieldParentInstance(value *InputStream) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(unixInputStreamObject, recv.Native(), "parent_instance", argValue)
}

var unixInputStreamNewFunction *gi.Function
var unixInputStreamNewFunction_Once sync.Once

func unixInputStreamNewFunction_Set() error {
	var err error
	unixInputStreamNewFunction_Once.Do(func() {
		err = unixInputStreamObject_Set()
		if err != nil {
			return
		}
		unixInputStreamNewFunction, err = unixInputStreamObject.InvokerNew("new")
	})
	return err
}

// UnixInputStreamNew is a representation of the C type g_unix_input_stream_new.
func UnixInputStreamNew(fd int32, closeFd bool) *UnixInputStream {
	var inArgs [2]gi.Argument
	inArgs[0].SetInt32(fd)
	inArgs[1].SetBoolean(closeFd)

	var ret gi.Argument

	err := unixInputStreamNewFunction_Set()
	if err == nil {
		ret = unixInputStreamNewFunction.Invoke(inArgs[:], nil)
	}

	retGo := UnixInputStreamNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

var unixInputStreamGetCloseFdFunction *gi.Function
var unixInputStreamGetCloseFdFunction_Once sync.Once

func unixInputStreamGetCloseFdFunction_Set() error {
	var err error
	unixInputStreamGetCloseFdFunction_Once.Do(func() {
		err = unixInputStreamObject_Set()
		if err != nil {
			return
		}
		unixInputStreamGetCloseFdFunction, err = unixInputStreamObject.InvokerNew("get_close_fd")
	})
	return err
}

// GetCloseFd is a representation of the C type g_unix_input_stream_get_close_fd.
func (recv *UnixInputStream) GetCloseFd() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := unixInputStreamGetCloseFdFunction_Set()
	if err == nil {
		ret = unixInputStreamGetCloseFdFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var unixInputStreamGetFdFunction *gi.Function
var unixInputStreamGetFdFunction_Once sync.Once

func unixInputStreamGetFdFunction_Set() error {
	var err error
	unixInputStreamGetFdFunction_Once.Do(func() {
		err = unixInputStreamObject_Set()
		if err != nil {
			return
		}
		unixInputStreamGetFdFunction, err = unixInputStreamObject.InvokerNew("get_fd")
	})
	return err
}

// GetFd is a representation of the C type g_unix_input_stream_get_fd.
func (recv *UnixInputStream) GetFd() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := unixInputStreamGetFdFunction_Set()
	if err == nil {
		ret = unixInputStreamGetFdFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var unixInputStreamSetCloseFdFunction *gi.Function
var unixInputStreamSetCloseFdFunction_Once sync.Once

func unixInputStreamSetCloseFdFunction_Set() error {
	var err error
	unixInputStreamSetCloseFdFunction_Once.Do(func() {
		err = unixInputStreamObject_Set()
		if err != nil {
			return
		}
		unixInputStreamSetCloseFdFunction, err = unixInputStreamObject.InvokerNew("set_close_fd")
	})
	return err
}

// SetCloseFd is a representation of the C type g_unix_input_stream_set_close_fd.
func (recv *UnixInputStream) SetCloseFd(closeFd bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetBoolean(closeFd)

	err := unixInputStreamSetCloseFdFunction_Set()
	if err == nil {
		unixInputStreamSetCloseFdFunction.Invoke(inArgs[:], nil)
	}

	return
}

var unixMountMonitorObject *gi.Object
var unixMountMonitorObject_Once sync.Once

func unixMountMonitorObject_Set() error {
	var err error
	unixMountMonitorObject_Once.Do(func() {
		unixMountMonitorObject, err = gi.ObjectNew("Gio", "UnixMountMonitor")
	})
	return err
}

type UnixMountMonitor struct {
	native unsafe.Pointer
}

func UnixMountMonitorNewFromNative(native unsafe.Pointer) *UnixMountMonitor {
	instance := &UnixMountMonitor{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *UnixMountMonitor) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToUnixMountMonitor down casts any arbitrary Object to UnixMountMonitor.
Exercise care, as this is a potentially dangerous function
if the Object is not a UnixMountMonitor.
*/
func CastToUnixMountMonitor(object *gobject.Object) *UnixMountMonitor {
	return UnixMountMonitorNewFromNative(object.Native())
}

// Equals compares this UnixMountMonitor with another UnixMountMonitor, and returns true if they represent the same GObject.
func (recv *UnixMountMonitor) Equals(other *UnixMountMonitor) bool {
	return other.Native() == recv.Native()
}

func (recv *UnixMountMonitor) Native() unsafe.Pointer {
	return recv.native
}

var unixMountMonitorNewFunction *gi.Function
var unixMountMonitorNewFunction_Once sync.Once

func unixMountMonitorNewFunction_Set() error {
	var err error
	unixMountMonitorNewFunction_Once.Do(func() {
		err = unixMountMonitorObject_Set()
		if err != nil {
			return
		}
		unixMountMonitorNewFunction, err = unixMountMonitorObject.InvokerNew("new")
	})
	return err
}

// UnixMountMonitorNew is a representation of the C type g_unix_mount_monitor_new.
func UnixMountMonitorNew() *UnixMountMonitor {

	var ret gi.Argument

	err := unixMountMonitorNewFunction_Set()
	if err == nil {
		ret = unixMountMonitorNewFunction.Invoke(nil, nil)
	}

	retGo := UnixMountMonitorNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

var unixMountMonitorSetRateLimitFunction *gi.Function
var unixMountMonitorSetRateLimitFunction_Once sync.Once

func unixMountMonitorSetRateLimitFunction_Set() error {
	var err error
	unixMountMonitorSetRateLimitFunction_Once.Do(func() {
		err = unixMountMonitorObject_Set()
		if err != nil {
			return
		}
		unixMountMonitorSetRateLimitFunction, err = unixMountMonitorObject.InvokerNew("set_rate_limit")
	})
	return err
}

// SetRateLimit is a representation of the C type g_unix_mount_monitor_set_rate_limit.
func (recv *UnixMountMonitor) SetRateLimit(limitMsec int32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(limitMsec)

	err := unixMountMonitorSetRateLimitFunction_Set()
	if err == nil {
		unixMountMonitorSetRateLimitFunction.Invoke(inArgs[:], nil)
	}

	return
}

var unixOutputStreamObject *gi.Object
var unixOutputStreamObject_Once sync.Once

func unixOutputStreamObject_Set() error {
	var err error
	unixOutputStreamObject_Once.Do(func() {
		unixOutputStreamObject, err = gi.ObjectNew("Gio", "UnixOutputStream")
	})
	return err
}

type UnixOutputStream struct {
	native unsafe.Pointer
}

func UnixOutputStreamNewFromNative(native unsafe.Pointer) *UnixOutputStream {
	instance := &UnixOutputStream{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// OutputStream upcasts to *OutputStream
func (recv *UnixOutputStream) OutputStream() *OutputStream {
	return OutputStreamNewFromNative(recv.Native())
}

// Object upcasts to *Object
func (recv *UnixOutputStream) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToUnixOutputStream down casts any arbitrary Object to UnixOutputStream.
Exercise care, as this is a potentially dangerous function
if the Object is not a UnixOutputStream.
*/
func CastToUnixOutputStream(object *gobject.Object) *UnixOutputStream {
	return UnixOutputStreamNewFromNative(object.Native())
}

// Equals compares this UnixOutputStream with another UnixOutputStream, and returns true if they represent the same GObject.
func (recv *UnixOutputStream) Equals(other *UnixOutputStream) bool {
	return other.Native() == recv.Native()
}

func (recv *UnixOutputStream) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *UnixOutputStream) FieldParentInstance() *OutputStream {
	argValue := gi.ObjectFieldGet(unixOutputStreamObject, recv.Native(), "parent_instance")
	value := OutputStreamNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *UnixOutputStream) SetFieldParentInstance(value *OutputStream) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(unixOutputStreamObject, recv.Native(), "parent_instance", argValue)
}

var unixOutputStreamNewFunction *gi.Function
var unixOutputStreamNewFunction_Once sync.Once

func unixOutputStreamNewFunction_Set() error {
	var err error
	unixOutputStreamNewFunction_Once.Do(func() {
		err = unixOutputStreamObject_Set()
		if err != nil {
			return
		}
		unixOutputStreamNewFunction, err = unixOutputStreamObject.InvokerNew("new")
	})
	return err
}

// UnixOutputStreamNew is a representation of the C type g_unix_output_stream_new.
func UnixOutputStreamNew(fd int32, closeFd bool) *UnixOutputStream {
	var inArgs [2]gi.Argument
	inArgs[0].SetInt32(fd)
	inArgs[1].SetBoolean(closeFd)

	var ret gi.Argument

	err := unixOutputStreamNewFunction_Set()
	if err == nil {
		ret = unixOutputStreamNewFunction.Invoke(inArgs[:], nil)
	}

	retGo := UnixOutputStreamNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

var unixOutputStreamGetCloseFdFunction *gi.Function
var unixOutputStreamGetCloseFdFunction_Once sync.Once

func unixOutputStreamGetCloseFdFunction_Set() error {
	var err error
	unixOutputStreamGetCloseFdFunction_Once.Do(func() {
		err = unixOutputStreamObject_Set()
		if err != nil {
			return
		}
		unixOutputStreamGetCloseFdFunction, err = unixOutputStreamObject.InvokerNew("get_close_fd")
	})
	return err
}

// GetCloseFd is a representation of the C type g_unix_output_stream_get_close_fd.
func (recv *UnixOutputStream) GetCloseFd() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := unixOutputStreamGetCloseFdFunction_Set()
	if err == nil {
		ret = unixOutputStreamGetCloseFdFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var unixOutputStreamGetFdFunction *gi.Function
var unixOutputStreamGetFdFunction_Once sync.Once

func unixOutputStreamGetFdFunction_Set() error {
	var err error
	unixOutputStreamGetFdFunction_Once.Do(func() {
		err = unixOutputStreamObject_Set()
		if err != nil {
			return
		}
		unixOutputStreamGetFdFunction, err = unixOutputStreamObject.InvokerNew("get_fd")
	})
	return err
}

// GetFd is a representation of the C type g_unix_output_stream_get_fd.
func (recv *UnixOutputStream) GetFd() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := unixOutputStreamGetFdFunction_Set()
	if err == nil {
		ret = unixOutputStreamGetFdFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var unixOutputStreamSetCloseFdFunction *gi.Function
var unixOutputStreamSetCloseFdFunction_Once sync.Once

func unixOutputStreamSetCloseFdFunction_Set() error {
	var err error
	unixOutputStreamSetCloseFdFunction_Once.Do(func() {
		err = unixOutputStreamObject_Set()
		if err != nil {
			return
		}
		unixOutputStreamSetCloseFdFunction, err = unixOutputStreamObject.InvokerNew("set_close_fd")
	})
	return err
}

// SetCloseFd is a representation of the C type g_unix_output_stream_set_close_fd.
func (recv *UnixOutputStream) SetCloseFd(closeFd bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetBoolean(closeFd)

	err := unixOutputStreamSetCloseFdFunction_Set()
	if err == nil {
		unixOutputStreamSetCloseFdFunction.Invoke(inArgs[:], nil)
	}

	return
}

var unixSocketAddressObject *gi.Object
var unixSocketAddressObject_Once sync.Once

func unixSocketAddressObject_Set() error {
	var err error
	unixSocketAddressObject_Once.Do(func() {
		unixSocketAddressObject, err = gi.ObjectNew("Gio", "UnixSocketAddress")
	})
	return err
}

type UnixSocketAddress struct {
	native unsafe.Pointer
}

func UnixSocketAddressNewFromNative(native unsafe.Pointer) *UnixSocketAddress {
	instance := &UnixSocketAddress{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// SocketAddress upcasts to *SocketAddress
func (recv *UnixSocketAddress) SocketAddress() *SocketAddress {
	return SocketAddressNewFromNative(recv.Native())
}

// Object upcasts to *Object
func (recv *UnixSocketAddress) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToUnixSocketAddress down casts any arbitrary Object to UnixSocketAddress.
Exercise care, as this is a potentially dangerous function
if the Object is not a UnixSocketAddress.
*/
func CastToUnixSocketAddress(object *gobject.Object) *UnixSocketAddress {
	return UnixSocketAddressNewFromNative(object.Native())
}

// Equals compares this UnixSocketAddress with another UnixSocketAddress, and returns true if they represent the same GObject.
func (recv *UnixSocketAddress) Equals(other *UnixSocketAddress) bool {
	return other.Native() == recv.Native()
}

func (recv *UnixSocketAddress) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *UnixSocketAddress) FieldParentInstance() *SocketAddress {
	argValue := gi.ObjectFieldGet(unixSocketAddressObject, recv.Native(), "parent_instance")
	value := SocketAddressNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *UnixSocketAddress) SetFieldParentInstance(value *SocketAddress) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(unixSocketAddressObject, recv.Native(), "parent_instance", argValue)
}

var unixSocketAddressNewFunction *gi.Function
var unixSocketAddressNewFunction_Once sync.Once

func unixSocketAddressNewFunction_Set() error {
	var err error
	unixSocketAddressNewFunction_Once.Do(func() {
		err = unixSocketAddressObject_Set()
		if err != nil {
			return
		}
		unixSocketAddressNewFunction, err = unixSocketAddressObject.InvokerNew("new")
	})
	return err
}

// UnixSocketAddressNew is a representation of the C type g_unix_socket_address_new.
func UnixSocketAddressNew(path string) *UnixSocketAddress {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(path)

	var ret gi.Argument

	err := unixSocketAddressNewFunction_Set()
	if err == nil {
		ret = unixSocketAddressNewFunction.Invoke(inArgs[:], nil)
	}

	retGo := UnixSocketAddressNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

// UNSUPPORTED : C value 'g_unix_socket_address_new_abstract' : parameter 'path' of type 'nil' not supported

// UNSUPPORTED : C value 'g_unix_socket_address_new_with_type' : parameter 'path' of type 'nil' not supported

var unixSocketAddressGetAddressTypeFunction *gi.Function
var unixSocketAddressGetAddressTypeFunction_Once sync.Once

func unixSocketAddressGetAddressTypeFunction_Set() error {
	var err error
	unixSocketAddressGetAddressTypeFunction_Once.Do(func() {
		err = unixSocketAddressObject_Set()
		if err != nil {
			return
		}
		unixSocketAddressGetAddressTypeFunction, err = unixSocketAddressObject.InvokerNew("get_address_type")
	})
	return err
}

// GetAddressType is a representation of the C type g_unix_socket_address_get_address_type.
func (recv *UnixSocketAddress) GetAddressType() UnixSocketAddressType {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := unixSocketAddressGetAddressTypeFunction_Set()
	if err == nil {
		ret = unixSocketAddressGetAddressTypeFunction.Invoke(inArgs[:], nil)
	}

	retGo := UnixSocketAddressType(ret.Int32())

	return retGo
}

var unixSocketAddressGetIsAbstractFunction *gi.Function
var unixSocketAddressGetIsAbstractFunction_Once sync.Once

func unixSocketAddressGetIsAbstractFunction_Set() error {
	var err error
	unixSocketAddressGetIsAbstractFunction_Once.Do(func() {
		err = unixSocketAddressObject_Set()
		if err != nil {
			return
		}
		unixSocketAddressGetIsAbstractFunction, err = unixSocketAddressObject.InvokerNew("get_is_abstract")
	})
	return err
}

// GetIsAbstract is a representation of the C type g_unix_socket_address_get_is_abstract.
func (recv *UnixSocketAddress) GetIsAbstract() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := unixSocketAddressGetIsAbstractFunction_Set()
	if err == nil {
		ret = unixSocketAddressGetIsAbstractFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var unixSocketAddressGetPathFunction *gi.Function
var unixSocketAddressGetPathFunction_Once sync.Once

func unixSocketAddressGetPathFunction_Set() error {
	var err error
	unixSocketAddressGetPathFunction_Once.Do(func() {
		err = unixSocketAddressObject_Set()
		if err != nil {
			return
		}
		unixSocketAddressGetPathFunction, err = unixSocketAddressObject.InvokerNew("get_path")
	})
	return err
}

// GetPath is a representation of the C type g_unix_socket_address_get_path.
func (recv *UnixSocketAddress) GetPath() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := unixSocketAddressGetPathFunction_Set()
	if err == nil {
		ret = unixSocketAddressGetPathFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var unixSocketAddressGetPathLenFunction *gi.Function
var unixSocketAddressGetPathLenFunction_Once sync.Once

func unixSocketAddressGetPathLenFunction_Set() error {
	var err error
	unixSocketAddressGetPathLenFunction_Once.Do(func() {
		err = unixSocketAddressObject_Set()
		if err != nil {
			return
		}
		unixSocketAddressGetPathLenFunction, err = unixSocketAddressObject.InvokerNew("get_path_len")
	})
	return err
}

// GetPathLen is a representation of the C type g_unix_socket_address_get_path_len.
func (recv *UnixSocketAddress) GetPathLen() uint64 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := unixSocketAddressGetPathLenFunction_Set()
	if err == nil {
		ret = unixSocketAddressGetPathLenFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint64()

	return retGo
}

var vfsObject *gi.Object
var vfsObject_Once sync.Once

func vfsObject_Set() error {
	var err error
	vfsObject_Once.Do(func() {
		vfsObject, err = gi.ObjectNew("Gio", "Vfs")
	})
	return err
}

type Vfs struct {
	native unsafe.Pointer
}

func VfsNewFromNative(native unsafe.Pointer) *Vfs {
	instance := &Vfs{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *Vfs) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToVfs down casts any arbitrary Object to Vfs.
Exercise care, as this is a potentially dangerous function
if the Object is not a Vfs.
*/
func CastToVfs(object *gobject.Object) *Vfs {
	return VfsNewFromNative(object.Native())
}

// Equals compares this Vfs with another Vfs, and returns true if they represent the same GObject.
func (recv *Vfs) Equals(other *Vfs) bool {
	return other.Native() == recv.Native()
}

func (recv *Vfs) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *Vfs) FieldParentInstance() *gobject.Object {
	argValue := gi.ObjectFieldGet(vfsObject, recv.Native(), "parent_instance")
	value := gobject.ObjectNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *Vfs) SetFieldParentInstance(value *gobject.Object) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(vfsObject, recv.Native(), "parent_instance", argValue)
}

// UNSUPPORTED : C value 'g_vfs_get_file_for_path' : return type 'File' not supported

// UNSUPPORTED : C value 'g_vfs_get_file_for_uri' : return type 'File' not supported

var vfsGetSupportedUriSchemesFunction *gi.Function
var vfsGetSupportedUriSchemesFunction_Once sync.Once

func vfsGetSupportedUriSchemesFunction_Set() error {
	var err error
	vfsGetSupportedUriSchemesFunction_Once.Do(func() {
		err = vfsObject_Set()
		if err != nil {
			return
		}
		vfsGetSupportedUriSchemesFunction, err = vfsObject.InvokerNew("get_supported_uri_schemes")
	})
	return err
}

// GetSupportedUriSchemes is a representation of the C type g_vfs_get_supported_uri_schemes.
func (recv *Vfs) GetSupportedUriSchemes() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := vfsGetSupportedUriSchemesFunction_Set()
	if err == nil {
		vfsGetSupportedUriSchemesFunction.Invoke(inArgs[:], nil)
	}

	return
}

var vfsIsActiveFunction *gi.Function
var vfsIsActiveFunction_Once sync.Once

func vfsIsActiveFunction_Set() error {
	var err error
	vfsIsActiveFunction_Once.Do(func() {
		err = vfsObject_Set()
		if err != nil {
			return
		}
		vfsIsActiveFunction, err = vfsObject.InvokerNew("is_active")
	})
	return err
}

// IsActive is a representation of the C type g_vfs_is_active.
func (recv *Vfs) IsActive() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := vfsIsActiveFunction_Set()
	if err == nil {
		ret = vfsIsActiveFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'g_vfs_parse_name' : return type 'File' not supported

// UNSUPPORTED : C value 'g_vfs_register_uri_scheme' : parameter 'uri_func' of type 'VfsFileLookupFunc' not supported

var vfsUnregisterUriSchemeFunction *gi.Function
var vfsUnregisterUriSchemeFunction_Once sync.Once

func vfsUnregisterUriSchemeFunction_Set() error {
	var err error
	vfsUnregisterUriSchemeFunction_Once.Do(func() {
		err = vfsObject_Set()
		if err != nil {
			return
		}
		vfsUnregisterUriSchemeFunction, err = vfsObject.InvokerNew("unregister_uri_scheme")
	})
	return err
}

// UnregisterUriScheme is a representation of the C type g_vfs_unregister_uri_scheme.
func (recv *Vfs) UnregisterUriScheme(scheme string) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(scheme)

	var ret gi.Argument

	err := vfsUnregisterUriSchemeFunction_Set()
	if err == nil {
		ret = vfsUnregisterUriSchemeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var volumeMonitorObject *gi.Object
var volumeMonitorObject_Once sync.Once

func volumeMonitorObject_Set() error {
	var err error
	volumeMonitorObject_Once.Do(func() {
		volumeMonitorObject, err = gi.ObjectNew("Gio", "VolumeMonitor")
	})
	return err
}

type VolumeMonitor struct {
	native unsafe.Pointer
}

func VolumeMonitorNewFromNative(native unsafe.Pointer) *VolumeMonitor {
	instance := &VolumeMonitor{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *VolumeMonitor) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToVolumeMonitor down casts any arbitrary Object to VolumeMonitor.
Exercise care, as this is a potentially dangerous function
if the Object is not a VolumeMonitor.
*/
func CastToVolumeMonitor(object *gobject.Object) *VolumeMonitor {
	return VolumeMonitorNewFromNative(object.Native())
}

// Equals compares this VolumeMonitor with another VolumeMonitor, and returns true if they represent the same GObject.
func (recv *VolumeMonitor) Equals(other *VolumeMonitor) bool {
	return other.Native() == recv.Native()
}

func (recv *VolumeMonitor) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *VolumeMonitor) FieldParentInstance() *gobject.Object {
	argValue := gi.ObjectFieldGet(volumeMonitorObject, recv.Native(), "parent_instance")
	value := gobject.ObjectNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *VolumeMonitor) SetFieldParentInstance(value *gobject.Object) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(volumeMonitorObject, recv.Native(), "parent_instance", argValue)
}

// UNSUPPORTED : C value 'g_volume_monitor_get_connected_drives' : return type 'GLib.List' not supported

// UNSUPPORTED : C value 'g_volume_monitor_get_mount_for_uuid' : return type 'Mount' not supported

// UNSUPPORTED : C value 'g_volume_monitor_get_mounts' : return type 'GLib.List' not supported

// UNSUPPORTED : C value 'g_volume_monitor_get_volume_for_uuid' : return type 'Volume' not supported

// UNSUPPORTED : C value 'g_volume_monitor_get_volumes' : return type 'GLib.List' not supported

var zlibCompressorObject *gi.Object
var zlibCompressorObject_Once sync.Once

func zlibCompressorObject_Set() error {
	var err error
	zlibCompressorObject_Once.Do(func() {
		zlibCompressorObject, err = gi.ObjectNew("Gio", "ZlibCompressor")
	})
	return err
}

type ZlibCompressor struct {
	native unsafe.Pointer
}

func ZlibCompressorNewFromNative(native unsafe.Pointer) *ZlibCompressor {
	instance := &ZlibCompressor{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *ZlibCompressor) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToZlibCompressor down casts any arbitrary Object to ZlibCompressor.
Exercise care, as this is a potentially dangerous function
if the Object is not a ZlibCompressor.
*/
func CastToZlibCompressor(object *gobject.Object) *ZlibCompressor {
	return ZlibCompressorNewFromNative(object.Native())
}

// Equals compares this ZlibCompressor with another ZlibCompressor, and returns true if they represent the same GObject.
func (recv *ZlibCompressor) Equals(other *ZlibCompressor) bool {
	return other.Native() == recv.Native()
}

func (recv *ZlibCompressor) Native() unsafe.Pointer {
	return recv.native
}

var zlibCompressorNewFunction *gi.Function
var zlibCompressorNewFunction_Once sync.Once

func zlibCompressorNewFunction_Set() error {
	var err error
	zlibCompressorNewFunction_Once.Do(func() {
		err = zlibCompressorObject_Set()
		if err != nil {
			return
		}
		zlibCompressorNewFunction, err = zlibCompressorObject.InvokerNew("new")
	})
	return err
}

// ZlibCompressorNew is a representation of the C type g_zlib_compressor_new.
func ZlibCompressorNew(format ZlibCompressorFormat, level int32) *ZlibCompressor {
	var inArgs [2]gi.Argument
	inArgs[0].SetInt32(int32(format))
	inArgs[1].SetInt32(level)

	var ret gi.Argument

	err := zlibCompressorNewFunction_Set()
	if err == nil {
		ret = zlibCompressorNewFunction.Invoke(inArgs[:], nil)
	}

	retGo := ZlibCompressorNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

var zlibCompressorGetFileInfoFunction *gi.Function
var zlibCompressorGetFileInfoFunction_Once sync.Once

func zlibCompressorGetFileInfoFunction_Set() error {
	var err error
	zlibCompressorGetFileInfoFunction_Once.Do(func() {
		err = zlibCompressorObject_Set()
		if err != nil {
			return
		}
		zlibCompressorGetFileInfoFunction, err = zlibCompressorObject.InvokerNew("get_file_info")
	})
	return err
}

// GetFileInfo is a representation of the C type g_zlib_compressor_get_file_info.
func (recv *ZlibCompressor) GetFileInfo() *FileInfo {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := zlibCompressorGetFileInfoFunction_Set()
	if err == nil {
		ret = zlibCompressorGetFileInfoFunction.Invoke(inArgs[:], nil)
	}

	retGo := FileInfoNewFromNative(ret.Pointer())

	return retGo
}

var zlibCompressorSetFileInfoFunction *gi.Function
var zlibCompressorSetFileInfoFunction_Once sync.Once

func zlibCompressorSetFileInfoFunction_Set() error {
	var err error
	zlibCompressorSetFileInfoFunction_Once.Do(func() {
		err = zlibCompressorObject_Set()
		if err != nil {
			return
		}
		zlibCompressorSetFileInfoFunction, err = zlibCompressorObject.InvokerNew("set_file_info")
	})
	return err
}

// SetFileInfo is a representation of the C type g_zlib_compressor_set_file_info.
func (recv *ZlibCompressor) SetFileInfo(fileInfo *FileInfo) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(fileInfo.Native())

	err := zlibCompressorSetFileInfoFunction_Set()
	if err == nil {
		zlibCompressorSetFileInfoFunction.Invoke(inArgs[:], nil)
	}

	return
}

var zlibDecompressorObject *gi.Object
var zlibDecompressorObject_Once sync.Once

func zlibDecompressorObject_Set() error {
	var err error
	zlibDecompressorObject_Once.Do(func() {
		zlibDecompressorObject, err = gi.ObjectNew("Gio", "ZlibDecompressor")
	})
	return err
}

type ZlibDecompressor struct {
	native unsafe.Pointer
}

func ZlibDecompressorNewFromNative(native unsafe.Pointer) *ZlibDecompressor {
	instance := &ZlibDecompressor{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *ZlibDecompressor) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToZlibDecompressor down casts any arbitrary Object to ZlibDecompressor.
Exercise care, as this is a potentially dangerous function
if the Object is not a ZlibDecompressor.
*/
func CastToZlibDecompressor(object *gobject.Object) *ZlibDecompressor {
	return ZlibDecompressorNewFromNative(object.Native())
}

// Equals compares this ZlibDecompressor with another ZlibDecompressor, and returns true if they represent the same GObject.
func (recv *ZlibDecompressor) Equals(other *ZlibDecompressor) bool {
	return other.Native() == recv.Native()
}

func (recv *ZlibDecompressor) Native() unsafe.Pointer {
	return recv.native
}

var zlibDecompressorNewFunction *gi.Function
var zlibDecompressorNewFunction_Once sync.Once

func zlibDecompressorNewFunction_Set() error {
	var err error
	zlibDecompressorNewFunction_Once.Do(func() {
		err = zlibDecompressorObject_Set()
		if err != nil {
			return
		}
		zlibDecompressorNewFunction, err = zlibDecompressorObject.InvokerNew("new")
	})
	return err
}

// ZlibDecompressorNew is a representation of the C type g_zlib_decompressor_new.
func ZlibDecompressorNew(format ZlibCompressorFormat) *ZlibDecompressor {
	var inArgs [1]gi.Argument
	inArgs[0].SetInt32(int32(format))

	var ret gi.Argument

	err := zlibDecompressorNewFunction_Set()
	if err == nil {
		ret = zlibDecompressorNewFunction.Invoke(inArgs[:], nil)
	}

	retGo := ZlibDecompressorNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

var zlibDecompressorGetFileInfoFunction *gi.Function
var zlibDecompressorGetFileInfoFunction_Once sync.Once

func zlibDecompressorGetFileInfoFunction_Set() error {
	var err error
	zlibDecompressorGetFileInfoFunction_Once.Do(func() {
		err = zlibDecompressorObject_Set()
		if err != nil {
			return
		}
		zlibDecompressorGetFileInfoFunction, err = zlibDecompressorObject.InvokerNew("get_file_info")
	})
	return err
}

// GetFileInfo is a representation of the C type g_zlib_decompressor_get_file_info.
func (recv *ZlibDecompressor) GetFileInfo() *FileInfo {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := zlibDecompressorGetFileInfoFunction_Set()
	if err == nil {
		ret = zlibDecompressorGetFileInfoFunction.Invoke(inArgs[:], nil)
	}

	retGo := FileInfoNewFromNative(ret.Pointer())

	return retGo
}
