// This is a generated file - DO NOT EDIT

package gtk

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

// Controls how a widget deals with extra space in a single (x or y)
// dimension.
//
// Alignment only matters if the widget receives a “too large” allocation,
// for example if you packed the widget with the #GtkWidget:expand
// flag inside a #GtkBox, then the widget might get extra space.  If
// you have for example a 16x16 icon inside a 32x32 space, the icon
// could be scaled and stretched, it could be centered, or it could be
// positioned to one side of the space.
//
// Note that in horizontal context @GTK_ALIGN_START and @GTK_ALIGN_END
// are interpreted relative to text direction.
//
// GTK_ALIGN_BASELINE support for it is optional for containers and widgets, and
// it is only supported for vertical alignment.  When its not supported by
// a child or a container it is treated as @GTK_ALIGN_FILL.
type Align C.GtkAlign

const (
	// stretch to fill all space if possible, center if
	// no meaningful way to stretch
	GTK_ALIGN_FILL Align = 0

	// snap to left or top side, leaving space on right
	// or bottom
	GTK_ALIGN_START Align = 1

	// snap to right or bottom side, leaving space on left
	// or top
	GTK_ALIGN_END Align = 2

	// center natural width of widget inside the
	// allocation
	GTK_ALIGN_CENTER Align = 3

	// align the widget according to the baseline. Since 3.10.
	GTK_ALIGN_BASELINE Align = 4
)

// Used to specify the placement of scroll arrows in scrolling menus.
type ArrowPlacement C.GtkArrowPlacement

const (
	// Place one arrow on each end of the menu.
	GTK_ARROWS_BOTH ArrowPlacement = 0

	// Place both arrows at the top of the menu.
	GTK_ARROWS_START ArrowPlacement = 1

	// Place both arrows at the bottom of the menu.
	GTK_ARROWS_END ArrowPlacement = 2
)

// Used to indicate the direction in which an arrow should point.
type ArrowType C.GtkArrowType

const (
	// Represents an upward pointing arrow.
	GTK_ARROW_UP ArrowType = 0

	// Represents a downward pointing arrow.
	GTK_ARROW_DOWN ArrowType = 1

	// Represents a left pointing arrow.
	GTK_ARROW_LEFT ArrowType = 2

	// Represents a right pointing arrow.
	GTK_ARROW_RIGHT ArrowType = 3

	// No arrow. Since 2.10.
	GTK_ARROW_NONE ArrowType = 4
)

// An enum for determining the page role inside the #GtkAssistant. It's
// used to handle buttons sensitivity and visibility.
//
// Note that an assistant needs to end its page flow with a page of type
// %GTK_ASSISTANT_PAGE_CONFIRM, %GTK_ASSISTANT_PAGE_SUMMARY or
// %GTK_ASSISTANT_PAGE_PROGRESS to be correct.
//
// The Cancel button will only be shown if the page isn’t “committed”.
// See gtk_assistant_commit() for details.
type AssistantPageType C.GtkAssistantPageType

const (
	// The page has regular contents. Both the
	// Back and forward buttons will be shown.
	GTK_ASSISTANT_PAGE_CONTENT AssistantPageType = 0

	// The page contains an introduction to the
	// assistant task. Only the Forward button will be shown if there is a
	// next page.
	GTK_ASSISTANT_PAGE_INTRO AssistantPageType = 1

	// The page lets the user confirm or deny the
	// changes. The Back and Apply buttons will be shown.
	GTK_ASSISTANT_PAGE_CONFIRM AssistantPageType = 2

	// The page informs the user of the changes
	// done. Only the Close button will be shown.
	GTK_ASSISTANT_PAGE_SUMMARY AssistantPageType = 3

	// Used for tasks that take a long time to
	// complete, blocks the assistant until the page is marked as complete.
	// Only the back button will be shown.
	GTK_ASSISTANT_PAGE_PROGRESS AssistantPageType = 4

	// Used for when other page types are not
	// appropriate. No buttons will be shown, and the application must
	// add its own buttons through gtk_assistant_add_action_widget().
	GTK_ASSISTANT_PAGE_CUSTOM AssistantPageType = 5
)

// Describes how the border of a UI element should be rendered.
type BorderStyle C.GtkBorderStyle

const (
	// No visible border
	GTK_BORDER_STYLE_NONE BorderStyle = 0

	// A single line segment
	GTK_BORDER_STYLE_SOLID BorderStyle = 1

	// Looks as if the content is sunken into the canvas
	GTK_BORDER_STYLE_INSET BorderStyle = 2

	// Looks as if the content is coming out of the canvas
	GTK_BORDER_STYLE_OUTSET BorderStyle = 3

	// Same as @GTK_BORDER_STYLE_NONE
	GTK_BORDER_STYLE_HIDDEN BorderStyle = 4

	// A series of round dots
	GTK_BORDER_STYLE_DOTTED BorderStyle = 5

	// A series of square-ended dashes
	GTK_BORDER_STYLE_DASHED BorderStyle = 6

	// Two parallel lines with some space between them
	GTK_BORDER_STYLE_DOUBLE BorderStyle = 7

	// Looks as if it were carved in the canvas
	GTK_BORDER_STYLE_GROOVE BorderStyle = 8

	// Looks as if it were coming out of the canvas
	GTK_BORDER_STYLE_RIDGE BorderStyle = 9
)

// Error codes that identify various errors that can occur while using
// #GtkBuilder.
type BuilderError C.GtkBuilderError

const (
	// A type-func attribute didn’t name
	// a function that returns a #GType.
	GTK_BUILDER_ERROR_INVALID_TYPE_FUNCTION BuilderError = 0

	// The input contained a tag that #GtkBuilder
	// can’t handle.
	GTK_BUILDER_ERROR_UNHANDLED_TAG BuilderError = 1

	// An attribute that is required by
	// #GtkBuilder was missing.
	GTK_BUILDER_ERROR_MISSING_ATTRIBUTE BuilderError = 2

	// #GtkBuilder found an attribute that
	// it doesn’t understand.
	GTK_BUILDER_ERROR_INVALID_ATTRIBUTE BuilderError = 3

	// #GtkBuilder found a tag that
	// it doesn’t understand.
	GTK_BUILDER_ERROR_INVALID_TAG BuilderError = 4

	// A required property value was
	// missing.
	GTK_BUILDER_ERROR_MISSING_PROPERTY_VALUE BuilderError = 5

	// #GtkBuilder couldn’t parse
	// some attribute value.
	GTK_BUILDER_ERROR_INVALID_VALUE BuilderError = 6

	// The input file requires a newer version
	// of GTK+.
	GTK_BUILDER_ERROR_VERSION_MISMATCH BuilderError = 7

	// An object id occurred twice.
	GTK_BUILDER_ERROR_DUPLICATE_ID BuilderError = 8

	// A specified object type is of the same type or
	// derived from the type of the composite class being extended with builder XML.
	GTK_BUILDER_ERROR_OBJECT_TYPE_REFUSED BuilderError = 9

	// The wrong type was specified in a composite class’s template XML
	GTK_BUILDER_ERROR_TEMPLATE_MISMATCH BuilderError = 10

	// The specified property is unknown for the object class.
	GTK_BUILDER_ERROR_INVALID_PROPERTY BuilderError = 11

	// The specified signal is unknown for the object class.
	GTK_BUILDER_ERROR_INVALID_SIGNAL BuilderError = 12

	// An object id is unknown
	GTK_BUILDER_ERROR_INVALID_ID BuilderError = 13
)

// Used to dictate the style that a #GtkButtonBox uses to layout the buttons it
// contains.
type ButtonBoxStyle C.GtkButtonBoxStyle

const (
	// Buttons are evenly spread across the box.
	GTK_BUTTONBOX_SPREAD ButtonBoxStyle = 1

	// Buttons are placed at the edges of the box.
	GTK_BUTTONBOX_EDGE ButtonBoxStyle = 2

	// Buttons are grouped towards the start of the box,
	// (on the left for a HBox, or the top for a VBox).
	GTK_BUTTONBOX_START ButtonBoxStyle = 3

	// Buttons are grouped towards the end of the box,
	// (on the right for a HBox, or the bottom for a VBox).
	GTK_BUTTONBOX_END ButtonBoxStyle = 4

	// Buttons are centered in the box. Since 2.12.
	GTK_BUTTONBOX_CENTER ButtonBoxStyle = 5

	// Buttons expand to fill the box. This entails giving
	// buttons a "linked" appearance, making button sizes homogeneous, and
	// setting spacing to 0 (same as calling gtk_box_set_homogeneous() and
	// gtk_box_set_spacing() manually). Since 3.12.
	GTK_BUTTONBOX_EXPAND ButtonBoxStyle = 6
)

// The role specifies the desired appearance of a #GtkModelButton.
type ButtonRole C.GtkButtonRole

const (
	// A plain button
	GTK_BUTTON_ROLE_NORMAL ButtonRole = 0

	// A check button
	GTK_BUTTON_ROLE_CHECK ButtonRole = 1

	// A radio button
	GTK_BUTTON_ROLE_RADIO ButtonRole = 2
)

// Prebuilt sets of buttons for the dialog. If
// none of these choices are appropriate, simply use %GTK_BUTTONS_NONE
// then call gtk_dialog_add_buttons().
//
// > Please note that %GTK_BUTTONS_OK, %GTK_BUTTONS_YES_NO
// > and %GTK_BUTTONS_OK_CANCEL are discouraged by the
// > [GNOME Human Interface Guidelines](http://library.gnome.org/devel/hig-book/stable/).
type ButtonsType C.GtkButtonsType

const (
	// no buttons at all
	GTK_BUTTONS_NONE ButtonsType = 0

	// an OK button
	GTK_BUTTONS_OK ButtonsType = 1

	// a Close button
	GTK_BUTTONS_CLOSE ButtonsType = 2

	// a Cancel button
	GTK_BUTTONS_CANCEL ButtonsType = 3

	// Yes and No buttons
	GTK_BUTTONS_YES_NO ButtonsType = 4

	// OK and Cancel buttons
	GTK_BUTTONS_OK_CANCEL ButtonsType = 5
)

// Determines if the edited accelerators are GTK+ accelerators. If
// they are, consumed modifiers are suppressed, only accelerators
// accepted by GTK+ are allowed, and the accelerators are rendered
// in the same way as they are in menus.
type CellRendererAccelMode C.GtkCellRendererAccelMode

const (
	// GTK+ accelerators mode
	GTK_CELL_RENDERER_ACCEL_MODE_GTK CellRendererAccelMode = 0

	// Other accelerator mode
	// GTK_CELL_RENDERER_ACCEL_MODE_MODIFIER_TAP: Bare modifiers mode
	GTK_CELL_RENDERER_ACCEL_MODE_OTHER CellRendererAccelMode = 1

	GTK_CELL_RENDERER_ACCEL_MODE_MODIFIER_TAP CellRendererAccelMode = 2
)

// Identifies how the user can interact with a particular cell.
type CellRendererMode C.GtkCellRendererMode

const (
	// The cell is just for display
	// and cannot be interacted with.  Note that this doesn’t mean that eg. the
	// row being drawn can’t be selected -- just that a particular element of
	// it cannot be individually modified.
	GTK_CELL_RENDERER_MODE_INERT CellRendererMode = 0

	// The cell can be clicked.
	GTK_CELL_RENDERER_MODE_ACTIVATABLE CellRendererMode = 1

	// The cell can be edited or otherwise modified.
	GTK_CELL_RENDERER_MODE_EDITABLE CellRendererMode = 2
)

// Specifies which corner a child widget should be placed in when packed into
// a #GtkScrolledWindow. This is effectively the opposite of where the scroll
// bars are placed.
type CornerType C.GtkCornerType

const (
	// Place the scrollbars on the right and bottom of the
	// widget (default behaviour).
	GTK_CORNER_TOP_LEFT CornerType = 0

	// Place the scrollbars on the top and right of the
	// widget.
	GTK_CORNER_BOTTOM_LEFT CornerType = 1

	// Place the scrollbars on the left and bottom of the
	// widget.
	GTK_CORNER_TOP_RIGHT CornerType = 2

	// Place the scrollbars on the top and left of the
	// widget.
	GTK_CORNER_BOTTOM_RIGHT CornerType = 3
)

// Error codes for %GTK_CSS_PROVIDER_ERROR.
type CssProviderError C.GtkCssProviderError

const (
	// Failed.
	GTK_CSS_PROVIDER_ERROR_FAILED CssProviderError = 0

	// Syntax error.
	GTK_CSS_PROVIDER_ERROR_SYNTAX CssProviderError = 1

	// Import error.
	GTK_CSS_PROVIDER_ERROR_IMPORT CssProviderError = 2

	// Name error.
	GTK_CSS_PROVIDER_ERROR_NAME CssProviderError = 3

	// Deprecation error.
	GTK_CSS_PROVIDER_ERROR_DEPRECATED CssProviderError = 4

	// Unknown value.
	GTK_CSS_PROVIDER_ERROR_UNKNOWN_VALUE CssProviderError = 5
)

// See also: #GtkEntry::delete-from-cursor.
type DeleteType C.GtkDeleteType

const (
	// Delete characters.
	GTK_DELETE_CHARS DeleteType = 0

	// Delete only the portion of the word to the
	// left/right of cursor if we’re in the middle of a word.
	GTK_DELETE_WORD_ENDS DeleteType = 1

	// Delete words.
	GTK_DELETE_WORDS DeleteType = 2

	// Delete display-lines. Display-lines
	// refers to the visible lines, with respect to to the current line
	// breaks. As opposed to paragraphs, which are defined by line
	// breaks in the input.
	GTK_DELETE_DISPLAY_LINES DeleteType = 3

	// Delete only the portion of the
	// display-line to the left/right of cursor.
	GTK_DELETE_DISPLAY_LINE_ENDS DeleteType = 4

	// Delete to the end of the
	// paragraph. Like C-k in Emacs (or its reverse).
	GTK_DELETE_PARAGRAPH_ENDS DeleteType = 5

	// Delete entire line. Like C-k in pico.
	GTK_DELETE_PARAGRAPHS DeleteType = 6

	// Delete only whitespace. Like M-\ in Emacs.
	GTK_DELETE_WHITESPACE DeleteType = 7
)

// Focus movement types.
type DirectionType C.GtkDirectionType

const (
	// Move forward.
	GTK_DIR_TAB_FORWARD DirectionType = 0

	// Move backward.
	GTK_DIR_TAB_BACKWARD DirectionType = 1

	// Move up.
	GTK_DIR_UP DirectionType = 2

	// Move down.
	GTK_DIR_DOWN DirectionType = 3

	// Move left.
	GTK_DIR_LEFT DirectionType = 4

	// Move right.
	GTK_DIR_RIGHT DirectionType = 5
)

// Gives an indication why a drag operation failed.
// The value can by obtained by connecting to the
// #GtkWidget::drag-failed signal.
type DragResult C.GtkDragResult

const (
	// The drag operation was successful.
	GTK_DRAG_RESULT_SUCCESS DragResult = 0

	// No suitable drag target.
	GTK_DRAG_RESULT_NO_TARGET DragResult = 1

	// The user cancelled the drag operation.
	GTK_DRAG_RESULT_USER_CANCELLED DragResult = 2

	// The drag operation timed out.
	GTK_DRAG_RESULT_TIMEOUT_EXPIRED DragResult = 3

	// The pointer or keyboard grab used
	// for the drag operation was broken.
	GTK_DRAG_RESULT_GRAB_BROKEN DragResult = 4

	// The drag operation failed due to some
	// unspecified error.
	GTK_DRAG_RESULT_ERROR DragResult = 5
)

// Used to specify the style of the expanders drawn by a #GtkTreeView.
type ExpanderStyle C.GtkExpanderStyle

const (
	// The style used for a collapsed subtree.
	GTK_EXPANDER_COLLAPSED ExpanderStyle = 0

	// Intermediate style used during animation.
	GTK_EXPANDER_SEMI_COLLAPSED ExpanderStyle = 1

	// Intermediate style used during animation.
	GTK_EXPANDER_SEMI_EXPANDED ExpanderStyle = 2

	// The style used for an expanded subtree.
	GTK_EXPANDER_EXPANDED ExpanderStyle = 3
)

// Describes whether a #GtkFileChooser is being used to open existing files
// or to save to a possibly new file.
type FileChooserAction C.GtkFileChooserAction

const (
	// Indicates open mode.  The file chooser
	// will only let the user pick an existing file.
	GTK_FILE_CHOOSER_ACTION_OPEN FileChooserAction = 0

	// Indicates save mode.  The file chooser
	// will let the user pick an existing file, or type in a new
	// filename.
	GTK_FILE_CHOOSER_ACTION_SAVE FileChooserAction = 1

	// Indicates an Open mode for
	// selecting folders.  The file chooser will let the user pick an
	// existing folder.
	GTK_FILE_CHOOSER_ACTION_SELECT_FOLDER FileChooserAction = 2

	// Indicates a mode for creating a
	// new folder.  The file chooser will let the user name an existing or
	// new folder.
	GTK_FILE_CHOOSER_ACTION_CREATE_FOLDER FileChooserAction = 3
)

// These identify the various errors that can occur while calling
// #GtkFileChooser functions.
type FileChooserError C.GtkFileChooserError

const (
	// Indicates that a file does not exist.
	GTK_FILE_CHOOSER_ERROR_NONEXISTENT FileChooserError = 0

	// Indicates a malformed filename.
	GTK_FILE_CHOOSER_ERROR_BAD_FILENAME FileChooserError = 1

	// Indicates a duplicate path (e.g. when
	// adding a bookmark).
	GTK_FILE_CHOOSER_ERROR_ALREADY_EXISTS FileChooserError = 2

	// Indicates an incomplete hostname (e.g. "http://foo" without a slash after that).
	GTK_FILE_CHOOSER_ERROR_INCOMPLETE_HOSTNAME FileChooserError = 3
)

// Style for input method preedit. See also
// #GtkSettings:gtk-im-preedit-style
type IMPreeditStyle C.GtkIMPreeditStyle

const (
	// Deprecated
	GTK_IM_PREEDIT_NOTHING IMPreeditStyle = 0

	// Deprecated
	GTK_IM_PREEDIT_CALLBACK IMPreeditStyle = 1

	// Deprecated
	GTK_IM_PREEDIT_NONE IMPreeditStyle = 2
)

// Style for input method status. See also
// #GtkSettings:gtk-im-status-style
type IMStatusStyle C.GtkIMStatusStyle

const (
	// Deprecated
	GTK_IM_STATUS_NOTHING IMStatusStyle = 0

	// Deprecated
	GTK_IM_STATUS_CALLBACK IMStatusStyle = 1

	// Deprecated
	GTK_IM_STATUS_NONE IMStatusStyle = 2
)

// Built-in stock icon sizes.
type IconSize C.GtkIconSize

const (
	// Invalid size.
	GTK_ICON_SIZE_INVALID IconSize = 0

	// Size appropriate for menus (16px).
	GTK_ICON_SIZE_MENU IconSize = 1

	// Size appropriate for small toolbars (16px).
	GTK_ICON_SIZE_SMALL_TOOLBAR IconSize = 2

	// Size appropriate for large toolbars (24px)
	GTK_ICON_SIZE_LARGE_TOOLBAR IconSize = 3

	// Size appropriate for buttons (16px)
	GTK_ICON_SIZE_BUTTON IconSize = 4

	// Size appropriate for drag and drop (32px)
	GTK_ICON_SIZE_DND IconSize = 5

	// Size appropriate for dialogs (48px)
	GTK_ICON_SIZE_DIALOG IconSize = 6
)

// Error codes for GtkIconTheme operations.
type IconThemeError C.GtkIconThemeError

const (
	// The icon specified does not exist in the theme
	GTK_ICON_THEME_NOT_FOUND IconThemeError = 0

	// An unspecified error occurred.
	GTK_ICON_THEME_FAILED IconThemeError = 1
)

// An enum for determining where a dropped item goes.
type IconViewDropPosition C.GtkIconViewDropPosition

const (
	// no drop possible
	GTK_ICON_VIEW_NO_DROP IconViewDropPosition = 0

	// dropped item replaces the item
	GTK_ICON_VIEW_DROP_INTO IconViewDropPosition = 1

	// droppped item is inserted to the left
	GTK_ICON_VIEW_DROP_LEFT IconViewDropPosition = 2

	// dropped item is inserted to the right
	GTK_ICON_VIEW_DROP_RIGHT IconViewDropPosition = 3

	// dropped item is inserted above
	GTK_ICON_VIEW_DROP_ABOVE IconViewDropPosition = 4

	// dropped item is inserted below
	GTK_ICON_VIEW_DROP_BELOW IconViewDropPosition = 5
)

// Describes the image data representation used by a #GtkImage. If you
// want to get the image from the widget, you can only get the
// currently-stored representation. e.g.  if the
// gtk_image_get_storage_type() returns #GTK_IMAGE_PIXBUF, then you can
// call gtk_image_get_pixbuf() but not gtk_image_get_stock().  For empty
// images, you can request any storage type (call any of the "get"
// functions), but they will all return %NULL values.
type ImageType C.GtkImageType

const (
	// there is no image displayed by the widget
	GTK_IMAGE_EMPTY ImageType = 0

	// the widget contains a #GdkPixbuf
	GTK_IMAGE_PIXBUF ImageType = 1

	// the widget contains a [stock item name][gtkstock]
	GTK_IMAGE_STOCK ImageType = 2

	// the widget contains a #GtkIconSet
	GTK_IMAGE_ICON_SET ImageType = 3

	// the widget contains a #GdkPixbufAnimation
	GTK_IMAGE_ANIMATION ImageType = 4

	// the widget contains a named icon.
	// This image type was added in GTK+ 2.6
	GTK_IMAGE_ICON_NAME ImageType = 5

	// the widget contains a #GIcon.
	// This image type was added in GTK+ 2.14
	GTK_IMAGE_GICON ImageType = 6

	// the widget contains a #cairo_surface_t.
	// This image type was added in GTK+ 3.10
	GTK_IMAGE_SURFACE ImageType = 7
)

// Used for justifying the text inside a #GtkLabel widget. (See also
// #GtkAlignment).
type Justification C.GtkJustification

const (
	// The text is placed at the left edge of the label.
	GTK_JUSTIFY_LEFT Justification = 0

	// The text is placed at the right edge of the label.
	GTK_JUSTIFY_RIGHT Justification = 1

	// The text is placed in the center of the label.
	GTK_JUSTIFY_CENTER Justification = 2

	// The text is placed is distributed across the label.
	GTK_JUSTIFY_FILL Justification = 3
)

// An enumeration representing directional movements within a menu.
type MenuDirectionType C.GtkMenuDirectionType

const (
	// To the parent menu shell
	GTK_MENU_DIR_PARENT MenuDirectionType = 0

	// To the submenu, if any, associated with the item
	GTK_MENU_DIR_CHILD MenuDirectionType = 1

	// To the next menu item
	GTK_MENU_DIR_NEXT MenuDirectionType = 2

	// To the previous menu item
	GTK_MENU_DIR_PREV MenuDirectionType = 3
)

// The type of message being displayed in the dialog.
type MessageType C.GtkMessageType

const (
	// Informational message
	GTK_MESSAGE_INFO MessageType = 0

	// Non-fatal warning message
	GTK_MESSAGE_WARNING MessageType = 1

	// Question requiring a choice
	GTK_MESSAGE_QUESTION MessageType = 2

	// Fatal error message
	GTK_MESSAGE_ERROR MessageType = 3

	// None of the above
	GTK_MESSAGE_OTHER MessageType = 4
)

type MovementStep C.GtkMovementStep

const (
	// Move forward or back by graphemes
	GTK_MOVEMENT_LOGICAL_POSITIONS MovementStep = 0

	// Move left or right by graphemes
	GTK_MOVEMENT_VISUAL_POSITIONS MovementStep = 1

	// Move forward or back by words
	GTK_MOVEMENT_WORDS MovementStep = 2

	// Move up or down lines (wrapped lines)
	GTK_MOVEMENT_DISPLAY_LINES MovementStep = 3

	// Move to either end of a line
	GTK_MOVEMENT_DISPLAY_LINE_ENDS MovementStep = 4

	// Move up or down paragraphs (newline-ended lines)
	GTK_MOVEMENT_PARAGRAPHS MovementStep = 5

	// Move to either end of a paragraph
	GTK_MOVEMENT_PARAGRAPH_ENDS MovementStep = 6

	// Move by pages
	GTK_MOVEMENT_PAGES MovementStep = 7

	// Move to ends of the buffer
	GTK_MOVEMENT_BUFFER_ENDS MovementStep = 8

	// Move horizontally by pages
	GTK_MOVEMENT_HORIZONTAL_PAGES MovementStep = 9
)

type NotebookTab C.GtkNotebookTab

const (
	GTK_NOTEBOOK_TAB_FIRST NotebookTab = 0

	GTK_NOTEBOOK_TAB_LAST NotebookTab = 1
)

// Used to determine the layout of pages on a sheet when printing
// multiple pages per sheet.
type NumberUpLayout C.GtkNumberUpLayout

const (
	// ![](layout-lrtb.png)
	GTK_NUMBER_UP_LAYOUT_LEFT_TO_RIGHT_TOP_TO_BOTTOM NumberUpLayout = 0

	// ![](layout-lrbt.png)
	GTK_NUMBER_UP_LAYOUT_LEFT_TO_RIGHT_BOTTOM_TO_TOP NumberUpLayout = 1

	// ![](layout-rltb.png)
	GTK_NUMBER_UP_LAYOUT_RIGHT_TO_LEFT_TOP_TO_BOTTOM NumberUpLayout = 2

	// ![](layout-rlbt.png)
	GTK_NUMBER_UP_LAYOUT_RIGHT_TO_LEFT_BOTTOM_TO_TOP NumberUpLayout = 3

	// ![](layout-tblr.png)
	GTK_NUMBER_UP_LAYOUT_TOP_TO_BOTTOM_LEFT_TO_RIGHT NumberUpLayout = 4

	// ![](layout-tbrl.png)
	GTK_NUMBER_UP_LAYOUT_TOP_TO_BOTTOM_RIGHT_TO_LEFT NumberUpLayout = 5

	// ![](layout-btlr.png)
	GTK_NUMBER_UP_LAYOUT_BOTTOM_TO_TOP_LEFT_TO_RIGHT NumberUpLayout = 6

	// ![](layout-btrl.png)
	GTK_NUMBER_UP_LAYOUT_BOTTOM_TO_TOP_RIGHT_TO_LEFT NumberUpLayout = 7
)

// Represents the orientation of widgets and other objects which can be switched
// between horizontal and vertical orientation on the fly, like #GtkToolbar or
// #GtkGesturePan.
type Orientation C.GtkOrientation

const (
	// The element is in horizontal orientation.
	GTK_ORIENTATION_HORIZONTAL Orientation = 0

	// The element is in vertical orientation.
	GTK_ORIENTATION_VERTICAL Orientation = 1
)

// Determines how widgets should be packed inside menubars
// and menuitems contained in menubars.
type PackDirection C.GtkPackDirection

const (
	// Widgets are packed left-to-right
	GTK_PACK_DIRECTION_LTR PackDirection = 0

	// Widgets are packed right-to-left
	GTK_PACK_DIRECTION_RTL PackDirection = 1

	// Widgets are packed top-to-bottom
	GTK_PACK_DIRECTION_TTB PackDirection = 2

	// Widgets are packed bottom-to-top
	GTK_PACK_DIRECTION_BTT PackDirection = 3
)

// Represents the packing location #GtkBox children. (See: #GtkVBox,
// #GtkHBox, and #GtkButtonBox).
type PackType C.GtkPackType

const (
	// The child is packed into the start of the box
	GTK_PACK_START PackType = 0

	// The child is packed into the end of the box
	GTK_PACK_END PackType = 1
)

// See also gtk_print_settings_set_orientation().
type PageOrientation C.GtkPageOrientation

const (
	// Portrait mode.
	GTK_PAGE_ORIENTATION_PORTRAIT PageOrientation = 0

	// Landscape mode.
	GTK_PAGE_ORIENTATION_LANDSCAPE PageOrientation = 1

	// Reverse portrait mode.
	GTK_PAGE_ORIENTATION_REVERSE_PORTRAIT PageOrientation = 2

	// Reverse landscape mode.
	GTK_PAGE_ORIENTATION_REVERSE_LANDSCAPE PageOrientation = 3
)

// See also gtk_print_job_set_page_set().
type PageSet C.GtkPageSet

const (
	// All pages.
	GTK_PAGE_SET_ALL PageSet = 0

	// Even pages.
	GTK_PAGE_SET_EVEN PageSet = 1

	// Odd pages.
	GTK_PAGE_SET_ODD PageSet = 2
)

// Priorities for path lookups.
// See also gtk_binding_set_add_path().
type PathPriorityType C.GtkPathPriorityType

const (
	// Deprecated
	GTK_PATH_PRIO_LOWEST PathPriorityType = 0

	// Deprecated
	GTK_PATH_PRIO_GTK PathPriorityType = 4

	// Deprecated
	GTK_PATH_PRIO_APPLICATION PathPriorityType = 8

	// Deprecated
	GTK_PATH_PRIO_THEME PathPriorityType = 10

	// Deprecated
	GTK_PATH_PRIO_RC PathPriorityType = 12

	// Deprecated
	GTK_PATH_PRIO_HIGHEST PathPriorityType = 15
)

// Widget path types.
// See also gtk_binding_set_add_path().
type PathType C.GtkPathType

const (
	// Deprecated
	GTK_PATH_WIDGET PathType = 0

	// Deprecated
	GTK_PATH_WIDGET_CLASS PathType = 1

	// Deprecated
	GTK_PATH_CLASS PathType = 2
)

// Determines how the size should be computed to achieve the one of the
// visibility mode for the scrollbars.
type PolicyType C.GtkPolicyType

const (
	// The scrollbar is always visible. The view size is
	// independent of the content.
	GTK_POLICY_ALWAYS PolicyType = 0

	// The scrollbar will appear and disappear as necessary.
	// For example, when all of a #GtkTreeView can not be seen.
	GTK_POLICY_AUTOMATIC PolicyType = 1

	// The scrollbar should never appear. In this mode the
	// content determines the size.
	GTK_POLICY_NEVER PolicyType = 2

	// Don't show a scrollbar, but don't force the
	// size to follow the content. This can be used e.g. to make multiple
	// scrolled windows share a scrollbar. Since: 3.16
	GTK_POLICY_EXTERNAL PolicyType = 3
)

// Describes which edge of a widget a certain feature is positioned at, e.g. the
// tabs of a #GtkNotebook, the handle of a #GtkHandleBox or the label of a
// #GtkScale.
type PositionType C.GtkPositionType

const (
	// The feature is at the left edge.
	GTK_POS_LEFT PositionType = 0

	// The feature is at the right edge.
	GTK_POS_RIGHT PositionType = 1

	// The feature is at the top edge.
	GTK_POS_TOP PositionType = 2

	// The feature is at the bottom edge.
	GTK_POS_BOTTOM PositionType = 3
)

// See also gtk_print_settings_set_duplex().
type PrintDuplex C.GtkPrintDuplex

const (
	// No duplex.
	GTK_PRINT_DUPLEX_SIMPLEX PrintDuplex = 0

	// Horizontal duplex.
	GTK_PRINT_DUPLEX_HORIZONTAL PrintDuplex = 1

	// Vertical duplex.
	GTK_PRINT_DUPLEX_VERTICAL PrintDuplex = 2
)

// Error codes that identify various errors that can occur while
// using the GTK+ printing support.
type PrintError C.GtkPrintError

const (
	// An unspecified error occurred.
	GTK_PRINT_ERROR_GENERAL PrintError = 0

	// An internal error occurred.
	GTK_PRINT_ERROR_INTERNAL_ERROR PrintError = 1

	// A memory allocation failed.
	GTK_PRINT_ERROR_NOMEM PrintError = 2

	// An error occurred while loading a page setup
	// or paper size from a key file.
	GTK_PRINT_ERROR_INVALID_FILE PrintError = 3
)

// The @action parameter to gtk_print_operation_run()
// determines what action the print operation should perform.
type PrintOperationAction C.GtkPrintOperationAction

const (
	// Show the print dialog.
	GTK_PRINT_OPERATION_ACTION_PRINT_DIALOG PrintOperationAction = 0

	// Start to print without showing
	// the print dialog, based on the current print settings.
	GTK_PRINT_OPERATION_ACTION_PRINT PrintOperationAction = 1

	// Show the print preview.
	GTK_PRINT_OPERATION_ACTION_PREVIEW PrintOperationAction = 2

	// Export to a file. This requires
	// the export-filename property to be set.
	GTK_PRINT_OPERATION_ACTION_EXPORT PrintOperationAction = 3
)

// A value of this type is returned by gtk_print_operation_run().
type PrintOperationResult C.GtkPrintOperationResult

const (
	// An error has occurred.
	GTK_PRINT_OPERATION_RESULT_ERROR PrintOperationResult = 0

	// The print settings should be stored.
	GTK_PRINT_OPERATION_RESULT_APPLY PrintOperationResult = 1

	// The print operation has been canceled,
	// the print settings should not be stored.
	GTK_PRINT_OPERATION_RESULT_CANCEL PrintOperationResult = 2

	// The print operation is not complete
	// yet. This value will only be returned when running asynchronously.
	GTK_PRINT_OPERATION_RESULT_IN_PROGRESS PrintOperationResult = 3
)

// See also gtk_print_job_set_pages()
type PrintPages C.GtkPrintPages

const (
	// All pages.
	GTK_PRINT_PAGES_ALL PrintPages = 0

	// Current page.
	GTK_PRINT_PAGES_CURRENT PrintPages = 1

	// Range of pages.
	GTK_PRINT_PAGES_RANGES PrintPages = 2

	// Selected pages.
	GTK_PRINT_PAGES_SELECTION PrintPages = 3
)

// See also gtk_print_settings_set_quality().
type PrintQuality C.GtkPrintQuality

const (
	// Low quality.
	GTK_PRINT_QUALITY_LOW PrintQuality = 0

	// Normal quality.
	GTK_PRINT_QUALITY_NORMAL PrintQuality = 1

	// High quality.
	GTK_PRINT_QUALITY_HIGH PrintQuality = 2

	// Draft quality.
	GTK_PRINT_QUALITY_DRAFT PrintQuality = 3
)

// The status gives a rough indication of the completion of a running
// print operation.
type PrintStatus C.GtkPrintStatus

const (
	// The printing has not started yet; this
	// status is set initially, and while the print dialog is shown.
	GTK_PRINT_STATUS_INITIAL PrintStatus = 0

	// This status is set while the begin-print
	// signal is emitted and during pagination.
	GTK_PRINT_STATUS_PREPARING PrintStatus = 1

	// This status is set while the
	// pages are being rendered.
	GTK_PRINT_STATUS_GENERATING_DATA PrintStatus = 2

	// The print job is being sent off to the
	// printer.
	GTK_PRINT_STATUS_SENDING_DATA PrintStatus = 3

	// The print job has been sent to the printer,
	// but is not printed for some reason, e.g. the printer may be stopped.
	GTK_PRINT_STATUS_PENDING PrintStatus = 4

	// Some problem has occurred during
	// printing, e.g. a paper jam.
	GTK_PRINT_STATUS_PENDING_ISSUE PrintStatus = 5

	// The printer is processing the print job.
	GTK_PRINT_STATUS_PRINTING PrintStatus = 6

	// The printing has been completed successfully.
	GTK_PRINT_STATUS_FINISHED PrintStatus = 7

	// The printing has been aborted.
	GTK_PRINT_STATUS_FINISHED_ABORTED PrintStatus = 8
)

// The #GtkRcTokenType enumeration represents the tokens
// in the RC file. It is exposed so that theme engines
// can reuse these tokens when parsing the theme-engine
// specific portions of a RC file.
type RcTokenType C.GtkRcTokenType

const (
	// Deprecated
	GTK_RC_TOKEN_INVALID RcTokenType = 270

	// Deprecated
	GTK_RC_TOKEN_INCLUDE RcTokenType = 271

	// Deprecated
	GTK_RC_TOKEN_NORMAL RcTokenType = 272

	// Deprecated
	GTK_RC_TOKEN_ACTIVE RcTokenType = 273

	// Deprecated
	GTK_RC_TOKEN_PRELIGHT RcTokenType = 274

	// Deprecated
	GTK_RC_TOKEN_SELECTED RcTokenType = 275

	// Deprecated
	GTK_RC_TOKEN_INSENSITIVE RcTokenType = 276

	// Deprecated
	GTK_RC_TOKEN_FG RcTokenType = 277

	// Deprecated
	GTK_RC_TOKEN_BG RcTokenType = 278

	// Deprecated
	GTK_RC_TOKEN_TEXT RcTokenType = 279

	// Deprecated
	GTK_RC_TOKEN_BASE RcTokenType = 280

	// Deprecated
	GTK_RC_TOKEN_XTHICKNESS RcTokenType = 281

	// Deprecated
	GTK_RC_TOKEN_YTHICKNESS RcTokenType = 282

	// Deprecated
	GTK_RC_TOKEN_FONT RcTokenType = 283

	// Deprecated
	GTK_RC_TOKEN_FONTSET RcTokenType = 284

	// Deprecated
	GTK_RC_TOKEN_FONT_NAME RcTokenType = 285

	// Deprecated
	GTK_RC_TOKEN_BG_PIXMAP RcTokenType = 286

	// Deprecated
	GTK_RC_TOKEN_PIXMAP_PATH RcTokenType = 287

	// Deprecated
	GTK_RC_TOKEN_STYLE RcTokenType = 288

	// Deprecated
	GTK_RC_TOKEN_BINDING RcTokenType = 289

	// Deprecated
	GTK_RC_TOKEN_BIND RcTokenType = 290

	// Deprecated
	GTK_RC_TOKEN_WIDGET RcTokenType = 291

	// Deprecated
	GTK_RC_TOKEN_WIDGET_CLASS RcTokenType = 292

	// Deprecated
	GTK_RC_TOKEN_CLASS RcTokenType = 293

	// Deprecated
	GTK_RC_TOKEN_LOWEST RcTokenType = 294

	// Deprecated
	GTK_RC_TOKEN_GTK RcTokenType = 295

	// Deprecated
	GTK_RC_TOKEN_APPLICATION RcTokenType = 296

	// Deprecated
	GTK_RC_TOKEN_THEME RcTokenType = 297

	// Deprecated
	GTK_RC_TOKEN_RC RcTokenType = 298

	// Deprecated
	GTK_RC_TOKEN_HIGHEST RcTokenType = 299

	// Deprecated
	GTK_RC_TOKEN_ENGINE RcTokenType = 300

	// Deprecated
	GTK_RC_TOKEN_MODULE_PATH RcTokenType = 301

	// Deprecated
	GTK_RC_TOKEN_IM_MODULE_PATH RcTokenType = 302

	// Deprecated
	GTK_RC_TOKEN_IM_MODULE_FILE RcTokenType = 303

	// Deprecated
	GTK_RC_TOKEN_STOCK RcTokenType = 304

	// Deprecated
	GTK_RC_TOKEN_LTR RcTokenType = 305

	// Deprecated
	GTK_RC_TOKEN_RTL RcTokenType = 306

	// Deprecated
	GTK_RC_TOKEN_COLOR RcTokenType = 307

	// Deprecated
	GTK_RC_TOKEN_UNBIND RcTokenType = 308

	// Deprecated
	GTK_RC_TOKEN_LAST RcTokenType = 309
)

// Indicated the relief to be drawn around a #GtkButton.
type ReliefStyle C.GtkReliefStyle

const (
	// Draw a normal relief.
	GTK_RELIEF_NORMAL ReliefStyle = 0

	// A half relief. Deprecated in 3.14, does the same as @GTK_RELIEF_NORMAL
	GTK_RELIEF_HALF ReliefStyle = 1

	// No relief.
	GTK_RELIEF_NONE ReliefStyle = 2
)

type ResizeMode C.GtkResizeMode

const (
	// Pass resize request to the parent
	GTK_RESIZE_PARENT ResizeMode = 0

	// Queue resizes on this widget
	GTK_RESIZE_QUEUE ResizeMode = 1

	// Resize immediately. Deprecated.
	GTK_RESIZE_IMMEDIATE ResizeMode = 2
)

// Predefined values for use as response ids in gtk_dialog_add_button().
// All predefined values are negative; GTK+ leaves values of 0 or greater for
// application-defined response ids.
type ResponseType C.GtkResponseType

const (
	// Returned if an action widget has no response id,
	// or if the dialog gets programmatically hidden or destroyed
	GTK_RESPONSE_NONE ResponseType = -1

	// Generic response id, not used by GTK+ dialogs
	GTK_RESPONSE_REJECT ResponseType = -2

	// Generic response id, not used by GTK+ dialogs
	GTK_RESPONSE_ACCEPT ResponseType = -3

	// Returned if the dialog is deleted
	GTK_RESPONSE_DELETE_EVENT ResponseType = -4

	// Returned by OK buttons in GTK+ dialogs
	GTK_RESPONSE_OK ResponseType = -5

	// Returned by Cancel buttons in GTK+ dialogs
	GTK_RESPONSE_CANCEL ResponseType = -6

	// Returned by Close buttons in GTK+ dialogs
	GTK_RESPONSE_CLOSE ResponseType = -7

	// Returned by Yes buttons in GTK+ dialogs
	GTK_RESPONSE_YES ResponseType = -8

	// Returned by No buttons in GTK+ dialogs
	GTK_RESPONSE_NO ResponseType = -9

	// Returned by Apply buttons in GTK+ dialogs
	GTK_RESPONSE_APPLY ResponseType = -10

	// Returned by Help buttons in GTK+ dialogs
	GTK_RESPONSE_HELP ResponseType = -11
)

// These enumeration values describe the possible transitions
// when the child of a #GtkRevealer widget is shown or hidden.
type RevealerTransitionType C.GtkRevealerTransitionType

const (
	// No transition
	GTK_REVEALER_TRANSITION_TYPE_NONE RevealerTransitionType = 0

	// Fade in
	GTK_REVEALER_TRANSITION_TYPE_CROSSFADE RevealerTransitionType = 1

	// Slide in from the left
	GTK_REVEALER_TRANSITION_TYPE_SLIDE_RIGHT RevealerTransitionType = 2

	// Slide in from the right
	GTK_REVEALER_TRANSITION_TYPE_SLIDE_LEFT RevealerTransitionType = 3

	// Slide in from the bottom
	GTK_REVEALER_TRANSITION_TYPE_SLIDE_UP RevealerTransitionType = 4

	// Slide in from the top
	GTK_REVEALER_TRANSITION_TYPE_SLIDE_DOWN RevealerTransitionType = 5
)

type ScrollStep C.GtkScrollStep

const (
	// Scroll in steps.
	GTK_SCROLL_STEPS ScrollStep = 0

	// Scroll by pages.
	GTK_SCROLL_PAGES ScrollStep = 1

	// Scroll to ends.
	GTK_SCROLL_ENDS ScrollStep = 2

	// Scroll in horizontal steps.
	GTK_SCROLL_HORIZONTAL_STEPS ScrollStep = 3

	// Scroll by horizontal pages.
	GTK_SCROLL_HORIZONTAL_PAGES ScrollStep = 4

	// Scroll to the horizontal ends.
	GTK_SCROLL_HORIZONTAL_ENDS ScrollStep = 5
)

// Scrolling types.
type ScrollType C.GtkScrollType

const (
	// No scrolling.
	GTK_SCROLL_NONE ScrollType = 0

	// Jump to new location.
	GTK_SCROLL_JUMP ScrollType = 1

	// Step backward.
	GTK_SCROLL_STEP_BACKWARD ScrollType = 2

	// Step forward.
	GTK_SCROLL_STEP_FORWARD ScrollType = 3

	// Page backward.
	GTK_SCROLL_PAGE_BACKWARD ScrollType = 4

	// Page forward.
	GTK_SCROLL_PAGE_FORWARD ScrollType = 5

	// Step up.
	GTK_SCROLL_STEP_UP ScrollType = 6

	// Step down.
	GTK_SCROLL_STEP_DOWN ScrollType = 7

	// Page up.
	GTK_SCROLL_PAGE_UP ScrollType = 8

	// Page down.
	GTK_SCROLL_PAGE_DOWN ScrollType = 9

	// Step to the left.
	GTK_SCROLL_STEP_LEFT ScrollType = 10

	// Step to the right.
	GTK_SCROLL_STEP_RIGHT ScrollType = 11

	// Page to the left.
	GTK_SCROLL_PAGE_LEFT ScrollType = 12

	// Page to the right.
	GTK_SCROLL_PAGE_RIGHT ScrollType = 13

	// Scroll to start.
	GTK_SCROLL_START ScrollType = 14

	// Scroll to end.
	GTK_SCROLL_END ScrollType = 15
)

// Defines the policy to be used in a scrollable widget when updating
// the scrolled window adjustments in a given orientation.
type ScrollablePolicy C.GtkScrollablePolicy

const (
	// Scrollable adjustments are based on the minimum size
	GTK_SCROLL_MINIMUM ScrollablePolicy = 0

	// Scrollable adjustments are based on the natural size
	GTK_SCROLL_NATURAL ScrollablePolicy = 1
)

// Used to control what selections users are allowed to make.
type SelectionMode C.GtkSelectionMode

const (
	// No selection is possible.
	GTK_SELECTION_NONE SelectionMode = 0

	// Zero or one element may be selected.
	GTK_SELECTION_SINGLE SelectionMode = 1

	// Exactly one element is selected.
	// In some circumstances, such as initially or during a search
	// operation, it’s possible for no element to be selected with
	// %GTK_SELECTION_BROWSE. What is really enforced is that the user
	// can’t deselect a currently selected element except by selecting
	// another element.
	GTK_SELECTION_BROWSE SelectionMode = 2

	// Any number of elements may be selected.
	// The Ctrl key may be used to enlarge the selection, and Shift
	// key to select between the focus and the child pointed to.
	// Some widgets may also allow Click-drag to select a range of elements.
	GTK_SELECTION_MULTIPLE SelectionMode = 3
)

// Determines how GTK+ handles the sensitivity of stepper arrows
// at the end of range widgets.
type SensitivityType C.GtkSensitivityType

const (
	// The arrow is made insensitive if the
	// thumb is at the end
	GTK_SENSITIVITY_AUTO SensitivityType = 0

	// The arrow is always sensitive
	GTK_SENSITIVITY_ON SensitivityType = 1

	// The arrow is always insensitive
	GTK_SENSITIVITY_OFF SensitivityType = 2
)

// Used to change the appearance of an outline typically provided by a #GtkFrame.
//
// Note that many themes do not differentiate the appearance of the
// various shadow types: Either their is no visible shadow (@GTK_SHADOW_NONE),
// or there is (any other value).
type ShadowType C.GtkShadowType

const (
	// No outline.
	GTK_SHADOW_NONE ShadowType = 0

	// The outline is bevelled inwards.
	GTK_SHADOW_IN ShadowType = 1

	// The outline is bevelled outwards like a button.
	GTK_SHADOW_OUT ShadowType = 2

	// The outline has a sunken 3d appearance.
	GTK_SHADOW_ETCHED_IN ShadowType = 3

	// The outline has a raised 3d appearance.
	GTK_SHADOW_ETCHED_OUT ShadowType = 4
)

// The mode of the size group determines the directions in which the size
// group affects the requested sizes of its component widgets.
type SizeGroupMode C.GtkSizeGroupMode

const (
	// group has no effect
	GTK_SIZE_GROUP_NONE SizeGroupMode = 0

	// group affects horizontal requisition
	GTK_SIZE_GROUP_HORIZONTAL SizeGroupMode = 1

	// group affects vertical requisition
	GTK_SIZE_GROUP_VERTICAL SizeGroupMode = 2

	// group affects both horizontal and vertical requisition
	GTK_SIZE_GROUP_BOTH SizeGroupMode = 3
)

// Specifies a preference for height-for-width or
// width-for-height geometry management.
type SizeRequestMode C.GtkSizeRequestMode

const (
	// Prefer height-for-width geometry management
	GTK_SIZE_REQUEST_HEIGHT_FOR_WIDTH SizeRequestMode = 0

	// Prefer width-for-height geometry management
	GTK_SIZE_REQUEST_WIDTH_FOR_HEIGHT SizeRequestMode = 1

	// Don’t trade height-for-width or width-for-height
	GTK_SIZE_REQUEST_CONSTANT_SIZE SizeRequestMode = 2
)

// Determines the direction of a sort.
type SortType C.GtkSortType

const (
	// Sorting is in ascending order.
	GTK_SORT_ASCENDING SortType = 0

	// Sorting is in descending order.
	GTK_SORT_DESCENDING SortType = 1
)

// The spin button update policy determines whether the spin button displays
// values even if they are outside the bounds of its adjustment.
// See gtk_spin_button_set_update_policy().
type SpinButtonUpdatePolicy C.GtkSpinButtonUpdatePolicy

const (
	// When refreshing your #GtkSpinButton, the value is
	// always displayed
	GTK_UPDATE_ALWAYS SpinButtonUpdatePolicy = 0

	// When refreshing your #GtkSpinButton, the value is
	// only displayed if it is valid within the bounds of the spin button's
	// adjustment
	GTK_UPDATE_IF_VALID SpinButtonUpdatePolicy = 1
)

// The values of the GtkSpinType enumeration are used to specify the
// change to make in gtk_spin_button_spin().
type SpinType C.GtkSpinType

const (
	// Increment by the adjustments step increment.
	GTK_SPIN_STEP_FORWARD SpinType = 0

	// Decrement by the adjustments step increment.
	GTK_SPIN_STEP_BACKWARD SpinType = 1

	// Increment by the adjustments page increment.
	GTK_SPIN_PAGE_FORWARD SpinType = 2

	// Decrement by the adjustments page increment.
	GTK_SPIN_PAGE_BACKWARD SpinType = 3

	// Go to the adjustments lower bound.
	GTK_SPIN_HOME SpinType = 4

	// Go to the adjustments upper bound.
	GTK_SPIN_END SpinType = 5

	// Change by a specified amount.
	GTK_SPIN_USER_DEFINED SpinType = 6
)

// These enumeration values describe the possible transitions
// between pages in a #GtkStack widget.
//
// New values may be added to this enumeration over time.
type StackTransitionType C.GtkStackTransitionType

const (
	// No transition
	GTK_STACK_TRANSITION_TYPE_NONE StackTransitionType = 0

	// A cross-fade
	GTK_STACK_TRANSITION_TYPE_CROSSFADE StackTransitionType = 1

	// Slide from left to right
	GTK_STACK_TRANSITION_TYPE_SLIDE_RIGHT StackTransitionType = 2

	// Slide from right to left
	GTK_STACK_TRANSITION_TYPE_SLIDE_LEFT StackTransitionType = 3

	// Slide from bottom up
	GTK_STACK_TRANSITION_TYPE_SLIDE_UP StackTransitionType = 4

	// Slide from top down
	GTK_STACK_TRANSITION_TYPE_SLIDE_DOWN StackTransitionType = 5

	// Slide from left or right according to the children order
	GTK_STACK_TRANSITION_TYPE_SLIDE_LEFT_RIGHT StackTransitionType = 6

	// Slide from top down or bottom up according to the order
	GTK_STACK_TRANSITION_TYPE_SLIDE_UP_DOWN StackTransitionType = 7

	// Cover the old page by sliding up. Since 3.12
	GTK_STACK_TRANSITION_TYPE_OVER_UP StackTransitionType = 8

	// Cover the old page by sliding down. Since: 3.12
	GTK_STACK_TRANSITION_TYPE_OVER_DOWN StackTransitionType = 9

	// Cover the old page by sliding to the left. Since: 3.12
	GTK_STACK_TRANSITION_TYPE_OVER_LEFT StackTransitionType = 10

	// Cover the old page by sliding to the right. Since: 3.12
	GTK_STACK_TRANSITION_TYPE_OVER_RIGHT StackTransitionType = 11

	// Uncover the new page by sliding up. Since 3.12
	GTK_STACK_TRANSITION_TYPE_UNDER_UP StackTransitionType = 12

	// Uncover the new page by sliding down. Since: 3.12
	GTK_STACK_TRANSITION_TYPE_UNDER_DOWN StackTransitionType = 13

	// Uncover the new page by sliding to the left. Since: 3.12
	GTK_STACK_TRANSITION_TYPE_UNDER_LEFT StackTransitionType = 14

	// Uncover the new page by sliding to the right. Since: 3.12
	GTK_STACK_TRANSITION_TYPE_UNDER_RIGHT StackTransitionType = 15

	// Cover the old page sliding up or uncover the new page sliding down, according to order. Since: 3.12
	GTK_STACK_TRANSITION_TYPE_OVER_UP_DOWN StackTransitionType = 16

	// Cover the old page sliding down or uncover the new page sliding up, according to order. Since: 3.14
	GTK_STACK_TRANSITION_TYPE_OVER_DOWN_UP StackTransitionType = 17

	// Cover the old page sliding left or uncover the new page sliding right, according to order. Since: 3.14
	GTK_STACK_TRANSITION_TYPE_OVER_LEFT_RIGHT StackTransitionType = 18

	// Cover the old page sliding right or uncover the new page sliding left, according to order. Since: 3.14
	GTK_STACK_TRANSITION_TYPE_OVER_RIGHT_LEFT StackTransitionType = 19
)

// This type indicates the current state of a widget; the state determines how
// the widget is drawn. The #GtkStateType enumeration is also used to
// identify different colors in a #GtkStyle for drawing, so states can be
// used for subparts of a widget as well as entire widgets.
type StateType C.GtkStateType

const (
	// State during normal operation.
	GTK_STATE_NORMAL StateType = 0

	// State of a currently active widget, such as a depressed button.
	GTK_STATE_ACTIVE StateType = 1

	// State indicating that the mouse pointer is over
	// the widget and the widget will respond to mouse clicks.
	GTK_STATE_PRELIGHT StateType = 2

	// State of a selected item, such the selected row in a list.
	GTK_STATE_SELECTED StateType = 3

	// State indicating that the widget is
	// unresponsive to user actions.
	GTK_STATE_INSENSITIVE StateType = 4

	// The widget is inconsistent, such as checkbuttons
	// or radiobuttons that aren’t either set to %TRUE nor %FALSE,
	// or buttons requiring the user attention.
	GTK_STATE_INCONSISTENT StateType = 5

	// The widget has the keyboard focus.
	GTK_STATE_FOCUSED StateType = 6
)

// These values are used as “info” for the targets contained in the
// lists returned by gtk_text_buffer_get_copy_target_list() and
// gtk_text_buffer_get_paste_target_list().
//
// The values counts down from `-1` to avoid clashes
// with application added drag destinations which usually start at 0.
type TextBufferTargetInfo C.GtkTextBufferTargetInfo

const (
	// Buffer contents
	GTK_TEXT_BUFFER_TARGET_INFO_BUFFER_CONTENTS TextBufferTargetInfo = -1

	// Rich text
	GTK_TEXT_BUFFER_TARGET_INFO_RICH_TEXT TextBufferTargetInfo = -2

	// Text
	GTK_TEXT_BUFFER_TARGET_INFO_TEXT TextBufferTargetInfo = -3
)

// Reading directions for text.
type TextDirection C.GtkTextDirection

const (
	// No direction.
	GTK_TEXT_DIR_NONE TextDirection = 0

	// Left to right text direction.
	GTK_TEXT_DIR_LTR TextDirection = 1

	// Right to left text direction.
	GTK_TEXT_DIR_RTL TextDirection = 2
)

// Used to reference the layers of #GtkTextView for the purpose of customized
// drawing with the ::draw_layer vfunc.
type TextViewLayer C.GtkTextViewLayer

const (
	// Old deprecated layer, use %GTK_TEXT_VIEW_LAYER_BELOW_TEXT instead
	GTK_TEXT_VIEW_LAYER_BELOW TextViewLayer = 0

	// Old deprecated layer, use %GTK_TEXT_VIEW_LAYER_ABOVE_TEXT instead
	GTK_TEXT_VIEW_LAYER_ABOVE TextViewLayer = 1

	// The layer rendered below the text (but above the background).  Since: 3.20
	GTK_TEXT_VIEW_LAYER_BELOW_TEXT TextViewLayer = 2

	// The layer rendered above the text.  Since: 3.20
	GTK_TEXT_VIEW_LAYER_ABOVE_TEXT TextViewLayer = 3
)

// Used to reference the parts of #GtkTextView.
type TextWindowType C.GtkTextWindowType

const (
	// Invalid value, used as a marker
	GTK_TEXT_WINDOW_PRIVATE TextWindowType = 0

	// Window that floats over scrolling areas.
	GTK_TEXT_WINDOW_WIDGET TextWindowType = 1

	// Scrollable text window.
	GTK_TEXT_WINDOW_TEXT TextWindowType = 2

	// Left side border window.
	GTK_TEXT_WINDOW_LEFT TextWindowType = 3

	// Right side border window.
	GTK_TEXT_WINDOW_RIGHT TextWindowType = 4

	// Top border window.
	GTK_TEXT_WINDOW_TOP TextWindowType = 5

	// Bottom border window.
	GTK_TEXT_WINDOW_BOTTOM TextWindowType = 6
)

// Whether spacers are vertical lines or just blank.
type ToolbarSpaceStyle C.GtkToolbarSpaceStyle

const (
	// Use blank spacers.
	GTK_TOOLBAR_SPACE_EMPTY ToolbarSpaceStyle = 0

	// Use vertical lines for spacers.
	GTK_TOOLBAR_SPACE_LINE ToolbarSpaceStyle = 1
)

// Used to customize the appearance of a #GtkToolbar. Note that
// setting the toolbar style overrides the user’s preferences
// for the default toolbar style.  Note that if the button has only
// a label set and GTK_TOOLBAR_ICONS is used, the label will be
// visible, and vice versa.
type ToolbarStyle C.GtkToolbarStyle

const (
	// Buttons display only icons in the toolbar.
	GTK_TOOLBAR_ICONS ToolbarStyle = 0

	// Buttons display only text labels in the toolbar.
	GTK_TOOLBAR_TEXT ToolbarStyle = 1

	// Buttons display text and icons in the toolbar.
	GTK_TOOLBAR_BOTH ToolbarStyle = 2

	// Buttons display icons and text alongside each
	// other, rather than vertically stacked
	GTK_TOOLBAR_BOTH_HORIZ ToolbarStyle = 3
)

// The sizing method the column uses to determine its width.  Please note
// that @GTK_TREE_VIEW_COLUMN_AUTOSIZE are inefficient for large views, and
// can make columns appear choppy.
type TreeViewColumnSizing C.GtkTreeViewColumnSizing

const (
	// Columns only get bigger in reaction to changes in the model
	GTK_TREE_VIEW_COLUMN_GROW_ONLY TreeViewColumnSizing = 0

	// Columns resize to be the optimal size everytime the model changes.
	GTK_TREE_VIEW_COLUMN_AUTOSIZE TreeViewColumnSizing = 1

	// Columns are a fixed numbers of pixels wide.
	GTK_TREE_VIEW_COLUMN_FIXED TreeViewColumnSizing = 2
)

// An enum for determining where a dropped row goes.
type TreeViewDropPosition C.GtkTreeViewDropPosition

const (
	// dropped row is inserted before
	GTK_TREE_VIEW_DROP_BEFORE TreeViewDropPosition = 0

	// dropped row is inserted after
	GTK_TREE_VIEW_DROP_AFTER TreeViewDropPosition = 1

	// dropped row becomes a child or is inserted before
	GTK_TREE_VIEW_DROP_INTO_OR_BEFORE TreeViewDropPosition = 2

	// dropped row becomes a child or is inserted after
	GTK_TREE_VIEW_DROP_INTO_OR_AFTER TreeViewDropPosition = 3
)

// Used to indicate which grid lines to draw in a tree view.
type TreeViewGridLines C.GtkTreeViewGridLines

const (
	// No grid lines.
	GTK_TREE_VIEW_GRID_LINES_NONE TreeViewGridLines = 0

	// Horizontal grid lines.
	GTK_TREE_VIEW_GRID_LINES_HORIZONTAL TreeViewGridLines = 1

	// Vertical grid lines.
	GTK_TREE_VIEW_GRID_LINES_VERTICAL TreeViewGridLines = 2

	// Horizontal and vertical grid lines.
	GTK_TREE_VIEW_GRID_LINES_BOTH TreeViewGridLines = 3
)

// See also gtk_print_settings_set_paper_width().
type Unit C.GtkUnit

const (
	// No units.
	GTK_UNIT_NONE Unit = 0

	// Dimensions in points.
	GTK_UNIT_POINTS Unit = 1

	// Dimensions in inches.
	GTK_UNIT_INCH Unit = 2

	// Dimensions in millimeters
	GTK_UNIT_MM Unit = 3
)

// Kinds of widget-specific help. Used by the ::show-help signal.
type WidgetHelpType C.GtkWidgetHelpType

const (
	// Tooltip.
	GTK_WIDGET_HELP_TOOLTIP WidgetHelpType = 0

	// What’s this.
	GTK_WIDGET_HELP_WHATS_THIS WidgetHelpType = 1
)

// Window placement can be influenced using this enumeration. Note that
// using #GTK_WIN_POS_CENTER_ALWAYS is almost always a bad idea.
// It won’t necessarily work well with all window managers or on all windowing systems.
type WindowPosition C.GtkWindowPosition

const (
	// No influence is made on placement.
	GTK_WIN_POS_NONE WindowPosition = 0

	// Windows should be placed in the center of the screen.
	GTK_WIN_POS_CENTER WindowPosition = 1

	// Windows should be placed at the current mouse position.
	GTK_WIN_POS_MOUSE WindowPosition = 2

	// Keep window centered as it changes size, etc.
	GTK_WIN_POS_CENTER_ALWAYS WindowPosition = 3

	// Center the window on its transient
	// parent (see gtk_window_set_transient_for()).
	GTK_WIN_POS_CENTER_ON_PARENT WindowPosition = 4
)

// A #GtkWindow can be one of these types. Most things you’d consider a
// “window” should have type #GTK_WINDOW_TOPLEVEL; windows with this type
// are managed by the window manager and have a frame by default (call
// gtk_window_set_decorated() to toggle the frame).  Windows with type
// #GTK_WINDOW_POPUP are ignored by the window manager; window manager
// keybindings won’t work on them, the window manager won’t decorate the
// window with a frame, many GTK+ features that rely on the window
// manager will not work (e.g. resize grips and
// maximization/minimization). #GTK_WINDOW_POPUP is used to implement
// widgets such as #GtkMenu or tooltips that you normally don’t think of
// as windows per se. Nearly all windows should be #GTK_WINDOW_TOPLEVEL.
// In particular, do not use #GTK_WINDOW_POPUP just to turn off
// the window borders; use gtk_window_set_decorated() for that.
type WindowType C.GtkWindowType

const (
	// A regular window, such as a dialog.
	GTK_WINDOW_TOPLEVEL WindowType = 0

	// A special window such as a tooltip.
	GTK_WINDOW_POPUP WindowType = 1
)

// Describes a type of line wrapping.
type WrapMode C.GtkWrapMode

const (
	// do not wrap lines; just make the text area wider
	GTK_WRAP_NONE WrapMode = 0

	// wrap text, breaking lines anywhere the cursor can
	// appear (between characters, usually - if you want to be technical,
	// between graphemes, see pango_get_log_attrs())
	GTK_WRAP_CHAR WrapMode = 1

	// wrap text, breaking lines in between words
	GTK_WRAP_WORD WrapMode = 2

	// wrap text, breaking lines in between words, or if
	// that is not enough, also between graphemes
	GTK_WRAP_WORD_CHAR WrapMode = 3
)
