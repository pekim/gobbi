// This is a generated file - DO NOT EDIT
// +build gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	gobject "github.com/pekim/gobbi/lib/gobject"
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

	GtkWidget* callback_listboxcreatewidgetfuncHandler(GObject *, gpointer, gpointer, gpointer);

*/
import "C"

var callbackListboxcreatewidgetfuncId int
var callbackListboxcreatewidgetfuncMap = make(map[int]ListboxcreatewidgetfuncCallback)
var callbackListboxcreatewidgetfuncLock sync.RWMutex

// ListboxcreatewidgetfuncCallback is a callback function for a 'ListBoxCreateWidgetFunc' callback.
type ListboxcreatewidgetfuncCallback func(item *gobject.Object) *Widget

//export callback_listboxcreatewidgetfuncHandler
func callback_listboxcreatewidgetfuncHandler(_ *C.GObject, c_item *C.GObject, data C.gpointer) **C.GtkWidget {
	callbackListboxcreatewidgetfuncLock.RLock()
	defer callbackListboxcreatewidgetfuncLock.RUnlock()

	item := gobject.ObjectNewFromC(unsafe.Pointer(c_item))

	index := int(uintptr(data))
	callback := callbackListboxcreatewidgetfuncMap[index].callback
	retGo := callback(item, userData)
	retC :=
		(*C.GtkWidget)(retGo.ToC())
	return retC
}
