// This is a generated file - DO NOT EDIT
// +build gobject_2.46 gobject_2.54

package gobject

import (
	"C"
	glib "github.com/pekim/gobbi/lib/glib"
)

// GetNameQuark is a wrapper around the C function g_param_spec_get_name_quark.
func (recv *ParamSpec) GetNameQuark() glib.Quark {
	retC := C.g_param_spec_get_name_quark((*C.GParamSpec)(recv.native))
	retGo := (glib.Quark)(retC)

	return retGo
}
