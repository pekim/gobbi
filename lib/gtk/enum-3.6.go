// This is a generated file - DO NOT EDIT
// +build gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

/*
Describes primary purpose of the input widget. This information is
useful for on-screen keyboards and similar input methods to decide
which keys should be presented to the user.

Note that the purpose is not meant to impose a totally strict rule
about allowed characters, and does not replace input validation.
It is fine for an on-screen keyboard to let the user override the
character set restriction that is expressed by the purpose. The
application is expected to validate the entry contents, even if
it specified a purpose.

The difference between @GTK_INPUT_PURPOSE_DIGITS and
@GTK_INPUT_PURPOSE_NUMBER is that the former accepts only digits
while the latter also some punctuation (like commas or points, plus,
minus) and “e” or “E” as in 3.14E+000.

This enumeration may be extended in the future; input methods should
interpret unknown values as “free form”.
*/
type InputPurpose C.GtkInputPurpose

const (
	// Allow any character
	GTK_INPUT_PURPOSE_FREE_FORM InputPurpose = 0
	// Allow only alphabetic characters
	GTK_INPUT_PURPOSE_ALPHA InputPurpose = 1
	// Allow only digits
	GTK_INPUT_PURPOSE_DIGITS InputPurpose = 2
	// Edited field expects numbers
	GTK_INPUT_PURPOSE_NUMBER InputPurpose = 3
	// Edited field expects phone number
	GTK_INPUT_PURPOSE_PHONE InputPurpose = 4
	// Edited field expects URL
	GTK_INPUT_PURPOSE_URL InputPurpose = 5
	// Edited field expects email address
	GTK_INPUT_PURPOSE_EMAIL InputPurpose = 6
	// Edited field expects the name of a person
	GTK_INPUT_PURPOSE_NAME InputPurpose = 7
	// Like @GTK_INPUT_PURPOSE_FREE_FORM, but characters are hidden
	GTK_INPUT_PURPOSE_PASSWORD InputPurpose = 8
	// Like @GTK_INPUT_PURPOSE_DIGITS, but characters are hidden
	GTK_INPUT_PURPOSE_PIN InputPurpose = 9
)

/*
Describes how #GtkLevelBar contents should be rendered.
Note that this enumeration could be extended with additional modes
in the future.
*/
type LevelBarMode C.GtkLevelBarMode

const (
	// the bar has a continuous mode
	GTK_LEVEL_BAR_MODE_CONTINUOUS LevelBarMode = 0
	// the bar has a discrete mode
	GTK_LEVEL_BAR_MODE_DISCRETE LevelBarMode = 1
)
