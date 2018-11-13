// This is a generated file - DO NOT EDIT
// +build gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

/*
Describes constraints to positioning of popovers. More values
may be added to this enumeration in the future.
*/
type PopoverConstraint C.GtkPopoverConstraint

const (
	/*
	   Don't constrain the popover position
	     beyond what is imposed by the implementation
	*/
	GTK_POPOVER_CONSTRAINT_NONE PopoverConstraint = 0
	/*
	   Constrain the popover to the boundaries
	     of the window that it is attached to
	*/
	GTK_POPOVER_CONSTRAINT_WINDOW PopoverConstraint = 1
)

/*
GtkShortcutType specifies the kind of shortcut that is being described.
More values may be added to this enumeration over time.
*/
type ShortcutType C.GtkShortcutType

const (
	/*
	   The shortcut is a keyboard accelerator. The #GtkShortcutsShortcut:accelerator
	     property will be used.
	*/
	GTK_SHORTCUT_ACCELERATOR ShortcutType = 0
	// The shortcut is a pinch gesture. GTK+ provides an icon and subtitle.
	GTK_SHORTCUT_GESTURE_PINCH ShortcutType = 1
	// The shortcut is a stretch gesture. GTK+ provides an icon and subtitle.
	GTK_SHORTCUT_GESTURE_STRETCH ShortcutType = 2
	// The shortcut is a clockwise rotation gesture. GTK+ provides an icon and subtitle.
	GTK_SHORTCUT_GESTURE_ROTATE_CLOCKWISE ShortcutType = 3
	// The shortcut is a counterclockwise rotation gesture. GTK+ provides an icon and subtitle.
	GTK_SHORTCUT_GESTURE_ROTATE_COUNTERCLOCKWISE ShortcutType = 4
	// The shortcut is a two-finger swipe gesture. GTK+ provides an icon and subtitle.
	GTK_SHORTCUT_GESTURE_TWO_FINGER_SWIPE_LEFT ShortcutType = 5
	// The shortcut is a two-finger swipe gesture. GTK+ provides an icon and subtitle.
	GTK_SHORTCUT_GESTURE_TWO_FINGER_SWIPE_RIGHT ShortcutType = 6
	/*
	   The shortcut is a gesture. The #GtkShortcutsShortcut:icon property will be
	     used.
	*/
	GTK_SHORTCUT_GESTURE ShortcutType = 7
)
