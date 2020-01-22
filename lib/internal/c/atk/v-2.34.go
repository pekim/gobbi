// Code generated - DO NOT EDIT.
// +build atk_2.34

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

// UNSUPPORTED : atk_add_focus_tracker : parameter 'focus_tracker' is callback

// UNSUPPORTED : atk_add_global_event_listener : parameter 'listener' is callback

// UNSUPPORTED : atk_add_key_event_listener : parameter 'listener' is callback

// UNSUPPORTED : atk_focus_tracker_init : parameter 'init' is callback

// UNSUPPORTED : atk_text_free_ranges : parameter 'ranges' is array parameter without length parameter

// UNSUPPORTED : atk_object_connect_property_change_handler : parameter 'handler' is callback

func Fn_atk_object_get_accessible_id(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.AtkObject)(unsafe.Pointer(paramInstance))

	ret := C.atk_object_get_accessible_id(cValueInstance)

	return C.GoString(ret)
}

func Fn_atk_object_set_accessible_id(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.AtkObject)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.atk_object_set_accessible_id(cValueInstance, cValue0)
}

// UNSUPPORTED : atk_relation_get_target : no array length

// UNSUPPORTED : atk_component_add_focus_handler : parameter 'handler' is callback

// UNSUPPORTED : atk_table_cell_get_column_header_cells : no array length

// UNSUPPORTED : atk_table_cell_get_row_header_cells : no array length

// UNSUPPORTED : atk_text_get_bounded_ranges : no array length

// UNSUPPORTED : atk_text_get_default_attributes : blacklisted

// UNSUPPORTED : atk_text_get_run_attributes : blacklisted
