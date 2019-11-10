// Code generated - DO NOT EDIT.

package gtk

type AccelFlags uint32

const (
	AccelFlags_Visible AccelFlags = int64(1)
	AccelFlags_Locked  AccelFlags = int64(2)
	AccelFlags_Mask    AccelFlags = int64(7)
)

type ApplicationInhibitFlags uint32

const (
	ApplicationInhibitFlags_Logout  ApplicationInhibitFlags = int64(1)
	ApplicationInhibitFlags_Switch  ApplicationInhibitFlags = int64(2)
	ApplicationInhibitFlags_Suspend ApplicationInhibitFlags = int64(4)
	ApplicationInhibitFlags_Idle    ApplicationInhibitFlags = int64(8)
)

type AttachOptions uint32

const (
	AttachOptions_Expand AttachOptions = int64(1)
	AttachOptions_Shrink AttachOptions = int64(2)
	AttachOptions_Fill   AttachOptions = int64(4)
)

type CalendarDisplayOptions uint32

const (
	CalendarDisplayOptions_ShowHeading     CalendarDisplayOptions = int64(1)
	CalendarDisplayOptions_ShowDayNames    CalendarDisplayOptions = int64(2)
	CalendarDisplayOptions_NoMonthChange   CalendarDisplayOptions = int64(4)
	CalendarDisplayOptions_ShowWeekNumbers CalendarDisplayOptions = int64(8)
	CalendarDisplayOptions_ShowDetails     CalendarDisplayOptions = int64(32)
)

type CellRendererState uint32

const (
	CellRendererState_Selected    CellRendererState = int64(1)
	CellRendererState_Prelit      CellRendererState = int64(2)
	CellRendererState_Insensitive CellRendererState = int64(4)
	CellRendererState_Sorted      CellRendererState = int64(8)
	CellRendererState_Focused     CellRendererState = int64(16)
	CellRendererState_Expandable  CellRendererState = int64(32)
	CellRendererState_Expanded    CellRendererState = int64(64)
)

type DebugFlag uint32

const (
	DebugFlag_Misc         DebugFlag = int64(1)
	DebugFlag_Plugsocket   DebugFlag = int64(2)
	DebugFlag_Text         DebugFlag = int64(4)
	DebugFlag_Tree         DebugFlag = int64(8)
	DebugFlag_Updates      DebugFlag = int64(16)
	DebugFlag_Keybindings  DebugFlag = int64(32)
	DebugFlag_Multihead    DebugFlag = int64(64)
	DebugFlag_Modules      DebugFlag = int64(128)
	DebugFlag_Geometry     DebugFlag = int64(256)
	DebugFlag_Icontheme    DebugFlag = int64(512)
	DebugFlag_Printing     DebugFlag = int64(1024)
	DebugFlag_Builder      DebugFlag = int64(2048)
	DebugFlag_SizeRequest  DebugFlag = int64(4096)
	DebugFlag_NoCssCache   DebugFlag = int64(8192)
	DebugFlag_Baselines    DebugFlag = int64(16384)
	DebugFlag_PixelCache   DebugFlag = int64(32768)
	DebugFlag_NoPixelCache DebugFlag = int64(65536)
	DebugFlag_Interactive  DebugFlag = int64(131072)
	DebugFlag_Touchscreen  DebugFlag = int64(262144)
	DebugFlag_Actions      DebugFlag = int64(524288)
	DebugFlag_Resize       DebugFlag = int64(1048576)
	DebugFlag_Layout       DebugFlag = int64(2097152)
)

type DestDefaults uint32

const (
	DestDefaults_Motion    DestDefaults = int64(1)
	DestDefaults_Highlight DestDefaults = int64(2)
	DestDefaults_Drop      DestDefaults = int64(4)
	DestDefaults_All       DestDefaults = int64(7)
)

type DialogFlags uint32

const (
	DialogFlags_Modal             DialogFlags = int64(1)
	DialogFlags_DestroyWithParent DialogFlags = int64(2)
	DialogFlags_UseHeaderBar      DialogFlags = int64(4)
)

type EventControllerScrollFlags uint32

const (
	EventControllerScrollFlags_None       EventControllerScrollFlags = int64(0)
	EventControllerScrollFlags_Vertical   EventControllerScrollFlags = int64(1)
	EventControllerScrollFlags_Horizontal EventControllerScrollFlags = int64(2)
	EventControllerScrollFlags_Discrete   EventControllerScrollFlags = int64(4)
	EventControllerScrollFlags_Kinetic    EventControllerScrollFlags = int64(8)
	EventControllerScrollFlags_BothAxes   EventControllerScrollFlags = int64(3)
)

type FileFilterFlags uint32

const (
	FileFilterFlags_Filename    FileFilterFlags = int64(1)
	FileFilterFlags_Uri         FileFilterFlags = int64(2)
	FileFilterFlags_DisplayName FileFilterFlags = int64(4)
	FileFilterFlags_MimeType    FileFilterFlags = int64(8)
)

type FontChooserLevel uint32

const (
	FontChooserLevel_Family     FontChooserLevel = int64(0)
	FontChooserLevel_Style      FontChooserLevel = int64(1)
	FontChooserLevel_Size       FontChooserLevel = int64(2)
	FontChooserLevel_Variations FontChooserLevel = int64(4)
	FontChooserLevel_Features   FontChooserLevel = int64(8)
)

type IconLookupFlags uint32

const (
	IconLookupFlags_NoSvg           IconLookupFlags = int64(1)
	IconLookupFlags_ForceSvg        IconLookupFlags = int64(2)
	IconLookupFlags_UseBuiltin      IconLookupFlags = int64(4)
	IconLookupFlags_GenericFallback IconLookupFlags = int64(8)
	IconLookupFlags_ForceSize       IconLookupFlags = int64(16)
	IconLookupFlags_ForceRegular    IconLookupFlags = int64(32)
	IconLookupFlags_ForceSymbolic   IconLookupFlags = int64(64)
	IconLookupFlags_DirLtr          IconLookupFlags = int64(128)
	IconLookupFlags_DirRtl          IconLookupFlags = int64(256)
)

type InputHints uint32

const (
	InputHints_None               InputHints = int64(0)
	InputHints_Spellcheck         InputHints = int64(1)
	InputHints_NoSpellcheck       InputHints = int64(2)
	InputHints_WordCompletion     InputHints = int64(4)
	InputHints_Lowercase          InputHints = int64(8)
	InputHints_UppercaseChars     InputHints = int64(16)
	InputHints_UppercaseWords     InputHints = int64(32)
	InputHints_UppercaseSentences InputHints = int64(64)
	InputHints_InhibitOsk         InputHints = int64(128)
	InputHints_VerticalWriting    InputHints = int64(256)
	InputHints_Emoji              InputHints = int64(512)
	InputHints_NoEmoji            InputHints = int64(1024)
)

type JunctionSides uint32

const (
	JunctionSides_None              JunctionSides = int64(0)
	JunctionSides_CornerTopleft     JunctionSides = int64(1)
	JunctionSides_CornerTopright    JunctionSides = int64(2)
	JunctionSides_CornerBottomleft  JunctionSides = int64(4)
	JunctionSides_CornerBottomright JunctionSides = int64(8)
	JunctionSides_Top               JunctionSides = int64(3)
	JunctionSides_Bottom            JunctionSides = int64(12)
	JunctionSides_Left              JunctionSides = int64(5)
	JunctionSides_Right             JunctionSides = int64(10)
)

type PlacesOpenFlags uint32

const (
	PlacesOpenFlags_Normal    PlacesOpenFlags = int64(1)
	PlacesOpenFlags_NewTab    PlacesOpenFlags = int64(2)
	PlacesOpenFlags_NewWindow PlacesOpenFlags = int64(4)
)

type RcFlags uint32

const (
	RcFlags_Fg   RcFlags = int64(1)
	RcFlags_Bg   RcFlags = int64(2)
	RcFlags_Text RcFlags = int64(4)
	RcFlags_Base RcFlags = int64(8)
)

type RecentFilterFlags uint32

const (
	RecentFilterFlags_Uri         RecentFilterFlags = int64(1)
	RecentFilterFlags_DisplayName RecentFilterFlags = int64(2)
	RecentFilterFlags_MimeType    RecentFilterFlags = int64(4)
	RecentFilterFlags_Application RecentFilterFlags = int64(8)
	RecentFilterFlags_Group       RecentFilterFlags = int64(16)
	RecentFilterFlags_Age         RecentFilterFlags = int64(32)
)

type RegionFlags uint32

const (
	RegionFlags_Even   RegionFlags = int64(1)
	RegionFlags_Odd    RegionFlags = int64(2)
	RegionFlags_First  RegionFlags = int64(4)
	RegionFlags_Last   RegionFlags = int64(8)
	RegionFlags_Only   RegionFlags = int64(16)
	RegionFlags_Sorted RegionFlags = int64(32)
)

type StateFlags uint32

const (
	StateFlags_Normal       StateFlags = int64(0)
	StateFlags_Active       StateFlags = int64(1)
	StateFlags_Prelight     StateFlags = int64(2)
	StateFlags_Selected     StateFlags = int64(4)
	StateFlags_Insensitive  StateFlags = int64(8)
	StateFlags_Inconsistent StateFlags = int64(16)
	StateFlags_Focused      StateFlags = int64(32)
	StateFlags_Backdrop     StateFlags = int64(64)
	StateFlags_DirLtr       StateFlags = int64(128)
	StateFlags_DirRtl       StateFlags = int64(256)
	StateFlags_Link         StateFlags = int64(512)
	StateFlags_Visited      StateFlags = int64(1024)
	StateFlags_Checked      StateFlags = int64(2048)
	StateFlags_DropActive   StateFlags = int64(4096)
)

type StyleContextPrintFlags uint32

const (
	StyleContextPrintFlags_None      StyleContextPrintFlags = int64(0)
	StyleContextPrintFlags_Recurse   StyleContextPrintFlags = int64(1)
	StyleContextPrintFlags_ShowStyle StyleContextPrintFlags = int64(2)
)

type TargetFlags uint32

const (
	TargetFlags_SameApp     TargetFlags = int64(1)
	TargetFlags_SameWidget  TargetFlags = int64(2)
	TargetFlags_OtherApp    TargetFlags = int64(4)
	TargetFlags_OtherWidget TargetFlags = int64(8)
)

type TextSearchFlags uint32

const (
	TextSearchFlags_VisibleOnly     TextSearchFlags = int64(1)
	TextSearchFlags_TextOnly        TextSearchFlags = int64(2)
	TextSearchFlags_CaseInsensitive TextSearchFlags = int64(4)
)

type ToolPaletteDragTargets uint32

const (
	ToolPaletteDragTargets_Items  ToolPaletteDragTargets = int64(1)
	ToolPaletteDragTargets_Groups ToolPaletteDragTargets = int64(2)
)

type TreeModelFlags uint32

const (
	TreeModelFlags_ItersPersist TreeModelFlags = int64(1)
	TreeModelFlags_ListOnly     TreeModelFlags = int64(2)
)

type UIManagerItemType uint32

const (
	UIManagerItemType_Auto            UIManagerItemType = int64(0)
	UIManagerItemType_Menubar         UIManagerItemType = int64(1)
	UIManagerItemType_Menu            UIManagerItemType = int64(2)
	UIManagerItemType_Toolbar         UIManagerItemType = int64(4)
	UIManagerItemType_Placeholder     UIManagerItemType = int64(8)
	UIManagerItemType_Popup           UIManagerItemType = int64(16)
	UIManagerItemType_Menuitem        UIManagerItemType = int64(32)
	UIManagerItemType_Toolitem        UIManagerItemType = int64(64)
	UIManagerItemType_Separator       UIManagerItemType = int64(128)
	UIManagerItemType_Accelerator     UIManagerItemType = int64(256)
	UIManagerItemType_PopupWithAccels UIManagerItemType = int64(512)
)
