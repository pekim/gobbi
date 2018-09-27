// This is a generated file - DO NOT EDIT
// +build gtk_2.10 gtk_2.12 gtk_2.14 gtk_2.16 gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.20 gtk_3.22

package gtk

import (
	glib "github.com/pekim/gobbi/lib/glib"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

// RecentManager is a wrapper around the C record GtkRecentManager.
type RecentManager struct {
	native *C.GtkRecentManager
	// Private : parent_instance
	// Private : priv
}

func RecentManagerNewFromC(u unsafe.Pointer) *RecentManager {
	c := (*C.GtkRecentManager)(u)
	if c == nil {
		return nil
	}

	g := &RecentManager{native: c}

	return g
}

func (recv *RecentManager) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// RecentManagerNew is a wrapper around the C function gtk_recent_manager_new.
func RecentManagerNew() *RecentManager {
	retC := C.gtk_recent_manager_new()
	retGo := RecentManagerNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AddFull is a wrapper around the C function gtk_recent_manager_add_full.
func (recv *RecentManager) AddFull(uri string, recentData *RecentData) bool {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	c_recent_data := (*C.GtkRecentData)(recentData.ToC())

	retC := C.gtk_recent_manager_add_full((*C.GtkRecentManager)(recv.native), c_uri, c_recent_data)
	retGo := retC == C.TRUE

	return retGo
}

// AddItem is a wrapper around the C function gtk_recent_manager_add_item.
func (recv *RecentManager) AddItem(uri string) bool {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	retC := C.gtk_recent_manager_add_item((*C.GtkRecentManager)(recv.native), c_uri)
	retGo := retC == C.TRUE

	return retGo
}

// GetItems is a wrapper around the C function gtk_recent_manager_get_items.
func (recv *RecentManager) GetItems() *glib.List {
	retC := C.gtk_recent_manager_get_items((*C.GtkRecentManager)(recv.native))
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// HasItem is a wrapper around the C function gtk_recent_manager_has_item.
func (recv *RecentManager) HasItem(uri string) bool {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	retC := C.gtk_recent_manager_has_item((*C.GtkRecentManager)(recv.native), c_uri)
	retGo := retC == C.TRUE

	return retGo
}

// LookupItem is a wrapper around the C function gtk_recent_manager_lookup_item.
func (recv *RecentManager) LookupItem(uri string) (*RecentInfo, error) {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	var cThrowableError *C.GError

	retC := C.gtk_recent_manager_lookup_item((*C.GtkRecentManager)(recv.native), c_uri, &cThrowableError)
	retGo := RecentInfoNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// MoveItem is a wrapper around the C function gtk_recent_manager_move_item.
func (recv *RecentManager) MoveItem(uri string, newUri string) (bool, error) {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	c_new_uri := C.CString(newUri)
	defer C.free(unsafe.Pointer(c_new_uri))

	var cThrowableError *C.GError

	retC := C.gtk_recent_manager_move_item((*C.GtkRecentManager)(recv.native), c_uri, c_new_uri, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// PurgeItems is a wrapper around the C function gtk_recent_manager_purge_items.
func (recv *RecentManager) PurgeItems() (int32, error) {
	var cThrowableError *C.GError

	retC := C.gtk_recent_manager_purge_items((*C.GtkRecentManager)(recv.native), &cThrowableError)
	retGo := (int32)(retC)

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// RemoveItem is a wrapper around the C function gtk_recent_manager_remove_item.
func (recv *RecentManager) RemoveItem(uri string) (bool, error) {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	var cThrowableError *C.GError

	retC := C.gtk_recent_manager_remove_item((*C.GtkRecentManager)(recv.native), c_uri, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}
