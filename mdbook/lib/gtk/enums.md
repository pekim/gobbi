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

## `ArrowPlacement`

Used to specify the placement of scroll arrows in scrolling menus.

C - `GtkArrowPlacement`

## `ArrowType`

Used to indicate the direction in which an arrow should point.

C - `GtkArrowType`

## `AssistantPageType`

An enum for determining the page role inside the #GtkAssistant. It's
used to handle buttons sensitivity and visibility.

Note that an assistant needs to end its page flow with a page of type
%GTK_ASSISTANT_PAGE_CONFIRM, %GTK_ASSISTANT_PAGE_SUMMARY or
%GTK_ASSISTANT_PAGE_PROGRESS to be correct.

The Cancel button will only be shown if the page isn’t “committed”.
See gtk_assistant_commit() for details.

C - `GtkAssistantPageType`

## `BaselinePosition`

Whenever a container has some form of natural row it may align
children in that row along a common typographical baseline. If
the amount of verical space in the row is taller than the total
requested height of the baseline-aligned children then it can use a
&num;GtkBaselinePosition to select where to put the baseline inside the
extra availible space.

C - `GtkBaselinePosition`

## `BorderStyle`

Describes how the border of a UI element should be rendered.

C - `GtkBorderStyle`

## `BuilderError`

Error codes that identify various errors that can occur while using
&num;GtkBuilder.

C - `GtkBuilderError`

## `ButtonBoxStyle`

Used to dictate the style that a #GtkButtonBox uses to layout the buttons it
contains.

C - `GtkButtonBoxStyle`

## `ButtonRole`

The role specifies the desired appearance of a #GtkModelButton.

C - `GtkButtonRole`

## `ButtonsType`

Prebuilt sets of buttons for the dialog. If
none of these choices are appropriate, simply use %GTK_BUTTONS_NONE
then call gtk_dialog_add_buttons().

> Please note that %GTK_BUTTONS_OK, %GTK_BUTTONS_YES_NO
> and %GTK_BUTTONS_OK_CANCEL are discouraged by the
> [GNOME Human Interface Guidelines](http://library.gnome.org/devel/hig-book/stable/).

C - `GtkButtonsType`

## `CellRendererAccelMode`

Determines if the edited accelerators are GTK+ accelerators. If
they are, consumed modifiers are suppressed, only accelerators
accepted by GTK+ are allowed, and the accelerators are rendered
in the same way as they are in menus.

C - `GtkCellRendererAccelMode`

## `CellRendererMode`

Identifies how the user can interact with a particular cell.

C - `GtkCellRendererMode`

## `CornerType`

Specifies which corner a child widget should be placed in when packed into
a #GtkScrolledWindow. This is effectively the opposite of where the scroll
bars are placed.

C - `GtkCornerType`

## `CssProviderError`

Error codes for %GTK_CSS_PROVIDER_ERROR.

C - `GtkCssProviderError`

## `CssSectionType`

The different types of sections indicate parts of a CSS document as
parsed by GTK’s CSS parser. They are oriented towards the
[CSS Grammar](http://www.w3.org/TR/CSS21/grammar.html),
but may contain extensions.

More types might be added in the future as the parser incorporates
more features.

C - `GtkCssSectionType`

## `DeleteType`

See also: #GtkEntry::delete-from-cursor.

C - `GtkDeleteType`

## `DirectionType`

Focus movement types.

C - `GtkDirectionType`

## `DragResult`

Gives an indication why a drag operation failed.
The value can by obtained by connecting to the
&num;GtkWidget::drag-failed signal.

C - `GtkDragResult`

## `EntryIconPosition`

Specifies the side of the entry at which an icon is placed.

C - `GtkEntryIconPosition`

## `EventSequenceState`

Describes the state of a #GdkEventSequence in a #GtkGesture.

C - `GtkEventSequenceState`

## `ExpanderStyle`

Used to specify the style of the expanders drawn by a #GtkTreeView.

C - `GtkExpanderStyle`

## `FileChooserAction`

Describes whether a #GtkFileChooser is being used to open existing files
or to save to a possibly new file.

C - `GtkFileChooserAction`

## `FileChooserConfirmation`

Used as a return value of handlers for the
&num;GtkFileChooser::confirm-overwrite signal of a #GtkFileChooser. This
value determines whether the file chooser will present the stock
confirmation dialog, accept the user’s choice of a filename, or
let the user choose another filename.

C - `GtkFileChooserConfirmation`

## `FileChooserError`

These identify the various errors that can occur while calling
&num;GtkFileChooser functions.

C - `GtkFileChooserError`

## `IMPreeditStyle`

Style for input method preedit. See also
&num;GtkSettings:gtk-im-preedit-style

C - `GtkIMPreeditStyle`

## `IMStatusStyle`

Style for input method status. See also
&num;GtkSettings:gtk-im-status-style

C - `GtkIMStatusStyle`

## `IconSize`

Built-in stock icon sizes.

C - `GtkIconSize`

## `IconThemeError`

Error codes for GtkIconTheme operations.

C - `GtkIconThemeError`

## `IconViewDropPosition`

An enum for determining where a dropped item goes.

C - `GtkIconViewDropPosition`

## `ImageType`

Describes the image data representation used by a #GtkImage. If you
want to get the image from the widget, you can only get the
currently-stored representation. e.g.  if the
gtk_image_get_storage_type() returns #GTK_IMAGE_PIXBUF, then you can
call gtk_image_get_pixbuf() but not gtk_image_get_stock().  For empty
images, you can request any storage type (call any of the "get"
functions), but they will all return %NULL values.

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

C - `GtkInputPurpose`

## `Justification`

Used for justifying the text inside a #GtkLabel widget. (See also
&num;GtkAlignment).

C - `GtkJustification`

## `LevelBarMode`

Describes how #GtkLevelBar contents should be rendered.
Note that this enumeration could be extended with additional modes
in the future.

C - `GtkLevelBarMode`

## `License`

The type of license for an application.

This enumeration can be expanded at later date.

C - `GtkLicense`

## `MenuDirectionType`

An enumeration representing directional movements within a menu.

C - `GtkMenuDirectionType`

## `MessageType`

The type of message being displayed in the dialog.

C - `GtkMessageType`

## `MovementStep`



C - `GtkMovementStep`

## `NotebookTab`



C - `GtkNotebookTab`

## `NumberUpLayout`

Used to determine the layout of pages on a sheet when printing
multiple pages per sheet.

C - `GtkNumberUpLayout`

## `Orientation`

Represents the orientation of widgets and other objects which can be switched
between horizontal and vertical orientation on the fly, like #GtkToolbar or
&num;GtkGesturePan.

C - `GtkOrientation`

## `PackDirection`

Determines how widgets should be packed inside menubars
and menuitems contained in menubars.

C - `GtkPackDirection`

## `PackType`

Represents the packing location #GtkBox children. (See: #GtkVBox,
&num;GtkHBox, and #GtkButtonBox).

C - `GtkPackType`

## `PadActionType`

The type of a pad action.

C - `GtkPadActionType`

## `PageOrientation`

See also gtk_print_settings_set_orientation().

C - `GtkPageOrientation`

## `PageSet`

See also gtk_print_job_set_page_set().

C - `GtkPageSet`

## `PanDirection`

Describes the panning direction of a #GtkGesturePan

C - `GtkPanDirection`

## `PathPriorityType`

Priorities for path lookups.
See also gtk_binding_set_add_path().

C - `GtkPathPriorityType`

## `PathType`

Widget path types.
See also gtk_binding_set_add_path().

C - `GtkPathType`

## `PolicyType`

Determines how the size should be computed to achieve the one of the
visibility mode for the scrollbars.

C - `GtkPolicyType`

## `PopoverConstraint`

Describes constraints to positioning of popovers. More values
may be added to this enumeration in the future.

C - `GtkPopoverConstraint`

## `PositionType`

Describes which edge of a widget a certain feature is positioned at, e.g. the
tabs of a #GtkNotebook, the handle of a #GtkHandleBox or the label of a
&num;GtkScale.

C - `GtkPositionType`

## `PrintDuplex`

See also gtk_print_settings_set_duplex().

C - `GtkPrintDuplex`

## `PrintError`

Error codes that identify various errors that can occur while
using the GTK+ printing support.

C - `GtkPrintError`

## `PrintOperationAction`

The @action parameter to gtk_print_operation_run()
determines what action the print operation should perform.

C - `GtkPrintOperationAction`

## `PrintOperationResult`

A value of this type is returned by gtk_print_operation_run().

C - `GtkPrintOperationResult`

## `PrintPages`

See also gtk_print_job_set_pages()

C - `GtkPrintPages`

## `PrintQuality`

See also gtk_print_settings_set_quality().

C - `GtkPrintQuality`

## `PrintStatus`

The status gives a rough indication of the completion of a running
print operation.

C - `GtkPrintStatus`

## `PropagationPhase`

Describes the stage at which events are fed into a #GtkEventController.

C - `GtkPropagationPhase`

## `RcTokenType`

The #GtkRcTokenType enumeration represents the tokens
in the RC file. It is exposed so that theme engines
can reuse these tokens when parsing the theme-engine
specific portions of a RC file.

C - `GtkRcTokenType`

## `RecentChooserError`

These identify the various errors that can occur while calling
&num;GtkRecentChooser functions.

C - `GtkRecentChooserError`

## `RecentManagerError`

Error codes for #GtkRecentManager operations

C - `GtkRecentManagerError`

## `RecentSortType`

Used to specify the sorting method to be applyed to the recently
used resource list.

C - `GtkRecentSortType`

## `ReliefStyle`

Indicated the relief to be drawn around a #GtkButton.

C - `GtkReliefStyle`

## `ResizeMode`



C - `GtkResizeMode`

## `ResponseType`

Predefined values for use as response ids in gtk_dialog_add_button().
All predefined values are negative; GTK+ leaves values of 0 or greater for
application-defined response ids.

C - `GtkResponseType`

## `RevealerTransitionType`

These enumeration values describe the possible transitions
when the child of a #GtkRevealer widget is shown or hidden.

C - `GtkRevealerTransitionType`

## `ScrollStep`



C - `GtkScrollStep`

## `ScrollType`

Scrolling types.

C - `GtkScrollType`

## `ScrollablePolicy`

Defines the policy to be used in a scrollable widget when updating
the scrolled window adjustments in a given orientation.

C - `GtkScrollablePolicy`

## `SelectionMode`

Used to control what selections users are allowed to make.

C - `GtkSelectionMode`

## `SensitivityType`

Determines how GTK+ handles the sensitivity of stepper arrows
at the end of range widgets.

C - `GtkSensitivityType`

## `ShadowType`

Used to change the appearance of an outline typically provided by a #GtkFrame.

Note that many themes do not differentiate the appearance of the
various shadow types: Either their is no visible shadow (@GTK_SHADOW_NONE),
or there is (any other value).

C - `GtkShadowType`

## `ShortcutType`

GtkShortcutType specifies the kind of shortcut that is being described.
More values may be added to this enumeration over time.

C - `GtkShortcutType`

## `SizeGroupMode`

The mode of the size group determines the directions in which the size
group affects the requested sizes of its component widgets.

C - `GtkSizeGroupMode`

## `SizeRequestMode`

Specifies a preference for height-for-width or
width-for-height geometry management.

C - `GtkSizeRequestMode`

## `SortType`

Determines the direction of a sort.

C - `GtkSortType`

## `SpinButtonUpdatePolicy`

The spin button update policy determines whether the spin button displays
values even if they are outside the bounds of its adjustment.
See gtk_spin_button_set_update_policy().

C - `GtkSpinButtonUpdatePolicy`

## `SpinType`

The values of the GtkSpinType enumeration are used to specify the
change to make in gtk_spin_button_spin().

C - `GtkSpinType`

## `StackTransitionType`

These enumeration values describe the possible transitions
between pages in a #GtkStack widget.

New values may be added to this enumeration over time.

C - `GtkStackTransitionType`

## `StateType`

This type indicates the current state of a widget; the state determines how
the widget is drawn. The #GtkStateType enumeration is also used to
identify different colors in a #GtkStyle for drawing, so states can be
used for subparts of a widget as well as entire widgets.

C - `GtkStateType`

## `TextBufferTargetInfo`

These values are used as “info” for the targets contained in the
lists returned by gtk_text_buffer_get_copy_target_list() and
gtk_text_buffer_get_paste_target_list().

The values counts down from `-1` to avoid clashes
with application added drag destinations which usually start at 0.

C - `GtkTextBufferTargetInfo`

## `TextDirection`

Reading directions for text.

C - `GtkTextDirection`

## `TextExtendSelection`

Granularity types that extend the text selection. Use the
&num;GtkTextView::extend-selection signal to customize the selection.

C - `GtkTextExtendSelection`

## `TextViewLayer`

Used to reference the layers of #GtkTextView for the purpose of customized
drawing with the ::draw_layer vfunc.

C - `GtkTextViewLayer`

## `TextWindowType`

Used to reference the parts of #GtkTextView.

C - `GtkTextWindowType`

## `ToolbarSpaceStyle`

Whether spacers are vertical lines or just blank.

C - `GtkToolbarSpaceStyle`

## `ToolbarStyle`

Used to customize the appearance of a #GtkToolbar. Note that
setting the toolbar style overrides the user’s preferences
for the default toolbar style.  Note that if the button has only
a label set and GTK_TOOLBAR_ICONS is used, the label will be
visible, and vice versa.

C - `GtkToolbarStyle`

## `TreeViewColumnSizing`

The sizing method the column uses to determine its width.  Please note
that @GTK_TREE_VIEW_COLUMN_AUTOSIZE are inefficient for large views, and
can make columns appear choppy.

C - `GtkTreeViewColumnSizing`

## `TreeViewDropPosition`

An enum for determining where a dropped row goes.

C - `GtkTreeViewDropPosition`

## `TreeViewGridLines`

Used to indicate which grid lines to draw in a tree view.

C - `GtkTreeViewGridLines`

## `Unit`

See also gtk_print_settings_set_paper_width().

C - `GtkUnit`

## `WidgetHelpType`

Kinds of widget-specific help. Used by the ::show-help signal.

C - `GtkWidgetHelpType`

## `WindowPosition`

Window placement can be influenced using this enumeration. Note that
using #GTK_WIN_POS_CENTER_ALWAYS is almost always a bad idea.
It won’t necessarily work well with all window managers or on all windowing systems.

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

C - `GtkWindowType`

## `WrapMode`

Describes a type of line wrapping.

C - `GtkWrapMode`

