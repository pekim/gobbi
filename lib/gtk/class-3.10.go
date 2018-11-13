// This is a generated file - DO NOT EDIT
// +build gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	cairo "github.com/pekim/gobbi/lib/cairo"
	gdk "github.com/pekim/gobbi/lib/gdk"
	gdkpixbuf "github.com/pekim/gobbi/lib/gdkpixbuf"
	gio "github.com/pekim/gobbi/lib/gio"
	glib "github.com/pekim/gobbi/lib/glib"
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

	void listbox_rowActivatedHandler(GObject *, GtkListBoxRow *, gpointer);

	static gulong ListBox_signal_connect_row_activated(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "row-activated", G_CALLBACK(listbox_rowActivatedHandler), data);
	}

*/
/*

	void listbox_rowSelectedHandler(GObject *, GtkListBoxRow *, gpointer);

	static gulong ListBox_signal_connect_row_selected(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "row-selected", G_CALLBACK(listbox_rowSelectedHandler), data);
	}

*/
/*

	void listboxrow_activateHandler(GObject *, gpointer);

	static gulong ListBoxRow_signal_connect_activate(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "activate", G_CALLBACK(listboxrow_activateHandler), data);
	}

*/
/*

	void placessidebar_populatePopupHandler(GObject *, GtkWidget *, GFile *, GVolume *, gpointer);

	static gulong PlacesSidebar_signal_connect_populate_popup(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "populate-popup", G_CALLBACK(placessidebar_populatePopupHandler), data);
	}

*/
/*

	void searchentry_searchChangedHandler(GObject *, gpointer);

	static gulong SearchEntry_signal_connect_search_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "search-changed", G_CALLBACK(searchentry_searchChangedHandler), data);
	}

*/
import "C"

// Gets the value set by gtk_box_set_baseline_position().
/*

C function : gtk_box_get_baseline_position
*/
func (recv *Box) GetBaselinePosition() BaselinePosition {
	retC := C.gtk_box_get_baseline_position((*C.GtkBox)(recv.native))
	retGo := (BaselinePosition)(retC)

	return retGo
}

// Sets the baseline position of a box. This affects
// only horizontal boxes with at least one baseline aligned
// child. If there is more vertical space available than requested,
// and the baseline is not allocated by the parent then
// @position is used to allocate the baseline wrt the
// extra space available.
/*

C function : gtk_box_set_baseline_position
*/
func (recv *Box) SetBaselinePosition(position BaselinePosition) {
	c_position := (C.GtkBaselinePosition)(position)

	C.gtk_box_set_baseline_position((*C.GtkBox)(recv.native), c_position)

	return
}

// Builds the [GtkBuilder UI definition][BUILDER-UI]
// in the file @filename.
//
// If there is an error opening the file or parsing the description then
// the program will be aborted.  You should only ever attempt to parse
// user interface descriptions that are shipped as part of your program.
/*

C function : gtk_builder_new_from_file
*/
func BuilderNewFromFile(filename string) *Builder {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))

	retC := C.gtk_builder_new_from_file(c_filename)
	retGo := BuilderNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Builds the [GtkBuilder UI definition][BUILDER-UI]
// at @resource_path.
//
// If there is an error locating the resource or parsing the
// description, then the program will be aborted.
/*

C function : gtk_builder_new_from_resource
*/
func BuilderNewFromResource(resourcePath string) *Builder {
	c_resource_path := C.CString(resourcePath)
	defer C.free(unsafe.Pointer(c_resource_path))

	retC := C.gtk_builder_new_from_resource(c_resource_path)
	retGo := BuilderNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Builds the user interface described by @string (in the
// [GtkBuilder UI definition][BUILDER-UI] format).
//
// If @string is %NULL-terminated, then @length should be -1.
// If @length is not -1, then it is the length of @string.
//
// If there is an error parsing @string then the program will be
// aborted. You should not attempt to parse user interface description
// from untrusted sources.
/*

C function : gtk_builder_new_from_string
*/
func BuilderNewFromString(string string) *Builder {
	c_string := C.CString(string)
	defer C.free(unsafe.Pointer(c_string))

	c_length := (C.gssize)(len(string))

	retC := C.gtk_builder_new_from_string(c_string, c_length)
	retGo := BuilderNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_builder_add_callback_symbol : unsupported parameter callback_symbol : no type generator for GObject.Callback (GCallback) for param callback_symbol

// Unsupported : gtk_builder_add_callback_symbols : unsupported parameter first_callback_symbol : no type generator for GObject.Callback (GCallback) for param first_callback_symbol

// Gets the #GtkApplication associated with the builder.
//
// The #GtkApplication is used for creating action proxies as requested
// from XML that the builder is loading.
//
// By default, the builder uses the default application: the one from
// g_application_get_default(). If you want to use another application
// for constructing proxies, use gtk_builder_set_application().
/*

C function : gtk_builder_get_application
*/
func (recv *Builder) GetApplication() *Application {
	retC := C.gtk_builder_get_application((*C.GtkBuilder)(recv.native))
	var retGo (*Application)
	if retC == nil {
		retGo = nil
	} else {
		retGo = ApplicationNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Unsupported : gtk_builder_lookup_callback_symbol : no return generator

// Sets the application associated with @builder.
//
// You only need this function if there is more than one #GApplication
// in your process. @application cannot be %NULL.
/*

C function : gtk_builder_set_application
*/
func (recv *Builder) SetApplication(application *Application) {
	c_application := (*C.GtkApplication)(C.NULL)
	if application != nil {
		c_application = (*C.GtkApplication)(application.ToC())
	}

	C.gtk_builder_set_application((*C.GtkBuilder)(recv.native), c_application)

	return
}

// Creates a new button containing an icon from the current icon theme.
//
// If the icon name isn’t known, a “broken image” icon will be
// displayed instead. If the current icon theme is changed, the icon
// will be updated appropriately.
//
// This function is a convenience wrapper around gtk_button_new() and
// gtk_button_set_image().
/*

C function : gtk_button_new_from_icon_name
*/
func ButtonNewFromIconName(iconName string, size IconSize) *Button {
	c_icon_name := C.CString(iconName)
	defer C.free(unsafe.Pointer(c_icon_name))

	c_size := (C.GtkIconSize)(size)

	retC := C.gtk_button_new_from_icon_name(c_icon_name, c_size)
	retGo := ButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the tabstops that were set on the entry using gtk_entry_set_tabs(), if
// any.
/*

C function : gtk_entry_get_tabs
*/
func (recv *Entry) GetTabs() *pango.TabArray {
	retC := C.gtk_entry_get_tabs((*C.GtkEntry)(recv.native))
	var retGo (*pango.TabArray)
	if retC == nil {
		retGo = nil
	} else {
		retGo = pango.TabArrayNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Sets a #PangoTabArray; the tabstops in the array are applied to the entry
// text.
/*

C function : gtk_entry_set_tabs
*/
func (recv *Entry) SetTabs(tabs *pango.TabArray) {
	c_tabs := (*C.PangoTabArray)(C.NULL)
	if tabs != nil {
		c_tabs = (*C.PangoTabArray)(tabs.ToC())
	}

	C.gtk_entry_set_tabs((*C.GtkEntry)(recv.native), c_tabs)

	return
}

// Returns which row defines the global baseline of @grid.
/*

C function : gtk_grid_get_baseline_row
*/
func (recv *Grid) GetBaselineRow() int32 {
	retC := C.gtk_grid_get_baseline_row((*C.GtkGrid)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Returns the baseline position of @row as set
// by gtk_grid_set_row_baseline_position() or the default value
// %GTK_BASELINE_POSITION_CENTER.
/*

C function : gtk_grid_get_row_baseline_position
*/
func (recv *Grid) GetRowBaselinePosition(row int32) BaselinePosition {
	c_row := (C.gint)(row)

	retC := C.gtk_grid_get_row_baseline_position((*C.GtkGrid)(recv.native), c_row)
	retGo := (BaselinePosition)(retC)

	return retGo
}

// Removes a column from the grid.
//
// Children that are placed in this column are removed,
// spanning children that overlap this column have their
// width reduced by one, and children after the column
// are moved to the left.
/*

C function : gtk_grid_remove_column
*/
func (recv *Grid) RemoveColumn(position int32) {
	c_position := (C.gint)(position)

	C.gtk_grid_remove_column((*C.GtkGrid)(recv.native), c_position)

	return
}

// Removes a row from the grid.
//
// Children that are placed in this row are removed,
// spanning children that overlap this row have their
// height reduced by one, and children below the row
// are moved up.
/*

C function : gtk_grid_remove_row
*/
func (recv *Grid) RemoveRow(position int32) {
	c_position := (C.gint)(position)

	C.gtk_grid_remove_row((*C.GtkGrid)(recv.native), c_position)

	return
}

// Sets which row defines the global baseline for the entire grid.
// Each row in the grid can have its own local baseline, but only
// one of those is global, meaning it will be the baseline in the
// parent of the @grid.
/*

C function : gtk_grid_set_baseline_row
*/
func (recv *Grid) SetBaselineRow(row int32) {
	c_row := (C.gint)(row)

	C.gtk_grid_set_baseline_row((*C.GtkGrid)(recv.native), c_row)

	return
}

// Sets how the baseline should be positioned on @row of the
// grid, in case that row is assigned more space than is requested.
/*

C function : gtk_grid_set_row_baseline_position
*/
func (recv *Grid) SetRowBaselinePosition(row int32, pos BaselinePosition) {
	c_row := (C.gint)(row)

	c_pos := (C.GtkBaselinePosition)(pos)

	C.gtk_grid_set_row_baseline_position((*C.GtkGrid)(recv.native), c_row, c_pos)

	return
}

// Creates a new #GtkHeaderBar widget.
/*

C function : gtk_header_bar_new
*/
func HeaderBarNew() *HeaderBar {
	retC := C.gtk_header_bar_new()
	retGo := HeaderBarNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Retrieves the custom title widget of the header. See
// gtk_header_bar_set_custom_title().
/*

C function : gtk_header_bar_get_custom_title
*/
func (recv *HeaderBar) GetCustomTitle() *Widget {
	retC := C.gtk_header_bar_get_custom_title((*C.GtkHeaderBar)(recv.native))
	var retGo (*Widget)
	if retC == nil {
		retGo = nil
	} else {
		retGo = WidgetNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Returns whether this header bar shows the standard window
// decorations.
/*

C function : gtk_header_bar_get_show_close_button
*/
func (recv *HeaderBar) GetShowCloseButton() bool {
	retC := C.gtk_header_bar_get_show_close_button((*C.GtkHeaderBar)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Retrieves the subtitle of the header. See gtk_header_bar_set_subtitle().
/*

C function : gtk_header_bar_get_subtitle
*/
func (recv *HeaderBar) GetSubtitle() string {
	retC := C.gtk_header_bar_get_subtitle((*C.GtkHeaderBar)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Retrieves the title of the header. See gtk_header_bar_set_title().
/*

C function : gtk_header_bar_get_title
*/
func (recv *HeaderBar) GetTitle() string {
	retC := C.gtk_header_bar_get_title((*C.GtkHeaderBar)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Adds @child to @bar, packed with reference to the
// end of the @bar.
/*

C function : gtk_header_bar_pack_end
*/
func (recv *HeaderBar) PackEnd(child *Widget) {
	c_child := (*C.GtkWidget)(C.NULL)
	if child != nil {
		c_child = (*C.GtkWidget)(child.ToC())
	}

	C.gtk_header_bar_pack_end((*C.GtkHeaderBar)(recv.native), c_child)

	return
}

// Adds @child to @bar, packed with reference to the
// start of the @bar.
/*

C function : gtk_header_bar_pack_start
*/
func (recv *HeaderBar) PackStart(child *Widget) {
	c_child := (*C.GtkWidget)(C.NULL)
	if child != nil {
		c_child = (*C.GtkWidget)(child.ToC())
	}

	C.gtk_header_bar_pack_start((*C.GtkHeaderBar)(recv.native), c_child)

	return
}

// Sets a custom title for the #GtkHeaderBar.
//
// The title should help a user identify the current view. This
// supersedes any title set by gtk_header_bar_set_title() or
// gtk_header_bar_set_subtitle(). To achieve the same style as
// the builtin title and subtitle, use the “title” and “subtitle”
// style classes.
//
// You should set the custom title to %NULL, for the header title
// label to be visible again.
/*

C function : gtk_header_bar_set_custom_title
*/
func (recv *HeaderBar) SetCustomTitle(titleWidget *Widget) {
	c_title_widget := (*C.GtkWidget)(C.NULL)
	if titleWidget != nil {
		c_title_widget = (*C.GtkWidget)(titleWidget.ToC())
	}

	C.gtk_header_bar_set_custom_title((*C.GtkHeaderBar)(recv.native), c_title_widget)

	return
}

// Sets whether this header bar shows the standard window decorations,
// including close, maximize, and minimize.
/*

C function : gtk_header_bar_set_show_close_button
*/
func (recv *HeaderBar) SetShowCloseButton(setting bool) {
	c_setting :=
		boolToGboolean(setting)

	C.gtk_header_bar_set_show_close_button((*C.GtkHeaderBar)(recv.native), c_setting)

	return
}

// Sets the subtitle of the #GtkHeaderBar. The title should give a user
// an additional detail to help him identify the current view.
//
// Note that GtkHeaderBar by default reserves room for the subtitle,
// even if none is currently set. If this is not desired, set the
// #GtkHeaderBar:has-subtitle property to %FALSE.
/*

C function : gtk_header_bar_set_subtitle
*/
func (recv *HeaderBar) SetSubtitle(subtitle string) {
	c_subtitle := C.CString(subtitle)
	defer C.free(unsafe.Pointer(c_subtitle))

	C.gtk_header_bar_set_subtitle((*C.GtkHeaderBar)(recv.native), c_subtitle)

	return
}

// Sets the title of the #GtkHeaderBar. The title should help a user
// identify the current view. A good title should not include the
// application name.
/*

C function : gtk_header_bar_set_title
*/
func (recv *HeaderBar) SetTitle(title string) {
	c_title := C.CString(title)
	defer C.free(unsafe.Pointer(c_title))

	C.gtk_header_bar_set_title((*C.GtkHeaderBar)(recv.native), c_title)

	return
}

// Gets the base scale for the icon. The base scale is a scale
// for the icon that was specified by the icon theme creator.
// For instance an icon drawn for a high-dpi screen with window
// scale 2 for a base size of 32 will be 64 pixels tall and have
// a base scale of 2.
/*

C function : gtk_icon_info_get_base_scale
*/
func (recv *IconInfo) GetBaseScale() int32 {
	retC := C.gtk_icon_info_get_base_scale((*C.GtkIconInfo)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Renders an icon previously looked up in an icon theme using
// gtk_icon_theme_lookup_icon(); the size will be based on the size
// passed to gtk_icon_theme_lookup_icon(). Note that the resulting
// surface may not be exactly this size; an icon theme may have icons
// that differ slightly from their nominal sizes, and in addition GTK+
// will avoid scaling icons that it considers sufficiently close to the
// requested size or for which the source image would have to be scaled
// up too far. (This maintains sharpness.). This behaviour can be changed
// by passing the %GTK_ICON_LOOKUP_FORCE_SIZE flag when obtaining
// the #GtkIconInfo. If this flag has been specified, the pixbuf
// returned by this function will be scaled to the exact size.
/*

C function : gtk_icon_info_load_surface
*/
func (recv *IconInfo) LoadSurface(forWindow *gdk.Window) (*cairo.Surface, error) {
	c_for_window := (*C.GdkWindow)(C.NULL)
	if forWindow != nil {
		c_for_window = (*C.GdkWindow)(forWindow.ToC())
	}

	var cThrowableError *C.GError

	retC := C.gtk_icon_info_load_surface((*C.GtkIconInfo)(recv.native), c_for_window, &cThrowableError)
	retGo := cairo.SurfaceNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : gtk_icon_theme_choose_icon_for_scale : unsupported parameter icon_names :

// Looks up an icon in an icon theme for a particular window scale,
// scales it to the given size and renders it into a pixbuf. This is a
// convenience function; if more details about the icon are needed,
// use gtk_icon_theme_lookup_icon() followed by
// gtk_icon_info_load_icon().
//
// Note that you probably want to listen for icon theme changes and
// update the icon. This is usually done by connecting to the
// GtkWidget::style-set signal. If for some reason you do not want to
// update the icon when the icon theme changes, you should consider
// using gdk_pixbuf_copy() to make a private copy of the pixbuf
// returned by this function. Otherwise GTK+ may need to keep the old
// icon theme loaded, which would be a waste of memory.
/*

C function : gtk_icon_theme_load_icon_for_scale
*/
func (recv *IconTheme) LoadIconForScale(iconName string, size int32, scale int32, flags IconLookupFlags) (*gdkpixbuf.Pixbuf, error) {
	c_icon_name := C.CString(iconName)
	defer C.free(unsafe.Pointer(c_icon_name))

	c_size := (C.gint)(size)

	c_scale := (C.gint)(scale)

	c_flags := (C.GtkIconLookupFlags)(flags)

	var cThrowableError *C.GError

	retC := C.gtk_icon_theme_load_icon_for_scale((*C.GtkIconTheme)(recv.native), c_icon_name, c_size, c_scale, c_flags, &cThrowableError)
	var retGo (*gdkpixbuf.Pixbuf)
	if retC == nil {
		retGo = nil
	} else {
		retGo = gdkpixbuf.PixbufNewFromC(unsafe.Pointer(retC))
	}

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Looks up an icon in an icon theme for a particular window scale,
// scales it to the given size and renders it into a cairo surface. This is a
// convenience function; if more details about the icon are needed,
// use gtk_icon_theme_lookup_icon() followed by
// gtk_icon_info_load_surface().
//
// Note that you probably want to listen for icon theme changes and
// update the icon. This is usually done by connecting to the
// GtkWidget::style-set signal.
/*

C function : gtk_icon_theme_load_surface
*/
func (recv *IconTheme) LoadSurface(iconName string, size int32, scale int32, forWindow *gdk.Window, flags IconLookupFlags) (*cairo.Surface, error) {
	c_icon_name := C.CString(iconName)
	defer C.free(unsafe.Pointer(c_icon_name))

	c_size := (C.gint)(size)

	c_scale := (C.gint)(scale)

	c_for_window := (*C.GdkWindow)(C.NULL)
	if forWindow != nil {
		c_for_window = (*C.GdkWindow)(forWindow.ToC())
	}

	c_flags := (C.GtkIconLookupFlags)(flags)

	var cThrowableError *C.GError

	retC := C.gtk_icon_theme_load_surface((*C.GtkIconTheme)(recv.native), c_icon_name, c_size, c_scale, c_for_window, c_flags, &cThrowableError)
	var retGo (*cairo.Surface)
	if retC == nil {
		retGo = nil
	} else {
		retGo = cairo.SurfaceNewFromC(unsafe.Pointer(retC))
	}

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Looks up an icon and returns a #GtkIconInfo containing information
// such as the filename of the icon. The icon can then be rendered into
// a pixbuf using gtk_icon_info_load_icon().
/*

C function : gtk_icon_theme_lookup_by_gicon_for_scale
*/
func (recv *IconTheme) LookupByGiconForScale(icon *gio.Icon, size int32, scale int32, flags IconLookupFlags) *IconInfo {
	c_icon := (*C.GIcon)(icon.ToC())

	c_size := (C.gint)(size)

	c_scale := (C.gint)(scale)

	c_flags := (C.GtkIconLookupFlags)(flags)

	retC := C.gtk_icon_theme_lookup_by_gicon_for_scale((*C.GtkIconTheme)(recv.native), c_icon, c_size, c_scale, c_flags)
	var retGo (*IconInfo)
	if retC == nil {
		retGo = nil
	} else {
		retGo = IconInfoNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Looks up a named icon for a particular window scale and returns a
// #GtkIconInfo containing information such as the filename of the
// icon. The icon can then be rendered into a pixbuf using
// gtk_icon_info_load_icon(). (gtk_icon_theme_load_icon() combines
// these two steps if all you need is the pixbuf.)
/*

C function : gtk_icon_theme_lookup_icon_for_scale
*/
func (recv *IconTheme) LookupIconForScale(iconName string, size int32, scale int32, flags IconLookupFlags) *IconInfo {
	c_icon_name := C.CString(iconName)
	defer C.free(unsafe.Pointer(c_icon_name))

	c_size := (C.gint)(size)

	c_scale := (C.gint)(scale)

	c_flags := (C.GtkIconLookupFlags)(flags)

	retC := C.gtk_icon_theme_lookup_icon_for_scale((*C.GtkIconTheme)(recv.native), c_icon_name, c_size, c_scale, c_flags)
	var retGo (*IconInfo)
	if retC == nil {
		retGo = nil
	} else {
		retGo = IconInfoNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Creates a new #GtkImage displaying @surface.
// The #GtkImage does not assume a reference to the
// surface; you still need to unref it if you own references.
// #GtkImage will add its own reference rather than adopting yours.
/*

C function : gtk_image_new_from_surface
*/
func ImageNewFromSurface(surface *cairo.Surface) *Image {
	c_surface := (*C.cairo_surface_t)(C.NULL)
	if surface != nil {
		c_surface = (*C.cairo_surface_t)(surface.ToC())
	}

	retC := C.gtk_image_new_from_surface(c_surface)
	retGo := ImageNewFromC(unsafe.Pointer(retC))

	return retGo
}

// See gtk_image_new_from_surface() for details.
/*

C function : gtk_image_set_from_surface
*/
func (recv *Image) SetFromSurface(surface *cairo.Surface) {
	c_surface := (*C.cairo_surface_t)(C.NULL)
	if surface != nil {
		c_surface = (*C.cairo_surface_t)(surface.ToC())
	}

	C.gtk_image_set_from_surface((*C.GtkImage)(recv.native), c_surface)

	return
}

// Returns whether the widget will display a standard close button.
/*

C function : gtk_info_bar_get_show_close_button
*/
func (recv *InfoBar) GetShowCloseButton() bool {
	retC := C.gtk_info_bar_get_show_close_button((*C.GtkInfoBar)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// If true, a standard close button is shown. When clicked it emits
// the response %GTK_RESPONSE_CLOSE.
/*

C function : gtk_info_bar_set_show_close_button
*/
func (recv *InfoBar) SetShowCloseButton(setting bool) {
	c_setting :=
		boolToGboolean(setting)

	C.gtk_info_bar_set_show_close_button((*C.GtkInfoBar)(recv.native), c_setting)

	return
}

// Gets the number of lines to which an ellipsized, wrapping
// label should be limited. See gtk_label_set_lines().
/*

C function : gtk_label_get_lines
*/
func (recv *Label) GetLines() int32 {
	retC := C.gtk_label_get_lines((*C.GtkLabel)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Sets the number of lines to which an ellipsized, wrapping label
// should be limited. This has no effect if the label is not wrapping
// or ellipsized. Set this to -1 if you don’t want to limit the
// number of lines.
/*

C function : gtk_label_set_lines
*/
func (recv *Label) SetLines(lines int32) {
	c_lines := (C.gint)(lines)

	C.gtk_label_set_lines((*C.GtkLabel)(recv.native), c_lines)

	return
}

type signalListBoxRowActivatedDetail struct {
	callback  ListBoxSignalRowActivatedCallback
	handlerID C.gulong
}

var signalListBoxRowActivatedId int
var signalListBoxRowActivatedMap = make(map[int]signalListBoxRowActivatedDetail)
var signalListBoxRowActivatedLock sync.Mutex

// ListBoxSignalRowActivatedCallback is a callback function for a 'row-activated' signal emitted from a ListBox.
type ListBoxSignalRowActivatedCallback func(row *ListBoxRow)

/*
ConnectRowActivated connects the callback to the 'row-activated' signal for the ListBox.

The returned value represents the connection, and may be passed to DisconnectRowActivated to remove it.
*/
func (recv *ListBox) ConnectRowActivated(callback ListBoxSignalRowActivatedCallback) int {
	signalListBoxRowActivatedLock.Lock()
	defer signalListBoxRowActivatedLock.Unlock()

	signalListBoxRowActivatedId++
	instance := C.gpointer(recv.native)
	handlerID := C.ListBox_signal_connect_row_activated(instance, C.gpointer(uintptr(signalListBoxRowActivatedId)))

	detail := signalListBoxRowActivatedDetail{callback, handlerID}
	signalListBoxRowActivatedMap[signalListBoxRowActivatedId] = detail

	return signalListBoxRowActivatedId
}

/*
DisconnectRowActivated disconnects a callback from the 'row-activated' signal for the ListBox.

The connectionID should be a value returned from a call to ConnectRowActivated.
*/
func (recv *ListBox) DisconnectRowActivated(connectionID int) {
	signalListBoxRowActivatedLock.Lock()
	defer signalListBoxRowActivatedLock.Unlock()

	detail, exists := signalListBoxRowActivatedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalListBoxRowActivatedMap, connectionID)
}

//export listbox_rowActivatedHandler
func listbox_rowActivatedHandler(_ *C.GObject, c_row *C.GtkListBoxRow, data C.gpointer) {
	row := ListBoxRowNewFromC(unsafe.Pointer(c_row))

	index := int(uintptr(data))
	callback := signalListBoxRowActivatedMap[index].callback
	callback(row)
}

type signalListBoxRowSelectedDetail struct {
	callback  ListBoxSignalRowSelectedCallback
	handlerID C.gulong
}

var signalListBoxRowSelectedId int
var signalListBoxRowSelectedMap = make(map[int]signalListBoxRowSelectedDetail)
var signalListBoxRowSelectedLock sync.Mutex

// ListBoxSignalRowSelectedCallback is a callback function for a 'row-selected' signal emitted from a ListBox.
type ListBoxSignalRowSelectedCallback func(row *ListBoxRow)

/*
ConnectRowSelected connects the callback to the 'row-selected' signal for the ListBox.

The returned value represents the connection, and may be passed to DisconnectRowSelected to remove it.
*/
func (recv *ListBox) ConnectRowSelected(callback ListBoxSignalRowSelectedCallback) int {
	signalListBoxRowSelectedLock.Lock()
	defer signalListBoxRowSelectedLock.Unlock()

	signalListBoxRowSelectedId++
	instance := C.gpointer(recv.native)
	handlerID := C.ListBox_signal_connect_row_selected(instance, C.gpointer(uintptr(signalListBoxRowSelectedId)))

	detail := signalListBoxRowSelectedDetail{callback, handlerID}
	signalListBoxRowSelectedMap[signalListBoxRowSelectedId] = detail

	return signalListBoxRowSelectedId
}

/*
DisconnectRowSelected disconnects a callback from the 'row-selected' signal for the ListBox.

The connectionID should be a value returned from a call to ConnectRowSelected.
*/
func (recv *ListBox) DisconnectRowSelected(connectionID int) {
	signalListBoxRowSelectedLock.Lock()
	defer signalListBoxRowSelectedLock.Unlock()

	detail, exists := signalListBoxRowSelectedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalListBoxRowSelectedMap, connectionID)
}

//export listbox_rowSelectedHandler
func listbox_rowSelectedHandler(_ *C.GObject, c_row *C.GtkListBoxRow, data C.gpointer) {
	row := ListBoxRowNewFromC(unsafe.Pointer(c_row))

	index := int(uintptr(data))
	callback := signalListBoxRowSelectedMap[index].callback
	callback(row)
}

// Creates a new #GtkListBox container.
/*

C function : gtk_list_box_new
*/
func ListBoxNew() *ListBox {
	retC := C.gtk_list_box_new()
	retGo := ListBoxNewFromC(unsafe.Pointer(retC))

	return retGo
}

// This is a helper function for implementing DnD onto a #GtkListBox.
// The passed in @row will be highlighted via gtk_drag_highlight(),
// and any previously highlighted row will be unhighlighted.
//
// The row will also be unhighlighted when the widget gets
// a drag leave event.
/*

C function : gtk_list_box_drag_highlight_row
*/
func (recv *ListBox) DragHighlightRow(row *ListBoxRow) {
	c_row := (*C.GtkListBoxRow)(C.NULL)
	if row != nil {
		c_row = (*C.GtkListBoxRow)(row.ToC())
	}

	C.gtk_list_box_drag_highlight_row((*C.GtkListBox)(recv.native), c_row)

	return
}

// If a row has previously been highlighted via gtk_list_box_drag_highlight_row()
// it will have the highlight removed.
/*

C function : gtk_list_box_drag_unhighlight_row
*/
func (recv *ListBox) DragUnhighlightRow() {
	C.gtk_list_box_drag_unhighlight_row((*C.GtkListBox)(recv.native))

	return
}

// Returns whether rows activate on single clicks.
/*

C function : gtk_list_box_get_activate_on_single_click
*/
func (recv *ListBox) GetActivateOnSingleClick() bool {
	retC := C.gtk_list_box_get_activate_on_single_click((*C.GtkListBox)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Gets the adjustment (if any) that the widget uses to
// for vertical scrolling.
/*

C function : gtk_list_box_get_adjustment
*/
func (recv *ListBox) GetAdjustment() *Adjustment {
	retC := C.gtk_list_box_get_adjustment((*C.GtkListBox)(recv.native))
	retGo := AdjustmentNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the n-th child in the list (not counting headers).
// If @_index is negative or larger than the number of items in the
// list, %NULL is returned.
/*

C function : gtk_list_box_get_row_at_index
*/
func (recv *ListBox) GetRowAtIndex(index int32) *ListBoxRow {
	c_index_ := (C.gint)(index)

	retC := C.gtk_list_box_get_row_at_index((*C.GtkListBox)(recv.native), c_index_)
	var retGo (*ListBoxRow)
	if retC == nil {
		retGo = nil
	} else {
		retGo = ListBoxRowNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Gets the row at the @y position.
/*

C function : gtk_list_box_get_row_at_y
*/
func (recv *ListBox) GetRowAtY(y int32) *ListBoxRow {
	c_y := (C.gint)(y)

	retC := C.gtk_list_box_get_row_at_y((*C.GtkListBox)(recv.native), c_y)
	var retGo (*ListBoxRow)
	if retC == nil {
		retGo = nil
	} else {
		retGo = ListBoxRowNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Gets the selected row.
//
// Note that the box may allow multiple selection, in which
// case you should use gtk_list_box_selected_foreach() to
// find all selected rows.
/*

C function : gtk_list_box_get_selected_row
*/
func (recv *ListBox) GetSelectedRow() *ListBoxRow {
	retC := C.gtk_list_box_get_selected_row((*C.GtkListBox)(recv.native))
	retGo := ListBoxRowNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the selection mode of the listbox.
/*

C function : gtk_list_box_get_selection_mode
*/
func (recv *ListBox) GetSelectionMode() SelectionMode {
	retC := C.gtk_list_box_get_selection_mode((*C.GtkListBox)(recv.native))
	retGo := (SelectionMode)(retC)

	return retGo
}

// Insert the @child into the @box at @position. If a sort function is
// set, the widget will actually be inserted at the calculated position and
// this function has the same effect of gtk_container_add().
//
// If @position is -1, or larger than the total number of items in the
// @box, then the @child will be appended to the end.
/*

C function : gtk_list_box_insert
*/
func (recv *ListBox) Insert(child *Widget, position int32) {
	c_child := (*C.GtkWidget)(C.NULL)
	if child != nil {
		c_child = (*C.GtkWidget)(child.ToC())
	}

	c_position := (C.gint)(position)

	C.gtk_list_box_insert((*C.GtkListBox)(recv.native), c_child, c_position)

	return
}

// Update the filtering for all rows. Call this when result
// of the filter function on the @box is changed due
// to an external factor. For instance, this would be used
// if the filter function just looked for a specific search
// string and the entry with the search string has changed.
/*

C function : gtk_list_box_invalidate_filter
*/
func (recv *ListBox) InvalidateFilter() {
	C.gtk_list_box_invalidate_filter((*C.GtkListBox)(recv.native))

	return
}

// Update the separators for all rows. Call this when result
// of the header function on the @box is changed due
// to an external factor.
/*

C function : gtk_list_box_invalidate_headers
*/
func (recv *ListBox) InvalidateHeaders() {
	C.gtk_list_box_invalidate_headers((*C.GtkListBox)(recv.native))

	return
}

// Update the sorting for all rows. Call this when result
// of the sort function on the @box is changed due
// to an external factor.
/*

C function : gtk_list_box_invalidate_sort
*/
func (recv *ListBox) InvalidateSort() {
	C.gtk_list_box_invalidate_sort((*C.GtkListBox)(recv.native))

	return
}

// Prepend a widget to the list. If a sort function is set, the widget will
// actually be inserted at the calculated position and this function has the
// same effect of gtk_container_add().
/*

C function : gtk_list_box_prepend
*/
func (recv *ListBox) Prepend(child *Widget) {
	c_child := (*C.GtkWidget)(C.NULL)
	if child != nil {
		c_child = (*C.GtkWidget)(child.ToC())
	}

	C.gtk_list_box_prepend((*C.GtkListBox)(recv.native), c_child)

	return
}

// Make @row the currently selected row.
/*

C function : gtk_list_box_select_row
*/
func (recv *ListBox) SelectRow(row *ListBoxRow) {
	c_row := (*C.GtkListBoxRow)(C.NULL)
	if row != nil {
		c_row = (*C.GtkListBoxRow)(row.ToC())
	}

	C.gtk_list_box_select_row((*C.GtkListBox)(recv.native), c_row)

	return
}

// If @single is %TRUE, rows will be activated when you click on them,
// otherwise you need to double-click.
/*

C function : gtk_list_box_set_activate_on_single_click
*/
func (recv *ListBox) SetActivateOnSingleClick(single bool) {
	c_single :=
		boolToGboolean(single)

	C.gtk_list_box_set_activate_on_single_click((*C.GtkListBox)(recv.native), c_single)

	return
}

// Sets the adjustment (if any) that the widget uses to
// for vertical scrolling. For instance, this is used
// to get the page size for PageUp/Down key handling.
//
// In the normal case when the @box is packed inside
// a #GtkScrolledWindow the adjustment from that will
// be picked up automatically, so there is no need
// to manually do that.
/*

C function : gtk_list_box_set_adjustment
*/
func (recv *ListBox) SetAdjustment(adjustment *Adjustment) {
	c_adjustment := (*C.GtkAdjustment)(C.NULL)
	if adjustment != nil {
		c_adjustment = (*C.GtkAdjustment)(adjustment.ToC())
	}

	C.gtk_list_box_set_adjustment((*C.GtkListBox)(recv.native), c_adjustment)

	return
}

// Unsupported : gtk_list_box_set_filter_func : unsupported parameter filter_func : no type generator for ListBoxFilterFunc (GtkListBoxFilterFunc) for param filter_func

// Unsupported : gtk_list_box_set_header_func : unsupported parameter update_header : no type generator for ListBoxUpdateHeaderFunc (GtkListBoxUpdateHeaderFunc) for param update_header

// Sets the placeholder widget that is shown in the list when
// it doesn't display any visible children.
/*

C function : gtk_list_box_set_placeholder
*/
func (recv *ListBox) SetPlaceholder(placeholder *Widget) {
	c_placeholder := (*C.GtkWidget)(C.NULL)
	if placeholder != nil {
		c_placeholder = (*C.GtkWidget)(placeholder.ToC())
	}

	C.gtk_list_box_set_placeholder((*C.GtkListBox)(recv.native), c_placeholder)

	return
}

// Sets how selection works in the listbox.
// See #GtkSelectionMode for details.
/*

C function : gtk_list_box_set_selection_mode
*/
func (recv *ListBox) SetSelectionMode(mode SelectionMode) {
	c_mode := (C.GtkSelectionMode)(mode)

	C.gtk_list_box_set_selection_mode((*C.GtkListBox)(recv.native), c_mode)

	return
}

// Unsupported : gtk_list_box_set_sort_func : unsupported parameter sort_func : no type generator for ListBoxSortFunc (GtkListBoxSortFunc) for param sort_func

type signalListBoxRowActivateDetail struct {
	callback  ListBoxRowSignalActivateCallback
	handlerID C.gulong
}

var signalListBoxRowActivateId int
var signalListBoxRowActivateMap = make(map[int]signalListBoxRowActivateDetail)
var signalListBoxRowActivateLock sync.Mutex

// ListBoxRowSignalActivateCallback is a callback function for a 'activate' signal emitted from a ListBoxRow.
type ListBoxRowSignalActivateCallback func()

/*
ConnectActivate connects the callback to the 'activate' signal for the ListBoxRow.

The returned value represents the connection, and may be passed to DisconnectActivate to remove it.
*/
func (recv *ListBoxRow) ConnectActivate(callback ListBoxRowSignalActivateCallback) int {
	signalListBoxRowActivateLock.Lock()
	defer signalListBoxRowActivateLock.Unlock()

	signalListBoxRowActivateId++
	instance := C.gpointer(recv.native)
	handlerID := C.ListBoxRow_signal_connect_activate(instance, C.gpointer(uintptr(signalListBoxRowActivateId)))

	detail := signalListBoxRowActivateDetail{callback, handlerID}
	signalListBoxRowActivateMap[signalListBoxRowActivateId] = detail

	return signalListBoxRowActivateId
}

/*
DisconnectActivate disconnects a callback from the 'activate' signal for the ListBoxRow.

The connectionID should be a value returned from a call to ConnectActivate.
*/
func (recv *ListBoxRow) DisconnectActivate(connectionID int) {
	signalListBoxRowActivateLock.Lock()
	defer signalListBoxRowActivateLock.Unlock()

	detail, exists := signalListBoxRowActivateMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalListBoxRowActivateMap, connectionID)
}

//export listboxrow_activateHandler
func listboxrow_activateHandler(_ *C.GObject, data C.gpointer) {
	index := int(uintptr(data))
	callback := signalListBoxRowActivateMap[index].callback
	callback()
}

// Creates a new #GtkListBoxRow, to be used as a child of a #GtkListBox.
/*

C function : gtk_list_box_row_new
*/
func ListBoxRowNew() *ListBoxRow {
	retC := C.gtk_list_box_row_new()
	retGo := ListBoxRowNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Marks @row as changed, causing any state that depends on this
// to be updated. This affects sorting, filtering and headers.
//
// Note that calls to this method must be in sync with the data
// used for the row functions. For instance, if the list is
// mirroring some external data set, and *two* rows changed in the
// external data set then when you call gtk_list_box_row_changed()
// on the first row the sort function must only read the new data
// for the first of the two changed rows, otherwise the resorting
// of the rows will be wrong.
//
// This generally means that if you don’t fully control the data
// model you have to duplicate the data that affects the listbox
// row functions into the row widgets themselves. Another alternative
// is to call gtk_list_box_invalidate_sort() on any model change,
// but that is more expensive.
/*

C function : gtk_list_box_row_changed
*/
func (recv *ListBoxRow) Changed() {
	C.gtk_list_box_row_changed((*C.GtkListBoxRow)(recv.native))

	return
}

// Returns the current header of the @row. This can be used
// in a #GtkListBoxUpdateHeaderFunc to see if there is a header
// set already, and if so to update the state of it.
/*

C function : gtk_list_box_row_get_header
*/
func (recv *ListBoxRow) GetHeader() *Widget {
	retC := C.gtk_list_box_row_get_header((*C.GtkListBoxRow)(recv.native))
	var retGo (*Widget)
	if retC == nil {
		retGo = nil
	} else {
		retGo = WidgetNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Gets the current index of the @row in its #GtkListBox container.
/*

C function : gtk_list_box_row_get_index
*/
func (recv *ListBoxRow) GetIndex() int32 {
	retC := C.gtk_list_box_row_get_index((*C.GtkListBoxRow)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Sets the current header of the @row. This is only allowed to be called
// from a #GtkListBoxUpdateHeaderFunc. It will replace any existing
// header in the row, and be shown in front of the row in the listbox.
/*

C function : gtk_list_box_row_set_header
*/
func (recv *ListBoxRow) SetHeader(header *Widget) {
	c_header := (*C.GtkWidget)(C.NULL)
	if header != nil {
		c_header = (*C.GtkWidget)(header.ToC())
	}

	C.gtk_list_box_row_set_header((*C.GtkListBoxRow)(recv.native), c_header)

	return
}

// Unsupported signal 'drag-action-ask' for PlacesSidebar : unsupported parameter actions : type gint :

// Unsupported signal 'drag-action-requested' for PlacesSidebar : unsupported parameter source_file_list : type GLib.List :

// Unsupported signal 'drag-perform-drop' for PlacesSidebar : unsupported parameter source_file_list : type GLib.List :

// Unsupported signal 'open-location' for PlacesSidebar : unsupported parameter open_flags : type PlacesOpenFlags :

type signalPlacesSidebarPopulatePopupDetail struct {
	callback  PlacesSidebarSignalPopulatePopupCallback
	handlerID C.gulong
}

var signalPlacesSidebarPopulatePopupId int
var signalPlacesSidebarPopulatePopupMap = make(map[int]signalPlacesSidebarPopulatePopupDetail)
var signalPlacesSidebarPopulatePopupLock sync.Mutex

// PlacesSidebarSignalPopulatePopupCallback is a callback function for a 'populate-popup' signal emitted from a PlacesSidebar.
type PlacesSidebarSignalPopulatePopupCallback func(container *Widget, selectedItem *gio.File, selectedVolume *gio.Volume)

/*
ConnectPopulatePopup connects the callback to the 'populate-popup' signal for the PlacesSidebar.

The returned value represents the connection, and may be passed to DisconnectPopulatePopup to remove it.
*/
func (recv *PlacesSidebar) ConnectPopulatePopup(callback PlacesSidebarSignalPopulatePopupCallback) int {
	signalPlacesSidebarPopulatePopupLock.Lock()
	defer signalPlacesSidebarPopulatePopupLock.Unlock()

	signalPlacesSidebarPopulatePopupId++
	instance := C.gpointer(recv.native)
	handlerID := C.PlacesSidebar_signal_connect_populate_popup(instance, C.gpointer(uintptr(signalPlacesSidebarPopulatePopupId)))

	detail := signalPlacesSidebarPopulatePopupDetail{callback, handlerID}
	signalPlacesSidebarPopulatePopupMap[signalPlacesSidebarPopulatePopupId] = detail

	return signalPlacesSidebarPopulatePopupId
}

/*
DisconnectPopulatePopup disconnects a callback from the 'populate-popup' signal for the PlacesSidebar.

The connectionID should be a value returned from a call to ConnectPopulatePopup.
*/
func (recv *PlacesSidebar) DisconnectPopulatePopup(connectionID int) {
	signalPlacesSidebarPopulatePopupLock.Lock()
	defer signalPlacesSidebarPopulatePopupLock.Unlock()

	detail, exists := signalPlacesSidebarPopulatePopupMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalPlacesSidebarPopulatePopupMap, connectionID)
}

//export placessidebar_populatePopupHandler
func placessidebar_populatePopupHandler(_ *C.GObject, c_container *C.GtkWidget, c_selected_item *C.GFile, c_selected_volume *C.GVolume, data C.gpointer) {
	container := WidgetNewFromC(unsafe.Pointer(c_container))

	selectedItem := gio.FileNewFromC(unsafe.Pointer(c_selected_item))

	selectedVolume := gio.VolumeNewFromC(unsafe.Pointer(c_selected_volume))

	index := int(uintptr(data))
	callback := signalPlacesSidebarPopulatePopupMap[index].callback
	callback(container, selectedItem, selectedVolume)
}

// Unsupported signal 'show-error-message' for PlacesSidebar : unsupported parameter primary : type utf8 :

// Creates a new #GtkPlacesSidebar widget.
//
// The application should connect to at least the
// #GtkPlacesSidebar::open-location signal to be notified
// when the user makes a selection in the sidebar.
/*

C function : gtk_places_sidebar_new
*/
func PlacesSidebarNew() *PlacesSidebar {
	retC := C.gtk_places_sidebar_new()
	retGo := PlacesSidebarNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Applications may want to present some folders in the places sidebar if
// they could be immediately useful to users.  For example, a drawing
// program could add a “/usr/share/clipart” location when the sidebar is
// being used in an “Insert Clipart” dialog box.
//
// This function adds the specified @location to a special place for immutable
// shortcuts.  The shortcuts are application-specific; they are not shared
// across applications, and they are not persistent.  If this function
// is called multiple times with different locations, then they are added
// to the sidebar’s list in the same order as the function is called.
/*

C function : gtk_places_sidebar_add_shortcut
*/
func (recv *PlacesSidebar) AddShortcut(location *gio.File) {
	c_location := (*C.GFile)(location.ToC())

	C.gtk_places_sidebar_add_shortcut((*C.GtkPlacesSidebar)(recv.native), c_location)

	return
}

// Gets the currently selected location in the @sidebar. This can be %NULL when
// nothing is selected, for example, when gtk_places_sidebar_set_location() has
// been called with a location that is not among the sidebar’s list of places to
// show.
//
// You can use this function to get the selection in the @sidebar.  Also, if you
// connect to the #GtkPlacesSidebar::populate-popup signal, you can use this
// function to get the location that is being referred to during the callbacks
// for your menu items.
/*

C function : gtk_places_sidebar_get_location
*/
func (recv *PlacesSidebar) GetLocation() *gio.File {
	retC := C.gtk_places_sidebar_get_location((*C.GtkPlacesSidebar)(recv.native))
	var retGo (*gio.File)
	if retC == nil {
		retGo = nil
	} else {
		retGo = gio.FileNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// This function queries the bookmarks added by the user to the places sidebar,
// and returns one of them.  This function is used by #GtkFileChooser to implement
// the “Alt-1”, “Alt-2”, etc. shortcuts, which activate the cooresponding bookmark.
/*

C function : gtk_places_sidebar_get_nth_bookmark
*/
func (recv *PlacesSidebar) GetNthBookmark(n int32) *gio.File {
	c_n := (C.gint)(n)

	retC := C.gtk_places_sidebar_get_nth_bookmark((*C.GtkPlacesSidebar)(recv.native), c_n)
	var retGo (*gio.File)
	if retC == nil {
		retGo = nil
	} else {
		retGo = gio.FileNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Gets the open flags.
/*

C function : gtk_places_sidebar_get_open_flags
*/
func (recv *PlacesSidebar) GetOpenFlags() PlacesOpenFlags {
	retC := C.gtk_places_sidebar_get_open_flags((*C.GtkPlacesSidebar)(recv.native))
	retGo := (PlacesOpenFlags)(retC)

	return retGo
}

// Returns the value previously set with gtk_places_sidebar_set_show_desktop()
/*

C function : gtk_places_sidebar_get_show_desktop
*/
func (recv *PlacesSidebar) GetShowDesktop() bool {
	retC := C.gtk_places_sidebar_get_show_desktop((*C.GtkPlacesSidebar)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Gets the list of shortcuts.
/*

C function : gtk_places_sidebar_list_shortcuts
*/
func (recv *PlacesSidebar) ListShortcuts() *glib.SList {
	retC := C.gtk_places_sidebar_list_shortcuts((*C.GtkPlacesSidebar)(recv.native))
	retGo := glib.SListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Removes an application-specific shortcut that has been previously been
// inserted with gtk_places_sidebar_add_shortcut().  If the @location is not a
// shortcut in the sidebar, then nothing is done.
/*

C function : gtk_places_sidebar_remove_shortcut
*/
func (recv *PlacesSidebar) RemoveShortcut(location *gio.File) {
	c_location := (*C.GFile)(location.ToC())

	C.gtk_places_sidebar_remove_shortcut((*C.GtkPlacesSidebar)(recv.native), c_location)

	return
}

// Sets the location that is being shown in the widgets surrounding the
// @sidebar, for example, in a folder view in a file manager.  In turn, the
// @sidebar will highlight that location if it is being shown in the list of
// places, or it will unhighlight everything if the @location is not among the
// places in the list.
/*

C function : gtk_places_sidebar_set_location
*/
func (recv *PlacesSidebar) SetLocation(location *gio.File) {
	c_location := (*C.GFile)(location.ToC())

	C.gtk_places_sidebar_set_location((*C.GtkPlacesSidebar)(recv.native), c_location)

	return
}

// Sets the way in which the calling application can open new locations from
// the places sidebar.  For example, some applications only open locations
// “directly” into their main view, while others may support opening locations
// in a new notebook tab or a new window.
//
// This function is used to tell the places @sidebar about the ways in which the
// application can open new locations, so that the sidebar can display (or not)
// the “Open in new tab” and “Open in new window” menu items as appropriate.
//
// When the #GtkPlacesSidebar::open-location signal is emitted, its flags
// argument will be set to one of the @flags that was passed in
// gtk_places_sidebar_set_open_flags().
//
// Passing 0 for @flags will cause #GTK_PLACES_OPEN_NORMAL to always be sent
// to callbacks for the “open-location” signal.
/*

C function : gtk_places_sidebar_set_open_flags
*/
func (recv *PlacesSidebar) SetOpenFlags(flags PlacesOpenFlags) {
	c_flags := (C.GtkPlacesOpenFlags)(flags)

	C.gtk_places_sidebar_set_open_flags((*C.GtkPlacesSidebar)(recv.native), c_flags)

	return
}

// Sets whether the @sidebar should show an item for connecting to a network server;
// this is off by default. An application may want to turn this on if it implements
// a way for the user to connect to network servers directly.
//
// If you enable this, you should connect to the
// #GtkPlacesSidebar::show-connect-to-server signal.
/*

C function : gtk_places_sidebar_set_show_connect_to_server
*/
func (recv *PlacesSidebar) SetShowConnectToServer(showConnectToServer bool) {
	c_show_connect_to_server :=
		boolToGboolean(showConnectToServer)

	C.gtk_places_sidebar_set_show_connect_to_server((*C.GtkPlacesSidebar)(recv.native), c_show_connect_to_server)

	return
}

// Sets whether the @sidebar should show an item for the Desktop folder.
// The default value for this option is determined by the desktop
// environment and the user’s configuration, but this function can be
// used to override it on a per-application basis.
/*

C function : gtk_places_sidebar_set_show_desktop
*/
func (recv *PlacesSidebar) SetShowDesktop(showDesktop bool) {
	c_show_desktop :=
		boolToGboolean(showDesktop)

	C.gtk_places_sidebar_set_show_desktop((*C.GtkPlacesSidebar)(recv.native), c_show_desktop)

	return
}

// Creates a new #GtkRevealer.
/*

C function : gtk_revealer_new
*/
func RevealerNew() *Revealer {
	retC := C.gtk_revealer_new()
	retGo := RevealerNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Returns whether the child is fully revealed, in other words whether
// the transition to the revealed state is completed.
/*

C function : gtk_revealer_get_child_revealed
*/
func (recv *Revealer) GetChildRevealed() bool {
	retC := C.gtk_revealer_get_child_revealed((*C.GtkRevealer)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Returns whether the child is currently
// revealed. See gtk_revealer_set_reveal_child().
//
// This function returns %TRUE as soon as the transition
// is to the revealed state is started. To learn whether
// the child is fully revealed (ie the transition is completed),
// use gtk_revealer_get_child_revealed().
/*

C function : gtk_revealer_get_reveal_child
*/
func (recv *Revealer) GetRevealChild() bool {
	retC := C.gtk_revealer_get_reveal_child((*C.GtkRevealer)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Returns the amount of time (in milliseconds) that
// transitions will take.
/*

C function : gtk_revealer_get_transition_duration
*/
func (recv *Revealer) GetTransitionDuration() uint32 {
	retC := C.gtk_revealer_get_transition_duration((*C.GtkRevealer)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// Gets the type of animation that will be used
// for transitions in @revealer.
/*

C function : gtk_revealer_get_transition_type
*/
func (recv *Revealer) GetTransitionType() RevealerTransitionType {
	retC := C.gtk_revealer_get_transition_type((*C.GtkRevealer)(recv.native))
	retGo := (RevealerTransitionType)(retC)

	return retGo
}

// Tells the #GtkRevealer to reveal or conceal its child.
//
// The transition will be animated with the current
// transition type of @revealer.
/*

C function : gtk_revealer_set_reveal_child
*/
func (recv *Revealer) SetRevealChild(revealChild bool) {
	c_reveal_child :=
		boolToGboolean(revealChild)

	C.gtk_revealer_set_reveal_child((*C.GtkRevealer)(recv.native), c_reveal_child)

	return
}

// Sets the duration that transitions will take.
/*

C function : gtk_revealer_set_transition_duration
*/
func (recv *Revealer) SetTransitionDuration(duration uint32) {
	c_duration := (C.guint)(duration)

	C.gtk_revealer_set_transition_duration((*C.GtkRevealer)(recv.native), c_duration)

	return
}

// Sets the type of animation that will be used for
// transitions in @revealer. Available types include
// various kinds of fades and slides.
/*

C function : gtk_revealer_set_transition_type
*/
func (recv *Revealer) SetTransitionType(transition RevealerTransitionType) {
	c_transition := (C.GtkRevealerTransitionType)(transition)

	C.gtk_revealer_set_transition_type((*C.GtkRevealer)(recv.native), c_transition)

	return
}

// Creates a #GtkSearchBar. You will need to tell it about
// which widget is going to be your text entry using
// gtk_search_bar_connect_entry().
/*

C function : gtk_search_bar_new
*/
func SearchBarNew() *SearchBar {
	retC := C.gtk_search_bar_new()
	retGo := SearchBarNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Connects the #GtkEntry widget passed as the one to be used in
// this search bar. The entry should be a descendant of the search bar.
// This is only required if the entry isn’t the direct child of the
// search bar (as in our main example).
/*

C function : gtk_search_bar_connect_entry
*/
func (recv *SearchBar) ConnectEntry(entry *Entry) {
	c_entry := (*C.GtkEntry)(C.NULL)
	if entry != nil {
		c_entry = (*C.GtkEntry)(entry.ToC())
	}

	C.gtk_search_bar_connect_entry((*C.GtkSearchBar)(recv.native), c_entry)

	return
}

// Returns whether the search mode is on or off.
/*

C function : gtk_search_bar_get_search_mode
*/
func (recv *SearchBar) GetSearchMode() bool {
	retC := C.gtk_search_bar_get_search_mode((*C.GtkSearchBar)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Returns whether the close button is shown.
/*

C function : gtk_search_bar_get_show_close_button
*/
func (recv *SearchBar) GetShowCloseButton() bool {
	retC := C.gtk_search_bar_get_show_close_button((*C.GtkSearchBar)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_search_bar_handle_event : unsupported parameter event : no type generator for Gdk.Event (GdkEvent*) for param event

// Switches the search mode on or off.
/*

C function : gtk_search_bar_set_search_mode
*/
func (recv *SearchBar) SetSearchMode(searchMode bool) {
	c_search_mode :=
		boolToGboolean(searchMode)

	C.gtk_search_bar_set_search_mode((*C.GtkSearchBar)(recv.native), c_search_mode)

	return
}

// Shows or hides the close button. Applications that
// already have a “search” toggle button should not show a close
// button in their search bar, as it duplicates the role of the
// toggle button.
/*

C function : gtk_search_bar_set_show_close_button
*/
func (recv *SearchBar) SetShowCloseButton(visible bool) {
	c_visible :=
		boolToGboolean(visible)

	C.gtk_search_bar_set_show_close_button((*C.GtkSearchBar)(recv.native), c_visible)

	return
}

type signalSearchEntrySearchChangedDetail struct {
	callback  SearchEntrySignalSearchChangedCallback
	handlerID C.gulong
}

var signalSearchEntrySearchChangedId int
var signalSearchEntrySearchChangedMap = make(map[int]signalSearchEntrySearchChangedDetail)
var signalSearchEntrySearchChangedLock sync.Mutex

// SearchEntrySignalSearchChangedCallback is a callback function for a 'search-changed' signal emitted from a SearchEntry.
type SearchEntrySignalSearchChangedCallback func()

/*
ConnectSearchChanged connects the callback to the 'search-changed' signal for the SearchEntry.

The returned value represents the connection, and may be passed to DisconnectSearchChanged to remove it.
*/
func (recv *SearchEntry) ConnectSearchChanged(callback SearchEntrySignalSearchChangedCallback) int {
	signalSearchEntrySearchChangedLock.Lock()
	defer signalSearchEntrySearchChangedLock.Unlock()

	signalSearchEntrySearchChangedId++
	instance := C.gpointer(recv.native)
	handlerID := C.SearchEntry_signal_connect_search_changed(instance, C.gpointer(uintptr(signalSearchEntrySearchChangedId)))

	detail := signalSearchEntrySearchChangedDetail{callback, handlerID}
	signalSearchEntrySearchChangedMap[signalSearchEntrySearchChangedId] = detail

	return signalSearchEntrySearchChangedId
}

/*
DisconnectSearchChanged disconnects a callback from the 'search-changed' signal for the SearchEntry.

The connectionID should be a value returned from a call to ConnectSearchChanged.
*/
func (recv *SearchEntry) DisconnectSearchChanged(connectionID int) {
	signalSearchEntrySearchChangedLock.Lock()
	defer signalSearchEntrySearchChangedLock.Unlock()

	detail, exists := signalSearchEntrySearchChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalSearchEntrySearchChangedMap, connectionID)
}

//export searchentry_searchChangedHandler
func searchentry_searchChangedHandler(_ *C.GObject, data C.gpointer) {
	index := int(uintptr(data))
	callback := signalSearchEntrySearchChangedMap[index].callback
	callback()
}

// Creates a new #GtkStack container.
/*

C function : gtk_stack_new
*/
func StackNew() *Stack {
	retC := C.gtk_stack_new()
	retGo := StackNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Adds a child to @stack.
// The child is identified by the @name.
/*

C function : gtk_stack_add_named
*/
func (recv *Stack) AddNamed(child *Widget, name string) {
	c_child := (*C.GtkWidget)(C.NULL)
	if child != nil {
		c_child = (*C.GtkWidget)(child.ToC())
	}

	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	C.gtk_stack_add_named((*C.GtkStack)(recv.native), c_child, c_name)

	return
}

// Adds a child to @stack.
// The child is identified by the @name. The @title
// will be used by #GtkStackSwitcher to represent
// @child in a tab bar, so it should be short.
/*

C function : gtk_stack_add_titled
*/
func (recv *Stack) AddTitled(child *Widget, name string, title string) {
	c_child := (*C.GtkWidget)(C.NULL)
	if child != nil {
		c_child = (*C.GtkWidget)(child.ToC())
	}

	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_title := C.CString(title)
	defer C.free(unsafe.Pointer(c_title))

	C.gtk_stack_add_titled((*C.GtkStack)(recv.native), c_child, c_name, c_title)

	return
}

// Gets whether @stack is homogeneous.
// See gtk_stack_set_homogeneous().
/*

C function : gtk_stack_get_homogeneous
*/
func (recv *Stack) GetHomogeneous() bool {
	retC := C.gtk_stack_get_homogeneous((*C.GtkStack)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Returns the amount of time (in milliseconds) that
// transitions between pages in @stack will take.
/*

C function : gtk_stack_get_transition_duration
*/
func (recv *Stack) GetTransitionDuration() uint32 {
	retC := C.gtk_stack_get_transition_duration((*C.GtkStack)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// Gets the type of animation that will be used
// for transitions between pages in @stack.
/*

C function : gtk_stack_get_transition_type
*/
func (recv *Stack) GetTransitionType() StackTransitionType {
	retC := C.gtk_stack_get_transition_type((*C.GtkStack)(recv.native))
	retGo := (StackTransitionType)(retC)

	return retGo
}

// Gets the currently visible child of @stack, or %NULL if
// there are no visible children.
/*

C function : gtk_stack_get_visible_child
*/
func (recv *Stack) GetVisibleChild() *Widget {
	retC := C.gtk_stack_get_visible_child((*C.GtkStack)(recv.native))
	var retGo (*Widget)
	if retC == nil {
		retGo = nil
	} else {
		retGo = WidgetNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Returns the name of the currently visible child of @stack, or
// %NULL if there is no visible child.
/*

C function : gtk_stack_get_visible_child_name
*/
func (recv *Stack) GetVisibleChildName() string {
	retC := C.gtk_stack_get_visible_child_name((*C.GtkStack)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Sets the #GtkStack to be homogeneous or not. If it
// is homogeneous, the #GtkStack will request the same
// size for all its children. If it isn't, the stack
// may change size when a different child becomes visible.
//
// Since 3.16, homogeneity can be controlled separately
// for horizontal and vertical size, with the
// #GtkStack:hhomogeneous and #GtkStack:vhomogeneous.
/*

C function : gtk_stack_set_homogeneous
*/
func (recv *Stack) SetHomogeneous(homogeneous bool) {
	c_homogeneous :=
		boolToGboolean(homogeneous)

	C.gtk_stack_set_homogeneous((*C.GtkStack)(recv.native), c_homogeneous)

	return
}

// Sets the duration that transitions between pages in @stack
// will take.
/*

C function : gtk_stack_set_transition_duration
*/
func (recv *Stack) SetTransitionDuration(duration uint32) {
	c_duration := (C.guint)(duration)

	C.gtk_stack_set_transition_duration((*C.GtkStack)(recv.native), c_duration)

	return
}

// Sets the type of animation that will be used for
// transitions between pages in @stack. Available
// types include various kinds of fades and slides.
//
// The transition type can be changed without problems
// at runtime, so it is possible to change the animation
// based on the page that is about to become current.
/*

C function : gtk_stack_set_transition_type
*/
func (recv *Stack) SetTransitionType(transition StackTransitionType) {
	c_transition := (C.GtkStackTransitionType)(transition)

	C.gtk_stack_set_transition_type((*C.GtkStack)(recv.native), c_transition)

	return
}

// Makes @child the visible child of @stack.
//
// If @child is different from the currently
// visible child, the transition between the
// two will be animated with the current
// transition type of @stack.
//
// Note that the @child widget has to be visible itself
// (see gtk_widget_show()) in order to become the visible
// child of @stack.
/*

C function : gtk_stack_set_visible_child
*/
func (recv *Stack) SetVisibleChild(child *Widget) {
	c_child := (*C.GtkWidget)(C.NULL)
	if child != nil {
		c_child = (*C.GtkWidget)(child.ToC())
	}

	C.gtk_stack_set_visible_child((*C.GtkStack)(recv.native), c_child)

	return
}

// Makes the child with the given name visible.
//
// Note that the child widget has to be visible itself
// (see gtk_widget_show()) in order to become the visible
// child of @stack.
/*

C function : gtk_stack_set_visible_child_full
*/
func (recv *Stack) SetVisibleChildFull(name string, transition StackTransitionType) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_transition := (C.GtkStackTransitionType)(transition)

	C.gtk_stack_set_visible_child_full((*C.GtkStack)(recv.native), c_name, c_transition)

	return
}

// Makes the child with the given name visible.
//
// If @child is different from the currently
// visible child, the transition between the
// two will be animated with the current
// transition type of @stack.
//
// Note that the child widget has to be visible itself
// (see gtk_widget_show()) in order to become the visible
// child of @stack.
/*

C function : gtk_stack_set_visible_child_name
*/
func (recv *Stack) SetVisibleChildName(name string) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	C.gtk_stack_set_visible_child_name((*C.GtkStack)(recv.native), c_name)

	return
}

// Create a new #GtkStackSwitcher.
/*

C function : gtk_stack_switcher_new
*/
func StackSwitcherNew() *StackSwitcher {
	retC := C.gtk_stack_switcher_new()
	retGo := StackSwitcherNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Retrieves the stack.
// See gtk_stack_switcher_set_stack().
/*

C function : gtk_stack_switcher_get_stack
*/
func (recv *StackSwitcher) GetStack() *Stack {
	retC := C.gtk_stack_switcher_get_stack((*C.GtkStackSwitcher)(recv.native))
	var retGo (*Stack)
	if retC == nil {
		retGo = nil
	} else {
		retGo = StackNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Sets the stack to control.
/*

C function : gtk_stack_switcher_set_stack
*/
func (recv *StackSwitcher) SetStack(stack *Stack) {
	c_stack := (*C.GtkStack)(C.NULL)
	if stack != nil {
		c_stack = (*C.GtkStack)(stack.ToC())
	}

	C.gtk_stack_switcher_set_stack((*C.GtkStackSwitcher)(recv.native), c_stack)

	return
}

// Returns the scale used for assets.
/*

C function : gtk_style_context_get_scale
*/
func (recv *StyleContext) GetScale() int32 {
	retC := C.gtk_style_context_get_scale((*C.GtkStyleContext)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Sets the scale to use when getting image assets for the style.
/*

C function : gtk_style_context_set_scale
*/
func (recv *StyleContext) SetScale(scale int32) {
	c_scale := (C.gint)(scale)

	C.gtk_style_context_set_scale((*C.GtkStyleContext)(recv.native), c_scale)

	return
}

// Unsupported : gtk_drag_begin_with_coordinates : unsupported parameter event : no type generator for Gdk.Event (GdkEvent*) for param event

// Returns the baseline that has currently been allocated to @widget.
// This function is intended to be used when implementing handlers
// for the #GtkWidget::draw function, and when allocating child
// widgets in #GtkWidget::size_allocate.
/*

C function : gtk_widget_get_allocated_baseline
*/
func (recv *Widget) GetAllocatedBaseline() int32 {
	retC := C.gtk_widget_get_allocated_baseline((*C.GtkWidget)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Retrieves a widget’s minimum and natural height and the corresponding baselines if it would be given
// the specified @width, or the default height if @width is -1. The baselines may be -1 which means
// that no baseline is requested for this widget.
//
// The returned request will be modified by the
// GtkWidgetClass::adjust_size_request and GtkWidgetClass::adjust_baseline_request virtual methods
// and by any #GtkSizeGroups that have been applied. That is, the returned request
// is the one that should be used for layout, not necessarily the one
// returned by the widget itself.
/*

C function : gtk_widget_get_preferred_height_and_baseline_for_width
*/
func (recv *Widget) GetPreferredHeightAndBaselineForWidth(width int32) (int32, int32, int32, int32) {
	c_width := (C.gint)(width)

	var c_minimum_height C.gint

	var c_natural_height C.gint

	var c_minimum_baseline C.gint

	var c_natural_baseline C.gint

	C.gtk_widget_get_preferred_height_and_baseline_for_width((*C.GtkWidget)(recv.native), c_width, &c_minimum_height, &c_natural_height, &c_minimum_baseline, &c_natural_baseline)

	minimumHeight := (int32)(c_minimum_height)

	naturalHeight := (int32)(c_natural_height)

	minimumBaseline := (int32)(c_minimum_baseline)

	naturalBaseline := (int32)(c_natural_baseline)

	return minimumHeight, naturalHeight, minimumBaseline, naturalBaseline
}

// Retrieves the internal scale factor that maps from window coordinates
// to the actual device pixels. On traditional systems this is 1, on
// high density outputs, it can be a higher value (typically 2).
//
// See gdk_window_get_scale_factor().
/*

C function : gtk_widget_get_scale_factor
*/
func (recv *Widget) GetScaleFactor() int32 {
	retC := C.gtk_widget_get_scale_factor((*C.GtkWidget)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Gets the value of the #GtkWidget:valign property, including
// %GTK_ALIGN_BASELINE.
/*

C function : gtk_widget_get_valign_with_baseline
*/
func (recv *Widget) GetValignWithBaseline() Align {
	retC := C.gtk_widget_get_valign_with_baseline((*C.GtkWidget)(recv.native))
	retGo := (Align)(retC)

	return retGo
}

// Creates and initializes child widgets defined in templates. This
// function must be called in the instance initializer for any
// class which assigned itself a template using gtk_widget_class_set_template()
//
// It is important to call this function in the instance initializer
// of a #GtkWidget subclass and not in #GObject.constructed() or
// #GObject.constructor() for two reasons.
//
// One reason is that generally derived widgets will assume that parent
// class composite widgets have been created in their instance
// initializers.
//
// Another reason is that when calling g_object_new() on a widget with
// composite templates, it’s important to build the composite widgets
// before the construct properties are set. Properties passed to g_object_new()
// should take precedence over properties set in the private template XML.
/*

C function : gtk_widget_init_template
*/
func (recv *Widget) InitTemplate() {
	C.gtk_widget_init_template((*C.GtkWidget)(recv.native))

	return
}

// Unsupported : gtk_widget_size_allocate_with_baseline : unsupported parameter allocation : Blacklisted record : GdkRectangle

// Requests that the window is closed, similar to what happens
// when a window manager close button is clicked.
//
// This function can be used with close buttons in custom
// titlebars.
/*

C function : gtk_window_close
*/
func (recv *Window) Close() {
	C.gtk_window_close((*C.GtkWindow)(recv.native))

	return
}

// Sets a custom titlebar for @window.
//
// A typical widget used here is #GtkHeaderBar, as it provides various features
// expected of a titlebar while allowing the addition of child widgets to it.
//
// If you set a custom titlebar, GTK+ will do its best to convince
// the window manager not to put its own titlebar on the window.
// Depending on the system, this function may not work for a window
// that is already visible, so you set the titlebar before calling
// gtk_widget_show().
/*

C function : gtk_window_set_titlebar
*/
func (recv *Window) SetTitlebar(titlebar *Widget) {
	c_titlebar := (*C.GtkWidget)(C.NULL)
	if titlebar != nil {
		c_titlebar = (*C.GtkWidget)(titlebar.ToC())
	}

	C.gtk_window_set_titlebar((*C.GtkWindow)(recv.native), c_titlebar)

	return
}
