// This is a generated file - DO NOT EDIT
// +build gtk_2.10 gtk_2.12 gtk_2.14 gtk_2.16 gtk_2.18 gtk_2.20 gtk_2.22 gtk_2.24 gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	gdk "github.com/pekim/gobbi/lib/gdk"
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

	void callback_clipboardrichtextreceivedfuncHandler(GObject *, GtkClipboard*, GdkAtom, guint8*, gsize, gpointer);

*/
import "C"

var callbackClipboardrichtextreceivedfuncId int
var callbackClipboardrichtextreceivedfuncMap = make(map[int]ClipboardrichtextreceivedfuncCallback)
var callbackClipboardrichtextreceivedfuncLock sync.RWMutex

// ClipboardrichtextreceivedfuncCallback is a callback function for a 'ClipboardRichTextReceivedFunc' callback.
type ClipboardrichtextreceivedfuncCallback func(clipboard *Clipboard, format *gdk.Atom, text string, length uint64)

//export callback_clipboardrichtextreceivedfuncHandler
func callback_clipboardrichtextreceivedfuncHandler(_ *C.GObject, c_clipboard *C.GtkClipboard, c_format *C.GdkAtom, c_text *C.guint8, c_length C.gsize, c_data C.gpointer) {
	callbackClipboardrichtextreceivedfuncLock.RLock()
	defer callbackClipboardrichtextreceivedfuncLock.RUnlock()

	clipboard := ClipboardNewFromC(unsafe.Pointer(c_clipboard))

	format := gdk.AtomNewFromC(unsafe.Pointer(c_format))

	text := C.GoString(c_text)

	length := uint64(c_length)

	index := int(uintptr(c_data))
	callback := callbackClipboardrichtextreceivedfuncMap[index]
	callback(clipboard, format, text, length)
}
