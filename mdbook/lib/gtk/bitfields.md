# `gtk` bitfields

## `AccelFlags`

Accelerator flags used with gtk_accel_group_connect().

- GTK_ACCEL_VISIBLE = 1
- GTK_ACCEL_LOCKED = 2
- GTK_ACCEL_MASK = 7

C - `GtkAccelFlags`

## `ApplicationInhibitFlags`

Types of user actions that may be blocked by gtk_application_inhibit().

- GTK_APPLICATION_INHIBIT_LOGOUT = 1
- GTK_APPLICATION_INHIBIT_SWITCH = 2
- GTK_APPLICATION_INHIBIT_SUSPEND = 4
- GTK_APPLICATION_INHIBIT_IDLE = 8

C - `GtkApplicationInhibitFlags`

## `AttachOptions`

Denotes the expansion properties that a widget will have when it (or its
parent) is resized.

- GTK_EXPAND = 1
- GTK_SHRINK = 2
- GTK_FILL = 4

C - `GtkAttachOptions`

## `CalendarDisplayOptions`

These options can be used to influence the display and behaviour of a #GtkCalendar.

- GTK_CALENDAR_SHOW_HEADING = 1
- GTK_CALENDAR_SHOW_DAY_NAMES = 2
- GTK_CALENDAR_NO_MONTH_CHANGE = 4
- GTK_CALENDAR_SHOW_WEEK_NUMBERS = 8
- GTK_CALENDAR_SHOW_DETAILS = 32

C - `GtkCalendarDisplayOptions`

## `CellRendererState`

Tells how a cell is to be rendered.

- GTK_CELL_RENDERER_SELECTED = 1
- GTK_CELL_RENDERER_PRELIT = 2
- GTK_CELL_RENDERER_INSENSITIVE = 4
- GTK_CELL_RENDERER_SORTED = 8
- GTK_CELL_RENDERER_FOCUSED = 16
- GTK_CELL_RENDERER_EXPANDABLE = 32
- GTK_CELL_RENDERER_EXPANDED = 64

C - `GtkCellRendererState`

## `DebugFlag`



- GTK_DEBUG_MISC = 1
- GTK_DEBUG_PLUGSOCKET = 2
- GTK_DEBUG_TEXT = 4
- GTK_DEBUG_TREE = 8
- GTK_DEBUG_UPDATES = 16
- GTK_DEBUG_KEYBINDINGS = 32
- GTK_DEBUG_MULTIHEAD = 64
- GTK_DEBUG_MODULES = 128
- GTK_DEBUG_GEOMETRY = 256
- GTK_DEBUG_ICONTHEME = 512
- GTK_DEBUG_PRINTING = 1024
- GTK_DEBUG_BUILDER = 2048
- GTK_DEBUG_SIZE_REQUEST = 4096
- GTK_DEBUG_NO_CSS_CACHE = 8192
- GTK_DEBUG_BASELINES = 16384
- GTK_DEBUG_PIXEL_CACHE = 32768
- GTK_DEBUG_NO_PIXEL_CACHE = 65536
- GTK_DEBUG_INTERACTIVE = 131072
- GTK_DEBUG_TOUCHSCREEN = 262144
- GTK_DEBUG_ACTIONS = 524288
- GTK_DEBUG_RESIZE = 1048576
- GTK_DEBUG_LAYOUT = 2097152

C - `GtkDebugFlag`

## `DestDefaults`

The #GtkDestDefaults enumeration specifies the various
types of action that will be taken on behalf
of the user for a drag destination site.

- GTK_DEST_DEFAULT_MOTION = 1
- GTK_DEST_DEFAULT_HIGHLIGHT = 2
- GTK_DEST_DEFAULT_DROP = 4
- GTK_DEST_DEFAULT_ALL = 7

C - `GtkDestDefaults`

## `DialogFlags`

Flags used to influence dialog construction.

- GTK_DIALOG_MODAL = 1
- GTK_DIALOG_DESTROY_WITH_PARENT = 2
- GTK_DIALOG_USE_HEADER_BAR = 4

C - `GtkDialogFlags`

## `FileFilterFlags`

These flags indicate what parts of a #GtkFileFilterInfo struct
are filled or need to be filled.

- GTK_FILE_FILTER_FILENAME = 1
- GTK_FILE_FILTER_URI = 2
- GTK_FILE_FILTER_DISPLAY_NAME = 4
- GTK_FILE_FILTER_MIME_TYPE = 8

C - `GtkFileFilterFlags`

## `IconLookupFlags`

Used to specify options for gtk_icon_theme_lookup_icon()

- GTK_ICON_LOOKUP_NO_SVG = 1
- GTK_ICON_LOOKUP_FORCE_SVG = 2
- GTK_ICON_LOOKUP_USE_BUILTIN = 4
- GTK_ICON_LOOKUP_GENERIC_FALLBACK = 8
- GTK_ICON_LOOKUP_FORCE_SIZE = 16
- GTK_ICON_LOOKUP_FORCE_REGULAR = 32
- GTK_ICON_LOOKUP_FORCE_SYMBOLIC = 64
- GTK_ICON_LOOKUP_DIR_LTR = 128
- GTK_ICON_LOOKUP_DIR_RTL = 256

C - `GtkIconLookupFlags`

## `InputHints`

Describes hints that might be taken into account by input methods
or applications. Note that input methods may already tailor their
behaviour according to the #GtkInputPurpose of the entry.

Some common sense is expected when using these flags - mixing
@GTK_INPUT_HINT_LOWERCASE with any of the uppercase hints makes no sense.

This enumeration may be extended in the future; input methods should
ignore unknown values.

- GTK_INPUT_HINT_NONE = 0
- GTK_INPUT_HINT_SPELLCHECK = 1
- GTK_INPUT_HINT_NO_SPELLCHECK = 2
- GTK_INPUT_HINT_WORD_COMPLETION = 4
- GTK_INPUT_HINT_LOWERCASE = 8
- GTK_INPUT_HINT_UPPERCASE_CHARS = 16
- GTK_INPUT_HINT_UPPERCASE_WORDS = 32
- GTK_INPUT_HINT_UPPERCASE_SENTENCES = 64
- GTK_INPUT_HINT_INHIBIT_OSK = 128
- GTK_INPUT_HINT_VERTICAL_WRITING = 256
- GTK_INPUT_HINT_EMOJI = 512
- GTK_INPUT_HINT_NO_EMOJI = 1024

C - `GtkInputHints`

## `JunctionSides`

Describes how a rendered element connects to adjacent elements.

- GTK_JUNCTION_NONE = 0
- GTK_JUNCTION_CORNER_TOPLEFT = 1
- GTK_JUNCTION_CORNER_TOPRIGHT = 2
- GTK_JUNCTION_CORNER_BOTTOMLEFT = 4
- GTK_JUNCTION_CORNER_BOTTOMRIGHT = 8
- GTK_JUNCTION_TOP = 3
- GTK_JUNCTION_BOTTOM = 12
- GTK_JUNCTION_LEFT = 5
- GTK_JUNCTION_RIGHT = 10

C - `GtkJunctionSides`

## `PlacesOpenFlags`

These flags serve two purposes.  First, the application can call gtk_places_sidebar_set_open_flags()
using these flags as a bitmask.  This tells the sidebar that the application is able to open
folders selected from the sidebar in various ways, for example, in new tabs or in new windows in
addition to the normal mode.

Second, when one of these values gets passed back to the application in the
&num;GtkPlacesSidebar::open-location signal, it means that the application should
open the selected location in the normal way, in a new tab, or in a new
window.  The sidebar takes care of determining the desired way to open the location,
based on the modifier keys that the user is pressing at the time the selection is made.

If the application never calls gtk_places_sidebar_set_open_flags(), then the sidebar will only
use #GTK_PLACES_OPEN_NORMAL in the #GtkPlacesSidebar::open-location signal.  This is the
default mode of operation.

- GTK_PLACES_OPEN_NORMAL = 1
- GTK_PLACES_OPEN_NEW_TAB = 2
- GTK_PLACES_OPEN_NEW_WINDOW = 4

C - `GtkPlacesOpenFlags`

## `RcFlags`

Deprecated

- GTK_RC_FG = 1
- GTK_RC_BG = 2
- GTK_RC_TEXT = 4
- GTK_RC_BASE = 8

C - `GtkRcFlags`

## `RecentFilterFlags`

These flags indicate what parts of a #GtkRecentFilterInfo struct
are filled or need to be filled.

- GTK_RECENT_FILTER_URI = 1
- GTK_RECENT_FILTER_DISPLAY_NAME = 2
- GTK_RECENT_FILTER_MIME_TYPE = 4
- GTK_RECENT_FILTER_APPLICATION = 8
- GTK_RECENT_FILTER_GROUP = 16
- GTK_RECENT_FILTER_AGE = 32

C - `GtkRecentFilterFlags`

## `RegionFlags`

Describes a region within a widget.

- GTK_REGION_EVEN = 1
- GTK_REGION_ODD = 2
- GTK_REGION_FIRST = 4
- GTK_REGION_LAST = 8
- GTK_REGION_ONLY = 16
- GTK_REGION_SORTED = 32

C - `GtkRegionFlags`

## `StateFlags`

Describes a widget state. Widget states are used to match the widget
against CSS pseudo-classes. Note that GTK extends the regular CSS
classes and sometimes uses different names.

- GTK_STATE_FLAG_NORMAL = 0
- GTK_STATE_FLAG_ACTIVE = 1
- GTK_STATE_FLAG_PRELIGHT = 2
- GTK_STATE_FLAG_SELECTED = 4
- GTK_STATE_FLAG_INSENSITIVE = 8
- GTK_STATE_FLAG_INCONSISTENT = 16
- GTK_STATE_FLAG_FOCUSED = 32
- GTK_STATE_FLAG_BACKDROP = 64
- GTK_STATE_FLAG_DIR_LTR = 128
- GTK_STATE_FLAG_DIR_RTL = 256
- GTK_STATE_FLAG_LINK = 512
- GTK_STATE_FLAG_VISITED = 1024
- GTK_STATE_FLAG_CHECKED = 2048
- GTK_STATE_FLAG_DROP_ACTIVE = 4096

C - `GtkStateFlags`

## `StyleContextPrintFlags`

Flags that modify the behavior of gtk_style_context_to_string().
New values may be added to this enumeration.

- GTK_STYLE_CONTEXT_PRINT_NONE = 0
- GTK_STYLE_CONTEXT_PRINT_RECURSE = 1
- GTK_STYLE_CONTEXT_PRINT_SHOW_STYLE = 2

C - `GtkStyleContextPrintFlags`

## `TargetFlags`

The #GtkTargetFlags enumeration is used to specify
constraints on a #GtkTargetEntry.

- GTK_TARGET_SAME_APP = 1
- GTK_TARGET_SAME_WIDGET = 2
- GTK_TARGET_OTHER_APP = 4
- GTK_TARGET_OTHER_WIDGET = 8

C - `GtkTargetFlags`

## `TextSearchFlags`

Flags affecting how a search is done.

If neither #GTK_TEXT_SEARCH_VISIBLE_ONLY nor #GTK_TEXT_SEARCH_TEXT_ONLY are
enabled, the match must be exact; the special 0xFFFC character will match
embedded pixbufs or child widgets.

- GTK_TEXT_SEARCH_VISIBLE_ONLY = 1
- GTK_TEXT_SEARCH_TEXT_ONLY = 2
- GTK_TEXT_SEARCH_CASE_INSENSITIVE = 4

C - `GtkTextSearchFlags`

## `ToolPaletteDragTargets`

Flags used to specify the supported drag targets.

- GTK_TOOL_PALETTE_DRAG_ITEMS = 1
- GTK_TOOL_PALETTE_DRAG_GROUPS = 2

C - `GtkToolPaletteDragTargets`

## `TreeModelFlags`

These flags indicate various properties of a #GtkTreeModel.

They are returned by gtk_tree_model_get_flags(), and must be
static for the lifetime of the object. A more complete description
of #GTK_TREE_MODEL_ITERS_PERSIST can be found in the overview of
this section.

- GTK_TREE_MODEL_ITERS_PERSIST = 1
- GTK_TREE_MODEL_LIST_ONLY = 2

C - `GtkTreeModelFlags`

## `UIManagerItemType`

These enumeration values are used by gtk_ui_manager_add_ui() to determine
what UI element to create.

- GTK_UI_MANAGER_AUTO = 0
- GTK_UI_MANAGER_MENUBAR = 1
- GTK_UI_MANAGER_MENU = 2
- GTK_UI_MANAGER_TOOLBAR = 4
- GTK_UI_MANAGER_PLACEHOLDER = 8
- GTK_UI_MANAGER_POPUP = 16
- GTK_UI_MANAGER_MENUITEM = 32
- GTK_UI_MANAGER_TOOLITEM = 64
- GTK_UI_MANAGER_SEPARATOR = 128
- GTK_UI_MANAGER_ACCELERATOR = 256
- GTK_UI_MANAGER_POPUP_WITH_ACCELS = 512

C - `GtkUIManagerItemType`

EXT_ONLY = %!s(int=2)
- GTK_TEXT_SEARCH_CASE_INSENSITIVE = %!s(int=4)
## `ToolPaletteDragTargets`

Flags used to specify the supported drag targets.

C - `GtkToolPaletteDragTargets`

- GTK_TOOL_PALETTE_DRAG_ITEMS = %!s(int=1)
- GTK_TOOL_PALETTE_DRAG_GROUPS = %!s(int=2)
## `TreeModelFlags`

These flags indicate various properties of a #GtkTreeModel.

They are returned by gtk_tree_model_get_flags(), and must be
static for the lifetime of the object. A more complete description
of #GTK_TREE_MODEL_ITERS_PERSIST can be found in the overview of
this section.

C - `GtkTreeModelFlags`

- GTK_TREE_MODEL_ITERS_PERSIST = %!s(int=1)
- GTK_TREE_MODEL_LIST_ONLY = %!s(int=2)
## `UIManagerItemType`

These enumeration values are used by gtk_ui_manager_add_ui() to determine
what UI element to create.

C - `GtkUIManagerItemType`

- GTK_UI_MANAGER_AUTO = %!s(int=0)
- GTK_UI_MANAGER_MENUBAR = %!s(int=1)
- GTK_UI_MANAGER_MENU = %!s(int=2)
- GTK_UI_MANAGER_TOOLBAR = %!s(int=4)
- GTK_UI_MANAGER_PLACEHOLDER = %!s(int=8)
- GTK_UI_MANAGER_POPUP = %!s(int=16)
- GTK_UI_MANAGER_MENUITEM = %!s(int=32)
- GTK_UI_MANAGER_TOOLITEM = %!s(int=64)
- GTK_UI_MANAGER_SEPARATOR = %!s(int=128)
- GTK_UI_MANAGER_ACCELERATOR = %!s(int=256)
- GTK_UI_MANAGER_POPUP_WITH_ACCELS = %!s(int=512)
