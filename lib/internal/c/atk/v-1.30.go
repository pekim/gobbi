// Code generated - DO NOT EDIT.
// +build atk_1.30

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

func Fn_atk_plug_new() unsafe.Pointer {
	ret := C.atk_plug_new()

	return unsafe.Pointer(ret)
}

func Fn_atk_plug_get_id(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.AtkPlug)(unsafe.Pointer(paramInstance))

	ret := C.atk_plug_get_id(cValueInstance)

	return C.GoString(ret)
}

// UNSUPPORTED : atk_relation_get_target : no array length

func Fn_atk_socket_embed(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.AtkSocket)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.atk_socket_embed(cValueInstance, cValue0)
}

func Fn_atk_socket_is_occupied(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.AtkSocket)(unsafe.Pointer(paramInstance))

	ret := C.atk_socket_is_occupied(cValueInstance)

	return toGoBool(ret)
}

// UNSUPPORTED : atk_component_add_focus_handler : parameter 'handler' is callback

// UNSUPPORTED : atk_table_cell_get_column_header_cells : no array length

// UNSUPPORTED : atk_table_cell_get_row_header_cells : no array length

// UNSUPPORTED : atk_text_get_bounded_ranges : no array length

// UNSUPPORTED : atk_text_get_default_attributes : blacklisted

// UNSUPPORTED : atk_text_get_run_attributes : blacklisted
