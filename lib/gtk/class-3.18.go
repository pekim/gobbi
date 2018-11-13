// This is a generated file - DO NOT EDIT
// +build gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	cairo "github.com/pekim/gobbi/lib/cairo"
	gdk "github.com/pekim/gobbi/lib/gdk"
	pango "github.com/pekim/gobbi/lib/pango"
	"sync"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
/*

	void placessidebar_showOtherLocationsHandler(GObject *, gpointer);

	static gulong PlacesSidebar_signal_connect_show_other_locations(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "show-other-locations", G_CALLBACK(placessidebar_showOtherLocationsHandler), data);
	}

*/
import "C"

// Gets whether page has padding.
/*

C function

gtk_assistant_get_page_has_padding
*/
func (recv *Assistant) GetPageHasPadding(page *Widget) bool {
	c_page := (*C.GtkWidget)(C.NULL)
	if page != nil {
		c_page = (*C.GtkWidget)(page.ToC())
	}

	retC := C.gtk_assistant_get_page_has_padding((*C.GtkAssistant)(recv.native), c_page)
	retGo := retC == C.TRUE

	return retGo
}

// Sets whether the assistant is adding padding around
// the page.
/*

C function

gtk_assistant_set_page_has_padding
*/
func (recv *Assistant) SetPageHasPadding(page *Widget, hasPadding bool) {
	c_page := (*C.GtkWidget)(C.NULL)
	if page != nil {
		c_page = (*C.GtkWidget)(page.ToC())
	}

	c_has_padding :=
		boolToGboolean(hasPadding)

	C.gtk_assistant_set_page_has_padding((*C.GtkAssistant)(recv.native), c_page, c_has_padding)

	return
}

// Unsupported : gtk_container_child_notify_by_pspec : unsupported parameter pspec : Blacklisted record : GParamSpec

// Unsupported : gtk_flow_box_bind_model : unsupported parameter create_widget_func : no type generator for FlowBoxCreateWidgetFunc (GtkFlowBoxCreateWidgetFunc) for param create_widget_func

// Convenience function to get the value of the #GtkOverlay:pass-through
// child property for @widget.
/*

C function

gtk_overlay_get_overlay_pass_through
*/
func (recv *Overlay) GetOverlayPassThrough(widget *Widget) bool {
	c_widget := (*C.GtkWidget)(C.NULL)
	if widget != nil {
		c_widget = (*C.GtkWidget)(widget.ToC())
	}

	retC := C.gtk_overlay_get_overlay_pass_through((*C.GtkOverlay)(recv.native), c_widget)
	retGo := retC == C.TRUE

	return retGo
}

// Moves @child to a new @index in the list of @overlay children.
// The list contains overlays in the order that these were
// added to @overlay.
//
// A widget’s index in the @overlay children list determines which order
// the children are drawn if they overlap. The first child is drawn at
// the bottom. It also affects the default focus chain order.
/*

C function

gtk_overlay_reorder_overlay
*/
func (recv *Overlay) ReorderOverlay(child *Widget, position int32) {
	c_child := (*C.GtkWidget)(C.NULL)
	if child != nil {
		c_child = (*C.GtkWidget)(child.ToC())
	}

	c_position := (C.gint)(position)

	C.gtk_overlay_reorder_overlay((*C.GtkOverlay)(recv.native), c_child, c_position)

	return
}

// Convenience function to set the value of the #GtkOverlay:pass-through
// child property for @widget.
/*

C function

gtk_overlay_set_overlay_pass_through
*/
func (recv *Overlay) SetOverlayPassThrough(widget *Widget, passThrough bool) {
	c_widget := (*C.GtkWidget)(C.NULL)
	if widget != nil {
		c_widget = (*C.GtkWidget)(widget.ToC())
	}

	c_pass_through :=
		boolToGboolean(passThrough)

	C.gtk_overlay_set_overlay_pass_through((*C.GtkOverlay)(recv.native), c_widget, c_pass_through)

	return
}

type signalPlacesSidebarShowOtherLocationsDetail struct {
	callback  PlacesSidebarSignalShowOtherLocationsCallback
	handlerID C.gulong
}

var signalPlacesSidebarShowOtherLocationsId int
var signalPlacesSidebarShowOtherLocationsMap = make(map[int]signalPlacesSidebarShowOtherLocationsDetail)
var signalPlacesSidebarShowOtherLocationsLock sync.Mutex

// PlacesSidebarSignalShowOtherLocationsCallback is a callback function for a 'show-other-locations' signal emitted from a PlacesSidebar.
type PlacesSidebarSignalShowOtherLocationsCallback func()

/*
ConnectShowOtherLocations connects the callback to the 'show-other-locations' signal for the PlacesSidebar.

The returned value represents the connection, and may be passed to DisconnectShowOtherLocations to remove it.
*/
func (recv *PlacesSidebar) ConnectShowOtherLocations(callback PlacesSidebarSignalShowOtherLocationsCallback) int {
	signalPlacesSidebarShowOtherLocationsLock.Lock()
	defer signalPlacesSidebarShowOtherLocationsLock.Unlock()

	signalPlacesSidebarShowOtherLocationsId++
	instance := C.gpointer(recv.native)
	handlerID := C.PlacesSidebar_signal_connect_show_other_locations(instance, C.gpointer(uintptr(signalPlacesSidebarShowOtherLocationsId)))

	detail := signalPlacesSidebarShowOtherLocationsDetail{callback, handlerID}
	signalPlacesSidebarShowOtherLocationsMap[signalPlacesSidebarShowOtherLocationsId] = detail

	return signalPlacesSidebarShowOtherLocationsId
}

/*
DisconnectShowOtherLocations disconnects a callback from the 'show-other-locations' signal for the PlacesSidebar.

The connectionID should be a value returned from a call to ConnectShowOtherLocations.
*/
func (recv *PlacesSidebar) DisconnectShowOtherLocations(connectionID int) {
	signalPlacesSidebarShowOtherLocationsLock.Lock()
	defer signalPlacesSidebarShowOtherLocationsLock.Unlock()

	detail, exists := signalPlacesSidebarShowOtherLocationsMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalPlacesSidebarShowOtherLocationsMap, connectionID)
}

//export placessidebar_showOtherLocationsHandler
func placessidebar_showOtherLocationsHandler(_ *C.GObject, data C.gpointer) {
	index := int(uintptr(data))
	callback := signalPlacesSidebarShowOtherLocationsMap[index].callback
	callback()
}

// Returns the value previously set with gtk_places_sidebar_set_show_other_locations()
/*

C function

gtk_places_sidebar_get_show_other_locations
*/
func (recv *PlacesSidebar) GetShowOtherLocations() bool {
	retC := C.gtk_places_sidebar_get_show_other_locations((*C.GtkPlacesSidebar)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Returns the value previously set with gtk_places_sidebar_set_show_recent()
/*

C function

gtk_places_sidebar_get_show_recent
*/
func (recv *PlacesSidebar) GetShowRecent() bool {
	retC := C.gtk_places_sidebar_get_show_recent((*C.GtkPlacesSidebar)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Returns the value previously set with gtk_places_sidebar_set_show_trash()
/*

C function

gtk_places_sidebar_get_show_trash
*/
func (recv *PlacesSidebar) GetShowTrash() bool {
	retC := C.gtk_places_sidebar_get_show_trash((*C.GtkPlacesSidebar)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Make the GtkPlacesSidebar show drop targets, so it can show the available
// drop targets and a "new bookmark" row. This improves the Drag-and-Drop
// experience of the user and allows applications to show all available
// drop targets at once.
//
// This needs to be called when the application is aware of an ongoing drag
// that might target the sidebar. The drop-targets-visible state will be unset
// automatically if the drag finishes in the GtkPlacesSidebar. You only need
// to unset the state when the drag ends on some other widget on your application.
/*

C function

gtk_places_sidebar_set_drop_targets_visible
*/
func (recv *PlacesSidebar) SetDropTargetsVisible(visible bool, context *gdk.DragContext) {
	c_visible :=
		boolToGboolean(visible)

	c_context := (*C.GdkDragContext)(C.NULL)
	if context != nil {
		c_context = (*C.GdkDragContext)(context.ToC())
	}

	C.gtk_places_sidebar_set_drop_targets_visible((*C.GtkPlacesSidebar)(recv.native), c_visible, c_context)

	return
}

// Sets whether the @sidebar should show an item for the application to show
// an Other Locations view; this is off by default. When set to %TRUE, persistent
// devices such as hard drives are hidden, otherwise they are shown in the sidebar.
// An application may want to turn this on if it implements a way for the user to
// see and interact with drives and network servers directly.
//
// If you enable this, you should connect to the
// #GtkPlacesSidebar::show-other-locations signal.
/*

C function

gtk_places_sidebar_set_show_other_locations
*/
func (recv *PlacesSidebar) SetShowOtherLocations(showOtherLocations bool) {
	c_show_other_locations :=
		boolToGboolean(showOtherLocations)

	C.gtk_places_sidebar_set_show_other_locations((*C.GtkPlacesSidebar)(recv.native), c_show_other_locations)

	return
}

// Sets whether the @sidebar should show an item for recent files.
// The default value for this option is determined by the desktop
// environment, but this function can be used to override it on a
// per-application basis.
/*

C function

gtk_places_sidebar_set_show_recent
*/
func (recv *PlacesSidebar) SetShowRecent(showRecent bool) {
	c_show_recent :=
		boolToGboolean(showRecent)

	C.gtk_places_sidebar_set_show_recent((*C.GtkPlacesSidebar)(recv.native), c_show_recent)

	return
}

// Sets whether the @sidebar should show an item for the Trash location.
/*

C function

gtk_places_sidebar_set_show_trash
*/
func (recv *PlacesSidebar) SetShowTrash(showTrash bool) {
	c_show_trash :=
		boolToGboolean(showTrash)

	C.gtk_places_sidebar_set_show_trash((*C.GtkPlacesSidebar)(recv.native), c_show_trash)

	return
}

// Gets the widget that should be set as the default while
// the popover is shown.
/*

C function

gtk_popover_get_default_widget
*/
func (recv *Popover) GetDefaultWidget() *Widget {
	retC := C.gtk_popover_get_default_widget((*C.GtkPopover)(recv.native))
	var retGo (*Widget)
	if retC == nil {
		retGo = nil
	} else {
		retGo = WidgetNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Sets the widget that should be set as default widget while
// the popover is shown (see gtk_window_set_default()). #GtkPopover
// remembers the previous default widget and reestablishes it
// when the popover is dismissed.
/*

C function

gtk_popover_set_default_widget
*/
func (recv *Popover) SetDefaultWidget(widget *Widget) {
	c_widget := (*C.GtkWidget)(C.NULL)
	if widget != nil {
		c_widget = (*C.GtkWidget)(widget.ToC())
	}

	C.gtk_popover_set_default_widget((*C.GtkPopover)(recv.native), c_widget)

	return
}

// Joins a #GtkRadioMenuItem object to the group of another #GtkRadioMenuItem
// object.
//
// This function should be used by language bindings to avoid the memory
// manangement of the opaque #GSList of gtk_radio_menu_item_get_group()
// and gtk_radio_menu_item_set_group().
//
// A common way to set up a group of #GtkRadioMenuItem instances is:
//
// |[
// GtkRadioMenuItem *last_item = NULL;
//
// while ( ...more items to add... )
// {
// GtkRadioMenuItem *radio_item;
//
// radio_item = gtk_radio_menu_item_new (...);
//
// gtk_radio_menu_item_join_group (radio_item, last_item);
// last_item = radio_item;
// }
// ]|
/*

C function

gtk_radio_menu_item_join_group
*/
func (recv *RadioMenuItem) JoinGroup(groupSource *RadioMenuItem) {
	c_group_source := (*C.GtkRadioMenuItem)(C.NULL)
	if groupSource != nil {
		c_group_source = (*C.GtkRadioMenuItem)(groupSource.ToC())
	}

	C.gtk_radio_menu_item_join_group((*C.GtkRadioMenuItem)(recv.native), c_group_source)

	return
}

// Returns wether the #GtkStack is set up to interpolate between
// the sizes of children on page switch.
/*

C function

gtk_stack_get_interpolate_size
*/
func (recv *Stack) GetInterpolateSize() bool {
	retC := C.gtk_stack_get_interpolate_size((*C.GtkStack)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Sets whether or not @stack will interpolate its size when
// changing the visible child. If the #GtkStack:interpolate-size
// property is set to %TRUE, @stack will interpolate its size between
// the current one and the one it'll take after changing the
// visible child, according to the set transition duration.
/*

C function

gtk_stack_set_interpolate_size
*/
func (recv *Stack) SetInterpolateSize(interpolateSize bool) {
	c_interpolate_size :=
		boolToGboolean(interpolateSize)

	C.gtk_stack_set_interpolate_size((*C.GtkStack)(recv.native), c_interpolate_size)

	return
}

// Gets the bottom margin for text in the @text_view.
/*

C function

gtk_text_view_get_bottom_margin
*/
func (recv *TextView) GetBottomMargin() int32 {
	retC := C.gtk_text_view_get_bottom_margin((*C.GtkTextView)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Gets the top margin for text in the @text_view.
/*

C function

gtk_text_view_get_top_margin
*/
func (recv *TextView) GetTopMargin() int32 {
	retC := C.gtk_text_view_get_top_margin((*C.GtkTextView)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Sets the bottom margin for text in @text_view.
//
// Note that this function is confusingly named.
// In CSS terms, the value set here is padding.
/*

C function

gtk_text_view_set_bottom_margin
*/
func (recv *TextView) SetBottomMargin(bottomMargin int32) {
	c_bottom_margin := (C.gint)(bottomMargin)

	C.gtk_text_view_set_bottom_margin((*C.GtkTextView)(recv.native), c_bottom_margin)

	return
}

// Sets the top margin for text in @text_view.
//
// Note that this function is confusingly named.
// In CSS terms, the value set here is padding.
/*

C function

gtk_text_view_set_top_margin
*/
func (recv *TextView) SetTopMargin(topMargin int32) {
	c_top_margin := (C.gint)(topMargin)

	C.gtk_text_view_set_top_margin((*C.GtkTextView)(recv.native), c_top_margin)

	return
}

// Gets the font map that has been set with gtk_widget_set_font_map().
/*

C function

gtk_widget_get_font_map
*/
func (recv *Widget) GetFontMap() *pango.FontMap {
	retC := C.gtk_widget_get_font_map((*C.GtkWidget)(recv.native))
	var retGo (*pango.FontMap)
	if retC == nil {
		retGo = nil
	} else {
		retGo = pango.FontMapNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Returns the #cairo_font_options_t used for Pango rendering. When not set,
// the defaults font options for the #GdkScreen will be used.
/*

C function

gtk_widget_get_font_options
*/
func (recv *Widget) GetFontOptions() *cairo.FontOptions {
	retC := C.gtk_widget_get_font_options((*C.GtkWidget)(recv.native))
	var retGo (*cairo.FontOptions)
	if retC == nil {
		retGo = nil
	} else {
		retGo = cairo.FontOptionsNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Sets the font map to use for Pango rendering. When not set, the widget
// will inherit the font map from its parent.
/*

C function

gtk_widget_set_font_map
*/
func (recv *Widget) SetFontMap(fontMap *pango.FontMap) {
	c_font_map := (*C.PangoFontMap)(C.NULL)
	if fontMap != nil {
		c_font_map = (*C.PangoFontMap)(fontMap.ToC())
	}

	C.gtk_widget_set_font_map((*C.GtkWidget)(recv.native), c_font_map)

	return
}

// Sets the #cairo_font_options_t used for Pango rendering in this widget.
// When not set, the default font options for the #GdkScreen will be used.
/*

C function

gtk_widget_set_font_options
*/
func (recv *Widget) SetFontOptions(options *cairo.FontOptions) {
	c_options := (*C.cairo_font_options_t)(C.NULL)
	if options != nil {
		c_options = (*C.cairo_font_options_t)(options.ToC())
	}

	C.gtk_widget_set_font_options((*C.GtkWidget)(recv.native), c_options)

	return
}

// Asks to place @window in the fullscreen state. Note that you shouldn't assume
// the window is definitely full screen afterward.
//
// You can track the fullscreen state via the "window-state-event" signal
// on #GtkWidget.
/*

C function

gtk_window_fullscreen_on_monitor
*/
func (recv *Window) FullscreenOnMonitor(screen *gdk.Screen, monitor int32) {
	c_screen := (*C.GdkScreen)(C.NULL)
	if screen != nil {
		c_screen = (*C.GdkScreen)(screen.ToC())
	}

	c_monitor := (C.gint)(monitor)

	C.gtk_window_fullscreen_on_monitor((*C.GtkWindow)(recv.native), c_screen, c_monitor)

	return
}
