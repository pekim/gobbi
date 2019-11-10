package generate

// #include <gobject-introspection-1.0/girepository.h>
import "C"

import (
	"github.com/dave/jennifer/jen"
	"unsafe"
)

func (r *repository) generateConstant(info *C.GIBaseInfo) {
	cName := C.g_base_info_get_name(info)
	name := C.GoString(cName)

	typ := C.g_constant_info_get_type(info)
	typeTag := C.g_type_info_get_tag(typ)

	constant := (*C.GIConstantInfo)(info)
	var value C.GIArgument
	ptrValue := unsafe.Pointer(&value)
	C.g_constant_info_get_value(constant, &value)

	var jenValue *jen.Statement
	var typeTagName string

	switch typeTag {
	case C.GI_TYPE_TAG_VOID:
		typeTagName = "void"
	case C.GI_TYPE_TAG_BOOLEAN:
		jenValue = jen.Lit(*((*C.gboolean)(ptrValue)) == C.TRUE)
	case C.GI_TYPE_TAG_INT8:
		jenValue = jen.Lit(int8(*((*C.gint8)(ptrValue))))
	case C.GI_TYPE_TAG_UINT8:
		jenValue = jen.Lit(uint8(*((*C.guint8)(ptrValue))))
	case C.GI_TYPE_TAG_INT16:
		jenValue = jen.Lit(int16(*((*C.gint16)(ptrValue))))
	case C.GI_TYPE_TAG_UINT16:
		jenValue = jen.Lit(uint16(*((*C.guint16)(ptrValue))))
	case C.GI_TYPE_TAG_INT32:
		jenValue = jen.Lit(int32(*((*C.gint32)(ptrValue))))
	case C.GI_TYPE_TAG_UINT32:
		jenValue = jen.Lit(uint32(*((*C.guint32)(ptrValue))))
	case C.GI_TYPE_TAG_INT64:
		jenValue = jen.Lit(int64(*((*C.gint64)(ptrValue))))
	case C.GI_TYPE_TAG_UINT64:
		jenValue = jen.Lit(uint64(*((*C.guint64)(ptrValue))))
	case C.GI_TYPE_TAG_FLOAT:
		jenValue = jen.Lit(float32(*((*C.gfloat)(ptrValue))))
	case C.GI_TYPE_TAG_DOUBLE:
		jenValue = jen.Lit(float64(*((*C.gdouble)(ptrValue))))
	case C.GI_TYPE_TAG_GTYPE:
		typeTagName = "a GType"
	case C.GI_TYPE_TAG_UTF8:
		jenValue = jen.Lit(C.GoString(*((**C.char)(ptrValue))))
	case C.GI_TYPE_TAG_FILENAME:
		typeTagName = "a filename, encoded in the same encoding as the native filesystem is using."
	case C.GI_TYPE_TAG_ARRAY:
		typeTagName = "an array"
	case C.GI_TYPE_TAG_INTERFACE:
		typeTagName = "an extended interface object"
	case C.GI_TYPE_TAG_GLIST:
		typeTagName = " GList"
	case C.GI_TYPE_TAG_GSLIST:
		typeTagName = " GSList"
	case C.GI_TYPE_TAG_GHASH:
		typeTagName = " GHashTable"
	case C.GI_TYPE_TAG_ERROR:
		typeTagName = " GError"
	case C.GI_TYPE_TAG_UNICHAR:
		typeTagName = "Unicode character"
	}

	if jenValue == nil {
		r.file.Commentf("Unsupported constant type '%s' (%d) for %s", typeTagName, typeTag, name)
		return
	}

	r.file.
		Const().
		Id(name).
		Op("=").
		Add(jenValue)
}
