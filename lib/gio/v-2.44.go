// This is a generated file - DO NOT EDIT
// +build gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import (
	glib "github.com/pekim/gobbi/lib/glib"
	gobject "github.com/pekim/gobbi/lib/gobject"
	"runtime"
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
// #include <gio/gnetworking.h>
// #include <stdlib.h>
/*

	void listmodel_itemsChangedHandler(GObject *, guint, guint, guint, gpointer);

	static gulong ListModel_signal_connect_items_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "items-changed", G_CALLBACK(listmodel_itemsChangedHandler), data);
	}

*/
import "C"

// BindBusyProperty is a wrapper around the C function g_application_bind_busy_property.
func (recv *Application) BindBusyProperty(object *gobject.Object, property string) {
	c_object := (C.gpointer)(C.NULL)
	if object != nil {
		c_object = (C.gpointer)(object.ToC())
	}

	c_property := C.CString(property)
	defer C.free(unsafe.Pointer(c_property))

	C.g_application_bind_busy_property((*C.GApplication)(recv.native), c_object, c_property)

	return
}

// GetIsBusy is a wrapper around the C function g_application_get_is_busy.
func (recv *Application) GetIsBusy() bool {
	retC := C.g_application_get_is_busy((*C.GApplication)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// UnbindBusyProperty is a wrapper around the C function g_application_unbind_busy_property.
func (recv *Application) UnbindBusyProperty(object *gobject.Object, property string) {
	c_object := (C.gpointer)(C.NULL)
	if object != nil {
		c_object = (C.gpointer)(object.ToC())
	}

	c_property := C.CString(property)
	defer C.free(unsafe.Pointer(c_property))

	C.g_application_unbind_busy_property((*C.GApplication)(recv.native), c_object, c_property)

	return
}

// Iterate is a wrapper around the C function g_file_enumerator_iterate.
func (recv *FileEnumerator) Iterate(cancellable *Cancellable) (bool, *FileInfo, *File, error) {
	var c_out_info *C.GFileInfo

	var c_out_child *C.GFile

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_file_enumerator_iterate((*C.GFileEnumerator)(recv.native), &c_out_info, &c_out_child, c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	outInfo := FileInfoNewFromC(unsafe.Pointer(c_out_info))

	outChild := FileNewFromC(unsafe.Pointer(c_out_child))

	return retGo, outInfo, outChild, goError
}

// Unsupported : g_input_stream_read_all_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// ReadAllFinish is a wrapper around the C function g_input_stream_read_all_finish.
func (recv *InputStream) ReadAllFinish(result *AsyncResult) (bool, uint64, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var c_bytes_read C.gsize

	var cThrowableError *C.GError

	retC := C.g_input_stream_read_all_finish((*C.GInputStream)(recv.native), c_result, &c_bytes_read, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	bytesRead := (uint64)(c_bytes_read)

	return retGo, bytesRead, goError
}

// ListStoreNew is a wrapper around the C function g_list_store_new.
func ListStoreNew(itemType gobject.Type) *ListStore {
	c_item_type := (C.GType)(itemType)

	retC := C.g_list_store_new(c_item_type)
	retGo := ListStoreNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// Append is a wrapper around the C function g_list_store_append.
func (recv *ListStore) Append(item *gobject.Object) {
	c_item := (C.gpointer)(C.NULL)
	if item != nil {
		c_item = (C.gpointer)(item.ToC())
	}

	C.g_list_store_append((*C.GListStore)(recv.native), c_item)

	return
}

// Insert is a wrapper around the C function g_list_store_insert.
func (recv *ListStore) Insert(position uint32, item *gobject.Object) {
	c_position := (C.guint)(position)

	c_item := (C.gpointer)(C.NULL)
	if item != nil {
		c_item = (C.gpointer)(item.ToC())
	}

	C.g_list_store_insert((*C.GListStore)(recv.native), c_position, c_item)

	return
}

// Unsupported : g_list_store_insert_sorted : unsupported parameter compare_func : no type generator for GLib.CompareDataFunc (GCompareDataFunc) for param compare_func

// Remove is a wrapper around the C function g_list_store_remove.
func (recv *ListStore) Remove(position uint32) {
	c_position := (C.guint)(position)

	C.g_list_store_remove((*C.GListStore)(recv.native), c_position)

	return
}

// RemoveAll is a wrapper around the C function g_list_store_remove_all.
func (recv *ListStore) RemoveAll() {
	C.g_list_store_remove_all((*C.GListStore)(recv.native))

	return
}

// Unsupported : g_list_store_splice : unsupported parameter additions :

// NetworkAddressNewLoopback is a wrapper around the C function g_network_address_new_loopback.
func NetworkAddressNewLoopback(port uint16) *NetworkAddress {
	c_port := (C.guint16)(port)

	retC := C.g_network_address_new_loopback(c_port)
	retGo := NetworkAddressNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// Unsupported : g_output_stream_write_all_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// WriteAllFinish is a wrapper around the C function g_output_stream_write_all_finish.
func (recv *OutputStream) WriteAllFinish(result *AsyncResult) (bool, uint64, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var c_bytes_written C.gsize

	var cThrowableError *C.GError

	retC := C.g_output_stream_write_all_finish((*C.GOutputStream)(recv.native), c_result, &c_bytes_written, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	bytesWritten := (uint64)(c_bytes_written)

	return retGo, bytesWritten, goError
}

// SetStateHint is a wrapper around the C function g_simple_action_set_state_hint.
func (recv *SimpleAction) SetStateHint(stateHint *glib.Variant) {
	c_state_hint := (*C.GVariant)(C.NULL)
	if stateHint != nil {
		c_state_hint = (*C.GVariant)(stateHint.ToC())
	}

	C.g_simple_action_set_state_hint((*C.GSimpleAction)(recv.native), c_state_hint)

	return
}

// SimpleIOStream is a wrapper around the C record GSimpleIOStream.
type SimpleIOStream struct {
	native *C.GSimpleIOStream
}

func SimpleIOStreamNewFromC(u unsafe.Pointer) *SimpleIOStream {
	c := (*C.GSimpleIOStream)(u)
	if c == nil {
		return nil
	}

	g := &SimpleIOStream{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *SimpleIOStream) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *SimpleIOStream) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this SimpleIOStream with another SimpleIOStream, and returns true if they represent the same GObject.
func (recv *SimpleIOStream) Equals(other *SimpleIOStream) bool {
	return other.ToC() == recv.ToC()
}

// IOStream upcasts to *IOStream
func (recv *SimpleIOStream) IOStream() *IOStream {
	return IOStreamNewFromC(unsafe.Pointer(recv.native))
}

// Object upcasts to *Object
func (recv *SimpleIOStream) Object() *gobject.Object {
	return recv.IOStream().Object()
}

// CastToWidget down casts any arbitrary Object to SimpleIOStream.
// Exercise care, as this is a potentially dangerous function if the Object is not a SimpleIOStream.
func CastToSimpleIOStream(object *gobject.Object) *SimpleIOStream {
	return SimpleIOStreamNewFromC(object.ToC())
}

// SimpleIOStreamNew is a wrapper around the C function g_simple_io_stream_new.
func SimpleIOStreamNew(inputStream *InputStream, outputStream *OutputStream) *SimpleIOStream {
	c_input_stream := (*C.GInputStream)(C.NULL)
	if inputStream != nil {
		c_input_stream = (*C.GInputStream)(inputStream.ToC())
	}

	c_output_stream := (*C.GOutputStream)(C.NULL)
	if outputStream != nil {
		c_output_stream = (*C.GOutputStream)(outputStream.ToC())
	}

	retC := C.g_simple_io_stream_new(c_input_stream, c_output_stream)
	retGo := SimpleIOStreamNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// Unsupported : g_socket_send_messages : unsupported parameter messages :

// GetCompleted is a wrapper around the C function g_task_get_completed.
func (recv *Task) GetCompleted() bool {
	retC := C.g_task_get_completed((*C.GTask)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// UnixMountMonitorGet is a wrapper around the C function g_unix_mount_monitor_get.
func UnixMountMonitorGet() *UnixMountMonitor {
	retC := C.g_unix_mount_monitor_get()
	retGo := UnixMountMonitorNewFromC(unsafe.Pointer(retC))

	return retGo
}

type NetworkConnectivity C.GNetworkConnectivity

const (
	NETWORK_CONNECTIVITY_LOCAL   NetworkConnectivity = 1
	NETWORK_CONNECTIVITY_LIMITED NetworkConnectivity = 2
	NETWORK_CONNECTIVITY_PORTAL  NetworkConnectivity = 3
	NETWORK_CONNECTIVITY_FULL    NetworkConnectivity = 4
)

type signalListModelItemsChangedDetail struct {
	callback  ListModelSignalItemsChangedCallback
	handlerID C.gulong
}

var signalListModelItemsChangedId int
var signalListModelItemsChangedMap = make(map[int]signalListModelItemsChangedDetail)
var signalListModelItemsChangedLock sync.RWMutex

// ListModelSignalItemsChangedCallback is a callback function for a 'items-changed' signal emitted from a ListModel.
type ListModelSignalItemsChangedCallback func(position uint32, removed uint32, added uint32)

/*
ConnectItemsChanged connects the callback to the 'items-changed' signal for the ListModel.

The returned value represents the connection, and may be passed to DisconnectItemsChanged to remove it.
*/
func (recv *ListModel) ConnectItemsChanged(callback ListModelSignalItemsChangedCallback) int {
	signalListModelItemsChangedLock.Lock()
	defer signalListModelItemsChangedLock.Unlock()

	signalListModelItemsChangedId++
	instance := C.gpointer(recv.native)
	handlerID := C.ListModel_signal_connect_items_changed(instance, C.gpointer(uintptr(signalListModelItemsChangedId)))

	detail := signalListModelItemsChangedDetail{callback, handlerID}
	signalListModelItemsChangedMap[signalListModelItemsChangedId] = detail

	return signalListModelItemsChangedId
}

/*
DisconnectItemsChanged disconnects a callback from the 'items-changed' signal for the ListModel.

The connectionID should be a value returned from a call to ConnectItemsChanged.
*/
func (recv *ListModel) DisconnectItemsChanged(connectionID int) {
	signalListModelItemsChangedLock.Lock()
	defer signalListModelItemsChangedLock.Unlock()

	detail, exists := signalListModelItemsChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalListModelItemsChangedMap, connectionID)
}

//export listmodel_itemsChangedHandler
func listmodel_itemsChangedHandler(_ *C.GObject, c_position C.guint, c_removed C.guint, c_added C.guint, data C.gpointer) {
	signalListModelItemsChangedLock.RLock()
	defer signalListModelItemsChangedLock.RUnlock()

	position := uint32(c_position)

	removed := uint32(c_removed)

	added := uint32(c_added)

	index := int(uintptr(data))
	callback := signalListModelItemsChangedMap[index].callback
	callback(position, removed, added)
}

// Unsupported : g_list_model_get_item : no return generator

// GetItemType is a wrapper around the C function g_list_model_get_item_type.
func (recv *ListModel) GetItemType() gobject.Type {
	retC := C.g_list_model_get_item_type((*C.GListModel)(recv.native))
	retGo := (gobject.Type)(retC)

	return retGo
}

// GetNItems is a wrapper around the C function g_list_model_get_n_items.
func (recv *ListModel) GetNItems() uint32 {
	retC := C.g_list_model_get_n_items((*C.GListModel)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// GetObject is a wrapper around the C function g_list_model_get_object.
func (recv *ListModel) GetObject(position uint32) *gobject.Object {
	c_position := (C.guint)(position)

	retC := C.g_list_model_get_object((*C.GListModel)(recv.native), c_position)
	var retGo (*gobject.Object)
	if retC == nil {
		retGo = nil
	} else {
		retGo = gobject.ObjectNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// ItemsChanged is a wrapper around the C function g_list_model_items_changed.
func (recv *ListModel) ItemsChanged(position uint32, removed uint32, added uint32) {
	c_position := (C.guint)(position)

	c_removed := (C.guint)(removed)

	c_added := (C.guint)(added)

	C.g_list_model_items_changed((*C.GListModel)(recv.native), c_position, c_removed, c_added)

	return
}

// GetConnectivity is a wrapper around the C function g_network_monitor_get_connectivity.
func (recv *NetworkMonitor) GetConnectivity() NetworkConnectivity {
	retC := C.g_network_monitor_get_connectivity((*C.GNetworkMonitor)(recv.native))
	retGo := (NetworkConnectivity)(retC)

	return retGo
}

// ListModelInterface is a wrapper around the C record GListModelInterface.
type ListModelInterface struct {
	native *C.GListModelInterface
	// g_iface : record
	// no type for get_item_type
	// no type for get_n_items
	// no type for get_item
}

func ListModelInterfaceNewFromC(u unsafe.Pointer) *ListModelInterface {
	c := (*C.GListModelInterface)(u)
	if c == nil {
		return nil
	}

	g := &ListModelInterface{native: c}

	return g
}

func (recv *ListModelInterface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ListModelInterface with another ListModelInterface, and returns true if they represent the same GObject.
func (recv *ListModelInterface) Equals(other *ListModelInterface) bool {
	return other.ToC() == recv.ToC()
}

// OutputMessage is a wrapper around the C record GOutputMessage.
type OutputMessage struct {
	native *C.GOutputMessage
	// address : record
	// vectors : record
	NumVectors uint32
	BytesSent  uint32
	// no type for control_messages
	NumControlMessages uint32
}

func OutputMessageNewFromC(u unsafe.Pointer) *OutputMessage {
	c := (*C.GOutputMessage)(u)
	if c == nil {
		return nil
	}

	g := &OutputMessage{
		BytesSent:          (uint32)(c.bytes_sent),
		NumControlMessages: (uint32)(c.num_control_messages),
		NumVectors:         (uint32)(c.num_vectors),
		native:             c,
	}

	return g
}

func (recv *OutputMessage) ToC() unsafe.Pointer {
	recv.native.num_vectors =
		(C.guint)(recv.NumVectors)
	recv.native.bytes_sent =
		(C.guint)(recv.BytesSent)
	recv.native.num_control_messages =
		(C.guint)(recv.NumControlMessages)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this OutputMessage with another OutputMessage, and returns true if they represent the same GObject.
func (recv *OutputMessage) Equals(other *OutputMessage) bool {
	return other.ToC() == recv.ToC()
}

// ListChildren is a wrapper around the C function g_settings_schema_list_children.
func (recv *SettingsSchema) ListChildren() []string {
	retC := C.g_settings_schema_list_children((*C.GSettingsSchema)(recv.native))
	retGo := []string{}
	for p := retC; *p != nil; p = (**C.char)(C.gpointer((uintptr(C.gpointer(p)) + uintptr(C.sizeof_gpointer)))) {
		s := C.GoString(*p)
		retGo = append(retGo, s)
	}
	defer C.g_strfreev(retC)

	return retGo
}

// GetName is a wrapper around the C function g_settings_schema_key_get_name.
func (recv *SettingsSchemaKey) GetName() string {
	retC := C.g_settings_schema_key_get_name((*C.GSettingsSchemaKey)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}
