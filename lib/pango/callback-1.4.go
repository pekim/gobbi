// This is a generated file - DO NOT EDIT
// +build pango_1.4 pango_1.6 pango_1.8 pango_1.10 pango_1.12 pango_1.14 pango_1.16 pango_1.18 pango_1.20 pango_1.22 pango_1.26 pango_1.30 pango_1.31.0 pango_1.32 pango_1.32.4 pango_1.34 pango_1.38

package pango

import (
	"sync"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <pango/pango.h>
// #include <stdlib.h>
/*

	gboolean callback_fontsetforeachfuncHandler(GObject *, PangoFontset*, PangoFont*, gpointer);

*/
import "C"

var callbackFontsetforeachfuncId int
var callbackFontsetforeachfuncMap = make(map[int]FontsetforeachfuncCallback)
var callbackFontsetforeachfuncLock sync.RWMutex

// FontsetforeachfuncCallback is a callback function for a 'FontsetForeachFunc' callback.
type FontsetforeachfuncCallback func(fontset *Fontset, font *Font) bool

//export callback_fontsetforeachfuncHandler
func callback_fontsetforeachfuncHandler(_ *C.GObject, c_fontset *C.PangoFontset, c_font *C.PangoFont, c_user_data C.gpointer) C.gboolean {
	callbackFontsetforeachfuncLock.RLock()
	defer callbackFontsetforeachfuncLock.RUnlock()

	fontset := FontsetNewFromC(unsafe.Pointer(c_fontset))

	font := FontNewFromC(unsafe.Pointer(c_font))

	index := int(uintptr(c_user_data))
	callback := callbackFontsetforeachfuncMap[index]
	retGo := callback(fontset, font)
	retC :=
		boolToGboolean(retGo)
	return retC
}
