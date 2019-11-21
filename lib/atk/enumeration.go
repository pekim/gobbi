// Code generated - DO NOT EDIT.

package atk

// CoordType is a representation of the C type AtkCoordType.
//
type CoordType int

const (
	// CoordType_Screen is a representation of the C type ATK_XY_SCREEN.
	CoordType_Screen CoordType = 0
	// CoordType_Window is a representation of the C type ATK_XY_WINDOW.
	CoordType_Window CoordType = 1
	// CoordType_Parent is a representation of the C type ATK_XY_PARENT.
	CoordType_Parent CoordType = 2
)

// KeyEventType is a representation of the C type AtkKeyEventType.
//
type KeyEventType int

const (
	// KeyEventType_Press is a representation of the C type ATK_KEY_EVENT_PRESS.
	KeyEventType_Press KeyEventType = 0
	// KeyEventType_Release is a representation of the C type ATK_KEY_EVENT_RELEASE.
	KeyEventType_Release KeyEventType = 1
	// KeyEventType_LastDefined is a representation of the C type ATK_KEY_EVENT_LAST_DEFINED.
	KeyEventType_LastDefined KeyEventType = 2
)

// Layer is a representation of the C type AtkLayer.
//
type Layer int

const (
	// Layer_Invalid is a representation of the C type ATK_LAYER_INVALID.
	Layer_Invalid Layer = 0
	// Layer_Background is a representation of the C type ATK_LAYER_BACKGROUND.
	Layer_Background Layer = 1
	// Layer_Canvas is a representation of the C type ATK_LAYER_CANVAS.
	Layer_Canvas Layer = 2
	// Layer_Widget is a representation of the C type ATK_LAYER_WIDGET.
	Layer_Widget Layer = 3
	// Layer_Mdi is a representation of the C type ATK_LAYER_MDI.
	Layer_Mdi Layer = 4
	// Layer_Popup is a representation of the C type ATK_LAYER_POPUP.
	Layer_Popup Layer = 5
	// Layer_Overlay is a representation of the C type ATK_LAYER_OVERLAY.
	Layer_Overlay Layer = 6
	// Layer_Window is a representation of the C type ATK_LAYER_WINDOW.
	Layer_Window Layer = 7
)

// RelationType is a representation of the C type AtkRelationType.
//
type RelationType int

const (
	// RelationType_Null is a representation of the C type ATK_RELATION_NULL.
	RelationType_Null RelationType = 0
	// RelationType_ControlledBy is a representation of the C type ATK_RELATION_CONTROLLED_BY.
	RelationType_ControlledBy RelationType = 1
	// RelationType_ControllerFor is a representation of the C type ATK_RELATION_CONTROLLER_FOR.
	RelationType_ControllerFor RelationType = 2
	// RelationType_LabelFor is a representation of the C type ATK_RELATION_LABEL_FOR.
	RelationType_LabelFor RelationType = 3
	// RelationType_LabelledBy is a representation of the C type ATK_RELATION_LABELLED_BY.
	RelationType_LabelledBy RelationType = 4
	// RelationType_MemberOf is a representation of the C type ATK_RELATION_MEMBER_OF.
	RelationType_MemberOf RelationType = 5
	// RelationType_NodeChildOf is a representation of the C type ATK_RELATION_NODE_CHILD_OF.
	RelationType_NodeChildOf RelationType = 6
	// RelationType_FlowsTo is a representation of the C type ATK_RELATION_FLOWS_TO.
	RelationType_FlowsTo RelationType = 7
	// RelationType_FlowsFrom is a representation of the C type ATK_RELATION_FLOWS_FROM.
	RelationType_FlowsFrom RelationType = 8
	// RelationType_SubwindowOf is a representation of the C type ATK_RELATION_SUBWINDOW_OF.
	RelationType_SubwindowOf RelationType = 9
	// RelationType_Embeds is a representation of the C type ATK_RELATION_EMBEDS.
	RelationType_Embeds RelationType = 10
	// RelationType_EmbeddedBy is a representation of the C type ATK_RELATION_EMBEDDED_BY.
	RelationType_EmbeddedBy RelationType = 11
	// RelationType_PopupFor is a representation of the C type ATK_RELATION_POPUP_FOR.
	RelationType_PopupFor RelationType = 12
	// RelationType_ParentWindowOf is a representation of the C type ATK_RELATION_PARENT_WINDOW_OF.
	RelationType_ParentWindowOf RelationType = 13
	// RelationType_DescribedBy is a representation of the C type ATK_RELATION_DESCRIBED_BY.
	RelationType_DescribedBy RelationType = 14
	// RelationType_DescriptionFor is a representation of the C type ATK_RELATION_DESCRIPTION_FOR.
	RelationType_DescriptionFor RelationType = 15
	// RelationType_NodeParentOf is a representation of the C type ATK_RELATION_NODE_PARENT_OF.
	RelationType_NodeParentOf RelationType = 16
	// RelationType_Details is a representation of the C type ATK_RELATION_DETAILS.
	RelationType_Details RelationType = 17
	// RelationType_DetailsFor is a representation of the C type ATK_RELATION_DETAILS_FOR.
	RelationType_DetailsFor RelationType = 18
	// RelationType_ErrorMessage is a representation of the C type ATK_RELATION_ERROR_MESSAGE.
	RelationType_ErrorMessage RelationType = 19
	// RelationType_ErrorFor is a representation of the C type ATK_RELATION_ERROR_FOR.
	RelationType_ErrorFor RelationType = 20
	// RelationType_LastDefined is a representation of the C type ATK_RELATION_LAST_DEFINED.
	RelationType_LastDefined RelationType = 21
)

// Role is a representation of the C type AtkRole.
//
type Role int

const (
	// Role_Invalid is a representation of the C type ATK_ROLE_INVALID.
	Role_Invalid Role = 0
	// Role_AcceleratorLabel is a representation of the C type ATK_ROLE_ACCEL_LABEL.
	Role_AcceleratorLabel Role = 1
	// Role_Alert is a representation of the C type ATK_ROLE_ALERT.
	Role_Alert Role = 2
	// Role_Animation is a representation of the C type ATK_ROLE_ANIMATION.
	Role_Animation Role = 3
	// Role_Arrow is a representation of the C type ATK_ROLE_ARROW.
	Role_Arrow Role = 4
	// Role_Calendar is a representation of the C type ATK_ROLE_CALENDAR.
	Role_Calendar Role = 5
	// Role_Canvas is a representation of the C type ATK_ROLE_CANVAS.
	Role_Canvas Role = 6
	// Role_CheckBox is a representation of the C type ATK_ROLE_CHECK_BOX.
	Role_CheckBox Role = 7
	// Role_CheckMenuItem is a representation of the C type ATK_ROLE_CHECK_MENU_ITEM.
	Role_CheckMenuItem Role = 8
	// Role_ColorChooser is a representation of the C type ATK_ROLE_COLOR_CHOOSER.
	Role_ColorChooser Role = 9
	// Role_ColumnHeader is a representation of the C type ATK_ROLE_COLUMN_HEADER.
	Role_ColumnHeader Role = 10
	// Role_ComboBox is a representation of the C type ATK_ROLE_COMBO_BOX.
	Role_ComboBox Role = 11
	// Role_DateEditor is a representation of the C type ATK_ROLE_DATE_EDITOR.
	Role_DateEditor Role = 12
	// Role_DesktopIcon is a representation of the C type ATK_ROLE_DESKTOP_ICON.
	Role_DesktopIcon Role = 13
	// Role_DesktopFrame is a representation of the C type ATK_ROLE_DESKTOP_FRAME.
	Role_DesktopFrame Role = 14
	// Role_Dial is a representation of the C type ATK_ROLE_DIAL.
	Role_Dial Role = 15
	// Role_Dialog is a representation of the C type ATK_ROLE_DIALOG.
	Role_Dialog Role = 16
	// Role_DirectoryPane is a representation of the C type ATK_ROLE_DIRECTORY_PANE.
	Role_DirectoryPane Role = 17
	// Role_DrawingArea is a representation of the C type ATK_ROLE_DRAWING_AREA.
	Role_DrawingArea Role = 18
	// Role_FileChooser is a representation of the C type ATK_ROLE_FILE_CHOOSER.
	Role_FileChooser Role = 19
	// Role_Filler is a representation of the C type ATK_ROLE_FILLER.
	Role_Filler Role = 20
	// Role_FontChooser is a representation of the C type ATK_ROLE_FONT_CHOOSER.
	Role_FontChooser Role = 21
	// Role_Frame is a representation of the C type ATK_ROLE_FRAME.
	Role_Frame Role = 22
	// Role_GlassPane is a representation of the C type ATK_ROLE_GLASS_PANE.
	Role_GlassPane Role = 23
	// Role_HtmlContainer is a representation of the C type ATK_ROLE_HTML_CONTAINER.
	Role_HtmlContainer Role = 24
	// Role_Icon is a representation of the C type ATK_ROLE_ICON.
	Role_Icon Role = 25
	// Role_Image is a representation of the C type ATK_ROLE_IMAGE.
	Role_Image Role = 26
	// Role_InternalFrame is a representation of the C type ATK_ROLE_INTERNAL_FRAME.
	Role_InternalFrame Role = 27
	// Role_Label is a representation of the C type ATK_ROLE_LABEL.
	Role_Label Role = 28
	// Role_LayeredPane is a representation of the C type ATK_ROLE_LAYERED_PANE.
	Role_LayeredPane Role = 29
	// Role_List is a representation of the C type ATK_ROLE_LIST.
	Role_List Role = 30
	// Role_ListItem is a representation of the C type ATK_ROLE_LIST_ITEM.
	Role_ListItem Role = 31
	// Role_Menu is a representation of the C type ATK_ROLE_MENU.
	Role_Menu Role = 32
	// Role_MenuBar is a representation of the C type ATK_ROLE_MENU_BAR.
	Role_MenuBar Role = 33
	// Role_MenuItem is a representation of the C type ATK_ROLE_MENU_ITEM.
	Role_MenuItem Role = 34
	// Role_OptionPane is a representation of the C type ATK_ROLE_OPTION_PANE.
	Role_OptionPane Role = 35
	// Role_PageTab is a representation of the C type ATK_ROLE_PAGE_TAB.
	Role_PageTab Role = 36
	// Role_PageTabList is a representation of the C type ATK_ROLE_PAGE_TAB_LIST.
	Role_PageTabList Role = 37
	// Role_Panel is a representation of the C type ATK_ROLE_PANEL.
	Role_Panel Role = 38
	// Role_PasswordText is a representation of the C type ATK_ROLE_PASSWORD_TEXT.
	Role_PasswordText Role = 39
	// Role_PopupMenu is a representation of the C type ATK_ROLE_POPUP_MENU.
	Role_PopupMenu Role = 40
	// Role_ProgressBar is a representation of the C type ATK_ROLE_PROGRESS_BAR.
	Role_ProgressBar Role = 41
	// Role_PushButton is a representation of the C type ATK_ROLE_PUSH_BUTTON.
	Role_PushButton Role = 42
	// Role_RadioButton is a representation of the C type ATK_ROLE_RADIO_BUTTON.
	Role_RadioButton Role = 43
	// Role_RadioMenuItem is a representation of the C type ATK_ROLE_RADIO_MENU_ITEM.
	Role_RadioMenuItem Role = 44
	// Role_RootPane is a representation of the C type ATK_ROLE_ROOT_PANE.
	Role_RootPane Role = 45
	// Role_RowHeader is a representation of the C type ATK_ROLE_ROW_HEADER.
	Role_RowHeader Role = 46
	// Role_ScrollBar is a representation of the C type ATK_ROLE_SCROLL_BAR.
	Role_ScrollBar Role = 47
	// Role_ScrollPane is a representation of the C type ATK_ROLE_SCROLL_PANE.
	Role_ScrollPane Role = 48
	// Role_Separator is a representation of the C type ATK_ROLE_SEPARATOR.
	Role_Separator Role = 49
	// Role_Slider is a representation of the C type ATK_ROLE_SLIDER.
	Role_Slider Role = 50
	// Role_SplitPane is a representation of the C type ATK_ROLE_SPLIT_PANE.
	Role_SplitPane Role = 51
	// Role_SpinButton is a representation of the C type ATK_ROLE_SPIN_BUTTON.
	Role_SpinButton Role = 52
	// Role_Statusbar is a representation of the C type ATK_ROLE_STATUSBAR.
	Role_Statusbar Role = 53
	// Role_Table is a representation of the C type ATK_ROLE_TABLE.
	Role_Table Role = 54
	// Role_TableCell is a representation of the C type ATK_ROLE_TABLE_CELL.
	Role_TableCell Role = 55
	// Role_TableColumnHeader is a representation of the C type ATK_ROLE_TABLE_COLUMN_HEADER.
	Role_TableColumnHeader Role = 56
	// Role_TableRowHeader is a representation of the C type ATK_ROLE_TABLE_ROW_HEADER.
	Role_TableRowHeader Role = 57
	// Role_TearOffMenuItem is a representation of the C type ATK_ROLE_TEAR_OFF_MENU_ITEM.
	Role_TearOffMenuItem Role = 58
	// Role_Terminal is a representation of the C type ATK_ROLE_TERMINAL.
	Role_Terminal Role = 59
	// Role_Text is a representation of the C type ATK_ROLE_TEXT.
	Role_Text Role = 60
	// Role_ToggleButton is a representation of the C type ATK_ROLE_TOGGLE_BUTTON.
	Role_ToggleButton Role = 61
	// Role_ToolBar is a representation of the C type ATK_ROLE_TOOL_BAR.
	Role_ToolBar Role = 62
	// Role_ToolTip is a representation of the C type ATK_ROLE_TOOL_TIP.
	Role_ToolTip Role = 63
	// Role_Tree is a representation of the C type ATK_ROLE_TREE.
	Role_Tree Role = 64
	// Role_TreeTable is a representation of the C type ATK_ROLE_TREE_TABLE.
	Role_TreeTable Role = 65
	// Role_Unknown is a representation of the C type ATK_ROLE_UNKNOWN.
	Role_Unknown Role = 66
	// Role_Viewport is a representation of the C type ATK_ROLE_VIEWPORT.
	Role_Viewport Role = 67
	// Role_Window is a representation of the C type ATK_ROLE_WINDOW.
	Role_Window Role = 68
	// Role_Header is a representation of the C type ATK_ROLE_HEADER.
	Role_Header Role = 69
	// Role_Footer is a representation of the C type ATK_ROLE_FOOTER.
	Role_Footer Role = 70
	// Role_Paragraph is a representation of the C type ATK_ROLE_PARAGRAPH.
	Role_Paragraph Role = 71
	// Role_Ruler is a representation of the C type ATK_ROLE_RULER.
	Role_Ruler Role = 72
	// Role_Application is a representation of the C type ATK_ROLE_APPLICATION.
	Role_Application Role = 73
	// Role_Autocomplete is a representation of the C type ATK_ROLE_AUTOCOMPLETE.
	Role_Autocomplete Role = 74
	// Role_EditBar is a representation of the C type ATK_ROLE_EDITBAR.
	Role_EditBar Role = 75
	// Role_Embedded is a representation of the C type ATK_ROLE_EMBEDDED.
	Role_Embedded Role = 76
	// Role_Entry is a representation of the C type ATK_ROLE_ENTRY.
	Role_Entry Role = 77
	// Role_Chart is a representation of the C type ATK_ROLE_CHART.
	Role_Chart Role = 78
	// Role_Caption is a representation of the C type ATK_ROLE_CAPTION.
	Role_Caption Role = 79
	// Role_DocumentFrame is a representation of the C type ATK_ROLE_DOCUMENT_FRAME.
	Role_DocumentFrame Role = 80
	// Role_Heading is a representation of the C type ATK_ROLE_HEADING.
	Role_Heading Role = 81
	// Role_Page is a representation of the C type ATK_ROLE_PAGE.
	Role_Page Role = 82
	// Role_Section is a representation of the C type ATK_ROLE_SECTION.
	Role_Section Role = 83
	// Role_RedundantObject is a representation of the C type ATK_ROLE_REDUNDANT_OBJECT.
	Role_RedundantObject Role = 84
	// Role_Form is a representation of the C type ATK_ROLE_FORM.
	Role_Form Role = 85
	// Role_Link is a representation of the C type ATK_ROLE_LINK.
	Role_Link Role = 86
	// Role_InputMethodWindow is a representation of the C type ATK_ROLE_INPUT_METHOD_WINDOW.
	Role_InputMethodWindow Role = 87
	// Role_TableRow is a representation of the C type ATK_ROLE_TABLE_ROW.
	Role_TableRow Role = 88
	// Role_TreeItem is a representation of the C type ATK_ROLE_TREE_ITEM.
	Role_TreeItem Role = 89
	// Role_DocumentSpreadsheet is a representation of the C type ATK_ROLE_DOCUMENT_SPREADSHEET.
	Role_DocumentSpreadsheet Role = 90
	// Role_DocumentPresentation is a representation of the C type ATK_ROLE_DOCUMENT_PRESENTATION.
	Role_DocumentPresentation Role = 91
	// Role_DocumentText is a representation of the C type ATK_ROLE_DOCUMENT_TEXT.
	Role_DocumentText Role = 92
	// Role_DocumentWeb is a representation of the C type ATK_ROLE_DOCUMENT_WEB.
	Role_DocumentWeb Role = 93
	// Role_DocumentEmail is a representation of the C type ATK_ROLE_DOCUMENT_EMAIL.
	Role_DocumentEmail Role = 94
	// Role_Comment is a representation of the C type ATK_ROLE_COMMENT.
	Role_Comment Role = 95
	// Role_ListBox is a representation of the C type ATK_ROLE_LIST_BOX.
	Role_ListBox Role = 96
	// Role_Grouping is a representation of the C type ATK_ROLE_GROUPING.
	Role_Grouping Role = 97
	// Role_ImageMap is a representation of the C type ATK_ROLE_IMAGE_MAP.
	Role_ImageMap Role = 98
	// Role_Notification is a representation of the C type ATK_ROLE_NOTIFICATION.
	Role_Notification Role = 99
	// Role_InfoBar is a representation of the C type ATK_ROLE_INFO_BAR.
	Role_InfoBar Role = 100
	// Role_LevelBar is a representation of the C type ATK_ROLE_LEVEL_BAR.
	Role_LevelBar Role = 101
	// Role_TitleBar is a representation of the C type ATK_ROLE_TITLE_BAR.
	Role_TitleBar Role = 102
	// Role_BlockQuote is a representation of the C type ATK_ROLE_BLOCK_QUOTE.
	Role_BlockQuote Role = 103
	// Role_Audio is a representation of the C type ATK_ROLE_AUDIO.
	Role_Audio Role = 104
	// Role_Video is a representation of the C type ATK_ROLE_VIDEO.
	Role_Video Role = 105
	// Role_Definition is a representation of the C type ATK_ROLE_DEFINITION.
	Role_Definition Role = 106
	// Role_Article is a representation of the C type ATK_ROLE_ARTICLE.
	Role_Article Role = 107
	// Role_Landmark is a representation of the C type ATK_ROLE_LANDMARK.
	Role_Landmark Role = 108
	// Role_Log is a representation of the C type ATK_ROLE_LOG.
	Role_Log Role = 109
	// Role_Marquee is a representation of the C type ATK_ROLE_MARQUEE.
	Role_Marquee Role = 110
	// Role_Math is a representation of the C type ATK_ROLE_MATH.
	Role_Math Role = 111
	// Role_Rating is a representation of the C type ATK_ROLE_RATING.
	Role_Rating Role = 112
	// Role_Timer is a representation of the C type ATK_ROLE_TIMER.
	Role_Timer Role = 113
	// Role_DescriptionList is a representation of the C type ATK_ROLE_DESCRIPTION_LIST.
	Role_DescriptionList Role = 114
	// Role_DescriptionTerm is a representation of the C type ATK_ROLE_DESCRIPTION_TERM.
	Role_DescriptionTerm Role = 115
	// Role_DescriptionValue is a representation of the C type ATK_ROLE_DESCRIPTION_VALUE.
	Role_DescriptionValue Role = 116
	// Role_Static is a representation of the C type ATK_ROLE_STATIC.
	Role_Static Role = 117
	// Role_MathFraction is a representation of the C type ATK_ROLE_MATH_FRACTION.
	Role_MathFraction Role = 118
	// Role_MathRoot is a representation of the C type ATK_ROLE_MATH_ROOT.
	Role_MathRoot Role = 119
	// Role_Subscript is a representation of the C type ATK_ROLE_SUBSCRIPT.
	Role_Subscript Role = 120
	// Role_Superscript is a representation of the C type ATK_ROLE_SUPERSCRIPT.
	Role_Superscript Role = 121
	// Role_Footnote is a representation of the C type ATK_ROLE_FOOTNOTE.
	Role_Footnote Role = 122
	// Role_ContentDeletion is a representation of the C type ATK_ROLE_CONTENT_DELETION.
	Role_ContentDeletion Role = 123
	// Role_ContentInsertion is a representation of the C type ATK_ROLE_CONTENT_INSERTION.
	Role_ContentInsertion Role = 124
	// Role_LastDefined is a representation of the C type ATK_ROLE_LAST_DEFINED.
	Role_LastDefined Role = 125
)

// ScrollType is a representation of the C type AtkScrollType.
//
// since 2.30
type ScrollType int

const (
	// ScrollType_TopLeft is a representation of the C type ATK_SCROLL_TOP_LEFT.
	ScrollType_TopLeft ScrollType = 0
	// ScrollType_BottomRight is a representation of the C type ATK_SCROLL_BOTTOM_RIGHT.
	ScrollType_BottomRight ScrollType = 1
	// ScrollType_TopEdge is a representation of the C type ATK_SCROLL_TOP_EDGE.
	ScrollType_TopEdge ScrollType = 2
	// ScrollType_BottomEdge is a representation of the C type ATK_SCROLL_BOTTOM_EDGE.
	ScrollType_BottomEdge ScrollType = 3
	// ScrollType_LeftEdge is a representation of the C type ATK_SCROLL_LEFT_EDGE.
	ScrollType_LeftEdge ScrollType = 4
	// ScrollType_RightEdge is a representation of the C type ATK_SCROLL_RIGHT_EDGE.
	ScrollType_RightEdge ScrollType = 5
	// ScrollType_Anywhere is a representation of the C type ATK_SCROLL_ANYWHERE.
	ScrollType_Anywhere ScrollType = 6
)

// StateType is a representation of the C type AtkStateType.
//
type StateType int

const (
	// StateType_Invalid is a representation of the C type ATK_STATE_INVALID.
	StateType_Invalid StateType = 0
	// StateType_Active is a representation of the C type ATK_STATE_ACTIVE.
	StateType_Active StateType = 1
	// StateType_Armed is a representation of the C type ATK_STATE_ARMED.
	StateType_Armed StateType = 2
	// StateType_Busy is a representation of the C type ATK_STATE_BUSY.
	StateType_Busy StateType = 3
	// StateType_Checked is a representation of the C type ATK_STATE_CHECKED.
	StateType_Checked StateType = 4
	// StateType_Defunct is a representation of the C type ATK_STATE_DEFUNCT.
	StateType_Defunct StateType = 5
	// StateType_Editable is a representation of the C type ATK_STATE_EDITABLE.
	StateType_Editable StateType = 6
	// StateType_Enabled is a representation of the C type ATK_STATE_ENABLED.
	StateType_Enabled StateType = 7
	// StateType_Expandable is a representation of the C type ATK_STATE_EXPANDABLE.
	StateType_Expandable StateType = 8
	// StateType_Expanded is a representation of the C type ATK_STATE_EXPANDED.
	StateType_Expanded StateType = 9
	// StateType_Focusable is a representation of the C type ATK_STATE_FOCUSABLE.
	StateType_Focusable StateType = 10
	// StateType_Focused is a representation of the C type ATK_STATE_FOCUSED.
	StateType_Focused StateType = 11
	// StateType_Horizontal is a representation of the C type ATK_STATE_HORIZONTAL.
	StateType_Horizontal StateType = 12
	// StateType_Iconified is a representation of the C type ATK_STATE_ICONIFIED.
	StateType_Iconified StateType = 13
	// StateType_Modal is a representation of the C type ATK_STATE_MODAL.
	StateType_Modal StateType = 14
	// StateType_MultiLine is a representation of the C type ATK_STATE_MULTI_LINE.
	StateType_MultiLine StateType = 15
	// StateType_Multiselectable is a representation of the C type ATK_STATE_MULTISELECTABLE.
	StateType_Multiselectable StateType = 16
	// StateType_Opaque is a representation of the C type ATK_STATE_OPAQUE.
	StateType_Opaque StateType = 17
	// StateType_Pressed is a representation of the C type ATK_STATE_PRESSED.
	StateType_Pressed StateType = 18
	// StateType_Resizable is a representation of the C type ATK_STATE_RESIZABLE.
	StateType_Resizable StateType = 19
	// StateType_Selectable is a representation of the C type ATK_STATE_SELECTABLE.
	StateType_Selectable StateType = 20
	// StateType_Selected is a representation of the C type ATK_STATE_SELECTED.
	StateType_Selected StateType = 21
	// StateType_Sensitive is a representation of the C type ATK_STATE_SENSITIVE.
	StateType_Sensitive StateType = 22
	// StateType_Showing is a representation of the C type ATK_STATE_SHOWING.
	StateType_Showing StateType = 23
	// StateType_SingleLine is a representation of the C type ATK_STATE_SINGLE_LINE.
	StateType_SingleLine StateType = 24
	// StateType_Stale is a representation of the C type ATK_STATE_STALE.
	StateType_Stale StateType = 25
	// StateType_Transient is a representation of the C type ATK_STATE_TRANSIENT.
	StateType_Transient StateType = 26
	// StateType_Vertical is a representation of the C type ATK_STATE_VERTICAL.
	StateType_Vertical StateType = 27
	// StateType_Visible is a representation of the C type ATK_STATE_VISIBLE.
	StateType_Visible StateType = 28
	// StateType_ManagesDescendants is a representation of the C type ATK_STATE_MANAGES_DESCENDANTS.
	StateType_ManagesDescendants StateType = 29
	// StateType_Indeterminate is a representation of the C type ATK_STATE_INDETERMINATE.
	StateType_Indeterminate StateType = 30
	// StateType_Truncated is a representation of the C type ATK_STATE_TRUNCATED.
	StateType_Truncated StateType = 31
	// StateType_Required is a representation of the C type ATK_STATE_REQUIRED.
	StateType_Required StateType = 32
	// StateType_InvalidEntry is a representation of the C type ATK_STATE_INVALID_ENTRY.
	StateType_InvalidEntry StateType = 33
	// StateType_SupportsAutocompletion is a representation of the C type ATK_STATE_SUPPORTS_AUTOCOMPLETION.
	StateType_SupportsAutocompletion StateType = 34
	// StateType_SelectableText is a representation of the C type ATK_STATE_SELECTABLE_TEXT.
	StateType_SelectableText StateType = 35
	// StateType_Default is a representation of the C type ATK_STATE_DEFAULT.
	StateType_Default StateType = 36
	// StateType_Animated is a representation of the C type ATK_STATE_ANIMATED.
	StateType_Animated StateType = 37
	// StateType_Visited is a representation of the C type ATK_STATE_VISITED.
	StateType_Visited StateType = 38
	// StateType_Checkable is a representation of the C type ATK_STATE_CHECKABLE.
	StateType_Checkable StateType = 39
	// StateType_HasPopup is a representation of the C type ATK_STATE_HAS_POPUP.
	StateType_HasPopup StateType = 40
	// StateType_HasTooltip is a representation of the C type ATK_STATE_HAS_TOOLTIP.
	StateType_HasTooltip StateType = 41
	// StateType_ReadOnly is a representation of the C type ATK_STATE_READ_ONLY.
	StateType_ReadOnly StateType = 42
	// StateType_LastDefined is a representation of the C type ATK_STATE_LAST_DEFINED.
	StateType_LastDefined StateType = 43
)

// TextAttribute is a representation of the C type AtkTextAttribute.
//
type TextAttribute int

const (
	// TextAttribute_Invalid is a representation of the C type ATK_TEXT_ATTR_INVALID.
	TextAttribute_Invalid TextAttribute = 0
	// TextAttribute_LeftMargin is a representation of the C type ATK_TEXT_ATTR_LEFT_MARGIN.
	TextAttribute_LeftMargin TextAttribute = 1
	// TextAttribute_RightMargin is a representation of the C type ATK_TEXT_ATTR_RIGHT_MARGIN.
	TextAttribute_RightMargin TextAttribute = 2
	// TextAttribute_Indent is a representation of the C type ATK_TEXT_ATTR_INDENT.
	TextAttribute_Indent TextAttribute = 3
	// TextAttribute_Invisible is a representation of the C type ATK_TEXT_ATTR_INVISIBLE.
	TextAttribute_Invisible TextAttribute = 4
	// TextAttribute_Editable is a representation of the C type ATK_TEXT_ATTR_EDITABLE.
	TextAttribute_Editable TextAttribute = 5
	// TextAttribute_PixelsAboveLines is a representation of the C type ATK_TEXT_ATTR_PIXELS_ABOVE_LINES.
	TextAttribute_PixelsAboveLines TextAttribute = 6
	// TextAttribute_PixelsBelowLines is a representation of the C type ATK_TEXT_ATTR_PIXELS_BELOW_LINES.
	TextAttribute_PixelsBelowLines TextAttribute = 7
	// TextAttribute_PixelsInsideWrap is a representation of the C type ATK_TEXT_ATTR_PIXELS_INSIDE_WRAP.
	TextAttribute_PixelsInsideWrap TextAttribute = 8
	// TextAttribute_BgFullHeight is a representation of the C type ATK_TEXT_ATTR_BG_FULL_HEIGHT.
	TextAttribute_BgFullHeight TextAttribute = 9
	// TextAttribute_Rise is a representation of the C type ATK_TEXT_ATTR_RISE.
	TextAttribute_Rise TextAttribute = 10
	// TextAttribute_Underline is a representation of the C type ATK_TEXT_ATTR_UNDERLINE.
	TextAttribute_Underline TextAttribute = 11
	// TextAttribute_Strikethrough is a representation of the C type ATK_TEXT_ATTR_STRIKETHROUGH.
	TextAttribute_Strikethrough TextAttribute = 12
	// TextAttribute_Size is a representation of the C type ATK_TEXT_ATTR_SIZE.
	TextAttribute_Size TextAttribute = 13
	// TextAttribute_Scale is a representation of the C type ATK_TEXT_ATTR_SCALE.
	TextAttribute_Scale TextAttribute = 14
	// TextAttribute_Weight is a representation of the C type ATK_TEXT_ATTR_WEIGHT.
	TextAttribute_Weight TextAttribute = 15
	// TextAttribute_Language is a representation of the C type ATK_TEXT_ATTR_LANGUAGE.
	TextAttribute_Language TextAttribute = 16
	// TextAttribute_FamilyName is a representation of the C type ATK_TEXT_ATTR_FAMILY_NAME.
	TextAttribute_FamilyName TextAttribute = 17
	// TextAttribute_BgColor is a representation of the C type ATK_TEXT_ATTR_BG_COLOR.
	TextAttribute_BgColor TextAttribute = 18
	// TextAttribute_FgColor is a representation of the C type ATK_TEXT_ATTR_FG_COLOR.
	TextAttribute_FgColor TextAttribute = 19
	// TextAttribute_BgStipple is a representation of the C type ATK_TEXT_ATTR_BG_STIPPLE.
	TextAttribute_BgStipple TextAttribute = 20
	// TextAttribute_FgStipple is a representation of the C type ATK_TEXT_ATTR_FG_STIPPLE.
	TextAttribute_FgStipple TextAttribute = 21
	// TextAttribute_WrapMode is a representation of the C type ATK_TEXT_ATTR_WRAP_MODE.
	TextAttribute_WrapMode TextAttribute = 22
	// TextAttribute_Direction is a representation of the C type ATK_TEXT_ATTR_DIRECTION.
	TextAttribute_Direction TextAttribute = 23
	// TextAttribute_Justification is a representation of the C type ATK_TEXT_ATTR_JUSTIFICATION.
	TextAttribute_Justification TextAttribute = 24
	// TextAttribute_Stretch is a representation of the C type ATK_TEXT_ATTR_STRETCH.
	TextAttribute_Stretch TextAttribute = 25
	// TextAttribute_Variant is a representation of the C type ATK_TEXT_ATTR_VARIANT.
	TextAttribute_Variant TextAttribute = 26
	// TextAttribute_Style is a representation of the C type ATK_TEXT_ATTR_STYLE.
	TextAttribute_Style TextAttribute = 27
	// TextAttribute_LastDefined is a representation of the C type ATK_TEXT_ATTR_LAST_DEFINED.
	TextAttribute_LastDefined TextAttribute = 28
)

// TextBoundary is a representation of the C type AtkTextBoundary.
//
type TextBoundary int

const (
	// TextBoundary_Char is a representation of the C type ATK_TEXT_BOUNDARY_CHAR.
	TextBoundary_Char TextBoundary = 0
	// TextBoundary_WordStart is a representation of the C type ATK_TEXT_BOUNDARY_WORD_START.
	TextBoundary_WordStart TextBoundary = 1
	// TextBoundary_WordEnd is a representation of the C type ATK_TEXT_BOUNDARY_WORD_END.
	TextBoundary_WordEnd TextBoundary = 2
	// TextBoundary_SentenceStart is a representation of the C type ATK_TEXT_BOUNDARY_SENTENCE_START.
	TextBoundary_SentenceStart TextBoundary = 3
	// TextBoundary_SentenceEnd is a representation of the C type ATK_TEXT_BOUNDARY_SENTENCE_END.
	TextBoundary_SentenceEnd TextBoundary = 4
	// TextBoundary_LineStart is a representation of the C type ATK_TEXT_BOUNDARY_LINE_START.
	TextBoundary_LineStart TextBoundary = 5
	// TextBoundary_LineEnd is a representation of the C type ATK_TEXT_BOUNDARY_LINE_END.
	TextBoundary_LineEnd TextBoundary = 6
)

// TextClipType is a representation of the C type AtkTextClipType.
//
type TextClipType int

const (
	// TextClipType_None is a representation of the C type ATK_TEXT_CLIP_NONE.
	TextClipType_None TextClipType = 0
	// TextClipType_Min is a representation of the C type ATK_TEXT_CLIP_MIN.
	TextClipType_Min TextClipType = 1
	// TextClipType_Max is a representation of the C type ATK_TEXT_CLIP_MAX.
	TextClipType_Max TextClipType = 2
	// TextClipType_Both is a representation of the C type ATK_TEXT_CLIP_BOTH.
	TextClipType_Both TextClipType = 3
)

// TextGranularity is a representation of the C type AtkTextGranularity.
//
type TextGranularity int

const (
	// TextGranularity_Char is a representation of the C type ATK_TEXT_GRANULARITY_CHAR.
	TextGranularity_Char TextGranularity = 0
	// TextGranularity_Word is a representation of the C type ATK_TEXT_GRANULARITY_WORD.
	TextGranularity_Word TextGranularity = 1
	// TextGranularity_Sentence is a representation of the C type ATK_TEXT_GRANULARITY_SENTENCE.
	TextGranularity_Sentence TextGranularity = 2
	// TextGranularity_Line is a representation of the C type ATK_TEXT_GRANULARITY_LINE.
	TextGranularity_Line TextGranularity = 3
	// TextGranularity_Paragraph is a representation of the C type ATK_TEXT_GRANULARITY_PARAGRAPH.
	TextGranularity_Paragraph TextGranularity = 4
)

// ValueType is a representation of the C type AtkValueType.
//
type ValueType int

const (
	// ValueType_VeryWeak is a representation of the C type ATK_VALUE_VERY_WEAK.
	ValueType_VeryWeak ValueType = 0
	// ValueType_Weak is a representation of the C type ATK_VALUE_WEAK.
	ValueType_Weak ValueType = 1
	// ValueType_Acceptable is a representation of the C type ATK_VALUE_ACCEPTABLE.
	ValueType_Acceptable ValueType = 2
	// ValueType_Strong is a representation of the C type ATK_VALUE_STRONG.
	ValueType_Strong ValueType = 3
	// ValueType_VeryStrong is a representation of the C type ATK_VALUE_VERY_STRONG.
	ValueType_VeryStrong ValueType = 4
	// ValueType_VeryLow is a representation of the C type ATK_VALUE_VERY_LOW.
	ValueType_VeryLow ValueType = 5
	// ValueType_Low is a representation of the C type ATK_VALUE_LOW.
	ValueType_Low ValueType = 6
	// ValueType_Medium is a representation of the C type ATK_VALUE_MEDIUM.
	ValueType_Medium ValueType = 7
	// ValueType_High is a representation of the C type ATK_VALUE_HIGH.
	ValueType_High ValueType = 8
	// ValueType_VeryHigh is a representation of the C type ATK_VALUE_VERY_HIGH.
	ValueType_VeryHigh ValueType = 9
	// ValueType_VeryBad is a representation of the C type ATK_VALUE_VERY_BAD.
	ValueType_VeryBad ValueType = 10
	// ValueType_Bad is a representation of the C type ATK_VALUE_BAD.
	ValueType_Bad ValueType = 11
	// ValueType_Good is a representation of the C type ATK_VALUE_GOOD.
	ValueType_Good ValueType = 12
	// ValueType_VeryGood is a representation of the C type ATK_VALUE_VERY_GOOD.
	ValueType_VeryGood ValueType = 13
	// ValueType_Best is a representation of the C type ATK_VALUE_BEST.
	ValueType_Best ValueType = 14
	// ValueType_LastDefined is a representation of the C type ATK_VALUE_LAST_DEFINED.
	ValueType_LastDefined ValueType = 15
)
