// This is a generated file - DO NOT EDIT
// +build pango_1.16 pango_1.18 pango_1.20 pango_1.22 pango_1.26 pango_1.30 pango_1.31.0 pango_1.32 pango_1.32.4 pango_1.34 pango_1.38
// +build pango_1.16 pango_1.18 pango_1.20 pango_1.22 pango_1.26 pango_1.30 pango_1.31.0 pango_1.32 pango_1.32.4 pango_1.34 pango_1.38
// +build pango_1.16 pango_1.18 pango_1.20 pango_1.22 pango_1.26 pango_1.30 pango_1.31.0 pango_1.32 pango_1.32.4 pango_1.34 pango_1.38
// +build pango_1.16 pango_1.18 pango_1.20 pango_1.22 pango_1.26 pango_1.30 pango_1.31.0 pango_1.32 pango_1.32.4 pango_1.34 pango_1.38
// +build pango_1.16 pango_1.18 pango_1.20 pango_1.22 pango_1.26 pango_1.30 pango_1.31.0 pango_1.32 pango_1.32.4 pango_1.34 pango_1.38
// +build pango_1.16 pango_1.18 pango_1.20 pango_1.22 pango_1.26 pango_1.30 pango_1.31.0 pango_1.32 pango_1.32.4 pango_1.34 pango_1.38
// +build pango_1.16 pango_1.18 pango_1.20 pango_1.22 pango_1.26 pango_1.30 pango_1.31.0 pango_1.32 pango_1.32.4 pango_1.34 pango_1.38
// +build pango_1.16 pango_1.18 pango_1.20 pango_1.22 pango_1.26 pango_1.30 pango_1.31.0 pango_1.32 pango_1.32.4 pango_1.34 pango_1.38
// +build pango_1.16 pango_1.18 pango_1.20 pango_1.22 pango_1.26 pango_1.30 pango_1.31.0 pango_1.32 pango_1.32.4 pango_1.34 pango_1.38

package pango

import call "github.com/pekim/gobbi/lib/internal/call"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <pango/pango.h>
// #include <stdlib.h>
import "C"

const ANALYSIS_FLAG_CENTERED_BASELINE int32 = 1

type Gravity int

const (
	PANGO_GRAVITY_SOUTH Gravity = 0
	PANGO_GRAVITY_EAST  Gravity = 1
	PANGO_GRAVITY_NORTH Gravity = 2
	PANGO_GRAVITY_WEST  Gravity = 3
	PANGO_GRAVITY_AUTO  Gravity = 4
)

// pango_gravity_get_for_matrix : return type :
// pango_gravity_get_for_script : return type :
type GravityHint int

const (
	PANGO_GRAVITY_HINT_NATURAL GravityHint = 0
	PANGO_GRAVITY_HINT_STRONG  GravityHint = 1
	PANGO_GRAVITY_HINT_LINE    GravityHint = 2
)

// Unsupported : pango_attr_gravity_hint_new : return type :

// Unsupported : pango_attr_gravity_new : return type :

// Unsupported : pango_gravity_get_for_matrix : return type :

// Unsupported : pango_gravity_get_for_script : return type :

// Unsupported : pango_language_get_default : return type :

// Unsupported : pango_parse_enum : return type :

// Version is a wrapper around the C function pango_version.
func Version() int32 {
	data := call.Data{
		Params: []call.Value{},
		Return: call.Value{Type: call.TYPE_INT},
	}
	call.Function(9015, &data)
	ret := data.Return.Int32()

	return ret
}

// Unsupported : pango_version_check : return type :

// Unsupported : pango_version_string : return type :
