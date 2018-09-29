// This is a generated file - DO NOT EDIT
// +build gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import "unsafe"

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

// Unsupported : gtk_css_section_get_file : no return generator

// FileChooserNativeClass is a wrapper around the C record GtkFileChooserNativeClass.
type FileChooserNativeClass struct {
	native *C.GtkFileChooserNativeClass
	// parent_class : record
}

func FileChooserNativeClassNewFromC(u unsafe.Pointer) *FileChooserNativeClass {
	c := (*C.GtkFileChooserNativeClass)(u)
	if c == nil {
		return nil
	}

	g := &FileChooserNativeClass{native: c}

	return g
}

func (recv *FileChooserNativeClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gtk_gradient_resolve : unsupported parameter resolved_gradient : record with indirection level of 2

// Unsupported : gtk_icon_set_get_sizes : unsupported parameter sizes : no param type

// Unsupported : gtk_icon_set_render_icon : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_icon_set_render_icon_pixbuf : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_icon_set_render_icon_surface : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_icon_source_get_size : no return generator

// Unsupported : gtk_icon_source_set_size : unsupported parameter size : no type generator for gint, GtkIconSize

// NativeDialogClass is a wrapper around the C record GtkNativeDialogClass.
type NativeDialogClass struct {
	native *C.GtkNativeDialogClass
	// parent_class : record
	// no type for response
	// no type for show
	// no type for hide
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func NativeDialogClassNewFromC(u unsafe.Pointer) *NativeDialogClass {
	c := (*C.GtkNativeDialogClass)(u)
	if c == nil {
		return nil
	}

	g := &NativeDialogClass{native: c}

	return g
}

func (recv *NativeDialogClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// PadActionEntry is a wrapper around the C record GtkPadActionEntry.
type PadActionEntry struct {
	native     *C.GtkPadActionEntry
	Type       PadActionType
	Index      int32
	Mode       int32
	Label      string
	ActionName string
}

func PadActionEntryNewFromC(u unsafe.Pointer) *PadActionEntry {
	c := (*C.GtkPadActionEntry)(u)
	if c == nil {
		return nil
	}

	g := &PadActionEntry{
		ActionName: C.GoString(c.action_name),
		Index:      (int32)(c.index),
		Label:      C.GoString(c.label),
		Mode:       (int32)(c.mode),
		Type:       (PadActionType)(c._type),
		native:     c,
	}

	return g
}

func (recv *PadActionEntry) ToC() unsafe.Pointer {
	recv.native._type =
		(C.GtkPadActionType)(recv.Type)
	recv.native.index =
		(C.gint)(recv.Index)
	recv.native.mode =
		(C.gint)(recv.Mode)
	recv.native.label =
		C.CString(recv.Label)
	recv.native.action_name =
		C.CString(recv.ActionName)

	return (unsafe.Pointer)(recv.native)
}

// PadControllerClass is a wrapper around the C record GtkPadControllerClass.
type PadControllerClass struct {
	native *C.GtkPadControllerClass
}

func PadControllerClassNewFromC(u unsafe.Pointer) *PadControllerClass {
	c := (*C.GtkPadControllerClass)(u)
	if c == nil {
		return nil
	}

	g := &PadControllerClass{native: c}

	return g
}

func (recv *PadControllerClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

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

// ShortcutLabelClass is a wrapper around the C record GtkShortcutLabelClass.
type ShortcutLabelClass struct {
	native *C.GtkShortcutLabelClass
}

func ShortcutLabelClassNewFromC(u unsafe.Pointer) *ShortcutLabelClass {
	c := (*C.GtkShortcutLabelClass)(u)
	if c == nil {
		return nil
	}

	g := &ShortcutLabelClass{native: c}

	return g
}

func (recv *ShortcutLabelClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ShortcutsGroupClass is a wrapper around the C record GtkShortcutsGroupClass.
type ShortcutsGroupClass struct {
	native *C.GtkShortcutsGroupClass
}

func ShortcutsGroupClassNewFromC(u unsafe.Pointer) *ShortcutsGroupClass {
	c := (*C.GtkShortcutsGroupClass)(u)
	if c == nil {
		return nil
	}

	g := &ShortcutsGroupClass{native: c}

	return g
}

func (recv *ShortcutsGroupClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ShortcutsSectionClass is a wrapper around the C record GtkShortcutsSectionClass.
type ShortcutsSectionClass struct {
	native *C.GtkShortcutsSectionClass
}

func ShortcutsSectionClassNewFromC(u unsafe.Pointer) *ShortcutsSectionClass {
	c := (*C.GtkShortcutsSectionClass)(u)
	if c == nil {
		return nil
	}

	g := &ShortcutsSectionClass{native: c}

	return g
}

func (recv *ShortcutsSectionClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ShortcutsShortcutClass is a wrapper around the C record GtkShortcutsShortcutClass.
type ShortcutsShortcutClass struct {
	native *C.GtkShortcutsShortcutClass
}

func ShortcutsShortcutClassNewFromC(u unsafe.Pointer) *ShortcutsShortcutClass {
	c := (*C.GtkShortcutsShortcutClass)(u)
	if c == nil {
		return nil
	}

	g := &ShortcutsShortcutClass{native: c}

	return g
}

func (recv *ShortcutsShortcutClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ShortcutsWindowClass is a wrapper around the C record GtkShortcutsWindowClass.
type ShortcutsWindowClass struct {
	native *C.GtkShortcutsWindowClass
	// parent_class : record
	// no type for close
	// no type for search
}

func ShortcutsWindowClassNewFromC(u unsafe.Pointer) *ShortcutsWindowClass {
	c := (*C.GtkShortcutsWindowClass)(u)
	if c == nil {
		return nil
	}

	g := &ShortcutsWindowClass{native: c}

	return g
}

func (recv *ShortcutsWindowClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gtk_target_list_new : unsupported parameter targets : no param type

// Unsupported : gtk_target_list_add : unsupported parameter target : Blacklisted record : GdkAtom

// Unsupported : gtk_target_list_add_table : unsupported parameter targets : no param type

// Unsupported : gtk_target_list_find : unsupported parameter target : Blacklisted record : GdkAtom

// Unsupported : gtk_target_list_remove : unsupported parameter target : Blacklisted record : GdkAtom

// Unsupported : gtk_text_iter_backward_find_char : unsupported parameter pred : no type generator for TextCharPredicate, GtkTextCharPredicate

// Unsupported : gtk_text_iter_forward_find_char : unsupported parameter pred : no type generator for TextCharPredicate, GtkTextCharPredicate

// StartsTag is a wrapper around the C function gtk_text_iter_starts_tag.
func (recv *TextIter) StartsTag(tag *TextTag) bool {
	c_tag := (*C.GtkTextTag)(tag.ToC())

	retC := C.gtk_text_iter_starts_tag((*C.GtkTextIter)(recv.native), c_tag)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_tree_path_new_from_indices : unsupported parameter ... : varargs

// Unsupported : gtk_tree_path_new_from_indicesv : unsupported parameter indices : no param type

// Unsupported : gtk_tree_path_get_indices : no return generator

// Unsupported : gtk_tree_path_get_indices_with_depth : unsupported parameter depth : no type generator for gint, gint*

// Unsupported : gtk_tree_row_reference_new : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

// Unsupported : gtk_tree_row_reference_new_proxy : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

// Unsupported : gtk_tree_row_reference_get_model : no return generator

// Unsupported : gtk_widget_class_bind_template_callback_full : unsupported parameter callback_symbol : no type generator for GObject.Callback, GCallback

// GetCssName is a wrapper around the C function gtk_widget_class_get_css_name.
func (recv *WidgetClass) GetCssName() string {
	retC := C.gtk_widget_class_get_css_name((*C.GtkWidgetClass)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : gtk_widget_class_install_style_property_parser : unsupported parameter parser : no type generator for RcPropertyParser, GtkRcPropertyParser

// Unsupported : gtk_widget_class_list_style_properties : unsupported parameter n_properties : no type generator for guint, guint*

// Unsupported : gtk_widget_class_set_accessible_type : unsupported parameter type : no type generator for GType, GType

// Unsupported : gtk_widget_class_set_connect_func : unsupported parameter connect_func : no type generator for BuilderConnectFunc, GtkBuilderConnectFunc

// SetCssName is a wrapper around the C function gtk_widget_class_set_css_name.
func (recv *WidgetClass) SetCssName(name string) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	C.gtk_widget_class_set_css_name((*C.GtkWidgetClass)(recv.native), c_name)

	return
}

// Unsupported : gtk_widget_path_append_type : unsupported parameter type : no type generator for GType, GType

// Unsupported : gtk_widget_path_get_object_type : no return generator

// Unsupported : gtk_widget_path_has_parent : unsupported parameter type : no type generator for GType, GType

// Unsupported : gtk_widget_path_is_type : unsupported parameter type : no type generator for GType, GType

// IterGetObjectName is a wrapper around the C function gtk_widget_path_iter_get_object_name.
func (recv *WidgetPath) IterGetObjectName(pos int32) string {
	c_pos := (C.gint)(pos)

	retC := C.gtk_widget_path_iter_get_object_name((*C.GtkWidgetPath)(recv.native), c_pos)
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : gtk_widget_path_iter_get_object_type : no return generator

// Unsupported : gtk_widget_path_iter_has_qregion : unsupported parameter flags : GtkRegionFlags* with indirection level of 1

// Unsupported : gtk_widget_path_iter_has_region : unsupported parameter flags : GtkRegionFlags* with indirection level of 1

// IterSetObjectName is a wrapper around the C function gtk_widget_path_iter_set_object_name.
func (recv *WidgetPath) IterSetObjectName(pos int32, name string) {
	c_pos := (C.gint)(pos)

	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	C.gtk_widget_path_iter_set_object_name((*C.GtkWidgetPath)(recv.native), c_pos, c_name)

	return
}

// Unsupported : gtk_widget_path_iter_set_object_type : unsupported parameter type : no type generator for GType, GType

// Unsupported : gtk_widget_path_prepend_type : unsupported parameter type : no type generator for GType, GType
