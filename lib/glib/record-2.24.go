// This is a generated file - DO NOT EDIT
// +build glib_2.24 glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

import (
	"fmt"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
/*

	static GVariant* _g_variant_new(const gchar* format_string) {
		return g_variant_new(format_string);
    }
*/
/*

	static GVariant* _g_variant_new_parsed(const gchar* format) {
		return g_variant_new_parsed(format);
    }
*/
/*

	static void _g_variant_get(GVariant* value, const gchar* format_string) {
		return g_variant_get(value, format_string);
    }
*/
/*

	static void _g_variant_get_child(GVariant* value, gsize index_, const gchar* format_string) {
		return g_variant_get_child(value, index_, format_string);
    }
*/
/*

	static void _g_variant_builder_add(GVariantBuilder* builder, const gchar* format_string) {
		return g_variant_builder_add(builder, format_string);
    }
*/
/*

	static gboolean _g_variant_iter_loop(GVariantIter* iter, const gchar* format_string) {
		return g_variant_iter_loop(iter, format_string);
    }
*/
/*

	static gboolean _g_variant_iter_next(GVariantIter* iter, const gchar* format_string) {
		return g_variant_iter_next(iter, format_string);
    }
*/
import "C"

// Variant is a wrapper around the C record GVariant.
type Variant struct {
	native *C.GVariant
}

func VariantNewFromC(u unsafe.Pointer) *Variant {
	c := (*C.GVariant)(u)
	if c == nil {
		return nil
	}

	g := &Variant{native: c}

	return g
}

func (recv *Variant) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Variant with another Variant, and returns true if they represent the same GObject.
func (recv *Variant) Equals(other *Variant) bool {
	return other.ToC() == recv.ToC()
}

// VariantNew is a wrapper around the C function g_variant_new.
func VariantNew(formatString string, args ...interface{}) *Variant {
	goFormattedString := fmt.Sprintf(formatString, args...)
	c_format_string := C.CString(goFormattedString)
	defer C.free(unsafe.Pointer(c_format_string))

	retC := C._g_variant_new(c_format_string)
	retGo := VariantNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_variant_new_array : unsupported parameter children :

// VariantNewBoolean is a wrapper around the C function g_variant_new_boolean.
func VariantNewBoolean(value bool) *Variant {
	c_value :=
		boolToGboolean(value)

	retC := C.g_variant_new_boolean(c_value)
	retGo := VariantNewFromC(unsafe.Pointer(retC))

	return retGo
}

// VariantNewByte is a wrapper around the C function g_variant_new_byte.
func VariantNewByte(value uint8) *Variant {
	c_value := (C.guchar)(value)

	retC := C.g_variant_new_byte(c_value)
	retGo := VariantNewFromC(unsafe.Pointer(retC))

	return retGo
}

// VariantNewDictEntry is a wrapper around the C function g_variant_new_dict_entry.
func VariantNewDictEntry(key *Variant, value *Variant) *Variant {
	c_key := (*C.GVariant)(C.NULL)
	if key != nil {
		c_key = (*C.GVariant)(key.ToC())
	}

	c_value := (*C.GVariant)(C.NULL)
	if value != nil {
		c_value = (*C.GVariant)(value.ToC())
	}

	retC := C.g_variant_new_dict_entry(c_key, c_value)
	retGo := VariantNewFromC(unsafe.Pointer(retC))

	return retGo
}

// VariantNewDouble is a wrapper around the C function g_variant_new_double.
func VariantNewDouble(value float64) *Variant {
	c_value := (C.gdouble)(value)

	retC := C.g_variant_new_double(c_value)
	retGo := VariantNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_variant_new_from_data : unsupported parameter notify : no type generator for DestroyNotify (GDestroyNotify) for param notify

// VariantNewHandle is a wrapper around the C function g_variant_new_handle.
func VariantNewHandle(value int32) *Variant {
	c_value := (C.gint32)(value)

	retC := C.g_variant_new_handle(c_value)
	retGo := VariantNewFromC(unsafe.Pointer(retC))

	return retGo
}

// VariantNewInt16 is a wrapper around the C function g_variant_new_int16.
func VariantNewInt16(value int16) *Variant {
	c_value := (C.gint16)(value)

	retC := C.g_variant_new_int16(c_value)
	retGo := VariantNewFromC(unsafe.Pointer(retC))

	return retGo
}

// VariantNewInt32 is a wrapper around the C function g_variant_new_int32.
func VariantNewInt32(value int32) *Variant {
	c_value := (C.gint32)(value)

	retC := C.g_variant_new_int32(c_value)
	retGo := VariantNewFromC(unsafe.Pointer(retC))

	return retGo
}

// VariantNewInt64 is a wrapper around the C function g_variant_new_int64.
func VariantNewInt64(value int64) *Variant {
	c_value := (C.gint64)(value)

	retC := C.g_variant_new_int64(c_value)
	retGo := VariantNewFromC(unsafe.Pointer(retC))

	return retGo
}

// VariantNewMaybe is a wrapper around the C function g_variant_new_maybe.
func VariantNewMaybe(childType *VariantType, child *Variant) *Variant {
	c_child_type := (*C.GVariantType)(C.NULL)
	if childType != nil {
		c_child_type = (*C.GVariantType)(childType.ToC())
	}

	c_child := (*C.GVariant)(C.NULL)
	if child != nil {
		c_child = (*C.GVariant)(child.ToC())
	}

	retC := C.g_variant_new_maybe(c_child_type, c_child)
	retGo := VariantNewFromC(unsafe.Pointer(retC))

	return retGo
}

// VariantNewObjectPath is a wrapper around the C function g_variant_new_object_path.
func VariantNewObjectPath(objectPath string) *Variant {
	c_object_path := C.CString(objectPath)
	defer C.free(unsafe.Pointer(c_object_path))

	retC := C.g_variant_new_object_path(c_object_path)
	retGo := VariantNewFromC(unsafe.Pointer(retC))

	return retGo
}

// VariantNewParsed is a wrapper around the C function g_variant_new_parsed.
func VariantNewParsed(format string, args ...interface{}) *Variant {
	goFormattedString := fmt.Sprintf(format, args...)
	c_format := C.CString(goFormattedString)
	defer C.free(unsafe.Pointer(c_format))

	retC := C._g_variant_new_parsed(c_format)
	retGo := VariantNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_variant_new_parsed_va : unsupported parameter app : no type generator for va_list (va_list*) for param app

// VariantNewSignature is a wrapper around the C function g_variant_new_signature.
func VariantNewSignature(signature string) *Variant {
	c_signature := C.CString(signature)
	defer C.free(unsafe.Pointer(c_signature))

	retC := C.g_variant_new_signature(c_signature)
	retGo := VariantNewFromC(unsafe.Pointer(retC))

	return retGo
}

// VariantNewString is a wrapper around the C function g_variant_new_string.
func VariantNewString(string_ string) *Variant {
	c_string := C.CString(string_)
	defer C.free(unsafe.Pointer(c_string))

	retC := C.g_variant_new_string(c_string)
	retGo := VariantNewFromC(unsafe.Pointer(retC))

	return retGo
}

// VariantNewStrv is a wrapper around the C function g_variant_new_strv.
func VariantNewStrv(strv []string) *Variant {
	c_strv_array := make([]*C.gchar, len(strv), len(strv))
	for i, item := range strv {
		c := C.CString(item)
		defer C.free(unsafe.Pointer(c))
		c_strv_array[i] = c
	}
	c_strv_arrayPtr := &c_strv_array[0]
	c_strv := (**C.gchar)(unsafe.Pointer(c_strv_arrayPtr))

	c_length := (C.gssize)(len(strv))

	retC := C.g_variant_new_strv(c_strv, c_length)
	retGo := VariantNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_variant_new_tuple : unsupported parameter children :

// VariantNewUint16 is a wrapper around the C function g_variant_new_uint16.
func VariantNewUint16(value uint16) *Variant {
	c_value := (C.guint16)(value)

	retC := C.g_variant_new_uint16(c_value)
	retGo := VariantNewFromC(unsafe.Pointer(retC))

	return retGo
}

// VariantNewUint32 is a wrapper around the C function g_variant_new_uint32.
func VariantNewUint32(value uint32) *Variant {
	c_value := (C.guint32)(value)

	retC := C.g_variant_new_uint32(c_value)
	retGo := VariantNewFromC(unsafe.Pointer(retC))

	return retGo
}

// VariantNewUint64 is a wrapper around the C function g_variant_new_uint64.
func VariantNewUint64(value uint64) *Variant {
	c_value := (C.guint64)(value)

	retC := C.g_variant_new_uint64(c_value)
	retGo := VariantNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_variant_new_va : unsupported parameter endptr : in string with indirection level of 2

// VariantNewVariant is a wrapper around the C function g_variant_new_variant.
func VariantNewVariant(value *Variant) *Variant {
	c_value := (*C.GVariant)(C.NULL)
	if value != nil {
		c_value = (*C.GVariant)(value.ToC())
	}

	retC := C.g_variant_new_variant(c_value)
	retGo := VariantNewFromC(unsafe.Pointer(retC))

	return retGo
}

// VariantIsObjectPath is a wrapper around the C function g_variant_is_object_path.
func VariantIsObjectPath(string_ string) bool {
	c_string := C.CString(string_)
	defer C.free(unsafe.Pointer(c_string))

	retC := C.g_variant_is_object_path(c_string)
	retGo := retC == C.TRUE

	return retGo
}

// VariantIsSignature is a wrapper around the C function g_variant_is_signature.
func VariantIsSignature(string_ string) bool {
	c_string := C.CString(string_)
	defer C.free(unsafe.Pointer(c_string))

	retC := C.g_variant_is_signature(c_string)
	retGo := retC == C.TRUE

	return retGo
}

// Byteswap is a wrapper around the C function g_variant_byteswap.
func (recv *Variant) Byteswap() *Variant {
	retC := C.g_variant_byteswap((*C.GVariant)(recv.native))
	retGo := VariantNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Classify is a wrapper around the C function g_variant_classify.
func (recv *Variant) Classify() VariantClass {
	retC := C.g_variant_classify((*C.GVariant)(recv.native))
	retGo := (VariantClass)(retC)

	return retGo
}

// DupString is a wrapper around the C function g_variant_dup_string.
func (recv *Variant) DupString() (string, uint64) {
	var c_length C.gsize

	retC := C.g_variant_dup_string((*C.GVariant)(recv.native), &c_length)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	length := (uint64)(c_length)

	return retGo, length
}

// DupStrv is a wrapper around the C function g_variant_dup_strv.
func (recv *Variant) DupStrv() ([]string, uint64) {
	var c_length C.gsize

	retC := C.g_variant_dup_strv((*C.GVariant)(recv.native), &c_length)
	retGo := []string{}
	for p := retC; *p != nil; p = (**C.char)(C.gpointer((uintptr(C.gpointer(p)) + uintptr(C.sizeof_gpointer)))) {
		s := C.GoString(*p)
		retGo = append(retGo, s)
	}
	defer C.g_strfreev(retC)

	length := (uint64)(c_length)

	return retGo, length
}

// Equal is a wrapper around the C function g_variant_equal.
func (recv *Variant) Equal(two *Variant) bool {
	c_two := (C.gconstpointer)(C.NULL)
	if two != nil {
		c_two = (C.gconstpointer)(two.ToC())
	}

	retC := C.g_variant_equal((C.gconstpointer)(recv.native), c_two)
	retGo := retC == C.TRUE

	return retGo
}

// Get is a wrapper around the C function g_variant_get.
func (recv *Variant) Get(formatString string, args ...interface{}) {
	goFormattedString := fmt.Sprintf(formatString, args...)
	c_format_string := C.CString(goFormattedString)
	defer C.free(unsafe.Pointer(c_format_string))

	C._g_variant_get((*C.GVariant)(recv.native), c_format_string)

	return
}

// GetBoolean is a wrapper around the C function g_variant_get_boolean.
func (recv *Variant) GetBoolean() bool {
	retC := C.g_variant_get_boolean((*C.GVariant)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetByte is a wrapper around the C function g_variant_get_byte.
func (recv *Variant) GetByte() uint8 {
	retC := C.g_variant_get_byte((*C.GVariant)(recv.native))
	retGo := (uint8)(retC)

	return retGo
}

// GetChild is a wrapper around the C function g_variant_get_child.
func (recv *Variant) GetChild(index uint64, formatString string, args ...interface{}) {
	c_index_ := (C.gsize)(index)

	goFormattedString := fmt.Sprintf(formatString, args...)
	c_format_string := C.CString(goFormattedString)
	defer C.free(unsafe.Pointer(c_format_string))

	C._g_variant_get_child((*C.GVariant)(recv.native), c_index_, c_format_string)

	return
}

// GetChildValue is a wrapper around the C function g_variant_get_child_value.
func (recv *Variant) GetChildValue(index uint64) *Variant {
	c_index_ := (C.gsize)(index)

	retC := C.g_variant_get_child_value((*C.GVariant)(recv.native), c_index_)
	retGo := VariantNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_variant_get_data : no return generator

// GetDouble is a wrapper around the C function g_variant_get_double.
func (recv *Variant) GetDouble() float64 {
	retC := C.g_variant_get_double((*C.GVariant)(recv.native))
	retGo := (float64)(retC)

	return retGo
}

// Unsupported : g_variant_get_fixed_array : no type generator for gpointer (gconstpointer) for array return

// GetHandle is a wrapper around the C function g_variant_get_handle.
func (recv *Variant) GetHandle() int32 {
	retC := C.g_variant_get_handle((*C.GVariant)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetInt16 is a wrapper around the C function g_variant_get_int16.
func (recv *Variant) GetInt16() int16 {
	retC := C.g_variant_get_int16((*C.GVariant)(recv.native))
	retGo := (int16)(retC)

	return retGo
}

// GetInt32 is a wrapper around the C function g_variant_get_int32.
func (recv *Variant) GetInt32() int32 {
	retC := C.g_variant_get_int32((*C.GVariant)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetInt64 is a wrapper around the C function g_variant_get_int64.
func (recv *Variant) GetInt64() int64 {
	retC := C.g_variant_get_int64((*C.GVariant)(recv.native))
	retGo := (int64)(retC)

	return retGo
}

// GetMaybe is a wrapper around the C function g_variant_get_maybe.
func (recv *Variant) GetMaybe() *Variant {
	retC := C.g_variant_get_maybe((*C.GVariant)(recv.native))
	var retGo (*Variant)
	if retC == nil {
		retGo = nil
	} else {
		retGo = VariantNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetNormalForm is a wrapper around the C function g_variant_get_normal_form.
func (recv *Variant) GetNormalForm() *Variant {
	retC := C.g_variant_get_normal_form((*C.GVariant)(recv.native))
	retGo := VariantNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetSize is a wrapper around the C function g_variant_get_size.
func (recv *Variant) GetSize() uint64 {
	retC := C.g_variant_get_size((*C.GVariant)(recv.native))
	retGo := (uint64)(retC)

	return retGo
}

// GetString is a wrapper around the C function g_variant_get_string.
func (recv *Variant) GetString() (string, uint64) {
	var c_length C.gsize

	retC := C.g_variant_get_string((*C.GVariant)(recv.native), &c_length)
	retGo := C.GoString(retC)

	length := (uint64)(c_length)

	return retGo, length
}

// GetStrv is a wrapper around the C function g_variant_get_strv.
func (recv *Variant) GetStrv() ([]string, uint64) {
	var c_length C.gsize

	retC := C.g_variant_get_strv((*C.GVariant)(recv.native), &c_length)
	retGo := []string{}
	for p := retC; *p != nil; p = (**C.char)(C.gpointer((uintptr(C.gpointer(p)) + uintptr(C.sizeof_gpointer)))) {
		s := C.GoString(*p)
		retGo = append(retGo, s)
	}

	length := (uint64)(c_length)

	return retGo, length
}

// GetType is a wrapper around the C function g_variant_get_type.
func (recv *Variant) GetType() *VariantType {
	retC := C.g_variant_get_type((*C.GVariant)(recv.native))
	retGo := VariantTypeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetTypeString is a wrapper around the C function g_variant_get_type_string.
func (recv *Variant) GetTypeString() string {
	retC := C.g_variant_get_type_string((*C.GVariant)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetUint16 is a wrapper around the C function g_variant_get_uint16.
func (recv *Variant) GetUint16() uint16 {
	retC := C.g_variant_get_uint16((*C.GVariant)(recv.native))
	retGo := (uint16)(retC)

	return retGo
}

// GetUint32 is a wrapper around the C function g_variant_get_uint32.
func (recv *Variant) GetUint32() uint32 {
	retC := C.g_variant_get_uint32((*C.GVariant)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// GetUint64 is a wrapper around the C function g_variant_get_uint64.
func (recv *Variant) GetUint64() uint64 {
	retC := C.g_variant_get_uint64((*C.GVariant)(recv.native))
	retGo := (uint64)(retC)

	return retGo
}

// Unsupported : g_variant_get_va : unsupported parameter endptr : in string with indirection level of 2

// GetVariant is a wrapper around the C function g_variant_get_variant.
func (recv *Variant) GetVariant() *Variant {
	retC := C.g_variant_get_variant((*C.GVariant)(recv.native))
	retGo := VariantNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Hash is a wrapper around the C function g_variant_hash.
func (recv *Variant) Hash() uint32 {
	retC := C.g_variant_hash((C.gconstpointer)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// IsContainer is a wrapper around the C function g_variant_is_container.
func (recv *Variant) IsContainer() bool {
	retC := C.g_variant_is_container((*C.GVariant)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// IsNormalForm is a wrapper around the C function g_variant_is_normal_form.
func (recv *Variant) IsNormalForm() bool {
	retC := C.g_variant_is_normal_form((*C.GVariant)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// IsOfType is a wrapper around the C function g_variant_is_of_type.
func (recv *Variant) IsOfType(type_ *VariantType) bool {
	c_type := (*C.GVariantType)(C.NULL)
	if type_ != nil {
		c_type = (*C.GVariantType)(type_.ToC())
	}

	retC := C.g_variant_is_of_type((*C.GVariant)(recv.native), c_type)
	retGo := retC == C.TRUE

	return retGo
}

// IterNew is a wrapper around the C function g_variant_iter_new.
func (recv *Variant) IterNew() *VariantIter {
	retC := C.g_variant_iter_new((*C.GVariant)(recv.native))
	retGo := VariantIterNewFromC(unsafe.Pointer(retC))

	return retGo
}

// NChildren is a wrapper around the C function g_variant_n_children.
func (recv *Variant) NChildren() uint64 {
	retC := C.g_variant_n_children((*C.GVariant)(recv.native))
	retGo := (uint64)(retC)

	return retGo
}

// Print is a wrapper around the C function g_variant_print.
func (recv *Variant) Print(typeAnnotate bool) string {
	c_type_annotate :=
		boolToGboolean(typeAnnotate)

	retC := C.g_variant_print((*C.GVariant)(recv.native), c_type_annotate)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// PrintString is a wrapper around the C function g_variant_print_string.
func (recv *Variant) PrintString(string_ *String, typeAnnotate bool) *String {
	c_string := (*C.GString)(C.NULL)
	if string_ != nil {
		c_string = (*C.GString)(string_.ToC())
	}

	c_type_annotate :=
		boolToGboolean(typeAnnotate)

	retC := C.g_variant_print_string((*C.GVariant)(recv.native), c_string, c_type_annotate)
	retGo := StringNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Ref is a wrapper around the C function g_variant_ref.
func (recv *Variant) Ref() *Variant {
	retC := C.g_variant_ref((*C.GVariant)(recv.native))
	retGo := VariantNewFromC(unsafe.Pointer(retC))

	return retGo
}

// RefSink is a wrapper around the C function g_variant_ref_sink.
func (recv *Variant) RefSink() *Variant {
	retC := C.g_variant_ref_sink((*C.GVariant)(recv.native))
	retGo := VariantNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_variant_store : unsupported parameter data : no type generator for gpointer (gpointer) for param data

// TakeRef is a wrapper around the C function g_variant_take_ref.
func (recv *Variant) TakeRef() *Variant {
	retC := C.g_variant_take_ref((*C.GVariant)(recv.native))
	retGo := VariantNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unref is a wrapper around the C function g_variant_unref.
func (recv *Variant) Unref() {
	C.g_variant_unref((*C.GVariant)(recv.native))

	return
}

// VariantBuilderNew is a wrapper around the C function g_variant_builder_new.
func VariantBuilderNew(type_ *VariantType) *VariantBuilder {
	c_type := (*C.GVariantType)(C.NULL)
	if type_ != nil {
		c_type = (*C.GVariantType)(type_.ToC())
	}

	retC := C.g_variant_builder_new(c_type)
	retGo := VariantBuilderNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Add is a wrapper around the C function g_variant_builder_add.
func (recv *VariantBuilder) Add(formatString string, args ...interface{}) {
	goFormattedString := fmt.Sprintf(formatString, args...)
	c_format_string := C.CString(goFormattedString)
	defer C.free(unsafe.Pointer(c_format_string))

	C._g_variant_builder_add((*C.GVariantBuilder)(recv.native), c_format_string)

	return
}

// AddValue is a wrapper around the C function g_variant_builder_add_value.
func (recv *VariantBuilder) AddValue(value *Variant) {
	c_value := (*C.GVariant)(C.NULL)
	if value != nil {
		c_value = (*C.GVariant)(value.ToC())
	}

	C.g_variant_builder_add_value((*C.GVariantBuilder)(recv.native), c_value)

	return
}

// Clear is a wrapper around the C function g_variant_builder_clear.
func (recv *VariantBuilder) Clear() {
	C.g_variant_builder_clear((*C.GVariantBuilder)(recv.native))

	return
}

// Close is a wrapper around the C function g_variant_builder_close.
func (recv *VariantBuilder) Close() {
	C.g_variant_builder_close((*C.GVariantBuilder)(recv.native))

	return
}

// End is a wrapper around the C function g_variant_builder_end.
func (recv *VariantBuilder) End() *Variant {
	retC := C.g_variant_builder_end((*C.GVariantBuilder)(recv.native))
	retGo := VariantNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Init is a wrapper around the C function g_variant_builder_init.
func (recv *VariantBuilder) Init(type_ *VariantType) {
	c_type := (*C.GVariantType)(C.NULL)
	if type_ != nil {
		c_type = (*C.GVariantType)(type_.ToC())
	}

	C.g_variant_builder_init((*C.GVariantBuilder)(recv.native), c_type)

	return
}

// Open is a wrapper around the C function g_variant_builder_open.
func (recv *VariantBuilder) Open(type_ *VariantType) {
	c_type := (*C.GVariantType)(C.NULL)
	if type_ != nil {
		c_type = (*C.GVariantType)(type_.ToC())
	}

	C.g_variant_builder_open((*C.GVariantBuilder)(recv.native), c_type)

	return
}

// Ref is a wrapper around the C function g_variant_builder_ref.
func (recv *VariantBuilder) Ref() *VariantBuilder {
	retC := C.g_variant_builder_ref((*C.GVariantBuilder)(recv.native))
	retGo := VariantBuilderNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unref is a wrapper around the C function g_variant_builder_unref.
func (recv *VariantBuilder) Unref() {
	C.g_variant_builder_unref((*C.GVariantBuilder)(recv.native))

	return
}

// Copy is a wrapper around the C function g_variant_iter_copy.
func (recv *VariantIter) Copy() *VariantIter {
	retC := C.g_variant_iter_copy((*C.GVariantIter)(recv.native))
	retGo := VariantIterNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Free is a wrapper around the C function g_variant_iter_free.
func (recv *VariantIter) Free() {
	C.g_variant_iter_free((*C.GVariantIter)(recv.native))

	return
}

// Init is a wrapper around the C function g_variant_iter_init.
func (recv *VariantIter) Init(value *Variant) uint64 {
	c_value := (*C.GVariant)(C.NULL)
	if value != nil {
		c_value = (*C.GVariant)(value.ToC())
	}

	retC := C.g_variant_iter_init((*C.GVariantIter)(recv.native), c_value)
	retGo := (uint64)(retC)

	return retGo
}

// Loop is a wrapper around the C function g_variant_iter_loop.
func (recv *VariantIter) Loop(formatString string, args ...interface{}) bool {
	goFormattedString := fmt.Sprintf(formatString, args...)
	c_format_string := C.CString(goFormattedString)
	defer C.free(unsafe.Pointer(c_format_string))

	retC := C._g_variant_iter_loop((*C.GVariantIter)(recv.native), c_format_string)
	retGo := retC == C.TRUE

	return retGo
}

// NChildren is a wrapper around the C function g_variant_iter_n_children.
func (recv *VariantIter) NChildren() uint64 {
	retC := C.g_variant_iter_n_children((*C.GVariantIter)(recv.native))
	retGo := (uint64)(retC)

	return retGo
}

// Next is a wrapper around the C function g_variant_iter_next.
func (recv *VariantIter) Next(formatString string, args ...interface{}) bool {
	goFormattedString := fmt.Sprintf(formatString, args...)
	c_format_string := C.CString(goFormattedString)
	defer C.free(unsafe.Pointer(c_format_string))

	retC := C._g_variant_iter_next((*C.GVariantIter)(recv.native), c_format_string)
	retGo := retC == C.TRUE

	return retGo
}

// NextValue is a wrapper around the C function g_variant_iter_next_value.
func (recv *VariantIter) NextValue() *Variant {
	retC := C.g_variant_iter_next_value((*C.GVariantIter)(recv.native))
	var retGo (*Variant)
	if retC == nil {
		retGo = nil
	} else {
		retGo = VariantNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// VariantTypeNew is a wrapper around the C function g_variant_type_new.
func VariantTypeNew(typeString string) *VariantType {
	c_type_string := C.CString(typeString)
	defer C.free(unsafe.Pointer(c_type_string))

	retC := C.g_variant_type_new(c_type_string)
	retGo := VariantTypeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// VariantTypeStringScan is a wrapper around the C function g_variant_type_string_scan.
func VariantTypeStringScan(string_ string, limit string) (bool, string) {
	c_string := C.CString(string_)
	defer C.free(unsafe.Pointer(c_string))

	c_limit := C.CString(limit)
	defer C.free(unsafe.Pointer(c_limit))

	var c_endptr *C.gchar

	retC := C.g_variant_type_string_scan(c_string, c_limit, &c_endptr)
	retGo := retC == C.TRUE

	endptr := C.GoString(c_endptr)
	defer C.free(unsafe.Pointer(c_endptr))

	return retGo, endptr
}
