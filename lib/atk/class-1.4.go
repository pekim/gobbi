// This is a generated file - DO NOT EDIT
// +build atk_1.4 atk_1.6 atk_1.9 atk_1.12 atk_1.13 atk_1.20 atk_1.30 atk_2.8 atk_2.12

package atk

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <atk/atk.h>
// #include <stdlib.h>
import "C"

// IsSelectedLink is a wrapper around the C function atk_hyperlink_is_selected_link.
func (recv *Hyperlink) IsSelectedLink() bool {
	retC := C.atk_hyperlink_is_selected_link((*C.AtkHyperlink)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : atk_relation_new : unsupported parameter targets : no type generator for Object (AtkObject*) for array param targets
