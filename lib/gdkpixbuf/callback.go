// This is a generated file - DO NOT EDIT

package gdkpixbuf

import "sync"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gdk-pixbuf/gdk-pixbuf.h>
// #include <stdlib.h>
/*

	void callback_pixbufdestroynotifyHandler(GObject *, guchar*, gpointer, gpointer);

*/
import "C"

var callbackPixbufdestroynotifyId int
var callbackPixbufdestroynotifyMap = make(map[int]PixbufdestroynotifyCallback)
var callbackPixbufdestroynotifyLock sync.RWMutex

// PixbufdestroynotifyCallback is a callback function for a 'PixbufDestroyNotify' callback.
type PixbufdestroynotifyCallback func(pixels []uint8)
