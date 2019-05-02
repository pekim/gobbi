// Code generated - DO NOT EDIT.
// +build gtk_3.22.26 gtk_3.22.29

package gtk

import "sync"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
/*

	void placessidebar_showStarredLocationHandler(GObject *, GtkPlacesOpenFlags, gpointer);

	static gulong PlacesSidebar_signal_connect_show_starred_location(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "show-starred-location", G_CALLBACK(placessidebar_showStarredLocationHandler), data);
	}

*/
import "C"

type signalPlacesSidebarShowStarredLocationDetail struct {
	callback  PlacesSidebarSignalShowStarredLocationCallback
	handlerID C.gulong
}

var signalPlacesSidebarShowStarredLocationId int
var signalPlacesSidebarShowStarredLocationMap = make(map[int]signalPlacesSidebarShowStarredLocationDetail)
var signalPlacesSidebarShowStarredLocationLock sync.RWMutex

// PlacesSidebarSignalShowStarredLocationCallback is a callback function for a 'show-starred-location' signal emitted from a PlacesSidebar.
type PlacesSidebarSignalShowStarredLocationCallback func(object PlacesOpenFlags)

/*
ConnectShowStarredLocation connects the callback to the 'show-starred-location' signal for the PlacesSidebar.

The returned value represents the connection, and may be passed to DisconnectShowStarredLocation to remove it.
*/
func (recv *PlacesSidebar) ConnectShowStarredLocation(callback PlacesSidebarSignalShowStarredLocationCallback) int {
	signalPlacesSidebarShowStarredLocationLock.Lock()
	defer signalPlacesSidebarShowStarredLocationLock.Unlock()

	signalPlacesSidebarShowStarredLocationId++
	instance := C.gpointer(recv.native)
	handlerID := C.PlacesSidebar_signal_connect_show_starred_location(instance, C.gpointer(uintptr(signalPlacesSidebarShowStarredLocationId)))

	detail := signalPlacesSidebarShowStarredLocationDetail{callback, handlerID}
	signalPlacesSidebarShowStarredLocationMap[signalPlacesSidebarShowStarredLocationId] = detail

	return signalPlacesSidebarShowStarredLocationId
}

/*
DisconnectShowStarredLocation disconnects a callback from the 'show-starred-location' signal for the PlacesSidebar.

The connectionID should be a value returned from a call to ConnectShowStarredLocation.
*/
func (recv *PlacesSidebar) DisconnectShowStarredLocation(connectionID int) {
	signalPlacesSidebarShowStarredLocationLock.Lock()
	defer signalPlacesSidebarShowStarredLocationLock.Unlock()

	detail, exists := signalPlacesSidebarShowStarredLocationMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalPlacesSidebarShowStarredLocationMap, connectionID)
}

//export placessidebar_showStarredLocationHandler
func placessidebar_showStarredLocationHandler(_ *C.GObject, c_object C.GtkPlacesOpenFlags, data C.gpointer) {
	signalPlacesSidebarShowStarredLocationLock.RLock()
	defer signalPlacesSidebarShowStarredLocationLock.RUnlock()

	object := PlacesOpenFlags(c_object)

	index := int(uintptr(data))
	callback := signalPlacesSidebarShowStarredLocationMap[index].callback
	callback(object)
}

// GetShowStarredLocation is a wrapper around the C function gtk_places_sidebar_get_show_starred_location.
func (recv *PlacesSidebar) GetShowStarredLocation() bool {
	retC := C.gtk_places_sidebar_get_show_starred_location((*C.GtkPlacesSidebar)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetShowStarredLocation is a wrapper around the C function gtk_places_sidebar_set_show_starred_location.
func (recv *PlacesSidebar) SetShowStarredLocation(showStarredLocation bool) {
	c_show_starred_location :=
		boolToGboolean(showStarredLocation)

	C.gtk_places_sidebar_set_show_starred_location((*C.GtkPlacesSidebar)(recv.native), c_show_starred_location)

	return
}
