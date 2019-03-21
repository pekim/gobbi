// This is a generated file - DO NOT EDIT
// +build gio_2.26 gio_2.28 gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import (
	glib "github.com/pekim/gobbi/lib/glib"
	"sync"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gio/gdesktopappinfo.h>
// #include <gio/gfiledescriptorbased.h>
// #include <gio/gio.h>
// #include <gio/gunixconnection.h>
// #include <gio/gunixcredentialsmessage.h>
// #include <gio/gunixfdlist.h>
// #include <gio/gunixfdmessage.h>
// #include <gio/gunixinputstream.h>
// #include <gio/gunixmounts.h>
// #include <gio/gunixoutputstream.h>
// #include <gio/gunixsocketaddress.h>
// #include <stdlib.h>
/*

	void callback_busacquiredcallbackHandler(GObject *, GDBusConnection*, gchar*, gpointer);

*/
/*

	void callback_busnameacquiredcallbackHandler(GObject *, GDBusConnection*, gchar*, gpointer);

*/
/*

	void callback_busnameappearedcallbackHandler(GObject *, GDBusConnection*, gchar*, gchar*, gpointer);

*/
/*

	void callback_busnamelostcallbackHandler(GObject *, GDBusConnection*, gchar*, gpointer);

*/
/*

	void callback_busnamevanishedcallbackHandler(GObject *, GDBusConnection*, gchar*, gpointer);

*/
/*

	GVariant* callback_dbusinterfacegetpropertyfuncHandler(GObject *, GDBusConnection*, gchar*, gchar*, gchar*, gchar*, GError**, gpointer);

*/
/*

	void callback_dbusinterfacemethodcallfuncHandler(GObject *, GDBusConnection*, gchar*, gchar*, gchar*, gchar*, GVariant*, GDBusMethodInvocation*, gpointer);

*/
/*

	gboolean callback_dbusinterfacesetpropertyfuncHandler(GObject *, GDBusConnection*, gchar*, gchar*, gchar*, gchar*, GVariant*, GError**, gpointer);

*/
/*

	GDBusMessage* callback_dbusmessagefilterfunctionHandler(GObject *, GDBusConnection*, GDBusMessage*, gboolean, gpointer);

*/
/*

	void callback_dbussignalcallbackHandler(GObject *, GDBusConnection*, gchar*, gchar*, gchar*, gchar*, GVariant*, gpointer);

*/
import "C"

var callbackBusacquiredcallbackId int
var callbackBusacquiredcallbackMap = make(map[int]BusacquiredcallbackCallback)
var callbackBusacquiredcallbackLock sync.RWMutex

// BusacquiredcallbackCallback is a callback function for a 'BusAcquiredCallback' callback.
type BusacquiredcallbackCallback func(connection *DBusConnection, name string)

//export callback_busacquiredcallbackHandler
func callback_busacquiredcallbackHandler(_ *C.GObject, c_connection *C.GDBusConnection, c_name *C.gchar, c_user_data C.gpointer) {
	callbackBusacquiredcallbackLock.RLock()
	defer callbackBusacquiredcallbackLock.RUnlock()

	connection := DBusConnectionNewFromC(unsafe.Pointer(c_connection))

	name := C.GoString(c_name)

	index := int(uintptr(c_user_data))
	callback := callbackBusacquiredcallbackMap[index]
	callback(connection, name)
}

var callbackBusnameacquiredcallbackId int
var callbackBusnameacquiredcallbackMap = make(map[int]BusnameacquiredcallbackCallback)
var callbackBusnameacquiredcallbackLock sync.RWMutex

// BusnameacquiredcallbackCallback is a callback function for a 'BusNameAcquiredCallback' callback.
type BusnameacquiredcallbackCallback func(connection *DBusConnection, name string)

//export callback_busnameacquiredcallbackHandler
func callback_busnameacquiredcallbackHandler(_ *C.GObject, c_connection *C.GDBusConnection, c_name *C.gchar, c_user_data C.gpointer) {
	callbackBusnameacquiredcallbackLock.RLock()
	defer callbackBusnameacquiredcallbackLock.RUnlock()

	connection := DBusConnectionNewFromC(unsafe.Pointer(c_connection))

	name := C.GoString(c_name)

	index := int(uintptr(c_user_data))
	callback := callbackBusnameacquiredcallbackMap[index]
	callback(connection, name)
}

var callbackBusnameappearedcallbackId int
var callbackBusnameappearedcallbackMap = make(map[int]BusnameappearedcallbackCallback)
var callbackBusnameappearedcallbackLock sync.RWMutex

// BusnameappearedcallbackCallback is a callback function for a 'BusNameAppearedCallback' callback.
type BusnameappearedcallbackCallback func(connection *DBusConnection, name string, nameOwner string)

//export callback_busnameappearedcallbackHandler
func callback_busnameappearedcallbackHandler(_ *C.GObject, c_connection *C.GDBusConnection, c_name *C.gchar, c_name_owner *C.gchar, c_user_data C.gpointer) {
	callbackBusnameappearedcallbackLock.RLock()
	defer callbackBusnameappearedcallbackLock.RUnlock()

	connection := DBusConnectionNewFromC(unsafe.Pointer(c_connection))

	name := C.GoString(c_name)

	nameOwner := C.GoString(c_name_owner)

	index := int(uintptr(c_user_data))
	callback := callbackBusnameappearedcallbackMap[index]
	callback(connection, name, nameOwner)
}

var callbackBusnamelostcallbackId int
var callbackBusnamelostcallbackMap = make(map[int]BusnamelostcallbackCallback)
var callbackBusnamelostcallbackLock sync.RWMutex

// BusnamelostcallbackCallback is a callback function for a 'BusNameLostCallback' callback.
type BusnamelostcallbackCallback func(connection *DBusConnection, name string)

//export callback_busnamelostcallbackHandler
func callback_busnamelostcallbackHandler(_ *C.GObject, c_connection *C.GDBusConnection, c_name *C.gchar, c_user_data C.gpointer) {
	callbackBusnamelostcallbackLock.RLock()
	defer callbackBusnamelostcallbackLock.RUnlock()

	connection := DBusConnectionNewFromC(unsafe.Pointer(c_connection))

	name := C.GoString(c_name)

	index := int(uintptr(c_user_data))
	callback := callbackBusnamelostcallbackMap[index]
	callback(connection, name)
}

var callbackBusnamevanishedcallbackId int
var callbackBusnamevanishedcallbackMap = make(map[int]BusnamevanishedcallbackCallback)
var callbackBusnamevanishedcallbackLock sync.RWMutex

// BusnamevanishedcallbackCallback is a callback function for a 'BusNameVanishedCallback' callback.
type BusnamevanishedcallbackCallback func(connection *DBusConnection, name string)

//export callback_busnamevanishedcallbackHandler
func callback_busnamevanishedcallbackHandler(_ *C.GObject, c_connection *C.GDBusConnection, c_name *C.gchar, c_user_data C.gpointer) {
	callbackBusnamevanishedcallbackLock.RLock()
	defer callbackBusnamevanishedcallbackLock.RUnlock()

	connection := DBusConnectionNewFromC(unsafe.Pointer(c_connection))

	name := C.GoString(c_name)

	index := int(uintptr(c_user_data))
	callback := callbackBusnamevanishedcallbackMap[index]
	callback(connection, name)
}

var callbackDbusinterfacegetpropertyfuncId int
var callbackDbusinterfacegetpropertyfuncMap = make(map[int]DbusinterfacegetpropertyfuncCallback)
var callbackDbusinterfacegetpropertyfuncLock sync.RWMutex

// DbusinterfacegetpropertyfuncCallback is a callback function for a 'DBusInterfaceGetPropertyFunc' callback.
type DbusinterfacegetpropertyfuncCallback func(connection *DBusConnection, sender string, objectPath string, interfaceName string, propertyName string, error *glib.Error) *glib.Variant

//export callback_dbusinterfacegetpropertyfuncHandler
func callback_dbusinterfacegetpropertyfuncHandler(_ *C.GObject, c_connection *C.GDBusConnection, c_sender *C.gchar, c_object_path *C.gchar, c_interface_name *C.gchar, c_property_name *C.gchar, c_error *C.GError, c_user_data C.gpointer) **C.GVariant {
	callbackDbusinterfacegetpropertyfuncLock.RLock()
	defer callbackDbusinterfacegetpropertyfuncLock.RUnlock()

	connection := DBusConnectionNewFromC(unsafe.Pointer(c_connection))

	sender := C.GoString(c_sender)

	objectPath := C.GoString(c_object_path)

	interfaceName := C.GoString(c_interface_name)

	propertyName := C.GoString(c_property_name)

	error := glib.ErrorNewFromC(unsafe.Pointer(c_error))

	index := int(uintptr(c_user_data))
	callback := callbackDbusinterfacegetpropertyfuncMap[index]
	retGo := callback(connection, sender, objectPath, interfaceName, propertyName, error)
	retC :=
		(*C.GVariant)(retGo.ToC())
	return retC
}

var callbackDbusinterfacemethodcallfuncId int
var callbackDbusinterfacemethodcallfuncMap = make(map[int]DbusinterfacemethodcallfuncCallback)
var callbackDbusinterfacemethodcallfuncLock sync.RWMutex

// DbusinterfacemethodcallfuncCallback is a callback function for a 'DBusInterfaceMethodCallFunc' callback.
type DbusinterfacemethodcallfuncCallback func(connection *DBusConnection, sender string, objectPath string, interfaceName string, methodName string, parameters *glib.Variant, invocation *DBusMethodInvocation)

//export callback_dbusinterfacemethodcallfuncHandler
func callback_dbusinterfacemethodcallfuncHandler(_ *C.GObject, c_connection *C.GDBusConnection, c_sender *C.gchar, c_object_path *C.gchar, c_interface_name *C.gchar, c_method_name *C.gchar, c_parameters *C.GVariant, c_invocation *C.GDBusMethodInvocation, c_user_data C.gpointer) {
	callbackDbusinterfacemethodcallfuncLock.RLock()
	defer callbackDbusinterfacemethodcallfuncLock.RUnlock()

	connection := DBusConnectionNewFromC(unsafe.Pointer(c_connection))

	sender := C.GoString(c_sender)

	objectPath := C.GoString(c_object_path)

	interfaceName := C.GoString(c_interface_name)

	methodName := C.GoString(c_method_name)

	parameters := glib.VariantNewFromC(unsafe.Pointer(c_parameters))

	invocation := DBusMethodInvocationNewFromC(unsafe.Pointer(c_invocation))

	index := int(uintptr(c_user_data))
	callback := callbackDbusinterfacemethodcallfuncMap[index]
	callback(connection, sender, objectPath, interfaceName, methodName, parameters, invocation)
}

var callbackDbusinterfacesetpropertyfuncId int
var callbackDbusinterfacesetpropertyfuncMap = make(map[int]DbusinterfacesetpropertyfuncCallback)
var callbackDbusinterfacesetpropertyfuncLock sync.RWMutex

// DbusinterfacesetpropertyfuncCallback is a callback function for a 'DBusInterfaceSetPropertyFunc' callback.
type DbusinterfacesetpropertyfuncCallback func(connection *DBusConnection, sender string, objectPath string, interfaceName string, propertyName string, value *glib.Variant, error *glib.Error) bool

//export callback_dbusinterfacesetpropertyfuncHandler
func callback_dbusinterfacesetpropertyfuncHandler(_ *C.GObject, c_connection *C.GDBusConnection, c_sender *C.gchar, c_object_path *C.gchar, c_interface_name *C.gchar, c_property_name *C.gchar, c_value *C.GVariant, c_error *C.GError, c_user_data C.gpointer) C.gboolean {
	callbackDbusinterfacesetpropertyfuncLock.RLock()
	defer callbackDbusinterfacesetpropertyfuncLock.RUnlock()

	connection := DBusConnectionNewFromC(unsafe.Pointer(c_connection))

	sender := C.GoString(c_sender)

	objectPath := C.GoString(c_object_path)

	interfaceName := C.GoString(c_interface_name)

	propertyName := C.GoString(c_property_name)

	value := glib.VariantNewFromC(unsafe.Pointer(c_value))

	error := glib.ErrorNewFromC(unsafe.Pointer(c_error))

	index := int(uintptr(c_user_data))
	callback := callbackDbusinterfacesetpropertyfuncMap[index]
	retGo := callback(connection, sender, objectPath, interfaceName, propertyName, value, error)
	retC :=
		boolToGboolean(retGo)
	return retC
}

var callbackDbusmessagefilterfunctionId int
var callbackDbusmessagefilterfunctionMap = make(map[int]DbusmessagefilterfunctionCallback)
var callbackDbusmessagefilterfunctionLock sync.RWMutex

// DbusmessagefilterfunctionCallback is a callback function for a 'DBusMessageFilterFunction' callback.
type DbusmessagefilterfunctionCallback func(connection *DBusConnection, message *DBusMessage, incoming bool) *DBusMessage

//export callback_dbusmessagefilterfunctionHandler
func callback_dbusmessagefilterfunctionHandler(_ *C.GObject, c_connection *C.GDBusConnection, c_message *C.GDBusMessage, c_incoming C.gboolean, c_user_data C.gpointer) **C.GDBusMessage {
	callbackDbusmessagefilterfunctionLock.RLock()
	defer callbackDbusmessagefilterfunctionLock.RUnlock()

	connection := DBusConnectionNewFromC(unsafe.Pointer(c_connection))

	message := DBusMessageNewFromC(unsafe.Pointer(c_message))

	incoming := c_incoming == C.TRUE

	index := int(uintptr(c_user_data))
	callback := callbackDbusmessagefilterfunctionMap[index]
	retGo := callback(connection, message, incoming)
	retC :=
		(*C.GDBusMessage)(retGo.ToC())
	return retC
}

var callbackDbussignalcallbackId int
var callbackDbussignalcallbackMap = make(map[int]DbussignalcallbackCallback)
var callbackDbussignalcallbackLock sync.RWMutex

// DbussignalcallbackCallback is a callback function for a 'DBusSignalCallback' callback.
type DbussignalcallbackCallback func(connection *DBusConnection, senderName string, objectPath string, interfaceName string, signalName string, parameters *glib.Variant)

//export callback_dbussignalcallbackHandler
func callback_dbussignalcallbackHandler(_ *C.GObject, c_connection *C.GDBusConnection, c_sender_name *C.gchar, c_object_path *C.gchar, c_interface_name *C.gchar, c_signal_name *C.gchar, c_parameters *C.GVariant, c_user_data C.gpointer) {
	callbackDbussignalcallbackLock.RLock()
	defer callbackDbussignalcallbackLock.RUnlock()

	connection := DBusConnectionNewFromC(unsafe.Pointer(c_connection))

	senderName := C.GoString(c_sender_name)

	objectPath := C.GoString(c_object_path)

	interfaceName := C.GoString(c_interface_name)

	signalName := C.GoString(c_signal_name)

	parameters := glib.VariantNewFromC(unsafe.Pointer(c_parameters))

	index := int(uintptr(c_user_data))
	callback := callbackDbussignalcallbackMap[index]
	callback(connection, senderName, objectPath, interfaceName, signalName, parameters)
}

// Unsupported : callback DBusSubtreeDispatchFunc : unsupported parameter out_user_data : no type generator for gpointer (gpointer*) for param out_user_data

// Unsupported : callback DBusSubtreeEnumerateFunc : unsupported return array

// Unsupported : callback DBusSubtreeIntrospectFunc : return type : record with indirection level of 2
