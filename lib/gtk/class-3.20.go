// This is a generated file - DO NOT EDIT
// +build gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	atk "github.com/pekim/gobbi/lib/atk"
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

// GetHelpOverlay is a wrapper around the C function gtk_application_window_get_help_overlay.
func (recv *ApplicationWindow) GetHelpOverlay() *ShortcutsWindow {
	retC := C.gtk_application_window_get_help_overlay((*C.GtkApplicationWindow)(recv.native))
	retGo := ShortcutsWindowNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SetHelpOverlay is a wrapper around the C function gtk_application_window_set_help_overlay.
func (recv *ApplicationWindow) SetHelpOverlay(helpOverlay *ShortcutsWindow) {
	c_help_overlay := (*C.GtkShortcutsWindow)(helpOverlay.ToC())

	C.gtk_application_window_set_help_overlay((*C.GtkApplicationWindow)(recv.native), c_help_overlay)

	return
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

// FileChooserNative is a wrapper around the C record GtkFileChooserNative.
type FileChooserNative struct {
	native *C.GtkFileChooserNative
}

func FileChooserNativeNewFromC(u unsafe.Pointer) *FileChooserNative {
	c := (*C.GtkFileChooserNative)(u)
	if c == nil {
		return nil
	}

	g := &FileChooserNative{native: c}

	return g
}

func (recv *FileChooserNative) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// FileChooserNativeNew is a wrapper around the C function gtk_file_chooser_native_new.
func FileChooserNativeNew(title string, parent *Window, action FileChooserAction, acceptLabel string, cancelLabel string) *FileChooserNative {
	c_title := C.CString(title)
	defer C.free(unsafe.Pointer(c_title))

	c_parent := (*C.GtkWindow)(parent.ToC())

	c_action := (C.GtkFileChooserAction)(action)

	c_accept_label := C.CString(acceptLabel)
	defer C.free(unsafe.Pointer(c_accept_label))

	c_cancel_label := C.CString(cancelLabel)
	defer C.free(unsafe.Pointer(c_cancel_label))

	retC := C.gtk_file_chooser_native_new(c_title, c_parent, c_action, c_accept_label, c_cancel_label)
	retGo := FileChooserNativeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetAcceptLabel is a wrapper around the C function gtk_file_chooser_native_get_accept_label.
func (recv *FileChooserNative) GetAcceptLabel() string {
	retC := C.gtk_file_chooser_native_get_accept_label((*C.GtkFileChooserNative)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetCancelLabel is a wrapper around the C function gtk_file_chooser_native_get_cancel_label.
func (recv *FileChooserNative) GetCancelLabel() string {
	retC := C.gtk_file_chooser_native_get_cancel_label((*C.GtkFileChooserNative)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// SetAcceptLabel is a wrapper around the C function gtk_file_chooser_native_set_accept_label.
func (recv *FileChooserNative) SetAcceptLabel(acceptLabel string) {
	c_accept_label := C.CString(acceptLabel)
	defer C.free(unsafe.Pointer(c_accept_label))

	C.gtk_file_chooser_native_set_accept_label((*C.GtkFileChooserNative)(recv.native), c_accept_label)

	return
}

// SetCancelLabel is a wrapper around the C function gtk_file_chooser_native_set_cancel_label.
func (recv *FileChooserNative) SetCancelLabel(cancelLabel string) {
	c_cancel_label := C.CString(cancelLabel)
	defer C.free(unsafe.Pointer(c_cancel_label))

	C.gtk_file_chooser_native_set_cancel_label((*C.GtkFileChooserNative)(recv.native), c_cancel_label)

	return
}

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

// NativeDialog is a wrapper around the C record GtkNativeDialog.
type NativeDialog struct {
	native *C.GtkNativeDialog
	// parent_instance : record
}

func NativeDialogNewFromC(u unsafe.Pointer) *NativeDialog {
	c := (*C.GtkNativeDialog)(u)
	if c == nil {
		return nil
	}

	g := &NativeDialog{native: c}

	return g
}

func (recv *NativeDialog) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Destroy is a wrapper around the C function gtk_native_dialog_destroy.
func (recv *NativeDialog) Destroy() {
	C.gtk_native_dialog_destroy((*C.GtkNativeDialog)(recv.native))

	return
}

// GetModal is a wrapper around the C function gtk_native_dialog_get_modal.
func (recv *NativeDialog) GetModal() bool {
	retC := C.gtk_native_dialog_get_modal((*C.GtkNativeDialog)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetTitle is a wrapper around the C function gtk_native_dialog_get_title.
func (recv *NativeDialog) GetTitle() string {
	retC := C.gtk_native_dialog_get_title((*C.GtkNativeDialog)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetTransientFor is a wrapper around the C function gtk_native_dialog_get_transient_for.
func (recv *NativeDialog) GetTransientFor() *Window {
	retC := C.gtk_native_dialog_get_transient_for((*C.GtkNativeDialog)(recv.native))
	retGo := WindowNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetVisible is a wrapper around the C function gtk_native_dialog_get_visible.
func (recv *NativeDialog) GetVisible() bool {
	retC := C.gtk_native_dialog_get_visible((*C.GtkNativeDialog)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Hide is a wrapper around the C function gtk_native_dialog_hide.
func (recv *NativeDialog) Hide() {
	C.gtk_native_dialog_hide((*C.GtkNativeDialog)(recv.native))

	return
}

// Run is a wrapper around the C function gtk_native_dialog_run.
func (recv *NativeDialog) Run() int32 {
	retC := C.gtk_native_dialog_run((*C.GtkNativeDialog)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// SetModal is a wrapper around the C function gtk_native_dialog_set_modal.
func (recv *NativeDialog) SetModal(modal bool) {
	c_modal :=
		boolToGboolean(modal)

	C.gtk_native_dialog_set_modal((*C.GtkNativeDialog)(recv.native), c_modal)

	return
}

// SetTitle is a wrapper around the C function gtk_native_dialog_set_title.
func (recv *NativeDialog) SetTitle(title string) {
	c_title := C.CString(title)
	defer C.free(unsafe.Pointer(c_title))

	C.gtk_native_dialog_set_title((*C.GtkNativeDialog)(recv.native), c_title)

	return
}

// SetTransientFor is a wrapper around the C function gtk_native_dialog_set_transient_for.
func (recv *NativeDialog) SetTransientFor(parent *Window) {
	c_parent := (*C.GtkWindow)(parent.ToC())

	C.gtk_native_dialog_set_transient_for((*C.GtkNativeDialog)(recv.native), c_parent)

	return
}

// Show is a wrapper around the C function gtk_native_dialog_show.
func (recv *NativeDialog) Show() {
	C.gtk_native_dialog_show((*C.GtkNativeDialog)(recv.native))

	return
}

func (recv *NativeDialog) Object() *gobject.Object {}

func (recv *Notebook) Container() *Container {}

func (recv *NotebookAccessible) ContainerAccessible() *ContainerAccessible {}

func (recv *NotebookPageAccessible) Object() *atk.Object {}

func (recv *NumerableIcon) EmblemedIcon() *gio.EmblemedIcon {}

func (recv *OffscreenWindow) Window() *Window {}

func (recv *Overlay) Bin() *Bin {}

// PadController is a wrapper around the C record GtkPadController.
type PadController struct {
	native *C.GtkPadController
}

func PadControllerNewFromC(u unsafe.Pointer) *PadController {
	c := (*C.GtkPadController)(u)
	if c == nil {
		return nil
	}

	g := &PadController{native: c}

	return g
}

func (recv *PadController) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gtk_pad_controller_new : unsupported parameter group : no type generator for Gio.ActionGroup, GActionGroup*

func (recv *PadController) EventController() *EventController {}

// Unsupported : gtk_page_setup_new_from_gvariant : unsupported parameter variant : Blacklisted record : GVariant

func (recv *PageSetup) Object() *gobject.Object {}

func (recv *Paned) Container() *Container {}

func (recv *PanedAccessible) ContainerAccessible() *ContainerAccessible {}

func (recv *PlacesSidebar) ScrolledWindow() *ScrolledWindow {}

// GetConstrainTo is a wrapper around the C function gtk_popover_get_constrain_to.
func (recv *Popover) GetConstrainTo() PopoverConstraint {
	retC := C.gtk_popover_get_constrain_to((*C.GtkPopover)(recv.native))
	retGo := (PopoverConstraint)(retC)

	return retGo
}

// SetConstrainTo is a wrapper around the C function gtk_popover_set_constrain_to.
func (recv *Popover) SetConstrainTo(constraint PopoverConstraint) {
	c_constraint := (C.GtkPopoverConstraint)(constraint)

	C.gtk_popover_set_constrain_to((*C.GtkPopover)(recv.native), c_constraint)

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

// ResetProperty is a wrapper around the C function gtk_settings_reset_property.
func (recv *Settings) ResetProperty(name string) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	C.gtk_settings_reset_property((*C.GtkSettings)(recv.native), c_name)

	return
}

func (recv *Settings) Object() *gobject.Object {}

// ShortcutLabel is a wrapper around the C record GtkShortcutLabel.
type ShortcutLabel struct {
	native *C.GtkShortcutLabel
}

func ShortcutLabelNewFromC(u unsafe.Pointer) *ShortcutLabel {
	c := (*C.GtkShortcutLabel)(u)
	if c == nil {
		return nil
	}

	g := &ShortcutLabel{native: c}

	return g
}

func (recv *ShortcutLabel) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

func (recv *ShortcutLabel) Box() *Box {}

// ShortcutsGroup is a wrapper around the C record GtkShortcutsGroup.
type ShortcutsGroup struct {
	native *C.GtkShortcutsGroup
}

func ShortcutsGroupNewFromC(u unsafe.Pointer) *ShortcutsGroup {
	c := (*C.GtkShortcutsGroup)(u)
	if c == nil {
		return nil
	}

	g := &ShortcutsGroup{native: c}

	return g
}

func (recv *ShortcutsGroup) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

func (recv *ShortcutsGroup) Box() *Box {}

// ShortcutsSection is a wrapper around the C record GtkShortcutsSection.
type ShortcutsSection struct {
	native *C.GtkShortcutsSection
}

func ShortcutsSectionNewFromC(u unsafe.Pointer) *ShortcutsSection {
	c := (*C.GtkShortcutsSection)(u)
	if c == nil {
		return nil
	}

	g := &ShortcutsSection{native: c}

	return g
}

func (recv *ShortcutsSection) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

func (recv *ShortcutsSection) Box() *Box {}

// ShortcutsShortcut is a wrapper around the C record GtkShortcutsShortcut.
type ShortcutsShortcut struct {
	native *C.GtkShortcutsShortcut
}

func ShortcutsShortcutNewFromC(u unsafe.Pointer) *ShortcutsShortcut {
	c := (*C.GtkShortcutsShortcut)(u)
	if c == nil {
		return nil
	}

	g := &ShortcutsShortcut{native: c}

	return g
}

func (recv *ShortcutsShortcut) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

func (recv *ShortcutsShortcut) Box() *Box {}

// ShortcutsWindow is a wrapper around the C record GtkShortcutsWindow.
type ShortcutsWindow struct {
	native *C.GtkShortcutsWindow
	// window : record
}

func ShortcutsWindowNewFromC(u unsafe.Pointer) *ShortcutsWindow {
	c := (*C.GtkShortcutsWindow)(u)
	if c == nil {
		return nil
	}

	g := &ShortcutsWindow{native: c}

	return g
}

func (recv *ShortcutsWindow) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

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

// ToString is a wrapper around the C function gtk_style_context_to_string.
func (recv *StyleContext) ToString(flags StyleContextPrintFlags) string {
	c_flags := (C.GtkStyleContextPrintFlags)(flags)

	retC := C.gtk_style_context_to_string((*C.GtkStyleContext)(recv.native), c_flags)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
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

// Changed is a wrapper around the C function gtk_text_tag_changed.
func (recv *TextTag) Changed(sizeChanged bool) {
	c_size_changed :=
		boolToGboolean(sizeChanged)

	C.gtk_text_tag_changed((*C.GtkTextTag)(recv.native), c_size_changed)

	return
}

func (recv *TextTag) Object() *gobject.Object {}

func (recv *TextTagTable) Object() *gobject.Object {}

// ResetCursorBlink is a wrapper around the C function gtk_text_view_reset_cursor_blink.
func (recv *TextView) ResetCursorBlink() {
	C.gtk_text_view_reset_cursor_blink((*C.GtkTextView)(recv.native))

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

// Unsupported : gtk_widget_get_allocated_size : unsupported parameter allocation : Blacklisted record : GdkRectangle

// GetFocusOnClick is a wrapper around the C function gtk_widget_get_focus_on_click.
func (recv *Widget) GetFocusOnClick() bool {
	retC := C.gtk_widget_get_focus_on_click((*C.GtkWidget)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// QueueAllocate is a wrapper around the C function gtk_widget_queue_allocate.
func (recv *Widget) QueueAllocate() {
	C.gtk_widget_queue_allocate((*C.GtkWidget)(recv.native))

	return
}

// SetFocusOnClick is a wrapper around the C function gtk_widget_set_focus_on_click.
func (recv *Widget) SetFocusOnClick(focusOnClick bool) {
	c_focus_on_click :=
		boolToGboolean(focusOnClick)

	C.gtk_widget_set_focus_on_click((*C.GtkWidget)(recv.native), c_focus_on_click)

	return
}

func (recv *Widget) InitiallyUnowned() *gobject.InitiallyUnowned {}

func (recv *WidgetAccessible) Accessible() *Accessible {}

func (recv *Window) Bin() *Bin {}

func (recv *WindowAccessible) ContainerAccessible() *ContainerAccessible {}

func (recv *WindowGroup) Object() *gobject.Object {}
