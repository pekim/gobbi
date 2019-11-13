// Code generated - DO NOT EDIT.

package atk

// Coordtype is a representation of the C type AtkCoordType.
type Coordtype int

const (
	// Coordtype_Screen is a representation of the C type ATK_XY_SCREEN.
	Coordtype_Screen Coordtype = 0
	// Coordtype_Window is a representation of the C type ATK_XY_WINDOW.
	Coordtype_Window Coordtype = 1
)

// Keyeventtype is a representation of the C type AtkKeyEventType.
type Keyeventtype int

const (
	// Keyeventtype_Press is a representation of the C type ATK_KEY_EVENT_PRESS.
	Keyeventtype_Press Keyeventtype = 0
	// Keyeventtype_Release is a representation of the C type ATK_KEY_EVENT_RELEASE.
	Keyeventtype_Release Keyeventtype = 1
	// Keyeventtype_LastDefined is a representation of the C type ATK_KEY_EVENT_LAST_DEFINED.
	Keyeventtype_LastDefined Keyeventtype = 2
)

// Layer is a representation of the C type AtkLayer.
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

// Relationtype is a representation of the C type AtkRelationType.
type Relationtype int

const (
	// Relationtype_Null is a representation of the C type ATK_RELATION_NULL.
	Relationtype_Null Relationtype = 0
	// Relationtype_ControlledBy is a representation of the C type ATK_RELATION_CONTROLLED_BY.
	Relationtype_ControlledBy Relationtype = 1
	// Relationtype_ControllerFor is a representation of the C type ATK_RELATION_CONTROLLER_FOR.
	Relationtype_ControllerFor Relationtype = 2
	// Relationtype_LabelFor is a representation of the C type ATK_RELATION_LABEL_FOR.
	Relationtype_LabelFor Relationtype = 3
	// Relationtype_LabelledBy is a representation of the C type ATK_RELATION_LABELLED_BY.
	Relationtype_LabelledBy Relationtype = 4
	// Relationtype_MemberOf is a representation of the C type ATK_RELATION_MEMBER_OF.
	Relationtype_MemberOf Relationtype = 5
	// Relationtype_NodeChildOf is a representation of the C type ATK_RELATION_NODE_CHILD_OF.
	Relationtype_NodeChildOf Relationtype = 6
	// Relationtype_FlowsTo is a representation of the C type ATK_RELATION_FLOWS_TO.
	Relationtype_FlowsTo Relationtype = 7
	// Relationtype_FlowsFrom is a representation of the C type ATK_RELATION_FLOWS_FROM.
	Relationtype_FlowsFrom Relationtype = 8
	// Relationtype_SubwindowOf is a representation of the C type ATK_RELATION_SUBWINDOW_OF.
	Relationtype_SubwindowOf Relationtype = 9
	// Relationtype_Embeds is a representation of the C type ATK_RELATION_EMBEDS.
	Relationtype_Embeds Relationtype = 10
	// Relationtype_EmbeddedBy is a representation of the C type ATK_RELATION_EMBEDDED_BY.
	Relationtype_EmbeddedBy Relationtype = 11
	// Relationtype_PopupFor is a representation of the C type ATK_RELATION_POPUP_FOR.
	Relationtype_PopupFor Relationtype = 12
	// Relationtype_ParentWindowOf is a representation of the C type ATK_RELATION_PARENT_WINDOW_OF.
	Relationtype_ParentWindowOf Relationtype = 13
	// Relationtype_DescribedBy is a representation of the C type ATK_RELATION_DESCRIBED_BY.
	Relationtype_DescribedBy Relationtype = 14
	// Relationtype_DescriptionFor is a representation of the C type ATK_RELATION_DESCRIPTION_FOR.
	Relationtype_DescriptionFor Relationtype = 15
	// Relationtype_NodeParentOf is a representation of the C type ATK_RELATION_NODE_PARENT_OF.
	Relationtype_NodeParentOf Relationtype = 16
	// Relationtype_Details is a representation of the C type ATK_RELATION_DETAILS.
	Relationtype_Details Relationtype = 17
	// Relationtype_DetailsFor is a representation of the C type ATK_RELATION_DETAILS_FOR.
	Relationtype_DetailsFor Relationtype = 18
	// Relationtype_ErrorMessage is a representation of the C type ATK_RELATION_ERROR_MESSAGE.
	Relationtype_ErrorMessage Relationtype = 19
	// Relationtype_ErrorFor is a representation of the C type ATK_RELATION_ERROR_FOR.
	Relationtype_ErrorFor Relationtype = 20
	// Relationtype_LastDefined is a representation of the C type ATK_RELATION_LAST_DEFINED.
	Relationtype_LastDefined Relationtype = 21
)

// Role is a representation of the C type AtkRole.
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
	// Role_LastDefined is a representation of the C type ATK_ROLE_LAST_DEFINED.
	Role_LastDefined Role = 123
)

// Statetype is a representation of the C type AtkStateType.
type Statetype int

const (
	// Statetype_Invalid is a representation of the C type ATK_STATE_INVALID.
	Statetype_Invalid Statetype = 0
	// Statetype_Active is a representation of the C type ATK_STATE_ACTIVE.
	Statetype_Active Statetype = 1
	// Statetype_Armed is a representation of the C type ATK_STATE_ARMED.
	Statetype_Armed Statetype = 2
	// Statetype_Busy is a representation of the C type ATK_STATE_BUSY.
	Statetype_Busy Statetype = 3
	// Statetype_Checked is a representation of the C type ATK_STATE_CHECKED.
	Statetype_Checked Statetype = 4
	// Statetype_Defunct is a representation of the C type ATK_STATE_DEFUNCT.
	Statetype_Defunct Statetype = 5
	// Statetype_Editable is a representation of the C type ATK_STATE_EDITABLE.
	Statetype_Editable Statetype = 6
	// Statetype_Enabled is a representation of the C type ATK_STATE_ENABLED.
	Statetype_Enabled Statetype = 7
	// Statetype_Expandable is a representation of the C type ATK_STATE_EXPANDABLE.
	Statetype_Expandable Statetype = 8
	// Statetype_Expanded is a representation of the C type ATK_STATE_EXPANDED.
	Statetype_Expanded Statetype = 9
	// Statetype_Focusable is a representation of the C type ATK_STATE_FOCUSABLE.
	Statetype_Focusable Statetype = 10
	// Statetype_Focused is a representation of the C type ATK_STATE_FOCUSED.
	Statetype_Focused Statetype = 11
	// Statetype_Horizontal is a representation of the C type ATK_STATE_HORIZONTAL.
	Statetype_Horizontal Statetype = 12
	// Statetype_Iconified is a representation of the C type ATK_STATE_ICONIFIED.
	Statetype_Iconified Statetype = 13
	// Statetype_Modal is a representation of the C type ATK_STATE_MODAL.
	Statetype_Modal Statetype = 14
	// Statetype_MultiLine is a representation of the C type ATK_STATE_MULTI_LINE.
	Statetype_MultiLine Statetype = 15
	// Statetype_Multiselectable is a representation of the C type ATK_STATE_MULTISELECTABLE.
	Statetype_Multiselectable Statetype = 16
	// Statetype_Opaque is a representation of the C type ATK_STATE_OPAQUE.
	Statetype_Opaque Statetype = 17
	// Statetype_Pressed is a representation of the C type ATK_STATE_PRESSED.
	Statetype_Pressed Statetype = 18
	// Statetype_Resizable is a representation of the C type ATK_STATE_RESIZABLE.
	Statetype_Resizable Statetype = 19
	// Statetype_Selectable is a representation of the C type ATK_STATE_SELECTABLE.
	Statetype_Selectable Statetype = 20
	// Statetype_Selected is a representation of the C type ATK_STATE_SELECTED.
	Statetype_Selected Statetype = 21
	// Statetype_Sensitive is a representation of the C type ATK_STATE_SENSITIVE.
	Statetype_Sensitive Statetype = 22
	// Statetype_Showing is a representation of the C type ATK_STATE_SHOWING.
	Statetype_Showing Statetype = 23
	// Statetype_SingleLine is a representation of the C type ATK_STATE_SINGLE_LINE.
	Statetype_SingleLine Statetype = 24
	// Statetype_Stale is a representation of the C type ATK_STATE_STALE.
	Statetype_Stale Statetype = 25
	// Statetype_Transient is a representation of the C type ATK_STATE_TRANSIENT.
	Statetype_Transient Statetype = 26
	// Statetype_Vertical is a representation of the C type ATK_STATE_VERTICAL.
	Statetype_Vertical Statetype = 27
	// Statetype_Visible is a representation of the C type ATK_STATE_VISIBLE.
	Statetype_Visible Statetype = 28
	// Statetype_ManagesDescendants is a representation of the C type ATK_STATE_MANAGES_DESCENDANTS.
	Statetype_ManagesDescendants Statetype = 29
	// Statetype_Indeterminate is a representation of the C type ATK_STATE_INDETERMINATE.
	Statetype_Indeterminate Statetype = 30
	// Statetype_Truncated is a representation of the C type ATK_STATE_TRUNCATED.
	Statetype_Truncated Statetype = 31
	// Statetype_Required is a representation of the C type ATK_STATE_REQUIRED.
	Statetype_Required Statetype = 32
	// Statetype_InvalidEntry is a representation of the C type ATK_STATE_INVALID_ENTRY.
	Statetype_InvalidEntry Statetype = 33
	// Statetype_SupportsAutocompletion is a representation of the C type ATK_STATE_SUPPORTS_AUTOCOMPLETION.
	Statetype_SupportsAutocompletion Statetype = 34
	// Statetype_SelectableText is a representation of the C type ATK_STATE_SELECTABLE_TEXT.
	Statetype_SelectableText Statetype = 35
	// Statetype_Default is a representation of the C type ATK_STATE_DEFAULT.
	Statetype_Default Statetype = 36
	// Statetype_Animated is a representation of the C type ATK_STATE_ANIMATED.
	Statetype_Animated Statetype = 37
	// Statetype_Visited is a representation of the C type ATK_STATE_VISITED.
	Statetype_Visited Statetype = 38
	// Statetype_Checkable is a representation of the C type ATK_STATE_CHECKABLE.
	Statetype_Checkable Statetype = 39
	// Statetype_HasPopup is a representation of the C type ATK_STATE_HAS_POPUP.
	Statetype_HasPopup Statetype = 40
	// Statetype_HasTooltip is a representation of the C type ATK_STATE_HAS_TOOLTIP.
	Statetype_HasTooltip Statetype = 41
	// Statetype_ReadOnly is a representation of the C type ATK_STATE_READ_ONLY.
	Statetype_ReadOnly Statetype = 42
	// Statetype_LastDefined is a representation of the C type ATK_STATE_LAST_DEFINED.
	Statetype_LastDefined Statetype = 43
)

// Textattribute is a representation of the C type AtkTextAttribute.
type Textattribute int

const (
	// Textattribute_Invalid is a representation of the C type ATK_TEXT_ATTR_INVALID.
	Textattribute_Invalid Textattribute = 0
	// Textattribute_LeftMargin is a representation of the C type ATK_TEXT_ATTR_LEFT_MARGIN.
	Textattribute_LeftMargin Textattribute = 1
	// Textattribute_RightMargin is a representation of the C type ATK_TEXT_ATTR_RIGHT_MARGIN.
	Textattribute_RightMargin Textattribute = 2
	// Textattribute_Indent is a representation of the C type ATK_TEXT_ATTR_INDENT.
	Textattribute_Indent Textattribute = 3
	// Textattribute_Invisible is a representation of the C type ATK_TEXT_ATTR_INVISIBLE.
	Textattribute_Invisible Textattribute = 4
	// Textattribute_Editable is a representation of the C type ATK_TEXT_ATTR_EDITABLE.
	Textattribute_Editable Textattribute = 5
	// Textattribute_PixelsAboveLines is a representation of the C type ATK_TEXT_ATTR_PIXELS_ABOVE_LINES.
	Textattribute_PixelsAboveLines Textattribute = 6
	// Textattribute_PixelsBelowLines is a representation of the C type ATK_TEXT_ATTR_PIXELS_BELOW_LINES.
	Textattribute_PixelsBelowLines Textattribute = 7
	// Textattribute_PixelsInsideWrap is a representation of the C type ATK_TEXT_ATTR_PIXELS_INSIDE_WRAP.
	Textattribute_PixelsInsideWrap Textattribute = 8
	// Textattribute_BgFullHeight is a representation of the C type ATK_TEXT_ATTR_BG_FULL_HEIGHT.
	Textattribute_BgFullHeight Textattribute = 9
	// Textattribute_Rise is a representation of the C type ATK_TEXT_ATTR_RISE.
	Textattribute_Rise Textattribute = 10
	// Textattribute_Underline is a representation of the C type ATK_TEXT_ATTR_UNDERLINE.
	Textattribute_Underline Textattribute = 11
	// Textattribute_Strikethrough is a representation of the C type ATK_TEXT_ATTR_STRIKETHROUGH.
	Textattribute_Strikethrough Textattribute = 12
	// Textattribute_Size is a representation of the C type ATK_TEXT_ATTR_SIZE.
	Textattribute_Size Textattribute = 13
	// Textattribute_Scale is a representation of the C type ATK_TEXT_ATTR_SCALE.
	Textattribute_Scale Textattribute = 14
	// Textattribute_Weight is a representation of the C type ATK_TEXT_ATTR_WEIGHT.
	Textattribute_Weight Textattribute = 15
	// Textattribute_Language is a representation of the C type ATK_TEXT_ATTR_LANGUAGE.
	Textattribute_Language Textattribute = 16
	// Textattribute_FamilyName is a representation of the C type ATK_TEXT_ATTR_FAMILY_NAME.
	Textattribute_FamilyName Textattribute = 17
	// Textattribute_BgColor is a representation of the C type ATK_TEXT_ATTR_BG_COLOR.
	Textattribute_BgColor Textattribute = 18
	// Textattribute_FgColor is a representation of the C type ATK_TEXT_ATTR_FG_COLOR.
	Textattribute_FgColor Textattribute = 19
	// Textattribute_BgStipple is a representation of the C type ATK_TEXT_ATTR_BG_STIPPLE.
	Textattribute_BgStipple Textattribute = 20
	// Textattribute_FgStipple is a representation of the C type ATK_TEXT_ATTR_FG_STIPPLE.
	Textattribute_FgStipple Textattribute = 21
	// Textattribute_WrapMode is a representation of the C type ATK_TEXT_ATTR_WRAP_MODE.
	Textattribute_WrapMode Textattribute = 22
	// Textattribute_Direction is a representation of the C type ATK_TEXT_ATTR_DIRECTION.
	Textattribute_Direction Textattribute = 23
	// Textattribute_Justification is a representation of the C type ATK_TEXT_ATTR_JUSTIFICATION.
	Textattribute_Justification Textattribute = 24
	// Textattribute_Stretch is a representation of the C type ATK_TEXT_ATTR_STRETCH.
	Textattribute_Stretch Textattribute = 25
	// Textattribute_Variant is a representation of the C type ATK_TEXT_ATTR_VARIANT.
	Textattribute_Variant Textattribute = 26
	// Textattribute_Style is a representation of the C type ATK_TEXT_ATTR_STYLE.
	Textattribute_Style Textattribute = 27
	// Textattribute_LastDefined is a representation of the C type ATK_TEXT_ATTR_LAST_DEFINED.
	Textattribute_LastDefined Textattribute = 28
)

// Textboundary is a representation of the C type AtkTextBoundary.
type Textboundary int

const (
	// Textboundary_Char is a representation of the C type ATK_TEXT_BOUNDARY_CHAR.
	Textboundary_Char Textboundary = 0
	// Textboundary_WordStart is a representation of the C type ATK_TEXT_BOUNDARY_WORD_START.
	Textboundary_WordStart Textboundary = 1
	// Textboundary_WordEnd is a representation of the C type ATK_TEXT_BOUNDARY_WORD_END.
	Textboundary_WordEnd Textboundary = 2
	// Textboundary_SentenceStart is a representation of the C type ATK_TEXT_BOUNDARY_SENTENCE_START.
	Textboundary_SentenceStart Textboundary = 3
	// Textboundary_SentenceEnd is a representation of the C type ATK_TEXT_BOUNDARY_SENTENCE_END.
	Textboundary_SentenceEnd Textboundary = 4
	// Textboundary_LineStart is a representation of the C type ATK_TEXT_BOUNDARY_LINE_START.
	Textboundary_LineStart Textboundary = 5
	// Textboundary_LineEnd is a representation of the C type ATK_TEXT_BOUNDARY_LINE_END.
	Textboundary_LineEnd Textboundary = 6
)

// Textcliptype is a representation of the C type AtkTextClipType.
type Textcliptype int

const (
	// Textcliptype_None is a representation of the C type ATK_TEXT_CLIP_NONE.
	Textcliptype_None Textcliptype = 0
	// Textcliptype_Min is a representation of the C type ATK_TEXT_CLIP_MIN.
	Textcliptype_Min Textcliptype = 1
	// Textcliptype_Max is a representation of the C type ATK_TEXT_CLIP_MAX.
	Textcliptype_Max Textcliptype = 2
	// Textcliptype_Both is a representation of the C type ATK_TEXT_CLIP_BOTH.
	Textcliptype_Both Textcliptype = 3
)

// Textgranularity is a representation of the C type AtkTextGranularity.
type Textgranularity int

const (
	// Textgranularity_Char is a representation of the C type ATK_TEXT_GRANULARITY_CHAR.
	Textgranularity_Char Textgranularity = 0
	// Textgranularity_Word is a representation of the C type ATK_TEXT_GRANULARITY_WORD.
	Textgranularity_Word Textgranularity = 1
	// Textgranularity_Sentence is a representation of the C type ATK_TEXT_GRANULARITY_SENTENCE.
	Textgranularity_Sentence Textgranularity = 2
	// Textgranularity_Line is a representation of the C type ATK_TEXT_GRANULARITY_LINE.
	Textgranularity_Line Textgranularity = 3
	// Textgranularity_Paragraph is a representation of the C type ATK_TEXT_GRANULARITY_PARAGRAPH.
	Textgranularity_Paragraph Textgranularity = 4
)

// Valuetype is a representation of the C type AtkValueType.
type Valuetype int

const (
	// Valuetype_VeryWeak is a representation of the C type ATK_VALUE_VERY_WEAK.
	Valuetype_VeryWeak Valuetype = 0
	// Valuetype_Weak is a representation of the C type ATK_VALUE_WEAK.
	Valuetype_Weak Valuetype = 1
	// Valuetype_Acceptable is a representation of the C type ATK_VALUE_ACCEPTABLE.
	Valuetype_Acceptable Valuetype = 2
	// Valuetype_Strong is a representation of the C type ATK_VALUE_STRONG.
	Valuetype_Strong Valuetype = 3
	// Valuetype_VeryStrong is a representation of the C type ATK_VALUE_VERY_STRONG.
	Valuetype_VeryStrong Valuetype = 4
	// Valuetype_VeryLow is a representation of the C type ATK_VALUE_VERY_LOW.
	Valuetype_VeryLow Valuetype = 5
	// Valuetype_Low is a representation of the C type ATK_VALUE_LOW.
	Valuetype_Low Valuetype = 6
	// Valuetype_Medium is a representation of the C type ATK_VALUE_MEDIUM.
	Valuetype_Medium Valuetype = 7
	// Valuetype_High is a representation of the C type ATK_VALUE_HIGH.
	Valuetype_High Valuetype = 8
	// Valuetype_VeryHigh is a representation of the C type ATK_VALUE_VERY_HIGH.
	Valuetype_VeryHigh Valuetype = 9
	// Valuetype_VeryBad is a representation of the C type ATK_VALUE_VERY_BAD.
	Valuetype_VeryBad Valuetype = 10
	// Valuetype_Bad is a representation of the C type ATK_VALUE_BAD.
	Valuetype_Bad Valuetype = 11
	// Valuetype_Good is a representation of the C type ATK_VALUE_GOOD.
	Valuetype_Good Valuetype = 12
	// Valuetype_VeryGood is a representation of the C type ATK_VALUE_VERY_GOOD.
	Valuetype_VeryGood Valuetype = 13
	// Valuetype_Best is a representation of the C type ATK_VALUE_BEST.
	Valuetype_Best Valuetype = 14
	// Valuetype_LastDefined is a representation of the C type ATK_VALUE_LAST_DEFINED.
	Valuetype_LastDefined Valuetype = 15
)
