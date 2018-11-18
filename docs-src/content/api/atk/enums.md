+++
title = "enums"
+++
<p class="api-heading">CoordType</p>
<p class="api-doc">Specifies how xy coordinates are to be interpreted. Used by functions such
as atk_component_get_position() and atk_text_get_character_extents()</p>
<div class="api-notes">
  <p class="api-ctype">AtkCoordType</p>
<table>
<tr>
<td class="name">ATK_XY_SCREEN</td>
<td class="value">0</td>
<td class="doc">specifies xy coordinates relative to the screen</td>
</tr>
<tr>
<td class="name">ATK_XY_WINDOW</td>
<td class="value">1</td>
<td class="doc">specifies xy coordinates relative to the widget's
top-level window</td>
</tr>
</table>
</div>
<p class="api-heading">KeyEventType</p>
<p class="api-doc">Specifies the type of a keyboard evemt.</p>
<div class="api-notes">
  <p class="api-ctype">AtkKeyEventType</p>
<table>
<tr>
<td class="name">ATK_KEY_EVENT_PRESS</td>
<td class="value">0</td>
<td class="doc">specifies a key press event</td>
</tr>
<tr>
<td class="name">ATK_KEY_EVENT_RELEASE</td>
<td class="value">1</td>
<td class="doc">specifies a key release event</td>
</tr>
<tr>
<td class="name">ATK_KEY_EVENT_LAST_DEFINED</td>
<td class="value">2</td>
<td class="doc">Not a valid value; specifies end of enumeration</td>
</tr>
</table>
</div>
<p class="api-heading">Layer</p>
<p class="api-doc">Describes the layer of a component

These enumerated "layer values" are used when determining which UI
rendering layer a component is drawn into, which can help in making
determinations of when components occlude one another.</p>
<div class="api-notes">
  <p class="api-ctype">AtkLayer</p>
<table>
<tr>
<td class="name">ATK_LAYER_INVALID</td>
<td class="value">0</td>
<td class="doc">The object does not have a layer</td>
</tr>
<tr>
<td class="name">ATK_LAYER_BACKGROUND</td>
<td class="value">1</td>
<td class="doc">This layer is reserved for the desktop background</td>
</tr>
<tr>
<td class="name">ATK_LAYER_CANVAS</td>
<td class="value">2</td>
<td class="doc">This layer is used for Canvas components</td>
</tr>
<tr>
<td class="name">ATK_LAYER_WIDGET</td>
<td class="value">3</td>
<td class="doc">This layer is normally used for components</td>
</tr>
<tr>
<td class="name">ATK_LAYER_MDI</td>
<td class="value">4</td>
<td class="doc">This layer is used for layered components</td>
</tr>
<tr>
<td class="name">ATK_LAYER_POPUP</td>
<td class="value">5</td>
<td class="doc">This layer is used for popup components, such as menus</td>
</tr>
<tr>
<td class="name">ATK_LAYER_OVERLAY</td>
<td class="value">6</td>
<td class="doc">This layer is reserved for future use.</td>
</tr>
<tr>
<td class="name">ATK_LAYER_WINDOW</td>
<td class="value">7</td>
<td class="doc">This layer is used for toplevel windows.</td>
</tr>
</table>
</div>
<p class="api-heading">RelationType</p>
<p class="api-doc">Describes the type of the relation</p>
<div class="api-notes">
  <p class="api-ctype">AtkRelationType</p>
<table>
<tr>
<td class="name">ATK_RELATION_NULL</td>
<td class="value">0</td>
<td class="doc">Not used, represens "no relationship" or an error condition.</td>
</tr>
<tr>
<td class="name">ATK_RELATION_CONTROLLED_BY</td>
<td class="value">1</td>
<td class="doc">Indicates an object controlled by one or more target objects.</td>
</tr>
<tr>
<td class="name">ATK_RELATION_CONTROLLER_FOR</td>
<td class="value">2</td>
<td class="doc">Indicates an object is an controller for one or more target objects.</td>
</tr>
<tr>
<td class="name">ATK_RELATION_LABEL_FOR</td>
<td class="value">3</td>
<td class="doc">Indicates an object is a label for one or more target objects.</td>
</tr>
<tr>
<td class="name">ATK_RELATION_LABELLED_BY</td>
<td class="value">4</td>
<td class="doc">Indicates an object is labelled by one or more target objects.</td>
</tr>
<tr>
<td class="name">ATK_RELATION_MEMBER_OF</td>
<td class="value">5</td>
<td class="doc">Indicates an object is a member of a group of one or more target objects.</td>
</tr>
<tr>
<td class="name">ATK_RELATION_NODE_CHILD_OF</td>
<td class="value">6</td>
<td class="doc">Indicates an object is a cell in a treetable which is displayed because a cell in the same column is expanded and identifies that cell.</td>
</tr>
<tr>
<td class="name">ATK_RELATION_FLOWS_TO</td>
<td class="value">7</td>
<td class="doc">Indicates that the object has content that flows logically to another
 AtkObject in a sequential way, (for instance text-flow).</td>
</tr>
<tr>
<td class="name">ATK_RELATION_FLOWS_FROM</td>
<td class="value">8</td>
<td class="doc">Indicates that the object has content that flows logically from
 another AtkObject in a sequential way, (for instance text-flow).</td>
</tr>
<tr>
<td class="name">ATK_RELATION_SUBWINDOW_OF</td>
<td class="value">9</td>
<td class="doc">Indicates a subwindow attached to a component but otherwise has no connection in  the UI heirarchy to that component.</td>
</tr>
<tr>
<td class="name">ATK_RELATION_EMBEDS</td>
<td class="value">10</td>
<td class="doc">Indicates that the object visually embeds
 another object's content, i.e. this object's content flows around
 another's content.</td>
</tr>
<tr>
<td class="name">ATK_RELATION_EMBEDDED_BY</td>
<td class="value">11</td>
<td class="doc">Reciprocal of %ATK_RELATION_EMBEDS, indicates that
 this object's content is visualy embedded in another object.</td>
</tr>
<tr>
<td class="name">ATK_RELATION_POPUP_FOR</td>
<td class="value">12</td>
<td class="doc">Indicates that an object is a popup for another object.</td>
</tr>
<tr>
<td class="name">ATK_RELATION_PARENT_WINDOW_OF</td>
<td class="value">13</td>
<td class="doc">Indicates that an object is a parent window of another object.</td>
</tr>
<tr>
<td class="name">ATK_RELATION_DESCRIBED_BY</td>
<td class="value">14</td>
<td class="doc">Reciprocal of %ATK_RELATION_DESCRIPTION_FOR. Indicates that one
or more target objects provide descriptive information about this object. This relation
type is most appropriate for information that is not essential as its presentation may
be user-configurable and/or limited to an on-demand mechanism such as an assistive
technology command. For brief, essential information such as can be found in a widget's
on-screen label, use %ATK_RELATION_LABELLED_BY. For an on-screen error message, use
%ATK_RELATION_ERROR_MESSAGE. For lengthy extended descriptive information contained in
an on-screen object, consider using %ATK_RELATION_DETAILS as assistive technologies may
provide a means for the user to navigate to objects containing detailed descriptions so
that their content can be more closely reviewed.</td>
</tr>
<tr>
<td class="name">ATK_RELATION_DESCRIPTION_FOR</td>
<td class="value">15</td>
<td class="doc">Reciprocal of %ATK_RELATION_DESCRIBED_BY. Indicates that this
object provides descriptive information about the target object(s). See also
%ATK_RELATION_DETAILS_FOR and %ATK_RELATION_ERROR_FOR.</td>
</tr>
<tr>
<td class="name">ATK_RELATION_NODE_PARENT_OF</td>
<td class="value">16</td>
<td class="doc">Indicates an object is a cell in a treetable and is expanded to display other cells in the same column.</td>
</tr>
<tr>
<td class="name">ATK_RELATION_DETAILS</td>
<td class="value">17</td>
<td class="doc">Reciprocal of %ATK_RELATION_DETAILS_FOR. Indicates that this object
has a detailed or extended description, the contents of which can be found in the target
object(s). This relation type is most appropriate for information that is sufficiently
lengthy as to make navigation to the container of that information desirable. For less
verbose information suitable for announcement only, see %ATK_RELATION_DESCRIBED_BY. If
the detailed information describes an error condition, %ATK_RELATION_ERROR_FOR should be
used instead. @Since: ATK-2.26.</td>
</tr>
<tr>
<td class="name">ATK_RELATION_DETAILS_FOR</td>
<td class="value">18</td>
<td class="doc">Reciprocal of %ATK_RELATION_DETAILS. Indicates that this object
provides a detailed or extended description about the target object(s). See also
%ATK_RELATION_DESCRIPTION_FOR and %ATK_RELATION_ERROR_FOR. @Since: ATK-2.26.</td>
</tr>
<tr>
<td class="name">ATK_RELATION_ERROR_MESSAGE</td>
<td class="value">19</td>
<td class="doc">Reciprocal of %ATK_RELATION_ERROR_FOR. Indicates that this object
has one or more errors, the nature of which is described in the contents of the target
object(s). Objects that have this relation type should also contain %ATK_STATE_INVALID_ENTRY
in their #AtkStateSet. @Since: ATK-2.26.</td>
</tr>
<tr>
<td class="name">ATK_RELATION_ERROR_FOR</td>
<td class="value">20</td>
<td class="doc">Reciprocal of %ATK_RELATION_ERROR_MESSAGE. Indicates that this object
contains an error message describing an invalid condition in the target object(s). @Since:
ATK_2.26.</td>
</tr>
<tr>
<td class="name">ATK_RELATION_LAST_DEFINED</td>
<td class="value">21</td>
<td class="doc">Not used, this value indicates the end of the enumeration.</td>
</tr>
</table>
</div>
<p class="api-heading">Role</p>
<p class="api-doc">Describes the role of an object

These are the built-in enumerated roles that UI components can have in
ATK.  Other roles may be added at runtime, so an AtkRole >=
ATK_ROLE_LAST_DEFINED is not necessarily an error.</p>
<div class="api-notes">
  <p class="api-ctype">AtkRole</p>
<table>
<tr>
<td class="name">ATK_ROLE_INVALID</td>
<td class="value">0</td>
<td class="doc">Invalid role</td>
</tr>
<tr>
<td class="name">ATK_ROLE_ACCEL_LABEL</td>
<td class="value">1</td>
<td class="doc">A label which represents an accelerator</td>
</tr>
<tr>
<td class="name">ATK_ROLE_ALERT</td>
<td class="value">2</td>
<td class="doc">An object which is an alert to the user. Assistive Technologies typically respond to ATK_ROLE_ALERT by reading the entire onscreen contents of containers advertising this role.  Should be used for warning dialogs, etc.</td>
</tr>
<tr>
<td class="name">ATK_ROLE_ANIMATION</td>
<td class="value">3</td>
<td class="doc">An object which is an animated image</td>
</tr>
<tr>
<td class="name">ATK_ROLE_ARROW</td>
<td class="value">4</td>
<td class="doc">An arrow in one of the four cardinal directions</td>
</tr>
<tr>
<td class="name">ATK_ROLE_CALENDAR</td>
<td class="value">5</td>
<td class="doc">An object that displays a calendar and allows the user to select a date</td>
</tr>
<tr>
<td class="name">ATK_ROLE_CANVAS</td>
<td class="value">6</td>
<td class="doc">An object that can be drawn into and is used to trap events</td>
</tr>
<tr>
<td class="name">ATK_ROLE_CHECK_BOX</td>
<td class="value">7</td>
<td class="doc">A choice that can be checked or unchecked and provides a separate indicator for the current state</td>
</tr>
<tr>
<td class="name">ATK_ROLE_CHECK_MENU_ITEM</td>
<td class="value">8</td>
<td class="doc">A menu item with a check box</td>
</tr>
<tr>
<td class="name">ATK_ROLE_COLOR_CHOOSER</td>
<td class="value">9</td>
<td class="doc">A specialized dialog that lets the user choose a color</td>
</tr>
<tr>
<td class="name">ATK_ROLE_COLUMN_HEADER</td>
<td class="value">10</td>
<td class="doc">The header for a column of data</td>
</tr>
<tr>
<td class="name">ATK_ROLE_COMBO_BOX</td>
<td class="value">11</td>
<td class="doc">A collapsible list of choices the user can select from</td>
</tr>
<tr>
<td class="name">ATK_ROLE_DATE_EDITOR</td>
<td class="value">12</td>
<td class="doc">An object whose purpose is to allow a user to edit a date</td>
</tr>
<tr>
<td class="name">ATK_ROLE_DESKTOP_ICON</td>
<td class="value">13</td>
<td class="doc">An inconifed internal frame within a DESKTOP_PANE</td>
</tr>
<tr>
<td class="name">ATK_ROLE_DESKTOP_FRAME</td>
<td class="value">14</td>
<td class="doc">A pane that supports internal frames and iconified versions of those internal frames</td>
</tr>
<tr>
<td class="name">ATK_ROLE_DIAL</td>
<td class="value">15</td>
<td class="doc">An object whose purpose is to allow a user to set a value</td>
</tr>
<tr>
<td class="name">ATK_ROLE_DIALOG</td>
<td class="value">16</td>
<td class="doc">A top level window with title bar and a border</td>
</tr>
<tr>
<td class="name">ATK_ROLE_DIRECTORY_PANE</td>
<td class="value">17</td>
<td class="doc">A pane that allows the user to navigate through and select the contents of a directory</td>
</tr>
<tr>
<td class="name">ATK_ROLE_DRAWING_AREA</td>
<td class="value">18</td>
<td class="doc">An object used for drawing custom user interface elements</td>
</tr>
<tr>
<td class="name">ATK_ROLE_FILE_CHOOSER</td>
<td class="value">19</td>
<td class="doc">A specialized dialog that lets the user choose a file</td>
</tr>
<tr>
<td class="name">ATK_ROLE_FILLER</td>
<td class="value">20</td>
<td class="doc">A object that fills up space in a user interface</td>
</tr>
<tr>
<td class="name">ATK_ROLE_FONT_CHOOSER</td>
<td class="value">21</td>
<td class="doc">A specialized dialog that lets the user choose a font</td>
</tr>
<tr>
<td class="name">ATK_ROLE_FRAME</td>
<td class="value">22</td>
<td class="doc">A top level window with a title bar, border, menubar, etc.</td>
</tr>
<tr>
<td class="name">ATK_ROLE_GLASS_PANE</td>
<td class="value">23</td>
<td class="doc">A pane that is guaranteed to be painted on top of all panes beneath it</td>
</tr>
<tr>
<td class="name">ATK_ROLE_HTML_CONTAINER</td>
<td class="value">24</td>
<td class="doc">A document container for HTML, whose children represent the document content</td>
</tr>
<tr>
<td class="name">ATK_ROLE_ICON</td>
<td class="value">25</td>
<td class="doc">A small fixed size picture, typically used to decorate components</td>
</tr>
<tr>
<td class="name">ATK_ROLE_IMAGE</td>
<td class="value">26</td>
<td class="doc">An object whose primary purpose is to display an image</td>
</tr>
<tr>
<td class="name">ATK_ROLE_INTERNAL_FRAME</td>
<td class="value">27</td>
<td class="doc">A frame-like object that is clipped by a desktop pane</td>
</tr>
<tr>
<td class="name">ATK_ROLE_LABEL</td>
<td class="value">28</td>
<td class="doc">An object used to present an icon or short string in an interface</td>
</tr>
<tr>
<td class="name">ATK_ROLE_LAYERED_PANE</td>
<td class="value">29</td>
<td class="doc">A specialized pane that allows its children to be drawn in layers, providing a form of stacking order</td>
</tr>
<tr>
<td class="name">ATK_ROLE_LIST</td>
<td class="value">30</td>
<td class="doc">An object that presents a list of objects to the user and allows the user to select one or more of them</td>
</tr>
<tr>
<td class="name">ATK_ROLE_LIST_ITEM</td>
<td class="value">31</td>
<td class="doc">An object that represents an element of a list</td>
</tr>
<tr>
<td class="name">ATK_ROLE_MENU</td>
<td class="value">32</td>
<td class="doc">An object usually found inside a menu bar that contains a list of actions the user can choose from</td>
</tr>
<tr>
<td class="name">ATK_ROLE_MENU_BAR</td>
<td class="value">33</td>
<td class="doc">An object usually drawn at the top of the primary dialog box of an application that contains a list of menus the user can choose from</td>
</tr>
<tr>
<td class="name">ATK_ROLE_MENU_ITEM</td>
<td class="value">34</td>
<td class="doc">An object usually contained in a menu that presents an action the user can choose</td>
</tr>
<tr>
<td class="name">ATK_ROLE_OPTION_PANE</td>
<td class="value">35</td>
<td class="doc">A specialized pane whose primary use is inside a DIALOG</td>
</tr>
<tr>
<td class="name">ATK_ROLE_PAGE_TAB</td>
<td class="value">36</td>
<td class="doc">An object that is a child of a page tab list</td>
</tr>
<tr>
<td class="name">ATK_ROLE_PAGE_TAB_LIST</td>
<td class="value">37</td>
<td class="doc">An object that presents a series of panels (or page tabs), one at a time, through some mechanism provided by the object</td>
</tr>
<tr>
<td class="name">ATK_ROLE_PANEL</td>
<td class="value">38</td>
<td class="doc">A generic container that is often used to group objects</td>
</tr>
<tr>
<td class="name">ATK_ROLE_PASSWORD_TEXT</td>
<td class="value">39</td>
<td class="doc">A text object uses for passwords, or other places where the text content is not shown visibly to the user</td>
</tr>
<tr>
<td class="name">ATK_ROLE_POPUP_MENU</td>
<td class="value">40</td>
<td class="doc">A temporary window that is usually used to offer the user a list of choices, and then hides when the user selects one of those choices</td>
</tr>
<tr>
<td class="name">ATK_ROLE_PROGRESS_BAR</td>
<td class="value">41</td>
<td class="doc">An object used to indicate how much of a task has been completed</td>
</tr>
<tr>
<td class="name">ATK_ROLE_PUSH_BUTTON</td>
<td class="value">42</td>
<td class="doc">An object the user can manipulate to tell the application to do something</td>
</tr>
<tr>
<td class="name">ATK_ROLE_RADIO_BUTTON</td>
<td class="value">43</td>
<td class="doc">A specialized check box that will cause other radio buttons in the same group to become unchecked when this one is checked</td>
</tr>
<tr>
<td class="name">ATK_ROLE_RADIO_MENU_ITEM</td>
<td class="value">44</td>
<td class="doc">A check menu item which belongs to a group. At each instant exactly one of the radio menu items from a group is selected</td>
</tr>
<tr>
<td class="name">ATK_ROLE_ROOT_PANE</td>
<td class="value">45</td>
<td class="doc">A specialized pane that has a glass pane and a layered pane as its children</td>
</tr>
<tr>
<td class="name">ATK_ROLE_ROW_HEADER</td>
<td class="value">46</td>
<td class="doc">The header for a row of data</td>
</tr>
<tr>
<td class="name">ATK_ROLE_SCROLL_BAR</td>
<td class="value">47</td>
<td class="doc">An object usually used to allow a user to incrementally view a large amount of data.</td>
</tr>
<tr>
<td class="name">ATK_ROLE_SCROLL_PANE</td>
<td class="value">48</td>
<td class="doc">An object that allows a user to incrementally view a large amount of information</td>
</tr>
<tr>
<td class="name">ATK_ROLE_SEPARATOR</td>
<td class="value">49</td>
<td class="doc">An object usually contained in a menu to provide a visible and logical separation of the contents in a menu</td>
</tr>
<tr>
<td class="name">ATK_ROLE_SLIDER</td>
<td class="value">50</td>
<td class="doc">An object that allows the user to select from a bounded range</td>
</tr>
<tr>
<td class="name">ATK_ROLE_SPLIT_PANE</td>
<td class="value">51</td>
<td class="doc">A specialized panel that presents two other panels at the same time</td>
</tr>
<tr>
<td class="name">ATK_ROLE_SPIN_BUTTON</td>
<td class="value">52</td>
<td class="doc">An object used to get an integer or floating point number from the user</td>
</tr>
<tr>
<td class="name">ATK_ROLE_STATUSBAR</td>
<td class="value">53</td>
<td class="doc">An object which reports messages of minor importance to the user</td>
</tr>
<tr>
<td class="name">ATK_ROLE_TABLE</td>
<td class="value">54</td>
<td class="doc">An object used to represent information in terms of rows and columns</td>
</tr>
<tr>
<td class="name">ATK_ROLE_TABLE_CELL</td>
<td class="value">55</td>
<td class="doc">A cell in a table</td>
</tr>
<tr>
<td class="name">ATK_ROLE_TABLE_COLUMN_HEADER</td>
<td class="value">56</td>
<td class="doc">The header for a column of a table</td>
</tr>
<tr>
<td class="name">ATK_ROLE_TABLE_ROW_HEADER</td>
<td class="value">57</td>
<td class="doc">The header for a row of a table</td>
</tr>
<tr>
<td class="name">ATK_ROLE_TEAR_OFF_MENU_ITEM</td>
<td class="value">58</td>
<td class="doc">A menu item used to tear off and reattach its menu</td>
</tr>
<tr>
<td class="name">ATK_ROLE_TERMINAL</td>
<td class="value">59</td>
<td class="doc">An object that represents an accessible terminal.  @Since: ATK-0.6</td>
</tr>
<tr>
<td class="name">ATK_ROLE_TEXT</td>
<td class="value">60</td>
<td class="doc">An interactive widget that supports multiple lines of text and
optionally accepts user input, but whose purpose is not to solicit user input.
Thus ATK_ROLE_TEXT is appropriate for the text view in a plain text editor
but inappropriate for an input field in a dialog box or web form. For widgets
whose purpose is to solicit input from the user, see ATK_ROLE_ENTRY and
ATK_ROLE_PASSWORD_TEXT. For generic objects which display a brief amount of
textual information, see ATK_ROLE_STATIC.</td>
</tr>
<tr>
<td class="name">ATK_ROLE_TOGGLE_BUTTON</td>
<td class="value">61</td>
<td class="doc">A specialized push button that can be checked or unchecked, but does not provide a separate indicator for the current state</td>
</tr>
<tr>
<td class="name">ATK_ROLE_TOOL_BAR</td>
<td class="value">62</td>
<td class="doc">A bar or palette usually composed of push buttons or toggle buttons</td>
</tr>
<tr>
<td class="name">ATK_ROLE_TOOL_TIP</td>
<td class="value">63</td>
<td class="doc">An object that provides information about another object</td>
</tr>
<tr>
<td class="name">ATK_ROLE_TREE</td>
<td class="value">64</td>
<td class="doc">An object used to represent hierarchical information to the user</td>
</tr>
<tr>
<td class="name">ATK_ROLE_TREE_TABLE</td>
<td class="value">65</td>
<td class="doc">An object capable of expanding and collapsing rows as well as showing multiple columns of data.   @Since: ATK-0.7</td>
</tr>
<tr>
<td class="name">ATK_ROLE_UNKNOWN</td>
<td class="value">66</td>
<td class="doc">The object contains some Accessible information, but its role is not known</td>
</tr>
<tr>
<td class="name">ATK_ROLE_VIEWPORT</td>
<td class="value">67</td>
<td class="doc">An object usually used in a scroll pane</td>
</tr>
<tr>
<td class="name">ATK_ROLE_WINDOW</td>
<td class="value">68</td>
<td class="doc">A top level window with no title or border.</td>
</tr>
<tr>
<td class="name">ATK_ROLE_HEADER</td>
<td class="value">69</td>
<td class="doc">An object that serves as a document header. @Since: ATK-1.1.1</td>
</tr>
<tr>
<td class="name">ATK_ROLE_FOOTER</td>
<td class="value">70</td>
<td class="doc">An object that serves as a document footer.  @Since: ATK-1.1.1</td>
</tr>
<tr>
<td class="name">ATK_ROLE_PARAGRAPH</td>
<td class="value">71</td>
<td class="doc">An object which is contains a paragraph of text content.   @Since: ATK-1.1.1</td>
</tr>
<tr>
<td class="name">ATK_ROLE_RULER</td>
<td class="value">72</td>
<td class="doc">An object which describes margins and tab stops, etc. for text objects which it controls (should have CONTROLLER_FOR relation to such).   @Since: ATK-1.1.1</td>
</tr>
<tr>
<td class="name">ATK_ROLE_APPLICATION</td>
<td class="value">73</td>
<td class="doc">The object is an application object, which may contain @ATK_ROLE_FRAME objects or other types of accessibles.  The root accessible of any application's ATK hierarchy should have ATK_ROLE_APPLICATION.   @Since: ATK-1.1.4</td>
</tr>
<tr>
<td class="name">ATK_ROLE_AUTOCOMPLETE</td>
<td class="value">74</td>
<td class="doc">The object is a dialog or list containing items for insertion into an entry widget, for instance a list of words for completion of a text entry.   @Since: ATK-1.3</td>
</tr>
<tr>
<td class="name">ATK_ROLE_EDITBAR</td>
<td class="value">75</td>
<td class="doc">The object is an editable text object in a toolbar.  @Since: ATK-1.5</td>
</tr>
<tr>
<td class="name">ATK_ROLE_EMBEDDED</td>
<td class="value">76</td>
<td class="doc">The object is an embedded container within a document or panel.  This role is a grouping "hint" indicating that the contained objects share a context.  @Since: ATK-1.7.2</td>
</tr>
<tr>
<td class="name">ATK_ROLE_ENTRY</td>
<td class="value">77</td>
<td class="doc">The object is a component whose textual content may be entered or modified by the user, provided @ATK_STATE_EDITABLE is present.   @Since: ATK-1.11</td>
</tr>
<tr>
<td class="name">ATK_ROLE_CHART</td>
<td class="value">78</td>
<td class="doc">The object is a graphical depiction of quantitative data. It may contain multiple subelements whose attributes and/or description may be queried to obtain both the quantitative data and information about how the data is being presented. The LABELLED_BY relation is particularly important in interpreting objects of this type, as is the accessible-description property.  @Since: ATK-1.11</td>
</tr>
<tr>
<td class="name">ATK_ROLE_CAPTION</td>
<td class="value">79</td>
<td class="doc">The object contains descriptive information, usually textual, about another user interface element such as a table, chart, or image.  @Since: ATK-1.11</td>
</tr>
<tr>
<td class="name">ATK_ROLE_DOCUMENT_FRAME</td>
<td class="value">80</td>
<td class="doc">The object is a visual frame or container which contains a view of document content. Document frames may occur within another Document instance, in which case the second document may be said to be embedded in the containing instance. HTML frames are often ROLE_DOCUMENT_FRAME. Either this object, or a singleton descendant, should implement the Document interface.  @Since: ATK-1.11</td>
</tr>
<tr>
<td class="name">ATK_ROLE_HEADING</td>
<td class="value">81</td>
<td class="doc">The object serves as a heading for content which follows it in a document. The 'heading level' of the heading, if availabe, may be obtained by querying the object's attributes.</td>
</tr>
<tr>
<td class="name">ATK_ROLE_PAGE</td>
<td class="value">82</td>
<td class="doc">The object is a containing instance which encapsulates a page of information. @ATK_ROLE_PAGE is used in documents and content which support a paginated navigation model.  @Since: ATK-1.11</td>
</tr>
<tr>
<td class="name">ATK_ROLE_SECTION</td>
<td class="value">83</td>
<td class="doc">The object is a containing instance of document content which constitutes a particular 'logical' section of the document. The type of content within a section, and the nature of the section division itself, may be obtained by querying the object's attributes. Sections may be nested. @Since: ATK-1.11</td>
</tr>
<tr>
<td class="name">ATK_ROLE_REDUNDANT_OBJECT</td>
<td class="value">84</td>
<td class="doc">The object is redundant with another object in the hierarchy, and is exposed for purely technical reasons.  Objects of this role should normally be ignored by clients. @Since: ATK-1.11</td>
</tr>
<tr>
<td class="name">ATK_ROLE_FORM</td>
<td class="value">85</td>
<td class="doc">The object is a container for form controls, for instance as part of a
web form or user-input form within a document.  This role is primarily a tag/convenience for
clients when navigating complex documents, it is not expected that ordinary GUI containers will
always have ATK_ROLE_FORM. @Since: ATK-1.12.0</td>
</tr>
<tr>
<td class="name">ATK_ROLE_LINK</td>
<td class="value">86</td>
<td class="doc">The object is a hypertext anchor, i.e. a "link" in a
hypertext document.  Such objects are distinct from 'inline'
content which may also use the Hypertext/Hyperlink interfaces
to indicate the range/location within a text object where
an inline or embedded object lies.  @Since: ATK-1.12.1</td>
</tr>
<tr>
<td class="name">ATK_ROLE_INPUT_METHOD_WINDOW</td>
<td class="value">87</td>
<td class="doc">The object is a window or similar viewport
which is used to allow composition or input of a 'complex character',
in other words it is an "input method window." @Since: ATK-1.12.1</td>
</tr>
<tr>
<td class="name">ATK_ROLE_TABLE_ROW</td>
<td class="value">88</td>
<td class="doc">A row in a table.  @Since: ATK-2.1.0</td>
</tr>
<tr>
<td class="name">ATK_ROLE_TREE_ITEM</td>
<td class="value">89</td>
<td class="doc">An object that represents an element of a tree.  @Since: ATK-2.1.0</td>
</tr>
<tr>
<td class="name">ATK_ROLE_DOCUMENT_SPREADSHEET</td>
<td class="value">90</td>
<td class="doc">A document frame which contains a spreadsheet.  @Since: ATK-2.1.0</td>
</tr>
<tr>
<td class="name">ATK_ROLE_DOCUMENT_PRESENTATION</td>
<td class="value">91</td>
<td class="doc">A document frame which contains a presentation or slide content.  @Since: ATK-2.1.0</td>
</tr>
<tr>
<td class="name">ATK_ROLE_DOCUMENT_TEXT</td>
<td class="value">92</td>
<td class="doc">A document frame which contains textual content, such as found in a word processing application.  @Since: ATK-2.1.0</td>
</tr>
<tr>
<td class="name">ATK_ROLE_DOCUMENT_WEB</td>
<td class="value">93</td>
<td class="doc">A document frame which contains HTML or other markup suitable for display in a web browser.  @Since: ATK-2.1.0</td>
</tr>
<tr>
<td class="name">ATK_ROLE_DOCUMENT_EMAIL</td>
<td class="value">94</td>
<td class="doc">A document frame which contains email content to be displayed or composed either in plain text or HTML.  @Since: ATK-2.1.0</td>
</tr>
<tr>
<td class="name">ATK_ROLE_COMMENT</td>
<td class="value">95</td>
<td class="doc">An object found within a document and designed to present a comment, note, or other annotation. In some cases, this object might not be visible until activated.  @Since: ATK-2.1.0</td>
</tr>
<tr>
<td class="name">ATK_ROLE_LIST_BOX</td>
<td class="value">96</td>
<td class="doc">A non-collapsible list of choices the user can select from. @Since: ATK-2.1.0</td>
</tr>
<tr>
<td class="name">ATK_ROLE_GROUPING</td>
<td class="value">97</td>
<td class="doc">A group of related widgets. This group typically has a label. @Since: ATK-2.1.0</td>
</tr>
<tr>
<td class="name">ATK_ROLE_IMAGE_MAP</td>
<td class="value">98</td>
<td class="doc">An image map object. Usually a graphic with multiple hotspots, where each hotspot can be activated resulting in the loading of another document or section of a document. @Since: ATK-2.1.0</td>
</tr>
<tr>
<td class="name">ATK_ROLE_NOTIFICATION</td>
<td class="value">99</td>
<td class="doc">A transitory object designed to present a message to the user, typically at the desktop level rather than inside a particular application.  @Since: ATK-2.1.0</td>
</tr>
<tr>
<td class="name">ATK_ROLE_INFO_BAR</td>
<td class="value">100</td>
<td class="doc">An object designed to present a message to the user within an existing window. @Since: ATK-2.1.0</td>
</tr>
<tr>
<td class="name">ATK_ROLE_LEVEL_BAR</td>
<td class="value">101</td>
<td class="doc">A bar that serves as a level indicator to, for instance, show the strength of a password or the state of a battery.  @Since: ATK-2.7.3</td>
</tr>
<tr>
<td class="name">ATK_ROLE_TITLE_BAR</td>
<td class="value">102</td>
<td class="doc">A bar that serves as the title of a window or a
dialog. @Since: ATK-2.12</td>
</tr>
<tr>
<td class="name">ATK_ROLE_BLOCK_QUOTE</td>
<td class="value">103</td>
<td class="doc">An object which contains a text section
that is quoted from another source. @Since: ATK-2.12</td>
</tr>
<tr>
<td class="name">ATK_ROLE_AUDIO</td>
<td class="value">104</td>
<td class="doc">An object which represents an audio element. @Since: ATK-2.12</td>
</tr>
<tr>
<td class="name">ATK_ROLE_VIDEO</td>
<td class="value">105</td>
<td class="doc">An object which represents a video element. @Since: ATK-2.12</td>
</tr>
<tr>
<td class="name">ATK_ROLE_DEFINITION</td>
<td class="value">106</td>
<td class="doc">A definition of a term or concept. @Since: ATK-2.12</td>
</tr>
<tr>
<td class="name">ATK_ROLE_ARTICLE</td>
<td class="value">107</td>
<td class="doc">A section of a page that consists of a
composition that forms an independent part of a document, page, or
site. Examples: A blog entry, a news story, a forum post. @Since:
ATK-2.12</td>
</tr>
<tr>
<td class="name">ATK_ROLE_LANDMARK</td>
<td class="value">108</td>
<td class="doc">A region of a web page intended as a
navigational landmark. This is designed to allow Assistive
Technologies to provide quick navigation among key regions within a
document. @Since: ATK-2.12</td>
</tr>
<tr>
<td class="name">ATK_ROLE_LOG</td>
<td class="value">109</td>
<td class="doc">A text widget or container holding log content, such
as chat history and error logs. In this role there is a
relationship between the arrival of new items in the log and the
reading order. The log contains a meaningful sequence and new
information is added only to the end of the log, not at arbitrary
points. @Since: ATK-2.12</td>
</tr>
<tr>
<td class="name">ATK_ROLE_MARQUEE</td>
<td class="value">110</td>
<td class="doc">A container where non-essential information
changes frequently. Common usages of marquee include stock tickers
and ad banners. The primary difference between a marquee and a log
is that logs usually have a meaningful order or sequence of
important content changes. @Since: ATK-2.12</td>
</tr>
<tr>
<td class="name">ATK_ROLE_MATH</td>
<td class="value">111</td>
<td class="doc">A text widget or container that holds a mathematical
expression. @Since: ATK-2.12</td>
</tr>
<tr>
<td class="name">ATK_ROLE_RATING</td>
<td class="value">112</td>
<td class="doc">A widget whose purpose is to display a rating,
such as the number of stars associated with a song in a media
player. Objects of this role should also implement
AtkValue. @Since: ATK-2.12</td>
</tr>
<tr>
<td class="name">ATK_ROLE_TIMER</td>
<td class="value">113</td>
<td class="doc">An object containing a numerical counter which
indicates an amount of elapsed time from a start point, or the time
remaining until an end point. @Since: ATK-2.12</td>
</tr>
<tr>
<td class="name">ATK_ROLE_DESCRIPTION_LIST</td>
<td class="value">114</td>
<td class="doc">An object that represents a list of
term-value groups. A term-value group represents a individual
description and consist of one or more names
(ATK_ROLE_DESCRIPTION_TERM) followed by one or more values
(ATK_ROLE_DESCRIPTION_VALUE). For each list, there should not be
more than one group with the same term name. @Since: ATK-2.12</td>
</tr>
<tr>
<td class="name">ATK_ROLE_DESCRIPTION_TERM</td>
<td class="value">115</td>
<td class="doc">An object that represents a term or phrase
with a corresponding definition. @Since: ATK-2.12</td>
</tr>
<tr>
<td class="name">ATK_ROLE_DESCRIPTION_VALUE</td>
<td class="value">116</td>
<td class="doc">An object that represents the
description, definition or value of a term. @Since: ATK-2.12</td>
</tr>
<tr>
<td class="name">ATK_ROLE_STATIC</td>
<td class="value">117</td>
<td class="doc">A generic non-container object whose purpose is to display a
brief amount of information to the user and whose role is known by the
implementor but lacks semantic value for the user. Examples in which
ATK_ROLE_STATIC is appropriate include the message displayed in a message box
and an image used as an alternative means to display text. ATK_ROLE_STATIC
should not be applied to widgets which are traditionally interactive, objects
which display a significant amount of content, or any object which has an
accessible relation pointing to another object. Implementors should expose the
displayed information through the accessible name of the object. If doing so seems
inappropriate, it may indicate that a different role should be used. For
labels which describe another widget, see ATK_ROLE_LABEL. For text views, see
ATK_ROLE_TEXT. For generic containers, see ATK_ROLE_PANEL. For objects whose
role is not known by the implementor, see ATK_ROLE_UNKNOWN. @Since: ATK-2.16.</td>
</tr>
<tr>
<td class="name">ATK_ROLE_MATH_FRACTION</td>
<td class="value">118</td>
<td class="doc">An object that represents a mathematical fraction.</td>
</tr>
<tr>
<td class="name">ATK_ROLE_MATH_ROOT</td>
<td class="value">119</td>
<td class="doc">An object that represents a mathematical expression
displayed with a radical. @Since: ATK-2.16.</td>
</tr>
<tr>
<td class="name">ATK_ROLE_SUBSCRIPT</td>
<td class="value">120</td>
<td class="doc">An object that contains text that is displayed as a
subscript. @Since: ATK-2.16.</td>
</tr>
<tr>
<td class="name">ATK_ROLE_SUPERSCRIPT</td>
<td class="value">121</td>
<td class="doc">An object that contains text that is displayed as a
superscript. @Since: ATK-2.16.</td>
</tr>
<tr>
<td class="name">ATK_ROLE_FOOTNOTE</td>
<td class="value">122</td>
<td class="doc">An object that contains the text of a footnote. @Since: ATK-2.26.</td>
</tr>
<tr>
<td class="name">ATK_ROLE_LAST_DEFINED</td>
<td class="value">123</td>
<td class="doc">not a valid role, used for finding end of the enumeration</td>
</tr>
</table>
</div>
<p class="api-heading">StateType</p>
<p class="api-doc">The possible types of states of an object</p>
<div class="api-notes">
  <p class="api-ctype">AtkStateType</p>
<table>
<tr>
<td class="name">ATK_STATE_INVALID</td>
<td class="value">0</td>
<td class="doc">Indicates an invalid state - probably an error condition.</td>
</tr>
<tr>
<td class="name">ATK_STATE_ACTIVE</td>
<td class="value">1</td>
<td class="doc">Indicates a window is currently the active window, or an object is the active subelement within a container or table. ATK_STATE_ACTIVE should not be used for objects which have ATK_STATE_FOCUSABLE or ATK_STATE_SELECTABLE: Those objects should use ATK_STATE_FOCUSED and ATK_STATE_SELECTED respectively. ATK_STATE_ACTIVE is a means to indicate that an object which is not focusable and not selectable is the currently-active item within its parent container.</td>
</tr>
<tr>
<td class="name">ATK_STATE_ARMED</td>
<td class="value">2</td>
<td class="doc">Indicates that the object is 'armed', i.e. will be activated by if a pointer button-release event occurs within its bounds.  Buttons often enter this state when a pointer click occurs within their bounds, as a precursor to activation. ATK_STATE_ARMED has been deprecated since ATK-2.16 and should not be used in newly-written code.</td>
</tr>
<tr>
<td class="name">ATK_STATE_BUSY</td>
<td class="value">3</td>
<td class="doc">Indicates the current object is busy, i.e. onscreen representation is in the process of changing, or the object is temporarily unavailable for interaction due to activity already in progress.  This state may be used by implementors of Document to indicate that content loading is underway.  It also may indicate other 'pending' conditions; clients may wish to interrogate this object when the ATK_STATE_BUSY flag is removed.</td>
</tr>
<tr>
<td class="name">ATK_STATE_CHECKED</td>
<td class="value">4</td>
<td class="doc">Indicates this object is currently checked, for instance a checkbox is 'non-empty'.</td>
</tr>
<tr>
<td class="name">ATK_STATE_DEFUNCT</td>
<td class="value">5</td>
<td class="doc">Indicates that this object no longer has a valid backing widget (for instance, if its peer object has been destroyed)</td>
</tr>
<tr>
<td class="name">ATK_STATE_EDITABLE</td>
<td class="value">6</td>
<td class="doc">Indicates that this object can contain text, and that the
user can change the textual contents of this object by editing those contents
directly. For an object which is expected to be editable due to its type, but
which cannot be edited due to the application or platform preventing the user
from doing so, that object's #AtkStateSet should lack ATK_STATE_EDITABLE and
should contain ATK_STATE_READ_ONLY.</td>
</tr>
<tr>
<td class="name">ATK_STATE_ENABLED</td>
<td class="value">7</td>
<td class="doc">Indicates that this object is enabled, i.e. that it currently reflects some application state. Objects that are "greyed out" may lack this state, and may lack the STATE_SENSITIVE if direct user interaction cannot cause them to acquire STATE_ENABLED. See also: ATK_STATE_SENSITIVE</td>
</tr>
<tr>
<td class="name">ATK_STATE_EXPANDABLE</td>
<td class="value">8</td>
<td class="doc">Indicates this object allows progressive disclosure of its children</td>
</tr>
<tr>
<td class="name">ATK_STATE_EXPANDED</td>
<td class="value">9</td>
<td class="doc">Indicates this object its expanded - see ATK_STATE_EXPANDABLE above</td>
</tr>
<tr>
<td class="name">ATK_STATE_FOCUSABLE</td>
<td class="value">10</td>
<td class="doc">Indicates this object can accept keyboard focus, which means all events resulting from typing on the keyboard will normally be passed to it when it has focus</td>
</tr>
<tr>
<td class="name">ATK_STATE_FOCUSED</td>
<td class="value">11</td>
<td class="doc">Indicates this object currently has the keyboard focus</td>
</tr>
<tr>
<td class="name">ATK_STATE_HORIZONTAL</td>
<td class="value">12</td>
<td class="doc">Indicates the orientation of this object is horizontal; used, for instance, by objects of ATK_ROLE_SCROLL_BAR.  For objects where vertical/horizontal orientation is especially meaningful.</td>
</tr>
<tr>
<td class="name">ATK_STATE_ICONIFIED</td>
<td class="value">13</td>
<td class="doc">Indicates this object is minimized and is represented only by an icon</td>
</tr>
<tr>
<td class="name">ATK_STATE_MODAL</td>
<td class="value">14</td>
<td class="doc">Indicates something must be done with this object before the user can interact with an object in a different window</td>
</tr>
<tr>
<td class="name">ATK_STATE_MULTI_LINE</td>
<td class="value">15</td>
<td class="doc">Indicates this (text) object can contain multiple lines of text</td>
</tr>
<tr>
<td class="name">ATK_STATE_MULTISELECTABLE</td>
<td class="value">16</td>
<td class="doc">Indicates this object allows more than one of its children to be selected at the same time, or in the case of text objects, that the object supports non-contiguous text selections.</td>
</tr>
<tr>
<td class="name">ATK_STATE_OPAQUE</td>
<td class="value">17</td>
<td class="doc">Indicates this object paints every pixel within its rectangular region.</td>
</tr>
<tr>
<td class="name">ATK_STATE_PRESSED</td>
<td class="value">18</td>
<td class="doc">Indicates this object is currently pressed.</td>
</tr>
<tr>
<td class="name">ATK_STATE_RESIZABLE</td>
<td class="value">19</td>
<td class="doc">Indicates the size of this object is not fixed</td>
</tr>
<tr>
<td class="name">ATK_STATE_SELECTABLE</td>
<td class="value">20</td>
<td class="doc">Indicates this object is the child of an object that allows its children to be selected and that this child is one of those children that can be selected</td>
</tr>
<tr>
<td class="name">ATK_STATE_SELECTED</td>
<td class="value">21</td>
<td class="doc">Indicates this object is the child of an object that allows its children to be selected and that this child is one of those children that has been selected</td>
</tr>
<tr>
<td class="name">ATK_STATE_SENSITIVE</td>
<td class="value">22</td>
<td class="doc">Indicates this object is sensitive, e.g. to user interaction.
STATE_SENSITIVE usually accompanies STATE_ENABLED for user-actionable controls,
but may be found in the absence of STATE_ENABLED if the current visible state of the
control is "disconnected" from the application state.  In such cases, direct user interaction
can often result in the object gaining STATE_SENSITIVE, for instance if a user makes
an explicit selection using an object whose current state is ambiguous or undefined.
@see STATE_ENABLED, STATE_INDETERMINATE.</td>
</tr>
<tr>
<td class="name">ATK_STATE_SHOWING</td>
<td class="value">23</td>
<td class="doc">Indicates this object, the object's parent, the object's parent's parent, and so on,
are all 'shown' to the end-user, i.e. subject to "exposure" if blocking or obscuring objects do not interpose
between this object and the top of the window stack.</td>
</tr>
<tr>
<td class="name">ATK_STATE_SINGLE_LINE</td>
<td class="value">24</td>
<td class="doc">Indicates this (text) object can contain only a single line of text</td>
</tr>
<tr>
<td class="name">ATK_STATE_STALE</td>
<td class="value">25</td>
<td class="doc">Indicates that the information returned for this object may no longer be
synchronized with the application state.  This is implied if the object has STATE_TRANSIENT,
and can also occur towards the end of the object peer's lifecycle. It can also be used to indicate that
the index associated with this object has changed since the user accessed the object (in lieu of
"index-in-parent-changed" events).</td>
</tr>
<tr>
<td class="name">ATK_STATE_TRANSIENT</td>
<td class="value">26</td>
<td class="doc">Indicates this object is transient, i.e. a snapshot which may not emit events when its
state changes.  Data from objects with ATK_STATE_TRANSIENT should not be cached, since there may be no
notification given when the cached data becomes obsolete.</td>
</tr>
<tr>
<td class="name">ATK_STATE_VERTICAL</td>
<td class="value">27</td>
<td class="doc">Indicates the orientation of this object is vertical</td>
</tr>
<tr>
<td class="name">ATK_STATE_VISIBLE</td>
<td class="value">28</td>
<td class="doc">Indicates this object is visible, e.g. has been explicitly marked for exposure to the user.</td>
</tr>
<tr>
<td class="name">ATK_STATE_MANAGES_DESCENDANTS</td>
<td class="value">29</td>
<td class="doc">Indicates that "active-descendant-changed" event
is sent when children become 'active' (i.e. are selected or navigated to onscreen).
Used to prevent need to enumerate all children in very large containers, like tables.
The presence of STATE_MANAGES_DESCENDANTS is an indication to the client.
that the children should not, and need not, be enumerated by the client.
Objects implementing this state are expected to provide relevant state
notifications to listening clients, for instance notifications of visibility
changes and activation of their contained child objects, without the client
having previously requested references to those children.</td>
</tr>
<tr>
<td class="name">ATK_STATE_INDETERMINATE</td>
<td class="value">30</td>
<td class="doc">Indicates that the value, or some other quantifiable
property, of this AtkObject cannot be fully determined. In the case of a large
data set in which the total number of items in that set is unknown (e.g. 1 of
999+), implementors should expose the currently-known set size (999) along
with this state. In the case of a check box, this state should be used to
indicate that the check box is a tri-state check box which is currently
neither checked nor unchecked.</td>
</tr>
<tr>
<td class="name">ATK_STATE_TRUNCATED</td>
<td class="value">31</td>
<td class="doc">Indicates that an object is truncated, e.g. a text value in a speradsheet cell.</td>
</tr>
<tr>
<td class="name">ATK_STATE_REQUIRED</td>
<td class="value">32</td>
<td class="doc">Indicates that explicit user interaction with an object is required by the user interface, e.g. a required field in a "web-form" interface.</td>
</tr>
<tr>
<td class="name">ATK_STATE_INVALID_ENTRY</td>
<td class="value">33</td>
<td class="doc">Indicates that the object has encountered an error condition due to failure of input validation. For instance, a form control may acquire this state in response to invalid or malformed user input.</td>
</tr>
<tr>
<td class="name">ATK_STATE_SUPPORTS_AUTOCOMPLETION</td>
<td class="value">34</td>
<td class="doc">Indicates that the object in question implements some form of ¨typeahead¨ or
pre-selection behavior whereby entering the first character of one or more sub-elements
causes those elements to scroll into view or become selected.  Subsequent character input
may narrow the selection further as long as one or more sub-elements match the string.
This state is normally only useful and encountered on objects that implement Selection.
In some cases the typeahead behavior may result in full or partial ¨completion¨ of
the data in the input field, in which case these input events may trigger text-changed
events from the AtkText interface.  This state supplants @ATK_ROLE_AUTOCOMPLETE.</td>
</tr>
<tr>
<td class="name">ATK_STATE_SELECTABLE_TEXT</td>
<td class="value">35</td>
<td class="doc">Indicates that the object in question supports text selection. It should only be exposed on objects which implement the Text interface, in order to distinguish this state from @ATK_STATE_SELECTABLE, which infers that the object in question is a selectable child of an object which implements Selection. While similar, text selection and subelement selection are distinct operations.</td>
</tr>
<tr>
<td class="name">ATK_STATE_DEFAULT</td>
<td class="value">36</td>
<td class="doc">Indicates that the object is the "default" active component, i.e. the object which is activated by an end-user press of the "Enter" or "Return" key.  Typically a "close" or "submit" button.</td>
</tr>
<tr>
<td class="name">ATK_STATE_ANIMATED</td>
<td class="value">37</td>
<td class="doc">Indicates that the object changes its appearance dynamically as an inherent part of its presentation.  This state may come and go if an object is only temporarily animated on the way to a 'final' onscreen presentation.
@note some applications, notably content viewers, may not be able to detect
all kinds of animated content.  Therefore the absence of this state should not
be taken as definitive evidence that the object's visual representation is
static; this state is advisory.</td>
</tr>
<tr>
<td class="name">ATK_STATE_VISITED</td>
<td class="value">38</td>
<td class="doc">Indicates that the object (typically a hyperlink) has already been 'activated', and/or its backing data has already been downloaded, rendered, or otherwise "visited".</td>
</tr>
<tr>
<td class="name">ATK_STATE_CHECKABLE</td>
<td class="value">39</td>
<td class="doc">Indicates this object has the potential to be
 checked, such as a checkbox or toggle-able table cell. @Since:
 ATK-2.12</td>
</tr>
<tr>
<td class="name">ATK_STATE_HAS_POPUP</td>
<td class="value">40</td>
<td class="doc">Indicates that the object has a popup context
menu or sub-level menu which may or may not be showing. This means
that activation renders conditional content.  Note that ordinary
tooltips are not considered popups in this context. @Since: ATK-2.12</td>
</tr>
<tr>
<td class="name">ATK_STATE_HAS_TOOLTIP</td>
<td class="value">41</td>
<td class="doc">Indicates this object has a tooltip. @Since: ATK-2.16</td>
</tr>
<tr>
<td class="name">ATK_STATE_READ_ONLY</td>
<td class="value">42</td>
<td class="doc">Indicates that a widget which is ENABLED and SENSITIVE
has a value which can be read, but not modified, by the user. Note that this
state should only be applied to widget types whose value is normally directly
user modifiable, such as check boxes, radio buttons, spin buttons, text input
fields, and combo boxes, as a means to convey that the expected interaction
with that widget is not possible. When the expected interaction with a
widget does not include modification by the user, as is the case with
labels and containers, ATK_STATE_READ_ONLY should not be applied. See also
ATK_STATE_EDITABLE. @Since: ATK-2-16</td>
</tr>
<tr>
<td class="name">ATK_STATE_LAST_DEFINED</td>
<td class="value">43</td>
<td class="doc">Not a valid state, used for finding end of enumeration</td>
</tr>
</table>
</div>
<p class="api-heading">TextAttribute</p>
<p class="api-doc">Describes the text attributes supported</p>
<div class="api-notes">
  <p class="api-ctype">AtkTextAttribute</p>
<table>
<tr>
<td class="name">ATK_TEXT_ATTR_INVALID</td>
<td class="value">0</td>
<td class="doc">Invalid attribute, like bad spelling or grammar.</td>
</tr>
<tr>
<td class="name">ATK_TEXT_ATTR_LEFT_MARGIN</td>
<td class="value">1</td>
<td class="doc">The pixel width of the left margin</td>
</tr>
<tr>
<td class="name">ATK_TEXT_ATTR_RIGHT_MARGIN</td>
<td class="value">2</td>
<td class="doc">The pixel width of the right margin</td>
</tr>
<tr>
<td class="name">ATK_TEXT_ATTR_INDENT</td>
<td class="value">3</td>
<td class="doc">The number of pixels that the text is indented</td>
</tr>
<tr>
<td class="name">ATK_TEXT_ATTR_INVISIBLE</td>
<td class="value">4</td>
<td class="doc">Either "true" or "false" indicating whether text is visible or not</td>
</tr>
<tr>
<td class="name">ATK_TEXT_ATTR_EDITABLE</td>
<td class="value">5</td>
<td class="doc">Either "true" or "false" indicating whether text is editable or not</td>
</tr>
<tr>
<td class="name">ATK_TEXT_ATTR_PIXELS_ABOVE_LINES</td>
<td class="value">6</td>
<td class="doc">Pixels of blank space to leave above each newline-terminated line.</td>
</tr>
<tr>
<td class="name">ATK_TEXT_ATTR_PIXELS_BELOW_LINES</td>
<td class="value">7</td>
<td class="doc">Pixels of blank space to leave below each newline-terminated line.</td>
</tr>
<tr>
<td class="name">ATK_TEXT_ATTR_PIXELS_INSIDE_WRAP</td>
<td class="value">8</td>
<td class="doc">Pixels of blank space to leave between wrapped lines inside the same newline-terminated line (paragraph).</td>
</tr>
<tr>
<td class="name">ATK_TEXT_ATTR_BG_FULL_HEIGHT</td>
<td class="value">9</td>
<td class="doc">"true" or "false" whether to make the background color for each character the height of the highest font used on the current line, or the height of the font used for the current character.</td>
</tr>
<tr>
<td class="name">ATK_TEXT_ATTR_RISE</td>
<td class="value">10</td>
<td class="doc">Number of pixels that the characters are risen above the baseline</td>
</tr>
<tr>
<td class="name">ATK_TEXT_ATTR_UNDERLINE</td>
<td class="value">11</td>
<td class="doc">"none", "single", "double", "low", or "error"</td>
</tr>
<tr>
<td class="name">ATK_TEXT_ATTR_STRIKETHROUGH</td>
<td class="value">12</td>
<td class="doc">"true" or "false" whether the text is strikethrough</td>
</tr>
<tr>
<td class="name">ATK_TEXT_ATTR_SIZE</td>
<td class="value">13</td>
<td class="doc">The size of the characters in points. eg: 10</td>
</tr>
<tr>
<td class="name">ATK_TEXT_ATTR_SCALE</td>
<td class="value">14</td>
<td class="doc">The scale of the characters. The value is a string representation of a double</td>
</tr>
<tr>
<td class="name">ATK_TEXT_ATTR_WEIGHT</td>
<td class="value">15</td>
<td class="doc">The weight of the characters.</td>
</tr>
<tr>
<td class="name">ATK_TEXT_ATTR_LANGUAGE</td>
<td class="value">16</td>
<td class="doc">The language used</td>
</tr>
<tr>
<td class="name">ATK_TEXT_ATTR_FAMILY_NAME</td>
<td class="value">17</td>
<td class="doc">The font family name</td>
</tr>
<tr>
<td class="name">ATK_TEXT_ATTR_BG_COLOR</td>
<td class="value">18</td>
<td class="doc">The background color. The value is an RGB value of the format "%u,%u,%u"</td>
</tr>
<tr>
<td class="name">ATK_TEXT_ATTR_FG_COLOR</td>
<td class="value">19</td>
<td class="doc">The foreground color. The value is an RGB value of the format "%u,%u,%u"</td>
</tr>
<tr>
<td class="name">ATK_TEXT_ATTR_BG_STIPPLE</td>
<td class="value">20</td>
<td class="doc">"true" if a #GdkBitmap is set for stippling the background color.</td>
</tr>
<tr>
<td class="name">ATK_TEXT_ATTR_FG_STIPPLE</td>
<td class="value">21</td>
<td class="doc">"true" if a #GdkBitmap is set for stippling the foreground color.</td>
</tr>
<tr>
<td class="name">ATK_TEXT_ATTR_WRAP_MODE</td>
<td class="value">22</td>
<td class="doc">The wrap mode of the text, if any. Values are "none", "char", "word", or "word_char".</td>
</tr>
<tr>
<td class="name">ATK_TEXT_ATTR_DIRECTION</td>
<td class="value">23</td>
<td class="doc">The direction of the text, if set. Values are "none", "ltr" or "rtl"</td>
</tr>
<tr>
<td class="name">ATK_TEXT_ATTR_JUSTIFICATION</td>
<td class="value">24</td>
<td class="doc">The justification of the text, if set. Values are "left", "right", "center" or "fill"</td>
</tr>
<tr>
<td class="name">ATK_TEXT_ATTR_STRETCH</td>
<td class="value">25</td>
<td class="doc">The stretch of the text, if set. Values are "ultra_condensed", "extra_condensed", "condensed", "semi_condensed", "normal", "semi_expanded", "expanded", "extra_expanded" or "ultra_expanded"</td>
</tr>
<tr>
<td class="name">ATK_TEXT_ATTR_VARIANT</td>
<td class="value">26</td>
<td class="doc">The capitalization variant of the text, if set. Values are "normal" or "small_caps"</td>
</tr>
<tr>
<td class="name">ATK_TEXT_ATTR_STYLE</td>
<td class="value">27</td>
<td class="doc">The slant style of the text, if set. Values are "normal", "oblique" or "italic"</td>
</tr>
<tr>
<td class="name">ATK_TEXT_ATTR_LAST_DEFINED</td>
<td class="value">28</td>
<td class="doc">not a valid text attribute, used for finding end of enumeration</td>
</tr>
</table>
</div>
<p class="api-heading">TextBoundary</p>
<p class="api-doc">Text boundary types used for specifying boundaries for regions of text.
This enumeration is deprecated since 2.9.4 and should not be used. Use
AtkTextGranularity with #atk_text_get_string_at_offset instead.</p>
<div class="api-notes">
  <p class="api-ctype">AtkTextBoundary</p>
<table>
<tr>
<td class="name">ATK_TEXT_BOUNDARY_CHAR</td>
<td class="value">0</td>
<td class="doc">Boundary is the boundary between characters
(including non-printing characters)</td>
</tr>
<tr>
<td class="name">ATK_TEXT_BOUNDARY_WORD_START</td>
<td class="value">1</td>
<td class="doc">Boundary is the start (i.e. first character) of a word.</td>
</tr>
<tr>
<td class="name">ATK_TEXT_BOUNDARY_WORD_END</td>
<td class="value">2</td>
<td class="doc">Boundary is the end (i.e. last
character) of a word.</td>
</tr>
<tr>
<td class="name">ATK_TEXT_BOUNDARY_SENTENCE_START</td>
<td class="value">3</td>
<td class="doc">Boundary is the first character in a sentence.</td>
</tr>
<tr>
<td class="name">ATK_TEXT_BOUNDARY_SENTENCE_END</td>
<td class="value">4</td>
<td class="doc">Boundary is the last (terminal)
character in a sentence; in languages which use "sentence stop"
punctuation such as English, the boundary is thus the '.', '?', or
similar terminal punctuation character.</td>
</tr>
<tr>
<td class="name">ATK_TEXT_BOUNDARY_LINE_START</td>
<td class="value">5</td>
<td class="doc">Boundary is the initial character of the content or a
character immediately following a newline, linefeed, or return character.</td>
</tr>
<tr>
<td class="name">ATK_TEXT_BOUNDARY_LINE_END</td>
<td class="value">6</td>
<td class="doc">Boundary is the linefeed, or return
character.</td>
</tr>
</table>
</div>
<p class="api-heading">TextClipType</p>
<p class="api-doc">Describes the type of clipping required.</p>
<div class="api-notes">
  <p class="api-ctype">AtkTextClipType</p>
<table>
<tr>
<td class="name">ATK_TEXT_CLIP_NONE</td>
<td class="value">0</td>
<td class="doc">No clipping to be done</td>
</tr>
<tr>
<td class="name">ATK_TEXT_CLIP_MIN</td>
<td class="value">1</td>
<td class="doc">Text clipped by min coordinate is omitted</td>
</tr>
<tr>
<td class="name">ATK_TEXT_CLIP_MAX</td>
<td class="value">2</td>
<td class="doc">Text clipped by max coordinate is omitted</td>
</tr>
<tr>
<td class="name">ATK_TEXT_CLIP_BOTH</td>
<td class="value">3</td>
<td class="doc">Only text fully within mix/max bound is retained</td>
</tr>
</table>
</div>
<p class="api-heading">TextGranularity</p>
<p class="api-doc">Text granularity types used for specifying the granularity of the region of
text we are interested in.</p>
<div class="api-notes">
  <p class="api-ctype">AtkTextGranularity</p>
<table>
<tr>
<td class="name">ATK_TEXT_GRANULARITY_CHAR</td>
<td class="value">0</td>
<td class="doc">Granularity is defined by the boundaries between characters
(including non-printing characters)</td>
</tr>
<tr>
<td class="name">ATK_TEXT_GRANULARITY_WORD</td>
<td class="value">1</td>
<td class="doc">Granularity is defined by the boundaries of a word,
starting at the beginning of the current word and finishing at the beginning of
the following one, if present.</td>
</tr>
<tr>
<td class="name">ATK_TEXT_GRANULARITY_SENTENCE</td>
<td class="value">2</td>
<td class="doc">Granularity is defined by the boundaries of a sentence,
starting at the beginning of the current sentence and finishing at the beginning of
the following one, if present.</td>
</tr>
<tr>
<td class="name">ATK_TEXT_GRANULARITY_LINE</td>
<td class="value">3</td>
<td class="doc">Granularity is defined by the boundaries of a line,
starting at the beginning of the current line and finishing at the beginning of
the following one, if present.</td>
</tr>
<tr>
<td class="name">ATK_TEXT_GRANULARITY_PARAGRAPH</td>
<td class="value">4</td>
<td class="doc">Granularity is defined by the boundaries of a paragraph,
starting at the beginning of the current paragraph and finishing at the beginning of
the following one, if present.</td>
</tr>
</table>
</div>
<p class="api-heading">ValueType</p>
<p class="api-doc">Default types for a given value. Those are defined in order to
easily get localized strings to describe a given value or a given
subrange, using atk_value_type_get_localized_name().</p>
<div class="api-notes">
  <p class="api-ctype">AtkValueType</p>
<table>
<tr>
<td class="name">ATK_VALUE_VERY_WEAK</td>
<td class="value">0</td>
</tr>
<tr>
<td class="name">ATK_VALUE_WEAK</td>
<td class="value">1</td>
</tr>
<tr>
<td class="name">ATK_VALUE_ACCEPTABLE</td>
<td class="value">2</td>
</tr>
<tr>
<td class="name">ATK_VALUE_STRONG</td>
<td class="value">3</td>
</tr>
<tr>
<td class="name">ATK_VALUE_VERY_STRONG</td>
<td class="value">4</td>
</tr>
<tr>
<td class="name">ATK_VALUE_VERY_LOW</td>
<td class="value">5</td>
</tr>
<tr>
<td class="name">ATK_VALUE_LOW</td>
<td class="value">6</td>
</tr>
<tr>
<td class="name">ATK_VALUE_MEDIUM</td>
<td class="value">7</td>
</tr>
<tr>
<td class="name">ATK_VALUE_HIGH</td>
<td class="value">8</td>
</tr>
<tr>
<td class="name">ATK_VALUE_VERY_HIGH</td>
<td class="value">9</td>
</tr>
<tr>
<td class="name">ATK_VALUE_VERY_BAD</td>
<td class="value">10</td>
</tr>
<tr>
<td class="name">ATK_VALUE_BAD</td>
<td class="value">11</td>
</tr>
<tr>
<td class="name">ATK_VALUE_GOOD</td>
<td class="value">12</td>
</tr>
<tr>
<td class="name">ATK_VALUE_VERY_GOOD</td>
<td class="value">13</td>
</tr>
<tr>
<td class="name">ATK_VALUE_BEST</td>
<td class="value">14</td>
</tr>
<tr>
<td class="name">ATK_VALUE_LAST_DEFINED</td>
<td class="value">15</td>
</tr>
</table>
</div>