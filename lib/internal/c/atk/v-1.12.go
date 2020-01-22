// Code generated - DO NOT EDIT.
// +build atk_1.12 atk_1.13 atk_1.20 atk_1.30 atk_2.7.4 atk_2.8 atk_2.10 atk_2.12 atk_2.14 atk_2.30 atk_2.32 atk_2.34

package atk

import (
	glib "github.com/pekim/gobbi/lib/internal/c/glib"
	"unsafe"
)

// #include <atk/atk.h>
// #include <stdlib.h>
import "C"

// UNSUPPORTED : atk_add_focus_tracker : parameter 'focus_tracker' is callback

// UNSUPPORTED : atk_add_global_event_listener : parameter 'listener' is callback

// UNSUPPORTED : atk_add_key_event_listener : parameter 'listener' is callback

// UNSUPPORTED : atk_focus_tracker_init : parameter 'init' is callback

// UNSUPPORTED : atk_text_free_ranges : parameter 'ranges' is array parameter without length parameter

// UNSUPPORTED : atk_object_connect_property_change_handler : parameter 'handler' is callback

func Fn_atk_object_get_attributes(paramInstance unsafe.Pointer) *glib.SList {
	cValueInstance := (*C.AtkObject)(unsafe.Pointer(paramInstance))

	ret := C.atk_object_get_attributes(cValueInstance)

	return (*glib.SList)(ret)
}

// UNSUPPORTED : atk_relation_get_target : no array length

// UNSUPPORTED : atk_component_add_focus_handler : parameter 'handler' is callback

func Fn_atk_component_get_alpha(paramInstance unsafe.Pointer) float64 {
	cValueInstance := (*C.AtkComponent)(unsafe.Pointer(paramInstance))

	ret := C.atk_component_get_alpha(cValueInstance)

	return (float64)(ret)
}

func Fn_atk_document_get_attribute_value(paramInstance unsafe.Pointer, param0 string) string {
	cValueInstance := (*C.AtkDocument)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.atk_document_get_attribute_value(cValueInstance, cValue0)

	return C.GoString(ret)
}

func Fn_atk_document_get_attributes(paramInstance unsafe.Pointer) *glib.SList {
	cValueInstance := (*C.AtkDocument)(unsafe.Pointer(paramInstance))

	ret := C.atk_document_get_attributes(cValueInstance)

	return (*glib.SList)(ret)
}

func Fn_atk_document_set_attribute_value(paramInstance unsafe.Pointer, param0 string, param1 string) bool {
	cValueInstance := (*C.AtkDocument)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	ret := C.atk_document_set_attribute_value(cValueInstance, cValue0, cValue1)

	return toGoBool(ret)
}

func Fn_atk_hyperlink_impl_get_hyperlink(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.AtkHyperlinkImpl)(unsafe.Pointer(paramInstance))

	ret := C.atk_hyperlink_impl_get_hyperlink(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_atk_image_get_image_locale(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.AtkImage)(unsafe.Pointer(paramInstance))

	ret := C.atk_image_get_image_locale(cValueInstance)

	return C.GoString(ret)
}

func Fn_atk_streamable_content_get_uri(paramInstance unsafe.Pointer, param0 string) string {
	cValueInstance := (*C.AtkStreamableContent)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.atk_streamable_content_get_uri(cValueInstance, cValue0)

	return C.GoString(ret)
}

// UNSUPPORTED : atk_table_cell_get_column_header_cells : no array length

// UNSUPPORTED : atk_table_cell_get_row_header_cells : no array length

// UNSUPPORTED : atk_text_get_bounded_ranges : no array length

// UNSUPPORTED : atk_text_get_default_attributes : blacklisted

// UNSUPPORTED : atk_text_get_run_attributes : blacklisted

func Fn_atk_value_get_minimum_increment(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.AtkValue)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GValue)(unsafe.Pointer(param0))

	C.atk_value_get_minimum_increment(cValueInstance, cValue0)
}
