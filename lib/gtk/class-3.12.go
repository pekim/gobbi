// This is a generated file - DO NOT EDIT
// +build gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	gio "github.com/pekim/gobbi/lib/gio"
	glib "github.com/pekim/gobbi/lib/glib"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

// Unsupported : gtk_accel_label_get_accel : unsupported parameter accelerator_mods : GdkModifierType* with indirection level of 1

// Creates a new #GtkActionBar widget.
/*

C function

gtk_action_bar_new
*/
func ActionBarNew() *ActionBar {
	retC := C.gtk_action_bar_new()
	retGo := ActionBarNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Retrieves the center bar widget of the bar.
/*

C function

gtk_action_bar_get_center_widget
*/
func (recv *ActionBar) GetCenterWidget() *Widget {
	retC := C.gtk_action_bar_get_center_widget((*C.GtkActionBar)(recv.native))
	var retGo (*Widget)
	if retC == nil {
		retGo = nil
	} else {
		retGo = WidgetNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Adds @child to @action_bar, packed with reference to the
// end of the @action_bar.
/*

C function

gtk_action_bar_pack_end
*/
func (recv *ActionBar) PackEnd(child *Widget) {
	c_child := (*C.GtkWidget)(C.NULL)
	if child != nil {
		c_child = (*C.GtkWidget)(child.ToC())
	}

	C.gtk_action_bar_pack_end((*C.GtkActionBar)(recv.native), c_child)

	return
}

// Adds @child to @action_bar, packed with reference to the
// start of the @action_bar.
/*

C function

gtk_action_bar_pack_start
*/
func (recv *ActionBar) PackStart(child *Widget) {
	c_child := (*C.GtkWidget)(C.NULL)
	if child != nil {
		c_child = (*C.GtkWidget)(child.ToC())
	}

	C.gtk_action_bar_pack_start((*C.GtkActionBar)(recv.native), c_child)

	return
}

// Sets the center widget for the #GtkActionBar.
/*

C function

gtk_action_bar_set_center_widget
*/
func (recv *ActionBar) SetCenterWidget(centerWidget *Widget) {
	c_center_widget := (*C.GtkWidget)(C.NULL)
	if centerWidget != nil {
		c_center_widget = (*C.GtkWidget)(centerWidget.ToC())
	}

	C.gtk_action_bar_set_center_widget((*C.GtkActionBar)(recv.native), c_center_widget)

	return
}

// Unsupported : gtk_application_get_accels_for_action : no return type

// Unsupported : gtk_application_list_action_descriptions : no return type

// Unsupported : gtk_application_set_accels_for_action : unsupported parameter accels :

// Retrieves the center widget of the box.
/*

C function

gtk_box_get_center_widget
*/
func (recv *Box) GetCenterWidget() *Widget {
	retC := C.gtk_box_get_center_widget((*C.GtkBox)(recv.native))
	var retGo (*Widget)
	if retC == nil {
		retGo = nil
	} else {
		retGo = WidgetNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Sets a center widget; that is a child widget that will be
// centered with respect to the full width of the box, even
// if the children at either side take up different amounts
// of space.
/*

C function

gtk_box_set_center_widget
*/
func (recv *Box) SetCenterWidget(widget *Widget) {
	c_widget := (*C.GtkWidget)(C.NULL)
	if widget != nil {
		c_widget = (*C.GtkWidget)(widget.ToC())
	}

	C.gtk_box_set_center_widget((*C.GtkBox)(recv.native), c_widget)

	return
}

// Returns the header bar of @dialog. Note that the
// headerbar is only used by the dialog if the
// #GtkDialog:use-header-bar property is %TRUE.
/*

C function

gtk_dialog_get_header_bar
*/
func (recv *Dialog) GetHeaderBar() *Widget {
	retC := C.gtk_dialog_get_header_bar((*C.GtkDialog)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Retrieves the desired maximum width of @entry, in characters.
// See gtk_entry_set_max_width_chars().
/*

C function

gtk_entry_get_max_width_chars
*/
func (recv *Entry) GetMaxWidthChars() int32 {
	retC := C.gtk_entry_get_max_width_chars((*C.GtkEntry)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Sets the desired maximum width in characters of @entry.
/*

C function

gtk_entry_set_max_width_chars
*/
func (recv *Entry) SetMaxWidthChars(nChars int32) {
	c_n_chars := (C.gint)(nChars)

	C.gtk_entry_set_max_width_chars((*C.GtkEntry)(recv.native), c_n_chars)

	return
}

// Creates a GtkFlowBox.
/*

C function

gtk_flow_box_new
*/
func FlowBoxNew() *FlowBox {
	retC := C.gtk_flow_box_new()
	retGo := FlowBoxNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Returns whether children activate on single clicks.
/*

C function

gtk_flow_box_get_activate_on_single_click
*/
func (recv *FlowBox) GetActivateOnSingleClick() bool {
	retC := C.gtk_flow_box_get_activate_on_single_click((*C.GtkFlowBox)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Gets the nth child in the @box.
/*

C function

gtk_flow_box_get_child_at_index
*/
func (recv *FlowBox) GetChildAtIndex(idx int32) *FlowBoxChild {
	c_idx := (C.gint)(idx)

	retC := C.gtk_flow_box_get_child_at_index((*C.GtkFlowBox)(recv.native), c_idx)
	var retGo (*FlowBoxChild)
	if retC == nil {
		retGo = nil
	} else {
		retGo = FlowBoxChildNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Gets the horizontal spacing.
/*

C function

gtk_flow_box_get_column_spacing
*/
func (recv *FlowBox) GetColumnSpacing() uint32 {
	retC := C.gtk_flow_box_get_column_spacing((*C.GtkFlowBox)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// Returns whether the box is homogeneous (all children are the
// same size). See gtk_box_set_homogeneous().
/*

C function

gtk_flow_box_get_homogeneous
*/
func (recv *FlowBox) GetHomogeneous() bool {
	retC := C.gtk_flow_box_get_homogeneous((*C.GtkFlowBox)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Gets the maximum number of children per line.
/*

C function

gtk_flow_box_get_max_children_per_line
*/
func (recv *FlowBox) GetMaxChildrenPerLine() uint32 {
	retC := C.gtk_flow_box_get_max_children_per_line((*C.GtkFlowBox)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// Gets the minimum number of children per line.
/*

C function

gtk_flow_box_get_min_children_per_line
*/
func (recv *FlowBox) GetMinChildrenPerLine() uint32 {
	retC := C.gtk_flow_box_get_min_children_per_line((*C.GtkFlowBox)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// Gets the vertical spacing.
/*

C function

gtk_flow_box_get_row_spacing
*/
func (recv *FlowBox) GetRowSpacing() uint32 {
	retC := C.gtk_flow_box_get_row_spacing((*C.GtkFlowBox)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// Creates a list of all selected children.
/*

C function

gtk_flow_box_get_selected_children
*/
func (recv *FlowBox) GetSelectedChildren() *glib.List {
	retC := C.gtk_flow_box_get_selected_children((*C.GtkFlowBox)(recv.native))
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the selection mode of @box.
/*

C function

gtk_flow_box_get_selection_mode
*/
func (recv *FlowBox) GetSelectionMode() SelectionMode {
	retC := C.gtk_flow_box_get_selection_mode((*C.GtkFlowBox)(recv.native))
	retGo := (SelectionMode)(retC)

	return retGo
}

// Inserts the @widget into @box at @position.
//
// If a sort function is set, the widget will actually be inserted
// at the calculated position and this function has the same effect
// as gtk_container_add().
//
// If @position is -1, or larger than the total number of children
// in the @box, then the @widget will be appended to the end.
/*

C function

gtk_flow_box_insert
*/
func (recv *FlowBox) Insert(widget *Widget, position int32) {
	c_widget := (*C.GtkWidget)(C.NULL)
	if widget != nil {
		c_widget = (*C.GtkWidget)(widget.ToC())
	}

	c_position := (C.gint)(position)

	C.gtk_flow_box_insert((*C.GtkFlowBox)(recv.native), c_widget, c_position)

	return
}

// Updates the filtering for all children.
//
// Call this function when the result of the filter
// function on the @box is changed due ot an external
// factor. For instance, this would be used if the
// filter function just looked for a specific search
// term, and the entry with the string has changed.
/*

C function

gtk_flow_box_invalidate_filter
*/
func (recv *FlowBox) InvalidateFilter() {
	C.gtk_flow_box_invalidate_filter((*C.GtkFlowBox)(recv.native))

	return
}

// Updates the sorting for all children.
//
// Call this when the result of the sort function on
// @box is changed due to an external factor.
/*

C function

gtk_flow_box_invalidate_sort
*/
func (recv *FlowBox) InvalidateSort() {
	C.gtk_flow_box_invalidate_sort((*C.GtkFlowBox)(recv.native))

	return
}

// Select all children of @box, if the selection
// mode allows it.
/*

C function

gtk_flow_box_select_all
*/
func (recv *FlowBox) SelectAll() {
	C.gtk_flow_box_select_all((*C.GtkFlowBox)(recv.native))

	return
}

// Selects a single child of @box, if the selection
// mode allows it.
/*

C function

gtk_flow_box_select_child
*/
func (recv *FlowBox) SelectChild(child *FlowBoxChild) {
	c_child := (*C.GtkFlowBoxChild)(C.NULL)
	if child != nil {
		c_child = (*C.GtkFlowBoxChild)(child.ToC())
	}

	C.gtk_flow_box_select_child((*C.GtkFlowBox)(recv.native), c_child)

	return
}

// Unsupported : gtk_flow_box_selected_foreach : unsupported parameter func : no type generator for FlowBoxForeachFunc (GtkFlowBoxForeachFunc) for param func

// If @single is %TRUE, children will be activated when you click
// on them, otherwise you need to double-click.
/*

C function

gtk_flow_box_set_activate_on_single_click
*/
func (recv *FlowBox) SetActivateOnSingleClick(single bool) {
	c_single :=
		boolToGboolean(single)

	C.gtk_flow_box_set_activate_on_single_click((*C.GtkFlowBox)(recv.native), c_single)

	return
}

// Sets the horizontal space to add between children.
// See the #GtkFlowBox:column-spacing property.
/*

C function

gtk_flow_box_set_column_spacing
*/
func (recv *FlowBox) SetColumnSpacing(spacing uint32) {
	c_spacing := (C.guint)(spacing)

	C.gtk_flow_box_set_column_spacing((*C.GtkFlowBox)(recv.native), c_spacing)

	return
}

// Unsupported : gtk_flow_box_set_filter_func : unsupported parameter filter_func : no type generator for FlowBoxFilterFunc (GtkFlowBoxFilterFunc) for param filter_func

// Hooks up an adjustment to focus handling in @box.
// The adjustment is also used for autoscrolling during
// rubberband selection. See gtk_scrolled_window_get_hadjustment()
// for a typical way of obtaining the adjustment, and
// gtk_flow_box_set_vadjustment()for setting the vertical
// adjustment.
//
// The adjustments have to be in pixel units and in the same
// coordinate system as the allocation for immediate children
// of the box.
/*

C function

gtk_flow_box_set_hadjustment
*/
func (recv *FlowBox) SetHadjustment(adjustment *Adjustment) {
	c_adjustment := (*C.GtkAdjustment)(C.NULL)
	if adjustment != nil {
		c_adjustment = (*C.GtkAdjustment)(adjustment.ToC())
	}

	C.gtk_flow_box_set_hadjustment((*C.GtkFlowBox)(recv.native), c_adjustment)

	return
}

// Sets the #GtkFlowBox:homogeneous property of @box, controlling
// whether or not all children of @box are given equal space
// in the box.
/*

C function

gtk_flow_box_set_homogeneous
*/
func (recv *FlowBox) SetHomogeneous(homogeneous bool) {
	c_homogeneous :=
		boolToGboolean(homogeneous)

	C.gtk_flow_box_set_homogeneous((*C.GtkFlowBox)(recv.native), c_homogeneous)

	return
}

// Sets the maximum number of children to request and
// allocate space for in @box’s orientation.
//
// Setting the maximum number of children per line
// limits the overall natural size request to be no more
// than @n_children children long in the given orientation.
/*

C function

gtk_flow_box_set_max_children_per_line
*/
func (recv *FlowBox) SetMaxChildrenPerLine(nChildren uint32) {
	c_n_children := (C.guint)(nChildren)

	C.gtk_flow_box_set_max_children_per_line((*C.GtkFlowBox)(recv.native), c_n_children)

	return
}

// Sets the minimum number of children to line up
// in @box’s orientation before flowing.
/*

C function

gtk_flow_box_set_min_children_per_line
*/
func (recv *FlowBox) SetMinChildrenPerLine(nChildren uint32) {
	c_n_children := (C.guint)(nChildren)

	C.gtk_flow_box_set_min_children_per_line((*C.GtkFlowBox)(recv.native), c_n_children)

	return
}

// Sets the vertical space to add between children.
// See the #GtkFlowBox:row-spacing property.
/*

C function

gtk_flow_box_set_row_spacing
*/
func (recv *FlowBox) SetRowSpacing(spacing uint32) {
	c_spacing := (C.guint)(spacing)

	C.gtk_flow_box_set_row_spacing((*C.GtkFlowBox)(recv.native), c_spacing)

	return
}

// Sets how selection works in @box.
// See #GtkSelectionMode for details.
/*

C function

gtk_flow_box_set_selection_mode
*/
func (recv *FlowBox) SetSelectionMode(mode SelectionMode) {
	c_mode := (C.GtkSelectionMode)(mode)

	C.gtk_flow_box_set_selection_mode((*C.GtkFlowBox)(recv.native), c_mode)

	return
}

// Unsupported : gtk_flow_box_set_sort_func : unsupported parameter sort_func : no type generator for FlowBoxSortFunc (GtkFlowBoxSortFunc) for param sort_func

// Hooks up an adjustment to focus handling in @box.
// The adjustment is also used for autoscrolling during
// rubberband selection. See gtk_scrolled_window_get_vadjustment()
// for a typical way of obtaining the adjustment, and
// gtk_flow_box_set_hadjustment()for setting the horizontal
// adjustment.
//
// The adjustments have to be in pixel units and in the same
// coordinate system as the allocation for immediate children
// of the box.
/*

C function

gtk_flow_box_set_vadjustment
*/
func (recv *FlowBox) SetVadjustment(adjustment *Adjustment) {
	c_adjustment := (*C.GtkAdjustment)(C.NULL)
	if adjustment != nil {
		c_adjustment = (*C.GtkAdjustment)(adjustment.ToC())
	}

	C.gtk_flow_box_set_vadjustment((*C.GtkFlowBox)(recv.native), c_adjustment)

	return
}

// Unselect all children of @box, if the selection
// mode allows it.
/*

C function

gtk_flow_box_unselect_all
*/
func (recv *FlowBox) UnselectAll() {
	C.gtk_flow_box_unselect_all((*C.GtkFlowBox)(recv.native))

	return
}

// Unselects a single child of @box, if the selection
// mode allows it.
/*

C function

gtk_flow_box_unselect_child
*/
func (recv *FlowBox) UnselectChild(child *FlowBoxChild) {
	c_child := (*C.GtkFlowBoxChild)(C.NULL)
	if child != nil {
		c_child = (*C.GtkFlowBoxChild)(child.ToC())
	}

	C.gtk_flow_box_unselect_child((*C.GtkFlowBox)(recv.native), c_child)

	return
}

// Creates a new #GtkFlowBoxChild, to be used as a child
// of a #GtkFlowBox.
/*

C function

gtk_flow_box_child_new
*/
func FlowBoxChildNew() *FlowBoxChild {
	retC := C.gtk_flow_box_child_new()
	retGo := FlowBoxChildNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Marks @child as changed, causing any state that depends on this
// to be updated. This affects sorting and filtering.
//
// Note that calls to this method must be in sync with the data
// used for the sorting and filtering functions. For instance, if
// the list is mirroring some external data set, and *two* children
// changed in the external data set when you call
// gtk_flow_box_child_changed() on the first child, the sort function
// must only read the new data for the first of the two changed
// children, otherwise the resorting of the children will be wrong.
//
// This generally means that if you don’t fully control the data
// model, you have to duplicate the data that affects the sorting
// and filtering functions into the widgets themselves. Another
// alternative is to call gtk_flow_box_invalidate_sort() on any
// model change, but that is more expensive.
/*

C function

gtk_flow_box_child_changed
*/
func (recv *FlowBoxChild) Changed() {
	C.gtk_flow_box_child_changed((*C.GtkFlowBoxChild)(recv.native))

	return
}

// Gets the current index of the @child in its #GtkFlowBox container.
/*

C function

gtk_flow_box_child_get_index
*/
func (recv *FlowBoxChild) GetIndex() int32 {
	retC := C.gtk_flow_box_child_get_index((*C.GtkFlowBoxChild)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Returns whether the @child is currently selected in its
// #GtkFlowBox container.
/*

C function

gtk_flow_box_child_is_selected
*/
func (recv *FlowBoxChild) IsSelected() bool {
	retC := C.gtk_flow_box_child_is_selected((*C.GtkFlowBoxChild)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Gets the decoration layout set with
// gtk_header_bar_set_decoration_layout().
/*

C function

gtk_header_bar_get_decoration_layout
*/
func (recv *HeaderBar) GetDecorationLayout() string {
	retC := C.gtk_header_bar_get_decoration_layout((*C.GtkHeaderBar)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Retrieves whether the header bar reserves space for
// a subtitle, regardless if one is currently set or not.
/*

C function

gtk_header_bar_get_has_subtitle
*/
func (recv *HeaderBar) GetHasSubtitle() bool {
	retC := C.gtk_header_bar_get_has_subtitle((*C.GtkHeaderBar)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Sets the decoration layout for this header bar, overriding
// the #GtkSettings:gtk-decoration-layout setting.
//
// There can be valid reasons for overriding the setting, such
// as a header bar design that does not allow for buttons to take
// room on the right, or only offers room for a single close button.
// Split header bars are another example for overriding the
// setting.
//
// The format of the string is button names, separated by commas.
// A colon separates the buttons that should appear on the left
// from those on the right. Recognized button names are minimize,
// maximize, close, icon (the window icon) and menu (a menu button
// for the fallback app menu).
//
// For example, “menu:minimize,maximize,close” specifies a menu
// on the left, and minimize, maximize and close buttons on the right.
/*

C function

gtk_header_bar_set_decoration_layout
*/
func (recv *HeaderBar) SetDecorationLayout(layout string) {
	c_layout := C.CString(layout)
	defer C.free(unsafe.Pointer(c_layout))

	C.gtk_header_bar_set_decoration_layout((*C.GtkHeaderBar)(recv.native), c_layout)

	return
}

// Sets whether the header bar should reserve space
// for a subtitle, even if none is currently set.
/*

C function

gtk_header_bar_set_has_subtitle
*/
func (recv *HeaderBar) SetHasSubtitle(setting bool) {
	c_setting :=
		boolToGboolean(setting)

	C.gtk_header_bar_set_has_subtitle((*C.GtkHeaderBar)(recv.native), c_setting)

	return
}

// Checks if the icon is symbolic or not. This currently uses only
// the file name and not the file contents for determining this.
// This behaviour may change in the future.
/*

C function

gtk_icon_info_is_symbolic
*/
func (recv *IconInfo) IsSymbolic() bool {
	retC := C.gtk_icon_info_is_symbolic((*C.GtkIconInfo)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Returns the #GtkPopover that pops out of the button.
// If the button is not using a #GtkPopover, this function
// returns %NULL.
/*

C function

gtk_menu_button_get_popover
*/
func (recv *MenuButton) GetPopover() *Popover {
	retC := C.gtk_menu_button_get_popover((*C.GtkMenuButton)(recv.native))
	var retGo (*Popover)
	if retC == nil {
		retGo = nil
	} else {
		retGo = PopoverNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Returns whether a #GtkPopover or a #GtkMenu will be constructed
// from the menu model.
/*

C function

gtk_menu_button_get_use_popover
*/
func (recv *MenuButton) GetUsePopover() bool {
	retC := C.gtk_menu_button_get_use_popover((*C.GtkMenuButton)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Sets the #GtkPopover that will be popped up when the button is
// clicked, or %NULL to disable the button. If #GtkMenuButton:menu-model
// or #GtkMenuButton:popup are set, they will be set to %NULL.
/*

C function

gtk_menu_button_set_popover
*/
func (recv *MenuButton) SetPopover(popover *Widget) {
	c_popover := (*C.GtkWidget)(C.NULL)
	if popover != nil {
		c_popover = (*C.GtkWidget)(popover.ToC())
	}

	C.gtk_menu_button_set_popover((*C.GtkMenuButton)(recv.native), c_popover)

	return
}

// Sets whether to construct a #GtkPopover instead of #GtkMenu
// when gtk_menu_button_set_menu_model() is called. Note that
// this property is only consulted when a new menu model is set.
/*

C function

gtk_menu_button_set_use_popover
*/
func (recv *MenuButton) SetUsePopover(usePopover bool) {
	c_use_popover :=
		boolToGboolean(usePopover)

	C.gtk_menu_button_set_use_popover((*C.GtkMenuButton)(recv.native), c_use_popover)

	return
}

// Returns the value previously set with gtk_places_sidebar_set_local_only().
/*

C function

gtk_places_sidebar_get_local_only
*/
func (recv *PlacesSidebar) GetLocalOnly() bool {
	retC := C.gtk_places_sidebar_get_local_only((*C.GtkPlacesSidebar)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Sets whether the @sidebar should only show local files.
/*

C function

gtk_places_sidebar_set_local_only
*/
func (recv *PlacesSidebar) SetLocalOnly(localOnly bool) {
	c_local_only :=
		boolToGboolean(localOnly)

	C.gtk_places_sidebar_set_local_only((*C.GtkPlacesSidebar)(recv.native), c_local_only)

	return
}

// Creates a new popover to point to @relative_to
/*

C function

gtk_popover_new
*/
func PopoverNew(relativeTo *Widget) *Popover {
	c_relative_to := (*C.GtkWidget)(C.NULL)
	if relativeTo != nil {
		c_relative_to = (*C.GtkWidget)(relativeTo.ToC())
	}

	retC := C.gtk_popover_new(c_relative_to)
	retGo := PopoverNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Creates a #GtkPopover and populates it according to
// @model. The popover is pointed to the @relative_to widget.
//
// The created buttons are connected to actions found in the
// #GtkApplicationWindow to which the popover belongs - typically
// by means of being attached to a widget that is contained within
// the #GtkApplicationWindows widget hierarchy.
//
// Actions can also be added using gtk_widget_insert_action_group()
// on the menus attach widget or on any of its parent widgets.
/*

C function

gtk_popover_new_from_model
*/
func PopoverNewFromModel(relativeTo *Widget, model *gio.MenuModel) *Popover {
	c_relative_to := (*C.GtkWidget)(C.NULL)
	if relativeTo != nil {
		c_relative_to = (*C.GtkWidget)(relativeTo.ToC())
	}

	c_model := (*C.GMenuModel)(C.NULL)
	if model != nil {
		c_model = (*C.GMenuModel)(model.ToC())
	}

	retC := C.gtk_popover_new_from_model(c_relative_to, c_model)
	retGo := PopoverNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Establishes a binding between a #GtkPopover and a #GMenuModel.
//
// The contents of @popover are removed and then refilled with menu items
// according to @model.  When @model changes, @popover is updated.
// Calling this function twice on @popover with different @model will
// cause the first binding to be replaced with a binding to the new
// model. If @model is %NULL then any previous binding is undone and
// all children are removed.
//
// If @action_namespace is non-%NULL then the effect is as if all
// actions mentioned in the @model have their names prefixed with the
// namespace, plus a dot.  For example, if the action “quit” is
// mentioned and @action_namespace is “app” then the effective action
// name is “app.quit”.
//
// This function uses #GtkActionable to define the action name and
// target values on the created menu items.  If you want to use an
// action group other than “app” and “win”, or if you want to use a
// #GtkMenuShell outside of a #GtkApplicationWindow, then you will need
// to attach your own action group to the widget hierarchy using
// gtk_widget_insert_action_group().  As an example, if you created a
// group with a “quit” action and inserted it with the name “mygroup”
// then you would use the action name “mygroup.quit” in your
// #GMenuModel.
/*

C function

gtk_popover_bind_model
*/
func (recv *Popover) BindModel(model *gio.MenuModel, actionNamespace string) {
	c_model := (*C.GMenuModel)(C.NULL)
	if model != nil {
		c_model = (*C.GMenuModel)(model.ToC())
	}

	c_action_namespace := C.CString(actionNamespace)
	defer C.free(unsafe.Pointer(c_action_namespace))

	C.gtk_popover_bind_model((*C.GtkPopover)(recv.native), c_model, c_action_namespace)

	return
}

// Returns whether the popover is modal, see gtk_popover_set_modal to
// see the implications of this.
/*

C function

gtk_popover_get_modal
*/
func (recv *Popover) GetModal() bool {
	retC := C.gtk_popover_get_modal((*C.GtkPopover)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Returns the widget @popover is currently attached to
/*

C function

gtk_popover_get_relative_to
*/
func (recv *Popover) GetRelativeTo() *Widget {
	retC := C.gtk_popover_get_relative_to((*C.GtkPopover)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Sets whether @popover is modal, a modal popover will grab all input
// within the toplevel and grab the keyboard focus on it when being
// displayed. Clicking outside the popover area or pressing Esc will
// dismiss the popover and ungrab input.
/*

C function

gtk_popover_set_modal
*/
func (recv *Popover) SetModal(modal bool) {
	c_modal :=
		boolToGboolean(modal)

	C.gtk_popover_set_modal((*C.GtkPopover)(recv.native), c_modal)

	return
}

// Unsupported : gtk_popover_set_pointing_to : unsupported parameter rect : Blacklisted record : GdkRectangle

// Sets the preferred position for @popover to appear. If the @popover
// is currently visible, it will be immediately updated.
//
// This preference will be respected where possible, although
// on lack of space (eg. if close to the window edges), the
// #GtkPopover may choose to appear on the opposite side
/*

C function

gtk_popover_set_position
*/
func (recv *Popover) SetPosition(position PositionType) {
	c_position := (C.GtkPositionType)(position)

	C.gtk_popover_set_position((*C.GtkPopover)(recv.native), c_position)

	return
}

// Sets a new widget to be attached to @popover. If @popover is
// visible, the position will be updated.
//
// Note: the ownership of popovers is always given to their @relative_to
// widget, so if @relative_to is set to %NULL on an attached @popover, it
// will be detached from its previous widget, and consequently destroyed
// unless extra references are kept.
/*

C function

gtk_popover_set_relative_to
*/
func (recv *Popover) SetRelativeTo(relativeTo *Widget) {
	c_relative_to := (*C.GtkWidget)(C.NULL)
	if relativeTo != nil {
		c_relative_to = (*C.GtkWidget)(relativeTo.ToC())
	}

	C.gtk_popover_set_relative_to((*C.GtkPopover)(recv.native), c_relative_to)

	return
}

// Finds the child of the #GtkStack with the name given as
// the argument. Returns %NULL if there is no child with this
// name.
/*

C function

gtk_stack_get_child_by_name
*/
func (recv *Stack) GetChildByName(name string) *Widget {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.gtk_stack_get_child_by_name((*C.GtkStack)(recv.native), c_name)
	var retGo (*Widget)
	if retC == nil {
		retGo = nil
	} else {
		retGo = WidgetNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Returns whether the @stack is currently in a transition from one page to
// another.
/*

C function

gtk_stack_get_transition_running
*/
func (recv *Stack) GetTransitionRunning() bool {
	retC := C.gtk_stack_get_transition_running((*C.GtkStack)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Gets the value of the #GtkWidget:margin-end property.
/*

C function

gtk_widget_get_margin_end
*/
func (recv *Widget) GetMarginEnd() int32 {
	retC := C.gtk_widget_get_margin_end((*C.GtkWidget)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Gets the value of the #GtkWidget:margin-start property.
/*

C function

gtk_widget_get_margin_start
*/
func (recv *Widget) GetMarginStart() int32 {
	retC := C.gtk_widget_get_margin_start((*C.GtkWidget)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Sets the end margin of @widget.
// See the #GtkWidget:margin-end property.
/*

C function

gtk_widget_set_margin_end
*/
func (recv *Widget) SetMarginEnd(margin int32) {
	c_margin := (C.gint)(margin)

	C.gtk_widget_set_margin_end((*C.GtkWidget)(recv.native), c_margin)

	return
}

// Sets the start margin of @widget.
// See the #GtkWidget:margin-start property.
/*

C function

gtk_widget_set_margin_start
*/
func (recv *Widget) SetMarginStart(margin int32) {
	c_margin := (C.gint)(margin)

	C.gtk_widget_set_margin_start((*C.GtkWidget)(recv.native), c_margin)

	return
}

// Retrieves the current maximized state of @window.
//
// Note that since maximization is ultimately handled by the window
// manager and happens asynchronously to an application request, you
// shouldn’t assume the return value of this function changing
// immediately (or at all), as an effect of calling
// gtk_window_maximize() or gtk_window_unmaximize().
/*

C function

gtk_window_is_maximized
*/
func (recv *Window) IsMaximized() bool {
	retC := C.gtk_window_is_maximized((*C.GtkWindow)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}
