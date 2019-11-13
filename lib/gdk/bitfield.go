// Code generated - DO NOT EDIT.

package gdk

// Anchorhints is a representation of the C type GdkAnchorHints.
//
// since 3.22
type Anchorhints int

const (
	// Anchorhints_FlipX is a representation of the C type GDK_ANCHOR_FLIP_X.
	Anchorhints_FlipX Anchorhints = 1
	// Anchorhints_FlipY is a representation of the C type GDK_ANCHOR_FLIP_Y.
	Anchorhints_FlipY Anchorhints = 2
	// Anchorhints_SlideX is a representation of the C type GDK_ANCHOR_SLIDE_X.
	Anchorhints_SlideX Anchorhints = 4
	// Anchorhints_SlideY is a representation of the C type GDK_ANCHOR_SLIDE_Y.
	Anchorhints_SlideY Anchorhints = 8
	// Anchorhints_ResizeX is a representation of the C type GDK_ANCHOR_RESIZE_X.
	Anchorhints_ResizeX Anchorhints = 16
	// Anchorhints_ResizeY is a representation of the C type GDK_ANCHOR_RESIZE_Y.
	Anchorhints_ResizeY Anchorhints = 32
	// Anchorhints_Flip is a representation of the C type GDK_ANCHOR_FLIP.
	Anchorhints_Flip Anchorhints = 3
	// Anchorhints_Slide is a representation of the C type GDK_ANCHOR_SLIDE.
	Anchorhints_Slide Anchorhints = 12
	// Anchorhints_Resize is a representation of the C type GDK_ANCHOR_RESIZE.
	Anchorhints_Resize Anchorhints = 48
)

// Axisflags is a representation of the C type GdkAxisFlags.
//
// since 3.22
type Axisflags int

const (
	// Axisflags_X is a representation of the C type GDK_AXIS_FLAG_X.
	Axisflags_X Axisflags = 2
	// Axisflags_Y is a representation of the C type GDK_AXIS_FLAG_Y.
	Axisflags_Y Axisflags = 4
	// Axisflags_Pressure is a representation of the C type GDK_AXIS_FLAG_PRESSURE.
	Axisflags_Pressure Axisflags = 8
	// Axisflags_Xtilt is a representation of the C type GDK_AXIS_FLAG_XTILT.
	Axisflags_Xtilt Axisflags = 16
	// Axisflags_Ytilt is a representation of the C type GDK_AXIS_FLAG_YTILT.
	Axisflags_Ytilt Axisflags = 32
	// Axisflags_Wheel is a representation of the C type GDK_AXIS_FLAG_WHEEL.
	Axisflags_Wheel Axisflags = 64
	// Axisflags_Distance is a representation of the C type GDK_AXIS_FLAG_DISTANCE.
	Axisflags_Distance Axisflags = 128
	// Axisflags_Rotation is a representation of the C type GDK_AXIS_FLAG_ROTATION.
	Axisflags_Rotation Axisflags = 256
	// Axisflags_Slider is a representation of the C type GDK_AXIS_FLAG_SLIDER.
	Axisflags_Slider Axisflags = 512
)

// Dragaction is a representation of the C type GdkDragAction.
type Dragaction int

const (
	// Dragaction_Default is a representation of the C type GDK_ACTION_DEFAULT.
	Dragaction_Default Dragaction = 1
	// Dragaction_Copy is a representation of the C type GDK_ACTION_COPY.
	Dragaction_Copy Dragaction = 2
	// Dragaction_Move is a representation of the C type GDK_ACTION_MOVE.
	Dragaction_Move Dragaction = 4
	// Dragaction_Link is a representation of the C type GDK_ACTION_LINK.
	Dragaction_Link Dragaction = 8
	// Dragaction_Private is a representation of the C type GDK_ACTION_PRIVATE.
	Dragaction_Private Dragaction = 16
	// Dragaction_Ask is a representation of the C type GDK_ACTION_ASK.
	Dragaction_Ask Dragaction = 32
)

// Eventmask is a representation of the C type GdkEventMask.
type Eventmask int

const (
	// Eventmask_ExposureMask is a representation of the C type GDK_EXPOSURE_MASK.
	Eventmask_ExposureMask Eventmask = 2
	// Eventmask_PointerMotionMask is a representation of the C type GDK_POINTER_MOTION_MASK.
	Eventmask_PointerMotionMask Eventmask = 4
	// Eventmask_PointerMotionHintMask is a representation of the C type GDK_POINTER_MOTION_HINT_MASK.
	Eventmask_PointerMotionHintMask Eventmask = 8
	// Eventmask_ButtonMotionMask is a representation of the C type GDK_BUTTON_MOTION_MASK.
	Eventmask_ButtonMotionMask Eventmask = 16
	// Eventmask_Button1MotionMask is a representation of the C type GDK_BUTTON1_MOTION_MASK.
	Eventmask_Button1MotionMask Eventmask = 32
	// Eventmask_Button2MotionMask is a representation of the C type GDK_BUTTON2_MOTION_MASK.
	Eventmask_Button2MotionMask Eventmask = 64
	// Eventmask_Button3MotionMask is a representation of the C type GDK_BUTTON3_MOTION_MASK.
	Eventmask_Button3MotionMask Eventmask = 128
	// Eventmask_ButtonPressMask is a representation of the C type GDK_BUTTON_PRESS_MASK.
	Eventmask_ButtonPressMask Eventmask = 256
	// Eventmask_ButtonReleaseMask is a representation of the C type GDK_BUTTON_RELEASE_MASK.
	Eventmask_ButtonReleaseMask Eventmask = 512
	// Eventmask_KeyPressMask is a representation of the C type GDK_KEY_PRESS_MASK.
	Eventmask_KeyPressMask Eventmask = 1024
	// Eventmask_KeyReleaseMask is a representation of the C type GDK_KEY_RELEASE_MASK.
	Eventmask_KeyReleaseMask Eventmask = 2048
	// Eventmask_EnterNotifyMask is a representation of the C type GDK_ENTER_NOTIFY_MASK.
	Eventmask_EnterNotifyMask Eventmask = 4096
	// Eventmask_LeaveNotifyMask is a representation of the C type GDK_LEAVE_NOTIFY_MASK.
	Eventmask_LeaveNotifyMask Eventmask = 8192
	// Eventmask_FocusChangeMask is a representation of the C type GDK_FOCUS_CHANGE_MASK.
	Eventmask_FocusChangeMask Eventmask = 16384
	// Eventmask_StructureMask is a representation of the C type GDK_STRUCTURE_MASK.
	Eventmask_StructureMask Eventmask = 32768
	// Eventmask_PropertyChangeMask is a representation of the C type GDK_PROPERTY_CHANGE_MASK.
	Eventmask_PropertyChangeMask Eventmask = 65536
	// Eventmask_VisibilityNotifyMask is a representation of the C type GDK_VISIBILITY_NOTIFY_MASK.
	Eventmask_VisibilityNotifyMask Eventmask = 131072
	// Eventmask_ProximityInMask is a representation of the C type GDK_PROXIMITY_IN_MASK.
	Eventmask_ProximityInMask Eventmask = 262144
	// Eventmask_ProximityOutMask is a representation of the C type GDK_PROXIMITY_OUT_MASK.
	Eventmask_ProximityOutMask Eventmask = 524288
	// Eventmask_SubstructureMask is a representation of the C type GDK_SUBSTRUCTURE_MASK.
	Eventmask_SubstructureMask Eventmask = 1048576
	// Eventmask_ScrollMask is a representation of the C type GDK_SCROLL_MASK.
	Eventmask_ScrollMask Eventmask = 2097152
	// Eventmask_TouchMask is a representation of the C type GDK_TOUCH_MASK.
	Eventmask_TouchMask Eventmask = 4194304
	// Eventmask_SmoothScrollMask is a representation of the C type GDK_SMOOTH_SCROLL_MASK.
	Eventmask_SmoothScrollMask Eventmask = 8388608
	// Eventmask_TouchpadGestureMask is a representation of the C type GDK_TOUCHPAD_GESTURE_MASK.
	Eventmask_TouchpadGestureMask Eventmask = 16777216
	// Eventmask_TabletPadMask is a representation of the C type GDK_TABLET_PAD_MASK.
	Eventmask_TabletPadMask Eventmask = 33554432
	// Eventmask_AllEventsMask is a representation of the C type GDK_ALL_EVENTS_MASK.
	Eventmask_AllEventsMask Eventmask = 67108862
)

// Frameclockphase is a representation of the C type GdkFrameClockPhase.
//
// since 3.8
type Frameclockphase int

const (
	// Frameclockphase_None is a representation of the C type GDK_FRAME_CLOCK_PHASE_NONE.
	Frameclockphase_None Frameclockphase = 0
	// Frameclockphase_FlushEvents is a representation of the C type GDK_FRAME_CLOCK_PHASE_FLUSH_EVENTS.
	Frameclockphase_FlushEvents Frameclockphase = 1
	// Frameclockphase_BeforePaint is a representation of the C type GDK_FRAME_CLOCK_PHASE_BEFORE_PAINT.
	Frameclockphase_BeforePaint Frameclockphase = 2
	// Frameclockphase_Update is a representation of the C type GDK_FRAME_CLOCK_PHASE_UPDATE.
	Frameclockphase_Update Frameclockphase = 4
	// Frameclockphase_Layout is a representation of the C type GDK_FRAME_CLOCK_PHASE_LAYOUT.
	Frameclockphase_Layout Frameclockphase = 8
	// Frameclockphase_Paint is a representation of the C type GDK_FRAME_CLOCK_PHASE_PAINT.
	Frameclockphase_Paint Frameclockphase = 16
	// Frameclockphase_ResumeEvents is a representation of the C type GDK_FRAME_CLOCK_PHASE_RESUME_EVENTS.
	Frameclockphase_ResumeEvents Frameclockphase = 32
	// Frameclockphase_AfterPaint is a representation of the C type GDK_FRAME_CLOCK_PHASE_AFTER_PAINT.
	Frameclockphase_AfterPaint Frameclockphase = 64
)

// Modifiertype is a representation of the C type GdkModifierType.
type Modifiertype int

const (
	// Modifiertype_ShiftMask is a representation of the C type GDK_SHIFT_MASK.
	Modifiertype_ShiftMask Modifiertype = 1
	// Modifiertype_LockMask is a representation of the C type GDK_LOCK_MASK.
	Modifiertype_LockMask Modifiertype = 2
	// Modifiertype_ControlMask is a representation of the C type GDK_CONTROL_MASK.
	Modifiertype_ControlMask Modifiertype = 4
	// Modifiertype_Mod1Mask is a representation of the C type GDK_MOD1_MASK.
	Modifiertype_Mod1Mask Modifiertype = 8
	// Modifiertype_Mod2Mask is a representation of the C type GDK_MOD2_MASK.
	Modifiertype_Mod2Mask Modifiertype = 16
	// Modifiertype_Mod3Mask is a representation of the C type GDK_MOD3_MASK.
	Modifiertype_Mod3Mask Modifiertype = 32
	// Modifiertype_Mod4Mask is a representation of the C type GDK_MOD4_MASK.
	Modifiertype_Mod4Mask Modifiertype = 64
	// Modifiertype_Mod5Mask is a representation of the C type GDK_MOD5_MASK.
	Modifiertype_Mod5Mask Modifiertype = 128
	// Modifiertype_Button1Mask is a representation of the C type GDK_BUTTON1_MASK.
	Modifiertype_Button1Mask Modifiertype = 256
	// Modifiertype_Button2Mask is a representation of the C type GDK_BUTTON2_MASK.
	Modifiertype_Button2Mask Modifiertype = 512
	// Modifiertype_Button3Mask is a representation of the C type GDK_BUTTON3_MASK.
	Modifiertype_Button3Mask Modifiertype = 1024
	// Modifiertype_Button4Mask is a representation of the C type GDK_BUTTON4_MASK.
	Modifiertype_Button4Mask Modifiertype = 2048
	// Modifiertype_Button5Mask is a representation of the C type GDK_BUTTON5_MASK.
	Modifiertype_Button5Mask Modifiertype = 4096
	// Modifiertype_ModifierReserved13Mask is a representation of the C type GDK_MODIFIER_RESERVED_13_MASK.
	Modifiertype_ModifierReserved13Mask Modifiertype = 8192
	// Modifiertype_ModifierReserved14Mask is a representation of the C type GDK_MODIFIER_RESERVED_14_MASK.
	Modifiertype_ModifierReserved14Mask Modifiertype = 16384
	// Modifiertype_ModifierReserved15Mask is a representation of the C type GDK_MODIFIER_RESERVED_15_MASK.
	Modifiertype_ModifierReserved15Mask Modifiertype = 32768
	// Modifiertype_ModifierReserved16Mask is a representation of the C type GDK_MODIFIER_RESERVED_16_MASK.
	Modifiertype_ModifierReserved16Mask Modifiertype = 65536
	// Modifiertype_ModifierReserved17Mask is a representation of the C type GDK_MODIFIER_RESERVED_17_MASK.
	Modifiertype_ModifierReserved17Mask Modifiertype = 131072
	// Modifiertype_ModifierReserved18Mask is a representation of the C type GDK_MODIFIER_RESERVED_18_MASK.
	Modifiertype_ModifierReserved18Mask Modifiertype = 262144
	// Modifiertype_ModifierReserved19Mask is a representation of the C type GDK_MODIFIER_RESERVED_19_MASK.
	Modifiertype_ModifierReserved19Mask Modifiertype = 524288
	// Modifiertype_ModifierReserved20Mask is a representation of the C type GDK_MODIFIER_RESERVED_20_MASK.
	Modifiertype_ModifierReserved20Mask Modifiertype = 1048576
	// Modifiertype_ModifierReserved21Mask is a representation of the C type GDK_MODIFIER_RESERVED_21_MASK.
	Modifiertype_ModifierReserved21Mask Modifiertype = 2097152
	// Modifiertype_ModifierReserved22Mask is a representation of the C type GDK_MODIFIER_RESERVED_22_MASK.
	Modifiertype_ModifierReserved22Mask Modifiertype = 4194304
	// Modifiertype_ModifierReserved23Mask is a representation of the C type GDK_MODIFIER_RESERVED_23_MASK.
	Modifiertype_ModifierReserved23Mask Modifiertype = 8388608
	// Modifiertype_ModifierReserved24Mask is a representation of the C type GDK_MODIFIER_RESERVED_24_MASK.
	Modifiertype_ModifierReserved24Mask Modifiertype = 16777216
	// Modifiertype_ModifierReserved25Mask is a representation of the C type GDK_MODIFIER_RESERVED_25_MASK.
	Modifiertype_ModifierReserved25Mask Modifiertype = 33554432
	// Modifiertype_SuperMask is a representation of the C type GDK_SUPER_MASK.
	Modifiertype_SuperMask Modifiertype = 67108864
	// Modifiertype_HyperMask is a representation of the C type GDK_HYPER_MASK.
	Modifiertype_HyperMask Modifiertype = 134217728
	// Modifiertype_MetaMask is a representation of the C type GDK_META_MASK.
	Modifiertype_MetaMask Modifiertype = 268435456
	// Modifiertype_ModifierReserved29Mask is a representation of the C type GDK_MODIFIER_RESERVED_29_MASK.
	Modifiertype_ModifierReserved29Mask Modifiertype = 536870912
	// Modifiertype_ReleaseMask is a representation of the C type GDK_RELEASE_MASK.
	Modifiertype_ReleaseMask Modifiertype = 1073741824
	// Modifiertype_ModifierMask is a representation of the C type GDK_MODIFIER_MASK.
	Modifiertype_ModifierMask Modifiertype = 1543512063
)

// Seatcapabilities is a representation of the C type GdkSeatCapabilities.
//
// since 3.20
type Seatcapabilities int

const (
	// Seatcapabilities_None is a representation of the C type GDK_SEAT_CAPABILITY_NONE.
	Seatcapabilities_None Seatcapabilities = 0
	// Seatcapabilities_Pointer is a representation of the C type GDK_SEAT_CAPABILITY_POINTER.
	Seatcapabilities_Pointer Seatcapabilities = 1
	// Seatcapabilities_Touch is a representation of the C type GDK_SEAT_CAPABILITY_TOUCH.
	Seatcapabilities_Touch Seatcapabilities = 2
	// Seatcapabilities_TabletStylus is a representation of the C type GDK_SEAT_CAPABILITY_TABLET_STYLUS.
	Seatcapabilities_TabletStylus Seatcapabilities = 4
	// Seatcapabilities_Keyboard is a representation of the C type GDK_SEAT_CAPABILITY_KEYBOARD.
	Seatcapabilities_Keyboard Seatcapabilities = 8
	// Seatcapabilities_AllPointing is a representation of the C type GDK_SEAT_CAPABILITY_ALL_POINTING.
	Seatcapabilities_AllPointing Seatcapabilities = 7
	// Seatcapabilities_All is a representation of the C type GDK_SEAT_CAPABILITY_ALL.
	Seatcapabilities_All Seatcapabilities = 15
)

// Wmdecoration is a representation of the C type GdkWMDecoration.
type Wmdecoration int

const (
	// Wmdecoration_All is a representation of the C type GDK_DECOR_ALL.
	Wmdecoration_All Wmdecoration = 1
	// Wmdecoration_Border is a representation of the C type GDK_DECOR_BORDER.
	Wmdecoration_Border Wmdecoration = 2
	// Wmdecoration_Resizeh is a representation of the C type GDK_DECOR_RESIZEH.
	Wmdecoration_Resizeh Wmdecoration = 4
	// Wmdecoration_Title is a representation of the C type GDK_DECOR_TITLE.
	Wmdecoration_Title Wmdecoration = 8
	// Wmdecoration_Menu is a representation of the C type GDK_DECOR_MENU.
	Wmdecoration_Menu Wmdecoration = 16
	// Wmdecoration_Minimize is a representation of the C type GDK_DECOR_MINIMIZE.
	Wmdecoration_Minimize Wmdecoration = 32
	// Wmdecoration_Maximize is a representation of the C type GDK_DECOR_MAXIMIZE.
	Wmdecoration_Maximize Wmdecoration = 64
)

// Wmfunction is a representation of the C type GdkWMFunction.
type Wmfunction int

const (
	// Wmfunction_All is a representation of the C type GDK_FUNC_ALL.
	Wmfunction_All Wmfunction = 1
	// Wmfunction_Resize is a representation of the C type GDK_FUNC_RESIZE.
	Wmfunction_Resize Wmfunction = 2
	// Wmfunction_Move is a representation of the C type GDK_FUNC_MOVE.
	Wmfunction_Move Wmfunction = 4
	// Wmfunction_Minimize is a representation of the C type GDK_FUNC_MINIMIZE.
	Wmfunction_Minimize Wmfunction = 8
	// Wmfunction_Maximize is a representation of the C type GDK_FUNC_MAXIMIZE.
	Wmfunction_Maximize Wmfunction = 16
	// Wmfunction_Close is a representation of the C type GDK_FUNC_CLOSE.
	Wmfunction_Close Wmfunction = 32
)

// Windowattributestype is a representation of the C type GdkWindowAttributesType.
type Windowattributestype int

const (
	// Windowattributestype_Title is a representation of the C type GDK_WA_TITLE.
	Windowattributestype_Title Windowattributestype = 2
	// Windowattributestype_X is a representation of the C type GDK_WA_X.
	Windowattributestype_X Windowattributestype = 4
	// Windowattributestype_Y is a representation of the C type GDK_WA_Y.
	Windowattributestype_Y Windowattributestype = 8
	// Windowattributestype_Cursor is a representation of the C type GDK_WA_CURSOR.
	Windowattributestype_Cursor Windowattributestype = 16
	// Windowattributestype_Visual is a representation of the C type GDK_WA_VISUAL.
	Windowattributestype_Visual Windowattributestype = 32
	// Windowattributestype_Wmclass is a representation of the C type GDK_WA_WMCLASS.
	Windowattributestype_Wmclass Windowattributestype = 64
	// Windowattributestype_Noredir is a representation of the C type GDK_WA_NOREDIR.
	Windowattributestype_Noredir Windowattributestype = 128
	// Windowattributestype_TypeHint is a representation of the C type GDK_WA_TYPE_HINT.
	Windowattributestype_TypeHint Windowattributestype = 256
)

// Windowhints is a representation of the C type GdkWindowHints.
type Windowhints int

const (
	// Windowhints_Pos is a representation of the C type GDK_HINT_POS.
	Windowhints_Pos Windowhints = 1
	// Windowhints_MinSize is a representation of the C type GDK_HINT_MIN_SIZE.
	Windowhints_MinSize Windowhints = 2
	// Windowhints_MaxSize is a representation of the C type GDK_HINT_MAX_SIZE.
	Windowhints_MaxSize Windowhints = 4
	// Windowhints_BaseSize is a representation of the C type GDK_HINT_BASE_SIZE.
	Windowhints_BaseSize Windowhints = 8
	// Windowhints_Aspect is a representation of the C type GDK_HINT_ASPECT.
	Windowhints_Aspect Windowhints = 16
	// Windowhints_ResizeInc is a representation of the C type GDK_HINT_RESIZE_INC.
	Windowhints_ResizeInc Windowhints = 32
	// Windowhints_WinGravity is a representation of the C type GDK_HINT_WIN_GRAVITY.
	Windowhints_WinGravity Windowhints = 64
	// Windowhints_UserPos is a representation of the C type GDK_HINT_USER_POS.
	Windowhints_UserPos Windowhints = 128
	// Windowhints_UserSize is a representation of the C type GDK_HINT_USER_SIZE.
	Windowhints_UserSize Windowhints = 256
)

// Windowstate is a representation of the C type GdkWindowState.
type Windowstate int

const (
	// Windowstate_Withdrawn is a representation of the C type GDK_WINDOW_STATE_WITHDRAWN.
	Windowstate_Withdrawn Windowstate = 1
	// Windowstate_Iconified is a representation of the C type GDK_WINDOW_STATE_ICONIFIED.
	Windowstate_Iconified Windowstate = 2
	// Windowstate_Maximized is a representation of the C type GDK_WINDOW_STATE_MAXIMIZED.
	Windowstate_Maximized Windowstate = 4
	// Windowstate_Sticky is a representation of the C type GDK_WINDOW_STATE_STICKY.
	Windowstate_Sticky Windowstate = 8
	// Windowstate_Fullscreen is a representation of the C type GDK_WINDOW_STATE_FULLSCREEN.
	Windowstate_Fullscreen Windowstate = 16
	// Windowstate_Above is a representation of the C type GDK_WINDOW_STATE_ABOVE.
	Windowstate_Above Windowstate = 32
	// Windowstate_Below is a representation of the C type GDK_WINDOW_STATE_BELOW.
	Windowstate_Below Windowstate = 64
	// Windowstate_Focused is a representation of the C type GDK_WINDOW_STATE_FOCUSED.
	Windowstate_Focused Windowstate = 128
	// Windowstate_Tiled is a representation of the C type GDK_WINDOW_STATE_TILED.
	Windowstate_Tiled Windowstate = 256
	// Windowstate_TopTiled is a representation of the C type GDK_WINDOW_STATE_TOP_TILED.
	Windowstate_TopTiled Windowstate = 512
	// Windowstate_TopResizable is a representation of the C type GDK_WINDOW_STATE_TOP_RESIZABLE.
	Windowstate_TopResizable Windowstate = 1024
	// Windowstate_RightTiled is a representation of the C type GDK_WINDOW_STATE_RIGHT_TILED.
	Windowstate_RightTiled Windowstate = 2048
	// Windowstate_RightResizable is a representation of the C type GDK_WINDOW_STATE_RIGHT_RESIZABLE.
	Windowstate_RightResizable Windowstate = 4096
	// Windowstate_BottomTiled is a representation of the C type GDK_WINDOW_STATE_BOTTOM_TILED.
	Windowstate_BottomTiled Windowstate = 8192
	// Windowstate_BottomResizable is a representation of the C type GDK_WINDOW_STATE_BOTTOM_RESIZABLE.
	Windowstate_BottomResizable Windowstate = 16384
	// Windowstate_LeftTiled is a representation of the C type GDK_WINDOW_STATE_LEFT_TILED.
	Windowstate_LeftTiled Windowstate = 32768
	// Windowstate_LeftResizable is a representation of the C type GDK_WINDOW_STATE_LEFT_RESIZABLE.
	Windowstate_LeftResizable Windowstate = 65536
)
