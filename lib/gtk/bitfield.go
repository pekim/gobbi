// Code generated - DO NOT EDIT.

package gtk

// Accelflags is a representation of the C type AccelFlags.
type Accelflags int

const (
	Accelflags_Visible Accelflags = 1
	Accelflags_Locked  Accelflags = 2
	Accelflags_Mask    Accelflags = 7
)

// Applicationinhibitflags is a representation of the C type ApplicationInhibitFlags.
//
// since 3.4
type Applicationinhibitflags int

const (
	Applicationinhibitflags_Logout  Applicationinhibitflags = 1
	Applicationinhibitflags_Switch  Applicationinhibitflags = 2
	Applicationinhibitflags_Suspend Applicationinhibitflags = 4
	Applicationinhibitflags_Idle    Applicationinhibitflags = 8
)

// Attachoptions is a representation of the C type AttachOptions.
type Attachoptions int

const (
	Attachoptions_Expand Attachoptions = 1
	Attachoptions_Shrink Attachoptions = 2
	Attachoptions_Fill   Attachoptions = 4
)

// Calendardisplayoptions is a representation of the C type CalendarDisplayOptions.
type Calendardisplayoptions int

const (
	Calendardisplayoptions_ShowHeading     Calendardisplayoptions = 1
	Calendardisplayoptions_ShowDayNames    Calendardisplayoptions = 2
	Calendardisplayoptions_NoMonthChange   Calendardisplayoptions = 4
	Calendardisplayoptions_ShowWeekNumbers Calendardisplayoptions = 8
	Calendardisplayoptions_ShowDetails     Calendardisplayoptions = 32
)

// Cellrendererstate is a representation of the C type CellRendererState.
type Cellrendererstate int

const (
	Cellrendererstate_Selected    Cellrendererstate = 1
	Cellrendererstate_Prelit      Cellrendererstate = 2
	Cellrendererstate_Insensitive Cellrendererstate = 4
	Cellrendererstate_Sorted      Cellrendererstate = 8
	Cellrendererstate_Focused     Cellrendererstate = 16
	Cellrendererstate_Expandable  Cellrendererstate = 32
	Cellrendererstate_Expanded    Cellrendererstate = 64
)

// Debugflag is a representation of the C type DebugFlag.
type Debugflag int

const (
	Debugflag_Misc         Debugflag = 1
	Debugflag_Plugsocket   Debugflag = 2
	Debugflag_Text         Debugflag = 4
	Debugflag_Tree         Debugflag = 8
	Debugflag_Updates      Debugflag = 16
	Debugflag_Keybindings  Debugflag = 32
	Debugflag_Multihead    Debugflag = 64
	Debugflag_Modules      Debugflag = 128
	Debugflag_Geometry     Debugflag = 256
	Debugflag_Icontheme    Debugflag = 512
	Debugflag_Printing     Debugflag = 1024
	Debugflag_Builder      Debugflag = 2048
	Debugflag_SizeRequest  Debugflag = 4096
	Debugflag_NoCssCache   Debugflag = 8192
	Debugflag_Baselines    Debugflag = 16384
	Debugflag_PixelCache   Debugflag = 32768
	Debugflag_NoPixelCache Debugflag = 65536
	Debugflag_Interactive  Debugflag = 131072
	Debugflag_Touchscreen  Debugflag = 262144
	Debugflag_Actions      Debugflag = 524288
	Debugflag_Resize       Debugflag = 1048576
	Debugflag_Layout       Debugflag = 2097152
)

// Destdefaults is a representation of the C type DestDefaults.
type Destdefaults int

const (
	Destdefaults_Motion    Destdefaults = 1
	Destdefaults_Highlight Destdefaults = 2
	Destdefaults_Drop      Destdefaults = 4
	Destdefaults_All       Destdefaults = 7
)

// Dialogflags is a representation of the C type DialogFlags.
type Dialogflags int

const (
	Dialogflags_Modal             Dialogflags = 1
	Dialogflags_DestroyWithParent Dialogflags = 2
	Dialogflags_UseHeaderBar      Dialogflags = 4
)

// Filefilterflags is a representation of the C type FileFilterFlags.
type Filefilterflags int

const (
	Filefilterflags_Filename    Filefilterflags = 1
	Filefilterflags_Uri         Filefilterflags = 2
	Filefilterflags_DisplayName Filefilterflags = 4
	Filefilterflags_MimeType    Filefilterflags = 8
)

// Iconlookupflags is a representation of the C type IconLookupFlags.
type Iconlookupflags int

const (
	Iconlookupflags_NoSvg           Iconlookupflags = 1
	Iconlookupflags_ForceSvg        Iconlookupflags = 2
	Iconlookupflags_UseBuiltin      Iconlookupflags = 4
	Iconlookupflags_GenericFallback Iconlookupflags = 8
	Iconlookupflags_ForceSize       Iconlookupflags = 16
	Iconlookupflags_ForceRegular    Iconlookupflags = 32
	Iconlookupflags_ForceSymbolic   Iconlookupflags = 64
	Iconlookupflags_DirLtr          Iconlookupflags = 128
	Iconlookupflags_DirRtl          Iconlookupflags = 256
)

// Inputhints is a representation of the C type InputHints.
//
// since 3.6
type Inputhints int

const (
	Inputhints_None               Inputhints = 0
	Inputhints_Spellcheck         Inputhints = 1
	Inputhints_NoSpellcheck       Inputhints = 2
	Inputhints_WordCompletion     Inputhints = 4
	Inputhints_Lowercase          Inputhints = 8
	Inputhints_UppercaseChars     Inputhints = 16
	Inputhints_UppercaseWords     Inputhints = 32
	Inputhints_UppercaseSentences Inputhints = 64
	Inputhints_InhibitOsk         Inputhints = 128
	Inputhints_VerticalWriting    Inputhints = 256
	Inputhints_Emoji              Inputhints = 512
	Inputhints_NoEmoji            Inputhints = 1024
)

// Junctionsides is a representation of the C type JunctionSides.
type Junctionsides int

const (
	Junctionsides_None              Junctionsides = 0
	Junctionsides_CornerTopleft     Junctionsides = 1
	Junctionsides_CornerTopright    Junctionsides = 2
	Junctionsides_CornerBottomleft  Junctionsides = 4
	Junctionsides_CornerBottomright Junctionsides = 8
	Junctionsides_Top               Junctionsides = 3
	Junctionsides_Bottom            Junctionsides = 12
	Junctionsides_Left              Junctionsides = 5
	Junctionsides_Right             Junctionsides = 10
)

// Placesopenflags is a representation of the C type PlacesOpenFlags.
type Placesopenflags int

const (
	Placesopenflags_Normal    Placesopenflags = 1
	Placesopenflags_NewTab    Placesopenflags = 2
	Placesopenflags_NewWindow Placesopenflags = 4
)

// Rcflags is a representation of the C type RcFlags.
type Rcflags int

const (
	Rcflags_Fg   Rcflags = 1
	Rcflags_Bg   Rcflags = 2
	Rcflags_Text Rcflags = 4
	Rcflags_Base Rcflags = 8
)

// Recentfilterflags is a representation of the C type RecentFilterFlags.
type Recentfilterflags int

const (
	Recentfilterflags_Uri         Recentfilterflags = 1
	Recentfilterflags_DisplayName Recentfilterflags = 2
	Recentfilterflags_MimeType    Recentfilterflags = 4
	Recentfilterflags_Application Recentfilterflags = 8
	Recentfilterflags_Group       Recentfilterflags = 16
	Recentfilterflags_Age         Recentfilterflags = 32
)

// Regionflags is a representation of the C type RegionFlags.
type Regionflags int

const (
	Regionflags_Even   Regionflags = 1
	Regionflags_Odd    Regionflags = 2
	Regionflags_First  Regionflags = 4
	Regionflags_Last   Regionflags = 8
	Regionflags_Only   Regionflags = 16
	Regionflags_Sorted Regionflags = 32
)

// Stateflags is a representation of the C type StateFlags.
type Stateflags int

const (
	Stateflags_Normal       Stateflags = 0
	Stateflags_Active       Stateflags = 1
	Stateflags_Prelight     Stateflags = 2
	Stateflags_Selected     Stateflags = 4
	Stateflags_Insensitive  Stateflags = 8
	Stateflags_Inconsistent Stateflags = 16
	Stateflags_Focused      Stateflags = 32
	Stateflags_Backdrop     Stateflags = 64
	Stateflags_DirLtr       Stateflags = 128
	Stateflags_DirRtl       Stateflags = 256
	Stateflags_Link         Stateflags = 512
	Stateflags_Visited      Stateflags = 1024
	Stateflags_Checked      Stateflags = 2048
	Stateflags_DropActive   Stateflags = 4096
)

// Stylecontextprintflags is a representation of the C type StyleContextPrintFlags.
type Stylecontextprintflags int

const (
	Stylecontextprintflags_None      Stylecontextprintflags = 0
	Stylecontextprintflags_Recurse   Stylecontextprintflags = 1
	Stylecontextprintflags_ShowStyle Stylecontextprintflags = 2
)

// Targetflags is a representation of the C type TargetFlags.
type Targetflags int

const (
	Targetflags_SameApp     Targetflags = 1
	Targetflags_SameWidget  Targetflags = 2
	Targetflags_OtherApp    Targetflags = 4
	Targetflags_OtherWidget Targetflags = 8
)

// Textsearchflags is a representation of the C type TextSearchFlags.
type Textsearchflags int

const (
	Textsearchflags_VisibleOnly     Textsearchflags = 1
	Textsearchflags_TextOnly        Textsearchflags = 2
	Textsearchflags_CaseInsensitive Textsearchflags = 4
)

// Toolpalettedragtargets is a representation of the C type ToolPaletteDragTargets.
type Toolpalettedragtargets int

const (
	Toolpalettedragtargets_Items  Toolpalettedragtargets = 1
	Toolpalettedragtargets_Groups Toolpalettedragtargets = 2
)

// Treemodelflags is a representation of the C type TreeModelFlags.
type Treemodelflags int

const (
	Treemodelflags_ItersPersist Treemodelflags = 1
	Treemodelflags_ListOnly     Treemodelflags = 2
)

// Uimanageritemtype is a representation of the C type UIManagerItemType.
type Uimanageritemtype int

const (
	Uimanageritemtype_Auto            Uimanageritemtype = 0
	Uimanageritemtype_Menubar         Uimanageritemtype = 1
	Uimanageritemtype_Menu            Uimanageritemtype = 2
	Uimanageritemtype_Toolbar         Uimanageritemtype = 4
	Uimanageritemtype_Placeholder     Uimanageritemtype = 8
	Uimanageritemtype_Popup           Uimanageritemtype = 16
	Uimanageritemtype_Menuitem        Uimanageritemtype = 32
	Uimanageritemtype_Toolitem        Uimanageritemtype = 64
	Uimanageritemtype_Separator       Uimanageritemtype = 128
	Uimanageritemtype_Accelerator     Uimanageritemtype = 256
	Uimanageritemtype_PopupWithAccels Uimanageritemtype = 512
)
