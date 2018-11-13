// This is a generated file - DO NOT EDIT

package gtk

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

// Like gtk_get_binary_age(), but from the headers used at
// application compile time, rather than from the library linked
// against at application run time.
const BINARY_AGE int = C.GTK_BINARY_AGE

// Constant to return from a signal handler for the #GtkSpinButton::input
// signal in case of conversion failure.
const INPUT_ERROR int = C.GTK_INPUT_ERROR

// Like gtk_get_interface_age(), but from the headers used at
// application compile time, rather than from the library linked
// against at application run time.
const INTERFACE_AGE int = C.GTK_INTERFACE_AGE

// Like gtk_get_major_version(), but from the headers used at
// application compile time, rather than from the library linked
// against at application run time.
const MAJOR_VERSION int = C.GTK_MAJOR_VERSION

// The maximum length of sequences in compose tables.
const MAX_COMPOSE_LEN int = C.GTK_MAX_COMPOSE_LEN

// Like gtk_get_micro_version(), but from the headers used at
// application compile time, rather than from the library linked
// against at application run time.
const MICRO_VERSION int = C.GTK_MICRO_VERSION

// Like gtk_get_minor_version(), but from the headers used at
// application compile time, rather than from the library linked
// against at application run time.
const MINOR_VERSION int = C.GTK_MINOR_VERSION

// Name for the A3 paper size.
const PAPER_NAME_A3 string = C.GTK_PAPER_NAME_A3

// Name for the A4 paper size.
const PAPER_NAME_A4 string = C.GTK_PAPER_NAME_A4

// Name for the A5 paper size.
const PAPER_NAME_A5 string = C.GTK_PAPER_NAME_A5

// Name for the B5 paper size.
const PAPER_NAME_B5 string = C.GTK_PAPER_NAME_B5

// Name for the Executive paper size.
const PAPER_NAME_EXECUTIVE string = C.GTK_PAPER_NAME_EXECUTIVE

// Name for the Legal paper size.
const PAPER_NAME_LEGAL string = C.GTK_PAPER_NAME_LEGAL

// Name for the Letter paper size.
const PAPER_NAME_LETTER string = C.GTK_PAPER_NAME_LETTER
const PATH_PRIO_MASK int = C.GTK_PATH_PRIO_MASK
const PRINT_SETTINGS_COLLATE string = C.GTK_PRINT_SETTINGS_COLLATE
const PRINT_SETTINGS_DEFAULT_SOURCE string = C.GTK_PRINT_SETTINGS_DEFAULT_SOURCE
const PRINT_SETTINGS_DITHER string = C.GTK_PRINT_SETTINGS_DITHER
const PRINT_SETTINGS_DUPLEX string = C.GTK_PRINT_SETTINGS_DUPLEX
const PRINT_SETTINGS_FINISHINGS string = C.GTK_PRINT_SETTINGS_FINISHINGS
const PRINT_SETTINGS_MEDIA_TYPE string = C.GTK_PRINT_SETTINGS_MEDIA_TYPE
const PRINT_SETTINGS_NUMBER_UP string = C.GTK_PRINT_SETTINGS_NUMBER_UP
const PRINT_SETTINGS_NUMBER_UP_LAYOUT string = C.GTK_PRINT_SETTINGS_NUMBER_UP_LAYOUT
const PRINT_SETTINGS_N_COPIES string = C.GTK_PRINT_SETTINGS_N_COPIES
const PRINT_SETTINGS_ORIENTATION string = C.GTK_PRINT_SETTINGS_ORIENTATION
const PRINT_SETTINGS_OUTPUT_BIN string = C.GTK_PRINT_SETTINGS_OUTPUT_BIN

// The key used by the “Print to file” printer to store the format
// of the output. The supported values are “PS” and “PDF”.
const PRINT_SETTINGS_OUTPUT_FILE_FORMAT string = C.GTK_PRINT_SETTINGS_OUTPUT_FILE_FORMAT

// The key used by the “Print to file” printer to store the URI
// to which the output should be written. GTK+ itself supports
// only “file://” URIs.
const PRINT_SETTINGS_OUTPUT_URI string = C.GTK_PRINT_SETTINGS_OUTPUT_URI
const PRINT_SETTINGS_PAGE_RANGES string = C.GTK_PRINT_SETTINGS_PAGE_RANGES
const PRINT_SETTINGS_PAGE_SET string = C.GTK_PRINT_SETTINGS_PAGE_SET
const PRINT_SETTINGS_PAPER_FORMAT string = C.GTK_PRINT_SETTINGS_PAPER_FORMAT
const PRINT_SETTINGS_PAPER_HEIGHT string = C.GTK_PRINT_SETTINGS_PAPER_HEIGHT
const PRINT_SETTINGS_PAPER_WIDTH string = C.GTK_PRINT_SETTINGS_PAPER_WIDTH
const PRINT_SETTINGS_PRINTER string = C.GTK_PRINT_SETTINGS_PRINTER
const PRINT_SETTINGS_PRINTER_LPI string = C.GTK_PRINT_SETTINGS_PRINTER_LPI
const PRINT_SETTINGS_PRINT_PAGES string = C.GTK_PRINT_SETTINGS_PRINT_PAGES
const PRINT_SETTINGS_QUALITY string = C.GTK_PRINT_SETTINGS_QUALITY
const PRINT_SETTINGS_RESOLUTION string = C.GTK_PRINT_SETTINGS_RESOLUTION
const PRINT_SETTINGS_RESOLUTION_X string = C.GTK_PRINT_SETTINGS_RESOLUTION_X
const PRINT_SETTINGS_RESOLUTION_Y string = C.GTK_PRINT_SETTINGS_RESOLUTION_Y
const PRINT_SETTINGS_REVERSE string = C.GTK_PRINT_SETTINGS_REVERSE
const PRINT_SETTINGS_SCALE string = C.GTK_PRINT_SETTINGS_SCALE
const PRINT_SETTINGS_USE_COLOR string = C.GTK_PRINT_SETTINGS_USE_COLOR
const PRINT_SETTINGS_WIN32_DRIVER_EXTRA string = C.GTK_PRINT_SETTINGS_WIN32_DRIVER_EXTRA
const PRINT_SETTINGS_WIN32_DRIVER_VERSION string = C.GTK_PRINT_SETTINGS_WIN32_DRIVER_VERSION

// Use this priority for functionality related to size allocation.
//
// It is used internally by GTK+ to compute the sizes of widgets.
// This priority is higher than %GDK_PRIORITY_REDRAW to avoid
// resizing a widget which was just redrawn.
const PRIORITY_RESIZE int = C.GTK_PRIORITY_RESIZE

// Blacklisted : STOCK_ADD

// Blacklisted : STOCK_APPLY

// Blacklisted : STOCK_BOLD

// Blacklisted : STOCK_CANCEL

// Blacklisted : STOCK_CDROM

// Blacklisted : STOCK_CLEAR

// Blacklisted : STOCK_CLOSE

// Blacklisted : STOCK_CONVERT

// Blacklisted : STOCK_COPY

// Blacklisted : STOCK_CUT

// Blacklisted : STOCK_DELETE

// Blacklisted : STOCK_DIALOG_ERROR

// Blacklisted : STOCK_DIALOG_INFO

// Blacklisted : STOCK_DIALOG_QUESTION

// Blacklisted : STOCK_DIALOG_WARNING

// Blacklisted : STOCK_DND

// Blacklisted : STOCK_DND_MULTIPLE

// Blacklisted : STOCK_EXECUTE

// Blacklisted : STOCK_FIND

// Blacklisted : STOCK_FIND_AND_REPLACE

// Blacklisted : STOCK_FLOPPY

// Blacklisted : STOCK_GOTO_BOTTOM

// Blacklisted : STOCK_GOTO_FIRST

// Blacklisted : STOCK_GOTO_LAST

// Blacklisted : STOCK_GOTO_TOP

// Blacklisted : STOCK_GO_BACK

// Blacklisted : STOCK_GO_DOWN

// Blacklisted : STOCK_GO_FORWARD

// Blacklisted : STOCK_GO_UP

// Blacklisted : STOCK_HELP

// Blacklisted : STOCK_HOME

// Blacklisted : STOCK_INDEX

// Blacklisted : STOCK_ITALIC

// Blacklisted : STOCK_JUMP_TO

// Blacklisted : STOCK_JUSTIFY_CENTER

// Blacklisted : STOCK_JUSTIFY_FILL

// Blacklisted : STOCK_JUSTIFY_LEFT

// Blacklisted : STOCK_JUSTIFY_RIGHT

// Blacklisted : STOCK_MISSING_IMAGE

// Blacklisted : STOCK_NEW

// Blacklisted : STOCK_NO

// Blacklisted : STOCK_OK

// Blacklisted : STOCK_OPEN

// Blacklisted : STOCK_PASTE

// Blacklisted : STOCK_PREFERENCES

// Blacklisted : STOCK_PRINT

// Blacklisted : STOCK_PRINT_PREVIEW

// Blacklisted : STOCK_PROPERTIES

// Blacklisted : STOCK_QUIT

// Blacklisted : STOCK_REDO

// Blacklisted : STOCK_REFRESH

// Blacklisted : STOCK_REMOVE

// Blacklisted : STOCK_REVERT_TO_SAVED

// Blacklisted : STOCK_SAVE

// Blacklisted : STOCK_SAVE_AS

// Blacklisted : STOCK_SELECT_COLOR

// Blacklisted : STOCK_SELECT_FONT

// Blacklisted : STOCK_SORT_ASCENDING

// Blacklisted : STOCK_SORT_DESCENDING

// Blacklisted : STOCK_SPELL_CHECK

// Blacklisted : STOCK_STOP

// Blacklisted : STOCK_STRIKETHROUGH

// Blacklisted : STOCK_UNDELETE

// Blacklisted : STOCK_UNDERLINE

// Blacklisted : STOCK_UNDO

// Blacklisted : STOCK_YES

// Blacklisted : STOCK_ZOOM_100

// Blacklisted : STOCK_ZOOM_FIT

// Blacklisted : STOCK_ZOOM_IN

// Blacklisted : STOCK_ZOOM_OUT

// A CSS class to match an accelerator.
//
// Refer to individual widget documentation for used style classes.
const STYLE_CLASS_ACCELERATOR string = C.GTK_STYLE_CLASS_ACCELERATOR

// A CSS class used when rendering an arrow element.
//
// Refer to individual widget documentation for used style classes.
const STYLE_CLASS_ARROW string = C.GTK_STYLE_CLASS_ARROW

// A CSS class to match the window background.
//
// Refer to individual widget documentation for used style classes.
const STYLE_CLASS_BACKGROUND string = C.GTK_STYLE_CLASS_BACKGROUND

// A CSS class to indicate an area at the bottom of a widget.
//
// Refer to individual widget documentation for used style classes.
const STYLE_CLASS_BOTTOM string = C.GTK_STYLE_CLASS_BOTTOM

// A CSS class to match buttons.
//
// Refer to individual widget documentation for used style classes.
const STYLE_CLASS_BUTTON string = C.GTK_STYLE_CLASS_BUTTON

// A CSS class to match calendars.
//
// Refer to individual widget documentation for used style classes.
const STYLE_CLASS_CALENDAR string = C.GTK_STYLE_CLASS_CALENDAR

// A CSS class to match content rendered in cell views.
//
// Refer to individual widget documentation for used style classes.
const STYLE_CLASS_CELL string = C.GTK_STYLE_CLASS_CELL

// A CSS class to match check boxes.
//
// Refer to individual widget documentation for used style classes.
const STYLE_CLASS_CHECK string = C.GTK_STYLE_CLASS_CHECK

// A CSS class to match combobox entries.
//
// Refer to individual widget documentation for used style classes.
const STYLE_CLASS_COMBOBOX_ENTRY string = C.GTK_STYLE_CLASS_COMBOBOX_ENTRY

// A CSS class to match context menus.
//
// Refer to individual widget documentation for used style classes.
const STYLE_CLASS_CONTEXT_MENU string = C.GTK_STYLE_CLASS_CONTEXT_MENU

// A CSS class used when rendering a drag handle for
// text selection.
//
// Refer to individual widget documentation for used style classes.
const STYLE_CLASS_CURSOR_HANDLE string = C.GTK_STYLE_CLASS_CURSOR_HANDLE

// A CSS class to match the default widget.
//
// Refer to individual widget documentation for used style classes.
const STYLE_CLASS_DEFAULT string = C.GTK_STYLE_CLASS_DEFAULT

// A CSS class to match dimmed labels.
//
// Refer to individual widget documentation for used style classes.
const STYLE_CLASS_DIM_LABEL string = C.GTK_STYLE_CLASS_DIM_LABEL

// A CSS class for a drag-and-drop indicator.
//
// Refer to individual widget documentation for used style classes.
const STYLE_CLASS_DND string = C.GTK_STYLE_CLASS_DND

// A CSS class defining a dock area.
//
// Refer to individual widget documentation for used style classes.
const STYLE_CLASS_DOCK string = C.GTK_STYLE_CLASS_DOCK

// A CSS class to match text entries.
//
// Refer to individual widget documentation for used style classes.
const STYLE_CLASS_ENTRY string = C.GTK_STYLE_CLASS_ENTRY

// A CSS class for an area displaying an error message,
// such as those in infobars.
//
// Refer to individual widget documentation for used style classes.
const STYLE_CLASS_ERROR string = C.GTK_STYLE_CLASS_ERROR

// A CSS class defining an expander, such as those in treeviews.
//
// Refer to individual widget documentation for used style classes.
const STYLE_CLASS_EXPANDER string = C.GTK_STYLE_CLASS_EXPANDER

// A CSS class defining a frame delimiting content, such as
// #GtkFrame or the scrolled window frame around the
// scrollable area.
//
// Refer to individual widget documentation for used style classes.
const STYLE_CLASS_FRAME string = C.GTK_STYLE_CLASS_FRAME

// A CSS class defining a resize grip.
//
// Refer to individual widget documentation for used style classes.
const STYLE_CLASS_GRIP string = C.GTK_STYLE_CLASS_GRIP

// A CSS class to match a header element.
//
// Refer to individual widget documentation for used style classes.
const STYLE_CLASS_HEADER string = C.GTK_STYLE_CLASS_HEADER

// A CSS class defining a highlighted area, such as headings in
// assistants and calendars.
//
// Refer to individual widget documentation for used style classes.
const STYLE_CLASS_HIGHLIGHT string = C.GTK_STYLE_CLASS_HIGHLIGHT

// A CSS class for horizontally layered widgets.
//
// Refer to individual widget documentation for used style classes.
const STYLE_CLASS_HORIZONTAL string = C.GTK_STYLE_CLASS_HORIZONTAL

// A CSS class defining an image, such as the icon in an entry.
//
// Refer to individual widget documentation for used style classes.
const STYLE_CLASS_IMAGE string = C.GTK_STYLE_CLASS_IMAGE

// A CSS class for an area displaying an informational message,
// such as those in infobars.
//
// Refer to individual widget documentation for used style classes.
const STYLE_CLASS_INFO string = C.GTK_STYLE_CLASS_INFO

// A CSS class to match inline toolbars.
//
// Refer to individual widget documentation for used style classes.
const STYLE_CLASS_INLINE_TOOLBAR string = C.GTK_STYLE_CLASS_INLINE_TOOLBAR

// A CSS class used when rendering a drag handle for
// the insertion cursor position.
//
// Refer to individual widget documentation for used style classes.
const STYLE_CLASS_INSERTION_CURSOR string = C.GTK_STYLE_CLASS_INSERTION_CURSOR

// A CSS class to indicate an area at the left of a widget.
//
// Refer to individual widget documentation for used style classes.
const STYLE_CLASS_LEFT string = C.GTK_STYLE_CLASS_LEFT

// A CSS class used when rendering a level indicator, such
// as a battery charge level, or a password strength.
//
// Refer to individual widget documentation for used style classes.
const STYLE_CLASS_LEVEL_BAR string = C.GTK_STYLE_CLASS_LEVEL_BAR

// A CSS class to match a linked area, such as a box containing buttons
// belonging to the same control.
//
// Refer to individual widget documentation for used style classes.
const STYLE_CLASS_LINKED string = C.GTK_STYLE_CLASS_LINKED

// A CSS class to match lists.
//
// Refer to individual widget documentation for used style classes.
const STYLE_CLASS_LIST string = C.GTK_STYLE_CLASS_LIST

// A CSS class to match list rows.
//
// Refer to individual widget documentation for used style classes.
const STYLE_CLASS_LIST_ROW string = C.GTK_STYLE_CLASS_LIST_ROW

// A CSS class defining marks in a widget, such as in scales.
//
// Refer to individual widget documentation for used style classes.
const STYLE_CLASS_MARK string = C.GTK_STYLE_CLASS_MARK

// A CSS class to match menus.
//
// Refer to individual widget documentation for used style classes.
const STYLE_CLASS_MENU string = C.GTK_STYLE_CLASS_MENU

// A CSS class to menubars.
//
// Refer to individual widget documentation for used style classes.
const STYLE_CLASS_MENUBAR string = C.GTK_STYLE_CLASS_MENUBAR

// A CSS class to match menu items.
//
// Refer to individual widget documentation for used style classes.
const STYLE_CLASS_MENUITEM string = C.GTK_STYLE_CLASS_MENUITEM

// A CSS class defining a notebook.
//
// Refer to individual widget documentation for used style classes.
const STYLE_CLASS_NOTEBOOK string = C.GTK_STYLE_CLASS_NOTEBOOK

// A CSS class used when rendering an OSD (On Screen Display) element,
// on top of another container.
//
// Refer to individual widget documentation for used style classes.
const STYLE_CLASS_OSD string = C.GTK_STYLE_CLASS_OSD

// A CSS class for a pane separator, such as those in #GtkPaned.
//
// Refer to individual widget documentation for used style classes.
const STYLE_CLASS_PANE_SEPARATOR string = C.GTK_STYLE_CLASS_PANE_SEPARATOR

// A CSS class to match primary toolbars.
//
// Refer to individual widget documentation for used style classes.
const STYLE_CLASS_PRIMARY_TOOLBAR string = C.GTK_STYLE_CLASS_PRIMARY_TOOLBAR

// A CSS class to use when rendering activity as a progressbar.
//
// Refer to individual widget documentation for used style classes.
const STYLE_CLASS_PROGRESSBAR string = C.GTK_STYLE_CLASS_PROGRESSBAR

// A CSS class to use when rendering a pulse in an indeterminate progress bar.
//
// Refer to individual widget documentation for used style classes.
const STYLE_CLASS_PULSE string = C.GTK_STYLE_CLASS_PULSE

// A CSS class for an area displaying a question to the user,
// such as those in infobars.
//
// Refer to individual widget documentation for used style classes.
const STYLE_CLASS_QUESTION string = C.GTK_STYLE_CLASS_QUESTION

// A CSS class to match radio buttons.
//
// Refer to individual widget documentation for used style classes.
const STYLE_CLASS_RADIO string = C.GTK_STYLE_CLASS_RADIO

// A CSS class to match a raised control, such as a raised
// button on a toolbar.
//
// Refer to individual widget documentation for used style classes.
const STYLE_CLASS_RAISED string = C.GTK_STYLE_CLASS_RAISED

// A CSS class used to indicate a read-only state.
//
// Refer to individual widget documentation for used style classes.
const STYLE_CLASS_READ_ONLY string = C.GTK_STYLE_CLASS_READ_ONLY

// A CSS class to indicate an area at the right of a widget.
//
// Refer to individual widget documentation for used style classes.
const STYLE_CLASS_RIGHT string = C.GTK_STYLE_CLASS_RIGHT

// A CSS class to match the rubberband selection rectangle.
//
// Refer to individual widget documentation for used style classes.
const STYLE_CLASS_RUBBERBAND string = C.GTK_STYLE_CLASS_RUBBERBAND

// A CSS class to match scale widgets.
//
// Refer to individual widget documentation for used style classes.
const STYLE_CLASS_SCALE string = C.GTK_STYLE_CLASS_SCALE

// A CSS class to match scale widgets with marks attached,
// all the marks are above for horizontal #GtkScale.
// left for vertical #GtkScale.
//
// Refer to individual widget documentation for used style classes.
const STYLE_CLASS_SCALE_HAS_MARKS_ABOVE string = C.GTK_STYLE_CLASS_SCALE_HAS_MARKS_ABOVE

// A CSS class to match scale widgets with marks attached,
// all the marks are below for horizontal #GtkScale,
// right for vertical #GtkScale.
//
// Refer to individual widget documentation for used style classes.
const STYLE_CLASS_SCALE_HAS_MARKS_BELOW string = C.GTK_STYLE_CLASS_SCALE_HAS_MARKS_BELOW

// A CSS class to match scrollbars.
//
// Refer to individual widget documentation for used style classes.
const STYLE_CLASS_SCROLLBAR string = C.GTK_STYLE_CLASS_SCROLLBAR

// A CSS class to match the junction area between an horizontal
// and vertical scrollbar, when they’re both shown.
//
// Refer to individual widget documentation for used style classes.
const STYLE_CLASS_SCROLLBARS_JUNCTION string = C.GTK_STYLE_CLASS_SCROLLBARS_JUNCTION

// A CSS class for a separator.
//
// Refer to individual widget documentation for used style classes.
const STYLE_CLASS_SEPARATOR string = C.GTK_STYLE_CLASS_SEPARATOR

// A CSS class defining a sidebar, such as the left side in
// a file chooser.
//
// Refer to individual widget documentation for used style classes.
const STYLE_CLASS_SIDEBAR string = C.GTK_STYLE_CLASS_SIDEBAR

// A CSS class to match sliders.
//
// Refer to individual widget documentation for used style classes.
const STYLE_CLASS_SLIDER string = C.GTK_STYLE_CLASS_SLIDER

// A CSS class defining an spinbutton.
//
// Refer to individual widget documentation for used style classes.
const STYLE_CLASS_SPINBUTTON string = C.GTK_STYLE_CLASS_SPINBUTTON

// A CSS class to use when rendering activity as a “spinner”.
//
// Refer to individual widget documentation for used style classes.
const STYLE_CLASS_SPINNER string = C.GTK_STYLE_CLASS_SPINNER

// A CSS class used when rendering a titlebar in a toplevel window.
//
// Refer to individual widget documentation for used style classes.
const STYLE_CLASS_TITLEBAR string = C.GTK_STYLE_CLASS_TITLEBAR

// A CSS class to match toolbars.
//
// Refer to individual widget documentation for used style classes.
const STYLE_CLASS_TOOLBAR string = C.GTK_STYLE_CLASS_TOOLBAR

// A CSS class to match tooltip windows.
//
// Refer to individual widget documentation for used style classes.
const STYLE_CLASS_TOOLTIP string = C.GTK_STYLE_CLASS_TOOLTIP

// A CSS class to indicate an area at the top of a widget.
//
// Refer to individual widget documentation for used style classes.
const STYLE_CLASS_TOP string = C.GTK_STYLE_CLASS_TOP

// A CSS class to match troughs, as in scrollbars and progressbars.
//
// Refer to individual widget documentation for used style classes.
const STYLE_CLASS_TROUGH string = C.GTK_STYLE_CLASS_TROUGH

// A CSS class for vertically layered widgets.
//
// Refer to individual widget documentation for used style classes.
const STYLE_CLASS_VERTICAL string = C.GTK_STYLE_CLASS_VERTICAL

// A CSS class defining a view, such as iconviews or treeviews.
//
// Refer to individual widget documentation for used style classes.
const STYLE_CLASS_VIEW string = C.GTK_STYLE_CLASS_VIEW

// A CSS class for an area displaying a warning message,
// such as those in infobars.
//
// Refer to individual widget documentation for used style classes.
const STYLE_CLASS_WARNING string = C.GTK_STYLE_CLASS_WARNING

// A property holding the background color of rendered elements as a #GdkRGBA.
const STYLE_PROPERTY_BACKGROUND_COLOR string = C.GTK_STYLE_PROPERTY_BACKGROUND_COLOR

// A property holding the element’s background as a #cairo_pattern_t.
const STYLE_PROPERTY_BACKGROUND_IMAGE string = C.GTK_STYLE_PROPERTY_BACKGROUND_IMAGE

// A property holding the element’s border color as a #GdkRGBA.
const STYLE_PROPERTY_BORDER_COLOR string = C.GTK_STYLE_PROPERTY_BORDER_COLOR

// A property holding the rendered element’s border radius in pixels as a #gint.
const STYLE_PROPERTY_BORDER_RADIUS string = C.GTK_STYLE_PROPERTY_BORDER_RADIUS

// A property holding the element’s border style as a #GtkBorderStyle.
const STYLE_PROPERTY_BORDER_STYLE string = C.GTK_STYLE_PROPERTY_BORDER_STYLE

// A property holding the rendered element’s border width in pixels as
// a #GtkBorder. The border is the intermediary spacing property of the
// padding/border/margin series.
//
// gtk_render_frame() uses this property to find out the frame line width,
// so #GtkWidgets rendering frames may need to add up this padding when
// requesting size
const STYLE_PROPERTY_BORDER_WIDTH string = C.GTK_STYLE_PROPERTY_BORDER_WIDTH

// A property holding the foreground color of rendered elements as a #GdkRGBA.
const STYLE_PROPERTY_COLOR string = C.GTK_STYLE_PROPERTY_COLOR

// A property holding the font properties used when rendering text
// as a #PangoFontDescription.
const STYLE_PROPERTY_FONT string = C.GTK_STYLE_PROPERTY_FONT

// A property holding the rendered element’s margin as a #GtkBorder. The
// margin is defined as the spacing between the border of the element
// and its surrounding elements. It is external to #GtkWidget's
// size allocations, and the most external spacing property of the
// padding/border/margin series.
const STYLE_PROPERTY_MARGIN string = C.GTK_STYLE_PROPERTY_MARGIN

// A property holding the rendered element’s padding as a #GtkBorder. The
// padding is defined as the spacing between the inner part of the element border
// and its child. It’s the innermost spacing property of the padding/border/margin
// series.
const STYLE_PROPERTY_PADDING string = C.GTK_STYLE_PROPERTY_PADDING

// A priority that can be used when adding a #GtkStyleProvider
// for application-specific style information.
const STYLE_PROVIDER_PRIORITY_APPLICATION int = C.GTK_STYLE_PROVIDER_PRIORITY_APPLICATION

// The priority used for default style information
// that is used in the absence of themes.
//
// Note that this is not very useful for providing default
// styling for custom style classes - themes are likely to
// override styling provided at this priority with
// catch-all `* {...}` rules.
const STYLE_PROVIDER_PRIORITY_FALLBACK int = C.GTK_STYLE_PROVIDER_PRIORITY_FALLBACK

// The priority used for style information provided
// via #GtkSettings.
//
// This priority is higher than #GTK_STYLE_PROVIDER_PRIORITY_THEME
// to let settings override themes.
const STYLE_PROVIDER_PRIORITY_SETTINGS int = C.GTK_STYLE_PROVIDER_PRIORITY_SETTINGS

// The priority used for style information provided
// by themes.
const STYLE_PROVIDER_PRIORITY_THEME int = C.GTK_STYLE_PROVIDER_PRIORITY_THEME

// The priority used for the style information from
// `~/.gtk-3.0.css`.
//
// You should not use priorities higher than this, to
// give the user the last word.
const STYLE_PROVIDER_PRIORITY_USER int = C.GTK_STYLE_PROVIDER_PRIORITY_USER

// A widget region name to define a treeview column.
const STYLE_REGION_COLUMN string = C.GTK_STYLE_REGION_COLUMN

// A widget region name to define a treeview column header.
const STYLE_REGION_COLUMN_HEADER string = C.GTK_STYLE_REGION_COLUMN_HEADER

// A widget region name to define a treeview row.
const STYLE_REGION_ROW string = C.GTK_STYLE_REGION_ROW

// A widget region name to define a notebook tab.
const STYLE_REGION_TAB string = C.GTK_STYLE_REGION_TAB

// The priority at which the text view validates onscreen lines
// in an idle job in the background.
const TEXT_VIEW_PRIORITY_VALIDATE int = C.GTK_TEXT_VIEW_PRIORITY_VALIDATE

// The GTK_TREE_SORTABLE_DEFAULT_SORT_COLUMN_ID can be used to make a
// #GtkTreeSortable use the default sort function.
//
// See also gtk_tree_sortable_set_sort_column_id()
const TREE_SORTABLE_DEFAULT_SORT_COLUMN_ID int = C.GTK_TREE_SORTABLE_DEFAULT_SORT_COLUMN_ID

// The GTK_TREE_SORTABLE_DEFAULT_SORT_COLUMN_ID can be used to make a
// #GtkTreeSortable use no sorting.
//
// See also gtk_tree_sortable_set_sort_column_id()
const TREE_SORTABLE_UNSORTED_SORT_COLUMN_ID int = C.GTK_TREE_SORTABLE_UNSORTED_SORT_COLUMN_ID
