// This is a generated file - DO NOT EDIT
// +build gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import (
	gobject "github.com/pekim/gobbi/lib/gobject"
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

	void listmodel_itemsChangedHandler(GObject *, guint, guint, guint, gpointer);

	static gulong ListModel_signal_connect_items_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "items-changed", G_CALLBACK(listmodel_itemsChangedHandler), data);
	}

*/
import "C"

type signalListModelItemsChangedDetail struct {
	callback  ListModelSignalItemsChangedCallback
	handlerID C.gulong
}

var signalListModelItemsChangedId int
var signalListModelItemsChangedMap = make(map[int]signalListModelItemsChangedDetail)
var signalListModelItemsChangedLock sync.Mutex

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
	position := uint32(c_position)

	removed := uint32(c_removed)

	added := uint32(c_added)

	index := int(uintptr(data))
	callback := signalListModelItemsChangedMap[index].callback
	callback(position, removed, added)
}

// GetItem is a wrapper around the C function g_list_model_get_item.
func (recv *ListModel) GetItem(position uint32) uintptr {
	c_position := (C.guint)(position)

	retC := C.g_list_model_get_item((*C.GListModel)(recv.native), c_position)
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

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
