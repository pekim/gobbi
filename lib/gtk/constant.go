// This is a generated file - DO NOT EDIT

package gtk

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

const BINARY_AGE int32 = 2230
const INPUT_ERROR int32 = -1
const INTERFACE_AGE int32 = 30
const MAJOR_VERSION int32 = 3
const MAX_COMPOSE_LEN int32 = 7
const MICRO_VERSION int32 = 30
const MINOR_VERSION int32 = 22
const PAPER_NAME_A3 string = "iso_a3"
const PAPER_NAME_A4 string = "iso_a4"
const PAPER_NAME_A5 string = "iso_a5"
const PAPER_NAME_B5 string = "iso_b5"
const PAPER_NAME_EXECUTIVE string = "na_executive"
const PAPER_NAME_LEGAL string = "na_legal"
const PAPER_NAME_LETTER string = "na_letter"
const PATH_PRIO_MASK int32 = 15
const PRINT_SETTINGS_COLLATE string = "collate"
const PRINT_SETTINGS_DEFAULT_SOURCE string = "default-source"
const PRINT_SETTINGS_DITHER string = "dither"
const PRINT_SETTINGS_DUPLEX string = "duplex"
const PRINT_SETTINGS_FINISHINGS string = "finishings"
const PRINT_SETTINGS_MEDIA_TYPE string = "media-type"
const PRINT_SETTINGS_NUMBER_UP string = "number-up"
const PRINT_SETTINGS_NUMBER_UP_LAYOUT string = "number-up-layout"
const PRINT_SETTINGS_N_COPIES string = "n-copies"
const PRINT_SETTINGS_ORIENTATION string = "orientation"
const PRINT_SETTINGS_OUTPUT_BIN string = "output-bin"
const PRINT_SETTINGS_OUTPUT_FILE_FORMAT string = "output-file-format"
const PRINT_SETTINGS_OUTPUT_URI string = "output-uri"
const PRINT_SETTINGS_PAGE_RANGES string = "page-ranges"
const PRINT_SETTINGS_PAGE_SET string = "page-set"
const PRINT_SETTINGS_PAPER_FORMAT string = "paper-format"
const PRINT_SETTINGS_PAPER_HEIGHT string = "paper-height"
const PRINT_SETTINGS_PAPER_WIDTH string = "paper-width"
const PRINT_SETTINGS_PRINTER string = "printer"
const PRINT_SETTINGS_PRINTER_LPI string = "printer-lpi"
const PRINT_SETTINGS_PRINT_PAGES string = "print-pages"
const PRINT_SETTINGS_QUALITY string = "quality"
const PRINT_SETTINGS_RESOLUTION string = "resolution"
const PRINT_SETTINGS_RESOLUTION_X string = "resolution-x"
const PRINT_SETTINGS_RESOLUTION_Y string = "resolution-y"
const PRINT_SETTINGS_REVERSE string = "reverse"
const PRINT_SETTINGS_SCALE string = "scale"
const PRINT_SETTINGS_USE_COLOR string = "use-color"
const PRINT_SETTINGS_WIN32_DRIVER_EXTRA string = "win32-driver-extra"
const PRINT_SETTINGS_WIN32_DRIVER_VERSION string = "win32-driver-version"
const PRIORITY_RESIZE int32 = 10

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

const STYLE_CLASS_ACCELERATOR string = "accelerator"
const STYLE_CLASS_ARROW string = "arrow"
const STYLE_CLASS_BACKGROUND string = "background"
const STYLE_CLASS_BOTTOM string = "bottom"
const STYLE_CLASS_BUTTON string = "button"
const STYLE_CLASS_CALENDAR string = "calendar"
const STYLE_CLASS_CELL string = "cell"
const STYLE_CLASS_CHECK string = "check"
const STYLE_CLASS_COMBOBOX_ENTRY string = "combobox-entry"
const STYLE_CLASS_CONTEXT_MENU string = "context-menu"
const STYLE_CLASS_CURSOR_HANDLE string = "cursor-handle"
const STYLE_CLASS_DEFAULT string = "default"
const STYLE_CLASS_DIM_LABEL string = "dim-label"
const STYLE_CLASS_DND string = "dnd"
const STYLE_CLASS_DOCK string = "dock"
const STYLE_CLASS_ENTRY string = "entry"
const STYLE_CLASS_ERROR string = "error"
const STYLE_CLASS_EXPANDER string = "expander"
const STYLE_CLASS_FRAME string = "frame"
const STYLE_CLASS_GRIP string = "grip"
const STYLE_CLASS_HEADER string = "header"
const STYLE_CLASS_HIGHLIGHT string = "highlight"
const STYLE_CLASS_HORIZONTAL string = "horizontal"
const STYLE_CLASS_IMAGE string = "image"
const STYLE_CLASS_INFO string = "info"
const STYLE_CLASS_INLINE_TOOLBAR string = "inline-toolbar"
const STYLE_CLASS_INSERTION_CURSOR string = "insertion-cursor"
const STYLE_CLASS_LEFT string = "left"
const STYLE_CLASS_LEVEL_BAR string = "level-bar"
const STYLE_CLASS_LINKED string = "linked"
const STYLE_CLASS_LIST string = "list"
const STYLE_CLASS_LIST_ROW string = "list-row"
const STYLE_CLASS_MARK string = "mark"
const STYLE_CLASS_MENU string = "menu"
const STYLE_CLASS_MENUBAR string = "menubar"
const STYLE_CLASS_MENUITEM string = "menuitem"
const STYLE_CLASS_NOTEBOOK string = "notebook"
const STYLE_CLASS_OSD string = "osd"
const STYLE_CLASS_PANE_SEPARATOR string = "pane-separator"
const STYLE_CLASS_PRIMARY_TOOLBAR string = "primary-toolbar"
const STYLE_CLASS_PROGRESSBAR string = "progressbar"
const STYLE_CLASS_PULSE string = "pulse"
const STYLE_CLASS_QUESTION string = "question"
const STYLE_CLASS_RADIO string = "radio"
const STYLE_CLASS_RAISED string = "raised"
const STYLE_CLASS_READ_ONLY string = "read-only"
const STYLE_CLASS_RIGHT string = "right"
const STYLE_CLASS_RUBBERBAND string = "rubberband"
const STYLE_CLASS_SCALE string = "scale"
const STYLE_CLASS_SCALE_HAS_MARKS_ABOVE string = "scale-has-marks-above"
const STYLE_CLASS_SCALE_HAS_MARKS_BELOW string = "scale-has-marks-below"
const STYLE_CLASS_SCROLLBAR string = "scrollbar"
const STYLE_CLASS_SCROLLBARS_JUNCTION string = "scrollbars-junction"
const STYLE_CLASS_SEPARATOR string = "separator"
const STYLE_CLASS_SIDEBAR string = "sidebar"
const STYLE_CLASS_SLIDER string = "slider"
const STYLE_CLASS_SPINBUTTON string = "spinbutton"
const STYLE_CLASS_SPINNER string = "spinner"
const STYLE_CLASS_TITLEBAR string = "titlebar"
const STYLE_CLASS_TOOLBAR string = "toolbar"
const STYLE_CLASS_TOOLTIP string = "tooltip"
const STYLE_CLASS_TOP string = "top"
const STYLE_CLASS_TROUGH string = "trough"
const STYLE_CLASS_VERTICAL string = "vertical"
const STYLE_CLASS_VIEW string = "view"
const STYLE_CLASS_WARNING string = "warning"
const STYLE_PROPERTY_BACKGROUND_COLOR string = "background-color"
const STYLE_PROPERTY_BACKGROUND_IMAGE string = "background-image"
const STYLE_PROPERTY_BORDER_COLOR string = "border-color"
const STYLE_PROPERTY_BORDER_RADIUS string = "border-radius"
const STYLE_PROPERTY_BORDER_STYLE string = "border-style"
const STYLE_PROPERTY_BORDER_WIDTH string = "border-width"
const STYLE_PROPERTY_COLOR string = "color"
const STYLE_PROPERTY_FONT string = "font"
const STYLE_PROPERTY_MARGIN string = "margin"
const STYLE_PROPERTY_PADDING string = "padding"
const STYLE_PROVIDER_PRIORITY_APPLICATION int32 = 600
const STYLE_PROVIDER_PRIORITY_FALLBACK int32 = 1
const STYLE_PROVIDER_PRIORITY_SETTINGS int32 = 400
const STYLE_PROVIDER_PRIORITY_THEME int32 = 200
const STYLE_PROVIDER_PRIORITY_USER int32 = 800
const STYLE_REGION_COLUMN string = "column"
const STYLE_REGION_COLUMN_HEADER string = "column-header"
const STYLE_REGION_ROW string = "row"
const STYLE_REGION_TAB string = "tab"
const TEXT_VIEW_PRIORITY_VALIDATE int32 = 5
const TREE_SORTABLE_DEFAULT_SORT_COLUMN_ID int32 = -1
const TREE_SORTABLE_UNSORTED_SORT_COLUMN_ID int32 = -2
