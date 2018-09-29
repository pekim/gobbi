// This is a generated file - DO NOT EDIT
// +build gtk_2.10 gtk_2.12 gtk_2.14 gtk_2.16 gtk_2.18 gtk_2.20 gtk_2.22 gtk_2.24 gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	gdkpixbuf "github.com/pekim/gobbi/lib/gdkpixbuf"
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

// Unsupported : gtk_css_section_get_file : no return generator

// Unsupported : gtk_gradient_resolve : unsupported parameter resolved_gradient : record with indirection level of 2

// Unsupported : gtk_icon_set_get_sizes : unsupported parameter sizes : no param type

// Unsupported : gtk_icon_set_render_icon : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_icon_set_render_icon_pixbuf : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_icon_set_render_icon_surface : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_icon_source_get_size : no return generator

// Unsupported : gtk_icon_source_set_size : unsupported parameter size : no type generator for gint, GtkIconSize

// PaperSizeNew is a wrapper around the C function gtk_paper_size_new.
func PaperSizeNew(name string) *PaperSize {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.gtk_paper_size_new(c_name)
	retGo := PaperSizeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// PaperSizeNewCustom is a wrapper around the C function gtk_paper_size_new_custom.
func PaperSizeNewCustom(name string, displayName string, width float64, height float64, unit Unit) *PaperSize {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_display_name := C.CString(displayName)
	defer C.free(unsafe.Pointer(c_display_name))

	c_width := (C.gdouble)(width)

	c_height := (C.gdouble)(height)

	c_unit := (C.GtkUnit)(unit)

	retC := C.gtk_paper_size_new_custom(c_name, c_display_name, c_width, c_height, c_unit)
	retGo := PaperSizeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_paper_size_new_from_gvariant : unsupported parameter variant : Blacklisted record : GVariant

// PaperSizeNewFromPpd is a wrapper around the C function gtk_paper_size_new_from_ppd.
func PaperSizeNewFromPpd(ppdName string, ppdDisplayName string, width float64, height float64) *PaperSize {
	c_ppd_name := C.CString(ppdName)
	defer C.free(unsafe.Pointer(c_ppd_name))

	c_ppd_display_name := C.CString(ppdDisplayName)
	defer C.free(unsafe.Pointer(c_ppd_display_name))

	c_width := (C.gdouble)(width)

	c_height := (C.gdouble)(height)

	retC := C.gtk_paper_size_new_from_ppd(c_ppd_name, c_ppd_display_name, c_width, c_height)
	retGo := PaperSizeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Copy is a wrapper around the C function gtk_paper_size_copy.
func (recv *PaperSize) Copy() *PaperSize {
	retC := C.gtk_paper_size_copy((*C.GtkPaperSize)(recv.native))
	retGo := PaperSizeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Free is a wrapper around the C function gtk_paper_size_free.
func (recv *PaperSize) Free() {
	C.gtk_paper_size_free((*C.GtkPaperSize)(recv.native))

	return
}

// GetDefaultBottomMargin is a wrapper around the C function gtk_paper_size_get_default_bottom_margin.
func (recv *PaperSize) GetDefaultBottomMargin(unit Unit) float64 {
	c_unit := (C.GtkUnit)(unit)

	retC := C.gtk_paper_size_get_default_bottom_margin((*C.GtkPaperSize)(recv.native), c_unit)
	retGo := (float64)(retC)

	return retGo
}

// GetDefaultLeftMargin is a wrapper around the C function gtk_paper_size_get_default_left_margin.
func (recv *PaperSize) GetDefaultLeftMargin(unit Unit) float64 {
	c_unit := (C.GtkUnit)(unit)

	retC := C.gtk_paper_size_get_default_left_margin((*C.GtkPaperSize)(recv.native), c_unit)
	retGo := (float64)(retC)

	return retGo
}

// GetDefaultRightMargin is a wrapper around the C function gtk_paper_size_get_default_right_margin.
func (recv *PaperSize) GetDefaultRightMargin(unit Unit) float64 {
	c_unit := (C.GtkUnit)(unit)

	retC := C.gtk_paper_size_get_default_right_margin((*C.GtkPaperSize)(recv.native), c_unit)
	retGo := (float64)(retC)

	return retGo
}

// GetDefaultTopMargin is a wrapper around the C function gtk_paper_size_get_default_top_margin.
func (recv *PaperSize) GetDefaultTopMargin(unit Unit) float64 {
	c_unit := (C.GtkUnit)(unit)

	retC := C.gtk_paper_size_get_default_top_margin((*C.GtkPaperSize)(recv.native), c_unit)
	retGo := (float64)(retC)

	return retGo
}

// GetDisplayName is a wrapper around the C function gtk_paper_size_get_display_name.
func (recv *PaperSize) GetDisplayName() string {
	retC := C.gtk_paper_size_get_display_name((*C.GtkPaperSize)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetHeight is a wrapper around the C function gtk_paper_size_get_height.
func (recv *PaperSize) GetHeight(unit Unit) float64 {
	c_unit := (C.GtkUnit)(unit)

	retC := C.gtk_paper_size_get_height((*C.GtkPaperSize)(recv.native), c_unit)
	retGo := (float64)(retC)

	return retGo
}

// GetName is a wrapper around the C function gtk_paper_size_get_name.
func (recv *PaperSize) GetName() string {
	retC := C.gtk_paper_size_get_name((*C.GtkPaperSize)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetPpdName is a wrapper around the C function gtk_paper_size_get_ppd_name.
func (recv *PaperSize) GetPpdName() string {
	retC := C.gtk_paper_size_get_ppd_name((*C.GtkPaperSize)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetWidth is a wrapper around the C function gtk_paper_size_get_width.
func (recv *PaperSize) GetWidth(unit Unit) float64 {
	c_unit := (C.GtkUnit)(unit)

	retC := C.gtk_paper_size_get_width((*C.GtkPaperSize)(recv.native), c_unit)
	retGo := (float64)(retC)

	return retGo
}

// IsEqual is a wrapper around the C function gtk_paper_size_is_equal.
func (recv *PaperSize) IsEqual(size2 *PaperSize) bool {
	c_size2 := (*C.GtkPaperSize)(size2.ToC())

	retC := C.gtk_paper_size_is_equal((*C.GtkPaperSize)(recv.native), c_size2)
	retGo := retC == C.TRUE

	return retGo
}

// SetSize is a wrapper around the C function gtk_paper_size_set_size.
func (recv *PaperSize) SetSize(width float64, height float64, unit Unit) {
	c_width := (C.gdouble)(width)

	c_height := (C.gdouble)(height)

	c_unit := (C.GtkUnit)(unit)

	C.gtk_paper_size_set_size((*C.GtkPaperSize)(recv.native), c_width, c_height, c_unit)

	return
}

// Unsupported : gtk_paper_size_to_gvariant : return type : Blacklisted record : GVariant

// RecentInfo is a wrapper around the C record GtkRecentInfo.
type RecentInfo struct {
	native *C.GtkRecentInfo
}

func RecentInfoNewFromC(u unsafe.Pointer) *RecentInfo {
	c := (*C.GtkRecentInfo)(u)
	if c == nil {
		return nil
	}

	g := &RecentInfo{native: c}

	return g
}

func (recv *RecentInfo) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gtk_recent_info_create_app_info : no return generator

// Exists is a wrapper around the C function gtk_recent_info_exists.
func (recv *RecentInfo) Exists() bool {
	retC := C.gtk_recent_info_exists((*C.GtkRecentInfo)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_recent_info_get_added : no return generator

// GetAge is a wrapper around the C function gtk_recent_info_get_age.
func (recv *RecentInfo) GetAge() int32 {
	retC := C.gtk_recent_info_get_age((*C.GtkRecentInfo)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : gtk_recent_info_get_application_info : unsupported parameter count : no type generator for guint, guint*

// Unsupported : gtk_recent_info_get_applications : unsupported parameter length : no type generator for gsize, gsize*

// GetDescription is a wrapper around the C function gtk_recent_info_get_description.
func (recv *RecentInfo) GetDescription() string {
	retC := C.gtk_recent_info_get_description((*C.GtkRecentInfo)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetDisplayName is a wrapper around the C function gtk_recent_info_get_display_name.
func (recv *RecentInfo) GetDisplayName() string {
	retC := C.gtk_recent_info_get_display_name((*C.GtkRecentInfo)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : gtk_recent_info_get_gicon : no return generator

// Unsupported : gtk_recent_info_get_groups : unsupported parameter length : no type generator for gsize, gsize*

// GetIcon is a wrapper around the C function gtk_recent_info_get_icon.
func (recv *RecentInfo) GetIcon(size int32) *gdkpixbuf.Pixbuf {
	c_size := (C.gint)(size)

	retC := C.gtk_recent_info_get_icon((*C.GtkRecentInfo)(recv.native), c_size)
	retGo := gdkpixbuf.PixbufNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetMimeType is a wrapper around the C function gtk_recent_info_get_mime_type.
func (recv *RecentInfo) GetMimeType() string {
	retC := C.gtk_recent_info_get_mime_type((*C.GtkRecentInfo)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : gtk_recent_info_get_modified : no return generator

// GetPrivateHint is a wrapper around the C function gtk_recent_info_get_private_hint.
func (recv *RecentInfo) GetPrivateHint() bool {
	retC := C.gtk_recent_info_get_private_hint((*C.GtkRecentInfo)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetShortName is a wrapper around the C function gtk_recent_info_get_short_name.
func (recv *RecentInfo) GetShortName() string {
	retC := C.gtk_recent_info_get_short_name((*C.GtkRecentInfo)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// GetUri is a wrapper around the C function gtk_recent_info_get_uri.
func (recv *RecentInfo) GetUri() string {
	retC := C.gtk_recent_info_get_uri((*C.GtkRecentInfo)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetUriDisplay is a wrapper around the C function gtk_recent_info_get_uri_display.
func (recv *RecentInfo) GetUriDisplay() string {
	retC := C.gtk_recent_info_get_uri_display((*C.GtkRecentInfo)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_recent_info_get_visited : no return generator

// HasApplication is a wrapper around the C function gtk_recent_info_has_application.
func (recv *RecentInfo) HasApplication(appName string) bool {
	c_app_name := C.CString(appName)
	defer C.free(unsafe.Pointer(c_app_name))

	retC := C.gtk_recent_info_has_application((*C.GtkRecentInfo)(recv.native), c_app_name)
	retGo := retC == C.TRUE

	return retGo
}

// HasGroup is a wrapper around the C function gtk_recent_info_has_group.
func (recv *RecentInfo) HasGroup(groupName string) bool {
	c_group_name := C.CString(groupName)
	defer C.free(unsafe.Pointer(c_group_name))

	retC := C.gtk_recent_info_has_group((*C.GtkRecentInfo)(recv.native), c_group_name)
	retGo := retC == C.TRUE

	return retGo
}

// IsLocal is a wrapper around the C function gtk_recent_info_is_local.
func (recv *RecentInfo) IsLocal() bool {
	retC := C.gtk_recent_info_is_local((*C.GtkRecentInfo)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// LastApplication is a wrapper around the C function gtk_recent_info_last_application.
func (recv *RecentInfo) LastApplication() string {
	retC := C.gtk_recent_info_last_application((*C.GtkRecentInfo)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Match is a wrapper around the C function gtk_recent_info_match.
func (recv *RecentInfo) Match(infoB *RecentInfo) bool {
	c_info_b := (*C.GtkRecentInfo)(infoB.ToC())

	retC := C.gtk_recent_info_match((*C.GtkRecentInfo)(recv.native), c_info_b)
	retGo := retC == C.TRUE

	return retGo
}

// Ref is a wrapper around the C function gtk_recent_info_ref.
func (recv *RecentInfo) Ref() *RecentInfo {
	retC := C.gtk_recent_info_ref((*C.GtkRecentInfo)(recv.native))
	retGo := RecentInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unref is a wrapper around the C function gtk_recent_info_unref.
func (recv *RecentInfo) Unref() {
	C.gtk_recent_info_unref((*C.GtkRecentInfo)(recv.native))

	return
}

// RecentManagerClass is a wrapper around the C record GtkRecentManagerClass.
type RecentManagerClass struct {
	native *C.GtkRecentManagerClass
	// Private : parent_class
	// no type for changed
	// no type for _gtk_recent1
	// no type for _gtk_recent2
	// no type for _gtk_recent3
	// no type for _gtk_recent4
}

func RecentManagerClassNewFromC(u unsafe.Pointer) *RecentManagerClass {
	c := (*C.GtkRecentManagerClass)(u)
	if c == nil {
		return nil
	}

	g := &RecentManagerClass{native: c}

	return g
}

func (recv *RecentManagerClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gtk_selection_data_get_data : no return type

// Unsupported : gtk_selection_data_get_data_type : return type : Blacklisted record : GdkAtom

// Unsupported : gtk_selection_data_get_data_with_length : unsupported parameter length : no type generator for gint, gint*

// Unsupported : gtk_selection_data_get_selection : return type : Blacklisted record : GdkAtom

// Unsupported : gtk_selection_data_get_target : return type : Blacklisted record : GdkAtom

// Unsupported : gtk_selection_data_get_targets : unsupported parameter targets : no param type

// Unsupported : gtk_selection_data_get_uris : no return type

// Unsupported : gtk_selection_data_set : unsupported parameter type : Blacklisted record : GdkAtom

// Unsupported : gtk_selection_data_set_uris : unsupported parameter uris : no param type

// TargetsIncludeRichText is a wrapper around the C function gtk_selection_data_targets_include_rich_text.
func (recv *SelectionData) TargetsIncludeRichText(buffer *TextBuffer) bool {
	c_buffer := (*C.GtkTextBuffer)(buffer.ToC())

	retC := C.gtk_selection_data_targets_include_rich_text((*C.GtkSelectionData)(recv.native), c_buffer)
	retGo := retC == C.TRUE

	return retGo
}

// TargetsIncludeUri is a wrapper around the C function gtk_selection_data_targets_include_uri.
func (recv *SelectionData) TargetsIncludeUri() bool {
	retC := C.gtk_selection_data_targets_include_uri((*C.GtkSelectionData)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_target_list_new : unsupported parameter targets : no param type

// Unsupported : gtk_target_list_add : unsupported parameter target : Blacklisted record : GdkAtom

// AddRichTextTargets is a wrapper around the C function gtk_target_list_add_rich_text_targets.
func (recv *TargetList) AddRichTextTargets(info uint32, deserializable bool, buffer *TextBuffer) {
	c_info := (C.guint)(info)

	c_deserializable :=
		boolToGboolean(deserializable)

	c_buffer := (*C.GtkTextBuffer)(buffer.ToC())

	C.gtk_target_list_add_rich_text_targets((*C.GtkTargetList)(recv.native), c_info, c_deserializable, c_buffer)

	return
}

// Unsupported : gtk_target_list_add_table : unsupported parameter targets : no param type

// Unsupported : gtk_target_list_find : unsupported parameter target : Blacklisted record : GdkAtom

// Unsupported : gtk_target_list_remove : unsupported parameter target : Blacklisted record : GdkAtom

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

// Unsupported : gtk_widget_class_set_accessible_type : unsupported parameter type : no type generator for GType, GType

// Unsupported : gtk_widget_class_set_connect_func : unsupported parameter connect_func : no type generator for BuilderConnectFunc, GtkBuilderConnectFunc

// Unsupported : gtk_widget_path_append_type : unsupported parameter type : no type generator for GType, GType

// Unsupported : gtk_widget_path_get_object_type : no return generator

// Unsupported : gtk_widget_path_has_parent : unsupported parameter type : no type generator for GType, GType

// Unsupported : gtk_widget_path_is_type : unsupported parameter type : no type generator for GType, GType

// Unsupported : gtk_widget_path_iter_get_object_type : no return generator

// Unsupported : gtk_widget_path_iter_has_qregion : unsupported parameter flags : GtkRegionFlags* with indirection level of 1

// Unsupported : gtk_widget_path_iter_has_region : unsupported parameter flags : GtkRegionFlags* with indirection level of 1

// Unsupported : gtk_widget_path_iter_set_object_type : unsupported parameter type : no type generator for GType, GType

// Unsupported : gtk_widget_path_prepend_type : unsupported parameter type : no type generator for GType, GType
