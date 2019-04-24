// This is a generated file - DO NOT EDIT
// +build gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import call "github.com/pekim/gobbi/lib/internal/call"

// Unsupported : gtk_binding_entry_add_signal_from_string : return type :

// Unsupported : gtk_cairo_should_draw_window : return type :

// GetBinaryAge is a wrapper around the C function gtk_get_binary_age.
func GetBinaryAge() uint32 {
	data := call.Data{Return: call.Return{Type: call.RT_UINT}}
	call.Function(5902, &data)
	ret := data.Return.Uint32()

	return ret
}

// GetInterfaceAge is a wrapper around the C function gtk_get_interface_age.
func GetInterfaceAge() uint32 {
	data := call.Data{Return: call.Return{Type: call.RT_UINT}}
	call.Function(5910, &data)
	ret := data.Return.Uint32()

	return ret
}

// GetMajorVersion is a wrapper around the C function gtk_get_major_version.
func GetMajorVersion() uint32 {
	data := call.Data{Return: call.Return{Type: call.RT_UINT}}
	call.Function(5912, &data)
	ret := data.Return.Uint32()

	return ret
}

// GetMicroVersion is a wrapper around the C function gtk_get_micro_version.
func GetMicroVersion() uint32 {
	data := call.Data{Return: call.Return{Type: call.RT_UINT}}
	call.Function(5913, &data)
	ret := data.Return.Uint32()

	return ret
}

// GetMinorVersion is a wrapper around the C function gtk_get_minor_version.
func GetMinorVersion() uint32 {
	data := call.Data{Return: call.Return{Type: call.RT_UINT}}
	call.Function(5914, &data)
	ret := data.Return.Uint32()

	return ret
}

// Unsupported : gtk_render_icon_pixbuf : return type :
