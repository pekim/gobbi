// This is a generated file - DO NOT EDIT
// +build gtk_2.6 gtk_2.8 gtk_2.10 gtk_2.12 gtk_2.14 gtk_2.16 gtk_2.18 gtk_2.20 gtk_2.22 gtk_2.24 gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	atk "github.com/pekim/gobbi/lib/atk"
	gdk "github.com/pekim/gobbi/lib/gdk"
	gdkpixbuf "github.com/pekim/gobbi/lib/gdkpixbuf"
	glib "github.com/pekim/gobbi/lib/glib"
	gobject "github.com/pekim/gobbi/lib/gobject"
	pango "github.com/pekim/gobbi/lib/pango"
	"sync"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
/*

	void cellrenderer_editingStartedHandler(GObject *, GtkCellEditable *, gchar*, gpointer);

	static gulong CellRenderer_signal_connect_editing_started(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "editing-started", G_CALLBACK(cellrenderer_editingStartedHandler), data);
	}

*/
/*

	void clipboard_ownerChangeHandler(GObject *, GdkEventOwnerChange *, gpointer);

	static gulong Clipboard_signal_connect_owner_change(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "owner-change", G_CALLBACK(clipboard_ownerChangeHandler), data);
	}

*/
/*

	gboolean entrycompletion_insertPrefixHandler(GObject *, gchar*, gpointer);

	static gulong EntryCompletion_signal_connect_insert_prefix(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "insert-prefix", G_CALLBACK(entrycompletion_insertPrefixHandler), data);
	}

*/
import "C"

// AboutDialogNew is a wrapper around the C function gtk_about_dialog_new.
func AboutDialogNew() *AboutDialog {
	retC := C.gtk_about_dialog_new()
	gobject.ObjectNewFromC(unsafe.Pointer(retC)).Take()
	retGo := AboutDialogNewFromC(unsafe.Pointer(retC))

	return retGo
}

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

// Unsupported : gtk_about_dialog_set_artists : unsupported parameter artists :

// Unsupported : gtk_about_dialog_set_authors : unsupported parameter authors :

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

// Unsupported : gtk_about_dialog_set_documenters : unsupported parameter documenters :

// SetLicense is a wrapper around the C function gtk_about_dialog_set_license.
func (recv *AboutDialog) SetLicense(license string) {
	c_license := C.CString(license)
	defer C.free(unsafe.Pointer(c_license))

	C.gtk_about_dialog_set_license((*C.GtkAboutDialog)(recv.native), c_license)

	return
}

// SetLogo is a wrapper around the C function gtk_about_dialog_set_logo.
func (recv *AboutDialog) SetLogo(logo *gdkpixbuf.Pixbuf) {
	c_logo := (*C.GdkPixbuf)(C.NULL)
	if logo != nil {
		c_logo = (*C.GdkPixbuf)(logo.ToC())
	}

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

// GetAccelPath is a wrapper around the C function gtk_action_get_accel_path.
func (recv *Action) GetAccelPath() string {
	retC := C.gtk_action_get_accel_path((*C.GtkAction)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

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

// TranslateString is a wrapper around the C function gtk_action_group_translate_string.
func (recv *ActionGroup) TranslateString(string string) string {
	c_string := C.CString(string)
	defer C.free(unsafe.Pointer(c_string))

	retC := C.gtk_action_group_translate_string((*C.GtkActionGroup)(recv.native), c_string)
	retGo := C.GoString(retC)

	return retGo
}

// GetImage is a wrapper around the C function gtk_button_get_image.
func (recv *Button) GetImage() *Widget {
	retC := C.gtk_button_get_image((*C.GtkButton)(recv.native))
	var retGo (*Widget)
	if retC == nil {
		retGo = nil
	} else {
		retGo = WidgetNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// SetImage is a wrapper around the C function gtk_button_set_image.
func (recv *Button) SetImage(image *Widget) {
	c_image := (*C.GtkWidget)(C.NULL)
	if image != nil {
		c_image = (*C.GtkWidget)(image.ToC())
	}

	C.gtk_button_set_image((*C.GtkButton)(recv.native), c_image)

	return
}

type signalCellRendererEditingStartedDetail struct {
	callback  CellRendererSignalEditingStartedCallback
	handlerID C.gulong
}

var signalCellRendererEditingStartedId int
var signalCellRendererEditingStartedMap = make(map[int]signalCellRendererEditingStartedDetail)
var signalCellRendererEditingStartedLock sync.RWMutex

// CellRendererSignalEditingStartedCallback is a callback function for a 'editing-started' signal emitted from a CellRenderer.
type CellRendererSignalEditingStartedCallback func(editable *CellEditable, path string)

/*
ConnectEditingStarted connects the callback to the 'editing-started' signal for the CellRenderer.

The returned value represents the connection, and may be passed to DisconnectEditingStarted to remove it.
*/
func (recv *CellRenderer) ConnectEditingStarted(callback CellRendererSignalEditingStartedCallback) int {
	signalCellRendererEditingStartedLock.Lock()
	defer signalCellRendererEditingStartedLock.Unlock()

	signalCellRendererEditingStartedId++
	instance := C.gpointer(recv.native)
	handlerID := C.CellRenderer_signal_connect_editing_started(instance, C.gpointer(uintptr(signalCellRendererEditingStartedId)))

	detail := signalCellRendererEditingStartedDetail{callback, handlerID}
	signalCellRendererEditingStartedMap[signalCellRendererEditingStartedId] = detail

	return signalCellRendererEditingStartedId
}

/*
DisconnectEditingStarted disconnects a callback from the 'editing-started' signal for the CellRenderer.

The connectionID should be a value returned from a call to ConnectEditingStarted.
*/
func (recv *CellRenderer) DisconnectEditingStarted(connectionID int) {
	signalCellRendererEditingStartedLock.Lock()
	defer signalCellRendererEditingStartedLock.Unlock()

	detail, exists := signalCellRendererEditingStartedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalCellRendererEditingStartedMap, connectionID)
}

//export cellrenderer_editingStartedHandler
func cellrenderer_editingStartedHandler(_ *C.GObject, c_editable *C.GtkCellEditable, c_path *C.gchar, data C.gpointer) {
	signalCellRendererEditingStartedLock.RLock()
	defer signalCellRendererEditingStartedLock.RUnlock()

	editable := CellEditableNewFromC(unsafe.Pointer(c_editable))

	path := C.GoString(c_path)

	index := int(uintptr(data))
	callback := signalCellRendererEditingStartedMap[index].callback
	callback(editable, path)
}

// StopEditing is a wrapper around the C function gtk_cell_renderer_stop_editing.
func (recv *CellRenderer) StopEditing(canceled bool) {
	c_canceled :=
		boolToGboolean(canceled)

	C.gtk_cell_renderer_stop_editing((*C.GtkCellRenderer)(recv.native), c_canceled)

	return
}

// CellRendererComboNew is a wrapper around the C function gtk_cell_renderer_combo_new.
func CellRendererComboNew() *CellRendererCombo {
	retC := C.gtk_cell_renderer_combo_new()
	gobject.ObjectNewFromC(unsafe.Pointer(retC)).Take()
	retGo := CellRendererComboNewFromC(unsafe.Pointer(retC))

	return retGo
}

// CellRendererProgressNew is a wrapper around the C function gtk_cell_renderer_progress_new.
func CellRendererProgressNew() *CellRendererProgress {
	retC := C.gtk_cell_renderer_progress_new()
	gobject.ObjectNewFromC(unsafe.Pointer(retC)).Take()
	retGo := CellRendererProgressNewFromC(unsafe.Pointer(retC))

	return retGo
}

// CellViewNew is a wrapper around the C function gtk_cell_view_new.
func CellViewNew() *CellView {
	retC := C.gtk_cell_view_new()
	gobject.ObjectNewFromC(unsafe.Pointer(retC)).Take()
	retGo := CellViewNewFromC(unsafe.Pointer(retC))

	return retGo
}

// CellViewNewWithContext is a wrapper around the C function gtk_cell_view_new_with_context.
func CellViewNewWithContext(area *CellArea, context *CellAreaContext) *CellView {
	c_area := (*C.GtkCellArea)(C.NULL)
	if area != nil {
		c_area = (*C.GtkCellArea)(area.ToC())
	}

	c_context := (*C.GtkCellAreaContext)(C.NULL)
	if context != nil {
		c_context = (*C.GtkCellAreaContext)(context.ToC())
	}

	retC := C.gtk_cell_view_new_with_context(c_area, c_context)
	gobject.ObjectNewFromC(unsafe.Pointer(retC)).Take()
	retGo := CellViewNewFromC(unsafe.Pointer(retC))

	return retGo
}

// CellViewNewWithMarkup is a wrapper around the C function gtk_cell_view_new_with_markup.
func CellViewNewWithMarkup(markup string) *CellView {
	c_markup := C.CString(markup)
	defer C.free(unsafe.Pointer(c_markup))

	retC := C.gtk_cell_view_new_with_markup(c_markup)
	gobject.ObjectNewFromC(unsafe.Pointer(retC)).Take()
	retGo := CellViewNewFromC(unsafe.Pointer(retC))

	return retGo
}

// CellViewNewWithPixbuf is a wrapper around the C function gtk_cell_view_new_with_pixbuf.
func CellViewNewWithPixbuf(pixbuf *gdkpixbuf.Pixbuf) *CellView {
	c_pixbuf := (*C.GdkPixbuf)(C.NULL)
	if pixbuf != nil {
		c_pixbuf = (*C.GdkPixbuf)(pixbuf.ToC())
	}

	retC := C.gtk_cell_view_new_with_pixbuf(c_pixbuf)
	gobject.ObjectNewFromC(unsafe.Pointer(retC)).Take()
	retGo := CellViewNewFromC(unsafe.Pointer(retC))

	return retGo
}

// CellViewNewWithText is a wrapper around the C function gtk_cell_view_new_with_text.
func CellViewNewWithText(text string) *CellView {
	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))

	retC := C.gtk_cell_view_new_with_text(c_text)
	gobject.ObjectNewFromC(unsafe.Pointer(retC)).Take()
	retGo := CellViewNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetDisplayedRow is a wrapper around the C function gtk_cell_view_get_displayed_row.
func (recv *CellView) GetDisplayedRow() *TreePath {
	retC := C.gtk_cell_view_get_displayed_row((*C.GtkCellView)(recv.native))
	var retGo (*TreePath)
	if retC == nil {
		retGo = nil
	} else {
		retGo = TreePathNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetSizeOfRow is a wrapper around the C function gtk_cell_view_get_size_of_row.
func (recv *CellView) GetSizeOfRow(path *TreePath) (bool, *Requisition) {
	c_path := (*C.GtkTreePath)(C.NULL)
	if path != nil {
		c_path = (*C.GtkTreePath)(path.ToC())
	}

	var c_requisition C.GtkRequisition

	retC := C.gtk_cell_view_get_size_of_row((*C.GtkCellView)(recv.native), c_path, &c_requisition)
	retGo := retC == C.TRUE

	requisition := RequisitionNewFromC(unsafe.Pointer(&c_requisition))

	return retGo, requisition
}

// SetBackgroundColor is a wrapper around the C function gtk_cell_view_set_background_color.
func (recv *CellView) SetBackgroundColor(color *gdk.Color) {
	c_color := (*C.GdkColor)(C.NULL)
	if color != nil {
		c_color = (*C.GdkColor)(color.ToC())
	}

	C.gtk_cell_view_set_background_color((*C.GtkCellView)(recv.native), c_color)

	return
}

// SetDisplayedRow is a wrapper around the C function gtk_cell_view_set_displayed_row.
func (recv *CellView) SetDisplayedRow(path *TreePath) {
	c_path := (*C.GtkTreePath)(C.NULL)
	if path != nil {
		c_path = (*C.GtkTreePath)(path.ToC())
	}

	C.gtk_cell_view_set_displayed_row((*C.GtkCellView)(recv.native), c_path)

	return
}

// SetModel is a wrapper around the C function gtk_cell_view_set_model.
func (recv *CellView) SetModel(model *TreeModel) {
	c_model := (*C.GtkTreeModel)(model.ToC())

	C.gtk_cell_view_set_model((*C.GtkCellView)(recv.native), c_model)

	return
}

type signalClipboardOwnerChangeDetail struct {
	callback  ClipboardSignalOwnerChangeCallback
	handlerID C.gulong
}

var signalClipboardOwnerChangeId int
var signalClipboardOwnerChangeMap = make(map[int]signalClipboardOwnerChangeDetail)
var signalClipboardOwnerChangeLock sync.RWMutex

// ClipboardSignalOwnerChangeCallback is a callback function for a 'owner-change' signal emitted from a Clipboard.
type ClipboardSignalOwnerChangeCallback func(event *gdk.EventOwnerChange)

/*
ConnectOwnerChange connects the callback to the 'owner-change' signal for the Clipboard.

The returned value represents the connection, and may be passed to DisconnectOwnerChange to remove it.
*/
func (recv *Clipboard) ConnectOwnerChange(callback ClipboardSignalOwnerChangeCallback) int {
	signalClipboardOwnerChangeLock.Lock()
	defer signalClipboardOwnerChangeLock.Unlock()

	signalClipboardOwnerChangeId++
	instance := C.gpointer(recv.native)
	handlerID := C.Clipboard_signal_connect_owner_change(instance, C.gpointer(uintptr(signalClipboardOwnerChangeId)))

	detail := signalClipboardOwnerChangeDetail{callback, handlerID}
	signalClipboardOwnerChangeMap[signalClipboardOwnerChangeId] = detail

	return signalClipboardOwnerChangeId
}

/*
DisconnectOwnerChange disconnects a callback from the 'owner-change' signal for the Clipboard.

The connectionID should be a value returned from a call to ConnectOwnerChange.
*/
func (recv *Clipboard) DisconnectOwnerChange(connectionID int) {
	signalClipboardOwnerChangeLock.Lock()
	defer signalClipboardOwnerChangeLock.Unlock()

	detail, exists := signalClipboardOwnerChangeMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalClipboardOwnerChangeMap, connectionID)
}

//export clipboard_ownerChangeHandler
func clipboard_ownerChangeHandler(_ *C.GObject, c_event *C.GdkEventOwnerChange, data C.gpointer) {
	signalClipboardOwnerChangeLock.RLock()
	defer signalClipboardOwnerChangeLock.RUnlock()

	event := gdk.EventOwnerChangeNewFromC(unsafe.Pointer(c_event))

	index := int(uintptr(data))
	callback := signalClipboardOwnerChangeMap[index].callback
	callback(event)
}

// Unsupported : gtk_clipboard_request_image : unsupported parameter callback : no type generator for ClipboardImageReceivedFunc (GtkClipboardImageReceivedFunc) for param callback

// Unsupported : gtk_clipboard_set_can_store : unsupported parameter targets :

// SetImage is a wrapper around the C function gtk_clipboard_set_image.
func (recv *Clipboard) SetImage(pixbuf *gdkpixbuf.Pixbuf) {
	c_pixbuf := (*C.GdkPixbuf)(C.NULL)
	if pixbuf != nil {
		c_pixbuf = (*C.GdkPixbuf)(pixbuf.ToC())
	}

	C.gtk_clipboard_set_image((*C.GtkClipboard)(recv.native), c_pixbuf)

	return
}

// Store is a wrapper around the C function gtk_clipboard_store.
func (recv *Clipboard) Store() {
	C.gtk_clipboard_store((*C.GtkClipboard)(recv.native))

	return
}

// WaitForImage is a wrapper around the C function gtk_clipboard_wait_for_image.
func (recv *Clipboard) WaitForImage() *gdkpixbuf.Pixbuf {
	retC := C.gtk_clipboard_wait_for_image((*C.GtkClipboard)(recv.native))
	var retGo (*gdkpixbuf.Pixbuf)
	if retC == nil {
		retGo = nil
	} else {
		retGo = gdkpixbuf.PixbufNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// WaitIsImageAvailable is a wrapper around the C function gtk_clipboard_wait_is_image_available.
func (recv *Clipboard) WaitIsImageAvailable() bool {
	retC := C.gtk_clipboard_wait_is_image_available((*C.GtkClipboard)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_clipboard_wait_is_target_available : unsupported parameter target : Blacklisted record : GdkAtom

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

// Unsupported : gtk_combo_box_set_row_separator_func : unsupported parameter func : no type generator for TreeViewRowSeparatorFunc (GtkTreeViewRowSeparatorFunc) for param func

// Unsupported : gtk_dialog_set_alternative_button_order : unsupported parameter ... : varargs

// SetAlternativeButtonOrderFromArray is a wrapper around the C function gtk_dialog_set_alternative_button_order_from_array.
func (recv *Dialog) SetAlternativeButtonOrderFromArray(newOrder []int32) {
	c_n_params := (C.gint)(len(newOrder))

	c_new_order := &newOrder[0]

	C.gtk_dialog_set_alternative_button_order_from_array((*C.GtkDialog)(recv.native), c_n_params, (*C.gint)(unsafe.Pointer(c_new_order)))

	return
}

type signalEntryCompletionInsertPrefixDetail struct {
	callback  EntryCompletionSignalInsertPrefixCallback
	handlerID C.gulong
}

var signalEntryCompletionInsertPrefixId int
var signalEntryCompletionInsertPrefixMap = make(map[int]signalEntryCompletionInsertPrefixDetail)
var signalEntryCompletionInsertPrefixLock sync.RWMutex

// EntryCompletionSignalInsertPrefixCallback is a callback function for a 'insert-prefix' signal emitted from a EntryCompletion.
type EntryCompletionSignalInsertPrefixCallback func(prefix string) bool

/*
ConnectInsertPrefix connects the callback to the 'insert-prefix' signal for the EntryCompletion.

The returned value represents the connection, and may be passed to DisconnectInsertPrefix to remove it.
*/
func (recv *EntryCompletion) ConnectInsertPrefix(callback EntryCompletionSignalInsertPrefixCallback) int {
	signalEntryCompletionInsertPrefixLock.Lock()
	defer signalEntryCompletionInsertPrefixLock.Unlock()

	signalEntryCompletionInsertPrefixId++
	instance := C.gpointer(recv.native)
	handlerID := C.EntryCompletion_signal_connect_insert_prefix(instance, C.gpointer(uintptr(signalEntryCompletionInsertPrefixId)))

	detail := signalEntryCompletionInsertPrefixDetail{callback, handlerID}
	signalEntryCompletionInsertPrefixMap[signalEntryCompletionInsertPrefixId] = detail

	return signalEntryCompletionInsertPrefixId
}

/*
DisconnectInsertPrefix disconnects a callback from the 'insert-prefix' signal for the EntryCompletion.

The connectionID should be a value returned from a call to ConnectInsertPrefix.
*/
func (recv *EntryCompletion) DisconnectInsertPrefix(connectionID int) {
	signalEntryCompletionInsertPrefixLock.Lock()
	defer signalEntryCompletionInsertPrefixLock.Unlock()

	detail, exists := signalEntryCompletionInsertPrefixMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalEntryCompletionInsertPrefixMap, connectionID)
}

//export entrycompletion_insertPrefixHandler
func entrycompletion_insertPrefixHandler(_ *C.GObject, c_prefix *C.gchar, data C.gpointer) C.gboolean {
	signalEntryCompletionInsertPrefixLock.RLock()
	defer signalEntryCompletionInsertPrefixLock.RUnlock()

	prefix := C.GoString(c_prefix)

	index := int(uintptr(data))
	callback := signalEntryCompletionInsertPrefixMap[index].callback
	retGo := callback(prefix)
	retC :=
		boolToGboolean(retGo)
	return retC
}

// GetInlineCompletion is a wrapper around the C function gtk_entry_completion_get_inline_completion.
func (recv *EntryCompletion) GetInlineCompletion() bool {
	retC := C.gtk_entry_completion_get_inline_completion((*C.GtkEntryCompletion)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

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

// SetPopupCompletion is a wrapper around the C function gtk_entry_completion_set_popup_completion.
func (recv *EntryCompletion) SetPopupCompletion(popupCompletion bool) {
	c_popup_completion :=
		boolToGboolean(popupCompletion)

	C.gtk_entry_completion_set_popup_completion((*C.GtkEntryCompletion)(recv.native), c_popup_completion)

	return
}

// FileChooserButtonNew is a wrapper around the C function gtk_file_chooser_button_new.
func FileChooserButtonNew(title string, action FileChooserAction) *FileChooserButton {
	c_title := C.CString(title)
	defer C.free(unsafe.Pointer(c_title))

	c_action := (C.GtkFileChooserAction)(action)

	retC := C.gtk_file_chooser_button_new(c_title, c_action)
	gobject.ObjectNewFromC(unsafe.Pointer(retC)).Take()
	retGo := FileChooserButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// FileChooserButtonNewWithDialog is a wrapper around the C function gtk_file_chooser_button_new_with_dialog.
func FileChooserButtonNewWithDialog(dialog *Dialog) *FileChooserButton {
	c_dialog := (*C.GtkWidget)(C.NULL)
	if dialog != nil {
		c_dialog = (*C.GtkWidget)(dialog.ToC())
	}

	retC := C.gtk_file_chooser_button_new_with_dialog(c_dialog)
	gobject.ObjectNewFromC(unsafe.Pointer(retC)).Take()
	retGo := FileChooserButtonNewFromC(unsafe.Pointer(retC))

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

// AddPixbufFormats is a wrapper around the C function gtk_file_filter_add_pixbuf_formats.
func (recv *FileFilter) AddPixbufFormats() {
	C.gtk_file_filter_add_pixbuf_formats((*C.GtkFileFilter)(recv.native))

	return
}

// Unsupported : gtk_icon_theme_get_icon_sizes : no return type

// IconViewNew is a wrapper around the C function gtk_icon_view_new.
func IconViewNew() *IconView {
	retC := C.gtk_icon_view_new()
	gobject.ObjectNewFromC(unsafe.Pointer(retC)).Take()
	retGo := IconViewNewFromC(unsafe.Pointer(retC))

	return retGo
}

// IconViewNewWithModel is a wrapper around the C function gtk_icon_view_new_with_model.
func IconViewNewWithModel(model *TreeModel) *IconView {
	c_model := (*C.GtkTreeModel)(model.ToC())

	retC := C.gtk_icon_view_new_with_model(c_model)
	gobject.ObjectNewFromC(unsafe.Pointer(retC)).Take()
	retGo := IconViewNewFromC(unsafe.Pointer(retC))

	return retGo
}

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

// GetModel is a wrapper around the C function gtk_icon_view_get_model.
func (recv *IconView) GetModel() *TreeModel {
	retC := C.gtk_icon_view_get_model((*C.GtkIconView)(recv.native))
	var retGo (*TreeModel)
	if retC == nil {
		retGo = nil
	} else {
		retGo = TreeModelNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetPathAtPos is a wrapper around the C function gtk_icon_view_get_path_at_pos.
func (recv *IconView) GetPathAtPos(x int32, y int32) *TreePath {
	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	retC := C.gtk_icon_view_get_path_at_pos((*C.GtkIconView)(recv.native), c_x, c_y)
	var retGo (*TreePath)
	if retC == nil {
		retGo = nil
	} else {
		retGo = TreePathNewFromC(unsafe.Pointer(retC))
	}

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

// ItemActivated is a wrapper around the C function gtk_icon_view_item_activated.
func (recv *IconView) ItemActivated(path *TreePath) {
	c_path := (*C.GtkTreePath)(C.NULL)
	if path != nil {
		c_path = (*C.GtkTreePath)(path.ToC())
	}

	C.gtk_icon_view_item_activated((*C.GtkIconView)(recv.native), c_path)

	return
}

// PathIsSelected is a wrapper around the C function gtk_icon_view_path_is_selected.
func (recv *IconView) PathIsSelected(path *TreePath) bool {
	c_path := (*C.GtkTreePath)(C.NULL)
	if path != nil {
		c_path = (*C.GtkTreePath)(path.ToC())
	}

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
	c_path := (*C.GtkTreePath)(C.NULL)
	if path != nil {
		c_path = (*C.GtkTreePath)(path.ToC())
	}

	C.gtk_icon_view_select_path((*C.GtkIconView)(recv.native), c_path)

	return
}

// Unsupported : gtk_icon_view_selected_foreach : unsupported parameter func : no type generator for IconViewForeachFunc (GtkIconViewForeachFunc) for param func

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

// SetModel is a wrapper around the C function gtk_icon_view_set_model.
func (recv *IconView) SetModel(model *TreeModel) {
	c_model := (*C.GtkTreeModel)(model.ToC())

	C.gtk_icon_view_set_model((*C.GtkIconView)(recv.native), c_model)

	return
}

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
	c_path := (*C.GtkTreePath)(C.NULL)
	if path != nil {
		c_path = (*C.GtkTreePath)(path.ToC())
	}

	C.gtk_icon_view_unselect_path((*C.GtkIconView)(recv.native), c_path)

	return
}

// ImageNewFromIconName is a wrapper around the C function gtk_image_new_from_icon_name.
func ImageNewFromIconName(iconName string, size IconSize) *Image {
	c_icon_name := C.CString(iconName)
	defer C.free(unsafe.Pointer(c_icon_name))

	c_size := (C.GtkIconSize)(size)

	retC := C.gtk_image_new_from_icon_name(c_icon_name, c_size)
	gobject.ObjectNewFromC(unsafe.Pointer(retC)).Take()
	retGo := ImageNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_image_get_icon_name : unsupported parameter size : no type generator for gint (GtkIconSize*) for param size

// GetPixelSize is a wrapper around the C function gtk_image_get_pixel_size.
func (recv *Image) GetPixelSize() int32 {
	retC := C.gtk_image_get_pixel_size((*C.GtkImage)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// SetFromIconName is a wrapper around the C function gtk_image_set_from_icon_name.
func (recv *Image) SetFromIconName(iconName string, size IconSize) {
	c_icon_name := C.CString(iconName)
	defer C.free(unsafe.Pointer(c_icon_name))

	c_size := (C.GtkIconSize)(size)

	C.gtk_image_set_from_icon_name((*C.GtkImage)(recv.native), c_icon_name, c_size)

	return
}

// SetPixelSize is a wrapper around the C function gtk_image_set_pixel_size.
func (recv *Image) SetPixelSize(pixelSize int32) {
	c_pixel_size := (C.gint)(pixelSize)

	C.gtk_image_set_pixel_size((*C.GtkImage)(recv.native), c_pixel_size)

	return
}

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

// GetMaxWidthChars is a wrapper around the C function gtk_label_get_max_width_chars.
func (recv *Label) GetMaxWidthChars() int32 {
	retC := C.gtk_label_get_max_width_chars((*C.GtkLabel)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

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

// Unsupported : gtk_list_store_insert_with_values : unsupported parameter ... : varargs

// Unsupported : gtk_list_store_insert_with_valuesv : unsupported parameter values :

// MenuGetForAttachWidget is a wrapper around the C function gtk_menu_get_for_attach_widget.
func MenuGetForAttachWidget(widget *Widget) *glib.List {
	c_widget := (*C.GtkWidget)(C.NULL)
	if widget != nil {
		c_widget = (*C.GtkWidget)(widget.ToC())
	}

	retC := C.gtk_menu_get_for_attach_widget(c_widget)
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// MenuToolButtonNew is a wrapper around the C function gtk_menu_tool_button_new.
func MenuToolButtonNew(iconWidget *Widget, label string) *MenuToolButton {
	c_icon_widget := (*C.GtkWidget)(C.NULL)
	if iconWidget != nil {
		c_icon_widget = (*C.GtkWidget)(iconWidget.ToC())
	}

	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	retC := C.gtk_menu_tool_button_new(c_icon_widget, c_label)
	gobject.ObjectNewFromC(unsafe.Pointer(retC)).Take()
	retGo := MenuToolButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// MenuToolButtonNewFromStock is a wrapper around the C function gtk_menu_tool_button_new_from_stock.
func MenuToolButtonNewFromStock(stockId string) *MenuToolButton {
	c_stock_id := C.CString(stockId)
	defer C.free(unsafe.Pointer(c_stock_id))

	retC := C.gtk_menu_tool_button_new_from_stock(c_stock_id)
	gobject.ObjectNewFromC(unsafe.Pointer(retC)).Take()
	retGo := MenuToolButtonNewFromC(unsafe.Pointer(retC))

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
	c_menu := (*C.GtkWidget)(C.NULL)
	if menu != nil {
		c_menu = (*C.GtkWidget)(menu.ToC())
	}

	C.gtk_menu_tool_button_set_menu((*C.GtkMenuToolButton)(recv.native), c_menu)

	return
}

// Unsupported : gtk_message_dialog_format_secondary_markup : unsupported parameter ... : varargs

// Unsupported : gtk_message_dialog_format_secondary_text : unsupported parameter ... : varargs

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

// Unsupported signal 'change-value' for Range : unsupported parameter scroll : type ScrollType :

// Backspace is a wrapper around the C function gtk_text_buffer_backspace.
func (recv *TextBuffer) Backspace(iter *TextIter, interactive bool, defaultEditable bool) bool {
	c_iter := (*C.GtkTextIter)(C.NULL)
	if iter != nil {
		c_iter = (*C.GtkTextIter)(iter.ToC())
	}

	c_interactive :=
		boolToGboolean(interactive)

	c_default_editable :=
		boolToGboolean(defaultEditable)

	retC := C.gtk_text_buffer_backspace((*C.GtkTextBuffer)(recv.native), c_iter, c_interactive, c_default_editable)
	retGo := retC == C.TRUE

	return retGo
}

// GetIterAtPosition is a wrapper around the C function gtk_text_view_get_iter_at_position.
func (recv *TextView) GetIterAtPosition(x int32, y int32) (*TextIter, int32) {
	var c_iter C.GtkTextIter

	var c_trailing C.gint

	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	C.gtk_text_view_get_iter_at_position((*C.GtkTextView)(recv.native), &c_iter, &c_trailing, c_x, c_y)

	iter := TextIterNewFromC(unsafe.Pointer(&c_iter))

	trailing := (int32)(c_trailing)

	return iter, trailing
}

// RebuildMenu is a wrapper around the C function gtk_tool_item_rebuild_menu.
func (recv *ToolItem) RebuildMenu() {
	C.gtk_tool_item_rebuild_menu((*C.GtkToolItem)(recv.native))

	return
}

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

// Unsupported : gtk_tree_view_get_row_separator_func : no return generator

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

// Unsupported : gtk_tree_view_set_row_separator_func : unsupported parameter func : no type generator for TreeViewRowSeparatorFunc (GtkTreeViewRowSeparatorFunc) for param func

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

// WindowSetDefaultIconName is a wrapper around the C function gtk_window_set_default_icon_name.
func WindowSetDefaultIconName(name string) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	C.gtk_window_set_default_icon_name(c_name)

	return
}

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
