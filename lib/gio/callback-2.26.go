// This is a generated file - DO NOT EDIT
// +build gio_2.26 gio_2.28 gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import (
	glib "github.com/pekim/gobbi/lib/glib"
	"sync"
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

	void callback_busacquiredcallbackHandler(GObject *, GDBusConnection*, const gchar*, gpointer, gpointer);

*/
/*

	void callback_busnameacquiredcallbackHandler(GObject *, GDBusConnection*, const gchar*, gpointer, gpointer);

*/
/*

	void callback_busnameappearedcallbackHandler(GObject *, GDBusConnection*, const gchar*, const gchar*, gpointer, gpointer);

*/
/*

	void callback_busnamelostcallbackHandler(GObject *, GDBusConnection*, const gchar*, gpointer, gpointer);

*/
/*

	void callback_busnamevanishedcallbackHandler(GObject *, GDBusConnection*, const gchar*, gpointer, gpointer);

*/
/*

	GVariant* callback_dbusinterfacegetpropertyfuncHandler(GObject *, GDBusConnection*, const gchar*, const gchar*, const gchar*, const gchar*, GError**, gpointer, gpointer);

*/
/*

	void callback_dbusinterfacemethodcallfuncHandler(GObject *, GDBusConnection*, const gchar*, const gchar*, const gchar*, const gchar*, GVariant*, GDBusMethodInvocation*, gpointer, gpointer);

*/
/*

	gboolean callback_dbusinterfacesetpropertyfuncHandler(GObject *, GDBusConnection*, const gchar*, const gchar*, const gchar*, const gchar*, GVariant*, GError**, gpointer, gpointer);

*/
/*

	GDBusMessage* callback_dbusmessagefilterfunctionHandler(GObject *, GDBusConnection*, GDBusMessage*, gboolean, gpointer, gpointer);

*/
/*

	void callback_dbussignalcallbackHandler(GObject *, GDBusConnection*, const gchar*, const gchar*, const gchar*, const gchar*, GVariant*, gpointer, gpointer);

*/
import "C"

var callbackBusacquiredcallbackId int
var callbackBusacquiredcallbackMap = make(map[int]BusacquiredcallbackCallback)
var callbackBusacquiredcallbackLock sync.RWMutex

// BusacquiredcallbackCallback is a callback function for a 'BusAcquiredCallback' callback.
type BusacquiredcallbackCallback func(connection *DBusConnection, name string)

var callbackBusnameacquiredcallbackId int
var callbackBusnameacquiredcallbackMap = make(map[int]BusnameacquiredcallbackCallback)
var callbackBusnameacquiredcallbackLock sync.RWMutex

// BusnameacquiredcallbackCallback is a callback function for a 'BusNameAcquiredCallback' callback.
type BusnameacquiredcallbackCallback func(connection *DBusConnection, name string)

var callbackBusnameappearedcallbackId int
var callbackBusnameappearedcallbackMap = make(map[int]BusnameappearedcallbackCallback)
var callbackBusnameappearedcallbackLock sync.RWMutex

// BusnameappearedcallbackCallback is a callback function for a 'BusNameAppearedCallback' callback.
type BusnameappearedcallbackCallback func(connection *DBusConnection, name string, nameOwner string)

var callbackBusnamelostcallbackId int
var callbackBusnamelostcallbackMap = make(map[int]BusnamelostcallbackCallback)
var callbackBusnamelostcallbackLock sync.RWMutex

// BusnamelostcallbackCallback is a callback function for a 'BusNameLostCallback' callback.
type BusnamelostcallbackCallback func(connection *DBusConnection, name string)

var callbackBusnamevanishedcallbackId int
var callbackBusnamevanishedcallbackMap = make(map[int]BusnamevanishedcallbackCallback)
var callbackBusnamevanishedcallbackLock sync.RWMutex

// BusnamevanishedcallbackCallback is a callback function for a 'BusNameVanishedCallback' callback.
type BusnamevanishedcallbackCallback func(connection *DBusConnection, name string)

var callbackDbusinterfacegetpropertyfuncId int
var callbackDbusinterfacegetpropertyfuncMap = make(map[int]DbusinterfacegetpropertyfuncCallback)
var callbackDbusinterfacegetpropertyfuncLock sync.RWMutex

// DbusinterfacegetpropertyfuncCallback is a callback function for a 'DBusInterfaceGetPropertyFunc' callback.
type DbusinterfacegetpropertyfuncCallback func(connection *DBusConnection, sender string, objectPath string, interfaceName string, propertyName string, error *glib.Error) *glib.Variant

var callbackDbusinterfacemethodcallfuncId int
var callbackDbusinterfacemethodcallfuncMap = make(map[int]DbusinterfacemethodcallfuncCallback)
var callbackDbusinterfacemethodcallfuncLock sync.RWMutex

// DbusinterfacemethodcallfuncCallback is a callback function for a 'DBusInterfaceMethodCallFunc' callback.
type DbusinterfacemethodcallfuncCallback func(connection *DBusConnection, sender string, objectPath string, interfaceName string, methodName string, parameters *glib.Variant, invocation *DBusMethodInvocation)

var callbackDbusinterfacesetpropertyfuncId int
var callbackDbusinterfacesetpropertyfuncMap = make(map[int]DbusinterfacesetpropertyfuncCallback)
var callbackDbusinterfacesetpropertyfuncLock sync.RWMutex

// DbusinterfacesetpropertyfuncCallback is a callback function for a 'DBusInterfaceSetPropertyFunc' callback.
type DbusinterfacesetpropertyfuncCallback func(connection *DBusConnection, sender string, objectPath string, interfaceName string, propertyName string, value *glib.Variant, error *glib.Error) bool

var callbackDbusmessagefilterfunctionId int
var callbackDbusmessagefilterfunctionMap = make(map[int]DbusmessagefilterfunctionCallback)
var callbackDbusmessagefilterfunctionLock sync.RWMutex

// DbusmessagefilterfunctionCallback is a callback function for a 'DBusMessageFilterFunction' callback.
type DbusmessagefilterfunctionCallback func(connection *DBusConnection, message *DBusMessage, incoming bool) *DBusMessage

var callbackDbussignalcallbackId int
var callbackDbussignalcallbackMap = make(map[int]DbussignalcallbackCallback)
var callbackDbussignalcallbackLock sync.RWMutex

// DbussignalcallbackCallback is a callback function for a 'DBusSignalCallback' callback.
type DbussignalcallbackCallback func(connection *DBusConnection, senderName string, objectPath string, interfaceName string, signalName string, parameters *glib.Variant)

// Unsupported : callback DBusSubtreeDispatchFunc : unsupported parameter out_user_data : no type generator for gpointer (gpointer*) for param out_user_data

// Unsupported : callback DBusSubtreeEnumerateFunc : unsupported return array

// Unsupported : callback DBusSubtreeIntrospectFunc : return type : record with indirection level of 2
