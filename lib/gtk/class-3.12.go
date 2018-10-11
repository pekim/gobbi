// This is a generated file - DO NOT EDIT
// +build gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	atk "github.com/pekim/gobbi/lib/atk"
	gio "github.com/pekim/gobbi/lib/gio"
	glib "github.com/pekim/gobbi/lib/glib"
	gobject "github.com/pekim/gobbi/lib/gobject"
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

// Unsupported : gtk_accel_label_get_accel : unsupported parameter accelerator_mods : GdkModifierType* with indirection level of 1

func (recv *AccelLabel) Label() *Label {}

func (recv *AccelMap) Object() *gobject.Object {}

func (recv *Accessible) Object() *atk.Object {}

func (recv *Action) Object() *gobject.Object {}

// ActionBarNew is a wrapper around the C function gtk_action_bar_new.
func ActionBarNew() *ActionBar {
	retC := C.gtk_action_bar_new()
	retGo := ActionBarNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetCenterWidget is a wrapper around the C function gtk_action_bar_get_center_widget.
func (recv *ActionBar) GetCenterWidget() *Widget {
	retC := C.gtk_action_bar_get_center_widget((*C.GtkActionBar)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// PackEnd is a wrapper around the C function gtk_action_bar_pack_end.
func (recv *ActionBar) PackEnd(child *Widget) {
	c_child := (*C.GtkWidget)(child.ToC())

	C.gtk_action_bar_pack_end((*C.GtkActionBar)(recv.native), c_child)

	return
}

// PackStart is a wrapper around the C function gtk_action_bar_pack_start.
func (recv *ActionBar) PackStart(child *Widget) {
	c_child := (*C.GtkWidget)(child.ToC())

	C.gtk_action_bar_pack_start((*C.GtkActionBar)(recv.native), c_child)

	return
}

// SetCenterWidget is a wrapper around the C function gtk_action_bar_set_center_widget.
func (recv *ActionBar) SetCenterWidget(centerWidget *Widget) {
	c_center_widget := (*C.GtkWidget)(centerWidget.ToC())

	C.gtk_action_bar_set_center_widget((*C.GtkActionBar)(recv.native), c_center_widget)

	return
}

func (recv *ActionBar) Bin() *Bin {}

func (recv *ActionGroup) Object() *gobject.Object {}

func (recv *Adjustment) InitiallyUnowned() *gobject.InitiallyUnowned {}

func (recv *Alignment) Bin() *Bin {}

func (recv *AppChooserButton) ComboBox() *ComboBox {}

// Unsupported : gtk_app_chooser_dialog_new : unsupported parameter file : no type generator for Gio.File, GFile*

func (recv *AppChooserDialog) Dialog() *Dialog {}

func (recv *AppChooserWidget) Box() *Box {}

// Unsupported : gtk_application_get_accels_for_action : no return type

// Unsupported : gtk_application_list_action_descriptions : no return type

// Unsupported : gtk_application_set_accels_for_action : unsupported parameter accels : no param type

func (recv *Application) Application() *gio.Application {}

func (recv *ApplicationWindow) Window() *Window {}

func (recv *Arrow) Misc() *Misc {}

func (recv *ArrowAccessible) WidgetAccessible() *WidgetAccessible {}

func (recv *AspectFrame) Frame() *Frame {}

func (recv *Assistant) Window() *Window {}

func (recv *Bin) Container() *Container {}

func (recv *BooleanCellAccessible) RendererCellAccessible() *RendererCellAccessible {}

// GetCenterWidget is a wrapper around the C function gtk_box_get_center_widget.
func (recv *Box) GetCenterWidget() *Widget {
	retC := C.gtk_box_get_center_widget((*C.GtkBox)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SetCenterWidget is a wrapper around the C function gtk_box_set_center_widget.
func (recv *Box) SetCenterWidget(widget *Widget) {
	c_widget := (*C.GtkWidget)(widget.ToC())

	C.gtk_box_set_center_widget((*C.GtkBox)(recv.native), c_widget)

	return
}

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

// GetHeaderBar is a wrapper around the C function gtk_dialog_get_header_bar.
func (recv *Dialog) GetHeaderBar() *Widget {
	retC := C.gtk_dialog_get_header_bar((*C.GtkDialog)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

func (recv *Dialog) Window() *Window {}

func (recv *DrawingArea) Widget() *Widget {}

// GetMaxWidthChars is a wrapper around the C function gtk_entry_get_max_width_chars.
func (recv *Entry) GetMaxWidthChars() int32 {
	retC := C.gtk_entry_get_max_width_chars((*C.GtkEntry)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// SetMaxWidthChars is a wrapper around the C function gtk_entry_set_max_width_chars.
func (recv *Entry) SetMaxWidthChars(nChars int32) {
	c_n_chars := (C.gint)(nChars)

	C.gtk_entry_set_max_width_chars((*C.GtkEntry)(recv.native), c_n_chars)

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

// FlowBoxNew is a wrapper around the C function gtk_flow_box_new.
func FlowBoxNew() *FlowBox {
	retC := C.gtk_flow_box_new()
	retGo := FlowBoxNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetActivateOnSingleClick is a wrapper around the C function gtk_flow_box_get_activate_on_single_click.
func (recv *FlowBox) GetActivateOnSingleClick() bool {
	retC := C.gtk_flow_box_get_activate_on_single_click((*C.GtkFlowBox)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetChildAtIndex is a wrapper around the C function gtk_flow_box_get_child_at_index.
func (recv *FlowBox) GetChildAtIndex(idx int32) *FlowBoxChild {
	c_idx := (C.gint)(idx)

	retC := C.gtk_flow_box_get_child_at_index((*C.GtkFlowBox)(recv.native), c_idx)
	retGo := FlowBoxChildNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetColumnSpacing is a wrapper around the C function gtk_flow_box_get_column_spacing.
func (recv *FlowBox) GetColumnSpacing() uint32 {
	retC := C.gtk_flow_box_get_column_spacing((*C.GtkFlowBox)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// GetHomogeneous is a wrapper around the C function gtk_flow_box_get_homogeneous.
func (recv *FlowBox) GetHomogeneous() bool {
	retC := C.gtk_flow_box_get_homogeneous((*C.GtkFlowBox)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetMaxChildrenPerLine is a wrapper around the C function gtk_flow_box_get_max_children_per_line.
func (recv *FlowBox) GetMaxChildrenPerLine() uint32 {
	retC := C.gtk_flow_box_get_max_children_per_line((*C.GtkFlowBox)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// GetMinChildrenPerLine is a wrapper around the C function gtk_flow_box_get_min_children_per_line.
func (recv *FlowBox) GetMinChildrenPerLine() uint32 {
	retC := C.gtk_flow_box_get_min_children_per_line((*C.GtkFlowBox)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// GetRowSpacing is a wrapper around the C function gtk_flow_box_get_row_spacing.
func (recv *FlowBox) GetRowSpacing() uint32 {
	retC := C.gtk_flow_box_get_row_spacing((*C.GtkFlowBox)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// GetSelectedChildren is a wrapper around the C function gtk_flow_box_get_selected_children.
func (recv *FlowBox) GetSelectedChildren() *glib.List {
	retC := C.gtk_flow_box_get_selected_children((*C.GtkFlowBox)(recv.native))
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetSelectionMode is a wrapper around the C function gtk_flow_box_get_selection_mode.
func (recv *FlowBox) GetSelectionMode() SelectionMode {
	retC := C.gtk_flow_box_get_selection_mode((*C.GtkFlowBox)(recv.native))
	retGo := (SelectionMode)(retC)

	return retGo
}

// Insert is a wrapper around the C function gtk_flow_box_insert.
func (recv *FlowBox) Insert(widget *Widget, position int32) {
	c_widget := (*C.GtkWidget)(widget.ToC())

	c_position := (C.gint)(position)

	C.gtk_flow_box_insert((*C.GtkFlowBox)(recv.native), c_widget, c_position)

	return
}

// InvalidateFilter is a wrapper around the C function gtk_flow_box_invalidate_filter.
func (recv *FlowBox) InvalidateFilter() {
	C.gtk_flow_box_invalidate_filter((*C.GtkFlowBox)(recv.native))

	return
}

// InvalidateSort is a wrapper around the C function gtk_flow_box_invalidate_sort.
func (recv *FlowBox) InvalidateSort() {
	C.gtk_flow_box_invalidate_sort((*C.GtkFlowBox)(recv.native))

	return
}

// SelectAll is a wrapper around the C function gtk_flow_box_select_all.
func (recv *FlowBox) SelectAll() {
	C.gtk_flow_box_select_all((*C.GtkFlowBox)(recv.native))

	return
}

// SelectChild is a wrapper around the C function gtk_flow_box_select_child.
func (recv *FlowBox) SelectChild(child *FlowBoxChild) {
	c_child := (*C.GtkFlowBoxChild)(child.ToC())

	C.gtk_flow_box_select_child((*C.GtkFlowBox)(recv.native), c_child)

	return
}

// Unsupported : gtk_flow_box_selected_foreach : unsupported parameter func : no type generator for FlowBoxForeachFunc, GtkFlowBoxForeachFunc

// SetActivateOnSingleClick is a wrapper around the C function gtk_flow_box_set_activate_on_single_click.
func (recv *FlowBox) SetActivateOnSingleClick(single bool) {
	c_single :=
		boolToGboolean(single)

	C.gtk_flow_box_set_activate_on_single_click((*C.GtkFlowBox)(recv.native), c_single)

	return
}

// SetColumnSpacing is a wrapper around the C function gtk_flow_box_set_column_spacing.
func (recv *FlowBox) SetColumnSpacing(spacing uint32) {
	c_spacing := (C.guint)(spacing)

	C.gtk_flow_box_set_column_spacing((*C.GtkFlowBox)(recv.native), c_spacing)

	return
}

// Unsupported : gtk_flow_box_set_filter_func : unsupported parameter filter_func : no type generator for FlowBoxFilterFunc, GtkFlowBoxFilterFunc

// SetHadjustment is a wrapper around the C function gtk_flow_box_set_hadjustment.
func (recv *FlowBox) SetHadjustment(adjustment *Adjustment) {
	c_adjustment := (*C.GtkAdjustment)(adjustment.ToC())

	C.gtk_flow_box_set_hadjustment((*C.GtkFlowBox)(recv.native), c_adjustment)

	return
}

// SetHomogeneous is a wrapper around the C function gtk_flow_box_set_homogeneous.
func (recv *FlowBox) SetHomogeneous(homogeneous bool) {
	c_homogeneous :=
		boolToGboolean(homogeneous)

	C.gtk_flow_box_set_homogeneous((*C.GtkFlowBox)(recv.native), c_homogeneous)

	return
}

// SetMaxChildrenPerLine is a wrapper around the C function gtk_flow_box_set_max_children_per_line.
func (recv *FlowBox) SetMaxChildrenPerLine(nChildren uint32) {
	c_n_children := (C.guint)(nChildren)

	C.gtk_flow_box_set_max_children_per_line((*C.GtkFlowBox)(recv.native), c_n_children)

	return
}

// SetMinChildrenPerLine is a wrapper around the C function gtk_flow_box_set_min_children_per_line.
func (recv *FlowBox) SetMinChildrenPerLine(nChildren uint32) {
	c_n_children := (C.guint)(nChildren)

	C.gtk_flow_box_set_min_children_per_line((*C.GtkFlowBox)(recv.native), c_n_children)

	return
}

// SetRowSpacing is a wrapper around the C function gtk_flow_box_set_row_spacing.
func (recv *FlowBox) SetRowSpacing(spacing uint32) {
	c_spacing := (C.guint)(spacing)

	C.gtk_flow_box_set_row_spacing((*C.GtkFlowBox)(recv.native), c_spacing)

	return
}

// SetSelectionMode is a wrapper around the C function gtk_flow_box_set_selection_mode.
func (recv *FlowBox) SetSelectionMode(mode SelectionMode) {
	c_mode := (C.GtkSelectionMode)(mode)

	C.gtk_flow_box_set_selection_mode((*C.GtkFlowBox)(recv.native), c_mode)

	return
}

// Unsupported : gtk_flow_box_set_sort_func : unsupported parameter sort_func : no type generator for FlowBoxSortFunc, GtkFlowBoxSortFunc

// SetVadjustment is a wrapper around the C function gtk_flow_box_set_vadjustment.
func (recv *FlowBox) SetVadjustment(adjustment *Adjustment) {
	c_adjustment := (*C.GtkAdjustment)(adjustment.ToC())

	C.gtk_flow_box_set_vadjustment((*C.GtkFlowBox)(recv.native), c_adjustment)

	return
}

// UnselectAll is a wrapper around the C function gtk_flow_box_unselect_all.
func (recv *FlowBox) UnselectAll() {
	C.gtk_flow_box_unselect_all((*C.GtkFlowBox)(recv.native))

	return
}

// UnselectChild is a wrapper around the C function gtk_flow_box_unselect_child.
func (recv *FlowBox) UnselectChild(child *FlowBoxChild) {
	c_child := (*C.GtkFlowBoxChild)(child.ToC())

	C.gtk_flow_box_unselect_child((*C.GtkFlowBox)(recv.native), c_child)

	return
}

func (recv *FlowBox) Container() *Container {}

func (recv *FlowBoxAccessible) ContainerAccessible() *ContainerAccessible {}

// FlowBoxChildNew is a wrapper around the C function gtk_flow_box_child_new.
func FlowBoxChildNew() *FlowBoxChild {
	retC := C.gtk_flow_box_child_new()
	retGo := FlowBoxChildNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Changed is a wrapper around the C function gtk_flow_box_child_changed.
func (recv *FlowBoxChild) Changed() {
	C.gtk_flow_box_child_changed((*C.GtkFlowBoxChild)(recv.native))

	return
}

// GetIndex is a wrapper around the C function gtk_flow_box_child_get_index.
func (recv *FlowBoxChild) GetIndex() int32 {
	retC := C.gtk_flow_box_child_get_index((*C.GtkFlowBoxChild)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// IsSelected is a wrapper around the C function gtk_flow_box_child_is_selected.
func (recv *FlowBoxChild) IsSelected() bool {
	retC := C.gtk_flow_box_child_is_selected((*C.GtkFlowBoxChild)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

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

// GetDecorationLayout is a wrapper around the C function gtk_header_bar_get_decoration_layout.
func (recv *HeaderBar) GetDecorationLayout() string {
	retC := C.gtk_header_bar_get_decoration_layout((*C.GtkHeaderBar)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetHasSubtitle is a wrapper around the C function gtk_header_bar_get_has_subtitle.
func (recv *HeaderBar) GetHasSubtitle() bool {
	retC := C.gtk_header_bar_get_has_subtitle((*C.GtkHeaderBar)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetDecorationLayout is a wrapper around the C function gtk_header_bar_set_decoration_layout.
func (recv *HeaderBar) SetDecorationLayout(layout string) {
	c_layout := C.CString(layout)
	defer C.free(unsafe.Pointer(c_layout))

	C.gtk_header_bar_set_decoration_layout((*C.GtkHeaderBar)(recv.native), c_layout)

	return
}

// SetHasSubtitle is a wrapper around the C function gtk_header_bar_set_has_subtitle.
func (recv *HeaderBar) SetHasSubtitle(setting bool) {
	c_setting :=
		boolToGboolean(setting)

	C.gtk_header_bar_set_has_subtitle((*C.GtkHeaderBar)(recv.native), c_setting)

	return
}

func (recv *HeaderBar) Container() *Container {}

func (recv *IMContext) Object() *gobject.Object {}

func (recv *IMContextSimple) IMContext() *IMContext {}

func (recv *IMMulticontext) IMContext() *IMContext {}

func (recv *IconFactory) Object() *gobject.Object {}

// IsSymbolic is a wrapper around the C function gtk_icon_info_is_symbolic.
func (recv *IconInfo) IsSymbolic() bool {
	retC := C.gtk_icon_info_is_symbolic((*C.GtkIconInfo)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

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

// GetPopover is a wrapper around the C function gtk_menu_button_get_popover.
func (recv *MenuButton) GetPopover() *Popover {
	retC := C.gtk_menu_button_get_popover((*C.GtkMenuButton)(recv.native))
	retGo := PopoverNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetUsePopover is a wrapper around the C function gtk_menu_button_get_use_popover.
func (recv *MenuButton) GetUsePopover() bool {
	retC := C.gtk_menu_button_get_use_popover((*C.GtkMenuButton)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetPopover is a wrapper around the C function gtk_menu_button_set_popover.
func (recv *MenuButton) SetPopover(popover *Widget) {
	c_popover := (*C.GtkWidget)(popover.ToC())

	C.gtk_menu_button_set_popover((*C.GtkMenuButton)(recv.native), c_popover)

	return
}

// SetUsePopover is a wrapper around the C function gtk_menu_button_set_use_popover.
func (recv *MenuButton) SetUsePopover(usePopover bool) {
	c_use_popover :=
		boolToGboolean(usePopover)

	C.gtk_menu_button_set_use_popover((*C.GtkMenuButton)(recv.native), c_use_popover)

	return
}

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

// GetLocalOnly is a wrapper around the C function gtk_places_sidebar_get_local_only.
func (recv *PlacesSidebar) GetLocalOnly() bool {
	retC := C.gtk_places_sidebar_get_local_only((*C.GtkPlacesSidebar)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetLocalOnly is a wrapper around the C function gtk_places_sidebar_set_local_only.
func (recv *PlacesSidebar) SetLocalOnly(localOnly bool) {
	c_local_only :=
		boolToGboolean(localOnly)

	C.gtk_places_sidebar_set_local_only((*C.GtkPlacesSidebar)(recv.native), c_local_only)

	return
}

func (recv *PlacesSidebar) ScrolledWindow() *ScrolledWindow {}

// PopoverNew is a wrapper around the C function gtk_popover_new.
func PopoverNew(relativeTo *Widget) *Popover {
	c_relative_to := (*C.GtkWidget)(relativeTo.ToC())

	retC := C.gtk_popover_new(c_relative_to)
	retGo := PopoverNewFromC(unsafe.Pointer(retC))

	return retGo
}

// PopoverNewFromModel is a wrapper around the C function gtk_popover_new_from_model.
func PopoverNewFromModel(relativeTo *Widget, model *gio.MenuModel) *Popover {
	c_relative_to := (*C.GtkWidget)(relativeTo.ToC())

	c_model := (*C.GMenuModel)(model.ToC())

	retC := C.gtk_popover_new_from_model(c_relative_to, c_model)
	retGo := PopoverNewFromC(unsafe.Pointer(retC))

	return retGo
}

// BindModel is a wrapper around the C function gtk_popover_bind_model.
func (recv *Popover) BindModel(model *gio.MenuModel, actionNamespace string) {
	c_model := (*C.GMenuModel)(model.ToC())

	c_action_namespace := C.CString(actionNamespace)
	defer C.free(unsafe.Pointer(c_action_namespace))

	C.gtk_popover_bind_model((*C.GtkPopover)(recv.native), c_model, c_action_namespace)

	return
}

// GetModal is a wrapper around the C function gtk_popover_get_modal.
func (recv *Popover) GetModal() bool {
	retC := C.gtk_popover_get_modal((*C.GtkPopover)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetRelativeTo is a wrapper around the C function gtk_popover_get_relative_to.
func (recv *Popover) GetRelativeTo() *Widget {
	retC := C.gtk_popover_get_relative_to((*C.GtkPopover)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SetModal is a wrapper around the C function gtk_popover_set_modal.
func (recv *Popover) SetModal(modal bool) {
	c_modal :=
		boolToGboolean(modal)

	C.gtk_popover_set_modal((*C.GtkPopover)(recv.native), c_modal)

	return
}

// Unsupported : gtk_popover_set_pointing_to : unsupported parameter rect : Blacklisted record : GdkRectangle

// SetPosition is a wrapper around the C function gtk_popover_set_position.
func (recv *Popover) SetPosition(position PositionType) {
	c_position := (C.GtkPositionType)(position)

	C.gtk_popover_set_position((*C.GtkPopover)(recv.native), c_position)

	return
}

// SetRelativeTo is a wrapper around the C function gtk_popover_set_relative_to.
func (recv *Popover) SetRelativeTo(relativeTo *Widget) {
	c_relative_to := (*C.GtkWidget)(relativeTo.ToC())

	C.gtk_popover_set_relative_to((*C.GtkPopover)(recv.native), c_relative_to)

	return
}

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

func (recv *Spinner) Widget() *Widget {}

func (recv *SpinnerAccessible) WidgetAccessible() *WidgetAccessible {}

// GetChildByName is a wrapper around the C function gtk_stack_get_child_by_name.
func (recv *Stack) GetChildByName(name string) *Widget {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.gtk_stack_get_child_by_name((*C.GtkStack)(recv.native), c_name)
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetTransitionRunning is a wrapper around the C function gtk_stack_get_transition_running.
func (recv *Stack) GetTransitionRunning() bool {
	retC := C.gtk_stack_get_transition_running((*C.GtkStack)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

func (recv *Stack) Container() *Container {}

func (recv *StackSidebar) Bin() *Bin {}

func (recv *StackSwitcher) Box() *Box {}

// Unsupported : gtk_status_icon_new_from_gicon : unsupported parameter icon : no type generator for Gio.Icon, GIcon*

func (recv *StatusIcon) Object() *gobject.Object {}

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

// GetMarginEnd is a wrapper around the C function gtk_widget_get_margin_end.
func (recv *Widget) GetMarginEnd() int32 {
	retC := C.gtk_widget_get_margin_end((*C.GtkWidget)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetMarginStart is a wrapper around the C function gtk_widget_get_margin_start.
func (recv *Widget) GetMarginStart() int32 {
	retC := C.gtk_widget_get_margin_start((*C.GtkWidget)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// SetMarginEnd is a wrapper around the C function gtk_widget_set_margin_end.
func (recv *Widget) SetMarginEnd(margin int32) {
	c_margin := (C.gint)(margin)

	C.gtk_widget_set_margin_end((*C.GtkWidget)(recv.native), c_margin)

	return
}

// SetMarginStart is a wrapper around the C function gtk_widget_set_margin_start.
func (recv *Widget) SetMarginStart(margin int32) {
	c_margin := (C.gint)(margin)

	C.gtk_widget_set_margin_start((*C.GtkWidget)(recv.native), c_margin)

	return
}

func (recv *Widget) InitiallyUnowned() *gobject.InitiallyUnowned {}

func (recv *WidgetAccessible) Accessible() *Accessible {}

// IsMaximized is a wrapper around the C function gtk_window_is_maximized.
func (recv *Window) IsMaximized() bool {
	retC := C.gtk_window_is_maximized((*C.GtkWindow)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

func (recv *Window) Bin() *Bin {}

func (recv *WindowAccessible) ContainerAccessible() *ContainerAccessible {}

func (recv *WindowGroup) Object() *gobject.Object {}
