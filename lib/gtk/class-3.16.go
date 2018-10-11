// This is a generated file - DO NOT EDIT
// +build gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

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

// LoadFromResource is a wrapper around the C function gtk_css_provider_load_from_resource.
func (recv *CssProvider) LoadFromResource(resourcePath string) {
	c_resource_path := C.CString(resourcePath)
	defer C.free(unsafe.Pointer(c_resource_path))

	C.gtk_css_provider_load_from_resource((*C.GtkCssProvider)(recv.native), c_resource_path)

	return
}

func (recv *CssProvider) Object() *gobject.Object {}

// Unsupported : gtk_dialog_new_with_buttons : unsupported parameter ... : varargs

func (recv *Dialog) Window() *Window {}

func (recv *DrawingArea) Widget() *Widget {}

// GrabFocusWithoutSelecting is a wrapper around the C function gtk_entry_grab_focus_without_selecting.
func (recv *Entry) GrabFocusWithoutSelecting() {
	C.gtk_entry_grab_focus_without_selecting((*C.GtkEntry)(recv.native))

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

// GLArea is a wrapper around the C record GtkGLArea.
type GLArea struct {
	native *C.GtkGLArea
	// Private : parent_instance
}

func GLAreaNewFromC(u unsafe.Pointer) *GLArea {
	c := (*C.GtkGLArea)(u)
	if c == nil {
		return nil
	}

	g := &GLArea{native: c}

	return g
}

func (recv *GLArea) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// GLAreaNew is a wrapper around the C function gtk_gl_area_new.
func GLAreaNew() *GLArea {
	retC := C.gtk_gl_area_new()
	retGo := GLAreaNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AttachBuffers is a wrapper around the C function gtk_gl_area_attach_buffers.
func (recv *GLArea) AttachBuffers() {
	C.gtk_gl_area_attach_buffers((*C.GtkGLArea)(recv.native))

	return
}

// GetAutoRender is a wrapper around the C function gtk_gl_area_get_auto_render.
func (recv *GLArea) GetAutoRender() bool {
	retC := C.gtk_gl_area_get_auto_render((*C.GtkGLArea)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetContext is a wrapper around the C function gtk_gl_area_get_context.
func (recv *GLArea) GetContext() *gdk.GLContext {
	retC := C.gtk_gl_area_get_context((*C.GtkGLArea)(recv.native))
	retGo := gdk.GLContextNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetError is a wrapper around the C function gtk_gl_area_get_error.
func (recv *GLArea) GetError() *glib.Error {
	retC := C.gtk_gl_area_get_error((*C.GtkGLArea)(recv.native))
	retGo := glib.ErrorNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetHasAlpha is a wrapper around the C function gtk_gl_area_get_has_alpha.
func (recv *GLArea) GetHasAlpha() bool {
	retC := C.gtk_gl_area_get_has_alpha((*C.GtkGLArea)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetHasDepthBuffer is a wrapper around the C function gtk_gl_area_get_has_depth_buffer.
func (recv *GLArea) GetHasDepthBuffer() bool {
	retC := C.gtk_gl_area_get_has_depth_buffer((*C.GtkGLArea)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetHasStencilBuffer is a wrapper around the C function gtk_gl_area_get_has_stencil_buffer.
func (recv *GLArea) GetHasStencilBuffer() bool {
	retC := C.gtk_gl_area_get_has_stencil_buffer((*C.GtkGLArea)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetRequiredVersion is a wrapper around the C function gtk_gl_area_get_required_version.
func (recv *GLArea) GetRequiredVersion() (*int32, *int32) {
	var c_major C.gint

	var c_minor C.gint

	C.gtk_gl_area_get_required_version((*C.GtkGLArea)(recv.native), &c_major, &c_minor)

	major := (*int32)(&c_major)

	minor := (*int32)(&c_minor)

	return major, minor
}

// MakeCurrent is a wrapper around the C function gtk_gl_area_make_current.
func (recv *GLArea) MakeCurrent() {
	C.gtk_gl_area_make_current((*C.GtkGLArea)(recv.native))

	return
}

// QueueRender is a wrapper around the C function gtk_gl_area_queue_render.
func (recv *GLArea) QueueRender() {
	C.gtk_gl_area_queue_render((*C.GtkGLArea)(recv.native))

	return
}

// SetAutoRender is a wrapper around the C function gtk_gl_area_set_auto_render.
func (recv *GLArea) SetAutoRender(autoRender bool) {
	c_auto_render :=
		boolToGboolean(autoRender)

	C.gtk_gl_area_set_auto_render((*C.GtkGLArea)(recv.native), c_auto_render)

	return
}

// SetError is a wrapper around the C function gtk_gl_area_set_error.
func (recv *GLArea) SetError(error *glib.Error) {
	c_error := (*C.GError)(error.ToC())

	C.gtk_gl_area_set_error((*C.GtkGLArea)(recv.native), c_error)

	return
}

// SetHasAlpha is a wrapper around the C function gtk_gl_area_set_has_alpha.
func (recv *GLArea) SetHasAlpha(hasAlpha bool) {
	c_has_alpha :=
		boolToGboolean(hasAlpha)

	C.gtk_gl_area_set_has_alpha((*C.GtkGLArea)(recv.native), c_has_alpha)

	return
}

// SetHasDepthBuffer is a wrapper around the C function gtk_gl_area_set_has_depth_buffer.
func (recv *GLArea) SetHasDepthBuffer(hasDepthBuffer bool) {
	c_has_depth_buffer :=
		boolToGboolean(hasDepthBuffer)

	C.gtk_gl_area_set_has_depth_buffer((*C.GtkGLArea)(recv.native), c_has_depth_buffer)

	return
}

// SetHasStencilBuffer is a wrapper around the C function gtk_gl_area_set_has_stencil_buffer.
func (recv *GLArea) SetHasStencilBuffer(hasStencilBuffer bool) {
	c_has_stencil_buffer :=
		boolToGboolean(hasStencilBuffer)

	C.gtk_gl_area_set_has_stencil_buffer((*C.GtkGLArea)(recv.native), c_has_stencil_buffer)

	return
}

// SetRequiredVersion is a wrapper around the C function gtk_gl_area_set_required_version.
func (recv *GLArea) SetRequiredVersion(major int32, minor int32) {
	c_major := (C.gint)(major)

	c_minor := (C.gint)(minor)

	C.gtk_gl_area_set_required_version((*C.GtkGLArea)(recv.native), c_major, c_minor)

	return
}

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

// GetXalign is a wrapper around the C function gtk_label_get_xalign.
func (recv *Label) GetXalign() float32 {
	retC := C.gtk_label_get_xalign((*C.GtkLabel)(recv.native))
	retGo := (float32)(retC)

	return retGo
}

// GetYalign is a wrapper around the C function gtk_label_get_yalign.
func (recv *Label) GetYalign() float32 {
	retC := C.gtk_label_get_yalign((*C.GtkLabel)(recv.native))
	retGo := (float32)(retC)

	return retGo
}

// SetXalign is a wrapper around the C function gtk_label_set_xalign.
func (recv *Label) SetXalign(xalign float32) {
	c_xalign := (C.gfloat)(xalign)

	C.gtk_label_set_xalign((*C.GtkLabel)(recv.native), c_xalign)

	return
}

// SetYalign is a wrapper around the C function gtk_label_set_yalign.
func (recv *Label) SetYalign(yalign float32) {
	c_yalign := (C.gfloat)(yalign)

	C.gtk_label_set_yalign((*C.GtkLabel)(recv.native), c_yalign)

	return
}

func (recv *Label) Misc() *Misc {}

func (recv *LabelAccessible) WidgetAccessible() *WidgetAccessible {}

func (recv *Layout) Container() *Container {}

func (recv *LevelBar) Widget() *Widget {}

func (recv *LevelBarAccessible) WidgetAccessible() *WidgetAccessible {}

func (recv *LinkButton) Button() *Button {}

func (recv *LinkButtonAccessible) ButtonAccessible() *ButtonAccessible {}

// Unsupported : gtk_list_box_bind_model : unsupported parameter model : no type generator for Gio.ListModel, GListModel*

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

// ModelButtonNew is a wrapper around the C function gtk_model_button_new.
func ModelButtonNew() *ModelButton {
	retC := C.gtk_model_button_new()
	retGo := ModelButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

func (recv *ModelButton) Button() *Button {}

func (recv *MountOperation) MountOperation() *gio.MountOperation {}

func (recv *NativeDialog) Object() *gobject.Object {}

// DetachTab is a wrapper around the C function gtk_notebook_detach_tab.
func (recv *Notebook) DetachTab(child *Widget) {
	c_child := (*C.GtkWidget)(child.ToC())

	C.gtk_notebook_detach_tab((*C.GtkNotebook)(recv.native), c_child)

	return
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

// GetWideHandle is a wrapper around the C function gtk_paned_get_wide_handle.
func (recv *Paned) GetWideHandle() bool {
	retC := C.gtk_paned_get_wide_handle((*C.GtkPaned)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetWideHandle is a wrapper around the C function gtk_paned_set_wide_handle.
func (recv *Paned) SetWideHandle(wide bool) {
	c_wide :=
		boolToGboolean(wide)

	C.gtk_paned_set_wide_handle((*C.GtkPaned)(recv.native), c_wide)

	return
}

func (recv *Paned) Container() *Container {}

func (recv *PanedAccessible) ContainerAccessible() *ContainerAccessible {}

func (recv *PlacesSidebar) ScrolledWindow() *ScrolledWindow {}

// GetTransitionsEnabled is a wrapper around the C function gtk_popover_get_transitions_enabled.
func (recv *Popover) GetTransitionsEnabled() bool {
	retC := C.gtk_popover_get_transitions_enabled((*C.GtkPopover)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetTransitionsEnabled is a wrapper around the C function gtk_popover_set_transitions_enabled.
func (recv *Popover) SetTransitionsEnabled(transitionsEnabled bool) {
	c_transitions_enabled :=
		boolToGboolean(transitionsEnabled)

	C.gtk_popover_set_transitions_enabled((*C.GtkPopover)(recv.native), c_transitions_enabled)

	return
}

func (recv *Popover) Bin() *Bin {}

func (recv *PopoverAccessible) ContainerAccessible() *ContainerAccessible {}

// PopoverMenuNew is a wrapper around the C function gtk_popover_menu_new.
func PopoverMenuNew() *PopoverMenu {
	retC := C.gtk_popover_menu_new()
	retGo := PopoverMenuNewFromC(unsafe.Pointer(retC))

	return retGo
}

// OpenSubmenu is a wrapper around the C function gtk_popover_menu_open_submenu.
func (recv *PopoverMenu) OpenSubmenu(name string) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	C.gtk_popover_menu_open_submenu((*C.GtkPopoverMenu)(recv.native), c_name)

	return
}

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

// GetOverlayScrolling is a wrapper around the C function gtk_scrolled_window_get_overlay_scrolling.
func (recv *ScrolledWindow) GetOverlayScrolling() bool {
	retC := C.gtk_scrolled_window_get_overlay_scrolling((*C.GtkScrolledWindow)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetOverlayScrolling is a wrapper around the C function gtk_scrolled_window_set_overlay_scrolling.
func (recv *ScrolledWindow) SetOverlayScrolling(overlayScrolling bool) {
	c_overlay_scrolling :=
		boolToGboolean(overlayScrolling)

	C.gtk_scrolled_window_set_overlay_scrolling((*C.GtkScrolledWindow)(recv.native), c_overlay_scrolling)

	return
}

func (recv *ScrolledWindow) Bin() *Bin {}

func (recv *ScrolledWindowAccessible) ContainerAccessible() *ContainerAccessible {}

func (recv *SearchBar) Bin() *Bin {}

// Unsupported : gtk_search_entry_handle_event : unsupported parameter event : no type generator for Gdk.Event, GdkEvent*

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

// GetHhomogeneous is a wrapper around the C function gtk_stack_get_hhomogeneous.
func (recv *Stack) GetHhomogeneous() bool {
	retC := C.gtk_stack_get_hhomogeneous((*C.GtkStack)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetVhomogeneous is a wrapper around the C function gtk_stack_get_vhomogeneous.
func (recv *Stack) GetVhomogeneous() bool {
	retC := C.gtk_stack_get_vhomogeneous((*C.GtkStack)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetHhomogeneous is a wrapper around the C function gtk_stack_set_hhomogeneous.
func (recv *Stack) SetHhomogeneous(hhomogeneous bool) {
	c_hhomogeneous :=
		boolToGboolean(hhomogeneous)

	C.gtk_stack_set_hhomogeneous((*C.GtkStack)(recv.native), c_hhomogeneous)

	return
}

// SetVhomogeneous is a wrapper around the C function gtk_stack_set_vhomogeneous.
func (recv *Stack) SetVhomogeneous(vhomogeneous bool) {
	c_vhomogeneous :=
		boolToGboolean(vhomogeneous)

	C.gtk_stack_set_vhomogeneous((*C.GtkStack)(recv.native), c_vhomogeneous)

	return
}

func (recv *Stack) Container() *Container {}

// StackSidebarNew is a wrapper around the C function gtk_stack_sidebar_new.
func StackSidebarNew() *StackSidebar {
	retC := C.gtk_stack_sidebar_new()
	retGo := StackSidebarNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetStack is a wrapper around the C function gtk_stack_sidebar_get_stack.
func (recv *StackSidebar) GetStack() *Stack {
	retC := C.gtk_stack_sidebar_get_stack((*C.GtkStackSidebar)(recv.native))
	retGo := StackNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SetStack is a wrapper around the C function gtk_stack_sidebar_set_stack.
func (recv *StackSidebar) SetStack(stack *Stack) {
	c_stack := (*C.GtkStack)(stack.ToC())

	C.gtk_stack_sidebar_set_stack((*C.GtkStackSidebar)(recv.native), c_stack)

	return
}

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

// InsertMarkup is a wrapper around the C function gtk_text_buffer_insert_markup.
func (recv *TextBuffer) InsertMarkup(iter *TextIter, markup string, len int32) {
	c_iter := (*C.GtkTextIter)(iter.ToC())

	c_markup := C.CString(markup)
	defer C.free(unsafe.Pointer(c_markup))

	c_len := (C.gint)(len)

	C.gtk_text_buffer_insert_markup((*C.GtkTextBuffer)(recv.native), c_iter, c_markup, c_len)

	return
}

func (recv *TextBuffer) Object() *gobject.Object {}

func (recv *TextCellAccessible) RendererCellAccessible() *RendererCellAccessible {}

func (recv *TextChildAnchor) Object() *gobject.Object {}

func (recv *TextMark) Object() *gobject.Object {}

func (recv *TextTag) Object() *gobject.Object {}

func (recv *TextTagTable) Object() *gobject.Object {}

// GetMonospace is a wrapper around the C function gtk_text_view_get_monospace.
func (recv *TextView) GetMonospace() bool {
	retC := C.gtk_text_view_get_monospace((*C.GtkTextView)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetMonospace is a wrapper around the C function gtk_text_view_set_monospace.
func (recv *TextView) SetMonospace(monospace bool) {
	c_monospace :=
		boolToGboolean(monospace)

	C.gtk_text_view_set_monospace((*C.GtkTextView)(recv.native), c_monospace)

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

// Unsupported : gtk_widget_get_action_group : no return generator

// Unsupported : gtk_widget_list_action_prefixes : no return type

func (recv *Widget) InitiallyUnowned() *gobject.InitiallyUnowned {}

func (recv *WidgetAccessible) Accessible() *Accessible {}

// GetTitlebar is a wrapper around the C function gtk_window_get_titlebar.
func (recv *Window) GetTitlebar() *Widget {
	retC := C.gtk_window_get_titlebar((*C.GtkWindow)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

func (recv *Window) Bin() *Bin {}

func (recv *WindowAccessible) ContainerAccessible() *ContainerAccessible {}

func (recv *WindowGroup) Object() *gobject.Object {}
