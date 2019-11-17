// Code generated - DO NOT EDIT.

package gtk

// AccelFlags is a representation of the C type GtkAccelFlags.
type AccelFlags int

const (
	// AccelFlags_Visible is a representation of the C type GTK_ACCEL_VISIBLE.
	AccelFlags_Visible AccelFlags = 1
	// AccelFlags_Locked is a representation of the C type GTK_ACCEL_LOCKED.
	AccelFlags_Locked AccelFlags = 2
	// AccelFlags_Mask is a representation of the C type GTK_ACCEL_MASK.
	AccelFlags_Mask AccelFlags = 7
)

// ApplicationInhibitFlags is a representation of the C type GtkApplicationInhibitFlags.
//
// since 3.4
type ApplicationInhibitFlags int

const (
	// ApplicationInhibitFlags_Logout is a representation of the C type GTK_APPLICATION_INHIBIT_LOGOUT.
	ApplicationInhibitFlags_Logout ApplicationInhibitFlags = 1
	// ApplicationInhibitFlags_Switch is a representation of the C type GTK_APPLICATION_INHIBIT_SWITCH.
	ApplicationInhibitFlags_Switch ApplicationInhibitFlags = 2
	// ApplicationInhibitFlags_Suspend is a representation of the C type GTK_APPLICATION_INHIBIT_SUSPEND.
	ApplicationInhibitFlags_Suspend ApplicationInhibitFlags = 4
	// ApplicationInhibitFlags_Idle is a representation of the C type GTK_APPLICATION_INHIBIT_IDLE.
	ApplicationInhibitFlags_Idle ApplicationInhibitFlags = 8
)

// AttachOptions is a representation of the C type GtkAttachOptions.
type AttachOptions int

const (
	// AttachOptions_Expand is a representation of the C type GTK_EXPAND.
	AttachOptions_Expand AttachOptions = 1
	// AttachOptions_Shrink is a representation of the C type GTK_SHRINK.
	AttachOptions_Shrink AttachOptions = 2
	// AttachOptions_Fill is a representation of the C type GTK_FILL.
	AttachOptions_Fill AttachOptions = 4
)

// CalendarDisplayOptions is a representation of the C type GtkCalendarDisplayOptions.
type CalendarDisplayOptions int

const (
	// CalendarDisplayOptions_ShowHeading is a representation of the C type GTK_CALENDAR_SHOW_HEADING.
	CalendarDisplayOptions_ShowHeading CalendarDisplayOptions = 1
	// CalendarDisplayOptions_ShowDayNames is a representation of the C type GTK_CALENDAR_SHOW_DAY_NAMES.
	CalendarDisplayOptions_ShowDayNames CalendarDisplayOptions = 2
	// CalendarDisplayOptions_NoMonthChange is a representation of the C type GTK_CALENDAR_NO_MONTH_CHANGE.
	CalendarDisplayOptions_NoMonthChange CalendarDisplayOptions = 4
	// CalendarDisplayOptions_ShowWeekNumbers is a representation of the C type GTK_CALENDAR_SHOW_WEEK_NUMBERS.
	CalendarDisplayOptions_ShowWeekNumbers CalendarDisplayOptions = 8
	// CalendarDisplayOptions_ShowDetails is a representation of the C type GTK_CALENDAR_SHOW_DETAILS.
	CalendarDisplayOptions_ShowDetails CalendarDisplayOptions = 32
)

// CellRendererState is a representation of the C type GtkCellRendererState.
type CellRendererState int

const (
	// CellRendererState_Selected is a representation of the C type GTK_CELL_RENDERER_SELECTED.
	CellRendererState_Selected CellRendererState = 1
	// CellRendererState_Prelit is a representation of the C type GTK_CELL_RENDERER_PRELIT.
	CellRendererState_Prelit CellRendererState = 2
	// CellRendererState_Insensitive is a representation of the C type GTK_CELL_RENDERER_INSENSITIVE.
	CellRendererState_Insensitive CellRendererState = 4
	// CellRendererState_Sorted is a representation of the C type GTK_CELL_RENDERER_SORTED.
	CellRendererState_Sorted CellRendererState = 8
	// CellRendererState_Focused is a representation of the C type GTK_CELL_RENDERER_FOCUSED.
	CellRendererState_Focused CellRendererState = 16
	// CellRendererState_Expandable is a representation of the C type GTK_CELL_RENDERER_EXPANDABLE.
	CellRendererState_Expandable CellRendererState = 32
	// CellRendererState_Expanded is a representation of the C type GTK_CELL_RENDERER_EXPANDED.
	CellRendererState_Expanded CellRendererState = 64
)

// DebugFlag is a representation of the C type GtkDebugFlag.
type DebugFlag int

const (
	// DebugFlag_Misc is a representation of the C type GTK_DEBUG_MISC.
	DebugFlag_Misc DebugFlag = 1
	// DebugFlag_Plugsocket is a representation of the C type GTK_DEBUG_PLUGSOCKET.
	DebugFlag_Plugsocket DebugFlag = 2
	// DebugFlag_Text is a representation of the C type GTK_DEBUG_TEXT.
	DebugFlag_Text DebugFlag = 4
	// DebugFlag_Tree is a representation of the C type GTK_DEBUG_TREE.
	DebugFlag_Tree DebugFlag = 8
	// DebugFlag_Updates is a representation of the C type GTK_DEBUG_UPDATES.
	DebugFlag_Updates DebugFlag = 16
	// DebugFlag_Keybindings is a representation of the C type GTK_DEBUG_KEYBINDINGS.
	DebugFlag_Keybindings DebugFlag = 32
	// DebugFlag_Multihead is a representation of the C type GTK_DEBUG_MULTIHEAD.
	DebugFlag_Multihead DebugFlag = 64
	// DebugFlag_Modules is a representation of the C type GTK_DEBUG_MODULES.
	DebugFlag_Modules DebugFlag = 128
	// DebugFlag_Geometry is a representation of the C type GTK_DEBUG_GEOMETRY.
	DebugFlag_Geometry DebugFlag = 256
	// DebugFlag_Icontheme is a representation of the C type GTK_DEBUG_ICONTHEME.
	DebugFlag_Icontheme DebugFlag = 512
	// DebugFlag_Printing is a representation of the C type GTK_DEBUG_PRINTING.
	DebugFlag_Printing DebugFlag = 1024
	// DebugFlag_Builder is a representation of the C type GTK_DEBUG_BUILDER.
	DebugFlag_Builder DebugFlag = 2048
	// DebugFlag_SizeRequest is a representation of the C type GTK_DEBUG_SIZE_REQUEST.
	DebugFlag_SizeRequest DebugFlag = 4096
	// DebugFlag_NoCssCache is a representation of the C type GTK_DEBUG_NO_CSS_CACHE.
	DebugFlag_NoCssCache DebugFlag = 8192
	// DebugFlag_Baselines is a representation of the C type GTK_DEBUG_BASELINES.
	DebugFlag_Baselines DebugFlag = 16384
	// DebugFlag_PixelCache is a representation of the C type GTK_DEBUG_PIXEL_CACHE.
	DebugFlag_PixelCache DebugFlag = 32768
	// DebugFlag_NoPixelCache is a representation of the C type GTK_DEBUG_NO_PIXEL_CACHE.
	DebugFlag_NoPixelCache DebugFlag = 65536
	// DebugFlag_Interactive is a representation of the C type GTK_DEBUG_INTERACTIVE.
	DebugFlag_Interactive DebugFlag = 131072
	// DebugFlag_Touchscreen is a representation of the C type GTK_DEBUG_TOUCHSCREEN.
	DebugFlag_Touchscreen DebugFlag = 262144
	// DebugFlag_Actions is a representation of the C type GTK_DEBUG_ACTIONS.
	DebugFlag_Actions DebugFlag = 524288
	// DebugFlag_Resize is a representation of the C type GTK_DEBUG_RESIZE.
	DebugFlag_Resize DebugFlag = 1048576
	// DebugFlag_Layout is a representation of the C type GTK_DEBUG_LAYOUT.
	DebugFlag_Layout DebugFlag = 2097152
)

// DestDefaults is a representation of the C type GtkDestDefaults.
type DestDefaults int

const (
	// DestDefaults_Motion is a representation of the C type GTK_DEST_DEFAULT_MOTION.
	DestDefaults_Motion DestDefaults = 1
	// DestDefaults_Highlight is a representation of the C type GTK_DEST_DEFAULT_HIGHLIGHT.
	DestDefaults_Highlight DestDefaults = 2
	// DestDefaults_Drop is a representation of the C type GTK_DEST_DEFAULT_DROP.
	DestDefaults_Drop DestDefaults = 4
	// DestDefaults_All is a representation of the C type GTK_DEST_DEFAULT_ALL.
	DestDefaults_All DestDefaults = 7
)

// DialogFlags is a representation of the C type GtkDialogFlags.
type DialogFlags int

const (
	// DialogFlags_Modal is a representation of the C type GTK_DIALOG_MODAL.
	DialogFlags_Modal DialogFlags = 1
	// DialogFlags_DestroyWithParent is a representation of the C type GTK_DIALOG_DESTROY_WITH_PARENT.
	DialogFlags_DestroyWithParent DialogFlags = 2
	// DialogFlags_UseHeaderBar is a representation of the C type GTK_DIALOG_USE_HEADER_BAR.
	DialogFlags_UseHeaderBar DialogFlags = 4
)

// EventControllerScrollFlags is a representation of the C type GtkEventControllerScrollFlags.
//
// since 3.24
type EventControllerScrollFlags int

const (
	// EventControllerScrollFlags_None is a representation of the C type GTK_EVENT_CONTROLLER_SCROLL_NONE.
	EventControllerScrollFlags_None EventControllerScrollFlags = 0
	// EventControllerScrollFlags_Vertical is a representation of the C type GTK_EVENT_CONTROLLER_SCROLL_VERTICAL.
	EventControllerScrollFlags_Vertical EventControllerScrollFlags = 1
	// EventControllerScrollFlags_Horizontal is a representation of the C type GTK_EVENT_CONTROLLER_SCROLL_HORIZONTAL.
	EventControllerScrollFlags_Horizontal EventControllerScrollFlags = 2
	// EventControllerScrollFlags_Discrete is a representation of the C type GTK_EVENT_CONTROLLER_SCROLL_DISCRETE.
	EventControllerScrollFlags_Discrete EventControllerScrollFlags = 4
	// EventControllerScrollFlags_Kinetic is a representation of the C type GTK_EVENT_CONTROLLER_SCROLL_KINETIC.
	EventControllerScrollFlags_Kinetic EventControllerScrollFlags = 8
	// EventControllerScrollFlags_BothAxes is a representation of the C type GTK_EVENT_CONTROLLER_SCROLL_BOTH_AXES.
	EventControllerScrollFlags_BothAxes EventControllerScrollFlags = 3
)

// FileFilterFlags is a representation of the C type GtkFileFilterFlags.
type FileFilterFlags int

const (
	// FileFilterFlags_Filename is a representation of the C type GTK_FILE_FILTER_FILENAME.
	FileFilterFlags_Filename FileFilterFlags = 1
	// FileFilterFlags_Uri is a representation of the C type GTK_FILE_FILTER_URI.
	FileFilterFlags_Uri FileFilterFlags = 2
	// FileFilterFlags_DisplayName is a representation of the C type GTK_FILE_FILTER_DISPLAY_NAME.
	FileFilterFlags_DisplayName FileFilterFlags = 4
	// FileFilterFlags_MimeType is a representation of the C type GTK_FILE_FILTER_MIME_TYPE.
	FileFilterFlags_MimeType FileFilterFlags = 8
)

// FontChooserLevel is a representation of the C type GtkFontChooserLevel.
type FontChooserLevel int

const (
	// FontChooserLevel_Family is a representation of the C type GTK_FONT_CHOOSER_LEVEL_FAMILY.
	FontChooserLevel_Family FontChooserLevel = 0
	// FontChooserLevel_Style is a representation of the C type GTK_FONT_CHOOSER_LEVEL_STYLE.
	FontChooserLevel_Style FontChooserLevel = 1
	// FontChooserLevel_Size is a representation of the C type GTK_FONT_CHOOSER_LEVEL_SIZE.
	FontChooserLevel_Size FontChooserLevel = 2
	// FontChooserLevel_Variations is a representation of the C type GTK_FONT_CHOOSER_LEVEL_VARIATIONS.
	FontChooserLevel_Variations FontChooserLevel = 4
	// FontChooserLevel_Features is a representation of the C type GTK_FONT_CHOOSER_LEVEL_FEATURES.
	FontChooserLevel_Features FontChooserLevel = 8
)

// IconLookupFlags is a representation of the C type GtkIconLookupFlags.
type IconLookupFlags int

const (
	// IconLookupFlags_NoSvg is a representation of the C type GTK_ICON_LOOKUP_NO_SVG.
	IconLookupFlags_NoSvg IconLookupFlags = 1
	// IconLookupFlags_ForceSvg is a representation of the C type GTK_ICON_LOOKUP_FORCE_SVG.
	IconLookupFlags_ForceSvg IconLookupFlags = 2
	// IconLookupFlags_UseBuiltin is a representation of the C type GTK_ICON_LOOKUP_USE_BUILTIN.
	IconLookupFlags_UseBuiltin IconLookupFlags = 4
	// IconLookupFlags_GenericFallback is a representation of the C type GTK_ICON_LOOKUP_GENERIC_FALLBACK.
	IconLookupFlags_GenericFallback IconLookupFlags = 8
	// IconLookupFlags_ForceSize is a representation of the C type GTK_ICON_LOOKUP_FORCE_SIZE.
	IconLookupFlags_ForceSize IconLookupFlags = 16
	// IconLookupFlags_ForceRegular is a representation of the C type GTK_ICON_LOOKUP_FORCE_REGULAR.
	IconLookupFlags_ForceRegular IconLookupFlags = 32
	// IconLookupFlags_ForceSymbolic is a representation of the C type GTK_ICON_LOOKUP_FORCE_SYMBOLIC.
	IconLookupFlags_ForceSymbolic IconLookupFlags = 64
	// IconLookupFlags_DirLtr is a representation of the C type GTK_ICON_LOOKUP_DIR_LTR.
	IconLookupFlags_DirLtr IconLookupFlags = 128
	// IconLookupFlags_DirRtl is a representation of the C type GTK_ICON_LOOKUP_DIR_RTL.
	IconLookupFlags_DirRtl IconLookupFlags = 256
)

// InputHints is a representation of the C type GtkInputHints.
//
// since 3.6
type InputHints int

const (
	// InputHints_None is a representation of the C type GTK_INPUT_HINT_NONE.
	InputHints_None InputHints = 0
	// InputHints_Spellcheck is a representation of the C type GTK_INPUT_HINT_SPELLCHECK.
	InputHints_Spellcheck InputHints = 1
	// InputHints_NoSpellcheck is a representation of the C type GTK_INPUT_HINT_NO_SPELLCHECK.
	InputHints_NoSpellcheck InputHints = 2
	// InputHints_WordCompletion is a representation of the C type GTK_INPUT_HINT_WORD_COMPLETION.
	InputHints_WordCompletion InputHints = 4
	// InputHints_Lowercase is a representation of the C type GTK_INPUT_HINT_LOWERCASE.
	InputHints_Lowercase InputHints = 8
	// InputHints_UppercaseChars is a representation of the C type GTK_INPUT_HINT_UPPERCASE_CHARS.
	InputHints_UppercaseChars InputHints = 16
	// InputHints_UppercaseWords is a representation of the C type GTK_INPUT_HINT_UPPERCASE_WORDS.
	InputHints_UppercaseWords InputHints = 32
	// InputHints_UppercaseSentences is a representation of the C type GTK_INPUT_HINT_UPPERCASE_SENTENCES.
	InputHints_UppercaseSentences InputHints = 64
	// InputHints_InhibitOsk is a representation of the C type GTK_INPUT_HINT_INHIBIT_OSK.
	InputHints_InhibitOsk InputHints = 128
	// InputHints_VerticalWriting is a representation of the C type GTK_INPUT_HINT_VERTICAL_WRITING.
	InputHints_VerticalWriting InputHints = 256
	// InputHints_Emoji is a representation of the C type GTK_INPUT_HINT_EMOJI.
	InputHints_Emoji InputHints = 512
	// InputHints_NoEmoji is a representation of the C type GTK_INPUT_HINT_NO_EMOJI.
	InputHints_NoEmoji InputHints = 1024
)

// JunctionSides is a representation of the C type GtkJunctionSides.
type JunctionSides int

const (
	// JunctionSides_None is a representation of the C type GTK_JUNCTION_NONE.
	JunctionSides_None JunctionSides = 0
	// JunctionSides_CornerTopleft is a representation of the C type GTK_JUNCTION_CORNER_TOPLEFT.
	JunctionSides_CornerTopleft JunctionSides = 1
	// JunctionSides_CornerTopright is a representation of the C type GTK_JUNCTION_CORNER_TOPRIGHT.
	JunctionSides_CornerTopright JunctionSides = 2
	// JunctionSides_CornerBottomleft is a representation of the C type GTK_JUNCTION_CORNER_BOTTOMLEFT.
	JunctionSides_CornerBottomleft JunctionSides = 4
	// JunctionSides_CornerBottomright is a representation of the C type GTK_JUNCTION_CORNER_BOTTOMRIGHT.
	JunctionSides_CornerBottomright JunctionSides = 8
	// JunctionSides_Top is a representation of the C type GTK_JUNCTION_TOP.
	JunctionSides_Top JunctionSides = 3
	// JunctionSides_Bottom is a representation of the C type GTK_JUNCTION_BOTTOM.
	JunctionSides_Bottom JunctionSides = 12
	// JunctionSides_Left is a representation of the C type GTK_JUNCTION_LEFT.
	JunctionSides_Left JunctionSides = 5
	// JunctionSides_Right is a representation of the C type GTK_JUNCTION_RIGHT.
	JunctionSides_Right JunctionSides = 10
)

// PlacesOpenFlags is a representation of the C type GtkPlacesOpenFlags.
type PlacesOpenFlags int

const (
	// PlacesOpenFlags_Normal is a representation of the C type GTK_PLACES_OPEN_NORMAL.
	PlacesOpenFlags_Normal PlacesOpenFlags = 1
	// PlacesOpenFlags_NewTab is a representation of the C type GTK_PLACES_OPEN_NEW_TAB.
	PlacesOpenFlags_NewTab PlacesOpenFlags = 2
	// PlacesOpenFlags_NewWindow is a representation of the C type GTK_PLACES_OPEN_NEW_WINDOW.
	PlacesOpenFlags_NewWindow PlacesOpenFlags = 4
)

// RcFlags is a representation of the C type GtkRcFlags.
type RcFlags int

const (
	// RcFlags_Fg is a representation of the C type GTK_RC_FG.
	RcFlags_Fg RcFlags = 1
	// RcFlags_Bg is a representation of the C type GTK_RC_BG.
	RcFlags_Bg RcFlags = 2
	// RcFlags_Text is a representation of the C type GTK_RC_TEXT.
	RcFlags_Text RcFlags = 4
	// RcFlags_Base is a representation of the C type GTK_RC_BASE.
	RcFlags_Base RcFlags = 8
)

// RecentFilterFlags is a representation of the C type GtkRecentFilterFlags.
type RecentFilterFlags int

const (
	// RecentFilterFlags_Uri is a representation of the C type GTK_RECENT_FILTER_URI.
	RecentFilterFlags_Uri RecentFilterFlags = 1
	// RecentFilterFlags_DisplayName is a representation of the C type GTK_RECENT_FILTER_DISPLAY_NAME.
	RecentFilterFlags_DisplayName RecentFilterFlags = 2
	// RecentFilterFlags_MimeType is a representation of the C type GTK_RECENT_FILTER_MIME_TYPE.
	RecentFilterFlags_MimeType RecentFilterFlags = 4
	// RecentFilterFlags_Application is a representation of the C type GTK_RECENT_FILTER_APPLICATION.
	RecentFilterFlags_Application RecentFilterFlags = 8
	// RecentFilterFlags_Group is a representation of the C type GTK_RECENT_FILTER_GROUP.
	RecentFilterFlags_Group RecentFilterFlags = 16
	// RecentFilterFlags_Age is a representation of the C type GTK_RECENT_FILTER_AGE.
	RecentFilterFlags_Age RecentFilterFlags = 32
)

// RegionFlags is a representation of the C type GtkRegionFlags.
type RegionFlags int

const (
	// RegionFlags_Even is a representation of the C type GTK_REGION_EVEN.
	RegionFlags_Even RegionFlags = 1
	// RegionFlags_Odd is a representation of the C type GTK_REGION_ODD.
	RegionFlags_Odd RegionFlags = 2
	// RegionFlags_First is a representation of the C type GTK_REGION_FIRST.
	RegionFlags_First RegionFlags = 4
	// RegionFlags_Last is a representation of the C type GTK_REGION_LAST.
	RegionFlags_Last RegionFlags = 8
	// RegionFlags_Only is a representation of the C type GTK_REGION_ONLY.
	RegionFlags_Only RegionFlags = 16
	// RegionFlags_Sorted is a representation of the C type GTK_REGION_SORTED.
	RegionFlags_Sorted RegionFlags = 32
)

// StateFlags is a representation of the C type GtkStateFlags.
type StateFlags int

const (
	// StateFlags_Normal is a representation of the C type GTK_STATE_FLAG_NORMAL.
	StateFlags_Normal StateFlags = 0
	// StateFlags_Active is a representation of the C type GTK_STATE_FLAG_ACTIVE.
	StateFlags_Active StateFlags = 1
	// StateFlags_Prelight is a representation of the C type GTK_STATE_FLAG_PRELIGHT.
	StateFlags_Prelight StateFlags = 2
	// StateFlags_Selected is a representation of the C type GTK_STATE_FLAG_SELECTED.
	StateFlags_Selected StateFlags = 4
	// StateFlags_Insensitive is a representation of the C type GTK_STATE_FLAG_INSENSITIVE.
	StateFlags_Insensitive StateFlags = 8
	// StateFlags_Inconsistent is a representation of the C type GTK_STATE_FLAG_INCONSISTENT.
	StateFlags_Inconsistent StateFlags = 16
	// StateFlags_Focused is a representation of the C type GTK_STATE_FLAG_FOCUSED.
	StateFlags_Focused StateFlags = 32
	// StateFlags_Backdrop is a representation of the C type GTK_STATE_FLAG_BACKDROP.
	StateFlags_Backdrop StateFlags = 64
	// StateFlags_DirLtr is a representation of the C type GTK_STATE_FLAG_DIR_LTR.
	StateFlags_DirLtr StateFlags = 128
	// StateFlags_DirRtl is a representation of the C type GTK_STATE_FLAG_DIR_RTL.
	StateFlags_DirRtl StateFlags = 256
	// StateFlags_Link is a representation of the C type GTK_STATE_FLAG_LINK.
	StateFlags_Link StateFlags = 512
	// StateFlags_Visited is a representation of the C type GTK_STATE_FLAG_VISITED.
	StateFlags_Visited StateFlags = 1024
	// StateFlags_Checked is a representation of the C type GTK_STATE_FLAG_CHECKED.
	StateFlags_Checked StateFlags = 2048
	// StateFlags_DropActive is a representation of the C type GTK_STATE_FLAG_DROP_ACTIVE.
	StateFlags_DropActive StateFlags = 4096
)

// StyleContextPrintFlags is a representation of the C type GtkStyleContextPrintFlags.
type StyleContextPrintFlags int

const (
	// StyleContextPrintFlags_None is a representation of the C type GTK_STYLE_CONTEXT_PRINT_NONE.
	StyleContextPrintFlags_None StyleContextPrintFlags = 0
	// StyleContextPrintFlags_Recurse is a representation of the C type GTK_STYLE_CONTEXT_PRINT_RECURSE.
	StyleContextPrintFlags_Recurse StyleContextPrintFlags = 1
	// StyleContextPrintFlags_ShowStyle is a representation of the C type GTK_STYLE_CONTEXT_PRINT_SHOW_STYLE.
	StyleContextPrintFlags_ShowStyle StyleContextPrintFlags = 2
)

// TargetFlags is a representation of the C type GtkTargetFlags.
type TargetFlags int

const (
	// TargetFlags_SameApp is a representation of the C type GTK_TARGET_SAME_APP.
	TargetFlags_SameApp TargetFlags = 1
	// TargetFlags_SameWidget is a representation of the C type GTK_TARGET_SAME_WIDGET.
	TargetFlags_SameWidget TargetFlags = 2
	// TargetFlags_OtherApp is a representation of the C type GTK_TARGET_OTHER_APP.
	TargetFlags_OtherApp TargetFlags = 4
	// TargetFlags_OtherWidget is a representation of the C type GTK_TARGET_OTHER_WIDGET.
	TargetFlags_OtherWidget TargetFlags = 8
)

// TextSearchFlags is a representation of the C type GtkTextSearchFlags.
type TextSearchFlags int

const (
	// TextSearchFlags_VisibleOnly is a representation of the C type GTK_TEXT_SEARCH_VISIBLE_ONLY.
	TextSearchFlags_VisibleOnly TextSearchFlags = 1
	// TextSearchFlags_TextOnly is a representation of the C type GTK_TEXT_SEARCH_TEXT_ONLY.
	TextSearchFlags_TextOnly TextSearchFlags = 2
	// TextSearchFlags_CaseInsensitive is a representation of the C type GTK_TEXT_SEARCH_CASE_INSENSITIVE.
	TextSearchFlags_CaseInsensitive TextSearchFlags = 4
)

// ToolPaletteDragTargets is a representation of the C type GtkToolPaletteDragTargets.
type ToolPaletteDragTargets int

const (
	// ToolPaletteDragTargets_Items is a representation of the C type GTK_TOOL_PALETTE_DRAG_ITEMS.
	ToolPaletteDragTargets_Items ToolPaletteDragTargets = 1
	// ToolPaletteDragTargets_Groups is a representation of the C type GTK_TOOL_PALETTE_DRAG_GROUPS.
	ToolPaletteDragTargets_Groups ToolPaletteDragTargets = 2
)

// TreeModelFlags is a representation of the C type GtkTreeModelFlags.
type TreeModelFlags int

const (
	// TreeModelFlags_ItersPersist is a representation of the C type GTK_TREE_MODEL_ITERS_PERSIST.
	TreeModelFlags_ItersPersist TreeModelFlags = 1
	// TreeModelFlags_ListOnly is a representation of the C type GTK_TREE_MODEL_LIST_ONLY.
	TreeModelFlags_ListOnly TreeModelFlags = 2
)

// UIManagerItemType is a representation of the C type GtkUIManagerItemType.
type UIManagerItemType int

const (
	// UIManagerItemType_Auto is a representation of the C type GTK_UI_MANAGER_AUTO.
	UIManagerItemType_Auto UIManagerItemType = 0
	// UIManagerItemType_Menubar is a representation of the C type GTK_UI_MANAGER_MENUBAR.
	UIManagerItemType_Menubar UIManagerItemType = 1
	// UIManagerItemType_Menu is a representation of the C type GTK_UI_MANAGER_MENU.
	UIManagerItemType_Menu UIManagerItemType = 2
	// UIManagerItemType_Toolbar is a representation of the C type GTK_UI_MANAGER_TOOLBAR.
	UIManagerItemType_Toolbar UIManagerItemType = 4
	// UIManagerItemType_Placeholder is a representation of the C type GTK_UI_MANAGER_PLACEHOLDER.
	UIManagerItemType_Placeholder UIManagerItemType = 8
	// UIManagerItemType_Popup is a representation of the C type GTK_UI_MANAGER_POPUP.
	UIManagerItemType_Popup UIManagerItemType = 16
	// UIManagerItemType_Menuitem is a representation of the C type GTK_UI_MANAGER_MENUITEM.
	UIManagerItemType_Menuitem UIManagerItemType = 32
	// UIManagerItemType_Toolitem is a representation of the C type GTK_UI_MANAGER_TOOLITEM.
	UIManagerItemType_Toolitem UIManagerItemType = 64
	// UIManagerItemType_Separator is a representation of the C type GTK_UI_MANAGER_SEPARATOR.
	UIManagerItemType_Separator UIManagerItemType = 128
	// UIManagerItemType_Accelerator is a representation of the C type GTK_UI_MANAGER_ACCELERATOR.
	UIManagerItemType_Accelerator UIManagerItemType = 256
	// UIManagerItemType_PopupWithAccels is a representation of the C type GTK_UI_MANAGER_POPUP_WITH_ACCELS.
	UIManagerItemType_PopupWithAccels UIManagerItemType = 512
)
