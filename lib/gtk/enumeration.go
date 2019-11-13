// Code generated - DO NOT EDIT.

package gtk

// Align is a representation of the C type Align.
type Align int

const (
	// fill
	GtkAlignFill Align = 0
	// start
	GtkAlignStart Align = 1
	// end
	GtkAlignEnd Align = 2
	// center
	GtkAlignCenter Align = 3
	// baseline
	GtkAlignBaseline Align = 4
)

// Arrowplacement is a representation of the C type ArrowPlacement.
type Arrowplacement int

const (
	// both
	GtkArrowsBoth Arrowplacement = 0
	// start
	GtkArrowsStart Arrowplacement = 1
	// end
	GtkArrowsEnd Arrowplacement = 2
)

// Arrowtype is a representation of the C type ArrowType.
type Arrowtype int

const (
	// up
	GtkArrowUp Arrowtype = 0
	// down
	GtkArrowDown Arrowtype = 1
	// left
	GtkArrowLeft Arrowtype = 2
	// right
	GtkArrowRight Arrowtype = 3
	// none
	GtkArrowNone Arrowtype = 4
)

// Assistantpagetype is a representation of the C type AssistantPageType.
type Assistantpagetype int

const (
	// content
	GtkAssistantPageContent Assistantpagetype = 0
	// intro
	GtkAssistantPageIntro Assistantpagetype = 1
	// confirm
	GtkAssistantPageConfirm Assistantpagetype = 2
	// summary
	GtkAssistantPageSummary Assistantpagetype = 3
	// progress
	GtkAssistantPageProgress Assistantpagetype = 4
	// custom
	GtkAssistantPageCustom Assistantpagetype = 5
)

// Baselineposition is a representation of the C type BaselinePosition.
//
// since 3.10
type Baselineposition int

const (
	// top
	GtkBaselinePositionTop Baselineposition = 0
	// center
	GtkBaselinePositionCenter Baselineposition = 1
	// bottom
	GtkBaselinePositionBottom Baselineposition = 2
)

// Borderstyle is a representation of the C type BorderStyle.
type Borderstyle int

const (
	// none
	GtkBorderStyleNone Borderstyle = 0
	// solid
	GtkBorderStyleSolid Borderstyle = 1
	// inset
	GtkBorderStyleInset Borderstyle = 2
	// outset
	GtkBorderStyleOutset Borderstyle = 3
	// hidden
	GtkBorderStyleHidden Borderstyle = 4
	// dotted
	GtkBorderStyleDotted Borderstyle = 5
	// dashed
	GtkBorderStyleDashed Borderstyle = 6
	// double
	GtkBorderStyleDouble Borderstyle = 7
	// groove
	GtkBorderStyleGroove Borderstyle = 8
	// ridge
	GtkBorderStyleRidge Borderstyle = 9
)

// Buildererror is a representation of the C type BuilderError.
type Buildererror int

const (
	// invalid_type_function
	GtkBuilderErrorInvalidTypeFunction Buildererror = 0
	// unhandled_tag
	GtkBuilderErrorUnhandledTag Buildererror = 1
	// missing_attribute
	GtkBuilderErrorMissingAttribute Buildererror = 2
	// invalid_attribute
	GtkBuilderErrorInvalidAttribute Buildererror = 3
	// invalid_tag
	GtkBuilderErrorInvalidTag Buildererror = 4
	// missing_property_value
	GtkBuilderErrorMissingPropertyValue Buildererror = 5
	// invalid_value
	GtkBuilderErrorInvalidValue Buildererror = 6
	// version_mismatch
	GtkBuilderErrorVersionMismatch Buildererror = 7
	// duplicate_id
	GtkBuilderErrorDuplicateId Buildererror = 8
	// object_type_refused
	GtkBuilderErrorObjectTypeRefused Buildererror = 9
	// template_mismatch
	GtkBuilderErrorTemplateMismatch Buildererror = 10
	// invalid_property
	GtkBuilderErrorInvalidProperty Buildererror = 11
	// invalid_signal
	GtkBuilderErrorInvalidSignal Buildererror = 12
	// invalid_id
	GtkBuilderErrorInvalidId Buildererror = 13
)

// Buttonboxstyle is a representation of the C type ButtonBoxStyle.
type Buttonboxstyle int

const (
	// spread
	GtkButtonboxSpread Buttonboxstyle = 1
	// edge
	GtkButtonboxEdge Buttonboxstyle = 2
	// start
	GtkButtonboxStart Buttonboxstyle = 3
	// end
	GtkButtonboxEnd Buttonboxstyle = 4
	// center
	GtkButtonboxCenter Buttonboxstyle = 5
	// expand
	GtkButtonboxExpand Buttonboxstyle = 6
)

// Buttonrole is a representation of the C type ButtonRole.
type Buttonrole int

const (
	// normal
	GtkButtonRoleNormal Buttonrole = 0
	// check
	GtkButtonRoleCheck Buttonrole = 1
	// radio
	GtkButtonRoleRadio Buttonrole = 2
)

// Buttonstype is a representation of the C type ButtonsType.
type Buttonstype int

const (
	// none
	GtkButtonsNone Buttonstype = 0
	// ok
	GtkButtonsOk Buttonstype = 1
	// close
	GtkButtonsClose Buttonstype = 2
	// cancel
	GtkButtonsCancel Buttonstype = 3
	// yes_no
	GtkButtonsYesNo Buttonstype = 4
	// ok_cancel
	GtkButtonsOkCancel Buttonstype = 5
)

// Cellrendereraccelmode is a representation of the C type CellRendererAccelMode.
type Cellrendereraccelmode int

const (
	// gtk
	GtkCellRendererAccelModeGtk Cellrendereraccelmode = 0
	// other
	GtkCellRendererAccelModeOther Cellrendereraccelmode = 1
	// modifier_tap
	GtkCellRendererAccelModeModifierTap Cellrendereraccelmode = 2
)

// Cellrenderermode is a representation of the C type CellRendererMode.
type Cellrenderermode int

const (
	// inert
	GtkCellRendererModeInert Cellrenderermode = 0
	// activatable
	GtkCellRendererModeActivatable Cellrenderermode = 1
	// editable
	GtkCellRendererModeEditable Cellrenderermode = 2
)

// Cornertype is a representation of the C type CornerType.
type Cornertype int

const (
	// top_left
	GtkCornerTopLeft Cornertype = 0
	// bottom_left
	GtkCornerBottomLeft Cornertype = 1
	// top_right
	GtkCornerTopRight Cornertype = 2
	// bottom_right
	GtkCornerBottomRight Cornertype = 3
)

// Cssprovidererror is a representation of the C type CssProviderError.
type Cssprovidererror int

const (
	// failed
	GtkCssProviderErrorFailed Cssprovidererror = 0
	// syntax
	GtkCssProviderErrorSyntax Cssprovidererror = 1
	// import
	GtkCssProviderErrorImport Cssprovidererror = 2
	// name
	GtkCssProviderErrorName Cssprovidererror = 3
	// deprecated
	GtkCssProviderErrorDeprecated Cssprovidererror = 4
	// unknown_value
	GtkCssProviderErrorUnknownValue Cssprovidererror = 5
)

// Csssectiontype is a representation of the C type CssSectionType.
//
// since 3.2
type Csssectiontype int

const (
	// document
	GtkCssSectionDocument Csssectiontype = 0
	// import
	GtkCssSectionImport Csssectiontype = 1
	// color_definition
	GtkCssSectionColorDefinition Csssectiontype = 2
	// binding_set
	GtkCssSectionBindingSet Csssectiontype = 3
	// ruleset
	GtkCssSectionRuleset Csssectiontype = 4
	// selector
	GtkCssSectionSelector Csssectiontype = 5
	// declaration
	GtkCssSectionDeclaration Csssectiontype = 6
	// value
	GtkCssSectionValue Csssectiontype = 7
	// keyframes
	GtkCssSectionKeyframes Csssectiontype = 8
)

// Deletetype is a representation of the C type DeleteType.
type Deletetype int

const (
	// chars
	GtkDeleteChars Deletetype = 0
	// word_ends
	GtkDeleteWordEnds Deletetype = 1
	// words
	GtkDeleteWords Deletetype = 2
	// display_lines
	GtkDeleteDisplayLines Deletetype = 3
	// display_line_ends
	GtkDeleteDisplayLineEnds Deletetype = 4
	// paragraph_ends
	GtkDeleteParagraphEnds Deletetype = 5
	// paragraphs
	GtkDeleteParagraphs Deletetype = 6
	// whitespace
	GtkDeleteWhitespace Deletetype = 7
)

// Directiontype is a representation of the C type DirectionType.
type Directiontype int

const (
	// tab_forward
	GtkDirTabForward Directiontype = 0
	// tab_backward
	GtkDirTabBackward Directiontype = 1
	// up
	GtkDirUp Directiontype = 2
	// down
	GtkDirDown Directiontype = 3
	// left
	GtkDirLeft Directiontype = 4
	// right
	GtkDirRight Directiontype = 5
)

// Dragresult is a representation of the C type DragResult.
type Dragresult int

const (
	// success
	GtkDragResultSuccess Dragresult = 0
	// no_target
	GtkDragResultNoTarget Dragresult = 1
	// user_cancelled
	GtkDragResultUserCancelled Dragresult = 2
	// timeout_expired
	GtkDragResultTimeoutExpired Dragresult = 3
	// grab_broken
	GtkDragResultGrabBroken Dragresult = 4
	// error
	GtkDragResultError Dragresult = 5
)

// Entryiconposition is a representation of the C type EntryIconPosition.
//
// since 2.16
type Entryiconposition int

const (
	// primary
	GtkEntryIconPrimary Entryiconposition = 0
	// secondary
	GtkEntryIconSecondary Entryiconposition = 1
)

// Eventsequencestate is a representation of the C type EventSequenceState.
//
// since 3.14
type Eventsequencestate int

const (
	// none
	GtkEventSequenceNone Eventsequencestate = 0
	// claimed
	GtkEventSequenceClaimed Eventsequencestate = 1
	// denied
	GtkEventSequenceDenied Eventsequencestate = 2
)

// Expanderstyle is a representation of the C type ExpanderStyle.
type Expanderstyle int

const (
	// collapsed
	GtkExpanderCollapsed Expanderstyle = 0
	// semi_collapsed
	GtkExpanderSemiCollapsed Expanderstyle = 1
	// semi_expanded
	GtkExpanderSemiExpanded Expanderstyle = 2
	// expanded
	GtkExpanderExpanded Expanderstyle = 3
)

// Filechooseraction is a representation of the C type FileChooserAction.
type Filechooseraction int

const (
	// open
	GtkFileChooserActionOpen Filechooseraction = 0
	// save
	GtkFileChooserActionSave Filechooseraction = 1
	// select_folder
	GtkFileChooserActionSelectFolder Filechooseraction = 2
	// create_folder
	GtkFileChooserActionCreateFolder Filechooseraction = 3
)

// Filechooserconfirmation is a representation of the C type FileChooserConfirmation.
//
// since 2.8
type Filechooserconfirmation int

const (
	// confirm
	GtkFileChooserConfirmationConfirm Filechooserconfirmation = 0
	// accept_filename
	GtkFileChooserConfirmationAcceptFilename Filechooserconfirmation = 1
	// select_again
	GtkFileChooserConfirmationSelectAgain Filechooserconfirmation = 2
)

// Filechoosererror is a representation of the C type FileChooserError.
type Filechoosererror int

const (
	// nonexistent
	GtkFileChooserErrorNonexistent Filechoosererror = 0
	// bad_filename
	GtkFileChooserErrorBadFilename Filechoosererror = 1
	// already_exists
	GtkFileChooserErrorAlreadyExists Filechoosererror = 2
	// incomplete_hostname
	GtkFileChooserErrorIncompleteHostname Filechoosererror = 3
)

// Impreeditstyle is a representation of the C type IMPreeditStyle.
type Impreeditstyle int

const (
	// nothing
	GtkImPreeditNothing Impreeditstyle = 0
	// callback
	GtkImPreeditCallback Impreeditstyle = 1
	// none
	GtkImPreeditNone Impreeditstyle = 2
)

// Imstatusstyle is a representation of the C type IMStatusStyle.
type Imstatusstyle int

const (
	// nothing
	GtkImStatusNothing Imstatusstyle = 0
	// callback
	GtkImStatusCallback Imstatusstyle = 1
	// none
	GtkImStatusNone Imstatusstyle = 2
)

// Iconsize is a representation of the C type IconSize.
type Iconsize int

const (
	// invalid
	GtkIconSizeInvalid Iconsize = 0
	// menu
	GtkIconSizeMenu Iconsize = 1
	// small_toolbar
	GtkIconSizeSmallToolbar Iconsize = 2
	// large_toolbar
	GtkIconSizeLargeToolbar Iconsize = 3
	// button
	GtkIconSizeButton Iconsize = 4
	// dnd
	GtkIconSizeDnd Iconsize = 5
	// dialog
	GtkIconSizeDialog Iconsize = 6
)

// Iconthemeerror is a representation of the C type IconThemeError.
type Iconthemeerror int

const (
	// not_found
	GtkIconThemeNotFound Iconthemeerror = 0
	// failed
	GtkIconThemeFailed Iconthemeerror = 1
)

// Iconviewdropposition is a representation of the C type IconViewDropPosition.
type Iconviewdropposition int

const (
	// no_drop
	GtkIconViewNoDrop Iconviewdropposition = 0
	// drop_into
	GtkIconViewDropInto Iconviewdropposition = 1
	// drop_left
	GtkIconViewDropLeft Iconviewdropposition = 2
	// drop_right
	GtkIconViewDropRight Iconviewdropposition = 3
	// drop_above
	GtkIconViewDropAbove Iconviewdropposition = 4
	// drop_below
	GtkIconViewDropBelow Iconviewdropposition = 5
)

// Imagetype is a representation of the C type ImageType.
type Imagetype int

const (
	// empty
	GtkImageEmpty Imagetype = 0
	// pixbuf
	GtkImagePixbuf Imagetype = 1
	// stock
	GtkImageStock Imagetype = 2
	// icon_set
	GtkImageIconSet Imagetype = 3
	// animation
	GtkImageAnimation Imagetype = 4
	// icon_name
	GtkImageIconName Imagetype = 5
	// gicon
	GtkImageGicon Imagetype = 6
	// surface
	GtkImageSurface Imagetype = 7
)

// Inputpurpose is a representation of the C type InputPurpose.
//
// since 3.6
type Inputpurpose int

const (
	// free_form
	GtkInputPurposeFreeForm Inputpurpose = 0
	// alpha
	GtkInputPurposeAlpha Inputpurpose = 1
	// digits
	GtkInputPurposeDigits Inputpurpose = 2
	// number
	GtkInputPurposeNumber Inputpurpose = 3
	// phone
	GtkInputPurposePhone Inputpurpose = 4
	// url
	GtkInputPurposeUrl Inputpurpose = 5
	// email
	GtkInputPurposeEmail Inputpurpose = 6
	// name
	GtkInputPurposeName Inputpurpose = 7
	// password
	GtkInputPurposePassword Inputpurpose = 8
	// pin
	GtkInputPurposePin Inputpurpose = 9
)

// Justification is a representation of the C type Justification.
type Justification int

const (
	// left
	GtkJustifyLeft Justification = 0
	// right
	GtkJustifyRight Justification = 1
	// center
	GtkJustifyCenter Justification = 2
	// fill
	GtkJustifyFill Justification = 3
)

// Levelbarmode is a representation of the C type LevelBarMode.
//
// since 3.6
type Levelbarmode int

const (
	// continuous
	GtkLevelBarModeContinuous Levelbarmode = 0
	// discrete
	GtkLevelBarModeDiscrete Levelbarmode = 1
)

// License is a representation of the C type License.
//
// since 3.0
type License int

const (
	// unknown
	GtkLicenseUnknown License = 0
	// custom
	GtkLicenseCustom License = 1
	// gpl_2_0
	GtkLicenseGpl20 License = 2
	// gpl_3_0
	GtkLicenseGpl30 License = 3
	// lgpl_2_1
	GtkLicenseLgpl21 License = 4
	// lgpl_3_0
	GtkLicenseLgpl30 License = 5
	// bsd
	GtkLicenseBsd License = 6
	// mit_x11
	GtkLicenseMitX11 License = 7
	// artistic
	GtkLicenseArtistic License = 8
	// gpl_2_0_only
	GtkLicenseGpl20Only License = 9
	// gpl_3_0_only
	GtkLicenseGpl30Only License = 10
	// lgpl_2_1_only
	GtkLicenseLgpl21Only License = 11
	// lgpl_3_0_only
	GtkLicenseLgpl30Only License = 12
	// agpl_3_0
	GtkLicenseAgpl30 License = 13
	// agpl_3_0_only
	GtkLicenseAgpl30Only License = 14
)

// Menudirectiontype is a representation of the C type MenuDirectionType.
type Menudirectiontype int

const (
	// parent
	GtkMenuDirParent Menudirectiontype = 0
	// child
	GtkMenuDirChild Menudirectiontype = 1
	// next
	GtkMenuDirNext Menudirectiontype = 2
	// prev
	GtkMenuDirPrev Menudirectiontype = 3
)

// Messagetype is a representation of the C type MessageType.
type Messagetype int

const (
	// info
	GtkMessageInfo Messagetype = 0
	// warning
	GtkMessageWarning Messagetype = 1
	// question
	GtkMessageQuestion Messagetype = 2
	// error
	GtkMessageError Messagetype = 3
	// other
	GtkMessageOther Messagetype = 4
)

// Movementstep is a representation of the C type MovementStep.
type Movementstep int

const (
	// logical_positions
	GtkMovementLogicalPositions Movementstep = 0
	// visual_positions
	GtkMovementVisualPositions Movementstep = 1
	// words
	GtkMovementWords Movementstep = 2
	// display_lines
	GtkMovementDisplayLines Movementstep = 3
	// display_line_ends
	GtkMovementDisplayLineEnds Movementstep = 4
	// paragraphs
	GtkMovementParagraphs Movementstep = 5
	// paragraph_ends
	GtkMovementParagraphEnds Movementstep = 6
	// pages
	GtkMovementPages Movementstep = 7
	// buffer_ends
	GtkMovementBufferEnds Movementstep = 8
	// horizontal_pages
	GtkMovementHorizontalPages Movementstep = 9
)

// Notebooktab is a representation of the C type NotebookTab.
type Notebooktab int

const (
	// first
	GtkNotebookTabFirst Notebooktab = 0
	// last
	GtkNotebookTabLast Notebooktab = 1
)

// Numberuplayout is a representation of the C type NumberUpLayout.
type Numberuplayout int

const (
	// lrtb
	GtkNumberUpLayoutLeftToRightTopToBottom Numberuplayout = 0
	// lrbt
	GtkNumberUpLayoutLeftToRightBottomToTop Numberuplayout = 1
	// rltb
	GtkNumberUpLayoutRightToLeftTopToBottom Numberuplayout = 2
	// rlbt
	GtkNumberUpLayoutRightToLeftBottomToTop Numberuplayout = 3
	// tblr
	GtkNumberUpLayoutTopToBottomLeftToRight Numberuplayout = 4
	// tbrl
	GtkNumberUpLayoutTopToBottomRightToLeft Numberuplayout = 5
	// btlr
	GtkNumberUpLayoutBottomToTopLeftToRight Numberuplayout = 6
	// btrl
	GtkNumberUpLayoutBottomToTopRightToLeft Numberuplayout = 7
)

// Orientation is a representation of the C type Orientation.
type Orientation int

const (
	// horizontal
	GtkOrientationHorizontal Orientation = 0
	// vertical
	GtkOrientationVertical Orientation = 1
)

// Packdirection is a representation of the C type PackDirection.
type Packdirection int

const (
	// ltr
	GtkPackDirectionLtr Packdirection = 0
	// rtl
	GtkPackDirectionRtl Packdirection = 1
	// ttb
	GtkPackDirectionTtb Packdirection = 2
	// btt
	GtkPackDirectionBtt Packdirection = 3
)

// Packtype is a representation of the C type PackType.
type Packtype int

const (
	// start
	GtkPackStart Packtype = 0
	// end
	GtkPackEnd Packtype = 1
)

// Padactiontype is a representation of the C type PadActionType.
type Padactiontype int

const (
	// button
	GtkPadActionButton Padactiontype = 0
	// ring
	GtkPadActionRing Padactiontype = 1
	// strip
	GtkPadActionStrip Padactiontype = 2
)

// Pageorientation is a representation of the C type PageOrientation.
type Pageorientation int

const (
	// portrait
	GtkPageOrientationPortrait Pageorientation = 0
	// landscape
	GtkPageOrientationLandscape Pageorientation = 1
	// reverse_portrait
	GtkPageOrientationReversePortrait Pageorientation = 2
	// reverse_landscape
	GtkPageOrientationReverseLandscape Pageorientation = 3
)

// Pageset is a representation of the C type PageSet.
type Pageset int

const (
	// all
	GtkPageSetAll Pageset = 0
	// even
	GtkPageSetEven Pageset = 1
	// odd
	GtkPageSetOdd Pageset = 2
)

// Pandirection is a representation of the C type PanDirection.
//
// since 3.14
type Pandirection int

const (
	// left
	GtkPanDirectionLeft Pandirection = 0
	// right
	GtkPanDirectionRight Pandirection = 1
	// up
	GtkPanDirectionUp Pandirection = 2
	// down
	GtkPanDirectionDown Pandirection = 3
)

// Pathprioritytype is a representation of the C type PathPriorityType.
type Pathprioritytype int

const (
	// lowest
	GtkPathPrioLowest Pathprioritytype = 0
	// gtk
	GtkPathPrioGtk Pathprioritytype = 4
	// application
	GtkPathPrioApplication Pathprioritytype = 8
	// theme
	GtkPathPrioTheme Pathprioritytype = 10
	// rc
	GtkPathPrioRc Pathprioritytype = 12
	// highest
	GtkPathPrioHighest Pathprioritytype = 15
)

// Pathtype is a representation of the C type PathType.
type Pathtype int

const (
	// widget
	GtkPathWidget Pathtype = 0
	// widget_class
	GtkPathWidgetClass Pathtype = 1
	// class
	GtkPathClass Pathtype = 2
)

// Policytype is a representation of the C type PolicyType.
type Policytype int

const (
	// always
	GtkPolicyAlways Policytype = 0
	// automatic
	GtkPolicyAutomatic Policytype = 1
	// never
	GtkPolicyNever Policytype = 2
	// external
	GtkPolicyExternal Policytype = 3
)

// Popoverconstraint is a representation of the C type PopoverConstraint.
//
// since 3.20
type Popoverconstraint int

const (
	// none
	GtkPopoverConstraintNone Popoverconstraint = 0
	// window
	GtkPopoverConstraintWindow Popoverconstraint = 1
)

// Positiontype is a representation of the C type PositionType.
type Positiontype int

const (
	// left
	GtkPosLeft Positiontype = 0
	// right
	GtkPosRight Positiontype = 1
	// top
	GtkPosTop Positiontype = 2
	// bottom
	GtkPosBottom Positiontype = 3
)

// Printduplex is a representation of the C type PrintDuplex.
type Printduplex int

const (
	// simplex
	GtkPrintDuplexSimplex Printduplex = 0
	// horizontal
	GtkPrintDuplexHorizontal Printduplex = 1
	// vertical
	GtkPrintDuplexVertical Printduplex = 2
)

// Printerror is a representation of the C type PrintError.
type Printerror int

const (
	// general
	GtkPrintErrorGeneral Printerror = 0
	// internal_error
	GtkPrintErrorInternalError Printerror = 1
	// nomem
	GtkPrintErrorNomem Printerror = 2
	// invalid_file
	GtkPrintErrorInvalidFile Printerror = 3
)

// Printoperationaction is a representation of the C type PrintOperationAction.
type Printoperationaction int

const (
	// print_dialog
	GtkPrintOperationActionPrintDialog Printoperationaction = 0
	// print
	GtkPrintOperationActionPrint Printoperationaction = 1
	// preview
	GtkPrintOperationActionPreview Printoperationaction = 2
	// export
	GtkPrintOperationActionExport Printoperationaction = 3
)

// Printoperationresult is a representation of the C type PrintOperationResult.
type Printoperationresult int

const (
	// error
	GtkPrintOperationResultError Printoperationresult = 0
	// apply
	GtkPrintOperationResultApply Printoperationresult = 1
	// cancel
	GtkPrintOperationResultCancel Printoperationresult = 2
	// in_progress
	GtkPrintOperationResultInProgress Printoperationresult = 3
)

// Printpages is a representation of the C type PrintPages.
type Printpages int

const (
	// all
	GtkPrintPagesAll Printpages = 0
	// current
	GtkPrintPagesCurrent Printpages = 1
	// ranges
	GtkPrintPagesRanges Printpages = 2
	// selection
	GtkPrintPagesSelection Printpages = 3
)

// Printquality is a representation of the C type PrintQuality.
type Printquality int

const (
	// low
	GtkPrintQualityLow Printquality = 0
	// normal
	GtkPrintQualityNormal Printquality = 1
	// high
	GtkPrintQualityHigh Printquality = 2
	// draft
	GtkPrintQualityDraft Printquality = 3
)

// Printstatus is a representation of the C type PrintStatus.
type Printstatus int

const (
	// initial
	GtkPrintStatusInitial Printstatus = 0
	// preparing
	GtkPrintStatusPreparing Printstatus = 1
	// generating_data
	GtkPrintStatusGeneratingData Printstatus = 2
	// sending_data
	GtkPrintStatusSendingData Printstatus = 3
	// pending
	GtkPrintStatusPending Printstatus = 4
	// pending_issue
	GtkPrintStatusPendingIssue Printstatus = 5
	// printing
	GtkPrintStatusPrinting Printstatus = 6
	// finished
	GtkPrintStatusFinished Printstatus = 7
	// finished_aborted
	GtkPrintStatusFinishedAborted Printstatus = 8
)

// Propagationphase is a representation of the C type PropagationPhase.
//
// since 3.14
type Propagationphase int

const (
	// none
	GtkPhaseNone Propagationphase = 0
	// capture
	GtkPhaseCapture Propagationphase = 1
	// bubble
	GtkPhaseBubble Propagationphase = 2
	// target
	GtkPhaseTarget Propagationphase = 3
)

// Rctokentype is a representation of the C type RcTokenType.
type Rctokentype int

const (
	// invalid
	GtkRcTokenInvalid Rctokentype = 270
	// include
	GtkRcTokenInclude Rctokentype = 271
	// normal
	GtkRcTokenNormal Rctokentype = 272
	// active
	GtkRcTokenActive Rctokentype = 273
	// prelight
	GtkRcTokenPrelight Rctokentype = 274
	// selected
	GtkRcTokenSelected Rctokentype = 275
	// insensitive
	GtkRcTokenInsensitive Rctokentype = 276
	// fg
	GtkRcTokenFg Rctokentype = 277
	// bg
	GtkRcTokenBg Rctokentype = 278
	// text
	GtkRcTokenText Rctokentype = 279
	// base
	GtkRcTokenBase Rctokentype = 280
	// xthickness
	GtkRcTokenXthickness Rctokentype = 281
	// ythickness
	GtkRcTokenYthickness Rctokentype = 282
	// font
	GtkRcTokenFont Rctokentype = 283
	// fontset
	GtkRcTokenFontset Rctokentype = 284
	// font_name
	GtkRcTokenFontName Rctokentype = 285
	// bg_pixmap
	GtkRcTokenBgPixmap Rctokentype = 286
	// pixmap_path
	GtkRcTokenPixmapPath Rctokentype = 287
	// style
	GtkRcTokenStyle Rctokentype = 288
	// binding
	GtkRcTokenBinding Rctokentype = 289
	// bind
	GtkRcTokenBind Rctokentype = 290
	// widget
	GtkRcTokenWidget Rctokentype = 291
	// widget_class
	GtkRcTokenWidgetClass Rctokentype = 292
	// class
	GtkRcTokenClass Rctokentype = 293
	// lowest
	GtkRcTokenLowest Rctokentype = 294
	// gtk
	GtkRcTokenGtk Rctokentype = 295
	// application
	GtkRcTokenApplication Rctokentype = 296
	// theme
	GtkRcTokenTheme Rctokentype = 297
	// rc
	GtkRcTokenRc Rctokentype = 298
	// highest
	GtkRcTokenHighest Rctokentype = 299
	// engine
	GtkRcTokenEngine Rctokentype = 300
	// module_path
	GtkRcTokenModulePath Rctokentype = 301
	// im_module_path
	GtkRcTokenImModulePath Rctokentype = 302
	// im_module_file
	GtkRcTokenImModuleFile Rctokentype = 303
	// stock
	GtkRcTokenStock Rctokentype = 304
	// ltr
	GtkRcTokenLtr Rctokentype = 305
	// rtl
	GtkRcTokenRtl Rctokentype = 306
	// color
	GtkRcTokenColor Rctokentype = 307
	// unbind
	GtkRcTokenUnbind Rctokentype = 308
	// last
	GtkRcTokenLast Rctokentype = 309
)

// Recentchoosererror is a representation of the C type RecentChooserError.
//
// since 2.10
type Recentchoosererror int

const (
	// not_found
	GtkRecentChooserErrorNotFound Recentchoosererror = 0
	// invalid_uri
	GtkRecentChooserErrorInvalidUri Recentchoosererror = 1
)

// Recentmanagererror is a representation of the C type RecentManagerError.
//
// since 2.10
type Recentmanagererror int

const (
	// not_found
	GtkRecentManagerErrorNotFound Recentmanagererror = 0
	// invalid_uri
	GtkRecentManagerErrorInvalidUri Recentmanagererror = 1
	// invalid_encoding
	GtkRecentManagerErrorInvalidEncoding Recentmanagererror = 2
	// not_registered
	GtkRecentManagerErrorNotRegistered Recentmanagererror = 3
	// read
	GtkRecentManagerErrorRead Recentmanagererror = 4
	// write
	GtkRecentManagerErrorWrite Recentmanagererror = 5
	// unknown
	GtkRecentManagerErrorUnknown Recentmanagererror = 6
)

// Recentsorttype is a representation of the C type RecentSortType.
//
// since 2.10
type Recentsorttype int

const (
	// none
	GtkRecentSortNone Recentsorttype = 0
	// mru
	GtkRecentSortMru Recentsorttype = 1
	// lru
	GtkRecentSortLru Recentsorttype = 2
	// custom
	GtkRecentSortCustom Recentsorttype = 3
)

// Reliefstyle is a representation of the C type ReliefStyle.
type Reliefstyle int

const (
	// normal
	GtkReliefNormal Reliefstyle = 0
	// half
	GtkReliefHalf Reliefstyle = 1
	// none
	GtkReliefNone Reliefstyle = 2
)

// Resizemode is a representation of the C type ResizeMode.
type Resizemode int

const (
	// parent
	GtkResizeParent Resizemode = 0
	// queue
	GtkResizeQueue Resizemode = 1
	// immediate
	GtkResizeImmediate Resizemode = 2
)

// Responsetype is a representation of the C type ResponseType.
type Responsetype int

const (
	// none
	GtkResponseNone Responsetype = -1
	// reject
	GtkResponseReject Responsetype = -2
	// accept
	GtkResponseAccept Responsetype = -3
	// delete_event
	GtkResponseDeleteEvent Responsetype = -4
	// ok
	GtkResponseOk Responsetype = -5
	// cancel
	GtkResponseCancel Responsetype = -6
	// close
	GtkResponseClose Responsetype = -7
	// yes
	GtkResponseYes Responsetype = -8
	// no
	GtkResponseNo Responsetype = -9
	// apply
	GtkResponseApply Responsetype = -10
	// help
	GtkResponseHelp Responsetype = -11
)

// Revealertransitiontype is a representation of the C type RevealerTransitionType.
type Revealertransitiontype int

const (
	// none
	GtkRevealerTransitionTypeNone Revealertransitiontype = 0
	// crossfade
	GtkRevealerTransitionTypeCrossfade Revealertransitiontype = 1
	// slide_right
	GtkRevealerTransitionTypeSlideRight Revealertransitiontype = 2
	// slide_left
	GtkRevealerTransitionTypeSlideLeft Revealertransitiontype = 3
	// slide_up
	GtkRevealerTransitionTypeSlideUp Revealertransitiontype = 4
	// slide_down
	GtkRevealerTransitionTypeSlideDown Revealertransitiontype = 5
)

// Scrollstep is a representation of the C type ScrollStep.
type Scrollstep int

const (
	// steps
	GtkScrollSteps Scrollstep = 0
	// pages
	GtkScrollPages Scrollstep = 1
	// ends
	GtkScrollEnds Scrollstep = 2
	// horizontal_steps
	GtkScrollHorizontalSteps Scrollstep = 3
	// horizontal_pages
	GtkScrollHorizontalPages Scrollstep = 4
	// horizontal_ends
	GtkScrollHorizontalEnds Scrollstep = 5
)

// Scrolltype is a representation of the C type ScrollType.
type Scrolltype int

const (
	// none
	GtkScrollNone Scrolltype = 0
	// jump
	GtkScrollJump Scrolltype = 1
	// step_backward
	GtkScrollStepBackward Scrolltype = 2
	// step_forward
	GtkScrollStepForward Scrolltype = 3
	// page_backward
	GtkScrollPageBackward Scrolltype = 4
	// page_forward
	GtkScrollPageForward Scrolltype = 5
	// step_up
	GtkScrollStepUp Scrolltype = 6
	// step_down
	GtkScrollStepDown Scrolltype = 7
	// page_up
	GtkScrollPageUp Scrolltype = 8
	// page_down
	GtkScrollPageDown Scrolltype = 9
	// step_left
	GtkScrollStepLeft Scrolltype = 10
	// step_right
	GtkScrollStepRight Scrolltype = 11
	// page_left
	GtkScrollPageLeft Scrolltype = 12
	// page_right
	GtkScrollPageRight Scrolltype = 13
	// start
	GtkScrollStart Scrolltype = 14
	// end
	GtkScrollEnd Scrolltype = 15
)

// Scrollablepolicy is a representation of the C type ScrollablePolicy.
type Scrollablepolicy int

const (
	// minimum
	GtkScrollMinimum Scrollablepolicy = 0
	// natural
	GtkScrollNatural Scrollablepolicy = 1
)

// Selectionmode is a representation of the C type SelectionMode.
type Selectionmode int

const (
	// none
	GtkSelectionNone Selectionmode = 0
	// single
	GtkSelectionSingle Selectionmode = 1
	// browse
	GtkSelectionBrowse Selectionmode = 2
	// multiple
	GtkSelectionMultiple Selectionmode = 3
)

// Sensitivitytype is a representation of the C type SensitivityType.
type Sensitivitytype int

const (
	// auto
	GtkSensitivityAuto Sensitivitytype = 0
	// on
	GtkSensitivityOn Sensitivitytype = 1
	// off
	GtkSensitivityOff Sensitivitytype = 2
)

// Shadowtype is a representation of the C type ShadowType.
type Shadowtype int

const (
	// none
	GtkShadowNone Shadowtype = 0
	// in
	GtkShadowIn Shadowtype = 1
	// out
	GtkShadowOut Shadowtype = 2
	// etched_in
	GtkShadowEtchedIn Shadowtype = 3
	// etched_out
	GtkShadowEtchedOut Shadowtype = 4
)

// Shortcuttype is a representation of the C type ShortcutType.
//
// since 3.20
type Shortcuttype int

const (
	// accelerator
	GtkShortcutAccelerator Shortcuttype = 0
	// gesture_pinch
	GtkShortcutGesturePinch Shortcuttype = 1
	// gesture_stretch
	GtkShortcutGestureStretch Shortcuttype = 2
	// gesture_rotate_clockwise
	GtkShortcutGestureRotateClockwise Shortcuttype = 3
	// gesture_rotate_counterclockwise
	GtkShortcutGestureRotateCounterclockwise Shortcuttype = 4
	// gesture_two_finger_swipe_left
	GtkShortcutGestureTwoFingerSwipeLeft Shortcuttype = 5
	// gesture_two_finger_swipe_right
	GtkShortcutGestureTwoFingerSwipeRight Shortcuttype = 6
	// gesture
	GtkShortcutGesture Shortcuttype = 7
)

// Sizegroupmode is a representation of the C type SizeGroupMode.
type Sizegroupmode int

const (
	// none
	GtkSizeGroupNone Sizegroupmode = 0
	// horizontal
	GtkSizeGroupHorizontal Sizegroupmode = 1
	// vertical
	GtkSizeGroupVertical Sizegroupmode = 2
	// both
	GtkSizeGroupBoth Sizegroupmode = 3
)

// Sizerequestmode is a representation of the C type SizeRequestMode.
type Sizerequestmode int

const (
	// height_for_width
	GtkSizeRequestHeightForWidth Sizerequestmode = 0
	// width_for_height
	GtkSizeRequestWidthForHeight Sizerequestmode = 1
	// constant_size
	GtkSizeRequestConstantSize Sizerequestmode = 2
)

// Sorttype is a representation of the C type SortType.
type Sorttype int

const (
	// ascending
	GtkSortAscending Sorttype = 0
	// descending
	GtkSortDescending Sorttype = 1
)

// Spinbuttonupdatepolicy is a representation of the C type SpinButtonUpdatePolicy.
type Spinbuttonupdatepolicy int

const (
	// always
	GtkUpdateAlways Spinbuttonupdatepolicy = 0
	// if_valid
	GtkUpdateIfValid Spinbuttonupdatepolicy = 1
)

// Spintype is a representation of the C type SpinType.
type Spintype int

const (
	// step_forward
	GtkSpinStepForward Spintype = 0
	// step_backward
	GtkSpinStepBackward Spintype = 1
	// page_forward
	GtkSpinPageForward Spintype = 2
	// page_backward
	GtkSpinPageBackward Spintype = 3
	// home
	GtkSpinHome Spintype = 4
	// end
	GtkSpinEnd Spintype = 5
	// user_defined
	GtkSpinUserDefined Spintype = 6
)

// Stacktransitiontype is a representation of the C type StackTransitionType.
type Stacktransitiontype int

const (
	// none
	GtkStackTransitionTypeNone Stacktransitiontype = 0
	// crossfade
	GtkStackTransitionTypeCrossfade Stacktransitiontype = 1
	// slide_right
	GtkStackTransitionTypeSlideRight Stacktransitiontype = 2
	// slide_left
	GtkStackTransitionTypeSlideLeft Stacktransitiontype = 3
	// slide_up
	GtkStackTransitionTypeSlideUp Stacktransitiontype = 4
	// slide_down
	GtkStackTransitionTypeSlideDown Stacktransitiontype = 5
	// slide_left_right
	GtkStackTransitionTypeSlideLeftRight Stacktransitiontype = 6
	// slide_up_down
	GtkStackTransitionTypeSlideUpDown Stacktransitiontype = 7
	// over_up
	GtkStackTransitionTypeOverUp Stacktransitiontype = 8
	// over_down
	GtkStackTransitionTypeOverDown Stacktransitiontype = 9
	// over_left
	GtkStackTransitionTypeOverLeft Stacktransitiontype = 10
	// over_right
	GtkStackTransitionTypeOverRight Stacktransitiontype = 11
	// under_up
	GtkStackTransitionTypeUnderUp Stacktransitiontype = 12
	// under_down
	GtkStackTransitionTypeUnderDown Stacktransitiontype = 13
	// under_left
	GtkStackTransitionTypeUnderLeft Stacktransitiontype = 14
	// under_right
	GtkStackTransitionTypeUnderRight Stacktransitiontype = 15
	// over_up_down
	GtkStackTransitionTypeOverUpDown Stacktransitiontype = 16
	// over_down_up
	GtkStackTransitionTypeOverDownUp Stacktransitiontype = 17
	// over_left_right
	GtkStackTransitionTypeOverLeftRight Stacktransitiontype = 18
	// over_right_left
	GtkStackTransitionTypeOverRightLeft Stacktransitiontype = 19
)

// Statetype is a representation of the C type StateType.
type Statetype int

const (
	// normal
	GtkStateNormal Statetype = 0
	// active
	GtkStateActive Statetype = 1
	// prelight
	GtkStatePrelight Statetype = 2
	// selected
	GtkStateSelected Statetype = 3
	// insensitive
	GtkStateInsensitive Statetype = 4
	// inconsistent
	GtkStateInconsistent Statetype = 5
	// focused
	GtkStateFocused Statetype = 6
)

// Textbuffertargetinfo is a representation of the C type TextBufferTargetInfo.
type Textbuffertargetinfo int

const (
	// buffer_contents
	GtkTextBufferTargetInfoBufferContents Textbuffertargetinfo = -1
	// rich_text
	GtkTextBufferTargetInfoRichText Textbuffertargetinfo = -2
	// text
	GtkTextBufferTargetInfoText Textbuffertargetinfo = -3
)

// Textdirection is a representation of the C type TextDirection.
type Textdirection int

const (
	// none
	GtkTextDirNone Textdirection = 0
	// ltr
	GtkTextDirLtr Textdirection = 1
	// rtl
	GtkTextDirRtl Textdirection = 2
)

// Textextendselection is a representation of the C type TextExtendSelection.
//
// since 3.16
type Textextendselection int

const (
	// word
	GtkTextExtendSelectionWord Textextendselection = 0
	// line
	GtkTextExtendSelectionLine Textextendselection = 1
)

// Textviewlayer is a representation of the C type TextViewLayer.
type Textviewlayer int

const (
	// below
	GtkTextViewLayerBelow Textviewlayer = 0
	// above
	GtkTextViewLayerAbove Textviewlayer = 1
	// below_text
	GtkTextViewLayerBelowText Textviewlayer = 2
	// above_text
	GtkTextViewLayerAboveText Textviewlayer = 3
)

// Textwindowtype is a representation of the C type TextWindowType.
type Textwindowtype int

const (
	// private
	GtkTextWindowPrivate Textwindowtype = 0
	// widget
	GtkTextWindowWidget Textwindowtype = 1
	// text
	GtkTextWindowText Textwindowtype = 2
	// left
	GtkTextWindowLeft Textwindowtype = 3
	// right
	GtkTextWindowRight Textwindowtype = 4
	// top
	GtkTextWindowTop Textwindowtype = 5
	// bottom
	GtkTextWindowBottom Textwindowtype = 6
)

// Toolbarspacestyle is a representation of the C type ToolbarSpaceStyle.
type Toolbarspacestyle int

const (
	// empty
	GtkToolbarSpaceEmpty Toolbarspacestyle = 0
	// line
	GtkToolbarSpaceLine Toolbarspacestyle = 1
)

// Toolbarstyle is a representation of the C type ToolbarStyle.
type Toolbarstyle int

const (
	// icons
	GtkToolbarIcons Toolbarstyle = 0
	// text
	GtkToolbarText Toolbarstyle = 1
	// both
	GtkToolbarBoth Toolbarstyle = 2
	// both_horiz
	GtkToolbarBothHoriz Toolbarstyle = 3
)

// Treeviewcolumnsizing is a representation of the C type TreeViewColumnSizing.
type Treeviewcolumnsizing int

const (
	// grow_only
	GtkTreeViewColumnGrowOnly Treeviewcolumnsizing = 0
	// autosize
	GtkTreeViewColumnAutosize Treeviewcolumnsizing = 1
	// fixed
	GtkTreeViewColumnFixed Treeviewcolumnsizing = 2
)

// Treeviewdropposition is a representation of the C type TreeViewDropPosition.
type Treeviewdropposition int

const (
	// before
	GtkTreeViewDropBefore Treeviewdropposition = 0
	// after
	GtkTreeViewDropAfter Treeviewdropposition = 1
	// into_or_before
	GtkTreeViewDropIntoOrBefore Treeviewdropposition = 2
	// into_or_after
	GtkTreeViewDropIntoOrAfter Treeviewdropposition = 3
)

// Treeviewgridlines is a representation of the C type TreeViewGridLines.
type Treeviewgridlines int

const (
	// none
	GtkTreeViewGridLinesNone Treeviewgridlines = 0
	// horizontal
	GtkTreeViewGridLinesHorizontal Treeviewgridlines = 1
	// vertical
	GtkTreeViewGridLinesVertical Treeviewgridlines = 2
	// both
	GtkTreeViewGridLinesBoth Treeviewgridlines = 3
)

// Unit is a representation of the C type Unit.
type Unit int

const (
	// none
	GtkUnitNone Unit = 0
	// points
	GtkUnitPoints Unit = 1
	// inch
	GtkUnitInch Unit = 2
	// mm
	GtkUnitMm Unit = 3
)

// Widgethelptype is a representation of the C type WidgetHelpType.
type Widgethelptype int

const (
	// tooltip
	GtkWidgetHelpTooltip Widgethelptype = 0
	// whats_this
	GtkWidgetHelpWhatsThis Widgethelptype = 1
)

// Windowposition is a representation of the C type WindowPosition.
type Windowposition int

const (
	// none
	GtkWinPosNone Windowposition = 0
	// center
	GtkWinPosCenter Windowposition = 1
	// mouse
	GtkWinPosMouse Windowposition = 2
	// center_always
	GtkWinPosCenterAlways Windowposition = 3
	// center_on_parent
	GtkWinPosCenterOnParent Windowposition = 4
)

// Windowtype is a representation of the C type WindowType.
type Windowtype int

const (
	// toplevel
	GtkWindowToplevel Windowtype = 0
	// popup
	GtkWindowPopup Windowtype = 1
)

// Wrapmode is a representation of the C type WrapMode.
type Wrapmode int

const (
	// none
	GtkWrapNone Wrapmode = 0
	// char
	GtkWrapChar Wrapmode = 1
	// word
	GtkWrapWord Wrapmode = 2
	// word_char
	GtkWrapWordChar Wrapmode = 3
)
