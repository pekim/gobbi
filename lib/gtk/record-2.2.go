// This is a generated file - DO NOT EDIT
// +build gtk_2.2 gtk_2.4 gtk_2.6 gtk_2.8 gtk_2.10 gtk_2.12 gtk_2.14 gtk_2.16 gtk_2.18 gtk_2.20 gtk_2.22 gtk_2.24 gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	gobject "github.com/pekim/gobbi/lib/gobject"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

// Unsupported : gtk_paper_size_new_from_gvariant : unsupported parameter variant : Blacklisted record : GVariant

// Unsupported : gtk_target_list_new : unsupported parameter targets : no param type

// Unsupported : gtk_tree_path_new_from_indices : unsupported parameter ... : varargs

// Unsupported : gtk_tree_path_new_from_indicesv : unsupported parameter indices : no param type

// Unsupported : gtk_tree_row_reference_new : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

// Unsupported : gtk_tree_row_reference_new_proxy : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

// Copy is a wrapper around the C function gtk_tree_row_reference_copy.
func (recv *TreeRowReference) Copy() *TreeRowReference {
	retC := C.gtk_tree_row_reference_copy((*C.GtkTreeRowReference)(recv.native))
	retGo := TreeRowReferenceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// FindStyleProperty is a wrapper around the C function gtk_widget_class_find_style_property.
func (recv *WidgetClass) FindStyleProperty(propertyName string) *gobject.ParamSpec {
	c_property_name := C.CString(propertyName)
	defer C.free(unsafe.Pointer(c_property_name))

	retC := C.gtk_widget_class_find_style_property((*C.GtkWidgetClass)(recv.native), c_property_name)
	retGo := gobject.ParamSpecNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_widget_class_list_style_properties : no return type