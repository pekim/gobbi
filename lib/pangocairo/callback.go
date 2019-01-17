// This is a generated file - DO NOT EDIT

package pangocairo

import (
	cairo "github.com/pekim/gobbi/lib/cairo"
	pango "github.com/pekim/gobbi/lib/pango"
	"sync"
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
