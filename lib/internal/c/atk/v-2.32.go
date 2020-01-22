// Code generated - DO NOT EDIT.
// +build atk_2.32

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

// UNSUPPORTED : atk_relation_get_target : no array length

// UNSUPPORTED : atk_component_add_focus_handler : parameter 'handler' is callback

// UNSUPPORTED : atk_table_cell_get_column_header_cells : no array length

// UNSUPPORTED : atk_table_cell_get_row_header_cells : no array length

// UNSUPPORTED : atk_text_get_bounded_ranges : no array length

// UNSUPPORTED : atk_text_get_default_attributes : blacklisted

// UNSUPPORTED : atk_text_get_run_attributes : blacklisted

func Fn_atk_text_scroll_substring_to(paramInstance unsafe.Pointer, param0 int, param1 int, param2 int) bool {
	cValueInstance := (*C.AtkText)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.gint)(param1)

	cValue2 := (C.AtkScrollType)(param2)

	ret := C.atk_text_scroll_substring_to(cValueInstance, cValue0, cValue1, cValue2)

	return toGoBool(ret)
}

func Fn_atk_text_scroll_substring_to_point(paramInstance unsafe.Pointer, param0 int, param1 int, param2 int, param3 int, param4 int) bool {
	cValueInstance := (*C.AtkText)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.gint)(param1)

	cValue2 := (C.AtkCoordType)(param2)

	cValue3 := (C.gint)(param3)

	cValue4 := (C.gint)(param4)

	ret := C.atk_text_scroll_substring_to_point(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4)

	return toGoBool(ret)
}
