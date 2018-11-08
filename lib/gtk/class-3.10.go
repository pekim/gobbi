// This is a generated file - DO NOT EDIT
// +build gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	cairo "github.com/pekim/gobbi/lib/cairo"
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

	void listbox_rowActivatedHandler(GObject *, GtkListBoxRow *, gpointer);

	static gulong ListBox_signal_connect_row_activated(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "row-activated", G_CALLBACK(listbox_rowActivatedHandler), data);
	}

*/
/*

	void listbox_rowSelectedHandler(GObject *, GtkListBoxRow *, gpointer);

	static gulong ListBox_signal_connect_row_selected(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "row-selected", G_CALLBACK(listbox_rowSelectedHandler), data);
	}

*/
/*

	void listboxrow_activateHandler(GObject *, gpointer);

	static gulong ListBoxRow_signal_connect_activate(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "activate", G_CALLBACK(listboxrow_activateHandler), data);
	}

*/
/*

	void searchentry_searchChangedHandler(GObject *, gpointer);

	static gulong SearchEntry_signal_connect_search_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "search-changed", G_CALLBACK(searchentry_searchChangedHandler), data);
	}

*/
import "C"

// Unsupported : gtk_app_chooser_dialog_new : unsupported parameter file : no type generator for Gio.File (GFile*) for param file

// GetBaselinePosition is a wrapper around the C function gtk_box_get_baseline_position.
func (recv *Box) GetBaselinePosition() BaselinePosition {
	retC := C.gtk_box_get_baseline_position((*C.GtkBox)(recv.native))
	retGo := (BaselinePosition)(retC)

	return retGo
}

// SetBaselinePosition is a wrapper around the C function gtk_box_set_baseline_position.
func (recv *Box) SetBaselinePosition(position BaselinePosition) {
	c_position := (C.GtkBaselinePosition)(position)

	C.gtk_box_set_baseline_position((*C.GtkBox)(recv.native), c_position)

	return
}

// BuilderNewFromFile is a wrapper around the C function gtk_builder_new_from_file.
func BuilderNewFromFile(filename string) *Builder {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))

	retC := C.gtk_builder_new_from_file(c_filename)
	retGo := BuilderNewFromC(unsafe.Pointer(retC))

	return retGo
}

// BuilderNewFromResource is a wrapper around the C function gtk_builder_new_from_resource.
func BuilderNewFromResource(resourcePath string) *Builder {
	c_resource_path := C.CString(resourcePath)
	defer C.free(unsafe.Pointer(c_resource_path))

	retC := C.gtk_builder_new_from_resource(c_resource_path)
	retGo := BuilderNewFromC(unsafe.Pointer(retC))

	return retGo
}

// BuilderNewFromString is a wrapper around the C function gtk_builder_new_from_string.
func BuilderNewFromString(string string, length int64) *Builder {
	c_string := C.CString(string)
	defer C.free(unsafe.Pointer(c_string))

	c_length := (C.gssize)(length)

	retC := C.gtk_builder_new_from_string(c_string, c_length)
	retGo := BuilderNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_builder_add_callback_symbol : unsupported parameter callback_symbol : no type generator for GObject.Callback (GCallback) for param callback_symbol

// Unsupported : gtk_builder_add_callback_symbols : unsupported parameter first_callback_symbol : no type generator for GObject.Callback (GCallback) for param first_callback_symbol

// GetApplication is a wrapper around the C function gtk_builder_get_application.
func (recv *Builder) GetApplication() *Application {
	retC := C.gtk_builder_get_application((*C.GtkBuilder)(recv.native))
	var retGo (*Application)
	if retC == nil {
		retGo = nil
	} else {
		retGo = ApplicationNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Unsupported : gtk_builder_lookup_callback_symbol : no return generator

// SetApplication is a wrapper around the C function gtk_builder_set_application.
func (recv *Builder) SetApplication(application *Application) {
	c_application := (*C.GtkApplication)(application.ToC())

	C.gtk_builder_set_application((*C.GtkBuilder)(recv.native), c_application)

	return
}

// Unsupported : gtk_button_new_from_icon_name : unsupported parameter size : no type generator for gint (GtkIconSize) for param size

// Unsupported : gtk_combo_box_new_with_model : unsupported parameter model : no type generator for TreeModel (GtkTreeModel*) for param model

// Unsupported : gtk_combo_box_new_with_model_and_entry : unsupported parameter model : no type generator for TreeModel (GtkTreeModel*) for param model

// Unsupported : gtk_dialog_new_with_buttons : unsupported parameter ... : varargs

// GetTabs is a wrapper around the C function gtk_entry_get_tabs.
func (recv *Entry) GetTabs() *pango.TabArray {
	retC := C.gtk_entry_get_tabs((*C.GtkEntry)(recv.native))
	var retGo (*pango.TabArray)
	if retC == nil {
		retGo = nil
	} else {
		retGo = pango.TabArrayNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// SetTabs is a wrapper around the C function gtk_entry_set_tabs.
func (recv *Entry) SetTabs(tabs *pango.TabArray) {
	c_tabs := (*C.PangoTabArray)(tabs.ToC())

	C.gtk_entry_set_tabs((*C.GtkEntry)(recv.native), c_tabs)

	return
}

// Unsupported : EntryIconAccessible : no CType

// Unsupported : gtk_file_chooser_dialog_new : unsupported parameter ... : varargs

// Unsupported : gtk_file_filter_new_from_gvariant : unsupported parameter variant : Blacklisted record : GVariant

// GetBaselineRow is a wrapper around the C function gtk_grid_get_baseline_row.
func (recv *Grid) GetBaselineRow() int32 {
	retC := C.gtk_grid_get_baseline_row((*C.GtkGrid)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetRowBaselinePosition is a wrapper around the C function gtk_grid_get_row_baseline_position.
func (recv *Grid) GetRowBaselinePosition(row int32) BaselinePosition {
	c_row := (C.gint)(row)

	retC := C.gtk_grid_get_row_baseline_position((*C.GtkGrid)(recv.native), c_row)
	retGo := (BaselinePosition)(retC)

	return retGo
}

// RemoveColumn is a wrapper around the C function gtk_grid_remove_column.
func (recv *Grid) RemoveColumn(position int32) {
	c_position := (C.gint)(position)

	C.gtk_grid_remove_column((*C.GtkGrid)(recv.native), c_position)

	return
}

// RemoveRow is a wrapper around the C function gtk_grid_remove_row.
func (recv *Grid) RemoveRow(position int32) {
	c_position := (C.gint)(position)

	C.gtk_grid_remove_row((*C.GtkGrid)(recv.native), c_position)

	return
}

// SetBaselineRow is a wrapper around the C function gtk_grid_set_baseline_row.
func (recv *Grid) SetBaselineRow(row int32) {
	c_row := (C.gint)(row)

	C.gtk_grid_set_baseline_row((*C.GtkGrid)(recv.native), c_row)

	return
}

// SetRowBaselinePosition is a wrapper around the C function gtk_grid_set_row_baseline_position.
func (recv *Grid) SetRowBaselinePosition(row int32, pos BaselinePosition) {
	c_row := (C.gint)(row)

	c_pos := (C.GtkBaselinePosition)(pos)

	C.gtk_grid_set_row_baseline_position((*C.GtkGrid)(recv.native), c_row, c_pos)

	return
}

// HeaderBarNew is a wrapper around the C function gtk_header_bar_new.
func HeaderBarNew() *HeaderBar {
	retC := C.gtk_header_bar_new()
	retGo := HeaderBarNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetCustomTitle is a wrapper around the C function gtk_header_bar_get_custom_title.
func (recv *HeaderBar) GetCustomTitle() *Widget {
	retC := C.gtk_header_bar_get_custom_title((*C.GtkHeaderBar)(recv.native))
	var retGo (*Widget)
	if retC == nil {
		retGo = nil
	} else {
		retGo = WidgetNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetShowCloseButton is a wrapper around the C function gtk_header_bar_get_show_close_button.
func (recv *HeaderBar) GetShowCloseButton() bool {
	retC := C.gtk_header_bar_get_show_close_button((*C.GtkHeaderBar)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetSubtitle is a wrapper around the C function gtk_header_bar_get_subtitle.
func (recv *HeaderBar) GetSubtitle() string {
	retC := C.gtk_header_bar_get_subtitle((*C.GtkHeaderBar)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetTitle is a wrapper around the C function gtk_header_bar_get_title.
func (recv *HeaderBar) GetTitle() string {
	retC := C.gtk_header_bar_get_title((*C.GtkHeaderBar)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// PackEnd is a wrapper around the C function gtk_header_bar_pack_end.
func (recv *HeaderBar) PackEnd(child *Widget) {
	c_child := (*C.GtkWidget)(child.ToC())

	C.gtk_header_bar_pack_end((*C.GtkHeaderBar)(recv.native), c_child)

	return
}

// PackStart is a wrapper around the C function gtk_header_bar_pack_start.
func (recv *HeaderBar) PackStart(child *Widget) {
	c_child := (*C.GtkWidget)(child.ToC())

	C.gtk_header_bar_pack_start((*C.GtkHeaderBar)(recv.native), c_child)

	return
}

// SetCustomTitle is a wrapper around the C function gtk_header_bar_set_custom_title.
func (recv *HeaderBar) SetCustomTitle(titleWidget *Widget) {
	c_title_widget := (*C.GtkWidget)(titleWidget.ToC())

	C.gtk_header_bar_set_custom_title((*C.GtkHeaderBar)(recv.native), c_title_widget)

	return
}

// SetShowCloseButton is a wrapper around the C function gtk_header_bar_set_show_close_button.
func (recv *HeaderBar) SetShowCloseButton(setting bool) {
	c_setting :=
		boolToGboolean(setting)

	C.gtk_header_bar_set_show_close_button((*C.GtkHeaderBar)(recv.native), c_setting)

	return
}

// SetSubtitle is a wrapper around the C function gtk_header_bar_set_subtitle.
func (recv *HeaderBar) SetSubtitle(subtitle string) {
	c_subtitle := C.CString(subtitle)
	defer C.free(unsafe.Pointer(c_subtitle))

	C.gtk_header_bar_set_subtitle((*C.GtkHeaderBar)(recv.native), c_subtitle)

	return
}

// SetTitle is a wrapper around the C function gtk_header_bar_set_title.
func (recv *HeaderBar) SetTitle(title string) {
	c_title := C.CString(title)
	defer C.free(unsafe.Pointer(c_title))

	C.gtk_header_bar_set_title((*C.GtkHeaderBar)(recv.native), c_title)

	return
}

// GetBaseScale is a wrapper around the C function gtk_icon_info_get_base_scale.
func (recv *IconInfo) GetBaseScale() int32 {
	retC := C.gtk_icon_info_get_base_scale((*C.GtkIconInfo)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// LoadSurface is a wrapper around the C function gtk_icon_info_load_surface.
func (recv *IconInfo) LoadSurface(forWindow *gdk.Window) (*cairo.Surface, error) {
	c_for_window := (*C.GdkWindow)(forWindow.ToC())

	var cThrowableError *C.GError

	retC := C.gtk_icon_info_load_surface((*C.GtkIconInfo)(recv.native), c_for_window, &cThrowableError)
	retGo := cairo.SurfaceNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : gtk_icon_theme_choose_icon_for_scale : unsupported parameter icon_names : no type generator for utf8 (gchar) for array param icon_names

// LoadIconForScale is a wrapper around the C function gtk_icon_theme_load_icon_for_scale.
func (recv *IconTheme) LoadIconForScale(iconName string, size int32, scale int32, flags IconLookupFlags) (*gdkpixbuf.Pixbuf, error) {
	c_icon_name := C.CString(iconName)
	defer C.free(unsafe.Pointer(c_icon_name))

	c_size := (C.gint)(size)

	c_scale := (C.gint)(scale)

	c_flags := (C.GtkIconLookupFlags)(flags)

	var cThrowableError *C.GError

	retC := C.gtk_icon_theme_load_icon_for_scale((*C.GtkIconTheme)(recv.native), c_icon_name, c_size, c_scale, c_flags, &cThrowableError)
	var retGo (*gdkpixbuf.Pixbuf)
	if retC == nil {
		retGo = nil
	} else {
		retGo = gdkpixbuf.PixbufNewFromC(unsafe.Pointer(retC))
	}

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// LoadSurface is a wrapper around the C function gtk_icon_theme_load_surface.
func (recv *IconTheme) LoadSurface(iconName string, size int32, scale int32, forWindow *gdk.Window, flags IconLookupFlags) (*cairo.Surface, error) {
	c_icon_name := C.CString(iconName)
	defer C.free(unsafe.Pointer(c_icon_name))

	c_size := (C.gint)(size)

	c_scale := (C.gint)(scale)

	c_for_window := (*C.GdkWindow)(forWindow.ToC())

	c_flags := (C.GtkIconLookupFlags)(flags)

	var cThrowableError *C.GError

	retC := C.gtk_icon_theme_load_surface((*C.GtkIconTheme)(recv.native), c_icon_name, c_size, c_scale, c_for_window, c_flags, &cThrowableError)
	var retGo (*cairo.Surface)
	if retC == nil {
		retGo = nil
	} else {
		retGo = cairo.SurfaceNewFromC(unsafe.Pointer(retC))
	}

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : gtk_icon_theme_lookup_by_gicon_for_scale : unsupported parameter icon : no type generator for Gio.Icon (GIcon*) for param icon

// LookupIconForScale is a wrapper around the C function gtk_icon_theme_lookup_icon_for_scale.
func (recv *IconTheme) LookupIconForScale(iconName string, size int32, scale int32, flags IconLookupFlags) *IconInfo {
	c_icon_name := C.CString(iconName)
	defer C.free(unsafe.Pointer(c_icon_name))

	c_size := (C.gint)(size)

	c_scale := (C.gint)(scale)

	c_flags := (C.GtkIconLookupFlags)(flags)

	retC := C.gtk_icon_theme_lookup_icon_for_scale((*C.GtkIconTheme)(recv.native), c_icon_name, c_size, c_scale, c_flags)
	var retGo (*IconInfo)
	if retC == nil {
		retGo = nil
	} else {
		retGo = IconInfoNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Unsupported : gtk_icon_view_new_with_model : unsupported parameter model : no type generator for TreeModel (GtkTreeModel*) for param model

// Unsupported : gtk_image_new_from_gicon : unsupported parameter icon : no type generator for Gio.Icon (GIcon*) for param icon

// Unsupported : gtk_image_new_from_icon_name : unsupported parameter size : no type generator for gint (GtkIconSize) for param size

// Unsupported : gtk_image_new_from_icon_set : unsupported parameter size : no type generator for gint (GtkIconSize) for param size

// Unsupported : gtk_image_new_from_stock : unsupported parameter size : no type generator for gint (GtkIconSize) for param size

// ImageNewFromSurface is a wrapper around the C function gtk_image_new_from_surface.
func ImageNewFromSurface(surface *cairo.Surface) *Image {
	c_surface := (*C.cairo_surface_t)(surface.ToC())

	retC := C.gtk_image_new_from_surface(c_surface)
	retGo := ImageNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SetFromSurface is a wrapper around the C function gtk_image_set_from_surface.
func (recv *Image) SetFromSurface(surface *cairo.Surface) {
	c_surface := (*C.cairo_surface_t)(surface.ToC())

	C.gtk_image_set_from_surface((*C.GtkImage)(recv.native), c_surface)

	return
}

// Unsupported : gtk_info_bar_new_with_buttons : unsupported parameter ... : varargs

// GetShowCloseButton is a wrapper around the C function gtk_info_bar_get_show_close_button.
func (recv *InfoBar) GetShowCloseButton() bool {
	retC := C.gtk_info_bar_get_show_close_button((*C.GtkInfoBar)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetShowCloseButton is a wrapper around the C function gtk_info_bar_set_show_close_button.
func (recv *InfoBar) SetShowCloseButton(setting bool) {
	c_setting :=
		boolToGboolean(setting)

	C.gtk_info_bar_set_show_close_button((*C.GtkInfoBar)(recv.native), c_setting)

	return
}

// GetLines is a wrapper around the C function gtk_label_get_lines.
func (recv *Label) GetLines() int32 {
	retC := C.gtk_label_get_lines((*C.GtkLabel)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// SetLines is a wrapper around the C function gtk_label_set_lines.
func (recv *Label) SetLines(lines int32) {
	c_lines := (C.gint)(lines)

	C.gtk_label_set_lines((*C.GtkLabel)(recv.native), c_lines)

	return
}

type signalListBoxRowActivatedDetail struct {
	callback  ListBoxSignalRowActivatedCallback
	handlerID C.gulong
}

var signalListBoxRowActivatedId int
var signalListBoxRowActivatedMap = make(map[int]signalListBoxRowActivatedDetail)
var signalListBoxRowActivatedLock sync.Mutex

// ListBoxSignalRowActivatedCallback is a callback function for a 'row-activated' signal emitted from a ListBox.
type ListBoxSignalRowActivatedCallback func(row *ListBoxRow)

/*
ConnectRowActivated connects the callback to the 'row-activated' signal for the ListBox.

The returned value represents the connection, and may be passed to DisconnectRowActivated to remove it.
*/
func (recv *ListBox) ConnectRowActivated(callback ListBoxSignalRowActivatedCallback) int {
	signalListBoxRowActivatedLock.Lock()
	defer signalListBoxRowActivatedLock.Unlock()

	signalListBoxRowActivatedId++
	instance := C.gpointer(recv.native)
	handlerID := C.ListBox_signal_connect_row_activated(instance, C.gpointer(uintptr(signalListBoxRowActivatedId)))

	detail := signalListBoxRowActivatedDetail{callback, handlerID}
	signalListBoxRowActivatedMap[signalListBoxRowActivatedId] = detail

	return signalListBoxRowActivatedId
}

/*
DisconnectRowActivated disconnects a callback from the 'row-activated' signal for the ListBox.

The connectionID should be a value returned from a call to ConnectRowActivated.
*/
func (recv *ListBox) DisconnectRowActivated(connectionID int) {
	signalListBoxRowActivatedLock.Lock()
	defer signalListBoxRowActivatedLock.Unlock()

	detail, exists := signalListBoxRowActivatedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalListBoxRowActivatedMap, connectionID)
}

//export listbox_rowActivatedHandler
func listbox_rowActivatedHandler(_ *C.GObject, c_row *C.GtkListBoxRow, data C.gpointer) {
	row := ListBoxRowNewFromC(unsafe.Pointer(c_row))

	index := int(uintptr(data))
	callback := signalListBoxRowActivatedMap[index].callback
	callback(row)
}

type signalListBoxRowSelectedDetail struct {
	callback  ListBoxSignalRowSelectedCallback
	handlerID C.gulong
}

var signalListBoxRowSelectedId int
var signalListBoxRowSelectedMap = make(map[int]signalListBoxRowSelectedDetail)
var signalListBoxRowSelectedLock sync.Mutex

// ListBoxSignalRowSelectedCallback is a callback function for a 'row-selected' signal emitted from a ListBox.
type ListBoxSignalRowSelectedCallback func(row *ListBoxRow)

/*
ConnectRowSelected connects the callback to the 'row-selected' signal for the ListBox.

The returned value represents the connection, and may be passed to DisconnectRowSelected to remove it.
*/
func (recv *ListBox) ConnectRowSelected(callback ListBoxSignalRowSelectedCallback) int {
	signalListBoxRowSelectedLock.Lock()
	defer signalListBoxRowSelectedLock.Unlock()

	signalListBoxRowSelectedId++
	instance := C.gpointer(recv.native)
	handlerID := C.ListBox_signal_connect_row_selected(instance, C.gpointer(uintptr(signalListBoxRowSelectedId)))

	detail := signalListBoxRowSelectedDetail{callback, handlerID}
	signalListBoxRowSelectedMap[signalListBoxRowSelectedId] = detail

	return signalListBoxRowSelectedId
}

/*
DisconnectRowSelected disconnects a callback from the 'row-selected' signal for the ListBox.

The connectionID should be a value returned from a call to ConnectRowSelected.
*/
func (recv *ListBox) DisconnectRowSelected(connectionID int) {
	signalListBoxRowSelectedLock.Lock()
	defer signalListBoxRowSelectedLock.Unlock()

	detail, exists := signalListBoxRowSelectedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalListBoxRowSelectedMap, connectionID)
}

//export listbox_rowSelectedHandler
func listbox_rowSelectedHandler(_ *C.GObject, c_row *C.GtkListBoxRow, data C.gpointer) {
	row := ListBoxRowNewFromC(unsafe.Pointer(c_row))

	index := int(uintptr(data))
	callback := signalListBoxRowSelectedMap[index].callback
	callback(row)
}

// ListBoxNew is a wrapper around the C function gtk_list_box_new.
func ListBoxNew() *ListBox {
	retC := C.gtk_list_box_new()
	retGo := ListBoxNewFromC(unsafe.Pointer(retC))

	return retGo
}

// DragHighlightRow is a wrapper around the C function gtk_list_box_drag_highlight_row.
func (recv *ListBox) DragHighlightRow(row *ListBoxRow) {
	c_row := (*C.GtkListBoxRow)(row.ToC())

	C.gtk_list_box_drag_highlight_row((*C.GtkListBox)(recv.native), c_row)

	return
}

// DragUnhighlightRow is a wrapper around the C function gtk_list_box_drag_unhighlight_row.
func (recv *ListBox) DragUnhighlightRow() {
	C.gtk_list_box_drag_unhighlight_row((*C.GtkListBox)(recv.native))

	return
}

// GetActivateOnSingleClick is a wrapper around the C function gtk_list_box_get_activate_on_single_click.
func (recv *ListBox) GetActivateOnSingleClick() bool {
	retC := C.gtk_list_box_get_activate_on_single_click((*C.GtkListBox)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetAdjustment is a wrapper around the C function gtk_list_box_get_adjustment.
func (recv *ListBox) GetAdjustment() *Adjustment {
	retC := C.gtk_list_box_get_adjustment((*C.GtkListBox)(recv.native))
	retGo := AdjustmentNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetRowAtIndex is a wrapper around the C function gtk_list_box_get_row_at_index.
func (recv *ListBox) GetRowAtIndex(index int32) *ListBoxRow {
	c_index_ := (C.gint)(index)

	retC := C.gtk_list_box_get_row_at_index((*C.GtkListBox)(recv.native), c_index_)
	var retGo (*ListBoxRow)
	if retC == nil {
		retGo = nil
	} else {
		retGo = ListBoxRowNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetRowAtY is a wrapper around the C function gtk_list_box_get_row_at_y.
func (recv *ListBox) GetRowAtY(y int32) *ListBoxRow {
	c_y := (C.gint)(y)

	retC := C.gtk_list_box_get_row_at_y((*C.GtkListBox)(recv.native), c_y)
	var retGo (*ListBoxRow)
	if retC == nil {
		retGo = nil
	} else {
		retGo = ListBoxRowNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetSelectedRow is a wrapper around the C function gtk_list_box_get_selected_row.
func (recv *ListBox) GetSelectedRow() *ListBoxRow {
	retC := C.gtk_list_box_get_selected_row((*C.GtkListBox)(recv.native))
	retGo := ListBoxRowNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetSelectionMode is a wrapper around the C function gtk_list_box_get_selection_mode.
func (recv *ListBox) GetSelectionMode() SelectionMode {
	retC := C.gtk_list_box_get_selection_mode((*C.GtkListBox)(recv.native))
	retGo := (SelectionMode)(retC)

	return retGo
}

// Insert is a wrapper around the C function gtk_list_box_insert.
func (recv *ListBox) Insert(child *Widget, position int32) {
	c_child := (*C.GtkWidget)(child.ToC())

	c_position := (C.gint)(position)

	C.gtk_list_box_insert((*C.GtkListBox)(recv.native), c_child, c_position)

	return
}

// InvalidateFilter is a wrapper around the C function gtk_list_box_invalidate_filter.
func (recv *ListBox) InvalidateFilter() {
	C.gtk_list_box_invalidate_filter((*C.GtkListBox)(recv.native))

	return
}

// InvalidateHeaders is a wrapper around the C function gtk_list_box_invalidate_headers.
func (recv *ListBox) InvalidateHeaders() {
	C.gtk_list_box_invalidate_headers((*C.GtkListBox)(recv.native))

	return
}

// InvalidateSort is a wrapper around the C function gtk_list_box_invalidate_sort.
func (recv *ListBox) InvalidateSort() {
	C.gtk_list_box_invalidate_sort((*C.GtkListBox)(recv.native))

	return
}

// Prepend is a wrapper around the C function gtk_list_box_prepend.
func (recv *ListBox) Prepend(child *Widget) {
	c_child := (*C.GtkWidget)(child.ToC())

	C.gtk_list_box_prepend((*C.GtkListBox)(recv.native), c_child)

	return
}

// SelectRow is a wrapper around the C function gtk_list_box_select_row.
func (recv *ListBox) SelectRow(row *ListBoxRow) {
	c_row := (*C.GtkListBoxRow)(row.ToC())

	C.gtk_list_box_select_row((*C.GtkListBox)(recv.native), c_row)

	return
}

// SetActivateOnSingleClick is a wrapper around the C function gtk_list_box_set_activate_on_single_click.
func (recv *ListBox) SetActivateOnSingleClick(single bool) {
	c_single :=
		boolToGboolean(single)

	C.gtk_list_box_set_activate_on_single_click((*C.GtkListBox)(recv.native), c_single)

	return
}

// SetAdjustment is a wrapper around the C function gtk_list_box_set_adjustment.
func (recv *ListBox) SetAdjustment(adjustment *Adjustment) {
	c_adjustment := (*C.GtkAdjustment)(adjustment.ToC())

	C.gtk_list_box_set_adjustment((*C.GtkListBox)(recv.native), c_adjustment)

	return
}

// Unsupported : gtk_list_box_set_filter_func : unsupported parameter filter_func : no type generator for ListBoxFilterFunc (GtkListBoxFilterFunc) for param filter_func

// Unsupported : gtk_list_box_set_header_func : unsupported parameter update_header : no type generator for ListBoxUpdateHeaderFunc (GtkListBoxUpdateHeaderFunc) for param update_header

// SetPlaceholder is a wrapper around the C function gtk_list_box_set_placeholder.
func (recv *ListBox) SetPlaceholder(placeholder *Widget) {
	c_placeholder := (*C.GtkWidget)(placeholder.ToC())

	C.gtk_list_box_set_placeholder((*C.GtkListBox)(recv.native), c_placeholder)

	return
}

// SetSelectionMode is a wrapper around the C function gtk_list_box_set_selection_mode.
func (recv *ListBox) SetSelectionMode(mode SelectionMode) {
	c_mode := (C.GtkSelectionMode)(mode)

	C.gtk_list_box_set_selection_mode((*C.GtkListBox)(recv.native), c_mode)

	return
}

// Unsupported : gtk_list_box_set_sort_func : unsupported parameter sort_func : no type generator for ListBoxSortFunc (GtkListBoxSortFunc) for param sort_func

type signalListBoxRowActivateDetail struct {
	callback  ListBoxRowSignalActivateCallback
	handlerID C.gulong
}

var signalListBoxRowActivateId int
var signalListBoxRowActivateMap = make(map[int]signalListBoxRowActivateDetail)
var signalListBoxRowActivateLock sync.Mutex

// ListBoxRowSignalActivateCallback is a callback function for a 'activate' signal emitted from a ListBoxRow.
type ListBoxRowSignalActivateCallback func()

/*
ConnectActivate connects the callback to the 'activate' signal for the ListBoxRow.

The returned value represents the connection, and may be passed to DisconnectActivate to remove it.
*/
func (recv *ListBoxRow) ConnectActivate(callback ListBoxRowSignalActivateCallback) int {
	signalListBoxRowActivateLock.Lock()
	defer signalListBoxRowActivateLock.Unlock()

	signalListBoxRowActivateId++
	instance := C.gpointer(recv.native)
	handlerID := C.ListBoxRow_signal_connect_activate(instance, C.gpointer(uintptr(signalListBoxRowActivateId)))

	detail := signalListBoxRowActivateDetail{callback, handlerID}
	signalListBoxRowActivateMap[signalListBoxRowActivateId] = detail

	return signalListBoxRowActivateId
}

/*
DisconnectActivate disconnects a callback from the 'activate' signal for the ListBoxRow.

The connectionID should be a value returned from a call to ConnectActivate.
*/
func (recv *ListBoxRow) DisconnectActivate(connectionID int) {
	signalListBoxRowActivateLock.Lock()
	defer signalListBoxRowActivateLock.Unlock()

	detail, exists := signalListBoxRowActivateMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalListBoxRowActivateMap, connectionID)
}

//export listboxrow_activateHandler
func listboxrow_activateHandler(_ *C.GObject, data C.gpointer) {
	index := int(uintptr(data))
	callback := signalListBoxRowActivateMap[index].callback
	callback()
}

// ListBoxRowNew is a wrapper around the C function gtk_list_box_row_new.
func ListBoxRowNew() *ListBoxRow {
	retC := C.gtk_list_box_row_new()
	retGo := ListBoxRowNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Changed is a wrapper around the C function gtk_list_box_row_changed.
func (recv *ListBoxRow) Changed() {
	C.gtk_list_box_row_changed((*C.GtkListBoxRow)(recv.native))

	return
}

// GetHeader is a wrapper around the C function gtk_list_box_row_get_header.
func (recv *ListBoxRow) GetHeader() *Widget {
	retC := C.gtk_list_box_row_get_header((*C.GtkListBoxRow)(recv.native))
	var retGo (*Widget)
	if retC == nil {
		retGo = nil
	} else {
		retGo = WidgetNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetIndex is a wrapper around the C function gtk_list_box_row_get_index.
func (recv *ListBoxRow) GetIndex() int32 {
	retC := C.gtk_list_box_row_get_index((*C.GtkListBoxRow)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// SetHeader is a wrapper around the C function gtk_list_box_row_set_header.
func (recv *ListBoxRow) SetHeader(header *Widget) {
	c_header := (*C.GtkWidget)(header.ToC())

	C.gtk_list_box_row_set_header((*C.GtkListBoxRow)(recv.native), c_header)

	return
}

// Unsupported : gtk_list_store_new : unsupported parameter ... : varargs

// Unsupported : gtk_list_store_newv : unsupported parameter types : no type generator for GType (GType) for array param types

// Unsupported : gtk_message_dialog_new : unsupported parameter ... : varargs

// Unsupported : gtk_message_dialog_new_with_markup : unsupported parameter ... : varargs

// Unsupported : gtk_page_setup_new_from_gvariant : unsupported parameter variant : Blacklisted record : GVariant

// Unsupported signal 'drag-action-ask' for PlacesSidebar : unsupported parameter actions : type gint :

// Unsupported signal 'drag-action-requested' for PlacesSidebar : unsupported parameter dest_file : no type generator for Gio.File,

// Unsupported signal 'drag-perform-drop' for PlacesSidebar : unsupported parameter dest_file : no type generator for Gio.File,

// Unsupported signal 'open-location' for PlacesSidebar : unsupported parameter location : no type generator for Gio.File,

// Unsupported signal 'populate-popup' for PlacesSidebar : unsupported parameter selected_item : no type generator for Gio.File,

// Unsupported signal 'show-error-message' for PlacesSidebar : unsupported parameter primary : type utf8 :

// PlacesSidebarNew is a wrapper around the C function gtk_places_sidebar_new.
func PlacesSidebarNew() *PlacesSidebar {
	retC := C.gtk_places_sidebar_new()
	retGo := PlacesSidebarNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_places_sidebar_add_shortcut : unsupported parameter location : no type generator for Gio.File (GFile*) for param location

// Unsupported : gtk_places_sidebar_get_location : no return generator

// Unsupported : gtk_places_sidebar_get_nth_bookmark : no return generator

// GetOpenFlags is a wrapper around the C function gtk_places_sidebar_get_open_flags.
func (recv *PlacesSidebar) GetOpenFlags() PlacesOpenFlags {
	retC := C.gtk_places_sidebar_get_open_flags((*C.GtkPlacesSidebar)(recv.native))
	retGo := (PlacesOpenFlags)(retC)

	return retGo
}

// GetShowDesktop is a wrapper around the C function gtk_places_sidebar_get_show_desktop.
func (recv *PlacesSidebar) GetShowDesktop() bool {
	retC := C.gtk_places_sidebar_get_show_desktop((*C.GtkPlacesSidebar)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// ListShortcuts is a wrapper around the C function gtk_places_sidebar_list_shortcuts.
func (recv *PlacesSidebar) ListShortcuts() *glib.SList {
	retC := C.gtk_places_sidebar_list_shortcuts((*C.GtkPlacesSidebar)(recv.native))
	retGo := glib.SListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_places_sidebar_remove_shortcut : unsupported parameter location : no type generator for Gio.File (GFile*) for param location

// Unsupported : gtk_places_sidebar_set_location : unsupported parameter location : no type generator for Gio.File (GFile*) for param location

// SetOpenFlags is a wrapper around the C function gtk_places_sidebar_set_open_flags.
func (recv *PlacesSidebar) SetOpenFlags(flags PlacesOpenFlags) {
	c_flags := (C.GtkPlacesOpenFlags)(flags)

	C.gtk_places_sidebar_set_open_flags((*C.GtkPlacesSidebar)(recv.native), c_flags)

	return
}

// SetShowConnectToServer is a wrapper around the C function gtk_places_sidebar_set_show_connect_to_server.
func (recv *PlacesSidebar) SetShowConnectToServer(showConnectToServer bool) {
	c_show_connect_to_server :=
		boolToGboolean(showConnectToServer)

	C.gtk_places_sidebar_set_show_connect_to_server((*C.GtkPlacesSidebar)(recv.native), c_show_connect_to_server)

	return
}

// SetShowDesktop is a wrapper around the C function gtk_places_sidebar_set_show_desktop.
func (recv *PlacesSidebar) SetShowDesktop(showDesktop bool) {
	c_show_desktop :=
		boolToGboolean(showDesktop)

	C.gtk_places_sidebar_set_show_desktop((*C.GtkPlacesSidebar)(recv.native), c_show_desktop)

	return
}

// Unsupported : gtk_print_settings_new_from_gvariant : unsupported parameter variant : Blacklisted record : GVariant

// Unsupported : gtk_recent_chooser_dialog_new : unsupported parameter ... : varargs

// Unsupported : gtk_recent_chooser_dialog_new_for_manager : unsupported parameter ... : varargs

// RevealerNew is a wrapper around the C function gtk_revealer_new.
func RevealerNew() *Revealer {
	retC := C.gtk_revealer_new()
	retGo := RevealerNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetChildRevealed is a wrapper around the C function gtk_revealer_get_child_revealed.
func (recv *Revealer) GetChildRevealed() bool {
	retC := C.gtk_revealer_get_child_revealed((*C.GtkRevealer)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetRevealChild is a wrapper around the C function gtk_revealer_get_reveal_child.
func (recv *Revealer) GetRevealChild() bool {
	retC := C.gtk_revealer_get_reveal_child((*C.GtkRevealer)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetTransitionDuration is a wrapper around the C function gtk_revealer_get_transition_duration.
func (recv *Revealer) GetTransitionDuration() uint32 {
	retC := C.gtk_revealer_get_transition_duration((*C.GtkRevealer)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// GetTransitionType is a wrapper around the C function gtk_revealer_get_transition_type.
func (recv *Revealer) GetTransitionType() RevealerTransitionType {
	retC := C.gtk_revealer_get_transition_type((*C.GtkRevealer)(recv.native))
	retGo := (RevealerTransitionType)(retC)

	return retGo
}

// SetRevealChild is a wrapper around the C function gtk_revealer_set_reveal_child.
func (recv *Revealer) SetRevealChild(revealChild bool) {
	c_reveal_child :=
		boolToGboolean(revealChild)

	C.gtk_revealer_set_reveal_child((*C.GtkRevealer)(recv.native), c_reveal_child)

	return
}

// SetTransitionDuration is a wrapper around the C function gtk_revealer_set_transition_duration.
func (recv *Revealer) SetTransitionDuration(duration uint32) {
	c_duration := (C.guint)(duration)

	C.gtk_revealer_set_transition_duration((*C.GtkRevealer)(recv.native), c_duration)

	return
}

// SetTransitionType is a wrapper around the C function gtk_revealer_set_transition_type.
func (recv *Revealer) SetTransitionType(transition RevealerTransitionType) {
	c_transition := (C.GtkRevealerTransitionType)(transition)

	C.gtk_revealer_set_transition_type((*C.GtkRevealer)(recv.native), c_transition)

	return
}

// Unsupported : gtk_scale_button_new : unsupported parameter size : no type generator for gint (GtkIconSize) for param size

// SearchBarNew is a wrapper around the C function gtk_search_bar_new.
func SearchBarNew() *SearchBar {
	retC := C.gtk_search_bar_new()
	retGo := SearchBarNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ConnectEntry is a wrapper around the C function gtk_search_bar_connect_entry.
func (recv *SearchBar) ConnectEntry(entry *Entry) {
	c_entry := (*C.GtkEntry)(entry.ToC())

	C.gtk_search_bar_connect_entry((*C.GtkSearchBar)(recv.native), c_entry)

	return
}

// GetSearchMode is a wrapper around the C function gtk_search_bar_get_search_mode.
func (recv *SearchBar) GetSearchMode() bool {
	retC := C.gtk_search_bar_get_search_mode((*C.GtkSearchBar)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetShowCloseButton is a wrapper around the C function gtk_search_bar_get_show_close_button.
func (recv *SearchBar) GetShowCloseButton() bool {
	retC := C.gtk_search_bar_get_show_close_button((*C.GtkSearchBar)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_search_bar_handle_event : unsupported parameter event : no type generator for Gdk.Event (GdkEvent*) for param event

// SetSearchMode is a wrapper around the C function gtk_search_bar_set_search_mode.
func (recv *SearchBar) SetSearchMode(searchMode bool) {
	c_search_mode :=
		boolToGboolean(searchMode)

	C.gtk_search_bar_set_search_mode((*C.GtkSearchBar)(recv.native), c_search_mode)

	return
}

// SetShowCloseButton is a wrapper around the C function gtk_search_bar_set_show_close_button.
func (recv *SearchBar) SetShowCloseButton(visible bool) {
	c_visible :=
		boolToGboolean(visible)

	C.gtk_search_bar_set_show_close_button((*C.GtkSearchBar)(recv.native), c_visible)

	return
}

type signalSearchEntrySearchChangedDetail struct {
	callback  SearchEntrySignalSearchChangedCallback
	handlerID C.gulong
}

var signalSearchEntrySearchChangedId int
var signalSearchEntrySearchChangedMap = make(map[int]signalSearchEntrySearchChangedDetail)
var signalSearchEntrySearchChangedLock sync.Mutex

// SearchEntrySignalSearchChangedCallback is a callback function for a 'search-changed' signal emitted from a SearchEntry.
type SearchEntrySignalSearchChangedCallback func()

/*
ConnectSearchChanged connects the callback to the 'search-changed' signal for the SearchEntry.

The returned value represents the connection, and may be passed to DisconnectSearchChanged to remove it.
*/
func (recv *SearchEntry) ConnectSearchChanged(callback SearchEntrySignalSearchChangedCallback) int {
	signalSearchEntrySearchChangedLock.Lock()
	defer signalSearchEntrySearchChangedLock.Unlock()

	signalSearchEntrySearchChangedId++
	instance := C.gpointer(recv.native)
	handlerID := C.SearchEntry_signal_connect_search_changed(instance, C.gpointer(uintptr(signalSearchEntrySearchChangedId)))

	detail := signalSearchEntrySearchChangedDetail{callback, handlerID}
	signalSearchEntrySearchChangedMap[signalSearchEntrySearchChangedId] = detail

	return signalSearchEntrySearchChangedId
}

/*
DisconnectSearchChanged disconnects a callback from the 'search-changed' signal for the SearchEntry.

The connectionID should be a value returned from a call to ConnectSearchChanged.
*/
func (recv *SearchEntry) DisconnectSearchChanged(connectionID int) {
	signalSearchEntrySearchChangedLock.Lock()
	defer signalSearchEntrySearchChangedLock.Unlock()

	detail, exists := signalSearchEntrySearchChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalSearchEntrySearchChangedMap, connectionID)
}

//export searchentry_searchChangedHandler
func searchentry_searchChangedHandler(_ *C.GObject, data C.gpointer) {
	index := int(uintptr(data))
	callback := signalSearchEntrySearchChangedMap[index].callback
	callback()
}

// StackNew is a wrapper around the C function gtk_stack_new.
func StackNew() *Stack {
	retC := C.gtk_stack_new()
	retGo := StackNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AddNamed is a wrapper around the C function gtk_stack_add_named.
func (recv *Stack) AddNamed(child *Widget, name string) {
	c_child := (*C.GtkWidget)(child.ToC())

	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	C.gtk_stack_add_named((*C.GtkStack)(recv.native), c_child, c_name)

	return
}

// AddTitled is a wrapper around the C function gtk_stack_add_titled.
func (recv *Stack) AddTitled(child *Widget, name string, title string) {
	c_child := (*C.GtkWidget)(child.ToC())

	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_title := C.CString(title)
	defer C.free(unsafe.Pointer(c_title))

	C.gtk_stack_add_titled((*C.GtkStack)(recv.native), c_child, c_name, c_title)

	return
}

// GetHomogeneous is a wrapper around the C function gtk_stack_get_homogeneous.
func (recv *Stack) GetHomogeneous() bool {
	retC := C.gtk_stack_get_homogeneous((*C.GtkStack)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetTransitionDuration is a wrapper around the C function gtk_stack_get_transition_duration.
func (recv *Stack) GetTransitionDuration() uint32 {
	retC := C.gtk_stack_get_transition_duration((*C.GtkStack)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// GetTransitionType is a wrapper around the C function gtk_stack_get_transition_type.
func (recv *Stack) GetTransitionType() StackTransitionType {
	retC := C.gtk_stack_get_transition_type((*C.GtkStack)(recv.native))
	retGo := (StackTransitionType)(retC)

	return retGo
}

// GetVisibleChild is a wrapper around the C function gtk_stack_get_visible_child.
func (recv *Stack) GetVisibleChild() *Widget {
	retC := C.gtk_stack_get_visible_child((*C.GtkStack)(recv.native))
	var retGo (*Widget)
	if retC == nil {
		retGo = nil
	} else {
		retGo = WidgetNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetVisibleChildName is a wrapper around the C function gtk_stack_get_visible_child_name.
func (recv *Stack) GetVisibleChildName() string {
	retC := C.gtk_stack_get_visible_child_name((*C.GtkStack)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// SetHomogeneous is a wrapper around the C function gtk_stack_set_homogeneous.
func (recv *Stack) SetHomogeneous(homogeneous bool) {
	c_homogeneous :=
		boolToGboolean(homogeneous)

	C.gtk_stack_set_homogeneous((*C.GtkStack)(recv.native), c_homogeneous)

	return
}

// SetTransitionDuration is a wrapper around the C function gtk_stack_set_transition_duration.
func (recv *Stack) SetTransitionDuration(duration uint32) {
	c_duration := (C.guint)(duration)

	C.gtk_stack_set_transition_duration((*C.GtkStack)(recv.native), c_duration)

	return
}

// SetTransitionType is a wrapper around the C function gtk_stack_set_transition_type.
func (recv *Stack) SetTransitionType(transition StackTransitionType) {
	c_transition := (C.GtkStackTransitionType)(transition)

	C.gtk_stack_set_transition_type((*C.GtkStack)(recv.native), c_transition)

	return
}

// SetVisibleChild is a wrapper around the C function gtk_stack_set_visible_child.
func (recv *Stack) SetVisibleChild(child *Widget) {
	c_child := (*C.GtkWidget)(child.ToC())

	C.gtk_stack_set_visible_child((*C.GtkStack)(recv.native), c_child)

	return
}

// SetVisibleChildFull is a wrapper around the C function gtk_stack_set_visible_child_full.
func (recv *Stack) SetVisibleChildFull(name string, transition StackTransitionType) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_transition := (C.GtkStackTransitionType)(transition)

	C.gtk_stack_set_visible_child_full((*C.GtkStack)(recv.native), c_name, c_transition)

	return
}

// SetVisibleChildName is a wrapper around the C function gtk_stack_set_visible_child_name.
func (recv *Stack) SetVisibleChildName(name string) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	C.gtk_stack_set_visible_child_name((*C.GtkStack)(recv.native), c_name)

	return
}

// StackSwitcherNew is a wrapper around the C function gtk_stack_switcher_new.
func StackSwitcherNew() *StackSwitcher {
	retC := C.gtk_stack_switcher_new()
	retGo := StackSwitcherNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetStack is a wrapper around the C function gtk_stack_switcher_get_stack.
func (recv *StackSwitcher) GetStack() *Stack {
	retC := C.gtk_stack_switcher_get_stack((*C.GtkStackSwitcher)(recv.native))
	var retGo (*Stack)
	if retC == nil {
		retGo = nil
	} else {
		retGo = StackNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// SetStack is a wrapper around the C function gtk_stack_switcher_set_stack.
func (recv *StackSwitcher) SetStack(stack *Stack) {
	c_stack := (*C.GtkStack)(stack.ToC())

	C.gtk_stack_switcher_set_stack((*C.GtkStackSwitcher)(recv.native), c_stack)

	return
}

// Unsupported : gtk_status_icon_new_from_gicon : unsupported parameter icon : no type generator for Gio.Icon (GIcon*) for param icon

// GetScale is a wrapper around the C function gtk_style_context_get_scale.
func (recv *StyleContext) GetScale() int32 {
	retC := C.gtk_style_context_get_scale((*C.GtkStyleContext)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// SetScale is a wrapper around the C function gtk_style_context_set_scale.
func (recv *StyleContext) SetScale(scale int32) {
	c_scale := (C.gint)(scale)

	C.gtk_style_context_set_scale((*C.GtkStyleContext)(recv.native), c_scale)

	return
}

// Unsupported : gtk_tree_store_new : unsupported parameter ... : varargs

// Unsupported : gtk_tree_store_newv : unsupported parameter types : no type generator for GType (GType) for array param types

// Unsupported : gtk_tree_view_new_with_model : unsupported parameter model : no type generator for TreeModel (GtkTreeModel*) for param model

// Unsupported : gtk_tree_view_column_new_with_attributes : unsupported parameter ... : varargs

// Unsupported : gtk_widget_new : unsupported parameter ... : varargs

// Unsupported : gtk_drag_begin_with_coordinates : unsupported parameter event : no type generator for Gdk.Event (GdkEvent*) for param event

// GetAllocatedBaseline is a wrapper around the C function gtk_widget_get_allocated_baseline.
func (recv *Widget) GetAllocatedBaseline() int32 {
	retC := C.gtk_widget_get_allocated_baseline((*C.GtkWidget)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetPreferredHeightAndBaselineForWidth is a wrapper around the C function gtk_widget_get_preferred_height_and_baseline_for_width.
func (recv *Widget) GetPreferredHeightAndBaselineForWidth(width int32) (int32, int32, int32, int32) {
	c_width := (C.gint)(width)

	var c_minimum_height C.gint

	var c_natural_height C.gint

	var c_minimum_baseline C.gint

	var c_natural_baseline C.gint

	C.gtk_widget_get_preferred_height_and_baseline_for_width((*C.GtkWidget)(recv.native), c_width, &c_minimum_height, &c_natural_height, &c_minimum_baseline, &c_natural_baseline)

	minimumHeight := (int32)(c_minimum_height)

	naturalHeight := (int32)(c_natural_height)

	minimumBaseline := (int32)(c_minimum_baseline)

	naturalBaseline := (int32)(c_natural_baseline)

	return minimumHeight, naturalHeight, minimumBaseline, naturalBaseline
}

// GetScaleFactor is a wrapper around the C function gtk_widget_get_scale_factor.
func (recv *Widget) GetScaleFactor() int32 {
	retC := C.gtk_widget_get_scale_factor((*C.GtkWidget)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetValignWithBaseline is a wrapper around the C function gtk_widget_get_valign_with_baseline.
func (recv *Widget) GetValignWithBaseline() Align {
	retC := C.gtk_widget_get_valign_with_baseline((*C.GtkWidget)(recv.native))
	retGo := (Align)(retC)

	return retGo
}

// InitTemplate is a wrapper around the C function gtk_widget_init_template.
func (recv *Widget) InitTemplate() {
	C.gtk_widget_init_template((*C.GtkWidget)(recv.native))

	return
}

// Unsupported : gtk_widget_size_allocate_with_baseline : unsupported parameter allocation : Blacklisted record : GdkRectangle

// Close is a wrapper around the C function gtk_window_close.
func (recv *Window) Close() {
	C.gtk_window_close((*C.GtkWindow)(recv.native))

	return
}

// SetTitlebar is a wrapper around the C function gtk_window_set_titlebar.
func (recv *Window) SetTitlebar(titlebar *Widget) {
	c_titlebar := (*C.GtkWidget)(titlebar.ToC())

	C.gtk_window_set_titlebar((*C.GtkWindow)(recv.native), c_titlebar)

	return
}
