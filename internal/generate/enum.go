package generate

// #include <gobject-introspection-1.0/girepository.h>
import "C"

import (
	"github.com/dave/jennifer/jen"
)

func (r *repository) generateEnum(info *C.GIBaseInfo) {
	cName := C.g_base_info_get_name(info)
	typeName := C.GoString(cName)

	enumInfo := (*C.GIConstantInfo)(info)
	typeTag := C.g_enum_info_get_storage_type(enumInfo)

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
		r.file.Commentf("Unsupported enum type %d for %s", typeTag, typeName)
		return
	}

	// generate type definition
	r.file.
		Type().
		Id(typeName).
		Add(jenType)

	// generate const values
	r.file.
		Const().
		DefsFunc(func(g *jen.Group) {
			n := C.g_enum_info_get_n_values(enumInfo)
			for i := C.int(0); i < n; i++ {
				valueInfo := C.g_enum_info_get_value(enumInfo, i)

				cName := C.g_base_info_get_name(valueInfo)
				valueName := C.GoString(cName)
				goValueName := typeName + "_" + makeExportedGoName(valueName)

				value := C.g_value_info_get_value(valueInfo)

				g.
					Id(goValueName).
					Id(typeName).
					Op("=").
					Lit(int64(value))

				C.g_base_info_unref(valueInfo)
			}
		})

	r.file.Line()
}
