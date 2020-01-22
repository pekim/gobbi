// Code generated - DO NOT EDIT.
// +build atk_2.12

package atk

import "unsafe"

// #include <atk/atk.h>
// #include <stdlib.h>
import "C"

func toCBool(b bool) C.gboolean {
	if b {
		return C.TRUE
	}
	return C.FALSE
}
func toGoBool(b C.gboolean) bool {
	return b == C.TRUE
}

type TableCellIface C.AtkTableCellIface

func Fn_atk_range_new(param0 float64, param1 float64, param2 string) unsafe.Pointer {
	cValue0 := (C.gdouble)(param0)

	cValue1 := (C.gdouble)(param1)

	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	ret := C.atk_range_new(cValue0, cValue1, cValue2)

	return unsafe.Pointer(ret)
}

func Fn_atk_range_copy(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.AtkRange)(unsafe.Pointer(paramInstance))

	ret := C.atk_range_copy(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_atk_range_free(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.AtkRange)(unsafe.Pointer(paramInstance))

	C.atk_range_free(cValueInstance)
}

func Fn_atk_range_get_description(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.AtkRange)(unsafe.Pointer(paramInstance))

	ret := C.atk_range_get_description(cValueInstance)

	return C.GoString(ret)
}

func Fn_atk_range_get_lower_limit(paramInstance unsafe.Pointer) float64 {
	cValueInstance := (*C.AtkRange)(unsafe.Pointer(paramInstance))

	ret := C.atk_range_get_lower_limit(cValueInstance)

	return (float64)(ret)
}

func Fn_atk_range_get_upper_limit(paramInstance unsafe.Pointer) float64 {
	cValueInstance := (*C.AtkRange)(unsafe.Pointer(paramInstance))

	ret := C.atk_range_get_upper_limit(cValueInstance)

	return (float64)(ret)
}

// UNSUPPORTED : atk_add_focus_tracker : parameter 'focus_tracker' is callback

// UNSUPPORTED : atk_add_global_event_listener : parameter 'listener' is callback

// UNSUPPORTED : atk_add_key_event_listener : parameter 'listener' is callback

// UNSUPPORTED : atk_focus_tracker_init : parameter 'init' is callback

// UNSUPPORTED : atk_text_free_ranges : parameter 'ranges' is array parameter without length parameter

// UNSUPPORTED : atk_object_connect_property_change_handler : parameter 'handler' is callback

// UNSUPPORTED : atk_relation_get_target : no array length

// UNSUPPORTED : atk_component_add_focus_handler : parameter 'handler' is callback

func Fn_atk_document_get_current_page_number(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.AtkDocument)(unsafe.Pointer(paramInstance))

	ret := C.atk_document_get_current_page_number(cValueInstance)

	return (int)(ret)
}

func Fn_atk_document_get_page_count(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.AtkDocument)(unsafe.Pointer(paramInstance))

	ret := C.atk_document_get_page_count(cValueInstance)

	return (int)(ret)
}

// UNSUPPORTED : atk_table_cell_get_column_header_cells : no array length

func Fn_atk_table_cell_get_column_span(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.AtkTableCell)(unsafe.Pointer(paramInstance))

	ret := C.atk_table_cell_get_column_span(cValueInstance)

	return (int)(ret)
}

func Fn_atk_table_cell_get_position(paramInstance unsafe.Pointer, param0 *int, param1 *int) bool {
	cValueInstance := (*C.AtkTableCell)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gint)(unsafe.Pointer(param0))

	cValue1 := (*C.gint)(unsafe.Pointer(param1))

	ret := C.atk_table_cell_get_position(cValueInstance, cValue0, cValue1)

	return toGoBool(ret)
}

func Fn_atk_table_cell_get_row_column_span(paramInstance unsafe.Pointer, param0 *int, param1 *int, param2 *int, param3 *int) bool {
	cValueInstance := (*C.AtkTableCell)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gint)(unsafe.Pointer(param0))

	cValue1 := (*C.gint)(unsafe.Pointer(param1))

	cValue2 := (*C.gint)(unsafe.Pointer(param2))

	cValue3 := (*C.gint)(unsafe.Pointer(param3))

	ret := C.atk_table_cell_get_row_column_span(cValueInstance, cValue0, cValue1, cValue2, cValue3)

	return toGoBool(ret)
}

// UNSUPPORTED : atk_table_cell_get_row_header_cells : no array length

func Fn_atk_table_cell_get_row_span(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.AtkTableCell)(unsafe.Pointer(paramInstance))

	ret := C.atk_table_cell_get_row_span(cValueInstance)

	return (int)(ret)
}

func Fn_atk_table_cell_get_table(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.AtkTableCell)(unsafe.Pointer(paramInstance))

	ret := C.atk_table_cell_get_table(cValueInstance)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : atk_text_get_bounded_ranges : no array length

// UNSUPPORTED : atk_text_get_default_attributes : blacklisted

// UNSUPPORTED : atk_text_get_run_attributes : blacklisted

func Fn_atk_value_get_increment(paramInstance unsafe.Pointer) float64 {
	cValueInstance := (*C.AtkValue)(unsafe.Pointer(paramInstance))

	ret := C.atk_value_get_increment(cValueInstance)

	return (float64)(ret)
}

func Fn_atk_value_get_range(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.AtkValue)(unsafe.Pointer(paramInstance))

	ret := C.atk_value_get_range(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_atk_value_get_sub_ranges(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.AtkValue)(unsafe.Pointer(paramInstance))

	ret := C.atk_value_get_sub_ranges(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_atk_value_get_value_and_text(paramInstance unsafe.Pointer, param0 *float64, param1 *string) {
	cValueInstance := (*C.AtkValue)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gdouble)(unsafe.Pointer(param0))

	var cValue1String *C.gchar
	cValue1 := &cValue1String

	C.atk_value_get_value_and_text(cValueInstance, cValue0, cValue1)

	param1String := C.GoString(cValue1String)
	*param1 = param1String
}

func Fn_atk_value_set_value(paramInstance unsafe.Pointer, param0 float64) {
	cValueInstance := (*C.AtkValue)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gdouble)(param0)

	C.atk_value_set_value(cValueInstance, cValue0)
}
