// Code generated - DO NOT EDIT.

package gdk

// Axisuse is a representation of the C type AxisUse.
type Axisuse int

const (
	// ignore
	GdkAxisIgnore Axisuse = 0
	// x
	GdkAxisX Axisuse = 1
	// y
	GdkAxisY Axisuse = 2
	// pressure
	GdkAxisPressure Axisuse = 3
	// xtilt
	GdkAxisXtilt Axisuse = 4
	// ytilt
	GdkAxisYtilt Axisuse = 5
	// wheel
	GdkAxisWheel Axisuse = 6
	// distance
	GdkAxisDistance Axisuse = 7
	// rotation
	GdkAxisRotation Axisuse = 8
	// slider
	GdkAxisSlider Axisuse = 9
	// last
	GdkAxisLast Axisuse = 10
)

// Byteorder is a representation of the C type ByteOrder.
type Byteorder int

const (
	// lsb_first
	GdkLsbFirst Byteorder = 0
	// msb_first
	GdkMsbFirst Byteorder = 1
)

// Crossingmode is a representation of the C type CrossingMode.
type Crossingmode int

const (
	// normal
	GdkCrossingNormal Crossingmode = 0
	// grab
	GdkCrossingGrab Crossingmode = 1
	// ungrab
	GdkCrossingUngrab Crossingmode = 2
	// gtk_grab
	GdkCrossingGtkGrab Crossingmode = 3
	// gtk_ungrab
	GdkCrossingGtkUngrab Crossingmode = 4
	// state_changed
	GdkCrossingStateChanged Crossingmode = 5
	// touch_begin
	GdkCrossingTouchBegin Crossingmode = 6
	// touch_end
	GdkCrossingTouchEnd Crossingmode = 7
	// device_switch
	GdkCrossingDeviceSwitch Crossingmode = 8
)

// Cursortype is a representation of the C type CursorType.
type Cursortype int

const (
	// x_cursor
	GdkXCursor Cursortype = 0
	// arrow
	GdkArrow Cursortype = 2
	// based_arrow_down
	GdkBasedArrowDown Cursortype = 4
	// based_arrow_up
	GdkBasedArrowUp Cursortype = 6
	// boat
	GdkBoat Cursortype = 8
	// bogosity
	GdkBogosity Cursortype = 10
	// bottom_left_corner
	GdkBottomLeftCorner Cursortype = 12
	// bottom_right_corner
	GdkBottomRightCorner Cursortype = 14
	// bottom_side
	GdkBottomSide Cursortype = 16
	// bottom_tee
	GdkBottomTee Cursortype = 18
	// box_spiral
	GdkBoxSpiral Cursortype = 20
	// center_ptr
	GdkCenterPtr Cursortype = 22
	// circle
	GdkCircle Cursortype = 24
	// clock
	GdkClock Cursortype = 26
	// coffee_mug
	GdkCoffeeMug Cursortype = 28
	// cross
	GdkCross Cursortype = 30
	// cross_reverse
	GdkCrossReverse Cursortype = 32
	// crosshair
	GdkCrosshair Cursortype = 34
	// diamond_cross
	GdkDiamondCross Cursortype = 36
	// dot
	GdkDot Cursortype = 38
	// dotbox
	GdkDotbox Cursortype = 40
	// double_arrow
	GdkDoubleArrow Cursortype = 42
	// draft_large
	GdkDraftLarge Cursortype = 44
	// draft_small
	GdkDraftSmall Cursortype = 46
	// draped_box
	GdkDrapedBox Cursortype = 48
	// exchange
	GdkExchange Cursortype = 50
	// fleur
	GdkFleur Cursortype = 52
	// gobbler
	GdkGobbler Cursortype = 54
	// gumby
	GdkGumby Cursortype = 56
	// hand1
	GdkHand1 Cursortype = 58
	// hand2
	GdkHand2 Cursortype = 60
	// heart
	GdkHeart Cursortype = 62
	// icon
	GdkIcon Cursortype = 64
	// iron_cross
	GdkIronCross Cursortype = 66
	// left_ptr
	GdkLeftPtr Cursortype = 68
	// left_side
	GdkLeftSide Cursortype = 70
	// left_tee
	GdkLeftTee Cursortype = 72
	// leftbutton
	GdkLeftbutton Cursortype = 74
	// ll_angle
	GdkLlAngle Cursortype = 76
	// lr_angle
	GdkLrAngle Cursortype = 78
	// man
	GdkMan Cursortype = 80
	// middlebutton
	GdkMiddlebutton Cursortype = 82
	// mouse
	GdkMouse Cursortype = 84
	// pencil
	GdkPencil Cursortype = 86
	// pirate
	GdkPirate Cursortype = 88
	// plus
	GdkPlus Cursortype = 90
	// question_arrow
	GdkQuestionArrow Cursortype = 92
	// right_ptr
	GdkRightPtr Cursortype = 94
	// right_side
	GdkRightSide Cursortype = 96
	// right_tee
	GdkRightTee Cursortype = 98
	// rightbutton
	GdkRightbutton Cursortype = 100
	// rtl_logo
	GdkRtlLogo Cursortype = 102
	// sailboat
	GdkSailboat Cursortype = 104
	// sb_down_arrow
	GdkSbDownArrow Cursortype = 106
	// sb_h_double_arrow
	GdkSbHDoubleArrow Cursortype = 108
	// sb_left_arrow
	GdkSbLeftArrow Cursortype = 110
	// sb_right_arrow
	GdkSbRightArrow Cursortype = 112
	// sb_up_arrow
	GdkSbUpArrow Cursortype = 114
	// sb_v_double_arrow
	GdkSbVDoubleArrow Cursortype = 116
	// shuttle
	GdkShuttle Cursortype = 118
	// sizing
	GdkSizing Cursortype = 120
	// spider
	GdkSpider Cursortype = 122
	// spraycan
	GdkSpraycan Cursortype = 124
	// star
	GdkStar Cursortype = 126
	// target
	GdkTarget Cursortype = 128
	// tcross
	GdkTcross Cursortype = 130
	// top_left_arrow
	GdkTopLeftArrow Cursortype = 132
	// top_left_corner
	GdkTopLeftCorner Cursortype = 134
	// top_right_corner
	GdkTopRightCorner Cursortype = 136
	// top_side
	GdkTopSide Cursortype = 138
	// top_tee
	GdkTopTee Cursortype = 140
	// trek
	GdkTrek Cursortype = 142
	// ul_angle
	GdkUlAngle Cursortype = 144
	// umbrella
	GdkUmbrella Cursortype = 146
	// ur_angle
	GdkUrAngle Cursortype = 148
	// watch
	GdkWatch Cursortype = 150
	// xterm
	GdkXterm Cursortype = 152
	// last_cursor
	GdkLastCursor Cursortype = 153
	// blank_cursor
	GdkBlankCursor Cursortype = -2
	// cursor_is_pixmap
	GdkCursorIsPixmap Cursortype = -1
)

// Devicepadfeature is a representation of the C type DevicePadFeature.
type Devicepadfeature int

const (
	// button
	GdkDevicePadFeatureButton Devicepadfeature = 0
	// ring
	GdkDevicePadFeatureRing Devicepadfeature = 1
	// strip
	GdkDevicePadFeatureStrip Devicepadfeature = 2
)

// Devicetooltype is a representation of the C type DeviceToolType.
//
// since 3.22
type Devicetooltype int

const (
	// unknown
	GdkDeviceToolTypeUnknown Devicetooltype = 0
	// pen
	GdkDeviceToolTypePen Devicetooltype = 1
	// eraser
	GdkDeviceToolTypeEraser Devicetooltype = 2
	// brush
	GdkDeviceToolTypeBrush Devicetooltype = 3
	// pencil
	GdkDeviceToolTypePencil Devicetooltype = 4
	// airbrush
	GdkDeviceToolTypeAirbrush Devicetooltype = 5
	// mouse
	GdkDeviceToolTypeMouse Devicetooltype = 6
	// lens
	GdkDeviceToolTypeLens Devicetooltype = 7
)

// Devicetype is a representation of the C type DeviceType.
type Devicetype int

const (
	// master
	GdkDeviceTypeMaster Devicetype = 0
	// slave
	GdkDeviceTypeSlave Devicetype = 1
	// floating
	GdkDeviceTypeFloating Devicetype = 2
)

// Dragcancelreason is a representation of the C type DragCancelReason.
//
// since 3.20
type Dragcancelreason int

const (
	// no_target
	GdkDragCancelNoTarget Dragcancelreason = 0
	// user_cancelled
	GdkDragCancelUserCancelled Dragcancelreason = 1
	// error
	GdkDragCancelError Dragcancelreason = 2
)

// Dragprotocol is a representation of the C type DragProtocol.
type Dragprotocol int

const (
	// none
	GdkDragProtoNone Dragprotocol = 0
	// motif
	GdkDragProtoMotif Dragprotocol = 1
	// xdnd
	GdkDragProtoXdnd Dragprotocol = 2
	// rootwin
	GdkDragProtoRootwin Dragprotocol = 3
	// win32_dropfiles
	GdkDragProtoWin32Dropfiles Dragprotocol = 4
	// ole2
	GdkDragProtoOle2 Dragprotocol = 5
	// local
	GdkDragProtoLocal Dragprotocol = 6
	// wayland
	GdkDragProtoWayland Dragprotocol = 7
)

// Eventtype is a representation of the C type EventType.
type Eventtype int

const (
	// nothing
	GdkNothing Eventtype = -1
	// delete
	GdkDelete Eventtype = 0
	// destroy
	GdkDestroy Eventtype = 1
	// expose
	GdkExpose Eventtype = 2
	// motion_notify
	GdkMotionNotify Eventtype = 3
	// button_press
	GdkButtonPress Eventtype = 4
	// 2button_press
	Gdk2buttonPress Eventtype = 5
	// double_button_press
	GdkDoubleButtonPress Eventtype = 5
	// 3button_press
	Gdk3buttonPress Eventtype = 6
	// triple_button_press
	GdkTripleButtonPress Eventtype = 6
	// button_release
	GdkButtonRelease Eventtype = 7
	// key_press
	GdkKeyPress Eventtype = 8
	// key_release
	GdkKeyRelease Eventtype = 9
	// enter_notify
	GdkEnterNotify Eventtype = 10
	// leave_notify
	GdkLeaveNotify Eventtype = 11
	// focus_change
	GdkFocusChange Eventtype = 12
	// configure
	GdkConfigure Eventtype = 13
	// map
	GdkMap Eventtype = 14
	// unmap
	GdkUnmap Eventtype = 15
	// property_notify
	GdkPropertyNotify Eventtype = 16
	// selection_clear
	GdkSelectionClear Eventtype = 17
	// selection_request
	GdkSelectionRequest Eventtype = 18
	// selection_notify
	GdkSelectionNotify Eventtype = 19
	// proximity_in
	GdkProximityIn Eventtype = 20
	// proximity_out
	GdkProximityOut Eventtype = 21
	// drag_enter
	GdkDragEnter Eventtype = 22
	// drag_leave
	GdkDragLeave Eventtype = 23
	// drag_motion
	GdkDragMotion Eventtype = 24
	// drag_status
	GdkDragStatus Eventtype = 25
	// drop_start
	GdkDropStart Eventtype = 26
	// drop_finished
	GdkDropFinished Eventtype = 27
	// client_event
	GdkClientEvent Eventtype = 28
	// visibility_notify
	GdkVisibilityNotify Eventtype = 29
	// scroll
	GdkScroll Eventtype = 31
	// window_state
	GdkWindowState Eventtype = 32
	// setting
	GdkSetting Eventtype = 33
	// owner_change
	GdkOwnerChange Eventtype = 34
	// grab_broken
	GdkGrabBroken Eventtype = 35
	// damage
	GdkDamage Eventtype = 36
	// touch_begin
	GdkTouchBegin Eventtype = 37
	// touch_update
	GdkTouchUpdate Eventtype = 38
	// touch_end
	GdkTouchEnd Eventtype = 39
	// touch_cancel
	GdkTouchCancel Eventtype = 40
	// touchpad_swipe
	GdkTouchpadSwipe Eventtype = 41
	// touchpad_pinch
	GdkTouchpadPinch Eventtype = 42
	// pad_button_press
	GdkPadButtonPress Eventtype = 43
	// pad_button_release
	GdkPadButtonRelease Eventtype = 44
	// pad_ring
	GdkPadRing Eventtype = 45
	// pad_strip
	GdkPadStrip Eventtype = 46
	// pad_group_mode
	GdkPadGroupMode Eventtype = 47
	// event_last
	GdkEventLast Eventtype = 48
)

// Filterreturn is a representation of the C type FilterReturn.
type Filterreturn int

const (
	// continue
	GdkFilterContinue Filterreturn = 0
	// translate
	GdkFilterTranslate Filterreturn = 1
	// remove
	GdkFilterRemove Filterreturn = 2
)

// Fullscreenmode is a representation of the C type FullscreenMode.
//
// since 3.8
type Fullscreenmode int

const (
	// current_monitor
	GdkFullscreenOnCurrentMonitor Fullscreenmode = 0
	// all_monitors
	GdkFullscreenOnAllMonitors Fullscreenmode = 1
)

// Glerror is a representation of the C type GLError.
//
// since 3.16
type Glerror int

const (
	// not_available
	GdkGlErrorNotAvailable Glerror = 0
	// unsupported_format
	GdkGlErrorUnsupportedFormat Glerror = 1
	// unsupported_profile
	GdkGlErrorUnsupportedProfile Glerror = 2
)

// Grabownership is a representation of the C type GrabOwnership.
type Grabownership int

const (
	// none
	GdkOwnershipNone Grabownership = 0
	// window
	GdkOwnershipWindow Grabownership = 1
	// application
	GdkOwnershipApplication Grabownership = 2
)

// Grabstatus is a representation of the C type GrabStatus.
type Grabstatus int

const (
	// success
	GdkGrabSuccess Grabstatus = 0
	// already_grabbed
	GdkGrabAlreadyGrabbed Grabstatus = 1
	// invalid_time
	GdkGrabInvalidTime Grabstatus = 2
	// not_viewable
	GdkGrabNotViewable Grabstatus = 3
	// frozen
	GdkGrabFrozen Grabstatus = 4
	// failed
	GdkGrabFailed Grabstatus = 5
)

// Gravity is a representation of the C type Gravity.
type Gravity int

const (
	// north_west
	GdkGravityNorthWest Gravity = 1
	// north
	GdkGravityNorth Gravity = 2
	// north_east
	GdkGravityNorthEast Gravity = 3
	// west
	GdkGravityWest Gravity = 4
	// center
	GdkGravityCenter Gravity = 5
	// east
	GdkGravityEast Gravity = 6
	// south_west
	GdkGravitySouthWest Gravity = 7
	// south
	GdkGravitySouth Gravity = 8
	// south_east
	GdkGravitySouthEast Gravity = 9
	// static
	GdkGravityStatic Gravity = 10
)

// Inputmode is a representation of the C type InputMode.
type Inputmode int

const (
	// disabled
	GdkModeDisabled Inputmode = 0
	// screen
	GdkModeScreen Inputmode = 1
	// window
	GdkModeWindow Inputmode = 2
)

// Inputsource is a representation of the C type InputSource.
type Inputsource int

const (
	// mouse
	GdkSourceMouse Inputsource = 0
	// pen
	GdkSourcePen Inputsource = 1
	// eraser
	GdkSourceEraser Inputsource = 2
	// cursor
	GdkSourceCursor Inputsource = 3
	// keyboard
	GdkSourceKeyboard Inputsource = 4
	// touchscreen
	GdkSourceTouchscreen Inputsource = 5
	// touchpad
	GdkSourceTouchpad Inputsource = 6
	// trackpoint
	GdkSourceTrackpoint Inputsource = 7
	// tablet_pad
	GdkSourceTabletPad Inputsource = 8
)

// Modifierintent is a representation of the C type ModifierIntent.
//
// since 3.4
type Modifierintent int

const (
	// primary_accelerator
	GdkModifierIntentPrimaryAccelerator Modifierintent = 0
	// context_menu
	GdkModifierIntentContextMenu Modifierintent = 1
	// extend_selection
	GdkModifierIntentExtendSelection Modifierintent = 2
	// modify_selection
	GdkModifierIntentModifySelection Modifierintent = 3
	// no_text_input
	GdkModifierIntentNoTextInput Modifierintent = 4
	// shift_group
	GdkModifierIntentShiftGroup Modifierintent = 5
	// default_mod_mask
	GdkModifierIntentDefaultModMask Modifierintent = 6
)

// Notifytype is a representation of the C type NotifyType.
type Notifytype int

const (
	// ancestor
	GdkNotifyAncestor Notifytype = 0
	// virtual
	GdkNotifyVirtual Notifytype = 1
	// inferior
	GdkNotifyInferior Notifytype = 2
	// nonlinear
	GdkNotifyNonlinear Notifytype = 3
	// nonlinear_virtual
	GdkNotifyNonlinearVirtual Notifytype = 4
	// unknown
	GdkNotifyUnknown Notifytype = 5
)

// Ownerchange is a representation of the C type OwnerChange.
type Ownerchange int

const (
	// new_owner
	GdkOwnerChangeNewOwner Ownerchange = 0
	// destroy
	GdkOwnerChangeDestroy Ownerchange = 1
	// close
	GdkOwnerChangeClose Ownerchange = 2
)

// Propmode is a representation of the C type PropMode.
type Propmode int

const (
	// replace
	GdkPropModeReplace Propmode = 0
	// prepend
	GdkPropModePrepend Propmode = 1
	// append
	GdkPropModeAppend Propmode = 2
)

// Propertystate is a representation of the C type PropertyState.
type Propertystate int

const (
	// new_value
	GdkPropertyNewValue Propertystate = 0
	// delete
	GdkPropertyDelete Propertystate = 1
)

// Scrolldirection is a representation of the C type ScrollDirection.
type Scrolldirection int

const (
	// up
	GdkScrollUp Scrolldirection = 0
	// down
	GdkScrollDown Scrolldirection = 1
	// left
	GdkScrollLeft Scrolldirection = 2
	// right
	GdkScrollRight Scrolldirection = 3
	// smooth
	GdkScrollSmooth Scrolldirection = 4
)

// Settingaction is a representation of the C type SettingAction.
type Settingaction int

const (
	// new
	GdkSettingActionNew Settingaction = 0
	// changed
	GdkSettingActionChanged Settingaction = 1
	// deleted
	GdkSettingActionDeleted Settingaction = 2
)

// Status is a representation of the C type Status.
type Status int

const (
	// ok
	GdkOk Status = 0
	// error
	GdkError Status = -1
	// error_param
	GdkErrorParam Status = -2
	// error_file
	GdkErrorFile Status = -3
	// error_mem
	GdkErrorMem Status = -4
)

// Subpixellayout is a representation of the C type SubpixelLayout.
//
// since 3.22
type Subpixellayout int

const (
	// unknown
	GdkSubpixelLayoutUnknown Subpixellayout = 0
	// none
	GdkSubpixelLayoutNone Subpixellayout = 1
	// horizontal_rgb
	GdkSubpixelLayoutHorizontalRgb Subpixellayout = 2
	// horizontal_bgr
	GdkSubpixelLayoutHorizontalBgr Subpixellayout = 3
	// vertical_rgb
	GdkSubpixelLayoutVerticalRgb Subpixellayout = 4
	// vertical_bgr
	GdkSubpixelLayoutVerticalBgr Subpixellayout = 5
)

// Touchpadgesturephase is a representation of the C type TouchpadGesturePhase.
type Touchpadgesturephase int

const (
	// begin
	GdkTouchpadGesturePhaseBegin Touchpadgesturephase = 0
	// update
	GdkTouchpadGesturePhaseUpdate Touchpadgesturephase = 1
	// end
	GdkTouchpadGesturePhaseEnd Touchpadgesturephase = 2
	// cancel
	GdkTouchpadGesturePhaseCancel Touchpadgesturephase = 3
)

// Visibilitystate is a representation of the C type VisibilityState.
type Visibilitystate int

const (
	// unobscured
	GdkVisibilityUnobscured Visibilitystate = 0
	// partial
	GdkVisibilityPartial Visibilitystate = 1
	// fully_obscured
	GdkVisibilityFullyObscured Visibilitystate = 2
)

// Visualtype is a representation of the C type VisualType.
type Visualtype int

const (
	// static_gray
	GdkVisualStaticGray Visualtype = 0
	// grayscale
	GdkVisualGrayscale Visualtype = 1
	// static_color
	GdkVisualStaticColor Visualtype = 2
	// pseudo_color
	GdkVisualPseudoColor Visualtype = 3
	// true_color
	GdkVisualTrueColor Visualtype = 4
	// direct_color
	GdkVisualDirectColor Visualtype = 5
)

// Windowedge is a representation of the C type WindowEdge.
type Windowedge int

const (
	// north_west
	GdkWindowEdgeNorthWest Windowedge = 0
	// north
	GdkWindowEdgeNorth Windowedge = 1
	// north_east
	GdkWindowEdgeNorthEast Windowedge = 2
	// west
	GdkWindowEdgeWest Windowedge = 3
	// east
	GdkWindowEdgeEast Windowedge = 4
	// south_west
	GdkWindowEdgeSouthWest Windowedge = 5
	// south
	GdkWindowEdgeSouth Windowedge = 6
	// south_east
	GdkWindowEdgeSouthEast Windowedge = 7
)

// Windowtype is a representation of the C type WindowType.
type Windowtype int

const (
	// root
	GdkWindowRoot Windowtype = 0
	// toplevel
	GdkWindowToplevel Windowtype = 1
	// child
	GdkWindowChild Windowtype = 2
	// temp
	GdkWindowTemp Windowtype = 3
	// foreign
	GdkWindowForeign Windowtype = 4
	// offscreen
	GdkWindowOffscreen Windowtype = 5
	// subsurface
	GdkWindowSubsurface Windowtype = 6
)

// Windowtypehint is a representation of the C type WindowTypeHint.
type Windowtypehint int

const (
	// normal
	GdkWindowTypeHintNormal Windowtypehint = 0
	// dialog
	GdkWindowTypeHintDialog Windowtypehint = 1
	// menu
	GdkWindowTypeHintMenu Windowtypehint = 2
	// toolbar
	GdkWindowTypeHintToolbar Windowtypehint = 3
	// splashscreen
	GdkWindowTypeHintSplashscreen Windowtypehint = 4
	// utility
	GdkWindowTypeHintUtility Windowtypehint = 5
	// dock
	GdkWindowTypeHintDock Windowtypehint = 6
	// desktop
	GdkWindowTypeHintDesktop Windowtypehint = 7
	// dropdown_menu
	GdkWindowTypeHintDropdownMenu Windowtypehint = 8
	// popup_menu
	GdkWindowTypeHintPopupMenu Windowtypehint = 9
	// tooltip
	GdkWindowTypeHintTooltip Windowtypehint = 10
	// notification
	GdkWindowTypeHintNotification Windowtypehint = 11
	// combo
	GdkWindowTypeHintCombo Windowtypehint = 12
	// dnd
	GdkWindowTypeHintDnd Windowtypehint = 13
)

// Windowwindowclass is a representation of the C type WindowWindowClass.
type Windowwindowclass int

const (
	// input_output
	GdkInputOutput Windowwindowclass = 0
	// input_only
	GdkInputOnly Windowwindowclass = 1
)
