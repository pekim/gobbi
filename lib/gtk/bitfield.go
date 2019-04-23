// This is a generated file - DO NOT EDIT

package gtk

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

// Blacklisted : GtkAccelFlags

// Blacklisted : GtkAttachOptions

// Blacklisted : GtkCalendarDisplayOptions

// Blacklisted : GtkCellRendererState

// Blacklisted : GtkDebugFlag

// Blacklisted : GtkDestDefaults

// Blacklisted : GtkDialogFlags

// Blacklisted : GtkFileFilterFlags

// Blacklisted : GtkIconLookupFlags

// Blacklisted : GtkJunctionSides

// Blacklisted : GtkPlacesOpenFlags

// Blacklisted : GtkRcFlags

// Blacklisted : GtkRecentFilterFlags

// Blacklisted : GtkRegionFlags

type StateFlags C.GtkStateFlags

const (
	GTK_STATE_FLAG_NORMAL       StateFlags = 0
	GTK_STATE_FLAG_ACTIVE       StateFlags = 1
	GTK_STATE_FLAG_PRELIGHT     StateFlags = 2
	GTK_STATE_FLAG_SELECTED     StateFlags = 4
	GTK_STATE_FLAG_INSENSITIVE  StateFlags = 8
	GTK_STATE_FLAG_INCONSISTENT StateFlags = 16
	GTK_STATE_FLAG_FOCUSED      StateFlags = 32
	GTK_STATE_FLAG_BACKDROP     StateFlags = 64
	GTK_STATE_FLAG_DIR_LTR      StateFlags = 128
	GTK_STATE_FLAG_DIR_RTL      StateFlags = 256
	GTK_STATE_FLAG_LINK         StateFlags = 512
	GTK_STATE_FLAG_VISITED      StateFlags = 1024
	GTK_STATE_FLAG_CHECKED      StateFlags = 2048
	GTK_STATE_FLAG_DROP_ACTIVE  StateFlags = 4096
)

// Blacklisted : GtkTargetFlags

// Blacklisted : GtkTextSearchFlags

// Blacklisted : GtkToolPaletteDragTargets

// Blacklisted : GtkTreeModelFlags

// Blacklisted : GtkUIManagerItemType
