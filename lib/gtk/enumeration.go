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

// ArrowPlacement is a representation of the C type GtkArrowPlacement.
type ArrowPlacement int

const (
	// ArrowPlacement_Both is a representation of the C type GTK_ARROWS_BOTH.
	ArrowPlacement_Both ArrowPlacement = 0
	// ArrowPlacement_Start is a representation of the C type GTK_ARROWS_START.
	ArrowPlacement_Start ArrowPlacement = 1
	// ArrowPlacement_End is a representation of the C type GTK_ARROWS_END.
	ArrowPlacement_End ArrowPlacement = 2
)

// ArrowType is a representation of the C type GtkArrowType.
type ArrowType int

const (
	// ArrowType_Up is a representation of the C type GTK_ARROW_UP.
	ArrowType_Up ArrowType = 0
	// ArrowType_Down is a representation of the C type GTK_ARROW_DOWN.
	ArrowType_Down ArrowType = 1
	// ArrowType_Left is a representation of the C type GTK_ARROW_LEFT.
	ArrowType_Left ArrowType = 2
	// ArrowType_Right is a representation of the C type GTK_ARROW_RIGHT.
	ArrowType_Right ArrowType = 3
	// ArrowType_None is a representation of the C type GTK_ARROW_NONE.
	ArrowType_None ArrowType = 4
)

// AssistantPageType is a representation of the C type GtkAssistantPageType.
type AssistantPageType int

const (
	// AssistantPageType_Content is a representation of the C type GTK_ASSISTANT_PAGE_CONTENT.
	AssistantPageType_Content AssistantPageType = 0
	// AssistantPageType_Intro is a representation of the C type GTK_ASSISTANT_PAGE_INTRO.
	AssistantPageType_Intro AssistantPageType = 1
	// AssistantPageType_Confirm is a representation of the C type GTK_ASSISTANT_PAGE_CONFIRM.
	AssistantPageType_Confirm AssistantPageType = 2
	// AssistantPageType_Summary is a representation of the C type GTK_ASSISTANT_PAGE_SUMMARY.
	AssistantPageType_Summary AssistantPageType = 3
	// AssistantPageType_Progress is a representation of the C type GTK_ASSISTANT_PAGE_PROGRESS.
	AssistantPageType_Progress AssistantPageType = 4
	// AssistantPageType_Custom is a representation of the C type GTK_ASSISTANT_PAGE_CUSTOM.
	AssistantPageType_Custom AssistantPageType = 5
)

// BaselinePosition is a representation of the C type GtkBaselinePosition.
//
// since 3.10
type BaselinePosition int

const (
	// BaselinePosition_Top is a representation of the C type GTK_BASELINE_POSITION_TOP.
	BaselinePosition_Top BaselinePosition = 0
	// BaselinePosition_Center is a representation of the C type GTK_BASELINE_POSITION_CENTER.
	BaselinePosition_Center BaselinePosition = 1
	// BaselinePosition_Bottom is a representation of the C type GTK_BASELINE_POSITION_BOTTOM.
	BaselinePosition_Bottom BaselinePosition = 2
)

// BorderStyle is a representation of the C type GtkBorderStyle.
type BorderStyle int

const (
	// BorderStyle_None is a representation of the C type GTK_BORDER_STYLE_NONE.
	BorderStyle_None BorderStyle = 0
	// BorderStyle_Solid is a representation of the C type GTK_BORDER_STYLE_SOLID.
	BorderStyle_Solid BorderStyle = 1
	// BorderStyle_Inset is a representation of the C type GTK_BORDER_STYLE_INSET.
	BorderStyle_Inset BorderStyle = 2
	// BorderStyle_Outset is a representation of the C type GTK_BORDER_STYLE_OUTSET.
	BorderStyle_Outset BorderStyle = 3
	// BorderStyle_Hidden is a representation of the C type GTK_BORDER_STYLE_HIDDEN.
	BorderStyle_Hidden BorderStyle = 4
	// BorderStyle_Dotted is a representation of the C type GTK_BORDER_STYLE_DOTTED.
	BorderStyle_Dotted BorderStyle = 5
	// BorderStyle_Dashed is a representation of the C type GTK_BORDER_STYLE_DASHED.
	BorderStyle_Dashed BorderStyle = 6
	// BorderStyle_Double is a representation of the C type GTK_BORDER_STYLE_DOUBLE.
	BorderStyle_Double BorderStyle = 7
	// BorderStyle_Groove is a representation of the C type GTK_BORDER_STYLE_GROOVE.
	BorderStyle_Groove BorderStyle = 8
	// BorderStyle_Ridge is a representation of the C type GTK_BORDER_STYLE_RIDGE.
	BorderStyle_Ridge BorderStyle = 9
)

// BuilderError is a representation of the C type GtkBuilderError.
type BuilderError int

const (
	// BuilderError_InvalidTypeFunction is a representation of the C type GTK_BUILDER_ERROR_INVALID_TYPE_FUNCTION.
	BuilderError_InvalidTypeFunction BuilderError = 0
	// BuilderError_UnhandledTag is a representation of the C type GTK_BUILDER_ERROR_UNHANDLED_TAG.
	BuilderError_UnhandledTag BuilderError = 1
	// BuilderError_MissingAttribute is a representation of the C type GTK_BUILDER_ERROR_MISSING_ATTRIBUTE.
	BuilderError_MissingAttribute BuilderError = 2
	// BuilderError_InvalidAttribute is a representation of the C type GTK_BUILDER_ERROR_INVALID_ATTRIBUTE.
	BuilderError_InvalidAttribute BuilderError = 3
	// BuilderError_InvalidTag is a representation of the C type GTK_BUILDER_ERROR_INVALID_TAG.
	BuilderError_InvalidTag BuilderError = 4
	// BuilderError_MissingPropertyValue is a representation of the C type GTK_BUILDER_ERROR_MISSING_PROPERTY_VALUE.
	BuilderError_MissingPropertyValue BuilderError = 5
	// BuilderError_InvalidValue is a representation of the C type GTK_BUILDER_ERROR_INVALID_VALUE.
	BuilderError_InvalidValue BuilderError = 6
	// BuilderError_VersionMismatch is a representation of the C type GTK_BUILDER_ERROR_VERSION_MISMATCH.
	BuilderError_VersionMismatch BuilderError = 7
	// BuilderError_DuplicateId is a representation of the C type GTK_BUILDER_ERROR_DUPLICATE_ID.
	BuilderError_DuplicateId BuilderError = 8
	// BuilderError_ObjectTypeRefused is a representation of the C type GTK_BUILDER_ERROR_OBJECT_TYPE_REFUSED.
	BuilderError_ObjectTypeRefused BuilderError = 9
	// BuilderError_TemplateMismatch is a representation of the C type GTK_BUILDER_ERROR_TEMPLATE_MISMATCH.
	BuilderError_TemplateMismatch BuilderError = 10
	// BuilderError_InvalidProperty is a representation of the C type GTK_BUILDER_ERROR_INVALID_PROPERTY.
	BuilderError_InvalidProperty BuilderError = 11
	// BuilderError_InvalidSignal is a representation of the C type GTK_BUILDER_ERROR_INVALID_SIGNAL.
	BuilderError_InvalidSignal BuilderError = 12
	// BuilderError_InvalidId is a representation of the C type GTK_BUILDER_ERROR_INVALID_ID.
	BuilderError_InvalidId BuilderError = 13
)

// ButtonBoxStyle is a representation of the C type GtkButtonBoxStyle.
type ButtonBoxStyle int

const (
	// ButtonBoxStyle_Spread is a representation of the C type GTK_BUTTONBOX_SPREAD.
	ButtonBoxStyle_Spread ButtonBoxStyle = 1
	// ButtonBoxStyle_Edge is a representation of the C type GTK_BUTTONBOX_EDGE.
	ButtonBoxStyle_Edge ButtonBoxStyle = 2
	// ButtonBoxStyle_Start is a representation of the C type GTK_BUTTONBOX_START.
	ButtonBoxStyle_Start ButtonBoxStyle = 3
	// ButtonBoxStyle_End is a representation of the C type GTK_BUTTONBOX_END.
	ButtonBoxStyle_End ButtonBoxStyle = 4
	// ButtonBoxStyle_Center is a representation of the C type GTK_BUTTONBOX_CENTER.
	ButtonBoxStyle_Center ButtonBoxStyle = 5
	// ButtonBoxStyle_Expand is a representation of the C type GTK_BUTTONBOX_EXPAND.
	ButtonBoxStyle_Expand ButtonBoxStyle = 6
)

// ButtonRole is a representation of the C type GtkButtonRole.
type ButtonRole int

const (
	// ButtonRole_Normal is a representation of the C type GTK_BUTTON_ROLE_NORMAL.
	ButtonRole_Normal ButtonRole = 0
	// ButtonRole_Check is a representation of the C type GTK_BUTTON_ROLE_CHECK.
	ButtonRole_Check ButtonRole = 1
	// ButtonRole_Radio is a representation of the C type GTK_BUTTON_ROLE_RADIO.
	ButtonRole_Radio ButtonRole = 2
)

// ButtonsType is a representation of the C type GtkButtonsType.
type ButtonsType int

const (
	// ButtonsType_None is a representation of the C type GTK_BUTTONS_NONE.
	ButtonsType_None ButtonsType = 0
	// ButtonsType_Ok is a representation of the C type GTK_BUTTONS_OK.
	ButtonsType_Ok ButtonsType = 1
	// ButtonsType_Close is a representation of the C type GTK_BUTTONS_CLOSE.
	ButtonsType_Close ButtonsType = 2
	// ButtonsType_Cancel is a representation of the C type GTK_BUTTONS_CANCEL.
	ButtonsType_Cancel ButtonsType = 3
	// ButtonsType_YesNo is a representation of the C type GTK_BUTTONS_YES_NO.
	ButtonsType_YesNo ButtonsType = 4
	// ButtonsType_OkCancel is a representation of the C type GTK_BUTTONS_OK_CANCEL.
	ButtonsType_OkCancel ButtonsType = 5
)

// CellRendererAccelMode is a representation of the C type GtkCellRendererAccelMode.
type CellRendererAccelMode int

const (
	// CellRendererAccelMode_Gtk is a representation of the C type GTK_CELL_RENDERER_ACCEL_MODE_GTK.
	CellRendererAccelMode_Gtk CellRendererAccelMode = 0
	// CellRendererAccelMode_Other is a representation of the C type GTK_CELL_RENDERER_ACCEL_MODE_OTHER.
	CellRendererAccelMode_Other CellRendererAccelMode = 1
	// CellRendererAccelMode_ModifierTap is a representation of the C type GTK_CELL_RENDERER_ACCEL_MODE_MODIFIER_TAP.
	CellRendererAccelMode_ModifierTap CellRendererAccelMode = 2
)

// CellRendererMode is a representation of the C type GtkCellRendererMode.
type CellRendererMode int

const (
	// CellRendererMode_Inert is a representation of the C type GTK_CELL_RENDERER_MODE_INERT.
	CellRendererMode_Inert CellRendererMode = 0
	// CellRendererMode_Activatable is a representation of the C type GTK_CELL_RENDERER_MODE_ACTIVATABLE.
	CellRendererMode_Activatable CellRendererMode = 1
	// CellRendererMode_Editable is a representation of the C type GTK_CELL_RENDERER_MODE_EDITABLE.
	CellRendererMode_Editable CellRendererMode = 2
)

// CornerType is a representation of the C type GtkCornerType.
type CornerType int

const (
	// CornerType_TopLeft is a representation of the C type GTK_CORNER_TOP_LEFT.
	CornerType_TopLeft CornerType = 0
	// CornerType_BottomLeft is a representation of the C type GTK_CORNER_BOTTOM_LEFT.
	CornerType_BottomLeft CornerType = 1
	// CornerType_TopRight is a representation of the C type GTK_CORNER_TOP_RIGHT.
	CornerType_TopRight CornerType = 2
	// CornerType_BottomRight is a representation of the C type GTK_CORNER_BOTTOM_RIGHT.
	CornerType_BottomRight CornerType = 3
)

// CssProviderError is a representation of the C type GtkCssProviderError.
type CssProviderError int

const (
	// CssProviderError_Failed is a representation of the C type GTK_CSS_PROVIDER_ERROR_FAILED.
	CssProviderError_Failed CssProviderError = 0
	// CssProviderError_Syntax is a representation of the C type GTK_CSS_PROVIDER_ERROR_SYNTAX.
	CssProviderError_Syntax CssProviderError = 1
	// CssProviderError_Import is a representation of the C type GTK_CSS_PROVIDER_ERROR_IMPORT.
	CssProviderError_Import CssProviderError = 2
	// CssProviderError_Name is a representation of the C type GTK_CSS_PROVIDER_ERROR_NAME.
	CssProviderError_Name CssProviderError = 3
	// CssProviderError_Deprecated is a representation of the C type GTK_CSS_PROVIDER_ERROR_DEPRECATED.
	CssProviderError_Deprecated CssProviderError = 4
	// CssProviderError_UnknownValue is a representation of the C type GTK_CSS_PROVIDER_ERROR_UNKNOWN_VALUE.
	CssProviderError_UnknownValue CssProviderError = 5
)

// CssSectionType is a representation of the C type GtkCssSectionType.
//
// since 3.2
type CssSectionType int

const (
	// CssSectionType_Document is a representation of the C type GTK_CSS_SECTION_DOCUMENT.
	CssSectionType_Document CssSectionType = 0
	// CssSectionType_Import is a representation of the C type GTK_CSS_SECTION_IMPORT.
	CssSectionType_Import CssSectionType = 1
	// CssSectionType_ColorDefinition is a representation of the C type GTK_CSS_SECTION_COLOR_DEFINITION.
	CssSectionType_ColorDefinition CssSectionType = 2
	// CssSectionType_BindingSet is a representation of the C type GTK_CSS_SECTION_BINDING_SET.
	CssSectionType_BindingSet CssSectionType = 3
	// CssSectionType_Ruleset is a representation of the C type GTK_CSS_SECTION_RULESET.
	CssSectionType_Ruleset CssSectionType = 4
	// CssSectionType_Selector is a representation of the C type GTK_CSS_SECTION_SELECTOR.
	CssSectionType_Selector CssSectionType = 5
	// CssSectionType_Declaration is a representation of the C type GTK_CSS_SECTION_DECLARATION.
	CssSectionType_Declaration CssSectionType = 6
	// CssSectionType_Value is a representation of the C type GTK_CSS_SECTION_VALUE.
	CssSectionType_Value CssSectionType = 7
	// CssSectionType_Keyframes is a representation of the C type GTK_CSS_SECTION_KEYFRAMES.
	CssSectionType_Keyframes CssSectionType = 8
)

// DeleteType is a representation of the C type GtkDeleteType.
type DeleteType int

const (
	// DeleteType_Chars is a representation of the C type GTK_DELETE_CHARS.
	DeleteType_Chars DeleteType = 0
	// DeleteType_WordEnds is a representation of the C type GTK_DELETE_WORD_ENDS.
	DeleteType_WordEnds DeleteType = 1
	// DeleteType_Words is a representation of the C type GTK_DELETE_WORDS.
	DeleteType_Words DeleteType = 2
	// DeleteType_DisplayLines is a representation of the C type GTK_DELETE_DISPLAY_LINES.
	DeleteType_DisplayLines DeleteType = 3
	// DeleteType_DisplayLineEnds is a representation of the C type GTK_DELETE_DISPLAY_LINE_ENDS.
	DeleteType_DisplayLineEnds DeleteType = 4
	// DeleteType_ParagraphEnds is a representation of the C type GTK_DELETE_PARAGRAPH_ENDS.
	DeleteType_ParagraphEnds DeleteType = 5
	// DeleteType_Paragraphs is a representation of the C type GTK_DELETE_PARAGRAPHS.
	DeleteType_Paragraphs DeleteType = 6
	// DeleteType_Whitespace is a representation of the C type GTK_DELETE_WHITESPACE.
	DeleteType_Whitespace DeleteType = 7
)

// DirectionType is a representation of the C type GtkDirectionType.
type DirectionType int

const (
	// DirectionType_TabForward is a representation of the C type GTK_DIR_TAB_FORWARD.
	DirectionType_TabForward DirectionType = 0
	// DirectionType_TabBackward is a representation of the C type GTK_DIR_TAB_BACKWARD.
	DirectionType_TabBackward DirectionType = 1
	// DirectionType_Up is a representation of the C type GTK_DIR_UP.
	DirectionType_Up DirectionType = 2
	// DirectionType_Down is a representation of the C type GTK_DIR_DOWN.
	DirectionType_Down DirectionType = 3
	// DirectionType_Left is a representation of the C type GTK_DIR_LEFT.
	DirectionType_Left DirectionType = 4
	// DirectionType_Right is a representation of the C type GTK_DIR_RIGHT.
	DirectionType_Right DirectionType = 5
)

// DragResult is a representation of the C type GtkDragResult.
type DragResult int

const (
	// DragResult_Success is a representation of the C type GTK_DRAG_RESULT_SUCCESS.
	DragResult_Success DragResult = 0
	// DragResult_NoTarget is a representation of the C type GTK_DRAG_RESULT_NO_TARGET.
	DragResult_NoTarget DragResult = 1
	// DragResult_UserCancelled is a representation of the C type GTK_DRAG_RESULT_USER_CANCELLED.
	DragResult_UserCancelled DragResult = 2
	// DragResult_TimeoutExpired is a representation of the C type GTK_DRAG_RESULT_TIMEOUT_EXPIRED.
	DragResult_TimeoutExpired DragResult = 3
	// DragResult_GrabBroken is a representation of the C type GTK_DRAG_RESULT_GRAB_BROKEN.
	DragResult_GrabBroken DragResult = 4
	// DragResult_Error is a representation of the C type GTK_DRAG_RESULT_ERROR.
	DragResult_Error DragResult = 5
)

// EntryIconPosition is a representation of the C type GtkEntryIconPosition.
//
// since 2.16
type EntryIconPosition int

const (
	// EntryIconPosition_Primary is a representation of the C type GTK_ENTRY_ICON_PRIMARY.
	EntryIconPosition_Primary EntryIconPosition = 0
	// EntryIconPosition_Secondary is a representation of the C type GTK_ENTRY_ICON_SECONDARY.
	EntryIconPosition_Secondary EntryIconPosition = 1
)

// EventSequenceState is a representation of the C type GtkEventSequenceState.
//
// since 3.14
type EventSequenceState int

const (
	// EventSequenceState_None is a representation of the C type GTK_EVENT_SEQUENCE_NONE.
	EventSequenceState_None EventSequenceState = 0
	// EventSequenceState_Claimed is a representation of the C type GTK_EVENT_SEQUENCE_CLAIMED.
	EventSequenceState_Claimed EventSequenceState = 1
	// EventSequenceState_Denied is a representation of the C type GTK_EVENT_SEQUENCE_DENIED.
	EventSequenceState_Denied EventSequenceState = 2
)

// ExpanderStyle is a representation of the C type GtkExpanderStyle.
type ExpanderStyle int

const (
	// ExpanderStyle_Collapsed is a representation of the C type GTK_EXPANDER_COLLAPSED.
	ExpanderStyle_Collapsed ExpanderStyle = 0
	// ExpanderStyle_SemiCollapsed is a representation of the C type GTK_EXPANDER_SEMI_COLLAPSED.
	ExpanderStyle_SemiCollapsed ExpanderStyle = 1
	// ExpanderStyle_SemiExpanded is a representation of the C type GTK_EXPANDER_SEMI_EXPANDED.
	ExpanderStyle_SemiExpanded ExpanderStyle = 2
	// ExpanderStyle_Expanded is a representation of the C type GTK_EXPANDER_EXPANDED.
	ExpanderStyle_Expanded ExpanderStyle = 3
)

// FileChooserAction is a representation of the C type GtkFileChooserAction.
type FileChooserAction int

const (
	// FileChooserAction_Open is a representation of the C type GTK_FILE_CHOOSER_ACTION_OPEN.
	FileChooserAction_Open FileChooserAction = 0
	// FileChooserAction_Save is a representation of the C type GTK_FILE_CHOOSER_ACTION_SAVE.
	FileChooserAction_Save FileChooserAction = 1
	// FileChooserAction_SelectFolder is a representation of the C type GTK_FILE_CHOOSER_ACTION_SELECT_FOLDER.
	FileChooserAction_SelectFolder FileChooserAction = 2
	// FileChooserAction_CreateFolder is a representation of the C type GTK_FILE_CHOOSER_ACTION_CREATE_FOLDER.
	FileChooserAction_CreateFolder FileChooserAction = 3
)

// FileChooserConfirmation is a representation of the C type GtkFileChooserConfirmation.
//
// since 2.8
type FileChooserConfirmation int

const (
	// FileChooserConfirmation_Confirm is a representation of the C type GTK_FILE_CHOOSER_CONFIRMATION_CONFIRM.
	FileChooserConfirmation_Confirm FileChooserConfirmation = 0
	// FileChooserConfirmation_AcceptFilename is a representation of the C type GTK_FILE_CHOOSER_CONFIRMATION_ACCEPT_FILENAME.
	FileChooserConfirmation_AcceptFilename FileChooserConfirmation = 1
	// FileChooserConfirmation_SelectAgain is a representation of the C type GTK_FILE_CHOOSER_CONFIRMATION_SELECT_AGAIN.
	FileChooserConfirmation_SelectAgain FileChooserConfirmation = 2
)

// FileChooserError is a representation of the C type GtkFileChooserError.
type FileChooserError int

const (
	// FileChooserError_Nonexistent is a representation of the C type GTK_FILE_CHOOSER_ERROR_NONEXISTENT.
	FileChooserError_Nonexistent FileChooserError = 0
	// FileChooserError_BadFilename is a representation of the C type GTK_FILE_CHOOSER_ERROR_BAD_FILENAME.
	FileChooserError_BadFilename FileChooserError = 1
	// FileChooserError_AlreadyExists is a representation of the C type GTK_FILE_CHOOSER_ERROR_ALREADY_EXISTS.
	FileChooserError_AlreadyExists FileChooserError = 2
	// FileChooserError_IncompleteHostname is a representation of the C type GTK_FILE_CHOOSER_ERROR_INCOMPLETE_HOSTNAME.
	FileChooserError_IncompleteHostname FileChooserError = 3
)

// IMPreeditStyle is a representation of the C type GtkIMPreeditStyle.
type IMPreeditStyle int

const (
	// IMPreeditStyle_Nothing is a representation of the C type GTK_IM_PREEDIT_NOTHING.
	IMPreeditStyle_Nothing IMPreeditStyle = 0
	// IMPreeditStyle_Callback is a representation of the C type GTK_IM_PREEDIT_CALLBACK.
	IMPreeditStyle_Callback IMPreeditStyle = 1
	// IMPreeditStyle_None is a representation of the C type GTK_IM_PREEDIT_NONE.
	IMPreeditStyle_None IMPreeditStyle = 2
)

// IMStatusStyle is a representation of the C type GtkIMStatusStyle.
type IMStatusStyle int

const (
	// IMStatusStyle_Nothing is a representation of the C type GTK_IM_STATUS_NOTHING.
	IMStatusStyle_Nothing IMStatusStyle = 0
	// IMStatusStyle_Callback is a representation of the C type GTK_IM_STATUS_CALLBACK.
	IMStatusStyle_Callback IMStatusStyle = 1
	// IMStatusStyle_None is a representation of the C type GTK_IM_STATUS_NONE.
	IMStatusStyle_None IMStatusStyle = 2
)

// IconSize is a representation of the C type GtkIconSize.
type IconSize int

const (
	// IconSize_Invalid is a representation of the C type GTK_ICON_SIZE_INVALID.
	IconSize_Invalid IconSize = 0
	// IconSize_Menu is a representation of the C type GTK_ICON_SIZE_MENU.
	IconSize_Menu IconSize = 1
	// IconSize_SmallToolbar is a representation of the C type GTK_ICON_SIZE_SMALL_TOOLBAR.
	IconSize_SmallToolbar IconSize = 2
	// IconSize_LargeToolbar is a representation of the C type GTK_ICON_SIZE_LARGE_TOOLBAR.
	IconSize_LargeToolbar IconSize = 3
	// IconSize_Button is a representation of the C type GTK_ICON_SIZE_BUTTON.
	IconSize_Button IconSize = 4
	// IconSize_Dnd is a representation of the C type GTK_ICON_SIZE_DND.
	IconSize_Dnd IconSize = 5
	// IconSize_Dialog is a representation of the C type GTK_ICON_SIZE_DIALOG.
	IconSize_Dialog IconSize = 6
)

// IconThemeError is a representation of the C type GtkIconThemeError.
type IconThemeError int

const (
	// IconThemeError_NotFound is a representation of the C type GTK_ICON_THEME_NOT_FOUND.
	IconThemeError_NotFound IconThemeError = 0
	// IconThemeError_Failed is a representation of the C type GTK_ICON_THEME_FAILED.
	IconThemeError_Failed IconThemeError = 1
)

// IconViewDropPosition is a representation of the C type GtkIconViewDropPosition.
type IconViewDropPosition int

const (
	// IconViewDropPosition_NoDrop is a representation of the C type GTK_ICON_VIEW_NO_DROP.
	IconViewDropPosition_NoDrop IconViewDropPosition = 0
	// IconViewDropPosition_DropInto is a representation of the C type GTK_ICON_VIEW_DROP_INTO.
	IconViewDropPosition_DropInto IconViewDropPosition = 1
	// IconViewDropPosition_DropLeft is a representation of the C type GTK_ICON_VIEW_DROP_LEFT.
	IconViewDropPosition_DropLeft IconViewDropPosition = 2
	// IconViewDropPosition_DropRight is a representation of the C type GTK_ICON_VIEW_DROP_RIGHT.
	IconViewDropPosition_DropRight IconViewDropPosition = 3
	// IconViewDropPosition_DropAbove is a representation of the C type GTK_ICON_VIEW_DROP_ABOVE.
	IconViewDropPosition_DropAbove IconViewDropPosition = 4
	// IconViewDropPosition_DropBelow is a representation of the C type GTK_ICON_VIEW_DROP_BELOW.
	IconViewDropPosition_DropBelow IconViewDropPosition = 5
)

// ImageType is a representation of the C type GtkImageType.
type ImageType int

const (
	// ImageType_Empty is a representation of the C type GTK_IMAGE_EMPTY.
	ImageType_Empty ImageType = 0
	// ImageType_Pixbuf is a representation of the C type GTK_IMAGE_PIXBUF.
	ImageType_Pixbuf ImageType = 1
	// ImageType_Stock is a representation of the C type GTK_IMAGE_STOCK.
	ImageType_Stock ImageType = 2
	// ImageType_IconSet is a representation of the C type GTK_IMAGE_ICON_SET.
	ImageType_IconSet ImageType = 3
	// ImageType_Animation is a representation of the C type GTK_IMAGE_ANIMATION.
	ImageType_Animation ImageType = 4
	// ImageType_IconName is a representation of the C type GTK_IMAGE_ICON_NAME.
	ImageType_IconName ImageType = 5
	// ImageType_Gicon is a representation of the C type GTK_IMAGE_GICON.
	ImageType_Gicon ImageType = 6
	// ImageType_Surface is a representation of the C type GTK_IMAGE_SURFACE.
	ImageType_Surface ImageType = 7
)

// InputPurpose is a representation of the C type GtkInputPurpose.
//
// since 3.6
type InputPurpose int

const (
	// InputPurpose_FreeForm is a representation of the C type GTK_INPUT_PURPOSE_FREE_FORM.
	InputPurpose_FreeForm InputPurpose = 0
	// InputPurpose_Alpha is a representation of the C type GTK_INPUT_PURPOSE_ALPHA.
	InputPurpose_Alpha InputPurpose = 1
	// InputPurpose_Digits is a representation of the C type GTK_INPUT_PURPOSE_DIGITS.
	InputPurpose_Digits InputPurpose = 2
	// InputPurpose_Number is a representation of the C type GTK_INPUT_PURPOSE_NUMBER.
	InputPurpose_Number InputPurpose = 3
	// InputPurpose_Phone is a representation of the C type GTK_INPUT_PURPOSE_PHONE.
	InputPurpose_Phone InputPurpose = 4
	// InputPurpose_Url is a representation of the C type GTK_INPUT_PURPOSE_URL.
	InputPurpose_Url InputPurpose = 5
	// InputPurpose_Email is a representation of the C type GTK_INPUT_PURPOSE_EMAIL.
	InputPurpose_Email InputPurpose = 6
	// InputPurpose_Name is a representation of the C type GTK_INPUT_PURPOSE_NAME.
	InputPurpose_Name InputPurpose = 7
	// InputPurpose_Password is a representation of the C type GTK_INPUT_PURPOSE_PASSWORD.
	InputPurpose_Password InputPurpose = 8
	// InputPurpose_Pin is a representation of the C type GTK_INPUT_PURPOSE_PIN.
	InputPurpose_Pin InputPurpose = 9
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

// LevelBarMode is a representation of the C type GtkLevelBarMode.
//
// since 3.6
type LevelBarMode int

const (
	// LevelBarMode_Continuous is a representation of the C type GTK_LEVEL_BAR_MODE_CONTINUOUS.
	LevelBarMode_Continuous LevelBarMode = 0
	// LevelBarMode_Discrete is a representation of the C type GTK_LEVEL_BAR_MODE_DISCRETE.
	LevelBarMode_Discrete LevelBarMode = 1
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

// MenuDirectionType is a representation of the C type GtkMenuDirectionType.
type MenuDirectionType int

const (
	// MenuDirectionType_Parent is a representation of the C type GTK_MENU_DIR_PARENT.
	MenuDirectionType_Parent MenuDirectionType = 0
	// MenuDirectionType_Child is a representation of the C type GTK_MENU_DIR_CHILD.
	MenuDirectionType_Child MenuDirectionType = 1
	// MenuDirectionType_Next is a representation of the C type GTK_MENU_DIR_NEXT.
	MenuDirectionType_Next MenuDirectionType = 2
	// MenuDirectionType_Prev is a representation of the C type GTK_MENU_DIR_PREV.
	MenuDirectionType_Prev MenuDirectionType = 3
)

// MessageType is a representation of the C type GtkMessageType.
type MessageType int

const (
	// MessageType_Info is a representation of the C type GTK_MESSAGE_INFO.
	MessageType_Info MessageType = 0
	// MessageType_Warning is a representation of the C type GTK_MESSAGE_WARNING.
	MessageType_Warning MessageType = 1
	// MessageType_Question is a representation of the C type GTK_MESSAGE_QUESTION.
	MessageType_Question MessageType = 2
	// MessageType_Error is a representation of the C type GTK_MESSAGE_ERROR.
	MessageType_Error MessageType = 3
	// MessageType_Other is a representation of the C type GTK_MESSAGE_OTHER.
	MessageType_Other MessageType = 4
)

// MovementStep is a representation of the C type GtkMovementStep.
type MovementStep int

const (
	// MovementStep_LogicalPositions is a representation of the C type GTK_MOVEMENT_LOGICAL_POSITIONS.
	MovementStep_LogicalPositions MovementStep = 0
	// MovementStep_VisualPositions is a representation of the C type GTK_MOVEMENT_VISUAL_POSITIONS.
	MovementStep_VisualPositions MovementStep = 1
	// MovementStep_Words is a representation of the C type GTK_MOVEMENT_WORDS.
	MovementStep_Words MovementStep = 2
	// MovementStep_DisplayLines is a representation of the C type GTK_MOVEMENT_DISPLAY_LINES.
	MovementStep_DisplayLines MovementStep = 3
	// MovementStep_DisplayLineEnds is a representation of the C type GTK_MOVEMENT_DISPLAY_LINE_ENDS.
	MovementStep_DisplayLineEnds MovementStep = 4
	// MovementStep_Paragraphs is a representation of the C type GTK_MOVEMENT_PARAGRAPHS.
	MovementStep_Paragraphs MovementStep = 5
	// MovementStep_ParagraphEnds is a representation of the C type GTK_MOVEMENT_PARAGRAPH_ENDS.
	MovementStep_ParagraphEnds MovementStep = 6
	// MovementStep_Pages is a representation of the C type GTK_MOVEMENT_PAGES.
	MovementStep_Pages MovementStep = 7
	// MovementStep_BufferEnds is a representation of the C type GTK_MOVEMENT_BUFFER_ENDS.
	MovementStep_BufferEnds MovementStep = 8
	// MovementStep_HorizontalPages is a representation of the C type GTK_MOVEMENT_HORIZONTAL_PAGES.
	MovementStep_HorizontalPages MovementStep = 9
)

// NotebookTab is a representation of the C type GtkNotebookTab.
type NotebookTab int

const (
	// NotebookTab_First is a representation of the C type GTK_NOTEBOOK_TAB_FIRST.
	NotebookTab_First NotebookTab = 0
	// NotebookTab_Last is a representation of the C type GTK_NOTEBOOK_TAB_LAST.
	NotebookTab_Last NotebookTab = 1
)

// NumberUpLayout is a representation of the C type GtkNumberUpLayout.
type NumberUpLayout int

const (
	// NumberUpLayout_Lrtb is a representation of the C type GTK_NUMBER_UP_LAYOUT_LEFT_TO_RIGHT_TOP_TO_BOTTOM.
	NumberUpLayout_Lrtb NumberUpLayout = 0
	// NumberUpLayout_Lrbt is a representation of the C type GTK_NUMBER_UP_LAYOUT_LEFT_TO_RIGHT_BOTTOM_TO_TOP.
	NumberUpLayout_Lrbt NumberUpLayout = 1
	// NumberUpLayout_Rltb is a representation of the C type GTK_NUMBER_UP_LAYOUT_RIGHT_TO_LEFT_TOP_TO_BOTTOM.
	NumberUpLayout_Rltb NumberUpLayout = 2
	// NumberUpLayout_Rlbt is a representation of the C type GTK_NUMBER_UP_LAYOUT_RIGHT_TO_LEFT_BOTTOM_TO_TOP.
	NumberUpLayout_Rlbt NumberUpLayout = 3
	// NumberUpLayout_Tblr is a representation of the C type GTK_NUMBER_UP_LAYOUT_TOP_TO_BOTTOM_LEFT_TO_RIGHT.
	NumberUpLayout_Tblr NumberUpLayout = 4
	// NumberUpLayout_Tbrl is a representation of the C type GTK_NUMBER_UP_LAYOUT_TOP_TO_BOTTOM_RIGHT_TO_LEFT.
	NumberUpLayout_Tbrl NumberUpLayout = 5
	// NumberUpLayout_Btlr is a representation of the C type GTK_NUMBER_UP_LAYOUT_BOTTOM_TO_TOP_LEFT_TO_RIGHT.
	NumberUpLayout_Btlr NumberUpLayout = 6
	// NumberUpLayout_Btrl is a representation of the C type GTK_NUMBER_UP_LAYOUT_BOTTOM_TO_TOP_RIGHT_TO_LEFT.
	NumberUpLayout_Btrl NumberUpLayout = 7
)

// Orientation is a representation of the C type GtkOrientation.
type Orientation int

const (
	// Orientation_Horizontal is a representation of the C type GTK_ORIENTATION_HORIZONTAL.
	Orientation_Horizontal Orientation = 0
	// Orientation_Vertical is a representation of the C type GTK_ORIENTATION_VERTICAL.
	Orientation_Vertical Orientation = 1
)

// PackDirection is a representation of the C type GtkPackDirection.
type PackDirection int

const (
	// PackDirection_Ltr is a representation of the C type GTK_PACK_DIRECTION_LTR.
	PackDirection_Ltr PackDirection = 0
	// PackDirection_Rtl is a representation of the C type GTK_PACK_DIRECTION_RTL.
	PackDirection_Rtl PackDirection = 1
	// PackDirection_Ttb is a representation of the C type GTK_PACK_DIRECTION_TTB.
	PackDirection_Ttb PackDirection = 2
	// PackDirection_Btt is a representation of the C type GTK_PACK_DIRECTION_BTT.
	PackDirection_Btt PackDirection = 3
)

// PackType is a representation of the C type GtkPackType.
type PackType int

const (
	// PackType_Start is a representation of the C type GTK_PACK_START.
	PackType_Start PackType = 0
	// PackType_End is a representation of the C type GTK_PACK_END.
	PackType_End PackType = 1
)

// PadActionType is a representation of the C type GtkPadActionType.
type PadActionType int

const (
	// PadActionType_Button is a representation of the C type GTK_PAD_ACTION_BUTTON.
	PadActionType_Button PadActionType = 0
	// PadActionType_Ring is a representation of the C type GTK_PAD_ACTION_RING.
	PadActionType_Ring PadActionType = 1
	// PadActionType_Strip is a representation of the C type GTK_PAD_ACTION_STRIP.
	PadActionType_Strip PadActionType = 2
)

// PageOrientation is a representation of the C type GtkPageOrientation.
type PageOrientation int

const (
	// PageOrientation_Portrait is a representation of the C type GTK_PAGE_ORIENTATION_PORTRAIT.
	PageOrientation_Portrait PageOrientation = 0
	// PageOrientation_Landscape is a representation of the C type GTK_PAGE_ORIENTATION_LANDSCAPE.
	PageOrientation_Landscape PageOrientation = 1
	// PageOrientation_ReversePortrait is a representation of the C type GTK_PAGE_ORIENTATION_REVERSE_PORTRAIT.
	PageOrientation_ReversePortrait PageOrientation = 2
	// PageOrientation_ReverseLandscape is a representation of the C type GTK_PAGE_ORIENTATION_REVERSE_LANDSCAPE.
	PageOrientation_ReverseLandscape PageOrientation = 3
)

// PageSet is a representation of the C type GtkPageSet.
type PageSet int

const (
	// PageSet_All is a representation of the C type GTK_PAGE_SET_ALL.
	PageSet_All PageSet = 0
	// PageSet_Even is a representation of the C type GTK_PAGE_SET_EVEN.
	PageSet_Even PageSet = 1
	// PageSet_Odd is a representation of the C type GTK_PAGE_SET_ODD.
	PageSet_Odd PageSet = 2
)

// PanDirection is a representation of the C type GtkPanDirection.
//
// since 3.14
type PanDirection int

const (
	// PanDirection_Left is a representation of the C type GTK_PAN_DIRECTION_LEFT.
	PanDirection_Left PanDirection = 0
	// PanDirection_Right is a representation of the C type GTK_PAN_DIRECTION_RIGHT.
	PanDirection_Right PanDirection = 1
	// PanDirection_Up is a representation of the C type GTK_PAN_DIRECTION_UP.
	PanDirection_Up PanDirection = 2
	// PanDirection_Down is a representation of the C type GTK_PAN_DIRECTION_DOWN.
	PanDirection_Down PanDirection = 3
)

// PathPriorityType is a representation of the C type GtkPathPriorityType.
type PathPriorityType int

const (
	// PathPriorityType_Lowest is a representation of the C type GTK_PATH_PRIO_LOWEST.
	PathPriorityType_Lowest PathPriorityType = 0
	// PathPriorityType_Gtk is a representation of the C type GTK_PATH_PRIO_GTK.
	PathPriorityType_Gtk PathPriorityType = 4
	// PathPriorityType_Application is a representation of the C type GTK_PATH_PRIO_APPLICATION.
	PathPriorityType_Application PathPriorityType = 8
	// PathPriorityType_Theme is a representation of the C type GTK_PATH_PRIO_THEME.
	PathPriorityType_Theme PathPriorityType = 10
	// PathPriorityType_Rc is a representation of the C type GTK_PATH_PRIO_RC.
	PathPriorityType_Rc PathPriorityType = 12
	// PathPriorityType_Highest is a representation of the C type GTK_PATH_PRIO_HIGHEST.
	PathPriorityType_Highest PathPriorityType = 15
)

// PathType is a representation of the C type GtkPathType.
type PathType int

const (
	// PathType_Widget is a representation of the C type GTK_PATH_WIDGET.
	PathType_Widget PathType = 0
	// PathType_WidgetClass is a representation of the C type GTK_PATH_WIDGET_CLASS.
	PathType_WidgetClass PathType = 1
	// PathType_Class is a representation of the C type GTK_PATH_CLASS.
	PathType_Class PathType = 2
)

// PolicyType is a representation of the C type GtkPolicyType.
type PolicyType int

const (
	// PolicyType_Always is a representation of the C type GTK_POLICY_ALWAYS.
	PolicyType_Always PolicyType = 0
	// PolicyType_Automatic is a representation of the C type GTK_POLICY_AUTOMATIC.
	PolicyType_Automatic PolicyType = 1
	// PolicyType_Never is a representation of the C type GTK_POLICY_NEVER.
	PolicyType_Never PolicyType = 2
	// PolicyType_External is a representation of the C type GTK_POLICY_EXTERNAL.
	PolicyType_External PolicyType = 3
)

// PopoverConstraint is a representation of the C type GtkPopoverConstraint.
//
// since 3.20
type PopoverConstraint int

const (
	// PopoverConstraint_None is a representation of the C type GTK_POPOVER_CONSTRAINT_NONE.
	PopoverConstraint_None PopoverConstraint = 0
	// PopoverConstraint_Window is a representation of the C type GTK_POPOVER_CONSTRAINT_WINDOW.
	PopoverConstraint_Window PopoverConstraint = 1
)

// PositionType is a representation of the C type GtkPositionType.
type PositionType int

const (
	// PositionType_Left is a representation of the C type GTK_POS_LEFT.
	PositionType_Left PositionType = 0
	// PositionType_Right is a representation of the C type GTK_POS_RIGHT.
	PositionType_Right PositionType = 1
	// PositionType_Top is a representation of the C type GTK_POS_TOP.
	PositionType_Top PositionType = 2
	// PositionType_Bottom is a representation of the C type GTK_POS_BOTTOM.
	PositionType_Bottom PositionType = 3
)

// PrintDuplex is a representation of the C type GtkPrintDuplex.
type PrintDuplex int

const (
	// PrintDuplex_Simplex is a representation of the C type GTK_PRINT_DUPLEX_SIMPLEX.
	PrintDuplex_Simplex PrintDuplex = 0
	// PrintDuplex_Horizontal is a representation of the C type GTK_PRINT_DUPLEX_HORIZONTAL.
	PrintDuplex_Horizontal PrintDuplex = 1
	// PrintDuplex_Vertical is a representation of the C type GTK_PRINT_DUPLEX_VERTICAL.
	PrintDuplex_Vertical PrintDuplex = 2
)

// PrintError is a representation of the C type GtkPrintError.
type PrintError int

const (
	// PrintError_General is a representation of the C type GTK_PRINT_ERROR_GENERAL.
	PrintError_General PrintError = 0
	// PrintError_InternalError is a representation of the C type GTK_PRINT_ERROR_INTERNAL_ERROR.
	PrintError_InternalError PrintError = 1
	// PrintError_Nomem is a representation of the C type GTK_PRINT_ERROR_NOMEM.
	PrintError_Nomem PrintError = 2
	// PrintError_InvalidFile is a representation of the C type GTK_PRINT_ERROR_INVALID_FILE.
	PrintError_InvalidFile PrintError = 3
)

// PrintOperationAction is a representation of the C type GtkPrintOperationAction.
type PrintOperationAction int

const (
	// PrintOperationAction_PrintDialog is a representation of the C type GTK_PRINT_OPERATION_ACTION_PRINT_DIALOG.
	PrintOperationAction_PrintDialog PrintOperationAction = 0
	// PrintOperationAction_Print is a representation of the C type GTK_PRINT_OPERATION_ACTION_PRINT.
	PrintOperationAction_Print PrintOperationAction = 1
	// PrintOperationAction_Preview is a representation of the C type GTK_PRINT_OPERATION_ACTION_PREVIEW.
	PrintOperationAction_Preview PrintOperationAction = 2
	// PrintOperationAction_Export is a representation of the C type GTK_PRINT_OPERATION_ACTION_EXPORT.
	PrintOperationAction_Export PrintOperationAction = 3
)

// PrintOperationResult is a representation of the C type GtkPrintOperationResult.
type PrintOperationResult int

const (
	// PrintOperationResult_Error is a representation of the C type GTK_PRINT_OPERATION_RESULT_ERROR.
	PrintOperationResult_Error PrintOperationResult = 0
	// PrintOperationResult_Apply is a representation of the C type GTK_PRINT_OPERATION_RESULT_APPLY.
	PrintOperationResult_Apply PrintOperationResult = 1
	// PrintOperationResult_Cancel is a representation of the C type GTK_PRINT_OPERATION_RESULT_CANCEL.
	PrintOperationResult_Cancel PrintOperationResult = 2
	// PrintOperationResult_InProgress is a representation of the C type GTK_PRINT_OPERATION_RESULT_IN_PROGRESS.
	PrintOperationResult_InProgress PrintOperationResult = 3
)

// PrintPages is a representation of the C type GtkPrintPages.
type PrintPages int

const (
	// PrintPages_All is a representation of the C type GTK_PRINT_PAGES_ALL.
	PrintPages_All PrintPages = 0
	// PrintPages_Current is a representation of the C type GTK_PRINT_PAGES_CURRENT.
	PrintPages_Current PrintPages = 1
	// PrintPages_Ranges is a representation of the C type GTK_PRINT_PAGES_RANGES.
	PrintPages_Ranges PrintPages = 2
	// PrintPages_Selection is a representation of the C type GTK_PRINT_PAGES_SELECTION.
	PrintPages_Selection PrintPages = 3
)

// PrintQuality is a representation of the C type GtkPrintQuality.
type PrintQuality int

const (
	// PrintQuality_Low is a representation of the C type GTK_PRINT_QUALITY_LOW.
	PrintQuality_Low PrintQuality = 0
	// PrintQuality_Normal is a representation of the C type GTK_PRINT_QUALITY_NORMAL.
	PrintQuality_Normal PrintQuality = 1
	// PrintQuality_High is a representation of the C type GTK_PRINT_QUALITY_HIGH.
	PrintQuality_High PrintQuality = 2
	// PrintQuality_Draft is a representation of the C type GTK_PRINT_QUALITY_DRAFT.
	PrintQuality_Draft PrintQuality = 3
)

// PrintStatus is a representation of the C type GtkPrintStatus.
type PrintStatus int

const (
	// PrintStatus_Initial is a representation of the C type GTK_PRINT_STATUS_INITIAL.
	PrintStatus_Initial PrintStatus = 0
	// PrintStatus_Preparing is a representation of the C type GTK_PRINT_STATUS_PREPARING.
	PrintStatus_Preparing PrintStatus = 1
	// PrintStatus_GeneratingData is a representation of the C type GTK_PRINT_STATUS_GENERATING_DATA.
	PrintStatus_GeneratingData PrintStatus = 2
	// PrintStatus_SendingData is a representation of the C type GTK_PRINT_STATUS_SENDING_DATA.
	PrintStatus_SendingData PrintStatus = 3
	// PrintStatus_Pending is a representation of the C type GTK_PRINT_STATUS_PENDING.
	PrintStatus_Pending PrintStatus = 4
	// PrintStatus_PendingIssue is a representation of the C type GTK_PRINT_STATUS_PENDING_ISSUE.
	PrintStatus_PendingIssue PrintStatus = 5
	// PrintStatus_Printing is a representation of the C type GTK_PRINT_STATUS_PRINTING.
	PrintStatus_Printing PrintStatus = 6
	// PrintStatus_Finished is a representation of the C type GTK_PRINT_STATUS_FINISHED.
	PrintStatus_Finished PrintStatus = 7
	// PrintStatus_FinishedAborted is a representation of the C type GTK_PRINT_STATUS_FINISHED_ABORTED.
	PrintStatus_FinishedAborted PrintStatus = 8
)

// PropagationPhase is a representation of the C type GtkPropagationPhase.
//
// since 3.14
type PropagationPhase int

const (
	// PropagationPhase_None is a representation of the C type GTK_PHASE_NONE.
	PropagationPhase_None PropagationPhase = 0
	// PropagationPhase_Capture is a representation of the C type GTK_PHASE_CAPTURE.
	PropagationPhase_Capture PropagationPhase = 1
	// PropagationPhase_Bubble is a representation of the C type GTK_PHASE_BUBBLE.
	PropagationPhase_Bubble PropagationPhase = 2
	// PropagationPhase_Target is a representation of the C type GTK_PHASE_TARGET.
	PropagationPhase_Target PropagationPhase = 3
)

// RcTokenType is a representation of the C type GtkRcTokenType.
type RcTokenType int

const (
	// RcTokenType_Invalid is a representation of the C type GTK_RC_TOKEN_INVALID.
	RcTokenType_Invalid RcTokenType = 270
	// RcTokenType_Include is a representation of the C type GTK_RC_TOKEN_INCLUDE.
	RcTokenType_Include RcTokenType = 271
	// RcTokenType_Normal is a representation of the C type GTK_RC_TOKEN_NORMAL.
	RcTokenType_Normal RcTokenType = 272
	// RcTokenType_Active is a representation of the C type GTK_RC_TOKEN_ACTIVE.
	RcTokenType_Active RcTokenType = 273
	// RcTokenType_Prelight is a representation of the C type GTK_RC_TOKEN_PRELIGHT.
	RcTokenType_Prelight RcTokenType = 274
	// RcTokenType_Selected is a representation of the C type GTK_RC_TOKEN_SELECTED.
	RcTokenType_Selected RcTokenType = 275
	// RcTokenType_Insensitive is a representation of the C type GTK_RC_TOKEN_INSENSITIVE.
	RcTokenType_Insensitive RcTokenType = 276
	// RcTokenType_Fg is a representation of the C type GTK_RC_TOKEN_FG.
	RcTokenType_Fg RcTokenType = 277
	// RcTokenType_Bg is a representation of the C type GTK_RC_TOKEN_BG.
	RcTokenType_Bg RcTokenType = 278
	// RcTokenType_Text is a representation of the C type GTK_RC_TOKEN_TEXT.
	RcTokenType_Text RcTokenType = 279
	// RcTokenType_Base is a representation of the C type GTK_RC_TOKEN_BASE.
	RcTokenType_Base RcTokenType = 280
	// RcTokenType_Xthickness is a representation of the C type GTK_RC_TOKEN_XTHICKNESS.
	RcTokenType_Xthickness RcTokenType = 281
	// RcTokenType_Ythickness is a representation of the C type GTK_RC_TOKEN_YTHICKNESS.
	RcTokenType_Ythickness RcTokenType = 282
	// RcTokenType_Font is a representation of the C type GTK_RC_TOKEN_FONT.
	RcTokenType_Font RcTokenType = 283
	// RcTokenType_Fontset is a representation of the C type GTK_RC_TOKEN_FONTSET.
	RcTokenType_Fontset RcTokenType = 284
	// RcTokenType_FontName is a representation of the C type GTK_RC_TOKEN_FONT_NAME.
	RcTokenType_FontName RcTokenType = 285
	// RcTokenType_BgPixmap is a representation of the C type GTK_RC_TOKEN_BG_PIXMAP.
	RcTokenType_BgPixmap RcTokenType = 286
	// RcTokenType_PixmapPath is a representation of the C type GTK_RC_TOKEN_PIXMAP_PATH.
	RcTokenType_PixmapPath RcTokenType = 287
	// RcTokenType_Style is a representation of the C type GTK_RC_TOKEN_STYLE.
	RcTokenType_Style RcTokenType = 288
	// RcTokenType_Binding is a representation of the C type GTK_RC_TOKEN_BINDING.
	RcTokenType_Binding RcTokenType = 289
	// RcTokenType_Bind is a representation of the C type GTK_RC_TOKEN_BIND.
	RcTokenType_Bind RcTokenType = 290
	// RcTokenType_Widget is a representation of the C type GTK_RC_TOKEN_WIDGET.
	RcTokenType_Widget RcTokenType = 291
	// RcTokenType_WidgetClass is a representation of the C type GTK_RC_TOKEN_WIDGET_CLASS.
	RcTokenType_WidgetClass RcTokenType = 292
	// RcTokenType_Class is a representation of the C type GTK_RC_TOKEN_CLASS.
	RcTokenType_Class RcTokenType = 293
	// RcTokenType_Lowest is a representation of the C type GTK_RC_TOKEN_LOWEST.
	RcTokenType_Lowest RcTokenType = 294
	// RcTokenType_Gtk is a representation of the C type GTK_RC_TOKEN_GTK.
	RcTokenType_Gtk RcTokenType = 295
	// RcTokenType_Application is a representation of the C type GTK_RC_TOKEN_APPLICATION.
	RcTokenType_Application RcTokenType = 296
	// RcTokenType_Theme is a representation of the C type GTK_RC_TOKEN_THEME.
	RcTokenType_Theme RcTokenType = 297
	// RcTokenType_Rc is a representation of the C type GTK_RC_TOKEN_RC.
	RcTokenType_Rc RcTokenType = 298
	// RcTokenType_Highest is a representation of the C type GTK_RC_TOKEN_HIGHEST.
	RcTokenType_Highest RcTokenType = 299
	// RcTokenType_Engine is a representation of the C type GTK_RC_TOKEN_ENGINE.
	RcTokenType_Engine RcTokenType = 300
	// RcTokenType_ModulePath is a representation of the C type GTK_RC_TOKEN_MODULE_PATH.
	RcTokenType_ModulePath RcTokenType = 301
	// RcTokenType_ImModulePath is a representation of the C type GTK_RC_TOKEN_IM_MODULE_PATH.
	RcTokenType_ImModulePath RcTokenType = 302
	// RcTokenType_ImModuleFile is a representation of the C type GTK_RC_TOKEN_IM_MODULE_FILE.
	RcTokenType_ImModuleFile RcTokenType = 303
	// RcTokenType_Stock is a representation of the C type GTK_RC_TOKEN_STOCK.
	RcTokenType_Stock RcTokenType = 304
	// RcTokenType_Ltr is a representation of the C type GTK_RC_TOKEN_LTR.
	RcTokenType_Ltr RcTokenType = 305
	// RcTokenType_Rtl is a representation of the C type GTK_RC_TOKEN_RTL.
	RcTokenType_Rtl RcTokenType = 306
	// RcTokenType_Color is a representation of the C type GTK_RC_TOKEN_COLOR.
	RcTokenType_Color RcTokenType = 307
	// RcTokenType_Unbind is a representation of the C type GTK_RC_TOKEN_UNBIND.
	RcTokenType_Unbind RcTokenType = 308
	// RcTokenType_Last is a representation of the C type GTK_RC_TOKEN_LAST.
	RcTokenType_Last RcTokenType = 309
)

// RecentChooserError is a representation of the C type GtkRecentChooserError.
//
// since 2.10
type RecentChooserError int

const (
	// RecentChooserError_NotFound is a representation of the C type GTK_RECENT_CHOOSER_ERROR_NOT_FOUND.
	RecentChooserError_NotFound RecentChooserError = 0
	// RecentChooserError_InvalidUri is a representation of the C type GTK_RECENT_CHOOSER_ERROR_INVALID_URI.
	RecentChooserError_InvalidUri RecentChooserError = 1
)

// RecentManagerError is a representation of the C type GtkRecentManagerError.
//
// since 2.10
type RecentManagerError int

const (
	// RecentManagerError_NotFound is a representation of the C type GTK_RECENT_MANAGER_ERROR_NOT_FOUND.
	RecentManagerError_NotFound RecentManagerError = 0
	// RecentManagerError_InvalidUri is a representation of the C type GTK_RECENT_MANAGER_ERROR_INVALID_URI.
	RecentManagerError_InvalidUri RecentManagerError = 1
	// RecentManagerError_InvalidEncoding is a representation of the C type GTK_RECENT_MANAGER_ERROR_INVALID_ENCODING.
	RecentManagerError_InvalidEncoding RecentManagerError = 2
	// RecentManagerError_NotRegistered is a representation of the C type GTK_RECENT_MANAGER_ERROR_NOT_REGISTERED.
	RecentManagerError_NotRegistered RecentManagerError = 3
	// RecentManagerError_Read is a representation of the C type GTK_RECENT_MANAGER_ERROR_READ.
	RecentManagerError_Read RecentManagerError = 4
	// RecentManagerError_Write is a representation of the C type GTK_RECENT_MANAGER_ERROR_WRITE.
	RecentManagerError_Write RecentManagerError = 5
	// RecentManagerError_Unknown is a representation of the C type GTK_RECENT_MANAGER_ERROR_UNKNOWN.
	RecentManagerError_Unknown RecentManagerError = 6
)

// RecentSortType is a representation of the C type GtkRecentSortType.
//
// since 2.10
type RecentSortType int

const (
	// RecentSortType_None is a representation of the C type GTK_RECENT_SORT_NONE.
	RecentSortType_None RecentSortType = 0
	// RecentSortType_Mru is a representation of the C type GTK_RECENT_SORT_MRU.
	RecentSortType_Mru RecentSortType = 1
	// RecentSortType_Lru is a representation of the C type GTK_RECENT_SORT_LRU.
	RecentSortType_Lru RecentSortType = 2
	// RecentSortType_Custom is a representation of the C type GTK_RECENT_SORT_CUSTOM.
	RecentSortType_Custom RecentSortType = 3
)

// ReliefStyle is a representation of the C type GtkReliefStyle.
type ReliefStyle int

const (
	// ReliefStyle_Normal is a representation of the C type GTK_RELIEF_NORMAL.
	ReliefStyle_Normal ReliefStyle = 0
	// ReliefStyle_Half is a representation of the C type GTK_RELIEF_HALF.
	ReliefStyle_Half ReliefStyle = 1
	// ReliefStyle_None is a representation of the C type GTK_RELIEF_NONE.
	ReliefStyle_None ReliefStyle = 2
)

// ResizeMode is a representation of the C type GtkResizeMode.
type ResizeMode int

const (
	// ResizeMode_Parent is a representation of the C type GTK_RESIZE_PARENT.
	ResizeMode_Parent ResizeMode = 0
	// ResizeMode_Queue is a representation of the C type GTK_RESIZE_QUEUE.
	ResizeMode_Queue ResizeMode = 1
	// ResizeMode_Immediate is a representation of the C type GTK_RESIZE_IMMEDIATE.
	ResizeMode_Immediate ResizeMode = 2
)

// ResponseType is a representation of the C type GtkResponseType.
type ResponseType int

const (
	// ResponseType_None is a representation of the C type GTK_RESPONSE_NONE.
	ResponseType_None ResponseType = -1
	// ResponseType_Reject is a representation of the C type GTK_RESPONSE_REJECT.
	ResponseType_Reject ResponseType = -2
	// ResponseType_Accept is a representation of the C type GTK_RESPONSE_ACCEPT.
	ResponseType_Accept ResponseType = -3
	// ResponseType_DeleteEvent is a representation of the C type GTK_RESPONSE_DELETE_EVENT.
	ResponseType_DeleteEvent ResponseType = -4
	// ResponseType_Ok is a representation of the C type GTK_RESPONSE_OK.
	ResponseType_Ok ResponseType = -5
	// ResponseType_Cancel is a representation of the C type GTK_RESPONSE_CANCEL.
	ResponseType_Cancel ResponseType = -6
	// ResponseType_Close is a representation of the C type GTK_RESPONSE_CLOSE.
	ResponseType_Close ResponseType = -7
	// ResponseType_Yes is a representation of the C type GTK_RESPONSE_YES.
	ResponseType_Yes ResponseType = -8
	// ResponseType_No is a representation of the C type GTK_RESPONSE_NO.
	ResponseType_No ResponseType = -9
	// ResponseType_Apply is a representation of the C type GTK_RESPONSE_APPLY.
	ResponseType_Apply ResponseType = -10
	// ResponseType_Help is a representation of the C type GTK_RESPONSE_HELP.
	ResponseType_Help ResponseType = -11
)

// RevealerTransitionType is a representation of the C type GtkRevealerTransitionType.
type RevealerTransitionType int

const (
	// RevealerTransitionType_None is a representation of the C type GTK_REVEALER_TRANSITION_TYPE_NONE.
	RevealerTransitionType_None RevealerTransitionType = 0
	// RevealerTransitionType_Crossfade is a representation of the C type GTK_REVEALER_TRANSITION_TYPE_CROSSFADE.
	RevealerTransitionType_Crossfade RevealerTransitionType = 1
	// RevealerTransitionType_SlideRight is a representation of the C type GTK_REVEALER_TRANSITION_TYPE_SLIDE_RIGHT.
	RevealerTransitionType_SlideRight RevealerTransitionType = 2
	// RevealerTransitionType_SlideLeft is a representation of the C type GTK_REVEALER_TRANSITION_TYPE_SLIDE_LEFT.
	RevealerTransitionType_SlideLeft RevealerTransitionType = 3
	// RevealerTransitionType_SlideUp is a representation of the C type GTK_REVEALER_TRANSITION_TYPE_SLIDE_UP.
	RevealerTransitionType_SlideUp RevealerTransitionType = 4
	// RevealerTransitionType_SlideDown is a representation of the C type GTK_REVEALER_TRANSITION_TYPE_SLIDE_DOWN.
	RevealerTransitionType_SlideDown RevealerTransitionType = 5
)

// ScrollStep is a representation of the C type GtkScrollStep.
type ScrollStep int

const (
	// ScrollStep_Steps is a representation of the C type GTK_SCROLL_STEPS.
	ScrollStep_Steps ScrollStep = 0
	// ScrollStep_Pages is a representation of the C type GTK_SCROLL_PAGES.
	ScrollStep_Pages ScrollStep = 1
	// ScrollStep_Ends is a representation of the C type GTK_SCROLL_ENDS.
	ScrollStep_Ends ScrollStep = 2
	// ScrollStep_HorizontalSteps is a representation of the C type GTK_SCROLL_HORIZONTAL_STEPS.
	ScrollStep_HorizontalSteps ScrollStep = 3
	// ScrollStep_HorizontalPages is a representation of the C type GTK_SCROLL_HORIZONTAL_PAGES.
	ScrollStep_HorizontalPages ScrollStep = 4
	// ScrollStep_HorizontalEnds is a representation of the C type GTK_SCROLL_HORIZONTAL_ENDS.
	ScrollStep_HorizontalEnds ScrollStep = 5
)

// ScrollType is a representation of the C type GtkScrollType.
type ScrollType int

const (
	// ScrollType_None is a representation of the C type GTK_SCROLL_NONE.
	ScrollType_None ScrollType = 0
	// ScrollType_Jump is a representation of the C type GTK_SCROLL_JUMP.
	ScrollType_Jump ScrollType = 1
	// ScrollType_StepBackward is a representation of the C type GTK_SCROLL_STEP_BACKWARD.
	ScrollType_StepBackward ScrollType = 2
	// ScrollType_StepForward is a representation of the C type GTK_SCROLL_STEP_FORWARD.
	ScrollType_StepForward ScrollType = 3
	// ScrollType_PageBackward is a representation of the C type GTK_SCROLL_PAGE_BACKWARD.
	ScrollType_PageBackward ScrollType = 4
	// ScrollType_PageForward is a representation of the C type GTK_SCROLL_PAGE_FORWARD.
	ScrollType_PageForward ScrollType = 5
	// ScrollType_StepUp is a representation of the C type GTK_SCROLL_STEP_UP.
	ScrollType_StepUp ScrollType = 6
	// ScrollType_StepDown is a representation of the C type GTK_SCROLL_STEP_DOWN.
	ScrollType_StepDown ScrollType = 7
	// ScrollType_PageUp is a representation of the C type GTK_SCROLL_PAGE_UP.
	ScrollType_PageUp ScrollType = 8
	// ScrollType_PageDown is a representation of the C type GTK_SCROLL_PAGE_DOWN.
	ScrollType_PageDown ScrollType = 9
	// ScrollType_StepLeft is a representation of the C type GTK_SCROLL_STEP_LEFT.
	ScrollType_StepLeft ScrollType = 10
	// ScrollType_StepRight is a representation of the C type GTK_SCROLL_STEP_RIGHT.
	ScrollType_StepRight ScrollType = 11
	// ScrollType_PageLeft is a representation of the C type GTK_SCROLL_PAGE_LEFT.
	ScrollType_PageLeft ScrollType = 12
	// ScrollType_PageRight is a representation of the C type GTK_SCROLL_PAGE_RIGHT.
	ScrollType_PageRight ScrollType = 13
	// ScrollType_Start is a representation of the C type GTK_SCROLL_START.
	ScrollType_Start ScrollType = 14
	// ScrollType_End is a representation of the C type GTK_SCROLL_END.
	ScrollType_End ScrollType = 15
)

// ScrollablePolicy is a representation of the C type GtkScrollablePolicy.
type ScrollablePolicy int

const (
	// ScrollablePolicy_Minimum is a representation of the C type GTK_SCROLL_MINIMUM.
	ScrollablePolicy_Minimum ScrollablePolicy = 0
	// ScrollablePolicy_Natural is a representation of the C type GTK_SCROLL_NATURAL.
	ScrollablePolicy_Natural ScrollablePolicy = 1
)

// SelectionMode is a representation of the C type GtkSelectionMode.
type SelectionMode int

const (
	// SelectionMode_None is a representation of the C type GTK_SELECTION_NONE.
	SelectionMode_None SelectionMode = 0
	// SelectionMode_Single is a representation of the C type GTK_SELECTION_SINGLE.
	SelectionMode_Single SelectionMode = 1
	// SelectionMode_Browse is a representation of the C type GTK_SELECTION_BROWSE.
	SelectionMode_Browse SelectionMode = 2
	// SelectionMode_Multiple is a representation of the C type GTK_SELECTION_MULTIPLE.
	SelectionMode_Multiple SelectionMode = 3
)

// SensitivityType is a representation of the C type GtkSensitivityType.
type SensitivityType int

const (
	// SensitivityType_Auto is a representation of the C type GTK_SENSITIVITY_AUTO.
	SensitivityType_Auto SensitivityType = 0
	// SensitivityType_On is a representation of the C type GTK_SENSITIVITY_ON.
	SensitivityType_On SensitivityType = 1
	// SensitivityType_Off is a representation of the C type GTK_SENSITIVITY_OFF.
	SensitivityType_Off SensitivityType = 2
)

// ShadowType is a representation of the C type GtkShadowType.
type ShadowType int

const (
	// ShadowType_None is a representation of the C type GTK_SHADOW_NONE.
	ShadowType_None ShadowType = 0
	// ShadowType_In is a representation of the C type GTK_SHADOW_IN.
	ShadowType_In ShadowType = 1
	// ShadowType_Out is a representation of the C type GTK_SHADOW_OUT.
	ShadowType_Out ShadowType = 2
	// ShadowType_EtchedIn is a representation of the C type GTK_SHADOW_ETCHED_IN.
	ShadowType_EtchedIn ShadowType = 3
	// ShadowType_EtchedOut is a representation of the C type GTK_SHADOW_ETCHED_OUT.
	ShadowType_EtchedOut ShadowType = 4
)

// ShortcutType is a representation of the C type GtkShortcutType.
//
// since 3.20
type ShortcutType int

const (
	// ShortcutType_Accelerator is a representation of the C type GTK_SHORTCUT_ACCELERATOR.
	ShortcutType_Accelerator ShortcutType = 0
	// ShortcutType_GesturePinch is a representation of the C type GTK_SHORTCUT_GESTURE_PINCH.
	ShortcutType_GesturePinch ShortcutType = 1
	// ShortcutType_GestureStretch is a representation of the C type GTK_SHORTCUT_GESTURE_STRETCH.
	ShortcutType_GestureStretch ShortcutType = 2
	// ShortcutType_GestureRotateClockwise is a representation of the C type GTK_SHORTCUT_GESTURE_ROTATE_CLOCKWISE.
	ShortcutType_GestureRotateClockwise ShortcutType = 3
	// ShortcutType_GestureRotateCounterclockwise is a representation of the C type GTK_SHORTCUT_GESTURE_ROTATE_COUNTERCLOCKWISE.
	ShortcutType_GestureRotateCounterclockwise ShortcutType = 4
	// ShortcutType_GestureTwoFingerSwipeLeft is a representation of the C type GTK_SHORTCUT_GESTURE_TWO_FINGER_SWIPE_LEFT.
	ShortcutType_GestureTwoFingerSwipeLeft ShortcutType = 5
	// ShortcutType_GestureTwoFingerSwipeRight is a representation of the C type GTK_SHORTCUT_GESTURE_TWO_FINGER_SWIPE_RIGHT.
	ShortcutType_GestureTwoFingerSwipeRight ShortcutType = 6
	// ShortcutType_Gesture is a representation of the C type GTK_SHORTCUT_GESTURE.
	ShortcutType_Gesture ShortcutType = 7
)

// SizeGroupMode is a representation of the C type GtkSizeGroupMode.
type SizeGroupMode int

const (
	// SizeGroupMode_None is a representation of the C type GTK_SIZE_GROUP_NONE.
	SizeGroupMode_None SizeGroupMode = 0
	// SizeGroupMode_Horizontal is a representation of the C type GTK_SIZE_GROUP_HORIZONTAL.
	SizeGroupMode_Horizontal SizeGroupMode = 1
	// SizeGroupMode_Vertical is a representation of the C type GTK_SIZE_GROUP_VERTICAL.
	SizeGroupMode_Vertical SizeGroupMode = 2
	// SizeGroupMode_Both is a representation of the C type GTK_SIZE_GROUP_BOTH.
	SizeGroupMode_Both SizeGroupMode = 3
)

// SizeRequestMode is a representation of the C type GtkSizeRequestMode.
type SizeRequestMode int

const (
	// SizeRequestMode_HeightForWidth is a representation of the C type GTK_SIZE_REQUEST_HEIGHT_FOR_WIDTH.
	SizeRequestMode_HeightForWidth SizeRequestMode = 0
	// SizeRequestMode_WidthForHeight is a representation of the C type GTK_SIZE_REQUEST_WIDTH_FOR_HEIGHT.
	SizeRequestMode_WidthForHeight SizeRequestMode = 1
	// SizeRequestMode_ConstantSize is a representation of the C type GTK_SIZE_REQUEST_CONSTANT_SIZE.
	SizeRequestMode_ConstantSize SizeRequestMode = 2
)

// SortType is a representation of the C type GtkSortType.
type SortType int

const (
	// SortType_Ascending is a representation of the C type GTK_SORT_ASCENDING.
	SortType_Ascending SortType = 0
	// SortType_Descending is a representation of the C type GTK_SORT_DESCENDING.
	SortType_Descending SortType = 1
)

// SpinButtonUpdatePolicy is a representation of the C type GtkSpinButtonUpdatePolicy.
type SpinButtonUpdatePolicy int

const (
	// SpinButtonUpdatePolicy_Always is a representation of the C type GTK_UPDATE_ALWAYS.
	SpinButtonUpdatePolicy_Always SpinButtonUpdatePolicy = 0
	// SpinButtonUpdatePolicy_IfValid is a representation of the C type GTK_UPDATE_IF_VALID.
	SpinButtonUpdatePolicy_IfValid SpinButtonUpdatePolicy = 1
)

// SpinType is a representation of the C type GtkSpinType.
type SpinType int

const (
	// SpinType_StepForward is a representation of the C type GTK_SPIN_STEP_FORWARD.
	SpinType_StepForward SpinType = 0
	// SpinType_StepBackward is a representation of the C type GTK_SPIN_STEP_BACKWARD.
	SpinType_StepBackward SpinType = 1
	// SpinType_PageForward is a representation of the C type GTK_SPIN_PAGE_FORWARD.
	SpinType_PageForward SpinType = 2
	// SpinType_PageBackward is a representation of the C type GTK_SPIN_PAGE_BACKWARD.
	SpinType_PageBackward SpinType = 3
	// SpinType_Home is a representation of the C type GTK_SPIN_HOME.
	SpinType_Home SpinType = 4
	// SpinType_End is a representation of the C type GTK_SPIN_END.
	SpinType_End SpinType = 5
	// SpinType_UserDefined is a representation of the C type GTK_SPIN_USER_DEFINED.
	SpinType_UserDefined SpinType = 6
)

// StackTransitionType is a representation of the C type GtkStackTransitionType.
type StackTransitionType int

const (
	// StackTransitionType_None is a representation of the C type GTK_STACK_TRANSITION_TYPE_NONE.
	StackTransitionType_None StackTransitionType = 0
	// StackTransitionType_Crossfade is a representation of the C type GTK_STACK_TRANSITION_TYPE_CROSSFADE.
	StackTransitionType_Crossfade StackTransitionType = 1
	// StackTransitionType_SlideRight is a representation of the C type GTK_STACK_TRANSITION_TYPE_SLIDE_RIGHT.
	StackTransitionType_SlideRight StackTransitionType = 2
	// StackTransitionType_SlideLeft is a representation of the C type GTK_STACK_TRANSITION_TYPE_SLIDE_LEFT.
	StackTransitionType_SlideLeft StackTransitionType = 3
	// StackTransitionType_SlideUp is a representation of the C type GTK_STACK_TRANSITION_TYPE_SLIDE_UP.
	StackTransitionType_SlideUp StackTransitionType = 4
	// StackTransitionType_SlideDown is a representation of the C type GTK_STACK_TRANSITION_TYPE_SLIDE_DOWN.
	StackTransitionType_SlideDown StackTransitionType = 5
	// StackTransitionType_SlideLeftRight is a representation of the C type GTK_STACK_TRANSITION_TYPE_SLIDE_LEFT_RIGHT.
	StackTransitionType_SlideLeftRight StackTransitionType = 6
	// StackTransitionType_SlideUpDown is a representation of the C type GTK_STACK_TRANSITION_TYPE_SLIDE_UP_DOWN.
	StackTransitionType_SlideUpDown StackTransitionType = 7
	// StackTransitionType_OverUp is a representation of the C type GTK_STACK_TRANSITION_TYPE_OVER_UP.
	StackTransitionType_OverUp StackTransitionType = 8
	// StackTransitionType_OverDown is a representation of the C type GTK_STACK_TRANSITION_TYPE_OVER_DOWN.
	StackTransitionType_OverDown StackTransitionType = 9
	// StackTransitionType_OverLeft is a representation of the C type GTK_STACK_TRANSITION_TYPE_OVER_LEFT.
	StackTransitionType_OverLeft StackTransitionType = 10
	// StackTransitionType_OverRight is a representation of the C type GTK_STACK_TRANSITION_TYPE_OVER_RIGHT.
	StackTransitionType_OverRight StackTransitionType = 11
	// StackTransitionType_UnderUp is a representation of the C type GTK_STACK_TRANSITION_TYPE_UNDER_UP.
	StackTransitionType_UnderUp StackTransitionType = 12
	// StackTransitionType_UnderDown is a representation of the C type GTK_STACK_TRANSITION_TYPE_UNDER_DOWN.
	StackTransitionType_UnderDown StackTransitionType = 13
	// StackTransitionType_UnderLeft is a representation of the C type GTK_STACK_TRANSITION_TYPE_UNDER_LEFT.
	StackTransitionType_UnderLeft StackTransitionType = 14
	// StackTransitionType_UnderRight is a representation of the C type GTK_STACK_TRANSITION_TYPE_UNDER_RIGHT.
	StackTransitionType_UnderRight StackTransitionType = 15
	// StackTransitionType_OverUpDown is a representation of the C type GTK_STACK_TRANSITION_TYPE_OVER_UP_DOWN.
	StackTransitionType_OverUpDown StackTransitionType = 16
	// StackTransitionType_OverDownUp is a representation of the C type GTK_STACK_TRANSITION_TYPE_OVER_DOWN_UP.
	StackTransitionType_OverDownUp StackTransitionType = 17
	// StackTransitionType_OverLeftRight is a representation of the C type GTK_STACK_TRANSITION_TYPE_OVER_LEFT_RIGHT.
	StackTransitionType_OverLeftRight StackTransitionType = 18
	// StackTransitionType_OverRightLeft is a representation of the C type GTK_STACK_TRANSITION_TYPE_OVER_RIGHT_LEFT.
	StackTransitionType_OverRightLeft StackTransitionType = 19
)

// StateType is a representation of the C type GtkStateType.
type StateType int

const (
	// StateType_Normal is a representation of the C type GTK_STATE_NORMAL.
	StateType_Normal StateType = 0
	// StateType_Active is a representation of the C type GTK_STATE_ACTIVE.
	StateType_Active StateType = 1
	// StateType_Prelight is a representation of the C type GTK_STATE_PRELIGHT.
	StateType_Prelight StateType = 2
	// StateType_Selected is a representation of the C type GTK_STATE_SELECTED.
	StateType_Selected StateType = 3
	// StateType_Insensitive is a representation of the C type GTK_STATE_INSENSITIVE.
	StateType_Insensitive StateType = 4
	// StateType_Inconsistent is a representation of the C type GTK_STATE_INCONSISTENT.
	StateType_Inconsistent StateType = 5
	// StateType_Focused is a representation of the C type GTK_STATE_FOCUSED.
	StateType_Focused StateType = 6
)

// TextBufferTargetInfo is a representation of the C type GtkTextBufferTargetInfo.
type TextBufferTargetInfo int

const (
	// TextBufferTargetInfo_BufferContents is a representation of the C type GTK_TEXT_BUFFER_TARGET_INFO_BUFFER_CONTENTS.
	TextBufferTargetInfo_BufferContents TextBufferTargetInfo = -1
	// TextBufferTargetInfo_RichText is a representation of the C type GTK_TEXT_BUFFER_TARGET_INFO_RICH_TEXT.
	TextBufferTargetInfo_RichText TextBufferTargetInfo = -2
	// TextBufferTargetInfo_Text is a representation of the C type GTK_TEXT_BUFFER_TARGET_INFO_TEXT.
	TextBufferTargetInfo_Text TextBufferTargetInfo = -3
)

// TextDirection is a representation of the C type GtkTextDirection.
type TextDirection int

const (
	// TextDirection_None is a representation of the C type GTK_TEXT_DIR_NONE.
	TextDirection_None TextDirection = 0
	// TextDirection_Ltr is a representation of the C type GTK_TEXT_DIR_LTR.
	TextDirection_Ltr TextDirection = 1
	// TextDirection_Rtl is a representation of the C type GTK_TEXT_DIR_RTL.
	TextDirection_Rtl TextDirection = 2
)

// TextExtendSelection is a representation of the C type GtkTextExtendSelection.
//
// since 3.16
type TextExtendSelection int

const (
	// TextExtendSelection_Word is a representation of the C type GTK_TEXT_EXTEND_SELECTION_WORD.
	TextExtendSelection_Word TextExtendSelection = 0
	// TextExtendSelection_Line is a representation of the C type GTK_TEXT_EXTEND_SELECTION_LINE.
	TextExtendSelection_Line TextExtendSelection = 1
)

// TextViewLayer is a representation of the C type GtkTextViewLayer.
type TextViewLayer int

const (
	// TextViewLayer_Below is a representation of the C type GTK_TEXT_VIEW_LAYER_BELOW.
	TextViewLayer_Below TextViewLayer = 0
	// TextViewLayer_Above is a representation of the C type GTK_TEXT_VIEW_LAYER_ABOVE.
	TextViewLayer_Above TextViewLayer = 1
	// TextViewLayer_BelowText is a representation of the C type GTK_TEXT_VIEW_LAYER_BELOW_TEXT.
	TextViewLayer_BelowText TextViewLayer = 2
	// TextViewLayer_AboveText is a representation of the C type GTK_TEXT_VIEW_LAYER_ABOVE_TEXT.
	TextViewLayer_AboveText TextViewLayer = 3
)

// TextWindowType is a representation of the C type GtkTextWindowType.
type TextWindowType int

const (
	// TextWindowType_Private is a representation of the C type GTK_TEXT_WINDOW_PRIVATE.
	TextWindowType_Private TextWindowType = 0
	// TextWindowType_Widget is a representation of the C type GTK_TEXT_WINDOW_WIDGET.
	TextWindowType_Widget TextWindowType = 1
	// TextWindowType_Text is a representation of the C type GTK_TEXT_WINDOW_TEXT.
	TextWindowType_Text TextWindowType = 2
	// TextWindowType_Left is a representation of the C type GTK_TEXT_WINDOW_LEFT.
	TextWindowType_Left TextWindowType = 3
	// TextWindowType_Right is a representation of the C type GTK_TEXT_WINDOW_RIGHT.
	TextWindowType_Right TextWindowType = 4
	// TextWindowType_Top is a representation of the C type GTK_TEXT_WINDOW_TOP.
	TextWindowType_Top TextWindowType = 5
	// TextWindowType_Bottom is a representation of the C type GTK_TEXT_WINDOW_BOTTOM.
	TextWindowType_Bottom TextWindowType = 6
)

// ToolbarSpaceStyle is a representation of the C type GtkToolbarSpaceStyle.
type ToolbarSpaceStyle int

const (
	// ToolbarSpaceStyle_Empty is a representation of the C type GTK_TOOLBAR_SPACE_EMPTY.
	ToolbarSpaceStyle_Empty ToolbarSpaceStyle = 0
	// ToolbarSpaceStyle_Line is a representation of the C type GTK_TOOLBAR_SPACE_LINE.
	ToolbarSpaceStyle_Line ToolbarSpaceStyle = 1
)

// ToolbarStyle is a representation of the C type GtkToolbarStyle.
type ToolbarStyle int

const (
	// ToolbarStyle_Icons is a representation of the C type GTK_TOOLBAR_ICONS.
	ToolbarStyle_Icons ToolbarStyle = 0
	// ToolbarStyle_Text is a representation of the C type GTK_TOOLBAR_TEXT.
	ToolbarStyle_Text ToolbarStyle = 1
	// ToolbarStyle_Both is a representation of the C type GTK_TOOLBAR_BOTH.
	ToolbarStyle_Both ToolbarStyle = 2
	// ToolbarStyle_BothHoriz is a representation of the C type GTK_TOOLBAR_BOTH_HORIZ.
	ToolbarStyle_BothHoriz ToolbarStyle = 3
)

// TreeViewColumnSizing is a representation of the C type GtkTreeViewColumnSizing.
type TreeViewColumnSizing int

const (
	// TreeViewColumnSizing_GrowOnly is a representation of the C type GTK_TREE_VIEW_COLUMN_GROW_ONLY.
	TreeViewColumnSizing_GrowOnly TreeViewColumnSizing = 0
	// TreeViewColumnSizing_Autosize is a representation of the C type GTK_TREE_VIEW_COLUMN_AUTOSIZE.
	TreeViewColumnSizing_Autosize TreeViewColumnSizing = 1
	// TreeViewColumnSizing_Fixed is a representation of the C type GTK_TREE_VIEW_COLUMN_FIXED.
	TreeViewColumnSizing_Fixed TreeViewColumnSizing = 2
)

// TreeViewDropPosition is a representation of the C type GtkTreeViewDropPosition.
type TreeViewDropPosition int

const (
	// TreeViewDropPosition_Before is a representation of the C type GTK_TREE_VIEW_DROP_BEFORE.
	TreeViewDropPosition_Before TreeViewDropPosition = 0
	// TreeViewDropPosition_After is a representation of the C type GTK_TREE_VIEW_DROP_AFTER.
	TreeViewDropPosition_After TreeViewDropPosition = 1
	// TreeViewDropPosition_IntoOrBefore is a representation of the C type GTK_TREE_VIEW_DROP_INTO_OR_BEFORE.
	TreeViewDropPosition_IntoOrBefore TreeViewDropPosition = 2
	// TreeViewDropPosition_IntoOrAfter is a representation of the C type GTK_TREE_VIEW_DROP_INTO_OR_AFTER.
	TreeViewDropPosition_IntoOrAfter TreeViewDropPosition = 3
)

// TreeViewGridLines is a representation of the C type GtkTreeViewGridLines.
type TreeViewGridLines int

const (
	// TreeViewGridLines_None is a representation of the C type GTK_TREE_VIEW_GRID_LINES_NONE.
	TreeViewGridLines_None TreeViewGridLines = 0
	// TreeViewGridLines_Horizontal is a representation of the C type GTK_TREE_VIEW_GRID_LINES_HORIZONTAL.
	TreeViewGridLines_Horizontal TreeViewGridLines = 1
	// TreeViewGridLines_Vertical is a representation of the C type GTK_TREE_VIEW_GRID_LINES_VERTICAL.
	TreeViewGridLines_Vertical TreeViewGridLines = 2
	// TreeViewGridLines_Both is a representation of the C type GTK_TREE_VIEW_GRID_LINES_BOTH.
	TreeViewGridLines_Both TreeViewGridLines = 3
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

// WidgetHelpType is a representation of the C type GtkWidgetHelpType.
type WidgetHelpType int

const (
	// WidgetHelpType_Tooltip is a representation of the C type GTK_WIDGET_HELP_TOOLTIP.
	WidgetHelpType_Tooltip WidgetHelpType = 0
	// WidgetHelpType_WhatsThis is a representation of the C type GTK_WIDGET_HELP_WHATS_THIS.
	WidgetHelpType_WhatsThis WidgetHelpType = 1
)

// WindowPosition is a representation of the C type GtkWindowPosition.
type WindowPosition int

const (
	// WindowPosition_None is a representation of the C type GTK_WIN_POS_NONE.
	WindowPosition_None WindowPosition = 0
	// WindowPosition_Center is a representation of the C type GTK_WIN_POS_CENTER.
	WindowPosition_Center WindowPosition = 1
	// WindowPosition_Mouse is a representation of the C type GTK_WIN_POS_MOUSE.
	WindowPosition_Mouse WindowPosition = 2
	// WindowPosition_CenterAlways is a representation of the C type GTK_WIN_POS_CENTER_ALWAYS.
	WindowPosition_CenterAlways WindowPosition = 3
	// WindowPosition_CenterOnParent is a representation of the C type GTK_WIN_POS_CENTER_ON_PARENT.
	WindowPosition_CenterOnParent WindowPosition = 4
)

// WindowType is a representation of the C type GtkWindowType.
type WindowType int

const (
	// WindowType_Toplevel is a representation of the C type GTK_WINDOW_TOPLEVEL.
	WindowType_Toplevel WindowType = 0
	// WindowType_Popup is a representation of the C type GTK_WINDOW_POPUP.
	WindowType_Popup WindowType = 1
)

// WrapMode is a representation of the C type GtkWrapMode.
type WrapMode int

const (
	// WrapMode_None is a representation of the C type GTK_WRAP_NONE.
	WrapMode_None WrapMode = 0
	// WrapMode_Char is a representation of the C type GTK_WRAP_CHAR.
	WrapMode_Char WrapMode = 1
	// WrapMode_Word is a representation of the C type GTK_WRAP_WORD.
	WrapMode_Word WrapMode = 2
	// WrapMode_WordChar is a representation of the C type GTK_WRAP_WORD_CHAR.
	WrapMode_WordChar WrapMode = 3
)
