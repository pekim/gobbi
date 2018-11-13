// This is a generated file - DO NOT EDIT
// +build gtk_2.10 gtk_2.12 gtk_2.14 gtk_2.16 gtk_2.18 gtk_2.20 gtk_2.22 gtk_2.24 gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	glib "github.com/pekim/gobbi/lib/glib"
	"sync"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
/*

	void recentchooser_itemActivatedHandler(GObject *, gpointer);

	static gulong RecentChooser_signal_connect_item_activated(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "item-activated", G_CALLBACK(recentchooser_itemActivatedHandler), data);
	}

*/
/*

	void recentchooser_selectionChangedHandler(GObject *, gpointer);

	static gulong RecentChooser_signal_connect_selection_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "selection-changed", G_CALLBACK(recentchooser_selectionChangedHandler), data);
	}

*/
import "C"

// Ends a preview.
//
// This function must be called to finish a custom print preview.
/*

C function : gtk_print_operation_preview_end_preview
*/
func (recv *PrintOperationPreview) EndPreview() {
	C.gtk_print_operation_preview_end_preview((*C.GtkPrintOperationPreview)(recv.native))

	return
}

// Returns whether the given page is included in the set of pages that
// have been selected for printing.
/*

C function : gtk_print_operation_preview_is_selected
*/
func (recv *PrintOperationPreview) IsSelected(pageNr int32) bool {
	c_page_nr := (C.gint)(pageNr)

	retC := C.gtk_print_operation_preview_is_selected((*C.GtkPrintOperationPreview)(recv.native), c_page_nr)
	retGo := retC == C.TRUE

	return retGo
}

// Renders a page to the preview, using the print context that
// was passed to the #GtkPrintOperation::preview handler together
// with @preview.
//
// A custom iprint preview should use this function in its ::expose
// handler to render the currently selected page.
//
// Note that this function requires a suitable cairo context to
// be associated with the print context.
/*

C function : gtk_print_operation_preview_render_page
*/
func (recv *PrintOperationPreview) RenderPage(pageNr int32) {
	c_page_nr := (C.gint)(pageNr)

	C.gtk_print_operation_preview_render_page((*C.GtkPrintOperationPreview)(recv.native), c_page_nr)

	return
}

type signalRecentChooserItemActivatedDetail struct {
	callback  RecentChooserSignalItemActivatedCallback
	handlerID C.gulong
}

var signalRecentChooserItemActivatedId int
var signalRecentChooserItemActivatedMap = make(map[int]signalRecentChooserItemActivatedDetail)
var signalRecentChooserItemActivatedLock sync.Mutex

// RecentChooserSignalItemActivatedCallback is a callback function for a 'item-activated' signal emitted from a RecentChooser.
type RecentChooserSignalItemActivatedCallback func()

/*
ConnectItemActivated connects the callback to the 'item-activated' signal for the RecentChooser.

The returned value represents the connection, and may be passed to DisconnectItemActivated to remove it.
*/
func (recv *RecentChooser) ConnectItemActivated(callback RecentChooserSignalItemActivatedCallback) int {
	signalRecentChooserItemActivatedLock.Lock()
	defer signalRecentChooserItemActivatedLock.Unlock()

	signalRecentChooserItemActivatedId++
	instance := C.gpointer(recv.native)
	handlerID := C.RecentChooser_signal_connect_item_activated(instance, C.gpointer(uintptr(signalRecentChooserItemActivatedId)))

	detail := signalRecentChooserItemActivatedDetail{callback, handlerID}
	signalRecentChooserItemActivatedMap[signalRecentChooserItemActivatedId] = detail

	return signalRecentChooserItemActivatedId
}

/*
DisconnectItemActivated disconnects a callback from the 'item-activated' signal for the RecentChooser.

The connectionID should be a value returned from a call to ConnectItemActivated.
*/
func (recv *RecentChooser) DisconnectItemActivated(connectionID int) {
	signalRecentChooserItemActivatedLock.Lock()
	defer signalRecentChooserItemActivatedLock.Unlock()

	detail, exists := signalRecentChooserItemActivatedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalRecentChooserItemActivatedMap, connectionID)
}

//export recentchooser_itemActivatedHandler
func recentchooser_itemActivatedHandler(_ *C.GObject, data C.gpointer) {
	index := int(uintptr(data))
	callback := signalRecentChooserItemActivatedMap[index].callback
	callback()
}

type signalRecentChooserSelectionChangedDetail struct {
	callback  RecentChooserSignalSelectionChangedCallback
	handlerID C.gulong
}

var signalRecentChooserSelectionChangedId int
var signalRecentChooserSelectionChangedMap = make(map[int]signalRecentChooserSelectionChangedDetail)
var signalRecentChooserSelectionChangedLock sync.Mutex

// RecentChooserSignalSelectionChangedCallback is a callback function for a 'selection-changed' signal emitted from a RecentChooser.
type RecentChooserSignalSelectionChangedCallback func()

/*
ConnectSelectionChanged connects the callback to the 'selection-changed' signal for the RecentChooser.

The returned value represents the connection, and may be passed to DisconnectSelectionChanged to remove it.
*/
func (recv *RecentChooser) ConnectSelectionChanged(callback RecentChooserSignalSelectionChangedCallback) int {
	signalRecentChooserSelectionChangedLock.Lock()
	defer signalRecentChooserSelectionChangedLock.Unlock()

	signalRecentChooserSelectionChangedId++
	instance := C.gpointer(recv.native)
	handlerID := C.RecentChooser_signal_connect_selection_changed(instance, C.gpointer(uintptr(signalRecentChooserSelectionChangedId)))

	detail := signalRecentChooserSelectionChangedDetail{callback, handlerID}
	signalRecentChooserSelectionChangedMap[signalRecentChooserSelectionChangedId] = detail

	return signalRecentChooserSelectionChangedId
}

/*
DisconnectSelectionChanged disconnects a callback from the 'selection-changed' signal for the RecentChooser.

The connectionID should be a value returned from a call to ConnectSelectionChanged.
*/
func (recv *RecentChooser) DisconnectSelectionChanged(connectionID int) {
	signalRecentChooserSelectionChangedLock.Lock()
	defer signalRecentChooserSelectionChangedLock.Unlock()

	detail, exists := signalRecentChooserSelectionChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalRecentChooserSelectionChangedMap, connectionID)
}

//export recentchooser_selectionChangedHandler
func recentchooser_selectionChangedHandler(_ *C.GObject, data C.gpointer) {
	index := int(uintptr(data))
	callback := signalRecentChooserSelectionChangedMap[index].callback
	callback()
}

// Adds @filter to the list of #GtkRecentFilter objects held by @chooser.
//
// If no previous filter objects were defined, this function will call
// gtk_recent_chooser_set_filter().
/*

C function : gtk_recent_chooser_add_filter
*/
func (recv *RecentChooser) AddFilter(filter *RecentFilter) {
	c_filter := (*C.GtkRecentFilter)(C.NULL)
	if filter != nil {
		c_filter = (*C.GtkRecentFilter)(filter.ToC())
	}

	C.gtk_recent_chooser_add_filter((*C.GtkRecentChooser)(recv.native), c_filter)

	return
}

// Gets the #GtkRecentInfo currently selected by @chooser.
/*

C function : gtk_recent_chooser_get_current_item
*/
func (recv *RecentChooser) GetCurrentItem() *RecentInfo {
	retC := C.gtk_recent_chooser_get_current_item((*C.GtkRecentChooser)(recv.native))
	retGo := RecentInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the URI currently selected by @chooser.
/*

C function : gtk_recent_chooser_get_current_uri
*/
func (recv *RecentChooser) GetCurrentUri() string {
	retC := C.gtk_recent_chooser_get_current_uri((*C.GtkRecentChooser)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Gets the #GtkRecentFilter object currently used by @chooser to affect
// the display of the recently used resources.
/*

C function : gtk_recent_chooser_get_filter
*/
func (recv *RecentChooser) GetFilter() *RecentFilter {
	retC := C.gtk_recent_chooser_get_filter((*C.GtkRecentChooser)(recv.native))
	retGo := RecentFilterNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the list of recently used resources in form of #GtkRecentInfo objects.
//
// The return value of this function is affected by the “sort-type” and
// “limit” properties of @chooser.
/*

C function : gtk_recent_chooser_get_items
*/
func (recv *RecentChooser) GetItems() *glib.List {
	retC := C.gtk_recent_chooser_get_items((*C.GtkRecentChooser)(recv.native))
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the number of items returned by gtk_recent_chooser_get_items()
// and gtk_recent_chooser_get_uris().
/*

C function : gtk_recent_chooser_get_limit
*/
func (recv *RecentChooser) GetLimit() int32 {
	retC := C.gtk_recent_chooser_get_limit((*C.GtkRecentChooser)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Gets whether only local resources should be shown in the recently used
// resources selector.  See gtk_recent_chooser_set_local_only()
/*

C function : gtk_recent_chooser_get_local_only
*/
func (recv *RecentChooser) GetLocalOnly() bool {
	retC := C.gtk_recent_chooser_get_local_only((*C.GtkRecentChooser)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Gets whether @chooser can select multiple items.
/*

C function : gtk_recent_chooser_get_select_multiple
*/
func (recv *RecentChooser) GetSelectMultiple() bool {
	retC := C.gtk_recent_chooser_get_select_multiple((*C.GtkRecentChooser)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Retrieves whether @chooser should show an icon near the resource.
/*

C function : gtk_recent_chooser_get_show_icons
*/
func (recv *RecentChooser) GetShowIcons() bool {
	retC := C.gtk_recent_chooser_get_show_icons((*C.GtkRecentChooser)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Retrieves whether @chooser should show the recently used resources that
// were not found.
/*

C function : gtk_recent_chooser_get_show_not_found
*/
func (recv *RecentChooser) GetShowNotFound() bool {
	retC := C.gtk_recent_chooser_get_show_not_found((*C.GtkRecentChooser)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Returns whether @chooser should display recently used resources
// registered as private.
/*

C function : gtk_recent_chooser_get_show_private
*/
func (recv *RecentChooser) GetShowPrivate() bool {
	retC := C.gtk_recent_chooser_get_show_private((*C.GtkRecentChooser)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Gets whether @chooser should display tooltips containing the full path
// of a recently user resource.
/*

C function : gtk_recent_chooser_get_show_tips
*/
func (recv *RecentChooser) GetShowTips() bool {
	retC := C.gtk_recent_chooser_get_show_tips((*C.GtkRecentChooser)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Gets the value set by gtk_recent_chooser_set_sort_type().
/*

C function : gtk_recent_chooser_get_sort_type
*/
func (recv *RecentChooser) GetSortType() RecentSortType {
	retC := C.gtk_recent_chooser_get_sort_type((*C.GtkRecentChooser)(recv.native))
	retGo := (RecentSortType)(retC)

	return retGo
}

// Unsupported : gtk_recent_chooser_get_uris : no return type

// Gets the #GtkRecentFilter objects held by @chooser.
/*

C function : gtk_recent_chooser_list_filters
*/
func (recv *RecentChooser) ListFilters() *glib.SList {
	retC := C.gtk_recent_chooser_list_filters((*C.GtkRecentChooser)(recv.native))
	retGo := glib.SListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Removes @filter from the list of #GtkRecentFilter objects held by @chooser.
/*

C function : gtk_recent_chooser_remove_filter
*/
func (recv *RecentChooser) RemoveFilter(filter *RecentFilter) {
	c_filter := (*C.GtkRecentFilter)(C.NULL)
	if filter != nil {
		c_filter = (*C.GtkRecentFilter)(filter.ToC())
	}

	C.gtk_recent_chooser_remove_filter((*C.GtkRecentChooser)(recv.native), c_filter)

	return
}

// Selects all the items inside @chooser, if the @chooser supports
// multiple selection.
/*

C function : gtk_recent_chooser_select_all
*/
func (recv *RecentChooser) SelectAll() {
	C.gtk_recent_chooser_select_all((*C.GtkRecentChooser)(recv.native))

	return
}

// Selects @uri inside @chooser.
/*

C function : gtk_recent_chooser_select_uri
*/
func (recv *RecentChooser) SelectUri(uri string) (bool, error) {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	var cThrowableError *C.GError

	retC := C.gtk_recent_chooser_select_uri((*C.GtkRecentChooser)(recv.native), c_uri, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Sets @uri as the current URI for @chooser.
/*

C function : gtk_recent_chooser_set_current_uri
*/
func (recv *RecentChooser) SetCurrentUri(uri string) (bool, error) {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	var cThrowableError *C.GError

	retC := C.gtk_recent_chooser_set_current_uri((*C.GtkRecentChooser)(recv.native), c_uri, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Sets @filter as the current #GtkRecentFilter object used by @chooser
// to affect the displayed recently used resources.
/*

C function : gtk_recent_chooser_set_filter
*/
func (recv *RecentChooser) SetFilter(filter *RecentFilter) {
	c_filter := (*C.GtkRecentFilter)(C.NULL)
	if filter != nil {
		c_filter = (*C.GtkRecentFilter)(filter.ToC())
	}

	C.gtk_recent_chooser_set_filter((*C.GtkRecentChooser)(recv.native), c_filter)

	return
}

// Sets the number of items that should be returned by
// gtk_recent_chooser_get_items() and gtk_recent_chooser_get_uris().
/*

C function : gtk_recent_chooser_set_limit
*/
func (recv *RecentChooser) SetLimit(limit int32) {
	c_limit := (C.gint)(limit)

	C.gtk_recent_chooser_set_limit((*C.GtkRecentChooser)(recv.native), c_limit)

	return
}

// Sets whether only local resources, that is resources using the file:// URI
// scheme, should be shown in the recently used resources selector.  If
// @local_only is %TRUE (the default) then the shown resources are guaranteed
// to be accessible through the operating system native file system.
/*

C function : gtk_recent_chooser_set_local_only
*/
func (recv *RecentChooser) SetLocalOnly(localOnly bool) {
	c_local_only :=
		boolToGboolean(localOnly)

	C.gtk_recent_chooser_set_local_only((*C.GtkRecentChooser)(recv.native), c_local_only)

	return
}

// Sets whether @chooser can select multiple items.
/*

C function : gtk_recent_chooser_set_select_multiple
*/
func (recv *RecentChooser) SetSelectMultiple(selectMultiple bool) {
	c_select_multiple :=
		boolToGboolean(selectMultiple)

	C.gtk_recent_chooser_set_select_multiple((*C.GtkRecentChooser)(recv.native), c_select_multiple)

	return
}

// Sets whether @chooser should show an icon near the resource when
// displaying it.
/*

C function : gtk_recent_chooser_set_show_icons
*/
func (recv *RecentChooser) SetShowIcons(showIcons bool) {
	c_show_icons :=
		boolToGboolean(showIcons)

	C.gtk_recent_chooser_set_show_icons((*C.GtkRecentChooser)(recv.native), c_show_icons)

	return
}

// Sets whether @chooser should display the recently used resources that
// it didn’t find.  This only applies to local resources.
/*

C function : gtk_recent_chooser_set_show_not_found
*/
func (recv *RecentChooser) SetShowNotFound(showNotFound bool) {
	c_show_not_found :=
		boolToGboolean(showNotFound)

	C.gtk_recent_chooser_set_show_not_found((*C.GtkRecentChooser)(recv.native), c_show_not_found)

	return
}

// Whether to show recently used resources marked registered as private.
/*

C function : gtk_recent_chooser_set_show_private
*/
func (recv *RecentChooser) SetShowPrivate(showPrivate bool) {
	c_show_private :=
		boolToGboolean(showPrivate)

	C.gtk_recent_chooser_set_show_private((*C.GtkRecentChooser)(recv.native), c_show_private)

	return
}

// Sets whether to show a tooltips containing the full path of each
// recently used resource in a #GtkRecentChooser widget.
/*

C function : gtk_recent_chooser_set_show_tips
*/
func (recv *RecentChooser) SetShowTips(showTips bool) {
	c_show_tips :=
		boolToGboolean(showTips)

	C.gtk_recent_chooser_set_show_tips((*C.GtkRecentChooser)(recv.native), c_show_tips)

	return
}

// Unsupported : gtk_recent_chooser_set_sort_func : unsupported parameter sort_func : no type generator for RecentSortFunc (GtkRecentSortFunc) for param sort_func

// Changes the sorting order of the recently used resources list displayed by
// @chooser.
/*

C function : gtk_recent_chooser_set_sort_type
*/
func (recv *RecentChooser) SetSortType(sortType RecentSortType) {
	c_sort_type := (C.GtkRecentSortType)(sortType)

	C.gtk_recent_chooser_set_sort_type((*C.GtkRecentChooser)(recv.native), c_sort_type)

	return
}

// Unselects all the items inside @chooser.
/*

C function : gtk_recent_chooser_unselect_all
*/
func (recv *RecentChooser) UnselectAll() {
	C.gtk_recent_chooser_unselect_all((*C.GtkRecentChooser)(recv.native))

	return
}

// Unselects @uri inside @chooser.
/*

C function : gtk_recent_chooser_unselect_uri
*/
func (recv *RecentChooser) UnselectUri(uri string) {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	C.gtk_recent_chooser_unselect_uri((*C.GtkRecentChooser)(recv.native), c_uri)

	return
}
