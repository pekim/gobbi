// Code generated - DO NOT EDIT.
// +build atk_1.9 atk_1.12 atk_1.13 atk_1.20 atk_1.30 atk_2.7.4 atk_2.8 atk_2.10 atk_2.12 atk_2.14 atk_2.30 atk_2.32 atk_2.34

package atk

import "unsafe"

// #include <atk/atk.h>
// #include <stdlib.h>
import "C"

// UNSUPPORTED : atk_add_focus_tracker : parameter 'focus_tracker' is callback

// UNSUPPORTED : atk_add_global_event_listener : parameter 'listener' is callback

// UNSUPPORTED : atk_add_key_event_listener : parameter 'listener' is callback

// UNSUPPORTED : atk_focus_tracker_init : parameter 'init' is callback

// UNSUPPORTED : atk_text_free_ranges : parameter 'ranges' is array parameter without length parameter

// UNSUPPORTED : atk_object_connect_property_change_handler : parameter 'handler' is callback

func Fn_atk_relation_add_target(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.AtkRelation)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.AtkObject)(unsafe.Pointer(param0))

	C.atk_relation_add_target(cValueInstance, cValue0)
}

// UNSUPPORTED : atk_relation_get_target : no array length

func Fn_atk_relation_set_add_relation_by_type(paramInstance unsafe.Pointer, param0 int, param1 unsafe.Pointer) {
	cValueInstance := (*C.AtkRelationSet)(unsafe.Pointer(paramInstance))

	cValue0 := (C.AtkRelationType)(param0)

	cValue1 := (*C.AtkObject)(unsafe.Pointer(param1))

	C.atk_relation_set_add_relation_by_type(cValueInstance, cValue0, cValue1)
}

// UNSUPPORTED : atk_component_add_focus_handler : parameter 'handler' is callback

// UNSUPPORTED : atk_table_cell_get_column_header_cells : no array length

// UNSUPPORTED : atk_table_cell_get_row_header_cells : no array length

// UNSUPPORTED : atk_text_get_bounded_ranges : no array length

// UNSUPPORTED : atk_text_get_default_attributes : blacklisted

// UNSUPPORTED : atk_text_get_run_attributes : blacklisted
