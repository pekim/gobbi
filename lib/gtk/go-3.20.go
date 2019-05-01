// This is a generated file - DO NOT EDIT
// +build gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29
// +build gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29
// +build gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29
// +build gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29
// +build gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29
// +build gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29
// +build gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29
// +build gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29
// +build gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

type StyleContextPrintFlags int

const (
	GTK_STYLE_CONTEXT_PRINT_NONE       StyleContextPrintFlags = 0
	GTK_STYLE_CONTEXT_PRINT_RECURSE    StyleContextPrintFlags = 1
	GTK_STYLE_CONTEXT_PRINT_SHOW_STYLE StyleContextPrintFlags = 2
)

// Unsupported : gtk_file_chooser_native_new : return type :

const LEVEL_BAR_OFFSET_FULL string = "full"

type PopoverConstraint int

const (
	GTK_POPOVER_CONSTRAINT_NONE   PopoverConstraint = 0
	GTK_POPOVER_CONSTRAINT_WINDOW PopoverConstraint = 1
)

type ShortcutType int

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
