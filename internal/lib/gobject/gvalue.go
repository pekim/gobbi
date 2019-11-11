package gobject

import (
	"unsafe"
)

// #include <glib-object.h>
import "C"

// ValueNew create a new Value initialised for the supplied Type.
func ValueNew(typ Type) *Value {
	var gvalue C.GValue
	value := ValueNewFromC(unsafe.Pointer(&gvalue))
	value.Init(typ)

	return value
}
