// This is a generated file - DO NOT EDIT
// +build gtk_2.2 gtk_2.4 gtk_2.6 gtk_2.8 gtk_2.10 gtk_2.12 gtk_2.14 gtk_2.16 gtk_2.18 gtk_2.20 gtk_2.22 gtk_2.24 gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	atk "github.com/pekim/gobbi/lib/atk"
	gdk "github.com/pekim/gobbi/lib/gdk"
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

// GetDisplay is a wrapper around the C function gtk_clipboard_get_display.
func (recv *Clipboard) GetDisplay() *gdk.Display {
	retC := C.gtk_clipboard_get_display((*C.GtkClipboard)(recv.native))
	retGo := gdk.DisplayNewFromC(unsafe.Pointer(retC))

	return retGo
}

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

// InvisibleNewForScreen is a wrapper around the C function gtk_invisible_new_for_screen.
func InvisibleNewForScreen(screen *gdk.Screen) *Invisible {
	c_screen := (*C.GdkScreen)(screen.ToC())

	retC := C.gtk_invisible_new_for_screen(c_screen)
	retGo := InvisibleNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetScreen is a wrapper around the C function gtk_invisible_get_screen.
func (recv *Invisible) GetScreen() *gdk.Screen {
	retC := C.gtk_invisible_get_screen((*C.GtkInvisible)(recv.native))
	retGo := gdk.ScreenNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SetScreen is a wrapper around the C function gtk_invisible_set_screen.
func (recv *Invisible) SetScreen(screen *gdk.Screen) {
	c_screen := (*C.GdkScreen)(screen.ToC())

	C.gtk_invisible_set_screen((*C.GtkInvisible)(recv.native), c_screen)

	return
}

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

// IterIsValid is a wrapper around the C function gtk_list_store_iter_is_valid.
func (recv *ListStore) IterIsValid(iter *TreeIter) bool {
	c_iter := (*C.GtkTreeIter)(iter.ToC())

	retC := C.gtk_list_store_iter_is_valid((*C.GtkListStore)(recv.native), c_iter)
	retGo := retC == C.TRUE

	return retGo
}

// MoveAfter is a wrapper around the C function gtk_list_store_move_after.
func (recv *ListStore) MoveAfter(iter *TreeIter, position *TreeIter) {
	c_iter := (*C.GtkTreeIter)(iter.ToC())

	c_position := (*C.GtkTreeIter)(position.ToC())

	C.gtk_list_store_move_after((*C.GtkListStore)(recv.native), c_iter, c_position)

	return
}

// MoveBefore is a wrapper around the C function gtk_list_store_move_before.
func (recv *ListStore) MoveBefore(iter *TreeIter, position *TreeIter) {
	c_iter := (*C.GtkTreeIter)(iter.ToC())

	c_position := (*C.GtkTreeIter)(position.ToC())

	C.gtk_list_store_move_before((*C.GtkListStore)(recv.native), c_iter, c_position)

	return
}

// Unsupported : gtk_list_store_reorder : unsupported parameter new_order : no param type

// Swap is a wrapper around the C function gtk_list_store_swap.
func (recv *ListStore) Swap(a *TreeIter, b *TreeIter) {
	c_a := (*C.GtkTreeIter)(a.ToC())

	c_b := (*C.GtkTreeIter)(b.ToC())

	C.gtk_list_store_swap((*C.GtkListStore)(recv.native), c_a, c_b)

	return
}

func (recv *ListStore) Object() *gobject.Object {}

func (recv *LockButton) Button() *Button {}

func (recv *LockButtonAccessible) ButtonAccessible() *ButtonAccessible {}

// SetScreen is a wrapper around the C function gtk_menu_set_screen.
func (recv *Menu) SetScreen(screen *gdk.Screen) {
	c_screen := (*C.GdkScreen)(screen.ToC())

	C.gtk_menu_set_screen((*C.GtkMenu)(recv.native), c_screen)

	return
}

func (recv *Menu) MenuShell() *MenuShell {}

func (recv *MenuAccessible) MenuShellAccessible() *MenuShellAccessible {}

func (recv *MenuBar) MenuShell() *MenuShell {}

func (recv *MenuButton) ToggleButton() *ToggleButton {}

func (recv *MenuButtonAccessible) ToggleButtonAccessible() *ToggleButtonAccessible {}

func (recv *MenuItem) Bin() *Bin {}

func (recv *MenuItemAccessible) ContainerAccessible() *ContainerAccessible {}

// SelectFirst is a wrapper around the C function gtk_menu_shell_select_first.
func (recv *MenuShell) SelectFirst(searchSensitive bool) {
	c_search_sensitive :=
		boolToGboolean(searchSensitive)

	C.gtk_menu_shell_select_first((*C.GtkMenuShell)(recv.native), c_search_sensitive)

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

// GetNPages is a wrapper around the C function gtk_notebook_get_n_pages.
func (recv *Notebook) GetNPages() int32 {
	retC := C.gtk_notebook_get_n_pages((*C.GtkNotebook)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

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

// IterIsValid is a wrapper around the C function gtk_tree_model_sort_iter_is_valid.
func (recv *TreeModelSort) IterIsValid(iter *TreeIter) bool {
	c_iter := (*C.GtkTreeIter)(iter.ToC())

	retC := C.gtk_tree_model_sort_iter_is_valid((*C.GtkTreeModelSort)(recv.native), c_iter)
	retGo := retC == C.TRUE

	return retGo
}

func (recv *TreeModelSort) Object() *gobject.Object {}

// CountSelectedRows is a wrapper around the C function gtk_tree_selection_count_selected_rows.
func (recv *TreeSelection) CountSelectedRows() int32 {
	retC := C.gtk_tree_selection_count_selected_rows((*C.GtkTreeSelection)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : gtk_tree_selection_get_selected_rows : unsupported parameter model : no type generator for TreeModel, GtkTreeModel**

// UnselectRange is a wrapper around the C function gtk_tree_selection_unselect_range.
func (recv *TreeSelection) UnselectRange(startPath *TreePath, endPath *TreePath) {
	c_start_path := (*C.GtkTreePath)(startPath.ToC())

	c_end_path := (*C.GtkTreePath)(endPath.ToC())

	C.gtk_tree_selection_unselect_range((*C.GtkTreeSelection)(recv.native), c_start_path, c_end_path)

	return
}

func (recv *TreeSelection) Object() *gobject.Object {}

// Unsupported : gtk_tree_store_new : unsupported parameter ... : varargs

// Unsupported : gtk_tree_store_newv : unsupported parameter types : no param type

// IterIsValid is a wrapper around the C function gtk_tree_store_iter_is_valid.
func (recv *TreeStore) IterIsValid(iter *TreeIter) bool {
	c_iter := (*C.GtkTreeIter)(iter.ToC())

	retC := C.gtk_tree_store_iter_is_valid((*C.GtkTreeStore)(recv.native), c_iter)
	retGo := retC == C.TRUE

	return retGo
}

// MoveAfter is a wrapper around the C function gtk_tree_store_move_after.
func (recv *TreeStore) MoveAfter(iter *TreeIter, position *TreeIter) {
	c_iter := (*C.GtkTreeIter)(iter.ToC())

	c_position := (*C.GtkTreeIter)(position.ToC())

	C.gtk_tree_store_move_after((*C.GtkTreeStore)(recv.native), c_iter, c_position)

	return
}

// MoveBefore is a wrapper around the C function gtk_tree_store_move_before.
func (recv *TreeStore) MoveBefore(iter *TreeIter, position *TreeIter) {
	c_iter := (*C.GtkTreeIter)(iter.ToC())

	c_position := (*C.GtkTreeIter)(position.ToC())

	C.gtk_tree_store_move_before((*C.GtkTreeStore)(recv.native), c_iter, c_position)

	return
}

// Unsupported : gtk_tree_store_reorder : unsupported parameter new_order : no param type

// Swap is a wrapper around the C function gtk_tree_store_swap.
func (recv *TreeStore) Swap(a *TreeIter, b *TreeIter) {
	c_a := (*C.GtkTreeIter)(a.ToC())

	c_b := (*C.GtkTreeIter)(b.ToC())

	C.gtk_tree_store_swap((*C.GtkTreeStore)(recv.native), c_a, c_b)

	return
}

func (recv *TreeStore) Object() *gobject.Object {}

// Unsupported : gtk_tree_view_new_with_model : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

// ExpandToPath is a wrapper around the C function gtk_tree_view_expand_to_path.
func (recv *TreeView) ExpandToPath(path *TreePath) {
	c_path := (*C.GtkTreePath)(path.ToC())

	C.gtk_tree_view_expand_to_path((*C.GtkTreeView)(recv.native), c_path)

	return
}

// SetCursorOnCell is a wrapper around the C function gtk_tree_view_set_cursor_on_cell.
func (recv *TreeView) SetCursorOnCell(path *TreePath, focusColumn *TreeViewColumn, focusCell *CellRenderer, startEditing bool) {
	c_path := (*C.GtkTreePath)(path.ToC())

	c_focus_column := (*C.GtkTreeViewColumn)(focusColumn.ToC())

	c_focus_cell := (*C.GtkCellRenderer)(focusCell.ToC())

	c_start_editing :=
		boolToGboolean(startEditing)

	C.gtk_tree_view_set_cursor_on_cell((*C.GtkTreeView)(recv.native), c_path, c_focus_column, c_focus_cell, c_start_editing)

	return
}

func (recv *TreeView) Container() *Container {}

func (recv *TreeViewAccessible) ContainerAccessible() *ContainerAccessible {}

// Unsupported : gtk_tree_view_column_new_with_attributes : unsupported parameter ... : varargs

// FocusCell is a wrapper around the C function gtk_tree_view_column_focus_cell.
func (recv *TreeViewColumn) FocusCell(cell *CellRenderer) {
	c_cell := (*C.GtkCellRenderer)(cell.ToC())

	C.gtk_tree_view_column_focus_cell((*C.GtkTreeViewColumn)(recv.native), c_cell)

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

// Unsupported : gtk_widget_get_clipboard : unsupported parameter selection : Blacklisted record : GdkAtom

// GetDisplay is a wrapper around the C function gtk_widget_get_display.
func (recv *Widget) GetDisplay() *gdk.Display {
	retC := C.gtk_widget_get_display((*C.GtkWidget)(recv.native))
	retGo := gdk.DisplayNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetRootWindow is a wrapper around the C function gtk_widget_get_root_window.
func (recv *Widget) GetRootWindow() *gdk.Window {
	retC := C.gtk_widget_get_root_window((*C.GtkWidget)(recv.native))
	retGo := gdk.WindowNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetScreen is a wrapper around the C function gtk_widget_get_screen.
func (recv *Widget) GetScreen() *gdk.Screen {
	retC := C.gtk_widget_get_screen((*C.GtkWidget)(recv.native))
	retGo := gdk.ScreenNewFromC(unsafe.Pointer(retC))

	return retGo
}

// HasScreen is a wrapper around the C function gtk_widget_has_screen.
func (recv *Widget) HasScreen() bool {
	retC := C.gtk_widget_has_screen((*C.GtkWidget)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

func (recv *Widget) InitiallyUnowned() *gobject.InitiallyUnowned {}

func (recv *WidgetAccessible) Accessible() *Accessible {}

// Fullscreen is a wrapper around the C function gtk_window_fullscreen.
func (recv *Window) Fullscreen() {
	C.gtk_window_fullscreen((*C.GtkWindow)(recv.native))

	return
}

// GetScreen is a wrapper around the C function gtk_window_get_screen.
func (recv *Window) GetScreen() *gdk.Screen {
	retC := C.gtk_window_get_screen((*C.GtkWindow)(recv.native))
	retGo := gdk.ScreenNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetSkipPagerHint is a wrapper around the C function gtk_window_get_skip_pager_hint.
func (recv *Window) GetSkipPagerHint() bool {
	retC := C.gtk_window_get_skip_pager_hint((*C.GtkWindow)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetSkipTaskbarHint is a wrapper around the C function gtk_window_get_skip_taskbar_hint.
func (recv *Window) GetSkipTaskbarHint() bool {
	retC := C.gtk_window_get_skip_taskbar_hint((*C.GtkWindow)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetIconFromFile is a wrapper around the C function gtk_window_set_icon_from_file.
func (recv *Window) SetIconFromFile(filename string) (bool, error) {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))

	var cThrowableError *C.GError

	retC := C.gtk_window_set_icon_from_file((*C.GtkWindow)(recv.native), c_filename, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// SetScreen is a wrapper around the C function gtk_window_set_screen.
func (recv *Window) SetScreen(screen *gdk.Screen) {
	c_screen := (*C.GdkScreen)(screen.ToC())

	C.gtk_window_set_screen((*C.GtkWindow)(recv.native), c_screen)

	return
}

// SetSkipPagerHint is a wrapper around the C function gtk_window_set_skip_pager_hint.
func (recv *Window) SetSkipPagerHint(setting bool) {
	c_setting :=
		boolToGboolean(setting)

	C.gtk_window_set_skip_pager_hint((*C.GtkWindow)(recv.native), c_setting)

	return
}

// SetSkipTaskbarHint is a wrapper around the C function gtk_window_set_skip_taskbar_hint.
func (recv *Window) SetSkipTaskbarHint(setting bool) {
	c_setting :=
		boolToGboolean(setting)

	C.gtk_window_set_skip_taskbar_hint((*C.GtkWindow)(recv.native), c_setting)

	return
}

// Unfullscreen is a wrapper around the C function gtk_window_unfullscreen.
func (recv *Window) Unfullscreen() {
	C.gtk_window_unfullscreen((*C.GtkWindow)(recv.native))

	return
}

func (recv *Window) Bin() *Bin {}

func (recv *WindowAccessible) ContainerAccessible() *ContainerAccessible {}

func (recv *WindowGroup) Object() *gobject.Object {}
