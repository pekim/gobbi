// Code generated - DO NOT EDIT.

package gdk

// Axisuse is a representation of the C type GdkAxisUse.
type Axisuse int

const (
	// Axisuse_Ignore is a representation of the C type GDK_AXIS_IGNORE.
	Axisuse_Ignore Axisuse = 0
	// Axisuse_X is a representation of the C type GDK_AXIS_X.
	Axisuse_X Axisuse = 1
	// Axisuse_Y is a representation of the C type GDK_AXIS_Y.
	Axisuse_Y Axisuse = 2
	// Axisuse_Pressure is a representation of the C type GDK_AXIS_PRESSURE.
	Axisuse_Pressure Axisuse = 3
	// Axisuse_Xtilt is a representation of the C type GDK_AXIS_XTILT.
	Axisuse_Xtilt Axisuse = 4
	// Axisuse_Ytilt is a representation of the C type GDK_AXIS_YTILT.
	Axisuse_Ytilt Axisuse = 5
	// Axisuse_Wheel is a representation of the C type GDK_AXIS_WHEEL.
	Axisuse_Wheel Axisuse = 6
	// Axisuse_Distance is a representation of the C type GDK_AXIS_DISTANCE.
	Axisuse_Distance Axisuse = 7
	// Axisuse_Rotation is a representation of the C type GDK_AXIS_ROTATION.
	Axisuse_Rotation Axisuse = 8
	// Axisuse_Slider is a representation of the C type GDK_AXIS_SLIDER.
	Axisuse_Slider Axisuse = 9
	// Axisuse_Last is a representation of the C type GDK_AXIS_LAST.
	Axisuse_Last Axisuse = 10
)

// Byteorder is a representation of the C type GdkByteOrder.
type Byteorder int

const (
	// Byteorder_LsbFirst is a representation of the C type GDK_LSB_FIRST.
	Byteorder_LsbFirst Byteorder = 0
	// Byteorder_MsbFirst is a representation of the C type GDK_MSB_FIRST.
	Byteorder_MsbFirst Byteorder = 1
)

// Crossingmode is a representation of the C type GdkCrossingMode.
type Crossingmode int

const (
	// Crossingmode_Normal is a representation of the C type GDK_CROSSING_NORMAL.
	Crossingmode_Normal Crossingmode = 0
	// Crossingmode_Grab is a representation of the C type GDK_CROSSING_GRAB.
	Crossingmode_Grab Crossingmode = 1
	// Crossingmode_Ungrab is a representation of the C type GDK_CROSSING_UNGRAB.
	Crossingmode_Ungrab Crossingmode = 2
	// Crossingmode_GtkGrab is a representation of the C type GDK_CROSSING_GTK_GRAB.
	Crossingmode_GtkGrab Crossingmode = 3
	// Crossingmode_GtkUngrab is a representation of the C type GDK_CROSSING_GTK_UNGRAB.
	Crossingmode_GtkUngrab Crossingmode = 4
	// Crossingmode_StateChanged is a representation of the C type GDK_CROSSING_STATE_CHANGED.
	Crossingmode_StateChanged Crossingmode = 5
	// Crossingmode_TouchBegin is a representation of the C type GDK_CROSSING_TOUCH_BEGIN.
	Crossingmode_TouchBegin Crossingmode = 6
	// Crossingmode_TouchEnd is a representation of the C type GDK_CROSSING_TOUCH_END.
	Crossingmode_TouchEnd Crossingmode = 7
	// Crossingmode_DeviceSwitch is a representation of the C type GDK_CROSSING_DEVICE_SWITCH.
	Crossingmode_DeviceSwitch Crossingmode = 8
)

// Cursortype is a representation of the C type GdkCursorType.
type Cursortype int

const (
	// Cursortype_XCursor is a representation of the C type GDK_X_CURSOR.
	Cursortype_XCursor Cursortype = 0
	// Cursortype_Arrow is a representation of the C type GDK_ARROW.
	Cursortype_Arrow Cursortype = 2
	// Cursortype_BasedArrowDown is a representation of the C type GDK_BASED_ARROW_DOWN.
	Cursortype_BasedArrowDown Cursortype = 4
	// Cursortype_BasedArrowUp is a representation of the C type GDK_BASED_ARROW_UP.
	Cursortype_BasedArrowUp Cursortype = 6
	// Cursortype_Boat is a representation of the C type GDK_BOAT.
	Cursortype_Boat Cursortype = 8
	// Cursortype_Bogosity is a representation of the C type GDK_BOGOSITY.
	Cursortype_Bogosity Cursortype = 10
	// Cursortype_BottomLeftCorner is a representation of the C type GDK_BOTTOM_LEFT_CORNER.
	Cursortype_BottomLeftCorner Cursortype = 12
	// Cursortype_BottomRightCorner is a representation of the C type GDK_BOTTOM_RIGHT_CORNER.
	Cursortype_BottomRightCorner Cursortype = 14
	// Cursortype_BottomSide is a representation of the C type GDK_BOTTOM_SIDE.
	Cursortype_BottomSide Cursortype = 16
	// Cursortype_BottomTee is a representation of the C type GDK_BOTTOM_TEE.
	Cursortype_BottomTee Cursortype = 18
	// Cursortype_BoxSpiral is a representation of the C type GDK_BOX_SPIRAL.
	Cursortype_BoxSpiral Cursortype = 20
	// Cursortype_CenterPtr is a representation of the C type GDK_CENTER_PTR.
	Cursortype_CenterPtr Cursortype = 22
	// Cursortype_Circle is a representation of the C type GDK_CIRCLE.
	Cursortype_Circle Cursortype = 24
	// Cursortype_Clock is a representation of the C type GDK_CLOCK.
	Cursortype_Clock Cursortype = 26
	// Cursortype_CoffeeMug is a representation of the C type GDK_COFFEE_MUG.
	Cursortype_CoffeeMug Cursortype = 28
	// Cursortype_Cross is a representation of the C type GDK_CROSS.
	Cursortype_Cross Cursortype = 30
	// Cursortype_CrossReverse is a representation of the C type GDK_CROSS_REVERSE.
	Cursortype_CrossReverse Cursortype = 32
	// Cursortype_Crosshair is a representation of the C type GDK_CROSSHAIR.
	Cursortype_Crosshair Cursortype = 34
	// Cursortype_DiamondCross is a representation of the C type GDK_DIAMOND_CROSS.
	Cursortype_DiamondCross Cursortype = 36
	// Cursortype_Dot is a representation of the C type GDK_DOT.
	Cursortype_Dot Cursortype = 38
	// Cursortype_Dotbox is a representation of the C type GDK_DOTBOX.
	Cursortype_Dotbox Cursortype = 40
	// Cursortype_DoubleArrow is a representation of the C type GDK_DOUBLE_ARROW.
	Cursortype_DoubleArrow Cursortype = 42
	// Cursortype_DraftLarge is a representation of the C type GDK_DRAFT_LARGE.
	Cursortype_DraftLarge Cursortype = 44
	// Cursortype_DraftSmall is a representation of the C type GDK_DRAFT_SMALL.
	Cursortype_DraftSmall Cursortype = 46
	// Cursortype_DrapedBox is a representation of the C type GDK_DRAPED_BOX.
	Cursortype_DrapedBox Cursortype = 48
	// Cursortype_Exchange is a representation of the C type GDK_EXCHANGE.
	Cursortype_Exchange Cursortype = 50
	// Cursortype_Fleur is a representation of the C type GDK_FLEUR.
	Cursortype_Fleur Cursortype = 52
	// Cursortype_Gobbler is a representation of the C type GDK_GOBBLER.
	Cursortype_Gobbler Cursortype = 54
	// Cursortype_Gumby is a representation of the C type GDK_GUMBY.
	Cursortype_Gumby Cursortype = 56
	// Cursortype_Hand1 is a representation of the C type GDK_HAND1.
	Cursortype_Hand1 Cursortype = 58
	// Cursortype_Hand2 is a representation of the C type GDK_HAND2.
	Cursortype_Hand2 Cursortype = 60
	// Cursortype_Heart is a representation of the C type GDK_HEART.
	Cursortype_Heart Cursortype = 62
	// Cursortype_Icon is a representation of the C type GDK_ICON.
	Cursortype_Icon Cursortype = 64
	// Cursortype_IronCross is a representation of the C type GDK_IRON_CROSS.
	Cursortype_IronCross Cursortype = 66
	// Cursortype_LeftPtr is a representation of the C type GDK_LEFT_PTR.
	Cursortype_LeftPtr Cursortype = 68
	// Cursortype_LeftSide is a representation of the C type GDK_LEFT_SIDE.
	Cursortype_LeftSide Cursortype = 70
	// Cursortype_LeftTee is a representation of the C type GDK_LEFT_TEE.
	Cursortype_LeftTee Cursortype = 72
	// Cursortype_Leftbutton is a representation of the C type GDK_LEFTBUTTON.
	Cursortype_Leftbutton Cursortype = 74
	// Cursortype_LlAngle is a representation of the C type GDK_LL_ANGLE.
	Cursortype_LlAngle Cursortype = 76
	// Cursortype_LrAngle is a representation of the C type GDK_LR_ANGLE.
	Cursortype_LrAngle Cursortype = 78
	// Cursortype_Man is a representation of the C type GDK_MAN.
	Cursortype_Man Cursortype = 80
	// Cursortype_Middlebutton is a representation of the C type GDK_MIDDLEBUTTON.
	Cursortype_Middlebutton Cursortype = 82
	// Cursortype_Mouse is a representation of the C type GDK_MOUSE.
	Cursortype_Mouse Cursortype = 84
	// Cursortype_Pencil is a representation of the C type GDK_PENCIL.
	Cursortype_Pencil Cursortype = 86
	// Cursortype_Pirate is a representation of the C type GDK_PIRATE.
	Cursortype_Pirate Cursortype = 88
	// Cursortype_Plus is a representation of the C type GDK_PLUS.
	Cursortype_Plus Cursortype = 90
	// Cursortype_QuestionArrow is a representation of the C type GDK_QUESTION_ARROW.
	Cursortype_QuestionArrow Cursortype = 92
	// Cursortype_RightPtr is a representation of the C type GDK_RIGHT_PTR.
	Cursortype_RightPtr Cursortype = 94
	// Cursortype_RightSide is a representation of the C type GDK_RIGHT_SIDE.
	Cursortype_RightSide Cursortype = 96
	// Cursortype_RightTee is a representation of the C type GDK_RIGHT_TEE.
	Cursortype_RightTee Cursortype = 98
	// Cursortype_Rightbutton is a representation of the C type GDK_RIGHTBUTTON.
	Cursortype_Rightbutton Cursortype = 100
	// Cursortype_RtlLogo is a representation of the C type GDK_RTL_LOGO.
	Cursortype_RtlLogo Cursortype = 102
	// Cursortype_Sailboat is a representation of the C type GDK_SAILBOAT.
	Cursortype_Sailboat Cursortype = 104
	// Cursortype_SbDownArrow is a representation of the C type GDK_SB_DOWN_ARROW.
	Cursortype_SbDownArrow Cursortype = 106
	// Cursortype_SbHDoubleArrow is a representation of the C type GDK_SB_H_DOUBLE_ARROW.
	Cursortype_SbHDoubleArrow Cursortype = 108
	// Cursortype_SbLeftArrow is a representation of the C type GDK_SB_LEFT_ARROW.
	Cursortype_SbLeftArrow Cursortype = 110
	// Cursortype_SbRightArrow is a representation of the C type GDK_SB_RIGHT_ARROW.
	Cursortype_SbRightArrow Cursortype = 112
	// Cursortype_SbUpArrow is a representation of the C type GDK_SB_UP_ARROW.
	Cursortype_SbUpArrow Cursortype = 114
	// Cursortype_SbVDoubleArrow is a representation of the C type GDK_SB_V_DOUBLE_ARROW.
	Cursortype_SbVDoubleArrow Cursortype = 116
	// Cursortype_Shuttle is a representation of the C type GDK_SHUTTLE.
	Cursortype_Shuttle Cursortype = 118
	// Cursortype_Sizing is a representation of the C type GDK_SIZING.
	Cursortype_Sizing Cursortype = 120
	// Cursortype_Spider is a representation of the C type GDK_SPIDER.
	Cursortype_Spider Cursortype = 122
	// Cursortype_Spraycan is a representation of the C type GDK_SPRAYCAN.
	Cursortype_Spraycan Cursortype = 124
	// Cursortype_Star is a representation of the C type GDK_STAR.
	Cursortype_Star Cursortype = 126
	// Cursortype_Target is a representation of the C type GDK_TARGET.
	Cursortype_Target Cursortype = 128
	// Cursortype_Tcross is a representation of the C type GDK_TCROSS.
	Cursortype_Tcross Cursortype = 130
	// Cursortype_TopLeftArrow is a representation of the C type GDK_TOP_LEFT_ARROW.
	Cursortype_TopLeftArrow Cursortype = 132
	// Cursortype_TopLeftCorner is a representation of the C type GDK_TOP_LEFT_CORNER.
	Cursortype_TopLeftCorner Cursortype = 134
	// Cursortype_TopRightCorner is a representation of the C type GDK_TOP_RIGHT_CORNER.
	Cursortype_TopRightCorner Cursortype = 136
	// Cursortype_TopSide is a representation of the C type GDK_TOP_SIDE.
	Cursortype_TopSide Cursortype = 138
	// Cursortype_TopTee is a representation of the C type GDK_TOP_TEE.
	Cursortype_TopTee Cursortype = 140
	// Cursortype_Trek is a representation of the C type GDK_TREK.
	Cursortype_Trek Cursortype = 142
	// Cursortype_UlAngle is a representation of the C type GDK_UL_ANGLE.
	Cursortype_UlAngle Cursortype = 144
	// Cursortype_Umbrella is a representation of the C type GDK_UMBRELLA.
	Cursortype_Umbrella Cursortype = 146
	// Cursortype_UrAngle is a representation of the C type GDK_UR_ANGLE.
	Cursortype_UrAngle Cursortype = 148
	// Cursortype_Watch is a representation of the C type GDK_WATCH.
	Cursortype_Watch Cursortype = 150
	// Cursortype_Xterm is a representation of the C type GDK_XTERM.
	Cursortype_Xterm Cursortype = 152
	// Cursortype_LastCursor is a representation of the C type GDK_LAST_CURSOR.
	Cursortype_LastCursor Cursortype = 153
	// Cursortype_BlankCursor is a representation of the C type GDK_BLANK_CURSOR.
	Cursortype_BlankCursor Cursortype = -2
	// Cursortype_CursorIsPixmap is a representation of the C type GDK_CURSOR_IS_PIXMAP.
	Cursortype_CursorIsPixmap Cursortype = -1
)

// Devicepadfeature is a representation of the C type GdkDevicePadFeature.
type Devicepadfeature int

const (
	// Devicepadfeature_Button is a representation of the C type GDK_DEVICE_PAD_FEATURE_BUTTON.
	Devicepadfeature_Button Devicepadfeature = 0
	// Devicepadfeature_Ring is a representation of the C type GDK_DEVICE_PAD_FEATURE_RING.
	Devicepadfeature_Ring Devicepadfeature = 1
	// Devicepadfeature_Strip is a representation of the C type GDK_DEVICE_PAD_FEATURE_STRIP.
	Devicepadfeature_Strip Devicepadfeature = 2
)

// Devicetooltype is a representation of the C type GdkDeviceToolType.
//
// since 3.22
type Devicetooltype int

const (
	// Devicetooltype_Unknown is a representation of the C type GDK_DEVICE_TOOL_TYPE_UNKNOWN.
	Devicetooltype_Unknown Devicetooltype = 0
	// Devicetooltype_Pen is a representation of the C type GDK_DEVICE_TOOL_TYPE_PEN.
	Devicetooltype_Pen Devicetooltype = 1
	// Devicetooltype_Eraser is a representation of the C type GDK_DEVICE_TOOL_TYPE_ERASER.
	Devicetooltype_Eraser Devicetooltype = 2
	// Devicetooltype_Brush is a representation of the C type GDK_DEVICE_TOOL_TYPE_BRUSH.
	Devicetooltype_Brush Devicetooltype = 3
	// Devicetooltype_Pencil is a representation of the C type GDK_DEVICE_TOOL_TYPE_PENCIL.
	Devicetooltype_Pencil Devicetooltype = 4
	// Devicetooltype_Airbrush is a representation of the C type GDK_DEVICE_TOOL_TYPE_AIRBRUSH.
	Devicetooltype_Airbrush Devicetooltype = 5
	// Devicetooltype_Mouse is a representation of the C type GDK_DEVICE_TOOL_TYPE_MOUSE.
	Devicetooltype_Mouse Devicetooltype = 6
	// Devicetooltype_Lens is a representation of the C type GDK_DEVICE_TOOL_TYPE_LENS.
	Devicetooltype_Lens Devicetooltype = 7
)

// Devicetype is a representation of the C type GdkDeviceType.
type Devicetype int

const (
	// Devicetype_Master is a representation of the C type GDK_DEVICE_TYPE_MASTER.
	Devicetype_Master Devicetype = 0
	// Devicetype_Slave is a representation of the C type GDK_DEVICE_TYPE_SLAVE.
	Devicetype_Slave Devicetype = 1
	// Devicetype_Floating is a representation of the C type GDK_DEVICE_TYPE_FLOATING.
	Devicetype_Floating Devicetype = 2
)

// Dragcancelreason is a representation of the C type GdkDragCancelReason.
//
// since 3.20
type Dragcancelreason int

const (
	// Dragcancelreason_NoTarget is a representation of the C type GDK_DRAG_CANCEL_NO_TARGET.
	Dragcancelreason_NoTarget Dragcancelreason = 0
	// Dragcancelreason_UserCancelled is a representation of the C type GDK_DRAG_CANCEL_USER_CANCELLED.
	Dragcancelreason_UserCancelled Dragcancelreason = 1
	// Dragcancelreason_Error is a representation of the C type GDK_DRAG_CANCEL_ERROR.
	Dragcancelreason_Error Dragcancelreason = 2
)

// Dragprotocol is a representation of the C type GdkDragProtocol.
type Dragprotocol int

const (
	// Dragprotocol_None is a representation of the C type GDK_DRAG_PROTO_NONE.
	Dragprotocol_None Dragprotocol = 0
	// Dragprotocol_Motif is a representation of the C type GDK_DRAG_PROTO_MOTIF.
	Dragprotocol_Motif Dragprotocol = 1
	// Dragprotocol_Xdnd is a representation of the C type GDK_DRAG_PROTO_XDND.
	Dragprotocol_Xdnd Dragprotocol = 2
	// Dragprotocol_Rootwin is a representation of the C type GDK_DRAG_PROTO_ROOTWIN.
	Dragprotocol_Rootwin Dragprotocol = 3
	// Dragprotocol_Win32Dropfiles is a representation of the C type GDK_DRAG_PROTO_WIN32_DROPFILES.
	Dragprotocol_Win32Dropfiles Dragprotocol = 4
	// Dragprotocol_Ole2 is a representation of the C type GDK_DRAG_PROTO_OLE2.
	Dragprotocol_Ole2 Dragprotocol = 5
	// Dragprotocol_Local is a representation of the C type GDK_DRAG_PROTO_LOCAL.
	Dragprotocol_Local Dragprotocol = 6
	// Dragprotocol_Wayland is a representation of the C type GDK_DRAG_PROTO_WAYLAND.
	Dragprotocol_Wayland Dragprotocol = 7
)

// Eventtype is a representation of the C type GdkEventType.
type Eventtype int

const (
	// Eventtype_Nothing is a representation of the C type GDK_NOTHING.
	Eventtype_Nothing Eventtype = -1
	// Eventtype_Delete is a representation of the C type GDK_DELETE.
	Eventtype_Delete Eventtype = 0
	// Eventtype_Destroy is a representation of the C type GDK_DESTROY.
	Eventtype_Destroy Eventtype = 1
	// Eventtype_Expose is a representation of the C type GDK_EXPOSE.
	Eventtype_Expose Eventtype = 2
	// Eventtype_MotionNotify is a representation of the C type GDK_MOTION_NOTIFY.
	Eventtype_MotionNotify Eventtype = 3
	// Eventtype_ButtonPress is a representation of the C type GDK_BUTTON_PRESS.
	Eventtype_ButtonPress Eventtype = 4
	// Eventtype_2buttonPress is a representation of the C type GDK_2BUTTON_PRESS.
	Eventtype_2buttonPress Eventtype = 5
	// Eventtype_DoubleButtonPress is a representation of the C type GDK_DOUBLE_BUTTON_PRESS.
	Eventtype_DoubleButtonPress Eventtype = 5
	// Eventtype_3buttonPress is a representation of the C type GDK_3BUTTON_PRESS.
	Eventtype_3buttonPress Eventtype = 6
	// Eventtype_TripleButtonPress is a representation of the C type GDK_TRIPLE_BUTTON_PRESS.
	Eventtype_TripleButtonPress Eventtype = 6
	// Eventtype_ButtonRelease is a representation of the C type GDK_BUTTON_RELEASE.
	Eventtype_ButtonRelease Eventtype = 7
	// Eventtype_KeyPress is a representation of the C type GDK_KEY_PRESS.
	Eventtype_KeyPress Eventtype = 8
	// Eventtype_KeyRelease is a representation of the C type GDK_KEY_RELEASE.
	Eventtype_KeyRelease Eventtype = 9
	// Eventtype_EnterNotify is a representation of the C type GDK_ENTER_NOTIFY.
	Eventtype_EnterNotify Eventtype = 10
	// Eventtype_LeaveNotify is a representation of the C type GDK_LEAVE_NOTIFY.
	Eventtype_LeaveNotify Eventtype = 11
	// Eventtype_FocusChange is a representation of the C type GDK_FOCUS_CHANGE.
	Eventtype_FocusChange Eventtype = 12
	// Eventtype_Configure is a representation of the C type GDK_CONFIGURE.
	Eventtype_Configure Eventtype = 13
	// Eventtype_Map is a representation of the C type GDK_MAP.
	Eventtype_Map Eventtype = 14
	// Eventtype_Unmap is a representation of the C type GDK_UNMAP.
	Eventtype_Unmap Eventtype = 15
	// Eventtype_PropertyNotify is a representation of the C type GDK_PROPERTY_NOTIFY.
	Eventtype_PropertyNotify Eventtype = 16
	// Eventtype_SelectionClear is a representation of the C type GDK_SELECTION_CLEAR.
	Eventtype_SelectionClear Eventtype = 17
	// Eventtype_SelectionRequest is a representation of the C type GDK_SELECTION_REQUEST.
	Eventtype_SelectionRequest Eventtype = 18
	// Eventtype_SelectionNotify is a representation of the C type GDK_SELECTION_NOTIFY.
	Eventtype_SelectionNotify Eventtype = 19
	// Eventtype_ProximityIn is a representation of the C type GDK_PROXIMITY_IN.
	Eventtype_ProximityIn Eventtype = 20
	// Eventtype_ProximityOut is a representation of the C type GDK_PROXIMITY_OUT.
	Eventtype_ProximityOut Eventtype = 21
	// Eventtype_DragEnter is a representation of the C type GDK_DRAG_ENTER.
	Eventtype_DragEnter Eventtype = 22
	// Eventtype_DragLeave is a representation of the C type GDK_DRAG_LEAVE.
	Eventtype_DragLeave Eventtype = 23
	// Eventtype_DragMotion is a representation of the C type GDK_DRAG_MOTION.
	Eventtype_DragMotion Eventtype = 24
	// Eventtype_DragStatus is a representation of the C type GDK_DRAG_STATUS.
	Eventtype_DragStatus Eventtype = 25
	// Eventtype_DropStart is a representation of the C type GDK_DROP_START.
	Eventtype_DropStart Eventtype = 26
	// Eventtype_DropFinished is a representation of the C type GDK_DROP_FINISHED.
	Eventtype_DropFinished Eventtype = 27
	// Eventtype_ClientEvent is a representation of the C type GDK_CLIENT_EVENT.
	Eventtype_ClientEvent Eventtype = 28
	// Eventtype_VisibilityNotify is a representation of the C type GDK_VISIBILITY_NOTIFY.
	Eventtype_VisibilityNotify Eventtype = 29
	// Eventtype_Scroll is a representation of the C type GDK_SCROLL.
	Eventtype_Scroll Eventtype = 31
	// Eventtype_WindowState is a representation of the C type GDK_WINDOW_STATE.
	Eventtype_WindowState Eventtype = 32
	// Eventtype_Setting is a representation of the C type GDK_SETTING.
	Eventtype_Setting Eventtype = 33
	// Eventtype_OwnerChange is a representation of the C type GDK_OWNER_CHANGE.
	Eventtype_OwnerChange Eventtype = 34
	// Eventtype_GrabBroken is a representation of the C type GDK_GRAB_BROKEN.
	Eventtype_GrabBroken Eventtype = 35
	// Eventtype_Damage is a representation of the C type GDK_DAMAGE.
	Eventtype_Damage Eventtype = 36
	// Eventtype_TouchBegin is a representation of the C type GDK_TOUCH_BEGIN.
	Eventtype_TouchBegin Eventtype = 37
	// Eventtype_TouchUpdate is a representation of the C type GDK_TOUCH_UPDATE.
	Eventtype_TouchUpdate Eventtype = 38
	// Eventtype_TouchEnd is a representation of the C type GDK_TOUCH_END.
	Eventtype_TouchEnd Eventtype = 39
	// Eventtype_TouchCancel is a representation of the C type GDK_TOUCH_CANCEL.
	Eventtype_TouchCancel Eventtype = 40
	// Eventtype_TouchpadSwipe is a representation of the C type GDK_TOUCHPAD_SWIPE.
	Eventtype_TouchpadSwipe Eventtype = 41
	// Eventtype_TouchpadPinch is a representation of the C type GDK_TOUCHPAD_PINCH.
	Eventtype_TouchpadPinch Eventtype = 42
	// Eventtype_PadButtonPress is a representation of the C type GDK_PAD_BUTTON_PRESS.
	Eventtype_PadButtonPress Eventtype = 43
	// Eventtype_PadButtonRelease is a representation of the C type GDK_PAD_BUTTON_RELEASE.
	Eventtype_PadButtonRelease Eventtype = 44
	// Eventtype_PadRing is a representation of the C type GDK_PAD_RING.
	Eventtype_PadRing Eventtype = 45
	// Eventtype_PadStrip is a representation of the C type GDK_PAD_STRIP.
	Eventtype_PadStrip Eventtype = 46
	// Eventtype_PadGroupMode is a representation of the C type GDK_PAD_GROUP_MODE.
	Eventtype_PadGroupMode Eventtype = 47
	// Eventtype_EventLast is a representation of the C type GDK_EVENT_LAST.
	Eventtype_EventLast Eventtype = 48
)

// Filterreturn is a representation of the C type GdkFilterReturn.
type Filterreturn int

const (
	// Filterreturn_Continue is a representation of the C type GDK_FILTER_CONTINUE.
	Filterreturn_Continue Filterreturn = 0
	// Filterreturn_Translate is a representation of the C type GDK_FILTER_TRANSLATE.
	Filterreturn_Translate Filterreturn = 1
	// Filterreturn_Remove is a representation of the C type GDK_FILTER_REMOVE.
	Filterreturn_Remove Filterreturn = 2
)

// Fullscreenmode is a representation of the C type GdkFullscreenMode.
//
// since 3.8
type Fullscreenmode int

const (
	// Fullscreenmode_CurrentMonitor is a representation of the C type GDK_FULLSCREEN_ON_CURRENT_MONITOR.
	Fullscreenmode_CurrentMonitor Fullscreenmode = 0
	// Fullscreenmode_AllMonitors is a representation of the C type GDK_FULLSCREEN_ON_ALL_MONITORS.
	Fullscreenmode_AllMonitors Fullscreenmode = 1
)

// Glerror is a representation of the C type GdkGLError.
//
// since 3.16
type Glerror int

const (
	// Glerror_NotAvailable is a representation of the C type GDK_GL_ERROR_NOT_AVAILABLE.
	Glerror_NotAvailable Glerror = 0
	// Glerror_UnsupportedFormat is a representation of the C type GDK_GL_ERROR_UNSUPPORTED_FORMAT.
	Glerror_UnsupportedFormat Glerror = 1
	// Glerror_UnsupportedProfile is a representation of the C type GDK_GL_ERROR_UNSUPPORTED_PROFILE.
	Glerror_UnsupportedProfile Glerror = 2
)

// Grabownership is a representation of the C type GdkGrabOwnership.
type Grabownership int

const (
	// Grabownership_None is a representation of the C type GDK_OWNERSHIP_NONE.
	Grabownership_None Grabownership = 0
	// Grabownership_Window is a representation of the C type GDK_OWNERSHIP_WINDOW.
	Grabownership_Window Grabownership = 1
	// Grabownership_Application is a representation of the C type GDK_OWNERSHIP_APPLICATION.
	Grabownership_Application Grabownership = 2
)

// Grabstatus is a representation of the C type GdkGrabStatus.
type Grabstatus int

const (
	// Grabstatus_Success is a representation of the C type GDK_GRAB_SUCCESS.
	Grabstatus_Success Grabstatus = 0
	// Grabstatus_AlreadyGrabbed is a representation of the C type GDK_GRAB_ALREADY_GRABBED.
	Grabstatus_AlreadyGrabbed Grabstatus = 1
	// Grabstatus_InvalidTime is a representation of the C type GDK_GRAB_INVALID_TIME.
	Grabstatus_InvalidTime Grabstatus = 2
	// Grabstatus_NotViewable is a representation of the C type GDK_GRAB_NOT_VIEWABLE.
	Grabstatus_NotViewable Grabstatus = 3
	// Grabstatus_Frozen is a representation of the C type GDK_GRAB_FROZEN.
	Grabstatus_Frozen Grabstatus = 4
	// Grabstatus_Failed is a representation of the C type GDK_GRAB_FAILED.
	Grabstatus_Failed Grabstatus = 5
)

// Gravity is a representation of the C type GdkGravity.
type Gravity int

const (
	// Gravity_NorthWest is a representation of the C type GDK_GRAVITY_NORTH_WEST.
	Gravity_NorthWest Gravity = 1
	// Gravity_North is a representation of the C type GDK_GRAVITY_NORTH.
	Gravity_North Gravity = 2
	// Gravity_NorthEast is a representation of the C type GDK_GRAVITY_NORTH_EAST.
	Gravity_NorthEast Gravity = 3
	// Gravity_West is a representation of the C type GDK_GRAVITY_WEST.
	Gravity_West Gravity = 4
	// Gravity_Center is a representation of the C type GDK_GRAVITY_CENTER.
	Gravity_Center Gravity = 5
	// Gravity_East is a representation of the C type GDK_GRAVITY_EAST.
	Gravity_East Gravity = 6
	// Gravity_SouthWest is a representation of the C type GDK_GRAVITY_SOUTH_WEST.
	Gravity_SouthWest Gravity = 7
	// Gravity_South is a representation of the C type GDK_GRAVITY_SOUTH.
	Gravity_South Gravity = 8
	// Gravity_SouthEast is a representation of the C type GDK_GRAVITY_SOUTH_EAST.
	Gravity_SouthEast Gravity = 9
	// Gravity_Static is a representation of the C type GDK_GRAVITY_STATIC.
	Gravity_Static Gravity = 10
)

// Inputmode is a representation of the C type GdkInputMode.
type Inputmode int

const (
	// Inputmode_Disabled is a representation of the C type GDK_MODE_DISABLED.
	Inputmode_Disabled Inputmode = 0
	// Inputmode_Screen is a representation of the C type GDK_MODE_SCREEN.
	Inputmode_Screen Inputmode = 1
	// Inputmode_Window is a representation of the C type GDK_MODE_WINDOW.
	Inputmode_Window Inputmode = 2
)

// Inputsource is a representation of the C type GdkInputSource.
type Inputsource int

const (
	// Inputsource_Mouse is a representation of the C type GDK_SOURCE_MOUSE.
	Inputsource_Mouse Inputsource = 0
	// Inputsource_Pen is a representation of the C type GDK_SOURCE_PEN.
	Inputsource_Pen Inputsource = 1
	// Inputsource_Eraser is a representation of the C type GDK_SOURCE_ERASER.
	Inputsource_Eraser Inputsource = 2
	// Inputsource_Cursor is a representation of the C type GDK_SOURCE_CURSOR.
	Inputsource_Cursor Inputsource = 3
	// Inputsource_Keyboard is a representation of the C type GDK_SOURCE_KEYBOARD.
	Inputsource_Keyboard Inputsource = 4
	// Inputsource_Touchscreen is a representation of the C type GDK_SOURCE_TOUCHSCREEN.
	Inputsource_Touchscreen Inputsource = 5
	// Inputsource_Touchpad is a representation of the C type GDK_SOURCE_TOUCHPAD.
	Inputsource_Touchpad Inputsource = 6
	// Inputsource_Trackpoint is a representation of the C type GDK_SOURCE_TRACKPOINT.
	Inputsource_Trackpoint Inputsource = 7
	// Inputsource_TabletPad is a representation of the C type GDK_SOURCE_TABLET_PAD.
	Inputsource_TabletPad Inputsource = 8
)

// Modifierintent is a representation of the C type GdkModifierIntent.
//
// since 3.4
type Modifierintent int

const (
	// Modifierintent_PrimaryAccelerator is a representation of the C type GDK_MODIFIER_INTENT_PRIMARY_ACCELERATOR.
	Modifierintent_PrimaryAccelerator Modifierintent = 0
	// Modifierintent_ContextMenu is a representation of the C type GDK_MODIFIER_INTENT_CONTEXT_MENU.
	Modifierintent_ContextMenu Modifierintent = 1
	// Modifierintent_ExtendSelection is a representation of the C type GDK_MODIFIER_INTENT_EXTEND_SELECTION.
	Modifierintent_ExtendSelection Modifierintent = 2
	// Modifierintent_ModifySelection is a representation of the C type GDK_MODIFIER_INTENT_MODIFY_SELECTION.
	Modifierintent_ModifySelection Modifierintent = 3
	// Modifierintent_NoTextInput is a representation of the C type GDK_MODIFIER_INTENT_NO_TEXT_INPUT.
	Modifierintent_NoTextInput Modifierintent = 4
	// Modifierintent_ShiftGroup is a representation of the C type GDK_MODIFIER_INTENT_SHIFT_GROUP.
	Modifierintent_ShiftGroup Modifierintent = 5
	// Modifierintent_DefaultModMask is a representation of the C type GDK_MODIFIER_INTENT_DEFAULT_MOD_MASK.
	Modifierintent_DefaultModMask Modifierintent = 6
)

// Notifytype is a representation of the C type GdkNotifyType.
type Notifytype int

const (
	// Notifytype_Ancestor is a representation of the C type GDK_NOTIFY_ANCESTOR.
	Notifytype_Ancestor Notifytype = 0
	// Notifytype_Virtual is a representation of the C type GDK_NOTIFY_VIRTUAL.
	Notifytype_Virtual Notifytype = 1
	// Notifytype_Inferior is a representation of the C type GDK_NOTIFY_INFERIOR.
	Notifytype_Inferior Notifytype = 2
	// Notifytype_Nonlinear is a representation of the C type GDK_NOTIFY_NONLINEAR.
	Notifytype_Nonlinear Notifytype = 3
	// Notifytype_NonlinearVirtual is a representation of the C type GDK_NOTIFY_NONLINEAR_VIRTUAL.
	Notifytype_NonlinearVirtual Notifytype = 4
	// Notifytype_Unknown is a representation of the C type GDK_NOTIFY_UNKNOWN.
	Notifytype_Unknown Notifytype = 5
)

// Ownerchange is a representation of the C type GdkOwnerChange.
type Ownerchange int

const (
	// Ownerchange_NewOwner is a representation of the C type GDK_OWNER_CHANGE_NEW_OWNER.
	Ownerchange_NewOwner Ownerchange = 0
	// Ownerchange_Destroy is a representation of the C type GDK_OWNER_CHANGE_DESTROY.
	Ownerchange_Destroy Ownerchange = 1
	// Ownerchange_Close is a representation of the C type GDK_OWNER_CHANGE_CLOSE.
	Ownerchange_Close Ownerchange = 2
)

// Propmode is a representation of the C type GdkPropMode.
type Propmode int

const (
	// Propmode_Replace is a representation of the C type GDK_PROP_MODE_REPLACE.
	Propmode_Replace Propmode = 0
	// Propmode_Prepend is a representation of the C type GDK_PROP_MODE_PREPEND.
	Propmode_Prepend Propmode = 1
	// Propmode_Append is a representation of the C type GDK_PROP_MODE_APPEND.
	Propmode_Append Propmode = 2
)

// Propertystate is a representation of the C type GdkPropertyState.
type Propertystate int

const (
	// Propertystate_NewValue is a representation of the C type GDK_PROPERTY_NEW_VALUE.
	Propertystate_NewValue Propertystate = 0
	// Propertystate_Delete is a representation of the C type GDK_PROPERTY_DELETE.
	Propertystate_Delete Propertystate = 1
)

// Scrolldirection is a representation of the C type GdkScrollDirection.
type Scrolldirection int

const (
	// Scrolldirection_Up is a representation of the C type GDK_SCROLL_UP.
	Scrolldirection_Up Scrolldirection = 0
	// Scrolldirection_Down is a representation of the C type GDK_SCROLL_DOWN.
	Scrolldirection_Down Scrolldirection = 1
	// Scrolldirection_Left is a representation of the C type GDK_SCROLL_LEFT.
	Scrolldirection_Left Scrolldirection = 2
	// Scrolldirection_Right is a representation of the C type GDK_SCROLL_RIGHT.
	Scrolldirection_Right Scrolldirection = 3
	// Scrolldirection_Smooth is a representation of the C type GDK_SCROLL_SMOOTH.
	Scrolldirection_Smooth Scrolldirection = 4
)

// Settingaction is a representation of the C type GdkSettingAction.
type Settingaction int

const (
	// Settingaction_New is a representation of the C type GDK_SETTING_ACTION_NEW.
	Settingaction_New Settingaction = 0
	// Settingaction_Changed is a representation of the C type GDK_SETTING_ACTION_CHANGED.
	Settingaction_Changed Settingaction = 1
	// Settingaction_Deleted is a representation of the C type GDK_SETTING_ACTION_DELETED.
	Settingaction_Deleted Settingaction = 2
)

// Status is a representation of the C type GdkStatus.
type Status int

const (
	// Status_Ok is a representation of the C type GDK_OK.
	Status_Ok Status = 0
	// Status_Error is a representation of the C type GDK_ERROR.
	Status_Error Status = -1
	// Status_ErrorParam is a representation of the C type GDK_ERROR_PARAM.
	Status_ErrorParam Status = -2
	// Status_ErrorFile is a representation of the C type GDK_ERROR_FILE.
	Status_ErrorFile Status = -3
	// Status_ErrorMem is a representation of the C type GDK_ERROR_MEM.
	Status_ErrorMem Status = -4
)

// Subpixellayout is a representation of the C type GdkSubpixelLayout.
//
// since 3.22
type Subpixellayout int

const (
	// Subpixellayout_Unknown is a representation of the C type GDK_SUBPIXEL_LAYOUT_UNKNOWN.
	Subpixellayout_Unknown Subpixellayout = 0
	// Subpixellayout_None is a representation of the C type GDK_SUBPIXEL_LAYOUT_NONE.
	Subpixellayout_None Subpixellayout = 1
	// Subpixellayout_HorizontalRgb is a representation of the C type GDK_SUBPIXEL_LAYOUT_HORIZONTAL_RGB.
	Subpixellayout_HorizontalRgb Subpixellayout = 2
	// Subpixellayout_HorizontalBgr is a representation of the C type GDK_SUBPIXEL_LAYOUT_HORIZONTAL_BGR.
	Subpixellayout_HorizontalBgr Subpixellayout = 3
	// Subpixellayout_VerticalRgb is a representation of the C type GDK_SUBPIXEL_LAYOUT_VERTICAL_RGB.
	Subpixellayout_VerticalRgb Subpixellayout = 4
	// Subpixellayout_VerticalBgr is a representation of the C type GDK_SUBPIXEL_LAYOUT_VERTICAL_BGR.
	Subpixellayout_VerticalBgr Subpixellayout = 5
)

// Touchpadgesturephase is a representation of the C type GdkTouchpadGesturePhase.
type Touchpadgesturephase int

const (
	// Touchpadgesturephase_Begin is a representation of the C type GDK_TOUCHPAD_GESTURE_PHASE_BEGIN.
	Touchpadgesturephase_Begin Touchpadgesturephase = 0
	// Touchpadgesturephase_Update is a representation of the C type GDK_TOUCHPAD_GESTURE_PHASE_UPDATE.
	Touchpadgesturephase_Update Touchpadgesturephase = 1
	// Touchpadgesturephase_End is a representation of the C type GDK_TOUCHPAD_GESTURE_PHASE_END.
	Touchpadgesturephase_End Touchpadgesturephase = 2
	// Touchpadgesturephase_Cancel is a representation of the C type GDK_TOUCHPAD_GESTURE_PHASE_CANCEL.
	Touchpadgesturephase_Cancel Touchpadgesturephase = 3
)

// Visibilitystate is a representation of the C type GdkVisibilityState.
type Visibilitystate int

const (
	// Visibilitystate_Unobscured is a representation of the C type GDK_VISIBILITY_UNOBSCURED.
	Visibilitystate_Unobscured Visibilitystate = 0
	// Visibilitystate_Partial is a representation of the C type GDK_VISIBILITY_PARTIAL.
	Visibilitystate_Partial Visibilitystate = 1
	// Visibilitystate_FullyObscured is a representation of the C type GDK_VISIBILITY_FULLY_OBSCURED.
	Visibilitystate_FullyObscured Visibilitystate = 2
)

// Visualtype is a representation of the C type GdkVisualType.
type Visualtype int

const (
	// Visualtype_StaticGray is a representation of the C type GDK_VISUAL_STATIC_GRAY.
	Visualtype_StaticGray Visualtype = 0
	// Visualtype_Grayscale is a representation of the C type GDK_VISUAL_GRAYSCALE.
	Visualtype_Grayscale Visualtype = 1
	// Visualtype_StaticColor is a representation of the C type GDK_VISUAL_STATIC_COLOR.
	Visualtype_StaticColor Visualtype = 2
	// Visualtype_PseudoColor is a representation of the C type GDK_VISUAL_PSEUDO_COLOR.
	Visualtype_PseudoColor Visualtype = 3
	// Visualtype_TrueColor is a representation of the C type GDK_VISUAL_TRUE_COLOR.
	Visualtype_TrueColor Visualtype = 4
	// Visualtype_DirectColor is a representation of the C type GDK_VISUAL_DIRECT_COLOR.
	Visualtype_DirectColor Visualtype = 5
)

// Windowedge is a representation of the C type GdkWindowEdge.
type Windowedge int

const (
	// Windowedge_NorthWest is a representation of the C type GDK_WINDOW_EDGE_NORTH_WEST.
	Windowedge_NorthWest Windowedge = 0
	// Windowedge_North is a representation of the C type GDK_WINDOW_EDGE_NORTH.
	Windowedge_North Windowedge = 1
	// Windowedge_NorthEast is a representation of the C type GDK_WINDOW_EDGE_NORTH_EAST.
	Windowedge_NorthEast Windowedge = 2
	// Windowedge_West is a representation of the C type GDK_WINDOW_EDGE_WEST.
	Windowedge_West Windowedge = 3
	// Windowedge_East is a representation of the C type GDK_WINDOW_EDGE_EAST.
	Windowedge_East Windowedge = 4
	// Windowedge_SouthWest is a representation of the C type GDK_WINDOW_EDGE_SOUTH_WEST.
	Windowedge_SouthWest Windowedge = 5
	// Windowedge_South is a representation of the C type GDK_WINDOW_EDGE_SOUTH.
	Windowedge_South Windowedge = 6
	// Windowedge_SouthEast is a representation of the C type GDK_WINDOW_EDGE_SOUTH_EAST.
	Windowedge_SouthEast Windowedge = 7
)

// Windowtype is a representation of the C type GdkWindowType.
type Windowtype int

const (
	// Windowtype_Root is a representation of the C type GDK_WINDOW_ROOT.
	Windowtype_Root Windowtype = 0
	// Windowtype_Toplevel is a representation of the C type GDK_WINDOW_TOPLEVEL.
	Windowtype_Toplevel Windowtype = 1
	// Windowtype_Child is a representation of the C type GDK_WINDOW_CHILD.
	Windowtype_Child Windowtype = 2
	// Windowtype_Temp is a representation of the C type GDK_WINDOW_TEMP.
	Windowtype_Temp Windowtype = 3
	// Windowtype_Foreign is a representation of the C type GDK_WINDOW_FOREIGN.
	Windowtype_Foreign Windowtype = 4
	// Windowtype_Offscreen is a representation of the C type GDK_WINDOW_OFFSCREEN.
	Windowtype_Offscreen Windowtype = 5
	// Windowtype_Subsurface is a representation of the C type GDK_WINDOW_SUBSURFACE.
	Windowtype_Subsurface Windowtype = 6
)

// Windowtypehint is a representation of the C type GdkWindowTypeHint.
type Windowtypehint int

const (
	// Windowtypehint_Normal is a representation of the C type GDK_WINDOW_TYPE_HINT_NORMAL.
	Windowtypehint_Normal Windowtypehint = 0
	// Windowtypehint_Dialog is a representation of the C type GDK_WINDOW_TYPE_HINT_DIALOG.
	Windowtypehint_Dialog Windowtypehint = 1
	// Windowtypehint_Menu is a representation of the C type GDK_WINDOW_TYPE_HINT_MENU.
	Windowtypehint_Menu Windowtypehint = 2
	// Windowtypehint_Toolbar is a representation of the C type GDK_WINDOW_TYPE_HINT_TOOLBAR.
	Windowtypehint_Toolbar Windowtypehint = 3
	// Windowtypehint_Splashscreen is a representation of the C type GDK_WINDOW_TYPE_HINT_SPLASHSCREEN.
	Windowtypehint_Splashscreen Windowtypehint = 4
	// Windowtypehint_Utility is a representation of the C type GDK_WINDOW_TYPE_HINT_UTILITY.
	Windowtypehint_Utility Windowtypehint = 5
	// Windowtypehint_Dock is a representation of the C type GDK_WINDOW_TYPE_HINT_DOCK.
	Windowtypehint_Dock Windowtypehint = 6
	// Windowtypehint_Desktop is a representation of the C type GDK_WINDOW_TYPE_HINT_DESKTOP.
	Windowtypehint_Desktop Windowtypehint = 7
	// Windowtypehint_DropdownMenu is a representation of the C type GDK_WINDOW_TYPE_HINT_DROPDOWN_MENU.
	Windowtypehint_DropdownMenu Windowtypehint = 8
	// Windowtypehint_PopupMenu is a representation of the C type GDK_WINDOW_TYPE_HINT_POPUP_MENU.
	Windowtypehint_PopupMenu Windowtypehint = 9
	// Windowtypehint_Tooltip is a representation of the C type GDK_WINDOW_TYPE_HINT_TOOLTIP.
	Windowtypehint_Tooltip Windowtypehint = 10
	// Windowtypehint_Notification is a representation of the C type GDK_WINDOW_TYPE_HINT_NOTIFICATION.
	Windowtypehint_Notification Windowtypehint = 11
	// Windowtypehint_Combo is a representation of the C type GDK_WINDOW_TYPE_HINT_COMBO.
	Windowtypehint_Combo Windowtypehint = 12
	// Windowtypehint_Dnd is a representation of the C type GDK_WINDOW_TYPE_HINT_DND.
	Windowtypehint_Dnd Windowtypehint = 13
)

// Windowwindowclass is a representation of the C type GdkWindowWindowClass.
type Windowwindowclass int

const (
	// Windowwindowclass_InputOutput is a representation of the C type GDK_INPUT_OUTPUT.
	Windowwindowclass_InputOutput Windowwindowclass = 0
	// Windowwindowclass_InputOnly is a representation of the C type GDK_INPUT_ONLY.
	Windowwindowclass_InputOnly Windowwindowclass = 1
)
