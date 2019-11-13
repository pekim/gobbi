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

// Messageflags is a representation of the C type SoupMessageFlags.
type Messageflags int

const (
	// Messageflags_NoRedirect is a representation of the C type SOUP_MESSAGE_NO_REDIRECT.
	Messageflags_NoRedirect Messageflags = 2
	// Messageflags_CanRebuild is a representation of the C type SOUP_MESSAGE_CAN_REBUILD.
	Messageflags_CanRebuild Messageflags = 4
	// Messageflags_OverwriteChunks is a representation of the C type SOUP_MESSAGE_OVERWRITE_CHUNKS.
	Messageflags_OverwriteChunks Messageflags = 8
	// Messageflags_ContentDecoded is a representation of the C type SOUP_MESSAGE_CONTENT_DECODED.
	Messageflags_ContentDecoded Messageflags = 16
	// Messageflags_CertificateTrusted is a representation of the C type SOUP_MESSAGE_CERTIFICATE_TRUSTED.
	Messageflags_CertificateTrusted Messageflags = 32
	// Messageflags_NewConnection is a representation of the C type SOUP_MESSAGE_NEW_CONNECTION.
	Messageflags_NewConnection Messageflags = 64
	// Messageflags_Idempotent is a representation of the C type SOUP_MESSAGE_IDEMPOTENT.
	Messageflags_Idempotent Messageflags = 128
	// Messageflags_IgnoreConnectionLimits is a representation of the C type SOUP_MESSAGE_IGNORE_CONNECTION_LIMITS.
	Messageflags_IgnoreConnectionLimits Messageflags = 256
	// Messageflags_DoNotUseAuthCache is a representation of the C type SOUP_MESSAGE_DO_NOT_USE_AUTH_CACHE.
	Messageflags_DoNotUseAuthCache Messageflags = 512
)

// Serverlistenoptions is a representation of the C type SoupServerListenOptions.
//
// since 2.48
type Serverlistenoptions int

const (
	// Serverlistenoptions_Https is a representation of the C type SOUP_SERVER_LISTEN_HTTPS.
	Serverlistenoptions_Https Serverlistenoptions = 1
	// Serverlistenoptions_Ipv4Only is a representation of the C type SOUP_SERVER_LISTEN_IPV4_ONLY.
	Serverlistenoptions_Ipv4Only Serverlistenoptions = 2
	// Serverlistenoptions_Ipv6Only is a representation of the C type SOUP_SERVER_LISTEN_IPV6_ONLY.
	Serverlistenoptions_Ipv6Only Serverlistenoptions = 4
)
