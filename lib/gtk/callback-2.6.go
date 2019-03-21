// This is a generated file - DO NOT EDIT
// +build gtk_2.6 gtk_2.8 gtk_2.10 gtk_2.12 gtk_2.14 gtk_2.16 gtk_2.18 gtk_2.20 gtk_2.22 gtk_2.24 gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	gdkpixbuf "github.com/pekim/gobbi/lib/gdkpixbuf"
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

	void callback_clipboardimagereceivedfuncHandler(GObject *, GtkClipboard*, GdkPixbuf*, gpointer);

*/
import "C"

var callbackClipboardimagereceivedfuncId int
var callbackClipboardimagereceivedfuncMap = make(map[int]ClipboardimagereceivedfuncCallback)
var callbackClipboardimagereceivedfuncLock sync.RWMutex

// ClipboardimagereceivedfuncCallback is a callback function for a 'ClipboardImageReceivedFunc' callback.
type ClipboardimagereceivedfuncCallback func(clipboard *Clipboard, pixbuf *gdkpixbuf.Pixbuf)

//export callback_clipboardimagereceivedfuncHandler
func callback_clipboardimagereceivedfuncHandler(_ *C.GObject, c_clipboard *C.GtkClipboard, c_pixbuf *C.GdkPixbuf, c_data C.gpointer) {
	callbackClipboardimagereceivedfuncLock.RLock()
	defer callbackClipboardimagereceivedfuncLock.RUnlock()

	clipboard := ClipboardNewFromC(unsafe.Pointer(c_clipboard))

	pixbuf := gdkpixbuf.PixbufNewFromC(unsafe.Pointer(c_pixbuf))

	index := int(uintptr(c_data))
	callback := callbackClipboardimagereceivedfuncMap[index]
	callback(clipboard, pixbuf)
}
