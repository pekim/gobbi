// Code generated - DO NOT EDIT.

package gtk

type Align uint32

const (
	Align_Fill     Align = int64(0)
	Align_Start    Align = int64(1)
	Align_End      Align = int64(2)
	Align_Center   Align = int64(3)
	Align_Baseline Align = int64(4)
)

type ArrowPlacement uint32

const (
	ArrowPlacement_Both  ArrowPlacement = int64(0)
	ArrowPlacement_Start ArrowPlacement = int64(1)
	ArrowPlacement_End   ArrowPlacement = int64(2)
)

type ArrowType uint32

const (
	ArrowType_Up    ArrowType = int64(0)
	ArrowType_Down  ArrowType = int64(1)
	ArrowType_Left  ArrowType = int64(2)
	ArrowType_Right ArrowType = int64(3)
	ArrowType_None  ArrowType = int64(4)
)

type AssistantPageType uint32

const (
	AssistantPageType_Content  AssistantPageType = int64(0)
	AssistantPageType_Intro    AssistantPageType = int64(1)
	AssistantPageType_Confirm  AssistantPageType = int64(2)
	AssistantPageType_Summary  AssistantPageType = int64(3)
	AssistantPageType_Progress AssistantPageType = int64(4)
	AssistantPageType_Custom   AssistantPageType = int64(5)
)

type BaselinePosition uint32

const (
	BaselinePosition_Top    BaselinePosition = int64(0)
	BaselinePosition_Center BaselinePosition = int64(1)
	BaselinePosition_Bottom BaselinePosition = int64(2)
)

type BorderStyle uint32

const (
	BorderStyle_None   BorderStyle = int64(0)
	BorderStyle_Solid  BorderStyle = int64(1)
	BorderStyle_Inset  BorderStyle = int64(2)
	BorderStyle_Outset BorderStyle = int64(3)
	BorderStyle_Hidden BorderStyle = int64(4)
	BorderStyle_Dotted BorderStyle = int64(5)
	BorderStyle_Dashed BorderStyle = int64(6)
	BorderStyle_Double BorderStyle = int64(7)
	BorderStyle_Groove BorderStyle = int64(8)
	BorderStyle_Ridge  BorderStyle = int64(9)
)

type BuilderError uint32

const (
	BuilderError_InvalidTypeFunction  BuilderError = int64(0)
	BuilderError_UnhandledTag         BuilderError = int64(1)
	BuilderError_MissingAttribute     BuilderError = int64(2)
	BuilderError_InvalidAttribute     BuilderError = int64(3)
	BuilderError_InvalidTag           BuilderError = int64(4)
	BuilderError_MissingPropertyValue BuilderError = int64(5)
	BuilderError_InvalidValue         BuilderError = int64(6)
	BuilderError_VersionMismatch      BuilderError = int64(7)
	BuilderError_DuplicateId          BuilderError = int64(8)
	BuilderError_ObjectTypeRefused    BuilderError = int64(9)
	BuilderError_TemplateMismatch     BuilderError = int64(10)
	BuilderError_InvalidProperty      BuilderError = int64(11)
	BuilderError_InvalidSignal        BuilderError = int64(12)
	BuilderError_InvalidId            BuilderError = int64(13)
)

type ButtonBoxStyle uint32

const (
	ButtonBoxStyle_Spread ButtonBoxStyle = int64(1)
	ButtonBoxStyle_Edge   ButtonBoxStyle = int64(2)
	ButtonBoxStyle_Start  ButtonBoxStyle = int64(3)
	ButtonBoxStyle_End    ButtonBoxStyle = int64(4)
	ButtonBoxStyle_Center ButtonBoxStyle = int64(5)
	ButtonBoxStyle_Expand ButtonBoxStyle = int64(6)
)

type ButtonRole uint32

const (
	ButtonRole_Normal ButtonRole = int64(0)
	ButtonRole_Check  ButtonRole = int64(1)
	ButtonRole_Radio  ButtonRole = int64(2)
)

type ButtonsType uint32

const (
	ButtonsType_None     ButtonsType = int64(0)
	ButtonsType_Ok       ButtonsType = int64(1)
	ButtonsType_Close    ButtonsType = int64(2)
	ButtonsType_Cancel   ButtonsType = int64(3)
	ButtonsType_YesNo    ButtonsType = int64(4)
	ButtonsType_OkCancel ButtonsType = int64(5)
)

type CellRendererAccelMode uint32

const (
	CellRendererAccelMode_Gtk         CellRendererAccelMode = int64(0)
	CellRendererAccelMode_Other       CellRendererAccelMode = int64(1)
	CellRendererAccelMode_ModifierTap CellRendererAccelMode = int64(2)
)

type CellRendererMode uint32

const (
	CellRendererMode_Inert       CellRendererMode = int64(0)
	CellRendererMode_Activatable CellRendererMode = int64(1)
	CellRendererMode_Editable    CellRendererMode = int64(2)
)

type CornerType uint32

const (
	CornerType_TopLeft     CornerType = int64(0)
	CornerType_BottomLeft  CornerType = int64(1)
	CornerType_TopRight    CornerType = int64(2)
	CornerType_BottomRight CornerType = int64(3)
)

type CssProviderError uint32

const (
	CssProviderError_Failed       CssProviderError = int64(0)
	CssProviderError_Syntax       CssProviderError = int64(1)
	CssProviderError_Import       CssProviderError = int64(2)
	CssProviderError_Name         CssProviderError = int64(3)
	CssProviderError_Deprecated   CssProviderError = int64(4)
	CssProviderError_UnknownValue CssProviderError = int64(5)
)

type CssSectionType uint32

const (
	CssSectionType_Document        CssSectionType = int64(0)
	CssSectionType_Import          CssSectionType = int64(1)
	CssSectionType_ColorDefinition CssSectionType = int64(2)
	CssSectionType_BindingSet      CssSectionType = int64(3)
	CssSectionType_Ruleset         CssSectionType = int64(4)
	CssSectionType_Selector        CssSectionType = int64(5)
	CssSectionType_Declaration     CssSectionType = int64(6)
	CssSectionType_Value           CssSectionType = int64(7)
	CssSectionType_Keyframes       CssSectionType = int64(8)
)

type DeleteType uint32

const (
	DeleteType_Chars           DeleteType = int64(0)
	DeleteType_WordEnds        DeleteType = int64(1)
	DeleteType_Words           DeleteType = int64(2)
	DeleteType_DisplayLines    DeleteType = int64(3)
	DeleteType_DisplayLineEnds DeleteType = int64(4)
	DeleteType_ParagraphEnds   DeleteType = int64(5)
	DeleteType_Paragraphs      DeleteType = int64(6)
	DeleteType_Whitespace      DeleteType = int64(7)
)

type DirectionType uint32

const (
	DirectionType_TabForward  DirectionType = int64(0)
	DirectionType_TabBackward DirectionType = int64(1)
	DirectionType_Up          DirectionType = int64(2)
	DirectionType_Down        DirectionType = int64(3)
	DirectionType_Left        DirectionType = int64(4)
	DirectionType_Right       DirectionType = int64(5)
)

type DragResult uint32

const (
	DragResult_Success        DragResult = int64(0)
	DragResult_NoTarget       DragResult = int64(1)
	DragResult_UserCancelled  DragResult = int64(2)
	DragResult_TimeoutExpired DragResult = int64(3)
	DragResult_GrabBroken     DragResult = int64(4)
	DragResult_Error          DragResult = int64(5)
)

type EntryIconPosition uint32

const (
	EntryIconPosition_Primary   EntryIconPosition = int64(0)
	EntryIconPosition_Secondary EntryIconPosition = int64(1)
)

type EventSequenceState uint32

const (
	EventSequenceState_None    EventSequenceState = int64(0)
	EventSequenceState_Claimed EventSequenceState = int64(1)
	EventSequenceState_Denied  EventSequenceState = int64(2)
)

type ExpanderStyle uint32

const (
	ExpanderStyle_Collapsed     ExpanderStyle = int64(0)
	ExpanderStyle_SemiCollapsed ExpanderStyle = int64(1)
	ExpanderStyle_SemiExpanded  ExpanderStyle = int64(2)
	ExpanderStyle_Expanded      ExpanderStyle = int64(3)
)

type FileChooserAction uint32

const (
	FileChooserAction_Open         FileChooserAction = int64(0)
	FileChooserAction_Save         FileChooserAction = int64(1)
	FileChooserAction_SelectFolder FileChooserAction = int64(2)
	FileChooserAction_CreateFolder FileChooserAction = int64(3)
)

type FileChooserConfirmation uint32

const (
	FileChooserConfirmation_Confirm        FileChooserConfirmation = int64(0)
	FileChooserConfirmation_AcceptFilename FileChooserConfirmation = int64(1)
	FileChooserConfirmation_SelectAgain    FileChooserConfirmation = int64(2)
)

type FileChooserError uint32

const (
	FileChooserError_Nonexistent        FileChooserError = int64(0)
	FileChooserError_BadFilename        FileChooserError = int64(1)
	FileChooserError_AlreadyExists      FileChooserError = int64(2)
	FileChooserError_IncompleteHostname FileChooserError = int64(3)
)

type IMPreeditStyle uint32

const (
	IMPreeditStyle_Nothing  IMPreeditStyle = int64(0)
	IMPreeditStyle_Callback IMPreeditStyle = int64(1)
	IMPreeditStyle_None     IMPreeditStyle = int64(2)
)

type IMStatusStyle uint32

const (
	IMStatusStyle_Nothing  IMStatusStyle = int64(0)
	IMStatusStyle_Callback IMStatusStyle = int64(1)
	IMStatusStyle_None     IMStatusStyle = int64(2)
)

type IconSize uint32

const (
	IconSize_Invalid      IconSize = int64(0)
	IconSize_Menu         IconSize = int64(1)
	IconSize_SmallToolbar IconSize = int64(2)
	IconSize_LargeToolbar IconSize = int64(3)
	IconSize_Button       IconSize = int64(4)
	IconSize_Dnd          IconSize = int64(5)
	IconSize_Dialog       IconSize = int64(6)
)

type IconThemeError uint32

const (
	IconThemeError_NotFound IconThemeError = int64(0)
	IconThemeError_Failed   IconThemeError = int64(1)
)

type IconViewDropPosition uint32

const (
	IconViewDropPosition_NoDrop    IconViewDropPosition = int64(0)
	IconViewDropPosition_DropInto  IconViewDropPosition = int64(1)
	IconViewDropPosition_DropLeft  IconViewDropPosition = int64(2)
	IconViewDropPosition_DropRight IconViewDropPosition = int64(3)
	IconViewDropPosition_DropAbove IconViewDropPosition = int64(4)
	IconViewDropPosition_DropBelow IconViewDropPosition = int64(5)
)

type ImageType uint32

const (
	ImageType_Empty     ImageType = int64(0)
	ImageType_Pixbuf    ImageType = int64(1)
	ImageType_Stock     ImageType = int64(2)
	ImageType_IconSet   ImageType = int64(3)
	ImageType_Animation ImageType = int64(4)
	ImageType_IconName  ImageType = int64(5)
	ImageType_Gicon     ImageType = int64(6)
	ImageType_Surface   ImageType = int64(7)
)

type InputPurpose uint32

const (
	InputPurpose_FreeForm InputPurpose = int64(0)
	InputPurpose_Alpha    InputPurpose = int64(1)
	InputPurpose_Digits   InputPurpose = int64(2)
	InputPurpose_Number   InputPurpose = int64(3)
	InputPurpose_Phone    InputPurpose = int64(4)
	InputPurpose_Url      InputPurpose = int64(5)
	InputPurpose_Email    InputPurpose = int64(6)
	InputPurpose_Name     InputPurpose = int64(7)
	InputPurpose_Password InputPurpose = int64(8)
	InputPurpose_Pin      InputPurpose = int64(9)
)

type Justification uint32

const (
	Justification_Left   Justification = int64(0)
	Justification_Right  Justification = int64(1)
	Justification_Center Justification = int64(2)
	Justification_Fill   Justification = int64(3)
)

type LevelBarMode uint32

const (
	LevelBarMode_Continuous LevelBarMode = int64(0)
	LevelBarMode_Discrete   LevelBarMode = int64(1)
)

type License uint32

const (
	License_Unknown    License = int64(0)
	License_Custom     License = int64(1)
	License_Gpl20      License = int64(2)
	License_Gpl30      License = int64(3)
	License_Lgpl21     License = int64(4)
	License_Lgpl30     License = int64(5)
	License_Bsd        License = int64(6)
	License_MitX11     License = int64(7)
	License_Artistic   License = int64(8)
	License_Gpl20Only  License = int64(9)
	License_Gpl30Only  License = int64(10)
	License_Lgpl21Only License = int64(11)
	License_Lgpl30Only License = int64(12)
	License_Agpl30     License = int64(13)
	License_Agpl30Only License = int64(14)
)

type MenuDirectionType uint32

const (
	MenuDirectionType_Parent MenuDirectionType = int64(0)
	MenuDirectionType_Child  MenuDirectionType = int64(1)
	MenuDirectionType_Next   MenuDirectionType = int64(2)
	MenuDirectionType_Prev   MenuDirectionType = int64(3)
)

type MessageType uint32

const (
	MessageType_Info     MessageType = int64(0)
	MessageType_Warning  MessageType = int64(1)
	MessageType_Question MessageType = int64(2)
	MessageType_Error    MessageType = int64(3)
	MessageType_Other    MessageType = int64(4)
)

type MovementStep uint32

const (
	MovementStep_LogicalPositions MovementStep = int64(0)
	MovementStep_VisualPositions  MovementStep = int64(1)
	MovementStep_Words            MovementStep = int64(2)
	MovementStep_DisplayLines     MovementStep = int64(3)
	MovementStep_DisplayLineEnds  MovementStep = int64(4)
	MovementStep_Paragraphs       MovementStep = int64(5)
	MovementStep_ParagraphEnds    MovementStep = int64(6)
	MovementStep_Pages            MovementStep = int64(7)
	MovementStep_BufferEnds       MovementStep = int64(8)
	MovementStep_HorizontalPages  MovementStep = int64(9)
)

type NotebookTab uint32

const (
	NotebookTab_First NotebookTab = int64(0)
	NotebookTab_Last  NotebookTab = int64(1)
)

type NumberUpLayout uint32

const (
	NumberUpLayout_Lrtb NumberUpLayout = int64(0)
	NumberUpLayout_Lrbt NumberUpLayout = int64(1)
	NumberUpLayout_Rltb NumberUpLayout = int64(2)
	NumberUpLayout_Rlbt NumberUpLayout = int64(3)
	NumberUpLayout_Tblr NumberUpLayout = int64(4)
	NumberUpLayout_Tbrl NumberUpLayout = int64(5)
	NumberUpLayout_Btlr NumberUpLayout = int64(6)
	NumberUpLayout_Btrl NumberUpLayout = int64(7)
)

type Orientation uint32

const (
	Orientation_Horizontal Orientation = int64(0)
	Orientation_Vertical   Orientation = int64(1)
)

type PackDirection uint32

const (
	PackDirection_Ltr PackDirection = int64(0)
	PackDirection_Rtl PackDirection = int64(1)
	PackDirection_Ttb PackDirection = int64(2)
	PackDirection_Btt PackDirection = int64(3)
)

type PackType uint32

const (
	PackType_Start PackType = int64(0)
	PackType_End   PackType = int64(1)
)

type PadActionType uint32

const (
	PadActionType_Button PadActionType = int64(0)
	PadActionType_Ring   PadActionType = int64(1)
	PadActionType_Strip  PadActionType = int64(2)
)

type PageOrientation uint32

const (
	PageOrientation_Portrait         PageOrientation = int64(0)
	PageOrientation_Landscape        PageOrientation = int64(1)
	PageOrientation_ReversePortrait  PageOrientation = int64(2)
	PageOrientation_ReverseLandscape PageOrientation = int64(3)
)

type PageSet uint32

const (
	PageSet_All  PageSet = int64(0)
	PageSet_Even PageSet = int64(1)
	PageSet_Odd  PageSet = int64(2)
)

type PanDirection uint32

const (
	PanDirection_Left  PanDirection = int64(0)
	PanDirection_Right PanDirection = int64(1)
	PanDirection_Up    PanDirection = int64(2)
	PanDirection_Down  PanDirection = int64(3)
)

type PathPriorityType uint32

const (
	PathPriorityType_Lowest      PathPriorityType = int64(0)
	PathPriorityType_Gtk         PathPriorityType = int64(4)
	PathPriorityType_Application PathPriorityType = int64(8)
	PathPriorityType_Theme       PathPriorityType = int64(10)
	PathPriorityType_Rc          PathPriorityType = int64(12)
	PathPriorityType_Highest     PathPriorityType = int64(15)
)

type PathType uint32

const (
	PathType_Widget      PathType = int64(0)
	PathType_WidgetClass PathType = int64(1)
	PathType_Class       PathType = int64(2)
)

type PolicyType uint32

const (
	PolicyType_Always    PolicyType = int64(0)
	PolicyType_Automatic PolicyType = int64(1)
	PolicyType_Never     PolicyType = int64(2)
	PolicyType_External  PolicyType = int64(3)
)

type PopoverConstraint uint32

const (
	PopoverConstraint_None   PopoverConstraint = int64(0)
	PopoverConstraint_Window PopoverConstraint = int64(1)
)

type PositionType uint32

const (
	PositionType_Left   PositionType = int64(0)
	PositionType_Right  PositionType = int64(1)
	PositionType_Top    PositionType = int64(2)
	PositionType_Bottom PositionType = int64(3)
)

type PrintDuplex uint32

const (
	PrintDuplex_Simplex    PrintDuplex = int64(0)
	PrintDuplex_Horizontal PrintDuplex = int64(1)
	PrintDuplex_Vertical   PrintDuplex = int64(2)
)

type PrintError uint32

const (
	PrintError_General       PrintError = int64(0)
	PrintError_InternalError PrintError = int64(1)
	PrintError_Nomem         PrintError = int64(2)
	PrintError_InvalidFile   PrintError = int64(3)
)

type PrintOperationAction uint32

const (
	PrintOperationAction_PrintDialog PrintOperationAction = int64(0)
	PrintOperationAction_Print       PrintOperationAction = int64(1)
	PrintOperationAction_Preview     PrintOperationAction = int64(2)
	PrintOperationAction_Export      PrintOperationAction = int64(3)
)

type PrintOperationResult uint32

const (
	PrintOperationResult_Error      PrintOperationResult = int64(0)
	PrintOperationResult_Apply      PrintOperationResult = int64(1)
	PrintOperationResult_Cancel     PrintOperationResult = int64(2)
	PrintOperationResult_InProgress PrintOperationResult = int64(3)
)

type PrintPages uint32

const (
	PrintPages_All       PrintPages = int64(0)
	PrintPages_Current   PrintPages = int64(1)
	PrintPages_Ranges    PrintPages = int64(2)
	PrintPages_Selection PrintPages = int64(3)
)

type PrintQuality uint32

const (
	PrintQuality_Low    PrintQuality = int64(0)
	PrintQuality_Normal PrintQuality = int64(1)
	PrintQuality_High   PrintQuality = int64(2)
	PrintQuality_Draft  PrintQuality = int64(3)
)

type PrintStatus uint32

const (
	PrintStatus_Initial         PrintStatus = int64(0)
	PrintStatus_Preparing       PrintStatus = int64(1)
	PrintStatus_GeneratingData  PrintStatus = int64(2)
	PrintStatus_SendingData     PrintStatus = int64(3)
	PrintStatus_Pending         PrintStatus = int64(4)
	PrintStatus_PendingIssue    PrintStatus = int64(5)
	PrintStatus_Printing        PrintStatus = int64(6)
	PrintStatus_Finished        PrintStatus = int64(7)
	PrintStatus_FinishedAborted PrintStatus = int64(8)
)

type PropagationPhase uint32

const (
	PropagationPhase_None    PropagationPhase = int64(0)
	PropagationPhase_Capture PropagationPhase = int64(1)
	PropagationPhase_Bubble  PropagationPhase = int64(2)
	PropagationPhase_Target  PropagationPhase = int64(3)
)

type RcTokenType uint32

const (
	RcTokenType_Invalid      RcTokenType = int64(270)
	RcTokenType_Include      RcTokenType = int64(271)
	RcTokenType_Normal       RcTokenType = int64(272)
	RcTokenType_Active       RcTokenType = int64(273)
	RcTokenType_Prelight     RcTokenType = int64(274)
	RcTokenType_Selected     RcTokenType = int64(275)
	RcTokenType_Insensitive  RcTokenType = int64(276)
	RcTokenType_Fg           RcTokenType = int64(277)
	RcTokenType_Bg           RcTokenType = int64(278)
	RcTokenType_Text         RcTokenType = int64(279)
	RcTokenType_Base         RcTokenType = int64(280)
	RcTokenType_Xthickness   RcTokenType = int64(281)
	RcTokenType_Ythickness   RcTokenType = int64(282)
	RcTokenType_Font         RcTokenType = int64(283)
	RcTokenType_Fontset      RcTokenType = int64(284)
	RcTokenType_FontName     RcTokenType = int64(285)
	RcTokenType_BgPixmap     RcTokenType = int64(286)
	RcTokenType_PixmapPath   RcTokenType = int64(287)
	RcTokenType_Style        RcTokenType = int64(288)
	RcTokenType_Binding      RcTokenType = int64(289)
	RcTokenType_Bind         RcTokenType = int64(290)
	RcTokenType_Widget       RcTokenType = int64(291)
	RcTokenType_WidgetClass  RcTokenType = int64(292)
	RcTokenType_Class        RcTokenType = int64(293)
	RcTokenType_Lowest       RcTokenType = int64(294)
	RcTokenType_Gtk          RcTokenType = int64(295)
	RcTokenType_Application  RcTokenType = int64(296)
	RcTokenType_Theme        RcTokenType = int64(297)
	RcTokenType_Rc           RcTokenType = int64(298)
	RcTokenType_Highest      RcTokenType = int64(299)
	RcTokenType_Engine       RcTokenType = int64(300)
	RcTokenType_ModulePath   RcTokenType = int64(301)
	RcTokenType_ImModulePath RcTokenType = int64(302)
	RcTokenType_ImModuleFile RcTokenType = int64(303)
	RcTokenType_Stock        RcTokenType = int64(304)
	RcTokenType_Ltr          RcTokenType = int64(305)
	RcTokenType_Rtl          RcTokenType = int64(306)
	RcTokenType_Color        RcTokenType = int64(307)
	RcTokenType_Unbind       RcTokenType = int64(308)
	RcTokenType_Last         RcTokenType = int64(309)
)

type RecentChooserError uint32

const (
	RecentChooserError_NotFound   RecentChooserError = int64(0)
	RecentChooserError_InvalidUri RecentChooserError = int64(1)
)

type RecentManagerError uint32

const (
	RecentManagerError_NotFound        RecentManagerError = int64(0)
	RecentManagerError_InvalidUri      RecentManagerError = int64(1)
	RecentManagerError_InvalidEncoding RecentManagerError = int64(2)
	RecentManagerError_NotRegistered   RecentManagerError = int64(3)
	RecentManagerError_Read            RecentManagerError = int64(4)
	RecentManagerError_Write           RecentManagerError = int64(5)
	RecentManagerError_Unknown         RecentManagerError = int64(6)
)

type RecentSortType uint32

const (
	RecentSortType_None   RecentSortType = int64(0)
	RecentSortType_Mru    RecentSortType = int64(1)
	RecentSortType_Lru    RecentSortType = int64(2)
	RecentSortType_Custom RecentSortType = int64(3)
)

type ReliefStyle uint32

const (
	ReliefStyle_Normal ReliefStyle = int64(0)
	ReliefStyle_Half   ReliefStyle = int64(1)
	ReliefStyle_None   ReliefStyle = int64(2)
)

type ResizeMode uint32

const (
	ResizeMode_Parent    ResizeMode = int64(0)
	ResizeMode_Queue     ResizeMode = int64(1)
	ResizeMode_Immediate ResizeMode = int64(2)
)

type ResponseType int32

const (
	ResponseType_None        ResponseType = int64(-1)
	ResponseType_Reject      ResponseType = int64(-2)
	ResponseType_Accept      ResponseType = int64(-3)
	ResponseType_DeleteEvent ResponseType = int64(-4)
	ResponseType_Ok          ResponseType = int64(-5)
	ResponseType_Cancel      ResponseType = int64(-6)
	ResponseType_Close       ResponseType = int64(-7)
	ResponseType_Yes         ResponseType = int64(-8)
	ResponseType_No          ResponseType = int64(-9)
	ResponseType_Apply       ResponseType = int64(-10)
	ResponseType_Help        ResponseType = int64(-11)
)

type RevealerTransitionType uint32

const (
	RevealerTransitionType_None       RevealerTransitionType = int64(0)
	RevealerTransitionType_Crossfade  RevealerTransitionType = int64(1)
	RevealerTransitionType_SlideRight RevealerTransitionType = int64(2)
	RevealerTransitionType_SlideLeft  RevealerTransitionType = int64(3)
	RevealerTransitionType_SlideUp    RevealerTransitionType = int64(4)
	RevealerTransitionType_SlideDown  RevealerTransitionType = int64(5)
)

type ScrollStep uint32

const (
	ScrollStep_Steps           ScrollStep = int64(0)
	ScrollStep_Pages           ScrollStep = int64(1)
	ScrollStep_Ends            ScrollStep = int64(2)
	ScrollStep_HorizontalSteps ScrollStep = int64(3)
	ScrollStep_HorizontalPages ScrollStep = int64(4)
	ScrollStep_HorizontalEnds  ScrollStep = int64(5)
)

type ScrollType uint32

const (
	ScrollType_None         ScrollType = int64(0)
	ScrollType_Jump         ScrollType = int64(1)
	ScrollType_StepBackward ScrollType = int64(2)
	ScrollType_StepForward  ScrollType = int64(3)
	ScrollType_PageBackward ScrollType = int64(4)
	ScrollType_PageForward  ScrollType = int64(5)
	ScrollType_StepUp       ScrollType = int64(6)
	ScrollType_StepDown     ScrollType = int64(7)
	ScrollType_PageUp       ScrollType = int64(8)
	ScrollType_PageDown     ScrollType = int64(9)
	ScrollType_StepLeft     ScrollType = int64(10)
	ScrollType_StepRight    ScrollType = int64(11)
	ScrollType_PageLeft     ScrollType = int64(12)
	ScrollType_PageRight    ScrollType = int64(13)
	ScrollType_Start        ScrollType = int64(14)
	ScrollType_End          ScrollType = int64(15)
)

type ScrollablePolicy uint32

const (
	ScrollablePolicy_Minimum ScrollablePolicy = int64(0)
	ScrollablePolicy_Natural ScrollablePolicy = int64(1)
)

type SelectionMode uint32

const (
	SelectionMode_None     SelectionMode = int64(0)
	SelectionMode_Single   SelectionMode = int64(1)
	SelectionMode_Browse   SelectionMode = int64(2)
	SelectionMode_Multiple SelectionMode = int64(3)
)

type SensitivityType uint32

const (
	SensitivityType_Auto SensitivityType = int64(0)
	SensitivityType_On   SensitivityType = int64(1)
	SensitivityType_Off  SensitivityType = int64(2)
)

type ShadowType uint32

const (
	ShadowType_None      ShadowType = int64(0)
	ShadowType_In        ShadowType = int64(1)
	ShadowType_Out       ShadowType = int64(2)
	ShadowType_EtchedIn  ShadowType = int64(3)
	ShadowType_EtchedOut ShadowType = int64(4)
)

type ShortcutType uint32

const (
	ShortcutType_Accelerator                   ShortcutType = int64(0)
	ShortcutType_GesturePinch                  ShortcutType = int64(1)
	ShortcutType_GestureStretch                ShortcutType = int64(2)
	ShortcutType_GestureRotateClockwise        ShortcutType = int64(3)
	ShortcutType_GestureRotateCounterclockwise ShortcutType = int64(4)
	ShortcutType_GestureTwoFingerSwipeLeft     ShortcutType = int64(5)
	ShortcutType_GestureTwoFingerSwipeRight    ShortcutType = int64(6)
	ShortcutType_Gesture                       ShortcutType = int64(7)
)

type SizeGroupMode uint32

const (
	SizeGroupMode_None       SizeGroupMode = int64(0)
	SizeGroupMode_Horizontal SizeGroupMode = int64(1)
	SizeGroupMode_Vertical   SizeGroupMode = int64(2)
	SizeGroupMode_Both       SizeGroupMode = int64(3)
)

type SizeRequestMode uint32

const (
	SizeRequestMode_HeightForWidth SizeRequestMode = int64(0)
	SizeRequestMode_WidthForHeight SizeRequestMode = int64(1)
	SizeRequestMode_ConstantSize   SizeRequestMode = int64(2)
)

type SortType uint32

const (
	SortType_Ascending  SortType = int64(0)
	SortType_Descending SortType = int64(1)
)

type SpinButtonUpdatePolicy uint32

const (
	SpinButtonUpdatePolicy_Always  SpinButtonUpdatePolicy = int64(0)
	SpinButtonUpdatePolicy_IfValid SpinButtonUpdatePolicy = int64(1)
)

type SpinType uint32

const (
	SpinType_StepForward  SpinType = int64(0)
	SpinType_StepBackward SpinType = int64(1)
	SpinType_PageForward  SpinType = int64(2)
	SpinType_PageBackward SpinType = int64(3)
	SpinType_Home         SpinType = int64(4)
	SpinType_End          SpinType = int64(5)
	SpinType_UserDefined  SpinType = int64(6)
)

type StackTransitionType uint32

const (
	StackTransitionType_None           StackTransitionType = int64(0)
	StackTransitionType_Crossfade      StackTransitionType = int64(1)
	StackTransitionType_SlideRight     StackTransitionType = int64(2)
	StackTransitionType_SlideLeft      StackTransitionType = int64(3)
	StackTransitionType_SlideUp        StackTransitionType = int64(4)
	StackTransitionType_SlideDown      StackTransitionType = int64(5)
	StackTransitionType_SlideLeftRight StackTransitionType = int64(6)
	StackTransitionType_SlideUpDown    StackTransitionType = int64(7)
	StackTransitionType_OverUp         StackTransitionType = int64(8)
	StackTransitionType_OverDown       StackTransitionType = int64(9)
	StackTransitionType_OverLeft       StackTransitionType = int64(10)
	StackTransitionType_OverRight      StackTransitionType = int64(11)
	StackTransitionType_UnderUp        StackTransitionType = int64(12)
	StackTransitionType_UnderDown      StackTransitionType = int64(13)
	StackTransitionType_UnderLeft      StackTransitionType = int64(14)
	StackTransitionType_UnderRight     StackTransitionType = int64(15)
	StackTransitionType_OverUpDown     StackTransitionType = int64(16)
	StackTransitionType_OverDownUp     StackTransitionType = int64(17)
	StackTransitionType_OverLeftRight  StackTransitionType = int64(18)
	StackTransitionType_OverRightLeft  StackTransitionType = int64(19)
)

type StateType uint32

const (
	StateType_Normal       StateType = int64(0)
	StateType_Active       StateType = int64(1)
	StateType_Prelight     StateType = int64(2)
	StateType_Selected     StateType = int64(3)
	StateType_Insensitive  StateType = int64(4)
	StateType_Inconsistent StateType = int64(5)
	StateType_Focused      StateType = int64(6)
)

type TextBufferTargetInfo int32

const (
	TextBufferTargetInfo_BufferContents TextBufferTargetInfo = int64(-1)
	TextBufferTargetInfo_RichText       TextBufferTargetInfo = int64(-2)
	TextBufferTargetInfo_Text           TextBufferTargetInfo = int64(-3)
)

type TextDirection uint32

const (
	TextDirection_None TextDirection = int64(0)
	TextDirection_Ltr  TextDirection = int64(1)
	TextDirection_Rtl  TextDirection = int64(2)
)

type TextExtendSelection uint32

const (
	TextExtendSelection_Word TextExtendSelection = int64(0)
	TextExtendSelection_Line TextExtendSelection = int64(1)
)

type TextViewLayer uint32

const (
	TextViewLayer_Below     TextViewLayer = int64(0)
	TextViewLayer_Above     TextViewLayer = int64(1)
	TextViewLayer_BelowText TextViewLayer = int64(2)
	TextViewLayer_AboveText TextViewLayer = int64(3)
)

type TextWindowType uint32

const (
	TextWindowType_Private TextWindowType = int64(0)
	TextWindowType_Widget  TextWindowType = int64(1)
	TextWindowType_Text    TextWindowType = int64(2)
	TextWindowType_Left    TextWindowType = int64(3)
	TextWindowType_Right   TextWindowType = int64(4)
	TextWindowType_Top     TextWindowType = int64(5)
	TextWindowType_Bottom  TextWindowType = int64(6)
)

type ToolbarSpaceStyle uint32

const (
	ToolbarSpaceStyle_Empty ToolbarSpaceStyle = int64(0)
	ToolbarSpaceStyle_Line  ToolbarSpaceStyle = int64(1)
)

type ToolbarStyle uint32

const (
	ToolbarStyle_Icons     ToolbarStyle = int64(0)
	ToolbarStyle_Text      ToolbarStyle = int64(1)
	ToolbarStyle_Both      ToolbarStyle = int64(2)
	ToolbarStyle_BothHoriz ToolbarStyle = int64(3)
)

type TreeViewColumnSizing uint32

const (
	TreeViewColumnSizing_GrowOnly TreeViewColumnSizing = int64(0)
	TreeViewColumnSizing_Autosize TreeViewColumnSizing = int64(1)
	TreeViewColumnSizing_Fixed    TreeViewColumnSizing = int64(2)
)

type TreeViewDropPosition uint32

const (
	TreeViewDropPosition_Before       TreeViewDropPosition = int64(0)
	TreeViewDropPosition_After        TreeViewDropPosition = int64(1)
	TreeViewDropPosition_IntoOrBefore TreeViewDropPosition = int64(2)
	TreeViewDropPosition_IntoOrAfter  TreeViewDropPosition = int64(3)
)

type TreeViewGridLines uint32

const (
	TreeViewGridLines_None       TreeViewGridLines = int64(0)
	TreeViewGridLines_Horizontal TreeViewGridLines = int64(1)
	TreeViewGridLines_Vertical   TreeViewGridLines = int64(2)
	TreeViewGridLines_Both       TreeViewGridLines = int64(3)
)

type Unit uint32

const (
	Unit_None   Unit = int64(0)
	Unit_Points Unit = int64(1)
	Unit_Inch   Unit = int64(2)
	Unit_Mm     Unit = int64(3)
)

type WidgetHelpType uint32

const (
	WidgetHelpType_Tooltip   WidgetHelpType = int64(0)
	WidgetHelpType_WhatsThis WidgetHelpType = int64(1)
)

type WindowPosition uint32

const (
	WindowPosition_None           WindowPosition = int64(0)
	WindowPosition_Center         WindowPosition = int64(1)
	WindowPosition_Mouse          WindowPosition = int64(2)
	WindowPosition_CenterAlways   WindowPosition = int64(3)
	WindowPosition_CenterOnParent WindowPosition = int64(4)
)

type WindowType uint32

const (
	WindowType_Toplevel WindowType = int64(0)
	WindowType_Popup    WindowType = int64(1)
)

type WrapMode uint32

const (
	WrapMode_None     WrapMode = int64(0)
	WrapMode_Char     WrapMode = int64(1)
	WrapMode_Word     WrapMode = int64(2)
	WrapMode_WordChar WrapMode = int64(3)
)
