// This is a generated file - DO NOT EDIT
// +build gtk_2.10 gtk_2.12 gtk_2.14 gtk_2.16 gtk_2.18 gtk_2.20 gtk_2.22 gtk_2.24 gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

// These identify the various errors that can occur while calling
// #GtkRecentChooser functions.
type RecentChooserError C.GtkRecentChooserError

const (
	// Indicates that a file does not exist
	GTK_RECENT_CHOOSER_ERROR_NOT_FOUND RecentChooserError = 0

	// Indicates a malformed URI
	GTK_RECENT_CHOOSER_ERROR_INVALID_URI RecentChooserError = 1
)

// Error codes for #GtkRecentManager operations
type RecentManagerError C.GtkRecentManagerError

const (
	// the URI specified does not exists in
	// the recently used resources list.
	GTK_RECENT_MANAGER_ERROR_NOT_FOUND RecentManagerError = 0

	// the URI specified is not valid.
	GTK_RECENT_MANAGER_ERROR_INVALID_URI RecentManagerError = 1

	// the supplied string is not
	// UTF-8 encoded.
	GTK_RECENT_MANAGER_ERROR_INVALID_ENCODING RecentManagerError = 2

	// no application has registered
	// the specified item.
	GTK_RECENT_MANAGER_ERROR_NOT_REGISTERED RecentManagerError = 3

	// failure while reading the recently used
	// resources file.
	GTK_RECENT_MANAGER_ERROR_READ RecentManagerError = 4

	// failure while writing the recently used
	// resources file.
	GTK_RECENT_MANAGER_ERROR_WRITE RecentManagerError = 5

	// unspecified error.
	GTK_RECENT_MANAGER_ERROR_UNKNOWN RecentManagerError = 6
)

// Used to specify the sorting method to be applyed to the recently
// used resource list.
type RecentSortType C.GtkRecentSortType

const (
	// Do not sort the returned list of recently used
	// resources.
	GTK_RECENT_SORT_NONE RecentSortType = 0

	// Sort the returned list with the most recently used
	// items first.
	GTK_RECENT_SORT_MRU RecentSortType = 1

	// Sort the returned list with the least recently used
	// items first.
	GTK_RECENT_SORT_LRU RecentSortType = 2

	// Sort the returned list using a custom sorting
	// function passed using gtk_recent_chooser_set_sort_func().
	GTK_RECENT_SORT_CUSTOM RecentSortType = 3
)
