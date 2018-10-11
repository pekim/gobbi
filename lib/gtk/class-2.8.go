// This is a generated file - DO NOT EDIT
// +build gtk_2.8 gtk_2.10 gtk_2.12 gtk_2.14 gtk_2.16 gtk_2.18 gtk_2.20 gtk_2.22 gtk_2.24 gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	atk "github.com/pekim/gobbi/lib/atk"
	cairo "github.com/pekim/gobbi/lib/cairo"
	gio "github.com/pekim/gobbi/lib/gio"
	gobject "github.com/pekim/gobbi/lib/gobject"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

// GetWrapLicense is a wrapper around the C function gtk_about_dialog_get_wrap_license.
func (recv *AboutDialog) GetWrapLicense() bool {
	retC := C.gtk_about_dialog_get_wrap_license((*C.GtkAboutDialog)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetWrapLicense is a wrapper around the C function gtk_about_dialog_set_wrap_license.
func (recv *AboutDialog) SetWrapLicense(wrapLicense bool) {
	c_wrap_license :=
		boolToGboolean(wrapLicense)

	C.gtk_about_dialog_set_wrap_license((*C.GtkAboutDialog)(recv.native), c_wrap_license)

	return
}

func (recv *AboutDialog) Dialog() *Dialog {}

func (recv *AccelGroup) Object() *gobject.Object {}

func (recv *AccelLabel) Label() *Label {}

func (recv *AccelMap) Object() *gobject.Object {}

func (recv *Accessible) Object() *atk.Object {}

// GetAccelClosure is a wrapper around the C function gtk_action_get_accel_closure.
func (recv *Action) GetAccelClosure() *gobject.Closure {
	retC := C.gtk_action_get_accel_closure((*C.GtkAction)(recv.native))
	retGo := gobject.ClosureNewFromC(unsafe.Pointer(retC))

	return retGo
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

// GetResponseForWidget is a wrapper around the C function gtk_dialog_get_response_for_widget.
func (recv *Dialog) GetResponseForWidget(widget *Widget) int32 {
	c_widget := (*C.GtkWidget)(widget.ToC())

	retC := C.gtk_dialog_get_response_for_widget((*C.GtkDialog)(recv.native), c_widget)
	retGo := (int32)(retC)

	return retGo
}

func (recv *Dialog) Window() *Window {}

func (recv *DrawingArea) Widget() *Widget {}

func (recv *Entry) Widget() *Widget {}

func (recv *EntryAccessible) WidgetAccessible() *WidgetAccessible {}

func (recv *EntryBuffer) Object() *gobject.Object {}

// GetPopupSetWidth is a wrapper around the C function gtk_entry_completion_get_popup_set_width.
func (recv *EntryCompletion) GetPopupSetWidth() bool {
	retC := C.gtk_entry_completion_get_popup_set_width((*C.GtkEntryCompletion)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetPopupSingleMatch is a wrapper around the C function gtk_entry_completion_get_popup_single_match.
func (recv *EntryCompletion) GetPopupSingleMatch() bool {
	retC := C.gtk_entry_completion_get_popup_single_match((*C.GtkEntryCompletion)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetPopupSetWidth is a wrapper around the C function gtk_entry_completion_set_popup_set_width.
func (recv *EntryCompletion) SetPopupSetWidth(popupSetWidth bool) {
	c_popup_set_width :=
		boolToGboolean(popupSetWidth)

	C.gtk_entry_completion_set_popup_set_width((*C.GtkEntryCompletion)(recv.native), c_popup_set_width)

	return
}

// SetPopupSingleMatch is a wrapper around the C function gtk_entry_completion_set_popup_single_match.
func (recv *EntryCompletion) SetPopupSingleMatch(popupSingleMatch bool) {
	c_popup_single_match :=
		boolToGboolean(popupSingleMatch)

	C.gtk_entry_completion_set_popup_single_match((*C.GtkEntryCompletion)(recv.native), c_popup_single_match)

	return
}

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

// CreateDragIcon is a wrapper around the C function gtk_icon_view_create_drag_icon.
func (recv *IconView) CreateDragIcon(path *TreePath) *cairo.Surface {
	c_path := (*C.GtkTreePath)(path.ToC())

	retC := C.gtk_icon_view_create_drag_icon((*C.GtkIconView)(recv.native), c_path)
	retGo := cairo.SurfaceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_icon_view_enable_model_drag_dest : unsupported parameter targets : no param type

// Unsupported : gtk_icon_view_enable_model_drag_source : unsupported parameter targets : no param type

// Unsupported : gtk_icon_view_get_cursor : unsupported parameter path : record with indirection level of 2

// Unsupported : gtk_icon_view_get_dest_item_at_pos : unsupported parameter path : record with indirection level of 2

// Unsupported : gtk_icon_view_get_drag_dest_item : unsupported parameter path : record with indirection level of 2

// Unsupported : gtk_icon_view_get_item_at_pos : unsupported parameter path : record with indirection level of 2

// GetReorderable is a wrapper around the C function gtk_icon_view_get_reorderable.
func (recv *IconView) GetReorderable() bool {
	retC := C.gtk_icon_view_get_reorderable((*C.GtkIconView)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_icon_view_get_visible_range : unsupported parameter start_path : record with indirection level of 2

// ScrollToPath is a wrapper around the C function gtk_icon_view_scroll_to_path.
func (recv *IconView) ScrollToPath(path *TreePath, useAlign bool, rowAlign float32, colAlign float32) {
	c_path := (*C.GtkTreePath)(path.ToC())

	c_use_align :=
		boolToGboolean(useAlign)

	c_row_align := (C.gfloat)(rowAlign)

	c_col_align := (C.gfloat)(colAlign)

	C.gtk_icon_view_scroll_to_path((*C.GtkIconView)(recv.native), c_path, c_use_align, c_row_align, c_col_align)

	return
}

// SetCursor is a wrapper around the C function gtk_icon_view_set_cursor.
func (recv *IconView) SetCursor(path *TreePath, cell *CellRenderer, startEditing bool) {
	c_path := (*C.GtkTreePath)(path.ToC())

	c_cell := (*C.GtkCellRenderer)(cell.ToC())

	c_start_editing :=
		boolToGboolean(startEditing)

	C.gtk_icon_view_set_cursor((*C.GtkIconView)(recv.native), c_path, c_cell, c_start_editing)

	return
}

// SetDragDestItem is a wrapper around the C function gtk_icon_view_set_drag_dest_item.
func (recv *IconView) SetDragDestItem(path *TreePath, pos IconViewDropPosition) {
	c_path := (*C.GtkTreePath)(path.ToC())

	c_pos := (C.GtkIconViewDropPosition)(pos)

	C.gtk_icon_view_set_drag_dest_item((*C.GtkIconView)(recv.native), c_path, c_pos)

	return
}

// SetReorderable is a wrapper around the C function gtk_icon_view_set_reorderable.
func (recv *IconView) SetReorderable(reorderable bool) {
	c_reorderable :=
		boolToGboolean(reorderable)

	C.gtk_icon_view_set_reorderable((*C.GtkIconView)(recv.native), c_reorderable)

	return
}

// UnsetModelDragDest is a wrapper around the C function gtk_icon_view_unset_model_drag_dest.
func (recv *IconView) UnsetModelDragDest() {
	C.gtk_icon_view_unset_model_drag_dest((*C.GtkIconView)(recv.native))

	return
}

// UnsetModelDragSource is a wrapper around the C function gtk_icon_view_unset_model_drag_source.
func (recv *IconView) UnsetModelDragSource() {
	C.gtk_icon_view_unset_model_drag_source((*C.GtkIconView)(recv.native))

	return
}

func (recv *IconView) Container() *Container {}

func (recv *IconViewAccessible) ContainerAccessible() *ContainerAccessible {}

// Unsupported : gtk_image_new_from_gicon : unsupported parameter icon : no type generator for Gio.Icon, GIcon*

// Unsupported : gtk_image_new_from_icon_name : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_image_new_from_icon_set : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_image_new_from_stock : unsupported parameter size : no type generator for gint, GtkIconSize

// Clear is a wrapper around the C function gtk_image_clear.
func (recv *Image) Clear() {
	C.gtk_image_clear((*C.GtkImage)(recv.native))

	return
}

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

// GetChildPackDirection is a wrapper around the C function gtk_menu_bar_get_child_pack_direction.
func (recv *MenuBar) GetChildPackDirection() PackDirection {
	retC := C.gtk_menu_bar_get_child_pack_direction((*C.GtkMenuBar)(recv.native))
	retGo := (PackDirection)(retC)

	return retGo
}

// GetPackDirection is a wrapper around the C function gtk_menu_bar_get_pack_direction.
func (recv *MenuBar) GetPackDirection() PackDirection {
	retC := C.gtk_menu_bar_get_pack_direction((*C.GtkMenuBar)(recv.native))
	retGo := (PackDirection)(retC)

	return retGo
}

// SetChildPackDirection is a wrapper around the C function gtk_menu_bar_set_child_pack_direction.
func (recv *MenuBar) SetChildPackDirection(childPackDir PackDirection) {
	c_child_pack_dir := (C.GtkPackDirection)(childPackDir)

	C.gtk_menu_bar_set_child_pack_direction((*C.GtkMenuBar)(recv.native), c_child_pack_dir)

	return
}

// SetPackDirection is a wrapper around the C function gtk_menu_bar_set_pack_direction.
func (recv *MenuBar) SetPackDirection(packDir PackDirection) {
	c_pack_dir := (C.GtkPackDirection)(packDir)

	C.gtk_menu_bar_set_pack_direction((*C.GtkMenuBar)(recv.native), c_pack_dir)

	return
}

func (recv *MenuBar) MenuShell() *MenuShell {}

func (recv *MenuButton) ToggleButton() *ToggleButton {}

func (recv *MenuButtonAccessible) ToggleButtonAccessible() *ToggleButtonAccessible {}

func (recv *MenuItem) Bin() *Bin {}

func (recv *MenuItemAccessible) ContainerAccessible() *ContainerAccessible {}

// GetTakeFocus is a wrapper around the C function gtk_menu_shell_get_take_focus.
func (recv *MenuShell) GetTakeFocus() bool {
	retC := C.gtk_menu_shell_get_take_focus((*C.GtkMenuShell)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetTakeFocus is a wrapper around the C function gtk_menu_shell_set_take_focus.
func (recv *MenuShell) SetTakeFocus(takeFocus bool) {
	c_take_focus :=
		boolToGboolean(takeFocus)

	C.gtk_menu_shell_set_take_focus((*C.GtkMenuShell)(recv.native), c_take_focus)

	return
}

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

func (recv *Revealer) Bin() *Bin {}

func (recv *Scale) Range() *Range {}

func (recv *ScaleAccessible) RangeAccessible() *RangeAccessible {}

// Unsupported : gtk_scale_button_new : unsupported parameter size : no type generator for gint, GtkIconSize

func (recv *ScaleButton) Button() *Button {}

func (recv *ScaleButtonAccessible) ButtonAccessible() *ButtonAccessible {}

func (recv *Scrollbar) Range() *Range {}

// GetHscrollbar is a wrapper around the C function gtk_scrolled_window_get_hscrollbar.
func (recv *ScrolledWindow) GetHscrollbar() *Widget {
	retC := C.gtk_scrolled_window_get_hscrollbar((*C.GtkScrolledWindow)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetVscrollbar is a wrapper around the C function gtk_scrolled_window_get_vscrollbar.
func (recv *ScrolledWindow) GetVscrollbar() *Widget {
	retC := C.gtk_scrolled_window_get_vscrollbar((*C.GtkScrolledWindow)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

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

// GetIgnoreHidden is a wrapper around the C function gtk_size_group_get_ignore_hidden.
func (recv *SizeGroup) GetIgnoreHidden() bool {
	retC := C.gtk_size_group_get_ignore_hidden((*C.GtkSizeGroup)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetIgnoreHidden is a wrapper around the C function gtk_size_group_set_ignore_hidden.
func (recv *SizeGroup) SetIgnoreHidden(ignoreHidden bool) {
	c_ignore_hidden :=
		boolToGboolean(ignoreHidden)

	C.gtk_size_group_set_ignore_hidden((*C.GtkSizeGroup)(recv.native), c_ignore_hidden)

	return
}

func (recv *SizeGroup) Object() *gobject.Object {}

func (recv *SpinButton) Entry() *Entry {}

func (recv *SpinButtonAccessible) EntryAccessible() *EntryAccessible {}

func (recv *Spinner) Widget() *Widget {}

func (recv *SpinnerAccessible) WidgetAccessible() *WidgetAccessible {}

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

// GetIconName is a wrapper around the C function gtk_tool_button_get_icon_name.
func (recv *ToolButton) GetIconName() string {
	retC := C.gtk_tool_button_get_icon_name((*C.GtkToolButton)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// SetIconName is a wrapper around the C function gtk_tool_button_set_icon_name.
func (recv *ToolButton) SetIconName(iconName string) {
	c_icon_name := C.CString(iconName)
	defer C.free(unsafe.Pointer(c_icon_name))

	C.gtk_tool_button_set_icon_name((*C.GtkToolButton)(recv.native), c_icon_name)

	return
}

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

// Unsupported : gtk_tree_view_get_visible_range : unsupported parameter start_path : record with indirection level of 2

func (recv *TreeView) Container() *Container {}

func (recv *TreeViewAccessible) ContainerAccessible() *ContainerAccessible {}

// Unsupported : gtk_tree_view_column_new_with_attributes : unsupported parameter ... : varargs

// QueueResize is a wrapper around the C function gtk_tree_view_column_queue_resize.
func (recv *TreeViewColumn) QueueResize() {
	C.gtk_tree_view_column_queue_resize((*C.GtkTreeViewColumn)(recv.native))

	return
}

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

// DragSourceSetIconName is a wrapper around the C function gtk_drag_source_set_icon_name.
func (recv *Widget) DragSourceSetIconName(iconName string) {
	c_icon_name := C.CString(iconName)
	defer C.free(unsafe.Pointer(c_icon_name))

	C.gtk_drag_source_set_icon_name((*C.GtkWidget)(recv.native), c_icon_name)

	return
}

func (recv *Widget) InitiallyUnowned() *gobject.InitiallyUnowned {}

func (recv *WidgetAccessible) Accessible() *Accessible {}

// GetUrgencyHint is a wrapper around the C function gtk_window_get_urgency_hint.
func (recv *Window) GetUrgencyHint() bool {
	retC := C.gtk_window_get_urgency_hint((*C.GtkWindow)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// PresentWithTime is a wrapper around the C function gtk_window_present_with_time.
func (recv *Window) PresentWithTime(timestamp uint32) {
	c_timestamp := (C.guint32)(timestamp)

	C.gtk_window_present_with_time((*C.GtkWindow)(recv.native), c_timestamp)

	return
}

// SetUrgencyHint is a wrapper around the C function gtk_window_set_urgency_hint.
func (recv *Window) SetUrgencyHint(setting bool) {
	c_setting :=
		boolToGboolean(setting)

	C.gtk_window_set_urgency_hint((*C.GtkWindow)(recv.native), c_setting)

	return
}

func (recv *Window) Bin() *Bin {}

func (recv *WindowAccessible) ContainerAccessible() *ContainerAccessible {}

func (recv *WindowGroup) Object() *gobject.Object {}
