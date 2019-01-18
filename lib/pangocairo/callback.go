// This is a generated file - DO NOT EDIT

package pangocairo

import (
	cairo "github.com/pekim/gobbi/lib/cairo"
	pango "github.com/pekim/gobbi/lib/pango"
	"sync"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <pango/pangocairo.h>
// #include <stdlib.h>
/*

	void callback_shaperendererfuncHandler(GObject *, cairo_t*, PangoAttrShape*, gboolean, gpointer, gpointer);

*/
import "C"

var callbackShaperendererfuncId int
var callbackShaperendererfuncMap = make(map[int]ShaperendererfuncCallback)
var callbackShaperendererfuncLock sync.RWMutex

// ShaperendererfuncCallback is a callback function for a 'ShapeRendererFunc' callback.
type ShaperendererfuncCallback func(cr *cairo.Context, attr *pango.AttrShape, doPath bool)

//export callback_shaperendererfuncHandler
func callback_shaperendererfuncHandler(_ *C.GObject, c_cr *C.cairo_t, c_attr *C.PangoAttrShape, c_do_path C.gboolean, data C.gpointer) {
	callbackShaperendererfuncLock.RLock()
	defer callbackShaperendererfuncLock.RUnlock()

	cr := cairo.ContextNewFromC(unsafe.Pointer(c_cr))

	attr := pango.AttrShapeNewFromC(unsafe.Pointer(c_attr))

	doPath := c_do_path == C.TRUE

	index := int(uintptr(data))
	callback := callbackShaperendererfuncMap[index].callback
	callback(cr, attr, doPath, data)
}
