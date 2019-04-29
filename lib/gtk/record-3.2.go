// This is a generated file - DO NOT EDIT
// +build gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import "unsafe"

// CssSection is a wrapper around the C record GtkCssSection.
type CssSection struct {
	native unsafe.Pointer
}

func CssSectionNewFromC(u unsafe.Pointer) *CssSection {
	if u == nil {
		return nil
	}

	g := &CssSection{native: u}

	return g
}

func (recv *CssSection) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}
