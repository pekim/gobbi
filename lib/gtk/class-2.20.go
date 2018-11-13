// This is a generated file - DO NOT EDIT
// +build gtk_2.20 gtk_2.22 gtk_2.24 gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	cairo "github.com/pekim/gobbi/lib/cairo"
	gdk "github.com/pekim/gobbi/lib/gdk"
	gdkpixbuf "github.com/pekim/gobbi/lib/gdkpixbuf"
	gio "github.com/pekim/gobbi/lib/gio"
	pango "github.com/pekim/gobbi/lib/pango"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

// Returns whether @action's menu item proxies will always
// show their image, if available.
/*

C function : gtk_action_get_always_show_image
*/
func (recv *Action) GetAlwaysShowImage() bool {
	retC := C.gtk_action_get_always_show_image((*C.GtkAction)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Sets whether @action's menu item proxies will ignore the
// #GtkSettings:gtk-menu-images setting and always show their image, if available.
//
// Use this if the menu item would be useless or hard to use
// without their image.
/*

C function : gtk_action_set_always_show_image
*/
func (recv *Action) SetAlwaysShowImage(alwaysShow bool) {
	c_always_show :=
		boolToGboolean(alwaysShow)

	C.gtk_action_set_always_show_image((*C.GtkAction)(recv.native), c_always_show)

	return
}

// Returns a new cell renderer which will show a spinner to indicate
// activity.
/*

C function : gtk_cell_renderer_spinner_new
*/
func CellRendererSpinnerNew() *CellRendererSpinner {
	retC := C.gtk_cell_renderer_spinner_new()
	retGo := CellRendererSpinnerNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the widget button that uses the given response ID in the action area
// of a dialog.
/*

C function : gtk_dialog_get_widget_for_response
*/
func (recv *Dialog) GetWidgetForResponse(responseId int32) *Widget {
	c_response_id := (C.gint)(responseId)

	retC := C.gtk_dialog_get_widget_for_response((*C.GtkDialog)(recv.native), c_response_id)
	var retGo (*Widget)
	if retC == nil {
		retGo = nil
	} else {
		retGo = WidgetNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Unsupported signal 'preedit-changed' for Entry : unsupported parameter preedit : type utf8 :

// Gets one of the action widgets. See gtk_notebook_set_action_widget().
/*

C function : gtk_notebook_get_action_widget
*/
func (recv *Notebook) GetActionWidget(packType PackType) *Widget {
	c_pack_type := (C.GtkPackType)(packType)

	retC := C.gtk_notebook_get_action_widget((*C.GtkNotebook)(recv.native), c_pack_type)
	var retGo (*Widget)
	if retC == nil {
		retGo = nil
	} else {
		retGo = WidgetNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Sets @widget as one of the action widgets. Depending on the pack type
// the widget will be placed before or after the tabs. You can use
// a #GtkBox if you need to pack more than one widget on the same side.
//
// Note that action widgets are “internal” children of the notebook and thus
// not included in the list returned from gtk_container_foreach().
/*

C function : gtk_notebook_set_action_widget
*/
func (recv *Notebook) SetActionWidget(widget *Widget, packType PackType) {
	c_widget := (*C.GtkWidget)(C.NULL)
	if widget != nil {
		c_widget = (*C.GtkWidget)(widget.ToC())
	}

	c_pack_type := (C.GtkPackType)(packType)

	C.gtk_notebook_set_action_widget((*C.GtkNotebook)(recv.native), c_widget, c_pack_type)

	return
}

// Creates a toplevel container widget that is used to retrieve
// snapshots of widgets without showing them on the screen.
/*

C function : gtk_offscreen_window_new
*/
func OffscreenWindowNew() *OffscreenWindow {
	retC := C.gtk_offscreen_window_new()
	retGo := OffscreenWindowNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Retrieves a snapshot of the contained widget in the form of
// a #GdkPixbuf.  This is a new pixbuf with a reference count of 1,
// and the application should unreference it once it is no longer
// needed.
/*

C function : gtk_offscreen_window_get_pixbuf
*/
func (recv *OffscreenWindow) GetPixbuf() *gdkpixbuf.Pixbuf {
	retC := C.gtk_offscreen_window_get_pixbuf((*C.GtkOffscreenWindow)(recv.native))
	var retGo (*gdkpixbuf.Pixbuf)
	if retC == nil {
		retGo = nil
	} else {
		retGo = gdkpixbuf.PixbufNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Retrieves a snapshot of the contained widget in the form of
// a #cairo_surface_t.  If you need to keep this around over window
// resizes then you should add a reference to it.
/*

C function : gtk_offscreen_window_get_surface
*/
func (recv *OffscreenWindow) GetSurface() *cairo.Surface {
	retC := C.gtk_offscreen_window_get_surface((*C.GtkOffscreenWindow)(recv.native))
	var retGo (*cairo.Surface)
	if retC == nil {
		retGo = nil
	} else {
		retGo = cairo.SurfaceNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Returns the #GdkWindow of the handle. This function is
// useful when handling button or motion events because it
// enables the callback to distinguish between the window
// of the paned, a child and the handle.
/*

C function : gtk_paned_get_handle_window
*/
func (recv *Paned) GetHandleWindow() *gdk.Window {
	retC := C.gtk_paned_get_handle_window((*C.GtkPaned)(recv.native))
	retGo := gdk.WindowNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Obtains the hardware printer margins of the #GtkPrintContext, in units.
/*

C function : gtk_print_context_get_hard_margins
*/
func (recv *PrintContext) GetHardMargins() (bool, float64, float64, float64, float64) {
	var c_top C.gdouble

	var c_bottom C.gdouble

	var c_left C.gdouble

	var c_right C.gdouble

	retC := C.gtk_print_context_get_hard_margins((*C.GtkPrintContext)(recv.native), &c_top, &c_bottom, &c_left, &c_right)
	retGo := retC == C.TRUE

	top := (float64)(c_top)

	bottom := (float64)(c_bottom)

	left := (float64)(c_left)

	right := (float64)(c_right)

	return retGo, top, bottom, left, right
}

// This function is useful mainly for #GtkRange subclasses.
//
// See gtk_range_set_min_slider_size().
/*

C function : gtk_range_get_min_slider_size
*/
func (recv *Range) GetMinSliderSize() int32 {
	retC := C.gtk_range_get_min_slider_size((*C.GtkRange)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : gtk_range_get_range_rect : unsupported parameter range_rect : Blacklisted record : GdkRectangle

// This function returns sliders range along the long dimension,
// in widget->window coordinates.
//
// This function is useful mainly for #GtkRange subclasses.
/*

C function : gtk_range_get_slider_range
*/
func (recv *Range) GetSliderRange() (int32, int32) {
	var c_slider_start C.gint

	var c_slider_end C.gint

	C.gtk_range_get_slider_range((*C.GtkRange)(recv.native), &c_slider_start, &c_slider_end)

	sliderStart := (int32)(c_slider_start)

	sliderEnd := (int32)(c_slider_end)

	return sliderStart, sliderEnd
}

// This function is useful mainly for #GtkRange subclasses.
//
// See gtk_range_set_slider_size_fixed().
/*

C function : gtk_range_get_slider_size_fixed
*/
func (recv *Range) GetSliderSizeFixed() bool {
	retC := C.gtk_range_get_slider_size_fixed((*C.GtkRange)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Sets the minimum size of the range’s slider.
//
// This function is useful mainly for #GtkRange subclasses.
/*

C function : gtk_range_set_min_slider_size
*/
func (recv *Range) SetMinSliderSize(minSize int32) {
	c_min_size := (C.gint)(minSize)

	C.gtk_range_set_min_slider_size((*C.GtkRange)(recv.native), c_min_size)

	return
}

// Sets whether the range’s slider has a fixed size, or a size that
// depends on its adjustment’s page size.
//
// This function is useful mainly for #GtkRange subclasses.
/*

C function : gtk_range_set_slider_size_fixed
*/
func (recv *Range) SetSliderSizeFixed(sizeFixed bool) {
	c_size_fixed :=
		boolToGboolean(sizeFixed)

	C.gtk_range_set_slider_size_fixed((*C.GtkRange)(recv.native), c_size_fixed)

	return
}

// Returns a new spinner widget. Not yet started.
/*

C function : gtk_spinner_new
*/
func SpinnerNew() *Spinner {
	retC := C.gtk_spinner_new()
	retGo := SpinnerNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Starts the animation of the spinner.
/*

C function : gtk_spinner_start
*/
func (recv *Spinner) Start() {
	C.gtk_spinner_start((*C.GtkSpinner)(recv.native))

	return
}

// Stops the animation of the spinner.
/*

C function : gtk_spinner_stop
*/
func (recv *Spinner) Stop() {
	C.gtk_spinner_stop((*C.GtkSpinner)(recv.native))

	return
}

// Sets the name of this tray icon.
// This should be a string identifying this icon. It is may be
// used for sorting the icons in the tray and will not be shown to
// the user.
/*

C function : gtk_status_icon_set_name
*/
func (recv *StatusIcon) SetName(name string) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	C.gtk_status_icon_set_name((*C.GtkStatusIcon)(recv.native), c_name)

	return
}

// Retrieves the box containing the label widget.
/*

C function : gtk_statusbar_get_message_area
*/
func (recv *Statusbar) GetMessageArea() *Box {
	retC := C.gtk_statusbar_get_message_area((*C.GtkStatusbar)(recv.native))
	retGo := BoxNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported signal 'preedit-changed' for TextView : unsupported parameter preedit : type utf8 :

// Returns the ellipsize mode used for @tool_item. Custom subclasses of
// #GtkToolItem should call this function to find out how text should
// be ellipsized.
/*

C function : gtk_tool_item_get_ellipsize_mode
*/
func (recv *ToolItem) GetEllipsizeMode() pango.EllipsizeMode {
	retC := C.gtk_tool_item_get_ellipsize_mode((*C.GtkToolItem)(recv.native))
	retGo := (pango.EllipsizeMode)(retC)

	return retGo
}

// Returns the text alignment used for @tool_item. Custom subclasses of
// #GtkToolItem should call this function to find out how text should
// be aligned.
/*

C function : gtk_tool_item_get_text_alignment
*/
func (recv *ToolItem) GetTextAlignment() float32 {
	retC := C.gtk_tool_item_get_text_alignment((*C.GtkToolItem)(recv.native))
	retGo := (float32)(retC)

	return retGo
}

// Returns the text orientation used for @tool_item. Custom subclasses of
// #GtkToolItem should call this function to find out how text should
// be orientated.
/*

C function : gtk_tool_item_get_text_orientation
*/
func (recv *ToolItem) GetTextOrientation() Orientation {
	retC := C.gtk_tool_item_get_text_orientation((*C.GtkToolItem)(recv.native))
	retGo := (Orientation)(retC)

	return retGo
}

// Returns the size group used for labels in @tool_item.
// Custom subclasses of #GtkToolItem should call this function
// and use the size group for labels.
/*

C function : gtk_tool_item_get_text_size_group
*/
func (recv *ToolItem) GetTextSizeGroup() *SizeGroup {
	retC := C.gtk_tool_item_get_text_size_group((*C.GtkToolItem)(recv.native))
	retGo := SizeGroupNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Creates a new tool item group with label @label.
/*

C function : gtk_tool_item_group_new
*/
func ToolItemGroupNew(label string) *ToolItemGroup {
	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	retC := C.gtk_tool_item_group_new(c_label)
	retGo := ToolItemGroupNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets whether @group is collapsed or expanded.
/*

C function : gtk_tool_item_group_get_collapsed
*/
func (recv *ToolItemGroup) GetCollapsed() bool {
	retC := C.gtk_tool_item_group_get_collapsed((*C.GtkToolItemGroup)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Gets the tool item at position (x, y).
/*

C function : gtk_tool_item_group_get_drop_item
*/
func (recv *ToolItemGroup) GetDropItem(x int32, y int32) *ToolItem {
	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	retC := C.gtk_tool_item_group_get_drop_item((*C.GtkToolItemGroup)(recv.native), c_x, c_y)
	retGo := ToolItemNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the ellipsization mode of @group.
/*

C function : gtk_tool_item_group_get_ellipsize
*/
func (recv *ToolItemGroup) GetEllipsize() pango.EllipsizeMode {
	retC := C.gtk_tool_item_group_get_ellipsize((*C.GtkToolItemGroup)(recv.native))
	retGo := (pango.EllipsizeMode)(retC)

	return retGo
}

// Gets the relief mode of the header button of @group.
/*

C function : gtk_tool_item_group_get_header_relief
*/
func (recv *ToolItemGroup) GetHeaderRelief() ReliefStyle {
	retC := C.gtk_tool_item_group_get_header_relief((*C.GtkToolItemGroup)(recv.native))
	retGo := (ReliefStyle)(retC)

	return retGo
}

// Gets the position of @item in @group as index.
/*

C function : gtk_tool_item_group_get_item_position
*/
func (recv *ToolItemGroup) GetItemPosition(item *ToolItem) int32 {
	c_item := (*C.GtkToolItem)(C.NULL)
	if item != nil {
		c_item = (*C.GtkToolItem)(item.ToC())
	}

	retC := C.gtk_tool_item_group_get_item_position((*C.GtkToolItemGroup)(recv.native), c_item)
	retGo := (int32)(retC)

	return retGo
}

// Gets the label of @group.
/*

C function : gtk_tool_item_group_get_label
*/
func (recv *ToolItemGroup) GetLabel() string {
	retC := C.gtk_tool_item_group_get_label((*C.GtkToolItemGroup)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Gets the label widget of @group.
// See gtk_tool_item_group_set_label_widget().
/*

C function : gtk_tool_item_group_get_label_widget
*/
func (recv *ToolItemGroup) GetLabelWidget() *Widget {
	retC := C.gtk_tool_item_group_get_label_widget((*C.GtkToolItemGroup)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the number of tool items in @group.
/*

C function : gtk_tool_item_group_get_n_items
*/
func (recv *ToolItemGroup) GetNItems() uint32 {
	retC := C.gtk_tool_item_group_get_n_items((*C.GtkToolItemGroup)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// Gets the tool item at @index in group.
/*

C function : gtk_tool_item_group_get_nth_item
*/
func (recv *ToolItemGroup) GetNthItem(index uint32) *ToolItem {
	c_index := (C.guint)(index)

	retC := C.gtk_tool_item_group_get_nth_item((*C.GtkToolItemGroup)(recv.native), c_index)
	retGo := ToolItemNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Inserts @item at @position in the list of children of @group.
/*

C function : gtk_tool_item_group_insert
*/
func (recv *ToolItemGroup) Insert(item *ToolItem, position int32) {
	c_item := (*C.GtkToolItem)(C.NULL)
	if item != nil {
		c_item = (*C.GtkToolItem)(item.ToC())
	}

	c_position := (C.gint)(position)

	C.gtk_tool_item_group_insert((*C.GtkToolItemGroup)(recv.native), c_item, c_position)

	return
}

// Sets whether the @group should be collapsed or expanded.
/*

C function : gtk_tool_item_group_set_collapsed
*/
func (recv *ToolItemGroup) SetCollapsed(collapsed bool) {
	c_collapsed :=
		boolToGboolean(collapsed)

	C.gtk_tool_item_group_set_collapsed((*C.GtkToolItemGroup)(recv.native), c_collapsed)

	return
}

// Sets the ellipsization mode which should be used by labels in @group.
/*

C function : gtk_tool_item_group_set_ellipsize
*/
func (recv *ToolItemGroup) SetEllipsize(ellipsize pango.EllipsizeMode) {
	c_ellipsize := (C.PangoEllipsizeMode)(ellipsize)

	C.gtk_tool_item_group_set_ellipsize((*C.GtkToolItemGroup)(recv.native), c_ellipsize)

	return
}

// Set the button relief of the group header.
// See gtk_button_set_relief() for details.
/*

C function : gtk_tool_item_group_set_header_relief
*/
func (recv *ToolItemGroup) SetHeaderRelief(style ReliefStyle) {
	c_style := (C.GtkReliefStyle)(style)

	C.gtk_tool_item_group_set_header_relief((*C.GtkToolItemGroup)(recv.native), c_style)

	return
}

// Sets the position of @item in the list of children of @group.
/*

C function : gtk_tool_item_group_set_item_position
*/
func (recv *ToolItemGroup) SetItemPosition(item *ToolItem, position int32) {
	c_item := (*C.GtkToolItem)(C.NULL)
	if item != nil {
		c_item = (*C.GtkToolItem)(item.ToC())
	}

	c_position := (C.gint)(position)

	C.gtk_tool_item_group_set_item_position((*C.GtkToolItemGroup)(recv.native), c_item, c_position)

	return
}

// Sets the label of the tool item group. The label is displayed in the header
// of the group.
/*

C function : gtk_tool_item_group_set_label
*/
func (recv *ToolItemGroup) SetLabel(label string) {
	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	C.gtk_tool_item_group_set_label((*C.GtkToolItemGroup)(recv.native), c_label)

	return
}

// Sets the label of the tool item group.
// The label widget is displayed in the header of the group, in place
// of the usual label.
/*

C function : gtk_tool_item_group_set_label_widget
*/
func (recv *ToolItemGroup) SetLabelWidget(labelWidget *Widget) {
	c_label_widget := (*C.GtkWidget)(C.NULL)
	if labelWidget != nil {
		c_label_widget = (*C.GtkWidget)(labelWidget.ToC())
	}

	C.gtk_tool_item_group_set_label_widget((*C.GtkToolItemGroup)(recv.native), c_label_widget)

	return
}

// Creates a new tool palette.
/*

C function : gtk_tool_palette_new
*/
func ToolPaletteNew() *ToolPalette {
	retC := C.gtk_tool_palette_new()
	retGo := ToolPaletteNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Sets @palette as drag source (see gtk_tool_palette_set_drag_source())
// and sets @widget as a drag destination for drags from @palette.
// See gtk_drag_dest_set().
/*

C function : gtk_tool_palette_add_drag_dest
*/
func (recv *ToolPalette) AddDragDest(widget *Widget, flags DestDefaults, targets ToolPaletteDragTargets, actions gdk.DragAction) {
	c_widget := (*C.GtkWidget)(C.NULL)
	if widget != nil {
		c_widget = (*C.GtkWidget)(widget.ToC())
	}

	c_flags := (C.GtkDestDefaults)(flags)

	c_targets := (C.GtkToolPaletteDragTargets)(targets)

	c_actions := (C.GdkDragAction)(actions)

	C.gtk_tool_palette_add_drag_dest((*C.GtkToolPalette)(recv.native), c_widget, c_flags, c_targets, c_actions)

	return
}

// Get the dragged item from the selection.
// This could be a #GtkToolItem or a #GtkToolItemGroup.
/*

C function : gtk_tool_palette_get_drag_item
*/
func (recv *ToolPalette) GetDragItem(selection *SelectionData) *Widget {
	c_selection := (*C.GtkSelectionData)(C.NULL)
	if selection != nil {
		c_selection = (*C.GtkSelectionData)(selection.ToC())
	}

	retC := C.gtk_tool_palette_get_drag_item((*C.GtkToolPalette)(recv.native), c_selection)
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the group at position (x, y).
/*

C function : gtk_tool_palette_get_drop_group
*/
func (recv *ToolPalette) GetDropGroup(x int32, y int32) *ToolItemGroup {
	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	retC := C.gtk_tool_palette_get_drop_group((*C.GtkToolPalette)(recv.native), c_x, c_y)
	var retGo (*ToolItemGroup)
	if retC == nil {
		retGo = nil
	} else {
		retGo = ToolItemGroupNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Gets the item at position (x, y).
// See gtk_tool_palette_get_drop_group().
/*

C function : gtk_tool_palette_get_drop_item
*/
func (recv *ToolPalette) GetDropItem(x int32, y int32) *ToolItem {
	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	retC := C.gtk_tool_palette_get_drop_item((*C.GtkToolPalette)(recv.native), c_x, c_y)
	var retGo (*ToolItem)
	if retC == nil {
		retGo = nil
	} else {
		retGo = ToolItemNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Gets whether @group is exclusive or not.
// See gtk_tool_palette_set_exclusive().
/*

C function : gtk_tool_palette_get_exclusive
*/
func (recv *ToolPalette) GetExclusive(group *ToolItemGroup) bool {
	c_group := (*C.GtkToolItemGroup)(C.NULL)
	if group != nil {
		c_group = (*C.GtkToolItemGroup)(group.ToC())
	}

	retC := C.gtk_tool_palette_get_exclusive((*C.GtkToolPalette)(recv.native), c_group)
	retGo := retC == C.TRUE

	return retGo
}

// Gets whether group should be given extra space.
// See gtk_tool_palette_set_expand().
/*

C function : gtk_tool_palette_get_expand
*/
func (recv *ToolPalette) GetExpand(group *ToolItemGroup) bool {
	c_group := (*C.GtkToolItemGroup)(C.NULL)
	if group != nil {
		c_group = (*C.GtkToolItemGroup)(group.ToC())
	}

	retC := C.gtk_tool_palette_get_expand((*C.GtkToolPalette)(recv.native), c_group)
	retGo := retC == C.TRUE

	return retGo
}

// Gets the position of @group in @palette as index.
// See gtk_tool_palette_set_group_position().
/*

C function : gtk_tool_palette_get_group_position
*/
func (recv *ToolPalette) GetGroupPosition(group *ToolItemGroup) int32 {
	c_group := (*C.GtkToolItemGroup)(C.NULL)
	if group != nil {
		c_group = (*C.GtkToolItemGroup)(group.ToC())
	}

	retC := C.gtk_tool_palette_get_group_position((*C.GtkToolPalette)(recv.native), c_group)
	retGo := (int32)(retC)

	return retGo
}

// Gets the horizontal adjustment of the tool palette.
/*

C function : gtk_tool_palette_get_hadjustment
*/
func (recv *ToolPalette) GetHadjustment() *Adjustment {
	retC := C.gtk_tool_palette_get_hadjustment((*C.GtkToolPalette)(recv.native))
	retGo := AdjustmentNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the size of icons in the tool palette.
// See gtk_tool_palette_set_icon_size().
/*

C function : gtk_tool_palette_get_icon_size
*/
func (recv *ToolPalette) GetIconSize() IconSize {
	retC := C.gtk_tool_palette_get_icon_size((*C.GtkToolPalette)(recv.native))
	retGo := (IconSize)(retC)

	return retGo
}

// Gets the style (icons, text or both) of items in the tool palette.
/*

C function : gtk_tool_palette_get_style
*/
func (recv *ToolPalette) GetStyle() ToolbarStyle {
	retC := C.gtk_tool_palette_get_style((*C.GtkToolPalette)(recv.native))
	retGo := (ToolbarStyle)(retC)

	return retGo
}

// Gets the vertical adjustment of the tool palette.
/*

C function : gtk_tool_palette_get_vadjustment
*/
func (recv *ToolPalette) GetVadjustment() *Adjustment {
	retC := C.gtk_tool_palette_get_vadjustment((*C.GtkToolPalette)(recv.native))
	retGo := AdjustmentNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Sets the tool palette as a drag source.
// Enables all groups and items in the tool palette as drag sources
// on button 1 and button 3 press with copy and move actions.
// See gtk_drag_source_set().
/*

C function : gtk_tool_palette_set_drag_source
*/
func (recv *ToolPalette) SetDragSource(targets ToolPaletteDragTargets) {
	c_targets := (C.GtkToolPaletteDragTargets)(targets)

	C.gtk_tool_palette_set_drag_source((*C.GtkToolPalette)(recv.native), c_targets)

	return
}

// Sets whether the group should be exclusive or not.
// If an exclusive group is expanded all other groups are collapsed.
/*

C function : gtk_tool_palette_set_exclusive
*/
func (recv *ToolPalette) SetExclusive(group *ToolItemGroup, exclusive bool) {
	c_group := (*C.GtkToolItemGroup)(C.NULL)
	if group != nil {
		c_group = (*C.GtkToolItemGroup)(group.ToC())
	}

	c_exclusive :=
		boolToGboolean(exclusive)

	C.gtk_tool_palette_set_exclusive((*C.GtkToolPalette)(recv.native), c_group, c_exclusive)

	return
}

// Sets whether the group should be given extra space.
/*

C function : gtk_tool_palette_set_expand
*/
func (recv *ToolPalette) SetExpand(group *ToolItemGroup, expand bool) {
	c_group := (*C.GtkToolItemGroup)(C.NULL)
	if group != nil {
		c_group = (*C.GtkToolItemGroup)(group.ToC())
	}

	c_expand :=
		boolToGboolean(expand)

	C.gtk_tool_palette_set_expand((*C.GtkToolPalette)(recv.native), c_group, c_expand)

	return
}

// Sets the position of the group as an index of the tool palette.
// If position is 0 the group will become the first child, if position is
// -1 it will become the last child.
/*

C function : gtk_tool_palette_set_group_position
*/
func (recv *ToolPalette) SetGroupPosition(group *ToolItemGroup, position int32) {
	c_group := (*C.GtkToolItemGroup)(C.NULL)
	if group != nil {
		c_group = (*C.GtkToolItemGroup)(group.ToC())
	}

	c_position := (C.gint)(position)

	C.gtk_tool_palette_set_group_position((*C.GtkToolPalette)(recv.native), c_group, c_position)

	return
}

// Sets the size of icons in the tool palette.
/*

C function : gtk_tool_palette_set_icon_size
*/
func (recv *ToolPalette) SetIconSize(iconSize IconSize) {
	c_icon_size := (C.GtkIconSize)(iconSize)

	C.gtk_tool_palette_set_icon_size((*C.GtkToolPalette)(recv.native), c_icon_size)

	return
}

// Sets the style (text, icons or both) of items in the tool palette.
/*

C function : gtk_tool_palette_set_style
*/
func (recv *ToolPalette) SetStyle(style ToolbarStyle) {
	c_style := (C.GtkToolbarStyle)(style)

	C.gtk_tool_palette_set_style((*C.GtkToolPalette)(recv.native), c_style)

	return
}

// Unsets the tool palette icon size set with gtk_tool_palette_set_icon_size(),
// so that user preferences will be used to determine the icon size.
/*

C function : gtk_tool_palette_unset_icon_size
*/
func (recv *ToolPalette) UnsetIconSize() {
	C.gtk_tool_palette_unset_icon_size((*C.GtkToolPalette)(recv.native))

	return
}

// Unsets a toolbar style set with gtk_tool_palette_set_style(),
// so that user preferences will be used to determine the toolbar style.
/*

C function : gtk_tool_palette_unset_style
*/
func (recv *ToolPalette) UnsetStyle() {
	C.gtk_tool_palette_unset_style((*C.GtkToolPalette)(recv.native))

	return
}

// Sets the icon of the tooltip (which is in front of the text)
// to be the icon indicated by @gicon with the size indicated
// by @size. If @gicon is %NULL, the image will be hidden.
/*

C function : gtk_tooltip_set_icon_from_gicon
*/
func (recv *Tooltip) SetIconFromGicon(gicon *gio.Icon, size IconSize) {
	c_gicon := (*C.GIcon)(gicon.ToC())

	c_size := (C.GtkIconSize)(size)

	C.gtk_tooltip_set_icon_from_gicon((*C.GtkTooltip)(recv.native), c_gicon, c_size)

	return
}

// Gets the bin window of the #GtkViewport.
/*

C function : gtk_viewport_get_bin_window
*/
func (recv *Viewport) GetBinWindow() *gdk.Window {
	retC := C.gtk_viewport_get_bin_window((*C.GtkViewport)(recv.native))
	retGo := gdk.WindowNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Whether the widget is mapped.
/*

C function : gtk_widget_get_mapped
*/
func (recv *Widget) GetMapped() bool {
	retC := C.gtk_widget_get_mapped((*C.GtkWidget)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Determines whether @widget is realized.
/*

C function : gtk_widget_get_realized
*/
func (recv *Widget) GetRealized() bool {
	retC := C.gtk_widget_get_realized((*C.GtkWidget)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Retrieves the widget’s requisition.
//
// This function should only be used by widget implementations in
// order to figure whether the widget’s requisition has actually
// changed after some internal state change (so that they can call
// gtk_widget_queue_resize() instead of gtk_widget_queue_draw()).
//
// Normally, gtk_widget_size_request() should be used.
/*

C function : gtk_widget_get_requisition
*/
func (recv *Widget) GetRequisition() *Requisition {
	var c_requisition C.GtkRequisition

	C.gtk_widget_get_requisition((*C.GtkWidget)(recv.native), &c_requisition)

	requisition := RequisitionNewFromC(unsafe.Pointer(&c_requisition))

	return requisition
}

// Determines if the widget style has been looked up through the rc mechanism.
/*

C function : gtk_widget_has_rc_style
*/
func (recv *Widget) HasRcStyle() bool {
	retC := C.gtk_widget_has_rc_style((*C.GtkWidget)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_widget_send_focus_change : unsupported parameter event : no type generator for Gdk.Event (GdkEvent*) for param event

// Marks the widget as being mapped.
//
// This function should only ever be called in a derived widget's
// “map” or “unmap” implementation.
/*

C function : gtk_widget_set_mapped
*/
func (recv *Widget) SetMapped(mapped bool) {
	c_mapped :=
		boolToGboolean(mapped)

	C.gtk_widget_set_mapped((*C.GtkWidget)(recv.native), c_mapped)

	return
}

// Marks the widget as being realized. This function must only be
// called after all #GdkWindows for the @widget have been created
// and registered.
//
// This function should only ever be called in a derived widget's
// “realize” or “unrealize” implementation.
/*

C function : gtk_widget_set_realized
*/
func (recv *Widget) SetRealized(realized bool) {
	c_realized :=
		boolToGboolean(realized)

	C.gtk_widget_set_realized((*C.GtkWidget)(recv.native), c_realized)

	return
}

// This function attaches the widget’s #GtkStyle to the widget's
// #GdkWindow. It is a replacement for
//
// |[
// widget->style = gtk_style_attach (widget->style, widget->window);
// ]|
//
// and should only ever be called in a derived widget’s “realize”
// implementation which does not chain up to its parent class'
// “realize” implementation, because one of the parent classes
// (finally #GtkWidget) would attach the style itself.
/*

C function : gtk_widget_style_attach
*/
func (recv *Widget) StyleAttach() {
	C.gtk_widget_style_attach((*C.GtkWidget)(recv.native))

	return
}

// Gets the value of the #GtkWindow:mnemonics-visible property.
/*

C function : gtk_window_get_mnemonics_visible
*/
func (recv *Window) GetMnemonicsVisible() bool {
	retC := C.gtk_window_get_mnemonics_visible((*C.GtkWindow)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Gets the type of the window. See #GtkWindowType.
/*

C function : gtk_window_get_window_type
*/
func (recv *Window) GetWindowType() WindowType {
	retC := C.gtk_window_get_window_type((*C.GtkWindow)(recv.native))
	retGo := (WindowType)(retC)

	return retGo
}

// Sets the #GtkWindow:mnemonics-visible property.
/*

C function : gtk_window_set_mnemonics_visible
*/
func (recv *Window) SetMnemonicsVisible(setting bool) {
	c_setting :=
		boolToGboolean(setting)

	C.gtk_window_set_mnemonics_visible((*C.GtkWindow)(recv.native), c_setting)

	return
}
