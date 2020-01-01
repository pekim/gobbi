// Code generated - DO NOT EDIT.
// +build gtk_2.22

package gtk

import (
	cairo "github.com/pekim/gobbi/lib/cairo"
	gdk "github.com/pekim/gobbi/lib/gdk"
	gdkpixbuf "github.com/pekim/gobbi/lib/gdkpixbuf"
	glib "github.com/pekim/gobbi/lib/glib"
	gobject "github.com/pekim/gobbi/lib/gobject"
	gdk1 "github.com/pekim/gobbi/lib/internal/c/gdk"
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

// STOCK_CAPS_LOCK_WARNING is a representation of the C constant GTK_STOCK_CAPS_LOCK_WARNING.
//
// since 2.16
const STOCK_CAPS_LOCK_WARNING = "gtk-caps-lock-warning"

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

// STOCK_DISCARD is a representation of the C constant GTK_STOCK_DISCARD.
//
// since 2.12
const STOCK_DISCARD = "gtk-discard"

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

// STOCK_FULLSCREEN is a representation of the C constant GTK_STOCK_FULLSCREEN.
//
// since 2.8
const STOCK_FULLSCREEN = "gtk-fullscreen"

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

// STOCK_INFO is a representation of the C constant GTK_STOCK_INFO.
//
// since 2.8
const STOCK_INFO = "gtk-info"

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

// STOCK_LEAVE_FULLSCREEN is a representation of the C constant GTK_STOCK_LEAVE_FULLSCREEN.
//
// since 2.8
const STOCK_LEAVE_FULLSCREEN = "gtk-leave-fullscreen"

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

// STOCK_ORIENTATION_LANDSCAPE is a representation of the C constant GTK_STOCK_ORIENTATION_LANDSCAPE.
//
// since 2.10
const STOCK_ORIENTATION_LANDSCAPE = "gtk-orientation-landscape"

// STOCK_ORIENTATION_PORTRAIT is a representation of the C constant GTK_STOCK_ORIENTATION_PORTRAIT.
//
// since 2.10
const STOCK_ORIENTATION_PORTRAIT = "gtk-orientation-portrait"

// STOCK_ORIENTATION_REVERSE_LANDSCAPE is a representation of the C constant GTK_STOCK_ORIENTATION_REVERSE_LANDSCAPE.
//
// since 2.10
const STOCK_ORIENTATION_REVERSE_LANDSCAPE = "gtk-orientation-reverse-landscape"

// STOCK_ORIENTATION_REVERSE_PORTRAIT is a representation of the C constant GTK_STOCK_ORIENTATION_REVERSE_PORTRAIT.
//
// since 2.10
const STOCK_ORIENTATION_REVERSE_PORTRAIT = "gtk-orientation-reverse-portrait"

// STOCK_PAGE_SETUP is a representation of the C constant GTK_STOCK_PAGE_SETUP.
//
// since 2.14
const STOCK_PAGE_SETUP = "gtk-page-setup"

// STOCK_PASTE is a representation of the C constant GTK_STOCK_PASTE.
const STOCK_PASTE = "gtk-paste"

// STOCK_PREFERENCES is a representation of the C constant GTK_STOCK_PREFERENCES.
const STOCK_PREFERENCES = "gtk-preferences"

// STOCK_PRINT is a representation of the C constant GTK_STOCK_PRINT.
const STOCK_PRINT = "gtk-print"

// STOCK_PRINT_ERROR is a representation of the C constant GTK_STOCK_PRINT_ERROR.
//
// since 2.14
const STOCK_PRINT_ERROR = "gtk-print-error"

// STOCK_PRINT_PAUSED is a representation of the C constant GTK_STOCK_PRINT_PAUSED.
//
// since 2.14
const STOCK_PRINT_PAUSED = "gtk-print-paused"

// STOCK_PRINT_PREVIEW is a representation of the C constant GTK_STOCK_PRINT_PREVIEW.
const STOCK_PRINT_PREVIEW = "gtk-print-preview"

// STOCK_PRINT_REPORT is a representation of the C constant GTK_STOCK_PRINT_REPORT.
//
// since 2.14
const STOCK_PRINT_REPORT = "gtk-print-report"

// STOCK_PRINT_WARNING is a representation of the C constant GTK_STOCK_PRINT_WARNING.
//
// since 2.14
const STOCK_PRINT_WARNING = "gtk-print-warning"

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

// STOCK_SELECT_ALL is a representation of the C constant GTK_STOCK_SELECT_ALL.
//
// since 2.10
const STOCK_SELECT_ALL = "gtk-select-all"

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

// EntryIconPosition is a representation of the C enumeration GtkEntryIconPosition.
type EntryIconPosition int

// EntryIconPosition_primary is a representation of the C enumeration member GTK_ENTRY_ICON_PRIMARY.
const EntryIconPosition_primary = EntryIconPosition(0)

// EntryIconPosition_secondary is a representation of the C enumeration member GTK_ENTRY_ICON_SECONDARY.
const EntryIconPosition_secondary = EntryIconPosition(1)

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

// FileChooserConfirmation is a representation of the C enumeration GtkFileChooserConfirmation.
type FileChooserConfirmation int

// FileChooserConfirmation_confirm is a representation of the C enumeration member GTK_FILE_CHOOSER_CONFIRMATION_CONFIRM.
const FileChooserConfirmation_confirm = FileChooserConfirmation(0)

// FileChooserConfirmation_accept_filename is a representation of the C enumeration member GTK_FILE_CHOOSER_CONFIRMATION_ACCEPT_FILENAME.
const FileChooserConfirmation_accept_filename = FileChooserConfirmation(1)

// FileChooserConfirmation_select_again is a representation of the C enumeration member GTK_FILE_CHOOSER_CONFIRMATION_SELECT_AGAIN.
const FileChooserConfirmation_select_again = FileChooserConfirmation(2)

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

// RecentChooserError is a representation of the C enumeration GtkRecentChooserError.
type RecentChooserError int

// RecentChooserError_not_found is a representation of the C enumeration member GTK_RECENT_CHOOSER_ERROR_NOT_FOUND.
const RecentChooserError_not_found = RecentChooserError(0)

// RecentChooserError_invalid_uri is a representation of the C enumeration member GTK_RECENT_CHOOSER_ERROR_INVALID_URI.
const RecentChooserError_invalid_uri = RecentChooserError(1)

// RecentManagerError is a representation of the C enumeration GtkRecentManagerError.
type RecentManagerError int

// RecentManagerError_not_found is a representation of the C enumeration member GTK_RECENT_MANAGER_ERROR_NOT_FOUND.
const RecentManagerError_not_found = RecentManagerError(0)

// RecentManagerError_invalid_uri is a representation of the C enumeration member GTK_RECENT_MANAGER_ERROR_INVALID_URI.
const RecentManagerError_invalid_uri = RecentManagerError(1)

// RecentManagerError_invalid_encoding is a representation of the C enumeration member GTK_RECENT_MANAGER_ERROR_INVALID_ENCODING.
const RecentManagerError_invalid_encoding = RecentManagerError(2)

// RecentManagerError_not_registered is a representation of the C enumeration member GTK_RECENT_MANAGER_ERROR_NOT_REGISTERED.
const RecentManagerError_not_registered = RecentManagerError(3)

// RecentManagerError_read is a representation of the C enumeration member GTK_RECENT_MANAGER_ERROR_READ.
const RecentManagerError_read = RecentManagerError(4)

// RecentManagerError_write is a representation of the C enumeration member GTK_RECENT_MANAGER_ERROR_WRITE.
const RecentManagerError_write = RecentManagerError(5)

// RecentManagerError_unknown is a representation of the C enumeration member GTK_RECENT_MANAGER_ERROR_UNKNOWN.
const RecentManagerError_unknown = RecentManagerError(6)

// RecentSortType is a representation of the C enumeration GtkRecentSortType.
type RecentSortType int

// RecentSortType_none is a representation of the C enumeration member GTK_RECENT_SORT_NONE.
const RecentSortType_none = RecentSortType(0)

// RecentSortType_mru is a representation of the C enumeration member GTK_RECENT_SORT_MRU.
const RecentSortType_mru = RecentSortType(1)

// RecentSortType_lru is a representation of the C enumeration member GTK_RECENT_SORT_LRU.
const RecentSortType_lru = RecentSortType(2)

// RecentSortType_custom is a representation of the C enumeration member GTK_RECENT_SORT_CUSTOM.
const RecentSortType_custom = RecentSortType(3)

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

func Fn_gtk_accel_groups_activate(object *gobject.Object, accelKey uint, accelMods int) {
	sys_object := unsafe.Pointer(object)
	sys_accelKey := uint(accelKey)
	sys_accelMods := int(accelMods)
}

func Fn_gtk_accel_groups_from_object(object *gobject.Object) {
	sys_object := unsafe.Pointer(object)
}

func Fn_gtk_accelerator_get_default_mod_mask() {}

func Fn_gtk_accelerator_get_label(acceleratorKey uint, acceleratorMods int) {
	sys_acceleratorKey := uint(acceleratorKey)
	sys_acceleratorMods := int(acceleratorMods)
}

func Fn_gtk_accelerator_name(acceleratorKey uint, acceleratorMods int) {
	sys_acceleratorKey := uint(acceleratorKey)
	sys_acceleratorMods := int(acceleratorMods)
}

func Fn_gtk_accelerator_parse(accelerator string) {
	sys_accelerator := string(accelerator)
}

// UNSUPPORTED : gtk_accelerator_parse_with_keycode : parameter 'accelerator_codes' is array parameter without length parameter

func Fn_gtk_accelerator_set_default_mod_mask(defaultModMask int) {
	sys_defaultModMask := int(defaultModMask)
}

func Fn_gtk_accelerator_valid(keyval uint, modifiers int) {
	sys_keyval := uint(keyval)
	sys_modifiers := int(modifiers)
}

func Fn_gtk_alternative_dialog_button_order(screen *gdk.Screen) {
	sys_screen := unsafe.Pointer(screen)
}

func Fn_gtk_bindings_activate(object *gobject.Object, keyval uint, modifiers int) {
	sys_object := unsafe.Pointer(object)
	sys_keyval := uint(keyval)
	sys_modifiers := int(modifiers)
}

func Fn_gtk_bindings_activate_event(object *gobject.Object, event *gdk.EventKey) {
	sys_object := unsafe.Pointer(object)
	sys_event := unsafe.Pointer(event)
}

func Fn_gtk_check_version(requiredMajor uint, requiredMinor uint, requiredMicro uint) {
	sys_requiredMajor := uint(requiredMajor)
	sys_requiredMinor := uint(requiredMinor)
	sys_requiredMicro := uint(requiredMicro)
}

func Fn_gtk_disable_setlocale() {}

func Fn_gtk_distribute_natural_allocation(extraSpace int, nRequestedSizes uint, sizes *RequestedSize) {
	sys_extraSpace := int(extraSpace)
	sys_nRequestedSizes := uint(nRequestedSizes)
	sys_sizes := unsafe.Pointer(sizes)
}

func Fn_gtk_drag_finish(context *gdk.DragContext, success bool, del bool, time uint32) {
	sys_context := unsafe.Pointer(context)
	sys_success := bool(success)
	sys_del := bool(del)
	sys_time := uint32(time)
}

func Fn_gtk_drag_get_source_widget(context *gdk.DragContext) {
	sys_context := unsafe.Pointer(context)
}

func Fn_gtk_drag_set_icon_default(context *gdk.DragContext) {
	sys_context := unsafe.Pointer(context)
}

func Fn_gtk_drag_set_icon_name(context *gdk.DragContext, iconName string, hotX int, hotY int) {
	sys_context := unsafe.Pointer(context)
	sys_iconName := string(iconName)
	sys_hotX := int(hotX)
	sys_hotY := int(hotY)
}

func Fn_gtk_drag_set_icon_pixbuf(context *gdk.DragContext, pixbuf *gdkpixbuf.Pixbuf, hotX int, hotY int) {
	sys_context := unsafe.Pointer(context)
	sys_pixbuf := unsafe.Pointer(pixbuf)
	sys_hotX := int(hotX)
	sys_hotY := int(hotY)
}

func Fn_gtk_drag_set_icon_stock(context *gdk.DragContext, stockId string, hotX int, hotY int) {
	sys_context := unsafe.Pointer(context)
	sys_stockId := string(stockId)
	sys_hotX := int(hotX)
	sys_hotY := int(hotY)
}

func Fn_gtk_drag_set_icon_surface(context *gdk.DragContext, surface *cairo.Surface) {
	sys_context := unsafe.Pointer(context)
	sys_surface := unsafe.Pointer(surface)
}

func Fn_gtk_drag_set_icon_widget(context *gdk.DragContext, widget *Widget, hotX int, hotY int) {
	sys_context := unsafe.Pointer(context)
	sys_widget := unsafe.Pointer(widget)
	sys_hotX := int(hotX)
	sys_hotY := int(hotY)
}

func Fn_gtk_events_pending() {}

func Fn_gtk_false() {}

func Fn_gtk_get_current_event() {}

func Fn_gtk_get_current_event_device() {}

func Fn_gtk_get_current_event_state() {}

func Fn_gtk_get_current_event_time() {}

func Fn_gtk_get_debug_flags() {}

func Fn_gtk_get_default_language() {}

func Fn_gtk_get_event_widget(event *gdk.Event) {
	sys_event := unsafe.Pointer(event)
}

func Fn_gtk_get_option_group(openDefaultDisplay bool) {
	sys_openDefaultDisplay := bool(openDefaultDisplay)
}

func Fn_gtk_grab_get_current() {}

// UNSUPPORTED : gtk_init : has array param, argv

// UNSUPPORTED : gtk_init_check : has array param, argv

// UNSUPPORTED : gtk_init_with_args : parameter 'entries' is array parameter without length parameter

// UNSUPPORTED : gtk_key_snooper_install : parameter 'snooper' is callback

func Fn_gtk_key_snooper_remove(snooperHandlerId uint) {
	sys_snooperHandlerId := uint(snooperHandlerId)
}

func Fn_gtk_main() {}

func Fn_gtk_main_do_event(event *gdk.Event) {
	sys_event := unsafe.Pointer(event)
}

func Fn_gtk_main_iteration() {}

func Fn_gtk_main_iteration_do(blocking bool) {
	sys_blocking := bool(blocking)
}

func Fn_gtk_main_level() {}

func Fn_gtk_main_quit() {}

func Fn_gtk_paint_arrow(style *Style, cr *cairo.Context, stateType int, shadowType int, widget *Widget, detail string, arrowType int, fill bool, x int, y int, width int, height int) {
	sys_style := unsafe.Pointer(style)
	sys_cr := unsafe.Pointer(cr)
	sys_stateType := int(stateType)
	sys_shadowType := int(shadowType)
	sys_widget := unsafe.Pointer(widget)
	sys_detail := string(detail)
	sys_arrowType := int(arrowType)
	sys_fill := bool(fill)
	sys_x := int(x)
	sys_y := int(y)
	sys_width := int(width)
	sys_height := int(height)
}

func Fn_gtk_paint_box(style *Style, cr *cairo.Context, stateType int, shadowType int, widget *Widget, detail string, x int, y int, width int, height int) {
	sys_style := unsafe.Pointer(style)
	sys_cr := unsafe.Pointer(cr)
	sys_stateType := int(stateType)
	sys_shadowType := int(shadowType)
	sys_widget := unsafe.Pointer(widget)
	sys_detail := string(detail)
	sys_x := int(x)
	sys_y := int(y)
	sys_width := int(width)
	sys_height := int(height)
}

func Fn_gtk_paint_box_gap(style *Style, cr *cairo.Context, stateType int, shadowType int, widget *Widget, detail string, x int, y int, width int, height int, gapSide int, gapX int, gapWidth int) {
	sys_style := unsafe.Pointer(style)
	sys_cr := unsafe.Pointer(cr)
	sys_stateType := int(stateType)
	sys_shadowType := int(shadowType)
	sys_widget := unsafe.Pointer(widget)
	sys_detail := string(detail)
	sys_x := int(x)
	sys_y := int(y)
	sys_width := int(width)
	sys_height := int(height)
	sys_gapSide := int(gapSide)
	sys_gapX := int(gapX)
	sys_gapWidth := int(gapWidth)
}

func Fn_gtk_paint_check(style *Style, cr *cairo.Context, stateType int, shadowType int, widget *Widget, detail string, x int, y int, width int, height int) {
	sys_style := unsafe.Pointer(style)
	sys_cr := unsafe.Pointer(cr)
	sys_stateType := int(stateType)
	sys_shadowType := int(shadowType)
	sys_widget := unsafe.Pointer(widget)
	sys_detail := string(detail)
	sys_x := int(x)
	sys_y := int(y)
	sys_width := int(width)
	sys_height := int(height)
}

func Fn_gtk_paint_diamond(style *Style, cr *cairo.Context, stateType int, shadowType int, widget *Widget, detail string, x int, y int, width int, height int) {
	sys_style := unsafe.Pointer(style)
	sys_cr := unsafe.Pointer(cr)
	sys_stateType := int(stateType)
	sys_shadowType := int(shadowType)
	sys_widget := unsafe.Pointer(widget)
	sys_detail := string(detail)
	sys_x := int(x)
	sys_y := int(y)
	sys_width := int(width)
	sys_height := int(height)
}

func Fn_gtk_paint_expander(style *Style, cr *cairo.Context, stateType int, widget *Widget, detail string, x int, y int, expanderStyle int) {
	sys_style := unsafe.Pointer(style)
	sys_cr := unsafe.Pointer(cr)
	sys_stateType := int(stateType)
	sys_widget := unsafe.Pointer(widget)
	sys_detail := string(detail)
	sys_x := int(x)
	sys_y := int(y)
	sys_expanderStyle := int(expanderStyle)
}

func Fn_gtk_paint_extension(style *Style, cr *cairo.Context, stateType int, shadowType int, widget *Widget, detail string, x int, y int, width int, height int, gapSide int) {
	sys_style := unsafe.Pointer(style)
	sys_cr := unsafe.Pointer(cr)
	sys_stateType := int(stateType)
	sys_shadowType := int(shadowType)
	sys_widget := unsafe.Pointer(widget)
	sys_detail := string(detail)
	sys_x := int(x)
	sys_y := int(y)
	sys_width := int(width)
	sys_height := int(height)
	sys_gapSide := int(gapSide)
}

func Fn_gtk_paint_flat_box(style *Style, cr *cairo.Context, stateType int, shadowType int, widget *Widget, detail string, x int, y int, width int, height int) {
	sys_style := unsafe.Pointer(style)
	sys_cr := unsafe.Pointer(cr)
	sys_stateType := int(stateType)
	sys_shadowType := int(shadowType)
	sys_widget := unsafe.Pointer(widget)
	sys_detail := string(detail)
	sys_x := int(x)
	sys_y := int(y)
	sys_width := int(width)
	sys_height := int(height)
}

func Fn_gtk_paint_focus(style *Style, cr *cairo.Context, stateType int, widget *Widget, detail string, x int, y int, width int, height int) {
	sys_style := unsafe.Pointer(style)
	sys_cr := unsafe.Pointer(cr)
	sys_stateType := int(stateType)
	sys_widget := unsafe.Pointer(widget)
	sys_detail := string(detail)
	sys_x := int(x)
	sys_y := int(y)
	sys_width := int(width)
	sys_height := int(height)
}

func Fn_gtk_paint_handle(style *Style, cr *cairo.Context, stateType int, shadowType int, widget *Widget, detail string, x int, y int, width int, height int, orientation int) {
	sys_style := unsafe.Pointer(style)
	sys_cr := unsafe.Pointer(cr)
	sys_stateType := int(stateType)
	sys_shadowType := int(shadowType)
	sys_widget := unsafe.Pointer(widget)
	sys_detail := string(detail)
	sys_x := int(x)
	sys_y := int(y)
	sys_width := int(width)
	sys_height := int(height)
	sys_orientation := int(orientation)
}

func Fn_gtk_paint_hline(style *Style, cr *cairo.Context, stateType int, widget *Widget, detail string, x1 int, x2 int, y int) {
	sys_style := unsafe.Pointer(style)
	sys_cr := unsafe.Pointer(cr)
	sys_stateType := int(stateType)
	sys_widget := unsafe.Pointer(widget)
	sys_detail := string(detail)
	sys_x1 := int(x1)
	sys_x2 := int(x2)
	sys_y := int(y)
}

func Fn_gtk_paint_layout(style *Style, cr *cairo.Context, stateType int, useText bool, widget *Widget, detail string, x int, y int, layout *pango.Layout) {
	sys_style := unsafe.Pointer(style)
	sys_cr := unsafe.Pointer(cr)
	sys_stateType := int(stateType)
	sys_useText := bool(useText)
	sys_widget := unsafe.Pointer(widget)
	sys_detail := string(detail)
	sys_x := int(x)
	sys_y := int(y)
	sys_layout := unsafe.Pointer(layout)
}

func Fn_gtk_paint_option(style *Style, cr *cairo.Context, stateType int, shadowType int, widget *Widget, detail string, x int, y int, width int, height int) {
	sys_style := unsafe.Pointer(style)
	sys_cr := unsafe.Pointer(cr)
	sys_stateType := int(stateType)
	sys_shadowType := int(shadowType)
	sys_widget := unsafe.Pointer(widget)
	sys_detail := string(detail)
	sys_x := int(x)
	sys_y := int(y)
	sys_width := int(width)
	sys_height := int(height)
}

func Fn_gtk_paint_resize_grip(style *Style, cr *cairo.Context, stateType int, widget *Widget, detail string, edge int, x int, y int, width int, height int) {
	sys_style := unsafe.Pointer(style)
	sys_cr := unsafe.Pointer(cr)
	sys_stateType := int(stateType)
	sys_widget := unsafe.Pointer(widget)
	sys_detail := string(detail)
	sys_edge := int(edge)
	sys_x := int(x)
	sys_y := int(y)
	sys_width := int(width)
	sys_height := int(height)
}

func Fn_gtk_paint_shadow(style *Style, cr *cairo.Context, stateType int, shadowType int, widget *Widget, detail string, x int, y int, width int, height int) {
	sys_style := unsafe.Pointer(style)
	sys_cr := unsafe.Pointer(cr)
	sys_stateType := int(stateType)
	sys_shadowType := int(shadowType)
	sys_widget := unsafe.Pointer(widget)
	sys_detail := string(detail)
	sys_x := int(x)
	sys_y := int(y)
	sys_width := int(width)
	sys_height := int(height)
}

func Fn_gtk_paint_shadow_gap(style *Style, cr *cairo.Context, stateType int, shadowType int, widget *Widget, detail string, x int, y int, width int, height int, gapSide int, gapX int, gapWidth int) {
	sys_style := unsafe.Pointer(style)
	sys_cr := unsafe.Pointer(cr)
	sys_stateType := int(stateType)
	sys_shadowType := int(shadowType)
	sys_widget := unsafe.Pointer(widget)
	sys_detail := string(detail)
	sys_x := int(x)
	sys_y := int(y)
	sys_width := int(width)
	sys_height := int(height)
	sys_gapSide := int(gapSide)
	sys_gapX := int(gapX)
	sys_gapWidth := int(gapWidth)
}

func Fn_gtk_paint_slider(style *Style, cr *cairo.Context, stateType int, shadowType int, widget *Widget, detail string, x int, y int, width int, height int, orientation int) {
	sys_style := unsafe.Pointer(style)
	sys_cr := unsafe.Pointer(cr)
	sys_stateType := int(stateType)
	sys_shadowType := int(shadowType)
	sys_widget := unsafe.Pointer(widget)
	sys_detail := string(detail)
	sys_x := int(x)
	sys_y := int(y)
	sys_width := int(width)
	sys_height := int(height)
	sys_orientation := int(orientation)
}

func Fn_gtk_paint_spinner(style *Style, cr *cairo.Context, stateType int, widget *Widget, detail string, step uint, x int, y int, width int, height int) {
	sys_style := unsafe.Pointer(style)
	sys_cr := unsafe.Pointer(cr)
	sys_stateType := int(stateType)
	sys_widget := unsafe.Pointer(widget)
	sys_detail := string(detail)
	sys_step := uint(step)
	sys_x := int(x)
	sys_y := int(y)
	sys_width := int(width)
	sys_height := int(height)
}

func Fn_gtk_paint_tab(style *Style, cr *cairo.Context, stateType int, shadowType int, widget *Widget, detail string, x int, y int, width int, height int) {
	sys_style := unsafe.Pointer(style)
	sys_cr := unsafe.Pointer(cr)
	sys_stateType := int(stateType)
	sys_shadowType := int(shadowType)
	sys_widget := unsafe.Pointer(widget)
	sys_detail := string(detail)
	sys_x := int(x)
	sys_y := int(y)
	sys_width := int(width)
	sys_height := int(height)
}

func Fn_gtk_paint_vline(style *Style, cr *cairo.Context, stateType int, widget *Widget, detail string, y1 int, y2 int, x int) {
	sys_style := unsafe.Pointer(style)
	sys_cr := unsafe.Pointer(cr)
	sys_stateType := int(stateType)
	sys_widget := unsafe.Pointer(widget)
	sys_detail := string(detail)
	sys_y1 := int(y1)
	sys_y2 := int(y2)
	sys_x := int(x)
}

// UNSUPPORTED : gtk_parse_args : has array param, argv

func Fn_gtk_print_run_page_setup_dialog(parent *Window, pageSetup *PageSetup, settings *PrintSettings) {
	sys_parent := unsafe.Pointer(parent)
	sys_pageSetup := unsafe.Pointer(pageSetup)
	sys_settings := unsafe.Pointer(settings)
}

// UNSUPPORTED : gtk_print_run_page_setup_dialog_async : parameter 'done_cb' is callback

func Fn_gtk_propagate_event(widget *Widget, event *gdk.Event) {
	sys_widget := unsafe.Pointer(widget)
	sys_event := unsafe.Pointer(event)
}

func Fn_gtk_rc_add_default_file(filename string) {
	sys_filename := string(filename)
}

func Fn_gtk_rc_find_module_in_path(moduleFile string) {
	sys_moduleFile := string(moduleFile)
}

func Fn_gtk_rc_find_pixmap_in_path(settings *Settings, scanner *glib.Scanner, pixmapFile string) {
	sys_settings := unsafe.Pointer(settings)
	sys_scanner := unsafe.Pointer(scanner)
	sys_pixmapFile := string(pixmapFile)
}

// UNSUPPORTED : gtk_rc_get_default_files : no array length

func Fn_gtk_rc_get_im_module_file() {}

func Fn_gtk_rc_get_im_module_path() {}

func Fn_gtk_rc_get_module_dir() {}

func Fn_gtk_rc_get_style(widget *Widget) {
	sys_widget := unsafe.Pointer(widget)
}

func Fn_gtk_rc_get_style_by_paths(settings *Settings, widgetPath string, classPath string, type_ uint64) {
	sys_settings := unsafe.Pointer(settings)
	sys_widgetPath := string(widgetPath)
	sys_classPath := string(classPath)
	sys_type_ := uint64(type_)
}

func Fn_gtk_rc_get_theme_dir() {}

func Fn_gtk_rc_parse(filename string) {
	sys_filename := string(filename)
}

func Fn_gtk_rc_parse_color(scanner *glib.Scanner) {
	sys_scanner := unsafe.Pointer(scanner)
}

func Fn_gtk_rc_parse_color_full(scanner *glib.Scanner, style *RcStyle) {
	sys_scanner := unsafe.Pointer(scanner)
	sys_style := unsafe.Pointer(style)
}

func Fn_gtk_rc_parse_priority(scanner *glib.Scanner, priority *int) {
	sys_scanner := unsafe.Pointer(scanner)
	sys_priority := *int(priority)
}

func Fn_gtk_rc_parse_state(scanner *glib.Scanner) {
	sys_scanner := unsafe.Pointer(scanner)
}

func Fn_gtk_rc_parse_string(rcString string) {
	sys_rcString := string(rcString)
}

func Fn_gtk_rc_reparse_all() {}

func Fn_gtk_rc_reparse_all_for_settings(settings *Settings, forceLoad bool) {
	sys_settings := unsafe.Pointer(settings)
	sys_forceLoad := bool(forceLoad)
}

func Fn_gtk_rc_reset_styles(settings *Settings) {
	sys_settings := unsafe.Pointer(settings)
}

func Fn_gtk_rc_scanner_new() {}

// UNSUPPORTED : gtk_rc_set_default_files : parameter 'filenames' is array parameter without length parameter

func Fn_gtk_rgb_to_hsv(r float64, g float64, b float64) {
	sys_r := float64(r)
	sys_g := float64(g)
	sys_b := float64(b)
}

func Fn_gtk_selection_add_target(widget *Widget, selection gdk.Atom, target gdk.Atom, info uint) {
	sys_widget := unsafe.Pointer(widget)
	sys_selection := gdk1.Atom(selection)
	sys_target := gdk1.Atom(target)
	sys_info := uint(info)
}

// UNSUPPORTED : gtk_selection_add_targets : has array param, targets

func Fn_gtk_selection_clear_targets(widget *Widget, selection gdk.Atom) {
	sys_widget := unsafe.Pointer(widget)
	sys_selection := gdk1.Atom(selection)
}

func Fn_gtk_selection_convert(widget *Widget, selection gdk.Atom, target gdk.Atom, time uint32) {
	sys_widget := unsafe.Pointer(widget)
	sys_selection := gdk1.Atom(selection)
	sys_target := gdk1.Atom(target)
	sys_time := uint32(time)
}

func Fn_gtk_selection_owner_set(widget *Widget, selection gdk.Atom, time uint32) {
	sys_widget := unsafe.Pointer(widget)
	sys_selection := gdk1.Atom(selection)
	sys_time := uint32(time)
}

func Fn_gtk_selection_owner_set_for_display(display *gdk.Display, widget *Widget, selection gdk.Atom, time uint32) {
	sys_display := unsafe.Pointer(display)
	sys_widget := unsafe.Pointer(widget)
	sys_selection := gdk1.Atom(selection)
	sys_time := uint32(time)
}

func Fn_gtk_selection_remove_all(widget *Widget) {
	sys_widget := unsafe.Pointer(widget)
}

func Fn_gtk_set_debug_flags(flags uint) {
	sys_flags := uint(flags)
}

func Fn_gtk_show_about_dialog(parent *Window, firstPropertyName string) {
	sys_parent := unsafe.Pointer(parent)
	sys_firstPropertyName := string(firstPropertyName)
}

func Fn_gtk_show_uri(screen *gdk.Screen, uri string, timestamp uint32) {
	sys_screen := unsafe.Pointer(screen)
	sys_uri := string(uri)
	sys_timestamp := uint32(timestamp)
}

// UNSUPPORTED : gtk_stock_add : has array param, items

// UNSUPPORTED : gtk_stock_add_static : has array param, items

func Fn_gtk_stock_list_ids() {}

func Fn_gtk_stock_lookup(stockId string) {
	sys_stockId := string(stockId)
}

// UNSUPPORTED : gtk_stock_set_translate_func : parameter 'func' is callback

// UNSUPPORTED : gtk_target_table_free : has array param, targets

func Fn_gtk_target_table_new_from_list(list *TargetList) {
	sys_list := unsafe.Pointer(list)
}

// UNSUPPORTED : gtk_targets_include_image : has array param, targets

// UNSUPPORTED : gtk_targets_include_rich_text : has array param, targets

// UNSUPPORTED : gtk_targets_include_text : has array param, targets

// UNSUPPORTED : gtk_targets_include_uri : has array param, targets

func Fn_gtk_test_create_simple_window(windowTitle string, dialogText string) {
	sys_windowTitle := string(windowTitle)
	sys_dialogText := string(dialogText)
}

func Fn_gtk_test_create_widget(widgetType uint64, firstPropertyName string) {
	sys_widgetType := uint64(widgetType)
	sys_firstPropertyName := string(firstPropertyName)
}

func Fn_gtk_test_display_button_window(windowTitle string, dialogText string) {
	sys_windowTitle := string(windowTitle)
	sys_dialogText := string(dialogText)
}

func Fn_gtk_test_find_label(widget *Widget, labelPattern string) {
	sys_widget := unsafe.Pointer(widget)
	sys_labelPattern := string(labelPattern)
}

func Fn_gtk_test_find_sibling(baseWidget *Widget, widgetType uint64) {
	sys_baseWidget := unsafe.Pointer(baseWidget)
	sys_widgetType := uint64(widgetType)
}

func Fn_gtk_test_find_widget(widget *Widget, labelPattern string, widgetType uint64) {
	sys_widget := unsafe.Pointer(widget)
	sys_labelPattern := string(labelPattern)
	sys_widgetType := uint64(widgetType)
}

// UNSUPPORTED : gtk_test_init : has array param, argvp

func Fn_gtk_test_list_all_types() {}

func Fn_gtk_test_register_all_types() {}

func Fn_gtk_test_slider_get_value(widget *Widget) {
	sys_widget := unsafe.Pointer(widget)
}

func Fn_gtk_test_slider_set_perc(widget *Widget, percentage float64) {
	sys_widget := unsafe.Pointer(widget)
	sys_percentage := float64(percentage)
}

func Fn_gtk_test_spin_button_click(spinner *SpinButton, button uint, upwards bool) {
	sys_spinner := unsafe.Pointer(spinner)
	sys_button := uint(button)
	sys_upwards := bool(upwards)
}

func Fn_gtk_test_text_get(widget *Widget) {
	sys_widget := unsafe.Pointer(widget)
}

func Fn_gtk_test_text_set(widget *Widget, string_ string) {
	sys_widget := unsafe.Pointer(widget)
	sys_string_ := string(string_)
}

func Fn_gtk_test_widget_click(widget *Widget, button uint, modifiers int) {
	sys_widget := unsafe.Pointer(widget)
	sys_button := uint(button)
	sys_modifiers := int(modifiers)
}

func Fn_gtk_test_widget_send_key(widget *Widget, keyval uint, modifiers int) {
	sys_widget := unsafe.Pointer(widget)
	sys_keyval := uint(keyval)
	sys_modifiers := int(modifiers)
}

func Fn_gtk_tree_get_row_drag_data(selectionData *SelectionData) {
	sys_selectionData := unsafe.Pointer(selectionData)
}

// UNSUPPORTED : gtk_tree_row_reference_reordered : parameter 'new_order' is array parameter without length parameter

func Fn_gtk_tree_set_row_drag_data(selectionData *SelectionData, treeModel *TreeModel, path *TreePath) {
	sys_selectionData := unsafe.Pointer(selectionData)
	sys_treeModel := unsafe.Pointer(treeModel)
	sys_path := unsafe.Pointer(path)
}

func Fn_gtk_true() {}

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

// ActivatableIface is a representation of the C record GtkActivatableIface.
//
// since 2.16
type ActivatableIface struct {
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

// RecentInfo is a representation of the C record GtkRecentInfo.
//
// since 2.10
type RecentInfo struct {
	native unsafe.Pointer
}

// RecentManagerClass is a representation of the C record GtkRecentManagerClass.
//
// since 2.10
type RecentManagerClass struct {
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

// AboutDialog is a representation of the C record GtkAboutDialog.
type AboutDialog struct {
	native unsafe.Pointer
}

// AccelGroup is a representation of the C record GtkAccelGroup.
type AccelGroup struct {
	native unsafe.Pointer
}

// AccelLabel is a representation of the C record GtkAccelLabel.
type AccelLabel struct {
	native unsafe.Pointer
}

// AccelMap is a representation of the C record GtkAccelMap.
type AccelMap struct {
	native unsafe.Pointer
}

// Accessible is a representation of the C record GtkAccessible.
type Accessible struct {
	native unsafe.Pointer
}

// Action is a representation of the C record GtkAction.
type Action struct {
	native unsafe.Pointer
}

// ActionBar is a representation of the C record GtkActionBar.
type ActionBar struct {
	native unsafe.Pointer
}

// ActionGroup is a representation of the C record GtkActionGroup.
type ActionGroup struct {
	native unsafe.Pointer
}

// Adjustment is a representation of the C record GtkAdjustment.
type Adjustment struct {
	native unsafe.Pointer
}

// Alignment is a representation of the C record GtkAlignment.
type Alignment struct {
	native unsafe.Pointer
}

// AppChooserButton is a representation of the C record GtkAppChooserButton.
type AppChooserButton struct {
	native unsafe.Pointer
}

// AppChooserDialog is a representation of the C record GtkAppChooserDialog.
type AppChooserDialog struct {
	native unsafe.Pointer
}

// AppChooserWidget is a representation of the C record GtkAppChooserWidget.
type AppChooserWidget struct {
	native unsafe.Pointer
}

// Application is a representation of the C record GtkApplication.
type Application struct {
	native unsafe.Pointer
}

// ApplicationWindow is a representation of the C record GtkApplicationWindow.
type ApplicationWindow struct {
	native unsafe.Pointer
}

// Arrow is a representation of the C record GtkArrow.
type Arrow struct {
	native unsafe.Pointer
}

// ArrowAccessible is a representation of the C record GtkArrowAccessible.
type ArrowAccessible struct {
	native unsafe.Pointer
}

// AspectFrame is a representation of the C record GtkAspectFrame.
type AspectFrame struct {
	native unsafe.Pointer
}

// Assistant is a representation of the C record GtkAssistant.
type Assistant struct {
	native unsafe.Pointer
}

// Bin is a representation of the C record GtkBin.
type Bin struct {
	native unsafe.Pointer
}

// BooleanCellAccessible is a representation of the C record GtkBooleanCellAccessible.
type BooleanCellAccessible struct {
	native unsafe.Pointer
}

// Box is a representation of the C record GtkBox.
type Box struct {
	native unsafe.Pointer
}

// Builder is a representation of the C record GtkBuilder.
type Builder struct {
	native unsafe.Pointer
}

// Button is a representation of the C record GtkButton.
type Button struct {
	native unsafe.Pointer
}

// ButtonAccessible is a representation of the C record GtkButtonAccessible.
type ButtonAccessible struct {
	native unsafe.Pointer
}

// ButtonBox is a representation of the C record GtkButtonBox.
type ButtonBox struct {
	native unsafe.Pointer
}

// Calendar is a representation of the C record GtkCalendar.
type Calendar struct {
	native unsafe.Pointer
}

// CellAccessible is a representation of the C record GtkCellAccessible.
type CellAccessible struct {
	native unsafe.Pointer
}

// CellArea is a representation of the C record GtkCellArea.
type CellArea struct {
	native unsafe.Pointer
}

// CellAreaBox is a representation of the C record GtkCellAreaBox.
type CellAreaBox struct {
	native unsafe.Pointer
}

// CellAreaContext is a representation of the C record GtkCellAreaContext.
type CellAreaContext struct {
	native unsafe.Pointer
}

// CellRenderer is a representation of the C record GtkCellRenderer.
type CellRenderer struct {
	native unsafe.Pointer
}

// CellRendererAccel is a representation of the C record GtkCellRendererAccel.
type CellRendererAccel struct {
	native unsafe.Pointer
}

// CellRendererCombo is a representation of the C record GtkCellRendererCombo.
type CellRendererCombo struct {
	native unsafe.Pointer
}

// CellRendererPixbuf is a representation of the C record GtkCellRendererPixbuf.
type CellRendererPixbuf struct {
	native unsafe.Pointer
}

// CellRendererProgress is a representation of the C record GtkCellRendererProgress.
type CellRendererProgress struct {
	native unsafe.Pointer
}

// CellRendererSpin is a representation of the C record GtkCellRendererSpin.
type CellRendererSpin struct {
	native unsafe.Pointer
}

// CellRendererSpinner is a representation of the C record GtkCellRendererSpinner.
type CellRendererSpinner struct {
	native unsafe.Pointer
}

// CellRendererText is a representation of the C record GtkCellRendererText.
type CellRendererText struct {
	native unsafe.Pointer
}

// CellRendererToggle is a representation of the C record GtkCellRendererToggle.
type CellRendererToggle struct {
	native unsafe.Pointer
}

// CellView is a representation of the C record GtkCellView.
type CellView struct {
	native unsafe.Pointer
}

// CheckButton is a representation of the C record GtkCheckButton.
type CheckButton struct {
	native unsafe.Pointer
}

// CheckMenuItem is a representation of the C record GtkCheckMenuItem.
type CheckMenuItem struct {
	native unsafe.Pointer
}

// CheckMenuItemAccessible is a representation of the C record GtkCheckMenuItemAccessible.
type CheckMenuItemAccessible struct {
	native unsafe.Pointer
}

// Clipboard is a representation of the C record GtkClipboard.
type Clipboard struct {
	native unsafe.Pointer
}

// ColorButton is a representation of the C record GtkColorButton.
type ColorButton struct {
	native unsafe.Pointer
}

// ColorSelection is a representation of the C record GtkColorSelection.
type ColorSelection struct {
	native unsafe.Pointer
}

// ColorSelectionDialog is a representation of the C record GtkColorSelectionDialog.
type ColorSelectionDialog struct {
	native unsafe.Pointer
}

// ComboBox is a representation of the C record GtkComboBox.
type ComboBox struct {
	native unsafe.Pointer
}

// ComboBoxAccessible is a representation of the C record GtkComboBoxAccessible.
type ComboBoxAccessible struct {
	native unsafe.Pointer
}

// ComboBoxText is a representation of the C record GtkComboBoxText.
type ComboBoxText struct {
	native unsafe.Pointer
}

// Container is a representation of the C record GtkContainer.
type Container struct {
	native unsafe.Pointer
}

// ContainerAccessible is a representation of the C record GtkContainerAccessible.
type ContainerAccessible struct {
	native unsafe.Pointer
}

// ContainerCellAccessible is a representation of the C record GtkContainerCellAccessible.
type ContainerCellAccessible struct {
	native unsafe.Pointer
}

// CssProvider is a representation of the C record GtkCssProvider.
type CssProvider struct {
	native unsafe.Pointer
}

// Dialog is a representation of the C record GtkDialog.
type Dialog struct {
	native unsafe.Pointer
}

// DrawingArea is a representation of the C record GtkDrawingArea.
type DrawingArea struct {
	native unsafe.Pointer
}

// Entry is a representation of the C record GtkEntry.
type Entry struct {
	native unsafe.Pointer
}

// EntryAccessible is a representation of the C record GtkEntryAccessible.
type EntryAccessible struct {
	native unsafe.Pointer
}

// EntryBuffer is a representation of the C record GtkEntryBuffer.
//
// since 2.18
type EntryBuffer struct {
	native unsafe.Pointer
}

// EntryCompletion is a representation of the C record GtkEntryCompletion.
type EntryCompletion struct {
	native unsafe.Pointer
}

// UNSUPPORTED : EntryIconAccessible : blacklisted

// EventBox is a representation of the C record GtkEventBox.
type EventBox struct {
	native unsafe.Pointer
}

// EventController is a representation of the C record GtkEventController.
type EventController struct {
	native unsafe.Pointer
}

// UNSUPPORTED : EventControllerKey : blacklisted

// UNSUPPORTED : EventControllerMotion : blacklisted

// UNSUPPORTED : EventControllerScroll : blacklisted

// Expander is a representation of the C record GtkExpander.
type Expander struct {
	native unsafe.Pointer
}

// ExpanderAccessible is a representation of the C record GtkExpanderAccessible.
type ExpanderAccessible struct {
	native unsafe.Pointer
}

// FileChooserButton is a representation of the C record GtkFileChooserButton.
type FileChooserButton struct {
	native unsafe.Pointer
}

// FileChooserDialog is a representation of the C record GtkFileChooserDialog.
type FileChooserDialog struct {
	native unsafe.Pointer
}

// FileChooserNative is a representation of the C record GtkFileChooserNative.
type FileChooserNative struct {
	native unsafe.Pointer
}

// FileChooserWidget is a representation of the C record GtkFileChooserWidget.
type FileChooserWidget struct {
	native unsafe.Pointer
}

// FileFilter is a representation of the C record GtkFileFilter.
type FileFilter struct {
	native unsafe.Pointer
}

// Fixed is a representation of the C record GtkFixed.
type Fixed struct {
	native unsafe.Pointer
}

// FlowBox is a representation of the C record GtkFlowBox.
type FlowBox struct {
	native unsafe.Pointer
}

// FlowBoxAccessible is a representation of the C record GtkFlowBoxAccessible.
type FlowBoxAccessible struct {
	native unsafe.Pointer
}

// FlowBoxChild is a representation of the C record GtkFlowBoxChild.
type FlowBoxChild struct {
	native unsafe.Pointer
}

// FlowBoxChildAccessible is a representation of the C record GtkFlowBoxChildAccessible.
type FlowBoxChildAccessible struct {
	native unsafe.Pointer
}

// FontButton is a representation of the C record GtkFontButton.
type FontButton struct {
	native unsafe.Pointer
}

// FontSelection is a representation of the C record GtkFontSelection.
type FontSelection struct {
	native unsafe.Pointer
}

// FontSelectionDialog is a representation of the C record GtkFontSelectionDialog.
type FontSelectionDialog struct {
	native unsafe.Pointer
}

// Frame is a representation of the C record GtkFrame.
type Frame struct {
	native unsafe.Pointer
}

// FrameAccessible is a representation of the C record GtkFrameAccessible.
type FrameAccessible struct {
	native unsafe.Pointer
}

// Gesture is a representation of the C record GtkGesture.
type Gesture struct {
	native unsafe.Pointer
}

// GestureDrag is a representation of the C record GtkGestureDrag.
type GestureDrag struct {
	native unsafe.Pointer
}

// GestureLongPress is a representation of the C record GtkGestureLongPress.
type GestureLongPress struct {
	native unsafe.Pointer
}

// GestureMultiPress is a representation of the C record GtkGestureMultiPress.
type GestureMultiPress struct {
	native unsafe.Pointer
}

// GesturePan is a representation of the C record GtkGesturePan.
type GesturePan struct {
	native unsafe.Pointer
}

// GestureRotate is a representation of the C record GtkGestureRotate.
type GestureRotate struct {
	native unsafe.Pointer
}

// GestureSingle is a representation of the C record GtkGestureSingle.
type GestureSingle struct {
	native unsafe.Pointer
}

// UNSUPPORTED : GestureStylus : blacklisted

// GestureSwipe is a representation of the C record GtkGestureSwipe.
type GestureSwipe struct {
	native unsafe.Pointer
}

// GestureZoom is a representation of the C record GtkGestureZoom.
type GestureZoom struct {
	native unsafe.Pointer
}

// Grid is a representation of the C record GtkGrid.
type Grid struct {
	native unsafe.Pointer
}

// HBox is a representation of the C record GtkHBox.
type HBox struct {
	native unsafe.Pointer
}

// HButtonBox is a representation of the C record GtkHButtonBox.
type HButtonBox struct {
	native unsafe.Pointer
}

// HPaned is a representation of the C record GtkHPaned.
type HPaned struct {
	native unsafe.Pointer
}

// HSV is a representation of the C record GtkHSV.
type HSV struct {
	native unsafe.Pointer
}

// HScale is a representation of the C record GtkHScale.
type HScale struct {
	native unsafe.Pointer
}

// HScrollbar is a representation of the C record GtkHScrollbar.
type HScrollbar struct {
	native unsafe.Pointer
}

// HSeparator is a representation of the C record GtkHSeparator.
type HSeparator struct {
	native unsafe.Pointer
}

// HandleBox is a representation of the C record GtkHandleBox.
type HandleBox struct {
	native unsafe.Pointer
}

// HeaderBar is a representation of the C record GtkHeaderBar.
type HeaderBar struct {
	native unsafe.Pointer
}

// UNSUPPORTED : HeaderBarAccessible : blacklisted

// IMContext is a representation of the C record GtkIMContext.
type IMContext struct {
	native unsafe.Pointer
}

// IMContextSimple is a representation of the C record GtkIMContextSimple.
type IMContextSimple struct {
	native unsafe.Pointer
}

// IMMulticontext is a representation of the C record GtkIMMulticontext.
type IMMulticontext struct {
	native unsafe.Pointer
}

// IconFactory is a representation of the C record GtkIconFactory.
type IconFactory struct {
	native unsafe.Pointer
}

// IconInfo is a representation of the C record GtkIconInfo.
type IconInfo struct {
	native unsafe.Pointer
}

// IconTheme is a representation of the C record GtkIconTheme.
type IconTheme struct {
	native unsafe.Pointer
}

// IconView is a representation of the C record GtkIconView.
type IconView struct {
	native unsafe.Pointer
}

// IconViewAccessible is a representation of the C record GtkIconViewAccessible.
type IconViewAccessible struct {
	native unsafe.Pointer
}

// Image is a representation of the C record GtkImage.
type Image struct {
	native unsafe.Pointer
}

// ImageAccessible is a representation of the C record GtkImageAccessible.
type ImageAccessible struct {
	native unsafe.Pointer
}

// ImageCellAccessible is a representation of the C record GtkImageCellAccessible.
type ImageCellAccessible struct {
	native unsafe.Pointer
}

// ImageMenuItem is a representation of the C record GtkImageMenuItem.
type ImageMenuItem struct {
	native unsafe.Pointer
}

// InfoBar is a representation of the C record GtkInfoBar.
type InfoBar struct {
	native unsafe.Pointer
}

// Invisible is a representation of the C record GtkInvisible.
type Invisible struct {
	native unsafe.Pointer
}

// Label is a representation of the C record GtkLabel.
type Label struct {
	native unsafe.Pointer
}

// LabelAccessible is a representation of the C record GtkLabelAccessible.
type LabelAccessible struct {
	native unsafe.Pointer
}

// Layout is a representation of the C record GtkLayout.
type Layout struct {
	native unsafe.Pointer
}

// LevelBar is a representation of the C record GtkLevelBar.
type LevelBar struct {
	native unsafe.Pointer
}

// LevelBarAccessible is a representation of the C record GtkLevelBarAccessible.
type LevelBarAccessible struct {
	native unsafe.Pointer
}

// LinkButton is a representation of the C record GtkLinkButton.
type LinkButton struct {
	native unsafe.Pointer
}

// LinkButtonAccessible is a representation of the C record GtkLinkButtonAccessible.
type LinkButtonAccessible struct {
	native unsafe.Pointer
}

// ListBox is a representation of the C record GtkListBox.
type ListBox struct {
	native unsafe.Pointer
}

// ListBoxAccessible is a representation of the C record GtkListBoxAccessible.
type ListBoxAccessible struct {
	native unsafe.Pointer
}

// ListBoxRow is a representation of the C record GtkListBoxRow.
type ListBoxRow struct {
	native unsafe.Pointer
}

// ListBoxRowAccessible is a representation of the C record GtkListBoxRowAccessible.
type ListBoxRowAccessible struct {
	native unsafe.Pointer
}

// ListStore is a representation of the C record GtkListStore.
type ListStore struct {
	native unsafe.Pointer
}

// LockButton is a representation of the C record GtkLockButton.
type LockButton struct {
	native unsafe.Pointer
}

// LockButtonAccessible is a representation of the C record GtkLockButtonAccessible.
type LockButtonAccessible struct {
	native unsafe.Pointer
}

// Menu is a representation of the C record GtkMenu.
type Menu struct {
	native unsafe.Pointer
}

// MenuAccessible is a representation of the C record GtkMenuAccessible.
type MenuAccessible struct {
	native unsafe.Pointer
}

// MenuBar is a representation of the C record GtkMenuBar.
type MenuBar struct {
	native unsafe.Pointer
}

// MenuButton is a representation of the C record GtkMenuButton.
type MenuButton struct {
	native unsafe.Pointer
}

// MenuButtonAccessible is a representation of the C record GtkMenuButtonAccessible.
type MenuButtonAccessible struct {
	native unsafe.Pointer
}

// MenuItem is a representation of the C record GtkMenuItem.
type MenuItem struct {
	native unsafe.Pointer
}

// MenuItemAccessible is a representation of the C record GtkMenuItemAccessible.
type MenuItemAccessible struct {
	native unsafe.Pointer
}

// MenuShell is a representation of the C record GtkMenuShell.
type MenuShell struct {
	native unsafe.Pointer
}

// MenuShellAccessible is a representation of the C record GtkMenuShellAccessible.
type MenuShellAccessible struct {
	native unsafe.Pointer
}

// MenuToolButton is a representation of the C record GtkMenuToolButton.
type MenuToolButton struct {
	native unsafe.Pointer
}

// MessageDialog is a representation of the C record GtkMessageDialog.
type MessageDialog struct {
	native unsafe.Pointer
}

// Misc is a representation of the C record GtkMisc.
type Misc struct {
	native unsafe.Pointer
}

// ModelButton is a representation of the C record GtkModelButton.
type ModelButton struct {
	native unsafe.Pointer
}

// MountOperation is a representation of the C record GtkMountOperation.
type MountOperation struct {
	native unsafe.Pointer
}

// NativeDialog is a representation of the C record GtkNativeDialog.
type NativeDialog struct {
	native unsafe.Pointer
}

// Notebook is a representation of the C record GtkNotebook.
type Notebook struct {
	native unsafe.Pointer
}

// NotebookAccessible is a representation of the C record GtkNotebookAccessible.
type NotebookAccessible struct {
	native unsafe.Pointer
}

// NotebookPageAccessible is a representation of the C record GtkNotebookPageAccessible.
type NotebookPageAccessible struct {
	native unsafe.Pointer
}

// NumerableIcon is a representation of the C record GtkNumerableIcon.
type NumerableIcon struct {
	native unsafe.Pointer
}

// OffscreenWindow is a representation of the C record GtkOffscreenWindow.
type OffscreenWindow struct {
	native unsafe.Pointer
}

// Overlay is a representation of the C record GtkOverlay.
type Overlay struct {
	native unsafe.Pointer
}

// PadController is a representation of the C record GtkPadController.
type PadController struct {
	native unsafe.Pointer
}

// PageSetup is a representation of the C record GtkPageSetup.
type PageSetup struct {
	native unsafe.Pointer
}

// Paned is a representation of the C record GtkPaned.
type Paned struct {
	native unsafe.Pointer
}

// PanedAccessible is a representation of the C record GtkPanedAccessible.
type PanedAccessible struct {
	native unsafe.Pointer
}

// PlacesSidebar is a representation of the C record GtkPlacesSidebar.
type PlacesSidebar struct {
	native unsafe.Pointer
}

// Plug is a representation of the C record GtkPlug.
type Plug struct {
	native unsafe.Pointer
}

// PopoverAccessible is a representation of the C record GtkPopoverAccessible.
type PopoverAccessible struct {
	native unsafe.Pointer
}

// PopoverMenu is a representation of the C record GtkPopoverMenu.
type PopoverMenu struct {
	native unsafe.Pointer
}

// PrintContext is a representation of the C record GtkPrintContext.
type PrintContext struct {
	native unsafe.Pointer
}

// PrintOperation is a representation of the C record GtkPrintOperation.
type PrintOperation struct {
	native unsafe.Pointer
}

// PrintSettings is a representation of the C record GtkPrintSettings.
type PrintSettings struct {
	native unsafe.Pointer
}

// ProgressBar is a representation of the C record GtkProgressBar.
type ProgressBar struct {
	native unsafe.Pointer
}

// ProgressBarAccessible is a representation of the C record GtkProgressBarAccessible.
type ProgressBarAccessible struct {
	native unsafe.Pointer
}

// RadioAction is a representation of the C record GtkRadioAction.
type RadioAction struct {
	native unsafe.Pointer
}

// RadioButton is a representation of the C record GtkRadioButton.
type RadioButton struct {
	native unsafe.Pointer
}

// RadioButtonAccessible is a representation of the C record GtkRadioButtonAccessible.
type RadioButtonAccessible struct {
	native unsafe.Pointer
}

// RadioMenuItem is a representation of the C record GtkRadioMenuItem.
type RadioMenuItem struct {
	native unsafe.Pointer
}

// RadioMenuItemAccessible is a representation of the C record GtkRadioMenuItemAccessible.
type RadioMenuItemAccessible struct {
	native unsafe.Pointer
}

// RadioToolButton is a representation of the C record GtkRadioToolButton.
type RadioToolButton struct {
	native unsafe.Pointer
}

// Range is a representation of the C record GtkRange.
type Range struct {
	native unsafe.Pointer
}

// RangeAccessible is a representation of the C record GtkRangeAccessible.
type RangeAccessible struct {
	native unsafe.Pointer
}

// RcStyle is a representation of the C record GtkRcStyle.
type RcStyle struct {
	native unsafe.Pointer
}

// RecentAction is a representation of the C record GtkRecentAction.
type RecentAction struct {
	native unsafe.Pointer
}

// RecentChooserDialog is a representation of the C record GtkRecentChooserDialog.
type RecentChooserDialog struct {
	native unsafe.Pointer
}

// RecentChooserMenu is a representation of the C record GtkRecentChooserMenu.
type RecentChooserMenu struct {
	native unsafe.Pointer
}

// RecentChooserWidget is a representation of the C record GtkRecentChooserWidget.
type RecentChooserWidget struct {
	native unsafe.Pointer
}

// RecentFilter is a representation of the C record GtkRecentFilter.
type RecentFilter struct {
	native unsafe.Pointer
}

// RecentManager is a representation of the C record GtkRecentManager.
//
// since 2.10
type RecentManager struct {
	native unsafe.Pointer
}

// RendererCellAccessible is a representation of the C record GtkRendererCellAccessible.
type RendererCellAccessible struct {
	native unsafe.Pointer
}

// Revealer is a representation of the C record GtkRevealer.
type Revealer struct {
	native unsafe.Pointer
}

// Scale is a representation of the C record GtkScale.
type Scale struct {
	native unsafe.Pointer
}

// ScaleAccessible is a representation of the C record GtkScaleAccessible.
type ScaleAccessible struct {
	native unsafe.Pointer
}

// ScaleButton is a representation of the C record GtkScaleButton.
type ScaleButton struct {
	native unsafe.Pointer
}

// ScaleButtonAccessible is a representation of the C record GtkScaleButtonAccessible.
type ScaleButtonAccessible struct {
	native unsafe.Pointer
}

// Scrollbar is a representation of the C record GtkScrollbar.
type Scrollbar struct {
	native unsafe.Pointer
}

// ScrolledWindow is a representation of the C record GtkScrolledWindow.
type ScrolledWindow struct {
	native unsafe.Pointer
}

// ScrolledWindowAccessible is a representation of the C record GtkScrolledWindowAccessible.
type ScrolledWindowAccessible struct {
	native unsafe.Pointer
}

// Separator is a representation of the C record GtkSeparator.
type Separator struct {
	native unsafe.Pointer
}

// SeparatorMenuItem is a representation of the C record GtkSeparatorMenuItem.
type SeparatorMenuItem struct {
	native unsafe.Pointer
}

// SeparatorToolItem is a representation of the C record GtkSeparatorToolItem.
type SeparatorToolItem struct {
	native unsafe.Pointer
}

// Settings is a representation of the C record GtkSettings.
type Settings struct {
	native unsafe.Pointer
}

// ShortcutLabel is a representation of the C record GtkShortcutLabel.
type ShortcutLabel struct {
	native unsafe.Pointer
}

// ShortcutsGroup is a representation of the C record GtkShortcutsGroup.
type ShortcutsGroup struct {
	native unsafe.Pointer
}

// ShortcutsSection is a representation of the C record GtkShortcutsSection.
type ShortcutsSection struct {
	native unsafe.Pointer
}

// ShortcutsShortcut is a representation of the C record GtkShortcutsShortcut.
type ShortcutsShortcut struct {
	native unsafe.Pointer
}

// ShortcutsWindow is a representation of the C record GtkShortcutsWindow.
type ShortcutsWindow struct {
	native unsafe.Pointer
}

// SizeGroup is a representation of the C record GtkSizeGroup.
type SizeGroup struct {
	native unsafe.Pointer
}

// Socket is a representation of the C record GtkSocket.
type Socket struct {
	native unsafe.Pointer
}

// SpinButton is a representation of the C record GtkSpinButton.
type SpinButton struct {
	native unsafe.Pointer
}

// SpinButtonAccessible is a representation of the C record GtkSpinButtonAccessible.
type SpinButtonAccessible struct {
	native unsafe.Pointer
}

// Spinner is a representation of the C record GtkSpinner.
type Spinner struct {
	native unsafe.Pointer
}

// SpinnerAccessible is a representation of the C record GtkSpinnerAccessible.
type SpinnerAccessible struct {
	native unsafe.Pointer
}

// Stack is a representation of the C record GtkStack.
type Stack struct {
	native unsafe.Pointer
}

// UNSUPPORTED : StackAccessible : blacklisted

// StackSwitcher is a representation of the C record GtkStackSwitcher.
type StackSwitcher struct {
	native unsafe.Pointer
}

// StatusIcon is a representation of the C record GtkStatusIcon.
type StatusIcon struct {
	native unsafe.Pointer
}

// Statusbar is a representation of the C record GtkStatusbar.
type Statusbar struct {
	native unsafe.Pointer
}

// StatusbarAccessible is a representation of the C record GtkStatusbarAccessible.
type StatusbarAccessible struct {
	native unsafe.Pointer
}

// Style is a representation of the C record GtkStyle.
type Style struct {
	native unsafe.Pointer
}

// StyleContext is a representation of the C record GtkStyleContext.
type StyleContext struct {
	native unsafe.Pointer
}

// StyleProperties is a representation of the C record GtkStyleProperties.
type StyleProperties struct {
	native unsafe.Pointer
}

// Switch is a representation of the C record GtkSwitch.
type Switch struct {
	native unsafe.Pointer
}

// SwitchAccessible is a representation of the C record GtkSwitchAccessible.
type SwitchAccessible struct {
	native unsafe.Pointer
}

// Table is a representation of the C record GtkTable.
type Table struct {
	native unsafe.Pointer
}

// TearoffMenuItem is a representation of the C record GtkTearoffMenuItem.
type TearoffMenuItem struct {
	native unsafe.Pointer
}

// TextBuffer is a representation of the C record GtkTextBuffer.
type TextBuffer struct {
	native unsafe.Pointer
}

// TextCellAccessible is a representation of the C record GtkTextCellAccessible.
type TextCellAccessible struct {
	native unsafe.Pointer
}

// TextChildAnchor is a representation of the C record GtkTextChildAnchor.
type TextChildAnchor struct {
	native unsafe.Pointer
}

// TextMark is a representation of the C record GtkTextMark.
type TextMark struct {
	native unsafe.Pointer
}

// TextTag is a representation of the C record GtkTextTag.
type TextTag struct {
	native unsafe.Pointer
}

// TextTagTable is a representation of the C record GtkTextTagTable.
type TextTagTable struct {
	native unsafe.Pointer
}

// TextView is a representation of the C record GtkTextView.
type TextView struct {
	native unsafe.Pointer
}

// TextViewAccessible is a representation of the C record GtkTextViewAccessible.
type TextViewAccessible struct {
	native unsafe.Pointer
}

// ThemingEngine is a representation of the C record GtkThemingEngine.
type ThemingEngine struct {
	native unsafe.Pointer
}

// ToggleAction is a representation of the C record GtkToggleAction.
type ToggleAction struct {
	native unsafe.Pointer
}

// ToggleButton is a representation of the C record GtkToggleButton.
type ToggleButton struct {
	native unsafe.Pointer
}

// ToggleButtonAccessible is a representation of the C record GtkToggleButtonAccessible.
type ToggleButtonAccessible struct {
	native unsafe.Pointer
}

// ToggleToolButton is a representation of the C record GtkToggleToolButton.
type ToggleToolButton struct {
	native unsafe.Pointer
}

// ToolButton is a representation of the C record GtkToolButton.
type ToolButton struct {
	native unsafe.Pointer
}

// ToolItem is a representation of the C record GtkToolItem.
type ToolItem struct {
	native unsafe.Pointer
}

// ToolItemGroup is a representation of the C record GtkToolItemGroup.
//
// since 2.20
type ToolItemGroup struct {
	native unsafe.Pointer
}

// ToolPalette is a representation of the C record GtkToolPalette.
//
// since 2.20
type ToolPalette struct {
	native unsafe.Pointer
}

// Toolbar is a representation of the C record GtkToolbar.
type Toolbar struct {
	native unsafe.Pointer
}

// Tooltip is a representation of the C record GtkTooltip.
type Tooltip struct {
	native unsafe.Pointer
}

// ToplevelAccessible is a representation of the C record GtkToplevelAccessible.
type ToplevelAccessible struct {
	native unsafe.Pointer
}

// TreeModelFilter is a representation of the C record GtkTreeModelFilter.
type TreeModelFilter struct {
	native unsafe.Pointer
}

// TreeModelSort is a representation of the C record GtkTreeModelSort.
type TreeModelSort struct {
	native unsafe.Pointer
}

// TreeSelection is a representation of the C record GtkTreeSelection.
type TreeSelection struct {
	native unsafe.Pointer
}

// TreeStore is a representation of the C record GtkTreeStore.
type TreeStore struct {
	native unsafe.Pointer
}

// TreeView is a representation of the C record GtkTreeView.
type TreeView struct {
	native unsafe.Pointer
}

// TreeViewAccessible is a representation of the C record GtkTreeViewAccessible.
type TreeViewAccessible struct {
	native unsafe.Pointer
}

// TreeViewColumn is a representation of the C record GtkTreeViewColumn.
type TreeViewColumn struct {
	native unsafe.Pointer
}

// UIManager is a representation of the C record GtkUIManager.
type UIManager struct {
	native unsafe.Pointer
}

// VBox is a representation of the C record GtkVBox.
type VBox struct {
	native unsafe.Pointer
}

// VButtonBox is a representation of the C record GtkVButtonBox.
type VButtonBox struct {
	native unsafe.Pointer
}

// VPaned is a representation of the C record GtkVPaned.
type VPaned struct {
	native unsafe.Pointer
}

// VScale is a representation of the C record GtkVScale.
type VScale struct {
	native unsafe.Pointer
}

// VScrollbar is a representation of the C record GtkVScrollbar.
type VScrollbar struct {
	native unsafe.Pointer
}

// VSeparator is a representation of the C record GtkVSeparator.
type VSeparator struct {
	native unsafe.Pointer
}

// Viewport is a representation of the C record GtkViewport.
type Viewport struct {
	native unsafe.Pointer
}

// VolumeButton is a representation of the C record GtkVolumeButton.
type VolumeButton struct {
	native unsafe.Pointer
}

// Widget is a representation of the C record GtkWidget.
type Widget struct {
	native unsafe.Pointer
}

// WidgetAccessible is a representation of the C record GtkWidgetAccessible.
type WidgetAccessible struct {
	native unsafe.Pointer
}

// Window is a representation of the C record GtkWindow.
type Window struct {
	native unsafe.Pointer
}

// WindowAccessible is a representation of the C record GtkWindowAccessible.
type WindowAccessible struct {
	native unsafe.Pointer
}

// WindowGroup is a representation of the C record GtkWindowGroup.
type WindowGroup struct {
	native unsafe.Pointer
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
