// Code generated - DO NOT EDIT.

package gobject

type BindingFlags uint32

const (
	BindingFlags_Default       BindingFlags = int64(0)
	BindingFlags_Bidirectional BindingFlags = int64(1)
	BindingFlags_SyncCreate    BindingFlags = int64(2)
	BindingFlags_InvertBoolean BindingFlags = int64(4)
)

type ConnectFlags uint32

const (
	ConnectFlags_After   ConnectFlags = int64(1)
	ConnectFlags_Swapped ConnectFlags = int64(2)
)

type ParamFlags uint32

const (
	ParamFlags_Readable       ParamFlags = int64(1)
	ParamFlags_Writable       ParamFlags = int64(2)
	ParamFlags_Readwrite      ParamFlags = int64(3)
	ParamFlags_Construct      ParamFlags = int64(4)
	ParamFlags_ConstructOnly  ParamFlags = int64(8)
	ParamFlags_LaxValidation  ParamFlags = int64(16)
	ParamFlags_StaticName     ParamFlags = int64(32)
	ParamFlags_Private        ParamFlags = int64(32)
	ParamFlags_StaticNick     ParamFlags = int64(64)
	ParamFlags_StaticBlurb    ParamFlags = int64(128)
	ParamFlags_ExplicitNotify ParamFlags = int64(1073741824)
	ParamFlags_Deprecated     ParamFlags = int64(2147483648)
)

type SignalFlags uint32

const (
	SignalFlags_RunFirst    SignalFlags = int64(1)
	SignalFlags_RunLast     SignalFlags = int64(2)
	SignalFlags_RunCleanup  SignalFlags = int64(4)
	SignalFlags_NoRecurse   SignalFlags = int64(8)
	SignalFlags_Detailed    SignalFlags = int64(16)
	SignalFlags_Action      SignalFlags = int64(32)
	SignalFlags_NoHooks     SignalFlags = int64(64)
	SignalFlags_MustCollect SignalFlags = int64(128)
	SignalFlags_Deprecated  SignalFlags = int64(256)
)

type SignalMatchType uint32

const (
	SignalMatchType_Id        SignalMatchType = int64(1)
	SignalMatchType_Detail    SignalMatchType = int64(2)
	SignalMatchType_Closure   SignalMatchType = int64(4)
	SignalMatchType_Func      SignalMatchType = int64(8)
	SignalMatchType_Data      SignalMatchType = int64(16)
	SignalMatchType_Unblocked SignalMatchType = int64(32)
)

type TypeDebugFlags uint32

const (
	TypeDebugFlags_None          TypeDebugFlags = int64(0)
	TypeDebugFlags_Objects       TypeDebugFlags = int64(1)
	TypeDebugFlags_Signals       TypeDebugFlags = int64(2)
	TypeDebugFlags_InstanceCount TypeDebugFlags = int64(4)
	TypeDebugFlags_Mask          TypeDebugFlags = int64(7)
)

type TypeFlags uint32

const (
	TypeFlags_Abstract      TypeFlags = int64(16)
	TypeFlags_ValueAbstract TypeFlags = int64(32)
)

type TypeFundamentalFlags uint32

const (
	TypeFundamentalFlags_Classed        TypeFundamentalFlags = int64(1)
	TypeFundamentalFlags_Instantiatable TypeFundamentalFlags = int64(2)
	TypeFundamentalFlags_Derivable      TypeFundamentalFlags = int64(4)
	TypeFundamentalFlags_DeepDerivable  TypeFundamentalFlags = int64(8)
)
