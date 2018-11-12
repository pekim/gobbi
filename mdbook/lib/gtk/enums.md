# `gtk` enums

## `Align`

Controls how a widget deals with extra space in a single (x or y)
dimension.

Alignment only matters if the widget receives a “too large” allocation,
for example if you packed the widget with the #GtkWidget:expand
flag inside a #GtkBox, then the widget might get extra space.  If
you have for example a 16x16 icon inside a 32x32 space, the icon
could be scaled and stretched, it could be centered, or it could be
positioned to one side of the space.

Note that in horizontal context @GTK_ALIGN_START and @GTK_ALIGN_END
are interpreted relative to text direction.

GTK_ALIGN_BASELINE support for it is optional for containers and widgets, and
it is only supported for vertical alignment.  When its not supported by
a child or a container it is treated as @GTK_ALIGN_FILL.

- GTK_ALIGN_FILL = 0
- GTK_ALIGN_START = 1
- GTK_ALIGN_END = 2
- GTK_ALIGN_CENTER = 3
- GTK_ALIGN_BASELINE = 4

C - `GtkAlign`

## `ArrowPlacement`

Used to specify the placement of scroll arrows in scrolling menus.

- GTK_ARROWS_BOTH = 0
- GTK_ARROWS_START = 1
- GTK_ARROWS_END = 2

C - `GtkArrowPlacement`

## `ArrowType`

Used to indicate the direction in which an arrow should point.

- GTK_ARROW_UP = 0
- GTK_ARROW_DOWN = 1
- GTK_ARROW_LEFT = 2
- GTK_ARROW_RIGHT = 3
- GTK_ARROW_NONE = 4

C - `GtkArrowType`

## `AssistantPageType`

An enum for determining the page role inside the #GtkAssistant. It's
used to handle buttons sensitivity and visibility.

Note that an assistant needs to end its page flow with a page of type
%GTK_ASSISTANT_PAGE_CONFIRM, %GTK_ASSISTANT_PAGE_SUMMARY or
%GTK_ASSISTANT_PAGE_PROGRESS to be correct.

The Cancel button will only be shown if the page isn’t “committed”.
See gtk_assistant_commit() for details.

- GTK_ASSISTANT_PAGE_CONTENT = 0
- GTK_ASSISTANT_PAGE_INTRO = 1
- GTK_ASSISTANT_PAGE_CONFIRM = 2
- GTK_ASSISTANT_PAGE_SUMMARY = 3
- GTK_ASSISTANT_PAGE_PROGRESS = 4
- GTK_ASSISTANT_PAGE_CUSTOM = 5

C - `GtkAssistantPageType`

## `BaselinePosition`

Whenever a container has some form of natural row it may align
children in that row along a common typographical baseline. If
the amount of verical space in the row is taller than the total
requested height of the baseline-aligned children then it can use a
&num;GtkBaselinePosition to select where to put the baseline inside the
extra availible space.

- GTK_BASELINE_POSITION_TOP = 0
- GTK_BASELINE_POSITION_CENTER = 1
- GTK_BASELINE_POSITION_BOTTOM = 2

C - `GtkBaselinePosition`

## `BorderStyle`

Describes how the border of a UI element should be rendered.

- GTK_BORDER_STYLE_NONE = 0
- GTK_BORDER_STYLE_SOLID = 1
- GTK_BORDER_STYLE_INSET = 2
- GTK_BORDER_STYLE_OUTSET = 3
- GTK_BORDER_STYLE_HIDDEN = 4
- GTK_BORDER_STYLE_DOTTED = 5
- GTK_BORDER_STYLE_DASHED = 6
- GTK_BORDER_STYLE_DOUBLE = 7
- GTK_BORDER_STYLE_GROOVE = 8
- GTK_BORDER_STYLE_RIDGE = 9

C - `GtkBorderStyle`

## `BuilderError`

Error codes that identify various errors that can occur while using
&num;GtkBuilder.

- GTK_BUILDER_ERROR_INVALID_TYPE_FUNCTION = 0
- GTK_BUILDER_ERROR_UNHANDLED_TAG = 1
- GTK_BUILDER_ERROR_MISSING_ATTRIBUTE = 2
- GTK_BUILDER_ERROR_INVALID_ATTRIBUTE = 3
- GTK_BUILDER_ERROR_INVALID_TAG = 4
- GTK_BUILDER_ERROR_MISSING_PROPERTY_VALUE = 5
- GTK_BUILDER_ERROR_INVALID_VALUE = 6
- GTK_BUILDER_ERROR_VERSION_MISMATCH = 7
- GTK_BUILDER_ERROR_DUPLICATE_ID = 8
- GTK_BUILDER_ERROR_OBJECT_TYPE_REFUSED = 9
- GTK_BUILDER_ERROR_TEMPLATE_MISMATCH = 10
- GTK_BUILDER_ERROR_INVALID_PROPERTY = 11
- GTK_BUILDER_ERROR_INVALID_SIGNAL = 12
- GTK_BUILDER_ERROR_INVALID_ID = 13

C - `GtkBuilderError`

## `ButtonBoxStyle`

Used to dictate the style that a #GtkButtonBox uses to layout the buttons it
contains.

- GTK_BUTTONBOX_SPREAD = 1
- GTK_BUTTONBOX_EDGE = 2
- GTK_BUTTONBOX_START = 3
- GTK_BUTTONBOX_END = 4
- GTK_BUTTONBOX_CENTER = 5
- GTK_BUTTONBOX_EXPAND = 6

C - `GtkButtonBoxStyle`

## `ButtonRole`

The role specifies the desired appearance of a #GtkModelButton.

- GTK_BUTTON_ROLE_NORMAL = 0
- GTK_BUTTON_ROLE_CHECK = 1
- GTK_BUTTON_ROLE_RADIO = 2

C - `GtkButtonRole`

## `ButtonsType`

Prebuilt sets of buttons for the dialog. If
none of these choices are appropriate, simply use %GTK_BUTTONS_NONE
then call gtk_dialog_add_buttons().

> Please note that %GTK_BUTTONS_OK, %GTK_BUTTONS_YES_NO
> and %GTK_BUTTONS_OK_CANCEL are discouraged by the
> [GNOME Human Interface Guidelines](http://library.gnome.org/devel/hig-book/stable/).

- GTK_BUTTONS_NONE = 0
- GTK_BUTTONS_OK = 1
- GTK_BUTTONS_CLOSE = 2
- GTK_BUTTONS_CANCEL = 3
- GTK_BUTTONS_YES_NO = 4
- GTK_BUTTONS_OK_CANCEL = 5

C - `GtkButtonsType`

## `CellRendererAccelMode`

Determines if the edited accelerators are GTK+ accelerators. If
they are, consumed modifiers are suppressed, only accelerators
accepted by GTK+ are allowed, and the accelerators are rendered
in the same way as they are in menus.

- GTK_CELL_RENDERER_ACCEL_MODE_GTK = 0
- GTK_CELL_RENDERER_ACCEL_MODE_OTHER = 1
- GTK_CELL_RENDERER_ACCEL_MODE_MODIFIER_TAP = 2

C - `GtkCellRendererAccelMode`

## `CellRendererMode`

Identifies how the user can interact with a particular cell.

- GTK_CELL_RENDERER_MODE_INERT = 0
- GTK_CELL_RENDERER_MODE_ACTIVATABLE = 1
- GTK_CELL_RENDERER_MODE_EDITABLE = 2

C - `GtkCellRendererMode`

## `CornerType`

Specifies which corner a child widget should be placed in when packed into
a #GtkScrolledWindow. This is effectively the opposite of where the scroll
bars are placed.

- GTK_CORNER_TOP_LEFT = 0
- GTK_CORNER_BOTTOM_LEFT = 1
- GTK_CORNER_TOP_RIGHT = 2
- GTK_CORNER_BOTTOM_RIGHT = 3

C - `GtkCornerType`

## `CssProviderError`

Error codes for %GTK_CSS_PROVIDER_ERROR.

- GTK_CSS_PROVIDER_ERROR_FAILED = 0
- GTK_CSS_PROVIDER_ERROR_SYNTAX = 1
- GTK_CSS_PROVIDER_ERROR_IMPORT = 2
- GTK_CSS_PROVIDER_ERROR_NAME = 3
- GTK_CSS_PROVIDER_ERROR_DEPRECATED = 4
- GTK_CSS_PROVIDER_ERROR_UNKNOWN_VALUE = 5

C - `GtkCssProviderError`

## `CssSectionType`

The different types of sections indicate parts of a CSS document as
parsed by GTK’s CSS parser. They are oriented towards the
[CSS Grammar](http://www.w3.org/TR/CSS21/grammar.html),
but may contain extensions.

More types might be added in the future as the parser incorporates
more features.

- GTK_CSS_SECTION_DOCUMENT = 0
- GTK_CSS_SECTION_IMPORT = 1
- GTK_CSS_SECTION_COLOR_DEFINITION = 2
- GTK_CSS_SECTION_BINDING_SET = 3
- GTK_CSS_SECTION_RULESET = 4
- GTK_CSS_SECTION_SELECTOR = 5
- GTK_CSS_SECTION_DECLARATION = 6
- GTK_CSS_SECTION_VALUE = 7
- GTK_CSS_SECTION_KEYFRAMES = 8

C - `GtkCssSectionType`

## `DeleteType`

See also: #GtkEntry::delete-from-cursor.

- GTK_DELETE_CHARS = 0
- GTK_DELETE_WORD_ENDS = 1
- GTK_DELETE_WORDS = 2
- GTK_DELETE_DISPLAY_LINES = 3
- GTK_DELETE_DISPLAY_LINE_ENDS = 4
- GTK_DELETE_PARAGRAPH_ENDS = 5
- GTK_DELETE_PARAGRAPHS = 6
- GTK_DELETE_WHITESPACE = 7

C - `GtkDeleteType`

## `DirectionType`

Focus movement types.

- GTK_DIR_TAB_FORWARD = 0
- GTK_DIR_TAB_BACKWARD = 1
- GTK_DIR_UP = 2
- GTK_DIR_DOWN = 3
- GTK_DIR_LEFT = 4
- GTK_DIR_RIGHT = 5

C - `GtkDirectionType`

## `DragResult`

Gives an indication why a drag operation failed.
The value can by obtained by connecting to the
&num;GtkWidget::drag-failed signal.

- GTK_DRAG_RESULT_SUCCESS = 0
- GTK_DRAG_RESULT_NO_TARGET = 1
- GTK_DRAG_RESULT_USER_CANCELLED = 2
- GTK_DRAG_RESULT_TIMEOUT_EXPIRED = 3
- GTK_DRAG_RESULT_GRAB_BROKEN = 4
- GTK_DRAG_RESULT_ERROR = 5

C - `GtkDragResult`

## `EntryIconPosition`

Specifies the side of the entry at which an icon is placed.

- GTK_ENTRY_ICON_PRIMARY = 0
- GTK_ENTRY_ICON_SECONDARY = 1

C - `GtkEntryIconPosition`

## `EventSequenceState`

Describes the state of a #GdkEventSequence in a #GtkGesture.

- GTK_EVENT_SEQUENCE_NONE = 0
- GTK_EVENT_SEQUENCE_CLAIMED = 1
- GTK_EVENT_SEQUENCE_DENIED = 2

C - `GtkEventSequenceState`

## `ExpanderStyle`

Used to specify the style of the expanders drawn by a #GtkTreeView.

- GTK_EXPANDER_COLLAPSED = 0
- GTK_EXPANDER_SEMI_COLLAPSED = 1
- GTK_EXPANDER_SEMI_EXPANDED = 2
- GTK_EXPANDER_EXPANDED = 3

C - `GtkExpanderStyle`

## `FileChooserAction`

Describes whether a #GtkFileChooser is being used to open existing files
or to save to a possibly new file.

- GTK_FILE_CHOOSER_ACTION_OPEN = 0
- GTK_FILE_CHOOSER_ACTION_SAVE = 1
- GTK_FILE_CHOOSER_ACTION_SELECT_FOLDER = 2
- GTK_FILE_CHOOSER_ACTION_CREATE_FOLDER = 3

C - `GtkFileChooserAction`

## `FileChooserConfirmation`

Used as a return value of handlers for the
&num;GtkFileChooser::confirm-overwrite signal of a #GtkFileChooser. This
value determines whether the file chooser will present the stock
confirmation dialog, accept the user’s choice of a filename, or
let the user choose another filename.

- GTK_FILE_CHOOSER_CONFIRMATION_CONFIRM = 0
- GTK_FILE_CHOOSER_CONFIRMATION_ACCEPT_FILENAME = 1
- GTK_FILE_CHOOSER_CONFIRMATION_SELECT_AGAIN = 2

C - `GtkFileChooserConfirmation`

## `FileChooserError`

These identify the various errors that can occur while calling
&num;GtkFileChooser functions.

- GTK_FILE_CHOOSER_ERROR_NONEXISTENT = 0
- GTK_FILE_CHOOSER_ERROR_BAD_FILENAME = 1
- GTK_FILE_CHOOSER_ERROR_ALREADY_EXISTS = 2
- GTK_FILE_CHOOSER_ERROR_INCOMPLETE_HOSTNAME = 3

C - `GtkFileChooserError`

## `IMPreeditStyle`

Style for input method preedit. See also
&num;GtkSettings:gtk-im-preedit-style

- GTK_IM_PREEDIT_NOTHING = 0
- GTK_IM_PREEDIT_CALLBACK = 1
- GTK_IM_PREEDIT_NONE = 2

C - `GtkIMPreeditStyle`

## `IMStatusStyle`

Style for input method status. See also
&num;GtkSettings:gtk-im-status-style

- GTK_IM_STATUS_NOTHING = 0
- GTK_IM_STATUS_CALLBACK = 1
- GTK_IM_STATUS_NONE = 2

C - `GtkIMStatusStyle`

## `IconSize`

Built-in stock icon sizes.

- GTK_ICON_SIZE_INVALID = 0
- GTK_ICON_SIZE_MENU = 1
- GTK_ICON_SIZE_SMALL_TOOLBAR = 2
- GTK_ICON_SIZE_LARGE_TOOLBAR = 3
- GTK_ICON_SIZE_BUTTON = 4
- GTK_ICON_SIZE_DND = 5
- GTK_ICON_SIZE_DIALOG = 6

C - `GtkIconSize`

## `IconThemeError`

Error codes for GtkIconTheme operations.

- GTK_ICON_THEME_NOT_FOUND = 0
- GTK_ICON_THEME_FAILED = 1

C - `GtkIconThemeError`

## `IconViewDropPosition`

An enum for determining where a dropped item goes.

- GTK_ICON_VIEW_NO_DROP = 0
- GTK_ICON_VIEW_DROP_INTO = 1
- GTK_ICON_VIEW_DROP_LEFT = 2
- GTK_ICON_VIEW_DROP_RIGHT = 3
- GTK_ICON_VIEW_DROP_ABOVE = 4
- GTK_ICON_VIEW_DROP_BELOW = 5

C - `GtkIconViewDropPosition`

## `ImageType`

Describes the image data representation used by a #GtkImage. If you
want to get the image from the widget, you can only get the
currently-stored representation. e.g.  if the
gtk_image_get_storage_type() returns #GTK_IMAGE_PIXBUF, then you can
call gtk_image_get_pixbuf() but not gtk_image_get_stock().  For empty
images, you can request any storage type (call any of the "get"
functions), but they will all return %NULL values.

- GTK_IMAGE_EMPTY = 0
- GTK_IMAGE_PIXBUF = 1
- GTK_IMAGE_STOCK = 2
- GTK_IMAGE_ICON_SET = 3
- GTK_IMAGE_ANIMATION = 4
- GTK_IMAGE_ICON_NAME = 5
- GTK_IMAGE_GICON = 6
- GTK_IMAGE_SURFACE = 7

C - `GtkImageType`

## `InputPurpose`

Describes primary purpose of the input widget. This information is
useful for on-screen keyboards and similar input methods to decide
which keys should be presented to the user.

Note that the purpose is not meant to impose a totally strict rule
about allowed characters, and does not replace input validation.
It is fine for an on-screen keyboard to let the user override the
character set restriction that is expressed by the purpose. The
application is expected to validate the entry contents, even if
it specified a purpose.

The difference between @GTK_INPUT_PURPOSE_DIGITS and
@GTK_INPUT_PURPOSE_NUMBER is that the former accepts only digits
while the latter also some punctuation (like commas or points, plus,
minus) and “e” or “E” as in 3.14E+000.

This enumeration may be extended in the future; input methods should
interpret unknown values as “free form”.

- GTK_INPUT_PURPOSE_FREE_FORM = 0
- GTK_INPUT_PURPOSE_ALPHA = 1
- GTK_INPUT_PURPOSE_DIGITS = 2
- GTK_INPUT_PURPOSE_NUMBER = 3
- GTK_INPUT_PURPOSE_PHONE = 4
- GTK_INPUT_PURPOSE_URL = 5
- GTK_INPUT_PURPOSE_EMAIL = 6
- GTK_INPUT_PURPOSE_NAME = 7
- GTK_INPUT_PURPOSE_PASSWORD = 8
- GTK_INPUT_PURPOSE_PIN = 9

C - `GtkInputPurpose`

## `Justification`

Used for justifying the text inside a #GtkLabel widget. (See also
&num;GtkAlignment).

- GTK_JUSTIFY_LEFT = 0
- GTK_JUSTIFY_RIGHT = 1
- GTK_JUSTIFY_CENTER = 2
- GTK_JUSTIFY_FILL = 3

C - `GtkJustification`

## `LevelBarMode`

Describes how #GtkLevelBar contents should be rendered.
Note that this enumeration could be extended with additional modes
in the future.

- GTK_LEVEL_BAR_MODE_CONTINUOUS = 0
- GTK_LEVEL_BAR_MODE_DISCRETE = 1

C - `GtkLevelBarMode`

## `License`

The type of license for an application.

This enumeration can be expanded at later date.

- GTK_LICENSE_UNKNOWN = 0
- GTK_LICENSE_CUSTOM = 1
- GTK_LICENSE_GPL_2_0 = 2
- GTK_LICENSE_GPL_3_0 = 3
- GTK_LICENSE_LGPL_2_1 = 4
- GTK_LICENSE_LGPL_3_0 = 5
- GTK_LICENSE_BSD = 6
- GTK_LICENSE_MIT_X11 = 7
- GTK_LICENSE_ARTISTIC = 8
- GTK_LICENSE_GPL_2_0_ONLY = 9
- GTK_LICENSE_GPL_3_0_ONLY = 10
- GTK_LICENSE_LGPL_2_1_ONLY = 11
- GTK_LICENSE_LGPL_3_0_ONLY = 12
- GTK_LICENSE_AGPL_3_0 = 13
- GTK_LICENSE_AGPL_3_0_ONLY = 14

C - `GtkLicense`

## `MenuDirectionType`

An enumeration representing directional movements within a menu.

- GTK_MENU_DIR_PARENT = 0
- GTK_MENU_DIR_CHILD = 1
- GTK_MENU_DIR_NEXT = 2
- GTK_MENU_DIR_PREV = 3

C - `GtkMenuDirectionType`

## `MessageType`

The type of message being displayed in the dialog.

- GTK_MESSAGE_INFO = 0
- GTK_MESSAGE_WARNING = 1
- GTK_MESSAGE_QUESTION = 2
- GTK_MESSAGE_ERROR = 3
- GTK_MESSAGE_OTHER = 4

C - `GtkMessageType`

## `MovementStep`



- GTK_MOVEMENT_LOGICAL_POSITIONS = 0
- GTK_MOVEMENT_VISUAL_POSITIONS = 1
- GTK_MOVEMENT_WORDS = 2
- GTK_MOVEMENT_DISPLAY_LINES = 3
- GTK_MOVEMENT_DISPLAY_LINE_ENDS = 4
- GTK_MOVEMENT_PARAGRAPHS = 5
- GTK_MOVEMENT_PARAGRAPH_ENDS = 6
- GTK_MOVEMENT_PAGES = 7
- GTK_MOVEMENT_BUFFER_ENDS = 8
- GTK_MOVEMENT_HORIZONTAL_PAGES = 9

C - `GtkMovementStep`

## `NotebookTab`



- GTK_NOTEBOOK_TAB_FIRST = 0
- GTK_NOTEBOOK_TAB_LAST = 1

C - `GtkNotebookTab`

## `NumberUpLayout`

Used to determine the layout of pages on a sheet when printing
multiple pages per sheet.

- GTK_NUMBER_UP_LAYOUT_LEFT_TO_RIGHT_TOP_TO_BOTTOM = 0
- GTK_NUMBER_UP_LAYOUT_LEFT_TO_RIGHT_BOTTOM_TO_TOP = 1
- GTK_NUMBER_UP_LAYOUT_RIGHT_TO_LEFT_TOP_TO_BOTTOM = 2
- GTK_NUMBER_UP_LAYOUT_RIGHT_TO_LEFT_BOTTOM_TO_TOP = 3
- GTK_NUMBER_UP_LAYOUT_TOP_TO_BOTTOM_LEFT_TO_RIGHT = 4
- GTK_NUMBER_UP_LAYOUT_TOP_TO_BOTTOM_RIGHT_TO_LEFT = 5
- GTK_NUMBER_UP_LAYOUT_BOTTOM_TO_TOP_LEFT_TO_RIGHT = 6
- GTK_NUMBER_UP_LAYOUT_BOTTOM_TO_TOP_RIGHT_TO_LEFT = 7

C - `GtkNumberUpLayout`

## `Orientation`

Represents the orientation of widgets and other objects which can be switched
between horizontal and vertical orientation on the fly, like #GtkToolbar or
&num;GtkGesturePan.

- GTK_ORIENTATION_HORIZONTAL = 0
- GTK_ORIENTATION_VERTICAL = 1

C - `GtkOrientation`

## `PackDirection`

Determines how widgets should be packed inside menubars
and menuitems contained in menubars.

- GTK_PACK_DIRECTION_LTR = 0
- GTK_PACK_DIRECTION_RTL = 1
- GTK_PACK_DIRECTION_TTB = 2
- GTK_PACK_DIRECTION_BTT = 3

C - `GtkPackDirection`

## `PackType`

Represents the packing location #GtkBox children. (See: #GtkVBox,
&num;GtkHBox, and #GtkButtonBox).

- GTK_PACK_START = 0
- GTK_PACK_END = 1

C - `GtkPackType`

## `PadActionType`

The type of a pad action.

- GTK_PAD_ACTION_BUTTON = 0
- GTK_PAD_ACTION_RING = 1
- GTK_PAD_ACTION_STRIP = 2

C - `GtkPadActionType`

## `PageOrientation`

See also gtk_print_settings_set_orientation().

- GTK_PAGE_ORIENTATION_PORTRAIT = 0
- GTK_PAGE_ORIENTATION_LANDSCAPE = 1
- GTK_PAGE_ORIENTATION_REVERSE_PORTRAIT = 2
- GTK_PAGE_ORIENTATION_REVERSE_LANDSCAPE = 3

C - `GtkPageOrientation`

## `PageSet`

See also gtk_print_job_set_page_set().

- GTK_PAGE_SET_ALL = 0
- GTK_PAGE_SET_EVEN = 1
- GTK_PAGE_SET_ODD = 2

C - `GtkPageSet`

## `PanDirection`

Describes the panning direction of a #GtkGesturePan

- GTK_PAN_DIRECTION_LEFT = 0
- GTK_PAN_DIRECTION_RIGHT = 1
- GTK_PAN_DIRECTION_UP = 2
- GTK_PAN_DIRECTION_DOWN = 3

C - `GtkPanDirection`

## `PathPriorityType`

Priorities for path lookups.
See also gtk_binding_set_add_path().

- GTK_PATH_PRIO_LOWEST = 0
- GTK_PATH_PRIO_GTK = 4
- GTK_PATH_PRIO_APPLICATION = 8
- GTK_PATH_PRIO_THEME = 10
- GTK_PATH_PRIO_RC = 12
- GTK_PATH_PRIO_HIGHEST = 15

C - `GtkPathPriorityType`

## `PathType`

Widget path types.
See also gtk_binding_set_add_path().

- GTK_PATH_WIDGET = 0
- GTK_PATH_WIDGET_CLASS = 1
- GTK_PATH_CLASS = 2

C - `GtkPathType`

## `PolicyType`

Determines how the size should be computed to achieve the one of the
visibility mode for the scrollbars.

- GTK_POLICY_ALWAYS = 0
- GTK_POLICY_AUTOMATIC = 1
- GTK_POLICY_NEVER = 2
- GTK_POLICY_EXTERNAL = 3

C - `GtkPolicyType`

## `PopoverConstraint`

Describes constraints to positioning of popovers. More values
may be added to this enumeration in the future.

- GTK_POPOVER_CONSTRAINT_NONE = 0
- GTK_POPOVER_CONSTRAINT_WINDOW = 1

C - `GtkPopoverConstraint`

## `PositionType`

Describes which edge of a widget a certain feature is positioned at, e.g. the
tabs of a #GtkNotebook, the handle of a #GtkHandleBox or the label of a
&num;GtkScale.

- GTK_POS_LEFT = 0
- GTK_POS_RIGHT = 1
- GTK_POS_TOP = 2
- GTK_POS_BOTTOM = 3

C - `GtkPositionType`

## `PrintDuplex`

See also gtk_print_settings_set_duplex().

- GTK_PRINT_DUPLEX_SIMPLEX = 0
- GTK_PRINT_DUPLEX_HORIZONTAL = 1
- GTK_PRINT_DUPLEX_VERTICAL = 2

C - `GtkPrintDuplex`

## `PrintError`

Error codes that identify various errors that can occur while
using the GTK+ printing support.

- GTK_PRINT_ERROR_GENERAL = 0
- GTK_PRINT_ERROR_INTERNAL_ERROR = 1
- GTK_PRINT_ERROR_NOMEM = 2
- GTK_PRINT_ERROR_INVALID_FILE = 3

C - `GtkPrintError`

## `PrintOperationAction`

The @action parameter to gtk_print_operation_run()
determines what action the print operation should perform.

- GTK_PRINT_OPERATION_ACTION_PRINT_DIALOG = 0
- GTK_PRINT_OPERATION_ACTION_PRINT = 1
- GTK_PRINT_OPERATION_ACTION_PREVIEW = 2
- GTK_PRINT_OPERATION_ACTION_EXPORT = 3

C - `GtkPrintOperationAction`

## `PrintOperationResult`

A value of this type is returned by gtk_print_operation_run().

- GTK_PRINT_OPERATION_RESULT_ERROR = 0
- GTK_PRINT_OPERATION_RESULT_APPLY = 1
- GTK_PRINT_OPERATION_RESULT_CANCEL = 2
- GTK_PRINT_OPERATION_RESULT_IN_PROGRESS = 3

C - `GtkPrintOperationResult`

## `PrintPages`

See also gtk_print_job_set_pages()

- GTK_PRINT_PAGES_ALL = 0
- GTK_PRINT_PAGES_CURRENT = 1
- GTK_PRINT_PAGES_RANGES = 2
- GTK_PRINT_PAGES_SELECTION = 3

C - `GtkPrintPages`

## `PrintQuality`

See also gtk_print_settings_set_quality().

- GTK_PRINT_QUALITY_LOW = 0
- GTK_PRINT_QUALITY_NORMAL = 1
- GTK_PRINT_QUALITY_HIGH = 2
- GTK_PRINT_QUALITY_DRAFT = 3

C - `GtkPrintQuality`

## `PrintStatus`

The status gives a rough indication of the completion of a running
print operation.

- GTK_PRINT_STATUS_INITIAL = 0
- GTK_PRINT_STATUS_PREPARING = 1
- GTK_PRINT_STATUS_GENERATING_DATA = 2
- GTK_PRINT_STATUS_SENDING_DATA = 3
- GTK_PRINT_STATUS_PENDING = 4
- GTK_PRINT_STATUS_PENDING_ISSUE = 5
- GTK_PRINT_STATUS_PRINTING = 6
- GTK_PRINT_STATUS_FINISHED = 7
- GTK_PRINT_STATUS_FINISHED_ABORTED = 8

C - `GtkPrintStatus`

## `PropagationPhase`

Describes the stage at which events are fed into a #GtkEventController.

- GTK_PHASE_NONE = 0
- GTK_PHASE_CAPTURE = 1
- GTK_PHASE_BUBBLE = 2
- GTK_PHASE_TARGET = 3

C - `GtkPropagationPhase`

## `RcTokenType`

The #GtkRcTokenType enumeration represents the tokens
in the RC file. It is exposed so that theme engines
can reuse these tokens when parsing the theme-engine
specific portions of a RC file.

- GTK_RC_TOKEN_INVALID = 270
- GTK_RC_TOKEN_INCLUDE = 271
- GTK_RC_TOKEN_NORMAL = 272
- GTK_RC_TOKEN_ACTIVE = 273
- GTK_RC_TOKEN_PRELIGHT = 274
- GTK_RC_TOKEN_SELECTED = 275
- GTK_RC_TOKEN_INSENSITIVE = 276
- GTK_RC_TOKEN_FG = 277
- GTK_RC_TOKEN_BG = 278
- GTK_RC_TOKEN_TEXT = 279
- GTK_RC_TOKEN_BASE = 280
- GTK_RC_TOKEN_XTHICKNESS = 281
- GTK_RC_TOKEN_YTHICKNESS = 282
- GTK_RC_TOKEN_FONT = 283
- GTK_RC_TOKEN_FONTSET = 284
- GTK_RC_TOKEN_FONT_NAME = 285
- GTK_RC_TOKEN_BG_PIXMAP = 286
- GTK_RC_TOKEN_PIXMAP_PATH = 287
- GTK_RC_TOKEN_STYLE = 288
- GTK_RC_TOKEN_BINDING = 289
- GTK_RC_TOKEN_BIND = 290
- GTK_RC_TOKEN_WIDGET = 291
- GTK_RC_TOKEN_WIDGET_CLASS = 292
- GTK_RC_TOKEN_CLASS = 293
- GTK_RC_TOKEN_LOWEST = 294
- GTK_RC_TOKEN_GTK = 295
- GTK_RC_TOKEN_APPLICATION = 296
- GTK_RC_TOKEN_THEME = 297
- GTK_RC_TOKEN_RC = 298
- GTK_RC_TOKEN_HIGHEST = 299
- GTK_RC_TOKEN_ENGINE = 300
- GTK_RC_TOKEN_MODULE_PATH = 301
- GTK_RC_TOKEN_IM_MODULE_PATH = 302
- GTK_RC_TOKEN_IM_MODULE_FILE = 303
- GTK_RC_TOKEN_STOCK = 304
- GTK_RC_TOKEN_LTR = 305
- GTK_RC_TOKEN_RTL = 306
- GTK_RC_TOKEN_COLOR = 307
- GTK_RC_TOKEN_UNBIND = 308
- GTK_RC_TOKEN_LAST = 309

C - `GtkRcTokenType`

## `RecentChooserError`

These identify the various errors that can occur while calling
&num;GtkRecentChooser functions.

- GTK_RECENT_CHOOSER_ERROR_NOT_FOUND = 0
- GTK_RECENT_CHOOSER_ERROR_INVALID_URI = 1

C - `GtkRecentChooserError`

## `RecentManagerError`

Error codes for #GtkRecentManager operations

- GTK_RECENT_MANAGER_ERROR_NOT_FOUND = 0
- GTK_RECENT_MANAGER_ERROR_INVALID_URI = 1
- GTK_RECENT_MANAGER_ERROR_INVALID_ENCODING = 2
- GTK_RECENT_MANAGER_ERROR_NOT_REGISTERED = 3
- GTK_RECENT_MANAGER_ERROR_READ = 4
- GTK_RECENT_MANAGER_ERROR_WRITE = 5
- GTK_RECENT_MANAGER_ERROR_UNKNOWN = 6

C - `GtkRecentManagerError`

## `RecentSortType`

Used to specify the sorting method to be applyed to the recently
used resource list.

- GTK_RECENT_SORT_NONE = 0
- GTK_RECENT_SORT_MRU = 1
- GTK_RECENT_SORT_LRU = 2
- GTK_RECENT_SORT_CUSTOM = 3

C - `GtkRecentSortType`

## `ReliefStyle`

Indicated the relief to be drawn around a #GtkButton.

- GTK_RELIEF_NORMAL = 0
- GTK_RELIEF_HALF = 1
- GTK_RELIEF_NONE = 2

C - `GtkReliefStyle`

## `ResizeMode`



- GTK_RESIZE_PARENT = 0
- GTK_RESIZE_QUEUE = 1
- GTK_RESIZE_IMMEDIATE = 2

C - `GtkResizeMode`

## `ResponseType`

Predefined values for use as response ids in gtk_dialog_add_button().
All predefined values are negative; GTK+ leaves values of 0 or greater for
application-defined response ids.

- GTK_RESPONSE_NONE = -1
- GTK_RESPONSE_REJECT = -2
- GTK_RESPONSE_ACCEPT = -3
- GTK_RESPONSE_DELETE_EVENT = -4
- GTK_RESPONSE_OK = -5
- GTK_RESPONSE_CANCEL = -6
- GTK_RESPONSE_CLOSE = -7
- GTK_RESPONSE_YES = -8
- GTK_RESPONSE_NO = -9
- GTK_RESPONSE_APPLY = -10
- GTK_RESPONSE_HELP = -11

C - `GtkResponseType`

## `RevealerTransitionType`

These enumeration values describe the possible transitions
when the child of a #GtkRevealer widget is shown or hidden.

- GTK_REVEALER_TRANSITION_TYPE_NONE = 0
- GTK_REVEALER_TRANSITION_TYPE_CROSSFADE = 1
- GTK_REVEALER_TRANSITION_TYPE_SLIDE_RIGHT = 2
- GTK_REVEALER_TRANSITION_TYPE_SLIDE_LEFT = 3
- GTK_REVEALER_TRANSITION_TYPE_SLIDE_UP = 4
- GTK_REVEALER_TRANSITION_TYPE_SLIDE_DOWN = 5

C - `GtkRevealerTransitionType`

## `ScrollStep`



- GTK_SCROLL_STEPS = 0
- GTK_SCROLL_PAGES = 1
- GTK_SCROLL_ENDS = 2
- GTK_SCROLL_HORIZONTAL_STEPS = 3
- GTK_SCROLL_HORIZONTAL_PAGES = 4
- GTK_SCROLL_HORIZONTAL_ENDS = 5

C - `GtkScrollStep`

## `ScrollType`

Scrolling types.

- GTK_SCROLL_NONE = 0
- GTK_SCROLL_JUMP = 1
- GTK_SCROLL_STEP_BACKWARD = 2
- GTK_SCROLL_STEP_FORWARD = 3
- GTK_SCROLL_PAGE_BACKWARD = 4
- GTK_SCROLL_PAGE_FORWARD = 5
- GTK_SCROLL_STEP_UP = 6
- GTK_SCROLL_STEP_DOWN = 7
- GTK_SCROLL_PAGE_UP = 8
- GTK_SCROLL_PAGE_DOWN = 9
- GTK_SCROLL_STEP_LEFT = 10
- GTK_SCROLL_STEP_RIGHT = 11
- GTK_SCROLL_PAGE_LEFT = 12
- GTK_SCROLL_PAGE_RIGHT = 13
- GTK_SCROLL_START = 14
- GTK_SCROLL_END = 15

C - `GtkScrollType`

## `ScrollablePolicy`

Defines the policy to be used in a scrollable widget when updating
the scrolled window adjustments in a given orientation.

- GTK_SCROLL_MINIMUM = 0
- GTK_SCROLL_NATURAL = 1

C - `GtkScrollablePolicy`

## `SelectionMode`

Used to control what selections users are allowed to make.

- GTK_SELECTION_NONE = 0
- GTK_SELECTION_SINGLE = 1
- GTK_SELECTION_BROWSE = 2
- GTK_SELECTION_MULTIPLE = 3

C - `GtkSelectionMode`

## `SensitivityType`

Determines how GTK+ handles the sensitivity of stepper arrows
at the end of range widgets.

- GTK_SENSITIVITY_AUTO = 0
- GTK_SENSITIVITY_ON = 1
- GTK_SENSITIVITY_OFF = 2

C - `GtkSensitivityType`

## `ShadowType`

Used to change the appearance of an outline typically provided by a #GtkFrame.

Note that many themes do not differentiate the appearance of the
various shadow types: Either their is no visible shadow (@GTK_SHADOW_NONE),
or there is (any other value).

- GTK_SHADOW_NONE = 0
- GTK_SHADOW_IN = 1
- GTK_SHADOW_OUT = 2
- GTK_SHADOW_ETCHED_IN = 3
- GTK_SHADOW_ETCHED_OUT = 4

C - `GtkShadowType`

## `ShortcutType`

GtkShortcutType specifies the kind of shortcut that is being described.
More values may be added to this enumeration over time.

- GTK_SHORTCUT_ACCELERATOR = 0
- GTK_SHORTCUT_GESTURE_PINCH = 1
- GTK_SHORTCUT_GESTURE_STRETCH = 2
- GTK_SHORTCUT_GESTURE_ROTATE_CLOCKWISE = 3
- GTK_SHORTCUT_GESTURE_ROTATE_COUNTERCLOCKWISE = 4
- GTK_SHORTCUT_GESTURE_TWO_FINGER_SWIPE_LEFT = 5
- GTK_SHORTCUT_GESTURE_TWO_FINGER_SWIPE_RIGHT = 6
- GTK_SHORTCUT_GESTURE = 7

C - `GtkShortcutType`

## `SizeGroupMode`

The mode of the size group determines the directions in which the size
group affects the requested sizes of its component widgets.

- GTK_SIZE_GROUP_NONE = 0
- GTK_SIZE_GROUP_HORIZONTAL = 1
- GTK_SIZE_GROUP_VERTICAL = 2
- GTK_SIZE_GROUP_BOTH = 3

C - `GtkSizeGroupMode`

## `SizeRequestMode`

Specifies a preference for height-for-width or
width-for-height geometry management.

- GTK_SIZE_REQUEST_HEIGHT_FOR_WIDTH = 0
- GTK_SIZE_REQUEST_WIDTH_FOR_HEIGHT = 1
- GTK_SIZE_REQUEST_CONSTANT_SIZE = 2

C - `GtkSizeRequestMode`

## `SortType`

Determines the direction of a sort.

- GTK_SORT_ASCENDING = 0
- GTK_SORT_DESCENDING = 1

C - `GtkSortType`

## `SpinButtonUpdatePolicy`

The spin button update policy determines whether the spin button displays
values even if they are outside the bounds of its adjustment.
See gtk_spin_button_set_update_policy().

- GTK_UPDATE_ALWAYS = 0
- GTK_UPDATE_IF_VALID = 1

C - `GtkSpinButtonUpdatePolicy`

## `SpinType`

The values of the GtkSpinType enumeration are used to specify the
change to make in gtk_spin_button_spin().

- GTK_SPIN_STEP_FORWARD = 0
- GTK_SPIN_STEP_BACKWARD = 1
- GTK_SPIN_PAGE_FORWARD = 2
- GTK_SPIN_PAGE_BACKWARD = 3
- GTK_SPIN_HOME = 4
- GTK_SPIN_END = 5
- GTK_SPIN_USER_DEFINED = 6

C - `GtkSpinType`

## `StackTransitionType`

These enumeration values describe the possible transitions
between pages in a #GtkStack widget.

New values may be added to this enumeration over time.

- GTK_STACK_TRANSITION_TYPE_NONE = 0
- GTK_STACK_TRANSITION_TYPE_CROSSFADE = 1
- GTK_STACK_TRANSITION_TYPE_SLIDE_RIGHT = 2
- GTK_STACK_TRANSITION_TYPE_SLIDE_LEFT = 3
- GTK_STACK_TRANSITION_TYPE_SLIDE_UP = 4
- GTK_STACK_TRANSITION_TYPE_SLIDE_DOWN = 5
- GTK_STACK_TRANSITION_TYPE_SLIDE_LEFT_RIGHT = 6
- GTK_STACK_TRANSITION_TYPE_SLIDE_UP_DOWN = 7
- GTK_STACK_TRANSITION_TYPE_OVER_UP = 8
- GTK_STACK_TRANSITION_TYPE_OVER_DOWN = 9
- GTK_STACK_TRANSITION_TYPE_OVER_LEFT = 10
- GTK_STACK_TRANSITION_TYPE_OVER_RIGHT = 11
- GTK_STACK_TRANSITION_TYPE_UNDER_UP = 12
- GTK_STACK_TRANSITION_TYPE_UNDER_DOWN = 13
- GTK_STACK_TRANSITION_TYPE_UNDER_LEFT = 14
- GTK_STACK_TRANSITION_TYPE_UNDER_RIGHT = 15
- GTK_STACK_TRANSITION_TYPE_OVER_UP_DOWN = 16
- GTK_STACK_TRANSITION_TYPE_OVER_DOWN_UP = 17
- GTK_STACK_TRANSITION_TYPE_OVER_LEFT_RIGHT = 18
- GTK_STACK_TRANSITION_TYPE_OVER_RIGHT_LEFT = 19

C - `GtkStackTransitionType`

## `StateType`

This type indicates the current state of a widget; the state determines how
the widget is drawn. The #GtkStateType enumeration is also used to
identify different colors in a #GtkStyle for drawing, so states can be
used for subparts of a widget as well as entire widgets.

- GTK_STATE_NORMAL = 0
- GTK_STATE_ACTIVE = 1
- GTK_STATE_PRELIGHT = 2
- GTK_STATE_SELECTED = 3
- GTK_STATE_INSENSITIVE = 4
- GTK_STATE_INCONSISTENT = 5
- GTK_STATE_FOCUSED = 6

C - `GtkStateType`

## `TextBufferTargetInfo`

These values are used as “info” for the targets contained in the
lists returned by gtk_text_buffer_get_copy_target_list() and
gtk_text_buffer_get_paste_target_list().

The values counts down from `-1` to avoid clashes
with application added drag destinations which usually start at 0.

- GTK_TEXT_BUFFER_TARGET_INFO_BUFFER_CONTENTS = -1
- GTK_TEXT_BUFFER_TARGET_INFO_RICH_TEXT = -2
- GTK_TEXT_BUFFER_TARGET_INFO_TEXT = -3

C - `GtkTextBufferTargetInfo`

## `TextDirection`

Reading directions for text.

- GTK_TEXT_DIR_NONE = 0
- GTK_TEXT_DIR_LTR = 1
- GTK_TEXT_DIR_RTL = 2

C - `GtkTextDirection`

## `TextExtendSelection`

Granularity types that extend the text selection. Use the
&num;GtkTextView::extend-selection signal to customize the selection.

- GTK_TEXT_EXTEND_SELECTION_WORD = 0
- GTK_TEXT_EXTEND_SELECTION_LINE = 1

C - `GtkTextExtendSelection`

## `TextViewLayer`

Used to reference the layers of #GtkTextView for the purpose of customized
drawing with the ::draw_layer vfunc.

- GTK_TEXT_VIEW_LAYER_BELOW = 0
- GTK_TEXT_VIEW_LAYER_ABOVE = 1
- GTK_TEXT_VIEW_LAYER_BELOW_TEXT = 2
- GTK_TEXT_VIEW_LAYER_ABOVE_TEXT = 3

C - `GtkTextViewLayer`

## `TextWindowType`

Used to reference the parts of #GtkTextView.

- GTK_TEXT_WINDOW_PRIVATE = 0
- GTK_TEXT_WINDOW_WIDGET = 1
- GTK_TEXT_WINDOW_TEXT = 2
- GTK_TEXT_WINDOW_LEFT = 3
- GTK_TEXT_WINDOW_RIGHT = 4
- GTK_TEXT_WINDOW_TOP = 5
- GTK_TEXT_WINDOW_BOTTOM = 6

C - `GtkTextWindowType`

## `ToolbarSpaceStyle`

Whether spacers are vertical lines or just blank.

- GTK_TOOLBAR_SPACE_EMPTY = 0
- GTK_TOOLBAR_SPACE_LINE = 1

C - `GtkToolbarSpaceStyle`

## `ToolbarStyle`

Used to customize the appearance of a #GtkToolbar. Note that
setting the toolbar style overrides the user’s preferences
for the default toolbar style.  Note that if the button has only
a label set and GTK_TOOLBAR_ICONS is used, the label will be
visible, and vice versa.

- GTK_TOOLBAR_ICONS = 0
- GTK_TOOLBAR_TEXT = 1
- GTK_TOOLBAR_BOTH = 2
- GTK_TOOLBAR_BOTH_HORIZ = 3

C - `GtkToolbarStyle`

## `TreeViewColumnSizing`

The sizing method the column uses to determine its width.  Please note
that @GTK_TREE_VIEW_COLUMN_AUTOSIZE are inefficient for large views, and
can make columns appear choppy.

- GTK_TREE_VIEW_COLUMN_GROW_ONLY = 0
- GTK_TREE_VIEW_COLUMN_AUTOSIZE = 1
- GTK_TREE_VIEW_COLUMN_FIXED = 2

C - `GtkTreeViewColumnSizing`

## `TreeViewDropPosition`

An enum for determining where a dropped row goes.

- GTK_TREE_VIEW_DROP_BEFORE = 0
- GTK_TREE_VIEW_DROP_AFTER = 1
- GTK_TREE_VIEW_DROP_INTO_OR_BEFORE = 2
- GTK_TREE_VIEW_DROP_INTO_OR_AFTER = 3

C - `GtkTreeViewDropPosition`

## `TreeViewGridLines`

Used to indicate which grid lines to draw in a tree view.

- GTK_TREE_VIEW_GRID_LINES_NONE = 0
- GTK_TREE_VIEW_GRID_LINES_HORIZONTAL = 1
- GTK_TREE_VIEW_GRID_LINES_VERTICAL = 2
- GTK_TREE_VIEW_GRID_LINES_BOTH = 3

C - `GtkTreeViewGridLines`

## `Unit`

See also gtk_print_settings_set_paper_width().

- GTK_UNIT_NONE = 0
- GTK_UNIT_POINTS = 1
- GTK_UNIT_INCH = 2
- GTK_UNIT_MM = 3

C - `GtkUnit`

## `WidgetHelpType`

Kinds of widget-specific help. Used by the ::show-help signal.

- GTK_WIDGET_HELP_TOOLTIP = 0
- GTK_WIDGET_HELP_WHATS_THIS = 1

C - `GtkWidgetHelpType`

## `WindowPosition`

Window placement can be influenced using this enumeration. Note that
using #GTK_WIN_POS_CENTER_ALWAYS is almost always a bad idea.
It won’t necessarily work well with all window managers or on all windowing systems.

- GTK_WIN_POS_NONE = 0
- GTK_WIN_POS_CENTER = 1
- GTK_WIN_POS_MOUSE = 2
- GTK_WIN_POS_CENTER_ALWAYS = 3
- GTK_WIN_POS_CENTER_ON_PARENT = 4

C - `GtkWindowPosition`

## `WindowType`

A #GtkWindow can be one of these types. Most things you’d consider a
“window” should have type #GTK_WINDOW_TOPLEVEL; windows with this type
are managed by the window manager and have a frame by default (call
gtk_window_set_decorated() to toggle the frame).  Windows with type
&num;GTK_WINDOW_POPUP are ignored by the window manager; window manager
keybindings won’t work on them, the window manager won’t decorate the
window with a frame, many GTK+ features that rely on the window
manager will not work (e.g. resize grips and
maximization/minimization). #GTK_WINDOW_POPUP is used to implement
widgets such as #GtkMenu or tooltips that you normally don’t think of
as windows per se. Nearly all windows should be #GTK_WINDOW_TOPLEVEL.
In particular, do not use #GTK_WINDOW_POPUP just to turn off
the window borders; use gtk_window_set_decorated() for that.

- GTK_WINDOW_TOPLEVEL = 0
- GTK_WINDOW_POPUP = 1

C - `GtkWindowType`

## `WrapMode`

Describes a type of line wrapping.

- GTK_WRAP_NONE = 0
- GTK_WRAP_CHAR = 1
- GTK_WRAP_WORD = 2
- GTK_WRAP_WORD_CHAR = 3

C - `GtkWrapMode`

t extend the text selection. Use the
&num;GtkTextView::extend-selection signal to customize the selection.

C - `GtkTextExtendSelection`

- GTK_TEXT_EXTEND_SELECTION_WORD = %!s(int=0)
- GTK_TEXT_EXTEND_SELECTION_LINE = %!s(int=1)
## `TextViewLayer`

Used to reference the layers of #GtkTextView for the purpose of customized
drawing with the ::draw_layer vfunc.

C - `GtkTextViewLayer`

- GTK_TEXT_VIEW_LAYER_BELOW = %!s(int=0)
- GTK_TEXT_VIEW_LAYER_ABOVE = %!s(int=1)
- GTK_TEXT_VIEW_LAYER_BELOW_TEXT = %!s(int=2)
- GTK_TEXT_VIEW_LAYER_ABOVE_TEXT = %!s(int=3)
## `TextWindowType`

Used to reference the parts of #GtkTextView.

C - `GtkTextWindowType`

- GTK_TEXT_WINDOW_PRIVATE = %!s(int=0)
- GTK_TEXT_WINDOW_WIDGET = %!s(int=1)
- GTK_TEXT_WINDOW_TEXT = %!s(int=2)
- GTK_TEXT_WINDOW_LEFT = %!s(int=3)
- GTK_TEXT_WINDOW_RIGHT = %!s(int=4)
- GTK_TEXT_WINDOW_TOP = %!s(int=5)
- GTK_TEXT_WINDOW_BOTTOM = %!s(int=6)
## `ToolbarSpaceStyle`

Whether spacers are vertical lines or just blank.

C - `GtkToolbarSpaceStyle`

- GTK_TOOLBAR_SPACE_EMPTY = %!s(int=0)
- GTK_TOOLBAR_SPACE_LINE = %!s(int=1)
## `ToolbarStyle`

Used to customize the appearance of a #GtkToolbar. Note that
setting the toolbar style overrides the user’s preferences
for the default toolbar style.  Note that if the button has only
a label set and GTK_TOOLBAR_ICONS is used, the label will be
visible, and vice versa.

C - `GtkToolbarStyle`

- GTK_TOOLBAR_ICONS = %!s(int=0)
- GTK_TOOLBAR_TEXT = %!s(int=1)
- GTK_TOOLBAR_BOTH = %!s(int=2)
- GTK_TOOLBAR_BOTH_HORIZ = %!s(int=3)
## `TreeViewColumnSizing`

The sizing method the column uses to determine its width.  Please note
that @GTK_TREE_VIEW_COLUMN_AUTOSIZE are inefficient for large views, and
can make columns appear choppy.

C - `GtkTreeViewColumnSizing`

- GTK_TREE_VIEW_COLUMN_GROW_ONLY = %!s(int=0)
- GTK_TREE_VIEW_COLUMN_AUTOSIZE = %!s(int=1)
- GTK_TREE_VIEW_COLUMN_FIXED = %!s(int=2)
## `TreeViewDropPosition`

An enum for determining where a dropped row goes.

C - `GtkTreeViewDropPosition`

- GTK_TREE_VIEW_DROP_BEFORE = %!s(int=0)
- GTK_TREE_VIEW_DROP_AFTER = %!s(int=1)
- GTK_TREE_VIEW_DROP_INTO_OR_BEFORE = %!s(int=2)
- GTK_TREE_VIEW_DROP_INTO_OR_AFTER = %!s(int=3)
## `TreeViewGridLines`

Used to indicate which grid lines to draw in a tree view.

C - `GtkTreeViewGridLines`

- GTK_TREE_VIEW_GRID_LINES_NONE = %!s(int=0)
- GTK_TREE_VIEW_GRID_LINES_HORIZONTAL = %!s(int=1)
- GTK_TREE_VIEW_GRID_LINES_VERTICAL = %!s(int=2)
- GTK_TREE_VIEW_GRID_LINES_BOTH = %!s(int=3)
## `Unit`

See also gtk_print_settings_set_paper_width().

C - `GtkUnit`

- GTK_UNIT_NONE = %!s(int=0)
- GTK_UNIT_POINTS = %!s(int=1)
- GTK_UNIT_INCH = %!s(int=2)
- GTK_UNIT_MM = %!s(int=3)
## `WidgetHelpType`

Kinds of widget-specific help. Used by the ::show-help signal.

C - `GtkWidgetHelpType`

- GTK_WIDGET_HELP_TOOLTIP = %!s(int=0)
- GTK_WIDGET_HELP_WHATS_THIS = %!s(int=1)
## `WindowPosition`

Window placement can be influenced using this enumeration. Note that
using #GTK_WIN_POS_CENTER_ALWAYS is almost always a bad idea.
It won’t necessarily work well with all window managers or on all windowing systems.

C - `GtkWindowPosition`

- GTK_WIN_POS_NONE = %!s(int=0)
- GTK_WIN_POS_CENTER = %!s(int=1)
- GTK_WIN_POS_MOUSE = %!s(int=2)
- GTK_WIN_POS_CENTER_ALWAYS = %!s(int=3)
- GTK_WIN_POS_CENTER_ON_PARENT = %!s(int=4)
## `WindowType`

A #GtkWindow can be one of these types. Most things you’d consider a
“window” should have type #GTK_WINDOW_TOPLEVEL; windows with this type
are managed by the window manager and have a frame by default (call
gtk_window_set_decorated() to toggle the frame).  Windows with type
&num;GTK_WINDOW_POPUP are ignored by the window manager; window manager
keybindings won’t work on them, the window manager won’t decorate the
window with a frame, many GTK+ features that rely on the window
manager will not work (e.g. resize grips and
maximization/minimization). #GTK_WINDOW_POPUP is used to implement
widgets such as #GtkMenu or tooltips that you normally don’t think of
as windows per se. Nearly all windows should be #GTK_WINDOW_TOPLEVEL.
In particular, do not use #GTK_WINDOW_POPUP just to turn off
the window borders; use gtk_window_set_decorated() for that.

C - `GtkWindowType`

- GTK_WINDOW_TOPLEVEL = %!s(int=0)
- GTK_WINDOW_POPUP = %!s(int=1)
## `WrapMode`

Describes a type of line wrapping.

C - `GtkWrapMode`

- GTK_WRAP_NONE = %!s(int=0)
- GTK_WRAP_CHAR = %!s(int=1)
- GTK_WRAP_WORD = %!s(int=2)
- GTK_WRAP_WORD_CHAR = %!s(int=3)
