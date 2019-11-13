// Code generated - DO NOT EDIT.

package atk

// Coordtype is a representation of the C type CoordType.
type Coordtype int

const (
	// screen
	AtkXyScreen Coordtype = 0
	// window
	AtkXyWindow Coordtype = 1
)

// Keyeventtype is a representation of the C type KeyEventType.
type Keyeventtype int

const (
	// press
	AtkKeyEventPress Keyeventtype = 0
	// release
	AtkKeyEventRelease Keyeventtype = 1
	// last_defined
	AtkKeyEventLastDefined Keyeventtype = 2
)

// Layer is a representation of the C type Layer.
type Layer int

const (
	// invalid
	AtkLayerInvalid Layer = 0
	// background
	AtkLayerBackground Layer = 1
	// canvas
	AtkLayerCanvas Layer = 2
	// widget
	AtkLayerWidget Layer = 3
	// mdi
	AtkLayerMdi Layer = 4
	// popup
	AtkLayerPopup Layer = 5
	// overlay
	AtkLayerOverlay Layer = 6
	// window
	AtkLayerWindow Layer = 7
)

// Relationtype is a representation of the C type RelationType.
type Relationtype int

const (
	// null
	AtkRelationNull Relationtype = 0
	// controlled_by
	AtkRelationControlledBy Relationtype = 1
	// controller_for
	AtkRelationControllerFor Relationtype = 2
	// label_for
	AtkRelationLabelFor Relationtype = 3
	// labelled_by
	AtkRelationLabelledBy Relationtype = 4
	// member_of
	AtkRelationMemberOf Relationtype = 5
	// node_child_of
	AtkRelationNodeChildOf Relationtype = 6
	// flows_to
	AtkRelationFlowsTo Relationtype = 7
	// flows_from
	AtkRelationFlowsFrom Relationtype = 8
	// subwindow_of
	AtkRelationSubwindowOf Relationtype = 9
	// embeds
	AtkRelationEmbeds Relationtype = 10
	// embedded_by
	AtkRelationEmbeddedBy Relationtype = 11
	// popup_for
	AtkRelationPopupFor Relationtype = 12
	// parent_window_of
	AtkRelationParentWindowOf Relationtype = 13
	// described_by
	AtkRelationDescribedBy Relationtype = 14
	// description_for
	AtkRelationDescriptionFor Relationtype = 15
	// node_parent_of
	AtkRelationNodeParentOf Relationtype = 16
	// details
	AtkRelationDetails Relationtype = 17
	// details_for
	AtkRelationDetailsFor Relationtype = 18
	// error_message
	AtkRelationErrorMessage Relationtype = 19
	// error_for
	AtkRelationErrorFor Relationtype = 20
	// last_defined
	AtkRelationLastDefined Relationtype = 21
)

// Role is a representation of the C type Role.
type Role int

const (
	// invalid
	AtkRoleInvalid Role = 0
	// accelerator_label
	AtkRoleAccelLabel Role = 1
	// alert
	AtkRoleAlert Role = 2
	// animation
	AtkRoleAnimation Role = 3
	// arrow
	AtkRoleArrow Role = 4
	// calendar
	AtkRoleCalendar Role = 5
	// canvas
	AtkRoleCanvas Role = 6
	// check_box
	AtkRoleCheckBox Role = 7
	// check_menu_item
	AtkRoleCheckMenuItem Role = 8
	// color_chooser
	AtkRoleColorChooser Role = 9
	// column_header
	AtkRoleColumnHeader Role = 10
	// combo_box
	AtkRoleComboBox Role = 11
	// date_editor
	AtkRoleDateEditor Role = 12
	// desktop_icon
	AtkRoleDesktopIcon Role = 13
	// desktop_frame
	AtkRoleDesktopFrame Role = 14
	// dial
	AtkRoleDial Role = 15
	// dialog
	AtkRoleDialog Role = 16
	// directory_pane
	AtkRoleDirectoryPane Role = 17
	// drawing_area
	AtkRoleDrawingArea Role = 18
	// file_chooser
	AtkRoleFileChooser Role = 19
	// filler
	AtkRoleFiller Role = 20
	// font_chooser
	AtkRoleFontChooser Role = 21
	// frame
	AtkRoleFrame Role = 22
	// glass_pane
	AtkRoleGlassPane Role = 23
	// html_container
	AtkRoleHtmlContainer Role = 24
	// icon
	AtkRoleIcon Role = 25
	// image
	AtkRoleImage Role = 26
	// internal_frame
	AtkRoleInternalFrame Role = 27
	// label
	AtkRoleLabel Role = 28
	// layered_pane
	AtkRoleLayeredPane Role = 29
	// list
	AtkRoleList Role = 30
	// list_item
	AtkRoleListItem Role = 31
	// menu
	AtkRoleMenu Role = 32
	// menu_bar
	AtkRoleMenuBar Role = 33
	// menu_item
	AtkRoleMenuItem Role = 34
	// option_pane
	AtkRoleOptionPane Role = 35
	// page_tab
	AtkRolePageTab Role = 36
	// page_tab_list
	AtkRolePageTabList Role = 37
	// panel
	AtkRolePanel Role = 38
	// password_text
	AtkRolePasswordText Role = 39
	// popup_menu
	AtkRolePopupMenu Role = 40
	// progress_bar
	AtkRoleProgressBar Role = 41
	// push_button
	AtkRolePushButton Role = 42
	// radio_button
	AtkRoleRadioButton Role = 43
	// radio_menu_item
	AtkRoleRadioMenuItem Role = 44
	// root_pane
	AtkRoleRootPane Role = 45
	// row_header
	AtkRoleRowHeader Role = 46
	// scroll_bar
	AtkRoleScrollBar Role = 47
	// scroll_pane
	AtkRoleScrollPane Role = 48
	// separator
	AtkRoleSeparator Role = 49
	// slider
	AtkRoleSlider Role = 50
	// split_pane
	AtkRoleSplitPane Role = 51
	// spin_button
	AtkRoleSpinButton Role = 52
	// statusbar
	AtkRoleStatusbar Role = 53
	// table
	AtkRoleTable Role = 54
	// table_cell
	AtkRoleTableCell Role = 55
	// table_column_header
	AtkRoleTableColumnHeader Role = 56
	// table_row_header
	AtkRoleTableRowHeader Role = 57
	// tear_off_menu_item
	AtkRoleTearOffMenuItem Role = 58
	// terminal
	AtkRoleTerminal Role = 59
	// text
	AtkRoleText Role = 60
	// toggle_button
	AtkRoleToggleButton Role = 61
	// tool_bar
	AtkRoleToolBar Role = 62
	// tool_tip
	AtkRoleToolTip Role = 63
	// tree
	AtkRoleTree Role = 64
	// tree_table
	AtkRoleTreeTable Role = 65
	// unknown
	AtkRoleUnknown Role = 66
	// viewport
	AtkRoleViewport Role = 67
	// window
	AtkRoleWindow Role = 68
	// header
	AtkRoleHeader Role = 69
	// footer
	AtkRoleFooter Role = 70
	// paragraph
	AtkRoleParagraph Role = 71
	// ruler
	AtkRoleRuler Role = 72
	// application
	AtkRoleApplication Role = 73
	// autocomplete
	AtkRoleAutocomplete Role = 74
	// edit_bar
	AtkRoleEditbar Role = 75
	// embedded
	AtkRoleEmbedded Role = 76
	// entry
	AtkRoleEntry Role = 77
	// chart
	AtkRoleChart Role = 78
	// caption
	AtkRoleCaption Role = 79
	// document_frame
	AtkRoleDocumentFrame Role = 80
	// heading
	AtkRoleHeading Role = 81
	// page
	AtkRolePage Role = 82
	// section
	AtkRoleSection Role = 83
	// redundant_object
	AtkRoleRedundantObject Role = 84
	// form
	AtkRoleForm Role = 85
	// link
	AtkRoleLink Role = 86
	// input_method_window
	AtkRoleInputMethodWindow Role = 87
	// table_row
	AtkRoleTableRow Role = 88
	// tree_item
	AtkRoleTreeItem Role = 89
	// document_spreadsheet
	AtkRoleDocumentSpreadsheet Role = 90
	// document_presentation
	AtkRoleDocumentPresentation Role = 91
	// document_text
	AtkRoleDocumentText Role = 92
	// document_web
	AtkRoleDocumentWeb Role = 93
	// document_email
	AtkRoleDocumentEmail Role = 94
	// comment
	AtkRoleComment Role = 95
	// list_box
	AtkRoleListBox Role = 96
	// grouping
	AtkRoleGrouping Role = 97
	// image_map
	AtkRoleImageMap Role = 98
	// notification
	AtkRoleNotification Role = 99
	// info_bar
	AtkRoleInfoBar Role = 100
	// level_bar
	AtkRoleLevelBar Role = 101
	// title_bar
	AtkRoleTitleBar Role = 102
	// block_quote
	AtkRoleBlockQuote Role = 103
	// audio
	AtkRoleAudio Role = 104
	// video
	AtkRoleVideo Role = 105
	// definition
	AtkRoleDefinition Role = 106
	// article
	AtkRoleArticle Role = 107
	// landmark
	AtkRoleLandmark Role = 108
	// log
	AtkRoleLog Role = 109
	// marquee
	AtkRoleMarquee Role = 110
	// math
	AtkRoleMath Role = 111
	// rating
	AtkRoleRating Role = 112
	// timer
	AtkRoleTimer Role = 113
	// description_list
	AtkRoleDescriptionList Role = 114
	// description_term
	AtkRoleDescriptionTerm Role = 115
	// description_value
	AtkRoleDescriptionValue Role = 116
	// static
	AtkRoleStatic Role = 117
	// math_fraction
	AtkRoleMathFraction Role = 118
	// math_root
	AtkRoleMathRoot Role = 119
	// subscript
	AtkRoleSubscript Role = 120
	// superscript
	AtkRoleSuperscript Role = 121
	// footnote
	AtkRoleFootnote Role = 122
	// last_defined
	AtkRoleLastDefined Role = 123
)

// Statetype is a representation of the C type StateType.
type Statetype int

const (
	// invalid
	AtkStateInvalid Statetype = 0
	// active
	AtkStateActive Statetype = 1
	// armed
	AtkStateArmed Statetype = 2
	// busy
	AtkStateBusy Statetype = 3
	// checked
	AtkStateChecked Statetype = 4
	// defunct
	AtkStateDefunct Statetype = 5
	// editable
	AtkStateEditable Statetype = 6
	// enabled
	AtkStateEnabled Statetype = 7
	// expandable
	AtkStateExpandable Statetype = 8
	// expanded
	AtkStateExpanded Statetype = 9
	// focusable
	AtkStateFocusable Statetype = 10
	// focused
	AtkStateFocused Statetype = 11
	// horizontal
	AtkStateHorizontal Statetype = 12
	// iconified
	AtkStateIconified Statetype = 13
	// modal
	AtkStateModal Statetype = 14
	// multi_line
	AtkStateMultiLine Statetype = 15
	// multiselectable
	AtkStateMultiselectable Statetype = 16
	// opaque
	AtkStateOpaque Statetype = 17
	// pressed
	AtkStatePressed Statetype = 18
	// resizable
	AtkStateResizable Statetype = 19
	// selectable
	AtkStateSelectable Statetype = 20
	// selected
	AtkStateSelected Statetype = 21
	// sensitive
	AtkStateSensitive Statetype = 22
	// showing
	AtkStateShowing Statetype = 23
	// single_line
	AtkStateSingleLine Statetype = 24
	// stale
	AtkStateStale Statetype = 25
	// transient
	AtkStateTransient Statetype = 26
	// vertical
	AtkStateVertical Statetype = 27
	// visible
	AtkStateVisible Statetype = 28
	// manages_descendants
	AtkStateManagesDescendants Statetype = 29
	// indeterminate
	AtkStateIndeterminate Statetype = 30
	// truncated
	AtkStateTruncated Statetype = 31
	// required
	AtkStateRequired Statetype = 32
	// invalid_entry
	AtkStateInvalidEntry Statetype = 33
	// supports_autocompletion
	AtkStateSupportsAutocompletion Statetype = 34
	// selectable_text
	AtkStateSelectableText Statetype = 35
	// default
	AtkStateDefault Statetype = 36
	// animated
	AtkStateAnimated Statetype = 37
	// visited
	AtkStateVisited Statetype = 38
	// checkable
	AtkStateCheckable Statetype = 39
	// has_popup
	AtkStateHasPopup Statetype = 40
	// has_tooltip
	AtkStateHasTooltip Statetype = 41
	// read_only
	AtkStateReadOnly Statetype = 42
	// last_defined
	AtkStateLastDefined Statetype = 43
)

// Textattribute is a representation of the C type TextAttribute.
type Textattribute int

const (
	// invalid
	AtkTextAttrInvalid Textattribute = 0
	// left_margin
	AtkTextAttrLeftMargin Textattribute = 1
	// right_margin
	AtkTextAttrRightMargin Textattribute = 2
	// indent
	AtkTextAttrIndent Textattribute = 3
	// invisible
	AtkTextAttrInvisible Textattribute = 4
	// editable
	AtkTextAttrEditable Textattribute = 5
	// pixels_above_lines
	AtkTextAttrPixelsAboveLines Textattribute = 6
	// pixels_below_lines
	AtkTextAttrPixelsBelowLines Textattribute = 7
	// pixels_inside_wrap
	AtkTextAttrPixelsInsideWrap Textattribute = 8
	// bg_full_height
	AtkTextAttrBgFullHeight Textattribute = 9
	// rise
	AtkTextAttrRise Textattribute = 10
	// underline
	AtkTextAttrUnderline Textattribute = 11
	// strikethrough
	AtkTextAttrStrikethrough Textattribute = 12
	// size
	AtkTextAttrSize Textattribute = 13
	// scale
	AtkTextAttrScale Textattribute = 14
	// weight
	AtkTextAttrWeight Textattribute = 15
	// language
	AtkTextAttrLanguage Textattribute = 16
	// family_name
	AtkTextAttrFamilyName Textattribute = 17
	// bg_color
	AtkTextAttrBgColor Textattribute = 18
	// fg_color
	AtkTextAttrFgColor Textattribute = 19
	// bg_stipple
	AtkTextAttrBgStipple Textattribute = 20
	// fg_stipple
	AtkTextAttrFgStipple Textattribute = 21
	// wrap_mode
	AtkTextAttrWrapMode Textattribute = 22
	// direction
	AtkTextAttrDirection Textattribute = 23
	// justification
	AtkTextAttrJustification Textattribute = 24
	// stretch
	AtkTextAttrStretch Textattribute = 25
	// variant
	AtkTextAttrVariant Textattribute = 26
	// style
	AtkTextAttrStyle Textattribute = 27
	// last_defined
	AtkTextAttrLastDefined Textattribute = 28
)

// Textboundary is a representation of the C type TextBoundary.
type Textboundary int

const (
	// char
	AtkTextBoundaryChar Textboundary = 0
	// word_start
	AtkTextBoundaryWordStart Textboundary = 1
	// word_end
	AtkTextBoundaryWordEnd Textboundary = 2
	// sentence_start
	AtkTextBoundarySentenceStart Textboundary = 3
	// sentence_end
	AtkTextBoundarySentenceEnd Textboundary = 4
	// line_start
	AtkTextBoundaryLineStart Textboundary = 5
	// line_end
	AtkTextBoundaryLineEnd Textboundary = 6
)

// Textcliptype is a representation of the C type TextClipType.
type Textcliptype int

const (
	// none
	AtkTextClipNone Textcliptype = 0
	// min
	AtkTextClipMin Textcliptype = 1
	// max
	AtkTextClipMax Textcliptype = 2
	// both
	AtkTextClipBoth Textcliptype = 3
)

// Textgranularity is a representation of the C type TextGranularity.
type Textgranularity int

const (
	// char
	AtkTextGranularityChar Textgranularity = 0
	// word
	AtkTextGranularityWord Textgranularity = 1
	// sentence
	AtkTextGranularitySentence Textgranularity = 2
	// line
	AtkTextGranularityLine Textgranularity = 3
	// paragraph
	AtkTextGranularityParagraph Textgranularity = 4
)

// Valuetype is a representation of the C type ValueType.
type Valuetype int

const (
	// very_weak
	AtkValueVeryWeak Valuetype = 0
	// weak
	AtkValueWeak Valuetype = 1
	// acceptable
	AtkValueAcceptable Valuetype = 2
	// strong
	AtkValueStrong Valuetype = 3
	// very_strong
	AtkValueVeryStrong Valuetype = 4
	// very_low
	AtkValueVeryLow Valuetype = 5
	// low
	AtkValueLow Valuetype = 6
	// medium
	AtkValueMedium Valuetype = 7
	// high
	AtkValueHigh Valuetype = 8
	// very_high
	AtkValueVeryHigh Valuetype = 9
	// very_bad
	AtkValueVeryBad Valuetype = 10
	// bad
	AtkValueBad Valuetype = 11
	// good
	AtkValueGood Valuetype = 12
	// very_good
	AtkValueVeryGood Valuetype = 13
	// best
	AtkValueBest Valuetype = 14
	// last_defined
	AtkValueLastDefined Valuetype = 15
)
