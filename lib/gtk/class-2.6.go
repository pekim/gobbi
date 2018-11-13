// This is a generated file - DO NOT EDIT
// +build gtk_2.6 gtk_2.8 gtk_2.10 gtk_2.12 gtk_2.14 gtk_2.16 gtk_2.18 gtk_2.20 gtk_2.22 gtk_2.24 gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	atk "github.com/pekim/gobbi/lib/atk"
	gdk "github.com/pekim/gobbi/lib/gdk"
	gdkpixbuf "github.com/pekim/gobbi/lib/gdkpixbuf"
	glib "github.com/pekim/gobbi/lib/glib"
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

	void clipboard_ownerChangeHandler(GObject *, GdkEventOwnerChange *, gpointer);

	static gulong Clipboard_signal_connect_owner_change(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "owner-change", G_CALLBACK(clipboard_ownerChangeHandler), data);
	}

*/
import "C"

// Creates a new #GtkAboutDialog.
/*

C function

gtk_about_dialog_new
*/
func AboutDialogNew() *AboutDialog {
	retC := C.gtk_about_dialog_new()
	retGo := AboutDialogNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_about_dialog_get_artists : no return type

// Unsupported : gtk_about_dialog_get_authors : no return type

// Returns the comments string.
/*

C function

gtk_about_dialog_get_comments
*/
func (recv *AboutDialog) GetComments() string {
	retC := C.gtk_about_dialog_get_comments((*C.GtkAboutDialog)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Returns the copyright string.
/*

C function

gtk_about_dialog_get_copyright
*/
func (recv *AboutDialog) GetCopyright() string {
	retC := C.gtk_about_dialog_get_copyright((*C.GtkAboutDialog)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : gtk_about_dialog_get_documenters : no return type

// Returns the license information.
/*

C function

gtk_about_dialog_get_license
*/
func (recv *AboutDialog) GetLicense() string {
	retC := C.gtk_about_dialog_get_license((*C.GtkAboutDialog)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Returns the pixbuf displayed as logo in the about dialog.
/*

C function

gtk_about_dialog_get_logo
*/
func (recv *AboutDialog) GetLogo() *gdkpixbuf.Pixbuf {
	retC := C.gtk_about_dialog_get_logo((*C.GtkAboutDialog)(recv.native))
	retGo := gdkpixbuf.PixbufNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Returns the icon name displayed as logo in the about dialog.
/*

C function

gtk_about_dialog_get_logo_icon_name
*/
func (recv *AboutDialog) GetLogoIconName() string {
	retC := C.gtk_about_dialog_get_logo_icon_name((*C.GtkAboutDialog)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Returns the translator credits string which is displayed
// in the translators tab of the secondary credits dialog.
/*

C function

gtk_about_dialog_get_translator_credits
*/
func (recv *AboutDialog) GetTranslatorCredits() string {
	retC := C.gtk_about_dialog_get_translator_credits((*C.GtkAboutDialog)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Returns the version string.
/*

C function

gtk_about_dialog_get_version
*/
func (recv *AboutDialog) GetVersion() string {
	retC := C.gtk_about_dialog_get_version((*C.GtkAboutDialog)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Returns the website URL.
/*

C function

gtk_about_dialog_get_website
*/
func (recv *AboutDialog) GetWebsite() string {
	retC := C.gtk_about_dialog_get_website((*C.GtkAboutDialog)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Returns the label used for the website link.
/*

C function

gtk_about_dialog_get_website_label
*/
func (recv *AboutDialog) GetWebsiteLabel() string {
	retC := C.gtk_about_dialog_get_website_label((*C.GtkAboutDialog)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : gtk_about_dialog_set_artists : unsupported parameter artists :

// Unsupported : gtk_about_dialog_set_authors : unsupported parameter authors :

// Sets the comments string to display in the about dialog.
// This should be a short string of one or two lines.
/*

C function

gtk_about_dialog_set_comments
*/
func (recv *AboutDialog) SetComments(comments string) {
	c_comments := C.CString(comments)
	defer C.free(unsafe.Pointer(c_comments))

	C.gtk_about_dialog_set_comments((*C.GtkAboutDialog)(recv.native), c_comments)

	return
}

// Sets the copyright string to display in the about dialog.
// This should be a short string of one or two lines.
/*

C function

gtk_about_dialog_set_copyright
*/
func (recv *AboutDialog) SetCopyright(copyright string) {
	c_copyright := C.CString(copyright)
	defer C.free(unsafe.Pointer(c_copyright))

	C.gtk_about_dialog_set_copyright((*C.GtkAboutDialog)(recv.native), c_copyright)

	return
}

// Unsupported : gtk_about_dialog_set_documenters : unsupported parameter documenters :

// Sets the license information to be displayed in the secondary
// license dialog. If @license is %NULL, the license button is
// hidden.
/*

C function

gtk_about_dialog_set_license
*/
func (recv *AboutDialog) SetLicense(license string) {
	c_license := C.CString(license)
	defer C.free(unsafe.Pointer(c_license))

	C.gtk_about_dialog_set_license((*C.GtkAboutDialog)(recv.native), c_license)

	return
}

// Sets the pixbuf to be displayed as logo in the about dialog.
// If it is %NULL, the default window icon set with
// gtk_window_set_default_icon() will be used.
/*

C function

gtk_about_dialog_set_logo
*/
func (recv *AboutDialog) SetLogo(logo *gdkpixbuf.Pixbuf) {
	c_logo := (*C.GdkPixbuf)(C.NULL)
	if logo != nil {
		c_logo = (*C.GdkPixbuf)(logo.ToC())
	}

	C.gtk_about_dialog_set_logo((*C.GtkAboutDialog)(recv.native), c_logo)

	return
}

// Sets the pixbuf to be displayed as logo in the about dialog.
// If it is %NULL, the default window icon set with
// gtk_window_set_default_icon() will be used.
/*

C function

gtk_about_dialog_set_logo_icon_name
*/
func (recv *AboutDialog) SetLogoIconName(iconName string) {
	c_icon_name := C.CString(iconName)
	defer C.free(unsafe.Pointer(c_icon_name))

	C.gtk_about_dialog_set_logo_icon_name((*C.GtkAboutDialog)(recv.native), c_icon_name)

	return
}

// Sets the translator credits string which is displayed in
// the translators tab of the secondary credits dialog.
//
// The intended use for this string is to display the translator
// of the language which is currently used in the user interface.
// Using gettext(), a simple way to achieve that is to mark the
// string for translation:
// |[<!-- language="C" -->
// GtkWidget *about = gtk_about_dialog_new ();
// gtk_about_dialog_set_translator_credits (GTK_ABOUT_DIALOG (about),
// _("translator-credits"));
// ]|
// It is a good idea to use the customary msgid “translator-credits” for this
// purpose, since translators will already know the purpose of that msgid, and
// since #GtkAboutDialog will detect if “translator-credits” is untranslated
// and hide the tab.
/*

C function

gtk_about_dialog_set_translator_credits
*/
func (recv *AboutDialog) SetTranslatorCredits(translatorCredits string) {
	c_translator_credits := C.CString(translatorCredits)
	defer C.free(unsafe.Pointer(c_translator_credits))

	C.gtk_about_dialog_set_translator_credits((*C.GtkAboutDialog)(recv.native), c_translator_credits)

	return
}

// Sets the version string to display in the about dialog.
/*

C function

gtk_about_dialog_set_version
*/
func (recv *AboutDialog) SetVersion(version string) {
	c_version := C.CString(version)
	defer C.free(unsafe.Pointer(c_version))

	C.gtk_about_dialog_set_version((*C.GtkAboutDialog)(recv.native), c_version)

	return
}

// Sets the URL to use for the website link.
/*

C function

gtk_about_dialog_set_website
*/
func (recv *AboutDialog) SetWebsite(website string) {
	c_website := C.CString(website)
	defer C.free(unsafe.Pointer(c_website))

	C.gtk_about_dialog_set_website((*C.GtkAboutDialog)(recv.native), c_website)

	return
}

// Sets the label to be used for the website link.
/*

C function

gtk_about_dialog_set_website_label
*/
func (recv *AboutDialog) SetWebsiteLabel(websiteLabel string) {
	c_website_label := C.CString(websiteLabel)
	defer C.free(unsafe.Pointer(c_website_label))

	C.gtk_about_dialog_set_website_label((*C.GtkAboutDialog)(recv.native), c_website_label)

	return
}

// Returns the accel path for this action.
/*

C function

gtk_action_get_accel_path
*/
func (recv *Action) GetAccelPath() string {
	retC := C.gtk_action_get_accel_path((*C.GtkAction)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Sets the :sensitive property of the action to @sensitive. Note that
// this doesn’t necessarily mean effective sensitivity. See
// gtk_action_is_sensitive()
// for that.
/*

C function

gtk_action_set_sensitive
*/
func (recv *Action) SetSensitive(sensitive bool) {
	c_sensitive :=
		boolToGboolean(sensitive)

	C.gtk_action_set_sensitive((*C.GtkAction)(recv.native), c_sensitive)

	return
}

// Sets the :visible property of the action to @visible. Note that
// this doesn’t necessarily mean effective visibility. See
// gtk_action_is_visible()
// for that.
/*

C function

gtk_action_set_visible
*/
func (recv *Action) SetVisible(visible bool) {
	c_visible :=
		boolToGboolean(visible)

	C.gtk_action_set_visible((*C.GtkAction)(recv.native), c_visible)

	return
}

// Translates a string using the function set with
// gtk_action_group_set_translate_func(). This
// is mainly intended for language bindings.
/*

C function

gtk_action_group_translate_string
*/
func (recv *ActionGroup) TranslateString(string string) string {
	c_string := C.CString(string)
	defer C.free(unsafe.Pointer(c_string))

	retC := C.gtk_action_group_translate_string((*C.GtkActionGroup)(recv.native), c_string)
	retGo := C.GoString(retC)

	return retGo
}

// Gets the widget that is currenty set as the image of @button.
// This may have been explicitly set by gtk_button_set_image()
// or constructed by gtk_button_new_from_stock().
/*

C function

gtk_button_get_image
*/
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

// Set the image of @button to the given widget. The image will be
// displayed if the label text is %NULL or if
// #GtkButton:always-show-image is %TRUE. You don’t have to call
// gtk_widget_show() on @image yourself.
/*

C function

gtk_button_set_image
*/
func (recv *Button) SetImage(image *Widget) {
	c_image := (*C.GtkWidget)(C.NULL)
	if image != nil {
		c_image = (*C.GtkWidget)(image.ToC())
	}

	C.gtk_button_set_image((*C.GtkButton)(recv.native), c_image)

	return
}

// Unsupported signal 'editing-started' for CellRenderer : unsupported parameter path : type utf8 :

// Informs the cell renderer that the editing is stopped.
// If @canceled is %TRUE, the cell renderer will emit the
// #GtkCellRenderer::editing-canceled signal.
//
// This function should be called by cell renderer implementations
// in response to the #GtkCellEditable::editing-done signal of
// #GtkCellEditable.
/*

C function

gtk_cell_renderer_stop_editing
*/
func (recv *CellRenderer) StopEditing(canceled bool) {
	c_canceled :=
		boolToGboolean(canceled)

	C.gtk_cell_renderer_stop_editing((*C.GtkCellRenderer)(recv.native), c_canceled)

	return
}

// Creates a new #GtkCellRendererCombo.
// Adjust how text is drawn using object properties.
// Object properties can be set globally (with g_object_set()).
// Also, with #GtkTreeViewColumn, you can bind a property to a value
// in a #GtkTreeModel. For example, you can bind the “text” property
// on the cell renderer to a string value in the model, thus rendering
// a different string in each row of the #GtkTreeView.
/*

C function

gtk_cell_renderer_combo_new
*/
func CellRendererComboNew() *CellRendererCombo {
	retC := C.gtk_cell_renderer_combo_new()
	retGo := CellRendererComboNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Creates a new #GtkCellRendererProgress.
/*

C function

gtk_cell_renderer_progress_new
*/
func CellRendererProgressNew() *CellRendererProgress {
	retC := C.gtk_cell_renderer_progress_new()
	retGo := CellRendererProgressNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Creates a new #GtkCellView widget.
/*

C function

gtk_cell_view_new
*/
func CellViewNew() *CellView {
	retC := C.gtk_cell_view_new()
	retGo := CellViewNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Creates a new #GtkCellView widget with a specific #GtkCellArea
// to layout cells and a specific #GtkCellAreaContext.
//
// Specifying the same context for a handfull of cells lets
// the underlying area synchronize the geometry for those cells,
// in this way alignments with cellviews for other rows are
// possible.
/*

C function

gtk_cell_view_new_with_context
*/
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
	retGo := CellViewNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Creates a new #GtkCellView widget, adds a #GtkCellRendererText
// to it, and makes it show @markup. The text can be
// marked up with the [Pango text markup language][PangoMarkupFormat].
/*

C function

gtk_cell_view_new_with_markup
*/
func CellViewNewWithMarkup(markup string) *CellView {
	c_markup := C.CString(markup)
	defer C.free(unsafe.Pointer(c_markup))

	retC := C.gtk_cell_view_new_with_markup(c_markup)
	retGo := CellViewNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Creates a new #GtkCellView widget, adds a #GtkCellRendererPixbuf
// to it, and makes it show @pixbuf.
/*

C function

gtk_cell_view_new_with_pixbuf
*/
func CellViewNewWithPixbuf(pixbuf *gdkpixbuf.Pixbuf) *CellView {
	c_pixbuf := (*C.GdkPixbuf)(C.NULL)
	if pixbuf != nil {
		c_pixbuf = (*C.GdkPixbuf)(pixbuf.ToC())
	}

	retC := C.gtk_cell_view_new_with_pixbuf(c_pixbuf)
	retGo := CellViewNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Creates a new #GtkCellView widget, adds a #GtkCellRendererText
// to it, and makes it show @text.
/*

C function

gtk_cell_view_new_with_text
*/
func CellViewNewWithText(text string) *CellView {
	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))

	retC := C.gtk_cell_view_new_with_text(c_text)
	retGo := CellViewNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Returns a #GtkTreePath referring to the currently
// displayed row. If no row is currently displayed,
// %NULL is returned.
/*

C function

gtk_cell_view_get_displayed_row
*/
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

// Sets @requisition to the size needed by @cell_view to display
// the model row pointed to by @path.
/*

C function

gtk_cell_view_get_size_of_row
*/
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

// Sets the background color of @view.
/*

C function

gtk_cell_view_set_background_color
*/
func (recv *CellView) SetBackgroundColor(color *gdk.Color) {
	c_color := (*C.GdkColor)(C.NULL)
	if color != nil {
		c_color = (*C.GdkColor)(color.ToC())
	}

	C.gtk_cell_view_set_background_color((*C.GtkCellView)(recv.native), c_color)

	return
}

// Sets the row of the model that is currently displayed
// by the #GtkCellView. If the path is unset, then the
// contents of the cellview “stick” at their last value;
// this is not normally a desired result, but may be
// a needed intermediate state if say, the model for
// the #GtkCellView becomes temporarily empty.
/*

C function

gtk_cell_view_set_displayed_row
*/
func (recv *CellView) SetDisplayedRow(path *TreePath) {
	c_path := (*C.GtkTreePath)(C.NULL)
	if path != nil {
		c_path = (*C.GtkTreePath)(path.ToC())
	}

	C.gtk_cell_view_set_displayed_row((*C.GtkCellView)(recv.native), c_path)

	return
}

// Sets the model for @cell_view.  If @cell_view already has a model
// set, it will remove it before setting the new model.  If @model is
// %NULL, then it will unset the old model.
/*

C function

gtk_cell_view_set_model
*/
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
var signalClipboardOwnerChangeLock sync.Mutex

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
	event := gdk.EventOwnerChangeNewFromC(unsafe.Pointer(c_event))

	index := int(uintptr(data))
	callback := signalClipboardOwnerChangeMap[index].callback
	callback(event)
}

// Unsupported : gtk_clipboard_request_image : unsupported parameter callback : no type generator for ClipboardImageReceivedFunc (GtkClipboardImageReceivedFunc) for param callback

// Unsupported : gtk_clipboard_set_can_store : unsupported parameter targets :

// Sets the contents of the clipboard to the given #GdkPixbuf.
// GTK+ will take responsibility for responding for requests
// for the image, and for converting the image into the
// requested format.
/*

C function

gtk_clipboard_set_image
*/
func (recv *Clipboard) SetImage(pixbuf *gdkpixbuf.Pixbuf) {
	c_pixbuf := (*C.GdkPixbuf)(C.NULL)
	if pixbuf != nil {
		c_pixbuf = (*C.GdkPixbuf)(pixbuf.ToC())
	}

	C.gtk_clipboard_set_image((*C.GtkClipboard)(recv.native), c_pixbuf)

	return
}

// Stores the current clipboard data somewhere so that it will stay
// around after the application has quit.
/*

C function

gtk_clipboard_store
*/
func (recv *Clipboard) Store() {
	C.gtk_clipboard_store((*C.GtkClipboard)(recv.native))

	return
}

// Requests the contents of the clipboard as image and converts
// the result to a #GdkPixbuf. This function waits for
// the data to be received using the main loop, so events,
// timeouts, etc, may be dispatched during the wait.
/*

C function

gtk_clipboard_wait_for_image
*/
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

// Test to see if there is an image available to be pasted
// This is done by requesting the TARGETS atom and checking
// if it contains any of the supported image targets. This function
// waits for the data to be received using the main loop, so events,
// timeouts, etc, may be dispatched during the wait.
//
// This function is a little faster than calling
// gtk_clipboard_wait_for_image() since it doesn’t need to retrieve
// the actual image data.
/*

C function

gtk_clipboard_wait_is_image_available
*/
func (recv *Clipboard) WaitIsImageAvailable() bool {
	retC := C.gtk_clipboard_wait_is_image_available((*C.GtkClipboard)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_clipboard_wait_is_target_available : unsupported parameter target : Blacklisted record : GdkAtom

// Returns the column with column span information for @combo_box.
/*

C function

gtk_combo_box_get_column_span_column
*/
func (recv *ComboBox) GetColumnSpanColumn() int32 {
	retC := C.gtk_combo_box_get_column_span_column((*C.GtkComboBox)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Returns whether the combo box grabs focus when it is clicked
// with the mouse. See gtk_combo_box_set_focus_on_click().
/*

C function

gtk_combo_box_get_focus_on_click
*/
func (recv *ComboBox) GetFocusOnClick() bool {
	retC := C.gtk_combo_box_get_focus_on_click((*C.GtkComboBox)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Gets the accessible object corresponding to the combo box’s popup.
//
// This function is mostly intended for use by accessibility technologies;
// applications should have little use for it.
/*

C function

gtk_combo_box_get_popup_accessible
*/
func (recv *ComboBox) GetPopupAccessible() *atk.Object {
	retC := C.gtk_combo_box_get_popup_accessible((*C.GtkComboBox)(recv.native))
	retGo := atk.ObjectNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_combo_box_get_row_separator_func : no return generator

// Returns the column with row span information for @combo_box.
/*

C function

gtk_combo_box_get_row_span_column
*/
func (recv *ComboBox) GetRowSpanColumn() int32 {
	retC := C.gtk_combo_box_get_row_span_column((*C.GtkComboBox)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Returns the wrap width which is used to determine the number of columns
// for the popup menu. If the wrap width is larger than 1, the combo box
// is in table mode.
/*

C function

gtk_combo_box_get_wrap_width
*/
func (recv *ComboBox) GetWrapWidth() int32 {
	retC := C.gtk_combo_box_get_wrap_width((*C.GtkComboBox)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Sets whether the popup menu should have a tearoff
// menu item.
/*

C function

gtk_combo_box_set_add_tearoffs
*/
func (recv *ComboBox) SetAddTearoffs(addTearoffs bool) {
	c_add_tearoffs :=
		boolToGboolean(addTearoffs)

	C.gtk_combo_box_set_add_tearoffs((*C.GtkComboBox)(recv.native), c_add_tearoffs)

	return
}

// Sets whether the combo box will grab focus when it is clicked with
// the mouse. Making mouse clicks not grab focus is useful in places
// like toolbars where you don’t want the keyboard focus removed from
// the main area of the application.
/*

C function

gtk_combo_box_set_focus_on_click
*/
func (recv *ComboBox) SetFocusOnClick(focusOnClick bool) {
	c_focus_on_click :=
		boolToGboolean(focusOnClick)

	C.gtk_combo_box_set_focus_on_click((*C.GtkComboBox)(recv.native), c_focus_on_click)

	return
}

// Unsupported : gtk_combo_box_set_row_separator_func : unsupported parameter func : no type generator for TreeViewRowSeparatorFunc (GtkTreeViewRowSeparatorFunc) for param func

// Unsupported : gtk_dialog_set_alternative_button_order : unsupported parameter ... : varargs

// Sets an alternative button order. If the
// #GtkSettings:gtk-alternative-button-order setting is set to %TRUE,
// the dialog buttons are reordered according to the order of the
// response ids in @new_order.
//
// See gtk_dialog_set_alternative_button_order() for more information.
//
// This function is for use by language bindings.
/*

C function

gtk_dialog_set_alternative_button_order_from_array
*/
func (recv *Dialog) SetAlternativeButtonOrderFromArray(newOrder []int32) {
	c_n_params := (C.gint)(len(newOrder))

	c_new_order := &newOrder[0]

	C.gtk_dialog_set_alternative_button_order_from_array((*C.GtkDialog)(recv.native), c_n_params, (*C.gint)(unsafe.Pointer(c_new_order)))

	return
}

// Unsupported signal 'insert-prefix' for EntryCompletion : unsupported parameter prefix : type utf8 :

// Returns whether the common prefix of the possible completions should
// be automatically inserted in the entry.
/*

C function

gtk_entry_completion_get_inline_completion
*/
func (recv *EntryCompletion) GetInlineCompletion() bool {
	retC := C.gtk_entry_completion_get_inline_completion((*C.GtkEntryCompletion)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Returns whether the completions should be presented in a popup window.
/*

C function

gtk_entry_completion_get_popup_completion
*/
func (recv *EntryCompletion) GetPopupCompletion() bool {
	retC := C.gtk_entry_completion_get_popup_completion((*C.GtkEntryCompletion)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Returns the column in the model of @completion to get strings from.
/*

C function

gtk_entry_completion_get_text_column
*/
func (recv *EntryCompletion) GetTextColumn() int32 {
	retC := C.gtk_entry_completion_get_text_column((*C.GtkEntryCompletion)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Requests a prefix insertion.
/*

C function

gtk_entry_completion_insert_prefix
*/
func (recv *EntryCompletion) InsertPrefix() {
	C.gtk_entry_completion_insert_prefix((*C.GtkEntryCompletion)(recv.native))

	return
}

// Sets whether the common prefix of the possible completions should
// be automatically inserted in the entry.
/*

C function

gtk_entry_completion_set_inline_completion
*/
func (recv *EntryCompletion) SetInlineCompletion(inlineCompletion bool) {
	c_inline_completion :=
		boolToGboolean(inlineCompletion)

	C.gtk_entry_completion_set_inline_completion((*C.GtkEntryCompletion)(recv.native), c_inline_completion)

	return
}

// Sets whether the completions should be presented in a popup window.
/*

C function

gtk_entry_completion_set_popup_completion
*/
func (recv *EntryCompletion) SetPopupCompletion(popupCompletion bool) {
	c_popup_completion :=
		boolToGboolean(popupCompletion)

	C.gtk_entry_completion_set_popup_completion((*C.GtkEntryCompletion)(recv.native), c_popup_completion)

	return
}

// Creates a new file-selecting button widget.
/*

C function

gtk_file_chooser_button_new
*/
func FileChooserButtonNew(title string, action FileChooserAction) *FileChooserButton {
	c_title := C.CString(title)
	defer C.free(unsafe.Pointer(c_title))

	c_action := (C.GtkFileChooserAction)(action)

	retC := C.gtk_file_chooser_button_new(c_title, c_action)
	retGo := FileChooserButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Creates a #GtkFileChooserButton widget which uses @dialog as its
// file-picking window.
//
// Note that @dialog must be a #GtkDialog (or subclass) which
// implements the #GtkFileChooser interface and must not have
// %GTK_DIALOG_DESTROY_WITH_PARENT set.
//
// Also note that the dialog needs to have its confirmative button
// added with response %GTK_RESPONSE_ACCEPT or %GTK_RESPONSE_OK in
// order for the button to take over the file selected in the dialog.
/*

C function

gtk_file_chooser_button_new_with_dialog
*/
func FileChooserButtonNewWithDialog(dialog *Dialog) *FileChooserButton {
	c_dialog := (*C.GtkWidget)(C.NULL)
	if dialog != nil {
		c_dialog = (*C.GtkWidget)(dialog.ToC())
	}

	retC := C.gtk_file_chooser_button_new_with_dialog(c_dialog)
	retGo := FileChooserButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Retrieves the title of the browse dialog used by @button. The returned value
// should not be modified or freed.
/*

C function

gtk_file_chooser_button_get_title
*/
func (recv *FileChooserButton) GetTitle() string {
	retC := C.gtk_file_chooser_button_get_title((*C.GtkFileChooserButton)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Retrieves the width in characters of the @button widget’s entry and/or label.
/*

C function

gtk_file_chooser_button_get_width_chars
*/
func (recv *FileChooserButton) GetWidthChars() int32 {
	retC := C.gtk_file_chooser_button_get_width_chars((*C.GtkFileChooserButton)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Modifies the @title of the browse dialog used by @button.
/*

C function

gtk_file_chooser_button_set_title
*/
func (recv *FileChooserButton) SetTitle(title string) {
	c_title := C.CString(title)
	defer C.free(unsafe.Pointer(c_title))

	C.gtk_file_chooser_button_set_title((*C.GtkFileChooserButton)(recv.native), c_title)

	return
}

// Sets the width (in characters) that @button will use to @n_chars.
/*

C function

gtk_file_chooser_button_set_width_chars
*/
func (recv *FileChooserButton) SetWidthChars(nChars int32) {
	c_n_chars := (C.gint)(nChars)

	C.gtk_file_chooser_button_set_width_chars((*C.GtkFileChooserButton)(recv.native), c_n_chars)

	return
}

// Adds a rule allowing image files in the formats supported
// by GdkPixbuf.
/*

C function

gtk_file_filter_add_pixbuf_formats
*/
func (recv *FileFilter) AddPixbufFormats() {
	C.gtk_file_filter_add_pixbuf_formats((*C.GtkFileFilter)(recv.native))

	return
}

// Unsupported : gtk_icon_theme_get_icon_sizes : no return type

// Creates a new #GtkIconView widget
/*

C function

gtk_icon_view_new
*/
func IconViewNew() *IconView {
	retC := C.gtk_icon_view_new()
	retGo := IconViewNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Creates a new #GtkIconView widget with the model @model.
/*

C function

gtk_icon_view_new_with_model
*/
func IconViewNewWithModel(model *TreeModel) *IconView {
	c_model := (*C.GtkTreeModel)(model.ToC())

	retC := C.gtk_icon_view_new_with_model(c_model)
	retGo := IconViewNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Returns the value of the ::column-spacing property.
/*

C function

gtk_icon_view_get_column_spacing
*/
func (recv *IconView) GetColumnSpacing() int32 {
	retC := C.gtk_icon_view_get_column_spacing((*C.GtkIconView)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Returns the value of the ::columns property.
/*

C function

gtk_icon_view_get_columns
*/
func (recv *IconView) GetColumns() int32 {
	retC := C.gtk_icon_view_get_columns((*C.GtkIconView)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Returns the value of the ::item-orientation property which determines
// whether the labels are drawn beside the icons instead of below.
/*

C function

gtk_icon_view_get_item_orientation
*/
func (recv *IconView) GetItemOrientation() Orientation {
	retC := C.gtk_icon_view_get_item_orientation((*C.GtkIconView)(recv.native))
	retGo := (Orientation)(retC)

	return retGo
}

// Returns the value of the ::item-width property.
/*

C function

gtk_icon_view_get_item_width
*/
func (recv *IconView) GetItemWidth() int32 {
	retC := C.gtk_icon_view_get_item_width((*C.GtkIconView)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Returns the value of the ::margin property.
/*

C function

gtk_icon_view_get_margin
*/
func (recv *IconView) GetMargin() int32 {
	retC := C.gtk_icon_view_get_margin((*C.GtkIconView)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Returns the column with markup text for @icon_view.
/*

C function

gtk_icon_view_get_markup_column
*/
func (recv *IconView) GetMarkupColumn() int32 {
	retC := C.gtk_icon_view_get_markup_column((*C.GtkIconView)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Returns the model the #GtkIconView is based on.  Returns %NULL if the
// model is unset.
/*

C function

gtk_icon_view_get_model
*/
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

// Finds the path at the point (@x, @y), relative to bin_window coordinates.
// See gtk_icon_view_get_item_at_pos(), if you are also interested in
// the cell at the specified position.
// See gtk_icon_view_convert_widget_to_bin_window_coords() for converting
// widget coordinates to bin_window coordinates.
/*

C function

gtk_icon_view_get_path_at_pos
*/
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

// Returns the column with pixbufs for @icon_view.
/*

C function

gtk_icon_view_get_pixbuf_column
*/
func (recv *IconView) GetPixbufColumn() int32 {
	retC := C.gtk_icon_view_get_pixbuf_column((*C.GtkIconView)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Returns the value of the ::row-spacing property.
/*

C function

gtk_icon_view_get_row_spacing
*/
func (recv *IconView) GetRowSpacing() int32 {
	retC := C.gtk_icon_view_get_row_spacing((*C.GtkIconView)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Creates a list of paths of all selected items. Additionally, if you are
// planning on modifying the model after calling this function, you may
// want to convert the returned list into a list of #GtkTreeRowReferences.
// To do this, you can use gtk_tree_row_reference_new().
//
// To free the return value, use:
// |[<!-- language="C" -->
// g_list_free_full (list, (GDestroyNotify) gtk_tree_path_free);
// ]|
/*

C function

gtk_icon_view_get_selected_items
*/
func (recv *IconView) GetSelectedItems() *glib.List {
	retC := C.gtk_icon_view_get_selected_items((*C.GtkIconView)(recv.native))
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the selection mode of the @icon_view.
/*

C function

gtk_icon_view_get_selection_mode
*/
func (recv *IconView) GetSelectionMode() SelectionMode {
	retC := C.gtk_icon_view_get_selection_mode((*C.GtkIconView)(recv.native))
	retGo := (SelectionMode)(retC)

	return retGo
}

// Returns the value of the ::spacing property.
/*

C function

gtk_icon_view_get_spacing
*/
func (recv *IconView) GetSpacing() int32 {
	retC := C.gtk_icon_view_get_spacing((*C.GtkIconView)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Returns the column with text for @icon_view.
/*

C function

gtk_icon_view_get_text_column
*/
func (recv *IconView) GetTextColumn() int32 {
	retC := C.gtk_icon_view_get_text_column((*C.GtkIconView)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Activates the item determined by @path.
/*

C function

gtk_icon_view_item_activated
*/
func (recv *IconView) ItemActivated(path *TreePath) {
	c_path := (*C.GtkTreePath)(C.NULL)
	if path != nil {
		c_path = (*C.GtkTreePath)(path.ToC())
	}

	C.gtk_icon_view_item_activated((*C.GtkIconView)(recv.native), c_path)

	return
}

// Returns %TRUE if the icon pointed to by @path is currently
// selected. If @path does not point to a valid location, %FALSE is returned.
/*

C function

gtk_icon_view_path_is_selected
*/
func (recv *IconView) PathIsSelected(path *TreePath) bool {
	c_path := (*C.GtkTreePath)(C.NULL)
	if path != nil {
		c_path = (*C.GtkTreePath)(path.ToC())
	}

	retC := C.gtk_icon_view_path_is_selected((*C.GtkIconView)(recv.native), c_path)
	retGo := retC == C.TRUE

	return retGo
}

// Selects all the icons. @icon_view must has its selection mode set
// to #GTK_SELECTION_MULTIPLE.
/*

C function

gtk_icon_view_select_all
*/
func (recv *IconView) SelectAll() {
	C.gtk_icon_view_select_all((*C.GtkIconView)(recv.native))

	return
}

// Selects the row at @path.
/*

C function

gtk_icon_view_select_path
*/
func (recv *IconView) SelectPath(path *TreePath) {
	c_path := (*C.GtkTreePath)(C.NULL)
	if path != nil {
		c_path = (*C.GtkTreePath)(path.ToC())
	}

	C.gtk_icon_view_select_path((*C.GtkIconView)(recv.native), c_path)

	return
}

// Unsupported : gtk_icon_view_selected_foreach : unsupported parameter func : no type generator for IconViewForeachFunc (GtkIconViewForeachFunc) for param func

// Sets the ::column-spacing property which specifies the space
// which is inserted between the columns of the icon view.
/*

C function

gtk_icon_view_set_column_spacing
*/
func (recv *IconView) SetColumnSpacing(columnSpacing int32) {
	c_column_spacing := (C.gint)(columnSpacing)

	C.gtk_icon_view_set_column_spacing((*C.GtkIconView)(recv.native), c_column_spacing)

	return
}

// Sets the ::columns property which determines in how
// many columns the icons are arranged. If @columns is
// -1, the number of columns will be chosen automatically
// to fill the available area.
/*

C function

gtk_icon_view_set_columns
*/
func (recv *IconView) SetColumns(columns int32) {
	c_columns := (C.gint)(columns)

	C.gtk_icon_view_set_columns((*C.GtkIconView)(recv.native), c_columns)

	return
}

// Sets the ::item-orientation property which determines whether the labels
// are drawn beside the icons instead of below.
/*

C function

gtk_icon_view_set_item_orientation
*/
func (recv *IconView) SetItemOrientation(orientation Orientation) {
	c_orientation := (C.GtkOrientation)(orientation)

	C.gtk_icon_view_set_item_orientation((*C.GtkIconView)(recv.native), c_orientation)

	return
}

// Sets the ::item-width property which specifies the width
// to use for each item. If it is set to -1, the icon view will
// automatically determine a suitable item size.
/*

C function

gtk_icon_view_set_item_width
*/
func (recv *IconView) SetItemWidth(itemWidth int32) {
	c_item_width := (C.gint)(itemWidth)

	C.gtk_icon_view_set_item_width((*C.GtkIconView)(recv.native), c_item_width)

	return
}

// Sets the ::margin property which specifies the space
// which is inserted at the top, bottom, left and right
// of the icon view.
/*

C function

gtk_icon_view_set_margin
*/
func (recv *IconView) SetMargin(margin int32) {
	c_margin := (C.gint)(margin)

	C.gtk_icon_view_set_margin((*C.GtkIconView)(recv.native), c_margin)

	return
}

// Sets the column with markup information for @icon_view to be
// @column. The markup column must be of type #G_TYPE_STRING.
// If the markup column is set to something, it overrides
// the text column set by gtk_icon_view_set_text_column().
/*

C function

gtk_icon_view_set_markup_column
*/
func (recv *IconView) SetMarkupColumn(column int32) {
	c_column := (C.gint)(column)

	C.gtk_icon_view_set_markup_column((*C.GtkIconView)(recv.native), c_column)

	return
}

// Sets the model for a #GtkIconView.
// If the @icon_view already has a model set, it will remove
// it before setting the new model.  If @model is %NULL, then
// it will unset the old model.
/*

C function

gtk_icon_view_set_model
*/
func (recv *IconView) SetModel(model *TreeModel) {
	c_model := (*C.GtkTreeModel)(model.ToC())

	C.gtk_icon_view_set_model((*C.GtkIconView)(recv.native), c_model)

	return
}

// Sets the column with pixbufs for @icon_view to be @column. The pixbuf
// column must be of type #GDK_TYPE_PIXBUF
/*

C function

gtk_icon_view_set_pixbuf_column
*/
func (recv *IconView) SetPixbufColumn(column int32) {
	c_column := (C.gint)(column)

	C.gtk_icon_view_set_pixbuf_column((*C.GtkIconView)(recv.native), c_column)

	return
}

// Sets the ::row-spacing property which specifies the space
// which is inserted between the rows of the icon view.
/*

C function

gtk_icon_view_set_row_spacing
*/
func (recv *IconView) SetRowSpacing(rowSpacing int32) {
	c_row_spacing := (C.gint)(rowSpacing)

	C.gtk_icon_view_set_row_spacing((*C.GtkIconView)(recv.native), c_row_spacing)

	return
}

// Sets the selection mode of the @icon_view.
/*

C function

gtk_icon_view_set_selection_mode
*/
func (recv *IconView) SetSelectionMode(mode SelectionMode) {
	c_mode := (C.GtkSelectionMode)(mode)

	C.gtk_icon_view_set_selection_mode((*C.GtkIconView)(recv.native), c_mode)

	return
}

// Sets the ::spacing property which specifies the space
// which is inserted between the cells (i.e. the icon and
// the text) of an item.
/*

C function

gtk_icon_view_set_spacing
*/
func (recv *IconView) SetSpacing(spacing int32) {
	c_spacing := (C.gint)(spacing)

	C.gtk_icon_view_set_spacing((*C.GtkIconView)(recv.native), c_spacing)

	return
}

// Sets the column with text for @icon_view to be @column. The text
// column must be of type #G_TYPE_STRING.
/*

C function

gtk_icon_view_set_text_column
*/
func (recv *IconView) SetTextColumn(column int32) {
	c_column := (C.gint)(column)

	C.gtk_icon_view_set_text_column((*C.GtkIconView)(recv.native), c_column)

	return
}

// Unselects all the icons.
/*

C function

gtk_icon_view_unselect_all
*/
func (recv *IconView) UnselectAll() {
	C.gtk_icon_view_unselect_all((*C.GtkIconView)(recv.native))

	return
}

// Unselects the row at @path.
/*

C function

gtk_icon_view_unselect_path
*/
func (recv *IconView) UnselectPath(path *TreePath) {
	c_path := (*C.GtkTreePath)(C.NULL)
	if path != nil {
		c_path = (*C.GtkTreePath)(path.ToC())
	}

	C.gtk_icon_view_unselect_path((*C.GtkIconView)(recv.native), c_path)

	return
}

// Creates a #GtkImage displaying an icon from the current icon theme.
// If the icon name isn’t known, a “broken image” icon will be
// displayed instead.  If the current icon theme is changed, the icon
// will be updated appropriately.
/*

C function

gtk_image_new_from_icon_name
*/
func ImageNewFromIconName(iconName string, size IconSize) *Image {
	c_icon_name := C.CString(iconName)
	defer C.free(unsafe.Pointer(c_icon_name))

	c_size := (C.GtkIconSize)(size)

	retC := C.gtk_image_new_from_icon_name(c_icon_name, c_size)
	retGo := ImageNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_image_get_icon_name : unsupported parameter size : no type generator for gint (GtkIconSize*) for param size

// Gets the pixel size used for named icons.
/*

C function

gtk_image_get_pixel_size
*/
func (recv *Image) GetPixelSize() int32 {
	retC := C.gtk_image_get_pixel_size((*C.GtkImage)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// See gtk_image_new_from_icon_name() for details.
/*

C function

gtk_image_set_from_icon_name
*/
func (recv *Image) SetFromIconName(iconName string, size IconSize) {
	c_icon_name := C.CString(iconName)
	defer C.free(unsafe.Pointer(c_icon_name))

	c_size := (C.GtkIconSize)(size)

	C.gtk_image_set_from_icon_name((*C.GtkImage)(recv.native), c_icon_name, c_size)

	return
}

// Sets the pixel size to use for named icons. If the pixel size is set
// to a value != -1, it is used instead of the icon size set by
// gtk_image_set_from_icon_name().
/*

C function

gtk_image_set_pixel_size
*/
func (recv *Image) SetPixelSize(pixelSize int32) {
	c_pixel_size := (C.gint)(pixelSize)

	C.gtk_image_set_pixel_size((*C.GtkImage)(recv.native), c_pixel_size)

	return
}

// Gets the angle of rotation for the label. See
// gtk_label_set_angle().
/*

C function

gtk_label_get_angle
*/
func (recv *Label) GetAngle() float64 {
	retC := C.gtk_label_get_angle((*C.GtkLabel)(recv.native))
	retGo := (float64)(retC)

	return retGo
}

// Returns the ellipsizing position of the label. See gtk_label_set_ellipsize().
/*

C function

gtk_label_get_ellipsize
*/
func (recv *Label) GetEllipsize() pango.EllipsizeMode {
	retC := C.gtk_label_get_ellipsize((*C.GtkLabel)(recv.native))
	retGo := (pango.EllipsizeMode)(retC)

	return retGo
}

// Retrieves the desired maximum width of @label, in characters. See
// gtk_label_set_width_chars().
/*

C function

gtk_label_get_max_width_chars
*/
func (recv *Label) GetMaxWidthChars() int32 {
	retC := C.gtk_label_get_max_width_chars((*C.GtkLabel)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Returns whether the label is in single line mode.
/*

C function

gtk_label_get_single_line_mode
*/
func (recv *Label) GetSingleLineMode() bool {
	retC := C.gtk_label_get_single_line_mode((*C.GtkLabel)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Retrieves the desired width of @label, in characters. See
// gtk_label_set_width_chars().
/*

C function

gtk_label_get_width_chars
*/
func (recv *Label) GetWidthChars() int32 {
	retC := C.gtk_label_get_width_chars((*C.GtkLabel)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Sets the angle of rotation for the label. An angle of 90 reads from
// from bottom to top, an angle of 270, from top to bottom. The angle
// setting for the label is ignored if the label is selectable,
// wrapped, or ellipsized.
/*

C function

gtk_label_set_angle
*/
func (recv *Label) SetAngle(angle float64) {
	c_angle := (C.gdouble)(angle)

	C.gtk_label_set_angle((*C.GtkLabel)(recv.native), c_angle)

	return
}

// Sets the mode used to ellipsize (add an ellipsis: "...") to the text
// if there is not enough space to render the entire string.
/*

C function

gtk_label_set_ellipsize
*/
func (recv *Label) SetEllipsize(mode pango.EllipsizeMode) {
	c_mode := (C.PangoEllipsizeMode)(mode)

	C.gtk_label_set_ellipsize((*C.GtkLabel)(recv.native), c_mode)

	return
}

// Sets the desired maximum width in characters of @label to @n_chars.
/*

C function

gtk_label_set_max_width_chars
*/
func (recv *Label) SetMaxWidthChars(nChars int32) {
	c_n_chars := (C.gint)(nChars)

	C.gtk_label_set_max_width_chars((*C.GtkLabel)(recv.native), c_n_chars)

	return
}

// Sets whether the label is in single line mode.
/*

C function

gtk_label_set_single_line_mode
*/
func (recv *Label) SetSingleLineMode(singleLineMode bool) {
	c_single_line_mode :=
		boolToGboolean(singleLineMode)

	C.gtk_label_set_single_line_mode((*C.GtkLabel)(recv.native), c_single_line_mode)

	return
}

// Sets the desired width in characters of @label to @n_chars.
/*

C function

gtk_label_set_width_chars
*/
func (recv *Label) SetWidthChars(nChars int32) {
	c_n_chars := (C.gint)(nChars)

	C.gtk_label_set_width_chars((*C.GtkLabel)(recv.native), c_n_chars)

	return
}

// Unsupported : gtk_list_store_insert_with_values : unsupported parameter ... : varargs

// Unsupported : gtk_list_store_insert_with_valuesv : unsupported parameter values :

// Creates a new #GtkMenuToolButton using @icon_widget as icon and
// @label as label.
/*

C function

gtk_menu_tool_button_new
*/
func MenuToolButtonNew(iconWidget *Widget, label string) *MenuToolButton {
	c_icon_widget := (*C.GtkWidget)(C.NULL)
	if iconWidget != nil {
		c_icon_widget = (*C.GtkWidget)(iconWidget.ToC())
	}

	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	retC := C.gtk_menu_tool_button_new(c_icon_widget, c_label)
	retGo := MenuToolButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Creates a new #GtkMenuToolButton.
// The new #GtkMenuToolButton will contain an icon and label from
// the stock item indicated by @stock_id.
/*

C function

gtk_menu_tool_button_new_from_stock
*/
func MenuToolButtonNewFromStock(stockId string) *MenuToolButton {
	c_stock_id := C.CString(stockId)
	defer C.free(unsafe.Pointer(c_stock_id))

	retC := C.gtk_menu_tool_button_new_from_stock(c_stock_id)
	retGo := MenuToolButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the #GtkMenu associated with #GtkMenuToolButton.
/*

C function

gtk_menu_tool_button_get_menu
*/
func (recv *MenuToolButton) GetMenu() *Widget {
	retC := C.gtk_menu_tool_button_get_menu((*C.GtkMenuToolButton)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Sets the #GtkMenu that is popped up when the user clicks on the arrow.
// If @menu is NULL, the arrow button becomes insensitive.
/*

C function

gtk_menu_tool_button_set_menu
*/
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

// Returns the ellipsizing position of the progress bar.
// See gtk_progress_bar_set_ellipsize().
/*

C function

gtk_progress_bar_get_ellipsize
*/
func (recv *ProgressBar) GetEllipsize() pango.EllipsizeMode {
	retC := C.gtk_progress_bar_get_ellipsize((*C.GtkProgressBar)(recv.native))
	retGo := (pango.EllipsizeMode)(retC)

	return retGo
}

// Sets the mode used to ellipsize (add an ellipsis: "...") the
// text if there is not enough space to render the entire string.
/*

C function

gtk_progress_bar_set_ellipsize
*/
func (recv *ProgressBar) SetEllipsize(mode pango.EllipsizeMode) {
	c_mode := (C.PangoEllipsizeMode)(mode)

	C.gtk_progress_bar_set_ellipsize((*C.GtkProgressBar)(recv.native), c_mode)

	return
}

// Unsupported signal 'change-value' for Range : unsupported parameter scroll : type ScrollType :

// Performs the appropriate action as if the user hit the delete
// key with the cursor at the position specified by @iter. In the
// normal case a single character will be deleted, but when
// combining accents are involved, more than one character can
// be deleted, and when precomposed character and accent combinations
// are involved, less than one character will be deleted.
//
// Because the buffer is modified, all outstanding iterators become
// invalid after calling this function; however, the @iter will be
// re-initialized to point to the location where text was deleted.
/*

C function

gtk_text_buffer_backspace
*/
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

// Retrieves the iterator pointing to the character at buffer
// coordinates @x and @y. Buffer coordinates are coordinates for
// the entire buffer, not just the currently-displayed portion.
// If you have coordinates from an event, you have to convert
// those to buffer coordinates with
// gtk_text_view_window_to_buffer_coords().
//
// Note that this is different from gtk_text_view_get_iter_at_location(),
// which returns cursor locations, i.e. positions between
// characters.
/*

C function

gtk_text_view_get_iter_at_position
*/
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

// Calling this function signals to the toolbar that the
// overflow menu item for @tool_item has changed. If the
// overflow menu is visible when this function it called,
// the menu will be rebuilt.
//
// The function must be called when the tool item changes what it
// will do in response to the #GtkToolItem::create-menu-proxy signal.
/*

C function

gtk_tool_item_rebuild_menu
*/
func (recv *ToolItem) RebuildMenu() {
	C.gtk_tool_item_rebuild_menu((*C.GtkToolItem)(recv.native))

	return
}

// Returns whether fixed height mode is turned on for @tree_view.
/*

C function

gtk_tree_view_get_fixed_height_mode
*/
func (recv *TreeView) GetFixedHeightMode() bool {
	retC := C.gtk_tree_view_get_fixed_height_mode((*C.GtkTreeView)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Returns whether hover expansion mode is turned on for @tree_view.
/*

C function

gtk_tree_view_get_hover_expand
*/
func (recv *TreeView) GetHoverExpand() bool {
	retC := C.gtk_tree_view_get_hover_expand((*C.GtkTreeView)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Returns whether hover selection mode is turned on for @tree_view.
/*

C function

gtk_tree_view_get_hover_selection
*/
func (recv *TreeView) GetHoverSelection() bool {
	retC := C.gtk_tree_view_get_hover_selection((*C.GtkTreeView)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_tree_view_get_row_separator_func : no return generator

// Enables or disables the fixed height mode of @tree_view.
// Fixed height mode speeds up #GtkTreeView by assuming that all
// rows have the same height.
// Only enable this option if all rows are the same height and all
// columns are of type %GTK_TREE_VIEW_COLUMN_FIXED.
/*

C function

gtk_tree_view_set_fixed_height_mode
*/
func (recv *TreeView) SetFixedHeightMode(enable bool) {
	c_enable :=
		boolToGboolean(enable)

	C.gtk_tree_view_set_fixed_height_mode((*C.GtkTreeView)(recv.native), c_enable)

	return
}

// Enables or disables the hover expansion mode of @tree_view.
// Hover expansion makes rows expand or collapse if the pointer
// moves over them.
/*

C function

gtk_tree_view_set_hover_expand
*/
func (recv *TreeView) SetHoverExpand(expand bool) {
	c_expand :=
		boolToGboolean(expand)

	C.gtk_tree_view_set_hover_expand((*C.GtkTreeView)(recv.native), c_expand)

	return
}

// Enables or disables the hover selection mode of @tree_view.
// Hover selection makes the selected row follow the pointer.
// Currently, this works only for the selection modes
// %GTK_SELECTION_SINGLE and %GTK_SELECTION_BROWSE.
/*

C function

gtk_tree_view_set_hover_selection
*/
func (recv *TreeView) SetHoverSelection(hover bool) {
	c_hover :=
		boolToGboolean(hover)

	C.gtk_tree_view_set_hover_selection((*C.GtkTreeView)(recv.native), c_hover)

	return
}

// Unsupported : gtk_tree_view_set_row_separator_func : unsupported parameter func : no type generator for TreeViewRowSeparatorFunc (GtkTreeViewRowSeparatorFunc) for param func

// Add the image targets supported by #GtkSelectionData to
// the target list of the drag destination. The targets
// are added with @info = 0. If you need another value,
// use gtk_target_list_add_image_targets() and
// gtk_drag_dest_set_target_list().
/*

C function

gtk_drag_dest_add_image_targets
*/
func (recv *Widget) DragDestAddImageTargets() {
	C.gtk_drag_dest_add_image_targets((*C.GtkWidget)(recv.native))

	return
}

// Add the text targets supported by #GtkSelectionData to
// the target list of the drag destination. The targets
// are added with @info = 0. If you need another value,
// use gtk_target_list_add_text_targets() and
// gtk_drag_dest_set_target_list().
/*

C function

gtk_drag_dest_add_text_targets
*/
func (recv *Widget) DragDestAddTextTargets() {
	C.gtk_drag_dest_add_text_targets((*C.GtkWidget)(recv.native))

	return
}

// Add the URI targets supported by #GtkSelectionData to
// the target list of the drag destination. The targets
// are added with @info = 0. If you need another value,
// use gtk_target_list_add_uri_targets() and
// gtk_drag_dest_set_target_list().
/*

C function

gtk_drag_dest_add_uri_targets
*/
func (recv *Widget) DragDestAddUriTargets() {
	C.gtk_drag_dest_add_uri_targets((*C.GtkWidget)(recv.native))

	return
}

// Add the writable image targets supported by #GtkSelectionData to
// the target list of the drag source. The targets
// are added with @info = 0. If you need another value,
// use gtk_target_list_add_image_targets() and
// gtk_drag_source_set_target_list().
/*

C function

gtk_drag_source_add_image_targets
*/
func (recv *Widget) DragSourceAddImageTargets() {
	C.gtk_drag_source_add_image_targets((*C.GtkWidget)(recv.native))

	return
}

// Add the text targets supported by #GtkSelectionData to
// the target list of the drag source.  The targets
// are added with @info = 0. If you need another value,
// use gtk_target_list_add_text_targets() and
// gtk_drag_source_set_target_list().
/*

C function

gtk_drag_source_add_text_targets
*/
func (recv *Widget) DragSourceAddTextTargets() {
	C.gtk_drag_source_add_text_targets((*C.GtkWidget)(recv.native))

	return
}

// Add the URI targets supported by #GtkSelectionData to
// the target list of the drag source.  The targets
// are added with @info = 0. If you need another value,
// use gtk_target_list_add_uri_targets() and
// gtk_drag_source_set_target_list().
/*

C function

gtk_drag_source_add_uri_targets
*/
func (recv *Widget) DragSourceAddUriTargets() {
	C.gtk_drag_source_add_uri_targets((*C.GtkWidget)(recv.native))

	return
}

// Gets the value set by gtk_window_set_focus_on_map().
/*

C function

gtk_window_get_focus_on_map
*/
func (recv *Window) GetFocusOnMap() bool {
	retC := C.gtk_window_get_focus_on_map((*C.GtkWindow)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Returns the name of the themed icon for the window,
// see gtk_window_set_icon_name().
/*

C function

gtk_window_get_icon_name
*/
func (recv *Window) GetIconName() string {
	retC := C.gtk_window_get_icon_name((*C.GtkWindow)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Windows may set a hint asking the desktop environment not to receive
// the input focus when the window is mapped.  This function sets this
// hint.
/*

C function

gtk_window_set_focus_on_map
*/
func (recv *Window) SetFocusOnMap(setting bool) {
	c_setting :=
		boolToGboolean(setting)

	C.gtk_window_set_focus_on_map((*C.GtkWindow)(recv.native), c_setting)

	return
}

// Sets the icon for the window from a named themed icon.
// See the docs for #GtkIconTheme for more details.
// On some platforms, the window icon is not used at all.
//
// Note that this has nothing to do with the WM_ICON_NAME
// property which is mentioned in the ICCCM.
/*

C function

gtk_window_set_icon_name
*/
func (recv *Window) SetIconName(name string) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	C.gtk_window_set_icon_name((*C.GtkWindow)(recv.native), c_name)

	return
}
