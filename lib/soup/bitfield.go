// Code generated - DO NOT EDIT.

package soup

// Cacheability is a representation of the C type SoupCacheability.
type Cacheability int

const (
	// Cacheability_Cacheable is a representation of the C type SOUP_CACHE_CACHEABLE.
	Cacheability_Cacheable Cacheability = 1
	// Cacheability_Uncacheable is a representation of the C type SOUP_CACHE_UNCACHEABLE.
	Cacheability_Uncacheable Cacheability = 2
	// Cacheability_Invalidates is a representation of the C type SOUP_CACHE_INVALIDATES.
	Cacheability_Invalidates Cacheability = 4
	// Cacheability_Validates is a representation of the C type SOUP_CACHE_VALIDATES.
	Cacheability_Validates Cacheability = 8
)

// Expectation is a representation of the C type SoupExpectation.
type Expectation int

const (
	// Expectation_Unrecognized is a representation of the C type SOUP_EXPECTATION_UNRECOGNIZED.
	Expectation_Unrecognized Expectation = 1
	// Expectation_Continue is a representation of the C type SOUP_EXPECTATION_CONTINUE.
	Expectation_Continue Expectation = 2
)

// MessageFlags is a representation of the C type SoupMessageFlags.
type MessageFlags int

const (
	// MessageFlags_NoRedirect is a representation of the C type SOUP_MESSAGE_NO_REDIRECT.
	MessageFlags_NoRedirect MessageFlags = 2
	// MessageFlags_CanRebuild is a representation of the C type SOUP_MESSAGE_CAN_REBUILD.
	MessageFlags_CanRebuild MessageFlags = 4
	// MessageFlags_OverwriteChunks is a representation of the C type SOUP_MESSAGE_OVERWRITE_CHUNKS.
	MessageFlags_OverwriteChunks MessageFlags = 8
	// MessageFlags_ContentDecoded is a representation of the C type SOUP_MESSAGE_CONTENT_DECODED.
	MessageFlags_ContentDecoded MessageFlags = 16
	// MessageFlags_CertificateTrusted is a representation of the C type SOUP_MESSAGE_CERTIFICATE_TRUSTED.
	MessageFlags_CertificateTrusted MessageFlags = 32
	// MessageFlags_NewConnection is a representation of the C type SOUP_MESSAGE_NEW_CONNECTION.
	MessageFlags_NewConnection MessageFlags = 64
	// MessageFlags_Idempotent is a representation of the C type SOUP_MESSAGE_IDEMPOTENT.
	MessageFlags_Idempotent MessageFlags = 128
	// MessageFlags_IgnoreConnectionLimits is a representation of the C type SOUP_MESSAGE_IGNORE_CONNECTION_LIMITS.
	MessageFlags_IgnoreConnectionLimits MessageFlags = 256
	// MessageFlags_DoNotUseAuthCache is a representation of the C type SOUP_MESSAGE_DO_NOT_USE_AUTH_CACHE.
	MessageFlags_DoNotUseAuthCache MessageFlags = 512
)

// ServerListenOptions is a representation of the C type SoupServerListenOptions.
//
// since 2.48
type ServerListenOptions int

const (
	// ServerListenOptions_Https is a representation of the C type SOUP_SERVER_LISTEN_HTTPS.
	ServerListenOptions_Https ServerListenOptions = 1
	// ServerListenOptions_Ipv4Only is a representation of the C type SOUP_SERVER_LISTEN_IPV4_ONLY.
	ServerListenOptions_Ipv4Only ServerListenOptions = 2
	// ServerListenOptions_Ipv6Only is a representation of the C type SOUP_SERVER_LISTEN_IPV6_ONLY.
	ServerListenOptions_Ipv6Only ServerListenOptions = 4
)
