// This is a generated file - DO NOT EDIT
// +build gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	atk "github.com/pekim/gobbi/lib/atk"
	gdk "github.com/pekim/gobbi/lib/gdk"
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

// SetAccel is a wrapper around the C function gtk_accel_label_set_accel.
func (recv *AccelLabel) SetAccel(acceleratorKey uint32, acceleratorMods gdk.ModifierType) {
	c_accelerator_key := (C.guint)(acceleratorKey)

	c_accelerator_mods := (C.GdkModifierType)(acceleratorMods)

	C.gtk_accel_label_set_accel((*C.GtkAccelLabel)(recv.native), c_accelerator_key, c_accelerator_mods)

	return
}

func (recv *AccelLabel) Label() *Label {}

func (recv *AccelMap) Object() *gobject.Object {}

func (recv *Accessible) Object() *atk.Object {}

func (recv *Action) Object() *gobject.Object {}

func (recv *ActionBar) Bin() *Bin {}

// GetAccelGroup is a wrapper around the C function gtk_action_group_get_accel_group.
func (recv *ActionGroup) GetAccelGroup() *AccelGroup {
	retC := C.gtk_action_group_get_accel_group((*C.GtkActionGroup)(recv.native))
	retGo := AccelGroupNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SetAccelGroup is a wrapper around the C function gtk_action_group_set_accel_group.
func (recv *ActionGroup) SetAccelGroup(accelGroup *AccelGroup) {
	c_accel_group := (*C.GtkAccelGroup)(accelGroup.ToC())

	C.gtk_action_group_set_accel_group((*C.GtkActionGroup)(recv.native), c_accel_group)

	return
}

func (recv *ActionGroup) Object() *gobject.Object {}

func (recv *Adjustment) InitiallyUnowned() *gobject.InitiallyUnowned {}

func (recv *Alignment) Bin() *Bin {}

func (recv *AppChooserButton) ComboBox() *ComboBox {}

// Unsupported : gtk_app_chooser_dialog_new : unsupported parameter file : no type generator for Gio.File, GFile*

func (recv *AppChooserDialog) Dialog() *Dialog {}

func (recv *AppChooserWidget) Box() *Box {}

// GetActiveWindow is a wrapper around the C function gtk_application_get_active_window.
func (recv *Application) GetActiveWindow() *Window {
	retC := C.gtk_application_get_active_window((*C.GtkApplication)(recv.native))
	retGo := WindowNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetWindowById is a wrapper around the C function gtk_application_get_window_by_id.
func (recv *Application) GetWindowById(id uint32) *Window {
	c_id := (C.guint)(id)

	retC := C.gtk_application_get_window_by_id((*C.GtkApplication)(recv.native), c_id)
	retGo := WindowNewFromC(unsafe.Pointer(retC))

	return retGo
}

func (recv *Application) Application() *gio.Application {}

// GetId is a wrapper around the C function gtk_application_window_get_id.
func (recv *ApplicationWindow) GetId() uint32 {
	retC := C.gtk_application_window_get_id((*C.GtkApplicationWindow)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

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

// GetAlwaysShowImage is a wrapper around the C function gtk_button_get_always_show_image.
func (recv *Button) GetAlwaysShowImage() bool {
	retC := C.gtk_button_get_always_show_image((*C.GtkButton)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetAlwaysShowImage is a wrapper around the C function gtk_button_set_always_show_image.
func (recv *Button) SetAlwaysShowImage(alwaysShow bool) {
	c_always_show :=
		boolToGboolean(alwaysShow)

	C.gtk_button_set_always_show_image((*C.GtkButton)(recv.native), c_always_show)

	return
}

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

// GetAttributes is a wrapper around the C function gtk_entry_get_attributes.
func (recv *Entry) GetAttributes() *pango.AttrList {
	retC := C.gtk_entry_get_attributes((*C.GtkEntry)(recv.native))
	retGo := pango.AttrListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetInputHints is a wrapper around the C function gtk_entry_get_input_hints.
func (recv *Entry) GetInputHints() InputHints {
	retC := C.gtk_entry_get_input_hints((*C.GtkEntry)(recv.native))
	retGo := (InputHints)(retC)

	return retGo
}

// GetInputPurpose is a wrapper around the C function gtk_entry_get_input_purpose.
func (recv *Entry) GetInputPurpose() InputPurpose {
	retC := C.gtk_entry_get_input_purpose((*C.GtkEntry)(recv.native))
	retGo := (InputPurpose)(retC)

	return retGo
}

// SetAttributes is a wrapper around the C function gtk_entry_set_attributes.
func (recv *Entry) SetAttributes(attrs *pango.AttrList) {
	c_attrs := (*C.PangoAttrList)(attrs.ToC())

	C.gtk_entry_set_attributes((*C.GtkEntry)(recv.native), c_attrs)

	return
}

// SetInputHints is a wrapper around the C function gtk_entry_set_input_hints.
func (recv *Entry) SetInputHints(hints InputHints) {
	c_hints := (C.GtkInputHints)(hints)

	C.gtk_entry_set_input_hints((*C.GtkEntry)(recv.native), c_hints)

	return
}

// SetInputPurpose is a wrapper around the C function gtk_entry_set_input_purpose.
func (recv *Entry) SetInputPurpose(purpose InputPurpose) {
	c_purpose := (C.GtkInputPurpose)(purpose)

	C.gtk_entry_set_input_purpose((*C.GtkEntry)(recv.native), c_purpose)

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

// Unsupported : gtk_icon_view_get_cell_rect : unsupported parameter rect : Blacklisted record : GdkRectangle

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

// LevelBarNew is a wrapper around the C function gtk_level_bar_new.
func LevelBarNew() *LevelBar {
	retC := C.gtk_level_bar_new()
	retGo := LevelBarNewFromC(unsafe.Pointer(retC))

	return retGo
}

// LevelBarNewForInterval is a wrapper around the C function gtk_level_bar_new_for_interval.
func LevelBarNewForInterval(minValue float64, maxValue float64) *LevelBar {
	c_min_value := (C.gdouble)(minValue)

	c_max_value := (C.gdouble)(maxValue)

	retC := C.gtk_level_bar_new_for_interval(c_min_value, c_max_value)
	retGo := LevelBarNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AddOffsetValue is a wrapper around the C function gtk_level_bar_add_offset_value.
func (recv *LevelBar) AddOffsetValue(name string, value float64) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_value := (C.gdouble)(value)

	C.gtk_level_bar_add_offset_value((*C.GtkLevelBar)(recv.native), c_name, c_value)

	return
}

// GetMaxValue is a wrapper around the C function gtk_level_bar_get_max_value.
func (recv *LevelBar) GetMaxValue() float64 {
	retC := C.gtk_level_bar_get_max_value((*C.GtkLevelBar)(recv.native))
	retGo := (float64)(retC)

	return retGo
}

// GetMinValue is a wrapper around the C function gtk_level_bar_get_min_value.
func (recv *LevelBar) GetMinValue() float64 {
	retC := C.gtk_level_bar_get_min_value((*C.GtkLevelBar)(recv.native))
	retGo := (float64)(retC)

	return retGo
}

// GetMode is a wrapper around the C function gtk_level_bar_get_mode.
func (recv *LevelBar) GetMode() LevelBarMode {
	retC := C.gtk_level_bar_get_mode((*C.GtkLevelBar)(recv.native))
	retGo := (LevelBarMode)(retC)

	return retGo
}

// GetOffsetValue is a wrapper around the C function gtk_level_bar_get_offset_value.
func (recv *LevelBar) GetOffsetValue(name string) (bool, *float64) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	var c_value C.gdouble

	retC := C.gtk_level_bar_get_offset_value((*C.GtkLevelBar)(recv.native), c_name, &c_value)
	retGo := retC == C.TRUE

	value := (*float64)(&c_value)

	return retGo, value
}

// GetValue is a wrapper around the C function gtk_level_bar_get_value.
func (recv *LevelBar) GetValue() float64 {
	retC := C.gtk_level_bar_get_value((*C.GtkLevelBar)(recv.native))
	retGo := (float64)(retC)

	return retGo
}

// RemoveOffsetValue is a wrapper around the C function gtk_level_bar_remove_offset_value.
func (recv *LevelBar) RemoveOffsetValue(name string) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	C.gtk_level_bar_remove_offset_value((*C.GtkLevelBar)(recv.native), c_name)

	return
}

// SetMaxValue is a wrapper around the C function gtk_level_bar_set_max_value.
func (recv *LevelBar) SetMaxValue(value float64) {
	c_value := (C.gdouble)(value)

	C.gtk_level_bar_set_max_value((*C.GtkLevelBar)(recv.native), c_value)

	return
}

// SetMinValue is a wrapper around the C function gtk_level_bar_set_min_value.
func (recv *LevelBar) SetMinValue(value float64) {
	c_value := (C.gdouble)(value)

	C.gtk_level_bar_set_min_value((*C.GtkLevelBar)(recv.native), c_value)

	return
}

// SetMode is a wrapper around the C function gtk_level_bar_set_mode.
func (recv *LevelBar) SetMode(mode LevelBarMode) {
	c_mode := (C.GtkLevelBarMode)(mode)

	C.gtk_level_bar_set_mode((*C.GtkLevelBar)(recv.native), c_mode)

	return
}

// SetValue is a wrapper around the C function gtk_level_bar_set_value.
func (recv *LevelBar) SetValue(value float64) {
	c_value := (C.gdouble)(value)

	C.gtk_level_bar_set_value((*C.GtkLevelBar)(recv.native), c_value)

	return
}

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

// MenuButtonNew is a wrapper around the C function gtk_menu_button_new.
func MenuButtonNew() *MenuButton {
	retC := C.gtk_menu_button_new()
	retGo := MenuButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetAlignWidget is a wrapper around the C function gtk_menu_button_get_align_widget.
func (recv *MenuButton) GetAlignWidget() *Widget {
	retC := C.gtk_menu_button_get_align_widget((*C.GtkMenuButton)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetDirection is a wrapper around the C function gtk_menu_button_get_direction.
func (recv *MenuButton) GetDirection() ArrowType {
	retC := C.gtk_menu_button_get_direction((*C.GtkMenuButton)(recv.native))
	retGo := (ArrowType)(retC)

	return retGo
}

// GetMenuModel is a wrapper around the C function gtk_menu_button_get_menu_model.
func (recv *MenuButton) GetMenuModel() *gio.MenuModel {
	retC := C.gtk_menu_button_get_menu_model((*C.GtkMenuButton)(recv.native))
	retGo := gio.MenuModelNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetPopup is a wrapper around the C function gtk_menu_button_get_popup.
func (recv *MenuButton) GetPopup() *Menu {
	retC := C.gtk_menu_button_get_popup((*C.GtkMenuButton)(recv.native))
	retGo := MenuNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SetAlignWidget is a wrapper around the C function gtk_menu_button_set_align_widget.
func (recv *MenuButton) SetAlignWidget(alignWidget *Widget) {
	c_align_widget := (*C.GtkWidget)(alignWidget.ToC())

	C.gtk_menu_button_set_align_widget((*C.GtkMenuButton)(recv.native), c_align_widget)

	return
}

// SetDirection is a wrapper around the C function gtk_menu_button_set_direction.
func (recv *MenuButton) SetDirection(direction ArrowType) {
	c_direction := (C.GtkArrowType)(direction)

	C.gtk_menu_button_set_direction((*C.GtkMenuButton)(recv.native), c_direction)

	return
}

// SetMenuModel is a wrapper around the C function gtk_menu_button_set_menu_model.
func (recv *MenuButton) SetMenuModel(menuModel *gio.MenuModel) {
	c_menu_model := (*C.GMenuModel)(menuModel.ToC())

	C.gtk_menu_button_set_menu_model((*C.GtkMenuButton)(recv.native), c_menu_model)

	return
}

// SetPopup is a wrapper around the C function gtk_menu_button_set_popup.
func (recv *MenuButton) SetPopup(menu *Widget) {
	c_menu := (*C.GtkWidget)(menu.ToC())

	C.gtk_menu_button_set_popup((*C.GtkMenuButton)(recv.native), c_menu)

	return
}

func (recv *MenuButton) ToggleButton() *ToggleButton {}

func (recv *MenuButtonAccessible) ToggleButtonAccessible() *ToggleButtonAccessible {}

func (recv *MenuItem) Bin() *Bin {}

func (recv *MenuItemAccessible) ContainerAccessible() *ContainerAccessible {}

// BindModel is a wrapper around the C function gtk_menu_shell_bind_model.
func (recv *MenuShell) BindModel(model *gio.MenuModel, actionNamespace string, withSeparators bool) {
	c_model := (*C.GMenuModel)(model.ToC())

	c_action_namespace := C.CString(actionNamespace)
	defer C.free(unsafe.Pointer(c_action_namespace))

	c_with_separators :=
		boolToGboolean(withSeparators)

	C.gtk_menu_shell_bind_model((*C.GtkMenuShell)(recv.native), c_model, c_action_namespace, c_with_separators)

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

func (recv *ScrolledWindow) Bin() *Bin {}

func (recv *ScrolledWindowAccessible) ContainerAccessible() *ContainerAccessible {}

func (recv *SearchBar) Bin() *Bin {}

// SearchEntryNew is a wrapper around the C function gtk_search_entry_new.
func SearchEntryNew() *SearchEntry {
	retC := C.gtk_search_entry_new()
	retGo := SearchEntryNewFromC(unsafe.Pointer(retC))

	return retGo
}

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

// GetInputHints is a wrapper around the C function gtk_text_view_get_input_hints.
func (recv *TextView) GetInputHints() InputHints {
	retC := C.gtk_text_view_get_input_hints((*C.GtkTextView)(recv.native))
	retGo := (InputHints)(retC)

	return retGo
}

// GetInputPurpose is a wrapper around the C function gtk_text_view_get_input_purpose.
func (recv *TextView) GetInputPurpose() InputPurpose {
	retC := C.gtk_text_view_get_input_purpose((*C.GtkTextView)(recv.native))
	retGo := (InputPurpose)(retC)

	return retGo
}

// SetInputHints is a wrapper around the C function gtk_text_view_set_input_hints.
func (recv *TextView) SetInputHints(hints InputHints) {
	c_hints := (C.GtkInputHints)(hints)

	C.gtk_text_view_set_input_hints((*C.GtkTextView)(recv.native), c_hints)

	return
}

// SetInputPurpose is a wrapper around the C function gtk_text_view_set_input_purpose.
func (recv *TextView) SetInputPurpose(purpose InputPurpose) {
	c_purpose := (C.GtkInputPurpose)(purpose)

	C.gtk_text_view_set_input_purpose((*C.GtkTextView)(recv.native), c_purpose)

	return
}

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

// Unsupported : gtk_widget_insert_action_group : unsupported parameter group : no type generator for Gio.ActionGroup, GActionGroup*

func (recv *Widget) InitiallyUnowned() *gobject.InitiallyUnowned {}

func (recv *WidgetAccessible) Accessible() *Accessible {}

func (recv *Window) Bin() *Bin {}

func (recv *WindowAccessible) ContainerAccessible() *ContainerAccessible {}

func (recv *WindowGroup) Object() *gobject.Object {}
