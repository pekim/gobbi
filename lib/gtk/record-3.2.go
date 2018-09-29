// This is a generated file - DO NOT EDIT
// +build gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	atk "github.com/pekim/gobbi/lib/atk"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

// Unsupported : gtk_cell_area_class_list_cell_properties : unsupported parameter n_properties : no type generator for guint, guint*

// Unsupported : gtk_cell_renderer_class_set_accessible_type : unsupported parameter type : no type generator for GType, GType

// Unsupported : gtk_container_class_install_child_properties : unsupported parameter pspecs : no param type

// Unsupported : gtk_container_class_list_child_properties : unsupported parameter n_properties : no type generator for guint, guint*

// CssSection is a wrapper around the C record GtkCssSection.
type CssSection struct {
	native *C.GtkCssSection
}

func CssSectionNewFromC(u unsafe.Pointer) *CssSection {
	c := (*C.GtkCssSection)(u)
	if c == nil {
		return nil
	}

	g := &CssSection{native: c}

	return g
}

func (recv *CssSection) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// GetEndLine is a wrapper around the C function gtk_css_section_get_end_line.
func (recv *CssSection) GetEndLine() uint32 {
	retC := C.gtk_css_section_get_end_line((*C.GtkCssSection)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// GetEndPosition is a wrapper around the C function gtk_css_section_get_end_position.
func (recv *CssSection) GetEndPosition() uint32 {
	retC := C.gtk_css_section_get_end_position((*C.GtkCssSection)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// Unsupported : gtk_css_section_get_file : no return generator

// GetParent is a wrapper around the C function gtk_css_section_get_parent.
func (recv *CssSection) GetParent() *CssSection {
	retC := C.gtk_css_section_get_parent((*C.GtkCssSection)(recv.native))
	retGo := CssSectionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetSectionType is a wrapper around the C function gtk_css_section_get_section_type.
func (recv *CssSection) GetSectionType() CssSectionType {
	retC := C.gtk_css_section_get_section_type((*C.GtkCssSection)(recv.native))
	retGo := (CssSectionType)(retC)

	return retGo
}

// GetStartLine is a wrapper around the C function gtk_css_section_get_start_line.
func (recv *CssSection) GetStartLine() uint32 {
	retC := C.gtk_css_section_get_start_line((*C.GtkCssSection)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// GetStartPosition is a wrapper around the C function gtk_css_section_get_start_position.
func (recv *CssSection) GetStartPosition() uint32 {
	retC := C.gtk_css_section_get_start_position((*C.GtkCssSection)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// Ref is a wrapper around the C function gtk_css_section_ref.
func (recv *CssSection) Ref() *CssSection {
	retC := C.gtk_css_section_ref((*C.GtkCssSection)(recv.native))
	retGo := CssSectionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unref is a wrapper around the C function gtk_css_section_unref.
func (recv *CssSection) Unref() {
	C.gtk_css_section_unref((*C.GtkCssSection)(recv.native))

	return
}

// Unsupported : gtk_gradient_resolve : unsupported parameter resolved_gradient : record with indirection level of 2

// Unsupported : gtk_icon_set_get_sizes : unsupported parameter sizes : no param type

// Unsupported : gtk_icon_set_render_icon : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_icon_set_render_icon_pixbuf : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_icon_set_render_icon_surface : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_icon_source_get_size : no return generator

// Unsupported : gtk_icon_source_set_size : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_paper_size_new_from_gvariant : unsupported parameter variant : Blacklisted record : GVariant

// Unsupported : gtk_paper_size_to_gvariant : return type : Blacklisted record : GVariant

// Unsupported : gtk_recent_info_create_app_info : no return generator

// Unsupported : gtk_recent_info_get_added : no return generator

// Unsupported : gtk_recent_info_get_application_info : unsupported parameter count : no type generator for guint, guint*

// Unsupported : gtk_recent_info_get_applications : unsupported parameter length : no type generator for gsize, gsize*

// Unsupported : gtk_recent_info_get_gicon : no return generator

// Unsupported : gtk_recent_info_get_groups : unsupported parameter length : no type generator for gsize, gsize*

// Unsupported : gtk_recent_info_get_modified : no return generator

// Unsupported : gtk_recent_info_get_visited : no return generator

// Unsupported : gtk_selection_data_get_data : no return type

// Unsupported : gtk_selection_data_get_data_type : return type : Blacklisted record : GdkAtom

// Unsupported : gtk_selection_data_get_data_with_length : unsupported parameter length : no type generator for gint, gint*

// Unsupported : gtk_selection_data_get_selection : return type : Blacklisted record : GdkAtom

// Unsupported : gtk_selection_data_get_target : return type : Blacklisted record : GdkAtom

// Unsupported : gtk_selection_data_get_targets : unsupported parameter targets : no param type

// Unsupported : gtk_selection_data_get_uris : no return type

// Unsupported : gtk_selection_data_set : unsupported parameter type : Blacklisted record : GdkAtom

// Unsupported : gtk_selection_data_set_uris : unsupported parameter uris : no param type

// Blacklisted : GtkStackAccessibleClass

// Unsupported : gtk_target_list_new : unsupported parameter targets : no param type

// Unsupported : gtk_target_list_add : unsupported parameter target : Blacklisted record : GdkAtom

// Unsupported : gtk_target_list_add_table : unsupported parameter targets : no param type

// Unsupported : gtk_target_list_find : unsupported parameter target : Blacklisted record : GdkAtom

// Unsupported : gtk_target_list_remove : unsupported parameter target : Blacklisted record : GdkAtom

// Assign is a wrapper around the C function gtk_text_iter_assign.
func (recv *TextIter) Assign(other *TextIter) {
	c_other := (*C.GtkTextIter)(other.ToC())

	C.gtk_text_iter_assign((*C.GtkTextIter)(recv.native), c_other)

	return
}

// Unsupported : gtk_text_iter_backward_find_char : unsupported parameter pred : no type generator for TextCharPredicate, GtkTextCharPredicate

// Unsupported : gtk_text_iter_forward_find_char : unsupported parameter pred : no type generator for TextCharPredicate, GtkTextCharPredicate

// Unsupported : gtk_tree_path_new_from_indices : unsupported parameter ... : varargs

// Unsupported : gtk_tree_path_new_from_indicesv : unsupported parameter indices : no param type

// Unsupported : gtk_tree_path_get_indices : no return generator

// Unsupported : gtk_tree_path_get_indices_with_depth : unsupported parameter depth : no type generator for gint, gint*

// Unsupported : gtk_tree_row_reference_new : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

// Unsupported : gtk_tree_row_reference_new_proxy : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

// Unsupported : gtk_tree_row_reference_get_model : no return generator

// Unsupported : gtk_widget_class_bind_template_callback_full : unsupported parameter callback_symbol : no type generator for GObject.Callback, GCallback

// Unsupported : gtk_widget_class_install_style_property_parser : unsupported parameter parser : no type generator for RcPropertyParser, GtkRcPropertyParser

// Unsupported : gtk_widget_class_list_style_properties : unsupported parameter n_properties : no type generator for guint, guint*

// SetAccessibleRole is a wrapper around the C function gtk_widget_class_set_accessible_role.
func (recv *WidgetClass) SetAccessibleRole(role atk.Role) {
	c_role := (C.AtkRole)(role)

	C.gtk_widget_class_set_accessible_role((*C.GtkWidgetClass)(recv.native), c_role)

	return
}

// Unsupported : gtk_widget_class_set_accessible_type : unsupported parameter type : no type generator for GType, GType

// Unsupported : gtk_widget_class_set_connect_func : unsupported parameter connect_func : no type generator for BuilderConnectFunc, GtkBuilderConnectFunc

// AppendForWidget is a wrapper around the C function gtk_widget_path_append_for_widget.
func (recv *WidgetPath) AppendForWidget(widget *Widget) int32 {
	c_widget := (*C.GtkWidget)(widget.ToC())

	retC := C.gtk_widget_path_append_for_widget((*C.GtkWidgetPath)(recv.native), c_widget)
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : gtk_widget_path_append_type : unsupported parameter type : no type generator for GType, GType

// AppendWithSiblings is a wrapper around the C function gtk_widget_path_append_with_siblings.
func (recv *WidgetPath) AppendWithSiblings(siblings *WidgetPath, siblingIndex uint32) int32 {
	c_siblings := (*C.GtkWidgetPath)(siblings.ToC())

	c_sibling_index := (C.guint)(siblingIndex)

	retC := C.gtk_widget_path_append_with_siblings((*C.GtkWidgetPath)(recv.native), c_siblings, c_sibling_index)
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : gtk_widget_path_get_object_type : no return generator

// Unsupported : gtk_widget_path_has_parent : unsupported parameter type : no type generator for GType, GType

// Unsupported : gtk_widget_path_is_type : unsupported parameter type : no type generator for GType, GType

// Unsupported : gtk_widget_path_iter_get_object_type : no return generator

// Unsupported : gtk_widget_path_iter_has_qregion : unsupported parameter flags : GtkRegionFlags* with indirection level of 1

// Unsupported : gtk_widget_path_iter_has_region : unsupported parameter flags : GtkRegionFlags* with indirection level of 1

// Unsupported : gtk_widget_path_iter_set_object_type : unsupported parameter type : no type generator for GType, GType

// Unsupported : gtk_widget_path_prepend_type : unsupported parameter type : no type generator for GType, GType

// Ref is a wrapper around the C function gtk_widget_path_ref.
func (recv *WidgetPath) Ref() *WidgetPath {
	retC := C.gtk_widget_path_ref((*C.GtkWidgetPath)(recv.native))
	retGo := WidgetPathNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ToString is a wrapper around the C function gtk_widget_path_to_string.
func (recv *WidgetPath) ToString() string {
	retC := C.gtk_widget_path_to_string((*C.GtkWidgetPath)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Unref is a wrapper around the C function gtk_widget_path_unref.
func (recv *WidgetPath) Unref() {
	C.gtk_widget_path_unref((*C.GtkWidgetPath)(recv.native))

	return
}
