// Code generated - DO NOT EDIT.

package soup

type Cacheability uint32

const (
	Cacheability_Cacheable   Cacheability = int64(1)
	Cacheability_Uncacheable Cacheability = int64(2)
	Cacheability_Invalidates Cacheability = int64(4)
	Cacheability_Validates   Cacheability = int64(8)
)

type Expectation uint32

const (
	Expectation_Unrecognized Expectation = int64(1)
	Expectation_Continue     Expectation = int64(2)
)

type MessageFlags uint32

const (
	MessageFlags_NoRedirect             MessageFlags = int64(2)
	MessageFlags_CanRebuild             MessageFlags = int64(4)
	MessageFlags_OverwriteChunks        MessageFlags = int64(8)
	MessageFlags_ContentDecoded         MessageFlags = int64(16)
	MessageFlags_CertificateTrusted     MessageFlags = int64(32)
	MessageFlags_NewConnection          MessageFlags = int64(64)
	MessageFlags_Idempotent             MessageFlags = int64(128)
	MessageFlags_IgnoreConnectionLimits MessageFlags = int64(256)
	MessageFlags_DoNotUseAuthCache      MessageFlags = int64(512)
)

type ServerListenOptions uint32

const (
	ServerListenOptions_Https    ServerListenOptions = int64(1)
	ServerListenOptions_Ipv4Only ServerListenOptions = int64(2)
	ServerListenOptions_Ipv6Only ServerListenOptions = int64(4)
)
