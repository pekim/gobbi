// Code generated - DO NOT EDIT.
// +build atk_1.6 atk_1.9 atk_1.12 atk_1.13 atk_1.20 atk_1.30 atk_2.7.4 atk_2.8 atk_2.10 atk_2.12 atk_2.14 atk_2.30 atk_2.32 atk_2.34

package atk

import atk "github.com/pekim/gobbi/lib/internal/c/atk"

// UNSUPPORTED : atk_add_focus_tracker : parameter 'focus_tracker' is callback

// UNSUPPORTED : atk_add_global_event_listener : parameter 'listener' is callback

// UNSUPPORTED : atk_add_key_event_listener : parameter 'listener' is callback

// UNSUPPORTED : atk_focus_tracker_init : parameter 'init' is callback

// GetFocusObject wraps the C function atk_get_focus_object.
//
// since 1.6
func GetFocusObject() *Object {
	retSys := atk.Fn_atk_get_focus_object()
	ret := ObjectNewFromC(retSys)

	return ret
}

// UNSUPPORTED : atk_text_free_ranges : parameter 'ranges' is array parameter without length parameter
