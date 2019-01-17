// This is a generated file - DO NOT EDIT

package pango

import "sync"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <pango/pango.h>
// #include <stdlib.h>
/*

	gboolean callback_attrfilterfuncHandler(GObject *, PangoAttribute*, gpointer, gpointer);

*/
import "C"

// Unsupported : callback AttrDataCopyFunc : no return generator

var callbackAttrfilterfuncId int
var callbackAttrfilterfuncMap = make(map[int]AttrfilterfuncCallback)
var callbackAttrfilterfuncLock sync.RWMutex

// AttrfilterfuncCallback is a callback function for a 'AttrFilterFunc' callback.
type AttrfilterfuncCallback func(attribute *Attribute) bool
