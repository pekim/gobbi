// Code generated - DO NOT EDIT.
// +build gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	cairo "github.com/pekim/gobbi/lib/cairo"
	gdk "github.com/pekim/gobbi/lib/gdk"
	gobject "github.com/pekim/gobbi/lib/gobject"
	pango "github.com/pekim/gobbi/lib/pango"
	"sync"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
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

// GetPageHasPadding is a wrapper around the C function gtk_assistant_get_page_has_padding.
func (recv *Assistant) GetPageHasPadding(page *Widget) bool {
	c_page := (*C.GtkWidget)(C.NULL)
	if page != nil {
		c_page = (*C.GtkWidget)(page.ToC())
	}

	retC := C.gtk_assistant_get_page_has_padding((*C.GtkAssistant)(recv.native), c_page)
	retGo := retC == C.TRUE

	return retGo
}

// SetPageHasPadding is a wrapper around the C function gtk_assistant_set_page_has_padding.
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

// ChildNotifyByPspec is a wrapper around the C function gtk_container_child_notify_by_pspec.
func (recv *Container) ChildNotifyByPspec(child *Widget, pspec *gobject.ParamSpec) {
	c_child := (*C.GtkWidget)(C.NULL)
	if child != nil {
		c_child = (*C.GtkWidget)(child.ToC())
	}

	c_pspec := (*C.GParamSpec)(C.NULL)
	if pspec != nil {
		c_pspec = (*C.GParamSpec)(pspec.ToC())
	}

	C.gtk_container_child_notify_by_pspec((*C.GtkContainer)(recv.native), c_child, c_pspec)

	return
}

// Unsupported : gtk_flow_box_bind_model : unsupported parameter create_widget_func : no type generator for FlowBoxCreateWidgetFunc (GtkFlowBoxCreateWidgetFunc) for param create_widget_func

// GetOverlayPassThrough is a wrapper around the C function gtk_overlay_get_overlay_pass_through.
func (recv *Overlay) GetOverlayPassThrough(widget *Widget) bool {
	c_widget := (*C.GtkWidget)(C.NULL)
	if widget != nil {
		c_widget = (*C.GtkWidget)(widget.ToC())
	}

	retC := C.gtk_overlay_get_overlay_pass_through((*C.GtkOverlay)(recv.native), c_widget)
	retGo := retC == C.TRUE

	return retGo
}

// ReorderOverlay is a wrapper around the C function gtk_overlay_reorder_overlay.
func (recv *Overlay) ReorderOverlay(child *Widget, position int32) {
	c_child := (*C.GtkWidget)(C.NULL)
	if child != nil {
		c_child = (*C.GtkWidget)(child.ToC())
	}

	c_position := (C.gint)(position)

	C.gtk_overlay_reorder_overlay((*C.GtkOverlay)(recv.native), c_child, c_position)

	return
}

// SetOverlayPassThrough is a wrapper around the C function gtk_overlay_set_overlay_pass_through.
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
var signalPlacesSidebarShowOtherLocationsLock sync.RWMutex

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
	signalPlacesSidebarShowOtherLocationsLock.RLock()
	defer signalPlacesSidebarShowOtherLocationsLock.RUnlock()

	index := int(uintptr(data))
	callback := signalPlacesSidebarShowOtherLocationsMap[index].callback
	callback()
}

// GetShowOtherLocations is a wrapper around the C function gtk_places_sidebar_get_show_other_locations.
func (recv *PlacesSidebar) GetShowOtherLocations() bool {
	retC := C.gtk_places_sidebar_get_show_other_locations((*C.GtkPlacesSidebar)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetShowRecent is a wrapper around the C function gtk_places_sidebar_get_show_recent.
func (recv *PlacesSidebar) GetShowRecent() bool {
	retC := C.gtk_places_sidebar_get_show_recent((*C.GtkPlacesSidebar)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetShowTrash is a wrapper around the C function gtk_places_sidebar_get_show_trash.
func (recv *PlacesSidebar) GetShowTrash() bool {
	retC := C.gtk_places_sidebar_get_show_trash((*C.GtkPlacesSidebar)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetDropTargetsVisible is a wrapper around the C function gtk_places_sidebar_set_drop_targets_visible.
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

// SetShowOtherLocations is a wrapper around the C function gtk_places_sidebar_set_show_other_locations.
func (recv *PlacesSidebar) SetShowOtherLocations(showOtherLocations bool) {
	c_show_other_locations :=
		boolToGboolean(showOtherLocations)

	C.gtk_places_sidebar_set_show_other_locations((*C.GtkPlacesSidebar)(recv.native), c_show_other_locations)

	return
}

// SetShowRecent is a wrapper around the C function gtk_places_sidebar_set_show_recent.
func (recv *PlacesSidebar) SetShowRecent(showRecent bool) {
	c_show_recent :=
		boolToGboolean(showRecent)

	C.gtk_places_sidebar_set_show_recent((*C.GtkPlacesSidebar)(recv.native), c_show_recent)

	return
}

// SetShowTrash is a wrapper around the C function gtk_places_sidebar_set_show_trash.
func (recv *PlacesSidebar) SetShowTrash(showTrash bool) {
	c_show_trash :=
		boolToGboolean(showTrash)

	C.gtk_places_sidebar_set_show_trash((*C.GtkPlacesSidebar)(recv.native), c_show_trash)

	return
}

// GetDefaultWidget is a wrapper around the C function gtk_popover_get_default_widget.
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

// SetDefaultWidget is a wrapper around the C function gtk_popover_set_default_widget.
func (recv *Popover) SetDefaultWidget(widget *Widget) {
	c_widget := (*C.GtkWidget)(C.NULL)
	if widget != nil {
		c_widget = (*C.GtkWidget)(widget.ToC())
	}

	C.gtk_popover_set_default_widget((*C.GtkPopover)(recv.native), c_widget)

	return
}

// JoinGroup is a wrapper around the C function gtk_radio_menu_item_join_group.
func (recv *RadioMenuItem) JoinGroup(groupSource *RadioMenuItem) {
	c_group_source := (*C.GtkRadioMenuItem)(C.NULL)
	if groupSource != nil {
		c_group_source = (*C.GtkRadioMenuItem)(groupSource.ToC())
	}

	C.gtk_radio_menu_item_join_group((*C.GtkRadioMenuItem)(recv.native), c_group_source)

	return
}

// GetInterpolateSize is a wrapper around the C function gtk_stack_get_interpolate_size.
func (recv *Stack) GetInterpolateSize() bool {
	retC := C.gtk_stack_get_interpolate_size((*C.GtkStack)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetInterpolateSize is a wrapper around the C function gtk_stack_set_interpolate_size.
func (recv *Stack) SetInterpolateSize(interpolateSize bool) {
	c_interpolate_size :=
		boolToGboolean(interpolateSize)

	C.gtk_stack_set_interpolate_size((*C.GtkStack)(recv.native), c_interpolate_size)

	return
}

// GetBottomMargin is a wrapper around the C function gtk_text_view_get_bottom_margin.
func (recv *TextView) GetBottomMargin() int32 {
	retC := C.gtk_text_view_get_bottom_margin((*C.GtkTextView)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetTopMargin is a wrapper around the C function gtk_text_view_get_top_margin.
func (recv *TextView) GetTopMargin() int32 {
	retC := C.gtk_text_view_get_top_margin((*C.GtkTextView)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// SetBottomMargin is a wrapper around the C function gtk_text_view_set_bottom_margin.
func (recv *TextView) SetBottomMargin(bottomMargin int32) {
	c_bottom_margin := (C.gint)(bottomMargin)

	C.gtk_text_view_set_bottom_margin((*C.GtkTextView)(recv.native), c_bottom_margin)

	return
}

// SetTopMargin is a wrapper around the C function gtk_text_view_set_top_margin.
func (recv *TextView) SetTopMargin(topMargin int32) {
	c_top_margin := (C.gint)(topMargin)

	C.gtk_text_view_set_top_margin((*C.GtkTextView)(recv.native), c_top_margin)

	return
}

// GetFontMap is a wrapper around the C function gtk_widget_get_font_map.
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

// GetFontOptions is a wrapper around the C function gtk_widget_get_font_options.
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

// SetFontMap is a wrapper around the C function gtk_widget_set_font_map.
func (recv *Widget) SetFontMap(fontMap *pango.FontMap) {
	c_font_map := (*C.PangoFontMap)(C.NULL)
	if fontMap != nil {
		c_font_map = (*C.PangoFontMap)(fontMap.ToC())
	}

	C.gtk_widget_set_font_map((*C.GtkWidget)(recv.native), c_font_map)

	return
}

// SetFontOptions is a wrapper around the C function gtk_widget_set_font_options.
func (recv *Widget) SetFontOptions(options *cairo.FontOptions) {
	c_options := (*C.cairo_font_options_t)(C.NULL)
	if options != nil {
		c_options = (*C.cairo_font_options_t)(options.ToC())
	}

	C.gtk_widget_set_font_options((*C.GtkWidget)(recv.native), c_options)

	return
}

// FullscreenOnMonitor is a wrapper around the C function gtk_window_fullscreen_on_monitor.
func (recv *Window) FullscreenOnMonitor(screen *gdk.Screen, monitor int32) {
	c_screen := (*C.GdkScreen)(C.NULL)
	if screen != nil {
		c_screen = (*C.GdkScreen)(screen.ToC())
	}

	c_monitor := (C.gint)(monitor)

	C.gtk_window_fullscreen_on_monitor((*C.GtkWindow)(recv.native), c_screen, c_monitor)

	return
}

// GetFontMap is a wrapper around the C function gtk_font_chooser_get_font_map.
func (recv *FontChooser) GetFontMap() *pango.FontMap {
	retC := C.gtk_font_chooser_get_font_map((*C.GtkFontChooser)(recv.native))
	var retGo (*pango.FontMap)
	if retC == nil {
		retGo = nil
	} else {
		retGo = pango.FontMapNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// SetFontMap is a wrapper around the C function gtk_font_chooser_set_font_map.
func (recv *FontChooser) SetFontMap(fontmap *pango.FontMap) {
	c_fontmap := (*C.PangoFontMap)(C.NULL)
	if fontmap != nil {
		c_fontmap = (*C.PangoFontMap)(fontmap.ToC())
	}

	C.gtk_font_chooser_set_font_map((*C.GtkFontChooser)(recv.native), c_fontmap)

	return
}

// Unsupported : gtk_container_class_install_child_properties : unsupported parameter pspecs :
