// This is a generated file - DO NOT EDIT
// +build gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	gdk "github.com/pekim/gobbi/lib/gdk"
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

// Manually sets a keyval and modifier mask as the accelerator rendered
// by @accel_label.
//
// If a keyval and modifier are explicitly set then these values are
// used regardless of any associated accel closure or widget.
//
// Providing an @accelerator_key of 0 removes the manual setting.
/*

C function

gtk_accel_label_set_accel
*/
func (recv *AccelLabel) SetAccel(acceleratorKey uint32, acceleratorMods gdk.ModifierType) {
	c_accelerator_key := (C.guint)(acceleratorKey)

	c_accelerator_mods := (C.GdkModifierType)(acceleratorMods)

	C.gtk_accel_label_set_accel((*C.GtkAccelLabel)(recv.native), c_accelerator_key, c_accelerator_mods)

	return
}

// Gets the accelerator group.
/*

C function

gtk_action_group_get_accel_group
*/
func (recv *ActionGroup) GetAccelGroup() *AccelGroup {
	retC := C.gtk_action_group_get_accel_group((*C.GtkActionGroup)(recv.native))
	retGo := AccelGroupNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Sets the accelerator group to be used by every action in this group.
/*

C function

gtk_action_group_set_accel_group
*/
func (recv *ActionGroup) SetAccelGroup(accelGroup *AccelGroup) {
	c_accel_group := (*C.GtkAccelGroup)(C.NULL)
	if accelGroup != nil {
		c_accel_group = (*C.GtkAccelGroup)(accelGroup.ToC())
	}

	C.gtk_action_group_set_accel_group((*C.GtkActionGroup)(recv.native), c_accel_group)

	return
}

// Gets the “active” window for the application.
//
// The active window is the one that was most recently focused (within
// the application).  This window may not have the focus at the moment
// if another application has it — this is just the most
// recently-focused window within this application.
/*

C function

gtk_application_get_active_window
*/
func (recv *Application) GetActiveWindow() *Window {
	retC := C.gtk_application_get_active_window((*C.GtkApplication)(recv.native))
	var retGo (*Window)
	if retC == nil {
		retGo = nil
	} else {
		retGo = WindowNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Returns the #GtkApplicationWindow with the given ID.
//
// The ID of a #GtkApplicationWindow can be retrieved with
// gtk_application_window_get_id().
/*

C function

gtk_application_get_window_by_id
*/
func (recv *Application) GetWindowById(id uint32) *Window {
	c_id := (C.guint)(id)

	retC := C.gtk_application_get_window_by_id((*C.GtkApplication)(recv.native), c_id)
	var retGo (*Window)
	if retC == nil {
		retGo = nil
	} else {
		retGo = WindowNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Returns the unique ID of the window. If the window has not yet been added to
// a #GtkApplication, returns `0`.
/*

C function

gtk_application_window_get_id
*/
func (recv *ApplicationWindow) GetId() uint32 {
	retC := C.gtk_application_window_get_id((*C.GtkApplicationWindow)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// Returns whether the button will ignore the #GtkSettings:gtk-button-images
// setting and always show the image, if available.
/*

C function

gtk_button_get_always_show_image
*/
func (recv *Button) GetAlwaysShowImage() bool {
	retC := C.gtk_button_get_always_show_image((*C.GtkButton)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// If %TRUE, the button will ignore the #GtkSettings:gtk-button-images
// setting and always show the image, if available.
//
// Use this property if the button  would be useless or hard to use
// without the image.
/*

C function

gtk_button_set_always_show_image
*/
func (recv *Button) SetAlwaysShowImage(alwaysShow bool) {
	c_always_show :=
		boolToGboolean(alwaysShow)

	C.gtk_button_set_always_show_image((*C.GtkButton)(recv.native), c_always_show)

	return
}

// Gets the attribute list that was set on the entry using
// gtk_entry_set_attributes(), if any.
/*

C function

gtk_entry_get_attributes
*/
func (recv *Entry) GetAttributes() *pango.AttrList {
	retC := C.gtk_entry_get_attributes((*C.GtkEntry)(recv.native))
	var retGo (*pango.AttrList)
	if retC == nil {
		retGo = nil
	} else {
		retGo = pango.AttrListNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Gets the value of the #GtkEntry:input-hints property.
/*

C function

gtk_entry_get_input_hints
*/
func (recv *Entry) GetInputHints() InputHints {
	retC := C.gtk_entry_get_input_hints((*C.GtkEntry)(recv.native))
	retGo := (InputHints)(retC)

	return retGo
}

// Gets the value of the #GtkEntry:input-purpose property.
/*

C function

gtk_entry_get_input_purpose
*/
func (recv *Entry) GetInputPurpose() InputPurpose {
	retC := C.gtk_entry_get_input_purpose((*C.GtkEntry)(recv.native))
	retGo := (InputPurpose)(retC)

	return retGo
}

// Sets a #PangoAttrList; the attributes in the list are applied to the
// entry text.
/*

C function

gtk_entry_set_attributes
*/
func (recv *Entry) SetAttributes(attrs *pango.AttrList) {
	c_attrs := (*C.PangoAttrList)(C.NULL)
	if attrs != nil {
		c_attrs = (*C.PangoAttrList)(attrs.ToC())
	}

	C.gtk_entry_set_attributes((*C.GtkEntry)(recv.native), c_attrs)

	return
}

// Sets the #GtkEntry:input-hints property, which
// allows input methods to fine-tune their behaviour.
/*

C function

gtk_entry_set_input_hints
*/
func (recv *Entry) SetInputHints(hints InputHints) {
	c_hints := (C.GtkInputHints)(hints)

	C.gtk_entry_set_input_hints((*C.GtkEntry)(recv.native), c_hints)

	return
}

// Sets the #GtkEntry:input-purpose property which
// can be used by on-screen keyboards and other input
// methods to adjust their behaviour.
/*

C function

gtk_entry_set_input_purpose
*/
func (recv *Entry) SetInputPurpose(purpose InputPurpose) {
	c_purpose := (C.GtkInputPurpose)(purpose)

	C.gtk_entry_set_input_purpose((*C.GtkEntry)(recv.native), c_purpose)

	return
}

// Unsupported : gtk_icon_view_get_cell_rect : unsupported parameter rect : Blacklisted record : GdkRectangle

// Unsupported signal 'offset-changed' for LevelBar : unsupported parameter name : type utf8 :

// Creates a new #GtkLevelBar.
/*

C function

gtk_level_bar_new
*/
func LevelBarNew() *LevelBar {
	retC := C.gtk_level_bar_new()
	retGo := LevelBarNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Utility constructor that creates a new #GtkLevelBar for the specified
// interval.
/*

C function

gtk_level_bar_new_for_interval
*/
func LevelBarNewForInterval(minValue float64, maxValue float64) *LevelBar {
	c_min_value := (C.gdouble)(minValue)

	c_max_value := (C.gdouble)(maxValue)

	retC := C.gtk_level_bar_new_for_interval(c_min_value, c_max_value)
	retGo := LevelBarNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Adds a new offset marker on @self at the position specified by @value.
// When the bar value is in the interval topped by @value (or between @value
// and #GtkLevelBar:max-value in case the offset is the last one on the bar)
// a style class named `level-`@name will be applied
// when rendering the level bar fill.
// If another offset marker named @name exists, its value will be
// replaced by @value.
/*

C function

gtk_level_bar_add_offset_value
*/
func (recv *LevelBar) AddOffsetValue(name string, value float64) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_value := (C.gdouble)(value)

	C.gtk_level_bar_add_offset_value((*C.GtkLevelBar)(recv.native), c_name, c_value)

	return
}

// Returns the value of the #GtkLevelBar:max-value property.
/*

C function

gtk_level_bar_get_max_value
*/
func (recv *LevelBar) GetMaxValue() float64 {
	retC := C.gtk_level_bar_get_max_value((*C.GtkLevelBar)(recv.native))
	retGo := (float64)(retC)

	return retGo
}

// Returns the value of the #GtkLevelBar:min-value property.
/*

C function

gtk_level_bar_get_min_value
*/
func (recv *LevelBar) GetMinValue() float64 {
	retC := C.gtk_level_bar_get_min_value((*C.GtkLevelBar)(recv.native))
	retGo := (float64)(retC)

	return retGo
}

// Returns the value of the #GtkLevelBar:mode property.
/*

C function

gtk_level_bar_get_mode
*/
func (recv *LevelBar) GetMode() LevelBarMode {
	retC := C.gtk_level_bar_get_mode((*C.GtkLevelBar)(recv.native))
	retGo := (LevelBarMode)(retC)

	return retGo
}

// Fetches the value specified for the offset marker @name in @self,
// returning %TRUE in case an offset named @name was found.
/*

C function

gtk_level_bar_get_offset_value
*/
func (recv *LevelBar) GetOffsetValue(name string) (bool, float64) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	var c_value C.gdouble

	retC := C.gtk_level_bar_get_offset_value((*C.GtkLevelBar)(recv.native), c_name, &c_value)
	retGo := retC == C.TRUE

	value := (float64)(c_value)

	return retGo, value
}

// Returns the value of the #GtkLevelBar:value property.
/*

C function

gtk_level_bar_get_value
*/
func (recv *LevelBar) GetValue() float64 {
	retC := C.gtk_level_bar_get_value((*C.GtkLevelBar)(recv.native))
	retGo := (float64)(retC)

	return retGo
}

// Removes an offset marker previously added with
// gtk_level_bar_add_offset_value().
/*

C function

gtk_level_bar_remove_offset_value
*/
func (recv *LevelBar) RemoveOffsetValue(name string) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	C.gtk_level_bar_remove_offset_value((*C.GtkLevelBar)(recv.native), c_name)

	return
}

// Sets the value of the #GtkLevelBar:max-value property.
//
// You probably want to update preexisting level offsets after calling
// this function.
/*

C function

gtk_level_bar_set_max_value
*/
func (recv *LevelBar) SetMaxValue(value float64) {
	c_value := (C.gdouble)(value)

	C.gtk_level_bar_set_max_value((*C.GtkLevelBar)(recv.native), c_value)

	return
}

// Sets the value of the #GtkLevelBar:min-value property.
//
// You probably want to update preexisting level offsets after calling
// this function.
/*

C function

gtk_level_bar_set_min_value
*/
func (recv *LevelBar) SetMinValue(value float64) {
	c_value := (C.gdouble)(value)

	C.gtk_level_bar_set_min_value((*C.GtkLevelBar)(recv.native), c_value)

	return
}

// Sets the value of the #GtkLevelBar:mode property.
/*

C function

gtk_level_bar_set_mode
*/
func (recv *LevelBar) SetMode(mode LevelBarMode) {
	c_mode := (C.GtkLevelBarMode)(mode)

	C.gtk_level_bar_set_mode((*C.GtkLevelBar)(recv.native), c_mode)

	return
}

// Sets the value of the #GtkLevelBar:value property.
/*

C function

gtk_level_bar_set_value
*/
func (recv *LevelBar) SetValue(value float64) {
	c_value := (C.gdouble)(value)

	C.gtk_level_bar_set_value((*C.GtkLevelBar)(recv.native), c_value)

	return
}

// Creates a new #GtkMenuButton widget with downwards-pointing
// arrow as the only child. You can replace the child widget
// with another #GtkWidget should you wish to.
/*

C function

gtk_menu_button_new
*/
func MenuButtonNew() *MenuButton {
	retC := C.gtk_menu_button_new()
	retGo := MenuButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Returns the parent #GtkWidget to use to line up with menu.
/*

C function

gtk_menu_button_get_align_widget
*/
func (recv *MenuButton) GetAlignWidget() *Widget {
	retC := C.gtk_menu_button_get_align_widget((*C.GtkMenuButton)(recv.native))
	var retGo (*Widget)
	if retC == nil {
		retGo = nil
	} else {
		retGo = WidgetNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Returns the direction the popup will be pointing at when popped up.
/*

C function

gtk_menu_button_get_direction
*/
func (recv *MenuButton) GetDirection() ArrowType {
	retC := C.gtk_menu_button_get_direction((*C.GtkMenuButton)(recv.native))
	retGo := (ArrowType)(retC)

	return retGo
}

// Returns the #GMenuModel used to generate the popup.
/*

C function

gtk_menu_button_get_menu_model
*/
func (recv *MenuButton) GetMenuModel() *gio.MenuModel {
	retC := C.gtk_menu_button_get_menu_model((*C.GtkMenuButton)(recv.native))
	var retGo (*gio.MenuModel)
	if retC == nil {
		retGo = nil
	} else {
		retGo = gio.MenuModelNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Returns the #GtkMenu that pops out of the button.
// If the button does not use a #GtkMenu, this function
// returns %NULL.
/*

C function

gtk_menu_button_get_popup
*/
func (recv *MenuButton) GetPopup() *Menu {
	retC := C.gtk_menu_button_get_popup((*C.GtkMenuButton)(recv.native))
	var retGo (*Menu)
	if retC == nil {
		retGo = nil
	} else {
		retGo = MenuNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Sets the #GtkWidget to use to line the menu with when popped up.
// Note that the @align_widget must contain the #GtkMenuButton itself.
//
// Setting it to %NULL means that the menu will be aligned with the
// button itself.
//
// Note that this property is only used with menus currently,
// and not for popovers.
/*

C function

gtk_menu_button_set_align_widget
*/
func (recv *MenuButton) SetAlignWidget(alignWidget *Widget) {
	c_align_widget := (*C.GtkWidget)(C.NULL)
	if alignWidget != nil {
		c_align_widget = (*C.GtkWidget)(alignWidget.ToC())
	}

	C.gtk_menu_button_set_align_widget((*C.GtkMenuButton)(recv.native), c_align_widget)

	return
}

// Sets the direction in which the popup will be popped up, as
// well as changing the arrow’s direction. The child will not
// be changed to an arrow if it was customized.
//
// If the does not fit in the available space in the given direction,
// GTK+ will its best to keep it inside the screen and fully visible.
//
// If you pass %GTK_ARROW_NONE for a @direction, the popup will behave
// as if you passed %GTK_ARROW_DOWN (although you won’t see any arrows).
/*

C function

gtk_menu_button_set_direction
*/
func (recv *MenuButton) SetDirection(direction ArrowType) {
	c_direction := (C.GtkArrowType)(direction)

	C.gtk_menu_button_set_direction((*C.GtkMenuButton)(recv.native), c_direction)

	return
}

// Sets the #GMenuModel from which the popup will be constructed,
// or %NULL to disable the button.
//
// Depending on the value of #GtkMenuButton:use-popover, either a
// #GtkMenu will be created with gtk_menu_new_from_model(), or a
// #GtkPopover with gtk_popover_new_from_model(). In either case,
// actions will be connected as documented for these functions.
//
// If #GtkMenuButton:popup or #GtkMenuButton:popover are already set,
// their content will be lost and replaced by the newly created popup.
/*

C function

gtk_menu_button_set_menu_model
*/
func (recv *MenuButton) SetMenuModel(menuModel *gio.MenuModel) {
	c_menu_model := (*C.GMenuModel)(C.NULL)
	if menuModel != nil {
		c_menu_model = (*C.GMenuModel)(menuModel.ToC())
	}

	C.gtk_menu_button_set_menu_model((*C.GtkMenuButton)(recv.native), c_menu_model)

	return
}

// Sets the #GtkMenu that will be popped up when the button is clicked,
// or %NULL to disable the button. If #GtkMenuButton:menu-model or
// #GtkMenuButton:popover are set, they will be set to %NULL.
/*

C function

gtk_menu_button_set_popup
*/
func (recv *MenuButton) SetPopup(menu *Widget) {
	c_menu := (*C.GtkWidget)(C.NULL)
	if menu != nil {
		c_menu = (*C.GtkWidget)(menu.ToC())
	}

	C.gtk_menu_button_set_popup((*C.GtkMenuButton)(recv.native), c_menu)

	return
}

// Establishes a binding between a #GtkMenuShell and a #GMenuModel.
//
// The contents of @shell are removed and then refilled with menu items
// according to @model.  When @model changes, @shell is updated.
// Calling this function twice on @shell with different @model will
// cause the first binding to be replaced with a binding to the new
// model. If @model is %NULL then any previous binding is undone and
// all children are removed.
//
// @with_separators determines if toplevel items (eg: sections) have
// separators inserted between them.  This is typically desired for
// menus but doesn’t make sense for menubars.
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
//
// For most cases you are probably better off using
// gtk_menu_new_from_model() or gtk_menu_bar_new_from_model() or just
// directly passing the #GMenuModel to gtk_application_set_app_menu() or
// gtk_application_set_menubar().
/*

C function

gtk_menu_shell_bind_model
*/
func (recv *MenuShell) BindModel(model *gio.MenuModel, actionNamespace string, withSeparators bool) {
	c_model := (*C.GMenuModel)(C.NULL)
	if model != nil {
		c_model = (*C.GMenuModel)(model.ToC())
	}

	c_action_namespace := C.CString(actionNamespace)
	defer C.free(unsafe.Pointer(c_action_namespace))

	c_with_separators :=
		boolToGboolean(withSeparators)

	C.gtk_menu_shell_bind_model((*C.GtkMenuShell)(recv.native), c_model, c_action_namespace, c_with_separators)

	return
}

// Creates a #GtkSearchEntry, with a find icon when the search field is
// empty, and a clear icon when it isn't.
/*

C function

gtk_search_entry_new
*/
func SearchEntryNew() *SearchEntry {
	retC := C.gtk_search_entry_new()
	retGo := SearchEntryNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the value of the #GtkTextView:input-hints property.
/*

C function

gtk_text_view_get_input_hints
*/
func (recv *TextView) GetInputHints() InputHints {
	retC := C.gtk_text_view_get_input_hints((*C.GtkTextView)(recv.native))
	retGo := (InputHints)(retC)

	return retGo
}

// Gets the value of the #GtkTextView:input-purpose property.
/*

C function

gtk_text_view_get_input_purpose
*/
func (recv *TextView) GetInputPurpose() InputPurpose {
	retC := C.gtk_text_view_get_input_purpose((*C.GtkTextView)(recv.native))
	retGo := (InputPurpose)(retC)

	return retGo
}

// Sets the #GtkTextView:input-hints property, which
// allows input methods to fine-tune their behaviour.
/*

C function

gtk_text_view_set_input_hints
*/
func (recv *TextView) SetInputHints(hints InputHints) {
	c_hints := (C.GtkInputHints)(hints)

	C.gtk_text_view_set_input_hints((*C.GtkTextView)(recv.native), c_hints)

	return
}

// Sets the #GtkTextView:input-purpose property which
// can be used by on-screen keyboards and other input
// methods to adjust their behaviour.
/*

C function

gtk_text_view_set_input_purpose
*/
func (recv *TextView) SetInputPurpose(purpose InputPurpose) {
	c_purpose := (C.GtkInputPurpose)(purpose)

	C.gtk_text_view_set_input_purpose((*C.GtkTextView)(recv.native), c_purpose)

	return
}

// Inserts @group into @widget. Children of @widget that implement
// #GtkActionable can then be associated with actions in @group by
// setting their “action-name” to
// @prefix.`action-name`.
//
// If @group is %NULL, a previously inserted group for @name is removed
// from @widget.
/*

C function

gtk_widget_insert_action_group
*/
func (recv *Widget) InsertActionGroup(name string, group *gio.ActionGroup) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_group := (*C.GActionGroup)(group.ToC())

	C.gtk_widget_insert_action_group((*C.GtkWidget)(recv.native), c_name, c_group)

	return
}
