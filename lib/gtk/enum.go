// This is a generated file - DO NOT EDIT

package gtk

import (
	glib "github.com/pekim/gobbi/lib/glib"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

type Align C.GtkAlign

const (
	GTK_ALIGN_FILL     Align = 0
	GTK_ALIGN_START    Align = 1
	GTK_ALIGN_END      Align = 2
	GTK_ALIGN_CENTER   Align = 3
	GTK_ALIGN_BASELINE Align = 4
)

type ArrowPlacement C.GtkArrowPlacement

const (
	GTK_ARROWS_BOTH  ArrowPlacement = 0
	GTK_ARROWS_START ArrowPlacement = 1
	GTK_ARROWS_END   ArrowPlacement = 2
)

type ArrowType C.GtkArrowType

const (
	GTK_ARROW_UP    ArrowType = 0
	GTK_ARROW_DOWN  ArrowType = 1
	GTK_ARROW_LEFT  ArrowType = 2
	GTK_ARROW_RIGHT ArrowType = 3
	GTK_ARROW_NONE  ArrowType = 4
)

type AssistantPageType C.GtkAssistantPageType

const (
	GTK_ASSISTANT_PAGE_CONTENT  AssistantPageType = 0
	GTK_ASSISTANT_PAGE_INTRO    AssistantPageType = 1
	GTK_ASSISTANT_PAGE_CONFIRM  AssistantPageType = 2
	GTK_ASSISTANT_PAGE_SUMMARY  AssistantPageType = 3
	GTK_ASSISTANT_PAGE_PROGRESS AssistantPageType = 4
	GTK_ASSISTANT_PAGE_CUSTOM   AssistantPageType = 5
)

type BorderStyle C.GtkBorderStyle

const (
	GTK_BORDER_STYLE_NONE   BorderStyle = 0
	GTK_BORDER_STYLE_SOLID  BorderStyle = 1
	GTK_BORDER_STYLE_INSET  BorderStyle = 2
	GTK_BORDER_STYLE_OUTSET BorderStyle = 3
	GTK_BORDER_STYLE_HIDDEN BorderStyle = 4
	GTK_BORDER_STYLE_DOTTED BorderStyle = 5
	GTK_BORDER_STYLE_DASHED BorderStyle = 6
	GTK_BORDER_STYLE_DOUBLE BorderStyle = 7
	GTK_BORDER_STYLE_GROOVE BorderStyle = 8
	GTK_BORDER_STYLE_RIDGE  BorderStyle = 9
)

type BuilderError C.GtkBuilderError

const (
	GTK_BUILDER_ERROR_INVALID_TYPE_FUNCTION  BuilderError = 0
	GTK_BUILDER_ERROR_UNHANDLED_TAG          BuilderError = 1
	GTK_BUILDER_ERROR_MISSING_ATTRIBUTE      BuilderError = 2
	GTK_BUILDER_ERROR_INVALID_ATTRIBUTE      BuilderError = 3
	GTK_BUILDER_ERROR_INVALID_TAG            BuilderError = 4
	GTK_BUILDER_ERROR_MISSING_PROPERTY_VALUE BuilderError = 5
	GTK_BUILDER_ERROR_INVALID_VALUE          BuilderError = 6
	GTK_BUILDER_ERROR_VERSION_MISMATCH       BuilderError = 7
	GTK_BUILDER_ERROR_DUPLICATE_ID           BuilderError = 8
	GTK_BUILDER_ERROR_OBJECT_TYPE_REFUSED    BuilderError = 9
	GTK_BUILDER_ERROR_TEMPLATE_MISMATCH      BuilderError = 10
	GTK_BUILDER_ERROR_INVALID_PROPERTY       BuilderError = 11
	GTK_BUILDER_ERROR_INVALID_SIGNAL         BuilderError = 12
	GTK_BUILDER_ERROR_INVALID_ID             BuilderError = 13
)

// BuilderErrorQuark is a wrapper around the C function gtk_builder_error_quark.
func BuilderErrorQuark() glib.Quark {
	retC := C.gtk_builder_error_quark()
	retGo := (glib.Quark)(retC)

	return retGo
}

type ButtonBoxStyle C.GtkButtonBoxStyle

const (
	GTK_BUTTONBOX_SPREAD ButtonBoxStyle = 1
	GTK_BUTTONBOX_EDGE   ButtonBoxStyle = 2
	GTK_BUTTONBOX_START  ButtonBoxStyle = 3
	GTK_BUTTONBOX_END    ButtonBoxStyle = 4
	GTK_BUTTONBOX_CENTER ButtonBoxStyle = 5
	GTK_BUTTONBOX_EXPAND ButtonBoxStyle = 6
)

type ButtonRole C.GtkButtonRole

const (
	GTK_BUTTON_ROLE_NORMAL ButtonRole = 0
	GTK_BUTTON_ROLE_CHECK  ButtonRole = 1
	GTK_BUTTON_ROLE_RADIO  ButtonRole = 2
)

type ButtonsType C.GtkButtonsType

const (
	GTK_BUTTONS_NONE      ButtonsType = 0
	GTK_BUTTONS_OK        ButtonsType = 1
	GTK_BUTTONS_CLOSE     ButtonsType = 2
	GTK_BUTTONS_CANCEL    ButtonsType = 3
	GTK_BUTTONS_YES_NO    ButtonsType = 4
	GTK_BUTTONS_OK_CANCEL ButtonsType = 5
)

type CellRendererAccelMode C.GtkCellRendererAccelMode

const (
	GTK_CELL_RENDERER_ACCEL_MODE_GTK          CellRendererAccelMode = 0
	GTK_CELL_RENDERER_ACCEL_MODE_OTHER        CellRendererAccelMode = 1
	GTK_CELL_RENDERER_ACCEL_MODE_MODIFIER_TAP CellRendererAccelMode = 2
)

type CellRendererMode C.GtkCellRendererMode

const (
	GTK_CELL_RENDERER_MODE_INERT       CellRendererMode = 0
	GTK_CELL_RENDERER_MODE_ACTIVATABLE CellRendererMode = 1
	GTK_CELL_RENDERER_MODE_EDITABLE    CellRendererMode = 2
)

type CornerType C.GtkCornerType

const (
	GTK_CORNER_TOP_LEFT     CornerType = 0
	GTK_CORNER_BOTTOM_LEFT  CornerType = 1
	GTK_CORNER_TOP_RIGHT    CornerType = 2
	GTK_CORNER_BOTTOM_RIGHT CornerType = 3
)

type CssProviderError C.GtkCssProviderError

const (
	GTK_CSS_PROVIDER_ERROR_FAILED        CssProviderError = 0
	GTK_CSS_PROVIDER_ERROR_SYNTAX        CssProviderError = 1
	GTK_CSS_PROVIDER_ERROR_IMPORT        CssProviderError = 2
	GTK_CSS_PROVIDER_ERROR_NAME          CssProviderError = 3
	GTK_CSS_PROVIDER_ERROR_DEPRECATED    CssProviderError = 4
	GTK_CSS_PROVIDER_ERROR_UNKNOWN_VALUE CssProviderError = 5
)

// CssProviderErrorQuark is a wrapper around the C function gtk_css_provider_error_quark.
func CssProviderErrorQuark() glib.Quark {
	retC := C.gtk_css_provider_error_quark()
	retGo := (glib.Quark)(retC)

	return retGo
}

type DeleteType C.GtkDeleteType

const (
	GTK_DELETE_CHARS             DeleteType = 0
	GTK_DELETE_WORD_ENDS         DeleteType = 1
	GTK_DELETE_WORDS             DeleteType = 2
	GTK_DELETE_DISPLAY_LINES     DeleteType = 3
	GTK_DELETE_DISPLAY_LINE_ENDS DeleteType = 4
	GTK_DELETE_PARAGRAPH_ENDS    DeleteType = 5
	GTK_DELETE_PARAGRAPHS        DeleteType = 6
	GTK_DELETE_WHITESPACE        DeleteType = 7
)

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

type ExpanderStyle C.GtkExpanderStyle

const (
	GTK_EXPANDER_COLLAPSED      ExpanderStyle = 0
	GTK_EXPANDER_SEMI_COLLAPSED ExpanderStyle = 1
	GTK_EXPANDER_SEMI_EXPANDED  ExpanderStyle = 2
	GTK_EXPANDER_EXPANDED       ExpanderStyle = 3
)

type FileChooserAction C.GtkFileChooserAction

const (
	GTK_FILE_CHOOSER_ACTION_OPEN          FileChooserAction = 0
	GTK_FILE_CHOOSER_ACTION_SAVE          FileChooserAction = 1
	GTK_FILE_CHOOSER_ACTION_SELECT_FOLDER FileChooserAction = 2
	GTK_FILE_CHOOSER_ACTION_CREATE_FOLDER FileChooserAction = 3
)

type FileChooserError C.GtkFileChooserError

const (
	GTK_FILE_CHOOSER_ERROR_NONEXISTENT         FileChooserError = 0
	GTK_FILE_CHOOSER_ERROR_BAD_FILENAME        FileChooserError = 1
	GTK_FILE_CHOOSER_ERROR_ALREADY_EXISTS      FileChooserError = 2
	GTK_FILE_CHOOSER_ERROR_INCOMPLETE_HOSTNAME FileChooserError = 3
)

type IMPreeditStyle C.GtkIMPreeditStyle

const (
	GTK_IM_PREEDIT_NOTHING  IMPreeditStyle = 0
	GTK_IM_PREEDIT_CALLBACK IMPreeditStyle = 1
	GTK_IM_PREEDIT_NONE     IMPreeditStyle = 2
)

type IMStatusStyle C.GtkIMStatusStyle

const (
	GTK_IM_STATUS_NOTHING  IMStatusStyle = 0
	GTK_IM_STATUS_CALLBACK IMStatusStyle = 1
	GTK_IM_STATUS_NONE     IMStatusStyle = 2
)

type IconSize C.GtkIconSize

const (
	GTK_ICON_SIZE_INVALID       IconSize = 0
	GTK_ICON_SIZE_MENU          IconSize = 1
	GTK_ICON_SIZE_SMALL_TOOLBAR IconSize = 2
	GTK_ICON_SIZE_LARGE_TOOLBAR IconSize = 3
	GTK_ICON_SIZE_BUTTON        IconSize = 4
	GTK_ICON_SIZE_DND           IconSize = 5
	GTK_ICON_SIZE_DIALOG        IconSize = 6
)

// IconSizeFromName is a wrapper around the C function gtk_icon_size_from_name.
func IconSizeFromName(name string) IconSize {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.gtk_icon_size_from_name(c_name)
	retGo := (IconSize)(retC)

	return retGo
}

// IconSizeGetName is a wrapper around the C function gtk_icon_size_get_name.
func IconSizeGetName(size IconSize) string {
	c_size := (C.GtkIconSize)(size)

	retC := C.gtk_icon_size_get_name(c_size)
	retGo := C.GoString(retC)

	return retGo
}

// IconSizeLookup is a wrapper around the C function gtk_icon_size_lookup.
func IconSizeLookup(size IconSize) (bool, int32, int32) {
	c_size := (C.GtkIconSize)(size)

	var c_width C.gint

	var c_height C.gint

	retC := C.gtk_icon_size_lookup(c_size, &c_width, &c_height)
	retGo := retC == C.TRUE

	width := (int32)(c_width)

	height := (int32)(c_height)

	return retGo, width, height
}

// IconSizeRegister is a wrapper around the C function gtk_icon_size_register.
func IconSizeRegister(name string, width int32, height int32) IconSize {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_width := (C.gint)(width)

	c_height := (C.gint)(height)

	retC := C.gtk_icon_size_register(c_name, c_width, c_height)
	retGo := (IconSize)(retC)

	return retGo
}

// IconSizeRegisterAlias is a wrapper around the C function gtk_icon_size_register_alias.
func IconSizeRegisterAlias(alias string, target IconSize) {
	c_alias := C.CString(alias)
	defer C.free(unsafe.Pointer(c_alias))

	c_target := (C.GtkIconSize)(target)

	C.gtk_icon_size_register_alias(c_alias, c_target)

	return
}

type IconThemeError C.GtkIconThemeError

const (
	GTK_ICON_THEME_NOT_FOUND IconThemeError = 0
	GTK_ICON_THEME_FAILED    IconThemeError = 1
)

// IconThemeErrorQuark is a wrapper around the C function gtk_icon_theme_error_quark.
func IconThemeErrorQuark() glib.Quark {
	retC := C.gtk_icon_theme_error_quark()
	retGo := (glib.Quark)(retC)

	return retGo
}

type IconViewDropPosition C.GtkIconViewDropPosition

const (
	GTK_ICON_VIEW_NO_DROP    IconViewDropPosition = 0
	GTK_ICON_VIEW_DROP_INTO  IconViewDropPosition = 1
	GTK_ICON_VIEW_DROP_LEFT  IconViewDropPosition = 2
	GTK_ICON_VIEW_DROP_RIGHT IconViewDropPosition = 3
	GTK_ICON_VIEW_DROP_ABOVE IconViewDropPosition = 4
	GTK_ICON_VIEW_DROP_BELOW IconViewDropPosition = 5
)

type ImageType C.GtkImageType

const (
	GTK_IMAGE_EMPTY     ImageType = 0
	GTK_IMAGE_PIXBUF    ImageType = 1
	GTK_IMAGE_STOCK     ImageType = 2
	GTK_IMAGE_ICON_SET  ImageType = 3
	GTK_IMAGE_ANIMATION ImageType = 4
	GTK_IMAGE_ICON_NAME ImageType = 5
	GTK_IMAGE_GICON     ImageType = 6
	GTK_IMAGE_SURFACE   ImageType = 7
)

type Justification C.GtkJustification

const (
	GTK_JUSTIFY_LEFT   Justification = 0
	GTK_JUSTIFY_RIGHT  Justification = 1
	GTK_JUSTIFY_CENTER Justification = 2
	GTK_JUSTIFY_FILL   Justification = 3
)

type MenuDirectionType C.GtkMenuDirectionType

const (
	GTK_MENU_DIR_PARENT MenuDirectionType = 0
	GTK_MENU_DIR_CHILD  MenuDirectionType = 1
	GTK_MENU_DIR_NEXT   MenuDirectionType = 2
	GTK_MENU_DIR_PREV   MenuDirectionType = 3
)

type MessageType C.GtkMessageType

const (
	GTK_MESSAGE_INFO     MessageType = 0
	GTK_MESSAGE_WARNING  MessageType = 1
	GTK_MESSAGE_QUESTION MessageType = 2
	GTK_MESSAGE_ERROR    MessageType = 3
	GTK_MESSAGE_OTHER    MessageType = 4
)

type MovementStep C.GtkMovementStep

const (
	GTK_MOVEMENT_LOGICAL_POSITIONS MovementStep = 0
	GTK_MOVEMENT_VISUAL_POSITIONS  MovementStep = 1
	GTK_MOVEMENT_WORDS             MovementStep = 2
	GTK_MOVEMENT_DISPLAY_LINES     MovementStep = 3
	GTK_MOVEMENT_DISPLAY_LINE_ENDS MovementStep = 4
	GTK_MOVEMENT_PARAGRAPHS        MovementStep = 5
	GTK_MOVEMENT_PARAGRAPH_ENDS    MovementStep = 6
	GTK_MOVEMENT_PAGES             MovementStep = 7
	GTK_MOVEMENT_BUFFER_ENDS       MovementStep = 8
	GTK_MOVEMENT_HORIZONTAL_PAGES  MovementStep = 9
)

type NotebookTab C.GtkNotebookTab

const (
	GTK_NOTEBOOK_TAB_FIRST NotebookTab = 0
	GTK_NOTEBOOK_TAB_LAST  NotebookTab = 1
)

type NumberUpLayout C.GtkNumberUpLayout

const (
	GTK_NUMBER_UP_LAYOUT_LEFT_TO_RIGHT_TOP_TO_BOTTOM NumberUpLayout = 0
	GTK_NUMBER_UP_LAYOUT_LEFT_TO_RIGHT_BOTTOM_TO_TOP NumberUpLayout = 1
	GTK_NUMBER_UP_LAYOUT_RIGHT_TO_LEFT_TOP_TO_BOTTOM NumberUpLayout = 2
	GTK_NUMBER_UP_LAYOUT_RIGHT_TO_LEFT_BOTTOM_TO_TOP NumberUpLayout = 3
	GTK_NUMBER_UP_LAYOUT_TOP_TO_BOTTOM_LEFT_TO_RIGHT NumberUpLayout = 4
	GTK_NUMBER_UP_LAYOUT_TOP_TO_BOTTOM_RIGHT_TO_LEFT NumberUpLayout = 5
	GTK_NUMBER_UP_LAYOUT_BOTTOM_TO_TOP_LEFT_TO_RIGHT NumberUpLayout = 6
	GTK_NUMBER_UP_LAYOUT_BOTTOM_TO_TOP_RIGHT_TO_LEFT NumberUpLayout = 7
)

type Orientation C.GtkOrientation

const (
	GTK_ORIENTATION_HORIZONTAL Orientation = 0
	GTK_ORIENTATION_VERTICAL   Orientation = 1
)

type PackDirection C.GtkPackDirection

const (
	GTK_PACK_DIRECTION_LTR PackDirection = 0
	GTK_PACK_DIRECTION_RTL PackDirection = 1
	GTK_PACK_DIRECTION_TTB PackDirection = 2
	GTK_PACK_DIRECTION_BTT PackDirection = 3
)

type PackType C.GtkPackType

const (
	GTK_PACK_START PackType = 0
	GTK_PACK_END   PackType = 1
)

type PageOrientation C.GtkPageOrientation

const (
	GTK_PAGE_ORIENTATION_PORTRAIT          PageOrientation = 0
	GTK_PAGE_ORIENTATION_LANDSCAPE         PageOrientation = 1
	GTK_PAGE_ORIENTATION_REVERSE_PORTRAIT  PageOrientation = 2
	GTK_PAGE_ORIENTATION_REVERSE_LANDSCAPE PageOrientation = 3
)

type PageSet C.GtkPageSet

const (
	GTK_PAGE_SET_ALL  PageSet = 0
	GTK_PAGE_SET_EVEN PageSet = 1
	GTK_PAGE_SET_ODD  PageSet = 2
)

type PathPriorityType C.GtkPathPriorityType

const (
	GTK_PATH_PRIO_LOWEST      PathPriorityType = 0
	GTK_PATH_PRIO_GTK         PathPriorityType = 4
	GTK_PATH_PRIO_APPLICATION PathPriorityType = 8
	GTK_PATH_PRIO_THEME       PathPriorityType = 10
	GTK_PATH_PRIO_RC          PathPriorityType = 12
	GTK_PATH_PRIO_HIGHEST     PathPriorityType = 15
)

type PathType C.GtkPathType

const (
	GTK_PATH_WIDGET       PathType = 0
	GTK_PATH_WIDGET_CLASS PathType = 1
	GTK_PATH_CLASS        PathType = 2
)

type PolicyType C.GtkPolicyType

const (
	GTK_POLICY_ALWAYS    PolicyType = 0
	GTK_POLICY_AUTOMATIC PolicyType = 1
	GTK_POLICY_NEVER     PolicyType = 2
	GTK_POLICY_EXTERNAL  PolicyType = 3
)

type PositionType C.GtkPositionType

const (
	GTK_POS_LEFT   PositionType = 0
	GTK_POS_RIGHT  PositionType = 1
	GTK_POS_TOP    PositionType = 2
	GTK_POS_BOTTOM PositionType = 3
)

type PrintDuplex C.GtkPrintDuplex

const (
	GTK_PRINT_DUPLEX_SIMPLEX    PrintDuplex = 0
	GTK_PRINT_DUPLEX_HORIZONTAL PrintDuplex = 1
	GTK_PRINT_DUPLEX_VERTICAL   PrintDuplex = 2
)

type PrintError C.GtkPrintError

const (
	GTK_PRINT_ERROR_GENERAL        PrintError = 0
	GTK_PRINT_ERROR_INTERNAL_ERROR PrintError = 1
	GTK_PRINT_ERROR_NOMEM          PrintError = 2
	GTK_PRINT_ERROR_INVALID_FILE   PrintError = 3
)

type PrintOperationAction C.GtkPrintOperationAction

const (
	GTK_PRINT_OPERATION_ACTION_PRINT_DIALOG PrintOperationAction = 0
	GTK_PRINT_OPERATION_ACTION_PRINT        PrintOperationAction = 1
	GTK_PRINT_OPERATION_ACTION_PREVIEW      PrintOperationAction = 2
	GTK_PRINT_OPERATION_ACTION_EXPORT       PrintOperationAction = 3
)

type PrintOperationResult C.GtkPrintOperationResult

const (
	GTK_PRINT_OPERATION_RESULT_ERROR       PrintOperationResult = 0
	GTK_PRINT_OPERATION_RESULT_APPLY       PrintOperationResult = 1
	GTK_PRINT_OPERATION_RESULT_CANCEL      PrintOperationResult = 2
	GTK_PRINT_OPERATION_RESULT_IN_PROGRESS PrintOperationResult = 3
)

type PrintPages C.GtkPrintPages

const (
	GTK_PRINT_PAGES_ALL       PrintPages = 0
	GTK_PRINT_PAGES_CURRENT   PrintPages = 1
	GTK_PRINT_PAGES_RANGES    PrintPages = 2
	GTK_PRINT_PAGES_SELECTION PrintPages = 3
)

type PrintQuality C.GtkPrintQuality

const (
	GTK_PRINT_QUALITY_LOW    PrintQuality = 0
	GTK_PRINT_QUALITY_NORMAL PrintQuality = 1
	GTK_PRINT_QUALITY_HIGH   PrintQuality = 2
	GTK_PRINT_QUALITY_DRAFT  PrintQuality = 3
)

type PrintStatus C.GtkPrintStatus

const (
	GTK_PRINT_STATUS_INITIAL          PrintStatus = 0
	GTK_PRINT_STATUS_PREPARING        PrintStatus = 1
	GTK_PRINT_STATUS_GENERATING_DATA  PrintStatus = 2
	GTK_PRINT_STATUS_SENDING_DATA     PrintStatus = 3
	GTK_PRINT_STATUS_PENDING          PrintStatus = 4
	GTK_PRINT_STATUS_PENDING_ISSUE    PrintStatus = 5
	GTK_PRINT_STATUS_PRINTING         PrintStatus = 6
	GTK_PRINT_STATUS_FINISHED         PrintStatus = 7
	GTK_PRINT_STATUS_FINISHED_ABORTED PrintStatus = 8
)

type RcTokenType C.GtkRcTokenType

const (
	GTK_RC_TOKEN_INVALID        RcTokenType = 270
	GTK_RC_TOKEN_INCLUDE        RcTokenType = 271
	GTK_RC_TOKEN_NORMAL         RcTokenType = 272
	GTK_RC_TOKEN_ACTIVE         RcTokenType = 273
	GTK_RC_TOKEN_PRELIGHT       RcTokenType = 274
	GTK_RC_TOKEN_SELECTED       RcTokenType = 275
	GTK_RC_TOKEN_INSENSITIVE    RcTokenType = 276
	GTK_RC_TOKEN_FG             RcTokenType = 277
	GTK_RC_TOKEN_BG             RcTokenType = 278
	GTK_RC_TOKEN_TEXT           RcTokenType = 279
	GTK_RC_TOKEN_BASE           RcTokenType = 280
	GTK_RC_TOKEN_XTHICKNESS     RcTokenType = 281
	GTK_RC_TOKEN_YTHICKNESS     RcTokenType = 282
	GTK_RC_TOKEN_FONT           RcTokenType = 283
	GTK_RC_TOKEN_FONTSET        RcTokenType = 284
	GTK_RC_TOKEN_FONT_NAME      RcTokenType = 285
	GTK_RC_TOKEN_BG_PIXMAP      RcTokenType = 286
	GTK_RC_TOKEN_PIXMAP_PATH    RcTokenType = 287
	GTK_RC_TOKEN_STYLE          RcTokenType = 288
	GTK_RC_TOKEN_BINDING        RcTokenType = 289
	GTK_RC_TOKEN_BIND           RcTokenType = 290
	GTK_RC_TOKEN_WIDGET         RcTokenType = 291
	GTK_RC_TOKEN_WIDGET_CLASS   RcTokenType = 292
	GTK_RC_TOKEN_CLASS          RcTokenType = 293
	GTK_RC_TOKEN_LOWEST         RcTokenType = 294
	GTK_RC_TOKEN_GTK            RcTokenType = 295
	GTK_RC_TOKEN_APPLICATION    RcTokenType = 296
	GTK_RC_TOKEN_THEME          RcTokenType = 297
	GTK_RC_TOKEN_RC             RcTokenType = 298
	GTK_RC_TOKEN_HIGHEST        RcTokenType = 299
	GTK_RC_TOKEN_ENGINE         RcTokenType = 300
	GTK_RC_TOKEN_MODULE_PATH    RcTokenType = 301
	GTK_RC_TOKEN_IM_MODULE_PATH RcTokenType = 302
	GTK_RC_TOKEN_IM_MODULE_FILE RcTokenType = 303
	GTK_RC_TOKEN_STOCK          RcTokenType = 304
	GTK_RC_TOKEN_LTR            RcTokenType = 305
	GTK_RC_TOKEN_RTL            RcTokenType = 306
	GTK_RC_TOKEN_COLOR          RcTokenType = 307
	GTK_RC_TOKEN_UNBIND         RcTokenType = 308
	GTK_RC_TOKEN_LAST           RcTokenType = 309
)

type ReliefStyle C.GtkReliefStyle

const (
	GTK_RELIEF_NORMAL ReliefStyle = 0
	GTK_RELIEF_HALF   ReliefStyle = 1
	GTK_RELIEF_NONE   ReliefStyle = 2
)

type ResizeMode C.GtkResizeMode

const (
	GTK_RESIZE_PARENT    ResizeMode = 0
	GTK_RESIZE_QUEUE     ResizeMode = 1
	GTK_RESIZE_IMMEDIATE ResizeMode = 2
)

type ResponseType C.GtkResponseType

const (
	GTK_RESPONSE_NONE         ResponseType = -1
	GTK_RESPONSE_REJECT       ResponseType = -2
	GTK_RESPONSE_ACCEPT       ResponseType = -3
	GTK_RESPONSE_DELETE_EVENT ResponseType = -4
	GTK_RESPONSE_OK           ResponseType = -5
	GTK_RESPONSE_CANCEL       ResponseType = -6
	GTK_RESPONSE_CLOSE        ResponseType = -7
	GTK_RESPONSE_YES          ResponseType = -8
	GTK_RESPONSE_NO           ResponseType = -9
	GTK_RESPONSE_APPLY        ResponseType = -10
	GTK_RESPONSE_HELP         ResponseType = -11
)

type RevealerTransitionType C.GtkRevealerTransitionType

const (
	GTK_REVEALER_TRANSITION_TYPE_NONE        RevealerTransitionType = 0
	GTK_REVEALER_TRANSITION_TYPE_CROSSFADE   RevealerTransitionType = 1
	GTK_REVEALER_TRANSITION_TYPE_SLIDE_RIGHT RevealerTransitionType = 2
	GTK_REVEALER_TRANSITION_TYPE_SLIDE_LEFT  RevealerTransitionType = 3
	GTK_REVEALER_TRANSITION_TYPE_SLIDE_UP    RevealerTransitionType = 4
	GTK_REVEALER_TRANSITION_TYPE_SLIDE_DOWN  RevealerTransitionType = 5
)

type ScrollStep C.GtkScrollStep

const (
	GTK_SCROLL_STEPS            ScrollStep = 0
	GTK_SCROLL_PAGES            ScrollStep = 1
	GTK_SCROLL_ENDS             ScrollStep = 2
	GTK_SCROLL_HORIZONTAL_STEPS ScrollStep = 3
	GTK_SCROLL_HORIZONTAL_PAGES ScrollStep = 4
	GTK_SCROLL_HORIZONTAL_ENDS  ScrollStep = 5
)

type ScrollType C.GtkScrollType

const (
	GTK_SCROLL_NONE          ScrollType = 0
	GTK_SCROLL_JUMP          ScrollType = 1
	GTK_SCROLL_STEP_BACKWARD ScrollType = 2
	GTK_SCROLL_STEP_FORWARD  ScrollType = 3
	GTK_SCROLL_PAGE_BACKWARD ScrollType = 4
	GTK_SCROLL_PAGE_FORWARD  ScrollType = 5
	GTK_SCROLL_STEP_UP       ScrollType = 6
	GTK_SCROLL_STEP_DOWN     ScrollType = 7
	GTK_SCROLL_PAGE_UP       ScrollType = 8
	GTK_SCROLL_PAGE_DOWN     ScrollType = 9
	GTK_SCROLL_STEP_LEFT     ScrollType = 10
	GTK_SCROLL_STEP_RIGHT    ScrollType = 11
	GTK_SCROLL_PAGE_LEFT     ScrollType = 12
	GTK_SCROLL_PAGE_RIGHT    ScrollType = 13
	GTK_SCROLL_START         ScrollType = 14
	GTK_SCROLL_END           ScrollType = 15
)

type ScrollablePolicy C.GtkScrollablePolicy

const (
	GTK_SCROLL_MINIMUM ScrollablePolicy = 0
	GTK_SCROLL_NATURAL ScrollablePolicy = 1
)

type SelectionMode C.GtkSelectionMode

const (
	GTK_SELECTION_NONE     SelectionMode = 0
	GTK_SELECTION_SINGLE   SelectionMode = 1
	GTK_SELECTION_BROWSE   SelectionMode = 2
	GTK_SELECTION_MULTIPLE SelectionMode = 3
)

type SensitivityType C.GtkSensitivityType

const (
	GTK_SENSITIVITY_AUTO SensitivityType = 0
	GTK_SENSITIVITY_ON   SensitivityType = 1
	GTK_SENSITIVITY_OFF  SensitivityType = 2
)

type ShadowType C.GtkShadowType

const (
	GTK_SHADOW_NONE       ShadowType = 0
	GTK_SHADOW_IN         ShadowType = 1
	GTK_SHADOW_OUT        ShadowType = 2
	GTK_SHADOW_ETCHED_IN  ShadowType = 3
	GTK_SHADOW_ETCHED_OUT ShadowType = 4
)

type SizeGroupMode C.GtkSizeGroupMode

const (
	GTK_SIZE_GROUP_NONE       SizeGroupMode = 0
	GTK_SIZE_GROUP_HORIZONTAL SizeGroupMode = 1
	GTK_SIZE_GROUP_VERTICAL   SizeGroupMode = 2
	GTK_SIZE_GROUP_BOTH       SizeGroupMode = 3
)

type SizeRequestMode C.GtkSizeRequestMode

const (
	GTK_SIZE_REQUEST_HEIGHT_FOR_WIDTH SizeRequestMode = 0
	GTK_SIZE_REQUEST_WIDTH_FOR_HEIGHT SizeRequestMode = 1
	GTK_SIZE_REQUEST_CONSTANT_SIZE    SizeRequestMode = 2
)

type SortType C.GtkSortType

const (
	GTK_SORT_ASCENDING  SortType = 0
	GTK_SORT_DESCENDING SortType = 1
)

type SpinButtonUpdatePolicy C.GtkSpinButtonUpdatePolicy

const (
	GTK_UPDATE_ALWAYS   SpinButtonUpdatePolicy = 0
	GTK_UPDATE_IF_VALID SpinButtonUpdatePolicy = 1
)

type SpinType C.GtkSpinType

const (
	GTK_SPIN_STEP_FORWARD  SpinType = 0
	GTK_SPIN_STEP_BACKWARD SpinType = 1
	GTK_SPIN_PAGE_FORWARD  SpinType = 2
	GTK_SPIN_PAGE_BACKWARD SpinType = 3
	GTK_SPIN_HOME          SpinType = 4
	GTK_SPIN_END           SpinType = 5
	GTK_SPIN_USER_DEFINED  SpinType = 6
)

type StackTransitionType C.GtkStackTransitionType

const (
	GTK_STACK_TRANSITION_TYPE_NONE             StackTransitionType = 0
	GTK_STACK_TRANSITION_TYPE_CROSSFADE        StackTransitionType = 1
	GTK_STACK_TRANSITION_TYPE_SLIDE_RIGHT      StackTransitionType = 2
	GTK_STACK_TRANSITION_TYPE_SLIDE_LEFT       StackTransitionType = 3
	GTK_STACK_TRANSITION_TYPE_SLIDE_UP         StackTransitionType = 4
	GTK_STACK_TRANSITION_TYPE_SLIDE_DOWN       StackTransitionType = 5
	GTK_STACK_TRANSITION_TYPE_SLIDE_LEFT_RIGHT StackTransitionType = 6
	GTK_STACK_TRANSITION_TYPE_SLIDE_UP_DOWN    StackTransitionType = 7
	GTK_STACK_TRANSITION_TYPE_OVER_UP          StackTransitionType = 8
	GTK_STACK_TRANSITION_TYPE_OVER_DOWN        StackTransitionType = 9
	GTK_STACK_TRANSITION_TYPE_OVER_LEFT        StackTransitionType = 10
	GTK_STACK_TRANSITION_TYPE_OVER_RIGHT       StackTransitionType = 11
	GTK_STACK_TRANSITION_TYPE_UNDER_UP         StackTransitionType = 12
	GTK_STACK_TRANSITION_TYPE_UNDER_DOWN       StackTransitionType = 13
	GTK_STACK_TRANSITION_TYPE_UNDER_LEFT       StackTransitionType = 14
	GTK_STACK_TRANSITION_TYPE_UNDER_RIGHT      StackTransitionType = 15
	GTK_STACK_TRANSITION_TYPE_OVER_UP_DOWN     StackTransitionType = 16
	GTK_STACK_TRANSITION_TYPE_OVER_DOWN_UP     StackTransitionType = 17
	GTK_STACK_TRANSITION_TYPE_OVER_LEFT_RIGHT  StackTransitionType = 18
	GTK_STACK_TRANSITION_TYPE_OVER_RIGHT_LEFT  StackTransitionType = 19
)

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

type TextBufferTargetInfo C.GtkTextBufferTargetInfo

const (
	GTK_TEXT_BUFFER_TARGET_INFO_BUFFER_CONTENTS TextBufferTargetInfo = -1
	GTK_TEXT_BUFFER_TARGET_INFO_RICH_TEXT       TextBufferTargetInfo = -2
	GTK_TEXT_BUFFER_TARGET_INFO_TEXT            TextBufferTargetInfo = -3
)

type TextDirection C.GtkTextDirection

const (
	GTK_TEXT_DIR_NONE TextDirection = 0
	GTK_TEXT_DIR_LTR  TextDirection = 1
	GTK_TEXT_DIR_RTL  TextDirection = 2
)

type TextViewLayer C.GtkTextViewLayer

const (
	GTK_TEXT_VIEW_LAYER_BELOW      TextViewLayer = 0
	GTK_TEXT_VIEW_LAYER_ABOVE      TextViewLayer = 1
	GTK_TEXT_VIEW_LAYER_BELOW_TEXT TextViewLayer = 2
	GTK_TEXT_VIEW_LAYER_ABOVE_TEXT TextViewLayer = 3
)

type TextWindowType C.GtkTextWindowType

const (
	GTK_TEXT_WINDOW_PRIVATE TextWindowType = 0
	GTK_TEXT_WINDOW_WIDGET  TextWindowType = 1
	GTK_TEXT_WINDOW_TEXT    TextWindowType = 2
	GTK_TEXT_WINDOW_LEFT    TextWindowType = 3
	GTK_TEXT_WINDOW_RIGHT   TextWindowType = 4
	GTK_TEXT_WINDOW_TOP     TextWindowType = 5
	GTK_TEXT_WINDOW_BOTTOM  TextWindowType = 6
)

type ToolbarSpaceStyle C.GtkToolbarSpaceStyle

const (
	GTK_TOOLBAR_SPACE_EMPTY ToolbarSpaceStyle = 0
	GTK_TOOLBAR_SPACE_LINE  ToolbarSpaceStyle = 1
)

type ToolbarStyle C.GtkToolbarStyle

const (
	GTK_TOOLBAR_ICONS      ToolbarStyle = 0
	GTK_TOOLBAR_TEXT       ToolbarStyle = 1
	GTK_TOOLBAR_BOTH       ToolbarStyle = 2
	GTK_TOOLBAR_BOTH_HORIZ ToolbarStyle = 3
)

type TreeViewColumnSizing C.GtkTreeViewColumnSizing

const (
	GTK_TREE_VIEW_COLUMN_GROW_ONLY TreeViewColumnSizing = 0
	GTK_TREE_VIEW_COLUMN_AUTOSIZE  TreeViewColumnSizing = 1
	GTK_TREE_VIEW_COLUMN_FIXED     TreeViewColumnSizing = 2
)

type TreeViewDropPosition C.GtkTreeViewDropPosition

const (
	GTK_TREE_VIEW_DROP_BEFORE         TreeViewDropPosition = 0
	GTK_TREE_VIEW_DROP_AFTER          TreeViewDropPosition = 1
	GTK_TREE_VIEW_DROP_INTO_OR_BEFORE TreeViewDropPosition = 2
	GTK_TREE_VIEW_DROP_INTO_OR_AFTER  TreeViewDropPosition = 3
)

type TreeViewGridLines C.GtkTreeViewGridLines

const (
	GTK_TREE_VIEW_GRID_LINES_NONE       TreeViewGridLines = 0
	GTK_TREE_VIEW_GRID_LINES_HORIZONTAL TreeViewGridLines = 1
	GTK_TREE_VIEW_GRID_LINES_VERTICAL   TreeViewGridLines = 2
	GTK_TREE_VIEW_GRID_LINES_BOTH       TreeViewGridLines = 3
)

type Unit C.GtkUnit

const (
	GTK_UNIT_NONE   Unit = 0
	GTK_UNIT_POINTS Unit = 1
	GTK_UNIT_INCH   Unit = 2
	GTK_UNIT_MM     Unit = 3
)

type WidgetHelpType C.GtkWidgetHelpType

const (
	GTK_WIDGET_HELP_TOOLTIP    WidgetHelpType = 0
	GTK_WIDGET_HELP_WHATS_THIS WidgetHelpType = 1
)

type WindowPosition C.GtkWindowPosition

const (
	GTK_WIN_POS_NONE             WindowPosition = 0
	GTK_WIN_POS_CENTER           WindowPosition = 1
	GTK_WIN_POS_MOUSE            WindowPosition = 2
	GTK_WIN_POS_CENTER_ALWAYS    WindowPosition = 3
	GTK_WIN_POS_CENTER_ON_PARENT WindowPosition = 4
)

type WindowType C.GtkWindowType

const (
	GTK_WINDOW_TOPLEVEL WindowType = 0
	GTK_WINDOW_POPUP    WindowType = 1
)

type WrapMode C.GtkWrapMode

const (
	GTK_WRAP_NONE      WrapMode = 0
	GTK_WRAP_CHAR      WrapMode = 1
	GTK_WRAP_WORD      WrapMode = 2
	GTK_WRAP_WORD_CHAR WrapMode = 3
)
