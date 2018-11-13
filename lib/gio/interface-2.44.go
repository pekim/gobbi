// This is a generated file - DO NOT EDIT
// +build gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import (
	gobject "github.com/pekim/gobbi/lib/gobject"
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
import "C"

// Unsupported signal 'items-changed' for ListModel : unsupported parameter position : type guint :

// Get the item at @position. If @position is greater than the number of
// items in @list, %NULL is returned.
//
// %NULL is never returned for an index that is smaller than the length
// of the list.  See g_list_model_get_n_items().
/*

C function : g_list_model_get_item
*/
func (recv *ListModel) GetItem(position uint32) uintptr {
	c_position := (C.guint)(position)

	retC := C.g_list_model_get_item((*C.GListModel)(recv.native), c_position)
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// Gets the type of the items in @list. All items returned from
// g_list_model_get_type() are of that type or a subtype, or are an
// implementation of that interface.
//
// The item type of a #GListModel can not change during the life of the
// model.
/*

C function : g_list_model_get_item_type
*/
func (recv *ListModel) GetItemType() gobject.Type {
	retC := C.g_list_model_get_item_type((*C.GListModel)(recv.native))
	retGo := (gobject.Type)(retC)

	return retGo
}

// Gets the number of items in @list.
//
// Depending on the model implementation, calling this function may be
// less efficient than iterating the list with increasing values for
// @position until g_list_model_get_item() returns %NULL.
/*

C function : g_list_model_get_n_items
*/
func (recv *ListModel) GetNItems() uint32 {
	retC := C.g_list_model_get_n_items((*C.GListModel)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// Get the item at @position. If @position is greater than the number of
// items in @list, %NULL is returned.
//
// %NULL is never returned for an index that is smaller than the length
// of the list.  See g_list_model_get_n_items().
/*

C function : g_list_model_get_object
*/
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

// Emits the #GListModel::items-changed signal on @list.
//
// This function should only be called by classes implementing
// #GListModel. It has to be called after the internal representation
// of @list has been updated, because handlers connected to this signal
// might query the new state of the list.
//
// Implementations must only make changes to the model (as visible to
// its consumer) in places that will not cause problems for that
// consumer.  For models that are driven directly by a write API (such
// as #GListStore), changes can be reported in response to uses of that
// API.  For models that represent remote data, changes should only be
// made from a fresh mainloop dispatch.  It is particularly not
// permitted to make changes in response to a call to the #GListModel
// consumer API.
//
// Stated another way: in general, it is assumed that code making a
// series of accesses to the model via the API, without returning to the
// mainloop, and without calling other code, will continue to view the
// same contents of the model.
/*

C function : g_list_model_items_changed
*/
func (recv *ListModel) ItemsChanged(position uint32, removed uint32, added uint32) {
	c_position := (C.guint)(position)

	c_removed := (C.guint)(removed)

	c_added := (C.guint)(added)

	C.g_list_model_items_changed((*C.GListModel)(recv.native), c_position, c_removed, c_added)

	return
}

// Gets a more detailed networking state than
// g_network_monitor_get_network_available().
//
// If #GNetworkMonitor:network-available is %FALSE, then the
// connectivity state will be %G_NETWORK_CONNECTIVITY_LOCAL.
//
// If #GNetworkMonitor:network-available is %TRUE, then the
// connectivity state will be %G_NETWORK_CONNECTIVITY_FULL (if there
// is full Internet connectivity), %G_NETWORK_CONNECTIVITY_LIMITED (if
// the host has a default route, but appears to be unable to actually
// reach the full Internet), or %G_NETWORK_CONNECTIVITY_PORTAL (if the
// host is trapped behind a "captive portal" that requires some sort
// of login or acknowledgement before allowing full Internet access).
//
// Note that in the case of %G_NETWORK_CONNECTIVITY_LIMITED and
// %G_NETWORK_CONNECTIVITY_PORTAL, it is possible that some sites are
// reachable but others are not. In this case, applications can
// attempt to connect to remote servers, but should gracefully fall
// back to their "offline" behavior if the connection attempt fails.
/*

C function : g_network_monitor_get_connectivity
*/
func (recv *NetworkMonitor) GetConnectivity() NetworkConnectivity {
	retC := C.g_network_monitor_get_connectivity((*C.GNetworkMonitor)(recv.native))
	retGo := (NetworkConnectivity)(retC)

	return retGo
}
