+++
title = "enums"
+++
<p class="api-heading">Align</p>
<p class="api-doc">Controls how a widget deals with extra space in a single (x or y)
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
a child or a container it is treated as @GTK_ALIGN_FILL.</p>
<div class="api-notes">
  <p class="api-ctype">GtkAlign</p>
<table>
<tr>
<td class="name">GTK_ALIGN_FILL</td>
<td class="value">0</td>
<td class="doc">stretch to fill all space if possible, center if
    no meaningful way to stretch</td>
</tr>
<tr>
<td class="name">GTK_ALIGN_START</td>
<td class="value">1</td>
<td class="doc">snap to left or top side, leaving space on right
    or bottom</td>
</tr>
<tr>
<td class="name">GTK_ALIGN_END</td>
<td class="value">2</td>
<td class="doc">snap to right or bottom side, leaving space on left
    or top</td>
</tr>
<tr>
<td class="name">GTK_ALIGN_CENTER</td>
<td class="value">3</td>
<td class="doc">center natural width of widget inside the
    allocation</td>
</tr>
<tr>
<td class="name">GTK_ALIGN_BASELINE</td>
<td class="value">4</td>
<td class="doc">align the widget according to the baseline. Since 3.10.</td>
</tr>
</table>
</div>
<p class="api-heading">ArrowPlacement</p>
<p class="api-doc">Used to specify the placement of scroll arrows in scrolling menus.</p>
<div class="api-notes">
  <p class="api-ctype">GtkArrowPlacement</p>
<table>
<tr>
<td class="name">GTK_ARROWS_BOTH</td>
<td class="value">0</td>
<td class="doc">Place one arrow on each end of the menu.</td>
</tr>
<tr>
<td class="name">GTK_ARROWS_START</td>
<td class="value">1</td>
<td class="doc">Place both arrows at the top of the menu.</td>
</tr>
<tr>
<td class="name">GTK_ARROWS_END</td>
<td class="value">2</td>
<td class="doc">Place both arrows at the bottom of the menu.</td>
</tr>
</table>
</div>
<p class="api-heading">ArrowType</p>
<p class="api-doc">Used to indicate the direction in which an arrow should point.</p>
<div class="api-notes">
  <p class="api-ctype">GtkArrowType</p>
<table>
<tr>
<td class="name">GTK_ARROW_UP</td>
<td class="value">0</td>
<td class="doc">Represents an upward pointing arrow.</td>
</tr>
<tr>
<td class="name">GTK_ARROW_DOWN</td>
<td class="value">1</td>
<td class="doc">Represents a downward pointing arrow.</td>
</tr>
<tr>
<td class="name">GTK_ARROW_LEFT</td>
<td class="value">2</td>
<td class="doc">Represents a left pointing arrow.</td>
</tr>
<tr>
<td class="name">GTK_ARROW_RIGHT</td>
<td class="value">3</td>
<td class="doc">Represents a right pointing arrow.</td>
</tr>
<tr>
<td class="name">GTK_ARROW_NONE</td>
<td class="value">4</td>
<td class="doc">No arrow. Since 2.10.</td>
</tr>
</table>
</div>
<p class="api-heading">AssistantPageType</p>
<p class="api-doc">An enum for determining the page role inside the #GtkAssistant. It's
used to handle buttons sensitivity and visibility.

Note that an assistant needs to end its page flow with a page of type
%GTK_ASSISTANT_PAGE_CONFIRM, %GTK_ASSISTANT_PAGE_SUMMARY or
%GTK_ASSISTANT_PAGE_PROGRESS to be correct.

The Cancel button will only be shown if the page isn’t “committed”.
See gtk_assistant_commit() for details.</p>
<div class="api-notes">
  <p class="api-ctype">GtkAssistantPageType</p>
<table>
<tr>
<td class="name">GTK_ASSISTANT_PAGE_CONTENT</td>
<td class="value">0</td>
<td class="doc">The page has regular contents. Both the
 Back and forward buttons will be shown.</td>
</tr>
<tr>
<td class="name">GTK_ASSISTANT_PAGE_INTRO</td>
<td class="value">1</td>
<td class="doc">The page contains an introduction to the
 assistant task. Only the Forward button will be shown if there is a
  next page.</td>
</tr>
<tr>
<td class="name">GTK_ASSISTANT_PAGE_CONFIRM</td>
<td class="value">2</td>
<td class="doc">The page lets the user confirm or deny the
 changes. The Back and Apply buttons will be shown.</td>
</tr>
<tr>
<td class="name">GTK_ASSISTANT_PAGE_SUMMARY</td>
<td class="value">3</td>
<td class="doc">The page informs the user of the changes
 done. Only the Close button will be shown.</td>
</tr>
<tr>
<td class="name">GTK_ASSISTANT_PAGE_PROGRESS</td>
<td class="value">4</td>
<td class="doc">Used for tasks that take a long time to
 complete, blocks the assistant until the page is marked as complete.
  Only the back button will be shown.</td>
</tr>
<tr>
<td class="name">GTK_ASSISTANT_PAGE_CUSTOM</td>
<td class="value">5</td>
<td class="doc">Used for when other page types are not
 appropriate. No buttons will be shown, and the application must
 add its own buttons through gtk_assistant_add_action_widget().</td>
</tr>
</table>
</div>
<p class="api-heading">BaselinePosition</p>
<p class="api-doc">Whenever a container has some form of natural row it may align
children in that row along a common typographical baseline. If
the amount of verical space in the row is taller than the total
requested height of the baseline-aligned children then it can use a
#GtkBaselinePosition to select where to put the baseline inside the
extra availible space.</p>
<div class="api-notes">
  <p class="api-ctype">GtkBaselinePosition</p>
  <p class="api-since">since 3.10</p>
<table>
<tr>
<td class="name">GTK_BASELINE_POSITION_TOP</td>
<td class="value">0</td>
<td class="doc">Align the baseline at the top</td>
</tr>
<tr>
<td class="name">GTK_BASELINE_POSITION_CENTER</td>
<td class="value">1</td>
<td class="doc">Center the baseline</td>
</tr>
<tr>
<td class="name">GTK_BASELINE_POSITION_BOTTOM</td>
<td class="value">2</td>
<td class="doc">Align the baseline at the bottom</td>
</tr>
</table>
</div>
<p class="api-heading">BorderStyle</p>
<p class="api-doc">Describes how the border of a UI element should be rendered.</p>
<div class="api-notes">
  <p class="api-ctype">GtkBorderStyle</p>
<table>
<tr>
<td class="name">GTK_BORDER_STYLE_NONE</td>
<td class="value">0</td>
<td class="doc">No visible border</td>
</tr>
<tr>
<td class="name">GTK_BORDER_STYLE_SOLID</td>
<td class="value">1</td>
<td class="doc">A single line segment</td>
</tr>
<tr>
<td class="name">GTK_BORDER_STYLE_INSET</td>
<td class="value">2</td>
<td class="doc">Looks as if the content is sunken into the canvas</td>
</tr>
<tr>
<td class="name">GTK_BORDER_STYLE_OUTSET</td>
<td class="value">3</td>
<td class="doc">Looks as if the content is coming out of the canvas</td>
</tr>
<tr>
<td class="name">GTK_BORDER_STYLE_HIDDEN</td>
<td class="value">4</td>
<td class="doc">Same as @GTK_BORDER_STYLE_NONE</td>
</tr>
<tr>
<td class="name">GTK_BORDER_STYLE_DOTTED</td>
<td class="value">5</td>
<td class="doc">A series of round dots</td>
</tr>
<tr>
<td class="name">GTK_BORDER_STYLE_DASHED</td>
<td class="value">6</td>
<td class="doc">A series of square-ended dashes</td>
</tr>
<tr>
<td class="name">GTK_BORDER_STYLE_DOUBLE</td>
<td class="value">7</td>
<td class="doc">Two parallel lines with some space between them</td>
</tr>
<tr>
<td class="name">GTK_BORDER_STYLE_GROOVE</td>
<td class="value">8</td>
<td class="doc">Looks as if it were carved in the canvas</td>
</tr>
<tr>
<td class="name">GTK_BORDER_STYLE_RIDGE</td>
<td class="value">9</td>
<td class="doc">Looks as if it were coming out of the canvas</td>
</tr>
</table>
</div>
<p class="api-heading">BuilderError</p>
<p class="api-doc">Error codes that identify various errors that can occur while using
#GtkBuilder.</p>
<div class="api-notes">
  <p class="api-ctype">GtkBuilderError</p>
<table>
<tr>
<td class="name">GTK_BUILDER_ERROR_INVALID_TYPE_FUNCTION</td>
<td class="value">0</td>
<td class="doc">A type-func attribute didn’t name
 a function that returns a #GType.</td>
</tr>
<tr>
<td class="name">GTK_BUILDER_ERROR_UNHANDLED_TAG</td>
<td class="value">1</td>
<td class="doc">The input contained a tag that #GtkBuilder
 can’t handle.</td>
</tr>
<tr>
<td class="name">GTK_BUILDER_ERROR_MISSING_ATTRIBUTE</td>
<td class="value">2</td>
<td class="doc">An attribute that is required by
 #GtkBuilder was missing.</td>
</tr>
<tr>
<td class="name">GTK_BUILDER_ERROR_INVALID_ATTRIBUTE</td>
<td class="value">3</td>
<td class="doc">#GtkBuilder found an attribute that
 it doesn’t understand.</td>
</tr>
<tr>
<td class="name">GTK_BUILDER_ERROR_INVALID_TAG</td>
<td class="value">4</td>
<td class="doc">#GtkBuilder found a tag that
 it doesn’t understand.</td>
</tr>
<tr>
<td class="name">GTK_BUILDER_ERROR_MISSING_PROPERTY_VALUE</td>
<td class="value">5</td>
<td class="doc">A required property value was
 missing.</td>
</tr>
<tr>
<td class="name">GTK_BUILDER_ERROR_INVALID_VALUE</td>
<td class="value">6</td>
<td class="doc">#GtkBuilder couldn’t parse
 some attribute value.</td>
</tr>
<tr>
<td class="name">GTK_BUILDER_ERROR_VERSION_MISMATCH</td>
<td class="value">7</td>
<td class="doc">The input file requires a newer version
 of GTK+.</td>
</tr>
<tr>
<td class="name">GTK_BUILDER_ERROR_DUPLICATE_ID</td>
<td class="value">8</td>
<td class="doc">An object id occurred twice.</td>
</tr>
<tr>
<td class="name">GTK_BUILDER_ERROR_OBJECT_TYPE_REFUSED</td>
<td class="value">9</td>
<td class="doc">A specified object type is of the same type or
 derived from the type of the composite class being extended with builder XML.</td>
</tr>
<tr>
<td class="name">GTK_BUILDER_ERROR_TEMPLATE_MISMATCH</td>
<td class="value">10</td>
<td class="doc">The wrong type was specified in a composite class’s template XML</td>
</tr>
<tr>
<td class="name">GTK_BUILDER_ERROR_INVALID_PROPERTY</td>
<td class="value">11</td>
<td class="doc">The specified property is unknown for the object class.</td>
</tr>
<tr>
<td class="name">GTK_BUILDER_ERROR_INVALID_SIGNAL</td>
<td class="value">12</td>
<td class="doc">The specified signal is unknown for the object class.</td>
</tr>
<tr>
<td class="name">GTK_BUILDER_ERROR_INVALID_ID</td>
<td class="value">13</td>
<td class="doc">An object id is unknown</td>
</tr>
</table>
</div>
<p class="api-heading">ButtonBoxStyle</p>
<p class="api-doc">Used to dictate the style that a #GtkButtonBox uses to layout the buttons it
contains.</p>
<div class="api-notes">
  <p class="api-ctype">GtkButtonBoxStyle</p>
<table>
<tr>
<td class="name">GTK_BUTTONBOX_SPREAD</td>
<td class="value">1</td>
<td class="doc">Buttons are evenly spread across the box.</td>
</tr>
<tr>
<td class="name">GTK_BUTTONBOX_EDGE</td>
<td class="value">2</td>
<td class="doc">Buttons are placed at the edges of the box.</td>
</tr>
<tr>
<td class="name">GTK_BUTTONBOX_START</td>
<td class="value">3</td>
<td class="doc">Buttons are grouped towards the start of the box,
  (on the left for a HBox, or the top for a VBox).</td>
</tr>
<tr>
<td class="name">GTK_BUTTONBOX_END</td>
<td class="value">4</td>
<td class="doc">Buttons are grouped towards the end of the box,
  (on the right for a HBox, or the bottom for a VBox).</td>
</tr>
<tr>
<td class="name">GTK_BUTTONBOX_CENTER</td>
<td class="value">5</td>
<td class="doc">Buttons are centered in the box. Since 2.12.</td>
</tr>
<tr>
<td class="name">GTK_BUTTONBOX_EXPAND</td>
<td class="value">6</td>
<td class="doc">Buttons expand to fill the box. This entails giving
  buttons a "linked" appearance, making button sizes homogeneous, and
  setting spacing to 0 (same as calling gtk_box_set_homogeneous() and
  gtk_box_set_spacing() manually). Since 3.12.</td>
</tr>
</table>
</div>
<p class="api-heading">ButtonRole</p>
<p class="api-doc">The role specifies the desired appearance of a #GtkModelButton.</p>
<div class="api-notes">
  <p class="api-ctype">GtkButtonRole</p>
<table>
<tr>
<td class="name">GTK_BUTTON_ROLE_NORMAL</td>
<td class="value">0</td>
<td class="doc">A plain button</td>
</tr>
<tr>
<td class="name">GTK_BUTTON_ROLE_CHECK</td>
<td class="value">1</td>
<td class="doc">A check button</td>
</tr>
<tr>
<td class="name">GTK_BUTTON_ROLE_RADIO</td>
<td class="value">2</td>
<td class="doc">A radio button</td>
</tr>
</table>
</div>
<p class="api-heading">ButtonsType</p>
<p class="api-doc">Prebuilt sets of buttons for the dialog. If
none of these choices are appropriate, simply use %GTK_BUTTONS_NONE
then call gtk_dialog_add_buttons().

> Please note that %GTK_BUTTONS_OK, %GTK_BUTTONS_YES_NO
> and %GTK_BUTTONS_OK_CANCEL are discouraged by the
> [GNOME Human Interface Guidelines](http://library.gnome.org/devel/hig-book/stable/).</p>
<div class="api-notes">
  <p class="api-ctype">GtkButtonsType</p>
<table>
<tr>
<td class="name">GTK_BUTTONS_NONE</td>
<td class="value">0</td>
<td class="doc">no buttons at all</td>
</tr>
<tr>
<td class="name">GTK_BUTTONS_OK</td>
<td class="value">1</td>
<td class="doc">an OK button</td>
</tr>
<tr>
<td class="name">GTK_BUTTONS_CLOSE</td>
<td class="value">2</td>
<td class="doc">a Close button</td>
</tr>
<tr>
<td class="name">GTK_BUTTONS_CANCEL</td>
<td class="value">3</td>
<td class="doc">a Cancel button</td>
</tr>
<tr>
<td class="name">GTK_BUTTONS_YES_NO</td>
<td class="value">4</td>
<td class="doc">Yes and No buttons</td>
</tr>
<tr>
<td class="name">GTK_BUTTONS_OK_CANCEL</td>
<td class="value">5</td>
<td class="doc">OK and Cancel buttons</td>
</tr>
</table>
</div>
<p class="api-heading">CellRendererAccelMode</p>
<p class="api-doc">Determines if the edited accelerators are GTK+ accelerators. If
they are, consumed modifiers are suppressed, only accelerators
accepted by GTK+ are allowed, and the accelerators are rendered
in the same way as they are in menus.</p>
<div class="api-notes">
  <p class="api-ctype">GtkCellRendererAccelMode</p>
<table>
<tr>
<td class="name">GTK_CELL_RENDERER_ACCEL_MODE_GTK</td>
<td class="value">0</td>
<td class="doc">GTK+ accelerators mode</td>
</tr>
<tr>
<td class="name">GTK_CELL_RENDERER_ACCEL_MODE_OTHER</td>
<td class="value">1</td>
<td class="doc">Other accelerator mode
GTK_CELL_RENDERER_ACCEL_MODE_MODIFIER_TAP: Bare modifiers mode</td>
</tr>
<tr>
<td class="name">GTK_CELL_RENDERER_ACCEL_MODE_MODIFIER_TAP</td>
<td class="value">2</td>
</tr>
</table>
</div>
<p class="api-heading">CellRendererMode</p>
<p class="api-doc">Identifies how the user can interact with a particular cell.</p>
<div class="api-notes">
  <p class="api-ctype">GtkCellRendererMode</p>
<table>
<tr>
<td class="name">GTK_CELL_RENDERER_MODE_INERT</td>
<td class="value">0</td>
<td class="doc">The cell is just for display
 and cannot be interacted with.  Note that this doesn’t mean that eg. the
 row being drawn can’t be selected -- just that a particular element of
 it cannot be individually modified.</td>
</tr>
<tr>
<td class="name">GTK_CELL_RENDERER_MODE_ACTIVATABLE</td>
<td class="value">1</td>
<td class="doc">The cell can be clicked.</td>
</tr>
<tr>
<td class="name">GTK_CELL_RENDERER_MODE_EDITABLE</td>
<td class="value">2</td>
<td class="doc">The cell can be edited or otherwise modified.</td>
</tr>
</table>
</div>
<p class="api-heading">CornerType</p>
<p class="api-doc">Specifies which corner a child widget should be placed in when packed into
a #GtkScrolledWindow. This is effectively the opposite of where the scroll
bars are placed.</p>
<div class="api-notes">
  <p class="api-ctype">GtkCornerType</p>
<table>
<tr>
<td class="name">GTK_CORNER_TOP_LEFT</td>
<td class="value">0</td>
<td class="doc">Place the scrollbars on the right and bottom of the
 widget (default behaviour).</td>
</tr>
<tr>
<td class="name">GTK_CORNER_BOTTOM_LEFT</td>
<td class="value">1</td>
<td class="doc">Place the scrollbars on the top and right of the
 widget.</td>
</tr>
<tr>
<td class="name">GTK_CORNER_TOP_RIGHT</td>
<td class="value">2</td>
<td class="doc">Place the scrollbars on the left and bottom of the
 widget.</td>
</tr>
<tr>
<td class="name">GTK_CORNER_BOTTOM_RIGHT</td>
<td class="value">3</td>
<td class="doc">Place the scrollbars on the top and left of the
 widget.</td>
</tr>
</table>
</div>
<p class="api-heading">CssProviderError</p>
<p class="api-doc">Error codes for %GTK_CSS_PROVIDER_ERROR.</p>
<div class="api-notes">
  <p class="api-ctype">GtkCssProviderError</p>
<table>
<tr>
<td class="name">GTK_CSS_PROVIDER_ERROR_FAILED</td>
<td class="value">0</td>
<td class="doc">Failed.</td>
</tr>
<tr>
<td class="name">GTK_CSS_PROVIDER_ERROR_SYNTAX</td>
<td class="value">1</td>
<td class="doc">Syntax error.</td>
</tr>
<tr>
<td class="name">GTK_CSS_PROVIDER_ERROR_IMPORT</td>
<td class="value">2</td>
<td class="doc">Import error.</td>
</tr>
<tr>
<td class="name">GTK_CSS_PROVIDER_ERROR_NAME</td>
<td class="value">3</td>
<td class="doc">Name error.</td>
</tr>
<tr>
<td class="name">GTK_CSS_PROVIDER_ERROR_DEPRECATED</td>
<td class="value">4</td>
<td class="doc">Deprecation error.</td>
</tr>
<tr>
<td class="name">GTK_CSS_PROVIDER_ERROR_UNKNOWN_VALUE</td>
<td class="value">5</td>
<td class="doc">Unknown value.</td>
</tr>
</table>
</div>
<p class="api-heading">CssSectionType</p>
<p class="api-doc">The different types of sections indicate parts of a CSS document as
parsed by GTK’s CSS parser. They are oriented towards the
[CSS Grammar](http://www.w3.org/TR/CSS21/grammar.html),
but may contain extensions.

More types might be added in the future as the parser incorporates
more features.</p>
<div class="api-notes">
  <p class="api-ctype">GtkCssSectionType</p>
  <p class="api-since">since 3.2</p>
<table>
<tr>
<td class="name">GTK_CSS_SECTION_DOCUMENT</td>
<td class="value">0</td>
<td class="doc">The section describes a complete document.
  This section time is the only one where gtk_css_section_get_parent()
  might return %NULL.</td>
</tr>
<tr>
<td class="name">GTK_CSS_SECTION_IMPORT</td>
<td class="value">1</td>
<td class="doc">The section defines an import rule.</td>
</tr>
<tr>
<td class="name">GTK_CSS_SECTION_COLOR_DEFINITION</td>
<td class="value">2</td>
<td class="doc">The section defines a color. This
  is a GTK extension to CSS.</td>
</tr>
<tr>
<td class="name">GTK_CSS_SECTION_BINDING_SET</td>
<td class="value">3</td>
<td class="doc">The section defines a binding set. This
  is a GTK extension to CSS.</td>
</tr>
<tr>
<td class="name">GTK_CSS_SECTION_RULESET</td>
<td class="value">4</td>
<td class="doc">The section defines a CSS ruleset.</td>
</tr>
<tr>
<td class="name">GTK_CSS_SECTION_SELECTOR</td>
<td class="value">5</td>
<td class="doc">The section defines a CSS selector.</td>
</tr>
<tr>
<td class="name">GTK_CSS_SECTION_DECLARATION</td>
<td class="value">6</td>
<td class="doc">The section defines the declaration of
  a CSS variable.</td>
</tr>
<tr>
<td class="name">GTK_CSS_SECTION_VALUE</td>
<td class="value">7</td>
<td class="doc">The section defines the value of a CSS declaration.</td>
</tr>
<tr>
<td class="name">GTK_CSS_SECTION_KEYFRAMES</td>
<td class="value">8</td>
<td class="doc">The section defines keyframes. See [CSS
  Animations](http://dev.w3.org/csswg/css3-animations/#keyframes) for details. Since 3.6</td>
</tr>
</table>
</div>
<p class="api-heading">DeleteType</p>
<p class="api-doc">See also: #GtkEntry::delete-from-cursor.</p>
<div class="api-notes">
  <p class="api-ctype">GtkDeleteType</p>
<table>
<tr>
<td class="name">GTK_DELETE_CHARS</td>
<td class="value">0</td>
<td class="doc">Delete characters.</td>
</tr>
<tr>
<td class="name">GTK_DELETE_WORD_ENDS</td>
<td class="value">1</td>
<td class="doc">Delete only the portion of the word to the
  left/right of cursor if we’re in the middle of a word.</td>
</tr>
<tr>
<td class="name">GTK_DELETE_WORDS</td>
<td class="value">2</td>
<td class="doc">Delete words.</td>
</tr>
<tr>
<td class="name">GTK_DELETE_DISPLAY_LINES</td>
<td class="value">3</td>
<td class="doc">Delete display-lines. Display-lines
  refers to the visible lines, with respect to to the current line
  breaks. As opposed to paragraphs, which are defined by line
  breaks in the input.</td>
</tr>
<tr>
<td class="name">GTK_DELETE_DISPLAY_LINE_ENDS</td>
<td class="value">4</td>
<td class="doc">Delete only the portion of the
  display-line to the left/right of cursor.</td>
</tr>
<tr>
<td class="name">GTK_DELETE_PARAGRAPH_ENDS</td>
<td class="value">5</td>
<td class="doc">Delete to the end of the
  paragraph. Like C-k in Emacs (or its reverse).</td>
</tr>
<tr>
<td class="name">GTK_DELETE_PARAGRAPHS</td>
<td class="value">6</td>
<td class="doc">Delete entire line. Like C-k in pico.</td>
</tr>
<tr>
<td class="name">GTK_DELETE_WHITESPACE</td>
<td class="value">7</td>
<td class="doc">Delete only whitespace. Like M-\ in Emacs.</td>
</tr>
</table>
</div>
<p class="api-heading">DirectionType</p>
<p class="api-doc">Focus movement types.</p>
<div class="api-notes">
  <p class="api-ctype">GtkDirectionType</p>
<table>
<tr>
<td class="name">GTK_DIR_TAB_FORWARD</td>
<td class="value">0</td>
<td class="doc">Move forward.</td>
</tr>
<tr>
<td class="name">GTK_DIR_TAB_BACKWARD</td>
<td class="value">1</td>
<td class="doc">Move backward.</td>
</tr>
<tr>
<td class="name">GTK_DIR_UP</td>
<td class="value">2</td>
<td class="doc">Move up.</td>
</tr>
<tr>
<td class="name">GTK_DIR_DOWN</td>
<td class="value">3</td>
<td class="doc">Move down.</td>
</tr>
<tr>
<td class="name">GTK_DIR_LEFT</td>
<td class="value">4</td>
<td class="doc">Move left.</td>
</tr>
<tr>
<td class="name">GTK_DIR_RIGHT</td>
<td class="value">5</td>
<td class="doc">Move right.</td>
</tr>
</table>
</div>
<p class="api-heading">DragResult</p>
<p class="api-doc">Gives an indication why a drag operation failed.
The value can by obtained by connecting to the
#GtkWidget::drag-failed signal.</p>
<div class="api-notes">
  <p class="api-ctype">GtkDragResult</p>
<table>
<tr>
<td class="name">GTK_DRAG_RESULT_SUCCESS</td>
<td class="value">0</td>
<td class="doc">The drag operation was successful.</td>
</tr>
<tr>
<td class="name">GTK_DRAG_RESULT_NO_TARGET</td>
<td class="value">1</td>
<td class="doc">No suitable drag target.</td>
</tr>
<tr>
<td class="name">GTK_DRAG_RESULT_USER_CANCELLED</td>
<td class="value">2</td>
<td class="doc">The user cancelled the drag operation.</td>
</tr>
<tr>
<td class="name">GTK_DRAG_RESULT_TIMEOUT_EXPIRED</td>
<td class="value">3</td>
<td class="doc">The drag operation timed out.</td>
</tr>
<tr>
<td class="name">GTK_DRAG_RESULT_GRAB_BROKEN</td>
<td class="value">4</td>
<td class="doc">The pointer or keyboard grab used
 for the drag operation was broken.</td>
</tr>
<tr>
<td class="name">GTK_DRAG_RESULT_ERROR</td>
<td class="value">5</td>
<td class="doc">The drag operation failed due to some
 unspecified error.</td>
</tr>
</table>
</div>
<p class="api-heading">EntryIconPosition</p>
<p class="api-doc">Specifies the side of the entry at which an icon is placed.</p>
<div class="api-notes">
  <p class="api-ctype">GtkEntryIconPosition</p>
  <p class="api-since">since 2.16</p>
<table>
<tr>
<td class="name">GTK_ENTRY_ICON_PRIMARY</td>
<td class="value">0</td>
<td class="doc">At the beginning of the entry (depending on the text direction).</td>
</tr>
<tr>
<td class="name">GTK_ENTRY_ICON_SECONDARY</td>
<td class="value">1</td>
<td class="doc">At the end of the entry (depending on the text direction).</td>
</tr>
</table>
</div>
<p class="api-heading">EventSequenceState</p>
<p class="api-doc">Describes the state of a #GdkEventSequence in a #GtkGesture.</p>
<div class="api-notes">
  <p class="api-ctype">GtkEventSequenceState</p>
  <p class="api-since">since 3.14</p>
<table>
<tr>
<td class="name">GTK_EVENT_SEQUENCE_NONE</td>
<td class="value">0</td>
<td class="doc">The sequence is handled, but not grabbed.</td>
</tr>
<tr>
<td class="name">GTK_EVENT_SEQUENCE_CLAIMED</td>
<td class="value">1</td>
<td class="doc">The sequence is handled and grabbed.</td>
</tr>
<tr>
<td class="name">GTK_EVENT_SEQUENCE_DENIED</td>
<td class="value">2</td>
<td class="doc">The sequence is denied.</td>
</tr>
</table>
</div>
<p class="api-heading">ExpanderStyle</p>
<p class="api-doc">Used to specify the style of the expanders drawn by a #GtkTreeView.</p>
<div class="api-notes">
  <p class="api-ctype">GtkExpanderStyle</p>
<table>
<tr>
<td class="name">GTK_EXPANDER_COLLAPSED</td>
<td class="value">0</td>
<td class="doc">The style used for a collapsed subtree.</td>
</tr>
<tr>
<td class="name">GTK_EXPANDER_SEMI_COLLAPSED</td>
<td class="value">1</td>
<td class="doc">Intermediate style used during animation.</td>
</tr>
<tr>
<td class="name">GTK_EXPANDER_SEMI_EXPANDED</td>
<td class="value">2</td>
<td class="doc">Intermediate style used during animation.</td>
</tr>
<tr>
<td class="name">GTK_EXPANDER_EXPANDED</td>
<td class="value">3</td>
<td class="doc">The style used for an expanded subtree.</td>
</tr>
</table>
</div>
<p class="api-heading">FileChooserAction</p>
<p class="api-doc">Describes whether a #GtkFileChooser is being used to open existing files
or to save to a possibly new file.</p>
<div class="api-notes">
  <p class="api-ctype">GtkFileChooserAction</p>
<table>
<tr>
<td class="name">GTK_FILE_CHOOSER_ACTION_OPEN</td>
<td class="value">0</td>
<td class="doc">Indicates open mode.  The file chooser
 will only let the user pick an existing file.</td>
</tr>
<tr>
<td class="name">GTK_FILE_CHOOSER_ACTION_SAVE</td>
<td class="value">1</td>
<td class="doc">Indicates save mode.  The file chooser
 will let the user pick an existing file, or type in a new
 filename.</td>
</tr>
<tr>
<td class="name">GTK_FILE_CHOOSER_ACTION_SELECT_FOLDER</td>
<td class="value">2</td>
<td class="doc">Indicates an Open mode for
 selecting folders.  The file chooser will let the user pick an
 existing folder.</td>
</tr>
<tr>
<td class="name">GTK_FILE_CHOOSER_ACTION_CREATE_FOLDER</td>
<td class="value">3</td>
<td class="doc">Indicates a mode for creating a
 new folder.  The file chooser will let the user name an existing or
 new folder.</td>
</tr>
</table>
</div>
<p class="api-heading">FileChooserConfirmation</p>
<p class="api-doc">Used as a return value of handlers for the
#GtkFileChooser::confirm-overwrite signal of a #GtkFileChooser. This
value determines whether the file chooser will present the stock
confirmation dialog, accept the user’s choice of a filename, or
let the user choose another filename.</p>
<div class="api-notes">
  <p class="api-ctype">GtkFileChooserConfirmation</p>
  <p class="api-since">since 2.8</p>
<table>
<tr>
<td class="name">GTK_FILE_CHOOSER_CONFIRMATION_CONFIRM</td>
<td class="value">0</td>
<td class="doc">The file chooser will present
 its stock dialog to confirm about overwriting an existing file.</td>
</tr>
<tr>
<td class="name">GTK_FILE_CHOOSER_CONFIRMATION_ACCEPT_FILENAME</td>
<td class="value">1</td>
<td class="doc">The file chooser will
 terminate and accept the user’s choice of a file name.</td>
</tr>
<tr>
<td class="name">GTK_FILE_CHOOSER_CONFIRMATION_SELECT_AGAIN</td>
<td class="value">2</td>
<td class="doc">The file chooser will
 continue running, so as to let the user select another file name.</td>
</tr>
</table>
</div>
<p class="api-heading">FileChooserError</p>
<p class="api-doc">These identify the various errors that can occur while calling
#GtkFileChooser functions.</p>
<div class="api-notes">
  <p class="api-ctype">GtkFileChooserError</p>
<table>
<tr>
<td class="name">GTK_FILE_CHOOSER_ERROR_NONEXISTENT</td>
<td class="value">0</td>
<td class="doc">Indicates that a file does not exist.</td>
</tr>
<tr>
<td class="name">GTK_FILE_CHOOSER_ERROR_BAD_FILENAME</td>
<td class="value">1</td>
<td class="doc">Indicates a malformed filename.</td>
</tr>
<tr>
<td class="name">GTK_FILE_CHOOSER_ERROR_ALREADY_EXISTS</td>
<td class="value">2</td>
<td class="doc">Indicates a duplicate path (e.g. when
 adding a bookmark).</td>
</tr>
<tr>
<td class="name">GTK_FILE_CHOOSER_ERROR_INCOMPLETE_HOSTNAME</td>
<td class="value">3</td>
<td class="doc">Indicates an incomplete hostname (e.g. "http://foo" without a slash after that).</td>
</tr>
</table>
</div>
<p class="api-heading">IMPreeditStyle</p>
<p class="api-doc">Style for input method preedit. See also
#GtkSettings:gtk-im-preedit-style</p>
<div class="api-notes">
  <p class="api-ctype">GtkIMPreeditStyle</p>
<table>
<tr>
<td class="name">GTK_IM_PREEDIT_NOTHING</td>
<td class="value">0</td>
<td class="doc">Deprecated</td>
</tr>
<tr>
<td class="name">GTK_IM_PREEDIT_CALLBACK</td>
<td class="value">1</td>
<td class="doc">Deprecated</td>
</tr>
<tr>
<td class="name">GTK_IM_PREEDIT_NONE</td>
<td class="value">2</td>
<td class="doc">Deprecated</td>
</tr>
</table>
</div>
<p class="api-heading">IMStatusStyle</p>
<p class="api-doc">Style for input method status. See also
#GtkSettings:gtk-im-status-style</p>
<div class="api-notes">
  <p class="api-ctype">GtkIMStatusStyle</p>
<table>
<tr>
<td class="name">GTK_IM_STATUS_NOTHING</td>
<td class="value">0</td>
<td class="doc">Deprecated</td>
</tr>
<tr>
<td class="name">GTK_IM_STATUS_CALLBACK</td>
<td class="value">1</td>
<td class="doc">Deprecated</td>
</tr>
<tr>
<td class="name">GTK_IM_STATUS_NONE</td>
<td class="value">2</td>
<td class="doc">Deprecated</td>
</tr>
</table>
</div>
<p class="api-heading">IconSize</p>
<p class="api-doc">Built-in stock icon sizes.</p>
<div class="api-notes">
  <p class="api-ctype">GtkIconSize</p>
<table>
<tr>
<td class="name">GTK_ICON_SIZE_INVALID</td>
<td class="value">0</td>
<td class="doc">Invalid size.</td>
</tr>
<tr>
<td class="name">GTK_ICON_SIZE_MENU</td>
<td class="value">1</td>
<td class="doc">Size appropriate for menus (16px).</td>
</tr>
<tr>
<td class="name">GTK_ICON_SIZE_SMALL_TOOLBAR</td>
<td class="value">2</td>
<td class="doc">Size appropriate for small toolbars (16px).</td>
</tr>
<tr>
<td class="name">GTK_ICON_SIZE_LARGE_TOOLBAR</td>
<td class="value">3</td>
<td class="doc">Size appropriate for large toolbars (24px)</td>
</tr>
<tr>
<td class="name">GTK_ICON_SIZE_BUTTON</td>
<td class="value">4</td>
<td class="doc">Size appropriate for buttons (16px)</td>
</tr>
<tr>
<td class="name">GTK_ICON_SIZE_DND</td>
<td class="value">5</td>
<td class="doc">Size appropriate for drag and drop (32px)</td>
</tr>
<tr>
<td class="name">GTK_ICON_SIZE_DIALOG</td>
<td class="value">6</td>
<td class="doc">Size appropriate for dialogs (48px)</td>
</tr>
</table>
</div>
<p class="api-heading">IconThemeError</p>
<p class="api-doc">Error codes for GtkIconTheme operations.</p>
<div class="api-notes">
  <p class="api-ctype">GtkIconThemeError</p>
<table>
<tr>
<td class="name">GTK_ICON_THEME_NOT_FOUND</td>
<td class="value">0</td>
<td class="doc">The icon specified does not exist in the theme</td>
</tr>
<tr>
<td class="name">GTK_ICON_THEME_FAILED</td>
<td class="value">1</td>
<td class="doc">An unspecified error occurred.</td>
</tr>
</table>
</div>
<p class="api-heading">IconViewDropPosition</p>
<p class="api-doc">An enum for determining where a dropped item goes.</p>
<div class="api-notes">
  <p class="api-ctype">GtkIconViewDropPosition</p>
<table>
<tr>
<td class="name">GTK_ICON_VIEW_NO_DROP</td>
<td class="value">0</td>
<td class="doc">no drop possible</td>
</tr>
<tr>
<td class="name">GTK_ICON_VIEW_DROP_INTO</td>
<td class="value">1</td>
<td class="doc">dropped item replaces the item</td>
</tr>
<tr>
<td class="name">GTK_ICON_VIEW_DROP_LEFT</td>
<td class="value">2</td>
<td class="doc">droppped item is inserted to the left</td>
</tr>
<tr>
<td class="name">GTK_ICON_VIEW_DROP_RIGHT</td>
<td class="value">3</td>
<td class="doc">dropped item is inserted to the right</td>
</tr>
<tr>
<td class="name">GTK_ICON_VIEW_DROP_ABOVE</td>
<td class="value">4</td>
<td class="doc">dropped item is inserted above</td>
</tr>
<tr>
<td class="name">GTK_ICON_VIEW_DROP_BELOW</td>
<td class="value">5</td>
<td class="doc">dropped item is inserted below</td>
</tr>
</table>
</div>
<p class="api-heading">ImageType</p>
<p class="api-doc">Describes the image data representation used by a #GtkImage. If you
want to get the image from the widget, you can only get the
currently-stored representation. e.g.  if the
gtk_image_get_storage_type() returns #GTK_IMAGE_PIXBUF, then you can
call gtk_image_get_pixbuf() but not gtk_image_get_stock().  For empty
images, you can request any storage type (call any of the "get"
functions), but they will all return %NULL values.</p>
<div class="api-notes">
  <p class="api-ctype">GtkImageType</p>
<table>
<tr>
<td class="name">GTK_IMAGE_EMPTY</td>
<td class="value">0</td>
<td class="doc">there is no image displayed by the widget</td>
</tr>
<tr>
<td class="name">GTK_IMAGE_PIXBUF</td>
<td class="value">1</td>
<td class="doc">the widget contains a #GdkPixbuf</td>
</tr>
<tr>
<td class="name">GTK_IMAGE_STOCK</td>
<td class="value">2</td>
<td class="doc">the widget contains a [stock item name][gtkstock]</td>
</tr>
<tr>
<td class="name">GTK_IMAGE_ICON_SET</td>
<td class="value">3</td>
<td class="doc">the widget contains a #GtkIconSet</td>
</tr>
<tr>
<td class="name">GTK_IMAGE_ANIMATION</td>
<td class="value">4</td>
<td class="doc">the widget contains a #GdkPixbufAnimation</td>
</tr>
<tr>
<td class="name">GTK_IMAGE_ICON_NAME</td>
<td class="value">5</td>
<td class="doc">the widget contains a named icon.
 This image type was added in GTK+ 2.6</td>
</tr>
<tr>
<td class="name">GTK_IMAGE_GICON</td>
<td class="value">6</td>
<td class="doc">the widget contains a #GIcon.
 This image type was added in GTK+ 2.14</td>
</tr>
<tr>
<td class="name">GTK_IMAGE_SURFACE</td>
<td class="value">7</td>
<td class="doc">the widget contains a #cairo_surface_t.
 This image type was added in GTK+ 3.10</td>
</tr>
</table>
</div>
<p class="api-heading">InputPurpose</p>
<p class="api-doc">Describes primary purpose of the input widget. This information is
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
interpret unknown values as “free form”.</p>
<div class="api-notes">
  <p class="api-ctype">GtkInputPurpose</p>
  <p class="api-since">since 3.6</p>
<table>
<tr>
<td class="name">GTK_INPUT_PURPOSE_FREE_FORM</td>
<td class="value">0</td>
<td class="doc">Allow any character</td>
</tr>
<tr>
<td class="name">GTK_INPUT_PURPOSE_ALPHA</td>
<td class="value">1</td>
<td class="doc">Allow only alphabetic characters</td>
</tr>
<tr>
<td class="name">GTK_INPUT_PURPOSE_DIGITS</td>
<td class="value">2</td>
<td class="doc">Allow only digits</td>
</tr>
<tr>
<td class="name">GTK_INPUT_PURPOSE_NUMBER</td>
<td class="value">3</td>
<td class="doc">Edited field expects numbers</td>
</tr>
<tr>
<td class="name">GTK_INPUT_PURPOSE_PHONE</td>
<td class="value">4</td>
<td class="doc">Edited field expects phone number</td>
</tr>
<tr>
<td class="name">GTK_INPUT_PURPOSE_URL</td>
<td class="value">5</td>
<td class="doc">Edited field expects URL</td>
</tr>
<tr>
<td class="name">GTK_INPUT_PURPOSE_EMAIL</td>
<td class="value">6</td>
<td class="doc">Edited field expects email address</td>
</tr>
<tr>
<td class="name">GTK_INPUT_PURPOSE_NAME</td>
<td class="value">7</td>
<td class="doc">Edited field expects the name of a person</td>
</tr>
<tr>
<td class="name">GTK_INPUT_PURPOSE_PASSWORD</td>
<td class="value">8</td>
<td class="doc">Like @GTK_INPUT_PURPOSE_FREE_FORM, but characters are hidden</td>
</tr>
<tr>
<td class="name">GTK_INPUT_PURPOSE_PIN</td>
<td class="value">9</td>
<td class="doc">Like @GTK_INPUT_PURPOSE_DIGITS, but characters are hidden</td>
</tr>
</table>
</div>
<p class="api-heading">Justification</p>
<p class="api-doc">Used for justifying the text inside a #GtkLabel widget. (See also
#GtkAlignment).</p>
<div class="api-notes">
  <p class="api-ctype">GtkJustification</p>
<table>
<tr>
<td class="name">GTK_JUSTIFY_LEFT</td>
<td class="value">0</td>
<td class="doc">The text is placed at the left edge of the label.</td>
</tr>
<tr>
<td class="name">GTK_JUSTIFY_RIGHT</td>
<td class="value">1</td>
<td class="doc">The text is placed at the right edge of the label.</td>
</tr>
<tr>
<td class="name">GTK_JUSTIFY_CENTER</td>
<td class="value">2</td>
<td class="doc">The text is placed in the center of the label.</td>
</tr>
<tr>
<td class="name">GTK_JUSTIFY_FILL</td>
<td class="value">3</td>
<td class="doc">The text is placed is distributed across the label.</td>
</tr>
</table>
</div>
<p class="api-heading">LevelBarMode</p>
<p class="api-doc">Describes how #GtkLevelBar contents should be rendered.
Note that this enumeration could be extended with additional modes
in the future.</p>
<div class="api-notes">
  <p class="api-ctype">GtkLevelBarMode</p>
  <p class="api-since">since 3.6</p>
<table>
<tr>
<td class="name">GTK_LEVEL_BAR_MODE_CONTINUOUS</td>
<td class="value">0</td>
<td class="doc">the bar has a continuous mode</td>
</tr>
<tr>
<td class="name">GTK_LEVEL_BAR_MODE_DISCRETE</td>
<td class="value">1</td>
<td class="doc">the bar has a discrete mode</td>
</tr>
</table>
</div>
<p class="api-heading">License</p>
<p class="api-doc">The type of license for an application.

This enumeration can be expanded at later date.</p>
<div class="api-notes">
  <p class="api-ctype">GtkLicense</p>
  <p class="api-since">since 3.0</p>
<table>
<tr>
<td class="name">GTK_LICENSE_UNKNOWN</td>
<td class="value">0</td>
<td class="doc">No license specified</td>
</tr>
<tr>
<td class="name">GTK_LICENSE_CUSTOM</td>
<td class="value">1</td>
<td class="doc">A license text is going to be specified by the
  developer</td>
</tr>
<tr>
<td class="name">GTK_LICENSE_GPL_2_0</td>
<td class="value">2</td>
<td class="doc">The GNU General Public License, version 2.0 or later</td>
</tr>
<tr>
<td class="name">GTK_LICENSE_GPL_3_0</td>
<td class="value">3</td>
<td class="doc">The GNU General Public License, version 3.0 or later</td>
</tr>
<tr>
<td class="name">GTK_LICENSE_LGPL_2_1</td>
<td class="value">4</td>
<td class="doc">The GNU Lesser General Public License, version 2.1 or later</td>
</tr>
<tr>
<td class="name">GTK_LICENSE_LGPL_3_0</td>
<td class="value">5</td>
<td class="doc">The GNU Lesser General Public License, version 3.0 or later</td>
</tr>
<tr>
<td class="name">GTK_LICENSE_BSD</td>
<td class="value">6</td>
<td class="doc">The BSD standard license</td>
</tr>
<tr>
<td class="name">GTK_LICENSE_MIT_X11</td>
<td class="value">7</td>
<td class="doc">The MIT/X11 standard license</td>
</tr>
<tr>
<td class="name">GTK_LICENSE_ARTISTIC</td>
<td class="value">8</td>
<td class="doc">The Artistic License, version 2.0</td>
</tr>
<tr>
<td class="name">GTK_LICENSE_GPL_2_0_ONLY</td>
<td class="value">9</td>
<td class="doc">The GNU General Public License, version 2.0 only. Since 3.12.</td>
</tr>
<tr>
<td class="name">GTK_LICENSE_GPL_3_0_ONLY</td>
<td class="value">10</td>
<td class="doc">The GNU General Public License, version 3.0 only. Since 3.12.</td>
</tr>
<tr>
<td class="name">GTK_LICENSE_LGPL_2_1_ONLY</td>
<td class="value">11</td>
<td class="doc">The GNU Lesser General Public License, version 2.1 only. Since 3.12.</td>
</tr>
<tr>
<td class="name">GTK_LICENSE_LGPL_3_0_ONLY</td>
<td class="value">12</td>
<td class="doc">The GNU Lesser General Public License, version 3.0 only. Since 3.12.</td>
</tr>
<tr>
<td class="name">GTK_LICENSE_AGPL_3_0</td>
<td class="value">13</td>
<td class="doc">The GNU Affero General Public License, version 3.0 or later. Since: 3.22.</td>
</tr>
<tr>
<td class="name">GTK_LICENSE_AGPL_3_0_ONLY</td>
<td class="value">14</td>
<td class="doc">The GNU Affero General Public License, version 3.0 only. Since: 3.22.27.</td>
</tr>
</table>
</div>
<p class="api-heading">MenuDirectionType</p>
<p class="api-doc">An enumeration representing directional movements within a menu.</p>
<div class="api-notes">
  <p class="api-ctype">GtkMenuDirectionType</p>
<table>
<tr>
<td class="name">GTK_MENU_DIR_PARENT</td>
<td class="value">0</td>
<td class="doc">To the parent menu shell</td>
</tr>
<tr>
<td class="name">GTK_MENU_DIR_CHILD</td>
<td class="value">1</td>
<td class="doc">To the submenu, if any, associated with the item</td>
</tr>
<tr>
<td class="name">GTK_MENU_DIR_NEXT</td>
<td class="value">2</td>
<td class="doc">To the next menu item</td>
</tr>
<tr>
<td class="name">GTK_MENU_DIR_PREV</td>
<td class="value">3</td>
<td class="doc">To the previous menu item</td>
</tr>
</table>
</div>
<p class="api-heading">MessageType</p>
<p class="api-doc">The type of message being displayed in the dialog.</p>
<div class="api-notes">
  <p class="api-ctype">GtkMessageType</p>
<table>
<tr>
<td class="name">GTK_MESSAGE_INFO</td>
<td class="value">0</td>
<td class="doc">Informational message</td>
</tr>
<tr>
<td class="name">GTK_MESSAGE_WARNING</td>
<td class="value">1</td>
<td class="doc">Non-fatal warning message</td>
</tr>
<tr>
<td class="name">GTK_MESSAGE_QUESTION</td>
<td class="value">2</td>
<td class="doc">Question requiring a choice</td>
</tr>
<tr>
<td class="name">GTK_MESSAGE_ERROR</td>
<td class="value">3</td>
<td class="doc">Fatal error message</td>
</tr>
<tr>
<td class="name">GTK_MESSAGE_OTHER</td>
<td class="value">4</td>
<td class="doc">None of the above</td>
</tr>
</table>
</div>
<p class="api-heading">MovementStep</p>
<div class="api-notes">
  <p class="api-ctype">GtkMovementStep</p>
<table>
<tr>
<td class="name">GTK_MOVEMENT_LOGICAL_POSITIONS</td>
<td class="value">0</td>
<td class="doc">Move forward or back by graphemes</td>
</tr>
<tr>
<td class="name">GTK_MOVEMENT_VISUAL_POSITIONS</td>
<td class="value">1</td>
<td class="doc">Move left or right by graphemes</td>
</tr>
<tr>
<td class="name">GTK_MOVEMENT_WORDS</td>
<td class="value">2</td>
<td class="doc">Move forward or back by words</td>
</tr>
<tr>
<td class="name">GTK_MOVEMENT_DISPLAY_LINES</td>
<td class="value">3</td>
<td class="doc">Move up or down lines (wrapped lines)</td>
</tr>
<tr>
<td class="name">GTK_MOVEMENT_DISPLAY_LINE_ENDS</td>
<td class="value">4</td>
<td class="doc">Move to either end of a line</td>
</tr>
<tr>
<td class="name">GTK_MOVEMENT_PARAGRAPHS</td>
<td class="value">5</td>
<td class="doc">Move up or down paragraphs (newline-ended lines)</td>
</tr>
<tr>
<td class="name">GTK_MOVEMENT_PARAGRAPH_ENDS</td>
<td class="value">6</td>
<td class="doc">Move to either end of a paragraph</td>
</tr>
<tr>
<td class="name">GTK_MOVEMENT_PAGES</td>
<td class="value">7</td>
<td class="doc">Move by pages</td>
</tr>
<tr>
<td class="name">GTK_MOVEMENT_BUFFER_ENDS</td>
<td class="value">8</td>
<td class="doc">Move to ends of the buffer</td>
</tr>
<tr>
<td class="name">GTK_MOVEMENT_HORIZONTAL_PAGES</td>
<td class="value">9</td>
<td class="doc">Move horizontally by pages</td>
</tr>
</table>
</div>
<p class="api-heading">NotebookTab</p>
<div class="api-notes">
  <p class="api-ctype">GtkNotebookTab</p>
<table>
<tr>
<td class="name">GTK_NOTEBOOK_TAB_FIRST</td>
<td class="value">0</td>
</tr>
<tr>
<td class="name">GTK_NOTEBOOK_TAB_LAST</td>
<td class="value">1</td>
</tr>
</table>
</div>
<p class="api-heading">NumberUpLayout</p>
<p class="api-doc">Used to determine the layout of pages on a sheet when printing
multiple pages per sheet.</p>
<div class="api-notes">
  <p class="api-ctype">GtkNumberUpLayout</p>
<table>
<tr>
<td class="name">GTK_NUMBER_UP_LAYOUT_LEFT_TO_RIGHT_TOP_TO_BOTTOM</td>
<td class="value">0</td>
<td class="doc">![](layout-lrtb.png)</td>
</tr>
<tr>
<td class="name">GTK_NUMBER_UP_LAYOUT_LEFT_TO_RIGHT_BOTTOM_TO_TOP</td>
<td class="value">1</td>
<td class="doc">![](layout-lrbt.png)</td>
</tr>
<tr>
<td class="name">GTK_NUMBER_UP_LAYOUT_RIGHT_TO_LEFT_TOP_TO_BOTTOM</td>
<td class="value">2</td>
<td class="doc">![](layout-rltb.png)</td>
</tr>
<tr>
<td class="name">GTK_NUMBER_UP_LAYOUT_RIGHT_TO_LEFT_BOTTOM_TO_TOP</td>
<td class="value">3</td>
<td class="doc">![](layout-rlbt.png)</td>
</tr>
<tr>
<td class="name">GTK_NUMBER_UP_LAYOUT_TOP_TO_BOTTOM_LEFT_TO_RIGHT</td>
<td class="value">4</td>
<td class="doc">![](layout-tblr.png)</td>
</tr>
<tr>
<td class="name">GTK_NUMBER_UP_LAYOUT_TOP_TO_BOTTOM_RIGHT_TO_LEFT</td>
<td class="value">5</td>
<td class="doc">![](layout-tbrl.png)</td>
</tr>
<tr>
<td class="name">GTK_NUMBER_UP_LAYOUT_BOTTOM_TO_TOP_LEFT_TO_RIGHT</td>
<td class="value">6</td>
<td class="doc">![](layout-btlr.png)</td>
</tr>
<tr>
<td class="name">GTK_NUMBER_UP_LAYOUT_BOTTOM_TO_TOP_RIGHT_TO_LEFT</td>
<td class="value">7</td>
<td class="doc">![](layout-btrl.png)</td>
</tr>
</table>
</div>
<p class="api-heading">Orientation</p>
<p class="api-doc">Represents the orientation of widgets and other objects which can be switched
between horizontal and vertical orientation on the fly, like #GtkToolbar or
#GtkGesturePan.</p>
<div class="api-notes">
  <p class="api-ctype">GtkOrientation</p>
<table>
<tr>
<td class="name">GTK_ORIENTATION_HORIZONTAL</td>
<td class="value">0</td>
<td class="doc">The element is in horizontal orientation.</td>
</tr>
<tr>
<td class="name">GTK_ORIENTATION_VERTICAL</td>
<td class="value">1</td>
<td class="doc">The element is in vertical orientation.</td>
</tr>
</table>
</div>
<p class="api-heading">PackDirection</p>
<p class="api-doc">Determines how widgets should be packed inside menubars
and menuitems contained in menubars.</p>
<div class="api-notes">
  <p class="api-ctype">GtkPackDirection</p>
<table>
<tr>
<td class="name">GTK_PACK_DIRECTION_LTR</td>
<td class="value">0</td>
<td class="doc">Widgets are packed left-to-right</td>
</tr>
<tr>
<td class="name">GTK_PACK_DIRECTION_RTL</td>
<td class="value">1</td>
<td class="doc">Widgets are packed right-to-left</td>
</tr>
<tr>
<td class="name">GTK_PACK_DIRECTION_TTB</td>
<td class="value">2</td>
<td class="doc">Widgets are packed top-to-bottom</td>
</tr>
<tr>
<td class="name">GTK_PACK_DIRECTION_BTT</td>
<td class="value">3</td>
<td class="doc">Widgets are packed bottom-to-top</td>
</tr>
</table>
</div>
<p class="api-heading">PackType</p>
<p class="api-doc">Represents the packing location #GtkBox children. (See: #GtkVBox,
#GtkHBox, and #GtkButtonBox).</p>
<div class="api-notes">
  <p class="api-ctype">GtkPackType</p>
<table>
<tr>
<td class="name">GTK_PACK_START</td>
<td class="value">0</td>
<td class="doc">The child is packed into the start of the box</td>
</tr>
<tr>
<td class="name">GTK_PACK_END</td>
<td class="value">1</td>
<td class="doc">The child is packed into the end of the box</td>
</tr>
</table>
</div>
<p class="api-heading">PadActionType</p>
<p class="api-doc">The type of a pad action.</p>
<div class="api-notes">
  <p class="api-ctype">GtkPadActionType</p>
  <p class="api-since">since 3.22</p>
<table>
<tr>
<td class="name">GTK_PAD_ACTION_BUTTON</td>
<td class="value">0</td>
<td class="doc">Action is triggered by a pad button</td>
</tr>
<tr>
<td class="name">GTK_PAD_ACTION_RING</td>
<td class="value">1</td>
<td class="doc">Action is triggered by a pad ring</td>
</tr>
<tr>
<td class="name">GTK_PAD_ACTION_STRIP</td>
<td class="value">2</td>
<td class="doc">Action is triggered by a pad strip</td>
</tr>
</table>
</div>
<p class="api-heading">PageOrientation</p>
<p class="api-doc">See also gtk_print_settings_set_orientation().</p>
<div class="api-notes">
  <p class="api-ctype">GtkPageOrientation</p>
<table>
<tr>
<td class="name">GTK_PAGE_ORIENTATION_PORTRAIT</td>
<td class="value">0</td>
<td class="doc">Portrait mode.</td>
</tr>
<tr>
<td class="name">GTK_PAGE_ORIENTATION_LANDSCAPE</td>
<td class="value">1</td>
<td class="doc">Landscape mode.</td>
</tr>
<tr>
<td class="name">GTK_PAGE_ORIENTATION_REVERSE_PORTRAIT</td>
<td class="value">2</td>
<td class="doc">Reverse portrait mode.</td>
</tr>
<tr>
<td class="name">GTK_PAGE_ORIENTATION_REVERSE_LANDSCAPE</td>
<td class="value">3</td>
<td class="doc">Reverse landscape mode.</td>
</tr>
</table>
</div>
<p class="api-heading">PageSet</p>
<p class="api-doc">See also gtk_print_job_set_page_set().</p>
<div class="api-notes">
  <p class="api-ctype">GtkPageSet</p>
<table>
<tr>
<td class="name">GTK_PAGE_SET_ALL</td>
<td class="value">0</td>
<td class="doc">All pages.</td>
</tr>
<tr>
<td class="name">GTK_PAGE_SET_EVEN</td>
<td class="value">1</td>
<td class="doc">Even pages.</td>
</tr>
<tr>
<td class="name">GTK_PAGE_SET_ODD</td>
<td class="value">2</td>
<td class="doc">Odd pages.</td>
</tr>
</table>
</div>
<p class="api-heading">PanDirection</p>
<p class="api-doc">Describes the panning direction of a #GtkGesturePan</p>
<div class="api-notes">
  <p class="api-ctype">GtkPanDirection</p>
  <p class="api-since">since 3.14</p>
<table>
<tr>
<td class="name">GTK_PAN_DIRECTION_LEFT</td>
<td class="value">0</td>
<td class="doc">panned towards the left</td>
</tr>
<tr>
<td class="name">GTK_PAN_DIRECTION_RIGHT</td>
<td class="value">1</td>
<td class="doc">panned towards the right</td>
</tr>
<tr>
<td class="name">GTK_PAN_DIRECTION_UP</td>
<td class="value">2</td>
<td class="doc">panned upwards</td>
</tr>
<tr>
<td class="name">GTK_PAN_DIRECTION_DOWN</td>
<td class="value">3</td>
<td class="doc">panned downwards</td>
</tr>
</table>
</div>
<p class="api-heading">PathPriorityType</p>
<p class="api-doc">Priorities for path lookups.
See also gtk_binding_set_add_path().</p>
<div class="api-notes">
  <p class="api-ctype">GtkPathPriorityType</p>
<table>
<tr>
<td class="name">GTK_PATH_PRIO_LOWEST</td>
<td class="value">0</td>
<td class="doc">Deprecated</td>
</tr>
<tr>
<td class="name">GTK_PATH_PRIO_GTK</td>
<td class="value">4</td>
<td class="doc">Deprecated</td>
</tr>
<tr>
<td class="name">GTK_PATH_PRIO_APPLICATION</td>
<td class="value">8</td>
<td class="doc">Deprecated</td>
</tr>
<tr>
<td class="name">GTK_PATH_PRIO_THEME</td>
<td class="value">10</td>
<td class="doc">Deprecated</td>
</tr>
<tr>
<td class="name">GTK_PATH_PRIO_RC</td>
<td class="value">12</td>
<td class="doc">Deprecated</td>
</tr>
<tr>
<td class="name">GTK_PATH_PRIO_HIGHEST</td>
<td class="value">15</td>
<td class="doc">Deprecated</td>
</tr>
</table>
</div>
<p class="api-heading">PathType</p>
<p class="api-doc">Widget path types.
See also gtk_binding_set_add_path().</p>
<div class="api-notes">
  <p class="api-ctype">GtkPathType</p>
<table>
<tr>
<td class="name">GTK_PATH_WIDGET</td>
<td class="value">0</td>
<td class="doc">Deprecated</td>
</tr>
<tr>
<td class="name">GTK_PATH_WIDGET_CLASS</td>
<td class="value">1</td>
<td class="doc">Deprecated</td>
</tr>
<tr>
<td class="name">GTK_PATH_CLASS</td>
<td class="value">2</td>
<td class="doc">Deprecated</td>
</tr>
</table>
</div>
<p class="api-heading">PolicyType</p>
<p class="api-doc">Determines how the size should be computed to achieve the one of the
visibility mode for the scrollbars.</p>
<div class="api-notes">
  <p class="api-ctype">GtkPolicyType</p>
<table>
<tr>
<td class="name">GTK_POLICY_ALWAYS</td>
<td class="value">0</td>
<td class="doc">The scrollbar is always visible. The view size is
 independent of the content.</td>
</tr>
<tr>
<td class="name">GTK_POLICY_AUTOMATIC</td>
<td class="value">1</td>
<td class="doc">The scrollbar will appear and disappear as necessary.
 For example, when all of a #GtkTreeView can not be seen.</td>
</tr>
<tr>
<td class="name">GTK_POLICY_NEVER</td>
<td class="value">2</td>
<td class="doc">The scrollbar should never appear. In this mode the
 content determines the size.</td>
</tr>
<tr>
<td class="name">GTK_POLICY_EXTERNAL</td>
<td class="value">3</td>
<td class="doc">Don't show a scrollbar, but don't force the
 size to follow the content. This can be used e.g. to make multiple
 scrolled windows share a scrollbar. Since: 3.16</td>
</tr>
</table>
</div>
<p class="api-heading">PopoverConstraint</p>
<p class="api-doc">Describes constraints to positioning of popovers. More values
may be added to this enumeration in the future.</p>
<div class="api-notes">
  <p class="api-ctype">GtkPopoverConstraint</p>
  <p class="api-since">since 3.20</p>
<table>
<tr>
<td class="name">GTK_POPOVER_CONSTRAINT_NONE</td>
<td class="value">0</td>
<td class="doc">Don't constrain the popover position
  beyond what is imposed by the implementation</td>
</tr>
<tr>
<td class="name">GTK_POPOVER_CONSTRAINT_WINDOW</td>
<td class="value">1</td>
<td class="doc">Constrain the popover to the boundaries
  of the window that it is attached to</td>
</tr>
</table>
</div>
<p class="api-heading">PositionType</p>
<p class="api-doc">Describes which edge of a widget a certain feature is positioned at, e.g. the
tabs of a #GtkNotebook, the handle of a #GtkHandleBox or the label of a
#GtkScale.</p>
<div class="api-notes">
  <p class="api-ctype">GtkPositionType</p>
<table>
<tr>
<td class="name">GTK_POS_LEFT</td>
<td class="value">0</td>
<td class="doc">The feature is at the left edge.</td>
</tr>
<tr>
<td class="name">GTK_POS_RIGHT</td>
<td class="value">1</td>
<td class="doc">The feature is at the right edge.</td>
</tr>
<tr>
<td class="name">GTK_POS_TOP</td>
<td class="value">2</td>
<td class="doc">The feature is at the top edge.</td>
</tr>
<tr>
<td class="name">GTK_POS_BOTTOM</td>
<td class="value">3</td>
<td class="doc">The feature is at the bottom edge.</td>
</tr>
</table>
</div>
<p class="api-heading">PrintDuplex</p>
<p class="api-doc">See also gtk_print_settings_set_duplex().</p>
<div class="api-notes">
  <p class="api-ctype">GtkPrintDuplex</p>
<table>
<tr>
<td class="name">GTK_PRINT_DUPLEX_SIMPLEX</td>
<td class="value">0</td>
<td class="doc">No duplex.</td>
</tr>
<tr>
<td class="name">GTK_PRINT_DUPLEX_HORIZONTAL</td>
<td class="value">1</td>
<td class="doc">Horizontal duplex.</td>
</tr>
<tr>
<td class="name">GTK_PRINT_DUPLEX_VERTICAL</td>
<td class="value">2</td>
<td class="doc">Vertical duplex.</td>
</tr>
</table>
</div>
<p class="api-heading">PrintError</p>
<p class="api-doc">Error codes that identify various errors that can occur while
using the GTK+ printing support.</p>
<div class="api-notes">
  <p class="api-ctype">GtkPrintError</p>
<table>
<tr>
<td class="name">GTK_PRINT_ERROR_GENERAL</td>
<td class="value">0</td>
<td class="doc">An unspecified error occurred.</td>
</tr>
<tr>
<td class="name">GTK_PRINT_ERROR_INTERNAL_ERROR</td>
<td class="value">1</td>
<td class="doc">An internal error occurred.</td>
</tr>
<tr>
<td class="name">GTK_PRINT_ERROR_NOMEM</td>
<td class="value">2</td>
<td class="doc">A memory allocation failed.</td>
</tr>
<tr>
<td class="name">GTK_PRINT_ERROR_INVALID_FILE</td>
<td class="value">3</td>
<td class="doc">An error occurred while loading a page setup
    or paper size from a key file.</td>
</tr>
</table>
</div>
<p class="api-heading">PrintOperationAction</p>
<p class="api-doc">The @action parameter to gtk_print_operation_run()
determines what action the print operation should perform.</p>
<div class="api-notes">
  <p class="api-ctype">GtkPrintOperationAction</p>
<table>
<tr>
<td class="name">GTK_PRINT_OPERATION_ACTION_PRINT_DIALOG</td>
<td class="value">0</td>
<td class="doc">Show the print dialog.</td>
</tr>
<tr>
<td class="name">GTK_PRINT_OPERATION_ACTION_PRINT</td>
<td class="value">1</td>
<td class="doc">Start to print without showing
    the print dialog, based on the current print settings.</td>
</tr>
<tr>
<td class="name">GTK_PRINT_OPERATION_ACTION_PREVIEW</td>
<td class="value">2</td>
<td class="doc">Show the print preview.</td>
</tr>
<tr>
<td class="name">GTK_PRINT_OPERATION_ACTION_EXPORT</td>
<td class="value">3</td>
<td class="doc">Export to a file. This requires
    the export-filename property to be set.</td>
</tr>
</table>
</div>
<p class="api-heading">PrintOperationResult</p>
<p class="api-doc">A value of this type is returned by gtk_print_operation_run().</p>
<div class="api-notes">
  <p class="api-ctype">GtkPrintOperationResult</p>
<table>
<tr>
<td class="name">GTK_PRINT_OPERATION_RESULT_ERROR</td>
<td class="value">0</td>
<td class="doc">An error has occurred.</td>
</tr>
<tr>
<td class="name">GTK_PRINT_OPERATION_RESULT_APPLY</td>
<td class="value">1</td>
<td class="doc">The print settings should be stored.</td>
</tr>
<tr>
<td class="name">GTK_PRINT_OPERATION_RESULT_CANCEL</td>
<td class="value">2</td>
<td class="doc">The print operation has been canceled,
    the print settings should not be stored.</td>
</tr>
<tr>
<td class="name">GTK_PRINT_OPERATION_RESULT_IN_PROGRESS</td>
<td class="value">3</td>
<td class="doc">The print operation is not complete
    yet. This value will only be returned when running asynchronously.</td>
</tr>
</table>
</div>
<p class="api-heading">PrintPages</p>
<p class="api-doc">See also gtk_print_job_set_pages()</p>
<div class="api-notes">
  <p class="api-ctype">GtkPrintPages</p>
<table>
<tr>
<td class="name">GTK_PRINT_PAGES_ALL</td>
<td class="value">0</td>
<td class="doc">All pages.</td>
</tr>
<tr>
<td class="name">GTK_PRINT_PAGES_CURRENT</td>
<td class="value">1</td>
<td class="doc">Current page.</td>
</tr>
<tr>
<td class="name">GTK_PRINT_PAGES_RANGES</td>
<td class="value">2</td>
<td class="doc">Range of pages.</td>
</tr>
<tr>
<td class="name">GTK_PRINT_PAGES_SELECTION</td>
<td class="value">3</td>
<td class="doc">Selected pages.</td>
</tr>
</table>
</div>
<p class="api-heading">PrintQuality</p>
<p class="api-doc">See also gtk_print_settings_set_quality().</p>
<div class="api-notes">
  <p class="api-ctype">GtkPrintQuality</p>
<table>
<tr>
<td class="name">GTK_PRINT_QUALITY_LOW</td>
<td class="value">0</td>
<td class="doc">Low quality.</td>
</tr>
<tr>
<td class="name">GTK_PRINT_QUALITY_NORMAL</td>
<td class="value">1</td>
<td class="doc">Normal quality.</td>
</tr>
<tr>
<td class="name">GTK_PRINT_QUALITY_HIGH</td>
<td class="value">2</td>
<td class="doc">High quality.</td>
</tr>
<tr>
<td class="name">GTK_PRINT_QUALITY_DRAFT</td>
<td class="value">3</td>
<td class="doc">Draft quality.</td>
</tr>
</table>
</div>
<p class="api-heading">PrintStatus</p>
<p class="api-doc">The status gives a rough indication of the completion of a running
print operation.</p>
<div class="api-notes">
  <p class="api-ctype">GtkPrintStatus</p>
<table>
<tr>
<td class="name">GTK_PRINT_STATUS_INITIAL</td>
<td class="value">0</td>
<td class="doc">The printing has not started yet; this
    status is set initially, and while the print dialog is shown.</td>
</tr>
<tr>
<td class="name">GTK_PRINT_STATUS_PREPARING</td>
<td class="value">1</td>
<td class="doc">This status is set while the begin-print
    signal is emitted and during pagination.</td>
</tr>
<tr>
<td class="name">GTK_PRINT_STATUS_GENERATING_DATA</td>
<td class="value">2</td>
<td class="doc">This status is set while the
    pages are being rendered.</td>
</tr>
<tr>
<td class="name">GTK_PRINT_STATUS_SENDING_DATA</td>
<td class="value">3</td>
<td class="doc">The print job is being sent off to the
    printer.</td>
</tr>
<tr>
<td class="name">GTK_PRINT_STATUS_PENDING</td>
<td class="value">4</td>
<td class="doc">The print job has been sent to the printer,
    but is not printed for some reason, e.g. the printer may be stopped.</td>
</tr>
<tr>
<td class="name">GTK_PRINT_STATUS_PENDING_ISSUE</td>
<td class="value">5</td>
<td class="doc">Some problem has occurred during
    printing, e.g. a paper jam.</td>
</tr>
<tr>
<td class="name">GTK_PRINT_STATUS_PRINTING</td>
<td class="value">6</td>
<td class="doc">The printer is processing the print job.</td>
</tr>
<tr>
<td class="name">GTK_PRINT_STATUS_FINISHED</td>
<td class="value">7</td>
<td class="doc">The printing has been completed successfully.</td>
</tr>
<tr>
<td class="name">GTK_PRINT_STATUS_FINISHED_ABORTED</td>
<td class="value">8</td>
<td class="doc">The printing has been aborted.</td>
</tr>
</table>
</div>
<p class="api-heading">PropagationPhase</p>
<p class="api-doc">Describes the stage at which events are fed into a #GtkEventController.</p>
<div class="api-notes">
  <p class="api-ctype">GtkPropagationPhase</p>
  <p class="api-since">since 3.14</p>
<table>
<tr>
<td class="name">GTK_PHASE_NONE</td>
<td class="value">0</td>
<td class="doc">Events are not delivered automatically. Those can be
  manually fed through gtk_event_controller_handle_event(). This should
  only be used when full control about when, or whether the controller
  handles the event is needed.</td>
</tr>
<tr>
<td class="name">GTK_PHASE_CAPTURE</td>
<td class="value">1</td>
<td class="doc">Events are delivered in the capture phase. The
  capture phase happens before the bubble phase, runs from the toplevel down
  to the event widget. This option should only be used on containers that
  might possibly handle events before their children do.</td>
</tr>
<tr>
<td class="name">GTK_PHASE_BUBBLE</td>
<td class="value">2</td>
<td class="doc">Events are delivered in the bubble phase. The bubble
  phase happens after the capture phase, and before the default handlers
  are run. This phase runs from the event widget, up to the toplevel.</td>
</tr>
<tr>
<td class="name">GTK_PHASE_TARGET</td>
<td class="value">3</td>
<td class="doc">Events are delivered in the default widget event handlers,
  note that widget implementations must chain up on button, motion, touch and
  grab broken handlers for controllers in this phase to be run.</td>
</tr>
</table>
</div>
<p class="api-heading">RcTokenType</p>
<p class="api-doc">The #GtkRcTokenType enumeration represents the tokens
in the RC file. It is exposed so that theme engines
can reuse these tokens when parsing the theme-engine
specific portions of a RC file.</p>
<div class="api-notes">
  <p class="api-ctype">GtkRcTokenType</p>
<table>
<tr>
<td class="name">GTK_RC_TOKEN_INVALID</td>
<td class="value">270</td>
<td class="doc">Deprecated</td>
</tr>
<tr>
<td class="name">GTK_RC_TOKEN_INCLUDE</td>
<td class="value">271</td>
<td class="doc">Deprecated</td>
</tr>
<tr>
<td class="name">GTK_RC_TOKEN_NORMAL</td>
<td class="value">272</td>
<td class="doc">Deprecated</td>
</tr>
<tr>
<td class="name">GTK_RC_TOKEN_ACTIVE</td>
<td class="value">273</td>
<td class="doc">Deprecated</td>
</tr>
<tr>
<td class="name">GTK_RC_TOKEN_PRELIGHT</td>
<td class="value">274</td>
<td class="doc">Deprecated</td>
</tr>
<tr>
<td class="name">GTK_RC_TOKEN_SELECTED</td>
<td class="value">275</td>
<td class="doc">Deprecated</td>
</tr>
<tr>
<td class="name">GTK_RC_TOKEN_INSENSITIVE</td>
<td class="value">276</td>
<td class="doc">Deprecated</td>
</tr>
<tr>
<td class="name">GTK_RC_TOKEN_FG</td>
<td class="value">277</td>
<td class="doc">Deprecated</td>
</tr>
<tr>
<td class="name">GTK_RC_TOKEN_BG</td>
<td class="value">278</td>
<td class="doc">Deprecated</td>
</tr>
<tr>
<td class="name">GTK_RC_TOKEN_TEXT</td>
<td class="value">279</td>
<td class="doc">Deprecated</td>
</tr>
<tr>
<td class="name">GTK_RC_TOKEN_BASE</td>
<td class="value">280</td>
<td class="doc">Deprecated</td>
</tr>
<tr>
<td class="name">GTK_RC_TOKEN_XTHICKNESS</td>
<td class="value">281</td>
<td class="doc">Deprecated</td>
</tr>
<tr>
<td class="name">GTK_RC_TOKEN_YTHICKNESS</td>
<td class="value">282</td>
<td class="doc">Deprecated</td>
</tr>
<tr>
<td class="name">GTK_RC_TOKEN_FONT</td>
<td class="value">283</td>
<td class="doc">Deprecated</td>
</tr>
<tr>
<td class="name">GTK_RC_TOKEN_FONTSET</td>
<td class="value">284</td>
<td class="doc">Deprecated</td>
</tr>
<tr>
<td class="name">GTK_RC_TOKEN_FONT_NAME</td>
<td class="value">285</td>
<td class="doc">Deprecated</td>
</tr>
<tr>
<td class="name">GTK_RC_TOKEN_BG_PIXMAP</td>
<td class="value">286</td>
<td class="doc">Deprecated</td>
</tr>
<tr>
<td class="name">GTK_RC_TOKEN_PIXMAP_PATH</td>
<td class="value">287</td>
<td class="doc">Deprecated</td>
</tr>
<tr>
<td class="name">GTK_RC_TOKEN_STYLE</td>
<td class="value">288</td>
<td class="doc">Deprecated</td>
</tr>
<tr>
<td class="name">GTK_RC_TOKEN_BINDING</td>
<td class="value">289</td>
<td class="doc">Deprecated</td>
</tr>
<tr>
<td class="name">GTK_RC_TOKEN_BIND</td>
<td class="value">290</td>
<td class="doc">Deprecated</td>
</tr>
<tr>
<td class="name">GTK_RC_TOKEN_WIDGET</td>
<td class="value">291</td>
<td class="doc">Deprecated</td>
</tr>
<tr>
<td class="name">GTK_RC_TOKEN_WIDGET_CLASS</td>
<td class="value">292</td>
<td class="doc">Deprecated</td>
</tr>
<tr>
<td class="name">GTK_RC_TOKEN_CLASS</td>
<td class="value">293</td>
<td class="doc">Deprecated</td>
</tr>
<tr>
<td class="name">GTK_RC_TOKEN_LOWEST</td>
<td class="value">294</td>
<td class="doc">Deprecated</td>
</tr>
<tr>
<td class="name">GTK_RC_TOKEN_GTK</td>
<td class="value">295</td>
<td class="doc">Deprecated</td>
</tr>
<tr>
<td class="name">GTK_RC_TOKEN_APPLICATION</td>
<td class="value">296</td>
<td class="doc">Deprecated</td>
</tr>
<tr>
<td class="name">GTK_RC_TOKEN_THEME</td>
<td class="value">297</td>
<td class="doc">Deprecated</td>
</tr>
<tr>
<td class="name">GTK_RC_TOKEN_RC</td>
<td class="value">298</td>
<td class="doc">Deprecated</td>
</tr>
<tr>
<td class="name">GTK_RC_TOKEN_HIGHEST</td>
<td class="value">299</td>
<td class="doc">Deprecated</td>
</tr>
<tr>
<td class="name">GTK_RC_TOKEN_ENGINE</td>
<td class="value">300</td>
<td class="doc">Deprecated</td>
</tr>
<tr>
<td class="name">GTK_RC_TOKEN_MODULE_PATH</td>
<td class="value">301</td>
<td class="doc">Deprecated</td>
</tr>
<tr>
<td class="name">GTK_RC_TOKEN_IM_MODULE_PATH</td>
<td class="value">302</td>
<td class="doc">Deprecated</td>
</tr>
<tr>
<td class="name">GTK_RC_TOKEN_IM_MODULE_FILE</td>
<td class="value">303</td>
<td class="doc">Deprecated</td>
</tr>
<tr>
<td class="name">GTK_RC_TOKEN_STOCK</td>
<td class="value">304</td>
<td class="doc">Deprecated</td>
</tr>
<tr>
<td class="name">GTK_RC_TOKEN_LTR</td>
<td class="value">305</td>
<td class="doc">Deprecated</td>
</tr>
<tr>
<td class="name">GTK_RC_TOKEN_RTL</td>
<td class="value">306</td>
<td class="doc">Deprecated</td>
</tr>
<tr>
<td class="name">GTK_RC_TOKEN_COLOR</td>
<td class="value">307</td>
<td class="doc">Deprecated</td>
</tr>
<tr>
<td class="name">GTK_RC_TOKEN_UNBIND</td>
<td class="value">308</td>
<td class="doc">Deprecated</td>
</tr>
<tr>
<td class="name">GTK_RC_TOKEN_LAST</td>
<td class="value">309</td>
<td class="doc">Deprecated</td>
</tr>
</table>
</div>
<p class="api-heading">RecentChooserError</p>
<p class="api-doc">These identify the various errors that can occur while calling
#GtkRecentChooser functions.</p>
<div class="api-notes">
  <p class="api-ctype">GtkRecentChooserError</p>
  <p class="api-since">since 2.10</p>
<table>
<tr>
<td class="name">GTK_RECENT_CHOOSER_ERROR_NOT_FOUND</td>
<td class="value">0</td>
<td class="doc">Indicates that a file does not exist</td>
</tr>
<tr>
<td class="name">GTK_RECENT_CHOOSER_ERROR_INVALID_URI</td>
<td class="value">1</td>
<td class="doc">Indicates a malformed URI</td>
</tr>
</table>
</div>
<p class="api-heading">RecentManagerError</p>
<p class="api-doc">Error codes for #GtkRecentManager operations</p>
<div class="api-notes">
  <p class="api-ctype">GtkRecentManagerError</p>
  <p class="api-since">since 2.10</p>
<table>
<tr>
<td class="name">GTK_RECENT_MANAGER_ERROR_NOT_FOUND</td>
<td class="value">0</td>
<td class="doc">the URI specified does not exists in
  the recently used resources list.</td>
</tr>
<tr>
<td class="name">GTK_RECENT_MANAGER_ERROR_INVALID_URI</td>
<td class="value">1</td>
<td class="doc">the URI specified is not valid.</td>
</tr>
<tr>
<td class="name">GTK_RECENT_MANAGER_ERROR_INVALID_ENCODING</td>
<td class="value">2</td>
<td class="doc">the supplied string is not
  UTF-8 encoded.</td>
</tr>
<tr>
<td class="name">GTK_RECENT_MANAGER_ERROR_NOT_REGISTERED</td>
<td class="value">3</td>
<td class="doc">no application has registered
  the specified item.</td>
</tr>
<tr>
<td class="name">GTK_RECENT_MANAGER_ERROR_READ</td>
<td class="value">4</td>
<td class="doc">failure while reading the recently used
  resources file.</td>
</tr>
<tr>
<td class="name">GTK_RECENT_MANAGER_ERROR_WRITE</td>
<td class="value">5</td>
<td class="doc">failure while writing the recently used
  resources file.</td>
</tr>
<tr>
<td class="name">GTK_RECENT_MANAGER_ERROR_UNKNOWN</td>
<td class="value">6</td>
<td class="doc">unspecified error.</td>
</tr>
</table>
</div>
<p class="api-heading">RecentSortType</p>
<p class="api-doc">Used to specify the sorting method to be applyed to the recently
used resource list.</p>
<div class="api-notes">
  <p class="api-ctype">GtkRecentSortType</p>
  <p class="api-since">since 2.10</p>
<table>
<tr>
<td class="name">GTK_RECENT_SORT_NONE</td>
<td class="value">0</td>
<td class="doc">Do not sort the returned list of recently used
  resources.</td>
</tr>
<tr>
<td class="name">GTK_RECENT_SORT_MRU</td>
<td class="value">1</td>
<td class="doc">Sort the returned list with the most recently used
  items first.</td>
</tr>
<tr>
<td class="name">GTK_RECENT_SORT_LRU</td>
<td class="value">2</td>
<td class="doc">Sort the returned list with the least recently used
  items first.</td>
</tr>
<tr>
<td class="name">GTK_RECENT_SORT_CUSTOM</td>
<td class="value">3</td>
<td class="doc">Sort the returned list using a custom sorting
  function passed using gtk_recent_chooser_set_sort_func().</td>
</tr>
</table>
</div>
<p class="api-heading">ReliefStyle</p>
<p class="api-doc">Indicated the relief to be drawn around a #GtkButton.</p>
<div class="api-notes">
  <p class="api-ctype">GtkReliefStyle</p>
<table>
<tr>
<td class="name">GTK_RELIEF_NORMAL</td>
<td class="value">0</td>
<td class="doc">Draw a normal relief.</td>
</tr>
<tr>
<td class="name">GTK_RELIEF_HALF</td>
<td class="value">1</td>
<td class="doc">A half relief. Deprecated in 3.14, does the same as @GTK_RELIEF_NORMAL</td>
</tr>
<tr>
<td class="name">GTK_RELIEF_NONE</td>
<td class="value">2</td>
<td class="doc">No relief.</td>
</tr>
</table>
</div>
<p class="api-heading">ResizeMode</p>
<div class="api-notes">
  <p class="api-ctype">GtkResizeMode</p>
<table>
<tr>
<td class="name">GTK_RESIZE_PARENT</td>
<td class="value">0</td>
<td class="doc">Pass resize request to the parent</td>
</tr>
<tr>
<td class="name">GTK_RESIZE_QUEUE</td>
<td class="value">1</td>
<td class="doc">Queue resizes on this widget</td>
</tr>
<tr>
<td class="name">GTK_RESIZE_IMMEDIATE</td>
<td class="value">2</td>
<td class="doc">Resize immediately. Deprecated.</td>
</tr>
</table>
</div>
<p class="api-heading">ResponseType</p>
<p class="api-doc">Predefined values for use as response ids in gtk_dialog_add_button().
All predefined values are negative; GTK+ leaves values of 0 or greater for
application-defined response ids.</p>
<div class="api-notes">
  <p class="api-ctype">GtkResponseType</p>
<table>
<tr>
<td class="name">GTK_RESPONSE_NONE</td>
<td class="value">-1</td>
<td class="doc">Returned if an action widget has no response id,
    or if the dialog gets programmatically hidden or destroyed</td>
</tr>
<tr>
<td class="name">GTK_RESPONSE_REJECT</td>
<td class="value">-2</td>
<td class="doc">Generic response id, not used by GTK+ dialogs</td>
</tr>
<tr>
<td class="name">GTK_RESPONSE_ACCEPT</td>
<td class="value">-3</td>
<td class="doc">Generic response id, not used by GTK+ dialogs</td>
</tr>
<tr>
<td class="name">GTK_RESPONSE_DELETE_EVENT</td>
<td class="value">-4</td>
<td class="doc">Returned if the dialog is deleted</td>
</tr>
<tr>
<td class="name">GTK_RESPONSE_OK</td>
<td class="value">-5</td>
<td class="doc">Returned by OK buttons in GTK+ dialogs</td>
</tr>
<tr>
<td class="name">GTK_RESPONSE_CANCEL</td>
<td class="value">-6</td>
<td class="doc">Returned by Cancel buttons in GTK+ dialogs</td>
</tr>
<tr>
<td class="name">GTK_RESPONSE_CLOSE</td>
<td class="value">-7</td>
<td class="doc">Returned by Close buttons in GTK+ dialogs</td>
</tr>
<tr>
<td class="name">GTK_RESPONSE_YES</td>
<td class="value">-8</td>
<td class="doc">Returned by Yes buttons in GTK+ dialogs</td>
</tr>
<tr>
<td class="name">GTK_RESPONSE_NO</td>
<td class="value">-9</td>
<td class="doc">Returned by No buttons in GTK+ dialogs</td>
</tr>
<tr>
<td class="name">GTK_RESPONSE_APPLY</td>
<td class="value">-10</td>
<td class="doc">Returned by Apply buttons in GTK+ dialogs</td>
</tr>
<tr>
<td class="name">GTK_RESPONSE_HELP</td>
<td class="value">-11</td>
<td class="doc">Returned by Help buttons in GTK+ dialogs</td>
</tr>
</table>
</div>
<p class="api-heading">RevealerTransitionType</p>
<p class="api-doc">These enumeration values describe the possible transitions
when the child of a #GtkRevealer widget is shown or hidden.</p>
<div class="api-notes">
  <p class="api-ctype">GtkRevealerTransitionType</p>
<table>
<tr>
<td class="name">GTK_REVEALER_TRANSITION_TYPE_NONE</td>
<td class="value">0</td>
<td class="doc">No transition</td>
</tr>
<tr>
<td class="name">GTK_REVEALER_TRANSITION_TYPE_CROSSFADE</td>
<td class="value">1</td>
<td class="doc">Fade in</td>
</tr>
<tr>
<td class="name">GTK_REVEALER_TRANSITION_TYPE_SLIDE_RIGHT</td>
<td class="value">2</td>
<td class="doc">Slide in from the left</td>
</tr>
<tr>
<td class="name">GTK_REVEALER_TRANSITION_TYPE_SLIDE_LEFT</td>
<td class="value">3</td>
<td class="doc">Slide in from the right</td>
</tr>
<tr>
<td class="name">GTK_REVEALER_TRANSITION_TYPE_SLIDE_UP</td>
<td class="value">4</td>
<td class="doc">Slide in from the bottom</td>
</tr>
<tr>
<td class="name">GTK_REVEALER_TRANSITION_TYPE_SLIDE_DOWN</td>
<td class="value">5</td>
<td class="doc">Slide in from the top</td>
</tr>
</table>
</div>
<p class="api-heading">ScrollStep</p>
<div class="api-notes">
  <p class="api-ctype">GtkScrollStep</p>
<table>
<tr>
<td class="name">GTK_SCROLL_STEPS</td>
<td class="value">0</td>
<td class="doc">Scroll in steps.</td>
</tr>
<tr>
<td class="name">GTK_SCROLL_PAGES</td>
<td class="value">1</td>
<td class="doc">Scroll by pages.</td>
</tr>
<tr>
<td class="name">GTK_SCROLL_ENDS</td>
<td class="value">2</td>
<td class="doc">Scroll to ends.</td>
</tr>
<tr>
<td class="name">GTK_SCROLL_HORIZONTAL_STEPS</td>
<td class="value">3</td>
<td class="doc">Scroll in horizontal steps.</td>
</tr>
<tr>
<td class="name">GTK_SCROLL_HORIZONTAL_PAGES</td>
<td class="value">4</td>
<td class="doc">Scroll by horizontal pages.</td>
</tr>
<tr>
<td class="name">GTK_SCROLL_HORIZONTAL_ENDS</td>
<td class="value">5</td>
<td class="doc">Scroll to the horizontal ends.</td>
</tr>
</table>
</div>
<p class="api-heading">ScrollType</p>
<p class="api-doc">Scrolling types.</p>
<div class="api-notes">
  <p class="api-ctype">GtkScrollType</p>
<table>
<tr>
<td class="name">GTK_SCROLL_NONE</td>
<td class="value">0</td>
<td class="doc">No scrolling.</td>
</tr>
<tr>
<td class="name">GTK_SCROLL_JUMP</td>
<td class="value">1</td>
<td class="doc">Jump to new location.</td>
</tr>
<tr>
<td class="name">GTK_SCROLL_STEP_BACKWARD</td>
<td class="value">2</td>
<td class="doc">Step backward.</td>
</tr>
<tr>
<td class="name">GTK_SCROLL_STEP_FORWARD</td>
<td class="value">3</td>
<td class="doc">Step forward.</td>
</tr>
<tr>
<td class="name">GTK_SCROLL_PAGE_BACKWARD</td>
<td class="value">4</td>
<td class="doc">Page backward.</td>
</tr>
<tr>
<td class="name">GTK_SCROLL_PAGE_FORWARD</td>
<td class="value">5</td>
<td class="doc">Page forward.</td>
</tr>
<tr>
<td class="name">GTK_SCROLL_STEP_UP</td>
<td class="value">6</td>
<td class="doc">Step up.</td>
</tr>
<tr>
<td class="name">GTK_SCROLL_STEP_DOWN</td>
<td class="value">7</td>
<td class="doc">Step down.</td>
</tr>
<tr>
<td class="name">GTK_SCROLL_PAGE_UP</td>
<td class="value">8</td>
<td class="doc">Page up.</td>
</tr>
<tr>
<td class="name">GTK_SCROLL_PAGE_DOWN</td>
<td class="value">9</td>
<td class="doc">Page down.</td>
</tr>
<tr>
<td class="name">GTK_SCROLL_STEP_LEFT</td>
<td class="value">10</td>
<td class="doc">Step to the left.</td>
</tr>
<tr>
<td class="name">GTK_SCROLL_STEP_RIGHT</td>
<td class="value">11</td>
<td class="doc">Step to the right.</td>
</tr>
<tr>
<td class="name">GTK_SCROLL_PAGE_LEFT</td>
<td class="value">12</td>
<td class="doc">Page to the left.</td>
</tr>
<tr>
<td class="name">GTK_SCROLL_PAGE_RIGHT</td>
<td class="value">13</td>
<td class="doc">Page to the right.</td>
</tr>
<tr>
<td class="name">GTK_SCROLL_START</td>
<td class="value">14</td>
<td class="doc">Scroll to start.</td>
</tr>
<tr>
<td class="name">GTK_SCROLL_END</td>
<td class="value">15</td>
<td class="doc">Scroll to end.</td>
</tr>
</table>
</div>
<p class="api-heading">ScrollablePolicy</p>
<p class="api-doc">Defines the policy to be used in a scrollable widget when updating
the scrolled window adjustments in a given orientation.</p>
<div class="api-notes">
  <p class="api-ctype">GtkScrollablePolicy</p>
<table>
<tr>
<td class="name">GTK_SCROLL_MINIMUM</td>
<td class="value">0</td>
<td class="doc">Scrollable adjustments are based on the minimum size</td>
</tr>
<tr>
<td class="name">GTK_SCROLL_NATURAL</td>
<td class="value">1</td>
<td class="doc">Scrollable adjustments are based on the natural size</td>
</tr>
</table>
</div>
<p class="api-heading">SelectionMode</p>
<p class="api-doc">Used to control what selections users are allowed to make.</p>
<div class="api-notes">
  <p class="api-ctype">GtkSelectionMode</p>
<table>
<tr>
<td class="name">GTK_SELECTION_NONE</td>
<td class="value">0</td>
<td class="doc">No selection is possible.</td>
</tr>
<tr>
<td class="name">GTK_SELECTION_SINGLE</td>
<td class="value">1</td>
<td class="doc">Zero or one element may be selected.</td>
</tr>
<tr>
<td class="name">GTK_SELECTION_BROWSE</td>
<td class="value">2</td>
<td class="doc">Exactly one element is selected.
    In some circumstances, such as initially or during a search
    operation, it’s possible for no element to be selected with
    %GTK_SELECTION_BROWSE. What is really enforced is that the user
    can’t deselect a currently selected element except by selecting
    another element.</td>
</tr>
<tr>
<td class="name">GTK_SELECTION_MULTIPLE</td>
<td class="value">3</td>
<td class="doc">Any number of elements may be selected.
     The Ctrl key may be used to enlarge the selection, and Shift
     key to select between the focus and the child pointed to.
     Some widgets may also allow Click-drag to select a range of elements.</td>
</tr>
</table>
</div>
<p class="api-heading">SensitivityType</p>
<p class="api-doc">Determines how GTK+ handles the sensitivity of stepper arrows
at the end of range widgets.</p>
<div class="api-notes">
  <p class="api-ctype">GtkSensitivityType</p>
<table>
<tr>
<td class="name">GTK_SENSITIVITY_AUTO</td>
<td class="value">0</td>
<td class="doc">The arrow is made insensitive if the
  thumb is at the end</td>
</tr>
<tr>
<td class="name">GTK_SENSITIVITY_ON</td>
<td class="value">1</td>
<td class="doc">The arrow is always sensitive</td>
</tr>
<tr>
<td class="name">GTK_SENSITIVITY_OFF</td>
<td class="value">2</td>
<td class="doc">The arrow is always insensitive</td>
</tr>
</table>
</div>
<p class="api-heading">ShadowType</p>
<p class="api-doc">Used to change the appearance of an outline typically provided by a #GtkFrame.

Note that many themes do not differentiate the appearance of the
various shadow types: Either their is no visible shadow (@GTK_SHADOW_NONE),
or there is (any other value).</p>
<div class="api-notes">
  <p class="api-ctype">GtkShadowType</p>
<table>
<tr>
<td class="name">GTK_SHADOW_NONE</td>
<td class="value">0</td>
<td class="doc">No outline.</td>
</tr>
<tr>
<td class="name">GTK_SHADOW_IN</td>
<td class="value">1</td>
<td class="doc">The outline is bevelled inwards.</td>
</tr>
<tr>
<td class="name">GTK_SHADOW_OUT</td>
<td class="value">2</td>
<td class="doc">The outline is bevelled outwards like a button.</td>
</tr>
<tr>
<td class="name">GTK_SHADOW_ETCHED_IN</td>
<td class="value">3</td>
<td class="doc">The outline has a sunken 3d appearance.</td>
</tr>
<tr>
<td class="name">GTK_SHADOW_ETCHED_OUT</td>
<td class="value">4</td>
<td class="doc">The outline has a raised 3d appearance.</td>
</tr>
</table>
</div>
<p class="api-heading">ShortcutType</p>
<p class="api-doc">GtkShortcutType specifies the kind of shortcut that is being described.
More values may be added to this enumeration over time.</p>
<div class="api-notes">
  <p class="api-ctype">GtkShortcutType</p>
  <p class="api-since">since 3.20</p>
<table>
<tr>
<td class="name">GTK_SHORTCUT_ACCELERATOR</td>
<td class="value">0</td>
<td class="doc">The shortcut is a keyboard accelerator. The #GtkShortcutsShortcut:accelerator
  property will be used.</td>
</tr>
<tr>
<td class="name">GTK_SHORTCUT_GESTURE_PINCH</td>
<td class="value">1</td>
<td class="doc">The shortcut is a pinch gesture. GTK+ provides an icon and subtitle.</td>
</tr>
<tr>
<td class="name">GTK_SHORTCUT_GESTURE_STRETCH</td>
<td class="value">2</td>
<td class="doc">The shortcut is a stretch gesture. GTK+ provides an icon and subtitle.</td>
</tr>
<tr>
<td class="name">GTK_SHORTCUT_GESTURE_ROTATE_CLOCKWISE</td>
<td class="value">3</td>
<td class="doc">The shortcut is a clockwise rotation gesture. GTK+ provides an icon and subtitle.</td>
</tr>
<tr>
<td class="name">GTK_SHORTCUT_GESTURE_ROTATE_COUNTERCLOCKWISE</td>
<td class="value">4</td>
<td class="doc">The shortcut is a counterclockwise rotation gesture. GTK+ provides an icon and subtitle.</td>
</tr>
<tr>
<td class="name">GTK_SHORTCUT_GESTURE_TWO_FINGER_SWIPE_LEFT</td>
<td class="value">5</td>
<td class="doc">The shortcut is a two-finger swipe gesture. GTK+ provides an icon and subtitle.</td>
</tr>
<tr>
<td class="name">GTK_SHORTCUT_GESTURE_TWO_FINGER_SWIPE_RIGHT</td>
<td class="value">6</td>
<td class="doc">The shortcut is a two-finger swipe gesture. GTK+ provides an icon and subtitle.</td>
</tr>
<tr>
<td class="name">GTK_SHORTCUT_GESTURE</td>
<td class="value">7</td>
<td class="doc">The shortcut is a gesture. The #GtkShortcutsShortcut:icon property will be
  used.</td>
</tr>
</table>
</div>
<p class="api-heading">SizeGroupMode</p>
<p class="api-doc">The mode of the size group determines the directions in which the size
group affects the requested sizes of its component widgets.</p>
<div class="api-notes">
  <p class="api-ctype">GtkSizeGroupMode</p>
<table>
<tr>
<td class="name">GTK_SIZE_GROUP_NONE</td>
<td class="value">0</td>
<td class="doc">group has no effect</td>
</tr>
<tr>
<td class="name">GTK_SIZE_GROUP_HORIZONTAL</td>
<td class="value">1</td>
<td class="doc">group affects horizontal requisition</td>
</tr>
<tr>
<td class="name">GTK_SIZE_GROUP_VERTICAL</td>
<td class="value">2</td>
<td class="doc">group affects vertical requisition</td>
</tr>
<tr>
<td class="name">GTK_SIZE_GROUP_BOTH</td>
<td class="value">3</td>
<td class="doc">group affects both horizontal and vertical requisition</td>
</tr>
</table>
</div>
<p class="api-heading">SizeRequestMode</p>
<p class="api-doc">Specifies a preference for height-for-width or
width-for-height geometry management.</p>
<div class="api-notes">
  <p class="api-ctype">GtkSizeRequestMode</p>
<table>
<tr>
<td class="name">GTK_SIZE_REQUEST_HEIGHT_FOR_WIDTH</td>
<td class="value">0</td>
<td class="doc">Prefer height-for-width geometry management</td>
</tr>
<tr>
<td class="name">GTK_SIZE_REQUEST_WIDTH_FOR_HEIGHT</td>
<td class="value">1</td>
<td class="doc">Prefer width-for-height geometry management</td>
</tr>
<tr>
<td class="name">GTK_SIZE_REQUEST_CONSTANT_SIZE</td>
<td class="value">2</td>
<td class="doc">Don’t trade height-for-width or width-for-height</td>
</tr>
</table>
</div>
<p class="api-heading">SortType</p>
<p class="api-doc">Determines the direction of a sort.</p>
<div class="api-notes">
  <p class="api-ctype">GtkSortType</p>
<table>
<tr>
<td class="name">GTK_SORT_ASCENDING</td>
<td class="value">0</td>
<td class="doc">Sorting is in ascending order.</td>
</tr>
<tr>
<td class="name">GTK_SORT_DESCENDING</td>
<td class="value">1</td>
<td class="doc">Sorting is in descending order.</td>
</tr>
</table>
</div>
<p class="api-heading">SpinButtonUpdatePolicy</p>
<p class="api-doc">The spin button update policy determines whether the spin button displays
values even if they are outside the bounds of its adjustment.
See gtk_spin_button_set_update_policy().</p>
<div class="api-notes">
  <p class="api-ctype">GtkSpinButtonUpdatePolicy</p>
<table>
<tr>
<td class="name">GTK_UPDATE_ALWAYS</td>
<td class="value">0</td>
<td class="doc">When refreshing your #GtkSpinButton, the value is
    always displayed</td>
</tr>
<tr>
<td class="name">GTK_UPDATE_IF_VALID</td>
<td class="value">1</td>
<td class="doc">When refreshing your #GtkSpinButton, the value is
    only displayed if it is valid within the bounds of the spin button's
    adjustment</td>
</tr>
</table>
</div>
<p class="api-heading">SpinType</p>
<p class="api-doc">The values of the GtkSpinType enumeration are used to specify the
change to make in gtk_spin_button_spin().</p>
<div class="api-notes">
  <p class="api-ctype">GtkSpinType</p>
<table>
<tr>
<td class="name">GTK_SPIN_STEP_FORWARD</td>
<td class="value">0</td>
<td class="doc">Increment by the adjustments step increment.</td>
</tr>
<tr>
<td class="name">GTK_SPIN_STEP_BACKWARD</td>
<td class="value">1</td>
<td class="doc">Decrement by the adjustments step increment.</td>
</tr>
<tr>
<td class="name">GTK_SPIN_PAGE_FORWARD</td>
<td class="value">2</td>
<td class="doc">Increment by the adjustments page increment.</td>
</tr>
<tr>
<td class="name">GTK_SPIN_PAGE_BACKWARD</td>
<td class="value">3</td>
<td class="doc">Decrement by the adjustments page increment.</td>
</tr>
<tr>
<td class="name">GTK_SPIN_HOME</td>
<td class="value">4</td>
<td class="doc">Go to the adjustments lower bound.</td>
</tr>
<tr>
<td class="name">GTK_SPIN_END</td>
<td class="value">5</td>
<td class="doc">Go to the adjustments upper bound.</td>
</tr>
<tr>
<td class="name">GTK_SPIN_USER_DEFINED</td>
<td class="value">6</td>
<td class="doc">Change by a specified amount.</td>
</tr>
</table>
</div>
<p class="api-heading">StackTransitionType</p>
<p class="api-doc">These enumeration values describe the possible transitions
between pages in a #GtkStack widget.

New values may be added to this enumeration over time.</p>
<div class="api-notes">
  <p class="api-ctype">GtkStackTransitionType</p>
<table>
<tr>
<td class="name">GTK_STACK_TRANSITION_TYPE_NONE</td>
<td class="value">0</td>
<td class="doc">No transition</td>
</tr>
<tr>
<td class="name">GTK_STACK_TRANSITION_TYPE_CROSSFADE</td>
<td class="value">1</td>
<td class="doc">A cross-fade</td>
</tr>
<tr>
<td class="name">GTK_STACK_TRANSITION_TYPE_SLIDE_RIGHT</td>
<td class="value">2</td>
<td class="doc">Slide from left to right</td>
</tr>
<tr>
<td class="name">GTK_STACK_TRANSITION_TYPE_SLIDE_LEFT</td>
<td class="value">3</td>
<td class="doc">Slide from right to left</td>
</tr>
<tr>
<td class="name">GTK_STACK_TRANSITION_TYPE_SLIDE_UP</td>
<td class="value">4</td>
<td class="doc">Slide from bottom up</td>
</tr>
<tr>
<td class="name">GTK_STACK_TRANSITION_TYPE_SLIDE_DOWN</td>
<td class="value">5</td>
<td class="doc">Slide from top down</td>
</tr>
<tr>
<td class="name">GTK_STACK_TRANSITION_TYPE_SLIDE_LEFT_RIGHT</td>
<td class="value">6</td>
<td class="doc">Slide from left or right according to the children order</td>
</tr>
<tr>
<td class="name">GTK_STACK_TRANSITION_TYPE_SLIDE_UP_DOWN</td>
<td class="value">7</td>
<td class="doc">Slide from top down or bottom up according to the order</td>
</tr>
<tr>
<td class="name">GTK_STACK_TRANSITION_TYPE_OVER_UP</td>
<td class="value">8</td>
<td class="doc">Cover the old page by sliding up. Since 3.12</td>
</tr>
<tr>
<td class="name">GTK_STACK_TRANSITION_TYPE_OVER_DOWN</td>
<td class="value">9</td>
<td class="doc">Cover the old page by sliding down. Since: 3.12</td>
</tr>
<tr>
<td class="name">GTK_STACK_TRANSITION_TYPE_OVER_LEFT</td>
<td class="value">10</td>
<td class="doc">Cover the old page by sliding to the left. Since: 3.12</td>
</tr>
<tr>
<td class="name">GTK_STACK_TRANSITION_TYPE_OVER_RIGHT</td>
<td class="value">11</td>
<td class="doc">Cover the old page by sliding to the right. Since: 3.12</td>
</tr>
<tr>
<td class="name">GTK_STACK_TRANSITION_TYPE_UNDER_UP</td>
<td class="value">12</td>
<td class="doc">Uncover the new page by sliding up. Since 3.12</td>
</tr>
<tr>
<td class="name">GTK_STACK_TRANSITION_TYPE_UNDER_DOWN</td>
<td class="value">13</td>
<td class="doc">Uncover the new page by sliding down. Since: 3.12</td>
</tr>
<tr>
<td class="name">GTK_STACK_TRANSITION_TYPE_UNDER_LEFT</td>
<td class="value">14</td>
<td class="doc">Uncover the new page by sliding to the left. Since: 3.12</td>
</tr>
<tr>
<td class="name">GTK_STACK_TRANSITION_TYPE_UNDER_RIGHT</td>
<td class="value">15</td>
<td class="doc">Uncover the new page by sliding to the right. Since: 3.12</td>
</tr>
<tr>
<td class="name">GTK_STACK_TRANSITION_TYPE_OVER_UP_DOWN</td>
<td class="value">16</td>
<td class="doc">Cover the old page sliding up or uncover the new page sliding down, according to order. Since: 3.12</td>
</tr>
<tr>
<td class="name">GTK_STACK_TRANSITION_TYPE_OVER_DOWN_UP</td>
<td class="value">17</td>
<td class="doc">Cover the old page sliding down or uncover the new page sliding up, according to order. Since: 3.14</td>
</tr>
<tr>
<td class="name">GTK_STACK_TRANSITION_TYPE_OVER_LEFT_RIGHT</td>
<td class="value">18</td>
<td class="doc">Cover the old page sliding left or uncover the new page sliding right, according to order. Since: 3.14</td>
</tr>
<tr>
<td class="name">GTK_STACK_TRANSITION_TYPE_OVER_RIGHT_LEFT</td>
<td class="value">19</td>
<td class="doc">Cover the old page sliding right or uncover the new page sliding left, according to order. Since: 3.14</td>
</tr>
</table>
</div>
<p class="api-heading">StateType</p>
<p class="api-doc">This type indicates the current state of a widget; the state determines how
the widget is drawn. The #GtkStateType enumeration is also used to
identify different colors in a #GtkStyle for drawing, so states can be
used for subparts of a widget as well as entire widgets.</p>
<div class="api-notes">
  <p class="api-ctype">GtkStateType</p>
<table>
<tr>
<td class="name">GTK_STATE_NORMAL</td>
<td class="value">0</td>
<td class="doc">State during normal operation.</td>
</tr>
<tr>
<td class="name">GTK_STATE_ACTIVE</td>
<td class="value">1</td>
<td class="doc">State of a currently active widget, such as a depressed button.</td>
</tr>
<tr>
<td class="name">GTK_STATE_PRELIGHT</td>
<td class="value">2</td>
<td class="doc">State indicating that the mouse pointer is over
                     the widget and the widget will respond to mouse clicks.</td>
</tr>
<tr>
<td class="name">GTK_STATE_SELECTED</td>
<td class="value">3</td>
<td class="doc">State of a selected item, such the selected row in a list.</td>
</tr>
<tr>
<td class="name">GTK_STATE_INSENSITIVE</td>
<td class="value">4</td>
<td class="doc">State indicating that the widget is
                        unresponsive to user actions.</td>
</tr>
<tr>
<td class="name">GTK_STATE_INCONSISTENT</td>
<td class="value">5</td>
<td class="doc">The widget is inconsistent, such as checkbuttons
                         or radiobuttons that aren’t either set to %TRUE nor %FALSE,
                         or buttons requiring the user attention.</td>
</tr>
<tr>
<td class="name">GTK_STATE_FOCUSED</td>
<td class="value">6</td>
<td class="doc">The widget has the keyboard focus.</td>
</tr>
</table>
</div>
<p class="api-heading">TextBufferTargetInfo</p>
<p class="api-doc">These values are used as “info” for the targets contained in the
lists returned by gtk_text_buffer_get_copy_target_list() and
gtk_text_buffer_get_paste_target_list().

The values counts down from `-1` to avoid clashes
with application added drag destinations which usually start at 0.</p>
<div class="api-notes">
  <p class="api-ctype">GtkTextBufferTargetInfo</p>
<table>
<tr>
<td class="name">GTK_TEXT_BUFFER_TARGET_INFO_BUFFER_CONTENTS</td>
<td class="value">-1</td>
<td class="doc">Buffer contents</td>
</tr>
<tr>
<td class="name">GTK_TEXT_BUFFER_TARGET_INFO_RICH_TEXT</td>
<td class="value">-2</td>
<td class="doc">Rich text</td>
</tr>
<tr>
<td class="name">GTK_TEXT_BUFFER_TARGET_INFO_TEXT</td>
<td class="value">-3</td>
<td class="doc">Text</td>
</tr>
</table>
</div>
<p class="api-heading">TextDirection</p>
<p class="api-doc">Reading directions for text.</p>
<div class="api-notes">
  <p class="api-ctype">GtkTextDirection</p>
<table>
<tr>
<td class="name">GTK_TEXT_DIR_NONE</td>
<td class="value">0</td>
<td class="doc">No direction.</td>
</tr>
<tr>
<td class="name">GTK_TEXT_DIR_LTR</td>
<td class="value">1</td>
<td class="doc">Left to right text direction.</td>
</tr>
<tr>
<td class="name">GTK_TEXT_DIR_RTL</td>
<td class="value">2</td>
<td class="doc">Right to left text direction.</td>
</tr>
</table>
</div>
<p class="api-heading">TextExtendSelection</p>
<p class="api-doc">Granularity types that extend the text selection. Use the
#GtkTextView::extend-selection signal to customize the selection.</p>
<div class="api-notes">
  <p class="api-ctype">GtkTextExtendSelection</p>
  <p class="api-since">since 3.16</p>
<table>
<tr>
<td class="name">GTK_TEXT_EXTEND_SELECTION_WORD</td>
<td class="value">0</td>
<td class="doc">Selects the current word. It is triggered by
  a double-click for example.</td>
</tr>
<tr>
<td class="name">GTK_TEXT_EXTEND_SELECTION_LINE</td>
<td class="value">1</td>
<td class="doc">Selects the current line. It is triggered by
  a triple-click for example.</td>
</tr>
</table>
</div>
<p class="api-heading">TextViewLayer</p>
<p class="api-doc">Used to reference the layers of #GtkTextView for the purpose of customized
drawing with the ::draw_layer vfunc.</p>
<div class="api-notes">
  <p class="api-ctype">GtkTextViewLayer</p>
<table>
<tr>
<td class="name">GTK_TEXT_VIEW_LAYER_BELOW</td>
<td class="value">0</td>
<td class="doc">Old deprecated layer, use %GTK_TEXT_VIEW_LAYER_BELOW_TEXT instead</td>
</tr>
<tr>
<td class="name">GTK_TEXT_VIEW_LAYER_ABOVE</td>
<td class="value">1</td>
<td class="doc">Old deprecated layer, use %GTK_TEXT_VIEW_LAYER_ABOVE_TEXT instead</td>
</tr>
<tr>
<td class="name">GTK_TEXT_VIEW_LAYER_BELOW_TEXT</td>
<td class="value">2</td>
<td class="doc">The layer rendered below the text (but above the background).  Since: 3.20</td>
</tr>
<tr>
<td class="name">GTK_TEXT_VIEW_LAYER_ABOVE_TEXT</td>
<td class="value">3</td>
<td class="doc">The layer rendered above the text.  Since: 3.20</td>
</tr>
</table>
</div>
<p class="api-heading">TextWindowType</p>
<p class="api-doc">Used to reference the parts of #GtkTextView.</p>
<div class="api-notes">
  <p class="api-ctype">GtkTextWindowType</p>
<table>
<tr>
<td class="name">GTK_TEXT_WINDOW_PRIVATE</td>
<td class="value">0</td>
<td class="doc">Invalid value, used as a marker</td>
</tr>
<tr>
<td class="name">GTK_TEXT_WINDOW_WIDGET</td>
<td class="value">1</td>
<td class="doc">Window that floats over scrolling areas.</td>
</tr>
<tr>
<td class="name">GTK_TEXT_WINDOW_TEXT</td>
<td class="value">2</td>
<td class="doc">Scrollable text window.</td>
</tr>
<tr>
<td class="name">GTK_TEXT_WINDOW_LEFT</td>
<td class="value">3</td>
<td class="doc">Left side border window.</td>
</tr>
<tr>
<td class="name">GTK_TEXT_WINDOW_RIGHT</td>
<td class="value">4</td>
<td class="doc">Right side border window.</td>
</tr>
<tr>
<td class="name">GTK_TEXT_WINDOW_TOP</td>
<td class="value">5</td>
<td class="doc">Top border window.</td>
</tr>
<tr>
<td class="name">GTK_TEXT_WINDOW_BOTTOM</td>
<td class="value">6</td>
<td class="doc">Bottom border window.</td>
</tr>
</table>
</div>
<p class="api-heading">ToolbarSpaceStyle</p>
<p class="api-doc">Whether spacers are vertical lines or just blank.</p>
<div class="api-notes">
  <p class="api-ctype">GtkToolbarSpaceStyle</p>
<table>
<tr>
<td class="name">GTK_TOOLBAR_SPACE_EMPTY</td>
<td class="value">0</td>
<td class="doc">Use blank spacers.</td>
</tr>
<tr>
<td class="name">GTK_TOOLBAR_SPACE_LINE</td>
<td class="value">1</td>
<td class="doc">Use vertical lines for spacers.</td>
</tr>
</table>
</div>
<p class="api-heading">ToolbarStyle</p>
<p class="api-doc">Used to customize the appearance of a #GtkToolbar. Note that
setting the toolbar style overrides the user’s preferences
for the default toolbar style.  Note that if the button has only
a label set and GTK_TOOLBAR_ICONS is used, the label will be
visible, and vice versa.</p>
<div class="api-notes">
  <p class="api-ctype">GtkToolbarStyle</p>
<table>
<tr>
<td class="name">GTK_TOOLBAR_ICONS</td>
<td class="value">0</td>
<td class="doc">Buttons display only icons in the toolbar.</td>
</tr>
<tr>
<td class="name">GTK_TOOLBAR_TEXT</td>
<td class="value">1</td>
<td class="doc">Buttons display only text labels in the toolbar.</td>
</tr>
<tr>
<td class="name">GTK_TOOLBAR_BOTH</td>
<td class="value">2</td>
<td class="doc">Buttons display text and icons in the toolbar.</td>
</tr>
<tr>
<td class="name">GTK_TOOLBAR_BOTH_HORIZ</td>
<td class="value">3</td>
<td class="doc">Buttons display icons and text alongside each
 other, rather than vertically stacked</td>
</tr>
</table>
</div>
<p class="api-heading">TreeViewColumnSizing</p>
<p class="api-doc">The sizing method the column uses to determine its width.  Please note
that @GTK_TREE_VIEW_COLUMN_AUTOSIZE are inefficient for large views, and
can make columns appear choppy.</p>
<div class="api-notes">
  <p class="api-ctype">GtkTreeViewColumnSizing</p>
<table>
<tr>
<td class="name">GTK_TREE_VIEW_COLUMN_GROW_ONLY</td>
<td class="value">0</td>
<td class="doc">Columns only get bigger in reaction to changes in the model</td>
</tr>
<tr>
<td class="name">GTK_TREE_VIEW_COLUMN_AUTOSIZE</td>
<td class="value">1</td>
<td class="doc">Columns resize to be the optimal size everytime the model changes.</td>
</tr>
<tr>
<td class="name">GTK_TREE_VIEW_COLUMN_FIXED</td>
<td class="value">2</td>
<td class="doc">Columns are a fixed numbers of pixels wide.</td>
</tr>
</table>
</div>
<p class="api-heading">TreeViewDropPosition</p>
<p class="api-doc">An enum for determining where a dropped row goes.</p>
<div class="api-notes">
  <p class="api-ctype">GtkTreeViewDropPosition</p>
<table>
<tr>
<td class="name">GTK_TREE_VIEW_DROP_BEFORE</td>
<td class="value">0</td>
<td class="doc">dropped row is inserted before</td>
</tr>
<tr>
<td class="name">GTK_TREE_VIEW_DROP_AFTER</td>
<td class="value">1</td>
<td class="doc">dropped row is inserted after</td>
</tr>
<tr>
<td class="name">GTK_TREE_VIEW_DROP_INTO_OR_BEFORE</td>
<td class="value">2</td>
<td class="doc">dropped row becomes a child or is inserted before</td>
</tr>
<tr>
<td class="name">GTK_TREE_VIEW_DROP_INTO_OR_AFTER</td>
<td class="value">3</td>
<td class="doc">dropped row becomes a child or is inserted after</td>
</tr>
</table>
</div>
<p class="api-heading">TreeViewGridLines</p>
<p class="api-doc">Used to indicate which grid lines to draw in a tree view.</p>
<div class="api-notes">
  <p class="api-ctype">GtkTreeViewGridLines</p>
<table>
<tr>
<td class="name">GTK_TREE_VIEW_GRID_LINES_NONE</td>
<td class="value">0</td>
<td class="doc">No grid lines.</td>
</tr>
<tr>
<td class="name">GTK_TREE_VIEW_GRID_LINES_HORIZONTAL</td>
<td class="value">1</td>
<td class="doc">Horizontal grid lines.</td>
</tr>
<tr>
<td class="name">GTK_TREE_VIEW_GRID_LINES_VERTICAL</td>
<td class="value">2</td>
<td class="doc">Vertical grid lines.</td>
</tr>
<tr>
<td class="name">GTK_TREE_VIEW_GRID_LINES_BOTH</td>
<td class="value">3</td>
<td class="doc">Horizontal and vertical grid lines.</td>
</tr>
</table>
</div>
<p class="api-heading">Unit</p>
<p class="api-doc">See also gtk_print_settings_set_paper_width().</p>
<div class="api-notes">
  <p class="api-ctype">GtkUnit</p>
<table>
<tr>
<td class="name">GTK_UNIT_NONE</td>
<td class="value">0</td>
<td class="doc">No units.</td>
</tr>
<tr>
<td class="name">GTK_UNIT_POINTS</td>
<td class="value">1</td>
<td class="doc">Dimensions in points.</td>
</tr>
<tr>
<td class="name">GTK_UNIT_INCH</td>
<td class="value">2</td>
<td class="doc">Dimensions in inches.</td>
</tr>
<tr>
<td class="name">GTK_UNIT_MM</td>
<td class="value">3</td>
<td class="doc">Dimensions in millimeters</td>
</tr>
</table>
</div>
<p class="api-heading">WidgetHelpType</p>
<p class="api-doc">Kinds of widget-specific help. Used by the ::show-help signal.</p>
<div class="api-notes">
  <p class="api-ctype">GtkWidgetHelpType</p>
<table>
<tr>
<td class="name">GTK_WIDGET_HELP_TOOLTIP</td>
<td class="value">0</td>
<td class="doc">Tooltip.</td>
</tr>
<tr>
<td class="name">GTK_WIDGET_HELP_WHATS_THIS</td>
<td class="value">1</td>
<td class="doc">What’s this.</td>
</tr>
</table>
</div>
<p class="api-heading">WindowPosition</p>
<p class="api-doc">Window placement can be influenced using this enumeration. Note that
using #GTK_WIN_POS_CENTER_ALWAYS is almost always a bad idea.
It won’t necessarily work well with all window managers or on all windowing systems.</p>
<div class="api-notes">
  <p class="api-ctype">GtkWindowPosition</p>
<table>
<tr>
<td class="name">GTK_WIN_POS_NONE</td>
<td class="value">0</td>
<td class="doc">No influence is made on placement.</td>
</tr>
<tr>
<td class="name">GTK_WIN_POS_CENTER</td>
<td class="value">1</td>
<td class="doc">Windows should be placed in the center of the screen.</td>
</tr>
<tr>
<td class="name">GTK_WIN_POS_MOUSE</td>
<td class="value">2</td>
<td class="doc">Windows should be placed at the current mouse position.</td>
</tr>
<tr>
<td class="name">GTK_WIN_POS_CENTER_ALWAYS</td>
<td class="value">3</td>
<td class="doc">Keep window centered as it changes size, etc.</td>
</tr>
<tr>
<td class="name">GTK_WIN_POS_CENTER_ON_PARENT</td>
<td class="value">4</td>
<td class="doc">Center the window on its transient
 parent (see gtk_window_set_transient_for()).</td>
</tr>
</table>
</div>
<p class="api-heading">WindowType</p>
<p class="api-doc">A #GtkWindow can be one of these types. Most things you’d consider a
“window” should have type #GTK_WINDOW_TOPLEVEL; windows with this type
are managed by the window manager and have a frame by default (call
gtk_window_set_decorated() to toggle the frame).  Windows with type
#GTK_WINDOW_POPUP are ignored by the window manager; window manager
keybindings won’t work on them, the window manager won’t decorate the
window with a frame, many GTK+ features that rely on the window
manager will not work (e.g. resize grips and
maximization/minimization). #GTK_WINDOW_POPUP is used to implement
widgets such as #GtkMenu or tooltips that you normally don’t think of
as windows per se. Nearly all windows should be #GTK_WINDOW_TOPLEVEL.
In particular, do not use #GTK_WINDOW_POPUP just to turn off
the window borders; use gtk_window_set_decorated() for that.</p>
<div class="api-notes">
  <p class="api-ctype">GtkWindowType</p>
<table>
<tr>
<td class="name">GTK_WINDOW_TOPLEVEL</td>
<td class="value">0</td>
<td class="doc">A regular window, such as a dialog.</td>
</tr>
<tr>
<td class="name">GTK_WINDOW_POPUP</td>
<td class="value">1</td>
<td class="doc">A special window such as a tooltip.</td>
</tr>
</table>
</div>
<p class="api-heading">WrapMode</p>
<p class="api-doc">Describes a type of line wrapping.</p>
<div class="api-notes">
  <p class="api-ctype">GtkWrapMode</p>
<table>
<tr>
<td class="name">GTK_WRAP_NONE</td>
<td class="value">0</td>
<td class="doc">do not wrap lines; just make the text area wider</td>
</tr>
<tr>
<td class="name">GTK_WRAP_CHAR</td>
<td class="value">1</td>
<td class="doc">wrap text, breaking lines anywhere the cursor can
    appear (between characters, usually - if you want to be technical,
    between graphemes, see pango_get_log_attrs())</td>
</tr>
<tr>
<td class="name">GTK_WRAP_WORD</td>
<td class="value">2</td>
<td class="doc">wrap text, breaking lines in between words</td>
</tr>
<tr>
<td class="name">GTK_WRAP_WORD_CHAR</td>
<td class="value">3</td>
<td class="doc">wrap text, breaking lines in between words, or if
    that is not enough, also between graphemes</td>
</tr>
</table>
</div>