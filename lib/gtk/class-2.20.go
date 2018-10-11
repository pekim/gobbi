// This is a generated file - DO NOT EDIT
// +build gtk_2.20 gtk_2.22 gtk_2.24 gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	atk "github.com/pekim/gobbi/lib/atk"
	cairo "github.com/pekim/gobbi/lib/cairo"
	gdk "github.com/pekim/gobbi/lib/gdk"
	gdkpixbuf "github.com/pekim/gobbi/lib/gdkpixbuf"
	gio "github.com/pekim/gobbi/lib/gio"
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

// GetAlwaysShowImage is a wrapper around the C function gtk_action_get_always_show_image.
func (recv *Action) GetAlwaysShowImage() bool {
	retC := C.gtk_action_get_always_show_image((*C.GtkAction)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetAlwaysShowImage is a wrapper around the C function gtk_action_set_always_show_image.
func (recv *Action) SetAlwaysShowImage(alwaysShow bool) {
	c_always_show :=
		boolToGboolean(alwaysShow)

	C.gtk_action_set_always_show_image((*C.GtkAction)(recv.native), c_always_show)

	return
}

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

func (recv *Box) Container() *Container {}

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

// CellRendererSpinnerNew is a wrapper around the C function gtk_cell_renderer_spinner_new.
func CellRendererSpinnerNew() *CellRendererSpinner {
	retC := C.gtk_cell_renderer_spinner_new()
	retGo := CellRendererSpinnerNewFromC(unsafe.Pointer(retC))

	return retGo
}

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

// GetWidgetForResponse is a wrapper around the C function gtk_dialog_get_widget_for_response.
func (recv *Dialog) GetWidgetForResponse(responseId int32) *Widget {
	c_response_id := (C.gint)(responseId)

	retC := C.gtk_dialog_get_widget_for_response((*C.GtkDialog)(recv.native), c_response_id)
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

func (recv *Dialog) Window() *Window {}

func (recv *DrawingArea) Widget() *Widget {}

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

func (recv *Grid) Container() *Container {}

func (recv *HBox) Box() *Box {}

func (recv *HButtonBox) ButtonBox() *ButtonBox {}

func (recv *HPaned) Paned() *Paned {}

func (recv *HSV) Widget() *Widget {}

func (recv *HScale) Scale() *Scale {}

func (recv *HScrollbar) Scrollbar() *Scrollbar {}

func (recv *HSeparator) Separator() *Separator {}

func (recv *HandleBox) Bin() *Bin {}

func (recv *HeaderBar) Container() *Container {}

func (recv *IMContext) Object() *gobject.Object {}

func (recv *IMContextSimple) IMContext() *IMContext {}

func (recv *IMMulticontext) IMContext() *IMContext {}

func (recv *IconFactory) Object() *gobject.Object {}

func (recv *IconInfo) Object() *gobject.Object {}

func (recv *IconTheme) Object() *gobject.Object {}

// Unsupported : gtk_icon_view_new_with_model : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

func (recv *IconView) Container() *Container {}

func (recv *IconViewAccessible) ContainerAccessible() *ContainerAccessible {}

// Unsupported : gtk_image_new_from_gicon : unsupported parameter icon : no type generator for Gio.Icon, GIcon*

// Unsupported : gtk_image_new_from_icon_name : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_image_new_from_icon_set : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_image_new_from_stock : unsupported parameter size : no type generator for gint, GtkIconSize

func (recv *Image) Misc() *Misc {}

func (recv *ImageAccessible) WidgetAccessible() *WidgetAccessible {}

func (recv *ImageCellAccessible) RendererCellAccessible() *RendererCellAccessible {}

func (recv *ImageMenuItem) MenuItem() *MenuItem {}

// Unsupported : gtk_info_bar_new_with_buttons : unsupported parameter ... : varargs

func (recv *InfoBar) Box() *Box {}

func (recv *Invisible) Widget() *Widget {}

func (recv *Label) Misc() *Misc {}

func (recv *LabelAccessible) WidgetAccessible() *WidgetAccessible {}

func (recv *Layout) Container() *Container {}

func (recv *LevelBar) Widget() *Widget {}

func (recv *LevelBarAccessible) WidgetAccessible() *WidgetAccessible {}

func (recv *LinkButton) Button() *Button {}

func (recv *LinkButtonAccessible) ButtonAccessible() *ButtonAccessible {}

func (recv *ListBox) Container() *Container {}

func (recv *ListBoxAccessible) ContainerAccessible() *ContainerAccessible {}

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

// GetActionWidget is a wrapper around the C function gtk_notebook_get_action_widget.
func (recv *Notebook) GetActionWidget(packType PackType) *Widget {
	c_pack_type := (C.GtkPackType)(packType)

	retC := C.gtk_notebook_get_action_widget((*C.GtkNotebook)(recv.native), c_pack_type)
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SetActionWidget is a wrapper around the C function gtk_notebook_set_action_widget.
func (recv *Notebook) SetActionWidget(widget *Widget, packType PackType) {
	c_widget := (*C.GtkWidget)(widget.ToC())

	c_pack_type := (C.GtkPackType)(packType)

	C.gtk_notebook_set_action_widget((*C.GtkNotebook)(recv.native), c_widget, c_pack_type)

	return
}

func (recv *Notebook) Container() *Container {}

func (recv *NotebookAccessible) ContainerAccessible() *ContainerAccessible {}

func (recv *NotebookPageAccessible) Object() *atk.Object {}

func (recv *NumerableIcon) EmblemedIcon() *gio.EmblemedIcon {}

// OffscreenWindowNew is a wrapper around the C function gtk_offscreen_window_new.
func OffscreenWindowNew() *OffscreenWindow {
	retC := C.gtk_offscreen_window_new()
	retGo := OffscreenWindowNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetPixbuf is a wrapper around the C function gtk_offscreen_window_get_pixbuf.
func (recv *OffscreenWindow) GetPixbuf() *gdkpixbuf.Pixbuf {
	retC := C.gtk_offscreen_window_get_pixbuf((*C.GtkOffscreenWindow)(recv.native))
	retGo := gdkpixbuf.PixbufNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetSurface is a wrapper around the C function gtk_offscreen_window_get_surface.
func (recv *OffscreenWindow) GetSurface() *cairo.Surface {
	retC := C.gtk_offscreen_window_get_surface((*C.GtkOffscreenWindow)(recv.native))
	retGo := cairo.SurfaceNewFromC(unsafe.Pointer(retC))

	return retGo
}

func (recv *OffscreenWindow) Window() *Window {}

func (recv *Overlay) Bin() *Bin {}

// Unsupported : gtk_pad_controller_new : unsupported parameter group : no type generator for Gio.ActionGroup, GActionGroup*

func (recv *PadController) EventController() *EventController {}

// Unsupported : gtk_page_setup_new_from_gvariant : unsupported parameter variant : Blacklisted record : GVariant

func (recv *PageSetup) Object() *gobject.Object {}

// GetHandleWindow is a wrapper around the C function gtk_paned_get_handle_window.
func (recv *Paned) GetHandleWindow() *gdk.Window {
	retC := C.gtk_paned_get_handle_window((*C.GtkPaned)(recv.native))
	retGo := gdk.WindowNewFromC(unsafe.Pointer(retC))

	return retGo
}

func (recv *Paned) Container() *Container {}

func (recv *PanedAccessible) ContainerAccessible() *ContainerAccessible {}

func (recv *PlacesSidebar) ScrolledWindow() *ScrolledWindow {}

func (recv *Popover) Bin() *Bin {}

func (recv *PopoverAccessible) ContainerAccessible() *ContainerAccessible {}

func (recv *PopoverMenu) Popover() *Popover {}

// GetHardMargins is a wrapper around the C function gtk_print_context_get_hard_margins.
func (recv *PrintContext) GetHardMargins() (bool, *float64, *float64, *float64, *float64) {
	var c_top C.gdouble

	var c_bottom C.gdouble

	var c_left C.gdouble

	var c_right C.gdouble

	retC := C.gtk_print_context_get_hard_margins((*C.GtkPrintContext)(recv.native), &c_top, &c_bottom, &c_left, &c_right)
	retGo := retC == C.TRUE

	top := (*float64)(&c_top)

	bottom := (*float64)(&c_bottom)

	left := (*float64)(&c_left)

	right := (*float64)(&c_right)

	return retGo, top, bottom, left, right
}

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

// GetMinSliderSize is a wrapper around the C function gtk_range_get_min_slider_size.
func (recv *Range) GetMinSliderSize() int32 {
	retC := C.gtk_range_get_min_slider_size((*C.GtkRange)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : gtk_range_get_range_rect : unsupported parameter range_rect : Blacklisted record : GdkRectangle

// GetSliderRange is a wrapper around the C function gtk_range_get_slider_range.
func (recv *Range) GetSliderRange() (*int32, *int32) {
	var c_slider_start C.gint

	var c_slider_end C.gint

	C.gtk_range_get_slider_range((*C.GtkRange)(recv.native), &c_slider_start, &c_slider_end)

	sliderStart := (*int32)(&c_slider_start)

	sliderEnd := (*int32)(&c_slider_end)

	return sliderStart, sliderEnd
}

// GetSliderSizeFixed is a wrapper around the C function gtk_range_get_slider_size_fixed.
func (recv *Range) GetSliderSizeFixed() bool {
	retC := C.gtk_range_get_slider_size_fixed((*C.GtkRange)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetMinSliderSize is a wrapper around the C function gtk_range_set_min_slider_size.
func (recv *Range) SetMinSliderSize(minSize int32) {
	c_min_size := (C.gint)(minSize)

	C.gtk_range_set_min_slider_size((*C.GtkRange)(recv.native), c_min_size)

	return
}

// SetSliderSizeFixed is a wrapper around the C function gtk_range_set_slider_size_fixed.
func (recv *Range) SetSliderSizeFixed(sizeFixed bool) {
	c_size_fixed :=
		boolToGboolean(sizeFixed)

	C.gtk_range_set_slider_size_fixed((*C.GtkRange)(recv.native), c_size_fixed)

	return
}

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

func (recv *Revealer) Bin() *Bin {}

func (recv *Scale) Range() *Range {}

func (recv *ScaleAccessible) RangeAccessible() *RangeAccessible {}

// Unsupported : gtk_scale_button_new : unsupported parameter size : no type generator for gint, GtkIconSize

func (recv *ScaleButton) Button() *Button {}

func (recv *ScaleButtonAccessible) ButtonAccessible() *ButtonAccessible {}

func (recv *Scrollbar) Range() *Range {}

func (recv *ScrolledWindow) Bin() *Bin {}

func (recv *ScrolledWindowAccessible) ContainerAccessible() *ContainerAccessible {}

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

// SpinnerNew is a wrapper around the C function gtk_spinner_new.
func SpinnerNew() *Spinner {
	retC := C.gtk_spinner_new()
	retGo := SpinnerNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Start is a wrapper around the C function gtk_spinner_start.
func (recv *Spinner) Start() {
	C.gtk_spinner_start((*C.GtkSpinner)(recv.native))

	return
}

// Stop is a wrapper around the C function gtk_spinner_stop.
func (recv *Spinner) Stop() {
	C.gtk_spinner_stop((*C.GtkSpinner)(recv.native))

	return
}

func (recv *Spinner) Widget() *Widget {}

func (recv *SpinnerAccessible) WidgetAccessible() *WidgetAccessible {}

func (recv *Stack) Container() *Container {}

func (recv *StackSidebar) Bin() *Bin {}

func (recv *StackSwitcher) Box() *Box {}

// Unsupported : gtk_status_icon_new_from_gicon : unsupported parameter icon : no type generator for Gio.Icon, GIcon*

// SetName is a wrapper around the C function gtk_status_icon_set_name.
func (recv *StatusIcon) SetName(name string) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	C.gtk_status_icon_set_name((*C.GtkStatusIcon)(recv.native), c_name)

	return
}

func (recv *StatusIcon) Object() *gobject.Object {}

// GetMessageArea is a wrapper around the C function gtk_statusbar_get_message_area.
func (recv *Statusbar) GetMessageArea() *Box {
	retC := C.gtk_statusbar_get_message_area((*C.GtkStatusbar)(recv.native))
	retGo := BoxNewFromC(unsafe.Pointer(retC))

	return retGo
}

func (recv *Statusbar) Box() *Box {}

func (recv *StatusbarAccessible) ContainerAccessible() *ContainerAccessible {}

func (recv *Style) Object() *gobject.Object {}

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

// GetEllipsizeMode is a wrapper around the C function gtk_tool_item_get_ellipsize_mode.
func (recv *ToolItem) GetEllipsizeMode() pango.EllipsizeMode {
	retC := C.gtk_tool_item_get_ellipsize_mode((*C.GtkToolItem)(recv.native))
	retGo := (pango.EllipsizeMode)(retC)

	return retGo
}

// GetTextAlignment is a wrapper around the C function gtk_tool_item_get_text_alignment.
func (recv *ToolItem) GetTextAlignment() float32 {
	retC := C.gtk_tool_item_get_text_alignment((*C.GtkToolItem)(recv.native))
	retGo := (float32)(retC)

	return retGo
}

// GetTextOrientation is a wrapper around the C function gtk_tool_item_get_text_orientation.
func (recv *ToolItem) GetTextOrientation() Orientation {
	retC := C.gtk_tool_item_get_text_orientation((*C.GtkToolItem)(recv.native))
	retGo := (Orientation)(retC)

	return retGo
}

// GetTextSizeGroup is a wrapper around the C function gtk_tool_item_get_text_size_group.
func (recv *ToolItem) GetTextSizeGroup() *SizeGroup {
	retC := C.gtk_tool_item_get_text_size_group((*C.GtkToolItem)(recv.native))
	retGo := SizeGroupNewFromC(unsafe.Pointer(retC))

	return retGo
}

func (recv *ToolItem) Bin() *Bin {}

// ToolItemGroupNew is a wrapper around the C function gtk_tool_item_group_new.
func ToolItemGroupNew(label string) *ToolItemGroup {
	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	retC := C.gtk_tool_item_group_new(c_label)
	retGo := ToolItemGroupNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetCollapsed is a wrapper around the C function gtk_tool_item_group_get_collapsed.
func (recv *ToolItemGroup) GetCollapsed() bool {
	retC := C.gtk_tool_item_group_get_collapsed((*C.GtkToolItemGroup)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetDropItem is a wrapper around the C function gtk_tool_item_group_get_drop_item.
func (recv *ToolItemGroup) GetDropItem(x int32, y int32) *ToolItem {
	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	retC := C.gtk_tool_item_group_get_drop_item((*C.GtkToolItemGroup)(recv.native), c_x, c_y)
	retGo := ToolItemNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetEllipsize is a wrapper around the C function gtk_tool_item_group_get_ellipsize.
func (recv *ToolItemGroup) GetEllipsize() pango.EllipsizeMode {
	retC := C.gtk_tool_item_group_get_ellipsize((*C.GtkToolItemGroup)(recv.native))
	retGo := (pango.EllipsizeMode)(retC)

	return retGo
}

// GetHeaderRelief is a wrapper around the C function gtk_tool_item_group_get_header_relief.
func (recv *ToolItemGroup) GetHeaderRelief() ReliefStyle {
	retC := C.gtk_tool_item_group_get_header_relief((*C.GtkToolItemGroup)(recv.native))
	retGo := (ReliefStyle)(retC)

	return retGo
}

// GetItemPosition is a wrapper around the C function gtk_tool_item_group_get_item_position.
func (recv *ToolItemGroup) GetItemPosition(item *ToolItem) int32 {
	c_item := (*C.GtkToolItem)(item.ToC())

	retC := C.gtk_tool_item_group_get_item_position((*C.GtkToolItemGroup)(recv.native), c_item)
	retGo := (int32)(retC)

	return retGo
}

// GetLabel is a wrapper around the C function gtk_tool_item_group_get_label.
func (recv *ToolItemGroup) GetLabel() string {
	retC := C.gtk_tool_item_group_get_label((*C.GtkToolItemGroup)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetLabelWidget is a wrapper around the C function gtk_tool_item_group_get_label_widget.
func (recv *ToolItemGroup) GetLabelWidget() *Widget {
	retC := C.gtk_tool_item_group_get_label_widget((*C.GtkToolItemGroup)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetNItems is a wrapper around the C function gtk_tool_item_group_get_n_items.
func (recv *ToolItemGroup) GetNItems() uint32 {
	retC := C.gtk_tool_item_group_get_n_items((*C.GtkToolItemGroup)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// GetNthItem is a wrapper around the C function gtk_tool_item_group_get_nth_item.
func (recv *ToolItemGroup) GetNthItem(index uint32) *ToolItem {
	c_index := (C.guint)(index)

	retC := C.gtk_tool_item_group_get_nth_item((*C.GtkToolItemGroup)(recv.native), c_index)
	retGo := ToolItemNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Insert is a wrapper around the C function gtk_tool_item_group_insert.
func (recv *ToolItemGroup) Insert(item *ToolItem, position int32) {
	c_item := (*C.GtkToolItem)(item.ToC())

	c_position := (C.gint)(position)

	C.gtk_tool_item_group_insert((*C.GtkToolItemGroup)(recv.native), c_item, c_position)

	return
}

// SetCollapsed is a wrapper around the C function gtk_tool_item_group_set_collapsed.
func (recv *ToolItemGroup) SetCollapsed(collapsed bool) {
	c_collapsed :=
		boolToGboolean(collapsed)

	C.gtk_tool_item_group_set_collapsed((*C.GtkToolItemGroup)(recv.native), c_collapsed)

	return
}

// SetEllipsize is a wrapper around the C function gtk_tool_item_group_set_ellipsize.
func (recv *ToolItemGroup) SetEllipsize(ellipsize pango.EllipsizeMode) {
	c_ellipsize := (C.PangoEllipsizeMode)(ellipsize)

	C.gtk_tool_item_group_set_ellipsize((*C.GtkToolItemGroup)(recv.native), c_ellipsize)

	return
}

// SetHeaderRelief is a wrapper around the C function gtk_tool_item_group_set_header_relief.
func (recv *ToolItemGroup) SetHeaderRelief(style ReliefStyle) {
	c_style := (C.GtkReliefStyle)(style)

	C.gtk_tool_item_group_set_header_relief((*C.GtkToolItemGroup)(recv.native), c_style)

	return
}

// SetItemPosition is a wrapper around the C function gtk_tool_item_group_set_item_position.
func (recv *ToolItemGroup) SetItemPosition(item *ToolItem, position int32) {
	c_item := (*C.GtkToolItem)(item.ToC())

	c_position := (C.gint)(position)

	C.gtk_tool_item_group_set_item_position((*C.GtkToolItemGroup)(recv.native), c_item, c_position)

	return
}

// SetLabel is a wrapper around the C function gtk_tool_item_group_set_label.
func (recv *ToolItemGroup) SetLabel(label string) {
	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	C.gtk_tool_item_group_set_label((*C.GtkToolItemGroup)(recv.native), c_label)

	return
}

// SetLabelWidget is a wrapper around the C function gtk_tool_item_group_set_label_widget.
func (recv *ToolItemGroup) SetLabelWidget(labelWidget *Widget) {
	c_label_widget := (*C.GtkWidget)(labelWidget.ToC())

	C.gtk_tool_item_group_set_label_widget((*C.GtkToolItemGroup)(recv.native), c_label_widget)

	return
}

func (recv *ToolItemGroup) Container() *Container {}

// ToolPaletteNew is a wrapper around the C function gtk_tool_palette_new.
func ToolPaletteNew() *ToolPalette {
	retC := C.gtk_tool_palette_new()
	retGo := ToolPaletteNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AddDragDest is a wrapper around the C function gtk_tool_palette_add_drag_dest.
func (recv *ToolPalette) AddDragDest(widget *Widget, flags DestDefaults, targets ToolPaletteDragTargets, actions gdk.DragAction) {
	c_widget := (*C.GtkWidget)(widget.ToC())

	c_flags := (C.GtkDestDefaults)(flags)

	c_targets := (C.GtkToolPaletteDragTargets)(targets)

	c_actions := (C.GdkDragAction)(actions)

	C.gtk_tool_palette_add_drag_dest((*C.GtkToolPalette)(recv.native), c_widget, c_flags, c_targets, c_actions)

	return
}

// GetDragItem is a wrapper around the C function gtk_tool_palette_get_drag_item.
func (recv *ToolPalette) GetDragItem(selection *SelectionData) *Widget {
	c_selection := (*C.GtkSelectionData)(selection.ToC())

	retC := C.gtk_tool_palette_get_drag_item((*C.GtkToolPalette)(recv.native), c_selection)
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetDropGroup is a wrapper around the C function gtk_tool_palette_get_drop_group.
func (recv *ToolPalette) GetDropGroup(x int32, y int32) *ToolItemGroup {
	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	retC := C.gtk_tool_palette_get_drop_group((*C.GtkToolPalette)(recv.native), c_x, c_y)
	retGo := ToolItemGroupNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetDropItem is a wrapper around the C function gtk_tool_palette_get_drop_item.
func (recv *ToolPalette) GetDropItem(x int32, y int32) *ToolItem {
	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	retC := C.gtk_tool_palette_get_drop_item((*C.GtkToolPalette)(recv.native), c_x, c_y)
	retGo := ToolItemNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetExclusive is a wrapper around the C function gtk_tool_palette_get_exclusive.
func (recv *ToolPalette) GetExclusive(group *ToolItemGroup) bool {
	c_group := (*C.GtkToolItemGroup)(group.ToC())

	retC := C.gtk_tool_palette_get_exclusive((*C.GtkToolPalette)(recv.native), c_group)
	retGo := retC == C.TRUE

	return retGo
}

// GetExpand is a wrapper around the C function gtk_tool_palette_get_expand.
func (recv *ToolPalette) GetExpand(group *ToolItemGroup) bool {
	c_group := (*C.GtkToolItemGroup)(group.ToC())

	retC := C.gtk_tool_palette_get_expand((*C.GtkToolPalette)(recv.native), c_group)
	retGo := retC == C.TRUE

	return retGo
}

// GetGroupPosition is a wrapper around the C function gtk_tool_palette_get_group_position.
func (recv *ToolPalette) GetGroupPosition(group *ToolItemGroup) int32 {
	c_group := (*C.GtkToolItemGroup)(group.ToC())

	retC := C.gtk_tool_palette_get_group_position((*C.GtkToolPalette)(recv.native), c_group)
	retGo := (int32)(retC)

	return retGo
}

// GetHadjustment is a wrapper around the C function gtk_tool_palette_get_hadjustment.
func (recv *ToolPalette) GetHadjustment() *Adjustment {
	retC := C.gtk_tool_palette_get_hadjustment((*C.GtkToolPalette)(recv.native))
	retGo := AdjustmentNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_tool_palette_get_icon_size : no return generator

// GetStyle is a wrapper around the C function gtk_tool_palette_get_style.
func (recv *ToolPalette) GetStyle() ToolbarStyle {
	retC := C.gtk_tool_palette_get_style((*C.GtkToolPalette)(recv.native))
	retGo := (ToolbarStyle)(retC)

	return retGo
}

// GetVadjustment is a wrapper around the C function gtk_tool_palette_get_vadjustment.
func (recv *ToolPalette) GetVadjustment() *Adjustment {
	retC := C.gtk_tool_palette_get_vadjustment((*C.GtkToolPalette)(recv.native))
	retGo := AdjustmentNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SetDragSource is a wrapper around the C function gtk_tool_palette_set_drag_source.
func (recv *ToolPalette) SetDragSource(targets ToolPaletteDragTargets) {
	c_targets := (C.GtkToolPaletteDragTargets)(targets)

	C.gtk_tool_palette_set_drag_source((*C.GtkToolPalette)(recv.native), c_targets)

	return
}

// SetExclusive is a wrapper around the C function gtk_tool_palette_set_exclusive.
func (recv *ToolPalette) SetExclusive(group *ToolItemGroup, exclusive bool) {
	c_group := (*C.GtkToolItemGroup)(group.ToC())

	c_exclusive :=
		boolToGboolean(exclusive)

	C.gtk_tool_palette_set_exclusive((*C.GtkToolPalette)(recv.native), c_group, c_exclusive)

	return
}

// SetExpand is a wrapper around the C function gtk_tool_palette_set_expand.
func (recv *ToolPalette) SetExpand(group *ToolItemGroup, expand bool) {
	c_group := (*C.GtkToolItemGroup)(group.ToC())

	c_expand :=
		boolToGboolean(expand)

	C.gtk_tool_palette_set_expand((*C.GtkToolPalette)(recv.native), c_group, c_expand)

	return
}

// SetGroupPosition is a wrapper around the C function gtk_tool_palette_set_group_position.
func (recv *ToolPalette) SetGroupPosition(group *ToolItemGroup, position int32) {
	c_group := (*C.GtkToolItemGroup)(group.ToC())

	c_position := (C.gint)(position)

	C.gtk_tool_palette_set_group_position((*C.GtkToolPalette)(recv.native), c_group, c_position)

	return
}

// Unsupported : gtk_tool_palette_set_icon_size : unsupported parameter icon_size : no type generator for gint, GtkIconSize

// SetStyle is a wrapper around the C function gtk_tool_palette_set_style.
func (recv *ToolPalette) SetStyle(style ToolbarStyle) {
	c_style := (C.GtkToolbarStyle)(style)

	C.gtk_tool_palette_set_style((*C.GtkToolPalette)(recv.native), c_style)

	return
}

// UnsetIconSize is a wrapper around the C function gtk_tool_palette_unset_icon_size.
func (recv *ToolPalette) UnsetIconSize() {
	C.gtk_tool_palette_unset_icon_size((*C.GtkToolPalette)(recv.native))

	return
}

// UnsetStyle is a wrapper around the C function gtk_tool_palette_unset_style.
func (recv *ToolPalette) UnsetStyle() {
	C.gtk_tool_palette_unset_style((*C.GtkToolPalette)(recv.native))

	return
}

func (recv *ToolPalette) Container() *Container {}

func (recv *Toolbar) Container() *Container {}

// Unsupported : gtk_tooltip_set_icon_from_gicon : unsupported parameter gicon : no type generator for Gio.Icon, GIcon*

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

// GetBinWindow is a wrapper around the C function gtk_viewport_get_bin_window.
func (recv *Viewport) GetBinWindow() *gdk.Window {
	retC := C.gtk_viewport_get_bin_window((*C.GtkViewport)(recv.native))
	retGo := gdk.WindowNewFromC(unsafe.Pointer(retC))

	return retGo
}

func (recv *Viewport) Bin() *Bin {}

func (recv *VolumeButton) ScaleButton() *ScaleButton {}

// Unsupported : gtk_widget_new : unsupported parameter type : no type generator for GType, GType

// GetMapped is a wrapper around the C function gtk_widget_get_mapped.
func (recv *Widget) GetMapped() bool {
	retC := C.gtk_widget_get_mapped((*C.GtkWidget)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetRealized is a wrapper around the C function gtk_widget_get_realized.
func (recv *Widget) GetRealized() bool {
	retC := C.gtk_widget_get_realized((*C.GtkWidget)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetRequisition is a wrapper around the C function gtk_widget_get_requisition.
func (recv *Widget) GetRequisition() *Requisition {
	var c_requisition C.GtkRequisition

	C.gtk_widget_get_requisition((*C.GtkWidget)(recv.native), &c_requisition)

	requisition := RequisitionNewFromC(unsafe.Pointer(&c_requisition))

	return requisition
}

// HasRcStyle is a wrapper around the C function gtk_widget_has_rc_style.
func (recv *Widget) HasRcStyle() bool {
	retC := C.gtk_widget_has_rc_style((*C.GtkWidget)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_widget_send_focus_change : unsupported parameter event : no type generator for Gdk.Event, GdkEvent*

// SetMapped is a wrapper around the C function gtk_widget_set_mapped.
func (recv *Widget) SetMapped(mapped bool) {
	c_mapped :=
		boolToGboolean(mapped)

	C.gtk_widget_set_mapped((*C.GtkWidget)(recv.native), c_mapped)

	return
}

// SetRealized is a wrapper around the C function gtk_widget_set_realized.
func (recv *Widget) SetRealized(realized bool) {
	c_realized :=
		boolToGboolean(realized)

	C.gtk_widget_set_realized((*C.GtkWidget)(recv.native), c_realized)

	return
}

// StyleAttach is a wrapper around the C function gtk_widget_style_attach.
func (recv *Widget) StyleAttach() {
	C.gtk_widget_style_attach((*C.GtkWidget)(recv.native))

	return
}

func (recv *Widget) InitiallyUnowned() *gobject.InitiallyUnowned {}

func (recv *WidgetAccessible) Accessible() *Accessible {}

// GetMnemonicsVisible is a wrapper around the C function gtk_window_get_mnemonics_visible.
func (recv *Window) GetMnemonicsVisible() bool {
	retC := C.gtk_window_get_mnemonics_visible((*C.GtkWindow)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetWindowType is a wrapper around the C function gtk_window_get_window_type.
func (recv *Window) GetWindowType() WindowType {
	retC := C.gtk_window_get_window_type((*C.GtkWindow)(recv.native))
	retGo := (WindowType)(retC)

	return retGo
}

// SetMnemonicsVisible is a wrapper around the C function gtk_window_set_mnemonics_visible.
func (recv *Window) SetMnemonicsVisible(setting bool) {
	c_setting :=
		boolToGboolean(setting)

	C.gtk_window_set_mnemonics_visible((*C.GtkWindow)(recv.native), c_setting)

	return
}

func (recv *Window) Bin() *Bin {}

func (recv *WindowAccessible) ContainerAccessible() *ContainerAccessible {}

func (recv *WindowGroup) Object() *gobject.Object {}
