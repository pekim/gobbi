// This is a generated file - DO NOT EDIT
// +build pango_1.4 pango_1.6 pango_1.8 pango_1.10 pango_1.12 pango_1.14 pango_1.16 pango_1.18 pango_1.20 pango_1.22 pango_1.26 pango_1.30 pango_1.31.0 pango_1.32 pango_1.32.4 pango_1.34 pango_1.38

package pango

import "sync"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <pango/pango.h>
// #include <stdlib.h>
/*

	gboolean callback_fontsetforeachfuncHandler(GObject *, PangoFontset*, PangoFont*, gpointer, gpointer);

*/
import "C"

var callbackFontsetforeachfuncId int
var callbackFontsetforeachfuncMap = make(map[int]FontsetforeachfuncCallback)
var callbackFontsetforeachfuncLock sync.RWMutex

// FontsetforeachfuncCallback is a callback function for a 'FontsetForeachFunc' callback.
type FontsetforeachfuncCallback func(fontset *Fontset, font *Font) bool
