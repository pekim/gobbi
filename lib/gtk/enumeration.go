// Code generated - DO NOT EDIT.

package gtk

// Align is a representation of the C type GtkAlign.
type Align int

const (
	// Align_Fill is a representation of the C type GTK_ALIGN_FILL.
	Align_Fill Align = 0
	// Align_Start is a representation of the C type GTK_ALIGN_START.
	Align_Start Align = 1
	// Align_End is a representation of the C type GTK_ALIGN_END.
	Align_End Align = 2
	// Align_Center is a representation of the C type GTK_ALIGN_CENTER.
	Align_Center Align = 3
	// Align_Baseline is a representation of the C type GTK_ALIGN_BASELINE.
	Align_Baseline Align = 4
)

// Arrowplacement is a representation of the C type GtkArrowPlacement.
type Arrowplacement int

const (
	// Arrowplacement_Both is a representation of the C type GTK_ARROWS_BOTH.
	Arrowplacement_Both Arrowplacement = 0
	// Arrowplacement_Start is a representation of the C type GTK_ARROWS_START.
	Arrowplacement_Start Arrowplacement = 1
	// Arrowplacement_End is a representation of the C type GTK_ARROWS_END.
	Arrowplacement_End Arrowplacement = 2
)

// Arrowtype is a representation of the C type GtkArrowType.
type Arrowtype int

const (
	// Arrowtype_Up is a representation of the C type GTK_ARROW_UP.
	Arrowtype_Up Arrowtype = 0
	// Arrowtype_Down is a representation of the C type GTK_ARROW_DOWN.
	Arrowtype_Down Arrowtype = 1
	// Arrowtype_Left is a representation of the C type GTK_ARROW_LEFT.
	Arrowtype_Left Arrowtype = 2
	// Arrowtype_Right is a representation of the C type GTK_ARROW_RIGHT.
	Arrowtype_Right Arrowtype = 3
	// Arrowtype_None is a representation of the C type GTK_ARROW_NONE.
	Arrowtype_None Arrowtype = 4
)

// Assistantpagetype is a representation of the C type GtkAssistantPageType.
type Assistantpagetype int

const (
	// Assistantpagetype_Content is a representation of the C type GTK_ASSISTANT_PAGE_CONTENT.
	Assistantpagetype_Content Assistantpagetype = 0
	// Assistantpagetype_Intro is a representation of the C type GTK_ASSISTANT_PAGE_INTRO.
	Assistantpagetype_Intro Assistantpagetype = 1
	// Assistantpagetype_Confirm is a representation of the C type GTK_ASSISTANT_PAGE_CONFIRM.
	Assistantpagetype_Confirm Assistantpagetype = 2
	// Assistantpagetype_Summary is a representation of the C type GTK_ASSISTANT_PAGE_SUMMARY.
	Assistantpagetype_Summary Assistantpagetype = 3
	// Assistantpagetype_Progress is a representation of the C type GTK_ASSISTANT_PAGE_PROGRESS.
	Assistantpagetype_Progress Assistantpagetype = 4
	// Assistantpagetype_Custom is a representation of the C type GTK_ASSISTANT_PAGE_CUSTOM.
	Assistantpagetype_Custom Assistantpagetype = 5
)

// Baselineposition is a representation of the C type GtkBaselinePosition.
//
// since 3.10
type Baselineposition int

const (
	// Baselineposition_Top is a representation of the C type GTK_BASELINE_POSITION_TOP.
	Baselineposition_Top Baselineposition = 0
	// Baselineposition_Center is a representation of the C type GTK_BASELINE_POSITION_CENTER.
	Baselineposition_Center Baselineposition = 1
	// Baselineposition_Bottom is a representation of the C type GTK_BASELINE_POSITION_BOTTOM.
	Baselineposition_Bottom Baselineposition = 2
)

// Borderstyle is a representation of the C type GtkBorderStyle.
type Borderstyle int

const (
	// Borderstyle_None is a representation of the C type GTK_BORDER_STYLE_NONE.
	Borderstyle_None Borderstyle = 0
	// Borderstyle_Solid is a representation of the C type GTK_BORDER_STYLE_SOLID.
	Borderstyle_Solid Borderstyle = 1
	// Borderstyle_Inset is a representation of the C type GTK_BORDER_STYLE_INSET.
	Borderstyle_Inset Borderstyle = 2
	// Borderstyle_Outset is a representation of the C type GTK_BORDER_STYLE_OUTSET.
	Borderstyle_Outset Borderstyle = 3
	// Borderstyle_Hidden is a representation of the C type GTK_BORDER_STYLE_HIDDEN.
	Borderstyle_Hidden Borderstyle = 4
	// Borderstyle_Dotted is a representation of the C type GTK_BORDER_STYLE_DOTTED.
	Borderstyle_Dotted Borderstyle = 5
	// Borderstyle_Dashed is a representation of the C type GTK_BORDER_STYLE_DASHED.
	Borderstyle_Dashed Borderstyle = 6
	// Borderstyle_Double is a representation of the C type GTK_BORDER_STYLE_DOUBLE.
	Borderstyle_Double Borderstyle = 7
	// Borderstyle_Groove is a representation of the C type GTK_BORDER_STYLE_GROOVE.
	Borderstyle_Groove Borderstyle = 8
	// Borderstyle_Ridge is a representation of the C type GTK_BORDER_STYLE_RIDGE.
	Borderstyle_Ridge Borderstyle = 9
)

// Buildererror is a representation of the C type GtkBuilderError.
type Buildererror int

const (
	// Buildererror_InvalidTypeFunction is a representation of the C type GTK_BUILDER_ERROR_INVALID_TYPE_FUNCTION.
	Buildererror_InvalidTypeFunction Buildererror = 0
	// Buildererror_UnhandledTag is a representation of the C type GTK_BUILDER_ERROR_UNHANDLED_TAG.
	Buildererror_UnhandledTag Buildererror = 1
	// Buildererror_MissingAttribute is a representation of the C type GTK_BUILDER_ERROR_MISSING_ATTRIBUTE.
	Buildererror_MissingAttribute Buildererror = 2
	// Buildererror_InvalidAttribute is a representation of the C type GTK_BUILDER_ERROR_INVALID_ATTRIBUTE.
	Buildererror_InvalidAttribute Buildererror = 3
	// Buildererror_InvalidTag is a representation of the C type GTK_BUILDER_ERROR_INVALID_TAG.
	Buildererror_InvalidTag Buildererror = 4
	// Buildererror_MissingPropertyValue is a representation of the C type GTK_BUILDER_ERROR_MISSING_PROPERTY_VALUE.
	Buildererror_MissingPropertyValue Buildererror = 5
	// Buildererror_InvalidValue is a representation of the C type GTK_BUILDER_ERROR_INVALID_VALUE.
	Buildererror_InvalidValue Buildererror = 6
	// Buildererror_VersionMismatch is a representation of the C type GTK_BUILDER_ERROR_VERSION_MISMATCH.
	Buildererror_VersionMismatch Buildererror = 7
	// Buildererror_DuplicateId is a representation of the C type GTK_BUILDER_ERROR_DUPLICATE_ID.
	Buildererror_DuplicateId Buildererror = 8
	// Buildererror_ObjectTypeRefused is a representation of the C type GTK_BUILDER_ERROR_OBJECT_TYPE_REFUSED.
	Buildererror_ObjectTypeRefused Buildererror = 9
	// Buildererror_TemplateMismatch is a representation of the C type GTK_BUILDER_ERROR_TEMPLATE_MISMATCH.
	Buildererror_TemplateMismatch Buildererror = 10
	// Buildererror_InvalidProperty is a representation of the C type GTK_BUILDER_ERROR_INVALID_PROPERTY.
	Buildererror_InvalidProperty Buildererror = 11
	// Buildererror_InvalidSignal is a representation of the C type GTK_BUILDER_ERROR_INVALID_SIGNAL.
	Buildererror_InvalidSignal Buildererror = 12
	// Buildererror_InvalidId is a representation of the C type GTK_BUILDER_ERROR_INVALID_ID.
	Buildererror_InvalidId Buildererror = 13
)

// Buttonboxstyle is a representation of the C type GtkButtonBoxStyle.
type Buttonboxstyle int

const (
	// Buttonboxstyle_Spread is a representation of the C type GTK_BUTTONBOX_SPREAD.
	Buttonboxstyle_Spread Buttonboxstyle = 1
	// Buttonboxstyle_Edge is a representation of the C type GTK_BUTTONBOX_EDGE.
	Buttonboxstyle_Edge Buttonboxstyle = 2
	// Buttonboxstyle_Start is a representation of the C type GTK_BUTTONBOX_START.
	Buttonboxstyle_Start Buttonboxstyle = 3
	// Buttonboxstyle_End is a representation of the C type GTK_BUTTONBOX_END.
	Buttonboxstyle_End Buttonboxstyle = 4
	// Buttonboxstyle_Center is a representation of the C type GTK_BUTTONBOX_CENTER.
	Buttonboxstyle_Center Buttonboxstyle = 5
	// Buttonboxstyle_Expand is a representation of the C type GTK_BUTTONBOX_EXPAND.
	Buttonboxstyle_Expand Buttonboxstyle = 6
)

// Buttonrole is a representation of the C type GtkButtonRole.
type Buttonrole int

const (
	// Buttonrole_Normal is a representation of the C type GTK_BUTTON_ROLE_NORMAL.
	Buttonrole_Normal Buttonrole = 0
	// Buttonrole_Check is a representation of the C type GTK_BUTTON_ROLE_CHECK.
	Buttonrole_Check Buttonrole = 1
	// Buttonrole_Radio is a representation of the C type GTK_BUTTON_ROLE_RADIO.
	Buttonrole_Radio Buttonrole = 2
)

// Buttonstype is a representation of the C type GtkButtonsType.
type Buttonstype int

const (
	// Buttonstype_None is a representation of the C type GTK_BUTTONS_NONE.
	Buttonstype_None Buttonstype = 0
	// Buttonstype_Ok is a representation of the C type GTK_BUTTONS_OK.
	Buttonstype_Ok Buttonstype = 1
	// Buttonstype_Close is a representation of the C type GTK_BUTTONS_CLOSE.
	Buttonstype_Close Buttonstype = 2
	// Buttonstype_Cancel is a representation of the C type GTK_BUTTONS_CANCEL.
	Buttonstype_Cancel Buttonstype = 3
	// Buttonstype_YesNo is a representation of the C type GTK_BUTTONS_YES_NO.
	Buttonstype_YesNo Buttonstype = 4
	// Buttonstype_OkCancel is a representation of the C type GTK_BUTTONS_OK_CANCEL.
	Buttonstype_OkCancel Buttonstype = 5
)

// Cellrendereraccelmode is a representation of the C type GtkCellRendererAccelMode.
type Cellrendereraccelmode int

const (
	// Cellrendereraccelmode_Gtk is a representation of the C type GTK_CELL_RENDERER_ACCEL_MODE_GTK.
	Cellrendereraccelmode_Gtk Cellrendereraccelmode = 0
	// Cellrendereraccelmode_Other is a representation of the C type GTK_CELL_RENDERER_ACCEL_MODE_OTHER.
	Cellrendereraccelmode_Other Cellrendereraccelmode = 1
	// Cellrendereraccelmode_ModifierTap is a representation of the C type GTK_CELL_RENDERER_ACCEL_MODE_MODIFIER_TAP.
	Cellrendereraccelmode_ModifierTap Cellrendereraccelmode = 2
)

// Cellrenderermode is a representation of the C type GtkCellRendererMode.
type Cellrenderermode int

const (
	// Cellrenderermode_Inert is a representation of the C type GTK_CELL_RENDERER_MODE_INERT.
	Cellrenderermode_Inert Cellrenderermode = 0
	// Cellrenderermode_Activatable is a representation of the C type GTK_CELL_RENDERER_MODE_ACTIVATABLE.
	Cellrenderermode_Activatable Cellrenderermode = 1
	// Cellrenderermode_Editable is a representation of the C type GTK_CELL_RENDERER_MODE_EDITABLE.
	Cellrenderermode_Editable Cellrenderermode = 2
)

// Cornertype is a representation of the C type GtkCornerType.
type Cornertype int

const (
	// Cornertype_TopLeft is a representation of the C type GTK_CORNER_TOP_LEFT.
	Cornertype_TopLeft Cornertype = 0
	// Cornertype_BottomLeft is a representation of the C type GTK_CORNER_BOTTOM_LEFT.
	Cornertype_BottomLeft Cornertype = 1
	// Cornertype_TopRight is a representation of the C type GTK_CORNER_TOP_RIGHT.
	Cornertype_TopRight Cornertype = 2
	// Cornertype_BottomRight is a representation of the C type GTK_CORNER_BOTTOM_RIGHT.
	Cornertype_BottomRight Cornertype = 3
)

// Cssprovidererror is a representation of the C type GtkCssProviderError.
type Cssprovidererror int

const (
	// Cssprovidererror_Failed is a representation of the C type GTK_CSS_PROVIDER_ERROR_FAILED.
	Cssprovidererror_Failed Cssprovidererror = 0
	// Cssprovidererror_Syntax is a representation of the C type GTK_CSS_PROVIDER_ERROR_SYNTAX.
	Cssprovidererror_Syntax Cssprovidererror = 1
	// Cssprovidererror_Import is a representation of the C type GTK_CSS_PROVIDER_ERROR_IMPORT.
	Cssprovidererror_Import Cssprovidererror = 2
	// Cssprovidererror_Name is a representation of the C type GTK_CSS_PROVIDER_ERROR_NAME.
	Cssprovidererror_Name Cssprovidererror = 3
	// Cssprovidererror_Deprecated is a representation of the C type GTK_CSS_PROVIDER_ERROR_DEPRECATED.
	Cssprovidererror_Deprecated Cssprovidererror = 4
	// Cssprovidererror_UnknownValue is a representation of the C type GTK_CSS_PROVIDER_ERROR_UNKNOWN_VALUE.
	Cssprovidererror_UnknownValue Cssprovidererror = 5
)

// Csssectiontype is a representation of the C type GtkCssSectionType.
//
// since 3.2
type Csssectiontype int

const (
	// Csssectiontype_Document is a representation of the C type GTK_CSS_SECTION_DOCUMENT.
	Csssectiontype_Document Csssectiontype = 0
	// Csssectiontype_Import is a representation of the C type GTK_CSS_SECTION_IMPORT.
	Csssectiontype_Import Csssectiontype = 1
	// Csssectiontype_ColorDefinition is a representation of the C type GTK_CSS_SECTION_COLOR_DEFINITION.
	Csssectiontype_ColorDefinition Csssectiontype = 2
	// Csssectiontype_BindingSet is a representation of the C type GTK_CSS_SECTION_BINDING_SET.
	Csssectiontype_BindingSet Csssectiontype = 3
	// Csssectiontype_Ruleset is a representation of the C type GTK_CSS_SECTION_RULESET.
	Csssectiontype_Ruleset Csssectiontype = 4
	// Csssectiontype_Selector is a representation of the C type GTK_CSS_SECTION_SELECTOR.
	Csssectiontype_Selector Csssectiontype = 5
	// Csssectiontype_Declaration is a representation of the C type GTK_CSS_SECTION_DECLARATION.
	Csssectiontype_Declaration Csssectiontype = 6
	// Csssectiontype_Value is a representation of the C type GTK_CSS_SECTION_VALUE.
	Csssectiontype_Value Csssectiontype = 7
	// Csssectiontype_Keyframes is a representation of the C type GTK_CSS_SECTION_KEYFRAMES.
	Csssectiontype_Keyframes Csssectiontype = 8
)

// Deletetype is a representation of the C type GtkDeleteType.
type Deletetype int

const (
	// Deletetype_Chars is a representation of the C type GTK_DELETE_CHARS.
	Deletetype_Chars Deletetype = 0
	// Deletetype_WordEnds is a representation of the C type GTK_DELETE_WORD_ENDS.
	Deletetype_WordEnds Deletetype = 1
	// Deletetype_Words is a representation of the C type GTK_DELETE_WORDS.
	Deletetype_Words Deletetype = 2
	// Deletetype_DisplayLines is a representation of the C type GTK_DELETE_DISPLAY_LINES.
	Deletetype_DisplayLines Deletetype = 3
	// Deletetype_DisplayLineEnds is a representation of the C type GTK_DELETE_DISPLAY_LINE_ENDS.
	Deletetype_DisplayLineEnds Deletetype = 4
	// Deletetype_ParagraphEnds is a representation of the C type GTK_DELETE_PARAGRAPH_ENDS.
	Deletetype_ParagraphEnds Deletetype = 5
	// Deletetype_Paragraphs is a representation of the C type GTK_DELETE_PARAGRAPHS.
	Deletetype_Paragraphs Deletetype = 6
	// Deletetype_Whitespace is a representation of the C type GTK_DELETE_WHITESPACE.
	Deletetype_Whitespace Deletetype = 7
)

// Directiontype is a representation of the C type GtkDirectionType.
type Directiontype int

const (
	// Directiontype_TabForward is a representation of the C type GTK_DIR_TAB_FORWARD.
	Directiontype_TabForward Directiontype = 0
	// Directiontype_TabBackward is a representation of the C type GTK_DIR_TAB_BACKWARD.
	Directiontype_TabBackward Directiontype = 1
	// Directiontype_Up is a representation of the C type GTK_DIR_UP.
	Directiontype_Up Directiontype = 2
	// Directiontype_Down is a representation of the C type GTK_DIR_DOWN.
	Directiontype_Down Directiontype = 3
	// Directiontype_Left is a representation of the C type GTK_DIR_LEFT.
	Directiontype_Left Directiontype = 4
	// Directiontype_Right is a representation of the C type GTK_DIR_RIGHT.
	Directiontype_Right Directiontype = 5
)

// Dragresult is a representation of the C type GtkDragResult.
type Dragresult int

const (
	// Dragresult_Success is a representation of the C type GTK_DRAG_RESULT_SUCCESS.
	Dragresult_Success Dragresult = 0
	// Dragresult_NoTarget is a representation of the C type GTK_DRAG_RESULT_NO_TARGET.
	Dragresult_NoTarget Dragresult = 1
	// Dragresult_UserCancelled is a representation of the C type GTK_DRAG_RESULT_USER_CANCELLED.
	Dragresult_UserCancelled Dragresult = 2
	// Dragresult_TimeoutExpired is a representation of the C type GTK_DRAG_RESULT_TIMEOUT_EXPIRED.
	Dragresult_TimeoutExpired Dragresult = 3
	// Dragresult_GrabBroken is a representation of the C type GTK_DRAG_RESULT_GRAB_BROKEN.
	Dragresult_GrabBroken Dragresult = 4
	// Dragresult_Error is a representation of the C type GTK_DRAG_RESULT_ERROR.
	Dragresult_Error Dragresult = 5
)

// Entryiconposition is a representation of the C type GtkEntryIconPosition.
//
// since 2.16
type Entryiconposition int

const (
	// Entryiconposition_Primary is a representation of the C type GTK_ENTRY_ICON_PRIMARY.
	Entryiconposition_Primary Entryiconposition = 0
	// Entryiconposition_Secondary is a representation of the C type GTK_ENTRY_ICON_SECONDARY.
	Entryiconposition_Secondary Entryiconposition = 1
)

// Eventsequencestate is a representation of the C type GtkEventSequenceState.
//
// since 3.14
type Eventsequencestate int

const (
	// Eventsequencestate_None is a representation of the C type GTK_EVENT_SEQUENCE_NONE.
	Eventsequencestate_None Eventsequencestate = 0
	// Eventsequencestate_Claimed is a representation of the C type GTK_EVENT_SEQUENCE_CLAIMED.
	Eventsequencestate_Claimed Eventsequencestate = 1
	// Eventsequencestate_Denied is a representation of the C type GTK_EVENT_SEQUENCE_DENIED.
	Eventsequencestate_Denied Eventsequencestate = 2
)

// Expanderstyle is a representation of the C type GtkExpanderStyle.
type Expanderstyle int

const (
	// Expanderstyle_Collapsed is a representation of the C type GTK_EXPANDER_COLLAPSED.
	Expanderstyle_Collapsed Expanderstyle = 0
	// Expanderstyle_SemiCollapsed is a representation of the C type GTK_EXPANDER_SEMI_COLLAPSED.
	Expanderstyle_SemiCollapsed Expanderstyle = 1
	// Expanderstyle_SemiExpanded is a representation of the C type GTK_EXPANDER_SEMI_EXPANDED.
	Expanderstyle_SemiExpanded Expanderstyle = 2
	// Expanderstyle_Expanded is a representation of the C type GTK_EXPANDER_EXPANDED.
	Expanderstyle_Expanded Expanderstyle = 3
)

// Filechooseraction is a representation of the C type GtkFileChooserAction.
type Filechooseraction int

const (
	// Filechooseraction_Open is a representation of the C type GTK_FILE_CHOOSER_ACTION_OPEN.
	Filechooseraction_Open Filechooseraction = 0
	// Filechooseraction_Save is a representation of the C type GTK_FILE_CHOOSER_ACTION_SAVE.
	Filechooseraction_Save Filechooseraction = 1
	// Filechooseraction_SelectFolder is a representation of the C type GTK_FILE_CHOOSER_ACTION_SELECT_FOLDER.
	Filechooseraction_SelectFolder Filechooseraction = 2
	// Filechooseraction_CreateFolder is a representation of the C type GTK_FILE_CHOOSER_ACTION_CREATE_FOLDER.
	Filechooseraction_CreateFolder Filechooseraction = 3
)

// Filechooserconfirmation is a representation of the C type GtkFileChooserConfirmation.
//
// since 2.8
type Filechooserconfirmation int

const (
	// Filechooserconfirmation_Confirm is a representation of the C type GTK_FILE_CHOOSER_CONFIRMATION_CONFIRM.
	Filechooserconfirmation_Confirm Filechooserconfirmation = 0
	// Filechooserconfirmation_AcceptFilename is a representation of the C type GTK_FILE_CHOOSER_CONFIRMATION_ACCEPT_FILENAME.
	Filechooserconfirmation_AcceptFilename Filechooserconfirmation = 1
	// Filechooserconfirmation_SelectAgain is a representation of the C type GTK_FILE_CHOOSER_CONFIRMATION_SELECT_AGAIN.
	Filechooserconfirmation_SelectAgain Filechooserconfirmation = 2
)

// Filechoosererror is a representation of the C type GtkFileChooserError.
type Filechoosererror int

const (
	// Filechoosererror_Nonexistent is a representation of the C type GTK_FILE_CHOOSER_ERROR_NONEXISTENT.
	Filechoosererror_Nonexistent Filechoosererror = 0
	// Filechoosererror_BadFilename is a representation of the C type GTK_FILE_CHOOSER_ERROR_BAD_FILENAME.
	Filechoosererror_BadFilename Filechoosererror = 1
	// Filechoosererror_AlreadyExists is a representation of the C type GTK_FILE_CHOOSER_ERROR_ALREADY_EXISTS.
	Filechoosererror_AlreadyExists Filechoosererror = 2
	// Filechoosererror_IncompleteHostname is a representation of the C type GTK_FILE_CHOOSER_ERROR_INCOMPLETE_HOSTNAME.
	Filechoosererror_IncompleteHostname Filechoosererror = 3
)

// Impreeditstyle is a representation of the C type GtkIMPreeditStyle.
type Impreeditstyle int

const (
	// Impreeditstyle_Nothing is a representation of the C type GTK_IM_PREEDIT_NOTHING.
	Impreeditstyle_Nothing Impreeditstyle = 0
	// Impreeditstyle_Callback is a representation of the C type GTK_IM_PREEDIT_CALLBACK.
	Impreeditstyle_Callback Impreeditstyle = 1
	// Impreeditstyle_None is a representation of the C type GTK_IM_PREEDIT_NONE.
	Impreeditstyle_None Impreeditstyle = 2
)

// Imstatusstyle is a representation of the C type GtkIMStatusStyle.
type Imstatusstyle int

const (
	// Imstatusstyle_Nothing is a representation of the C type GTK_IM_STATUS_NOTHING.
	Imstatusstyle_Nothing Imstatusstyle = 0
	// Imstatusstyle_Callback is a representation of the C type GTK_IM_STATUS_CALLBACK.
	Imstatusstyle_Callback Imstatusstyle = 1
	// Imstatusstyle_None is a representation of the C type GTK_IM_STATUS_NONE.
	Imstatusstyle_None Imstatusstyle = 2
)

// Iconsize is a representation of the C type GtkIconSize.
type Iconsize int

const (
	// Iconsize_Invalid is a representation of the C type GTK_ICON_SIZE_INVALID.
	Iconsize_Invalid Iconsize = 0
	// Iconsize_Menu is a representation of the C type GTK_ICON_SIZE_MENU.
	Iconsize_Menu Iconsize = 1
	// Iconsize_SmallToolbar is a representation of the C type GTK_ICON_SIZE_SMALL_TOOLBAR.
	Iconsize_SmallToolbar Iconsize = 2
	// Iconsize_LargeToolbar is a representation of the C type GTK_ICON_SIZE_LARGE_TOOLBAR.
	Iconsize_LargeToolbar Iconsize = 3
	// Iconsize_Button is a representation of the C type GTK_ICON_SIZE_BUTTON.
	Iconsize_Button Iconsize = 4
	// Iconsize_Dnd is a representation of the C type GTK_ICON_SIZE_DND.
	Iconsize_Dnd Iconsize = 5
	// Iconsize_Dialog is a representation of the C type GTK_ICON_SIZE_DIALOG.
	Iconsize_Dialog Iconsize = 6
)

// Iconthemeerror is a representation of the C type GtkIconThemeError.
type Iconthemeerror int

const (
	// Iconthemeerror_NotFound is a representation of the C type GTK_ICON_THEME_NOT_FOUND.
	Iconthemeerror_NotFound Iconthemeerror = 0
	// Iconthemeerror_Failed is a representation of the C type GTK_ICON_THEME_FAILED.
	Iconthemeerror_Failed Iconthemeerror = 1
)

// Iconviewdropposition is a representation of the C type GtkIconViewDropPosition.
type Iconviewdropposition int

const (
	// Iconviewdropposition_NoDrop is a representation of the C type GTK_ICON_VIEW_NO_DROP.
	Iconviewdropposition_NoDrop Iconviewdropposition = 0
	// Iconviewdropposition_DropInto is a representation of the C type GTK_ICON_VIEW_DROP_INTO.
	Iconviewdropposition_DropInto Iconviewdropposition = 1
	// Iconviewdropposition_DropLeft is a representation of the C type GTK_ICON_VIEW_DROP_LEFT.
	Iconviewdropposition_DropLeft Iconviewdropposition = 2
	// Iconviewdropposition_DropRight is a representation of the C type GTK_ICON_VIEW_DROP_RIGHT.
	Iconviewdropposition_DropRight Iconviewdropposition = 3
	// Iconviewdropposition_DropAbove is a representation of the C type GTK_ICON_VIEW_DROP_ABOVE.
	Iconviewdropposition_DropAbove Iconviewdropposition = 4
	// Iconviewdropposition_DropBelow is a representation of the C type GTK_ICON_VIEW_DROP_BELOW.
	Iconviewdropposition_DropBelow Iconviewdropposition = 5
)

// Imagetype is a representation of the C type GtkImageType.
type Imagetype int

const (
	// Imagetype_Empty is a representation of the C type GTK_IMAGE_EMPTY.
	Imagetype_Empty Imagetype = 0
	// Imagetype_Pixbuf is a representation of the C type GTK_IMAGE_PIXBUF.
	Imagetype_Pixbuf Imagetype = 1
	// Imagetype_Stock is a representation of the C type GTK_IMAGE_STOCK.
	Imagetype_Stock Imagetype = 2
	// Imagetype_IconSet is a representation of the C type GTK_IMAGE_ICON_SET.
	Imagetype_IconSet Imagetype = 3
	// Imagetype_Animation is a representation of the C type GTK_IMAGE_ANIMATION.
	Imagetype_Animation Imagetype = 4
	// Imagetype_IconName is a representation of the C type GTK_IMAGE_ICON_NAME.
	Imagetype_IconName Imagetype = 5
	// Imagetype_Gicon is a representation of the C type GTK_IMAGE_GICON.
	Imagetype_Gicon Imagetype = 6
	// Imagetype_Surface is a representation of the C type GTK_IMAGE_SURFACE.
	Imagetype_Surface Imagetype = 7
)

// Inputpurpose is a representation of the C type GtkInputPurpose.
//
// since 3.6
type Inputpurpose int

const (
	// Inputpurpose_FreeForm is a representation of the C type GTK_INPUT_PURPOSE_FREE_FORM.
	Inputpurpose_FreeForm Inputpurpose = 0
	// Inputpurpose_Alpha is a representation of the C type GTK_INPUT_PURPOSE_ALPHA.
	Inputpurpose_Alpha Inputpurpose = 1
	// Inputpurpose_Digits is a representation of the C type GTK_INPUT_PURPOSE_DIGITS.
	Inputpurpose_Digits Inputpurpose = 2
	// Inputpurpose_Number is a representation of the C type GTK_INPUT_PURPOSE_NUMBER.
	Inputpurpose_Number Inputpurpose = 3
	// Inputpurpose_Phone is a representation of the C type GTK_INPUT_PURPOSE_PHONE.
	Inputpurpose_Phone Inputpurpose = 4
	// Inputpurpose_Url is a representation of the C type GTK_INPUT_PURPOSE_URL.
	Inputpurpose_Url Inputpurpose = 5
	// Inputpurpose_Email is a representation of the C type GTK_INPUT_PURPOSE_EMAIL.
	Inputpurpose_Email Inputpurpose = 6
	// Inputpurpose_Name is a representation of the C type GTK_INPUT_PURPOSE_NAME.
	Inputpurpose_Name Inputpurpose = 7
	// Inputpurpose_Password is a representation of the C type GTK_INPUT_PURPOSE_PASSWORD.
	Inputpurpose_Password Inputpurpose = 8
	// Inputpurpose_Pin is a representation of the C type GTK_INPUT_PURPOSE_PIN.
	Inputpurpose_Pin Inputpurpose = 9
)

// Justification is a representation of the C type GtkJustification.
type Justification int

const (
	// Justification_Left is a representation of the C type GTK_JUSTIFY_LEFT.
	Justification_Left Justification = 0
	// Justification_Right is a representation of the C type GTK_JUSTIFY_RIGHT.
	Justification_Right Justification = 1
	// Justification_Center is a representation of the C type GTK_JUSTIFY_CENTER.
	Justification_Center Justification = 2
	// Justification_Fill is a representation of the C type GTK_JUSTIFY_FILL.
	Justification_Fill Justification = 3
)

// Levelbarmode is a representation of the C type GtkLevelBarMode.
//
// since 3.6
type Levelbarmode int

const (
	// Levelbarmode_Continuous is a representation of the C type GTK_LEVEL_BAR_MODE_CONTINUOUS.
	Levelbarmode_Continuous Levelbarmode = 0
	// Levelbarmode_Discrete is a representation of the C type GTK_LEVEL_BAR_MODE_DISCRETE.
	Levelbarmode_Discrete Levelbarmode = 1
)

// License is a representation of the C type GtkLicense.
//
// since 3.0
type License int

const (
	// License_Unknown is a representation of the C type GTK_LICENSE_UNKNOWN.
	License_Unknown License = 0
	// License_Custom is a representation of the C type GTK_LICENSE_CUSTOM.
	License_Custom License = 1
	// License_Gpl20 is a representation of the C type GTK_LICENSE_GPL_2_0.
	License_Gpl20 License = 2
	// License_Gpl30 is a representation of the C type GTK_LICENSE_GPL_3_0.
	License_Gpl30 License = 3
	// License_Lgpl21 is a representation of the C type GTK_LICENSE_LGPL_2_1.
	License_Lgpl21 License = 4
	// License_Lgpl30 is a representation of the C type GTK_LICENSE_LGPL_3_0.
	License_Lgpl30 License = 5
	// License_Bsd is a representation of the C type GTK_LICENSE_BSD.
	License_Bsd License = 6
	// License_MitX11 is a representation of the C type GTK_LICENSE_MIT_X11.
	License_MitX11 License = 7
	// License_Artistic is a representation of the C type GTK_LICENSE_ARTISTIC.
	License_Artistic License = 8
	// License_Gpl20Only is a representation of the C type GTK_LICENSE_GPL_2_0_ONLY.
	License_Gpl20Only License = 9
	// License_Gpl30Only is a representation of the C type GTK_LICENSE_GPL_3_0_ONLY.
	License_Gpl30Only License = 10
	// License_Lgpl21Only is a representation of the C type GTK_LICENSE_LGPL_2_1_ONLY.
	License_Lgpl21Only License = 11
	// License_Lgpl30Only is a representation of the C type GTK_LICENSE_LGPL_3_0_ONLY.
	License_Lgpl30Only License = 12
	// License_Agpl30 is a representation of the C type GTK_LICENSE_AGPL_3_0.
	License_Agpl30 License = 13
	// License_Agpl30Only is a representation of the C type GTK_LICENSE_AGPL_3_0_ONLY.
	License_Agpl30Only License = 14
)

// Menudirectiontype is a representation of the C type GtkMenuDirectionType.
type Menudirectiontype int

const (
	// Menudirectiontype_Parent is a representation of the C type GTK_MENU_DIR_PARENT.
	Menudirectiontype_Parent Menudirectiontype = 0
	// Menudirectiontype_Child is a representation of the C type GTK_MENU_DIR_CHILD.
	Menudirectiontype_Child Menudirectiontype = 1
	// Menudirectiontype_Next is a representation of the C type GTK_MENU_DIR_NEXT.
	Menudirectiontype_Next Menudirectiontype = 2
	// Menudirectiontype_Prev is a representation of the C type GTK_MENU_DIR_PREV.
	Menudirectiontype_Prev Menudirectiontype = 3
)

// Messagetype is a representation of the C type GtkMessageType.
type Messagetype int

const (
	// Messagetype_Info is a representation of the C type GTK_MESSAGE_INFO.
	Messagetype_Info Messagetype = 0
	// Messagetype_Warning is a representation of the C type GTK_MESSAGE_WARNING.
	Messagetype_Warning Messagetype = 1
	// Messagetype_Question is a representation of the C type GTK_MESSAGE_QUESTION.
	Messagetype_Question Messagetype = 2
	// Messagetype_Error is a representation of the C type GTK_MESSAGE_ERROR.
	Messagetype_Error Messagetype = 3
	// Messagetype_Other is a representation of the C type GTK_MESSAGE_OTHER.
	Messagetype_Other Messagetype = 4
)

// Movementstep is a representation of the C type GtkMovementStep.
type Movementstep int

const (
	// Movementstep_LogicalPositions is a representation of the C type GTK_MOVEMENT_LOGICAL_POSITIONS.
	Movementstep_LogicalPositions Movementstep = 0
	// Movementstep_VisualPositions is a representation of the C type GTK_MOVEMENT_VISUAL_POSITIONS.
	Movementstep_VisualPositions Movementstep = 1
	// Movementstep_Words is a representation of the C type GTK_MOVEMENT_WORDS.
	Movementstep_Words Movementstep = 2
	// Movementstep_DisplayLines is a representation of the C type GTK_MOVEMENT_DISPLAY_LINES.
	Movementstep_DisplayLines Movementstep = 3
	// Movementstep_DisplayLineEnds is a representation of the C type GTK_MOVEMENT_DISPLAY_LINE_ENDS.
	Movementstep_DisplayLineEnds Movementstep = 4
	// Movementstep_Paragraphs is a representation of the C type GTK_MOVEMENT_PARAGRAPHS.
	Movementstep_Paragraphs Movementstep = 5
	// Movementstep_ParagraphEnds is a representation of the C type GTK_MOVEMENT_PARAGRAPH_ENDS.
	Movementstep_ParagraphEnds Movementstep = 6
	// Movementstep_Pages is a representation of the C type GTK_MOVEMENT_PAGES.
	Movementstep_Pages Movementstep = 7
	// Movementstep_BufferEnds is a representation of the C type GTK_MOVEMENT_BUFFER_ENDS.
	Movementstep_BufferEnds Movementstep = 8
	// Movementstep_HorizontalPages is a representation of the C type GTK_MOVEMENT_HORIZONTAL_PAGES.
	Movementstep_HorizontalPages Movementstep = 9
)

// Notebooktab is a representation of the C type GtkNotebookTab.
type Notebooktab int

const (
	// Notebooktab_First is a representation of the C type GTK_NOTEBOOK_TAB_FIRST.
	Notebooktab_First Notebooktab = 0
	// Notebooktab_Last is a representation of the C type GTK_NOTEBOOK_TAB_LAST.
	Notebooktab_Last Notebooktab = 1
)

// Numberuplayout is a representation of the C type GtkNumberUpLayout.
type Numberuplayout int

const (
	// Numberuplayout_Lrtb is a representation of the C type GTK_NUMBER_UP_LAYOUT_LEFT_TO_RIGHT_TOP_TO_BOTTOM.
	Numberuplayout_Lrtb Numberuplayout = 0
	// Numberuplayout_Lrbt is a representation of the C type GTK_NUMBER_UP_LAYOUT_LEFT_TO_RIGHT_BOTTOM_TO_TOP.
	Numberuplayout_Lrbt Numberuplayout = 1
	// Numberuplayout_Rltb is a representation of the C type GTK_NUMBER_UP_LAYOUT_RIGHT_TO_LEFT_TOP_TO_BOTTOM.
	Numberuplayout_Rltb Numberuplayout = 2
	// Numberuplayout_Rlbt is a representation of the C type GTK_NUMBER_UP_LAYOUT_RIGHT_TO_LEFT_BOTTOM_TO_TOP.
	Numberuplayout_Rlbt Numberuplayout = 3
	// Numberuplayout_Tblr is a representation of the C type GTK_NUMBER_UP_LAYOUT_TOP_TO_BOTTOM_LEFT_TO_RIGHT.
	Numberuplayout_Tblr Numberuplayout = 4
	// Numberuplayout_Tbrl is a representation of the C type GTK_NUMBER_UP_LAYOUT_TOP_TO_BOTTOM_RIGHT_TO_LEFT.
	Numberuplayout_Tbrl Numberuplayout = 5
	// Numberuplayout_Btlr is a representation of the C type GTK_NUMBER_UP_LAYOUT_BOTTOM_TO_TOP_LEFT_TO_RIGHT.
	Numberuplayout_Btlr Numberuplayout = 6
	// Numberuplayout_Btrl is a representation of the C type GTK_NUMBER_UP_LAYOUT_BOTTOM_TO_TOP_RIGHT_TO_LEFT.
	Numberuplayout_Btrl Numberuplayout = 7
)

// Orientation is a representation of the C type GtkOrientation.
type Orientation int

const (
	// Orientation_Horizontal is a representation of the C type GTK_ORIENTATION_HORIZONTAL.
	Orientation_Horizontal Orientation = 0
	// Orientation_Vertical is a representation of the C type GTK_ORIENTATION_VERTICAL.
	Orientation_Vertical Orientation = 1
)

// Packdirection is a representation of the C type GtkPackDirection.
type Packdirection int

const (
	// Packdirection_Ltr is a representation of the C type GTK_PACK_DIRECTION_LTR.
	Packdirection_Ltr Packdirection = 0
	// Packdirection_Rtl is a representation of the C type GTK_PACK_DIRECTION_RTL.
	Packdirection_Rtl Packdirection = 1
	// Packdirection_Ttb is a representation of the C type GTK_PACK_DIRECTION_TTB.
	Packdirection_Ttb Packdirection = 2
	// Packdirection_Btt is a representation of the C type GTK_PACK_DIRECTION_BTT.
	Packdirection_Btt Packdirection = 3
)

// Packtype is a representation of the C type GtkPackType.
type Packtype int

const (
	// Packtype_Start is a representation of the C type GTK_PACK_START.
	Packtype_Start Packtype = 0
	// Packtype_End is a representation of the C type GTK_PACK_END.
	Packtype_End Packtype = 1
)

// Padactiontype is a representation of the C type GtkPadActionType.
type Padactiontype int

const (
	// Padactiontype_Button is a representation of the C type GTK_PAD_ACTION_BUTTON.
	Padactiontype_Button Padactiontype = 0
	// Padactiontype_Ring is a representation of the C type GTK_PAD_ACTION_RING.
	Padactiontype_Ring Padactiontype = 1
	// Padactiontype_Strip is a representation of the C type GTK_PAD_ACTION_STRIP.
	Padactiontype_Strip Padactiontype = 2
)

// Pageorientation is a representation of the C type GtkPageOrientation.
type Pageorientation int

const (
	// Pageorientation_Portrait is a representation of the C type GTK_PAGE_ORIENTATION_PORTRAIT.
	Pageorientation_Portrait Pageorientation = 0
	// Pageorientation_Landscape is a representation of the C type GTK_PAGE_ORIENTATION_LANDSCAPE.
	Pageorientation_Landscape Pageorientation = 1
	// Pageorientation_ReversePortrait is a representation of the C type GTK_PAGE_ORIENTATION_REVERSE_PORTRAIT.
	Pageorientation_ReversePortrait Pageorientation = 2
	// Pageorientation_ReverseLandscape is a representation of the C type GTK_PAGE_ORIENTATION_REVERSE_LANDSCAPE.
	Pageorientation_ReverseLandscape Pageorientation = 3
)

// Pageset is a representation of the C type GtkPageSet.
type Pageset int

const (
	// Pageset_All is a representation of the C type GTK_PAGE_SET_ALL.
	Pageset_All Pageset = 0
	// Pageset_Even is a representation of the C type GTK_PAGE_SET_EVEN.
	Pageset_Even Pageset = 1
	// Pageset_Odd is a representation of the C type GTK_PAGE_SET_ODD.
	Pageset_Odd Pageset = 2
)

// Pandirection is a representation of the C type GtkPanDirection.
//
// since 3.14
type Pandirection int

const (
	// Pandirection_Left is a representation of the C type GTK_PAN_DIRECTION_LEFT.
	Pandirection_Left Pandirection = 0
	// Pandirection_Right is a representation of the C type GTK_PAN_DIRECTION_RIGHT.
	Pandirection_Right Pandirection = 1
	// Pandirection_Up is a representation of the C type GTK_PAN_DIRECTION_UP.
	Pandirection_Up Pandirection = 2
	// Pandirection_Down is a representation of the C type GTK_PAN_DIRECTION_DOWN.
	Pandirection_Down Pandirection = 3
)

// Pathprioritytype is a representation of the C type GtkPathPriorityType.
type Pathprioritytype int

const (
	// Pathprioritytype_Lowest is a representation of the C type GTK_PATH_PRIO_LOWEST.
	Pathprioritytype_Lowest Pathprioritytype = 0
	// Pathprioritytype_Gtk is a representation of the C type GTK_PATH_PRIO_GTK.
	Pathprioritytype_Gtk Pathprioritytype = 4
	// Pathprioritytype_Application is a representation of the C type GTK_PATH_PRIO_APPLICATION.
	Pathprioritytype_Application Pathprioritytype = 8
	// Pathprioritytype_Theme is a representation of the C type GTK_PATH_PRIO_THEME.
	Pathprioritytype_Theme Pathprioritytype = 10
	// Pathprioritytype_Rc is a representation of the C type GTK_PATH_PRIO_RC.
	Pathprioritytype_Rc Pathprioritytype = 12
	// Pathprioritytype_Highest is a representation of the C type GTK_PATH_PRIO_HIGHEST.
	Pathprioritytype_Highest Pathprioritytype = 15
)

// Pathtype is a representation of the C type GtkPathType.
type Pathtype int

const (
	// Pathtype_Widget is a representation of the C type GTK_PATH_WIDGET.
	Pathtype_Widget Pathtype = 0
	// Pathtype_WidgetClass is a representation of the C type GTK_PATH_WIDGET_CLASS.
	Pathtype_WidgetClass Pathtype = 1
	// Pathtype_Class is a representation of the C type GTK_PATH_CLASS.
	Pathtype_Class Pathtype = 2
)

// Policytype is a representation of the C type GtkPolicyType.
type Policytype int

const (
	// Policytype_Always is a representation of the C type GTK_POLICY_ALWAYS.
	Policytype_Always Policytype = 0
	// Policytype_Automatic is a representation of the C type GTK_POLICY_AUTOMATIC.
	Policytype_Automatic Policytype = 1
	// Policytype_Never is a representation of the C type GTK_POLICY_NEVER.
	Policytype_Never Policytype = 2
	// Policytype_External is a representation of the C type GTK_POLICY_EXTERNAL.
	Policytype_External Policytype = 3
)

// Popoverconstraint is a representation of the C type GtkPopoverConstraint.
//
// since 3.20
type Popoverconstraint int

const (
	// Popoverconstraint_None is a representation of the C type GTK_POPOVER_CONSTRAINT_NONE.
	Popoverconstraint_None Popoverconstraint = 0
	// Popoverconstraint_Window is a representation of the C type GTK_POPOVER_CONSTRAINT_WINDOW.
	Popoverconstraint_Window Popoverconstraint = 1
)

// Positiontype is a representation of the C type GtkPositionType.
type Positiontype int

const (
	// Positiontype_Left is a representation of the C type GTK_POS_LEFT.
	Positiontype_Left Positiontype = 0
	// Positiontype_Right is a representation of the C type GTK_POS_RIGHT.
	Positiontype_Right Positiontype = 1
	// Positiontype_Top is a representation of the C type GTK_POS_TOP.
	Positiontype_Top Positiontype = 2
	// Positiontype_Bottom is a representation of the C type GTK_POS_BOTTOM.
	Positiontype_Bottom Positiontype = 3
)

// Printduplex is a representation of the C type GtkPrintDuplex.
type Printduplex int

const (
	// Printduplex_Simplex is a representation of the C type GTK_PRINT_DUPLEX_SIMPLEX.
	Printduplex_Simplex Printduplex = 0
	// Printduplex_Horizontal is a representation of the C type GTK_PRINT_DUPLEX_HORIZONTAL.
	Printduplex_Horizontal Printduplex = 1
	// Printduplex_Vertical is a representation of the C type GTK_PRINT_DUPLEX_VERTICAL.
	Printduplex_Vertical Printduplex = 2
)

// Printerror is a representation of the C type GtkPrintError.
type Printerror int

const (
	// Printerror_General is a representation of the C type GTK_PRINT_ERROR_GENERAL.
	Printerror_General Printerror = 0
	// Printerror_InternalError is a representation of the C type GTK_PRINT_ERROR_INTERNAL_ERROR.
	Printerror_InternalError Printerror = 1
	// Printerror_Nomem is a representation of the C type GTK_PRINT_ERROR_NOMEM.
	Printerror_Nomem Printerror = 2
	// Printerror_InvalidFile is a representation of the C type GTK_PRINT_ERROR_INVALID_FILE.
	Printerror_InvalidFile Printerror = 3
)

// Printoperationaction is a representation of the C type GtkPrintOperationAction.
type Printoperationaction int

const (
	// Printoperationaction_PrintDialog is a representation of the C type GTK_PRINT_OPERATION_ACTION_PRINT_DIALOG.
	Printoperationaction_PrintDialog Printoperationaction = 0
	// Printoperationaction_Print is a representation of the C type GTK_PRINT_OPERATION_ACTION_PRINT.
	Printoperationaction_Print Printoperationaction = 1
	// Printoperationaction_Preview is a representation of the C type GTK_PRINT_OPERATION_ACTION_PREVIEW.
	Printoperationaction_Preview Printoperationaction = 2
	// Printoperationaction_Export is a representation of the C type GTK_PRINT_OPERATION_ACTION_EXPORT.
	Printoperationaction_Export Printoperationaction = 3
)

// Printoperationresult is a representation of the C type GtkPrintOperationResult.
type Printoperationresult int

const (
	// Printoperationresult_Error is a representation of the C type GTK_PRINT_OPERATION_RESULT_ERROR.
	Printoperationresult_Error Printoperationresult = 0
	// Printoperationresult_Apply is a representation of the C type GTK_PRINT_OPERATION_RESULT_APPLY.
	Printoperationresult_Apply Printoperationresult = 1
	// Printoperationresult_Cancel is a representation of the C type GTK_PRINT_OPERATION_RESULT_CANCEL.
	Printoperationresult_Cancel Printoperationresult = 2
	// Printoperationresult_InProgress is a representation of the C type GTK_PRINT_OPERATION_RESULT_IN_PROGRESS.
	Printoperationresult_InProgress Printoperationresult = 3
)

// Printpages is a representation of the C type GtkPrintPages.
type Printpages int

const (
	// Printpages_All is a representation of the C type GTK_PRINT_PAGES_ALL.
	Printpages_All Printpages = 0
	// Printpages_Current is a representation of the C type GTK_PRINT_PAGES_CURRENT.
	Printpages_Current Printpages = 1
	// Printpages_Ranges is a representation of the C type GTK_PRINT_PAGES_RANGES.
	Printpages_Ranges Printpages = 2
	// Printpages_Selection is a representation of the C type GTK_PRINT_PAGES_SELECTION.
	Printpages_Selection Printpages = 3
)

// Printquality is a representation of the C type GtkPrintQuality.
type Printquality int

const (
	// Printquality_Low is a representation of the C type GTK_PRINT_QUALITY_LOW.
	Printquality_Low Printquality = 0
	// Printquality_Normal is a representation of the C type GTK_PRINT_QUALITY_NORMAL.
	Printquality_Normal Printquality = 1
	// Printquality_High is a representation of the C type GTK_PRINT_QUALITY_HIGH.
	Printquality_High Printquality = 2
	// Printquality_Draft is a representation of the C type GTK_PRINT_QUALITY_DRAFT.
	Printquality_Draft Printquality = 3
)

// Printstatus is a representation of the C type GtkPrintStatus.
type Printstatus int

const (
	// Printstatus_Initial is a representation of the C type GTK_PRINT_STATUS_INITIAL.
	Printstatus_Initial Printstatus = 0
	// Printstatus_Preparing is a representation of the C type GTK_PRINT_STATUS_PREPARING.
	Printstatus_Preparing Printstatus = 1
	// Printstatus_GeneratingData is a representation of the C type GTK_PRINT_STATUS_GENERATING_DATA.
	Printstatus_GeneratingData Printstatus = 2
	// Printstatus_SendingData is a representation of the C type GTK_PRINT_STATUS_SENDING_DATA.
	Printstatus_SendingData Printstatus = 3
	// Printstatus_Pending is a representation of the C type GTK_PRINT_STATUS_PENDING.
	Printstatus_Pending Printstatus = 4
	// Printstatus_PendingIssue is a representation of the C type GTK_PRINT_STATUS_PENDING_ISSUE.
	Printstatus_PendingIssue Printstatus = 5
	// Printstatus_Printing is a representation of the C type GTK_PRINT_STATUS_PRINTING.
	Printstatus_Printing Printstatus = 6
	// Printstatus_Finished is a representation of the C type GTK_PRINT_STATUS_FINISHED.
	Printstatus_Finished Printstatus = 7
	// Printstatus_FinishedAborted is a representation of the C type GTK_PRINT_STATUS_FINISHED_ABORTED.
	Printstatus_FinishedAborted Printstatus = 8
)

// Propagationphase is a representation of the C type GtkPropagationPhase.
//
// since 3.14
type Propagationphase int

const (
	// Propagationphase_None is a representation of the C type GTK_PHASE_NONE.
	Propagationphase_None Propagationphase = 0
	// Propagationphase_Capture is a representation of the C type GTK_PHASE_CAPTURE.
	Propagationphase_Capture Propagationphase = 1
	// Propagationphase_Bubble is a representation of the C type GTK_PHASE_BUBBLE.
	Propagationphase_Bubble Propagationphase = 2
	// Propagationphase_Target is a representation of the C type GTK_PHASE_TARGET.
	Propagationphase_Target Propagationphase = 3
)

// Rctokentype is a representation of the C type GtkRcTokenType.
type Rctokentype int

const (
	// Rctokentype_Invalid is a representation of the C type GTK_RC_TOKEN_INVALID.
	Rctokentype_Invalid Rctokentype = 270
	// Rctokentype_Include is a representation of the C type GTK_RC_TOKEN_INCLUDE.
	Rctokentype_Include Rctokentype = 271
	// Rctokentype_Normal is a representation of the C type GTK_RC_TOKEN_NORMAL.
	Rctokentype_Normal Rctokentype = 272
	// Rctokentype_Active is a representation of the C type GTK_RC_TOKEN_ACTIVE.
	Rctokentype_Active Rctokentype = 273
	// Rctokentype_Prelight is a representation of the C type GTK_RC_TOKEN_PRELIGHT.
	Rctokentype_Prelight Rctokentype = 274
	// Rctokentype_Selected is a representation of the C type GTK_RC_TOKEN_SELECTED.
	Rctokentype_Selected Rctokentype = 275
	// Rctokentype_Insensitive is a representation of the C type GTK_RC_TOKEN_INSENSITIVE.
	Rctokentype_Insensitive Rctokentype = 276
	// Rctokentype_Fg is a representation of the C type GTK_RC_TOKEN_FG.
	Rctokentype_Fg Rctokentype = 277
	// Rctokentype_Bg is a representation of the C type GTK_RC_TOKEN_BG.
	Rctokentype_Bg Rctokentype = 278
	// Rctokentype_Text is a representation of the C type GTK_RC_TOKEN_TEXT.
	Rctokentype_Text Rctokentype = 279
	// Rctokentype_Base is a representation of the C type GTK_RC_TOKEN_BASE.
	Rctokentype_Base Rctokentype = 280
	// Rctokentype_Xthickness is a representation of the C type GTK_RC_TOKEN_XTHICKNESS.
	Rctokentype_Xthickness Rctokentype = 281
	// Rctokentype_Ythickness is a representation of the C type GTK_RC_TOKEN_YTHICKNESS.
	Rctokentype_Ythickness Rctokentype = 282
	// Rctokentype_Font is a representation of the C type GTK_RC_TOKEN_FONT.
	Rctokentype_Font Rctokentype = 283
	// Rctokentype_Fontset is a representation of the C type GTK_RC_TOKEN_FONTSET.
	Rctokentype_Fontset Rctokentype = 284
	// Rctokentype_FontName is a representation of the C type GTK_RC_TOKEN_FONT_NAME.
	Rctokentype_FontName Rctokentype = 285
	// Rctokentype_BgPixmap is a representation of the C type GTK_RC_TOKEN_BG_PIXMAP.
	Rctokentype_BgPixmap Rctokentype = 286
	// Rctokentype_PixmapPath is a representation of the C type GTK_RC_TOKEN_PIXMAP_PATH.
	Rctokentype_PixmapPath Rctokentype = 287
	// Rctokentype_Style is a representation of the C type GTK_RC_TOKEN_STYLE.
	Rctokentype_Style Rctokentype = 288
	// Rctokentype_Binding is a representation of the C type GTK_RC_TOKEN_BINDING.
	Rctokentype_Binding Rctokentype = 289
	// Rctokentype_Bind is a representation of the C type GTK_RC_TOKEN_BIND.
	Rctokentype_Bind Rctokentype = 290
	// Rctokentype_Widget is a representation of the C type GTK_RC_TOKEN_WIDGET.
	Rctokentype_Widget Rctokentype = 291
	// Rctokentype_WidgetClass is a representation of the C type GTK_RC_TOKEN_WIDGET_CLASS.
	Rctokentype_WidgetClass Rctokentype = 292
	// Rctokentype_Class is a representation of the C type GTK_RC_TOKEN_CLASS.
	Rctokentype_Class Rctokentype = 293
	// Rctokentype_Lowest is a representation of the C type GTK_RC_TOKEN_LOWEST.
	Rctokentype_Lowest Rctokentype = 294
	// Rctokentype_Gtk is a representation of the C type GTK_RC_TOKEN_GTK.
	Rctokentype_Gtk Rctokentype = 295
	// Rctokentype_Application is a representation of the C type GTK_RC_TOKEN_APPLICATION.
	Rctokentype_Application Rctokentype = 296
	// Rctokentype_Theme is a representation of the C type GTK_RC_TOKEN_THEME.
	Rctokentype_Theme Rctokentype = 297
	// Rctokentype_Rc is a representation of the C type GTK_RC_TOKEN_RC.
	Rctokentype_Rc Rctokentype = 298
	// Rctokentype_Highest is a representation of the C type GTK_RC_TOKEN_HIGHEST.
	Rctokentype_Highest Rctokentype = 299
	// Rctokentype_Engine is a representation of the C type GTK_RC_TOKEN_ENGINE.
	Rctokentype_Engine Rctokentype = 300
	// Rctokentype_ModulePath is a representation of the C type GTK_RC_TOKEN_MODULE_PATH.
	Rctokentype_ModulePath Rctokentype = 301
	// Rctokentype_ImModulePath is a representation of the C type GTK_RC_TOKEN_IM_MODULE_PATH.
	Rctokentype_ImModulePath Rctokentype = 302
	// Rctokentype_ImModuleFile is a representation of the C type GTK_RC_TOKEN_IM_MODULE_FILE.
	Rctokentype_ImModuleFile Rctokentype = 303
	// Rctokentype_Stock is a representation of the C type GTK_RC_TOKEN_STOCK.
	Rctokentype_Stock Rctokentype = 304
	// Rctokentype_Ltr is a representation of the C type GTK_RC_TOKEN_LTR.
	Rctokentype_Ltr Rctokentype = 305
	// Rctokentype_Rtl is a representation of the C type GTK_RC_TOKEN_RTL.
	Rctokentype_Rtl Rctokentype = 306
	// Rctokentype_Color is a representation of the C type GTK_RC_TOKEN_COLOR.
	Rctokentype_Color Rctokentype = 307
	// Rctokentype_Unbind is a representation of the C type GTK_RC_TOKEN_UNBIND.
	Rctokentype_Unbind Rctokentype = 308
	// Rctokentype_Last is a representation of the C type GTK_RC_TOKEN_LAST.
	Rctokentype_Last Rctokentype = 309
)

// Recentchoosererror is a representation of the C type GtkRecentChooserError.
//
// since 2.10
type Recentchoosererror int

const (
	// Recentchoosererror_NotFound is a representation of the C type GTK_RECENT_CHOOSER_ERROR_NOT_FOUND.
	Recentchoosererror_NotFound Recentchoosererror = 0
	// Recentchoosererror_InvalidUri is a representation of the C type GTK_RECENT_CHOOSER_ERROR_INVALID_URI.
	Recentchoosererror_InvalidUri Recentchoosererror = 1
)

// Recentmanagererror is a representation of the C type GtkRecentManagerError.
//
// since 2.10
type Recentmanagererror int

const (
	// Recentmanagererror_NotFound is a representation of the C type GTK_RECENT_MANAGER_ERROR_NOT_FOUND.
	Recentmanagererror_NotFound Recentmanagererror = 0
	// Recentmanagererror_InvalidUri is a representation of the C type GTK_RECENT_MANAGER_ERROR_INVALID_URI.
	Recentmanagererror_InvalidUri Recentmanagererror = 1
	// Recentmanagererror_InvalidEncoding is a representation of the C type GTK_RECENT_MANAGER_ERROR_INVALID_ENCODING.
	Recentmanagererror_InvalidEncoding Recentmanagererror = 2
	// Recentmanagererror_NotRegistered is a representation of the C type GTK_RECENT_MANAGER_ERROR_NOT_REGISTERED.
	Recentmanagererror_NotRegistered Recentmanagererror = 3
	// Recentmanagererror_Read is a representation of the C type GTK_RECENT_MANAGER_ERROR_READ.
	Recentmanagererror_Read Recentmanagererror = 4
	// Recentmanagererror_Write is a representation of the C type GTK_RECENT_MANAGER_ERROR_WRITE.
	Recentmanagererror_Write Recentmanagererror = 5
	// Recentmanagererror_Unknown is a representation of the C type GTK_RECENT_MANAGER_ERROR_UNKNOWN.
	Recentmanagererror_Unknown Recentmanagererror = 6
)

// Recentsorttype is a representation of the C type GtkRecentSortType.
//
// since 2.10
type Recentsorttype int

const (
	// Recentsorttype_None is a representation of the C type GTK_RECENT_SORT_NONE.
	Recentsorttype_None Recentsorttype = 0
	// Recentsorttype_Mru is a representation of the C type GTK_RECENT_SORT_MRU.
	Recentsorttype_Mru Recentsorttype = 1
	// Recentsorttype_Lru is a representation of the C type GTK_RECENT_SORT_LRU.
	Recentsorttype_Lru Recentsorttype = 2
	// Recentsorttype_Custom is a representation of the C type GTK_RECENT_SORT_CUSTOM.
	Recentsorttype_Custom Recentsorttype = 3
)

// Reliefstyle is a representation of the C type GtkReliefStyle.
type Reliefstyle int

const (
	// Reliefstyle_Normal is a representation of the C type GTK_RELIEF_NORMAL.
	Reliefstyle_Normal Reliefstyle = 0
	// Reliefstyle_Half is a representation of the C type GTK_RELIEF_HALF.
	Reliefstyle_Half Reliefstyle = 1
	// Reliefstyle_None is a representation of the C type GTK_RELIEF_NONE.
	Reliefstyle_None Reliefstyle = 2
)

// Resizemode is a representation of the C type GtkResizeMode.
type Resizemode int

const (
	// Resizemode_Parent is a representation of the C type GTK_RESIZE_PARENT.
	Resizemode_Parent Resizemode = 0
	// Resizemode_Queue is a representation of the C type GTK_RESIZE_QUEUE.
	Resizemode_Queue Resizemode = 1
	// Resizemode_Immediate is a representation of the C type GTK_RESIZE_IMMEDIATE.
	Resizemode_Immediate Resizemode = 2
)

// Responsetype is a representation of the C type GtkResponseType.
type Responsetype int

const (
	// Responsetype_None is a representation of the C type GTK_RESPONSE_NONE.
	Responsetype_None Responsetype = -1
	// Responsetype_Reject is a representation of the C type GTK_RESPONSE_REJECT.
	Responsetype_Reject Responsetype = -2
	// Responsetype_Accept is a representation of the C type GTK_RESPONSE_ACCEPT.
	Responsetype_Accept Responsetype = -3
	// Responsetype_DeleteEvent is a representation of the C type GTK_RESPONSE_DELETE_EVENT.
	Responsetype_DeleteEvent Responsetype = -4
	// Responsetype_Ok is a representation of the C type GTK_RESPONSE_OK.
	Responsetype_Ok Responsetype = -5
	// Responsetype_Cancel is a representation of the C type GTK_RESPONSE_CANCEL.
	Responsetype_Cancel Responsetype = -6
	// Responsetype_Close is a representation of the C type GTK_RESPONSE_CLOSE.
	Responsetype_Close Responsetype = -7
	// Responsetype_Yes is a representation of the C type GTK_RESPONSE_YES.
	Responsetype_Yes Responsetype = -8
	// Responsetype_No is a representation of the C type GTK_RESPONSE_NO.
	Responsetype_No Responsetype = -9
	// Responsetype_Apply is a representation of the C type GTK_RESPONSE_APPLY.
	Responsetype_Apply Responsetype = -10
	// Responsetype_Help is a representation of the C type GTK_RESPONSE_HELP.
	Responsetype_Help Responsetype = -11
)

// Revealertransitiontype is a representation of the C type GtkRevealerTransitionType.
type Revealertransitiontype int

const (
	// Revealertransitiontype_None is a representation of the C type GTK_REVEALER_TRANSITION_TYPE_NONE.
	Revealertransitiontype_None Revealertransitiontype = 0
	// Revealertransitiontype_Crossfade is a representation of the C type GTK_REVEALER_TRANSITION_TYPE_CROSSFADE.
	Revealertransitiontype_Crossfade Revealertransitiontype = 1
	// Revealertransitiontype_SlideRight is a representation of the C type GTK_REVEALER_TRANSITION_TYPE_SLIDE_RIGHT.
	Revealertransitiontype_SlideRight Revealertransitiontype = 2
	// Revealertransitiontype_SlideLeft is a representation of the C type GTK_REVEALER_TRANSITION_TYPE_SLIDE_LEFT.
	Revealertransitiontype_SlideLeft Revealertransitiontype = 3
	// Revealertransitiontype_SlideUp is a representation of the C type GTK_REVEALER_TRANSITION_TYPE_SLIDE_UP.
	Revealertransitiontype_SlideUp Revealertransitiontype = 4
	// Revealertransitiontype_SlideDown is a representation of the C type GTK_REVEALER_TRANSITION_TYPE_SLIDE_DOWN.
	Revealertransitiontype_SlideDown Revealertransitiontype = 5
)

// Scrollstep is a representation of the C type GtkScrollStep.
type Scrollstep int

const (
	// Scrollstep_Steps is a representation of the C type GTK_SCROLL_STEPS.
	Scrollstep_Steps Scrollstep = 0
	// Scrollstep_Pages is a representation of the C type GTK_SCROLL_PAGES.
	Scrollstep_Pages Scrollstep = 1
	// Scrollstep_Ends is a representation of the C type GTK_SCROLL_ENDS.
	Scrollstep_Ends Scrollstep = 2
	// Scrollstep_HorizontalSteps is a representation of the C type GTK_SCROLL_HORIZONTAL_STEPS.
	Scrollstep_HorizontalSteps Scrollstep = 3
	// Scrollstep_HorizontalPages is a representation of the C type GTK_SCROLL_HORIZONTAL_PAGES.
	Scrollstep_HorizontalPages Scrollstep = 4
	// Scrollstep_HorizontalEnds is a representation of the C type GTK_SCROLL_HORIZONTAL_ENDS.
	Scrollstep_HorizontalEnds Scrollstep = 5
)

// Scrolltype is a representation of the C type GtkScrollType.
type Scrolltype int

const (
	// Scrolltype_None is a representation of the C type GTK_SCROLL_NONE.
	Scrolltype_None Scrolltype = 0
	// Scrolltype_Jump is a representation of the C type GTK_SCROLL_JUMP.
	Scrolltype_Jump Scrolltype = 1
	// Scrolltype_StepBackward is a representation of the C type GTK_SCROLL_STEP_BACKWARD.
	Scrolltype_StepBackward Scrolltype = 2
	// Scrolltype_StepForward is a representation of the C type GTK_SCROLL_STEP_FORWARD.
	Scrolltype_StepForward Scrolltype = 3
	// Scrolltype_PageBackward is a representation of the C type GTK_SCROLL_PAGE_BACKWARD.
	Scrolltype_PageBackward Scrolltype = 4
	// Scrolltype_PageForward is a representation of the C type GTK_SCROLL_PAGE_FORWARD.
	Scrolltype_PageForward Scrolltype = 5
	// Scrolltype_StepUp is a representation of the C type GTK_SCROLL_STEP_UP.
	Scrolltype_StepUp Scrolltype = 6
	// Scrolltype_StepDown is a representation of the C type GTK_SCROLL_STEP_DOWN.
	Scrolltype_StepDown Scrolltype = 7
	// Scrolltype_PageUp is a representation of the C type GTK_SCROLL_PAGE_UP.
	Scrolltype_PageUp Scrolltype = 8
	// Scrolltype_PageDown is a representation of the C type GTK_SCROLL_PAGE_DOWN.
	Scrolltype_PageDown Scrolltype = 9
	// Scrolltype_StepLeft is a representation of the C type GTK_SCROLL_STEP_LEFT.
	Scrolltype_StepLeft Scrolltype = 10
	// Scrolltype_StepRight is a representation of the C type GTK_SCROLL_STEP_RIGHT.
	Scrolltype_StepRight Scrolltype = 11
	// Scrolltype_PageLeft is a representation of the C type GTK_SCROLL_PAGE_LEFT.
	Scrolltype_PageLeft Scrolltype = 12
	// Scrolltype_PageRight is a representation of the C type GTK_SCROLL_PAGE_RIGHT.
	Scrolltype_PageRight Scrolltype = 13
	// Scrolltype_Start is a representation of the C type GTK_SCROLL_START.
	Scrolltype_Start Scrolltype = 14
	// Scrolltype_End is a representation of the C type GTK_SCROLL_END.
	Scrolltype_End Scrolltype = 15
)

// Scrollablepolicy is a representation of the C type GtkScrollablePolicy.
type Scrollablepolicy int

const (
	// Scrollablepolicy_Minimum is a representation of the C type GTK_SCROLL_MINIMUM.
	Scrollablepolicy_Minimum Scrollablepolicy = 0
	// Scrollablepolicy_Natural is a representation of the C type GTK_SCROLL_NATURAL.
	Scrollablepolicy_Natural Scrollablepolicy = 1
)

// Selectionmode is a representation of the C type GtkSelectionMode.
type Selectionmode int

const (
	// Selectionmode_None is a representation of the C type GTK_SELECTION_NONE.
	Selectionmode_None Selectionmode = 0
	// Selectionmode_Single is a representation of the C type GTK_SELECTION_SINGLE.
	Selectionmode_Single Selectionmode = 1
	// Selectionmode_Browse is a representation of the C type GTK_SELECTION_BROWSE.
	Selectionmode_Browse Selectionmode = 2
	// Selectionmode_Multiple is a representation of the C type GTK_SELECTION_MULTIPLE.
	Selectionmode_Multiple Selectionmode = 3
)

// Sensitivitytype is a representation of the C type GtkSensitivityType.
type Sensitivitytype int

const (
	// Sensitivitytype_Auto is a representation of the C type GTK_SENSITIVITY_AUTO.
	Sensitivitytype_Auto Sensitivitytype = 0
	// Sensitivitytype_On is a representation of the C type GTK_SENSITIVITY_ON.
	Sensitivitytype_On Sensitivitytype = 1
	// Sensitivitytype_Off is a representation of the C type GTK_SENSITIVITY_OFF.
	Sensitivitytype_Off Sensitivitytype = 2
)

// Shadowtype is a representation of the C type GtkShadowType.
type Shadowtype int

const (
	// Shadowtype_None is a representation of the C type GTK_SHADOW_NONE.
	Shadowtype_None Shadowtype = 0
	// Shadowtype_In is a representation of the C type GTK_SHADOW_IN.
	Shadowtype_In Shadowtype = 1
	// Shadowtype_Out is a representation of the C type GTK_SHADOW_OUT.
	Shadowtype_Out Shadowtype = 2
	// Shadowtype_EtchedIn is a representation of the C type GTK_SHADOW_ETCHED_IN.
	Shadowtype_EtchedIn Shadowtype = 3
	// Shadowtype_EtchedOut is a representation of the C type GTK_SHADOW_ETCHED_OUT.
	Shadowtype_EtchedOut Shadowtype = 4
)

// Shortcuttype is a representation of the C type GtkShortcutType.
//
// since 3.20
type Shortcuttype int

const (
	// Shortcuttype_Accelerator is a representation of the C type GTK_SHORTCUT_ACCELERATOR.
	Shortcuttype_Accelerator Shortcuttype = 0
	// Shortcuttype_GesturePinch is a representation of the C type GTK_SHORTCUT_GESTURE_PINCH.
	Shortcuttype_GesturePinch Shortcuttype = 1
	// Shortcuttype_GestureStretch is a representation of the C type GTK_SHORTCUT_GESTURE_STRETCH.
	Shortcuttype_GestureStretch Shortcuttype = 2
	// Shortcuttype_GestureRotateClockwise is a representation of the C type GTK_SHORTCUT_GESTURE_ROTATE_CLOCKWISE.
	Shortcuttype_GestureRotateClockwise Shortcuttype = 3
	// Shortcuttype_GestureRotateCounterclockwise is a representation of the C type GTK_SHORTCUT_GESTURE_ROTATE_COUNTERCLOCKWISE.
	Shortcuttype_GestureRotateCounterclockwise Shortcuttype = 4
	// Shortcuttype_GestureTwoFingerSwipeLeft is a representation of the C type GTK_SHORTCUT_GESTURE_TWO_FINGER_SWIPE_LEFT.
	Shortcuttype_GestureTwoFingerSwipeLeft Shortcuttype = 5
	// Shortcuttype_GestureTwoFingerSwipeRight is a representation of the C type GTK_SHORTCUT_GESTURE_TWO_FINGER_SWIPE_RIGHT.
	Shortcuttype_GestureTwoFingerSwipeRight Shortcuttype = 6
	// Shortcuttype_Gesture is a representation of the C type GTK_SHORTCUT_GESTURE.
	Shortcuttype_Gesture Shortcuttype = 7
)

// Sizegroupmode is a representation of the C type GtkSizeGroupMode.
type Sizegroupmode int

const (
	// Sizegroupmode_None is a representation of the C type GTK_SIZE_GROUP_NONE.
	Sizegroupmode_None Sizegroupmode = 0
	// Sizegroupmode_Horizontal is a representation of the C type GTK_SIZE_GROUP_HORIZONTAL.
	Sizegroupmode_Horizontal Sizegroupmode = 1
	// Sizegroupmode_Vertical is a representation of the C type GTK_SIZE_GROUP_VERTICAL.
	Sizegroupmode_Vertical Sizegroupmode = 2
	// Sizegroupmode_Both is a representation of the C type GTK_SIZE_GROUP_BOTH.
	Sizegroupmode_Both Sizegroupmode = 3
)

// Sizerequestmode is a representation of the C type GtkSizeRequestMode.
type Sizerequestmode int

const (
	// Sizerequestmode_HeightForWidth is a representation of the C type GTK_SIZE_REQUEST_HEIGHT_FOR_WIDTH.
	Sizerequestmode_HeightForWidth Sizerequestmode = 0
	// Sizerequestmode_WidthForHeight is a representation of the C type GTK_SIZE_REQUEST_WIDTH_FOR_HEIGHT.
	Sizerequestmode_WidthForHeight Sizerequestmode = 1
	// Sizerequestmode_ConstantSize is a representation of the C type GTK_SIZE_REQUEST_CONSTANT_SIZE.
	Sizerequestmode_ConstantSize Sizerequestmode = 2
)

// Sorttype is a representation of the C type GtkSortType.
type Sorttype int

const (
	// Sorttype_Ascending is a representation of the C type GTK_SORT_ASCENDING.
	Sorttype_Ascending Sorttype = 0
	// Sorttype_Descending is a representation of the C type GTK_SORT_DESCENDING.
	Sorttype_Descending Sorttype = 1
)

// Spinbuttonupdatepolicy is a representation of the C type GtkSpinButtonUpdatePolicy.
type Spinbuttonupdatepolicy int

const (
	// Spinbuttonupdatepolicy_Always is a representation of the C type GTK_UPDATE_ALWAYS.
	Spinbuttonupdatepolicy_Always Spinbuttonupdatepolicy = 0
	// Spinbuttonupdatepolicy_IfValid is a representation of the C type GTK_UPDATE_IF_VALID.
	Spinbuttonupdatepolicy_IfValid Spinbuttonupdatepolicy = 1
)

// Spintype is a representation of the C type GtkSpinType.
type Spintype int

const (
	// Spintype_StepForward is a representation of the C type GTK_SPIN_STEP_FORWARD.
	Spintype_StepForward Spintype = 0
	// Spintype_StepBackward is a representation of the C type GTK_SPIN_STEP_BACKWARD.
	Spintype_StepBackward Spintype = 1
	// Spintype_PageForward is a representation of the C type GTK_SPIN_PAGE_FORWARD.
	Spintype_PageForward Spintype = 2
	// Spintype_PageBackward is a representation of the C type GTK_SPIN_PAGE_BACKWARD.
	Spintype_PageBackward Spintype = 3
	// Spintype_Home is a representation of the C type GTK_SPIN_HOME.
	Spintype_Home Spintype = 4
	// Spintype_End is a representation of the C type GTK_SPIN_END.
	Spintype_End Spintype = 5
	// Spintype_UserDefined is a representation of the C type GTK_SPIN_USER_DEFINED.
	Spintype_UserDefined Spintype = 6
)

// Stacktransitiontype is a representation of the C type GtkStackTransitionType.
type Stacktransitiontype int

const (
	// Stacktransitiontype_None is a representation of the C type GTK_STACK_TRANSITION_TYPE_NONE.
	Stacktransitiontype_None Stacktransitiontype = 0
	// Stacktransitiontype_Crossfade is a representation of the C type GTK_STACK_TRANSITION_TYPE_CROSSFADE.
	Stacktransitiontype_Crossfade Stacktransitiontype = 1
	// Stacktransitiontype_SlideRight is a representation of the C type GTK_STACK_TRANSITION_TYPE_SLIDE_RIGHT.
	Stacktransitiontype_SlideRight Stacktransitiontype = 2
	// Stacktransitiontype_SlideLeft is a representation of the C type GTK_STACK_TRANSITION_TYPE_SLIDE_LEFT.
	Stacktransitiontype_SlideLeft Stacktransitiontype = 3
	// Stacktransitiontype_SlideUp is a representation of the C type GTK_STACK_TRANSITION_TYPE_SLIDE_UP.
	Stacktransitiontype_SlideUp Stacktransitiontype = 4
	// Stacktransitiontype_SlideDown is a representation of the C type GTK_STACK_TRANSITION_TYPE_SLIDE_DOWN.
	Stacktransitiontype_SlideDown Stacktransitiontype = 5
	// Stacktransitiontype_SlideLeftRight is a representation of the C type GTK_STACK_TRANSITION_TYPE_SLIDE_LEFT_RIGHT.
	Stacktransitiontype_SlideLeftRight Stacktransitiontype = 6
	// Stacktransitiontype_SlideUpDown is a representation of the C type GTK_STACK_TRANSITION_TYPE_SLIDE_UP_DOWN.
	Stacktransitiontype_SlideUpDown Stacktransitiontype = 7
	// Stacktransitiontype_OverUp is a representation of the C type GTK_STACK_TRANSITION_TYPE_OVER_UP.
	Stacktransitiontype_OverUp Stacktransitiontype = 8
	// Stacktransitiontype_OverDown is a representation of the C type GTK_STACK_TRANSITION_TYPE_OVER_DOWN.
	Stacktransitiontype_OverDown Stacktransitiontype = 9
	// Stacktransitiontype_OverLeft is a representation of the C type GTK_STACK_TRANSITION_TYPE_OVER_LEFT.
	Stacktransitiontype_OverLeft Stacktransitiontype = 10
	// Stacktransitiontype_OverRight is a representation of the C type GTK_STACK_TRANSITION_TYPE_OVER_RIGHT.
	Stacktransitiontype_OverRight Stacktransitiontype = 11
	// Stacktransitiontype_UnderUp is a representation of the C type GTK_STACK_TRANSITION_TYPE_UNDER_UP.
	Stacktransitiontype_UnderUp Stacktransitiontype = 12
	// Stacktransitiontype_UnderDown is a representation of the C type GTK_STACK_TRANSITION_TYPE_UNDER_DOWN.
	Stacktransitiontype_UnderDown Stacktransitiontype = 13
	// Stacktransitiontype_UnderLeft is a representation of the C type GTK_STACK_TRANSITION_TYPE_UNDER_LEFT.
	Stacktransitiontype_UnderLeft Stacktransitiontype = 14
	// Stacktransitiontype_UnderRight is a representation of the C type GTK_STACK_TRANSITION_TYPE_UNDER_RIGHT.
	Stacktransitiontype_UnderRight Stacktransitiontype = 15
	// Stacktransitiontype_OverUpDown is a representation of the C type GTK_STACK_TRANSITION_TYPE_OVER_UP_DOWN.
	Stacktransitiontype_OverUpDown Stacktransitiontype = 16
	// Stacktransitiontype_OverDownUp is a representation of the C type GTK_STACK_TRANSITION_TYPE_OVER_DOWN_UP.
	Stacktransitiontype_OverDownUp Stacktransitiontype = 17
	// Stacktransitiontype_OverLeftRight is a representation of the C type GTK_STACK_TRANSITION_TYPE_OVER_LEFT_RIGHT.
	Stacktransitiontype_OverLeftRight Stacktransitiontype = 18
	// Stacktransitiontype_OverRightLeft is a representation of the C type GTK_STACK_TRANSITION_TYPE_OVER_RIGHT_LEFT.
	Stacktransitiontype_OverRightLeft Stacktransitiontype = 19
)

// Statetype is a representation of the C type GtkStateType.
type Statetype int

const (
	// Statetype_Normal is a representation of the C type GTK_STATE_NORMAL.
	Statetype_Normal Statetype = 0
	// Statetype_Active is a representation of the C type GTK_STATE_ACTIVE.
	Statetype_Active Statetype = 1
	// Statetype_Prelight is a representation of the C type GTK_STATE_PRELIGHT.
	Statetype_Prelight Statetype = 2
	// Statetype_Selected is a representation of the C type GTK_STATE_SELECTED.
	Statetype_Selected Statetype = 3
	// Statetype_Insensitive is a representation of the C type GTK_STATE_INSENSITIVE.
	Statetype_Insensitive Statetype = 4
	// Statetype_Inconsistent is a representation of the C type GTK_STATE_INCONSISTENT.
	Statetype_Inconsistent Statetype = 5
	// Statetype_Focused is a representation of the C type GTK_STATE_FOCUSED.
	Statetype_Focused Statetype = 6
)

// Textbuffertargetinfo is a representation of the C type GtkTextBufferTargetInfo.
type Textbuffertargetinfo int

const (
	// Textbuffertargetinfo_BufferContents is a representation of the C type GTK_TEXT_BUFFER_TARGET_INFO_BUFFER_CONTENTS.
	Textbuffertargetinfo_BufferContents Textbuffertargetinfo = -1
	// Textbuffertargetinfo_RichText is a representation of the C type GTK_TEXT_BUFFER_TARGET_INFO_RICH_TEXT.
	Textbuffertargetinfo_RichText Textbuffertargetinfo = -2
	// Textbuffertargetinfo_Text is a representation of the C type GTK_TEXT_BUFFER_TARGET_INFO_TEXT.
	Textbuffertargetinfo_Text Textbuffertargetinfo = -3
)

// Textdirection is a representation of the C type GtkTextDirection.
type Textdirection int

const (
	// Textdirection_None is a representation of the C type GTK_TEXT_DIR_NONE.
	Textdirection_None Textdirection = 0
	// Textdirection_Ltr is a representation of the C type GTK_TEXT_DIR_LTR.
	Textdirection_Ltr Textdirection = 1
	// Textdirection_Rtl is a representation of the C type GTK_TEXT_DIR_RTL.
	Textdirection_Rtl Textdirection = 2
)

// Textextendselection is a representation of the C type GtkTextExtendSelection.
//
// since 3.16
type Textextendselection int

const (
	// Textextendselection_Word is a representation of the C type GTK_TEXT_EXTEND_SELECTION_WORD.
	Textextendselection_Word Textextendselection = 0
	// Textextendselection_Line is a representation of the C type GTK_TEXT_EXTEND_SELECTION_LINE.
	Textextendselection_Line Textextendselection = 1
)

// Textviewlayer is a representation of the C type GtkTextViewLayer.
type Textviewlayer int

const (
	// Textviewlayer_Below is a representation of the C type GTK_TEXT_VIEW_LAYER_BELOW.
	Textviewlayer_Below Textviewlayer = 0
	// Textviewlayer_Above is a representation of the C type GTK_TEXT_VIEW_LAYER_ABOVE.
	Textviewlayer_Above Textviewlayer = 1
	// Textviewlayer_BelowText is a representation of the C type GTK_TEXT_VIEW_LAYER_BELOW_TEXT.
	Textviewlayer_BelowText Textviewlayer = 2
	// Textviewlayer_AboveText is a representation of the C type GTK_TEXT_VIEW_LAYER_ABOVE_TEXT.
	Textviewlayer_AboveText Textviewlayer = 3
)

// Textwindowtype is a representation of the C type GtkTextWindowType.
type Textwindowtype int

const (
	// Textwindowtype_Private is a representation of the C type GTK_TEXT_WINDOW_PRIVATE.
	Textwindowtype_Private Textwindowtype = 0
	// Textwindowtype_Widget is a representation of the C type GTK_TEXT_WINDOW_WIDGET.
	Textwindowtype_Widget Textwindowtype = 1
	// Textwindowtype_Text is a representation of the C type GTK_TEXT_WINDOW_TEXT.
	Textwindowtype_Text Textwindowtype = 2
	// Textwindowtype_Left is a representation of the C type GTK_TEXT_WINDOW_LEFT.
	Textwindowtype_Left Textwindowtype = 3
	// Textwindowtype_Right is a representation of the C type GTK_TEXT_WINDOW_RIGHT.
	Textwindowtype_Right Textwindowtype = 4
	// Textwindowtype_Top is a representation of the C type GTK_TEXT_WINDOW_TOP.
	Textwindowtype_Top Textwindowtype = 5
	// Textwindowtype_Bottom is a representation of the C type GTK_TEXT_WINDOW_BOTTOM.
	Textwindowtype_Bottom Textwindowtype = 6
)

// Toolbarspacestyle is a representation of the C type GtkToolbarSpaceStyle.
type Toolbarspacestyle int

const (
	// Toolbarspacestyle_Empty is a representation of the C type GTK_TOOLBAR_SPACE_EMPTY.
	Toolbarspacestyle_Empty Toolbarspacestyle = 0
	// Toolbarspacestyle_Line is a representation of the C type GTK_TOOLBAR_SPACE_LINE.
	Toolbarspacestyle_Line Toolbarspacestyle = 1
)

// Toolbarstyle is a representation of the C type GtkToolbarStyle.
type Toolbarstyle int

const (
	// Toolbarstyle_Icons is a representation of the C type GTK_TOOLBAR_ICONS.
	Toolbarstyle_Icons Toolbarstyle = 0
	// Toolbarstyle_Text is a representation of the C type GTK_TOOLBAR_TEXT.
	Toolbarstyle_Text Toolbarstyle = 1
	// Toolbarstyle_Both is a representation of the C type GTK_TOOLBAR_BOTH.
	Toolbarstyle_Both Toolbarstyle = 2
	// Toolbarstyle_BothHoriz is a representation of the C type GTK_TOOLBAR_BOTH_HORIZ.
	Toolbarstyle_BothHoriz Toolbarstyle = 3
)

// Treeviewcolumnsizing is a representation of the C type GtkTreeViewColumnSizing.
type Treeviewcolumnsizing int

const (
	// Treeviewcolumnsizing_GrowOnly is a representation of the C type GTK_TREE_VIEW_COLUMN_GROW_ONLY.
	Treeviewcolumnsizing_GrowOnly Treeviewcolumnsizing = 0
	// Treeviewcolumnsizing_Autosize is a representation of the C type GTK_TREE_VIEW_COLUMN_AUTOSIZE.
	Treeviewcolumnsizing_Autosize Treeviewcolumnsizing = 1
	// Treeviewcolumnsizing_Fixed is a representation of the C type GTK_TREE_VIEW_COLUMN_FIXED.
	Treeviewcolumnsizing_Fixed Treeviewcolumnsizing = 2
)

// Treeviewdropposition is a representation of the C type GtkTreeViewDropPosition.
type Treeviewdropposition int

const (
	// Treeviewdropposition_Before is a representation of the C type GTK_TREE_VIEW_DROP_BEFORE.
	Treeviewdropposition_Before Treeviewdropposition = 0
	// Treeviewdropposition_After is a representation of the C type GTK_TREE_VIEW_DROP_AFTER.
	Treeviewdropposition_After Treeviewdropposition = 1
	// Treeviewdropposition_IntoOrBefore is a representation of the C type GTK_TREE_VIEW_DROP_INTO_OR_BEFORE.
	Treeviewdropposition_IntoOrBefore Treeviewdropposition = 2
	// Treeviewdropposition_IntoOrAfter is a representation of the C type GTK_TREE_VIEW_DROP_INTO_OR_AFTER.
	Treeviewdropposition_IntoOrAfter Treeviewdropposition = 3
)

// Treeviewgridlines is a representation of the C type GtkTreeViewGridLines.
type Treeviewgridlines int

const (
	// Treeviewgridlines_None is a representation of the C type GTK_TREE_VIEW_GRID_LINES_NONE.
	Treeviewgridlines_None Treeviewgridlines = 0
	// Treeviewgridlines_Horizontal is a representation of the C type GTK_TREE_VIEW_GRID_LINES_HORIZONTAL.
	Treeviewgridlines_Horizontal Treeviewgridlines = 1
	// Treeviewgridlines_Vertical is a representation of the C type GTK_TREE_VIEW_GRID_LINES_VERTICAL.
	Treeviewgridlines_Vertical Treeviewgridlines = 2
	// Treeviewgridlines_Both is a representation of the C type GTK_TREE_VIEW_GRID_LINES_BOTH.
	Treeviewgridlines_Both Treeviewgridlines = 3
)

// Unit is a representation of the C type GtkUnit.
type Unit int

const (
	// Unit_None is a representation of the C type GTK_UNIT_NONE.
	Unit_None Unit = 0
	// Unit_Points is a representation of the C type GTK_UNIT_POINTS.
	Unit_Points Unit = 1
	// Unit_Inch is a representation of the C type GTK_UNIT_INCH.
	Unit_Inch Unit = 2
	// Unit_Mm is a representation of the C type GTK_UNIT_MM.
	Unit_Mm Unit = 3
)

// Widgethelptype is a representation of the C type GtkWidgetHelpType.
type Widgethelptype int

const (
	// Widgethelptype_Tooltip is a representation of the C type GTK_WIDGET_HELP_TOOLTIP.
	Widgethelptype_Tooltip Widgethelptype = 0
	// Widgethelptype_WhatsThis is a representation of the C type GTK_WIDGET_HELP_WHATS_THIS.
	Widgethelptype_WhatsThis Widgethelptype = 1
)

// Windowposition is a representation of the C type GtkWindowPosition.
type Windowposition int

const (
	// Windowposition_None is a representation of the C type GTK_WIN_POS_NONE.
	Windowposition_None Windowposition = 0
	// Windowposition_Center is a representation of the C type GTK_WIN_POS_CENTER.
	Windowposition_Center Windowposition = 1
	// Windowposition_Mouse is a representation of the C type GTK_WIN_POS_MOUSE.
	Windowposition_Mouse Windowposition = 2
	// Windowposition_CenterAlways is a representation of the C type GTK_WIN_POS_CENTER_ALWAYS.
	Windowposition_CenterAlways Windowposition = 3
	// Windowposition_CenterOnParent is a representation of the C type GTK_WIN_POS_CENTER_ON_PARENT.
	Windowposition_CenterOnParent Windowposition = 4
)

// Windowtype is a representation of the C type GtkWindowType.
type Windowtype int

const (
	// Windowtype_Toplevel is a representation of the C type GTK_WINDOW_TOPLEVEL.
	Windowtype_Toplevel Windowtype = 0
	// Windowtype_Popup is a representation of the C type GTK_WINDOW_POPUP.
	Windowtype_Popup Windowtype = 1
)

// Wrapmode is a representation of the C type GtkWrapMode.
type Wrapmode int

const (
	// Wrapmode_None is a representation of the C type GTK_WRAP_NONE.
	Wrapmode_None Wrapmode = 0
	// Wrapmode_Char is a representation of the C type GTK_WRAP_CHAR.
	Wrapmode_Char Wrapmode = 1
	// Wrapmode_Word is a representation of the C type GTK_WRAP_WORD.
	Wrapmode_Word Wrapmode = 2
	// Wrapmode_WordChar is a representation of the C type GTK_WRAP_WORD_CHAR.
	Wrapmode_WordChar Wrapmode = 3
)
