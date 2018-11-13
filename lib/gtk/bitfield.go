// This is a generated file - DO NOT EDIT

package gtk

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

// Accelerator flags used with gtk_accel_group_connect().
type AccelFlags C.GtkAccelFlags

const (
	// Accelerator is visible
	GTK_ACCEL_VISIBLE AccelFlags = 1

	// Accelerator not removable
	GTK_ACCEL_LOCKED AccelFlags = 2

	// Mask
	GTK_ACCEL_MASK AccelFlags = 7
)

// Denotes the expansion properties that a widget will have when it (or its
// parent) is resized.
type AttachOptions C.GtkAttachOptions

const (
	// the widget should expand to take up any extra space in its
	// container that has been allocated.
	GTK_EXPAND AttachOptions = 1

	// the widget should shrink as and when possible.
	GTK_SHRINK AttachOptions = 2

	// the widget should fill the space allocated to it.
	GTK_FILL AttachOptions = 4
)

// These options can be used to influence the display and behaviour of a #GtkCalendar.
type CalendarDisplayOptions C.GtkCalendarDisplayOptions

const (
	// Specifies that the month and year should be displayed.
	GTK_CALENDAR_SHOW_HEADING CalendarDisplayOptions = 1

	// Specifies that three letter day descriptions should be present.
	GTK_CALENDAR_SHOW_DAY_NAMES CalendarDisplayOptions = 2

	// Prevents the user from switching months with the calendar.
	GTK_CALENDAR_NO_MONTH_CHANGE CalendarDisplayOptions = 4

	// Displays each week numbers of the current year, down the
	// left side of the calendar.
	GTK_CALENDAR_SHOW_WEEK_NUMBERS CalendarDisplayOptions = 8

	// Just show an indicator, not the full details
	// text when details are provided. See gtk_calendar_set_detail_func().
	GTK_CALENDAR_SHOW_DETAILS CalendarDisplayOptions = 32
)

// Tells how a cell is to be rendered.
type CellRendererState C.GtkCellRendererState

const (
	// The cell is currently selected, and
	// probably has a selection colored background to render to.
	GTK_CELL_RENDERER_SELECTED CellRendererState = 1

	// The mouse is hovering over the cell.
	GTK_CELL_RENDERER_PRELIT CellRendererState = 2

	// The cell is drawn in an insensitive manner
	GTK_CELL_RENDERER_INSENSITIVE CellRendererState = 4

	// The cell is in a sorted row
	GTK_CELL_RENDERER_SORTED CellRendererState = 8

	// The cell is in the focus row.
	GTK_CELL_RENDERER_FOCUSED CellRendererState = 16

	// The cell is in a row that can be expanded. Since 3.4
	GTK_CELL_RENDERER_EXPANDABLE CellRendererState = 32

	// The cell is in a row that is expanded. Since 3.4
	GTK_CELL_RENDERER_EXPANDED CellRendererState = 64
)

type DebugFlag C.GtkDebugFlag

const (
	GTK_DEBUG_MISC DebugFlag = 1

	GTK_DEBUG_PLUGSOCKET DebugFlag = 2

	GTK_DEBUG_TEXT DebugFlag = 4

	GTK_DEBUG_TREE DebugFlag = 8

	GTK_DEBUG_UPDATES DebugFlag = 16

	GTK_DEBUG_KEYBINDINGS DebugFlag = 32

	GTK_DEBUG_MULTIHEAD DebugFlag = 64

	GTK_DEBUG_MODULES DebugFlag = 128

	GTK_DEBUG_GEOMETRY DebugFlag = 256

	GTK_DEBUG_ICONTHEME DebugFlag = 512

	GTK_DEBUG_PRINTING DebugFlag = 1024

	GTK_DEBUG_BUILDER DebugFlag = 2048

	GTK_DEBUG_SIZE_REQUEST DebugFlag = 4096

	GTK_DEBUG_NO_CSS_CACHE DebugFlag = 8192

	GTK_DEBUG_BASELINES DebugFlag = 16384

	GTK_DEBUG_PIXEL_CACHE DebugFlag = 32768

	GTK_DEBUG_NO_PIXEL_CACHE DebugFlag = 65536

	GTK_DEBUG_INTERACTIVE DebugFlag = 131072

	GTK_DEBUG_TOUCHSCREEN DebugFlag = 262144

	GTK_DEBUG_ACTIONS DebugFlag = 524288

	GTK_DEBUG_RESIZE DebugFlag = 1048576

	GTK_DEBUG_LAYOUT DebugFlag = 2097152
)

// The #GtkDestDefaults enumeration specifies the various
// types of action that will be taken on behalf
// of the user for a drag destination site.
type DestDefaults C.GtkDestDefaults

const (
	// If set for a widget, GTK+, during a drag over this
	// widget will check if the drag matches this widget’s list of possible targets
	// and actions.
	// GTK+ will then call gdk_drag_status() as appropriate.
	GTK_DEST_DEFAULT_MOTION DestDefaults = 1

	// If set for a widget, GTK+ will draw a highlight on
	// this widget as long as a drag is over this widget and the widget drag format
	// and action are acceptable.
	GTK_DEST_DEFAULT_HIGHLIGHT DestDefaults = 2

	// If set for a widget, when a drop occurs, GTK+ will
	// will check if the drag matches this widget’s list of possible targets and
	// actions. If so, GTK+ will call gtk_drag_get_data() on behalf of the widget.
	// Whether or not the drop is successful, GTK+ will call gtk_drag_finish(). If
	// the action was a move, then if the drag was successful, then %TRUE will be
	// passed for the @delete parameter to gtk_drag_finish().
	GTK_DEST_DEFAULT_DROP DestDefaults = 4

	// If set, specifies that all default actions should
	// be taken.
	GTK_DEST_DEFAULT_ALL DestDefaults = 7
)

// Flags used to influence dialog construction.
type DialogFlags C.GtkDialogFlags

const (
	// Make the constructed dialog modal,
	// see gtk_window_set_modal()
	GTK_DIALOG_MODAL DialogFlags = 1

	// Destroy the dialog when its
	// parent is destroyed, see gtk_window_set_destroy_with_parent()
	GTK_DIALOG_DESTROY_WITH_PARENT DialogFlags = 2

	// Create dialog with actions in header
	// bar instead of action area. Since 3.12.
	GTK_DIALOG_USE_HEADER_BAR DialogFlags = 4
)

// These flags indicate what parts of a #GtkFileFilterInfo struct
// are filled or need to be filled.
type FileFilterFlags C.GtkFileFilterFlags

const (
	// the filename of the file being tested
	GTK_FILE_FILTER_FILENAME FileFilterFlags = 1

	// the URI for the file being tested
	GTK_FILE_FILTER_URI FileFilterFlags = 2

	// the string that will be used to
	// display the file in the file chooser
	GTK_FILE_FILTER_DISPLAY_NAME FileFilterFlags = 4

	// the mime type of the file
	GTK_FILE_FILTER_MIME_TYPE FileFilterFlags = 8
)

// Used to specify options for gtk_icon_theme_lookup_icon()
type IconLookupFlags C.GtkIconLookupFlags

const (
	// Never get SVG icons, even if gdk-pixbuf
	// supports them. Cannot be used together with %GTK_ICON_LOOKUP_FORCE_SVG.
	GTK_ICON_LOOKUP_NO_SVG IconLookupFlags = 1

	// Get SVG icons, even if gdk-pixbuf
	// doesn’t support them.
	// Cannot be used together with %GTK_ICON_LOOKUP_NO_SVG.
	GTK_ICON_LOOKUP_FORCE_SVG IconLookupFlags = 2

	// When passed to
	// gtk_icon_theme_lookup_icon() includes builtin icons
	// as well as files. For a builtin icon, gtk_icon_info_get_filename()
	// is %NULL and you need to call gtk_icon_info_get_builtin_pixbuf().
	GTK_ICON_LOOKUP_USE_BUILTIN IconLookupFlags = 4

	// Try to shorten icon name at '-'
	// characters before looking at inherited themes. This flag is only
	// supported in functions that take a single icon name. For more general
	// fallback, see gtk_icon_theme_choose_icon(). Since 2.12.
	GTK_ICON_LOOKUP_GENERIC_FALLBACK IconLookupFlags = 8

	// Always get the icon scaled to the
	// requested size. Since 2.14.
	GTK_ICON_LOOKUP_FORCE_SIZE IconLookupFlags = 16

	// Try to always load regular icons, even
	// when symbolic icon names are given. Since 3.14.
	GTK_ICON_LOOKUP_FORCE_REGULAR IconLookupFlags = 32

	// Try to always load symbolic icons, even
	// when regular icon names are given. Since 3.14.
	GTK_ICON_LOOKUP_FORCE_SYMBOLIC IconLookupFlags = 64

	// Try to load a variant of the icon for left-to-right
	// text direction. Since 3.14.
	GTK_ICON_LOOKUP_DIR_LTR IconLookupFlags = 128

	// Try to load a variant of the icon for right-to-left
	// text direction. Since 3.14.
	GTK_ICON_LOOKUP_DIR_RTL IconLookupFlags = 256
)

// Describes how a rendered element connects to adjacent elements.
type JunctionSides C.GtkJunctionSides

const (
	// No junctions.
	GTK_JUNCTION_NONE JunctionSides = 0

	// Element connects on the top-left corner.
	GTK_JUNCTION_CORNER_TOPLEFT JunctionSides = 1

	// Element connects on the top-right corner.
	GTK_JUNCTION_CORNER_TOPRIGHT JunctionSides = 2

	// Element connects on the bottom-left corner.
	GTK_JUNCTION_CORNER_BOTTOMLEFT JunctionSides = 4

	// Element connects on the bottom-right corner.
	GTK_JUNCTION_CORNER_BOTTOMRIGHT JunctionSides = 8

	// Element connects on the top side.
	GTK_JUNCTION_TOP JunctionSides = 3

	// Element connects on the bottom side.
	GTK_JUNCTION_BOTTOM JunctionSides = 12

	// Element connects on the left side.
	GTK_JUNCTION_LEFT JunctionSides = 5

	// Element connects on the right side.
	GTK_JUNCTION_RIGHT JunctionSides = 10
)

// These flags serve two purposes.  First, the application can call gtk_places_sidebar_set_open_flags()
// using these flags as a bitmask.  This tells the sidebar that the application is able to open
// folders selected from the sidebar in various ways, for example, in new tabs or in new windows in
// addition to the normal mode.
//
// Second, when one of these values gets passed back to the application in the
// #GtkPlacesSidebar::open-location signal, it means that the application should
// open the selected location in the normal way, in a new tab, or in a new
// window.  The sidebar takes care of determining the desired way to open the location,
// based on the modifier keys that the user is pressing at the time the selection is made.
//
// If the application never calls gtk_places_sidebar_set_open_flags(), then the sidebar will only
// use #GTK_PLACES_OPEN_NORMAL in the #GtkPlacesSidebar::open-location signal.  This is the
// default mode of operation.
type PlacesOpenFlags C.GtkPlacesOpenFlags

const (
	// This is the default mode that #GtkPlacesSidebar uses if no other flags
	// are specified.  It indicates that the calling application should open the selected location
	// in the normal way, for example, in the folder view beside the sidebar.
	GTK_PLACES_OPEN_NORMAL PlacesOpenFlags = 1

	// When passed to gtk_places_sidebar_set_open_flags(), this indicates
	// that the application can open folders selected from the sidebar in new tabs.  This value
	// will be passed to the #GtkPlacesSidebar::open-location signal when the user selects
	// that a location be opened in a new tab instead of in the standard fashion.
	GTK_PLACES_OPEN_NEW_TAB PlacesOpenFlags = 2

	// Similar to @GTK_PLACES_OPEN_NEW_TAB, but indicates that the application
	// can open folders in new windows.
	GTK_PLACES_OPEN_NEW_WINDOW PlacesOpenFlags = 4
)

// Deprecated
type RcFlags C.GtkRcFlags

const (
	// Deprecated
	GTK_RC_FG RcFlags = 1

	// Deprecated
	GTK_RC_BG RcFlags = 2

	// Deprecated
	GTK_RC_TEXT RcFlags = 4

	// Deprecated
	GTK_RC_BASE RcFlags = 8
)

// These flags indicate what parts of a #GtkRecentFilterInfo struct
// are filled or need to be filled.
type RecentFilterFlags C.GtkRecentFilterFlags

const (
	// the URI of the file being tested
	GTK_RECENT_FILTER_URI RecentFilterFlags = 1

	// the string that will be used to
	// display the file in the recent chooser
	GTK_RECENT_FILTER_DISPLAY_NAME RecentFilterFlags = 2

	// the mime type of the file
	GTK_RECENT_FILTER_MIME_TYPE RecentFilterFlags = 4

	// the list of applications that have
	// registered the file
	GTK_RECENT_FILTER_APPLICATION RecentFilterFlags = 8

	// the groups to which the file belongs to
	GTK_RECENT_FILTER_GROUP RecentFilterFlags = 16

	// the number of days elapsed since the file
	// has been registered
	GTK_RECENT_FILTER_AGE RecentFilterFlags = 32
)

// Describes a region within a widget.
type RegionFlags C.GtkRegionFlags

const (
	// Region has an even number within a set.
	GTK_REGION_EVEN RegionFlags = 1

	// Region has an odd number within a set.
	GTK_REGION_ODD RegionFlags = 2

	// Region is the first one within a set.
	GTK_REGION_FIRST RegionFlags = 4

	// Region is the last one within a set.
	GTK_REGION_LAST RegionFlags = 8

	// Region is the only one within a set.
	GTK_REGION_ONLY RegionFlags = 16

	// Region is part of a sorted area.
	GTK_REGION_SORTED RegionFlags = 32
)

// Describes a widget state. Widget states are used to match the widget
// against CSS pseudo-classes. Note that GTK extends the regular CSS
// classes and sometimes uses different names.
type StateFlags C.GtkStateFlags

const (
	// State during normal operation.
	GTK_STATE_FLAG_NORMAL StateFlags = 0

	// Widget is active.
	GTK_STATE_FLAG_ACTIVE StateFlags = 1

	// Widget has a mouse pointer over it.
	GTK_STATE_FLAG_PRELIGHT StateFlags = 2

	// Widget is selected.
	GTK_STATE_FLAG_SELECTED StateFlags = 4

	// Widget is insensitive.
	GTK_STATE_FLAG_INSENSITIVE StateFlags = 8

	// Widget is inconsistent.
	GTK_STATE_FLAG_INCONSISTENT StateFlags = 16

	// Widget has the keyboard focus.
	GTK_STATE_FLAG_FOCUSED StateFlags = 32

	// Widget is in a background toplevel window.
	GTK_STATE_FLAG_BACKDROP StateFlags = 64

	// Widget is in left-to-right text direction. Since 3.8
	GTK_STATE_FLAG_DIR_LTR StateFlags = 128

	// Widget is in right-to-left text direction. Since 3.8
	GTK_STATE_FLAG_DIR_RTL StateFlags = 256

	// Widget is a link. Since 3.12
	GTK_STATE_FLAG_LINK StateFlags = 512

	// The location the widget points to has already been visited. Since 3.12
	GTK_STATE_FLAG_VISITED StateFlags = 1024

	// Widget is checked. Since 3.14
	GTK_STATE_FLAG_CHECKED StateFlags = 2048

	// Widget is highlighted as a drop target for DND. Since 3.20
	GTK_STATE_FLAG_DROP_ACTIVE StateFlags = 4096
)

// The #GtkTargetFlags enumeration is used to specify
// constraints on a #GtkTargetEntry.
type TargetFlags C.GtkTargetFlags

const (
	// If this is set, the target will only be selected
	// for drags within a single application.
	GTK_TARGET_SAME_APP TargetFlags = 1

	// If this is set, the target will only be selected
	// for drags within a single widget.
	GTK_TARGET_SAME_WIDGET TargetFlags = 2

	// If this is set, the target will not be selected
	// for drags within a single application.
	GTK_TARGET_OTHER_APP TargetFlags = 4

	// If this is set, the target will not be selected
	// for drags withing a single widget.
	GTK_TARGET_OTHER_WIDGET TargetFlags = 8
)

// Flags affecting how a search is done.
//
// If neither #GTK_TEXT_SEARCH_VISIBLE_ONLY nor #GTK_TEXT_SEARCH_TEXT_ONLY are
// enabled, the match must be exact; the special 0xFFFC character will match
// embedded pixbufs or child widgets.
type TextSearchFlags C.GtkTextSearchFlags

const (
	// Search only visible data. A search match may
	// have invisible text interspersed.
	GTK_TEXT_SEARCH_VISIBLE_ONLY TextSearchFlags = 1

	// Search only text. A match may have pixbufs or
	// child widgets mixed inside the matched range.
	GTK_TEXT_SEARCH_TEXT_ONLY TextSearchFlags = 2

	// The text will be matched regardless of
	// what case it is in.
	GTK_TEXT_SEARCH_CASE_INSENSITIVE TextSearchFlags = 4
)

// Flags used to specify the supported drag targets.
type ToolPaletteDragTargets C.GtkToolPaletteDragTargets

const (
	// Support drag of items.
	GTK_TOOL_PALETTE_DRAG_ITEMS ToolPaletteDragTargets = 1

	// Support drag of groups.
	GTK_TOOL_PALETTE_DRAG_GROUPS ToolPaletteDragTargets = 2
)

// These flags indicate various properties of a #GtkTreeModel.
//
// They are returned by gtk_tree_model_get_flags(), and must be
// static for the lifetime of the object. A more complete description
// of #GTK_TREE_MODEL_ITERS_PERSIST can be found in the overview of
// this section.
type TreeModelFlags C.GtkTreeModelFlags

const (
	// iterators survive all signals
	// emitted by the tree
	GTK_TREE_MODEL_ITERS_PERSIST TreeModelFlags = 1

	// the model is a list only, and never
	// has children
	GTK_TREE_MODEL_LIST_ONLY TreeModelFlags = 2
)

// These enumeration values are used by gtk_ui_manager_add_ui() to determine
// what UI element to create.
type UIManagerItemType C.GtkUIManagerItemType

const (
	// Pick the type of the UI element according to context.
	GTK_UI_MANAGER_AUTO UIManagerItemType = 0

	// Create a menubar.
	GTK_UI_MANAGER_MENUBAR UIManagerItemType = 1

	// Create a menu.
	GTK_UI_MANAGER_MENU UIManagerItemType = 2

	// Create a toolbar.
	GTK_UI_MANAGER_TOOLBAR UIManagerItemType = 4

	// Insert a placeholder.
	GTK_UI_MANAGER_PLACEHOLDER UIManagerItemType = 8

	// Create a popup menu.
	GTK_UI_MANAGER_POPUP UIManagerItemType = 16

	// Create a menuitem.
	GTK_UI_MANAGER_MENUITEM UIManagerItemType = 32

	// Create a toolitem.
	GTK_UI_MANAGER_TOOLITEM UIManagerItemType = 64

	// Create a separator.
	GTK_UI_MANAGER_SEPARATOR UIManagerItemType = 128

	// Install an accelerator.
	GTK_UI_MANAGER_ACCELERATOR UIManagerItemType = 256

	// Same as %GTK_UI_MANAGER_POPUP, but the
	// actions’ accelerators are shown.
	GTK_UI_MANAGER_POPUP_WITH_ACCELS UIManagerItemType = 512
)
