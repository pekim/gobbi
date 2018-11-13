// This is a generated file - DO NOT EDIT

package atk

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <atk/atk.h>
// #include <stdlib.h>
import "C"

// Specifies how xy coordinates are to be interpreted. Used by functions such
// as atk_component_get_position() and atk_text_get_character_extents()
type CoordType C.AtkCoordType

const (
	// specifies xy coordinates relative to the screen
	ATK_XY_SCREEN CoordType = 0
	// specifies xy coordinates relative to the widget's
	// top-level window
	ATK_XY_WINDOW CoordType = 1
)

// Specifies the type of a keyboard evemt.
type KeyEventType C.AtkKeyEventType

const (
	// specifies a key press event
	ATK_KEY_EVENT_PRESS KeyEventType = 0
	// specifies a key release event
	ATK_KEY_EVENT_RELEASE KeyEventType = 1
	// Not a valid value; specifies end of enumeration
	ATK_KEY_EVENT_LAST_DEFINED KeyEventType = 2
)

// Describes the layer of a component
//
// These enumerated "layer values" are used when determining which UI
// rendering layer a component is drawn into, which can help in making
// determinations of when components occlude one another.
type Layer C.AtkLayer

const (
	// The object does not have a layer
	ATK_LAYER_INVALID Layer = 0
	// This layer is reserved for the desktop background
	ATK_LAYER_BACKGROUND Layer = 1
	// This layer is used for Canvas components
	ATK_LAYER_CANVAS Layer = 2
	// This layer is normally used for components
	ATK_LAYER_WIDGET Layer = 3
	// This layer is used for layered components
	ATK_LAYER_MDI Layer = 4
	// This layer is used for popup components, such as menus
	ATK_LAYER_POPUP Layer = 5
	// This layer is reserved for future use.
	ATK_LAYER_OVERLAY Layer = 6
	// This layer is used for toplevel windows.
	ATK_LAYER_WINDOW Layer = 7
)

// Describes the type of the relation
type RelationType C.AtkRelationType

const (
	// Not used, represens "no relationship" or an error condition.
	ATK_RELATION_NULL RelationType = 0
	// Indicates an object controlled by one or more target objects.
	ATK_RELATION_CONTROLLED_BY RelationType = 1
	// Indicates an object is an controller for one or more target objects.
	ATK_RELATION_CONTROLLER_FOR RelationType = 2
	// Indicates an object is a label for one or more target objects.
	ATK_RELATION_LABEL_FOR RelationType = 3
	// Indicates an object is labelled by one or more target objects.
	ATK_RELATION_LABELLED_BY RelationType = 4
	// Indicates an object is a member of a group of one or more target objects.
	ATK_RELATION_MEMBER_OF RelationType = 5
	// Indicates an object is a cell in a treetable which is displayed because a cell in the same column is expanded and identifies that cell.
	ATK_RELATION_NODE_CHILD_OF RelationType = 6
	// Indicates that the object has content that flows logically to another
	// AtkObject in a sequential way, (for instance text-flow).
	ATK_RELATION_FLOWS_TO RelationType = 7
	// Indicates that the object has content that flows logically from
	// another AtkObject in a sequential way, (for instance text-flow).
	ATK_RELATION_FLOWS_FROM RelationType = 8
	// Indicates a subwindow attached to a component but otherwise has no connection in  the UI heirarchy to that component.
	ATK_RELATION_SUBWINDOW_OF RelationType = 9
	// Indicates that the object visually embeds
	// another object's content, i.e. this object's content flows around
	// another's content.
	ATK_RELATION_EMBEDS RelationType = 10
	// Reciprocal of %ATK_RELATION_EMBEDS, indicates that
	// this object's content is visualy embedded in another object.
	ATK_RELATION_EMBEDDED_BY RelationType = 11
	// Indicates that an object is a popup for another object.
	ATK_RELATION_POPUP_FOR RelationType = 12
	// Indicates that an object is a parent window of another object.
	ATK_RELATION_PARENT_WINDOW_OF RelationType = 13
	// Reciprocal of %ATK_RELATION_DESCRIPTION_FOR. Indicates that one
	// or more target objects provide descriptive information about this object. This relation
	// type is most appropriate for information that is not essential as its presentation may
	// be user-configurable and/or limited to an on-demand mechanism such as an assistive
	// technology command. For brief, essential information such as can be found in a widget's
	// on-screen label, use %ATK_RELATION_LABELLED_BY. For an on-screen error message, use
	// %ATK_RELATION_ERROR_MESSAGE. For lengthy extended descriptive information contained in
	// an on-screen object, consider using %ATK_RELATION_DETAILS as assistive technologies may
	// provide a means for the user to navigate to objects containing detailed descriptions so
	// that their content can be more closely reviewed.
	ATK_RELATION_DESCRIBED_BY RelationType = 14
	// Reciprocal of %ATK_RELATION_DESCRIBED_BY. Indicates that this
	// object provides descriptive information about the target object(s). See also
	// %ATK_RELATION_DETAILS_FOR and %ATK_RELATION_ERROR_FOR.
	ATK_RELATION_DESCRIPTION_FOR RelationType = 15
	// Indicates an object is a cell in a treetable and is expanded to display other cells in the same column.
	ATK_RELATION_NODE_PARENT_OF RelationType = 16
	// Reciprocal of %ATK_RELATION_DETAILS_FOR. Indicates that this object
	// has a detailed or extended description, the contents of which can be found in the target
	// object(s). This relation type is most appropriate for information that is sufficiently
	// lengthy as to make navigation to the container of that information desirable. For less
	// verbose information suitable for announcement only, see %ATK_RELATION_DESCRIBED_BY. If
	// the detailed information describes an error condition, %ATK_RELATION_ERROR_FOR should be
	// used instead. @Since: ATK-2.26.
	ATK_RELATION_DETAILS RelationType = 17
	// Reciprocal of %ATK_RELATION_DETAILS. Indicates that this object
	// provides a detailed or extended description about the target object(s). See also
	// %ATK_RELATION_DESCRIPTION_FOR and %ATK_RELATION_ERROR_FOR. @Since: ATK-2.26.
	ATK_RELATION_DETAILS_FOR RelationType = 18
	// Reciprocal of %ATK_RELATION_ERROR_FOR. Indicates that this object
	// has one or more errors, the nature of which is described in the contents of the target
	// object(s). Objects that have this relation type should also contain %ATK_STATE_INVALID_ENTRY
	// in their #AtkStateSet. @Since: ATK-2.26.
	ATK_RELATION_ERROR_MESSAGE RelationType = 19
	// Reciprocal of %ATK_RELATION_ERROR_MESSAGE. Indicates that this object
	// contains an error message describing an invalid condition in the target object(s). @Since:
	// ATK_2.26.
	ATK_RELATION_ERROR_FOR RelationType = 20
	// Not used, this value indicates the end of the enumeration.
	ATK_RELATION_LAST_DEFINED RelationType = 21
)

// Describes the role of an object
//
// These are the built-in enumerated roles that UI components can have in
// ATK.  Other roles may be added at runtime, so an AtkRole >=
// ATK_ROLE_LAST_DEFINED is not necessarily an error.
type Role C.AtkRole

const (
	// Invalid role
	ATK_ROLE_INVALID Role = 0
	// A label which represents an accelerator
	ATK_ROLE_ACCEL_LABEL Role = 1
	// An object which is an alert to the user. Assistive Technologies typically respond to ATK_ROLE_ALERT by reading the entire onscreen contents of containers advertising this role.  Should be used for warning dialogs, etc.
	ATK_ROLE_ALERT Role = 2
	// An object which is an animated image
	ATK_ROLE_ANIMATION Role = 3
	// An arrow in one of the four cardinal directions
	ATK_ROLE_ARROW Role = 4
	// An object that displays a calendar and allows the user to select a date
	ATK_ROLE_CALENDAR Role = 5
	// An object that can be drawn into and is used to trap events
	ATK_ROLE_CANVAS Role = 6
	// A choice that can be checked or unchecked and provides a separate indicator for the current state
	ATK_ROLE_CHECK_BOX Role = 7
	// A menu item with a check box
	ATK_ROLE_CHECK_MENU_ITEM Role = 8
	// A specialized dialog that lets the user choose a color
	ATK_ROLE_COLOR_CHOOSER Role = 9
	// The header for a column of data
	ATK_ROLE_COLUMN_HEADER Role = 10
	// A collapsible list of choices the user can select from
	ATK_ROLE_COMBO_BOX Role = 11
	// An object whose purpose is to allow a user to edit a date
	ATK_ROLE_DATE_EDITOR Role = 12
	// An inconifed internal frame within a DESKTOP_PANE
	ATK_ROLE_DESKTOP_ICON Role = 13
	// A pane that supports internal frames and iconified versions of those internal frames
	ATK_ROLE_DESKTOP_FRAME Role = 14
	// An object whose purpose is to allow a user to set a value
	ATK_ROLE_DIAL Role = 15
	// A top level window with title bar and a border
	ATK_ROLE_DIALOG Role = 16
	// A pane that allows the user to navigate through and select the contents of a directory
	ATK_ROLE_DIRECTORY_PANE Role = 17
	// An object used for drawing custom user interface elements
	ATK_ROLE_DRAWING_AREA Role = 18
	// A specialized dialog that lets the user choose a file
	ATK_ROLE_FILE_CHOOSER Role = 19
	// A object that fills up space in a user interface
	ATK_ROLE_FILLER Role = 20
	// A specialized dialog that lets the user choose a font
	ATK_ROLE_FONT_CHOOSER Role = 21
	// A top level window with a title bar, border, menubar, etc.
	ATK_ROLE_FRAME Role = 22
	// A pane that is guaranteed to be painted on top of all panes beneath it
	ATK_ROLE_GLASS_PANE Role = 23
	// A document container for HTML, whose children represent the document content
	ATK_ROLE_HTML_CONTAINER Role = 24
	// A small fixed size picture, typically used to decorate components
	ATK_ROLE_ICON Role = 25
	// An object whose primary purpose is to display an image
	ATK_ROLE_IMAGE Role = 26
	// A frame-like object that is clipped by a desktop pane
	ATK_ROLE_INTERNAL_FRAME Role = 27
	// An object used to present an icon or short string in an interface
	ATK_ROLE_LABEL Role = 28
	// A specialized pane that allows its children to be drawn in layers, providing a form of stacking order
	ATK_ROLE_LAYERED_PANE Role = 29
	// An object that presents a list of objects to the user and allows the user to select one or more of them
	ATK_ROLE_LIST Role = 30
	// An object that represents an element of a list
	ATK_ROLE_LIST_ITEM Role = 31
	// An object usually found inside a menu bar that contains a list of actions the user can choose from
	ATK_ROLE_MENU Role = 32
	// An object usually drawn at the top of the primary dialog box of an application that contains a list of menus the user can choose from
	ATK_ROLE_MENU_BAR Role = 33
	// An object usually contained in a menu that presents an action the user can choose
	ATK_ROLE_MENU_ITEM Role = 34
	// A specialized pane whose primary use is inside a DIALOG
	ATK_ROLE_OPTION_PANE Role = 35
	// An object that is a child of a page tab list
	ATK_ROLE_PAGE_TAB Role = 36
	// An object that presents a series of panels (or page tabs), one at a time, through some mechanism provided by the object
	ATK_ROLE_PAGE_TAB_LIST Role = 37
	// A generic container that is often used to group objects
	ATK_ROLE_PANEL Role = 38
	// A text object uses for passwords, or other places where the text content is not shown visibly to the user
	ATK_ROLE_PASSWORD_TEXT Role = 39
	// A temporary window that is usually used to offer the user a list of choices, and then hides when the user selects one of those choices
	ATK_ROLE_POPUP_MENU Role = 40
	// An object used to indicate how much of a task has been completed
	ATK_ROLE_PROGRESS_BAR Role = 41
	// An object the user can manipulate to tell the application to do something
	ATK_ROLE_PUSH_BUTTON Role = 42
	// A specialized check box that will cause other radio buttons in the same group to become unchecked when this one is checked
	ATK_ROLE_RADIO_BUTTON Role = 43
	// A check menu item which belongs to a group. At each instant exactly one of the radio menu items from a group is selected
	ATK_ROLE_RADIO_MENU_ITEM Role = 44
	// A specialized pane that has a glass pane and a layered pane as its children
	ATK_ROLE_ROOT_PANE Role = 45
	// The header for a row of data
	ATK_ROLE_ROW_HEADER Role = 46
	// An object usually used to allow a user to incrementally view a large amount of data.
	ATK_ROLE_SCROLL_BAR Role = 47
	// An object that allows a user to incrementally view a large amount of information
	ATK_ROLE_SCROLL_PANE Role = 48
	// An object usually contained in a menu to provide a visible and logical separation of the contents in a menu
	ATK_ROLE_SEPARATOR Role = 49
	// An object that allows the user to select from a bounded range
	ATK_ROLE_SLIDER Role = 50
	// A specialized panel that presents two other panels at the same time
	ATK_ROLE_SPLIT_PANE Role = 51
	// An object used to get an integer or floating point number from the user
	ATK_ROLE_SPIN_BUTTON Role = 52
	// An object which reports messages of minor importance to the user
	ATK_ROLE_STATUSBAR Role = 53
	// An object used to represent information in terms of rows and columns
	ATK_ROLE_TABLE Role = 54
	// A cell in a table
	ATK_ROLE_TABLE_CELL Role = 55
	// The header for a column of a table
	ATK_ROLE_TABLE_COLUMN_HEADER Role = 56
	// The header for a row of a table
	ATK_ROLE_TABLE_ROW_HEADER Role = 57
	// A menu item used to tear off and reattach its menu
	ATK_ROLE_TEAR_OFF_MENU_ITEM Role = 58
	// An object that represents an accessible terminal.  @Since: ATK-0.6
	ATK_ROLE_TERMINAL Role = 59
	// An interactive widget that supports multiple lines of text and
	// optionally accepts user input, but whose purpose is not to solicit user input.
	// Thus ATK_ROLE_TEXT is appropriate for the text view in a plain text editor
	// but inappropriate for an input field in a dialog box or web form. For widgets
	// whose purpose is to solicit input from the user, see ATK_ROLE_ENTRY and
	// ATK_ROLE_PASSWORD_TEXT. For generic objects which display a brief amount of
	// textual information, see ATK_ROLE_STATIC.
	ATK_ROLE_TEXT Role = 60
	// A specialized push button that can be checked or unchecked, but does not provide a separate indicator for the current state
	ATK_ROLE_TOGGLE_BUTTON Role = 61
	// A bar or palette usually composed of push buttons or toggle buttons
	ATK_ROLE_TOOL_BAR Role = 62
	// An object that provides information about another object
	ATK_ROLE_TOOL_TIP Role = 63
	// An object used to represent hierarchical information to the user
	ATK_ROLE_TREE Role = 64
	// An object capable of expanding and collapsing rows as well as showing multiple columns of data.   @Since: ATK-0.7
	ATK_ROLE_TREE_TABLE Role = 65
	// The object contains some Accessible information, but its role is not known
	ATK_ROLE_UNKNOWN Role = 66
	// An object usually used in a scroll pane
	ATK_ROLE_VIEWPORT Role = 67
	// A top level window with no title or border.
	ATK_ROLE_WINDOW Role = 68
	// An object that serves as a document header. @Since: ATK-1.1.1
	ATK_ROLE_HEADER Role = 69
	// An object that serves as a document footer.  @Since: ATK-1.1.1
	ATK_ROLE_FOOTER Role = 70
	// An object which is contains a paragraph of text content.   @Since: ATK-1.1.1
	ATK_ROLE_PARAGRAPH Role = 71
	// An object which describes margins and tab stops, etc. for text objects which it controls (should have CONTROLLER_FOR relation to such).   @Since: ATK-1.1.1
	ATK_ROLE_RULER Role = 72
	// The object is an application object, which may contain @ATK_ROLE_FRAME objects or other types of accessibles.  The root accessible of any application's ATK hierarchy should have ATK_ROLE_APPLICATION.   @Since: ATK-1.1.4
	ATK_ROLE_APPLICATION Role = 73
	// The object is a dialog or list containing items for insertion into an entry widget, for instance a list of words for completion of a text entry.   @Since: ATK-1.3
	ATK_ROLE_AUTOCOMPLETE Role = 74
	// The object is an editable text object in a toolbar.  @Since: ATK-1.5
	ATK_ROLE_EDITBAR Role = 75
	// The object is an embedded container within a document or panel.  This role is a grouping "hint" indicating that the contained objects share a context.  @Since: ATK-1.7.2
	ATK_ROLE_EMBEDDED Role = 76
	// The object is a component whose textual content may be entered or modified by the user, provided @ATK_STATE_EDITABLE is present.   @Since: ATK-1.11
	ATK_ROLE_ENTRY Role = 77
	// The object is a graphical depiction of quantitative data. It may contain multiple subelements whose attributes and/or description may be queried to obtain both the quantitative data and information about how the data is being presented. The LABELLED_BY relation is particularly important in interpreting objects of this type, as is the accessible-description property.  @Since: ATK-1.11
	ATK_ROLE_CHART Role = 78
	// The object contains descriptive information, usually textual, about another user interface element such as a table, chart, or image.  @Since: ATK-1.11
	ATK_ROLE_CAPTION Role = 79
	// The object is a visual frame or container which contains a view of document content. Document frames may occur within another Document instance, in which case the second document may be said to be embedded in the containing instance. HTML frames are often ROLE_DOCUMENT_FRAME. Either this object, or a singleton descendant, should implement the Document interface.  @Since: ATK-1.11
	ATK_ROLE_DOCUMENT_FRAME Role = 80
	// The object serves as a heading for content which follows it in a document. The 'heading level' of the heading, if availabe, may be obtained by querying the object's attributes.
	ATK_ROLE_HEADING Role = 81
	// The object is a containing instance which encapsulates a page of information. @ATK_ROLE_PAGE is used in documents and content which support a paginated navigation model.  @Since: ATK-1.11
	ATK_ROLE_PAGE Role = 82
	// The object is a containing instance of document content which constitutes a particular 'logical' section of the document. The type of content within a section, and the nature of the section division itself, may be obtained by querying the object's attributes. Sections may be nested. @Since: ATK-1.11
	ATK_ROLE_SECTION Role = 83
	// The object is redundant with another object in the hierarchy, and is exposed for purely technical reasons.  Objects of this role should normally be ignored by clients. @Since: ATK-1.11
	ATK_ROLE_REDUNDANT_OBJECT Role = 84
	// The object is a container for form controls, for instance as part of a
	// web form or user-input form within a document.  This role is primarily a tag/convenience for
	// clients when navigating complex documents, it is not expected that ordinary GUI containers will
	// always have ATK_ROLE_FORM. @Since: ATK-1.12.0
	ATK_ROLE_FORM Role = 85
	// The object is a hypertext anchor, i.e. a "link" in a
	// hypertext document.  Such objects are distinct from 'inline'
	// content which may also use the Hypertext/Hyperlink interfaces
	// to indicate the range/location within a text object where
	// an inline or embedded object lies.  @Since: ATK-1.12.1
	ATK_ROLE_LINK Role = 86
	// The object is a window or similar viewport
	// which is used to allow composition or input of a 'complex character',
	// in other words it is an "input method window." @Since: ATK-1.12.1
	ATK_ROLE_INPUT_METHOD_WINDOW Role = 87
	// A row in a table.  @Since: ATK-2.1.0
	ATK_ROLE_TABLE_ROW Role = 88
	// An object that represents an element of a tree.  @Since: ATK-2.1.0
	ATK_ROLE_TREE_ITEM Role = 89
	// A document frame which contains a spreadsheet.  @Since: ATK-2.1.0
	ATK_ROLE_DOCUMENT_SPREADSHEET Role = 90
	// A document frame which contains a presentation or slide content.  @Since: ATK-2.1.0
	ATK_ROLE_DOCUMENT_PRESENTATION Role = 91
	// A document frame which contains textual content, such as found in a word processing application.  @Since: ATK-2.1.0
	ATK_ROLE_DOCUMENT_TEXT Role = 92
	// A document frame which contains HTML or other markup suitable for display in a web browser.  @Since: ATK-2.1.0
	ATK_ROLE_DOCUMENT_WEB Role = 93
	// A document frame which contains email content to be displayed or composed either in plain text or HTML.  @Since: ATK-2.1.0
	ATK_ROLE_DOCUMENT_EMAIL Role = 94
	// An object found within a document and designed to present a comment, note, or other annotation. In some cases, this object might not be visible until activated.  @Since: ATK-2.1.0
	ATK_ROLE_COMMENT Role = 95
	// A non-collapsible list of choices the user can select from. @Since: ATK-2.1.0
	ATK_ROLE_LIST_BOX Role = 96
	// A group of related widgets. This group typically has a label. @Since: ATK-2.1.0
	ATK_ROLE_GROUPING Role = 97
	// An image map object. Usually a graphic with multiple hotspots, where each hotspot can be activated resulting in the loading of another document or section of a document. @Since: ATK-2.1.0
	ATK_ROLE_IMAGE_MAP Role = 98
	// A transitory object designed to present a message to the user, typically at the desktop level rather than inside a particular application.  @Since: ATK-2.1.0
	ATK_ROLE_NOTIFICATION Role = 99
	// An object designed to present a message to the user within an existing window. @Since: ATK-2.1.0
	ATK_ROLE_INFO_BAR Role = 100
	// A bar that serves as a level indicator to, for instance, show the strength of a password or the state of a battery.  @Since: ATK-2.7.3
	ATK_ROLE_LEVEL_BAR Role = 101
	// A bar that serves as the title of a window or a
	// dialog. @Since: ATK-2.12
	ATK_ROLE_TITLE_BAR Role = 102
	// An object which contains a text section
	// that is quoted from another source. @Since: ATK-2.12
	ATK_ROLE_BLOCK_QUOTE Role = 103
	// An object which represents an audio element. @Since: ATK-2.12
	ATK_ROLE_AUDIO Role = 104
	// An object which represents a video element. @Since: ATK-2.12
	ATK_ROLE_VIDEO Role = 105
	// A definition of a term or concept. @Since: ATK-2.12
	ATK_ROLE_DEFINITION Role = 106
	// A section of a page that consists of a
	// composition that forms an independent part of a document, page, or
	// site. Examples: A blog entry, a news story, a forum post. @Since:
	// ATK-2.12
	ATK_ROLE_ARTICLE Role = 107
	// A region of a web page intended as a
	// navigational landmark. This is designed to allow Assistive
	// Technologies to provide quick navigation among key regions within a
	// document. @Since: ATK-2.12
	ATK_ROLE_LANDMARK Role = 108
	// A text widget or container holding log content, such
	// as chat history and error logs. In this role there is a
	// relationship between the arrival of new items in the log and the
	// reading order. The log contains a meaningful sequence and new
	// information is added only to the end of the log, not at arbitrary
	// points. @Since: ATK-2.12
	ATK_ROLE_LOG Role = 109
	// A container where non-essential information
	// changes frequently. Common usages of marquee include stock tickers
	// and ad banners. The primary difference between a marquee and a log
	// is that logs usually have a meaningful order or sequence of
	// important content changes. @Since: ATK-2.12
	ATK_ROLE_MARQUEE Role = 110
	// A text widget or container that holds a mathematical
	// expression. @Since: ATK-2.12
	ATK_ROLE_MATH Role = 111
	// A widget whose purpose is to display a rating,
	// such as the number of stars associated with a song in a media
	// player. Objects of this role should also implement
	// AtkValue. @Since: ATK-2.12
	ATK_ROLE_RATING Role = 112
	// An object containing a numerical counter which
	// indicates an amount of elapsed time from a start point, or the time
	// remaining until an end point. @Since: ATK-2.12
	ATK_ROLE_TIMER Role = 113
	// An object that represents a list of
	// term-value groups. A term-value group represents a individual
	// description and consist of one or more names
	// (ATK_ROLE_DESCRIPTION_TERM) followed by one or more values
	// (ATK_ROLE_DESCRIPTION_VALUE). For each list, there should not be
	// more than one group with the same term name. @Since: ATK-2.12
	ATK_ROLE_DESCRIPTION_LIST Role = 114
	// An object that represents a term or phrase
	// with a corresponding definition. @Since: ATK-2.12
	ATK_ROLE_DESCRIPTION_TERM Role = 115
	// An object that represents the
	// description, definition or value of a term. @Since: ATK-2.12
	ATK_ROLE_DESCRIPTION_VALUE Role = 116
	// A generic non-container object whose purpose is to display a
	// brief amount of information to the user and whose role is known by the
	// implementor but lacks semantic value for the user. Examples in which
	// ATK_ROLE_STATIC is appropriate include the message displayed in a message box
	// and an image used as an alternative means to display text. ATK_ROLE_STATIC
	// should not be applied to widgets which are traditionally interactive, objects
	// which display a significant amount of content, or any object which has an
	// accessible relation pointing to another object. Implementors should expose the
	// displayed information through the accessible name of the object. If doing so seems
	// inappropriate, it may indicate that a different role should be used. For
	// labels which describe another widget, see ATK_ROLE_LABEL. For text views, see
	// ATK_ROLE_TEXT. For generic containers, see ATK_ROLE_PANEL. For objects whose
	// role is not known by the implementor, see ATK_ROLE_UNKNOWN. @Since: ATK-2.16.
	ATK_ROLE_STATIC Role = 117
	// An object that represents a mathematical fraction.
	ATK_ROLE_MATH_FRACTION Role = 118
	// An object that represents a mathematical expression
	// displayed with a radical. @Since: ATK-2.16.
	ATK_ROLE_MATH_ROOT Role = 119
	// An object that contains text that is displayed as a
	// subscript. @Since: ATK-2.16.
	ATK_ROLE_SUBSCRIPT Role = 120
	// An object that contains text that is displayed as a
	// superscript. @Since: ATK-2.16.
	ATK_ROLE_SUPERSCRIPT Role = 121
	// An object that contains the text of a footnote. @Since: ATK-2.26.
	ATK_ROLE_FOOTNOTE Role = 122
	// not a valid role, used for finding end of the enumeration
	ATK_ROLE_LAST_DEFINED Role = 123
)

// The possible types of states of an object
type StateType C.AtkStateType

const (
	// Indicates an invalid state - probably an error condition.
	ATK_STATE_INVALID StateType = 0
	// Indicates a window is currently the active window, or an object is the active subelement within a container or table. ATK_STATE_ACTIVE should not be used for objects which have ATK_STATE_FOCUSABLE or ATK_STATE_SELECTABLE: Those objects should use ATK_STATE_FOCUSED and ATK_STATE_SELECTED respectively. ATK_STATE_ACTIVE is a means to indicate that an object which is not focusable and not selectable is the currently-active item within its parent container.
	ATK_STATE_ACTIVE StateType = 1
	// Indicates that the object is 'armed', i.e. will be activated by if a pointer button-release event occurs within its bounds.  Buttons often enter this state when a pointer click occurs within their bounds, as a precursor to activation. ATK_STATE_ARMED has been deprecated since ATK-2.16 and should not be used in newly-written code.
	ATK_STATE_ARMED StateType = 2
	// Indicates the current object is busy, i.e. onscreen representation is in the process of changing, or the object is temporarily unavailable for interaction due to activity already in progress.  This state may be used by implementors of Document to indicate that content loading is underway.  It also may indicate other 'pending' conditions; clients may wish to interrogate this object when the ATK_STATE_BUSY flag is removed.
	ATK_STATE_BUSY StateType = 3
	// Indicates this object is currently checked, for instance a checkbox is 'non-empty'.
	ATK_STATE_CHECKED StateType = 4
	// Indicates that this object no longer has a valid backing widget (for instance, if its peer object has been destroyed)
	ATK_STATE_DEFUNCT StateType = 5
	// Indicates that this object can contain text, and that the
	// user can change the textual contents of this object by editing those contents
	// directly. For an object which is expected to be editable due to its type, but
	// which cannot be edited due to the application or platform preventing the user
	// from doing so, that object's #AtkStateSet should lack ATK_STATE_EDITABLE and
	// should contain ATK_STATE_READ_ONLY.
	ATK_STATE_EDITABLE StateType = 6
	// Indicates that this object is enabled, i.e. that it currently reflects some application state. Objects that are "greyed out" may lack this state, and may lack the STATE_SENSITIVE if direct user interaction cannot cause them to acquire STATE_ENABLED. See also: ATK_STATE_SENSITIVE
	ATK_STATE_ENABLED StateType = 7
	// Indicates this object allows progressive disclosure of its children
	ATK_STATE_EXPANDABLE StateType = 8
	// Indicates this object its expanded - see ATK_STATE_EXPANDABLE above
	ATK_STATE_EXPANDED StateType = 9
	// Indicates this object can accept keyboard focus, which means all events resulting from typing on the keyboard will normally be passed to it when it has focus
	ATK_STATE_FOCUSABLE StateType = 10
	// Indicates this object currently has the keyboard focus
	ATK_STATE_FOCUSED StateType = 11
	// Indicates the orientation of this object is horizontal; used, for instance, by objects of ATK_ROLE_SCROLL_BAR.  For objects where vertical/horizontal orientation is especially meaningful.
	ATK_STATE_HORIZONTAL StateType = 12
	// Indicates this object is minimized and is represented only by an icon
	ATK_STATE_ICONIFIED StateType = 13
	// Indicates something must be done with this object before the user can interact with an object in a different window
	ATK_STATE_MODAL StateType = 14
	// Indicates this (text) object can contain multiple lines of text
	ATK_STATE_MULTI_LINE StateType = 15
	// Indicates this object allows more than one of its children to be selected at the same time, or in the case of text objects, that the object supports non-contiguous text selections.
	ATK_STATE_MULTISELECTABLE StateType = 16
	// Indicates this object paints every pixel within its rectangular region.
	ATK_STATE_OPAQUE StateType = 17
	// Indicates this object is currently pressed.
	ATK_STATE_PRESSED StateType = 18
	// Indicates the size of this object is not fixed
	ATK_STATE_RESIZABLE StateType = 19
	// Indicates this object is the child of an object that allows its children to be selected and that this child is one of those children that can be selected
	ATK_STATE_SELECTABLE StateType = 20
	// Indicates this object is the child of an object that allows its children to be selected and that this child is one of those children that has been selected
	ATK_STATE_SELECTED StateType = 21
	// Indicates this object is sensitive, e.g. to user interaction.
	// STATE_SENSITIVE usually accompanies STATE_ENABLED for user-actionable controls,
	// but may be found in the absence of STATE_ENABLED if the current visible state of the
	// control is "disconnected" from the application state.  In such cases, direct user interaction
	// can often result in the object gaining STATE_SENSITIVE, for instance if a user makes
	// an explicit selection using an object whose current state is ambiguous or undefined.
	// @see STATE_ENABLED, STATE_INDETERMINATE.
	ATK_STATE_SENSITIVE StateType = 22
	// Indicates this object, the object's parent, the object's parent's parent, and so on,
	// are all 'shown' to the end-user, i.e. subject to "exposure" if blocking or obscuring objects do not interpose
	// between this object and the top of the window stack.
	ATK_STATE_SHOWING StateType = 23
	// Indicates this (text) object can contain only a single line of text
	ATK_STATE_SINGLE_LINE StateType = 24
	// Indicates that the information returned for this object may no longer be
	// synchronized with the application state.  This is implied if the object has STATE_TRANSIENT,
	// and can also occur towards the end of the object peer's lifecycle. It can also be used to indicate that
	// the index associated with this object has changed since the user accessed the object (in lieu of
	// "index-in-parent-changed" events).
	ATK_STATE_STALE StateType = 25
	// Indicates this object is transient, i.e. a snapshot which may not emit events when its
	// state changes.  Data from objects with ATK_STATE_TRANSIENT should not be cached, since there may be no
	// notification given when the cached data becomes obsolete.
	ATK_STATE_TRANSIENT StateType = 26
	// Indicates the orientation of this object is vertical
	ATK_STATE_VERTICAL StateType = 27
	// Indicates this object is visible, e.g. has been explicitly marked for exposure to the user.
	ATK_STATE_VISIBLE StateType = 28
	// Indicates that "active-descendant-changed" event
	// is sent when children become 'active' (i.e. are selected or navigated to onscreen).
	// Used to prevent need to enumerate all children in very large containers, like tables.
	// The presence of STATE_MANAGES_DESCENDANTS is an indication to the client.
	// that the children should not, and need not, be enumerated by the client.
	// Objects implementing this state are expected to provide relevant state
	// notifications to listening clients, for instance notifications of visibility
	// changes and activation of their contained child objects, without the client
	// having previously requested references to those children.
	ATK_STATE_MANAGES_DESCENDANTS StateType = 29
	// Indicates that the value, or some other quantifiable
	// property, of this AtkObject cannot be fully determined. In the case of a large
	// data set in which the total number of items in that set is unknown (e.g. 1 of
	// 999+), implementors should expose the currently-known set size (999) along
	// with this state. In the case of a check box, this state should be used to
	// indicate that the check box is a tri-state check box which is currently
	// neither checked nor unchecked.
	ATK_STATE_INDETERMINATE StateType = 30
	// Indicates that an object is truncated, e.g. a text value in a speradsheet cell.
	ATK_STATE_TRUNCATED StateType = 31
	// Indicates that explicit user interaction with an object is required by the user interface, e.g. a required field in a "web-form" interface.
	ATK_STATE_REQUIRED StateType = 32
	// Indicates that the object has encountered an error condition due to failure of input validation. For instance, a form control may acquire this state in response to invalid or malformed user input.
	ATK_STATE_INVALID_ENTRY StateType = 33
	// Indicates that the object in question implements some form of ¨typeahead¨ or
	// pre-selection behavior whereby entering the first character of one or more sub-elements
	// causes those elements to scroll into view or become selected.  Subsequent character input
	// may narrow the selection further as long as one or more sub-elements match the string.
	// This state is normally only useful and encountered on objects that implement Selection.
	// In some cases the typeahead behavior may result in full or partial ¨completion¨ of
	// the data in the input field, in which case these input events may trigger text-changed
	// events from the AtkText interface.  This state supplants @ATK_ROLE_AUTOCOMPLETE.
	ATK_STATE_SUPPORTS_AUTOCOMPLETION StateType = 34
	// Indicates that the object in question supports text selection. It should only be exposed on objects which implement the Text interface, in order to distinguish this state from @ATK_STATE_SELECTABLE, which infers that the object in question is a selectable child of an object which implements Selection. While similar, text selection and subelement selection are distinct operations.
	ATK_STATE_SELECTABLE_TEXT StateType = 35
	// Indicates that the object is the "default" active component, i.e. the object which is activated by an end-user press of the "Enter" or "Return" key.  Typically a "close" or "submit" button.
	ATK_STATE_DEFAULT StateType = 36
	// Indicates that the object changes its appearance dynamically as an inherent part of its presentation.  This state may come and go if an object is only temporarily animated on the way to a 'final' onscreen presentation.
	// @note some applications, notably content viewers, may not be able to detect
	// all kinds of animated content.  Therefore the absence of this state should not
	// be taken as definitive evidence that the object's visual representation is
	// static; this state is advisory.
	ATK_STATE_ANIMATED StateType = 37
	// Indicates that the object (typically a hyperlink) has already been 'activated', and/or its backing data has already been downloaded, rendered, or otherwise "visited".
	ATK_STATE_VISITED StateType = 38
	// Indicates this object has the potential to be
	// checked, such as a checkbox or toggle-able table cell. @Since:
	// ATK-2.12
	ATK_STATE_CHECKABLE StateType = 39
	// Indicates that the object has a popup context
	// menu or sub-level menu which may or may not be showing. This means
	// that activation renders conditional content.  Note that ordinary
	// tooltips are not considered popups in this context. @Since: ATK-2.12
	ATK_STATE_HAS_POPUP StateType = 40
	// Indicates this object has a tooltip. @Since: ATK-2.16
	ATK_STATE_HAS_TOOLTIP StateType = 41
	// Indicates that a widget which is ENABLED and SENSITIVE
	// has a value which can be read, but not modified, by the user. Note that this
	// state should only be applied to widget types whose value is normally directly
	// user modifiable, such as check boxes, radio buttons, spin buttons, text input
	// fields, and combo boxes, as a means to convey that the expected interaction
	// with that widget is not possible. When the expected interaction with a
	// widget does not include modification by the user, as is the case with
	// labels and containers, ATK_STATE_READ_ONLY should not be applied. See also
	// ATK_STATE_EDITABLE. @Since: ATK-2-16
	ATK_STATE_READ_ONLY StateType = 42
	// Not a valid state, used for finding end of enumeration
	ATK_STATE_LAST_DEFINED StateType = 43
)

// Describes the text attributes supported
type TextAttribute C.AtkTextAttribute

const (
	// Invalid attribute, like bad spelling or grammar.
	ATK_TEXT_ATTR_INVALID TextAttribute = 0
	// The pixel width of the left margin
	ATK_TEXT_ATTR_LEFT_MARGIN TextAttribute = 1
	// The pixel width of the right margin
	ATK_TEXT_ATTR_RIGHT_MARGIN TextAttribute = 2
	// The number of pixels that the text is indented
	ATK_TEXT_ATTR_INDENT TextAttribute = 3
	// Either "true" or "false" indicating whether text is visible or not
	ATK_TEXT_ATTR_INVISIBLE TextAttribute = 4
	// Either "true" or "false" indicating whether text is editable or not
	ATK_TEXT_ATTR_EDITABLE TextAttribute = 5
	// Pixels of blank space to leave above each newline-terminated line.
	ATK_TEXT_ATTR_PIXELS_ABOVE_LINES TextAttribute = 6
	// Pixels of blank space to leave below each newline-terminated line.
	ATK_TEXT_ATTR_PIXELS_BELOW_LINES TextAttribute = 7
	// Pixels of blank space to leave between wrapped lines inside the same newline-terminated line (paragraph).
	ATK_TEXT_ATTR_PIXELS_INSIDE_WRAP TextAttribute = 8
	// "true" or "false" whether to make the background color for each character the height of the highest font used on the current line, or the height of the font used for the current character.
	ATK_TEXT_ATTR_BG_FULL_HEIGHT TextAttribute = 9
	// Number of pixels that the characters are risen above the baseline
	ATK_TEXT_ATTR_RISE TextAttribute = 10
	// "none", "single", "double", "low", or "error"
	ATK_TEXT_ATTR_UNDERLINE TextAttribute = 11
	// "true" or "false" whether the text is strikethrough
	ATK_TEXT_ATTR_STRIKETHROUGH TextAttribute = 12
	// The size of the characters in points. eg: 10
	ATK_TEXT_ATTR_SIZE TextAttribute = 13
	// The scale of the characters. The value is a string representation of a double
	ATK_TEXT_ATTR_SCALE TextAttribute = 14
	// The weight of the characters.
	ATK_TEXT_ATTR_WEIGHT TextAttribute = 15
	// The language used
	ATK_TEXT_ATTR_LANGUAGE TextAttribute = 16
	// The font family name
	ATK_TEXT_ATTR_FAMILY_NAME TextAttribute = 17
	// The background color. The value is an RGB value of the format "%u,%u,%u"
	ATK_TEXT_ATTR_BG_COLOR TextAttribute = 18
	// The foreground color. The value is an RGB value of the format "%u,%u,%u"
	ATK_TEXT_ATTR_FG_COLOR TextAttribute = 19
	// "true" if a #GdkBitmap is set for stippling the background color.
	ATK_TEXT_ATTR_BG_STIPPLE TextAttribute = 20
	// "true" if a #GdkBitmap is set for stippling the foreground color.
	ATK_TEXT_ATTR_FG_STIPPLE TextAttribute = 21
	// The wrap mode of the text, if any. Values are "none", "char", "word", or "word_char".
	ATK_TEXT_ATTR_WRAP_MODE TextAttribute = 22
	// The direction of the text, if set. Values are "none", "ltr" or "rtl"
	ATK_TEXT_ATTR_DIRECTION TextAttribute = 23
	// The justification of the text, if set. Values are "left", "right", "center" or "fill"
	ATK_TEXT_ATTR_JUSTIFICATION TextAttribute = 24
	// The stretch of the text, if set. Values are "ultra_condensed", "extra_condensed", "condensed", "semi_condensed", "normal", "semi_expanded", "expanded", "extra_expanded" or "ultra_expanded"
	ATK_TEXT_ATTR_STRETCH TextAttribute = 25
	// The capitalization variant of the text, if set. Values are "normal" or "small_caps"
	ATK_TEXT_ATTR_VARIANT TextAttribute = 26
	// The slant style of the text, if set. Values are "normal", "oblique" or "italic"
	ATK_TEXT_ATTR_STYLE TextAttribute = 27
	// not a valid text attribute, used for finding end of enumeration
	ATK_TEXT_ATTR_LAST_DEFINED TextAttribute = 28
)

// Text boundary types used for specifying boundaries for regions of text.
// This enumeration is deprecated since 2.9.4 and should not be used. Use
// AtkTextGranularity with #atk_text_get_string_at_offset instead.
type TextBoundary C.AtkTextBoundary

const (
	// Boundary is the boundary between characters
	// (including non-printing characters)
	ATK_TEXT_BOUNDARY_CHAR TextBoundary = 0
	// Boundary is the start (i.e. first character) of a word.
	ATK_TEXT_BOUNDARY_WORD_START TextBoundary = 1
	// Boundary is the end (i.e. last
	// character) of a word.
	ATK_TEXT_BOUNDARY_WORD_END TextBoundary = 2
	// Boundary is the first character in a sentence.
	ATK_TEXT_BOUNDARY_SENTENCE_START TextBoundary = 3
	// Boundary is the last (terminal)
	// character in a sentence; in languages which use "sentence stop"
	// punctuation such as English, the boundary is thus the '.', '?', or
	// similar terminal punctuation character.
	ATK_TEXT_BOUNDARY_SENTENCE_END TextBoundary = 4
	// Boundary is the initial character of the content or a
	// character immediately following a newline, linefeed, or return character.
	ATK_TEXT_BOUNDARY_LINE_START TextBoundary = 5
	// Boundary is the linefeed, or return
	// character.
	ATK_TEXT_BOUNDARY_LINE_END TextBoundary = 6
)

// Describes the type of clipping required.
type TextClipType C.AtkTextClipType

const (
	// No clipping to be done
	ATK_TEXT_CLIP_NONE TextClipType = 0
	// Text clipped by min coordinate is omitted
	ATK_TEXT_CLIP_MIN TextClipType = 1
	// Text clipped by max coordinate is omitted
	ATK_TEXT_CLIP_MAX TextClipType = 2
	// Only text fully within mix/max bound is retained
	ATK_TEXT_CLIP_BOTH TextClipType = 3
)

// Text granularity types used for specifying the granularity of the region of
// text we are interested in.
type TextGranularity C.AtkTextGranularity

const (
	// Granularity is defined by the boundaries between characters
	// (including non-printing characters)
	ATK_TEXT_GRANULARITY_CHAR TextGranularity = 0
	// Granularity is defined by the boundaries of a word,
	// starting at the beginning of the current word and finishing at the beginning of
	// the following one, if present.
	ATK_TEXT_GRANULARITY_WORD TextGranularity = 1
	// Granularity is defined by the boundaries of a sentence,
	// starting at the beginning of the current sentence and finishing at the beginning of
	// the following one, if present.
	ATK_TEXT_GRANULARITY_SENTENCE TextGranularity = 2
	// Granularity is defined by the boundaries of a line,
	// starting at the beginning of the current line and finishing at the beginning of
	// the following one, if present.
	ATK_TEXT_GRANULARITY_LINE TextGranularity = 3
	// Granularity is defined by the boundaries of a paragraph,
	// starting at the beginning of the current paragraph and finishing at the beginning of
	// the following one, if present.
	ATK_TEXT_GRANULARITY_PARAGRAPH TextGranularity = 4
)

// Default types for a given value. Those are defined in order to
// easily get localized strings to describe a given value or a given
// subrange, using atk_value_type_get_localized_name().
type ValueType C.AtkValueType

const (
	ATK_VALUE_VERY_WEAK    ValueType = 0
	ATK_VALUE_WEAK         ValueType = 1
	ATK_VALUE_ACCEPTABLE   ValueType = 2
	ATK_VALUE_STRONG       ValueType = 3
	ATK_VALUE_VERY_STRONG  ValueType = 4
	ATK_VALUE_VERY_LOW     ValueType = 5
	ATK_VALUE_LOW          ValueType = 6
	ATK_VALUE_MEDIUM       ValueType = 7
	ATK_VALUE_HIGH         ValueType = 8
	ATK_VALUE_VERY_HIGH    ValueType = 9
	ATK_VALUE_VERY_BAD     ValueType = 10
	ATK_VALUE_BAD          ValueType = 11
	ATK_VALUE_GOOD         ValueType = 12
	ATK_VALUE_VERY_GOOD    ValueType = 13
	ATK_VALUE_BEST         ValueType = 14
	ATK_VALUE_LAST_DEFINED ValueType = 15
)
