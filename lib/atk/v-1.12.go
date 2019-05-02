// This is a generated file - DO NOT EDIT
// +build atk_1.12 atk_1.13 atk_1.20 atk_1.30 atk_2.8 atk_2.12

package atk

import (
	gobject "github.com/pekim/gobbi/lib/gobject"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <atk/atk.h>
// #include <stdlib.h>
import "C"

// GetAttributes is a wrapper around the C function atk_object_get_attributes.
func (recv *Object) GetAttributes() *AttributeSet {
	retC := C.atk_object_get_attributes((*C.AtkObject)(recv.native))
	retGo := SListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetAlpha is a wrapper around the C function atk_component_get_alpha.
func (recv *Component) GetAlpha() float64 {
	retC := C.atk_component_get_alpha((*C.AtkComponent)(recv.native))
	retGo := (float64)(retC)

	return retGo
}

// GetAttributeValue is a wrapper around the C function atk_document_get_attribute_value.
func (recv *Document) GetAttributeValue(attributeName string) string {
	c_attribute_name := C.CString(attributeName)
	defer C.free(unsafe.Pointer(c_attribute_name))

	retC := C.atk_document_get_attribute_value((*C.AtkDocument)(recv.native), c_attribute_name)
	retGo := C.GoString(retC)

	return retGo
}

// GetAttributes is a wrapper around the C function atk_document_get_attributes.
func (recv *Document) GetAttributes() *AttributeSet {
	retC := C.atk_document_get_attributes((*C.AtkDocument)(recv.native))
	retGo := SListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SetAttributeValue is a wrapper around the C function atk_document_set_attribute_value.
func (recv *Document) SetAttributeValue(attributeName string, attributeValue string) bool {
	c_attribute_name := C.CString(attributeName)
	defer C.free(unsafe.Pointer(c_attribute_name))

	c_attribute_value := C.CString(attributeValue)
	defer C.free(unsafe.Pointer(c_attribute_value))

	retC := C.atk_document_set_attribute_value((*C.AtkDocument)(recv.native), c_attribute_name, c_attribute_value)
	retGo := retC == C.TRUE

	return retGo
}

// GetHyperlink is a wrapper around the C function atk_hyperlink_impl_get_hyperlink.
func (recv *HyperlinkImpl) GetHyperlink() *Hyperlink {
	retC := C.atk_hyperlink_impl_get_hyperlink((*C.AtkHyperlinkImpl)(recv.native))
	retGo := HyperlinkNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetImageLocale is a wrapper around the C function atk_image_get_image_locale.
func (recv *Image) GetImageLocale() string {
	retC := C.atk_image_get_image_locale((*C.AtkImage)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetUri is a wrapper around the C function atk_streamable_content_get_uri.
func (recv *StreamableContent) GetUri(mimeType string) string {
	c_mime_type := C.CString(mimeType)
	defer C.free(unsafe.Pointer(c_mime_type))

	retC := C.atk_streamable_content_get_uri((*C.AtkStreamableContent)(recv.native), c_mime_type)
	retGo := C.GoString(retC)

	return retGo
}

// GetMinimumIncrement is a wrapper around the C function atk_value_get_minimum_increment.
func (recv *Value) GetMinimumIncrement() *gobject.Value {
	var c_value C.GValue

	C.atk_value_get_minimum_increment((*C.AtkValue)(recv.native), &c_value)

	value := gobject.ValueNewFromC(unsafe.Pointer(&c_value))

	return value
}
