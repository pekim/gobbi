// Code generated - DO NOT EDIT.
// +build !gtk_2.2,!gtk_2.4,!gtk_2.6,!gtk_2.8,!gtk_2.10,!gtk_2.12,!gtk_2.14,!gtk_2.16,!gtk_2.18,!gtk_2.20,!gtk_2.22,!gtk_2.24,!gtk_3.0,!gtk_3.2,!gtk_3.4,!gtk_3.6,!gtk_3.8,!gtk_3.10,!gtk_3.12,!gtk_3.14,!gtk_3.16,!gtk_3.18,!gtk_3.20,!gtk_3.22,!gtk_3.22.6,!gtk_3.22.26,!gtk_3.22.29,!gtk_3.24

package gtk

import (
	gdk "github.com/pekim/gobbi/lib/gdk"
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

// AboutDialogClass is a representation of the C record GtkAboutDialogClass.
type AboutDialogClass struct {
	native unsafe.Pointer
}

// AboutDialogPrivate is a representation of the C record GtkAboutDialogPrivate.
type AboutDialogPrivate struct {
	native unsafe.Pointer
}

// AccelGroupClass is a representation of the C record GtkAccelGroupClass.
type AccelGroupClass struct {
	native unsafe.Pointer
}

// AccelGroupEntry is a representation of the C record GtkAccelGroupEntry.
type AccelGroupEntry struct {
	native unsafe.Pointer
}

// AccelGroupPrivate is a representation of the C record GtkAccelGroupPrivate.
type AccelGroupPrivate struct {
	native unsafe.Pointer
}

// AccelKey is a representation of the C record GtkAccelKey.
type AccelKey struct {
	native unsafe.Pointer
}

// AccelLabelClass is a representation of the C record GtkAccelLabelClass.
type AccelLabelClass struct {
	native unsafe.Pointer
}

// AccelLabelPrivate is a representation of the C record GtkAccelLabelPrivate.
type AccelLabelPrivate struct {
	native unsafe.Pointer
}

// AccelMapClass is a representation of the C record GtkAccelMapClass.
type AccelMapClass struct {
	native unsafe.Pointer
}

// AccessibleClass is a representation of the C record GtkAccessibleClass.
type AccessibleClass struct {
	native unsafe.Pointer
}

// AccessiblePrivate is a representation of the C record GtkAccessiblePrivate.
type AccessiblePrivate struct {
	native unsafe.Pointer
}

// ActionBarClass is a representation of the C record GtkActionBarClass.
type ActionBarClass struct {
	native unsafe.Pointer
}

// ActionBarPrivate is a representation of the C record GtkActionBarPrivate.
type ActionBarPrivate struct {
	native unsafe.Pointer
}

// ActionClass is a representation of the C record GtkActionClass.
type ActionClass struct {
	native unsafe.Pointer
}

// ActionEntry is a representation of the C record GtkActionEntry.
type ActionEntry struct {
	native unsafe.Pointer
}

// ActionGroupClass is a representation of the C record GtkActionGroupClass.
type ActionGroupClass struct {
	native unsafe.Pointer
}

// ActionGroupPrivate is a representation of the C record GtkActionGroupPrivate.
type ActionGroupPrivate struct {
	native unsafe.Pointer
}

// ActionPrivate is a representation of the C record GtkActionPrivate.
type ActionPrivate struct {
	native unsafe.Pointer
}

// ActionableInterface is a representation of the C record GtkActionableInterface.
type ActionableInterface struct {
	native unsafe.Pointer
}

// AdjustmentClass is a representation of the C record GtkAdjustmentClass.
type AdjustmentClass struct {
	native unsafe.Pointer
}

// AdjustmentPrivate is a representation of the C record GtkAdjustmentPrivate.
type AdjustmentPrivate struct {
	native unsafe.Pointer
}

// AlignmentClass is a representation of the C record GtkAlignmentClass.
type AlignmentClass struct {
	native unsafe.Pointer
}

// AlignmentPrivate is a representation of the C record GtkAlignmentPrivate.
type AlignmentPrivate struct {
	native unsafe.Pointer
}

// AppChooserButtonClass is a representation of the C record GtkAppChooserButtonClass.
type AppChooserButtonClass struct {
	native unsafe.Pointer
}

// AppChooserButtonPrivate is a representation of the C record GtkAppChooserButtonPrivate.
type AppChooserButtonPrivate struct {
	native unsafe.Pointer
}

// AppChooserDialogClass is a representation of the C record GtkAppChooserDialogClass.
type AppChooserDialogClass struct {
	native unsafe.Pointer
}

// AppChooserDialogPrivate is a representation of the C record GtkAppChooserDialogPrivate.
type AppChooserDialogPrivate struct {
	native unsafe.Pointer
}

// AppChooserWidgetClass is a representation of the C record GtkAppChooserWidgetClass.
type AppChooserWidgetClass struct {
	native unsafe.Pointer
}

// AppChooserWidgetPrivate is a representation of the C record GtkAppChooserWidgetPrivate.
type AppChooserWidgetPrivate struct {
	native unsafe.Pointer
}

// ApplicationClass is a representation of the C record GtkApplicationClass.
type ApplicationClass struct {
	native unsafe.Pointer
}

// ApplicationPrivate is a representation of the C record GtkApplicationPrivate.
type ApplicationPrivate struct {
	native unsafe.Pointer
}

// ApplicationWindowClass is a representation of the C record GtkApplicationWindowClass.
type ApplicationWindowClass struct {
	native unsafe.Pointer
}

// ApplicationWindowPrivate is a representation of the C record GtkApplicationWindowPrivate.
type ApplicationWindowPrivate struct {
	native unsafe.Pointer
}

// ArrowAccessibleClass is a representation of the C record GtkArrowAccessibleClass.
type ArrowAccessibleClass struct {
	native unsafe.Pointer
}

// ArrowAccessiblePrivate is a representation of the C record GtkArrowAccessiblePrivate.
type ArrowAccessiblePrivate struct {
	native unsafe.Pointer
}

// ArrowClass is a representation of the C record GtkArrowClass.
type ArrowClass struct {
	native unsafe.Pointer
}

// ArrowPrivate is a representation of the C record GtkArrowPrivate.
type ArrowPrivate struct {
	native unsafe.Pointer
}

// AspectFrameClass is a representation of the C record GtkAspectFrameClass.
type AspectFrameClass struct {
	native unsafe.Pointer
}

// AspectFramePrivate is a representation of the C record GtkAspectFramePrivate.
type AspectFramePrivate struct {
	native unsafe.Pointer
}

// AssistantClass is a representation of the C record GtkAssistantClass.
type AssistantClass struct {
	native unsafe.Pointer
}

// AssistantPrivate is a representation of the C record GtkAssistantPrivate.
type AssistantPrivate struct {
	native unsafe.Pointer
}

// BinClass is a representation of the C record GtkBinClass.
type BinClass struct {
	native unsafe.Pointer
}

// BinPrivate is a representation of the C record GtkBinPrivate.
type BinPrivate struct {
	native unsafe.Pointer
}

// BindingArg is a representation of the C record GtkBindingArg.
type BindingArg struct {
	native unsafe.Pointer
}

// BindingEntry is a representation of the C record GtkBindingEntry.
type BindingEntry struct {
	native unsafe.Pointer
}

// BindingSet is a representation of the C record GtkBindingSet.
type BindingSet struct {
	native unsafe.Pointer
}

// BindingSignal is a representation of the C record GtkBindingSignal.
type BindingSignal struct {
	native unsafe.Pointer
}

// BooleanCellAccessibleClass is a representation of the C record GtkBooleanCellAccessibleClass.
type BooleanCellAccessibleClass struct {
	native unsafe.Pointer
}

// BooleanCellAccessiblePrivate is a representation of the C record GtkBooleanCellAccessiblePrivate.
type BooleanCellAccessiblePrivate struct {
	native unsafe.Pointer
}

// Border is a representation of the C record GtkBorder.
type Border struct {
	native unsafe.Pointer
}

// BoxClass is a representation of the C record GtkBoxClass.
type BoxClass struct {
	native unsafe.Pointer
}

// BoxPrivate is a representation of the C record GtkBoxPrivate.
type BoxPrivate struct {
	native unsafe.Pointer
}

// BuildableIface is a representation of the C record GtkBuildableIface.
type BuildableIface struct {
	native unsafe.Pointer
}

// BuilderClass is a representation of the C record GtkBuilderClass.
type BuilderClass struct {
	native unsafe.Pointer
}

// BuilderPrivate is a representation of the C record GtkBuilderPrivate.
type BuilderPrivate struct {
	native unsafe.Pointer
}

// ButtonAccessibleClass is a representation of the C record GtkButtonAccessibleClass.
type ButtonAccessibleClass struct {
	native unsafe.Pointer
}

// ButtonAccessiblePrivate is a representation of the C record GtkButtonAccessiblePrivate.
type ButtonAccessiblePrivate struct {
	native unsafe.Pointer
}

// ButtonBoxClass is a representation of the C record GtkButtonBoxClass.
type ButtonBoxClass struct {
	native unsafe.Pointer
}

// ButtonBoxPrivate is a representation of the C record GtkButtonBoxPrivate.
type ButtonBoxPrivate struct {
	native unsafe.Pointer
}

// ButtonClass is a representation of the C record GtkButtonClass.
type ButtonClass struct {
	native unsafe.Pointer
}

// ButtonPrivate is a representation of the C record GtkButtonPrivate.
type ButtonPrivate struct {
	native unsafe.Pointer
}

// CalendarClass is a representation of the C record GtkCalendarClass.
type CalendarClass struct {
	native unsafe.Pointer
}

// CalendarPrivate is a representation of the C record GtkCalendarPrivate.
type CalendarPrivate struct {
	native unsafe.Pointer
}

// CellAccessibleClass is a representation of the C record GtkCellAccessibleClass.
type CellAccessibleClass struct {
	native unsafe.Pointer
}

// CellAccessibleParentIface is a representation of the C record GtkCellAccessibleParentIface.
type CellAccessibleParentIface struct {
	native unsafe.Pointer
}

// CellAccessiblePrivate is a representation of the C record GtkCellAccessiblePrivate.
type CellAccessiblePrivate struct {
	native unsafe.Pointer
}

// CellAreaBoxClass is a representation of the C record GtkCellAreaBoxClass.
type CellAreaBoxClass struct {
	native unsafe.Pointer
}

// CellAreaBoxPrivate is a representation of the C record GtkCellAreaBoxPrivate.
type CellAreaBoxPrivate struct {
	native unsafe.Pointer
}

// CellAreaClass is a representation of the C record GtkCellAreaClass.
type CellAreaClass struct {
	native unsafe.Pointer
}

// CellAreaContextClass is a representation of the C record GtkCellAreaContextClass.
type CellAreaContextClass struct {
	native unsafe.Pointer
}

// CellAreaContextPrivate is a representation of the C record GtkCellAreaContextPrivate.
type CellAreaContextPrivate struct {
	native unsafe.Pointer
}

// CellAreaPrivate is a representation of the C record GtkCellAreaPrivate.
type CellAreaPrivate struct {
	native unsafe.Pointer
}

// CellEditableIface is a representation of the C record GtkCellEditableIface.
type CellEditableIface struct {
	native unsafe.Pointer
}

// CellLayoutIface is a representation of the C record GtkCellLayoutIface.
type CellLayoutIface struct {
	native unsafe.Pointer
}

// CellRendererAccelClass is a representation of the C record GtkCellRendererAccelClass.
type CellRendererAccelClass struct {
	native unsafe.Pointer
}

// CellRendererAccelPrivate is a representation of the C record GtkCellRendererAccelPrivate.
type CellRendererAccelPrivate struct {
	native unsafe.Pointer
}

// CellRendererClass is a representation of the C record GtkCellRendererClass.
type CellRendererClass struct {
	native unsafe.Pointer
}

// CellRendererClassPrivate is a representation of the C record GtkCellRendererClassPrivate.
type CellRendererClassPrivate struct {
	native unsafe.Pointer
}

// CellRendererComboClass is a representation of the C record GtkCellRendererComboClass.
type CellRendererComboClass struct {
	native unsafe.Pointer
}

// CellRendererComboPrivate is a representation of the C record GtkCellRendererComboPrivate.
type CellRendererComboPrivate struct {
	native unsafe.Pointer
}

// CellRendererPixbufClass is a representation of the C record GtkCellRendererPixbufClass.
type CellRendererPixbufClass struct {
	native unsafe.Pointer
}

// CellRendererPixbufPrivate is a representation of the C record GtkCellRendererPixbufPrivate.
type CellRendererPixbufPrivate struct {
	native unsafe.Pointer
}

// CellRendererPrivate is a representation of the C record GtkCellRendererPrivate.
type CellRendererPrivate struct {
	native unsafe.Pointer
}

// CellRendererProgressClass is a representation of the C record GtkCellRendererProgressClass.
type CellRendererProgressClass struct {
	native unsafe.Pointer
}

// CellRendererProgressPrivate is a representation of the C record GtkCellRendererProgressPrivate.
type CellRendererProgressPrivate struct {
	native unsafe.Pointer
}

// CellRendererSpinClass is a representation of the C record GtkCellRendererSpinClass.
type CellRendererSpinClass struct {
	native unsafe.Pointer
}

// CellRendererSpinPrivate is a representation of the C record GtkCellRendererSpinPrivate.
type CellRendererSpinPrivate struct {
	native unsafe.Pointer
}

// CellRendererSpinnerClass is a representation of the C record GtkCellRendererSpinnerClass.
type CellRendererSpinnerClass struct {
	native unsafe.Pointer
}

// CellRendererSpinnerPrivate is a representation of the C record GtkCellRendererSpinnerPrivate.
type CellRendererSpinnerPrivate struct {
	native unsafe.Pointer
}

// CellRendererTextClass is a representation of the C record GtkCellRendererTextClass.
type CellRendererTextClass struct {
	native unsafe.Pointer
}

// CellRendererTextPrivate is a representation of the C record GtkCellRendererTextPrivate.
type CellRendererTextPrivate struct {
	native unsafe.Pointer
}

// CellRendererToggleClass is a representation of the C record GtkCellRendererToggleClass.
type CellRendererToggleClass struct {
	native unsafe.Pointer
}

// CellRendererTogglePrivate is a representation of the C record GtkCellRendererTogglePrivate.
type CellRendererTogglePrivate struct {
	native unsafe.Pointer
}

// CellViewClass is a representation of the C record GtkCellViewClass.
type CellViewClass struct {
	native unsafe.Pointer
}

// CellViewPrivate is a representation of the C record GtkCellViewPrivate.
type CellViewPrivate struct {
	native unsafe.Pointer
}

// CheckButtonClass is a representation of the C record GtkCheckButtonClass.
type CheckButtonClass struct {
	native unsafe.Pointer
}

// CheckMenuItemAccessibleClass is a representation of the C record GtkCheckMenuItemAccessibleClass.
type CheckMenuItemAccessibleClass struct {
	native unsafe.Pointer
}

// CheckMenuItemAccessiblePrivate is a representation of the C record GtkCheckMenuItemAccessiblePrivate.
type CheckMenuItemAccessiblePrivate struct {
	native unsafe.Pointer
}

// CheckMenuItemClass is a representation of the C record GtkCheckMenuItemClass.
type CheckMenuItemClass struct {
	native unsafe.Pointer
}

// CheckMenuItemPrivate is a representation of the C record GtkCheckMenuItemPrivate.
type CheckMenuItemPrivate struct {
	native unsafe.Pointer
}

// ColorButtonClass is a representation of the C record GtkColorButtonClass.
type ColorButtonClass struct {
	native unsafe.Pointer
}

// ColorButtonPrivate is a representation of the C record GtkColorButtonPrivate.
type ColorButtonPrivate struct {
	native unsafe.Pointer
}

// ColorChooserDialogClass is a representation of the C record GtkColorChooserDialogClass.
type ColorChooserDialogClass struct {
	native unsafe.Pointer
}

// ColorChooserDialogPrivate is a representation of the C record GtkColorChooserDialogPrivate.
type ColorChooserDialogPrivate struct {
	native unsafe.Pointer
}

// ColorChooserInterface is a representation of the C record GtkColorChooserInterface.
type ColorChooserInterface struct {
	native unsafe.Pointer
}

// ColorChooserWidgetClass is a representation of the C record GtkColorChooserWidgetClass.
type ColorChooserWidgetClass struct {
	native unsafe.Pointer
}

// ColorChooserWidgetPrivate is a representation of the C record GtkColorChooserWidgetPrivate.
type ColorChooserWidgetPrivate struct {
	native unsafe.Pointer
}

// ColorSelectionClass is a representation of the C record GtkColorSelectionClass.
type ColorSelectionClass struct {
	native unsafe.Pointer
}

// ColorSelectionDialogClass is a representation of the C record GtkColorSelectionDialogClass.
type ColorSelectionDialogClass struct {
	native unsafe.Pointer
}

// ColorSelectionDialogPrivate is a representation of the C record GtkColorSelectionDialogPrivate.
type ColorSelectionDialogPrivate struct {
	native unsafe.Pointer
}

// ColorSelectionPrivate is a representation of the C record GtkColorSelectionPrivate.
type ColorSelectionPrivate struct {
	native unsafe.Pointer
}

// ComboBoxAccessibleClass is a representation of the C record GtkComboBoxAccessibleClass.
type ComboBoxAccessibleClass struct {
	native unsafe.Pointer
}

// ComboBoxAccessiblePrivate is a representation of the C record GtkComboBoxAccessiblePrivate.
type ComboBoxAccessiblePrivate struct {
	native unsafe.Pointer
}

// ComboBoxClass is a representation of the C record GtkComboBoxClass.
type ComboBoxClass struct {
	native unsafe.Pointer
}

// ComboBoxPrivate is a representation of the C record GtkComboBoxPrivate.
type ComboBoxPrivate struct {
	native unsafe.Pointer
}

// ComboBoxTextClass is a representation of the C record GtkComboBoxTextClass.
type ComboBoxTextClass struct {
	native unsafe.Pointer
}

// ComboBoxTextPrivate is a representation of the C record GtkComboBoxTextPrivate.
type ComboBoxTextPrivate struct {
	native unsafe.Pointer
}

// ContainerAccessibleClass is a representation of the C record GtkContainerAccessibleClass.
type ContainerAccessibleClass struct {
	native unsafe.Pointer
}

// ContainerAccessiblePrivate is a representation of the C record GtkContainerAccessiblePrivate.
type ContainerAccessiblePrivate struct {
	native unsafe.Pointer
}

// ContainerCellAccessibleClass is a representation of the C record GtkContainerCellAccessibleClass.
type ContainerCellAccessibleClass struct {
	native unsafe.Pointer
}

// ContainerCellAccessiblePrivate is a representation of the C record GtkContainerCellAccessiblePrivate.
type ContainerCellAccessiblePrivate struct {
	native unsafe.Pointer
}

// ContainerClass is a representation of the C record GtkContainerClass.
type ContainerClass struct {
	native unsafe.Pointer
}

// ContainerPrivate is a representation of the C record GtkContainerPrivate.
type ContainerPrivate struct {
	native unsafe.Pointer
}

// CssProviderClass is a representation of the C record GtkCssProviderClass.
type CssProviderClass struct {
	native unsafe.Pointer
}

// CssProviderPrivate is a representation of the C record GtkCssProviderPrivate.
type CssProviderPrivate struct {
	native unsafe.Pointer
}

// DialogClass is a representation of the C record GtkDialogClass.
type DialogClass struct {
	native unsafe.Pointer
}

// DialogPrivate is a representation of the C record GtkDialogPrivate.
type DialogPrivate struct {
	native unsafe.Pointer
}

// DrawingAreaClass is a representation of the C record GtkDrawingAreaClass.
type DrawingAreaClass struct {
	native unsafe.Pointer
}

// EditableInterface is a representation of the C record GtkEditableInterface.
type EditableInterface struct {
	native unsafe.Pointer
}

// EntryAccessibleClass is a representation of the C record GtkEntryAccessibleClass.
type EntryAccessibleClass struct {
	native unsafe.Pointer
}

// EntryAccessiblePrivate is a representation of the C record GtkEntryAccessiblePrivate.
type EntryAccessiblePrivate struct {
	native unsafe.Pointer
}

// EntryBufferClass is a representation of the C record GtkEntryBufferClass.
type EntryBufferClass struct {
	native unsafe.Pointer
}

// EntryBufferPrivate is a representation of the C record GtkEntryBufferPrivate.
type EntryBufferPrivate struct {
	native unsafe.Pointer
}

// EntryClass is a representation of the C record GtkEntryClass.
type EntryClass struct {
	native unsafe.Pointer
}

// EntryCompletionClass is a representation of the C record GtkEntryCompletionClass.
type EntryCompletionClass struct {
	native unsafe.Pointer
}

// EntryCompletionPrivate is a representation of the C record GtkEntryCompletionPrivate.
type EntryCompletionPrivate struct {
	native unsafe.Pointer
}

// EntryPrivate is a representation of the C record GtkEntryPrivate.
type EntryPrivate struct {
	native unsafe.Pointer
}

// EventBoxClass is a representation of the C record GtkEventBoxClass.
type EventBoxClass struct {
	native unsafe.Pointer
}

// EventBoxPrivate is a representation of the C record GtkEventBoxPrivate.
type EventBoxPrivate struct {
	native unsafe.Pointer
}

// EventControllerClass is a representation of the C record GtkEventControllerClass.
type EventControllerClass struct {
	native unsafe.Pointer
}

// UNSUPPORTED : EventControllerMotionClass : blacklisted

// UNSUPPORTED : EventControllerScrollClass : blacklisted

// ExpanderAccessibleClass is a representation of the C record GtkExpanderAccessibleClass.
type ExpanderAccessibleClass struct {
	native unsafe.Pointer
}

// ExpanderAccessiblePrivate is a representation of the C record GtkExpanderAccessiblePrivate.
type ExpanderAccessiblePrivate struct {
	native unsafe.Pointer
}

// ExpanderClass is a representation of the C record GtkExpanderClass.
type ExpanderClass struct {
	native unsafe.Pointer
}

// ExpanderPrivate is a representation of the C record GtkExpanderPrivate.
type ExpanderPrivate struct {
	native unsafe.Pointer
}

// FileChooserButtonClass is a representation of the C record GtkFileChooserButtonClass.
type FileChooserButtonClass struct {
	native unsafe.Pointer
}

// FileChooserButtonPrivate is a representation of the C record GtkFileChooserButtonPrivate.
type FileChooserButtonPrivate struct {
	native unsafe.Pointer
}

// FileChooserDialogClass is a representation of the C record GtkFileChooserDialogClass.
type FileChooserDialogClass struct {
	native unsafe.Pointer
}

// FileChooserDialogPrivate is a representation of the C record GtkFileChooserDialogPrivate.
type FileChooserDialogPrivate struct {
	native unsafe.Pointer
}

// FileChooserNativeClass is a representation of the C record GtkFileChooserNativeClass.
type FileChooserNativeClass struct {
	native unsafe.Pointer
}

// FileChooserWidgetClass is a representation of the C record GtkFileChooserWidgetClass.
type FileChooserWidgetClass struct {
	native unsafe.Pointer
}

// FileChooserWidgetPrivate is a representation of the C record GtkFileChooserWidgetPrivate.
type FileChooserWidgetPrivate struct {
	native unsafe.Pointer
}

// FileFilterInfo is a representation of the C record GtkFileFilterInfo.
type FileFilterInfo struct {
	native unsafe.Pointer
}

// FixedChild is a representation of the C record GtkFixedChild.
type FixedChild struct {
	native unsafe.Pointer
}

// FixedClass is a representation of the C record GtkFixedClass.
type FixedClass struct {
	native unsafe.Pointer
}

// FixedPrivate is a representation of the C record GtkFixedPrivate.
type FixedPrivate struct {
	native unsafe.Pointer
}

// FlowBoxAccessibleClass is a representation of the C record GtkFlowBoxAccessibleClass.
type FlowBoxAccessibleClass struct {
	native unsafe.Pointer
}

// FlowBoxAccessiblePrivate is a representation of the C record GtkFlowBoxAccessiblePrivate.
type FlowBoxAccessiblePrivate struct {
	native unsafe.Pointer
}

// FlowBoxChildAccessibleClass is a representation of the C record GtkFlowBoxChildAccessibleClass.
type FlowBoxChildAccessibleClass struct {
	native unsafe.Pointer
}

// FlowBoxChildClass is a representation of the C record GtkFlowBoxChildClass.
type FlowBoxChildClass struct {
	native unsafe.Pointer
}

// FlowBoxClass is a representation of the C record GtkFlowBoxClass.
type FlowBoxClass struct {
	native unsafe.Pointer
}

// FontButtonClass is a representation of the C record GtkFontButtonClass.
type FontButtonClass struct {
	native unsafe.Pointer
}

// FontButtonPrivate is a representation of the C record GtkFontButtonPrivate.
type FontButtonPrivate struct {
	native unsafe.Pointer
}

// FontChooserDialogClass is a representation of the C record GtkFontChooserDialogClass.
type FontChooserDialogClass struct {
	native unsafe.Pointer
}

// FontChooserDialogPrivate is a representation of the C record GtkFontChooserDialogPrivate.
type FontChooserDialogPrivate struct {
	native unsafe.Pointer
}

// FontChooserIface is a representation of the C record GtkFontChooserIface.
type FontChooserIface struct {
	native unsafe.Pointer
}

// FontChooserWidgetClass is a representation of the C record GtkFontChooserWidgetClass.
type FontChooserWidgetClass struct {
	native unsafe.Pointer
}

// FontChooserWidgetPrivate is a representation of the C record GtkFontChooserWidgetPrivate.
type FontChooserWidgetPrivate struct {
	native unsafe.Pointer
}

// FontSelectionClass is a representation of the C record GtkFontSelectionClass.
type FontSelectionClass struct {
	native unsafe.Pointer
}

// FontSelectionDialogClass is a representation of the C record GtkFontSelectionDialogClass.
type FontSelectionDialogClass struct {
	native unsafe.Pointer
}

// FontSelectionDialogPrivate is a representation of the C record GtkFontSelectionDialogPrivate.
type FontSelectionDialogPrivate struct {
	native unsafe.Pointer
}

// FontSelectionPrivate is a representation of the C record GtkFontSelectionPrivate.
type FontSelectionPrivate struct {
	native unsafe.Pointer
}

// FrameAccessibleClass is a representation of the C record GtkFrameAccessibleClass.
type FrameAccessibleClass struct {
	native unsafe.Pointer
}

// FrameAccessiblePrivate is a representation of the C record GtkFrameAccessiblePrivate.
type FrameAccessiblePrivate struct {
	native unsafe.Pointer
}

// FrameClass is a representation of the C record GtkFrameClass.
type FrameClass struct {
	native unsafe.Pointer
}

// FramePrivate is a representation of the C record GtkFramePrivate.
type FramePrivate struct {
	native unsafe.Pointer
}

// GestureClass is a representation of the C record GtkGestureClass.
type GestureClass struct {
	native unsafe.Pointer
}

// GestureDragClass is a representation of the C record GtkGestureDragClass.
type GestureDragClass struct {
	native unsafe.Pointer
}

// GestureLongPressClass is a representation of the C record GtkGestureLongPressClass.
type GestureLongPressClass struct {
	native unsafe.Pointer
}

// GestureMultiPressClass is a representation of the C record GtkGestureMultiPressClass.
type GestureMultiPressClass struct {
	native unsafe.Pointer
}

// GesturePanClass is a representation of the C record GtkGesturePanClass.
type GesturePanClass struct {
	native unsafe.Pointer
}

// GestureRotateClass is a representation of the C record GtkGestureRotateClass.
type GestureRotateClass struct {
	native unsafe.Pointer
}

// GestureSingleClass is a representation of the C record GtkGestureSingleClass.
type GestureSingleClass struct {
	native unsafe.Pointer
}

// UNSUPPORTED : GestureStylusClass : blacklisted

// GestureSwipeClass is a representation of the C record GtkGestureSwipeClass.
type GestureSwipeClass struct {
	native unsafe.Pointer
}

// GestureZoomClass is a representation of the C record GtkGestureZoomClass.
type GestureZoomClass struct {
	native unsafe.Pointer
}

// Gradient is a representation of the C record GtkGradient.
type Gradient struct {
	native unsafe.Pointer
}

// GridClass is a representation of the C record GtkGridClass.
type GridClass struct {
	native unsafe.Pointer
}

// GridPrivate is a representation of the C record GtkGridPrivate.
type GridPrivate struct {
	native unsafe.Pointer
}

// HBoxClass is a representation of the C record GtkHBoxClass.
type HBoxClass struct {
	native unsafe.Pointer
}

// HButtonBoxClass is a representation of the C record GtkHButtonBoxClass.
type HButtonBoxClass struct {
	native unsafe.Pointer
}

// HPanedClass is a representation of the C record GtkHPanedClass.
type HPanedClass struct {
	native unsafe.Pointer
}

// HSVClass is a representation of the C record GtkHSVClass.
type HSVClass struct {
	native unsafe.Pointer
}

// HSVPrivate is a representation of the C record GtkHSVPrivate.
type HSVPrivate struct {
	native unsafe.Pointer
}

// HScaleClass is a representation of the C record GtkHScaleClass.
type HScaleClass struct {
	native unsafe.Pointer
}

// HScrollbarClass is a representation of the C record GtkHScrollbarClass.
type HScrollbarClass struct {
	native unsafe.Pointer
}

// HSeparatorClass is a representation of the C record GtkHSeparatorClass.
type HSeparatorClass struct {
	native unsafe.Pointer
}

// HandleBoxClass is a representation of the C record GtkHandleBoxClass.
type HandleBoxClass struct {
	native unsafe.Pointer
}

// HandleBoxPrivate is a representation of the C record GtkHandleBoxPrivate.
type HandleBoxPrivate struct {
	native unsafe.Pointer
}

// UNSUPPORTED : HeaderBarAccessibleClass : blacklisted

// UNSUPPORTED : HeaderBarAccessiblePrivate : blacklisted

// HeaderBarClass is a representation of the C record GtkHeaderBarClass.
type HeaderBarClass struct {
	native unsafe.Pointer
}

// HeaderBarPrivate is a representation of the C record GtkHeaderBarPrivate.
type HeaderBarPrivate struct {
	native unsafe.Pointer
}

// IMContextClass is a representation of the C record GtkIMContextClass.
type IMContextClass struct {
	native unsafe.Pointer
}

// IMContextInfo is a representation of the C record GtkIMContextInfo.
type IMContextInfo struct {
	native unsafe.Pointer
}

// IMContextSimpleClass is a representation of the C record GtkIMContextSimpleClass.
type IMContextSimpleClass struct {
	native unsafe.Pointer
}

// IMContextSimplePrivate is a representation of the C record GtkIMContextSimplePrivate.
type IMContextSimplePrivate struct {
	native unsafe.Pointer
}

// IMMulticontextClass is a representation of the C record GtkIMMulticontextClass.
type IMMulticontextClass struct {
	native unsafe.Pointer
}

// IMMulticontextPrivate is a representation of the C record GtkIMMulticontextPrivate.
type IMMulticontextPrivate struct {
	native unsafe.Pointer
}

// IconFactoryClass is a representation of the C record GtkIconFactoryClass.
type IconFactoryClass struct {
	native unsafe.Pointer
}

// IconFactoryPrivate is a representation of the C record GtkIconFactoryPrivate.
type IconFactoryPrivate struct {
	native unsafe.Pointer
}

// IconInfoClass is a representation of the C record GtkIconInfoClass.
type IconInfoClass struct {
	native unsafe.Pointer
}

// IconSet is a representation of the C record GtkIconSet.
type IconSet struct {
	native unsafe.Pointer
}

// IconSource is a representation of the C record GtkIconSource.
type IconSource struct {
	native unsafe.Pointer
}

// IconThemeClass is a representation of the C record GtkIconThemeClass.
type IconThemeClass struct {
	native unsafe.Pointer
}

// IconThemePrivate is a representation of the C record GtkIconThemePrivate.
type IconThemePrivate struct {
	native unsafe.Pointer
}

// IconViewAccessibleClass is a representation of the C record GtkIconViewAccessibleClass.
type IconViewAccessibleClass struct {
	native unsafe.Pointer
}

// IconViewAccessiblePrivate is a representation of the C record GtkIconViewAccessiblePrivate.
type IconViewAccessiblePrivate struct {
	native unsafe.Pointer
}

// IconViewClass is a representation of the C record GtkIconViewClass.
type IconViewClass struct {
	native unsafe.Pointer
}

// IconViewPrivate is a representation of the C record GtkIconViewPrivate.
type IconViewPrivate struct {
	native unsafe.Pointer
}

// ImageAccessibleClass is a representation of the C record GtkImageAccessibleClass.
type ImageAccessibleClass struct {
	native unsafe.Pointer
}

// ImageAccessiblePrivate is a representation of the C record GtkImageAccessiblePrivate.
type ImageAccessiblePrivate struct {
	native unsafe.Pointer
}

// ImageCellAccessibleClass is a representation of the C record GtkImageCellAccessibleClass.
type ImageCellAccessibleClass struct {
	native unsafe.Pointer
}

// ImageCellAccessiblePrivate is a representation of the C record GtkImageCellAccessiblePrivate.
type ImageCellAccessiblePrivate struct {
	native unsafe.Pointer
}

// ImageClass is a representation of the C record GtkImageClass.
type ImageClass struct {
	native unsafe.Pointer
}

// ImageMenuItemClass is a representation of the C record GtkImageMenuItemClass.
type ImageMenuItemClass struct {
	native unsafe.Pointer
}

// ImageMenuItemPrivate is a representation of the C record GtkImageMenuItemPrivate.
type ImageMenuItemPrivate struct {
	native unsafe.Pointer
}

// ImagePrivate is a representation of the C record GtkImagePrivate.
type ImagePrivate struct {
	native unsafe.Pointer
}

// InfoBarClass is a representation of the C record GtkInfoBarClass.
type InfoBarClass struct {
	native unsafe.Pointer
}

// InfoBarPrivate is a representation of the C record GtkInfoBarPrivate.
type InfoBarPrivate struct {
	native unsafe.Pointer
}

// InvisibleClass is a representation of the C record GtkInvisibleClass.
type InvisibleClass struct {
	native unsafe.Pointer
}

// InvisiblePrivate is a representation of the C record GtkInvisiblePrivate.
type InvisiblePrivate struct {
	native unsafe.Pointer
}

// LabelAccessibleClass is a representation of the C record GtkLabelAccessibleClass.
type LabelAccessibleClass struct {
	native unsafe.Pointer
}

// LabelAccessiblePrivate is a representation of the C record GtkLabelAccessiblePrivate.
type LabelAccessiblePrivate struct {
	native unsafe.Pointer
}

// LabelClass is a representation of the C record GtkLabelClass.
type LabelClass struct {
	native unsafe.Pointer
}

// LabelPrivate is a representation of the C record GtkLabelPrivate.
type LabelPrivate struct {
	native unsafe.Pointer
}

// LabelSelectionInfo is a representation of the C record GtkLabelSelectionInfo.
type LabelSelectionInfo struct {
	native unsafe.Pointer
}

// LayoutClass is a representation of the C record GtkLayoutClass.
type LayoutClass struct {
	native unsafe.Pointer
}

// LayoutPrivate is a representation of the C record GtkLayoutPrivate.
type LayoutPrivate struct {
	native unsafe.Pointer
}

// LevelBarAccessibleClass is a representation of the C record GtkLevelBarAccessibleClass.
type LevelBarAccessibleClass struct {
	native unsafe.Pointer
}

// LevelBarAccessiblePrivate is a representation of the C record GtkLevelBarAccessiblePrivate.
type LevelBarAccessiblePrivate struct {
	native unsafe.Pointer
}

// LevelBarClass is a representation of the C record GtkLevelBarClass.
type LevelBarClass struct {
	native unsafe.Pointer
}

// LevelBarPrivate is a representation of the C record GtkLevelBarPrivate.
type LevelBarPrivate struct {
	native unsafe.Pointer
}

// LinkButtonAccessibleClass is a representation of the C record GtkLinkButtonAccessibleClass.
type LinkButtonAccessibleClass struct {
	native unsafe.Pointer
}

// LinkButtonAccessiblePrivate is a representation of the C record GtkLinkButtonAccessiblePrivate.
type LinkButtonAccessiblePrivate struct {
	native unsafe.Pointer
}

// LinkButtonClass is a representation of the C record GtkLinkButtonClass.
type LinkButtonClass struct {
	native unsafe.Pointer
}

// LinkButtonPrivate is a representation of the C record GtkLinkButtonPrivate.
type LinkButtonPrivate struct {
	native unsafe.Pointer
}

// ListBoxAccessibleClass is a representation of the C record GtkListBoxAccessibleClass.
type ListBoxAccessibleClass struct {
	native unsafe.Pointer
}

// ListBoxAccessiblePrivate is a representation of the C record GtkListBoxAccessiblePrivate.
type ListBoxAccessiblePrivate struct {
	native unsafe.Pointer
}

// ListBoxClass is a representation of the C record GtkListBoxClass.
type ListBoxClass struct {
	native unsafe.Pointer
}

// ListBoxRowAccessibleClass is a representation of the C record GtkListBoxRowAccessibleClass.
type ListBoxRowAccessibleClass struct {
	native unsafe.Pointer
}

// ListBoxRowClass is a representation of the C record GtkListBoxRowClass.
type ListBoxRowClass struct {
	native unsafe.Pointer
}

// ListStoreClass is a representation of the C record GtkListStoreClass.
type ListStoreClass struct {
	native unsafe.Pointer
}

// ListStorePrivate is a representation of the C record GtkListStorePrivate.
type ListStorePrivate struct {
	native unsafe.Pointer
}

// LockButtonAccessibleClass is a representation of the C record GtkLockButtonAccessibleClass.
type LockButtonAccessibleClass struct {
	native unsafe.Pointer
}

// LockButtonAccessiblePrivate is a representation of the C record GtkLockButtonAccessiblePrivate.
type LockButtonAccessiblePrivate struct {
	native unsafe.Pointer
}

// LockButtonClass is a representation of the C record GtkLockButtonClass.
type LockButtonClass struct {
	native unsafe.Pointer
}

// LockButtonPrivate is a representation of the C record GtkLockButtonPrivate.
type LockButtonPrivate struct {
	native unsafe.Pointer
}

// MenuAccessibleClass is a representation of the C record GtkMenuAccessibleClass.
type MenuAccessibleClass struct {
	native unsafe.Pointer
}

// MenuAccessiblePrivate is a representation of the C record GtkMenuAccessiblePrivate.
type MenuAccessiblePrivate struct {
	native unsafe.Pointer
}

// MenuBarClass is a representation of the C record GtkMenuBarClass.
type MenuBarClass struct {
	native unsafe.Pointer
}

// MenuBarPrivate is a representation of the C record GtkMenuBarPrivate.
type MenuBarPrivate struct {
	native unsafe.Pointer
}

// MenuButtonAccessibleClass is a representation of the C record GtkMenuButtonAccessibleClass.
type MenuButtonAccessibleClass struct {
	native unsafe.Pointer
}

// MenuButtonAccessiblePrivate is a representation of the C record GtkMenuButtonAccessiblePrivate.
type MenuButtonAccessiblePrivate struct {
	native unsafe.Pointer
}

// MenuButtonClass is a representation of the C record GtkMenuButtonClass.
type MenuButtonClass struct {
	native unsafe.Pointer
}

// MenuButtonPrivate is a representation of the C record GtkMenuButtonPrivate.
type MenuButtonPrivate struct {
	native unsafe.Pointer
}

// MenuClass is a representation of the C record GtkMenuClass.
type MenuClass struct {
	native unsafe.Pointer
}

// MenuItemAccessibleClass is a representation of the C record GtkMenuItemAccessibleClass.
type MenuItemAccessibleClass struct {
	native unsafe.Pointer
}

// MenuItemAccessiblePrivate is a representation of the C record GtkMenuItemAccessiblePrivate.
type MenuItemAccessiblePrivate struct {
	native unsafe.Pointer
}

// MenuItemClass is a representation of the C record GtkMenuItemClass.
type MenuItemClass struct {
	native unsafe.Pointer
}

// MenuItemPrivate is a representation of the C record GtkMenuItemPrivate.
type MenuItemPrivate struct {
	native unsafe.Pointer
}

// MenuPrivate is a representation of the C record GtkMenuPrivate.
type MenuPrivate struct {
	native unsafe.Pointer
}

// MenuShellAccessibleClass is a representation of the C record GtkMenuShellAccessibleClass.
type MenuShellAccessibleClass struct {
	native unsafe.Pointer
}

// MenuShellAccessiblePrivate is a representation of the C record GtkMenuShellAccessiblePrivate.
type MenuShellAccessiblePrivate struct {
	native unsafe.Pointer
}

// MenuShellClass is a representation of the C record GtkMenuShellClass.
type MenuShellClass struct {
	native unsafe.Pointer
}

// MenuShellPrivate is a representation of the C record GtkMenuShellPrivate.
type MenuShellPrivate struct {
	native unsafe.Pointer
}

// MenuToolButtonClass is a representation of the C record GtkMenuToolButtonClass.
type MenuToolButtonClass struct {
	native unsafe.Pointer
}

// MenuToolButtonPrivate is a representation of the C record GtkMenuToolButtonPrivate.
type MenuToolButtonPrivate struct {
	native unsafe.Pointer
}

// MessageDialogClass is a representation of the C record GtkMessageDialogClass.
type MessageDialogClass struct {
	native unsafe.Pointer
}

// MessageDialogPrivate is a representation of the C record GtkMessageDialogPrivate.
type MessageDialogPrivate struct {
	native unsafe.Pointer
}

// MiscClass is a representation of the C record GtkMiscClass.
type MiscClass struct {
	native unsafe.Pointer
}

// MiscPrivate is a representation of the C record GtkMiscPrivate.
type MiscPrivate struct {
	native unsafe.Pointer
}

// MountOperationClass is a representation of the C record GtkMountOperationClass.
type MountOperationClass struct {
	native unsafe.Pointer
}

// MountOperationPrivate is a representation of the C record GtkMountOperationPrivate.
type MountOperationPrivate struct {
	native unsafe.Pointer
}

// NativeDialogClass is a representation of the C record GtkNativeDialogClass.
type NativeDialogClass struct {
	native unsafe.Pointer
}

// NotebookAccessibleClass is a representation of the C record GtkNotebookAccessibleClass.
type NotebookAccessibleClass struct {
	native unsafe.Pointer
}

// NotebookAccessiblePrivate is a representation of the C record GtkNotebookAccessiblePrivate.
type NotebookAccessiblePrivate struct {
	native unsafe.Pointer
}

// NotebookClass is a representation of the C record GtkNotebookClass.
type NotebookClass struct {
	native unsafe.Pointer
}

// NotebookPageAccessibleClass is a representation of the C record GtkNotebookPageAccessibleClass.
type NotebookPageAccessibleClass struct {
	native unsafe.Pointer
}

// NotebookPageAccessiblePrivate is a representation of the C record GtkNotebookPageAccessiblePrivate.
type NotebookPageAccessiblePrivate struct {
	native unsafe.Pointer
}

// NotebookPrivate is a representation of the C record GtkNotebookPrivate.
type NotebookPrivate struct {
	native unsafe.Pointer
}

// NumerableIconClass is a representation of the C record GtkNumerableIconClass.
type NumerableIconClass struct {
	native unsafe.Pointer
}

// NumerableIconPrivate is a representation of the C record GtkNumerableIconPrivate.
type NumerableIconPrivate struct {
	native unsafe.Pointer
}

// OffscreenWindowClass is a representation of the C record GtkOffscreenWindowClass.
type OffscreenWindowClass struct {
	native unsafe.Pointer
}

// OrientableIface is a representation of the C record GtkOrientableIface.
type OrientableIface struct {
	native unsafe.Pointer
}

// OverlayClass is a representation of the C record GtkOverlayClass.
type OverlayClass struct {
	native unsafe.Pointer
}

// OverlayPrivate is a representation of the C record GtkOverlayPrivate.
type OverlayPrivate struct {
	native unsafe.Pointer
}

// PadActionEntry is a representation of the C record GtkPadActionEntry.
type PadActionEntry struct {
	native unsafe.Pointer
}

// PadControllerClass is a representation of the C record GtkPadControllerClass.
type PadControllerClass struct {
	native unsafe.Pointer
}

// PageRange is a representation of the C record GtkPageRange.
type PageRange struct {
	native unsafe.Pointer
}

// PanedAccessibleClass is a representation of the C record GtkPanedAccessibleClass.
type PanedAccessibleClass struct {
	native unsafe.Pointer
}

// PanedAccessiblePrivate is a representation of the C record GtkPanedAccessiblePrivate.
type PanedAccessiblePrivate struct {
	native unsafe.Pointer
}

// PanedClass is a representation of the C record GtkPanedClass.
type PanedClass struct {
	native unsafe.Pointer
}

// PanedPrivate is a representation of the C record GtkPanedPrivate.
type PanedPrivate struct {
	native unsafe.Pointer
}

// PaperSize is a representation of the C record GtkPaperSize.
type PaperSize struct {
	native unsafe.Pointer
}

// PlacesSidebarClass is a representation of the C record GtkPlacesSidebarClass.
type PlacesSidebarClass struct {
	native unsafe.Pointer
}

// PlugClass is a representation of the C record GtkPlugClass.
type PlugClass struct {
	native unsafe.Pointer
}

// PlugPrivate is a representation of the C record GtkPlugPrivate.
type PlugPrivate struct {
	native unsafe.Pointer
}

// PopoverAccessibleClass is a representation of the C record GtkPopoverAccessibleClass.
type PopoverAccessibleClass struct {
	native unsafe.Pointer
}

// PopoverClass is a representation of the C record GtkPopoverClass.
type PopoverClass struct {
	native unsafe.Pointer
}

// PopoverMenuClass is a representation of the C record GtkPopoverMenuClass.
type PopoverMenuClass struct {
	native unsafe.Pointer
}

// PopoverPrivate is a representation of the C record GtkPopoverPrivate.
type PopoverPrivate struct {
	native unsafe.Pointer
}

// PrintOperationClass is a representation of the C record GtkPrintOperationClass.
type PrintOperationClass struct {
	native unsafe.Pointer
}

// PrintOperationPreviewIface is a representation of the C record GtkPrintOperationPreviewIface.
type PrintOperationPreviewIface struct {
	native unsafe.Pointer
}

// PrintOperationPrivate is a representation of the C record GtkPrintOperationPrivate.
type PrintOperationPrivate struct {
	native unsafe.Pointer
}

// ProgressBarAccessibleClass is a representation of the C record GtkProgressBarAccessibleClass.
type ProgressBarAccessibleClass struct {
	native unsafe.Pointer
}

// ProgressBarAccessiblePrivate is a representation of the C record GtkProgressBarAccessiblePrivate.
type ProgressBarAccessiblePrivate struct {
	native unsafe.Pointer
}

// ProgressBarClass is a representation of the C record GtkProgressBarClass.
type ProgressBarClass struct {
	native unsafe.Pointer
}

// ProgressBarPrivate is a representation of the C record GtkProgressBarPrivate.
type ProgressBarPrivate struct {
	native unsafe.Pointer
}

// RadioActionClass is a representation of the C record GtkRadioActionClass.
type RadioActionClass struct {
	native unsafe.Pointer
}

// RadioActionEntry is a representation of the C record GtkRadioActionEntry.
type RadioActionEntry struct {
	native unsafe.Pointer
}

// RadioActionPrivate is a representation of the C record GtkRadioActionPrivate.
type RadioActionPrivate struct {
	native unsafe.Pointer
}

// RadioButtonAccessibleClass is a representation of the C record GtkRadioButtonAccessibleClass.
type RadioButtonAccessibleClass struct {
	native unsafe.Pointer
}

// RadioButtonAccessiblePrivate is a representation of the C record GtkRadioButtonAccessiblePrivate.
type RadioButtonAccessiblePrivate struct {
	native unsafe.Pointer
}

// RadioButtonClass is a representation of the C record GtkRadioButtonClass.
type RadioButtonClass struct {
	native unsafe.Pointer
}

// RadioButtonPrivate is a representation of the C record GtkRadioButtonPrivate.
type RadioButtonPrivate struct {
	native unsafe.Pointer
}

// RadioMenuItemAccessibleClass is a representation of the C record GtkRadioMenuItemAccessibleClass.
type RadioMenuItemAccessibleClass struct {
	native unsafe.Pointer
}

// RadioMenuItemAccessiblePrivate is a representation of the C record GtkRadioMenuItemAccessiblePrivate.
type RadioMenuItemAccessiblePrivate struct {
	native unsafe.Pointer
}

// RadioMenuItemClass is a representation of the C record GtkRadioMenuItemClass.
type RadioMenuItemClass struct {
	native unsafe.Pointer
}

// RadioMenuItemPrivate is a representation of the C record GtkRadioMenuItemPrivate.
type RadioMenuItemPrivate struct {
	native unsafe.Pointer
}

// RadioToolButtonClass is a representation of the C record GtkRadioToolButtonClass.
type RadioToolButtonClass struct {
	native unsafe.Pointer
}

// RangeAccessibleClass is a representation of the C record GtkRangeAccessibleClass.
type RangeAccessibleClass struct {
	native unsafe.Pointer
}

// RangeAccessiblePrivate is a representation of the C record GtkRangeAccessiblePrivate.
type RangeAccessiblePrivate struct {
	native unsafe.Pointer
}

// RangeClass is a representation of the C record GtkRangeClass.
type RangeClass struct {
	native unsafe.Pointer
}

// RangePrivate is a representation of the C record GtkRangePrivate.
type RangePrivate struct {
	native unsafe.Pointer
}

// RcContext is a representation of the C record GtkRcContext.
type RcContext struct {
	native unsafe.Pointer
}

// RcProperty is a representation of the C record GtkRcProperty.
type RcProperty struct {
	native unsafe.Pointer
}

// RcStyleClass is a representation of the C record GtkRcStyleClass.
type RcStyleClass struct {
	native unsafe.Pointer
}

// RecentActionClass is a representation of the C record GtkRecentActionClass.
type RecentActionClass struct {
	native unsafe.Pointer
}

// RecentActionPrivate is a representation of the C record GtkRecentActionPrivate.
type RecentActionPrivate struct {
	native unsafe.Pointer
}

// RecentChooserDialogClass is a representation of the C record GtkRecentChooserDialogClass.
type RecentChooserDialogClass struct {
	native unsafe.Pointer
}

// RecentChooserDialogPrivate is a representation of the C record GtkRecentChooserDialogPrivate.
type RecentChooserDialogPrivate struct {
	native unsafe.Pointer
}

// RecentChooserIface is a representation of the C record GtkRecentChooserIface.
type RecentChooserIface struct {
	native unsafe.Pointer
}

// RecentChooserMenuClass is a representation of the C record GtkRecentChooserMenuClass.
type RecentChooserMenuClass struct {
	native unsafe.Pointer
}

// RecentChooserMenuPrivate is a representation of the C record GtkRecentChooserMenuPrivate.
type RecentChooserMenuPrivate struct {
	native unsafe.Pointer
}

// RecentChooserWidgetClass is a representation of the C record GtkRecentChooserWidgetClass.
type RecentChooserWidgetClass struct {
	native unsafe.Pointer
}

// RecentChooserWidgetPrivate is a representation of the C record GtkRecentChooserWidgetPrivate.
type RecentChooserWidgetPrivate struct {
	native unsafe.Pointer
}

// RecentData is a representation of the C record GtkRecentData.
type RecentData struct {
	native unsafe.Pointer
}

// RecentFilterInfo is a representation of the C record GtkRecentFilterInfo.
type RecentFilterInfo struct {
	native unsafe.Pointer
}

// RecentManagerPrivate is a representation of the C record GtkRecentManagerPrivate.
type RecentManagerPrivate struct {
	native unsafe.Pointer
}

// RendererCellAccessibleClass is a representation of the C record GtkRendererCellAccessibleClass.
type RendererCellAccessibleClass struct {
	native unsafe.Pointer
}

// RendererCellAccessiblePrivate is a representation of the C record GtkRendererCellAccessiblePrivate.
type RendererCellAccessiblePrivate struct {
	native unsafe.Pointer
}

// RequestedSize is a representation of the C record GtkRequestedSize.
type RequestedSize struct {
	native unsafe.Pointer
}

// Requisition is a representation of the C record GtkRequisition.
type Requisition struct {
	native unsafe.Pointer
}

// RevealerClass is a representation of the C record GtkRevealerClass.
type RevealerClass struct {
	native unsafe.Pointer
}

// ScaleAccessibleClass is a representation of the C record GtkScaleAccessibleClass.
type ScaleAccessibleClass struct {
	native unsafe.Pointer
}

// ScaleAccessiblePrivate is a representation of the C record GtkScaleAccessiblePrivate.
type ScaleAccessiblePrivate struct {
	native unsafe.Pointer
}

// ScaleButtonAccessibleClass is a representation of the C record GtkScaleButtonAccessibleClass.
type ScaleButtonAccessibleClass struct {
	native unsafe.Pointer
}

// ScaleButtonAccessiblePrivate is a representation of the C record GtkScaleButtonAccessiblePrivate.
type ScaleButtonAccessiblePrivate struct {
	native unsafe.Pointer
}

// ScaleButtonClass is a representation of the C record GtkScaleButtonClass.
type ScaleButtonClass struct {
	native unsafe.Pointer
}

// ScaleButtonPrivate is a representation of the C record GtkScaleButtonPrivate.
type ScaleButtonPrivate struct {
	native unsafe.Pointer
}

// ScaleClass is a representation of the C record GtkScaleClass.
type ScaleClass struct {
	native unsafe.Pointer
}

// ScalePrivate is a representation of the C record GtkScalePrivate.
type ScalePrivate struct {
	native unsafe.Pointer
}

// ScrollableInterface is a representation of the C record GtkScrollableInterface.
type ScrollableInterface struct {
	native unsafe.Pointer
}

// ScrollbarClass is a representation of the C record GtkScrollbarClass.
type ScrollbarClass struct {
	native unsafe.Pointer
}

// ScrolledWindowAccessibleClass is a representation of the C record GtkScrolledWindowAccessibleClass.
type ScrolledWindowAccessibleClass struct {
	native unsafe.Pointer
}

// ScrolledWindowAccessiblePrivate is a representation of the C record GtkScrolledWindowAccessiblePrivate.
type ScrolledWindowAccessiblePrivate struct {
	native unsafe.Pointer
}

// ScrolledWindowClass is a representation of the C record GtkScrolledWindowClass.
type ScrolledWindowClass struct {
	native unsafe.Pointer
}

// ScrolledWindowPrivate is a representation of the C record GtkScrolledWindowPrivate.
type ScrolledWindowPrivate struct {
	native unsafe.Pointer
}

// SearchBarClass is a representation of the C record GtkSearchBarClass.
type SearchBarClass struct {
	native unsafe.Pointer
}

// SearchEntryClass is a representation of the C record GtkSearchEntryClass.
type SearchEntryClass struct {
	native unsafe.Pointer
}

// SelectionData is a representation of the C record GtkSelectionData.
type SelectionData struct {
	native unsafe.Pointer
}

// SeparatorClass is a representation of the C record GtkSeparatorClass.
type SeparatorClass struct {
	native unsafe.Pointer
}

// SeparatorMenuItemClass is a representation of the C record GtkSeparatorMenuItemClass.
type SeparatorMenuItemClass struct {
	native unsafe.Pointer
}

// SeparatorPrivate is a representation of the C record GtkSeparatorPrivate.
type SeparatorPrivate struct {
	native unsafe.Pointer
}

// SeparatorToolItemClass is a representation of the C record GtkSeparatorToolItemClass.
type SeparatorToolItemClass struct {
	native unsafe.Pointer
}

// SeparatorToolItemPrivate is a representation of the C record GtkSeparatorToolItemPrivate.
type SeparatorToolItemPrivate struct {
	native unsafe.Pointer
}

// SettingsClass is a representation of the C record GtkSettingsClass.
type SettingsClass struct {
	native unsafe.Pointer
}

// SettingsPrivate is a representation of the C record GtkSettingsPrivate.
type SettingsPrivate struct {
	native unsafe.Pointer
}

// SettingsValue is a representation of the C record GtkSettingsValue.
type SettingsValue struct {
	native unsafe.Pointer
}

// ShortcutLabelClass is a representation of the C record GtkShortcutLabelClass.
type ShortcutLabelClass struct {
	native unsafe.Pointer
}

// ShortcutsGroupClass is a representation of the C record GtkShortcutsGroupClass.
type ShortcutsGroupClass struct {
	native unsafe.Pointer
}

// ShortcutsSectionClass is a representation of the C record GtkShortcutsSectionClass.
type ShortcutsSectionClass struct {
	native unsafe.Pointer
}

// ShortcutsShortcutClass is a representation of the C record GtkShortcutsShortcutClass.
type ShortcutsShortcutClass struct {
	native unsafe.Pointer
}

// ShortcutsWindowClass is a representation of the C record GtkShortcutsWindowClass.
type ShortcutsWindowClass struct {
	native unsafe.Pointer
}

// SizeGroupClass is a representation of the C record GtkSizeGroupClass.
type SizeGroupClass struct {
	native unsafe.Pointer
}

// SizeGroupPrivate is a representation of the C record GtkSizeGroupPrivate.
type SizeGroupPrivate struct {
	native unsafe.Pointer
}

// SocketClass is a representation of the C record GtkSocketClass.
type SocketClass struct {
	native unsafe.Pointer
}

// SocketPrivate is a representation of the C record GtkSocketPrivate.
type SocketPrivate struct {
	native unsafe.Pointer
}

// SpinButtonAccessibleClass is a representation of the C record GtkSpinButtonAccessibleClass.
type SpinButtonAccessibleClass struct {
	native unsafe.Pointer
}

// SpinButtonAccessiblePrivate is a representation of the C record GtkSpinButtonAccessiblePrivate.
type SpinButtonAccessiblePrivate struct {
	native unsafe.Pointer
}

// SpinButtonClass is a representation of the C record GtkSpinButtonClass.
type SpinButtonClass struct {
	native unsafe.Pointer
}

// SpinButtonPrivate is a representation of the C record GtkSpinButtonPrivate.
type SpinButtonPrivate struct {
	native unsafe.Pointer
}

// SpinnerAccessibleClass is a representation of the C record GtkSpinnerAccessibleClass.
type SpinnerAccessibleClass struct {
	native unsafe.Pointer
}

// SpinnerAccessiblePrivate is a representation of the C record GtkSpinnerAccessiblePrivate.
type SpinnerAccessiblePrivate struct {
	native unsafe.Pointer
}

// SpinnerClass is a representation of the C record GtkSpinnerClass.
type SpinnerClass struct {
	native unsafe.Pointer
}

// SpinnerPrivate is a representation of the C record GtkSpinnerPrivate.
type SpinnerPrivate struct {
	native unsafe.Pointer
}

// UNSUPPORTED : StackAccessibleClass : blacklisted

// StackClass is a representation of the C record GtkStackClass.
type StackClass struct {
	native unsafe.Pointer
}

// StackSidebarClass is a representation of the C record GtkStackSidebarClass.
type StackSidebarClass struct {
	native unsafe.Pointer
}

// StackSidebarPrivate is a representation of the C record GtkStackSidebarPrivate.
type StackSidebarPrivate struct {
	native unsafe.Pointer
}

// StackSwitcherClass is a representation of the C record GtkStackSwitcherClass.
type StackSwitcherClass struct {
	native unsafe.Pointer
}

// StatusIconClass is a representation of the C record GtkStatusIconClass.
type StatusIconClass struct {
	native unsafe.Pointer
}

// StatusIconPrivate is a representation of the C record GtkStatusIconPrivate.
type StatusIconPrivate struct {
	native unsafe.Pointer
}

// StatusbarAccessibleClass is a representation of the C record GtkStatusbarAccessibleClass.
type StatusbarAccessibleClass struct {
	native unsafe.Pointer
}

// StatusbarAccessiblePrivate is a representation of the C record GtkStatusbarAccessiblePrivate.
type StatusbarAccessiblePrivate struct {
	native unsafe.Pointer
}

// StatusbarClass is a representation of the C record GtkStatusbarClass.
type StatusbarClass struct {
	native unsafe.Pointer
}

// StatusbarPrivate is a representation of the C record GtkStatusbarPrivate.
type StatusbarPrivate struct {
	native unsafe.Pointer
}

// StockItem is a representation of the C record GtkStockItem.
type StockItem struct {
	native unsafe.Pointer
}

// StyleClass is a representation of the C record GtkStyleClass.
type StyleClass struct {
	native unsafe.Pointer
}

// StyleContextClass is a representation of the C record GtkStyleContextClass.
type StyleContextClass struct {
	native unsafe.Pointer
}

// StyleContextPrivate is a representation of the C record GtkStyleContextPrivate.
type StyleContextPrivate struct {
	native unsafe.Pointer
}

// StylePropertiesClass is a representation of the C record GtkStylePropertiesClass.
type StylePropertiesClass struct {
	native unsafe.Pointer
}

// StylePropertiesPrivate is a representation of the C record GtkStylePropertiesPrivate.
type StylePropertiesPrivate struct {
	native unsafe.Pointer
}

// StyleProviderIface is a representation of the C record GtkStyleProviderIface.
type StyleProviderIface struct {
	native unsafe.Pointer
}

// SwitchAccessibleClass is a representation of the C record GtkSwitchAccessibleClass.
type SwitchAccessibleClass struct {
	native unsafe.Pointer
}

// SwitchAccessiblePrivate is a representation of the C record GtkSwitchAccessiblePrivate.
type SwitchAccessiblePrivate struct {
	native unsafe.Pointer
}

// SwitchClass is a representation of the C record GtkSwitchClass.
type SwitchClass struct {
	native unsafe.Pointer
}

// SwitchPrivate is a representation of the C record GtkSwitchPrivate.
type SwitchPrivate struct {
	native unsafe.Pointer
}

// SymbolicColor is a representation of the C record GtkSymbolicColor.
type SymbolicColor struct {
	native unsafe.Pointer
}

// TableChild is a representation of the C record GtkTableChild.
type TableChild struct {
	native unsafe.Pointer
}

// TableClass is a representation of the C record GtkTableClass.
type TableClass struct {
	native unsafe.Pointer
}

// TablePrivate is a representation of the C record GtkTablePrivate.
type TablePrivate struct {
	native unsafe.Pointer
}

// TableRowCol is a representation of the C record GtkTableRowCol.
type TableRowCol struct {
	native unsafe.Pointer
}

// TargetEntry is a representation of the C record GtkTargetEntry.
type TargetEntry struct {
	native unsafe.Pointer
}

// TargetList is a representation of the C record GtkTargetList.
type TargetList struct {
	native unsafe.Pointer
}

// TargetPair is a representation of the C record GtkTargetPair.
type TargetPair struct {
	native unsafe.Pointer
}

// TearoffMenuItemClass is a representation of the C record GtkTearoffMenuItemClass.
type TearoffMenuItemClass struct {
	native unsafe.Pointer
}

// TearoffMenuItemPrivate is a representation of the C record GtkTearoffMenuItemPrivate.
type TearoffMenuItemPrivate struct {
	native unsafe.Pointer
}

// TextAppearance is a representation of the C record GtkTextAppearance.
type TextAppearance struct {
	native unsafe.Pointer
}

// TextAttributes is a representation of the C record GtkTextAttributes.
type TextAttributes struct {
	native unsafe.Pointer
}

// TextBTree is a representation of the C record GtkTextBTree.
type TextBTree struct {
	native unsafe.Pointer
}

// TextBufferClass is a representation of the C record GtkTextBufferClass.
type TextBufferClass struct {
	native unsafe.Pointer
}

// TextBufferPrivate is a representation of the C record GtkTextBufferPrivate.
type TextBufferPrivate struct {
	native unsafe.Pointer
}

// TextCellAccessibleClass is a representation of the C record GtkTextCellAccessibleClass.
type TextCellAccessibleClass struct {
	native unsafe.Pointer
}

// TextCellAccessiblePrivate is a representation of the C record GtkTextCellAccessiblePrivate.
type TextCellAccessiblePrivate struct {
	native unsafe.Pointer
}

// TextChildAnchorClass is a representation of the C record GtkTextChildAnchorClass.
type TextChildAnchorClass struct {
	native unsafe.Pointer
}

// TextIter is a representation of the C record GtkTextIter.
type TextIter struct {
	native unsafe.Pointer
}

// TextMarkClass is a representation of the C record GtkTextMarkClass.
type TextMarkClass struct {
	native unsafe.Pointer
}

// TextTagClass is a representation of the C record GtkTextTagClass.
type TextTagClass struct {
	native unsafe.Pointer
}

// TextTagPrivate is a representation of the C record GtkTextTagPrivate.
type TextTagPrivate struct {
	native unsafe.Pointer
}

// TextTagTableClass is a representation of the C record GtkTextTagTableClass.
type TextTagTableClass struct {
	native unsafe.Pointer
}

// TextTagTablePrivate is a representation of the C record GtkTextTagTablePrivate.
type TextTagTablePrivate struct {
	native unsafe.Pointer
}

// TextViewAccessibleClass is a representation of the C record GtkTextViewAccessibleClass.
type TextViewAccessibleClass struct {
	native unsafe.Pointer
}

// TextViewAccessiblePrivate is a representation of the C record GtkTextViewAccessiblePrivate.
type TextViewAccessiblePrivate struct {
	native unsafe.Pointer
}

// TextViewClass is a representation of the C record GtkTextViewClass.
type TextViewClass struct {
	native unsafe.Pointer
}

// TextViewPrivate is a representation of the C record GtkTextViewPrivate.
type TextViewPrivate struct {
	native unsafe.Pointer
}

// ThemeEngine is a representation of the C record GtkThemeEngine.
type ThemeEngine struct {
	native unsafe.Pointer
}

// ThemingEngineClass is a representation of the C record GtkThemingEngineClass.
type ThemingEngineClass struct {
	native unsafe.Pointer
}

// ThemingEnginePrivate is a representation of the C record GtkThemingEnginePrivate.
type ThemingEnginePrivate struct {
	native unsafe.Pointer
}

// ToggleActionClass is a representation of the C record GtkToggleActionClass.
type ToggleActionClass struct {
	native unsafe.Pointer
}

// ToggleActionEntry is a representation of the C record GtkToggleActionEntry.
type ToggleActionEntry struct {
	native unsafe.Pointer
}

// ToggleActionPrivate is a representation of the C record GtkToggleActionPrivate.
type ToggleActionPrivate struct {
	native unsafe.Pointer
}

// ToggleButtonAccessibleClass is a representation of the C record GtkToggleButtonAccessibleClass.
type ToggleButtonAccessibleClass struct {
	native unsafe.Pointer
}

// ToggleButtonAccessiblePrivate is a representation of the C record GtkToggleButtonAccessiblePrivate.
type ToggleButtonAccessiblePrivate struct {
	native unsafe.Pointer
}

// ToggleButtonClass is a representation of the C record GtkToggleButtonClass.
type ToggleButtonClass struct {
	native unsafe.Pointer
}

// ToggleButtonPrivate is a representation of the C record GtkToggleButtonPrivate.
type ToggleButtonPrivate struct {
	native unsafe.Pointer
}

// ToggleToolButtonClass is a representation of the C record GtkToggleToolButtonClass.
type ToggleToolButtonClass struct {
	native unsafe.Pointer
}

// ToggleToolButtonPrivate is a representation of the C record GtkToggleToolButtonPrivate.
type ToggleToolButtonPrivate struct {
	native unsafe.Pointer
}

// ToolButtonClass is a representation of the C record GtkToolButtonClass.
type ToolButtonClass struct {
	native unsafe.Pointer
}

// ToolButtonPrivate is a representation of the C record GtkToolButtonPrivate.
type ToolButtonPrivate struct {
	native unsafe.Pointer
}

// ToolItemClass is a representation of the C record GtkToolItemClass.
type ToolItemClass struct {
	native unsafe.Pointer
}

// ToolItemGroupClass is a representation of the C record GtkToolItemGroupClass.
type ToolItemGroupClass struct {
	native unsafe.Pointer
}

// ToolItemGroupPrivate is a representation of the C record GtkToolItemGroupPrivate.
type ToolItemGroupPrivate struct {
	native unsafe.Pointer
}

// ToolItemPrivate is a representation of the C record GtkToolItemPrivate.
type ToolItemPrivate struct {
	native unsafe.Pointer
}

// ToolPaletteClass is a representation of the C record GtkToolPaletteClass.
type ToolPaletteClass struct {
	native unsafe.Pointer
}

// ToolPalettePrivate is a representation of the C record GtkToolPalettePrivate.
type ToolPalettePrivate struct {
	native unsafe.Pointer
}

// ToolShellIface is a representation of the C record GtkToolShellIface.
type ToolShellIface struct {
	native unsafe.Pointer
}

// ToolbarClass is a representation of the C record GtkToolbarClass.
type ToolbarClass struct {
	native unsafe.Pointer
}

// ToolbarPrivate is a representation of the C record GtkToolbarPrivate.
type ToolbarPrivate struct {
	native unsafe.Pointer
}

// ToplevelAccessibleClass is a representation of the C record GtkToplevelAccessibleClass.
type ToplevelAccessibleClass struct {
	native unsafe.Pointer
}

// ToplevelAccessiblePrivate is a representation of the C record GtkToplevelAccessiblePrivate.
type ToplevelAccessiblePrivate struct {
	native unsafe.Pointer
}

// TreeDragDestIface is a representation of the C record GtkTreeDragDestIface.
type TreeDragDestIface struct {
	native unsafe.Pointer
}

// TreeDragSourceIface is a representation of the C record GtkTreeDragSourceIface.
type TreeDragSourceIface struct {
	native unsafe.Pointer
}

// TreeIter is a representation of the C record GtkTreeIter.
type TreeIter struct {
	native unsafe.Pointer
}

// TreeModelFilterClass is a representation of the C record GtkTreeModelFilterClass.
type TreeModelFilterClass struct {
	native unsafe.Pointer
}

// TreeModelFilterPrivate is a representation of the C record GtkTreeModelFilterPrivate.
type TreeModelFilterPrivate struct {
	native unsafe.Pointer
}

// TreeModelIface is a representation of the C record GtkTreeModelIface.
type TreeModelIface struct {
	native unsafe.Pointer
}

// TreeModelSortClass is a representation of the C record GtkTreeModelSortClass.
type TreeModelSortClass struct {
	native unsafe.Pointer
}

// TreeModelSortPrivate is a representation of the C record GtkTreeModelSortPrivate.
type TreeModelSortPrivate struct {
	native unsafe.Pointer
}

// TreePath is a representation of the C record GtkTreePath.
type TreePath struct {
	native unsafe.Pointer
}

// TreeRowReference is a representation of the C record GtkTreeRowReference.
type TreeRowReference struct {
	native unsafe.Pointer
}

// TreeSelectionClass is a representation of the C record GtkTreeSelectionClass.
type TreeSelectionClass struct {
	native unsafe.Pointer
}

// TreeSelectionPrivate is a representation of the C record GtkTreeSelectionPrivate.
type TreeSelectionPrivate struct {
	native unsafe.Pointer
}

// TreeSortableIface is a representation of the C record GtkTreeSortableIface.
type TreeSortableIface struct {
	native unsafe.Pointer
}

// TreeStoreClass is a representation of the C record GtkTreeStoreClass.
type TreeStoreClass struct {
	native unsafe.Pointer
}

// TreeStorePrivate is a representation of the C record GtkTreeStorePrivate.
type TreeStorePrivate struct {
	native unsafe.Pointer
}

// TreeViewAccessibleClass is a representation of the C record GtkTreeViewAccessibleClass.
type TreeViewAccessibleClass struct {
	native unsafe.Pointer
}

// TreeViewAccessiblePrivate is a representation of the C record GtkTreeViewAccessiblePrivate.
type TreeViewAccessiblePrivate struct {
	native unsafe.Pointer
}

// TreeViewClass is a representation of the C record GtkTreeViewClass.
type TreeViewClass struct {
	native unsafe.Pointer
}

// TreeViewColumnClass is a representation of the C record GtkTreeViewColumnClass.
type TreeViewColumnClass struct {
	native unsafe.Pointer
}

// TreeViewColumnPrivate is a representation of the C record GtkTreeViewColumnPrivate.
type TreeViewColumnPrivate struct {
	native unsafe.Pointer
}

// TreeViewPrivate is a representation of the C record GtkTreeViewPrivate.
type TreeViewPrivate struct {
	native unsafe.Pointer
}

// UIManagerClass is a representation of the C record GtkUIManagerClass.
type UIManagerClass struct {
	native unsafe.Pointer
}

// UIManagerPrivate is a representation of the C record GtkUIManagerPrivate.
type UIManagerPrivate struct {
	native unsafe.Pointer
}

// VBoxClass is a representation of the C record GtkVBoxClass.
type VBoxClass struct {
	native unsafe.Pointer
}

// VButtonBoxClass is a representation of the C record GtkVButtonBoxClass.
type VButtonBoxClass struct {
	native unsafe.Pointer
}

// VPanedClass is a representation of the C record GtkVPanedClass.
type VPanedClass struct {
	native unsafe.Pointer
}

// VScaleClass is a representation of the C record GtkVScaleClass.
type VScaleClass struct {
	native unsafe.Pointer
}

// VScrollbarClass is a representation of the C record GtkVScrollbarClass.
type VScrollbarClass struct {
	native unsafe.Pointer
}

// VSeparatorClass is a representation of the C record GtkVSeparatorClass.
type VSeparatorClass struct {
	native unsafe.Pointer
}

// ViewportClass is a representation of the C record GtkViewportClass.
type ViewportClass struct {
	native unsafe.Pointer
}

// ViewportPrivate is a representation of the C record GtkViewportPrivate.
type ViewportPrivate struct {
	native unsafe.Pointer
}

// VolumeButtonClass is a representation of the C record GtkVolumeButtonClass.
type VolumeButtonClass struct {
	native unsafe.Pointer
}

// WidgetAccessibleClass is a representation of the C record GtkWidgetAccessibleClass.
type WidgetAccessibleClass struct {
	native unsafe.Pointer
}

// WidgetAccessiblePrivate is a representation of the C record GtkWidgetAccessiblePrivate.
type WidgetAccessiblePrivate struct {
	native unsafe.Pointer
}

// WidgetClass is a representation of the C record GtkWidgetClass.
type WidgetClass struct {
	native unsafe.Pointer
}

// WidgetClassPrivate is a representation of the C record GtkWidgetClassPrivate.
type WidgetClassPrivate struct {
	native unsafe.Pointer
}

// WidgetPath is a representation of the C record GtkWidgetPath.
type WidgetPath struct {
	native unsafe.Pointer
}

// WidgetPrivate is a representation of the C record GtkWidgetPrivate.
type WidgetPrivate struct {
	native unsafe.Pointer
}

// WindowAccessibleClass is a representation of the C record GtkWindowAccessibleClass.
type WindowAccessibleClass struct {
	native unsafe.Pointer
}

// WindowAccessiblePrivate is a representation of the C record GtkWindowAccessiblePrivate.
type WindowAccessiblePrivate struct {
	native unsafe.Pointer
}

// WindowClass is a representation of the C record GtkWindowClass.
type WindowClass struct {
	native unsafe.Pointer
}

// WindowGeometryInfo is a representation of the C record GtkWindowGeometryInfo.
type WindowGeometryInfo struct {
	native unsafe.Pointer
}

// WindowGroupClass is a representation of the C record GtkWindowGroupClass.
type WindowGroupClass struct {
	native unsafe.Pointer
}

// WindowGroupPrivate is a representation of the C record GtkWindowGroupPrivate.
type WindowGroupPrivate struct {
	native unsafe.Pointer
}

// WindowPrivate is a representation of the C record GtkWindowPrivate.
type WindowPrivate struct {
	native unsafe.Pointer
}
