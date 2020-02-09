package gi

// #include <girepository.h>
// #include <stdlib.h>
import "C"

import (
	"fmt"
	"unsafe"
)

// setValue takes a C.GIArgument with a value returned from a gi
// function invocation, and sets its value in the value field.
func (a *Arg) setValue(value C.GIArgument) {
	valuePtr := unsafe.Pointer(&value)

	switch a.typ {
	case ArgType_boolean:
		a.value = (*(*C.gboolean)(valuePtr)) == C.TRUE
	case ArgType_int8:
		a.value = (int8)(*(*C.gint8)(valuePtr))
	case ArgType_uint8:
		a.value = (uint8)(*(*C.guint8)(valuePtr))
	case ArgType_int16:
		a.value = (int16)(*(*C.gint16)(valuePtr))
	case ArgType_uint16:
		a.value = (uint16)(*(*C.guint16)(valuePtr))
	case ArgType_int32:
		a.value = (int32)(*(*C.gint32)(valuePtr))
	case ArgType_uint32:
		a.value = (uint32)(*(*C.guint32)(valuePtr))
	case ArgType_int64:
		a.value = (int64)(*(*C.gint64)(valuePtr))
	case ArgType_uint64:
		a.value = (uint64)(*(*C.guint64)(valuePtr))
	case ArgType_float:
		a.value = (float32)(*(*C.gfloat)(valuePtr))
	case ArgType_double:
		a.value = (float64)(*(*C.gdouble)(valuePtr))
	case ArgType_short:
		a.value = (int16)(*(*C.gshort)(valuePtr))
	case ArgType_ushort:
		a.value = (uint16)(*(*C.gushort)(valuePtr))
	case ArgType_int:
		a.value = (int32)(*(*C.gint)(valuePtr))
	case ArgType_uint:
		a.value = (uint32)(*(*C.guint)(valuePtr))
	case ArgType_long:
		a.value = (int64)(*(*C.glong)(valuePtr))
	case ArgType_ulong:
		a.value = (uint64)(*(*C.gulong)(valuePtr))
	case ArgType_ssize:
		a.value = (int)(*(*C.gssize)(valuePtr))
	case ArgType_size:
		a.value = (uint)(*(*C.gsize)(valuePtr))
	case ArgType_string:
		var cString *C.gchar
		if a.out {
			cString = *(**C.gchar)(unsafe.Pointer(&a.outPtr))
		} else {
			cString = *(**C.gchar)(valuePtr)
		}

		a.value = C.GoString(cString)

		if a.transferOwnership != TransferOwnershipNone {
			C.free(unsafe.Pointer(cString))
		}
	case ArgType_pointer:
		a.value = unsafe.Pointer(*(*C.gpointer)(valuePtr))
	default:
		panic(fmt.Sprintf("Unhandle arg type, %#v", a))
	}
}
