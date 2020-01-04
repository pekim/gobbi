// Code generated - DO NOT EDIT.
// +build atk_2.8

package atk

import (
	glib "github.com/pekim/gobbi/lib/glib"
	atk "github.com/pekim/gobbi/lib/internal/c/atk"
	"unsafe"
)

// AttributeSet is a representation of the C alias AtkAttributeSet.
type AttributeSet glib.SList

// State is a representation of the C alias AtkState.
type State uint64

// BINARY_AGE is a representation of the C constant ATK_BINARY_AGE.
//
// since 2.7.4
const BINARY_AGE = 23313

// INTERFACE_AGE is a representation of the C constant ATK_INTERFACE_AGE.
//
// since 2.7.4
const INTERFACE_AGE = 1

// MAJOR_VERSION is a representation of the C constant ATK_MAJOR_VERSION.
//
// since 2.7.4
const MAJOR_VERSION = 2

// MICRO_VERSION is a representation of the C constant ATK_MICRO_VERSION.
//
// since 2.7.4
const MICRO_VERSION = 3

// MINOR_VERSION is a representation of the C constant ATK_MINOR_VERSION.
//
// since 2.7.4
const MINOR_VERSION = 33

// HyperlinkStateFlags is a representation of the C bitfield AtkHyperlinkStateFlags.
type HyperlinkStateFlags int

// HyperlinkStateFlags_inline is a representation of the C bitfield member ATK_HYPERLINK_IS_INLINE.
const HyperlinkStateFlags_inline = HyperlinkStateFlags(1)

// CoordType is a representation of the C enumeration AtkCoordType.
type CoordType int

// CoordType_screen is a representation of the C enumeration member ATK_XY_SCREEN.
const CoordType_screen = CoordType(0)

// CoordType_window is a representation of the C enumeration member ATK_XY_WINDOW.
const CoordType_window = CoordType(1)

// CoordType_parent is a representation of the C enumeration member ATK_XY_PARENT.
const CoordType_parent = CoordType(2)

// KeyEventType is a representation of the C enumeration AtkKeyEventType.
type KeyEventType int

// KeyEventType_press is a representation of the C enumeration member ATK_KEY_EVENT_PRESS.
const KeyEventType_press = KeyEventType(0)

// KeyEventType_release is a representation of the C enumeration member ATK_KEY_EVENT_RELEASE.
const KeyEventType_release = KeyEventType(1)

// KeyEventType_last_defined is a representation of the C enumeration member ATK_KEY_EVENT_LAST_DEFINED.
const KeyEventType_last_defined = KeyEventType(2)

// Layer is a representation of the C enumeration AtkLayer.
type Layer int

// Layer_invalid is a representation of the C enumeration member ATK_LAYER_INVALID.
const Layer_invalid = Layer(0)

// Layer_background is a representation of the C enumeration member ATK_LAYER_BACKGROUND.
const Layer_background = Layer(1)

// Layer_canvas is a representation of the C enumeration member ATK_LAYER_CANVAS.
const Layer_canvas = Layer(2)

// Layer_widget is a representation of the C enumeration member ATK_LAYER_WIDGET.
const Layer_widget = Layer(3)

// Layer_mdi is a representation of the C enumeration member ATK_LAYER_MDI.
const Layer_mdi = Layer(4)

// Layer_popup is a representation of the C enumeration member ATK_LAYER_POPUP.
const Layer_popup = Layer(5)

// Layer_overlay is a representation of the C enumeration member ATK_LAYER_OVERLAY.
const Layer_overlay = Layer(6)

// Layer_window is a representation of the C enumeration member ATK_LAYER_WINDOW.
const Layer_window = Layer(7)

// RelationType is a representation of the C enumeration AtkRelationType.
type RelationType int

// RelationType_null is a representation of the C enumeration member ATK_RELATION_NULL.
const RelationType_null = RelationType(0)

// RelationType_controlled_by is a representation of the C enumeration member ATK_RELATION_CONTROLLED_BY.
const RelationType_controlled_by = RelationType(1)

// RelationType_controller_for is a representation of the C enumeration member ATK_RELATION_CONTROLLER_FOR.
const RelationType_controller_for = RelationType(2)

// RelationType_label_for is a representation of the C enumeration member ATK_RELATION_LABEL_FOR.
const RelationType_label_for = RelationType(3)

// RelationType_labelled_by is a representation of the C enumeration member ATK_RELATION_LABELLED_BY.
const RelationType_labelled_by = RelationType(4)

// RelationType_member_of is a representation of the C enumeration member ATK_RELATION_MEMBER_OF.
const RelationType_member_of = RelationType(5)

// RelationType_node_child_of is a representation of the C enumeration member ATK_RELATION_NODE_CHILD_OF.
const RelationType_node_child_of = RelationType(6)

// RelationType_flows_to is a representation of the C enumeration member ATK_RELATION_FLOWS_TO.
const RelationType_flows_to = RelationType(7)

// RelationType_flows_from is a representation of the C enumeration member ATK_RELATION_FLOWS_FROM.
const RelationType_flows_from = RelationType(8)

// RelationType_subwindow_of is a representation of the C enumeration member ATK_RELATION_SUBWINDOW_OF.
const RelationType_subwindow_of = RelationType(9)

// RelationType_embeds is a representation of the C enumeration member ATK_RELATION_EMBEDS.
const RelationType_embeds = RelationType(10)

// RelationType_embedded_by is a representation of the C enumeration member ATK_RELATION_EMBEDDED_BY.
const RelationType_embedded_by = RelationType(11)

// RelationType_popup_for is a representation of the C enumeration member ATK_RELATION_POPUP_FOR.
const RelationType_popup_for = RelationType(12)

// RelationType_parent_window_of is a representation of the C enumeration member ATK_RELATION_PARENT_WINDOW_OF.
const RelationType_parent_window_of = RelationType(13)

// RelationType_described_by is a representation of the C enumeration member ATK_RELATION_DESCRIBED_BY.
const RelationType_described_by = RelationType(14)

// RelationType_description_for is a representation of the C enumeration member ATK_RELATION_DESCRIPTION_FOR.
const RelationType_description_for = RelationType(15)

// RelationType_node_parent_of is a representation of the C enumeration member ATK_RELATION_NODE_PARENT_OF.
const RelationType_node_parent_of = RelationType(16)

// RelationType_details is a representation of the C enumeration member ATK_RELATION_DETAILS.
const RelationType_details = RelationType(17)

// RelationType_details_for is a representation of the C enumeration member ATK_RELATION_DETAILS_FOR.
const RelationType_details_for = RelationType(18)

// RelationType_error_message is a representation of the C enumeration member ATK_RELATION_ERROR_MESSAGE.
const RelationType_error_message = RelationType(19)

// RelationType_error_for is a representation of the C enumeration member ATK_RELATION_ERROR_FOR.
const RelationType_error_for = RelationType(20)

// RelationType_last_defined is a representation of the C enumeration member ATK_RELATION_LAST_DEFINED.
const RelationType_last_defined = RelationType(21)

// Role is a representation of the C enumeration AtkRole.
type Role int

// Role_invalid is a representation of the C enumeration member ATK_ROLE_INVALID.
const Role_invalid = Role(0)

// Role_accelerator_label is a representation of the C enumeration member ATK_ROLE_ACCEL_LABEL.
const Role_accelerator_label = Role(1)

// Role_alert is a representation of the C enumeration member ATK_ROLE_ALERT.
const Role_alert = Role(2)

// Role_animation is a representation of the C enumeration member ATK_ROLE_ANIMATION.
const Role_animation = Role(3)

// Role_arrow is a representation of the C enumeration member ATK_ROLE_ARROW.
const Role_arrow = Role(4)

// Role_calendar is a representation of the C enumeration member ATK_ROLE_CALENDAR.
const Role_calendar = Role(5)

// Role_canvas is a representation of the C enumeration member ATK_ROLE_CANVAS.
const Role_canvas = Role(6)

// Role_check_box is a representation of the C enumeration member ATK_ROLE_CHECK_BOX.
const Role_check_box = Role(7)

// Role_check_menu_item is a representation of the C enumeration member ATK_ROLE_CHECK_MENU_ITEM.
const Role_check_menu_item = Role(8)

// Role_color_chooser is a representation of the C enumeration member ATK_ROLE_COLOR_CHOOSER.
const Role_color_chooser = Role(9)

// Role_column_header is a representation of the C enumeration member ATK_ROLE_COLUMN_HEADER.
const Role_column_header = Role(10)

// Role_combo_box is a representation of the C enumeration member ATK_ROLE_COMBO_BOX.
const Role_combo_box = Role(11)

// Role_date_editor is a representation of the C enumeration member ATK_ROLE_DATE_EDITOR.
const Role_date_editor = Role(12)

// Role_desktop_icon is a representation of the C enumeration member ATK_ROLE_DESKTOP_ICON.
const Role_desktop_icon = Role(13)

// Role_desktop_frame is a representation of the C enumeration member ATK_ROLE_DESKTOP_FRAME.
const Role_desktop_frame = Role(14)

// Role_dial is a representation of the C enumeration member ATK_ROLE_DIAL.
const Role_dial = Role(15)

// Role_dialog is a representation of the C enumeration member ATK_ROLE_DIALOG.
const Role_dialog = Role(16)

// Role_directory_pane is a representation of the C enumeration member ATK_ROLE_DIRECTORY_PANE.
const Role_directory_pane = Role(17)

// Role_drawing_area is a representation of the C enumeration member ATK_ROLE_DRAWING_AREA.
const Role_drawing_area = Role(18)

// Role_file_chooser is a representation of the C enumeration member ATK_ROLE_FILE_CHOOSER.
const Role_file_chooser = Role(19)

// Role_filler is a representation of the C enumeration member ATK_ROLE_FILLER.
const Role_filler = Role(20)

// Role_font_chooser is a representation of the C enumeration member ATK_ROLE_FONT_CHOOSER.
const Role_font_chooser = Role(21)

// Role_frame is a representation of the C enumeration member ATK_ROLE_FRAME.
const Role_frame = Role(22)

// Role_glass_pane is a representation of the C enumeration member ATK_ROLE_GLASS_PANE.
const Role_glass_pane = Role(23)

// Role_html_container is a representation of the C enumeration member ATK_ROLE_HTML_CONTAINER.
const Role_html_container = Role(24)

// Role_icon is a representation of the C enumeration member ATK_ROLE_ICON.
const Role_icon = Role(25)

// Role_image is a representation of the C enumeration member ATK_ROLE_IMAGE.
const Role_image = Role(26)

// Role_internal_frame is a representation of the C enumeration member ATK_ROLE_INTERNAL_FRAME.
const Role_internal_frame = Role(27)

// Role_label is a representation of the C enumeration member ATK_ROLE_LABEL.
const Role_label = Role(28)

// Role_layered_pane is a representation of the C enumeration member ATK_ROLE_LAYERED_PANE.
const Role_layered_pane = Role(29)

// Role_list is a representation of the C enumeration member ATK_ROLE_LIST.
const Role_list = Role(30)

// Role_list_item is a representation of the C enumeration member ATK_ROLE_LIST_ITEM.
const Role_list_item = Role(31)

// Role_menu is a representation of the C enumeration member ATK_ROLE_MENU.
const Role_menu = Role(32)

// Role_menu_bar is a representation of the C enumeration member ATK_ROLE_MENU_BAR.
const Role_menu_bar = Role(33)

// Role_menu_item is a representation of the C enumeration member ATK_ROLE_MENU_ITEM.
const Role_menu_item = Role(34)

// Role_option_pane is a representation of the C enumeration member ATK_ROLE_OPTION_PANE.
const Role_option_pane = Role(35)

// Role_page_tab is a representation of the C enumeration member ATK_ROLE_PAGE_TAB.
const Role_page_tab = Role(36)

// Role_page_tab_list is a representation of the C enumeration member ATK_ROLE_PAGE_TAB_LIST.
const Role_page_tab_list = Role(37)

// Role_panel is a representation of the C enumeration member ATK_ROLE_PANEL.
const Role_panel = Role(38)

// Role_password_text is a representation of the C enumeration member ATK_ROLE_PASSWORD_TEXT.
const Role_password_text = Role(39)

// Role_popup_menu is a representation of the C enumeration member ATK_ROLE_POPUP_MENU.
const Role_popup_menu = Role(40)

// Role_progress_bar is a representation of the C enumeration member ATK_ROLE_PROGRESS_BAR.
const Role_progress_bar = Role(41)

// Role_push_button is a representation of the C enumeration member ATK_ROLE_PUSH_BUTTON.
const Role_push_button = Role(42)

// Role_radio_button is a representation of the C enumeration member ATK_ROLE_RADIO_BUTTON.
const Role_radio_button = Role(43)

// Role_radio_menu_item is a representation of the C enumeration member ATK_ROLE_RADIO_MENU_ITEM.
const Role_radio_menu_item = Role(44)

// Role_root_pane is a representation of the C enumeration member ATK_ROLE_ROOT_PANE.
const Role_root_pane = Role(45)

// Role_row_header is a representation of the C enumeration member ATK_ROLE_ROW_HEADER.
const Role_row_header = Role(46)

// Role_scroll_bar is a representation of the C enumeration member ATK_ROLE_SCROLL_BAR.
const Role_scroll_bar = Role(47)

// Role_scroll_pane is a representation of the C enumeration member ATK_ROLE_SCROLL_PANE.
const Role_scroll_pane = Role(48)

// Role_separator is a representation of the C enumeration member ATK_ROLE_SEPARATOR.
const Role_separator = Role(49)

// Role_slider is a representation of the C enumeration member ATK_ROLE_SLIDER.
const Role_slider = Role(50)

// Role_split_pane is a representation of the C enumeration member ATK_ROLE_SPLIT_PANE.
const Role_split_pane = Role(51)

// Role_spin_button is a representation of the C enumeration member ATK_ROLE_SPIN_BUTTON.
const Role_spin_button = Role(52)

// Role_statusbar is a representation of the C enumeration member ATK_ROLE_STATUSBAR.
const Role_statusbar = Role(53)

// Role_table is a representation of the C enumeration member ATK_ROLE_TABLE.
const Role_table = Role(54)

// Role_table_cell is a representation of the C enumeration member ATK_ROLE_TABLE_CELL.
const Role_table_cell = Role(55)

// Role_table_column_header is a representation of the C enumeration member ATK_ROLE_TABLE_COLUMN_HEADER.
const Role_table_column_header = Role(56)

// Role_table_row_header is a representation of the C enumeration member ATK_ROLE_TABLE_ROW_HEADER.
const Role_table_row_header = Role(57)

// Role_tear_off_menu_item is a representation of the C enumeration member ATK_ROLE_TEAR_OFF_MENU_ITEM.
const Role_tear_off_menu_item = Role(58)

// Role_terminal is a representation of the C enumeration member ATK_ROLE_TERMINAL.
const Role_terminal = Role(59)

// Role_text is a representation of the C enumeration member ATK_ROLE_TEXT.
const Role_text = Role(60)

// Role_toggle_button is a representation of the C enumeration member ATK_ROLE_TOGGLE_BUTTON.
const Role_toggle_button = Role(61)

// Role_tool_bar is a representation of the C enumeration member ATK_ROLE_TOOL_BAR.
const Role_tool_bar = Role(62)

// Role_tool_tip is a representation of the C enumeration member ATK_ROLE_TOOL_TIP.
const Role_tool_tip = Role(63)

// Role_tree is a representation of the C enumeration member ATK_ROLE_TREE.
const Role_tree = Role(64)

// Role_tree_table is a representation of the C enumeration member ATK_ROLE_TREE_TABLE.
const Role_tree_table = Role(65)

// Role_unknown is a representation of the C enumeration member ATK_ROLE_UNKNOWN.
const Role_unknown = Role(66)

// Role_viewport is a representation of the C enumeration member ATK_ROLE_VIEWPORT.
const Role_viewport = Role(67)

// Role_window is a representation of the C enumeration member ATK_ROLE_WINDOW.
const Role_window = Role(68)

// Role_header is a representation of the C enumeration member ATK_ROLE_HEADER.
const Role_header = Role(69)

// Role_footer is a representation of the C enumeration member ATK_ROLE_FOOTER.
const Role_footer = Role(70)

// Role_paragraph is a representation of the C enumeration member ATK_ROLE_PARAGRAPH.
const Role_paragraph = Role(71)

// Role_ruler is a representation of the C enumeration member ATK_ROLE_RULER.
const Role_ruler = Role(72)

// Role_application is a representation of the C enumeration member ATK_ROLE_APPLICATION.
const Role_application = Role(73)

// Role_autocomplete is a representation of the C enumeration member ATK_ROLE_AUTOCOMPLETE.
const Role_autocomplete = Role(74)

// Role_edit_bar is a representation of the C enumeration member ATK_ROLE_EDITBAR.
const Role_edit_bar = Role(75)

// Role_embedded is a representation of the C enumeration member ATK_ROLE_EMBEDDED.
const Role_embedded = Role(76)

// Role_entry is a representation of the C enumeration member ATK_ROLE_ENTRY.
const Role_entry = Role(77)

// Role_chart is a representation of the C enumeration member ATK_ROLE_CHART.
const Role_chart = Role(78)

// Role_caption is a representation of the C enumeration member ATK_ROLE_CAPTION.
const Role_caption = Role(79)

// Role_document_frame is a representation of the C enumeration member ATK_ROLE_DOCUMENT_FRAME.
const Role_document_frame = Role(80)

// Role_heading is a representation of the C enumeration member ATK_ROLE_HEADING.
const Role_heading = Role(81)

// Role_page is a representation of the C enumeration member ATK_ROLE_PAGE.
const Role_page = Role(82)

// Role_section is a representation of the C enumeration member ATK_ROLE_SECTION.
const Role_section = Role(83)

// Role_redundant_object is a representation of the C enumeration member ATK_ROLE_REDUNDANT_OBJECT.
const Role_redundant_object = Role(84)

// Role_form is a representation of the C enumeration member ATK_ROLE_FORM.
const Role_form = Role(85)

// Role_link is a representation of the C enumeration member ATK_ROLE_LINK.
const Role_link = Role(86)

// Role_input_method_window is a representation of the C enumeration member ATK_ROLE_INPUT_METHOD_WINDOW.
const Role_input_method_window = Role(87)

// Role_table_row is a representation of the C enumeration member ATK_ROLE_TABLE_ROW.
const Role_table_row = Role(88)

// Role_tree_item is a representation of the C enumeration member ATK_ROLE_TREE_ITEM.
const Role_tree_item = Role(89)

// Role_document_spreadsheet is a representation of the C enumeration member ATK_ROLE_DOCUMENT_SPREADSHEET.
const Role_document_spreadsheet = Role(90)

// Role_document_presentation is a representation of the C enumeration member ATK_ROLE_DOCUMENT_PRESENTATION.
const Role_document_presentation = Role(91)

// Role_document_text is a representation of the C enumeration member ATK_ROLE_DOCUMENT_TEXT.
const Role_document_text = Role(92)

// Role_document_web is a representation of the C enumeration member ATK_ROLE_DOCUMENT_WEB.
const Role_document_web = Role(93)

// Role_document_email is a representation of the C enumeration member ATK_ROLE_DOCUMENT_EMAIL.
const Role_document_email = Role(94)

// Role_comment is a representation of the C enumeration member ATK_ROLE_COMMENT.
const Role_comment = Role(95)

// Role_list_box is a representation of the C enumeration member ATK_ROLE_LIST_BOX.
const Role_list_box = Role(96)

// Role_grouping is a representation of the C enumeration member ATK_ROLE_GROUPING.
const Role_grouping = Role(97)

// Role_image_map is a representation of the C enumeration member ATK_ROLE_IMAGE_MAP.
const Role_image_map = Role(98)

// Role_notification is a representation of the C enumeration member ATK_ROLE_NOTIFICATION.
const Role_notification = Role(99)

// Role_info_bar is a representation of the C enumeration member ATK_ROLE_INFO_BAR.
const Role_info_bar = Role(100)

// Role_level_bar is a representation of the C enumeration member ATK_ROLE_LEVEL_BAR.
const Role_level_bar = Role(101)

// Role_title_bar is a representation of the C enumeration member ATK_ROLE_TITLE_BAR.
const Role_title_bar = Role(102)

// Role_block_quote is a representation of the C enumeration member ATK_ROLE_BLOCK_QUOTE.
const Role_block_quote = Role(103)

// Role_audio is a representation of the C enumeration member ATK_ROLE_AUDIO.
const Role_audio = Role(104)

// Role_video is a representation of the C enumeration member ATK_ROLE_VIDEO.
const Role_video = Role(105)

// Role_definition is a representation of the C enumeration member ATK_ROLE_DEFINITION.
const Role_definition = Role(106)

// Role_article is a representation of the C enumeration member ATK_ROLE_ARTICLE.
const Role_article = Role(107)

// Role_landmark is a representation of the C enumeration member ATK_ROLE_LANDMARK.
const Role_landmark = Role(108)

// Role_log is a representation of the C enumeration member ATK_ROLE_LOG.
const Role_log = Role(109)

// Role_marquee is a representation of the C enumeration member ATK_ROLE_MARQUEE.
const Role_marquee = Role(110)

// Role_math is a representation of the C enumeration member ATK_ROLE_MATH.
const Role_math = Role(111)

// Role_rating is a representation of the C enumeration member ATK_ROLE_RATING.
const Role_rating = Role(112)

// Role_timer is a representation of the C enumeration member ATK_ROLE_TIMER.
const Role_timer = Role(113)

// Role_description_list is a representation of the C enumeration member ATK_ROLE_DESCRIPTION_LIST.
const Role_description_list = Role(114)

// Role_description_term is a representation of the C enumeration member ATK_ROLE_DESCRIPTION_TERM.
const Role_description_term = Role(115)

// Role_description_value is a representation of the C enumeration member ATK_ROLE_DESCRIPTION_VALUE.
const Role_description_value = Role(116)

// Role_static is a representation of the C enumeration member ATK_ROLE_STATIC.
const Role_static = Role(117)

// Role_math_fraction is a representation of the C enumeration member ATK_ROLE_MATH_FRACTION.
const Role_math_fraction = Role(118)

// Role_math_root is a representation of the C enumeration member ATK_ROLE_MATH_ROOT.
const Role_math_root = Role(119)

// Role_subscript is a representation of the C enumeration member ATK_ROLE_SUBSCRIPT.
const Role_subscript = Role(120)

// Role_superscript is a representation of the C enumeration member ATK_ROLE_SUPERSCRIPT.
const Role_superscript = Role(121)

// Role_footnote is a representation of the C enumeration member ATK_ROLE_FOOTNOTE.
const Role_footnote = Role(122)

// Role_content_deletion is a representation of the C enumeration member ATK_ROLE_CONTENT_DELETION.
const Role_content_deletion = Role(123)

// Role_content_insertion is a representation of the C enumeration member ATK_ROLE_CONTENT_INSERTION.
const Role_content_insertion = Role(124)

// Role_last_defined is a representation of the C enumeration member ATK_ROLE_LAST_DEFINED.
const Role_last_defined = Role(125)

// StateType is a representation of the C enumeration AtkStateType.
type StateType int

// StateType_invalid is a representation of the C enumeration member ATK_STATE_INVALID.
const StateType_invalid = StateType(0)

// StateType_active is a representation of the C enumeration member ATK_STATE_ACTIVE.
const StateType_active = StateType(1)

// StateType_armed is a representation of the C enumeration member ATK_STATE_ARMED.
const StateType_armed = StateType(2)

// StateType_busy is a representation of the C enumeration member ATK_STATE_BUSY.
const StateType_busy = StateType(3)

// StateType_checked is a representation of the C enumeration member ATK_STATE_CHECKED.
const StateType_checked = StateType(4)

// StateType_defunct is a representation of the C enumeration member ATK_STATE_DEFUNCT.
const StateType_defunct = StateType(5)

// StateType_editable is a representation of the C enumeration member ATK_STATE_EDITABLE.
const StateType_editable = StateType(6)

// StateType_enabled is a representation of the C enumeration member ATK_STATE_ENABLED.
const StateType_enabled = StateType(7)

// StateType_expandable is a representation of the C enumeration member ATK_STATE_EXPANDABLE.
const StateType_expandable = StateType(8)

// StateType_expanded is a representation of the C enumeration member ATK_STATE_EXPANDED.
const StateType_expanded = StateType(9)

// StateType_focusable is a representation of the C enumeration member ATK_STATE_FOCUSABLE.
const StateType_focusable = StateType(10)

// StateType_focused is a representation of the C enumeration member ATK_STATE_FOCUSED.
const StateType_focused = StateType(11)

// StateType_horizontal is a representation of the C enumeration member ATK_STATE_HORIZONTAL.
const StateType_horizontal = StateType(12)

// StateType_iconified is a representation of the C enumeration member ATK_STATE_ICONIFIED.
const StateType_iconified = StateType(13)

// StateType_modal is a representation of the C enumeration member ATK_STATE_MODAL.
const StateType_modal = StateType(14)

// StateType_multi_line is a representation of the C enumeration member ATK_STATE_MULTI_LINE.
const StateType_multi_line = StateType(15)

// StateType_multiselectable is a representation of the C enumeration member ATK_STATE_MULTISELECTABLE.
const StateType_multiselectable = StateType(16)

// StateType_opaque is a representation of the C enumeration member ATK_STATE_OPAQUE.
const StateType_opaque = StateType(17)

// StateType_pressed is a representation of the C enumeration member ATK_STATE_PRESSED.
const StateType_pressed = StateType(18)

// StateType_resizable is a representation of the C enumeration member ATK_STATE_RESIZABLE.
const StateType_resizable = StateType(19)

// StateType_selectable is a representation of the C enumeration member ATK_STATE_SELECTABLE.
const StateType_selectable = StateType(20)

// StateType_selected is a representation of the C enumeration member ATK_STATE_SELECTED.
const StateType_selected = StateType(21)

// StateType_sensitive is a representation of the C enumeration member ATK_STATE_SENSITIVE.
const StateType_sensitive = StateType(22)

// StateType_showing is a representation of the C enumeration member ATK_STATE_SHOWING.
const StateType_showing = StateType(23)

// StateType_single_line is a representation of the C enumeration member ATK_STATE_SINGLE_LINE.
const StateType_single_line = StateType(24)

// StateType_stale is a representation of the C enumeration member ATK_STATE_STALE.
const StateType_stale = StateType(25)

// StateType_transient is a representation of the C enumeration member ATK_STATE_TRANSIENT.
const StateType_transient = StateType(26)

// StateType_vertical is a representation of the C enumeration member ATK_STATE_VERTICAL.
const StateType_vertical = StateType(27)

// StateType_visible is a representation of the C enumeration member ATK_STATE_VISIBLE.
const StateType_visible = StateType(28)

// StateType_manages_descendants is a representation of the C enumeration member ATK_STATE_MANAGES_DESCENDANTS.
const StateType_manages_descendants = StateType(29)

// StateType_indeterminate is a representation of the C enumeration member ATK_STATE_INDETERMINATE.
const StateType_indeterminate = StateType(30)

// StateType_truncated is a representation of the C enumeration member ATK_STATE_TRUNCATED.
const StateType_truncated = StateType(31)

// StateType_required is a representation of the C enumeration member ATK_STATE_REQUIRED.
const StateType_required = StateType(32)

// StateType_invalid_entry is a representation of the C enumeration member ATK_STATE_INVALID_ENTRY.
const StateType_invalid_entry = StateType(33)

// StateType_supports_autocompletion is a representation of the C enumeration member ATK_STATE_SUPPORTS_AUTOCOMPLETION.
const StateType_supports_autocompletion = StateType(34)

// StateType_selectable_text is a representation of the C enumeration member ATK_STATE_SELECTABLE_TEXT.
const StateType_selectable_text = StateType(35)

// StateType_default is a representation of the C enumeration member ATK_STATE_DEFAULT.
const StateType_default = StateType(36)

// StateType_animated is a representation of the C enumeration member ATK_STATE_ANIMATED.
const StateType_animated = StateType(37)

// StateType_visited is a representation of the C enumeration member ATK_STATE_VISITED.
const StateType_visited = StateType(38)

// StateType_checkable is a representation of the C enumeration member ATK_STATE_CHECKABLE.
const StateType_checkable = StateType(39)

// StateType_has_popup is a representation of the C enumeration member ATK_STATE_HAS_POPUP.
const StateType_has_popup = StateType(40)

// StateType_has_tooltip is a representation of the C enumeration member ATK_STATE_HAS_TOOLTIP.
const StateType_has_tooltip = StateType(41)

// StateType_read_only is a representation of the C enumeration member ATK_STATE_READ_ONLY.
const StateType_read_only = StateType(42)

// StateType_last_defined is a representation of the C enumeration member ATK_STATE_LAST_DEFINED.
const StateType_last_defined = StateType(43)

// TextAttribute is a representation of the C enumeration AtkTextAttribute.
type TextAttribute int

// TextAttribute_invalid is a representation of the C enumeration member ATK_TEXT_ATTR_INVALID.
const TextAttribute_invalid = TextAttribute(0)

// TextAttribute_left_margin is a representation of the C enumeration member ATK_TEXT_ATTR_LEFT_MARGIN.
const TextAttribute_left_margin = TextAttribute(1)

// TextAttribute_right_margin is a representation of the C enumeration member ATK_TEXT_ATTR_RIGHT_MARGIN.
const TextAttribute_right_margin = TextAttribute(2)

// TextAttribute_indent is a representation of the C enumeration member ATK_TEXT_ATTR_INDENT.
const TextAttribute_indent = TextAttribute(3)

// TextAttribute_invisible is a representation of the C enumeration member ATK_TEXT_ATTR_INVISIBLE.
const TextAttribute_invisible = TextAttribute(4)

// TextAttribute_editable is a representation of the C enumeration member ATK_TEXT_ATTR_EDITABLE.
const TextAttribute_editable = TextAttribute(5)

// TextAttribute_pixels_above_lines is a representation of the C enumeration member ATK_TEXT_ATTR_PIXELS_ABOVE_LINES.
const TextAttribute_pixels_above_lines = TextAttribute(6)

// TextAttribute_pixels_below_lines is a representation of the C enumeration member ATK_TEXT_ATTR_PIXELS_BELOW_LINES.
const TextAttribute_pixels_below_lines = TextAttribute(7)

// TextAttribute_pixels_inside_wrap is a representation of the C enumeration member ATK_TEXT_ATTR_PIXELS_INSIDE_WRAP.
const TextAttribute_pixels_inside_wrap = TextAttribute(8)

// TextAttribute_bg_full_height is a representation of the C enumeration member ATK_TEXT_ATTR_BG_FULL_HEIGHT.
const TextAttribute_bg_full_height = TextAttribute(9)

// TextAttribute_rise is a representation of the C enumeration member ATK_TEXT_ATTR_RISE.
const TextAttribute_rise = TextAttribute(10)

// TextAttribute_underline is a representation of the C enumeration member ATK_TEXT_ATTR_UNDERLINE.
const TextAttribute_underline = TextAttribute(11)

// TextAttribute_strikethrough is a representation of the C enumeration member ATK_TEXT_ATTR_STRIKETHROUGH.
const TextAttribute_strikethrough = TextAttribute(12)

// TextAttribute_size is a representation of the C enumeration member ATK_TEXT_ATTR_SIZE.
const TextAttribute_size = TextAttribute(13)

// TextAttribute_scale is a representation of the C enumeration member ATK_TEXT_ATTR_SCALE.
const TextAttribute_scale = TextAttribute(14)

// TextAttribute_weight is a representation of the C enumeration member ATK_TEXT_ATTR_WEIGHT.
const TextAttribute_weight = TextAttribute(15)

// TextAttribute_language is a representation of the C enumeration member ATK_TEXT_ATTR_LANGUAGE.
const TextAttribute_language = TextAttribute(16)

// TextAttribute_family_name is a representation of the C enumeration member ATK_TEXT_ATTR_FAMILY_NAME.
const TextAttribute_family_name = TextAttribute(17)

// TextAttribute_bg_color is a representation of the C enumeration member ATK_TEXT_ATTR_BG_COLOR.
const TextAttribute_bg_color = TextAttribute(18)

// TextAttribute_fg_color is a representation of the C enumeration member ATK_TEXT_ATTR_FG_COLOR.
const TextAttribute_fg_color = TextAttribute(19)

// TextAttribute_bg_stipple is a representation of the C enumeration member ATK_TEXT_ATTR_BG_STIPPLE.
const TextAttribute_bg_stipple = TextAttribute(20)

// TextAttribute_fg_stipple is a representation of the C enumeration member ATK_TEXT_ATTR_FG_STIPPLE.
const TextAttribute_fg_stipple = TextAttribute(21)

// TextAttribute_wrap_mode is a representation of the C enumeration member ATK_TEXT_ATTR_WRAP_MODE.
const TextAttribute_wrap_mode = TextAttribute(22)

// TextAttribute_direction is a representation of the C enumeration member ATK_TEXT_ATTR_DIRECTION.
const TextAttribute_direction = TextAttribute(23)

// TextAttribute_justification is a representation of the C enumeration member ATK_TEXT_ATTR_JUSTIFICATION.
const TextAttribute_justification = TextAttribute(24)

// TextAttribute_stretch is a representation of the C enumeration member ATK_TEXT_ATTR_STRETCH.
const TextAttribute_stretch = TextAttribute(25)

// TextAttribute_variant is a representation of the C enumeration member ATK_TEXT_ATTR_VARIANT.
const TextAttribute_variant = TextAttribute(26)

// TextAttribute_style is a representation of the C enumeration member ATK_TEXT_ATTR_STYLE.
const TextAttribute_style = TextAttribute(27)

// TextAttribute_last_defined is a representation of the C enumeration member ATK_TEXT_ATTR_LAST_DEFINED.
const TextAttribute_last_defined = TextAttribute(28)

// TextBoundary is a representation of the C enumeration AtkTextBoundary.
type TextBoundary int

// TextBoundary_char is a representation of the C enumeration member ATK_TEXT_BOUNDARY_CHAR.
const TextBoundary_char = TextBoundary(0)

// TextBoundary_word_start is a representation of the C enumeration member ATK_TEXT_BOUNDARY_WORD_START.
const TextBoundary_word_start = TextBoundary(1)

// TextBoundary_word_end is a representation of the C enumeration member ATK_TEXT_BOUNDARY_WORD_END.
const TextBoundary_word_end = TextBoundary(2)

// TextBoundary_sentence_start is a representation of the C enumeration member ATK_TEXT_BOUNDARY_SENTENCE_START.
const TextBoundary_sentence_start = TextBoundary(3)

// TextBoundary_sentence_end is a representation of the C enumeration member ATK_TEXT_BOUNDARY_SENTENCE_END.
const TextBoundary_sentence_end = TextBoundary(4)

// TextBoundary_line_start is a representation of the C enumeration member ATK_TEXT_BOUNDARY_LINE_START.
const TextBoundary_line_start = TextBoundary(5)

// TextBoundary_line_end is a representation of the C enumeration member ATK_TEXT_BOUNDARY_LINE_END.
const TextBoundary_line_end = TextBoundary(6)

// TextClipType is a representation of the C enumeration AtkTextClipType.
type TextClipType int

// TextClipType_none is a representation of the C enumeration member ATK_TEXT_CLIP_NONE.
const TextClipType_none = TextClipType(0)

// TextClipType_min is a representation of the C enumeration member ATK_TEXT_CLIP_MIN.
const TextClipType_min = TextClipType(1)

// TextClipType_max is a representation of the C enumeration member ATK_TEXT_CLIP_MAX.
const TextClipType_max = TextClipType(2)

// TextClipType_both is a representation of the C enumeration member ATK_TEXT_CLIP_BOTH.
const TextClipType_both = TextClipType(3)

// TextGranularity is a representation of the C enumeration AtkTextGranularity.
type TextGranularity int

// TextGranularity_char is a representation of the C enumeration member ATK_TEXT_GRANULARITY_CHAR.
const TextGranularity_char = TextGranularity(0)

// TextGranularity_word is a representation of the C enumeration member ATK_TEXT_GRANULARITY_WORD.
const TextGranularity_word = TextGranularity(1)

// TextGranularity_sentence is a representation of the C enumeration member ATK_TEXT_GRANULARITY_SENTENCE.
const TextGranularity_sentence = TextGranularity(2)

// TextGranularity_line is a representation of the C enumeration member ATK_TEXT_GRANULARITY_LINE.
const TextGranularity_line = TextGranularity(3)

// TextGranularity_paragraph is a representation of the C enumeration member ATK_TEXT_GRANULARITY_PARAGRAPH.
const TextGranularity_paragraph = TextGranularity(4)

// ValueType is a representation of the C enumeration AtkValueType.
type ValueType int

// ValueType_very_weak is a representation of the C enumeration member ATK_VALUE_VERY_WEAK.
const ValueType_very_weak = ValueType(0)

// ValueType_weak is a representation of the C enumeration member ATK_VALUE_WEAK.
const ValueType_weak = ValueType(1)

// ValueType_acceptable is a representation of the C enumeration member ATK_VALUE_ACCEPTABLE.
const ValueType_acceptable = ValueType(2)

// ValueType_strong is a representation of the C enumeration member ATK_VALUE_STRONG.
const ValueType_strong = ValueType(3)

// ValueType_very_strong is a representation of the C enumeration member ATK_VALUE_VERY_STRONG.
const ValueType_very_strong = ValueType(4)

// ValueType_very_low is a representation of the C enumeration member ATK_VALUE_VERY_LOW.
const ValueType_very_low = ValueType(5)

// ValueType_low is a representation of the C enumeration member ATK_VALUE_LOW.
const ValueType_low = ValueType(6)

// ValueType_medium is a representation of the C enumeration member ATK_VALUE_MEDIUM.
const ValueType_medium = ValueType(7)

// ValueType_high is a representation of the C enumeration member ATK_VALUE_HIGH.
const ValueType_high = ValueType(8)

// ValueType_very_high is a representation of the C enumeration member ATK_VALUE_VERY_HIGH.
const ValueType_very_high = ValueType(9)

// ValueType_very_bad is a representation of the C enumeration member ATK_VALUE_VERY_BAD.
const ValueType_very_bad = ValueType(10)

// ValueType_bad is a representation of the C enumeration member ATK_VALUE_BAD.
const ValueType_bad = ValueType(11)

// ValueType_good is a representation of the C enumeration member ATK_VALUE_GOOD.
const ValueType_good = ValueType(12)

// ValueType_very_good is a representation of the C enumeration member ATK_VALUE_VERY_GOOD.
const ValueType_very_good = ValueType(13)

// ValueType_best is a representation of the C enumeration member ATK_VALUE_BEST.
const ValueType_best = ValueType(14)

// ValueType_last_defined is a representation of the C enumeration member ATK_VALUE_LAST_DEFINED.
const ValueType_last_defined = ValueType(15)

// UNSUPPORTED : atk_add_focus_tracker : parameter 'focus_tracker' is callback

// UNSUPPORTED : atk_add_global_event_listener : parameter 'listener' is callback

// UNSUPPORTED : atk_add_key_event_listener : parameter 'listener' is callback

// UNSUPPORTED : atk_focus_tracker_init : parameter 'init' is callback

// FocusTrackerNotify is analogous to the C function atk_focus_tracker_notify.
func FocusTrackerNotify(object *Object) {
	sys_object := object.ToC()
	atk.Fn_atk_focus_tracker_notify(sys_object)
}

// GetBinaryAge is analogous to the C function atk_get_binary_age.
//
// since 2.8
func GetBinaryAge() uint {
	retSys := atk.Fn_atk_get_binary_age()
	ret := retSys

	return ret
}

// GetDefaultRegistry is analogous to the C function atk_get_default_registry.
func GetDefaultRegistry() *Registry {
	retSys := atk.Fn_atk_get_default_registry()
	ret := RegistryNewFromC(retSys)

	return ret
}

// GetFocusObject is analogous to the C function atk_get_focus_object.
//
// since 1.6
func GetFocusObject() *Object {
	retSys := atk.Fn_atk_get_focus_object()
	ret := ObjectNewFromC(retSys)

	return ret
}

// GetInterfaceAge is analogous to the C function atk_get_interface_age.
//
// since 2.8
func GetInterfaceAge() uint {
	retSys := atk.Fn_atk_get_interface_age()
	ret := retSys

	return ret
}

// GetMajorVersion is analogous to the C function atk_get_major_version.
//
// since 2.8
func GetMajorVersion() uint {
	retSys := atk.Fn_atk_get_major_version()
	ret := retSys

	return ret
}

// GetMicroVersion is analogous to the C function atk_get_micro_version.
//
// since 2.8
func GetMicroVersion() uint {
	retSys := atk.Fn_atk_get_micro_version()
	ret := retSys

	return ret
}

// GetMinorVersion is analogous to the C function atk_get_minor_version.
//
// since 2.8
func GetMinorVersion() uint {
	retSys := atk.Fn_atk_get_minor_version()
	ret := retSys

	return ret
}

// GetRoot is analogous to the C function atk_get_root.
func GetRoot() *Object {
	retSys := atk.Fn_atk_get_root()
	ret := ObjectNewFromC(retSys)

	return ret
}

// GetToolkitName is analogous to the C function atk_get_toolkit_name.
func GetToolkitName() string {
	retSys := atk.Fn_atk_get_toolkit_name()
	ret := retSys

	return ret
}

// GetToolkitVersion is analogous to the C function atk_get_toolkit_version.
func GetToolkitVersion() string {
	retSys := atk.Fn_atk_get_toolkit_version()
	ret := retSys

	return ret
}

// GetVersion is analogous to the C function atk_get_version.
//
// since 1.20
func GetVersion() string {
	retSys := atk.Fn_atk_get_version()
	ret := retSys

	return ret
}

// RemoveFocusTracker is analogous to the C function atk_remove_focus_tracker.
func RemoveFocusTracker(trackerId uint) {
	sys_trackerId := trackerId
	atk.Fn_atk_remove_focus_tracker(sys_trackerId)
}

// RemoveGlobalEventListener is analogous to the C function atk_remove_global_event_listener.
func RemoveGlobalEventListener(listenerId uint) {
	sys_listenerId := listenerId
	atk.Fn_atk_remove_global_event_listener(sys_listenerId)
}

// RemoveKeyEventListener is analogous to the C function atk_remove_key_event_listener.
func RemoveKeyEventListener(listenerId uint) {
	sys_listenerId := listenerId
	atk.Fn_atk_remove_key_event_listener(sys_listenerId)
}

// UNSUPPORTED : atk_text_free_ranges : parameter 'ranges' is array parameter without length parameter

// ActionIface is a representation of the C record AtkActionIface.
type ActionIface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C AtkActionIface that represents the ActionIface.
func (recv *ActionIface) ToC() unsafe.Pointer {
	return recv.native
}

// ActionIfaceNewFromC creates a new ActionIface from a pointer to the C AtkActionIface that represents the ActionIface.
func ActionIfaceNewFromC(native unsafe.Pointer) *ActionIface {
	return &ActionIface{native: native}
}

// Attribute is a representation of the C record AtkAttribute.
type Attribute struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C AtkAttribute that represents the Attribute.
func (recv *Attribute) ToC() unsafe.Pointer {
	return recv.native
}

// AttributeNewFromC creates a new Attribute from a pointer to the C AtkAttribute that represents the Attribute.
func AttributeNewFromC(native unsafe.Pointer) *Attribute {
	return &Attribute{native: native}
}

// ComponentIface is a representation of the C record AtkComponentIface.
type ComponentIface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C AtkComponentIface that represents the ComponentIface.
func (recv *ComponentIface) ToC() unsafe.Pointer {
	return recv.native
}

// ComponentIfaceNewFromC creates a new ComponentIface from a pointer to the C AtkComponentIface that represents the ComponentIface.
func ComponentIfaceNewFromC(native unsafe.Pointer) *ComponentIface {
	return &ComponentIface{native: native}
}

// DocumentIface is a representation of the C record AtkDocumentIface.
type DocumentIface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C AtkDocumentIface that represents the DocumentIface.
func (recv *DocumentIface) ToC() unsafe.Pointer {
	return recv.native
}

// DocumentIfaceNewFromC creates a new DocumentIface from a pointer to the C AtkDocumentIface that represents the DocumentIface.
func DocumentIfaceNewFromC(native unsafe.Pointer) *DocumentIface {
	return &DocumentIface{native: native}
}

// EditableTextIface is a representation of the C record AtkEditableTextIface.
type EditableTextIface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C AtkEditableTextIface that represents the EditableTextIface.
func (recv *EditableTextIface) ToC() unsafe.Pointer {
	return recv.native
}

// EditableTextIfaceNewFromC creates a new EditableTextIface from a pointer to the C AtkEditableTextIface that represents the EditableTextIface.
func EditableTextIfaceNewFromC(native unsafe.Pointer) *EditableTextIface {
	return &EditableTextIface{native: native}
}

// GObjectAccessibleClass is a representation of the C record AtkGObjectAccessibleClass.
type GObjectAccessibleClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C AtkGObjectAccessibleClass that represents the GObjectAccessibleClass.
func (recv *GObjectAccessibleClass) ToC() unsafe.Pointer {
	return recv.native
}

// GObjectAccessibleClassNewFromC creates a new GObjectAccessibleClass from a pointer to the C AtkGObjectAccessibleClass that represents the GObjectAccessibleClass.
func GObjectAccessibleClassNewFromC(native unsafe.Pointer) *GObjectAccessibleClass {
	return &GObjectAccessibleClass{native: native}
}

// HyperlinkClass is a representation of the C record AtkHyperlinkClass.
type HyperlinkClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C AtkHyperlinkClass that represents the HyperlinkClass.
func (recv *HyperlinkClass) ToC() unsafe.Pointer {
	return recv.native
}

// HyperlinkClassNewFromC creates a new HyperlinkClass from a pointer to the C AtkHyperlinkClass that represents the HyperlinkClass.
func HyperlinkClassNewFromC(native unsafe.Pointer) *HyperlinkClass {
	return &HyperlinkClass{native: native}
}

// HyperlinkImplIface is a representation of the C record AtkHyperlinkImplIface.
type HyperlinkImplIface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C AtkHyperlinkImplIface that represents the HyperlinkImplIface.
func (recv *HyperlinkImplIface) ToC() unsafe.Pointer {
	return recv.native
}

// HyperlinkImplIfaceNewFromC creates a new HyperlinkImplIface from a pointer to the C AtkHyperlinkImplIface that represents the HyperlinkImplIface.
func HyperlinkImplIfaceNewFromC(native unsafe.Pointer) *HyperlinkImplIface {
	return &HyperlinkImplIface{native: native}
}

// HypertextIface is a representation of the C record AtkHypertextIface.
type HypertextIface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C AtkHypertextIface that represents the HypertextIface.
func (recv *HypertextIface) ToC() unsafe.Pointer {
	return recv.native
}

// HypertextIfaceNewFromC creates a new HypertextIface from a pointer to the C AtkHypertextIface that represents the HypertextIface.
func HypertextIfaceNewFromC(native unsafe.Pointer) *HypertextIface {
	return &HypertextIface{native: native}
}

// ImageIface is a representation of the C record AtkImageIface.
type ImageIface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C AtkImageIface that represents the ImageIface.
func (recv *ImageIface) ToC() unsafe.Pointer {
	return recv.native
}

// ImageIfaceNewFromC creates a new ImageIface from a pointer to the C AtkImageIface that represents the ImageIface.
func ImageIfaceNewFromC(native unsafe.Pointer) *ImageIface {
	return &ImageIface{native: native}
}

// Implementor is a representation of the C record AtkImplementor.
type Implementor struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C AtkImplementor that represents the Implementor.
func (recv *Implementor) ToC() unsafe.Pointer {
	return recv.native
}

// ImplementorNewFromC creates a new Implementor from a pointer to the C AtkImplementor that represents the Implementor.
func ImplementorNewFromC(native unsafe.Pointer) *Implementor {
	return &Implementor{native: native}
}

// KeyEventStruct is a representation of the C record AtkKeyEventStruct.
type KeyEventStruct struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C AtkKeyEventStruct that represents the KeyEventStruct.
func (recv *KeyEventStruct) ToC() unsafe.Pointer {
	return recv.native
}

// KeyEventStructNewFromC creates a new KeyEventStruct from a pointer to the C AtkKeyEventStruct that represents the KeyEventStruct.
func KeyEventStructNewFromC(native unsafe.Pointer) *KeyEventStruct {
	return &KeyEventStruct{native: native}
}

// MiscClass is a representation of the C record AtkMiscClass.
type MiscClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C AtkMiscClass that represents the MiscClass.
func (recv *MiscClass) ToC() unsafe.Pointer {
	return recv.native
}

// MiscClassNewFromC creates a new MiscClass from a pointer to the C AtkMiscClass that represents the MiscClass.
func MiscClassNewFromC(native unsafe.Pointer) *MiscClass {
	return &MiscClass{native: native}
}

// NoOpObjectClass is a representation of the C record AtkNoOpObjectClass.
type NoOpObjectClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C AtkNoOpObjectClass that represents the NoOpObjectClass.
func (recv *NoOpObjectClass) ToC() unsafe.Pointer {
	return recv.native
}

// NoOpObjectClassNewFromC creates a new NoOpObjectClass from a pointer to the C AtkNoOpObjectClass that represents the NoOpObjectClass.
func NoOpObjectClassNewFromC(native unsafe.Pointer) *NoOpObjectClass {
	return &NoOpObjectClass{native: native}
}

// NoOpObjectFactoryClass is a representation of the C record AtkNoOpObjectFactoryClass.
type NoOpObjectFactoryClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C AtkNoOpObjectFactoryClass that represents the NoOpObjectFactoryClass.
func (recv *NoOpObjectFactoryClass) ToC() unsafe.Pointer {
	return recv.native
}

// NoOpObjectFactoryClassNewFromC creates a new NoOpObjectFactoryClass from a pointer to the C AtkNoOpObjectFactoryClass that represents the NoOpObjectFactoryClass.
func NoOpObjectFactoryClassNewFromC(native unsafe.Pointer) *NoOpObjectFactoryClass {
	return &NoOpObjectFactoryClass{native: native}
}

// ObjectClass is a representation of the C record AtkObjectClass.
type ObjectClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C AtkObjectClass that represents the ObjectClass.
func (recv *ObjectClass) ToC() unsafe.Pointer {
	return recv.native
}

// ObjectClassNewFromC creates a new ObjectClass from a pointer to the C AtkObjectClass that represents the ObjectClass.
func ObjectClassNewFromC(native unsafe.Pointer) *ObjectClass {
	return &ObjectClass{native: native}
}

// ObjectFactoryClass is a representation of the C record AtkObjectFactoryClass.
type ObjectFactoryClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C AtkObjectFactoryClass that represents the ObjectFactoryClass.
func (recv *ObjectFactoryClass) ToC() unsafe.Pointer {
	return recv.native
}

// ObjectFactoryClassNewFromC creates a new ObjectFactoryClass from a pointer to the C AtkObjectFactoryClass that represents the ObjectFactoryClass.
func ObjectFactoryClassNewFromC(native unsafe.Pointer) *ObjectFactoryClass {
	return &ObjectFactoryClass{native: native}
}

// PlugClass is a representation of the C record AtkPlugClass.
type PlugClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C AtkPlugClass that represents the PlugClass.
func (recv *PlugClass) ToC() unsafe.Pointer {
	return recv.native
}

// PlugClassNewFromC creates a new PlugClass from a pointer to the C AtkPlugClass that represents the PlugClass.
func PlugClassNewFromC(native unsafe.Pointer) *PlugClass {
	return &PlugClass{native: native}
}

// PropertyValues is a representation of the C record AtkPropertyValues.
type PropertyValues struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C AtkPropertyValues that represents the PropertyValues.
func (recv *PropertyValues) ToC() unsafe.Pointer {
	return recv.native
}

// PropertyValuesNewFromC creates a new PropertyValues from a pointer to the C AtkPropertyValues that represents the PropertyValues.
func PropertyValuesNewFromC(native unsafe.Pointer) *PropertyValues {
	return &PropertyValues{native: native}
}

// Range is a representation of the C record AtkRange.
type Range struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C AtkRange that represents the Range.
func (recv *Range) ToC() unsafe.Pointer {
	return recv.native
}

// RangeNewFromC creates a new Range from a pointer to the C AtkRange that represents the Range.
func RangeNewFromC(native unsafe.Pointer) *Range {
	return &Range{native: native}
}

// Rectangle is a representation of the C record AtkRectangle.
type Rectangle struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C AtkRectangle that represents the Rectangle.
func (recv *Rectangle) ToC() unsafe.Pointer {
	return recv.native
}

// RectangleNewFromC creates a new Rectangle from a pointer to the C AtkRectangle that represents the Rectangle.
func RectangleNewFromC(native unsafe.Pointer) *Rectangle {
	return &Rectangle{native: native}
}

// RegistryClass is a representation of the C record AtkRegistryClass.
type RegistryClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C AtkRegistryClass that represents the RegistryClass.
func (recv *RegistryClass) ToC() unsafe.Pointer {
	return recv.native
}

// RegistryClassNewFromC creates a new RegistryClass from a pointer to the C AtkRegistryClass that represents the RegistryClass.
func RegistryClassNewFromC(native unsafe.Pointer) *RegistryClass {
	return &RegistryClass{native: native}
}

// RelationClass is a representation of the C record AtkRelationClass.
type RelationClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C AtkRelationClass that represents the RelationClass.
func (recv *RelationClass) ToC() unsafe.Pointer {
	return recv.native
}

// RelationClassNewFromC creates a new RelationClass from a pointer to the C AtkRelationClass that represents the RelationClass.
func RelationClassNewFromC(native unsafe.Pointer) *RelationClass {
	return &RelationClass{native: native}
}

// RelationSetClass is a representation of the C record AtkRelationSetClass.
type RelationSetClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C AtkRelationSetClass that represents the RelationSetClass.
func (recv *RelationSetClass) ToC() unsafe.Pointer {
	return recv.native
}

// RelationSetClassNewFromC creates a new RelationSetClass from a pointer to the C AtkRelationSetClass that represents the RelationSetClass.
func RelationSetClassNewFromC(native unsafe.Pointer) *RelationSetClass {
	return &RelationSetClass{native: native}
}

// SelectionIface is a representation of the C record AtkSelectionIface.
type SelectionIface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C AtkSelectionIface that represents the SelectionIface.
func (recv *SelectionIface) ToC() unsafe.Pointer {
	return recv.native
}

// SelectionIfaceNewFromC creates a new SelectionIface from a pointer to the C AtkSelectionIface that represents the SelectionIface.
func SelectionIfaceNewFromC(native unsafe.Pointer) *SelectionIface {
	return &SelectionIface{native: native}
}

// SocketClass is a representation of the C record AtkSocketClass.
type SocketClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C AtkSocketClass that represents the SocketClass.
func (recv *SocketClass) ToC() unsafe.Pointer {
	return recv.native
}

// SocketClassNewFromC creates a new SocketClass from a pointer to the C AtkSocketClass that represents the SocketClass.
func SocketClassNewFromC(native unsafe.Pointer) *SocketClass {
	return &SocketClass{native: native}
}

// StateSetClass is a representation of the C record AtkStateSetClass.
type StateSetClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C AtkStateSetClass that represents the StateSetClass.
func (recv *StateSetClass) ToC() unsafe.Pointer {
	return recv.native
}

// StateSetClassNewFromC creates a new StateSetClass from a pointer to the C AtkStateSetClass that represents the StateSetClass.
func StateSetClassNewFromC(native unsafe.Pointer) *StateSetClass {
	return &StateSetClass{native: native}
}

// StreamableContentIface is a representation of the C record AtkStreamableContentIface.
type StreamableContentIface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C AtkStreamableContentIface that represents the StreamableContentIface.
func (recv *StreamableContentIface) ToC() unsafe.Pointer {
	return recv.native
}

// StreamableContentIfaceNewFromC creates a new StreamableContentIface from a pointer to the C AtkStreamableContentIface that represents the StreamableContentIface.
func StreamableContentIfaceNewFromC(native unsafe.Pointer) *StreamableContentIface {
	return &StreamableContentIface{native: native}
}

// TableIface is a representation of the C record AtkTableIface.
type TableIface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C AtkTableIface that represents the TableIface.
func (recv *TableIface) ToC() unsafe.Pointer {
	return recv.native
}

// TableIfaceNewFromC creates a new TableIface from a pointer to the C AtkTableIface that represents the TableIface.
func TableIfaceNewFromC(native unsafe.Pointer) *TableIface {
	return &TableIface{native: native}
}

// TextIface is a representation of the C record AtkTextIface.
type TextIface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C AtkTextIface that represents the TextIface.
func (recv *TextIface) ToC() unsafe.Pointer {
	return recv.native
}

// TextIfaceNewFromC creates a new TextIface from a pointer to the C AtkTextIface that represents the TextIface.
func TextIfaceNewFromC(native unsafe.Pointer) *TextIface {
	return &TextIface{native: native}
}

// TextRange is a representation of the C record AtkTextRange.
type TextRange struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C AtkTextRange that represents the TextRange.
func (recv *TextRange) ToC() unsafe.Pointer {
	return recv.native
}

// TextRangeNewFromC creates a new TextRange from a pointer to the C AtkTextRange that represents the TextRange.
func TextRangeNewFromC(native unsafe.Pointer) *TextRange {
	return &TextRange{native: native}
}

// TextRectangle is a representation of the C record AtkTextRectangle.
type TextRectangle struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C AtkTextRectangle that represents the TextRectangle.
func (recv *TextRectangle) ToC() unsafe.Pointer {
	return recv.native
}

// TextRectangleNewFromC creates a new TextRectangle from a pointer to the C AtkTextRectangle that represents the TextRectangle.
func TextRectangleNewFromC(native unsafe.Pointer) *TextRectangle {
	return &TextRectangle{native: native}
}

// UtilClass is a representation of the C record AtkUtilClass.
type UtilClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C AtkUtilClass that represents the UtilClass.
func (recv *UtilClass) ToC() unsafe.Pointer {
	return recv.native
}

// UtilClassNewFromC creates a new UtilClass from a pointer to the C AtkUtilClass that represents the UtilClass.
func UtilClassNewFromC(native unsafe.Pointer) *UtilClass {
	return &UtilClass{native: native}
}

// ValueIface is a representation of the C record AtkValueIface.
type ValueIface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C AtkValueIface that represents the ValueIface.
func (recv *ValueIface) ToC() unsafe.Pointer {
	return recv.native
}

// ValueIfaceNewFromC creates a new ValueIface from a pointer to the C AtkValueIface that represents the ValueIface.
func ValueIfaceNewFromC(native unsafe.Pointer) *ValueIface {
	return &ValueIface{native: native}
}

// WindowIface is a representation of the C record AtkWindowIface.
type WindowIface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C AtkWindowIface that represents the WindowIface.
func (recv *WindowIface) ToC() unsafe.Pointer {
	return recv.native
}

// WindowIfaceNewFromC creates a new WindowIface from a pointer to the C AtkWindowIface that represents the WindowIface.
func WindowIfaceNewFromC(native unsafe.Pointer) *WindowIface {
	return &WindowIface{native: native}
}

// GObjectAccessible is a representation of the C record AtkGObjectAccessible.
type GObjectAccessible struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C AtkGObjectAccessible that represents the GObjectAccessible.
func (recv *GObjectAccessible) ToC() unsafe.Pointer {
	return recv.native
}

// GObjectAccessibleNewFromC creates a new GObjectAccessible from a pointer to the C AtkGObjectAccessible that represents the GObjectAccessible.
func GObjectAccessibleNewFromC(native unsafe.Pointer) *GObjectAccessible {
	return &GObjectAccessible{native: native}
}

// Hyperlink is a representation of the C record AtkHyperlink.
type Hyperlink struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C AtkHyperlink that represents the Hyperlink.
func (recv *Hyperlink) ToC() unsafe.Pointer {
	return recv.native
}

// HyperlinkNewFromC creates a new Hyperlink from a pointer to the C AtkHyperlink that represents the Hyperlink.
func HyperlinkNewFromC(native unsafe.Pointer) *Hyperlink {
	return &Hyperlink{native: native}
}

// Misc is a representation of the C record AtkMisc.
type Misc struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C AtkMisc that represents the Misc.
func (recv *Misc) ToC() unsafe.Pointer {
	return recv.native
}

// MiscNewFromC creates a new Misc from a pointer to the C AtkMisc that represents the Misc.
func MiscNewFromC(native unsafe.Pointer) *Misc {
	return &Misc{native: native}
}

// NoOpObject is a representation of the C record AtkNoOpObject.
type NoOpObject struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C AtkNoOpObject that represents the NoOpObject.
func (recv *NoOpObject) ToC() unsafe.Pointer {
	return recv.native
}

// NoOpObjectNewFromC creates a new NoOpObject from a pointer to the C AtkNoOpObject that represents the NoOpObject.
func NoOpObjectNewFromC(native unsafe.Pointer) *NoOpObject {
	return &NoOpObject{native: native}
}

// NoOpObjectFactory is a representation of the C record AtkNoOpObjectFactory.
type NoOpObjectFactory struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C AtkNoOpObjectFactory that represents the NoOpObjectFactory.
func (recv *NoOpObjectFactory) ToC() unsafe.Pointer {
	return recv.native
}

// NoOpObjectFactoryNewFromC creates a new NoOpObjectFactory from a pointer to the C AtkNoOpObjectFactory that represents the NoOpObjectFactory.
func NoOpObjectFactoryNewFromC(native unsafe.Pointer) *NoOpObjectFactory {
	return &NoOpObjectFactory{native: native}
}

// Object is a representation of the C record AtkObject.
type Object struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C AtkObject that represents the Object.
func (recv *Object) ToC() unsafe.Pointer {
	return recv.native
}

// ObjectNewFromC creates a new Object from a pointer to the C AtkObject that represents the Object.
func ObjectNewFromC(native unsafe.Pointer) *Object {
	return &Object{native: native}
}

// ObjectFactory is a representation of the C record AtkObjectFactory.
type ObjectFactory struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C AtkObjectFactory that represents the ObjectFactory.
func (recv *ObjectFactory) ToC() unsafe.Pointer {
	return recv.native
}

// ObjectFactoryNewFromC creates a new ObjectFactory from a pointer to the C AtkObjectFactory that represents the ObjectFactory.
func ObjectFactoryNewFromC(native unsafe.Pointer) *ObjectFactory {
	return &ObjectFactory{native: native}
}

// Plug is a representation of the C record AtkPlug.
type Plug struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C AtkPlug that represents the Plug.
func (recv *Plug) ToC() unsafe.Pointer {
	return recv.native
}

// PlugNewFromC creates a new Plug from a pointer to the C AtkPlug that represents the Plug.
func PlugNewFromC(native unsafe.Pointer) *Plug {
	return &Plug{native: native}
}

// Registry is a representation of the C record AtkRegistry.
type Registry struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C AtkRegistry that represents the Registry.
func (recv *Registry) ToC() unsafe.Pointer {
	return recv.native
}

// RegistryNewFromC creates a new Registry from a pointer to the C AtkRegistry that represents the Registry.
func RegistryNewFromC(native unsafe.Pointer) *Registry {
	return &Registry{native: native}
}

// Relation is a representation of the C record AtkRelation.
type Relation struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C AtkRelation that represents the Relation.
func (recv *Relation) ToC() unsafe.Pointer {
	return recv.native
}

// RelationNewFromC creates a new Relation from a pointer to the C AtkRelation that represents the Relation.
func RelationNewFromC(native unsafe.Pointer) *Relation {
	return &Relation{native: native}
}

// RelationSet is a representation of the C record AtkRelationSet.
type RelationSet struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C AtkRelationSet that represents the RelationSet.
func (recv *RelationSet) ToC() unsafe.Pointer {
	return recv.native
}

// RelationSetNewFromC creates a new RelationSet from a pointer to the C AtkRelationSet that represents the RelationSet.
func RelationSetNewFromC(native unsafe.Pointer) *RelationSet {
	return &RelationSet{native: native}
}

// Socket is a representation of the C record AtkSocket.
type Socket struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C AtkSocket that represents the Socket.
func (recv *Socket) ToC() unsafe.Pointer {
	return recv.native
}

// SocketNewFromC creates a new Socket from a pointer to the C AtkSocket that represents the Socket.
func SocketNewFromC(native unsafe.Pointer) *Socket {
	return &Socket{native: native}
}

// StateSet is a representation of the C record AtkStateSet.
type StateSet struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C AtkStateSet that represents the StateSet.
func (recv *StateSet) ToC() unsafe.Pointer {
	return recv.native
}

// StateSetNewFromC creates a new StateSet from a pointer to the C AtkStateSet that represents the StateSet.
func StateSetNewFromC(native unsafe.Pointer) *StateSet {
	return &StateSet{native: native}
}

// Util is a representation of the C record AtkUtil.
type Util struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C AtkUtil that represents the Util.
func (recv *Util) ToC() unsafe.Pointer {
	return recv.native
}

// UtilNewFromC creates a new Util from a pointer to the C AtkUtil that represents the Util.
func UtilNewFromC(native unsafe.Pointer) *Util {
	return &Util{native: native}
}

// Action is a representation of the C interface AtkAction.
type Action struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C AtkAction that represents the Action.
func (recv *Action) ToC() unsafe.Pointer {
	return recv.native
}

// ActionNewFromC creates a new Action from a pointer to the C AtkAction that represents the Action.
func ActionNewFromC(native unsafe.Pointer) *Action {
	return &Action{native: native}
}

// Component is a representation of the C interface AtkComponent.
type Component struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C AtkComponent that represents the Component.
func (recv *Component) ToC() unsafe.Pointer {
	return recv.native
}

// ComponentNewFromC creates a new Component from a pointer to the C AtkComponent that represents the Component.
func ComponentNewFromC(native unsafe.Pointer) *Component {
	return &Component{native: native}
}

// Document is a representation of the C interface AtkDocument.
type Document struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C AtkDocument that represents the Document.
func (recv *Document) ToC() unsafe.Pointer {
	return recv.native
}

// DocumentNewFromC creates a new Document from a pointer to the C AtkDocument that represents the Document.
func DocumentNewFromC(native unsafe.Pointer) *Document {
	return &Document{native: native}
}

// EditableText is a representation of the C interface AtkEditableText.
type EditableText struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C AtkEditableText that represents the EditableText.
func (recv *EditableText) ToC() unsafe.Pointer {
	return recv.native
}

// EditableTextNewFromC creates a new EditableText from a pointer to the C AtkEditableText that represents the EditableText.
func EditableTextNewFromC(native unsafe.Pointer) *EditableText {
	return &EditableText{native: native}
}

// HyperlinkImpl is a representation of the C interface AtkHyperlinkImpl.
type HyperlinkImpl struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C AtkHyperlinkImpl that represents the HyperlinkImpl.
func (recv *HyperlinkImpl) ToC() unsafe.Pointer {
	return recv.native
}

// HyperlinkImplNewFromC creates a new HyperlinkImpl from a pointer to the C AtkHyperlinkImpl that represents the HyperlinkImpl.
func HyperlinkImplNewFromC(native unsafe.Pointer) *HyperlinkImpl {
	return &HyperlinkImpl{native: native}
}

// Hypertext is a representation of the C interface AtkHypertext.
type Hypertext struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C AtkHypertext that represents the Hypertext.
func (recv *Hypertext) ToC() unsafe.Pointer {
	return recv.native
}

// HypertextNewFromC creates a new Hypertext from a pointer to the C AtkHypertext that represents the Hypertext.
func HypertextNewFromC(native unsafe.Pointer) *Hypertext {
	return &Hypertext{native: native}
}

// Image is a representation of the C interface AtkImage.
type Image struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C AtkImage that represents the Image.
func (recv *Image) ToC() unsafe.Pointer {
	return recv.native
}

// ImageNewFromC creates a new Image from a pointer to the C AtkImage that represents the Image.
func ImageNewFromC(native unsafe.Pointer) *Image {
	return &Image{native: native}
}

// ImplementorIface is a representation of the C interface AtkImplementorIface.
type ImplementorIface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C AtkImplementorIface that represents the ImplementorIface.
func (recv *ImplementorIface) ToC() unsafe.Pointer {
	return recv.native
}

// ImplementorIfaceNewFromC creates a new ImplementorIface from a pointer to the C AtkImplementorIface that represents the ImplementorIface.
func ImplementorIfaceNewFromC(native unsafe.Pointer) *ImplementorIface {
	return &ImplementorIface{native: native}
}

// Selection is a representation of the C interface AtkSelection.
type Selection struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C AtkSelection that represents the Selection.
func (recv *Selection) ToC() unsafe.Pointer {
	return recv.native
}

// SelectionNewFromC creates a new Selection from a pointer to the C AtkSelection that represents the Selection.
func SelectionNewFromC(native unsafe.Pointer) *Selection {
	return &Selection{native: native}
}

// StreamableContent is a representation of the C interface AtkStreamableContent.
type StreamableContent struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C AtkStreamableContent that represents the StreamableContent.
func (recv *StreamableContent) ToC() unsafe.Pointer {
	return recv.native
}

// StreamableContentNewFromC creates a new StreamableContent from a pointer to the C AtkStreamableContent that represents the StreamableContent.
func StreamableContentNewFromC(native unsafe.Pointer) *StreamableContent {
	return &StreamableContent{native: native}
}

// Table is a representation of the C interface AtkTable.
type Table struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C AtkTable that represents the Table.
func (recv *Table) ToC() unsafe.Pointer {
	return recv.native
}

// TableNewFromC creates a new Table from a pointer to the C AtkTable that represents the Table.
func TableNewFromC(native unsafe.Pointer) *Table {
	return &Table{native: native}
}

// TableCell is a representation of the C interface AtkTableCell.
type TableCell struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C AtkTableCell that represents the TableCell.
func (recv *TableCell) ToC() unsafe.Pointer {
	return recv.native
}

// TableCellNewFromC creates a new TableCell from a pointer to the C AtkTableCell that represents the TableCell.
func TableCellNewFromC(native unsafe.Pointer) *TableCell {
	return &TableCell{native: native}
}

// Text is a representation of the C interface AtkText.
type Text struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C AtkText that represents the Text.
func (recv *Text) ToC() unsafe.Pointer {
	return recv.native
}

// TextNewFromC creates a new Text from a pointer to the C AtkText that represents the Text.
func TextNewFromC(native unsafe.Pointer) *Text {
	return &Text{native: native}
}

// Value is a representation of the C interface AtkValue.
type Value struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C AtkValue that represents the Value.
func (recv *Value) ToC() unsafe.Pointer {
	return recv.native
}

// ValueNewFromC creates a new Value from a pointer to the C AtkValue that represents the Value.
func ValueNewFromC(native unsafe.Pointer) *Value {
	return &Value{native: native}
}

// Window is a representation of the C interface AtkWindow.
type Window struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C AtkWindow that represents the Window.
func (recv *Window) ToC() unsafe.Pointer {
	return recv.native
}

// WindowNewFromC creates a new Window from a pointer to the C AtkWindow that represents the Window.
func WindowNewFromC(native unsafe.Pointer) *Window {
	return &Window{native: native}
}
