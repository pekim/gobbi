// Code generated - DO NOT EDIT.
// +build atk_1.20

package atk

import atk "github.com/pekim/gobbi/lib/internal/c/atk"

// UNSUPPORTED : atk_add_focus_tracker : parameter 'focus_tracker' is callback

// UNSUPPORTED : atk_add_global_event_listener : parameter 'listener' is callback

// UNSUPPORTED : atk_add_key_event_listener : parameter 'listener' is callback

// UNSUPPORTED : atk_focus_tracker_init : parameter 'init' is callback

// GetVersion wraps the C function atk_get_version.
//
// since 1.20
func GetVersion() string {
	retSys := atk.Fn_atk_get_version()
	ret := retSys

	return ret
}

// UNSUPPORTED : atk_text_free_ranges : parameter 'ranges' is array parameter without length parameter
