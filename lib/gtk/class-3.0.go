// This is a generated file - DO NOT EDIT
// +build gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	cairo "github.com/pekim/gobbi/lib/cairo"
	gdk "github.com/pekim/gobbi/lib/gdk"
	gdkpixbuf "github.com/pekim/gobbi/lib/gdkpixbuf"
	gio "github.com/pekim/gobbi/lib/gio"
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

	void cellarea_applyAttributesHandler(GObject *, GtkTreeModel *, GtkTreeIter *, gboolean, gboolean, gpointer);

	static gulong CellArea_signal_connect_apply_attributes(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "apply-attributes", G_CALLBACK(cellarea_applyAttributesHandler), data);
	}

*/
/*

	void cellarea_removeEditableHandler(GObject *, GtkCellRenderer *, GtkCellEditable *, gpointer);

	static gulong CellArea_signal_connect_remove_editable(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "remove-editable", G_CALLBACK(cellarea_removeEditableHandler), data);
	}

*/
/*

	void stylecontext_changedHandler(GObject *, gpointer);

	static gulong StyleContext_signal_connect_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "changed", G_CALLBACK(stylecontext_changedHandler), data);
	}

*/
/*

	gboolean widget_drawHandler(GObject *, cairo_t *, gpointer);

	static gulong Widget_signal_connect_draw(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "draw", G_CALLBACK(widget_drawHandler), data);
	}

*/
/*

	void widget_styleUpdatedHandler(GObject *, gpointer);

	static gulong Widget_signal_connect_style_updated(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "style-updated", G_CALLBACK(widget_styleUpdatedHandler), data);
	}

*/
import "C"

// Retrieves the license set using gtk_about_dialog_set_license_type()
/*

C function

gtk_about_dialog_get_license_type
*/
func (recv *AboutDialog) GetLicenseType() License {
	retC := C.gtk_about_dialog_get_license_type((*C.GtkAboutDialog)(recv.native))
	retGo := (License)(retC)

	return retGo
}

// Sets the license of the application showing the @about dialog from a
// list of known licenses.
//
// This function overrides the license set using
// gtk_about_dialog_set_license().
/*

C function

gtk_about_dialog_set_license_type
*/
func (recv *AboutDialog) SetLicenseType(licenseType License) {
	c_license_type := (C.GtkLicense)(licenseType)

	C.gtk_about_dialog_set_license_type((*C.GtkAboutDialog)(recv.native), c_license_type)

	return
}

// Creates a new #GtkAppChooserButton for applications
// that can handle content of the given type.
/*

C function

gtk_app_chooser_button_new
*/
func AppChooserButtonNew(contentType string) *AppChooserButton {
	c_content_type := C.CString(contentType)
	defer C.free(unsafe.Pointer(c_content_type))

	retC := C.gtk_app_chooser_button_new(c_content_type)
	retGo := AppChooserButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Appends a custom item to the list of applications that is shown
// in the popup; the item name must be unique per-widget.
// Clients can use the provided name as a detail for the
// #GtkAppChooserButton::custom-item-activated signal, to add a
// callback for the activation of a particular custom item in the list.
// See also gtk_app_chooser_button_append_separator().
/*

C function

gtk_app_chooser_button_append_custom_item
*/
func (recv *AppChooserButton) AppendCustomItem(name string, label string, icon *gio.Icon) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	c_icon := (*C.GIcon)(icon.ToC())

	C.gtk_app_chooser_button_append_custom_item((*C.GtkAppChooserButton)(recv.native), c_name, c_label, c_icon)

	return
}

// Appends a separator to the list of applications that is shown
// in the popup.
/*

C function

gtk_app_chooser_button_append_separator
*/
func (recv *AppChooserButton) AppendSeparator() {
	C.gtk_app_chooser_button_append_separator((*C.GtkAppChooserButton)(recv.native))

	return
}

// Returns the current value of the #GtkAppChooserButton:show-dialog-item
// property.
/*

C function

gtk_app_chooser_button_get_show_dialog_item
*/
func (recv *AppChooserButton) GetShowDialogItem() bool {
	retC := C.gtk_app_chooser_button_get_show_dialog_item((*C.GtkAppChooserButton)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Selects a custom item previously added with
// gtk_app_chooser_button_append_custom_item().
//
// Use gtk_app_chooser_refresh() to bring the selection
// to its initial state.
/*

C function

gtk_app_chooser_button_set_active_custom_item
*/
func (recv *AppChooserButton) SetActiveCustomItem(name string) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	C.gtk_app_chooser_button_set_active_custom_item((*C.GtkAppChooserButton)(recv.native), c_name)

	return
}

// Sets whether the dropdown menu of this button should show an
// entry to trigger a #GtkAppChooserDialog.
/*

C function

gtk_app_chooser_button_set_show_dialog_item
*/
func (recv *AppChooserButton) SetShowDialogItem(setting bool) {
	c_setting :=
		boolToGboolean(setting)

	C.gtk_app_chooser_button_set_show_dialog_item((*C.GtkAppChooserButton)(recv.native), c_setting)

	return
}

// Creates a new #GtkAppChooserDialog for the provided #GFile,
// to allow the user to select an application for it.
/*

C function

gtk_app_chooser_dialog_new
*/
func AppChooserDialogNew(parent *Window, flags DialogFlags, file *gio.File) *AppChooserDialog {
	c_parent := (*C.GtkWindow)(C.NULL)
	if parent != nil {
		c_parent = (*C.GtkWindow)(parent.ToC())
	}

	c_flags := (C.GtkDialogFlags)(flags)

	c_file := (*C.GFile)(file.ToC())

	retC := C.gtk_app_chooser_dialog_new(c_parent, c_flags, c_file)
	retGo := AppChooserDialogNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Creates a new #GtkAppChooserDialog for the provided content type,
// to allow the user to select an application for it.
/*

C function

gtk_app_chooser_dialog_new_for_content_type
*/
func AppChooserDialogNewForContentType(parent *Window, flags DialogFlags, contentType string) *AppChooserDialog {
	c_parent := (*C.GtkWindow)(C.NULL)
	if parent != nil {
		c_parent = (*C.GtkWindow)(parent.ToC())
	}

	c_flags := (C.GtkDialogFlags)(flags)

	c_content_type := C.CString(contentType)
	defer C.free(unsafe.Pointer(c_content_type))

	retC := C.gtk_app_chooser_dialog_new_for_content_type(c_parent, c_flags, c_content_type)
	retGo := AppChooserDialogNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Returns the #GtkAppChooserWidget of this dialog.
/*

C function

gtk_app_chooser_dialog_get_widget
*/
func (recv *AppChooserDialog) GetWidget() *Widget {
	retC := C.gtk_app_chooser_dialog_get_widget((*C.GtkAppChooserDialog)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Creates a new #GtkAppChooserWidget for applications
// that can handle content of the given type.
/*

C function

gtk_app_chooser_widget_new
*/
func AppChooserWidgetNew(contentType string) *AppChooserWidget {
	c_content_type := C.CString(contentType)
	defer C.free(unsafe.Pointer(c_content_type))

	retC := C.gtk_app_chooser_widget_new(c_content_type)
	retGo := AppChooserWidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Returns the text that is shown if there are not applications
// that can handle the content type.
/*

C function

gtk_app_chooser_widget_get_default_text
*/
func (recv *AppChooserWidget) GetDefaultText() string {
	retC := C.gtk_app_chooser_widget_get_default_text((*C.GtkAppChooserWidget)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Returns the current value of the #GtkAppChooserWidget:show-all
// property.
/*

C function

gtk_app_chooser_widget_get_show_all
*/
func (recv *AppChooserWidget) GetShowAll() bool {
	retC := C.gtk_app_chooser_widget_get_show_all((*C.GtkAppChooserWidget)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Returns the current value of the #GtkAppChooserWidget:show-default
// property.
/*

C function

gtk_app_chooser_widget_get_show_default
*/
func (recv *AppChooserWidget) GetShowDefault() bool {
	retC := C.gtk_app_chooser_widget_get_show_default((*C.GtkAppChooserWidget)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Returns the current value of the #GtkAppChooserWidget:show-fallback
// property.
/*

C function

gtk_app_chooser_widget_get_show_fallback
*/
func (recv *AppChooserWidget) GetShowFallback() bool {
	retC := C.gtk_app_chooser_widget_get_show_fallback((*C.GtkAppChooserWidget)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Returns the current value of the #GtkAppChooserWidget:show-other
// property.
/*

C function

gtk_app_chooser_widget_get_show_other
*/
func (recv *AppChooserWidget) GetShowOther() bool {
	retC := C.gtk_app_chooser_widget_get_show_other((*C.GtkAppChooserWidget)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Returns the current value of the #GtkAppChooserWidget:show-recommended
// property.
/*

C function

gtk_app_chooser_widget_get_show_recommended
*/
func (recv *AppChooserWidget) GetShowRecommended() bool {
	retC := C.gtk_app_chooser_widget_get_show_recommended((*C.GtkAppChooserWidget)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Sets whether the app chooser should show all applications
// in a flat list.
/*

C function

gtk_app_chooser_widget_set_show_all
*/
func (recv *AppChooserWidget) SetShowAll(setting bool) {
	c_setting :=
		boolToGboolean(setting)

	C.gtk_app_chooser_widget_set_show_all((*C.GtkAppChooserWidget)(recv.native), c_setting)

	return
}

// Sets whether the app chooser should show the default handler
// for the content type in a separate section.
/*

C function

gtk_app_chooser_widget_set_show_default
*/
func (recv *AppChooserWidget) SetShowDefault(setting bool) {
	c_setting :=
		boolToGboolean(setting)

	C.gtk_app_chooser_widget_set_show_default((*C.GtkAppChooserWidget)(recv.native), c_setting)

	return
}

// Sets whether the app chooser should show related applications
// for the content type in a separate section.
/*

C function

gtk_app_chooser_widget_set_show_fallback
*/
func (recv *AppChooserWidget) SetShowFallback(setting bool) {
	c_setting :=
		boolToGboolean(setting)

	C.gtk_app_chooser_widget_set_show_fallback((*C.GtkAppChooserWidget)(recv.native), c_setting)

	return
}

// Sets whether the app chooser should show applications
// which are unrelated to the content type.
/*

C function

gtk_app_chooser_widget_set_show_other
*/
func (recv *AppChooserWidget) SetShowOther(setting bool) {
	c_setting :=
		boolToGboolean(setting)

	C.gtk_app_chooser_widget_set_show_other((*C.GtkAppChooserWidget)(recv.native), c_setting)

	return
}

// Sets whether the app chooser should show recommended applications
// for the content type in a separate section.
/*

C function

gtk_app_chooser_widget_set_show_recommended
*/
func (recv *AppChooserWidget) SetShowRecommended(setting bool) {
	c_setting :=
		boolToGboolean(setting)

	C.gtk_app_chooser_widget_set_show_recommended((*C.GtkAppChooserWidget)(recv.native), c_setting)

	return
}

// Creates a new #GtkApplication instance.
//
// When using #GtkApplication, it is not necessary to call gtk_init()
// manually. It is called as soon as the application gets registered as
// the primary instance.
//
// Concretely, gtk_init() is called in the default handler for the
// #GApplication::startup signal. Therefore, #GtkApplication subclasses should
// chain up in their #GApplication::startup handler before using any GTK+ API.
//
// Note that commandline arguments are not passed to gtk_init().
// All GTK+ functionality that is available via commandline arguments
// can also be achieved by setting suitable environment variables
// such as `G_DEBUG`, so this should not be a big
// problem. If you absolutely must support GTK+ commandline arguments,
// you can explicitly call gtk_init() before creating the application
// instance.
//
// If non-%NULL, the application ID must be valid.  See
// g_application_id_is_valid().
//
// If no application ID is given then some features (most notably application
// uniqueness) will be disabled. A null application ID is only allowed with
// GTK+ 3.6 or later.
/*

C function

gtk_application_new
*/
func ApplicationNew(applicationId string, flags gio.ApplicationFlags) *Application {
	c_application_id := C.CString(applicationId)
	defer C.free(unsafe.Pointer(c_application_id))

	c_flags := (C.GApplicationFlags)(flags)

	retC := C.gtk_application_new(c_application_id, c_flags)
	retGo := ApplicationNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Adds a window to @application.
//
// This call can only happen after the @application has started;
// typically, you should add new application windows in response
// to the emission of the #GApplication::activate signal.
//
// This call is equivalent to setting the #GtkWindow:application
// property of @window to @application.
//
// Normally, the connection between the application and the window
// will remain until the window is destroyed, but you can explicitly
// remove it with gtk_application_remove_window().
//
// GTK+ will keep the @application running as long as it has
// any windows.
/*

C function

gtk_application_add_window
*/
func (recv *Application) AddWindow(window *Window) {
	c_window := (*C.GtkWindow)(C.NULL)
	if window != nil {
		c_window = (*C.GtkWindow)(window.ToC())
	}

	C.gtk_application_add_window((*C.GtkApplication)(recv.native), c_window)

	return
}

// Gets a list of the #GtkWindows associated with @application.
//
// The list is sorted by most recently focused window, such that the first
// element is the currently focused window. (Useful for choosing a parent
// for a transient window.)
//
// The list that is returned should not be modified in any way. It will
// only remain valid until the next focus change or window creation or
// deletion.
/*

C function

gtk_application_get_windows
*/
func (recv *Application) GetWindows() *glib.List {
	retC := C.gtk_application_get_windows((*C.GtkApplication)(recv.native))
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Remove a window from @application.
//
// If @window belongs to @application then this call is equivalent to
// setting the #GtkWindow:application property of @window to
// %NULL.
//
// The application may stop running as a result of a call to this
// function.
/*

C function

gtk_application_remove_window
*/
func (recv *Application) RemoveWindow(window *Window) {
	c_window := (*C.GtkWindow)(C.NULL)
	if window != nil {
		c_window = (*C.GtkWindow)(window.ToC())
	}

	C.gtk_application_remove_window((*C.GtkApplication)(recv.native), c_window)

	return
}

// Navigate to the next page.
//
// It is a programming error to call this function when
// there is no next page.
//
// This function is for use when creating pages of the
// #GTK_ASSISTANT_PAGE_CUSTOM type.
/*

C function

gtk_assistant_next_page
*/
func (recv *Assistant) NextPage() {
	C.gtk_assistant_next_page((*C.GtkAssistant)(recv.native))

	return
}

// Navigate to the previous visited page.
//
// It is a programming error to call this function when
// no previous page is available.
//
// This function is for use when creating pages of the
// #GTK_ASSISTANT_PAGE_CUSTOM type.
/*

C function

gtk_assistant_previous_page
*/
func (recv *Assistant) PreviousPage() {
	C.gtk_assistant_previous_page((*C.GtkAssistant)(recv.native))

	return
}

// Creates a new #GtkBox.
/*

C function

gtk_box_new
*/
func BoxNew(orientation Orientation, spacing int32) *Box {
	c_orientation := (C.GtkOrientation)(orientation)

	c_spacing := (C.gint)(spacing)

	retC := C.gtk_box_new(c_orientation, c_spacing)
	retGo := BoxNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Creates a new #GtkButtonBox.
/*

C function

gtk_button_box_new
*/
func ButtonBoxNew(orientation Orientation) *ButtonBox {
	c_orientation := (C.GtkOrientation)(orientation)

	retC := C.gtk_button_box_new(c_orientation)
	retGo := ButtonBoxNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Returns if the @day of the @calendar is already marked.
/*

C function

gtk_calendar_get_day_is_marked
*/
func (recv *Calendar) GetDayIsMarked(day uint32) bool {
	c_day := (C.guint)(day)

	retC := C.gtk_calendar_get_day_is_marked((*C.GtkCalendar)(recv.native), c_day)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported signal 'add-editable' for CellArea : unsupported parameter cell_area : type Gdk.Rectangle : Blacklisted record : GdkRectangle

type signalCellAreaApplyAttributesDetail struct {
	callback  CellAreaSignalApplyAttributesCallback
	handlerID C.gulong
}

var signalCellAreaApplyAttributesId int
var signalCellAreaApplyAttributesMap = make(map[int]signalCellAreaApplyAttributesDetail)
var signalCellAreaApplyAttributesLock sync.Mutex

// CellAreaSignalApplyAttributesCallback is a callback function for a 'apply-attributes' signal emitted from a CellArea.
type CellAreaSignalApplyAttributesCallback func(model *TreeModel, iter *TreeIter, isExpander bool, isExpanded bool)

/*
ConnectApplyAttributes connects the callback to the 'apply-attributes' signal for the CellArea.

The returned value represents the connection, and may be passed to DisconnectApplyAttributes to remove it.
*/
func (recv *CellArea) ConnectApplyAttributes(callback CellAreaSignalApplyAttributesCallback) int {
	signalCellAreaApplyAttributesLock.Lock()
	defer signalCellAreaApplyAttributesLock.Unlock()

	signalCellAreaApplyAttributesId++
	instance := C.gpointer(recv.native)
	handlerID := C.CellArea_signal_connect_apply_attributes(instance, C.gpointer(uintptr(signalCellAreaApplyAttributesId)))

	detail := signalCellAreaApplyAttributesDetail{callback, handlerID}
	signalCellAreaApplyAttributesMap[signalCellAreaApplyAttributesId] = detail

	return signalCellAreaApplyAttributesId
}

/*
DisconnectApplyAttributes disconnects a callback from the 'apply-attributes' signal for the CellArea.

The connectionID should be a value returned from a call to ConnectApplyAttributes.
*/
func (recv *CellArea) DisconnectApplyAttributes(connectionID int) {
	signalCellAreaApplyAttributesLock.Lock()
	defer signalCellAreaApplyAttributesLock.Unlock()

	detail, exists := signalCellAreaApplyAttributesMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalCellAreaApplyAttributesMap, connectionID)
}

//export cellarea_applyAttributesHandler
func cellarea_applyAttributesHandler(_ *C.GObject, c_model *C.GtkTreeModel, c_iter *C.GtkTreeIter, c_is_expander C.gboolean, c_is_expanded C.gboolean, data C.gpointer) {
	model := TreeModelNewFromC(unsafe.Pointer(c_model))

	iter := TreeIterNewFromC(unsafe.Pointer(c_iter))

	isExpander := c_is_expander == C.TRUE

	isExpanded := c_is_expanded == C.TRUE

	index := int(uintptr(data))
	callback := signalCellAreaApplyAttributesMap[index].callback
	callback(model, iter, isExpander, isExpanded)
}

// Unsupported signal 'focus-changed' for CellArea : unsupported parameter path : type utf8 :

type signalCellAreaRemoveEditableDetail struct {
	callback  CellAreaSignalRemoveEditableCallback
	handlerID C.gulong
}

var signalCellAreaRemoveEditableId int
var signalCellAreaRemoveEditableMap = make(map[int]signalCellAreaRemoveEditableDetail)
var signalCellAreaRemoveEditableLock sync.Mutex

// CellAreaSignalRemoveEditableCallback is a callback function for a 'remove-editable' signal emitted from a CellArea.
type CellAreaSignalRemoveEditableCallback func(renderer *CellRenderer, editable *CellEditable)

/*
ConnectRemoveEditable connects the callback to the 'remove-editable' signal for the CellArea.

The returned value represents the connection, and may be passed to DisconnectRemoveEditable to remove it.
*/
func (recv *CellArea) ConnectRemoveEditable(callback CellAreaSignalRemoveEditableCallback) int {
	signalCellAreaRemoveEditableLock.Lock()
	defer signalCellAreaRemoveEditableLock.Unlock()

	signalCellAreaRemoveEditableId++
	instance := C.gpointer(recv.native)
	handlerID := C.CellArea_signal_connect_remove_editable(instance, C.gpointer(uintptr(signalCellAreaRemoveEditableId)))

	detail := signalCellAreaRemoveEditableDetail{callback, handlerID}
	signalCellAreaRemoveEditableMap[signalCellAreaRemoveEditableId] = detail

	return signalCellAreaRemoveEditableId
}

/*
DisconnectRemoveEditable disconnects a callback from the 'remove-editable' signal for the CellArea.

The connectionID should be a value returned from a call to ConnectRemoveEditable.
*/
func (recv *CellArea) DisconnectRemoveEditable(connectionID int) {
	signalCellAreaRemoveEditableLock.Lock()
	defer signalCellAreaRemoveEditableLock.Unlock()

	detail, exists := signalCellAreaRemoveEditableMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalCellAreaRemoveEditableMap, connectionID)
}

//export cellarea_removeEditableHandler
func cellarea_removeEditableHandler(_ *C.GObject, c_renderer *C.GtkCellRenderer, c_editable *C.GtkCellEditable, data C.gpointer) {
	renderer := CellRendererNewFromC(unsafe.Pointer(c_renderer))

	editable := CellEditableNewFromC(unsafe.Pointer(c_editable))

	index := int(uintptr(data))
	callback := signalCellAreaRemoveEditableMap[index].callback
	callback(renderer, editable)
}

// Unsupported : gtk_cell_area_activate : unsupported parameter cell_area : Blacklisted record : GdkRectangle

// Unsupported : gtk_cell_area_activate_cell : unsupported parameter event : no type generator for Gdk.Event (GdkEvent*) for param event

// Adds @renderer to @area with the default child cell properties.
/*

C function

gtk_cell_area_add
*/
func (recv *CellArea) Add(renderer *CellRenderer) {
	c_renderer := (*C.GtkCellRenderer)(C.NULL)
	if renderer != nil {
		c_renderer = (*C.GtkCellRenderer)(renderer.ToC())
	}

	C.gtk_cell_area_add((*C.GtkCellArea)(recv.native), c_renderer)

	return
}

// Adds @sibling to @renderer’s focusable area, focus will be drawn
// around @renderer and all of its siblings if @renderer can
// focus for a given row.
//
// Events handled by focus siblings can also activate the given
// focusable @renderer.
/*

C function

gtk_cell_area_add_focus_sibling
*/
func (recv *CellArea) AddFocusSibling(renderer *CellRenderer, sibling *CellRenderer) {
	c_renderer := (*C.GtkCellRenderer)(C.NULL)
	if renderer != nil {
		c_renderer = (*C.GtkCellRenderer)(renderer.ToC())
	}

	c_sibling := (*C.GtkCellRenderer)(C.NULL)
	if sibling != nil {
		c_sibling = (*C.GtkCellRenderer)(sibling.ToC())
	}

	C.gtk_cell_area_add_focus_sibling((*C.GtkCellArea)(recv.native), c_renderer, c_sibling)

	return
}

// Unsupported : gtk_cell_area_add_with_properties : unsupported parameter ... : varargs

// Applies any connected attributes to the renderers in
// @area by pulling the values from @tree_model.
/*

C function

gtk_cell_area_apply_attributes
*/
func (recv *CellArea) ApplyAttributes(treeModel *TreeModel, iter *TreeIter, isExpander bool, isExpanded bool) {
	c_tree_model := (*C.GtkTreeModel)(treeModel.ToC())

	c_iter := (*C.GtkTreeIter)(C.NULL)
	if iter != nil {
		c_iter = (*C.GtkTreeIter)(iter.ToC())
	}

	c_is_expander :=
		boolToGboolean(isExpander)

	c_is_expanded :=
		boolToGboolean(isExpanded)

	C.gtk_cell_area_apply_attributes((*C.GtkCellArea)(recv.native), c_tree_model, c_iter, c_is_expander, c_is_expanded)

	return
}

// Connects an @attribute to apply values from @column for the
// #GtkTreeModel in use.
/*

C function

gtk_cell_area_attribute_connect
*/
func (recv *CellArea) AttributeConnect(renderer *CellRenderer, attribute string, column int32) {
	c_renderer := (*C.GtkCellRenderer)(C.NULL)
	if renderer != nil {
		c_renderer = (*C.GtkCellRenderer)(renderer.ToC())
	}

	c_attribute := C.CString(attribute)
	defer C.free(unsafe.Pointer(c_attribute))

	c_column := (C.gint)(column)

	C.gtk_cell_area_attribute_connect((*C.GtkCellArea)(recv.native), c_renderer, c_attribute, c_column)

	return
}

// Disconnects @attribute for the @renderer in @area so that
// attribute will no longer be updated with values from the
// model.
/*

C function

gtk_cell_area_attribute_disconnect
*/
func (recv *CellArea) AttributeDisconnect(renderer *CellRenderer, attribute string) {
	c_renderer := (*C.GtkCellRenderer)(C.NULL)
	if renderer != nil {
		c_renderer = (*C.GtkCellRenderer)(renderer.ToC())
	}

	c_attribute := C.CString(attribute)
	defer C.free(unsafe.Pointer(c_attribute))

	C.gtk_cell_area_attribute_disconnect((*C.GtkCellArea)(recv.native), c_renderer, c_attribute)

	return
}

// Unsupported : gtk_cell_area_cell_get : unsupported parameter ... : varargs

// Gets the value of a cell property for @renderer in @area.
/*

C function

gtk_cell_area_cell_get_property
*/
func (recv *CellArea) CellGetProperty(renderer *CellRenderer, propertyName string, value *gobject.Value) {
	c_renderer := (*C.GtkCellRenderer)(C.NULL)
	if renderer != nil {
		c_renderer = (*C.GtkCellRenderer)(renderer.ToC())
	}

	c_property_name := C.CString(propertyName)
	defer C.free(unsafe.Pointer(c_property_name))

	c_value := (*C.GValue)(C.NULL)
	if value != nil {
		c_value = (*C.GValue)(value.ToC())
	}

	C.gtk_cell_area_cell_get_property((*C.GtkCellArea)(recv.native), c_renderer, c_property_name, c_value)

	return
}

// Unsupported : gtk_cell_area_cell_get_valist : unsupported parameter var_args : no type generator for va_list (va_list) for param var_args

// Unsupported : gtk_cell_area_cell_set : unsupported parameter ... : varargs

// Sets a cell property for @renderer in @area.
/*

C function

gtk_cell_area_cell_set_property
*/
func (recv *CellArea) CellSetProperty(renderer *CellRenderer, propertyName string, value *gobject.Value) {
	c_renderer := (*C.GtkCellRenderer)(C.NULL)
	if renderer != nil {
		c_renderer = (*C.GtkCellRenderer)(renderer.ToC())
	}

	c_property_name := C.CString(propertyName)
	defer C.free(unsafe.Pointer(c_property_name))

	c_value := (*C.GValue)(C.NULL)
	if value != nil {
		c_value = (*C.GValue)(value.ToC())
	}

	C.gtk_cell_area_cell_set_property((*C.GtkCellArea)(recv.native), c_renderer, c_property_name, c_value)

	return
}

// Unsupported : gtk_cell_area_cell_set_valist : unsupported parameter var_args : no type generator for va_list (va_list) for param var_args

// This is sometimes needed for cases where rows need to share
// alignments in one orientation but may be separately grouped
// in the opposing orientation.
//
// For instance, #GtkIconView creates all icons (rows) to have
// the same width and the cells theirin to have the same
// horizontal alignments. However each row of icons may have
// a separate collective height. #GtkIconView uses this to
// request the heights of each row based on a context which
// was already used to request all the row widths that are
// to be displayed.
/*

C function

gtk_cell_area_copy_context
*/
func (recv *CellArea) CopyContext(context *CellAreaContext) *CellAreaContext {
	c_context := (*C.GtkCellAreaContext)(C.NULL)
	if context != nil {
		c_context = (*C.GtkCellAreaContext)(context.ToC())
	}

	retC := C.gtk_cell_area_copy_context((*C.GtkCellArea)(recv.native), c_context)
	retGo := CellAreaContextNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Creates a #GtkCellAreaContext to be used with @area for
// all purposes. #GtkCellAreaContext stores geometry information
// for rows for which it was operated on, it is important to use
// the same context for the same row of data at all times (i.e.
// one should render and handle events with the same #GtkCellAreaContext
// which was used to request the size of those rows of data).
/*

C function

gtk_cell_area_create_context
*/
func (recv *CellArea) CreateContext() *CellAreaContext {
	retC := C.gtk_cell_area_create_context((*C.GtkCellArea)(recv.native))
	retGo := CellAreaContextNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_cell_area_event : unsupported parameter event : no type generator for Gdk.Event (GdkEvent*) for param event

// This should be called by the @area’s owning layout widget
// when focus is to be passed to @area, or moved within @area
// for a given @direction and row data.
//
// Implementing #GtkCellArea classes should implement this
// method to receive and navigate focus in its own way particular
// to how it lays out cells.
/*

C function

gtk_cell_area_focus
*/
func (recv *CellArea) Focus(direction DirectionType) bool {
	c_direction := (C.GtkDirectionType)(direction)

	retC := C.gtk_cell_area_focus((*C.GtkCellArea)(recv.native), c_direction)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_cell_area_foreach : unsupported parameter callback : no type generator for CellCallback (GtkCellCallback) for param callback

// Unsupported : gtk_cell_area_foreach_alloc : unsupported parameter cell_area : Blacklisted record : GdkRectangle

// Unsupported : gtk_cell_area_get_cell_allocation : unsupported parameter cell_area : Blacklisted record : GdkRectangle

// Unsupported : gtk_cell_area_get_cell_at_position : unsupported parameter cell_area : Blacklisted record : GdkRectangle

// Gets the current #GtkTreePath string for the currently
// applied #GtkTreeIter, this is implicitly updated when
// gtk_cell_area_apply_attributes() is called and can be
// used to interact with renderers from #GtkCellArea
// subclasses.
/*

C function

gtk_cell_area_get_current_path_string
*/
func (recv *CellArea) GetCurrentPathString() string {
	retC := C.gtk_cell_area_get_current_path_string((*C.GtkCellArea)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Gets the #GtkCellEditable widget currently used
// to edit the currently edited cell.
/*

C function

gtk_cell_area_get_edit_widget
*/
func (recv *CellArea) GetEditWidget() *CellEditable {
	retC := C.gtk_cell_area_get_edit_widget((*C.GtkCellArea)(recv.native))
	retGo := CellEditableNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the #GtkCellRenderer in @area that is currently
// being edited.
/*

C function

gtk_cell_area_get_edited_cell
*/
func (recv *CellArea) GetEditedCell() *CellRenderer {
	retC := C.gtk_cell_area_get_edited_cell((*C.GtkCellArea)(recv.native))
	retGo := CellRendererNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Retrieves the currently focused cell for @area
/*

C function

gtk_cell_area_get_focus_cell
*/
func (recv *CellArea) GetFocusCell() *CellRenderer {
	retC := C.gtk_cell_area_get_focus_cell((*C.GtkCellArea)(recv.native))
	retGo := CellRendererNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the #GtkCellRenderer which is expected to be focusable
// for which @renderer is, or may be a sibling.
//
// This is handy for #GtkCellArea subclasses when handling events,
// after determining the renderer at the event location it can
// then chose to activate the focus cell for which the event
// cell may have been a sibling.
/*

C function

gtk_cell_area_get_focus_from_sibling
*/
func (recv *CellArea) GetFocusFromSibling(renderer *CellRenderer) *CellRenderer {
	c_renderer := (*C.GtkCellRenderer)(C.NULL)
	if renderer != nil {
		c_renderer = (*C.GtkCellRenderer)(renderer.ToC())
	}

	retC := C.gtk_cell_area_get_focus_from_sibling((*C.GtkCellArea)(recv.native), c_renderer)
	var retGo (*CellRenderer)
	if retC == nil {
		retGo = nil
	} else {
		retGo = CellRendererNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Gets the focus sibling cell renderers for @renderer.
/*

C function

gtk_cell_area_get_focus_siblings
*/
func (recv *CellArea) GetFocusSiblings(renderer *CellRenderer) *glib.List {
	c_renderer := (*C.GtkCellRenderer)(C.NULL)
	if renderer != nil {
		c_renderer = (*C.GtkCellRenderer)(renderer.ToC())
	}

	retC := C.gtk_cell_area_get_focus_siblings((*C.GtkCellArea)(recv.native), c_renderer)
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Retrieves a cell area’s initial minimum and natural height.
//
// @area will store some geometrical information in @context along the way;
// when requesting sizes over an arbitrary number of rows, it’s not important
// to check the @minimum_height and @natural_height of this call but rather to
// consult gtk_cell_area_context_get_preferred_height() after a series of
// requests.
/*

C function

gtk_cell_area_get_preferred_height
*/
func (recv *CellArea) GetPreferredHeight(context *CellAreaContext, widget *Widget) (int32, int32) {
	c_context := (*C.GtkCellAreaContext)(C.NULL)
	if context != nil {
		c_context = (*C.GtkCellAreaContext)(context.ToC())
	}

	c_widget := (*C.GtkWidget)(C.NULL)
	if widget != nil {
		c_widget = (*C.GtkWidget)(widget.ToC())
	}

	var c_minimum_height C.gint

	var c_natural_height C.gint

	C.gtk_cell_area_get_preferred_height((*C.GtkCellArea)(recv.native), c_context, c_widget, &c_minimum_height, &c_natural_height)

	minimumHeight := (int32)(c_minimum_height)

	naturalHeight := (int32)(c_natural_height)

	return minimumHeight, naturalHeight
}

// Retrieves a cell area’s minimum and natural height if it would be given
// the specified @width.
//
// @area stores some geometrical information in @context along the way
// while calling gtk_cell_area_get_preferred_width(). It’s important to
// perform a series of gtk_cell_area_get_preferred_width() requests with
// @context first and then call gtk_cell_area_get_preferred_height_for_width()
// on each cell area individually to get the height for width of each
// fully requested row.
//
// If at some point, the width of a single row changes, it should be
// requested with gtk_cell_area_get_preferred_width() again and then
// the full width of the requested rows checked again with
// gtk_cell_area_context_get_preferred_width().
/*

C function

gtk_cell_area_get_preferred_height_for_width
*/
func (recv *CellArea) GetPreferredHeightForWidth(context *CellAreaContext, widget *Widget, width int32) (int32, int32) {
	c_context := (*C.GtkCellAreaContext)(C.NULL)
	if context != nil {
		c_context = (*C.GtkCellAreaContext)(context.ToC())
	}

	c_widget := (*C.GtkWidget)(C.NULL)
	if widget != nil {
		c_widget = (*C.GtkWidget)(widget.ToC())
	}

	c_width := (C.gint)(width)

	var c_minimum_height C.gint

	var c_natural_height C.gint

	C.gtk_cell_area_get_preferred_height_for_width((*C.GtkCellArea)(recv.native), c_context, c_widget, c_width, &c_minimum_height, &c_natural_height)

	minimumHeight := (int32)(c_minimum_height)

	naturalHeight := (int32)(c_natural_height)

	return minimumHeight, naturalHeight
}

// Retrieves a cell area’s initial minimum and natural width.
//
// @area will store some geometrical information in @context along the way;
// when requesting sizes over an arbitrary number of rows, it’s not important
// to check the @minimum_width and @natural_width of this call but rather to
// consult gtk_cell_area_context_get_preferred_width() after a series of
// requests.
/*

C function

gtk_cell_area_get_preferred_width
*/
func (recv *CellArea) GetPreferredWidth(context *CellAreaContext, widget *Widget) (int32, int32) {
	c_context := (*C.GtkCellAreaContext)(C.NULL)
	if context != nil {
		c_context = (*C.GtkCellAreaContext)(context.ToC())
	}

	c_widget := (*C.GtkWidget)(C.NULL)
	if widget != nil {
		c_widget = (*C.GtkWidget)(widget.ToC())
	}

	var c_minimum_width C.gint

	var c_natural_width C.gint

	C.gtk_cell_area_get_preferred_width((*C.GtkCellArea)(recv.native), c_context, c_widget, &c_minimum_width, &c_natural_width)

	minimumWidth := (int32)(c_minimum_width)

	naturalWidth := (int32)(c_natural_width)

	return minimumWidth, naturalWidth
}

// Retrieves a cell area’s minimum and natural width if it would be given
// the specified @height.
//
// @area stores some geometrical information in @context along the way
// while calling gtk_cell_area_get_preferred_height(). It’s important to
// perform a series of gtk_cell_area_get_preferred_height() requests with
// @context first and then call gtk_cell_area_get_preferred_width_for_height()
// on each cell area individually to get the height for width of each
// fully requested row.
//
// If at some point, the height of a single row changes, it should be
// requested with gtk_cell_area_get_preferred_height() again and then
// the full height of the requested rows checked again with
// gtk_cell_area_context_get_preferred_height().
/*

C function

gtk_cell_area_get_preferred_width_for_height
*/
func (recv *CellArea) GetPreferredWidthForHeight(context *CellAreaContext, widget *Widget, height int32) (int32, int32) {
	c_context := (*C.GtkCellAreaContext)(C.NULL)
	if context != nil {
		c_context = (*C.GtkCellAreaContext)(context.ToC())
	}

	c_widget := (*C.GtkWidget)(C.NULL)
	if widget != nil {
		c_widget = (*C.GtkWidget)(widget.ToC())
	}

	c_height := (C.gint)(height)

	var c_minimum_width C.gint

	var c_natural_width C.gint

	C.gtk_cell_area_get_preferred_width_for_height((*C.GtkCellArea)(recv.native), c_context, c_widget, c_height, &c_minimum_width, &c_natural_width)

	minimumWidth := (int32)(c_minimum_width)

	naturalWidth := (int32)(c_natural_width)

	return minimumWidth, naturalWidth
}

// Gets whether the area prefers a height-for-width layout
// or a width-for-height layout.
/*

C function

gtk_cell_area_get_request_mode
*/
func (recv *CellArea) GetRequestMode() SizeRequestMode {
	retC := C.gtk_cell_area_get_request_mode((*C.GtkCellArea)(recv.native))
	retGo := (SizeRequestMode)(retC)

	return retGo
}

// Checks if @area contains @renderer.
/*

C function

gtk_cell_area_has_renderer
*/
func (recv *CellArea) HasRenderer(renderer *CellRenderer) bool {
	c_renderer := (*C.GtkCellRenderer)(C.NULL)
	if renderer != nil {
		c_renderer = (*C.GtkCellRenderer)(renderer.ToC())
	}

	retC := C.gtk_cell_area_has_renderer((*C.GtkCellArea)(recv.native), c_renderer)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_cell_area_inner_cell_area : unsupported parameter cell_area : Blacklisted record : GdkRectangle

// Returns whether the area can do anything when activated,
// after applying new attributes to @area.
/*

C function

gtk_cell_area_is_activatable
*/
func (recv *CellArea) IsActivatable() bool {
	retC := C.gtk_cell_area_is_activatable((*C.GtkCellArea)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Returns whether @sibling is one of @renderer’s focus siblings
// (see gtk_cell_area_add_focus_sibling()).
/*

C function

gtk_cell_area_is_focus_sibling
*/
func (recv *CellArea) IsFocusSibling(renderer *CellRenderer, sibling *CellRenderer) bool {
	c_renderer := (*C.GtkCellRenderer)(C.NULL)
	if renderer != nil {
		c_renderer = (*C.GtkCellRenderer)(renderer.ToC())
	}

	c_sibling := (*C.GtkCellRenderer)(C.NULL)
	if sibling != nil {
		c_sibling = (*C.GtkCellRenderer)(sibling.ToC())
	}

	retC := C.gtk_cell_area_is_focus_sibling((*C.GtkCellArea)(recv.native), c_renderer, c_sibling)
	retGo := retC == C.TRUE

	return retGo
}

// Removes @renderer from @area.
/*

C function

gtk_cell_area_remove
*/
func (recv *CellArea) Remove(renderer *CellRenderer) {
	c_renderer := (*C.GtkCellRenderer)(C.NULL)
	if renderer != nil {
		c_renderer = (*C.GtkCellRenderer)(renderer.ToC())
	}

	C.gtk_cell_area_remove((*C.GtkCellArea)(recv.native), c_renderer)

	return
}

// Removes @sibling from @renderer’s focus sibling list
// (see gtk_cell_area_add_focus_sibling()).
/*

C function

gtk_cell_area_remove_focus_sibling
*/
func (recv *CellArea) RemoveFocusSibling(renderer *CellRenderer, sibling *CellRenderer) {
	c_renderer := (*C.GtkCellRenderer)(C.NULL)
	if renderer != nil {
		c_renderer = (*C.GtkCellRenderer)(renderer.ToC())
	}

	c_sibling := (*C.GtkCellRenderer)(C.NULL)
	if sibling != nil {
		c_sibling = (*C.GtkCellRenderer)(sibling.ToC())
	}

	C.gtk_cell_area_remove_focus_sibling((*C.GtkCellArea)(recv.native), c_renderer, c_sibling)

	return
}

// Unsupported : gtk_cell_area_render : unsupported parameter background_area : Blacklisted record : GdkRectangle

// This is a convenience function for #GtkCellArea implementations
// to request size for cell renderers. It’s important to use this
// function to request size and then use gtk_cell_area_inner_cell_area()
// at render and event time since this function will add padding
// around the cell for focus painting.
/*

C function

gtk_cell_area_request_renderer
*/
func (recv *CellArea) RequestRenderer(renderer *CellRenderer, orientation Orientation, widget *Widget, forSize int32) (int32, int32) {
	c_renderer := (*C.GtkCellRenderer)(C.NULL)
	if renderer != nil {
		c_renderer = (*C.GtkCellRenderer)(renderer.ToC())
	}

	c_orientation := (C.GtkOrientation)(orientation)

	c_widget := (*C.GtkWidget)(C.NULL)
	if widget != nil {
		c_widget = (*C.GtkWidget)(widget.ToC())
	}

	c_for_size := (C.gint)(forSize)

	var c_minimum_size C.gint

	var c_natural_size C.gint

	C.gtk_cell_area_request_renderer((*C.GtkCellArea)(recv.native), c_renderer, c_orientation, c_widget, c_for_size, &c_minimum_size, &c_natural_size)

	minimumSize := (int32)(c_minimum_size)

	naturalSize := (int32)(c_natural_size)

	return minimumSize, naturalSize
}

// Explicitly sets the currently focused cell to @renderer.
//
// This is generally called by implementations of
// #GtkCellAreaClass.focus() or #GtkCellAreaClass.event(),
// however it can also be used to implement functions such
// as gtk_tree_view_set_cursor_on_cell().
/*

C function

gtk_cell_area_set_focus_cell
*/
func (recv *CellArea) SetFocusCell(renderer *CellRenderer) {
	c_renderer := (*C.GtkCellRenderer)(C.NULL)
	if renderer != nil {
		c_renderer = (*C.GtkCellRenderer)(renderer.ToC())
	}

	C.gtk_cell_area_set_focus_cell((*C.GtkCellArea)(recv.native), c_renderer)

	return
}

// Explicitly stops the editing of the currently edited cell.
//
// If @canceled is %TRUE, the currently edited cell renderer
// will emit the ::editing-canceled signal, otherwise the
// the ::editing-done signal will be emitted on the current
// edit widget.
//
// See gtk_cell_area_get_edited_cell() and gtk_cell_area_get_edit_widget().
/*

C function

gtk_cell_area_stop_editing
*/
func (recv *CellArea) StopEditing(canceled bool) {
	c_canceled :=
		boolToGboolean(canceled)

	C.gtk_cell_area_stop_editing((*C.GtkCellArea)(recv.native), c_canceled)

	return
}

// Creates a new #GtkCellAreaBox.
/*

C function

gtk_cell_area_box_new
*/
func CellAreaBoxNew() *CellAreaBox {
	retC := C.gtk_cell_area_box_new()
	retGo := CellAreaBoxNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the spacing added between cell renderers.
/*

C function

gtk_cell_area_box_get_spacing
*/
func (recv *CellAreaBox) GetSpacing() int32 {
	retC := C.gtk_cell_area_box_get_spacing((*C.GtkCellAreaBox)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Adds @renderer to @box, packed with reference to the end of @box.
//
// The @renderer is packed after (away from end of) any other
// #GtkCellRenderer packed with reference to the end of @box.
/*

C function

gtk_cell_area_box_pack_end
*/
func (recv *CellAreaBox) PackEnd(renderer *CellRenderer, expand bool, align bool, fixed bool) {
	c_renderer := (*C.GtkCellRenderer)(C.NULL)
	if renderer != nil {
		c_renderer = (*C.GtkCellRenderer)(renderer.ToC())
	}

	c_expand :=
		boolToGboolean(expand)

	c_align :=
		boolToGboolean(align)

	c_fixed :=
		boolToGboolean(fixed)

	C.gtk_cell_area_box_pack_end((*C.GtkCellAreaBox)(recv.native), c_renderer, c_expand, c_align, c_fixed)

	return
}

// Adds @renderer to @box, packed with reference to the start of @box.
//
// The @renderer is packed after any other #GtkCellRenderer packed
// with reference to the start of @box.
/*

C function

gtk_cell_area_box_pack_start
*/
func (recv *CellAreaBox) PackStart(renderer *CellRenderer, expand bool, align bool, fixed bool) {
	c_renderer := (*C.GtkCellRenderer)(C.NULL)
	if renderer != nil {
		c_renderer = (*C.GtkCellRenderer)(renderer.ToC())
	}

	c_expand :=
		boolToGboolean(expand)

	c_align :=
		boolToGboolean(align)

	c_fixed :=
		boolToGboolean(fixed)

	C.gtk_cell_area_box_pack_start((*C.GtkCellAreaBox)(recv.native), c_renderer, c_expand, c_align, c_fixed)

	return
}

// Sets the spacing to add between cell renderers in @box.
/*

C function

gtk_cell_area_box_set_spacing
*/
func (recv *CellAreaBox) SetSpacing(spacing int32) {
	c_spacing := (C.gint)(spacing)

	C.gtk_cell_area_box_set_spacing((*C.GtkCellAreaBox)(recv.native), c_spacing)

	return
}

// Fetches the current allocation size for @context.
//
// If the context was not allocated in width or height, or if the
// context was recently reset with gtk_cell_area_context_reset(),
// the returned value will be -1.
/*

C function

gtk_cell_area_context_get_allocation
*/
func (recv *CellAreaContext) GetAllocation() (int32, int32) {
	var c_width C.gint

	var c_height C.gint

	C.gtk_cell_area_context_get_allocation((*C.GtkCellAreaContext)(recv.native), &c_width, &c_height)

	width := (int32)(c_width)

	height := (int32)(c_height)

	return width, height
}

// Fetches the #GtkCellArea this @context was created by.
//
// This is generally unneeded by layouting widgets; however,
// it is important for the context implementation itself to
// fetch information about the area it is being used for.
//
// For instance at #GtkCellAreaContextClass.allocate() time
// it’s important to know details about any cell spacing
// that the #GtkCellArea is configured with in order to
// compute a proper allocation.
/*

C function

gtk_cell_area_context_get_area
*/
func (recv *CellAreaContext) GetArea() *CellArea {
	retC := C.gtk_cell_area_context_get_area((*C.GtkCellAreaContext)(recv.native))
	retGo := CellAreaNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the accumulative preferred height for all rows which have been
// requested with this context.
//
// After gtk_cell_area_context_reset() is called and/or before ever
// requesting the size of a #GtkCellArea, the returned values are 0.
/*

C function

gtk_cell_area_context_get_preferred_height
*/
func (recv *CellAreaContext) GetPreferredHeight() (int32, int32) {
	var c_minimum_height C.gint

	var c_natural_height C.gint

	C.gtk_cell_area_context_get_preferred_height((*C.GtkCellAreaContext)(recv.native), &c_minimum_height, &c_natural_height)

	minimumHeight := (int32)(c_minimum_height)

	naturalHeight := (int32)(c_natural_height)

	return minimumHeight, naturalHeight
}

// Gets the accumulative preferred height for @width for all rows
// which have been requested for the same said @width with this context.
//
// After gtk_cell_area_context_reset() is called and/or before ever
// requesting the size of a #GtkCellArea, the returned values are -1.
/*

C function

gtk_cell_area_context_get_preferred_height_for_width
*/
func (recv *CellAreaContext) GetPreferredHeightForWidth(width int32) (int32, int32) {
	c_width := (C.gint)(width)

	var c_minimum_height C.gint

	var c_natural_height C.gint

	C.gtk_cell_area_context_get_preferred_height_for_width((*C.GtkCellAreaContext)(recv.native), c_width, &c_minimum_height, &c_natural_height)

	minimumHeight := (int32)(c_minimum_height)

	naturalHeight := (int32)(c_natural_height)

	return minimumHeight, naturalHeight
}

// Gets the accumulative preferred width for all rows which have been
// requested with this context.
//
// After gtk_cell_area_context_reset() is called and/or before ever
// requesting the size of a #GtkCellArea, the returned values are 0.
/*

C function

gtk_cell_area_context_get_preferred_width
*/
func (recv *CellAreaContext) GetPreferredWidth() (int32, int32) {
	var c_minimum_width C.gint

	var c_natural_width C.gint

	C.gtk_cell_area_context_get_preferred_width((*C.GtkCellAreaContext)(recv.native), &c_minimum_width, &c_natural_width)

	minimumWidth := (int32)(c_minimum_width)

	naturalWidth := (int32)(c_natural_width)

	return minimumWidth, naturalWidth
}

// Gets the accumulative preferred width for @height for all rows which
// have been requested for the same said @height with this context.
//
// After gtk_cell_area_context_reset() is called and/or before ever
// requesting the size of a #GtkCellArea, the returned values are -1.
/*

C function

gtk_cell_area_context_get_preferred_width_for_height
*/
func (recv *CellAreaContext) GetPreferredWidthForHeight(height int32) (int32, int32) {
	c_height := (C.gint)(height)

	var c_minimum_width C.gint

	var c_natural_width C.gint

	C.gtk_cell_area_context_get_preferred_width_for_height((*C.GtkCellAreaContext)(recv.native), c_height, &c_minimum_width, &c_natural_width)

	minimumWidth := (int32)(c_minimum_width)

	naturalWidth := (int32)(c_natural_width)

	return minimumWidth, naturalWidth
}

// Causes the minimum and/or natural height to grow if the new
// proposed sizes exceed the current minimum and natural height.
//
// This is used by #GtkCellAreaContext implementations during
// the request process over a series of #GtkTreeModel rows to
// progressively push the requested height over a series of
// gtk_cell_area_get_preferred_height() requests.
/*

C function

gtk_cell_area_context_push_preferred_height
*/
func (recv *CellAreaContext) PushPreferredHeight(minimumHeight int32, naturalHeight int32) {
	c_minimum_height := (C.gint)(minimumHeight)

	c_natural_height := (C.gint)(naturalHeight)

	C.gtk_cell_area_context_push_preferred_height((*C.GtkCellAreaContext)(recv.native), c_minimum_height, c_natural_height)

	return
}

// Causes the minimum and/or natural width to grow if the new
// proposed sizes exceed the current minimum and natural width.
//
// This is used by #GtkCellAreaContext implementations during
// the request process over a series of #GtkTreeModel rows to
// progressively push the requested width over a series of
// gtk_cell_area_get_preferred_width() requests.
/*

C function

gtk_cell_area_context_push_preferred_width
*/
func (recv *CellAreaContext) PushPreferredWidth(minimumWidth int32, naturalWidth int32) {
	c_minimum_width := (C.gint)(minimumWidth)

	c_natural_width := (C.gint)(naturalWidth)

	C.gtk_cell_area_context_push_preferred_width((*C.GtkCellAreaContext)(recv.native), c_minimum_width, c_natural_width)

	return
}

// Unsupported : gtk_cell_renderer_get_aligned_area : unsupported parameter cell_area : Blacklisted record : GdkRectangle

// Retreives a renderer’s natural size when rendered to @widget.
/*

C function

gtk_cell_renderer_get_preferred_height
*/
func (recv *CellRenderer) GetPreferredHeight(widget *Widget) (int32, int32) {
	c_widget := (*C.GtkWidget)(C.NULL)
	if widget != nil {
		c_widget = (*C.GtkWidget)(widget.ToC())
	}

	var c_minimum_size C.gint

	var c_natural_size C.gint

	C.gtk_cell_renderer_get_preferred_height((*C.GtkCellRenderer)(recv.native), c_widget, &c_minimum_size, &c_natural_size)

	minimumSize := (int32)(c_minimum_size)

	naturalSize := (int32)(c_natural_size)

	return minimumSize, naturalSize
}

// Retreives a cell renderers’s minimum and natural height if it were rendered to
// @widget with the specified @width.
/*

C function

gtk_cell_renderer_get_preferred_height_for_width
*/
func (recv *CellRenderer) GetPreferredHeightForWidth(widget *Widget, width int32) (int32, int32) {
	c_widget := (*C.GtkWidget)(C.NULL)
	if widget != nil {
		c_widget = (*C.GtkWidget)(widget.ToC())
	}

	c_width := (C.gint)(width)

	var c_minimum_height C.gint

	var c_natural_height C.gint

	C.gtk_cell_renderer_get_preferred_height_for_width((*C.GtkCellRenderer)(recv.native), c_widget, c_width, &c_minimum_height, &c_natural_height)

	minimumHeight := (int32)(c_minimum_height)

	naturalHeight := (int32)(c_natural_height)

	return minimumHeight, naturalHeight
}

// Retrieves the minimum and natural size of a cell taking
// into account the widget’s preference for height-for-width management.
/*

C function

gtk_cell_renderer_get_preferred_size
*/
func (recv *CellRenderer) GetPreferredSize(widget *Widget) (*Requisition, *Requisition) {
	c_widget := (*C.GtkWidget)(C.NULL)
	if widget != nil {
		c_widget = (*C.GtkWidget)(widget.ToC())
	}

	var c_minimum_size C.GtkRequisition

	var c_natural_size C.GtkRequisition

	C.gtk_cell_renderer_get_preferred_size((*C.GtkCellRenderer)(recv.native), c_widget, &c_minimum_size, &c_natural_size)

	minimumSize := RequisitionNewFromC(unsafe.Pointer(&c_minimum_size))

	naturalSize := RequisitionNewFromC(unsafe.Pointer(&c_natural_size))

	return minimumSize, naturalSize
}

// Retreives a renderer’s natural size when rendered to @widget.
/*

C function

gtk_cell_renderer_get_preferred_width
*/
func (recv *CellRenderer) GetPreferredWidth(widget *Widget) (int32, int32) {
	c_widget := (*C.GtkWidget)(C.NULL)
	if widget != nil {
		c_widget = (*C.GtkWidget)(widget.ToC())
	}

	var c_minimum_size C.gint

	var c_natural_size C.gint

	C.gtk_cell_renderer_get_preferred_width((*C.GtkCellRenderer)(recv.native), c_widget, &c_minimum_size, &c_natural_size)

	minimumSize := (int32)(c_minimum_size)

	naturalSize := (int32)(c_natural_size)

	return minimumSize, naturalSize
}

// Retreives a cell renderers’s minimum and natural width if it were rendered to
// @widget with the specified @height.
/*

C function

gtk_cell_renderer_get_preferred_width_for_height
*/
func (recv *CellRenderer) GetPreferredWidthForHeight(widget *Widget, height int32) (int32, int32) {
	c_widget := (*C.GtkWidget)(C.NULL)
	if widget != nil {
		c_widget = (*C.GtkWidget)(widget.ToC())
	}

	c_height := (C.gint)(height)

	var c_minimum_width C.gint

	var c_natural_width C.gint

	C.gtk_cell_renderer_get_preferred_width_for_height((*C.GtkCellRenderer)(recv.native), c_widget, c_height, &c_minimum_width, &c_natural_width)

	minimumWidth := (int32)(c_minimum_width)

	naturalWidth := (int32)(c_natural_width)

	return minimumWidth, naturalWidth
}

// Gets whether the cell renderer prefers a height-for-width layout
// or a width-for-height layout.
/*

C function

gtk_cell_renderer_get_request_mode
*/
func (recv *CellRenderer) GetRequestMode() SizeRequestMode {
	retC := C.gtk_cell_renderer_get_request_mode((*C.GtkCellRenderer)(recv.native))
	retGo := (SizeRequestMode)(retC)

	return retGo
}

// Translates the cell renderer state to #GtkStateFlags,
// based on the cell renderer and widget sensitivity, and
// the given #GtkCellRendererState.
/*

C function

gtk_cell_renderer_get_state
*/
func (recv *CellRenderer) GetState(widget *Widget, cellState CellRendererState) StateFlags {
	c_widget := (*C.GtkWidget)(C.NULL)
	if widget != nil {
		c_widget = (*C.GtkWidget)(widget.ToC())
	}

	c_cell_state := (C.GtkCellRendererState)(cellState)

	retC := C.gtk_cell_renderer_get_state((*C.GtkCellRenderer)(recv.native), c_widget, c_cell_state)
	retGo := (StateFlags)(retC)

	return retGo
}

// Checks whether the cell renderer can do something when activated.
/*

C function

gtk_cell_renderer_is_activatable
*/
func (recv *CellRenderer) IsActivatable() bool {
	retC := C.gtk_cell_renderer_is_activatable((*C.GtkCellRenderer)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Gets whether @cell_view is configured to draw all of its
// cells in a sensitive state.
/*

C function

gtk_cell_view_get_draw_sensitive
*/
func (recv *CellView) GetDrawSensitive() bool {
	retC := C.gtk_cell_view_get_draw_sensitive((*C.GtkCellView)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Gets whether @cell_view is configured to request space
// to fit the entire #GtkTreeModel.
/*

C function

gtk_cell_view_get_fit_model
*/
func (recv *CellView) GetFitModel() bool {
	retC := C.gtk_cell_view_get_fit_model((*C.GtkCellView)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Sets the background color of @cell_view.
/*

C function

gtk_cell_view_set_background_rgba
*/
func (recv *CellView) SetBackgroundRgba(rgba *gdk.RGBA) {
	c_rgba := (*C.GdkRGBA)(C.NULL)
	if rgba != nil {
		c_rgba = (*C.GdkRGBA)(rgba.ToC())
	}

	C.gtk_cell_view_set_background_rgba((*C.GtkCellView)(recv.native), c_rgba)

	return
}

// Sets whether @cell_view should draw all of its
// cells in a sensitive state, this is used by #GtkComboBox menus
// to ensure that rows with insensitive cells that contain
// children appear sensitive in the parent menu item.
/*

C function

gtk_cell_view_set_draw_sensitive
*/
func (recv *CellView) SetDrawSensitive(drawSensitive bool) {
	c_draw_sensitive :=
		boolToGboolean(drawSensitive)

	C.gtk_cell_view_set_draw_sensitive((*C.GtkCellView)(recv.native), c_draw_sensitive)

	return
}

// Sets whether @cell_view should request space to fit the entire #GtkTreeModel.
//
// This is used by #GtkComboBox to ensure that the cell view displayed on
// the combo box’s button always gets enough space and does not resize
// when selection changes.
/*

C function

gtk_cell_view_set_fit_model
*/
func (recv *CellView) SetFitModel(fitModel bool) {
	c_fit_model :=
		boolToGboolean(fitModel)

	C.gtk_cell_view_set_fit_model((*C.GtkCellView)(recv.native), c_fit_model)

	return
}

// Creates a new color button.
/*

C function

gtk_color_button_new_with_rgba
*/
func ColorButtonNewWithRgba(rgba *gdk.RGBA) *ColorButton {
	c_rgba := (*C.GdkRGBA)(C.NULL)
	if rgba != nil {
		c_rgba = (*C.GdkRGBA)(rgba.ToC())
	}

	retC := C.gtk_color_button_new_with_rgba(c_rgba)
	retGo := ColorButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Sets @rgba to be the current color in the #GtkColorButton widget.
/*

C function

gtk_color_button_get_rgba
*/
func (recv *ColorButton) GetRgba() *gdk.RGBA {
	var c_rgba C.GdkRGBA

	C.gtk_color_button_get_rgba((*C.GtkColorButton)(recv.native), &c_rgba)

	rgba := gdk.RGBANewFromC(unsafe.Pointer(&c_rgba))

	return rgba
}

// Sets the current color to be @rgba.
/*

C function

gtk_color_button_set_rgba
*/
func (recv *ColorButton) SetRgba(rgba *gdk.RGBA) {
	c_rgba := (*C.GdkRGBA)(C.NULL)
	if rgba != nil {
		c_rgba = (*C.GdkRGBA)(rgba.ToC())
	}

	C.gtk_color_button_set_rgba((*C.GtkColorButton)(recv.native), c_rgba)

	return
}

// Sets @rgba to be the current color in the GtkColorSelection widget.
/*

C function

gtk_color_selection_get_current_rgba
*/
func (recv *ColorSelection) GetCurrentRgba() *gdk.RGBA {
	var c_rgba C.GdkRGBA

	C.gtk_color_selection_get_current_rgba((*C.GtkColorSelection)(recv.native), &c_rgba)

	rgba := gdk.RGBANewFromC(unsafe.Pointer(&c_rgba))

	return rgba
}

// Fills @rgba in with the original color value.
/*

C function

gtk_color_selection_get_previous_rgba
*/
func (recv *ColorSelection) GetPreviousRgba() *gdk.RGBA {
	var c_rgba C.GdkRGBA

	C.gtk_color_selection_get_previous_rgba((*C.GtkColorSelection)(recv.native), &c_rgba)

	rgba := gdk.RGBANewFromC(unsafe.Pointer(&c_rgba))

	return rgba
}

// Sets the current color to be @rgba.
//
// The first time this is called, it will also set
// the original color to be @rgba too.
/*

C function

gtk_color_selection_set_current_rgba
*/
func (recv *ColorSelection) SetCurrentRgba(rgba *gdk.RGBA) {
	c_rgba := (*C.GdkRGBA)(C.NULL)
	if rgba != nil {
		c_rgba = (*C.GdkRGBA)(rgba.ToC())
	}

	C.gtk_color_selection_set_current_rgba((*C.GtkColorSelection)(recv.native), c_rgba)

	return
}

// Sets the “previous” color to be @rgba.
//
// This function should be called with some hesitations,
// as it might seem confusing to have that color change.
// Calling gtk_color_selection_set_current_rgba() will also
// set this color the first time it is called.
/*

C function

gtk_color_selection_set_previous_rgba
*/
func (recv *ColorSelection) SetPreviousRgba(rgba *gdk.RGBA) {
	c_rgba := (*C.GdkRGBA)(C.NULL)
	if rgba != nil {
		c_rgba = (*C.GdkRGBA)(rgba.ToC())
	}

	C.gtk_color_selection_set_previous_rgba((*C.GtkColorSelection)(recv.native), c_rgba)

	return
}

// Returns the ID of the active row of @combo_box.  This value is taken
// from the active row and the column specified by the #GtkComboBox:id-column
// property of @combo_box (see gtk_combo_box_set_id_column()).
//
// The returned value is an interned string which means that you can
// compare the pointer by value to other interned strings and that you
// must not free it.
//
// If the #GtkComboBox:id-column property of @combo_box is not set, or if
// no row is active, or if the active row has a %NULL ID value, then %NULL
// is returned.
/*

C function

gtk_combo_box_get_active_id
*/
func (recv *ComboBox) GetActiveId() string {
	retC := C.gtk_combo_box_get_active_id((*C.GtkComboBox)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Returns the column which @combo_box is using to get string IDs
// for values from.
/*

C function

gtk_combo_box_get_id_column
*/
func (recv *ComboBox) GetIdColumn() int32 {
	retC := C.gtk_combo_box_get_id_column((*C.GtkComboBox)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Gets whether the popup uses a fixed width matching
// the allocated width of the combo box.
/*

C function

gtk_combo_box_get_popup_fixed_width
*/
func (recv *ComboBox) GetPopupFixedWidth() bool {
	retC := C.gtk_combo_box_get_popup_fixed_width((*C.GtkComboBox)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Pops up the menu or dropdown list of @combo_box, the popup window
// will be grabbed so only @device and its associated pointer/keyboard
// are the only #GdkDevices able to send events to it.
/*

C function

gtk_combo_box_popup_for_device
*/
func (recv *ComboBox) PopupForDevice(device *gdk.Device) {
	c_device := (*C.GdkDevice)(C.NULL)
	if device != nil {
		c_device = (*C.GdkDevice)(device.ToC())
	}

	C.gtk_combo_box_popup_for_device((*C.GtkComboBox)(recv.native), c_device)

	return
}

// Changes the active row of @combo_box to the one that has an ID equal to
// @active_id, or unsets the active row if @active_id is %NULL.  Rows having
// a %NULL ID string cannot be made active by this function.
//
// If the #GtkComboBox:id-column property of @combo_box is unset or if no
// row has the given ID then the function does nothing and returns %FALSE.
/*

C function

gtk_combo_box_set_active_id
*/
func (recv *ComboBox) SetActiveId(activeId string) bool {
	c_active_id := C.CString(activeId)
	defer C.free(unsafe.Pointer(c_active_id))

	retC := C.gtk_combo_box_set_active_id((*C.GtkComboBox)(recv.native), c_active_id)
	retGo := retC == C.TRUE

	return retGo
}

// Sets the model column which @combo_box should use to get string IDs
// for values from. The column @id_column in the model of @combo_box
// must be of type %G_TYPE_STRING.
/*

C function

gtk_combo_box_set_id_column
*/
func (recv *ComboBox) SetIdColumn(idColumn int32) {
	c_id_column := (C.gint)(idColumn)

	C.gtk_combo_box_set_id_column((*C.GtkComboBox)(recv.native), c_id_column)

	return
}

// Specifies whether the popup’s width should be a fixed width
// matching the allocated width of the combo box.
/*

C function

gtk_combo_box_set_popup_fixed_width
*/
func (recv *ComboBox) SetPopupFixedWidth(fixed bool) {
	c_fixed :=
		boolToGboolean(fixed)

	C.gtk_combo_box_set_popup_fixed_width((*C.GtkComboBox)(recv.native), c_fixed)

	return
}

// Inserts @text at @position in the list of strings stored in @combo_box.
// If @id is non-%NULL then it is used as the ID of the row.  See
// #GtkComboBox:id-column.
//
// If @position is negative then @text is appended.
/*

C function

gtk_combo_box_text_insert
*/
func (recv *ComboBoxText) Insert(position int32, id string, text string) {
	c_position := (C.gint)(position)

	c_id := C.CString(id)
	defer C.free(unsafe.Pointer(c_id))

	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))

	C.gtk_combo_box_text_insert((*C.GtkComboBoxText)(recv.native), c_position, c_id, c_text)

	return
}

// Removes all the text entries from the combo box.
/*

C function

gtk_combo_box_text_remove_all
*/
func (recv *ComboBoxText) RemoveAll() {
	C.gtk_combo_box_text_remove_all((*C.GtkComboBoxText)(recv.native))

	return
}

// Unsupported : gtk_entry_get_icon_area : unsupported parameter icon_area : Blacklisted record : GdkRectangle

// Unsupported : gtk_entry_get_text_area : unsupported parameter text_area : Blacklisted record : GdkRectangle

// Creates a new #GtkEntryCompletion object using the
// specified @area to layout cells in the underlying
// #GtkTreeViewColumn for the drop-down menu.
/*

C function

gtk_entry_completion_new_with_area
*/
func EntryCompletionNewWithArea(area *CellArea) *EntryCompletion {
	c_area := (*C.GtkCellArea)(C.NULL)
	if area != nil {
		c_area = (*C.GtkCellArea)(area.ToC())
	}

	retC := C.gtk_entry_completion_new_with_area(c_area)
	retGo := EntryCompletionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Loads an icon, modifying it to match the system colours for the foreground,
// success, warning and error colors provided. If the icon is not a symbolic
// one, the function will return the result from gtk_icon_info_load_icon().
//
// This allows loading symbolic icons that will match the system theme.
//
// Unless you are implementing a widget, you will want to use
// g_themed_icon_new_with_default_fallbacks() to load the icon.
//
// As implementation details, the icon loaded needs to be of SVG type,
// contain the “symbolic” term as the last component of the icon name,
// and use the “fg”, “success”, “warning” and “error” CSS styles in the
// SVG file itself.
//
// See the [Symbolic Icons Specification](http://www.freedesktop.org/wiki/SymbolicIcons)
// for more information about symbolic icons.
/*

C function

gtk_icon_info_load_symbolic
*/
func (recv *IconInfo) LoadSymbolic(fg *gdk.RGBA, successColor *gdk.RGBA, warningColor *gdk.RGBA, errorColor *gdk.RGBA) (*gdkpixbuf.Pixbuf, bool, error) {
	c_fg := (*C.GdkRGBA)(C.NULL)
	if fg != nil {
		c_fg = (*C.GdkRGBA)(fg.ToC())
	}

	c_success_color := (*C.GdkRGBA)(C.NULL)
	if successColor != nil {
		c_success_color = (*C.GdkRGBA)(successColor.ToC())
	}

	c_warning_color := (*C.GdkRGBA)(C.NULL)
	if warningColor != nil {
		c_warning_color = (*C.GdkRGBA)(warningColor.ToC())
	}

	c_error_color := (*C.GdkRGBA)(C.NULL)
	if errorColor != nil {
		c_error_color = (*C.GdkRGBA)(errorColor.ToC())
	}

	var c_was_symbolic C.gboolean

	var cThrowableError *C.GError

	retC := C.gtk_icon_info_load_symbolic((*C.GtkIconInfo)(recv.native), c_fg, c_success_color, c_warning_color, c_error_color, &c_was_symbolic, &cThrowableError)
	retGo := gdkpixbuf.PixbufNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	wasSymbolic := c_was_symbolic == C.TRUE

	return retGo, wasSymbolic, goThrowableError
}

// Loads an icon, modifying it to match the system colors for the foreground,
// success, warning and error colors provided. If the icon is not a symbolic
// one, the function will return the result from gtk_icon_info_load_icon().
// This function uses the regular foreground color and the symbolic colors
// with the names “success_color”, “warning_color” and “error_color” from
// the context.
//
// This allows loading symbolic icons that will match the system theme.
//
// See gtk_icon_info_load_symbolic() for more details.
/*

C function

gtk_icon_info_load_symbolic_for_context
*/
func (recv *IconInfo) LoadSymbolicForContext(context *StyleContext) (*gdkpixbuf.Pixbuf, bool, error) {
	c_context := (*C.GtkStyleContext)(C.NULL)
	if context != nil {
		c_context = (*C.GtkStyleContext)(context.ToC())
	}

	var c_was_symbolic C.gboolean

	var cThrowableError *C.GError

	retC := C.gtk_icon_info_load_symbolic_for_context((*C.GtkIconInfo)(recv.native), c_context, &c_was_symbolic, &cThrowableError)
	retGo := gdkpixbuf.PixbufNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	wasSymbolic := c_was_symbolic == C.TRUE

	return retGo, wasSymbolic, goThrowableError
}

// Loads an icon, modifying it to match the system colours for the foreground,
// success, warning and error colors provided. If the icon is not a symbolic
// one, the function will return the result from gtk_icon_info_load_icon().
//
// This allows loading symbolic icons that will match the system theme.
//
// See gtk_icon_info_load_symbolic() for more details.
/*

C function

gtk_icon_info_load_symbolic_for_style
*/
func (recv *IconInfo) LoadSymbolicForStyle(style *Style, state StateType) (*gdkpixbuf.Pixbuf, bool, error) {
	c_style := (*C.GtkStyle)(C.NULL)
	if style != nil {
		c_style = (*C.GtkStyle)(style.ToC())
	}

	c_state := (C.GtkStateType)(state)

	var c_was_symbolic C.gboolean

	var cThrowableError *C.GError

	retC := C.gtk_icon_info_load_symbolic_for_style((*C.GtkIconInfo)(recv.native), c_style, c_state, &c_was_symbolic, &cThrowableError)
	retGo := gdkpixbuf.PixbufNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	wasSymbolic := c_was_symbolic == C.TRUE

	return retGo, wasSymbolic, goThrowableError
}

// Creates a new #GtkIconView widget using the
// specified @area to layout cells inside the icons.
/*

C function

gtk_icon_view_new_with_area
*/
func IconViewNewWithArea(area *CellArea) *IconView {
	c_area := (*C.GtkCellArea)(C.NULL)
	if area != nil {
		c_area = (*C.GtkCellArea)(area.ToC())
	}

	retC := C.gtk_icon_view_new_with_area(c_area)
	retGo := IconViewNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_menu_popup_for_device : unsupported parameter func : no type generator for MenuPositionFunc (GtkMenuPositionFunc) for param func

// Returns whether the @menu_item reserves space for
// the submenu indicator, regardless if it has a submenu
// or not.
/*

C function

gtk_menu_item_get_reserve_indicator
*/
func (recv *MenuItem) GetReserveIndicator() bool {
	retC := C.gtk_menu_item_get_reserve_indicator((*C.GtkMenuItem)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Sets whether the @menu_item should reserve space for
// the submenu indicator, regardless if it actually has
// a submenu or not.
//
// There should be little need for applications to call
// this functions.
/*

C function

gtk_menu_item_set_reserve_indicator
*/
func (recv *MenuItem) SetReserveIndicator(reserve bool) {
	c_reserve :=
		boolToGboolean(reserve)

	C.gtk_menu_item_set_reserve_indicator((*C.GtkMenuItem)(recv.native), c_reserve)

	return
}

// Gets the parent menu shell.
//
// The parent menu shell of a submenu is the #GtkMenu or #GtkMenuBar
// from which it was opened up.
/*

C function

gtk_menu_shell_get_parent_shell
*/
func (recv *MenuShell) GetParentShell() *Widget {
	retC := C.gtk_menu_shell_get_parent_shell((*C.GtkMenuShell)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the currently selected item.
/*

C function

gtk_menu_shell_get_selected_item
*/
func (recv *MenuShell) GetSelectedItem() *Widget {
	retC := C.gtk_menu_shell_get_selected_item((*C.GtkMenuShell)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Returns the #GIcon that was set as the base background image, or
// %NULL if there’s none. The caller of this function does not own
// a reference to the returned #GIcon.
/*

C function

gtk_numerable_icon_get_background_gicon
*/
func (recv *NumerableIcon) GetBackgroundGicon() *gio.Icon {
	retC := C.gtk_numerable_icon_get_background_gicon((*C.GtkNumerableIcon)(recv.native))
	var retGo (*gio.Icon)
	if retC == nil {
		retGo = nil
	} else {
		retGo = gio.IconNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Returns the icon name used as the base background image,
// or %NULL if there’s none.
/*

C function

gtk_numerable_icon_get_background_icon_name
*/
func (recv *NumerableIcon) GetBackgroundIconName() string {
	retC := C.gtk_numerable_icon_get_background_icon_name((*C.GtkNumerableIcon)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Returns the value currently displayed by @self.
/*

C function

gtk_numerable_icon_get_count
*/
func (recv *NumerableIcon) GetCount() int32 {
	retC := C.gtk_numerable_icon_get_count((*C.GtkNumerableIcon)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Returns the currently displayed label of the icon, or %NULL.
/*

C function

gtk_numerable_icon_get_label
*/
func (recv *NumerableIcon) GetLabel() string {
	retC := C.gtk_numerable_icon_get_label((*C.GtkNumerableIcon)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Returns the #GtkStyleContext used by the icon for theming,
// or %NULL if there’s none.
/*

C function

gtk_numerable_icon_get_style_context
*/
func (recv *NumerableIcon) GetStyleContext() *StyleContext {
	retC := C.gtk_numerable_icon_get_style_context((*C.GtkNumerableIcon)(recv.native))
	var retGo (*StyleContext)
	if retC == nil {
		retGo = nil
	} else {
		retGo = StyleContextNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Updates the icon to use @icon as the base background image.
// If @icon is %NULL, @self will go back using style information
// or default theming for its background image.
//
// If this method is called and an icon name was already set as
// background for the icon, @icon will be used, i.e. the last method
// called between gtk_numerable_icon_set_background_gicon() and
// gtk_numerable_icon_set_background_icon_name() has always priority.
/*

C function

gtk_numerable_icon_set_background_gicon
*/
func (recv *NumerableIcon) SetBackgroundGicon(icon *gio.Icon) {
	c_icon := (*C.GIcon)(icon.ToC())

	C.gtk_numerable_icon_set_background_gicon((*C.GtkNumerableIcon)(recv.native), c_icon)

	return
}

// Updates the icon to use the icon named @icon_name from the
// current icon theme as the base background image. If @icon_name
// is %NULL, @self will go back using style information or default
// theming for its background image.
//
// If this method is called and a #GIcon was already set as
// background for the icon, @icon_name will be used, i.e. the
// last method called between gtk_numerable_icon_set_background_icon_name()
// and gtk_numerable_icon_set_background_gicon() has always priority.
/*

C function

gtk_numerable_icon_set_background_icon_name
*/
func (recv *NumerableIcon) SetBackgroundIconName(iconName string) {
	c_icon_name := C.CString(iconName)
	defer C.free(unsafe.Pointer(c_icon_name))

	C.gtk_numerable_icon_set_background_icon_name((*C.GtkNumerableIcon)(recv.native), c_icon_name)

	return
}

// Sets the currently displayed value of @self to @count.
//
// The numeric value is always clamped to make it two digits, i.e.
// between -99 and 99. Setting a count of zero removes the emblem.
// If this method is called, and a label was already set on the icon,
// it will automatically be reset to %NULL before rendering the number,
// i.e. the last method called between gtk_numerable_icon_set_count()
// and gtk_numerable_icon_set_label() has always priority.
/*

C function

gtk_numerable_icon_set_count
*/
func (recv *NumerableIcon) SetCount(count int32) {
	c_count := (C.gint)(count)

	C.gtk_numerable_icon_set_count((*C.GtkNumerableIcon)(recv.native), c_count)

	return
}

// Sets the currently displayed value of @self to the string
// in @label. Setting an empty label removes the emblem.
//
// Note that this is meant for displaying short labels, such as
// roman numbers, or single letters. For roman numbers, consider
// using the Unicode characters U+2160 - U+217F. Strings longer
// than two characters will likely not be rendered very well.
//
// If this method is called, and a number was already set on the
// icon, it will automatically be reset to zero before rendering
// the label, i.e. the last method called between
// gtk_numerable_icon_set_label() and gtk_numerable_icon_set_count()
// has always priority.
/*

C function

gtk_numerable_icon_set_label
*/
func (recv *NumerableIcon) SetLabel(label string) {
	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	C.gtk_numerable_icon_set_label((*C.GtkNumerableIcon)(recv.native), c_label)

	return
}

// Updates the icon to fetch theme information from the
// given #GtkStyleContext.
/*

C function

gtk_numerable_icon_set_style_context
*/
func (recv *NumerableIcon) SetStyleContext(style *StyleContext) {
	c_style := (*C.GtkStyleContext)(C.NULL)
	if style != nil {
		c_style = (*C.GtkStyleContext)(style.ToC())
	}

	C.gtk_numerable_icon_set_style_context((*C.GtkNumerableIcon)(recv.native), c_style)

	return
}

// Creates a new #GtkPaned widget.
/*

C function

gtk_paned_new
*/
func PanedNew(orientation Orientation) *Paned {
	c_orientation := (C.GtkOrientation)(orientation)

	retC := C.gtk_paned_new(c_orientation)
	retGo := PanedNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the value of the #GtkProgressBar:show-text property.
// See gtk_progress_bar_set_show_text().
/*

C function

gtk_progress_bar_get_show_text
*/
func (recv *ProgressBar) GetShowText() bool {
	retC := C.gtk_progress_bar_get_show_text((*C.GtkProgressBar)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Sets whether the progress bar will show text next to the bar.
// The shown text is either the value of the #GtkProgressBar:text
// property or, if that is %NULL, the #GtkProgressBar:fraction value,
// as a percentage.
//
// To make a progress bar that is styled and sized suitably for containing
// text (even if the actual text is blank), set #GtkProgressBar:show-text to
// %TRUE and #GtkProgressBar:text to the empty string (not %NULL).
/*

C function

gtk_progress_bar_set_show_text
*/
func (recv *ProgressBar) SetShowText(showText bool) {
	c_show_text :=
		boolToGboolean(showText)

	C.gtk_progress_bar_set_show_text((*C.GtkProgressBar)(recv.native), c_show_text)

	return
}

// Joins a radio action object to the group of another radio action object.
//
// Use this in language bindings instead of the gtk_radio_action_get_group()
// and gtk_radio_action_set_group() methods
//
// A common way to set up a group of radio actions is the following:
// |[<!-- language="C" -->
// GtkRadioAction *action;
// GtkRadioAction *last_action;
//
// while ( ...more actions to add... /)
// {
// action = gtk_radio_action_new (...);
//
// gtk_radio_action_join_group (action, last_action);
// last_action = action;
// }
// ]|
/*

C function

gtk_radio_action_join_group
*/
func (recv *RadioAction) JoinGroup(groupSource *RadioAction) {
	c_group_source := (*C.GtkRadioAction)(C.NULL)
	if groupSource != nil {
		c_group_source = (*C.GtkRadioAction)(groupSource.ToC())
	}

	C.gtk_radio_action_join_group((*C.GtkRadioAction)(recv.native), c_group_source)

	return
}

// Joins a #GtkRadioButton object to the group of another #GtkRadioButton object
//
// Use this in language bindings instead of the gtk_radio_button_get_group()
// and gtk_radio_button_set_group() methods
//
// A common way to set up a group of radio buttons is the following:
// |[<!-- language="C" -->
// GtkRadioButton *radio_button;
// GtkRadioButton *last_button;
//
// while (some_condition)
// {
// radio_button = gtk_radio_button_new (NULL);
//
// gtk_radio_button_join_group (radio_button, last_button);
// last_button = radio_button;
// }
// ]|
/*

C function

gtk_radio_button_join_group
*/
func (recv *RadioButton) JoinGroup(groupSource *RadioButton) {
	c_group_source := (*C.GtkRadioButton)(C.NULL)
	if groupSource != nil {
		c_group_source = (*C.GtkRadioButton)(groupSource.ToC())
	}

	C.gtk_radio_button_join_group((*C.GtkRadioButton)(recv.native), c_group_source)

	return
}

// Creates a new #GtkScale.
/*

C function

gtk_scale_new
*/
func ScaleNew(orientation Orientation, adjustment *Adjustment) *Scale {
	c_orientation := (C.GtkOrientation)(orientation)

	c_adjustment := (*C.GtkAdjustment)(C.NULL)
	if adjustment != nil {
		c_adjustment = (*C.GtkAdjustment)(adjustment.ToC())
	}

	retC := C.gtk_scale_new(c_orientation, c_adjustment)
	retGo := ScaleNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Creates a new scale widget with the given orientation that lets the
// user input a number between @min and @max (including @min and @max)
// with the increment @step.  @step must be nonzero; it’s the distance
// the slider moves when using the arrow keys to adjust the scale
// value.
//
// Note that the way in which the precision is derived works best if @step
// is a power of ten. If the resulting precision is not suitable for your
// needs, use gtk_scale_set_digits() to correct it.
/*

C function

gtk_scale_new_with_range
*/
func ScaleNewWithRange(orientation Orientation, min float64, max float64, step float64) *Scale {
	c_orientation := (C.GtkOrientation)(orientation)

	c_min := (C.gdouble)(min)

	c_max := (C.gdouble)(max)

	c_step := (C.gdouble)(step)

	retC := C.gtk_scale_new_with_range(c_orientation, c_min, c_max, c_step)
	retGo := ScaleNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Creates a new scrollbar with the given orientation.
/*

C function

gtk_scrollbar_new
*/
func ScrollbarNew(orientation Orientation, adjustment *Adjustment) *Scrollbar {
	c_orientation := (C.GtkOrientation)(orientation)

	c_adjustment := (*C.GtkAdjustment)(C.NULL)
	if adjustment != nil {
		c_adjustment = (*C.GtkAdjustment)(adjustment.ToC())
	}

	retC := C.gtk_scrollbar_new(c_orientation, c_adjustment)
	retGo := ScrollbarNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the minimal content height of @scrolled_window, or -1 if not set.
/*

C function

gtk_scrolled_window_get_min_content_height
*/
func (recv *ScrolledWindow) GetMinContentHeight() int32 {
	retC := C.gtk_scrolled_window_get_min_content_height((*C.GtkScrolledWindow)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Gets the minimum content width of @scrolled_window, or -1 if not set.
/*

C function

gtk_scrolled_window_get_min_content_width
*/
func (recv *ScrolledWindow) GetMinContentWidth() int32 {
	retC := C.gtk_scrolled_window_get_min_content_width((*C.GtkScrolledWindow)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Sets the minimum height that @scrolled_window should keep visible.
// Note that this can and (usually will) be smaller than the minimum
// size of the content.
//
// It is a programming error to set the minimum content height to a
// value greater than #GtkScrolledWindow:max-content-height.
/*

C function

gtk_scrolled_window_set_min_content_height
*/
func (recv *ScrolledWindow) SetMinContentHeight(height int32) {
	c_height := (C.gint)(height)

	C.gtk_scrolled_window_set_min_content_height((*C.GtkScrolledWindow)(recv.native), c_height)

	return
}

// Sets the minimum width that @scrolled_window should keep visible.
// Note that this can and (usually will) be smaller than the minimum
// size of the content.
//
// It is a programming error to set the minimum content width to a
// value greater than #GtkScrolledWindow:max-content-width.
/*

C function

gtk_scrolled_window_set_min_content_width
*/
func (recv *ScrolledWindow) SetMinContentWidth(width int32) {
	c_width := (C.gint)(width)

	C.gtk_scrolled_window_set_min_content_width((*C.GtkScrolledWindow)(recv.native), c_width)

	return
}

// Creates a new #GtkSeparator with the given orientation.
/*

C function

gtk_separator_new
*/
func SeparatorNew(orientation Orientation) *Separator {
	c_orientation := (C.GtkOrientation)(orientation)

	retC := C.gtk_separator_new(c_orientation)
	retGo := SeparatorNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Returns whether @style has an associated #GtkStyleContext.
/*

C function

gtk_style_has_context
*/
func (recv *Style) HasContext() bool {
	retC := C.gtk_style_has_context((*C.GtkStyle)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

type signalStyleContextChangedDetail struct {
	callback  StyleContextSignalChangedCallback
	handlerID C.gulong
}

var signalStyleContextChangedId int
var signalStyleContextChangedMap = make(map[int]signalStyleContextChangedDetail)
var signalStyleContextChangedLock sync.Mutex

// StyleContextSignalChangedCallback is a callback function for a 'changed' signal emitted from a StyleContext.
type StyleContextSignalChangedCallback func()

/*
ConnectChanged connects the callback to the 'changed' signal for the StyleContext.

The returned value represents the connection, and may be passed to DisconnectChanged to remove it.
*/
func (recv *StyleContext) ConnectChanged(callback StyleContextSignalChangedCallback) int {
	signalStyleContextChangedLock.Lock()
	defer signalStyleContextChangedLock.Unlock()

	signalStyleContextChangedId++
	instance := C.gpointer(recv.native)
	handlerID := C.StyleContext_signal_connect_changed(instance, C.gpointer(uintptr(signalStyleContextChangedId)))

	detail := signalStyleContextChangedDetail{callback, handlerID}
	signalStyleContextChangedMap[signalStyleContextChangedId] = detail

	return signalStyleContextChangedId
}

/*
DisconnectChanged disconnects a callback from the 'changed' signal for the StyleContext.

The connectionID should be a value returned from a call to ConnectChanged.
*/
func (recv *StyleContext) DisconnectChanged(connectionID int) {
	signalStyleContextChangedLock.Lock()
	defer signalStyleContextChangedLock.Unlock()

	detail, exists := signalStyleContextChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalStyleContextChangedMap, connectionID)
}

//export stylecontext_changedHandler
func stylecontext_changedHandler(_ *C.GObject, data C.gpointer) {
	index := int(uintptr(data))
	callback := signalStyleContextChangedMap[index].callback
	callback()
}

// Adds a style class to @context, so posterior calls to
// gtk_style_context_get() or any of the gtk_render_*()
// functions will make use of this new class for styling.
//
// In the CSS file format, a #GtkEntry defining a “search”
// class, would be matched by:
//
// |[ <!-- language="CSS" -->
// entry.search { ... }
// ]|
//
// While any widget defining a “search” class would be
// matched by:
// |[ <!-- language="CSS" -->
// .search { ... }
// ]|
/*

C function

gtk_style_context_add_class
*/
func (recv *StyleContext) AddClass(className string) {
	c_class_name := C.CString(className)
	defer C.free(unsafe.Pointer(c_class_name))

	C.gtk_style_context_add_class((*C.GtkStyleContext)(recv.native), c_class_name)

	return
}

// Adds a style provider to @context, to be used in style construction.
// Note that a style provider added by this function only affects
// the style of the widget to which @context belongs. If you want
// to affect the style of all widgets, use
// gtk_style_context_add_provider_for_screen().
//
// Note: If both priorities are the same, a #GtkStyleProvider
// added through this function takes precedence over another added
// through gtk_style_context_add_provider_for_screen().
/*

C function

gtk_style_context_add_provider
*/
func (recv *StyleContext) AddProvider(provider *StyleProvider, priority uint32) {
	c_provider := (*C.GtkStyleProvider)(provider.ToC())

	c_priority := (C.guint)(priority)

	C.gtk_style_context_add_provider((*C.GtkStyleContext)(recv.native), c_provider, c_priority)

	return
}

// Adds a region to @context, so posterior calls to
// gtk_style_context_get() or any of the gtk_render_*()
// functions will make use of this new region for styling.
//
// In the CSS file format, a #GtkTreeView defining a “row”
// region, would be matched by:
//
// |[ <!-- language="CSS" -->
// treeview row { ... }
// ]|
//
// Pseudo-classes are used for matching @flags, so the two
// following rules:
// |[ <!-- language="CSS" -->
// treeview row:nth-child(even) { ... }
// treeview row:nth-child(odd) { ... }
// ]|
//
// would apply to even and odd rows, respectively.
//
// Region names must only contain lowercase letters
// and “-”, starting always with a lowercase letter.
/*

C function

gtk_style_context_add_region
*/
func (recv *StyleContext) AddRegion(regionName string, flags RegionFlags) {
	c_region_name := C.CString(regionName)
	defer C.free(unsafe.Pointer(c_region_name))

	c_flags := (C.GtkRegionFlags)(flags)

	C.gtk_style_context_add_region((*C.GtkStyleContext)(recv.native), c_region_name, c_flags)

	return
}

// Stops all running animations for @region_id and all animatable
// regions underneath.
//
// A %NULL @region_id will stop all ongoing animations in @context,
// when dealing with a #GtkStyleContext obtained through
// gtk_widget_get_style_context(), this is normally done for you
// in all circumstances you would expect all widget to be stopped,
// so this should be only used in complex widgets with different
// animatable regions.
/*

C function

gtk_style_context_cancel_animations
*/
func (recv *StyleContext) CancelAnimations(regionId uintptr) {
	c_region_id := (C.gpointer)(regionId)

	C.gtk_style_context_cancel_animations((*C.GtkStyleContext)(recv.native), c_region_id)

	return
}

// Unsupported : gtk_style_context_get : unsupported parameter ... : varargs

// Gets the background color for a given state.
//
// This function is far less useful than it seems, and it should not be used in
// newly written code. CSS has no concept of "background color", as a background
// can be an image, or a gradient, or any other pattern including solid colors.
//
// The only reason why you would call gtk_style_context_get_background_color() is
// to use the returned value to draw the background with it; the correct way to
// achieve this result is to use gtk_render_background() instead, along with CSS
// style classes to modify the color to be rendered.
/*

C function

gtk_style_context_get_background_color
*/
func (recv *StyleContext) GetBackgroundColor(state StateFlags) *gdk.RGBA {
	c_state := (C.GtkStateFlags)(state)

	var c_color C.GdkRGBA

	C.gtk_style_context_get_background_color((*C.GtkStyleContext)(recv.native), c_state, &c_color)

	color := gdk.RGBANewFromC(unsafe.Pointer(&c_color))

	return color
}

// Gets the border for a given state as a #GtkBorder.
//
// See gtk_style_context_get_property() and
// #GTK_STYLE_PROPERTY_BORDER_WIDTH for details.
/*

C function

gtk_style_context_get_border
*/
func (recv *StyleContext) GetBorder(state StateFlags) *Border {
	c_state := (C.GtkStateFlags)(state)

	var c_border C.GtkBorder

	C.gtk_style_context_get_border((*C.GtkStyleContext)(recv.native), c_state, &c_border)

	border := BorderNewFromC(unsafe.Pointer(&c_border))

	return border
}

// Gets the border color for a given state.
/*

C function

gtk_style_context_get_border_color
*/
func (recv *StyleContext) GetBorderColor(state StateFlags) *gdk.RGBA {
	c_state := (C.GtkStateFlags)(state)

	var c_color C.GdkRGBA

	C.gtk_style_context_get_border_color((*C.GtkStyleContext)(recv.native), c_state, &c_color)

	color := gdk.RGBANewFromC(unsafe.Pointer(&c_color))

	return color
}

// Gets the foreground color for a given state.
//
// See gtk_style_context_get_property() and
// #GTK_STYLE_PROPERTY_COLOR for details.
/*

C function

gtk_style_context_get_color
*/
func (recv *StyleContext) GetColor(state StateFlags) *gdk.RGBA {
	c_state := (C.GtkStateFlags)(state)

	var c_color C.GdkRGBA

	C.gtk_style_context_get_color((*C.GtkStyleContext)(recv.native), c_state, &c_color)

	color := gdk.RGBANewFromC(unsafe.Pointer(&c_color))

	return color
}

// Returns the widget direction used for rendering.
/*

C function

gtk_style_context_get_direction
*/
func (recv *StyleContext) GetDirection() TextDirection {
	retC := C.gtk_style_context_get_direction((*C.GtkStyleContext)(recv.native))
	retGo := (TextDirection)(retC)

	return retGo
}

// Returns the font description for a given state. The returned
// object is const and will remain valid until the
// #GtkStyleContext::changed signal happens.
/*

C function

gtk_style_context_get_font
*/
func (recv *StyleContext) GetFont(state StateFlags) *pango.FontDescription {
	c_state := (C.GtkStateFlags)(state)

	retC := C.gtk_style_context_get_font((*C.GtkStyleContext)(recv.native), c_state)
	retGo := pango.FontDescriptionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Returns the sides where rendered elements connect visually with others.
/*

C function

gtk_style_context_get_junction_sides
*/
func (recv *StyleContext) GetJunctionSides() JunctionSides {
	retC := C.gtk_style_context_get_junction_sides((*C.GtkStyleContext)(recv.native))
	retGo := (JunctionSides)(retC)

	return retGo
}

// Gets the margin for a given state as a #GtkBorder.
// See gtk_style_property_get() and #GTK_STYLE_PROPERTY_MARGIN
// for details.
/*

C function

gtk_style_context_get_margin
*/
func (recv *StyleContext) GetMargin(state StateFlags) *Border {
	c_state := (C.GtkStateFlags)(state)

	var c_margin C.GtkBorder

	C.gtk_style_context_get_margin((*C.GtkStyleContext)(recv.native), c_state, &c_margin)

	margin := BorderNewFromC(unsafe.Pointer(&c_margin))

	return margin
}

// Gets the padding for a given state as a #GtkBorder.
// See gtk_style_context_get() and #GTK_STYLE_PROPERTY_PADDING
// for details.
/*

C function

gtk_style_context_get_padding
*/
func (recv *StyleContext) GetPadding(state StateFlags) *Border {
	c_state := (C.GtkStateFlags)(state)

	var c_padding C.GtkBorder

	C.gtk_style_context_get_padding((*C.GtkStyleContext)(recv.native), c_state, &c_padding)

	padding := BorderNewFromC(unsafe.Pointer(&c_padding))

	return padding
}

// Returns the widget path used for style matching.
/*

C function

gtk_style_context_get_path
*/
func (recv *StyleContext) GetPath() *WidgetPath {
	retC := C.gtk_style_context_get_path((*C.GtkStyleContext)(recv.native))
	retGo := WidgetPathNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets a style property from @context for the given state.
//
// Note that not all CSS properties that are supported by GTK+ can be
// retrieved in this way, since they may not be representable as #GValue.
// GTK+ defines macros for a number of properties that can be used
// with this function.
//
// Note that passing a state other than the current state of @context
// is not recommended unless the style context has been saved with
// gtk_style_context_save().
//
// When @value is no longer needed, g_value_unset() must be called
// to free any allocated memory.
/*

C function

gtk_style_context_get_property
*/
func (recv *StyleContext) GetProperty(property string, state StateFlags) *gobject.Value {
	c_property := C.CString(property)
	defer C.free(unsafe.Pointer(c_property))

	c_state := (C.GtkStateFlags)(state)

	var c_value C.GValue

	C.gtk_style_context_get_property((*C.GtkStyleContext)(recv.native), c_property, c_state, &c_value)

	value := gobject.ValueNewFromC(unsafe.Pointer(&c_value))

	return value
}

// Returns the state used for style matching.
//
// This method should only be used to retrieve the #GtkStateFlags
// to pass to #GtkStyleContext methods, like gtk_style_context_get_padding().
// If you need to retrieve the current state of a #GtkWidget, use
// gtk_widget_get_state_flags().
/*

C function

gtk_style_context_get_state
*/
func (recv *StyleContext) GetState() StateFlags {
	retC := C.gtk_style_context_get_state((*C.GtkStyleContext)(recv.native))
	retGo := (StateFlags)(retC)

	return retGo
}

// Unsupported : gtk_style_context_get_style : unsupported parameter ... : varargs

// Unsupported : gtk_style_context_get_style_valist : unsupported parameter args : no type generator for va_list (va_list) for param args

// Unsupported : gtk_style_context_get_valist : unsupported parameter args : no type generator for va_list (va_list) for param args

// Returns %TRUE if @context currently has defined the
// given class name.
/*

C function

gtk_style_context_has_class
*/
func (recv *StyleContext) HasClass(className string) bool {
	c_class_name := C.CString(className)
	defer C.free(unsafe.Pointer(c_class_name))

	retC := C.gtk_style_context_has_class((*C.GtkStyleContext)(recv.native), c_class_name)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_style_context_has_region : unsupported parameter flags_return : GtkRegionFlags* with indirection level of 1

// Invalidates @context style information, so it will be reconstructed
// again. It is useful if you modify the @context and need the new
// information immediately.
/*

C function

gtk_style_context_invalidate
*/
func (recv *StyleContext) Invalidate() {
	C.gtk_style_context_invalidate((*C.GtkStyleContext)(recv.native))

	return
}

// Returns the list of classes currently defined in @context.
/*

C function

gtk_style_context_list_classes
*/
func (recv *StyleContext) ListClasses() *glib.List {
	retC := C.gtk_style_context_list_classes((*C.GtkStyleContext)(recv.native))
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Returns the list of regions currently defined in @context.
/*

C function

gtk_style_context_list_regions
*/
func (recv *StyleContext) ListRegions() *glib.List {
	retC := C.gtk_style_context_list_regions((*C.GtkStyleContext)(recv.native))
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Notifies a state change on @context, so if the current style makes use
// of transition animations, one will be started so all rendered elements
// under @region_id are animated for state @state being set to value
// @state_value.
//
// The @window parameter is used in order to invalidate the rendered area
// as the animation runs, so make sure it is the same window that is being
// rendered on by the gtk_render_*() functions.
//
// If @region_id is %NULL, all rendered elements using @context will be
// affected by this state transition.
//
// As a practical example, a #GtkButton notifying a state transition on
// the prelight state:
// |[ <!-- language="C" -->
// gtk_style_context_notify_state_change (context,
// gtk_widget_get_window (widget),
// NULL,
// GTK_STATE_PRELIGHT,
// button->in_button);
// ]|
//
// Can be handled in the CSS file like this:
// |[ <!-- language="CSS" -->
// button {
// background-color: #f00
// }
//
// button:hover {
// background-color: #fff;
// transition: 200ms linear
// }
// ]|
//
// This combination will animate the button background from red to white
// if a pointer enters the button, and back to red if the pointer leaves
// the button.
//
// Note that @state is used when finding the transition parameters, which
// is why the style places the transition under the :hover pseudo-class.
/*

C function

gtk_style_context_notify_state_change
*/
func (recv *StyleContext) NotifyStateChange(window *gdk.Window, regionId uintptr, state StateType, stateValue bool) {
	c_window := (*C.GdkWindow)(C.NULL)
	if window != nil {
		c_window = (*C.GdkWindow)(window.ToC())
	}

	c_region_id := (C.gpointer)(regionId)

	c_state := (C.GtkStateType)(state)

	c_state_value :=
		boolToGboolean(stateValue)

	C.gtk_style_context_notify_state_change((*C.GtkStyleContext)(recv.native), c_window, c_region_id, c_state, c_state_value)

	return
}

// Pops an animatable region from @context.
// See gtk_style_context_push_animatable_region().
/*

C function

gtk_style_context_pop_animatable_region
*/
func (recv *StyleContext) PopAnimatableRegion() {
	C.gtk_style_context_pop_animatable_region((*C.GtkStyleContext)(recv.native))

	return
}

// Pushes an animatable region, so all further gtk_render_*() calls between
// this call and the following gtk_style_context_pop_animatable_region()
// will potentially show transition animations for this region if
// gtk_style_context_notify_state_change() is called for a given state,
// and the current theme/style defines transition animations for state
// changes.
//
// The @region_id used must be unique in @context so the themes
// can uniquely identify rendered elements subject to a state transition.
/*

C function

gtk_style_context_push_animatable_region
*/
func (recv *StyleContext) PushAnimatableRegion(regionId uintptr) {
	c_region_id := (C.gpointer)(regionId)

	C.gtk_style_context_push_animatable_region((*C.GtkStyleContext)(recv.native), c_region_id)

	return
}

// Removes @class_name from @context.
/*

C function

gtk_style_context_remove_class
*/
func (recv *StyleContext) RemoveClass(className string) {
	c_class_name := C.CString(className)
	defer C.free(unsafe.Pointer(c_class_name))

	C.gtk_style_context_remove_class((*C.GtkStyleContext)(recv.native), c_class_name)

	return
}

// Removes @provider from the style providers list in @context.
/*

C function

gtk_style_context_remove_provider
*/
func (recv *StyleContext) RemoveProvider(provider *StyleProvider) {
	c_provider := (*C.GtkStyleProvider)(provider.ToC())

	C.gtk_style_context_remove_provider((*C.GtkStyleContext)(recv.native), c_provider)

	return
}

// Removes a region from @context.
/*

C function

gtk_style_context_remove_region
*/
func (recv *StyleContext) RemoveRegion(regionName string) {
	c_region_name := C.CString(regionName)
	defer C.free(unsafe.Pointer(c_region_name))

	C.gtk_style_context_remove_region((*C.GtkStyleContext)(recv.native), c_region_name)

	return
}

// Restores @context state to a previous stage.
// See gtk_style_context_save().
/*

C function

gtk_style_context_restore
*/
func (recv *StyleContext) Restore() {
	C.gtk_style_context_restore((*C.GtkStyleContext)(recv.native))

	return
}

// Saves the @context state, so temporary modifications done through
// gtk_style_context_add_class(), gtk_style_context_remove_class(),
// gtk_style_context_set_state(), etc. can quickly be reverted
// in one go through gtk_style_context_restore().
//
// The matching call to gtk_style_context_restore() must be done
// before GTK returns to the main loop.
/*

C function

gtk_style_context_save
*/
func (recv *StyleContext) Save() {
	C.gtk_style_context_save((*C.GtkStyleContext)(recv.native))

	return
}

// This function is analogous to gdk_window_scroll(), and
// should be called together with it so the invalidation
// areas for any ongoing animation are scrolled together
// with it.
/*

C function

gtk_style_context_scroll_animations
*/
func (recv *StyleContext) ScrollAnimations(window *gdk.Window, dx int32, dy int32) {
	c_window := (*C.GdkWindow)(C.NULL)
	if window != nil {
		c_window = (*C.GdkWindow)(window.ToC())
	}

	c_dx := (C.gint)(dx)

	c_dy := (C.gint)(dy)

	C.gtk_style_context_scroll_animations((*C.GtkStyleContext)(recv.native), c_window, c_dx, c_dy)

	return
}

// Sets the background of @window to the background pattern or
// color specified in @context for its current state.
/*

C function

gtk_style_context_set_background
*/
func (recv *StyleContext) SetBackground(window *gdk.Window) {
	c_window := (*C.GdkWindow)(C.NULL)
	if window != nil {
		c_window = (*C.GdkWindow)(window.ToC())
	}

	C.gtk_style_context_set_background((*C.GtkStyleContext)(recv.native), c_window)

	return
}

// Sets the reading direction for rendering purposes.
//
// If you are using a #GtkStyleContext returned from
// gtk_widget_get_style_context(), you do not need to
// call this yourself.
/*

C function

gtk_style_context_set_direction
*/
func (recv *StyleContext) SetDirection(direction TextDirection) {
	c_direction := (C.GtkTextDirection)(direction)

	C.gtk_style_context_set_direction((*C.GtkStyleContext)(recv.native), c_direction)

	return
}

// Sets the sides where rendered elements (mostly through
// gtk_render_frame()) will visually connect with other visual elements.
//
// This is merely a hint that may or may not be honored
// by themes.
//
// Container widgets are expected to set junction hints as appropriate
// for their children, so it should not normally be necessary to call
// this function manually.
/*

C function

gtk_style_context_set_junction_sides
*/
func (recv *StyleContext) SetJunctionSides(sides JunctionSides) {
	c_sides := (C.GtkJunctionSides)(sides)

	C.gtk_style_context_set_junction_sides((*C.GtkStyleContext)(recv.native), c_sides)

	return
}

// Sets the #GtkWidgetPath used for style matching. As a
// consequence, the style will be regenerated to match
// the new given path.
//
// If you are using a #GtkStyleContext returned from
// gtk_widget_get_style_context(), you do not need to call
// this yourself.
/*

C function

gtk_style_context_set_path
*/
func (recv *StyleContext) SetPath(path *WidgetPath) {
	c_path := (*C.GtkWidgetPath)(C.NULL)
	if path != nil {
		c_path = (*C.GtkWidgetPath)(path.ToC())
	}

	C.gtk_style_context_set_path((*C.GtkStyleContext)(recv.native), c_path)

	return
}

// Attaches @context to the given screen.
//
// The screen is used to add style information from “global” style
// providers, such as the screens #GtkSettings instance.
//
// If you are using a #GtkStyleContext returned from
// gtk_widget_get_style_context(), you do not need to
// call this yourself.
/*

C function

gtk_style_context_set_screen
*/
func (recv *StyleContext) SetScreen(screen *gdk.Screen) {
	c_screen := (*C.GdkScreen)(C.NULL)
	if screen != nil {
		c_screen = (*C.GdkScreen)(screen.ToC())
	}

	C.gtk_style_context_set_screen((*C.GtkStyleContext)(recv.native), c_screen)

	return
}

// Sets the state to be used for style matching.
/*

C function

gtk_style_context_set_state
*/
func (recv *StyleContext) SetState(flags StateFlags) {
	c_flags := (C.GtkStateFlags)(flags)

	C.gtk_style_context_set_state((*C.GtkStyleContext)(recv.native), c_flags)

	return
}

// Returns %TRUE if there is a transition animation running for the
// current region (see gtk_style_context_push_animatable_region()).
//
// If @progress is not %NULL, the animation progress will be returned
// there, 0.0 means the state is closest to being unset, while 1.0 means
// it’s closest to being set. This means transition animation will
// run from 0 to 1 when @state is being set and from 1 to 0 when
// it’s being unset.
/*

C function

gtk_style_context_state_is_running
*/
func (recv *StyleContext) StateIsRunning(state StateType) (bool, float64) {
	c_state := (C.GtkStateType)(state)

	var c_progress C.gdouble

	retC := C.gtk_style_context_state_is_running((*C.GtkStyleContext)(recv.native), c_state, &c_progress)
	retGo := retC == C.TRUE

	progress := (float64)(c_progress)

	return retGo, progress
}

// Unsupported : gtk_style_properties_get : unsupported parameter ... : varargs

// Gets a style property from @props for the given state. When done with @value,
// g_value_unset() needs to be called to free any allocated memory.
/*

C function

gtk_style_properties_get_property
*/
func (recv *StyleProperties) GetProperty(property string, state StateFlags) (bool, *gobject.Value) {
	c_property := C.CString(property)
	defer C.free(unsafe.Pointer(c_property))

	c_state := (C.GtkStateFlags)(state)

	var c_value C.GValue

	retC := C.gtk_style_properties_get_property((*C.GtkStyleProperties)(recv.native), c_property, c_state, &c_value)
	retGo := retC == C.TRUE

	value := gobject.ValueNewFromC(unsafe.Pointer(&c_value))

	return retGo, value
}

// Unsupported : gtk_style_properties_get_valist : unsupported parameter args : no type generator for va_list (va_list) for param args

// Returns the symbolic color that is mapped
// to @name.
/*

C function

gtk_style_properties_lookup_color
*/
func (recv *StyleProperties) LookupColor(name string) *SymbolicColor {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.gtk_style_properties_lookup_color((*C.GtkStyleProperties)(recv.native), c_name)
	retGo := SymbolicColorNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Maps @color so it can be referenced by @name. See
// gtk_style_properties_lookup_color()
/*

C function

gtk_style_properties_map_color
*/
func (recv *StyleProperties) MapColor(name string, color *SymbolicColor) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_color := (*C.GtkSymbolicColor)(C.NULL)
	if color != nil {
		c_color = (*C.GtkSymbolicColor)(color.ToC())
	}

	C.gtk_style_properties_map_color((*C.GtkStyleProperties)(recv.native), c_name, c_color)

	return
}

// Merges into @props all the style information contained
// in @props_to_merge. If @replace is %TRUE, the values
// will be overwritten, if it is %FALSE, the older values
// will prevail.
/*

C function

gtk_style_properties_merge
*/
func (recv *StyleProperties) Merge(propsToMerge *StyleProperties, replace bool) {
	c_props_to_merge := (*C.GtkStyleProperties)(C.NULL)
	if propsToMerge != nil {
		c_props_to_merge = (*C.GtkStyleProperties)(propsToMerge.ToC())
	}

	c_replace :=
		boolToGboolean(replace)

	C.gtk_style_properties_merge((*C.GtkStyleProperties)(recv.native), c_props_to_merge, c_replace)

	return
}

// Unsupported : gtk_style_properties_set : unsupported parameter ... : varargs

// Sets a styling property in @props.
/*

C function

gtk_style_properties_set_property
*/
func (recv *StyleProperties) SetProperty(property string, state StateFlags, value *gobject.Value) {
	c_property := C.CString(property)
	defer C.free(unsafe.Pointer(c_property))

	c_state := (C.GtkStateFlags)(state)

	c_value := (*C.GValue)(C.NULL)
	if value != nil {
		c_value = (*C.GValue)(value.ToC())
	}

	C.gtk_style_properties_set_property((*C.GtkStyleProperties)(recv.native), c_property, c_state, c_value)

	return
}

// Unsupported : gtk_style_properties_set_valist : unsupported parameter args : no type generator for va_list (va_list) for param args

// Unsets a style property in @props.
/*

C function

gtk_style_properties_unset_property
*/
func (recv *StyleProperties) UnsetProperty(property string, state StateFlags) {
	c_property := C.CString(property)
	defer C.free(unsafe.Pointer(c_property))

	c_state := (C.GtkStateFlags)(state)

	C.gtk_style_properties_unset_property((*C.GtkStyleProperties)(recv.native), c_property, c_state)

	return
}

// Creates a new #GtkSwitch widget.
/*

C function

gtk_switch_new
*/
func SwitchNew() *Switch {
	retC := C.gtk_switch_new()
	retGo := SwitchNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets whether the #GtkSwitch is in its “on” or “off” state.
/*

C function

gtk_switch_get_active
*/
func (recv *Switch) GetActive() bool {
	retC := C.gtk_switch_get_active((*C.GtkSwitch)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Changes the state of @sw to the desired one.
/*

C function

gtk_switch_set_active
*/
func (recv *Switch) SetActive(isActive bool) {
	c_is_active :=
		boolToGboolean(isActive)

	C.gtk_switch_set_active((*C.GtkSwitch)(recv.native), c_is_active)

	return
}

// Unsupported : gtk_text_view_get_cursor_locations : unsupported parameter strong : Blacklisted record : GdkRectangle

// Unsupported : gtk_theming_engine_get : unsupported parameter ... : varargs

// Gets the background color for a given state.
/*

C function

gtk_theming_engine_get_background_color
*/
func (recv *ThemingEngine) GetBackgroundColor(state StateFlags) *gdk.RGBA {
	c_state := (C.GtkStateFlags)(state)

	var c_color C.GdkRGBA

	C.gtk_theming_engine_get_background_color((*C.GtkThemingEngine)(recv.native), c_state, &c_color)

	color := gdk.RGBANewFromC(unsafe.Pointer(&c_color))

	return color
}

// Gets the border for a given state as a #GtkBorder.
/*

C function

gtk_theming_engine_get_border
*/
func (recv *ThemingEngine) GetBorder(state StateFlags) *Border {
	c_state := (C.GtkStateFlags)(state)

	var c_border C.GtkBorder

	C.gtk_theming_engine_get_border((*C.GtkThemingEngine)(recv.native), c_state, &c_border)

	border := BorderNewFromC(unsafe.Pointer(&c_border))

	return border
}

// Gets the border color for a given state.
/*

C function

gtk_theming_engine_get_border_color
*/
func (recv *ThemingEngine) GetBorderColor(state StateFlags) *gdk.RGBA {
	c_state := (C.GtkStateFlags)(state)

	var c_color C.GdkRGBA

	C.gtk_theming_engine_get_border_color((*C.GtkThemingEngine)(recv.native), c_state, &c_color)

	color := gdk.RGBANewFromC(unsafe.Pointer(&c_color))

	return color
}

// Gets the foreground color for a given state.
/*

C function

gtk_theming_engine_get_color
*/
func (recv *ThemingEngine) GetColor(state StateFlags) *gdk.RGBA {
	c_state := (C.GtkStateFlags)(state)

	var c_color C.GdkRGBA

	C.gtk_theming_engine_get_color((*C.GtkThemingEngine)(recv.native), c_state, &c_color)

	color := gdk.RGBANewFromC(unsafe.Pointer(&c_color))

	return color
}

// Returns the widget direction used for rendering.
/*

C function

gtk_theming_engine_get_direction
*/
func (recv *ThemingEngine) GetDirection() TextDirection {
	retC := C.gtk_theming_engine_get_direction((*C.GtkThemingEngine)(recv.native))
	retGo := (TextDirection)(retC)

	return retGo
}

// Returns the font description for a given state.
/*

C function

gtk_theming_engine_get_font
*/
func (recv *ThemingEngine) GetFont(state StateFlags) *pango.FontDescription {
	c_state := (C.GtkStateFlags)(state)

	retC := C.gtk_theming_engine_get_font((*C.GtkThemingEngine)(recv.native), c_state)
	retGo := pango.FontDescriptionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Returns the widget direction used for rendering.
/*

C function

gtk_theming_engine_get_junction_sides
*/
func (recv *ThemingEngine) GetJunctionSides() JunctionSides {
	retC := C.gtk_theming_engine_get_junction_sides((*C.GtkThemingEngine)(recv.native))
	retGo := (JunctionSides)(retC)

	return retGo
}

// Gets the margin for a given state as a #GtkBorder.
/*

C function

gtk_theming_engine_get_margin
*/
func (recv *ThemingEngine) GetMargin(state StateFlags) *Border {
	c_state := (C.GtkStateFlags)(state)

	var c_margin C.GtkBorder

	C.gtk_theming_engine_get_margin((*C.GtkThemingEngine)(recv.native), c_state, &c_margin)

	margin := BorderNewFromC(unsafe.Pointer(&c_margin))

	return margin
}

// Gets the padding for a given state as a #GtkBorder.
/*

C function

gtk_theming_engine_get_padding
*/
func (recv *ThemingEngine) GetPadding(state StateFlags) *Border {
	c_state := (C.GtkStateFlags)(state)

	var c_padding C.GtkBorder

	C.gtk_theming_engine_get_padding((*C.GtkThemingEngine)(recv.native), c_state, &c_padding)

	padding := BorderNewFromC(unsafe.Pointer(&c_padding))

	return padding
}

// Returns the widget path used for style matching.
/*

C function

gtk_theming_engine_get_path
*/
func (recv *ThemingEngine) GetPath() *WidgetPath {
	retC := C.gtk_theming_engine_get_path((*C.GtkThemingEngine)(recv.native))
	retGo := WidgetPathNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets a property value as retrieved from the style settings that apply
// to the currently rendered element.
/*

C function

gtk_theming_engine_get_property
*/
func (recv *ThemingEngine) GetProperty(property string, state StateFlags) *gobject.Value {
	c_property := C.CString(property)
	defer C.free(unsafe.Pointer(c_property))

	c_state := (C.GtkStateFlags)(state)

	var c_value C.GValue

	C.gtk_theming_engine_get_property((*C.GtkThemingEngine)(recv.native), c_property, c_state, &c_value)

	value := gobject.ValueNewFromC(unsafe.Pointer(&c_value))

	return value
}

// returns the state used when rendering.
/*

C function

gtk_theming_engine_get_state
*/
func (recv *ThemingEngine) GetState() StateFlags {
	retC := C.gtk_theming_engine_get_state((*C.GtkThemingEngine)(recv.native))
	retGo := (StateFlags)(retC)

	return retGo
}

// Unsupported : gtk_theming_engine_get_style : unsupported parameter ... : varargs

// Gets the value for a widget style property.
/*

C function

gtk_theming_engine_get_style_property
*/
func (recv *ThemingEngine) GetStyleProperty(propertyName string) *gobject.Value {
	c_property_name := C.CString(propertyName)
	defer C.free(unsafe.Pointer(c_property_name))

	var c_value C.GValue

	C.gtk_theming_engine_get_style_property((*C.GtkThemingEngine)(recv.native), c_property_name, &c_value)

	value := gobject.ValueNewFromC(unsafe.Pointer(&c_value))

	return value
}

// Unsupported : gtk_theming_engine_get_style_valist : unsupported parameter args : no type generator for va_list (va_list) for param args

// Unsupported : gtk_theming_engine_get_valist : unsupported parameter args : no type generator for va_list (va_list) for param args

// Returns %TRUE if the currently rendered contents have
// defined the given class name.
/*

C function

gtk_theming_engine_has_class
*/
func (recv *ThemingEngine) HasClass(styleClass string) bool {
	c_style_class := C.CString(styleClass)
	defer C.free(unsafe.Pointer(c_style_class))

	retC := C.gtk_theming_engine_has_class((*C.GtkThemingEngine)(recv.native), c_style_class)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_theming_engine_has_region : unsupported parameter flags : GtkRegionFlags* with indirection level of 1

// Looks up and resolves a color name in the current style’s color map.
/*

C function

gtk_theming_engine_lookup_color
*/
func (recv *ThemingEngine) LookupColor(colorName string) (bool, *gdk.RGBA) {
	c_color_name := C.CString(colorName)
	defer C.free(unsafe.Pointer(c_color_name))

	var c_color C.GdkRGBA

	retC := C.gtk_theming_engine_lookup_color((*C.GtkThemingEngine)(recv.native), c_color_name, &c_color)
	retGo := retC == C.TRUE

	color := gdk.RGBANewFromC(unsafe.Pointer(&c_color))

	return retGo, color
}

// Returns %TRUE if there is a transition animation running for the
// current region (see gtk_style_context_push_animatable_region()).
//
// If @progress is not %NULL, the animation progress will be returned
// there, 0.0 means the state is closest to being %FALSE, while 1.0 means
// it’s closest to being %TRUE. This means transition animations will
// run from 0 to 1 when @state is being set to %TRUE and from 1 to 0 when
// it’s being set to %FALSE.
/*

C function

gtk_theming_engine_state_is_running
*/
func (recv *ThemingEngine) StateIsRunning(state StateType) (bool, float64) {
	c_state := (C.GtkStateType)(state)

	var c_progress C.gdouble

	retC := C.gtk_theming_engine_state_is_running((*C.GtkThemingEngine)(recv.native), c_state, &c_progress)
	retGo := retC == C.TRUE

	progress := (float64)(c_progress)

	return retGo, progress
}

// Determine whether the point (@x, @y) in @tree_view is blank, that is no
// cell content nor an expander arrow is drawn at the location. If so, the
// location can be considered as the background. You might wish to take
// special action on clicks on the background, such as clearing a current
// selection, having a custom context menu or starting rubber banding.
//
// The @x and @y coordinate that are provided must be relative to bin_window
// coordinates.  That is, @x and @y must come from an event on @tree_view
// where `event->window == gtk_tree_view_get_bin_window ()`.
//
// For converting widget coordinates (eg. the ones you get from
// GtkWidget::query-tooltip), please see
// gtk_tree_view_convert_widget_to_bin_window_coords().
//
// The @path, @column, @cell_x and @cell_y arguments will be filled in
// likewise as for gtk_tree_view_get_path_at_pos().  Please see
// gtk_tree_view_get_path_at_pos() for more information.
/*

C function

gtk_tree_view_is_blank_at_pos
*/
func (recv *TreeView) IsBlankAtPos(x int32, y int32) (bool, *TreePath, *TreeViewColumn, int32, int32) {
	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	var c_path *C.GtkTreePath

	var c_column *C.GtkTreeViewColumn

	var c_cell_x C.gint

	var c_cell_y C.gint

	retC := C.gtk_tree_view_is_blank_at_pos((*C.GtkTreeView)(recv.native), c_x, c_y, &c_path, &c_column, &c_cell_x, &c_cell_y)
	retGo := retC == C.TRUE

	path := TreePathNewFromC(unsafe.Pointer(c_path))

	column := TreeViewColumnNewFromC(unsafe.Pointer(c_column))

	cellX := (int32)(c_cell_x)

	cellY := (int32)(c_cell_y)

	return retGo, path, column, cellX, cellY
}

// Creates a new #GtkTreeViewColumn using @area to render its cells.
/*

C function

gtk_tree_view_column_new_with_area
*/
func TreeViewColumnNewWithArea(area *CellArea) *TreeViewColumn {
	c_area := (*C.GtkCellArea)(C.NULL)
	if area != nil {
		c_area = (*C.GtkCellArea)(area.ToC())
	}

	retC := C.gtk_tree_view_column_new_with_area(c_area)
	retGo := TreeViewColumnNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Returns the button used in the treeview column header
/*

C function

gtk_tree_view_column_get_button
*/
func (recv *TreeViewColumn) GetButton() *Widget {
	retC := C.gtk_tree_view_column_get_button((*C.GtkTreeViewColumn)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

type signalWidgetDrawDetail struct {
	callback  WidgetSignalDrawCallback
	handlerID C.gulong
}

var signalWidgetDrawId int
var signalWidgetDrawMap = make(map[int]signalWidgetDrawDetail)
var signalWidgetDrawLock sync.Mutex

// WidgetSignalDrawCallback is a callback function for a 'draw' signal emitted from a Widget.
type WidgetSignalDrawCallback func(cr *cairo.Context) bool

/*
ConnectDraw connects the callback to the 'draw' signal for the Widget.

The returned value represents the connection, and may be passed to DisconnectDraw to remove it.
*/
func (recv *Widget) ConnectDraw(callback WidgetSignalDrawCallback) int {
	signalWidgetDrawLock.Lock()
	defer signalWidgetDrawLock.Unlock()

	signalWidgetDrawId++
	instance := C.gpointer(recv.native)
	handlerID := C.Widget_signal_connect_draw(instance, C.gpointer(uintptr(signalWidgetDrawId)))

	detail := signalWidgetDrawDetail{callback, handlerID}
	signalWidgetDrawMap[signalWidgetDrawId] = detail

	return signalWidgetDrawId
}

/*
DisconnectDraw disconnects a callback from the 'draw' signal for the Widget.

The connectionID should be a value returned from a call to ConnectDraw.
*/
func (recv *Widget) DisconnectDraw(connectionID int) {
	signalWidgetDrawLock.Lock()
	defer signalWidgetDrawLock.Unlock()

	detail, exists := signalWidgetDrawMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWidgetDrawMap, connectionID)
}

//export widget_drawHandler
func widget_drawHandler(_ *C.GObject, c_cr *C.cairo_t, data C.gpointer) C.gboolean {
	cr := cairo.ContextNewFromC(unsafe.Pointer(c_cr))

	index := int(uintptr(data))
	callback := signalWidgetDrawMap[index].callback
	retGo := callback(cr)
	retC :=
		boolToGboolean(retGo)
	return retC
}

// Unsupported signal 'state-flags-changed' for Widget : unsupported parameter flags : type StateFlags :

type signalWidgetStyleUpdatedDetail struct {
	callback  WidgetSignalStyleUpdatedCallback
	handlerID C.gulong
}

var signalWidgetStyleUpdatedId int
var signalWidgetStyleUpdatedMap = make(map[int]signalWidgetStyleUpdatedDetail)
var signalWidgetStyleUpdatedLock sync.Mutex

// WidgetSignalStyleUpdatedCallback is a callback function for a 'style-updated' signal emitted from a Widget.
type WidgetSignalStyleUpdatedCallback func()

/*
ConnectStyleUpdated connects the callback to the 'style-updated' signal for the Widget.

The returned value represents the connection, and may be passed to DisconnectStyleUpdated to remove it.
*/
func (recv *Widget) ConnectStyleUpdated(callback WidgetSignalStyleUpdatedCallback) int {
	signalWidgetStyleUpdatedLock.Lock()
	defer signalWidgetStyleUpdatedLock.Unlock()

	signalWidgetStyleUpdatedId++
	instance := C.gpointer(recv.native)
	handlerID := C.Widget_signal_connect_style_updated(instance, C.gpointer(uintptr(signalWidgetStyleUpdatedId)))

	detail := signalWidgetStyleUpdatedDetail{callback, handlerID}
	signalWidgetStyleUpdatedMap[signalWidgetStyleUpdatedId] = detail

	return signalWidgetStyleUpdatedId
}

/*
DisconnectStyleUpdated disconnects a callback from the 'style-updated' signal for the Widget.

The connectionID should be a value returned from a call to ConnectStyleUpdated.
*/
func (recv *Widget) DisconnectStyleUpdated(connectionID int) {
	signalWidgetStyleUpdatedLock.Lock()
	defer signalWidgetStyleUpdatedLock.Unlock()

	detail, exists := signalWidgetStyleUpdatedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWidgetStyleUpdatedMap, connectionID)
}

//export widget_styleUpdatedHandler
func widget_styleUpdatedHandler(_ *C.GObject, data C.gpointer) {
	index := int(uintptr(data))
	callback := signalWidgetStyleUpdatedMap[index].callback
	callback()
}

// Adds the device events in the bitfield @events to the event mask for
// @widget. See gtk_widget_set_device_events() for details.
/*

C function

gtk_widget_add_device_events
*/
func (recv *Widget) AddDeviceEvents(device *gdk.Device, events gdk.EventMask) {
	c_device := (*C.GdkDevice)(C.NULL)
	if device != nil {
		c_device = (*C.GdkDevice)(device.ToC())
	}

	c_events := (C.GdkEventMask)(events)

	C.gtk_widget_add_device_events((*C.GtkWidget)(recv.native), c_device, c_events)

	return
}

// Returns %TRUE if @device has been shadowed by a GTK+
// device grab on another widget, so it would stop sending
// events to @widget. This may be used in the
// #GtkWidget::grab-notify signal to check for specific
// devices. See gtk_device_grab_add().
/*

C function

gtk_widget_device_is_shadowed
*/
func (recv *Widget) DeviceIsShadowed(device *gdk.Device) bool {
	c_device := (*C.GdkDevice)(C.NULL)
	if device != nil {
		c_device = (*C.GdkDevice)(device.ToC())
	}

	retC := C.gtk_widget_device_is_shadowed((*C.GtkWidget)(recv.native), c_device)
	retGo := retC == C.TRUE

	return retGo
}

// Draws @widget to @cr. The top left corner of the widget will be
// drawn to the currently set origin point of @cr.
//
// You should pass a cairo context as @cr argument that is in an
// original state. Otherwise the resulting drawing is undefined. For
// example changing the operator using cairo_set_operator() or the
// line width using cairo_set_line_width() might have unwanted side
// effects.
// You may however change the context’s transform matrix - like with
// cairo_scale(), cairo_translate() or cairo_set_matrix() and clip
// region with cairo_clip() prior to calling this function. Also, it
// is fine to modify the context with cairo_save() and
// cairo_push_group() prior to calling this function.
//
// Note that special-purpose widgets may contain special code for
// rendering to the screen and might appear differently on screen
// and when rendered using gtk_widget_draw().
/*

C function

gtk_widget_draw
*/
func (recv *Widget) Draw(cr *cairo.Context) {
	c_cr := (*C.cairo_t)(C.NULL)
	if cr != nil {
		c_cr = (*C.cairo_t)(cr.ToC())
	}

	C.gtk_widget_draw((*C.GtkWidget)(recv.native), c_cr)

	return
}

// Returns whether @device can interact with @widget and its
// children. See gtk_widget_set_device_enabled().
/*

C function

gtk_widget_get_device_enabled
*/
func (recv *Widget) GetDeviceEnabled(device *gdk.Device) bool {
	c_device := (*C.GdkDevice)(C.NULL)
	if device != nil {
		c_device = (*C.GdkDevice)(device.ToC())
	}

	retC := C.gtk_widget_get_device_enabled((*C.GtkWidget)(recv.native), c_device)
	retGo := retC == C.TRUE

	return retGo
}

// Returns the events mask for the widget corresponding to an specific device. These
// are the events that the widget will receive when @device operates on it.
/*

C function

gtk_widget_get_device_events
*/
func (recv *Widget) GetDeviceEvents(device *gdk.Device) gdk.EventMask {
	c_device := (*C.GdkDevice)(C.NULL)
	if device != nil {
		c_device = (*C.GdkDevice)(device.ToC())
	}

	retC := C.gtk_widget_get_device_events((*C.GtkWidget)(recv.native), c_device)
	retGo := (gdk.EventMask)(retC)

	return retGo
}

// Gets the value of the #GtkWidget:margin-bottom property.
/*

C function

gtk_widget_get_margin_bottom
*/
func (recv *Widget) GetMarginBottom() int32 {
	retC := C.gtk_widget_get_margin_bottom((*C.GtkWidget)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Gets the value of the #GtkWidget:margin-left property.
/*

C function

gtk_widget_get_margin_left
*/
func (recv *Widget) GetMarginLeft() int32 {
	retC := C.gtk_widget_get_margin_left((*C.GtkWidget)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Gets the value of the #GtkWidget:margin-right property.
/*

C function

gtk_widget_get_margin_right
*/
func (recv *Widget) GetMarginRight() int32 {
	retC := C.gtk_widget_get_margin_right((*C.GtkWidget)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Gets the value of the #GtkWidget:margin-top property.
/*

C function

gtk_widget_get_margin_top
*/
func (recv *Widget) GetMarginTop() int32 {
	retC := C.gtk_widget_get_margin_top((*C.GtkWidget)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Retrieves a widget’s initial minimum and natural height.
//
// This call is specific to width-for-height requests.
//
// The returned request will be modified by the
// GtkWidgetClass::adjust_size_request virtual method and by any
// #GtkSizeGroups that have been applied. That is, the returned request
// is the one that should be used for layout, not necessarily the one
// returned by the widget itself.
/*

C function

gtk_widget_get_preferred_height
*/
func (recv *Widget) GetPreferredHeight() (int32, int32) {
	var c_minimum_height C.gint

	var c_natural_height C.gint

	C.gtk_widget_get_preferred_height((*C.GtkWidget)(recv.native), &c_minimum_height, &c_natural_height)

	minimumHeight := (int32)(c_minimum_height)

	naturalHeight := (int32)(c_natural_height)

	return minimumHeight, naturalHeight
}

// Retrieves a widget’s minimum and natural height if it would be given
// the specified @width.
//
// The returned request will be modified by the
// GtkWidgetClass::adjust_size_request virtual method and by any
// #GtkSizeGroups that have been applied. That is, the returned request
// is the one that should be used for layout, not necessarily the one
// returned by the widget itself.
/*

C function

gtk_widget_get_preferred_height_for_width
*/
func (recv *Widget) GetPreferredHeightForWidth(width int32) (int32, int32) {
	c_width := (C.gint)(width)

	var c_minimum_height C.gint

	var c_natural_height C.gint

	C.gtk_widget_get_preferred_height_for_width((*C.GtkWidget)(recv.native), c_width, &c_minimum_height, &c_natural_height)

	minimumHeight := (int32)(c_minimum_height)

	naturalHeight := (int32)(c_natural_height)

	return minimumHeight, naturalHeight
}

// Retrieves the minimum and natural size of a widget, taking
// into account the widget’s preference for height-for-width management.
//
// This is used to retrieve a suitable size by container widgets which do
// not impose any restrictions on the child placement. It can be used
// to deduce toplevel window and menu sizes as well as child widgets in
// free-form containers such as GtkLayout.
//
// Handle with care. Note that the natural height of a height-for-width
// widget will generally be a smaller size than the minimum height, since the required
// height for the natural width is generally smaller than the required height for
// the minimum width.
//
// Use gtk_widget_get_preferred_height_and_baseline_for_width() if you want to support
// baseline alignment.
/*

C function

gtk_widget_get_preferred_size
*/
func (recv *Widget) GetPreferredSize() (*Requisition, *Requisition) {
	var c_minimum_size C.GtkRequisition

	var c_natural_size C.GtkRequisition

	C.gtk_widget_get_preferred_size((*C.GtkWidget)(recv.native), &c_minimum_size, &c_natural_size)

	minimumSize := RequisitionNewFromC(unsafe.Pointer(&c_minimum_size))

	naturalSize := RequisitionNewFromC(unsafe.Pointer(&c_natural_size))

	return minimumSize, naturalSize
}

// Retrieves a widget’s initial minimum and natural width.
//
// This call is specific to height-for-width requests.
//
// The returned request will be modified by the
// GtkWidgetClass::adjust_size_request virtual method and by any
// #GtkSizeGroups that have been applied. That is, the returned request
// is the one that should be used for layout, not necessarily the one
// returned by the widget itself.
/*

C function

gtk_widget_get_preferred_width
*/
func (recv *Widget) GetPreferredWidth() (int32, int32) {
	var c_minimum_width C.gint

	var c_natural_width C.gint

	C.gtk_widget_get_preferred_width((*C.GtkWidget)(recv.native), &c_minimum_width, &c_natural_width)

	minimumWidth := (int32)(c_minimum_width)

	naturalWidth := (int32)(c_natural_width)

	return minimumWidth, naturalWidth
}

// Retrieves a widget’s minimum and natural width if it would be given
// the specified @height.
//
// The returned request will be modified by the
// GtkWidgetClass::adjust_size_request virtual method and by any
// #GtkSizeGroups that have been applied. That is, the returned request
// is the one that should be used for layout, not necessarily the one
// returned by the widget itself.
/*

C function

gtk_widget_get_preferred_width_for_height
*/
func (recv *Widget) GetPreferredWidthForHeight(height int32) (int32, int32) {
	c_height := (C.gint)(height)

	var c_minimum_width C.gint

	var c_natural_width C.gint

	C.gtk_widget_get_preferred_width_for_height((*C.GtkWidget)(recv.native), c_height, &c_minimum_width, &c_natural_width)

	minimumWidth := (int32)(c_minimum_width)

	naturalWidth := (int32)(c_natural_width)

	return minimumWidth, naturalWidth
}

// Gets whether the widget prefers a height-for-width layout
// or a width-for-height layout.
//
// #GtkBin widgets generally propagate the preference of
// their child, container widgets need to request something either in
// context of their children or in context of their allocation
// capabilities.
/*

C function

gtk_widget_get_request_mode
*/
func (recv *Widget) GetRequestMode() SizeRequestMode {
	retC := C.gtk_widget_get_request_mode((*C.GtkWidget)(recv.native))
	retGo := (SizeRequestMode)(retC)

	return retGo
}

// Returns the widget state as a flag set. It is worth mentioning
// that the effective %GTK_STATE_FLAG_INSENSITIVE state will be
// returned, that is, also based on parent insensitivity, even if
// @widget itself is sensitive.
//
// Also note that if you are looking for a way to obtain the
// #GtkStateFlags to pass to a #GtkStyleContext method, you
// should look at gtk_style_context_get_state().
/*

C function

gtk_widget_get_state_flags
*/
func (recv *Widget) GetStateFlags() StateFlags {
	retC := C.gtk_widget_get_state_flags((*C.GtkWidget)(recv.native))
	retGo := (StateFlags)(retC)

	return retGo
}

// Sets an input shape for this widget’s GDK window. This allows for
// windows which react to mouse click in a nonrectangular region, see
// gdk_window_input_shape_combine_region() for more information.
/*

C function

gtk_widget_input_shape_combine_region
*/
func (recv *Widget) InputShapeCombineRegion(region *cairo.Region) {
	c_region := (*C.cairo_region_t)(C.NULL)
	if region != nil {
		c_region = (*C.cairo_region_t)(region.ToC())
	}

	C.gtk_widget_input_shape_combine_region((*C.GtkWidget)(recv.native), c_region)

	return
}

// Sets the background color to use for a widget.
//
// All other style values are left untouched.
// See gtk_widget_override_color().
/*

C function

gtk_widget_override_background_color
*/
func (recv *Widget) OverrideBackgroundColor(state StateFlags, color *gdk.RGBA) {
	c_state := (C.GtkStateFlags)(state)

	c_color := (*C.GdkRGBA)(C.NULL)
	if color != nil {
		c_color = (*C.GdkRGBA)(color.ToC())
	}

	C.gtk_widget_override_background_color((*C.GtkWidget)(recv.native), c_state, c_color)

	return
}

// Sets the color to use for a widget.
//
// All other style values are left untouched.
//
// This function does not act recursively. Setting the color of a
// container does not affect its children. Note that some widgets that
// you may not think of as containers, for instance #GtkButtons,
// are actually containers.
//
// This API is mostly meant as a quick way for applications to
// change a widget appearance. If you are developing a widgets
// library and intend this change to be themeable, it is better
// done by setting meaningful CSS classes in your
// widget/container implementation through gtk_style_context_add_class().
//
// This way, your widget library can install a #GtkCssProvider
// with the %GTK_STYLE_PROVIDER_PRIORITY_FALLBACK priority in order
// to provide a default styling for those widgets that need so, and
// this theming may fully overridden by the user’s theme.
//
// Note that for complex widgets this may bring in undesired
// results (such as uniform background color everywhere), in
// these cases it is better to fully style such widgets through a
// #GtkCssProvider with the %GTK_STYLE_PROVIDER_PRIORITY_APPLICATION
// priority.
/*

C function

gtk_widget_override_color
*/
func (recv *Widget) OverrideColor(state StateFlags, color *gdk.RGBA) {
	c_state := (C.GtkStateFlags)(state)

	c_color := (*C.GdkRGBA)(C.NULL)
	if color != nil {
		c_color = (*C.GdkRGBA)(color.ToC())
	}

	C.gtk_widget_override_color((*C.GtkWidget)(recv.native), c_state, c_color)

	return
}

// Sets the cursor color to use in a widget, overriding the
// cursor-color and secondary-cursor-color
// style properties. All other style values are left untouched.
// See also gtk_widget_modify_style().
//
// Note that the underlying properties have the #GdkColor type,
// so the alpha value in @primary and @secondary will be ignored.
/*

C function

gtk_widget_override_cursor
*/
func (recv *Widget) OverrideCursor(cursor *gdk.RGBA, secondaryCursor *gdk.RGBA) {
	c_cursor := (*C.GdkRGBA)(C.NULL)
	if cursor != nil {
		c_cursor = (*C.GdkRGBA)(cursor.ToC())
	}

	c_secondary_cursor := (*C.GdkRGBA)(C.NULL)
	if secondaryCursor != nil {
		c_secondary_cursor = (*C.GdkRGBA)(secondaryCursor.ToC())
	}

	C.gtk_widget_override_cursor((*C.GtkWidget)(recv.native), c_cursor, c_secondary_cursor)

	return
}

// Sets the font to use for a widget. All other style values are
// left untouched. See gtk_widget_override_color().
/*

C function

gtk_widget_override_font
*/
func (recv *Widget) OverrideFont(fontDesc *pango.FontDescription) {
	c_font_desc := (*C.PangoFontDescription)(C.NULL)
	if fontDesc != nil {
		c_font_desc = (*C.PangoFontDescription)(fontDesc.ToC())
	}

	C.gtk_widget_override_font((*C.GtkWidget)(recv.native), c_font_desc)

	return
}

// Sets a symbolic color for a widget.
//
// All other style values are left untouched.
// See gtk_widget_override_color() for overriding the foreground
// or background color.
/*

C function

gtk_widget_override_symbolic_color
*/
func (recv *Widget) OverrideSymbolicColor(name string, color *gdk.RGBA) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_color := (*C.GdkRGBA)(C.NULL)
	if color != nil {
		c_color = (*C.GdkRGBA)(color.ToC())
	}

	C.gtk_widget_override_symbolic_color((*C.GtkWidget)(recv.native), c_name, c_color)

	return
}

// Invalidates the area of @widget defined by @region by calling
// gdk_window_invalidate_region() on the widget’s window and all its
// child windows. Once the main loop becomes idle (after the current
// batch of events has been processed, roughly), the window will
// receive expose events for the union of all regions that have been
// invalidated.
//
// Normally you would only use this function in widget
// implementations. You might also use it to schedule a redraw of a
// #GtkDrawingArea or some portion thereof.
/*

C function

gtk_widget_queue_draw_region
*/
func (recv *Widget) QueueDrawRegion(region *cairo.Region) {
	c_region := (*C.cairo_region_t)(C.NULL)
	if region != nil {
		c_region = (*C.cairo_region_t)(region.ToC())
	}

	C.gtk_widget_queue_draw_region((*C.GtkWidget)(recv.native), c_region)

	return
}

// A convenience function that uses the theme engine and style
// settings for @widget to look up @stock_id and render it to
// a pixbuf. @stock_id should be a stock icon ID such as
// #GTK_STOCK_OPEN or #GTK_STOCK_OK. @size should be a size
// such as #GTK_ICON_SIZE_MENU.
//
// The pixels in the returned #GdkPixbuf are shared with the rest of
// the application and should not be modified. The pixbuf should be freed
// after use with g_object_unref().
/*

C function

gtk_widget_render_icon_pixbuf
*/
func (recv *Widget) RenderIconPixbuf(stockId string, size IconSize) *gdkpixbuf.Pixbuf {
	c_stock_id := C.CString(stockId)
	defer C.free(unsafe.Pointer(c_stock_id))

	c_size := (C.GtkIconSize)(size)

	retC := C.gtk_widget_render_icon_pixbuf((*C.GtkWidget)(recv.native), c_stock_id, c_size)
	var retGo (*gdkpixbuf.Pixbuf)
	if retC == nil {
		retGo = nil
	} else {
		retGo = gdkpixbuf.PixbufNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Updates the style context of @widget and all descendants
// by updating its widget path. #GtkContainers may want
// to use this on a child when reordering it in a way that a different
// style might apply to it. See also gtk_container_get_path_for_child().
/*

C function

gtk_widget_reset_style
*/
func (recv *Widget) ResetStyle() {
	C.gtk_widget_reset_style((*C.GtkWidget)(recv.native))

	return
}

// Enables or disables a #GdkDevice to interact with @widget
// and all its children.
//
// It does so by descending through the #GdkWindow hierarchy
// and enabling the same mask that is has for core events
// (i.e. the one that gdk_window_get_events() returns).
/*

C function

gtk_widget_set_device_enabled
*/
func (recv *Widget) SetDeviceEnabled(device *gdk.Device, enabled bool) {
	c_device := (*C.GdkDevice)(C.NULL)
	if device != nil {
		c_device = (*C.GdkDevice)(device.ToC())
	}

	c_enabled :=
		boolToGboolean(enabled)

	C.gtk_widget_set_device_enabled((*C.GtkWidget)(recv.native), c_device, c_enabled)

	return
}

// Sets the device event mask (see #GdkEventMask) for a widget. The event
// mask determines which events a widget will receive from @device. Keep
// in mind that different widgets have different default event masks, and by
// changing the event mask you may disrupt a widget’s functionality,
// so be careful. This function must be called while a widget is
// unrealized. Consider gtk_widget_add_device_events() for widgets that are
// already realized, or if you want to preserve the existing event
// mask. This function can’t be used with windowless widgets (which return
// %FALSE from gtk_widget_get_has_window());
// to get events on those widgets, place them inside a #GtkEventBox
// and receive events on the event box.
/*

C function

gtk_widget_set_device_events
*/
func (recv *Widget) SetDeviceEvents(device *gdk.Device, events gdk.EventMask) {
	c_device := (*C.GdkDevice)(C.NULL)
	if device != nil {
		c_device = (*C.GdkDevice)(device.ToC())
	}

	c_events := (C.GdkEventMask)(events)

	C.gtk_widget_set_device_events((*C.GtkWidget)(recv.native), c_device, c_events)

	return
}

// Sets the bottom margin of @widget.
// See the #GtkWidget:margin-bottom property.
/*

C function

gtk_widget_set_margin_bottom
*/
func (recv *Widget) SetMarginBottom(margin int32) {
	c_margin := (C.gint)(margin)

	C.gtk_widget_set_margin_bottom((*C.GtkWidget)(recv.native), c_margin)

	return
}

// Sets the left margin of @widget.
// See the #GtkWidget:margin-left property.
/*

C function

gtk_widget_set_margin_left
*/
func (recv *Widget) SetMarginLeft(margin int32) {
	c_margin := (C.gint)(margin)

	C.gtk_widget_set_margin_left((*C.GtkWidget)(recv.native), c_margin)

	return
}

// Sets the right margin of @widget.
// See the #GtkWidget:margin-right property.
/*

C function

gtk_widget_set_margin_right
*/
func (recv *Widget) SetMarginRight(margin int32) {
	c_margin := (C.gint)(margin)

	C.gtk_widget_set_margin_right((*C.GtkWidget)(recv.native), c_margin)

	return
}

// Sets the top margin of @widget.
// See the #GtkWidget:margin-top property.
/*

C function

gtk_widget_set_margin_top
*/
func (recv *Widget) SetMarginTop(margin int32) {
	c_margin := (C.gint)(margin)

	C.gtk_widget_set_margin_top((*C.GtkWidget)(recv.native), c_margin)

	return
}

// This function is for use in widget implementations. Turns on flag
// values in the current widget state (insensitive, prelighted, etc.).
//
// This function accepts the values %GTK_STATE_FLAG_DIR_LTR and
// %GTK_STATE_FLAG_DIR_RTL but ignores them. If you want to set the widget's
// direction, use gtk_widget_set_direction().
//
// It is worth mentioning that any other state than %GTK_STATE_FLAG_INSENSITIVE,
// will be propagated down to all non-internal children if @widget is a
// #GtkContainer, while %GTK_STATE_FLAG_INSENSITIVE itself will be propagated
// down to all #GtkContainer children by different means than turning on the
// state flag down the hierarchy, both gtk_widget_get_state_flags() and
// gtk_widget_is_sensitive() will make use of these.
/*

C function

gtk_widget_set_state_flags
*/
func (recv *Widget) SetStateFlags(flags StateFlags, clear bool) {
	c_flags := (C.GtkStateFlags)(flags)

	c_clear :=
		boolToGboolean(clear)

	C.gtk_widget_set_state_flags((*C.GtkWidget)(recv.native), c_flags, c_clear)

	return
}

// Enables or disables multiple pointer awareness. If this setting is %TRUE,
// @widget will start receiving multiple, per device enter/leave events. Note
// that if custom #GdkWindows are created in #GtkWidget::realize,
// gdk_window_set_support_multidevice() will have to be called manually on them.
/*

C function

gtk_widget_set_support_multidevice
*/
func (recv *Widget) SetSupportMultidevice(supportMultidevice bool) {
	c_support_multidevice :=
		boolToGboolean(supportMultidevice)

	C.gtk_widget_set_support_multidevice((*C.GtkWidget)(recv.native), c_support_multidevice)

	return
}

// Sets a shape for this widget’s GDK window. This allows for
// transparent windows etc., see gdk_window_shape_combine_region()
// for more information.
/*

C function

gtk_widget_shape_combine_region
*/
func (recv *Widget) ShapeCombineRegion(region *cairo.Region) {
	c_region := (*C.cairo_region_t)(C.NULL)
	if region != nil {
		c_region = (*C.cairo_region_t)(region.ToC())
	}

	C.gtk_widget_shape_combine_region((*C.GtkWidget)(recv.native), c_region)

	return
}

// This function is for use in widget implementations. Turns off flag
// values for the current widget state (insensitive, prelighted, etc.).
// See gtk_widget_set_state_flags().
/*

C function

gtk_widget_unset_state_flags
*/
func (recv *Widget) UnsetStateFlags(flags StateFlags) {
	c_flags := (C.GtkStateFlags)(flags)

	C.gtk_widget_unset_state_flags((*C.GtkWidget)(recv.native), c_flags)

	return
}

// Gets the #GtkApplication associated with the window (if any).
/*

C function

gtk_window_get_application
*/
func (recv *Window) GetApplication() *Application {
	retC := C.gtk_window_get_application((*C.GtkWindow)(recv.native))
	var retGo (*Application)
	if retC == nil {
		retGo = nil
	} else {
		retGo = ApplicationNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Determines whether the window may have a resize grip.
/*

C function

gtk_window_get_has_resize_grip
*/
func (recv *Window) GetHasResizeGrip() bool {
	retC := C.gtk_window_get_has_resize_grip((*C.GtkWindow)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_window_get_resize_grip_area : unsupported parameter rect : Blacklisted record : GdkRectangle

// Determines whether a resize grip is visible for the specified window.
/*

C function

gtk_window_resize_grip_is_visible
*/
func (recv *Window) ResizeGripIsVisible() bool {
	retC := C.gtk_window_resize_grip_is_visible((*C.GtkWindow)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Like gtk_window_resize(), but @width and @height are interpreted
// in terms of the base size and increment set with
// gtk_window_set_geometry_hints.
/*

C function

gtk_window_resize_to_geometry
*/
func (recv *Window) ResizeToGeometry(width int32, height int32) {
	c_width := (C.gint)(width)

	c_height := (C.gint)(height)

	C.gtk_window_resize_to_geometry((*C.GtkWindow)(recv.native), c_width, c_height)

	return
}

// Sets or unsets the #GtkApplication associated with the window.
//
// The application will be kept alive for at least as long as it has any windows
// associated with it (see g_application_hold() for a way to keep it alive
// without windows).
//
// Normally, the connection between the application and the window will remain
// until the window is destroyed, but you can explicitly remove it by setting
// the @application to %NULL.
//
// This is equivalent to calling gtk_application_remove_window() and/or
// gtk_application_add_window() on the old/new applications as relevant.
/*

C function

gtk_window_set_application
*/
func (recv *Window) SetApplication(application *Application) {
	c_application := (*C.GtkApplication)(C.NULL)
	if application != nil {
		c_application = (*C.GtkApplication)(application.ToC())
	}

	C.gtk_window_set_application((*C.GtkWindow)(recv.native), c_application)

	return
}

// Like gtk_window_set_default_size(), but @width and @height are interpreted
// in terms of the base size and increment set with
// gtk_window_set_geometry_hints.
/*

C function

gtk_window_set_default_geometry
*/
func (recv *Window) SetDefaultGeometry(width int32, height int32) {
	c_width := (C.gint)(width)

	c_height := (C.gint)(height)

	C.gtk_window_set_default_geometry((*C.GtkWindow)(recv.native), c_width, c_height)

	return
}

// Sets whether @window has a corner resize grip.
//
// Note that the resize grip is only shown if the window
// is actually resizable and not maximized. Use
// gtk_window_resize_grip_is_visible() to find out if the
// resize grip is currently shown.
/*

C function

gtk_window_set_has_resize_grip
*/
func (recv *Window) SetHasResizeGrip(value bool) {
	c_value :=
		boolToGboolean(value)

	C.gtk_window_set_has_resize_grip((*C.GtkWindow)(recv.native), c_value)

	return
}

// Tells GTK+ whether to drop its extra reference to the window
// when gtk_widget_destroy() is called.
//
// This function is only exported for the benefit of language
// bindings which may need to keep the window alive until their
// wrapper object is garbage collected. There is no justification
// for ever calling this function in an application.
/*

C function

gtk_window_set_has_user_ref_count
*/
func (recv *Window) SetHasUserRefCount(setting bool) {
	c_setting :=
		boolToGboolean(setting)

	C.gtk_window_set_has_user_ref_count((*C.GtkWindow)(recv.native), c_setting)

	return
}

// Returns the current grab widget for @device, or %NULL if none.
/*

C function

gtk_window_group_get_current_device_grab
*/
func (recv *WindowGroup) GetCurrentDeviceGrab(device *gdk.Device) *Widget {
	c_device := (*C.GdkDevice)(C.NULL)
	if device != nil {
		c_device = (*C.GdkDevice)(device.ToC())
	}

	retC := C.gtk_window_group_get_current_device_grab((*C.GtkWindowGroup)(recv.native), c_device)
	var retGo (*Widget)
	if retC == nil {
		retGo = nil
	} else {
		retGo = WidgetNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}
