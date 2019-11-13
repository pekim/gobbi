// Code generated - DO NOT EDIT.

package gobject

// Bindingflags is a representation of the C type GBindingFlags.
//
// since 2.26
type Bindingflags int

const (
	// Bindingflags_Default is a representation of the C type G_BINDING_DEFAULT.
	Bindingflags_Default Bindingflags = 0
	// Bindingflags_Bidirectional is a representation of the C type G_BINDING_BIDIRECTIONAL.
	Bindingflags_Bidirectional Bindingflags = 1
	// Bindingflags_SyncCreate is a representation of the C type G_BINDING_SYNC_CREATE.
	Bindingflags_SyncCreate Bindingflags = 2
	// Bindingflags_InvertBoolean is a representation of the C type G_BINDING_INVERT_BOOLEAN.
	Bindingflags_InvertBoolean Bindingflags = 4
)

// Connectflags is a representation of the C type GConnectFlags.
type Connectflags int

const (
	// Connectflags_After is a representation of the C type G_CONNECT_AFTER.
	Connectflags_After Connectflags = 1
	// Connectflags_Swapped is a representation of the C type G_CONNECT_SWAPPED.
	Connectflags_Swapped Connectflags = 2
)

// Paramflags is a representation of the C type GParamFlags.
type Paramflags int

const (
	// Paramflags_Readable is a representation of the C type G_PARAM_READABLE.
	Paramflags_Readable Paramflags = 1
	// Paramflags_Writable is a representation of the C type G_PARAM_WRITABLE.
	Paramflags_Writable Paramflags = 2
	// Paramflags_Readwrite is a representation of the C type G_PARAM_READWRITE.
	Paramflags_Readwrite Paramflags = 3
	// Paramflags_Construct is a representation of the C type G_PARAM_CONSTRUCT.
	Paramflags_Construct Paramflags = 4
	// Paramflags_ConstructOnly is a representation of the C type G_PARAM_CONSTRUCT_ONLY.
	Paramflags_ConstructOnly Paramflags = 8
	// Paramflags_LaxValidation is a representation of the C type G_PARAM_LAX_VALIDATION.
	Paramflags_LaxValidation Paramflags = 16
	// Paramflags_StaticName is a representation of the C type G_PARAM_STATIC_NAME.
	Paramflags_StaticName Paramflags = 32
	// Paramflags_Private is a representation of the C type G_PARAM_PRIVATE.
	Paramflags_Private Paramflags = 32
	// Paramflags_StaticNick is a representation of the C type G_PARAM_STATIC_NICK.
	Paramflags_StaticNick Paramflags = 64
	// Paramflags_StaticBlurb is a representation of the C type G_PARAM_STATIC_BLURB.
	Paramflags_StaticBlurb Paramflags = 128
	// Paramflags_ExplicitNotify is a representation of the C type G_PARAM_EXPLICIT_NOTIFY.
	Paramflags_ExplicitNotify Paramflags = 1073741824
	// Paramflags_Deprecated is a representation of the C type G_PARAM_DEPRECATED.
	Paramflags_Deprecated Paramflags = 2147483648
)

// Signalflags is a representation of the C type GSignalFlags.
type Signalflags int

const (
	// Signalflags_RunFirst is a representation of the C type G_SIGNAL_RUN_FIRST.
	Signalflags_RunFirst Signalflags = 1
	// Signalflags_RunLast is a representation of the C type G_SIGNAL_RUN_LAST.
	Signalflags_RunLast Signalflags = 2
	// Signalflags_RunCleanup is a representation of the C type G_SIGNAL_RUN_CLEANUP.
	Signalflags_RunCleanup Signalflags = 4
	// Signalflags_NoRecurse is a representation of the C type G_SIGNAL_NO_RECURSE.
	Signalflags_NoRecurse Signalflags = 8
	// Signalflags_Detailed is a representation of the C type G_SIGNAL_DETAILED.
	Signalflags_Detailed Signalflags = 16
	// Signalflags_Action is a representation of the C type G_SIGNAL_ACTION.
	Signalflags_Action Signalflags = 32
	// Signalflags_NoHooks is a representation of the C type G_SIGNAL_NO_HOOKS.
	Signalflags_NoHooks Signalflags = 64
	// Signalflags_MustCollect is a representation of the C type G_SIGNAL_MUST_COLLECT.
	Signalflags_MustCollect Signalflags = 128
	// Signalflags_Deprecated is a representation of the C type G_SIGNAL_DEPRECATED.
	Signalflags_Deprecated Signalflags = 256
)

// Signalmatchtype is a representation of the C type GSignalMatchType.
type Signalmatchtype int

const (
	// Signalmatchtype_Id is a representation of the C type G_SIGNAL_MATCH_ID.
	Signalmatchtype_Id Signalmatchtype = 1
	// Signalmatchtype_Detail is a representation of the C type G_SIGNAL_MATCH_DETAIL.
	Signalmatchtype_Detail Signalmatchtype = 2
	// Signalmatchtype_Closure is a representation of the C type G_SIGNAL_MATCH_CLOSURE.
	Signalmatchtype_Closure Signalmatchtype = 4
	// Signalmatchtype_Func is a representation of the C type G_SIGNAL_MATCH_FUNC.
	Signalmatchtype_Func Signalmatchtype = 8
	// Signalmatchtype_Data is a representation of the C type G_SIGNAL_MATCH_DATA.
	Signalmatchtype_Data Signalmatchtype = 16
	// Signalmatchtype_Unblocked is a representation of the C type G_SIGNAL_MATCH_UNBLOCKED.
	Signalmatchtype_Unblocked Signalmatchtype = 32
)

// Typedebugflags is a representation of the C type GTypeDebugFlags.
type Typedebugflags int

const (
	// Typedebugflags_None is a representation of the C type G_TYPE_DEBUG_NONE.
	Typedebugflags_None Typedebugflags = 0
	// Typedebugflags_Objects is a representation of the C type G_TYPE_DEBUG_OBJECTS.
	Typedebugflags_Objects Typedebugflags = 1
	// Typedebugflags_Signals is a representation of the C type G_TYPE_DEBUG_SIGNALS.
	Typedebugflags_Signals Typedebugflags = 2
	// Typedebugflags_InstanceCount is a representation of the C type G_TYPE_DEBUG_INSTANCE_COUNT.
	Typedebugflags_InstanceCount Typedebugflags = 4
	// Typedebugflags_Mask is a representation of the C type G_TYPE_DEBUG_MASK.
	Typedebugflags_Mask Typedebugflags = 7
)

// Typeflags is a representation of the C type GTypeFlags.
type Typeflags int

const (
	// Typeflags_Abstract is a representation of the C type G_TYPE_FLAG_ABSTRACT.
	Typeflags_Abstract Typeflags = 16
	// Typeflags_ValueAbstract is a representation of the C type G_TYPE_FLAG_VALUE_ABSTRACT.
	Typeflags_ValueAbstract Typeflags = 32
)

// Typefundamentalflags is a representation of the C type GTypeFundamentalFlags.
type Typefundamentalflags int

const (
	// Typefundamentalflags_Classed is a representation of the C type G_TYPE_FLAG_CLASSED.
	Typefundamentalflags_Classed Typefundamentalflags = 1
	// Typefundamentalflags_Instantiatable is a representation of the C type G_TYPE_FLAG_INSTANTIATABLE.
	Typefundamentalflags_Instantiatable Typefundamentalflags = 2
	// Typefundamentalflags_Derivable is a representation of the C type G_TYPE_FLAG_DERIVABLE.
	Typefundamentalflags_Derivable Typefundamentalflags = 4
	// Typefundamentalflags_DeepDerivable is a representation of the C type G_TYPE_FLAG_DEEP_DERIVABLE.
	Typefundamentalflags_DeepDerivable Typefundamentalflags = 8
)
