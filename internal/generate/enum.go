package generate

// #include <gobject-introspection-1.0/girepository.h>
import "C"

import (
	"fmt"
	"github.com/dave/jennifer/jen"
)

func (r *repository) generateEnum(info *C.GIBaseInfo) {
	cName := C.g_base_info_get_name(info)
	name := C.GoString(cName)

	enum := (*C.GIConstantInfo)(info)

	typeTag := C.g_enum_info_get_storage_type(enum)

	var jenType *jen.Statement
	switch typeTag {
	case C.GI_TYPE_TAG_INT8:
		jenType = jen.Int8()
	case C.GI_TYPE_TAG_UINT8:
		jenType = jen.Uint8()
	case C.GI_TYPE_TAG_INT16:
		jenType = jen.Int16()
	case C.GI_TYPE_TAG_UINT16:
		jenType = jen.Uint16()
	case C.GI_TYPE_TAG_INT32:
		jenType = jen.Int32()
	case C.GI_TYPE_TAG_UINT32:
		jenType = jen.Uint32()
	case C.GI_TYPE_TAG_INT64:
		jenType = jen.Int64()
	case C.GI_TYPE_TAG_UINT64:
		jenType = jen.Uint64()
	default:
		// unsupported
	}

	if jenType == nil {
		r.file.Commentf("Unsupported enum type %d for %s", typeTag, name)
		return
	}

	r.file.
		Type().
		Id(name).
		Add(jenType)
}
