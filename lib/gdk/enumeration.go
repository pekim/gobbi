// Code generated - DO NOT EDIT.

package gdk

// AxisUse is a representation of the C type GdkAxisUse.
//
type AxisUse int

const (
	// AxisUse_Ignore is a representation of the C type GDK_AXIS_IGNORE.
	AxisUse_Ignore AxisUse = 0
	// AxisUse_X is a representation of the C type GDK_AXIS_X.
	AxisUse_X AxisUse = 1
	// AxisUse_Y is a representation of the C type GDK_AXIS_Y.
	AxisUse_Y AxisUse = 2
	// AxisUse_Pressure is a representation of the C type GDK_AXIS_PRESSURE.
	AxisUse_Pressure AxisUse = 3
	// AxisUse_Xtilt is a representation of the C type GDK_AXIS_XTILT.
	AxisUse_Xtilt AxisUse = 4
	// AxisUse_Ytilt is a representation of the C type GDK_AXIS_YTILT.
	AxisUse_Ytilt AxisUse = 5
	// AxisUse_Wheel is a representation of the C type GDK_AXIS_WHEEL.
	AxisUse_Wheel AxisUse = 6
	// AxisUse_Distance is a representation of the C type GDK_AXIS_DISTANCE.
	AxisUse_Distance AxisUse = 7
	// AxisUse_Rotation is a representation of the C type GDK_AXIS_ROTATION.
	AxisUse_Rotation AxisUse = 8
	// AxisUse_Slider is a representation of the C type GDK_AXIS_SLIDER.
	AxisUse_Slider AxisUse = 9
	// AxisUse_Last is a representation of the C type GDK_AXIS_LAST.
	AxisUse_Last AxisUse = 10
)

// ByteOrder is a representation of the C type GdkByteOrder.
//
type ByteOrder int

const (
	// ByteOrder_LsbFirst is a representation of the C type GDK_LSB_FIRST.
	ByteOrder_LsbFirst ByteOrder = 0
	// ByteOrder_MsbFirst is a representation of the C type GDK_MSB_FIRST.
	ByteOrder_MsbFirst ByteOrder = 1
)

// CrossingMode is a representation of the C type GdkCrossingMode.
//
type CrossingMode int

const (
	// CrossingMode_Normal is a representation of the C type GDK_CROSSING_NORMAL.
	CrossingMode_Normal CrossingMode = 0
	// CrossingMode_Grab is a representation of the C type GDK_CROSSING_GRAB.
	CrossingMode_Grab CrossingMode = 1
	// CrossingMode_Ungrab is a representation of the C type GDK_CROSSING_UNGRAB.
	CrossingMode_Ungrab CrossingMode = 2
	// CrossingMode_GtkGrab is a representation of the C type GDK_CROSSING_GTK_GRAB.
	CrossingMode_GtkGrab CrossingMode = 3
	// CrossingMode_GtkUngrab is a representation of the C type GDK_CROSSING_GTK_UNGRAB.
	CrossingMode_GtkUngrab CrossingMode = 4
	// CrossingMode_StateChanged is a representation of the C type GDK_CROSSING_STATE_CHANGED.
	CrossingMode_StateChanged CrossingMode = 5
	// CrossingMode_TouchBegin is a representation of the C type GDK_CROSSING_TOUCH_BEGIN.
	CrossingMode_TouchBegin CrossingMode = 6
	// CrossingMode_TouchEnd is a representation of the C type GDK_CROSSING_TOUCH_END.
	CrossingMode_TouchEnd CrossingMode = 7
	// CrossingMode_DeviceSwitch is a representation of the C type GDK_CROSSING_DEVICE_SWITCH.
	CrossingMode_DeviceSwitch CrossingMode = 8
)

// CursorType is a representation of the C type GdkCursorType.
//
type CursorType int

const (
	// CursorType_XCursor is a representation of the C type GDK_X_CURSOR.
	CursorType_XCursor CursorType = 0
	// CursorType_Arrow is a representation of the C type GDK_ARROW.
	CursorType_Arrow CursorType = 2
	// CursorType_BasedArrowDown is a representation of the C type GDK_BASED_ARROW_DOWN.
	CursorType_BasedArrowDown CursorType = 4
	// CursorType_BasedArrowUp is a representation of the C type GDK_BASED_ARROW_UP.
	CursorType_BasedArrowUp CursorType = 6
	// CursorType_Boat is a representation of the C type GDK_BOAT.
	CursorType_Boat CursorType = 8
	// CursorType_Bogosity is a representation of the C type GDK_BOGOSITY.
	CursorType_Bogosity CursorType = 10
	// CursorType_BottomLeftCorner is a representation of the C type GDK_BOTTOM_LEFT_CORNER.
	CursorType_BottomLeftCorner CursorType = 12
	// CursorType_BottomRightCorner is a representation of the C type GDK_BOTTOM_RIGHT_CORNER.
	CursorType_BottomRightCorner CursorType = 14
	// CursorType_BottomSide is a representation of the C type GDK_BOTTOM_SIDE.
	CursorType_BottomSide CursorType = 16
	// CursorType_BottomTee is a representation of the C type GDK_BOTTOM_TEE.
	CursorType_BottomTee CursorType = 18
	// CursorType_BoxSpiral is a representation of the C type GDK_BOX_SPIRAL.
	CursorType_BoxSpiral CursorType = 20
	// CursorType_CenterPtr is a representation of the C type GDK_CENTER_PTR.
	CursorType_CenterPtr CursorType = 22
	// CursorType_Circle is a representation of the C type GDK_CIRCLE.
	CursorType_Circle CursorType = 24
	// CursorType_Clock is a representation of the C type GDK_CLOCK.
	CursorType_Clock CursorType = 26
	// CursorType_CoffeeMug is a representation of the C type GDK_COFFEE_MUG.
	CursorType_CoffeeMug CursorType = 28
	// CursorType_Cross is a representation of the C type GDK_CROSS.
	CursorType_Cross CursorType = 30
	// CursorType_CrossReverse is a representation of the C type GDK_CROSS_REVERSE.
	CursorType_CrossReverse CursorType = 32
	// CursorType_Crosshair is a representation of the C type GDK_CROSSHAIR.
	CursorType_Crosshair CursorType = 34
	// CursorType_DiamondCross is a representation of the C type GDK_DIAMOND_CROSS.
	CursorType_DiamondCross CursorType = 36
	// CursorType_Dot is a representation of the C type GDK_DOT.
	CursorType_Dot CursorType = 38
	// CursorType_Dotbox is a representation of the C type GDK_DOTBOX.
	CursorType_Dotbox CursorType = 40
	// CursorType_DoubleArrow is a representation of the C type GDK_DOUBLE_ARROW.
	CursorType_DoubleArrow CursorType = 42
	// CursorType_DraftLarge is a representation of the C type GDK_DRAFT_LARGE.
	CursorType_DraftLarge CursorType = 44
	// CursorType_DraftSmall is a representation of the C type GDK_DRAFT_SMALL.
	CursorType_DraftSmall CursorType = 46
	// CursorType_DrapedBox is a representation of the C type GDK_DRAPED_BOX.
	CursorType_DrapedBox CursorType = 48
	// CursorType_Exchange is a representation of the C type GDK_EXCHANGE.
	CursorType_Exchange CursorType = 50
	// CursorType_Fleur is a representation of the C type GDK_FLEUR.
	CursorType_Fleur CursorType = 52
	// CursorType_Gobbler is a representation of the C type GDK_GOBBLER.
	CursorType_Gobbler CursorType = 54
	// CursorType_Gumby is a representation of the C type GDK_GUMBY.
	CursorType_Gumby CursorType = 56
	// CursorType_Hand1 is a representation of the C type GDK_HAND1.
	CursorType_Hand1 CursorType = 58
	// CursorType_Hand2 is a representation of the C type GDK_HAND2.
	CursorType_Hand2 CursorType = 60
	// CursorType_Heart is a representation of the C type GDK_HEART.
	CursorType_Heart CursorType = 62
	// CursorType_Icon is a representation of the C type GDK_ICON.
	CursorType_Icon CursorType = 64
	// CursorType_IronCross is a representation of the C type GDK_IRON_CROSS.
	CursorType_IronCross CursorType = 66
	// CursorType_LeftPtr is a representation of the C type GDK_LEFT_PTR.
	CursorType_LeftPtr CursorType = 68
	// CursorType_LeftSide is a representation of the C type GDK_LEFT_SIDE.
	CursorType_LeftSide CursorType = 70
	// CursorType_LeftTee is a representation of the C type GDK_LEFT_TEE.
	CursorType_LeftTee CursorType = 72
	// CursorType_Leftbutton is a representation of the C type GDK_LEFTBUTTON.
	CursorType_Leftbutton CursorType = 74
	// CursorType_LlAngle is a representation of the C type GDK_LL_ANGLE.
	CursorType_LlAngle CursorType = 76
	// CursorType_LrAngle is a representation of the C type GDK_LR_ANGLE.
	CursorType_LrAngle CursorType = 78
	// CursorType_Man is a representation of the C type GDK_MAN.
	CursorType_Man CursorType = 80
	// CursorType_Middlebutton is a representation of the C type GDK_MIDDLEBUTTON.
	CursorType_Middlebutton CursorType = 82
	// CursorType_Mouse is a representation of the C type GDK_MOUSE.
	CursorType_Mouse CursorType = 84
	// CursorType_Pencil is a representation of the C type GDK_PENCIL.
	CursorType_Pencil CursorType = 86
	// CursorType_Pirate is a representation of the C type GDK_PIRATE.
	CursorType_Pirate CursorType = 88
	// CursorType_Plus is a representation of the C type GDK_PLUS.
	CursorType_Plus CursorType = 90
	// CursorType_QuestionArrow is a representation of the C type GDK_QUESTION_ARROW.
	CursorType_QuestionArrow CursorType = 92
	// CursorType_RightPtr is a representation of the C type GDK_RIGHT_PTR.
	CursorType_RightPtr CursorType = 94
	// CursorType_RightSide is a representation of the C type GDK_RIGHT_SIDE.
	CursorType_RightSide CursorType = 96
	// CursorType_RightTee is a representation of the C type GDK_RIGHT_TEE.
	CursorType_RightTee CursorType = 98
	// CursorType_Rightbutton is a representation of the C type GDK_RIGHTBUTTON.
	CursorType_Rightbutton CursorType = 100
	// CursorType_RtlLogo is a representation of the C type GDK_RTL_LOGO.
	CursorType_RtlLogo CursorType = 102
	// CursorType_Sailboat is a representation of the C type GDK_SAILBOAT.
	CursorType_Sailboat CursorType = 104
	// CursorType_SbDownArrow is a representation of the C type GDK_SB_DOWN_ARROW.
	CursorType_SbDownArrow CursorType = 106
	// CursorType_SbHDoubleArrow is a representation of the C type GDK_SB_H_DOUBLE_ARROW.
	CursorType_SbHDoubleArrow CursorType = 108
	// CursorType_SbLeftArrow is a representation of the C type GDK_SB_LEFT_ARROW.
	CursorType_SbLeftArrow CursorType = 110
	// CursorType_SbRightArrow is a representation of the C type GDK_SB_RIGHT_ARROW.
	CursorType_SbRightArrow CursorType = 112
	// CursorType_SbUpArrow is a representation of the C type GDK_SB_UP_ARROW.
	CursorType_SbUpArrow CursorType = 114
	// CursorType_SbVDoubleArrow is a representation of the C type GDK_SB_V_DOUBLE_ARROW.
	CursorType_SbVDoubleArrow CursorType = 116
	// CursorType_Shuttle is a representation of the C type GDK_SHUTTLE.
	CursorType_Shuttle CursorType = 118
	// CursorType_Sizing is a representation of the C type GDK_SIZING.
	CursorType_Sizing CursorType = 120
	// CursorType_Spider is a representation of the C type GDK_SPIDER.
	CursorType_Spider CursorType = 122
	// CursorType_Spraycan is a representation of the C type GDK_SPRAYCAN.
	CursorType_Spraycan CursorType = 124
	// CursorType_Star is a representation of the C type GDK_STAR.
	CursorType_Star CursorType = 126
	// CursorType_Target is a representation of the C type GDK_TARGET.
	CursorType_Target CursorType = 128
	// CursorType_Tcross is a representation of the C type GDK_TCROSS.
	CursorType_Tcross CursorType = 130
	// CursorType_TopLeftArrow is a representation of the C type GDK_TOP_LEFT_ARROW.
	CursorType_TopLeftArrow CursorType = 132
	// CursorType_TopLeftCorner is a representation of the C type GDK_TOP_LEFT_CORNER.
	CursorType_TopLeftCorner CursorType = 134
	// CursorType_TopRightCorner is a representation of the C type GDK_TOP_RIGHT_CORNER.
	CursorType_TopRightCorner CursorType = 136
	// CursorType_TopSide is a representation of the C type GDK_TOP_SIDE.
	CursorType_TopSide CursorType = 138
	// CursorType_TopTee is a representation of the C type GDK_TOP_TEE.
	CursorType_TopTee CursorType = 140
	// CursorType_Trek is a representation of the C type GDK_TREK.
	CursorType_Trek CursorType = 142
	// CursorType_UlAngle is a representation of the C type GDK_UL_ANGLE.
	CursorType_UlAngle CursorType = 144
	// CursorType_Umbrella is a representation of the C type GDK_UMBRELLA.
	CursorType_Umbrella CursorType = 146
	// CursorType_UrAngle is a representation of the C type GDK_UR_ANGLE.
	CursorType_UrAngle CursorType = 148
	// CursorType_Watch is a representation of the C type GDK_WATCH.
	CursorType_Watch CursorType = 150
	// CursorType_Xterm is a representation of the C type GDK_XTERM.
	CursorType_Xterm CursorType = 152
	// CursorType_LastCursor is a representation of the C type GDK_LAST_CURSOR.
	CursorType_LastCursor CursorType = 153
	// CursorType_BlankCursor is a representation of the C type GDK_BLANK_CURSOR.
	CursorType_BlankCursor CursorType = -2
	// CursorType_CursorIsPixmap is a representation of the C type GDK_CURSOR_IS_PIXMAP.
	CursorType_CursorIsPixmap CursorType = -1
)

// DevicePadFeature is a representation of the C type GdkDevicePadFeature.
//
type DevicePadFeature int

const (
	// DevicePadFeature_Button is a representation of the C type GDK_DEVICE_PAD_FEATURE_BUTTON.
	DevicePadFeature_Button DevicePadFeature = 0
	// DevicePadFeature_Ring is a representation of the C type GDK_DEVICE_PAD_FEATURE_RING.
	DevicePadFeature_Ring DevicePadFeature = 1
	// DevicePadFeature_Strip is a representation of the C type GDK_DEVICE_PAD_FEATURE_STRIP.
	DevicePadFeature_Strip DevicePadFeature = 2
)

// DeviceToolType is a representation of the C type GdkDeviceToolType.
//
// since 3.22
type DeviceToolType int

const (
	// DeviceToolType_Unknown is a representation of the C type GDK_DEVICE_TOOL_TYPE_UNKNOWN.
	DeviceToolType_Unknown DeviceToolType = 0
	// DeviceToolType_Pen is a representation of the C type GDK_DEVICE_TOOL_TYPE_PEN.
	DeviceToolType_Pen DeviceToolType = 1
	// DeviceToolType_Eraser is a representation of the C type GDK_DEVICE_TOOL_TYPE_ERASER.
	DeviceToolType_Eraser DeviceToolType = 2
	// DeviceToolType_Brush is a representation of the C type GDK_DEVICE_TOOL_TYPE_BRUSH.
	DeviceToolType_Brush DeviceToolType = 3
	// DeviceToolType_Pencil is a representation of the C type GDK_DEVICE_TOOL_TYPE_PENCIL.
	DeviceToolType_Pencil DeviceToolType = 4
	// DeviceToolType_Airbrush is a representation of the C type GDK_DEVICE_TOOL_TYPE_AIRBRUSH.
	DeviceToolType_Airbrush DeviceToolType = 5
	// DeviceToolType_Mouse is a representation of the C type GDK_DEVICE_TOOL_TYPE_MOUSE.
	DeviceToolType_Mouse DeviceToolType = 6
	// DeviceToolType_Lens is a representation of the C type GDK_DEVICE_TOOL_TYPE_LENS.
	DeviceToolType_Lens DeviceToolType = 7
)

// DeviceType is a representation of the C type GdkDeviceType.
//
type DeviceType int

const (
	// DeviceType_Master is a representation of the C type GDK_DEVICE_TYPE_MASTER.
	DeviceType_Master DeviceType = 0
	// DeviceType_Slave is a representation of the C type GDK_DEVICE_TYPE_SLAVE.
	DeviceType_Slave DeviceType = 1
	// DeviceType_Floating is a representation of the C type GDK_DEVICE_TYPE_FLOATING.
	DeviceType_Floating DeviceType = 2
)

// DragCancelReason is a representation of the C type GdkDragCancelReason.
//
// since 3.20
type DragCancelReason int

const (
	// DragCancelReason_NoTarget is a representation of the C type GDK_DRAG_CANCEL_NO_TARGET.
	DragCancelReason_NoTarget DragCancelReason = 0
	// DragCancelReason_UserCancelled is a representation of the C type GDK_DRAG_CANCEL_USER_CANCELLED.
	DragCancelReason_UserCancelled DragCancelReason = 1
	// DragCancelReason_Error is a representation of the C type GDK_DRAG_CANCEL_ERROR.
	DragCancelReason_Error DragCancelReason = 2
)

// DragProtocol is a representation of the C type GdkDragProtocol.
//
type DragProtocol int

const (
	// DragProtocol_None is a representation of the C type GDK_DRAG_PROTO_NONE.
	DragProtocol_None DragProtocol = 0
	// DragProtocol_Motif is a representation of the C type GDK_DRAG_PROTO_MOTIF.
	DragProtocol_Motif DragProtocol = 1
	// DragProtocol_Xdnd is a representation of the C type GDK_DRAG_PROTO_XDND.
	DragProtocol_Xdnd DragProtocol = 2
	// DragProtocol_Rootwin is a representation of the C type GDK_DRAG_PROTO_ROOTWIN.
	DragProtocol_Rootwin DragProtocol = 3
	// DragProtocol_Win32Dropfiles is a representation of the C type GDK_DRAG_PROTO_WIN32_DROPFILES.
	DragProtocol_Win32Dropfiles DragProtocol = 4
	// DragProtocol_Ole2 is a representation of the C type GDK_DRAG_PROTO_OLE2.
	DragProtocol_Ole2 DragProtocol = 5
	// DragProtocol_Local is a representation of the C type GDK_DRAG_PROTO_LOCAL.
	DragProtocol_Local DragProtocol = 6
	// DragProtocol_Wayland is a representation of the C type GDK_DRAG_PROTO_WAYLAND.
	DragProtocol_Wayland DragProtocol = 7
)

// EventType is a representation of the C type GdkEventType.
//
type EventType int

const (
	// EventType_Nothing is a representation of the C type GDK_NOTHING.
	EventType_Nothing EventType = -1
	// EventType_Delete is a representation of the C type GDK_DELETE.
	EventType_Delete EventType = 0
	// EventType_Destroy is a representation of the C type GDK_DESTROY.
	EventType_Destroy EventType = 1
	// EventType_Expose is a representation of the C type GDK_EXPOSE.
	EventType_Expose EventType = 2
	// EventType_MotionNotify is a representation of the C type GDK_MOTION_NOTIFY.
	EventType_MotionNotify EventType = 3
	// EventType_ButtonPress is a representation of the C type GDK_BUTTON_PRESS.
	EventType_ButtonPress EventType = 4
	// EventType_2buttonPress is a representation of the C type GDK_2BUTTON_PRESS.
	EventType_2buttonPress EventType = 5
	// EventType_DoubleButtonPress is a representation of the C type GDK_DOUBLE_BUTTON_PRESS.
	EventType_DoubleButtonPress EventType = 5
	// EventType_3buttonPress is a representation of the C type GDK_3BUTTON_PRESS.
	EventType_3buttonPress EventType = 6
	// EventType_TripleButtonPress is a representation of the C type GDK_TRIPLE_BUTTON_PRESS.
	EventType_TripleButtonPress EventType = 6
	// EventType_ButtonRelease is a representation of the C type GDK_BUTTON_RELEASE.
	EventType_ButtonRelease EventType = 7
	// EventType_KeyPress is a representation of the C type GDK_KEY_PRESS.
	EventType_KeyPress EventType = 8
	// EventType_KeyRelease is a representation of the C type GDK_KEY_RELEASE.
	EventType_KeyRelease EventType = 9
	// EventType_EnterNotify is a representation of the C type GDK_ENTER_NOTIFY.
	EventType_EnterNotify EventType = 10
	// EventType_LeaveNotify is a representation of the C type GDK_LEAVE_NOTIFY.
	EventType_LeaveNotify EventType = 11
	// EventType_FocusChange is a representation of the C type GDK_FOCUS_CHANGE.
	EventType_FocusChange EventType = 12
	// EventType_Configure is a representation of the C type GDK_CONFIGURE.
	EventType_Configure EventType = 13
	// EventType_Map is a representation of the C type GDK_MAP.
	EventType_Map EventType = 14
	// EventType_Unmap is a representation of the C type GDK_UNMAP.
	EventType_Unmap EventType = 15
	// EventType_PropertyNotify is a representation of the C type GDK_PROPERTY_NOTIFY.
	EventType_PropertyNotify EventType = 16
	// EventType_SelectionClear is a representation of the C type GDK_SELECTION_CLEAR.
	EventType_SelectionClear EventType = 17
	// EventType_SelectionRequest is a representation of the C type GDK_SELECTION_REQUEST.
	EventType_SelectionRequest EventType = 18
	// EventType_SelectionNotify is a representation of the C type GDK_SELECTION_NOTIFY.
	EventType_SelectionNotify EventType = 19
	// EventType_ProximityIn is a representation of the C type GDK_PROXIMITY_IN.
	EventType_ProximityIn EventType = 20
	// EventType_ProximityOut is a representation of the C type GDK_PROXIMITY_OUT.
	EventType_ProximityOut EventType = 21
	// EventType_DragEnter is a representation of the C type GDK_DRAG_ENTER.
	EventType_DragEnter EventType = 22
	// EventType_DragLeave is a representation of the C type GDK_DRAG_LEAVE.
	EventType_DragLeave EventType = 23
	// EventType_DragMotion is a representation of the C type GDK_DRAG_MOTION.
	EventType_DragMotion EventType = 24
	// EventType_DragStatus is a representation of the C type GDK_DRAG_STATUS.
	EventType_DragStatus EventType = 25
	// EventType_DropStart is a representation of the C type GDK_DROP_START.
	EventType_DropStart EventType = 26
	// EventType_DropFinished is a representation of the C type GDK_DROP_FINISHED.
	EventType_DropFinished EventType = 27
	// EventType_ClientEvent is a representation of the C type GDK_CLIENT_EVENT.
	EventType_ClientEvent EventType = 28
	// EventType_VisibilityNotify is a representation of the C type GDK_VISIBILITY_NOTIFY.
	EventType_VisibilityNotify EventType = 29
	// EventType_Scroll is a representation of the C type GDK_SCROLL.
	EventType_Scroll EventType = 31
	// EventType_WindowState is a representation of the C type GDK_WINDOW_STATE.
	EventType_WindowState EventType = 32
	// EventType_Setting is a representation of the C type GDK_SETTING.
	EventType_Setting EventType = 33
	// EventType_OwnerChange is a representation of the C type GDK_OWNER_CHANGE.
	EventType_OwnerChange EventType = 34
	// EventType_GrabBroken is a representation of the C type GDK_GRAB_BROKEN.
	EventType_GrabBroken EventType = 35
	// EventType_Damage is a representation of the C type GDK_DAMAGE.
	EventType_Damage EventType = 36
	// EventType_TouchBegin is a representation of the C type GDK_TOUCH_BEGIN.
	EventType_TouchBegin EventType = 37
	// EventType_TouchUpdate is a representation of the C type GDK_TOUCH_UPDATE.
	EventType_TouchUpdate EventType = 38
	// EventType_TouchEnd is a representation of the C type GDK_TOUCH_END.
	EventType_TouchEnd EventType = 39
	// EventType_TouchCancel is a representation of the C type GDK_TOUCH_CANCEL.
	EventType_TouchCancel EventType = 40
	// EventType_TouchpadSwipe is a representation of the C type GDK_TOUCHPAD_SWIPE.
	EventType_TouchpadSwipe EventType = 41
	// EventType_TouchpadPinch is a representation of the C type GDK_TOUCHPAD_PINCH.
	EventType_TouchpadPinch EventType = 42
	// EventType_PadButtonPress is a representation of the C type GDK_PAD_BUTTON_PRESS.
	EventType_PadButtonPress EventType = 43
	// EventType_PadButtonRelease is a representation of the C type GDK_PAD_BUTTON_RELEASE.
	EventType_PadButtonRelease EventType = 44
	// EventType_PadRing is a representation of the C type GDK_PAD_RING.
	EventType_PadRing EventType = 45
	// EventType_PadStrip is a representation of the C type GDK_PAD_STRIP.
	EventType_PadStrip EventType = 46
	// EventType_PadGroupMode is a representation of the C type GDK_PAD_GROUP_MODE.
	EventType_PadGroupMode EventType = 47
	// EventType_EventLast is a representation of the C type GDK_EVENT_LAST.
	EventType_EventLast EventType = 48
)

// FilterReturn is a representation of the C type GdkFilterReturn.
//
type FilterReturn int

const (
	// FilterReturn_Continue is a representation of the C type GDK_FILTER_CONTINUE.
	FilterReturn_Continue FilterReturn = 0
	// FilterReturn_Translate is a representation of the C type GDK_FILTER_TRANSLATE.
	FilterReturn_Translate FilterReturn = 1
	// FilterReturn_Remove is a representation of the C type GDK_FILTER_REMOVE.
	FilterReturn_Remove FilterReturn = 2
)

// FullscreenMode is a representation of the C type GdkFullscreenMode.
//
// since 3.8
type FullscreenMode int

const (
	// FullscreenMode_CurrentMonitor is a representation of the C type GDK_FULLSCREEN_ON_CURRENT_MONITOR.
	FullscreenMode_CurrentMonitor FullscreenMode = 0
	// FullscreenMode_AllMonitors is a representation of the C type GDK_FULLSCREEN_ON_ALL_MONITORS.
	FullscreenMode_AllMonitors FullscreenMode = 1
)

// GLError is a representation of the C type GdkGLError.
//
// since 3.16
type GLError int

const (
	// GLError_NotAvailable is a representation of the C type GDK_GL_ERROR_NOT_AVAILABLE.
	GLError_NotAvailable GLError = 0
	// GLError_UnsupportedFormat is a representation of the C type GDK_GL_ERROR_UNSUPPORTED_FORMAT.
	GLError_UnsupportedFormat GLError = 1
	// GLError_UnsupportedProfile is a representation of the C type GDK_GL_ERROR_UNSUPPORTED_PROFILE.
	GLError_UnsupportedProfile GLError = 2
)

// GrabOwnership is a representation of the C type GdkGrabOwnership.
//
type GrabOwnership int

const (
	// GrabOwnership_None is a representation of the C type GDK_OWNERSHIP_NONE.
	GrabOwnership_None GrabOwnership = 0
	// GrabOwnership_Window is a representation of the C type GDK_OWNERSHIP_WINDOW.
	GrabOwnership_Window GrabOwnership = 1
	// GrabOwnership_Application is a representation of the C type GDK_OWNERSHIP_APPLICATION.
	GrabOwnership_Application GrabOwnership = 2
)

// GrabStatus is a representation of the C type GdkGrabStatus.
//
type GrabStatus int

const (
	// GrabStatus_Success is a representation of the C type GDK_GRAB_SUCCESS.
	GrabStatus_Success GrabStatus = 0
	// GrabStatus_AlreadyGrabbed is a representation of the C type GDK_GRAB_ALREADY_GRABBED.
	GrabStatus_AlreadyGrabbed GrabStatus = 1
	// GrabStatus_InvalidTime is a representation of the C type GDK_GRAB_INVALID_TIME.
	GrabStatus_InvalidTime GrabStatus = 2
	// GrabStatus_NotViewable is a representation of the C type GDK_GRAB_NOT_VIEWABLE.
	GrabStatus_NotViewable GrabStatus = 3
	// GrabStatus_Frozen is a representation of the C type GDK_GRAB_FROZEN.
	GrabStatus_Frozen GrabStatus = 4
	// GrabStatus_Failed is a representation of the C type GDK_GRAB_FAILED.
	GrabStatus_Failed GrabStatus = 5
)

// Gravity is a representation of the C type GdkGravity.
//
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

// InputMode is a representation of the C type GdkInputMode.
//
type InputMode int

const (
	// InputMode_Disabled is a representation of the C type GDK_MODE_DISABLED.
	InputMode_Disabled InputMode = 0
	// InputMode_Screen is a representation of the C type GDK_MODE_SCREEN.
	InputMode_Screen InputMode = 1
	// InputMode_Window is a representation of the C type GDK_MODE_WINDOW.
	InputMode_Window InputMode = 2
)

// InputSource is a representation of the C type GdkInputSource.
//
type InputSource int

const (
	// InputSource_Mouse is a representation of the C type GDK_SOURCE_MOUSE.
	InputSource_Mouse InputSource = 0
	// InputSource_Pen is a representation of the C type GDK_SOURCE_PEN.
	InputSource_Pen InputSource = 1
	// InputSource_Eraser is a representation of the C type GDK_SOURCE_ERASER.
	InputSource_Eraser InputSource = 2
	// InputSource_Cursor is a representation of the C type GDK_SOURCE_CURSOR.
	InputSource_Cursor InputSource = 3
	// InputSource_Keyboard is a representation of the C type GDK_SOURCE_KEYBOARD.
	InputSource_Keyboard InputSource = 4
	// InputSource_Touchscreen is a representation of the C type GDK_SOURCE_TOUCHSCREEN.
	InputSource_Touchscreen InputSource = 5
	// InputSource_Touchpad is a representation of the C type GDK_SOURCE_TOUCHPAD.
	InputSource_Touchpad InputSource = 6
	// InputSource_Trackpoint is a representation of the C type GDK_SOURCE_TRACKPOINT.
	InputSource_Trackpoint InputSource = 7
	// InputSource_TabletPad is a representation of the C type GDK_SOURCE_TABLET_PAD.
	InputSource_TabletPad InputSource = 8
)

// ModifierIntent is a representation of the C type GdkModifierIntent.
//
// since 3.4
type ModifierIntent int

const (
	// ModifierIntent_PrimaryAccelerator is a representation of the C type GDK_MODIFIER_INTENT_PRIMARY_ACCELERATOR.
	ModifierIntent_PrimaryAccelerator ModifierIntent = 0
	// ModifierIntent_ContextMenu is a representation of the C type GDK_MODIFIER_INTENT_CONTEXT_MENU.
	ModifierIntent_ContextMenu ModifierIntent = 1
	// ModifierIntent_ExtendSelection is a representation of the C type GDK_MODIFIER_INTENT_EXTEND_SELECTION.
	ModifierIntent_ExtendSelection ModifierIntent = 2
	// ModifierIntent_ModifySelection is a representation of the C type GDK_MODIFIER_INTENT_MODIFY_SELECTION.
	ModifierIntent_ModifySelection ModifierIntent = 3
	// ModifierIntent_NoTextInput is a representation of the C type GDK_MODIFIER_INTENT_NO_TEXT_INPUT.
	ModifierIntent_NoTextInput ModifierIntent = 4
	// ModifierIntent_ShiftGroup is a representation of the C type GDK_MODIFIER_INTENT_SHIFT_GROUP.
	ModifierIntent_ShiftGroup ModifierIntent = 5
	// ModifierIntent_DefaultModMask is a representation of the C type GDK_MODIFIER_INTENT_DEFAULT_MOD_MASK.
	ModifierIntent_DefaultModMask ModifierIntent = 6
)

// NotifyType is a representation of the C type GdkNotifyType.
//
type NotifyType int

const (
	// NotifyType_Ancestor is a representation of the C type GDK_NOTIFY_ANCESTOR.
	NotifyType_Ancestor NotifyType = 0
	// NotifyType_Virtual is a representation of the C type GDK_NOTIFY_VIRTUAL.
	NotifyType_Virtual NotifyType = 1
	// NotifyType_Inferior is a representation of the C type GDK_NOTIFY_INFERIOR.
	NotifyType_Inferior NotifyType = 2
	// NotifyType_Nonlinear is a representation of the C type GDK_NOTIFY_NONLINEAR.
	NotifyType_Nonlinear NotifyType = 3
	// NotifyType_NonlinearVirtual is a representation of the C type GDK_NOTIFY_NONLINEAR_VIRTUAL.
	NotifyType_NonlinearVirtual NotifyType = 4
	// NotifyType_Unknown is a representation of the C type GDK_NOTIFY_UNKNOWN.
	NotifyType_Unknown NotifyType = 5
)

// OwnerChange is a representation of the C type GdkOwnerChange.
//
type OwnerChange int

const (
	// OwnerChange_NewOwner is a representation of the C type GDK_OWNER_CHANGE_NEW_OWNER.
	OwnerChange_NewOwner OwnerChange = 0
	// OwnerChange_Destroy is a representation of the C type GDK_OWNER_CHANGE_DESTROY.
	OwnerChange_Destroy OwnerChange = 1
	// OwnerChange_Close is a representation of the C type GDK_OWNER_CHANGE_CLOSE.
	OwnerChange_Close OwnerChange = 2
)

// PropMode is a representation of the C type GdkPropMode.
//
type PropMode int

const (
	// PropMode_Replace is a representation of the C type GDK_PROP_MODE_REPLACE.
	PropMode_Replace PropMode = 0
	// PropMode_Prepend is a representation of the C type GDK_PROP_MODE_PREPEND.
	PropMode_Prepend PropMode = 1
	// PropMode_Append is a representation of the C type GDK_PROP_MODE_APPEND.
	PropMode_Append PropMode = 2
)

// PropertyState is a representation of the C type GdkPropertyState.
//
type PropertyState int

const (
	// PropertyState_NewValue is a representation of the C type GDK_PROPERTY_NEW_VALUE.
	PropertyState_NewValue PropertyState = 0
	// PropertyState_Delete is a representation of the C type GDK_PROPERTY_DELETE.
	PropertyState_Delete PropertyState = 1
)

// ScrollDirection is a representation of the C type GdkScrollDirection.
//
type ScrollDirection int

const (
	// ScrollDirection_Up is a representation of the C type GDK_SCROLL_UP.
	ScrollDirection_Up ScrollDirection = 0
	// ScrollDirection_Down is a representation of the C type GDK_SCROLL_DOWN.
	ScrollDirection_Down ScrollDirection = 1
	// ScrollDirection_Left is a representation of the C type GDK_SCROLL_LEFT.
	ScrollDirection_Left ScrollDirection = 2
	// ScrollDirection_Right is a representation of the C type GDK_SCROLL_RIGHT.
	ScrollDirection_Right ScrollDirection = 3
	// ScrollDirection_Smooth is a representation of the C type GDK_SCROLL_SMOOTH.
	ScrollDirection_Smooth ScrollDirection = 4
)

// SettingAction is a representation of the C type GdkSettingAction.
//
type SettingAction int

const (
	// SettingAction_New is a representation of the C type GDK_SETTING_ACTION_NEW.
	SettingAction_New SettingAction = 0
	// SettingAction_Changed is a representation of the C type GDK_SETTING_ACTION_CHANGED.
	SettingAction_Changed SettingAction = 1
	// SettingAction_Deleted is a representation of the C type GDK_SETTING_ACTION_DELETED.
	SettingAction_Deleted SettingAction = 2
)

// Status is a representation of the C type GdkStatus.
//
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

// SubpixelLayout is a representation of the C type GdkSubpixelLayout.
//
// since 3.22
type SubpixelLayout int

const (
	// SubpixelLayout_Unknown is a representation of the C type GDK_SUBPIXEL_LAYOUT_UNKNOWN.
	SubpixelLayout_Unknown SubpixelLayout = 0
	// SubpixelLayout_None is a representation of the C type GDK_SUBPIXEL_LAYOUT_NONE.
	SubpixelLayout_None SubpixelLayout = 1
	// SubpixelLayout_HorizontalRgb is a representation of the C type GDK_SUBPIXEL_LAYOUT_HORIZONTAL_RGB.
	SubpixelLayout_HorizontalRgb SubpixelLayout = 2
	// SubpixelLayout_HorizontalBgr is a representation of the C type GDK_SUBPIXEL_LAYOUT_HORIZONTAL_BGR.
	SubpixelLayout_HorizontalBgr SubpixelLayout = 3
	// SubpixelLayout_VerticalRgb is a representation of the C type GDK_SUBPIXEL_LAYOUT_VERTICAL_RGB.
	SubpixelLayout_VerticalRgb SubpixelLayout = 4
	// SubpixelLayout_VerticalBgr is a representation of the C type GDK_SUBPIXEL_LAYOUT_VERTICAL_BGR.
	SubpixelLayout_VerticalBgr SubpixelLayout = 5
)

// TouchpadGesturePhase is a representation of the C type GdkTouchpadGesturePhase.
//
type TouchpadGesturePhase int

const (
	// TouchpadGesturePhase_Begin is a representation of the C type GDK_TOUCHPAD_GESTURE_PHASE_BEGIN.
	TouchpadGesturePhase_Begin TouchpadGesturePhase = 0
	// TouchpadGesturePhase_Update is a representation of the C type GDK_TOUCHPAD_GESTURE_PHASE_UPDATE.
	TouchpadGesturePhase_Update TouchpadGesturePhase = 1
	// TouchpadGesturePhase_End is a representation of the C type GDK_TOUCHPAD_GESTURE_PHASE_END.
	TouchpadGesturePhase_End TouchpadGesturePhase = 2
	// TouchpadGesturePhase_Cancel is a representation of the C type GDK_TOUCHPAD_GESTURE_PHASE_CANCEL.
	TouchpadGesturePhase_Cancel TouchpadGesturePhase = 3
)

// VisibilityState is a representation of the C type GdkVisibilityState.
//
type VisibilityState int

const (
	// VisibilityState_Unobscured is a representation of the C type GDK_VISIBILITY_UNOBSCURED.
	VisibilityState_Unobscured VisibilityState = 0
	// VisibilityState_Partial is a representation of the C type GDK_VISIBILITY_PARTIAL.
	VisibilityState_Partial VisibilityState = 1
	// VisibilityState_FullyObscured is a representation of the C type GDK_VISIBILITY_FULLY_OBSCURED.
	VisibilityState_FullyObscured VisibilityState = 2
)

// VisualType is a representation of the C type GdkVisualType.
//
type VisualType int

const (
	// VisualType_StaticGray is a representation of the C type GDK_VISUAL_STATIC_GRAY.
	VisualType_StaticGray VisualType = 0
	// VisualType_Grayscale is a representation of the C type GDK_VISUAL_GRAYSCALE.
	VisualType_Grayscale VisualType = 1
	// VisualType_StaticColor is a representation of the C type GDK_VISUAL_STATIC_COLOR.
	VisualType_StaticColor VisualType = 2
	// VisualType_PseudoColor is a representation of the C type GDK_VISUAL_PSEUDO_COLOR.
	VisualType_PseudoColor VisualType = 3
	// VisualType_TrueColor is a representation of the C type GDK_VISUAL_TRUE_COLOR.
	VisualType_TrueColor VisualType = 4
	// VisualType_DirectColor is a representation of the C type GDK_VISUAL_DIRECT_COLOR.
	VisualType_DirectColor VisualType = 5
)

// WindowEdge is a representation of the C type GdkWindowEdge.
//
type WindowEdge int

const (
	// WindowEdge_NorthWest is a representation of the C type GDK_WINDOW_EDGE_NORTH_WEST.
	WindowEdge_NorthWest WindowEdge = 0
	// WindowEdge_North is a representation of the C type GDK_WINDOW_EDGE_NORTH.
	WindowEdge_North WindowEdge = 1
	// WindowEdge_NorthEast is a representation of the C type GDK_WINDOW_EDGE_NORTH_EAST.
	WindowEdge_NorthEast WindowEdge = 2
	// WindowEdge_West is a representation of the C type GDK_WINDOW_EDGE_WEST.
	WindowEdge_West WindowEdge = 3
	// WindowEdge_East is a representation of the C type GDK_WINDOW_EDGE_EAST.
	WindowEdge_East WindowEdge = 4
	// WindowEdge_SouthWest is a representation of the C type GDK_WINDOW_EDGE_SOUTH_WEST.
	WindowEdge_SouthWest WindowEdge = 5
	// WindowEdge_South is a representation of the C type GDK_WINDOW_EDGE_SOUTH.
	WindowEdge_South WindowEdge = 6
	// WindowEdge_SouthEast is a representation of the C type GDK_WINDOW_EDGE_SOUTH_EAST.
	WindowEdge_SouthEast WindowEdge = 7
)

// WindowType is a representation of the C type GdkWindowType.
//
type WindowType int

const (
	// WindowType_Root is a representation of the C type GDK_WINDOW_ROOT.
	WindowType_Root WindowType = 0
	// WindowType_Toplevel is a representation of the C type GDK_WINDOW_TOPLEVEL.
	WindowType_Toplevel WindowType = 1
	// WindowType_Child is a representation of the C type GDK_WINDOW_CHILD.
	WindowType_Child WindowType = 2
	// WindowType_Temp is a representation of the C type GDK_WINDOW_TEMP.
	WindowType_Temp WindowType = 3
	// WindowType_Foreign is a representation of the C type GDK_WINDOW_FOREIGN.
	WindowType_Foreign WindowType = 4
	// WindowType_Offscreen is a representation of the C type GDK_WINDOW_OFFSCREEN.
	WindowType_Offscreen WindowType = 5
	// WindowType_Subsurface is a representation of the C type GDK_WINDOW_SUBSURFACE.
	WindowType_Subsurface WindowType = 6
)

// WindowTypeHint is a representation of the C type GdkWindowTypeHint.
//
type WindowTypeHint int

const (
	// WindowTypeHint_Normal is a representation of the C type GDK_WINDOW_TYPE_HINT_NORMAL.
	WindowTypeHint_Normal WindowTypeHint = 0
	// WindowTypeHint_Dialog is a representation of the C type GDK_WINDOW_TYPE_HINT_DIALOG.
	WindowTypeHint_Dialog WindowTypeHint = 1
	// WindowTypeHint_Menu is a representation of the C type GDK_WINDOW_TYPE_HINT_MENU.
	WindowTypeHint_Menu WindowTypeHint = 2
	// WindowTypeHint_Toolbar is a representation of the C type GDK_WINDOW_TYPE_HINT_TOOLBAR.
	WindowTypeHint_Toolbar WindowTypeHint = 3
	// WindowTypeHint_Splashscreen is a representation of the C type GDK_WINDOW_TYPE_HINT_SPLASHSCREEN.
	WindowTypeHint_Splashscreen WindowTypeHint = 4
	// WindowTypeHint_Utility is a representation of the C type GDK_WINDOW_TYPE_HINT_UTILITY.
	WindowTypeHint_Utility WindowTypeHint = 5
	// WindowTypeHint_Dock is a representation of the C type GDK_WINDOW_TYPE_HINT_DOCK.
	WindowTypeHint_Dock WindowTypeHint = 6
	// WindowTypeHint_Desktop is a representation of the C type GDK_WINDOW_TYPE_HINT_DESKTOP.
	WindowTypeHint_Desktop WindowTypeHint = 7
	// WindowTypeHint_DropdownMenu is a representation of the C type GDK_WINDOW_TYPE_HINT_DROPDOWN_MENU.
	WindowTypeHint_DropdownMenu WindowTypeHint = 8
	// WindowTypeHint_PopupMenu is a representation of the C type GDK_WINDOW_TYPE_HINT_POPUP_MENU.
	WindowTypeHint_PopupMenu WindowTypeHint = 9
	// WindowTypeHint_Tooltip is a representation of the C type GDK_WINDOW_TYPE_HINT_TOOLTIP.
	WindowTypeHint_Tooltip WindowTypeHint = 10
	// WindowTypeHint_Notification is a representation of the C type GDK_WINDOW_TYPE_HINT_NOTIFICATION.
	WindowTypeHint_Notification WindowTypeHint = 11
	// WindowTypeHint_Combo is a representation of the C type GDK_WINDOW_TYPE_HINT_COMBO.
	WindowTypeHint_Combo WindowTypeHint = 12
	// WindowTypeHint_Dnd is a representation of the C type GDK_WINDOW_TYPE_HINT_DND.
	WindowTypeHint_Dnd WindowTypeHint = 13
)

// WindowWindowClass is a representation of the C type GdkWindowWindowClass.
//
type WindowWindowClass int

const (
	// WindowWindowClass_InputOutput is a representation of the C type GDK_INPUT_OUTPUT.
	WindowWindowClass_InputOutput WindowWindowClass = 0
	// WindowWindowClass_InputOnly is a representation of the C type GDK_INPUT_ONLY.
	WindowWindowClass_InputOnly WindowWindowClass = 1
)
