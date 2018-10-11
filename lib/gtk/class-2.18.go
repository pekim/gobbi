// This is a generated file - DO NOT EDIT
// +build gtk_2.18 gtk_2.20 gtk_2.22 gtk_2.24 gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	atk "github.com/pekim/gobbi/lib/atk"
	gdk "github.com/pekim/gobbi/lib/gdk"
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

// GetAlignment is a wrapper around the C function gtk_cell_renderer_get_alignment.
func (recv *CellRenderer) GetAlignment() (*float32, *float32) {
	var c_xalign C.gfloat

	var c_yalign C.gfloat

	C.gtk_cell_renderer_get_alignment((*C.GtkCellRenderer)(recv.native), &c_xalign, &c_yalign)

	xalign := (*float32)(&c_xalign)

	yalign := (*float32)(&c_yalign)

	return xalign, yalign
}

// GetPadding is a wrapper around the C function gtk_cell_renderer_get_padding.
func (recv *CellRenderer) GetPadding() (*int32, *int32) {
	var c_xpad C.gint

	var c_ypad C.gint

	C.gtk_cell_renderer_get_padding((*C.GtkCellRenderer)(recv.native), &c_xpad, &c_ypad)

	xpad := (*int32)(&c_xpad)

	ypad := (*int32)(&c_ypad)

	return xpad, ypad
}

// GetSensitive is a wrapper around the C function gtk_cell_renderer_get_sensitive.
func (recv *CellRenderer) GetSensitive() bool {
	retC := C.gtk_cell_renderer_get_sensitive((*C.GtkCellRenderer)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetVisible is a wrapper around the C function gtk_cell_renderer_get_visible.
func (recv *CellRenderer) GetVisible() bool {
	retC := C.gtk_cell_renderer_get_visible((*C.GtkCellRenderer)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetAlignment is a wrapper around the C function gtk_cell_renderer_set_alignment.
func (recv *CellRenderer) SetAlignment(xalign float32, yalign float32) {
	c_xalign := (C.gfloat)(xalign)

	c_yalign := (C.gfloat)(yalign)

	C.gtk_cell_renderer_set_alignment((*C.GtkCellRenderer)(recv.native), c_xalign, c_yalign)

	return
}

// SetPadding is a wrapper around the C function gtk_cell_renderer_set_padding.
func (recv *CellRenderer) SetPadding(xpad int32, ypad int32) {
	c_xpad := (C.gint)(xpad)

	c_ypad := (C.gint)(ypad)

	C.gtk_cell_renderer_set_padding((*C.GtkCellRenderer)(recv.native), c_xpad, c_ypad)

	return
}

// SetSensitive is a wrapper around the C function gtk_cell_renderer_set_sensitive.
func (recv *CellRenderer) SetSensitive(sensitive bool) {
	c_sensitive :=
		boolToGboolean(sensitive)

	C.gtk_cell_renderer_set_sensitive((*C.GtkCellRenderer)(recv.native), c_sensitive)

	return
}

// SetVisible is a wrapper around the C function gtk_cell_renderer_set_visible.
func (recv *CellRenderer) SetVisible(visible bool) {
	c_visible :=
		boolToGboolean(visible)

	C.gtk_cell_renderer_set_visible((*C.GtkCellRenderer)(recv.native), c_visible)

	return
}

func (recv *CellRenderer) InitiallyUnowned() *gobject.InitiallyUnowned {}

func (recv *CellRendererAccel) CellRendererText() *CellRendererText {}

func (recv *CellRendererCombo) CellRendererText() *CellRendererText {}

func (recv *CellRendererPixbuf) CellRenderer() *CellRenderer {}

func (recv *CellRendererProgress) CellRenderer() *CellRenderer {}

func (recv *CellRendererSpin) CellRendererText() *CellRendererText {}

func (recv *CellRendererSpinner) CellRenderer() *CellRenderer {}

func (recv *CellRendererText) CellRenderer() *CellRenderer {}

// GetActivatable is a wrapper around the C function gtk_cell_renderer_toggle_get_activatable.
func (recv *CellRendererToggle) GetActivatable() bool {
	retC := C.gtk_cell_renderer_toggle_get_activatable((*C.GtkCellRendererToggle)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetActivatable is a wrapper around the C function gtk_cell_renderer_toggle_set_activatable.
func (recv *CellRendererToggle) SetActivatable(setting bool) {
	c_setting :=
		boolToGboolean(setting)

	C.gtk_cell_renderer_toggle_set_activatable((*C.GtkCellRendererToggle)(recv.native), c_setting)

	return
}

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

// EntryNewWithBuffer is a wrapper around the C function gtk_entry_new_with_buffer.
func EntryNewWithBuffer(buffer *EntryBuffer) *Entry {
	c_buffer := (*C.GtkEntryBuffer)(buffer.ToC())

	retC := C.gtk_entry_new_with_buffer(c_buffer)
	retGo := EntryNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetBuffer is a wrapper around the C function gtk_entry_get_buffer.
func (recv *Entry) GetBuffer() *EntryBuffer {
	retC := C.gtk_entry_get_buffer((*C.GtkEntry)(recv.native))
	retGo := EntryBufferNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SetBuffer is a wrapper around the C function gtk_entry_set_buffer.
func (recv *Entry) SetBuffer(buffer *EntryBuffer) {
	c_buffer := (*C.GtkEntryBuffer)(buffer.ToC())

	C.gtk_entry_set_buffer((*C.GtkEntry)(recv.native), c_buffer)

	return
}

func (recv *Entry) Widget() *Widget {}

func (recv *EntryAccessible) WidgetAccessible() *WidgetAccessible {}

// EntryBufferNew is a wrapper around the C function gtk_entry_buffer_new.
func EntryBufferNew(initialChars string, nInitialChars int32) *EntryBuffer {
	c_initial_chars := C.CString(initialChars)
	defer C.free(unsafe.Pointer(c_initial_chars))

	c_n_initial_chars := (C.gint)(nInitialChars)

	retC := C.gtk_entry_buffer_new(c_initial_chars, c_n_initial_chars)
	retGo := EntryBufferNewFromC(unsafe.Pointer(retC))

	return retGo
}

// DeleteText is a wrapper around the C function gtk_entry_buffer_delete_text.
func (recv *EntryBuffer) DeleteText(position uint32, nChars int32) uint32 {
	c_position := (C.guint)(position)

	c_n_chars := (C.gint)(nChars)

	retC := C.gtk_entry_buffer_delete_text((*C.GtkEntryBuffer)(recv.native), c_position, c_n_chars)
	retGo := (uint32)(retC)

	return retGo
}

// EmitDeletedText is a wrapper around the C function gtk_entry_buffer_emit_deleted_text.
func (recv *EntryBuffer) EmitDeletedText(position uint32, nChars uint32) {
	c_position := (C.guint)(position)

	c_n_chars := (C.guint)(nChars)

	C.gtk_entry_buffer_emit_deleted_text((*C.GtkEntryBuffer)(recv.native), c_position, c_n_chars)

	return
}

// EmitInsertedText is a wrapper around the C function gtk_entry_buffer_emit_inserted_text.
func (recv *EntryBuffer) EmitInsertedText(position uint32, chars string, nChars uint32) {
	c_position := (C.guint)(position)

	c_chars := C.CString(chars)
	defer C.free(unsafe.Pointer(c_chars))

	c_n_chars := (C.guint)(nChars)

	C.gtk_entry_buffer_emit_inserted_text((*C.GtkEntryBuffer)(recv.native), c_position, c_chars, c_n_chars)

	return
}

// GetBytes is a wrapper around the C function gtk_entry_buffer_get_bytes.
func (recv *EntryBuffer) GetBytes() uint64 {
	retC := C.gtk_entry_buffer_get_bytes((*C.GtkEntryBuffer)(recv.native))
	retGo := (uint64)(retC)

	return retGo
}

// GetLength is a wrapper around the C function gtk_entry_buffer_get_length.
func (recv *EntryBuffer) GetLength() uint32 {
	retC := C.gtk_entry_buffer_get_length((*C.GtkEntryBuffer)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// GetMaxLength is a wrapper around the C function gtk_entry_buffer_get_max_length.
func (recv *EntryBuffer) GetMaxLength() int32 {
	retC := C.gtk_entry_buffer_get_max_length((*C.GtkEntryBuffer)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetText is a wrapper around the C function gtk_entry_buffer_get_text.
func (recv *EntryBuffer) GetText() string {
	retC := C.gtk_entry_buffer_get_text((*C.GtkEntryBuffer)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// InsertText is a wrapper around the C function gtk_entry_buffer_insert_text.
func (recv *EntryBuffer) InsertText(position uint32, chars string, nChars int32) uint32 {
	c_position := (C.guint)(position)

	c_chars := C.CString(chars)
	defer C.free(unsafe.Pointer(c_chars))

	c_n_chars := (C.gint)(nChars)

	retC := C.gtk_entry_buffer_insert_text((*C.GtkEntryBuffer)(recv.native), c_position, c_chars, c_n_chars)
	retGo := (uint32)(retC)

	return retGo
}

// SetMaxLength is a wrapper around the C function gtk_entry_buffer_set_max_length.
func (recv *EntryBuffer) SetMaxLength(maxLength int32) {
	c_max_length := (C.gint)(maxLength)

	C.gtk_entry_buffer_set_max_length((*C.GtkEntryBuffer)(recv.native), c_max_length)

	return
}

// SetText is a wrapper around the C function gtk_entry_buffer_set_text.
func (recv *EntryBuffer) SetText(chars string, nChars int32) {
	c_chars := C.CString(chars)
	defer C.free(unsafe.Pointer(c_chars))

	c_n_chars := (C.gint)(nChars)

	C.gtk_entry_buffer_set_text((*C.GtkEntryBuffer)(recv.native), c_chars, c_n_chars)

	return
}

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

// GetItemPadding is a wrapper around the C function gtk_icon_view_get_item_padding.
func (recv *IconView) GetItemPadding() int32 {
	retC := C.gtk_icon_view_get_item_padding((*C.GtkIconView)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// SetItemPadding is a wrapper around the C function gtk_icon_view_set_item_padding.
func (recv *IconView) SetItemPadding(itemPadding int32) {
	c_item_padding := (C.gint)(itemPadding)

	C.gtk_icon_view_set_item_padding((*C.GtkIconView)(recv.native), c_item_padding)

	return
}

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

// InfoBarNew is a wrapper around the C function gtk_info_bar_new.
func InfoBarNew() *InfoBar {
	retC := C.gtk_info_bar_new()
	retGo := InfoBarNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_info_bar_new_with_buttons : unsupported parameter ... : varargs

// AddActionWidget is a wrapper around the C function gtk_info_bar_add_action_widget.
func (recv *InfoBar) AddActionWidget(child *Widget, responseId int32) {
	c_child := (*C.GtkWidget)(child.ToC())

	c_response_id := (C.gint)(responseId)

	C.gtk_info_bar_add_action_widget((*C.GtkInfoBar)(recv.native), c_child, c_response_id)

	return
}

// AddButton is a wrapper around the C function gtk_info_bar_add_button.
func (recv *InfoBar) AddButton(buttonText string, responseId int32) *Button {
	c_button_text := C.CString(buttonText)
	defer C.free(unsafe.Pointer(c_button_text))

	c_response_id := (C.gint)(responseId)

	retC := C.gtk_info_bar_add_button((*C.GtkInfoBar)(recv.native), c_button_text, c_response_id)
	retGo := ButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_info_bar_add_buttons : unsupported parameter ... : varargs

// GetActionArea is a wrapper around the C function gtk_info_bar_get_action_area.
func (recv *InfoBar) GetActionArea() *Widget {
	retC := C.gtk_info_bar_get_action_area((*C.GtkInfoBar)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetContentArea is a wrapper around the C function gtk_info_bar_get_content_area.
func (recv *InfoBar) GetContentArea() *Widget {
	retC := C.gtk_info_bar_get_content_area((*C.GtkInfoBar)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetMessageType is a wrapper around the C function gtk_info_bar_get_message_type.
func (recv *InfoBar) GetMessageType() MessageType {
	retC := C.gtk_info_bar_get_message_type((*C.GtkInfoBar)(recv.native))
	retGo := (MessageType)(retC)

	return retGo
}

// Response is a wrapper around the C function gtk_info_bar_response.
func (recv *InfoBar) Response(responseId int32) {
	c_response_id := (C.gint)(responseId)

	C.gtk_info_bar_response((*C.GtkInfoBar)(recv.native), c_response_id)

	return
}

// SetDefaultResponse is a wrapper around the C function gtk_info_bar_set_default_response.
func (recv *InfoBar) SetDefaultResponse(responseId int32) {
	c_response_id := (C.gint)(responseId)

	C.gtk_info_bar_set_default_response((*C.GtkInfoBar)(recv.native), c_response_id)

	return
}

// SetMessageType is a wrapper around the C function gtk_info_bar_set_message_type.
func (recv *InfoBar) SetMessageType(messageType MessageType) {
	c_message_type := (C.GtkMessageType)(messageType)

	C.gtk_info_bar_set_message_type((*C.GtkInfoBar)(recv.native), c_message_type)

	return
}

// SetResponseSensitive is a wrapper around the C function gtk_info_bar_set_response_sensitive.
func (recv *InfoBar) SetResponseSensitive(responseId int32, setting bool) {
	c_response_id := (C.gint)(responseId)

	c_setting :=
		boolToGboolean(setting)

	C.gtk_info_bar_set_response_sensitive((*C.GtkInfoBar)(recv.native), c_response_id, c_setting)

	return
}

func (recv *InfoBar) Box() *Box {}

func (recv *Invisible) Widget() *Widget {}

// GetCurrentUri is a wrapper around the C function gtk_label_get_current_uri.
func (recv *Label) GetCurrentUri() string {
	retC := C.gtk_label_get_current_uri((*C.GtkLabel)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetTrackVisitedLinks is a wrapper around the C function gtk_label_get_track_visited_links.
func (recv *Label) GetTrackVisitedLinks() bool {
	retC := C.gtk_label_get_track_visited_links((*C.GtkLabel)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetTrackVisitedLinks is a wrapper around the C function gtk_label_set_track_visited_links.
func (recv *Label) SetTrackVisitedLinks(trackLinks bool) {
	c_track_links :=
		boolToGboolean(trackLinks)

	C.gtk_label_set_track_visited_links((*C.GtkLabel)(recv.native), c_track_links)

	return
}

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

// GetReserveToggleSize is a wrapper around the C function gtk_menu_get_reserve_toggle_size.
func (recv *Menu) GetReserveToggleSize() bool {
	retC := C.gtk_menu_get_reserve_toggle_size((*C.GtkMenu)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetReserveToggleSize is a wrapper around the C function gtk_menu_set_reserve_toggle_size.
func (recv *Menu) SetReserveToggleSize(reserveToggleSize bool) {
	c_reserve_toggle_size :=
		boolToGboolean(reserveToggleSize)

	C.gtk_menu_set_reserve_toggle_size((*C.GtkMenu)(recv.native), c_reserve_toggle_size)

	return
}

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

func (recv *PlacesSidebar) ScrolledWindow() *ScrolledWindow {}

func (recv *Popover) Bin() *Bin {}

func (recv *PopoverAccessible) ContainerAccessible() *ContainerAccessible {}

func (recv *PopoverMenu) Popover() *Popover {}

func (recv *PrintContext) Object() *gobject.Object {}

// GetEmbedPageSetup is a wrapper around the C function gtk_print_operation_get_embed_page_setup.
func (recv *PrintOperation) GetEmbedPageSetup() bool {
	retC := C.gtk_print_operation_get_embed_page_setup((*C.GtkPrintOperation)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetHasSelection is a wrapper around the C function gtk_print_operation_get_has_selection.
func (recv *PrintOperation) GetHasSelection() bool {
	retC := C.gtk_print_operation_get_has_selection((*C.GtkPrintOperation)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetNPagesToPrint is a wrapper around the C function gtk_print_operation_get_n_pages_to_print.
func (recv *PrintOperation) GetNPagesToPrint() int32 {
	retC := C.gtk_print_operation_get_n_pages_to_print((*C.GtkPrintOperation)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetSupportSelection is a wrapper around the C function gtk_print_operation_get_support_selection.
func (recv *PrintOperation) GetSupportSelection() bool {
	retC := C.gtk_print_operation_get_support_selection((*C.GtkPrintOperation)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetEmbedPageSetup is a wrapper around the C function gtk_print_operation_set_embed_page_setup.
func (recv *PrintOperation) SetEmbedPageSetup(embed bool) {
	c_embed :=
		boolToGboolean(embed)

	C.gtk_print_operation_set_embed_page_setup((*C.GtkPrintOperation)(recv.native), c_embed)

	return
}

// SetHasSelection is a wrapper around the C function gtk_print_operation_set_has_selection.
func (recv *PrintOperation) SetHasSelection(hasSelection bool) {
	c_has_selection :=
		boolToGboolean(hasSelection)

	C.gtk_print_operation_set_has_selection((*C.GtkPrintOperation)(recv.native), c_has_selection)

	return
}

// SetSupportSelection is a wrapper around the C function gtk_print_operation_set_support_selection.
func (recv *PrintOperation) SetSupportSelection(supportSelection bool) {
	c_support_selection :=
		boolToGboolean(supportSelection)

	C.gtk_print_operation_set_support_selection((*C.GtkPrintOperation)(recv.native), c_support_selection)

	return
}

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

// GetFlippable is a wrapper around the C function gtk_range_get_flippable.
func (recv *Range) GetFlippable() bool {
	retC := C.gtk_range_get_flippable((*C.GtkRange)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetFlippable is a wrapper around the C function gtk_range_set_flippable.
func (recv *Range) SetFlippable(flippable bool) {
	c_flippable :=
		boolToGboolean(flippable)

	C.gtk_range_set_flippable((*C.GtkRange)(recv.native), c_flippable)

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

func (recv *Spinner) Widget() *Widget {}

func (recv *SpinnerAccessible) WidgetAccessible() *WidgetAccessible {}

func (recv *Stack) Container() *Container {}

func (recv *StackSidebar) Bin() *Bin {}

func (recv *StackSwitcher) Box() *Box {}

// Unsupported : gtk_status_icon_new_from_gicon : unsupported parameter icon : no type generator for Gio.Icon, GIcon*

// GetTitle is a wrapper around the C function gtk_status_icon_get_title.
func (recv *StatusIcon) GetTitle() string {
	retC := C.gtk_status_icon_get_title((*C.GtkStatusIcon)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// SetTitle is a wrapper around the C function gtk_status_icon_set_title.
func (recv *StatusIcon) SetTitle(title string) {
	c_title := C.CString(title)
	defer C.free(unsafe.Pointer(c_title))

	C.gtk_status_icon_set_title((*C.GtkStatusIcon)(recv.native), c_title)

	return
}

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

// Unsupported : gtk_widget_get_allocation : unsupported parameter allocation : Blacklisted record : GdkRectangle

// GetAppPaintable is a wrapper around the C function gtk_widget_get_app_paintable.
func (recv *Widget) GetAppPaintable() bool {
	retC := C.gtk_widget_get_app_paintable((*C.GtkWidget)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetCanDefault is a wrapper around the C function gtk_widget_get_can_default.
func (recv *Widget) GetCanDefault() bool {
	retC := C.gtk_widget_get_can_default((*C.GtkWidget)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetCanFocus is a wrapper around the C function gtk_widget_get_can_focus.
func (recv *Widget) GetCanFocus() bool {
	retC := C.gtk_widget_get_can_focus((*C.GtkWidget)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetDoubleBuffered is a wrapper around the C function gtk_widget_get_double_buffered.
func (recv *Widget) GetDoubleBuffered() bool {
	retC := C.gtk_widget_get_double_buffered((*C.GtkWidget)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetHasWindow is a wrapper around the C function gtk_widget_get_has_window.
func (recv *Widget) GetHasWindow() bool {
	retC := C.gtk_widget_get_has_window((*C.GtkWidget)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetReceivesDefault is a wrapper around the C function gtk_widget_get_receives_default.
func (recv *Widget) GetReceivesDefault() bool {
	retC := C.gtk_widget_get_receives_default((*C.GtkWidget)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetSensitive is a wrapper around the C function gtk_widget_get_sensitive.
func (recv *Widget) GetSensitive() bool {
	retC := C.gtk_widget_get_sensitive((*C.GtkWidget)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetState is a wrapper around the C function gtk_widget_get_state.
func (recv *Widget) GetState() StateType {
	retC := C.gtk_widget_get_state((*C.GtkWidget)(recv.native))
	retGo := (StateType)(retC)

	return retGo
}

// GetVisible is a wrapper around the C function gtk_widget_get_visible.
func (recv *Widget) GetVisible() bool {
	retC := C.gtk_widget_get_visible((*C.GtkWidget)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// HasDefault is a wrapper around the C function gtk_widget_has_default.
func (recv *Widget) HasDefault() bool {
	retC := C.gtk_widget_has_default((*C.GtkWidget)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// HasFocus is a wrapper around the C function gtk_widget_has_focus.
func (recv *Widget) HasFocus() bool {
	retC := C.gtk_widget_has_focus((*C.GtkWidget)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// HasGrab is a wrapper around the C function gtk_widget_has_grab.
func (recv *Widget) HasGrab() bool {
	retC := C.gtk_widget_has_grab((*C.GtkWidget)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// IsDrawable is a wrapper around the C function gtk_widget_is_drawable.
func (recv *Widget) IsDrawable() bool {
	retC := C.gtk_widget_is_drawable((*C.GtkWidget)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// IsSensitive is a wrapper around the C function gtk_widget_is_sensitive.
func (recv *Widget) IsSensitive() bool {
	retC := C.gtk_widget_is_sensitive((*C.GtkWidget)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// IsToplevel is a wrapper around the C function gtk_widget_is_toplevel.
func (recv *Widget) IsToplevel() bool {
	retC := C.gtk_widget_is_toplevel((*C.GtkWidget)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_widget_set_allocation : unsupported parameter allocation : Blacklisted record : GdkRectangle

// SetCanDefault is a wrapper around the C function gtk_widget_set_can_default.
func (recv *Widget) SetCanDefault(canDefault bool) {
	c_can_default :=
		boolToGboolean(canDefault)

	C.gtk_widget_set_can_default((*C.GtkWidget)(recv.native), c_can_default)

	return
}

// SetCanFocus is a wrapper around the C function gtk_widget_set_can_focus.
func (recv *Widget) SetCanFocus(canFocus bool) {
	c_can_focus :=
		boolToGboolean(canFocus)

	C.gtk_widget_set_can_focus((*C.GtkWidget)(recv.native), c_can_focus)

	return
}

// SetHasWindow is a wrapper around the C function gtk_widget_set_has_window.
func (recv *Widget) SetHasWindow(hasWindow bool) {
	c_has_window :=
		boolToGboolean(hasWindow)

	C.gtk_widget_set_has_window((*C.GtkWidget)(recv.native), c_has_window)

	return
}

// SetReceivesDefault is a wrapper around the C function gtk_widget_set_receives_default.
func (recv *Widget) SetReceivesDefault(receivesDefault bool) {
	c_receives_default :=
		boolToGboolean(receivesDefault)

	C.gtk_widget_set_receives_default((*C.GtkWidget)(recv.native), c_receives_default)

	return
}

// SetVisible is a wrapper around the C function gtk_widget_set_visible.
func (recv *Widget) SetVisible(visible bool) {
	c_visible :=
		boolToGboolean(visible)

	C.gtk_widget_set_visible((*C.GtkWidget)(recv.native), c_visible)

	return
}

// SetWindow is a wrapper around the C function gtk_widget_set_window.
func (recv *Widget) SetWindow(window *gdk.Window) {
	c_window := (*C.GdkWindow)(window.ToC())

	C.gtk_widget_set_window((*C.GtkWidget)(recv.native), c_window)

	return
}

func (recv *Widget) InitiallyUnowned() *gobject.InitiallyUnowned {}

func (recv *WidgetAccessible) Accessible() *Accessible {}

func (recv *Window) Bin() *Bin {}

func (recv *WindowAccessible) ContainerAccessible() *ContainerAccessible {}

func (recv *WindowGroup) Object() *gobject.Object {}
