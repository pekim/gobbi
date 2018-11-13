// This is a generated file - DO NOT EDIT
// +build gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

// Describes hints that might be taken into account by input methods
// or applications. Note that input methods may already tailor their
// behaviour according to the #GtkInputPurpose of the entry.
//
// Some common sense is expected when using these flags - mixing
// @GTK_INPUT_HINT_LOWERCASE with any of the uppercase hints makes no sense.
//
// This enumeration may be extended in the future; input methods should
// ignore unknown values.
type InputHints C.GtkInputHints

const (
	// No special behaviour suggested
	GTK_INPUT_HINT_NONE InputHints = 0
	// Suggest checking for typos
	GTK_INPUT_HINT_SPELLCHECK InputHints = 1
	// Suggest not checking for typos
	GTK_INPUT_HINT_NO_SPELLCHECK InputHints = 2
	// Suggest word completion
	GTK_INPUT_HINT_WORD_COMPLETION InputHints = 4
	// Suggest to convert all text to lowercase
	GTK_INPUT_HINT_LOWERCASE InputHints = 8
	// Suggest to capitalize all text
	GTK_INPUT_HINT_UPPERCASE_CHARS InputHints = 16
	// Suggest to capitalize the first
	// character of each word
	GTK_INPUT_HINT_UPPERCASE_WORDS InputHints = 32
	// Suggest to capitalize the
	// first word of each sentence
	GTK_INPUT_HINT_UPPERCASE_SENTENCES InputHints = 64
	// Suggest to not show an onscreen keyboard
	// (e.g for a calculator that already has all the keys).
	GTK_INPUT_HINT_INHIBIT_OSK InputHints = 128
	// The text is vertical. Since 3.18
	GTK_INPUT_HINT_VERTICAL_WRITING InputHints = 256
	// Suggest offering Emoji support. Since 3.22.20
	GTK_INPUT_HINT_EMOJI InputHints = 512
	// Suggest not offering Emoji support. Since 3.22.20
	GTK_INPUT_HINT_NO_EMOJI InputHints = 1024
)
