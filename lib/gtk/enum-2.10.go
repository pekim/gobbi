// This is a generated file - DO NOT EDIT
// +build gtk_2.10 gtk_2.12 gtk_2.14 gtk_2.16 gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.20 gtk_3.22

package gtk

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

type RecentChooserError C.GtkRecentChooserError

const (
	GTK_RECENT_CHOOSER_ERROR_NOT_FOUND   RecentChooserError = 0
	GTK_RECENT_CHOOSER_ERROR_INVALID_URI RecentChooserError = 1
)

type RecentManagerError C.GtkRecentManagerError

const (
	GTK_RECENT_MANAGER_ERROR_NOT_FOUND        RecentManagerError = 0
	GTK_RECENT_MANAGER_ERROR_INVALID_URI      RecentManagerError = 1
	GTK_RECENT_MANAGER_ERROR_INVALID_ENCODING RecentManagerError = 2
	GTK_RECENT_MANAGER_ERROR_NOT_REGISTERED   RecentManagerError = 3
	GTK_RECENT_MANAGER_ERROR_READ             RecentManagerError = 4
	GTK_RECENT_MANAGER_ERROR_WRITE            RecentManagerError = 5
	GTK_RECENT_MANAGER_ERROR_UNKNOWN          RecentManagerError = 6
)

type RecentSortType C.GtkRecentSortType

const (
	GTK_RECENT_SORT_NONE   RecentSortType = 0
	GTK_RECENT_SORT_MRU    RecentSortType = 1
	GTK_RECENT_SORT_LRU    RecentSortType = 2
	GTK_RECENT_SORT_CUSTOM RecentSortType = 3
)
