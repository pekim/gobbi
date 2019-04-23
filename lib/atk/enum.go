// This is a generated file - DO NOT EDIT

package atk

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <atk/atk.h>
// #include <stdlib.h>
import "C"

// Blacklisted : AtkCoordType

// Blacklisted : AtkKeyEventType

type Layer C.AtkLayer

const (
	ATK_LAYER_INVALID    Layer = 0
	ATK_LAYER_BACKGROUND Layer = 1
	ATK_LAYER_CANVAS     Layer = 2
	ATK_LAYER_WIDGET     Layer = 3
	ATK_LAYER_MDI        Layer = 4
	ATK_LAYER_POPUP      Layer = 5
	ATK_LAYER_OVERLAY    Layer = 6
	ATK_LAYER_WINDOW     Layer = 7
)

// Blacklisted : AtkRelationType

type Role C.AtkRole

const (
	ATK_ROLE_INVALID               Role = 0
	ATK_ROLE_ACCEL_LABEL           Role = 1
	ATK_ROLE_ALERT                 Role = 2
	ATK_ROLE_ANIMATION             Role = 3
	ATK_ROLE_ARROW                 Role = 4
	ATK_ROLE_CALENDAR              Role = 5
	ATK_ROLE_CANVAS                Role = 6
	ATK_ROLE_CHECK_BOX             Role = 7
	ATK_ROLE_CHECK_MENU_ITEM       Role = 8
	ATK_ROLE_COLOR_CHOOSER         Role = 9
	ATK_ROLE_COLUMN_HEADER         Role = 10
	ATK_ROLE_COMBO_BOX             Role = 11
	ATK_ROLE_DATE_EDITOR           Role = 12
	ATK_ROLE_DESKTOP_ICON          Role = 13
	ATK_ROLE_DESKTOP_FRAME         Role = 14
	ATK_ROLE_DIAL                  Role = 15
	ATK_ROLE_DIALOG                Role = 16
	ATK_ROLE_DIRECTORY_PANE        Role = 17
	ATK_ROLE_DRAWING_AREA          Role = 18
	ATK_ROLE_FILE_CHOOSER          Role = 19
	ATK_ROLE_FILLER                Role = 20
	ATK_ROLE_FONT_CHOOSER          Role = 21
	ATK_ROLE_FRAME                 Role = 22
	ATK_ROLE_GLASS_PANE            Role = 23
	ATK_ROLE_HTML_CONTAINER        Role = 24
	ATK_ROLE_ICON                  Role = 25
	ATK_ROLE_IMAGE                 Role = 26
	ATK_ROLE_INTERNAL_FRAME        Role = 27
	ATK_ROLE_LABEL                 Role = 28
	ATK_ROLE_LAYERED_PANE          Role = 29
	ATK_ROLE_LIST                  Role = 30
	ATK_ROLE_LIST_ITEM             Role = 31
	ATK_ROLE_MENU                  Role = 32
	ATK_ROLE_MENU_BAR              Role = 33
	ATK_ROLE_MENU_ITEM             Role = 34
	ATK_ROLE_OPTION_PANE           Role = 35
	ATK_ROLE_PAGE_TAB              Role = 36
	ATK_ROLE_PAGE_TAB_LIST         Role = 37
	ATK_ROLE_PANEL                 Role = 38
	ATK_ROLE_PASSWORD_TEXT         Role = 39
	ATK_ROLE_POPUP_MENU            Role = 40
	ATK_ROLE_PROGRESS_BAR          Role = 41
	ATK_ROLE_PUSH_BUTTON           Role = 42
	ATK_ROLE_RADIO_BUTTON          Role = 43
	ATK_ROLE_RADIO_MENU_ITEM       Role = 44
	ATK_ROLE_ROOT_PANE             Role = 45
	ATK_ROLE_ROW_HEADER            Role = 46
	ATK_ROLE_SCROLL_BAR            Role = 47
	ATK_ROLE_SCROLL_PANE           Role = 48
	ATK_ROLE_SEPARATOR             Role = 49
	ATK_ROLE_SLIDER                Role = 50
	ATK_ROLE_SPLIT_PANE            Role = 51
	ATK_ROLE_SPIN_BUTTON           Role = 52
	ATK_ROLE_STATUSBAR             Role = 53
	ATK_ROLE_TABLE                 Role = 54
	ATK_ROLE_TABLE_CELL            Role = 55
	ATK_ROLE_TABLE_COLUMN_HEADER   Role = 56
	ATK_ROLE_TABLE_ROW_HEADER      Role = 57
	ATK_ROLE_TEAR_OFF_MENU_ITEM    Role = 58
	ATK_ROLE_TERMINAL              Role = 59
	ATK_ROLE_TEXT                  Role = 60
	ATK_ROLE_TOGGLE_BUTTON         Role = 61
	ATK_ROLE_TOOL_BAR              Role = 62
	ATK_ROLE_TOOL_TIP              Role = 63
	ATK_ROLE_TREE                  Role = 64
	ATK_ROLE_TREE_TABLE            Role = 65
	ATK_ROLE_UNKNOWN               Role = 66
	ATK_ROLE_VIEWPORT              Role = 67
	ATK_ROLE_WINDOW                Role = 68
	ATK_ROLE_HEADER                Role = 69
	ATK_ROLE_FOOTER                Role = 70
	ATK_ROLE_PARAGRAPH             Role = 71
	ATK_ROLE_RULER                 Role = 72
	ATK_ROLE_APPLICATION           Role = 73
	ATK_ROLE_AUTOCOMPLETE          Role = 74
	ATK_ROLE_EDITBAR               Role = 75
	ATK_ROLE_EMBEDDED              Role = 76
	ATK_ROLE_ENTRY                 Role = 77
	ATK_ROLE_CHART                 Role = 78
	ATK_ROLE_CAPTION               Role = 79
	ATK_ROLE_DOCUMENT_FRAME        Role = 80
	ATK_ROLE_HEADING               Role = 81
	ATK_ROLE_PAGE                  Role = 82
	ATK_ROLE_SECTION               Role = 83
	ATK_ROLE_REDUNDANT_OBJECT      Role = 84
	ATK_ROLE_FORM                  Role = 85
	ATK_ROLE_LINK                  Role = 86
	ATK_ROLE_INPUT_METHOD_WINDOW   Role = 87
	ATK_ROLE_TABLE_ROW             Role = 88
	ATK_ROLE_TREE_ITEM             Role = 89
	ATK_ROLE_DOCUMENT_SPREADSHEET  Role = 90
	ATK_ROLE_DOCUMENT_PRESENTATION Role = 91
	ATK_ROLE_DOCUMENT_TEXT         Role = 92
	ATK_ROLE_DOCUMENT_WEB          Role = 93
	ATK_ROLE_DOCUMENT_EMAIL        Role = 94
	ATK_ROLE_COMMENT               Role = 95
	ATK_ROLE_LIST_BOX              Role = 96
	ATK_ROLE_GROUPING              Role = 97
	ATK_ROLE_IMAGE_MAP             Role = 98
	ATK_ROLE_NOTIFICATION          Role = 99
	ATK_ROLE_INFO_BAR              Role = 100
	ATK_ROLE_LEVEL_BAR             Role = 101
	ATK_ROLE_TITLE_BAR             Role = 102
	ATK_ROLE_BLOCK_QUOTE           Role = 103
	ATK_ROLE_AUDIO                 Role = 104
	ATK_ROLE_VIDEO                 Role = 105
	ATK_ROLE_DEFINITION            Role = 106
	ATK_ROLE_ARTICLE               Role = 107
	ATK_ROLE_LANDMARK              Role = 108
	ATK_ROLE_LOG                   Role = 109
	ATK_ROLE_MARQUEE               Role = 110
	ATK_ROLE_MATH                  Role = 111
	ATK_ROLE_RATING                Role = 112
	ATK_ROLE_TIMER                 Role = 113
	ATK_ROLE_DESCRIPTION_LIST      Role = 114
	ATK_ROLE_DESCRIPTION_TERM      Role = 115
	ATK_ROLE_DESCRIPTION_VALUE     Role = 116
	ATK_ROLE_STATIC                Role = 117
	ATK_ROLE_MATH_FRACTION         Role = 118
	ATK_ROLE_MATH_ROOT             Role = 119
	ATK_ROLE_SUBSCRIPT             Role = 120
	ATK_ROLE_SUPERSCRIPT           Role = 121
	ATK_ROLE_FOOTNOTE              Role = 122
	ATK_ROLE_LAST_DEFINED          Role = 123
)

// Blacklisted : atk_role_for_name

// Blacklisted : atk_role_get_localized_name

// Blacklisted : atk_role_get_name

// Blacklisted : atk_role_register

type StateType C.AtkStateType

const (
	ATK_STATE_INVALID                 StateType = 0
	ATK_STATE_ACTIVE                  StateType = 1
	ATK_STATE_ARMED                   StateType = 2
	ATK_STATE_BUSY                    StateType = 3
	ATK_STATE_CHECKED                 StateType = 4
	ATK_STATE_DEFUNCT                 StateType = 5
	ATK_STATE_EDITABLE                StateType = 6
	ATK_STATE_ENABLED                 StateType = 7
	ATK_STATE_EXPANDABLE              StateType = 8
	ATK_STATE_EXPANDED                StateType = 9
	ATK_STATE_FOCUSABLE               StateType = 10
	ATK_STATE_FOCUSED                 StateType = 11
	ATK_STATE_HORIZONTAL              StateType = 12
	ATK_STATE_ICONIFIED               StateType = 13
	ATK_STATE_MODAL                   StateType = 14
	ATK_STATE_MULTI_LINE              StateType = 15
	ATK_STATE_MULTISELECTABLE         StateType = 16
	ATK_STATE_OPAQUE                  StateType = 17
	ATK_STATE_PRESSED                 StateType = 18
	ATK_STATE_RESIZABLE               StateType = 19
	ATK_STATE_SELECTABLE              StateType = 20
	ATK_STATE_SELECTED                StateType = 21
	ATK_STATE_SENSITIVE               StateType = 22
	ATK_STATE_SHOWING                 StateType = 23
	ATK_STATE_SINGLE_LINE             StateType = 24
	ATK_STATE_STALE                   StateType = 25
	ATK_STATE_TRANSIENT               StateType = 26
	ATK_STATE_VERTICAL                StateType = 27
	ATK_STATE_VISIBLE                 StateType = 28
	ATK_STATE_MANAGES_DESCENDANTS     StateType = 29
	ATK_STATE_INDETERMINATE           StateType = 30
	ATK_STATE_TRUNCATED               StateType = 31
	ATK_STATE_REQUIRED                StateType = 32
	ATK_STATE_INVALID_ENTRY           StateType = 33
	ATK_STATE_SUPPORTS_AUTOCOMPLETION StateType = 34
	ATK_STATE_SELECTABLE_TEXT         StateType = 35
	ATK_STATE_DEFAULT                 StateType = 36
	ATK_STATE_ANIMATED                StateType = 37
	ATK_STATE_VISITED                 StateType = 38
	ATK_STATE_CHECKABLE               StateType = 39
	ATK_STATE_HAS_POPUP               StateType = 40
	ATK_STATE_HAS_TOOLTIP             StateType = 41
	ATK_STATE_READ_ONLY               StateType = 42
	ATK_STATE_LAST_DEFINED            StateType = 43
)

// Blacklisted : atk_state_type_for_name

// Blacklisted : atk_state_type_get_name

// Blacklisted : atk_state_type_register

// Blacklisted : AtkTextAttribute

// Blacklisted : AtkTextBoundary

// Blacklisted : AtkTextClipType

// Blacklisted : AtkTextGranularity

// Blacklisted : AtkValueType
