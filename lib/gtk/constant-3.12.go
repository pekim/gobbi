// This is a generated file - DO NOT EDIT
// +build gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

// A CSS class used when an action (usually a button) is
// one that is expected to remove or destroy something visible
// to the user.
//
// Refer to individual widget documentation for used style classes.
const STYLE_CLASS_DESTRUCTIVE_ACTION string = C.GTK_STYLE_CLASS_DESTRUCTIVE_ACTION

// A CSS class used when an element needs the user attention,
// for instance a button in a stack switcher corresponding to
// a hidden page that changed state.
//
// Refer to individual widget documentation for used style classes.
const STYLE_CLASS_NEEDS_ATTENTION string = C.GTK_STYLE_CLASS_NEEDS_ATTENTION

// A CSS class used when an action (usually a button) is the
// primary suggested action in a specific context.
//
// Refer to individual widget documentation for used style classes.
const STYLE_CLASS_SUGGESTED_ACTION string = C.GTK_STYLE_CLASS_SUGGESTED_ACTION
