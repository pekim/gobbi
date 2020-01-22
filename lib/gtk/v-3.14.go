// Code generated - DO NOT EDIT.
// +build gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29 gtk_3.24

package gtk

// STYLE_CLASS_CSD is a representation of the C constant GTK_STYLE_CLASS_CSD.
//
// since 3.14
const STYLE_CLASS_CSD = "csd"

// STYLE_CLASS_FLAT is a representation of the C constant GTK_STYLE_CLASS_FLAT.
//
// since 3.14
const STYLE_CLASS_FLAT = "flat"

// STYLE_CLASS_MESSAGE_DIALOG is a representation of the C constant GTK_STYLE_CLASS_MESSAGE_DIALOG.
//
// since 3.14
const STYLE_CLASS_MESSAGE_DIALOG = "message-dialog"

// STYLE_CLASS_OVERSHOOT is a representation of the C constant GTK_STYLE_CLASS_OVERSHOOT.
//
// since 3.14
const STYLE_CLASS_OVERSHOOT = "overshoot"

// STYLE_CLASS_POPOVER is a representation of the C constant GTK_STYLE_CLASS_POPOVER.
//
// since 3.14
const STYLE_CLASS_POPOVER = "popover"

// STYLE_CLASS_POPUP is a representation of the C constant GTK_STYLE_CLASS_POPUP.
//
// since 3.14
const STYLE_CLASS_POPUP = "popup"

// STYLE_CLASS_SUBTITLE is a representation of the C constant GTK_STYLE_CLASS_SUBTITLE.
//
// since 3.14
const STYLE_CLASS_SUBTITLE = "subtitle"

// STYLE_CLASS_TITLE is a representation of the C constant GTK_STYLE_CLASS_TITLE.
//
// since 3.14
const STYLE_CLASS_TITLE = "title"

// EventSequenceState is a representation of the C enumeration GtkEventSequenceState.
type EventSequenceState int

// EventSequenceState_none is a representation of the C enumeration member GTK_EVENT_SEQUENCE_NONE.
const EventSequenceState_none = EventSequenceState(0)

// EventSequenceState_claimed is a representation of the C enumeration member GTK_EVENT_SEQUENCE_CLAIMED.
const EventSequenceState_claimed = EventSequenceState(1)

// EventSequenceState_denied is a representation of the C enumeration member GTK_EVENT_SEQUENCE_DENIED.
const EventSequenceState_denied = EventSequenceState(2)

// PanDirection is a representation of the C enumeration GtkPanDirection.
type PanDirection int

// PanDirection_left is a representation of the C enumeration member GTK_PAN_DIRECTION_LEFT.
const PanDirection_left = PanDirection(0)

// PanDirection_right is a representation of the C enumeration member GTK_PAN_DIRECTION_RIGHT.
const PanDirection_right = PanDirection(1)

// PanDirection_up is a representation of the C enumeration member GTK_PAN_DIRECTION_UP.
const PanDirection_up = PanDirection(2)

// PanDirection_down is a representation of the C enumeration member GTK_PAN_DIRECTION_DOWN.
const PanDirection_down = PanDirection(3)

// PropagationPhase is a representation of the C enumeration GtkPropagationPhase.
type PropagationPhase int

// PropagationPhase_none is a representation of the C enumeration member GTK_PHASE_NONE.
const PropagationPhase_none = PropagationPhase(0)

// PropagationPhase_capture is a representation of the C enumeration member GTK_PHASE_CAPTURE.
const PropagationPhase_capture = PropagationPhase(1)

// PropagationPhase_bubble is a representation of the C enumeration member GTK_PHASE_BUBBLE.
const PropagationPhase_bubble = PropagationPhase(2)

// PropagationPhase_target is a representation of the C enumeration member GTK_PHASE_TARGET.
const PropagationPhase_target = PropagationPhase(3)

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

// UNSUPPORTED : StackAccessible : blacklisted
