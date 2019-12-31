// Code generated - DO NOT EDIT.
// +build atk_2.34

package atk

import (
	glib "github.com/pekim/gobbi/lib/glib"
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

// VERSION_MIN_REQUIRED is a representation of the C constant ATK_VERSION_MIN_REQUIRED.
//
// since 2.14
const VERSION_MIN_REQUIRED = 2

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

// ScrollType is a representation of the C enumeration AtkScrollType.
type ScrollType int

// ScrollType_top_left is a representation of the C enumeration member ATK_SCROLL_TOP_LEFT.
const ScrollType_top_left = ScrollType(0)

// ScrollType_bottom_right is a representation of the C enumeration member ATK_SCROLL_BOTTOM_RIGHT.
const ScrollType_bottom_right = ScrollType(1)

// ScrollType_top_edge is a representation of the C enumeration member ATK_SCROLL_TOP_EDGE.
const ScrollType_top_edge = ScrollType(2)

// ScrollType_bottom_edge is a representation of the C enumeration member ATK_SCROLL_BOTTOM_EDGE.
const ScrollType_bottom_edge = ScrollType(3)

// ScrollType_left_edge is a representation of the C enumeration member ATK_SCROLL_LEFT_EDGE.
const ScrollType_left_edge = ScrollType(4)

// ScrollType_right_edge is a representation of the C enumeration member ATK_SCROLL_RIGHT_EDGE.
const ScrollType_right_edge = ScrollType(5)

// ScrollType_anywhere is a representation of the C enumeration member ATK_SCROLL_ANYWHERE.
const ScrollType_anywhere = ScrollType(6)

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

// ActionIface is a representation of the C record AtkActionIface.
type ActionIface struct {
	native unsafe.Pointer
}

// Attribute is a representation of the C record AtkAttribute.
type Attribute struct {
	native unsafe.Pointer
}

// ComponentIface is a representation of the C record AtkComponentIface.
type ComponentIface struct {
	native unsafe.Pointer
}

// DocumentIface is a representation of the C record AtkDocumentIface.
type DocumentIface struct {
	native unsafe.Pointer
}

// EditableTextIface is a representation of the C record AtkEditableTextIface.
type EditableTextIface struct {
	native unsafe.Pointer
}

// GObjectAccessibleClass is a representation of the C record AtkGObjectAccessibleClass.
type GObjectAccessibleClass struct {
	native unsafe.Pointer
}

// HyperlinkClass is a representation of the C record AtkHyperlinkClass.
type HyperlinkClass struct {
	native unsafe.Pointer
}

// HyperlinkImplIface is a representation of the C record AtkHyperlinkImplIface.
type HyperlinkImplIface struct {
	native unsafe.Pointer
}

// HypertextIface is a representation of the C record AtkHypertextIface.
type HypertextIface struct {
	native unsafe.Pointer
}

// ImageIface is a representation of the C record AtkImageIface.
type ImageIface struct {
	native unsafe.Pointer
}

// Implementor is a representation of the C record AtkImplementor.
type Implementor struct {
	native unsafe.Pointer
}

// KeyEventStruct is a representation of the C record AtkKeyEventStruct.
type KeyEventStruct struct {
	native unsafe.Pointer
}

// MiscClass is a representation of the C record AtkMiscClass.
type MiscClass struct {
	native unsafe.Pointer
}

// NoOpObjectClass is a representation of the C record AtkNoOpObjectClass.
type NoOpObjectClass struct {
	native unsafe.Pointer
}

// NoOpObjectFactoryClass is a representation of the C record AtkNoOpObjectFactoryClass.
type NoOpObjectFactoryClass struct {
	native unsafe.Pointer
}

// ObjectClass is a representation of the C record AtkObjectClass.
type ObjectClass struct {
	native unsafe.Pointer
}

// ObjectFactoryClass is a representation of the C record AtkObjectFactoryClass.
type ObjectFactoryClass struct {
	native unsafe.Pointer
}

// PlugClass is a representation of the C record AtkPlugClass.
type PlugClass struct {
	native unsafe.Pointer
}

// PropertyValues is a representation of the C record AtkPropertyValues.
type PropertyValues struct {
	native unsafe.Pointer
}

// Range is a representation of the C record AtkRange.
type Range struct {
	native unsafe.Pointer
}

// Rectangle is a representation of the C record AtkRectangle.
type Rectangle struct {
	native unsafe.Pointer
}

// RegistryClass is a representation of the C record AtkRegistryClass.
type RegistryClass struct {
	native unsafe.Pointer
}

// RelationClass is a representation of the C record AtkRelationClass.
type RelationClass struct {
	native unsafe.Pointer
}

// RelationSetClass is a representation of the C record AtkRelationSetClass.
type RelationSetClass struct {
	native unsafe.Pointer
}

// SelectionIface is a representation of the C record AtkSelectionIface.
type SelectionIface struct {
	native unsafe.Pointer
}

// SocketClass is a representation of the C record AtkSocketClass.
type SocketClass struct {
	native unsafe.Pointer
}

// StateSetClass is a representation of the C record AtkStateSetClass.
type StateSetClass struct {
	native unsafe.Pointer
}

// StreamableContentIface is a representation of the C record AtkStreamableContentIface.
type StreamableContentIface struct {
	native unsafe.Pointer
}

// TableCellIface is a representation of the C record AtkTableCellIface.
//
// since 2.12
type TableCellIface struct {
	native unsafe.Pointer
}

// TableIface is a representation of the C record AtkTableIface.
type TableIface struct {
	native unsafe.Pointer
}

// TextIface is a representation of the C record AtkTextIface.
type TextIface struct {
	native unsafe.Pointer
}

// TextRange is a representation of the C record AtkTextRange.
type TextRange struct {
	native unsafe.Pointer
}

// TextRectangle is a representation of the C record AtkTextRectangle.
type TextRectangle struct {
	native unsafe.Pointer
}

// UtilClass is a representation of the C record AtkUtilClass.
type UtilClass struct {
	native unsafe.Pointer
}

// ValueIface is a representation of the C record AtkValueIface.
type ValueIface struct {
	native unsafe.Pointer
}

// WindowIface is a representation of the C record AtkWindowIface.
type WindowIface struct {
	native unsafe.Pointer
}
