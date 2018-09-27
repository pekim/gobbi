// This is a generated file - DO NOT EDIT
// +build gtk_3.20 gtk_3.22

package gtk

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

type PopoverConstraint C.GtkPopoverConstraint

const (
	GTK_POPOVER_CONSTRAINT_NONE   PopoverConstraint = 0
	GTK_POPOVER_CONSTRAINT_WINDOW PopoverConstraint = 1
)

type ShortcutType C.GtkShortcutType

const (
	GTK_SHORTCUT_ACCELERATOR                     ShortcutType = 0
	GTK_SHORTCUT_GESTURE_PINCH                   ShortcutType = 1
	GTK_SHORTCUT_GESTURE_STRETCH                 ShortcutType = 2
	GTK_SHORTCUT_GESTURE_ROTATE_CLOCKWISE        ShortcutType = 3
	GTK_SHORTCUT_GESTURE_ROTATE_COUNTERCLOCKWISE ShortcutType = 4
	GTK_SHORTCUT_GESTURE_TWO_FINGER_SWIPE_LEFT   ShortcutType = 5
	GTK_SHORTCUT_GESTURE_TWO_FINGER_SWIPE_RIGHT  ShortcutType = 6
	GTK_SHORTCUT_GESTURE                         ShortcutType = 7
)
