// This is a generated file - DO NOT EDIT
// +build gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import "sync"

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

// Blacklisted : g_list_model_get_item_type

// Blacklisted : g_list_model_get_n_items

// Blacklisted : g_list_model_get_object

// Blacklisted : g_list_model_items_changed

// Blacklisted : g_network_monitor_get_connectivity
