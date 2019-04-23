// This is a generated file - DO NOT EDIT

package gtk

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

// Blacklisted : GtkAlign

// Blacklisted : GtkArrowPlacement

// Blacklisted : GtkArrowType

// Blacklisted : GtkAssistantPageType

// Blacklisted : GtkBorderStyle

// Blacklisted : GtkBuilderError

// Blacklisted : GtkButtonBoxStyle

// Blacklisted : GtkButtonRole

// Blacklisted : GtkButtonsType

// Blacklisted : GtkCellRendererAccelMode

// Blacklisted : GtkCellRendererMode

// Blacklisted : GtkCornerType

// Blacklisted : GtkCssProviderError

// Blacklisted : GtkDeleteType

type DirectionType C.GtkDirectionType

const (
	GTK_DIR_TAB_FORWARD  DirectionType = 0
	GTK_DIR_TAB_BACKWARD DirectionType = 1
	GTK_DIR_UP           DirectionType = 2
	GTK_DIR_DOWN         DirectionType = 3
	GTK_DIR_LEFT         DirectionType = 4
	GTK_DIR_RIGHT        DirectionType = 5
)

type DragResult C.GtkDragResult

const (
	GTK_DRAG_RESULT_SUCCESS         DragResult = 0
	GTK_DRAG_RESULT_NO_TARGET       DragResult = 1
	GTK_DRAG_RESULT_USER_CANCELLED  DragResult = 2
	GTK_DRAG_RESULT_TIMEOUT_EXPIRED DragResult = 3
	GTK_DRAG_RESULT_GRAB_BROKEN     DragResult = 4
	GTK_DRAG_RESULT_ERROR           DragResult = 5
)

// Blacklisted : GtkExpanderStyle

// Blacklisted : GtkFileChooserAction

// Blacklisted : GtkFileChooserError

// Blacklisted : GtkIMPreeditStyle

// Blacklisted : GtkIMStatusStyle

// Blacklisted : GtkIconSize

// Blacklisted : GtkIconThemeError

// Blacklisted : GtkIconViewDropPosition

// Blacklisted : GtkImageType

// Blacklisted : GtkJustification

// Blacklisted : GtkMenuDirectionType

// Blacklisted : GtkMessageType

// Blacklisted : GtkMovementStep

// Blacklisted : GtkNotebookTab

// Blacklisted : GtkNumberUpLayout

// Blacklisted : GtkOrientation

// Blacklisted : GtkPackDirection

// Blacklisted : GtkPackType

// Blacklisted : GtkPageOrientation

// Blacklisted : GtkPageSet

// Blacklisted : GtkPathPriorityType

// Blacklisted : GtkPathType

// Blacklisted : GtkPolicyType

// Blacklisted : GtkPositionType

// Blacklisted : GtkPrintDuplex

// Blacklisted : GtkPrintError

// Blacklisted : GtkPrintOperationAction

// Blacklisted : GtkPrintOperationResult

// Blacklisted : GtkPrintPages

// Blacklisted : GtkPrintQuality

// Blacklisted : GtkPrintStatus

// Blacklisted : GtkRcTokenType

// Blacklisted : GtkReliefStyle

// Blacklisted : GtkResizeMode

// Blacklisted : GtkResponseType

// Blacklisted : GtkRevealerTransitionType

// Blacklisted : GtkScrollStep

// Blacklisted : GtkScrollType

// Blacklisted : GtkScrollablePolicy

// Blacklisted : GtkSelectionMode

// Blacklisted : GtkSensitivityType

// Blacklisted : GtkShadowType

// Blacklisted : GtkSizeGroupMode

// Blacklisted : GtkSizeRequestMode

// Blacklisted : GtkSortType

// Blacklisted : GtkSpinButtonUpdatePolicy

// Blacklisted : GtkSpinType

// Blacklisted : GtkStackTransitionType

type StateType C.GtkStateType

const (
	GTK_STATE_NORMAL       StateType = 0
	GTK_STATE_ACTIVE       StateType = 1
	GTK_STATE_PRELIGHT     StateType = 2
	GTK_STATE_SELECTED     StateType = 3
	GTK_STATE_INSENSITIVE  StateType = 4
	GTK_STATE_INCONSISTENT StateType = 5
	GTK_STATE_FOCUSED      StateType = 6
)

// Blacklisted : GtkTextBufferTargetInfo

type TextDirection C.GtkTextDirection

const (
	GTK_TEXT_DIR_NONE TextDirection = 0
	GTK_TEXT_DIR_LTR  TextDirection = 1
	GTK_TEXT_DIR_RTL  TextDirection = 2
)

// Blacklisted : GtkTextViewLayer

// Blacklisted : GtkTextWindowType

// Blacklisted : GtkToolbarSpaceStyle

// Blacklisted : GtkToolbarStyle

// Blacklisted : GtkTreeViewColumnSizing

// Blacklisted : GtkTreeViewDropPosition

// Blacklisted : GtkTreeViewGridLines

// Blacklisted : GtkUnit

type WidgetHelpType C.GtkWidgetHelpType

const (
	GTK_WIDGET_HELP_TOOLTIP    WidgetHelpType = 0
	GTK_WIDGET_HELP_WHATS_THIS WidgetHelpType = 1
)

// Blacklisted : GtkWindowPosition

type WindowType C.GtkWindowType

const (
	GTK_WINDOW_TOPLEVEL WindowType = 0
	GTK_WINDOW_POPUP    WindowType = 1
)

// Blacklisted : GtkWrapMode
