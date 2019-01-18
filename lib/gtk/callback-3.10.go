// This is a generated file - DO NOT EDIT
// +build gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	"sync"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
/*

	gboolean callback_listboxfilterfuncHandler(GObject *, GtkListBoxRow*, gpointer, gpointer);

*/
/*

	gint callback_listboxsortfuncHandler(GObject *, GtkListBoxRow*, GtkListBoxRow*, gpointer, gpointer);

*/
/*

	void callback_listboxupdateheaderfuncHandler(GObject *, GtkListBoxRow*, GtkListBoxRow*, gpointer, gpointer);

*/
import "C"

var callbackListboxfilterfuncId int
var callbackListboxfilterfuncMap = make(map[int]ListboxfilterfuncCallback)
var callbackListboxfilterfuncLock sync.RWMutex

// ListboxfilterfuncCallback is a callback function for a 'ListBoxFilterFunc' callback.
type ListboxfilterfuncCallback func(row *ListBoxRow) bool

//export callback_listboxfilterfuncHandler
func callback_listboxfilterfuncHandler(_ *C.GObject, c_row *C.GtkListBoxRow, data C.gpointer) C.gboolean {
	callbackListboxfilterfuncLock.RLock()
	defer callbackListboxfilterfuncLock.RUnlock()

	row := ListBoxRowNewFromC(unsafe.Pointer(c_row))

	index := int(uintptr(data))
	callback := callbackListboxfilterfuncMap[index].callback
	retGo := callback(row, userData)
	retC :=
		boolToGboolean(retGo)
	return retC
}

var callbackListboxsortfuncId int
var callbackListboxsortfuncMap = make(map[int]ListboxsortfuncCallback)
var callbackListboxsortfuncLock sync.RWMutex

// ListboxsortfuncCallback is a callback function for a 'ListBoxSortFunc' callback.
type ListboxsortfuncCallback func(row1 *ListBoxRow, row2 *ListBoxRow) int32

//export callback_listboxsortfuncHandler
func callback_listboxsortfuncHandler(_ *C.GObject, c_row1 *C.GtkListBoxRow, c_row2 *C.GtkListBoxRow, data C.gpointer) C.gint {
	callbackListboxsortfuncLock.RLock()
	defer callbackListboxsortfuncLock.RUnlock()

	row1 := ListBoxRowNewFromC(unsafe.Pointer(c_row1))

	row2 := ListBoxRowNewFromC(unsafe.Pointer(c_row2))

	index := int(uintptr(data))
	callback := callbackListboxsortfuncMap[index].callback
	retGo := callback(row1, row2, userData)
	retC :=
		(C.gint)(retGo)
	return retC
}

var callbackListboxupdateheaderfuncId int
var callbackListboxupdateheaderfuncMap = make(map[int]ListboxupdateheaderfuncCallback)
var callbackListboxupdateheaderfuncLock sync.RWMutex

// ListboxupdateheaderfuncCallback is a callback function for a 'ListBoxUpdateHeaderFunc' callback.
type ListboxupdateheaderfuncCallback func(row *ListBoxRow, before *ListBoxRow)

//export callback_listboxupdateheaderfuncHandler
func callback_listboxupdateheaderfuncHandler(_ *C.GObject, c_row *C.GtkListBoxRow, c_before *C.GtkListBoxRow, data C.gpointer) {
	callbackListboxupdateheaderfuncLock.RLock()
	defer callbackListboxupdateheaderfuncLock.RUnlock()

	row := ListBoxRowNewFromC(unsafe.Pointer(c_row))

	before := ListBoxRowNewFromC(unsafe.Pointer(c_before))

	index := int(uintptr(data))
	callback := callbackListboxupdateheaderfuncMap[index].callback
	callback(row, before, userData)
}
