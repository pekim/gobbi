// This is a generated file - DO NOT EDIT
// +build gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import (
	glib "github.com/pekim/gobbi/lib/glib"
	gobject "github.com/pekim/gobbi/lib/gobject"
	"runtime"
	"sync"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
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

	void mountoperation_showUnmountProgressHandler(GObject *, gchar*, gint64, gint64, gpointer);

	static gulong MountOperation_signal_connect_show_unmount_progress(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "show-unmount-progress", G_CALLBACK(mountoperation_showUnmountProgressHandler), data);
	}

*/
import "C"

// GetDbusConnection is a wrapper around the C function g_application_get_dbus_connection.
func (recv *Application) GetDbusConnection() *DBusConnection {
	retC := C.g_application_get_dbus_connection((*C.GApplication)(recv.native))
	retGo := DBusConnectionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetDbusObjectPath is a wrapper around the C function g_application_get_dbus_object_path.
func (recv *Application) GetDbusObjectPath() string {
	retC := C.g_application_get_dbus_object_path((*C.GApplication)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetStdin is a wrapper around the C function g_application_command_line_get_stdin.
func (recv *ApplicationCommandLine) GetStdin() *InputStream {
	retC := C.g_application_command_line_get_stdin((*C.GApplicationCommandLine)(recv.native))
	retGo := InputStreamNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AllowMechanism is a wrapper around the C function g_dbus_auth_observer_allow_mechanism.
func (recv *DBusAuthObserver) AllowMechanism(mechanism string) bool {
	c_mechanism := C.CString(mechanism)
	defer C.free(unsafe.Pointer(c_mechanism))

	retC := C.g_dbus_auth_observer_allow_mechanism((*C.GDBusAuthObserver)(recv.native), c_mechanism)
	retGo := retC == C.TRUE

	return retGo
}

// GetLastSerial is a wrapper around the C function g_dbus_connection_get_last_serial.
func (recv *DBusConnection) GetLastSerial() uint32 {
	retC := C.g_dbus_connection_get_last_serial((*C.GDBusConnection)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// IsExported is a wrapper around the C function g_dbus_object_manager_server_is_exported.
func (recv *DBusObjectManagerServer) IsExported(object *DBusObjectSkeleton) bool {
	c_object := (*C.GDBusObjectSkeleton)(C.NULL)
	if object != nil {
		c_object = (*C.GDBusObjectSkeleton)(object.ToC())
	}

	retC := C.g_dbus_object_manager_server_is_exported((*C.GDBusObjectManagerServer)(recv.native), c_object)
	retGo := retC == C.TRUE

	return retGo
}

// GetStartupWmClass is a wrapper around the C function g_desktop_app_info_get_startup_wm_class.
func (recv *DesktopAppInfo) GetStartupWmClass() string {
	retC := C.g_desktop_app_info_get_startup_wm_class((*C.GDesktopAppInfo)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetSymbolicIcon is a wrapper around the C function g_file_info_get_symbolic_icon.
func (recv *FileInfo) GetSymbolicIcon() *Icon {
	retC := C.g_file_info_get_symbolic_icon((*C.GFileInfo)(recv.native))
	retGo := IconNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SetSymbolicIcon is a wrapper around the C function g_file_info_set_symbolic_icon.
func (recv *FileInfo) SetSymbolicIcon(icon *Icon) {
	c_icon := (*C.GIcon)(icon.ToC())

	C.g_file_info_set_symbolic_icon((*C.GFileInfo)(recv.native), c_icon)

	return
}

// ReadBytes is a wrapper around the C function g_input_stream_read_bytes.
func (recv *InputStream) ReadBytes(count uint64, cancellable *Cancellable) (*glib.Bytes, error) {
	c_count := (C.gsize)(count)

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_input_stream_read_bytes((*C.GInputStream)(recv.native), c_count, c_cancellable, &cThrowableError)
	retGo := glib.BytesNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_input_stream_read_bytes_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// ReadBytesFinish is a wrapper around the C function g_input_stream_read_bytes_finish.
func (recv *InputStream) ReadBytesFinish(result *AsyncResult) (*glib.Bytes, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_input_stream_read_bytes_finish((*C.GInputStream)(recv.native), c_result, &cThrowableError)
	retGo := glib.BytesNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// MemoryInputStreamNewFromBytes is a wrapper around the C function g_memory_input_stream_new_from_bytes.
func MemoryInputStreamNewFromBytes(bytes *glib.Bytes) *MemoryInputStream {
	c_bytes := (*C.GBytes)(C.NULL)
	if bytes != nil {
		c_bytes = (*C.GBytes)(bytes.ToC())
	}

	retC := C.g_memory_input_stream_new_from_bytes(c_bytes)
	retGPointer := (C.gpointer)(retC)
	nonFloatingRef := C.g_object_is_floating(retGPointer) == C.FALSE
	retGo := MemoryInputStreamNewFromC(unsafe.Pointer(retC))

	if nonFloatingRef {
		C.g_object_unref(retGPointer)
	}

	return retGo
}

// AddBytes is a wrapper around the C function g_memory_input_stream_add_bytes.
func (recv *MemoryInputStream) AddBytes(bytes *glib.Bytes) {
	c_bytes := (*C.GBytes)(C.NULL)
	if bytes != nil {
		c_bytes = (*C.GBytes)(bytes.ToC())
	}

	C.g_memory_input_stream_add_bytes((*C.GMemoryInputStream)(recv.native), c_bytes)

	return
}

// StealAsBytes is a wrapper around the C function g_memory_output_stream_steal_as_bytes.
func (recv *MemoryOutputStream) StealAsBytes() *glib.Bytes {
	retC := C.g_memory_output_stream_steal_as_bytes((*C.GMemoryOutputStream)(recv.native))
	retGo := glib.BytesNewFromC(unsafe.Pointer(retC))

	return retGo
}

// MenuItemNewFromModel is a wrapper around the C function g_menu_item_new_from_model.
func MenuItemNewFromModel(model *MenuModel, itemIndex int32) *MenuItem {
	c_model := (*C.GMenuModel)(C.NULL)
	if model != nil {
		c_model = (*C.GMenuModel)(model.ToC())
	}

	c_item_index := (C.gint)(itemIndex)

	retC := C.g_menu_item_new_from_model(c_model, c_item_index)
	retGPointer := (C.gpointer)(retC)
	nonFloatingRef := C.g_object_is_floating(retGPointer) == C.FALSE
	retGo := MenuItemNewFromC(unsafe.Pointer(retC))

	if nonFloatingRef {
		C.g_object_unref(retGPointer)
	}

	return retGo
}

// Unsupported : g_menu_item_get_attribute : unsupported parameter ... : varargs

// Unsupported : g_menu_item_get_attribute_value : unsupported parameter expected_type : Blacklisted record : GVariantType

// GetLink is a wrapper around the C function g_menu_item_get_link.
func (recv *MenuItem) GetLink(link string) *MenuModel {
	c_link := C.CString(link)
	defer C.free(unsafe.Pointer(c_link))

	retC := C.g_menu_item_get_link((*C.GMenuItem)(recv.native), c_link)
	retGo := MenuModelNewFromC(unsafe.Pointer(retC))

	return retGo
}

type signalMountOperationShowUnmountProgressDetail struct {
	callback  MountOperationSignalShowUnmountProgressCallback
	handlerID C.gulong
}

var signalMountOperationShowUnmountProgressId int
var signalMountOperationShowUnmountProgressMap = make(map[int]signalMountOperationShowUnmountProgressDetail)
var signalMountOperationShowUnmountProgressLock sync.RWMutex

// MountOperationSignalShowUnmountProgressCallback is a callback function for a 'show-unmount-progress' signal emitted from a MountOperation.
type MountOperationSignalShowUnmountProgressCallback func(message string, timeLeft int64, bytesLeft int64)

/*
ConnectShowUnmountProgress connects the callback to the 'show-unmount-progress' signal for the MountOperation.

The returned value represents the connection, and may be passed to DisconnectShowUnmountProgress to remove it.
*/
func (recv *MountOperation) ConnectShowUnmountProgress(callback MountOperationSignalShowUnmountProgressCallback) int {
	signalMountOperationShowUnmountProgressLock.Lock()
	defer signalMountOperationShowUnmountProgressLock.Unlock()

	signalMountOperationShowUnmountProgressId++
	instance := C.gpointer(recv.native)
	handlerID := C.MountOperation_signal_connect_show_unmount_progress(instance, C.gpointer(uintptr(signalMountOperationShowUnmountProgressId)))

	detail := signalMountOperationShowUnmountProgressDetail{callback, handlerID}
	signalMountOperationShowUnmountProgressMap[signalMountOperationShowUnmountProgressId] = detail

	return signalMountOperationShowUnmountProgressId
}

/*
DisconnectShowUnmountProgress disconnects a callback from the 'show-unmount-progress' signal for the MountOperation.

The connectionID should be a value returned from a call to ConnectShowUnmountProgress.
*/
func (recv *MountOperation) DisconnectShowUnmountProgress(connectionID int) {
	signalMountOperationShowUnmountProgressLock.Lock()
	defer signalMountOperationShowUnmountProgressLock.Unlock()

	detail, exists := signalMountOperationShowUnmountProgressMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalMountOperationShowUnmountProgressMap, connectionID)
}

//export mountoperation_showUnmountProgressHandler
func mountoperation_showUnmountProgressHandler(_ *C.GObject, c_message *C.gchar, c_time_left C.gint64, c_bytes_left C.gint64, data C.gpointer) {
	signalMountOperationShowUnmountProgressLock.RLock()
	defer signalMountOperationShowUnmountProgressLock.RUnlock()

	message := C.GoString(c_message)

	timeLeft := int64(c_time_left)

	bytesLeft := int64(c_bytes_left)

	index := int(uintptr(data))
	callback := signalMountOperationShowUnmountProgressMap[index].callback
	callback(message, timeLeft, bytesLeft)
}

// GetDestinationProtocol is a wrapper around the C function g_proxy_address_get_destination_protocol.
func (recv *ProxyAddress) GetDestinationProtocol() string {
	retC := C.g_proxy_address_get_destination_protocol((*C.GProxyAddress)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetUri is a wrapper around the C function g_proxy_address_get_uri.
func (recv *ProxyAddress) GetUri() string {
	retC := C.g_proxy_address_get_uri((*C.GProxyAddress)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// LookupRecords is a wrapper around the C function g_resolver_lookup_records.
func (recv *Resolver) LookupRecords(rrname string, recordType ResolverRecordType, cancellable *Cancellable) (*glib.List, error) {
	c_rrname := C.CString(rrname)
	defer C.free(unsafe.Pointer(c_rrname))

	c_record_type := (C.GResolverRecordType)(recordType)

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_resolver_lookup_records((*C.GResolver)(recv.native), c_rrname, c_record_type, c_cancellable, &cThrowableError)
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_resolver_lookup_records_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// LookupRecordsFinish is a wrapper around the C function g_resolver_lookup_records_finish.
func (recv *Resolver) LookupRecordsFinish(result *AsyncResult) (*glib.List, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_resolver_lookup_records_finish((*C.GResolver)(recv.native), c_result, &cThrowableError)
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// TestDBus is a wrapper around the C record GTestDBus.
type TestDBus struct {
	native *C.GTestDBus
}

func TestDBusNewFromC(u unsafe.Pointer) *TestDBus {
	c := (*C.GTestDBus)(u)
	if c == nil {
		return nil
	}

	g := &TestDBus{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *TestDBus) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *TestDBus) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this TestDBus with another TestDBus, and returns true if they represent the same GObject.
func (recv *TestDBus) Equals(other *TestDBus) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *TestDBus) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to TestDBus.
// Exercise care, as this is a potentially dangerous function if the Object is not a TestDBus.
func CastToTestDBus(object *gobject.Object) *TestDBus {
	return TestDBusNewFromC(object.ToC())
}

// TestDBusNew is a wrapper around the C function g_test_dbus_new.
func TestDBusNew(flags TestDBusFlags) *TestDBus {
	c_flags := (C.GTestDBusFlags)(flags)

	retC := C.g_test_dbus_new(c_flags)
	retGPointer := (C.gpointer)(retC)
	nonFloatingRef := C.g_object_is_floating(retGPointer) == C.FALSE
	retGo := TestDBusNewFromC(unsafe.Pointer(retC))

	if nonFloatingRef {
		C.g_object_unref(retGPointer)
	}

	return retGo
}

// AddServiceDir is a wrapper around the C function g_test_dbus_add_service_dir.
func (recv *TestDBus) AddServiceDir(path string) {
	c_path := C.CString(path)
	defer C.free(unsafe.Pointer(c_path))

	C.g_test_dbus_add_service_dir((*C.GTestDBus)(recv.native), c_path)

	return
}

// Down is a wrapper around the C function g_test_dbus_down.
func (recv *TestDBus) Down() {
	C.g_test_dbus_down((*C.GTestDBus)(recv.native))

	return
}

// GetBusAddress is a wrapper around the C function g_test_dbus_get_bus_address.
func (recv *TestDBus) GetBusAddress() string {
	retC := C.g_test_dbus_get_bus_address((*C.GTestDBus)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetFlags is a wrapper around the C function g_test_dbus_get_flags.
func (recv *TestDBus) GetFlags() TestDBusFlags {
	retC := C.g_test_dbus_get_flags((*C.GTestDBus)(recv.native))
	retGo := (TestDBusFlags)(retC)

	return retGo
}

// Stop is a wrapper around the C function g_test_dbus_stop.
func (recv *TestDBus) Stop() {
	C.g_test_dbus_stop((*C.GTestDBus)(recv.native))

	return
}

// Up is a wrapper around the C function g_test_dbus_up.
func (recv *TestDBus) Up() {
	C.g_test_dbus_up((*C.GTestDBus)(recv.native))

	return
}

// IsSame is a wrapper around the C function g_tls_certificate_is_same.
func (recv *TlsCertificate) IsSame(certTwo *TlsCertificate) bool {
	c_cert_two := (*C.GTlsCertificate)(C.NULL)
	if certTwo != nil {
		c_cert_two = (*C.GTlsCertificate)(certTwo.ToC())
	}

	retC := C.g_tls_certificate_is_same((*C.GTlsCertificate)(recv.native), c_cert_two)
	retGo := retC == C.TRUE

	return retGo
}
