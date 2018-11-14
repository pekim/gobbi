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

C - `GtkAlign`

### GTK_ALIGN_FILL = 0
stretch to fill all space if possible, center if
    no meaningful way to stretch

### GTK_ALIGN_START = 1
snap to left or top side, leaving space on right
    or bottom

### GTK_ALIGN_END = 2
snap to right or bottom side, leaving space on left
    or top

### GTK_ALIGN_CENTER = 3
center natural width of widget inside the
    allocation

### GTK_ALIGN_BASELINE = 4
align the widget according to the baseline. Since 3.10.


## `ArrowPlacement`

Used to specify the placement of scroll arrows in scrolling menus.

C - `GtkArrowPlacement`

### GTK_ARROWS_BOTH = 0
Place one arrow on each end of the menu.

### GTK_ARROWS_START = 1
Place both arrows at the top of the menu.

### GTK_ARROWS_END = 2
Place both arrows at the bottom of the menu.


## `ArrowType`

Used to indicate the direction in which an arrow should point.

C - `GtkArrowType`

### GTK_ARROW_UP = 0
Represents an upward pointing arrow.

### GTK_ARROW_DOWN = 1
Represents a downward pointing arrow.

### GTK_ARROW_LEFT = 2
Represents a left pointing arrow.

### GTK_ARROW_RIGHT = 3
Represents a right pointing arrow.

### GTK_ARROW_NONE = 4
No arrow. Since 2.10.


## `AssistantPageType`

An enum for determining the page role inside the #GtkAssistant. It's
used to handle buttons sensitivity and visibility.

Note that an assistant needs to end its page flow with a page of type
%GTK_ASSISTANT_PAGE_CONFIRM, %GTK_ASSISTANT_PAGE_SUMMARY or
%GTK_ASSISTANT_PAGE_PROGRESS to be correct.

The Cancel button will only be shown if the page isn’t “committed”.
See gtk_assistant_commit() for details.

C - `GtkAssistantPageType`

### GTK_ASSISTANT_PAGE_CONTENT = 0
The page has regular contents. Both the
 Back and forward buttons will be shown.

### GTK_ASSISTANT_PAGE_INTRO = 1
The page contains an introduction to the
 assistant task. Only the Forward button will be shown if there is a
  next page.

### GTK_ASSISTANT_PAGE_CONFIRM = 2
The page lets the user confirm or deny the
 changes. The Back and Apply buttons will be shown.

### GTK_ASSISTANT_PAGE_SUMMARY = 3
The page informs the user of the changes
 done. Only the Close button will be shown.

### GTK_ASSISTANT_PAGE_PROGRESS = 4
Used for tasks that take a long time to
 complete, blocks the assistant until the page is marked as complete.
  Only the back button will be shown.

### GTK_ASSISTANT_PAGE_CUSTOM = 5
Used for when other page types are not
 appropriate. No buttons will be shown, and the application must
 add its own buttons through gtk_assistant_add_action_widget().


## `BaselinePosition`

Whenever a container has some form of natural row it may align
children in that row along a common typographical baseline. If
the amount of verical space in the row is taller than the total
requested height of the baseline-aligned children then it can use a
&num;GtkBaselinePosition to select where to put the baseline inside the
extra availible space.

C - `GtkBaselinePosition`

### GTK_BASELINE_POSITION_TOP = 0
Align the baseline at the top

### GTK_BASELINE_POSITION_CENTER = 1
Center the baseline

### GTK_BASELINE_POSITION_BOTTOM = 2
Align the baseline at the bottom


## `BorderStyle`

Describes how the border of a UI element should be rendered.

C - `GtkBorderStyle`

### GTK_BORDER_STYLE_NONE = 0
No visible border

### GTK_BORDER_STYLE_SOLID = 1
A single line segment

### GTK_BORDER_STYLE_INSET = 2
Looks as if the content is sunken into the canvas

### GTK_BORDER_STYLE_OUTSET = 3
Looks as if the content is coming out of the canvas

### GTK_BORDER_STYLE_HIDDEN = 4
Same as @GTK_BORDER_STYLE_NONE

### GTK_BORDER_STYLE_DOTTED = 5
A series of round dots

### GTK_BORDER_STYLE_DASHED = 6
A series of square-ended dashes

### GTK_BORDER_STYLE_DOUBLE = 7
Two parallel lines with some space between them

### GTK_BORDER_STYLE_GROOVE = 8
Looks as if it were carved in the canvas

### GTK_BORDER_STYLE_RIDGE = 9
Looks as if it were coming out of the canvas


## `BuilderError`

Error codes that identify various errors that can occur while using
&num;GtkBuilder.

C - `GtkBuilderError`

### GTK_BUILDER_ERROR_INVALID_TYPE_FUNCTION = 0
A type-func attribute didn’t name
 a function that returns a #GType.

### GTK_BUILDER_ERROR_UNHANDLED_TAG = 1
The input contained a tag that #GtkBuilder
 can’t handle.

### GTK_BUILDER_ERROR_MISSING_ATTRIBUTE = 2
An attribute that is required by
 #GtkBuilder was missing.

### GTK_BUILDER_ERROR_INVALID_ATTRIBUTE = 3
#GtkBuilder found an attribute that
 it doesn’t understand.

### GTK_BUILDER_ERROR_INVALID_TAG = 4
#GtkBuilder found a tag that
 it doesn’t understand.

### GTK_BUILDER_ERROR_MISSING_PROPERTY_VALUE = 5
A required property value was
 missing.

### GTK_BUILDER_ERROR_INVALID_VALUE = 6
#GtkBuilder couldn’t parse
 some attribute value.

### GTK_BUILDER_ERROR_VERSION_MISMATCH = 7
The input file requires a newer version
 of GTK+.

### GTK_BUILDER_ERROR_DUPLICATE_ID = 8
An object id occurred twice.

### GTK_BUILDER_ERROR_OBJECT_TYPE_REFUSED = 9
A specified object type is of the same type or
 derived from the type of the composite class being extended with builder XML.

### GTK_BUILDER_ERROR_TEMPLATE_MISMATCH = 10
The wrong type was specified in a composite class’s template XML

### GTK_BUILDER_ERROR_INVALID_PROPERTY = 11
The specified property is unknown for the object class.

### GTK_BUILDER_ERROR_INVALID_SIGNAL = 12
The specified signal is unknown for the object class.

### GTK_BUILDER_ERROR_INVALID_ID = 13
An object id is unknown


## `ButtonBoxStyle`

Used to dictate the style that a #GtkButtonBox uses to layout the buttons it
contains.

C - `GtkButtonBoxStyle`

### GTK_BUTTONBOX_SPREAD = 1
Buttons are evenly spread across the box.

### GTK_BUTTONBOX_EDGE = 2
Buttons are placed at the edges of the box.

### GTK_BUTTONBOX_START = 3
Buttons are grouped towards the start of the box,
  (on the left for a HBox, or the top for a VBox).

### GTK_BUTTONBOX_END = 4
Buttons are grouped towards the end of the box,
  (on the right for a HBox, or the bottom for a VBox).

### GTK_BUTTONBOX_CENTER = 5
Buttons are centered in the box. Since 2.12.

### GTK_BUTTONBOX_EXPAND = 6
Buttons expand to fill the box. This entails giving
  buttons a "linked" appearance, making button sizes homogeneous, and
  setting spacing to 0 (same as calling gtk_box_set_homogeneous() and
  gtk_box_set_spacing() manually). Since 3.12.


## `ButtonRole`

The role specifies the desired appearance of a #GtkModelButton.

C - `GtkButtonRole`

### GTK_BUTTON_ROLE_NORMAL = 0
A plain button

### GTK_BUTTON_ROLE_CHECK = 1
A check button

### GTK_BUTTON_ROLE_RADIO = 2
A radio button


## `ButtonsType`

Prebuilt sets of buttons for the dialog. If
none of these choices are appropriate, simply use %GTK_BUTTONS_NONE
then call gtk_dialog_add_buttons().

> Please note that %GTK_BUTTONS_OK, %GTK_BUTTONS_YES_NO
> and %GTK_BUTTONS_OK_CANCEL are discouraged by the
> [GNOME Human Interface Guidelines](http://library.gnome.org/devel/hig-book/stable/).

C - `GtkButtonsType`

### GTK_BUTTONS_NONE = 0
no buttons at all

### GTK_BUTTONS_OK = 1
an OK button

### GTK_BUTTONS_CLOSE = 2
a Close button

### GTK_BUTTONS_CANCEL = 3
a Cancel button

### GTK_BUTTONS_YES_NO = 4
Yes and No buttons

### GTK_BUTTONS_OK_CANCEL = 5
OK and Cancel buttons


## `CellRendererAccelMode`

Determines if the edited accelerators are GTK+ accelerators. If
they are, consumed modifiers are suppressed, only accelerators
accepted by GTK+ are allowed, and the accelerators are rendered
in the same way as they are in menus.

C - `GtkCellRendererAccelMode`

### GTK_CELL_RENDERER_ACCEL_MODE_GTK = 0
GTK+ accelerators mode

### GTK_CELL_RENDERER_ACCEL_MODE_OTHER = 1
Other accelerator mode
GTK_CELL_RENDERER_ACCEL_MODE_MODIFIER_TAP: Bare modifiers mode

### GTK_CELL_RENDERER_ACCEL_MODE_MODIFIER_TAP = 2



## `CellRendererMode`

Identifies how the user can interact with a particular cell.

C - `GtkCellRendererMode`

### GTK_CELL_RENDERER_MODE_INERT = 0
The cell is just for display
 and cannot be interacted with.  Note that this doesn’t mean that eg. the
 row being drawn can’t be selected -- just that a particular element of
 it cannot be individually modified.

### GTK_CELL_RENDERER_MODE_ACTIVATABLE = 1
The cell can be clicked.

### GTK_CELL_RENDERER_MODE_EDITABLE = 2
The cell can be edited or otherwise modified.


## `CornerType`

Specifies which corner a child widget should be placed in when packed into
a #GtkScrolledWindow. This is effectively the opposite of where the scroll
bars are placed.

C - `GtkCornerType`

### GTK_CORNER_TOP_LEFT = 0
Place the scrollbars on the right and bottom of the
 widget (default behaviour).

### GTK_CORNER_BOTTOM_LEFT = 1
Place the scrollbars on the top and right of the
 widget.

### GTK_CORNER_TOP_RIGHT = 2
Place the scrollbars on the left and bottom of the
 widget.

### GTK_CORNER_BOTTOM_RIGHT = 3
Place the scrollbars on the top and left of the
 widget.


## `CssProviderError`

Error codes for %GTK_CSS_PROVIDER_ERROR.

C - `GtkCssProviderError`

### GTK_CSS_PROVIDER_ERROR_FAILED = 0
Failed.

### GTK_CSS_PROVIDER_ERROR_SYNTAX = 1
Syntax error.

### GTK_CSS_PROVIDER_ERROR_IMPORT = 2
Import error.

### GTK_CSS_PROVIDER_ERROR_NAME = 3
Name error.

### GTK_CSS_PROVIDER_ERROR_DEPRECATED = 4
Deprecation error.

### GTK_CSS_PROVIDER_ERROR_UNKNOWN_VALUE = 5
Unknown value.


## `CssSectionType`

The different types of sections indicate parts of a CSS document as
parsed by GTK’s CSS parser. They are oriented towards the
[CSS Grammar](http://www.w3.org/TR/CSS21/grammar.html),
but may contain extensions.

More types might be added in the future as the parser incorporates
more features.

C - `GtkCssSectionType`

### GTK_CSS_SECTION_DOCUMENT = 0
The section describes a complete document.
  This section time is the only one where gtk_css_section_get_parent()
  might return %NULL.

### GTK_CSS_SECTION_IMPORT = 1
The section defines an import rule.

### GTK_CSS_SECTION_COLOR_DEFINITION = 2
The section defines a color. This
  is a GTK extension to CSS.

### GTK_CSS_SECTION_BINDING_SET = 3
The section defines a binding set. This
  is a GTK extension to CSS.

### GTK_CSS_SECTION_RULESET = 4
The section defines a CSS ruleset.

### GTK_CSS_SECTION_SELECTOR = 5
The section defines a CSS selector.

### GTK_CSS_SECTION_DECLARATION = 6
The section defines the declaration of
  a CSS variable.

### GTK_CSS_SECTION_VALUE = 7
The section defines the value of a CSS declaration.

### GTK_CSS_SECTION_KEYFRAMES = 8
The section defines keyframes. See [CSS
  Animations](http://dev.w3.org/csswg/css3-animations/#keyframes) for details. Since 3.6


## `DeleteType`

See also: #GtkEntry::delete-from-cursor.

C - `GtkDeleteType`

### GTK_DELETE_CHARS = 0
Delete characters.

### GTK_DELETE_WORD_ENDS = 1
Delete only the portion of the word to the
  left/right of cursor if we’re in the middle of a word.

### GTK_DELETE_WORDS = 2
Delete words.

### GTK_DELETE_DISPLAY_LINES = 3
Delete display-lines. Display-lines
  refers to the visible lines, with respect to to the current line
  breaks. As opposed to paragraphs, which are defined by line
  breaks in the input.

### GTK_DELETE_DISPLAY_LINE_ENDS = 4
Delete only the portion of the
  display-line to the left/right of cursor.

### GTK_DELETE_PARAGRAPH_ENDS = 5
Delete to the end of the
  paragraph. Like C-k in Emacs (or its reverse).

### GTK_DELETE_PARAGRAPHS = 6
Delete entire line. Like C-k in pico.

### GTK_DELETE_WHITESPACE = 7
Delete only whitespace. Like M-\ in Emacs.


## `DirectionType`

Focus movement types.

C - `GtkDirectionType`

### GTK_DIR_TAB_FORWARD = 0
Move forward.

### GTK_DIR_TAB_BACKWARD = 1
Move backward.

### GTK_DIR_UP = 2
Move up.

### GTK_DIR_DOWN = 3
Move down.

### GTK_DIR_LEFT = 4
Move left.

### GTK_DIR_RIGHT = 5
Move right.


## `DragResult`

Gives an indication why a drag operation failed.
The value can by obtained by connecting to the
&num;GtkWidget::drag-failed signal.

C - `GtkDragResult`

### GTK_DRAG_RESULT_SUCCESS = 0
The drag operation was successful.

### GTK_DRAG_RESULT_NO_TARGET = 1
No suitable drag target.

### GTK_DRAG_RESULT_USER_CANCELLED = 2
The user cancelled the drag operation.

### GTK_DRAG_RESULT_TIMEOUT_EXPIRED = 3
The drag operation timed out.

### GTK_DRAG_RESULT_GRAB_BROKEN = 4
The pointer or keyboard grab used
 for the drag operation was broken.

### GTK_DRAG_RESULT_ERROR = 5
The drag operation failed due to some
 unspecified error.


## `EntryIconPosition`

Specifies the side of the entry at which an icon is placed.

C - `GtkEntryIconPosition`

### GTK_ENTRY_ICON_PRIMARY = 0
At the beginning of the entry (depending on the text direction).

### GTK_ENTRY_ICON_SECONDARY = 1
At the end of the entry (depending on the text direction).


## `EventSequenceState`

Describes the state of a #GdkEventSequence in a #GtkGesture.

C - `GtkEventSequenceState`

### GTK_EVENT_SEQUENCE_NONE = 0
The sequence is handled, but not grabbed.

### GTK_EVENT_SEQUENCE_CLAIMED = 1
The sequence is handled and grabbed.

### GTK_EVENT_SEQUENCE_DENIED = 2
The sequence is denied.


## `ExpanderStyle`

Used to specify the style of the expanders drawn by a #GtkTreeView.

C - `GtkExpanderStyle`

### GTK_EXPANDER_COLLAPSED = 0
The style used for a collapsed subtree.

### GTK_EXPANDER_SEMI_COLLAPSED = 1
Intermediate style used during animation.

### GTK_EXPANDER_SEMI_EXPANDED = 2
Intermediate style used during animation.

### GTK_EXPANDER_EXPANDED = 3
The style used for an expanded subtree.


## `FileChooserAction`

Describes whether a #GtkFileChooser is being used to open existing files
or to save to a possibly new file.

C - `GtkFileChooserAction`

### GTK_FILE_CHOOSER_ACTION_OPEN = 0
Indicates open mode.  The file chooser
 will only let the user pick an existing file.

### GTK_FILE_CHOOSER_ACTION_SAVE = 1
Indicates save mode.  The file chooser
 will let the user pick an existing file, or type in a new
 filename.

### GTK_FILE_CHOOSER_ACTION_SELECT_FOLDER = 2
Indicates an Open mode for
 selecting folders.  The file chooser will let the user pick an
 existing folder.

### GTK_FILE_CHOOSER_ACTION_CREATE_FOLDER = 3
Indicates a mode for creating a
 new folder.  The file chooser will let the user name an existing or
 new folder.


## `FileChooserConfirmation`

Used as a return value of handlers for the
&num;GtkFileChooser::confirm-overwrite signal of a #GtkFileChooser. This
value determines whether the file chooser will present the stock
confirmation dialog, accept the user’s choice of a filename, or
let the user choose another filename.

C - `GtkFileChooserConfirmation`

### GTK_FILE_CHOOSER_CONFIRMATION_CONFIRM = 0
The file chooser will present
 its stock dialog to confirm about overwriting an existing file.

### GTK_FILE_CHOOSER_CONFIRMATION_ACCEPT_FILENAME = 1
The file chooser will
 terminate and accept the user’s choice of a file name.

### GTK_FILE_CHOOSER_CONFIRMATION_SELECT_AGAIN = 2
The file chooser will
 continue running, so as to let the user select another file name.


## `FileChooserError`

These identify the various errors that can occur while calling
&num;GtkFileChooser functions.

C - `GtkFileChooserError`

### GTK_FILE_CHOOSER_ERROR_NONEXISTENT = 0
Indicates that a file does not exist.

### GTK_FILE_CHOOSER_ERROR_BAD_FILENAME = 1
Indicates a malformed filename.

### GTK_FILE_CHOOSER_ERROR_ALREADY_EXISTS = 2
Indicates a duplicate path (e.g. when
 adding a bookmark).

### GTK_FILE_CHOOSER_ERROR_INCOMPLETE_HOSTNAME = 3
Indicates an incomplete hostname (e.g. "http://foo" without a slash after that).


## `IMPreeditStyle`

Style for input method preedit. See also
&num;GtkSettings:gtk-im-preedit-style

C - `GtkIMPreeditStyle`

### GTK_IM_PREEDIT_NOTHING = 0
Deprecated

### GTK_IM_PREEDIT_CALLBACK = 1
Deprecated

### GTK_IM_PREEDIT_NONE = 2
Deprecated


## `IMStatusStyle`

Style for input method status. See also
&num;GtkSettings:gtk-im-status-style

C - `GtkIMStatusStyle`

### GTK_IM_STATUS_NOTHING = 0
Deprecated

### GTK_IM_STATUS_CALLBACK = 1
Deprecated

### GTK_IM_STATUS_NONE = 2
Deprecated


## `IconSize`

Built-in stock icon sizes.

C - `GtkIconSize`

### GTK_ICON_SIZE_INVALID = 0
Invalid size.

### GTK_ICON_SIZE_MENU = 1
Size appropriate for menus (16px).

### GTK_ICON_SIZE_SMALL_TOOLBAR = 2
Size appropriate for small toolbars (16px).

### GTK_ICON_SIZE_LARGE_TOOLBAR = 3
Size appropriate for large toolbars (24px)

### GTK_ICON_SIZE_BUTTON = 4
Size appropriate for buttons (16px)

### GTK_ICON_SIZE_DND = 5
Size appropriate for drag and drop (32px)

### GTK_ICON_SIZE_DIALOG = 6
Size appropriate for dialogs (48px)


## `IconThemeError`

Error codes for GtkIconTheme operations.

C - `GtkIconThemeError`

### GTK_ICON_THEME_NOT_FOUND = 0
The icon specified does not exist in the theme

### GTK_ICON_THEME_FAILED = 1
An unspecified error occurred.


## `IconViewDropPosition`

An enum for determining where a dropped item goes.

C - `GtkIconViewDropPosition`

### GTK_ICON_VIEW_NO_DROP = 0
no drop possible

### GTK_ICON_VIEW_DROP_INTO = 1
dropped item replaces the item

### GTK_ICON_VIEW_DROP_LEFT = 2
droppped item is inserted to the left

### GTK_ICON_VIEW_DROP_RIGHT = 3
dropped item is inserted to the right

### GTK_ICON_VIEW_DROP_ABOVE = 4
dropped item is inserted above

### GTK_ICON_VIEW_DROP_BELOW = 5
dropped item is inserted below


## `ImageType`

Describes the image data representation used by a #GtkImage. If you
want to get the image from the widget, you can only get the
currently-stored representation. e.g.  if the
gtk_image_get_storage_type() returns #GTK_IMAGE_PIXBUF, then you can
call gtk_image_get_pixbuf() but not gtk_image_get_stock().  For empty
images, you can request any storage type (call any of the "get"
functions), but they will all return %NULL values.

C - `GtkImageType`

### GTK_IMAGE_EMPTY = 0
there is no image displayed by the widget

### GTK_IMAGE_PIXBUF = 1
the widget contains a #GdkPixbuf

### GTK_IMAGE_STOCK = 2
the widget contains a [stock item name][gtkstock]

### GTK_IMAGE_ICON_SET = 3
the widget contains a #GtkIconSet

### GTK_IMAGE_ANIMATION = 4
the widget contains a #GdkPixbufAnimation

### GTK_IMAGE_ICON_NAME = 5
the widget contains a named icon.
 This image type was added in GTK+ 2.6

### GTK_IMAGE_GICON = 6
the widget contains a #GIcon.
 This image type was added in GTK+ 2.14

### GTK_IMAGE_SURFACE = 7
the widget contains a #cairo_surface_t.
 This image type was added in GTK+ 3.10


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

C - `GtkInputPurpose`

### GTK_INPUT_PURPOSE_FREE_FORM = 0
Allow any character

### GTK_INPUT_PURPOSE_ALPHA = 1
Allow only alphabetic characters

### GTK_INPUT_PURPOSE_DIGITS = 2
Allow only digits

### GTK_INPUT_PURPOSE_NUMBER = 3
Edited field expects numbers

### GTK_INPUT_PURPOSE_PHONE = 4
Edited field expects phone number

### GTK_INPUT_PURPOSE_URL = 5
Edited field expects URL

### GTK_INPUT_PURPOSE_EMAIL = 6
Edited field expects email address

### GTK_INPUT_PURPOSE_NAME = 7
Edited field expects the name of a person

### GTK_INPUT_PURPOSE_PASSWORD = 8
Like @GTK_INPUT_PURPOSE_FREE_FORM, but characters are hidden

### GTK_INPUT_PURPOSE_PIN = 9
Like @GTK_INPUT_PURPOSE_DIGITS, but characters are hidden


## `Justification`

Used for justifying the text inside a #GtkLabel widget. (See also
&num;GtkAlignment).

C - `GtkJustification`

### GTK_JUSTIFY_LEFT = 0
The text is placed at the left edge of the label.

### GTK_JUSTIFY_RIGHT = 1
The text is placed at the right edge of the label.

### GTK_JUSTIFY_CENTER = 2
The text is placed in the center of the label.

### GTK_JUSTIFY_FILL = 3
The text is placed is distributed across the label.


## `LevelBarMode`

Describes how #GtkLevelBar contents should be rendered.
Note that this enumeration could be extended with additional modes
in the future.

C - `GtkLevelBarMode`

### GTK_LEVEL_BAR_MODE_CONTINUOUS = 0
the bar has a continuous mode

### GTK_LEVEL_BAR_MODE_DISCRETE = 1
the bar has a discrete mode


## `License`

The type of license for an application.

This enumeration can be expanded at later date.

C - `GtkLicense`

### GTK_LICENSE_UNKNOWN = 0
No license specified

### GTK_LICENSE_CUSTOM = 1
A license text is going to be specified by the
  developer

### GTK_LICENSE_GPL_2_0 = 2
The GNU General Public License, version 2.0 or later

### GTK_LICENSE_GPL_3_0 = 3
The GNU General Public License, version 3.0 or later

### GTK_LICENSE_LGPL_2_1 = 4
The GNU Lesser General Public License, version 2.1 or later

### GTK_LICENSE_LGPL_3_0 = 5
The GNU Lesser General Public License, version 3.0 or later

### GTK_LICENSE_BSD = 6
The BSD standard license

### GTK_LICENSE_MIT_X11 = 7
The MIT/X11 standard license

### GTK_LICENSE_ARTISTIC = 8
The Artistic License, version 2.0

### GTK_LICENSE_GPL_2_0_ONLY = 9
The GNU General Public License, version 2.0 only. Since 3.12.

### GTK_LICENSE_GPL_3_0_ONLY = 10
The GNU General Public License, version 3.0 only. Since 3.12.

### GTK_LICENSE_LGPL_2_1_ONLY = 11
The GNU Lesser General Public License, version 2.1 only. Since 3.12.

### GTK_LICENSE_LGPL_3_0_ONLY = 12
The GNU Lesser General Public License, version 3.0 only. Since 3.12.

### GTK_LICENSE_AGPL_3_0 = 13
The GNU Affero General Public License, version 3.0 or later. Since: 3.22.

### GTK_LICENSE_AGPL_3_0_ONLY = 14
The GNU Affero General Public License, version 3.0 only. Since: 3.22.27.


## `MenuDirectionType`

An enumeration representing directional movements within a menu.

C - `GtkMenuDirectionType`

### GTK_MENU_DIR_PARENT = 0
To the parent menu shell

### GTK_MENU_DIR_CHILD = 1
To the submenu, if any, associated with the item

### GTK_MENU_DIR_NEXT = 2
To the next menu item

### GTK_MENU_DIR_PREV = 3
To the previous menu item


## `MessageType`

The type of message being displayed in the dialog.

C - `GtkMessageType`

### GTK_MESSAGE_INFO = 0
Informational message

### GTK_MESSAGE_WARNING = 1
Non-fatal warning message

### GTK_MESSAGE_QUESTION = 2
Question requiring a choice

### GTK_MESSAGE_ERROR = 3
Fatal error message

### GTK_MESSAGE_OTHER = 4
None of the above


## `MovementStep`



C - `GtkMovementStep`

### GTK_MOVEMENT_LOGICAL_POSITIONS = 0
Move forward or back by graphemes

### GTK_MOVEMENT_VISUAL_POSITIONS = 1
Move left or right by graphemes

### GTK_MOVEMENT_WORDS = 2
Move forward or back by words

### GTK_MOVEMENT_DISPLAY_LINES = 3
Move up or down lines (wrapped lines)

### GTK_MOVEMENT_DISPLAY_LINE_ENDS = 4
Move to either end of a line

### GTK_MOVEMENT_PARAGRAPHS = 5
Move up or down paragraphs (newline-ended lines)

### GTK_MOVEMENT_PARAGRAPH_ENDS = 6
Move to either end of a paragraph

### GTK_MOVEMENT_PAGES = 7
Move by pages

### GTK_MOVEMENT_BUFFER_ENDS = 8
Move to ends of the buffer

### GTK_MOVEMENT_HORIZONTAL_PAGES = 9
Move horizontally by pages


## `NotebookTab`



C - `GtkNotebookTab`

### GTK_NOTEBOOK_TAB_FIRST = 0


### GTK_NOTEBOOK_TAB_LAST = 1



## `NumberUpLayout`

Used to determine the layout of pages on a sheet when printing
multiple pages per sheet.

C - `GtkNumberUpLayout`

### GTK_NUMBER_UP_LAYOUT_LEFT_TO_RIGHT_TOP_TO_BOTTOM = 0
![](layout-lrtb.png)

### GTK_NUMBER_UP_LAYOUT_LEFT_TO_RIGHT_BOTTOM_TO_TOP = 1
![](layout-lrbt.png)

### GTK_NUMBER_UP_LAYOUT_RIGHT_TO_LEFT_TOP_TO_BOTTOM = 2
![](layout-rltb.png)

### GTK_NUMBER_UP_LAYOUT_RIGHT_TO_LEFT_BOTTOM_TO_TOP = 3
![](layout-rlbt.png)

### GTK_NUMBER_UP_LAYOUT_TOP_TO_BOTTOM_LEFT_TO_RIGHT = 4
![](layout-tblr.png)

### GTK_NUMBER_UP_LAYOUT_TOP_TO_BOTTOM_RIGHT_TO_LEFT = 5
![](layout-tbrl.png)

### GTK_NUMBER_UP_LAYOUT_BOTTOM_TO_TOP_LEFT_TO_RIGHT = 6
![](layout-btlr.png)

### GTK_NUMBER_UP_LAYOUT_BOTTOM_TO_TOP_RIGHT_TO_LEFT = 7
![](layout-btrl.png)


## `Orientation`

Represents the orientation of widgets and other objects which can be switched
between horizontal and vertical orientation on the fly, like #GtkToolbar or
&num;GtkGesturePan.

C - `GtkOrientation`

### GTK_ORIENTATION_HORIZONTAL = 0
The element is in horizontal orientation.

### GTK_ORIENTATION_VERTICAL = 1
The element is in vertical orientation.


## `PackDirection`

Determines how widgets should be packed inside menubars
and menuitems contained in menubars.

C - `GtkPackDirection`

### GTK_PACK_DIRECTION_LTR = 0
Widgets are packed left-to-right

### GTK_PACK_DIRECTION_RTL = 1
Widgets are packed right-to-left

### GTK_PACK_DIRECTION_TTB = 2
Widgets are packed top-to-bottom

### GTK_PACK_DIRECTION_BTT = 3
Widgets are packed bottom-to-top


## `PackType`

Represents the packing location #GtkBox children. (See: #GtkVBox,
&num;GtkHBox, and #GtkButtonBox).

C - `GtkPackType`

### GTK_PACK_START = 0
The child is packed into the start of the box

### GTK_PACK_END = 1
The child is packed into the end of the box


## `PadActionType`

The type of a pad action.

C - `GtkPadActionType`

### GTK_PAD_ACTION_BUTTON = 0
Action is triggered by a pad button

### GTK_PAD_ACTION_RING = 1
Action is triggered by a pad ring

### GTK_PAD_ACTION_STRIP = 2
Action is triggered by a pad strip


## `PageOrientation`

See also gtk_print_settings_set_orientation().

C - `GtkPageOrientation`

### GTK_PAGE_ORIENTATION_PORTRAIT = 0
Portrait mode.

### GTK_PAGE_ORIENTATION_LANDSCAPE = 1
Landscape mode.

### GTK_PAGE_ORIENTATION_REVERSE_PORTRAIT = 2
Reverse portrait mode.

### GTK_PAGE_ORIENTATION_REVERSE_LANDSCAPE = 3
Reverse landscape mode.


## `PageSet`

See also gtk_print_job_set_page_set().

C - `GtkPageSet`

### GTK_PAGE_SET_ALL = 0
All pages.

### GTK_PAGE_SET_EVEN = 1
Even pages.

### GTK_PAGE_SET_ODD = 2
Odd pages.


## `PanDirection`

Describes the panning direction of a #GtkGesturePan

C - `GtkPanDirection`

### GTK_PAN_DIRECTION_LEFT = 0
panned towards the left

### GTK_PAN_DIRECTION_RIGHT = 1
panned towards the right

### GTK_PAN_DIRECTION_UP = 2
panned upwards

### GTK_PAN_DIRECTION_DOWN = 3
panned downwards


## `PathPriorityType`

Priorities for path lookups.
See also gtk_binding_set_add_path().

C - `GtkPathPriorityType`

### GTK_PATH_PRIO_LOWEST = 0
Deprecated

### GTK_PATH_PRIO_GTK = 4
Deprecated

### GTK_PATH_PRIO_APPLICATION = 8
Deprecated

### GTK_PATH_PRIO_THEME = 10
Deprecated

### GTK_PATH_PRIO_RC = 12
Deprecated

### GTK_PATH_PRIO_HIGHEST = 15
Deprecated


## `PathType`

Widget path types.
See also gtk_binding_set_add_path().

C - `GtkPathType`

### GTK_PATH_WIDGET = 0
Deprecated

### GTK_PATH_WIDGET_CLASS = 1
Deprecated

### GTK_PATH_CLASS = 2
Deprecated


## `PolicyType`

Determines how the size should be computed to achieve the one of the
visibility mode for the scrollbars.

C - `GtkPolicyType`

### GTK_POLICY_ALWAYS = 0
The scrollbar is always visible. The view size is
 independent of the content.

### GTK_POLICY_AUTOMATIC = 1
The scrollbar will appear and disappear as necessary.
 For example, when all of a #GtkTreeView can not be seen.

### GTK_POLICY_NEVER = 2
The scrollbar should never appear. In this mode the
 content determines the size.

### GTK_POLICY_EXTERNAL = 3
Don't show a scrollbar, but don't force the
 size to follow the content. This can be used e.g. to make multiple
 scrolled windows share a scrollbar. Since: 3.16


## `PopoverConstraint`

Describes constraints to positioning of popovers. More values
may be added to this enumeration in the future.

C - `GtkPopoverConstraint`

### GTK_POPOVER_CONSTRAINT_NONE = 0
Don't constrain the popover position
  beyond what is imposed by the implementation

### GTK_POPOVER_CONSTRAINT_WINDOW = 1
Constrain the popover to the boundaries
  of the window that it is attached to


## `PositionType`

Describes which edge of a widget a certain feature is positioned at, e.g. the
tabs of a #GtkNotebook, the handle of a #GtkHandleBox or the label of a
&num;GtkScale.

C - `GtkPositionType`

### GTK_POS_LEFT = 0
The feature is at the left edge.

### GTK_POS_RIGHT = 1
The feature is at the right edge.

### GTK_POS_TOP = 2
The feature is at the top edge.

### GTK_POS_BOTTOM = 3
The feature is at the bottom edge.


## `PrintDuplex`

See also gtk_print_settings_set_duplex().

C - `GtkPrintDuplex`

### GTK_PRINT_DUPLEX_SIMPLEX = 0
No duplex.

### GTK_PRINT_DUPLEX_HORIZONTAL = 1
Horizontal duplex.

### GTK_PRINT_DUPLEX_VERTICAL = 2
Vertical duplex.


## `PrintError`

Error codes that identify various errors that can occur while
using the GTK+ printing support.

C - `GtkPrintError`

### GTK_PRINT_ERROR_GENERAL = 0
An unspecified error occurred.

### GTK_PRINT_ERROR_INTERNAL_ERROR = 1
An internal error occurred.

### GTK_PRINT_ERROR_NOMEM = 2
A memory allocation failed.

### GTK_PRINT_ERROR_INVALID_FILE = 3
An error occurred while loading a page setup
    or paper size from a key file.


## `PrintOperationAction`

The @action parameter to gtk_print_operation_run()
determines what action the print operation should perform.

C - `GtkPrintOperationAction`

### GTK_PRINT_OPERATION_ACTION_PRINT_DIALOG = 0
Show the print dialog.

### GTK_PRINT_OPERATION_ACTION_PRINT = 1
Start to print without showing
    the print dialog, based on the current print settings.

### GTK_PRINT_OPERATION_ACTION_PREVIEW = 2
Show the print preview.

### GTK_PRINT_OPERATION_ACTION_EXPORT = 3
Export to a file. This requires
    the export-filename property to be set.


## `PrintOperationResult`

A value of this type is returned by gtk_print_operation_run().

C - `GtkPrintOperationResult`

### GTK_PRINT_OPERATION_RESULT_ERROR = 0
An error has occurred.

### GTK_PRINT_OPERATION_RESULT_APPLY = 1
The print settings should be stored.

### GTK_PRINT_OPERATION_RESULT_CANCEL = 2
The print operation has been canceled,
    the print settings should not be stored.

### GTK_PRINT_OPERATION_RESULT_IN_PROGRESS = 3
The print operation is not complete
    yet. This value will only be returned when running asynchronously.


## `PrintPages`

See also gtk_print_job_set_pages()

C - `GtkPrintPages`

### GTK_PRINT_PAGES_ALL = 0
All pages.

### GTK_PRINT_PAGES_CURRENT = 1
Current page.

### GTK_PRINT_PAGES_RANGES = 2
Range of pages.

### GTK_PRINT_PAGES_SELECTION = 3
Selected pages.


## `PrintQuality`

See also gtk_print_settings_set_quality().

C - `GtkPrintQuality`

### GTK_PRINT_QUALITY_LOW = 0
Low quality.

### GTK_PRINT_QUALITY_NORMAL = 1
Normal quality.

### GTK_PRINT_QUALITY_HIGH = 2
High quality.

### GTK_PRINT_QUALITY_DRAFT = 3
Draft quality.


## `PrintStatus`

The status gives a rough indication of the completion of a running
print operation.

C - `GtkPrintStatus`

### GTK_PRINT_STATUS_INITIAL = 0
The printing has not started yet; this
    status is set initially, and while the print dialog is shown.

### GTK_PRINT_STATUS_PREPARING = 1
This status is set while the begin-print
    signal is emitted and during pagination.

### GTK_PRINT_STATUS_GENERATING_DATA = 2
This status is set while the
    pages are being rendered.

### GTK_PRINT_STATUS_SENDING_DATA = 3
The print job is being sent off to the
    printer.

### GTK_PRINT_STATUS_PENDING = 4
The print job has been sent to the printer,
    but is not printed for some reason, e.g. the printer may be stopped.

### GTK_PRINT_STATUS_PENDING_ISSUE = 5
Some problem has occurred during
    printing, e.g. a paper jam.

### GTK_PRINT_STATUS_PRINTING = 6
The printer is processing the print job.

### GTK_PRINT_STATUS_FINISHED = 7
The printing has been completed successfully.

### GTK_PRINT_STATUS_FINISHED_ABORTED = 8
The printing has been aborted.


## `PropagationPhase`

Describes the stage at which events are fed into a #GtkEventController.

C - `GtkPropagationPhase`

### GTK_PHASE_NONE = 0
Events are not delivered automatically. Those can be
  manually fed through gtk_event_controller_handle_event(). This should
  only be used when full control about when, or whether the controller
  handles the event is needed.

### GTK_PHASE_CAPTURE = 1
Events are delivered in the capture phase. The
  capture phase happens before the bubble phase, runs from the toplevel down
  to the event widget. This option should only be used on containers that
  might possibly handle events before their children do.

### GTK_PHASE_BUBBLE = 2
Events are delivered in the bubble phase. The bubble
  phase happens after the capture phase, and before the default handlers
  are run. This phase runs from the event widget, up to the toplevel.

### GTK_PHASE_TARGET = 3
Events are delivered in the default widget event handlers,
  note that widget implementations must chain up on button, motion, touch and
  grab broken handlers for controllers in this phase to be run.


## `RcTokenType`

The #GtkRcTokenType enumeration represents the tokens
in the RC file. It is exposed so that theme engines
can reuse these tokens when parsing the theme-engine
specific portions of a RC file.

C - `GtkRcTokenType`

### GTK_RC_TOKEN_INVALID = 270
Deprecated

### GTK_RC_TOKEN_INCLUDE = 271
Deprecated

### GTK_RC_TOKEN_NORMAL = 272
Deprecated

### GTK_RC_TOKEN_ACTIVE = 273
Deprecated

### GTK_RC_TOKEN_PRELIGHT = 274
Deprecated

### GTK_RC_TOKEN_SELECTED = 275
Deprecated

### GTK_RC_TOKEN_INSENSITIVE = 276
Deprecated

### GTK_RC_TOKEN_FG = 277
Deprecated

### GTK_RC_TOKEN_BG = 278
Deprecated

### GTK_RC_TOKEN_TEXT = 279
Deprecated

### GTK_RC_TOKEN_BASE = 280
Deprecated

### GTK_RC_TOKEN_XTHICKNESS = 281
Deprecated

### GTK_RC_TOKEN_YTHICKNESS = 282
Deprecated

### GTK_RC_TOKEN_FONT = 283
Deprecated

### GTK_RC_TOKEN_FONTSET = 284
Deprecated

### GTK_RC_TOKEN_FONT_NAME = 285
Deprecated

### GTK_RC_TOKEN_BG_PIXMAP = 286
Deprecated

### GTK_RC_TOKEN_PIXMAP_PATH = 287
Deprecated

### GTK_RC_TOKEN_STYLE = 288
Deprecated

### GTK_RC_TOKEN_BINDING = 289
Deprecated

### GTK_RC_TOKEN_BIND = 290
Deprecated

### GTK_RC_TOKEN_WIDGET = 291
Deprecated

### GTK_RC_TOKEN_WIDGET_CLASS = 292
Deprecated

### GTK_RC_TOKEN_CLASS = 293
Deprecated

### GTK_RC_TOKEN_LOWEST = 294
Deprecated

### GTK_RC_TOKEN_GTK = 295
Deprecated

### GTK_RC_TOKEN_APPLICATION = 296
Deprecated

### GTK_RC_TOKEN_THEME = 297
Deprecated

### GTK_RC_TOKEN_RC = 298
Deprecated

### GTK_RC_TOKEN_HIGHEST = 299
Deprecated

### GTK_RC_TOKEN_ENGINE = 300
Deprecated

### GTK_RC_TOKEN_MODULE_PATH = 301
Deprecated

### GTK_RC_TOKEN_IM_MODULE_PATH = 302
Deprecated

### GTK_RC_TOKEN_IM_MODULE_FILE = 303
Deprecated

### GTK_RC_TOKEN_STOCK = 304
Deprecated

### GTK_RC_TOKEN_LTR = 305
Deprecated

### GTK_RC_TOKEN_RTL = 306
Deprecated

### GTK_RC_TOKEN_COLOR = 307
Deprecated

### GTK_RC_TOKEN_UNBIND = 308
Deprecated

### GTK_RC_TOKEN_LAST = 309
Deprecated


## `RecentChooserError`

These identify the various errors that can occur while calling
&num;GtkRecentChooser functions.

C - `GtkRecentChooserError`

### GTK_RECENT_CHOOSER_ERROR_NOT_FOUND = 0
Indicates that a file does not exist

### GTK_RECENT_CHOOSER_ERROR_INVALID_URI = 1
Indicates a malformed URI


## `RecentManagerError`

Error codes for #GtkRecentManager operations

C - `GtkRecentManagerError`

### GTK_RECENT_MANAGER_ERROR_NOT_FOUND = 0
the URI specified does not exists in
  the recently used resources list.

### GTK_RECENT_MANAGER_ERROR_INVALID_URI = 1
the URI specified is not valid.

### GTK_RECENT_MANAGER_ERROR_INVALID_ENCODING = 2
the supplied string is not
  UTF-8 encoded.

### GTK_RECENT_MANAGER_ERROR_NOT_REGISTERED = 3
no application has registered
  the specified item.

### GTK_RECENT_MANAGER_ERROR_READ = 4
failure while reading the recently used
  resources file.

### GTK_RECENT_MANAGER_ERROR_WRITE = 5
failure while writing the recently used
  resources file.

### GTK_RECENT_MANAGER_ERROR_UNKNOWN = 6
unspecified error.


## `RecentSortType`

Used to specify the sorting method to be applyed to the recently
used resource list.

C - `GtkRecentSortType`

### GTK_RECENT_SORT_NONE = 0
Do not sort the returned list of recently used
  resources.

### GTK_RECENT_SORT_MRU = 1
Sort the returned list with the most recently used
  items first.

### GTK_RECENT_SORT_LRU = 2
Sort the returned list with the least recently used
  items first.

### GTK_RECENT_SORT_CUSTOM = 3
Sort the returned list using a custom sorting
  function passed using gtk_recent_chooser_set_sort_func().


## `ReliefStyle`

Indicated the relief to be drawn around a #GtkButton.

C - `GtkReliefStyle`

### GTK_RELIEF_NORMAL = 0
Draw a normal relief.

### GTK_RELIEF_HALF = 1
A half relief. Deprecated in 3.14, does the same as @GTK_RELIEF_NORMAL

### GTK_RELIEF_NONE = 2
No relief.


## `ResizeMode`



C - `GtkResizeMode`

### GTK_RESIZE_PARENT = 0
Pass resize request to the parent

### GTK_RESIZE_QUEUE = 1
Queue resizes on this widget

### GTK_RESIZE_IMMEDIATE = 2
Resize immediately. Deprecated.


## `ResponseType`

Predefined values for use as response ids in gtk_dialog_add_button().
All predefined values are negative; GTK+ leaves values of 0 or greater for
application-defined response ids.

C - `GtkResponseType`

### GTK_RESPONSE_NONE = -1
Returned if an action widget has no response id,
    or if the dialog gets programmatically hidden or destroyed

### GTK_RESPONSE_REJECT = -2
Generic response id, not used by GTK+ dialogs

### GTK_RESPONSE_ACCEPT = -3
Generic response id, not used by GTK+ dialogs

### GTK_RESPONSE_DELETE_EVENT = -4
Returned if the dialog is deleted

### GTK_RESPONSE_OK = -5
Returned by OK buttons in GTK+ dialogs

### GTK_RESPONSE_CANCEL = -6
Returned by Cancel buttons in GTK+ dialogs

### GTK_RESPONSE_CLOSE = -7
Returned by Close buttons in GTK+ dialogs

### GTK_RESPONSE_YES = -8
Returned by Yes buttons in GTK+ dialogs

### GTK_RESPONSE_NO = -9
Returned by No buttons in GTK+ dialogs

### GTK_RESPONSE_APPLY = -10
Returned by Apply buttons in GTK+ dialogs

### GTK_RESPONSE_HELP = -11
Returned by Help buttons in GTK+ dialogs


## `RevealerTransitionType`

These enumeration values describe the possible transitions
when the child of a #GtkRevealer widget is shown or hidden.

C - `GtkRevealerTransitionType`

### GTK_REVEALER_TRANSITION_TYPE_NONE = 0
No transition

### GTK_REVEALER_TRANSITION_TYPE_CROSSFADE = 1
Fade in

### GTK_REVEALER_TRANSITION_TYPE_SLIDE_RIGHT = 2
Slide in from the left

### GTK_REVEALER_TRANSITION_TYPE_SLIDE_LEFT = 3
Slide in from the right

### GTK_REVEALER_TRANSITION_TYPE_SLIDE_UP = 4
Slide in from the bottom

### GTK_REVEALER_TRANSITION_TYPE_SLIDE_DOWN = 5
Slide in from the top


## `ScrollStep`



C - `GtkScrollStep`

### GTK_SCROLL_STEPS = 0
Scroll in steps.

### GTK_SCROLL_PAGES = 1
Scroll by pages.

### GTK_SCROLL_ENDS = 2
Scroll to ends.

### GTK_SCROLL_HORIZONTAL_STEPS = 3
Scroll in horizontal steps.

### GTK_SCROLL_HORIZONTAL_PAGES = 4
Scroll by horizontal pages.

### GTK_SCROLL_HORIZONTAL_ENDS = 5
Scroll to the horizontal ends.


## `ScrollType`

Scrolling types.

C - `GtkScrollType`

### GTK_SCROLL_NONE = 0
No scrolling.

### GTK_SCROLL_JUMP = 1
Jump to new location.

### GTK_SCROLL_STEP_BACKWARD = 2
Step backward.

### GTK_SCROLL_STEP_FORWARD = 3
Step forward.

### GTK_SCROLL_PAGE_BACKWARD = 4
Page backward.

### GTK_SCROLL_PAGE_FORWARD = 5
Page forward.

### GTK_SCROLL_STEP_UP = 6
Step up.

### GTK_SCROLL_STEP_DOWN = 7
Step down.

### GTK_SCROLL_PAGE_UP = 8
Page up.

### GTK_SCROLL_PAGE_DOWN = 9
Page down.

### GTK_SCROLL_STEP_LEFT = 10
Step to the left.

### GTK_SCROLL_STEP_RIGHT = 11
Step to the right.

### GTK_SCROLL_PAGE_LEFT = 12
Page to the left.

### GTK_SCROLL_PAGE_RIGHT = 13
Page to the right.

### GTK_SCROLL_START = 14
Scroll to start.

### GTK_SCROLL_END = 15
Scroll to end.


## `ScrollablePolicy`

Defines the policy to be used in a scrollable widget when updating
the scrolled window adjustments in a given orientation.

C - `GtkScrollablePolicy`

### GTK_SCROLL_MINIMUM = 0
Scrollable adjustments are based on the minimum size

### GTK_SCROLL_NATURAL = 1
Scrollable adjustments are based on the natural size


## `SelectionMode`

Used to control what selections users are allowed to make.

C - `GtkSelectionMode`

### GTK_SELECTION_NONE = 0
No selection is possible.

### GTK_SELECTION_SINGLE = 1
Zero or one element may be selected.

### GTK_SELECTION_BROWSE = 2
Exactly one element is selected.
    In some circumstances, such as initially or during a search
    operation, it’s possible for no element to be selected with
    %GTK_SELECTION_BROWSE. What is really enforced is that the user
    can’t deselect a currently selected element except by selecting
    another element.

### GTK_SELECTION_MULTIPLE = 3
Any number of elements may be selected.
     The Ctrl key may be used to enlarge the selection, and Shift
     key to select between the focus and the child pointed to.
     Some widgets may also allow Click-drag to select a range of elements.


## `SensitivityType`

Determines how GTK+ handles the sensitivity of stepper arrows
at the end of range widgets.

C - `GtkSensitivityType`

### GTK_SENSITIVITY_AUTO = 0
The arrow is made insensitive if the
  thumb is at the end

### GTK_SENSITIVITY_ON = 1
The arrow is always sensitive

### GTK_SENSITIVITY_OFF = 2
The arrow is always insensitive


## `ShadowType`

Used to change the appearance of an outline typically provided by a #GtkFrame.

Note that many themes do not differentiate the appearance of the
various shadow types: Either their is no visible shadow (@GTK_SHADOW_NONE),
or there is (any other value).

C - `GtkShadowType`

### GTK_SHADOW_NONE = 0
No outline.

### GTK_SHADOW_IN = 1
The outline is bevelled inwards.

### GTK_SHADOW_OUT = 2
The outline is bevelled outwards like a button.

### GTK_SHADOW_ETCHED_IN = 3
The outline has a sunken 3d appearance.

### GTK_SHADOW_ETCHED_OUT = 4
The outline has a raised 3d appearance.


## `ShortcutType`

GtkShortcutType specifies the kind of shortcut that is being described.
More values may be added to this enumeration over time.

C - `GtkShortcutType`

### GTK_SHORTCUT_ACCELERATOR = 0
The shortcut is a keyboard accelerator. The #GtkShortcutsShortcut:accelerator
  property will be used.

### GTK_SHORTCUT_GESTURE_PINCH = 1
The shortcut is a pinch gesture. GTK+ provides an icon and subtitle.

### GTK_SHORTCUT_GESTURE_STRETCH = 2
The shortcut is a stretch gesture. GTK+ provides an icon and subtitle.

### GTK_SHORTCUT_GESTURE_ROTATE_CLOCKWISE = 3
The shortcut is a clockwise rotation gesture. GTK+ provides an icon and subtitle.

### GTK_SHORTCUT_GESTURE_ROTATE_COUNTERCLOCKWISE = 4
The shortcut is a counterclockwise rotation gesture. GTK+ provides an icon and subtitle.

### GTK_SHORTCUT_GESTURE_TWO_FINGER_SWIPE_LEFT = 5
The shortcut is a two-finger swipe gesture. GTK+ provides an icon and subtitle.

### GTK_SHORTCUT_GESTURE_TWO_FINGER_SWIPE_RIGHT = 6
The shortcut is a two-finger swipe gesture. GTK+ provides an icon and subtitle.

### GTK_SHORTCUT_GESTURE = 7
The shortcut is a gesture. The #GtkShortcutsShortcut:icon property will be
  used.


## `SizeGroupMode`

The mode of the size group determines the directions in which the size
group affects the requested sizes of its component widgets.

C - `GtkSizeGroupMode`

### GTK_SIZE_GROUP_NONE = 0
group has no effect

### GTK_SIZE_GROUP_HORIZONTAL = 1
group affects horizontal requisition

### GTK_SIZE_GROUP_VERTICAL = 2
group affects vertical requisition

### GTK_SIZE_GROUP_BOTH = 3
group affects both horizontal and vertical requisition


## `SizeRequestMode`

Specifies a preference for height-for-width or
width-for-height geometry management.

C - `GtkSizeRequestMode`

### GTK_SIZE_REQUEST_HEIGHT_FOR_WIDTH = 0
Prefer height-for-width geometry management

### GTK_SIZE_REQUEST_WIDTH_FOR_HEIGHT = 1
Prefer width-for-height geometry management

### GTK_SIZE_REQUEST_CONSTANT_SIZE = 2
Don’t trade height-for-width or width-for-height


## `SortType`

Determines the direction of a sort.

C - `GtkSortType`

### GTK_SORT_ASCENDING = 0
Sorting is in ascending order.

### GTK_SORT_DESCENDING = 1
Sorting is in descending order.


## `SpinButtonUpdatePolicy`

The spin button update policy determines whether the spin button displays
values even if they are outside the bounds of its adjustment.
See gtk_spin_button_set_update_policy().

C - `GtkSpinButtonUpdatePolicy`

### GTK_UPDATE_ALWAYS = 0
When refreshing your #GtkSpinButton, the value is
    always displayed

### GTK_UPDATE_IF_VALID = 1
When refreshing your #GtkSpinButton, the value is
    only displayed if it is valid within the bounds of the spin button's
    adjustment


## `SpinType`

The values of the GtkSpinType enumeration are used to specify the
change to make in gtk_spin_button_spin().

C - `GtkSpinType`

### GTK_SPIN_STEP_FORWARD = 0
Increment by the adjustments step increment.

### GTK_SPIN_STEP_BACKWARD = 1
Decrement by the adjustments step increment.

### GTK_SPIN_PAGE_FORWARD = 2
Increment by the adjustments page increment.

### GTK_SPIN_PAGE_BACKWARD = 3
Decrement by the adjustments page increment.

### GTK_SPIN_HOME = 4
Go to the adjustments lower bound.

### GTK_SPIN_END = 5
Go to the adjustments upper bound.

### GTK_SPIN_USER_DEFINED = 6
Change by a specified amount.


## `StackTransitionType`

These enumeration values describe the possible transitions
between pages in a #GtkStack widget.

New values may be added to this enumeration over time.

C - `GtkStackTransitionType`

### GTK_STACK_TRANSITION_TYPE_NONE = 0
No transition

### GTK_STACK_TRANSITION_TYPE_CROSSFADE = 1
A cross-fade

### GTK_STACK_TRANSITION_TYPE_SLIDE_RIGHT = 2
Slide from left to right

### GTK_STACK_TRANSITION_TYPE_SLIDE_LEFT = 3
Slide from right to left

### GTK_STACK_TRANSITION_TYPE_SLIDE_UP = 4
Slide from bottom up

### GTK_STACK_TRANSITION_TYPE_SLIDE_DOWN = 5
Slide from top down

### GTK_STACK_TRANSITION_TYPE_SLIDE_LEFT_RIGHT = 6
Slide from left or right according to the children order

### GTK_STACK_TRANSITION_TYPE_SLIDE_UP_DOWN = 7
Slide from top down or bottom up according to the order

### GTK_STACK_TRANSITION_TYPE_OVER_UP = 8
Cover the old page by sliding up. Since 3.12

### GTK_STACK_TRANSITION_TYPE_OVER_DOWN = 9
Cover the old page by sliding down. Since: 3.12

### GTK_STACK_TRANSITION_TYPE_OVER_LEFT = 10
Cover the old page by sliding to the left. Since: 3.12

### GTK_STACK_TRANSITION_TYPE_OVER_RIGHT = 11
Cover the old page by sliding to the right. Since: 3.12

### GTK_STACK_TRANSITION_TYPE_UNDER_UP = 12
Uncover the new page by sliding up. Since 3.12

### GTK_STACK_TRANSITION_TYPE_UNDER_DOWN = 13
Uncover the new page by sliding down. Since: 3.12

### GTK_STACK_TRANSITION_TYPE_UNDER_LEFT = 14
Uncover the new page by sliding to the left. Since: 3.12

### GTK_STACK_TRANSITION_TYPE_UNDER_RIGHT = 15
Uncover the new page by sliding to the right. Since: 3.12

### GTK_STACK_TRANSITION_TYPE_OVER_UP_DOWN = 16
Cover the old page sliding up or uncover the new page sliding down, according to order. Since: 3.12

### GTK_STACK_TRANSITION_TYPE_OVER_DOWN_UP = 17
Cover the old page sliding down or uncover the new page sliding up, according to order. Since: 3.14

### GTK_STACK_TRANSITION_TYPE_OVER_LEFT_RIGHT = 18
Cover the old page sliding left or uncover the new page sliding right, according to order. Since: 3.14

### GTK_STACK_TRANSITION_TYPE_OVER_RIGHT_LEFT = 19
Cover the old page sliding right or uncover the new page sliding left, according to order. Since: 3.14


## `StateType`

This type indicates the current state of a widget; the state determines how
the widget is drawn. The #GtkStateType enumeration is also used to
identify different colors in a #GtkStyle for drawing, so states can be
used for subparts of a widget as well as entire widgets.

C - `GtkStateType`

### GTK_STATE_NORMAL = 0
State during normal operation.

### GTK_STATE_ACTIVE = 1
State of a currently active widget, such as a depressed button.

### GTK_STATE_PRELIGHT = 2
State indicating that the mouse pointer is over
                     the widget and the widget will respond to mouse clicks.

### GTK_STATE_SELECTED = 3
State of a selected item, such the selected row in a list.

### GTK_STATE_INSENSITIVE = 4
State indicating that the widget is
                        unresponsive to user actions.

### GTK_STATE_INCONSISTENT = 5
The widget is inconsistent, such as checkbuttons
                         or radiobuttons that aren’t either set to %TRUE nor %FALSE,
                         or buttons requiring the user attention.

### GTK_STATE_FOCUSED = 6
The widget has the keyboard focus.


## `TextBufferTargetInfo`

These values are used as “info” for the targets contained in the
lists returned by gtk_text_buffer_get_copy_target_list() and
gtk_text_buffer_get_paste_target_list().

The values counts down from `-1` to avoid clashes
with application added drag destinations which usually start at 0.

C - `GtkTextBufferTargetInfo`

### GTK_TEXT_BUFFER_TARGET_INFO_BUFFER_CONTENTS = -1
Buffer contents

### GTK_TEXT_BUFFER_TARGET_INFO_RICH_TEXT = -2
Rich text

### GTK_TEXT_BUFFER_TARGET_INFO_TEXT = -3
Text


## `TextDirection`

Reading directions for text.

C - `GtkTextDirection`

### GTK_TEXT_DIR_NONE = 0
No direction.

### GTK_TEXT_DIR_LTR = 1
Left to right text direction.

### GTK_TEXT_DIR_RTL = 2
Right to left text direction.


## `TextExtendSelection`

Granularity types that extend the text selection. Use the
&num;GtkTextView::extend-selection signal to customize the selection.

C - `GtkTextExtendSelection`

### GTK_TEXT_EXTEND_SELECTION_WORD = 0
Selects the current word. It is triggered by
  a double-click for example.

### GTK_TEXT_EXTEND_SELECTION_LINE = 1
Selects the current line. It is triggered by
  a triple-click for example.


## `TextViewLayer`

Used to reference the layers of #GtkTextView for the purpose of customized
drawing with the ::draw_layer vfunc.

C - `GtkTextViewLayer`

### GTK_TEXT_VIEW_LAYER_BELOW = 0
Old deprecated layer, use %GTK_TEXT_VIEW_LAYER_BELOW_TEXT instead

### GTK_TEXT_VIEW_LAYER_ABOVE = 1
Old deprecated layer, use %GTK_TEXT_VIEW_LAYER_ABOVE_TEXT instead

### GTK_TEXT_VIEW_LAYER_BELOW_TEXT = 2
The layer rendered below the text (but above the background).  Since: 3.20

### GTK_TEXT_VIEW_LAYER_ABOVE_TEXT = 3
The layer rendered above the text.  Since: 3.20


## `TextWindowType`

Used to reference the parts of #GtkTextView.

C - `GtkTextWindowType`

### GTK_TEXT_WINDOW_PRIVATE = 0
Invalid value, used as a marker

### GTK_TEXT_WINDOW_WIDGET = 1
Window that floats over scrolling areas.

### GTK_TEXT_WINDOW_TEXT = 2
Scrollable text window.

### GTK_TEXT_WINDOW_LEFT = 3
Left side border window.

### GTK_TEXT_WINDOW_RIGHT = 4
Right side border window.

### GTK_TEXT_WINDOW_TOP = 5
Top border window.

### GTK_TEXT_WINDOW_BOTTOM = 6
Bottom border window.


## `ToolbarSpaceStyle`

Whether spacers are vertical lines or just blank.

C - `GtkToolbarSpaceStyle`

### GTK_TOOLBAR_SPACE_EMPTY = 0
Use blank spacers.

### GTK_TOOLBAR_SPACE_LINE = 1
Use vertical lines for spacers.


## `ToolbarStyle`

Used to customize the appearance of a #GtkToolbar. Note that
setting the toolbar style overrides the user’s preferences
for the default toolbar style.  Note that if the button has only
a label set and GTK_TOOLBAR_ICONS is used, the label will be
visible, and vice versa.

C - `GtkToolbarStyle`

### GTK_TOOLBAR_ICONS = 0
Buttons display only icons in the toolbar.

### GTK_TOOLBAR_TEXT = 1
Buttons display only text labels in the toolbar.

### GTK_TOOLBAR_BOTH = 2
Buttons display text and icons in the toolbar.

### GTK_TOOLBAR_BOTH_HORIZ = 3
Buttons display icons and text alongside each
 other, rather than vertically stacked


## `TreeViewColumnSizing`

The sizing method the column uses to determine its width.  Please note
that @GTK_TREE_VIEW_COLUMN_AUTOSIZE are inefficient for large views, and
can make columns appear choppy.

C - `GtkTreeViewColumnSizing`

### GTK_TREE_VIEW_COLUMN_GROW_ONLY = 0
Columns only get bigger in reaction to changes in the model

### GTK_TREE_VIEW_COLUMN_AUTOSIZE = 1
Columns resize to be the optimal size everytime the model changes.

### GTK_TREE_VIEW_COLUMN_FIXED = 2
Columns are a fixed numbers of pixels wide.


## `TreeViewDropPosition`

An enum for determining where a dropped row goes.

C - `GtkTreeViewDropPosition`

### GTK_TREE_VIEW_DROP_BEFORE = 0
dropped row is inserted before

### GTK_TREE_VIEW_DROP_AFTER = 1
dropped row is inserted after

### GTK_TREE_VIEW_DROP_INTO_OR_BEFORE = 2
dropped row becomes a child or is inserted before

### GTK_TREE_VIEW_DROP_INTO_OR_AFTER = 3
dropped row becomes a child or is inserted after


## `TreeViewGridLines`

Used to indicate which grid lines to draw in a tree view.

C - `GtkTreeViewGridLines`

### GTK_TREE_VIEW_GRID_LINES_NONE = 0
No grid lines.

### GTK_TREE_VIEW_GRID_LINES_HORIZONTAL = 1
Horizontal grid lines.

### GTK_TREE_VIEW_GRID_LINES_VERTICAL = 2
Vertical grid lines.

### GTK_TREE_VIEW_GRID_LINES_BOTH = 3
Horizontal and vertical grid lines.


## `Unit`

See also gtk_print_settings_set_paper_width().

C - `GtkUnit`

### GTK_UNIT_NONE = 0
No units.

### GTK_UNIT_POINTS = 1
Dimensions in points.

### GTK_UNIT_INCH = 2
Dimensions in inches.

### GTK_UNIT_MM = 3
Dimensions in millimeters


## `WidgetHelpType`

Kinds of widget-specific help. Used by the ::show-help signal.

C - `GtkWidgetHelpType`

### GTK_WIDGET_HELP_TOOLTIP = 0
Tooltip.

### GTK_WIDGET_HELP_WHATS_THIS = 1
What’s this.


## `WindowPosition`

Window placement can be influenced using this enumeration. Note that
using #GTK_WIN_POS_CENTER_ALWAYS is almost always a bad idea.
It won’t necessarily work well with all window managers or on all windowing systems.

C - `GtkWindowPosition`

### GTK_WIN_POS_NONE = 0
No influence is made on placement.

### GTK_WIN_POS_CENTER = 1
Windows should be placed in the center of the screen.

### GTK_WIN_POS_MOUSE = 2
Windows should be placed at the current mouse position.

### GTK_WIN_POS_CENTER_ALWAYS = 3
Keep window centered as it changes size, etc.

### GTK_WIN_POS_CENTER_ON_PARENT = 4
Center the window on its transient
 parent (see gtk_window_set_transient_for()).


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

### GTK_WINDOW_TOPLEVEL = 0
A regular window, such as a dialog.

### GTK_WINDOW_POPUP = 1
A special window such as a tooltip.


## `WrapMode`

Describes a type of line wrapping.

C - `GtkWrapMode`

### GTK_WRAP_NONE = 0
do not wrap lines; just make the text area wider

### GTK_WRAP_CHAR = 1
wrap text, breaking lines anywhere the cursor can
    appear (between characters, usually - if you want to be technical,
    between graphemes, see pango_get_log_attrs())

### GTK_WRAP_WORD = 2
wrap text, breaking lines in between words

### GTK_WRAP_WORD_CHAR = 3
wrap text, breaking lines in between words, or if
    that is not enough, also between graphemes


