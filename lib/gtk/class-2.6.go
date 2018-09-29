// This is a generated file - DO NOT EDIT
// +build gtk_2.6 gtk_2.8 gtk_2.10 gtk_2.12 gtk_2.14 gtk_2.16 gtk_2.18 gtk_2.20 gtk_2.22 gtk_2.24 gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	atk "github.com/pekim/gobbi/lib/atk"
	gdk "github.com/pekim/gobbi/lib/gdk"
	gdkpixbuf "github.com/pekim/gobbi/lib/gdkpixbuf"
	glib "github.com/pekim/gobbi/lib/glib"
	pango "github.com/pekim/gobbi/lib/pango"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

// AboutDialogNew is a wrapper around the C function gtk_about_dialog_new.
func AboutDialogNew() *Widget {
	retC := C.gtk_about_dialog_new()
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_about_dialog_add_credit_section : unsupported parameter people : no param type

// Unsupported : gtk_about_dialog_get_artists : no return type

// Unsupported : gtk_about_dialog_get_authors : no return type

// GetComments is a wrapper around the C function gtk_about_dialog_get_comments.
func (recv *AboutDialog) GetComments() string {
	retC := C.gtk_about_dialog_get_comments((*C.GtkAboutDialog)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetCopyright is a wrapper around the C function gtk_about_dialog_get_copyright.
func (recv *AboutDialog) GetCopyright() string {
	retC := C.gtk_about_dialog_get_copyright((*C.GtkAboutDialog)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : gtk_about_dialog_get_documenters : no return type

// GetLicense is a wrapper around the C function gtk_about_dialog_get_license.
func (recv *AboutDialog) GetLicense() string {
	retC := C.gtk_about_dialog_get_license((*C.GtkAboutDialog)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetLogo is a wrapper around the C function gtk_about_dialog_get_logo.
func (recv *AboutDialog) GetLogo() *gdkpixbuf.Pixbuf {
	retC := C.gtk_about_dialog_get_logo((*C.GtkAboutDialog)(recv.native))
	retGo := gdkpixbuf.PixbufNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetLogoIconName is a wrapper around the C function gtk_about_dialog_get_logo_icon_name.
func (recv *AboutDialog) GetLogoIconName() string {
	retC := C.gtk_about_dialog_get_logo_icon_name((*C.GtkAboutDialog)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetTranslatorCredits is a wrapper around the C function gtk_about_dialog_get_translator_credits.
func (recv *AboutDialog) GetTranslatorCredits() string {
	retC := C.gtk_about_dialog_get_translator_credits((*C.GtkAboutDialog)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetVersion is a wrapper around the C function gtk_about_dialog_get_version.
func (recv *AboutDialog) GetVersion() string {
	retC := C.gtk_about_dialog_get_version((*C.GtkAboutDialog)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetWebsite is a wrapper around the C function gtk_about_dialog_get_website.
func (recv *AboutDialog) GetWebsite() string {
	retC := C.gtk_about_dialog_get_website((*C.GtkAboutDialog)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetWebsiteLabel is a wrapper around the C function gtk_about_dialog_get_website_label.
func (recv *AboutDialog) GetWebsiteLabel() string {
	retC := C.gtk_about_dialog_get_website_label((*C.GtkAboutDialog)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : gtk_about_dialog_set_artists : unsupported parameter artists : no param type

// Unsupported : gtk_about_dialog_set_authors : unsupported parameter authors : no param type

// SetComments is a wrapper around the C function gtk_about_dialog_set_comments.
func (recv *AboutDialog) SetComments(comments string) {
	c_comments := C.CString(comments)
	defer C.free(unsafe.Pointer(c_comments))

	C.gtk_about_dialog_set_comments((*C.GtkAboutDialog)(recv.native), c_comments)

	return
}

// SetCopyright is a wrapper around the C function gtk_about_dialog_set_copyright.
func (recv *AboutDialog) SetCopyright(copyright string) {
	c_copyright := C.CString(copyright)
	defer C.free(unsafe.Pointer(c_copyright))

	C.gtk_about_dialog_set_copyright((*C.GtkAboutDialog)(recv.native), c_copyright)

	return
}

// Unsupported : gtk_about_dialog_set_documenters : unsupported parameter documenters : no param type

// SetLicense is a wrapper around the C function gtk_about_dialog_set_license.
func (recv *AboutDialog) SetLicense(license string) {
	c_license := C.CString(license)
	defer C.free(unsafe.Pointer(c_license))

	C.gtk_about_dialog_set_license((*C.GtkAboutDialog)(recv.native), c_license)

	return
}

// SetLogo is a wrapper around the C function gtk_about_dialog_set_logo.
func (recv *AboutDialog) SetLogo(logo *gdkpixbuf.Pixbuf) {
	c_logo := (*C.GdkPixbuf)(logo.ToC())

	C.gtk_about_dialog_set_logo((*C.GtkAboutDialog)(recv.native), c_logo)

	return
}

// SetLogoIconName is a wrapper around the C function gtk_about_dialog_set_logo_icon_name.
func (recv *AboutDialog) SetLogoIconName(iconName string) {
	c_icon_name := C.CString(iconName)
	defer C.free(unsafe.Pointer(c_icon_name))

	C.gtk_about_dialog_set_logo_icon_name((*C.GtkAboutDialog)(recv.native), c_icon_name)

	return
}

// SetTranslatorCredits is a wrapper around the C function gtk_about_dialog_set_translator_credits.
func (recv *AboutDialog) SetTranslatorCredits(translatorCredits string) {
	c_translator_credits := C.CString(translatorCredits)
	defer C.free(unsafe.Pointer(c_translator_credits))

	C.gtk_about_dialog_set_translator_credits((*C.GtkAboutDialog)(recv.native), c_translator_credits)

	return
}

// SetVersion is a wrapper around the C function gtk_about_dialog_set_version.
func (recv *AboutDialog) SetVersion(version string) {
	c_version := C.CString(version)
	defer C.free(unsafe.Pointer(c_version))

	C.gtk_about_dialog_set_version((*C.GtkAboutDialog)(recv.native), c_version)

	return
}

// SetWebsite is a wrapper around the C function gtk_about_dialog_set_website.
func (recv *AboutDialog) SetWebsite(website string) {
	c_website := C.CString(website)
	defer C.free(unsafe.Pointer(c_website))

	C.gtk_about_dialog_set_website((*C.GtkAboutDialog)(recv.native), c_website)

	return
}

// SetWebsiteLabel is a wrapper around the C function gtk_about_dialog_set_website_label.
func (recv *AboutDialog) SetWebsiteLabel(websiteLabel string) {
	c_website_label := C.CString(websiteLabel)
	defer C.free(unsafe.Pointer(c_website_label))

	C.gtk_about_dialog_set_website_label((*C.GtkAboutDialog)(recv.native), c_website_label)

	return
}

// Unsupported : gtk_accel_group_find : unsupported parameter find_func : no type generator for AccelGroupFindFunc, GtkAccelGroupFindFunc

// Unsupported : gtk_accel_group_query : unsupported parameter n_entries : no type generator for guint, guint*

// Unsupported : gtk_accel_label_get_accel : unsupported parameter accelerator_key : no type generator for guint, guint*

// Unsupported : gtk_action_create_icon : unsupported parameter icon_size : no type generator for gint, GtkIconSize

// GetAccelPath is a wrapper around the C function gtk_action_get_accel_path.
func (recv *Action) GetAccelPath() string {
	retC := C.gtk_action_get_accel_path((*C.GtkAction)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : gtk_action_get_gicon : no return generator

// Unsupported : gtk_action_set_gicon : unsupported parameter icon : no type generator for Gio.Icon, GIcon*

// SetSensitive is a wrapper around the C function gtk_action_set_sensitive.
func (recv *Action) SetSensitive(sensitive bool) {
	c_sensitive :=
		boolToGboolean(sensitive)

	C.gtk_action_set_sensitive((*C.GtkAction)(recv.native), c_sensitive)

	return
}

// SetVisible is a wrapper around the C function gtk_action_set_visible.
func (recv *Action) SetVisible(visible bool) {
	c_visible :=
		boolToGboolean(visible)

	C.gtk_action_set_visible((*C.GtkAction)(recv.native), c_visible)

	return
}

// Unsupported : gtk_action_group_add_actions : unsupported parameter entries : no param type

// Unsupported : gtk_action_group_add_actions_full : unsupported parameter entries : no param type

// Unsupported : gtk_action_group_add_radio_actions : unsupported parameter entries : no param type

// Unsupported : gtk_action_group_add_radio_actions_full : unsupported parameter entries : no param type

// Unsupported : gtk_action_group_add_toggle_actions : unsupported parameter entries : no param type

// Unsupported : gtk_action_group_add_toggle_actions_full : unsupported parameter entries : no param type

// Unsupported : gtk_action_group_set_translate_func : unsupported parameter func : no type generator for TranslateFunc, GtkTranslateFunc

// TranslateString is a wrapper around the C function gtk_action_group_translate_string.
func (recv *ActionGroup) TranslateString(string string) string {
	c_string := C.CString(string)
	defer C.free(unsafe.Pointer(c_string))

	retC := C.gtk_action_group_translate_string((*C.GtkActionGroup)(recv.native), c_string)
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : gtk_alignment_get_padding : unsupported parameter padding_top : no type generator for guint, guint*

// Unsupported : gtk_app_chooser_button_append_custom_item : unsupported parameter icon : no type generator for Gio.Icon, GIcon*

// Unsupported : gtk_app_chooser_dialog_new : unsupported parameter file : no type generator for Gio.File, GFile*

// Unsupported : gtk_application_add_accelerator : unsupported parameter parameter : Blacklisted record : GVariant

// Unsupported : gtk_application_get_accels_for_action : no return type

// Unsupported : gtk_application_get_actions_for_accel : no return type

// Unsupported : gtk_application_list_action_descriptions : no return type

// Unsupported : gtk_application_remove_accelerator : unsupported parameter parameter : Blacklisted record : GVariant

// Unsupported : gtk_application_set_accels_for_action : unsupported parameter accels : no param type

// Unsupported : gtk_assistant_set_forward_page_func : unsupported parameter page_func : no type generator for AssistantPageFunc, GtkAssistantPageFunc

// Unsupported : gtk_box_query_child_packing : unsupported parameter padding : no type generator for guint, guint*

// Unsupported : gtk_builder_add_callback_symbol : unsupported parameter callback_symbol : no type generator for GObject.Callback, GCallback

// Unsupported : gtk_builder_add_callback_symbols : unsupported parameter first_callback_symbol : no type generator for GObject.Callback, GCallback

// Unsupported : gtk_builder_add_objects_from_file : unsupported parameter object_ids : no param type

// Unsupported : gtk_builder_add_objects_from_resource : unsupported parameter object_ids : no param type

// Unsupported : gtk_builder_add_objects_from_string : unsupported parameter object_ids : no param type

// Unsupported : gtk_builder_connect_signals_full : unsupported parameter func : no type generator for BuilderConnectFunc, GtkBuilderConnectFunc

// Unsupported : gtk_builder_extend_with_template : unsupported parameter template_type : no type generator for GType, GType

// Unsupported : gtk_builder_get_type_from_name : no return generator

// Unsupported : gtk_builder_lookup_callback_symbol : no return generator

// Unsupported : gtk_builder_value_from_string_type : unsupported parameter type : no type generator for GType, GType

// Unsupported : gtk_button_new_from_icon_name : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_button_get_alignment : unsupported parameter xalign : no type generator for gfloat, gfloat*

// GetImage is a wrapper around the C function gtk_button_get_image.
func (recv *Button) GetImage() *Widget {
	retC := C.gtk_button_get_image((*C.GtkButton)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SetImage is a wrapper around the C function gtk_button_set_image.
func (recv *Button) SetImage(image *Widget) {
	c_image := (*C.GtkWidget)(image.ToC())

	C.gtk_button_set_image((*C.GtkButton)(recv.native), c_image)

	return
}

// Unsupported : gtk_calendar_get_date : unsupported parameter year : no type generator for guint, guint*

// Unsupported : gtk_calendar_set_detail_func : unsupported parameter func : no type generator for CalendarDetailFunc, GtkCalendarDetailFunc

// Unsupported : gtk_cell_area_activate : unsupported parameter cell_area : Blacklisted record : GdkRectangle

// Unsupported : gtk_cell_area_activate_cell : unsupported parameter event : no type generator for Gdk.Event, GdkEvent*

// Unsupported : gtk_cell_area_add_with_properties : unsupported parameter ... : varargs

// Unsupported : gtk_cell_area_apply_attributes : unsupported parameter tree_model : no type generator for TreeModel, GtkTreeModel*

// Unsupported : gtk_cell_area_cell_get : unsupported parameter ... : varargs

// Unsupported : gtk_cell_area_cell_get_valist : unsupported parameter var_args : no type generator for va_list, va_list

// Unsupported : gtk_cell_area_cell_set : unsupported parameter ... : varargs

// Unsupported : gtk_cell_area_cell_set_valist : unsupported parameter var_args : no type generator for va_list, va_list

// Unsupported : gtk_cell_area_event : unsupported parameter event : no type generator for Gdk.Event, GdkEvent*

// Unsupported : gtk_cell_area_foreach : unsupported parameter callback : no type generator for CellCallback, GtkCellCallback

// Unsupported : gtk_cell_area_foreach_alloc : unsupported parameter cell_area : Blacklisted record : GdkRectangle

// Unsupported : gtk_cell_area_get_cell_allocation : unsupported parameter cell_area : Blacklisted record : GdkRectangle

// Unsupported : gtk_cell_area_get_cell_at_position : unsupported parameter cell_area : Blacklisted record : GdkRectangle

// Unsupported : gtk_cell_area_get_edit_widget : no return generator

// Unsupported : gtk_cell_area_get_preferred_height : unsupported parameter minimum_height : no type generator for gint, gint*

// Unsupported : gtk_cell_area_get_preferred_height_for_width : unsupported parameter minimum_height : no type generator for gint, gint*

// Unsupported : gtk_cell_area_get_preferred_width : unsupported parameter minimum_width : no type generator for gint, gint*

// Unsupported : gtk_cell_area_get_preferred_width_for_height : unsupported parameter minimum_width : no type generator for gint, gint*

// Unsupported : gtk_cell_area_inner_cell_area : unsupported parameter cell_area : Blacklisted record : GdkRectangle

// Unsupported : gtk_cell_area_render : unsupported parameter background_area : Blacklisted record : GdkRectangle

// Unsupported : gtk_cell_area_request_renderer : unsupported parameter minimum_size : no type generator for gint, gint*

// Unsupported : gtk_cell_area_context_get_allocation : unsupported parameter width : no type generator for gint, gint*

// Unsupported : gtk_cell_area_context_get_preferred_height : unsupported parameter minimum_height : no type generator for gint, gint*

// Unsupported : gtk_cell_area_context_get_preferred_height_for_width : unsupported parameter minimum_height : no type generator for gint, gint*

// Unsupported : gtk_cell_area_context_get_preferred_width : unsupported parameter minimum_width : no type generator for gint, gint*

// Unsupported : gtk_cell_area_context_get_preferred_width_for_height : unsupported parameter minimum_width : no type generator for gint, gint*

// Unsupported : gtk_cell_renderer_activate : unsupported parameter event : no type generator for Gdk.Event, GdkEvent*

// Unsupported : gtk_cell_renderer_get_aligned_area : unsupported parameter cell_area : Blacklisted record : GdkRectangle

// Unsupported : gtk_cell_renderer_get_alignment : unsupported parameter xalign : no type generator for gfloat, gfloat*

// Unsupported : gtk_cell_renderer_get_fixed_size : unsupported parameter width : no type generator for gint, gint*

// Unsupported : gtk_cell_renderer_get_padding : unsupported parameter xpad : no type generator for gint, gint*

// Unsupported : gtk_cell_renderer_get_preferred_height : unsupported parameter minimum_size : no type generator for gint, gint*

// Unsupported : gtk_cell_renderer_get_preferred_height_for_width : unsupported parameter minimum_height : no type generator for gint, gint*

// Unsupported : gtk_cell_renderer_get_preferred_width : unsupported parameter minimum_size : no type generator for gint, gint*

// Unsupported : gtk_cell_renderer_get_preferred_width_for_height : unsupported parameter minimum_width : no type generator for gint, gint*

// Unsupported : gtk_cell_renderer_get_size : unsupported parameter cell_area : Blacklisted record : GdkRectangle

// Unsupported : gtk_cell_renderer_render : unsupported parameter background_area : Blacklisted record : GdkRectangle

// Unsupported : gtk_cell_renderer_start_editing : unsupported parameter event : no type generator for Gdk.Event, GdkEvent*

// StopEditing is a wrapper around the C function gtk_cell_renderer_stop_editing.
func (recv *CellRenderer) StopEditing(canceled bool) {
	c_canceled :=
		boolToGboolean(canceled)

	C.gtk_cell_renderer_stop_editing((*C.GtkCellRenderer)(recv.native), c_canceled)

	return
}

// CellRendererComboNew is a wrapper around the C function gtk_cell_renderer_combo_new.
func CellRendererComboNew() *CellRenderer {
	retC := C.gtk_cell_renderer_combo_new()
	retGo := CellRendererNewFromC(unsafe.Pointer(retC))

	return retGo
}

// CellRendererProgressNew is a wrapper around the C function gtk_cell_renderer_progress_new.
func CellRendererProgressNew() *CellRenderer {
	retC := C.gtk_cell_renderer_progress_new()
	retGo := CellRendererNewFromC(unsafe.Pointer(retC))

	return retGo
}

// CellViewNew is a wrapper around the C function gtk_cell_view_new.
func CellViewNew() *Widget {
	retC := C.gtk_cell_view_new()
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// CellViewNewWithContext is a wrapper around the C function gtk_cell_view_new_with_context.
func CellViewNewWithContext(area *CellArea, context *CellAreaContext) *Widget {
	c_area := (*C.GtkCellArea)(area.ToC())

	c_context := (*C.GtkCellAreaContext)(context.ToC())

	retC := C.gtk_cell_view_new_with_context(c_area, c_context)
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// CellViewNewWithMarkup is a wrapper around the C function gtk_cell_view_new_with_markup.
func CellViewNewWithMarkup(markup string) *Widget {
	c_markup := C.CString(markup)
	defer C.free(unsafe.Pointer(c_markup))

	retC := C.gtk_cell_view_new_with_markup(c_markup)
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// CellViewNewWithPixbuf is a wrapper around the C function gtk_cell_view_new_with_pixbuf.
func CellViewNewWithPixbuf(pixbuf *gdkpixbuf.Pixbuf) *Widget {
	c_pixbuf := (*C.GdkPixbuf)(pixbuf.ToC())

	retC := C.gtk_cell_view_new_with_pixbuf(c_pixbuf)
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// CellViewNewWithText is a wrapper around the C function gtk_cell_view_new_with_text.
func CellViewNewWithText(text string) *Widget {
	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))

	retC := C.gtk_cell_view_new_with_text(c_text)
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetDisplayedRow is a wrapper around the C function gtk_cell_view_get_displayed_row.
func (recv *CellView) GetDisplayedRow() *TreePath {
	retC := C.gtk_cell_view_get_displayed_row((*C.GtkCellView)(recv.native))
	retGo := TreePathNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_cell_view_get_model : no return generator

// GetSizeOfRow is a wrapper around the C function gtk_cell_view_get_size_of_row.
func (recv *CellView) GetSizeOfRow(path *TreePath) (bool, *Requisition) {
	c_path := (*C.GtkTreePath)(path.ToC())

	var c_requisition C.GtkRequisition

	retC := C.gtk_cell_view_get_size_of_row((*C.GtkCellView)(recv.native), c_path, &c_requisition)
	retGo := retC == C.TRUE

	requisition := RequisitionNewFromC(unsafe.Pointer(&c_requisition))

	return retGo, requisition
}

// SetBackgroundColor is a wrapper around the C function gtk_cell_view_set_background_color.
func (recv *CellView) SetBackgroundColor(color *gdk.Color) {
	c_color := (*C.GdkColor)(color.ToC())

	C.gtk_cell_view_set_background_color((*C.GtkCellView)(recv.native), c_color)

	return
}

// SetDisplayedRow is a wrapper around the C function gtk_cell_view_set_displayed_row.
func (recv *CellView) SetDisplayedRow(path *TreePath) {
	c_path := (*C.GtkTreePath)(path.ToC())

	C.gtk_cell_view_set_displayed_row((*C.GtkCellView)(recv.native), c_path)

	return
}

// Unsupported : gtk_cell_view_set_model : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

// Unsupported : gtk_clipboard_get_selection : return type : Blacklisted record : GdkAtom

// Unsupported : gtk_clipboard_request_contents : unsupported parameter target : Blacklisted record : GdkAtom

// Unsupported : gtk_clipboard_request_image : unsupported parameter callback : no type generator for ClipboardImageReceivedFunc, GtkClipboardImageReceivedFunc

// Unsupported : gtk_clipboard_request_rich_text : unsupported parameter callback : no type generator for ClipboardRichTextReceivedFunc, GtkClipboardRichTextReceivedFunc

// Unsupported : gtk_clipboard_request_targets : unsupported parameter callback : no type generator for ClipboardTargetsReceivedFunc, GtkClipboardTargetsReceivedFunc

// Unsupported : gtk_clipboard_request_text : unsupported parameter callback : no type generator for ClipboardTextReceivedFunc, GtkClipboardTextReceivedFunc

// Unsupported : gtk_clipboard_request_uris : unsupported parameter callback : no type generator for ClipboardURIReceivedFunc, GtkClipboardURIReceivedFunc

// Unsupported : gtk_clipboard_set_can_store : unsupported parameter targets : no param type

// SetImage is a wrapper around the C function gtk_clipboard_set_image.
func (recv *Clipboard) SetImage(pixbuf *gdkpixbuf.Pixbuf) {
	c_pixbuf := (*C.GdkPixbuf)(pixbuf.ToC())

	C.gtk_clipboard_set_image((*C.GtkClipboard)(recv.native), c_pixbuf)

	return
}

// Unsupported : gtk_clipboard_set_with_data : unsupported parameter targets : no param type

// Unsupported : gtk_clipboard_set_with_owner : unsupported parameter targets : no param type

// Store is a wrapper around the C function gtk_clipboard_store.
func (recv *Clipboard) Store() {
	C.gtk_clipboard_store((*C.GtkClipboard)(recv.native))

	return
}

// Unsupported : gtk_clipboard_wait_for_contents : unsupported parameter target : Blacklisted record : GdkAtom

// WaitForImage is a wrapper around the C function gtk_clipboard_wait_for_image.
func (recv *Clipboard) WaitForImage() *gdkpixbuf.Pixbuf {
	retC := C.gtk_clipboard_wait_for_image((*C.GtkClipboard)(recv.native))
	retGo := gdkpixbuf.PixbufNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_clipboard_wait_for_rich_text : unsupported parameter format : Blacklisted record : GdkAtom

// Unsupported : gtk_clipboard_wait_for_targets : unsupported parameter targets : no param type

// Unsupported : gtk_clipboard_wait_for_uris : no return type

// WaitIsImageAvailable is a wrapper around the C function gtk_clipboard_wait_is_image_available.
func (recv *Clipboard) WaitIsImageAvailable() bool {
	retC := C.gtk_clipboard_wait_is_image_available((*C.GtkClipboard)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_clipboard_wait_is_target_available : unsupported parameter target : Blacklisted record : GdkAtom

// Unsupported : gtk_combo_box_new_with_model : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

// Unsupported : gtk_combo_box_new_with_model_and_entry : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

// GetColumnSpanColumn is a wrapper around the C function gtk_combo_box_get_column_span_column.
func (recv *ComboBox) GetColumnSpanColumn() int32 {
	retC := C.gtk_combo_box_get_column_span_column((*C.GtkComboBox)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetFocusOnClick is a wrapper around the C function gtk_combo_box_get_focus_on_click.
func (recv *ComboBox) GetFocusOnClick() bool {
	retC := C.gtk_combo_box_get_focus_on_click((*C.GtkComboBox)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_combo_box_get_model : no return generator

// GetPopupAccessible is a wrapper around the C function gtk_combo_box_get_popup_accessible.
func (recv *ComboBox) GetPopupAccessible() *atk.Object {
	retC := C.gtk_combo_box_get_popup_accessible((*C.GtkComboBox)(recv.native))
	retGo := atk.ObjectNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_combo_box_get_row_separator_func : no return generator

// GetRowSpanColumn is a wrapper around the C function gtk_combo_box_get_row_span_column.
func (recv *ComboBox) GetRowSpanColumn() int32 {
	retC := C.gtk_combo_box_get_row_span_column((*C.GtkComboBox)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetWrapWidth is a wrapper around the C function gtk_combo_box_get_wrap_width.
func (recv *ComboBox) GetWrapWidth() int32 {
	retC := C.gtk_combo_box_get_wrap_width((*C.GtkComboBox)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// SetAddTearoffs is a wrapper around the C function gtk_combo_box_set_add_tearoffs.
func (recv *ComboBox) SetAddTearoffs(addTearoffs bool) {
	c_add_tearoffs :=
		boolToGboolean(addTearoffs)

	C.gtk_combo_box_set_add_tearoffs((*C.GtkComboBox)(recv.native), c_add_tearoffs)

	return
}

// SetFocusOnClick is a wrapper around the C function gtk_combo_box_set_focus_on_click.
func (recv *ComboBox) SetFocusOnClick(focusOnClick bool) {
	c_focus_on_click :=
		boolToGboolean(focusOnClick)

	C.gtk_combo_box_set_focus_on_click((*C.GtkComboBox)(recv.native), c_focus_on_click)

	return
}

// Unsupported : gtk_combo_box_set_model : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

// Unsupported : gtk_combo_box_set_row_separator_func : unsupported parameter func : no type generator for TreeViewRowSeparatorFunc, GtkTreeViewRowSeparatorFunc

// Unsupported : gtk_container_add_with_properties : unsupported parameter ... : varargs

// Unsupported : gtk_container_child_get : unsupported parameter ... : varargs

// Unsupported : gtk_container_child_get_valist : unsupported parameter var_args : no type generator for va_list, va_list

// Unsupported : gtk_container_child_set : unsupported parameter ... : varargs

// Unsupported : gtk_container_child_set_valist : unsupported parameter var_args : no type generator for va_list, va_list

// Unsupported : gtk_container_child_type : no return generator

// Unsupported : gtk_container_forall : unsupported parameter callback : no type generator for Callback, GtkCallback

// Unsupported : gtk_container_foreach : unsupported parameter callback : no type generator for Callback, GtkCallback

// Unsupported : gtk_container_get_focus_chain : unsupported parameter focusable_widgets : record with indirection level of 2

// Unsupported : gtk_css_provider_load_from_data : unsupported parameter data : no param type

// Unsupported : gtk_css_provider_load_from_file : unsupported parameter file : no type generator for Gio.File, GFile*

// Unsupported : gtk_dialog_new_with_buttons : unsupported parameter ... : varargs

// Unsupported : gtk_dialog_add_buttons : unsupported parameter ... : varargs

// Unsupported : gtk_dialog_set_alternative_button_order : unsupported parameter ... : varargs

// Unsupported : gtk_dialog_set_alternative_button_order_from_array : unsupported parameter new_order : no param type

// Unsupported : gtk_entry_get_icon_area : unsupported parameter icon_area : Blacklisted record : GdkRectangle

// Unsupported : gtk_entry_get_icon_gicon : no return generator

// Unsupported : gtk_entry_get_layout_offsets : unsupported parameter x : no type generator for gint, gint*

// Unsupported : gtk_entry_get_text_area : unsupported parameter text_area : Blacklisted record : GdkRectangle

// Unsupported : gtk_entry_set_icon_from_gicon : unsupported parameter icon : no type generator for Gio.Icon, GIcon*

// GetInlineCompletion is a wrapper around the C function gtk_entry_completion_get_inline_completion.
func (recv *EntryCompletion) GetInlineCompletion() bool {
	retC := C.gtk_entry_completion_get_inline_completion((*C.GtkEntryCompletion)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_entry_completion_get_model : no return generator

// GetPopupCompletion is a wrapper around the C function gtk_entry_completion_get_popup_completion.
func (recv *EntryCompletion) GetPopupCompletion() bool {
	retC := C.gtk_entry_completion_get_popup_completion((*C.GtkEntryCompletion)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetTextColumn is a wrapper around the C function gtk_entry_completion_get_text_column.
func (recv *EntryCompletion) GetTextColumn() int32 {
	retC := C.gtk_entry_completion_get_text_column((*C.GtkEntryCompletion)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// InsertPrefix is a wrapper around the C function gtk_entry_completion_insert_prefix.
func (recv *EntryCompletion) InsertPrefix() {
	C.gtk_entry_completion_insert_prefix((*C.GtkEntryCompletion)(recv.native))

	return
}

// SetInlineCompletion is a wrapper around the C function gtk_entry_completion_set_inline_completion.
func (recv *EntryCompletion) SetInlineCompletion(inlineCompletion bool) {
	c_inline_completion :=
		boolToGboolean(inlineCompletion)

	C.gtk_entry_completion_set_inline_completion((*C.GtkEntryCompletion)(recv.native), c_inline_completion)

	return
}

// Unsupported : gtk_entry_completion_set_match_func : unsupported parameter func : no type generator for EntryCompletionMatchFunc, GtkEntryCompletionMatchFunc

// Unsupported : gtk_entry_completion_set_model : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

// SetPopupCompletion is a wrapper around the C function gtk_entry_completion_set_popup_completion.
func (recv *EntryCompletion) SetPopupCompletion(popupCompletion bool) {
	c_popup_completion :=
		boolToGboolean(popupCompletion)

	C.gtk_entry_completion_set_popup_completion((*C.GtkEntryCompletion)(recv.native), c_popup_completion)

	return
}

// Unsupported : EntryIconAccessible : no CType

// Unsupported : gtk_event_controller_handle_event : unsupported parameter event : no type generator for Gdk.Event, const GdkEvent*

// FileChooserButtonNew is a wrapper around the C function gtk_file_chooser_button_new.
func FileChooserButtonNew(title string, action FileChooserAction) *Widget {
	c_title := C.CString(title)
	defer C.free(unsafe.Pointer(c_title))

	c_action := (C.GtkFileChooserAction)(action)

	retC := C.gtk_file_chooser_button_new(c_title, c_action)
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// FileChooserButtonNewWithDialog is a wrapper around the C function gtk_file_chooser_button_new_with_dialog.
func FileChooserButtonNewWithDialog(dialog *Dialog) *Widget {
	c_dialog := (*C.GtkWidget)(dialog.ToC())

	retC := C.gtk_file_chooser_button_new_with_dialog(c_dialog)
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetTitle is a wrapper around the C function gtk_file_chooser_button_get_title.
func (recv *FileChooserButton) GetTitle() string {
	retC := C.gtk_file_chooser_button_get_title((*C.GtkFileChooserButton)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetWidthChars is a wrapper around the C function gtk_file_chooser_button_get_width_chars.
func (recv *FileChooserButton) GetWidthChars() int32 {
	retC := C.gtk_file_chooser_button_get_width_chars((*C.GtkFileChooserButton)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// SetTitle is a wrapper around the C function gtk_file_chooser_button_set_title.
func (recv *FileChooserButton) SetTitle(title string) {
	c_title := C.CString(title)
	defer C.free(unsafe.Pointer(c_title))

	C.gtk_file_chooser_button_set_title((*C.GtkFileChooserButton)(recv.native), c_title)

	return
}

// SetWidthChars is a wrapper around the C function gtk_file_chooser_button_set_width_chars.
func (recv *FileChooserButton) SetWidthChars(nChars int32) {
	c_n_chars := (C.gint)(nChars)

	C.gtk_file_chooser_button_set_width_chars((*C.GtkFileChooserButton)(recv.native), c_n_chars)

	return
}

// Unsupported : gtk_file_chooser_dialog_new : unsupported parameter ... : varargs

// Unsupported : gtk_file_filter_new_from_gvariant : unsupported parameter variant : Blacklisted record : GVariant

// Unsupported : gtk_file_filter_add_custom : unsupported parameter func : no type generator for FileFilterFunc, GtkFileFilterFunc

// AddPixbufFormats is a wrapper around the C function gtk_file_filter_add_pixbuf_formats.
func (recv *FileFilter) AddPixbufFormats() {
	C.gtk_file_filter_add_pixbuf_formats((*C.GtkFileFilter)(recv.native))

	return
}

// Unsupported : gtk_file_filter_to_gvariant : return type : Blacklisted record : GVariant

// Unsupported : gtk_flow_box_bind_model : unsupported parameter model : no type generator for Gio.ListModel, GListModel*

// Unsupported : gtk_flow_box_selected_foreach : unsupported parameter func : no type generator for FlowBoxForeachFunc, GtkFlowBoxForeachFunc

// Unsupported : gtk_flow_box_set_filter_func : unsupported parameter filter_func : no type generator for FlowBoxFilterFunc, GtkFlowBoxFilterFunc

// Unsupported : gtk_flow_box_set_sort_func : unsupported parameter sort_func : no type generator for FlowBoxSortFunc, GtkFlowBoxSortFunc

// Unsupported : gtk_frame_get_label_align : unsupported parameter xalign : no type generator for gfloat, gfloat*

// Unsupported : gtk_gl_area_get_required_version : unsupported parameter major : no type generator for gint, gint*

// Unsupported : gtk_gesture_get_bounding_box : unsupported parameter rect : Blacklisted record : GdkRectangle

// Unsupported : gtk_gesture_get_bounding_box_center : unsupported parameter x : no type generator for gdouble, gdouble*

// Unsupported : gtk_gesture_get_last_event : no return generator

// Unsupported : gtk_gesture_get_point : unsupported parameter x : no type generator for gdouble, gdouble*

// Unsupported : gtk_gesture_drag_get_offset : unsupported parameter x : no type generator for gdouble, gdouble*

// Unsupported : gtk_gesture_drag_get_start_point : unsupported parameter x : no type generator for gdouble, gdouble*

// Unsupported : gtk_gesture_multi_press_get_area : unsupported parameter rect : Blacklisted record : GdkRectangle

// Unsupported : gtk_gesture_multi_press_set_area : unsupported parameter rect : Blacklisted record : GdkRectangle

// Unsupported : gtk_gesture_swipe_get_velocity : unsupported parameter velocity_x : no type generator for gdouble, gdouble*

// Unsupported : gtk_hsv_get_color : unsupported parameter h : no type generator for gdouble, gdouble*

// Unsupported : gtk_hsv_get_metrics : unsupported parameter size : no type generator for gint, gint*

// Unsupported : gtk_im_context_get_preedit_string : unsupported parameter attrs : record with indirection level of 2

// Unsupported : gtk_im_context_get_surrounding : unsupported parameter cursor_index : no type generator for gint, gint*

// Unsupported : gtk_im_context_set_cursor_location : unsupported parameter area : Blacklisted record : GdkRectangle

// Unsupported : gtk_im_context_simple_add_table : unsupported parameter data : no param type

// Unsupported : gtk_icon_info_get_attach_points : unsupported parameter points : no param type

// Unsupported : gtk_icon_info_get_embedded_rect : unsupported parameter rectangle : Blacklisted record : GdkRectangle

// Unsupported : gtk_icon_info_load_icon_async : unsupported parameter callback : no type generator for Gio.AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : gtk_icon_info_load_icon_finish : unsupported parameter res : no type generator for Gio.AsyncResult, GAsyncResult*

// Unsupported : gtk_icon_info_load_symbolic_async : unsupported parameter callback : no type generator for Gio.AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : gtk_icon_info_load_symbolic_finish : unsupported parameter res : no type generator for Gio.AsyncResult, GAsyncResult*

// Unsupported : gtk_icon_info_load_symbolic_for_context_async : unsupported parameter callback : no type generator for Gio.AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : gtk_icon_info_load_symbolic_for_context_finish : unsupported parameter res : no type generator for Gio.AsyncResult, GAsyncResult*

// Unsupported : gtk_icon_theme_choose_icon : unsupported parameter icon_names : no param type

// Unsupported : gtk_icon_theme_choose_icon_for_scale : unsupported parameter icon_names : no param type

// Unsupported : gtk_icon_theme_get_icon_sizes : no return type

// Unsupported : gtk_icon_theme_get_search_path : unsupported parameter path : no param type

// Unsupported : gtk_icon_theme_lookup_by_gicon : unsupported parameter icon : no type generator for Gio.Icon, GIcon*

// Unsupported : gtk_icon_theme_lookup_by_gicon_for_scale : unsupported parameter icon : no type generator for Gio.Icon, GIcon*

// Unsupported : gtk_icon_theme_set_search_path : unsupported parameter path : no param type

// IconViewNew is a wrapper around the C function gtk_icon_view_new.
func IconViewNew() *Widget {
	retC := C.gtk_icon_view_new()
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_icon_view_new_with_model : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

// Unsupported : gtk_icon_view_convert_widget_to_bin_window_coords : unsupported parameter bx : no type generator for gint, gint*

// Unsupported : gtk_icon_view_enable_model_drag_dest : unsupported parameter targets : no param type

// Unsupported : gtk_icon_view_enable_model_drag_source : unsupported parameter targets : no param type

// Unsupported : gtk_icon_view_get_cell_rect : unsupported parameter rect : Blacklisted record : GdkRectangle

// GetColumnSpacing is a wrapper around the C function gtk_icon_view_get_column_spacing.
func (recv *IconView) GetColumnSpacing() int32 {
	retC := C.gtk_icon_view_get_column_spacing((*C.GtkIconView)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetColumns is a wrapper around the C function gtk_icon_view_get_columns.
func (recv *IconView) GetColumns() int32 {
	retC := C.gtk_icon_view_get_columns((*C.GtkIconView)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : gtk_icon_view_get_cursor : unsupported parameter path : record with indirection level of 2

// Unsupported : gtk_icon_view_get_dest_item_at_pos : unsupported parameter path : record with indirection level of 2

// Unsupported : gtk_icon_view_get_drag_dest_item : unsupported parameter path : record with indirection level of 2

// Unsupported : gtk_icon_view_get_item_at_pos : unsupported parameter path : record with indirection level of 2

// GetItemOrientation is a wrapper around the C function gtk_icon_view_get_item_orientation.
func (recv *IconView) GetItemOrientation() Orientation {
	retC := C.gtk_icon_view_get_item_orientation((*C.GtkIconView)(recv.native))
	retGo := (Orientation)(retC)

	return retGo
}

// GetItemWidth is a wrapper around the C function gtk_icon_view_get_item_width.
func (recv *IconView) GetItemWidth() int32 {
	retC := C.gtk_icon_view_get_item_width((*C.GtkIconView)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetMargin is a wrapper around the C function gtk_icon_view_get_margin.
func (recv *IconView) GetMargin() int32 {
	retC := C.gtk_icon_view_get_margin((*C.GtkIconView)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetMarkupColumn is a wrapper around the C function gtk_icon_view_get_markup_column.
func (recv *IconView) GetMarkupColumn() int32 {
	retC := C.gtk_icon_view_get_markup_column((*C.GtkIconView)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : gtk_icon_view_get_model : no return generator

// GetPathAtPos is a wrapper around the C function gtk_icon_view_get_path_at_pos.
func (recv *IconView) GetPathAtPos(x int32, y int32) *TreePath {
	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	retC := C.gtk_icon_view_get_path_at_pos((*C.GtkIconView)(recv.native), c_x, c_y)
	retGo := TreePathNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetPixbufColumn is a wrapper around the C function gtk_icon_view_get_pixbuf_column.
func (recv *IconView) GetPixbufColumn() int32 {
	retC := C.gtk_icon_view_get_pixbuf_column((*C.GtkIconView)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetRowSpacing is a wrapper around the C function gtk_icon_view_get_row_spacing.
func (recv *IconView) GetRowSpacing() int32 {
	retC := C.gtk_icon_view_get_row_spacing((*C.GtkIconView)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetSelectedItems is a wrapper around the C function gtk_icon_view_get_selected_items.
func (recv *IconView) GetSelectedItems() *glib.List {
	retC := C.gtk_icon_view_get_selected_items((*C.GtkIconView)(recv.native))
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetSelectionMode is a wrapper around the C function gtk_icon_view_get_selection_mode.
func (recv *IconView) GetSelectionMode() SelectionMode {
	retC := C.gtk_icon_view_get_selection_mode((*C.GtkIconView)(recv.native))
	retGo := (SelectionMode)(retC)

	return retGo
}

// GetSpacing is a wrapper around the C function gtk_icon_view_get_spacing.
func (recv *IconView) GetSpacing() int32 {
	retC := C.gtk_icon_view_get_spacing((*C.GtkIconView)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetTextColumn is a wrapper around the C function gtk_icon_view_get_text_column.
func (recv *IconView) GetTextColumn() int32 {
	retC := C.gtk_icon_view_get_text_column((*C.GtkIconView)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : gtk_icon_view_get_tooltip_context : unsupported parameter x : no type generator for gint, gint*

// Unsupported : gtk_icon_view_get_visible_range : unsupported parameter start_path : record with indirection level of 2

// ItemActivated is a wrapper around the C function gtk_icon_view_item_activated.
func (recv *IconView) ItemActivated(path *TreePath) {
	c_path := (*C.GtkTreePath)(path.ToC())

	C.gtk_icon_view_item_activated((*C.GtkIconView)(recv.native), c_path)

	return
}

// PathIsSelected is a wrapper around the C function gtk_icon_view_path_is_selected.
func (recv *IconView) PathIsSelected(path *TreePath) bool {
	c_path := (*C.GtkTreePath)(path.ToC())

	retC := C.gtk_icon_view_path_is_selected((*C.GtkIconView)(recv.native), c_path)
	retGo := retC == C.TRUE

	return retGo
}

// SelectAll is a wrapper around the C function gtk_icon_view_select_all.
func (recv *IconView) SelectAll() {
	C.gtk_icon_view_select_all((*C.GtkIconView)(recv.native))

	return
}

// SelectPath is a wrapper around the C function gtk_icon_view_select_path.
func (recv *IconView) SelectPath(path *TreePath) {
	c_path := (*C.GtkTreePath)(path.ToC())

	C.gtk_icon_view_select_path((*C.GtkIconView)(recv.native), c_path)

	return
}

// Unsupported : gtk_icon_view_selected_foreach : unsupported parameter func : no type generator for IconViewForeachFunc, GtkIconViewForeachFunc

// SetColumnSpacing is a wrapper around the C function gtk_icon_view_set_column_spacing.
func (recv *IconView) SetColumnSpacing(columnSpacing int32) {
	c_column_spacing := (C.gint)(columnSpacing)

	C.gtk_icon_view_set_column_spacing((*C.GtkIconView)(recv.native), c_column_spacing)

	return
}

// SetColumns is a wrapper around the C function gtk_icon_view_set_columns.
func (recv *IconView) SetColumns(columns int32) {
	c_columns := (C.gint)(columns)

	C.gtk_icon_view_set_columns((*C.GtkIconView)(recv.native), c_columns)

	return
}

// SetItemOrientation is a wrapper around the C function gtk_icon_view_set_item_orientation.
func (recv *IconView) SetItemOrientation(orientation Orientation) {
	c_orientation := (C.GtkOrientation)(orientation)

	C.gtk_icon_view_set_item_orientation((*C.GtkIconView)(recv.native), c_orientation)

	return
}

// SetItemWidth is a wrapper around the C function gtk_icon_view_set_item_width.
func (recv *IconView) SetItemWidth(itemWidth int32) {
	c_item_width := (C.gint)(itemWidth)

	C.gtk_icon_view_set_item_width((*C.GtkIconView)(recv.native), c_item_width)

	return
}

// SetMargin is a wrapper around the C function gtk_icon_view_set_margin.
func (recv *IconView) SetMargin(margin int32) {
	c_margin := (C.gint)(margin)

	C.gtk_icon_view_set_margin((*C.GtkIconView)(recv.native), c_margin)

	return
}

// SetMarkupColumn is a wrapper around the C function gtk_icon_view_set_markup_column.
func (recv *IconView) SetMarkupColumn(column int32) {
	c_column := (C.gint)(column)

	C.gtk_icon_view_set_markup_column((*C.GtkIconView)(recv.native), c_column)

	return
}

// Unsupported : gtk_icon_view_set_model : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

// SetPixbufColumn is a wrapper around the C function gtk_icon_view_set_pixbuf_column.
func (recv *IconView) SetPixbufColumn(column int32) {
	c_column := (C.gint)(column)

	C.gtk_icon_view_set_pixbuf_column((*C.GtkIconView)(recv.native), c_column)

	return
}

// SetRowSpacing is a wrapper around the C function gtk_icon_view_set_row_spacing.
func (recv *IconView) SetRowSpacing(rowSpacing int32) {
	c_row_spacing := (C.gint)(rowSpacing)

	C.gtk_icon_view_set_row_spacing((*C.GtkIconView)(recv.native), c_row_spacing)

	return
}

// SetSelectionMode is a wrapper around the C function gtk_icon_view_set_selection_mode.
func (recv *IconView) SetSelectionMode(mode SelectionMode) {
	c_mode := (C.GtkSelectionMode)(mode)

	C.gtk_icon_view_set_selection_mode((*C.GtkIconView)(recv.native), c_mode)

	return
}

// SetSpacing is a wrapper around the C function gtk_icon_view_set_spacing.
func (recv *IconView) SetSpacing(spacing int32) {
	c_spacing := (C.gint)(spacing)

	C.gtk_icon_view_set_spacing((*C.GtkIconView)(recv.native), c_spacing)

	return
}

// SetTextColumn is a wrapper around the C function gtk_icon_view_set_text_column.
func (recv *IconView) SetTextColumn(column int32) {
	c_column := (C.gint)(column)

	C.gtk_icon_view_set_text_column((*C.GtkIconView)(recv.native), c_column)

	return
}

// UnselectAll is a wrapper around the C function gtk_icon_view_unselect_all.
func (recv *IconView) UnselectAll() {
	C.gtk_icon_view_unselect_all((*C.GtkIconView)(recv.native))

	return
}

// UnselectPath is a wrapper around the C function gtk_icon_view_unselect_path.
func (recv *IconView) UnselectPath(path *TreePath) {
	c_path := (*C.GtkTreePath)(path.ToC())

	C.gtk_icon_view_unselect_path((*C.GtkIconView)(recv.native), c_path)

	return
}

// Unsupported : gtk_image_new_from_gicon : unsupported parameter icon : no type generator for Gio.Icon, GIcon*

// Unsupported : gtk_image_new_from_icon_name : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_image_new_from_icon_set : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_image_new_from_stock : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_image_get_gicon : unsupported parameter gicon : no type generator for Gio.Icon, GIcon**

// Unsupported : gtk_image_get_icon_name : unsupported parameter size : no type generator for gint, GtkIconSize*

// Unsupported : gtk_image_get_icon_set : unsupported parameter icon_set : record with indirection level of 2

// GetPixelSize is a wrapper around the C function gtk_image_get_pixel_size.
func (recv *Image) GetPixelSize() int32 {
	retC := C.gtk_image_get_pixel_size((*C.GtkImage)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : gtk_image_get_stock : unsupported parameter size : no type generator for gint, GtkIconSize*

// Unsupported : gtk_image_set_from_gicon : unsupported parameter icon : no type generator for Gio.Icon, GIcon*

// Unsupported : gtk_image_set_from_icon_name : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_image_set_from_icon_set : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_image_set_from_stock : unsupported parameter size : no type generator for gint, GtkIconSize

// SetPixelSize is a wrapper around the C function gtk_image_set_pixel_size.
func (recv *Image) SetPixelSize(pixelSize int32) {
	c_pixel_size := (C.gint)(pixelSize)

	C.gtk_image_set_pixel_size((*C.GtkImage)(recv.native), c_pixel_size)

	return
}

// Unsupported : gtk_info_bar_new_with_buttons : unsupported parameter ... : varargs

// Unsupported : gtk_info_bar_add_buttons : unsupported parameter ... : varargs

// GetAngle is a wrapper around the C function gtk_label_get_angle.
func (recv *Label) GetAngle() float64 {
	retC := C.gtk_label_get_angle((*C.GtkLabel)(recv.native))
	retGo := (float64)(retC)

	return retGo
}

// GetEllipsize is a wrapper around the C function gtk_label_get_ellipsize.
func (recv *Label) GetEllipsize() pango.EllipsizeMode {
	retC := C.gtk_label_get_ellipsize((*C.GtkLabel)(recv.native))
	retGo := (pango.EllipsizeMode)(retC)

	return retGo
}

// Unsupported : gtk_label_get_layout_offsets : unsupported parameter x : no type generator for gint, gint*

// GetMaxWidthChars is a wrapper around the C function gtk_label_get_max_width_chars.
func (recv *Label) GetMaxWidthChars() int32 {
	retC := C.gtk_label_get_max_width_chars((*C.GtkLabel)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : gtk_label_get_selection_bounds : unsupported parameter start : no type generator for gint, gint*

// GetSingleLineMode is a wrapper around the C function gtk_label_get_single_line_mode.
func (recv *Label) GetSingleLineMode() bool {
	retC := C.gtk_label_get_single_line_mode((*C.GtkLabel)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetWidthChars is a wrapper around the C function gtk_label_get_width_chars.
func (recv *Label) GetWidthChars() int32 {
	retC := C.gtk_label_get_width_chars((*C.GtkLabel)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// SetAngle is a wrapper around the C function gtk_label_set_angle.
func (recv *Label) SetAngle(angle float64) {
	c_angle := (C.gdouble)(angle)

	C.gtk_label_set_angle((*C.GtkLabel)(recv.native), c_angle)

	return
}

// SetEllipsize is a wrapper around the C function gtk_label_set_ellipsize.
func (recv *Label) SetEllipsize(mode pango.EllipsizeMode) {
	c_mode := (C.PangoEllipsizeMode)(mode)

	C.gtk_label_set_ellipsize((*C.GtkLabel)(recv.native), c_mode)

	return
}

// SetMaxWidthChars is a wrapper around the C function gtk_label_set_max_width_chars.
func (recv *Label) SetMaxWidthChars(nChars int32) {
	c_n_chars := (C.gint)(nChars)

	C.gtk_label_set_max_width_chars((*C.GtkLabel)(recv.native), c_n_chars)

	return
}

// SetSingleLineMode is a wrapper around the C function gtk_label_set_single_line_mode.
func (recv *Label) SetSingleLineMode(singleLineMode bool) {
	c_single_line_mode :=
		boolToGboolean(singleLineMode)

	C.gtk_label_set_single_line_mode((*C.GtkLabel)(recv.native), c_single_line_mode)

	return
}

// SetWidthChars is a wrapper around the C function gtk_label_set_width_chars.
func (recv *Label) SetWidthChars(nChars int32) {
	c_n_chars := (C.gint)(nChars)

	C.gtk_label_set_width_chars((*C.GtkLabel)(recv.native), c_n_chars)

	return
}

// Unsupported : gtk_layout_get_size : unsupported parameter width : no type generator for guint, guint*

// Unsupported : gtk_level_bar_get_offset_value : unsupported parameter value : no type generator for gdouble, gdouble*

// Unsupported : gtk_list_box_bind_model : unsupported parameter model : no type generator for Gio.ListModel, GListModel*

// Unsupported : gtk_list_box_selected_foreach : unsupported parameter func : no type generator for ListBoxForeachFunc, GtkListBoxForeachFunc

// Unsupported : gtk_list_box_set_filter_func : unsupported parameter filter_func : no type generator for ListBoxFilterFunc, GtkListBoxFilterFunc

// Unsupported : gtk_list_box_set_header_func : unsupported parameter update_header : no type generator for ListBoxUpdateHeaderFunc, GtkListBoxUpdateHeaderFunc

// Unsupported : gtk_list_box_set_sort_func : unsupported parameter sort_func : no type generator for ListBoxSortFunc, GtkListBoxSortFunc

// Unsupported : gtk_list_store_new : unsupported parameter ... : varargs

// Unsupported : gtk_list_store_newv : unsupported parameter types : no param type

// Unsupported : gtk_list_store_insert_with_values : unsupported parameter ... : varargs

// Unsupported : gtk_list_store_insert_with_valuesv : unsupported parameter columns : no param type

// Unsupported : gtk_list_store_reorder : unsupported parameter new_order : no param type

// Unsupported : gtk_list_store_set : unsupported parameter ... : varargs

// Unsupported : gtk_list_store_set_column_types : unsupported parameter types : no param type

// Unsupported : gtk_list_store_set_valist : unsupported parameter var_args : no type generator for va_list, va_list

// Unsupported : gtk_list_store_set_valuesv : unsupported parameter columns : no param type

// Unsupported : gtk_menu_attach_to_widget : unsupported parameter detacher : no type generator for MenuDetachFunc, GtkMenuDetachFunc

// Unsupported : gtk_menu_popup : unsupported parameter func : no type generator for MenuPositionFunc, GtkMenuPositionFunc

// Unsupported : gtk_menu_popup_at_pointer : unsupported parameter trigger_event : no type generator for Gdk.Event, const GdkEvent*

// Unsupported : gtk_menu_popup_at_rect : unsupported parameter rect : Blacklisted record : GdkRectangle

// Unsupported : gtk_menu_popup_at_widget : unsupported parameter trigger_event : no type generator for Gdk.Event, const GdkEvent*

// Unsupported : gtk_menu_popup_for_device : unsupported parameter func : no type generator for MenuPositionFunc, GtkMenuPositionFunc

// Unsupported : gtk_menu_item_toggle_size_request : unsupported parameter requisition : no type generator for gint, gint*

// MenuToolButtonNew is a wrapper around the C function gtk_menu_tool_button_new.
func MenuToolButtonNew(iconWidget *Widget, label string) *ToolItem {
	c_icon_widget := (*C.GtkWidget)(iconWidget.ToC())

	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	retC := C.gtk_menu_tool_button_new(c_icon_widget, c_label)
	retGo := ToolItemNewFromC(unsafe.Pointer(retC))

	return retGo
}

// MenuToolButtonNewFromStock is a wrapper around the C function gtk_menu_tool_button_new_from_stock.
func MenuToolButtonNewFromStock(stockId string) *ToolItem {
	c_stock_id := C.CString(stockId)
	defer C.free(unsafe.Pointer(c_stock_id))

	retC := C.gtk_menu_tool_button_new_from_stock(c_stock_id)
	retGo := ToolItemNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetMenu is a wrapper around the C function gtk_menu_tool_button_get_menu.
func (recv *MenuToolButton) GetMenu() *Widget {
	retC := C.gtk_menu_tool_button_get_menu((*C.GtkMenuToolButton)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SetMenu is a wrapper around the C function gtk_menu_tool_button_set_menu.
func (recv *MenuToolButton) SetMenu(menu *Widget) {
	c_menu := (*C.GtkWidget)(menu.ToC())

	C.gtk_menu_tool_button_set_menu((*C.GtkMenuToolButton)(recv.native), c_menu)

	return
}

// Unsupported : gtk_message_dialog_new : unsupported parameter ... : varargs

// Unsupported : gtk_message_dialog_new_with_markup : unsupported parameter ... : varargs

// Unsupported : gtk_message_dialog_format_secondary_markup : unsupported parameter ... : varargs

// Unsupported : gtk_message_dialog_format_secondary_text : unsupported parameter ... : varargs

// Unsupported : gtk_misc_get_alignment : unsupported parameter xalign : no type generator for gfloat, gfloat*

// Unsupported : gtk_misc_get_padding : unsupported parameter xpad : no type generator for gint, gint*

// Unsupported : gtk_numerable_icon_get_background_gicon : no return generator

// Unsupported : gtk_numerable_icon_set_background_gicon : unsupported parameter icon : no type generator for Gio.Icon, GIcon*

// Unsupported : gtk_pad_controller_new : unsupported parameter group : no type generator for Gio.ActionGroup, GActionGroup*

// Unsupported : gtk_pad_controller_set_action_entries : unsupported parameter entries : no param type

// Unsupported : gtk_page_setup_new_from_gvariant : unsupported parameter variant : Blacklisted record : GVariant

// Unsupported : gtk_page_setup_to_gvariant : return type : Blacklisted record : GVariant

// Unsupported : gtk_places_sidebar_add_shortcut : unsupported parameter location : no type generator for Gio.File, GFile*

// Unsupported : gtk_places_sidebar_get_location : no return generator

// Unsupported : gtk_places_sidebar_get_nth_bookmark : no return generator

// Unsupported : gtk_places_sidebar_remove_shortcut : unsupported parameter location : no type generator for Gio.File, GFile*

// Unsupported : gtk_places_sidebar_set_location : unsupported parameter location : no type generator for Gio.File, GFile*

// Blacklisted : GtkPlug

// Unsupported : gtk_popover_get_pointing_to : unsupported parameter rect : Blacklisted record : GdkRectangle

// Unsupported : gtk_popover_set_pointing_to : unsupported parameter rect : Blacklisted record : GdkRectangle

// Unsupported : gtk_print_context_get_hard_margins : unsupported parameter top : no type generator for gdouble, gdouble*

// Unsupported : gtk_print_settings_new_from_gvariant : unsupported parameter variant : Blacklisted record : GVariant

// Unsupported : gtk_print_settings_foreach : unsupported parameter func : no type generator for PrintSettingsFunc, GtkPrintSettingsFunc

// Unsupported : gtk_print_settings_get_page_ranges : unsupported parameter num_ranges : no type generator for gint, gint*

// Unsupported : gtk_print_settings_set_page_ranges : unsupported parameter page_ranges : no param type

// Unsupported : gtk_print_settings_to_gvariant : return type : Blacklisted record : GVariant

// GetEllipsize is a wrapper around the C function gtk_progress_bar_get_ellipsize.
func (recv *ProgressBar) GetEllipsize() pango.EllipsizeMode {
	retC := C.gtk_progress_bar_get_ellipsize((*C.GtkProgressBar)(recv.native))
	retGo := (pango.EllipsizeMode)(retC)

	return retGo
}

// SetEllipsize is a wrapper around the C function gtk_progress_bar_set_ellipsize.
func (recv *ProgressBar) SetEllipsize(mode pango.EllipsizeMode) {
	c_mode := (C.PangoEllipsizeMode)(mode)

	C.gtk_progress_bar_set_ellipsize((*C.GtkProgressBar)(recv.native), c_mode)

	return
}

// Unsupported : gtk_range_get_range_rect : unsupported parameter range_rect : Blacklisted record : GdkRectangle

// Unsupported : gtk_range_get_slider_range : unsupported parameter slider_start : no type generator for gint, gint*

// Unsupported : gtk_recent_chooser_dialog_new : unsupported parameter ... : varargs

// Unsupported : gtk_recent_chooser_dialog_new_for_manager : unsupported parameter ... : varargs

// Unsupported : gtk_recent_filter_add_custom : unsupported parameter func : no type generator for RecentFilterFunc, GtkRecentFilterFunc

// Unsupported : gtk_scale_get_layout_offsets : unsupported parameter x : no type generator for gint, gint*

// Unsupported : gtk_scale_button_new : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_scale_button_set_icons : unsupported parameter icons : no param type

// Unsupported : gtk_scrolled_window_get_policy : unsupported parameter hscrollbar_policy : GtkPolicyType* with indirection level of 1

// Unsupported : gtk_search_bar_handle_event : unsupported parameter event : no type generator for Gdk.Event, GdkEvent*

// Unsupported : gtk_search_entry_handle_event : unsupported parameter event : no type generator for Gdk.Event, GdkEvent*

// Blacklisted : GtkSocket

// Unsupported : gtk_spin_button_get_increments : unsupported parameter step : no type generator for gdouble, gdouble*

// Unsupported : gtk_spin_button_get_range : unsupported parameter min : no type generator for gdouble, gdouble*

// Blacklisted : GtkStackAccessible

// Unsupported : gtk_status_icon_new_from_gicon : unsupported parameter icon : no type generator for Gio.Icon, GIcon*

// Unsupported : gtk_status_icon_get_geometry : unsupported parameter screen : record with indirection level of 2

// Unsupported : gtk_status_icon_get_gicon : no return generator

// Unsupported : gtk_status_icon_set_from_gicon : unsupported parameter icon : no type generator for Gio.Icon, GIcon*

// Unsupported : gtk_style_get : unsupported parameter widget_type : no type generator for GType, GType

// Unsupported : gtk_style_get_style_property : unsupported parameter widget_type : no type generator for GType, GType

// Unsupported : gtk_style_get_valist : unsupported parameter widget_type : no type generator for GType, GType

// Unsupported : gtk_style_render_icon : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_style_context_add_provider : unsupported parameter provider : no type generator for StyleProvider, GtkStyleProvider*

// Unsupported : gtk_style_context_get : unsupported parameter ... : varargs

// Unsupported : gtk_style_context_get_style : unsupported parameter ... : varargs

// Unsupported : gtk_style_context_get_style_valist : unsupported parameter args : no type generator for va_list, va_list

// Unsupported : gtk_style_context_get_valist : unsupported parameter args : no type generator for va_list, va_list

// Unsupported : gtk_style_context_has_region : unsupported parameter flags_return : GtkRegionFlags* with indirection level of 1

// Unsupported : gtk_style_context_remove_provider : unsupported parameter provider : no type generator for StyleProvider, GtkStyleProvider*

// Unsupported : gtk_style_context_state_is_running : unsupported parameter progress : no type generator for gdouble, gdouble*

// Unsupported : gtk_style_properties_get : unsupported parameter ... : varargs

// Unsupported : gtk_style_properties_get_valist : unsupported parameter args : no type generator for va_list, va_list

// Unsupported : gtk_style_properties_set : unsupported parameter ... : varargs

// Unsupported : gtk_style_properties_set_valist : unsupported parameter args : no type generator for va_list, va_list

// Unsupported : gtk_table_get_size : unsupported parameter rows : no type generator for guint, guint*

// Backspace is a wrapper around the C function gtk_text_buffer_backspace.
func (recv *TextBuffer) Backspace(iter *TextIter, interactive bool, defaultEditable bool) bool {
	c_iter := (*C.GtkTextIter)(iter.ToC())

	c_interactive :=
		boolToGboolean(interactive)

	c_default_editable :=
		boolToGboolean(defaultEditable)

	retC := C.gtk_text_buffer_backspace((*C.GtkTextBuffer)(recv.native), c_iter, c_interactive, c_default_editable)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_text_buffer_create_tag : unsupported parameter ... : varargs

// Unsupported : gtk_text_buffer_deserialize : unsupported parameter format : Blacklisted record : GdkAtom

// Unsupported : gtk_text_buffer_deserialize_get_can_create_tags : unsupported parameter format : Blacklisted record : GdkAtom

// Unsupported : gtk_text_buffer_deserialize_set_can_create_tags : unsupported parameter format : Blacklisted record : GdkAtom

// Unsupported : gtk_text_buffer_get_deserialize_formats : unsupported parameter n_formats : no type generator for gint, gint*

// Unsupported : gtk_text_buffer_get_serialize_formats : unsupported parameter n_formats : no type generator for gint, gint*

// Unsupported : gtk_text_buffer_insert_with_tags : unsupported parameter ... : varargs

// Unsupported : gtk_text_buffer_insert_with_tags_by_name : unsupported parameter ... : varargs

// Unsupported : gtk_text_buffer_register_deserialize_format : unsupported parameter function : no type generator for TextBufferDeserializeFunc, GtkTextBufferDeserializeFunc

// Unsupported : gtk_text_buffer_register_deserialize_tagset : return type : Blacklisted record : GdkAtom

// Unsupported : gtk_text_buffer_register_serialize_format : unsupported parameter function : no type generator for TextBufferSerializeFunc, GtkTextBufferSerializeFunc

// Unsupported : gtk_text_buffer_register_serialize_tagset : return type : Blacklisted record : GdkAtom

// Unsupported : gtk_text_buffer_serialize : unsupported parameter format : Blacklisted record : GdkAtom

// Unsupported : gtk_text_buffer_unregister_deserialize_format : unsupported parameter format : Blacklisted record : GdkAtom

// Unsupported : gtk_text_buffer_unregister_serialize_format : unsupported parameter format : Blacklisted record : GdkAtom

// Unsupported : gtk_text_tag_event : unsupported parameter event : no type generator for Gdk.Event, GdkEvent*

// Unsupported : gtk_text_tag_table_foreach : unsupported parameter func : no type generator for TextTagTableForeach, GtkTextTagTableForeach

// Unsupported : gtk_text_view_buffer_to_window_coords : unsupported parameter window_x : no type generator for gint, gint*

// Unsupported : gtk_text_view_get_cursor_locations : unsupported parameter strong : Blacklisted record : GdkRectangle

// Unsupported : gtk_text_view_get_iter_at_position : unsupported parameter trailing : no type generator for gint, gint*

// Unsupported : gtk_text_view_get_iter_location : unsupported parameter location : Blacklisted record : GdkRectangle

// Unsupported : gtk_text_view_get_line_at_y : unsupported parameter line_top : no type generator for gint, gint*

// Unsupported : gtk_text_view_get_line_yrange : unsupported parameter y : no type generator for gint, gint*

// Unsupported : gtk_text_view_get_visible_rect : unsupported parameter visible_rect : Blacklisted record : GdkRectangle

// Unsupported : gtk_text_view_window_to_buffer_coords : unsupported parameter buffer_x : no type generator for gint, gint*

// Unsupported : gtk_theming_engine_get : unsupported parameter ... : varargs

// Unsupported : gtk_theming_engine_get_style : unsupported parameter ... : varargs

// Unsupported : gtk_theming_engine_get_style_valist : unsupported parameter args : no type generator for va_list, va_list

// Unsupported : gtk_theming_engine_get_valist : unsupported parameter args : no type generator for va_list, va_list

// Unsupported : gtk_theming_engine_has_region : unsupported parameter flags : GtkRegionFlags* with indirection level of 1

// Unsupported : gtk_theming_engine_state_is_running : unsupported parameter progress : no type generator for gdouble, gdouble*

// Unsupported : gtk_tool_item_get_icon_size : no return generator

// RebuildMenu is a wrapper around the C function gtk_tool_item_rebuild_menu.
func (recv *ToolItem) RebuildMenu() {
	C.gtk_tool_item_rebuild_menu((*C.GtkToolItem)(recv.native))

	return
}

// Unsupported : gtk_tool_palette_get_icon_size : no return generator

// Unsupported : gtk_tool_palette_set_icon_size : unsupported parameter icon_size : no type generator for gint, GtkIconSize

// Unsupported : gtk_tooltip_set_icon_from_gicon : unsupported parameter gicon : no type generator for Gio.Icon, GIcon*

// Unsupported : gtk_tooltip_set_icon_from_icon_name : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_tooltip_set_icon_from_stock : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_tooltip_set_tip_area : unsupported parameter rect : Blacklisted record : GdkRectangle

// Unsupported : gtk_tree_model_filter_get_model : no return generator

// Unsupported : gtk_tree_model_filter_set_modify_func : unsupported parameter types : no param type

// Unsupported : gtk_tree_model_filter_set_visible_func : unsupported parameter func : no type generator for TreeModelFilterVisibleFunc, GtkTreeModelFilterVisibleFunc

// Unsupported : gtk_tree_model_sort_get_model : no return generator

// Unsupported : gtk_tree_selection_get_select_function : no return generator

// Unsupported : gtk_tree_selection_get_selected : unsupported parameter model : no type generator for TreeModel, GtkTreeModel**

// Unsupported : gtk_tree_selection_get_selected_rows : unsupported parameter model : no type generator for TreeModel, GtkTreeModel**

// Unsupported : gtk_tree_selection_selected_foreach : unsupported parameter func : no type generator for TreeSelectionForeachFunc, GtkTreeSelectionForeachFunc

// Unsupported : gtk_tree_selection_set_select_function : unsupported parameter func : no type generator for TreeSelectionFunc, GtkTreeSelectionFunc

// Unsupported : gtk_tree_store_new : unsupported parameter ... : varargs

// Unsupported : gtk_tree_store_newv : unsupported parameter types : no param type

// Unsupported : gtk_tree_store_insert_with_values : unsupported parameter ... : varargs

// Unsupported : gtk_tree_store_insert_with_valuesv : unsupported parameter columns : no param type

// Unsupported : gtk_tree_store_reorder : unsupported parameter new_order : no param type

// Unsupported : gtk_tree_store_set : unsupported parameter ... : varargs

// Unsupported : gtk_tree_store_set_column_types : unsupported parameter types : no param type

// Unsupported : gtk_tree_store_set_valist : unsupported parameter var_args : no type generator for va_list, va_list

// Unsupported : gtk_tree_store_set_valuesv : unsupported parameter columns : no param type

// Unsupported : gtk_tree_view_new_with_model : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

// Unsupported : gtk_tree_view_convert_bin_window_to_tree_coords : unsupported parameter tx : no type generator for gint, gint*

// Unsupported : gtk_tree_view_convert_bin_window_to_widget_coords : unsupported parameter wx : no type generator for gint, gint*

// Unsupported : gtk_tree_view_convert_tree_to_bin_window_coords : unsupported parameter bx : no type generator for gint, gint*

// Unsupported : gtk_tree_view_convert_tree_to_widget_coords : unsupported parameter wx : no type generator for gint, gint*

// Unsupported : gtk_tree_view_convert_widget_to_bin_window_coords : unsupported parameter bx : no type generator for gint, gint*

// Unsupported : gtk_tree_view_convert_widget_to_tree_coords : unsupported parameter tx : no type generator for gint, gint*

// Unsupported : gtk_tree_view_enable_model_drag_dest : unsupported parameter targets : no param type

// Unsupported : gtk_tree_view_enable_model_drag_source : unsupported parameter targets : no param type

// Unsupported : gtk_tree_view_get_background_area : unsupported parameter rect : Blacklisted record : GdkRectangle

// Unsupported : gtk_tree_view_get_cell_area : unsupported parameter rect : Blacklisted record : GdkRectangle

// Unsupported : gtk_tree_view_get_cursor : unsupported parameter path : record with indirection level of 2

// Unsupported : gtk_tree_view_get_dest_row_at_pos : unsupported parameter path : record with indirection level of 2

// Unsupported : gtk_tree_view_get_drag_dest_row : unsupported parameter path : record with indirection level of 2

// GetFixedHeightMode is a wrapper around the C function gtk_tree_view_get_fixed_height_mode.
func (recv *TreeView) GetFixedHeightMode() bool {
	retC := C.gtk_tree_view_get_fixed_height_mode((*C.GtkTreeView)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetHoverExpand is a wrapper around the C function gtk_tree_view_get_hover_expand.
func (recv *TreeView) GetHoverExpand() bool {
	retC := C.gtk_tree_view_get_hover_expand((*C.GtkTreeView)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetHoverSelection is a wrapper around the C function gtk_tree_view_get_hover_selection.
func (recv *TreeView) GetHoverSelection() bool {
	retC := C.gtk_tree_view_get_hover_selection((*C.GtkTreeView)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_tree_view_get_model : no return generator

// Unsupported : gtk_tree_view_get_path_at_pos : unsupported parameter path : record with indirection level of 2

// Unsupported : gtk_tree_view_get_row_separator_func : no return generator

// Unsupported : gtk_tree_view_get_search_equal_func : no return generator

// Unsupported : gtk_tree_view_get_search_position_func : no return generator

// Unsupported : gtk_tree_view_get_tooltip_context : unsupported parameter x : no type generator for gint, gint*

// Unsupported : gtk_tree_view_get_visible_range : unsupported parameter start_path : record with indirection level of 2

// Unsupported : gtk_tree_view_get_visible_rect : unsupported parameter visible_rect : Blacklisted record : GdkRectangle

// Unsupported : gtk_tree_view_insert_column_with_attributes : unsupported parameter ... : varargs

// Unsupported : gtk_tree_view_insert_column_with_data_func : unsupported parameter func : no type generator for TreeCellDataFunc, GtkTreeCellDataFunc

// Unsupported : gtk_tree_view_is_blank_at_pos : unsupported parameter path : record with indirection level of 2

// Unsupported : gtk_tree_view_map_expanded_rows : unsupported parameter func : no type generator for TreeViewMappingFunc, GtkTreeViewMappingFunc

// Unsupported : gtk_tree_view_set_column_drag_function : unsupported parameter func : no type generator for TreeViewColumnDropFunc, GtkTreeViewColumnDropFunc

// Unsupported : gtk_tree_view_set_destroy_count_func : unsupported parameter func : no type generator for TreeDestroyCountFunc, GtkTreeDestroyCountFunc

// SetFixedHeightMode is a wrapper around the C function gtk_tree_view_set_fixed_height_mode.
func (recv *TreeView) SetFixedHeightMode(enable bool) {
	c_enable :=
		boolToGboolean(enable)

	C.gtk_tree_view_set_fixed_height_mode((*C.GtkTreeView)(recv.native), c_enable)

	return
}

// SetHoverExpand is a wrapper around the C function gtk_tree_view_set_hover_expand.
func (recv *TreeView) SetHoverExpand(expand bool) {
	c_expand :=
		boolToGboolean(expand)

	C.gtk_tree_view_set_hover_expand((*C.GtkTreeView)(recv.native), c_expand)

	return
}

// SetHoverSelection is a wrapper around the C function gtk_tree_view_set_hover_selection.
func (recv *TreeView) SetHoverSelection(hover bool) {
	c_hover :=
		boolToGboolean(hover)

	C.gtk_tree_view_set_hover_selection((*C.GtkTreeView)(recv.native), c_hover)

	return
}

// Unsupported : gtk_tree_view_set_model : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

// Unsupported : gtk_tree_view_set_row_separator_func : unsupported parameter func : no type generator for TreeViewRowSeparatorFunc, GtkTreeViewRowSeparatorFunc

// Unsupported : gtk_tree_view_set_search_equal_func : unsupported parameter search_equal_func : no type generator for TreeViewSearchEqualFunc, GtkTreeViewSearchEqualFunc

// Unsupported : gtk_tree_view_set_search_position_func : unsupported parameter func : no type generator for TreeViewSearchPositionFunc, GtkTreeViewSearchPositionFunc

// Unsupported : gtk_tree_view_column_new_with_attributes : unsupported parameter ... : varargs

// Unsupported : gtk_tree_view_column_cell_get_position : unsupported parameter x_offset : no type generator for gint, gint*

// Unsupported : gtk_tree_view_column_cell_get_size : unsupported parameter cell_area : Blacklisted record : GdkRectangle

// Unsupported : gtk_tree_view_column_cell_set_cell_data : unsupported parameter tree_model : no type generator for TreeModel, GtkTreeModel*

// Unsupported : gtk_tree_view_column_set_attributes : unsupported parameter ... : varargs

// Unsupported : gtk_tree_view_column_set_cell_data_func : unsupported parameter func : no type generator for TreeCellDataFunc, GtkTreeCellDataFunc

// Unsupported : gtk_widget_new : unsupported parameter type : no type generator for GType, GType

// Unsupported : gtk_widget_add_tick_callback : unsupported parameter callback : no type generator for TickCallback, GtkTickCallback

// Unsupported : gtk_widget_class_path : unsupported parameter path_length : no type generator for guint, guint*

// Unsupported : gtk_widget_destroyed : unsupported parameter widget_pointer : record with indirection level of 2

// Unsupported : gtk_drag_begin : unsupported parameter event : no type generator for Gdk.Event, GdkEvent*

// Unsupported : gtk_drag_begin_with_coordinates : unsupported parameter event : no type generator for Gdk.Event, GdkEvent*

// DragDestAddImageTargets is a wrapper around the C function gtk_drag_dest_add_image_targets.
func (recv *Widget) DragDestAddImageTargets() {
	C.gtk_drag_dest_add_image_targets((*C.GtkWidget)(recv.native))

	return
}

// DragDestAddTextTargets is a wrapper around the C function gtk_drag_dest_add_text_targets.
func (recv *Widget) DragDestAddTextTargets() {
	C.gtk_drag_dest_add_text_targets((*C.GtkWidget)(recv.native))

	return
}

// DragDestAddUriTargets is a wrapper around the C function gtk_drag_dest_add_uri_targets.
func (recv *Widget) DragDestAddUriTargets() {
	C.gtk_drag_dest_add_uri_targets((*C.GtkWidget)(recv.native))

	return
}

// Unsupported : gtk_drag_dest_find_target : return type : Blacklisted record : GdkAtom

// Unsupported : gtk_drag_dest_set : unsupported parameter targets : no param type

// Unsupported : gtk_drag_get_data : unsupported parameter target : Blacklisted record : GdkAtom

// DragSourceAddImageTargets is a wrapper around the C function gtk_drag_source_add_image_targets.
func (recv *Widget) DragSourceAddImageTargets() {
	C.gtk_drag_source_add_image_targets((*C.GtkWidget)(recv.native))

	return
}

// DragSourceAddTextTargets is a wrapper around the C function gtk_drag_source_add_text_targets.
func (recv *Widget) DragSourceAddTextTargets() {
	C.gtk_drag_source_add_text_targets((*C.GtkWidget)(recv.native))

	return
}

// DragSourceAddUriTargets is a wrapper around the C function gtk_drag_source_add_uri_targets.
func (recv *Widget) DragSourceAddUriTargets() {
	C.gtk_drag_source_add_uri_targets((*C.GtkWidget)(recv.native))

	return
}

// Unsupported : gtk_drag_source_set : unsupported parameter targets : no param type

// Unsupported : gtk_drag_source_set_icon_gicon : unsupported parameter icon : no type generator for Gio.Icon, GIcon*

// Unsupported : gtk_widget_event : unsupported parameter event : no type generator for Gdk.Event, GdkEvent*

// Unsupported : gtk_widget_get_action_group : no return generator

// Unsupported : gtk_widget_get_allocated_size : unsupported parameter allocation : Blacklisted record : GdkRectangle

// Unsupported : gtk_widget_get_allocation : unsupported parameter allocation : Blacklisted record : GdkRectangle

// Unsupported : gtk_widget_get_ancestor : unsupported parameter widget_type : no type generator for GType, GType

// Unsupported : gtk_widget_get_clip : unsupported parameter clip : Blacklisted record : GdkRectangle

// Unsupported : gtk_widget_get_clipboard : unsupported parameter selection : Blacklisted record : GdkAtom

// Unsupported : gtk_widget_get_pointer : unsupported parameter x : no type generator for gint, gint*

// Unsupported : gtk_widget_get_preferred_height : unsupported parameter minimum_height : no type generator for gint, gint*

// Unsupported : gtk_widget_get_preferred_height_and_baseline_for_width : unsupported parameter minimum_height : no type generator for gint, gint*

// Unsupported : gtk_widget_get_preferred_height_for_width : unsupported parameter minimum_height : no type generator for gint, gint*

// Unsupported : gtk_widget_get_preferred_width : unsupported parameter minimum_width : no type generator for gint, gint*

// Unsupported : gtk_widget_get_preferred_width_for_height : unsupported parameter minimum_width : no type generator for gint, gint*

// Unsupported : gtk_widget_get_size_request : unsupported parameter width : no type generator for gint, gint*

// Unsupported : gtk_widget_get_template_child : unsupported parameter widget_type : no type generator for GType, GType

// Unsupported : gtk_widget_insert_action_group : unsupported parameter group : no type generator for Gio.ActionGroup, GActionGroup*

// Unsupported : gtk_widget_intersect : unsupported parameter area : Blacklisted record : GdkRectangle

// Unsupported : gtk_widget_list_action_prefixes : no return type

// Unsupported : gtk_widget_path : unsupported parameter path_length : no type generator for guint, guint*

// Unsupported : gtk_widget_render_icon : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_widget_render_icon_pixbuf : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_widget_send_expose : unsupported parameter event : no type generator for Gdk.Event, GdkEvent*

// Unsupported : gtk_widget_send_focus_change : unsupported parameter event : no type generator for Gdk.Event, GdkEvent*

// Unsupported : gtk_widget_set_allocation : unsupported parameter allocation : Blacklisted record : GdkRectangle

// Unsupported : gtk_widget_set_clip : unsupported parameter clip : Blacklisted record : GdkRectangle

// Unsupported : gtk_widget_size_allocate : unsupported parameter allocation : Blacklisted record : GdkRectangle

// Unsupported : gtk_widget_size_allocate_with_baseline : unsupported parameter allocation : Blacklisted record : GdkRectangle

// Unsupported : gtk_widget_style_get : unsupported parameter ... : varargs

// Unsupported : gtk_widget_style_get_valist : unsupported parameter var_args : no type generator for va_list, va_list

// Unsupported : gtk_widget_translate_coordinates : unsupported parameter dest_x : no type generator for gint, gint*

// Unsupported : gtk_window_get_default_size : unsupported parameter width : no type generator for gint, gint*

// GetFocusOnMap is a wrapper around the C function gtk_window_get_focus_on_map.
func (recv *Window) GetFocusOnMap() bool {
	retC := C.gtk_window_get_focus_on_map((*C.GtkWindow)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetIconName is a wrapper around the C function gtk_window_get_icon_name.
func (recv *Window) GetIconName() string {
	retC := C.gtk_window_get_icon_name((*C.GtkWindow)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : gtk_window_get_position : unsupported parameter root_x : no type generator for gint, gint*

// Unsupported : gtk_window_get_resize_grip_area : unsupported parameter rect : Blacklisted record : GdkRectangle

// Unsupported : gtk_window_get_size : unsupported parameter width : no type generator for gint, gint*

// SetFocusOnMap is a wrapper around the C function gtk_window_set_focus_on_map.
func (recv *Window) SetFocusOnMap(setting bool) {
	c_setting :=
		boolToGboolean(setting)

	C.gtk_window_set_focus_on_map((*C.GtkWindow)(recv.native), c_setting)

	return
}

// SetIconName is a wrapper around the C function gtk_window_set_icon_name.
func (recv *Window) SetIconName(name string) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	C.gtk_window_set_icon_name((*C.GtkWindow)(recv.native), c_name)

	return
}
