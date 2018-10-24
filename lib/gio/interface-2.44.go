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

// GetItem is a wrapper around the C function g_list_model_get_item.
func (recv *ListModel) GetItem(position uint32) uintptr {
	c_position := (C.guint)(position)

	retC := C.g_list_model_get_item((*C.GListModel)(recv.native), c_position)
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_list_model_get_item_type : no return generator

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
