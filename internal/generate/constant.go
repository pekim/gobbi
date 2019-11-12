package generate

import (
	"fmt"
)

type Constant struct {
	Name string `xml:"name,attr"`
	//Blacklist bool   `xml:"blacklist,attr"`
	Value   string `xml:"value,attr"`
	Version string `xml:"version,attr"`
	CType   string `xml:"http://www.gtk.org/introspection/c/1.0 type,attr"`
	//Doc       *Doc   `xml:"doc"`
	Type *Type `xml:"type"`

	Namespace *Namespace
}

func (c *Constant) init(ns *Namespace) {
	c.Namespace = ns
}

func (c Constant) generate(f *file) {
	if c.Type == nil {
		fmt.Println(c)
		fmt.Println("  ", c.CType)
		fmt.Println("  ", c.Type)
		return
	}

	goName := makeExportedGoName(c.Name)

	goType := c.Type.jenGoType()
	if goType == nil {
		f.unsupported(c.Name, "")
		return
	}

	f.
		Const().
		Id(goName).
		Add(goType).
		Op("=").
		Lit(c.Value)
}

//// #include <gobject-introspection-1.0/girepository.h>
//import "C"
//
//import (
//	"github.com/dave/jennifer/jen"
//	"unsafe"
//)
//
//func (r *repository) generateConstant(info *C.GIBaseInfo) {
//	cName := C.g_base_info_get_name(info)
//	name := C.GoString(cName)
//
//	constantInfo := (*C.GIConstantInfo)(info)
//
//	typ := C.g_constant_info_get_type(constantInfo)
//	typeTag := C.g_type_info_get_tag(typ)
//
//	var value C.GIArgument
//	ptrValue := unsafe.Pointer(&value)
//	C.g_constant_info_get_value(constantInfo, &value)
//
//	var jenValue *jen.Statement
//	switch typeTag {
//	case C.GI_TYPE_TAG_BOOLEAN:
//		jenValue = jen.Lit(*((*C.gboolean)(ptrValue)) == C.TRUE)
//	case C.GI_TYPE_TAG_INT8:
//		jenValue = jen.Lit(int8(*((*C.gint8)(ptrValue))))
//	case C.GI_TYPE_TAG_UINT8:
//		jenValue = jen.Lit(uint8(*((*C.guint8)(ptrValue))))
//	case C.GI_TYPE_TAG_INT16:
//		jenValue = jen.Lit(int16(*((*C.gint16)(ptrValue))))
//	case C.GI_TYPE_TAG_UINT16:
//		jenValue = jen.Lit(uint16(*((*C.guint16)(ptrValue))))
//	case C.GI_TYPE_TAG_INT32:
//		jenValue = jen.Lit(int32(*((*C.gint32)(ptrValue))))
//	case C.GI_TYPE_TAG_UINT32:
//		jenValue = jen.Lit(uint32(*((*C.guint32)(ptrValue))))
//	case C.GI_TYPE_TAG_INT64:
//		jenValue = jen.Lit(int64(*((*C.gint64)(ptrValue))))
//	case C.GI_TYPE_TAG_UINT64:
//		jenValue = jen.Lit(uint64(*((*C.guint64)(ptrValue))))
//	case C.GI_TYPE_TAG_FLOAT:
//		jenValue = jen.Lit(float32(*((*C.gfloat)(ptrValue))))
//	case C.GI_TYPE_TAG_DOUBLE:
//		jenValue = jen.Lit(float64(*((*C.gdouble)(ptrValue))))
//	case C.GI_TYPE_TAG_UTF8:
//		jenValue = jen.Lit(C.GoString(*((**C.char)(ptrValue))))
//	default:
//		r.file.Commentf("Unsupported constant type %d for %s", typeTag, name)
//		return
//	}
//
//	r.file.
//		Const().
//		Id(name).
//		Op("=").
//		Add(jenValue)
//
//	C.g_base_info_unref(typ)
//	C.g_constant_info_free_value(constantInfo, &value)
//}
