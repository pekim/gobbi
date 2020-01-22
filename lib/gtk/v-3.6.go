// Code generated - DO NOT EDIT.
// +build gtk_3.6

package gtk

import "unsafe"

// LEVEL_BAR_OFFSET_HIGH is a representation of the C constant GTK_LEVEL_BAR_OFFSET_HIGH.
//
// since 3.6
const LEVEL_BAR_OFFSET_HIGH = "high"

// LEVEL_BAR_OFFSET_LOW is a representation of the C constant GTK_LEVEL_BAR_OFFSET_LOW.
//
// since 3.6
const LEVEL_BAR_OFFSET_LOW = "low"

// PRINT_SETTINGS_OUTPUT_BASENAME is a representation of the C constant GTK_PRINT_SETTINGS_OUTPUT_BASENAME.
//
// since 3.6
const PRINT_SETTINGS_OUTPUT_BASENAME = "output-basename"

// PRINT_SETTINGS_OUTPUT_DIR is a representation of the C constant GTK_PRINT_SETTINGS_OUTPUT_DIR.
//
// since 3.6
const PRINT_SETTINGS_OUTPUT_DIR = "output-dir"

// InputHints is a representation of the C bitfield GtkInputHints.
type InputHints int

// InputHints_none is a representation of the C bitfield member GTK_INPUT_HINT_NONE.
const InputHints_none = InputHints(0)

// InputHints_spellcheck is a representation of the C bitfield member GTK_INPUT_HINT_SPELLCHECK.
const InputHints_spellcheck = InputHints(1)

// InputHints_no_spellcheck is a representation of the C bitfield member GTK_INPUT_HINT_NO_SPELLCHECK.
const InputHints_no_spellcheck = InputHints(2)

// InputHints_word_completion is a representation of the C bitfield member GTK_INPUT_HINT_WORD_COMPLETION.
const InputHints_word_completion = InputHints(4)

// InputHints_lowercase is a representation of the C bitfield member GTK_INPUT_HINT_LOWERCASE.
const InputHints_lowercase = InputHints(8)

// InputHints_uppercase_chars is a representation of the C bitfield member GTK_INPUT_HINT_UPPERCASE_CHARS.
const InputHints_uppercase_chars = InputHints(16)

// InputHints_uppercase_words is a representation of the C bitfield member GTK_INPUT_HINT_UPPERCASE_WORDS.
const InputHints_uppercase_words = InputHints(32)

// InputHints_uppercase_sentences is a representation of the C bitfield member GTK_INPUT_HINT_UPPERCASE_SENTENCES.
const InputHints_uppercase_sentences = InputHints(64)

// InputHints_inhibit_osk is a representation of the C bitfield member GTK_INPUT_HINT_INHIBIT_OSK.
const InputHints_inhibit_osk = InputHints(128)

// InputHints_vertical_writing is a representation of the C bitfield member GTK_INPUT_HINT_VERTICAL_WRITING.
const InputHints_vertical_writing = InputHints(256)

// InputHints_emoji is a representation of the C bitfield member GTK_INPUT_HINT_EMOJI.
const InputHints_emoji = InputHints(512)

// InputHints_no_emoji is a representation of the C bitfield member GTK_INPUT_HINT_NO_EMOJI.
const InputHints_no_emoji = InputHints(1024)

// InputPurpose is a representation of the C enumeration GtkInputPurpose.
type InputPurpose int

// InputPurpose_free_form is a representation of the C enumeration member GTK_INPUT_PURPOSE_FREE_FORM.
const InputPurpose_free_form = InputPurpose(0)

// InputPurpose_alpha is a representation of the C enumeration member GTK_INPUT_PURPOSE_ALPHA.
const InputPurpose_alpha = InputPurpose(1)

// InputPurpose_digits is a representation of the C enumeration member GTK_INPUT_PURPOSE_DIGITS.
const InputPurpose_digits = InputPurpose(2)

// InputPurpose_number is a representation of the C enumeration member GTK_INPUT_PURPOSE_NUMBER.
const InputPurpose_number = InputPurpose(3)

// InputPurpose_phone is a representation of the C enumeration member GTK_INPUT_PURPOSE_PHONE.
const InputPurpose_phone = InputPurpose(4)

// InputPurpose_url is a representation of the C enumeration member GTK_INPUT_PURPOSE_URL.
const InputPurpose_url = InputPurpose(5)

// InputPurpose_email is a representation of the C enumeration member GTK_INPUT_PURPOSE_EMAIL.
const InputPurpose_email = InputPurpose(6)

// InputPurpose_name is a representation of the C enumeration member GTK_INPUT_PURPOSE_NAME.
const InputPurpose_name = InputPurpose(7)

// InputPurpose_password is a representation of the C enumeration member GTK_INPUT_PURPOSE_PASSWORD.
const InputPurpose_password = InputPurpose(8)

// InputPurpose_pin is a representation of the C enumeration member GTK_INPUT_PURPOSE_PIN.
const InputPurpose_pin = InputPurpose(9)

// LevelBarMode is a representation of the C enumeration GtkLevelBarMode.
type LevelBarMode int

// LevelBarMode_continuous is a representation of the C enumeration member GTK_LEVEL_BAR_MODE_CONTINUOUS.
const LevelBarMode_continuous = LevelBarMode(0)

// LevelBarMode_discrete is a representation of the C enumeration member GTK_LEVEL_BAR_MODE_DISCRETE.
const LevelBarMode_discrete = LevelBarMode(1)

// UNSUPPORTED : gtk_accelerator_parse : has [in]out param, accelerator_key

// UNSUPPORTED : gtk_accelerator_parse_with_keycode : parameter 'accelerator_codes' is array parameter without length parameter

// UNSUPPORTED : gtk_get_current_event_state : has [in]out param, state

// UNSUPPORTED : gtk_icon_size_lookup : has [in]out param, width

// UNSUPPORTED : gtk_icon_size_lookup_for_settings : has [in]out param, width

// UNSUPPORTED : gtk_init : has array param, argv

// UNSUPPORTED : gtk_init_check : has array param, argv

// UNSUPPORTED : gtk_init_with_args : parameter 'entries' is array parameter without length parameter

// UNSUPPORTED : gtk_key_snooper_install : parameter 'snooper' is callback

// UNSUPPORTED : gtk_parse_args : has array param, argv

// UNSUPPORTED : gtk_print_run_page_setup_dialog_async : parameter 'done_cb' is callback

// UNSUPPORTED : gtk_rc_get_default_files : no array length

// UNSUPPORTED : gtk_rc_parse_color : has [in]out param, color

// UNSUPPORTED : gtk_rc_parse_color_full : has [in]out param, color

// UNSUPPORTED : gtk_rc_parse_state : has [in]out param, state

// UNSUPPORTED : gtk_rc_set_default_files : parameter 'filenames' is array parameter without length parameter

// UNSUPPORTED : gtk_render_background_get_clip : has [in]out param, out_clip

// UNSUPPORTED : gtk_rgb_to_hsv : has [in]out param, h

// UNSUPPORTED : gtk_selection_add_targets : has array param, targets

// UNSUPPORTED : gtk_show_uri : throws

// UNSUPPORTED : gtk_show_uri_on_window : throws

// UNSUPPORTED : gtk_stock_add : has array param, items

// UNSUPPORTED : gtk_stock_add_static : has array param, items

// UNSUPPORTED : gtk_stock_lookup : has [in]out param, item

// UNSUPPORTED : gtk_stock_set_translate_func : parameter 'func' is callback

// UNSUPPORTED : gtk_target_table_free : has array param, targets

// UNSUPPORTED : gtk_target_table_new_from_list : has [in]out param, n_targets

// UNSUPPORTED : gtk_targets_include_image : has array param, targets

// UNSUPPORTED : gtk_targets_include_rich_text : has array param, targets

// UNSUPPORTED : gtk_targets_include_text : has array param, targets

// UNSUPPORTED : gtk_targets_include_uri : has array param, targets

// UNSUPPORTED : gtk_test_init : has array param, argvp

// UNSUPPORTED : gtk_test_list_all_types : has [in]out param, n_types

// UNSUPPORTED : gtk_tree_get_row_drag_data : has [in]out param, tree_model

// UNSUPPORTED : gtk_tree_row_reference_reordered : parameter 'new_order' is array parameter without length parameter

// UNSUPPORTED : EventControllerMotionClass : blacklisted

// UNSUPPORTED : EventControllerScrollClass : blacklisted

// UNSUPPORTED : GestureStylusClass : blacklisted

// UNSUPPORTED : HeaderBarAccessibleClass : blacklisted

// UNSUPPORTED : HeaderBarAccessiblePrivate : blacklisted

// UNSUPPORTED : StackAccessibleClass : blacklisted

// UNSUPPORTED : EntryIconAccessible : blacklisted

// UNSUPPORTED : EventControllerKey : blacklisted

// UNSUPPORTED : EventControllerMotion : blacklisted

// UNSUPPORTED : EventControllerScroll : blacklisted

// UNSUPPORTED : GestureStylus : blacklisted

// UNSUPPORTED : HeaderBarAccessible : blacklisted

// SearchEntry is a representation of the C record GtkSearchEntry.
//
// since 3.6
type SearchEntry struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkSearchEntry that represents the SearchEntry.
func (recv *SearchEntry) ToC() unsafe.Pointer {
	return recv.native
}

// SearchEntryNewFromC creates a new SearchEntry from a pointer to the C GtkSearchEntry that represents the SearchEntry.
func SearchEntryNewFromC(native unsafe.Pointer) *SearchEntry {
	return &SearchEntry{native: native}
}

// UNSUPPORTED : StackAccessible : blacklisted
