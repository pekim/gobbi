// This is a generated file - DO NOT EDIT

package gtk

import "C"

type AccelFlags C.GtkAccelFlags

const (
	GTK_ACCEL_VISIBLE AccelFlags = 1
	GTK_ACCEL_LOCKED  AccelFlags = 2
	GTK_ACCEL_MASK    AccelFlags = 7
)

type AttachOptions C.GtkAttachOptions

const (
	GTK_EXPAND AttachOptions = 1
	GTK_SHRINK AttachOptions = 2
	GTK_FILL   AttachOptions = 4
)

type CalendarDisplayOptions C.GtkCalendarDisplayOptions

const (
	GTK_CALENDAR_SHOW_HEADING      CalendarDisplayOptions = 1
	GTK_CALENDAR_SHOW_DAY_NAMES    CalendarDisplayOptions = 2
	GTK_CALENDAR_NO_MONTH_CHANGE   CalendarDisplayOptions = 4
	GTK_CALENDAR_SHOW_WEEK_NUMBERS CalendarDisplayOptions = 8
	GTK_CALENDAR_SHOW_DETAILS      CalendarDisplayOptions = 32
)

type CellRendererState C.GtkCellRendererState

const (
	GTK_CELL_RENDERER_SELECTED    CellRendererState = 1
	GTK_CELL_RENDERER_PRELIT      CellRendererState = 2
	GTK_CELL_RENDERER_INSENSITIVE CellRendererState = 4
	GTK_CELL_RENDERER_SORTED      CellRendererState = 8
	GTK_CELL_RENDERER_FOCUSED     CellRendererState = 16
	GTK_CELL_RENDERER_EXPANDABLE  CellRendererState = 32
	GTK_CELL_RENDERER_EXPANDED    CellRendererState = 64
)

type DebugFlag C.GtkDebugFlag

const (
	GTK_DEBUG_MISC           DebugFlag = 1
	GTK_DEBUG_PLUGSOCKET     DebugFlag = 2
	GTK_DEBUG_TEXT           DebugFlag = 4
	GTK_DEBUG_TREE           DebugFlag = 8
	GTK_DEBUG_UPDATES        DebugFlag = 16
	GTK_DEBUG_KEYBINDINGS    DebugFlag = 32
	GTK_DEBUG_MULTIHEAD      DebugFlag = 64
	GTK_DEBUG_MODULES        DebugFlag = 128
	GTK_DEBUG_GEOMETRY       DebugFlag = 256
	GTK_DEBUG_ICONTHEME      DebugFlag = 512
	GTK_DEBUG_PRINTING       DebugFlag = 1024
	GTK_DEBUG_BUILDER        DebugFlag = 2048
	GTK_DEBUG_SIZE_REQUEST   DebugFlag = 4096
	GTK_DEBUG_NO_CSS_CACHE   DebugFlag = 8192
	GTK_DEBUG_BASELINES      DebugFlag = 16384
	GTK_DEBUG_PIXEL_CACHE    DebugFlag = 32768
	GTK_DEBUG_NO_PIXEL_CACHE DebugFlag = 65536
	GTK_DEBUG_INTERACTIVE    DebugFlag = 131072
	GTK_DEBUG_TOUCHSCREEN    DebugFlag = 262144
	GTK_DEBUG_ACTIONS        DebugFlag = 524288
	GTK_DEBUG_RESIZE         DebugFlag = 1048576
	GTK_DEBUG_LAYOUT         DebugFlag = 2097152
)

type DestDefaults C.GtkDestDefaults

const (
	GTK_DEST_DEFAULT_MOTION    DestDefaults = 1
	GTK_DEST_DEFAULT_HIGHLIGHT DestDefaults = 2
	GTK_DEST_DEFAULT_DROP      DestDefaults = 4
	GTK_DEST_DEFAULT_ALL       DestDefaults = 7
)

type DialogFlags C.GtkDialogFlags

const (
	GTK_DIALOG_MODAL               DialogFlags = 1
	GTK_DIALOG_DESTROY_WITH_PARENT DialogFlags = 2
	GTK_DIALOG_USE_HEADER_BAR      DialogFlags = 4
)

type FileFilterFlags C.GtkFileFilterFlags

const (
	GTK_FILE_FILTER_FILENAME     FileFilterFlags = 1
	GTK_FILE_FILTER_URI          FileFilterFlags = 2
	GTK_FILE_FILTER_DISPLAY_NAME FileFilterFlags = 4
	GTK_FILE_FILTER_MIME_TYPE    FileFilterFlags = 8
)

type IconLookupFlags C.GtkIconLookupFlags

const (
	GTK_ICON_LOOKUP_NO_SVG           IconLookupFlags = 1
	GTK_ICON_LOOKUP_FORCE_SVG        IconLookupFlags = 2
	GTK_ICON_LOOKUP_USE_BUILTIN      IconLookupFlags = 4
	GTK_ICON_LOOKUP_GENERIC_FALLBACK IconLookupFlags = 8
	GTK_ICON_LOOKUP_FORCE_SIZE       IconLookupFlags = 16
	GTK_ICON_LOOKUP_FORCE_REGULAR    IconLookupFlags = 32
	GTK_ICON_LOOKUP_FORCE_SYMBOLIC   IconLookupFlags = 64
	GTK_ICON_LOOKUP_DIR_LTR          IconLookupFlags = 128
	GTK_ICON_LOOKUP_DIR_RTL          IconLookupFlags = 256
)

type JunctionSides C.GtkJunctionSides

const (
	GTK_JUNCTION_NONE               JunctionSides = 0
	GTK_JUNCTION_CORNER_TOPLEFT     JunctionSides = 1
	GTK_JUNCTION_CORNER_TOPRIGHT    JunctionSides = 2
	GTK_JUNCTION_CORNER_BOTTOMLEFT  JunctionSides = 4
	GTK_JUNCTION_CORNER_BOTTOMRIGHT JunctionSides = 8
	GTK_JUNCTION_TOP                JunctionSides = 3
	GTK_JUNCTION_BOTTOM             JunctionSides = 12
	GTK_JUNCTION_LEFT               JunctionSides = 5
	GTK_JUNCTION_RIGHT              JunctionSides = 10
)

type PlacesOpenFlags C.GtkPlacesOpenFlags

const (
	GTK_PLACES_OPEN_NORMAL     PlacesOpenFlags = 1
	GTK_PLACES_OPEN_NEW_TAB    PlacesOpenFlags = 2
	GTK_PLACES_OPEN_NEW_WINDOW PlacesOpenFlags = 4
)

type RcFlags C.GtkRcFlags

const (
	GTK_RC_FG   RcFlags = 1
	GTK_RC_BG   RcFlags = 2
	GTK_RC_TEXT RcFlags = 4
	GTK_RC_BASE RcFlags = 8
)

type RecentFilterFlags C.GtkRecentFilterFlags

const (
	GTK_RECENT_FILTER_URI          RecentFilterFlags = 1
	GTK_RECENT_FILTER_DISPLAY_NAME RecentFilterFlags = 2
	GTK_RECENT_FILTER_MIME_TYPE    RecentFilterFlags = 4
	GTK_RECENT_FILTER_APPLICATION  RecentFilterFlags = 8
	GTK_RECENT_FILTER_GROUP        RecentFilterFlags = 16
	GTK_RECENT_FILTER_AGE          RecentFilterFlags = 32
)

type RegionFlags C.GtkRegionFlags

const (
	GTK_REGION_EVEN   RegionFlags = 1
	GTK_REGION_ODD    RegionFlags = 2
	GTK_REGION_FIRST  RegionFlags = 4
	GTK_REGION_LAST   RegionFlags = 8
	GTK_REGION_ONLY   RegionFlags = 16
	GTK_REGION_SORTED RegionFlags = 32
)

type StateFlags C.GtkStateFlags

const (
	GTK_STATE_FLAG_NORMAL       StateFlags = 0
	GTK_STATE_FLAG_ACTIVE       StateFlags = 1
	GTK_STATE_FLAG_PRELIGHT     StateFlags = 2
	GTK_STATE_FLAG_SELECTED     StateFlags = 4
	GTK_STATE_FLAG_INSENSITIVE  StateFlags = 8
	GTK_STATE_FLAG_INCONSISTENT StateFlags = 16
	GTK_STATE_FLAG_FOCUSED      StateFlags = 32
	GTK_STATE_FLAG_BACKDROP     StateFlags = 64
	GTK_STATE_FLAG_DIR_LTR      StateFlags = 128
	GTK_STATE_FLAG_DIR_RTL      StateFlags = 256
	GTK_STATE_FLAG_LINK         StateFlags = 512
	GTK_STATE_FLAG_VISITED      StateFlags = 1024
	GTK_STATE_FLAG_CHECKED      StateFlags = 2048
	GTK_STATE_FLAG_DROP_ACTIVE  StateFlags = 4096
)

type TargetFlags C.GtkTargetFlags

const (
	GTK_TARGET_SAME_APP     TargetFlags = 1
	GTK_TARGET_SAME_WIDGET  TargetFlags = 2
	GTK_TARGET_OTHER_APP    TargetFlags = 4
	GTK_TARGET_OTHER_WIDGET TargetFlags = 8
)

type TextSearchFlags C.GtkTextSearchFlags

const (
	GTK_TEXT_SEARCH_VISIBLE_ONLY     TextSearchFlags = 1
	GTK_TEXT_SEARCH_TEXT_ONLY        TextSearchFlags = 2
	GTK_TEXT_SEARCH_CASE_INSENSITIVE TextSearchFlags = 4
)

type ToolPaletteDragTargets C.GtkToolPaletteDragTargets

const (
	GTK_TOOL_PALETTE_DRAG_ITEMS  ToolPaletteDragTargets = 1
	GTK_TOOL_PALETTE_DRAG_GROUPS ToolPaletteDragTargets = 2
)

type TreeModelFlags C.GtkTreeModelFlags

const (
	GTK_TREE_MODEL_ITERS_PERSIST TreeModelFlags = 1
	GTK_TREE_MODEL_LIST_ONLY     TreeModelFlags = 2
)

type UIManagerItemType C.GtkUIManagerItemType

const (
	GTK_UI_MANAGER_AUTO              UIManagerItemType = 0
	GTK_UI_MANAGER_MENUBAR           UIManagerItemType = 1
	GTK_UI_MANAGER_MENU              UIManagerItemType = 2
	GTK_UI_MANAGER_TOOLBAR           UIManagerItemType = 4
	GTK_UI_MANAGER_PLACEHOLDER       UIManagerItemType = 8
	GTK_UI_MANAGER_POPUP             UIManagerItemType = 16
	GTK_UI_MANAGER_MENUITEM          UIManagerItemType = 32
	GTK_UI_MANAGER_TOOLITEM          UIManagerItemType = 64
	GTK_UI_MANAGER_SEPARATOR         UIManagerItemType = 128
	GTK_UI_MANAGER_ACCELERATOR       UIManagerItemType = 256
	GTK_UI_MANAGER_POPUP_WITH_ACCELS UIManagerItemType = 512
)
