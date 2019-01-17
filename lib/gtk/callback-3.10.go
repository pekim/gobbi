// This is a generated file - DO NOT EDIT
// +build gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import "sync"

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

var callbackListboxsortfuncId int
var callbackListboxsortfuncMap = make(map[int]ListboxsortfuncCallback)
var callbackListboxsortfuncLock sync.RWMutex

// ListboxsortfuncCallback is a callback function for a 'ListBoxSortFunc' callback.
type ListboxsortfuncCallback func(row1 *ListBoxRow, row2 *ListBoxRow) int32

var callbackListboxupdateheaderfuncId int
var callbackListboxupdateheaderfuncMap = make(map[int]ListboxupdateheaderfuncCallback)
var callbackListboxupdateheaderfuncLock sync.RWMutex

// ListboxupdateheaderfuncCallback is a callback function for a 'ListBoxUpdateHeaderFunc' callback.
type ListboxupdateheaderfuncCallback func(row *ListBoxRow, before *ListBoxRow)
