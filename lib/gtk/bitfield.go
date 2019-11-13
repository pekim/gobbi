// Code generated - DO NOT EDIT.

package gtk

// Accelflags is a representation of the C type GtkAccelFlags.
type Accelflags int

const (
	// Accelflags_Visible is a representation of the C type GTK_ACCEL_VISIBLE.
	Accelflags_Visible Accelflags = 1
	// Accelflags_Locked is a representation of the C type GTK_ACCEL_LOCKED.
	Accelflags_Locked Accelflags = 2
	// Accelflags_Mask is a representation of the C type GTK_ACCEL_MASK.
	Accelflags_Mask Accelflags = 7
)

// Applicationinhibitflags is a representation of the C type GtkApplicationInhibitFlags.
//
// since 3.4
type Applicationinhibitflags int

const (
	// Applicationinhibitflags_Logout is a representation of the C type GTK_APPLICATION_INHIBIT_LOGOUT.
	Applicationinhibitflags_Logout Applicationinhibitflags = 1
	// Applicationinhibitflags_Switch is a representation of the C type GTK_APPLICATION_INHIBIT_SWITCH.
	Applicationinhibitflags_Switch Applicationinhibitflags = 2
	// Applicationinhibitflags_Suspend is a representation of the C type GTK_APPLICATION_INHIBIT_SUSPEND.
	Applicationinhibitflags_Suspend Applicationinhibitflags = 4
	// Applicationinhibitflags_Idle is a representation of the C type GTK_APPLICATION_INHIBIT_IDLE.
	Applicationinhibitflags_Idle Applicationinhibitflags = 8
)

// Attachoptions is a representation of the C type GtkAttachOptions.
type Attachoptions int

const (
	// Attachoptions_Expand is a representation of the C type GTK_EXPAND.
	Attachoptions_Expand Attachoptions = 1
	// Attachoptions_Shrink is a representation of the C type GTK_SHRINK.
	Attachoptions_Shrink Attachoptions = 2
	// Attachoptions_Fill is a representation of the C type GTK_FILL.
	Attachoptions_Fill Attachoptions = 4
)

// Calendardisplayoptions is a representation of the C type GtkCalendarDisplayOptions.
type Calendardisplayoptions int

const (
	// Calendardisplayoptions_ShowHeading is a representation of the C type GTK_CALENDAR_SHOW_HEADING.
	Calendardisplayoptions_ShowHeading Calendardisplayoptions = 1
	// Calendardisplayoptions_ShowDayNames is a representation of the C type GTK_CALENDAR_SHOW_DAY_NAMES.
	Calendardisplayoptions_ShowDayNames Calendardisplayoptions = 2
	// Calendardisplayoptions_NoMonthChange is a representation of the C type GTK_CALENDAR_NO_MONTH_CHANGE.
	Calendardisplayoptions_NoMonthChange Calendardisplayoptions = 4
	// Calendardisplayoptions_ShowWeekNumbers is a representation of the C type GTK_CALENDAR_SHOW_WEEK_NUMBERS.
	Calendardisplayoptions_ShowWeekNumbers Calendardisplayoptions = 8
	// Calendardisplayoptions_ShowDetails is a representation of the C type GTK_CALENDAR_SHOW_DETAILS.
	Calendardisplayoptions_ShowDetails Calendardisplayoptions = 32
)

// Cellrendererstate is a representation of the C type GtkCellRendererState.
type Cellrendererstate int

const (
	// Cellrendererstate_Selected is a representation of the C type GTK_CELL_RENDERER_SELECTED.
	Cellrendererstate_Selected Cellrendererstate = 1
	// Cellrendererstate_Prelit is a representation of the C type GTK_CELL_RENDERER_PRELIT.
	Cellrendererstate_Prelit Cellrendererstate = 2
	// Cellrendererstate_Insensitive is a representation of the C type GTK_CELL_RENDERER_INSENSITIVE.
	Cellrendererstate_Insensitive Cellrendererstate = 4
	// Cellrendererstate_Sorted is a representation of the C type GTK_CELL_RENDERER_SORTED.
	Cellrendererstate_Sorted Cellrendererstate = 8
	// Cellrendererstate_Focused is a representation of the C type GTK_CELL_RENDERER_FOCUSED.
	Cellrendererstate_Focused Cellrendererstate = 16
	// Cellrendererstate_Expandable is a representation of the C type GTK_CELL_RENDERER_EXPANDABLE.
	Cellrendererstate_Expandable Cellrendererstate = 32
	// Cellrendererstate_Expanded is a representation of the C type GTK_CELL_RENDERER_EXPANDED.
	Cellrendererstate_Expanded Cellrendererstate = 64
)

// Debugflag is a representation of the C type GtkDebugFlag.
type Debugflag int

const (
	// Debugflag_Misc is a representation of the C type GTK_DEBUG_MISC.
	Debugflag_Misc Debugflag = 1
	// Debugflag_Plugsocket is a representation of the C type GTK_DEBUG_PLUGSOCKET.
	Debugflag_Plugsocket Debugflag = 2
	// Debugflag_Text is a representation of the C type GTK_DEBUG_TEXT.
	Debugflag_Text Debugflag = 4
	// Debugflag_Tree is a representation of the C type GTK_DEBUG_TREE.
	Debugflag_Tree Debugflag = 8
	// Debugflag_Updates is a representation of the C type GTK_DEBUG_UPDATES.
	Debugflag_Updates Debugflag = 16
	// Debugflag_Keybindings is a representation of the C type GTK_DEBUG_KEYBINDINGS.
	Debugflag_Keybindings Debugflag = 32
	// Debugflag_Multihead is a representation of the C type GTK_DEBUG_MULTIHEAD.
	Debugflag_Multihead Debugflag = 64
	// Debugflag_Modules is a representation of the C type GTK_DEBUG_MODULES.
	Debugflag_Modules Debugflag = 128
	// Debugflag_Geometry is a representation of the C type GTK_DEBUG_GEOMETRY.
	Debugflag_Geometry Debugflag = 256
	// Debugflag_Icontheme is a representation of the C type GTK_DEBUG_ICONTHEME.
	Debugflag_Icontheme Debugflag = 512
	// Debugflag_Printing is a representation of the C type GTK_DEBUG_PRINTING.
	Debugflag_Printing Debugflag = 1024
	// Debugflag_Builder is a representation of the C type GTK_DEBUG_BUILDER.
	Debugflag_Builder Debugflag = 2048
	// Debugflag_SizeRequest is a representation of the C type GTK_DEBUG_SIZE_REQUEST.
	Debugflag_SizeRequest Debugflag = 4096
	// Debugflag_NoCssCache is a representation of the C type GTK_DEBUG_NO_CSS_CACHE.
	Debugflag_NoCssCache Debugflag = 8192
	// Debugflag_Baselines is a representation of the C type GTK_DEBUG_BASELINES.
	Debugflag_Baselines Debugflag = 16384
	// Debugflag_PixelCache is a representation of the C type GTK_DEBUG_PIXEL_CACHE.
	Debugflag_PixelCache Debugflag = 32768
	// Debugflag_NoPixelCache is a representation of the C type GTK_DEBUG_NO_PIXEL_CACHE.
	Debugflag_NoPixelCache Debugflag = 65536
	// Debugflag_Interactive is a representation of the C type GTK_DEBUG_INTERACTIVE.
	Debugflag_Interactive Debugflag = 131072
	// Debugflag_Touchscreen is a representation of the C type GTK_DEBUG_TOUCHSCREEN.
	Debugflag_Touchscreen Debugflag = 262144
	// Debugflag_Actions is a representation of the C type GTK_DEBUG_ACTIONS.
	Debugflag_Actions Debugflag = 524288
	// Debugflag_Resize is a representation of the C type GTK_DEBUG_RESIZE.
	Debugflag_Resize Debugflag = 1048576
	// Debugflag_Layout is a representation of the C type GTK_DEBUG_LAYOUT.
	Debugflag_Layout Debugflag = 2097152
)

// Destdefaults is a representation of the C type GtkDestDefaults.
type Destdefaults int

const (
	// Destdefaults_Motion is a representation of the C type GTK_DEST_DEFAULT_MOTION.
	Destdefaults_Motion Destdefaults = 1
	// Destdefaults_Highlight is a representation of the C type GTK_DEST_DEFAULT_HIGHLIGHT.
	Destdefaults_Highlight Destdefaults = 2
	// Destdefaults_Drop is a representation of the C type GTK_DEST_DEFAULT_DROP.
	Destdefaults_Drop Destdefaults = 4
	// Destdefaults_All is a representation of the C type GTK_DEST_DEFAULT_ALL.
	Destdefaults_All Destdefaults = 7
)

// Dialogflags is a representation of the C type GtkDialogFlags.
type Dialogflags int

const (
	// Dialogflags_Modal is a representation of the C type GTK_DIALOG_MODAL.
	Dialogflags_Modal Dialogflags = 1
	// Dialogflags_DestroyWithParent is a representation of the C type GTK_DIALOG_DESTROY_WITH_PARENT.
	Dialogflags_DestroyWithParent Dialogflags = 2
	// Dialogflags_UseHeaderBar is a representation of the C type GTK_DIALOG_USE_HEADER_BAR.
	Dialogflags_UseHeaderBar Dialogflags = 4
)

// Filefilterflags is a representation of the C type GtkFileFilterFlags.
type Filefilterflags int

const (
	// Filefilterflags_Filename is a representation of the C type GTK_FILE_FILTER_FILENAME.
	Filefilterflags_Filename Filefilterflags = 1
	// Filefilterflags_Uri is a representation of the C type GTK_FILE_FILTER_URI.
	Filefilterflags_Uri Filefilterflags = 2
	// Filefilterflags_DisplayName is a representation of the C type GTK_FILE_FILTER_DISPLAY_NAME.
	Filefilterflags_DisplayName Filefilterflags = 4
	// Filefilterflags_MimeType is a representation of the C type GTK_FILE_FILTER_MIME_TYPE.
	Filefilterflags_MimeType Filefilterflags = 8
)

// Iconlookupflags is a representation of the C type GtkIconLookupFlags.
type Iconlookupflags int

const (
	// Iconlookupflags_NoSvg is a representation of the C type GTK_ICON_LOOKUP_NO_SVG.
	Iconlookupflags_NoSvg Iconlookupflags = 1
	// Iconlookupflags_ForceSvg is a representation of the C type GTK_ICON_LOOKUP_FORCE_SVG.
	Iconlookupflags_ForceSvg Iconlookupflags = 2
	// Iconlookupflags_UseBuiltin is a representation of the C type GTK_ICON_LOOKUP_USE_BUILTIN.
	Iconlookupflags_UseBuiltin Iconlookupflags = 4
	// Iconlookupflags_GenericFallback is a representation of the C type GTK_ICON_LOOKUP_GENERIC_FALLBACK.
	Iconlookupflags_GenericFallback Iconlookupflags = 8
	// Iconlookupflags_ForceSize is a representation of the C type GTK_ICON_LOOKUP_FORCE_SIZE.
	Iconlookupflags_ForceSize Iconlookupflags = 16
	// Iconlookupflags_ForceRegular is a representation of the C type GTK_ICON_LOOKUP_FORCE_REGULAR.
	Iconlookupflags_ForceRegular Iconlookupflags = 32
	// Iconlookupflags_ForceSymbolic is a representation of the C type GTK_ICON_LOOKUP_FORCE_SYMBOLIC.
	Iconlookupflags_ForceSymbolic Iconlookupflags = 64
	// Iconlookupflags_DirLtr is a representation of the C type GTK_ICON_LOOKUP_DIR_LTR.
	Iconlookupflags_DirLtr Iconlookupflags = 128
	// Iconlookupflags_DirRtl is a representation of the C type GTK_ICON_LOOKUP_DIR_RTL.
	Iconlookupflags_DirRtl Iconlookupflags = 256
)

// Inputhints is a representation of the C type GtkInputHints.
//
// since 3.6
type Inputhints int

const (
	// Inputhints_None is a representation of the C type GTK_INPUT_HINT_NONE.
	Inputhints_None Inputhints = 0
	// Inputhints_Spellcheck is a representation of the C type GTK_INPUT_HINT_SPELLCHECK.
	Inputhints_Spellcheck Inputhints = 1
	// Inputhints_NoSpellcheck is a representation of the C type GTK_INPUT_HINT_NO_SPELLCHECK.
	Inputhints_NoSpellcheck Inputhints = 2
	// Inputhints_WordCompletion is a representation of the C type GTK_INPUT_HINT_WORD_COMPLETION.
	Inputhints_WordCompletion Inputhints = 4
	// Inputhints_Lowercase is a representation of the C type GTK_INPUT_HINT_LOWERCASE.
	Inputhints_Lowercase Inputhints = 8
	// Inputhints_UppercaseChars is a representation of the C type GTK_INPUT_HINT_UPPERCASE_CHARS.
	Inputhints_UppercaseChars Inputhints = 16
	// Inputhints_UppercaseWords is a representation of the C type GTK_INPUT_HINT_UPPERCASE_WORDS.
	Inputhints_UppercaseWords Inputhints = 32
	// Inputhints_UppercaseSentences is a representation of the C type GTK_INPUT_HINT_UPPERCASE_SENTENCES.
	Inputhints_UppercaseSentences Inputhints = 64
	// Inputhints_InhibitOsk is a representation of the C type GTK_INPUT_HINT_INHIBIT_OSK.
	Inputhints_InhibitOsk Inputhints = 128
	// Inputhints_VerticalWriting is a representation of the C type GTK_INPUT_HINT_VERTICAL_WRITING.
	Inputhints_VerticalWriting Inputhints = 256
	// Inputhints_Emoji is a representation of the C type GTK_INPUT_HINT_EMOJI.
	Inputhints_Emoji Inputhints = 512
	// Inputhints_NoEmoji is a representation of the C type GTK_INPUT_HINT_NO_EMOJI.
	Inputhints_NoEmoji Inputhints = 1024
)

// Junctionsides is a representation of the C type GtkJunctionSides.
type Junctionsides int

const (
	// Junctionsides_None is a representation of the C type GTK_JUNCTION_NONE.
	Junctionsides_None Junctionsides = 0
	// Junctionsides_CornerTopleft is a representation of the C type GTK_JUNCTION_CORNER_TOPLEFT.
	Junctionsides_CornerTopleft Junctionsides = 1
	// Junctionsides_CornerTopright is a representation of the C type GTK_JUNCTION_CORNER_TOPRIGHT.
	Junctionsides_CornerTopright Junctionsides = 2
	// Junctionsides_CornerBottomleft is a representation of the C type GTK_JUNCTION_CORNER_BOTTOMLEFT.
	Junctionsides_CornerBottomleft Junctionsides = 4
	// Junctionsides_CornerBottomright is a representation of the C type GTK_JUNCTION_CORNER_BOTTOMRIGHT.
	Junctionsides_CornerBottomright Junctionsides = 8
	// Junctionsides_Top is a representation of the C type GTK_JUNCTION_TOP.
	Junctionsides_Top Junctionsides = 3
	// Junctionsides_Bottom is a representation of the C type GTK_JUNCTION_BOTTOM.
	Junctionsides_Bottom Junctionsides = 12
	// Junctionsides_Left is a representation of the C type GTK_JUNCTION_LEFT.
	Junctionsides_Left Junctionsides = 5
	// Junctionsides_Right is a representation of the C type GTK_JUNCTION_RIGHT.
	Junctionsides_Right Junctionsides = 10
)

// Placesopenflags is a representation of the C type GtkPlacesOpenFlags.
type Placesopenflags int

const (
	// Placesopenflags_Normal is a representation of the C type GTK_PLACES_OPEN_NORMAL.
	Placesopenflags_Normal Placesopenflags = 1
	// Placesopenflags_NewTab is a representation of the C type GTK_PLACES_OPEN_NEW_TAB.
	Placesopenflags_NewTab Placesopenflags = 2
	// Placesopenflags_NewWindow is a representation of the C type GTK_PLACES_OPEN_NEW_WINDOW.
	Placesopenflags_NewWindow Placesopenflags = 4
)

// Rcflags is a representation of the C type GtkRcFlags.
type Rcflags int

const (
	// Rcflags_Fg is a representation of the C type GTK_RC_FG.
	Rcflags_Fg Rcflags = 1
	// Rcflags_Bg is a representation of the C type GTK_RC_BG.
	Rcflags_Bg Rcflags = 2
	// Rcflags_Text is a representation of the C type GTK_RC_TEXT.
	Rcflags_Text Rcflags = 4
	// Rcflags_Base is a representation of the C type GTK_RC_BASE.
	Rcflags_Base Rcflags = 8
)

// Recentfilterflags is a representation of the C type GtkRecentFilterFlags.
type Recentfilterflags int

const (
	// Recentfilterflags_Uri is a representation of the C type GTK_RECENT_FILTER_URI.
	Recentfilterflags_Uri Recentfilterflags = 1
	// Recentfilterflags_DisplayName is a representation of the C type GTK_RECENT_FILTER_DISPLAY_NAME.
	Recentfilterflags_DisplayName Recentfilterflags = 2
	// Recentfilterflags_MimeType is a representation of the C type GTK_RECENT_FILTER_MIME_TYPE.
	Recentfilterflags_MimeType Recentfilterflags = 4
	// Recentfilterflags_Application is a representation of the C type GTK_RECENT_FILTER_APPLICATION.
	Recentfilterflags_Application Recentfilterflags = 8
	// Recentfilterflags_Group is a representation of the C type GTK_RECENT_FILTER_GROUP.
	Recentfilterflags_Group Recentfilterflags = 16
	// Recentfilterflags_Age is a representation of the C type GTK_RECENT_FILTER_AGE.
	Recentfilterflags_Age Recentfilterflags = 32
)

// Regionflags is a representation of the C type GtkRegionFlags.
type Regionflags int

const (
	// Regionflags_Even is a representation of the C type GTK_REGION_EVEN.
	Regionflags_Even Regionflags = 1
	// Regionflags_Odd is a representation of the C type GTK_REGION_ODD.
	Regionflags_Odd Regionflags = 2
	// Regionflags_First is a representation of the C type GTK_REGION_FIRST.
	Regionflags_First Regionflags = 4
	// Regionflags_Last is a representation of the C type GTK_REGION_LAST.
	Regionflags_Last Regionflags = 8
	// Regionflags_Only is a representation of the C type GTK_REGION_ONLY.
	Regionflags_Only Regionflags = 16
	// Regionflags_Sorted is a representation of the C type GTK_REGION_SORTED.
	Regionflags_Sorted Regionflags = 32
)

// Stateflags is a representation of the C type GtkStateFlags.
type Stateflags int

const (
	// Stateflags_Normal is a representation of the C type GTK_STATE_FLAG_NORMAL.
	Stateflags_Normal Stateflags = 0
	// Stateflags_Active is a representation of the C type GTK_STATE_FLAG_ACTIVE.
	Stateflags_Active Stateflags = 1
	// Stateflags_Prelight is a representation of the C type GTK_STATE_FLAG_PRELIGHT.
	Stateflags_Prelight Stateflags = 2
	// Stateflags_Selected is a representation of the C type GTK_STATE_FLAG_SELECTED.
	Stateflags_Selected Stateflags = 4
	// Stateflags_Insensitive is a representation of the C type GTK_STATE_FLAG_INSENSITIVE.
	Stateflags_Insensitive Stateflags = 8
	// Stateflags_Inconsistent is a representation of the C type GTK_STATE_FLAG_INCONSISTENT.
	Stateflags_Inconsistent Stateflags = 16
	// Stateflags_Focused is a representation of the C type GTK_STATE_FLAG_FOCUSED.
	Stateflags_Focused Stateflags = 32
	// Stateflags_Backdrop is a representation of the C type GTK_STATE_FLAG_BACKDROP.
	Stateflags_Backdrop Stateflags = 64
	// Stateflags_DirLtr is a representation of the C type GTK_STATE_FLAG_DIR_LTR.
	Stateflags_DirLtr Stateflags = 128
	// Stateflags_DirRtl is a representation of the C type GTK_STATE_FLAG_DIR_RTL.
	Stateflags_DirRtl Stateflags = 256
	// Stateflags_Link is a representation of the C type GTK_STATE_FLAG_LINK.
	Stateflags_Link Stateflags = 512
	// Stateflags_Visited is a representation of the C type GTK_STATE_FLAG_VISITED.
	Stateflags_Visited Stateflags = 1024
	// Stateflags_Checked is a representation of the C type GTK_STATE_FLAG_CHECKED.
	Stateflags_Checked Stateflags = 2048
	// Stateflags_DropActive is a representation of the C type GTK_STATE_FLAG_DROP_ACTIVE.
	Stateflags_DropActive Stateflags = 4096
)

// Stylecontextprintflags is a representation of the C type GtkStyleContextPrintFlags.
type Stylecontextprintflags int

const (
	// Stylecontextprintflags_None is a representation of the C type GTK_STYLE_CONTEXT_PRINT_NONE.
	Stylecontextprintflags_None Stylecontextprintflags = 0
	// Stylecontextprintflags_Recurse is a representation of the C type GTK_STYLE_CONTEXT_PRINT_RECURSE.
	Stylecontextprintflags_Recurse Stylecontextprintflags = 1
	// Stylecontextprintflags_ShowStyle is a representation of the C type GTK_STYLE_CONTEXT_PRINT_SHOW_STYLE.
	Stylecontextprintflags_ShowStyle Stylecontextprintflags = 2
)

// Targetflags is a representation of the C type GtkTargetFlags.
type Targetflags int

const (
	// Targetflags_SameApp is a representation of the C type GTK_TARGET_SAME_APP.
	Targetflags_SameApp Targetflags = 1
	// Targetflags_SameWidget is a representation of the C type GTK_TARGET_SAME_WIDGET.
	Targetflags_SameWidget Targetflags = 2
	// Targetflags_OtherApp is a representation of the C type GTK_TARGET_OTHER_APP.
	Targetflags_OtherApp Targetflags = 4
	// Targetflags_OtherWidget is a representation of the C type GTK_TARGET_OTHER_WIDGET.
	Targetflags_OtherWidget Targetflags = 8
)

// Textsearchflags is a representation of the C type GtkTextSearchFlags.
type Textsearchflags int

const (
	// Textsearchflags_VisibleOnly is a representation of the C type GTK_TEXT_SEARCH_VISIBLE_ONLY.
	Textsearchflags_VisibleOnly Textsearchflags = 1
	// Textsearchflags_TextOnly is a representation of the C type GTK_TEXT_SEARCH_TEXT_ONLY.
	Textsearchflags_TextOnly Textsearchflags = 2
	// Textsearchflags_CaseInsensitive is a representation of the C type GTK_TEXT_SEARCH_CASE_INSENSITIVE.
	Textsearchflags_CaseInsensitive Textsearchflags = 4
)

// Toolpalettedragtargets is a representation of the C type GtkToolPaletteDragTargets.
type Toolpalettedragtargets int

const (
	// Toolpalettedragtargets_Items is a representation of the C type GTK_TOOL_PALETTE_DRAG_ITEMS.
	Toolpalettedragtargets_Items Toolpalettedragtargets = 1
	// Toolpalettedragtargets_Groups is a representation of the C type GTK_TOOL_PALETTE_DRAG_GROUPS.
	Toolpalettedragtargets_Groups Toolpalettedragtargets = 2
)

// Treemodelflags is a representation of the C type GtkTreeModelFlags.
type Treemodelflags int

const (
	// Treemodelflags_ItersPersist is a representation of the C type GTK_TREE_MODEL_ITERS_PERSIST.
	Treemodelflags_ItersPersist Treemodelflags = 1
	// Treemodelflags_ListOnly is a representation of the C type GTK_TREE_MODEL_LIST_ONLY.
	Treemodelflags_ListOnly Treemodelflags = 2
)

// Uimanageritemtype is a representation of the C type GtkUIManagerItemType.
type Uimanageritemtype int

const (
	// Uimanageritemtype_Auto is a representation of the C type GTK_UI_MANAGER_AUTO.
	Uimanageritemtype_Auto Uimanageritemtype = 0
	// Uimanageritemtype_Menubar is a representation of the C type GTK_UI_MANAGER_MENUBAR.
	Uimanageritemtype_Menubar Uimanageritemtype = 1
	// Uimanageritemtype_Menu is a representation of the C type GTK_UI_MANAGER_MENU.
	Uimanageritemtype_Menu Uimanageritemtype = 2
	// Uimanageritemtype_Toolbar is a representation of the C type GTK_UI_MANAGER_TOOLBAR.
	Uimanageritemtype_Toolbar Uimanageritemtype = 4
	// Uimanageritemtype_Placeholder is a representation of the C type GTK_UI_MANAGER_PLACEHOLDER.
	Uimanageritemtype_Placeholder Uimanageritemtype = 8
	// Uimanageritemtype_Popup is a representation of the C type GTK_UI_MANAGER_POPUP.
	Uimanageritemtype_Popup Uimanageritemtype = 16
	// Uimanageritemtype_Menuitem is a representation of the C type GTK_UI_MANAGER_MENUITEM.
	Uimanageritemtype_Menuitem Uimanageritemtype = 32
	// Uimanageritemtype_Toolitem is a representation of the C type GTK_UI_MANAGER_TOOLITEM.
	Uimanageritemtype_Toolitem Uimanageritemtype = 64
	// Uimanageritemtype_Separator is a representation of the C type GTK_UI_MANAGER_SEPARATOR.
	Uimanageritemtype_Separator Uimanageritemtype = 128
	// Uimanageritemtype_Accelerator is a representation of the C type GTK_UI_MANAGER_ACCELERATOR.
	Uimanageritemtype_Accelerator Uimanageritemtype = 256
	// Uimanageritemtype_PopupWithAccels is a representation of the C type GTK_UI_MANAGER_POPUP_WITH_ACCELS.
	Uimanageritemtype_PopupWithAccels Uimanageritemtype = 512
)
