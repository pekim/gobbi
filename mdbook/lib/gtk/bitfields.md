# `gtk` bitfields

## `AccelFlags`

Accelerator flags used with gtk_accel_group_connect().

C - `GtkAccelFlags`

## `ApplicationInhibitFlags`

Types of user actions that may be blocked by gtk_application_inhibit().

C - `GtkApplicationInhibitFlags`

## `AttachOptions`

Denotes the expansion properties that a widget will have when it (or its
parent) is resized.

C - `GtkAttachOptions`

## `CalendarDisplayOptions`

These options can be used to influence the display and behaviour of a #GtkCalendar.

C - `GtkCalendarDisplayOptions`

## `CellRendererState`

Tells how a cell is to be rendered.

C - `GtkCellRendererState`

## `DebugFlag`



C - `GtkDebugFlag`

## `DestDefaults`

The #GtkDestDefaults enumeration specifies the various
types of action that will be taken on behalf
of the user for a drag destination site.

C - `GtkDestDefaults`

## `DialogFlags`

Flags used to influence dialog construction.

C - `GtkDialogFlags`

## `FileFilterFlags`

These flags indicate what parts of a #GtkFileFilterInfo struct
are filled or need to be filled.

C - `GtkFileFilterFlags`

## `IconLookupFlags`

Used to specify options for gtk_icon_theme_lookup_icon()

C - `GtkIconLookupFlags`

## `InputHints`

Describes hints that might be taken into account by input methods
or applications. Note that input methods may already tailor their
behaviour according to the #GtkInputPurpose of the entry.

Some common sense is expected when using these flags - mixing
@GTK_INPUT_HINT_LOWERCASE with any of the uppercase hints makes no sense.

This enumeration may be extended in the future; input methods should
ignore unknown values.

C - `GtkInputHints`

## `JunctionSides`

Describes how a rendered element connects to adjacent elements.

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

C - `GtkPlacesOpenFlags`

## `RcFlags`

Deprecated

C - `GtkRcFlags`

## `RecentFilterFlags`

These flags indicate what parts of a #GtkRecentFilterInfo struct
are filled or need to be filled.

C - `GtkRecentFilterFlags`

## `RegionFlags`

Describes a region within a widget.

C - `GtkRegionFlags`

## `StateFlags`

Describes a widget state. Widget states are used to match the widget
against CSS pseudo-classes. Note that GTK extends the regular CSS
classes and sometimes uses different names.

C - `GtkStateFlags`

## `StyleContextPrintFlags`

Flags that modify the behavior of gtk_style_context_to_string().
New values may be added to this enumeration.

C - `GtkStyleContextPrintFlags`

## `TargetFlags`

The #GtkTargetFlags enumeration is used to specify
constraints on a #GtkTargetEntry.

C - `GtkTargetFlags`

## `TextSearchFlags`

Flags affecting how a search is done.

If neither #GTK_TEXT_SEARCH_VISIBLE_ONLY nor #GTK_TEXT_SEARCH_TEXT_ONLY are
enabled, the match must be exact; the special 0xFFFC character will match
embedded pixbufs or child widgets.

C - `GtkTextSearchFlags`

## `ToolPaletteDragTargets`

Flags used to specify the supported drag targets.

C - `GtkToolPaletteDragTargets`

## `TreeModelFlags`

These flags indicate various properties of a #GtkTreeModel.

They are returned by gtk_tree_model_get_flags(), and must be
static for the lifetime of the object. A more complete description
of #GTK_TREE_MODEL_ITERS_PERSIST can be found in the overview of
this section.

C - `GtkTreeModelFlags`

## `UIManagerItemType`

These enumeration values are used by gtk_ui_manager_add_ui() to determine
what UI element to create.

C - `GtkUIManagerItemType`

