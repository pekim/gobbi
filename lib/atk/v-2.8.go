// Code generated - DO NOT EDIT.
// +build atk_2.8 atk_2.10 atk_2.12 atk_2.14 atk_2.30 atk_2.32 atk_2.34

package atk

import atk "github.com/pekim/gobbi/lib/internal/c/atk"

// UNSUPPORTED : atk_add_focus_tracker : parameter 'focus_tracker' is callback

// UNSUPPORTED : atk_add_global_event_listener : parameter 'listener' is callback

// UNSUPPORTED : atk_add_key_event_listener : parameter 'listener' is callback

// UNSUPPORTED : atk_focus_tracker_init : parameter 'init' is callback

// GetBinaryAge wraps the C function atk_get_binary_age.
//
// since 2.8
func GetBinaryAge() uint {
	retSys := atk.Fn_atk_get_binary_age()
	ret := retSys

	return ret
}

// GetInterfaceAge wraps the C function atk_get_interface_age.
//
// since 2.8
func GetInterfaceAge() uint {
	retSys := atk.Fn_atk_get_interface_age()
	ret := retSys

	return ret
}

// GetMajorVersion wraps the C function atk_get_major_version.
//
// since 2.8
func GetMajorVersion() uint {
	retSys := atk.Fn_atk_get_major_version()
	ret := retSys

	return ret
}

// GetMicroVersion wraps the C function atk_get_micro_version.
//
// since 2.8
func GetMicroVersion() uint {
	retSys := atk.Fn_atk_get_micro_version()
	ret := retSys

	return ret
}

// GetMinorVersion wraps the C function atk_get_minor_version.
//
// since 2.8
func GetMinorVersion() uint {
	retSys := atk.Fn_atk_get_minor_version()
	ret := retSys

	return ret
}

// UNSUPPORTED : atk_text_free_ranges : parameter 'ranges' is array parameter without length parameter
