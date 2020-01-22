// Code generated - DO NOT EDIT.
// +build atk_2.12 atk_2.14 atk_2.30 atk_2.32 atk_2.34

package atk

import "unsafe"

// UNSUPPORTED : atk_add_focus_tracker : parameter 'focus_tracker' is callback

// UNSUPPORTED : atk_add_global_event_listener : parameter 'listener' is callback

// UNSUPPORTED : atk_add_key_event_listener : parameter 'listener' is callback

// UNSUPPORTED : atk_focus_tracker_init : parameter 'init' is callback

// UNSUPPORTED : atk_text_free_ranges : parameter 'ranges' is array parameter without length parameter

// TableCellIface is a representation of the C record AtkTableCellIface.
//
// since 2.12
type TableCellIface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C AtkTableCellIface that represents the TableCellIface.
func (recv *TableCellIface) ToC() unsafe.Pointer {
	return recv.native
}

// TableCellIfaceNewFromC creates a new TableCellIface from a pointer to the C AtkTableCellIface that represents the TableCellIface.
func TableCellIfaceNewFromC(native unsafe.Pointer) *TableCellIface {
	return &TableCellIface{native: native}
}
