# `gtk` bitfields

## `AccelFlags`

Accelerator flags used with gtk_accel_group_connect().

C - `GtkAccelFlags`

### GTK_ACCEL_VISIBLE = 1
Accelerator is visible

### GTK_ACCEL_LOCKED = 2
Accelerator not removable

### GTK_ACCEL_MASK = 7
Mask


## `ApplicationInhibitFlags`

Types of user actions that may be blocked by gtk_application_inhibit().

C - `GtkApplicationInhibitFlags`

### GTK_APPLICATION_INHIBIT_LOGOUT = 1
Inhibit ending the user session
    by logging out or by shutting down the computer

### GTK_APPLICATION_INHIBIT_SWITCH = 2
Inhibit user switching

### GTK_APPLICATION_INHIBIT_SUSPEND = 4
Inhibit suspending the
    session or computer

### GTK_APPLICATION_INHIBIT_IDLE = 8
Inhibit the session being
    marked as idle (and possibly locked)


## `AttachOptions`

Denotes the expansion properties that a widget will have when it (or its
parent) is resized.

C - `GtkAttachOptions`

### GTK_EXPAND = 1
the widget should expand to take up any extra space in its
container that has been allocated.

### GTK_SHRINK = 2
the widget should shrink as and when possible.

### GTK_FILL = 4
the widget should fill the space allocated to it.


## `CalendarDisplayOptions`

These options can be used to influence the display and behaviour of a #GtkCalendar.

C - `GtkCalendarDisplayOptions`

### GTK_CALENDAR_SHOW_HEADING = 1
Specifies that the month and year should be displayed.

### GTK_CALENDAR_SHOW_DAY_NAMES = 2
Specifies that three letter day descriptions should be present.

### GTK_CALENDAR_NO_MONTH_CHANGE = 4
Prevents the user from switching months with the calendar.

### GTK_CALENDAR_SHOW_WEEK_NUMBERS = 8
Displays each week numbers of the current year, down the
left side of the calendar.

### GTK_CALENDAR_SHOW_DETAILS = 32
Just show an indicator, not the full details
text when details are provided. See gtk_calendar_set_detail_func().


## `CellRendererState`

Tells how a cell is to be rendered.

C - `GtkCellRendererState`

### GTK_CELL_RENDERER_SELECTED = 1
The cell is currently selected, and
 probably has a selection colored background to render to.

### GTK_CELL_RENDERER_PRELIT = 2
The mouse is hovering over the cell.

### GTK_CELL_RENDERER_INSENSITIVE = 4
The cell is drawn in an insensitive manner

### GTK_CELL_RENDERER_SORTED = 8
The cell is in a sorted row

### GTK_CELL_RENDERER_FOCUSED = 16
The cell is in the focus row.

### GTK_CELL_RENDERER_EXPANDABLE = 32
The cell is in a row that can be expanded. Since 3.4

### GTK_CELL_RENDERER_EXPANDED = 64
The cell is in a row that is expanded. Since 3.4


## `DebugFlag`



C - `GtkDebugFlag`

### GTK_DEBUG_MISC = 1


### GTK_DEBUG_PLUGSOCKET = 2


### GTK_DEBUG_TEXT = 4


### GTK_DEBUG_TREE = 8


### GTK_DEBUG_UPDATES = 16


### GTK_DEBUG_KEYBINDINGS = 32


### GTK_DEBUG_MULTIHEAD = 64


### GTK_DEBUG_MODULES = 128


### GTK_DEBUG_GEOMETRY = 256


### GTK_DEBUG_ICONTHEME = 512


### GTK_DEBUG_PRINTING = 1024


### GTK_DEBUG_BUILDER = 2048


### GTK_DEBUG_SIZE_REQUEST = 4096


### GTK_DEBUG_NO_CSS_CACHE = 8192


### GTK_DEBUG_BASELINES = 16384


### GTK_DEBUG_PIXEL_CACHE = 32768


### GTK_DEBUG_NO_PIXEL_CACHE = 65536


### GTK_DEBUG_INTERACTIVE = 131072


### GTK_DEBUG_TOUCHSCREEN = 262144


### GTK_DEBUG_ACTIONS = 524288


### GTK_DEBUG_RESIZE = 1048576


### GTK_DEBUG_LAYOUT = 2097152



## `DestDefaults`

The #GtkDestDefaults enumeration specifies the various
types of action that will be taken on behalf
of the user for a drag destination site.

C - `GtkDestDefaults`

### GTK_DEST_DEFAULT_MOTION = 1
If set for a widget, GTK+, during a drag over this
  widget will check if the drag matches this widget’s list of possible targets
  and actions.
  GTK+ will then call gdk_drag_status() as appropriate.

### GTK_DEST_DEFAULT_HIGHLIGHT = 2
If set for a widget, GTK+ will draw a highlight on
  this widget as long as a drag is over this widget and the widget drag format
  and action are acceptable.

### GTK_DEST_DEFAULT_DROP = 4
If set for a widget, when a drop occurs, GTK+ will
  will check if the drag matches this widget’s list of possible targets and
  actions. If so, GTK+ will call gtk_drag_get_data() on behalf of the widget.
  Whether or not the drop is successful, GTK+ will call gtk_drag_finish(). If
  the action was a move, then if the drag was successful, then %TRUE will be
  passed for the @delete parameter to gtk_drag_finish().

### GTK_DEST_DEFAULT_ALL = 7
If set, specifies that all default actions should
  be taken.


## `DialogFlags`

Flags used to influence dialog construction.

C - `GtkDialogFlags`

### GTK_DIALOG_MODAL = 1
Make the constructed dialog modal,
    see gtk_window_set_modal()

### GTK_DIALOG_DESTROY_WITH_PARENT = 2
Destroy the dialog when its
    parent is destroyed, see gtk_window_set_destroy_with_parent()

### GTK_DIALOG_USE_HEADER_BAR = 4
Create dialog with actions in header
    bar instead of action area. Since 3.12.


## `FileFilterFlags`

These flags indicate what parts of a #GtkFileFilterInfo struct
are filled or need to be filled.

C - `GtkFileFilterFlags`

### GTK_FILE_FILTER_FILENAME = 1
the filename of the file being tested

### GTK_FILE_FILTER_URI = 2
the URI for the file being tested

### GTK_FILE_FILTER_DISPLAY_NAME = 4
the string that will be used to
  display the file in the file chooser

### GTK_FILE_FILTER_MIME_TYPE = 8
the mime type of the file


## `IconLookupFlags`

Used to specify options for gtk_icon_theme_lookup_icon()

C - `GtkIconLookupFlags`

### GTK_ICON_LOOKUP_NO_SVG = 1
Never get SVG icons, even if gdk-pixbuf
  supports them. Cannot be used together with %GTK_ICON_LOOKUP_FORCE_SVG.

### GTK_ICON_LOOKUP_FORCE_SVG = 2
Get SVG icons, even if gdk-pixbuf
  doesn’t support them.
  Cannot be used together with %GTK_ICON_LOOKUP_NO_SVG.

### GTK_ICON_LOOKUP_USE_BUILTIN = 4
When passed to
  gtk_icon_theme_lookup_icon() includes builtin icons
  as well as files. For a builtin icon, gtk_icon_info_get_filename()
  is %NULL and you need to call gtk_icon_info_get_builtin_pixbuf().

### GTK_ICON_LOOKUP_GENERIC_FALLBACK = 8
Try to shorten icon name at '-'
  characters before looking at inherited themes. This flag is only
  supported in functions that take a single icon name. For more general
  fallback, see gtk_icon_theme_choose_icon(). Since 2.12.

### GTK_ICON_LOOKUP_FORCE_SIZE = 16
Always get the icon scaled to the
  requested size. Since 2.14.

### GTK_ICON_LOOKUP_FORCE_REGULAR = 32
Try to always load regular icons, even
  when symbolic icon names are given. Since 3.14.

### GTK_ICON_LOOKUP_FORCE_SYMBOLIC = 64
Try to always load symbolic icons, even
  when regular icon names are given. Since 3.14.

### GTK_ICON_LOOKUP_DIR_LTR = 128
Try to load a variant of the icon for left-to-right
  text direction. Since 3.14.

### GTK_ICON_LOOKUP_DIR_RTL = 256
Try to load a variant of the icon for right-to-left
  text direction. Since 3.14.


## `InputHints`

Describes hints that might be taken into account by input methods
or applications. Note that input methods may already tailor their
behaviour according to the #GtkInputPurpose of the entry.

Some common sense is expected when using these flags - mixing
@GTK_INPUT_HINT_LOWERCASE with any of the uppercase hints makes no sense.

This enumeration may be extended in the future; input methods should
ignore unknown values.

C - `GtkInputHints`

### GTK_INPUT_HINT_NONE = 0
No special behaviour suggested

### GTK_INPUT_HINT_SPELLCHECK = 1
Suggest checking for typos

### GTK_INPUT_HINT_NO_SPELLCHECK = 2
Suggest not checking for typos

### GTK_INPUT_HINT_WORD_COMPLETION = 4
Suggest word completion

### GTK_INPUT_HINT_LOWERCASE = 8
Suggest to convert all text to lowercase

### GTK_INPUT_HINT_UPPERCASE_CHARS = 16
Suggest to capitalize all text

### GTK_INPUT_HINT_UPPERCASE_WORDS = 32
Suggest to capitalize the first
    character of each word

### GTK_INPUT_HINT_UPPERCASE_SENTENCES = 64
Suggest to capitalize the
    first word of each sentence

### GTK_INPUT_HINT_INHIBIT_OSK = 128
Suggest to not show an onscreen keyboard
    (e.g for a calculator that already has all the keys).

### GTK_INPUT_HINT_VERTICAL_WRITING = 256
The text is vertical. Since 3.18

### GTK_INPUT_HINT_EMOJI = 512
Suggest offering Emoji support. Since 3.22.20

### GTK_INPUT_HINT_NO_EMOJI = 1024
Suggest not offering Emoji support. Since 3.22.20


## `JunctionSides`

Describes how a rendered element connects to adjacent elements.

C - `GtkJunctionSides`

### GTK_JUNCTION_NONE = 0
No junctions.

### GTK_JUNCTION_CORNER_TOPLEFT = 1
Element connects on the top-left corner.

### GTK_JUNCTION_CORNER_TOPRIGHT = 2
Element connects on the top-right corner.

### GTK_JUNCTION_CORNER_BOTTOMLEFT = 4
Element connects on the bottom-left corner.

### GTK_JUNCTION_CORNER_BOTTOMRIGHT = 8
Element connects on the bottom-right corner.

### GTK_JUNCTION_TOP = 3
Element connects on the top side.

### GTK_JUNCTION_BOTTOM = 12
Element connects on the bottom side.

### GTK_JUNCTION_LEFT = 5
Element connects on the left side.

### GTK_JUNCTION_RIGHT = 10
Element connects on the right side.


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

C - `GtkPlacesOpenFlags`

### GTK_PLACES_OPEN_NORMAL = 1
This is the default mode that #GtkPlacesSidebar uses if no other flags
 are specified.  It indicates that the calling application should open the selected location
 in the normal way, for example, in the folder view beside the sidebar.

### GTK_PLACES_OPEN_NEW_TAB = 2
When passed to gtk_places_sidebar_set_open_flags(), this indicates
 that the application can open folders selected from the sidebar in new tabs.  This value
 will be passed to the #GtkPlacesSidebar::open-location signal when the user selects
 that a location be opened in a new tab instead of in the standard fashion.

### GTK_PLACES_OPEN_NEW_WINDOW = 4
Similar to @GTK_PLACES_OPEN_NEW_TAB, but indicates that the application
 can open folders in new windows.


## `RcFlags`

Deprecated

C - `GtkRcFlags`

### GTK_RC_FG = 1
Deprecated

### GTK_RC_BG = 2
Deprecated

### GTK_RC_TEXT = 4
Deprecated

### GTK_RC_BASE = 8
Deprecated


## `RecentFilterFlags`

These flags indicate what parts of a #GtkRecentFilterInfo struct
are filled or need to be filled.

C - `GtkRecentFilterFlags`

### GTK_RECENT_FILTER_URI = 1
the URI of the file being tested

### GTK_RECENT_FILTER_DISPLAY_NAME = 2
the string that will be used to
 display the file in the recent chooser

### GTK_RECENT_FILTER_MIME_TYPE = 4
the mime type of the file

### GTK_RECENT_FILTER_APPLICATION = 8
the list of applications that have
 registered the file

### GTK_RECENT_FILTER_GROUP = 16
the groups to which the file belongs to

### GTK_RECENT_FILTER_AGE = 32
the number of days elapsed since the file
 has been registered


## `RegionFlags`

Describes a region within a widget.

C - `GtkRegionFlags`

### GTK_REGION_EVEN = 1
Region has an even number within a set.

### GTK_REGION_ODD = 2
Region has an odd number within a set.

### GTK_REGION_FIRST = 4
Region is the first one within a set.

### GTK_REGION_LAST = 8
Region is the last one within a set.

### GTK_REGION_ONLY = 16
Region is the only one within a set.

### GTK_REGION_SORTED = 32
Region is part of a sorted area.


## `StateFlags`

Describes a widget state. Widget states are used to match the widget
against CSS pseudo-classes. Note that GTK extends the regular CSS
classes and sometimes uses different names.

C - `GtkStateFlags`

### GTK_STATE_FLAG_NORMAL = 0
State during normal operation.

### GTK_STATE_FLAG_ACTIVE = 1
Widget is active.

### GTK_STATE_FLAG_PRELIGHT = 2
Widget has a mouse pointer over it.

### GTK_STATE_FLAG_SELECTED = 4
Widget is selected.

### GTK_STATE_FLAG_INSENSITIVE = 8
Widget is insensitive.

### GTK_STATE_FLAG_INCONSISTENT = 16
Widget is inconsistent.

### GTK_STATE_FLAG_FOCUSED = 32
Widget has the keyboard focus.

### GTK_STATE_FLAG_BACKDROP = 64
Widget is in a background toplevel window.

### GTK_STATE_FLAG_DIR_LTR = 128
Widget is in left-to-right text direction. Since 3.8

### GTK_STATE_FLAG_DIR_RTL = 256
Widget is in right-to-left text direction. Since 3.8

### GTK_STATE_FLAG_LINK = 512
Widget is a link. Since 3.12

### GTK_STATE_FLAG_VISITED = 1024
The location the widget points to has already been visited. Since 3.12

### GTK_STATE_FLAG_CHECKED = 2048
Widget is checked. Since 3.14

### GTK_STATE_FLAG_DROP_ACTIVE = 4096
Widget is highlighted as a drop target for DND. Since 3.20


## `StyleContextPrintFlags`

Flags that modify the behavior of gtk_style_context_to_string().
New values may be added to this enumeration.

C - `GtkStyleContextPrintFlags`

### GTK_STYLE_CONTEXT_PRINT_NONE = 0


### GTK_STYLE_CONTEXT_PRINT_RECURSE = 1
Print the entire tree of
    CSS nodes starting at the style context's node

### GTK_STYLE_CONTEXT_PRINT_SHOW_STYLE = 2
Show the values of the
    CSS properties for each node


## `TargetFlags`

The #GtkTargetFlags enumeration is used to specify
constraints on a #GtkTargetEntry.

C - `GtkTargetFlags`

### GTK_TARGET_SAME_APP = 1
If this is set, the target will only be selected
  for drags within a single application.

### GTK_TARGET_SAME_WIDGET = 2
If this is set, the target will only be selected
  for drags within a single widget.

### GTK_TARGET_OTHER_APP = 4
If this is set, the target will not be selected
  for drags within a single application.

### GTK_TARGET_OTHER_WIDGET = 8
If this is set, the target will not be selected
  for drags withing a single widget.


## `TextSearchFlags`

Flags affecting how a search is done.

If neither #GTK_TEXT_SEARCH_VISIBLE_ONLY nor #GTK_TEXT_SEARCH_TEXT_ONLY are
enabled, the match must be exact; the special 0xFFFC character will match
embedded pixbufs or child widgets.

C - `GtkTextSearchFlags`

### GTK_TEXT_SEARCH_VISIBLE_ONLY = 1
Search only visible data. A search match may
have invisible text interspersed.

### GTK_TEXT_SEARCH_TEXT_ONLY = 2
Search only text. A match may have pixbufs or
child widgets mixed inside the matched range.

### GTK_TEXT_SEARCH_CASE_INSENSITIVE = 4
The text will be matched regardless of
what case it is in.


## `ToolPaletteDragTargets`

Flags used to specify the supported drag targets.

C - `GtkToolPaletteDragTargets`

### GTK_TOOL_PALETTE_DRAG_ITEMS = 1
Support drag of items.

### GTK_TOOL_PALETTE_DRAG_GROUPS = 2
Support drag of groups.


## `TreeModelFlags`

These flags indicate various properties of a #GtkTreeModel.

They are returned by gtk_tree_model_get_flags(), and must be
static for the lifetime of the object. A more complete description
of #GTK_TREE_MODEL_ITERS_PERSIST can be found in the overview of
this section.

C - `GtkTreeModelFlags`

### GTK_TREE_MODEL_ITERS_PERSIST = 1
iterators survive all signals
    emitted by the tree

### GTK_TREE_MODEL_LIST_ONLY = 2
the model is a list only, and never
    has children


## `UIManagerItemType`

These enumeration values are used by gtk_ui_manager_add_ui() to determine
what UI element to create.

C - `GtkUIManagerItemType`

### GTK_UI_MANAGER_AUTO = 0
Pick the type of the UI element according to context.

### GTK_UI_MANAGER_MENUBAR = 1
Create a menubar.

### GTK_UI_MANAGER_MENU = 2
Create a menu.

### GTK_UI_MANAGER_TOOLBAR = 4
Create a toolbar.

### GTK_UI_MANAGER_PLACEHOLDER = 8
Insert a placeholder.

### GTK_UI_MANAGER_POPUP = 16
Create a popup menu.

### GTK_UI_MANAGER_MENUITEM = 32
Create a menuitem.

### GTK_UI_MANAGER_TOOLITEM = 64
Create a toolitem.

### GTK_UI_MANAGER_SEPARATOR = 128
Create a separator.

### GTK_UI_MANAGER_ACCELERATOR = 256
Install an accelerator.

### GTK_UI_MANAGER_POPUP_WITH_ACCELS = 512
Same as %GTK_UI_MANAGER_POPUP, but the
  actions’ accelerators are shown.


