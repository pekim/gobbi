// Code generated - DO NOT EDIT.

package soup

// Cacheability is a representation of the C type Cacheability.
type Cacheability int

const (
	Cacheability_Cacheable   Cacheability = 1
	Cacheability_Uncacheable Cacheability = 2
	Cacheability_Invalidates Cacheability = 4
	Cacheability_Validates   Cacheability = 8
)

// Expectation is a representation of the C type Expectation.
type Expectation int

const (
	Expectation_Unrecognized Expectation = 1
	Expectation_Continue     Expectation = 2
)

// Messageflags is a representation of the C type MessageFlags.
type Messageflags int

const (
	Messageflags_NoRedirect             Messageflags = 2
	Messageflags_CanRebuild             Messageflags = 4
	Messageflags_OverwriteChunks        Messageflags = 8
	Messageflags_ContentDecoded         Messageflags = 16
	Messageflags_CertificateTrusted     Messageflags = 32
	Messageflags_NewConnection          Messageflags = 64
	Messageflags_Idempotent             Messageflags = 128
	Messageflags_IgnoreConnectionLimits Messageflags = 256
	Messageflags_DoNotUseAuthCache      Messageflags = 512
)

// Serverlistenoptions is a representation of the C type ServerListenOptions.
//
// since 2.48
type Serverlistenoptions int

const (
	Serverlistenoptions_Https    Serverlistenoptions = 1
	Serverlistenoptions_Ipv4Only Serverlistenoptions = 2
	Serverlistenoptions_Ipv6Only Serverlistenoptions = 4
)
