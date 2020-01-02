// Code generated - DO NOT EDIT.
// +build gtk_2.6

package gtk

import (
	cairo "github.com/pekim/gobbi/lib/cairo"
	gdk "github.com/pekim/gobbi/lib/gdk"
	gdkpixbuf "github.com/pekim/gobbi/lib/gdkpixbuf"
	glib "github.com/pekim/gobbi/lib/glib"
	gobject "github.com/pekim/gobbi/lib/gobject"
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

// STOCK_ABOUT is a representation of the C constant GTK_STOCK_ABOUT.
//
// since 2.6
const STOCK_ABOUT = "gtk-about"

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

// STOCK_CONNECT is a representation of the C constant GTK_STOCK_CONNECT.
//
// since 2.6
const STOCK_CONNECT = "gtk-connect"

// STOCK_CONVERT is a representation of the C constant GTK_STOCK_CONVERT.
const STOCK_CONVERT = "gtk-convert"

// STOCK_COPY is a representation of the C constant GTK_STOCK_COPY.
const STOCK_COPY = "gtk-copy"

// STOCK_CUT is a representation of the C constant GTK_STOCK_CUT.
const STOCK_CUT = "gtk-cut"

// STOCK_DELETE is a representation of the C constant GTK_STOCK_DELETE.
const STOCK_DELETE = "gtk-delete"

// STOCK_DIALOG_AUTHENTICATION is a representation of the C constant GTK_STOCK_DIALOG_AUTHENTICATION.
//
// since 2.4
const STOCK_DIALOG_AUTHENTICATION = "gtk-dialog-authentication"

// STOCK_DIALOG_ERROR is a representation of the C constant GTK_STOCK_DIALOG_ERROR.
const STOCK_DIALOG_ERROR = "gtk-dialog-error"

// STOCK_DIALOG_INFO is a representation of the C constant GTK_STOCK_DIALOG_INFO.
const STOCK_DIALOG_INFO = "gtk-dialog-info"

// STOCK_DIALOG_QUESTION is a representation of the C constant GTK_STOCK_DIALOG_QUESTION.
const STOCK_DIALOG_QUESTION = "gtk-dialog-question"

// STOCK_DIALOG_WARNING is a representation of the C constant GTK_STOCK_DIALOG_WARNING.
const STOCK_DIALOG_WARNING = "gtk-dialog-warning"

// STOCK_DIRECTORY is a representation of the C constant GTK_STOCK_DIRECTORY.
//
// since 2.6
const STOCK_DIRECTORY = "gtk-directory"

// STOCK_DISCONNECT is a representation of the C constant GTK_STOCK_DISCONNECT.
//
// since 2.6
const STOCK_DISCONNECT = "gtk-disconnect"

// STOCK_DND is a representation of the C constant GTK_STOCK_DND.
const STOCK_DND = "gtk-dnd"

// STOCK_DND_MULTIPLE is a representation of the C constant GTK_STOCK_DND_MULTIPLE.
const STOCK_DND_MULTIPLE = "gtk-dnd-multiple"

// STOCK_EDIT is a representation of the C constant GTK_STOCK_EDIT.
//
// since 2.6
const STOCK_EDIT = "gtk-edit"

// STOCK_EXECUTE is a representation of the C constant GTK_STOCK_EXECUTE.
const STOCK_EXECUTE = "gtk-execute"

// STOCK_FILE is a representation of the C constant GTK_STOCK_FILE.
//
// since 2.6
const STOCK_FILE = "gtk-file"

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

// STOCK_HARDDISK is a representation of the C constant GTK_STOCK_HARDDISK.
//
// since 2.4
const STOCK_HARDDISK = "gtk-harddisk"

// STOCK_HELP is a representation of the C constant GTK_STOCK_HELP.
const STOCK_HELP = "gtk-help"

// STOCK_HOME is a representation of the C constant GTK_STOCK_HOME.
const STOCK_HOME = "gtk-home"

// STOCK_INDENT is a representation of the C constant GTK_STOCK_INDENT.
//
// since 2.4
const STOCK_INDENT = "gtk-indent"

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

// STOCK_MEDIA_FORWARD is a representation of the C constant GTK_STOCK_MEDIA_FORWARD.
//
// since 2.6
const STOCK_MEDIA_FORWARD = "gtk-media-forward"

// STOCK_MEDIA_NEXT is a representation of the C constant GTK_STOCK_MEDIA_NEXT.
//
// since 2.6
const STOCK_MEDIA_NEXT = "gtk-media-next"

// STOCK_MEDIA_PAUSE is a representation of the C constant GTK_STOCK_MEDIA_PAUSE.
//
// since 2.6
const STOCK_MEDIA_PAUSE = "gtk-media-pause"

// STOCK_MEDIA_PLAY is a representation of the C constant GTK_STOCK_MEDIA_PLAY.
//
// since 2.6
const STOCK_MEDIA_PLAY = "gtk-media-play"

// STOCK_MEDIA_PREVIOUS is a representation of the C constant GTK_STOCK_MEDIA_PREVIOUS.
//
// since 2.6
const STOCK_MEDIA_PREVIOUS = "gtk-media-previous"

// STOCK_MEDIA_RECORD is a representation of the C constant GTK_STOCK_MEDIA_RECORD.
//
// since 2.6
const STOCK_MEDIA_RECORD = "gtk-media-record"

// STOCK_MEDIA_REWIND is a representation of the C constant GTK_STOCK_MEDIA_REWIND.
//
// since 2.6
const STOCK_MEDIA_REWIND = "gtk-media-rewind"

// STOCK_MEDIA_STOP is a representation of the C constant GTK_STOCK_MEDIA_STOP.
//
// since 2.6
const STOCK_MEDIA_STOP = "gtk-media-stop"

// STOCK_MISSING_IMAGE is a representation of the C constant GTK_STOCK_MISSING_IMAGE.
const STOCK_MISSING_IMAGE = "gtk-missing-image"

// STOCK_NETWORK is a representation of the C constant GTK_STOCK_NETWORK.
//
// since 2.4
const STOCK_NETWORK = "gtk-network"

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

// STOCK_UNINDENT is a representation of the C constant GTK_STOCK_UNINDENT.
//
// since 2.4
const STOCK_UNINDENT = "gtk-unindent"

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

func AccelGroupsActivate(object *gobject.Object, accelKey uint, accelMods int) {
	sys_object := object.ToC()
	sys_accelKey := accelKey
	sys_accelMods := accelMods
}

func AccelGroupsFromObject(object *gobject.Object) {
	sys_object := object.ToC()
}

func AcceleratorGetDefaultModMask() {}

func AcceleratorGetLabel(acceleratorKey uint, acceleratorMods int) {
	sys_acceleratorKey := acceleratorKey
	sys_acceleratorMods := acceleratorMods
}

func AcceleratorName(acceleratorKey uint, acceleratorMods int) {
	sys_acceleratorKey := acceleratorKey
	sys_acceleratorMods := acceleratorMods
}

func AcceleratorParse(accelerator string) {
	sys_accelerator := accelerator
}

// UNSUPPORTED : gtk_accelerator_parse_with_keycode : parameter 'accelerator_codes' is array parameter without length parameter

func AcceleratorSetDefaultModMask(defaultModMask int) {
	sys_defaultModMask := defaultModMask
}

func AcceleratorValid(keyval uint, modifiers int) {
	sys_keyval := keyval
	sys_modifiers := modifiers
}

func AlternativeDialogButtonOrder(screen *gdk.Screen) {
	sys_screen := screen.ToC()
}

func BindingsActivate(object *gobject.Object, keyval uint, modifiers int) {
	sys_object := object.ToC()
	sys_keyval := keyval
	sys_modifiers := modifiers
}

func BindingsActivateEvent(object *gobject.Object, event *gdk.EventKey) {
	sys_object := object.ToC()
	sys_event := event.ToC()
}

func CheckVersion(requiredMajor uint, requiredMinor uint, requiredMicro uint) {
	sys_requiredMajor := requiredMajor
	sys_requiredMinor := requiredMinor
	sys_requiredMicro := requiredMicro
}

func DisableSetlocale() {}

func DistributeNaturalAllocation(extraSpace int, nRequestedSizes uint, sizes *RequestedSize) {
	sys_extraSpace := extraSpace
	sys_nRequestedSizes := nRequestedSizes
	sys_sizes := sizes.ToC()
}

func DragFinish(context *gdk.DragContext, success bool, del bool, time uint32) {
	sys_context := context.ToC()
	sys_success := success
	sys_del := del
	sys_time := time
}

func DragGetSourceWidget(context *gdk.DragContext) {
	sys_context := context.ToC()
}

func DragSetIconDefault(context *gdk.DragContext) {
	sys_context := context.ToC()
}

func DragSetIconPixbuf(context *gdk.DragContext, pixbuf *gdkpixbuf.Pixbuf, hotX int, hotY int) {
	sys_context := context.ToC()
	sys_pixbuf := pixbuf.ToC()
	sys_hotX := hotX
	sys_hotY := hotY
}

func DragSetIconStock(context *gdk.DragContext, stockId string, hotX int, hotY int) {
	sys_context := context.ToC()
	sys_stockId := stockId
	sys_hotX := hotX
	sys_hotY := hotY
}

func DragSetIconSurface(context *gdk.DragContext, surface *cairo.Surface) {
	sys_context := context.ToC()
	sys_surface := surface.ToC()
}

func DragSetIconWidget(context *gdk.DragContext, widget *Widget, hotX int, hotY int) {
	sys_context := context.ToC()
	sys_widget := widget.ToC()
	sys_hotX := hotX
	sys_hotY := hotY
}

func EventsPending() {}

func False() {}

func GetCurrentEvent() {}

func GetCurrentEventDevice() {}

func GetCurrentEventState() {}

func GetCurrentEventTime() {}

func GetDebugFlags() {}

func GetDefaultLanguage() {}

func GetEventWidget(event *gdk.Event) {
	sys_event := event.ToC()
}

func GetOptionGroup(openDefaultDisplay bool) {
	sys_openDefaultDisplay := openDefaultDisplay
}

func GrabGetCurrent() {}

// UNSUPPORTED : gtk_init : has array param, argv

// UNSUPPORTED : gtk_init_check : has array param, argv

// UNSUPPORTED : gtk_init_with_args : parameter 'entries' is array parameter without length parameter

// UNSUPPORTED : gtk_key_snooper_install : parameter 'snooper' is callback

func KeySnooperRemove(snooperHandlerId uint) {
	sys_snooperHandlerId := snooperHandlerId
}

func Main() {}

func MainDoEvent(event *gdk.Event) {
	sys_event := event.ToC()
}

func MainIteration() {}

func MainIterationDo(blocking bool) {
	sys_blocking := blocking
}

func MainLevel() {}

func MainQuit() {}

func PaintArrow(style *Style, cr *cairo.Context, stateType int, shadowType int, widget *Widget, detail string, arrowType int, fill bool, x int, y int, width int, height int) {
	sys_style := style.ToC()
	sys_cr := cr.ToC()
	sys_stateType := stateType
	sys_shadowType := shadowType
	sys_widget := widget.ToC()
	sys_detail := detail
	sys_arrowType := arrowType
	sys_fill := fill
	sys_x := x
	sys_y := y
	sys_width := width
	sys_height := height
}

func PaintBox(style *Style, cr *cairo.Context, stateType int, shadowType int, widget *Widget, detail string, x int, y int, width int, height int) {
	sys_style := style.ToC()
	sys_cr := cr.ToC()
	sys_stateType := stateType
	sys_shadowType := shadowType
	sys_widget := widget.ToC()
	sys_detail := detail
	sys_x := x
	sys_y := y
	sys_width := width
	sys_height := height
}

func PaintBoxGap(style *Style, cr *cairo.Context, stateType int, shadowType int, widget *Widget, detail string, x int, y int, width int, height int, gapSide int, gapX int, gapWidth int) {
	sys_style := style.ToC()
	sys_cr := cr.ToC()
	sys_stateType := stateType
	sys_shadowType := shadowType
	sys_widget := widget.ToC()
	sys_detail := detail
	sys_x := x
	sys_y := y
	sys_width := width
	sys_height := height
	sys_gapSide := gapSide
	sys_gapX := gapX
	sys_gapWidth := gapWidth
}

func PaintCheck(style *Style, cr *cairo.Context, stateType int, shadowType int, widget *Widget, detail string, x int, y int, width int, height int) {
	sys_style := style.ToC()
	sys_cr := cr.ToC()
	sys_stateType := stateType
	sys_shadowType := shadowType
	sys_widget := widget.ToC()
	sys_detail := detail
	sys_x := x
	sys_y := y
	sys_width := width
	sys_height := height
}

func PaintDiamond(style *Style, cr *cairo.Context, stateType int, shadowType int, widget *Widget, detail string, x int, y int, width int, height int) {
	sys_style := style.ToC()
	sys_cr := cr.ToC()
	sys_stateType := stateType
	sys_shadowType := shadowType
	sys_widget := widget.ToC()
	sys_detail := detail
	sys_x := x
	sys_y := y
	sys_width := width
	sys_height := height
}

func PaintExpander(style *Style, cr *cairo.Context, stateType int, widget *Widget, detail string, x int, y int, expanderStyle int) {
	sys_style := style.ToC()
	sys_cr := cr.ToC()
	sys_stateType := stateType
	sys_widget := widget.ToC()
	sys_detail := detail
	sys_x := x
	sys_y := y
	sys_expanderStyle := expanderStyle
}

func PaintExtension(style *Style, cr *cairo.Context, stateType int, shadowType int, widget *Widget, detail string, x int, y int, width int, height int, gapSide int) {
	sys_style := style.ToC()
	sys_cr := cr.ToC()
	sys_stateType := stateType
	sys_shadowType := shadowType
	sys_widget := widget.ToC()
	sys_detail := detail
	sys_x := x
	sys_y := y
	sys_width := width
	sys_height := height
	sys_gapSide := gapSide
}

func PaintFlatBox(style *Style, cr *cairo.Context, stateType int, shadowType int, widget *Widget, detail string, x int, y int, width int, height int) {
	sys_style := style.ToC()
	sys_cr := cr.ToC()
	sys_stateType := stateType
	sys_shadowType := shadowType
	sys_widget := widget.ToC()
	sys_detail := detail
	sys_x := x
	sys_y := y
	sys_width := width
	sys_height := height
}

func PaintFocus(style *Style, cr *cairo.Context, stateType int, widget *Widget, detail string, x int, y int, width int, height int) {
	sys_style := style.ToC()
	sys_cr := cr.ToC()
	sys_stateType := stateType
	sys_widget := widget.ToC()
	sys_detail := detail
	sys_x := x
	sys_y := y
	sys_width := width
	sys_height := height
}

func PaintHandle(style *Style, cr *cairo.Context, stateType int, shadowType int, widget *Widget, detail string, x int, y int, width int, height int, orientation int) {
	sys_style := style.ToC()
	sys_cr := cr.ToC()
	sys_stateType := stateType
	sys_shadowType := shadowType
	sys_widget := widget.ToC()
	sys_detail := detail
	sys_x := x
	sys_y := y
	sys_width := width
	sys_height := height
	sys_orientation := orientation
}

func PaintHline(style *Style, cr *cairo.Context, stateType int, widget *Widget, detail string, x1 int, x2 int, y int) {
	sys_style := style.ToC()
	sys_cr := cr.ToC()
	sys_stateType := stateType
	sys_widget := widget.ToC()
	sys_detail := detail
	sys_x1 := x1
	sys_x2 := x2
	sys_y := y
}

func PaintLayout(style *Style, cr *cairo.Context, stateType int, useText bool, widget *Widget, detail string, x int, y int, layout *pango.Layout) {
	sys_style := style.ToC()
	sys_cr := cr.ToC()
	sys_stateType := stateType
	sys_useText := useText
	sys_widget := widget.ToC()
	sys_detail := detail
	sys_x := x
	sys_y := y
	sys_layout := layout.ToC()
}

func PaintOption(style *Style, cr *cairo.Context, stateType int, shadowType int, widget *Widget, detail string, x int, y int, width int, height int) {
	sys_style := style.ToC()
	sys_cr := cr.ToC()
	sys_stateType := stateType
	sys_shadowType := shadowType
	sys_widget := widget.ToC()
	sys_detail := detail
	sys_x := x
	sys_y := y
	sys_width := width
	sys_height := height
}

func PaintResizeGrip(style *Style, cr *cairo.Context, stateType int, widget *Widget, detail string, edge int, x int, y int, width int, height int) {
	sys_style := style.ToC()
	sys_cr := cr.ToC()
	sys_stateType := stateType
	sys_widget := widget.ToC()
	sys_detail := detail
	sys_edge := edge
	sys_x := x
	sys_y := y
	sys_width := width
	sys_height := height
}

func PaintShadow(style *Style, cr *cairo.Context, stateType int, shadowType int, widget *Widget, detail string, x int, y int, width int, height int) {
	sys_style := style.ToC()
	sys_cr := cr.ToC()
	sys_stateType := stateType
	sys_shadowType := shadowType
	sys_widget := widget.ToC()
	sys_detail := detail
	sys_x := x
	sys_y := y
	sys_width := width
	sys_height := height
}

func PaintShadowGap(style *Style, cr *cairo.Context, stateType int, shadowType int, widget *Widget, detail string, x int, y int, width int, height int, gapSide int, gapX int, gapWidth int) {
	sys_style := style.ToC()
	sys_cr := cr.ToC()
	sys_stateType := stateType
	sys_shadowType := shadowType
	sys_widget := widget.ToC()
	sys_detail := detail
	sys_x := x
	sys_y := y
	sys_width := width
	sys_height := height
	sys_gapSide := gapSide
	sys_gapX := gapX
	sys_gapWidth := gapWidth
}

func PaintSlider(style *Style, cr *cairo.Context, stateType int, shadowType int, widget *Widget, detail string, x int, y int, width int, height int, orientation int) {
	sys_style := style.ToC()
	sys_cr := cr.ToC()
	sys_stateType := stateType
	sys_shadowType := shadowType
	sys_widget := widget.ToC()
	sys_detail := detail
	sys_x := x
	sys_y := y
	sys_width := width
	sys_height := height
	sys_orientation := orientation
}

func PaintSpinner(style *Style, cr *cairo.Context, stateType int, widget *Widget, detail string, step uint, x int, y int, width int, height int) {
	sys_style := style.ToC()
	sys_cr := cr.ToC()
	sys_stateType := stateType
	sys_widget := widget.ToC()
	sys_detail := detail
	sys_step := step
	sys_x := x
	sys_y := y
	sys_width := width
	sys_height := height
}

func PaintTab(style *Style, cr *cairo.Context, stateType int, shadowType int, widget *Widget, detail string, x int, y int, width int, height int) {
	sys_style := style.ToC()
	sys_cr := cr.ToC()
	sys_stateType := stateType
	sys_shadowType := shadowType
	sys_widget := widget.ToC()
	sys_detail := detail
	sys_x := x
	sys_y := y
	sys_width := width
	sys_height := height
}

func PaintVline(style *Style, cr *cairo.Context, stateType int, widget *Widget, detail string, y1 int, y2 int, x int) {
	sys_style := style.ToC()
	sys_cr := cr.ToC()
	sys_stateType := stateType
	sys_widget := widget.ToC()
	sys_detail := detail
	sys_y1 := y1
	sys_y2 := y2
	sys_x := x
}

// UNSUPPORTED : gtk_parse_args : has array param, argv

// UNSUPPORTED : gtk_print_run_page_setup_dialog_async : parameter 'done_cb' is callback

func PropagateEvent(widget *Widget, event *gdk.Event) {
	sys_widget := widget.ToC()
	sys_event := event.ToC()
}

func RcAddDefaultFile(filename string) {
	sys_filename := filename
}

func RcFindModuleInPath(moduleFile string) {
	sys_moduleFile := moduleFile
}

func RcFindPixmapInPath(settings *Settings, scanner *glib.Scanner, pixmapFile string) {
	sys_settings := settings.ToC()
	sys_scanner := scanner.ToC()
	sys_pixmapFile := pixmapFile
}

// UNSUPPORTED : gtk_rc_get_default_files : no array length

func RcGetImModuleFile() {}

func RcGetImModulePath() {}

func RcGetModuleDir() {}

func RcGetStyle(widget *Widget) {
	sys_widget := widget.ToC()
}

func RcGetStyleByPaths(settings *Settings, widgetPath string, classPath string, type_ uint64) {
	sys_settings := settings.ToC()
	sys_widgetPath := widgetPath
	sys_classPath := classPath
	sys_type_ := type_
}

func RcGetThemeDir() {}

func RcParse(filename string) {
	sys_filename := filename
}

func RcParseColor(scanner *glib.Scanner) {
	sys_scanner := scanner.ToC()
}

func RcParsePriority(scanner *glib.Scanner, priority *int) {
	sys_scanner := scanner.ToC()
	sys_priority := priority
}

func RcParseState(scanner *glib.Scanner) {
	sys_scanner := scanner.ToC()
}

func RcParseString(rcString string) {
	sys_rcString := rcString
}

func RcReparseAll() {}

func RcReparseAllForSettings(settings *Settings, forceLoad bool) {
	sys_settings := settings.ToC()
	sys_forceLoad := forceLoad
}

func RcResetStyles(settings *Settings) {
	sys_settings := settings.ToC()
}

func RcScannerNew() {}

// UNSUPPORTED : gtk_rc_set_default_files : parameter 'filenames' is array parameter without length parameter

func SelectionAddTarget(widget *Widget, selection gdk.Atom, target gdk.Atom, info uint) {
	sys_widget := widget.ToC()
	sys_selection := selection.ToC()
	sys_target := target.ToC()
	sys_info := info
}

// UNSUPPORTED : gtk_selection_add_targets : has array param, targets

func SelectionClearTargets(widget *Widget, selection gdk.Atom) {
	sys_widget := widget.ToC()
	sys_selection := selection.ToC()
}

func SelectionConvert(widget *Widget, selection gdk.Atom, target gdk.Atom, time uint32) {
	sys_widget := widget.ToC()
	sys_selection := selection.ToC()
	sys_target := target.ToC()
	sys_time := time
}

func SelectionOwnerSet(widget *Widget, selection gdk.Atom, time uint32) {
	sys_widget := widget.ToC()
	sys_selection := selection.ToC()
	sys_time := time
}

func SelectionOwnerSetForDisplay(display *gdk.Display, widget *Widget, selection gdk.Atom, time uint32) {
	sys_display := display.ToC()
	sys_widget := widget.ToC()
	sys_selection := selection.ToC()
	sys_time := time
}

func SelectionRemoveAll(widget *Widget) {
	sys_widget := widget.ToC()
}

func SetDebugFlags(flags uint) {
	sys_flags := flags
}

func ShowAboutDialog(parent *Window, firstPropertyName string) {
	sys_parent := parent.ToC()
	sys_firstPropertyName := firstPropertyName
}

// UNSUPPORTED : gtk_stock_add : has array param, items

// UNSUPPORTED : gtk_stock_add_static : has array param, items

func StockListIds() {}

func StockLookup(stockId string) {
	sys_stockId := stockId
}

// UNSUPPORTED : gtk_stock_set_translate_func : parameter 'func' is callback

// UNSUPPORTED : gtk_target_table_free : has array param, targets

// UNSUPPORTED : gtk_targets_include_image : has array param, targets

// UNSUPPORTED : gtk_targets_include_rich_text : has array param, targets

// UNSUPPORTED : gtk_targets_include_text : has array param, targets

// UNSUPPORTED : gtk_targets_include_uri : has array param, targets

// UNSUPPORTED : gtk_test_init : has array param, argvp

func TreeGetRowDragData(selectionData *SelectionData) {
	sys_selectionData := selectionData.ToC()
}

// UNSUPPORTED : gtk_tree_row_reference_reordered : parameter 'new_order' is array parameter without length parameter

func TreeSetRowDragData(selectionData *SelectionData, treeModel *TreeModel, path *TreePath) {
	sys_selectionData := selectionData.ToC()
	sys_treeModel := treeModel.ToC()
	sys_path := path.ToC()
}

func True() {}

// AboutDialogClass is a representation of the C record GtkAboutDialogClass.
type AboutDialogClass struct {
	native unsafe.Pointer
}

func (recv *AboutDialogClass) ToC() unsafe.Pointer {
	return recv.native
}

// AboutDialogPrivate is a representation of the C record GtkAboutDialogPrivate.
type AboutDialogPrivate struct {
	native unsafe.Pointer
}

func (recv *AboutDialogPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// AccelGroupClass is a representation of the C record GtkAccelGroupClass.
type AccelGroupClass struct {
	native unsafe.Pointer
}

func (recv *AccelGroupClass) ToC() unsafe.Pointer {
	return recv.native
}

// AccelGroupEntry is a representation of the C record GtkAccelGroupEntry.
type AccelGroupEntry struct {
	native unsafe.Pointer
}

func (recv *AccelGroupEntry) ToC() unsafe.Pointer {
	return recv.native
}

// AccelGroupPrivate is a representation of the C record GtkAccelGroupPrivate.
type AccelGroupPrivate struct {
	native unsafe.Pointer
}

func (recv *AccelGroupPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// AccelKey is a representation of the C record GtkAccelKey.
type AccelKey struct {
	native unsafe.Pointer
}

func (recv *AccelKey) ToC() unsafe.Pointer {
	return recv.native
}

// AccelLabelClass is a representation of the C record GtkAccelLabelClass.
type AccelLabelClass struct {
	native unsafe.Pointer
}

func (recv *AccelLabelClass) ToC() unsafe.Pointer {
	return recv.native
}

// AccelLabelPrivate is a representation of the C record GtkAccelLabelPrivate.
type AccelLabelPrivate struct {
	native unsafe.Pointer
}

func (recv *AccelLabelPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// AccelMapClass is a representation of the C record GtkAccelMapClass.
type AccelMapClass struct {
	native unsafe.Pointer
}

func (recv *AccelMapClass) ToC() unsafe.Pointer {
	return recv.native
}

// AccessibleClass is a representation of the C record GtkAccessibleClass.
type AccessibleClass struct {
	native unsafe.Pointer
}

func (recv *AccessibleClass) ToC() unsafe.Pointer {
	return recv.native
}

// AccessiblePrivate is a representation of the C record GtkAccessiblePrivate.
type AccessiblePrivate struct {
	native unsafe.Pointer
}

func (recv *AccessiblePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ActionBarClass is a representation of the C record GtkActionBarClass.
type ActionBarClass struct {
	native unsafe.Pointer
}

func (recv *ActionBarClass) ToC() unsafe.Pointer {
	return recv.native
}

// ActionBarPrivate is a representation of the C record GtkActionBarPrivate.
type ActionBarPrivate struct {
	native unsafe.Pointer
}

func (recv *ActionBarPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ActionClass is a representation of the C record GtkActionClass.
type ActionClass struct {
	native unsafe.Pointer
}

func (recv *ActionClass) ToC() unsafe.Pointer {
	return recv.native
}

// ActionEntry is a representation of the C record GtkActionEntry.
type ActionEntry struct {
	native unsafe.Pointer
}

func (recv *ActionEntry) ToC() unsafe.Pointer {
	return recv.native
}

// ActionGroupClass is a representation of the C record GtkActionGroupClass.
type ActionGroupClass struct {
	native unsafe.Pointer
}

func (recv *ActionGroupClass) ToC() unsafe.Pointer {
	return recv.native
}

// ActionGroupPrivate is a representation of the C record GtkActionGroupPrivate.
type ActionGroupPrivate struct {
	native unsafe.Pointer
}

func (recv *ActionGroupPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ActionPrivate is a representation of the C record GtkActionPrivate.
type ActionPrivate struct {
	native unsafe.Pointer
}

func (recv *ActionPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ActionableInterface is a representation of the C record GtkActionableInterface.
type ActionableInterface struct {
	native unsafe.Pointer
}

func (recv *ActionableInterface) ToC() unsafe.Pointer {
	return recv.native
}

// AdjustmentClass is a representation of the C record GtkAdjustmentClass.
type AdjustmentClass struct {
	native unsafe.Pointer
}

func (recv *AdjustmentClass) ToC() unsafe.Pointer {
	return recv.native
}

// AdjustmentPrivate is a representation of the C record GtkAdjustmentPrivate.
type AdjustmentPrivate struct {
	native unsafe.Pointer
}

func (recv *AdjustmentPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// AlignmentClass is a representation of the C record GtkAlignmentClass.
type AlignmentClass struct {
	native unsafe.Pointer
}

func (recv *AlignmentClass) ToC() unsafe.Pointer {
	return recv.native
}

// AlignmentPrivate is a representation of the C record GtkAlignmentPrivate.
type AlignmentPrivate struct {
	native unsafe.Pointer
}

func (recv *AlignmentPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// AppChooserButtonClass is a representation of the C record GtkAppChooserButtonClass.
type AppChooserButtonClass struct {
	native unsafe.Pointer
}

func (recv *AppChooserButtonClass) ToC() unsafe.Pointer {
	return recv.native
}

// AppChooserButtonPrivate is a representation of the C record GtkAppChooserButtonPrivate.
type AppChooserButtonPrivate struct {
	native unsafe.Pointer
}

func (recv *AppChooserButtonPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// AppChooserDialogClass is a representation of the C record GtkAppChooserDialogClass.
type AppChooserDialogClass struct {
	native unsafe.Pointer
}

func (recv *AppChooserDialogClass) ToC() unsafe.Pointer {
	return recv.native
}

// AppChooserDialogPrivate is a representation of the C record GtkAppChooserDialogPrivate.
type AppChooserDialogPrivate struct {
	native unsafe.Pointer
}

func (recv *AppChooserDialogPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// AppChooserWidgetClass is a representation of the C record GtkAppChooserWidgetClass.
type AppChooserWidgetClass struct {
	native unsafe.Pointer
}

func (recv *AppChooserWidgetClass) ToC() unsafe.Pointer {
	return recv.native
}

// AppChooserWidgetPrivate is a representation of the C record GtkAppChooserWidgetPrivate.
type AppChooserWidgetPrivate struct {
	native unsafe.Pointer
}

func (recv *AppChooserWidgetPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ApplicationClass is a representation of the C record GtkApplicationClass.
type ApplicationClass struct {
	native unsafe.Pointer
}

func (recv *ApplicationClass) ToC() unsafe.Pointer {
	return recv.native
}

// ApplicationPrivate is a representation of the C record GtkApplicationPrivate.
type ApplicationPrivate struct {
	native unsafe.Pointer
}

func (recv *ApplicationPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ApplicationWindowClass is a representation of the C record GtkApplicationWindowClass.
type ApplicationWindowClass struct {
	native unsafe.Pointer
}

func (recv *ApplicationWindowClass) ToC() unsafe.Pointer {
	return recv.native
}

// ApplicationWindowPrivate is a representation of the C record GtkApplicationWindowPrivate.
type ApplicationWindowPrivate struct {
	native unsafe.Pointer
}

func (recv *ApplicationWindowPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ArrowAccessibleClass is a representation of the C record GtkArrowAccessibleClass.
type ArrowAccessibleClass struct {
	native unsafe.Pointer
}

func (recv *ArrowAccessibleClass) ToC() unsafe.Pointer {
	return recv.native
}

// ArrowAccessiblePrivate is a representation of the C record GtkArrowAccessiblePrivate.
type ArrowAccessiblePrivate struct {
	native unsafe.Pointer
}

func (recv *ArrowAccessiblePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ArrowClass is a representation of the C record GtkArrowClass.
type ArrowClass struct {
	native unsafe.Pointer
}

func (recv *ArrowClass) ToC() unsafe.Pointer {
	return recv.native
}

// ArrowPrivate is a representation of the C record GtkArrowPrivate.
type ArrowPrivate struct {
	native unsafe.Pointer
}

func (recv *ArrowPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// AspectFrameClass is a representation of the C record GtkAspectFrameClass.
type AspectFrameClass struct {
	native unsafe.Pointer
}

func (recv *AspectFrameClass) ToC() unsafe.Pointer {
	return recv.native
}

// AspectFramePrivate is a representation of the C record GtkAspectFramePrivate.
type AspectFramePrivate struct {
	native unsafe.Pointer
}

func (recv *AspectFramePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// AssistantClass is a representation of the C record GtkAssistantClass.
type AssistantClass struct {
	native unsafe.Pointer
}

func (recv *AssistantClass) ToC() unsafe.Pointer {
	return recv.native
}

// AssistantPrivate is a representation of the C record GtkAssistantPrivate.
type AssistantPrivate struct {
	native unsafe.Pointer
}

func (recv *AssistantPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// BinClass is a representation of the C record GtkBinClass.
type BinClass struct {
	native unsafe.Pointer
}

func (recv *BinClass) ToC() unsafe.Pointer {
	return recv.native
}

// BinPrivate is a representation of the C record GtkBinPrivate.
type BinPrivate struct {
	native unsafe.Pointer
}

func (recv *BinPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// BindingArg is a representation of the C record GtkBindingArg.
type BindingArg struct {
	native unsafe.Pointer
}

func (recv *BindingArg) ToC() unsafe.Pointer {
	return recv.native
}

// BindingEntry is a representation of the C record GtkBindingEntry.
type BindingEntry struct {
	native unsafe.Pointer
}

func (recv *BindingEntry) ToC() unsafe.Pointer {
	return recv.native
}

// BindingSet is a representation of the C record GtkBindingSet.
type BindingSet struct {
	native unsafe.Pointer
}

func (recv *BindingSet) ToC() unsafe.Pointer {
	return recv.native
}

// BindingSignal is a representation of the C record GtkBindingSignal.
type BindingSignal struct {
	native unsafe.Pointer
}

func (recv *BindingSignal) ToC() unsafe.Pointer {
	return recv.native
}

// BooleanCellAccessibleClass is a representation of the C record GtkBooleanCellAccessibleClass.
type BooleanCellAccessibleClass struct {
	native unsafe.Pointer
}

func (recv *BooleanCellAccessibleClass) ToC() unsafe.Pointer {
	return recv.native
}

// BooleanCellAccessiblePrivate is a representation of the C record GtkBooleanCellAccessiblePrivate.
type BooleanCellAccessiblePrivate struct {
	native unsafe.Pointer
}

func (recv *BooleanCellAccessiblePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// Border is a representation of the C record GtkBorder.
type Border struct {
	native unsafe.Pointer
}

func (recv *Border) ToC() unsafe.Pointer {
	return recv.native
}

// BoxClass is a representation of the C record GtkBoxClass.
type BoxClass struct {
	native unsafe.Pointer
}

func (recv *BoxClass) ToC() unsafe.Pointer {
	return recv.native
}

// BoxPrivate is a representation of the C record GtkBoxPrivate.
type BoxPrivate struct {
	native unsafe.Pointer
}

func (recv *BoxPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// BuildableIface is a representation of the C record GtkBuildableIface.
type BuildableIface struct {
	native unsafe.Pointer
}

func (recv *BuildableIface) ToC() unsafe.Pointer {
	return recv.native
}

// BuilderClass is a representation of the C record GtkBuilderClass.
type BuilderClass struct {
	native unsafe.Pointer
}

func (recv *BuilderClass) ToC() unsafe.Pointer {
	return recv.native
}

// BuilderPrivate is a representation of the C record GtkBuilderPrivate.
type BuilderPrivate struct {
	native unsafe.Pointer
}

func (recv *BuilderPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ButtonAccessibleClass is a representation of the C record GtkButtonAccessibleClass.
type ButtonAccessibleClass struct {
	native unsafe.Pointer
}

func (recv *ButtonAccessibleClass) ToC() unsafe.Pointer {
	return recv.native
}

// ButtonAccessiblePrivate is a representation of the C record GtkButtonAccessiblePrivate.
type ButtonAccessiblePrivate struct {
	native unsafe.Pointer
}

func (recv *ButtonAccessiblePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ButtonBoxClass is a representation of the C record GtkButtonBoxClass.
type ButtonBoxClass struct {
	native unsafe.Pointer
}

func (recv *ButtonBoxClass) ToC() unsafe.Pointer {
	return recv.native
}

// ButtonBoxPrivate is a representation of the C record GtkButtonBoxPrivate.
type ButtonBoxPrivate struct {
	native unsafe.Pointer
}

func (recv *ButtonBoxPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ButtonClass is a representation of the C record GtkButtonClass.
type ButtonClass struct {
	native unsafe.Pointer
}

func (recv *ButtonClass) ToC() unsafe.Pointer {
	return recv.native
}

// ButtonPrivate is a representation of the C record GtkButtonPrivate.
type ButtonPrivate struct {
	native unsafe.Pointer
}

func (recv *ButtonPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// CalendarClass is a representation of the C record GtkCalendarClass.
type CalendarClass struct {
	native unsafe.Pointer
}

func (recv *CalendarClass) ToC() unsafe.Pointer {
	return recv.native
}

// CalendarPrivate is a representation of the C record GtkCalendarPrivate.
type CalendarPrivate struct {
	native unsafe.Pointer
}

func (recv *CalendarPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// CellAccessibleClass is a representation of the C record GtkCellAccessibleClass.
type CellAccessibleClass struct {
	native unsafe.Pointer
}

func (recv *CellAccessibleClass) ToC() unsafe.Pointer {
	return recv.native
}

// CellAccessibleParentIface is a representation of the C record GtkCellAccessibleParentIface.
type CellAccessibleParentIface struct {
	native unsafe.Pointer
}

func (recv *CellAccessibleParentIface) ToC() unsafe.Pointer {
	return recv.native
}

// CellAccessiblePrivate is a representation of the C record GtkCellAccessiblePrivate.
type CellAccessiblePrivate struct {
	native unsafe.Pointer
}

func (recv *CellAccessiblePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// CellAreaBoxClass is a representation of the C record GtkCellAreaBoxClass.
type CellAreaBoxClass struct {
	native unsafe.Pointer
}

func (recv *CellAreaBoxClass) ToC() unsafe.Pointer {
	return recv.native
}

// CellAreaBoxPrivate is a representation of the C record GtkCellAreaBoxPrivate.
type CellAreaBoxPrivate struct {
	native unsafe.Pointer
}

func (recv *CellAreaBoxPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// CellAreaClass is a representation of the C record GtkCellAreaClass.
type CellAreaClass struct {
	native unsafe.Pointer
}

func (recv *CellAreaClass) ToC() unsafe.Pointer {
	return recv.native
}

// CellAreaContextClass is a representation of the C record GtkCellAreaContextClass.
type CellAreaContextClass struct {
	native unsafe.Pointer
}

func (recv *CellAreaContextClass) ToC() unsafe.Pointer {
	return recv.native
}

// CellAreaContextPrivate is a representation of the C record GtkCellAreaContextPrivate.
type CellAreaContextPrivate struct {
	native unsafe.Pointer
}

func (recv *CellAreaContextPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// CellAreaPrivate is a representation of the C record GtkCellAreaPrivate.
type CellAreaPrivate struct {
	native unsafe.Pointer
}

func (recv *CellAreaPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// CellEditableIface is a representation of the C record GtkCellEditableIface.
type CellEditableIface struct {
	native unsafe.Pointer
}

func (recv *CellEditableIface) ToC() unsafe.Pointer {
	return recv.native
}

// CellLayoutIface is a representation of the C record GtkCellLayoutIface.
type CellLayoutIface struct {
	native unsafe.Pointer
}

func (recv *CellLayoutIface) ToC() unsafe.Pointer {
	return recv.native
}

// CellRendererAccelClass is a representation of the C record GtkCellRendererAccelClass.
type CellRendererAccelClass struct {
	native unsafe.Pointer
}

func (recv *CellRendererAccelClass) ToC() unsafe.Pointer {
	return recv.native
}

// CellRendererAccelPrivate is a representation of the C record GtkCellRendererAccelPrivate.
type CellRendererAccelPrivate struct {
	native unsafe.Pointer
}

func (recv *CellRendererAccelPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// CellRendererClass is a representation of the C record GtkCellRendererClass.
type CellRendererClass struct {
	native unsafe.Pointer
}

func (recv *CellRendererClass) ToC() unsafe.Pointer {
	return recv.native
}

// CellRendererClassPrivate is a representation of the C record GtkCellRendererClassPrivate.
type CellRendererClassPrivate struct {
	native unsafe.Pointer
}

func (recv *CellRendererClassPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// CellRendererComboClass is a representation of the C record GtkCellRendererComboClass.
type CellRendererComboClass struct {
	native unsafe.Pointer
}

func (recv *CellRendererComboClass) ToC() unsafe.Pointer {
	return recv.native
}

// CellRendererComboPrivate is a representation of the C record GtkCellRendererComboPrivate.
type CellRendererComboPrivate struct {
	native unsafe.Pointer
}

func (recv *CellRendererComboPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// CellRendererPixbufClass is a representation of the C record GtkCellRendererPixbufClass.
type CellRendererPixbufClass struct {
	native unsafe.Pointer
}

func (recv *CellRendererPixbufClass) ToC() unsafe.Pointer {
	return recv.native
}

// CellRendererPixbufPrivate is a representation of the C record GtkCellRendererPixbufPrivate.
type CellRendererPixbufPrivate struct {
	native unsafe.Pointer
}

func (recv *CellRendererPixbufPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// CellRendererPrivate is a representation of the C record GtkCellRendererPrivate.
type CellRendererPrivate struct {
	native unsafe.Pointer
}

func (recv *CellRendererPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// CellRendererProgressClass is a representation of the C record GtkCellRendererProgressClass.
type CellRendererProgressClass struct {
	native unsafe.Pointer
}

func (recv *CellRendererProgressClass) ToC() unsafe.Pointer {
	return recv.native
}

// CellRendererProgressPrivate is a representation of the C record GtkCellRendererProgressPrivate.
type CellRendererProgressPrivate struct {
	native unsafe.Pointer
}

func (recv *CellRendererProgressPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// CellRendererSpinClass is a representation of the C record GtkCellRendererSpinClass.
type CellRendererSpinClass struct {
	native unsafe.Pointer
}

func (recv *CellRendererSpinClass) ToC() unsafe.Pointer {
	return recv.native
}

// CellRendererSpinPrivate is a representation of the C record GtkCellRendererSpinPrivate.
type CellRendererSpinPrivate struct {
	native unsafe.Pointer
}

func (recv *CellRendererSpinPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// CellRendererSpinnerClass is a representation of the C record GtkCellRendererSpinnerClass.
type CellRendererSpinnerClass struct {
	native unsafe.Pointer
}

func (recv *CellRendererSpinnerClass) ToC() unsafe.Pointer {
	return recv.native
}

// CellRendererSpinnerPrivate is a representation of the C record GtkCellRendererSpinnerPrivate.
type CellRendererSpinnerPrivate struct {
	native unsafe.Pointer
}

func (recv *CellRendererSpinnerPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// CellRendererTextClass is a representation of the C record GtkCellRendererTextClass.
type CellRendererTextClass struct {
	native unsafe.Pointer
}

func (recv *CellRendererTextClass) ToC() unsafe.Pointer {
	return recv.native
}

// CellRendererTextPrivate is a representation of the C record GtkCellRendererTextPrivate.
type CellRendererTextPrivate struct {
	native unsafe.Pointer
}

func (recv *CellRendererTextPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// CellRendererToggleClass is a representation of the C record GtkCellRendererToggleClass.
type CellRendererToggleClass struct {
	native unsafe.Pointer
}

func (recv *CellRendererToggleClass) ToC() unsafe.Pointer {
	return recv.native
}

// CellRendererTogglePrivate is a representation of the C record GtkCellRendererTogglePrivate.
type CellRendererTogglePrivate struct {
	native unsafe.Pointer
}

func (recv *CellRendererTogglePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// CellViewClass is a representation of the C record GtkCellViewClass.
type CellViewClass struct {
	native unsafe.Pointer
}

func (recv *CellViewClass) ToC() unsafe.Pointer {
	return recv.native
}

// CellViewPrivate is a representation of the C record GtkCellViewPrivate.
type CellViewPrivate struct {
	native unsafe.Pointer
}

func (recv *CellViewPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// CheckButtonClass is a representation of the C record GtkCheckButtonClass.
type CheckButtonClass struct {
	native unsafe.Pointer
}

func (recv *CheckButtonClass) ToC() unsafe.Pointer {
	return recv.native
}

// CheckMenuItemAccessibleClass is a representation of the C record GtkCheckMenuItemAccessibleClass.
type CheckMenuItemAccessibleClass struct {
	native unsafe.Pointer
}

func (recv *CheckMenuItemAccessibleClass) ToC() unsafe.Pointer {
	return recv.native
}

// CheckMenuItemAccessiblePrivate is a representation of the C record GtkCheckMenuItemAccessiblePrivate.
type CheckMenuItemAccessiblePrivate struct {
	native unsafe.Pointer
}

func (recv *CheckMenuItemAccessiblePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// CheckMenuItemClass is a representation of the C record GtkCheckMenuItemClass.
type CheckMenuItemClass struct {
	native unsafe.Pointer
}

func (recv *CheckMenuItemClass) ToC() unsafe.Pointer {
	return recv.native
}

// CheckMenuItemPrivate is a representation of the C record GtkCheckMenuItemPrivate.
type CheckMenuItemPrivate struct {
	native unsafe.Pointer
}

func (recv *CheckMenuItemPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ColorButtonClass is a representation of the C record GtkColorButtonClass.
type ColorButtonClass struct {
	native unsafe.Pointer
}

func (recv *ColorButtonClass) ToC() unsafe.Pointer {
	return recv.native
}

// ColorButtonPrivate is a representation of the C record GtkColorButtonPrivate.
type ColorButtonPrivate struct {
	native unsafe.Pointer
}

func (recv *ColorButtonPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ColorChooserDialogClass is a representation of the C record GtkColorChooserDialogClass.
type ColorChooserDialogClass struct {
	native unsafe.Pointer
}

func (recv *ColorChooserDialogClass) ToC() unsafe.Pointer {
	return recv.native
}

// ColorChooserDialogPrivate is a representation of the C record GtkColorChooserDialogPrivate.
type ColorChooserDialogPrivate struct {
	native unsafe.Pointer
}

func (recv *ColorChooserDialogPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ColorChooserInterface is a representation of the C record GtkColorChooserInterface.
type ColorChooserInterface struct {
	native unsafe.Pointer
}

func (recv *ColorChooserInterface) ToC() unsafe.Pointer {
	return recv.native
}

// ColorChooserWidgetClass is a representation of the C record GtkColorChooserWidgetClass.
type ColorChooserWidgetClass struct {
	native unsafe.Pointer
}

func (recv *ColorChooserWidgetClass) ToC() unsafe.Pointer {
	return recv.native
}

// ColorChooserWidgetPrivate is a representation of the C record GtkColorChooserWidgetPrivate.
type ColorChooserWidgetPrivate struct {
	native unsafe.Pointer
}

func (recv *ColorChooserWidgetPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ColorSelectionClass is a representation of the C record GtkColorSelectionClass.
type ColorSelectionClass struct {
	native unsafe.Pointer
}

func (recv *ColorSelectionClass) ToC() unsafe.Pointer {
	return recv.native
}

// ColorSelectionDialogClass is a representation of the C record GtkColorSelectionDialogClass.
type ColorSelectionDialogClass struct {
	native unsafe.Pointer
}

func (recv *ColorSelectionDialogClass) ToC() unsafe.Pointer {
	return recv.native
}

// ColorSelectionDialogPrivate is a representation of the C record GtkColorSelectionDialogPrivate.
type ColorSelectionDialogPrivate struct {
	native unsafe.Pointer
}

func (recv *ColorSelectionDialogPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ColorSelectionPrivate is a representation of the C record GtkColorSelectionPrivate.
type ColorSelectionPrivate struct {
	native unsafe.Pointer
}

func (recv *ColorSelectionPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ComboBoxAccessibleClass is a representation of the C record GtkComboBoxAccessibleClass.
type ComboBoxAccessibleClass struct {
	native unsafe.Pointer
}

func (recv *ComboBoxAccessibleClass) ToC() unsafe.Pointer {
	return recv.native
}

// ComboBoxAccessiblePrivate is a representation of the C record GtkComboBoxAccessiblePrivate.
type ComboBoxAccessiblePrivate struct {
	native unsafe.Pointer
}

func (recv *ComboBoxAccessiblePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ComboBoxClass is a representation of the C record GtkComboBoxClass.
type ComboBoxClass struct {
	native unsafe.Pointer
}

func (recv *ComboBoxClass) ToC() unsafe.Pointer {
	return recv.native
}

// ComboBoxPrivate is a representation of the C record GtkComboBoxPrivate.
type ComboBoxPrivate struct {
	native unsafe.Pointer
}

func (recv *ComboBoxPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ComboBoxTextClass is a representation of the C record GtkComboBoxTextClass.
type ComboBoxTextClass struct {
	native unsafe.Pointer
}

func (recv *ComboBoxTextClass) ToC() unsafe.Pointer {
	return recv.native
}

// ComboBoxTextPrivate is a representation of the C record GtkComboBoxTextPrivate.
type ComboBoxTextPrivate struct {
	native unsafe.Pointer
}

func (recv *ComboBoxTextPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ContainerAccessibleClass is a representation of the C record GtkContainerAccessibleClass.
type ContainerAccessibleClass struct {
	native unsafe.Pointer
}

func (recv *ContainerAccessibleClass) ToC() unsafe.Pointer {
	return recv.native
}

// ContainerAccessiblePrivate is a representation of the C record GtkContainerAccessiblePrivate.
type ContainerAccessiblePrivate struct {
	native unsafe.Pointer
}

func (recv *ContainerAccessiblePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ContainerCellAccessibleClass is a representation of the C record GtkContainerCellAccessibleClass.
type ContainerCellAccessibleClass struct {
	native unsafe.Pointer
}

func (recv *ContainerCellAccessibleClass) ToC() unsafe.Pointer {
	return recv.native
}

// ContainerCellAccessiblePrivate is a representation of the C record GtkContainerCellAccessiblePrivate.
type ContainerCellAccessiblePrivate struct {
	native unsafe.Pointer
}

func (recv *ContainerCellAccessiblePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ContainerClass is a representation of the C record GtkContainerClass.
type ContainerClass struct {
	native unsafe.Pointer
}

func (recv *ContainerClass) ToC() unsafe.Pointer {
	return recv.native
}

// ContainerPrivate is a representation of the C record GtkContainerPrivate.
type ContainerPrivate struct {
	native unsafe.Pointer
}

func (recv *ContainerPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// CssProviderClass is a representation of the C record GtkCssProviderClass.
type CssProviderClass struct {
	native unsafe.Pointer
}

func (recv *CssProviderClass) ToC() unsafe.Pointer {
	return recv.native
}

// CssProviderPrivate is a representation of the C record GtkCssProviderPrivate.
type CssProviderPrivate struct {
	native unsafe.Pointer
}

func (recv *CssProviderPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// DialogClass is a representation of the C record GtkDialogClass.
type DialogClass struct {
	native unsafe.Pointer
}

func (recv *DialogClass) ToC() unsafe.Pointer {
	return recv.native
}

// DialogPrivate is a representation of the C record GtkDialogPrivate.
type DialogPrivate struct {
	native unsafe.Pointer
}

func (recv *DialogPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// DrawingAreaClass is a representation of the C record GtkDrawingAreaClass.
type DrawingAreaClass struct {
	native unsafe.Pointer
}

func (recv *DrawingAreaClass) ToC() unsafe.Pointer {
	return recv.native
}

// EditableInterface is a representation of the C record GtkEditableInterface.
type EditableInterface struct {
	native unsafe.Pointer
}

func (recv *EditableInterface) ToC() unsafe.Pointer {
	return recv.native
}

// EntryAccessibleClass is a representation of the C record GtkEntryAccessibleClass.
type EntryAccessibleClass struct {
	native unsafe.Pointer
}

func (recv *EntryAccessibleClass) ToC() unsafe.Pointer {
	return recv.native
}

// EntryAccessiblePrivate is a representation of the C record GtkEntryAccessiblePrivate.
type EntryAccessiblePrivate struct {
	native unsafe.Pointer
}

func (recv *EntryAccessiblePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// EntryBufferClass is a representation of the C record GtkEntryBufferClass.
type EntryBufferClass struct {
	native unsafe.Pointer
}

func (recv *EntryBufferClass) ToC() unsafe.Pointer {
	return recv.native
}

// EntryBufferPrivate is a representation of the C record GtkEntryBufferPrivate.
type EntryBufferPrivate struct {
	native unsafe.Pointer
}

func (recv *EntryBufferPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// EntryClass is a representation of the C record GtkEntryClass.
type EntryClass struct {
	native unsafe.Pointer
}

func (recv *EntryClass) ToC() unsafe.Pointer {
	return recv.native
}

// EntryCompletionClass is a representation of the C record GtkEntryCompletionClass.
type EntryCompletionClass struct {
	native unsafe.Pointer
}

func (recv *EntryCompletionClass) ToC() unsafe.Pointer {
	return recv.native
}

// EntryCompletionPrivate is a representation of the C record GtkEntryCompletionPrivate.
type EntryCompletionPrivate struct {
	native unsafe.Pointer
}

func (recv *EntryCompletionPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// EntryPrivate is a representation of the C record GtkEntryPrivate.
type EntryPrivate struct {
	native unsafe.Pointer
}

func (recv *EntryPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// EventBoxClass is a representation of the C record GtkEventBoxClass.
type EventBoxClass struct {
	native unsafe.Pointer
}

func (recv *EventBoxClass) ToC() unsafe.Pointer {
	return recv.native
}

// EventBoxPrivate is a representation of the C record GtkEventBoxPrivate.
type EventBoxPrivate struct {
	native unsafe.Pointer
}

func (recv *EventBoxPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// EventControllerClass is a representation of the C record GtkEventControllerClass.
type EventControllerClass struct {
	native unsafe.Pointer
}

func (recv *EventControllerClass) ToC() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : EventControllerMotionClass : blacklisted

// UNSUPPORTED : EventControllerScrollClass : blacklisted

// ExpanderAccessibleClass is a representation of the C record GtkExpanderAccessibleClass.
type ExpanderAccessibleClass struct {
	native unsafe.Pointer
}

func (recv *ExpanderAccessibleClass) ToC() unsafe.Pointer {
	return recv.native
}

// ExpanderAccessiblePrivate is a representation of the C record GtkExpanderAccessiblePrivate.
type ExpanderAccessiblePrivate struct {
	native unsafe.Pointer
}

func (recv *ExpanderAccessiblePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ExpanderClass is a representation of the C record GtkExpanderClass.
type ExpanderClass struct {
	native unsafe.Pointer
}

func (recv *ExpanderClass) ToC() unsafe.Pointer {
	return recv.native
}

// ExpanderPrivate is a representation of the C record GtkExpanderPrivate.
type ExpanderPrivate struct {
	native unsafe.Pointer
}

func (recv *ExpanderPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// FileChooserButtonClass is a representation of the C record GtkFileChooserButtonClass.
type FileChooserButtonClass struct {
	native unsafe.Pointer
}

func (recv *FileChooserButtonClass) ToC() unsafe.Pointer {
	return recv.native
}

// FileChooserButtonPrivate is a representation of the C record GtkFileChooserButtonPrivate.
type FileChooserButtonPrivate struct {
	native unsafe.Pointer
}

func (recv *FileChooserButtonPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// FileChooserDialogClass is a representation of the C record GtkFileChooserDialogClass.
type FileChooserDialogClass struct {
	native unsafe.Pointer
}

func (recv *FileChooserDialogClass) ToC() unsafe.Pointer {
	return recv.native
}

// FileChooserDialogPrivate is a representation of the C record GtkFileChooserDialogPrivate.
type FileChooserDialogPrivate struct {
	native unsafe.Pointer
}

func (recv *FileChooserDialogPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// FileChooserNativeClass is a representation of the C record GtkFileChooserNativeClass.
type FileChooserNativeClass struct {
	native unsafe.Pointer
}

func (recv *FileChooserNativeClass) ToC() unsafe.Pointer {
	return recv.native
}

// FileChooserWidgetClass is a representation of the C record GtkFileChooserWidgetClass.
type FileChooserWidgetClass struct {
	native unsafe.Pointer
}

func (recv *FileChooserWidgetClass) ToC() unsafe.Pointer {
	return recv.native
}

// FileChooserWidgetPrivate is a representation of the C record GtkFileChooserWidgetPrivate.
type FileChooserWidgetPrivate struct {
	native unsafe.Pointer
}

func (recv *FileChooserWidgetPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// FileFilterInfo is a representation of the C record GtkFileFilterInfo.
type FileFilterInfo struct {
	native unsafe.Pointer
}

func (recv *FileFilterInfo) ToC() unsafe.Pointer {
	return recv.native
}

// FixedChild is a representation of the C record GtkFixedChild.
type FixedChild struct {
	native unsafe.Pointer
}

func (recv *FixedChild) ToC() unsafe.Pointer {
	return recv.native
}

// FixedClass is a representation of the C record GtkFixedClass.
type FixedClass struct {
	native unsafe.Pointer
}

func (recv *FixedClass) ToC() unsafe.Pointer {
	return recv.native
}

// FixedPrivate is a representation of the C record GtkFixedPrivate.
type FixedPrivate struct {
	native unsafe.Pointer
}

func (recv *FixedPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// FlowBoxAccessibleClass is a representation of the C record GtkFlowBoxAccessibleClass.
type FlowBoxAccessibleClass struct {
	native unsafe.Pointer
}

func (recv *FlowBoxAccessibleClass) ToC() unsafe.Pointer {
	return recv.native
}

// FlowBoxAccessiblePrivate is a representation of the C record GtkFlowBoxAccessiblePrivate.
type FlowBoxAccessiblePrivate struct {
	native unsafe.Pointer
}

func (recv *FlowBoxAccessiblePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// FlowBoxChildAccessibleClass is a representation of the C record GtkFlowBoxChildAccessibleClass.
type FlowBoxChildAccessibleClass struct {
	native unsafe.Pointer
}

func (recv *FlowBoxChildAccessibleClass) ToC() unsafe.Pointer {
	return recv.native
}

// FlowBoxChildClass is a representation of the C record GtkFlowBoxChildClass.
type FlowBoxChildClass struct {
	native unsafe.Pointer
}

func (recv *FlowBoxChildClass) ToC() unsafe.Pointer {
	return recv.native
}

// FlowBoxClass is a representation of the C record GtkFlowBoxClass.
type FlowBoxClass struct {
	native unsafe.Pointer
}

func (recv *FlowBoxClass) ToC() unsafe.Pointer {
	return recv.native
}

// FontButtonClass is a representation of the C record GtkFontButtonClass.
type FontButtonClass struct {
	native unsafe.Pointer
}

func (recv *FontButtonClass) ToC() unsafe.Pointer {
	return recv.native
}

// FontButtonPrivate is a representation of the C record GtkFontButtonPrivate.
type FontButtonPrivate struct {
	native unsafe.Pointer
}

func (recv *FontButtonPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// FontChooserDialogClass is a representation of the C record GtkFontChooserDialogClass.
type FontChooserDialogClass struct {
	native unsafe.Pointer
}

func (recv *FontChooserDialogClass) ToC() unsafe.Pointer {
	return recv.native
}

// FontChooserDialogPrivate is a representation of the C record GtkFontChooserDialogPrivate.
type FontChooserDialogPrivate struct {
	native unsafe.Pointer
}

func (recv *FontChooserDialogPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// FontChooserIface is a representation of the C record GtkFontChooserIface.
type FontChooserIface struct {
	native unsafe.Pointer
}

func (recv *FontChooserIface) ToC() unsafe.Pointer {
	return recv.native
}

// FontChooserWidgetClass is a representation of the C record GtkFontChooserWidgetClass.
type FontChooserWidgetClass struct {
	native unsafe.Pointer
}

func (recv *FontChooserWidgetClass) ToC() unsafe.Pointer {
	return recv.native
}

// FontChooserWidgetPrivate is a representation of the C record GtkFontChooserWidgetPrivate.
type FontChooserWidgetPrivate struct {
	native unsafe.Pointer
}

func (recv *FontChooserWidgetPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// FontSelectionClass is a representation of the C record GtkFontSelectionClass.
type FontSelectionClass struct {
	native unsafe.Pointer
}

func (recv *FontSelectionClass) ToC() unsafe.Pointer {
	return recv.native
}

// FontSelectionDialogClass is a representation of the C record GtkFontSelectionDialogClass.
type FontSelectionDialogClass struct {
	native unsafe.Pointer
}

func (recv *FontSelectionDialogClass) ToC() unsafe.Pointer {
	return recv.native
}

// FontSelectionDialogPrivate is a representation of the C record GtkFontSelectionDialogPrivate.
type FontSelectionDialogPrivate struct {
	native unsafe.Pointer
}

func (recv *FontSelectionDialogPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// FontSelectionPrivate is a representation of the C record GtkFontSelectionPrivate.
type FontSelectionPrivate struct {
	native unsafe.Pointer
}

func (recv *FontSelectionPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// FrameAccessibleClass is a representation of the C record GtkFrameAccessibleClass.
type FrameAccessibleClass struct {
	native unsafe.Pointer
}

func (recv *FrameAccessibleClass) ToC() unsafe.Pointer {
	return recv.native
}

// FrameAccessiblePrivate is a representation of the C record GtkFrameAccessiblePrivate.
type FrameAccessiblePrivate struct {
	native unsafe.Pointer
}

func (recv *FrameAccessiblePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// FrameClass is a representation of the C record GtkFrameClass.
type FrameClass struct {
	native unsafe.Pointer
}

func (recv *FrameClass) ToC() unsafe.Pointer {
	return recv.native
}

// FramePrivate is a representation of the C record GtkFramePrivate.
type FramePrivate struct {
	native unsafe.Pointer
}

func (recv *FramePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// GestureClass is a representation of the C record GtkGestureClass.
type GestureClass struct {
	native unsafe.Pointer
}

func (recv *GestureClass) ToC() unsafe.Pointer {
	return recv.native
}

// GestureDragClass is a representation of the C record GtkGestureDragClass.
type GestureDragClass struct {
	native unsafe.Pointer
}

func (recv *GestureDragClass) ToC() unsafe.Pointer {
	return recv.native
}

// GestureLongPressClass is a representation of the C record GtkGestureLongPressClass.
type GestureLongPressClass struct {
	native unsafe.Pointer
}

func (recv *GestureLongPressClass) ToC() unsafe.Pointer {
	return recv.native
}

// GestureMultiPressClass is a representation of the C record GtkGestureMultiPressClass.
type GestureMultiPressClass struct {
	native unsafe.Pointer
}

func (recv *GestureMultiPressClass) ToC() unsafe.Pointer {
	return recv.native
}

// GesturePanClass is a representation of the C record GtkGesturePanClass.
type GesturePanClass struct {
	native unsafe.Pointer
}

func (recv *GesturePanClass) ToC() unsafe.Pointer {
	return recv.native
}

// GestureRotateClass is a representation of the C record GtkGestureRotateClass.
type GestureRotateClass struct {
	native unsafe.Pointer
}

func (recv *GestureRotateClass) ToC() unsafe.Pointer {
	return recv.native
}

// GestureSingleClass is a representation of the C record GtkGestureSingleClass.
type GestureSingleClass struct {
	native unsafe.Pointer
}

func (recv *GestureSingleClass) ToC() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : GestureStylusClass : blacklisted

// GestureSwipeClass is a representation of the C record GtkGestureSwipeClass.
type GestureSwipeClass struct {
	native unsafe.Pointer
}

func (recv *GestureSwipeClass) ToC() unsafe.Pointer {
	return recv.native
}

// GestureZoomClass is a representation of the C record GtkGestureZoomClass.
type GestureZoomClass struct {
	native unsafe.Pointer
}

func (recv *GestureZoomClass) ToC() unsafe.Pointer {
	return recv.native
}

// Gradient is a representation of the C record GtkGradient.
type Gradient struct {
	native unsafe.Pointer
}

func (recv *Gradient) ToC() unsafe.Pointer {
	return recv.native
}

// GridClass is a representation of the C record GtkGridClass.
type GridClass struct {
	native unsafe.Pointer
}

func (recv *GridClass) ToC() unsafe.Pointer {
	return recv.native
}

// GridPrivate is a representation of the C record GtkGridPrivate.
type GridPrivate struct {
	native unsafe.Pointer
}

func (recv *GridPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// HBoxClass is a representation of the C record GtkHBoxClass.
type HBoxClass struct {
	native unsafe.Pointer
}

func (recv *HBoxClass) ToC() unsafe.Pointer {
	return recv.native
}

// HButtonBoxClass is a representation of the C record GtkHButtonBoxClass.
type HButtonBoxClass struct {
	native unsafe.Pointer
}

func (recv *HButtonBoxClass) ToC() unsafe.Pointer {
	return recv.native
}

// HPanedClass is a representation of the C record GtkHPanedClass.
type HPanedClass struct {
	native unsafe.Pointer
}

func (recv *HPanedClass) ToC() unsafe.Pointer {
	return recv.native
}

// HSVClass is a representation of the C record GtkHSVClass.
type HSVClass struct {
	native unsafe.Pointer
}

func (recv *HSVClass) ToC() unsafe.Pointer {
	return recv.native
}

// HSVPrivate is a representation of the C record GtkHSVPrivate.
type HSVPrivate struct {
	native unsafe.Pointer
}

func (recv *HSVPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// HScaleClass is a representation of the C record GtkHScaleClass.
type HScaleClass struct {
	native unsafe.Pointer
}

func (recv *HScaleClass) ToC() unsafe.Pointer {
	return recv.native
}

// HScrollbarClass is a representation of the C record GtkHScrollbarClass.
type HScrollbarClass struct {
	native unsafe.Pointer
}

func (recv *HScrollbarClass) ToC() unsafe.Pointer {
	return recv.native
}

// HSeparatorClass is a representation of the C record GtkHSeparatorClass.
type HSeparatorClass struct {
	native unsafe.Pointer
}

func (recv *HSeparatorClass) ToC() unsafe.Pointer {
	return recv.native
}

// HandleBoxClass is a representation of the C record GtkHandleBoxClass.
type HandleBoxClass struct {
	native unsafe.Pointer
}

func (recv *HandleBoxClass) ToC() unsafe.Pointer {
	return recv.native
}

// HandleBoxPrivate is a representation of the C record GtkHandleBoxPrivate.
type HandleBoxPrivate struct {
	native unsafe.Pointer
}

func (recv *HandleBoxPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : HeaderBarAccessibleClass : blacklisted

// UNSUPPORTED : HeaderBarAccessiblePrivate : blacklisted

// HeaderBarClass is a representation of the C record GtkHeaderBarClass.
type HeaderBarClass struct {
	native unsafe.Pointer
}

func (recv *HeaderBarClass) ToC() unsafe.Pointer {
	return recv.native
}

// HeaderBarPrivate is a representation of the C record GtkHeaderBarPrivate.
type HeaderBarPrivate struct {
	native unsafe.Pointer
}

func (recv *HeaderBarPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// IMContextClass is a representation of the C record GtkIMContextClass.
type IMContextClass struct {
	native unsafe.Pointer
}

func (recv *IMContextClass) ToC() unsafe.Pointer {
	return recv.native
}

// IMContextInfo is a representation of the C record GtkIMContextInfo.
type IMContextInfo struct {
	native unsafe.Pointer
}

func (recv *IMContextInfo) ToC() unsafe.Pointer {
	return recv.native
}

// IMContextSimpleClass is a representation of the C record GtkIMContextSimpleClass.
type IMContextSimpleClass struct {
	native unsafe.Pointer
}

func (recv *IMContextSimpleClass) ToC() unsafe.Pointer {
	return recv.native
}

// IMContextSimplePrivate is a representation of the C record GtkIMContextSimplePrivate.
type IMContextSimplePrivate struct {
	native unsafe.Pointer
}

func (recv *IMContextSimplePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// IMMulticontextClass is a representation of the C record GtkIMMulticontextClass.
type IMMulticontextClass struct {
	native unsafe.Pointer
}

func (recv *IMMulticontextClass) ToC() unsafe.Pointer {
	return recv.native
}

// IMMulticontextPrivate is a representation of the C record GtkIMMulticontextPrivate.
type IMMulticontextPrivate struct {
	native unsafe.Pointer
}

func (recv *IMMulticontextPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// IconFactoryClass is a representation of the C record GtkIconFactoryClass.
type IconFactoryClass struct {
	native unsafe.Pointer
}

func (recv *IconFactoryClass) ToC() unsafe.Pointer {
	return recv.native
}

// IconFactoryPrivate is a representation of the C record GtkIconFactoryPrivate.
type IconFactoryPrivate struct {
	native unsafe.Pointer
}

func (recv *IconFactoryPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// IconInfoClass is a representation of the C record GtkIconInfoClass.
type IconInfoClass struct {
	native unsafe.Pointer
}

func (recv *IconInfoClass) ToC() unsafe.Pointer {
	return recv.native
}

// IconSet is a representation of the C record GtkIconSet.
type IconSet struct {
	native unsafe.Pointer
}

func (recv *IconSet) ToC() unsafe.Pointer {
	return recv.native
}

// IconSource is a representation of the C record GtkIconSource.
type IconSource struct {
	native unsafe.Pointer
}

func (recv *IconSource) ToC() unsafe.Pointer {
	return recv.native
}

// IconThemeClass is a representation of the C record GtkIconThemeClass.
type IconThemeClass struct {
	native unsafe.Pointer
}

func (recv *IconThemeClass) ToC() unsafe.Pointer {
	return recv.native
}

// IconThemePrivate is a representation of the C record GtkIconThemePrivate.
type IconThemePrivate struct {
	native unsafe.Pointer
}

func (recv *IconThemePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// IconViewAccessibleClass is a representation of the C record GtkIconViewAccessibleClass.
type IconViewAccessibleClass struct {
	native unsafe.Pointer
}

func (recv *IconViewAccessibleClass) ToC() unsafe.Pointer {
	return recv.native
}

// IconViewAccessiblePrivate is a representation of the C record GtkIconViewAccessiblePrivate.
type IconViewAccessiblePrivate struct {
	native unsafe.Pointer
}

func (recv *IconViewAccessiblePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// IconViewClass is a representation of the C record GtkIconViewClass.
type IconViewClass struct {
	native unsafe.Pointer
}

func (recv *IconViewClass) ToC() unsafe.Pointer {
	return recv.native
}

// IconViewPrivate is a representation of the C record GtkIconViewPrivate.
type IconViewPrivate struct {
	native unsafe.Pointer
}

func (recv *IconViewPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ImageAccessibleClass is a representation of the C record GtkImageAccessibleClass.
type ImageAccessibleClass struct {
	native unsafe.Pointer
}

func (recv *ImageAccessibleClass) ToC() unsafe.Pointer {
	return recv.native
}

// ImageAccessiblePrivate is a representation of the C record GtkImageAccessiblePrivate.
type ImageAccessiblePrivate struct {
	native unsafe.Pointer
}

func (recv *ImageAccessiblePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ImageCellAccessibleClass is a representation of the C record GtkImageCellAccessibleClass.
type ImageCellAccessibleClass struct {
	native unsafe.Pointer
}

func (recv *ImageCellAccessibleClass) ToC() unsafe.Pointer {
	return recv.native
}

// ImageCellAccessiblePrivate is a representation of the C record GtkImageCellAccessiblePrivate.
type ImageCellAccessiblePrivate struct {
	native unsafe.Pointer
}

func (recv *ImageCellAccessiblePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ImageClass is a representation of the C record GtkImageClass.
type ImageClass struct {
	native unsafe.Pointer
}

func (recv *ImageClass) ToC() unsafe.Pointer {
	return recv.native
}

// ImageMenuItemClass is a representation of the C record GtkImageMenuItemClass.
type ImageMenuItemClass struct {
	native unsafe.Pointer
}

func (recv *ImageMenuItemClass) ToC() unsafe.Pointer {
	return recv.native
}

// ImageMenuItemPrivate is a representation of the C record GtkImageMenuItemPrivate.
type ImageMenuItemPrivate struct {
	native unsafe.Pointer
}

func (recv *ImageMenuItemPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ImagePrivate is a representation of the C record GtkImagePrivate.
type ImagePrivate struct {
	native unsafe.Pointer
}

func (recv *ImagePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// InfoBarClass is a representation of the C record GtkInfoBarClass.
type InfoBarClass struct {
	native unsafe.Pointer
}

func (recv *InfoBarClass) ToC() unsafe.Pointer {
	return recv.native
}

// InfoBarPrivate is a representation of the C record GtkInfoBarPrivate.
type InfoBarPrivate struct {
	native unsafe.Pointer
}

func (recv *InfoBarPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// InvisibleClass is a representation of the C record GtkInvisibleClass.
type InvisibleClass struct {
	native unsafe.Pointer
}

func (recv *InvisibleClass) ToC() unsafe.Pointer {
	return recv.native
}

// InvisiblePrivate is a representation of the C record GtkInvisiblePrivate.
type InvisiblePrivate struct {
	native unsafe.Pointer
}

func (recv *InvisiblePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// LabelAccessibleClass is a representation of the C record GtkLabelAccessibleClass.
type LabelAccessibleClass struct {
	native unsafe.Pointer
}

func (recv *LabelAccessibleClass) ToC() unsafe.Pointer {
	return recv.native
}

// LabelAccessiblePrivate is a representation of the C record GtkLabelAccessiblePrivate.
type LabelAccessiblePrivate struct {
	native unsafe.Pointer
}

func (recv *LabelAccessiblePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// LabelClass is a representation of the C record GtkLabelClass.
type LabelClass struct {
	native unsafe.Pointer
}

func (recv *LabelClass) ToC() unsafe.Pointer {
	return recv.native
}

// LabelPrivate is a representation of the C record GtkLabelPrivate.
type LabelPrivate struct {
	native unsafe.Pointer
}

func (recv *LabelPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// LabelSelectionInfo is a representation of the C record GtkLabelSelectionInfo.
type LabelSelectionInfo struct {
	native unsafe.Pointer
}

func (recv *LabelSelectionInfo) ToC() unsafe.Pointer {
	return recv.native
}

// LayoutClass is a representation of the C record GtkLayoutClass.
type LayoutClass struct {
	native unsafe.Pointer
}

func (recv *LayoutClass) ToC() unsafe.Pointer {
	return recv.native
}

// LayoutPrivate is a representation of the C record GtkLayoutPrivate.
type LayoutPrivate struct {
	native unsafe.Pointer
}

func (recv *LayoutPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// LevelBarAccessibleClass is a representation of the C record GtkLevelBarAccessibleClass.
type LevelBarAccessibleClass struct {
	native unsafe.Pointer
}

func (recv *LevelBarAccessibleClass) ToC() unsafe.Pointer {
	return recv.native
}

// LevelBarAccessiblePrivate is a representation of the C record GtkLevelBarAccessiblePrivate.
type LevelBarAccessiblePrivate struct {
	native unsafe.Pointer
}

func (recv *LevelBarAccessiblePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// LevelBarClass is a representation of the C record GtkLevelBarClass.
type LevelBarClass struct {
	native unsafe.Pointer
}

func (recv *LevelBarClass) ToC() unsafe.Pointer {
	return recv.native
}

// LevelBarPrivate is a representation of the C record GtkLevelBarPrivate.
type LevelBarPrivate struct {
	native unsafe.Pointer
}

func (recv *LevelBarPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// LinkButtonAccessibleClass is a representation of the C record GtkLinkButtonAccessibleClass.
type LinkButtonAccessibleClass struct {
	native unsafe.Pointer
}

func (recv *LinkButtonAccessibleClass) ToC() unsafe.Pointer {
	return recv.native
}

// LinkButtonAccessiblePrivate is a representation of the C record GtkLinkButtonAccessiblePrivate.
type LinkButtonAccessiblePrivate struct {
	native unsafe.Pointer
}

func (recv *LinkButtonAccessiblePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// LinkButtonClass is a representation of the C record GtkLinkButtonClass.
type LinkButtonClass struct {
	native unsafe.Pointer
}

func (recv *LinkButtonClass) ToC() unsafe.Pointer {
	return recv.native
}

// LinkButtonPrivate is a representation of the C record GtkLinkButtonPrivate.
type LinkButtonPrivate struct {
	native unsafe.Pointer
}

func (recv *LinkButtonPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ListBoxAccessibleClass is a representation of the C record GtkListBoxAccessibleClass.
type ListBoxAccessibleClass struct {
	native unsafe.Pointer
}

func (recv *ListBoxAccessibleClass) ToC() unsafe.Pointer {
	return recv.native
}

// ListBoxAccessiblePrivate is a representation of the C record GtkListBoxAccessiblePrivate.
type ListBoxAccessiblePrivate struct {
	native unsafe.Pointer
}

func (recv *ListBoxAccessiblePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ListBoxClass is a representation of the C record GtkListBoxClass.
type ListBoxClass struct {
	native unsafe.Pointer
}

func (recv *ListBoxClass) ToC() unsafe.Pointer {
	return recv.native
}

// ListBoxRowAccessibleClass is a representation of the C record GtkListBoxRowAccessibleClass.
type ListBoxRowAccessibleClass struct {
	native unsafe.Pointer
}

func (recv *ListBoxRowAccessibleClass) ToC() unsafe.Pointer {
	return recv.native
}

// ListBoxRowClass is a representation of the C record GtkListBoxRowClass.
type ListBoxRowClass struct {
	native unsafe.Pointer
}

func (recv *ListBoxRowClass) ToC() unsafe.Pointer {
	return recv.native
}

// ListStoreClass is a representation of the C record GtkListStoreClass.
type ListStoreClass struct {
	native unsafe.Pointer
}

func (recv *ListStoreClass) ToC() unsafe.Pointer {
	return recv.native
}

// ListStorePrivate is a representation of the C record GtkListStorePrivate.
type ListStorePrivate struct {
	native unsafe.Pointer
}

func (recv *ListStorePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// LockButtonAccessibleClass is a representation of the C record GtkLockButtonAccessibleClass.
type LockButtonAccessibleClass struct {
	native unsafe.Pointer
}

func (recv *LockButtonAccessibleClass) ToC() unsafe.Pointer {
	return recv.native
}

// LockButtonAccessiblePrivate is a representation of the C record GtkLockButtonAccessiblePrivate.
type LockButtonAccessiblePrivate struct {
	native unsafe.Pointer
}

func (recv *LockButtonAccessiblePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// LockButtonClass is a representation of the C record GtkLockButtonClass.
type LockButtonClass struct {
	native unsafe.Pointer
}

func (recv *LockButtonClass) ToC() unsafe.Pointer {
	return recv.native
}

// LockButtonPrivate is a representation of the C record GtkLockButtonPrivate.
type LockButtonPrivate struct {
	native unsafe.Pointer
}

func (recv *LockButtonPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// MenuAccessibleClass is a representation of the C record GtkMenuAccessibleClass.
type MenuAccessibleClass struct {
	native unsafe.Pointer
}

func (recv *MenuAccessibleClass) ToC() unsafe.Pointer {
	return recv.native
}

// MenuAccessiblePrivate is a representation of the C record GtkMenuAccessiblePrivate.
type MenuAccessiblePrivate struct {
	native unsafe.Pointer
}

func (recv *MenuAccessiblePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// MenuBarClass is a representation of the C record GtkMenuBarClass.
type MenuBarClass struct {
	native unsafe.Pointer
}

func (recv *MenuBarClass) ToC() unsafe.Pointer {
	return recv.native
}

// MenuBarPrivate is a representation of the C record GtkMenuBarPrivate.
type MenuBarPrivate struct {
	native unsafe.Pointer
}

func (recv *MenuBarPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// MenuButtonAccessibleClass is a representation of the C record GtkMenuButtonAccessibleClass.
type MenuButtonAccessibleClass struct {
	native unsafe.Pointer
}

func (recv *MenuButtonAccessibleClass) ToC() unsafe.Pointer {
	return recv.native
}

// MenuButtonAccessiblePrivate is a representation of the C record GtkMenuButtonAccessiblePrivate.
type MenuButtonAccessiblePrivate struct {
	native unsafe.Pointer
}

func (recv *MenuButtonAccessiblePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// MenuButtonClass is a representation of the C record GtkMenuButtonClass.
type MenuButtonClass struct {
	native unsafe.Pointer
}

func (recv *MenuButtonClass) ToC() unsafe.Pointer {
	return recv.native
}

// MenuButtonPrivate is a representation of the C record GtkMenuButtonPrivate.
type MenuButtonPrivate struct {
	native unsafe.Pointer
}

func (recv *MenuButtonPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// MenuClass is a representation of the C record GtkMenuClass.
type MenuClass struct {
	native unsafe.Pointer
}

func (recv *MenuClass) ToC() unsafe.Pointer {
	return recv.native
}

// MenuItemAccessibleClass is a representation of the C record GtkMenuItemAccessibleClass.
type MenuItemAccessibleClass struct {
	native unsafe.Pointer
}

func (recv *MenuItemAccessibleClass) ToC() unsafe.Pointer {
	return recv.native
}

// MenuItemAccessiblePrivate is a representation of the C record GtkMenuItemAccessiblePrivate.
type MenuItemAccessiblePrivate struct {
	native unsafe.Pointer
}

func (recv *MenuItemAccessiblePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// MenuItemClass is a representation of the C record GtkMenuItemClass.
type MenuItemClass struct {
	native unsafe.Pointer
}

func (recv *MenuItemClass) ToC() unsafe.Pointer {
	return recv.native
}

// MenuItemPrivate is a representation of the C record GtkMenuItemPrivate.
type MenuItemPrivate struct {
	native unsafe.Pointer
}

func (recv *MenuItemPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// MenuPrivate is a representation of the C record GtkMenuPrivate.
type MenuPrivate struct {
	native unsafe.Pointer
}

func (recv *MenuPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// MenuShellAccessibleClass is a representation of the C record GtkMenuShellAccessibleClass.
type MenuShellAccessibleClass struct {
	native unsafe.Pointer
}

func (recv *MenuShellAccessibleClass) ToC() unsafe.Pointer {
	return recv.native
}

// MenuShellAccessiblePrivate is a representation of the C record GtkMenuShellAccessiblePrivate.
type MenuShellAccessiblePrivate struct {
	native unsafe.Pointer
}

func (recv *MenuShellAccessiblePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// MenuShellClass is a representation of the C record GtkMenuShellClass.
type MenuShellClass struct {
	native unsafe.Pointer
}

func (recv *MenuShellClass) ToC() unsafe.Pointer {
	return recv.native
}

// MenuShellPrivate is a representation of the C record GtkMenuShellPrivate.
type MenuShellPrivate struct {
	native unsafe.Pointer
}

func (recv *MenuShellPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// MenuToolButtonClass is a representation of the C record GtkMenuToolButtonClass.
type MenuToolButtonClass struct {
	native unsafe.Pointer
}

func (recv *MenuToolButtonClass) ToC() unsafe.Pointer {
	return recv.native
}

// MenuToolButtonPrivate is a representation of the C record GtkMenuToolButtonPrivate.
type MenuToolButtonPrivate struct {
	native unsafe.Pointer
}

func (recv *MenuToolButtonPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// MessageDialogClass is a representation of the C record GtkMessageDialogClass.
type MessageDialogClass struct {
	native unsafe.Pointer
}

func (recv *MessageDialogClass) ToC() unsafe.Pointer {
	return recv.native
}

// MessageDialogPrivate is a representation of the C record GtkMessageDialogPrivate.
type MessageDialogPrivate struct {
	native unsafe.Pointer
}

func (recv *MessageDialogPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// MiscClass is a representation of the C record GtkMiscClass.
type MiscClass struct {
	native unsafe.Pointer
}

func (recv *MiscClass) ToC() unsafe.Pointer {
	return recv.native
}

// MiscPrivate is a representation of the C record GtkMiscPrivate.
type MiscPrivate struct {
	native unsafe.Pointer
}

func (recv *MiscPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// MountOperationClass is a representation of the C record GtkMountOperationClass.
type MountOperationClass struct {
	native unsafe.Pointer
}

func (recv *MountOperationClass) ToC() unsafe.Pointer {
	return recv.native
}

// MountOperationPrivate is a representation of the C record GtkMountOperationPrivate.
type MountOperationPrivate struct {
	native unsafe.Pointer
}

func (recv *MountOperationPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// NativeDialogClass is a representation of the C record GtkNativeDialogClass.
type NativeDialogClass struct {
	native unsafe.Pointer
}

func (recv *NativeDialogClass) ToC() unsafe.Pointer {
	return recv.native
}

// NotebookAccessibleClass is a representation of the C record GtkNotebookAccessibleClass.
type NotebookAccessibleClass struct {
	native unsafe.Pointer
}

func (recv *NotebookAccessibleClass) ToC() unsafe.Pointer {
	return recv.native
}

// NotebookAccessiblePrivate is a representation of the C record GtkNotebookAccessiblePrivate.
type NotebookAccessiblePrivate struct {
	native unsafe.Pointer
}

func (recv *NotebookAccessiblePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// NotebookClass is a representation of the C record GtkNotebookClass.
type NotebookClass struct {
	native unsafe.Pointer
}

func (recv *NotebookClass) ToC() unsafe.Pointer {
	return recv.native
}

// NotebookPageAccessibleClass is a representation of the C record GtkNotebookPageAccessibleClass.
type NotebookPageAccessibleClass struct {
	native unsafe.Pointer
}

func (recv *NotebookPageAccessibleClass) ToC() unsafe.Pointer {
	return recv.native
}

// NotebookPageAccessiblePrivate is a representation of the C record GtkNotebookPageAccessiblePrivate.
type NotebookPageAccessiblePrivate struct {
	native unsafe.Pointer
}

func (recv *NotebookPageAccessiblePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// NotebookPrivate is a representation of the C record GtkNotebookPrivate.
type NotebookPrivate struct {
	native unsafe.Pointer
}

func (recv *NotebookPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// NumerableIconClass is a representation of the C record GtkNumerableIconClass.
type NumerableIconClass struct {
	native unsafe.Pointer
}

func (recv *NumerableIconClass) ToC() unsafe.Pointer {
	return recv.native
}

// NumerableIconPrivate is a representation of the C record GtkNumerableIconPrivate.
type NumerableIconPrivate struct {
	native unsafe.Pointer
}

func (recv *NumerableIconPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// OffscreenWindowClass is a representation of the C record GtkOffscreenWindowClass.
type OffscreenWindowClass struct {
	native unsafe.Pointer
}

func (recv *OffscreenWindowClass) ToC() unsafe.Pointer {
	return recv.native
}

// OrientableIface is a representation of the C record GtkOrientableIface.
type OrientableIface struct {
	native unsafe.Pointer
}

func (recv *OrientableIface) ToC() unsafe.Pointer {
	return recv.native
}

// OverlayClass is a representation of the C record GtkOverlayClass.
type OverlayClass struct {
	native unsafe.Pointer
}

func (recv *OverlayClass) ToC() unsafe.Pointer {
	return recv.native
}

// OverlayPrivate is a representation of the C record GtkOverlayPrivate.
type OverlayPrivate struct {
	native unsafe.Pointer
}

func (recv *OverlayPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// PadActionEntry is a representation of the C record GtkPadActionEntry.
type PadActionEntry struct {
	native unsafe.Pointer
}

func (recv *PadActionEntry) ToC() unsafe.Pointer {
	return recv.native
}

// PadControllerClass is a representation of the C record GtkPadControllerClass.
type PadControllerClass struct {
	native unsafe.Pointer
}

func (recv *PadControllerClass) ToC() unsafe.Pointer {
	return recv.native
}

// PageRange is a representation of the C record GtkPageRange.
type PageRange struct {
	native unsafe.Pointer
}

func (recv *PageRange) ToC() unsafe.Pointer {
	return recv.native
}

// PanedAccessibleClass is a representation of the C record GtkPanedAccessibleClass.
type PanedAccessibleClass struct {
	native unsafe.Pointer
}

func (recv *PanedAccessibleClass) ToC() unsafe.Pointer {
	return recv.native
}

// PanedAccessiblePrivate is a representation of the C record GtkPanedAccessiblePrivate.
type PanedAccessiblePrivate struct {
	native unsafe.Pointer
}

func (recv *PanedAccessiblePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// PanedClass is a representation of the C record GtkPanedClass.
type PanedClass struct {
	native unsafe.Pointer
}

func (recv *PanedClass) ToC() unsafe.Pointer {
	return recv.native
}

// PanedPrivate is a representation of the C record GtkPanedPrivate.
type PanedPrivate struct {
	native unsafe.Pointer
}

func (recv *PanedPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// PaperSize is a representation of the C record GtkPaperSize.
type PaperSize struct {
	native unsafe.Pointer
}

func (recv *PaperSize) ToC() unsafe.Pointer {
	return recv.native
}

// PlacesSidebarClass is a representation of the C record GtkPlacesSidebarClass.
type PlacesSidebarClass struct {
	native unsafe.Pointer
}

func (recv *PlacesSidebarClass) ToC() unsafe.Pointer {
	return recv.native
}

// PlugClass is a representation of the C record GtkPlugClass.
type PlugClass struct {
	native unsafe.Pointer
}

func (recv *PlugClass) ToC() unsafe.Pointer {
	return recv.native
}

// PlugPrivate is a representation of the C record GtkPlugPrivate.
type PlugPrivate struct {
	native unsafe.Pointer
}

func (recv *PlugPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// PopoverAccessibleClass is a representation of the C record GtkPopoverAccessibleClass.
type PopoverAccessibleClass struct {
	native unsafe.Pointer
}

func (recv *PopoverAccessibleClass) ToC() unsafe.Pointer {
	return recv.native
}

// PopoverClass is a representation of the C record GtkPopoverClass.
type PopoverClass struct {
	native unsafe.Pointer
}

func (recv *PopoverClass) ToC() unsafe.Pointer {
	return recv.native
}

// PopoverMenuClass is a representation of the C record GtkPopoverMenuClass.
type PopoverMenuClass struct {
	native unsafe.Pointer
}

func (recv *PopoverMenuClass) ToC() unsafe.Pointer {
	return recv.native
}

// PopoverPrivate is a representation of the C record GtkPopoverPrivate.
type PopoverPrivate struct {
	native unsafe.Pointer
}

func (recv *PopoverPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// PrintOperationClass is a representation of the C record GtkPrintOperationClass.
type PrintOperationClass struct {
	native unsafe.Pointer
}

func (recv *PrintOperationClass) ToC() unsafe.Pointer {
	return recv.native
}

// PrintOperationPreviewIface is a representation of the C record GtkPrintOperationPreviewIface.
type PrintOperationPreviewIface struct {
	native unsafe.Pointer
}

func (recv *PrintOperationPreviewIface) ToC() unsafe.Pointer {
	return recv.native
}

// PrintOperationPrivate is a representation of the C record GtkPrintOperationPrivate.
type PrintOperationPrivate struct {
	native unsafe.Pointer
}

func (recv *PrintOperationPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ProgressBarAccessibleClass is a representation of the C record GtkProgressBarAccessibleClass.
type ProgressBarAccessibleClass struct {
	native unsafe.Pointer
}

func (recv *ProgressBarAccessibleClass) ToC() unsafe.Pointer {
	return recv.native
}

// ProgressBarAccessiblePrivate is a representation of the C record GtkProgressBarAccessiblePrivate.
type ProgressBarAccessiblePrivate struct {
	native unsafe.Pointer
}

func (recv *ProgressBarAccessiblePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ProgressBarClass is a representation of the C record GtkProgressBarClass.
type ProgressBarClass struct {
	native unsafe.Pointer
}

func (recv *ProgressBarClass) ToC() unsafe.Pointer {
	return recv.native
}

// ProgressBarPrivate is a representation of the C record GtkProgressBarPrivate.
type ProgressBarPrivate struct {
	native unsafe.Pointer
}

func (recv *ProgressBarPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// RadioActionClass is a representation of the C record GtkRadioActionClass.
type RadioActionClass struct {
	native unsafe.Pointer
}

func (recv *RadioActionClass) ToC() unsafe.Pointer {
	return recv.native
}

// RadioActionEntry is a representation of the C record GtkRadioActionEntry.
type RadioActionEntry struct {
	native unsafe.Pointer
}

func (recv *RadioActionEntry) ToC() unsafe.Pointer {
	return recv.native
}

// RadioActionPrivate is a representation of the C record GtkRadioActionPrivate.
type RadioActionPrivate struct {
	native unsafe.Pointer
}

func (recv *RadioActionPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// RadioButtonAccessibleClass is a representation of the C record GtkRadioButtonAccessibleClass.
type RadioButtonAccessibleClass struct {
	native unsafe.Pointer
}

func (recv *RadioButtonAccessibleClass) ToC() unsafe.Pointer {
	return recv.native
}

// RadioButtonAccessiblePrivate is a representation of the C record GtkRadioButtonAccessiblePrivate.
type RadioButtonAccessiblePrivate struct {
	native unsafe.Pointer
}

func (recv *RadioButtonAccessiblePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// RadioButtonClass is a representation of the C record GtkRadioButtonClass.
type RadioButtonClass struct {
	native unsafe.Pointer
}

func (recv *RadioButtonClass) ToC() unsafe.Pointer {
	return recv.native
}

// RadioButtonPrivate is a representation of the C record GtkRadioButtonPrivate.
type RadioButtonPrivate struct {
	native unsafe.Pointer
}

func (recv *RadioButtonPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// RadioMenuItemAccessibleClass is a representation of the C record GtkRadioMenuItemAccessibleClass.
type RadioMenuItemAccessibleClass struct {
	native unsafe.Pointer
}

func (recv *RadioMenuItemAccessibleClass) ToC() unsafe.Pointer {
	return recv.native
}

// RadioMenuItemAccessiblePrivate is a representation of the C record GtkRadioMenuItemAccessiblePrivate.
type RadioMenuItemAccessiblePrivate struct {
	native unsafe.Pointer
}

func (recv *RadioMenuItemAccessiblePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// RadioMenuItemClass is a representation of the C record GtkRadioMenuItemClass.
type RadioMenuItemClass struct {
	native unsafe.Pointer
}

func (recv *RadioMenuItemClass) ToC() unsafe.Pointer {
	return recv.native
}

// RadioMenuItemPrivate is a representation of the C record GtkRadioMenuItemPrivate.
type RadioMenuItemPrivate struct {
	native unsafe.Pointer
}

func (recv *RadioMenuItemPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// RadioToolButtonClass is a representation of the C record GtkRadioToolButtonClass.
type RadioToolButtonClass struct {
	native unsafe.Pointer
}

func (recv *RadioToolButtonClass) ToC() unsafe.Pointer {
	return recv.native
}

// RangeAccessibleClass is a representation of the C record GtkRangeAccessibleClass.
type RangeAccessibleClass struct {
	native unsafe.Pointer
}

func (recv *RangeAccessibleClass) ToC() unsafe.Pointer {
	return recv.native
}

// RangeAccessiblePrivate is a representation of the C record GtkRangeAccessiblePrivate.
type RangeAccessiblePrivate struct {
	native unsafe.Pointer
}

func (recv *RangeAccessiblePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// RangeClass is a representation of the C record GtkRangeClass.
type RangeClass struct {
	native unsafe.Pointer
}

func (recv *RangeClass) ToC() unsafe.Pointer {
	return recv.native
}

// RangePrivate is a representation of the C record GtkRangePrivate.
type RangePrivate struct {
	native unsafe.Pointer
}

func (recv *RangePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// RcContext is a representation of the C record GtkRcContext.
type RcContext struct {
	native unsafe.Pointer
}

func (recv *RcContext) ToC() unsafe.Pointer {
	return recv.native
}

// RcProperty is a representation of the C record GtkRcProperty.
type RcProperty struct {
	native unsafe.Pointer
}

func (recv *RcProperty) ToC() unsafe.Pointer {
	return recv.native
}

// RcStyleClass is a representation of the C record GtkRcStyleClass.
type RcStyleClass struct {
	native unsafe.Pointer
}

func (recv *RcStyleClass) ToC() unsafe.Pointer {
	return recv.native
}

// RecentActionClass is a representation of the C record GtkRecentActionClass.
type RecentActionClass struct {
	native unsafe.Pointer
}

func (recv *RecentActionClass) ToC() unsafe.Pointer {
	return recv.native
}

// RecentActionPrivate is a representation of the C record GtkRecentActionPrivate.
type RecentActionPrivate struct {
	native unsafe.Pointer
}

func (recv *RecentActionPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// RecentChooserDialogClass is a representation of the C record GtkRecentChooserDialogClass.
type RecentChooserDialogClass struct {
	native unsafe.Pointer
}

func (recv *RecentChooserDialogClass) ToC() unsafe.Pointer {
	return recv.native
}

// RecentChooserDialogPrivate is a representation of the C record GtkRecentChooserDialogPrivate.
type RecentChooserDialogPrivate struct {
	native unsafe.Pointer
}

func (recv *RecentChooserDialogPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// RecentChooserIface is a representation of the C record GtkRecentChooserIface.
type RecentChooserIface struct {
	native unsafe.Pointer
}

func (recv *RecentChooserIface) ToC() unsafe.Pointer {
	return recv.native
}

// RecentChooserMenuClass is a representation of the C record GtkRecentChooserMenuClass.
type RecentChooserMenuClass struct {
	native unsafe.Pointer
}

func (recv *RecentChooserMenuClass) ToC() unsafe.Pointer {
	return recv.native
}

// RecentChooserMenuPrivate is a representation of the C record GtkRecentChooserMenuPrivate.
type RecentChooserMenuPrivate struct {
	native unsafe.Pointer
}

func (recv *RecentChooserMenuPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// RecentChooserWidgetClass is a representation of the C record GtkRecentChooserWidgetClass.
type RecentChooserWidgetClass struct {
	native unsafe.Pointer
}

func (recv *RecentChooserWidgetClass) ToC() unsafe.Pointer {
	return recv.native
}

// RecentChooserWidgetPrivate is a representation of the C record GtkRecentChooserWidgetPrivate.
type RecentChooserWidgetPrivate struct {
	native unsafe.Pointer
}

func (recv *RecentChooserWidgetPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// RecentData is a representation of the C record GtkRecentData.
type RecentData struct {
	native unsafe.Pointer
}

func (recv *RecentData) ToC() unsafe.Pointer {
	return recv.native
}

// RecentFilterInfo is a representation of the C record GtkRecentFilterInfo.
type RecentFilterInfo struct {
	native unsafe.Pointer
}

func (recv *RecentFilterInfo) ToC() unsafe.Pointer {
	return recv.native
}

// RecentManagerPrivate is a representation of the C record GtkRecentManagerPrivate.
type RecentManagerPrivate struct {
	native unsafe.Pointer
}

func (recv *RecentManagerPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// RendererCellAccessibleClass is a representation of the C record GtkRendererCellAccessibleClass.
type RendererCellAccessibleClass struct {
	native unsafe.Pointer
}

func (recv *RendererCellAccessibleClass) ToC() unsafe.Pointer {
	return recv.native
}

// RendererCellAccessiblePrivate is a representation of the C record GtkRendererCellAccessiblePrivate.
type RendererCellAccessiblePrivate struct {
	native unsafe.Pointer
}

func (recv *RendererCellAccessiblePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// RequestedSize is a representation of the C record GtkRequestedSize.
type RequestedSize struct {
	native unsafe.Pointer
}

func (recv *RequestedSize) ToC() unsafe.Pointer {
	return recv.native
}

// Requisition is a representation of the C record GtkRequisition.
type Requisition struct {
	native unsafe.Pointer
}

func (recv *Requisition) ToC() unsafe.Pointer {
	return recv.native
}

// RevealerClass is a representation of the C record GtkRevealerClass.
type RevealerClass struct {
	native unsafe.Pointer
}

func (recv *RevealerClass) ToC() unsafe.Pointer {
	return recv.native
}

// ScaleAccessibleClass is a representation of the C record GtkScaleAccessibleClass.
type ScaleAccessibleClass struct {
	native unsafe.Pointer
}

func (recv *ScaleAccessibleClass) ToC() unsafe.Pointer {
	return recv.native
}

// ScaleAccessiblePrivate is a representation of the C record GtkScaleAccessiblePrivate.
type ScaleAccessiblePrivate struct {
	native unsafe.Pointer
}

func (recv *ScaleAccessiblePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ScaleButtonAccessibleClass is a representation of the C record GtkScaleButtonAccessibleClass.
type ScaleButtonAccessibleClass struct {
	native unsafe.Pointer
}

func (recv *ScaleButtonAccessibleClass) ToC() unsafe.Pointer {
	return recv.native
}

// ScaleButtonAccessiblePrivate is a representation of the C record GtkScaleButtonAccessiblePrivate.
type ScaleButtonAccessiblePrivate struct {
	native unsafe.Pointer
}

func (recv *ScaleButtonAccessiblePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ScaleButtonClass is a representation of the C record GtkScaleButtonClass.
type ScaleButtonClass struct {
	native unsafe.Pointer
}

func (recv *ScaleButtonClass) ToC() unsafe.Pointer {
	return recv.native
}

// ScaleButtonPrivate is a representation of the C record GtkScaleButtonPrivate.
type ScaleButtonPrivate struct {
	native unsafe.Pointer
}

func (recv *ScaleButtonPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ScaleClass is a representation of the C record GtkScaleClass.
type ScaleClass struct {
	native unsafe.Pointer
}

func (recv *ScaleClass) ToC() unsafe.Pointer {
	return recv.native
}

// ScalePrivate is a representation of the C record GtkScalePrivate.
type ScalePrivate struct {
	native unsafe.Pointer
}

func (recv *ScalePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ScrollableInterface is a representation of the C record GtkScrollableInterface.
type ScrollableInterface struct {
	native unsafe.Pointer
}

func (recv *ScrollableInterface) ToC() unsafe.Pointer {
	return recv.native
}

// ScrollbarClass is a representation of the C record GtkScrollbarClass.
type ScrollbarClass struct {
	native unsafe.Pointer
}

func (recv *ScrollbarClass) ToC() unsafe.Pointer {
	return recv.native
}

// ScrolledWindowAccessibleClass is a representation of the C record GtkScrolledWindowAccessibleClass.
type ScrolledWindowAccessibleClass struct {
	native unsafe.Pointer
}

func (recv *ScrolledWindowAccessibleClass) ToC() unsafe.Pointer {
	return recv.native
}

// ScrolledWindowAccessiblePrivate is a representation of the C record GtkScrolledWindowAccessiblePrivate.
type ScrolledWindowAccessiblePrivate struct {
	native unsafe.Pointer
}

func (recv *ScrolledWindowAccessiblePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ScrolledWindowClass is a representation of the C record GtkScrolledWindowClass.
type ScrolledWindowClass struct {
	native unsafe.Pointer
}

func (recv *ScrolledWindowClass) ToC() unsafe.Pointer {
	return recv.native
}

// ScrolledWindowPrivate is a representation of the C record GtkScrolledWindowPrivate.
type ScrolledWindowPrivate struct {
	native unsafe.Pointer
}

func (recv *ScrolledWindowPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// SearchBarClass is a representation of the C record GtkSearchBarClass.
type SearchBarClass struct {
	native unsafe.Pointer
}

func (recv *SearchBarClass) ToC() unsafe.Pointer {
	return recv.native
}

// SearchEntryClass is a representation of the C record GtkSearchEntryClass.
type SearchEntryClass struct {
	native unsafe.Pointer
}

func (recv *SearchEntryClass) ToC() unsafe.Pointer {
	return recv.native
}

// SelectionData is a representation of the C record GtkSelectionData.
type SelectionData struct {
	native unsafe.Pointer
}

func (recv *SelectionData) ToC() unsafe.Pointer {
	return recv.native
}

// SeparatorClass is a representation of the C record GtkSeparatorClass.
type SeparatorClass struct {
	native unsafe.Pointer
}

func (recv *SeparatorClass) ToC() unsafe.Pointer {
	return recv.native
}

// SeparatorMenuItemClass is a representation of the C record GtkSeparatorMenuItemClass.
type SeparatorMenuItemClass struct {
	native unsafe.Pointer
}

func (recv *SeparatorMenuItemClass) ToC() unsafe.Pointer {
	return recv.native
}

// SeparatorPrivate is a representation of the C record GtkSeparatorPrivate.
type SeparatorPrivate struct {
	native unsafe.Pointer
}

func (recv *SeparatorPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// SeparatorToolItemClass is a representation of the C record GtkSeparatorToolItemClass.
type SeparatorToolItemClass struct {
	native unsafe.Pointer
}

func (recv *SeparatorToolItemClass) ToC() unsafe.Pointer {
	return recv.native
}

// SeparatorToolItemPrivate is a representation of the C record GtkSeparatorToolItemPrivate.
type SeparatorToolItemPrivate struct {
	native unsafe.Pointer
}

func (recv *SeparatorToolItemPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// SettingsClass is a representation of the C record GtkSettingsClass.
type SettingsClass struct {
	native unsafe.Pointer
}

func (recv *SettingsClass) ToC() unsafe.Pointer {
	return recv.native
}

// SettingsPrivate is a representation of the C record GtkSettingsPrivate.
type SettingsPrivate struct {
	native unsafe.Pointer
}

func (recv *SettingsPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// SettingsValue is a representation of the C record GtkSettingsValue.
type SettingsValue struct {
	native unsafe.Pointer
}

func (recv *SettingsValue) ToC() unsafe.Pointer {
	return recv.native
}

// ShortcutLabelClass is a representation of the C record GtkShortcutLabelClass.
type ShortcutLabelClass struct {
	native unsafe.Pointer
}

func (recv *ShortcutLabelClass) ToC() unsafe.Pointer {
	return recv.native
}

// ShortcutsGroupClass is a representation of the C record GtkShortcutsGroupClass.
type ShortcutsGroupClass struct {
	native unsafe.Pointer
}

func (recv *ShortcutsGroupClass) ToC() unsafe.Pointer {
	return recv.native
}

// ShortcutsSectionClass is a representation of the C record GtkShortcutsSectionClass.
type ShortcutsSectionClass struct {
	native unsafe.Pointer
}

func (recv *ShortcutsSectionClass) ToC() unsafe.Pointer {
	return recv.native
}

// ShortcutsShortcutClass is a representation of the C record GtkShortcutsShortcutClass.
type ShortcutsShortcutClass struct {
	native unsafe.Pointer
}

func (recv *ShortcutsShortcutClass) ToC() unsafe.Pointer {
	return recv.native
}

// ShortcutsWindowClass is a representation of the C record GtkShortcutsWindowClass.
type ShortcutsWindowClass struct {
	native unsafe.Pointer
}

func (recv *ShortcutsWindowClass) ToC() unsafe.Pointer {
	return recv.native
}

// SizeGroupClass is a representation of the C record GtkSizeGroupClass.
type SizeGroupClass struct {
	native unsafe.Pointer
}

func (recv *SizeGroupClass) ToC() unsafe.Pointer {
	return recv.native
}

// SizeGroupPrivate is a representation of the C record GtkSizeGroupPrivate.
type SizeGroupPrivate struct {
	native unsafe.Pointer
}

func (recv *SizeGroupPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// SocketClass is a representation of the C record GtkSocketClass.
type SocketClass struct {
	native unsafe.Pointer
}

func (recv *SocketClass) ToC() unsafe.Pointer {
	return recv.native
}

// SocketPrivate is a representation of the C record GtkSocketPrivate.
type SocketPrivate struct {
	native unsafe.Pointer
}

func (recv *SocketPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// SpinButtonAccessibleClass is a representation of the C record GtkSpinButtonAccessibleClass.
type SpinButtonAccessibleClass struct {
	native unsafe.Pointer
}

func (recv *SpinButtonAccessibleClass) ToC() unsafe.Pointer {
	return recv.native
}

// SpinButtonAccessiblePrivate is a representation of the C record GtkSpinButtonAccessiblePrivate.
type SpinButtonAccessiblePrivate struct {
	native unsafe.Pointer
}

func (recv *SpinButtonAccessiblePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// SpinButtonClass is a representation of the C record GtkSpinButtonClass.
type SpinButtonClass struct {
	native unsafe.Pointer
}

func (recv *SpinButtonClass) ToC() unsafe.Pointer {
	return recv.native
}

// SpinButtonPrivate is a representation of the C record GtkSpinButtonPrivate.
type SpinButtonPrivate struct {
	native unsafe.Pointer
}

func (recv *SpinButtonPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// SpinnerAccessibleClass is a representation of the C record GtkSpinnerAccessibleClass.
type SpinnerAccessibleClass struct {
	native unsafe.Pointer
}

func (recv *SpinnerAccessibleClass) ToC() unsafe.Pointer {
	return recv.native
}

// SpinnerAccessiblePrivate is a representation of the C record GtkSpinnerAccessiblePrivate.
type SpinnerAccessiblePrivate struct {
	native unsafe.Pointer
}

func (recv *SpinnerAccessiblePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// SpinnerClass is a representation of the C record GtkSpinnerClass.
type SpinnerClass struct {
	native unsafe.Pointer
}

func (recv *SpinnerClass) ToC() unsafe.Pointer {
	return recv.native
}

// SpinnerPrivate is a representation of the C record GtkSpinnerPrivate.
type SpinnerPrivate struct {
	native unsafe.Pointer
}

func (recv *SpinnerPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : StackAccessibleClass : blacklisted

// StackClass is a representation of the C record GtkStackClass.
type StackClass struct {
	native unsafe.Pointer
}

func (recv *StackClass) ToC() unsafe.Pointer {
	return recv.native
}

// StackSidebarClass is a representation of the C record GtkStackSidebarClass.
type StackSidebarClass struct {
	native unsafe.Pointer
}

func (recv *StackSidebarClass) ToC() unsafe.Pointer {
	return recv.native
}

// StackSidebarPrivate is a representation of the C record GtkStackSidebarPrivate.
type StackSidebarPrivate struct {
	native unsafe.Pointer
}

func (recv *StackSidebarPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// StackSwitcherClass is a representation of the C record GtkStackSwitcherClass.
type StackSwitcherClass struct {
	native unsafe.Pointer
}

func (recv *StackSwitcherClass) ToC() unsafe.Pointer {
	return recv.native
}

// StatusIconClass is a representation of the C record GtkStatusIconClass.
type StatusIconClass struct {
	native unsafe.Pointer
}

func (recv *StatusIconClass) ToC() unsafe.Pointer {
	return recv.native
}

// StatusIconPrivate is a representation of the C record GtkStatusIconPrivate.
type StatusIconPrivate struct {
	native unsafe.Pointer
}

func (recv *StatusIconPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// StatusbarAccessibleClass is a representation of the C record GtkStatusbarAccessibleClass.
type StatusbarAccessibleClass struct {
	native unsafe.Pointer
}

func (recv *StatusbarAccessibleClass) ToC() unsafe.Pointer {
	return recv.native
}

// StatusbarAccessiblePrivate is a representation of the C record GtkStatusbarAccessiblePrivate.
type StatusbarAccessiblePrivate struct {
	native unsafe.Pointer
}

func (recv *StatusbarAccessiblePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// StatusbarClass is a representation of the C record GtkStatusbarClass.
type StatusbarClass struct {
	native unsafe.Pointer
}

func (recv *StatusbarClass) ToC() unsafe.Pointer {
	return recv.native
}

// StatusbarPrivate is a representation of the C record GtkStatusbarPrivate.
type StatusbarPrivate struct {
	native unsafe.Pointer
}

func (recv *StatusbarPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// StockItem is a representation of the C record GtkStockItem.
type StockItem struct {
	native unsafe.Pointer
}

func (recv *StockItem) ToC() unsafe.Pointer {
	return recv.native
}

// StyleClass is a representation of the C record GtkStyleClass.
type StyleClass struct {
	native unsafe.Pointer
}

func (recv *StyleClass) ToC() unsafe.Pointer {
	return recv.native
}

// StyleContextClass is a representation of the C record GtkStyleContextClass.
type StyleContextClass struct {
	native unsafe.Pointer
}

func (recv *StyleContextClass) ToC() unsafe.Pointer {
	return recv.native
}

// StyleContextPrivate is a representation of the C record GtkStyleContextPrivate.
type StyleContextPrivate struct {
	native unsafe.Pointer
}

func (recv *StyleContextPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// StylePropertiesClass is a representation of the C record GtkStylePropertiesClass.
type StylePropertiesClass struct {
	native unsafe.Pointer
}

func (recv *StylePropertiesClass) ToC() unsafe.Pointer {
	return recv.native
}

// StylePropertiesPrivate is a representation of the C record GtkStylePropertiesPrivate.
type StylePropertiesPrivate struct {
	native unsafe.Pointer
}

func (recv *StylePropertiesPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// StyleProviderIface is a representation of the C record GtkStyleProviderIface.
type StyleProviderIface struct {
	native unsafe.Pointer
}

func (recv *StyleProviderIface) ToC() unsafe.Pointer {
	return recv.native
}

// SwitchAccessibleClass is a representation of the C record GtkSwitchAccessibleClass.
type SwitchAccessibleClass struct {
	native unsafe.Pointer
}

func (recv *SwitchAccessibleClass) ToC() unsafe.Pointer {
	return recv.native
}

// SwitchAccessiblePrivate is a representation of the C record GtkSwitchAccessiblePrivate.
type SwitchAccessiblePrivate struct {
	native unsafe.Pointer
}

func (recv *SwitchAccessiblePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// SwitchClass is a representation of the C record GtkSwitchClass.
type SwitchClass struct {
	native unsafe.Pointer
}

func (recv *SwitchClass) ToC() unsafe.Pointer {
	return recv.native
}

// SwitchPrivate is a representation of the C record GtkSwitchPrivate.
type SwitchPrivate struct {
	native unsafe.Pointer
}

func (recv *SwitchPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// SymbolicColor is a representation of the C record GtkSymbolicColor.
type SymbolicColor struct {
	native unsafe.Pointer
}

func (recv *SymbolicColor) ToC() unsafe.Pointer {
	return recv.native
}

// TableChild is a representation of the C record GtkTableChild.
type TableChild struct {
	native unsafe.Pointer
}

func (recv *TableChild) ToC() unsafe.Pointer {
	return recv.native
}

// TableClass is a representation of the C record GtkTableClass.
type TableClass struct {
	native unsafe.Pointer
}

func (recv *TableClass) ToC() unsafe.Pointer {
	return recv.native
}

// TablePrivate is a representation of the C record GtkTablePrivate.
type TablePrivate struct {
	native unsafe.Pointer
}

func (recv *TablePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// TableRowCol is a representation of the C record GtkTableRowCol.
type TableRowCol struct {
	native unsafe.Pointer
}

func (recv *TableRowCol) ToC() unsafe.Pointer {
	return recv.native
}

// TargetEntry is a representation of the C record GtkTargetEntry.
type TargetEntry struct {
	native unsafe.Pointer
}

func (recv *TargetEntry) ToC() unsafe.Pointer {
	return recv.native
}

// TargetList is a representation of the C record GtkTargetList.
type TargetList struct {
	native unsafe.Pointer
}

func (recv *TargetList) ToC() unsafe.Pointer {
	return recv.native
}

// TargetPair is a representation of the C record GtkTargetPair.
type TargetPair struct {
	native unsafe.Pointer
}

func (recv *TargetPair) ToC() unsafe.Pointer {
	return recv.native
}

// TearoffMenuItemClass is a representation of the C record GtkTearoffMenuItemClass.
type TearoffMenuItemClass struct {
	native unsafe.Pointer
}

func (recv *TearoffMenuItemClass) ToC() unsafe.Pointer {
	return recv.native
}

// TearoffMenuItemPrivate is a representation of the C record GtkTearoffMenuItemPrivate.
type TearoffMenuItemPrivate struct {
	native unsafe.Pointer
}

func (recv *TearoffMenuItemPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// TextAppearance is a representation of the C record GtkTextAppearance.
type TextAppearance struct {
	native unsafe.Pointer
}

func (recv *TextAppearance) ToC() unsafe.Pointer {
	return recv.native
}

// TextAttributes is a representation of the C record GtkTextAttributes.
type TextAttributes struct {
	native unsafe.Pointer
}

func (recv *TextAttributes) ToC() unsafe.Pointer {
	return recv.native
}

// TextBTree is a representation of the C record GtkTextBTree.
type TextBTree struct {
	native unsafe.Pointer
}

func (recv *TextBTree) ToC() unsafe.Pointer {
	return recv.native
}

// TextBufferClass is a representation of the C record GtkTextBufferClass.
type TextBufferClass struct {
	native unsafe.Pointer
}

func (recv *TextBufferClass) ToC() unsafe.Pointer {
	return recv.native
}

// TextBufferPrivate is a representation of the C record GtkTextBufferPrivate.
type TextBufferPrivate struct {
	native unsafe.Pointer
}

func (recv *TextBufferPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// TextCellAccessibleClass is a representation of the C record GtkTextCellAccessibleClass.
type TextCellAccessibleClass struct {
	native unsafe.Pointer
}

func (recv *TextCellAccessibleClass) ToC() unsafe.Pointer {
	return recv.native
}

// TextCellAccessiblePrivate is a representation of the C record GtkTextCellAccessiblePrivate.
type TextCellAccessiblePrivate struct {
	native unsafe.Pointer
}

func (recv *TextCellAccessiblePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// TextChildAnchorClass is a representation of the C record GtkTextChildAnchorClass.
type TextChildAnchorClass struct {
	native unsafe.Pointer
}

func (recv *TextChildAnchorClass) ToC() unsafe.Pointer {
	return recv.native
}

// TextIter is a representation of the C record GtkTextIter.
type TextIter struct {
	native unsafe.Pointer
}

func (recv *TextIter) ToC() unsafe.Pointer {
	return recv.native
}

// TextMarkClass is a representation of the C record GtkTextMarkClass.
type TextMarkClass struct {
	native unsafe.Pointer
}

func (recv *TextMarkClass) ToC() unsafe.Pointer {
	return recv.native
}

// TextTagClass is a representation of the C record GtkTextTagClass.
type TextTagClass struct {
	native unsafe.Pointer
}

func (recv *TextTagClass) ToC() unsafe.Pointer {
	return recv.native
}

// TextTagPrivate is a representation of the C record GtkTextTagPrivate.
type TextTagPrivate struct {
	native unsafe.Pointer
}

func (recv *TextTagPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// TextTagTableClass is a representation of the C record GtkTextTagTableClass.
type TextTagTableClass struct {
	native unsafe.Pointer
}

func (recv *TextTagTableClass) ToC() unsafe.Pointer {
	return recv.native
}

// TextTagTablePrivate is a representation of the C record GtkTextTagTablePrivate.
type TextTagTablePrivate struct {
	native unsafe.Pointer
}

func (recv *TextTagTablePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// TextViewAccessibleClass is a representation of the C record GtkTextViewAccessibleClass.
type TextViewAccessibleClass struct {
	native unsafe.Pointer
}

func (recv *TextViewAccessibleClass) ToC() unsafe.Pointer {
	return recv.native
}

// TextViewAccessiblePrivate is a representation of the C record GtkTextViewAccessiblePrivate.
type TextViewAccessiblePrivate struct {
	native unsafe.Pointer
}

func (recv *TextViewAccessiblePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// TextViewClass is a representation of the C record GtkTextViewClass.
type TextViewClass struct {
	native unsafe.Pointer
}

func (recv *TextViewClass) ToC() unsafe.Pointer {
	return recv.native
}

// TextViewPrivate is a representation of the C record GtkTextViewPrivate.
type TextViewPrivate struct {
	native unsafe.Pointer
}

func (recv *TextViewPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ThemeEngine is a representation of the C record GtkThemeEngine.
type ThemeEngine struct {
	native unsafe.Pointer
}

func (recv *ThemeEngine) ToC() unsafe.Pointer {
	return recv.native
}

// ThemingEngineClass is a representation of the C record GtkThemingEngineClass.
type ThemingEngineClass struct {
	native unsafe.Pointer
}

func (recv *ThemingEngineClass) ToC() unsafe.Pointer {
	return recv.native
}

// ThemingEnginePrivate is a representation of the C record GtkThemingEnginePrivate.
type ThemingEnginePrivate struct {
	native unsafe.Pointer
}

func (recv *ThemingEnginePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ToggleActionClass is a representation of the C record GtkToggleActionClass.
type ToggleActionClass struct {
	native unsafe.Pointer
}

func (recv *ToggleActionClass) ToC() unsafe.Pointer {
	return recv.native
}

// ToggleActionEntry is a representation of the C record GtkToggleActionEntry.
type ToggleActionEntry struct {
	native unsafe.Pointer
}

func (recv *ToggleActionEntry) ToC() unsafe.Pointer {
	return recv.native
}

// ToggleActionPrivate is a representation of the C record GtkToggleActionPrivate.
type ToggleActionPrivate struct {
	native unsafe.Pointer
}

func (recv *ToggleActionPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ToggleButtonAccessibleClass is a representation of the C record GtkToggleButtonAccessibleClass.
type ToggleButtonAccessibleClass struct {
	native unsafe.Pointer
}

func (recv *ToggleButtonAccessibleClass) ToC() unsafe.Pointer {
	return recv.native
}

// ToggleButtonAccessiblePrivate is a representation of the C record GtkToggleButtonAccessiblePrivate.
type ToggleButtonAccessiblePrivate struct {
	native unsafe.Pointer
}

func (recv *ToggleButtonAccessiblePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ToggleButtonClass is a representation of the C record GtkToggleButtonClass.
type ToggleButtonClass struct {
	native unsafe.Pointer
}

func (recv *ToggleButtonClass) ToC() unsafe.Pointer {
	return recv.native
}

// ToggleButtonPrivate is a representation of the C record GtkToggleButtonPrivate.
type ToggleButtonPrivate struct {
	native unsafe.Pointer
}

func (recv *ToggleButtonPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ToggleToolButtonClass is a representation of the C record GtkToggleToolButtonClass.
type ToggleToolButtonClass struct {
	native unsafe.Pointer
}

func (recv *ToggleToolButtonClass) ToC() unsafe.Pointer {
	return recv.native
}

// ToggleToolButtonPrivate is a representation of the C record GtkToggleToolButtonPrivate.
type ToggleToolButtonPrivate struct {
	native unsafe.Pointer
}

func (recv *ToggleToolButtonPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ToolButtonClass is a representation of the C record GtkToolButtonClass.
type ToolButtonClass struct {
	native unsafe.Pointer
}

func (recv *ToolButtonClass) ToC() unsafe.Pointer {
	return recv.native
}

// ToolButtonPrivate is a representation of the C record GtkToolButtonPrivate.
type ToolButtonPrivate struct {
	native unsafe.Pointer
}

func (recv *ToolButtonPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ToolItemClass is a representation of the C record GtkToolItemClass.
type ToolItemClass struct {
	native unsafe.Pointer
}

func (recv *ToolItemClass) ToC() unsafe.Pointer {
	return recv.native
}

// ToolItemGroupClass is a representation of the C record GtkToolItemGroupClass.
type ToolItemGroupClass struct {
	native unsafe.Pointer
}

func (recv *ToolItemGroupClass) ToC() unsafe.Pointer {
	return recv.native
}

// ToolItemGroupPrivate is a representation of the C record GtkToolItemGroupPrivate.
type ToolItemGroupPrivate struct {
	native unsafe.Pointer
}

func (recv *ToolItemGroupPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ToolItemPrivate is a representation of the C record GtkToolItemPrivate.
type ToolItemPrivate struct {
	native unsafe.Pointer
}

func (recv *ToolItemPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ToolPaletteClass is a representation of the C record GtkToolPaletteClass.
type ToolPaletteClass struct {
	native unsafe.Pointer
}

func (recv *ToolPaletteClass) ToC() unsafe.Pointer {
	return recv.native
}

// ToolPalettePrivate is a representation of the C record GtkToolPalettePrivate.
type ToolPalettePrivate struct {
	native unsafe.Pointer
}

func (recv *ToolPalettePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ToolShellIface is a representation of the C record GtkToolShellIface.
type ToolShellIface struct {
	native unsafe.Pointer
}

func (recv *ToolShellIface) ToC() unsafe.Pointer {
	return recv.native
}

// ToolbarClass is a representation of the C record GtkToolbarClass.
type ToolbarClass struct {
	native unsafe.Pointer
}

func (recv *ToolbarClass) ToC() unsafe.Pointer {
	return recv.native
}

// ToolbarPrivate is a representation of the C record GtkToolbarPrivate.
type ToolbarPrivate struct {
	native unsafe.Pointer
}

func (recv *ToolbarPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ToplevelAccessibleClass is a representation of the C record GtkToplevelAccessibleClass.
type ToplevelAccessibleClass struct {
	native unsafe.Pointer
}

func (recv *ToplevelAccessibleClass) ToC() unsafe.Pointer {
	return recv.native
}

// ToplevelAccessiblePrivate is a representation of the C record GtkToplevelAccessiblePrivate.
type ToplevelAccessiblePrivate struct {
	native unsafe.Pointer
}

func (recv *ToplevelAccessiblePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// TreeDragDestIface is a representation of the C record GtkTreeDragDestIface.
type TreeDragDestIface struct {
	native unsafe.Pointer
}

func (recv *TreeDragDestIface) ToC() unsafe.Pointer {
	return recv.native
}

// TreeDragSourceIface is a representation of the C record GtkTreeDragSourceIface.
type TreeDragSourceIface struct {
	native unsafe.Pointer
}

func (recv *TreeDragSourceIface) ToC() unsafe.Pointer {
	return recv.native
}

// TreeIter is a representation of the C record GtkTreeIter.
type TreeIter struct {
	native unsafe.Pointer
}

func (recv *TreeIter) ToC() unsafe.Pointer {
	return recv.native
}

// TreeModelFilterClass is a representation of the C record GtkTreeModelFilterClass.
type TreeModelFilterClass struct {
	native unsafe.Pointer
}

func (recv *TreeModelFilterClass) ToC() unsafe.Pointer {
	return recv.native
}

// TreeModelFilterPrivate is a representation of the C record GtkTreeModelFilterPrivate.
type TreeModelFilterPrivate struct {
	native unsafe.Pointer
}

func (recv *TreeModelFilterPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// TreeModelIface is a representation of the C record GtkTreeModelIface.
type TreeModelIface struct {
	native unsafe.Pointer
}

func (recv *TreeModelIface) ToC() unsafe.Pointer {
	return recv.native
}

// TreeModelSortClass is a representation of the C record GtkTreeModelSortClass.
type TreeModelSortClass struct {
	native unsafe.Pointer
}

func (recv *TreeModelSortClass) ToC() unsafe.Pointer {
	return recv.native
}

// TreeModelSortPrivate is a representation of the C record GtkTreeModelSortPrivate.
type TreeModelSortPrivate struct {
	native unsafe.Pointer
}

func (recv *TreeModelSortPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// TreePath is a representation of the C record GtkTreePath.
type TreePath struct {
	native unsafe.Pointer
}

func (recv *TreePath) ToC() unsafe.Pointer {
	return recv.native
}

// TreeRowReference is a representation of the C record GtkTreeRowReference.
type TreeRowReference struct {
	native unsafe.Pointer
}

func (recv *TreeRowReference) ToC() unsafe.Pointer {
	return recv.native
}

// TreeSelectionClass is a representation of the C record GtkTreeSelectionClass.
type TreeSelectionClass struct {
	native unsafe.Pointer
}

func (recv *TreeSelectionClass) ToC() unsafe.Pointer {
	return recv.native
}

// TreeSelectionPrivate is a representation of the C record GtkTreeSelectionPrivate.
type TreeSelectionPrivate struct {
	native unsafe.Pointer
}

func (recv *TreeSelectionPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// TreeSortableIface is a representation of the C record GtkTreeSortableIface.
type TreeSortableIface struct {
	native unsafe.Pointer
}

func (recv *TreeSortableIface) ToC() unsafe.Pointer {
	return recv.native
}

// TreeStoreClass is a representation of the C record GtkTreeStoreClass.
type TreeStoreClass struct {
	native unsafe.Pointer
}

func (recv *TreeStoreClass) ToC() unsafe.Pointer {
	return recv.native
}

// TreeStorePrivate is a representation of the C record GtkTreeStorePrivate.
type TreeStorePrivate struct {
	native unsafe.Pointer
}

func (recv *TreeStorePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// TreeViewAccessibleClass is a representation of the C record GtkTreeViewAccessibleClass.
type TreeViewAccessibleClass struct {
	native unsafe.Pointer
}

func (recv *TreeViewAccessibleClass) ToC() unsafe.Pointer {
	return recv.native
}

// TreeViewAccessiblePrivate is a representation of the C record GtkTreeViewAccessiblePrivate.
type TreeViewAccessiblePrivate struct {
	native unsafe.Pointer
}

func (recv *TreeViewAccessiblePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// TreeViewClass is a representation of the C record GtkTreeViewClass.
type TreeViewClass struct {
	native unsafe.Pointer
}

func (recv *TreeViewClass) ToC() unsafe.Pointer {
	return recv.native
}

// TreeViewColumnClass is a representation of the C record GtkTreeViewColumnClass.
type TreeViewColumnClass struct {
	native unsafe.Pointer
}

func (recv *TreeViewColumnClass) ToC() unsafe.Pointer {
	return recv.native
}

// TreeViewColumnPrivate is a representation of the C record GtkTreeViewColumnPrivate.
type TreeViewColumnPrivate struct {
	native unsafe.Pointer
}

func (recv *TreeViewColumnPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// TreeViewPrivate is a representation of the C record GtkTreeViewPrivate.
type TreeViewPrivate struct {
	native unsafe.Pointer
}

func (recv *TreeViewPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// UIManagerClass is a representation of the C record GtkUIManagerClass.
type UIManagerClass struct {
	native unsafe.Pointer
}

func (recv *UIManagerClass) ToC() unsafe.Pointer {
	return recv.native
}

// UIManagerPrivate is a representation of the C record GtkUIManagerPrivate.
type UIManagerPrivate struct {
	native unsafe.Pointer
}

func (recv *UIManagerPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// VBoxClass is a representation of the C record GtkVBoxClass.
type VBoxClass struct {
	native unsafe.Pointer
}

func (recv *VBoxClass) ToC() unsafe.Pointer {
	return recv.native
}

// VButtonBoxClass is a representation of the C record GtkVButtonBoxClass.
type VButtonBoxClass struct {
	native unsafe.Pointer
}

func (recv *VButtonBoxClass) ToC() unsafe.Pointer {
	return recv.native
}

// VPanedClass is a representation of the C record GtkVPanedClass.
type VPanedClass struct {
	native unsafe.Pointer
}

func (recv *VPanedClass) ToC() unsafe.Pointer {
	return recv.native
}

// VScaleClass is a representation of the C record GtkVScaleClass.
type VScaleClass struct {
	native unsafe.Pointer
}

func (recv *VScaleClass) ToC() unsafe.Pointer {
	return recv.native
}

// VScrollbarClass is a representation of the C record GtkVScrollbarClass.
type VScrollbarClass struct {
	native unsafe.Pointer
}

func (recv *VScrollbarClass) ToC() unsafe.Pointer {
	return recv.native
}

// VSeparatorClass is a representation of the C record GtkVSeparatorClass.
type VSeparatorClass struct {
	native unsafe.Pointer
}

func (recv *VSeparatorClass) ToC() unsafe.Pointer {
	return recv.native
}

// ViewportClass is a representation of the C record GtkViewportClass.
type ViewportClass struct {
	native unsafe.Pointer
}

func (recv *ViewportClass) ToC() unsafe.Pointer {
	return recv.native
}

// ViewportPrivate is a representation of the C record GtkViewportPrivate.
type ViewportPrivate struct {
	native unsafe.Pointer
}

func (recv *ViewportPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// VolumeButtonClass is a representation of the C record GtkVolumeButtonClass.
type VolumeButtonClass struct {
	native unsafe.Pointer
}

func (recv *VolumeButtonClass) ToC() unsafe.Pointer {
	return recv.native
}

// WidgetAccessibleClass is a representation of the C record GtkWidgetAccessibleClass.
type WidgetAccessibleClass struct {
	native unsafe.Pointer
}

func (recv *WidgetAccessibleClass) ToC() unsafe.Pointer {
	return recv.native
}

// WidgetAccessiblePrivate is a representation of the C record GtkWidgetAccessiblePrivate.
type WidgetAccessiblePrivate struct {
	native unsafe.Pointer
}

func (recv *WidgetAccessiblePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// WidgetClass is a representation of the C record GtkWidgetClass.
type WidgetClass struct {
	native unsafe.Pointer
}

func (recv *WidgetClass) ToC() unsafe.Pointer {
	return recv.native
}

// WidgetClassPrivate is a representation of the C record GtkWidgetClassPrivate.
type WidgetClassPrivate struct {
	native unsafe.Pointer
}

func (recv *WidgetClassPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// WidgetPath is a representation of the C record GtkWidgetPath.
type WidgetPath struct {
	native unsafe.Pointer
}

func (recv *WidgetPath) ToC() unsafe.Pointer {
	return recv.native
}

// WidgetPrivate is a representation of the C record GtkWidgetPrivate.
type WidgetPrivate struct {
	native unsafe.Pointer
}

func (recv *WidgetPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// WindowAccessibleClass is a representation of the C record GtkWindowAccessibleClass.
type WindowAccessibleClass struct {
	native unsafe.Pointer
}

func (recv *WindowAccessibleClass) ToC() unsafe.Pointer {
	return recv.native
}

// WindowAccessiblePrivate is a representation of the C record GtkWindowAccessiblePrivate.
type WindowAccessiblePrivate struct {
	native unsafe.Pointer
}

func (recv *WindowAccessiblePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// WindowClass is a representation of the C record GtkWindowClass.
type WindowClass struct {
	native unsafe.Pointer
}

func (recv *WindowClass) ToC() unsafe.Pointer {
	return recv.native
}

// WindowGeometryInfo is a representation of the C record GtkWindowGeometryInfo.
type WindowGeometryInfo struct {
	native unsafe.Pointer
}

func (recv *WindowGeometryInfo) ToC() unsafe.Pointer {
	return recv.native
}

// WindowGroupClass is a representation of the C record GtkWindowGroupClass.
type WindowGroupClass struct {
	native unsafe.Pointer
}

func (recv *WindowGroupClass) ToC() unsafe.Pointer {
	return recv.native
}

// WindowGroupPrivate is a representation of the C record GtkWindowGroupPrivate.
type WindowGroupPrivate struct {
	native unsafe.Pointer
}

func (recv *WindowGroupPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// WindowPrivate is a representation of the C record GtkWindowPrivate.
type WindowPrivate struct {
	native unsafe.Pointer
}

func (recv *WindowPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// AboutDialog is a representation of the C record GtkAboutDialog.
type AboutDialog struct {
	native unsafe.Pointer
}

func (recv *AboutDialog) ToC() unsafe.Pointer {
	return recv.native
}

// AccelGroup is a representation of the C record GtkAccelGroup.
type AccelGroup struct {
	native unsafe.Pointer
}

func (recv *AccelGroup) ToC() unsafe.Pointer {
	return recv.native
}

// AccelLabel is a representation of the C record GtkAccelLabel.
type AccelLabel struct {
	native unsafe.Pointer
}

func (recv *AccelLabel) ToC() unsafe.Pointer {
	return recv.native
}

// AccelMap is a representation of the C record GtkAccelMap.
type AccelMap struct {
	native unsafe.Pointer
}

func (recv *AccelMap) ToC() unsafe.Pointer {
	return recv.native
}

// Accessible is a representation of the C record GtkAccessible.
type Accessible struct {
	native unsafe.Pointer
}

func (recv *Accessible) ToC() unsafe.Pointer {
	return recv.native
}

// Action is a representation of the C record GtkAction.
type Action struct {
	native unsafe.Pointer
}

func (recv *Action) ToC() unsafe.Pointer {
	return recv.native
}

// ActionBar is a representation of the C record GtkActionBar.
type ActionBar struct {
	native unsafe.Pointer
}

func (recv *ActionBar) ToC() unsafe.Pointer {
	return recv.native
}

// ActionGroup is a representation of the C record GtkActionGroup.
type ActionGroup struct {
	native unsafe.Pointer
}

func (recv *ActionGroup) ToC() unsafe.Pointer {
	return recv.native
}

// Adjustment is a representation of the C record GtkAdjustment.
type Adjustment struct {
	native unsafe.Pointer
}

func (recv *Adjustment) ToC() unsafe.Pointer {
	return recv.native
}

// Alignment is a representation of the C record GtkAlignment.
type Alignment struct {
	native unsafe.Pointer
}

func (recv *Alignment) ToC() unsafe.Pointer {
	return recv.native
}

// AppChooserButton is a representation of the C record GtkAppChooserButton.
type AppChooserButton struct {
	native unsafe.Pointer
}

func (recv *AppChooserButton) ToC() unsafe.Pointer {
	return recv.native
}

// AppChooserDialog is a representation of the C record GtkAppChooserDialog.
type AppChooserDialog struct {
	native unsafe.Pointer
}

func (recv *AppChooserDialog) ToC() unsafe.Pointer {
	return recv.native
}

// AppChooserWidget is a representation of the C record GtkAppChooserWidget.
type AppChooserWidget struct {
	native unsafe.Pointer
}

func (recv *AppChooserWidget) ToC() unsafe.Pointer {
	return recv.native
}

// Application is a representation of the C record GtkApplication.
type Application struct {
	native unsafe.Pointer
}

func (recv *Application) ToC() unsafe.Pointer {
	return recv.native
}

// ApplicationWindow is a representation of the C record GtkApplicationWindow.
type ApplicationWindow struct {
	native unsafe.Pointer
}

func (recv *ApplicationWindow) ToC() unsafe.Pointer {
	return recv.native
}

// Arrow is a representation of the C record GtkArrow.
type Arrow struct {
	native unsafe.Pointer
}

func (recv *Arrow) ToC() unsafe.Pointer {
	return recv.native
}

// ArrowAccessible is a representation of the C record GtkArrowAccessible.
type ArrowAccessible struct {
	native unsafe.Pointer
}

func (recv *ArrowAccessible) ToC() unsafe.Pointer {
	return recv.native
}

// AspectFrame is a representation of the C record GtkAspectFrame.
type AspectFrame struct {
	native unsafe.Pointer
}

func (recv *AspectFrame) ToC() unsafe.Pointer {
	return recv.native
}

// Assistant is a representation of the C record GtkAssistant.
type Assistant struct {
	native unsafe.Pointer
}

func (recv *Assistant) ToC() unsafe.Pointer {
	return recv.native
}

// Bin is a representation of the C record GtkBin.
type Bin struct {
	native unsafe.Pointer
}

func (recv *Bin) ToC() unsafe.Pointer {
	return recv.native
}

// BooleanCellAccessible is a representation of the C record GtkBooleanCellAccessible.
type BooleanCellAccessible struct {
	native unsafe.Pointer
}

func (recv *BooleanCellAccessible) ToC() unsafe.Pointer {
	return recv.native
}

// Box is a representation of the C record GtkBox.
type Box struct {
	native unsafe.Pointer
}

func (recv *Box) ToC() unsafe.Pointer {
	return recv.native
}

// Builder is a representation of the C record GtkBuilder.
type Builder struct {
	native unsafe.Pointer
}

func (recv *Builder) ToC() unsafe.Pointer {
	return recv.native
}

// Button is a representation of the C record GtkButton.
type Button struct {
	native unsafe.Pointer
}

func (recv *Button) ToC() unsafe.Pointer {
	return recv.native
}

// ButtonAccessible is a representation of the C record GtkButtonAccessible.
type ButtonAccessible struct {
	native unsafe.Pointer
}

func (recv *ButtonAccessible) ToC() unsafe.Pointer {
	return recv.native
}

// ButtonBox is a representation of the C record GtkButtonBox.
type ButtonBox struct {
	native unsafe.Pointer
}

func (recv *ButtonBox) ToC() unsafe.Pointer {
	return recv.native
}

// Calendar is a representation of the C record GtkCalendar.
type Calendar struct {
	native unsafe.Pointer
}

func (recv *Calendar) ToC() unsafe.Pointer {
	return recv.native
}

// CellAccessible is a representation of the C record GtkCellAccessible.
type CellAccessible struct {
	native unsafe.Pointer
}

func (recv *CellAccessible) ToC() unsafe.Pointer {
	return recv.native
}

// CellArea is a representation of the C record GtkCellArea.
type CellArea struct {
	native unsafe.Pointer
}

func (recv *CellArea) ToC() unsafe.Pointer {
	return recv.native
}

// CellAreaBox is a representation of the C record GtkCellAreaBox.
type CellAreaBox struct {
	native unsafe.Pointer
}

func (recv *CellAreaBox) ToC() unsafe.Pointer {
	return recv.native
}

// CellAreaContext is a representation of the C record GtkCellAreaContext.
type CellAreaContext struct {
	native unsafe.Pointer
}

func (recv *CellAreaContext) ToC() unsafe.Pointer {
	return recv.native
}

// CellRenderer is a representation of the C record GtkCellRenderer.
type CellRenderer struct {
	native unsafe.Pointer
}

func (recv *CellRenderer) ToC() unsafe.Pointer {
	return recv.native
}

// CellRendererAccel is a representation of the C record GtkCellRendererAccel.
type CellRendererAccel struct {
	native unsafe.Pointer
}

func (recv *CellRendererAccel) ToC() unsafe.Pointer {
	return recv.native
}

// CellRendererCombo is a representation of the C record GtkCellRendererCombo.
type CellRendererCombo struct {
	native unsafe.Pointer
}

func (recv *CellRendererCombo) ToC() unsafe.Pointer {
	return recv.native
}

// CellRendererPixbuf is a representation of the C record GtkCellRendererPixbuf.
type CellRendererPixbuf struct {
	native unsafe.Pointer
}

func (recv *CellRendererPixbuf) ToC() unsafe.Pointer {
	return recv.native
}

// CellRendererProgress is a representation of the C record GtkCellRendererProgress.
type CellRendererProgress struct {
	native unsafe.Pointer
}

func (recv *CellRendererProgress) ToC() unsafe.Pointer {
	return recv.native
}

// CellRendererSpin is a representation of the C record GtkCellRendererSpin.
type CellRendererSpin struct {
	native unsafe.Pointer
}

func (recv *CellRendererSpin) ToC() unsafe.Pointer {
	return recv.native
}

// CellRendererSpinner is a representation of the C record GtkCellRendererSpinner.
type CellRendererSpinner struct {
	native unsafe.Pointer
}

func (recv *CellRendererSpinner) ToC() unsafe.Pointer {
	return recv.native
}

// CellRendererText is a representation of the C record GtkCellRendererText.
type CellRendererText struct {
	native unsafe.Pointer
}

func (recv *CellRendererText) ToC() unsafe.Pointer {
	return recv.native
}

// CellRendererToggle is a representation of the C record GtkCellRendererToggle.
type CellRendererToggle struct {
	native unsafe.Pointer
}

func (recv *CellRendererToggle) ToC() unsafe.Pointer {
	return recv.native
}

// CellView is a representation of the C record GtkCellView.
type CellView struct {
	native unsafe.Pointer
}

func (recv *CellView) ToC() unsafe.Pointer {
	return recv.native
}

// CheckButton is a representation of the C record GtkCheckButton.
type CheckButton struct {
	native unsafe.Pointer
}

func (recv *CheckButton) ToC() unsafe.Pointer {
	return recv.native
}

// CheckMenuItem is a representation of the C record GtkCheckMenuItem.
type CheckMenuItem struct {
	native unsafe.Pointer
}

func (recv *CheckMenuItem) ToC() unsafe.Pointer {
	return recv.native
}

// CheckMenuItemAccessible is a representation of the C record GtkCheckMenuItemAccessible.
type CheckMenuItemAccessible struct {
	native unsafe.Pointer
}

func (recv *CheckMenuItemAccessible) ToC() unsafe.Pointer {
	return recv.native
}

// Clipboard is a representation of the C record GtkClipboard.
type Clipboard struct {
	native unsafe.Pointer
}

func (recv *Clipboard) ToC() unsafe.Pointer {
	return recv.native
}

// ColorButton is a representation of the C record GtkColorButton.
type ColorButton struct {
	native unsafe.Pointer
}

func (recv *ColorButton) ToC() unsafe.Pointer {
	return recv.native
}

// ColorSelection is a representation of the C record GtkColorSelection.
type ColorSelection struct {
	native unsafe.Pointer
}

func (recv *ColorSelection) ToC() unsafe.Pointer {
	return recv.native
}

// ColorSelectionDialog is a representation of the C record GtkColorSelectionDialog.
type ColorSelectionDialog struct {
	native unsafe.Pointer
}

func (recv *ColorSelectionDialog) ToC() unsafe.Pointer {
	return recv.native
}

// ComboBox is a representation of the C record GtkComboBox.
type ComboBox struct {
	native unsafe.Pointer
}

func (recv *ComboBox) ToC() unsafe.Pointer {
	return recv.native
}

// ComboBoxAccessible is a representation of the C record GtkComboBoxAccessible.
type ComboBoxAccessible struct {
	native unsafe.Pointer
}

func (recv *ComboBoxAccessible) ToC() unsafe.Pointer {
	return recv.native
}

// ComboBoxText is a representation of the C record GtkComboBoxText.
type ComboBoxText struct {
	native unsafe.Pointer
}

func (recv *ComboBoxText) ToC() unsafe.Pointer {
	return recv.native
}

// Container is a representation of the C record GtkContainer.
type Container struct {
	native unsafe.Pointer
}

func (recv *Container) ToC() unsafe.Pointer {
	return recv.native
}

// ContainerAccessible is a representation of the C record GtkContainerAccessible.
type ContainerAccessible struct {
	native unsafe.Pointer
}

func (recv *ContainerAccessible) ToC() unsafe.Pointer {
	return recv.native
}

// ContainerCellAccessible is a representation of the C record GtkContainerCellAccessible.
type ContainerCellAccessible struct {
	native unsafe.Pointer
}

func (recv *ContainerCellAccessible) ToC() unsafe.Pointer {
	return recv.native
}

// CssProvider is a representation of the C record GtkCssProvider.
type CssProvider struct {
	native unsafe.Pointer
}

func (recv *CssProvider) ToC() unsafe.Pointer {
	return recv.native
}

// Dialog is a representation of the C record GtkDialog.
type Dialog struct {
	native unsafe.Pointer
}

func (recv *Dialog) ToC() unsafe.Pointer {
	return recv.native
}

// DrawingArea is a representation of the C record GtkDrawingArea.
type DrawingArea struct {
	native unsafe.Pointer
}

func (recv *DrawingArea) ToC() unsafe.Pointer {
	return recv.native
}

// Entry is a representation of the C record GtkEntry.
type Entry struct {
	native unsafe.Pointer
}

func (recv *Entry) ToC() unsafe.Pointer {
	return recv.native
}

// EntryAccessible is a representation of the C record GtkEntryAccessible.
type EntryAccessible struct {
	native unsafe.Pointer
}

func (recv *EntryAccessible) ToC() unsafe.Pointer {
	return recv.native
}

// EntryCompletion is a representation of the C record GtkEntryCompletion.
type EntryCompletion struct {
	native unsafe.Pointer
}

func (recv *EntryCompletion) ToC() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : EntryIconAccessible : blacklisted

// EventBox is a representation of the C record GtkEventBox.
type EventBox struct {
	native unsafe.Pointer
}

func (recv *EventBox) ToC() unsafe.Pointer {
	return recv.native
}

// EventController is a representation of the C record GtkEventController.
type EventController struct {
	native unsafe.Pointer
}

func (recv *EventController) ToC() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : EventControllerKey : blacklisted

// UNSUPPORTED : EventControllerMotion : blacklisted

// UNSUPPORTED : EventControllerScroll : blacklisted

// Expander is a representation of the C record GtkExpander.
type Expander struct {
	native unsafe.Pointer
}

func (recv *Expander) ToC() unsafe.Pointer {
	return recv.native
}

// ExpanderAccessible is a representation of the C record GtkExpanderAccessible.
type ExpanderAccessible struct {
	native unsafe.Pointer
}

func (recv *ExpanderAccessible) ToC() unsafe.Pointer {
	return recv.native
}

// FileChooserButton is a representation of the C record GtkFileChooserButton.
type FileChooserButton struct {
	native unsafe.Pointer
}

func (recv *FileChooserButton) ToC() unsafe.Pointer {
	return recv.native
}

// FileChooserDialog is a representation of the C record GtkFileChooserDialog.
type FileChooserDialog struct {
	native unsafe.Pointer
}

func (recv *FileChooserDialog) ToC() unsafe.Pointer {
	return recv.native
}

// FileChooserNative is a representation of the C record GtkFileChooserNative.
type FileChooserNative struct {
	native unsafe.Pointer
}

func (recv *FileChooserNative) ToC() unsafe.Pointer {
	return recv.native
}

// FileChooserWidget is a representation of the C record GtkFileChooserWidget.
type FileChooserWidget struct {
	native unsafe.Pointer
}

func (recv *FileChooserWidget) ToC() unsafe.Pointer {
	return recv.native
}

// FileFilter is a representation of the C record GtkFileFilter.
type FileFilter struct {
	native unsafe.Pointer
}

func (recv *FileFilter) ToC() unsafe.Pointer {
	return recv.native
}

// Fixed is a representation of the C record GtkFixed.
type Fixed struct {
	native unsafe.Pointer
}

func (recv *Fixed) ToC() unsafe.Pointer {
	return recv.native
}

// FlowBox is a representation of the C record GtkFlowBox.
type FlowBox struct {
	native unsafe.Pointer
}

func (recv *FlowBox) ToC() unsafe.Pointer {
	return recv.native
}

// FlowBoxAccessible is a representation of the C record GtkFlowBoxAccessible.
type FlowBoxAccessible struct {
	native unsafe.Pointer
}

func (recv *FlowBoxAccessible) ToC() unsafe.Pointer {
	return recv.native
}

// FlowBoxChild is a representation of the C record GtkFlowBoxChild.
type FlowBoxChild struct {
	native unsafe.Pointer
}

func (recv *FlowBoxChild) ToC() unsafe.Pointer {
	return recv.native
}

// FlowBoxChildAccessible is a representation of the C record GtkFlowBoxChildAccessible.
type FlowBoxChildAccessible struct {
	native unsafe.Pointer
}

func (recv *FlowBoxChildAccessible) ToC() unsafe.Pointer {
	return recv.native
}

// FontButton is a representation of the C record GtkFontButton.
type FontButton struct {
	native unsafe.Pointer
}

func (recv *FontButton) ToC() unsafe.Pointer {
	return recv.native
}

// FontSelection is a representation of the C record GtkFontSelection.
type FontSelection struct {
	native unsafe.Pointer
}

func (recv *FontSelection) ToC() unsafe.Pointer {
	return recv.native
}

// FontSelectionDialog is a representation of the C record GtkFontSelectionDialog.
type FontSelectionDialog struct {
	native unsafe.Pointer
}

func (recv *FontSelectionDialog) ToC() unsafe.Pointer {
	return recv.native
}

// Frame is a representation of the C record GtkFrame.
type Frame struct {
	native unsafe.Pointer
}

func (recv *Frame) ToC() unsafe.Pointer {
	return recv.native
}

// FrameAccessible is a representation of the C record GtkFrameAccessible.
type FrameAccessible struct {
	native unsafe.Pointer
}

func (recv *FrameAccessible) ToC() unsafe.Pointer {
	return recv.native
}

// Gesture is a representation of the C record GtkGesture.
type Gesture struct {
	native unsafe.Pointer
}

func (recv *Gesture) ToC() unsafe.Pointer {
	return recv.native
}

// GestureDrag is a representation of the C record GtkGestureDrag.
type GestureDrag struct {
	native unsafe.Pointer
}

func (recv *GestureDrag) ToC() unsafe.Pointer {
	return recv.native
}

// GestureLongPress is a representation of the C record GtkGestureLongPress.
type GestureLongPress struct {
	native unsafe.Pointer
}

func (recv *GestureLongPress) ToC() unsafe.Pointer {
	return recv.native
}

// GestureMultiPress is a representation of the C record GtkGestureMultiPress.
type GestureMultiPress struct {
	native unsafe.Pointer
}

func (recv *GestureMultiPress) ToC() unsafe.Pointer {
	return recv.native
}

// GesturePan is a representation of the C record GtkGesturePan.
type GesturePan struct {
	native unsafe.Pointer
}

func (recv *GesturePan) ToC() unsafe.Pointer {
	return recv.native
}

// GestureRotate is a representation of the C record GtkGestureRotate.
type GestureRotate struct {
	native unsafe.Pointer
}

func (recv *GestureRotate) ToC() unsafe.Pointer {
	return recv.native
}

// GestureSingle is a representation of the C record GtkGestureSingle.
type GestureSingle struct {
	native unsafe.Pointer
}

func (recv *GestureSingle) ToC() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : GestureStylus : blacklisted

// GestureSwipe is a representation of the C record GtkGestureSwipe.
type GestureSwipe struct {
	native unsafe.Pointer
}

func (recv *GestureSwipe) ToC() unsafe.Pointer {
	return recv.native
}

// GestureZoom is a representation of the C record GtkGestureZoom.
type GestureZoom struct {
	native unsafe.Pointer
}

func (recv *GestureZoom) ToC() unsafe.Pointer {
	return recv.native
}

// Grid is a representation of the C record GtkGrid.
type Grid struct {
	native unsafe.Pointer
}

func (recv *Grid) ToC() unsafe.Pointer {
	return recv.native
}

// HBox is a representation of the C record GtkHBox.
type HBox struct {
	native unsafe.Pointer
}

func (recv *HBox) ToC() unsafe.Pointer {
	return recv.native
}

// HButtonBox is a representation of the C record GtkHButtonBox.
type HButtonBox struct {
	native unsafe.Pointer
}

func (recv *HButtonBox) ToC() unsafe.Pointer {
	return recv.native
}

// HPaned is a representation of the C record GtkHPaned.
type HPaned struct {
	native unsafe.Pointer
}

func (recv *HPaned) ToC() unsafe.Pointer {
	return recv.native
}

// HSV is a representation of the C record GtkHSV.
type HSV struct {
	native unsafe.Pointer
}

func (recv *HSV) ToC() unsafe.Pointer {
	return recv.native
}

// HScale is a representation of the C record GtkHScale.
type HScale struct {
	native unsafe.Pointer
}

func (recv *HScale) ToC() unsafe.Pointer {
	return recv.native
}

// HScrollbar is a representation of the C record GtkHScrollbar.
type HScrollbar struct {
	native unsafe.Pointer
}

func (recv *HScrollbar) ToC() unsafe.Pointer {
	return recv.native
}

// HSeparator is a representation of the C record GtkHSeparator.
type HSeparator struct {
	native unsafe.Pointer
}

func (recv *HSeparator) ToC() unsafe.Pointer {
	return recv.native
}

// HandleBox is a representation of the C record GtkHandleBox.
type HandleBox struct {
	native unsafe.Pointer
}

func (recv *HandleBox) ToC() unsafe.Pointer {
	return recv.native
}

// HeaderBar is a representation of the C record GtkHeaderBar.
type HeaderBar struct {
	native unsafe.Pointer
}

func (recv *HeaderBar) ToC() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : HeaderBarAccessible : blacklisted

// IMContext is a representation of the C record GtkIMContext.
type IMContext struct {
	native unsafe.Pointer
}

func (recv *IMContext) ToC() unsafe.Pointer {
	return recv.native
}

// IMContextSimple is a representation of the C record GtkIMContextSimple.
type IMContextSimple struct {
	native unsafe.Pointer
}

func (recv *IMContextSimple) ToC() unsafe.Pointer {
	return recv.native
}

// IMMulticontext is a representation of the C record GtkIMMulticontext.
type IMMulticontext struct {
	native unsafe.Pointer
}

func (recv *IMMulticontext) ToC() unsafe.Pointer {
	return recv.native
}

// IconFactory is a representation of the C record GtkIconFactory.
type IconFactory struct {
	native unsafe.Pointer
}

func (recv *IconFactory) ToC() unsafe.Pointer {
	return recv.native
}

// IconInfo is a representation of the C record GtkIconInfo.
type IconInfo struct {
	native unsafe.Pointer
}

func (recv *IconInfo) ToC() unsafe.Pointer {
	return recv.native
}

// IconTheme is a representation of the C record GtkIconTheme.
type IconTheme struct {
	native unsafe.Pointer
}

func (recv *IconTheme) ToC() unsafe.Pointer {
	return recv.native
}

// IconView is a representation of the C record GtkIconView.
type IconView struct {
	native unsafe.Pointer
}

func (recv *IconView) ToC() unsafe.Pointer {
	return recv.native
}

// IconViewAccessible is a representation of the C record GtkIconViewAccessible.
type IconViewAccessible struct {
	native unsafe.Pointer
}

func (recv *IconViewAccessible) ToC() unsafe.Pointer {
	return recv.native
}

// Image is a representation of the C record GtkImage.
type Image struct {
	native unsafe.Pointer
}

func (recv *Image) ToC() unsafe.Pointer {
	return recv.native
}

// ImageAccessible is a representation of the C record GtkImageAccessible.
type ImageAccessible struct {
	native unsafe.Pointer
}

func (recv *ImageAccessible) ToC() unsafe.Pointer {
	return recv.native
}

// ImageCellAccessible is a representation of the C record GtkImageCellAccessible.
type ImageCellAccessible struct {
	native unsafe.Pointer
}

func (recv *ImageCellAccessible) ToC() unsafe.Pointer {
	return recv.native
}

// ImageMenuItem is a representation of the C record GtkImageMenuItem.
type ImageMenuItem struct {
	native unsafe.Pointer
}

func (recv *ImageMenuItem) ToC() unsafe.Pointer {
	return recv.native
}

// InfoBar is a representation of the C record GtkInfoBar.
type InfoBar struct {
	native unsafe.Pointer
}

func (recv *InfoBar) ToC() unsafe.Pointer {
	return recv.native
}

// Invisible is a representation of the C record GtkInvisible.
type Invisible struct {
	native unsafe.Pointer
}

func (recv *Invisible) ToC() unsafe.Pointer {
	return recv.native
}

// Label is a representation of the C record GtkLabel.
type Label struct {
	native unsafe.Pointer
}

func (recv *Label) ToC() unsafe.Pointer {
	return recv.native
}

// LabelAccessible is a representation of the C record GtkLabelAccessible.
type LabelAccessible struct {
	native unsafe.Pointer
}

func (recv *LabelAccessible) ToC() unsafe.Pointer {
	return recv.native
}

// Layout is a representation of the C record GtkLayout.
type Layout struct {
	native unsafe.Pointer
}

func (recv *Layout) ToC() unsafe.Pointer {
	return recv.native
}

// LevelBar is a representation of the C record GtkLevelBar.
type LevelBar struct {
	native unsafe.Pointer
}

func (recv *LevelBar) ToC() unsafe.Pointer {
	return recv.native
}

// LevelBarAccessible is a representation of the C record GtkLevelBarAccessible.
type LevelBarAccessible struct {
	native unsafe.Pointer
}

func (recv *LevelBarAccessible) ToC() unsafe.Pointer {
	return recv.native
}

// LinkButton is a representation of the C record GtkLinkButton.
type LinkButton struct {
	native unsafe.Pointer
}

func (recv *LinkButton) ToC() unsafe.Pointer {
	return recv.native
}

// LinkButtonAccessible is a representation of the C record GtkLinkButtonAccessible.
type LinkButtonAccessible struct {
	native unsafe.Pointer
}

func (recv *LinkButtonAccessible) ToC() unsafe.Pointer {
	return recv.native
}

// ListBox is a representation of the C record GtkListBox.
type ListBox struct {
	native unsafe.Pointer
}

func (recv *ListBox) ToC() unsafe.Pointer {
	return recv.native
}

// ListBoxAccessible is a representation of the C record GtkListBoxAccessible.
type ListBoxAccessible struct {
	native unsafe.Pointer
}

func (recv *ListBoxAccessible) ToC() unsafe.Pointer {
	return recv.native
}

// ListBoxRow is a representation of the C record GtkListBoxRow.
type ListBoxRow struct {
	native unsafe.Pointer
}

func (recv *ListBoxRow) ToC() unsafe.Pointer {
	return recv.native
}

// ListBoxRowAccessible is a representation of the C record GtkListBoxRowAccessible.
type ListBoxRowAccessible struct {
	native unsafe.Pointer
}

func (recv *ListBoxRowAccessible) ToC() unsafe.Pointer {
	return recv.native
}

// ListStore is a representation of the C record GtkListStore.
type ListStore struct {
	native unsafe.Pointer
}

func (recv *ListStore) ToC() unsafe.Pointer {
	return recv.native
}

// LockButton is a representation of the C record GtkLockButton.
type LockButton struct {
	native unsafe.Pointer
}

func (recv *LockButton) ToC() unsafe.Pointer {
	return recv.native
}

// LockButtonAccessible is a representation of the C record GtkLockButtonAccessible.
type LockButtonAccessible struct {
	native unsafe.Pointer
}

func (recv *LockButtonAccessible) ToC() unsafe.Pointer {
	return recv.native
}

// Menu is a representation of the C record GtkMenu.
type Menu struct {
	native unsafe.Pointer
}

func (recv *Menu) ToC() unsafe.Pointer {
	return recv.native
}

// MenuAccessible is a representation of the C record GtkMenuAccessible.
type MenuAccessible struct {
	native unsafe.Pointer
}

func (recv *MenuAccessible) ToC() unsafe.Pointer {
	return recv.native
}

// MenuBar is a representation of the C record GtkMenuBar.
type MenuBar struct {
	native unsafe.Pointer
}

func (recv *MenuBar) ToC() unsafe.Pointer {
	return recv.native
}

// MenuButton is a representation of the C record GtkMenuButton.
type MenuButton struct {
	native unsafe.Pointer
}

func (recv *MenuButton) ToC() unsafe.Pointer {
	return recv.native
}

// MenuButtonAccessible is a representation of the C record GtkMenuButtonAccessible.
type MenuButtonAccessible struct {
	native unsafe.Pointer
}

func (recv *MenuButtonAccessible) ToC() unsafe.Pointer {
	return recv.native
}

// MenuItem is a representation of the C record GtkMenuItem.
type MenuItem struct {
	native unsafe.Pointer
}

func (recv *MenuItem) ToC() unsafe.Pointer {
	return recv.native
}

// MenuItemAccessible is a representation of the C record GtkMenuItemAccessible.
type MenuItemAccessible struct {
	native unsafe.Pointer
}

func (recv *MenuItemAccessible) ToC() unsafe.Pointer {
	return recv.native
}

// MenuShell is a representation of the C record GtkMenuShell.
type MenuShell struct {
	native unsafe.Pointer
}

func (recv *MenuShell) ToC() unsafe.Pointer {
	return recv.native
}

// MenuShellAccessible is a representation of the C record GtkMenuShellAccessible.
type MenuShellAccessible struct {
	native unsafe.Pointer
}

func (recv *MenuShellAccessible) ToC() unsafe.Pointer {
	return recv.native
}

// MenuToolButton is a representation of the C record GtkMenuToolButton.
type MenuToolButton struct {
	native unsafe.Pointer
}

func (recv *MenuToolButton) ToC() unsafe.Pointer {
	return recv.native
}

// MessageDialog is a representation of the C record GtkMessageDialog.
type MessageDialog struct {
	native unsafe.Pointer
}

func (recv *MessageDialog) ToC() unsafe.Pointer {
	return recv.native
}

// Misc is a representation of the C record GtkMisc.
type Misc struct {
	native unsafe.Pointer
}

func (recv *Misc) ToC() unsafe.Pointer {
	return recv.native
}

// ModelButton is a representation of the C record GtkModelButton.
type ModelButton struct {
	native unsafe.Pointer
}

func (recv *ModelButton) ToC() unsafe.Pointer {
	return recv.native
}

// MountOperation is a representation of the C record GtkMountOperation.
type MountOperation struct {
	native unsafe.Pointer
}

func (recv *MountOperation) ToC() unsafe.Pointer {
	return recv.native
}

// NativeDialog is a representation of the C record GtkNativeDialog.
type NativeDialog struct {
	native unsafe.Pointer
}

func (recv *NativeDialog) ToC() unsafe.Pointer {
	return recv.native
}

// Notebook is a representation of the C record GtkNotebook.
type Notebook struct {
	native unsafe.Pointer
}

func (recv *Notebook) ToC() unsafe.Pointer {
	return recv.native
}

// NotebookAccessible is a representation of the C record GtkNotebookAccessible.
type NotebookAccessible struct {
	native unsafe.Pointer
}

func (recv *NotebookAccessible) ToC() unsafe.Pointer {
	return recv.native
}

// NotebookPageAccessible is a representation of the C record GtkNotebookPageAccessible.
type NotebookPageAccessible struct {
	native unsafe.Pointer
}

func (recv *NotebookPageAccessible) ToC() unsafe.Pointer {
	return recv.native
}

// NumerableIcon is a representation of the C record GtkNumerableIcon.
type NumerableIcon struct {
	native unsafe.Pointer
}

func (recv *NumerableIcon) ToC() unsafe.Pointer {
	return recv.native
}

// OffscreenWindow is a representation of the C record GtkOffscreenWindow.
type OffscreenWindow struct {
	native unsafe.Pointer
}

func (recv *OffscreenWindow) ToC() unsafe.Pointer {
	return recv.native
}

// Overlay is a representation of the C record GtkOverlay.
type Overlay struct {
	native unsafe.Pointer
}

func (recv *Overlay) ToC() unsafe.Pointer {
	return recv.native
}

// PadController is a representation of the C record GtkPadController.
type PadController struct {
	native unsafe.Pointer
}

func (recv *PadController) ToC() unsafe.Pointer {
	return recv.native
}

// PageSetup is a representation of the C record GtkPageSetup.
type PageSetup struct {
	native unsafe.Pointer
}

func (recv *PageSetup) ToC() unsafe.Pointer {
	return recv.native
}

// Paned is a representation of the C record GtkPaned.
type Paned struct {
	native unsafe.Pointer
}

func (recv *Paned) ToC() unsafe.Pointer {
	return recv.native
}

// PanedAccessible is a representation of the C record GtkPanedAccessible.
type PanedAccessible struct {
	native unsafe.Pointer
}

func (recv *PanedAccessible) ToC() unsafe.Pointer {
	return recv.native
}

// PlacesSidebar is a representation of the C record GtkPlacesSidebar.
type PlacesSidebar struct {
	native unsafe.Pointer
}

func (recv *PlacesSidebar) ToC() unsafe.Pointer {
	return recv.native
}

// Plug is a representation of the C record GtkPlug.
type Plug struct {
	native unsafe.Pointer
}

func (recv *Plug) ToC() unsafe.Pointer {
	return recv.native
}

// PopoverAccessible is a representation of the C record GtkPopoverAccessible.
type PopoverAccessible struct {
	native unsafe.Pointer
}

func (recv *PopoverAccessible) ToC() unsafe.Pointer {
	return recv.native
}

// PopoverMenu is a representation of the C record GtkPopoverMenu.
type PopoverMenu struct {
	native unsafe.Pointer
}

func (recv *PopoverMenu) ToC() unsafe.Pointer {
	return recv.native
}

// PrintContext is a representation of the C record GtkPrintContext.
type PrintContext struct {
	native unsafe.Pointer
}

func (recv *PrintContext) ToC() unsafe.Pointer {
	return recv.native
}

// PrintOperation is a representation of the C record GtkPrintOperation.
type PrintOperation struct {
	native unsafe.Pointer
}

func (recv *PrintOperation) ToC() unsafe.Pointer {
	return recv.native
}

// PrintSettings is a representation of the C record GtkPrintSettings.
type PrintSettings struct {
	native unsafe.Pointer
}

func (recv *PrintSettings) ToC() unsafe.Pointer {
	return recv.native
}

// ProgressBar is a representation of the C record GtkProgressBar.
type ProgressBar struct {
	native unsafe.Pointer
}

func (recv *ProgressBar) ToC() unsafe.Pointer {
	return recv.native
}

// ProgressBarAccessible is a representation of the C record GtkProgressBarAccessible.
type ProgressBarAccessible struct {
	native unsafe.Pointer
}

func (recv *ProgressBarAccessible) ToC() unsafe.Pointer {
	return recv.native
}

// RadioAction is a representation of the C record GtkRadioAction.
type RadioAction struct {
	native unsafe.Pointer
}

func (recv *RadioAction) ToC() unsafe.Pointer {
	return recv.native
}

// RadioButton is a representation of the C record GtkRadioButton.
type RadioButton struct {
	native unsafe.Pointer
}

func (recv *RadioButton) ToC() unsafe.Pointer {
	return recv.native
}

// RadioButtonAccessible is a representation of the C record GtkRadioButtonAccessible.
type RadioButtonAccessible struct {
	native unsafe.Pointer
}

func (recv *RadioButtonAccessible) ToC() unsafe.Pointer {
	return recv.native
}

// RadioMenuItem is a representation of the C record GtkRadioMenuItem.
type RadioMenuItem struct {
	native unsafe.Pointer
}

func (recv *RadioMenuItem) ToC() unsafe.Pointer {
	return recv.native
}

// RadioMenuItemAccessible is a representation of the C record GtkRadioMenuItemAccessible.
type RadioMenuItemAccessible struct {
	native unsafe.Pointer
}

func (recv *RadioMenuItemAccessible) ToC() unsafe.Pointer {
	return recv.native
}

// RadioToolButton is a representation of the C record GtkRadioToolButton.
type RadioToolButton struct {
	native unsafe.Pointer
}

func (recv *RadioToolButton) ToC() unsafe.Pointer {
	return recv.native
}

// Range is a representation of the C record GtkRange.
type Range struct {
	native unsafe.Pointer
}

func (recv *Range) ToC() unsafe.Pointer {
	return recv.native
}

// RangeAccessible is a representation of the C record GtkRangeAccessible.
type RangeAccessible struct {
	native unsafe.Pointer
}

func (recv *RangeAccessible) ToC() unsafe.Pointer {
	return recv.native
}

// RcStyle is a representation of the C record GtkRcStyle.
type RcStyle struct {
	native unsafe.Pointer
}

func (recv *RcStyle) ToC() unsafe.Pointer {
	return recv.native
}

// RecentAction is a representation of the C record GtkRecentAction.
type RecentAction struct {
	native unsafe.Pointer
}

func (recv *RecentAction) ToC() unsafe.Pointer {
	return recv.native
}

// RecentChooserDialog is a representation of the C record GtkRecentChooserDialog.
type RecentChooserDialog struct {
	native unsafe.Pointer
}

func (recv *RecentChooserDialog) ToC() unsafe.Pointer {
	return recv.native
}

// RecentChooserMenu is a representation of the C record GtkRecentChooserMenu.
type RecentChooserMenu struct {
	native unsafe.Pointer
}

func (recv *RecentChooserMenu) ToC() unsafe.Pointer {
	return recv.native
}

// RecentChooserWidget is a representation of the C record GtkRecentChooserWidget.
type RecentChooserWidget struct {
	native unsafe.Pointer
}

func (recv *RecentChooserWidget) ToC() unsafe.Pointer {
	return recv.native
}

// RecentFilter is a representation of the C record GtkRecentFilter.
type RecentFilter struct {
	native unsafe.Pointer
}

func (recv *RecentFilter) ToC() unsafe.Pointer {
	return recv.native
}

// RendererCellAccessible is a representation of the C record GtkRendererCellAccessible.
type RendererCellAccessible struct {
	native unsafe.Pointer
}

func (recv *RendererCellAccessible) ToC() unsafe.Pointer {
	return recv.native
}

// Revealer is a representation of the C record GtkRevealer.
type Revealer struct {
	native unsafe.Pointer
}

func (recv *Revealer) ToC() unsafe.Pointer {
	return recv.native
}

// Scale is a representation of the C record GtkScale.
type Scale struct {
	native unsafe.Pointer
}

func (recv *Scale) ToC() unsafe.Pointer {
	return recv.native
}

// ScaleAccessible is a representation of the C record GtkScaleAccessible.
type ScaleAccessible struct {
	native unsafe.Pointer
}

func (recv *ScaleAccessible) ToC() unsafe.Pointer {
	return recv.native
}

// ScaleButton is a representation of the C record GtkScaleButton.
type ScaleButton struct {
	native unsafe.Pointer
}

func (recv *ScaleButton) ToC() unsafe.Pointer {
	return recv.native
}

// ScaleButtonAccessible is a representation of the C record GtkScaleButtonAccessible.
type ScaleButtonAccessible struct {
	native unsafe.Pointer
}

func (recv *ScaleButtonAccessible) ToC() unsafe.Pointer {
	return recv.native
}

// Scrollbar is a representation of the C record GtkScrollbar.
type Scrollbar struct {
	native unsafe.Pointer
}

func (recv *Scrollbar) ToC() unsafe.Pointer {
	return recv.native
}

// ScrolledWindow is a representation of the C record GtkScrolledWindow.
type ScrolledWindow struct {
	native unsafe.Pointer
}

func (recv *ScrolledWindow) ToC() unsafe.Pointer {
	return recv.native
}

// ScrolledWindowAccessible is a representation of the C record GtkScrolledWindowAccessible.
type ScrolledWindowAccessible struct {
	native unsafe.Pointer
}

func (recv *ScrolledWindowAccessible) ToC() unsafe.Pointer {
	return recv.native
}

// Separator is a representation of the C record GtkSeparator.
type Separator struct {
	native unsafe.Pointer
}

func (recv *Separator) ToC() unsafe.Pointer {
	return recv.native
}

// SeparatorMenuItem is a representation of the C record GtkSeparatorMenuItem.
type SeparatorMenuItem struct {
	native unsafe.Pointer
}

func (recv *SeparatorMenuItem) ToC() unsafe.Pointer {
	return recv.native
}

// SeparatorToolItem is a representation of the C record GtkSeparatorToolItem.
type SeparatorToolItem struct {
	native unsafe.Pointer
}

func (recv *SeparatorToolItem) ToC() unsafe.Pointer {
	return recv.native
}

// Settings is a representation of the C record GtkSettings.
type Settings struct {
	native unsafe.Pointer
}

func (recv *Settings) ToC() unsafe.Pointer {
	return recv.native
}

// ShortcutLabel is a representation of the C record GtkShortcutLabel.
type ShortcutLabel struct {
	native unsafe.Pointer
}

func (recv *ShortcutLabel) ToC() unsafe.Pointer {
	return recv.native
}

// ShortcutsGroup is a representation of the C record GtkShortcutsGroup.
type ShortcutsGroup struct {
	native unsafe.Pointer
}

func (recv *ShortcutsGroup) ToC() unsafe.Pointer {
	return recv.native
}

// ShortcutsSection is a representation of the C record GtkShortcutsSection.
type ShortcutsSection struct {
	native unsafe.Pointer
}

func (recv *ShortcutsSection) ToC() unsafe.Pointer {
	return recv.native
}

// ShortcutsShortcut is a representation of the C record GtkShortcutsShortcut.
type ShortcutsShortcut struct {
	native unsafe.Pointer
}

func (recv *ShortcutsShortcut) ToC() unsafe.Pointer {
	return recv.native
}

// ShortcutsWindow is a representation of the C record GtkShortcutsWindow.
type ShortcutsWindow struct {
	native unsafe.Pointer
}

func (recv *ShortcutsWindow) ToC() unsafe.Pointer {
	return recv.native
}

// SizeGroup is a representation of the C record GtkSizeGroup.
type SizeGroup struct {
	native unsafe.Pointer
}

func (recv *SizeGroup) ToC() unsafe.Pointer {
	return recv.native
}

// Socket is a representation of the C record GtkSocket.
type Socket struct {
	native unsafe.Pointer
}

func (recv *Socket) ToC() unsafe.Pointer {
	return recv.native
}

// SpinButton is a representation of the C record GtkSpinButton.
type SpinButton struct {
	native unsafe.Pointer
}

func (recv *SpinButton) ToC() unsafe.Pointer {
	return recv.native
}

// SpinButtonAccessible is a representation of the C record GtkSpinButtonAccessible.
type SpinButtonAccessible struct {
	native unsafe.Pointer
}

func (recv *SpinButtonAccessible) ToC() unsafe.Pointer {
	return recv.native
}

// Spinner is a representation of the C record GtkSpinner.
type Spinner struct {
	native unsafe.Pointer
}

func (recv *Spinner) ToC() unsafe.Pointer {
	return recv.native
}

// SpinnerAccessible is a representation of the C record GtkSpinnerAccessible.
type SpinnerAccessible struct {
	native unsafe.Pointer
}

func (recv *SpinnerAccessible) ToC() unsafe.Pointer {
	return recv.native
}

// Stack is a representation of the C record GtkStack.
type Stack struct {
	native unsafe.Pointer
}

func (recv *Stack) ToC() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : StackAccessible : blacklisted

// StackSwitcher is a representation of the C record GtkStackSwitcher.
type StackSwitcher struct {
	native unsafe.Pointer
}

func (recv *StackSwitcher) ToC() unsafe.Pointer {
	return recv.native
}

// StatusIcon is a representation of the C record GtkStatusIcon.
type StatusIcon struct {
	native unsafe.Pointer
}

func (recv *StatusIcon) ToC() unsafe.Pointer {
	return recv.native
}

// Statusbar is a representation of the C record GtkStatusbar.
type Statusbar struct {
	native unsafe.Pointer
}

func (recv *Statusbar) ToC() unsafe.Pointer {
	return recv.native
}

// StatusbarAccessible is a representation of the C record GtkStatusbarAccessible.
type StatusbarAccessible struct {
	native unsafe.Pointer
}

func (recv *StatusbarAccessible) ToC() unsafe.Pointer {
	return recv.native
}

// Style is a representation of the C record GtkStyle.
type Style struct {
	native unsafe.Pointer
}

func (recv *Style) ToC() unsafe.Pointer {
	return recv.native
}

// StyleContext is a representation of the C record GtkStyleContext.
type StyleContext struct {
	native unsafe.Pointer
}

func (recv *StyleContext) ToC() unsafe.Pointer {
	return recv.native
}

// StyleProperties is a representation of the C record GtkStyleProperties.
type StyleProperties struct {
	native unsafe.Pointer
}

func (recv *StyleProperties) ToC() unsafe.Pointer {
	return recv.native
}

// Switch is a representation of the C record GtkSwitch.
type Switch struct {
	native unsafe.Pointer
}

func (recv *Switch) ToC() unsafe.Pointer {
	return recv.native
}

// SwitchAccessible is a representation of the C record GtkSwitchAccessible.
type SwitchAccessible struct {
	native unsafe.Pointer
}

func (recv *SwitchAccessible) ToC() unsafe.Pointer {
	return recv.native
}

// Table is a representation of the C record GtkTable.
type Table struct {
	native unsafe.Pointer
}

func (recv *Table) ToC() unsafe.Pointer {
	return recv.native
}

// TearoffMenuItem is a representation of the C record GtkTearoffMenuItem.
type TearoffMenuItem struct {
	native unsafe.Pointer
}

func (recv *TearoffMenuItem) ToC() unsafe.Pointer {
	return recv.native
}

// TextBuffer is a representation of the C record GtkTextBuffer.
type TextBuffer struct {
	native unsafe.Pointer
}

func (recv *TextBuffer) ToC() unsafe.Pointer {
	return recv.native
}

// TextCellAccessible is a representation of the C record GtkTextCellAccessible.
type TextCellAccessible struct {
	native unsafe.Pointer
}

func (recv *TextCellAccessible) ToC() unsafe.Pointer {
	return recv.native
}

// TextChildAnchor is a representation of the C record GtkTextChildAnchor.
type TextChildAnchor struct {
	native unsafe.Pointer
}

func (recv *TextChildAnchor) ToC() unsafe.Pointer {
	return recv.native
}

// TextMark is a representation of the C record GtkTextMark.
type TextMark struct {
	native unsafe.Pointer
}

func (recv *TextMark) ToC() unsafe.Pointer {
	return recv.native
}

// TextTag is a representation of the C record GtkTextTag.
type TextTag struct {
	native unsafe.Pointer
}

func (recv *TextTag) ToC() unsafe.Pointer {
	return recv.native
}

// TextTagTable is a representation of the C record GtkTextTagTable.
type TextTagTable struct {
	native unsafe.Pointer
}

func (recv *TextTagTable) ToC() unsafe.Pointer {
	return recv.native
}

// TextView is a representation of the C record GtkTextView.
type TextView struct {
	native unsafe.Pointer
}

func (recv *TextView) ToC() unsafe.Pointer {
	return recv.native
}

// TextViewAccessible is a representation of the C record GtkTextViewAccessible.
type TextViewAccessible struct {
	native unsafe.Pointer
}

func (recv *TextViewAccessible) ToC() unsafe.Pointer {
	return recv.native
}

// ThemingEngine is a representation of the C record GtkThemingEngine.
type ThemingEngine struct {
	native unsafe.Pointer
}

func (recv *ThemingEngine) ToC() unsafe.Pointer {
	return recv.native
}

// ToggleAction is a representation of the C record GtkToggleAction.
type ToggleAction struct {
	native unsafe.Pointer
}

func (recv *ToggleAction) ToC() unsafe.Pointer {
	return recv.native
}

// ToggleButton is a representation of the C record GtkToggleButton.
type ToggleButton struct {
	native unsafe.Pointer
}

func (recv *ToggleButton) ToC() unsafe.Pointer {
	return recv.native
}

// ToggleButtonAccessible is a representation of the C record GtkToggleButtonAccessible.
type ToggleButtonAccessible struct {
	native unsafe.Pointer
}

func (recv *ToggleButtonAccessible) ToC() unsafe.Pointer {
	return recv.native
}

// ToggleToolButton is a representation of the C record GtkToggleToolButton.
type ToggleToolButton struct {
	native unsafe.Pointer
}

func (recv *ToggleToolButton) ToC() unsafe.Pointer {
	return recv.native
}

// ToolButton is a representation of the C record GtkToolButton.
type ToolButton struct {
	native unsafe.Pointer
}

func (recv *ToolButton) ToC() unsafe.Pointer {
	return recv.native
}

// ToolItem is a representation of the C record GtkToolItem.
type ToolItem struct {
	native unsafe.Pointer
}

func (recv *ToolItem) ToC() unsafe.Pointer {
	return recv.native
}

// Toolbar is a representation of the C record GtkToolbar.
type Toolbar struct {
	native unsafe.Pointer
}

func (recv *Toolbar) ToC() unsafe.Pointer {
	return recv.native
}

// Tooltip is a representation of the C record GtkTooltip.
type Tooltip struct {
	native unsafe.Pointer
}

func (recv *Tooltip) ToC() unsafe.Pointer {
	return recv.native
}

// ToplevelAccessible is a representation of the C record GtkToplevelAccessible.
type ToplevelAccessible struct {
	native unsafe.Pointer
}

func (recv *ToplevelAccessible) ToC() unsafe.Pointer {
	return recv.native
}

// TreeModelFilter is a representation of the C record GtkTreeModelFilter.
type TreeModelFilter struct {
	native unsafe.Pointer
}

func (recv *TreeModelFilter) ToC() unsafe.Pointer {
	return recv.native
}

// TreeModelSort is a representation of the C record GtkTreeModelSort.
type TreeModelSort struct {
	native unsafe.Pointer
}

func (recv *TreeModelSort) ToC() unsafe.Pointer {
	return recv.native
}

// TreeSelection is a representation of the C record GtkTreeSelection.
type TreeSelection struct {
	native unsafe.Pointer
}

func (recv *TreeSelection) ToC() unsafe.Pointer {
	return recv.native
}

// TreeStore is a representation of the C record GtkTreeStore.
type TreeStore struct {
	native unsafe.Pointer
}

func (recv *TreeStore) ToC() unsafe.Pointer {
	return recv.native
}

// TreeView is a representation of the C record GtkTreeView.
type TreeView struct {
	native unsafe.Pointer
}

func (recv *TreeView) ToC() unsafe.Pointer {
	return recv.native
}

// TreeViewAccessible is a representation of the C record GtkTreeViewAccessible.
type TreeViewAccessible struct {
	native unsafe.Pointer
}

func (recv *TreeViewAccessible) ToC() unsafe.Pointer {
	return recv.native
}

// TreeViewColumn is a representation of the C record GtkTreeViewColumn.
type TreeViewColumn struct {
	native unsafe.Pointer
}

func (recv *TreeViewColumn) ToC() unsafe.Pointer {
	return recv.native
}

// UIManager is a representation of the C record GtkUIManager.
type UIManager struct {
	native unsafe.Pointer
}

func (recv *UIManager) ToC() unsafe.Pointer {
	return recv.native
}

// VBox is a representation of the C record GtkVBox.
type VBox struct {
	native unsafe.Pointer
}

func (recv *VBox) ToC() unsafe.Pointer {
	return recv.native
}

// VButtonBox is a representation of the C record GtkVButtonBox.
type VButtonBox struct {
	native unsafe.Pointer
}

func (recv *VButtonBox) ToC() unsafe.Pointer {
	return recv.native
}

// VPaned is a representation of the C record GtkVPaned.
type VPaned struct {
	native unsafe.Pointer
}

func (recv *VPaned) ToC() unsafe.Pointer {
	return recv.native
}

// VScale is a representation of the C record GtkVScale.
type VScale struct {
	native unsafe.Pointer
}

func (recv *VScale) ToC() unsafe.Pointer {
	return recv.native
}

// VScrollbar is a representation of the C record GtkVScrollbar.
type VScrollbar struct {
	native unsafe.Pointer
}

func (recv *VScrollbar) ToC() unsafe.Pointer {
	return recv.native
}

// VSeparator is a representation of the C record GtkVSeparator.
type VSeparator struct {
	native unsafe.Pointer
}

func (recv *VSeparator) ToC() unsafe.Pointer {
	return recv.native
}

// Viewport is a representation of the C record GtkViewport.
type Viewport struct {
	native unsafe.Pointer
}

func (recv *Viewport) ToC() unsafe.Pointer {
	return recv.native
}

// VolumeButton is a representation of the C record GtkVolumeButton.
type VolumeButton struct {
	native unsafe.Pointer
}

func (recv *VolumeButton) ToC() unsafe.Pointer {
	return recv.native
}

// Widget is a representation of the C record GtkWidget.
type Widget struct {
	native unsafe.Pointer
}

func (recv *Widget) ToC() unsafe.Pointer {
	return recv.native
}

// WidgetAccessible is a representation of the C record GtkWidgetAccessible.
type WidgetAccessible struct {
	native unsafe.Pointer
}

func (recv *WidgetAccessible) ToC() unsafe.Pointer {
	return recv.native
}

// Window is a representation of the C record GtkWindow.
type Window struct {
	native unsafe.Pointer
}

func (recv *Window) ToC() unsafe.Pointer {
	return recv.native
}

// WindowAccessible is a representation of the C record GtkWindowAccessible.
type WindowAccessible struct {
	native unsafe.Pointer
}

func (recv *WindowAccessible) ToC() unsafe.Pointer {
	return recv.native
}

// WindowGroup is a representation of the C record GtkWindowGroup.
type WindowGroup struct {
	native unsafe.Pointer
}

func (recv *WindowGroup) ToC() unsafe.Pointer {
	return recv.native
}

// Activatable is a representation of the C interface GtkActivatable.
type Activatable struct {
	native unsafe.Pointer
}

// AppChooser is a representation of the C interface GtkAppChooser.
type AppChooser struct {
	native unsafe.Pointer
}

// Buildable is a representation of the C interface GtkBuildable.
type Buildable struct {
	native unsafe.Pointer
}

// CellAccessibleParent is a representation of the C interface GtkCellAccessibleParent.
type CellAccessibleParent struct {
	native unsafe.Pointer
}

// CellEditable is a representation of the C interface GtkCellEditable.
type CellEditable struct {
	native unsafe.Pointer
}

// CellLayout is a representation of the C interface GtkCellLayout.
type CellLayout struct {
	native unsafe.Pointer
}

// Editable is a representation of the C interface GtkEditable.
type Editable struct {
	native unsafe.Pointer
}

// FileChooser is a representation of the C interface GtkFileChooser.
type FileChooser struct {
	native unsafe.Pointer
}

// FontChooser is a representation of the C interface GtkFontChooser.
type FontChooser struct {
	native unsafe.Pointer
}

// Orientable is a representation of the C interface GtkOrientable.
type Orientable struct {
	native unsafe.Pointer
}

// PrintOperationPreview is a representation of the C interface GtkPrintOperationPreview.
type PrintOperationPreview struct {
	native unsafe.Pointer
}

// RecentChooser is a representation of the C interface GtkRecentChooser.
type RecentChooser struct {
	native unsafe.Pointer
}

// Scrollable is a representation of the C interface GtkScrollable.
type Scrollable struct {
	native unsafe.Pointer
}

// StyleProvider is a representation of the C interface GtkStyleProvider.
type StyleProvider struct {
	native unsafe.Pointer
}

// ToolShell is a representation of the C interface GtkToolShell.
type ToolShell struct {
	native unsafe.Pointer
}

// TreeDragDest is a representation of the C interface GtkTreeDragDest.
type TreeDragDest struct {
	native unsafe.Pointer
}

// TreeDragSource is a representation of the C interface GtkTreeDragSource.
type TreeDragSource struct {
	native unsafe.Pointer
}

// TreeModel is a representation of the C interface GtkTreeModel.
type TreeModel struct {
	native unsafe.Pointer
}

// TreeSortable is a representation of the C interface GtkTreeSortable.
type TreeSortable struct {
	native unsafe.Pointer
}
