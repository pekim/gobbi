// Code generated - DO NOT EDIT.

package gtk

// Align is a representation of the C type Align.
type Align int

const (
	Align_Fill     Align = 0
	Align_Start    Align = 1
	Align_End      Align = 2
	Align_Center   Align = 3
	Align_Baseline Align = 4
)

// Arrowplacement is a representation of the C type ArrowPlacement.
type Arrowplacement int

const (
	Arrowplacement_Both  Arrowplacement = 0
	Arrowplacement_Start Arrowplacement = 1
	Arrowplacement_End   Arrowplacement = 2
)

// Arrowtype is a representation of the C type ArrowType.
type Arrowtype int

const (
	Arrowtype_Up    Arrowtype = 0
	Arrowtype_Down  Arrowtype = 1
	Arrowtype_Left  Arrowtype = 2
	Arrowtype_Right Arrowtype = 3
	Arrowtype_None  Arrowtype = 4
)

// Assistantpagetype is a representation of the C type AssistantPageType.
type Assistantpagetype int

const (
	Assistantpagetype_Content  Assistantpagetype = 0
	Assistantpagetype_Intro    Assistantpagetype = 1
	Assistantpagetype_Confirm  Assistantpagetype = 2
	Assistantpagetype_Summary  Assistantpagetype = 3
	Assistantpagetype_Progress Assistantpagetype = 4
	Assistantpagetype_Custom   Assistantpagetype = 5
)

// Baselineposition is a representation of the C type BaselinePosition.
//
// since 3.10
type Baselineposition int

const (
	Baselineposition_Top    Baselineposition = 0
	Baselineposition_Center Baselineposition = 1
	Baselineposition_Bottom Baselineposition = 2
)

// Borderstyle is a representation of the C type BorderStyle.
type Borderstyle int

const (
	Borderstyle_None   Borderstyle = 0
	Borderstyle_Solid  Borderstyle = 1
	Borderstyle_Inset  Borderstyle = 2
	Borderstyle_Outset Borderstyle = 3
	Borderstyle_Hidden Borderstyle = 4
	Borderstyle_Dotted Borderstyle = 5
	Borderstyle_Dashed Borderstyle = 6
	Borderstyle_Double Borderstyle = 7
	Borderstyle_Groove Borderstyle = 8
	Borderstyle_Ridge  Borderstyle = 9
)

// Buildererror is a representation of the C type BuilderError.
type Buildererror int

const (
	Buildererror_InvalidTypeFunction  Buildererror = 0
	Buildererror_UnhandledTag         Buildererror = 1
	Buildererror_MissingAttribute     Buildererror = 2
	Buildererror_InvalidAttribute     Buildererror = 3
	Buildererror_InvalidTag           Buildererror = 4
	Buildererror_MissingPropertyValue Buildererror = 5
	Buildererror_InvalidValue         Buildererror = 6
	Buildererror_VersionMismatch      Buildererror = 7
	Buildererror_DuplicateId          Buildererror = 8
	Buildererror_ObjectTypeRefused    Buildererror = 9
	Buildererror_TemplateMismatch     Buildererror = 10
	Buildererror_InvalidProperty      Buildererror = 11
	Buildererror_InvalidSignal        Buildererror = 12
	Buildererror_InvalidId            Buildererror = 13
)

// Buttonboxstyle is a representation of the C type ButtonBoxStyle.
type Buttonboxstyle int

const (
	Buttonboxstyle_Spread Buttonboxstyle = 1
	Buttonboxstyle_Edge   Buttonboxstyle = 2
	Buttonboxstyle_Start  Buttonboxstyle = 3
	Buttonboxstyle_End    Buttonboxstyle = 4
	Buttonboxstyle_Center Buttonboxstyle = 5
	Buttonboxstyle_Expand Buttonboxstyle = 6
)

// Buttonrole is a representation of the C type ButtonRole.
type Buttonrole int

const (
	Buttonrole_Normal Buttonrole = 0
	Buttonrole_Check  Buttonrole = 1
	Buttonrole_Radio  Buttonrole = 2
)

// Buttonstype is a representation of the C type ButtonsType.
type Buttonstype int

const (
	Buttonstype_None     Buttonstype = 0
	Buttonstype_Ok       Buttonstype = 1
	Buttonstype_Close    Buttonstype = 2
	Buttonstype_Cancel   Buttonstype = 3
	Buttonstype_YesNo    Buttonstype = 4
	Buttonstype_OkCancel Buttonstype = 5
)

// Cellrendereraccelmode is a representation of the C type CellRendererAccelMode.
type Cellrendereraccelmode int

const (
	Cellrendereraccelmode_Gtk         Cellrendereraccelmode = 0
	Cellrendereraccelmode_Other       Cellrendereraccelmode = 1
	Cellrendereraccelmode_ModifierTap Cellrendereraccelmode = 2
)

// Cellrenderermode is a representation of the C type CellRendererMode.
type Cellrenderermode int

const (
	Cellrenderermode_Inert       Cellrenderermode = 0
	Cellrenderermode_Activatable Cellrenderermode = 1
	Cellrenderermode_Editable    Cellrenderermode = 2
)

// Cornertype is a representation of the C type CornerType.
type Cornertype int

const (
	Cornertype_TopLeft     Cornertype = 0
	Cornertype_BottomLeft  Cornertype = 1
	Cornertype_TopRight    Cornertype = 2
	Cornertype_BottomRight Cornertype = 3
)

// Cssprovidererror is a representation of the C type CssProviderError.
type Cssprovidererror int

const (
	Cssprovidererror_Failed       Cssprovidererror = 0
	Cssprovidererror_Syntax       Cssprovidererror = 1
	Cssprovidererror_Import       Cssprovidererror = 2
	Cssprovidererror_Name         Cssprovidererror = 3
	Cssprovidererror_Deprecated   Cssprovidererror = 4
	Cssprovidererror_UnknownValue Cssprovidererror = 5
)

// Csssectiontype is a representation of the C type CssSectionType.
//
// since 3.2
type Csssectiontype int

const (
	Csssectiontype_Document        Csssectiontype = 0
	Csssectiontype_Import          Csssectiontype = 1
	Csssectiontype_ColorDefinition Csssectiontype = 2
	Csssectiontype_BindingSet      Csssectiontype = 3
	Csssectiontype_Ruleset         Csssectiontype = 4
	Csssectiontype_Selector        Csssectiontype = 5
	Csssectiontype_Declaration     Csssectiontype = 6
	Csssectiontype_Value           Csssectiontype = 7
	Csssectiontype_Keyframes       Csssectiontype = 8
)

// Deletetype is a representation of the C type DeleteType.
type Deletetype int

const (
	Deletetype_Chars           Deletetype = 0
	Deletetype_WordEnds        Deletetype = 1
	Deletetype_Words           Deletetype = 2
	Deletetype_DisplayLines    Deletetype = 3
	Deletetype_DisplayLineEnds Deletetype = 4
	Deletetype_ParagraphEnds   Deletetype = 5
	Deletetype_Paragraphs      Deletetype = 6
	Deletetype_Whitespace      Deletetype = 7
)

// Directiontype is a representation of the C type DirectionType.
type Directiontype int

const (
	Directiontype_TabForward  Directiontype = 0
	Directiontype_TabBackward Directiontype = 1
	Directiontype_Up          Directiontype = 2
	Directiontype_Down        Directiontype = 3
	Directiontype_Left        Directiontype = 4
	Directiontype_Right       Directiontype = 5
)

// Dragresult is a representation of the C type DragResult.
type Dragresult int

const (
	Dragresult_Success        Dragresult = 0
	Dragresult_NoTarget       Dragresult = 1
	Dragresult_UserCancelled  Dragresult = 2
	Dragresult_TimeoutExpired Dragresult = 3
	Dragresult_GrabBroken     Dragresult = 4
	Dragresult_Error          Dragresult = 5
)

// Entryiconposition is a representation of the C type EntryIconPosition.
//
// since 2.16
type Entryiconposition int

const (
	Entryiconposition_Primary   Entryiconposition = 0
	Entryiconposition_Secondary Entryiconposition = 1
)

// Eventsequencestate is a representation of the C type EventSequenceState.
//
// since 3.14
type Eventsequencestate int

const (
	Eventsequencestate_None    Eventsequencestate = 0
	Eventsequencestate_Claimed Eventsequencestate = 1
	Eventsequencestate_Denied  Eventsequencestate = 2
)

// Expanderstyle is a representation of the C type ExpanderStyle.
type Expanderstyle int

const (
	Expanderstyle_Collapsed     Expanderstyle = 0
	Expanderstyle_SemiCollapsed Expanderstyle = 1
	Expanderstyle_SemiExpanded  Expanderstyle = 2
	Expanderstyle_Expanded      Expanderstyle = 3
)

// Filechooseraction is a representation of the C type FileChooserAction.
type Filechooseraction int

const (
	Filechooseraction_Open         Filechooseraction = 0
	Filechooseraction_Save         Filechooseraction = 1
	Filechooseraction_SelectFolder Filechooseraction = 2
	Filechooseraction_CreateFolder Filechooseraction = 3
)

// Filechooserconfirmation is a representation of the C type FileChooserConfirmation.
//
// since 2.8
type Filechooserconfirmation int

const (
	Filechooserconfirmation_Confirm        Filechooserconfirmation = 0
	Filechooserconfirmation_AcceptFilename Filechooserconfirmation = 1
	Filechooserconfirmation_SelectAgain    Filechooserconfirmation = 2
)

// Filechoosererror is a representation of the C type FileChooserError.
type Filechoosererror int

const (
	Filechoosererror_Nonexistent        Filechoosererror = 0
	Filechoosererror_BadFilename        Filechoosererror = 1
	Filechoosererror_AlreadyExists      Filechoosererror = 2
	Filechoosererror_IncompleteHostname Filechoosererror = 3
)

// Impreeditstyle is a representation of the C type IMPreeditStyle.
type Impreeditstyle int

const (
	Impreeditstyle_Nothing  Impreeditstyle = 0
	Impreeditstyle_Callback Impreeditstyle = 1
	Impreeditstyle_None     Impreeditstyle = 2
)

// Imstatusstyle is a representation of the C type IMStatusStyle.
type Imstatusstyle int

const (
	Imstatusstyle_Nothing  Imstatusstyle = 0
	Imstatusstyle_Callback Imstatusstyle = 1
	Imstatusstyle_None     Imstatusstyle = 2
)

// Iconsize is a representation of the C type IconSize.
type Iconsize int

const (
	Iconsize_Invalid      Iconsize = 0
	Iconsize_Menu         Iconsize = 1
	Iconsize_SmallToolbar Iconsize = 2
	Iconsize_LargeToolbar Iconsize = 3
	Iconsize_Button       Iconsize = 4
	Iconsize_Dnd          Iconsize = 5
	Iconsize_Dialog       Iconsize = 6
)

// Iconthemeerror is a representation of the C type IconThemeError.
type Iconthemeerror int

const (
	Iconthemeerror_NotFound Iconthemeerror = 0
	Iconthemeerror_Failed   Iconthemeerror = 1
)

// Iconviewdropposition is a representation of the C type IconViewDropPosition.
type Iconviewdropposition int

const (
	Iconviewdropposition_NoDrop    Iconviewdropposition = 0
	Iconviewdropposition_DropInto  Iconviewdropposition = 1
	Iconviewdropposition_DropLeft  Iconviewdropposition = 2
	Iconviewdropposition_DropRight Iconviewdropposition = 3
	Iconviewdropposition_DropAbove Iconviewdropposition = 4
	Iconviewdropposition_DropBelow Iconviewdropposition = 5
)

// Imagetype is a representation of the C type ImageType.
type Imagetype int

const (
	Imagetype_Empty     Imagetype = 0
	Imagetype_Pixbuf    Imagetype = 1
	Imagetype_Stock     Imagetype = 2
	Imagetype_IconSet   Imagetype = 3
	Imagetype_Animation Imagetype = 4
	Imagetype_IconName  Imagetype = 5
	Imagetype_Gicon     Imagetype = 6
	Imagetype_Surface   Imagetype = 7
)

// Inputpurpose is a representation of the C type InputPurpose.
//
// since 3.6
type Inputpurpose int

const (
	Inputpurpose_FreeForm Inputpurpose = 0
	Inputpurpose_Alpha    Inputpurpose = 1
	Inputpurpose_Digits   Inputpurpose = 2
	Inputpurpose_Number   Inputpurpose = 3
	Inputpurpose_Phone    Inputpurpose = 4
	Inputpurpose_Url      Inputpurpose = 5
	Inputpurpose_Email    Inputpurpose = 6
	Inputpurpose_Name     Inputpurpose = 7
	Inputpurpose_Password Inputpurpose = 8
	Inputpurpose_Pin      Inputpurpose = 9
)

// Justification is a representation of the C type Justification.
type Justification int

const (
	Justification_Left   Justification = 0
	Justification_Right  Justification = 1
	Justification_Center Justification = 2
	Justification_Fill   Justification = 3
)

// Levelbarmode is a representation of the C type LevelBarMode.
//
// since 3.6
type Levelbarmode int

const (
	Levelbarmode_Continuous Levelbarmode = 0
	Levelbarmode_Discrete   Levelbarmode = 1
)

// License is a representation of the C type License.
//
// since 3.0
type License int

const (
	License_Unknown    License = 0
	License_Custom     License = 1
	License_Gpl20      License = 2
	License_Gpl30      License = 3
	License_Lgpl21     License = 4
	License_Lgpl30     License = 5
	License_Bsd        License = 6
	License_MitX11     License = 7
	License_Artistic   License = 8
	License_Gpl20Only  License = 9
	License_Gpl30Only  License = 10
	License_Lgpl21Only License = 11
	License_Lgpl30Only License = 12
	License_Agpl30     License = 13
	License_Agpl30Only License = 14
)

// Menudirectiontype is a representation of the C type MenuDirectionType.
type Menudirectiontype int

const (
	Menudirectiontype_Parent Menudirectiontype = 0
	Menudirectiontype_Child  Menudirectiontype = 1
	Menudirectiontype_Next   Menudirectiontype = 2
	Menudirectiontype_Prev   Menudirectiontype = 3
)

// Messagetype is a representation of the C type MessageType.
type Messagetype int

const (
	Messagetype_Info     Messagetype = 0
	Messagetype_Warning  Messagetype = 1
	Messagetype_Question Messagetype = 2
	Messagetype_Error    Messagetype = 3
	Messagetype_Other    Messagetype = 4
)

// Movementstep is a representation of the C type MovementStep.
type Movementstep int

const (
	Movementstep_LogicalPositions Movementstep = 0
	Movementstep_VisualPositions  Movementstep = 1
	Movementstep_Words            Movementstep = 2
	Movementstep_DisplayLines     Movementstep = 3
	Movementstep_DisplayLineEnds  Movementstep = 4
	Movementstep_Paragraphs       Movementstep = 5
	Movementstep_ParagraphEnds    Movementstep = 6
	Movementstep_Pages            Movementstep = 7
	Movementstep_BufferEnds       Movementstep = 8
	Movementstep_HorizontalPages  Movementstep = 9
)

// Notebooktab is a representation of the C type NotebookTab.
type Notebooktab int

const (
	Notebooktab_First Notebooktab = 0
	Notebooktab_Last  Notebooktab = 1
)

// Numberuplayout is a representation of the C type NumberUpLayout.
type Numberuplayout int

const (
	Numberuplayout_Lrtb Numberuplayout = 0
	Numberuplayout_Lrbt Numberuplayout = 1
	Numberuplayout_Rltb Numberuplayout = 2
	Numberuplayout_Rlbt Numberuplayout = 3
	Numberuplayout_Tblr Numberuplayout = 4
	Numberuplayout_Tbrl Numberuplayout = 5
	Numberuplayout_Btlr Numberuplayout = 6
	Numberuplayout_Btrl Numberuplayout = 7
)

// Orientation is a representation of the C type Orientation.
type Orientation int

const (
	Orientation_Horizontal Orientation = 0
	Orientation_Vertical   Orientation = 1
)

// Packdirection is a representation of the C type PackDirection.
type Packdirection int

const (
	Packdirection_Ltr Packdirection = 0
	Packdirection_Rtl Packdirection = 1
	Packdirection_Ttb Packdirection = 2
	Packdirection_Btt Packdirection = 3
)

// Packtype is a representation of the C type PackType.
type Packtype int

const (
	Packtype_Start Packtype = 0
	Packtype_End   Packtype = 1
)

// Padactiontype is a representation of the C type PadActionType.
type Padactiontype int

const (
	Padactiontype_Button Padactiontype = 0
	Padactiontype_Ring   Padactiontype = 1
	Padactiontype_Strip  Padactiontype = 2
)

// Pageorientation is a representation of the C type PageOrientation.
type Pageorientation int

const (
	Pageorientation_Portrait         Pageorientation = 0
	Pageorientation_Landscape        Pageorientation = 1
	Pageorientation_ReversePortrait  Pageorientation = 2
	Pageorientation_ReverseLandscape Pageorientation = 3
)

// Pageset is a representation of the C type PageSet.
type Pageset int

const (
	Pageset_All  Pageset = 0
	Pageset_Even Pageset = 1
	Pageset_Odd  Pageset = 2
)

// Pandirection is a representation of the C type PanDirection.
//
// since 3.14
type Pandirection int

const (
	Pandirection_Left  Pandirection = 0
	Pandirection_Right Pandirection = 1
	Pandirection_Up    Pandirection = 2
	Pandirection_Down  Pandirection = 3
)

// Pathprioritytype is a representation of the C type PathPriorityType.
type Pathprioritytype int

const (
	Pathprioritytype_Lowest      Pathprioritytype = 0
	Pathprioritytype_Gtk         Pathprioritytype = 4
	Pathprioritytype_Application Pathprioritytype = 8
	Pathprioritytype_Theme       Pathprioritytype = 10
	Pathprioritytype_Rc          Pathprioritytype = 12
	Pathprioritytype_Highest     Pathprioritytype = 15
)

// Pathtype is a representation of the C type PathType.
type Pathtype int

const (
	Pathtype_Widget      Pathtype = 0
	Pathtype_WidgetClass Pathtype = 1
	Pathtype_Class       Pathtype = 2
)

// Policytype is a representation of the C type PolicyType.
type Policytype int

const (
	Policytype_Always    Policytype = 0
	Policytype_Automatic Policytype = 1
	Policytype_Never     Policytype = 2
	Policytype_External  Policytype = 3
)

// Popoverconstraint is a representation of the C type PopoverConstraint.
//
// since 3.20
type Popoverconstraint int

const (
	Popoverconstraint_None   Popoverconstraint = 0
	Popoverconstraint_Window Popoverconstraint = 1
)

// Positiontype is a representation of the C type PositionType.
type Positiontype int

const (
	Positiontype_Left   Positiontype = 0
	Positiontype_Right  Positiontype = 1
	Positiontype_Top    Positiontype = 2
	Positiontype_Bottom Positiontype = 3
)

// Printduplex is a representation of the C type PrintDuplex.
type Printduplex int

const (
	Printduplex_Simplex    Printduplex = 0
	Printduplex_Horizontal Printduplex = 1
	Printduplex_Vertical   Printduplex = 2
)

// Printerror is a representation of the C type PrintError.
type Printerror int

const (
	Printerror_General       Printerror = 0
	Printerror_InternalError Printerror = 1
	Printerror_Nomem         Printerror = 2
	Printerror_InvalidFile   Printerror = 3
)

// Printoperationaction is a representation of the C type PrintOperationAction.
type Printoperationaction int

const (
	Printoperationaction_PrintDialog Printoperationaction = 0
	Printoperationaction_Print       Printoperationaction = 1
	Printoperationaction_Preview     Printoperationaction = 2
	Printoperationaction_Export      Printoperationaction = 3
)

// Printoperationresult is a representation of the C type PrintOperationResult.
type Printoperationresult int

const (
	Printoperationresult_Error      Printoperationresult = 0
	Printoperationresult_Apply      Printoperationresult = 1
	Printoperationresult_Cancel     Printoperationresult = 2
	Printoperationresult_InProgress Printoperationresult = 3
)

// Printpages is a representation of the C type PrintPages.
type Printpages int

const (
	Printpages_All       Printpages = 0
	Printpages_Current   Printpages = 1
	Printpages_Ranges    Printpages = 2
	Printpages_Selection Printpages = 3
)

// Printquality is a representation of the C type PrintQuality.
type Printquality int

const (
	Printquality_Low    Printquality = 0
	Printquality_Normal Printquality = 1
	Printquality_High   Printquality = 2
	Printquality_Draft  Printquality = 3
)

// Printstatus is a representation of the C type PrintStatus.
type Printstatus int

const (
	Printstatus_Initial         Printstatus = 0
	Printstatus_Preparing       Printstatus = 1
	Printstatus_GeneratingData  Printstatus = 2
	Printstatus_SendingData     Printstatus = 3
	Printstatus_Pending         Printstatus = 4
	Printstatus_PendingIssue    Printstatus = 5
	Printstatus_Printing        Printstatus = 6
	Printstatus_Finished        Printstatus = 7
	Printstatus_FinishedAborted Printstatus = 8
)

// Propagationphase is a representation of the C type PropagationPhase.
//
// since 3.14
type Propagationphase int

const (
	Propagationphase_None    Propagationphase = 0
	Propagationphase_Capture Propagationphase = 1
	Propagationphase_Bubble  Propagationphase = 2
	Propagationphase_Target  Propagationphase = 3
)

// Rctokentype is a representation of the C type RcTokenType.
type Rctokentype int

const (
	Rctokentype_Invalid      Rctokentype = 270
	Rctokentype_Include      Rctokentype = 271
	Rctokentype_Normal       Rctokentype = 272
	Rctokentype_Active       Rctokentype = 273
	Rctokentype_Prelight     Rctokentype = 274
	Rctokentype_Selected     Rctokentype = 275
	Rctokentype_Insensitive  Rctokentype = 276
	Rctokentype_Fg           Rctokentype = 277
	Rctokentype_Bg           Rctokentype = 278
	Rctokentype_Text         Rctokentype = 279
	Rctokentype_Base         Rctokentype = 280
	Rctokentype_Xthickness   Rctokentype = 281
	Rctokentype_Ythickness   Rctokentype = 282
	Rctokentype_Font         Rctokentype = 283
	Rctokentype_Fontset      Rctokentype = 284
	Rctokentype_FontName     Rctokentype = 285
	Rctokentype_BgPixmap     Rctokentype = 286
	Rctokentype_PixmapPath   Rctokentype = 287
	Rctokentype_Style        Rctokentype = 288
	Rctokentype_Binding      Rctokentype = 289
	Rctokentype_Bind         Rctokentype = 290
	Rctokentype_Widget       Rctokentype = 291
	Rctokentype_WidgetClass  Rctokentype = 292
	Rctokentype_Class        Rctokentype = 293
	Rctokentype_Lowest       Rctokentype = 294
	Rctokentype_Gtk          Rctokentype = 295
	Rctokentype_Application  Rctokentype = 296
	Rctokentype_Theme        Rctokentype = 297
	Rctokentype_Rc           Rctokentype = 298
	Rctokentype_Highest      Rctokentype = 299
	Rctokentype_Engine       Rctokentype = 300
	Rctokentype_ModulePath   Rctokentype = 301
	Rctokentype_ImModulePath Rctokentype = 302
	Rctokentype_ImModuleFile Rctokentype = 303
	Rctokentype_Stock        Rctokentype = 304
	Rctokentype_Ltr          Rctokentype = 305
	Rctokentype_Rtl          Rctokentype = 306
	Rctokentype_Color        Rctokentype = 307
	Rctokentype_Unbind       Rctokentype = 308
	Rctokentype_Last         Rctokentype = 309
)

// Recentchoosererror is a representation of the C type RecentChooserError.
//
// since 2.10
type Recentchoosererror int

const (
	Recentchoosererror_NotFound   Recentchoosererror = 0
	Recentchoosererror_InvalidUri Recentchoosererror = 1
)

// Recentmanagererror is a representation of the C type RecentManagerError.
//
// since 2.10
type Recentmanagererror int

const (
	Recentmanagererror_NotFound        Recentmanagererror = 0
	Recentmanagererror_InvalidUri      Recentmanagererror = 1
	Recentmanagererror_InvalidEncoding Recentmanagererror = 2
	Recentmanagererror_NotRegistered   Recentmanagererror = 3
	Recentmanagererror_Read            Recentmanagererror = 4
	Recentmanagererror_Write           Recentmanagererror = 5
	Recentmanagererror_Unknown         Recentmanagererror = 6
)

// Recentsorttype is a representation of the C type RecentSortType.
//
// since 2.10
type Recentsorttype int

const (
	Recentsorttype_None   Recentsorttype = 0
	Recentsorttype_Mru    Recentsorttype = 1
	Recentsorttype_Lru    Recentsorttype = 2
	Recentsorttype_Custom Recentsorttype = 3
)

// Reliefstyle is a representation of the C type ReliefStyle.
type Reliefstyle int

const (
	Reliefstyle_Normal Reliefstyle = 0
	Reliefstyle_Half   Reliefstyle = 1
	Reliefstyle_None   Reliefstyle = 2
)

// Resizemode is a representation of the C type ResizeMode.
type Resizemode int

const (
	Resizemode_Parent    Resizemode = 0
	Resizemode_Queue     Resizemode = 1
	Resizemode_Immediate Resizemode = 2
)

// Responsetype is a representation of the C type ResponseType.
type Responsetype int

const (
	Responsetype_None        Responsetype = -1
	Responsetype_Reject      Responsetype = -2
	Responsetype_Accept      Responsetype = -3
	Responsetype_DeleteEvent Responsetype = -4
	Responsetype_Ok          Responsetype = -5
	Responsetype_Cancel      Responsetype = -6
	Responsetype_Close       Responsetype = -7
	Responsetype_Yes         Responsetype = -8
	Responsetype_No          Responsetype = -9
	Responsetype_Apply       Responsetype = -10
	Responsetype_Help        Responsetype = -11
)

// Revealertransitiontype is a representation of the C type RevealerTransitionType.
type Revealertransitiontype int

const (
	Revealertransitiontype_None       Revealertransitiontype = 0
	Revealertransitiontype_Crossfade  Revealertransitiontype = 1
	Revealertransitiontype_SlideRight Revealertransitiontype = 2
	Revealertransitiontype_SlideLeft  Revealertransitiontype = 3
	Revealertransitiontype_SlideUp    Revealertransitiontype = 4
	Revealertransitiontype_SlideDown  Revealertransitiontype = 5
)

// Scrollstep is a representation of the C type ScrollStep.
type Scrollstep int

const (
	Scrollstep_Steps           Scrollstep = 0
	Scrollstep_Pages           Scrollstep = 1
	Scrollstep_Ends            Scrollstep = 2
	Scrollstep_HorizontalSteps Scrollstep = 3
	Scrollstep_HorizontalPages Scrollstep = 4
	Scrollstep_HorizontalEnds  Scrollstep = 5
)

// Scrolltype is a representation of the C type ScrollType.
type Scrolltype int

const (
	Scrolltype_None         Scrolltype = 0
	Scrolltype_Jump         Scrolltype = 1
	Scrolltype_StepBackward Scrolltype = 2
	Scrolltype_StepForward  Scrolltype = 3
	Scrolltype_PageBackward Scrolltype = 4
	Scrolltype_PageForward  Scrolltype = 5
	Scrolltype_StepUp       Scrolltype = 6
	Scrolltype_StepDown     Scrolltype = 7
	Scrolltype_PageUp       Scrolltype = 8
	Scrolltype_PageDown     Scrolltype = 9
	Scrolltype_StepLeft     Scrolltype = 10
	Scrolltype_StepRight    Scrolltype = 11
	Scrolltype_PageLeft     Scrolltype = 12
	Scrolltype_PageRight    Scrolltype = 13
	Scrolltype_Start        Scrolltype = 14
	Scrolltype_End          Scrolltype = 15
)

// Scrollablepolicy is a representation of the C type ScrollablePolicy.
type Scrollablepolicy int

const (
	Scrollablepolicy_Minimum Scrollablepolicy = 0
	Scrollablepolicy_Natural Scrollablepolicy = 1
)

// Selectionmode is a representation of the C type SelectionMode.
type Selectionmode int

const (
	Selectionmode_None     Selectionmode = 0
	Selectionmode_Single   Selectionmode = 1
	Selectionmode_Browse   Selectionmode = 2
	Selectionmode_Multiple Selectionmode = 3
)

// Sensitivitytype is a representation of the C type SensitivityType.
type Sensitivitytype int

const (
	Sensitivitytype_Auto Sensitivitytype = 0
	Sensitivitytype_On   Sensitivitytype = 1
	Sensitivitytype_Off  Sensitivitytype = 2
)

// Shadowtype is a representation of the C type ShadowType.
type Shadowtype int

const (
	Shadowtype_None      Shadowtype = 0
	Shadowtype_In        Shadowtype = 1
	Shadowtype_Out       Shadowtype = 2
	Shadowtype_EtchedIn  Shadowtype = 3
	Shadowtype_EtchedOut Shadowtype = 4
)

// Shortcuttype is a representation of the C type ShortcutType.
//
// since 3.20
type Shortcuttype int

const (
	Shortcuttype_Accelerator                   Shortcuttype = 0
	Shortcuttype_GesturePinch                  Shortcuttype = 1
	Shortcuttype_GestureStretch                Shortcuttype = 2
	Shortcuttype_GestureRotateClockwise        Shortcuttype = 3
	Shortcuttype_GestureRotateCounterclockwise Shortcuttype = 4
	Shortcuttype_GestureTwoFingerSwipeLeft     Shortcuttype = 5
	Shortcuttype_GestureTwoFingerSwipeRight    Shortcuttype = 6
	Shortcuttype_Gesture                       Shortcuttype = 7
)

// Sizegroupmode is a representation of the C type SizeGroupMode.
type Sizegroupmode int

const (
	Sizegroupmode_None       Sizegroupmode = 0
	Sizegroupmode_Horizontal Sizegroupmode = 1
	Sizegroupmode_Vertical   Sizegroupmode = 2
	Sizegroupmode_Both       Sizegroupmode = 3
)

// Sizerequestmode is a representation of the C type SizeRequestMode.
type Sizerequestmode int

const (
	Sizerequestmode_HeightForWidth Sizerequestmode = 0
	Sizerequestmode_WidthForHeight Sizerequestmode = 1
	Sizerequestmode_ConstantSize   Sizerequestmode = 2
)

// Sorttype is a representation of the C type SortType.
type Sorttype int

const (
	Sorttype_Ascending  Sorttype = 0
	Sorttype_Descending Sorttype = 1
)

// Spinbuttonupdatepolicy is a representation of the C type SpinButtonUpdatePolicy.
type Spinbuttonupdatepolicy int

const (
	Spinbuttonupdatepolicy_Always  Spinbuttonupdatepolicy = 0
	Spinbuttonupdatepolicy_IfValid Spinbuttonupdatepolicy = 1
)

// Spintype is a representation of the C type SpinType.
type Spintype int

const (
	Spintype_StepForward  Spintype = 0
	Spintype_StepBackward Spintype = 1
	Spintype_PageForward  Spintype = 2
	Spintype_PageBackward Spintype = 3
	Spintype_Home         Spintype = 4
	Spintype_End          Spintype = 5
	Spintype_UserDefined  Spintype = 6
)

// Stacktransitiontype is a representation of the C type StackTransitionType.
type Stacktransitiontype int

const (
	Stacktransitiontype_None           Stacktransitiontype = 0
	Stacktransitiontype_Crossfade      Stacktransitiontype = 1
	Stacktransitiontype_SlideRight     Stacktransitiontype = 2
	Stacktransitiontype_SlideLeft      Stacktransitiontype = 3
	Stacktransitiontype_SlideUp        Stacktransitiontype = 4
	Stacktransitiontype_SlideDown      Stacktransitiontype = 5
	Stacktransitiontype_SlideLeftRight Stacktransitiontype = 6
	Stacktransitiontype_SlideUpDown    Stacktransitiontype = 7
	Stacktransitiontype_OverUp         Stacktransitiontype = 8
	Stacktransitiontype_OverDown       Stacktransitiontype = 9
	Stacktransitiontype_OverLeft       Stacktransitiontype = 10
	Stacktransitiontype_OverRight      Stacktransitiontype = 11
	Stacktransitiontype_UnderUp        Stacktransitiontype = 12
	Stacktransitiontype_UnderDown      Stacktransitiontype = 13
	Stacktransitiontype_UnderLeft      Stacktransitiontype = 14
	Stacktransitiontype_UnderRight     Stacktransitiontype = 15
	Stacktransitiontype_OverUpDown     Stacktransitiontype = 16
	Stacktransitiontype_OverDownUp     Stacktransitiontype = 17
	Stacktransitiontype_OverLeftRight  Stacktransitiontype = 18
	Stacktransitiontype_OverRightLeft  Stacktransitiontype = 19
)

// Statetype is a representation of the C type StateType.
type Statetype int

const (
	Statetype_Normal       Statetype = 0
	Statetype_Active       Statetype = 1
	Statetype_Prelight     Statetype = 2
	Statetype_Selected     Statetype = 3
	Statetype_Insensitive  Statetype = 4
	Statetype_Inconsistent Statetype = 5
	Statetype_Focused      Statetype = 6
)

// Textbuffertargetinfo is a representation of the C type TextBufferTargetInfo.
type Textbuffertargetinfo int

const (
	Textbuffertargetinfo_BufferContents Textbuffertargetinfo = -1
	Textbuffertargetinfo_RichText       Textbuffertargetinfo = -2
	Textbuffertargetinfo_Text           Textbuffertargetinfo = -3
)

// Textdirection is a representation of the C type TextDirection.
type Textdirection int

const (
	Textdirection_None Textdirection = 0
	Textdirection_Ltr  Textdirection = 1
	Textdirection_Rtl  Textdirection = 2
)

// Textextendselection is a representation of the C type TextExtendSelection.
//
// since 3.16
type Textextendselection int

const (
	Textextendselection_Word Textextendselection = 0
	Textextendselection_Line Textextendselection = 1
)

// Textviewlayer is a representation of the C type TextViewLayer.
type Textviewlayer int

const (
	Textviewlayer_Below     Textviewlayer = 0
	Textviewlayer_Above     Textviewlayer = 1
	Textviewlayer_BelowText Textviewlayer = 2
	Textviewlayer_AboveText Textviewlayer = 3
)

// Textwindowtype is a representation of the C type TextWindowType.
type Textwindowtype int

const (
	Textwindowtype_Private Textwindowtype = 0
	Textwindowtype_Widget  Textwindowtype = 1
	Textwindowtype_Text    Textwindowtype = 2
	Textwindowtype_Left    Textwindowtype = 3
	Textwindowtype_Right   Textwindowtype = 4
	Textwindowtype_Top     Textwindowtype = 5
	Textwindowtype_Bottom  Textwindowtype = 6
)

// Toolbarspacestyle is a representation of the C type ToolbarSpaceStyle.
type Toolbarspacestyle int

const (
	Toolbarspacestyle_Empty Toolbarspacestyle = 0
	Toolbarspacestyle_Line  Toolbarspacestyle = 1
)

// Toolbarstyle is a representation of the C type ToolbarStyle.
type Toolbarstyle int

const (
	Toolbarstyle_Icons     Toolbarstyle = 0
	Toolbarstyle_Text      Toolbarstyle = 1
	Toolbarstyle_Both      Toolbarstyle = 2
	Toolbarstyle_BothHoriz Toolbarstyle = 3
)

// Treeviewcolumnsizing is a representation of the C type TreeViewColumnSizing.
type Treeviewcolumnsizing int

const (
	Treeviewcolumnsizing_GrowOnly Treeviewcolumnsizing = 0
	Treeviewcolumnsizing_Autosize Treeviewcolumnsizing = 1
	Treeviewcolumnsizing_Fixed    Treeviewcolumnsizing = 2
)

// Treeviewdropposition is a representation of the C type TreeViewDropPosition.
type Treeviewdropposition int

const (
	Treeviewdropposition_Before       Treeviewdropposition = 0
	Treeviewdropposition_After        Treeviewdropposition = 1
	Treeviewdropposition_IntoOrBefore Treeviewdropposition = 2
	Treeviewdropposition_IntoOrAfter  Treeviewdropposition = 3
)

// Treeviewgridlines is a representation of the C type TreeViewGridLines.
type Treeviewgridlines int

const (
	Treeviewgridlines_None       Treeviewgridlines = 0
	Treeviewgridlines_Horizontal Treeviewgridlines = 1
	Treeviewgridlines_Vertical   Treeviewgridlines = 2
	Treeviewgridlines_Both       Treeviewgridlines = 3
)

// Unit is a representation of the C type Unit.
type Unit int

const (
	Unit_None   Unit = 0
	Unit_Points Unit = 1
	Unit_Inch   Unit = 2
	Unit_Mm     Unit = 3
)

// Widgethelptype is a representation of the C type WidgetHelpType.
type Widgethelptype int

const (
	Widgethelptype_Tooltip   Widgethelptype = 0
	Widgethelptype_WhatsThis Widgethelptype = 1
)

// Windowposition is a representation of the C type WindowPosition.
type Windowposition int

const (
	Windowposition_None           Windowposition = 0
	Windowposition_Center         Windowposition = 1
	Windowposition_Mouse          Windowposition = 2
	Windowposition_CenterAlways   Windowposition = 3
	Windowposition_CenterOnParent Windowposition = 4
)

// Windowtype is a representation of the C type WindowType.
type Windowtype int

const (
	Windowtype_Toplevel Windowtype = 0
	Windowtype_Popup    Windowtype = 1
)

// Wrapmode is a representation of the C type WrapMode.
type Wrapmode int

const (
	Wrapmode_None     Wrapmode = 0
	Wrapmode_Char     Wrapmode = 1
	Wrapmode_Word     Wrapmode = 2
	Wrapmode_WordChar Wrapmode = 3
)
