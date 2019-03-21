// This is a generated file - DO NOT EDIT
// +build gtk_2.2 gtk_2.4 gtk_2.6 gtk_2.8 gtk_2.10 gtk_2.12 gtk_2.14 gtk_2.16 gtk_2.18 gtk_2.20 gtk_2.22 gtk_2.24 gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

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

	gboolean callback_accelgroupfindfuncHandler(GObject *, GtkAccelKey*, GClosure*, gpointer);

*/
import "C"

var callbackAccelgroupfindfuncId int
var callbackAccelgroupfindfuncMap = make(map[int]AccelgroupfindfuncCallback)
var callbackAccelgroupfindfuncLock sync.RWMutex

// AccelgroupfindfuncCallback is a callback function for a 'AccelGroupFindFunc' callback.
type AccelgroupfindfuncCallback func(key *AccelKey, closure *gobject.Closure) bool

//export callback_accelgroupfindfuncHandler
func callback_accelgroupfindfuncHandler(_ *C.GObject, c_key *C.GtkAccelKey, c_closure *C.GClosure, c_data C.gpointer) C.gboolean {
	callbackAccelgroupfindfuncLock.RLock()
	defer callbackAccelgroupfindfuncLock.RUnlock()

	key := AccelKeyNewFromC(unsafe.Pointer(c_key))

	closure := gobject.ClosureNewFromC(unsafe.Pointer(c_closure))

	index := int(uintptr(c_data))
	callback := callbackAccelgroupfindfuncMap[index]
	retGo := callback(key, closure)
	retC :=
		boolToGboolean(retGo)
	return retC
}

// Unsupported : callback ColorSelectionChangePaletteWithScreenFunc : no [user_]data param

// Unsupported : callback ModuleDisplayInitFunc : no [user_]data param
