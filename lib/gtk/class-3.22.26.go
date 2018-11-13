// This is a generated file - DO NOT EDIT
// +build gtk_3.22.26 gtk_3.22.29

package gtk

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

// Unsupported signal 'show-starred-location' for PlacesSidebar : unsupported parameter object : type PlacesOpenFlags :

// Returns the value previously set with gtk_places_sidebar_set_show_starred_location()
/*

C function

gtk_places_sidebar_get_show_starred_location
*/
func (recv *PlacesSidebar) GetShowStarredLocation() bool {
	retC := C.gtk_places_sidebar_get_show_starred_location((*C.GtkPlacesSidebar)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// If you enable this, you should connect to the
// #GtkPlacesSidebar::show-starred-location signal.
/*

C function

gtk_places_sidebar_set_show_starred_location
*/
func (recv *PlacesSidebar) SetShowStarredLocation(showStarredLocation bool) {
	c_show_starred_location :=
		boolToGboolean(showStarredLocation)

	C.gtk_places_sidebar_set_show_starred_location((*C.GtkPlacesSidebar)(recv.native), c_show_starred_location)

	return
}
