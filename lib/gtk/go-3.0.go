// This is a generated file - DO NOT EDIT
// +build gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29
// +build gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29
// +build gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29
// +build gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29
// +build gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29
// +build gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29
// +build gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29
// +build gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29
// +build gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import call "github.com/pekim/gobbi/lib/internal/call"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

// Unsupported : gtk_app_chooser_button_new : return type :

// Unsupported : gtk_app_chooser_dialog_new : return type :

// Unsupported : gtk_app_chooser_dialog_new_for_content_type : return type :

// Unsupported : gtk_app_chooser_widget_new : return type :

// Unsupported : gtk_application_new : return type :

// Unsupported : gtk_box_new : return type :

// Unsupported : gtk_button_box_new : return type :

// Unsupported : gtk_cell_area_box_new : return type :

// Unsupported : gtk_color_button_new_with_rgba : return type :

// Unsupported : gtk_entry_completion_new_with_area : return type :

// Unsupported : gtk_icon_view_new_with_area : return type :

// Unsupported : gtk_paned_new : return type :

// Unsupported : gtk_scale_new : return type :

// Unsupported : gtk_scale_new_with_range : return type :

// Unsupported : gtk_scrollbar_new : return type :

// Unsupported : gtk_separator_new : return type :

// Unsupported : gtk_switch_new : return type :

// Unsupported : gtk_tree_view_column_new_with_area : return type :

type License int

const (
	GTK_LICENSE_UNKNOWN       License = 0
	GTK_LICENSE_CUSTOM        License = 1
	GTK_LICENSE_GPL_2_0       License = 2
	GTK_LICENSE_GPL_3_0       License = 3
	GTK_LICENSE_LGPL_2_1      License = 4
	GTK_LICENSE_LGPL_3_0      License = 5
	GTK_LICENSE_BSD           License = 6
	GTK_LICENSE_MIT_X11       License = 7
	GTK_LICENSE_ARTISTIC      License = 8
	GTK_LICENSE_GPL_2_0_ONLY  License = 9
	GTK_LICENSE_GPL_3_0_ONLY  License = 10
	GTK_LICENSE_LGPL_2_1_ONLY License = 11
	GTK_LICENSE_LGPL_3_0_ONLY License = 12
	GTK_LICENSE_AGPL_3_0      License = 13
	GTK_LICENSE_AGPL_3_0_ONLY License = 14
)

// Unsupported : gtk_binding_entry_add_signal_from_string : return type :

// Unsupported : gtk_cairo_should_draw_window : return type :

// GetBinaryAge is a wrapper around the C function gtk_get_binary_age.
func GetBinaryAge() uint32 {
	data := call.Data{
		Params: []call.Value{},
		Return: call.Value{Type: call.TYPE_UINT},
	}
	call.Function(5902, &data)
	ret := data.Return.Uint32()

	return ret
}

// GetInterfaceAge is a wrapper around the C function gtk_get_interface_age.
func GetInterfaceAge() uint32 {
	data := call.Data{
		Params: []call.Value{},
		Return: call.Value{Type: call.TYPE_UINT},
	}
	call.Function(5910, &data)
	ret := data.Return.Uint32()

	return ret
}

// GetMajorVersion is a wrapper around the C function gtk_get_major_version.
func GetMajorVersion() uint32 {
	data := call.Data{
		Params: []call.Value{},
		Return: call.Value{Type: call.TYPE_UINT},
	}
	call.Function(5912, &data)
	ret := data.Return.Uint32()

	return ret
}

// GetMicroVersion is a wrapper around the C function gtk_get_micro_version.
func GetMicroVersion() uint32 {
	data := call.Data{
		Params: []call.Value{},
		Return: call.Value{Type: call.TYPE_UINT},
	}
	call.Function(5913, &data)
	ret := data.Return.Uint32()

	return ret
}

// GetMinorVersion is a wrapper around the C function gtk_get_minor_version.
func GetMinorVersion() uint32 {
	data := call.Data{
		Params: []call.Value{},
		Return: call.Value{Type: call.TYPE_UINT},
	}
	call.Function(5914, &data)
	ret := data.Return.Uint32()

	return ret
}

// Unsupported : gtk_render_icon_pixbuf : return type :

// Unsupported : gtk_gradient_new_linear : return type :

// Unsupported : gtk_gradient_new_radial : return type :

// Unsupported : gtk_requisition_new : return type :

// Unsupported : gtk_symbolic_color_new_alpha : return type :

// Unsupported : gtk_symbolic_color_new_literal : return type :

// Unsupported : gtk_symbolic_color_new_mix : return type :

// Unsupported : gtk_symbolic_color_new_name : return type :

// Unsupported : gtk_symbolic_color_new_shade : return type :

// Unsupported : gtk_widget_path_new : return type :
