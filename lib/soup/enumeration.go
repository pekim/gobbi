// Code generated - DO NOT EDIT.

package soup

// AddressFamily is a representation of the C type SoupAddressFamily.
//
type AddressFamily int

const (
	// AddressFamily_Invalid is a representation of the C type SOUP_ADDRESS_FAMILY_INVALID.
	AddressFamily_Invalid AddressFamily = -1
	// AddressFamily_Ipv4 is a representation of the C type SOUP_ADDRESS_FAMILY_IPV4.
	AddressFamily_Ipv4 AddressFamily = 2
	// AddressFamily_Ipv6 is a representation of the C type SOUP_ADDRESS_FAMILY_IPV6.
	AddressFamily_Ipv6 AddressFamily = 10
)

// CacheResponse is a representation of the C type SoupCacheResponse.
//
type CacheResponse int

const (
	// CacheResponse_Fresh is a representation of the C type SOUP_CACHE_RESPONSE_FRESH.
	CacheResponse_Fresh CacheResponse = 0
	// CacheResponse_NeedsValidation is a representation of the C type SOUP_CACHE_RESPONSE_NEEDS_VALIDATION.
	CacheResponse_NeedsValidation CacheResponse = 1
	// CacheResponse_Stale is a representation of the C type SOUP_CACHE_RESPONSE_STALE.
	CacheResponse_Stale CacheResponse = 2
)

// CacheType is a representation of the C type SoupCacheType.
//
// since 2.34
type CacheType int

const (
	// CacheType_SingleUser is a representation of the C type SOUP_CACHE_SINGLE_USER.
	CacheType_SingleUser CacheType = 0
	// CacheType_Shared is a representation of the C type SOUP_CACHE_SHARED.
	CacheType_Shared CacheType = 1
)

// ConnectionState is a representation of the C type SoupConnectionState.
//
type ConnectionState int

const (
	// ConnectionState_New is a representation of the C type SOUP_CONNECTION_NEW.
	ConnectionState_New ConnectionState = 0
	// ConnectionState_Connecting is a representation of the C type SOUP_CONNECTION_CONNECTING.
	ConnectionState_Connecting ConnectionState = 1
	// ConnectionState_Idle is a representation of the C type SOUP_CONNECTION_IDLE.
	ConnectionState_Idle ConnectionState = 2
	// ConnectionState_InUse is a representation of the C type SOUP_CONNECTION_IN_USE.
	ConnectionState_InUse ConnectionState = 3
	// ConnectionState_RemoteDisconnected is a representation of the C type SOUP_CONNECTION_REMOTE_DISCONNECTED.
	ConnectionState_RemoteDisconnected ConnectionState = 4
	// ConnectionState_Disconnected is a representation of the C type SOUP_CONNECTION_DISCONNECTED.
	ConnectionState_Disconnected ConnectionState = 5
)

// CookieJarAcceptPolicy is a representation of the C type SoupCookieJarAcceptPolicy.
//
// since 2.30
type CookieJarAcceptPolicy int

const (
	// CookieJarAcceptPolicy_Always is a representation of the C type SOUP_COOKIE_JAR_ACCEPT_ALWAYS.
	CookieJarAcceptPolicy_Always CookieJarAcceptPolicy = 0
	// CookieJarAcceptPolicy_Never is a representation of the C type SOUP_COOKIE_JAR_ACCEPT_NEVER.
	CookieJarAcceptPolicy_Never CookieJarAcceptPolicy = 1
	// CookieJarAcceptPolicy_NoThirdParty is a representation of the C type SOUP_COOKIE_JAR_ACCEPT_NO_THIRD_PARTY.
	CookieJarAcceptPolicy_NoThirdParty CookieJarAcceptPolicy = 2
)

// DateFormat is a representation of the C type SoupDateFormat.
//
type DateFormat int

const (
	// DateFormat_Http is a representation of the C type SOUP_DATE_HTTP.
	DateFormat_Http DateFormat = 1
	// DateFormat_Cookie is a representation of the C type SOUP_DATE_COOKIE.
	DateFormat_Cookie DateFormat = 2
	// DateFormat_Rfc2822 is a representation of the C type SOUP_DATE_RFC2822.
	DateFormat_Rfc2822 DateFormat = 3
	// DateFormat_Iso8601Compact is a representation of the C type SOUP_DATE_ISO8601_COMPACT.
	DateFormat_Iso8601Compact DateFormat = 4
	// DateFormat_Iso8601Full is a representation of the C type SOUP_DATE_ISO8601_FULL.
	DateFormat_Iso8601Full DateFormat = 5
	// DateFormat_Iso8601 is a representation of the C type SOUP_DATE_ISO8601.
	DateFormat_Iso8601 DateFormat = 5
	// DateFormat_Iso8601Xmlrpc is a representation of the C type SOUP_DATE_ISO8601_XMLRPC.
	DateFormat_Iso8601Xmlrpc DateFormat = 6
)

// Encoding is a representation of the C type SoupEncoding.
//
type Encoding int

const (
	// Encoding_Unrecognized is a representation of the C type SOUP_ENCODING_UNRECOGNIZED.
	Encoding_Unrecognized Encoding = 0
	// Encoding_None is a representation of the C type SOUP_ENCODING_NONE.
	Encoding_None Encoding = 1
	// Encoding_ContentLength is a representation of the C type SOUP_ENCODING_CONTENT_LENGTH.
	Encoding_ContentLength Encoding = 2
	// Encoding_Eof is a representation of the C type SOUP_ENCODING_EOF.
	Encoding_Eof Encoding = 3
	// Encoding_Chunked is a representation of the C type SOUP_ENCODING_CHUNKED.
	Encoding_Chunked Encoding = 4
	// Encoding_Byteranges is a representation of the C type SOUP_ENCODING_BYTERANGES.
	Encoding_Byteranges Encoding = 5
)

// HTTPVersion is a representation of the C type SoupHTTPVersion.
//
type HTTPVersion int

const (
	// HTTPVersion_Http10 is a representation of the C type SOUP_HTTP_1_0.
	HTTPVersion_Http10 HTTPVersion = 0
	// HTTPVersion_Http11 is a representation of the C type SOUP_HTTP_1_1.
	HTTPVersion_Http11 HTTPVersion = 1
)

// KnownStatusCode is a representation of the C type SoupKnownStatusCode.
//
type KnownStatusCode int

const (
	// KnownStatusCode_None is a representation of the C type SOUP_KNOWN_STATUS_CODE_NONE.
	KnownStatusCode_None KnownStatusCode = 0
	// KnownStatusCode_Cancelled is a representation of the C type SOUP_KNOWN_STATUS_CODE_CANCELLED.
	KnownStatusCode_Cancelled KnownStatusCode = 1
	// KnownStatusCode_CantResolve is a representation of the C type SOUP_KNOWN_STATUS_CODE_CANT_RESOLVE.
	KnownStatusCode_CantResolve KnownStatusCode = 2
	// KnownStatusCode_CantResolveProxy is a representation of the C type SOUP_KNOWN_STATUS_CODE_CANT_RESOLVE_PROXY.
	KnownStatusCode_CantResolveProxy KnownStatusCode = 3
	// KnownStatusCode_CantConnect is a representation of the C type SOUP_KNOWN_STATUS_CODE_CANT_CONNECT.
	KnownStatusCode_CantConnect KnownStatusCode = 4
	// KnownStatusCode_CantConnectProxy is a representation of the C type SOUP_KNOWN_STATUS_CODE_CANT_CONNECT_PROXY.
	KnownStatusCode_CantConnectProxy KnownStatusCode = 5
	// KnownStatusCode_SslFailed is a representation of the C type SOUP_KNOWN_STATUS_CODE_SSL_FAILED.
	KnownStatusCode_SslFailed KnownStatusCode = 6
	// KnownStatusCode_IoError is a representation of the C type SOUP_KNOWN_STATUS_CODE_IO_ERROR.
	KnownStatusCode_IoError KnownStatusCode = 7
	// KnownStatusCode_Malformed is a representation of the C type SOUP_KNOWN_STATUS_CODE_MALFORMED.
	KnownStatusCode_Malformed KnownStatusCode = 8
	// KnownStatusCode_TryAgain is a representation of the C type SOUP_KNOWN_STATUS_CODE_TRY_AGAIN.
	KnownStatusCode_TryAgain KnownStatusCode = 9
	// KnownStatusCode_TooManyRedirects is a representation of the C type SOUP_KNOWN_STATUS_CODE_TOO_MANY_REDIRECTS.
	KnownStatusCode_TooManyRedirects KnownStatusCode = 10
	// KnownStatusCode_TlsFailed is a representation of the C type SOUP_KNOWN_STATUS_CODE_TLS_FAILED.
	KnownStatusCode_TlsFailed KnownStatusCode = 11
	// KnownStatusCode_Continue is a representation of the C type SOUP_KNOWN_STATUS_CODE_CONTINUE.
	KnownStatusCode_Continue KnownStatusCode = 100
	// KnownStatusCode_SwitchingProtocols is a representation of the C type SOUP_KNOWN_STATUS_CODE_SWITCHING_PROTOCOLS.
	KnownStatusCode_SwitchingProtocols KnownStatusCode = 101
	// KnownStatusCode_Processing is a representation of the C type SOUP_KNOWN_STATUS_CODE_PROCESSING.
	KnownStatusCode_Processing KnownStatusCode = 102
	// KnownStatusCode_Ok is a representation of the C type SOUP_KNOWN_STATUS_CODE_OK.
	KnownStatusCode_Ok KnownStatusCode = 200
	// KnownStatusCode_Created is a representation of the C type SOUP_KNOWN_STATUS_CODE_CREATED.
	KnownStatusCode_Created KnownStatusCode = 201
	// KnownStatusCode_Accepted is a representation of the C type SOUP_KNOWN_STATUS_CODE_ACCEPTED.
	KnownStatusCode_Accepted KnownStatusCode = 202
	// KnownStatusCode_NonAuthoritative is a representation of the C type SOUP_KNOWN_STATUS_CODE_NON_AUTHORITATIVE.
	KnownStatusCode_NonAuthoritative KnownStatusCode = 203
	// KnownStatusCode_NoContent is a representation of the C type SOUP_KNOWN_STATUS_CODE_NO_CONTENT.
	KnownStatusCode_NoContent KnownStatusCode = 204
	// KnownStatusCode_ResetContent is a representation of the C type SOUP_KNOWN_STATUS_CODE_RESET_CONTENT.
	KnownStatusCode_ResetContent KnownStatusCode = 205
	// KnownStatusCode_PartialContent is a representation of the C type SOUP_KNOWN_STATUS_CODE_PARTIAL_CONTENT.
	KnownStatusCode_PartialContent KnownStatusCode = 206
	// KnownStatusCode_MultiStatus is a representation of the C type SOUP_KNOWN_STATUS_CODE_MULTI_STATUS.
	KnownStatusCode_MultiStatus KnownStatusCode = 207
	// KnownStatusCode_MultipleChoices is a representation of the C type SOUP_KNOWN_STATUS_CODE_MULTIPLE_CHOICES.
	KnownStatusCode_MultipleChoices KnownStatusCode = 300
	// KnownStatusCode_MovedPermanently is a representation of the C type SOUP_KNOWN_STATUS_CODE_MOVED_PERMANENTLY.
	KnownStatusCode_MovedPermanently KnownStatusCode = 301
	// KnownStatusCode_Found is a representation of the C type SOUP_KNOWN_STATUS_CODE_FOUND.
	KnownStatusCode_Found KnownStatusCode = 302
	// KnownStatusCode_MovedTemporarily is a representation of the C type SOUP_KNOWN_STATUS_CODE_MOVED_TEMPORARILY.
	KnownStatusCode_MovedTemporarily KnownStatusCode = 302
	// KnownStatusCode_SeeOther is a representation of the C type SOUP_KNOWN_STATUS_CODE_SEE_OTHER.
	KnownStatusCode_SeeOther KnownStatusCode = 303
	// KnownStatusCode_NotModified is a representation of the C type SOUP_KNOWN_STATUS_CODE_NOT_MODIFIED.
	KnownStatusCode_NotModified KnownStatusCode = 304
	// KnownStatusCode_UseProxy is a representation of the C type SOUP_KNOWN_STATUS_CODE_USE_PROXY.
	KnownStatusCode_UseProxy KnownStatusCode = 305
	// KnownStatusCode_NotAppearingInThisProtocol is a representation of the C type SOUP_KNOWN_STATUS_CODE_NOT_APPEARING_IN_THIS_PROTOCOL.
	KnownStatusCode_NotAppearingInThisProtocol KnownStatusCode = 306
	// KnownStatusCode_TemporaryRedirect is a representation of the C type SOUP_KNOWN_STATUS_CODE_TEMPORARY_REDIRECT.
	KnownStatusCode_TemporaryRedirect KnownStatusCode = 307
	// KnownStatusCode_BadRequest is a representation of the C type SOUP_KNOWN_STATUS_CODE_BAD_REQUEST.
	KnownStatusCode_BadRequest KnownStatusCode = 400
	// KnownStatusCode_Unauthorized is a representation of the C type SOUP_KNOWN_STATUS_CODE_UNAUTHORIZED.
	KnownStatusCode_Unauthorized KnownStatusCode = 401
	// KnownStatusCode_PaymentRequired is a representation of the C type SOUP_KNOWN_STATUS_CODE_PAYMENT_REQUIRED.
	KnownStatusCode_PaymentRequired KnownStatusCode = 402
	// KnownStatusCode_Forbidden is a representation of the C type SOUP_KNOWN_STATUS_CODE_FORBIDDEN.
	KnownStatusCode_Forbidden KnownStatusCode = 403
	// KnownStatusCode_NotFound is a representation of the C type SOUP_KNOWN_STATUS_CODE_NOT_FOUND.
	KnownStatusCode_NotFound KnownStatusCode = 404
	// KnownStatusCode_MethodNotAllowed is a representation of the C type SOUP_KNOWN_STATUS_CODE_METHOD_NOT_ALLOWED.
	KnownStatusCode_MethodNotAllowed KnownStatusCode = 405
	// KnownStatusCode_NotAcceptable is a representation of the C type SOUP_KNOWN_STATUS_CODE_NOT_ACCEPTABLE.
	KnownStatusCode_NotAcceptable KnownStatusCode = 406
	// KnownStatusCode_ProxyAuthenticationRequired is a representation of the C type SOUP_KNOWN_STATUS_CODE_PROXY_AUTHENTICATION_REQUIRED.
	KnownStatusCode_ProxyAuthenticationRequired KnownStatusCode = 407
	// KnownStatusCode_ProxyUnauthorized is a representation of the C type SOUP_KNOWN_STATUS_CODE_PROXY_UNAUTHORIZED.
	KnownStatusCode_ProxyUnauthorized KnownStatusCode = 407
	// KnownStatusCode_RequestTimeout is a representation of the C type SOUP_KNOWN_STATUS_CODE_REQUEST_TIMEOUT.
	KnownStatusCode_RequestTimeout KnownStatusCode = 408
	// KnownStatusCode_Conflict is a representation of the C type SOUP_KNOWN_STATUS_CODE_CONFLICT.
	KnownStatusCode_Conflict KnownStatusCode = 409
	// KnownStatusCode_Gone is a representation of the C type SOUP_KNOWN_STATUS_CODE_GONE.
	KnownStatusCode_Gone KnownStatusCode = 410
	// KnownStatusCode_LengthRequired is a representation of the C type SOUP_KNOWN_STATUS_CODE_LENGTH_REQUIRED.
	KnownStatusCode_LengthRequired KnownStatusCode = 411
	// KnownStatusCode_PreconditionFailed is a representation of the C type SOUP_KNOWN_STATUS_CODE_PRECONDITION_FAILED.
	KnownStatusCode_PreconditionFailed KnownStatusCode = 412
	// KnownStatusCode_RequestEntityTooLarge is a representation of the C type SOUP_KNOWN_STATUS_CODE_REQUEST_ENTITY_TOO_LARGE.
	KnownStatusCode_RequestEntityTooLarge KnownStatusCode = 413
	// KnownStatusCode_RequestUriTooLong is a representation of the C type SOUP_KNOWN_STATUS_CODE_REQUEST_URI_TOO_LONG.
	KnownStatusCode_RequestUriTooLong KnownStatusCode = 414
	// KnownStatusCode_UnsupportedMediaType is a representation of the C type SOUP_KNOWN_STATUS_CODE_UNSUPPORTED_MEDIA_TYPE.
	KnownStatusCode_UnsupportedMediaType KnownStatusCode = 415
	// KnownStatusCode_RequestedRangeNotSatisfiable is a representation of the C type SOUP_KNOWN_STATUS_CODE_REQUESTED_RANGE_NOT_SATISFIABLE.
	KnownStatusCode_RequestedRangeNotSatisfiable KnownStatusCode = 416
	// KnownStatusCode_InvalidRange is a representation of the C type SOUP_KNOWN_STATUS_CODE_INVALID_RANGE.
	KnownStatusCode_InvalidRange KnownStatusCode = 416
	// KnownStatusCode_ExpectationFailed is a representation of the C type SOUP_KNOWN_STATUS_CODE_EXPECTATION_FAILED.
	KnownStatusCode_ExpectationFailed KnownStatusCode = 417
	// KnownStatusCode_UnprocessableEntity is a representation of the C type SOUP_KNOWN_STATUS_CODE_UNPROCESSABLE_ENTITY.
	KnownStatusCode_UnprocessableEntity KnownStatusCode = 422
	// KnownStatusCode_Locked is a representation of the C type SOUP_KNOWN_STATUS_CODE_LOCKED.
	KnownStatusCode_Locked KnownStatusCode = 423
	// KnownStatusCode_FailedDependency is a representation of the C type SOUP_KNOWN_STATUS_CODE_FAILED_DEPENDENCY.
	KnownStatusCode_FailedDependency KnownStatusCode = 424
	// KnownStatusCode_InternalServerError is a representation of the C type SOUP_KNOWN_STATUS_CODE_INTERNAL_SERVER_ERROR.
	KnownStatusCode_InternalServerError KnownStatusCode = 500
	// KnownStatusCode_NotImplemented is a representation of the C type SOUP_KNOWN_STATUS_CODE_NOT_IMPLEMENTED.
	KnownStatusCode_NotImplemented KnownStatusCode = 501
	// KnownStatusCode_BadGateway is a representation of the C type SOUP_KNOWN_STATUS_CODE_BAD_GATEWAY.
	KnownStatusCode_BadGateway KnownStatusCode = 502
	// KnownStatusCode_ServiceUnavailable is a representation of the C type SOUP_KNOWN_STATUS_CODE_SERVICE_UNAVAILABLE.
	KnownStatusCode_ServiceUnavailable KnownStatusCode = 503
	// KnownStatusCode_GatewayTimeout is a representation of the C type SOUP_KNOWN_STATUS_CODE_GATEWAY_TIMEOUT.
	KnownStatusCode_GatewayTimeout KnownStatusCode = 504
	// KnownStatusCode_HttpVersionNotSupported is a representation of the C type SOUP_KNOWN_STATUS_CODE_HTTP_VERSION_NOT_SUPPORTED.
	KnownStatusCode_HttpVersionNotSupported KnownStatusCode = 505
	// KnownStatusCode_InsufficientStorage is a representation of the C type SOUP_KNOWN_STATUS_CODE_INSUFFICIENT_STORAGE.
	KnownStatusCode_InsufficientStorage KnownStatusCode = 507
	// KnownStatusCode_NotExtended is a representation of the C type SOUP_KNOWN_STATUS_CODE_NOT_EXTENDED.
	KnownStatusCode_NotExtended KnownStatusCode = 510
)

// LoggerLogLevel is a representation of the C type SoupLoggerLogLevel.
//
type LoggerLogLevel int

const (
	// LoggerLogLevel_None is a representation of the C type SOUP_LOGGER_LOG_NONE.
	LoggerLogLevel_None LoggerLogLevel = 0
	// LoggerLogLevel_Minimal is a representation of the C type SOUP_LOGGER_LOG_MINIMAL.
	LoggerLogLevel_Minimal LoggerLogLevel = 1
	// LoggerLogLevel_Headers is a representation of the C type SOUP_LOGGER_LOG_HEADERS.
	LoggerLogLevel_Headers LoggerLogLevel = 2
	// LoggerLogLevel_Body is a representation of the C type SOUP_LOGGER_LOG_BODY.
	LoggerLogLevel_Body LoggerLogLevel = 3
)

// MemoryUse is a representation of the C type SoupMemoryUse.
//
type MemoryUse int

const (
	// MemoryUse_Static is a representation of the C type SOUP_MEMORY_STATIC.
	MemoryUse_Static MemoryUse = 0
	// MemoryUse_Take is a representation of the C type SOUP_MEMORY_TAKE.
	MemoryUse_Take MemoryUse = 1
	// MemoryUse_Copy is a representation of the C type SOUP_MEMORY_COPY.
	MemoryUse_Copy MemoryUse = 2
	// MemoryUse_Temporary is a representation of the C type SOUP_MEMORY_TEMPORARY.
	MemoryUse_Temporary MemoryUse = 3
)

// MessageHeadersType is a representation of the C type SoupMessageHeadersType.
//
type MessageHeadersType int

const (
	// MessageHeadersType_Request is a representation of the C type SOUP_MESSAGE_HEADERS_REQUEST.
	MessageHeadersType_Request MessageHeadersType = 0
	// MessageHeadersType_Response is a representation of the C type SOUP_MESSAGE_HEADERS_RESPONSE.
	MessageHeadersType_Response MessageHeadersType = 1
	// MessageHeadersType_Multipart is a representation of the C type SOUP_MESSAGE_HEADERS_MULTIPART.
	MessageHeadersType_Multipart MessageHeadersType = 2
)

// MessagePriority is a representation of the C type SoupMessagePriority.
//
type MessagePriority int

const (
	// MessagePriority_VeryLow is a representation of the C type SOUP_MESSAGE_PRIORITY_VERY_LOW.
	MessagePriority_VeryLow MessagePriority = 0
	// MessagePriority_Low is a representation of the C type SOUP_MESSAGE_PRIORITY_LOW.
	MessagePriority_Low MessagePriority = 1
	// MessagePriority_Normal is a representation of the C type SOUP_MESSAGE_PRIORITY_NORMAL.
	MessagePriority_Normal MessagePriority = 2
	// MessagePriority_High is a representation of the C type SOUP_MESSAGE_PRIORITY_HIGH.
	MessagePriority_High MessagePriority = 3
	// MessagePriority_VeryHigh is a representation of the C type SOUP_MESSAGE_PRIORITY_VERY_HIGH.
	MessagePriority_VeryHigh MessagePriority = 4
)

// RequestError is a representation of the C type SoupRequestError.
//
// since 2.42
type RequestError int

const (
	// RequestError_BadUri is a representation of the C type SOUP_REQUEST_ERROR_BAD_URI.
	RequestError_BadUri RequestError = 0
	// RequestError_UnsupportedUriScheme is a representation of the C type SOUP_REQUEST_ERROR_UNSUPPORTED_URI_SCHEME.
	RequestError_UnsupportedUriScheme RequestError = 1
	// RequestError_Parsing is a representation of the C type SOUP_REQUEST_ERROR_PARSING.
	RequestError_Parsing RequestError = 2
	// RequestError_Encoding is a representation of the C type SOUP_REQUEST_ERROR_ENCODING.
	RequestError_Encoding RequestError = 3
)

// RequesterError is a representation of the C type SoupRequesterError.
//
type RequesterError int

const (
	// RequesterError_BadUri is a representation of the C type SOUP_REQUESTER_ERROR_BAD_URI.
	RequesterError_BadUri RequesterError = 0
	// RequesterError_UnsupportedUriScheme is a representation of the C type SOUP_REQUESTER_ERROR_UNSUPPORTED_URI_SCHEME.
	RequesterError_UnsupportedUriScheme RequesterError = 1
)

// SocketIOStatus is a representation of the C type SoupSocketIOStatus.
//
type SocketIOStatus int

const (
	// SocketIOStatus_Ok is a representation of the C type SOUP_SOCKET_OK.
	SocketIOStatus_Ok SocketIOStatus = 0
	// SocketIOStatus_WouldBlock is a representation of the C type SOUP_SOCKET_WOULD_BLOCK.
	SocketIOStatus_WouldBlock SocketIOStatus = 1
	// SocketIOStatus_Eof is a representation of the C type SOUP_SOCKET_EOF.
	SocketIOStatus_Eof SocketIOStatus = 2
	// SocketIOStatus_Error is a representation of the C type SOUP_SOCKET_ERROR.
	SocketIOStatus_Error SocketIOStatus = 3
)

// Status is a representation of the C type SoupStatus.
//
type Status int

const (
	// Status_None is a representation of the C type SOUP_STATUS_NONE.
	Status_None Status = 0
	// Status_Cancelled is a representation of the C type SOUP_STATUS_CANCELLED.
	Status_Cancelled Status = 1
	// Status_CantResolve is a representation of the C type SOUP_STATUS_CANT_RESOLVE.
	Status_CantResolve Status = 2
	// Status_CantResolveProxy is a representation of the C type SOUP_STATUS_CANT_RESOLVE_PROXY.
	Status_CantResolveProxy Status = 3
	// Status_CantConnect is a representation of the C type SOUP_STATUS_CANT_CONNECT.
	Status_CantConnect Status = 4
	// Status_CantConnectProxy is a representation of the C type SOUP_STATUS_CANT_CONNECT_PROXY.
	Status_CantConnectProxy Status = 5
	// Status_SslFailed is a representation of the C type SOUP_STATUS_SSL_FAILED.
	Status_SslFailed Status = 6
	// Status_IoError is a representation of the C type SOUP_STATUS_IO_ERROR.
	Status_IoError Status = 7
	// Status_Malformed is a representation of the C type SOUP_STATUS_MALFORMED.
	Status_Malformed Status = 8
	// Status_TryAgain is a representation of the C type SOUP_STATUS_TRY_AGAIN.
	Status_TryAgain Status = 9
	// Status_TooManyRedirects is a representation of the C type SOUP_STATUS_TOO_MANY_REDIRECTS.
	Status_TooManyRedirects Status = 10
	// Status_TlsFailed is a representation of the C type SOUP_STATUS_TLS_FAILED.
	Status_TlsFailed Status = 11
	// Status_Continue is a representation of the C type SOUP_STATUS_CONTINUE.
	Status_Continue Status = 100
	// Status_SwitchingProtocols is a representation of the C type SOUP_STATUS_SWITCHING_PROTOCOLS.
	Status_SwitchingProtocols Status = 101
	// Status_Processing is a representation of the C type SOUP_STATUS_PROCESSING.
	Status_Processing Status = 102
	// Status_Ok is a representation of the C type SOUP_STATUS_OK.
	Status_Ok Status = 200
	// Status_Created is a representation of the C type SOUP_STATUS_CREATED.
	Status_Created Status = 201
	// Status_Accepted is a representation of the C type SOUP_STATUS_ACCEPTED.
	Status_Accepted Status = 202
	// Status_NonAuthoritative is a representation of the C type SOUP_STATUS_NON_AUTHORITATIVE.
	Status_NonAuthoritative Status = 203
	// Status_NoContent is a representation of the C type SOUP_STATUS_NO_CONTENT.
	Status_NoContent Status = 204
	// Status_ResetContent is a representation of the C type SOUP_STATUS_RESET_CONTENT.
	Status_ResetContent Status = 205
	// Status_PartialContent is a representation of the C type SOUP_STATUS_PARTIAL_CONTENT.
	Status_PartialContent Status = 206
	// Status_MultiStatus is a representation of the C type SOUP_STATUS_MULTI_STATUS.
	Status_MultiStatus Status = 207
	// Status_MultipleChoices is a representation of the C type SOUP_STATUS_MULTIPLE_CHOICES.
	Status_MultipleChoices Status = 300
	// Status_MovedPermanently is a representation of the C type SOUP_STATUS_MOVED_PERMANENTLY.
	Status_MovedPermanently Status = 301
	// Status_Found is a representation of the C type SOUP_STATUS_FOUND.
	Status_Found Status = 302
	// Status_MovedTemporarily is a representation of the C type SOUP_STATUS_MOVED_TEMPORARILY.
	Status_MovedTemporarily Status = 302
	// Status_SeeOther is a representation of the C type SOUP_STATUS_SEE_OTHER.
	Status_SeeOther Status = 303
	// Status_NotModified is a representation of the C type SOUP_STATUS_NOT_MODIFIED.
	Status_NotModified Status = 304
	// Status_UseProxy is a representation of the C type SOUP_STATUS_USE_PROXY.
	Status_UseProxy Status = 305
	// Status_NotAppearingInThisProtocol is a representation of the C type SOUP_STATUS_NOT_APPEARING_IN_THIS_PROTOCOL.
	Status_NotAppearingInThisProtocol Status = 306
	// Status_TemporaryRedirect is a representation of the C type SOUP_STATUS_TEMPORARY_REDIRECT.
	Status_TemporaryRedirect Status = 307
	// Status_BadRequest is a representation of the C type SOUP_STATUS_BAD_REQUEST.
	Status_BadRequest Status = 400
	// Status_Unauthorized is a representation of the C type SOUP_STATUS_UNAUTHORIZED.
	Status_Unauthorized Status = 401
	// Status_PaymentRequired is a representation of the C type SOUP_STATUS_PAYMENT_REQUIRED.
	Status_PaymentRequired Status = 402
	// Status_Forbidden is a representation of the C type SOUP_STATUS_FORBIDDEN.
	Status_Forbidden Status = 403
	// Status_NotFound is a representation of the C type SOUP_STATUS_NOT_FOUND.
	Status_NotFound Status = 404
	// Status_MethodNotAllowed is a representation of the C type SOUP_STATUS_METHOD_NOT_ALLOWED.
	Status_MethodNotAllowed Status = 405
	// Status_NotAcceptable is a representation of the C type SOUP_STATUS_NOT_ACCEPTABLE.
	Status_NotAcceptable Status = 406
	// Status_ProxyAuthenticationRequired is a representation of the C type SOUP_STATUS_PROXY_AUTHENTICATION_REQUIRED.
	Status_ProxyAuthenticationRequired Status = 407
	// Status_ProxyUnauthorized is a representation of the C type SOUP_STATUS_PROXY_UNAUTHORIZED.
	Status_ProxyUnauthorized Status = 407
	// Status_RequestTimeout is a representation of the C type SOUP_STATUS_REQUEST_TIMEOUT.
	Status_RequestTimeout Status = 408
	// Status_Conflict is a representation of the C type SOUP_STATUS_CONFLICT.
	Status_Conflict Status = 409
	// Status_Gone is a representation of the C type SOUP_STATUS_GONE.
	Status_Gone Status = 410
	// Status_LengthRequired is a representation of the C type SOUP_STATUS_LENGTH_REQUIRED.
	Status_LengthRequired Status = 411
	// Status_PreconditionFailed is a representation of the C type SOUP_STATUS_PRECONDITION_FAILED.
	Status_PreconditionFailed Status = 412
	// Status_RequestEntityTooLarge is a representation of the C type SOUP_STATUS_REQUEST_ENTITY_TOO_LARGE.
	Status_RequestEntityTooLarge Status = 413
	// Status_RequestUriTooLong is a representation of the C type SOUP_STATUS_REQUEST_URI_TOO_LONG.
	Status_RequestUriTooLong Status = 414
	// Status_UnsupportedMediaType is a representation of the C type SOUP_STATUS_UNSUPPORTED_MEDIA_TYPE.
	Status_UnsupportedMediaType Status = 415
	// Status_RequestedRangeNotSatisfiable is a representation of the C type SOUP_STATUS_REQUESTED_RANGE_NOT_SATISFIABLE.
	Status_RequestedRangeNotSatisfiable Status = 416
	// Status_InvalidRange is a representation of the C type SOUP_STATUS_INVALID_RANGE.
	Status_InvalidRange Status = 416
	// Status_ExpectationFailed is a representation of the C type SOUP_STATUS_EXPECTATION_FAILED.
	Status_ExpectationFailed Status = 417
	// Status_UnprocessableEntity is a representation of the C type SOUP_STATUS_UNPROCESSABLE_ENTITY.
	Status_UnprocessableEntity Status = 422
	// Status_Locked is a representation of the C type SOUP_STATUS_LOCKED.
	Status_Locked Status = 423
	// Status_FailedDependency is a representation of the C type SOUP_STATUS_FAILED_DEPENDENCY.
	Status_FailedDependency Status = 424
	// Status_InternalServerError is a representation of the C type SOUP_STATUS_INTERNAL_SERVER_ERROR.
	Status_InternalServerError Status = 500
	// Status_NotImplemented is a representation of the C type SOUP_STATUS_NOT_IMPLEMENTED.
	Status_NotImplemented Status = 501
	// Status_BadGateway is a representation of the C type SOUP_STATUS_BAD_GATEWAY.
	Status_BadGateway Status = 502
	// Status_ServiceUnavailable is a representation of the C type SOUP_STATUS_SERVICE_UNAVAILABLE.
	Status_ServiceUnavailable Status = 503
	// Status_GatewayTimeout is a representation of the C type SOUP_STATUS_GATEWAY_TIMEOUT.
	Status_GatewayTimeout Status = 504
	// Status_HttpVersionNotSupported is a representation of the C type SOUP_STATUS_HTTP_VERSION_NOT_SUPPORTED.
	Status_HttpVersionNotSupported Status = 505
	// Status_InsufficientStorage is a representation of the C type SOUP_STATUS_INSUFFICIENT_STORAGE.
	Status_InsufficientStorage Status = 507
	// Status_NotExtended is a representation of the C type SOUP_STATUS_NOT_EXTENDED.
	Status_NotExtended Status = 510
)

// TLDError is a representation of the C type SoupTLDError.
//
// since 2.40
type TLDError int

const (
	// TLDError_InvalidHostname is a representation of the C type SOUP_TLD_ERROR_INVALID_HOSTNAME.
	TLDError_InvalidHostname TLDError = 0
	// TLDError_IsIpAddress is a representation of the C type SOUP_TLD_ERROR_IS_IP_ADDRESS.
	TLDError_IsIpAddress TLDError = 1
	// TLDError_NotEnoughDomains is a representation of the C type SOUP_TLD_ERROR_NOT_ENOUGH_DOMAINS.
	TLDError_NotEnoughDomains TLDError = 2
	// TLDError_NoBaseDomain is a representation of the C type SOUP_TLD_ERROR_NO_BASE_DOMAIN.
	TLDError_NoBaseDomain TLDError = 3
	// TLDError_NoPslData is a representation of the C type SOUP_TLD_ERROR_NO_PSL_DATA.
	TLDError_NoPslData TLDError = 4
)

// WebsocketCloseCode is a representation of the C type SoupWebsocketCloseCode.
//
// since 2.50
type WebsocketCloseCode int

const (
	// WebsocketCloseCode_Normal is a representation of the C type SOUP_WEBSOCKET_CLOSE_NORMAL.
	WebsocketCloseCode_Normal WebsocketCloseCode = 1000
	// WebsocketCloseCode_GoingAway is a representation of the C type SOUP_WEBSOCKET_CLOSE_GOING_AWAY.
	WebsocketCloseCode_GoingAway WebsocketCloseCode = 1001
	// WebsocketCloseCode_ProtocolError is a representation of the C type SOUP_WEBSOCKET_CLOSE_PROTOCOL_ERROR.
	WebsocketCloseCode_ProtocolError WebsocketCloseCode = 1002
	// WebsocketCloseCode_UnsupportedData is a representation of the C type SOUP_WEBSOCKET_CLOSE_UNSUPPORTED_DATA.
	WebsocketCloseCode_UnsupportedData WebsocketCloseCode = 1003
	// WebsocketCloseCode_NoStatus is a representation of the C type SOUP_WEBSOCKET_CLOSE_NO_STATUS.
	WebsocketCloseCode_NoStatus WebsocketCloseCode = 1005
	// WebsocketCloseCode_Abnormal is a representation of the C type SOUP_WEBSOCKET_CLOSE_ABNORMAL.
	WebsocketCloseCode_Abnormal WebsocketCloseCode = 1006
	// WebsocketCloseCode_BadData is a representation of the C type SOUP_WEBSOCKET_CLOSE_BAD_DATA.
	WebsocketCloseCode_BadData WebsocketCloseCode = 1007
	// WebsocketCloseCode_PolicyViolation is a representation of the C type SOUP_WEBSOCKET_CLOSE_POLICY_VIOLATION.
	WebsocketCloseCode_PolicyViolation WebsocketCloseCode = 1008
	// WebsocketCloseCode_TooBig is a representation of the C type SOUP_WEBSOCKET_CLOSE_TOO_BIG.
	WebsocketCloseCode_TooBig WebsocketCloseCode = 1009
	// WebsocketCloseCode_NoExtension is a representation of the C type SOUP_WEBSOCKET_CLOSE_NO_EXTENSION.
	WebsocketCloseCode_NoExtension WebsocketCloseCode = 1010
	// WebsocketCloseCode_ServerError is a representation of the C type SOUP_WEBSOCKET_CLOSE_SERVER_ERROR.
	WebsocketCloseCode_ServerError WebsocketCloseCode = 1011
	// WebsocketCloseCode_TlsHandshake is a representation of the C type SOUP_WEBSOCKET_CLOSE_TLS_HANDSHAKE.
	WebsocketCloseCode_TlsHandshake WebsocketCloseCode = 1015
)

// WebsocketConnectionType is a representation of the C type SoupWebsocketConnectionType.
//
// since 2.50
type WebsocketConnectionType int

const (
	// WebsocketConnectionType_Unknown is a representation of the C type SOUP_WEBSOCKET_CONNECTION_UNKNOWN.
	WebsocketConnectionType_Unknown WebsocketConnectionType = 0
	// WebsocketConnectionType_Client is a representation of the C type SOUP_WEBSOCKET_CONNECTION_CLIENT.
	WebsocketConnectionType_Client WebsocketConnectionType = 1
	// WebsocketConnectionType_Server is a representation of the C type SOUP_WEBSOCKET_CONNECTION_SERVER.
	WebsocketConnectionType_Server WebsocketConnectionType = 2
)

// WebsocketDataType is a representation of the C type SoupWebsocketDataType.
//
// since 2.50
type WebsocketDataType int

const (
	// WebsocketDataType_Text is a representation of the C type SOUP_WEBSOCKET_DATA_TEXT.
	WebsocketDataType_Text WebsocketDataType = 1
	// WebsocketDataType_Binary is a representation of the C type SOUP_WEBSOCKET_DATA_BINARY.
	WebsocketDataType_Binary WebsocketDataType = 2
)

// WebsocketError is a representation of the C type SoupWebsocketError.
//
// since 2.50
type WebsocketError int

const (
	// WebsocketError_Failed is a representation of the C type SOUP_WEBSOCKET_ERROR_FAILED.
	WebsocketError_Failed WebsocketError = 0
	// WebsocketError_NotWebsocket is a representation of the C type SOUP_WEBSOCKET_ERROR_NOT_WEBSOCKET.
	WebsocketError_NotWebsocket WebsocketError = 1
	// WebsocketError_BadHandshake is a representation of the C type SOUP_WEBSOCKET_ERROR_BAD_HANDSHAKE.
	WebsocketError_BadHandshake WebsocketError = 2
	// WebsocketError_BadOrigin is a representation of the C type SOUP_WEBSOCKET_ERROR_BAD_ORIGIN.
	WebsocketError_BadOrigin WebsocketError = 3
)

// WebsocketState is a representation of the C type SoupWebsocketState.
//
// since 2.50
type WebsocketState int

const (
	// WebsocketState_Open is a representation of the C type SOUP_WEBSOCKET_STATE_OPEN.
	WebsocketState_Open WebsocketState = 1
	// WebsocketState_Closing is a representation of the C type SOUP_WEBSOCKET_STATE_CLOSING.
	WebsocketState_Closing WebsocketState = 2
	// WebsocketState_Closed is a representation of the C type SOUP_WEBSOCKET_STATE_CLOSED.
	WebsocketState_Closed WebsocketState = 3
)

// XMLRPCError is a representation of the C type SoupXMLRPCError.
//
type XMLRPCError int

const (
	// XMLRPCError_Arguments is a representation of the C type SOUP_XMLRPC_ERROR_ARGUMENTS.
	XMLRPCError_Arguments XMLRPCError = 0
	// XMLRPCError_Retval is a representation of the C type SOUP_XMLRPC_ERROR_RETVAL.
	XMLRPCError_Retval XMLRPCError = 1
)

// XMLRPCFault is a representation of the C type SoupXMLRPCFault.
//
type XMLRPCFault int

const (
	// XMLRPCFault_ParseErrorNotWellFormed is a representation of the C type SOUP_XMLRPC_FAULT_PARSE_ERROR_NOT_WELL_FORMED.
	XMLRPCFault_ParseErrorNotWellFormed XMLRPCFault = -32700
	// XMLRPCFault_ParseErrorUnsupportedEncoding is a representation of the C type SOUP_XMLRPC_FAULT_PARSE_ERROR_UNSUPPORTED_ENCODING.
	XMLRPCFault_ParseErrorUnsupportedEncoding XMLRPCFault = -32701
	// XMLRPCFault_ParseErrorInvalidCharacterForEncoding is a representation of the C type SOUP_XMLRPC_FAULT_PARSE_ERROR_INVALID_CHARACTER_FOR_ENCODING.
	XMLRPCFault_ParseErrorInvalidCharacterForEncoding XMLRPCFault = -32702
	// XMLRPCFault_ServerErrorInvalidXmlRpc is a representation of the C type SOUP_XMLRPC_FAULT_SERVER_ERROR_INVALID_XML_RPC.
	XMLRPCFault_ServerErrorInvalidXmlRpc XMLRPCFault = -32600
	// XMLRPCFault_ServerErrorRequestedMethodNotFound is a representation of the C type SOUP_XMLRPC_FAULT_SERVER_ERROR_REQUESTED_METHOD_NOT_FOUND.
	XMLRPCFault_ServerErrorRequestedMethodNotFound XMLRPCFault = -32601
	// XMLRPCFault_ServerErrorInvalidMethodParameters is a representation of the C type SOUP_XMLRPC_FAULT_SERVER_ERROR_INVALID_METHOD_PARAMETERS.
	XMLRPCFault_ServerErrorInvalidMethodParameters XMLRPCFault = -32602
	// XMLRPCFault_ServerErrorInternalXmlRpcError is a representation of the C type SOUP_XMLRPC_FAULT_SERVER_ERROR_INTERNAL_XML_RPC_ERROR.
	XMLRPCFault_ServerErrorInternalXmlRpcError XMLRPCFault = -32603
	// XMLRPCFault_ApplicationError is a representation of the C type SOUP_XMLRPC_FAULT_APPLICATION_ERROR.
	XMLRPCFault_ApplicationError XMLRPCFault = -32500
	// XMLRPCFault_SystemError is a representation of the C type SOUP_XMLRPC_FAULT_SYSTEM_ERROR.
	XMLRPCFault_SystemError XMLRPCFault = -32400
	// XMLRPCFault_TransportError is a representation of the C type SOUP_XMLRPC_FAULT_TRANSPORT_ERROR.
	XMLRPCFault_TransportError XMLRPCFault = -32300
)
