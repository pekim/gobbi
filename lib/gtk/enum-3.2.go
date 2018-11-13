// This is a generated file - DO NOT EDIT
// +build gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

// The different types of sections indicate parts of a CSS document as
// parsed by GTK’s CSS parser. They are oriented towards the
// [CSS Grammar](http://www.w3.org/TR/CSS21/grammar.html),
// but may contain extensions.
//
// More types might be added in the future as the parser incorporates
// more features.
type CssSectionType C.GtkCssSectionType

const (
	// The section describes a complete document.
	// This section time is the only one where gtk_css_section_get_parent()
	// might return %NULL.
	GTK_CSS_SECTION_DOCUMENT CssSectionType = 0

	// The section defines an import rule.
	GTK_CSS_SECTION_IMPORT CssSectionType = 1

	// The section defines a color. This
	// is a GTK extension to CSS.
	GTK_CSS_SECTION_COLOR_DEFINITION CssSectionType = 2

	// The section defines a binding set. This
	// is a GTK extension to CSS.
	GTK_CSS_SECTION_BINDING_SET CssSectionType = 3

	// The section defines a CSS ruleset.
	GTK_CSS_SECTION_RULESET CssSectionType = 4

	// The section defines a CSS selector.
	GTK_CSS_SECTION_SELECTOR CssSectionType = 5

	// The section defines the declaration of
	// a CSS variable.
	GTK_CSS_SECTION_DECLARATION CssSectionType = 6

	// The section defines the value of a CSS declaration.
	GTK_CSS_SECTION_VALUE CssSectionType = 7

	// The section defines keyframes. See [CSS
	// Animations](http://dev.w3.org/csswg/css3-animations/#keyframes) for details. Since 3.6
	GTK_CSS_SECTION_KEYFRAMES CssSectionType = 8
)
