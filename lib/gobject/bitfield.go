// Code generated - DO NOT EDIT.

package gobject

// BindingFlags is a representation of the C type GBindingFlags.
//
// since 2.26
type BindingFlags int

const (
	// BindingFlags_Default is a representation of the C type G_BINDING_DEFAULT.
	BindingFlags_Default BindingFlags = 0
	// BindingFlags_Bidirectional is a representation of the C type G_BINDING_BIDIRECTIONAL.
	BindingFlags_Bidirectional BindingFlags = 1
	// BindingFlags_SyncCreate is a representation of the C type G_BINDING_SYNC_CREATE.
	BindingFlags_SyncCreate BindingFlags = 2
	// BindingFlags_InvertBoolean is a representation of the C type G_BINDING_INVERT_BOOLEAN.
	BindingFlags_InvertBoolean BindingFlags = 4
)

// ConnectFlags is a representation of the C type GConnectFlags.
type ConnectFlags int

const (
	// ConnectFlags_After is a representation of the C type G_CONNECT_AFTER.
	ConnectFlags_After ConnectFlags = 1
	// ConnectFlags_Swapped is a representation of the C type G_CONNECT_SWAPPED.
	ConnectFlags_Swapped ConnectFlags = 2
)

// ParamFlags is a representation of the C type GParamFlags.
type ParamFlags int

const (
	// ParamFlags_Readable is a representation of the C type G_PARAM_READABLE.
	ParamFlags_Readable ParamFlags = 1
	// ParamFlags_Writable is a representation of the C type G_PARAM_WRITABLE.
	ParamFlags_Writable ParamFlags = 2
	// ParamFlags_Readwrite is a representation of the C type G_PARAM_READWRITE.
	ParamFlags_Readwrite ParamFlags = 3
	// ParamFlags_Construct is a representation of the C type G_PARAM_CONSTRUCT.
	ParamFlags_Construct ParamFlags = 4
	// ParamFlags_ConstructOnly is a representation of the C type G_PARAM_CONSTRUCT_ONLY.
	ParamFlags_ConstructOnly ParamFlags = 8
	// ParamFlags_LaxValidation is a representation of the C type G_PARAM_LAX_VALIDATION.
	ParamFlags_LaxValidation ParamFlags = 16
	// ParamFlags_StaticName is a representation of the C type G_PARAM_STATIC_NAME.
	ParamFlags_StaticName ParamFlags = 32
	// ParamFlags_Private is a representation of the C type G_PARAM_PRIVATE.
	ParamFlags_Private ParamFlags = 32
	// ParamFlags_StaticNick is a representation of the C type G_PARAM_STATIC_NICK.
	ParamFlags_StaticNick ParamFlags = 64
	// ParamFlags_StaticBlurb is a representation of the C type G_PARAM_STATIC_BLURB.
	ParamFlags_StaticBlurb ParamFlags = 128
	// ParamFlags_ExplicitNotify is a representation of the C type G_PARAM_EXPLICIT_NOTIFY.
	ParamFlags_ExplicitNotify ParamFlags = 1073741824
	// ParamFlags_Deprecated is a representation of the C type G_PARAM_DEPRECATED.
	ParamFlags_Deprecated ParamFlags = 2147483648
)

// SignalFlags is a representation of the C type GSignalFlags.
type SignalFlags int

const (
	// SignalFlags_RunFirst is a representation of the C type G_SIGNAL_RUN_FIRST.
	SignalFlags_RunFirst SignalFlags = 1
	// SignalFlags_RunLast is a representation of the C type G_SIGNAL_RUN_LAST.
	SignalFlags_RunLast SignalFlags = 2
	// SignalFlags_RunCleanup is a representation of the C type G_SIGNAL_RUN_CLEANUP.
	SignalFlags_RunCleanup SignalFlags = 4
	// SignalFlags_NoRecurse is a representation of the C type G_SIGNAL_NO_RECURSE.
	SignalFlags_NoRecurse SignalFlags = 8
	// SignalFlags_Detailed is a representation of the C type G_SIGNAL_DETAILED.
	SignalFlags_Detailed SignalFlags = 16
	// SignalFlags_Action is a representation of the C type G_SIGNAL_ACTION.
	SignalFlags_Action SignalFlags = 32
	// SignalFlags_NoHooks is a representation of the C type G_SIGNAL_NO_HOOKS.
	SignalFlags_NoHooks SignalFlags = 64
	// SignalFlags_MustCollect is a representation of the C type G_SIGNAL_MUST_COLLECT.
	SignalFlags_MustCollect SignalFlags = 128
	// SignalFlags_Deprecated is a representation of the C type G_SIGNAL_DEPRECATED.
	SignalFlags_Deprecated SignalFlags = 256
)

// SignalMatchType is a representation of the C type GSignalMatchType.
type SignalMatchType int

const (
	// SignalMatchType_Id is a representation of the C type G_SIGNAL_MATCH_ID.
	SignalMatchType_Id SignalMatchType = 1
	// SignalMatchType_Detail is a representation of the C type G_SIGNAL_MATCH_DETAIL.
	SignalMatchType_Detail SignalMatchType = 2
	// SignalMatchType_Closure is a representation of the C type G_SIGNAL_MATCH_CLOSURE.
	SignalMatchType_Closure SignalMatchType = 4
	// SignalMatchType_Func is a representation of the C type G_SIGNAL_MATCH_FUNC.
	SignalMatchType_Func SignalMatchType = 8
	// SignalMatchType_Data is a representation of the C type G_SIGNAL_MATCH_DATA.
	SignalMatchType_Data SignalMatchType = 16
	// SignalMatchType_Unblocked is a representation of the C type G_SIGNAL_MATCH_UNBLOCKED.
	SignalMatchType_Unblocked SignalMatchType = 32
)

// TypeDebugFlags is a representation of the C type GTypeDebugFlags.
type TypeDebugFlags int

const (
	// TypeDebugFlags_None is a representation of the C type G_TYPE_DEBUG_NONE.
	TypeDebugFlags_None TypeDebugFlags = 0
	// TypeDebugFlags_Objects is a representation of the C type G_TYPE_DEBUG_OBJECTS.
	TypeDebugFlags_Objects TypeDebugFlags = 1
	// TypeDebugFlags_Signals is a representation of the C type G_TYPE_DEBUG_SIGNALS.
	TypeDebugFlags_Signals TypeDebugFlags = 2
	// TypeDebugFlags_InstanceCount is a representation of the C type G_TYPE_DEBUG_INSTANCE_COUNT.
	TypeDebugFlags_InstanceCount TypeDebugFlags = 4
	// TypeDebugFlags_Mask is a representation of the C type G_TYPE_DEBUG_MASK.
	TypeDebugFlags_Mask TypeDebugFlags = 7
)

// TypeFlags is a representation of the C type GTypeFlags.
type TypeFlags int

const (
	// TypeFlags_Abstract is a representation of the C type G_TYPE_FLAG_ABSTRACT.
	TypeFlags_Abstract TypeFlags = 16
	// TypeFlags_ValueAbstract is a representation of the C type G_TYPE_FLAG_VALUE_ABSTRACT.
	TypeFlags_ValueAbstract TypeFlags = 32
)

// TypeFundamentalFlags is a representation of the C type GTypeFundamentalFlags.
type TypeFundamentalFlags int

const (
	// TypeFundamentalFlags_Classed is a representation of the C type G_TYPE_FLAG_CLASSED.
	TypeFundamentalFlags_Classed TypeFundamentalFlags = 1
	// TypeFundamentalFlags_Instantiatable is a representation of the C type G_TYPE_FLAG_INSTANTIATABLE.
	TypeFundamentalFlags_Instantiatable TypeFundamentalFlags = 2
	// TypeFundamentalFlags_Derivable is a representation of the C type G_TYPE_FLAG_DERIVABLE.
	TypeFundamentalFlags_Derivable TypeFundamentalFlags = 4
	// TypeFundamentalFlags_DeepDerivable is a representation of the C type G_TYPE_FLAG_DEEP_DERIVABLE.
	TypeFundamentalFlags_DeepDerivable TypeFundamentalFlags = 8
)
