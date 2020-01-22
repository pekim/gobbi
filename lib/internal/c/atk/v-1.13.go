// Code generated - DO NOT EDIT.
// +build atk_1.13 atk_1.20 atk_1.30 atk_2.7.4 atk_2.8 atk_2.10 atk_2.12 atk_2.14 atk_2.30 atk_2.32 atk_2.34

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

func Fn_atk_misc_threads_enter(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.AtkMisc)(unsafe.Pointer(paramInstance))

	C.atk_misc_threads_enter(cValueInstance)
}

func Fn_atk_misc_threads_leave(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.AtkMisc)(unsafe.Pointer(paramInstance))

	C.atk_misc_threads_leave(cValueInstance)
}

func Fn_atk_misc_get_instance() unsafe.Pointer {
	ret := C.atk_misc_get_instance()

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : atk_object_connect_property_change_handler : parameter 'handler' is callback

// UNSUPPORTED : atk_relation_get_target : no array length

// UNSUPPORTED : atk_component_add_focus_handler : parameter 'handler' is callback

// UNSUPPORTED : atk_table_cell_get_column_header_cells : no array length

// UNSUPPORTED : atk_table_cell_get_row_header_cells : no array length

// UNSUPPORTED : atk_text_get_bounded_ranges : no array length

// UNSUPPORTED : atk_text_get_default_attributes : blacklisted

// UNSUPPORTED : atk_text_get_run_attributes : blacklisted
