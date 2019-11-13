// Code generated - DO NOT EDIT.

package gobject

// Bindingflags is a representation of the C type BindingFlags.
//
// since 2.26
type Bindingflags int

const (
	Bindingflags_Default       Bindingflags = 0
	Bindingflags_Bidirectional Bindingflags = 1
	Bindingflags_SyncCreate    Bindingflags = 2
	Bindingflags_InvertBoolean Bindingflags = 4
)

// Connectflags is a representation of the C type ConnectFlags.
type Connectflags int

const (
	Connectflags_After   Connectflags = 1
	Connectflags_Swapped Connectflags = 2
)

// Paramflags is a representation of the C type ParamFlags.
type Paramflags int

const (
	Paramflags_Readable       Paramflags = 1
	Paramflags_Writable       Paramflags = 2
	Paramflags_Readwrite      Paramflags = 3
	Paramflags_Construct      Paramflags = 4
	Paramflags_ConstructOnly  Paramflags = 8
	Paramflags_LaxValidation  Paramflags = 16
	Paramflags_StaticName     Paramflags = 32
	Paramflags_Private        Paramflags = 32
	Paramflags_StaticNick     Paramflags = 64
	Paramflags_StaticBlurb    Paramflags = 128
	Paramflags_ExplicitNotify Paramflags = 1073741824
	Paramflags_Deprecated     Paramflags = 2147483648
)

// Signalflags is a representation of the C type SignalFlags.
type Signalflags int

const (
	Signalflags_RunFirst    Signalflags = 1
	Signalflags_RunLast     Signalflags = 2
	Signalflags_RunCleanup  Signalflags = 4
	Signalflags_NoRecurse   Signalflags = 8
	Signalflags_Detailed    Signalflags = 16
	Signalflags_Action      Signalflags = 32
	Signalflags_NoHooks     Signalflags = 64
	Signalflags_MustCollect Signalflags = 128
	Signalflags_Deprecated  Signalflags = 256
)

// Signalmatchtype is a representation of the C type SignalMatchType.
type Signalmatchtype int

const (
	Signalmatchtype_Id        Signalmatchtype = 1
	Signalmatchtype_Detail    Signalmatchtype = 2
	Signalmatchtype_Closure   Signalmatchtype = 4
	Signalmatchtype_Func      Signalmatchtype = 8
	Signalmatchtype_Data      Signalmatchtype = 16
	Signalmatchtype_Unblocked Signalmatchtype = 32
)

// Typedebugflags is a representation of the C type TypeDebugFlags.
type Typedebugflags int

const (
	Typedebugflags_None          Typedebugflags = 0
	Typedebugflags_Objects       Typedebugflags = 1
	Typedebugflags_Signals       Typedebugflags = 2
	Typedebugflags_InstanceCount Typedebugflags = 4
	Typedebugflags_Mask          Typedebugflags = 7
)

// Typeflags is a representation of the C type TypeFlags.
type Typeflags int

const (
	Typeflags_Abstract      Typeflags = 16
	Typeflags_ValueAbstract Typeflags = 32
)

// Typefundamentalflags is a representation of the C type TypeFundamentalFlags.
type Typefundamentalflags int

const (
	Typefundamentalflags_Classed        Typefundamentalflags = 1
	Typefundamentalflags_Instantiatable Typefundamentalflags = 2
	Typefundamentalflags_Derivable      Typefundamentalflags = 4
	Typefundamentalflags_DeepDerivable  Typefundamentalflags = 8
)
