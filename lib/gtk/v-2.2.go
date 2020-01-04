// Code generated - DO NOT EDIT.
// +build gtk_2.2

package gtk

import (
	cairo "github.com/pekim/gobbi/lib/cairo"
	gdk "github.com/pekim/gobbi/lib/gdk"
	gdkpixbuf "github.com/pekim/gobbi/lib/gdkpixbuf"
	glib "github.com/pekim/gobbi/lib/glib"
	gobject "github.com/pekim/gobbi/lib/gobject"
	gtk "github.com/pekim/gobbi/lib/internal/c/gtk"
	pango "github.com/pekim/gobbi/lib/pango"
	"unsafe"
)

// Allocation is a representation of the C alias GtkAllocation.
type Allocation gdk.Rectangle

// Stock is a representation of the C alias GtkStock.
type Stock string

// BINARY_AGE is a representation of the C constant GTK_BINARY_AGE.
const BINARY_AGE = 2412

// INPUT_ERROR is a representation of the C constant GTK_INPUT_ERROR.
const INPUT_ERROR = -1

// INTERFACE_AGE is a representation of the C constant GTK_INTERFACE_AGE.
const INTERFACE_AGE = 8

// MAJOR_VERSION is a representation of the C constant GTK_MAJOR_VERSION.
const MAJOR_VERSION = 3

// MAX_COMPOSE_LEN is a representation of the C constant GTK_MAX_COMPOSE_LEN.
const MAX_COMPOSE_LEN = 7

// MICRO_VERSION is a representation of the C constant GTK_MICRO_VERSION.
const MICRO_VERSION = 12

// MINOR_VERSION is a representation of the C constant GTK_MINOR_VERSION.
const MINOR_VERSION = 24

// PAPER_NAME_A3 is a representation of the C constant GTK_PAPER_NAME_A3.
const PAPER_NAME_A3 = "iso_a3"

// PAPER_NAME_A4 is a representation of the C constant GTK_PAPER_NAME_A4.
const PAPER_NAME_A4 = "iso_a4"

// PAPER_NAME_A5 is a representation of the C constant GTK_PAPER_NAME_A5.
const PAPER_NAME_A5 = "iso_a5"

// PAPER_NAME_B5 is a representation of the C constant GTK_PAPER_NAME_B5.
const PAPER_NAME_B5 = "iso_b5"

// PAPER_NAME_EXECUTIVE is a representation of the C constant GTK_PAPER_NAME_EXECUTIVE.
const PAPER_NAME_EXECUTIVE = "na_executive"

// PAPER_NAME_LEGAL is a representation of the C constant GTK_PAPER_NAME_LEGAL.
const PAPER_NAME_LEGAL = "na_legal"

// PAPER_NAME_LETTER is a representation of the C constant GTK_PAPER_NAME_LETTER.
const PAPER_NAME_LETTER = "na_letter"

// PATH_PRIO_MASK is a representation of the C constant GTK_PATH_PRIO_MASK.
const PATH_PRIO_MASK = 15

// PRINT_SETTINGS_COLLATE is a representation of the C constant GTK_PRINT_SETTINGS_COLLATE.
const PRINT_SETTINGS_COLLATE = "collate"

// PRINT_SETTINGS_DEFAULT_SOURCE is a representation of the C constant GTK_PRINT_SETTINGS_DEFAULT_SOURCE.
const PRINT_SETTINGS_DEFAULT_SOURCE = "default-source"

// PRINT_SETTINGS_DITHER is a representation of the C constant GTK_PRINT_SETTINGS_DITHER.
const PRINT_SETTINGS_DITHER = "dither"

// PRINT_SETTINGS_DUPLEX is a representation of the C constant GTK_PRINT_SETTINGS_DUPLEX.
const PRINT_SETTINGS_DUPLEX = "duplex"

// PRINT_SETTINGS_FINISHINGS is a representation of the C constant GTK_PRINT_SETTINGS_FINISHINGS.
const PRINT_SETTINGS_FINISHINGS = "finishings"

// PRINT_SETTINGS_MEDIA_TYPE is a representation of the C constant GTK_PRINT_SETTINGS_MEDIA_TYPE.
const PRINT_SETTINGS_MEDIA_TYPE = "media-type"

// PRINT_SETTINGS_NUMBER_UP is a representation of the C constant GTK_PRINT_SETTINGS_NUMBER_UP.
const PRINT_SETTINGS_NUMBER_UP = "number-up"

// PRINT_SETTINGS_NUMBER_UP_LAYOUT is a representation of the C constant GTK_PRINT_SETTINGS_NUMBER_UP_LAYOUT.
const PRINT_SETTINGS_NUMBER_UP_LAYOUT = "number-up-layout"

// PRINT_SETTINGS_N_COPIES is a representation of the C constant GTK_PRINT_SETTINGS_N_COPIES.
const PRINT_SETTINGS_N_COPIES = "n-copies"

// PRINT_SETTINGS_ORIENTATION is a representation of the C constant GTK_PRINT_SETTINGS_ORIENTATION.
const PRINT_SETTINGS_ORIENTATION = "orientation"

// PRINT_SETTINGS_OUTPUT_BIN is a representation of the C constant GTK_PRINT_SETTINGS_OUTPUT_BIN.
const PRINT_SETTINGS_OUTPUT_BIN = "output-bin"

// PRINT_SETTINGS_OUTPUT_FILE_FORMAT is a representation of the C constant GTK_PRINT_SETTINGS_OUTPUT_FILE_FORMAT.
const PRINT_SETTINGS_OUTPUT_FILE_FORMAT = "output-file-format"

// PRINT_SETTINGS_OUTPUT_URI is a representation of the C constant GTK_PRINT_SETTINGS_OUTPUT_URI.
const PRINT_SETTINGS_OUTPUT_URI = "output-uri"

// PRINT_SETTINGS_PAGE_RANGES is a representation of the C constant GTK_PRINT_SETTINGS_PAGE_RANGES.
const PRINT_SETTINGS_PAGE_RANGES = "page-ranges"

// PRINT_SETTINGS_PAGE_SET is a representation of the C constant GTK_PRINT_SETTINGS_PAGE_SET.
const PRINT_SETTINGS_PAGE_SET = "page-set"

// PRINT_SETTINGS_PAPER_FORMAT is a representation of the C constant GTK_PRINT_SETTINGS_PAPER_FORMAT.
const PRINT_SETTINGS_PAPER_FORMAT = "paper-format"

// PRINT_SETTINGS_PAPER_HEIGHT is a representation of the C constant GTK_PRINT_SETTINGS_PAPER_HEIGHT.
const PRINT_SETTINGS_PAPER_HEIGHT = "paper-height"

// PRINT_SETTINGS_PAPER_WIDTH is a representation of the C constant GTK_PRINT_SETTINGS_PAPER_WIDTH.
const PRINT_SETTINGS_PAPER_WIDTH = "paper-width"

// PRINT_SETTINGS_PRINTER is a representation of the C constant GTK_PRINT_SETTINGS_PRINTER.
const PRINT_SETTINGS_PRINTER = "printer"

// PRINT_SETTINGS_PRINTER_LPI is a representation of the C constant GTK_PRINT_SETTINGS_PRINTER_LPI.
const PRINT_SETTINGS_PRINTER_LPI = "printer-lpi"

// PRINT_SETTINGS_PRINT_PAGES is a representation of the C constant GTK_PRINT_SETTINGS_PRINT_PAGES.
const PRINT_SETTINGS_PRINT_PAGES = "print-pages"

// PRINT_SETTINGS_QUALITY is a representation of the C constant GTK_PRINT_SETTINGS_QUALITY.
const PRINT_SETTINGS_QUALITY = "quality"

// PRINT_SETTINGS_RESOLUTION is a representation of the C constant GTK_PRINT_SETTINGS_RESOLUTION.
const PRINT_SETTINGS_RESOLUTION = "resolution"

// PRINT_SETTINGS_RESOLUTION_X is a representation of the C constant GTK_PRINT_SETTINGS_RESOLUTION_X.
const PRINT_SETTINGS_RESOLUTION_X = "resolution-x"

// PRINT_SETTINGS_RESOLUTION_Y is a representation of the C constant GTK_PRINT_SETTINGS_RESOLUTION_Y.
const PRINT_SETTINGS_RESOLUTION_Y = "resolution-y"

// PRINT_SETTINGS_REVERSE is a representation of the C constant GTK_PRINT_SETTINGS_REVERSE.
const PRINT_SETTINGS_REVERSE = "reverse"

// PRINT_SETTINGS_SCALE is a representation of the C constant GTK_PRINT_SETTINGS_SCALE.
const PRINT_SETTINGS_SCALE = "scale"

// PRINT_SETTINGS_USE_COLOR is a representation of the C constant GTK_PRINT_SETTINGS_USE_COLOR.
const PRINT_SETTINGS_USE_COLOR = "use-color"

// PRINT_SETTINGS_WIN32_DRIVER_EXTRA is a representation of the C constant GTK_PRINT_SETTINGS_WIN32_DRIVER_EXTRA.
const PRINT_SETTINGS_WIN32_DRIVER_EXTRA = "win32-driver-extra"

// PRINT_SETTINGS_WIN32_DRIVER_VERSION is a representation of the C constant GTK_PRINT_SETTINGS_WIN32_DRIVER_VERSION.
const PRINT_SETTINGS_WIN32_DRIVER_VERSION = "win32-driver-version"

// PRIORITY_RESIZE is a representation of the C constant GTK_PRIORITY_RESIZE.
const PRIORITY_RESIZE = 110

// STOCK_ADD is a representation of the C constant GTK_STOCK_ADD.
const STOCK_ADD = "gtk-add"

// STOCK_APPLY is a representation of the C constant GTK_STOCK_APPLY.
const STOCK_APPLY = "gtk-apply"

// STOCK_BOLD is a representation of the C constant GTK_STOCK_BOLD.
const STOCK_BOLD = "gtk-bold"

// STOCK_CANCEL is a representation of the C constant GTK_STOCK_CANCEL.
const STOCK_CANCEL = "gtk-cancel"

// STOCK_CDROM is a representation of the C constant GTK_STOCK_CDROM.
const STOCK_CDROM = "gtk-cdrom"

// STOCK_CLEAR is a representation of the C constant GTK_STOCK_CLEAR.
const STOCK_CLEAR = "gtk-clear"

// STOCK_CLOSE is a representation of the C constant GTK_STOCK_CLOSE.
const STOCK_CLOSE = "gtk-close"

// STOCK_COLOR_PICKER is a representation of the C constant GTK_STOCK_COLOR_PICKER.
//
// since 2.2
const STOCK_COLOR_PICKER = "gtk-color-picker"

// STOCK_CONVERT is a representation of the C constant GTK_STOCK_CONVERT.
const STOCK_CONVERT = "gtk-convert"

// STOCK_COPY is a representation of the C constant GTK_STOCK_COPY.
const STOCK_COPY = "gtk-copy"

// STOCK_CUT is a representation of the C constant GTK_STOCK_CUT.
const STOCK_CUT = "gtk-cut"

// STOCK_DELETE is a representation of the C constant GTK_STOCK_DELETE.
const STOCK_DELETE = "gtk-delete"

// STOCK_DIALOG_ERROR is a representation of the C constant GTK_STOCK_DIALOG_ERROR.
const STOCK_DIALOG_ERROR = "gtk-dialog-error"

// STOCK_DIALOG_INFO is a representation of the C constant GTK_STOCK_DIALOG_INFO.
const STOCK_DIALOG_INFO = "gtk-dialog-info"

// STOCK_DIALOG_QUESTION is a representation of the C constant GTK_STOCK_DIALOG_QUESTION.
const STOCK_DIALOG_QUESTION = "gtk-dialog-question"

// STOCK_DIALOG_WARNING is a representation of the C constant GTK_STOCK_DIALOG_WARNING.
const STOCK_DIALOG_WARNING = "gtk-dialog-warning"

// STOCK_DND is a representation of the C constant GTK_STOCK_DND.
const STOCK_DND = "gtk-dnd"

// STOCK_DND_MULTIPLE is a representation of the C constant GTK_STOCK_DND_MULTIPLE.
const STOCK_DND_MULTIPLE = "gtk-dnd-multiple"

// STOCK_EXECUTE is a representation of the C constant GTK_STOCK_EXECUTE.
const STOCK_EXECUTE = "gtk-execute"

// STOCK_FIND is a representation of the C constant GTK_STOCK_FIND.
const STOCK_FIND = "gtk-find"

// STOCK_FIND_AND_REPLACE is a representation of the C constant GTK_STOCK_FIND_AND_REPLACE.
const STOCK_FIND_AND_REPLACE = "gtk-find-and-replace"

// STOCK_FLOPPY is a representation of the C constant GTK_STOCK_FLOPPY.
const STOCK_FLOPPY = "gtk-floppy"

// STOCK_GOTO_BOTTOM is a representation of the C constant GTK_STOCK_GOTO_BOTTOM.
const STOCK_GOTO_BOTTOM = "gtk-goto-bottom"

// STOCK_GOTO_FIRST is a representation of the C constant GTK_STOCK_GOTO_FIRST.
const STOCK_GOTO_FIRST = "gtk-goto-first"

// STOCK_GOTO_LAST is a representation of the C constant GTK_STOCK_GOTO_LAST.
const STOCK_GOTO_LAST = "gtk-goto-last"

// STOCK_GOTO_TOP is a representation of the C constant GTK_STOCK_GOTO_TOP.
const STOCK_GOTO_TOP = "gtk-goto-top"

// STOCK_GO_BACK is a representation of the C constant GTK_STOCK_GO_BACK.
const STOCK_GO_BACK = "gtk-go-back"

// STOCK_GO_DOWN is a representation of the C constant GTK_STOCK_GO_DOWN.
const STOCK_GO_DOWN = "gtk-go-down"

// STOCK_GO_FORWARD is a representation of the C constant GTK_STOCK_GO_FORWARD.
const STOCK_GO_FORWARD = "gtk-go-forward"

// STOCK_GO_UP is a representation of the C constant GTK_STOCK_GO_UP.
const STOCK_GO_UP = "gtk-go-up"

// STOCK_HELP is a representation of the C constant GTK_STOCK_HELP.
const STOCK_HELP = "gtk-help"

// STOCK_HOME is a representation of the C constant GTK_STOCK_HOME.
const STOCK_HOME = "gtk-home"

// STOCK_INDEX is a representation of the C constant GTK_STOCK_INDEX.
const STOCK_INDEX = "gtk-index"

// STOCK_ITALIC is a representation of the C constant GTK_STOCK_ITALIC.
const STOCK_ITALIC = "gtk-italic"

// STOCK_JUMP_TO is a representation of the C constant GTK_STOCK_JUMP_TO.
const STOCK_JUMP_TO = "gtk-jump-to"

// STOCK_JUSTIFY_CENTER is a representation of the C constant GTK_STOCK_JUSTIFY_CENTER.
const STOCK_JUSTIFY_CENTER = "gtk-justify-center"

// STOCK_JUSTIFY_FILL is a representation of the C constant GTK_STOCK_JUSTIFY_FILL.
const STOCK_JUSTIFY_FILL = "gtk-justify-fill"

// STOCK_JUSTIFY_LEFT is a representation of the C constant GTK_STOCK_JUSTIFY_LEFT.
const STOCK_JUSTIFY_LEFT = "gtk-justify-left"

// STOCK_JUSTIFY_RIGHT is a representation of the C constant GTK_STOCK_JUSTIFY_RIGHT.
const STOCK_JUSTIFY_RIGHT = "gtk-justify-right"

// STOCK_MISSING_IMAGE is a representation of the C constant GTK_STOCK_MISSING_IMAGE.
const STOCK_MISSING_IMAGE = "gtk-missing-image"

// STOCK_NEW is a representation of the C constant GTK_STOCK_NEW.
const STOCK_NEW = "gtk-new"

// STOCK_NO is a representation of the C constant GTK_STOCK_NO.
const STOCK_NO = "gtk-no"

// STOCK_OK is a representation of the C constant GTK_STOCK_OK.
const STOCK_OK = "gtk-ok"

// STOCK_OPEN is a representation of the C constant GTK_STOCK_OPEN.
const STOCK_OPEN = "gtk-open"

// STOCK_PASTE is a representation of the C constant GTK_STOCK_PASTE.
const STOCK_PASTE = "gtk-paste"

// STOCK_PREFERENCES is a representation of the C constant GTK_STOCK_PREFERENCES.
const STOCK_PREFERENCES = "gtk-preferences"

// STOCK_PRINT is a representation of the C constant GTK_STOCK_PRINT.
const STOCK_PRINT = "gtk-print"

// STOCK_PRINT_PREVIEW is a representation of the C constant GTK_STOCK_PRINT_PREVIEW.
const STOCK_PRINT_PREVIEW = "gtk-print-preview"

// STOCK_PROPERTIES is a representation of the C constant GTK_STOCK_PROPERTIES.
const STOCK_PROPERTIES = "gtk-properties"

// STOCK_QUIT is a representation of the C constant GTK_STOCK_QUIT.
const STOCK_QUIT = "gtk-quit"

// STOCK_REDO is a representation of the C constant GTK_STOCK_REDO.
const STOCK_REDO = "gtk-redo"

// STOCK_REFRESH is a representation of the C constant GTK_STOCK_REFRESH.
const STOCK_REFRESH = "gtk-refresh"

// STOCK_REMOVE is a representation of the C constant GTK_STOCK_REMOVE.
const STOCK_REMOVE = "gtk-remove"

// STOCK_REVERT_TO_SAVED is a representation of the C constant GTK_STOCK_REVERT_TO_SAVED.
const STOCK_REVERT_TO_SAVED = "gtk-revert-to-saved"

// STOCK_SAVE is a representation of the C constant GTK_STOCK_SAVE.
const STOCK_SAVE = "gtk-save"

// STOCK_SAVE_AS is a representation of the C constant GTK_STOCK_SAVE_AS.
const STOCK_SAVE_AS = "gtk-save-as"

// STOCK_SELECT_COLOR is a representation of the C constant GTK_STOCK_SELECT_COLOR.
const STOCK_SELECT_COLOR = "gtk-select-color"

// STOCK_SELECT_FONT is a representation of the C constant GTK_STOCK_SELECT_FONT.
const STOCK_SELECT_FONT = "gtk-select-font"

// STOCK_SORT_ASCENDING is a representation of the C constant GTK_STOCK_SORT_ASCENDING.
const STOCK_SORT_ASCENDING = "gtk-sort-ascending"

// STOCK_SORT_DESCENDING is a representation of the C constant GTK_STOCK_SORT_DESCENDING.
const STOCK_SORT_DESCENDING = "gtk-sort-descending"

// STOCK_SPELL_CHECK is a representation of the C constant GTK_STOCK_SPELL_CHECK.
const STOCK_SPELL_CHECK = "gtk-spell-check"

// STOCK_STOP is a representation of the C constant GTK_STOCK_STOP.
const STOCK_STOP = "gtk-stop"

// STOCK_STRIKETHROUGH is a representation of the C constant GTK_STOCK_STRIKETHROUGH.
const STOCK_STRIKETHROUGH = "gtk-strikethrough"

// STOCK_UNDELETE is a representation of the C constant GTK_STOCK_UNDELETE.
const STOCK_UNDELETE = "gtk-undelete"

// STOCK_UNDERLINE is a representation of the C constant GTK_STOCK_UNDERLINE.
const STOCK_UNDERLINE = "gtk-underline"

// STOCK_UNDO is a representation of the C constant GTK_STOCK_UNDO.
const STOCK_UNDO = "gtk-undo"

// STOCK_YES is a representation of the C constant GTK_STOCK_YES.
const STOCK_YES = "gtk-yes"

// STOCK_ZOOM_100 is a representation of the C constant GTK_STOCK_ZOOM_100.
const STOCK_ZOOM_100 = "gtk-zoom-100"

// STOCK_ZOOM_FIT is a representation of the C constant GTK_STOCK_ZOOM_FIT.
const STOCK_ZOOM_FIT = "gtk-zoom-fit"

// STOCK_ZOOM_IN is a representation of the C constant GTK_STOCK_ZOOM_IN.
const STOCK_ZOOM_IN = "gtk-zoom-in"

// STOCK_ZOOM_OUT is a representation of the C constant GTK_STOCK_ZOOM_OUT.
const STOCK_ZOOM_OUT = "gtk-zoom-out"

// STYLE_CLASS_ACCELERATOR is a representation of the C constant GTK_STYLE_CLASS_ACCELERATOR.
const STYLE_CLASS_ACCELERATOR = "accelerator"

// STYLE_CLASS_ARROW is a representation of the C constant GTK_STYLE_CLASS_ARROW.
const STYLE_CLASS_ARROW = "arrow"

// STYLE_CLASS_BACKGROUND is a representation of the C constant GTK_STYLE_CLASS_BACKGROUND.
const STYLE_CLASS_BACKGROUND = "background"

// STYLE_CLASS_BOTTOM is a representation of the C constant GTK_STYLE_CLASS_BOTTOM.
const STYLE_CLASS_BOTTOM = "bottom"

// STYLE_CLASS_BUTTON is a representation of the C constant GTK_STYLE_CLASS_BUTTON.
const STYLE_CLASS_BUTTON = "button"

// STYLE_CLASS_CALENDAR is a representation of the C constant GTK_STYLE_CLASS_CALENDAR.
const STYLE_CLASS_CALENDAR = "calendar"

// STYLE_CLASS_CELL is a representation of the C constant GTK_STYLE_CLASS_CELL.
const STYLE_CLASS_CELL = "cell"

// STYLE_CLASS_CHECK is a representation of the C constant GTK_STYLE_CLASS_CHECK.
const STYLE_CLASS_CHECK = "check"

// STYLE_CLASS_COMBOBOX_ENTRY is a representation of the C constant GTK_STYLE_CLASS_COMBOBOX_ENTRY.
const STYLE_CLASS_COMBOBOX_ENTRY = "combobox-entry"

// STYLE_CLASS_CONTEXT_MENU is a representation of the C constant GTK_STYLE_CLASS_CONTEXT_MENU.
const STYLE_CLASS_CONTEXT_MENU = "context-menu"

// STYLE_CLASS_CURSOR_HANDLE is a representation of the C constant GTK_STYLE_CLASS_CURSOR_HANDLE.
const STYLE_CLASS_CURSOR_HANDLE = "cursor-handle"

// STYLE_CLASS_DEFAULT is a representation of the C constant GTK_STYLE_CLASS_DEFAULT.
const STYLE_CLASS_DEFAULT = "default"

// STYLE_CLASS_DIM_LABEL is a representation of the C constant GTK_STYLE_CLASS_DIM_LABEL.
const STYLE_CLASS_DIM_LABEL = "dim-label"

// STYLE_CLASS_DND is a representation of the C constant GTK_STYLE_CLASS_DND.
const STYLE_CLASS_DND = "dnd"

// STYLE_CLASS_DOCK is a representation of the C constant GTK_STYLE_CLASS_DOCK.
const STYLE_CLASS_DOCK = "dock"

// STYLE_CLASS_ENTRY is a representation of the C constant GTK_STYLE_CLASS_ENTRY.
const STYLE_CLASS_ENTRY = "entry"

// STYLE_CLASS_ERROR is a representation of the C constant GTK_STYLE_CLASS_ERROR.
const STYLE_CLASS_ERROR = "error"

// STYLE_CLASS_EXPANDER is a representation of the C constant GTK_STYLE_CLASS_EXPANDER.
const STYLE_CLASS_EXPANDER = "expander"

// STYLE_CLASS_FRAME is a representation of the C constant GTK_STYLE_CLASS_FRAME.
const STYLE_CLASS_FRAME = "frame"

// STYLE_CLASS_GRIP is a representation of the C constant GTK_STYLE_CLASS_GRIP.
const STYLE_CLASS_GRIP = "grip"

// STYLE_CLASS_HEADER is a representation of the C constant GTK_STYLE_CLASS_HEADER.
const STYLE_CLASS_HEADER = "header"

// STYLE_CLASS_HIGHLIGHT is a representation of the C constant GTK_STYLE_CLASS_HIGHLIGHT.
const STYLE_CLASS_HIGHLIGHT = "highlight"

// STYLE_CLASS_HORIZONTAL is a representation of the C constant GTK_STYLE_CLASS_HORIZONTAL.
const STYLE_CLASS_HORIZONTAL = "horizontal"

// STYLE_CLASS_IMAGE is a representation of the C constant GTK_STYLE_CLASS_IMAGE.
const STYLE_CLASS_IMAGE = "image"

// STYLE_CLASS_INFO is a representation of the C constant GTK_STYLE_CLASS_INFO.
const STYLE_CLASS_INFO = "info"

// STYLE_CLASS_INLINE_TOOLBAR is a representation of the C constant GTK_STYLE_CLASS_INLINE_TOOLBAR.
const STYLE_CLASS_INLINE_TOOLBAR = "inline-toolbar"

// STYLE_CLASS_INSERTION_CURSOR is a representation of the C constant GTK_STYLE_CLASS_INSERTION_CURSOR.
const STYLE_CLASS_INSERTION_CURSOR = "insertion-cursor"

// STYLE_CLASS_LEFT is a representation of the C constant GTK_STYLE_CLASS_LEFT.
const STYLE_CLASS_LEFT = "left"

// STYLE_CLASS_LEVEL_BAR is a representation of the C constant GTK_STYLE_CLASS_LEVEL_BAR.
const STYLE_CLASS_LEVEL_BAR = "level-bar"

// STYLE_CLASS_LINKED is a representation of the C constant GTK_STYLE_CLASS_LINKED.
const STYLE_CLASS_LINKED = "linked"

// STYLE_CLASS_LIST is a representation of the C constant GTK_STYLE_CLASS_LIST.
const STYLE_CLASS_LIST = "list"

// STYLE_CLASS_LIST_ROW is a representation of the C constant GTK_STYLE_CLASS_LIST_ROW.
const STYLE_CLASS_LIST_ROW = "list-row"

// STYLE_CLASS_MARK is a representation of the C constant GTK_STYLE_CLASS_MARK.
const STYLE_CLASS_MARK = "mark"

// STYLE_CLASS_MENU is a representation of the C constant GTK_STYLE_CLASS_MENU.
const STYLE_CLASS_MENU = "menu"

// STYLE_CLASS_MENUBAR is a representation of the C constant GTK_STYLE_CLASS_MENUBAR.
const STYLE_CLASS_MENUBAR = "menubar"

// STYLE_CLASS_MENUITEM is a representation of the C constant GTK_STYLE_CLASS_MENUITEM.
const STYLE_CLASS_MENUITEM = "menuitem"

// STYLE_CLASS_NOTEBOOK is a representation of the C constant GTK_STYLE_CLASS_NOTEBOOK.
const STYLE_CLASS_NOTEBOOK = "notebook"

// STYLE_CLASS_OSD is a representation of the C constant GTK_STYLE_CLASS_OSD.
const STYLE_CLASS_OSD = "osd"

// STYLE_CLASS_PANE_SEPARATOR is a representation of the C constant GTK_STYLE_CLASS_PANE_SEPARATOR.
const STYLE_CLASS_PANE_SEPARATOR = "pane-separator"

// STYLE_CLASS_PRIMARY_TOOLBAR is a representation of the C constant GTK_STYLE_CLASS_PRIMARY_TOOLBAR.
const STYLE_CLASS_PRIMARY_TOOLBAR = "primary-toolbar"

// STYLE_CLASS_PROGRESSBAR is a representation of the C constant GTK_STYLE_CLASS_PROGRESSBAR.
const STYLE_CLASS_PROGRESSBAR = "progressbar"

// STYLE_CLASS_PULSE is a representation of the C constant GTK_STYLE_CLASS_PULSE.
const STYLE_CLASS_PULSE = "pulse"

// STYLE_CLASS_QUESTION is a representation of the C constant GTK_STYLE_CLASS_QUESTION.
const STYLE_CLASS_QUESTION = "question"

// STYLE_CLASS_RADIO is a representation of the C constant GTK_STYLE_CLASS_RADIO.
const STYLE_CLASS_RADIO = "radio"

// STYLE_CLASS_RAISED is a representation of the C constant GTK_STYLE_CLASS_RAISED.
const STYLE_CLASS_RAISED = "raised"

// STYLE_CLASS_READ_ONLY is a representation of the C constant GTK_STYLE_CLASS_READ_ONLY.
const STYLE_CLASS_READ_ONLY = "read-only"

// STYLE_CLASS_RIGHT is a representation of the C constant GTK_STYLE_CLASS_RIGHT.
const STYLE_CLASS_RIGHT = "right"

// STYLE_CLASS_RUBBERBAND is a representation of the C constant GTK_STYLE_CLASS_RUBBERBAND.
const STYLE_CLASS_RUBBERBAND = "rubberband"

// STYLE_CLASS_SCALE is a representation of the C constant GTK_STYLE_CLASS_SCALE.
const STYLE_CLASS_SCALE = "scale"

// STYLE_CLASS_SCALE_HAS_MARKS_ABOVE is a representation of the C constant GTK_STYLE_CLASS_SCALE_HAS_MARKS_ABOVE.
const STYLE_CLASS_SCALE_HAS_MARKS_ABOVE = "scale-has-marks-above"

// STYLE_CLASS_SCALE_HAS_MARKS_BELOW is a representation of the C constant GTK_STYLE_CLASS_SCALE_HAS_MARKS_BELOW.
const STYLE_CLASS_SCALE_HAS_MARKS_BELOW = "scale-has-marks-below"

// STYLE_CLASS_SCROLLBAR is a representation of the C constant GTK_STYLE_CLASS_SCROLLBAR.
const STYLE_CLASS_SCROLLBAR = "scrollbar"

// STYLE_CLASS_SCROLLBARS_JUNCTION is a representation of the C constant GTK_STYLE_CLASS_SCROLLBARS_JUNCTION.
const STYLE_CLASS_SCROLLBARS_JUNCTION = "scrollbars-junction"

// STYLE_CLASS_SEPARATOR is a representation of the C constant GTK_STYLE_CLASS_SEPARATOR.
const STYLE_CLASS_SEPARATOR = "separator"

// STYLE_CLASS_SIDEBAR is a representation of the C constant GTK_STYLE_CLASS_SIDEBAR.
const STYLE_CLASS_SIDEBAR = "sidebar"

// STYLE_CLASS_SLIDER is a representation of the C constant GTK_STYLE_CLASS_SLIDER.
const STYLE_CLASS_SLIDER = "slider"

// STYLE_CLASS_SPINBUTTON is a representation of the C constant GTK_STYLE_CLASS_SPINBUTTON.
const STYLE_CLASS_SPINBUTTON = "spinbutton"

// STYLE_CLASS_SPINNER is a representation of the C constant GTK_STYLE_CLASS_SPINNER.
const STYLE_CLASS_SPINNER = "spinner"

// STYLE_CLASS_TITLEBAR is a representation of the C constant GTK_STYLE_CLASS_TITLEBAR.
const STYLE_CLASS_TITLEBAR = "titlebar"

// STYLE_CLASS_TOOLBAR is a representation of the C constant GTK_STYLE_CLASS_TOOLBAR.
const STYLE_CLASS_TOOLBAR = "toolbar"

// STYLE_CLASS_TOOLTIP is a representation of the C constant GTK_STYLE_CLASS_TOOLTIP.
const STYLE_CLASS_TOOLTIP = "tooltip"

// STYLE_CLASS_TOP is a representation of the C constant GTK_STYLE_CLASS_TOP.
const STYLE_CLASS_TOP = "top"

// STYLE_CLASS_TROUGH is a representation of the C constant GTK_STYLE_CLASS_TROUGH.
const STYLE_CLASS_TROUGH = "trough"

// STYLE_CLASS_VERTICAL is a representation of the C constant GTK_STYLE_CLASS_VERTICAL.
const STYLE_CLASS_VERTICAL = "vertical"

// STYLE_CLASS_VIEW is a representation of the C constant GTK_STYLE_CLASS_VIEW.
const STYLE_CLASS_VIEW = "view"

// STYLE_CLASS_WARNING is a representation of the C constant GTK_STYLE_CLASS_WARNING.
const STYLE_CLASS_WARNING = "warning"

// STYLE_PROPERTY_BACKGROUND_COLOR is a representation of the C constant GTK_STYLE_PROPERTY_BACKGROUND_COLOR.
const STYLE_PROPERTY_BACKGROUND_COLOR = "background-color"

// STYLE_PROPERTY_BACKGROUND_IMAGE is a representation of the C constant GTK_STYLE_PROPERTY_BACKGROUND_IMAGE.
const STYLE_PROPERTY_BACKGROUND_IMAGE = "background-image"

// STYLE_PROPERTY_BORDER_COLOR is a representation of the C constant GTK_STYLE_PROPERTY_BORDER_COLOR.
const STYLE_PROPERTY_BORDER_COLOR = "border-color"

// STYLE_PROPERTY_BORDER_RADIUS is a representation of the C constant GTK_STYLE_PROPERTY_BORDER_RADIUS.
const STYLE_PROPERTY_BORDER_RADIUS = "border-radius"

// STYLE_PROPERTY_BORDER_STYLE is a representation of the C constant GTK_STYLE_PROPERTY_BORDER_STYLE.
const STYLE_PROPERTY_BORDER_STYLE = "border-style"

// STYLE_PROPERTY_BORDER_WIDTH is a representation of the C constant GTK_STYLE_PROPERTY_BORDER_WIDTH.
const STYLE_PROPERTY_BORDER_WIDTH = "border-width"

// STYLE_PROPERTY_COLOR is a representation of the C constant GTK_STYLE_PROPERTY_COLOR.
const STYLE_PROPERTY_COLOR = "color"

// STYLE_PROPERTY_FONT is a representation of the C constant GTK_STYLE_PROPERTY_FONT.
const STYLE_PROPERTY_FONT = "font"

// STYLE_PROPERTY_MARGIN is a representation of the C constant GTK_STYLE_PROPERTY_MARGIN.
const STYLE_PROPERTY_MARGIN = "margin"

// STYLE_PROPERTY_PADDING is a representation of the C constant GTK_STYLE_PROPERTY_PADDING.
const STYLE_PROPERTY_PADDING = "padding"

// STYLE_PROVIDER_PRIORITY_APPLICATION is a representation of the C constant GTK_STYLE_PROVIDER_PRIORITY_APPLICATION.
const STYLE_PROVIDER_PRIORITY_APPLICATION = 600

// STYLE_PROVIDER_PRIORITY_FALLBACK is a representation of the C constant GTK_STYLE_PROVIDER_PRIORITY_FALLBACK.
const STYLE_PROVIDER_PRIORITY_FALLBACK = 1

// STYLE_PROVIDER_PRIORITY_SETTINGS is a representation of the C constant GTK_STYLE_PROVIDER_PRIORITY_SETTINGS.
const STYLE_PROVIDER_PRIORITY_SETTINGS = 400

// STYLE_PROVIDER_PRIORITY_THEME is a representation of the C constant GTK_STYLE_PROVIDER_PRIORITY_THEME.
const STYLE_PROVIDER_PRIORITY_THEME = 200

// STYLE_PROVIDER_PRIORITY_USER is a representation of the C constant GTK_STYLE_PROVIDER_PRIORITY_USER.
const STYLE_PROVIDER_PRIORITY_USER = 800

// STYLE_REGION_COLUMN is a representation of the C constant GTK_STYLE_REGION_COLUMN.
const STYLE_REGION_COLUMN = "column"

// STYLE_REGION_COLUMN_HEADER is a representation of the C constant GTK_STYLE_REGION_COLUMN_HEADER.
const STYLE_REGION_COLUMN_HEADER = "column-header"

// STYLE_REGION_ROW is a representation of the C constant GTK_STYLE_REGION_ROW.
const STYLE_REGION_ROW = "row"

// STYLE_REGION_TAB is a representation of the C constant GTK_STYLE_REGION_TAB.
const STYLE_REGION_TAB = "tab"

// TEXT_VIEW_PRIORITY_VALIDATE is a representation of the C constant GTK_TEXT_VIEW_PRIORITY_VALIDATE.
const TEXT_VIEW_PRIORITY_VALIDATE = 125

// TREE_SORTABLE_DEFAULT_SORT_COLUMN_ID is a representation of the C constant GTK_TREE_SORTABLE_DEFAULT_SORT_COLUMN_ID.
const TREE_SORTABLE_DEFAULT_SORT_COLUMN_ID = -1

// TREE_SORTABLE_UNSORTED_SORT_COLUMN_ID is a representation of the C constant GTK_TREE_SORTABLE_UNSORTED_SORT_COLUMN_ID.
const TREE_SORTABLE_UNSORTED_SORT_COLUMN_ID = -2

// AccelFlags is a representation of the C bitfield GtkAccelFlags.
type AccelFlags int

// AccelFlags_visible is a representation of the C bitfield member GTK_ACCEL_VISIBLE.
const AccelFlags_visible = AccelFlags(1)

// AccelFlags_locked is a representation of the C bitfield member GTK_ACCEL_LOCKED.
const AccelFlags_locked = AccelFlags(2)

// AccelFlags_mask is a representation of the C bitfield member GTK_ACCEL_MASK.
const AccelFlags_mask = AccelFlags(7)

// AttachOptions is a representation of the C bitfield GtkAttachOptions.
type AttachOptions int

// AttachOptions_expand is a representation of the C bitfield member GTK_EXPAND.
const AttachOptions_expand = AttachOptions(1)

// AttachOptions_shrink is a representation of the C bitfield member GTK_SHRINK.
const AttachOptions_shrink = AttachOptions(2)

// AttachOptions_fill is a representation of the C bitfield member GTK_FILL.
const AttachOptions_fill = AttachOptions(4)

// CalendarDisplayOptions is a representation of the C bitfield GtkCalendarDisplayOptions.
type CalendarDisplayOptions int

// CalendarDisplayOptions_show_heading is a representation of the C bitfield member GTK_CALENDAR_SHOW_HEADING.
const CalendarDisplayOptions_show_heading = CalendarDisplayOptions(1)

// CalendarDisplayOptions_show_day_names is a representation of the C bitfield member GTK_CALENDAR_SHOW_DAY_NAMES.
const CalendarDisplayOptions_show_day_names = CalendarDisplayOptions(2)

// CalendarDisplayOptions_no_month_change is a representation of the C bitfield member GTK_CALENDAR_NO_MONTH_CHANGE.
const CalendarDisplayOptions_no_month_change = CalendarDisplayOptions(4)

// CalendarDisplayOptions_show_week_numbers is a representation of the C bitfield member GTK_CALENDAR_SHOW_WEEK_NUMBERS.
const CalendarDisplayOptions_show_week_numbers = CalendarDisplayOptions(8)

// CalendarDisplayOptions_show_details is a representation of the C bitfield member GTK_CALENDAR_SHOW_DETAILS.
const CalendarDisplayOptions_show_details = CalendarDisplayOptions(32)

// CellRendererState is a representation of the C bitfield GtkCellRendererState.
type CellRendererState int

// CellRendererState_selected is a representation of the C bitfield member GTK_CELL_RENDERER_SELECTED.
const CellRendererState_selected = CellRendererState(1)

// CellRendererState_prelit is a representation of the C bitfield member GTK_CELL_RENDERER_PRELIT.
const CellRendererState_prelit = CellRendererState(2)

// CellRendererState_insensitive is a representation of the C bitfield member GTK_CELL_RENDERER_INSENSITIVE.
const CellRendererState_insensitive = CellRendererState(4)

// CellRendererState_sorted is a representation of the C bitfield member GTK_CELL_RENDERER_SORTED.
const CellRendererState_sorted = CellRendererState(8)

// CellRendererState_focused is a representation of the C bitfield member GTK_CELL_RENDERER_FOCUSED.
const CellRendererState_focused = CellRendererState(16)

// CellRendererState_expandable is a representation of the C bitfield member GTK_CELL_RENDERER_EXPANDABLE.
const CellRendererState_expandable = CellRendererState(32)

// CellRendererState_expanded is a representation of the C bitfield member GTK_CELL_RENDERER_EXPANDED.
const CellRendererState_expanded = CellRendererState(64)

// DebugFlag is a representation of the C bitfield GtkDebugFlag.
type DebugFlag int

// DebugFlag_misc is a representation of the C bitfield member GTK_DEBUG_MISC.
const DebugFlag_misc = DebugFlag(1)

// DebugFlag_plugsocket is a representation of the C bitfield member GTK_DEBUG_PLUGSOCKET.
const DebugFlag_plugsocket = DebugFlag(2)

// DebugFlag_text is a representation of the C bitfield member GTK_DEBUG_TEXT.
const DebugFlag_text = DebugFlag(4)

// DebugFlag_tree is a representation of the C bitfield member GTK_DEBUG_TREE.
const DebugFlag_tree = DebugFlag(8)

// DebugFlag_updates is a representation of the C bitfield member GTK_DEBUG_UPDATES.
const DebugFlag_updates = DebugFlag(16)

// DebugFlag_keybindings is a representation of the C bitfield member GTK_DEBUG_KEYBINDINGS.
const DebugFlag_keybindings = DebugFlag(32)

// DebugFlag_multihead is a representation of the C bitfield member GTK_DEBUG_MULTIHEAD.
const DebugFlag_multihead = DebugFlag(64)

// DebugFlag_modules is a representation of the C bitfield member GTK_DEBUG_MODULES.
const DebugFlag_modules = DebugFlag(128)

// DebugFlag_geometry is a representation of the C bitfield member GTK_DEBUG_GEOMETRY.
const DebugFlag_geometry = DebugFlag(256)

// DebugFlag_icontheme is a representation of the C bitfield member GTK_DEBUG_ICONTHEME.
const DebugFlag_icontheme = DebugFlag(512)

// DebugFlag_printing is a representation of the C bitfield member GTK_DEBUG_PRINTING.
const DebugFlag_printing = DebugFlag(1024)

// DebugFlag_builder is a representation of the C bitfield member GTK_DEBUG_BUILDER.
const DebugFlag_builder = DebugFlag(2048)

// DebugFlag_size_request is a representation of the C bitfield member GTK_DEBUG_SIZE_REQUEST.
const DebugFlag_size_request = DebugFlag(4096)

// DebugFlag_no_css_cache is a representation of the C bitfield member GTK_DEBUG_NO_CSS_CACHE.
const DebugFlag_no_css_cache = DebugFlag(8192)

// DebugFlag_baselines is a representation of the C bitfield member GTK_DEBUG_BASELINES.
const DebugFlag_baselines = DebugFlag(16384)

// DebugFlag_pixel_cache is a representation of the C bitfield member GTK_DEBUG_PIXEL_CACHE.
const DebugFlag_pixel_cache = DebugFlag(32768)

// DebugFlag_no_pixel_cache is a representation of the C bitfield member GTK_DEBUG_NO_PIXEL_CACHE.
const DebugFlag_no_pixel_cache = DebugFlag(65536)

// DebugFlag_interactive is a representation of the C bitfield member GTK_DEBUG_INTERACTIVE.
const DebugFlag_interactive = DebugFlag(131072)

// DebugFlag_touchscreen is a representation of the C bitfield member GTK_DEBUG_TOUCHSCREEN.
const DebugFlag_touchscreen = DebugFlag(262144)

// DebugFlag_actions is a representation of the C bitfield member GTK_DEBUG_ACTIONS.
const DebugFlag_actions = DebugFlag(524288)

// DebugFlag_resize is a representation of the C bitfield member GTK_DEBUG_RESIZE.
const DebugFlag_resize = DebugFlag(1048576)

// DebugFlag_layout is a representation of the C bitfield member GTK_DEBUG_LAYOUT.
const DebugFlag_layout = DebugFlag(2097152)

// DestDefaults is a representation of the C bitfield GtkDestDefaults.
type DestDefaults int

// DestDefaults_motion is a representation of the C bitfield member GTK_DEST_DEFAULT_MOTION.
const DestDefaults_motion = DestDefaults(1)

// DestDefaults_highlight is a representation of the C bitfield member GTK_DEST_DEFAULT_HIGHLIGHT.
const DestDefaults_highlight = DestDefaults(2)

// DestDefaults_drop is a representation of the C bitfield member GTK_DEST_DEFAULT_DROP.
const DestDefaults_drop = DestDefaults(4)

// DestDefaults_all is a representation of the C bitfield member GTK_DEST_DEFAULT_ALL.
const DestDefaults_all = DestDefaults(7)

// DialogFlags is a representation of the C bitfield GtkDialogFlags.
type DialogFlags int

// DialogFlags_modal is a representation of the C bitfield member GTK_DIALOG_MODAL.
const DialogFlags_modal = DialogFlags(1)

// DialogFlags_destroy_with_parent is a representation of the C bitfield member GTK_DIALOG_DESTROY_WITH_PARENT.
const DialogFlags_destroy_with_parent = DialogFlags(2)

// DialogFlags_use_header_bar is a representation of the C bitfield member GTK_DIALOG_USE_HEADER_BAR.
const DialogFlags_use_header_bar = DialogFlags(4)

// FileFilterFlags is a representation of the C bitfield GtkFileFilterFlags.
type FileFilterFlags int

// FileFilterFlags_filename is a representation of the C bitfield member GTK_FILE_FILTER_FILENAME.
const FileFilterFlags_filename = FileFilterFlags(1)

// FileFilterFlags_uri is a representation of the C bitfield member GTK_FILE_FILTER_URI.
const FileFilterFlags_uri = FileFilterFlags(2)

// FileFilterFlags_display_name is a representation of the C bitfield member GTK_FILE_FILTER_DISPLAY_NAME.
const FileFilterFlags_display_name = FileFilterFlags(4)

// FileFilterFlags_mime_type is a representation of the C bitfield member GTK_FILE_FILTER_MIME_TYPE.
const FileFilterFlags_mime_type = FileFilterFlags(8)

// IconLookupFlags is a representation of the C bitfield GtkIconLookupFlags.
type IconLookupFlags int

// IconLookupFlags_no_svg is a representation of the C bitfield member GTK_ICON_LOOKUP_NO_SVG.
const IconLookupFlags_no_svg = IconLookupFlags(1)

// IconLookupFlags_force_svg is a representation of the C bitfield member GTK_ICON_LOOKUP_FORCE_SVG.
const IconLookupFlags_force_svg = IconLookupFlags(2)

// IconLookupFlags_use_builtin is a representation of the C bitfield member GTK_ICON_LOOKUP_USE_BUILTIN.
const IconLookupFlags_use_builtin = IconLookupFlags(4)

// IconLookupFlags_generic_fallback is a representation of the C bitfield member GTK_ICON_LOOKUP_GENERIC_FALLBACK.
const IconLookupFlags_generic_fallback = IconLookupFlags(8)

// IconLookupFlags_force_size is a representation of the C bitfield member GTK_ICON_LOOKUP_FORCE_SIZE.
const IconLookupFlags_force_size = IconLookupFlags(16)

// IconLookupFlags_force_regular is a representation of the C bitfield member GTK_ICON_LOOKUP_FORCE_REGULAR.
const IconLookupFlags_force_regular = IconLookupFlags(32)

// IconLookupFlags_force_symbolic is a representation of the C bitfield member GTK_ICON_LOOKUP_FORCE_SYMBOLIC.
const IconLookupFlags_force_symbolic = IconLookupFlags(64)

// IconLookupFlags_dir_ltr is a representation of the C bitfield member GTK_ICON_LOOKUP_DIR_LTR.
const IconLookupFlags_dir_ltr = IconLookupFlags(128)

// IconLookupFlags_dir_rtl is a representation of the C bitfield member GTK_ICON_LOOKUP_DIR_RTL.
const IconLookupFlags_dir_rtl = IconLookupFlags(256)

// JunctionSides is a representation of the C bitfield GtkJunctionSides.
type JunctionSides int

// JunctionSides_none is a representation of the C bitfield member GTK_JUNCTION_NONE.
const JunctionSides_none = JunctionSides(0)

// JunctionSides_corner_topleft is a representation of the C bitfield member GTK_JUNCTION_CORNER_TOPLEFT.
const JunctionSides_corner_topleft = JunctionSides(1)

// JunctionSides_corner_topright is a representation of the C bitfield member GTK_JUNCTION_CORNER_TOPRIGHT.
const JunctionSides_corner_topright = JunctionSides(2)

// JunctionSides_corner_bottomleft is a representation of the C bitfield member GTK_JUNCTION_CORNER_BOTTOMLEFT.
const JunctionSides_corner_bottomleft = JunctionSides(4)

// JunctionSides_corner_bottomright is a representation of the C bitfield member GTK_JUNCTION_CORNER_BOTTOMRIGHT.
const JunctionSides_corner_bottomright = JunctionSides(8)

// JunctionSides_top is a representation of the C bitfield member GTK_JUNCTION_TOP.
const JunctionSides_top = JunctionSides(3)

// JunctionSides_bottom is a representation of the C bitfield member GTK_JUNCTION_BOTTOM.
const JunctionSides_bottom = JunctionSides(12)

// JunctionSides_left is a representation of the C bitfield member GTK_JUNCTION_LEFT.
const JunctionSides_left = JunctionSides(5)

// JunctionSides_right is a representation of the C bitfield member GTK_JUNCTION_RIGHT.
const JunctionSides_right = JunctionSides(10)

// PlacesOpenFlags is a representation of the C bitfield GtkPlacesOpenFlags.
type PlacesOpenFlags int

// PlacesOpenFlags_normal is a representation of the C bitfield member GTK_PLACES_OPEN_NORMAL.
const PlacesOpenFlags_normal = PlacesOpenFlags(1)

// PlacesOpenFlags_new_tab is a representation of the C bitfield member GTK_PLACES_OPEN_NEW_TAB.
const PlacesOpenFlags_new_tab = PlacesOpenFlags(2)

// PlacesOpenFlags_new_window is a representation of the C bitfield member GTK_PLACES_OPEN_NEW_WINDOW.
const PlacesOpenFlags_new_window = PlacesOpenFlags(4)

// RcFlags is a representation of the C bitfield GtkRcFlags.
type RcFlags int

// RcFlags_fg is a representation of the C bitfield member GTK_RC_FG.
const RcFlags_fg = RcFlags(1)

// RcFlags_bg is a representation of the C bitfield member GTK_RC_BG.
const RcFlags_bg = RcFlags(2)

// RcFlags_text is a representation of the C bitfield member GTK_RC_TEXT.
const RcFlags_text = RcFlags(4)

// RcFlags_base is a representation of the C bitfield member GTK_RC_BASE.
const RcFlags_base = RcFlags(8)

// RecentFilterFlags is a representation of the C bitfield GtkRecentFilterFlags.
type RecentFilterFlags int

// RecentFilterFlags_uri is a representation of the C bitfield member GTK_RECENT_FILTER_URI.
const RecentFilterFlags_uri = RecentFilterFlags(1)

// RecentFilterFlags_display_name is a representation of the C bitfield member GTK_RECENT_FILTER_DISPLAY_NAME.
const RecentFilterFlags_display_name = RecentFilterFlags(2)

// RecentFilterFlags_mime_type is a representation of the C bitfield member GTK_RECENT_FILTER_MIME_TYPE.
const RecentFilterFlags_mime_type = RecentFilterFlags(4)

// RecentFilterFlags_application is a representation of the C bitfield member GTK_RECENT_FILTER_APPLICATION.
const RecentFilterFlags_application = RecentFilterFlags(8)

// RecentFilterFlags_group is a representation of the C bitfield member GTK_RECENT_FILTER_GROUP.
const RecentFilterFlags_group = RecentFilterFlags(16)

// RecentFilterFlags_age is a representation of the C bitfield member GTK_RECENT_FILTER_AGE.
const RecentFilterFlags_age = RecentFilterFlags(32)

// RegionFlags is a representation of the C bitfield GtkRegionFlags.
type RegionFlags int

// RegionFlags_even is a representation of the C bitfield member GTK_REGION_EVEN.
const RegionFlags_even = RegionFlags(1)

// RegionFlags_odd is a representation of the C bitfield member GTK_REGION_ODD.
const RegionFlags_odd = RegionFlags(2)

// RegionFlags_first is a representation of the C bitfield member GTK_REGION_FIRST.
const RegionFlags_first = RegionFlags(4)

// RegionFlags_last is a representation of the C bitfield member GTK_REGION_LAST.
const RegionFlags_last = RegionFlags(8)

// RegionFlags_only is a representation of the C bitfield member GTK_REGION_ONLY.
const RegionFlags_only = RegionFlags(16)

// RegionFlags_sorted is a representation of the C bitfield member GTK_REGION_SORTED.
const RegionFlags_sorted = RegionFlags(32)

// StateFlags is a representation of the C bitfield GtkStateFlags.
type StateFlags int

// StateFlags_normal is a representation of the C bitfield member GTK_STATE_FLAG_NORMAL.
const StateFlags_normal = StateFlags(0)

// StateFlags_active is a representation of the C bitfield member GTK_STATE_FLAG_ACTIVE.
const StateFlags_active = StateFlags(1)

// StateFlags_prelight is a representation of the C bitfield member GTK_STATE_FLAG_PRELIGHT.
const StateFlags_prelight = StateFlags(2)

// StateFlags_selected is a representation of the C bitfield member GTK_STATE_FLAG_SELECTED.
const StateFlags_selected = StateFlags(4)

// StateFlags_insensitive is a representation of the C bitfield member GTK_STATE_FLAG_INSENSITIVE.
const StateFlags_insensitive = StateFlags(8)

// StateFlags_inconsistent is a representation of the C bitfield member GTK_STATE_FLAG_INCONSISTENT.
const StateFlags_inconsistent = StateFlags(16)

// StateFlags_focused is a representation of the C bitfield member GTK_STATE_FLAG_FOCUSED.
const StateFlags_focused = StateFlags(32)

// StateFlags_backdrop is a representation of the C bitfield member GTK_STATE_FLAG_BACKDROP.
const StateFlags_backdrop = StateFlags(64)

// StateFlags_dir_ltr is a representation of the C bitfield member GTK_STATE_FLAG_DIR_LTR.
const StateFlags_dir_ltr = StateFlags(128)

// StateFlags_dir_rtl is a representation of the C bitfield member GTK_STATE_FLAG_DIR_RTL.
const StateFlags_dir_rtl = StateFlags(256)

// StateFlags_link is a representation of the C bitfield member GTK_STATE_FLAG_LINK.
const StateFlags_link = StateFlags(512)

// StateFlags_visited is a representation of the C bitfield member GTK_STATE_FLAG_VISITED.
const StateFlags_visited = StateFlags(1024)

// StateFlags_checked is a representation of the C bitfield member GTK_STATE_FLAG_CHECKED.
const StateFlags_checked = StateFlags(2048)

// StateFlags_drop_active is a representation of the C bitfield member GTK_STATE_FLAG_DROP_ACTIVE.
const StateFlags_drop_active = StateFlags(4096)

// StyleContextPrintFlags is a representation of the C bitfield GtkStyleContextPrintFlags.
type StyleContextPrintFlags int

// StyleContextPrintFlags_none is a representation of the C bitfield member GTK_STYLE_CONTEXT_PRINT_NONE.
const StyleContextPrintFlags_none = StyleContextPrintFlags(0)

// StyleContextPrintFlags_recurse is a representation of the C bitfield member GTK_STYLE_CONTEXT_PRINT_RECURSE.
const StyleContextPrintFlags_recurse = StyleContextPrintFlags(1)

// StyleContextPrintFlags_show_style is a representation of the C bitfield member GTK_STYLE_CONTEXT_PRINT_SHOW_STYLE.
const StyleContextPrintFlags_show_style = StyleContextPrintFlags(2)

// TargetFlags is a representation of the C bitfield GtkTargetFlags.
type TargetFlags int

// TargetFlags_same_app is a representation of the C bitfield member GTK_TARGET_SAME_APP.
const TargetFlags_same_app = TargetFlags(1)

// TargetFlags_same_widget is a representation of the C bitfield member GTK_TARGET_SAME_WIDGET.
const TargetFlags_same_widget = TargetFlags(2)

// TargetFlags_other_app is a representation of the C bitfield member GTK_TARGET_OTHER_APP.
const TargetFlags_other_app = TargetFlags(4)

// TargetFlags_other_widget is a representation of the C bitfield member GTK_TARGET_OTHER_WIDGET.
const TargetFlags_other_widget = TargetFlags(8)

// TextSearchFlags is a representation of the C bitfield GtkTextSearchFlags.
type TextSearchFlags int

// TextSearchFlags_visible_only is a representation of the C bitfield member GTK_TEXT_SEARCH_VISIBLE_ONLY.
const TextSearchFlags_visible_only = TextSearchFlags(1)

// TextSearchFlags_text_only is a representation of the C bitfield member GTK_TEXT_SEARCH_TEXT_ONLY.
const TextSearchFlags_text_only = TextSearchFlags(2)

// TextSearchFlags_case_insensitive is a representation of the C bitfield member GTK_TEXT_SEARCH_CASE_INSENSITIVE.
const TextSearchFlags_case_insensitive = TextSearchFlags(4)

// ToolPaletteDragTargets is a representation of the C bitfield GtkToolPaletteDragTargets.
type ToolPaletteDragTargets int

// ToolPaletteDragTargets_items is a representation of the C bitfield member GTK_TOOL_PALETTE_DRAG_ITEMS.
const ToolPaletteDragTargets_items = ToolPaletteDragTargets(1)

// ToolPaletteDragTargets_groups is a representation of the C bitfield member GTK_TOOL_PALETTE_DRAG_GROUPS.
const ToolPaletteDragTargets_groups = ToolPaletteDragTargets(2)

// TreeModelFlags is a representation of the C bitfield GtkTreeModelFlags.
type TreeModelFlags int

// TreeModelFlags_iters_persist is a representation of the C bitfield member GTK_TREE_MODEL_ITERS_PERSIST.
const TreeModelFlags_iters_persist = TreeModelFlags(1)

// TreeModelFlags_list_only is a representation of the C bitfield member GTK_TREE_MODEL_LIST_ONLY.
const TreeModelFlags_list_only = TreeModelFlags(2)

// UIManagerItemType is a representation of the C bitfield GtkUIManagerItemType.
type UIManagerItemType int

// UIManagerItemType_auto is a representation of the C bitfield member GTK_UI_MANAGER_AUTO.
const UIManagerItemType_auto = UIManagerItemType(0)

// UIManagerItemType_menubar is a representation of the C bitfield member GTK_UI_MANAGER_MENUBAR.
const UIManagerItemType_menubar = UIManagerItemType(1)

// UIManagerItemType_menu is a representation of the C bitfield member GTK_UI_MANAGER_MENU.
const UIManagerItemType_menu = UIManagerItemType(2)

// UIManagerItemType_toolbar is a representation of the C bitfield member GTK_UI_MANAGER_TOOLBAR.
const UIManagerItemType_toolbar = UIManagerItemType(4)

// UIManagerItemType_placeholder is a representation of the C bitfield member GTK_UI_MANAGER_PLACEHOLDER.
const UIManagerItemType_placeholder = UIManagerItemType(8)

// UIManagerItemType_popup is a representation of the C bitfield member GTK_UI_MANAGER_POPUP.
const UIManagerItemType_popup = UIManagerItemType(16)

// UIManagerItemType_menuitem is a representation of the C bitfield member GTK_UI_MANAGER_MENUITEM.
const UIManagerItemType_menuitem = UIManagerItemType(32)

// UIManagerItemType_toolitem is a representation of the C bitfield member GTK_UI_MANAGER_TOOLITEM.
const UIManagerItemType_toolitem = UIManagerItemType(64)

// UIManagerItemType_separator is a representation of the C bitfield member GTK_UI_MANAGER_SEPARATOR.
const UIManagerItemType_separator = UIManagerItemType(128)

// UIManagerItemType_accelerator is a representation of the C bitfield member GTK_UI_MANAGER_ACCELERATOR.
const UIManagerItemType_accelerator = UIManagerItemType(256)

// UIManagerItemType_popup_with_accels is a representation of the C bitfield member GTK_UI_MANAGER_POPUP_WITH_ACCELS.
const UIManagerItemType_popup_with_accels = UIManagerItemType(512)

// Align is a representation of the C enumeration GtkAlign.
type Align int

// Align_fill is a representation of the C enumeration member GTK_ALIGN_FILL.
const Align_fill = Align(0)

// Align_start is a representation of the C enumeration member GTK_ALIGN_START.
const Align_start = Align(1)

// Align_end is a representation of the C enumeration member GTK_ALIGN_END.
const Align_end = Align(2)

// Align_center is a representation of the C enumeration member GTK_ALIGN_CENTER.
const Align_center = Align(3)

// Align_baseline is a representation of the C enumeration member GTK_ALIGN_BASELINE.
const Align_baseline = Align(4)

// ArrowPlacement is a representation of the C enumeration GtkArrowPlacement.
type ArrowPlacement int

// ArrowPlacement_both is a representation of the C enumeration member GTK_ARROWS_BOTH.
const ArrowPlacement_both = ArrowPlacement(0)

// ArrowPlacement_start is a representation of the C enumeration member GTK_ARROWS_START.
const ArrowPlacement_start = ArrowPlacement(1)

// ArrowPlacement_end is a representation of the C enumeration member GTK_ARROWS_END.
const ArrowPlacement_end = ArrowPlacement(2)

// ArrowType is a representation of the C enumeration GtkArrowType.
type ArrowType int

// ArrowType_up is a representation of the C enumeration member GTK_ARROW_UP.
const ArrowType_up = ArrowType(0)

// ArrowType_down is a representation of the C enumeration member GTK_ARROW_DOWN.
const ArrowType_down = ArrowType(1)

// ArrowType_left is a representation of the C enumeration member GTK_ARROW_LEFT.
const ArrowType_left = ArrowType(2)

// ArrowType_right is a representation of the C enumeration member GTK_ARROW_RIGHT.
const ArrowType_right = ArrowType(3)

// ArrowType_none is a representation of the C enumeration member GTK_ARROW_NONE.
const ArrowType_none = ArrowType(4)

// AssistantPageType is a representation of the C enumeration GtkAssistantPageType.
type AssistantPageType int

// AssistantPageType_content is a representation of the C enumeration member GTK_ASSISTANT_PAGE_CONTENT.
const AssistantPageType_content = AssistantPageType(0)

// AssistantPageType_intro is a representation of the C enumeration member GTK_ASSISTANT_PAGE_INTRO.
const AssistantPageType_intro = AssistantPageType(1)

// AssistantPageType_confirm is a representation of the C enumeration member GTK_ASSISTANT_PAGE_CONFIRM.
const AssistantPageType_confirm = AssistantPageType(2)

// AssistantPageType_summary is a representation of the C enumeration member GTK_ASSISTANT_PAGE_SUMMARY.
const AssistantPageType_summary = AssistantPageType(3)

// AssistantPageType_progress is a representation of the C enumeration member GTK_ASSISTANT_PAGE_PROGRESS.
const AssistantPageType_progress = AssistantPageType(4)

// AssistantPageType_custom is a representation of the C enumeration member GTK_ASSISTANT_PAGE_CUSTOM.
const AssistantPageType_custom = AssistantPageType(5)

// BorderStyle is a representation of the C enumeration GtkBorderStyle.
type BorderStyle int

// BorderStyle_none is a representation of the C enumeration member GTK_BORDER_STYLE_NONE.
const BorderStyle_none = BorderStyle(0)

// BorderStyle_solid is a representation of the C enumeration member GTK_BORDER_STYLE_SOLID.
const BorderStyle_solid = BorderStyle(1)

// BorderStyle_inset is a representation of the C enumeration member GTK_BORDER_STYLE_INSET.
const BorderStyle_inset = BorderStyle(2)

// BorderStyle_outset is a representation of the C enumeration member GTK_BORDER_STYLE_OUTSET.
const BorderStyle_outset = BorderStyle(3)

// BorderStyle_hidden is a representation of the C enumeration member GTK_BORDER_STYLE_HIDDEN.
const BorderStyle_hidden = BorderStyle(4)

// BorderStyle_dotted is a representation of the C enumeration member GTK_BORDER_STYLE_DOTTED.
const BorderStyle_dotted = BorderStyle(5)

// BorderStyle_dashed is a representation of the C enumeration member GTK_BORDER_STYLE_DASHED.
const BorderStyle_dashed = BorderStyle(6)

// BorderStyle_double is a representation of the C enumeration member GTK_BORDER_STYLE_DOUBLE.
const BorderStyle_double = BorderStyle(7)

// BorderStyle_groove is a representation of the C enumeration member GTK_BORDER_STYLE_GROOVE.
const BorderStyle_groove = BorderStyle(8)

// BorderStyle_ridge is a representation of the C enumeration member GTK_BORDER_STYLE_RIDGE.
const BorderStyle_ridge = BorderStyle(9)

// BuilderError is a representation of the C enumeration GtkBuilderError.
type BuilderError int

// BuilderError_invalid_type_function is a representation of the C enumeration member GTK_BUILDER_ERROR_INVALID_TYPE_FUNCTION.
const BuilderError_invalid_type_function = BuilderError(0)

// BuilderError_unhandled_tag is a representation of the C enumeration member GTK_BUILDER_ERROR_UNHANDLED_TAG.
const BuilderError_unhandled_tag = BuilderError(1)

// BuilderError_missing_attribute is a representation of the C enumeration member GTK_BUILDER_ERROR_MISSING_ATTRIBUTE.
const BuilderError_missing_attribute = BuilderError(2)

// BuilderError_invalid_attribute is a representation of the C enumeration member GTK_BUILDER_ERROR_INVALID_ATTRIBUTE.
const BuilderError_invalid_attribute = BuilderError(3)

// BuilderError_invalid_tag is a representation of the C enumeration member GTK_BUILDER_ERROR_INVALID_TAG.
const BuilderError_invalid_tag = BuilderError(4)

// BuilderError_missing_property_value is a representation of the C enumeration member GTK_BUILDER_ERROR_MISSING_PROPERTY_VALUE.
const BuilderError_missing_property_value = BuilderError(5)

// BuilderError_invalid_value is a representation of the C enumeration member GTK_BUILDER_ERROR_INVALID_VALUE.
const BuilderError_invalid_value = BuilderError(6)

// BuilderError_version_mismatch is a representation of the C enumeration member GTK_BUILDER_ERROR_VERSION_MISMATCH.
const BuilderError_version_mismatch = BuilderError(7)

// BuilderError_duplicate_id is a representation of the C enumeration member GTK_BUILDER_ERROR_DUPLICATE_ID.
const BuilderError_duplicate_id = BuilderError(8)

// BuilderError_object_type_refused is a representation of the C enumeration member GTK_BUILDER_ERROR_OBJECT_TYPE_REFUSED.
const BuilderError_object_type_refused = BuilderError(9)

// BuilderError_template_mismatch is a representation of the C enumeration member GTK_BUILDER_ERROR_TEMPLATE_MISMATCH.
const BuilderError_template_mismatch = BuilderError(10)

// BuilderError_invalid_property is a representation of the C enumeration member GTK_BUILDER_ERROR_INVALID_PROPERTY.
const BuilderError_invalid_property = BuilderError(11)

// BuilderError_invalid_signal is a representation of the C enumeration member GTK_BUILDER_ERROR_INVALID_SIGNAL.
const BuilderError_invalid_signal = BuilderError(12)

// BuilderError_invalid_id is a representation of the C enumeration member GTK_BUILDER_ERROR_INVALID_ID.
const BuilderError_invalid_id = BuilderError(13)

// ButtonBoxStyle is a representation of the C enumeration GtkButtonBoxStyle.
type ButtonBoxStyle int

// ButtonBoxStyle_spread is a representation of the C enumeration member GTK_BUTTONBOX_SPREAD.
const ButtonBoxStyle_spread = ButtonBoxStyle(1)

// ButtonBoxStyle_edge is a representation of the C enumeration member GTK_BUTTONBOX_EDGE.
const ButtonBoxStyle_edge = ButtonBoxStyle(2)

// ButtonBoxStyle_start is a representation of the C enumeration member GTK_BUTTONBOX_START.
const ButtonBoxStyle_start = ButtonBoxStyle(3)

// ButtonBoxStyle_end is a representation of the C enumeration member GTK_BUTTONBOX_END.
const ButtonBoxStyle_end = ButtonBoxStyle(4)

// ButtonBoxStyle_center is a representation of the C enumeration member GTK_BUTTONBOX_CENTER.
const ButtonBoxStyle_center = ButtonBoxStyle(5)

// ButtonBoxStyle_expand is a representation of the C enumeration member GTK_BUTTONBOX_EXPAND.
const ButtonBoxStyle_expand = ButtonBoxStyle(6)

// ButtonRole is a representation of the C enumeration GtkButtonRole.
type ButtonRole int

// ButtonRole_normal is a representation of the C enumeration member GTK_BUTTON_ROLE_NORMAL.
const ButtonRole_normal = ButtonRole(0)

// ButtonRole_check is a representation of the C enumeration member GTK_BUTTON_ROLE_CHECK.
const ButtonRole_check = ButtonRole(1)

// ButtonRole_radio is a representation of the C enumeration member GTK_BUTTON_ROLE_RADIO.
const ButtonRole_radio = ButtonRole(2)

// ButtonsType is a representation of the C enumeration GtkButtonsType.
type ButtonsType int

// ButtonsType_none is a representation of the C enumeration member GTK_BUTTONS_NONE.
const ButtonsType_none = ButtonsType(0)

// ButtonsType_ok is a representation of the C enumeration member GTK_BUTTONS_OK.
const ButtonsType_ok = ButtonsType(1)

// ButtonsType_close is a representation of the C enumeration member GTK_BUTTONS_CLOSE.
const ButtonsType_close = ButtonsType(2)

// ButtonsType_cancel is a representation of the C enumeration member GTK_BUTTONS_CANCEL.
const ButtonsType_cancel = ButtonsType(3)

// ButtonsType_yes_no is a representation of the C enumeration member GTK_BUTTONS_YES_NO.
const ButtonsType_yes_no = ButtonsType(4)

// ButtonsType_ok_cancel is a representation of the C enumeration member GTK_BUTTONS_OK_CANCEL.
const ButtonsType_ok_cancel = ButtonsType(5)

// CellRendererAccelMode is a representation of the C enumeration GtkCellRendererAccelMode.
type CellRendererAccelMode int

// CellRendererAccelMode_gtk is a representation of the C enumeration member GTK_CELL_RENDERER_ACCEL_MODE_GTK.
const CellRendererAccelMode_gtk = CellRendererAccelMode(0)

// CellRendererAccelMode_other is a representation of the C enumeration member GTK_CELL_RENDERER_ACCEL_MODE_OTHER.
const CellRendererAccelMode_other = CellRendererAccelMode(1)

// CellRendererAccelMode_modifier_tap is a representation of the C enumeration member GTK_CELL_RENDERER_ACCEL_MODE_MODIFIER_TAP.
const CellRendererAccelMode_modifier_tap = CellRendererAccelMode(2)

// CellRendererMode is a representation of the C enumeration GtkCellRendererMode.
type CellRendererMode int

// CellRendererMode_inert is a representation of the C enumeration member GTK_CELL_RENDERER_MODE_INERT.
const CellRendererMode_inert = CellRendererMode(0)

// CellRendererMode_activatable is a representation of the C enumeration member GTK_CELL_RENDERER_MODE_ACTIVATABLE.
const CellRendererMode_activatable = CellRendererMode(1)

// CellRendererMode_editable is a representation of the C enumeration member GTK_CELL_RENDERER_MODE_EDITABLE.
const CellRendererMode_editable = CellRendererMode(2)

// CornerType is a representation of the C enumeration GtkCornerType.
type CornerType int

// CornerType_top_left is a representation of the C enumeration member GTK_CORNER_TOP_LEFT.
const CornerType_top_left = CornerType(0)

// CornerType_bottom_left is a representation of the C enumeration member GTK_CORNER_BOTTOM_LEFT.
const CornerType_bottom_left = CornerType(1)

// CornerType_top_right is a representation of the C enumeration member GTK_CORNER_TOP_RIGHT.
const CornerType_top_right = CornerType(2)

// CornerType_bottom_right is a representation of the C enumeration member GTK_CORNER_BOTTOM_RIGHT.
const CornerType_bottom_right = CornerType(3)

// CssProviderError is a representation of the C enumeration GtkCssProviderError.
type CssProviderError int

// CssProviderError_failed is a representation of the C enumeration member GTK_CSS_PROVIDER_ERROR_FAILED.
const CssProviderError_failed = CssProviderError(0)

// CssProviderError_syntax is a representation of the C enumeration member GTK_CSS_PROVIDER_ERROR_SYNTAX.
const CssProviderError_syntax = CssProviderError(1)

// CssProviderError_import is a representation of the C enumeration member GTK_CSS_PROVIDER_ERROR_IMPORT.
const CssProviderError_import = CssProviderError(2)

// CssProviderError_name is a representation of the C enumeration member GTK_CSS_PROVIDER_ERROR_NAME.
const CssProviderError_name = CssProviderError(3)

// CssProviderError_deprecated is a representation of the C enumeration member GTK_CSS_PROVIDER_ERROR_DEPRECATED.
const CssProviderError_deprecated = CssProviderError(4)

// CssProviderError_unknown_value is a representation of the C enumeration member GTK_CSS_PROVIDER_ERROR_UNKNOWN_VALUE.
const CssProviderError_unknown_value = CssProviderError(5)

// DeleteType is a representation of the C enumeration GtkDeleteType.
type DeleteType int

// DeleteType_chars is a representation of the C enumeration member GTK_DELETE_CHARS.
const DeleteType_chars = DeleteType(0)

// DeleteType_word_ends is a representation of the C enumeration member GTK_DELETE_WORD_ENDS.
const DeleteType_word_ends = DeleteType(1)

// DeleteType_words is a representation of the C enumeration member GTK_DELETE_WORDS.
const DeleteType_words = DeleteType(2)

// DeleteType_display_lines is a representation of the C enumeration member GTK_DELETE_DISPLAY_LINES.
const DeleteType_display_lines = DeleteType(3)

// DeleteType_display_line_ends is a representation of the C enumeration member GTK_DELETE_DISPLAY_LINE_ENDS.
const DeleteType_display_line_ends = DeleteType(4)

// DeleteType_paragraph_ends is a representation of the C enumeration member GTK_DELETE_PARAGRAPH_ENDS.
const DeleteType_paragraph_ends = DeleteType(5)

// DeleteType_paragraphs is a representation of the C enumeration member GTK_DELETE_PARAGRAPHS.
const DeleteType_paragraphs = DeleteType(6)

// DeleteType_whitespace is a representation of the C enumeration member GTK_DELETE_WHITESPACE.
const DeleteType_whitespace = DeleteType(7)

// DirectionType is a representation of the C enumeration GtkDirectionType.
type DirectionType int

// DirectionType_tab_forward is a representation of the C enumeration member GTK_DIR_TAB_FORWARD.
const DirectionType_tab_forward = DirectionType(0)

// DirectionType_tab_backward is a representation of the C enumeration member GTK_DIR_TAB_BACKWARD.
const DirectionType_tab_backward = DirectionType(1)

// DirectionType_up is a representation of the C enumeration member GTK_DIR_UP.
const DirectionType_up = DirectionType(2)

// DirectionType_down is a representation of the C enumeration member GTK_DIR_DOWN.
const DirectionType_down = DirectionType(3)

// DirectionType_left is a representation of the C enumeration member GTK_DIR_LEFT.
const DirectionType_left = DirectionType(4)

// DirectionType_right is a representation of the C enumeration member GTK_DIR_RIGHT.
const DirectionType_right = DirectionType(5)

// DragResult is a representation of the C enumeration GtkDragResult.
type DragResult int

// DragResult_success is a representation of the C enumeration member GTK_DRAG_RESULT_SUCCESS.
const DragResult_success = DragResult(0)

// DragResult_no_target is a representation of the C enumeration member GTK_DRAG_RESULT_NO_TARGET.
const DragResult_no_target = DragResult(1)

// DragResult_user_cancelled is a representation of the C enumeration member GTK_DRAG_RESULT_USER_CANCELLED.
const DragResult_user_cancelled = DragResult(2)

// DragResult_timeout_expired is a representation of the C enumeration member GTK_DRAG_RESULT_TIMEOUT_EXPIRED.
const DragResult_timeout_expired = DragResult(3)

// DragResult_grab_broken is a representation of the C enumeration member GTK_DRAG_RESULT_GRAB_BROKEN.
const DragResult_grab_broken = DragResult(4)

// DragResult_error is a representation of the C enumeration member GTK_DRAG_RESULT_ERROR.
const DragResult_error = DragResult(5)

// ExpanderStyle is a representation of the C enumeration GtkExpanderStyle.
type ExpanderStyle int

// ExpanderStyle_collapsed is a representation of the C enumeration member GTK_EXPANDER_COLLAPSED.
const ExpanderStyle_collapsed = ExpanderStyle(0)

// ExpanderStyle_semi_collapsed is a representation of the C enumeration member GTK_EXPANDER_SEMI_COLLAPSED.
const ExpanderStyle_semi_collapsed = ExpanderStyle(1)

// ExpanderStyle_semi_expanded is a representation of the C enumeration member GTK_EXPANDER_SEMI_EXPANDED.
const ExpanderStyle_semi_expanded = ExpanderStyle(2)

// ExpanderStyle_expanded is a representation of the C enumeration member GTK_EXPANDER_EXPANDED.
const ExpanderStyle_expanded = ExpanderStyle(3)

// FileChooserAction is a representation of the C enumeration GtkFileChooserAction.
type FileChooserAction int

// FileChooserAction_open is a representation of the C enumeration member GTK_FILE_CHOOSER_ACTION_OPEN.
const FileChooserAction_open = FileChooserAction(0)

// FileChooserAction_save is a representation of the C enumeration member GTK_FILE_CHOOSER_ACTION_SAVE.
const FileChooserAction_save = FileChooserAction(1)

// FileChooserAction_select_folder is a representation of the C enumeration member GTK_FILE_CHOOSER_ACTION_SELECT_FOLDER.
const FileChooserAction_select_folder = FileChooserAction(2)

// FileChooserAction_create_folder is a representation of the C enumeration member GTK_FILE_CHOOSER_ACTION_CREATE_FOLDER.
const FileChooserAction_create_folder = FileChooserAction(3)

// FileChooserError is a representation of the C enumeration GtkFileChooserError.
type FileChooserError int

// FileChooserError_nonexistent is a representation of the C enumeration member GTK_FILE_CHOOSER_ERROR_NONEXISTENT.
const FileChooserError_nonexistent = FileChooserError(0)

// FileChooserError_bad_filename is a representation of the C enumeration member GTK_FILE_CHOOSER_ERROR_BAD_FILENAME.
const FileChooserError_bad_filename = FileChooserError(1)

// FileChooserError_already_exists is a representation of the C enumeration member GTK_FILE_CHOOSER_ERROR_ALREADY_EXISTS.
const FileChooserError_already_exists = FileChooserError(2)

// FileChooserError_incomplete_hostname is a representation of the C enumeration member GTK_FILE_CHOOSER_ERROR_INCOMPLETE_HOSTNAME.
const FileChooserError_incomplete_hostname = FileChooserError(3)

// IMPreeditStyle is a representation of the C enumeration GtkIMPreeditStyle.
type IMPreeditStyle int

// IMPreeditStyle_nothing is a representation of the C enumeration member GTK_IM_PREEDIT_NOTHING.
const IMPreeditStyle_nothing = IMPreeditStyle(0)

// IMPreeditStyle_callback is a representation of the C enumeration member GTK_IM_PREEDIT_CALLBACK.
const IMPreeditStyle_callback = IMPreeditStyle(1)

// IMPreeditStyle_none is a representation of the C enumeration member GTK_IM_PREEDIT_NONE.
const IMPreeditStyle_none = IMPreeditStyle(2)

// IMStatusStyle is a representation of the C enumeration GtkIMStatusStyle.
type IMStatusStyle int

// IMStatusStyle_nothing is a representation of the C enumeration member GTK_IM_STATUS_NOTHING.
const IMStatusStyle_nothing = IMStatusStyle(0)

// IMStatusStyle_callback is a representation of the C enumeration member GTK_IM_STATUS_CALLBACK.
const IMStatusStyle_callback = IMStatusStyle(1)

// IMStatusStyle_none is a representation of the C enumeration member GTK_IM_STATUS_NONE.
const IMStatusStyle_none = IMStatusStyle(2)

// IconSize is a representation of the C enumeration GtkIconSize.
type IconSize int

// IconSize_invalid is a representation of the C enumeration member GTK_ICON_SIZE_INVALID.
const IconSize_invalid = IconSize(0)

// IconSize_menu is a representation of the C enumeration member GTK_ICON_SIZE_MENU.
const IconSize_menu = IconSize(1)

// IconSize_small_toolbar is a representation of the C enumeration member GTK_ICON_SIZE_SMALL_TOOLBAR.
const IconSize_small_toolbar = IconSize(2)

// IconSize_large_toolbar is a representation of the C enumeration member GTK_ICON_SIZE_LARGE_TOOLBAR.
const IconSize_large_toolbar = IconSize(3)

// IconSize_button is a representation of the C enumeration member GTK_ICON_SIZE_BUTTON.
const IconSize_button = IconSize(4)

// IconSize_dnd is a representation of the C enumeration member GTK_ICON_SIZE_DND.
const IconSize_dnd = IconSize(5)

// IconSize_dialog is a representation of the C enumeration member GTK_ICON_SIZE_DIALOG.
const IconSize_dialog = IconSize(6)

// IconThemeError is a representation of the C enumeration GtkIconThemeError.
type IconThemeError int

// IconThemeError_not_found is a representation of the C enumeration member GTK_ICON_THEME_NOT_FOUND.
const IconThemeError_not_found = IconThemeError(0)

// IconThemeError_failed is a representation of the C enumeration member GTK_ICON_THEME_FAILED.
const IconThemeError_failed = IconThemeError(1)

// IconViewDropPosition is a representation of the C enumeration GtkIconViewDropPosition.
type IconViewDropPosition int

// IconViewDropPosition_no_drop is a representation of the C enumeration member GTK_ICON_VIEW_NO_DROP.
const IconViewDropPosition_no_drop = IconViewDropPosition(0)

// IconViewDropPosition_drop_into is a representation of the C enumeration member GTK_ICON_VIEW_DROP_INTO.
const IconViewDropPosition_drop_into = IconViewDropPosition(1)

// IconViewDropPosition_drop_left is a representation of the C enumeration member GTK_ICON_VIEW_DROP_LEFT.
const IconViewDropPosition_drop_left = IconViewDropPosition(2)

// IconViewDropPosition_drop_right is a representation of the C enumeration member GTK_ICON_VIEW_DROP_RIGHT.
const IconViewDropPosition_drop_right = IconViewDropPosition(3)

// IconViewDropPosition_drop_above is a representation of the C enumeration member GTK_ICON_VIEW_DROP_ABOVE.
const IconViewDropPosition_drop_above = IconViewDropPosition(4)

// IconViewDropPosition_drop_below is a representation of the C enumeration member GTK_ICON_VIEW_DROP_BELOW.
const IconViewDropPosition_drop_below = IconViewDropPosition(5)

// ImageType is a representation of the C enumeration GtkImageType.
type ImageType int

// ImageType_empty is a representation of the C enumeration member GTK_IMAGE_EMPTY.
const ImageType_empty = ImageType(0)

// ImageType_pixbuf is a representation of the C enumeration member GTK_IMAGE_PIXBUF.
const ImageType_pixbuf = ImageType(1)

// ImageType_stock is a representation of the C enumeration member GTK_IMAGE_STOCK.
const ImageType_stock = ImageType(2)

// ImageType_icon_set is a representation of the C enumeration member GTK_IMAGE_ICON_SET.
const ImageType_icon_set = ImageType(3)

// ImageType_animation is a representation of the C enumeration member GTK_IMAGE_ANIMATION.
const ImageType_animation = ImageType(4)

// ImageType_icon_name is a representation of the C enumeration member GTK_IMAGE_ICON_NAME.
const ImageType_icon_name = ImageType(5)

// ImageType_gicon is a representation of the C enumeration member GTK_IMAGE_GICON.
const ImageType_gicon = ImageType(6)

// ImageType_surface is a representation of the C enumeration member GTK_IMAGE_SURFACE.
const ImageType_surface = ImageType(7)

// Justification is a representation of the C enumeration GtkJustification.
type Justification int

// Justification_left is a representation of the C enumeration member GTK_JUSTIFY_LEFT.
const Justification_left = Justification(0)

// Justification_right is a representation of the C enumeration member GTK_JUSTIFY_RIGHT.
const Justification_right = Justification(1)

// Justification_center is a representation of the C enumeration member GTK_JUSTIFY_CENTER.
const Justification_center = Justification(2)

// Justification_fill is a representation of the C enumeration member GTK_JUSTIFY_FILL.
const Justification_fill = Justification(3)

// MenuDirectionType is a representation of the C enumeration GtkMenuDirectionType.
type MenuDirectionType int

// MenuDirectionType_parent is a representation of the C enumeration member GTK_MENU_DIR_PARENT.
const MenuDirectionType_parent = MenuDirectionType(0)

// MenuDirectionType_child is a representation of the C enumeration member GTK_MENU_DIR_CHILD.
const MenuDirectionType_child = MenuDirectionType(1)

// MenuDirectionType_next is a representation of the C enumeration member GTK_MENU_DIR_NEXT.
const MenuDirectionType_next = MenuDirectionType(2)

// MenuDirectionType_prev is a representation of the C enumeration member GTK_MENU_DIR_PREV.
const MenuDirectionType_prev = MenuDirectionType(3)

// MessageType is a representation of the C enumeration GtkMessageType.
type MessageType int

// MessageType_info is a representation of the C enumeration member GTK_MESSAGE_INFO.
const MessageType_info = MessageType(0)

// MessageType_warning is a representation of the C enumeration member GTK_MESSAGE_WARNING.
const MessageType_warning = MessageType(1)

// MessageType_question is a representation of the C enumeration member GTK_MESSAGE_QUESTION.
const MessageType_question = MessageType(2)

// MessageType_error is a representation of the C enumeration member GTK_MESSAGE_ERROR.
const MessageType_error = MessageType(3)

// MessageType_other is a representation of the C enumeration member GTK_MESSAGE_OTHER.
const MessageType_other = MessageType(4)

// MovementStep is a representation of the C enumeration GtkMovementStep.
type MovementStep int

// MovementStep_logical_positions is a representation of the C enumeration member GTK_MOVEMENT_LOGICAL_POSITIONS.
const MovementStep_logical_positions = MovementStep(0)

// MovementStep_visual_positions is a representation of the C enumeration member GTK_MOVEMENT_VISUAL_POSITIONS.
const MovementStep_visual_positions = MovementStep(1)

// MovementStep_words is a representation of the C enumeration member GTK_MOVEMENT_WORDS.
const MovementStep_words = MovementStep(2)

// MovementStep_display_lines is a representation of the C enumeration member GTK_MOVEMENT_DISPLAY_LINES.
const MovementStep_display_lines = MovementStep(3)

// MovementStep_display_line_ends is a representation of the C enumeration member GTK_MOVEMENT_DISPLAY_LINE_ENDS.
const MovementStep_display_line_ends = MovementStep(4)

// MovementStep_paragraphs is a representation of the C enumeration member GTK_MOVEMENT_PARAGRAPHS.
const MovementStep_paragraphs = MovementStep(5)

// MovementStep_paragraph_ends is a representation of the C enumeration member GTK_MOVEMENT_PARAGRAPH_ENDS.
const MovementStep_paragraph_ends = MovementStep(6)

// MovementStep_pages is a representation of the C enumeration member GTK_MOVEMENT_PAGES.
const MovementStep_pages = MovementStep(7)

// MovementStep_buffer_ends is a representation of the C enumeration member GTK_MOVEMENT_BUFFER_ENDS.
const MovementStep_buffer_ends = MovementStep(8)

// MovementStep_horizontal_pages is a representation of the C enumeration member GTK_MOVEMENT_HORIZONTAL_PAGES.
const MovementStep_horizontal_pages = MovementStep(9)

// NotebookTab is a representation of the C enumeration GtkNotebookTab.
type NotebookTab int

// NotebookTab_first is a representation of the C enumeration member GTK_NOTEBOOK_TAB_FIRST.
const NotebookTab_first = NotebookTab(0)

// NotebookTab_last is a representation of the C enumeration member GTK_NOTEBOOK_TAB_LAST.
const NotebookTab_last = NotebookTab(1)

// NumberUpLayout is a representation of the C enumeration GtkNumberUpLayout.
type NumberUpLayout int

// NumberUpLayout_lrtb is a representation of the C enumeration member GTK_NUMBER_UP_LAYOUT_LEFT_TO_RIGHT_TOP_TO_BOTTOM.
const NumberUpLayout_lrtb = NumberUpLayout(0)

// NumberUpLayout_lrbt is a representation of the C enumeration member GTK_NUMBER_UP_LAYOUT_LEFT_TO_RIGHT_BOTTOM_TO_TOP.
const NumberUpLayout_lrbt = NumberUpLayout(1)

// NumberUpLayout_rltb is a representation of the C enumeration member GTK_NUMBER_UP_LAYOUT_RIGHT_TO_LEFT_TOP_TO_BOTTOM.
const NumberUpLayout_rltb = NumberUpLayout(2)

// NumberUpLayout_rlbt is a representation of the C enumeration member GTK_NUMBER_UP_LAYOUT_RIGHT_TO_LEFT_BOTTOM_TO_TOP.
const NumberUpLayout_rlbt = NumberUpLayout(3)

// NumberUpLayout_tblr is a representation of the C enumeration member GTK_NUMBER_UP_LAYOUT_TOP_TO_BOTTOM_LEFT_TO_RIGHT.
const NumberUpLayout_tblr = NumberUpLayout(4)

// NumberUpLayout_tbrl is a representation of the C enumeration member GTK_NUMBER_UP_LAYOUT_TOP_TO_BOTTOM_RIGHT_TO_LEFT.
const NumberUpLayout_tbrl = NumberUpLayout(5)

// NumberUpLayout_btlr is a representation of the C enumeration member GTK_NUMBER_UP_LAYOUT_BOTTOM_TO_TOP_LEFT_TO_RIGHT.
const NumberUpLayout_btlr = NumberUpLayout(6)

// NumberUpLayout_btrl is a representation of the C enumeration member GTK_NUMBER_UP_LAYOUT_BOTTOM_TO_TOP_RIGHT_TO_LEFT.
const NumberUpLayout_btrl = NumberUpLayout(7)

// Orientation is a representation of the C enumeration GtkOrientation.
type Orientation int

// Orientation_horizontal is a representation of the C enumeration member GTK_ORIENTATION_HORIZONTAL.
const Orientation_horizontal = Orientation(0)

// Orientation_vertical is a representation of the C enumeration member GTK_ORIENTATION_VERTICAL.
const Orientation_vertical = Orientation(1)

// PackDirection is a representation of the C enumeration GtkPackDirection.
type PackDirection int

// PackDirection_ltr is a representation of the C enumeration member GTK_PACK_DIRECTION_LTR.
const PackDirection_ltr = PackDirection(0)

// PackDirection_rtl is a representation of the C enumeration member GTK_PACK_DIRECTION_RTL.
const PackDirection_rtl = PackDirection(1)

// PackDirection_ttb is a representation of the C enumeration member GTK_PACK_DIRECTION_TTB.
const PackDirection_ttb = PackDirection(2)

// PackDirection_btt is a representation of the C enumeration member GTK_PACK_DIRECTION_BTT.
const PackDirection_btt = PackDirection(3)

// PackType is a representation of the C enumeration GtkPackType.
type PackType int

// PackType_start is a representation of the C enumeration member GTK_PACK_START.
const PackType_start = PackType(0)

// PackType_end is a representation of the C enumeration member GTK_PACK_END.
const PackType_end = PackType(1)

// PadActionType is a representation of the C enumeration GtkPadActionType.
type PadActionType int

// PadActionType_button is a representation of the C enumeration member GTK_PAD_ACTION_BUTTON.
const PadActionType_button = PadActionType(0)

// PadActionType_ring is a representation of the C enumeration member GTK_PAD_ACTION_RING.
const PadActionType_ring = PadActionType(1)

// PadActionType_strip is a representation of the C enumeration member GTK_PAD_ACTION_STRIP.
const PadActionType_strip = PadActionType(2)

// PageOrientation is a representation of the C enumeration GtkPageOrientation.
type PageOrientation int

// PageOrientation_portrait is a representation of the C enumeration member GTK_PAGE_ORIENTATION_PORTRAIT.
const PageOrientation_portrait = PageOrientation(0)

// PageOrientation_landscape is a representation of the C enumeration member GTK_PAGE_ORIENTATION_LANDSCAPE.
const PageOrientation_landscape = PageOrientation(1)

// PageOrientation_reverse_portrait is a representation of the C enumeration member GTK_PAGE_ORIENTATION_REVERSE_PORTRAIT.
const PageOrientation_reverse_portrait = PageOrientation(2)

// PageOrientation_reverse_landscape is a representation of the C enumeration member GTK_PAGE_ORIENTATION_REVERSE_LANDSCAPE.
const PageOrientation_reverse_landscape = PageOrientation(3)

// PageSet is a representation of the C enumeration GtkPageSet.
type PageSet int

// PageSet_all is a representation of the C enumeration member GTK_PAGE_SET_ALL.
const PageSet_all = PageSet(0)

// PageSet_even is a representation of the C enumeration member GTK_PAGE_SET_EVEN.
const PageSet_even = PageSet(1)

// PageSet_odd is a representation of the C enumeration member GTK_PAGE_SET_ODD.
const PageSet_odd = PageSet(2)

// PathPriorityType is a representation of the C enumeration GtkPathPriorityType.
type PathPriorityType int

// PathPriorityType_lowest is a representation of the C enumeration member GTK_PATH_PRIO_LOWEST.
const PathPriorityType_lowest = PathPriorityType(0)

// PathPriorityType_gtk is a representation of the C enumeration member GTK_PATH_PRIO_GTK.
const PathPriorityType_gtk = PathPriorityType(4)

// PathPriorityType_application is a representation of the C enumeration member GTK_PATH_PRIO_APPLICATION.
const PathPriorityType_application = PathPriorityType(8)

// PathPriorityType_theme is a representation of the C enumeration member GTK_PATH_PRIO_THEME.
const PathPriorityType_theme = PathPriorityType(10)

// PathPriorityType_rc is a representation of the C enumeration member GTK_PATH_PRIO_RC.
const PathPriorityType_rc = PathPriorityType(12)

// PathPriorityType_highest is a representation of the C enumeration member GTK_PATH_PRIO_HIGHEST.
const PathPriorityType_highest = PathPriorityType(15)

// PathType is a representation of the C enumeration GtkPathType.
type PathType int

// PathType_widget is a representation of the C enumeration member GTK_PATH_WIDGET.
const PathType_widget = PathType(0)

// PathType_widget_class is a representation of the C enumeration member GTK_PATH_WIDGET_CLASS.
const PathType_widget_class = PathType(1)

// PathType_class is a representation of the C enumeration member GTK_PATH_CLASS.
const PathType_class = PathType(2)

// PolicyType is a representation of the C enumeration GtkPolicyType.
type PolicyType int

// PolicyType_always is a representation of the C enumeration member GTK_POLICY_ALWAYS.
const PolicyType_always = PolicyType(0)

// PolicyType_automatic is a representation of the C enumeration member GTK_POLICY_AUTOMATIC.
const PolicyType_automatic = PolicyType(1)

// PolicyType_never is a representation of the C enumeration member GTK_POLICY_NEVER.
const PolicyType_never = PolicyType(2)

// PolicyType_external is a representation of the C enumeration member GTK_POLICY_EXTERNAL.
const PolicyType_external = PolicyType(3)

// PositionType is a representation of the C enumeration GtkPositionType.
type PositionType int

// PositionType_left is a representation of the C enumeration member GTK_POS_LEFT.
const PositionType_left = PositionType(0)

// PositionType_right is a representation of the C enumeration member GTK_POS_RIGHT.
const PositionType_right = PositionType(1)

// PositionType_top is a representation of the C enumeration member GTK_POS_TOP.
const PositionType_top = PositionType(2)

// PositionType_bottom is a representation of the C enumeration member GTK_POS_BOTTOM.
const PositionType_bottom = PositionType(3)

// PrintDuplex is a representation of the C enumeration GtkPrintDuplex.
type PrintDuplex int

// PrintDuplex_simplex is a representation of the C enumeration member GTK_PRINT_DUPLEX_SIMPLEX.
const PrintDuplex_simplex = PrintDuplex(0)

// PrintDuplex_horizontal is a representation of the C enumeration member GTK_PRINT_DUPLEX_HORIZONTAL.
const PrintDuplex_horizontal = PrintDuplex(1)

// PrintDuplex_vertical is a representation of the C enumeration member GTK_PRINT_DUPLEX_VERTICAL.
const PrintDuplex_vertical = PrintDuplex(2)

// PrintError is a representation of the C enumeration GtkPrintError.
type PrintError int

// PrintError_general is a representation of the C enumeration member GTK_PRINT_ERROR_GENERAL.
const PrintError_general = PrintError(0)

// PrintError_internal_error is a representation of the C enumeration member GTK_PRINT_ERROR_INTERNAL_ERROR.
const PrintError_internal_error = PrintError(1)

// PrintError_nomem is a representation of the C enumeration member GTK_PRINT_ERROR_NOMEM.
const PrintError_nomem = PrintError(2)

// PrintError_invalid_file is a representation of the C enumeration member GTK_PRINT_ERROR_INVALID_FILE.
const PrintError_invalid_file = PrintError(3)

// PrintOperationAction is a representation of the C enumeration GtkPrintOperationAction.
type PrintOperationAction int

// PrintOperationAction_print_dialog is a representation of the C enumeration member GTK_PRINT_OPERATION_ACTION_PRINT_DIALOG.
const PrintOperationAction_print_dialog = PrintOperationAction(0)

// PrintOperationAction_print is a representation of the C enumeration member GTK_PRINT_OPERATION_ACTION_PRINT.
const PrintOperationAction_print = PrintOperationAction(1)

// PrintOperationAction_preview is a representation of the C enumeration member GTK_PRINT_OPERATION_ACTION_PREVIEW.
const PrintOperationAction_preview = PrintOperationAction(2)

// PrintOperationAction_export is a representation of the C enumeration member GTK_PRINT_OPERATION_ACTION_EXPORT.
const PrintOperationAction_export = PrintOperationAction(3)

// PrintOperationResult is a representation of the C enumeration GtkPrintOperationResult.
type PrintOperationResult int

// PrintOperationResult_error is a representation of the C enumeration member GTK_PRINT_OPERATION_RESULT_ERROR.
const PrintOperationResult_error = PrintOperationResult(0)

// PrintOperationResult_apply is a representation of the C enumeration member GTK_PRINT_OPERATION_RESULT_APPLY.
const PrintOperationResult_apply = PrintOperationResult(1)

// PrintOperationResult_cancel is a representation of the C enumeration member GTK_PRINT_OPERATION_RESULT_CANCEL.
const PrintOperationResult_cancel = PrintOperationResult(2)

// PrintOperationResult_in_progress is a representation of the C enumeration member GTK_PRINT_OPERATION_RESULT_IN_PROGRESS.
const PrintOperationResult_in_progress = PrintOperationResult(3)

// PrintPages is a representation of the C enumeration GtkPrintPages.
type PrintPages int

// PrintPages_all is a representation of the C enumeration member GTK_PRINT_PAGES_ALL.
const PrintPages_all = PrintPages(0)

// PrintPages_current is a representation of the C enumeration member GTK_PRINT_PAGES_CURRENT.
const PrintPages_current = PrintPages(1)

// PrintPages_ranges is a representation of the C enumeration member GTK_PRINT_PAGES_RANGES.
const PrintPages_ranges = PrintPages(2)

// PrintPages_selection is a representation of the C enumeration member GTK_PRINT_PAGES_SELECTION.
const PrintPages_selection = PrintPages(3)

// PrintQuality is a representation of the C enumeration GtkPrintQuality.
type PrintQuality int

// PrintQuality_low is a representation of the C enumeration member GTK_PRINT_QUALITY_LOW.
const PrintQuality_low = PrintQuality(0)

// PrintQuality_normal is a representation of the C enumeration member GTK_PRINT_QUALITY_NORMAL.
const PrintQuality_normal = PrintQuality(1)

// PrintQuality_high is a representation of the C enumeration member GTK_PRINT_QUALITY_HIGH.
const PrintQuality_high = PrintQuality(2)

// PrintQuality_draft is a representation of the C enumeration member GTK_PRINT_QUALITY_DRAFT.
const PrintQuality_draft = PrintQuality(3)

// PrintStatus is a representation of the C enumeration GtkPrintStatus.
type PrintStatus int

// PrintStatus_initial is a representation of the C enumeration member GTK_PRINT_STATUS_INITIAL.
const PrintStatus_initial = PrintStatus(0)

// PrintStatus_preparing is a representation of the C enumeration member GTK_PRINT_STATUS_PREPARING.
const PrintStatus_preparing = PrintStatus(1)

// PrintStatus_generating_data is a representation of the C enumeration member GTK_PRINT_STATUS_GENERATING_DATA.
const PrintStatus_generating_data = PrintStatus(2)

// PrintStatus_sending_data is a representation of the C enumeration member GTK_PRINT_STATUS_SENDING_DATA.
const PrintStatus_sending_data = PrintStatus(3)

// PrintStatus_pending is a representation of the C enumeration member GTK_PRINT_STATUS_PENDING.
const PrintStatus_pending = PrintStatus(4)

// PrintStatus_pending_issue is a representation of the C enumeration member GTK_PRINT_STATUS_PENDING_ISSUE.
const PrintStatus_pending_issue = PrintStatus(5)

// PrintStatus_printing is a representation of the C enumeration member GTK_PRINT_STATUS_PRINTING.
const PrintStatus_printing = PrintStatus(6)

// PrintStatus_finished is a representation of the C enumeration member GTK_PRINT_STATUS_FINISHED.
const PrintStatus_finished = PrintStatus(7)

// PrintStatus_finished_aborted is a representation of the C enumeration member GTK_PRINT_STATUS_FINISHED_ABORTED.
const PrintStatus_finished_aborted = PrintStatus(8)

// RcTokenType is a representation of the C enumeration GtkRcTokenType.
type RcTokenType int

// RcTokenType_invalid is a representation of the C enumeration member GTK_RC_TOKEN_INVALID.
const RcTokenType_invalid = RcTokenType(270)

// RcTokenType_include is a representation of the C enumeration member GTK_RC_TOKEN_INCLUDE.
const RcTokenType_include = RcTokenType(271)

// RcTokenType_normal is a representation of the C enumeration member GTK_RC_TOKEN_NORMAL.
const RcTokenType_normal = RcTokenType(272)

// RcTokenType_active is a representation of the C enumeration member GTK_RC_TOKEN_ACTIVE.
const RcTokenType_active = RcTokenType(273)

// RcTokenType_prelight is a representation of the C enumeration member GTK_RC_TOKEN_PRELIGHT.
const RcTokenType_prelight = RcTokenType(274)

// RcTokenType_selected is a representation of the C enumeration member GTK_RC_TOKEN_SELECTED.
const RcTokenType_selected = RcTokenType(275)

// RcTokenType_insensitive is a representation of the C enumeration member GTK_RC_TOKEN_INSENSITIVE.
const RcTokenType_insensitive = RcTokenType(276)

// RcTokenType_fg is a representation of the C enumeration member GTK_RC_TOKEN_FG.
const RcTokenType_fg = RcTokenType(277)

// RcTokenType_bg is a representation of the C enumeration member GTK_RC_TOKEN_BG.
const RcTokenType_bg = RcTokenType(278)

// RcTokenType_text is a representation of the C enumeration member GTK_RC_TOKEN_TEXT.
const RcTokenType_text = RcTokenType(279)

// RcTokenType_base is a representation of the C enumeration member GTK_RC_TOKEN_BASE.
const RcTokenType_base = RcTokenType(280)

// RcTokenType_xthickness is a representation of the C enumeration member GTK_RC_TOKEN_XTHICKNESS.
const RcTokenType_xthickness = RcTokenType(281)

// RcTokenType_ythickness is a representation of the C enumeration member GTK_RC_TOKEN_YTHICKNESS.
const RcTokenType_ythickness = RcTokenType(282)

// RcTokenType_font is a representation of the C enumeration member GTK_RC_TOKEN_FONT.
const RcTokenType_font = RcTokenType(283)

// RcTokenType_fontset is a representation of the C enumeration member GTK_RC_TOKEN_FONTSET.
const RcTokenType_fontset = RcTokenType(284)

// RcTokenType_font_name is a representation of the C enumeration member GTK_RC_TOKEN_FONT_NAME.
const RcTokenType_font_name = RcTokenType(285)

// RcTokenType_bg_pixmap is a representation of the C enumeration member GTK_RC_TOKEN_BG_PIXMAP.
const RcTokenType_bg_pixmap = RcTokenType(286)

// RcTokenType_pixmap_path is a representation of the C enumeration member GTK_RC_TOKEN_PIXMAP_PATH.
const RcTokenType_pixmap_path = RcTokenType(287)

// RcTokenType_style is a representation of the C enumeration member GTK_RC_TOKEN_STYLE.
const RcTokenType_style = RcTokenType(288)

// RcTokenType_binding is a representation of the C enumeration member GTK_RC_TOKEN_BINDING.
const RcTokenType_binding = RcTokenType(289)

// RcTokenType_bind is a representation of the C enumeration member GTK_RC_TOKEN_BIND.
const RcTokenType_bind = RcTokenType(290)

// RcTokenType_widget is a representation of the C enumeration member GTK_RC_TOKEN_WIDGET.
const RcTokenType_widget = RcTokenType(291)

// RcTokenType_widget_class is a representation of the C enumeration member GTK_RC_TOKEN_WIDGET_CLASS.
const RcTokenType_widget_class = RcTokenType(292)

// RcTokenType_class is a representation of the C enumeration member GTK_RC_TOKEN_CLASS.
const RcTokenType_class = RcTokenType(293)

// RcTokenType_lowest is a representation of the C enumeration member GTK_RC_TOKEN_LOWEST.
const RcTokenType_lowest = RcTokenType(294)

// RcTokenType_gtk is a representation of the C enumeration member GTK_RC_TOKEN_GTK.
const RcTokenType_gtk = RcTokenType(295)

// RcTokenType_application is a representation of the C enumeration member GTK_RC_TOKEN_APPLICATION.
const RcTokenType_application = RcTokenType(296)

// RcTokenType_theme is a representation of the C enumeration member GTK_RC_TOKEN_THEME.
const RcTokenType_theme = RcTokenType(297)

// RcTokenType_rc is a representation of the C enumeration member GTK_RC_TOKEN_RC.
const RcTokenType_rc = RcTokenType(298)

// RcTokenType_highest is a representation of the C enumeration member GTK_RC_TOKEN_HIGHEST.
const RcTokenType_highest = RcTokenType(299)

// RcTokenType_engine is a representation of the C enumeration member GTK_RC_TOKEN_ENGINE.
const RcTokenType_engine = RcTokenType(300)

// RcTokenType_module_path is a representation of the C enumeration member GTK_RC_TOKEN_MODULE_PATH.
const RcTokenType_module_path = RcTokenType(301)

// RcTokenType_im_module_path is a representation of the C enumeration member GTK_RC_TOKEN_IM_MODULE_PATH.
const RcTokenType_im_module_path = RcTokenType(302)

// RcTokenType_im_module_file is a representation of the C enumeration member GTK_RC_TOKEN_IM_MODULE_FILE.
const RcTokenType_im_module_file = RcTokenType(303)

// RcTokenType_stock is a representation of the C enumeration member GTK_RC_TOKEN_STOCK.
const RcTokenType_stock = RcTokenType(304)

// RcTokenType_ltr is a representation of the C enumeration member GTK_RC_TOKEN_LTR.
const RcTokenType_ltr = RcTokenType(305)

// RcTokenType_rtl is a representation of the C enumeration member GTK_RC_TOKEN_RTL.
const RcTokenType_rtl = RcTokenType(306)

// RcTokenType_color is a representation of the C enumeration member GTK_RC_TOKEN_COLOR.
const RcTokenType_color = RcTokenType(307)

// RcTokenType_unbind is a representation of the C enumeration member GTK_RC_TOKEN_UNBIND.
const RcTokenType_unbind = RcTokenType(308)

// RcTokenType_last is a representation of the C enumeration member GTK_RC_TOKEN_LAST.
const RcTokenType_last = RcTokenType(309)

// ReliefStyle is a representation of the C enumeration GtkReliefStyle.
type ReliefStyle int

// ReliefStyle_normal is a representation of the C enumeration member GTK_RELIEF_NORMAL.
const ReliefStyle_normal = ReliefStyle(0)

// ReliefStyle_half is a representation of the C enumeration member GTK_RELIEF_HALF.
const ReliefStyle_half = ReliefStyle(1)

// ReliefStyle_none is a representation of the C enumeration member GTK_RELIEF_NONE.
const ReliefStyle_none = ReliefStyle(2)

// ResizeMode is a representation of the C enumeration GtkResizeMode.
type ResizeMode int

// ResizeMode_parent is a representation of the C enumeration member GTK_RESIZE_PARENT.
const ResizeMode_parent = ResizeMode(0)

// ResizeMode_queue is a representation of the C enumeration member GTK_RESIZE_QUEUE.
const ResizeMode_queue = ResizeMode(1)

// ResizeMode_immediate is a representation of the C enumeration member GTK_RESIZE_IMMEDIATE.
const ResizeMode_immediate = ResizeMode(2)

// ResponseType is a representation of the C enumeration GtkResponseType.
type ResponseType int

// ResponseType_none is a representation of the C enumeration member GTK_RESPONSE_NONE.
const ResponseType_none = ResponseType(-1)

// ResponseType_reject is a representation of the C enumeration member GTK_RESPONSE_REJECT.
const ResponseType_reject = ResponseType(-2)

// ResponseType_accept is a representation of the C enumeration member GTK_RESPONSE_ACCEPT.
const ResponseType_accept = ResponseType(-3)

// ResponseType_delete_event is a representation of the C enumeration member GTK_RESPONSE_DELETE_EVENT.
const ResponseType_delete_event = ResponseType(-4)

// ResponseType_ok is a representation of the C enumeration member GTK_RESPONSE_OK.
const ResponseType_ok = ResponseType(-5)

// ResponseType_cancel is a representation of the C enumeration member GTK_RESPONSE_CANCEL.
const ResponseType_cancel = ResponseType(-6)

// ResponseType_close is a representation of the C enumeration member GTK_RESPONSE_CLOSE.
const ResponseType_close = ResponseType(-7)

// ResponseType_yes is a representation of the C enumeration member GTK_RESPONSE_YES.
const ResponseType_yes = ResponseType(-8)

// ResponseType_no is a representation of the C enumeration member GTK_RESPONSE_NO.
const ResponseType_no = ResponseType(-9)

// ResponseType_apply is a representation of the C enumeration member GTK_RESPONSE_APPLY.
const ResponseType_apply = ResponseType(-10)

// ResponseType_help is a representation of the C enumeration member GTK_RESPONSE_HELP.
const ResponseType_help = ResponseType(-11)

// RevealerTransitionType is a representation of the C enumeration GtkRevealerTransitionType.
type RevealerTransitionType int

// RevealerTransitionType_none is a representation of the C enumeration member GTK_REVEALER_TRANSITION_TYPE_NONE.
const RevealerTransitionType_none = RevealerTransitionType(0)

// RevealerTransitionType_crossfade is a representation of the C enumeration member GTK_REVEALER_TRANSITION_TYPE_CROSSFADE.
const RevealerTransitionType_crossfade = RevealerTransitionType(1)

// RevealerTransitionType_slide_right is a representation of the C enumeration member GTK_REVEALER_TRANSITION_TYPE_SLIDE_RIGHT.
const RevealerTransitionType_slide_right = RevealerTransitionType(2)

// RevealerTransitionType_slide_left is a representation of the C enumeration member GTK_REVEALER_TRANSITION_TYPE_SLIDE_LEFT.
const RevealerTransitionType_slide_left = RevealerTransitionType(3)

// RevealerTransitionType_slide_up is a representation of the C enumeration member GTK_REVEALER_TRANSITION_TYPE_SLIDE_UP.
const RevealerTransitionType_slide_up = RevealerTransitionType(4)

// RevealerTransitionType_slide_down is a representation of the C enumeration member GTK_REVEALER_TRANSITION_TYPE_SLIDE_DOWN.
const RevealerTransitionType_slide_down = RevealerTransitionType(5)

// ScrollStep is a representation of the C enumeration GtkScrollStep.
type ScrollStep int

// ScrollStep_steps is a representation of the C enumeration member GTK_SCROLL_STEPS.
const ScrollStep_steps = ScrollStep(0)

// ScrollStep_pages is a representation of the C enumeration member GTK_SCROLL_PAGES.
const ScrollStep_pages = ScrollStep(1)

// ScrollStep_ends is a representation of the C enumeration member GTK_SCROLL_ENDS.
const ScrollStep_ends = ScrollStep(2)

// ScrollStep_horizontal_steps is a representation of the C enumeration member GTK_SCROLL_HORIZONTAL_STEPS.
const ScrollStep_horizontal_steps = ScrollStep(3)

// ScrollStep_horizontal_pages is a representation of the C enumeration member GTK_SCROLL_HORIZONTAL_PAGES.
const ScrollStep_horizontal_pages = ScrollStep(4)

// ScrollStep_horizontal_ends is a representation of the C enumeration member GTK_SCROLL_HORIZONTAL_ENDS.
const ScrollStep_horizontal_ends = ScrollStep(5)

// ScrollType is a representation of the C enumeration GtkScrollType.
type ScrollType int

// ScrollType_none is a representation of the C enumeration member GTK_SCROLL_NONE.
const ScrollType_none = ScrollType(0)

// ScrollType_jump is a representation of the C enumeration member GTK_SCROLL_JUMP.
const ScrollType_jump = ScrollType(1)

// ScrollType_step_backward is a representation of the C enumeration member GTK_SCROLL_STEP_BACKWARD.
const ScrollType_step_backward = ScrollType(2)

// ScrollType_step_forward is a representation of the C enumeration member GTK_SCROLL_STEP_FORWARD.
const ScrollType_step_forward = ScrollType(3)

// ScrollType_page_backward is a representation of the C enumeration member GTK_SCROLL_PAGE_BACKWARD.
const ScrollType_page_backward = ScrollType(4)

// ScrollType_page_forward is a representation of the C enumeration member GTK_SCROLL_PAGE_FORWARD.
const ScrollType_page_forward = ScrollType(5)

// ScrollType_step_up is a representation of the C enumeration member GTK_SCROLL_STEP_UP.
const ScrollType_step_up = ScrollType(6)

// ScrollType_step_down is a representation of the C enumeration member GTK_SCROLL_STEP_DOWN.
const ScrollType_step_down = ScrollType(7)

// ScrollType_page_up is a representation of the C enumeration member GTK_SCROLL_PAGE_UP.
const ScrollType_page_up = ScrollType(8)

// ScrollType_page_down is a representation of the C enumeration member GTK_SCROLL_PAGE_DOWN.
const ScrollType_page_down = ScrollType(9)

// ScrollType_step_left is a representation of the C enumeration member GTK_SCROLL_STEP_LEFT.
const ScrollType_step_left = ScrollType(10)

// ScrollType_step_right is a representation of the C enumeration member GTK_SCROLL_STEP_RIGHT.
const ScrollType_step_right = ScrollType(11)

// ScrollType_page_left is a representation of the C enumeration member GTK_SCROLL_PAGE_LEFT.
const ScrollType_page_left = ScrollType(12)

// ScrollType_page_right is a representation of the C enumeration member GTK_SCROLL_PAGE_RIGHT.
const ScrollType_page_right = ScrollType(13)

// ScrollType_start is a representation of the C enumeration member GTK_SCROLL_START.
const ScrollType_start = ScrollType(14)

// ScrollType_end is a representation of the C enumeration member GTK_SCROLL_END.
const ScrollType_end = ScrollType(15)

// ScrollablePolicy is a representation of the C enumeration GtkScrollablePolicy.
type ScrollablePolicy int

// ScrollablePolicy_minimum is a representation of the C enumeration member GTK_SCROLL_MINIMUM.
const ScrollablePolicy_minimum = ScrollablePolicy(0)

// ScrollablePolicy_natural is a representation of the C enumeration member GTK_SCROLL_NATURAL.
const ScrollablePolicy_natural = ScrollablePolicy(1)

// SelectionMode is a representation of the C enumeration GtkSelectionMode.
type SelectionMode int

// SelectionMode_none is a representation of the C enumeration member GTK_SELECTION_NONE.
const SelectionMode_none = SelectionMode(0)

// SelectionMode_single is a representation of the C enumeration member GTK_SELECTION_SINGLE.
const SelectionMode_single = SelectionMode(1)

// SelectionMode_browse is a representation of the C enumeration member GTK_SELECTION_BROWSE.
const SelectionMode_browse = SelectionMode(2)

// SelectionMode_multiple is a representation of the C enumeration member GTK_SELECTION_MULTIPLE.
const SelectionMode_multiple = SelectionMode(3)

// SensitivityType is a representation of the C enumeration GtkSensitivityType.
type SensitivityType int

// SensitivityType_auto is a representation of the C enumeration member GTK_SENSITIVITY_AUTO.
const SensitivityType_auto = SensitivityType(0)

// SensitivityType_on is a representation of the C enumeration member GTK_SENSITIVITY_ON.
const SensitivityType_on = SensitivityType(1)

// SensitivityType_off is a representation of the C enumeration member GTK_SENSITIVITY_OFF.
const SensitivityType_off = SensitivityType(2)

// ShadowType is a representation of the C enumeration GtkShadowType.
type ShadowType int

// ShadowType_none is a representation of the C enumeration member GTK_SHADOW_NONE.
const ShadowType_none = ShadowType(0)

// ShadowType_in is a representation of the C enumeration member GTK_SHADOW_IN.
const ShadowType_in = ShadowType(1)

// ShadowType_out is a representation of the C enumeration member GTK_SHADOW_OUT.
const ShadowType_out = ShadowType(2)

// ShadowType_etched_in is a representation of the C enumeration member GTK_SHADOW_ETCHED_IN.
const ShadowType_etched_in = ShadowType(3)

// ShadowType_etched_out is a representation of the C enumeration member GTK_SHADOW_ETCHED_OUT.
const ShadowType_etched_out = ShadowType(4)

// SizeGroupMode is a representation of the C enumeration GtkSizeGroupMode.
type SizeGroupMode int

// SizeGroupMode_none is a representation of the C enumeration member GTK_SIZE_GROUP_NONE.
const SizeGroupMode_none = SizeGroupMode(0)

// SizeGroupMode_horizontal is a representation of the C enumeration member GTK_SIZE_GROUP_HORIZONTAL.
const SizeGroupMode_horizontal = SizeGroupMode(1)

// SizeGroupMode_vertical is a representation of the C enumeration member GTK_SIZE_GROUP_VERTICAL.
const SizeGroupMode_vertical = SizeGroupMode(2)

// SizeGroupMode_both is a representation of the C enumeration member GTK_SIZE_GROUP_BOTH.
const SizeGroupMode_both = SizeGroupMode(3)

// SizeRequestMode is a representation of the C enumeration GtkSizeRequestMode.
type SizeRequestMode int

// SizeRequestMode_height_for_width is a representation of the C enumeration member GTK_SIZE_REQUEST_HEIGHT_FOR_WIDTH.
const SizeRequestMode_height_for_width = SizeRequestMode(0)

// SizeRequestMode_width_for_height is a representation of the C enumeration member GTK_SIZE_REQUEST_WIDTH_FOR_HEIGHT.
const SizeRequestMode_width_for_height = SizeRequestMode(1)

// SizeRequestMode_constant_size is a representation of the C enumeration member GTK_SIZE_REQUEST_CONSTANT_SIZE.
const SizeRequestMode_constant_size = SizeRequestMode(2)

// SortType is a representation of the C enumeration GtkSortType.
type SortType int

// SortType_ascending is a representation of the C enumeration member GTK_SORT_ASCENDING.
const SortType_ascending = SortType(0)

// SortType_descending is a representation of the C enumeration member GTK_SORT_DESCENDING.
const SortType_descending = SortType(1)

// SpinButtonUpdatePolicy is a representation of the C enumeration GtkSpinButtonUpdatePolicy.
type SpinButtonUpdatePolicy int

// SpinButtonUpdatePolicy_always is a representation of the C enumeration member GTK_UPDATE_ALWAYS.
const SpinButtonUpdatePolicy_always = SpinButtonUpdatePolicy(0)

// SpinButtonUpdatePolicy_if_valid is a representation of the C enumeration member GTK_UPDATE_IF_VALID.
const SpinButtonUpdatePolicy_if_valid = SpinButtonUpdatePolicy(1)

// SpinType is a representation of the C enumeration GtkSpinType.
type SpinType int

// SpinType_step_forward is a representation of the C enumeration member GTK_SPIN_STEP_FORWARD.
const SpinType_step_forward = SpinType(0)

// SpinType_step_backward is a representation of the C enumeration member GTK_SPIN_STEP_BACKWARD.
const SpinType_step_backward = SpinType(1)

// SpinType_page_forward is a representation of the C enumeration member GTK_SPIN_PAGE_FORWARD.
const SpinType_page_forward = SpinType(2)

// SpinType_page_backward is a representation of the C enumeration member GTK_SPIN_PAGE_BACKWARD.
const SpinType_page_backward = SpinType(3)

// SpinType_home is a representation of the C enumeration member GTK_SPIN_HOME.
const SpinType_home = SpinType(4)

// SpinType_end is a representation of the C enumeration member GTK_SPIN_END.
const SpinType_end = SpinType(5)

// SpinType_user_defined is a representation of the C enumeration member GTK_SPIN_USER_DEFINED.
const SpinType_user_defined = SpinType(6)

// StackTransitionType is a representation of the C enumeration GtkStackTransitionType.
type StackTransitionType int

// StackTransitionType_none is a representation of the C enumeration member GTK_STACK_TRANSITION_TYPE_NONE.
const StackTransitionType_none = StackTransitionType(0)

// StackTransitionType_crossfade is a representation of the C enumeration member GTK_STACK_TRANSITION_TYPE_CROSSFADE.
const StackTransitionType_crossfade = StackTransitionType(1)

// StackTransitionType_slide_right is a representation of the C enumeration member GTK_STACK_TRANSITION_TYPE_SLIDE_RIGHT.
const StackTransitionType_slide_right = StackTransitionType(2)

// StackTransitionType_slide_left is a representation of the C enumeration member GTK_STACK_TRANSITION_TYPE_SLIDE_LEFT.
const StackTransitionType_slide_left = StackTransitionType(3)

// StackTransitionType_slide_up is a representation of the C enumeration member GTK_STACK_TRANSITION_TYPE_SLIDE_UP.
const StackTransitionType_slide_up = StackTransitionType(4)

// StackTransitionType_slide_down is a representation of the C enumeration member GTK_STACK_TRANSITION_TYPE_SLIDE_DOWN.
const StackTransitionType_slide_down = StackTransitionType(5)

// StackTransitionType_slide_left_right is a representation of the C enumeration member GTK_STACK_TRANSITION_TYPE_SLIDE_LEFT_RIGHT.
const StackTransitionType_slide_left_right = StackTransitionType(6)

// StackTransitionType_slide_up_down is a representation of the C enumeration member GTK_STACK_TRANSITION_TYPE_SLIDE_UP_DOWN.
const StackTransitionType_slide_up_down = StackTransitionType(7)

// StackTransitionType_over_up is a representation of the C enumeration member GTK_STACK_TRANSITION_TYPE_OVER_UP.
const StackTransitionType_over_up = StackTransitionType(8)

// StackTransitionType_over_down is a representation of the C enumeration member GTK_STACK_TRANSITION_TYPE_OVER_DOWN.
const StackTransitionType_over_down = StackTransitionType(9)

// StackTransitionType_over_left is a representation of the C enumeration member GTK_STACK_TRANSITION_TYPE_OVER_LEFT.
const StackTransitionType_over_left = StackTransitionType(10)

// StackTransitionType_over_right is a representation of the C enumeration member GTK_STACK_TRANSITION_TYPE_OVER_RIGHT.
const StackTransitionType_over_right = StackTransitionType(11)

// StackTransitionType_under_up is a representation of the C enumeration member GTK_STACK_TRANSITION_TYPE_UNDER_UP.
const StackTransitionType_under_up = StackTransitionType(12)

// StackTransitionType_under_down is a representation of the C enumeration member GTK_STACK_TRANSITION_TYPE_UNDER_DOWN.
const StackTransitionType_under_down = StackTransitionType(13)

// StackTransitionType_under_left is a representation of the C enumeration member GTK_STACK_TRANSITION_TYPE_UNDER_LEFT.
const StackTransitionType_under_left = StackTransitionType(14)

// StackTransitionType_under_right is a representation of the C enumeration member GTK_STACK_TRANSITION_TYPE_UNDER_RIGHT.
const StackTransitionType_under_right = StackTransitionType(15)

// StackTransitionType_over_up_down is a representation of the C enumeration member GTK_STACK_TRANSITION_TYPE_OVER_UP_DOWN.
const StackTransitionType_over_up_down = StackTransitionType(16)

// StackTransitionType_over_down_up is a representation of the C enumeration member GTK_STACK_TRANSITION_TYPE_OVER_DOWN_UP.
const StackTransitionType_over_down_up = StackTransitionType(17)

// StackTransitionType_over_left_right is a representation of the C enumeration member GTK_STACK_TRANSITION_TYPE_OVER_LEFT_RIGHT.
const StackTransitionType_over_left_right = StackTransitionType(18)

// StackTransitionType_over_right_left is a representation of the C enumeration member GTK_STACK_TRANSITION_TYPE_OVER_RIGHT_LEFT.
const StackTransitionType_over_right_left = StackTransitionType(19)

// StateType is a representation of the C enumeration GtkStateType.
type StateType int

// StateType_normal is a representation of the C enumeration member GTK_STATE_NORMAL.
const StateType_normal = StateType(0)

// StateType_active is a representation of the C enumeration member GTK_STATE_ACTIVE.
const StateType_active = StateType(1)

// StateType_prelight is a representation of the C enumeration member GTK_STATE_PRELIGHT.
const StateType_prelight = StateType(2)

// StateType_selected is a representation of the C enumeration member GTK_STATE_SELECTED.
const StateType_selected = StateType(3)

// StateType_insensitive is a representation of the C enumeration member GTK_STATE_INSENSITIVE.
const StateType_insensitive = StateType(4)

// StateType_inconsistent is a representation of the C enumeration member GTK_STATE_INCONSISTENT.
const StateType_inconsistent = StateType(5)

// StateType_focused is a representation of the C enumeration member GTK_STATE_FOCUSED.
const StateType_focused = StateType(6)

// TextBufferTargetInfo is a representation of the C enumeration GtkTextBufferTargetInfo.
type TextBufferTargetInfo int

// TextBufferTargetInfo_buffer_contents is a representation of the C enumeration member GTK_TEXT_BUFFER_TARGET_INFO_BUFFER_CONTENTS.
const TextBufferTargetInfo_buffer_contents = TextBufferTargetInfo(-1)

// TextBufferTargetInfo_rich_text is a representation of the C enumeration member GTK_TEXT_BUFFER_TARGET_INFO_RICH_TEXT.
const TextBufferTargetInfo_rich_text = TextBufferTargetInfo(-2)

// TextBufferTargetInfo_text is a representation of the C enumeration member GTK_TEXT_BUFFER_TARGET_INFO_TEXT.
const TextBufferTargetInfo_text = TextBufferTargetInfo(-3)

// TextDirection is a representation of the C enumeration GtkTextDirection.
type TextDirection int

// TextDirection_none is a representation of the C enumeration member GTK_TEXT_DIR_NONE.
const TextDirection_none = TextDirection(0)

// TextDirection_ltr is a representation of the C enumeration member GTK_TEXT_DIR_LTR.
const TextDirection_ltr = TextDirection(1)

// TextDirection_rtl is a representation of the C enumeration member GTK_TEXT_DIR_RTL.
const TextDirection_rtl = TextDirection(2)

// TextViewLayer is a representation of the C enumeration GtkTextViewLayer.
type TextViewLayer int

// TextViewLayer_below is a representation of the C enumeration member GTK_TEXT_VIEW_LAYER_BELOW.
const TextViewLayer_below = TextViewLayer(0)

// TextViewLayer_above is a representation of the C enumeration member GTK_TEXT_VIEW_LAYER_ABOVE.
const TextViewLayer_above = TextViewLayer(1)

// TextViewLayer_below_text is a representation of the C enumeration member GTK_TEXT_VIEW_LAYER_BELOW_TEXT.
const TextViewLayer_below_text = TextViewLayer(2)

// TextViewLayer_above_text is a representation of the C enumeration member GTK_TEXT_VIEW_LAYER_ABOVE_TEXT.
const TextViewLayer_above_text = TextViewLayer(3)

// TextWindowType is a representation of the C enumeration GtkTextWindowType.
type TextWindowType int

// TextWindowType_private is a representation of the C enumeration member GTK_TEXT_WINDOW_PRIVATE.
const TextWindowType_private = TextWindowType(0)

// TextWindowType_widget is a representation of the C enumeration member GTK_TEXT_WINDOW_WIDGET.
const TextWindowType_widget = TextWindowType(1)

// TextWindowType_text is a representation of the C enumeration member GTK_TEXT_WINDOW_TEXT.
const TextWindowType_text = TextWindowType(2)

// TextWindowType_left is a representation of the C enumeration member GTK_TEXT_WINDOW_LEFT.
const TextWindowType_left = TextWindowType(3)

// TextWindowType_right is a representation of the C enumeration member GTK_TEXT_WINDOW_RIGHT.
const TextWindowType_right = TextWindowType(4)

// TextWindowType_top is a representation of the C enumeration member GTK_TEXT_WINDOW_TOP.
const TextWindowType_top = TextWindowType(5)

// TextWindowType_bottom is a representation of the C enumeration member GTK_TEXT_WINDOW_BOTTOM.
const TextWindowType_bottom = TextWindowType(6)

// ToolbarSpaceStyle is a representation of the C enumeration GtkToolbarSpaceStyle.
type ToolbarSpaceStyle int

// ToolbarSpaceStyle_empty is a representation of the C enumeration member GTK_TOOLBAR_SPACE_EMPTY.
const ToolbarSpaceStyle_empty = ToolbarSpaceStyle(0)

// ToolbarSpaceStyle_line is a representation of the C enumeration member GTK_TOOLBAR_SPACE_LINE.
const ToolbarSpaceStyle_line = ToolbarSpaceStyle(1)

// ToolbarStyle is a representation of the C enumeration GtkToolbarStyle.
type ToolbarStyle int

// ToolbarStyle_icons is a representation of the C enumeration member GTK_TOOLBAR_ICONS.
const ToolbarStyle_icons = ToolbarStyle(0)

// ToolbarStyle_text is a representation of the C enumeration member GTK_TOOLBAR_TEXT.
const ToolbarStyle_text = ToolbarStyle(1)

// ToolbarStyle_both is a representation of the C enumeration member GTK_TOOLBAR_BOTH.
const ToolbarStyle_both = ToolbarStyle(2)

// ToolbarStyle_both_horiz is a representation of the C enumeration member GTK_TOOLBAR_BOTH_HORIZ.
const ToolbarStyle_both_horiz = ToolbarStyle(3)

// TreeViewColumnSizing is a representation of the C enumeration GtkTreeViewColumnSizing.
type TreeViewColumnSizing int

// TreeViewColumnSizing_grow_only is a representation of the C enumeration member GTK_TREE_VIEW_COLUMN_GROW_ONLY.
const TreeViewColumnSizing_grow_only = TreeViewColumnSizing(0)

// TreeViewColumnSizing_autosize is a representation of the C enumeration member GTK_TREE_VIEW_COLUMN_AUTOSIZE.
const TreeViewColumnSizing_autosize = TreeViewColumnSizing(1)

// TreeViewColumnSizing_fixed is a representation of the C enumeration member GTK_TREE_VIEW_COLUMN_FIXED.
const TreeViewColumnSizing_fixed = TreeViewColumnSizing(2)

// TreeViewDropPosition is a representation of the C enumeration GtkTreeViewDropPosition.
type TreeViewDropPosition int

// TreeViewDropPosition_before is a representation of the C enumeration member GTK_TREE_VIEW_DROP_BEFORE.
const TreeViewDropPosition_before = TreeViewDropPosition(0)

// TreeViewDropPosition_after is a representation of the C enumeration member GTK_TREE_VIEW_DROP_AFTER.
const TreeViewDropPosition_after = TreeViewDropPosition(1)

// TreeViewDropPosition_into_or_before is a representation of the C enumeration member GTK_TREE_VIEW_DROP_INTO_OR_BEFORE.
const TreeViewDropPosition_into_or_before = TreeViewDropPosition(2)

// TreeViewDropPosition_into_or_after is a representation of the C enumeration member GTK_TREE_VIEW_DROP_INTO_OR_AFTER.
const TreeViewDropPosition_into_or_after = TreeViewDropPosition(3)

// TreeViewGridLines is a representation of the C enumeration GtkTreeViewGridLines.
type TreeViewGridLines int

// TreeViewGridLines_none is a representation of the C enumeration member GTK_TREE_VIEW_GRID_LINES_NONE.
const TreeViewGridLines_none = TreeViewGridLines(0)

// TreeViewGridLines_horizontal is a representation of the C enumeration member GTK_TREE_VIEW_GRID_LINES_HORIZONTAL.
const TreeViewGridLines_horizontal = TreeViewGridLines(1)

// TreeViewGridLines_vertical is a representation of the C enumeration member GTK_TREE_VIEW_GRID_LINES_VERTICAL.
const TreeViewGridLines_vertical = TreeViewGridLines(2)

// TreeViewGridLines_both is a representation of the C enumeration member GTK_TREE_VIEW_GRID_LINES_BOTH.
const TreeViewGridLines_both = TreeViewGridLines(3)

// Unit is a representation of the C enumeration GtkUnit.
type Unit int

// Unit_none is a representation of the C enumeration member GTK_UNIT_NONE.
const Unit_none = Unit(0)

// Unit_points is a representation of the C enumeration member GTK_UNIT_POINTS.
const Unit_points = Unit(1)

// Unit_inch is a representation of the C enumeration member GTK_UNIT_INCH.
const Unit_inch = Unit(2)

// Unit_mm is a representation of the C enumeration member GTK_UNIT_MM.
const Unit_mm = Unit(3)

// WidgetHelpType is a representation of the C enumeration GtkWidgetHelpType.
type WidgetHelpType int

// WidgetHelpType_tooltip is a representation of the C enumeration member GTK_WIDGET_HELP_TOOLTIP.
const WidgetHelpType_tooltip = WidgetHelpType(0)

// WidgetHelpType_whats_this is a representation of the C enumeration member GTK_WIDGET_HELP_WHATS_THIS.
const WidgetHelpType_whats_this = WidgetHelpType(1)

// WindowPosition is a representation of the C enumeration GtkWindowPosition.
type WindowPosition int

// WindowPosition_none is a representation of the C enumeration member GTK_WIN_POS_NONE.
const WindowPosition_none = WindowPosition(0)

// WindowPosition_center is a representation of the C enumeration member GTK_WIN_POS_CENTER.
const WindowPosition_center = WindowPosition(1)

// WindowPosition_mouse is a representation of the C enumeration member GTK_WIN_POS_MOUSE.
const WindowPosition_mouse = WindowPosition(2)

// WindowPosition_center_always is a representation of the C enumeration member GTK_WIN_POS_CENTER_ALWAYS.
const WindowPosition_center_always = WindowPosition(3)

// WindowPosition_center_on_parent is a representation of the C enumeration member GTK_WIN_POS_CENTER_ON_PARENT.
const WindowPosition_center_on_parent = WindowPosition(4)

// WindowType is a representation of the C enumeration GtkWindowType.
type WindowType int

// WindowType_toplevel is a representation of the C enumeration member GTK_WINDOW_TOPLEVEL.
const WindowType_toplevel = WindowType(0)

// WindowType_popup is a representation of the C enumeration member GTK_WINDOW_POPUP.
const WindowType_popup = WindowType(1)

// WrapMode is a representation of the C enumeration GtkWrapMode.
type WrapMode int

// WrapMode_none is a representation of the C enumeration member GTK_WRAP_NONE.
const WrapMode_none = WrapMode(0)

// WrapMode_char is a representation of the C enumeration member GTK_WRAP_CHAR.
const WrapMode_char = WrapMode(1)

// WrapMode_word is a representation of the C enumeration member GTK_WRAP_WORD.
const WrapMode_word = WrapMode(2)

// WrapMode_word_char is a representation of the C enumeration member GTK_WRAP_WORD_CHAR.
const WrapMode_word_char = WrapMode(3)

// AccelGroupsActivate is analogous to the C function gtk_accel_groups_activate.
func AccelGroupsActivate(object *gobject.Object, accelKey uint, accelMods gdk.ModifierType) bool {
	sys_object := object.ToC()
	sys_accelKey := accelKey
	sys_accelMods := (int)(accelMods)
	retSys := gtk.Fn_gtk_accel_groups_activate(sys_object, sys_accelKey, sys_accelMods)
	ret := retSys

	return ret
}

// AccelGroupsFromObject is analogous to the C function gtk_accel_groups_from_object.
func AccelGroupsFromObject(object *gobject.Object) *glib.SList {
	sys_object := object.ToC()
	retSys := gtk.Fn_gtk_accel_groups_from_object(sys_object)
	ret := glib.SListNewFromC(retSys)

	return ret
}

// AcceleratorGetDefaultModMask is analogous to the C function gtk_accelerator_get_default_mod_mask.
func AcceleratorGetDefaultModMask() int {
	retSys := gtk.Fn_gtk_accelerator_get_default_mod_mask()
	ret := retSys

	return ret
}

// AcceleratorName is analogous to the C function gtk_accelerator_name.
func AcceleratorName(acceleratorKey uint, acceleratorMods gdk.ModifierType) string {
	sys_acceleratorKey := acceleratorKey
	sys_acceleratorMods := (int)(acceleratorMods)
	retSys := gtk.Fn_gtk_accelerator_name(sys_acceleratorKey, sys_acceleratorMods)
	ret := retSys

	return ret
}

// UNSUPPORTED : gtk_accelerator_parse : has [in]out param, accelerator_key

// UNSUPPORTED : gtk_accelerator_parse_with_keycode : parameter 'accelerator_codes' is array parameter without length parameter

// AcceleratorSetDefaultModMask is analogous to the C function gtk_accelerator_set_default_mod_mask.
func AcceleratorSetDefaultModMask(defaultModMask gdk.ModifierType) {
	sys_defaultModMask := (int)(defaultModMask)
	gtk.Fn_gtk_accelerator_set_default_mod_mask(sys_defaultModMask)
}

// AcceleratorValid is analogous to the C function gtk_accelerator_valid.
func AcceleratorValid(keyval uint, modifiers gdk.ModifierType) bool {
	sys_keyval := keyval
	sys_modifiers := (int)(modifiers)
	retSys := gtk.Fn_gtk_accelerator_valid(sys_keyval, sys_modifiers)
	ret := retSys

	return ret
}

// BindingsActivate is analogous to the C function gtk_bindings_activate.
func BindingsActivate(object *gobject.Object, keyval uint, modifiers gdk.ModifierType) bool {
	sys_object := object.ToC()
	sys_keyval := keyval
	sys_modifiers := (int)(modifiers)
	retSys := gtk.Fn_gtk_bindings_activate(sys_object, sys_keyval, sys_modifiers)
	ret := retSys

	return ret
}

// CheckVersion is analogous to the C function gtk_check_version.
func CheckVersion(requiredMajor uint, requiredMinor uint, requiredMicro uint) string {
	sys_requiredMajor := requiredMajor
	sys_requiredMinor := requiredMinor
	sys_requiredMicro := requiredMicro
	retSys := gtk.Fn_gtk_check_version(sys_requiredMajor, sys_requiredMinor, sys_requiredMicro)
	ret := retSys

	return ret
}

// DisableSetlocale is analogous to the C function gtk_disable_setlocale.
func DisableSetlocale() {
	gtk.Fn_gtk_disable_setlocale()
}

// DistributeNaturalAllocation is analogous to the C function gtk_distribute_natural_allocation.
func DistributeNaturalAllocation(extraSpace int, nRequestedSizes uint, sizes *RequestedSize) int {
	sys_extraSpace := extraSpace
	sys_nRequestedSizes := nRequestedSizes
	sys_sizes := sizes.ToC()
	retSys := gtk.Fn_gtk_distribute_natural_allocation(sys_extraSpace, sys_nRequestedSizes, sys_sizes)
	ret := retSys

	return ret
}

// DragFinish is analogous to the C function gtk_drag_finish.
func DragFinish(context *gdk.DragContext, success bool, del bool, time uint32) {
	sys_context := context.ToC()
	sys_success := success
	sys_del := del
	sys_time := time
	gtk.Fn_gtk_drag_finish(sys_context, sys_success, sys_del, sys_time)
}

// DragGetSourceWidget is analogous to the C function gtk_drag_get_source_widget.
func DragGetSourceWidget(context *gdk.DragContext) *Widget {
	sys_context := context.ToC()
	retSys := gtk.Fn_gtk_drag_get_source_widget(sys_context)
	ret := WidgetNewFromC(retSys)

	return ret
}

// DragSetIconDefault is analogous to the C function gtk_drag_set_icon_default.
func DragSetIconDefault(context *gdk.DragContext) {
	sys_context := context.ToC()
	gtk.Fn_gtk_drag_set_icon_default(sys_context)
}

// DragSetIconPixbuf is analogous to the C function gtk_drag_set_icon_pixbuf.
func DragSetIconPixbuf(context *gdk.DragContext, pixbuf *gdkpixbuf.Pixbuf, hotX int, hotY int) {
	sys_context := context.ToC()
	sys_pixbuf := pixbuf.ToC()
	sys_hotX := hotX
	sys_hotY := hotY
	gtk.Fn_gtk_drag_set_icon_pixbuf(sys_context, sys_pixbuf, sys_hotX, sys_hotY)
}

// DragSetIconStock is analogous to the C function gtk_drag_set_icon_stock.
func DragSetIconStock(context *gdk.DragContext, stockId string, hotX int, hotY int) {
	sys_context := context.ToC()
	sys_stockId := stockId
	sys_hotX := hotX
	sys_hotY := hotY
	gtk.Fn_gtk_drag_set_icon_stock(sys_context, sys_stockId, sys_hotX, sys_hotY)
}

// DragSetIconSurface is analogous to the C function gtk_drag_set_icon_surface.
func DragSetIconSurface(context *gdk.DragContext, surface *cairo.Surface) {
	sys_context := context.ToC()
	sys_surface := surface.ToC()
	gtk.Fn_gtk_drag_set_icon_surface(sys_context, sys_surface)
}

// DragSetIconWidget is analogous to the C function gtk_drag_set_icon_widget.
func DragSetIconWidget(context *gdk.DragContext, widget *Widget, hotX int, hotY int) {
	sys_context := context.ToC()
	sys_widget := widget.ToC()
	sys_hotX := hotX
	sys_hotY := hotY
	gtk.Fn_gtk_drag_set_icon_widget(sys_context, sys_widget, sys_hotX, sys_hotY)
}

// EventsPending is analogous to the C function gtk_events_pending.
func EventsPending() bool {
	retSys := gtk.Fn_gtk_events_pending()
	ret := retSys

	return ret
}

// False is analogous to the C function gtk_false.
func False() bool {
	retSys := gtk.Fn_gtk_false()
	ret := retSys

	return ret
}

// GetCurrentEvent is analogous to the C function gtk_get_current_event.
func GetCurrentEvent() *gdk.Event {
	retSys := gtk.Fn_gtk_get_current_event()
	ret := gdk.EventNewFromC(retSys)

	return ret
}

// GetCurrentEventDevice is analogous to the C function gtk_get_current_event_device.
func GetCurrentEventDevice() *gdk.Device {
	retSys := gtk.Fn_gtk_get_current_event_device()
	ret := gdk.DeviceNewFromC(retSys)

	return ret
}

// UNSUPPORTED : gtk_get_current_event_state : has [in]out param, state

// GetCurrentEventTime is analogous to the C function gtk_get_current_event_time.
func GetCurrentEventTime() uint32 {
	retSys := gtk.Fn_gtk_get_current_event_time()
	ret := retSys

	return ret
}

// GetDebugFlags is analogous to the C function gtk_get_debug_flags.
func GetDebugFlags() uint {
	retSys := gtk.Fn_gtk_get_debug_flags()
	ret := retSys

	return ret
}

// GetDefaultLanguage is analogous to the C function gtk_get_default_language.
func GetDefaultLanguage() *pango.Language {
	retSys := gtk.Fn_gtk_get_default_language()
	ret := pango.LanguageNewFromC(retSys)

	return ret
}

// GetEventWidget is analogous to the C function gtk_get_event_widget.
func GetEventWidget(event *gdk.Event) *Widget {
	sys_event := event.ToC()
	retSys := gtk.Fn_gtk_get_event_widget(sys_event)
	ret := WidgetNewFromC(retSys)

	return ret
}

// GrabGetCurrent is analogous to the C function gtk_grab_get_current.
func GrabGetCurrent() *Widget {
	retSys := gtk.Fn_gtk_grab_get_current()
	ret := WidgetNewFromC(retSys)

	return ret
}

// UNSUPPORTED : gtk_icon_size_lookup : has [in]out param, width

// UNSUPPORTED : gtk_icon_size_lookup_for_settings : has [in]out param, width

// UNSUPPORTED : gtk_init : has array param, argv

// UNSUPPORTED : gtk_init_check : has array param, argv

// UNSUPPORTED : gtk_init_with_args : parameter 'entries' is array parameter without length parameter

// UNSUPPORTED : gtk_key_snooper_install : parameter 'snooper' is callback

// KeySnooperRemove is analogous to the C function gtk_key_snooper_remove.
func KeySnooperRemove(snooperHandlerId uint) {
	sys_snooperHandlerId := snooperHandlerId
	gtk.Fn_gtk_key_snooper_remove(sys_snooperHandlerId)
}

// Main is analogous to the C function gtk_main.
func Main() {
	gtk.Fn_gtk_main()
}

// MainDoEvent is analogous to the C function gtk_main_do_event.
func MainDoEvent(event *gdk.Event) {
	sys_event := event.ToC()
	gtk.Fn_gtk_main_do_event(sys_event)
}

// MainIteration is analogous to the C function gtk_main_iteration.
func MainIteration() bool {
	retSys := gtk.Fn_gtk_main_iteration()
	ret := retSys

	return ret
}

// MainIterationDo is analogous to the C function gtk_main_iteration_do.
func MainIterationDo(blocking bool) bool {
	sys_blocking := blocking
	retSys := gtk.Fn_gtk_main_iteration_do(sys_blocking)
	ret := retSys

	return ret
}

// MainLevel is analogous to the C function gtk_main_level.
func MainLevel() uint {
	retSys := gtk.Fn_gtk_main_level()
	ret := retSys

	return ret
}

// MainQuit is analogous to the C function gtk_main_quit.
func MainQuit() {
	gtk.Fn_gtk_main_quit()
}

// PaintArrow is analogous to the C function gtk_paint_arrow.
func PaintArrow(style *Style, cr *cairo.Context, stateType StateType, shadowType ShadowType, widget *Widget, detail string, arrowType ArrowType, fill bool, x int, y int, width int, height int) {
	sys_style := style.ToC()
	sys_cr := cr.ToC()
	sys_stateType := (int)(stateType)
	sys_shadowType := (int)(shadowType)
	sys_widget := widget.ToC()
	sys_detail := detail
	sys_arrowType := (int)(arrowType)
	sys_fill := fill
	sys_x := x
	sys_y := y
	sys_width := width
	sys_height := height
	gtk.Fn_gtk_paint_arrow(sys_style, sys_cr, sys_stateType, sys_shadowType, sys_widget, sys_detail, sys_arrowType, sys_fill, sys_x, sys_y, sys_width, sys_height)
}

// PaintBox is analogous to the C function gtk_paint_box.
func PaintBox(style *Style, cr *cairo.Context, stateType StateType, shadowType ShadowType, widget *Widget, detail string, x int, y int, width int, height int) {
	sys_style := style.ToC()
	sys_cr := cr.ToC()
	sys_stateType := (int)(stateType)
	sys_shadowType := (int)(shadowType)
	sys_widget := widget.ToC()
	sys_detail := detail
	sys_x := x
	sys_y := y
	sys_width := width
	sys_height := height
	gtk.Fn_gtk_paint_box(sys_style, sys_cr, sys_stateType, sys_shadowType, sys_widget, sys_detail, sys_x, sys_y, sys_width, sys_height)
}

// PaintBoxGap is analogous to the C function gtk_paint_box_gap.
func PaintBoxGap(style *Style, cr *cairo.Context, stateType StateType, shadowType ShadowType, widget *Widget, detail string, x int, y int, width int, height int, gapSide PositionType, gapX int, gapWidth int) {
	sys_style := style.ToC()
	sys_cr := cr.ToC()
	sys_stateType := (int)(stateType)
	sys_shadowType := (int)(shadowType)
	sys_widget := widget.ToC()
	sys_detail := detail
	sys_x := x
	sys_y := y
	sys_width := width
	sys_height := height
	sys_gapSide := (int)(gapSide)
	sys_gapX := gapX
	sys_gapWidth := gapWidth
	gtk.Fn_gtk_paint_box_gap(sys_style, sys_cr, sys_stateType, sys_shadowType, sys_widget, sys_detail, sys_x, sys_y, sys_width, sys_height, sys_gapSide, sys_gapX, sys_gapWidth)
}

// PaintCheck is analogous to the C function gtk_paint_check.
func PaintCheck(style *Style, cr *cairo.Context, stateType StateType, shadowType ShadowType, widget *Widget, detail string, x int, y int, width int, height int) {
	sys_style := style.ToC()
	sys_cr := cr.ToC()
	sys_stateType := (int)(stateType)
	sys_shadowType := (int)(shadowType)
	sys_widget := widget.ToC()
	sys_detail := detail
	sys_x := x
	sys_y := y
	sys_width := width
	sys_height := height
	gtk.Fn_gtk_paint_check(sys_style, sys_cr, sys_stateType, sys_shadowType, sys_widget, sys_detail, sys_x, sys_y, sys_width, sys_height)
}

// PaintDiamond is analogous to the C function gtk_paint_diamond.
func PaintDiamond(style *Style, cr *cairo.Context, stateType StateType, shadowType ShadowType, widget *Widget, detail string, x int, y int, width int, height int) {
	sys_style := style.ToC()
	sys_cr := cr.ToC()
	sys_stateType := (int)(stateType)
	sys_shadowType := (int)(shadowType)
	sys_widget := widget.ToC()
	sys_detail := detail
	sys_x := x
	sys_y := y
	sys_width := width
	sys_height := height
	gtk.Fn_gtk_paint_diamond(sys_style, sys_cr, sys_stateType, sys_shadowType, sys_widget, sys_detail, sys_x, sys_y, sys_width, sys_height)
}

// PaintExpander is analogous to the C function gtk_paint_expander.
func PaintExpander(style *Style, cr *cairo.Context, stateType StateType, widget *Widget, detail string, x int, y int, expanderStyle ExpanderStyle) {
	sys_style := style.ToC()
	sys_cr := cr.ToC()
	sys_stateType := (int)(stateType)
	sys_widget := widget.ToC()
	sys_detail := detail
	sys_x := x
	sys_y := y
	sys_expanderStyle := (int)(expanderStyle)
	gtk.Fn_gtk_paint_expander(sys_style, sys_cr, sys_stateType, sys_widget, sys_detail, sys_x, sys_y, sys_expanderStyle)
}

// PaintExtension is analogous to the C function gtk_paint_extension.
func PaintExtension(style *Style, cr *cairo.Context, stateType StateType, shadowType ShadowType, widget *Widget, detail string, x int, y int, width int, height int, gapSide PositionType) {
	sys_style := style.ToC()
	sys_cr := cr.ToC()
	sys_stateType := (int)(stateType)
	sys_shadowType := (int)(shadowType)
	sys_widget := widget.ToC()
	sys_detail := detail
	sys_x := x
	sys_y := y
	sys_width := width
	sys_height := height
	sys_gapSide := (int)(gapSide)
	gtk.Fn_gtk_paint_extension(sys_style, sys_cr, sys_stateType, sys_shadowType, sys_widget, sys_detail, sys_x, sys_y, sys_width, sys_height, sys_gapSide)
}

// PaintFlatBox is analogous to the C function gtk_paint_flat_box.
func PaintFlatBox(style *Style, cr *cairo.Context, stateType StateType, shadowType ShadowType, widget *Widget, detail string, x int, y int, width int, height int) {
	sys_style := style.ToC()
	sys_cr := cr.ToC()
	sys_stateType := (int)(stateType)
	sys_shadowType := (int)(shadowType)
	sys_widget := widget.ToC()
	sys_detail := detail
	sys_x := x
	sys_y := y
	sys_width := width
	sys_height := height
	gtk.Fn_gtk_paint_flat_box(sys_style, sys_cr, sys_stateType, sys_shadowType, sys_widget, sys_detail, sys_x, sys_y, sys_width, sys_height)
}

// PaintFocus is analogous to the C function gtk_paint_focus.
func PaintFocus(style *Style, cr *cairo.Context, stateType StateType, widget *Widget, detail string, x int, y int, width int, height int) {
	sys_style := style.ToC()
	sys_cr := cr.ToC()
	sys_stateType := (int)(stateType)
	sys_widget := widget.ToC()
	sys_detail := detail
	sys_x := x
	sys_y := y
	sys_width := width
	sys_height := height
	gtk.Fn_gtk_paint_focus(sys_style, sys_cr, sys_stateType, sys_widget, sys_detail, sys_x, sys_y, sys_width, sys_height)
}

// PaintHandle is analogous to the C function gtk_paint_handle.
func PaintHandle(style *Style, cr *cairo.Context, stateType StateType, shadowType ShadowType, widget *Widget, detail string, x int, y int, width int, height int, orientation Orientation) {
	sys_style := style.ToC()
	sys_cr := cr.ToC()
	sys_stateType := (int)(stateType)
	sys_shadowType := (int)(shadowType)
	sys_widget := widget.ToC()
	sys_detail := detail
	sys_x := x
	sys_y := y
	sys_width := width
	sys_height := height
	sys_orientation := (int)(orientation)
	gtk.Fn_gtk_paint_handle(sys_style, sys_cr, sys_stateType, sys_shadowType, sys_widget, sys_detail, sys_x, sys_y, sys_width, sys_height, sys_orientation)
}

// PaintHline is analogous to the C function gtk_paint_hline.
func PaintHline(style *Style, cr *cairo.Context, stateType StateType, widget *Widget, detail string, x1 int, x2 int, y int) {
	sys_style := style.ToC()
	sys_cr := cr.ToC()
	sys_stateType := (int)(stateType)
	sys_widget := widget.ToC()
	sys_detail := detail
	sys_x1 := x1
	sys_x2 := x2
	sys_y := y
	gtk.Fn_gtk_paint_hline(sys_style, sys_cr, sys_stateType, sys_widget, sys_detail, sys_x1, sys_x2, sys_y)
}

// PaintLayout is analogous to the C function gtk_paint_layout.
func PaintLayout(style *Style, cr *cairo.Context, stateType StateType, useText bool, widget *Widget, detail string, x int, y int, layout *pango.Layout) {
	sys_style := style.ToC()
	sys_cr := cr.ToC()
	sys_stateType := (int)(stateType)
	sys_useText := useText
	sys_widget := widget.ToC()
	sys_detail := detail
	sys_x := x
	sys_y := y
	sys_layout := layout.ToC()
	gtk.Fn_gtk_paint_layout(sys_style, sys_cr, sys_stateType, sys_useText, sys_widget, sys_detail, sys_x, sys_y, sys_layout)
}

// PaintOption is analogous to the C function gtk_paint_option.
func PaintOption(style *Style, cr *cairo.Context, stateType StateType, shadowType ShadowType, widget *Widget, detail string, x int, y int, width int, height int) {
	sys_style := style.ToC()
	sys_cr := cr.ToC()
	sys_stateType := (int)(stateType)
	sys_shadowType := (int)(shadowType)
	sys_widget := widget.ToC()
	sys_detail := detail
	sys_x := x
	sys_y := y
	sys_width := width
	sys_height := height
	gtk.Fn_gtk_paint_option(sys_style, sys_cr, sys_stateType, sys_shadowType, sys_widget, sys_detail, sys_x, sys_y, sys_width, sys_height)
}

// PaintResizeGrip is analogous to the C function gtk_paint_resize_grip.
func PaintResizeGrip(style *Style, cr *cairo.Context, stateType StateType, widget *Widget, detail string, edge gdk.WindowEdge, x int, y int, width int, height int) {
	sys_style := style.ToC()
	sys_cr := cr.ToC()
	sys_stateType := (int)(stateType)
	sys_widget := widget.ToC()
	sys_detail := detail
	sys_edge := (int)(edge)
	sys_x := x
	sys_y := y
	sys_width := width
	sys_height := height
	gtk.Fn_gtk_paint_resize_grip(sys_style, sys_cr, sys_stateType, sys_widget, sys_detail, sys_edge, sys_x, sys_y, sys_width, sys_height)
}

// PaintShadow is analogous to the C function gtk_paint_shadow.
func PaintShadow(style *Style, cr *cairo.Context, stateType StateType, shadowType ShadowType, widget *Widget, detail string, x int, y int, width int, height int) {
	sys_style := style.ToC()
	sys_cr := cr.ToC()
	sys_stateType := (int)(stateType)
	sys_shadowType := (int)(shadowType)
	sys_widget := widget.ToC()
	sys_detail := detail
	sys_x := x
	sys_y := y
	sys_width := width
	sys_height := height
	gtk.Fn_gtk_paint_shadow(sys_style, sys_cr, sys_stateType, sys_shadowType, sys_widget, sys_detail, sys_x, sys_y, sys_width, sys_height)
}

// PaintShadowGap is analogous to the C function gtk_paint_shadow_gap.
func PaintShadowGap(style *Style, cr *cairo.Context, stateType StateType, shadowType ShadowType, widget *Widget, detail string, x int, y int, width int, height int, gapSide PositionType, gapX int, gapWidth int) {
	sys_style := style.ToC()
	sys_cr := cr.ToC()
	sys_stateType := (int)(stateType)
	sys_shadowType := (int)(shadowType)
	sys_widget := widget.ToC()
	sys_detail := detail
	sys_x := x
	sys_y := y
	sys_width := width
	sys_height := height
	sys_gapSide := (int)(gapSide)
	sys_gapX := gapX
	sys_gapWidth := gapWidth
	gtk.Fn_gtk_paint_shadow_gap(sys_style, sys_cr, sys_stateType, sys_shadowType, sys_widget, sys_detail, sys_x, sys_y, sys_width, sys_height, sys_gapSide, sys_gapX, sys_gapWidth)
}

// PaintSlider is analogous to the C function gtk_paint_slider.
func PaintSlider(style *Style, cr *cairo.Context, stateType StateType, shadowType ShadowType, widget *Widget, detail string, x int, y int, width int, height int, orientation Orientation) {
	sys_style := style.ToC()
	sys_cr := cr.ToC()
	sys_stateType := (int)(stateType)
	sys_shadowType := (int)(shadowType)
	sys_widget := widget.ToC()
	sys_detail := detail
	sys_x := x
	sys_y := y
	sys_width := width
	sys_height := height
	sys_orientation := (int)(orientation)
	gtk.Fn_gtk_paint_slider(sys_style, sys_cr, sys_stateType, sys_shadowType, sys_widget, sys_detail, sys_x, sys_y, sys_width, sys_height, sys_orientation)
}

// PaintSpinner is analogous to the C function gtk_paint_spinner.
func PaintSpinner(style *Style, cr *cairo.Context, stateType StateType, widget *Widget, detail string, step uint, x int, y int, width int, height int) {
	sys_style := style.ToC()
	sys_cr := cr.ToC()
	sys_stateType := (int)(stateType)
	sys_widget := widget.ToC()
	sys_detail := detail
	sys_step := step
	sys_x := x
	sys_y := y
	sys_width := width
	sys_height := height
	gtk.Fn_gtk_paint_spinner(sys_style, sys_cr, sys_stateType, sys_widget, sys_detail, sys_step, sys_x, sys_y, sys_width, sys_height)
}

// PaintTab is analogous to the C function gtk_paint_tab.
func PaintTab(style *Style, cr *cairo.Context, stateType StateType, shadowType ShadowType, widget *Widget, detail string, x int, y int, width int, height int) {
	sys_style := style.ToC()
	sys_cr := cr.ToC()
	sys_stateType := (int)(stateType)
	sys_shadowType := (int)(shadowType)
	sys_widget := widget.ToC()
	sys_detail := detail
	sys_x := x
	sys_y := y
	sys_width := width
	sys_height := height
	gtk.Fn_gtk_paint_tab(sys_style, sys_cr, sys_stateType, sys_shadowType, sys_widget, sys_detail, sys_x, sys_y, sys_width, sys_height)
}

// PaintVline is analogous to the C function gtk_paint_vline.
func PaintVline(style *Style, cr *cairo.Context, stateType StateType, widget *Widget, detail string, y1 int, y2 int, x int) {
	sys_style := style.ToC()
	sys_cr := cr.ToC()
	sys_stateType := (int)(stateType)
	sys_widget := widget.ToC()
	sys_detail := detail
	sys_y1 := y1
	sys_y2 := y2
	sys_x := x
	gtk.Fn_gtk_paint_vline(sys_style, sys_cr, sys_stateType, sys_widget, sys_detail, sys_y1, sys_y2, sys_x)
}

// UNSUPPORTED : gtk_parse_args : has array param, argv

// UNSUPPORTED : gtk_print_run_page_setup_dialog_async : parameter 'done_cb' is callback

// PropagateEvent is analogous to the C function gtk_propagate_event.
func PropagateEvent(widget *Widget, event *gdk.Event) {
	sys_widget := widget.ToC()
	sys_event := event.ToC()
	gtk.Fn_gtk_propagate_event(sys_widget, sys_event)
}

// RcAddDefaultFile is analogous to the C function gtk_rc_add_default_file.
func RcAddDefaultFile(filename string) {
	sys_filename := filename
	gtk.Fn_gtk_rc_add_default_file(sys_filename)
}

// RcFindModuleInPath is analogous to the C function gtk_rc_find_module_in_path.
func RcFindModuleInPath(moduleFile string) string {
	sys_moduleFile := moduleFile
	retSys := gtk.Fn_gtk_rc_find_module_in_path(sys_moduleFile)
	ret := retSys

	return ret
}

// RcFindPixmapInPath is analogous to the C function gtk_rc_find_pixmap_in_path.
func RcFindPixmapInPath(settings *Settings, scanner *glib.Scanner, pixmapFile string) string {
	sys_settings := settings.ToC()
	sys_scanner := scanner.ToC()
	sys_pixmapFile := pixmapFile
	retSys := gtk.Fn_gtk_rc_find_pixmap_in_path(sys_settings, sys_scanner, sys_pixmapFile)
	ret := retSys

	return ret
}

// UNSUPPORTED : gtk_rc_get_default_files : no array length

// RcGetImModuleFile is analogous to the C function gtk_rc_get_im_module_file.
func RcGetImModuleFile() string {
	retSys := gtk.Fn_gtk_rc_get_im_module_file()
	ret := retSys

	return ret
}

// RcGetImModulePath is analogous to the C function gtk_rc_get_im_module_path.
func RcGetImModulePath() string {
	retSys := gtk.Fn_gtk_rc_get_im_module_path()
	ret := retSys

	return ret
}

// RcGetModuleDir is analogous to the C function gtk_rc_get_module_dir.
func RcGetModuleDir() string {
	retSys := gtk.Fn_gtk_rc_get_module_dir()
	ret := retSys

	return ret
}

// RcGetStyle is analogous to the C function gtk_rc_get_style.
func RcGetStyle(widget *Widget) *Style {
	sys_widget := widget.ToC()
	retSys := gtk.Fn_gtk_rc_get_style(sys_widget)
	ret := StyleNewFromC(retSys)

	return ret
}

// RcGetStyleByPaths is analogous to the C function gtk_rc_get_style_by_paths.
func RcGetStyleByPaths(settings *Settings, widgetPath string, classPath string, type_ uint64) *Style {
	sys_settings := settings.ToC()
	sys_widgetPath := widgetPath
	sys_classPath := classPath
	sys_type_ := type_
	retSys := gtk.Fn_gtk_rc_get_style_by_paths(sys_settings, sys_widgetPath, sys_classPath, sys_type_)
	ret := StyleNewFromC(retSys)

	return ret
}

// RcGetThemeDir is analogous to the C function gtk_rc_get_theme_dir.
func RcGetThemeDir() string {
	retSys := gtk.Fn_gtk_rc_get_theme_dir()
	ret := retSys

	return ret
}

// RcParse is analogous to the C function gtk_rc_parse.
func RcParse(filename string) {
	sys_filename := filename
	gtk.Fn_gtk_rc_parse(sys_filename)
}

// UNSUPPORTED : gtk_rc_parse_color : has [in]out param, color

// UNSUPPORTED : gtk_rc_parse_color_full : has [in]out param, color

// RcParsePriority is analogous to the C function gtk_rc_parse_priority.
func RcParsePriority(scanner *glib.Scanner, priority *PathPriorityType) uint {
	sys_scanner := scanner.ToC()
	sys_priority := (*int)(priority)
	retSys := gtk.Fn_gtk_rc_parse_priority(sys_scanner, sys_priority)
	ret := retSys

	return ret
}

// UNSUPPORTED : gtk_rc_parse_state : has [in]out param, state

// RcParseString is analogous to the C function gtk_rc_parse_string.
func RcParseString(rcString string) {
	sys_rcString := rcString
	gtk.Fn_gtk_rc_parse_string(sys_rcString)
}

// RcReparseAll is analogous to the C function gtk_rc_reparse_all.
func RcReparseAll() bool {
	retSys := gtk.Fn_gtk_rc_reparse_all()
	ret := retSys

	return ret
}

// RcReparseAllForSettings is analogous to the C function gtk_rc_reparse_all_for_settings.
func RcReparseAllForSettings(settings *Settings, forceLoad bool) bool {
	sys_settings := settings.ToC()
	sys_forceLoad := forceLoad
	retSys := gtk.Fn_gtk_rc_reparse_all_for_settings(sys_settings, sys_forceLoad)
	ret := retSys

	return ret
}

// RcScannerNew is analogous to the C function gtk_rc_scanner_new.
func RcScannerNew() *glib.Scanner {
	retSys := gtk.Fn_gtk_rc_scanner_new()
	ret := glib.ScannerNewFromC(retSys)

	return ret
}

// UNSUPPORTED : gtk_rc_set_default_files : parameter 'filenames' is array parameter without length parameter

// UNSUPPORTED : gtk_render_background_get_clip : has [in]out param, out_clip

// UNSUPPORTED : gtk_rgb_to_hsv : has [in]out param, h

// SelectionAddTarget is analogous to the C function gtk_selection_add_target.
func SelectionAddTarget(widget *Widget, selection gdk.Atom, target gdk.Atom, info uint) {
	sys_widget := widget.ToC()
	sys_selection := selection.ToC()
	sys_target := target.ToC()
	sys_info := info
	gtk.Fn_gtk_selection_add_target(sys_widget, sys_selection, sys_target, sys_info)
}

// UNSUPPORTED : gtk_selection_add_targets : has array param, targets

// SelectionClearTargets is analogous to the C function gtk_selection_clear_targets.
func SelectionClearTargets(widget *Widget, selection gdk.Atom) {
	sys_widget := widget.ToC()
	sys_selection := selection.ToC()
	gtk.Fn_gtk_selection_clear_targets(sys_widget, sys_selection)
}

// SelectionConvert is analogous to the C function gtk_selection_convert.
func SelectionConvert(widget *Widget, selection gdk.Atom, target gdk.Atom, time uint32) bool {
	sys_widget := widget.ToC()
	sys_selection := selection.ToC()
	sys_target := target.ToC()
	sys_time := time
	retSys := gtk.Fn_gtk_selection_convert(sys_widget, sys_selection, sys_target, sys_time)
	ret := retSys

	return ret
}

// SelectionOwnerSet is analogous to the C function gtk_selection_owner_set.
func SelectionOwnerSet(widget *Widget, selection gdk.Atom, time uint32) bool {
	sys_widget := widget.ToC()
	sys_selection := selection.ToC()
	sys_time := time
	retSys := gtk.Fn_gtk_selection_owner_set(sys_widget, sys_selection, sys_time)
	ret := retSys

	return ret
}

// SelectionOwnerSetForDisplay is analogous to the C function gtk_selection_owner_set_for_display.
//
// since 2.2
func SelectionOwnerSetForDisplay(display *gdk.Display, widget *Widget, selection gdk.Atom, time uint32) bool {
	sys_display := display.ToC()
	sys_widget := widget.ToC()
	sys_selection := selection.ToC()
	sys_time := time
	retSys := gtk.Fn_gtk_selection_owner_set_for_display(sys_display, sys_widget, sys_selection, sys_time)
	ret := retSys

	return ret
}

// SelectionRemoveAll is analogous to the C function gtk_selection_remove_all.
func SelectionRemoveAll(widget *Widget) {
	sys_widget := widget.ToC()
	gtk.Fn_gtk_selection_remove_all(sys_widget)
}

// SetDebugFlags is analogous to the C function gtk_set_debug_flags.
func SetDebugFlags(flags uint) {
	sys_flags := flags
	gtk.Fn_gtk_set_debug_flags(sys_flags)
}

// UNSUPPORTED : gtk_show_uri : throws

// UNSUPPORTED : gtk_show_uri_on_window : throws

// UNSUPPORTED : gtk_stock_add : has array param, items

// UNSUPPORTED : gtk_stock_add_static : has array param, items

// StockListIds is analogous to the C function gtk_stock_list_ids.
func StockListIds() *glib.SList {
	retSys := gtk.Fn_gtk_stock_list_ids()
	ret := glib.SListNewFromC(retSys)

	return ret
}

// UNSUPPORTED : gtk_stock_lookup : has [in]out param, item

// UNSUPPORTED : gtk_stock_set_translate_func : parameter 'func' is callback

// UNSUPPORTED : gtk_target_table_free : has array param, targets

// UNSUPPORTED : gtk_target_table_new_from_list : has [in]out param, n_targets

// UNSUPPORTED : gtk_targets_include_image : has array param, targets

// UNSUPPORTED : gtk_targets_include_rich_text : has array param, targets

// UNSUPPORTED : gtk_targets_include_text : has array param, targets

// UNSUPPORTED : gtk_targets_include_uri : has array param, targets

// UNSUPPORTED : gtk_test_init : has array param, argvp

// UNSUPPORTED : gtk_test_list_all_types : has [in]out param, n_types

// UNSUPPORTED : gtk_tree_get_row_drag_data : has [in]out param, tree_model

// UNSUPPORTED : gtk_tree_row_reference_reordered : parameter 'new_order' is array parameter without length parameter

// TreeSetRowDragData is analogous to the C function gtk_tree_set_row_drag_data.
func TreeSetRowDragData(selectionData *SelectionData, treeModel *TreeModel, path *TreePath) bool {
	sys_selectionData := selectionData.ToC()
	sys_treeModel := treeModel.ToC()
	sys_path := path.ToC()
	retSys := gtk.Fn_gtk_tree_set_row_drag_data(sys_selectionData, sys_treeModel, sys_path)
	ret := retSys

	return ret
}

// True is analogous to the C function gtk_true.
func True() bool {
	retSys := gtk.Fn_gtk_true()
	ret := retSys

	return ret
}

// AboutDialogClass is a representation of the C record GtkAboutDialogClass.
type AboutDialogClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkAboutDialogClass that represents the AboutDialogClass.
func (recv *AboutDialogClass) ToC() unsafe.Pointer {
	return recv.native
}

// AboutDialogClassNewFromC creates a new AboutDialogClass from a pointer to the C GtkAboutDialogClass that represents the AboutDialogClass.
func AboutDialogClassNewFromC(native unsafe.Pointer) *AboutDialogClass {
	return &AboutDialogClass{native: native}
}

// AboutDialogPrivate is a representation of the C record GtkAboutDialogPrivate.
type AboutDialogPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkAboutDialogPrivate that represents the AboutDialogPrivate.
func (recv *AboutDialogPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// AboutDialogPrivateNewFromC creates a new AboutDialogPrivate from a pointer to the C GtkAboutDialogPrivate that represents the AboutDialogPrivate.
func AboutDialogPrivateNewFromC(native unsafe.Pointer) *AboutDialogPrivate {
	return &AboutDialogPrivate{native: native}
}

// AccelGroupClass is a representation of the C record GtkAccelGroupClass.
type AccelGroupClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkAccelGroupClass that represents the AccelGroupClass.
func (recv *AccelGroupClass) ToC() unsafe.Pointer {
	return recv.native
}

// AccelGroupClassNewFromC creates a new AccelGroupClass from a pointer to the C GtkAccelGroupClass that represents the AccelGroupClass.
func AccelGroupClassNewFromC(native unsafe.Pointer) *AccelGroupClass {
	return &AccelGroupClass{native: native}
}

// AccelGroupEntry is a representation of the C record GtkAccelGroupEntry.
type AccelGroupEntry struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkAccelGroupEntry that represents the AccelGroupEntry.
func (recv *AccelGroupEntry) ToC() unsafe.Pointer {
	return recv.native
}

// AccelGroupEntryNewFromC creates a new AccelGroupEntry from a pointer to the C GtkAccelGroupEntry that represents the AccelGroupEntry.
func AccelGroupEntryNewFromC(native unsafe.Pointer) *AccelGroupEntry {
	return &AccelGroupEntry{native: native}
}

// AccelGroupPrivate is a representation of the C record GtkAccelGroupPrivate.
type AccelGroupPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkAccelGroupPrivate that represents the AccelGroupPrivate.
func (recv *AccelGroupPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// AccelGroupPrivateNewFromC creates a new AccelGroupPrivate from a pointer to the C GtkAccelGroupPrivate that represents the AccelGroupPrivate.
func AccelGroupPrivateNewFromC(native unsafe.Pointer) *AccelGroupPrivate {
	return &AccelGroupPrivate{native: native}
}

// AccelKey is a representation of the C record GtkAccelKey.
type AccelKey struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkAccelKey that represents the AccelKey.
func (recv *AccelKey) ToC() unsafe.Pointer {
	return recv.native
}

// AccelKeyNewFromC creates a new AccelKey from a pointer to the C GtkAccelKey that represents the AccelKey.
func AccelKeyNewFromC(native unsafe.Pointer) *AccelKey {
	return &AccelKey{native: native}
}

// AccelLabelClass is a representation of the C record GtkAccelLabelClass.
type AccelLabelClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkAccelLabelClass that represents the AccelLabelClass.
func (recv *AccelLabelClass) ToC() unsafe.Pointer {
	return recv.native
}

// AccelLabelClassNewFromC creates a new AccelLabelClass from a pointer to the C GtkAccelLabelClass that represents the AccelLabelClass.
func AccelLabelClassNewFromC(native unsafe.Pointer) *AccelLabelClass {
	return &AccelLabelClass{native: native}
}

// AccelLabelPrivate is a representation of the C record GtkAccelLabelPrivate.
type AccelLabelPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkAccelLabelPrivate that represents the AccelLabelPrivate.
func (recv *AccelLabelPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// AccelLabelPrivateNewFromC creates a new AccelLabelPrivate from a pointer to the C GtkAccelLabelPrivate that represents the AccelLabelPrivate.
func AccelLabelPrivateNewFromC(native unsafe.Pointer) *AccelLabelPrivate {
	return &AccelLabelPrivate{native: native}
}

// AccelMapClass is a representation of the C record GtkAccelMapClass.
type AccelMapClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkAccelMapClass that represents the AccelMapClass.
func (recv *AccelMapClass) ToC() unsafe.Pointer {
	return recv.native
}

// AccelMapClassNewFromC creates a new AccelMapClass from a pointer to the C GtkAccelMapClass that represents the AccelMapClass.
func AccelMapClassNewFromC(native unsafe.Pointer) *AccelMapClass {
	return &AccelMapClass{native: native}
}

// AccessibleClass is a representation of the C record GtkAccessibleClass.
type AccessibleClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkAccessibleClass that represents the AccessibleClass.
func (recv *AccessibleClass) ToC() unsafe.Pointer {
	return recv.native
}

// AccessibleClassNewFromC creates a new AccessibleClass from a pointer to the C GtkAccessibleClass that represents the AccessibleClass.
func AccessibleClassNewFromC(native unsafe.Pointer) *AccessibleClass {
	return &AccessibleClass{native: native}
}

// AccessiblePrivate is a representation of the C record GtkAccessiblePrivate.
type AccessiblePrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkAccessiblePrivate that represents the AccessiblePrivate.
func (recv *AccessiblePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// AccessiblePrivateNewFromC creates a new AccessiblePrivate from a pointer to the C GtkAccessiblePrivate that represents the AccessiblePrivate.
func AccessiblePrivateNewFromC(native unsafe.Pointer) *AccessiblePrivate {
	return &AccessiblePrivate{native: native}
}

// ActionBarClass is a representation of the C record GtkActionBarClass.
type ActionBarClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkActionBarClass that represents the ActionBarClass.
func (recv *ActionBarClass) ToC() unsafe.Pointer {
	return recv.native
}

// ActionBarClassNewFromC creates a new ActionBarClass from a pointer to the C GtkActionBarClass that represents the ActionBarClass.
func ActionBarClassNewFromC(native unsafe.Pointer) *ActionBarClass {
	return &ActionBarClass{native: native}
}

// ActionBarPrivate is a representation of the C record GtkActionBarPrivate.
type ActionBarPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkActionBarPrivate that represents the ActionBarPrivate.
func (recv *ActionBarPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ActionBarPrivateNewFromC creates a new ActionBarPrivate from a pointer to the C GtkActionBarPrivate that represents the ActionBarPrivate.
func ActionBarPrivateNewFromC(native unsafe.Pointer) *ActionBarPrivate {
	return &ActionBarPrivate{native: native}
}

// ActionClass is a representation of the C record GtkActionClass.
type ActionClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkActionClass that represents the ActionClass.
func (recv *ActionClass) ToC() unsafe.Pointer {
	return recv.native
}

// ActionClassNewFromC creates a new ActionClass from a pointer to the C GtkActionClass that represents the ActionClass.
func ActionClassNewFromC(native unsafe.Pointer) *ActionClass {
	return &ActionClass{native: native}
}

// ActionEntry is a representation of the C record GtkActionEntry.
type ActionEntry struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkActionEntry that represents the ActionEntry.
func (recv *ActionEntry) ToC() unsafe.Pointer {
	return recv.native
}

// ActionEntryNewFromC creates a new ActionEntry from a pointer to the C GtkActionEntry that represents the ActionEntry.
func ActionEntryNewFromC(native unsafe.Pointer) *ActionEntry {
	return &ActionEntry{native: native}
}

// ActionGroupClass is a representation of the C record GtkActionGroupClass.
type ActionGroupClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkActionGroupClass that represents the ActionGroupClass.
func (recv *ActionGroupClass) ToC() unsafe.Pointer {
	return recv.native
}

// ActionGroupClassNewFromC creates a new ActionGroupClass from a pointer to the C GtkActionGroupClass that represents the ActionGroupClass.
func ActionGroupClassNewFromC(native unsafe.Pointer) *ActionGroupClass {
	return &ActionGroupClass{native: native}
}

// ActionGroupPrivate is a representation of the C record GtkActionGroupPrivate.
type ActionGroupPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkActionGroupPrivate that represents the ActionGroupPrivate.
func (recv *ActionGroupPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ActionGroupPrivateNewFromC creates a new ActionGroupPrivate from a pointer to the C GtkActionGroupPrivate that represents the ActionGroupPrivate.
func ActionGroupPrivateNewFromC(native unsafe.Pointer) *ActionGroupPrivate {
	return &ActionGroupPrivate{native: native}
}

// ActionPrivate is a representation of the C record GtkActionPrivate.
type ActionPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkActionPrivate that represents the ActionPrivate.
func (recv *ActionPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ActionPrivateNewFromC creates a new ActionPrivate from a pointer to the C GtkActionPrivate that represents the ActionPrivate.
func ActionPrivateNewFromC(native unsafe.Pointer) *ActionPrivate {
	return &ActionPrivate{native: native}
}

// ActionableInterface is a representation of the C record GtkActionableInterface.
type ActionableInterface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkActionableInterface that represents the ActionableInterface.
func (recv *ActionableInterface) ToC() unsafe.Pointer {
	return recv.native
}

// ActionableInterfaceNewFromC creates a new ActionableInterface from a pointer to the C GtkActionableInterface that represents the ActionableInterface.
func ActionableInterfaceNewFromC(native unsafe.Pointer) *ActionableInterface {
	return &ActionableInterface{native: native}
}

// AdjustmentClass is a representation of the C record GtkAdjustmentClass.
type AdjustmentClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkAdjustmentClass that represents the AdjustmentClass.
func (recv *AdjustmentClass) ToC() unsafe.Pointer {
	return recv.native
}

// AdjustmentClassNewFromC creates a new AdjustmentClass from a pointer to the C GtkAdjustmentClass that represents the AdjustmentClass.
func AdjustmentClassNewFromC(native unsafe.Pointer) *AdjustmentClass {
	return &AdjustmentClass{native: native}
}

// AdjustmentPrivate is a representation of the C record GtkAdjustmentPrivate.
type AdjustmentPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkAdjustmentPrivate that represents the AdjustmentPrivate.
func (recv *AdjustmentPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// AdjustmentPrivateNewFromC creates a new AdjustmentPrivate from a pointer to the C GtkAdjustmentPrivate that represents the AdjustmentPrivate.
func AdjustmentPrivateNewFromC(native unsafe.Pointer) *AdjustmentPrivate {
	return &AdjustmentPrivate{native: native}
}

// AlignmentClass is a representation of the C record GtkAlignmentClass.
type AlignmentClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkAlignmentClass that represents the AlignmentClass.
func (recv *AlignmentClass) ToC() unsafe.Pointer {
	return recv.native
}

// AlignmentClassNewFromC creates a new AlignmentClass from a pointer to the C GtkAlignmentClass that represents the AlignmentClass.
func AlignmentClassNewFromC(native unsafe.Pointer) *AlignmentClass {
	return &AlignmentClass{native: native}
}

// AlignmentPrivate is a representation of the C record GtkAlignmentPrivate.
type AlignmentPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkAlignmentPrivate that represents the AlignmentPrivate.
func (recv *AlignmentPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// AlignmentPrivateNewFromC creates a new AlignmentPrivate from a pointer to the C GtkAlignmentPrivate that represents the AlignmentPrivate.
func AlignmentPrivateNewFromC(native unsafe.Pointer) *AlignmentPrivate {
	return &AlignmentPrivate{native: native}
}

// AppChooserButtonClass is a representation of the C record GtkAppChooserButtonClass.
type AppChooserButtonClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkAppChooserButtonClass that represents the AppChooserButtonClass.
func (recv *AppChooserButtonClass) ToC() unsafe.Pointer {
	return recv.native
}

// AppChooserButtonClassNewFromC creates a new AppChooserButtonClass from a pointer to the C GtkAppChooserButtonClass that represents the AppChooserButtonClass.
func AppChooserButtonClassNewFromC(native unsafe.Pointer) *AppChooserButtonClass {
	return &AppChooserButtonClass{native: native}
}

// AppChooserButtonPrivate is a representation of the C record GtkAppChooserButtonPrivate.
type AppChooserButtonPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkAppChooserButtonPrivate that represents the AppChooserButtonPrivate.
func (recv *AppChooserButtonPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// AppChooserButtonPrivateNewFromC creates a new AppChooserButtonPrivate from a pointer to the C GtkAppChooserButtonPrivate that represents the AppChooserButtonPrivate.
func AppChooserButtonPrivateNewFromC(native unsafe.Pointer) *AppChooserButtonPrivate {
	return &AppChooserButtonPrivate{native: native}
}

// AppChooserDialogClass is a representation of the C record GtkAppChooserDialogClass.
type AppChooserDialogClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkAppChooserDialogClass that represents the AppChooserDialogClass.
func (recv *AppChooserDialogClass) ToC() unsafe.Pointer {
	return recv.native
}

// AppChooserDialogClassNewFromC creates a new AppChooserDialogClass from a pointer to the C GtkAppChooserDialogClass that represents the AppChooserDialogClass.
func AppChooserDialogClassNewFromC(native unsafe.Pointer) *AppChooserDialogClass {
	return &AppChooserDialogClass{native: native}
}

// AppChooserDialogPrivate is a representation of the C record GtkAppChooserDialogPrivate.
type AppChooserDialogPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkAppChooserDialogPrivate that represents the AppChooserDialogPrivate.
func (recv *AppChooserDialogPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// AppChooserDialogPrivateNewFromC creates a new AppChooserDialogPrivate from a pointer to the C GtkAppChooserDialogPrivate that represents the AppChooserDialogPrivate.
func AppChooserDialogPrivateNewFromC(native unsafe.Pointer) *AppChooserDialogPrivate {
	return &AppChooserDialogPrivate{native: native}
}

// AppChooserWidgetClass is a representation of the C record GtkAppChooserWidgetClass.
type AppChooserWidgetClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkAppChooserWidgetClass that represents the AppChooserWidgetClass.
func (recv *AppChooserWidgetClass) ToC() unsafe.Pointer {
	return recv.native
}

// AppChooserWidgetClassNewFromC creates a new AppChooserWidgetClass from a pointer to the C GtkAppChooserWidgetClass that represents the AppChooserWidgetClass.
func AppChooserWidgetClassNewFromC(native unsafe.Pointer) *AppChooserWidgetClass {
	return &AppChooserWidgetClass{native: native}
}

// AppChooserWidgetPrivate is a representation of the C record GtkAppChooserWidgetPrivate.
type AppChooserWidgetPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkAppChooserWidgetPrivate that represents the AppChooserWidgetPrivate.
func (recv *AppChooserWidgetPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// AppChooserWidgetPrivateNewFromC creates a new AppChooserWidgetPrivate from a pointer to the C GtkAppChooserWidgetPrivate that represents the AppChooserWidgetPrivate.
func AppChooserWidgetPrivateNewFromC(native unsafe.Pointer) *AppChooserWidgetPrivate {
	return &AppChooserWidgetPrivate{native: native}
}

// ApplicationClass is a representation of the C record GtkApplicationClass.
type ApplicationClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkApplicationClass that represents the ApplicationClass.
func (recv *ApplicationClass) ToC() unsafe.Pointer {
	return recv.native
}

// ApplicationClassNewFromC creates a new ApplicationClass from a pointer to the C GtkApplicationClass that represents the ApplicationClass.
func ApplicationClassNewFromC(native unsafe.Pointer) *ApplicationClass {
	return &ApplicationClass{native: native}
}

// ApplicationPrivate is a representation of the C record GtkApplicationPrivate.
type ApplicationPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkApplicationPrivate that represents the ApplicationPrivate.
func (recv *ApplicationPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ApplicationPrivateNewFromC creates a new ApplicationPrivate from a pointer to the C GtkApplicationPrivate that represents the ApplicationPrivate.
func ApplicationPrivateNewFromC(native unsafe.Pointer) *ApplicationPrivate {
	return &ApplicationPrivate{native: native}
}

// ApplicationWindowClass is a representation of the C record GtkApplicationWindowClass.
type ApplicationWindowClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkApplicationWindowClass that represents the ApplicationWindowClass.
func (recv *ApplicationWindowClass) ToC() unsafe.Pointer {
	return recv.native
}

// ApplicationWindowClassNewFromC creates a new ApplicationWindowClass from a pointer to the C GtkApplicationWindowClass that represents the ApplicationWindowClass.
func ApplicationWindowClassNewFromC(native unsafe.Pointer) *ApplicationWindowClass {
	return &ApplicationWindowClass{native: native}
}

// ApplicationWindowPrivate is a representation of the C record GtkApplicationWindowPrivate.
type ApplicationWindowPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkApplicationWindowPrivate that represents the ApplicationWindowPrivate.
func (recv *ApplicationWindowPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ApplicationWindowPrivateNewFromC creates a new ApplicationWindowPrivate from a pointer to the C GtkApplicationWindowPrivate that represents the ApplicationWindowPrivate.
func ApplicationWindowPrivateNewFromC(native unsafe.Pointer) *ApplicationWindowPrivate {
	return &ApplicationWindowPrivate{native: native}
}

// ArrowAccessibleClass is a representation of the C record GtkArrowAccessibleClass.
type ArrowAccessibleClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkArrowAccessibleClass that represents the ArrowAccessibleClass.
func (recv *ArrowAccessibleClass) ToC() unsafe.Pointer {
	return recv.native
}

// ArrowAccessibleClassNewFromC creates a new ArrowAccessibleClass from a pointer to the C GtkArrowAccessibleClass that represents the ArrowAccessibleClass.
func ArrowAccessibleClassNewFromC(native unsafe.Pointer) *ArrowAccessibleClass {
	return &ArrowAccessibleClass{native: native}
}

// ArrowAccessiblePrivate is a representation of the C record GtkArrowAccessiblePrivate.
type ArrowAccessiblePrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkArrowAccessiblePrivate that represents the ArrowAccessiblePrivate.
func (recv *ArrowAccessiblePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ArrowAccessiblePrivateNewFromC creates a new ArrowAccessiblePrivate from a pointer to the C GtkArrowAccessiblePrivate that represents the ArrowAccessiblePrivate.
func ArrowAccessiblePrivateNewFromC(native unsafe.Pointer) *ArrowAccessiblePrivate {
	return &ArrowAccessiblePrivate{native: native}
}

// ArrowClass is a representation of the C record GtkArrowClass.
type ArrowClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkArrowClass that represents the ArrowClass.
func (recv *ArrowClass) ToC() unsafe.Pointer {
	return recv.native
}

// ArrowClassNewFromC creates a new ArrowClass from a pointer to the C GtkArrowClass that represents the ArrowClass.
func ArrowClassNewFromC(native unsafe.Pointer) *ArrowClass {
	return &ArrowClass{native: native}
}

// ArrowPrivate is a representation of the C record GtkArrowPrivate.
type ArrowPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkArrowPrivate that represents the ArrowPrivate.
func (recv *ArrowPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ArrowPrivateNewFromC creates a new ArrowPrivate from a pointer to the C GtkArrowPrivate that represents the ArrowPrivate.
func ArrowPrivateNewFromC(native unsafe.Pointer) *ArrowPrivate {
	return &ArrowPrivate{native: native}
}

// AspectFrameClass is a representation of the C record GtkAspectFrameClass.
type AspectFrameClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkAspectFrameClass that represents the AspectFrameClass.
func (recv *AspectFrameClass) ToC() unsafe.Pointer {
	return recv.native
}

// AspectFrameClassNewFromC creates a new AspectFrameClass from a pointer to the C GtkAspectFrameClass that represents the AspectFrameClass.
func AspectFrameClassNewFromC(native unsafe.Pointer) *AspectFrameClass {
	return &AspectFrameClass{native: native}
}

// AspectFramePrivate is a representation of the C record GtkAspectFramePrivate.
type AspectFramePrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkAspectFramePrivate that represents the AspectFramePrivate.
func (recv *AspectFramePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// AspectFramePrivateNewFromC creates a new AspectFramePrivate from a pointer to the C GtkAspectFramePrivate that represents the AspectFramePrivate.
func AspectFramePrivateNewFromC(native unsafe.Pointer) *AspectFramePrivate {
	return &AspectFramePrivate{native: native}
}

// AssistantClass is a representation of the C record GtkAssistantClass.
type AssistantClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkAssistantClass that represents the AssistantClass.
func (recv *AssistantClass) ToC() unsafe.Pointer {
	return recv.native
}

// AssistantClassNewFromC creates a new AssistantClass from a pointer to the C GtkAssistantClass that represents the AssistantClass.
func AssistantClassNewFromC(native unsafe.Pointer) *AssistantClass {
	return &AssistantClass{native: native}
}

// AssistantPrivate is a representation of the C record GtkAssistantPrivate.
type AssistantPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkAssistantPrivate that represents the AssistantPrivate.
func (recv *AssistantPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// AssistantPrivateNewFromC creates a new AssistantPrivate from a pointer to the C GtkAssistantPrivate that represents the AssistantPrivate.
func AssistantPrivateNewFromC(native unsafe.Pointer) *AssistantPrivate {
	return &AssistantPrivate{native: native}
}

// BinClass is a representation of the C record GtkBinClass.
type BinClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkBinClass that represents the BinClass.
func (recv *BinClass) ToC() unsafe.Pointer {
	return recv.native
}

// BinClassNewFromC creates a new BinClass from a pointer to the C GtkBinClass that represents the BinClass.
func BinClassNewFromC(native unsafe.Pointer) *BinClass {
	return &BinClass{native: native}
}

// BinPrivate is a representation of the C record GtkBinPrivate.
type BinPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkBinPrivate that represents the BinPrivate.
func (recv *BinPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// BinPrivateNewFromC creates a new BinPrivate from a pointer to the C GtkBinPrivate that represents the BinPrivate.
func BinPrivateNewFromC(native unsafe.Pointer) *BinPrivate {
	return &BinPrivate{native: native}
}

// BindingArg is a representation of the C record GtkBindingArg.
type BindingArg struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkBindingArg that represents the BindingArg.
func (recv *BindingArg) ToC() unsafe.Pointer {
	return recv.native
}

// BindingArgNewFromC creates a new BindingArg from a pointer to the C GtkBindingArg that represents the BindingArg.
func BindingArgNewFromC(native unsafe.Pointer) *BindingArg {
	return &BindingArg{native: native}
}

// BindingEntry is a representation of the C record GtkBindingEntry.
type BindingEntry struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkBindingEntry that represents the BindingEntry.
func (recv *BindingEntry) ToC() unsafe.Pointer {
	return recv.native
}

// BindingEntryNewFromC creates a new BindingEntry from a pointer to the C GtkBindingEntry that represents the BindingEntry.
func BindingEntryNewFromC(native unsafe.Pointer) *BindingEntry {
	return &BindingEntry{native: native}
}

// BindingSet is a representation of the C record GtkBindingSet.
type BindingSet struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkBindingSet that represents the BindingSet.
func (recv *BindingSet) ToC() unsafe.Pointer {
	return recv.native
}

// BindingSetNewFromC creates a new BindingSet from a pointer to the C GtkBindingSet that represents the BindingSet.
func BindingSetNewFromC(native unsafe.Pointer) *BindingSet {
	return &BindingSet{native: native}
}

// BindingSignal is a representation of the C record GtkBindingSignal.
type BindingSignal struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkBindingSignal that represents the BindingSignal.
func (recv *BindingSignal) ToC() unsafe.Pointer {
	return recv.native
}

// BindingSignalNewFromC creates a new BindingSignal from a pointer to the C GtkBindingSignal that represents the BindingSignal.
func BindingSignalNewFromC(native unsafe.Pointer) *BindingSignal {
	return &BindingSignal{native: native}
}

// BooleanCellAccessibleClass is a representation of the C record GtkBooleanCellAccessibleClass.
type BooleanCellAccessibleClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkBooleanCellAccessibleClass that represents the BooleanCellAccessibleClass.
func (recv *BooleanCellAccessibleClass) ToC() unsafe.Pointer {
	return recv.native
}

// BooleanCellAccessibleClassNewFromC creates a new BooleanCellAccessibleClass from a pointer to the C GtkBooleanCellAccessibleClass that represents the BooleanCellAccessibleClass.
func BooleanCellAccessibleClassNewFromC(native unsafe.Pointer) *BooleanCellAccessibleClass {
	return &BooleanCellAccessibleClass{native: native}
}

// BooleanCellAccessiblePrivate is a representation of the C record GtkBooleanCellAccessiblePrivate.
type BooleanCellAccessiblePrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkBooleanCellAccessiblePrivate that represents the BooleanCellAccessiblePrivate.
func (recv *BooleanCellAccessiblePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// BooleanCellAccessiblePrivateNewFromC creates a new BooleanCellAccessiblePrivate from a pointer to the C GtkBooleanCellAccessiblePrivate that represents the BooleanCellAccessiblePrivate.
func BooleanCellAccessiblePrivateNewFromC(native unsafe.Pointer) *BooleanCellAccessiblePrivate {
	return &BooleanCellAccessiblePrivate{native: native}
}

// Border is a representation of the C record GtkBorder.
type Border struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkBorder that represents the Border.
func (recv *Border) ToC() unsafe.Pointer {
	return recv.native
}

// BorderNewFromC creates a new Border from a pointer to the C GtkBorder that represents the Border.
func BorderNewFromC(native unsafe.Pointer) *Border {
	return &Border{native: native}
}

// BoxClass is a representation of the C record GtkBoxClass.
type BoxClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkBoxClass that represents the BoxClass.
func (recv *BoxClass) ToC() unsafe.Pointer {
	return recv.native
}

// BoxClassNewFromC creates a new BoxClass from a pointer to the C GtkBoxClass that represents the BoxClass.
func BoxClassNewFromC(native unsafe.Pointer) *BoxClass {
	return &BoxClass{native: native}
}

// BoxPrivate is a representation of the C record GtkBoxPrivate.
type BoxPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkBoxPrivate that represents the BoxPrivate.
func (recv *BoxPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// BoxPrivateNewFromC creates a new BoxPrivate from a pointer to the C GtkBoxPrivate that represents the BoxPrivate.
func BoxPrivateNewFromC(native unsafe.Pointer) *BoxPrivate {
	return &BoxPrivate{native: native}
}

// BuildableIface is a representation of the C record GtkBuildableIface.
type BuildableIface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkBuildableIface that represents the BuildableIface.
func (recv *BuildableIface) ToC() unsafe.Pointer {
	return recv.native
}

// BuildableIfaceNewFromC creates a new BuildableIface from a pointer to the C GtkBuildableIface that represents the BuildableIface.
func BuildableIfaceNewFromC(native unsafe.Pointer) *BuildableIface {
	return &BuildableIface{native: native}
}

// BuilderClass is a representation of the C record GtkBuilderClass.
type BuilderClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkBuilderClass that represents the BuilderClass.
func (recv *BuilderClass) ToC() unsafe.Pointer {
	return recv.native
}

// BuilderClassNewFromC creates a new BuilderClass from a pointer to the C GtkBuilderClass that represents the BuilderClass.
func BuilderClassNewFromC(native unsafe.Pointer) *BuilderClass {
	return &BuilderClass{native: native}
}

// BuilderPrivate is a representation of the C record GtkBuilderPrivate.
type BuilderPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkBuilderPrivate that represents the BuilderPrivate.
func (recv *BuilderPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// BuilderPrivateNewFromC creates a new BuilderPrivate from a pointer to the C GtkBuilderPrivate that represents the BuilderPrivate.
func BuilderPrivateNewFromC(native unsafe.Pointer) *BuilderPrivate {
	return &BuilderPrivate{native: native}
}

// ButtonAccessibleClass is a representation of the C record GtkButtonAccessibleClass.
type ButtonAccessibleClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkButtonAccessibleClass that represents the ButtonAccessibleClass.
func (recv *ButtonAccessibleClass) ToC() unsafe.Pointer {
	return recv.native
}

// ButtonAccessibleClassNewFromC creates a new ButtonAccessibleClass from a pointer to the C GtkButtonAccessibleClass that represents the ButtonAccessibleClass.
func ButtonAccessibleClassNewFromC(native unsafe.Pointer) *ButtonAccessibleClass {
	return &ButtonAccessibleClass{native: native}
}

// ButtonAccessiblePrivate is a representation of the C record GtkButtonAccessiblePrivate.
type ButtonAccessiblePrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkButtonAccessiblePrivate that represents the ButtonAccessiblePrivate.
func (recv *ButtonAccessiblePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ButtonAccessiblePrivateNewFromC creates a new ButtonAccessiblePrivate from a pointer to the C GtkButtonAccessiblePrivate that represents the ButtonAccessiblePrivate.
func ButtonAccessiblePrivateNewFromC(native unsafe.Pointer) *ButtonAccessiblePrivate {
	return &ButtonAccessiblePrivate{native: native}
}

// ButtonBoxClass is a representation of the C record GtkButtonBoxClass.
type ButtonBoxClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkButtonBoxClass that represents the ButtonBoxClass.
func (recv *ButtonBoxClass) ToC() unsafe.Pointer {
	return recv.native
}

// ButtonBoxClassNewFromC creates a new ButtonBoxClass from a pointer to the C GtkButtonBoxClass that represents the ButtonBoxClass.
func ButtonBoxClassNewFromC(native unsafe.Pointer) *ButtonBoxClass {
	return &ButtonBoxClass{native: native}
}

// ButtonBoxPrivate is a representation of the C record GtkButtonBoxPrivate.
type ButtonBoxPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkButtonBoxPrivate that represents the ButtonBoxPrivate.
func (recv *ButtonBoxPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ButtonBoxPrivateNewFromC creates a new ButtonBoxPrivate from a pointer to the C GtkButtonBoxPrivate that represents the ButtonBoxPrivate.
func ButtonBoxPrivateNewFromC(native unsafe.Pointer) *ButtonBoxPrivate {
	return &ButtonBoxPrivate{native: native}
}

// ButtonClass is a representation of the C record GtkButtonClass.
type ButtonClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkButtonClass that represents the ButtonClass.
func (recv *ButtonClass) ToC() unsafe.Pointer {
	return recv.native
}

// ButtonClassNewFromC creates a new ButtonClass from a pointer to the C GtkButtonClass that represents the ButtonClass.
func ButtonClassNewFromC(native unsafe.Pointer) *ButtonClass {
	return &ButtonClass{native: native}
}

// ButtonPrivate is a representation of the C record GtkButtonPrivate.
type ButtonPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkButtonPrivate that represents the ButtonPrivate.
func (recv *ButtonPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ButtonPrivateNewFromC creates a new ButtonPrivate from a pointer to the C GtkButtonPrivate that represents the ButtonPrivate.
func ButtonPrivateNewFromC(native unsafe.Pointer) *ButtonPrivate {
	return &ButtonPrivate{native: native}
}

// CalendarClass is a representation of the C record GtkCalendarClass.
type CalendarClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkCalendarClass that represents the CalendarClass.
func (recv *CalendarClass) ToC() unsafe.Pointer {
	return recv.native
}

// CalendarClassNewFromC creates a new CalendarClass from a pointer to the C GtkCalendarClass that represents the CalendarClass.
func CalendarClassNewFromC(native unsafe.Pointer) *CalendarClass {
	return &CalendarClass{native: native}
}

// CalendarPrivate is a representation of the C record GtkCalendarPrivate.
type CalendarPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkCalendarPrivate that represents the CalendarPrivate.
func (recv *CalendarPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// CalendarPrivateNewFromC creates a new CalendarPrivate from a pointer to the C GtkCalendarPrivate that represents the CalendarPrivate.
func CalendarPrivateNewFromC(native unsafe.Pointer) *CalendarPrivate {
	return &CalendarPrivate{native: native}
}

// CellAccessibleClass is a representation of the C record GtkCellAccessibleClass.
type CellAccessibleClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkCellAccessibleClass that represents the CellAccessibleClass.
func (recv *CellAccessibleClass) ToC() unsafe.Pointer {
	return recv.native
}

// CellAccessibleClassNewFromC creates a new CellAccessibleClass from a pointer to the C GtkCellAccessibleClass that represents the CellAccessibleClass.
func CellAccessibleClassNewFromC(native unsafe.Pointer) *CellAccessibleClass {
	return &CellAccessibleClass{native: native}
}

// CellAccessibleParentIface is a representation of the C record GtkCellAccessibleParentIface.
type CellAccessibleParentIface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkCellAccessibleParentIface that represents the CellAccessibleParentIface.
func (recv *CellAccessibleParentIface) ToC() unsafe.Pointer {
	return recv.native
}

// CellAccessibleParentIfaceNewFromC creates a new CellAccessibleParentIface from a pointer to the C GtkCellAccessibleParentIface that represents the CellAccessibleParentIface.
func CellAccessibleParentIfaceNewFromC(native unsafe.Pointer) *CellAccessibleParentIface {
	return &CellAccessibleParentIface{native: native}
}

// CellAccessiblePrivate is a representation of the C record GtkCellAccessiblePrivate.
type CellAccessiblePrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkCellAccessiblePrivate that represents the CellAccessiblePrivate.
func (recv *CellAccessiblePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// CellAccessiblePrivateNewFromC creates a new CellAccessiblePrivate from a pointer to the C GtkCellAccessiblePrivate that represents the CellAccessiblePrivate.
func CellAccessiblePrivateNewFromC(native unsafe.Pointer) *CellAccessiblePrivate {
	return &CellAccessiblePrivate{native: native}
}

// CellAreaBoxClass is a representation of the C record GtkCellAreaBoxClass.
type CellAreaBoxClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkCellAreaBoxClass that represents the CellAreaBoxClass.
func (recv *CellAreaBoxClass) ToC() unsafe.Pointer {
	return recv.native
}

// CellAreaBoxClassNewFromC creates a new CellAreaBoxClass from a pointer to the C GtkCellAreaBoxClass that represents the CellAreaBoxClass.
func CellAreaBoxClassNewFromC(native unsafe.Pointer) *CellAreaBoxClass {
	return &CellAreaBoxClass{native: native}
}

// CellAreaBoxPrivate is a representation of the C record GtkCellAreaBoxPrivate.
type CellAreaBoxPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkCellAreaBoxPrivate that represents the CellAreaBoxPrivate.
func (recv *CellAreaBoxPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// CellAreaBoxPrivateNewFromC creates a new CellAreaBoxPrivate from a pointer to the C GtkCellAreaBoxPrivate that represents the CellAreaBoxPrivate.
func CellAreaBoxPrivateNewFromC(native unsafe.Pointer) *CellAreaBoxPrivate {
	return &CellAreaBoxPrivate{native: native}
}

// CellAreaClass is a representation of the C record GtkCellAreaClass.
type CellAreaClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkCellAreaClass that represents the CellAreaClass.
func (recv *CellAreaClass) ToC() unsafe.Pointer {
	return recv.native
}

// CellAreaClassNewFromC creates a new CellAreaClass from a pointer to the C GtkCellAreaClass that represents the CellAreaClass.
func CellAreaClassNewFromC(native unsafe.Pointer) *CellAreaClass {
	return &CellAreaClass{native: native}
}

// CellAreaContextClass is a representation of the C record GtkCellAreaContextClass.
type CellAreaContextClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkCellAreaContextClass that represents the CellAreaContextClass.
func (recv *CellAreaContextClass) ToC() unsafe.Pointer {
	return recv.native
}

// CellAreaContextClassNewFromC creates a new CellAreaContextClass from a pointer to the C GtkCellAreaContextClass that represents the CellAreaContextClass.
func CellAreaContextClassNewFromC(native unsafe.Pointer) *CellAreaContextClass {
	return &CellAreaContextClass{native: native}
}

// CellAreaContextPrivate is a representation of the C record GtkCellAreaContextPrivate.
type CellAreaContextPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkCellAreaContextPrivate that represents the CellAreaContextPrivate.
func (recv *CellAreaContextPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// CellAreaContextPrivateNewFromC creates a new CellAreaContextPrivate from a pointer to the C GtkCellAreaContextPrivate that represents the CellAreaContextPrivate.
func CellAreaContextPrivateNewFromC(native unsafe.Pointer) *CellAreaContextPrivate {
	return &CellAreaContextPrivate{native: native}
}

// CellAreaPrivate is a representation of the C record GtkCellAreaPrivate.
type CellAreaPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkCellAreaPrivate that represents the CellAreaPrivate.
func (recv *CellAreaPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// CellAreaPrivateNewFromC creates a new CellAreaPrivate from a pointer to the C GtkCellAreaPrivate that represents the CellAreaPrivate.
func CellAreaPrivateNewFromC(native unsafe.Pointer) *CellAreaPrivate {
	return &CellAreaPrivate{native: native}
}

// CellEditableIface is a representation of the C record GtkCellEditableIface.
type CellEditableIface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkCellEditableIface that represents the CellEditableIface.
func (recv *CellEditableIface) ToC() unsafe.Pointer {
	return recv.native
}

// CellEditableIfaceNewFromC creates a new CellEditableIface from a pointer to the C GtkCellEditableIface that represents the CellEditableIface.
func CellEditableIfaceNewFromC(native unsafe.Pointer) *CellEditableIface {
	return &CellEditableIface{native: native}
}

// CellLayoutIface is a representation of the C record GtkCellLayoutIface.
type CellLayoutIface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkCellLayoutIface that represents the CellLayoutIface.
func (recv *CellLayoutIface) ToC() unsafe.Pointer {
	return recv.native
}

// CellLayoutIfaceNewFromC creates a new CellLayoutIface from a pointer to the C GtkCellLayoutIface that represents the CellLayoutIface.
func CellLayoutIfaceNewFromC(native unsafe.Pointer) *CellLayoutIface {
	return &CellLayoutIface{native: native}
}

// CellRendererAccelClass is a representation of the C record GtkCellRendererAccelClass.
type CellRendererAccelClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkCellRendererAccelClass that represents the CellRendererAccelClass.
func (recv *CellRendererAccelClass) ToC() unsafe.Pointer {
	return recv.native
}

// CellRendererAccelClassNewFromC creates a new CellRendererAccelClass from a pointer to the C GtkCellRendererAccelClass that represents the CellRendererAccelClass.
func CellRendererAccelClassNewFromC(native unsafe.Pointer) *CellRendererAccelClass {
	return &CellRendererAccelClass{native: native}
}

// CellRendererAccelPrivate is a representation of the C record GtkCellRendererAccelPrivate.
type CellRendererAccelPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkCellRendererAccelPrivate that represents the CellRendererAccelPrivate.
func (recv *CellRendererAccelPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// CellRendererAccelPrivateNewFromC creates a new CellRendererAccelPrivate from a pointer to the C GtkCellRendererAccelPrivate that represents the CellRendererAccelPrivate.
func CellRendererAccelPrivateNewFromC(native unsafe.Pointer) *CellRendererAccelPrivate {
	return &CellRendererAccelPrivate{native: native}
}

// CellRendererClass is a representation of the C record GtkCellRendererClass.
type CellRendererClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkCellRendererClass that represents the CellRendererClass.
func (recv *CellRendererClass) ToC() unsafe.Pointer {
	return recv.native
}

// CellRendererClassNewFromC creates a new CellRendererClass from a pointer to the C GtkCellRendererClass that represents the CellRendererClass.
func CellRendererClassNewFromC(native unsafe.Pointer) *CellRendererClass {
	return &CellRendererClass{native: native}
}

// CellRendererClassPrivate is a representation of the C record GtkCellRendererClassPrivate.
type CellRendererClassPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkCellRendererClassPrivate that represents the CellRendererClassPrivate.
func (recv *CellRendererClassPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// CellRendererClassPrivateNewFromC creates a new CellRendererClassPrivate from a pointer to the C GtkCellRendererClassPrivate that represents the CellRendererClassPrivate.
func CellRendererClassPrivateNewFromC(native unsafe.Pointer) *CellRendererClassPrivate {
	return &CellRendererClassPrivate{native: native}
}

// CellRendererComboClass is a representation of the C record GtkCellRendererComboClass.
type CellRendererComboClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkCellRendererComboClass that represents the CellRendererComboClass.
func (recv *CellRendererComboClass) ToC() unsafe.Pointer {
	return recv.native
}

// CellRendererComboClassNewFromC creates a new CellRendererComboClass from a pointer to the C GtkCellRendererComboClass that represents the CellRendererComboClass.
func CellRendererComboClassNewFromC(native unsafe.Pointer) *CellRendererComboClass {
	return &CellRendererComboClass{native: native}
}

// CellRendererComboPrivate is a representation of the C record GtkCellRendererComboPrivate.
type CellRendererComboPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkCellRendererComboPrivate that represents the CellRendererComboPrivate.
func (recv *CellRendererComboPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// CellRendererComboPrivateNewFromC creates a new CellRendererComboPrivate from a pointer to the C GtkCellRendererComboPrivate that represents the CellRendererComboPrivate.
func CellRendererComboPrivateNewFromC(native unsafe.Pointer) *CellRendererComboPrivate {
	return &CellRendererComboPrivate{native: native}
}

// CellRendererPixbufClass is a representation of the C record GtkCellRendererPixbufClass.
type CellRendererPixbufClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkCellRendererPixbufClass that represents the CellRendererPixbufClass.
func (recv *CellRendererPixbufClass) ToC() unsafe.Pointer {
	return recv.native
}

// CellRendererPixbufClassNewFromC creates a new CellRendererPixbufClass from a pointer to the C GtkCellRendererPixbufClass that represents the CellRendererPixbufClass.
func CellRendererPixbufClassNewFromC(native unsafe.Pointer) *CellRendererPixbufClass {
	return &CellRendererPixbufClass{native: native}
}

// CellRendererPixbufPrivate is a representation of the C record GtkCellRendererPixbufPrivate.
type CellRendererPixbufPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkCellRendererPixbufPrivate that represents the CellRendererPixbufPrivate.
func (recv *CellRendererPixbufPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// CellRendererPixbufPrivateNewFromC creates a new CellRendererPixbufPrivate from a pointer to the C GtkCellRendererPixbufPrivate that represents the CellRendererPixbufPrivate.
func CellRendererPixbufPrivateNewFromC(native unsafe.Pointer) *CellRendererPixbufPrivate {
	return &CellRendererPixbufPrivate{native: native}
}

// CellRendererPrivate is a representation of the C record GtkCellRendererPrivate.
type CellRendererPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkCellRendererPrivate that represents the CellRendererPrivate.
func (recv *CellRendererPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// CellRendererPrivateNewFromC creates a new CellRendererPrivate from a pointer to the C GtkCellRendererPrivate that represents the CellRendererPrivate.
func CellRendererPrivateNewFromC(native unsafe.Pointer) *CellRendererPrivate {
	return &CellRendererPrivate{native: native}
}

// CellRendererProgressClass is a representation of the C record GtkCellRendererProgressClass.
type CellRendererProgressClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkCellRendererProgressClass that represents the CellRendererProgressClass.
func (recv *CellRendererProgressClass) ToC() unsafe.Pointer {
	return recv.native
}

// CellRendererProgressClassNewFromC creates a new CellRendererProgressClass from a pointer to the C GtkCellRendererProgressClass that represents the CellRendererProgressClass.
func CellRendererProgressClassNewFromC(native unsafe.Pointer) *CellRendererProgressClass {
	return &CellRendererProgressClass{native: native}
}

// CellRendererProgressPrivate is a representation of the C record GtkCellRendererProgressPrivate.
type CellRendererProgressPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkCellRendererProgressPrivate that represents the CellRendererProgressPrivate.
func (recv *CellRendererProgressPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// CellRendererProgressPrivateNewFromC creates a new CellRendererProgressPrivate from a pointer to the C GtkCellRendererProgressPrivate that represents the CellRendererProgressPrivate.
func CellRendererProgressPrivateNewFromC(native unsafe.Pointer) *CellRendererProgressPrivate {
	return &CellRendererProgressPrivate{native: native}
}

// CellRendererSpinClass is a representation of the C record GtkCellRendererSpinClass.
type CellRendererSpinClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkCellRendererSpinClass that represents the CellRendererSpinClass.
func (recv *CellRendererSpinClass) ToC() unsafe.Pointer {
	return recv.native
}

// CellRendererSpinClassNewFromC creates a new CellRendererSpinClass from a pointer to the C GtkCellRendererSpinClass that represents the CellRendererSpinClass.
func CellRendererSpinClassNewFromC(native unsafe.Pointer) *CellRendererSpinClass {
	return &CellRendererSpinClass{native: native}
}

// CellRendererSpinPrivate is a representation of the C record GtkCellRendererSpinPrivate.
type CellRendererSpinPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkCellRendererSpinPrivate that represents the CellRendererSpinPrivate.
func (recv *CellRendererSpinPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// CellRendererSpinPrivateNewFromC creates a new CellRendererSpinPrivate from a pointer to the C GtkCellRendererSpinPrivate that represents the CellRendererSpinPrivate.
func CellRendererSpinPrivateNewFromC(native unsafe.Pointer) *CellRendererSpinPrivate {
	return &CellRendererSpinPrivate{native: native}
}

// CellRendererSpinnerClass is a representation of the C record GtkCellRendererSpinnerClass.
type CellRendererSpinnerClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkCellRendererSpinnerClass that represents the CellRendererSpinnerClass.
func (recv *CellRendererSpinnerClass) ToC() unsafe.Pointer {
	return recv.native
}

// CellRendererSpinnerClassNewFromC creates a new CellRendererSpinnerClass from a pointer to the C GtkCellRendererSpinnerClass that represents the CellRendererSpinnerClass.
func CellRendererSpinnerClassNewFromC(native unsafe.Pointer) *CellRendererSpinnerClass {
	return &CellRendererSpinnerClass{native: native}
}

// CellRendererSpinnerPrivate is a representation of the C record GtkCellRendererSpinnerPrivate.
type CellRendererSpinnerPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkCellRendererSpinnerPrivate that represents the CellRendererSpinnerPrivate.
func (recv *CellRendererSpinnerPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// CellRendererSpinnerPrivateNewFromC creates a new CellRendererSpinnerPrivate from a pointer to the C GtkCellRendererSpinnerPrivate that represents the CellRendererSpinnerPrivate.
func CellRendererSpinnerPrivateNewFromC(native unsafe.Pointer) *CellRendererSpinnerPrivate {
	return &CellRendererSpinnerPrivate{native: native}
}

// CellRendererTextClass is a representation of the C record GtkCellRendererTextClass.
type CellRendererTextClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkCellRendererTextClass that represents the CellRendererTextClass.
func (recv *CellRendererTextClass) ToC() unsafe.Pointer {
	return recv.native
}

// CellRendererTextClassNewFromC creates a new CellRendererTextClass from a pointer to the C GtkCellRendererTextClass that represents the CellRendererTextClass.
func CellRendererTextClassNewFromC(native unsafe.Pointer) *CellRendererTextClass {
	return &CellRendererTextClass{native: native}
}

// CellRendererTextPrivate is a representation of the C record GtkCellRendererTextPrivate.
type CellRendererTextPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkCellRendererTextPrivate that represents the CellRendererTextPrivate.
func (recv *CellRendererTextPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// CellRendererTextPrivateNewFromC creates a new CellRendererTextPrivate from a pointer to the C GtkCellRendererTextPrivate that represents the CellRendererTextPrivate.
func CellRendererTextPrivateNewFromC(native unsafe.Pointer) *CellRendererTextPrivate {
	return &CellRendererTextPrivate{native: native}
}

// CellRendererToggleClass is a representation of the C record GtkCellRendererToggleClass.
type CellRendererToggleClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkCellRendererToggleClass that represents the CellRendererToggleClass.
func (recv *CellRendererToggleClass) ToC() unsafe.Pointer {
	return recv.native
}

// CellRendererToggleClassNewFromC creates a new CellRendererToggleClass from a pointer to the C GtkCellRendererToggleClass that represents the CellRendererToggleClass.
func CellRendererToggleClassNewFromC(native unsafe.Pointer) *CellRendererToggleClass {
	return &CellRendererToggleClass{native: native}
}

// CellRendererTogglePrivate is a representation of the C record GtkCellRendererTogglePrivate.
type CellRendererTogglePrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkCellRendererTogglePrivate that represents the CellRendererTogglePrivate.
func (recv *CellRendererTogglePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// CellRendererTogglePrivateNewFromC creates a new CellRendererTogglePrivate from a pointer to the C GtkCellRendererTogglePrivate that represents the CellRendererTogglePrivate.
func CellRendererTogglePrivateNewFromC(native unsafe.Pointer) *CellRendererTogglePrivate {
	return &CellRendererTogglePrivate{native: native}
}

// CellViewClass is a representation of the C record GtkCellViewClass.
type CellViewClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkCellViewClass that represents the CellViewClass.
func (recv *CellViewClass) ToC() unsafe.Pointer {
	return recv.native
}

// CellViewClassNewFromC creates a new CellViewClass from a pointer to the C GtkCellViewClass that represents the CellViewClass.
func CellViewClassNewFromC(native unsafe.Pointer) *CellViewClass {
	return &CellViewClass{native: native}
}

// CellViewPrivate is a representation of the C record GtkCellViewPrivate.
type CellViewPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkCellViewPrivate that represents the CellViewPrivate.
func (recv *CellViewPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// CellViewPrivateNewFromC creates a new CellViewPrivate from a pointer to the C GtkCellViewPrivate that represents the CellViewPrivate.
func CellViewPrivateNewFromC(native unsafe.Pointer) *CellViewPrivate {
	return &CellViewPrivate{native: native}
}

// CheckButtonClass is a representation of the C record GtkCheckButtonClass.
type CheckButtonClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkCheckButtonClass that represents the CheckButtonClass.
func (recv *CheckButtonClass) ToC() unsafe.Pointer {
	return recv.native
}

// CheckButtonClassNewFromC creates a new CheckButtonClass from a pointer to the C GtkCheckButtonClass that represents the CheckButtonClass.
func CheckButtonClassNewFromC(native unsafe.Pointer) *CheckButtonClass {
	return &CheckButtonClass{native: native}
}

// CheckMenuItemAccessibleClass is a representation of the C record GtkCheckMenuItemAccessibleClass.
type CheckMenuItemAccessibleClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkCheckMenuItemAccessibleClass that represents the CheckMenuItemAccessibleClass.
func (recv *CheckMenuItemAccessibleClass) ToC() unsafe.Pointer {
	return recv.native
}

// CheckMenuItemAccessibleClassNewFromC creates a new CheckMenuItemAccessibleClass from a pointer to the C GtkCheckMenuItemAccessibleClass that represents the CheckMenuItemAccessibleClass.
func CheckMenuItemAccessibleClassNewFromC(native unsafe.Pointer) *CheckMenuItemAccessibleClass {
	return &CheckMenuItemAccessibleClass{native: native}
}

// CheckMenuItemAccessiblePrivate is a representation of the C record GtkCheckMenuItemAccessiblePrivate.
type CheckMenuItemAccessiblePrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkCheckMenuItemAccessiblePrivate that represents the CheckMenuItemAccessiblePrivate.
func (recv *CheckMenuItemAccessiblePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// CheckMenuItemAccessiblePrivateNewFromC creates a new CheckMenuItemAccessiblePrivate from a pointer to the C GtkCheckMenuItemAccessiblePrivate that represents the CheckMenuItemAccessiblePrivate.
func CheckMenuItemAccessiblePrivateNewFromC(native unsafe.Pointer) *CheckMenuItemAccessiblePrivate {
	return &CheckMenuItemAccessiblePrivate{native: native}
}

// CheckMenuItemClass is a representation of the C record GtkCheckMenuItemClass.
type CheckMenuItemClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkCheckMenuItemClass that represents the CheckMenuItemClass.
func (recv *CheckMenuItemClass) ToC() unsafe.Pointer {
	return recv.native
}

// CheckMenuItemClassNewFromC creates a new CheckMenuItemClass from a pointer to the C GtkCheckMenuItemClass that represents the CheckMenuItemClass.
func CheckMenuItemClassNewFromC(native unsafe.Pointer) *CheckMenuItemClass {
	return &CheckMenuItemClass{native: native}
}

// CheckMenuItemPrivate is a representation of the C record GtkCheckMenuItemPrivate.
type CheckMenuItemPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkCheckMenuItemPrivate that represents the CheckMenuItemPrivate.
func (recv *CheckMenuItemPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// CheckMenuItemPrivateNewFromC creates a new CheckMenuItemPrivate from a pointer to the C GtkCheckMenuItemPrivate that represents the CheckMenuItemPrivate.
func CheckMenuItemPrivateNewFromC(native unsafe.Pointer) *CheckMenuItemPrivate {
	return &CheckMenuItemPrivate{native: native}
}

// ColorButtonClass is a representation of the C record GtkColorButtonClass.
type ColorButtonClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkColorButtonClass that represents the ColorButtonClass.
func (recv *ColorButtonClass) ToC() unsafe.Pointer {
	return recv.native
}

// ColorButtonClassNewFromC creates a new ColorButtonClass from a pointer to the C GtkColorButtonClass that represents the ColorButtonClass.
func ColorButtonClassNewFromC(native unsafe.Pointer) *ColorButtonClass {
	return &ColorButtonClass{native: native}
}

// ColorButtonPrivate is a representation of the C record GtkColorButtonPrivate.
type ColorButtonPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkColorButtonPrivate that represents the ColorButtonPrivate.
func (recv *ColorButtonPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ColorButtonPrivateNewFromC creates a new ColorButtonPrivate from a pointer to the C GtkColorButtonPrivate that represents the ColorButtonPrivate.
func ColorButtonPrivateNewFromC(native unsafe.Pointer) *ColorButtonPrivate {
	return &ColorButtonPrivate{native: native}
}

// ColorChooserDialogClass is a representation of the C record GtkColorChooserDialogClass.
type ColorChooserDialogClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkColorChooserDialogClass that represents the ColorChooserDialogClass.
func (recv *ColorChooserDialogClass) ToC() unsafe.Pointer {
	return recv.native
}

// ColorChooserDialogClassNewFromC creates a new ColorChooserDialogClass from a pointer to the C GtkColorChooserDialogClass that represents the ColorChooserDialogClass.
func ColorChooserDialogClassNewFromC(native unsafe.Pointer) *ColorChooserDialogClass {
	return &ColorChooserDialogClass{native: native}
}

// ColorChooserDialogPrivate is a representation of the C record GtkColorChooserDialogPrivate.
type ColorChooserDialogPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkColorChooserDialogPrivate that represents the ColorChooserDialogPrivate.
func (recv *ColorChooserDialogPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ColorChooserDialogPrivateNewFromC creates a new ColorChooserDialogPrivate from a pointer to the C GtkColorChooserDialogPrivate that represents the ColorChooserDialogPrivate.
func ColorChooserDialogPrivateNewFromC(native unsafe.Pointer) *ColorChooserDialogPrivate {
	return &ColorChooserDialogPrivate{native: native}
}

// ColorChooserInterface is a representation of the C record GtkColorChooserInterface.
type ColorChooserInterface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkColorChooserInterface that represents the ColorChooserInterface.
func (recv *ColorChooserInterface) ToC() unsafe.Pointer {
	return recv.native
}

// ColorChooserInterfaceNewFromC creates a new ColorChooserInterface from a pointer to the C GtkColorChooserInterface that represents the ColorChooserInterface.
func ColorChooserInterfaceNewFromC(native unsafe.Pointer) *ColorChooserInterface {
	return &ColorChooserInterface{native: native}
}

// ColorChooserWidgetClass is a representation of the C record GtkColorChooserWidgetClass.
type ColorChooserWidgetClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkColorChooserWidgetClass that represents the ColorChooserWidgetClass.
func (recv *ColorChooserWidgetClass) ToC() unsafe.Pointer {
	return recv.native
}

// ColorChooserWidgetClassNewFromC creates a new ColorChooserWidgetClass from a pointer to the C GtkColorChooserWidgetClass that represents the ColorChooserWidgetClass.
func ColorChooserWidgetClassNewFromC(native unsafe.Pointer) *ColorChooserWidgetClass {
	return &ColorChooserWidgetClass{native: native}
}

// ColorChooserWidgetPrivate is a representation of the C record GtkColorChooserWidgetPrivate.
type ColorChooserWidgetPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkColorChooserWidgetPrivate that represents the ColorChooserWidgetPrivate.
func (recv *ColorChooserWidgetPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ColorChooserWidgetPrivateNewFromC creates a new ColorChooserWidgetPrivate from a pointer to the C GtkColorChooserWidgetPrivate that represents the ColorChooserWidgetPrivate.
func ColorChooserWidgetPrivateNewFromC(native unsafe.Pointer) *ColorChooserWidgetPrivate {
	return &ColorChooserWidgetPrivate{native: native}
}

// ColorSelectionClass is a representation of the C record GtkColorSelectionClass.
type ColorSelectionClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkColorSelectionClass that represents the ColorSelectionClass.
func (recv *ColorSelectionClass) ToC() unsafe.Pointer {
	return recv.native
}

// ColorSelectionClassNewFromC creates a new ColorSelectionClass from a pointer to the C GtkColorSelectionClass that represents the ColorSelectionClass.
func ColorSelectionClassNewFromC(native unsafe.Pointer) *ColorSelectionClass {
	return &ColorSelectionClass{native: native}
}

// ColorSelectionDialogClass is a representation of the C record GtkColorSelectionDialogClass.
type ColorSelectionDialogClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkColorSelectionDialogClass that represents the ColorSelectionDialogClass.
func (recv *ColorSelectionDialogClass) ToC() unsafe.Pointer {
	return recv.native
}

// ColorSelectionDialogClassNewFromC creates a new ColorSelectionDialogClass from a pointer to the C GtkColorSelectionDialogClass that represents the ColorSelectionDialogClass.
func ColorSelectionDialogClassNewFromC(native unsafe.Pointer) *ColorSelectionDialogClass {
	return &ColorSelectionDialogClass{native: native}
}

// ColorSelectionDialogPrivate is a representation of the C record GtkColorSelectionDialogPrivate.
type ColorSelectionDialogPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkColorSelectionDialogPrivate that represents the ColorSelectionDialogPrivate.
func (recv *ColorSelectionDialogPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ColorSelectionDialogPrivateNewFromC creates a new ColorSelectionDialogPrivate from a pointer to the C GtkColorSelectionDialogPrivate that represents the ColorSelectionDialogPrivate.
func ColorSelectionDialogPrivateNewFromC(native unsafe.Pointer) *ColorSelectionDialogPrivate {
	return &ColorSelectionDialogPrivate{native: native}
}

// ColorSelectionPrivate is a representation of the C record GtkColorSelectionPrivate.
type ColorSelectionPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkColorSelectionPrivate that represents the ColorSelectionPrivate.
func (recv *ColorSelectionPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ColorSelectionPrivateNewFromC creates a new ColorSelectionPrivate from a pointer to the C GtkColorSelectionPrivate that represents the ColorSelectionPrivate.
func ColorSelectionPrivateNewFromC(native unsafe.Pointer) *ColorSelectionPrivate {
	return &ColorSelectionPrivate{native: native}
}

// ComboBoxAccessibleClass is a representation of the C record GtkComboBoxAccessibleClass.
type ComboBoxAccessibleClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkComboBoxAccessibleClass that represents the ComboBoxAccessibleClass.
func (recv *ComboBoxAccessibleClass) ToC() unsafe.Pointer {
	return recv.native
}

// ComboBoxAccessibleClassNewFromC creates a new ComboBoxAccessibleClass from a pointer to the C GtkComboBoxAccessibleClass that represents the ComboBoxAccessibleClass.
func ComboBoxAccessibleClassNewFromC(native unsafe.Pointer) *ComboBoxAccessibleClass {
	return &ComboBoxAccessibleClass{native: native}
}

// ComboBoxAccessiblePrivate is a representation of the C record GtkComboBoxAccessiblePrivate.
type ComboBoxAccessiblePrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkComboBoxAccessiblePrivate that represents the ComboBoxAccessiblePrivate.
func (recv *ComboBoxAccessiblePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ComboBoxAccessiblePrivateNewFromC creates a new ComboBoxAccessiblePrivate from a pointer to the C GtkComboBoxAccessiblePrivate that represents the ComboBoxAccessiblePrivate.
func ComboBoxAccessiblePrivateNewFromC(native unsafe.Pointer) *ComboBoxAccessiblePrivate {
	return &ComboBoxAccessiblePrivate{native: native}
}

// ComboBoxClass is a representation of the C record GtkComboBoxClass.
type ComboBoxClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkComboBoxClass that represents the ComboBoxClass.
func (recv *ComboBoxClass) ToC() unsafe.Pointer {
	return recv.native
}

// ComboBoxClassNewFromC creates a new ComboBoxClass from a pointer to the C GtkComboBoxClass that represents the ComboBoxClass.
func ComboBoxClassNewFromC(native unsafe.Pointer) *ComboBoxClass {
	return &ComboBoxClass{native: native}
}

// ComboBoxPrivate is a representation of the C record GtkComboBoxPrivate.
type ComboBoxPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkComboBoxPrivate that represents the ComboBoxPrivate.
func (recv *ComboBoxPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ComboBoxPrivateNewFromC creates a new ComboBoxPrivate from a pointer to the C GtkComboBoxPrivate that represents the ComboBoxPrivate.
func ComboBoxPrivateNewFromC(native unsafe.Pointer) *ComboBoxPrivate {
	return &ComboBoxPrivate{native: native}
}

// ComboBoxTextClass is a representation of the C record GtkComboBoxTextClass.
type ComboBoxTextClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkComboBoxTextClass that represents the ComboBoxTextClass.
func (recv *ComboBoxTextClass) ToC() unsafe.Pointer {
	return recv.native
}

// ComboBoxTextClassNewFromC creates a new ComboBoxTextClass from a pointer to the C GtkComboBoxTextClass that represents the ComboBoxTextClass.
func ComboBoxTextClassNewFromC(native unsafe.Pointer) *ComboBoxTextClass {
	return &ComboBoxTextClass{native: native}
}

// ComboBoxTextPrivate is a representation of the C record GtkComboBoxTextPrivate.
type ComboBoxTextPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkComboBoxTextPrivate that represents the ComboBoxTextPrivate.
func (recv *ComboBoxTextPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ComboBoxTextPrivateNewFromC creates a new ComboBoxTextPrivate from a pointer to the C GtkComboBoxTextPrivate that represents the ComboBoxTextPrivate.
func ComboBoxTextPrivateNewFromC(native unsafe.Pointer) *ComboBoxTextPrivate {
	return &ComboBoxTextPrivate{native: native}
}

// ContainerAccessibleClass is a representation of the C record GtkContainerAccessibleClass.
type ContainerAccessibleClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkContainerAccessibleClass that represents the ContainerAccessibleClass.
func (recv *ContainerAccessibleClass) ToC() unsafe.Pointer {
	return recv.native
}

// ContainerAccessibleClassNewFromC creates a new ContainerAccessibleClass from a pointer to the C GtkContainerAccessibleClass that represents the ContainerAccessibleClass.
func ContainerAccessibleClassNewFromC(native unsafe.Pointer) *ContainerAccessibleClass {
	return &ContainerAccessibleClass{native: native}
}

// ContainerAccessiblePrivate is a representation of the C record GtkContainerAccessiblePrivate.
type ContainerAccessiblePrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkContainerAccessiblePrivate that represents the ContainerAccessiblePrivate.
func (recv *ContainerAccessiblePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ContainerAccessiblePrivateNewFromC creates a new ContainerAccessiblePrivate from a pointer to the C GtkContainerAccessiblePrivate that represents the ContainerAccessiblePrivate.
func ContainerAccessiblePrivateNewFromC(native unsafe.Pointer) *ContainerAccessiblePrivate {
	return &ContainerAccessiblePrivate{native: native}
}

// ContainerCellAccessibleClass is a representation of the C record GtkContainerCellAccessibleClass.
type ContainerCellAccessibleClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkContainerCellAccessibleClass that represents the ContainerCellAccessibleClass.
func (recv *ContainerCellAccessibleClass) ToC() unsafe.Pointer {
	return recv.native
}

// ContainerCellAccessibleClassNewFromC creates a new ContainerCellAccessibleClass from a pointer to the C GtkContainerCellAccessibleClass that represents the ContainerCellAccessibleClass.
func ContainerCellAccessibleClassNewFromC(native unsafe.Pointer) *ContainerCellAccessibleClass {
	return &ContainerCellAccessibleClass{native: native}
}

// ContainerCellAccessiblePrivate is a representation of the C record GtkContainerCellAccessiblePrivate.
type ContainerCellAccessiblePrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkContainerCellAccessiblePrivate that represents the ContainerCellAccessiblePrivate.
func (recv *ContainerCellAccessiblePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ContainerCellAccessiblePrivateNewFromC creates a new ContainerCellAccessiblePrivate from a pointer to the C GtkContainerCellAccessiblePrivate that represents the ContainerCellAccessiblePrivate.
func ContainerCellAccessiblePrivateNewFromC(native unsafe.Pointer) *ContainerCellAccessiblePrivate {
	return &ContainerCellAccessiblePrivate{native: native}
}

// ContainerClass is a representation of the C record GtkContainerClass.
type ContainerClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkContainerClass that represents the ContainerClass.
func (recv *ContainerClass) ToC() unsafe.Pointer {
	return recv.native
}

// ContainerClassNewFromC creates a new ContainerClass from a pointer to the C GtkContainerClass that represents the ContainerClass.
func ContainerClassNewFromC(native unsafe.Pointer) *ContainerClass {
	return &ContainerClass{native: native}
}

// ContainerPrivate is a representation of the C record GtkContainerPrivate.
type ContainerPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkContainerPrivate that represents the ContainerPrivate.
func (recv *ContainerPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ContainerPrivateNewFromC creates a new ContainerPrivate from a pointer to the C GtkContainerPrivate that represents the ContainerPrivate.
func ContainerPrivateNewFromC(native unsafe.Pointer) *ContainerPrivate {
	return &ContainerPrivate{native: native}
}

// CssProviderClass is a representation of the C record GtkCssProviderClass.
type CssProviderClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkCssProviderClass that represents the CssProviderClass.
func (recv *CssProviderClass) ToC() unsafe.Pointer {
	return recv.native
}

// CssProviderClassNewFromC creates a new CssProviderClass from a pointer to the C GtkCssProviderClass that represents the CssProviderClass.
func CssProviderClassNewFromC(native unsafe.Pointer) *CssProviderClass {
	return &CssProviderClass{native: native}
}

// CssProviderPrivate is a representation of the C record GtkCssProviderPrivate.
type CssProviderPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkCssProviderPrivate that represents the CssProviderPrivate.
func (recv *CssProviderPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// CssProviderPrivateNewFromC creates a new CssProviderPrivate from a pointer to the C GtkCssProviderPrivate that represents the CssProviderPrivate.
func CssProviderPrivateNewFromC(native unsafe.Pointer) *CssProviderPrivate {
	return &CssProviderPrivate{native: native}
}

// DialogClass is a representation of the C record GtkDialogClass.
type DialogClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkDialogClass that represents the DialogClass.
func (recv *DialogClass) ToC() unsafe.Pointer {
	return recv.native
}

// DialogClassNewFromC creates a new DialogClass from a pointer to the C GtkDialogClass that represents the DialogClass.
func DialogClassNewFromC(native unsafe.Pointer) *DialogClass {
	return &DialogClass{native: native}
}

// DialogPrivate is a representation of the C record GtkDialogPrivate.
type DialogPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkDialogPrivate that represents the DialogPrivate.
func (recv *DialogPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// DialogPrivateNewFromC creates a new DialogPrivate from a pointer to the C GtkDialogPrivate that represents the DialogPrivate.
func DialogPrivateNewFromC(native unsafe.Pointer) *DialogPrivate {
	return &DialogPrivate{native: native}
}

// DrawingAreaClass is a representation of the C record GtkDrawingAreaClass.
type DrawingAreaClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkDrawingAreaClass that represents the DrawingAreaClass.
func (recv *DrawingAreaClass) ToC() unsafe.Pointer {
	return recv.native
}

// DrawingAreaClassNewFromC creates a new DrawingAreaClass from a pointer to the C GtkDrawingAreaClass that represents the DrawingAreaClass.
func DrawingAreaClassNewFromC(native unsafe.Pointer) *DrawingAreaClass {
	return &DrawingAreaClass{native: native}
}

// EditableInterface is a representation of the C record GtkEditableInterface.
type EditableInterface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkEditableInterface that represents the EditableInterface.
func (recv *EditableInterface) ToC() unsafe.Pointer {
	return recv.native
}

// EditableInterfaceNewFromC creates a new EditableInterface from a pointer to the C GtkEditableInterface that represents the EditableInterface.
func EditableInterfaceNewFromC(native unsafe.Pointer) *EditableInterface {
	return &EditableInterface{native: native}
}

// EntryAccessibleClass is a representation of the C record GtkEntryAccessibleClass.
type EntryAccessibleClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkEntryAccessibleClass that represents the EntryAccessibleClass.
func (recv *EntryAccessibleClass) ToC() unsafe.Pointer {
	return recv.native
}

// EntryAccessibleClassNewFromC creates a new EntryAccessibleClass from a pointer to the C GtkEntryAccessibleClass that represents the EntryAccessibleClass.
func EntryAccessibleClassNewFromC(native unsafe.Pointer) *EntryAccessibleClass {
	return &EntryAccessibleClass{native: native}
}

// EntryAccessiblePrivate is a representation of the C record GtkEntryAccessiblePrivate.
type EntryAccessiblePrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkEntryAccessiblePrivate that represents the EntryAccessiblePrivate.
func (recv *EntryAccessiblePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// EntryAccessiblePrivateNewFromC creates a new EntryAccessiblePrivate from a pointer to the C GtkEntryAccessiblePrivate that represents the EntryAccessiblePrivate.
func EntryAccessiblePrivateNewFromC(native unsafe.Pointer) *EntryAccessiblePrivate {
	return &EntryAccessiblePrivate{native: native}
}

// EntryBufferClass is a representation of the C record GtkEntryBufferClass.
type EntryBufferClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkEntryBufferClass that represents the EntryBufferClass.
func (recv *EntryBufferClass) ToC() unsafe.Pointer {
	return recv.native
}

// EntryBufferClassNewFromC creates a new EntryBufferClass from a pointer to the C GtkEntryBufferClass that represents the EntryBufferClass.
func EntryBufferClassNewFromC(native unsafe.Pointer) *EntryBufferClass {
	return &EntryBufferClass{native: native}
}

// EntryBufferPrivate is a representation of the C record GtkEntryBufferPrivate.
type EntryBufferPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkEntryBufferPrivate that represents the EntryBufferPrivate.
func (recv *EntryBufferPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// EntryBufferPrivateNewFromC creates a new EntryBufferPrivate from a pointer to the C GtkEntryBufferPrivate that represents the EntryBufferPrivate.
func EntryBufferPrivateNewFromC(native unsafe.Pointer) *EntryBufferPrivate {
	return &EntryBufferPrivate{native: native}
}

// EntryClass is a representation of the C record GtkEntryClass.
type EntryClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkEntryClass that represents the EntryClass.
func (recv *EntryClass) ToC() unsafe.Pointer {
	return recv.native
}

// EntryClassNewFromC creates a new EntryClass from a pointer to the C GtkEntryClass that represents the EntryClass.
func EntryClassNewFromC(native unsafe.Pointer) *EntryClass {
	return &EntryClass{native: native}
}

// EntryCompletionClass is a representation of the C record GtkEntryCompletionClass.
type EntryCompletionClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkEntryCompletionClass that represents the EntryCompletionClass.
func (recv *EntryCompletionClass) ToC() unsafe.Pointer {
	return recv.native
}

// EntryCompletionClassNewFromC creates a new EntryCompletionClass from a pointer to the C GtkEntryCompletionClass that represents the EntryCompletionClass.
func EntryCompletionClassNewFromC(native unsafe.Pointer) *EntryCompletionClass {
	return &EntryCompletionClass{native: native}
}

// EntryCompletionPrivate is a representation of the C record GtkEntryCompletionPrivate.
type EntryCompletionPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkEntryCompletionPrivate that represents the EntryCompletionPrivate.
func (recv *EntryCompletionPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// EntryCompletionPrivateNewFromC creates a new EntryCompletionPrivate from a pointer to the C GtkEntryCompletionPrivate that represents the EntryCompletionPrivate.
func EntryCompletionPrivateNewFromC(native unsafe.Pointer) *EntryCompletionPrivate {
	return &EntryCompletionPrivate{native: native}
}

// EntryPrivate is a representation of the C record GtkEntryPrivate.
type EntryPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkEntryPrivate that represents the EntryPrivate.
func (recv *EntryPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// EntryPrivateNewFromC creates a new EntryPrivate from a pointer to the C GtkEntryPrivate that represents the EntryPrivate.
func EntryPrivateNewFromC(native unsafe.Pointer) *EntryPrivate {
	return &EntryPrivate{native: native}
}

// EventBoxClass is a representation of the C record GtkEventBoxClass.
type EventBoxClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkEventBoxClass that represents the EventBoxClass.
func (recv *EventBoxClass) ToC() unsafe.Pointer {
	return recv.native
}

// EventBoxClassNewFromC creates a new EventBoxClass from a pointer to the C GtkEventBoxClass that represents the EventBoxClass.
func EventBoxClassNewFromC(native unsafe.Pointer) *EventBoxClass {
	return &EventBoxClass{native: native}
}

// EventBoxPrivate is a representation of the C record GtkEventBoxPrivate.
type EventBoxPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkEventBoxPrivate that represents the EventBoxPrivate.
func (recv *EventBoxPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// EventBoxPrivateNewFromC creates a new EventBoxPrivate from a pointer to the C GtkEventBoxPrivate that represents the EventBoxPrivate.
func EventBoxPrivateNewFromC(native unsafe.Pointer) *EventBoxPrivate {
	return &EventBoxPrivate{native: native}
}

// EventControllerClass is a representation of the C record GtkEventControllerClass.
type EventControllerClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkEventControllerClass that represents the EventControllerClass.
func (recv *EventControllerClass) ToC() unsafe.Pointer {
	return recv.native
}

// EventControllerClassNewFromC creates a new EventControllerClass from a pointer to the C GtkEventControllerClass that represents the EventControllerClass.
func EventControllerClassNewFromC(native unsafe.Pointer) *EventControllerClass {
	return &EventControllerClass{native: native}
}

// UNSUPPORTED : EventControllerMotionClass : blacklisted

// UNSUPPORTED : EventControllerScrollClass : blacklisted

// ExpanderAccessibleClass is a representation of the C record GtkExpanderAccessibleClass.
type ExpanderAccessibleClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkExpanderAccessibleClass that represents the ExpanderAccessibleClass.
func (recv *ExpanderAccessibleClass) ToC() unsafe.Pointer {
	return recv.native
}

// ExpanderAccessibleClassNewFromC creates a new ExpanderAccessibleClass from a pointer to the C GtkExpanderAccessibleClass that represents the ExpanderAccessibleClass.
func ExpanderAccessibleClassNewFromC(native unsafe.Pointer) *ExpanderAccessibleClass {
	return &ExpanderAccessibleClass{native: native}
}

// ExpanderAccessiblePrivate is a representation of the C record GtkExpanderAccessiblePrivate.
type ExpanderAccessiblePrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkExpanderAccessiblePrivate that represents the ExpanderAccessiblePrivate.
func (recv *ExpanderAccessiblePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ExpanderAccessiblePrivateNewFromC creates a new ExpanderAccessiblePrivate from a pointer to the C GtkExpanderAccessiblePrivate that represents the ExpanderAccessiblePrivate.
func ExpanderAccessiblePrivateNewFromC(native unsafe.Pointer) *ExpanderAccessiblePrivate {
	return &ExpanderAccessiblePrivate{native: native}
}

// ExpanderClass is a representation of the C record GtkExpanderClass.
type ExpanderClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkExpanderClass that represents the ExpanderClass.
func (recv *ExpanderClass) ToC() unsafe.Pointer {
	return recv.native
}

// ExpanderClassNewFromC creates a new ExpanderClass from a pointer to the C GtkExpanderClass that represents the ExpanderClass.
func ExpanderClassNewFromC(native unsafe.Pointer) *ExpanderClass {
	return &ExpanderClass{native: native}
}

// ExpanderPrivate is a representation of the C record GtkExpanderPrivate.
type ExpanderPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkExpanderPrivate that represents the ExpanderPrivate.
func (recv *ExpanderPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ExpanderPrivateNewFromC creates a new ExpanderPrivate from a pointer to the C GtkExpanderPrivate that represents the ExpanderPrivate.
func ExpanderPrivateNewFromC(native unsafe.Pointer) *ExpanderPrivate {
	return &ExpanderPrivate{native: native}
}

// FileChooserButtonClass is a representation of the C record GtkFileChooserButtonClass.
type FileChooserButtonClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkFileChooserButtonClass that represents the FileChooserButtonClass.
func (recv *FileChooserButtonClass) ToC() unsafe.Pointer {
	return recv.native
}

// FileChooserButtonClassNewFromC creates a new FileChooserButtonClass from a pointer to the C GtkFileChooserButtonClass that represents the FileChooserButtonClass.
func FileChooserButtonClassNewFromC(native unsafe.Pointer) *FileChooserButtonClass {
	return &FileChooserButtonClass{native: native}
}

// FileChooserButtonPrivate is a representation of the C record GtkFileChooserButtonPrivate.
type FileChooserButtonPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkFileChooserButtonPrivate that represents the FileChooserButtonPrivate.
func (recv *FileChooserButtonPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// FileChooserButtonPrivateNewFromC creates a new FileChooserButtonPrivate from a pointer to the C GtkFileChooserButtonPrivate that represents the FileChooserButtonPrivate.
func FileChooserButtonPrivateNewFromC(native unsafe.Pointer) *FileChooserButtonPrivate {
	return &FileChooserButtonPrivate{native: native}
}

// FileChooserDialogClass is a representation of the C record GtkFileChooserDialogClass.
type FileChooserDialogClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkFileChooserDialogClass that represents the FileChooserDialogClass.
func (recv *FileChooserDialogClass) ToC() unsafe.Pointer {
	return recv.native
}

// FileChooserDialogClassNewFromC creates a new FileChooserDialogClass from a pointer to the C GtkFileChooserDialogClass that represents the FileChooserDialogClass.
func FileChooserDialogClassNewFromC(native unsafe.Pointer) *FileChooserDialogClass {
	return &FileChooserDialogClass{native: native}
}

// FileChooserDialogPrivate is a representation of the C record GtkFileChooserDialogPrivate.
type FileChooserDialogPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkFileChooserDialogPrivate that represents the FileChooserDialogPrivate.
func (recv *FileChooserDialogPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// FileChooserDialogPrivateNewFromC creates a new FileChooserDialogPrivate from a pointer to the C GtkFileChooserDialogPrivate that represents the FileChooserDialogPrivate.
func FileChooserDialogPrivateNewFromC(native unsafe.Pointer) *FileChooserDialogPrivate {
	return &FileChooserDialogPrivate{native: native}
}

// FileChooserNativeClass is a representation of the C record GtkFileChooserNativeClass.
type FileChooserNativeClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkFileChooserNativeClass that represents the FileChooserNativeClass.
func (recv *FileChooserNativeClass) ToC() unsafe.Pointer {
	return recv.native
}

// FileChooserNativeClassNewFromC creates a new FileChooserNativeClass from a pointer to the C GtkFileChooserNativeClass that represents the FileChooserNativeClass.
func FileChooserNativeClassNewFromC(native unsafe.Pointer) *FileChooserNativeClass {
	return &FileChooserNativeClass{native: native}
}

// FileChooserWidgetClass is a representation of the C record GtkFileChooserWidgetClass.
type FileChooserWidgetClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkFileChooserWidgetClass that represents the FileChooserWidgetClass.
func (recv *FileChooserWidgetClass) ToC() unsafe.Pointer {
	return recv.native
}

// FileChooserWidgetClassNewFromC creates a new FileChooserWidgetClass from a pointer to the C GtkFileChooserWidgetClass that represents the FileChooserWidgetClass.
func FileChooserWidgetClassNewFromC(native unsafe.Pointer) *FileChooserWidgetClass {
	return &FileChooserWidgetClass{native: native}
}

// FileChooserWidgetPrivate is a representation of the C record GtkFileChooserWidgetPrivate.
type FileChooserWidgetPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkFileChooserWidgetPrivate that represents the FileChooserWidgetPrivate.
func (recv *FileChooserWidgetPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// FileChooserWidgetPrivateNewFromC creates a new FileChooserWidgetPrivate from a pointer to the C GtkFileChooserWidgetPrivate that represents the FileChooserWidgetPrivate.
func FileChooserWidgetPrivateNewFromC(native unsafe.Pointer) *FileChooserWidgetPrivate {
	return &FileChooserWidgetPrivate{native: native}
}

// FileFilterInfo is a representation of the C record GtkFileFilterInfo.
type FileFilterInfo struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkFileFilterInfo that represents the FileFilterInfo.
func (recv *FileFilterInfo) ToC() unsafe.Pointer {
	return recv.native
}

// FileFilterInfoNewFromC creates a new FileFilterInfo from a pointer to the C GtkFileFilterInfo that represents the FileFilterInfo.
func FileFilterInfoNewFromC(native unsafe.Pointer) *FileFilterInfo {
	return &FileFilterInfo{native: native}
}

// FixedChild is a representation of the C record GtkFixedChild.
type FixedChild struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkFixedChild that represents the FixedChild.
func (recv *FixedChild) ToC() unsafe.Pointer {
	return recv.native
}

// FixedChildNewFromC creates a new FixedChild from a pointer to the C GtkFixedChild that represents the FixedChild.
func FixedChildNewFromC(native unsafe.Pointer) *FixedChild {
	return &FixedChild{native: native}
}

// FixedClass is a representation of the C record GtkFixedClass.
type FixedClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkFixedClass that represents the FixedClass.
func (recv *FixedClass) ToC() unsafe.Pointer {
	return recv.native
}

// FixedClassNewFromC creates a new FixedClass from a pointer to the C GtkFixedClass that represents the FixedClass.
func FixedClassNewFromC(native unsafe.Pointer) *FixedClass {
	return &FixedClass{native: native}
}

// FixedPrivate is a representation of the C record GtkFixedPrivate.
type FixedPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkFixedPrivate that represents the FixedPrivate.
func (recv *FixedPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// FixedPrivateNewFromC creates a new FixedPrivate from a pointer to the C GtkFixedPrivate that represents the FixedPrivate.
func FixedPrivateNewFromC(native unsafe.Pointer) *FixedPrivate {
	return &FixedPrivate{native: native}
}

// FlowBoxAccessibleClass is a representation of the C record GtkFlowBoxAccessibleClass.
type FlowBoxAccessibleClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkFlowBoxAccessibleClass that represents the FlowBoxAccessibleClass.
func (recv *FlowBoxAccessibleClass) ToC() unsafe.Pointer {
	return recv.native
}

// FlowBoxAccessibleClassNewFromC creates a new FlowBoxAccessibleClass from a pointer to the C GtkFlowBoxAccessibleClass that represents the FlowBoxAccessibleClass.
func FlowBoxAccessibleClassNewFromC(native unsafe.Pointer) *FlowBoxAccessibleClass {
	return &FlowBoxAccessibleClass{native: native}
}

// FlowBoxAccessiblePrivate is a representation of the C record GtkFlowBoxAccessiblePrivate.
type FlowBoxAccessiblePrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkFlowBoxAccessiblePrivate that represents the FlowBoxAccessiblePrivate.
func (recv *FlowBoxAccessiblePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// FlowBoxAccessiblePrivateNewFromC creates a new FlowBoxAccessiblePrivate from a pointer to the C GtkFlowBoxAccessiblePrivate that represents the FlowBoxAccessiblePrivate.
func FlowBoxAccessiblePrivateNewFromC(native unsafe.Pointer) *FlowBoxAccessiblePrivate {
	return &FlowBoxAccessiblePrivate{native: native}
}

// FlowBoxChildAccessibleClass is a representation of the C record GtkFlowBoxChildAccessibleClass.
type FlowBoxChildAccessibleClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkFlowBoxChildAccessibleClass that represents the FlowBoxChildAccessibleClass.
func (recv *FlowBoxChildAccessibleClass) ToC() unsafe.Pointer {
	return recv.native
}

// FlowBoxChildAccessibleClassNewFromC creates a new FlowBoxChildAccessibleClass from a pointer to the C GtkFlowBoxChildAccessibleClass that represents the FlowBoxChildAccessibleClass.
func FlowBoxChildAccessibleClassNewFromC(native unsafe.Pointer) *FlowBoxChildAccessibleClass {
	return &FlowBoxChildAccessibleClass{native: native}
}

// FlowBoxChildClass is a representation of the C record GtkFlowBoxChildClass.
type FlowBoxChildClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkFlowBoxChildClass that represents the FlowBoxChildClass.
func (recv *FlowBoxChildClass) ToC() unsafe.Pointer {
	return recv.native
}

// FlowBoxChildClassNewFromC creates a new FlowBoxChildClass from a pointer to the C GtkFlowBoxChildClass that represents the FlowBoxChildClass.
func FlowBoxChildClassNewFromC(native unsafe.Pointer) *FlowBoxChildClass {
	return &FlowBoxChildClass{native: native}
}

// FlowBoxClass is a representation of the C record GtkFlowBoxClass.
type FlowBoxClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkFlowBoxClass that represents the FlowBoxClass.
func (recv *FlowBoxClass) ToC() unsafe.Pointer {
	return recv.native
}

// FlowBoxClassNewFromC creates a new FlowBoxClass from a pointer to the C GtkFlowBoxClass that represents the FlowBoxClass.
func FlowBoxClassNewFromC(native unsafe.Pointer) *FlowBoxClass {
	return &FlowBoxClass{native: native}
}

// FontButtonClass is a representation of the C record GtkFontButtonClass.
type FontButtonClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkFontButtonClass that represents the FontButtonClass.
func (recv *FontButtonClass) ToC() unsafe.Pointer {
	return recv.native
}

// FontButtonClassNewFromC creates a new FontButtonClass from a pointer to the C GtkFontButtonClass that represents the FontButtonClass.
func FontButtonClassNewFromC(native unsafe.Pointer) *FontButtonClass {
	return &FontButtonClass{native: native}
}

// FontButtonPrivate is a representation of the C record GtkFontButtonPrivate.
type FontButtonPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkFontButtonPrivate that represents the FontButtonPrivate.
func (recv *FontButtonPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// FontButtonPrivateNewFromC creates a new FontButtonPrivate from a pointer to the C GtkFontButtonPrivate that represents the FontButtonPrivate.
func FontButtonPrivateNewFromC(native unsafe.Pointer) *FontButtonPrivate {
	return &FontButtonPrivate{native: native}
}

// FontChooserDialogClass is a representation of the C record GtkFontChooserDialogClass.
type FontChooserDialogClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkFontChooserDialogClass that represents the FontChooserDialogClass.
func (recv *FontChooserDialogClass) ToC() unsafe.Pointer {
	return recv.native
}

// FontChooserDialogClassNewFromC creates a new FontChooserDialogClass from a pointer to the C GtkFontChooserDialogClass that represents the FontChooserDialogClass.
func FontChooserDialogClassNewFromC(native unsafe.Pointer) *FontChooserDialogClass {
	return &FontChooserDialogClass{native: native}
}

// FontChooserDialogPrivate is a representation of the C record GtkFontChooserDialogPrivate.
type FontChooserDialogPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkFontChooserDialogPrivate that represents the FontChooserDialogPrivate.
func (recv *FontChooserDialogPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// FontChooserDialogPrivateNewFromC creates a new FontChooserDialogPrivate from a pointer to the C GtkFontChooserDialogPrivate that represents the FontChooserDialogPrivate.
func FontChooserDialogPrivateNewFromC(native unsafe.Pointer) *FontChooserDialogPrivate {
	return &FontChooserDialogPrivate{native: native}
}

// FontChooserIface is a representation of the C record GtkFontChooserIface.
type FontChooserIface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkFontChooserIface that represents the FontChooserIface.
func (recv *FontChooserIface) ToC() unsafe.Pointer {
	return recv.native
}

// FontChooserIfaceNewFromC creates a new FontChooserIface from a pointer to the C GtkFontChooserIface that represents the FontChooserIface.
func FontChooserIfaceNewFromC(native unsafe.Pointer) *FontChooserIface {
	return &FontChooserIface{native: native}
}

// FontChooserWidgetClass is a representation of the C record GtkFontChooserWidgetClass.
type FontChooserWidgetClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkFontChooserWidgetClass that represents the FontChooserWidgetClass.
func (recv *FontChooserWidgetClass) ToC() unsafe.Pointer {
	return recv.native
}

// FontChooserWidgetClassNewFromC creates a new FontChooserWidgetClass from a pointer to the C GtkFontChooserWidgetClass that represents the FontChooserWidgetClass.
func FontChooserWidgetClassNewFromC(native unsafe.Pointer) *FontChooserWidgetClass {
	return &FontChooserWidgetClass{native: native}
}

// FontChooserWidgetPrivate is a representation of the C record GtkFontChooserWidgetPrivate.
type FontChooserWidgetPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkFontChooserWidgetPrivate that represents the FontChooserWidgetPrivate.
func (recv *FontChooserWidgetPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// FontChooserWidgetPrivateNewFromC creates a new FontChooserWidgetPrivate from a pointer to the C GtkFontChooserWidgetPrivate that represents the FontChooserWidgetPrivate.
func FontChooserWidgetPrivateNewFromC(native unsafe.Pointer) *FontChooserWidgetPrivate {
	return &FontChooserWidgetPrivate{native: native}
}

// FontSelectionClass is a representation of the C record GtkFontSelectionClass.
type FontSelectionClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkFontSelectionClass that represents the FontSelectionClass.
func (recv *FontSelectionClass) ToC() unsafe.Pointer {
	return recv.native
}

// FontSelectionClassNewFromC creates a new FontSelectionClass from a pointer to the C GtkFontSelectionClass that represents the FontSelectionClass.
func FontSelectionClassNewFromC(native unsafe.Pointer) *FontSelectionClass {
	return &FontSelectionClass{native: native}
}

// FontSelectionDialogClass is a representation of the C record GtkFontSelectionDialogClass.
type FontSelectionDialogClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkFontSelectionDialogClass that represents the FontSelectionDialogClass.
func (recv *FontSelectionDialogClass) ToC() unsafe.Pointer {
	return recv.native
}

// FontSelectionDialogClassNewFromC creates a new FontSelectionDialogClass from a pointer to the C GtkFontSelectionDialogClass that represents the FontSelectionDialogClass.
func FontSelectionDialogClassNewFromC(native unsafe.Pointer) *FontSelectionDialogClass {
	return &FontSelectionDialogClass{native: native}
}

// FontSelectionDialogPrivate is a representation of the C record GtkFontSelectionDialogPrivate.
type FontSelectionDialogPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkFontSelectionDialogPrivate that represents the FontSelectionDialogPrivate.
func (recv *FontSelectionDialogPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// FontSelectionDialogPrivateNewFromC creates a new FontSelectionDialogPrivate from a pointer to the C GtkFontSelectionDialogPrivate that represents the FontSelectionDialogPrivate.
func FontSelectionDialogPrivateNewFromC(native unsafe.Pointer) *FontSelectionDialogPrivate {
	return &FontSelectionDialogPrivate{native: native}
}

// FontSelectionPrivate is a representation of the C record GtkFontSelectionPrivate.
type FontSelectionPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkFontSelectionPrivate that represents the FontSelectionPrivate.
func (recv *FontSelectionPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// FontSelectionPrivateNewFromC creates a new FontSelectionPrivate from a pointer to the C GtkFontSelectionPrivate that represents the FontSelectionPrivate.
func FontSelectionPrivateNewFromC(native unsafe.Pointer) *FontSelectionPrivate {
	return &FontSelectionPrivate{native: native}
}

// FrameAccessibleClass is a representation of the C record GtkFrameAccessibleClass.
type FrameAccessibleClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkFrameAccessibleClass that represents the FrameAccessibleClass.
func (recv *FrameAccessibleClass) ToC() unsafe.Pointer {
	return recv.native
}

// FrameAccessibleClassNewFromC creates a new FrameAccessibleClass from a pointer to the C GtkFrameAccessibleClass that represents the FrameAccessibleClass.
func FrameAccessibleClassNewFromC(native unsafe.Pointer) *FrameAccessibleClass {
	return &FrameAccessibleClass{native: native}
}

// FrameAccessiblePrivate is a representation of the C record GtkFrameAccessiblePrivate.
type FrameAccessiblePrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkFrameAccessiblePrivate that represents the FrameAccessiblePrivate.
func (recv *FrameAccessiblePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// FrameAccessiblePrivateNewFromC creates a new FrameAccessiblePrivate from a pointer to the C GtkFrameAccessiblePrivate that represents the FrameAccessiblePrivate.
func FrameAccessiblePrivateNewFromC(native unsafe.Pointer) *FrameAccessiblePrivate {
	return &FrameAccessiblePrivate{native: native}
}

// FrameClass is a representation of the C record GtkFrameClass.
type FrameClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkFrameClass that represents the FrameClass.
func (recv *FrameClass) ToC() unsafe.Pointer {
	return recv.native
}

// FrameClassNewFromC creates a new FrameClass from a pointer to the C GtkFrameClass that represents the FrameClass.
func FrameClassNewFromC(native unsafe.Pointer) *FrameClass {
	return &FrameClass{native: native}
}

// FramePrivate is a representation of the C record GtkFramePrivate.
type FramePrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkFramePrivate that represents the FramePrivate.
func (recv *FramePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// FramePrivateNewFromC creates a new FramePrivate from a pointer to the C GtkFramePrivate that represents the FramePrivate.
func FramePrivateNewFromC(native unsafe.Pointer) *FramePrivate {
	return &FramePrivate{native: native}
}

// GestureClass is a representation of the C record GtkGestureClass.
type GestureClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkGestureClass that represents the GestureClass.
func (recv *GestureClass) ToC() unsafe.Pointer {
	return recv.native
}

// GestureClassNewFromC creates a new GestureClass from a pointer to the C GtkGestureClass that represents the GestureClass.
func GestureClassNewFromC(native unsafe.Pointer) *GestureClass {
	return &GestureClass{native: native}
}

// GestureDragClass is a representation of the C record GtkGestureDragClass.
type GestureDragClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkGestureDragClass that represents the GestureDragClass.
func (recv *GestureDragClass) ToC() unsafe.Pointer {
	return recv.native
}

// GestureDragClassNewFromC creates a new GestureDragClass from a pointer to the C GtkGestureDragClass that represents the GestureDragClass.
func GestureDragClassNewFromC(native unsafe.Pointer) *GestureDragClass {
	return &GestureDragClass{native: native}
}

// GestureLongPressClass is a representation of the C record GtkGestureLongPressClass.
type GestureLongPressClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkGestureLongPressClass that represents the GestureLongPressClass.
func (recv *GestureLongPressClass) ToC() unsafe.Pointer {
	return recv.native
}

// GestureLongPressClassNewFromC creates a new GestureLongPressClass from a pointer to the C GtkGestureLongPressClass that represents the GestureLongPressClass.
func GestureLongPressClassNewFromC(native unsafe.Pointer) *GestureLongPressClass {
	return &GestureLongPressClass{native: native}
}

// GestureMultiPressClass is a representation of the C record GtkGestureMultiPressClass.
type GestureMultiPressClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkGestureMultiPressClass that represents the GestureMultiPressClass.
func (recv *GestureMultiPressClass) ToC() unsafe.Pointer {
	return recv.native
}

// GestureMultiPressClassNewFromC creates a new GestureMultiPressClass from a pointer to the C GtkGestureMultiPressClass that represents the GestureMultiPressClass.
func GestureMultiPressClassNewFromC(native unsafe.Pointer) *GestureMultiPressClass {
	return &GestureMultiPressClass{native: native}
}

// GesturePanClass is a representation of the C record GtkGesturePanClass.
type GesturePanClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkGesturePanClass that represents the GesturePanClass.
func (recv *GesturePanClass) ToC() unsafe.Pointer {
	return recv.native
}

// GesturePanClassNewFromC creates a new GesturePanClass from a pointer to the C GtkGesturePanClass that represents the GesturePanClass.
func GesturePanClassNewFromC(native unsafe.Pointer) *GesturePanClass {
	return &GesturePanClass{native: native}
}

// GestureRotateClass is a representation of the C record GtkGestureRotateClass.
type GestureRotateClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkGestureRotateClass that represents the GestureRotateClass.
func (recv *GestureRotateClass) ToC() unsafe.Pointer {
	return recv.native
}

// GestureRotateClassNewFromC creates a new GestureRotateClass from a pointer to the C GtkGestureRotateClass that represents the GestureRotateClass.
func GestureRotateClassNewFromC(native unsafe.Pointer) *GestureRotateClass {
	return &GestureRotateClass{native: native}
}

// GestureSingleClass is a representation of the C record GtkGestureSingleClass.
type GestureSingleClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkGestureSingleClass that represents the GestureSingleClass.
func (recv *GestureSingleClass) ToC() unsafe.Pointer {
	return recv.native
}

// GestureSingleClassNewFromC creates a new GestureSingleClass from a pointer to the C GtkGestureSingleClass that represents the GestureSingleClass.
func GestureSingleClassNewFromC(native unsafe.Pointer) *GestureSingleClass {
	return &GestureSingleClass{native: native}
}

// UNSUPPORTED : GestureStylusClass : blacklisted

// GestureSwipeClass is a representation of the C record GtkGestureSwipeClass.
type GestureSwipeClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkGestureSwipeClass that represents the GestureSwipeClass.
func (recv *GestureSwipeClass) ToC() unsafe.Pointer {
	return recv.native
}

// GestureSwipeClassNewFromC creates a new GestureSwipeClass from a pointer to the C GtkGestureSwipeClass that represents the GestureSwipeClass.
func GestureSwipeClassNewFromC(native unsafe.Pointer) *GestureSwipeClass {
	return &GestureSwipeClass{native: native}
}

// GestureZoomClass is a representation of the C record GtkGestureZoomClass.
type GestureZoomClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkGestureZoomClass that represents the GestureZoomClass.
func (recv *GestureZoomClass) ToC() unsafe.Pointer {
	return recv.native
}

// GestureZoomClassNewFromC creates a new GestureZoomClass from a pointer to the C GtkGestureZoomClass that represents the GestureZoomClass.
func GestureZoomClassNewFromC(native unsafe.Pointer) *GestureZoomClass {
	return &GestureZoomClass{native: native}
}

// Gradient is a representation of the C record GtkGradient.
type Gradient struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkGradient that represents the Gradient.
func (recv *Gradient) ToC() unsafe.Pointer {
	return recv.native
}

// GradientNewFromC creates a new Gradient from a pointer to the C GtkGradient that represents the Gradient.
func GradientNewFromC(native unsafe.Pointer) *Gradient {
	return &Gradient{native: native}
}

// GridClass is a representation of the C record GtkGridClass.
type GridClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkGridClass that represents the GridClass.
func (recv *GridClass) ToC() unsafe.Pointer {
	return recv.native
}

// GridClassNewFromC creates a new GridClass from a pointer to the C GtkGridClass that represents the GridClass.
func GridClassNewFromC(native unsafe.Pointer) *GridClass {
	return &GridClass{native: native}
}

// GridPrivate is a representation of the C record GtkGridPrivate.
type GridPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkGridPrivate that represents the GridPrivate.
func (recv *GridPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// GridPrivateNewFromC creates a new GridPrivate from a pointer to the C GtkGridPrivate that represents the GridPrivate.
func GridPrivateNewFromC(native unsafe.Pointer) *GridPrivate {
	return &GridPrivate{native: native}
}

// HBoxClass is a representation of the C record GtkHBoxClass.
type HBoxClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkHBoxClass that represents the HBoxClass.
func (recv *HBoxClass) ToC() unsafe.Pointer {
	return recv.native
}

// HBoxClassNewFromC creates a new HBoxClass from a pointer to the C GtkHBoxClass that represents the HBoxClass.
func HBoxClassNewFromC(native unsafe.Pointer) *HBoxClass {
	return &HBoxClass{native: native}
}

// HButtonBoxClass is a representation of the C record GtkHButtonBoxClass.
type HButtonBoxClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkHButtonBoxClass that represents the HButtonBoxClass.
func (recv *HButtonBoxClass) ToC() unsafe.Pointer {
	return recv.native
}

// HButtonBoxClassNewFromC creates a new HButtonBoxClass from a pointer to the C GtkHButtonBoxClass that represents the HButtonBoxClass.
func HButtonBoxClassNewFromC(native unsafe.Pointer) *HButtonBoxClass {
	return &HButtonBoxClass{native: native}
}

// HPanedClass is a representation of the C record GtkHPanedClass.
type HPanedClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkHPanedClass that represents the HPanedClass.
func (recv *HPanedClass) ToC() unsafe.Pointer {
	return recv.native
}

// HPanedClassNewFromC creates a new HPanedClass from a pointer to the C GtkHPanedClass that represents the HPanedClass.
func HPanedClassNewFromC(native unsafe.Pointer) *HPanedClass {
	return &HPanedClass{native: native}
}

// HSVClass is a representation of the C record GtkHSVClass.
type HSVClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkHSVClass that represents the HSVClass.
func (recv *HSVClass) ToC() unsafe.Pointer {
	return recv.native
}

// HSVClassNewFromC creates a new HSVClass from a pointer to the C GtkHSVClass that represents the HSVClass.
func HSVClassNewFromC(native unsafe.Pointer) *HSVClass {
	return &HSVClass{native: native}
}

// HSVPrivate is a representation of the C record GtkHSVPrivate.
type HSVPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkHSVPrivate that represents the HSVPrivate.
func (recv *HSVPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// HSVPrivateNewFromC creates a new HSVPrivate from a pointer to the C GtkHSVPrivate that represents the HSVPrivate.
func HSVPrivateNewFromC(native unsafe.Pointer) *HSVPrivate {
	return &HSVPrivate{native: native}
}

// HScaleClass is a representation of the C record GtkHScaleClass.
type HScaleClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkHScaleClass that represents the HScaleClass.
func (recv *HScaleClass) ToC() unsafe.Pointer {
	return recv.native
}

// HScaleClassNewFromC creates a new HScaleClass from a pointer to the C GtkHScaleClass that represents the HScaleClass.
func HScaleClassNewFromC(native unsafe.Pointer) *HScaleClass {
	return &HScaleClass{native: native}
}

// HScrollbarClass is a representation of the C record GtkHScrollbarClass.
type HScrollbarClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkHScrollbarClass that represents the HScrollbarClass.
func (recv *HScrollbarClass) ToC() unsafe.Pointer {
	return recv.native
}

// HScrollbarClassNewFromC creates a new HScrollbarClass from a pointer to the C GtkHScrollbarClass that represents the HScrollbarClass.
func HScrollbarClassNewFromC(native unsafe.Pointer) *HScrollbarClass {
	return &HScrollbarClass{native: native}
}

// HSeparatorClass is a representation of the C record GtkHSeparatorClass.
type HSeparatorClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkHSeparatorClass that represents the HSeparatorClass.
func (recv *HSeparatorClass) ToC() unsafe.Pointer {
	return recv.native
}

// HSeparatorClassNewFromC creates a new HSeparatorClass from a pointer to the C GtkHSeparatorClass that represents the HSeparatorClass.
func HSeparatorClassNewFromC(native unsafe.Pointer) *HSeparatorClass {
	return &HSeparatorClass{native: native}
}

// HandleBoxClass is a representation of the C record GtkHandleBoxClass.
type HandleBoxClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkHandleBoxClass that represents the HandleBoxClass.
func (recv *HandleBoxClass) ToC() unsafe.Pointer {
	return recv.native
}

// HandleBoxClassNewFromC creates a new HandleBoxClass from a pointer to the C GtkHandleBoxClass that represents the HandleBoxClass.
func HandleBoxClassNewFromC(native unsafe.Pointer) *HandleBoxClass {
	return &HandleBoxClass{native: native}
}

// HandleBoxPrivate is a representation of the C record GtkHandleBoxPrivate.
type HandleBoxPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkHandleBoxPrivate that represents the HandleBoxPrivate.
func (recv *HandleBoxPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// HandleBoxPrivateNewFromC creates a new HandleBoxPrivate from a pointer to the C GtkHandleBoxPrivate that represents the HandleBoxPrivate.
func HandleBoxPrivateNewFromC(native unsafe.Pointer) *HandleBoxPrivate {
	return &HandleBoxPrivate{native: native}
}

// UNSUPPORTED : HeaderBarAccessibleClass : blacklisted

// UNSUPPORTED : HeaderBarAccessiblePrivate : blacklisted

// HeaderBarClass is a representation of the C record GtkHeaderBarClass.
type HeaderBarClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkHeaderBarClass that represents the HeaderBarClass.
func (recv *HeaderBarClass) ToC() unsafe.Pointer {
	return recv.native
}

// HeaderBarClassNewFromC creates a new HeaderBarClass from a pointer to the C GtkHeaderBarClass that represents the HeaderBarClass.
func HeaderBarClassNewFromC(native unsafe.Pointer) *HeaderBarClass {
	return &HeaderBarClass{native: native}
}

// HeaderBarPrivate is a representation of the C record GtkHeaderBarPrivate.
type HeaderBarPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkHeaderBarPrivate that represents the HeaderBarPrivate.
func (recv *HeaderBarPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// HeaderBarPrivateNewFromC creates a new HeaderBarPrivate from a pointer to the C GtkHeaderBarPrivate that represents the HeaderBarPrivate.
func HeaderBarPrivateNewFromC(native unsafe.Pointer) *HeaderBarPrivate {
	return &HeaderBarPrivate{native: native}
}

// IMContextClass is a representation of the C record GtkIMContextClass.
type IMContextClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkIMContextClass that represents the IMContextClass.
func (recv *IMContextClass) ToC() unsafe.Pointer {
	return recv.native
}

// IMContextClassNewFromC creates a new IMContextClass from a pointer to the C GtkIMContextClass that represents the IMContextClass.
func IMContextClassNewFromC(native unsafe.Pointer) *IMContextClass {
	return &IMContextClass{native: native}
}

// IMContextInfo is a representation of the C record GtkIMContextInfo.
type IMContextInfo struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkIMContextInfo that represents the IMContextInfo.
func (recv *IMContextInfo) ToC() unsafe.Pointer {
	return recv.native
}

// IMContextInfoNewFromC creates a new IMContextInfo from a pointer to the C GtkIMContextInfo that represents the IMContextInfo.
func IMContextInfoNewFromC(native unsafe.Pointer) *IMContextInfo {
	return &IMContextInfo{native: native}
}

// IMContextSimpleClass is a representation of the C record GtkIMContextSimpleClass.
type IMContextSimpleClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkIMContextSimpleClass that represents the IMContextSimpleClass.
func (recv *IMContextSimpleClass) ToC() unsafe.Pointer {
	return recv.native
}

// IMContextSimpleClassNewFromC creates a new IMContextSimpleClass from a pointer to the C GtkIMContextSimpleClass that represents the IMContextSimpleClass.
func IMContextSimpleClassNewFromC(native unsafe.Pointer) *IMContextSimpleClass {
	return &IMContextSimpleClass{native: native}
}

// IMContextSimplePrivate is a representation of the C record GtkIMContextSimplePrivate.
type IMContextSimplePrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkIMContextSimplePrivate that represents the IMContextSimplePrivate.
func (recv *IMContextSimplePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// IMContextSimplePrivateNewFromC creates a new IMContextSimplePrivate from a pointer to the C GtkIMContextSimplePrivate that represents the IMContextSimplePrivate.
func IMContextSimplePrivateNewFromC(native unsafe.Pointer) *IMContextSimplePrivate {
	return &IMContextSimplePrivate{native: native}
}

// IMMulticontextClass is a representation of the C record GtkIMMulticontextClass.
type IMMulticontextClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkIMMulticontextClass that represents the IMMulticontextClass.
func (recv *IMMulticontextClass) ToC() unsafe.Pointer {
	return recv.native
}

// IMMulticontextClassNewFromC creates a new IMMulticontextClass from a pointer to the C GtkIMMulticontextClass that represents the IMMulticontextClass.
func IMMulticontextClassNewFromC(native unsafe.Pointer) *IMMulticontextClass {
	return &IMMulticontextClass{native: native}
}

// IMMulticontextPrivate is a representation of the C record GtkIMMulticontextPrivate.
type IMMulticontextPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkIMMulticontextPrivate that represents the IMMulticontextPrivate.
func (recv *IMMulticontextPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// IMMulticontextPrivateNewFromC creates a new IMMulticontextPrivate from a pointer to the C GtkIMMulticontextPrivate that represents the IMMulticontextPrivate.
func IMMulticontextPrivateNewFromC(native unsafe.Pointer) *IMMulticontextPrivate {
	return &IMMulticontextPrivate{native: native}
}

// IconFactoryClass is a representation of the C record GtkIconFactoryClass.
type IconFactoryClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkIconFactoryClass that represents the IconFactoryClass.
func (recv *IconFactoryClass) ToC() unsafe.Pointer {
	return recv.native
}

// IconFactoryClassNewFromC creates a new IconFactoryClass from a pointer to the C GtkIconFactoryClass that represents the IconFactoryClass.
func IconFactoryClassNewFromC(native unsafe.Pointer) *IconFactoryClass {
	return &IconFactoryClass{native: native}
}

// IconFactoryPrivate is a representation of the C record GtkIconFactoryPrivate.
type IconFactoryPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkIconFactoryPrivate that represents the IconFactoryPrivate.
func (recv *IconFactoryPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// IconFactoryPrivateNewFromC creates a new IconFactoryPrivate from a pointer to the C GtkIconFactoryPrivate that represents the IconFactoryPrivate.
func IconFactoryPrivateNewFromC(native unsafe.Pointer) *IconFactoryPrivate {
	return &IconFactoryPrivate{native: native}
}

// IconInfoClass is a representation of the C record GtkIconInfoClass.
type IconInfoClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkIconInfoClass that represents the IconInfoClass.
func (recv *IconInfoClass) ToC() unsafe.Pointer {
	return recv.native
}

// IconInfoClassNewFromC creates a new IconInfoClass from a pointer to the C GtkIconInfoClass that represents the IconInfoClass.
func IconInfoClassNewFromC(native unsafe.Pointer) *IconInfoClass {
	return &IconInfoClass{native: native}
}

// IconSet is a representation of the C record GtkIconSet.
type IconSet struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkIconSet that represents the IconSet.
func (recv *IconSet) ToC() unsafe.Pointer {
	return recv.native
}

// IconSetNewFromC creates a new IconSet from a pointer to the C GtkIconSet that represents the IconSet.
func IconSetNewFromC(native unsafe.Pointer) *IconSet {
	return &IconSet{native: native}
}

// IconSource is a representation of the C record GtkIconSource.
type IconSource struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkIconSource that represents the IconSource.
func (recv *IconSource) ToC() unsafe.Pointer {
	return recv.native
}

// IconSourceNewFromC creates a new IconSource from a pointer to the C GtkIconSource that represents the IconSource.
func IconSourceNewFromC(native unsafe.Pointer) *IconSource {
	return &IconSource{native: native}
}

// IconThemeClass is a representation of the C record GtkIconThemeClass.
type IconThemeClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkIconThemeClass that represents the IconThemeClass.
func (recv *IconThemeClass) ToC() unsafe.Pointer {
	return recv.native
}

// IconThemeClassNewFromC creates a new IconThemeClass from a pointer to the C GtkIconThemeClass that represents the IconThemeClass.
func IconThemeClassNewFromC(native unsafe.Pointer) *IconThemeClass {
	return &IconThemeClass{native: native}
}

// IconThemePrivate is a representation of the C record GtkIconThemePrivate.
type IconThemePrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkIconThemePrivate that represents the IconThemePrivate.
func (recv *IconThemePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// IconThemePrivateNewFromC creates a new IconThemePrivate from a pointer to the C GtkIconThemePrivate that represents the IconThemePrivate.
func IconThemePrivateNewFromC(native unsafe.Pointer) *IconThemePrivate {
	return &IconThemePrivate{native: native}
}

// IconViewAccessibleClass is a representation of the C record GtkIconViewAccessibleClass.
type IconViewAccessibleClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkIconViewAccessibleClass that represents the IconViewAccessibleClass.
func (recv *IconViewAccessibleClass) ToC() unsafe.Pointer {
	return recv.native
}

// IconViewAccessibleClassNewFromC creates a new IconViewAccessibleClass from a pointer to the C GtkIconViewAccessibleClass that represents the IconViewAccessibleClass.
func IconViewAccessibleClassNewFromC(native unsafe.Pointer) *IconViewAccessibleClass {
	return &IconViewAccessibleClass{native: native}
}

// IconViewAccessiblePrivate is a representation of the C record GtkIconViewAccessiblePrivate.
type IconViewAccessiblePrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkIconViewAccessiblePrivate that represents the IconViewAccessiblePrivate.
func (recv *IconViewAccessiblePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// IconViewAccessiblePrivateNewFromC creates a new IconViewAccessiblePrivate from a pointer to the C GtkIconViewAccessiblePrivate that represents the IconViewAccessiblePrivate.
func IconViewAccessiblePrivateNewFromC(native unsafe.Pointer) *IconViewAccessiblePrivate {
	return &IconViewAccessiblePrivate{native: native}
}

// IconViewClass is a representation of the C record GtkIconViewClass.
type IconViewClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkIconViewClass that represents the IconViewClass.
func (recv *IconViewClass) ToC() unsafe.Pointer {
	return recv.native
}

// IconViewClassNewFromC creates a new IconViewClass from a pointer to the C GtkIconViewClass that represents the IconViewClass.
func IconViewClassNewFromC(native unsafe.Pointer) *IconViewClass {
	return &IconViewClass{native: native}
}

// IconViewPrivate is a representation of the C record GtkIconViewPrivate.
type IconViewPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkIconViewPrivate that represents the IconViewPrivate.
func (recv *IconViewPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// IconViewPrivateNewFromC creates a new IconViewPrivate from a pointer to the C GtkIconViewPrivate that represents the IconViewPrivate.
func IconViewPrivateNewFromC(native unsafe.Pointer) *IconViewPrivate {
	return &IconViewPrivate{native: native}
}

// ImageAccessibleClass is a representation of the C record GtkImageAccessibleClass.
type ImageAccessibleClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkImageAccessibleClass that represents the ImageAccessibleClass.
func (recv *ImageAccessibleClass) ToC() unsafe.Pointer {
	return recv.native
}

// ImageAccessibleClassNewFromC creates a new ImageAccessibleClass from a pointer to the C GtkImageAccessibleClass that represents the ImageAccessibleClass.
func ImageAccessibleClassNewFromC(native unsafe.Pointer) *ImageAccessibleClass {
	return &ImageAccessibleClass{native: native}
}

// ImageAccessiblePrivate is a representation of the C record GtkImageAccessiblePrivate.
type ImageAccessiblePrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkImageAccessiblePrivate that represents the ImageAccessiblePrivate.
func (recv *ImageAccessiblePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ImageAccessiblePrivateNewFromC creates a new ImageAccessiblePrivate from a pointer to the C GtkImageAccessiblePrivate that represents the ImageAccessiblePrivate.
func ImageAccessiblePrivateNewFromC(native unsafe.Pointer) *ImageAccessiblePrivate {
	return &ImageAccessiblePrivate{native: native}
}

// ImageCellAccessibleClass is a representation of the C record GtkImageCellAccessibleClass.
type ImageCellAccessibleClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkImageCellAccessibleClass that represents the ImageCellAccessibleClass.
func (recv *ImageCellAccessibleClass) ToC() unsafe.Pointer {
	return recv.native
}

// ImageCellAccessibleClassNewFromC creates a new ImageCellAccessibleClass from a pointer to the C GtkImageCellAccessibleClass that represents the ImageCellAccessibleClass.
func ImageCellAccessibleClassNewFromC(native unsafe.Pointer) *ImageCellAccessibleClass {
	return &ImageCellAccessibleClass{native: native}
}

// ImageCellAccessiblePrivate is a representation of the C record GtkImageCellAccessiblePrivate.
type ImageCellAccessiblePrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkImageCellAccessiblePrivate that represents the ImageCellAccessiblePrivate.
func (recv *ImageCellAccessiblePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ImageCellAccessiblePrivateNewFromC creates a new ImageCellAccessiblePrivate from a pointer to the C GtkImageCellAccessiblePrivate that represents the ImageCellAccessiblePrivate.
func ImageCellAccessiblePrivateNewFromC(native unsafe.Pointer) *ImageCellAccessiblePrivate {
	return &ImageCellAccessiblePrivate{native: native}
}

// ImageClass is a representation of the C record GtkImageClass.
type ImageClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkImageClass that represents the ImageClass.
func (recv *ImageClass) ToC() unsafe.Pointer {
	return recv.native
}

// ImageClassNewFromC creates a new ImageClass from a pointer to the C GtkImageClass that represents the ImageClass.
func ImageClassNewFromC(native unsafe.Pointer) *ImageClass {
	return &ImageClass{native: native}
}

// ImageMenuItemClass is a representation of the C record GtkImageMenuItemClass.
type ImageMenuItemClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkImageMenuItemClass that represents the ImageMenuItemClass.
func (recv *ImageMenuItemClass) ToC() unsafe.Pointer {
	return recv.native
}

// ImageMenuItemClassNewFromC creates a new ImageMenuItemClass from a pointer to the C GtkImageMenuItemClass that represents the ImageMenuItemClass.
func ImageMenuItemClassNewFromC(native unsafe.Pointer) *ImageMenuItemClass {
	return &ImageMenuItemClass{native: native}
}

// ImageMenuItemPrivate is a representation of the C record GtkImageMenuItemPrivate.
type ImageMenuItemPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkImageMenuItemPrivate that represents the ImageMenuItemPrivate.
func (recv *ImageMenuItemPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ImageMenuItemPrivateNewFromC creates a new ImageMenuItemPrivate from a pointer to the C GtkImageMenuItemPrivate that represents the ImageMenuItemPrivate.
func ImageMenuItemPrivateNewFromC(native unsafe.Pointer) *ImageMenuItemPrivate {
	return &ImageMenuItemPrivate{native: native}
}

// ImagePrivate is a representation of the C record GtkImagePrivate.
type ImagePrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkImagePrivate that represents the ImagePrivate.
func (recv *ImagePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ImagePrivateNewFromC creates a new ImagePrivate from a pointer to the C GtkImagePrivate that represents the ImagePrivate.
func ImagePrivateNewFromC(native unsafe.Pointer) *ImagePrivate {
	return &ImagePrivate{native: native}
}

// InfoBarClass is a representation of the C record GtkInfoBarClass.
type InfoBarClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkInfoBarClass that represents the InfoBarClass.
func (recv *InfoBarClass) ToC() unsafe.Pointer {
	return recv.native
}

// InfoBarClassNewFromC creates a new InfoBarClass from a pointer to the C GtkInfoBarClass that represents the InfoBarClass.
func InfoBarClassNewFromC(native unsafe.Pointer) *InfoBarClass {
	return &InfoBarClass{native: native}
}

// InfoBarPrivate is a representation of the C record GtkInfoBarPrivate.
type InfoBarPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkInfoBarPrivate that represents the InfoBarPrivate.
func (recv *InfoBarPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// InfoBarPrivateNewFromC creates a new InfoBarPrivate from a pointer to the C GtkInfoBarPrivate that represents the InfoBarPrivate.
func InfoBarPrivateNewFromC(native unsafe.Pointer) *InfoBarPrivate {
	return &InfoBarPrivate{native: native}
}

// InvisibleClass is a representation of the C record GtkInvisibleClass.
type InvisibleClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkInvisibleClass that represents the InvisibleClass.
func (recv *InvisibleClass) ToC() unsafe.Pointer {
	return recv.native
}

// InvisibleClassNewFromC creates a new InvisibleClass from a pointer to the C GtkInvisibleClass that represents the InvisibleClass.
func InvisibleClassNewFromC(native unsafe.Pointer) *InvisibleClass {
	return &InvisibleClass{native: native}
}

// InvisiblePrivate is a representation of the C record GtkInvisiblePrivate.
type InvisiblePrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkInvisiblePrivate that represents the InvisiblePrivate.
func (recv *InvisiblePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// InvisiblePrivateNewFromC creates a new InvisiblePrivate from a pointer to the C GtkInvisiblePrivate that represents the InvisiblePrivate.
func InvisiblePrivateNewFromC(native unsafe.Pointer) *InvisiblePrivate {
	return &InvisiblePrivate{native: native}
}

// LabelAccessibleClass is a representation of the C record GtkLabelAccessibleClass.
type LabelAccessibleClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkLabelAccessibleClass that represents the LabelAccessibleClass.
func (recv *LabelAccessibleClass) ToC() unsafe.Pointer {
	return recv.native
}

// LabelAccessibleClassNewFromC creates a new LabelAccessibleClass from a pointer to the C GtkLabelAccessibleClass that represents the LabelAccessibleClass.
func LabelAccessibleClassNewFromC(native unsafe.Pointer) *LabelAccessibleClass {
	return &LabelAccessibleClass{native: native}
}

// LabelAccessiblePrivate is a representation of the C record GtkLabelAccessiblePrivate.
type LabelAccessiblePrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkLabelAccessiblePrivate that represents the LabelAccessiblePrivate.
func (recv *LabelAccessiblePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// LabelAccessiblePrivateNewFromC creates a new LabelAccessiblePrivate from a pointer to the C GtkLabelAccessiblePrivate that represents the LabelAccessiblePrivate.
func LabelAccessiblePrivateNewFromC(native unsafe.Pointer) *LabelAccessiblePrivate {
	return &LabelAccessiblePrivate{native: native}
}

// LabelClass is a representation of the C record GtkLabelClass.
type LabelClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkLabelClass that represents the LabelClass.
func (recv *LabelClass) ToC() unsafe.Pointer {
	return recv.native
}

// LabelClassNewFromC creates a new LabelClass from a pointer to the C GtkLabelClass that represents the LabelClass.
func LabelClassNewFromC(native unsafe.Pointer) *LabelClass {
	return &LabelClass{native: native}
}

// LabelPrivate is a representation of the C record GtkLabelPrivate.
type LabelPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkLabelPrivate that represents the LabelPrivate.
func (recv *LabelPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// LabelPrivateNewFromC creates a new LabelPrivate from a pointer to the C GtkLabelPrivate that represents the LabelPrivate.
func LabelPrivateNewFromC(native unsafe.Pointer) *LabelPrivate {
	return &LabelPrivate{native: native}
}

// LabelSelectionInfo is a representation of the C record GtkLabelSelectionInfo.
type LabelSelectionInfo struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkLabelSelectionInfo that represents the LabelSelectionInfo.
func (recv *LabelSelectionInfo) ToC() unsafe.Pointer {
	return recv.native
}

// LabelSelectionInfoNewFromC creates a new LabelSelectionInfo from a pointer to the C GtkLabelSelectionInfo that represents the LabelSelectionInfo.
func LabelSelectionInfoNewFromC(native unsafe.Pointer) *LabelSelectionInfo {
	return &LabelSelectionInfo{native: native}
}

// LayoutClass is a representation of the C record GtkLayoutClass.
type LayoutClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkLayoutClass that represents the LayoutClass.
func (recv *LayoutClass) ToC() unsafe.Pointer {
	return recv.native
}

// LayoutClassNewFromC creates a new LayoutClass from a pointer to the C GtkLayoutClass that represents the LayoutClass.
func LayoutClassNewFromC(native unsafe.Pointer) *LayoutClass {
	return &LayoutClass{native: native}
}

// LayoutPrivate is a representation of the C record GtkLayoutPrivate.
type LayoutPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkLayoutPrivate that represents the LayoutPrivate.
func (recv *LayoutPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// LayoutPrivateNewFromC creates a new LayoutPrivate from a pointer to the C GtkLayoutPrivate that represents the LayoutPrivate.
func LayoutPrivateNewFromC(native unsafe.Pointer) *LayoutPrivate {
	return &LayoutPrivate{native: native}
}

// LevelBarAccessibleClass is a representation of the C record GtkLevelBarAccessibleClass.
type LevelBarAccessibleClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkLevelBarAccessibleClass that represents the LevelBarAccessibleClass.
func (recv *LevelBarAccessibleClass) ToC() unsafe.Pointer {
	return recv.native
}

// LevelBarAccessibleClassNewFromC creates a new LevelBarAccessibleClass from a pointer to the C GtkLevelBarAccessibleClass that represents the LevelBarAccessibleClass.
func LevelBarAccessibleClassNewFromC(native unsafe.Pointer) *LevelBarAccessibleClass {
	return &LevelBarAccessibleClass{native: native}
}

// LevelBarAccessiblePrivate is a representation of the C record GtkLevelBarAccessiblePrivate.
type LevelBarAccessiblePrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkLevelBarAccessiblePrivate that represents the LevelBarAccessiblePrivate.
func (recv *LevelBarAccessiblePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// LevelBarAccessiblePrivateNewFromC creates a new LevelBarAccessiblePrivate from a pointer to the C GtkLevelBarAccessiblePrivate that represents the LevelBarAccessiblePrivate.
func LevelBarAccessiblePrivateNewFromC(native unsafe.Pointer) *LevelBarAccessiblePrivate {
	return &LevelBarAccessiblePrivate{native: native}
}

// LevelBarClass is a representation of the C record GtkLevelBarClass.
type LevelBarClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkLevelBarClass that represents the LevelBarClass.
func (recv *LevelBarClass) ToC() unsafe.Pointer {
	return recv.native
}

// LevelBarClassNewFromC creates a new LevelBarClass from a pointer to the C GtkLevelBarClass that represents the LevelBarClass.
func LevelBarClassNewFromC(native unsafe.Pointer) *LevelBarClass {
	return &LevelBarClass{native: native}
}

// LevelBarPrivate is a representation of the C record GtkLevelBarPrivate.
type LevelBarPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkLevelBarPrivate that represents the LevelBarPrivate.
func (recv *LevelBarPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// LevelBarPrivateNewFromC creates a new LevelBarPrivate from a pointer to the C GtkLevelBarPrivate that represents the LevelBarPrivate.
func LevelBarPrivateNewFromC(native unsafe.Pointer) *LevelBarPrivate {
	return &LevelBarPrivate{native: native}
}

// LinkButtonAccessibleClass is a representation of the C record GtkLinkButtonAccessibleClass.
type LinkButtonAccessibleClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkLinkButtonAccessibleClass that represents the LinkButtonAccessibleClass.
func (recv *LinkButtonAccessibleClass) ToC() unsafe.Pointer {
	return recv.native
}

// LinkButtonAccessibleClassNewFromC creates a new LinkButtonAccessibleClass from a pointer to the C GtkLinkButtonAccessibleClass that represents the LinkButtonAccessibleClass.
func LinkButtonAccessibleClassNewFromC(native unsafe.Pointer) *LinkButtonAccessibleClass {
	return &LinkButtonAccessibleClass{native: native}
}

// LinkButtonAccessiblePrivate is a representation of the C record GtkLinkButtonAccessiblePrivate.
type LinkButtonAccessiblePrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkLinkButtonAccessiblePrivate that represents the LinkButtonAccessiblePrivate.
func (recv *LinkButtonAccessiblePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// LinkButtonAccessiblePrivateNewFromC creates a new LinkButtonAccessiblePrivate from a pointer to the C GtkLinkButtonAccessiblePrivate that represents the LinkButtonAccessiblePrivate.
func LinkButtonAccessiblePrivateNewFromC(native unsafe.Pointer) *LinkButtonAccessiblePrivate {
	return &LinkButtonAccessiblePrivate{native: native}
}

// LinkButtonClass is a representation of the C record GtkLinkButtonClass.
type LinkButtonClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkLinkButtonClass that represents the LinkButtonClass.
func (recv *LinkButtonClass) ToC() unsafe.Pointer {
	return recv.native
}

// LinkButtonClassNewFromC creates a new LinkButtonClass from a pointer to the C GtkLinkButtonClass that represents the LinkButtonClass.
func LinkButtonClassNewFromC(native unsafe.Pointer) *LinkButtonClass {
	return &LinkButtonClass{native: native}
}

// LinkButtonPrivate is a representation of the C record GtkLinkButtonPrivate.
type LinkButtonPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkLinkButtonPrivate that represents the LinkButtonPrivate.
func (recv *LinkButtonPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// LinkButtonPrivateNewFromC creates a new LinkButtonPrivate from a pointer to the C GtkLinkButtonPrivate that represents the LinkButtonPrivate.
func LinkButtonPrivateNewFromC(native unsafe.Pointer) *LinkButtonPrivate {
	return &LinkButtonPrivate{native: native}
}

// ListBoxAccessibleClass is a representation of the C record GtkListBoxAccessibleClass.
type ListBoxAccessibleClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkListBoxAccessibleClass that represents the ListBoxAccessibleClass.
func (recv *ListBoxAccessibleClass) ToC() unsafe.Pointer {
	return recv.native
}

// ListBoxAccessibleClassNewFromC creates a new ListBoxAccessibleClass from a pointer to the C GtkListBoxAccessibleClass that represents the ListBoxAccessibleClass.
func ListBoxAccessibleClassNewFromC(native unsafe.Pointer) *ListBoxAccessibleClass {
	return &ListBoxAccessibleClass{native: native}
}

// ListBoxAccessiblePrivate is a representation of the C record GtkListBoxAccessiblePrivate.
type ListBoxAccessiblePrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkListBoxAccessiblePrivate that represents the ListBoxAccessiblePrivate.
func (recv *ListBoxAccessiblePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ListBoxAccessiblePrivateNewFromC creates a new ListBoxAccessiblePrivate from a pointer to the C GtkListBoxAccessiblePrivate that represents the ListBoxAccessiblePrivate.
func ListBoxAccessiblePrivateNewFromC(native unsafe.Pointer) *ListBoxAccessiblePrivate {
	return &ListBoxAccessiblePrivate{native: native}
}

// ListBoxClass is a representation of the C record GtkListBoxClass.
type ListBoxClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkListBoxClass that represents the ListBoxClass.
func (recv *ListBoxClass) ToC() unsafe.Pointer {
	return recv.native
}

// ListBoxClassNewFromC creates a new ListBoxClass from a pointer to the C GtkListBoxClass that represents the ListBoxClass.
func ListBoxClassNewFromC(native unsafe.Pointer) *ListBoxClass {
	return &ListBoxClass{native: native}
}

// ListBoxRowAccessibleClass is a representation of the C record GtkListBoxRowAccessibleClass.
type ListBoxRowAccessibleClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkListBoxRowAccessibleClass that represents the ListBoxRowAccessibleClass.
func (recv *ListBoxRowAccessibleClass) ToC() unsafe.Pointer {
	return recv.native
}

// ListBoxRowAccessibleClassNewFromC creates a new ListBoxRowAccessibleClass from a pointer to the C GtkListBoxRowAccessibleClass that represents the ListBoxRowAccessibleClass.
func ListBoxRowAccessibleClassNewFromC(native unsafe.Pointer) *ListBoxRowAccessibleClass {
	return &ListBoxRowAccessibleClass{native: native}
}

// ListBoxRowClass is a representation of the C record GtkListBoxRowClass.
type ListBoxRowClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkListBoxRowClass that represents the ListBoxRowClass.
func (recv *ListBoxRowClass) ToC() unsafe.Pointer {
	return recv.native
}

// ListBoxRowClassNewFromC creates a new ListBoxRowClass from a pointer to the C GtkListBoxRowClass that represents the ListBoxRowClass.
func ListBoxRowClassNewFromC(native unsafe.Pointer) *ListBoxRowClass {
	return &ListBoxRowClass{native: native}
}

// ListStoreClass is a representation of the C record GtkListStoreClass.
type ListStoreClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkListStoreClass that represents the ListStoreClass.
func (recv *ListStoreClass) ToC() unsafe.Pointer {
	return recv.native
}

// ListStoreClassNewFromC creates a new ListStoreClass from a pointer to the C GtkListStoreClass that represents the ListStoreClass.
func ListStoreClassNewFromC(native unsafe.Pointer) *ListStoreClass {
	return &ListStoreClass{native: native}
}

// ListStorePrivate is a representation of the C record GtkListStorePrivate.
type ListStorePrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkListStorePrivate that represents the ListStorePrivate.
func (recv *ListStorePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ListStorePrivateNewFromC creates a new ListStorePrivate from a pointer to the C GtkListStorePrivate that represents the ListStorePrivate.
func ListStorePrivateNewFromC(native unsafe.Pointer) *ListStorePrivate {
	return &ListStorePrivate{native: native}
}

// LockButtonAccessibleClass is a representation of the C record GtkLockButtonAccessibleClass.
type LockButtonAccessibleClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkLockButtonAccessibleClass that represents the LockButtonAccessibleClass.
func (recv *LockButtonAccessibleClass) ToC() unsafe.Pointer {
	return recv.native
}

// LockButtonAccessibleClassNewFromC creates a new LockButtonAccessibleClass from a pointer to the C GtkLockButtonAccessibleClass that represents the LockButtonAccessibleClass.
func LockButtonAccessibleClassNewFromC(native unsafe.Pointer) *LockButtonAccessibleClass {
	return &LockButtonAccessibleClass{native: native}
}

// LockButtonAccessiblePrivate is a representation of the C record GtkLockButtonAccessiblePrivate.
type LockButtonAccessiblePrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkLockButtonAccessiblePrivate that represents the LockButtonAccessiblePrivate.
func (recv *LockButtonAccessiblePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// LockButtonAccessiblePrivateNewFromC creates a new LockButtonAccessiblePrivate from a pointer to the C GtkLockButtonAccessiblePrivate that represents the LockButtonAccessiblePrivate.
func LockButtonAccessiblePrivateNewFromC(native unsafe.Pointer) *LockButtonAccessiblePrivate {
	return &LockButtonAccessiblePrivate{native: native}
}

// LockButtonClass is a representation of the C record GtkLockButtonClass.
type LockButtonClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkLockButtonClass that represents the LockButtonClass.
func (recv *LockButtonClass) ToC() unsafe.Pointer {
	return recv.native
}

// LockButtonClassNewFromC creates a new LockButtonClass from a pointer to the C GtkLockButtonClass that represents the LockButtonClass.
func LockButtonClassNewFromC(native unsafe.Pointer) *LockButtonClass {
	return &LockButtonClass{native: native}
}

// LockButtonPrivate is a representation of the C record GtkLockButtonPrivate.
type LockButtonPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkLockButtonPrivate that represents the LockButtonPrivate.
func (recv *LockButtonPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// LockButtonPrivateNewFromC creates a new LockButtonPrivate from a pointer to the C GtkLockButtonPrivate that represents the LockButtonPrivate.
func LockButtonPrivateNewFromC(native unsafe.Pointer) *LockButtonPrivate {
	return &LockButtonPrivate{native: native}
}

// MenuAccessibleClass is a representation of the C record GtkMenuAccessibleClass.
type MenuAccessibleClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkMenuAccessibleClass that represents the MenuAccessibleClass.
func (recv *MenuAccessibleClass) ToC() unsafe.Pointer {
	return recv.native
}

// MenuAccessibleClassNewFromC creates a new MenuAccessibleClass from a pointer to the C GtkMenuAccessibleClass that represents the MenuAccessibleClass.
func MenuAccessibleClassNewFromC(native unsafe.Pointer) *MenuAccessibleClass {
	return &MenuAccessibleClass{native: native}
}

// MenuAccessiblePrivate is a representation of the C record GtkMenuAccessiblePrivate.
type MenuAccessiblePrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkMenuAccessiblePrivate that represents the MenuAccessiblePrivate.
func (recv *MenuAccessiblePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// MenuAccessiblePrivateNewFromC creates a new MenuAccessiblePrivate from a pointer to the C GtkMenuAccessiblePrivate that represents the MenuAccessiblePrivate.
func MenuAccessiblePrivateNewFromC(native unsafe.Pointer) *MenuAccessiblePrivate {
	return &MenuAccessiblePrivate{native: native}
}

// MenuBarClass is a representation of the C record GtkMenuBarClass.
type MenuBarClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkMenuBarClass that represents the MenuBarClass.
func (recv *MenuBarClass) ToC() unsafe.Pointer {
	return recv.native
}

// MenuBarClassNewFromC creates a new MenuBarClass from a pointer to the C GtkMenuBarClass that represents the MenuBarClass.
func MenuBarClassNewFromC(native unsafe.Pointer) *MenuBarClass {
	return &MenuBarClass{native: native}
}

// MenuBarPrivate is a representation of the C record GtkMenuBarPrivate.
type MenuBarPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkMenuBarPrivate that represents the MenuBarPrivate.
func (recv *MenuBarPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// MenuBarPrivateNewFromC creates a new MenuBarPrivate from a pointer to the C GtkMenuBarPrivate that represents the MenuBarPrivate.
func MenuBarPrivateNewFromC(native unsafe.Pointer) *MenuBarPrivate {
	return &MenuBarPrivate{native: native}
}

// MenuButtonAccessibleClass is a representation of the C record GtkMenuButtonAccessibleClass.
type MenuButtonAccessibleClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkMenuButtonAccessibleClass that represents the MenuButtonAccessibleClass.
func (recv *MenuButtonAccessibleClass) ToC() unsafe.Pointer {
	return recv.native
}

// MenuButtonAccessibleClassNewFromC creates a new MenuButtonAccessibleClass from a pointer to the C GtkMenuButtonAccessibleClass that represents the MenuButtonAccessibleClass.
func MenuButtonAccessibleClassNewFromC(native unsafe.Pointer) *MenuButtonAccessibleClass {
	return &MenuButtonAccessibleClass{native: native}
}

// MenuButtonAccessiblePrivate is a representation of the C record GtkMenuButtonAccessiblePrivate.
type MenuButtonAccessiblePrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkMenuButtonAccessiblePrivate that represents the MenuButtonAccessiblePrivate.
func (recv *MenuButtonAccessiblePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// MenuButtonAccessiblePrivateNewFromC creates a new MenuButtonAccessiblePrivate from a pointer to the C GtkMenuButtonAccessiblePrivate that represents the MenuButtonAccessiblePrivate.
func MenuButtonAccessiblePrivateNewFromC(native unsafe.Pointer) *MenuButtonAccessiblePrivate {
	return &MenuButtonAccessiblePrivate{native: native}
}

// MenuButtonClass is a representation of the C record GtkMenuButtonClass.
type MenuButtonClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkMenuButtonClass that represents the MenuButtonClass.
func (recv *MenuButtonClass) ToC() unsafe.Pointer {
	return recv.native
}

// MenuButtonClassNewFromC creates a new MenuButtonClass from a pointer to the C GtkMenuButtonClass that represents the MenuButtonClass.
func MenuButtonClassNewFromC(native unsafe.Pointer) *MenuButtonClass {
	return &MenuButtonClass{native: native}
}

// MenuButtonPrivate is a representation of the C record GtkMenuButtonPrivate.
type MenuButtonPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkMenuButtonPrivate that represents the MenuButtonPrivate.
func (recv *MenuButtonPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// MenuButtonPrivateNewFromC creates a new MenuButtonPrivate from a pointer to the C GtkMenuButtonPrivate that represents the MenuButtonPrivate.
func MenuButtonPrivateNewFromC(native unsafe.Pointer) *MenuButtonPrivate {
	return &MenuButtonPrivate{native: native}
}

// MenuClass is a representation of the C record GtkMenuClass.
type MenuClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkMenuClass that represents the MenuClass.
func (recv *MenuClass) ToC() unsafe.Pointer {
	return recv.native
}

// MenuClassNewFromC creates a new MenuClass from a pointer to the C GtkMenuClass that represents the MenuClass.
func MenuClassNewFromC(native unsafe.Pointer) *MenuClass {
	return &MenuClass{native: native}
}

// MenuItemAccessibleClass is a representation of the C record GtkMenuItemAccessibleClass.
type MenuItemAccessibleClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkMenuItemAccessibleClass that represents the MenuItemAccessibleClass.
func (recv *MenuItemAccessibleClass) ToC() unsafe.Pointer {
	return recv.native
}

// MenuItemAccessibleClassNewFromC creates a new MenuItemAccessibleClass from a pointer to the C GtkMenuItemAccessibleClass that represents the MenuItemAccessibleClass.
func MenuItemAccessibleClassNewFromC(native unsafe.Pointer) *MenuItemAccessibleClass {
	return &MenuItemAccessibleClass{native: native}
}

// MenuItemAccessiblePrivate is a representation of the C record GtkMenuItemAccessiblePrivate.
type MenuItemAccessiblePrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkMenuItemAccessiblePrivate that represents the MenuItemAccessiblePrivate.
func (recv *MenuItemAccessiblePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// MenuItemAccessiblePrivateNewFromC creates a new MenuItemAccessiblePrivate from a pointer to the C GtkMenuItemAccessiblePrivate that represents the MenuItemAccessiblePrivate.
func MenuItemAccessiblePrivateNewFromC(native unsafe.Pointer) *MenuItemAccessiblePrivate {
	return &MenuItemAccessiblePrivate{native: native}
}

// MenuItemClass is a representation of the C record GtkMenuItemClass.
type MenuItemClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkMenuItemClass that represents the MenuItemClass.
func (recv *MenuItemClass) ToC() unsafe.Pointer {
	return recv.native
}

// MenuItemClassNewFromC creates a new MenuItemClass from a pointer to the C GtkMenuItemClass that represents the MenuItemClass.
func MenuItemClassNewFromC(native unsafe.Pointer) *MenuItemClass {
	return &MenuItemClass{native: native}
}

// MenuItemPrivate is a representation of the C record GtkMenuItemPrivate.
type MenuItemPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkMenuItemPrivate that represents the MenuItemPrivate.
func (recv *MenuItemPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// MenuItemPrivateNewFromC creates a new MenuItemPrivate from a pointer to the C GtkMenuItemPrivate that represents the MenuItemPrivate.
func MenuItemPrivateNewFromC(native unsafe.Pointer) *MenuItemPrivate {
	return &MenuItemPrivate{native: native}
}

// MenuPrivate is a representation of the C record GtkMenuPrivate.
type MenuPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkMenuPrivate that represents the MenuPrivate.
func (recv *MenuPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// MenuPrivateNewFromC creates a new MenuPrivate from a pointer to the C GtkMenuPrivate that represents the MenuPrivate.
func MenuPrivateNewFromC(native unsafe.Pointer) *MenuPrivate {
	return &MenuPrivate{native: native}
}

// MenuShellAccessibleClass is a representation of the C record GtkMenuShellAccessibleClass.
type MenuShellAccessibleClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkMenuShellAccessibleClass that represents the MenuShellAccessibleClass.
func (recv *MenuShellAccessibleClass) ToC() unsafe.Pointer {
	return recv.native
}

// MenuShellAccessibleClassNewFromC creates a new MenuShellAccessibleClass from a pointer to the C GtkMenuShellAccessibleClass that represents the MenuShellAccessibleClass.
func MenuShellAccessibleClassNewFromC(native unsafe.Pointer) *MenuShellAccessibleClass {
	return &MenuShellAccessibleClass{native: native}
}

// MenuShellAccessiblePrivate is a representation of the C record GtkMenuShellAccessiblePrivate.
type MenuShellAccessiblePrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkMenuShellAccessiblePrivate that represents the MenuShellAccessiblePrivate.
func (recv *MenuShellAccessiblePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// MenuShellAccessiblePrivateNewFromC creates a new MenuShellAccessiblePrivate from a pointer to the C GtkMenuShellAccessiblePrivate that represents the MenuShellAccessiblePrivate.
func MenuShellAccessiblePrivateNewFromC(native unsafe.Pointer) *MenuShellAccessiblePrivate {
	return &MenuShellAccessiblePrivate{native: native}
}

// MenuShellClass is a representation of the C record GtkMenuShellClass.
type MenuShellClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkMenuShellClass that represents the MenuShellClass.
func (recv *MenuShellClass) ToC() unsafe.Pointer {
	return recv.native
}

// MenuShellClassNewFromC creates a new MenuShellClass from a pointer to the C GtkMenuShellClass that represents the MenuShellClass.
func MenuShellClassNewFromC(native unsafe.Pointer) *MenuShellClass {
	return &MenuShellClass{native: native}
}

// MenuShellPrivate is a representation of the C record GtkMenuShellPrivate.
type MenuShellPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkMenuShellPrivate that represents the MenuShellPrivate.
func (recv *MenuShellPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// MenuShellPrivateNewFromC creates a new MenuShellPrivate from a pointer to the C GtkMenuShellPrivate that represents the MenuShellPrivate.
func MenuShellPrivateNewFromC(native unsafe.Pointer) *MenuShellPrivate {
	return &MenuShellPrivate{native: native}
}

// MenuToolButtonClass is a representation of the C record GtkMenuToolButtonClass.
type MenuToolButtonClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkMenuToolButtonClass that represents the MenuToolButtonClass.
func (recv *MenuToolButtonClass) ToC() unsafe.Pointer {
	return recv.native
}

// MenuToolButtonClassNewFromC creates a new MenuToolButtonClass from a pointer to the C GtkMenuToolButtonClass that represents the MenuToolButtonClass.
func MenuToolButtonClassNewFromC(native unsafe.Pointer) *MenuToolButtonClass {
	return &MenuToolButtonClass{native: native}
}

// MenuToolButtonPrivate is a representation of the C record GtkMenuToolButtonPrivate.
type MenuToolButtonPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkMenuToolButtonPrivate that represents the MenuToolButtonPrivate.
func (recv *MenuToolButtonPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// MenuToolButtonPrivateNewFromC creates a new MenuToolButtonPrivate from a pointer to the C GtkMenuToolButtonPrivate that represents the MenuToolButtonPrivate.
func MenuToolButtonPrivateNewFromC(native unsafe.Pointer) *MenuToolButtonPrivate {
	return &MenuToolButtonPrivate{native: native}
}

// MessageDialogClass is a representation of the C record GtkMessageDialogClass.
type MessageDialogClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkMessageDialogClass that represents the MessageDialogClass.
func (recv *MessageDialogClass) ToC() unsafe.Pointer {
	return recv.native
}

// MessageDialogClassNewFromC creates a new MessageDialogClass from a pointer to the C GtkMessageDialogClass that represents the MessageDialogClass.
func MessageDialogClassNewFromC(native unsafe.Pointer) *MessageDialogClass {
	return &MessageDialogClass{native: native}
}

// MessageDialogPrivate is a representation of the C record GtkMessageDialogPrivate.
type MessageDialogPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkMessageDialogPrivate that represents the MessageDialogPrivate.
func (recv *MessageDialogPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// MessageDialogPrivateNewFromC creates a new MessageDialogPrivate from a pointer to the C GtkMessageDialogPrivate that represents the MessageDialogPrivate.
func MessageDialogPrivateNewFromC(native unsafe.Pointer) *MessageDialogPrivate {
	return &MessageDialogPrivate{native: native}
}

// MiscClass is a representation of the C record GtkMiscClass.
type MiscClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkMiscClass that represents the MiscClass.
func (recv *MiscClass) ToC() unsafe.Pointer {
	return recv.native
}

// MiscClassNewFromC creates a new MiscClass from a pointer to the C GtkMiscClass that represents the MiscClass.
func MiscClassNewFromC(native unsafe.Pointer) *MiscClass {
	return &MiscClass{native: native}
}

// MiscPrivate is a representation of the C record GtkMiscPrivate.
type MiscPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkMiscPrivate that represents the MiscPrivate.
func (recv *MiscPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// MiscPrivateNewFromC creates a new MiscPrivate from a pointer to the C GtkMiscPrivate that represents the MiscPrivate.
func MiscPrivateNewFromC(native unsafe.Pointer) *MiscPrivate {
	return &MiscPrivate{native: native}
}

// MountOperationClass is a representation of the C record GtkMountOperationClass.
type MountOperationClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkMountOperationClass that represents the MountOperationClass.
func (recv *MountOperationClass) ToC() unsafe.Pointer {
	return recv.native
}

// MountOperationClassNewFromC creates a new MountOperationClass from a pointer to the C GtkMountOperationClass that represents the MountOperationClass.
func MountOperationClassNewFromC(native unsafe.Pointer) *MountOperationClass {
	return &MountOperationClass{native: native}
}

// MountOperationPrivate is a representation of the C record GtkMountOperationPrivate.
type MountOperationPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkMountOperationPrivate that represents the MountOperationPrivate.
func (recv *MountOperationPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// MountOperationPrivateNewFromC creates a new MountOperationPrivate from a pointer to the C GtkMountOperationPrivate that represents the MountOperationPrivate.
func MountOperationPrivateNewFromC(native unsafe.Pointer) *MountOperationPrivate {
	return &MountOperationPrivate{native: native}
}

// NativeDialogClass is a representation of the C record GtkNativeDialogClass.
type NativeDialogClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkNativeDialogClass that represents the NativeDialogClass.
func (recv *NativeDialogClass) ToC() unsafe.Pointer {
	return recv.native
}

// NativeDialogClassNewFromC creates a new NativeDialogClass from a pointer to the C GtkNativeDialogClass that represents the NativeDialogClass.
func NativeDialogClassNewFromC(native unsafe.Pointer) *NativeDialogClass {
	return &NativeDialogClass{native: native}
}

// NotebookAccessibleClass is a representation of the C record GtkNotebookAccessibleClass.
type NotebookAccessibleClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkNotebookAccessibleClass that represents the NotebookAccessibleClass.
func (recv *NotebookAccessibleClass) ToC() unsafe.Pointer {
	return recv.native
}

// NotebookAccessibleClassNewFromC creates a new NotebookAccessibleClass from a pointer to the C GtkNotebookAccessibleClass that represents the NotebookAccessibleClass.
func NotebookAccessibleClassNewFromC(native unsafe.Pointer) *NotebookAccessibleClass {
	return &NotebookAccessibleClass{native: native}
}

// NotebookAccessiblePrivate is a representation of the C record GtkNotebookAccessiblePrivate.
type NotebookAccessiblePrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkNotebookAccessiblePrivate that represents the NotebookAccessiblePrivate.
func (recv *NotebookAccessiblePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// NotebookAccessiblePrivateNewFromC creates a new NotebookAccessiblePrivate from a pointer to the C GtkNotebookAccessiblePrivate that represents the NotebookAccessiblePrivate.
func NotebookAccessiblePrivateNewFromC(native unsafe.Pointer) *NotebookAccessiblePrivate {
	return &NotebookAccessiblePrivate{native: native}
}

// NotebookClass is a representation of the C record GtkNotebookClass.
type NotebookClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkNotebookClass that represents the NotebookClass.
func (recv *NotebookClass) ToC() unsafe.Pointer {
	return recv.native
}

// NotebookClassNewFromC creates a new NotebookClass from a pointer to the C GtkNotebookClass that represents the NotebookClass.
func NotebookClassNewFromC(native unsafe.Pointer) *NotebookClass {
	return &NotebookClass{native: native}
}

// NotebookPageAccessibleClass is a representation of the C record GtkNotebookPageAccessibleClass.
type NotebookPageAccessibleClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkNotebookPageAccessibleClass that represents the NotebookPageAccessibleClass.
func (recv *NotebookPageAccessibleClass) ToC() unsafe.Pointer {
	return recv.native
}

// NotebookPageAccessibleClassNewFromC creates a new NotebookPageAccessibleClass from a pointer to the C GtkNotebookPageAccessibleClass that represents the NotebookPageAccessibleClass.
func NotebookPageAccessibleClassNewFromC(native unsafe.Pointer) *NotebookPageAccessibleClass {
	return &NotebookPageAccessibleClass{native: native}
}

// NotebookPageAccessiblePrivate is a representation of the C record GtkNotebookPageAccessiblePrivate.
type NotebookPageAccessiblePrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkNotebookPageAccessiblePrivate that represents the NotebookPageAccessiblePrivate.
func (recv *NotebookPageAccessiblePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// NotebookPageAccessiblePrivateNewFromC creates a new NotebookPageAccessiblePrivate from a pointer to the C GtkNotebookPageAccessiblePrivate that represents the NotebookPageAccessiblePrivate.
func NotebookPageAccessiblePrivateNewFromC(native unsafe.Pointer) *NotebookPageAccessiblePrivate {
	return &NotebookPageAccessiblePrivate{native: native}
}

// NotebookPrivate is a representation of the C record GtkNotebookPrivate.
type NotebookPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkNotebookPrivate that represents the NotebookPrivate.
func (recv *NotebookPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// NotebookPrivateNewFromC creates a new NotebookPrivate from a pointer to the C GtkNotebookPrivate that represents the NotebookPrivate.
func NotebookPrivateNewFromC(native unsafe.Pointer) *NotebookPrivate {
	return &NotebookPrivate{native: native}
}

// NumerableIconClass is a representation of the C record GtkNumerableIconClass.
type NumerableIconClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkNumerableIconClass that represents the NumerableIconClass.
func (recv *NumerableIconClass) ToC() unsafe.Pointer {
	return recv.native
}

// NumerableIconClassNewFromC creates a new NumerableIconClass from a pointer to the C GtkNumerableIconClass that represents the NumerableIconClass.
func NumerableIconClassNewFromC(native unsafe.Pointer) *NumerableIconClass {
	return &NumerableIconClass{native: native}
}

// NumerableIconPrivate is a representation of the C record GtkNumerableIconPrivate.
type NumerableIconPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkNumerableIconPrivate that represents the NumerableIconPrivate.
func (recv *NumerableIconPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// NumerableIconPrivateNewFromC creates a new NumerableIconPrivate from a pointer to the C GtkNumerableIconPrivate that represents the NumerableIconPrivate.
func NumerableIconPrivateNewFromC(native unsafe.Pointer) *NumerableIconPrivate {
	return &NumerableIconPrivate{native: native}
}

// OffscreenWindowClass is a representation of the C record GtkOffscreenWindowClass.
type OffscreenWindowClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkOffscreenWindowClass that represents the OffscreenWindowClass.
func (recv *OffscreenWindowClass) ToC() unsafe.Pointer {
	return recv.native
}

// OffscreenWindowClassNewFromC creates a new OffscreenWindowClass from a pointer to the C GtkOffscreenWindowClass that represents the OffscreenWindowClass.
func OffscreenWindowClassNewFromC(native unsafe.Pointer) *OffscreenWindowClass {
	return &OffscreenWindowClass{native: native}
}

// OrientableIface is a representation of the C record GtkOrientableIface.
type OrientableIface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkOrientableIface that represents the OrientableIface.
func (recv *OrientableIface) ToC() unsafe.Pointer {
	return recv.native
}

// OrientableIfaceNewFromC creates a new OrientableIface from a pointer to the C GtkOrientableIface that represents the OrientableIface.
func OrientableIfaceNewFromC(native unsafe.Pointer) *OrientableIface {
	return &OrientableIface{native: native}
}

// OverlayClass is a representation of the C record GtkOverlayClass.
type OverlayClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkOverlayClass that represents the OverlayClass.
func (recv *OverlayClass) ToC() unsafe.Pointer {
	return recv.native
}

// OverlayClassNewFromC creates a new OverlayClass from a pointer to the C GtkOverlayClass that represents the OverlayClass.
func OverlayClassNewFromC(native unsafe.Pointer) *OverlayClass {
	return &OverlayClass{native: native}
}

// OverlayPrivate is a representation of the C record GtkOverlayPrivate.
type OverlayPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkOverlayPrivate that represents the OverlayPrivate.
func (recv *OverlayPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// OverlayPrivateNewFromC creates a new OverlayPrivate from a pointer to the C GtkOverlayPrivate that represents the OverlayPrivate.
func OverlayPrivateNewFromC(native unsafe.Pointer) *OverlayPrivate {
	return &OverlayPrivate{native: native}
}

// PadActionEntry is a representation of the C record GtkPadActionEntry.
type PadActionEntry struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkPadActionEntry that represents the PadActionEntry.
func (recv *PadActionEntry) ToC() unsafe.Pointer {
	return recv.native
}

// PadActionEntryNewFromC creates a new PadActionEntry from a pointer to the C GtkPadActionEntry that represents the PadActionEntry.
func PadActionEntryNewFromC(native unsafe.Pointer) *PadActionEntry {
	return &PadActionEntry{native: native}
}

// PadControllerClass is a representation of the C record GtkPadControllerClass.
type PadControllerClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkPadControllerClass that represents the PadControllerClass.
func (recv *PadControllerClass) ToC() unsafe.Pointer {
	return recv.native
}

// PadControllerClassNewFromC creates a new PadControllerClass from a pointer to the C GtkPadControllerClass that represents the PadControllerClass.
func PadControllerClassNewFromC(native unsafe.Pointer) *PadControllerClass {
	return &PadControllerClass{native: native}
}

// PageRange is a representation of the C record GtkPageRange.
type PageRange struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkPageRange that represents the PageRange.
func (recv *PageRange) ToC() unsafe.Pointer {
	return recv.native
}

// PageRangeNewFromC creates a new PageRange from a pointer to the C GtkPageRange that represents the PageRange.
func PageRangeNewFromC(native unsafe.Pointer) *PageRange {
	return &PageRange{native: native}
}

// PanedAccessibleClass is a representation of the C record GtkPanedAccessibleClass.
type PanedAccessibleClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkPanedAccessibleClass that represents the PanedAccessibleClass.
func (recv *PanedAccessibleClass) ToC() unsafe.Pointer {
	return recv.native
}

// PanedAccessibleClassNewFromC creates a new PanedAccessibleClass from a pointer to the C GtkPanedAccessibleClass that represents the PanedAccessibleClass.
func PanedAccessibleClassNewFromC(native unsafe.Pointer) *PanedAccessibleClass {
	return &PanedAccessibleClass{native: native}
}

// PanedAccessiblePrivate is a representation of the C record GtkPanedAccessiblePrivate.
type PanedAccessiblePrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkPanedAccessiblePrivate that represents the PanedAccessiblePrivate.
func (recv *PanedAccessiblePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// PanedAccessiblePrivateNewFromC creates a new PanedAccessiblePrivate from a pointer to the C GtkPanedAccessiblePrivate that represents the PanedAccessiblePrivate.
func PanedAccessiblePrivateNewFromC(native unsafe.Pointer) *PanedAccessiblePrivate {
	return &PanedAccessiblePrivate{native: native}
}

// PanedClass is a representation of the C record GtkPanedClass.
type PanedClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkPanedClass that represents the PanedClass.
func (recv *PanedClass) ToC() unsafe.Pointer {
	return recv.native
}

// PanedClassNewFromC creates a new PanedClass from a pointer to the C GtkPanedClass that represents the PanedClass.
func PanedClassNewFromC(native unsafe.Pointer) *PanedClass {
	return &PanedClass{native: native}
}

// PanedPrivate is a representation of the C record GtkPanedPrivate.
type PanedPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkPanedPrivate that represents the PanedPrivate.
func (recv *PanedPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// PanedPrivateNewFromC creates a new PanedPrivate from a pointer to the C GtkPanedPrivate that represents the PanedPrivate.
func PanedPrivateNewFromC(native unsafe.Pointer) *PanedPrivate {
	return &PanedPrivate{native: native}
}

// PaperSize is a representation of the C record GtkPaperSize.
type PaperSize struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkPaperSize that represents the PaperSize.
func (recv *PaperSize) ToC() unsafe.Pointer {
	return recv.native
}

// PaperSizeNewFromC creates a new PaperSize from a pointer to the C GtkPaperSize that represents the PaperSize.
func PaperSizeNewFromC(native unsafe.Pointer) *PaperSize {
	return &PaperSize{native: native}
}

// PlacesSidebarClass is a representation of the C record GtkPlacesSidebarClass.
type PlacesSidebarClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkPlacesSidebarClass that represents the PlacesSidebarClass.
func (recv *PlacesSidebarClass) ToC() unsafe.Pointer {
	return recv.native
}

// PlacesSidebarClassNewFromC creates a new PlacesSidebarClass from a pointer to the C GtkPlacesSidebarClass that represents the PlacesSidebarClass.
func PlacesSidebarClassNewFromC(native unsafe.Pointer) *PlacesSidebarClass {
	return &PlacesSidebarClass{native: native}
}

// PlugClass is a representation of the C record GtkPlugClass.
type PlugClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkPlugClass that represents the PlugClass.
func (recv *PlugClass) ToC() unsafe.Pointer {
	return recv.native
}

// PlugClassNewFromC creates a new PlugClass from a pointer to the C GtkPlugClass that represents the PlugClass.
func PlugClassNewFromC(native unsafe.Pointer) *PlugClass {
	return &PlugClass{native: native}
}

// PlugPrivate is a representation of the C record GtkPlugPrivate.
type PlugPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkPlugPrivate that represents the PlugPrivate.
func (recv *PlugPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// PlugPrivateNewFromC creates a new PlugPrivate from a pointer to the C GtkPlugPrivate that represents the PlugPrivate.
func PlugPrivateNewFromC(native unsafe.Pointer) *PlugPrivate {
	return &PlugPrivate{native: native}
}

// PopoverAccessibleClass is a representation of the C record GtkPopoverAccessibleClass.
type PopoverAccessibleClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkPopoverAccessibleClass that represents the PopoverAccessibleClass.
func (recv *PopoverAccessibleClass) ToC() unsafe.Pointer {
	return recv.native
}

// PopoverAccessibleClassNewFromC creates a new PopoverAccessibleClass from a pointer to the C GtkPopoverAccessibleClass that represents the PopoverAccessibleClass.
func PopoverAccessibleClassNewFromC(native unsafe.Pointer) *PopoverAccessibleClass {
	return &PopoverAccessibleClass{native: native}
}

// PopoverClass is a representation of the C record GtkPopoverClass.
type PopoverClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkPopoverClass that represents the PopoverClass.
func (recv *PopoverClass) ToC() unsafe.Pointer {
	return recv.native
}

// PopoverClassNewFromC creates a new PopoverClass from a pointer to the C GtkPopoverClass that represents the PopoverClass.
func PopoverClassNewFromC(native unsafe.Pointer) *PopoverClass {
	return &PopoverClass{native: native}
}

// PopoverMenuClass is a representation of the C record GtkPopoverMenuClass.
type PopoverMenuClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkPopoverMenuClass that represents the PopoverMenuClass.
func (recv *PopoverMenuClass) ToC() unsafe.Pointer {
	return recv.native
}

// PopoverMenuClassNewFromC creates a new PopoverMenuClass from a pointer to the C GtkPopoverMenuClass that represents the PopoverMenuClass.
func PopoverMenuClassNewFromC(native unsafe.Pointer) *PopoverMenuClass {
	return &PopoverMenuClass{native: native}
}

// PopoverPrivate is a representation of the C record GtkPopoverPrivate.
type PopoverPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkPopoverPrivate that represents the PopoverPrivate.
func (recv *PopoverPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// PopoverPrivateNewFromC creates a new PopoverPrivate from a pointer to the C GtkPopoverPrivate that represents the PopoverPrivate.
func PopoverPrivateNewFromC(native unsafe.Pointer) *PopoverPrivate {
	return &PopoverPrivate{native: native}
}

// PrintOperationClass is a representation of the C record GtkPrintOperationClass.
type PrintOperationClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkPrintOperationClass that represents the PrintOperationClass.
func (recv *PrintOperationClass) ToC() unsafe.Pointer {
	return recv.native
}

// PrintOperationClassNewFromC creates a new PrintOperationClass from a pointer to the C GtkPrintOperationClass that represents the PrintOperationClass.
func PrintOperationClassNewFromC(native unsafe.Pointer) *PrintOperationClass {
	return &PrintOperationClass{native: native}
}

// PrintOperationPreviewIface is a representation of the C record GtkPrintOperationPreviewIface.
type PrintOperationPreviewIface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkPrintOperationPreviewIface that represents the PrintOperationPreviewIface.
func (recv *PrintOperationPreviewIface) ToC() unsafe.Pointer {
	return recv.native
}

// PrintOperationPreviewIfaceNewFromC creates a new PrintOperationPreviewIface from a pointer to the C GtkPrintOperationPreviewIface that represents the PrintOperationPreviewIface.
func PrintOperationPreviewIfaceNewFromC(native unsafe.Pointer) *PrintOperationPreviewIface {
	return &PrintOperationPreviewIface{native: native}
}

// PrintOperationPrivate is a representation of the C record GtkPrintOperationPrivate.
type PrintOperationPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkPrintOperationPrivate that represents the PrintOperationPrivate.
func (recv *PrintOperationPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// PrintOperationPrivateNewFromC creates a new PrintOperationPrivate from a pointer to the C GtkPrintOperationPrivate that represents the PrintOperationPrivate.
func PrintOperationPrivateNewFromC(native unsafe.Pointer) *PrintOperationPrivate {
	return &PrintOperationPrivate{native: native}
}

// ProgressBarAccessibleClass is a representation of the C record GtkProgressBarAccessibleClass.
type ProgressBarAccessibleClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkProgressBarAccessibleClass that represents the ProgressBarAccessibleClass.
func (recv *ProgressBarAccessibleClass) ToC() unsafe.Pointer {
	return recv.native
}

// ProgressBarAccessibleClassNewFromC creates a new ProgressBarAccessibleClass from a pointer to the C GtkProgressBarAccessibleClass that represents the ProgressBarAccessibleClass.
func ProgressBarAccessibleClassNewFromC(native unsafe.Pointer) *ProgressBarAccessibleClass {
	return &ProgressBarAccessibleClass{native: native}
}

// ProgressBarAccessiblePrivate is a representation of the C record GtkProgressBarAccessiblePrivate.
type ProgressBarAccessiblePrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkProgressBarAccessiblePrivate that represents the ProgressBarAccessiblePrivate.
func (recv *ProgressBarAccessiblePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ProgressBarAccessiblePrivateNewFromC creates a new ProgressBarAccessiblePrivate from a pointer to the C GtkProgressBarAccessiblePrivate that represents the ProgressBarAccessiblePrivate.
func ProgressBarAccessiblePrivateNewFromC(native unsafe.Pointer) *ProgressBarAccessiblePrivate {
	return &ProgressBarAccessiblePrivate{native: native}
}

// ProgressBarClass is a representation of the C record GtkProgressBarClass.
type ProgressBarClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkProgressBarClass that represents the ProgressBarClass.
func (recv *ProgressBarClass) ToC() unsafe.Pointer {
	return recv.native
}

// ProgressBarClassNewFromC creates a new ProgressBarClass from a pointer to the C GtkProgressBarClass that represents the ProgressBarClass.
func ProgressBarClassNewFromC(native unsafe.Pointer) *ProgressBarClass {
	return &ProgressBarClass{native: native}
}

// ProgressBarPrivate is a representation of the C record GtkProgressBarPrivate.
type ProgressBarPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkProgressBarPrivate that represents the ProgressBarPrivate.
func (recv *ProgressBarPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ProgressBarPrivateNewFromC creates a new ProgressBarPrivate from a pointer to the C GtkProgressBarPrivate that represents the ProgressBarPrivate.
func ProgressBarPrivateNewFromC(native unsafe.Pointer) *ProgressBarPrivate {
	return &ProgressBarPrivate{native: native}
}

// RadioActionClass is a representation of the C record GtkRadioActionClass.
type RadioActionClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkRadioActionClass that represents the RadioActionClass.
func (recv *RadioActionClass) ToC() unsafe.Pointer {
	return recv.native
}

// RadioActionClassNewFromC creates a new RadioActionClass from a pointer to the C GtkRadioActionClass that represents the RadioActionClass.
func RadioActionClassNewFromC(native unsafe.Pointer) *RadioActionClass {
	return &RadioActionClass{native: native}
}

// RadioActionEntry is a representation of the C record GtkRadioActionEntry.
type RadioActionEntry struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkRadioActionEntry that represents the RadioActionEntry.
func (recv *RadioActionEntry) ToC() unsafe.Pointer {
	return recv.native
}

// RadioActionEntryNewFromC creates a new RadioActionEntry from a pointer to the C GtkRadioActionEntry that represents the RadioActionEntry.
func RadioActionEntryNewFromC(native unsafe.Pointer) *RadioActionEntry {
	return &RadioActionEntry{native: native}
}

// RadioActionPrivate is a representation of the C record GtkRadioActionPrivate.
type RadioActionPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkRadioActionPrivate that represents the RadioActionPrivate.
func (recv *RadioActionPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// RadioActionPrivateNewFromC creates a new RadioActionPrivate from a pointer to the C GtkRadioActionPrivate that represents the RadioActionPrivate.
func RadioActionPrivateNewFromC(native unsafe.Pointer) *RadioActionPrivate {
	return &RadioActionPrivate{native: native}
}

// RadioButtonAccessibleClass is a representation of the C record GtkRadioButtonAccessibleClass.
type RadioButtonAccessibleClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkRadioButtonAccessibleClass that represents the RadioButtonAccessibleClass.
func (recv *RadioButtonAccessibleClass) ToC() unsafe.Pointer {
	return recv.native
}

// RadioButtonAccessibleClassNewFromC creates a new RadioButtonAccessibleClass from a pointer to the C GtkRadioButtonAccessibleClass that represents the RadioButtonAccessibleClass.
func RadioButtonAccessibleClassNewFromC(native unsafe.Pointer) *RadioButtonAccessibleClass {
	return &RadioButtonAccessibleClass{native: native}
}

// RadioButtonAccessiblePrivate is a representation of the C record GtkRadioButtonAccessiblePrivate.
type RadioButtonAccessiblePrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkRadioButtonAccessiblePrivate that represents the RadioButtonAccessiblePrivate.
func (recv *RadioButtonAccessiblePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// RadioButtonAccessiblePrivateNewFromC creates a new RadioButtonAccessiblePrivate from a pointer to the C GtkRadioButtonAccessiblePrivate that represents the RadioButtonAccessiblePrivate.
func RadioButtonAccessiblePrivateNewFromC(native unsafe.Pointer) *RadioButtonAccessiblePrivate {
	return &RadioButtonAccessiblePrivate{native: native}
}

// RadioButtonClass is a representation of the C record GtkRadioButtonClass.
type RadioButtonClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkRadioButtonClass that represents the RadioButtonClass.
func (recv *RadioButtonClass) ToC() unsafe.Pointer {
	return recv.native
}

// RadioButtonClassNewFromC creates a new RadioButtonClass from a pointer to the C GtkRadioButtonClass that represents the RadioButtonClass.
func RadioButtonClassNewFromC(native unsafe.Pointer) *RadioButtonClass {
	return &RadioButtonClass{native: native}
}

// RadioButtonPrivate is a representation of the C record GtkRadioButtonPrivate.
type RadioButtonPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkRadioButtonPrivate that represents the RadioButtonPrivate.
func (recv *RadioButtonPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// RadioButtonPrivateNewFromC creates a new RadioButtonPrivate from a pointer to the C GtkRadioButtonPrivate that represents the RadioButtonPrivate.
func RadioButtonPrivateNewFromC(native unsafe.Pointer) *RadioButtonPrivate {
	return &RadioButtonPrivate{native: native}
}

// RadioMenuItemAccessibleClass is a representation of the C record GtkRadioMenuItemAccessibleClass.
type RadioMenuItemAccessibleClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkRadioMenuItemAccessibleClass that represents the RadioMenuItemAccessibleClass.
func (recv *RadioMenuItemAccessibleClass) ToC() unsafe.Pointer {
	return recv.native
}

// RadioMenuItemAccessibleClassNewFromC creates a new RadioMenuItemAccessibleClass from a pointer to the C GtkRadioMenuItemAccessibleClass that represents the RadioMenuItemAccessibleClass.
func RadioMenuItemAccessibleClassNewFromC(native unsafe.Pointer) *RadioMenuItemAccessibleClass {
	return &RadioMenuItemAccessibleClass{native: native}
}

// RadioMenuItemAccessiblePrivate is a representation of the C record GtkRadioMenuItemAccessiblePrivate.
type RadioMenuItemAccessiblePrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkRadioMenuItemAccessiblePrivate that represents the RadioMenuItemAccessiblePrivate.
func (recv *RadioMenuItemAccessiblePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// RadioMenuItemAccessiblePrivateNewFromC creates a new RadioMenuItemAccessiblePrivate from a pointer to the C GtkRadioMenuItemAccessiblePrivate that represents the RadioMenuItemAccessiblePrivate.
func RadioMenuItemAccessiblePrivateNewFromC(native unsafe.Pointer) *RadioMenuItemAccessiblePrivate {
	return &RadioMenuItemAccessiblePrivate{native: native}
}

// RadioMenuItemClass is a representation of the C record GtkRadioMenuItemClass.
type RadioMenuItemClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkRadioMenuItemClass that represents the RadioMenuItemClass.
func (recv *RadioMenuItemClass) ToC() unsafe.Pointer {
	return recv.native
}

// RadioMenuItemClassNewFromC creates a new RadioMenuItemClass from a pointer to the C GtkRadioMenuItemClass that represents the RadioMenuItemClass.
func RadioMenuItemClassNewFromC(native unsafe.Pointer) *RadioMenuItemClass {
	return &RadioMenuItemClass{native: native}
}

// RadioMenuItemPrivate is a representation of the C record GtkRadioMenuItemPrivate.
type RadioMenuItemPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkRadioMenuItemPrivate that represents the RadioMenuItemPrivate.
func (recv *RadioMenuItemPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// RadioMenuItemPrivateNewFromC creates a new RadioMenuItemPrivate from a pointer to the C GtkRadioMenuItemPrivate that represents the RadioMenuItemPrivate.
func RadioMenuItemPrivateNewFromC(native unsafe.Pointer) *RadioMenuItemPrivate {
	return &RadioMenuItemPrivate{native: native}
}

// RadioToolButtonClass is a representation of the C record GtkRadioToolButtonClass.
type RadioToolButtonClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkRadioToolButtonClass that represents the RadioToolButtonClass.
func (recv *RadioToolButtonClass) ToC() unsafe.Pointer {
	return recv.native
}

// RadioToolButtonClassNewFromC creates a new RadioToolButtonClass from a pointer to the C GtkRadioToolButtonClass that represents the RadioToolButtonClass.
func RadioToolButtonClassNewFromC(native unsafe.Pointer) *RadioToolButtonClass {
	return &RadioToolButtonClass{native: native}
}

// RangeAccessibleClass is a representation of the C record GtkRangeAccessibleClass.
type RangeAccessibleClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkRangeAccessibleClass that represents the RangeAccessibleClass.
func (recv *RangeAccessibleClass) ToC() unsafe.Pointer {
	return recv.native
}

// RangeAccessibleClassNewFromC creates a new RangeAccessibleClass from a pointer to the C GtkRangeAccessibleClass that represents the RangeAccessibleClass.
func RangeAccessibleClassNewFromC(native unsafe.Pointer) *RangeAccessibleClass {
	return &RangeAccessibleClass{native: native}
}

// RangeAccessiblePrivate is a representation of the C record GtkRangeAccessiblePrivate.
type RangeAccessiblePrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkRangeAccessiblePrivate that represents the RangeAccessiblePrivate.
func (recv *RangeAccessiblePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// RangeAccessiblePrivateNewFromC creates a new RangeAccessiblePrivate from a pointer to the C GtkRangeAccessiblePrivate that represents the RangeAccessiblePrivate.
func RangeAccessiblePrivateNewFromC(native unsafe.Pointer) *RangeAccessiblePrivate {
	return &RangeAccessiblePrivate{native: native}
}

// RangeClass is a representation of the C record GtkRangeClass.
type RangeClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkRangeClass that represents the RangeClass.
func (recv *RangeClass) ToC() unsafe.Pointer {
	return recv.native
}

// RangeClassNewFromC creates a new RangeClass from a pointer to the C GtkRangeClass that represents the RangeClass.
func RangeClassNewFromC(native unsafe.Pointer) *RangeClass {
	return &RangeClass{native: native}
}

// RangePrivate is a representation of the C record GtkRangePrivate.
type RangePrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkRangePrivate that represents the RangePrivate.
func (recv *RangePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// RangePrivateNewFromC creates a new RangePrivate from a pointer to the C GtkRangePrivate that represents the RangePrivate.
func RangePrivateNewFromC(native unsafe.Pointer) *RangePrivate {
	return &RangePrivate{native: native}
}

// RcContext is a representation of the C record GtkRcContext.
type RcContext struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkRcContext that represents the RcContext.
func (recv *RcContext) ToC() unsafe.Pointer {
	return recv.native
}

// RcContextNewFromC creates a new RcContext from a pointer to the C GtkRcContext that represents the RcContext.
func RcContextNewFromC(native unsafe.Pointer) *RcContext {
	return &RcContext{native: native}
}

// RcProperty is a representation of the C record GtkRcProperty.
type RcProperty struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkRcProperty that represents the RcProperty.
func (recv *RcProperty) ToC() unsafe.Pointer {
	return recv.native
}

// RcPropertyNewFromC creates a new RcProperty from a pointer to the C GtkRcProperty that represents the RcProperty.
func RcPropertyNewFromC(native unsafe.Pointer) *RcProperty {
	return &RcProperty{native: native}
}

// RcStyleClass is a representation of the C record GtkRcStyleClass.
type RcStyleClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkRcStyleClass that represents the RcStyleClass.
func (recv *RcStyleClass) ToC() unsafe.Pointer {
	return recv.native
}

// RcStyleClassNewFromC creates a new RcStyleClass from a pointer to the C GtkRcStyleClass that represents the RcStyleClass.
func RcStyleClassNewFromC(native unsafe.Pointer) *RcStyleClass {
	return &RcStyleClass{native: native}
}

// RecentActionClass is a representation of the C record GtkRecentActionClass.
type RecentActionClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkRecentActionClass that represents the RecentActionClass.
func (recv *RecentActionClass) ToC() unsafe.Pointer {
	return recv.native
}

// RecentActionClassNewFromC creates a new RecentActionClass from a pointer to the C GtkRecentActionClass that represents the RecentActionClass.
func RecentActionClassNewFromC(native unsafe.Pointer) *RecentActionClass {
	return &RecentActionClass{native: native}
}

// RecentActionPrivate is a representation of the C record GtkRecentActionPrivate.
type RecentActionPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkRecentActionPrivate that represents the RecentActionPrivate.
func (recv *RecentActionPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// RecentActionPrivateNewFromC creates a new RecentActionPrivate from a pointer to the C GtkRecentActionPrivate that represents the RecentActionPrivate.
func RecentActionPrivateNewFromC(native unsafe.Pointer) *RecentActionPrivate {
	return &RecentActionPrivate{native: native}
}

// RecentChooserDialogClass is a representation of the C record GtkRecentChooserDialogClass.
type RecentChooserDialogClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkRecentChooserDialogClass that represents the RecentChooserDialogClass.
func (recv *RecentChooserDialogClass) ToC() unsafe.Pointer {
	return recv.native
}

// RecentChooserDialogClassNewFromC creates a new RecentChooserDialogClass from a pointer to the C GtkRecentChooserDialogClass that represents the RecentChooserDialogClass.
func RecentChooserDialogClassNewFromC(native unsafe.Pointer) *RecentChooserDialogClass {
	return &RecentChooserDialogClass{native: native}
}

// RecentChooserDialogPrivate is a representation of the C record GtkRecentChooserDialogPrivate.
type RecentChooserDialogPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkRecentChooserDialogPrivate that represents the RecentChooserDialogPrivate.
func (recv *RecentChooserDialogPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// RecentChooserDialogPrivateNewFromC creates a new RecentChooserDialogPrivate from a pointer to the C GtkRecentChooserDialogPrivate that represents the RecentChooserDialogPrivate.
func RecentChooserDialogPrivateNewFromC(native unsafe.Pointer) *RecentChooserDialogPrivate {
	return &RecentChooserDialogPrivate{native: native}
}

// RecentChooserIface is a representation of the C record GtkRecentChooserIface.
type RecentChooserIface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkRecentChooserIface that represents the RecentChooserIface.
func (recv *RecentChooserIface) ToC() unsafe.Pointer {
	return recv.native
}

// RecentChooserIfaceNewFromC creates a new RecentChooserIface from a pointer to the C GtkRecentChooserIface that represents the RecentChooserIface.
func RecentChooserIfaceNewFromC(native unsafe.Pointer) *RecentChooserIface {
	return &RecentChooserIface{native: native}
}

// RecentChooserMenuClass is a representation of the C record GtkRecentChooserMenuClass.
type RecentChooserMenuClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkRecentChooserMenuClass that represents the RecentChooserMenuClass.
func (recv *RecentChooserMenuClass) ToC() unsafe.Pointer {
	return recv.native
}

// RecentChooserMenuClassNewFromC creates a new RecentChooserMenuClass from a pointer to the C GtkRecentChooserMenuClass that represents the RecentChooserMenuClass.
func RecentChooserMenuClassNewFromC(native unsafe.Pointer) *RecentChooserMenuClass {
	return &RecentChooserMenuClass{native: native}
}

// RecentChooserMenuPrivate is a representation of the C record GtkRecentChooserMenuPrivate.
type RecentChooserMenuPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkRecentChooserMenuPrivate that represents the RecentChooserMenuPrivate.
func (recv *RecentChooserMenuPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// RecentChooserMenuPrivateNewFromC creates a new RecentChooserMenuPrivate from a pointer to the C GtkRecentChooserMenuPrivate that represents the RecentChooserMenuPrivate.
func RecentChooserMenuPrivateNewFromC(native unsafe.Pointer) *RecentChooserMenuPrivate {
	return &RecentChooserMenuPrivate{native: native}
}

// RecentChooserWidgetClass is a representation of the C record GtkRecentChooserWidgetClass.
type RecentChooserWidgetClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkRecentChooserWidgetClass that represents the RecentChooserWidgetClass.
func (recv *RecentChooserWidgetClass) ToC() unsafe.Pointer {
	return recv.native
}

// RecentChooserWidgetClassNewFromC creates a new RecentChooserWidgetClass from a pointer to the C GtkRecentChooserWidgetClass that represents the RecentChooserWidgetClass.
func RecentChooserWidgetClassNewFromC(native unsafe.Pointer) *RecentChooserWidgetClass {
	return &RecentChooserWidgetClass{native: native}
}

// RecentChooserWidgetPrivate is a representation of the C record GtkRecentChooserWidgetPrivate.
type RecentChooserWidgetPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkRecentChooserWidgetPrivate that represents the RecentChooserWidgetPrivate.
func (recv *RecentChooserWidgetPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// RecentChooserWidgetPrivateNewFromC creates a new RecentChooserWidgetPrivate from a pointer to the C GtkRecentChooserWidgetPrivate that represents the RecentChooserWidgetPrivate.
func RecentChooserWidgetPrivateNewFromC(native unsafe.Pointer) *RecentChooserWidgetPrivate {
	return &RecentChooserWidgetPrivate{native: native}
}

// RecentData is a representation of the C record GtkRecentData.
type RecentData struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkRecentData that represents the RecentData.
func (recv *RecentData) ToC() unsafe.Pointer {
	return recv.native
}

// RecentDataNewFromC creates a new RecentData from a pointer to the C GtkRecentData that represents the RecentData.
func RecentDataNewFromC(native unsafe.Pointer) *RecentData {
	return &RecentData{native: native}
}

// RecentFilterInfo is a representation of the C record GtkRecentFilterInfo.
type RecentFilterInfo struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkRecentFilterInfo that represents the RecentFilterInfo.
func (recv *RecentFilterInfo) ToC() unsafe.Pointer {
	return recv.native
}

// RecentFilterInfoNewFromC creates a new RecentFilterInfo from a pointer to the C GtkRecentFilterInfo that represents the RecentFilterInfo.
func RecentFilterInfoNewFromC(native unsafe.Pointer) *RecentFilterInfo {
	return &RecentFilterInfo{native: native}
}

// RecentManagerPrivate is a representation of the C record GtkRecentManagerPrivate.
type RecentManagerPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkRecentManagerPrivate that represents the RecentManagerPrivate.
func (recv *RecentManagerPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// RecentManagerPrivateNewFromC creates a new RecentManagerPrivate from a pointer to the C GtkRecentManagerPrivate that represents the RecentManagerPrivate.
func RecentManagerPrivateNewFromC(native unsafe.Pointer) *RecentManagerPrivate {
	return &RecentManagerPrivate{native: native}
}

// RendererCellAccessibleClass is a representation of the C record GtkRendererCellAccessibleClass.
type RendererCellAccessibleClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkRendererCellAccessibleClass that represents the RendererCellAccessibleClass.
func (recv *RendererCellAccessibleClass) ToC() unsafe.Pointer {
	return recv.native
}

// RendererCellAccessibleClassNewFromC creates a new RendererCellAccessibleClass from a pointer to the C GtkRendererCellAccessibleClass that represents the RendererCellAccessibleClass.
func RendererCellAccessibleClassNewFromC(native unsafe.Pointer) *RendererCellAccessibleClass {
	return &RendererCellAccessibleClass{native: native}
}

// RendererCellAccessiblePrivate is a representation of the C record GtkRendererCellAccessiblePrivate.
type RendererCellAccessiblePrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkRendererCellAccessiblePrivate that represents the RendererCellAccessiblePrivate.
func (recv *RendererCellAccessiblePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// RendererCellAccessiblePrivateNewFromC creates a new RendererCellAccessiblePrivate from a pointer to the C GtkRendererCellAccessiblePrivate that represents the RendererCellAccessiblePrivate.
func RendererCellAccessiblePrivateNewFromC(native unsafe.Pointer) *RendererCellAccessiblePrivate {
	return &RendererCellAccessiblePrivate{native: native}
}

// RequestedSize is a representation of the C record GtkRequestedSize.
type RequestedSize struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkRequestedSize that represents the RequestedSize.
func (recv *RequestedSize) ToC() unsafe.Pointer {
	return recv.native
}

// RequestedSizeNewFromC creates a new RequestedSize from a pointer to the C GtkRequestedSize that represents the RequestedSize.
func RequestedSizeNewFromC(native unsafe.Pointer) *RequestedSize {
	return &RequestedSize{native: native}
}

// Requisition is a representation of the C record GtkRequisition.
type Requisition struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkRequisition that represents the Requisition.
func (recv *Requisition) ToC() unsafe.Pointer {
	return recv.native
}

// RequisitionNewFromC creates a new Requisition from a pointer to the C GtkRequisition that represents the Requisition.
func RequisitionNewFromC(native unsafe.Pointer) *Requisition {
	return &Requisition{native: native}
}

// RevealerClass is a representation of the C record GtkRevealerClass.
type RevealerClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkRevealerClass that represents the RevealerClass.
func (recv *RevealerClass) ToC() unsafe.Pointer {
	return recv.native
}

// RevealerClassNewFromC creates a new RevealerClass from a pointer to the C GtkRevealerClass that represents the RevealerClass.
func RevealerClassNewFromC(native unsafe.Pointer) *RevealerClass {
	return &RevealerClass{native: native}
}

// ScaleAccessibleClass is a representation of the C record GtkScaleAccessibleClass.
type ScaleAccessibleClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkScaleAccessibleClass that represents the ScaleAccessibleClass.
func (recv *ScaleAccessibleClass) ToC() unsafe.Pointer {
	return recv.native
}

// ScaleAccessibleClassNewFromC creates a new ScaleAccessibleClass from a pointer to the C GtkScaleAccessibleClass that represents the ScaleAccessibleClass.
func ScaleAccessibleClassNewFromC(native unsafe.Pointer) *ScaleAccessibleClass {
	return &ScaleAccessibleClass{native: native}
}

// ScaleAccessiblePrivate is a representation of the C record GtkScaleAccessiblePrivate.
type ScaleAccessiblePrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkScaleAccessiblePrivate that represents the ScaleAccessiblePrivate.
func (recv *ScaleAccessiblePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ScaleAccessiblePrivateNewFromC creates a new ScaleAccessiblePrivate from a pointer to the C GtkScaleAccessiblePrivate that represents the ScaleAccessiblePrivate.
func ScaleAccessiblePrivateNewFromC(native unsafe.Pointer) *ScaleAccessiblePrivate {
	return &ScaleAccessiblePrivate{native: native}
}

// ScaleButtonAccessibleClass is a representation of the C record GtkScaleButtonAccessibleClass.
type ScaleButtonAccessibleClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkScaleButtonAccessibleClass that represents the ScaleButtonAccessibleClass.
func (recv *ScaleButtonAccessibleClass) ToC() unsafe.Pointer {
	return recv.native
}

// ScaleButtonAccessibleClassNewFromC creates a new ScaleButtonAccessibleClass from a pointer to the C GtkScaleButtonAccessibleClass that represents the ScaleButtonAccessibleClass.
func ScaleButtonAccessibleClassNewFromC(native unsafe.Pointer) *ScaleButtonAccessibleClass {
	return &ScaleButtonAccessibleClass{native: native}
}

// ScaleButtonAccessiblePrivate is a representation of the C record GtkScaleButtonAccessiblePrivate.
type ScaleButtonAccessiblePrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkScaleButtonAccessiblePrivate that represents the ScaleButtonAccessiblePrivate.
func (recv *ScaleButtonAccessiblePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ScaleButtonAccessiblePrivateNewFromC creates a new ScaleButtonAccessiblePrivate from a pointer to the C GtkScaleButtonAccessiblePrivate that represents the ScaleButtonAccessiblePrivate.
func ScaleButtonAccessiblePrivateNewFromC(native unsafe.Pointer) *ScaleButtonAccessiblePrivate {
	return &ScaleButtonAccessiblePrivate{native: native}
}

// ScaleButtonClass is a representation of the C record GtkScaleButtonClass.
type ScaleButtonClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkScaleButtonClass that represents the ScaleButtonClass.
func (recv *ScaleButtonClass) ToC() unsafe.Pointer {
	return recv.native
}

// ScaleButtonClassNewFromC creates a new ScaleButtonClass from a pointer to the C GtkScaleButtonClass that represents the ScaleButtonClass.
func ScaleButtonClassNewFromC(native unsafe.Pointer) *ScaleButtonClass {
	return &ScaleButtonClass{native: native}
}

// ScaleButtonPrivate is a representation of the C record GtkScaleButtonPrivate.
type ScaleButtonPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkScaleButtonPrivate that represents the ScaleButtonPrivate.
func (recv *ScaleButtonPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ScaleButtonPrivateNewFromC creates a new ScaleButtonPrivate from a pointer to the C GtkScaleButtonPrivate that represents the ScaleButtonPrivate.
func ScaleButtonPrivateNewFromC(native unsafe.Pointer) *ScaleButtonPrivate {
	return &ScaleButtonPrivate{native: native}
}

// ScaleClass is a representation of the C record GtkScaleClass.
type ScaleClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkScaleClass that represents the ScaleClass.
func (recv *ScaleClass) ToC() unsafe.Pointer {
	return recv.native
}

// ScaleClassNewFromC creates a new ScaleClass from a pointer to the C GtkScaleClass that represents the ScaleClass.
func ScaleClassNewFromC(native unsafe.Pointer) *ScaleClass {
	return &ScaleClass{native: native}
}

// ScalePrivate is a representation of the C record GtkScalePrivate.
type ScalePrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkScalePrivate that represents the ScalePrivate.
func (recv *ScalePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ScalePrivateNewFromC creates a new ScalePrivate from a pointer to the C GtkScalePrivate that represents the ScalePrivate.
func ScalePrivateNewFromC(native unsafe.Pointer) *ScalePrivate {
	return &ScalePrivate{native: native}
}

// ScrollableInterface is a representation of the C record GtkScrollableInterface.
type ScrollableInterface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkScrollableInterface that represents the ScrollableInterface.
func (recv *ScrollableInterface) ToC() unsafe.Pointer {
	return recv.native
}

// ScrollableInterfaceNewFromC creates a new ScrollableInterface from a pointer to the C GtkScrollableInterface that represents the ScrollableInterface.
func ScrollableInterfaceNewFromC(native unsafe.Pointer) *ScrollableInterface {
	return &ScrollableInterface{native: native}
}

// ScrollbarClass is a representation of the C record GtkScrollbarClass.
type ScrollbarClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkScrollbarClass that represents the ScrollbarClass.
func (recv *ScrollbarClass) ToC() unsafe.Pointer {
	return recv.native
}

// ScrollbarClassNewFromC creates a new ScrollbarClass from a pointer to the C GtkScrollbarClass that represents the ScrollbarClass.
func ScrollbarClassNewFromC(native unsafe.Pointer) *ScrollbarClass {
	return &ScrollbarClass{native: native}
}

// ScrolledWindowAccessibleClass is a representation of the C record GtkScrolledWindowAccessibleClass.
type ScrolledWindowAccessibleClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkScrolledWindowAccessibleClass that represents the ScrolledWindowAccessibleClass.
func (recv *ScrolledWindowAccessibleClass) ToC() unsafe.Pointer {
	return recv.native
}

// ScrolledWindowAccessibleClassNewFromC creates a new ScrolledWindowAccessibleClass from a pointer to the C GtkScrolledWindowAccessibleClass that represents the ScrolledWindowAccessibleClass.
func ScrolledWindowAccessibleClassNewFromC(native unsafe.Pointer) *ScrolledWindowAccessibleClass {
	return &ScrolledWindowAccessibleClass{native: native}
}

// ScrolledWindowAccessiblePrivate is a representation of the C record GtkScrolledWindowAccessiblePrivate.
type ScrolledWindowAccessiblePrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkScrolledWindowAccessiblePrivate that represents the ScrolledWindowAccessiblePrivate.
func (recv *ScrolledWindowAccessiblePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ScrolledWindowAccessiblePrivateNewFromC creates a new ScrolledWindowAccessiblePrivate from a pointer to the C GtkScrolledWindowAccessiblePrivate that represents the ScrolledWindowAccessiblePrivate.
func ScrolledWindowAccessiblePrivateNewFromC(native unsafe.Pointer) *ScrolledWindowAccessiblePrivate {
	return &ScrolledWindowAccessiblePrivate{native: native}
}

// ScrolledWindowClass is a representation of the C record GtkScrolledWindowClass.
type ScrolledWindowClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkScrolledWindowClass that represents the ScrolledWindowClass.
func (recv *ScrolledWindowClass) ToC() unsafe.Pointer {
	return recv.native
}

// ScrolledWindowClassNewFromC creates a new ScrolledWindowClass from a pointer to the C GtkScrolledWindowClass that represents the ScrolledWindowClass.
func ScrolledWindowClassNewFromC(native unsafe.Pointer) *ScrolledWindowClass {
	return &ScrolledWindowClass{native: native}
}

// ScrolledWindowPrivate is a representation of the C record GtkScrolledWindowPrivate.
type ScrolledWindowPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkScrolledWindowPrivate that represents the ScrolledWindowPrivate.
func (recv *ScrolledWindowPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ScrolledWindowPrivateNewFromC creates a new ScrolledWindowPrivate from a pointer to the C GtkScrolledWindowPrivate that represents the ScrolledWindowPrivate.
func ScrolledWindowPrivateNewFromC(native unsafe.Pointer) *ScrolledWindowPrivate {
	return &ScrolledWindowPrivate{native: native}
}

// SearchBarClass is a representation of the C record GtkSearchBarClass.
type SearchBarClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkSearchBarClass that represents the SearchBarClass.
func (recv *SearchBarClass) ToC() unsafe.Pointer {
	return recv.native
}

// SearchBarClassNewFromC creates a new SearchBarClass from a pointer to the C GtkSearchBarClass that represents the SearchBarClass.
func SearchBarClassNewFromC(native unsafe.Pointer) *SearchBarClass {
	return &SearchBarClass{native: native}
}

// SearchEntryClass is a representation of the C record GtkSearchEntryClass.
type SearchEntryClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkSearchEntryClass that represents the SearchEntryClass.
func (recv *SearchEntryClass) ToC() unsafe.Pointer {
	return recv.native
}

// SearchEntryClassNewFromC creates a new SearchEntryClass from a pointer to the C GtkSearchEntryClass that represents the SearchEntryClass.
func SearchEntryClassNewFromC(native unsafe.Pointer) *SearchEntryClass {
	return &SearchEntryClass{native: native}
}

// SelectionData is a representation of the C record GtkSelectionData.
type SelectionData struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkSelectionData that represents the SelectionData.
func (recv *SelectionData) ToC() unsafe.Pointer {
	return recv.native
}

// SelectionDataNewFromC creates a new SelectionData from a pointer to the C GtkSelectionData that represents the SelectionData.
func SelectionDataNewFromC(native unsafe.Pointer) *SelectionData {
	return &SelectionData{native: native}
}

// SeparatorClass is a representation of the C record GtkSeparatorClass.
type SeparatorClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkSeparatorClass that represents the SeparatorClass.
func (recv *SeparatorClass) ToC() unsafe.Pointer {
	return recv.native
}

// SeparatorClassNewFromC creates a new SeparatorClass from a pointer to the C GtkSeparatorClass that represents the SeparatorClass.
func SeparatorClassNewFromC(native unsafe.Pointer) *SeparatorClass {
	return &SeparatorClass{native: native}
}

// SeparatorMenuItemClass is a representation of the C record GtkSeparatorMenuItemClass.
type SeparatorMenuItemClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkSeparatorMenuItemClass that represents the SeparatorMenuItemClass.
func (recv *SeparatorMenuItemClass) ToC() unsafe.Pointer {
	return recv.native
}

// SeparatorMenuItemClassNewFromC creates a new SeparatorMenuItemClass from a pointer to the C GtkSeparatorMenuItemClass that represents the SeparatorMenuItemClass.
func SeparatorMenuItemClassNewFromC(native unsafe.Pointer) *SeparatorMenuItemClass {
	return &SeparatorMenuItemClass{native: native}
}

// SeparatorPrivate is a representation of the C record GtkSeparatorPrivate.
type SeparatorPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkSeparatorPrivate that represents the SeparatorPrivate.
func (recv *SeparatorPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// SeparatorPrivateNewFromC creates a new SeparatorPrivate from a pointer to the C GtkSeparatorPrivate that represents the SeparatorPrivate.
func SeparatorPrivateNewFromC(native unsafe.Pointer) *SeparatorPrivate {
	return &SeparatorPrivate{native: native}
}

// SeparatorToolItemClass is a representation of the C record GtkSeparatorToolItemClass.
type SeparatorToolItemClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkSeparatorToolItemClass that represents the SeparatorToolItemClass.
func (recv *SeparatorToolItemClass) ToC() unsafe.Pointer {
	return recv.native
}

// SeparatorToolItemClassNewFromC creates a new SeparatorToolItemClass from a pointer to the C GtkSeparatorToolItemClass that represents the SeparatorToolItemClass.
func SeparatorToolItemClassNewFromC(native unsafe.Pointer) *SeparatorToolItemClass {
	return &SeparatorToolItemClass{native: native}
}

// SeparatorToolItemPrivate is a representation of the C record GtkSeparatorToolItemPrivate.
type SeparatorToolItemPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkSeparatorToolItemPrivate that represents the SeparatorToolItemPrivate.
func (recv *SeparatorToolItemPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// SeparatorToolItemPrivateNewFromC creates a new SeparatorToolItemPrivate from a pointer to the C GtkSeparatorToolItemPrivate that represents the SeparatorToolItemPrivate.
func SeparatorToolItemPrivateNewFromC(native unsafe.Pointer) *SeparatorToolItemPrivate {
	return &SeparatorToolItemPrivate{native: native}
}

// SettingsClass is a representation of the C record GtkSettingsClass.
type SettingsClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkSettingsClass that represents the SettingsClass.
func (recv *SettingsClass) ToC() unsafe.Pointer {
	return recv.native
}

// SettingsClassNewFromC creates a new SettingsClass from a pointer to the C GtkSettingsClass that represents the SettingsClass.
func SettingsClassNewFromC(native unsafe.Pointer) *SettingsClass {
	return &SettingsClass{native: native}
}

// SettingsPrivate is a representation of the C record GtkSettingsPrivate.
type SettingsPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkSettingsPrivate that represents the SettingsPrivate.
func (recv *SettingsPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// SettingsPrivateNewFromC creates a new SettingsPrivate from a pointer to the C GtkSettingsPrivate that represents the SettingsPrivate.
func SettingsPrivateNewFromC(native unsafe.Pointer) *SettingsPrivate {
	return &SettingsPrivate{native: native}
}

// SettingsValue is a representation of the C record GtkSettingsValue.
type SettingsValue struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkSettingsValue that represents the SettingsValue.
func (recv *SettingsValue) ToC() unsafe.Pointer {
	return recv.native
}

// SettingsValueNewFromC creates a new SettingsValue from a pointer to the C GtkSettingsValue that represents the SettingsValue.
func SettingsValueNewFromC(native unsafe.Pointer) *SettingsValue {
	return &SettingsValue{native: native}
}

// ShortcutLabelClass is a representation of the C record GtkShortcutLabelClass.
type ShortcutLabelClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkShortcutLabelClass that represents the ShortcutLabelClass.
func (recv *ShortcutLabelClass) ToC() unsafe.Pointer {
	return recv.native
}

// ShortcutLabelClassNewFromC creates a new ShortcutLabelClass from a pointer to the C GtkShortcutLabelClass that represents the ShortcutLabelClass.
func ShortcutLabelClassNewFromC(native unsafe.Pointer) *ShortcutLabelClass {
	return &ShortcutLabelClass{native: native}
}

// ShortcutsGroupClass is a representation of the C record GtkShortcutsGroupClass.
type ShortcutsGroupClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkShortcutsGroupClass that represents the ShortcutsGroupClass.
func (recv *ShortcutsGroupClass) ToC() unsafe.Pointer {
	return recv.native
}

// ShortcutsGroupClassNewFromC creates a new ShortcutsGroupClass from a pointer to the C GtkShortcutsGroupClass that represents the ShortcutsGroupClass.
func ShortcutsGroupClassNewFromC(native unsafe.Pointer) *ShortcutsGroupClass {
	return &ShortcutsGroupClass{native: native}
}

// ShortcutsSectionClass is a representation of the C record GtkShortcutsSectionClass.
type ShortcutsSectionClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkShortcutsSectionClass that represents the ShortcutsSectionClass.
func (recv *ShortcutsSectionClass) ToC() unsafe.Pointer {
	return recv.native
}

// ShortcutsSectionClassNewFromC creates a new ShortcutsSectionClass from a pointer to the C GtkShortcutsSectionClass that represents the ShortcutsSectionClass.
func ShortcutsSectionClassNewFromC(native unsafe.Pointer) *ShortcutsSectionClass {
	return &ShortcutsSectionClass{native: native}
}

// ShortcutsShortcutClass is a representation of the C record GtkShortcutsShortcutClass.
type ShortcutsShortcutClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkShortcutsShortcutClass that represents the ShortcutsShortcutClass.
func (recv *ShortcutsShortcutClass) ToC() unsafe.Pointer {
	return recv.native
}

// ShortcutsShortcutClassNewFromC creates a new ShortcutsShortcutClass from a pointer to the C GtkShortcutsShortcutClass that represents the ShortcutsShortcutClass.
func ShortcutsShortcutClassNewFromC(native unsafe.Pointer) *ShortcutsShortcutClass {
	return &ShortcutsShortcutClass{native: native}
}

// ShortcutsWindowClass is a representation of the C record GtkShortcutsWindowClass.
type ShortcutsWindowClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkShortcutsWindowClass that represents the ShortcutsWindowClass.
func (recv *ShortcutsWindowClass) ToC() unsafe.Pointer {
	return recv.native
}

// ShortcutsWindowClassNewFromC creates a new ShortcutsWindowClass from a pointer to the C GtkShortcutsWindowClass that represents the ShortcutsWindowClass.
func ShortcutsWindowClassNewFromC(native unsafe.Pointer) *ShortcutsWindowClass {
	return &ShortcutsWindowClass{native: native}
}

// SizeGroupClass is a representation of the C record GtkSizeGroupClass.
type SizeGroupClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkSizeGroupClass that represents the SizeGroupClass.
func (recv *SizeGroupClass) ToC() unsafe.Pointer {
	return recv.native
}

// SizeGroupClassNewFromC creates a new SizeGroupClass from a pointer to the C GtkSizeGroupClass that represents the SizeGroupClass.
func SizeGroupClassNewFromC(native unsafe.Pointer) *SizeGroupClass {
	return &SizeGroupClass{native: native}
}

// SizeGroupPrivate is a representation of the C record GtkSizeGroupPrivate.
type SizeGroupPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkSizeGroupPrivate that represents the SizeGroupPrivate.
func (recv *SizeGroupPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// SizeGroupPrivateNewFromC creates a new SizeGroupPrivate from a pointer to the C GtkSizeGroupPrivate that represents the SizeGroupPrivate.
func SizeGroupPrivateNewFromC(native unsafe.Pointer) *SizeGroupPrivate {
	return &SizeGroupPrivate{native: native}
}

// SocketClass is a representation of the C record GtkSocketClass.
type SocketClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkSocketClass that represents the SocketClass.
func (recv *SocketClass) ToC() unsafe.Pointer {
	return recv.native
}

// SocketClassNewFromC creates a new SocketClass from a pointer to the C GtkSocketClass that represents the SocketClass.
func SocketClassNewFromC(native unsafe.Pointer) *SocketClass {
	return &SocketClass{native: native}
}

// SocketPrivate is a representation of the C record GtkSocketPrivate.
type SocketPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkSocketPrivate that represents the SocketPrivate.
func (recv *SocketPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// SocketPrivateNewFromC creates a new SocketPrivate from a pointer to the C GtkSocketPrivate that represents the SocketPrivate.
func SocketPrivateNewFromC(native unsafe.Pointer) *SocketPrivate {
	return &SocketPrivate{native: native}
}

// SpinButtonAccessibleClass is a representation of the C record GtkSpinButtonAccessibleClass.
type SpinButtonAccessibleClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkSpinButtonAccessibleClass that represents the SpinButtonAccessibleClass.
func (recv *SpinButtonAccessibleClass) ToC() unsafe.Pointer {
	return recv.native
}

// SpinButtonAccessibleClassNewFromC creates a new SpinButtonAccessibleClass from a pointer to the C GtkSpinButtonAccessibleClass that represents the SpinButtonAccessibleClass.
func SpinButtonAccessibleClassNewFromC(native unsafe.Pointer) *SpinButtonAccessibleClass {
	return &SpinButtonAccessibleClass{native: native}
}

// SpinButtonAccessiblePrivate is a representation of the C record GtkSpinButtonAccessiblePrivate.
type SpinButtonAccessiblePrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkSpinButtonAccessiblePrivate that represents the SpinButtonAccessiblePrivate.
func (recv *SpinButtonAccessiblePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// SpinButtonAccessiblePrivateNewFromC creates a new SpinButtonAccessiblePrivate from a pointer to the C GtkSpinButtonAccessiblePrivate that represents the SpinButtonAccessiblePrivate.
func SpinButtonAccessiblePrivateNewFromC(native unsafe.Pointer) *SpinButtonAccessiblePrivate {
	return &SpinButtonAccessiblePrivate{native: native}
}

// SpinButtonClass is a representation of the C record GtkSpinButtonClass.
type SpinButtonClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkSpinButtonClass that represents the SpinButtonClass.
func (recv *SpinButtonClass) ToC() unsafe.Pointer {
	return recv.native
}

// SpinButtonClassNewFromC creates a new SpinButtonClass from a pointer to the C GtkSpinButtonClass that represents the SpinButtonClass.
func SpinButtonClassNewFromC(native unsafe.Pointer) *SpinButtonClass {
	return &SpinButtonClass{native: native}
}

// SpinButtonPrivate is a representation of the C record GtkSpinButtonPrivate.
type SpinButtonPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkSpinButtonPrivate that represents the SpinButtonPrivate.
func (recv *SpinButtonPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// SpinButtonPrivateNewFromC creates a new SpinButtonPrivate from a pointer to the C GtkSpinButtonPrivate that represents the SpinButtonPrivate.
func SpinButtonPrivateNewFromC(native unsafe.Pointer) *SpinButtonPrivate {
	return &SpinButtonPrivate{native: native}
}

// SpinnerAccessibleClass is a representation of the C record GtkSpinnerAccessibleClass.
type SpinnerAccessibleClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkSpinnerAccessibleClass that represents the SpinnerAccessibleClass.
func (recv *SpinnerAccessibleClass) ToC() unsafe.Pointer {
	return recv.native
}

// SpinnerAccessibleClassNewFromC creates a new SpinnerAccessibleClass from a pointer to the C GtkSpinnerAccessibleClass that represents the SpinnerAccessibleClass.
func SpinnerAccessibleClassNewFromC(native unsafe.Pointer) *SpinnerAccessibleClass {
	return &SpinnerAccessibleClass{native: native}
}

// SpinnerAccessiblePrivate is a representation of the C record GtkSpinnerAccessiblePrivate.
type SpinnerAccessiblePrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkSpinnerAccessiblePrivate that represents the SpinnerAccessiblePrivate.
func (recv *SpinnerAccessiblePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// SpinnerAccessiblePrivateNewFromC creates a new SpinnerAccessiblePrivate from a pointer to the C GtkSpinnerAccessiblePrivate that represents the SpinnerAccessiblePrivate.
func SpinnerAccessiblePrivateNewFromC(native unsafe.Pointer) *SpinnerAccessiblePrivate {
	return &SpinnerAccessiblePrivate{native: native}
}

// SpinnerClass is a representation of the C record GtkSpinnerClass.
type SpinnerClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkSpinnerClass that represents the SpinnerClass.
func (recv *SpinnerClass) ToC() unsafe.Pointer {
	return recv.native
}

// SpinnerClassNewFromC creates a new SpinnerClass from a pointer to the C GtkSpinnerClass that represents the SpinnerClass.
func SpinnerClassNewFromC(native unsafe.Pointer) *SpinnerClass {
	return &SpinnerClass{native: native}
}

// SpinnerPrivate is a representation of the C record GtkSpinnerPrivate.
type SpinnerPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkSpinnerPrivate that represents the SpinnerPrivate.
func (recv *SpinnerPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// SpinnerPrivateNewFromC creates a new SpinnerPrivate from a pointer to the C GtkSpinnerPrivate that represents the SpinnerPrivate.
func SpinnerPrivateNewFromC(native unsafe.Pointer) *SpinnerPrivate {
	return &SpinnerPrivate{native: native}
}

// UNSUPPORTED : StackAccessibleClass : blacklisted

// StackClass is a representation of the C record GtkStackClass.
type StackClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkStackClass that represents the StackClass.
func (recv *StackClass) ToC() unsafe.Pointer {
	return recv.native
}

// StackClassNewFromC creates a new StackClass from a pointer to the C GtkStackClass that represents the StackClass.
func StackClassNewFromC(native unsafe.Pointer) *StackClass {
	return &StackClass{native: native}
}

// StackSidebarClass is a representation of the C record GtkStackSidebarClass.
type StackSidebarClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkStackSidebarClass that represents the StackSidebarClass.
func (recv *StackSidebarClass) ToC() unsafe.Pointer {
	return recv.native
}

// StackSidebarClassNewFromC creates a new StackSidebarClass from a pointer to the C GtkStackSidebarClass that represents the StackSidebarClass.
func StackSidebarClassNewFromC(native unsafe.Pointer) *StackSidebarClass {
	return &StackSidebarClass{native: native}
}

// StackSidebarPrivate is a representation of the C record GtkStackSidebarPrivate.
type StackSidebarPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkStackSidebarPrivate that represents the StackSidebarPrivate.
func (recv *StackSidebarPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// StackSidebarPrivateNewFromC creates a new StackSidebarPrivate from a pointer to the C GtkStackSidebarPrivate that represents the StackSidebarPrivate.
func StackSidebarPrivateNewFromC(native unsafe.Pointer) *StackSidebarPrivate {
	return &StackSidebarPrivate{native: native}
}

// StackSwitcherClass is a representation of the C record GtkStackSwitcherClass.
type StackSwitcherClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkStackSwitcherClass that represents the StackSwitcherClass.
func (recv *StackSwitcherClass) ToC() unsafe.Pointer {
	return recv.native
}

// StackSwitcherClassNewFromC creates a new StackSwitcherClass from a pointer to the C GtkStackSwitcherClass that represents the StackSwitcherClass.
func StackSwitcherClassNewFromC(native unsafe.Pointer) *StackSwitcherClass {
	return &StackSwitcherClass{native: native}
}

// StatusIconClass is a representation of the C record GtkStatusIconClass.
type StatusIconClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkStatusIconClass that represents the StatusIconClass.
func (recv *StatusIconClass) ToC() unsafe.Pointer {
	return recv.native
}

// StatusIconClassNewFromC creates a new StatusIconClass from a pointer to the C GtkStatusIconClass that represents the StatusIconClass.
func StatusIconClassNewFromC(native unsafe.Pointer) *StatusIconClass {
	return &StatusIconClass{native: native}
}

// StatusIconPrivate is a representation of the C record GtkStatusIconPrivate.
type StatusIconPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkStatusIconPrivate that represents the StatusIconPrivate.
func (recv *StatusIconPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// StatusIconPrivateNewFromC creates a new StatusIconPrivate from a pointer to the C GtkStatusIconPrivate that represents the StatusIconPrivate.
func StatusIconPrivateNewFromC(native unsafe.Pointer) *StatusIconPrivate {
	return &StatusIconPrivate{native: native}
}

// StatusbarAccessibleClass is a representation of the C record GtkStatusbarAccessibleClass.
type StatusbarAccessibleClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkStatusbarAccessibleClass that represents the StatusbarAccessibleClass.
func (recv *StatusbarAccessibleClass) ToC() unsafe.Pointer {
	return recv.native
}

// StatusbarAccessibleClassNewFromC creates a new StatusbarAccessibleClass from a pointer to the C GtkStatusbarAccessibleClass that represents the StatusbarAccessibleClass.
func StatusbarAccessibleClassNewFromC(native unsafe.Pointer) *StatusbarAccessibleClass {
	return &StatusbarAccessibleClass{native: native}
}

// StatusbarAccessiblePrivate is a representation of the C record GtkStatusbarAccessiblePrivate.
type StatusbarAccessiblePrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkStatusbarAccessiblePrivate that represents the StatusbarAccessiblePrivate.
func (recv *StatusbarAccessiblePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// StatusbarAccessiblePrivateNewFromC creates a new StatusbarAccessiblePrivate from a pointer to the C GtkStatusbarAccessiblePrivate that represents the StatusbarAccessiblePrivate.
func StatusbarAccessiblePrivateNewFromC(native unsafe.Pointer) *StatusbarAccessiblePrivate {
	return &StatusbarAccessiblePrivate{native: native}
}

// StatusbarClass is a representation of the C record GtkStatusbarClass.
type StatusbarClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkStatusbarClass that represents the StatusbarClass.
func (recv *StatusbarClass) ToC() unsafe.Pointer {
	return recv.native
}

// StatusbarClassNewFromC creates a new StatusbarClass from a pointer to the C GtkStatusbarClass that represents the StatusbarClass.
func StatusbarClassNewFromC(native unsafe.Pointer) *StatusbarClass {
	return &StatusbarClass{native: native}
}

// StatusbarPrivate is a representation of the C record GtkStatusbarPrivate.
type StatusbarPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkStatusbarPrivate that represents the StatusbarPrivate.
func (recv *StatusbarPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// StatusbarPrivateNewFromC creates a new StatusbarPrivate from a pointer to the C GtkStatusbarPrivate that represents the StatusbarPrivate.
func StatusbarPrivateNewFromC(native unsafe.Pointer) *StatusbarPrivate {
	return &StatusbarPrivate{native: native}
}

// StockItem is a representation of the C record GtkStockItem.
type StockItem struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkStockItem that represents the StockItem.
func (recv *StockItem) ToC() unsafe.Pointer {
	return recv.native
}

// StockItemNewFromC creates a new StockItem from a pointer to the C GtkStockItem that represents the StockItem.
func StockItemNewFromC(native unsafe.Pointer) *StockItem {
	return &StockItem{native: native}
}

// StyleClass is a representation of the C record GtkStyleClass.
type StyleClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkStyleClass that represents the StyleClass.
func (recv *StyleClass) ToC() unsafe.Pointer {
	return recv.native
}

// StyleClassNewFromC creates a new StyleClass from a pointer to the C GtkStyleClass that represents the StyleClass.
func StyleClassNewFromC(native unsafe.Pointer) *StyleClass {
	return &StyleClass{native: native}
}

// StyleContextClass is a representation of the C record GtkStyleContextClass.
type StyleContextClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkStyleContextClass that represents the StyleContextClass.
func (recv *StyleContextClass) ToC() unsafe.Pointer {
	return recv.native
}

// StyleContextClassNewFromC creates a new StyleContextClass from a pointer to the C GtkStyleContextClass that represents the StyleContextClass.
func StyleContextClassNewFromC(native unsafe.Pointer) *StyleContextClass {
	return &StyleContextClass{native: native}
}

// StyleContextPrivate is a representation of the C record GtkStyleContextPrivate.
type StyleContextPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkStyleContextPrivate that represents the StyleContextPrivate.
func (recv *StyleContextPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// StyleContextPrivateNewFromC creates a new StyleContextPrivate from a pointer to the C GtkStyleContextPrivate that represents the StyleContextPrivate.
func StyleContextPrivateNewFromC(native unsafe.Pointer) *StyleContextPrivate {
	return &StyleContextPrivate{native: native}
}

// StylePropertiesClass is a representation of the C record GtkStylePropertiesClass.
type StylePropertiesClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkStylePropertiesClass that represents the StylePropertiesClass.
func (recv *StylePropertiesClass) ToC() unsafe.Pointer {
	return recv.native
}

// StylePropertiesClassNewFromC creates a new StylePropertiesClass from a pointer to the C GtkStylePropertiesClass that represents the StylePropertiesClass.
func StylePropertiesClassNewFromC(native unsafe.Pointer) *StylePropertiesClass {
	return &StylePropertiesClass{native: native}
}

// StylePropertiesPrivate is a representation of the C record GtkStylePropertiesPrivate.
type StylePropertiesPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkStylePropertiesPrivate that represents the StylePropertiesPrivate.
func (recv *StylePropertiesPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// StylePropertiesPrivateNewFromC creates a new StylePropertiesPrivate from a pointer to the C GtkStylePropertiesPrivate that represents the StylePropertiesPrivate.
func StylePropertiesPrivateNewFromC(native unsafe.Pointer) *StylePropertiesPrivate {
	return &StylePropertiesPrivate{native: native}
}

// StyleProviderIface is a representation of the C record GtkStyleProviderIface.
type StyleProviderIface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkStyleProviderIface that represents the StyleProviderIface.
func (recv *StyleProviderIface) ToC() unsafe.Pointer {
	return recv.native
}

// StyleProviderIfaceNewFromC creates a new StyleProviderIface from a pointer to the C GtkStyleProviderIface that represents the StyleProviderIface.
func StyleProviderIfaceNewFromC(native unsafe.Pointer) *StyleProviderIface {
	return &StyleProviderIface{native: native}
}

// SwitchAccessibleClass is a representation of the C record GtkSwitchAccessibleClass.
type SwitchAccessibleClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkSwitchAccessibleClass that represents the SwitchAccessibleClass.
func (recv *SwitchAccessibleClass) ToC() unsafe.Pointer {
	return recv.native
}

// SwitchAccessibleClassNewFromC creates a new SwitchAccessibleClass from a pointer to the C GtkSwitchAccessibleClass that represents the SwitchAccessibleClass.
func SwitchAccessibleClassNewFromC(native unsafe.Pointer) *SwitchAccessibleClass {
	return &SwitchAccessibleClass{native: native}
}

// SwitchAccessiblePrivate is a representation of the C record GtkSwitchAccessiblePrivate.
type SwitchAccessiblePrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkSwitchAccessiblePrivate that represents the SwitchAccessiblePrivate.
func (recv *SwitchAccessiblePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// SwitchAccessiblePrivateNewFromC creates a new SwitchAccessiblePrivate from a pointer to the C GtkSwitchAccessiblePrivate that represents the SwitchAccessiblePrivate.
func SwitchAccessiblePrivateNewFromC(native unsafe.Pointer) *SwitchAccessiblePrivate {
	return &SwitchAccessiblePrivate{native: native}
}

// SwitchClass is a representation of the C record GtkSwitchClass.
type SwitchClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkSwitchClass that represents the SwitchClass.
func (recv *SwitchClass) ToC() unsafe.Pointer {
	return recv.native
}

// SwitchClassNewFromC creates a new SwitchClass from a pointer to the C GtkSwitchClass that represents the SwitchClass.
func SwitchClassNewFromC(native unsafe.Pointer) *SwitchClass {
	return &SwitchClass{native: native}
}

// SwitchPrivate is a representation of the C record GtkSwitchPrivate.
type SwitchPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkSwitchPrivate that represents the SwitchPrivate.
func (recv *SwitchPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// SwitchPrivateNewFromC creates a new SwitchPrivate from a pointer to the C GtkSwitchPrivate that represents the SwitchPrivate.
func SwitchPrivateNewFromC(native unsafe.Pointer) *SwitchPrivate {
	return &SwitchPrivate{native: native}
}

// SymbolicColor is a representation of the C record GtkSymbolicColor.
type SymbolicColor struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkSymbolicColor that represents the SymbolicColor.
func (recv *SymbolicColor) ToC() unsafe.Pointer {
	return recv.native
}

// SymbolicColorNewFromC creates a new SymbolicColor from a pointer to the C GtkSymbolicColor that represents the SymbolicColor.
func SymbolicColorNewFromC(native unsafe.Pointer) *SymbolicColor {
	return &SymbolicColor{native: native}
}

// TableChild is a representation of the C record GtkTableChild.
type TableChild struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkTableChild that represents the TableChild.
func (recv *TableChild) ToC() unsafe.Pointer {
	return recv.native
}

// TableChildNewFromC creates a new TableChild from a pointer to the C GtkTableChild that represents the TableChild.
func TableChildNewFromC(native unsafe.Pointer) *TableChild {
	return &TableChild{native: native}
}

// TableClass is a representation of the C record GtkTableClass.
type TableClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkTableClass that represents the TableClass.
func (recv *TableClass) ToC() unsafe.Pointer {
	return recv.native
}

// TableClassNewFromC creates a new TableClass from a pointer to the C GtkTableClass that represents the TableClass.
func TableClassNewFromC(native unsafe.Pointer) *TableClass {
	return &TableClass{native: native}
}

// TablePrivate is a representation of the C record GtkTablePrivate.
type TablePrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkTablePrivate that represents the TablePrivate.
func (recv *TablePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// TablePrivateNewFromC creates a new TablePrivate from a pointer to the C GtkTablePrivate that represents the TablePrivate.
func TablePrivateNewFromC(native unsafe.Pointer) *TablePrivate {
	return &TablePrivate{native: native}
}

// TableRowCol is a representation of the C record GtkTableRowCol.
type TableRowCol struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkTableRowCol that represents the TableRowCol.
func (recv *TableRowCol) ToC() unsafe.Pointer {
	return recv.native
}

// TableRowColNewFromC creates a new TableRowCol from a pointer to the C GtkTableRowCol that represents the TableRowCol.
func TableRowColNewFromC(native unsafe.Pointer) *TableRowCol {
	return &TableRowCol{native: native}
}

// TargetEntry is a representation of the C record GtkTargetEntry.
type TargetEntry struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkTargetEntry that represents the TargetEntry.
func (recv *TargetEntry) ToC() unsafe.Pointer {
	return recv.native
}

// TargetEntryNewFromC creates a new TargetEntry from a pointer to the C GtkTargetEntry that represents the TargetEntry.
func TargetEntryNewFromC(native unsafe.Pointer) *TargetEntry {
	return &TargetEntry{native: native}
}

// TargetList is a representation of the C record GtkTargetList.
type TargetList struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkTargetList that represents the TargetList.
func (recv *TargetList) ToC() unsafe.Pointer {
	return recv.native
}

// TargetListNewFromC creates a new TargetList from a pointer to the C GtkTargetList that represents the TargetList.
func TargetListNewFromC(native unsafe.Pointer) *TargetList {
	return &TargetList{native: native}
}

// TargetPair is a representation of the C record GtkTargetPair.
type TargetPair struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkTargetPair that represents the TargetPair.
func (recv *TargetPair) ToC() unsafe.Pointer {
	return recv.native
}

// TargetPairNewFromC creates a new TargetPair from a pointer to the C GtkTargetPair that represents the TargetPair.
func TargetPairNewFromC(native unsafe.Pointer) *TargetPair {
	return &TargetPair{native: native}
}

// TearoffMenuItemClass is a representation of the C record GtkTearoffMenuItemClass.
type TearoffMenuItemClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkTearoffMenuItemClass that represents the TearoffMenuItemClass.
func (recv *TearoffMenuItemClass) ToC() unsafe.Pointer {
	return recv.native
}

// TearoffMenuItemClassNewFromC creates a new TearoffMenuItemClass from a pointer to the C GtkTearoffMenuItemClass that represents the TearoffMenuItemClass.
func TearoffMenuItemClassNewFromC(native unsafe.Pointer) *TearoffMenuItemClass {
	return &TearoffMenuItemClass{native: native}
}

// TearoffMenuItemPrivate is a representation of the C record GtkTearoffMenuItemPrivate.
type TearoffMenuItemPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkTearoffMenuItemPrivate that represents the TearoffMenuItemPrivate.
func (recv *TearoffMenuItemPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// TearoffMenuItemPrivateNewFromC creates a new TearoffMenuItemPrivate from a pointer to the C GtkTearoffMenuItemPrivate that represents the TearoffMenuItemPrivate.
func TearoffMenuItemPrivateNewFromC(native unsafe.Pointer) *TearoffMenuItemPrivate {
	return &TearoffMenuItemPrivate{native: native}
}

// TextAppearance is a representation of the C record GtkTextAppearance.
type TextAppearance struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkTextAppearance that represents the TextAppearance.
func (recv *TextAppearance) ToC() unsafe.Pointer {
	return recv.native
}

// TextAppearanceNewFromC creates a new TextAppearance from a pointer to the C GtkTextAppearance that represents the TextAppearance.
func TextAppearanceNewFromC(native unsafe.Pointer) *TextAppearance {
	return &TextAppearance{native: native}
}

// TextAttributes is a representation of the C record GtkTextAttributes.
type TextAttributes struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkTextAttributes that represents the TextAttributes.
func (recv *TextAttributes) ToC() unsafe.Pointer {
	return recv.native
}

// TextAttributesNewFromC creates a new TextAttributes from a pointer to the C GtkTextAttributes that represents the TextAttributes.
func TextAttributesNewFromC(native unsafe.Pointer) *TextAttributes {
	return &TextAttributes{native: native}
}

// TextBTree is a representation of the C record GtkTextBTree.
type TextBTree struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkTextBTree that represents the TextBTree.
func (recv *TextBTree) ToC() unsafe.Pointer {
	return recv.native
}

// TextBTreeNewFromC creates a new TextBTree from a pointer to the C GtkTextBTree that represents the TextBTree.
func TextBTreeNewFromC(native unsafe.Pointer) *TextBTree {
	return &TextBTree{native: native}
}

// TextBufferClass is a representation of the C record GtkTextBufferClass.
type TextBufferClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkTextBufferClass that represents the TextBufferClass.
func (recv *TextBufferClass) ToC() unsafe.Pointer {
	return recv.native
}

// TextBufferClassNewFromC creates a new TextBufferClass from a pointer to the C GtkTextBufferClass that represents the TextBufferClass.
func TextBufferClassNewFromC(native unsafe.Pointer) *TextBufferClass {
	return &TextBufferClass{native: native}
}

// TextBufferPrivate is a representation of the C record GtkTextBufferPrivate.
type TextBufferPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkTextBufferPrivate that represents the TextBufferPrivate.
func (recv *TextBufferPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// TextBufferPrivateNewFromC creates a new TextBufferPrivate from a pointer to the C GtkTextBufferPrivate that represents the TextBufferPrivate.
func TextBufferPrivateNewFromC(native unsafe.Pointer) *TextBufferPrivate {
	return &TextBufferPrivate{native: native}
}

// TextCellAccessibleClass is a representation of the C record GtkTextCellAccessibleClass.
type TextCellAccessibleClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkTextCellAccessibleClass that represents the TextCellAccessibleClass.
func (recv *TextCellAccessibleClass) ToC() unsafe.Pointer {
	return recv.native
}

// TextCellAccessibleClassNewFromC creates a new TextCellAccessibleClass from a pointer to the C GtkTextCellAccessibleClass that represents the TextCellAccessibleClass.
func TextCellAccessibleClassNewFromC(native unsafe.Pointer) *TextCellAccessibleClass {
	return &TextCellAccessibleClass{native: native}
}

// TextCellAccessiblePrivate is a representation of the C record GtkTextCellAccessiblePrivate.
type TextCellAccessiblePrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkTextCellAccessiblePrivate that represents the TextCellAccessiblePrivate.
func (recv *TextCellAccessiblePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// TextCellAccessiblePrivateNewFromC creates a new TextCellAccessiblePrivate from a pointer to the C GtkTextCellAccessiblePrivate that represents the TextCellAccessiblePrivate.
func TextCellAccessiblePrivateNewFromC(native unsafe.Pointer) *TextCellAccessiblePrivate {
	return &TextCellAccessiblePrivate{native: native}
}

// TextChildAnchorClass is a representation of the C record GtkTextChildAnchorClass.
type TextChildAnchorClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkTextChildAnchorClass that represents the TextChildAnchorClass.
func (recv *TextChildAnchorClass) ToC() unsafe.Pointer {
	return recv.native
}

// TextChildAnchorClassNewFromC creates a new TextChildAnchorClass from a pointer to the C GtkTextChildAnchorClass that represents the TextChildAnchorClass.
func TextChildAnchorClassNewFromC(native unsafe.Pointer) *TextChildAnchorClass {
	return &TextChildAnchorClass{native: native}
}

// TextIter is a representation of the C record GtkTextIter.
type TextIter struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkTextIter that represents the TextIter.
func (recv *TextIter) ToC() unsafe.Pointer {
	return recv.native
}

// TextIterNewFromC creates a new TextIter from a pointer to the C GtkTextIter that represents the TextIter.
func TextIterNewFromC(native unsafe.Pointer) *TextIter {
	return &TextIter{native: native}
}

// TextMarkClass is a representation of the C record GtkTextMarkClass.
type TextMarkClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkTextMarkClass that represents the TextMarkClass.
func (recv *TextMarkClass) ToC() unsafe.Pointer {
	return recv.native
}

// TextMarkClassNewFromC creates a new TextMarkClass from a pointer to the C GtkTextMarkClass that represents the TextMarkClass.
func TextMarkClassNewFromC(native unsafe.Pointer) *TextMarkClass {
	return &TextMarkClass{native: native}
}

// TextTagClass is a representation of the C record GtkTextTagClass.
type TextTagClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkTextTagClass that represents the TextTagClass.
func (recv *TextTagClass) ToC() unsafe.Pointer {
	return recv.native
}

// TextTagClassNewFromC creates a new TextTagClass from a pointer to the C GtkTextTagClass that represents the TextTagClass.
func TextTagClassNewFromC(native unsafe.Pointer) *TextTagClass {
	return &TextTagClass{native: native}
}

// TextTagPrivate is a representation of the C record GtkTextTagPrivate.
type TextTagPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkTextTagPrivate that represents the TextTagPrivate.
func (recv *TextTagPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// TextTagPrivateNewFromC creates a new TextTagPrivate from a pointer to the C GtkTextTagPrivate that represents the TextTagPrivate.
func TextTagPrivateNewFromC(native unsafe.Pointer) *TextTagPrivate {
	return &TextTagPrivate{native: native}
}

// TextTagTableClass is a representation of the C record GtkTextTagTableClass.
type TextTagTableClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkTextTagTableClass that represents the TextTagTableClass.
func (recv *TextTagTableClass) ToC() unsafe.Pointer {
	return recv.native
}

// TextTagTableClassNewFromC creates a new TextTagTableClass from a pointer to the C GtkTextTagTableClass that represents the TextTagTableClass.
func TextTagTableClassNewFromC(native unsafe.Pointer) *TextTagTableClass {
	return &TextTagTableClass{native: native}
}

// TextTagTablePrivate is a representation of the C record GtkTextTagTablePrivate.
type TextTagTablePrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkTextTagTablePrivate that represents the TextTagTablePrivate.
func (recv *TextTagTablePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// TextTagTablePrivateNewFromC creates a new TextTagTablePrivate from a pointer to the C GtkTextTagTablePrivate that represents the TextTagTablePrivate.
func TextTagTablePrivateNewFromC(native unsafe.Pointer) *TextTagTablePrivate {
	return &TextTagTablePrivate{native: native}
}

// TextViewAccessibleClass is a representation of the C record GtkTextViewAccessibleClass.
type TextViewAccessibleClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkTextViewAccessibleClass that represents the TextViewAccessibleClass.
func (recv *TextViewAccessibleClass) ToC() unsafe.Pointer {
	return recv.native
}

// TextViewAccessibleClassNewFromC creates a new TextViewAccessibleClass from a pointer to the C GtkTextViewAccessibleClass that represents the TextViewAccessibleClass.
func TextViewAccessibleClassNewFromC(native unsafe.Pointer) *TextViewAccessibleClass {
	return &TextViewAccessibleClass{native: native}
}

// TextViewAccessiblePrivate is a representation of the C record GtkTextViewAccessiblePrivate.
type TextViewAccessiblePrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkTextViewAccessiblePrivate that represents the TextViewAccessiblePrivate.
func (recv *TextViewAccessiblePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// TextViewAccessiblePrivateNewFromC creates a new TextViewAccessiblePrivate from a pointer to the C GtkTextViewAccessiblePrivate that represents the TextViewAccessiblePrivate.
func TextViewAccessiblePrivateNewFromC(native unsafe.Pointer) *TextViewAccessiblePrivate {
	return &TextViewAccessiblePrivate{native: native}
}

// TextViewClass is a representation of the C record GtkTextViewClass.
type TextViewClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkTextViewClass that represents the TextViewClass.
func (recv *TextViewClass) ToC() unsafe.Pointer {
	return recv.native
}

// TextViewClassNewFromC creates a new TextViewClass from a pointer to the C GtkTextViewClass that represents the TextViewClass.
func TextViewClassNewFromC(native unsafe.Pointer) *TextViewClass {
	return &TextViewClass{native: native}
}

// TextViewPrivate is a representation of the C record GtkTextViewPrivate.
type TextViewPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkTextViewPrivate that represents the TextViewPrivate.
func (recv *TextViewPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// TextViewPrivateNewFromC creates a new TextViewPrivate from a pointer to the C GtkTextViewPrivate that represents the TextViewPrivate.
func TextViewPrivateNewFromC(native unsafe.Pointer) *TextViewPrivate {
	return &TextViewPrivate{native: native}
}

// ThemeEngine is a representation of the C record GtkThemeEngine.
type ThemeEngine struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkThemeEngine that represents the ThemeEngine.
func (recv *ThemeEngine) ToC() unsafe.Pointer {
	return recv.native
}

// ThemeEngineNewFromC creates a new ThemeEngine from a pointer to the C GtkThemeEngine that represents the ThemeEngine.
func ThemeEngineNewFromC(native unsafe.Pointer) *ThemeEngine {
	return &ThemeEngine{native: native}
}

// ThemingEngineClass is a representation of the C record GtkThemingEngineClass.
type ThemingEngineClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkThemingEngineClass that represents the ThemingEngineClass.
func (recv *ThemingEngineClass) ToC() unsafe.Pointer {
	return recv.native
}

// ThemingEngineClassNewFromC creates a new ThemingEngineClass from a pointer to the C GtkThemingEngineClass that represents the ThemingEngineClass.
func ThemingEngineClassNewFromC(native unsafe.Pointer) *ThemingEngineClass {
	return &ThemingEngineClass{native: native}
}

// ThemingEnginePrivate is a representation of the C record GtkThemingEnginePrivate.
type ThemingEnginePrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkThemingEnginePrivate that represents the ThemingEnginePrivate.
func (recv *ThemingEnginePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ThemingEnginePrivateNewFromC creates a new ThemingEnginePrivate from a pointer to the C GtkThemingEnginePrivate that represents the ThemingEnginePrivate.
func ThemingEnginePrivateNewFromC(native unsafe.Pointer) *ThemingEnginePrivate {
	return &ThemingEnginePrivate{native: native}
}

// ToggleActionClass is a representation of the C record GtkToggleActionClass.
type ToggleActionClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkToggleActionClass that represents the ToggleActionClass.
func (recv *ToggleActionClass) ToC() unsafe.Pointer {
	return recv.native
}

// ToggleActionClassNewFromC creates a new ToggleActionClass from a pointer to the C GtkToggleActionClass that represents the ToggleActionClass.
func ToggleActionClassNewFromC(native unsafe.Pointer) *ToggleActionClass {
	return &ToggleActionClass{native: native}
}

// ToggleActionEntry is a representation of the C record GtkToggleActionEntry.
type ToggleActionEntry struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkToggleActionEntry that represents the ToggleActionEntry.
func (recv *ToggleActionEntry) ToC() unsafe.Pointer {
	return recv.native
}

// ToggleActionEntryNewFromC creates a new ToggleActionEntry from a pointer to the C GtkToggleActionEntry that represents the ToggleActionEntry.
func ToggleActionEntryNewFromC(native unsafe.Pointer) *ToggleActionEntry {
	return &ToggleActionEntry{native: native}
}

// ToggleActionPrivate is a representation of the C record GtkToggleActionPrivate.
type ToggleActionPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkToggleActionPrivate that represents the ToggleActionPrivate.
func (recv *ToggleActionPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ToggleActionPrivateNewFromC creates a new ToggleActionPrivate from a pointer to the C GtkToggleActionPrivate that represents the ToggleActionPrivate.
func ToggleActionPrivateNewFromC(native unsafe.Pointer) *ToggleActionPrivate {
	return &ToggleActionPrivate{native: native}
}

// ToggleButtonAccessibleClass is a representation of the C record GtkToggleButtonAccessibleClass.
type ToggleButtonAccessibleClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkToggleButtonAccessibleClass that represents the ToggleButtonAccessibleClass.
func (recv *ToggleButtonAccessibleClass) ToC() unsafe.Pointer {
	return recv.native
}

// ToggleButtonAccessibleClassNewFromC creates a new ToggleButtonAccessibleClass from a pointer to the C GtkToggleButtonAccessibleClass that represents the ToggleButtonAccessibleClass.
func ToggleButtonAccessibleClassNewFromC(native unsafe.Pointer) *ToggleButtonAccessibleClass {
	return &ToggleButtonAccessibleClass{native: native}
}

// ToggleButtonAccessiblePrivate is a representation of the C record GtkToggleButtonAccessiblePrivate.
type ToggleButtonAccessiblePrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkToggleButtonAccessiblePrivate that represents the ToggleButtonAccessiblePrivate.
func (recv *ToggleButtonAccessiblePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ToggleButtonAccessiblePrivateNewFromC creates a new ToggleButtonAccessiblePrivate from a pointer to the C GtkToggleButtonAccessiblePrivate that represents the ToggleButtonAccessiblePrivate.
func ToggleButtonAccessiblePrivateNewFromC(native unsafe.Pointer) *ToggleButtonAccessiblePrivate {
	return &ToggleButtonAccessiblePrivate{native: native}
}

// ToggleButtonClass is a representation of the C record GtkToggleButtonClass.
type ToggleButtonClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkToggleButtonClass that represents the ToggleButtonClass.
func (recv *ToggleButtonClass) ToC() unsafe.Pointer {
	return recv.native
}

// ToggleButtonClassNewFromC creates a new ToggleButtonClass from a pointer to the C GtkToggleButtonClass that represents the ToggleButtonClass.
func ToggleButtonClassNewFromC(native unsafe.Pointer) *ToggleButtonClass {
	return &ToggleButtonClass{native: native}
}

// ToggleButtonPrivate is a representation of the C record GtkToggleButtonPrivate.
type ToggleButtonPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkToggleButtonPrivate that represents the ToggleButtonPrivate.
func (recv *ToggleButtonPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ToggleButtonPrivateNewFromC creates a new ToggleButtonPrivate from a pointer to the C GtkToggleButtonPrivate that represents the ToggleButtonPrivate.
func ToggleButtonPrivateNewFromC(native unsafe.Pointer) *ToggleButtonPrivate {
	return &ToggleButtonPrivate{native: native}
}

// ToggleToolButtonClass is a representation of the C record GtkToggleToolButtonClass.
type ToggleToolButtonClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkToggleToolButtonClass that represents the ToggleToolButtonClass.
func (recv *ToggleToolButtonClass) ToC() unsafe.Pointer {
	return recv.native
}

// ToggleToolButtonClassNewFromC creates a new ToggleToolButtonClass from a pointer to the C GtkToggleToolButtonClass that represents the ToggleToolButtonClass.
func ToggleToolButtonClassNewFromC(native unsafe.Pointer) *ToggleToolButtonClass {
	return &ToggleToolButtonClass{native: native}
}

// ToggleToolButtonPrivate is a representation of the C record GtkToggleToolButtonPrivate.
type ToggleToolButtonPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkToggleToolButtonPrivate that represents the ToggleToolButtonPrivate.
func (recv *ToggleToolButtonPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ToggleToolButtonPrivateNewFromC creates a new ToggleToolButtonPrivate from a pointer to the C GtkToggleToolButtonPrivate that represents the ToggleToolButtonPrivate.
func ToggleToolButtonPrivateNewFromC(native unsafe.Pointer) *ToggleToolButtonPrivate {
	return &ToggleToolButtonPrivate{native: native}
}

// ToolButtonClass is a representation of the C record GtkToolButtonClass.
type ToolButtonClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkToolButtonClass that represents the ToolButtonClass.
func (recv *ToolButtonClass) ToC() unsafe.Pointer {
	return recv.native
}

// ToolButtonClassNewFromC creates a new ToolButtonClass from a pointer to the C GtkToolButtonClass that represents the ToolButtonClass.
func ToolButtonClassNewFromC(native unsafe.Pointer) *ToolButtonClass {
	return &ToolButtonClass{native: native}
}

// ToolButtonPrivate is a representation of the C record GtkToolButtonPrivate.
type ToolButtonPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkToolButtonPrivate that represents the ToolButtonPrivate.
func (recv *ToolButtonPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ToolButtonPrivateNewFromC creates a new ToolButtonPrivate from a pointer to the C GtkToolButtonPrivate that represents the ToolButtonPrivate.
func ToolButtonPrivateNewFromC(native unsafe.Pointer) *ToolButtonPrivate {
	return &ToolButtonPrivate{native: native}
}

// ToolItemClass is a representation of the C record GtkToolItemClass.
type ToolItemClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkToolItemClass that represents the ToolItemClass.
func (recv *ToolItemClass) ToC() unsafe.Pointer {
	return recv.native
}

// ToolItemClassNewFromC creates a new ToolItemClass from a pointer to the C GtkToolItemClass that represents the ToolItemClass.
func ToolItemClassNewFromC(native unsafe.Pointer) *ToolItemClass {
	return &ToolItemClass{native: native}
}

// ToolItemGroupClass is a representation of the C record GtkToolItemGroupClass.
type ToolItemGroupClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkToolItemGroupClass that represents the ToolItemGroupClass.
func (recv *ToolItemGroupClass) ToC() unsafe.Pointer {
	return recv.native
}

// ToolItemGroupClassNewFromC creates a new ToolItemGroupClass from a pointer to the C GtkToolItemGroupClass that represents the ToolItemGroupClass.
func ToolItemGroupClassNewFromC(native unsafe.Pointer) *ToolItemGroupClass {
	return &ToolItemGroupClass{native: native}
}

// ToolItemGroupPrivate is a representation of the C record GtkToolItemGroupPrivate.
type ToolItemGroupPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkToolItemGroupPrivate that represents the ToolItemGroupPrivate.
func (recv *ToolItemGroupPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ToolItemGroupPrivateNewFromC creates a new ToolItemGroupPrivate from a pointer to the C GtkToolItemGroupPrivate that represents the ToolItemGroupPrivate.
func ToolItemGroupPrivateNewFromC(native unsafe.Pointer) *ToolItemGroupPrivate {
	return &ToolItemGroupPrivate{native: native}
}

// ToolItemPrivate is a representation of the C record GtkToolItemPrivate.
type ToolItemPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkToolItemPrivate that represents the ToolItemPrivate.
func (recv *ToolItemPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ToolItemPrivateNewFromC creates a new ToolItemPrivate from a pointer to the C GtkToolItemPrivate that represents the ToolItemPrivate.
func ToolItemPrivateNewFromC(native unsafe.Pointer) *ToolItemPrivate {
	return &ToolItemPrivate{native: native}
}

// ToolPaletteClass is a representation of the C record GtkToolPaletteClass.
type ToolPaletteClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkToolPaletteClass that represents the ToolPaletteClass.
func (recv *ToolPaletteClass) ToC() unsafe.Pointer {
	return recv.native
}

// ToolPaletteClassNewFromC creates a new ToolPaletteClass from a pointer to the C GtkToolPaletteClass that represents the ToolPaletteClass.
func ToolPaletteClassNewFromC(native unsafe.Pointer) *ToolPaletteClass {
	return &ToolPaletteClass{native: native}
}

// ToolPalettePrivate is a representation of the C record GtkToolPalettePrivate.
type ToolPalettePrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkToolPalettePrivate that represents the ToolPalettePrivate.
func (recv *ToolPalettePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ToolPalettePrivateNewFromC creates a new ToolPalettePrivate from a pointer to the C GtkToolPalettePrivate that represents the ToolPalettePrivate.
func ToolPalettePrivateNewFromC(native unsafe.Pointer) *ToolPalettePrivate {
	return &ToolPalettePrivate{native: native}
}

// ToolShellIface is a representation of the C record GtkToolShellIface.
type ToolShellIface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkToolShellIface that represents the ToolShellIface.
func (recv *ToolShellIface) ToC() unsafe.Pointer {
	return recv.native
}

// ToolShellIfaceNewFromC creates a new ToolShellIface from a pointer to the C GtkToolShellIface that represents the ToolShellIface.
func ToolShellIfaceNewFromC(native unsafe.Pointer) *ToolShellIface {
	return &ToolShellIface{native: native}
}

// ToolbarClass is a representation of the C record GtkToolbarClass.
type ToolbarClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkToolbarClass that represents the ToolbarClass.
func (recv *ToolbarClass) ToC() unsafe.Pointer {
	return recv.native
}

// ToolbarClassNewFromC creates a new ToolbarClass from a pointer to the C GtkToolbarClass that represents the ToolbarClass.
func ToolbarClassNewFromC(native unsafe.Pointer) *ToolbarClass {
	return &ToolbarClass{native: native}
}

// ToolbarPrivate is a representation of the C record GtkToolbarPrivate.
type ToolbarPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkToolbarPrivate that represents the ToolbarPrivate.
func (recv *ToolbarPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ToolbarPrivateNewFromC creates a new ToolbarPrivate from a pointer to the C GtkToolbarPrivate that represents the ToolbarPrivate.
func ToolbarPrivateNewFromC(native unsafe.Pointer) *ToolbarPrivate {
	return &ToolbarPrivate{native: native}
}

// ToplevelAccessibleClass is a representation of the C record GtkToplevelAccessibleClass.
type ToplevelAccessibleClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkToplevelAccessibleClass that represents the ToplevelAccessibleClass.
func (recv *ToplevelAccessibleClass) ToC() unsafe.Pointer {
	return recv.native
}

// ToplevelAccessibleClassNewFromC creates a new ToplevelAccessibleClass from a pointer to the C GtkToplevelAccessibleClass that represents the ToplevelAccessibleClass.
func ToplevelAccessibleClassNewFromC(native unsafe.Pointer) *ToplevelAccessibleClass {
	return &ToplevelAccessibleClass{native: native}
}

// ToplevelAccessiblePrivate is a representation of the C record GtkToplevelAccessiblePrivate.
type ToplevelAccessiblePrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkToplevelAccessiblePrivate that represents the ToplevelAccessiblePrivate.
func (recv *ToplevelAccessiblePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ToplevelAccessiblePrivateNewFromC creates a new ToplevelAccessiblePrivate from a pointer to the C GtkToplevelAccessiblePrivate that represents the ToplevelAccessiblePrivate.
func ToplevelAccessiblePrivateNewFromC(native unsafe.Pointer) *ToplevelAccessiblePrivate {
	return &ToplevelAccessiblePrivate{native: native}
}

// TreeDragDestIface is a representation of the C record GtkTreeDragDestIface.
type TreeDragDestIface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkTreeDragDestIface that represents the TreeDragDestIface.
func (recv *TreeDragDestIface) ToC() unsafe.Pointer {
	return recv.native
}

// TreeDragDestIfaceNewFromC creates a new TreeDragDestIface from a pointer to the C GtkTreeDragDestIface that represents the TreeDragDestIface.
func TreeDragDestIfaceNewFromC(native unsafe.Pointer) *TreeDragDestIface {
	return &TreeDragDestIface{native: native}
}

// TreeDragSourceIface is a representation of the C record GtkTreeDragSourceIface.
type TreeDragSourceIface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkTreeDragSourceIface that represents the TreeDragSourceIface.
func (recv *TreeDragSourceIface) ToC() unsafe.Pointer {
	return recv.native
}

// TreeDragSourceIfaceNewFromC creates a new TreeDragSourceIface from a pointer to the C GtkTreeDragSourceIface that represents the TreeDragSourceIface.
func TreeDragSourceIfaceNewFromC(native unsafe.Pointer) *TreeDragSourceIface {
	return &TreeDragSourceIface{native: native}
}

// TreeIter is a representation of the C record GtkTreeIter.
type TreeIter struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkTreeIter that represents the TreeIter.
func (recv *TreeIter) ToC() unsafe.Pointer {
	return recv.native
}

// TreeIterNewFromC creates a new TreeIter from a pointer to the C GtkTreeIter that represents the TreeIter.
func TreeIterNewFromC(native unsafe.Pointer) *TreeIter {
	return &TreeIter{native: native}
}

// TreeModelFilterClass is a representation of the C record GtkTreeModelFilterClass.
type TreeModelFilterClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkTreeModelFilterClass that represents the TreeModelFilterClass.
func (recv *TreeModelFilterClass) ToC() unsafe.Pointer {
	return recv.native
}

// TreeModelFilterClassNewFromC creates a new TreeModelFilterClass from a pointer to the C GtkTreeModelFilterClass that represents the TreeModelFilterClass.
func TreeModelFilterClassNewFromC(native unsafe.Pointer) *TreeModelFilterClass {
	return &TreeModelFilterClass{native: native}
}

// TreeModelFilterPrivate is a representation of the C record GtkTreeModelFilterPrivate.
type TreeModelFilterPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkTreeModelFilterPrivate that represents the TreeModelFilterPrivate.
func (recv *TreeModelFilterPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// TreeModelFilterPrivateNewFromC creates a new TreeModelFilterPrivate from a pointer to the C GtkTreeModelFilterPrivate that represents the TreeModelFilterPrivate.
func TreeModelFilterPrivateNewFromC(native unsafe.Pointer) *TreeModelFilterPrivate {
	return &TreeModelFilterPrivate{native: native}
}

// TreeModelIface is a representation of the C record GtkTreeModelIface.
type TreeModelIface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkTreeModelIface that represents the TreeModelIface.
func (recv *TreeModelIface) ToC() unsafe.Pointer {
	return recv.native
}

// TreeModelIfaceNewFromC creates a new TreeModelIface from a pointer to the C GtkTreeModelIface that represents the TreeModelIface.
func TreeModelIfaceNewFromC(native unsafe.Pointer) *TreeModelIface {
	return &TreeModelIface{native: native}
}

// TreeModelSortClass is a representation of the C record GtkTreeModelSortClass.
type TreeModelSortClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkTreeModelSortClass that represents the TreeModelSortClass.
func (recv *TreeModelSortClass) ToC() unsafe.Pointer {
	return recv.native
}

// TreeModelSortClassNewFromC creates a new TreeModelSortClass from a pointer to the C GtkTreeModelSortClass that represents the TreeModelSortClass.
func TreeModelSortClassNewFromC(native unsafe.Pointer) *TreeModelSortClass {
	return &TreeModelSortClass{native: native}
}

// TreeModelSortPrivate is a representation of the C record GtkTreeModelSortPrivate.
type TreeModelSortPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkTreeModelSortPrivate that represents the TreeModelSortPrivate.
func (recv *TreeModelSortPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// TreeModelSortPrivateNewFromC creates a new TreeModelSortPrivate from a pointer to the C GtkTreeModelSortPrivate that represents the TreeModelSortPrivate.
func TreeModelSortPrivateNewFromC(native unsafe.Pointer) *TreeModelSortPrivate {
	return &TreeModelSortPrivate{native: native}
}

// TreePath is a representation of the C record GtkTreePath.
type TreePath struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkTreePath that represents the TreePath.
func (recv *TreePath) ToC() unsafe.Pointer {
	return recv.native
}

// TreePathNewFromC creates a new TreePath from a pointer to the C GtkTreePath that represents the TreePath.
func TreePathNewFromC(native unsafe.Pointer) *TreePath {
	return &TreePath{native: native}
}

// TreeRowReference is a representation of the C record GtkTreeRowReference.
type TreeRowReference struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkTreeRowReference that represents the TreeRowReference.
func (recv *TreeRowReference) ToC() unsafe.Pointer {
	return recv.native
}

// TreeRowReferenceNewFromC creates a new TreeRowReference from a pointer to the C GtkTreeRowReference that represents the TreeRowReference.
func TreeRowReferenceNewFromC(native unsafe.Pointer) *TreeRowReference {
	return &TreeRowReference{native: native}
}

// TreeSelectionClass is a representation of the C record GtkTreeSelectionClass.
type TreeSelectionClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkTreeSelectionClass that represents the TreeSelectionClass.
func (recv *TreeSelectionClass) ToC() unsafe.Pointer {
	return recv.native
}

// TreeSelectionClassNewFromC creates a new TreeSelectionClass from a pointer to the C GtkTreeSelectionClass that represents the TreeSelectionClass.
func TreeSelectionClassNewFromC(native unsafe.Pointer) *TreeSelectionClass {
	return &TreeSelectionClass{native: native}
}

// TreeSelectionPrivate is a representation of the C record GtkTreeSelectionPrivate.
type TreeSelectionPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkTreeSelectionPrivate that represents the TreeSelectionPrivate.
func (recv *TreeSelectionPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// TreeSelectionPrivateNewFromC creates a new TreeSelectionPrivate from a pointer to the C GtkTreeSelectionPrivate that represents the TreeSelectionPrivate.
func TreeSelectionPrivateNewFromC(native unsafe.Pointer) *TreeSelectionPrivate {
	return &TreeSelectionPrivate{native: native}
}

// TreeSortableIface is a representation of the C record GtkTreeSortableIface.
type TreeSortableIface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkTreeSortableIface that represents the TreeSortableIface.
func (recv *TreeSortableIface) ToC() unsafe.Pointer {
	return recv.native
}

// TreeSortableIfaceNewFromC creates a new TreeSortableIface from a pointer to the C GtkTreeSortableIface that represents the TreeSortableIface.
func TreeSortableIfaceNewFromC(native unsafe.Pointer) *TreeSortableIface {
	return &TreeSortableIface{native: native}
}

// TreeStoreClass is a representation of the C record GtkTreeStoreClass.
type TreeStoreClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkTreeStoreClass that represents the TreeStoreClass.
func (recv *TreeStoreClass) ToC() unsafe.Pointer {
	return recv.native
}

// TreeStoreClassNewFromC creates a new TreeStoreClass from a pointer to the C GtkTreeStoreClass that represents the TreeStoreClass.
func TreeStoreClassNewFromC(native unsafe.Pointer) *TreeStoreClass {
	return &TreeStoreClass{native: native}
}

// TreeStorePrivate is a representation of the C record GtkTreeStorePrivate.
type TreeStorePrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkTreeStorePrivate that represents the TreeStorePrivate.
func (recv *TreeStorePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// TreeStorePrivateNewFromC creates a new TreeStorePrivate from a pointer to the C GtkTreeStorePrivate that represents the TreeStorePrivate.
func TreeStorePrivateNewFromC(native unsafe.Pointer) *TreeStorePrivate {
	return &TreeStorePrivate{native: native}
}

// TreeViewAccessibleClass is a representation of the C record GtkTreeViewAccessibleClass.
type TreeViewAccessibleClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkTreeViewAccessibleClass that represents the TreeViewAccessibleClass.
func (recv *TreeViewAccessibleClass) ToC() unsafe.Pointer {
	return recv.native
}

// TreeViewAccessibleClassNewFromC creates a new TreeViewAccessibleClass from a pointer to the C GtkTreeViewAccessibleClass that represents the TreeViewAccessibleClass.
func TreeViewAccessibleClassNewFromC(native unsafe.Pointer) *TreeViewAccessibleClass {
	return &TreeViewAccessibleClass{native: native}
}

// TreeViewAccessiblePrivate is a representation of the C record GtkTreeViewAccessiblePrivate.
type TreeViewAccessiblePrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkTreeViewAccessiblePrivate that represents the TreeViewAccessiblePrivate.
func (recv *TreeViewAccessiblePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// TreeViewAccessiblePrivateNewFromC creates a new TreeViewAccessiblePrivate from a pointer to the C GtkTreeViewAccessiblePrivate that represents the TreeViewAccessiblePrivate.
func TreeViewAccessiblePrivateNewFromC(native unsafe.Pointer) *TreeViewAccessiblePrivate {
	return &TreeViewAccessiblePrivate{native: native}
}

// TreeViewClass is a representation of the C record GtkTreeViewClass.
type TreeViewClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkTreeViewClass that represents the TreeViewClass.
func (recv *TreeViewClass) ToC() unsafe.Pointer {
	return recv.native
}

// TreeViewClassNewFromC creates a new TreeViewClass from a pointer to the C GtkTreeViewClass that represents the TreeViewClass.
func TreeViewClassNewFromC(native unsafe.Pointer) *TreeViewClass {
	return &TreeViewClass{native: native}
}

// TreeViewColumnClass is a representation of the C record GtkTreeViewColumnClass.
type TreeViewColumnClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkTreeViewColumnClass that represents the TreeViewColumnClass.
func (recv *TreeViewColumnClass) ToC() unsafe.Pointer {
	return recv.native
}

// TreeViewColumnClassNewFromC creates a new TreeViewColumnClass from a pointer to the C GtkTreeViewColumnClass that represents the TreeViewColumnClass.
func TreeViewColumnClassNewFromC(native unsafe.Pointer) *TreeViewColumnClass {
	return &TreeViewColumnClass{native: native}
}

// TreeViewColumnPrivate is a representation of the C record GtkTreeViewColumnPrivate.
type TreeViewColumnPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkTreeViewColumnPrivate that represents the TreeViewColumnPrivate.
func (recv *TreeViewColumnPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// TreeViewColumnPrivateNewFromC creates a new TreeViewColumnPrivate from a pointer to the C GtkTreeViewColumnPrivate that represents the TreeViewColumnPrivate.
func TreeViewColumnPrivateNewFromC(native unsafe.Pointer) *TreeViewColumnPrivate {
	return &TreeViewColumnPrivate{native: native}
}

// TreeViewPrivate is a representation of the C record GtkTreeViewPrivate.
type TreeViewPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkTreeViewPrivate that represents the TreeViewPrivate.
func (recv *TreeViewPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// TreeViewPrivateNewFromC creates a new TreeViewPrivate from a pointer to the C GtkTreeViewPrivate that represents the TreeViewPrivate.
func TreeViewPrivateNewFromC(native unsafe.Pointer) *TreeViewPrivate {
	return &TreeViewPrivate{native: native}
}

// UIManagerClass is a representation of the C record GtkUIManagerClass.
type UIManagerClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkUIManagerClass that represents the UIManagerClass.
func (recv *UIManagerClass) ToC() unsafe.Pointer {
	return recv.native
}

// UIManagerClassNewFromC creates a new UIManagerClass from a pointer to the C GtkUIManagerClass that represents the UIManagerClass.
func UIManagerClassNewFromC(native unsafe.Pointer) *UIManagerClass {
	return &UIManagerClass{native: native}
}

// UIManagerPrivate is a representation of the C record GtkUIManagerPrivate.
type UIManagerPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkUIManagerPrivate that represents the UIManagerPrivate.
func (recv *UIManagerPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// UIManagerPrivateNewFromC creates a new UIManagerPrivate from a pointer to the C GtkUIManagerPrivate that represents the UIManagerPrivate.
func UIManagerPrivateNewFromC(native unsafe.Pointer) *UIManagerPrivate {
	return &UIManagerPrivate{native: native}
}

// VBoxClass is a representation of the C record GtkVBoxClass.
type VBoxClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkVBoxClass that represents the VBoxClass.
func (recv *VBoxClass) ToC() unsafe.Pointer {
	return recv.native
}

// VBoxClassNewFromC creates a new VBoxClass from a pointer to the C GtkVBoxClass that represents the VBoxClass.
func VBoxClassNewFromC(native unsafe.Pointer) *VBoxClass {
	return &VBoxClass{native: native}
}

// VButtonBoxClass is a representation of the C record GtkVButtonBoxClass.
type VButtonBoxClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkVButtonBoxClass that represents the VButtonBoxClass.
func (recv *VButtonBoxClass) ToC() unsafe.Pointer {
	return recv.native
}

// VButtonBoxClassNewFromC creates a new VButtonBoxClass from a pointer to the C GtkVButtonBoxClass that represents the VButtonBoxClass.
func VButtonBoxClassNewFromC(native unsafe.Pointer) *VButtonBoxClass {
	return &VButtonBoxClass{native: native}
}

// VPanedClass is a representation of the C record GtkVPanedClass.
type VPanedClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkVPanedClass that represents the VPanedClass.
func (recv *VPanedClass) ToC() unsafe.Pointer {
	return recv.native
}

// VPanedClassNewFromC creates a new VPanedClass from a pointer to the C GtkVPanedClass that represents the VPanedClass.
func VPanedClassNewFromC(native unsafe.Pointer) *VPanedClass {
	return &VPanedClass{native: native}
}

// VScaleClass is a representation of the C record GtkVScaleClass.
type VScaleClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkVScaleClass that represents the VScaleClass.
func (recv *VScaleClass) ToC() unsafe.Pointer {
	return recv.native
}

// VScaleClassNewFromC creates a new VScaleClass from a pointer to the C GtkVScaleClass that represents the VScaleClass.
func VScaleClassNewFromC(native unsafe.Pointer) *VScaleClass {
	return &VScaleClass{native: native}
}

// VScrollbarClass is a representation of the C record GtkVScrollbarClass.
type VScrollbarClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkVScrollbarClass that represents the VScrollbarClass.
func (recv *VScrollbarClass) ToC() unsafe.Pointer {
	return recv.native
}

// VScrollbarClassNewFromC creates a new VScrollbarClass from a pointer to the C GtkVScrollbarClass that represents the VScrollbarClass.
func VScrollbarClassNewFromC(native unsafe.Pointer) *VScrollbarClass {
	return &VScrollbarClass{native: native}
}

// VSeparatorClass is a representation of the C record GtkVSeparatorClass.
type VSeparatorClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkVSeparatorClass that represents the VSeparatorClass.
func (recv *VSeparatorClass) ToC() unsafe.Pointer {
	return recv.native
}

// VSeparatorClassNewFromC creates a new VSeparatorClass from a pointer to the C GtkVSeparatorClass that represents the VSeparatorClass.
func VSeparatorClassNewFromC(native unsafe.Pointer) *VSeparatorClass {
	return &VSeparatorClass{native: native}
}

// ViewportClass is a representation of the C record GtkViewportClass.
type ViewportClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkViewportClass that represents the ViewportClass.
func (recv *ViewportClass) ToC() unsafe.Pointer {
	return recv.native
}

// ViewportClassNewFromC creates a new ViewportClass from a pointer to the C GtkViewportClass that represents the ViewportClass.
func ViewportClassNewFromC(native unsafe.Pointer) *ViewportClass {
	return &ViewportClass{native: native}
}

// ViewportPrivate is a representation of the C record GtkViewportPrivate.
type ViewportPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkViewportPrivate that represents the ViewportPrivate.
func (recv *ViewportPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ViewportPrivateNewFromC creates a new ViewportPrivate from a pointer to the C GtkViewportPrivate that represents the ViewportPrivate.
func ViewportPrivateNewFromC(native unsafe.Pointer) *ViewportPrivate {
	return &ViewportPrivate{native: native}
}

// VolumeButtonClass is a representation of the C record GtkVolumeButtonClass.
type VolumeButtonClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkVolumeButtonClass that represents the VolumeButtonClass.
func (recv *VolumeButtonClass) ToC() unsafe.Pointer {
	return recv.native
}

// VolumeButtonClassNewFromC creates a new VolumeButtonClass from a pointer to the C GtkVolumeButtonClass that represents the VolumeButtonClass.
func VolumeButtonClassNewFromC(native unsafe.Pointer) *VolumeButtonClass {
	return &VolumeButtonClass{native: native}
}

// WidgetAccessibleClass is a representation of the C record GtkWidgetAccessibleClass.
type WidgetAccessibleClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkWidgetAccessibleClass that represents the WidgetAccessibleClass.
func (recv *WidgetAccessibleClass) ToC() unsafe.Pointer {
	return recv.native
}

// WidgetAccessibleClassNewFromC creates a new WidgetAccessibleClass from a pointer to the C GtkWidgetAccessibleClass that represents the WidgetAccessibleClass.
func WidgetAccessibleClassNewFromC(native unsafe.Pointer) *WidgetAccessibleClass {
	return &WidgetAccessibleClass{native: native}
}

// WidgetAccessiblePrivate is a representation of the C record GtkWidgetAccessiblePrivate.
type WidgetAccessiblePrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkWidgetAccessiblePrivate that represents the WidgetAccessiblePrivate.
func (recv *WidgetAccessiblePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// WidgetAccessiblePrivateNewFromC creates a new WidgetAccessiblePrivate from a pointer to the C GtkWidgetAccessiblePrivate that represents the WidgetAccessiblePrivate.
func WidgetAccessiblePrivateNewFromC(native unsafe.Pointer) *WidgetAccessiblePrivate {
	return &WidgetAccessiblePrivate{native: native}
}

// WidgetClass is a representation of the C record GtkWidgetClass.
type WidgetClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkWidgetClass that represents the WidgetClass.
func (recv *WidgetClass) ToC() unsafe.Pointer {
	return recv.native
}

// WidgetClassNewFromC creates a new WidgetClass from a pointer to the C GtkWidgetClass that represents the WidgetClass.
func WidgetClassNewFromC(native unsafe.Pointer) *WidgetClass {
	return &WidgetClass{native: native}
}

// WidgetClassPrivate is a representation of the C record GtkWidgetClassPrivate.
type WidgetClassPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkWidgetClassPrivate that represents the WidgetClassPrivate.
func (recv *WidgetClassPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// WidgetClassPrivateNewFromC creates a new WidgetClassPrivate from a pointer to the C GtkWidgetClassPrivate that represents the WidgetClassPrivate.
func WidgetClassPrivateNewFromC(native unsafe.Pointer) *WidgetClassPrivate {
	return &WidgetClassPrivate{native: native}
}

// WidgetPath is a representation of the C record GtkWidgetPath.
type WidgetPath struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkWidgetPath that represents the WidgetPath.
func (recv *WidgetPath) ToC() unsafe.Pointer {
	return recv.native
}

// WidgetPathNewFromC creates a new WidgetPath from a pointer to the C GtkWidgetPath that represents the WidgetPath.
func WidgetPathNewFromC(native unsafe.Pointer) *WidgetPath {
	return &WidgetPath{native: native}
}

// WidgetPrivate is a representation of the C record GtkWidgetPrivate.
type WidgetPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkWidgetPrivate that represents the WidgetPrivate.
func (recv *WidgetPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// WidgetPrivateNewFromC creates a new WidgetPrivate from a pointer to the C GtkWidgetPrivate that represents the WidgetPrivate.
func WidgetPrivateNewFromC(native unsafe.Pointer) *WidgetPrivate {
	return &WidgetPrivate{native: native}
}

// WindowAccessibleClass is a representation of the C record GtkWindowAccessibleClass.
type WindowAccessibleClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkWindowAccessibleClass that represents the WindowAccessibleClass.
func (recv *WindowAccessibleClass) ToC() unsafe.Pointer {
	return recv.native
}

// WindowAccessibleClassNewFromC creates a new WindowAccessibleClass from a pointer to the C GtkWindowAccessibleClass that represents the WindowAccessibleClass.
func WindowAccessibleClassNewFromC(native unsafe.Pointer) *WindowAccessibleClass {
	return &WindowAccessibleClass{native: native}
}

// WindowAccessiblePrivate is a representation of the C record GtkWindowAccessiblePrivate.
type WindowAccessiblePrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkWindowAccessiblePrivate that represents the WindowAccessiblePrivate.
func (recv *WindowAccessiblePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// WindowAccessiblePrivateNewFromC creates a new WindowAccessiblePrivate from a pointer to the C GtkWindowAccessiblePrivate that represents the WindowAccessiblePrivate.
func WindowAccessiblePrivateNewFromC(native unsafe.Pointer) *WindowAccessiblePrivate {
	return &WindowAccessiblePrivate{native: native}
}

// WindowClass is a representation of the C record GtkWindowClass.
type WindowClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkWindowClass that represents the WindowClass.
func (recv *WindowClass) ToC() unsafe.Pointer {
	return recv.native
}

// WindowClassNewFromC creates a new WindowClass from a pointer to the C GtkWindowClass that represents the WindowClass.
func WindowClassNewFromC(native unsafe.Pointer) *WindowClass {
	return &WindowClass{native: native}
}

// WindowGeometryInfo is a representation of the C record GtkWindowGeometryInfo.
type WindowGeometryInfo struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkWindowGeometryInfo that represents the WindowGeometryInfo.
func (recv *WindowGeometryInfo) ToC() unsafe.Pointer {
	return recv.native
}

// WindowGeometryInfoNewFromC creates a new WindowGeometryInfo from a pointer to the C GtkWindowGeometryInfo that represents the WindowGeometryInfo.
func WindowGeometryInfoNewFromC(native unsafe.Pointer) *WindowGeometryInfo {
	return &WindowGeometryInfo{native: native}
}

// WindowGroupClass is a representation of the C record GtkWindowGroupClass.
type WindowGroupClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkWindowGroupClass that represents the WindowGroupClass.
func (recv *WindowGroupClass) ToC() unsafe.Pointer {
	return recv.native
}

// WindowGroupClassNewFromC creates a new WindowGroupClass from a pointer to the C GtkWindowGroupClass that represents the WindowGroupClass.
func WindowGroupClassNewFromC(native unsafe.Pointer) *WindowGroupClass {
	return &WindowGroupClass{native: native}
}

// WindowGroupPrivate is a representation of the C record GtkWindowGroupPrivate.
type WindowGroupPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkWindowGroupPrivate that represents the WindowGroupPrivate.
func (recv *WindowGroupPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// WindowGroupPrivateNewFromC creates a new WindowGroupPrivate from a pointer to the C GtkWindowGroupPrivate that represents the WindowGroupPrivate.
func WindowGroupPrivateNewFromC(native unsafe.Pointer) *WindowGroupPrivate {
	return &WindowGroupPrivate{native: native}
}

// WindowPrivate is a representation of the C record GtkWindowPrivate.
type WindowPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkWindowPrivate that represents the WindowPrivate.
func (recv *WindowPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// WindowPrivateNewFromC creates a new WindowPrivate from a pointer to the C GtkWindowPrivate that represents the WindowPrivate.
func WindowPrivateNewFromC(native unsafe.Pointer) *WindowPrivate {
	return &WindowPrivate{native: native}
}

// AboutDialog is a representation of the C record GtkAboutDialog.
type AboutDialog struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkAboutDialog that represents the AboutDialog.
func (recv *AboutDialog) ToC() unsafe.Pointer {
	return recv.native
}

// AboutDialogNewFromC creates a new AboutDialog from a pointer to the C GtkAboutDialog that represents the AboutDialog.
func AboutDialogNewFromC(native unsafe.Pointer) *AboutDialog {
	return &AboutDialog{native: native}
}

// AccelGroup is a representation of the C record GtkAccelGroup.
type AccelGroup struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkAccelGroup that represents the AccelGroup.
func (recv *AccelGroup) ToC() unsafe.Pointer {
	return recv.native
}

// AccelGroupNewFromC creates a new AccelGroup from a pointer to the C GtkAccelGroup that represents the AccelGroup.
func AccelGroupNewFromC(native unsafe.Pointer) *AccelGroup {
	return &AccelGroup{native: native}
}

// AccelLabel is a representation of the C record GtkAccelLabel.
type AccelLabel struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkAccelLabel that represents the AccelLabel.
func (recv *AccelLabel) ToC() unsafe.Pointer {
	return recv.native
}

// AccelLabelNewFromC creates a new AccelLabel from a pointer to the C GtkAccelLabel that represents the AccelLabel.
func AccelLabelNewFromC(native unsafe.Pointer) *AccelLabel {
	return &AccelLabel{native: native}
}

// AccelMap is a representation of the C record GtkAccelMap.
type AccelMap struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkAccelMap that represents the AccelMap.
func (recv *AccelMap) ToC() unsafe.Pointer {
	return recv.native
}

// AccelMapNewFromC creates a new AccelMap from a pointer to the C GtkAccelMap that represents the AccelMap.
func AccelMapNewFromC(native unsafe.Pointer) *AccelMap {
	return &AccelMap{native: native}
}

// Accessible is a representation of the C record GtkAccessible.
type Accessible struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkAccessible that represents the Accessible.
func (recv *Accessible) ToC() unsafe.Pointer {
	return recv.native
}

// AccessibleNewFromC creates a new Accessible from a pointer to the C GtkAccessible that represents the Accessible.
func AccessibleNewFromC(native unsafe.Pointer) *Accessible {
	return &Accessible{native: native}
}

// Action is a representation of the C record GtkAction.
type Action struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkAction that represents the Action.
func (recv *Action) ToC() unsafe.Pointer {
	return recv.native
}

// ActionNewFromC creates a new Action from a pointer to the C GtkAction that represents the Action.
func ActionNewFromC(native unsafe.Pointer) *Action {
	return &Action{native: native}
}

// ActionBar is a representation of the C record GtkActionBar.
type ActionBar struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkActionBar that represents the ActionBar.
func (recv *ActionBar) ToC() unsafe.Pointer {
	return recv.native
}

// ActionBarNewFromC creates a new ActionBar from a pointer to the C GtkActionBar that represents the ActionBar.
func ActionBarNewFromC(native unsafe.Pointer) *ActionBar {
	return &ActionBar{native: native}
}

// ActionGroup is a representation of the C record GtkActionGroup.
type ActionGroup struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkActionGroup that represents the ActionGroup.
func (recv *ActionGroup) ToC() unsafe.Pointer {
	return recv.native
}

// ActionGroupNewFromC creates a new ActionGroup from a pointer to the C GtkActionGroup that represents the ActionGroup.
func ActionGroupNewFromC(native unsafe.Pointer) *ActionGroup {
	return &ActionGroup{native: native}
}

// Adjustment is a representation of the C record GtkAdjustment.
type Adjustment struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkAdjustment that represents the Adjustment.
func (recv *Adjustment) ToC() unsafe.Pointer {
	return recv.native
}

// AdjustmentNewFromC creates a new Adjustment from a pointer to the C GtkAdjustment that represents the Adjustment.
func AdjustmentNewFromC(native unsafe.Pointer) *Adjustment {
	return &Adjustment{native: native}
}

// Alignment is a representation of the C record GtkAlignment.
type Alignment struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkAlignment that represents the Alignment.
func (recv *Alignment) ToC() unsafe.Pointer {
	return recv.native
}

// AlignmentNewFromC creates a new Alignment from a pointer to the C GtkAlignment that represents the Alignment.
func AlignmentNewFromC(native unsafe.Pointer) *Alignment {
	return &Alignment{native: native}
}

// AppChooserButton is a representation of the C record GtkAppChooserButton.
type AppChooserButton struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkAppChooserButton that represents the AppChooserButton.
func (recv *AppChooserButton) ToC() unsafe.Pointer {
	return recv.native
}

// AppChooserButtonNewFromC creates a new AppChooserButton from a pointer to the C GtkAppChooserButton that represents the AppChooserButton.
func AppChooserButtonNewFromC(native unsafe.Pointer) *AppChooserButton {
	return &AppChooserButton{native: native}
}

// AppChooserDialog is a representation of the C record GtkAppChooserDialog.
type AppChooserDialog struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkAppChooserDialog that represents the AppChooserDialog.
func (recv *AppChooserDialog) ToC() unsafe.Pointer {
	return recv.native
}

// AppChooserDialogNewFromC creates a new AppChooserDialog from a pointer to the C GtkAppChooserDialog that represents the AppChooserDialog.
func AppChooserDialogNewFromC(native unsafe.Pointer) *AppChooserDialog {
	return &AppChooserDialog{native: native}
}

// AppChooserWidget is a representation of the C record GtkAppChooserWidget.
type AppChooserWidget struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkAppChooserWidget that represents the AppChooserWidget.
func (recv *AppChooserWidget) ToC() unsafe.Pointer {
	return recv.native
}

// AppChooserWidgetNewFromC creates a new AppChooserWidget from a pointer to the C GtkAppChooserWidget that represents the AppChooserWidget.
func AppChooserWidgetNewFromC(native unsafe.Pointer) *AppChooserWidget {
	return &AppChooserWidget{native: native}
}

// Application is a representation of the C record GtkApplication.
type Application struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkApplication that represents the Application.
func (recv *Application) ToC() unsafe.Pointer {
	return recv.native
}

// ApplicationNewFromC creates a new Application from a pointer to the C GtkApplication that represents the Application.
func ApplicationNewFromC(native unsafe.Pointer) *Application {
	return &Application{native: native}
}

// ApplicationWindow is a representation of the C record GtkApplicationWindow.
type ApplicationWindow struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkApplicationWindow that represents the ApplicationWindow.
func (recv *ApplicationWindow) ToC() unsafe.Pointer {
	return recv.native
}

// ApplicationWindowNewFromC creates a new ApplicationWindow from a pointer to the C GtkApplicationWindow that represents the ApplicationWindow.
func ApplicationWindowNewFromC(native unsafe.Pointer) *ApplicationWindow {
	return &ApplicationWindow{native: native}
}

// Arrow is a representation of the C record GtkArrow.
type Arrow struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkArrow that represents the Arrow.
func (recv *Arrow) ToC() unsafe.Pointer {
	return recv.native
}

// ArrowNewFromC creates a new Arrow from a pointer to the C GtkArrow that represents the Arrow.
func ArrowNewFromC(native unsafe.Pointer) *Arrow {
	return &Arrow{native: native}
}

// ArrowAccessible is a representation of the C record GtkArrowAccessible.
type ArrowAccessible struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkArrowAccessible that represents the ArrowAccessible.
func (recv *ArrowAccessible) ToC() unsafe.Pointer {
	return recv.native
}

// ArrowAccessibleNewFromC creates a new ArrowAccessible from a pointer to the C GtkArrowAccessible that represents the ArrowAccessible.
func ArrowAccessibleNewFromC(native unsafe.Pointer) *ArrowAccessible {
	return &ArrowAccessible{native: native}
}

// AspectFrame is a representation of the C record GtkAspectFrame.
type AspectFrame struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkAspectFrame that represents the AspectFrame.
func (recv *AspectFrame) ToC() unsafe.Pointer {
	return recv.native
}

// AspectFrameNewFromC creates a new AspectFrame from a pointer to the C GtkAspectFrame that represents the AspectFrame.
func AspectFrameNewFromC(native unsafe.Pointer) *AspectFrame {
	return &AspectFrame{native: native}
}

// Assistant is a representation of the C record GtkAssistant.
type Assistant struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkAssistant that represents the Assistant.
func (recv *Assistant) ToC() unsafe.Pointer {
	return recv.native
}

// AssistantNewFromC creates a new Assistant from a pointer to the C GtkAssistant that represents the Assistant.
func AssistantNewFromC(native unsafe.Pointer) *Assistant {
	return &Assistant{native: native}
}

// Bin is a representation of the C record GtkBin.
type Bin struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkBin that represents the Bin.
func (recv *Bin) ToC() unsafe.Pointer {
	return recv.native
}

// BinNewFromC creates a new Bin from a pointer to the C GtkBin that represents the Bin.
func BinNewFromC(native unsafe.Pointer) *Bin {
	return &Bin{native: native}
}

// BooleanCellAccessible is a representation of the C record GtkBooleanCellAccessible.
type BooleanCellAccessible struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkBooleanCellAccessible that represents the BooleanCellAccessible.
func (recv *BooleanCellAccessible) ToC() unsafe.Pointer {
	return recv.native
}

// BooleanCellAccessibleNewFromC creates a new BooleanCellAccessible from a pointer to the C GtkBooleanCellAccessible that represents the BooleanCellAccessible.
func BooleanCellAccessibleNewFromC(native unsafe.Pointer) *BooleanCellAccessible {
	return &BooleanCellAccessible{native: native}
}

// Box is a representation of the C record GtkBox.
type Box struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkBox that represents the Box.
func (recv *Box) ToC() unsafe.Pointer {
	return recv.native
}

// BoxNewFromC creates a new Box from a pointer to the C GtkBox that represents the Box.
func BoxNewFromC(native unsafe.Pointer) *Box {
	return &Box{native: native}
}

// Builder is a representation of the C record GtkBuilder.
type Builder struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkBuilder that represents the Builder.
func (recv *Builder) ToC() unsafe.Pointer {
	return recv.native
}

// BuilderNewFromC creates a new Builder from a pointer to the C GtkBuilder that represents the Builder.
func BuilderNewFromC(native unsafe.Pointer) *Builder {
	return &Builder{native: native}
}

// Button is a representation of the C record GtkButton.
type Button struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkButton that represents the Button.
func (recv *Button) ToC() unsafe.Pointer {
	return recv.native
}

// ButtonNewFromC creates a new Button from a pointer to the C GtkButton that represents the Button.
func ButtonNewFromC(native unsafe.Pointer) *Button {
	return &Button{native: native}
}

// ButtonAccessible is a representation of the C record GtkButtonAccessible.
type ButtonAccessible struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkButtonAccessible that represents the ButtonAccessible.
func (recv *ButtonAccessible) ToC() unsafe.Pointer {
	return recv.native
}

// ButtonAccessibleNewFromC creates a new ButtonAccessible from a pointer to the C GtkButtonAccessible that represents the ButtonAccessible.
func ButtonAccessibleNewFromC(native unsafe.Pointer) *ButtonAccessible {
	return &ButtonAccessible{native: native}
}

// ButtonBox is a representation of the C record GtkButtonBox.
type ButtonBox struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkButtonBox that represents the ButtonBox.
func (recv *ButtonBox) ToC() unsafe.Pointer {
	return recv.native
}

// ButtonBoxNewFromC creates a new ButtonBox from a pointer to the C GtkButtonBox that represents the ButtonBox.
func ButtonBoxNewFromC(native unsafe.Pointer) *ButtonBox {
	return &ButtonBox{native: native}
}

// Calendar is a representation of the C record GtkCalendar.
type Calendar struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkCalendar that represents the Calendar.
func (recv *Calendar) ToC() unsafe.Pointer {
	return recv.native
}

// CalendarNewFromC creates a new Calendar from a pointer to the C GtkCalendar that represents the Calendar.
func CalendarNewFromC(native unsafe.Pointer) *Calendar {
	return &Calendar{native: native}
}

// CellAccessible is a representation of the C record GtkCellAccessible.
type CellAccessible struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkCellAccessible that represents the CellAccessible.
func (recv *CellAccessible) ToC() unsafe.Pointer {
	return recv.native
}

// CellAccessibleNewFromC creates a new CellAccessible from a pointer to the C GtkCellAccessible that represents the CellAccessible.
func CellAccessibleNewFromC(native unsafe.Pointer) *CellAccessible {
	return &CellAccessible{native: native}
}

// CellArea is a representation of the C record GtkCellArea.
type CellArea struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkCellArea that represents the CellArea.
func (recv *CellArea) ToC() unsafe.Pointer {
	return recv.native
}

// CellAreaNewFromC creates a new CellArea from a pointer to the C GtkCellArea that represents the CellArea.
func CellAreaNewFromC(native unsafe.Pointer) *CellArea {
	return &CellArea{native: native}
}

// CellAreaBox is a representation of the C record GtkCellAreaBox.
type CellAreaBox struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkCellAreaBox that represents the CellAreaBox.
func (recv *CellAreaBox) ToC() unsafe.Pointer {
	return recv.native
}

// CellAreaBoxNewFromC creates a new CellAreaBox from a pointer to the C GtkCellAreaBox that represents the CellAreaBox.
func CellAreaBoxNewFromC(native unsafe.Pointer) *CellAreaBox {
	return &CellAreaBox{native: native}
}

// CellAreaContext is a representation of the C record GtkCellAreaContext.
type CellAreaContext struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkCellAreaContext that represents the CellAreaContext.
func (recv *CellAreaContext) ToC() unsafe.Pointer {
	return recv.native
}

// CellAreaContextNewFromC creates a new CellAreaContext from a pointer to the C GtkCellAreaContext that represents the CellAreaContext.
func CellAreaContextNewFromC(native unsafe.Pointer) *CellAreaContext {
	return &CellAreaContext{native: native}
}

// CellRenderer is a representation of the C record GtkCellRenderer.
type CellRenderer struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkCellRenderer that represents the CellRenderer.
func (recv *CellRenderer) ToC() unsafe.Pointer {
	return recv.native
}

// CellRendererNewFromC creates a new CellRenderer from a pointer to the C GtkCellRenderer that represents the CellRenderer.
func CellRendererNewFromC(native unsafe.Pointer) *CellRenderer {
	return &CellRenderer{native: native}
}

// CellRendererAccel is a representation of the C record GtkCellRendererAccel.
type CellRendererAccel struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkCellRendererAccel that represents the CellRendererAccel.
func (recv *CellRendererAccel) ToC() unsafe.Pointer {
	return recv.native
}

// CellRendererAccelNewFromC creates a new CellRendererAccel from a pointer to the C GtkCellRendererAccel that represents the CellRendererAccel.
func CellRendererAccelNewFromC(native unsafe.Pointer) *CellRendererAccel {
	return &CellRendererAccel{native: native}
}

// CellRendererCombo is a representation of the C record GtkCellRendererCombo.
type CellRendererCombo struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkCellRendererCombo that represents the CellRendererCombo.
func (recv *CellRendererCombo) ToC() unsafe.Pointer {
	return recv.native
}

// CellRendererComboNewFromC creates a new CellRendererCombo from a pointer to the C GtkCellRendererCombo that represents the CellRendererCombo.
func CellRendererComboNewFromC(native unsafe.Pointer) *CellRendererCombo {
	return &CellRendererCombo{native: native}
}

// CellRendererPixbuf is a representation of the C record GtkCellRendererPixbuf.
type CellRendererPixbuf struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkCellRendererPixbuf that represents the CellRendererPixbuf.
func (recv *CellRendererPixbuf) ToC() unsafe.Pointer {
	return recv.native
}

// CellRendererPixbufNewFromC creates a new CellRendererPixbuf from a pointer to the C GtkCellRendererPixbuf that represents the CellRendererPixbuf.
func CellRendererPixbufNewFromC(native unsafe.Pointer) *CellRendererPixbuf {
	return &CellRendererPixbuf{native: native}
}

// CellRendererProgress is a representation of the C record GtkCellRendererProgress.
type CellRendererProgress struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkCellRendererProgress that represents the CellRendererProgress.
func (recv *CellRendererProgress) ToC() unsafe.Pointer {
	return recv.native
}

// CellRendererProgressNewFromC creates a new CellRendererProgress from a pointer to the C GtkCellRendererProgress that represents the CellRendererProgress.
func CellRendererProgressNewFromC(native unsafe.Pointer) *CellRendererProgress {
	return &CellRendererProgress{native: native}
}

// CellRendererSpin is a representation of the C record GtkCellRendererSpin.
type CellRendererSpin struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkCellRendererSpin that represents the CellRendererSpin.
func (recv *CellRendererSpin) ToC() unsafe.Pointer {
	return recv.native
}

// CellRendererSpinNewFromC creates a new CellRendererSpin from a pointer to the C GtkCellRendererSpin that represents the CellRendererSpin.
func CellRendererSpinNewFromC(native unsafe.Pointer) *CellRendererSpin {
	return &CellRendererSpin{native: native}
}

// CellRendererSpinner is a representation of the C record GtkCellRendererSpinner.
type CellRendererSpinner struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkCellRendererSpinner that represents the CellRendererSpinner.
func (recv *CellRendererSpinner) ToC() unsafe.Pointer {
	return recv.native
}

// CellRendererSpinnerNewFromC creates a new CellRendererSpinner from a pointer to the C GtkCellRendererSpinner that represents the CellRendererSpinner.
func CellRendererSpinnerNewFromC(native unsafe.Pointer) *CellRendererSpinner {
	return &CellRendererSpinner{native: native}
}

// CellRendererText is a representation of the C record GtkCellRendererText.
type CellRendererText struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkCellRendererText that represents the CellRendererText.
func (recv *CellRendererText) ToC() unsafe.Pointer {
	return recv.native
}

// CellRendererTextNewFromC creates a new CellRendererText from a pointer to the C GtkCellRendererText that represents the CellRendererText.
func CellRendererTextNewFromC(native unsafe.Pointer) *CellRendererText {
	return &CellRendererText{native: native}
}

// CellRendererToggle is a representation of the C record GtkCellRendererToggle.
type CellRendererToggle struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkCellRendererToggle that represents the CellRendererToggle.
func (recv *CellRendererToggle) ToC() unsafe.Pointer {
	return recv.native
}

// CellRendererToggleNewFromC creates a new CellRendererToggle from a pointer to the C GtkCellRendererToggle that represents the CellRendererToggle.
func CellRendererToggleNewFromC(native unsafe.Pointer) *CellRendererToggle {
	return &CellRendererToggle{native: native}
}

// CellView is a representation of the C record GtkCellView.
type CellView struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkCellView that represents the CellView.
func (recv *CellView) ToC() unsafe.Pointer {
	return recv.native
}

// CellViewNewFromC creates a new CellView from a pointer to the C GtkCellView that represents the CellView.
func CellViewNewFromC(native unsafe.Pointer) *CellView {
	return &CellView{native: native}
}

// CheckButton is a representation of the C record GtkCheckButton.
type CheckButton struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkCheckButton that represents the CheckButton.
func (recv *CheckButton) ToC() unsafe.Pointer {
	return recv.native
}

// CheckButtonNewFromC creates a new CheckButton from a pointer to the C GtkCheckButton that represents the CheckButton.
func CheckButtonNewFromC(native unsafe.Pointer) *CheckButton {
	return &CheckButton{native: native}
}

// CheckMenuItem is a representation of the C record GtkCheckMenuItem.
type CheckMenuItem struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkCheckMenuItem that represents the CheckMenuItem.
func (recv *CheckMenuItem) ToC() unsafe.Pointer {
	return recv.native
}

// CheckMenuItemNewFromC creates a new CheckMenuItem from a pointer to the C GtkCheckMenuItem that represents the CheckMenuItem.
func CheckMenuItemNewFromC(native unsafe.Pointer) *CheckMenuItem {
	return &CheckMenuItem{native: native}
}

// CheckMenuItemAccessible is a representation of the C record GtkCheckMenuItemAccessible.
type CheckMenuItemAccessible struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkCheckMenuItemAccessible that represents the CheckMenuItemAccessible.
func (recv *CheckMenuItemAccessible) ToC() unsafe.Pointer {
	return recv.native
}

// CheckMenuItemAccessibleNewFromC creates a new CheckMenuItemAccessible from a pointer to the C GtkCheckMenuItemAccessible that represents the CheckMenuItemAccessible.
func CheckMenuItemAccessibleNewFromC(native unsafe.Pointer) *CheckMenuItemAccessible {
	return &CheckMenuItemAccessible{native: native}
}

// Clipboard is a representation of the C record GtkClipboard.
type Clipboard struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkClipboard that represents the Clipboard.
func (recv *Clipboard) ToC() unsafe.Pointer {
	return recv.native
}

// ClipboardNewFromC creates a new Clipboard from a pointer to the C GtkClipboard that represents the Clipboard.
func ClipboardNewFromC(native unsafe.Pointer) *Clipboard {
	return &Clipboard{native: native}
}

// ColorButton is a representation of the C record GtkColorButton.
type ColorButton struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkColorButton that represents the ColorButton.
func (recv *ColorButton) ToC() unsafe.Pointer {
	return recv.native
}

// ColorButtonNewFromC creates a new ColorButton from a pointer to the C GtkColorButton that represents the ColorButton.
func ColorButtonNewFromC(native unsafe.Pointer) *ColorButton {
	return &ColorButton{native: native}
}

// ColorSelection is a representation of the C record GtkColorSelection.
type ColorSelection struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkColorSelection that represents the ColorSelection.
func (recv *ColorSelection) ToC() unsafe.Pointer {
	return recv.native
}

// ColorSelectionNewFromC creates a new ColorSelection from a pointer to the C GtkColorSelection that represents the ColorSelection.
func ColorSelectionNewFromC(native unsafe.Pointer) *ColorSelection {
	return &ColorSelection{native: native}
}

// ColorSelectionDialog is a representation of the C record GtkColorSelectionDialog.
type ColorSelectionDialog struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkColorSelectionDialog that represents the ColorSelectionDialog.
func (recv *ColorSelectionDialog) ToC() unsafe.Pointer {
	return recv.native
}

// ColorSelectionDialogNewFromC creates a new ColorSelectionDialog from a pointer to the C GtkColorSelectionDialog that represents the ColorSelectionDialog.
func ColorSelectionDialogNewFromC(native unsafe.Pointer) *ColorSelectionDialog {
	return &ColorSelectionDialog{native: native}
}

// ComboBox is a representation of the C record GtkComboBox.
type ComboBox struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkComboBox that represents the ComboBox.
func (recv *ComboBox) ToC() unsafe.Pointer {
	return recv.native
}

// ComboBoxNewFromC creates a new ComboBox from a pointer to the C GtkComboBox that represents the ComboBox.
func ComboBoxNewFromC(native unsafe.Pointer) *ComboBox {
	return &ComboBox{native: native}
}

// ComboBoxAccessible is a representation of the C record GtkComboBoxAccessible.
type ComboBoxAccessible struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkComboBoxAccessible that represents the ComboBoxAccessible.
func (recv *ComboBoxAccessible) ToC() unsafe.Pointer {
	return recv.native
}

// ComboBoxAccessibleNewFromC creates a new ComboBoxAccessible from a pointer to the C GtkComboBoxAccessible that represents the ComboBoxAccessible.
func ComboBoxAccessibleNewFromC(native unsafe.Pointer) *ComboBoxAccessible {
	return &ComboBoxAccessible{native: native}
}

// ComboBoxText is a representation of the C record GtkComboBoxText.
type ComboBoxText struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkComboBoxText that represents the ComboBoxText.
func (recv *ComboBoxText) ToC() unsafe.Pointer {
	return recv.native
}

// ComboBoxTextNewFromC creates a new ComboBoxText from a pointer to the C GtkComboBoxText that represents the ComboBoxText.
func ComboBoxTextNewFromC(native unsafe.Pointer) *ComboBoxText {
	return &ComboBoxText{native: native}
}

// Container is a representation of the C record GtkContainer.
type Container struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkContainer that represents the Container.
func (recv *Container) ToC() unsafe.Pointer {
	return recv.native
}

// ContainerNewFromC creates a new Container from a pointer to the C GtkContainer that represents the Container.
func ContainerNewFromC(native unsafe.Pointer) *Container {
	return &Container{native: native}
}

// ContainerAccessible is a representation of the C record GtkContainerAccessible.
type ContainerAccessible struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkContainerAccessible that represents the ContainerAccessible.
func (recv *ContainerAccessible) ToC() unsafe.Pointer {
	return recv.native
}

// ContainerAccessibleNewFromC creates a new ContainerAccessible from a pointer to the C GtkContainerAccessible that represents the ContainerAccessible.
func ContainerAccessibleNewFromC(native unsafe.Pointer) *ContainerAccessible {
	return &ContainerAccessible{native: native}
}

// ContainerCellAccessible is a representation of the C record GtkContainerCellAccessible.
type ContainerCellAccessible struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkContainerCellAccessible that represents the ContainerCellAccessible.
func (recv *ContainerCellAccessible) ToC() unsafe.Pointer {
	return recv.native
}

// ContainerCellAccessibleNewFromC creates a new ContainerCellAccessible from a pointer to the C GtkContainerCellAccessible that represents the ContainerCellAccessible.
func ContainerCellAccessibleNewFromC(native unsafe.Pointer) *ContainerCellAccessible {
	return &ContainerCellAccessible{native: native}
}

// CssProvider is a representation of the C record GtkCssProvider.
type CssProvider struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkCssProvider that represents the CssProvider.
func (recv *CssProvider) ToC() unsafe.Pointer {
	return recv.native
}

// CssProviderNewFromC creates a new CssProvider from a pointer to the C GtkCssProvider that represents the CssProvider.
func CssProviderNewFromC(native unsafe.Pointer) *CssProvider {
	return &CssProvider{native: native}
}

// Dialog is a representation of the C record GtkDialog.
type Dialog struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkDialog that represents the Dialog.
func (recv *Dialog) ToC() unsafe.Pointer {
	return recv.native
}

// DialogNewFromC creates a new Dialog from a pointer to the C GtkDialog that represents the Dialog.
func DialogNewFromC(native unsafe.Pointer) *Dialog {
	return &Dialog{native: native}
}

// DrawingArea is a representation of the C record GtkDrawingArea.
type DrawingArea struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkDrawingArea that represents the DrawingArea.
func (recv *DrawingArea) ToC() unsafe.Pointer {
	return recv.native
}

// DrawingAreaNewFromC creates a new DrawingArea from a pointer to the C GtkDrawingArea that represents the DrawingArea.
func DrawingAreaNewFromC(native unsafe.Pointer) *DrawingArea {
	return &DrawingArea{native: native}
}

// Entry is a representation of the C record GtkEntry.
type Entry struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkEntry that represents the Entry.
func (recv *Entry) ToC() unsafe.Pointer {
	return recv.native
}

// EntryNewFromC creates a new Entry from a pointer to the C GtkEntry that represents the Entry.
func EntryNewFromC(native unsafe.Pointer) *Entry {
	return &Entry{native: native}
}

// EntryAccessible is a representation of the C record GtkEntryAccessible.
type EntryAccessible struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkEntryAccessible that represents the EntryAccessible.
func (recv *EntryAccessible) ToC() unsafe.Pointer {
	return recv.native
}

// EntryAccessibleNewFromC creates a new EntryAccessible from a pointer to the C GtkEntryAccessible that represents the EntryAccessible.
func EntryAccessibleNewFromC(native unsafe.Pointer) *EntryAccessible {
	return &EntryAccessible{native: native}
}

// EntryCompletion is a representation of the C record GtkEntryCompletion.
type EntryCompletion struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkEntryCompletion that represents the EntryCompletion.
func (recv *EntryCompletion) ToC() unsafe.Pointer {
	return recv.native
}

// EntryCompletionNewFromC creates a new EntryCompletion from a pointer to the C GtkEntryCompletion that represents the EntryCompletion.
func EntryCompletionNewFromC(native unsafe.Pointer) *EntryCompletion {
	return &EntryCompletion{native: native}
}

// UNSUPPORTED : EntryIconAccessible : blacklisted

// EventBox is a representation of the C record GtkEventBox.
type EventBox struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkEventBox that represents the EventBox.
func (recv *EventBox) ToC() unsafe.Pointer {
	return recv.native
}

// EventBoxNewFromC creates a new EventBox from a pointer to the C GtkEventBox that represents the EventBox.
func EventBoxNewFromC(native unsafe.Pointer) *EventBox {
	return &EventBox{native: native}
}

// EventController is a representation of the C record GtkEventController.
type EventController struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkEventController that represents the EventController.
func (recv *EventController) ToC() unsafe.Pointer {
	return recv.native
}

// EventControllerNewFromC creates a new EventController from a pointer to the C GtkEventController that represents the EventController.
func EventControllerNewFromC(native unsafe.Pointer) *EventController {
	return &EventController{native: native}
}

// UNSUPPORTED : EventControllerKey : blacklisted

// UNSUPPORTED : EventControllerMotion : blacklisted

// UNSUPPORTED : EventControllerScroll : blacklisted

// Expander is a representation of the C record GtkExpander.
type Expander struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkExpander that represents the Expander.
func (recv *Expander) ToC() unsafe.Pointer {
	return recv.native
}

// ExpanderNewFromC creates a new Expander from a pointer to the C GtkExpander that represents the Expander.
func ExpanderNewFromC(native unsafe.Pointer) *Expander {
	return &Expander{native: native}
}

// ExpanderAccessible is a representation of the C record GtkExpanderAccessible.
type ExpanderAccessible struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkExpanderAccessible that represents the ExpanderAccessible.
func (recv *ExpanderAccessible) ToC() unsafe.Pointer {
	return recv.native
}

// ExpanderAccessibleNewFromC creates a new ExpanderAccessible from a pointer to the C GtkExpanderAccessible that represents the ExpanderAccessible.
func ExpanderAccessibleNewFromC(native unsafe.Pointer) *ExpanderAccessible {
	return &ExpanderAccessible{native: native}
}

// FileChooserButton is a representation of the C record GtkFileChooserButton.
type FileChooserButton struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkFileChooserButton that represents the FileChooserButton.
func (recv *FileChooserButton) ToC() unsafe.Pointer {
	return recv.native
}

// FileChooserButtonNewFromC creates a new FileChooserButton from a pointer to the C GtkFileChooserButton that represents the FileChooserButton.
func FileChooserButtonNewFromC(native unsafe.Pointer) *FileChooserButton {
	return &FileChooserButton{native: native}
}

// FileChooserDialog is a representation of the C record GtkFileChooserDialog.
type FileChooserDialog struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkFileChooserDialog that represents the FileChooserDialog.
func (recv *FileChooserDialog) ToC() unsafe.Pointer {
	return recv.native
}

// FileChooserDialogNewFromC creates a new FileChooserDialog from a pointer to the C GtkFileChooserDialog that represents the FileChooserDialog.
func FileChooserDialogNewFromC(native unsafe.Pointer) *FileChooserDialog {
	return &FileChooserDialog{native: native}
}

// FileChooserNative is a representation of the C record GtkFileChooserNative.
type FileChooserNative struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkFileChooserNative that represents the FileChooserNative.
func (recv *FileChooserNative) ToC() unsafe.Pointer {
	return recv.native
}

// FileChooserNativeNewFromC creates a new FileChooserNative from a pointer to the C GtkFileChooserNative that represents the FileChooserNative.
func FileChooserNativeNewFromC(native unsafe.Pointer) *FileChooserNative {
	return &FileChooserNative{native: native}
}

// FileChooserWidget is a representation of the C record GtkFileChooserWidget.
type FileChooserWidget struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkFileChooserWidget that represents the FileChooserWidget.
func (recv *FileChooserWidget) ToC() unsafe.Pointer {
	return recv.native
}

// FileChooserWidgetNewFromC creates a new FileChooserWidget from a pointer to the C GtkFileChooserWidget that represents the FileChooserWidget.
func FileChooserWidgetNewFromC(native unsafe.Pointer) *FileChooserWidget {
	return &FileChooserWidget{native: native}
}

// FileFilter is a representation of the C record GtkFileFilter.
type FileFilter struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkFileFilter that represents the FileFilter.
func (recv *FileFilter) ToC() unsafe.Pointer {
	return recv.native
}

// FileFilterNewFromC creates a new FileFilter from a pointer to the C GtkFileFilter that represents the FileFilter.
func FileFilterNewFromC(native unsafe.Pointer) *FileFilter {
	return &FileFilter{native: native}
}

// Fixed is a representation of the C record GtkFixed.
type Fixed struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkFixed that represents the Fixed.
func (recv *Fixed) ToC() unsafe.Pointer {
	return recv.native
}

// FixedNewFromC creates a new Fixed from a pointer to the C GtkFixed that represents the Fixed.
func FixedNewFromC(native unsafe.Pointer) *Fixed {
	return &Fixed{native: native}
}

// FlowBox is a representation of the C record GtkFlowBox.
type FlowBox struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkFlowBox that represents the FlowBox.
func (recv *FlowBox) ToC() unsafe.Pointer {
	return recv.native
}

// FlowBoxNewFromC creates a new FlowBox from a pointer to the C GtkFlowBox that represents the FlowBox.
func FlowBoxNewFromC(native unsafe.Pointer) *FlowBox {
	return &FlowBox{native: native}
}

// FlowBoxAccessible is a representation of the C record GtkFlowBoxAccessible.
type FlowBoxAccessible struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkFlowBoxAccessible that represents the FlowBoxAccessible.
func (recv *FlowBoxAccessible) ToC() unsafe.Pointer {
	return recv.native
}

// FlowBoxAccessibleNewFromC creates a new FlowBoxAccessible from a pointer to the C GtkFlowBoxAccessible that represents the FlowBoxAccessible.
func FlowBoxAccessibleNewFromC(native unsafe.Pointer) *FlowBoxAccessible {
	return &FlowBoxAccessible{native: native}
}

// FlowBoxChild is a representation of the C record GtkFlowBoxChild.
type FlowBoxChild struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkFlowBoxChild that represents the FlowBoxChild.
func (recv *FlowBoxChild) ToC() unsafe.Pointer {
	return recv.native
}

// FlowBoxChildNewFromC creates a new FlowBoxChild from a pointer to the C GtkFlowBoxChild that represents the FlowBoxChild.
func FlowBoxChildNewFromC(native unsafe.Pointer) *FlowBoxChild {
	return &FlowBoxChild{native: native}
}

// FlowBoxChildAccessible is a representation of the C record GtkFlowBoxChildAccessible.
type FlowBoxChildAccessible struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkFlowBoxChildAccessible that represents the FlowBoxChildAccessible.
func (recv *FlowBoxChildAccessible) ToC() unsafe.Pointer {
	return recv.native
}

// FlowBoxChildAccessibleNewFromC creates a new FlowBoxChildAccessible from a pointer to the C GtkFlowBoxChildAccessible that represents the FlowBoxChildAccessible.
func FlowBoxChildAccessibleNewFromC(native unsafe.Pointer) *FlowBoxChildAccessible {
	return &FlowBoxChildAccessible{native: native}
}

// FontButton is a representation of the C record GtkFontButton.
type FontButton struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkFontButton that represents the FontButton.
func (recv *FontButton) ToC() unsafe.Pointer {
	return recv.native
}

// FontButtonNewFromC creates a new FontButton from a pointer to the C GtkFontButton that represents the FontButton.
func FontButtonNewFromC(native unsafe.Pointer) *FontButton {
	return &FontButton{native: native}
}

// FontSelection is a representation of the C record GtkFontSelection.
type FontSelection struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkFontSelection that represents the FontSelection.
func (recv *FontSelection) ToC() unsafe.Pointer {
	return recv.native
}

// FontSelectionNewFromC creates a new FontSelection from a pointer to the C GtkFontSelection that represents the FontSelection.
func FontSelectionNewFromC(native unsafe.Pointer) *FontSelection {
	return &FontSelection{native: native}
}

// FontSelectionDialog is a representation of the C record GtkFontSelectionDialog.
type FontSelectionDialog struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkFontSelectionDialog that represents the FontSelectionDialog.
func (recv *FontSelectionDialog) ToC() unsafe.Pointer {
	return recv.native
}

// FontSelectionDialogNewFromC creates a new FontSelectionDialog from a pointer to the C GtkFontSelectionDialog that represents the FontSelectionDialog.
func FontSelectionDialogNewFromC(native unsafe.Pointer) *FontSelectionDialog {
	return &FontSelectionDialog{native: native}
}

// Frame is a representation of the C record GtkFrame.
type Frame struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkFrame that represents the Frame.
func (recv *Frame) ToC() unsafe.Pointer {
	return recv.native
}

// FrameNewFromC creates a new Frame from a pointer to the C GtkFrame that represents the Frame.
func FrameNewFromC(native unsafe.Pointer) *Frame {
	return &Frame{native: native}
}

// FrameAccessible is a representation of the C record GtkFrameAccessible.
type FrameAccessible struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkFrameAccessible that represents the FrameAccessible.
func (recv *FrameAccessible) ToC() unsafe.Pointer {
	return recv.native
}

// FrameAccessibleNewFromC creates a new FrameAccessible from a pointer to the C GtkFrameAccessible that represents the FrameAccessible.
func FrameAccessibleNewFromC(native unsafe.Pointer) *FrameAccessible {
	return &FrameAccessible{native: native}
}

// Gesture is a representation of the C record GtkGesture.
type Gesture struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkGesture that represents the Gesture.
func (recv *Gesture) ToC() unsafe.Pointer {
	return recv.native
}

// GestureNewFromC creates a new Gesture from a pointer to the C GtkGesture that represents the Gesture.
func GestureNewFromC(native unsafe.Pointer) *Gesture {
	return &Gesture{native: native}
}

// GestureDrag is a representation of the C record GtkGestureDrag.
type GestureDrag struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkGestureDrag that represents the GestureDrag.
func (recv *GestureDrag) ToC() unsafe.Pointer {
	return recv.native
}

// GestureDragNewFromC creates a new GestureDrag from a pointer to the C GtkGestureDrag that represents the GestureDrag.
func GestureDragNewFromC(native unsafe.Pointer) *GestureDrag {
	return &GestureDrag{native: native}
}

// GestureLongPress is a representation of the C record GtkGestureLongPress.
type GestureLongPress struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkGestureLongPress that represents the GestureLongPress.
func (recv *GestureLongPress) ToC() unsafe.Pointer {
	return recv.native
}

// GestureLongPressNewFromC creates a new GestureLongPress from a pointer to the C GtkGestureLongPress that represents the GestureLongPress.
func GestureLongPressNewFromC(native unsafe.Pointer) *GestureLongPress {
	return &GestureLongPress{native: native}
}

// GestureMultiPress is a representation of the C record GtkGestureMultiPress.
type GestureMultiPress struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkGestureMultiPress that represents the GestureMultiPress.
func (recv *GestureMultiPress) ToC() unsafe.Pointer {
	return recv.native
}

// GestureMultiPressNewFromC creates a new GestureMultiPress from a pointer to the C GtkGestureMultiPress that represents the GestureMultiPress.
func GestureMultiPressNewFromC(native unsafe.Pointer) *GestureMultiPress {
	return &GestureMultiPress{native: native}
}

// GesturePan is a representation of the C record GtkGesturePan.
type GesturePan struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkGesturePan that represents the GesturePan.
func (recv *GesturePan) ToC() unsafe.Pointer {
	return recv.native
}

// GesturePanNewFromC creates a new GesturePan from a pointer to the C GtkGesturePan that represents the GesturePan.
func GesturePanNewFromC(native unsafe.Pointer) *GesturePan {
	return &GesturePan{native: native}
}

// GestureRotate is a representation of the C record GtkGestureRotate.
type GestureRotate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkGestureRotate that represents the GestureRotate.
func (recv *GestureRotate) ToC() unsafe.Pointer {
	return recv.native
}

// GestureRotateNewFromC creates a new GestureRotate from a pointer to the C GtkGestureRotate that represents the GestureRotate.
func GestureRotateNewFromC(native unsafe.Pointer) *GestureRotate {
	return &GestureRotate{native: native}
}

// GestureSingle is a representation of the C record GtkGestureSingle.
type GestureSingle struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkGestureSingle that represents the GestureSingle.
func (recv *GestureSingle) ToC() unsafe.Pointer {
	return recv.native
}

// GestureSingleNewFromC creates a new GestureSingle from a pointer to the C GtkGestureSingle that represents the GestureSingle.
func GestureSingleNewFromC(native unsafe.Pointer) *GestureSingle {
	return &GestureSingle{native: native}
}

// UNSUPPORTED : GestureStylus : blacklisted

// GestureSwipe is a representation of the C record GtkGestureSwipe.
type GestureSwipe struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkGestureSwipe that represents the GestureSwipe.
func (recv *GestureSwipe) ToC() unsafe.Pointer {
	return recv.native
}

// GestureSwipeNewFromC creates a new GestureSwipe from a pointer to the C GtkGestureSwipe that represents the GestureSwipe.
func GestureSwipeNewFromC(native unsafe.Pointer) *GestureSwipe {
	return &GestureSwipe{native: native}
}

// GestureZoom is a representation of the C record GtkGestureZoom.
type GestureZoom struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkGestureZoom that represents the GestureZoom.
func (recv *GestureZoom) ToC() unsafe.Pointer {
	return recv.native
}

// GestureZoomNewFromC creates a new GestureZoom from a pointer to the C GtkGestureZoom that represents the GestureZoom.
func GestureZoomNewFromC(native unsafe.Pointer) *GestureZoom {
	return &GestureZoom{native: native}
}

// Grid is a representation of the C record GtkGrid.
type Grid struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkGrid that represents the Grid.
func (recv *Grid) ToC() unsafe.Pointer {
	return recv.native
}

// GridNewFromC creates a new Grid from a pointer to the C GtkGrid that represents the Grid.
func GridNewFromC(native unsafe.Pointer) *Grid {
	return &Grid{native: native}
}

// HBox is a representation of the C record GtkHBox.
type HBox struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkHBox that represents the HBox.
func (recv *HBox) ToC() unsafe.Pointer {
	return recv.native
}

// HBoxNewFromC creates a new HBox from a pointer to the C GtkHBox that represents the HBox.
func HBoxNewFromC(native unsafe.Pointer) *HBox {
	return &HBox{native: native}
}

// HButtonBox is a representation of the C record GtkHButtonBox.
type HButtonBox struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkHButtonBox that represents the HButtonBox.
func (recv *HButtonBox) ToC() unsafe.Pointer {
	return recv.native
}

// HButtonBoxNewFromC creates a new HButtonBox from a pointer to the C GtkHButtonBox that represents the HButtonBox.
func HButtonBoxNewFromC(native unsafe.Pointer) *HButtonBox {
	return &HButtonBox{native: native}
}

// HPaned is a representation of the C record GtkHPaned.
type HPaned struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkHPaned that represents the HPaned.
func (recv *HPaned) ToC() unsafe.Pointer {
	return recv.native
}

// HPanedNewFromC creates a new HPaned from a pointer to the C GtkHPaned that represents the HPaned.
func HPanedNewFromC(native unsafe.Pointer) *HPaned {
	return &HPaned{native: native}
}

// HSV is a representation of the C record GtkHSV.
type HSV struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkHSV that represents the HSV.
func (recv *HSV) ToC() unsafe.Pointer {
	return recv.native
}

// HSVNewFromC creates a new HSV from a pointer to the C GtkHSV that represents the HSV.
func HSVNewFromC(native unsafe.Pointer) *HSV {
	return &HSV{native: native}
}

// HScale is a representation of the C record GtkHScale.
type HScale struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkHScale that represents the HScale.
func (recv *HScale) ToC() unsafe.Pointer {
	return recv.native
}

// HScaleNewFromC creates a new HScale from a pointer to the C GtkHScale that represents the HScale.
func HScaleNewFromC(native unsafe.Pointer) *HScale {
	return &HScale{native: native}
}

// HScrollbar is a representation of the C record GtkHScrollbar.
type HScrollbar struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkHScrollbar that represents the HScrollbar.
func (recv *HScrollbar) ToC() unsafe.Pointer {
	return recv.native
}

// HScrollbarNewFromC creates a new HScrollbar from a pointer to the C GtkHScrollbar that represents the HScrollbar.
func HScrollbarNewFromC(native unsafe.Pointer) *HScrollbar {
	return &HScrollbar{native: native}
}

// HSeparator is a representation of the C record GtkHSeparator.
type HSeparator struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkHSeparator that represents the HSeparator.
func (recv *HSeparator) ToC() unsafe.Pointer {
	return recv.native
}

// HSeparatorNewFromC creates a new HSeparator from a pointer to the C GtkHSeparator that represents the HSeparator.
func HSeparatorNewFromC(native unsafe.Pointer) *HSeparator {
	return &HSeparator{native: native}
}

// HandleBox is a representation of the C record GtkHandleBox.
type HandleBox struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkHandleBox that represents the HandleBox.
func (recv *HandleBox) ToC() unsafe.Pointer {
	return recv.native
}

// HandleBoxNewFromC creates a new HandleBox from a pointer to the C GtkHandleBox that represents the HandleBox.
func HandleBoxNewFromC(native unsafe.Pointer) *HandleBox {
	return &HandleBox{native: native}
}

// HeaderBar is a representation of the C record GtkHeaderBar.
type HeaderBar struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkHeaderBar that represents the HeaderBar.
func (recv *HeaderBar) ToC() unsafe.Pointer {
	return recv.native
}

// HeaderBarNewFromC creates a new HeaderBar from a pointer to the C GtkHeaderBar that represents the HeaderBar.
func HeaderBarNewFromC(native unsafe.Pointer) *HeaderBar {
	return &HeaderBar{native: native}
}

// UNSUPPORTED : HeaderBarAccessible : blacklisted

// IMContext is a representation of the C record GtkIMContext.
type IMContext struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkIMContext that represents the IMContext.
func (recv *IMContext) ToC() unsafe.Pointer {
	return recv.native
}

// IMContextNewFromC creates a new IMContext from a pointer to the C GtkIMContext that represents the IMContext.
func IMContextNewFromC(native unsafe.Pointer) *IMContext {
	return &IMContext{native: native}
}

// IMContextSimple is a representation of the C record GtkIMContextSimple.
type IMContextSimple struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkIMContextSimple that represents the IMContextSimple.
func (recv *IMContextSimple) ToC() unsafe.Pointer {
	return recv.native
}

// IMContextSimpleNewFromC creates a new IMContextSimple from a pointer to the C GtkIMContextSimple that represents the IMContextSimple.
func IMContextSimpleNewFromC(native unsafe.Pointer) *IMContextSimple {
	return &IMContextSimple{native: native}
}

// IMMulticontext is a representation of the C record GtkIMMulticontext.
type IMMulticontext struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkIMMulticontext that represents the IMMulticontext.
func (recv *IMMulticontext) ToC() unsafe.Pointer {
	return recv.native
}

// IMMulticontextNewFromC creates a new IMMulticontext from a pointer to the C GtkIMMulticontext that represents the IMMulticontext.
func IMMulticontextNewFromC(native unsafe.Pointer) *IMMulticontext {
	return &IMMulticontext{native: native}
}

// IconFactory is a representation of the C record GtkIconFactory.
type IconFactory struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkIconFactory that represents the IconFactory.
func (recv *IconFactory) ToC() unsafe.Pointer {
	return recv.native
}

// IconFactoryNewFromC creates a new IconFactory from a pointer to the C GtkIconFactory that represents the IconFactory.
func IconFactoryNewFromC(native unsafe.Pointer) *IconFactory {
	return &IconFactory{native: native}
}

// IconInfo is a representation of the C record GtkIconInfo.
type IconInfo struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkIconInfo that represents the IconInfo.
func (recv *IconInfo) ToC() unsafe.Pointer {
	return recv.native
}

// IconInfoNewFromC creates a new IconInfo from a pointer to the C GtkIconInfo that represents the IconInfo.
func IconInfoNewFromC(native unsafe.Pointer) *IconInfo {
	return &IconInfo{native: native}
}

// IconTheme is a representation of the C record GtkIconTheme.
type IconTheme struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkIconTheme that represents the IconTheme.
func (recv *IconTheme) ToC() unsafe.Pointer {
	return recv.native
}

// IconThemeNewFromC creates a new IconTheme from a pointer to the C GtkIconTheme that represents the IconTheme.
func IconThemeNewFromC(native unsafe.Pointer) *IconTheme {
	return &IconTheme{native: native}
}

// IconView is a representation of the C record GtkIconView.
type IconView struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkIconView that represents the IconView.
func (recv *IconView) ToC() unsafe.Pointer {
	return recv.native
}

// IconViewNewFromC creates a new IconView from a pointer to the C GtkIconView that represents the IconView.
func IconViewNewFromC(native unsafe.Pointer) *IconView {
	return &IconView{native: native}
}

// IconViewAccessible is a representation of the C record GtkIconViewAccessible.
type IconViewAccessible struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkIconViewAccessible that represents the IconViewAccessible.
func (recv *IconViewAccessible) ToC() unsafe.Pointer {
	return recv.native
}

// IconViewAccessibleNewFromC creates a new IconViewAccessible from a pointer to the C GtkIconViewAccessible that represents the IconViewAccessible.
func IconViewAccessibleNewFromC(native unsafe.Pointer) *IconViewAccessible {
	return &IconViewAccessible{native: native}
}

// Image is a representation of the C record GtkImage.
type Image struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkImage that represents the Image.
func (recv *Image) ToC() unsafe.Pointer {
	return recv.native
}

// ImageNewFromC creates a new Image from a pointer to the C GtkImage that represents the Image.
func ImageNewFromC(native unsafe.Pointer) *Image {
	return &Image{native: native}
}

// ImageAccessible is a representation of the C record GtkImageAccessible.
type ImageAccessible struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkImageAccessible that represents the ImageAccessible.
func (recv *ImageAccessible) ToC() unsafe.Pointer {
	return recv.native
}

// ImageAccessibleNewFromC creates a new ImageAccessible from a pointer to the C GtkImageAccessible that represents the ImageAccessible.
func ImageAccessibleNewFromC(native unsafe.Pointer) *ImageAccessible {
	return &ImageAccessible{native: native}
}

// ImageCellAccessible is a representation of the C record GtkImageCellAccessible.
type ImageCellAccessible struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkImageCellAccessible that represents the ImageCellAccessible.
func (recv *ImageCellAccessible) ToC() unsafe.Pointer {
	return recv.native
}

// ImageCellAccessibleNewFromC creates a new ImageCellAccessible from a pointer to the C GtkImageCellAccessible that represents the ImageCellAccessible.
func ImageCellAccessibleNewFromC(native unsafe.Pointer) *ImageCellAccessible {
	return &ImageCellAccessible{native: native}
}

// ImageMenuItem is a representation of the C record GtkImageMenuItem.
type ImageMenuItem struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkImageMenuItem that represents the ImageMenuItem.
func (recv *ImageMenuItem) ToC() unsafe.Pointer {
	return recv.native
}

// ImageMenuItemNewFromC creates a new ImageMenuItem from a pointer to the C GtkImageMenuItem that represents the ImageMenuItem.
func ImageMenuItemNewFromC(native unsafe.Pointer) *ImageMenuItem {
	return &ImageMenuItem{native: native}
}

// InfoBar is a representation of the C record GtkInfoBar.
type InfoBar struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkInfoBar that represents the InfoBar.
func (recv *InfoBar) ToC() unsafe.Pointer {
	return recv.native
}

// InfoBarNewFromC creates a new InfoBar from a pointer to the C GtkInfoBar that represents the InfoBar.
func InfoBarNewFromC(native unsafe.Pointer) *InfoBar {
	return &InfoBar{native: native}
}

// Invisible is a representation of the C record GtkInvisible.
type Invisible struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkInvisible that represents the Invisible.
func (recv *Invisible) ToC() unsafe.Pointer {
	return recv.native
}

// InvisibleNewFromC creates a new Invisible from a pointer to the C GtkInvisible that represents the Invisible.
func InvisibleNewFromC(native unsafe.Pointer) *Invisible {
	return &Invisible{native: native}
}

// Label is a representation of the C record GtkLabel.
type Label struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkLabel that represents the Label.
func (recv *Label) ToC() unsafe.Pointer {
	return recv.native
}

// LabelNewFromC creates a new Label from a pointer to the C GtkLabel that represents the Label.
func LabelNewFromC(native unsafe.Pointer) *Label {
	return &Label{native: native}
}

// LabelAccessible is a representation of the C record GtkLabelAccessible.
type LabelAccessible struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkLabelAccessible that represents the LabelAccessible.
func (recv *LabelAccessible) ToC() unsafe.Pointer {
	return recv.native
}

// LabelAccessibleNewFromC creates a new LabelAccessible from a pointer to the C GtkLabelAccessible that represents the LabelAccessible.
func LabelAccessibleNewFromC(native unsafe.Pointer) *LabelAccessible {
	return &LabelAccessible{native: native}
}

// Layout is a representation of the C record GtkLayout.
type Layout struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkLayout that represents the Layout.
func (recv *Layout) ToC() unsafe.Pointer {
	return recv.native
}

// LayoutNewFromC creates a new Layout from a pointer to the C GtkLayout that represents the Layout.
func LayoutNewFromC(native unsafe.Pointer) *Layout {
	return &Layout{native: native}
}

// LevelBar is a representation of the C record GtkLevelBar.
type LevelBar struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkLevelBar that represents the LevelBar.
func (recv *LevelBar) ToC() unsafe.Pointer {
	return recv.native
}

// LevelBarNewFromC creates a new LevelBar from a pointer to the C GtkLevelBar that represents the LevelBar.
func LevelBarNewFromC(native unsafe.Pointer) *LevelBar {
	return &LevelBar{native: native}
}

// LevelBarAccessible is a representation of the C record GtkLevelBarAccessible.
type LevelBarAccessible struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkLevelBarAccessible that represents the LevelBarAccessible.
func (recv *LevelBarAccessible) ToC() unsafe.Pointer {
	return recv.native
}

// LevelBarAccessibleNewFromC creates a new LevelBarAccessible from a pointer to the C GtkLevelBarAccessible that represents the LevelBarAccessible.
func LevelBarAccessibleNewFromC(native unsafe.Pointer) *LevelBarAccessible {
	return &LevelBarAccessible{native: native}
}

// LinkButton is a representation of the C record GtkLinkButton.
type LinkButton struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkLinkButton that represents the LinkButton.
func (recv *LinkButton) ToC() unsafe.Pointer {
	return recv.native
}

// LinkButtonNewFromC creates a new LinkButton from a pointer to the C GtkLinkButton that represents the LinkButton.
func LinkButtonNewFromC(native unsafe.Pointer) *LinkButton {
	return &LinkButton{native: native}
}

// LinkButtonAccessible is a representation of the C record GtkLinkButtonAccessible.
type LinkButtonAccessible struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkLinkButtonAccessible that represents the LinkButtonAccessible.
func (recv *LinkButtonAccessible) ToC() unsafe.Pointer {
	return recv.native
}

// LinkButtonAccessibleNewFromC creates a new LinkButtonAccessible from a pointer to the C GtkLinkButtonAccessible that represents the LinkButtonAccessible.
func LinkButtonAccessibleNewFromC(native unsafe.Pointer) *LinkButtonAccessible {
	return &LinkButtonAccessible{native: native}
}

// ListBox is a representation of the C record GtkListBox.
type ListBox struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkListBox that represents the ListBox.
func (recv *ListBox) ToC() unsafe.Pointer {
	return recv.native
}

// ListBoxNewFromC creates a new ListBox from a pointer to the C GtkListBox that represents the ListBox.
func ListBoxNewFromC(native unsafe.Pointer) *ListBox {
	return &ListBox{native: native}
}

// ListBoxAccessible is a representation of the C record GtkListBoxAccessible.
type ListBoxAccessible struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkListBoxAccessible that represents the ListBoxAccessible.
func (recv *ListBoxAccessible) ToC() unsafe.Pointer {
	return recv.native
}

// ListBoxAccessibleNewFromC creates a new ListBoxAccessible from a pointer to the C GtkListBoxAccessible that represents the ListBoxAccessible.
func ListBoxAccessibleNewFromC(native unsafe.Pointer) *ListBoxAccessible {
	return &ListBoxAccessible{native: native}
}

// ListBoxRow is a representation of the C record GtkListBoxRow.
type ListBoxRow struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkListBoxRow that represents the ListBoxRow.
func (recv *ListBoxRow) ToC() unsafe.Pointer {
	return recv.native
}

// ListBoxRowNewFromC creates a new ListBoxRow from a pointer to the C GtkListBoxRow that represents the ListBoxRow.
func ListBoxRowNewFromC(native unsafe.Pointer) *ListBoxRow {
	return &ListBoxRow{native: native}
}

// ListBoxRowAccessible is a representation of the C record GtkListBoxRowAccessible.
type ListBoxRowAccessible struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkListBoxRowAccessible that represents the ListBoxRowAccessible.
func (recv *ListBoxRowAccessible) ToC() unsafe.Pointer {
	return recv.native
}

// ListBoxRowAccessibleNewFromC creates a new ListBoxRowAccessible from a pointer to the C GtkListBoxRowAccessible that represents the ListBoxRowAccessible.
func ListBoxRowAccessibleNewFromC(native unsafe.Pointer) *ListBoxRowAccessible {
	return &ListBoxRowAccessible{native: native}
}

// ListStore is a representation of the C record GtkListStore.
type ListStore struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkListStore that represents the ListStore.
func (recv *ListStore) ToC() unsafe.Pointer {
	return recv.native
}

// ListStoreNewFromC creates a new ListStore from a pointer to the C GtkListStore that represents the ListStore.
func ListStoreNewFromC(native unsafe.Pointer) *ListStore {
	return &ListStore{native: native}
}

// LockButton is a representation of the C record GtkLockButton.
type LockButton struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkLockButton that represents the LockButton.
func (recv *LockButton) ToC() unsafe.Pointer {
	return recv.native
}

// LockButtonNewFromC creates a new LockButton from a pointer to the C GtkLockButton that represents the LockButton.
func LockButtonNewFromC(native unsafe.Pointer) *LockButton {
	return &LockButton{native: native}
}

// LockButtonAccessible is a representation of the C record GtkLockButtonAccessible.
type LockButtonAccessible struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkLockButtonAccessible that represents the LockButtonAccessible.
func (recv *LockButtonAccessible) ToC() unsafe.Pointer {
	return recv.native
}

// LockButtonAccessibleNewFromC creates a new LockButtonAccessible from a pointer to the C GtkLockButtonAccessible that represents the LockButtonAccessible.
func LockButtonAccessibleNewFromC(native unsafe.Pointer) *LockButtonAccessible {
	return &LockButtonAccessible{native: native}
}

// Menu is a representation of the C record GtkMenu.
type Menu struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkMenu that represents the Menu.
func (recv *Menu) ToC() unsafe.Pointer {
	return recv.native
}

// MenuNewFromC creates a new Menu from a pointer to the C GtkMenu that represents the Menu.
func MenuNewFromC(native unsafe.Pointer) *Menu {
	return &Menu{native: native}
}

// MenuAccessible is a representation of the C record GtkMenuAccessible.
type MenuAccessible struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkMenuAccessible that represents the MenuAccessible.
func (recv *MenuAccessible) ToC() unsafe.Pointer {
	return recv.native
}

// MenuAccessibleNewFromC creates a new MenuAccessible from a pointer to the C GtkMenuAccessible that represents the MenuAccessible.
func MenuAccessibleNewFromC(native unsafe.Pointer) *MenuAccessible {
	return &MenuAccessible{native: native}
}

// MenuBar is a representation of the C record GtkMenuBar.
type MenuBar struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkMenuBar that represents the MenuBar.
func (recv *MenuBar) ToC() unsafe.Pointer {
	return recv.native
}

// MenuBarNewFromC creates a new MenuBar from a pointer to the C GtkMenuBar that represents the MenuBar.
func MenuBarNewFromC(native unsafe.Pointer) *MenuBar {
	return &MenuBar{native: native}
}

// MenuButton is a representation of the C record GtkMenuButton.
type MenuButton struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkMenuButton that represents the MenuButton.
func (recv *MenuButton) ToC() unsafe.Pointer {
	return recv.native
}

// MenuButtonNewFromC creates a new MenuButton from a pointer to the C GtkMenuButton that represents the MenuButton.
func MenuButtonNewFromC(native unsafe.Pointer) *MenuButton {
	return &MenuButton{native: native}
}

// MenuButtonAccessible is a representation of the C record GtkMenuButtonAccessible.
type MenuButtonAccessible struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkMenuButtonAccessible that represents the MenuButtonAccessible.
func (recv *MenuButtonAccessible) ToC() unsafe.Pointer {
	return recv.native
}

// MenuButtonAccessibleNewFromC creates a new MenuButtonAccessible from a pointer to the C GtkMenuButtonAccessible that represents the MenuButtonAccessible.
func MenuButtonAccessibleNewFromC(native unsafe.Pointer) *MenuButtonAccessible {
	return &MenuButtonAccessible{native: native}
}

// MenuItem is a representation of the C record GtkMenuItem.
type MenuItem struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkMenuItem that represents the MenuItem.
func (recv *MenuItem) ToC() unsafe.Pointer {
	return recv.native
}

// MenuItemNewFromC creates a new MenuItem from a pointer to the C GtkMenuItem that represents the MenuItem.
func MenuItemNewFromC(native unsafe.Pointer) *MenuItem {
	return &MenuItem{native: native}
}

// MenuItemAccessible is a representation of the C record GtkMenuItemAccessible.
type MenuItemAccessible struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkMenuItemAccessible that represents the MenuItemAccessible.
func (recv *MenuItemAccessible) ToC() unsafe.Pointer {
	return recv.native
}

// MenuItemAccessibleNewFromC creates a new MenuItemAccessible from a pointer to the C GtkMenuItemAccessible that represents the MenuItemAccessible.
func MenuItemAccessibleNewFromC(native unsafe.Pointer) *MenuItemAccessible {
	return &MenuItemAccessible{native: native}
}

// MenuShell is a representation of the C record GtkMenuShell.
type MenuShell struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkMenuShell that represents the MenuShell.
func (recv *MenuShell) ToC() unsafe.Pointer {
	return recv.native
}

// MenuShellNewFromC creates a new MenuShell from a pointer to the C GtkMenuShell that represents the MenuShell.
func MenuShellNewFromC(native unsafe.Pointer) *MenuShell {
	return &MenuShell{native: native}
}

// MenuShellAccessible is a representation of the C record GtkMenuShellAccessible.
type MenuShellAccessible struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkMenuShellAccessible that represents the MenuShellAccessible.
func (recv *MenuShellAccessible) ToC() unsafe.Pointer {
	return recv.native
}

// MenuShellAccessibleNewFromC creates a new MenuShellAccessible from a pointer to the C GtkMenuShellAccessible that represents the MenuShellAccessible.
func MenuShellAccessibleNewFromC(native unsafe.Pointer) *MenuShellAccessible {
	return &MenuShellAccessible{native: native}
}

// MenuToolButton is a representation of the C record GtkMenuToolButton.
type MenuToolButton struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkMenuToolButton that represents the MenuToolButton.
func (recv *MenuToolButton) ToC() unsafe.Pointer {
	return recv.native
}

// MenuToolButtonNewFromC creates a new MenuToolButton from a pointer to the C GtkMenuToolButton that represents the MenuToolButton.
func MenuToolButtonNewFromC(native unsafe.Pointer) *MenuToolButton {
	return &MenuToolButton{native: native}
}

// MessageDialog is a representation of the C record GtkMessageDialog.
type MessageDialog struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkMessageDialog that represents the MessageDialog.
func (recv *MessageDialog) ToC() unsafe.Pointer {
	return recv.native
}

// MessageDialogNewFromC creates a new MessageDialog from a pointer to the C GtkMessageDialog that represents the MessageDialog.
func MessageDialogNewFromC(native unsafe.Pointer) *MessageDialog {
	return &MessageDialog{native: native}
}

// Misc is a representation of the C record GtkMisc.
type Misc struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkMisc that represents the Misc.
func (recv *Misc) ToC() unsafe.Pointer {
	return recv.native
}

// MiscNewFromC creates a new Misc from a pointer to the C GtkMisc that represents the Misc.
func MiscNewFromC(native unsafe.Pointer) *Misc {
	return &Misc{native: native}
}

// ModelButton is a representation of the C record GtkModelButton.
type ModelButton struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkModelButton that represents the ModelButton.
func (recv *ModelButton) ToC() unsafe.Pointer {
	return recv.native
}

// ModelButtonNewFromC creates a new ModelButton from a pointer to the C GtkModelButton that represents the ModelButton.
func ModelButtonNewFromC(native unsafe.Pointer) *ModelButton {
	return &ModelButton{native: native}
}

// MountOperation is a representation of the C record GtkMountOperation.
type MountOperation struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkMountOperation that represents the MountOperation.
func (recv *MountOperation) ToC() unsafe.Pointer {
	return recv.native
}

// MountOperationNewFromC creates a new MountOperation from a pointer to the C GtkMountOperation that represents the MountOperation.
func MountOperationNewFromC(native unsafe.Pointer) *MountOperation {
	return &MountOperation{native: native}
}

// NativeDialog is a representation of the C record GtkNativeDialog.
type NativeDialog struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkNativeDialog that represents the NativeDialog.
func (recv *NativeDialog) ToC() unsafe.Pointer {
	return recv.native
}

// NativeDialogNewFromC creates a new NativeDialog from a pointer to the C GtkNativeDialog that represents the NativeDialog.
func NativeDialogNewFromC(native unsafe.Pointer) *NativeDialog {
	return &NativeDialog{native: native}
}

// Notebook is a representation of the C record GtkNotebook.
type Notebook struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkNotebook that represents the Notebook.
func (recv *Notebook) ToC() unsafe.Pointer {
	return recv.native
}

// NotebookNewFromC creates a new Notebook from a pointer to the C GtkNotebook that represents the Notebook.
func NotebookNewFromC(native unsafe.Pointer) *Notebook {
	return &Notebook{native: native}
}

// NotebookAccessible is a representation of the C record GtkNotebookAccessible.
type NotebookAccessible struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkNotebookAccessible that represents the NotebookAccessible.
func (recv *NotebookAccessible) ToC() unsafe.Pointer {
	return recv.native
}

// NotebookAccessibleNewFromC creates a new NotebookAccessible from a pointer to the C GtkNotebookAccessible that represents the NotebookAccessible.
func NotebookAccessibleNewFromC(native unsafe.Pointer) *NotebookAccessible {
	return &NotebookAccessible{native: native}
}

// NotebookPageAccessible is a representation of the C record GtkNotebookPageAccessible.
type NotebookPageAccessible struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkNotebookPageAccessible that represents the NotebookPageAccessible.
func (recv *NotebookPageAccessible) ToC() unsafe.Pointer {
	return recv.native
}

// NotebookPageAccessibleNewFromC creates a new NotebookPageAccessible from a pointer to the C GtkNotebookPageAccessible that represents the NotebookPageAccessible.
func NotebookPageAccessibleNewFromC(native unsafe.Pointer) *NotebookPageAccessible {
	return &NotebookPageAccessible{native: native}
}

// NumerableIcon is a representation of the C record GtkNumerableIcon.
type NumerableIcon struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkNumerableIcon that represents the NumerableIcon.
func (recv *NumerableIcon) ToC() unsafe.Pointer {
	return recv.native
}

// NumerableIconNewFromC creates a new NumerableIcon from a pointer to the C GtkNumerableIcon that represents the NumerableIcon.
func NumerableIconNewFromC(native unsafe.Pointer) *NumerableIcon {
	return &NumerableIcon{native: native}
}

// OffscreenWindow is a representation of the C record GtkOffscreenWindow.
type OffscreenWindow struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkOffscreenWindow that represents the OffscreenWindow.
func (recv *OffscreenWindow) ToC() unsafe.Pointer {
	return recv.native
}

// OffscreenWindowNewFromC creates a new OffscreenWindow from a pointer to the C GtkOffscreenWindow that represents the OffscreenWindow.
func OffscreenWindowNewFromC(native unsafe.Pointer) *OffscreenWindow {
	return &OffscreenWindow{native: native}
}

// Overlay is a representation of the C record GtkOverlay.
type Overlay struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkOverlay that represents the Overlay.
func (recv *Overlay) ToC() unsafe.Pointer {
	return recv.native
}

// OverlayNewFromC creates a new Overlay from a pointer to the C GtkOverlay that represents the Overlay.
func OverlayNewFromC(native unsafe.Pointer) *Overlay {
	return &Overlay{native: native}
}

// PadController is a representation of the C record GtkPadController.
type PadController struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkPadController that represents the PadController.
func (recv *PadController) ToC() unsafe.Pointer {
	return recv.native
}

// PadControllerNewFromC creates a new PadController from a pointer to the C GtkPadController that represents the PadController.
func PadControllerNewFromC(native unsafe.Pointer) *PadController {
	return &PadController{native: native}
}

// PageSetup is a representation of the C record GtkPageSetup.
type PageSetup struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkPageSetup that represents the PageSetup.
func (recv *PageSetup) ToC() unsafe.Pointer {
	return recv.native
}

// PageSetupNewFromC creates a new PageSetup from a pointer to the C GtkPageSetup that represents the PageSetup.
func PageSetupNewFromC(native unsafe.Pointer) *PageSetup {
	return &PageSetup{native: native}
}

// Paned is a representation of the C record GtkPaned.
type Paned struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkPaned that represents the Paned.
func (recv *Paned) ToC() unsafe.Pointer {
	return recv.native
}

// PanedNewFromC creates a new Paned from a pointer to the C GtkPaned that represents the Paned.
func PanedNewFromC(native unsafe.Pointer) *Paned {
	return &Paned{native: native}
}

// PanedAccessible is a representation of the C record GtkPanedAccessible.
type PanedAccessible struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkPanedAccessible that represents the PanedAccessible.
func (recv *PanedAccessible) ToC() unsafe.Pointer {
	return recv.native
}

// PanedAccessibleNewFromC creates a new PanedAccessible from a pointer to the C GtkPanedAccessible that represents the PanedAccessible.
func PanedAccessibleNewFromC(native unsafe.Pointer) *PanedAccessible {
	return &PanedAccessible{native: native}
}

// PlacesSidebar is a representation of the C record GtkPlacesSidebar.
type PlacesSidebar struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkPlacesSidebar that represents the PlacesSidebar.
func (recv *PlacesSidebar) ToC() unsafe.Pointer {
	return recv.native
}

// PlacesSidebarNewFromC creates a new PlacesSidebar from a pointer to the C GtkPlacesSidebar that represents the PlacesSidebar.
func PlacesSidebarNewFromC(native unsafe.Pointer) *PlacesSidebar {
	return &PlacesSidebar{native: native}
}

// Plug is a representation of the C record GtkPlug.
type Plug struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkPlug that represents the Plug.
func (recv *Plug) ToC() unsafe.Pointer {
	return recv.native
}

// PlugNewFromC creates a new Plug from a pointer to the C GtkPlug that represents the Plug.
func PlugNewFromC(native unsafe.Pointer) *Plug {
	return &Plug{native: native}
}

// PopoverAccessible is a representation of the C record GtkPopoverAccessible.
type PopoverAccessible struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkPopoverAccessible that represents the PopoverAccessible.
func (recv *PopoverAccessible) ToC() unsafe.Pointer {
	return recv.native
}

// PopoverAccessibleNewFromC creates a new PopoverAccessible from a pointer to the C GtkPopoverAccessible that represents the PopoverAccessible.
func PopoverAccessibleNewFromC(native unsafe.Pointer) *PopoverAccessible {
	return &PopoverAccessible{native: native}
}

// PopoverMenu is a representation of the C record GtkPopoverMenu.
type PopoverMenu struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkPopoverMenu that represents the PopoverMenu.
func (recv *PopoverMenu) ToC() unsafe.Pointer {
	return recv.native
}

// PopoverMenuNewFromC creates a new PopoverMenu from a pointer to the C GtkPopoverMenu that represents the PopoverMenu.
func PopoverMenuNewFromC(native unsafe.Pointer) *PopoverMenu {
	return &PopoverMenu{native: native}
}

// PrintContext is a representation of the C record GtkPrintContext.
type PrintContext struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkPrintContext that represents the PrintContext.
func (recv *PrintContext) ToC() unsafe.Pointer {
	return recv.native
}

// PrintContextNewFromC creates a new PrintContext from a pointer to the C GtkPrintContext that represents the PrintContext.
func PrintContextNewFromC(native unsafe.Pointer) *PrintContext {
	return &PrintContext{native: native}
}

// PrintOperation is a representation of the C record GtkPrintOperation.
type PrintOperation struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkPrintOperation that represents the PrintOperation.
func (recv *PrintOperation) ToC() unsafe.Pointer {
	return recv.native
}

// PrintOperationNewFromC creates a new PrintOperation from a pointer to the C GtkPrintOperation that represents the PrintOperation.
func PrintOperationNewFromC(native unsafe.Pointer) *PrintOperation {
	return &PrintOperation{native: native}
}

// PrintSettings is a representation of the C record GtkPrintSettings.
type PrintSettings struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkPrintSettings that represents the PrintSettings.
func (recv *PrintSettings) ToC() unsafe.Pointer {
	return recv.native
}

// PrintSettingsNewFromC creates a new PrintSettings from a pointer to the C GtkPrintSettings that represents the PrintSettings.
func PrintSettingsNewFromC(native unsafe.Pointer) *PrintSettings {
	return &PrintSettings{native: native}
}

// ProgressBar is a representation of the C record GtkProgressBar.
type ProgressBar struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkProgressBar that represents the ProgressBar.
func (recv *ProgressBar) ToC() unsafe.Pointer {
	return recv.native
}

// ProgressBarNewFromC creates a new ProgressBar from a pointer to the C GtkProgressBar that represents the ProgressBar.
func ProgressBarNewFromC(native unsafe.Pointer) *ProgressBar {
	return &ProgressBar{native: native}
}

// ProgressBarAccessible is a representation of the C record GtkProgressBarAccessible.
type ProgressBarAccessible struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkProgressBarAccessible that represents the ProgressBarAccessible.
func (recv *ProgressBarAccessible) ToC() unsafe.Pointer {
	return recv.native
}

// ProgressBarAccessibleNewFromC creates a new ProgressBarAccessible from a pointer to the C GtkProgressBarAccessible that represents the ProgressBarAccessible.
func ProgressBarAccessibleNewFromC(native unsafe.Pointer) *ProgressBarAccessible {
	return &ProgressBarAccessible{native: native}
}

// RadioAction is a representation of the C record GtkRadioAction.
type RadioAction struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkRadioAction that represents the RadioAction.
func (recv *RadioAction) ToC() unsafe.Pointer {
	return recv.native
}

// RadioActionNewFromC creates a new RadioAction from a pointer to the C GtkRadioAction that represents the RadioAction.
func RadioActionNewFromC(native unsafe.Pointer) *RadioAction {
	return &RadioAction{native: native}
}

// RadioButton is a representation of the C record GtkRadioButton.
type RadioButton struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkRadioButton that represents the RadioButton.
func (recv *RadioButton) ToC() unsafe.Pointer {
	return recv.native
}

// RadioButtonNewFromC creates a new RadioButton from a pointer to the C GtkRadioButton that represents the RadioButton.
func RadioButtonNewFromC(native unsafe.Pointer) *RadioButton {
	return &RadioButton{native: native}
}

// RadioButtonAccessible is a representation of the C record GtkRadioButtonAccessible.
type RadioButtonAccessible struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkRadioButtonAccessible that represents the RadioButtonAccessible.
func (recv *RadioButtonAccessible) ToC() unsafe.Pointer {
	return recv.native
}

// RadioButtonAccessibleNewFromC creates a new RadioButtonAccessible from a pointer to the C GtkRadioButtonAccessible that represents the RadioButtonAccessible.
func RadioButtonAccessibleNewFromC(native unsafe.Pointer) *RadioButtonAccessible {
	return &RadioButtonAccessible{native: native}
}

// RadioMenuItem is a representation of the C record GtkRadioMenuItem.
type RadioMenuItem struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkRadioMenuItem that represents the RadioMenuItem.
func (recv *RadioMenuItem) ToC() unsafe.Pointer {
	return recv.native
}

// RadioMenuItemNewFromC creates a new RadioMenuItem from a pointer to the C GtkRadioMenuItem that represents the RadioMenuItem.
func RadioMenuItemNewFromC(native unsafe.Pointer) *RadioMenuItem {
	return &RadioMenuItem{native: native}
}

// RadioMenuItemAccessible is a representation of the C record GtkRadioMenuItemAccessible.
type RadioMenuItemAccessible struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkRadioMenuItemAccessible that represents the RadioMenuItemAccessible.
func (recv *RadioMenuItemAccessible) ToC() unsafe.Pointer {
	return recv.native
}

// RadioMenuItemAccessibleNewFromC creates a new RadioMenuItemAccessible from a pointer to the C GtkRadioMenuItemAccessible that represents the RadioMenuItemAccessible.
func RadioMenuItemAccessibleNewFromC(native unsafe.Pointer) *RadioMenuItemAccessible {
	return &RadioMenuItemAccessible{native: native}
}

// RadioToolButton is a representation of the C record GtkRadioToolButton.
type RadioToolButton struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkRadioToolButton that represents the RadioToolButton.
func (recv *RadioToolButton) ToC() unsafe.Pointer {
	return recv.native
}

// RadioToolButtonNewFromC creates a new RadioToolButton from a pointer to the C GtkRadioToolButton that represents the RadioToolButton.
func RadioToolButtonNewFromC(native unsafe.Pointer) *RadioToolButton {
	return &RadioToolButton{native: native}
}

// Range is a representation of the C record GtkRange.
type Range struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkRange that represents the Range.
func (recv *Range) ToC() unsafe.Pointer {
	return recv.native
}

// RangeNewFromC creates a new Range from a pointer to the C GtkRange that represents the Range.
func RangeNewFromC(native unsafe.Pointer) *Range {
	return &Range{native: native}
}

// RangeAccessible is a representation of the C record GtkRangeAccessible.
type RangeAccessible struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkRangeAccessible that represents the RangeAccessible.
func (recv *RangeAccessible) ToC() unsafe.Pointer {
	return recv.native
}

// RangeAccessibleNewFromC creates a new RangeAccessible from a pointer to the C GtkRangeAccessible that represents the RangeAccessible.
func RangeAccessibleNewFromC(native unsafe.Pointer) *RangeAccessible {
	return &RangeAccessible{native: native}
}

// RcStyle is a representation of the C record GtkRcStyle.
type RcStyle struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkRcStyle that represents the RcStyle.
func (recv *RcStyle) ToC() unsafe.Pointer {
	return recv.native
}

// RcStyleNewFromC creates a new RcStyle from a pointer to the C GtkRcStyle that represents the RcStyle.
func RcStyleNewFromC(native unsafe.Pointer) *RcStyle {
	return &RcStyle{native: native}
}

// RecentAction is a representation of the C record GtkRecentAction.
type RecentAction struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkRecentAction that represents the RecentAction.
func (recv *RecentAction) ToC() unsafe.Pointer {
	return recv.native
}

// RecentActionNewFromC creates a new RecentAction from a pointer to the C GtkRecentAction that represents the RecentAction.
func RecentActionNewFromC(native unsafe.Pointer) *RecentAction {
	return &RecentAction{native: native}
}

// RecentChooserDialog is a representation of the C record GtkRecentChooserDialog.
type RecentChooserDialog struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkRecentChooserDialog that represents the RecentChooserDialog.
func (recv *RecentChooserDialog) ToC() unsafe.Pointer {
	return recv.native
}

// RecentChooserDialogNewFromC creates a new RecentChooserDialog from a pointer to the C GtkRecentChooserDialog that represents the RecentChooserDialog.
func RecentChooserDialogNewFromC(native unsafe.Pointer) *RecentChooserDialog {
	return &RecentChooserDialog{native: native}
}

// RecentChooserMenu is a representation of the C record GtkRecentChooserMenu.
type RecentChooserMenu struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkRecentChooserMenu that represents the RecentChooserMenu.
func (recv *RecentChooserMenu) ToC() unsafe.Pointer {
	return recv.native
}

// RecentChooserMenuNewFromC creates a new RecentChooserMenu from a pointer to the C GtkRecentChooserMenu that represents the RecentChooserMenu.
func RecentChooserMenuNewFromC(native unsafe.Pointer) *RecentChooserMenu {
	return &RecentChooserMenu{native: native}
}

// RecentChooserWidget is a representation of the C record GtkRecentChooserWidget.
type RecentChooserWidget struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkRecentChooserWidget that represents the RecentChooserWidget.
func (recv *RecentChooserWidget) ToC() unsafe.Pointer {
	return recv.native
}

// RecentChooserWidgetNewFromC creates a new RecentChooserWidget from a pointer to the C GtkRecentChooserWidget that represents the RecentChooserWidget.
func RecentChooserWidgetNewFromC(native unsafe.Pointer) *RecentChooserWidget {
	return &RecentChooserWidget{native: native}
}

// RecentFilter is a representation of the C record GtkRecentFilter.
type RecentFilter struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkRecentFilter that represents the RecentFilter.
func (recv *RecentFilter) ToC() unsafe.Pointer {
	return recv.native
}

// RecentFilterNewFromC creates a new RecentFilter from a pointer to the C GtkRecentFilter that represents the RecentFilter.
func RecentFilterNewFromC(native unsafe.Pointer) *RecentFilter {
	return &RecentFilter{native: native}
}

// RendererCellAccessible is a representation of the C record GtkRendererCellAccessible.
type RendererCellAccessible struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkRendererCellAccessible that represents the RendererCellAccessible.
func (recv *RendererCellAccessible) ToC() unsafe.Pointer {
	return recv.native
}

// RendererCellAccessibleNewFromC creates a new RendererCellAccessible from a pointer to the C GtkRendererCellAccessible that represents the RendererCellAccessible.
func RendererCellAccessibleNewFromC(native unsafe.Pointer) *RendererCellAccessible {
	return &RendererCellAccessible{native: native}
}

// Revealer is a representation of the C record GtkRevealer.
type Revealer struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkRevealer that represents the Revealer.
func (recv *Revealer) ToC() unsafe.Pointer {
	return recv.native
}

// RevealerNewFromC creates a new Revealer from a pointer to the C GtkRevealer that represents the Revealer.
func RevealerNewFromC(native unsafe.Pointer) *Revealer {
	return &Revealer{native: native}
}

// Scale is a representation of the C record GtkScale.
type Scale struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkScale that represents the Scale.
func (recv *Scale) ToC() unsafe.Pointer {
	return recv.native
}

// ScaleNewFromC creates a new Scale from a pointer to the C GtkScale that represents the Scale.
func ScaleNewFromC(native unsafe.Pointer) *Scale {
	return &Scale{native: native}
}

// ScaleAccessible is a representation of the C record GtkScaleAccessible.
type ScaleAccessible struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkScaleAccessible that represents the ScaleAccessible.
func (recv *ScaleAccessible) ToC() unsafe.Pointer {
	return recv.native
}

// ScaleAccessibleNewFromC creates a new ScaleAccessible from a pointer to the C GtkScaleAccessible that represents the ScaleAccessible.
func ScaleAccessibleNewFromC(native unsafe.Pointer) *ScaleAccessible {
	return &ScaleAccessible{native: native}
}

// ScaleButton is a representation of the C record GtkScaleButton.
type ScaleButton struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkScaleButton that represents the ScaleButton.
func (recv *ScaleButton) ToC() unsafe.Pointer {
	return recv.native
}

// ScaleButtonNewFromC creates a new ScaleButton from a pointer to the C GtkScaleButton that represents the ScaleButton.
func ScaleButtonNewFromC(native unsafe.Pointer) *ScaleButton {
	return &ScaleButton{native: native}
}

// ScaleButtonAccessible is a representation of the C record GtkScaleButtonAccessible.
type ScaleButtonAccessible struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkScaleButtonAccessible that represents the ScaleButtonAccessible.
func (recv *ScaleButtonAccessible) ToC() unsafe.Pointer {
	return recv.native
}

// ScaleButtonAccessibleNewFromC creates a new ScaleButtonAccessible from a pointer to the C GtkScaleButtonAccessible that represents the ScaleButtonAccessible.
func ScaleButtonAccessibleNewFromC(native unsafe.Pointer) *ScaleButtonAccessible {
	return &ScaleButtonAccessible{native: native}
}

// Scrollbar is a representation of the C record GtkScrollbar.
type Scrollbar struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkScrollbar that represents the Scrollbar.
func (recv *Scrollbar) ToC() unsafe.Pointer {
	return recv.native
}

// ScrollbarNewFromC creates a new Scrollbar from a pointer to the C GtkScrollbar that represents the Scrollbar.
func ScrollbarNewFromC(native unsafe.Pointer) *Scrollbar {
	return &Scrollbar{native: native}
}

// ScrolledWindow is a representation of the C record GtkScrolledWindow.
type ScrolledWindow struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkScrolledWindow that represents the ScrolledWindow.
func (recv *ScrolledWindow) ToC() unsafe.Pointer {
	return recv.native
}

// ScrolledWindowNewFromC creates a new ScrolledWindow from a pointer to the C GtkScrolledWindow that represents the ScrolledWindow.
func ScrolledWindowNewFromC(native unsafe.Pointer) *ScrolledWindow {
	return &ScrolledWindow{native: native}
}

// ScrolledWindowAccessible is a representation of the C record GtkScrolledWindowAccessible.
type ScrolledWindowAccessible struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkScrolledWindowAccessible that represents the ScrolledWindowAccessible.
func (recv *ScrolledWindowAccessible) ToC() unsafe.Pointer {
	return recv.native
}

// ScrolledWindowAccessibleNewFromC creates a new ScrolledWindowAccessible from a pointer to the C GtkScrolledWindowAccessible that represents the ScrolledWindowAccessible.
func ScrolledWindowAccessibleNewFromC(native unsafe.Pointer) *ScrolledWindowAccessible {
	return &ScrolledWindowAccessible{native: native}
}

// Separator is a representation of the C record GtkSeparator.
type Separator struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkSeparator that represents the Separator.
func (recv *Separator) ToC() unsafe.Pointer {
	return recv.native
}

// SeparatorNewFromC creates a new Separator from a pointer to the C GtkSeparator that represents the Separator.
func SeparatorNewFromC(native unsafe.Pointer) *Separator {
	return &Separator{native: native}
}

// SeparatorMenuItem is a representation of the C record GtkSeparatorMenuItem.
type SeparatorMenuItem struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkSeparatorMenuItem that represents the SeparatorMenuItem.
func (recv *SeparatorMenuItem) ToC() unsafe.Pointer {
	return recv.native
}

// SeparatorMenuItemNewFromC creates a new SeparatorMenuItem from a pointer to the C GtkSeparatorMenuItem that represents the SeparatorMenuItem.
func SeparatorMenuItemNewFromC(native unsafe.Pointer) *SeparatorMenuItem {
	return &SeparatorMenuItem{native: native}
}

// SeparatorToolItem is a representation of the C record GtkSeparatorToolItem.
type SeparatorToolItem struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkSeparatorToolItem that represents the SeparatorToolItem.
func (recv *SeparatorToolItem) ToC() unsafe.Pointer {
	return recv.native
}

// SeparatorToolItemNewFromC creates a new SeparatorToolItem from a pointer to the C GtkSeparatorToolItem that represents the SeparatorToolItem.
func SeparatorToolItemNewFromC(native unsafe.Pointer) *SeparatorToolItem {
	return &SeparatorToolItem{native: native}
}

// Settings is a representation of the C record GtkSettings.
type Settings struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkSettings that represents the Settings.
func (recv *Settings) ToC() unsafe.Pointer {
	return recv.native
}

// SettingsNewFromC creates a new Settings from a pointer to the C GtkSettings that represents the Settings.
func SettingsNewFromC(native unsafe.Pointer) *Settings {
	return &Settings{native: native}
}

// ShortcutLabel is a representation of the C record GtkShortcutLabel.
type ShortcutLabel struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkShortcutLabel that represents the ShortcutLabel.
func (recv *ShortcutLabel) ToC() unsafe.Pointer {
	return recv.native
}

// ShortcutLabelNewFromC creates a new ShortcutLabel from a pointer to the C GtkShortcutLabel that represents the ShortcutLabel.
func ShortcutLabelNewFromC(native unsafe.Pointer) *ShortcutLabel {
	return &ShortcutLabel{native: native}
}

// ShortcutsGroup is a representation of the C record GtkShortcutsGroup.
type ShortcutsGroup struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkShortcutsGroup that represents the ShortcutsGroup.
func (recv *ShortcutsGroup) ToC() unsafe.Pointer {
	return recv.native
}

// ShortcutsGroupNewFromC creates a new ShortcutsGroup from a pointer to the C GtkShortcutsGroup that represents the ShortcutsGroup.
func ShortcutsGroupNewFromC(native unsafe.Pointer) *ShortcutsGroup {
	return &ShortcutsGroup{native: native}
}

// ShortcutsSection is a representation of the C record GtkShortcutsSection.
type ShortcutsSection struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkShortcutsSection that represents the ShortcutsSection.
func (recv *ShortcutsSection) ToC() unsafe.Pointer {
	return recv.native
}

// ShortcutsSectionNewFromC creates a new ShortcutsSection from a pointer to the C GtkShortcutsSection that represents the ShortcutsSection.
func ShortcutsSectionNewFromC(native unsafe.Pointer) *ShortcutsSection {
	return &ShortcutsSection{native: native}
}

// ShortcutsShortcut is a representation of the C record GtkShortcutsShortcut.
type ShortcutsShortcut struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkShortcutsShortcut that represents the ShortcutsShortcut.
func (recv *ShortcutsShortcut) ToC() unsafe.Pointer {
	return recv.native
}

// ShortcutsShortcutNewFromC creates a new ShortcutsShortcut from a pointer to the C GtkShortcutsShortcut that represents the ShortcutsShortcut.
func ShortcutsShortcutNewFromC(native unsafe.Pointer) *ShortcutsShortcut {
	return &ShortcutsShortcut{native: native}
}

// ShortcutsWindow is a representation of the C record GtkShortcutsWindow.
type ShortcutsWindow struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkShortcutsWindow that represents the ShortcutsWindow.
func (recv *ShortcutsWindow) ToC() unsafe.Pointer {
	return recv.native
}

// ShortcutsWindowNewFromC creates a new ShortcutsWindow from a pointer to the C GtkShortcutsWindow that represents the ShortcutsWindow.
func ShortcutsWindowNewFromC(native unsafe.Pointer) *ShortcutsWindow {
	return &ShortcutsWindow{native: native}
}

// SizeGroup is a representation of the C record GtkSizeGroup.
type SizeGroup struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkSizeGroup that represents the SizeGroup.
func (recv *SizeGroup) ToC() unsafe.Pointer {
	return recv.native
}

// SizeGroupNewFromC creates a new SizeGroup from a pointer to the C GtkSizeGroup that represents the SizeGroup.
func SizeGroupNewFromC(native unsafe.Pointer) *SizeGroup {
	return &SizeGroup{native: native}
}

// Socket is a representation of the C record GtkSocket.
type Socket struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkSocket that represents the Socket.
func (recv *Socket) ToC() unsafe.Pointer {
	return recv.native
}

// SocketNewFromC creates a new Socket from a pointer to the C GtkSocket that represents the Socket.
func SocketNewFromC(native unsafe.Pointer) *Socket {
	return &Socket{native: native}
}

// SpinButton is a representation of the C record GtkSpinButton.
type SpinButton struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkSpinButton that represents the SpinButton.
func (recv *SpinButton) ToC() unsafe.Pointer {
	return recv.native
}

// SpinButtonNewFromC creates a new SpinButton from a pointer to the C GtkSpinButton that represents the SpinButton.
func SpinButtonNewFromC(native unsafe.Pointer) *SpinButton {
	return &SpinButton{native: native}
}

// SpinButtonAccessible is a representation of the C record GtkSpinButtonAccessible.
type SpinButtonAccessible struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkSpinButtonAccessible that represents the SpinButtonAccessible.
func (recv *SpinButtonAccessible) ToC() unsafe.Pointer {
	return recv.native
}

// SpinButtonAccessibleNewFromC creates a new SpinButtonAccessible from a pointer to the C GtkSpinButtonAccessible that represents the SpinButtonAccessible.
func SpinButtonAccessibleNewFromC(native unsafe.Pointer) *SpinButtonAccessible {
	return &SpinButtonAccessible{native: native}
}

// Spinner is a representation of the C record GtkSpinner.
type Spinner struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkSpinner that represents the Spinner.
func (recv *Spinner) ToC() unsafe.Pointer {
	return recv.native
}

// SpinnerNewFromC creates a new Spinner from a pointer to the C GtkSpinner that represents the Spinner.
func SpinnerNewFromC(native unsafe.Pointer) *Spinner {
	return &Spinner{native: native}
}

// SpinnerAccessible is a representation of the C record GtkSpinnerAccessible.
type SpinnerAccessible struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkSpinnerAccessible that represents the SpinnerAccessible.
func (recv *SpinnerAccessible) ToC() unsafe.Pointer {
	return recv.native
}

// SpinnerAccessibleNewFromC creates a new SpinnerAccessible from a pointer to the C GtkSpinnerAccessible that represents the SpinnerAccessible.
func SpinnerAccessibleNewFromC(native unsafe.Pointer) *SpinnerAccessible {
	return &SpinnerAccessible{native: native}
}

// Stack is a representation of the C record GtkStack.
type Stack struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkStack that represents the Stack.
func (recv *Stack) ToC() unsafe.Pointer {
	return recv.native
}

// StackNewFromC creates a new Stack from a pointer to the C GtkStack that represents the Stack.
func StackNewFromC(native unsafe.Pointer) *Stack {
	return &Stack{native: native}
}

// UNSUPPORTED : StackAccessible : blacklisted

// StackSwitcher is a representation of the C record GtkStackSwitcher.
type StackSwitcher struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkStackSwitcher that represents the StackSwitcher.
func (recv *StackSwitcher) ToC() unsafe.Pointer {
	return recv.native
}

// StackSwitcherNewFromC creates a new StackSwitcher from a pointer to the C GtkStackSwitcher that represents the StackSwitcher.
func StackSwitcherNewFromC(native unsafe.Pointer) *StackSwitcher {
	return &StackSwitcher{native: native}
}

// StatusIcon is a representation of the C record GtkStatusIcon.
type StatusIcon struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkStatusIcon that represents the StatusIcon.
func (recv *StatusIcon) ToC() unsafe.Pointer {
	return recv.native
}

// StatusIconNewFromC creates a new StatusIcon from a pointer to the C GtkStatusIcon that represents the StatusIcon.
func StatusIconNewFromC(native unsafe.Pointer) *StatusIcon {
	return &StatusIcon{native: native}
}

// Statusbar is a representation of the C record GtkStatusbar.
type Statusbar struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkStatusbar that represents the Statusbar.
func (recv *Statusbar) ToC() unsafe.Pointer {
	return recv.native
}

// StatusbarNewFromC creates a new Statusbar from a pointer to the C GtkStatusbar that represents the Statusbar.
func StatusbarNewFromC(native unsafe.Pointer) *Statusbar {
	return &Statusbar{native: native}
}

// StatusbarAccessible is a representation of the C record GtkStatusbarAccessible.
type StatusbarAccessible struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkStatusbarAccessible that represents the StatusbarAccessible.
func (recv *StatusbarAccessible) ToC() unsafe.Pointer {
	return recv.native
}

// StatusbarAccessibleNewFromC creates a new StatusbarAccessible from a pointer to the C GtkStatusbarAccessible that represents the StatusbarAccessible.
func StatusbarAccessibleNewFromC(native unsafe.Pointer) *StatusbarAccessible {
	return &StatusbarAccessible{native: native}
}

// Style is a representation of the C record GtkStyle.
type Style struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkStyle that represents the Style.
func (recv *Style) ToC() unsafe.Pointer {
	return recv.native
}

// StyleNewFromC creates a new Style from a pointer to the C GtkStyle that represents the Style.
func StyleNewFromC(native unsafe.Pointer) *Style {
	return &Style{native: native}
}

// StyleContext is a representation of the C record GtkStyleContext.
type StyleContext struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkStyleContext that represents the StyleContext.
func (recv *StyleContext) ToC() unsafe.Pointer {
	return recv.native
}

// StyleContextNewFromC creates a new StyleContext from a pointer to the C GtkStyleContext that represents the StyleContext.
func StyleContextNewFromC(native unsafe.Pointer) *StyleContext {
	return &StyleContext{native: native}
}

// StyleProperties is a representation of the C record GtkStyleProperties.
type StyleProperties struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkStyleProperties that represents the StyleProperties.
func (recv *StyleProperties) ToC() unsafe.Pointer {
	return recv.native
}

// StylePropertiesNewFromC creates a new StyleProperties from a pointer to the C GtkStyleProperties that represents the StyleProperties.
func StylePropertiesNewFromC(native unsafe.Pointer) *StyleProperties {
	return &StyleProperties{native: native}
}

// Switch is a representation of the C record GtkSwitch.
type Switch struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkSwitch that represents the Switch.
func (recv *Switch) ToC() unsafe.Pointer {
	return recv.native
}

// SwitchNewFromC creates a new Switch from a pointer to the C GtkSwitch that represents the Switch.
func SwitchNewFromC(native unsafe.Pointer) *Switch {
	return &Switch{native: native}
}

// SwitchAccessible is a representation of the C record GtkSwitchAccessible.
type SwitchAccessible struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkSwitchAccessible that represents the SwitchAccessible.
func (recv *SwitchAccessible) ToC() unsafe.Pointer {
	return recv.native
}

// SwitchAccessibleNewFromC creates a new SwitchAccessible from a pointer to the C GtkSwitchAccessible that represents the SwitchAccessible.
func SwitchAccessibleNewFromC(native unsafe.Pointer) *SwitchAccessible {
	return &SwitchAccessible{native: native}
}

// Table is a representation of the C record GtkTable.
type Table struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkTable that represents the Table.
func (recv *Table) ToC() unsafe.Pointer {
	return recv.native
}

// TableNewFromC creates a new Table from a pointer to the C GtkTable that represents the Table.
func TableNewFromC(native unsafe.Pointer) *Table {
	return &Table{native: native}
}

// TearoffMenuItem is a representation of the C record GtkTearoffMenuItem.
type TearoffMenuItem struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkTearoffMenuItem that represents the TearoffMenuItem.
func (recv *TearoffMenuItem) ToC() unsafe.Pointer {
	return recv.native
}

// TearoffMenuItemNewFromC creates a new TearoffMenuItem from a pointer to the C GtkTearoffMenuItem that represents the TearoffMenuItem.
func TearoffMenuItemNewFromC(native unsafe.Pointer) *TearoffMenuItem {
	return &TearoffMenuItem{native: native}
}

// TextBuffer is a representation of the C record GtkTextBuffer.
type TextBuffer struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkTextBuffer that represents the TextBuffer.
func (recv *TextBuffer) ToC() unsafe.Pointer {
	return recv.native
}

// TextBufferNewFromC creates a new TextBuffer from a pointer to the C GtkTextBuffer that represents the TextBuffer.
func TextBufferNewFromC(native unsafe.Pointer) *TextBuffer {
	return &TextBuffer{native: native}
}

// TextCellAccessible is a representation of the C record GtkTextCellAccessible.
type TextCellAccessible struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkTextCellAccessible that represents the TextCellAccessible.
func (recv *TextCellAccessible) ToC() unsafe.Pointer {
	return recv.native
}

// TextCellAccessibleNewFromC creates a new TextCellAccessible from a pointer to the C GtkTextCellAccessible that represents the TextCellAccessible.
func TextCellAccessibleNewFromC(native unsafe.Pointer) *TextCellAccessible {
	return &TextCellAccessible{native: native}
}

// TextChildAnchor is a representation of the C record GtkTextChildAnchor.
type TextChildAnchor struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkTextChildAnchor that represents the TextChildAnchor.
func (recv *TextChildAnchor) ToC() unsafe.Pointer {
	return recv.native
}

// TextChildAnchorNewFromC creates a new TextChildAnchor from a pointer to the C GtkTextChildAnchor that represents the TextChildAnchor.
func TextChildAnchorNewFromC(native unsafe.Pointer) *TextChildAnchor {
	return &TextChildAnchor{native: native}
}

// TextMark is a representation of the C record GtkTextMark.
type TextMark struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkTextMark that represents the TextMark.
func (recv *TextMark) ToC() unsafe.Pointer {
	return recv.native
}

// TextMarkNewFromC creates a new TextMark from a pointer to the C GtkTextMark that represents the TextMark.
func TextMarkNewFromC(native unsafe.Pointer) *TextMark {
	return &TextMark{native: native}
}

// TextTag is a representation of the C record GtkTextTag.
type TextTag struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkTextTag that represents the TextTag.
func (recv *TextTag) ToC() unsafe.Pointer {
	return recv.native
}

// TextTagNewFromC creates a new TextTag from a pointer to the C GtkTextTag that represents the TextTag.
func TextTagNewFromC(native unsafe.Pointer) *TextTag {
	return &TextTag{native: native}
}

// TextTagTable is a representation of the C record GtkTextTagTable.
type TextTagTable struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkTextTagTable that represents the TextTagTable.
func (recv *TextTagTable) ToC() unsafe.Pointer {
	return recv.native
}

// TextTagTableNewFromC creates a new TextTagTable from a pointer to the C GtkTextTagTable that represents the TextTagTable.
func TextTagTableNewFromC(native unsafe.Pointer) *TextTagTable {
	return &TextTagTable{native: native}
}

// TextView is a representation of the C record GtkTextView.
type TextView struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkTextView that represents the TextView.
func (recv *TextView) ToC() unsafe.Pointer {
	return recv.native
}

// TextViewNewFromC creates a new TextView from a pointer to the C GtkTextView that represents the TextView.
func TextViewNewFromC(native unsafe.Pointer) *TextView {
	return &TextView{native: native}
}

// TextViewAccessible is a representation of the C record GtkTextViewAccessible.
type TextViewAccessible struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkTextViewAccessible that represents the TextViewAccessible.
func (recv *TextViewAccessible) ToC() unsafe.Pointer {
	return recv.native
}

// TextViewAccessibleNewFromC creates a new TextViewAccessible from a pointer to the C GtkTextViewAccessible that represents the TextViewAccessible.
func TextViewAccessibleNewFromC(native unsafe.Pointer) *TextViewAccessible {
	return &TextViewAccessible{native: native}
}

// ThemingEngine is a representation of the C record GtkThemingEngine.
type ThemingEngine struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkThemingEngine that represents the ThemingEngine.
func (recv *ThemingEngine) ToC() unsafe.Pointer {
	return recv.native
}

// ThemingEngineNewFromC creates a new ThemingEngine from a pointer to the C GtkThemingEngine that represents the ThemingEngine.
func ThemingEngineNewFromC(native unsafe.Pointer) *ThemingEngine {
	return &ThemingEngine{native: native}
}

// ToggleAction is a representation of the C record GtkToggleAction.
type ToggleAction struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkToggleAction that represents the ToggleAction.
func (recv *ToggleAction) ToC() unsafe.Pointer {
	return recv.native
}

// ToggleActionNewFromC creates a new ToggleAction from a pointer to the C GtkToggleAction that represents the ToggleAction.
func ToggleActionNewFromC(native unsafe.Pointer) *ToggleAction {
	return &ToggleAction{native: native}
}

// ToggleButton is a representation of the C record GtkToggleButton.
type ToggleButton struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkToggleButton that represents the ToggleButton.
func (recv *ToggleButton) ToC() unsafe.Pointer {
	return recv.native
}

// ToggleButtonNewFromC creates a new ToggleButton from a pointer to the C GtkToggleButton that represents the ToggleButton.
func ToggleButtonNewFromC(native unsafe.Pointer) *ToggleButton {
	return &ToggleButton{native: native}
}

// ToggleButtonAccessible is a representation of the C record GtkToggleButtonAccessible.
type ToggleButtonAccessible struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkToggleButtonAccessible that represents the ToggleButtonAccessible.
func (recv *ToggleButtonAccessible) ToC() unsafe.Pointer {
	return recv.native
}

// ToggleButtonAccessibleNewFromC creates a new ToggleButtonAccessible from a pointer to the C GtkToggleButtonAccessible that represents the ToggleButtonAccessible.
func ToggleButtonAccessibleNewFromC(native unsafe.Pointer) *ToggleButtonAccessible {
	return &ToggleButtonAccessible{native: native}
}

// ToggleToolButton is a representation of the C record GtkToggleToolButton.
type ToggleToolButton struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkToggleToolButton that represents the ToggleToolButton.
func (recv *ToggleToolButton) ToC() unsafe.Pointer {
	return recv.native
}

// ToggleToolButtonNewFromC creates a new ToggleToolButton from a pointer to the C GtkToggleToolButton that represents the ToggleToolButton.
func ToggleToolButtonNewFromC(native unsafe.Pointer) *ToggleToolButton {
	return &ToggleToolButton{native: native}
}

// ToolButton is a representation of the C record GtkToolButton.
type ToolButton struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkToolButton that represents the ToolButton.
func (recv *ToolButton) ToC() unsafe.Pointer {
	return recv.native
}

// ToolButtonNewFromC creates a new ToolButton from a pointer to the C GtkToolButton that represents the ToolButton.
func ToolButtonNewFromC(native unsafe.Pointer) *ToolButton {
	return &ToolButton{native: native}
}

// ToolItem is a representation of the C record GtkToolItem.
type ToolItem struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkToolItem that represents the ToolItem.
func (recv *ToolItem) ToC() unsafe.Pointer {
	return recv.native
}

// ToolItemNewFromC creates a new ToolItem from a pointer to the C GtkToolItem that represents the ToolItem.
func ToolItemNewFromC(native unsafe.Pointer) *ToolItem {
	return &ToolItem{native: native}
}

// Toolbar is a representation of the C record GtkToolbar.
type Toolbar struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkToolbar that represents the Toolbar.
func (recv *Toolbar) ToC() unsafe.Pointer {
	return recv.native
}

// ToolbarNewFromC creates a new Toolbar from a pointer to the C GtkToolbar that represents the Toolbar.
func ToolbarNewFromC(native unsafe.Pointer) *Toolbar {
	return &Toolbar{native: native}
}

// Tooltip is a representation of the C record GtkTooltip.
type Tooltip struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkTooltip that represents the Tooltip.
func (recv *Tooltip) ToC() unsafe.Pointer {
	return recv.native
}

// TooltipNewFromC creates a new Tooltip from a pointer to the C GtkTooltip that represents the Tooltip.
func TooltipNewFromC(native unsafe.Pointer) *Tooltip {
	return &Tooltip{native: native}
}

// ToplevelAccessible is a representation of the C record GtkToplevelAccessible.
type ToplevelAccessible struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkToplevelAccessible that represents the ToplevelAccessible.
func (recv *ToplevelAccessible) ToC() unsafe.Pointer {
	return recv.native
}

// ToplevelAccessibleNewFromC creates a new ToplevelAccessible from a pointer to the C GtkToplevelAccessible that represents the ToplevelAccessible.
func ToplevelAccessibleNewFromC(native unsafe.Pointer) *ToplevelAccessible {
	return &ToplevelAccessible{native: native}
}

// TreeModelFilter is a representation of the C record GtkTreeModelFilter.
type TreeModelFilter struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkTreeModelFilter that represents the TreeModelFilter.
func (recv *TreeModelFilter) ToC() unsafe.Pointer {
	return recv.native
}

// TreeModelFilterNewFromC creates a new TreeModelFilter from a pointer to the C GtkTreeModelFilter that represents the TreeModelFilter.
func TreeModelFilterNewFromC(native unsafe.Pointer) *TreeModelFilter {
	return &TreeModelFilter{native: native}
}

// TreeModelSort is a representation of the C record GtkTreeModelSort.
type TreeModelSort struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkTreeModelSort that represents the TreeModelSort.
func (recv *TreeModelSort) ToC() unsafe.Pointer {
	return recv.native
}

// TreeModelSortNewFromC creates a new TreeModelSort from a pointer to the C GtkTreeModelSort that represents the TreeModelSort.
func TreeModelSortNewFromC(native unsafe.Pointer) *TreeModelSort {
	return &TreeModelSort{native: native}
}

// TreeSelection is a representation of the C record GtkTreeSelection.
type TreeSelection struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkTreeSelection that represents the TreeSelection.
func (recv *TreeSelection) ToC() unsafe.Pointer {
	return recv.native
}

// TreeSelectionNewFromC creates a new TreeSelection from a pointer to the C GtkTreeSelection that represents the TreeSelection.
func TreeSelectionNewFromC(native unsafe.Pointer) *TreeSelection {
	return &TreeSelection{native: native}
}

// TreeStore is a representation of the C record GtkTreeStore.
type TreeStore struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkTreeStore that represents the TreeStore.
func (recv *TreeStore) ToC() unsafe.Pointer {
	return recv.native
}

// TreeStoreNewFromC creates a new TreeStore from a pointer to the C GtkTreeStore that represents the TreeStore.
func TreeStoreNewFromC(native unsafe.Pointer) *TreeStore {
	return &TreeStore{native: native}
}

// TreeView is a representation of the C record GtkTreeView.
type TreeView struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkTreeView that represents the TreeView.
func (recv *TreeView) ToC() unsafe.Pointer {
	return recv.native
}

// TreeViewNewFromC creates a new TreeView from a pointer to the C GtkTreeView that represents the TreeView.
func TreeViewNewFromC(native unsafe.Pointer) *TreeView {
	return &TreeView{native: native}
}

// TreeViewAccessible is a representation of the C record GtkTreeViewAccessible.
type TreeViewAccessible struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkTreeViewAccessible that represents the TreeViewAccessible.
func (recv *TreeViewAccessible) ToC() unsafe.Pointer {
	return recv.native
}

// TreeViewAccessibleNewFromC creates a new TreeViewAccessible from a pointer to the C GtkTreeViewAccessible that represents the TreeViewAccessible.
func TreeViewAccessibleNewFromC(native unsafe.Pointer) *TreeViewAccessible {
	return &TreeViewAccessible{native: native}
}

// TreeViewColumn is a representation of the C record GtkTreeViewColumn.
type TreeViewColumn struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkTreeViewColumn that represents the TreeViewColumn.
func (recv *TreeViewColumn) ToC() unsafe.Pointer {
	return recv.native
}

// TreeViewColumnNewFromC creates a new TreeViewColumn from a pointer to the C GtkTreeViewColumn that represents the TreeViewColumn.
func TreeViewColumnNewFromC(native unsafe.Pointer) *TreeViewColumn {
	return &TreeViewColumn{native: native}
}

// UIManager is a representation of the C record GtkUIManager.
type UIManager struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkUIManager that represents the UIManager.
func (recv *UIManager) ToC() unsafe.Pointer {
	return recv.native
}

// UIManagerNewFromC creates a new UIManager from a pointer to the C GtkUIManager that represents the UIManager.
func UIManagerNewFromC(native unsafe.Pointer) *UIManager {
	return &UIManager{native: native}
}

// VBox is a representation of the C record GtkVBox.
type VBox struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkVBox that represents the VBox.
func (recv *VBox) ToC() unsafe.Pointer {
	return recv.native
}

// VBoxNewFromC creates a new VBox from a pointer to the C GtkVBox that represents the VBox.
func VBoxNewFromC(native unsafe.Pointer) *VBox {
	return &VBox{native: native}
}

// VButtonBox is a representation of the C record GtkVButtonBox.
type VButtonBox struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkVButtonBox that represents the VButtonBox.
func (recv *VButtonBox) ToC() unsafe.Pointer {
	return recv.native
}

// VButtonBoxNewFromC creates a new VButtonBox from a pointer to the C GtkVButtonBox that represents the VButtonBox.
func VButtonBoxNewFromC(native unsafe.Pointer) *VButtonBox {
	return &VButtonBox{native: native}
}

// VPaned is a representation of the C record GtkVPaned.
type VPaned struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkVPaned that represents the VPaned.
func (recv *VPaned) ToC() unsafe.Pointer {
	return recv.native
}

// VPanedNewFromC creates a new VPaned from a pointer to the C GtkVPaned that represents the VPaned.
func VPanedNewFromC(native unsafe.Pointer) *VPaned {
	return &VPaned{native: native}
}

// VScale is a representation of the C record GtkVScale.
type VScale struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkVScale that represents the VScale.
func (recv *VScale) ToC() unsafe.Pointer {
	return recv.native
}

// VScaleNewFromC creates a new VScale from a pointer to the C GtkVScale that represents the VScale.
func VScaleNewFromC(native unsafe.Pointer) *VScale {
	return &VScale{native: native}
}

// VScrollbar is a representation of the C record GtkVScrollbar.
type VScrollbar struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkVScrollbar that represents the VScrollbar.
func (recv *VScrollbar) ToC() unsafe.Pointer {
	return recv.native
}

// VScrollbarNewFromC creates a new VScrollbar from a pointer to the C GtkVScrollbar that represents the VScrollbar.
func VScrollbarNewFromC(native unsafe.Pointer) *VScrollbar {
	return &VScrollbar{native: native}
}

// VSeparator is a representation of the C record GtkVSeparator.
type VSeparator struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkVSeparator that represents the VSeparator.
func (recv *VSeparator) ToC() unsafe.Pointer {
	return recv.native
}

// VSeparatorNewFromC creates a new VSeparator from a pointer to the C GtkVSeparator that represents the VSeparator.
func VSeparatorNewFromC(native unsafe.Pointer) *VSeparator {
	return &VSeparator{native: native}
}

// Viewport is a representation of the C record GtkViewport.
type Viewport struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkViewport that represents the Viewport.
func (recv *Viewport) ToC() unsafe.Pointer {
	return recv.native
}

// ViewportNewFromC creates a new Viewport from a pointer to the C GtkViewport that represents the Viewport.
func ViewportNewFromC(native unsafe.Pointer) *Viewport {
	return &Viewport{native: native}
}

// VolumeButton is a representation of the C record GtkVolumeButton.
type VolumeButton struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkVolumeButton that represents the VolumeButton.
func (recv *VolumeButton) ToC() unsafe.Pointer {
	return recv.native
}

// VolumeButtonNewFromC creates a new VolumeButton from a pointer to the C GtkVolumeButton that represents the VolumeButton.
func VolumeButtonNewFromC(native unsafe.Pointer) *VolumeButton {
	return &VolumeButton{native: native}
}

// Widget is a representation of the C record GtkWidget.
type Widget struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkWidget that represents the Widget.
func (recv *Widget) ToC() unsafe.Pointer {
	return recv.native
}

// WidgetNewFromC creates a new Widget from a pointer to the C GtkWidget that represents the Widget.
func WidgetNewFromC(native unsafe.Pointer) *Widget {
	return &Widget{native: native}
}

// WidgetAccessible is a representation of the C record GtkWidgetAccessible.
type WidgetAccessible struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkWidgetAccessible that represents the WidgetAccessible.
func (recv *WidgetAccessible) ToC() unsafe.Pointer {
	return recv.native
}

// WidgetAccessibleNewFromC creates a new WidgetAccessible from a pointer to the C GtkWidgetAccessible that represents the WidgetAccessible.
func WidgetAccessibleNewFromC(native unsafe.Pointer) *WidgetAccessible {
	return &WidgetAccessible{native: native}
}

// Window is a representation of the C record GtkWindow.
type Window struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkWindow that represents the Window.
func (recv *Window) ToC() unsafe.Pointer {
	return recv.native
}

// WindowNewFromC creates a new Window from a pointer to the C GtkWindow that represents the Window.
func WindowNewFromC(native unsafe.Pointer) *Window {
	return &Window{native: native}
}

// WindowAccessible is a representation of the C record GtkWindowAccessible.
type WindowAccessible struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkWindowAccessible that represents the WindowAccessible.
func (recv *WindowAccessible) ToC() unsafe.Pointer {
	return recv.native
}

// WindowAccessibleNewFromC creates a new WindowAccessible from a pointer to the C GtkWindowAccessible that represents the WindowAccessible.
func WindowAccessibleNewFromC(native unsafe.Pointer) *WindowAccessible {
	return &WindowAccessible{native: native}
}

// WindowGroup is a representation of the C record GtkWindowGroup.
type WindowGroup struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkWindowGroup that represents the WindowGroup.
func (recv *WindowGroup) ToC() unsafe.Pointer {
	return recv.native
}

// WindowGroupNewFromC creates a new WindowGroup from a pointer to the C GtkWindowGroup that represents the WindowGroup.
func WindowGroupNewFromC(native unsafe.Pointer) *WindowGroup {
	return &WindowGroup{native: native}
}

// Activatable is a representation of the C interface GtkActivatable.
type Activatable struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkActivatable that represents the Activatable.
func (recv *Activatable) ToC() unsafe.Pointer {
	return recv.native
}

// ActivatableNewFromC creates a new Activatable from a pointer to the C GtkActivatable that represents the Activatable.
func ActivatableNewFromC(native unsafe.Pointer) *Activatable {
	return &Activatable{native: native}
}

// AppChooser is a representation of the C interface GtkAppChooser.
type AppChooser struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkAppChooser that represents the AppChooser.
func (recv *AppChooser) ToC() unsafe.Pointer {
	return recv.native
}

// AppChooserNewFromC creates a new AppChooser from a pointer to the C GtkAppChooser that represents the AppChooser.
func AppChooserNewFromC(native unsafe.Pointer) *AppChooser {
	return &AppChooser{native: native}
}

// Buildable is a representation of the C interface GtkBuildable.
type Buildable struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkBuildable that represents the Buildable.
func (recv *Buildable) ToC() unsafe.Pointer {
	return recv.native
}

// BuildableNewFromC creates a new Buildable from a pointer to the C GtkBuildable that represents the Buildable.
func BuildableNewFromC(native unsafe.Pointer) *Buildable {
	return &Buildable{native: native}
}

// CellAccessibleParent is a representation of the C interface GtkCellAccessibleParent.
type CellAccessibleParent struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkCellAccessibleParent that represents the CellAccessibleParent.
func (recv *CellAccessibleParent) ToC() unsafe.Pointer {
	return recv.native
}

// CellAccessibleParentNewFromC creates a new CellAccessibleParent from a pointer to the C GtkCellAccessibleParent that represents the CellAccessibleParent.
func CellAccessibleParentNewFromC(native unsafe.Pointer) *CellAccessibleParent {
	return &CellAccessibleParent{native: native}
}

// CellEditable is a representation of the C interface GtkCellEditable.
type CellEditable struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkCellEditable that represents the CellEditable.
func (recv *CellEditable) ToC() unsafe.Pointer {
	return recv.native
}

// CellEditableNewFromC creates a new CellEditable from a pointer to the C GtkCellEditable that represents the CellEditable.
func CellEditableNewFromC(native unsafe.Pointer) *CellEditable {
	return &CellEditable{native: native}
}

// CellLayout is a representation of the C interface GtkCellLayout.
type CellLayout struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkCellLayout that represents the CellLayout.
func (recv *CellLayout) ToC() unsafe.Pointer {
	return recv.native
}

// CellLayoutNewFromC creates a new CellLayout from a pointer to the C GtkCellLayout that represents the CellLayout.
func CellLayoutNewFromC(native unsafe.Pointer) *CellLayout {
	return &CellLayout{native: native}
}

// Editable is a representation of the C interface GtkEditable.
type Editable struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkEditable that represents the Editable.
func (recv *Editable) ToC() unsafe.Pointer {
	return recv.native
}

// EditableNewFromC creates a new Editable from a pointer to the C GtkEditable that represents the Editable.
func EditableNewFromC(native unsafe.Pointer) *Editable {
	return &Editable{native: native}
}

// FileChooser is a representation of the C interface GtkFileChooser.
type FileChooser struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkFileChooser that represents the FileChooser.
func (recv *FileChooser) ToC() unsafe.Pointer {
	return recv.native
}

// FileChooserNewFromC creates a new FileChooser from a pointer to the C GtkFileChooser that represents the FileChooser.
func FileChooserNewFromC(native unsafe.Pointer) *FileChooser {
	return &FileChooser{native: native}
}

// FontChooser is a representation of the C interface GtkFontChooser.
type FontChooser struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkFontChooser that represents the FontChooser.
func (recv *FontChooser) ToC() unsafe.Pointer {
	return recv.native
}

// FontChooserNewFromC creates a new FontChooser from a pointer to the C GtkFontChooser that represents the FontChooser.
func FontChooserNewFromC(native unsafe.Pointer) *FontChooser {
	return &FontChooser{native: native}
}

// Orientable is a representation of the C interface GtkOrientable.
type Orientable struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkOrientable that represents the Orientable.
func (recv *Orientable) ToC() unsafe.Pointer {
	return recv.native
}

// OrientableNewFromC creates a new Orientable from a pointer to the C GtkOrientable that represents the Orientable.
func OrientableNewFromC(native unsafe.Pointer) *Orientable {
	return &Orientable{native: native}
}

// PrintOperationPreview is a representation of the C interface GtkPrintOperationPreview.
type PrintOperationPreview struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkPrintOperationPreview that represents the PrintOperationPreview.
func (recv *PrintOperationPreview) ToC() unsafe.Pointer {
	return recv.native
}

// PrintOperationPreviewNewFromC creates a new PrintOperationPreview from a pointer to the C GtkPrintOperationPreview that represents the PrintOperationPreview.
func PrintOperationPreviewNewFromC(native unsafe.Pointer) *PrintOperationPreview {
	return &PrintOperationPreview{native: native}
}

// RecentChooser is a representation of the C interface GtkRecentChooser.
type RecentChooser struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkRecentChooser that represents the RecentChooser.
func (recv *RecentChooser) ToC() unsafe.Pointer {
	return recv.native
}

// RecentChooserNewFromC creates a new RecentChooser from a pointer to the C GtkRecentChooser that represents the RecentChooser.
func RecentChooserNewFromC(native unsafe.Pointer) *RecentChooser {
	return &RecentChooser{native: native}
}

// Scrollable is a representation of the C interface GtkScrollable.
type Scrollable struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkScrollable that represents the Scrollable.
func (recv *Scrollable) ToC() unsafe.Pointer {
	return recv.native
}

// ScrollableNewFromC creates a new Scrollable from a pointer to the C GtkScrollable that represents the Scrollable.
func ScrollableNewFromC(native unsafe.Pointer) *Scrollable {
	return &Scrollable{native: native}
}

// StyleProvider is a representation of the C interface GtkStyleProvider.
type StyleProvider struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkStyleProvider that represents the StyleProvider.
func (recv *StyleProvider) ToC() unsafe.Pointer {
	return recv.native
}

// StyleProviderNewFromC creates a new StyleProvider from a pointer to the C GtkStyleProvider that represents the StyleProvider.
func StyleProviderNewFromC(native unsafe.Pointer) *StyleProvider {
	return &StyleProvider{native: native}
}

// ToolShell is a representation of the C interface GtkToolShell.
type ToolShell struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkToolShell that represents the ToolShell.
func (recv *ToolShell) ToC() unsafe.Pointer {
	return recv.native
}

// ToolShellNewFromC creates a new ToolShell from a pointer to the C GtkToolShell that represents the ToolShell.
func ToolShellNewFromC(native unsafe.Pointer) *ToolShell {
	return &ToolShell{native: native}
}

// TreeDragDest is a representation of the C interface GtkTreeDragDest.
type TreeDragDest struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkTreeDragDest that represents the TreeDragDest.
func (recv *TreeDragDest) ToC() unsafe.Pointer {
	return recv.native
}

// TreeDragDestNewFromC creates a new TreeDragDest from a pointer to the C GtkTreeDragDest that represents the TreeDragDest.
func TreeDragDestNewFromC(native unsafe.Pointer) *TreeDragDest {
	return &TreeDragDest{native: native}
}

// TreeDragSource is a representation of the C interface GtkTreeDragSource.
type TreeDragSource struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkTreeDragSource that represents the TreeDragSource.
func (recv *TreeDragSource) ToC() unsafe.Pointer {
	return recv.native
}

// TreeDragSourceNewFromC creates a new TreeDragSource from a pointer to the C GtkTreeDragSource that represents the TreeDragSource.
func TreeDragSourceNewFromC(native unsafe.Pointer) *TreeDragSource {
	return &TreeDragSource{native: native}
}

// TreeModel is a representation of the C interface GtkTreeModel.
type TreeModel struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkTreeModel that represents the TreeModel.
func (recv *TreeModel) ToC() unsafe.Pointer {
	return recv.native
}

// TreeModelNewFromC creates a new TreeModel from a pointer to the C GtkTreeModel that represents the TreeModel.
func TreeModelNewFromC(native unsafe.Pointer) *TreeModel {
	return &TreeModel{native: native}
}

// TreeSortable is a representation of the C interface GtkTreeSortable.
type TreeSortable struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkTreeSortable that represents the TreeSortable.
func (recv *TreeSortable) ToC() unsafe.Pointer {
	return recv.native
}

// TreeSortableNewFromC creates a new TreeSortable from a pointer to the C GtkTreeSortable that represents the TreeSortable.
func TreeSortableNewFromC(native unsafe.Pointer) *TreeSortable {
	return &TreeSortable{native: native}
}
