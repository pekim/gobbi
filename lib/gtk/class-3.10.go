// This is a generated file - DO NOT EDIT
// +build gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	atk "github.com/pekim/gobbi/lib/atk"
	cairo "github.com/pekim/gobbi/lib/cairo"
	gdk "github.com/pekim/gobbi/lib/gdk"
	gdkpixbuf "github.com/pekim/gobbi/lib/gdkpixbuf"
	gio "github.com/pekim/gobbi/lib/gio"
	glib "github.com/pekim/gobbi/lib/glib"
	gobject "github.com/pekim/gobbi/lib/gobject"
	pango "github.com/pekim/gobbi/lib/pango"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

func (recv *AboutDialog) Dialog() *Dialog {}

func (recv *AccelGroup) Object() *gobject.Object {}

func (recv *AccelLabel) Label() *Label {}

func (recv *AccelMap) Object() *gobject.Object {}

func (recv *Accessible) Object() *atk.Object {}

func (recv *Action) Object() *gobject.Object {}

func (recv *ActionBar) Bin() *Bin {}

func (recv *ActionGroup) Object() *gobject.Object {}

func (recv *Adjustment) InitiallyUnowned() *gobject.InitiallyUnowned {}

func (recv *Alignment) Bin() *Bin {}

func (recv *AppChooserButton) ComboBox() *ComboBox {}

// Unsupported : gtk_app_chooser_dialog_new : unsupported parameter file : no type generator for Gio.File, GFile*

func (recv *AppChooserDialog) Dialog() *Dialog {}

func (recv *AppChooserWidget) Box() *Box {}

func (recv *Application) Application() *gio.Application {}

func (recv *ApplicationWindow) Window() *Window {}

func (recv *Arrow) Misc() *Misc {}

func (recv *ArrowAccessible) WidgetAccessible() *WidgetAccessible {}

func (recv *AspectFrame) Frame() *Frame {}

func (recv *Assistant) Window() *Window {}

func (recv *Bin) Container() *Container {}

func (recv *BooleanCellAccessible) RendererCellAccessible() *RendererCellAccessible {}

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

func (recv *Box) Container() *Container {}

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

// Unsupported : gtk_builder_add_callback_symbol : unsupported parameter callback_symbol : no type generator for GObject.Callback, GCallback

// Unsupported : gtk_builder_add_callback_symbols : unsupported parameter first_callback_symbol : no type generator for GObject.Callback, GCallback

// GetApplication is a wrapper around the C function gtk_builder_get_application.
func (recv *Builder) GetApplication() *Application {
	retC := C.gtk_builder_get_application((*C.GtkBuilder)(recv.native))
	retGo := ApplicationNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_builder_lookup_callback_symbol : no return generator

// SetApplication is a wrapper around the C function gtk_builder_set_application.
func (recv *Builder) SetApplication(application *Application) {
	c_application := (*C.GtkApplication)(application.ToC())

	C.gtk_builder_set_application((*C.GtkBuilder)(recv.native), c_application)

	return
}

func (recv *Builder) Object() *gobject.Object {}

// Unsupported : gtk_button_new_from_icon_name : unsupported parameter size : no type generator for gint, GtkIconSize

func (recv *Button) Bin() *Bin {}

func (recv *ButtonAccessible) ContainerAccessible() *ContainerAccessible {}

func (recv *ButtonBox) Box() *Box {}

func (recv *Calendar) Widget() *Widget {}

func (recv *CellAccessible) Accessible() *Accessible {}

func (recv *CellArea) InitiallyUnowned() *gobject.InitiallyUnowned {}

func (recv *CellAreaBox) CellArea() *CellArea {}

func (recv *CellAreaContext) Object() *gobject.Object {}

func (recv *CellRenderer) InitiallyUnowned() *gobject.InitiallyUnowned {}

func (recv *CellRendererAccel) CellRendererText() *CellRendererText {}

func (recv *CellRendererCombo) CellRendererText() *CellRendererText {}

func (recv *CellRendererPixbuf) CellRenderer() *CellRenderer {}

func (recv *CellRendererProgress) CellRenderer() *CellRenderer {}

func (recv *CellRendererSpin) CellRendererText() *CellRendererText {}

func (recv *CellRendererSpinner) CellRenderer() *CellRenderer {}

func (recv *CellRendererText) CellRenderer() *CellRenderer {}

func (recv *CellRendererToggle) CellRenderer() *CellRenderer {}

func (recv *CellView) Widget() *Widget {}

func (recv *CheckButton) ToggleButton() *ToggleButton {}

func (recv *CheckMenuItem) MenuItem() *MenuItem {}

func (recv *CheckMenuItemAccessible) MenuItemAccessible() *MenuItemAccessible {}

func (recv *Clipboard) Object() *gobject.Object {}

func (recv *ColorButton) Button() *Button {}

func (recv *ColorChooserDialog) Dialog() *Dialog {}

func (recv *ColorChooserWidget) Box() *Box {}

func (recv *ColorSelection) Box() *Box {}

func (recv *ColorSelectionDialog) Dialog() *Dialog {}

// Unsupported : gtk_combo_box_new_with_model : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

// Unsupported : gtk_combo_box_new_with_model_and_entry : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

func (recv *ComboBox) Bin() *Bin {}

func (recv *ComboBoxAccessible) ContainerAccessible() *ContainerAccessible {}

func (recv *ComboBoxText) ComboBox() *ComboBox {}

func (recv *Container) Widget() *Widget {}

func (recv *ContainerAccessible) WidgetAccessible() *WidgetAccessible {}

func (recv *ContainerCellAccessible) CellAccessible() *CellAccessible {}

func (recv *CssProvider) Object() *gobject.Object {}

// Unsupported : gtk_dialog_new_with_buttons : unsupported parameter ... : varargs

func (recv *Dialog) Window() *Window {}

func (recv *DrawingArea) Widget() *Widget {}

// GetTabs is a wrapper around the C function gtk_entry_get_tabs.
func (recv *Entry) GetTabs() *pango.TabArray {
	retC := C.gtk_entry_get_tabs((*C.GtkEntry)(recv.native))
	retGo := pango.TabArrayNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SetTabs is a wrapper around the C function gtk_entry_set_tabs.
func (recv *Entry) SetTabs(tabs *pango.TabArray) {
	c_tabs := (*C.PangoTabArray)(tabs.ToC())

	C.gtk_entry_set_tabs((*C.GtkEntry)(recv.native), c_tabs)

	return
}

func (recv *Entry) Widget() *Widget {}

func (recv *EntryAccessible) WidgetAccessible() *WidgetAccessible {}

func (recv *EntryBuffer) Object() *gobject.Object {}

func (recv *EntryCompletion) Object() *gobject.Object {}

// Unsupported : EntryIconAccessible : no CType

func (recv *EventBox) Bin() *Bin {}

func (recv *EventController) Object() *gobject.Object {}

func (recv *Expander) Bin() *Bin {}

func (recv *ExpanderAccessible) ContainerAccessible() *ContainerAccessible {}

func (recv *FileChooserButton) Box() *Box {}

// Unsupported : gtk_file_chooser_dialog_new : unsupported parameter ... : varargs

func (recv *FileChooserDialog) Dialog() *Dialog {}

func (recv *FileChooserNative) NativeDialog() *NativeDialog {}

func (recv *FileChooserWidget) Box() *Box {}

// Unsupported : gtk_file_filter_new_from_gvariant : unsupported parameter variant : Blacklisted record : GVariant

func (recv *FileFilter) InitiallyUnowned() *gobject.InitiallyUnowned {}

func (recv *Fixed) Container() *Container {}

func (recv *FlowBox) Container() *Container {}

func (recv *FlowBoxAccessible) ContainerAccessible() *ContainerAccessible {}

func (recv *FlowBoxChild) Bin() *Bin {}

func (recv *FlowBoxChildAccessible) ContainerAccessible() *ContainerAccessible {}

func (recv *FontButton) Button() *Button {}

func (recv *FontChooserDialog) Dialog() *Dialog {}

func (recv *FontChooserWidget) Box() *Box {}

func (recv *FontSelection) Box() *Box {}

func (recv *FontSelectionDialog) Dialog() *Dialog {}

func (recv *Frame) Bin() *Bin {}

func (recv *FrameAccessible) ContainerAccessible() *ContainerAccessible {}

func (recv *GLArea) Widget() *Widget {}

func (recv *Gesture) EventController() *EventController {}

func (recv *GestureDrag) GestureSingle() *GestureSingle {}

func (recv *GestureLongPress) GestureSingle() *GestureSingle {}

func (recv *GestureMultiPress) GestureSingle() *GestureSingle {}

func (recv *GesturePan) GestureDrag() *GestureDrag {}

func (recv *GestureRotate) Gesture() *Gesture {}

func (recv *GestureSingle) Gesture() *Gesture {}

func (recv *GestureSwipe) GestureSingle() *GestureSingle {}

func (recv *GestureZoom) Gesture() *Gesture {}

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

func (recv *Grid) Container() *Container {}

func (recv *HBox) Box() *Box {}

func (recv *HButtonBox) ButtonBox() *ButtonBox {}

func (recv *HPaned) Paned() *Paned {}

func (recv *HSV) Widget() *Widget {}

func (recv *HScale) Scale() *Scale {}

func (recv *HScrollbar) Scrollbar() *Scrollbar {}

func (recv *HSeparator) Separator() *Separator {}

func (recv *HandleBox) Bin() *Bin {}

// HeaderBarNew is a wrapper around the C function gtk_header_bar_new.
func HeaderBarNew() *HeaderBar {
	retC := C.gtk_header_bar_new()
	retGo := HeaderBarNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetCustomTitle is a wrapper around the C function gtk_header_bar_get_custom_title.
func (recv *HeaderBar) GetCustomTitle() *Widget {
	retC := C.gtk_header_bar_get_custom_title((*C.GtkHeaderBar)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

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

func (recv *HeaderBar) Container() *Container {}

func (recv *IMContext) Object() *gobject.Object {}

func (recv *IMContextSimple) IMContext() *IMContext {}

func (recv *IMMulticontext) IMContext() *IMContext {}

func (recv *IconFactory) Object() *gobject.Object {}

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

func (recv *IconInfo) Object() *gobject.Object {}

// Unsupported : gtk_icon_theme_choose_icon_for_scale : unsupported parameter icon_names : no param type

// LoadIconForScale is a wrapper around the C function gtk_icon_theme_load_icon_for_scale.
func (recv *IconTheme) LoadIconForScale(iconName string, size int32, scale int32, flags IconLookupFlags) (*gdkpixbuf.Pixbuf, error) {
	c_icon_name := C.CString(iconName)
	defer C.free(unsafe.Pointer(c_icon_name))

	c_size := (C.gint)(size)

	c_scale := (C.gint)(scale)

	c_flags := (C.GtkIconLookupFlags)(flags)

	var cThrowableError *C.GError

	retC := C.gtk_icon_theme_load_icon_for_scale((*C.GtkIconTheme)(recv.native), c_icon_name, c_size, c_scale, c_flags, &cThrowableError)
	retGo := gdkpixbuf.PixbufNewFromC(unsafe.Pointer(retC))

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
	retGo := cairo.SurfaceNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : gtk_icon_theme_lookup_by_gicon_for_scale : unsupported parameter icon : no type generator for Gio.Icon, GIcon*

// LookupIconForScale is a wrapper around the C function gtk_icon_theme_lookup_icon_for_scale.
func (recv *IconTheme) LookupIconForScale(iconName string, size int32, scale int32, flags IconLookupFlags) *IconInfo {
	c_icon_name := C.CString(iconName)
	defer C.free(unsafe.Pointer(c_icon_name))

	c_size := (C.gint)(size)

	c_scale := (C.gint)(scale)

	c_flags := (C.GtkIconLookupFlags)(flags)

	retC := C.gtk_icon_theme_lookup_icon_for_scale((*C.GtkIconTheme)(recv.native), c_icon_name, c_size, c_scale, c_flags)
	retGo := IconInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

func (recv *IconTheme) Object() *gobject.Object {}

// Unsupported : gtk_icon_view_new_with_model : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

func (recv *IconView) Container() *Container {}

func (recv *IconViewAccessible) ContainerAccessible() *ContainerAccessible {}

// Unsupported : gtk_image_new_from_gicon : unsupported parameter icon : no type generator for Gio.Icon, GIcon*

// Unsupported : gtk_image_new_from_icon_name : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_image_new_from_icon_set : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_image_new_from_stock : unsupported parameter size : no type generator for gint, GtkIconSize

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

func (recv *Image) Misc() *Misc {}

func (recv *ImageAccessible) WidgetAccessible() *WidgetAccessible {}

func (recv *ImageCellAccessible) RendererCellAccessible() *RendererCellAccessible {}

func (recv *ImageMenuItem) MenuItem() *MenuItem {}

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

func (recv *InfoBar) Box() *Box {}

func (recv *Invisible) Widget() *Widget {}

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

func (recv *Label) Misc() *Misc {}

func (recv *LabelAccessible) WidgetAccessible() *WidgetAccessible {}

func (recv *Layout) Container() *Container {}

func (recv *LevelBar) Widget() *Widget {}

func (recv *LevelBarAccessible) WidgetAccessible() *WidgetAccessible {}

func (recv *LinkButton) Button() *Button {}

func (recv *LinkButtonAccessible) ButtonAccessible() *ButtonAccessible {}

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
	retGo := ListBoxRowNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetRowAtY is a wrapper around the C function gtk_list_box_get_row_at_y.
func (recv *ListBox) GetRowAtY(y int32) *ListBoxRow {
	c_y := (C.gint)(y)

	retC := C.gtk_list_box_get_row_at_y((*C.GtkListBox)(recv.native), c_y)
	retGo := ListBoxRowNewFromC(unsafe.Pointer(retC))

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

// Unsupported : gtk_list_box_set_filter_func : unsupported parameter filter_func : no type generator for ListBoxFilterFunc, GtkListBoxFilterFunc

// Unsupported : gtk_list_box_set_header_func : unsupported parameter update_header : no type generator for ListBoxUpdateHeaderFunc, GtkListBoxUpdateHeaderFunc

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

// Unsupported : gtk_list_box_set_sort_func : unsupported parameter sort_func : no type generator for ListBoxSortFunc, GtkListBoxSortFunc

func (recv *ListBox) Container() *Container {}

func (recv *ListBoxAccessible) ContainerAccessible() *ContainerAccessible {}

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
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

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

func (recv *ListBoxRow) Bin() *Bin {}

func (recv *ListBoxRowAccessible) ContainerAccessible() *ContainerAccessible {}

// Unsupported : gtk_list_store_new : unsupported parameter ... : varargs

// Unsupported : gtk_list_store_newv : unsupported parameter types : no param type

func (recv *ListStore) Object() *gobject.Object {}

func (recv *LockButton) Button() *Button {}

func (recv *LockButtonAccessible) ButtonAccessible() *ButtonAccessible {}

func (recv *Menu) MenuShell() *MenuShell {}

func (recv *MenuAccessible) MenuShellAccessible() *MenuShellAccessible {}

func (recv *MenuBar) MenuShell() *MenuShell {}

func (recv *MenuButton) ToggleButton() *ToggleButton {}

func (recv *MenuButtonAccessible) ToggleButtonAccessible() *ToggleButtonAccessible {}

func (recv *MenuItem) Bin() *Bin {}

func (recv *MenuItemAccessible) ContainerAccessible() *ContainerAccessible {}

func (recv *MenuShell) Container() *Container {}

func (recv *MenuShellAccessible) ContainerAccessible() *ContainerAccessible {}

func (recv *MenuToolButton) ToolButton() *ToolButton {}

// Unsupported : gtk_message_dialog_new : unsupported parameter ... : varargs

// Unsupported : gtk_message_dialog_new_with_markup : unsupported parameter ... : varargs

func (recv *MessageDialog) Dialog() *Dialog {}

func (recv *Misc) Widget() *Widget {}

func (recv *ModelButton) Button() *Button {}

func (recv *MountOperation) MountOperation() *gio.MountOperation {}

func (recv *NativeDialog) Object() *gobject.Object {}

func (recv *Notebook) Container() *Container {}

func (recv *NotebookAccessible) ContainerAccessible() *ContainerAccessible {}

func (recv *NotebookPageAccessible) Object() *atk.Object {}

func (recv *NumerableIcon) EmblemedIcon() *gio.EmblemedIcon {}

func (recv *OffscreenWindow) Window() *Window {}

func (recv *Overlay) Bin() *Bin {}

// Unsupported : gtk_pad_controller_new : unsupported parameter group : no type generator for Gio.ActionGroup, GActionGroup*

func (recv *PadController) EventController() *EventController {}

// Unsupported : gtk_page_setup_new_from_gvariant : unsupported parameter variant : Blacklisted record : GVariant

func (recv *PageSetup) Object() *gobject.Object {}

func (recv *Paned) Container() *Container {}

func (recv *PanedAccessible) ContainerAccessible() *ContainerAccessible {}

// PlacesSidebarNew is a wrapper around the C function gtk_places_sidebar_new.
func PlacesSidebarNew() *PlacesSidebar {
	retC := C.gtk_places_sidebar_new()
	retGo := PlacesSidebarNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_places_sidebar_add_shortcut : unsupported parameter location : no type generator for Gio.File, GFile*

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

// Unsupported : gtk_places_sidebar_remove_shortcut : unsupported parameter location : no type generator for Gio.File, GFile*

// Unsupported : gtk_places_sidebar_set_location : unsupported parameter location : no type generator for Gio.File, GFile*

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

func (recv *PlacesSidebar) ScrolledWindow() *ScrolledWindow {}

func (recv *Popover) Bin() *Bin {}

func (recv *PopoverAccessible) ContainerAccessible() *ContainerAccessible {}

func (recv *PopoverMenu) Popover() *Popover {}

func (recv *PrintContext) Object() *gobject.Object {}

func (recv *PrintOperation) Object() *gobject.Object {}

// Unsupported : gtk_print_settings_new_from_gvariant : unsupported parameter variant : Blacklisted record : GVariant

func (recv *PrintSettings) Object() *gobject.Object {}

func (recv *ProgressBar) Widget() *Widget {}

func (recv *ProgressBarAccessible) WidgetAccessible() *WidgetAccessible {}

func (recv *RadioAction) ToggleAction() *ToggleAction {}

func (recv *RadioButton) CheckButton() *CheckButton {}

func (recv *RadioButtonAccessible) ToggleButtonAccessible() *ToggleButtonAccessible {}

func (recv *RadioMenuItem) CheckMenuItem() *CheckMenuItem {}

func (recv *RadioMenuItemAccessible) CheckMenuItemAccessible() *CheckMenuItemAccessible {}

func (recv *RadioToolButton) ToggleToolButton() *ToggleToolButton {}

func (recv *Range) Widget() *Widget {}

func (recv *RangeAccessible) WidgetAccessible() *WidgetAccessible {}

func (recv *RcStyle) Object() *gobject.Object {}

func (recv *RecentAction) Action() *Action {}

// Unsupported : gtk_recent_chooser_dialog_new : unsupported parameter ... : varargs

// Unsupported : gtk_recent_chooser_dialog_new_for_manager : unsupported parameter ... : varargs

func (recv *RecentChooserDialog) Dialog() *Dialog {}

func (recv *RecentChooserMenu) Menu() *Menu {}

func (recv *RecentChooserWidget) Box() *Box {}

func (recv *RecentFilter) InitiallyUnowned() *gobject.InitiallyUnowned {}

func (recv *RecentManager) Object() *gobject.Object {}

func (recv *RendererCellAccessible) CellAccessible() *CellAccessible {}

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

func (recv *Revealer) Bin() *Bin {}

func (recv *Scale) Range() *Range {}

func (recv *ScaleAccessible) RangeAccessible() *RangeAccessible {}

// Unsupported : gtk_scale_button_new : unsupported parameter size : no type generator for gint, GtkIconSize

func (recv *ScaleButton) Button() *Button {}

func (recv *ScaleButtonAccessible) ButtonAccessible() *ButtonAccessible {}

func (recv *Scrollbar) Range() *Range {}

func (recv *ScrolledWindow) Bin() *Bin {}

func (recv *ScrolledWindowAccessible) ContainerAccessible() *ContainerAccessible {}

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

// Unsupported : gtk_search_bar_handle_event : unsupported parameter event : no type generator for Gdk.Event, GdkEvent*

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

func (recv *SearchBar) Bin() *Bin {}

func (recv *SearchEntry) Entry() *Entry {}

func (recv *Separator) Widget() *Widget {}

func (recv *SeparatorMenuItem) MenuItem() *MenuItem {}

func (recv *SeparatorToolItem) ToolItem() *ToolItem {}

func (recv *Settings) Object() *gobject.Object {}

func (recv *ShortcutLabel) Box() *Box {}

func (recv *ShortcutsGroup) Box() *Box {}

func (recv *ShortcutsSection) Box() *Box {}

func (recv *ShortcutsShortcut) Box() *Box {}

func (recv *ShortcutsWindow) Window() *Window {}

func (recv *SizeGroup) Object() *gobject.Object {}

func (recv *SpinButton) Entry() *Entry {}

func (recv *SpinButtonAccessible) EntryAccessible() *EntryAccessible {}

func (recv *Spinner) Widget() *Widget {}

func (recv *SpinnerAccessible) WidgetAccessible() *WidgetAccessible {}

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
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

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

func (recv *Stack) Container() *Container {}

func (recv *StackSidebar) Bin() *Bin {}

// StackSwitcherNew is a wrapper around the C function gtk_stack_switcher_new.
func StackSwitcherNew() *StackSwitcher {
	retC := C.gtk_stack_switcher_new()
	retGo := StackSwitcherNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetStack is a wrapper around the C function gtk_stack_switcher_get_stack.
func (recv *StackSwitcher) GetStack() *Stack {
	retC := C.gtk_stack_switcher_get_stack((*C.GtkStackSwitcher)(recv.native))
	retGo := StackNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SetStack is a wrapper around the C function gtk_stack_switcher_set_stack.
func (recv *StackSwitcher) SetStack(stack *Stack) {
	c_stack := (*C.GtkStack)(stack.ToC())

	C.gtk_stack_switcher_set_stack((*C.GtkStackSwitcher)(recv.native), c_stack)

	return
}

func (recv *StackSwitcher) Box() *Box {}

// Unsupported : gtk_status_icon_new_from_gicon : unsupported parameter icon : no type generator for Gio.Icon, GIcon*

func (recv *StatusIcon) Object() *gobject.Object {}

func (recv *Statusbar) Box() *Box {}

func (recv *StatusbarAccessible) ContainerAccessible() *ContainerAccessible {}

func (recv *Style) Object() *gobject.Object {}

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

func (recv *StyleContext) Object() *gobject.Object {}

func (recv *StyleProperties) Object() *gobject.Object {}

func (recv *Switch) Widget() *Widget {}

func (recv *SwitchAccessible) WidgetAccessible() *WidgetAccessible {}

func (recv *Table) Container() *Container {}

func (recv *TearoffMenuItem) MenuItem() *MenuItem {}

func (recv *TextBuffer) Object() *gobject.Object {}

func (recv *TextCellAccessible) RendererCellAccessible() *RendererCellAccessible {}

func (recv *TextChildAnchor) Object() *gobject.Object {}

func (recv *TextMark) Object() *gobject.Object {}

func (recv *TextTag) Object() *gobject.Object {}

func (recv *TextTagTable) Object() *gobject.Object {}

func (recv *TextView) Container() *Container {}

func (recv *TextViewAccessible) ContainerAccessible() *ContainerAccessible {}

func (recv *ThemingEngine) Object() *gobject.Object {}

func (recv *ToggleAction) Action() *Action {}

func (recv *ToggleButton) Button() *Button {}

func (recv *ToggleButtonAccessible) ButtonAccessible() *ButtonAccessible {}

func (recv *ToggleToolButton) ToolButton() *ToolButton {}

func (recv *ToolButton) ToolItem() *ToolItem {}

func (recv *ToolItem) Bin() *Bin {}

func (recv *ToolItemGroup) Container() *Container {}

func (recv *ToolPalette) Container() *Container {}

func (recv *Toolbar) Container() *Container {}

func (recv *Tooltip) Object() *gobject.Object {}

func (recv *ToplevelAccessible) Object() *atk.Object {}

func (recv *TreeModelFilter) Object() *gobject.Object {}

func (recv *TreeModelSort) Object() *gobject.Object {}

func (recv *TreeSelection) Object() *gobject.Object {}

// Unsupported : gtk_tree_store_new : unsupported parameter ... : varargs

// Unsupported : gtk_tree_store_newv : unsupported parameter types : no param type

func (recv *TreeStore) Object() *gobject.Object {}

// Unsupported : gtk_tree_view_new_with_model : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

func (recv *TreeView) Container() *Container {}

func (recv *TreeViewAccessible) ContainerAccessible() *ContainerAccessible {}

// Unsupported : gtk_tree_view_column_new_with_attributes : unsupported parameter ... : varargs

func (recv *TreeViewColumn) InitiallyUnowned() *gobject.InitiallyUnowned {}

func (recv *UIManager) Object() *gobject.Object {}

func (recv *VBox) Box() *Box {}

func (recv *VButtonBox) ButtonBox() *ButtonBox {}

func (recv *VPaned) Paned() *Paned {}

func (recv *VScale) Scale() *Scale {}

func (recv *VScrollbar) Scrollbar() *Scrollbar {}

func (recv *VSeparator) Separator() *Separator {}

func (recv *Viewport) Bin() *Bin {}

func (recv *VolumeButton) ScaleButton() *ScaleButton {}

// Unsupported : gtk_widget_new : unsupported parameter type : no type generator for GType, GType

// Unsupported : gtk_drag_begin_with_coordinates : unsupported parameter event : no type generator for Gdk.Event, GdkEvent*

// GetAllocatedBaseline is a wrapper around the C function gtk_widget_get_allocated_baseline.
func (recv *Widget) GetAllocatedBaseline() int32 {
	retC := C.gtk_widget_get_allocated_baseline((*C.GtkWidget)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetPreferredHeightAndBaselineForWidth is a wrapper around the C function gtk_widget_get_preferred_height_and_baseline_for_width.
func (recv *Widget) GetPreferredHeightAndBaselineForWidth(width int32) (*int32, *int32, *int32, *int32) {
	c_width := (C.gint)(width)

	var c_minimum_height C.gint

	var c_natural_height C.gint

	var c_minimum_baseline C.gint

	var c_natural_baseline C.gint

	C.gtk_widget_get_preferred_height_and_baseline_for_width((*C.GtkWidget)(recv.native), c_width, &c_minimum_height, &c_natural_height, &c_minimum_baseline, &c_natural_baseline)

	minimumHeight := (*int32)(&c_minimum_height)

	naturalHeight := (*int32)(&c_natural_height)

	minimumBaseline := (*int32)(&c_minimum_baseline)

	naturalBaseline := (*int32)(&c_natural_baseline)

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

func (recv *Widget) InitiallyUnowned() *gobject.InitiallyUnowned {}

func (recv *WidgetAccessible) Accessible() *Accessible {}

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

func (recv *Window) Bin() *Bin {}

func (recv *WindowAccessible) ContainerAccessible() *ContainerAccessible {}

func (recv *WindowGroup) Object() *gobject.Object {}
