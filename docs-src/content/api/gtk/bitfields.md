+++
title = "bitfields"
+++
<p class="api-heading">AccelFlags</p>
<p class="api-doc">Accelerator flags used with gtk_accel_group_connect().</p>
<table>
<tr>
<td class="name">GTK_ACCEL_VISIBLE</td>
<td class="value">1</td>
<td class="doc">Accelerator is visible</td>
</tr>
<tr>
<td class="name">GTK_ACCEL_LOCKED</td>
<td class="value">2</td>
<td class="doc">Accelerator not removable</td>
</tr>
<tr>
<td class="name">GTK_ACCEL_MASK</td>
<td class="value">7</td>
<td class="doc">Mask</td>
</tr>
</table>
<div class="api-notes">
  <p class="api-ctype">C type: GtkAccelFlags</p>
</div>
<p class="api-heading">ApplicationInhibitFlags</p>
<p class="api-doc">Types of user actions that may be blocked by gtk_application_inhibit().</p>
<table>
<tr>
<td class="name">GTK_APPLICATION_INHIBIT_LOGOUT</td>
<td class="value">1</td>
<td class="doc">Inhibit ending the user session
    by logging out or by shutting down the computer</td>
</tr>
<tr>
<td class="name">GTK_APPLICATION_INHIBIT_SWITCH</td>
<td class="value">2</td>
<td class="doc">Inhibit user switching</td>
</tr>
<tr>
<td class="name">GTK_APPLICATION_INHIBIT_SUSPEND</td>
<td class="value">4</td>
<td class="doc">Inhibit suspending the
    session or computer</td>
</tr>
<tr>
<td class="name">GTK_APPLICATION_INHIBIT_IDLE</td>
<td class="value">8</td>
<td class="doc">Inhibit the session being
    marked as idle (and possibly locked)</td>
</tr>
</table>
<div class="api-notes">
  <p class="api-since">since 3.4</p>
  <p class="api-ctype">C type: GtkApplicationInhibitFlags</p>
</div>
<p class="api-heading">AttachOptions</p>
<p class="api-doc">Denotes the expansion properties that a widget will have when it (or its
parent) is resized.</p>
<table>
<tr>
<td class="name">GTK_EXPAND</td>
<td class="value">1</td>
<td class="doc">the widget should expand to take up any extra space in its
container that has been allocated.</td>
</tr>
<tr>
<td class="name">GTK_SHRINK</td>
<td class="value">2</td>
<td class="doc">the widget should shrink as and when possible.</td>
</tr>
<tr>
<td class="name">GTK_FILL</td>
<td class="value">4</td>
<td class="doc">the widget should fill the space allocated to it.</td>
</tr>
</table>
<div class="api-notes">
  <p class="api-ctype">C type: GtkAttachOptions</p>
</div>
<p class="api-heading">CalendarDisplayOptions</p>
<p class="api-doc">These options can be used to influence the display and behaviour of a #GtkCalendar.</p>
<table>
<tr>
<td class="name">GTK_CALENDAR_SHOW_HEADING</td>
<td class="value">1</td>
<td class="doc">Specifies that the month and year should be displayed.</td>
</tr>
<tr>
<td class="name">GTK_CALENDAR_SHOW_DAY_NAMES</td>
<td class="value">2</td>
<td class="doc">Specifies that three letter day descriptions should be present.</td>
</tr>
<tr>
<td class="name">GTK_CALENDAR_NO_MONTH_CHANGE</td>
<td class="value">4</td>
<td class="doc">Prevents the user from switching months with the calendar.</td>
</tr>
<tr>
<td class="name">GTK_CALENDAR_SHOW_WEEK_NUMBERS</td>
<td class="value">8</td>
<td class="doc">Displays each week numbers of the current year, down the
left side of the calendar.</td>
</tr>
<tr>
<td class="name">GTK_CALENDAR_SHOW_DETAILS</td>
<td class="value">32</td>
<td class="doc">Just show an indicator, not the full details
text when details are provided. See gtk_calendar_set_detail_func().</td>
</tr>
</table>
<div class="api-notes">
  <p class="api-ctype">C type: GtkCalendarDisplayOptions</p>
</div>
<p class="api-heading">CellRendererState</p>
<p class="api-doc">Tells how a cell is to be rendered.</p>
<table>
<tr>
<td class="name">GTK_CELL_RENDERER_SELECTED</td>
<td class="value">1</td>
<td class="doc">The cell is currently selected, and
 probably has a selection colored background to render to.</td>
</tr>
<tr>
<td class="name">GTK_CELL_RENDERER_PRELIT</td>
<td class="value">2</td>
<td class="doc">The mouse is hovering over the cell.</td>
</tr>
<tr>
<td class="name">GTK_CELL_RENDERER_INSENSITIVE</td>
<td class="value">4</td>
<td class="doc">The cell is drawn in an insensitive manner</td>
</tr>
<tr>
<td class="name">GTK_CELL_RENDERER_SORTED</td>
<td class="value">8</td>
<td class="doc">The cell is in a sorted row</td>
</tr>
<tr>
<td class="name">GTK_CELL_RENDERER_FOCUSED</td>
<td class="value">16</td>
<td class="doc">The cell is in the focus row.</td>
</tr>
<tr>
<td class="name">GTK_CELL_RENDERER_EXPANDABLE</td>
<td class="value">32</td>
<td class="doc">The cell is in a row that can be expanded. Since 3.4</td>
</tr>
<tr>
<td class="name">GTK_CELL_RENDERER_EXPANDED</td>
<td class="value">64</td>
<td class="doc">The cell is in a row that is expanded. Since 3.4</td>
</tr>
</table>
<div class="api-notes">
  <p class="api-ctype">C type: GtkCellRendererState</p>
</div>
<p class="api-heading">DebugFlag</p>
<table>
<tr>
<td class="name">GTK_DEBUG_MISC</td>
<td class="value">1</td>
</tr>
<tr>
<td class="name">GTK_DEBUG_PLUGSOCKET</td>
<td class="value">2</td>
</tr>
<tr>
<td class="name">GTK_DEBUG_TEXT</td>
<td class="value">4</td>
</tr>
<tr>
<td class="name">GTK_DEBUG_TREE</td>
<td class="value">8</td>
</tr>
<tr>
<td class="name">GTK_DEBUG_UPDATES</td>
<td class="value">16</td>
</tr>
<tr>
<td class="name">GTK_DEBUG_KEYBINDINGS</td>
<td class="value">32</td>
</tr>
<tr>
<td class="name">GTK_DEBUG_MULTIHEAD</td>
<td class="value">64</td>
</tr>
<tr>
<td class="name">GTK_DEBUG_MODULES</td>
<td class="value">128</td>
</tr>
<tr>
<td class="name">GTK_DEBUG_GEOMETRY</td>
<td class="value">256</td>
</tr>
<tr>
<td class="name">GTK_DEBUG_ICONTHEME</td>
<td class="value">512</td>
</tr>
<tr>
<td class="name">GTK_DEBUG_PRINTING</td>
<td class="value">1024</td>
</tr>
<tr>
<td class="name">GTK_DEBUG_BUILDER</td>
<td class="value">2048</td>
</tr>
<tr>
<td class="name">GTK_DEBUG_SIZE_REQUEST</td>
<td class="value">4096</td>
</tr>
<tr>
<td class="name">GTK_DEBUG_NO_CSS_CACHE</td>
<td class="value">8192</td>
</tr>
<tr>
<td class="name">GTK_DEBUG_BASELINES</td>
<td class="value">16384</td>
</tr>
<tr>
<td class="name">GTK_DEBUG_PIXEL_CACHE</td>
<td class="value">32768</td>
</tr>
<tr>
<td class="name">GTK_DEBUG_NO_PIXEL_CACHE</td>
<td class="value">65536</td>
</tr>
<tr>
<td class="name">GTK_DEBUG_INTERACTIVE</td>
<td class="value">131072</td>
</tr>
<tr>
<td class="name">GTK_DEBUG_TOUCHSCREEN</td>
<td class="value">262144</td>
</tr>
<tr>
<td class="name">GTK_DEBUG_ACTIONS</td>
<td class="value">524288</td>
</tr>
<tr>
<td class="name">GTK_DEBUG_RESIZE</td>
<td class="value">1048576</td>
</tr>
<tr>
<td class="name">GTK_DEBUG_LAYOUT</td>
<td class="value">2097152</td>
</tr>
</table>
<div class="api-notes">
  <p class="api-ctype">C type: GtkDebugFlag</p>
</div>
<p class="api-heading">DestDefaults</p>
<p class="api-doc">The #GtkDestDefaults enumeration specifies the various
types of action that will be taken on behalf
of the user for a drag destination site.</p>
<table>
<tr>
<td class="name">GTK_DEST_DEFAULT_MOTION</td>
<td class="value">1</td>
<td class="doc">If set for a widget, GTK+, during a drag over this
  widget will check if the drag matches this widget’s list of possible targets
  and actions.
  GTK+ will then call gdk_drag_status() as appropriate.</td>
</tr>
<tr>
<td class="name">GTK_DEST_DEFAULT_HIGHLIGHT</td>
<td class="value">2</td>
<td class="doc">If set for a widget, GTK+ will draw a highlight on
  this widget as long as a drag is over this widget and the widget drag format
  and action are acceptable.</td>
</tr>
<tr>
<td class="name">GTK_DEST_DEFAULT_DROP</td>
<td class="value">4</td>
<td class="doc">If set for a widget, when a drop occurs, GTK+ will
  will check if the drag matches this widget’s list of possible targets and
  actions. If so, GTK+ will call gtk_drag_get_data() on behalf of the widget.
  Whether or not the drop is successful, GTK+ will call gtk_drag_finish(). If
  the action was a move, then if the drag was successful, then %TRUE will be
  passed for the @delete parameter to gtk_drag_finish().</td>
</tr>
<tr>
<td class="name">GTK_DEST_DEFAULT_ALL</td>
<td class="value">7</td>
<td class="doc">If set, specifies that all default actions should
  be taken.</td>
</tr>
</table>
<div class="api-notes">
  <p class="api-ctype">C type: GtkDestDefaults</p>
</div>
<p class="api-heading">DialogFlags</p>
<p class="api-doc">Flags used to influence dialog construction.</p>
<table>
<tr>
<td class="name">GTK_DIALOG_MODAL</td>
<td class="value">1</td>
<td class="doc">Make the constructed dialog modal,
    see gtk_window_set_modal()</td>
</tr>
<tr>
<td class="name">GTK_DIALOG_DESTROY_WITH_PARENT</td>
<td class="value">2</td>
<td class="doc">Destroy the dialog when its
    parent is destroyed, see gtk_window_set_destroy_with_parent()</td>
</tr>
<tr>
<td class="name">GTK_DIALOG_USE_HEADER_BAR</td>
<td class="value">4</td>
<td class="doc">Create dialog with actions in header
    bar instead of action area. Since 3.12.</td>
</tr>
</table>
<div class="api-notes">
  <p class="api-ctype">C type: GtkDialogFlags</p>
</div>
<p class="api-heading">FileFilterFlags</p>
<p class="api-doc">These flags indicate what parts of a #GtkFileFilterInfo struct
are filled or need to be filled.</p>
<table>
<tr>
<td class="name">GTK_FILE_FILTER_FILENAME</td>
<td class="value">1</td>
<td class="doc">the filename of the file being tested</td>
</tr>
<tr>
<td class="name">GTK_FILE_FILTER_URI</td>
<td class="value">2</td>
<td class="doc">the URI for the file being tested</td>
</tr>
<tr>
<td class="name">GTK_FILE_FILTER_DISPLAY_NAME</td>
<td class="value">4</td>
<td class="doc">the string that will be used to
  display the file in the file chooser</td>
</tr>
<tr>
<td class="name">GTK_FILE_FILTER_MIME_TYPE</td>
<td class="value">8</td>
<td class="doc">the mime type of the file</td>
</tr>
</table>
<div class="api-notes">
  <p class="api-ctype">C type: GtkFileFilterFlags</p>
</div>
<p class="api-heading">IconLookupFlags</p>
<p class="api-doc">Used to specify options for gtk_icon_theme_lookup_icon()</p>
<table>
<tr>
<td class="name">GTK_ICON_LOOKUP_NO_SVG</td>
<td class="value">1</td>
<td class="doc">Never get SVG icons, even if gdk-pixbuf
  supports them. Cannot be used together with %GTK_ICON_LOOKUP_FORCE_SVG.</td>
</tr>
<tr>
<td class="name">GTK_ICON_LOOKUP_FORCE_SVG</td>
<td class="value">2</td>
<td class="doc">Get SVG icons, even if gdk-pixbuf
  doesn’t support them.
  Cannot be used together with %GTK_ICON_LOOKUP_NO_SVG.</td>
</tr>
<tr>
<td class="name">GTK_ICON_LOOKUP_USE_BUILTIN</td>
<td class="value">4</td>
<td class="doc">When passed to
  gtk_icon_theme_lookup_icon() includes builtin icons
  as well as files. For a builtin icon, gtk_icon_info_get_filename()
  is %NULL and you need to call gtk_icon_info_get_builtin_pixbuf().</td>
</tr>
<tr>
<td class="name">GTK_ICON_LOOKUP_GENERIC_FALLBACK</td>
<td class="value">8</td>
<td class="doc">Try to shorten icon name at '-'
  characters before looking at inherited themes. This flag is only
  supported in functions that take a single icon name. For more general
  fallback, see gtk_icon_theme_choose_icon(). Since 2.12.</td>
</tr>
<tr>
<td class="name">GTK_ICON_LOOKUP_FORCE_SIZE</td>
<td class="value">16</td>
<td class="doc">Always get the icon scaled to the
  requested size. Since 2.14.</td>
</tr>
<tr>
<td class="name">GTK_ICON_LOOKUP_FORCE_REGULAR</td>
<td class="value">32</td>
<td class="doc">Try to always load regular icons, even
  when symbolic icon names are given. Since 3.14.</td>
</tr>
<tr>
<td class="name">GTK_ICON_LOOKUP_FORCE_SYMBOLIC</td>
<td class="value">64</td>
<td class="doc">Try to always load symbolic icons, even
  when regular icon names are given. Since 3.14.</td>
</tr>
<tr>
<td class="name">GTK_ICON_LOOKUP_DIR_LTR</td>
<td class="value">128</td>
<td class="doc">Try to load a variant of the icon for left-to-right
  text direction. Since 3.14.</td>
</tr>
<tr>
<td class="name">GTK_ICON_LOOKUP_DIR_RTL</td>
<td class="value">256</td>
<td class="doc">Try to load a variant of the icon for right-to-left
  text direction. Since 3.14.</td>
</tr>
</table>
<div class="api-notes">
  <p class="api-ctype">C type: GtkIconLookupFlags</p>
</div>
<p class="api-heading">InputHints</p>
<p class="api-doc">Describes hints that might be taken into account by input methods
or applications. Note that input methods may already tailor their
behaviour according to the #GtkInputPurpose of the entry.

Some common sense is expected when using these flags - mixing
@GTK_INPUT_HINT_LOWERCASE with any of the uppercase hints makes no sense.

This enumeration may be extended in the future; input methods should
ignore unknown values.</p>
<table>
<tr>
<td class="name">GTK_INPUT_HINT_NONE</td>
<td class="value">0</td>
<td class="doc">No special behaviour suggested</td>
</tr>
<tr>
<td class="name">GTK_INPUT_HINT_SPELLCHECK</td>
<td class="value">1</td>
<td class="doc">Suggest checking for typos</td>
</tr>
<tr>
<td class="name">GTK_INPUT_HINT_NO_SPELLCHECK</td>
<td class="value">2</td>
<td class="doc">Suggest not checking for typos</td>
</tr>
<tr>
<td class="name">GTK_INPUT_HINT_WORD_COMPLETION</td>
<td class="value">4</td>
<td class="doc">Suggest word completion</td>
</tr>
<tr>
<td class="name">GTK_INPUT_HINT_LOWERCASE</td>
<td class="value">8</td>
<td class="doc">Suggest to convert all text to lowercase</td>
</tr>
<tr>
<td class="name">GTK_INPUT_HINT_UPPERCASE_CHARS</td>
<td class="value">16</td>
<td class="doc">Suggest to capitalize all text</td>
</tr>
<tr>
<td class="name">GTK_INPUT_HINT_UPPERCASE_WORDS</td>
<td class="value">32</td>
<td class="doc">Suggest to capitalize the first
    character of each word</td>
</tr>
<tr>
<td class="name">GTK_INPUT_HINT_UPPERCASE_SENTENCES</td>
<td class="value">64</td>
<td class="doc">Suggest to capitalize the
    first word of each sentence</td>
</tr>
<tr>
<td class="name">GTK_INPUT_HINT_INHIBIT_OSK</td>
<td class="value">128</td>
<td class="doc">Suggest to not show an onscreen keyboard
    (e.g for a calculator that already has all the keys).</td>
</tr>
<tr>
<td class="name">GTK_INPUT_HINT_VERTICAL_WRITING</td>
<td class="value">256</td>
<td class="doc">The text is vertical. Since 3.18</td>
</tr>
<tr>
<td class="name">GTK_INPUT_HINT_EMOJI</td>
<td class="value">512</td>
<td class="doc">Suggest offering Emoji support. Since 3.22.20</td>
</tr>
<tr>
<td class="name">GTK_INPUT_HINT_NO_EMOJI</td>
<td class="value">1024</td>
<td class="doc">Suggest not offering Emoji support. Since 3.22.20</td>
</tr>
</table>
<div class="api-notes">
  <p class="api-since">since 3.6</p>
  <p class="api-ctype">C type: GtkInputHints</p>
</div>
<p class="api-heading">JunctionSides</p>
<p class="api-doc">Describes how a rendered element connects to adjacent elements.</p>
<table>
<tr>
<td class="name">GTK_JUNCTION_NONE</td>
<td class="value">0</td>
<td class="doc">No junctions.</td>
</tr>
<tr>
<td class="name">GTK_JUNCTION_CORNER_TOPLEFT</td>
<td class="value">1</td>
<td class="doc">Element connects on the top-left corner.</td>
</tr>
<tr>
<td class="name">GTK_JUNCTION_CORNER_TOPRIGHT</td>
<td class="value">2</td>
<td class="doc">Element connects on the top-right corner.</td>
</tr>
<tr>
<td class="name">GTK_JUNCTION_CORNER_BOTTOMLEFT</td>
<td class="value">4</td>
<td class="doc">Element connects on the bottom-left corner.</td>
</tr>
<tr>
<td class="name">GTK_JUNCTION_CORNER_BOTTOMRIGHT</td>
<td class="value">8</td>
<td class="doc">Element connects on the bottom-right corner.</td>
</tr>
<tr>
<td class="name">GTK_JUNCTION_TOP</td>
<td class="value">3</td>
<td class="doc">Element connects on the top side.</td>
</tr>
<tr>
<td class="name">GTK_JUNCTION_BOTTOM</td>
<td class="value">12</td>
<td class="doc">Element connects on the bottom side.</td>
</tr>
<tr>
<td class="name">GTK_JUNCTION_LEFT</td>
<td class="value">5</td>
<td class="doc">Element connects on the left side.</td>
</tr>
<tr>
<td class="name">GTK_JUNCTION_RIGHT</td>
<td class="value">10</td>
<td class="doc">Element connects on the right side.</td>
</tr>
</table>
<div class="api-notes">
  <p class="api-ctype">C type: GtkJunctionSides</p>
</div>
<p class="api-heading">PlacesOpenFlags</p>
<p class="api-doc">These flags serve two purposes.  First, the application can call gtk_places_sidebar_set_open_flags()
using these flags as a bitmask.  This tells the sidebar that the application is able to open
folders selected from the sidebar in various ways, for example, in new tabs or in new windows in
addition to the normal mode.

Second, when one of these values gets passed back to the application in the
#GtkPlacesSidebar::open-location signal, it means that the application should
open the selected location in the normal way, in a new tab, or in a new
window.  The sidebar takes care of determining the desired way to open the location,
based on the modifier keys that the user is pressing at the time the selection is made.

If the application never calls gtk_places_sidebar_set_open_flags(), then the sidebar will only
use #GTK_PLACES_OPEN_NORMAL in the #GtkPlacesSidebar::open-location signal.  This is the
default mode of operation.</p>
<table>
<tr>
<td class="name">GTK_PLACES_OPEN_NORMAL</td>
<td class="value">1</td>
<td class="doc">This is the default mode that #GtkPlacesSidebar uses if no other flags
 are specified.  It indicates that the calling application should open the selected location
 in the normal way, for example, in the folder view beside the sidebar.</td>
</tr>
<tr>
<td class="name">GTK_PLACES_OPEN_NEW_TAB</td>
<td class="value">2</td>
<td class="doc">When passed to gtk_places_sidebar_set_open_flags(), this indicates
 that the application can open folders selected from the sidebar in new tabs.  This value
 will be passed to the #GtkPlacesSidebar::open-location signal when the user selects
 that a location be opened in a new tab instead of in the standard fashion.</td>
</tr>
<tr>
<td class="name">GTK_PLACES_OPEN_NEW_WINDOW</td>
<td class="value">4</td>
<td class="doc">Similar to @GTK_PLACES_OPEN_NEW_TAB, but indicates that the application
 can open folders in new windows.</td>
</tr>
</table>
<div class="api-notes">
  <p class="api-ctype">C type: GtkPlacesOpenFlags</p>
</div>
<p class="api-heading">RcFlags</p>
<p class="api-doc">Deprecated</p>
<table>
<tr>
<td class="name">GTK_RC_FG</td>
<td class="value">1</td>
<td class="doc">Deprecated</td>
</tr>
<tr>
<td class="name">GTK_RC_BG</td>
<td class="value">2</td>
<td class="doc">Deprecated</td>
</tr>
<tr>
<td class="name">GTK_RC_TEXT</td>
<td class="value">4</td>
<td class="doc">Deprecated</td>
</tr>
<tr>
<td class="name">GTK_RC_BASE</td>
<td class="value">8</td>
<td class="doc">Deprecated</td>
</tr>
</table>
<div class="api-notes">
  <p class="api-ctype">C type: GtkRcFlags</p>
</div>
<p class="api-heading">RecentFilterFlags</p>
<p class="api-doc">These flags indicate what parts of a #GtkRecentFilterInfo struct
are filled or need to be filled.</p>
<table>
<tr>
<td class="name">GTK_RECENT_FILTER_URI</td>
<td class="value">1</td>
<td class="doc">the URI of the file being tested</td>
</tr>
<tr>
<td class="name">GTK_RECENT_FILTER_DISPLAY_NAME</td>
<td class="value">2</td>
<td class="doc">the string that will be used to
 display the file in the recent chooser</td>
</tr>
<tr>
<td class="name">GTK_RECENT_FILTER_MIME_TYPE</td>
<td class="value">4</td>
<td class="doc">the mime type of the file</td>
</tr>
<tr>
<td class="name">GTK_RECENT_FILTER_APPLICATION</td>
<td class="value">8</td>
<td class="doc">the list of applications that have
 registered the file</td>
</tr>
<tr>
<td class="name">GTK_RECENT_FILTER_GROUP</td>
<td class="value">16</td>
<td class="doc">the groups to which the file belongs to</td>
</tr>
<tr>
<td class="name">GTK_RECENT_FILTER_AGE</td>
<td class="value">32</td>
<td class="doc">the number of days elapsed since the file
 has been registered</td>
</tr>
</table>
<div class="api-notes">
  <p class="api-ctype">C type: GtkRecentFilterFlags</p>
</div>
<p class="api-heading">RegionFlags</p>
<p class="api-doc">Describes a region within a widget.</p>
<table>
<tr>
<td class="name">GTK_REGION_EVEN</td>
<td class="value">1</td>
<td class="doc">Region has an even number within a set.</td>
</tr>
<tr>
<td class="name">GTK_REGION_ODD</td>
<td class="value">2</td>
<td class="doc">Region has an odd number within a set.</td>
</tr>
<tr>
<td class="name">GTK_REGION_FIRST</td>
<td class="value">4</td>
<td class="doc">Region is the first one within a set.</td>
</tr>
<tr>
<td class="name">GTK_REGION_LAST</td>
<td class="value">8</td>
<td class="doc">Region is the last one within a set.</td>
</tr>
<tr>
<td class="name">GTK_REGION_ONLY</td>
<td class="value">16</td>
<td class="doc">Region is the only one within a set.</td>
</tr>
<tr>
<td class="name">GTK_REGION_SORTED</td>
<td class="value">32</td>
<td class="doc">Region is part of a sorted area.</td>
</tr>
</table>
<div class="api-notes">
  <p class="api-ctype">C type: GtkRegionFlags</p>
</div>
<p class="api-heading">StateFlags</p>
<p class="api-doc">Describes a widget state. Widget states are used to match the widget
against CSS pseudo-classes. Note that GTK extends the regular CSS
classes and sometimes uses different names.</p>
<table>
<tr>
<td class="name">GTK_STATE_FLAG_NORMAL</td>
<td class="value">0</td>
<td class="doc">State during normal operation.</td>
</tr>
<tr>
<td class="name">GTK_STATE_FLAG_ACTIVE</td>
<td class="value">1</td>
<td class="doc">Widget is active.</td>
</tr>
<tr>
<td class="name">GTK_STATE_FLAG_PRELIGHT</td>
<td class="value">2</td>
<td class="doc">Widget has a mouse pointer over it.</td>
</tr>
<tr>
<td class="name">GTK_STATE_FLAG_SELECTED</td>
<td class="value">4</td>
<td class="doc">Widget is selected.</td>
</tr>
<tr>
<td class="name">GTK_STATE_FLAG_INSENSITIVE</td>
<td class="value">8</td>
<td class="doc">Widget is insensitive.</td>
</tr>
<tr>
<td class="name">GTK_STATE_FLAG_INCONSISTENT</td>
<td class="value">16</td>
<td class="doc">Widget is inconsistent.</td>
</tr>
<tr>
<td class="name">GTK_STATE_FLAG_FOCUSED</td>
<td class="value">32</td>
<td class="doc">Widget has the keyboard focus.</td>
</tr>
<tr>
<td class="name">GTK_STATE_FLAG_BACKDROP</td>
<td class="value">64</td>
<td class="doc">Widget is in a background toplevel window.</td>
</tr>
<tr>
<td class="name">GTK_STATE_FLAG_DIR_LTR</td>
<td class="value">128</td>
<td class="doc">Widget is in left-to-right text direction. Since 3.8</td>
</tr>
<tr>
<td class="name">GTK_STATE_FLAG_DIR_RTL</td>
<td class="value">256</td>
<td class="doc">Widget is in right-to-left text direction. Since 3.8</td>
</tr>
<tr>
<td class="name">GTK_STATE_FLAG_LINK</td>
<td class="value">512</td>
<td class="doc">Widget is a link. Since 3.12</td>
</tr>
<tr>
<td class="name">GTK_STATE_FLAG_VISITED</td>
<td class="value">1024</td>
<td class="doc">The location the widget points to has already been visited. Since 3.12</td>
</tr>
<tr>
<td class="name">GTK_STATE_FLAG_CHECKED</td>
<td class="value">2048</td>
<td class="doc">Widget is checked. Since 3.14</td>
</tr>
<tr>
<td class="name">GTK_STATE_FLAG_DROP_ACTIVE</td>
<td class="value">4096</td>
<td class="doc">Widget is highlighted as a drop target for DND. Since 3.20</td>
</tr>
</table>
<div class="api-notes">
  <p class="api-ctype">C type: GtkStateFlags</p>
</div>
<p class="api-heading">StyleContextPrintFlags</p>
<p class="api-doc">Flags that modify the behavior of gtk_style_context_to_string().
New values may be added to this enumeration.</p>
<table>
<tr>
<td class="name">GTK_STYLE_CONTEXT_PRINT_NONE</td>
<td class="value">0</td>
</tr>
<tr>
<td class="name">GTK_STYLE_CONTEXT_PRINT_RECURSE</td>
<td class="value">1</td>
<td class="doc">Print the entire tree of
    CSS nodes starting at the style context's node</td>
</tr>
<tr>
<td class="name">GTK_STYLE_CONTEXT_PRINT_SHOW_STYLE</td>
<td class="value">2</td>
<td class="doc">Show the values of the
    CSS properties for each node</td>
</tr>
</table>
<div class="api-notes">
  <p class="api-since">since 3.20</p>
  <p class="api-ctype">C type: GtkStyleContextPrintFlags</p>
</div>
<p class="api-heading">TargetFlags</p>
<p class="api-doc">The #GtkTargetFlags enumeration is used to specify
constraints on a #GtkTargetEntry.</p>
<table>
<tr>
<td class="name">GTK_TARGET_SAME_APP</td>
<td class="value">1</td>
<td class="doc">If this is set, the target will only be selected
  for drags within a single application.</td>
</tr>
<tr>
<td class="name">GTK_TARGET_SAME_WIDGET</td>
<td class="value">2</td>
<td class="doc">If this is set, the target will only be selected
  for drags within a single widget.</td>
</tr>
<tr>
<td class="name">GTK_TARGET_OTHER_APP</td>
<td class="value">4</td>
<td class="doc">If this is set, the target will not be selected
  for drags within a single application.</td>
</tr>
<tr>
<td class="name">GTK_TARGET_OTHER_WIDGET</td>
<td class="value">8</td>
<td class="doc">If this is set, the target will not be selected
  for drags withing a single widget.</td>
</tr>
</table>
<div class="api-notes">
  <p class="api-ctype">C type: GtkTargetFlags</p>
</div>
<p class="api-heading">TextSearchFlags</p>
<p class="api-doc">Flags affecting how a search is done.

If neither #GTK_TEXT_SEARCH_VISIBLE_ONLY nor #GTK_TEXT_SEARCH_TEXT_ONLY are
enabled, the match must be exact; the special 0xFFFC character will match
embedded pixbufs or child widgets.</p>
<table>
<tr>
<td class="name">GTK_TEXT_SEARCH_VISIBLE_ONLY</td>
<td class="value">1</td>
<td class="doc">Search only visible data. A search match may
have invisible text interspersed.</td>
</tr>
<tr>
<td class="name">GTK_TEXT_SEARCH_TEXT_ONLY</td>
<td class="value">2</td>
<td class="doc">Search only text. A match may have pixbufs or
child widgets mixed inside the matched range.</td>
</tr>
<tr>
<td class="name">GTK_TEXT_SEARCH_CASE_INSENSITIVE</td>
<td class="value">4</td>
<td class="doc">The text will be matched regardless of
what case it is in.</td>
</tr>
</table>
<div class="api-notes">
  <p class="api-ctype">C type: GtkTextSearchFlags</p>
</div>
<p class="api-heading">ToolPaletteDragTargets</p>
<p class="api-doc">Flags used to specify the supported drag targets.</p>
<table>
<tr>
<td class="name">GTK_TOOL_PALETTE_DRAG_ITEMS</td>
<td class="value">1</td>
<td class="doc">Support drag of items.</td>
</tr>
<tr>
<td class="name">GTK_TOOL_PALETTE_DRAG_GROUPS</td>
<td class="value">2</td>
<td class="doc">Support drag of groups.</td>
</tr>
</table>
<div class="api-notes">
  <p class="api-ctype">C type: GtkToolPaletteDragTargets</p>
</div>
<p class="api-heading">TreeModelFlags</p>
<p class="api-doc">These flags indicate various properties of a #GtkTreeModel.

They are returned by gtk_tree_model_get_flags(), and must be
static for the lifetime of the object. A more complete description
of #GTK_TREE_MODEL_ITERS_PERSIST can be found in the overview of
this section.</p>
<table>
<tr>
<td class="name">GTK_TREE_MODEL_ITERS_PERSIST</td>
<td class="value">1</td>
<td class="doc">iterators survive all signals
    emitted by the tree</td>
</tr>
<tr>
<td class="name">GTK_TREE_MODEL_LIST_ONLY</td>
<td class="value">2</td>
<td class="doc">the model is a list only, and never
    has children</td>
</tr>
</table>
<div class="api-notes">
  <p class="api-ctype">C type: GtkTreeModelFlags</p>
</div>
<p class="api-heading">UIManagerItemType</p>
<p class="api-doc">These enumeration values are used by gtk_ui_manager_add_ui() to determine
what UI element to create.</p>
<table>
<tr>
<td class="name">GTK_UI_MANAGER_AUTO</td>
<td class="value">0</td>
<td class="doc">Pick the type of the UI element according to context.</td>
</tr>
<tr>
<td class="name">GTK_UI_MANAGER_MENUBAR</td>
<td class="value">1</td>
<td class="doc">Create a menubar.</td>
</tr>
<tr>
<td class="name">GTK_UI_MANAGER_MENU</td>
<td class="value">2</td>
<td class="doc">Create a menu.</td>
</tr>
<tr>
<td class="name">GTK_UI_MANAGER_TOOLBAR</td>
<td class="value">4</td>
<td class="doc">Create a toolbar.</td>
</tr>
<tr>
<td class="name">GTK_UI_MANAGER_PLACEHOLDER</td>
<td class="value">8</td>
<td class="doc">Insert a placeholder.</td>
</tr>
<tr>
<td class="name">GTK_UI_MANAGER_POPUP</td>
<td class="value">16</td>
<td class="doc">Create a popup menu.</td>
</tr>
<tr>
<td class="name">GTK_UI_MANAGER_MENUITEM</td>
<td class="value">32</td>
<td class="doc">Create a menuitem.</td>
</tr>
<tr>
<td class="name">GTK_UI_MANAGER_TOOLITEM</td>
<td class="value">64</td>
<td class="doc">Create a toolitem.</td>
</tr>
<tr>
<td class="name">GTK_UI_MANAGER_SEPARATOR</td>
<td class="value">128</td>
<td class="doc">Create a separator.</td>
</tr>
<tr>
<td class="name">GTK_UI_MANAGER_ACCELERATOR</td>
<td class="value">256</td>
<td class="doc">Install an accelerator.</td>
</tr>
<tr>
<td class="name">GTK_UI_MANAGER_POPUP_WITH_ACCELS</td>
<td class="value">512</td>
<td class="doc">Same as %GTK_UI_MANAGER_POPUP, but the
  actions’ accelerators are shown.</td>
</tr>
</table>
<div class="api-notes">
  <p class="api-ctype">C type: GtkUIManagerItemType</p>
</div>
