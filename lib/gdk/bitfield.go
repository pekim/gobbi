// Code generated - DO NOT EDIT.

package gdk

// AnchorHints is a representation of the C type GdkAnchorHints.
//
// since 3.22
type AnchorHints int

const (
	// AnchorHints_FlipX is a representation of the C type GDK_ANCHOR_FLIP_X.
	AnchorHints_FlipX AnchorHints = 1
	// AnchorHints_FlipY is a representation of the C type GDK_ANCHOR_FLIP_Y.
	AnchorHints_FlipY AnchorHints = 2
	// AnchorHints_SlideX is a representation of the C type GDK_ANCHOR_SLIDE_X.
	AnchorHints_SlideX AnchorHints = 4
	// AnchorHints_SlideY is a representation of the C type GDK_ANCHOR_SLIDE_Y.
	AnchorHints_SlideY AnchorHints = 8
	// AnchorHints_ResizeX is a representation of the C type GDK_ANCHOR_RESIZE_X.
	AnchorHints_ResizeX AnchorHints = 16
	// AnchorHints_ResizeY is a representation of the C type GDK_ANCHOR_RESIZE_Y.
	AnchorHints_ResizeY AnchorHints = 32
	// AnchorHints_Flip is a representation of the C type GDK_ANCHOR_FLIP.
	AnchorHints_Flip AnchorHints = 3
	// AnchorHints_Slide is a representation of the C type GDK_ANCHOR_SLIDE.
	AnchorHints_Slide AnchorHints = 12
	// AnchorHints_Resize is a representation of the C type GDK_ANCHOR_RESIZE.
	AnchorHints_Resize AnchorHints = 48
)

// AxisFlags is a representation of the C type GdkAxisFlags.
//
// since 3.22
type AxisFlags int

const (
	// AxisFlags_X is a representation of the C type GDK_AXIS_FLAG_X.
	AxisFlags_X AxisFlags = 2
	// AxisFlags_Y is a representation of the C type GDK_AXIS_FLAG_Y.
	AxisFlags_Y AxisFlags = 4
	// AxisFlags_Pressure is a representation of the C type GDK_AXIS_FLAG_PRESSURE.
	AxisFlags_Pressure AxisFlags = 8
	// AxisFlags_Xtilt is a representation of the C type GDK_AXIS_FLAG_XTILT.
	AxisFlags_Xtilt AxisFlags = 16
	// AxisFlags_Ytilt is a representation of the C type GDK_AXIS_FLAG_YTILT.
	AxisFlags_Ytilt AxisFlags = 32
	// AxisFlags_Wheel is a representation of the C type GDK_AXIS_FLAG_WHEEL.
	AxisFlags_Wheel AxisFlags = 64
	// AxisFlags_Distance is a representation of the C type GDK_AXIS_FLAG_DISTANCE.
	AxisFlags_Distance AxisFlags = 128
	// AxisFlags_Rotation is a representation of the C type GDK_AXIS_FLAG_ROTATION.
	AxisFlags_Rotation AxisFlags = 256
	// AxisFlags_Slider is a representation of the C type GDK_AXIS_FLAG_SLIDER.
	AxisFlags_Slider AxisFlags = 512
)

// DragAction is a representation of the C type GdkDragAction.
type DragAction int

const (
	// DragAction_Default is a representation of the C type GDK_ACTION_DEFAULT.
	DragAction_Default DragAction = 1
	// DragAction_Copy is a representation of the C type GDK_ACTION_COPY.
	DragAction_Copy DragAction = 2
	// DragAction_Move is a representation of the C type GDK_ACTION_MOVE.
	DragAction_Move DragAction = 4
	// DragAction_Link is a representation of the C type GDK_ACTION_LINK.
	DragAction_Link DragAction = 8
	// DragAction_Private is a representation of the C type GDK_ACTION_PRIVATE.
	DragAction_Private DragAction = 16
	// DragAction_Ask is a representation of the C type GDK_ACTION_ASK.
	DragAction_Ask DragAction = 32
)

// EventMask is a representation of the C type GdkEventMask.
type EventMask int

const (
	// EventMask_ExposureMask is a representation of the C type GDK_EXPOSURE_MASK.
	EventMask_ExposureMask EventMask = 2
	// EventMask_PointerMotionMask is a representation of the C type GDK_POINTER_MOTION_MASK.
	EventMask_PointerMotionMask EventMask = 4
	// EventMask_PointerMotionHintMask is a representation of the C type GDK_POINTER_MOTION_HINT_MASK.
	EventMask_PointerMotionHintMask EventMask = 8
	// EventMask_ButtonMotionMask is a representation of the C type GDK_BUTTON_MOTION_MASK.
	EventMask_ButtonMotionMask EventMask = 16
	// EventMask_Button1MotionMask is a representation of the C type GDK_BUTTON1_MOTION_MASK.
	EventMask_Button1MotionMask EventMask = 32
	// EventMask_Button2MotionMask is a representation of the C type GDK_BUTTON2_MOTION_MASK.
	EventMask_Button2MotionMask EventMask = 64
	// EventMask_Button3MotionMask is a representation of the C type GDK_BUTTON3_MOTION_MASK.
	EventMask_Button3MotionMask EventMask = 128
	// EventMask_ButtonPressMask is a representation of the C type GDK_BUTTON_PRESS_MASK.
	EventMask_ButtonPressMask EventMask = 256
	// EventMask_ButtonReleaseMask is a representation of the C type GDK_BUTTON_RELEASE_MASK.
	EventMask_ButtonReleaseMask EventMask = 512
	// EventMask_KeyPressMask is a representation of the C type GDK_KEY_PRESS_MASK.
	EventMask_KeyPressMask EventMask = 1024
	// EventMask_KeyReleaseMask is a representation of the C type GDK_KEY_RELEASE_MASK.
	EventMask_KeyReleaseMask EventMask = 2048
	// EventMask_EnterNotifyMask is a representation of the C type GDK_ENTER_NOTIFY_MASK.
	EventMask_EnterNotifyMask EventMask = 4096
	// EventMask_LeaveNotifyMask is a representation of the C type GDK_LEAVE_NOTIFY_MASK.
	EventMask_LeaveNotifyMask EventMask = 8192
	// EventMask_FocusChangeMask is a representation of the C type GDK_FOCUS_CHANGE_MASK.
	EventMask_FocusChangeMask EventMask = 16384
	// EventMask_StructureMask is a representation of the C type GDK_STRUCTURE_MASK.
	EventMask_StructureMask EventMask = 32768
	// EventMask_PropertyChangeMask is a representation of the C type GDK_PROPERTY_CHANGE_MASK.
	EventMask_PropertyChangeMask EventMask = 65536
	// EventMask_VisibilityNotifyMask is a representation of the C type GDK_VISIBILITY_NOTIFY_MASK.
	EventMask_VisibilityNotifyMask EventMask = 131072
	// EventMask_ProximityInMask is a representation of the C type GDK_PROXIMITY_IN_MASK.
	EventMask_ProximityInMask EventMask = 262144
	// EventMask_ProximityOutMask is a representation of the C type GDK_PROXIMITY_OUT_MASK.
	EventMask_ProximityOutMask EventMask = 524288
	// EventMask_SubstructureMask is a representation of the C type GDK_SUBSTRUCTURE_MASK.
	EventMask_SubstructureMask EventMask = 1048576
	// EventMask_ScrollMask is a representation of the C type GDK_SCROLL_MASK.
	EventMask_ScrollMask EventMask = 2097152
	// EventMask_TouchMask is a representation of the C type GDK_TOUCH_MASK.
	EventMask_TouchMask EventMask = 4194304
	// EventMask_SmoothScrollMask is a representation of the C type GDK_SMOOTH_SCROLL_MASK.
	EventMask_SmoothScrollMask EventMask = 8388608
	// EventMask_TouchpadGestureMask is a representation of the C type GDK_TOUCHPAD_GESTURE_MASK.
	EventMask_TouchpadGestureMask EventMask = 16777216
	// EventMask_TabletPadMask is a representation of the C type GDK_TABLET_PAD_MASK.
	EventMask_TabletPadMask EventMask = 33554432
	// EventMask_AllEventsMask is a representation of the C type GDK_ALL_EVENTS_MASK.
	EventMask_AllEventsMask EventMask = 67108862
)

// FrameClockPhase is a representation of the C type GdkFrameClockPhase.
//
// since 3.8
type FrameClockPhase int

const (
	// FrameClockPhase_None is a representation of the C type GDK_FRAME_CLOCK_PHASE_NONE.
	FrameClockPhase_None FrameClockPhase = 0
	// FrameClockPhase_FlushEvents is a representation of the C type GDK_FRAME_CLOCK_PHASE_FLUSH_EVENTS.
	FrameClockPhase_FlushEvents FrameClockPhase = 1
	// FrameClockPhase_BeforePaint is a representation of the C type GDK_FRAME_CLOCK_PHASE_BEFORE_PAINT.
	FrameClockPhase_BeforePaint FrameClockPhase = 2
	// FrameClockPhase_Update is a representation of the C type GDK_FRAME_CLOCK_PHASE_UPDATE.
	FrameClockPhase_Update FrameClockPhase = 4
	// FrameClockPhase_Layout is a representation of the C type GDK_FRAME_CLOCK_PHASE_LAYOUT.
	FrameClockPhase_Layout FrameClockPhase = 8
	// FrameClockPhase_Paint is a representation of the C type GDK_FRAME_CLOCK_PHASE_PAINT.
	FrameClockPhase_Paint FrameClockPhase = 16
	// FrameClockPhase_ResumeEvents is a representation of the C type GDK_FRAME_CLOCK_PHASE_RESUME_EVENTS.
	FrameClockPhase_ResumeEvents FrameClockPhase = 32
	// FrameClockPhase_AfterPaint is a representation of the C type GDK_FRAME_CLOCK_PHASE_AFTER_PAINT.
	FrameClockPhase_AfterPaint FrameClockPhase = 64
)

// ModifierType is a representation of the C type GdkModifierType.
type ModifierType int

const (
	// ModifierType_ShiftMask is a representation of the C type GDK_SHIFT_MASK.
	ModifierType_ShiftMask ModifierType = 1
	// ModifierType_LockMask is a representation of the C type GDK_LOCK_MASK.
	ModifierType_LockMask ModifierType = 2
	// ModifierType_ControlMask is a representation of the C type GDK_CONTROL_MASK.
	ModifierType_ControlMask ModifierType = 4
	// ModifierType_Mod1Mask is a representation of the C type GDK_MOD1_MASK.
	ModifierType_Mod1Mask ModifierType = 8
	// ModifierType_Mod2Mask is a representation of the C type GDK_MOD2_MASK.
	ModifierType_Mod2Mask ModifierType = 16
	// ModifierType_Mod3Mask is a representation of the C type GDK_MOD3_MASK.
	ModifierType_Mod3Mask ModifierType = 32
	// ModifierType_Mod4Mask is a representation of the C type GDK_MOD4_MASK.
	ModifierType_Mod4Mask ModifierType = 64
	// ModifierType_Mod5Mask is a representation of the C type GDK_MOD5_MASK.
	ModifierType_Mod5Mask ModifierType = 128
	// ModifierType_Button1Mask is a representation of the C type GDK_BUTTON1_MASK.
	ModifierType_Button1Mask ModifierType = 256
	// ModifierType_Button2Mask is a representation of the C type GDK_BUTTON2_MASK.
	ModifierType_Button2Mask ModifierType = 512
	// ModifierType_Button3Mask is a representation of the C type GDK_BUTTON3_MASK.
	ModifierType_Button3Mask ModifierType = 1024
	// ModifierType_Button4Mask is a representation of the C type GDK_BUTTON4_MASK.
	ModifierType_Button4Mask ModifierType = 2048
	// ModifierType_Button5Mask is a representation of the C type GDK_BUTTON5_MASK.
	ModifierType_Button5Mask ModifierType = 4096
	// ModifierType_ModifierReserved13Mask is a representation of the C type GDK_MODIFIER_RESERVED_13_MASK.
	ModifierType_ModifierReserved13Mask ModifierType = 8192
	// ModifierType_ModifierReserved14Mask is a representation of the C type GDK_MODIFIER_RESERVED_14_MASK.
	ModifierType_ModifierReserved14Mask ModifierType = 16384
	// ModifierType_ModifierReserved15Mask is a representation of the C type GDK_MODIFIER_RESERVED_15_MASK.
	ModifierType_ModifierReserved15Mask ModifierType = 32768
	// ModifierType_ModifierReserved16Mask is a representation of the C type GDK_MODIFIER_RESERVED_16_MASK.
	ModifierType_ModifierReserved16Mask ModifierType = 65536
	// ModifierType_ModifierReserved17Mask is a representation of the C type GDK_MODIFIER_RESERVED_17_MASK.
	ModifierType_ModifierReserved17Mask ModifierType = 131072
	// ModifierType_ModifierReserved18Mask is a representation of the C type GDK_MODIFIER_RESERVED_18_MASK.
	ModifierType_ModifierReserved18Mask ModifierType = 262144
	// ModifierType_ModifierReserved19Mask is a representation of the C type GDK_MODIFIER_RESERVED_19_MASK.
	ModifierType_ModifierReserved19Mask ModifierType = 524288
	// ModifierType_ModifierReserved20Mask is a representation of the C type GDK_MODIFIER_RESERVED_20_MASK.
	ModifierType_ModifierReserved20Mask ModifierType = 1048576
	// ModifierType_ModifierReserved21Mask is a representation of the C type GDK_MODIFIER_RESERVED_21_MASK.
	ModifierType_ModifierReserved21Mask ModifierType = 2097152
	// ModifierType_ModifierReserved22Mask is a representation of the C type GDK_MODIFIER_RESERVED_22_MASK.
	ModifierType_ModifierReserved22Mask ModifierType = 4194304
	// ModifierType_ModifierReserved23Mask is a representation of the C type GDK_MODIFIER_RESERVED_23_MASK.
	ModifierType_ModifierReserved23Mask ModifierType = 8388608
	// ModifierType_ModifierReserved24Mask is a representation of the C type GDK_MODIFIER_RESERVED_24_MASK.
	ModifierType_ModifierReserved24Mask ModifierType = 16777216
	// ModifierType_ModifierReserved25Mask is a representation of the C type GDK_MODIFIER_RESERVED_25_MASK.
	ModifierType_ModifierReserved25Mask ModifierType = 33554432
	// ModifierType_SuperMask is a representation of the C type GDK_SUPER_MASK.
	ModifierType_SuperMask ModifierType = 67108864
	// ModifierType_HyperMask is a representation of the C type GDK_HYPER_MASK.
	ModifierType_HyperMask ModifierType = 134217728
	// ModifierType_MetaMask is a representation of the C type GDK_META_MASK.
	ModifierType_MetaMask ModifierType = 268435456
	// ModifierType_ModifierReserved29Mask is a representation of the C type GDK_MODIFIER_RESERVED_29_MASK.
	ModifierType_ModifierReserved29Mask ModifierType = 536870912
	// ModifierType_ReleaseMask is a representation of the C type GDK_RELEASE_MASK.
	ModifierType_ReleaseMask ModifierType = 1073741824
	// ModifierType_ModifierMask is a representation of the C type GDK_MODIFIER_MASK.
	ModifierType_ModifierMask ModifierType = 1543512063
)

// SeatCapabilities is a representation of the C type GdkSeatCapabilities.
//
// since 3.20
type SeatCapabilities int

const (
	// SeatCapabilities_None is a representation of the C type GDK_SEAT_CAPABILITY_NONE.
	SeatCapabilities_None SeatCapabilities = 0
	// SeatCapabilities_Pointer is a representation of the C type GDK_SEAT_CAPABILITY_POINTER.
	SeatCapabilities_Pointer SeatCapabilities = 1
	// SeatCapabilities_Touch is a representation of the C type GDK_SEAT_CAPABILITY_TOUCH.
	SeatCapabilities_Touch SeatCapabilities = 2
	// SeatCapabilities_TabletStylus is a representation of the C type GDK_SEAT_CAPABILITY_TABLET_STYLUS.
	SeatCapabilities_TabletStylus SeatCapabilities = 4
	// SeatCapabilities_Keyboard is a representation of the C type GDK_SEAT_CAPABILITY_KEYBOARD.
	SeatCapabilities_Keyboard SeatCapabilities = 8
	// SeatCapabilities_AllPointing is a representation of the C type GDK_SEAT_CAPABILITY_ALL_POINTING.
	SeatCapabilities_AllPointing SeatCapabilities = 7
	// SeatCapabilities_All is a representation of the C type GDK_SEAT_CAPABILITY_ALL.
	SeatCapabilities_All SeatCapabilities = 15
)

// WMDecoration is a representation of the C type GdkWMDecoration.
type WMDecoration int

const (
	// WMDecoration_All is a representation of the C type GDK_DECOR_ALL.
	WMDecoration_All WMDecoration = 1
	// WMDecoration_Border is a representation of the C type GDK_DECOR_BORDER.
	WMDecoration_Border WMDecoration = 2
	// WMDecoration_Resizeh is a representation of the C type GDK_DECOR_RESIZEH.
	WMDecoration_Resizeh WMDecoration = 4
	// WMDecoration_Title is a representation of the C type GDK_DECOR_TITLE.
	WMDecoration_Title WMDecoration = 8
	// WMDecoration_Menu is a representation of the C type GDK_DECOR_MENU.
	WMDecoration_Menu WMDecoration = 16
	// WMDecoration_Minimize is a representation of the C type GDK_DECOR_MINIMIZE.
	WMDecoration_Minimize WMDecoration = 32
	// WMDecoration_Maximize is a representation of the C type GDK_DECOR_MAXIMIZE.
	WMDecoration_Maximize WMDecoration = 64
)

// WMFunction is a representation of the C type GdkWMFunction.
type WMFunction int

const (
	// WMFunction_All is a representation of the C type GDK_FUNC_ALL.
	WMFunction_All WMFunction = 1
	// WMFunction_Resize is a representation of the C type GDK_FUNC_RESIZE.
	WMFunction_Resize WMFunction = 2
	// WMFunction_Move is a representation of the C type GDK_FUNC_MOVE.
	WMFunction_Move WMFunction = 4
	// WMFunction_Minimize is a representation of the C type GDK_FUNC_MINIMIZE.
	WMFunction_Minimize WMFunction = 8
	// WMFunction_Maximize is a representation of the C type GDK_FUNC_MAXIMIZE.
	WMFunction_Maximize WMFunction = 16
	// WMFunction_Close is a representation of the C type GDK_FUNC_CLOSE.
	WMFunction_Close WMFunction = 32
)

// WindowAttributesType is a representation of the C type GdkWindowAttributesType.
type WindowAttributesType int

const (
	// WindowAttributesType_Title is a representation of the C type GDK_WA_TITLE.
	WindowAttributesType_Title WindowAttributesType = 2
	// WindowAttributesType_X is a representation of the C type GDK_WA_X.
	WindowAttributesType_X WindowAttributesType = 4
	// WindowAttributesType_Y is a representation of the C type GDK_WA_Y.
	WindowAttributesType_Y WindowAttributesType = 8
	// WindowAttributesType_Cursor is a representation of the C type GDK_WA_CURSOR.
	WindowAttributesType_Cursor WindowAttributesType = 16
	// WindowAttributesType_Visual is a representation of the C type GDK_WA_VISUAL.
	WindowAttributesType_Visual WindowAttributesType = 32
	// WindowAttributesType_Wmclass is a representation of the C type GDK_WA_WMCLASS.
	WindowAttributesType_Wmclass WindowAttributesType = 64
	// WindowAttributesType_Noredir is a representation of the C type GDK_WA_NOREDIR.
	WindowAttributesType_Noredir WindowAttributesType = 128
	// WindowAttributesType_TypeHint is a representation of the C type GDK_WA_TYPE_HINT.
	WindowAttributesType_TypeHint WindowAttributesType = 256
)

// WindowHints is a representation of the C type GdkWindowHints.
type WindowHints int

const (
	// WindowHints_Pos is a representation of the C type GDK_HINT_POS.
	WindowHints_Pos WindowHints = 1
	// WindowHints_MinSize is a representation of the C type GDK_HINT_MIN_SIZE.
	WindowHints_MinSize WindowHints = 2
	// WindowHints_MaxSize is a representation of the C type GDK_HINT_MAX_SIZE.
	WindowHints_MaxSize WindowHints = 4
	// WindowHints_BaseSize is a representation of the C type GDK_HINT_BASE_SIZE.
	WindowHints_BaseSize WindowHints = 8
	// WindowHints_Aspect is a representation of the C type GDK_HINT_ASPECT.
	WindowHints_Aspect WindowHints = 16
	// WindowHints_ResizeInc is a representation of the C type GDK_HINT_RESIZE_INC.
	WindowHints_ResizeInc WindowHints = 32
	// WindowHints_WinGravity is a representation of the C type GDK_HINT_WIN_GRAVITY.
	WindowHints_WinGravity WindowHints = 64
	// WindowHints_UserPos is a representation of the C type GDK_HINT_USER_POS.
	WindowHints_UserPos WindowHints = 128
	// WindowHints_UserSize is a representation of the C type GDK_HINT_USER_SIZE.
	WindowHints_UserSize WindowHints = 256
)

// WindowState is a representation of the C type GdkWindowState.
type WindowState int

const (
	// WindowState_Withdrawn is a representation of the C type GDK_WINDOW_STATE_WITHDRAWN.
	WindowState_Withdrawn WindowState = 1
	// WindowState_Iconified is a representation of the C type GDK_WINDOW_STATE_ICONIFIED.
	WindowState_Iconified WindowState = 2
	// WindowState_Maximized is a representation of the C type GDK_WINDOW_STATE_MAXIMIZED.
	WindowState_Maximized WindowState = 4
	// WindowState_Sticky is a representation of the C type GDK_WINDOW_STATE_STICKY.
	WindowState_Sticky WindowState = 8
	// WindowState_Fullscreen is a representation of the C type GDK_WINDOW_STATE_FULLSCREEN.
	WindowState_Fullscreen WindowState = 16
	// WindowState_Above is a representation of the C type GDK_WINDOW_STATE_ABOVE.
	WindowState_Above WindowState = 32
	// WindowState_Below is a representation of the C type GDK_WINDOW_STATE_BELOW.
	WindowState_Below WindowState = 64
	// WindowState_Focused is a representation of the C type GDK_WINDOW_STATE_FOCUSED.
	WindowState_Focused WindowState = 128
	// WindowState_Tiled is a representation of the C type GDK_WINDOW_STATE_TILED.
	WindowState_Tiled WindowState = 256
	// WindowState_TopTiled is a representation of the C type GDK_WINDOW_STATE_TOP_TILED.
	WindowState_TopTiled WindowState = 512
	// WindowState_TopResizable is a representation of the C type GDK_WINDOW_STATE_TOP_RESIZABLE.
	WindowState_TopResizable WindowState = 1024
	// WindowState_RightTiled is a representation of the C type GDK_WINDOW_STATE_RIGHT_TILED.
	WindowState_RightTiled WindowState = 2048
	// WindowState_RightResizable is a representation of the C type GDK_WINDOW_STATE_RIGHT_RESIZABLE.
	WindowState_RightResizable WindowState = 4096
	// WindowState_BottomTiled is a representation of the C type GDK_WINDOW_STATE_BOTTOM_TILED.
	WindowState_BottomTiled WindowState = 8192
	// WindowState_BottomResizable is a representation of the C type GDK_WINDOW_STATE_BOTTOM_RESIZABLE.
	WindowState_BottomResizable WindowState = 16384
	// WindowState_LeftTiled is a representation of the C type GDK_WINDOW_STATE_LEFT_TILED.
	WindowState_LeftTiled WindowState = 32768
	// WindowState_LeftResizable is a representation of the C type GDK_WINDOW_STATE_LEFT_RESIZABLE.
	WindowState_LeftResizable WindowState = 65536
)
