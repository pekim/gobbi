// This is a generated file - DO NOT EDIT
// +build gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	glib "github.com/pekim/gobbi/lib/glib"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

// Unsupported : gtk_icon_set_render_icon_surface : unsupported parameter size : no type generator for gint (GtkIconSize) for param size

// Unsupported : gtk_paper_size_new_from_gvariant : unsupported parameter variant : Blacklisted record : GVariant

// Unsupported : gtk_target_list_new : unsupported parameter targets : no type generator for TargetEntry (GtkTargetEntry) for array param targets

// Unsupported : gtk_tree_path_new_from_indices : unsupported parameter ... : varargs

// Unsupported : gtk_tree_path_new_from_indicesv : unsupported parameter indices : no type generator for gint (gint) for array param indices

// Unsupported : gtk_tree_row_reference_new : unsupported parameter model : no type generator for TreeModel (GtkTreeModel*) for param model

// Unsupported : gtk_tree_row_reference_new_proxy : unsupported parameter model : no type generator for TreeModel (GtkTreeModel*) for param model

// Unsupported : gtk_widget_class_bind_template_callback_full : unsupported parameter callback_symbol : no type generator for GObject.Callback (GCallback) for param callback_symbol

// BindTemplateChildFull is a wrapper around the C function gtk_widget_class_bind_template_child_full.
func (recv *WidgetClass) BindTemplateChildFull(name string, internalChild bool, structOffset int64) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_internal_child :=
		boolToGboolean(internalChild)

	c_struct_offset := (C.gssize)(structOffset)

	C.gtk_widget_class_bind_template_child_full((*C.GtkWidgetClass)(recv.native), c_name, c_internal_child, c_struct_offset)

	return
}

// Unsupported : gtk_widget_class_set_connect_func : unsupported parameter connect_func : no type generator for BuilderConnectFunc (GtkBuilderConnectFunc) for param connect_func

// SetTemplate is a wrapper around the C function gtk_widget_class_set_template.
func (recv *WidgetClass) SetTemplate(templateBytes *glib.Bytes) {
	c_template_bytes := (*C.GBytes)(templateBytes.ToC())

	C.gtk_widget_class_set_template((*C.GtkWidgetClass)(recv.native), c_template_bytes)

	return
}

// SetTemplateFromResource is a wrapper around the C function gtk_widget_class_set_template_from_resource.
func (recv *WidgetClass) SetTemplateFromResource(resourceName string) {
	c_resource_name := C.CString(resourceName)
	defer C.free(unsafe.Pointer(c_resource_name))

	C.gtk_widget_class_set_template_from_resource((*C.GtkWidgetClass)(recv.native), c_resource_name)

	return
}
