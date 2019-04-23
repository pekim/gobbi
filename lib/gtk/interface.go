// This is a generated file - DO NOT EDIT

package gtk

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

// Blacklisted : GtkActionable

// Blacklisted : GtkActivatable

// Blacklisted : GtkAppChooser

// Buildable is a wrapper around the C record GtkBuildable.
type Buildable struct {
	native *C.GtkBuildable
}

func BuildableNewFromC(u unsafe.Pointer) *Buildable {
	c := (*C.GtkBuildable)(u)
	if c == nil {
		return nil
	}

	g := &Buildable{native: c}

	return g
}

func (recv *Buildable) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Buildable with another Buildable, and returns true if they represent the same GObject.
func (recv *Buildable) Equals(other *Buildable) bool {
	return other.ToC() == recv.ToC()
}

// Blacklisted : GtkCellAccessibleParent

// Blacklisted : GtkCellEditable

// Blacklisted : GtkCellLayout

// Blacklisted : GtkColorChooser

// Blacklisted : GtkEditable

// Blacklisted : GtkFileChooser

// Blacklisted : GtkFontChooser

// Blacklisted : GtkOrientable

// Blacklisted : GtkPrintOperationPreview

// Blacklisted : GtkRecentChooser

// Blacklisted : GtkScrollable

// Blacklisted : GtkStyleProvider

// Blacklisted : GtkToolShell

// Blacklisted : GtkTreeDragDest

// Blacklisted : GtkTreeDragSource

// Blacklisted : GtkTreeModel

// Blacklisted : GtkTreeSortable
