// This is a generated file - DO NOT EDIT
// +build gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29
// +build gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29
// +build gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29
// +build gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29
// +build gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29
// +build gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29
// +build gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29
// +build gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29
// +build gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

type InputHints int

const (
	GTK_INPUT_HINT_NONE                InputHints = 0
	GTK_INPUT_HINT_SPELLCHECK          InputHints = 1
	GTK_INPUT_HINT_NO_SPELLCHECK       InputHints = 2
	GTK_INPUT_HINT_WORD_COMPLETION     InputHints = 4
	GTK_INPUT_HINT_LOWERCASE           InputHints = 8
	GTK_INPUT_HINT_UPPERCASE_CHARS     InputHints = 16
	GTK_INPUT_HINT_UPPERCASE_WORDS     InputHints = 32
	GTK_INPUT_HINT_UPPERCASE_SENTENCES InputHints = 64
	GTK_INPUT_HINT_INHIBIT_OSK         InputHints = 128
	GTK_INPUT_HINT_VERTICAL_WRITING    InputHints = 256
	GTK_INPUT_HINT_EMOJI               InputHints = 512
	GTK_INPUT_HINT_NO_EMOJI            InputHints = 1024
)

// Unsupported : gtk_level_bar_new : return type :

// Unsupported : gtk_level_bar_new_for_interval : return type :

// Unsupported : gtk_menu_button_new : return type :

// Unsupported : gtk_search_entry_new : return type :

const LEVEL_BAR_OFFSET_HIGH string = "high"
const LEVEL_BAR_OFFSET_LOW string = "low"
const PRINT_SETTINGS_OUTPUT_BASENAME string = "output-basename"
const PRINT_SETTINGS_OUTPUT_DIR string = "output-dir"

type InputPurpose int

const (
	GTK_INPUT_PURPOSE_FREE_FORM InputPurpose = 0
	GTK_INPUT_PURPOSE_ALPHA     InputPurpose = 1
	GTK_INPUT_PURPOSE_DIGITS    InputPurpose = 2
	GTK_INPUT_PURPOSE_NUMBER    InputPurpose = 3
	GTK_INPUT_PURPOSE_PHONE     InputPurpose = 4
	GTK_INPUT_PURPOSE_URL       InputPurpose = 5
	GTK_INPUT_PURPOSE_EMAIL     InputPurpose = 6
	GTK_INPUT_PURPOSE_NAME      InputPurpose = 7
	GTK_INPUT_PURPOSE_PASSWORD  InputPurpose = 8
	GTK_INPUT_PURPOSE_PIN       InputPurpose = 9
)

type LevelBarMode int

const (
	GTK_LEVEL_BAR_MODE_CONTINUOUS LevelBarMode = 0
	GTK_LEVEL_BAR_MODE_DISCRETE   LevelBarMode = 1
)
