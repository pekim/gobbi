// Code generated - DO NOT EDIT.

package soup

// Addressfamily is a representation of the C type SoupAddressFamily.
type Addressfamily int

const (
	// Addressfamily_Invalid is a representation of the C type SOUP_ADDRESS_FAMILY_INVALID.
	Addressfamily_Invalid Addressfamily = -1
	// Addressfamily_Ipv4 is a representation of the C type SOUP_ADDRESS_FAMILY_IPV4.
	Addressfamily_Ipv4 Addressfamily = 2
	// Addressfamily_Ipv6 is a representation of the C type SOUP_ADDRESS_FAMILY_IPV6.
	Addressfamily_Ipv6 Addressfamily = 10
)

// Cacheresponse is a representation of the C type SoupCacheResponse.
type Cacheresponse int

const (
	// Cacheresponse_Fresh is a representation of the C type SOUP_CACHE_RESPONSE_FRESH.
	Cacheresponse_Fresh Cacheresponse = 0
	// Cacheresponse_NeedsValidation is a representation of the C type SOUP_CACHE_RESPONSE_NEEDS_VALIDATION.
	Cacheresponse_NeedsValidation Cacheresponse = 1
	// Cacheresponse_Stale is a representation of the C type SOUP_CACHE_RESPONSE_STALE.
	Cacheresponse_Stale Cacheresponse = 2
)

// Cachetype is a representation of the C type SoupCacheType.
//
// since 2.34
type Cachetype int

const (
	// Cachetype_SingleUser is a representation of the C type SOUP_CACHE_SINGLE_USER.
	Cachetype_SingleUser Cachetype = 0
	// Cachetype_Shared is a representation of the C type SOUP_CACHE_SHARED.
	Cachetype_Shared Cachetype = 1
)

// Connectionstate is a representation of the C type SoupConnectionState.
type Connectionstate int

const (
	// Connectionstate_New is a representation of the C type SOUP_CONNECTION_NEW.
	Connectionstate_New Connectionstate = 0
	// Connectionstate_Connecting is a representation of the C type SOUP_CONNECTION_CONNECTING.
	Connectionstate_Connecting Connectionstate = 1
	// Connectionstate_Idle is a representation of the C type SOUP_CONNECTION_IDLE.
	Connectionstate_Idle Connectionstate = 2
	// Connectionstate_InUse is a representation of the C type SOUP_CONNECTION_IN_USE.
	Connectionstate_InUse Connectionstate = 3
	// Connectionstate_RemoteDisconnected is a representation of the C type SOUP_CONNECTION_REMOTE_DISCONNECTED.
	Connectionstate_RemoteDisconnected Connectionstate = 4
	// Connectionstate_Disconnected is a representation of the C type SOUP_CONNECTION_DISCONNECTED.
	Connectionstate_Disconnected Connectionstate = 5
)

// Cookiejaracceptpolicy is a representation of the C type SoupCookieJarAcceptPolicy.
//
// since 2.30
type Cookiejaracceptpolicy int

const (
	// Cookiejaracceptpolicy_Always is a representation of the C type SOUP_COOKIE_JAR_ACCEPT_ALWAYS.
	Cookiejaracceptpolicy_Always Cookiejaracceptpolicy = 0
	// Cookiejaracceptpolicy_Never is a representation of the C type SOUP_COOKIE_JAR_ACCEPT_NEVER.
	Cookiejaracceptpolicy_Never Cookiejaracceptpolicy = 1
	// Cookiejaracceptpolicy_NoThirdParty is a representation of the C type SOUP_COOKIE_JAR_ACCEPT_NO_THIRD_PARTY.
	Cookiejaracceptpolicy_NoThirdParty Cookiejaracceptpolicy = 2
)

// Dateformat is a representation of the C type SoupDateFormat.
type Dateformat int

const (
	// Dateformat_Http is a representation of the C type SOUP_DATE_HTTP.
	Dateformat_Http Dateformat = 1
	// Dateformat_Cookie is a representation of the C type SOUP_DATE_COOKIE.
	Dateformat_Cookie Dateformat = 2
	// Dateformat_Rfc2822 is a representation of the C type SOUP_DATE_RFC2822.
	Dateformat_Rfc2822 Dateformat = 3
	// Dateformat_Iso8601Compact is a representation of the C type SOUP_DATE_ISO8601_COMPACT.
	Dateformat_Iso8601Compact Dateformat = 4
	// Dateformat_Iso8601Full is a representation of the C type SOUP_DATE_ISO8601_FULL.
	Dateformat_Iso8601Full Dateformat = 5
	// Dateformat_Iso8601 is a representation of the C type SOUP_DATE_ISO8601.
	Dateformat_Iso8601 Dateformat = 5
	// Dateformat_Iso8601Xmlrpc is a representation of the C type SOUP_DATE_ISO8601_XMLRPC.
	Dateformat_Iso8601Xmlrpc Dateformat = 6
)

// Encoding is a representation of the C type SoupEncoding.
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

// Httpversion is a representation of the C type SoupHTTPVersion.
type Httpversion int

const (
	// Httpversion_Http10 is a representation of the C type SOUP_HTTP_1_0.
	Httpversion_Http10 Httpversion = 0
	// Httpversion_Http11 is a representation of the C type SOUP_HTTP_1_1.
	Httpversion_Http11 Httpversion = 1
)

// Knownstatuscode is a representation of the C type SoupKnownStatusCode.
type Knownstatuscode int

const (
	// Knownstatuscode_None is a representation of the C type SOUP_KNOWN_STATUS_CODE_NONE.
	Knownstatuscode_None Knownstatuscode = 0
	// Knownstatuscode_Cancelled is a representation of the C type SOUP_KNOWN_STATUS_CODE_CANCELLED.
	Knownstatuscode_Cancelled Knownstatuscode = 1
	// Knownstatuscode_CantResolve is a representation of the C type SOUP_KNOWN_STATUS_CODE_CANT_RESOLVE.
	Knownstatuscode_CantResolve Knownstatuscode = 2
	// Knownstatuscode_CantResolveProxy is a representation of the C type SOUP_KNOWN_STATUS_CODE_CANT_RESOLVE_PROXY.
	Knownstatuscode_CantResolveProxy Knownstatuscode = 3
	// Knownstatuscode_CantConnect is a representation of the C type SOUP_KNOWN_STATUS_CODE_CANT_CONNECT.
	Knownstatuscode_CantConnect Knownstatuscode = 4
	// Knownstatuscode_CantConnectProxy is a representation of the C type SOUP_KNOWN_STATUS_CODE_CANT_CONNECT_PROXY.
	Knownstatuscode_CantConnectProxy Knownstatuscode = 5
	// Knownstatuscode_SslFailed is a representation of the C type SOUP_KNOWN_STATUS_CODE_SSL_FAILED.
	Knownstatuscode_SslFailed Knownstatuscode = 6
	// Knownstatuscode_IoError is a representation of the C type SOUP_KNOWN_STATUS_CODE_IO_ERROR.
	Knownstatuscode_IoError Knownstatuscode = 7
	// Knownstatuscode_Malformed is a representation of the C type SOUP_KNOWN_STATUS_CODE_MALFORMED.
	Knownstatuscode_Malformed Knownstatuscode = 8
	// Knownstatuscode_TryAgain is a representation of the C type SOUP_KNOWN_STATUS_CODE_TRY_AGAIN.
	Knownstatuscode_TryAgain Knownstatuscode = 9
	// Knownstatuscode_TooManyRedirects is a representation of the C type SOUP_KNOWN_STATUS_CODE_TOO_MANY_REDIRECTS.
	Knownstatuscode_TooManyRedirects Knownstatuscode = 10
	// Knownstatuscode_TlsFailed is a representation of the C type SOUP_KNOWN_STATUS_CODE_TLS_FAILED.
	Knownstatuscode_TlsFailed Knownstatuscode = 11
	// Knownstatuscode_Continue is a representation of the C type SOUP_KNOWN_STATUS_CODE_CONTINUE.
	Knownstatuscode_Continue Knownstatuscode = 100
	// Knownstatuscode_SwitchingProtocols is a representation of the C type SOUP_KNOWN_STATUS_CODE_SWITCHING_PROTOCOLS.
	Knownstatuscode_SwitchingProtocols Knownstatuscode = 101
	// Knownstatuscode_Processing is a representation of the C type SOUP_KNOWN_STATUS_CODE_PROCESSING.
	Knownstatuscode_Processing Knownstatuscode = 102
	// Knownstatuscode_Ok is a representation of the C type SOUP_KNOWN_STATUS_CODE_OK.
	Knownstatuscode_Ok Knownstatuscode = 200
	// Knownstatuscode_Created is a representation of the C type SOUP_KNOWN_STATUS_CODE_CREATED.
	Knownstatuscode_Created Knownstatuscode = 201
	// Knownstatuscode_Accepted is a representation of the C type SOUP_KNOWN_STATUS_CODE_ACCEPTED.
	Knownstatuscode_Accepted Knownstatuscode = 202
	// Knownstatuscode_NonAuthoritative is a representation of the C type SOUP_KNOWN_STATUS_CODE_NON_AUTHORITATIVE.
	Knownstatuscode_NonAuthoritative Knownstatuscode = 203
	// Knownstatuscode_NoContent is a representation of the C type SOUP_KNOWN_STATUS_CODE_NO_CONTENT.
	Knownstatuscode_NoContent Knownstatuscode = 204
	// Knownstatuscode_ResetContent is a representation of the C type SOUP_KNOWN_STATUS_CODE_RESET_CONTENT.
	Knownstatuscode_ResetContent Knownstatuscode = 205
	// Knownstatuscode_PartialContent is a representation of the C type SOUP_KNOWN_STATUS_CODE_PARTIAL_CONTENT.
	Knownstatuscode_PartialContent Knownstatuscode = 206
	// Knownstatuscode_MultiStatus is a representation of the C type SOUP_KNOWN_STATUS_CODE_MULTI_STATUS.
	Knownstatuscode_MultiStatus Knownstatuscode = 207
	// Knownstatuscode_MultipleChoices is a representation of the C type SOUP_KNOWN_STATUS_CODE_MULTIPLE_CHOICES.
	Knownstatuscode_MultipleChoices Knownstatuscode = 300
	// Knownstatuscode_MovedPermanently is a representation of the C type SOUP_KNOWN_STATUS_CODE_MOVED_PERMANENTLY.
	Knownstatuscode_MovedPermanently Knownstatuscode = 301
	// Knownstatuscode_Found is a representation of the C type SOUP_KNOWN_STATUS_CODE_FOUND.
	Knownstatuscode_Found Knownstatuscode = 302
	// Knownstatuscode_MovedTemporarily is a representation of the C type SOUP_KNOWN_STATUS_CODE_MOVED_TEMPORARILY.
	Knownstatuscode_MovedTemporarily Knownstatuscode = 302
	// Knownstatuscode_SeeOther is a representation of the C type SOUP_KNOWN_STATUS_CODE_SEE_OTHER.
	Knownstatuscode_SeeOther Knownstatuscode = 303
	// Knownstatuscode_NotModified is a representation of the C type SOUP_KNOWN_STATUS_CODE_NOT_MODIFIED.
	Knownstatuscode_NotModified Knownstatuscode = 304
	// Knownstatuscode_UseProxy is a representation of the C type SOUP_KNOWN_STATUS_CODE_USE_PROXY.
	Knownstatuscode_UseProxy Knownstatuscode = 305
	// Knownstatuscode_NotAppearingInThisProtocol is a representation of the C type SOUP_KNOWN_STATUS_CODE_NOT_APPEARING_IN_THIS_PROTOCOL.
	Knownstatuscode_NotAppearingInThisProtocol Knownstatuscode = 306
	// Knownstatuscode_TemporaryRedirect is a representation of the C type SOUP_KNOWN_STATUS_CODE_TEMPORARY_REDIRECT.
	Knownstatuscode_TemporaryRedirect Knownstatuscode = 307
	// Knownstatuscode_BadRequest is a representation of the C type SOUP_KNOWN_STATUS_CODE_BAD_REQUEST.
	Knownstatuscode_BadRequest Knownstatuscode = 400
	// Knownstatuscode_Unauthorized is a representation of the C type SOUP_KNOWN_STATUS_CODE_UNAUTHORIZED.
	Knownstatuscode_Unauthorized Knownstatuscode = 401
	// Knownstatuscode_PaymentRequired is a representation of the C type SOUP_KNOWN_STATUS_CODE_PAYMENT_REQUIRED.
	Knownstatuscode_PaymentRequired Knownstatuscode = 402
	// Knownstatuscode_Forbidden is a representation of the C type SOUP_KNOWN_STATUS_CODE_FORBIDDEN.
	Knownstatuscode_Forbidden Knownstatuscode = 403
	// Knownstatuscode_NotFound is a representation of the C type SOUP_KNOWN_STATUS_CODE_NOT_FOUND.
	Knownstatuscode_NotFound Knownstatuscode = 404
	// Knownstatuscode_MethodNotAllowed is a representation of the C type SOUP_KNOWN_STATUS_CODE_METHOD_NOT_ALLOWED.
	Knownstatuscode_MethodNotAllowed Knownstatuscode = 405
	// Knownstatuscode_NotAcceptable is a representation of the C type SOUP_KNOWN_STATUS_CODE_NOT_ACCEPTABLE.
	Knownstatuscode_NotAcceptable Knownstatuscode = 406
	// Knownstatuscode_ProxyAuthenticationRequired is a representation of the C type SOUP_KNOWN_STATUS_CODE_PROXY_AUTHENTICATION_REQUIRED.
	Knownstatuscode_ProxyAuthenticationRequired Knownstatuscode = 407
	// Knownstatuscode_ProxyUnauthorized is a representation of the C type SOUP_KNOWN_STATUS_CODE_PROXY_UNAUTHORIZED.
	Knownstatuscode_ProxyUnauthorized Knownstatuscode = 407
	// Knownstatuscode_RequestTimeout is a representation of the C type SOUP_KNOWN_STATUS_CODE_REQUEST_TIMEOUT.
	Knownstatuscode_RequestTimeout Knownstatuscode = 408
	// Knownstatuscode_Conflict is a representation of the C type SOUP_KNOWN_STATUS_CODE_CONFLICT.
	Knownstatuscode_Conflict Knownstatuscode = 409
	// Knownstatuscode_Gone is a representation of the C type SOUP_KNOWN_STATUS_CODE_GONE.
	Knownstatuscode_Gone Knownstatuscode = 410
	// Knownstatuscode_LengthRequired is a representation of the C type SOUP_KNOWN_STATUS_CODE_LENGTH_REQUIRED.
	Knownstatuscode_LengthRequired Knownstatuscode = 411
	// Knownstatuscode_PreconditionFailed is a representation of the C type SOUP_KNOWN_STATUS_CODE_PRECONDITION_FAILED.
	Knownstatuscode_PreconditionFailed Knownstatuscode = 412
	// Knownstatuscode_RequestEntityTooLarge is a representation of the C type SOUP_KNOWN_STATUS_CODE_REQUEST_ENTITY_TOO_LARGE.
	Knownstatuscode_RequestEntityTooLarge Knownstatuscode = 413
	// Knownstatuscode_RequestUriTooLong is a representation of the C type SOUP_KNOWN_STATUS_CODE_REQUEST_URI_TOO_LONG.
	Knownstatuscode_RequestUriTooLong Knownstatuscode = 414
	// Knownstatuscode_UnsupportedMediaType is a representation of the C type SOUP_KNOWN_STATUS_CODE_UNSUPPORTED_MEDIA_TYPE.
	Knownstatuscode_UnsupportedMediaType Knownstatuscode = 415
	// Knownstatuscode_RequestedRangeNotSatisfiable is a representation of the C type SOUP_KNOWN_STATUS_CODE_REQUESTED_RANGE_NOT_SATISFIABLE.
	Knownstatuscode_RequestedRangeNotSatisfiable Knownstatuscode = 416
	// Knownstatuscode_InvalidRange is a representation of the C type SOUP_KNOWN_STATUS_CODE_INVALID_RANGE.
	Knownstatuscode_InvalidRange Knownstatuscode = 416
	// Knownstatuscode_ExpectationFailed is a representation of the C type SOUP_KNOWN_STATUS_CODE_EXPECTATION_FAILED.
	Knownstatuscode_ExpectationFailed Knownstatuscode = 417
	// Knownstatuscode_UnprocessableEntity is a representation of the C type SOUP_KNOWN_STATUS_CODE_UNPROCESSABLE_ENTITY.
	Knownstatuscode_UnprocessableEntity Knownstatuscode = 422
	// Knownstatuscode_Locked is a representation of the C type SOUP_KNOWN_STATUS_CODE_LOCKED.
	Knownstatuscode_Locked Knownstatuscode = 423
	// Knownstatuscode_FailedDependency is a representation of the C type SOUP_KNOWN_STATUS_CODE_FAILED_DEPENDENCY.
	Knownstatuscode_FailedDependency Knownstatuscode = 424
	// Knownstatuscode_InternalServerError is a representation of the C type SOUP_KNOWN_STATUS_CODE_INTERNAL_SERVER_ERROR.
	Knownstatuscode_InternalServerError Knownstatuscode = 500
	// Knownstatuscode_NotImplemented is a representation of the C type SOUP_KNOWN_STATUS_CODE_NOT_IMPLEMENTED.
	Knownstatuscode_NotImplemented Knownstatuscode = 501
	// Knownstatuscode_BadGateway is a representation of the C type SOUP_KNOWN_STATUS_CODE_BAD_GATEWAY.
	Knownstatuscode_BadGateway Knownstatuscode = 502
	// Knownstatuscode_ServiceUnavailable is a representation of the C type SOUP_KNOWN_STATUS_CODE_SERVICE_UNAVAILABLE.
	Knownstatuscode_ServiceUnavailable Knownstatuscode = 503
	// Knownstatuscode_GatewayTimeout is a representation of the C type SOUP_KNOWN_STATUS_CODE_GATEWAY_TIMEOUT.
	Knownstatuscode_GatewayTimeout Knownstatuscode = 504
	// Knownstatuscode_HttpVersionNotSupported is a representation of the C type SOUP_KNOWN_STATUS_CODE_HTTP_VERSION_NOT_SUPPORTED.
	Knownstatuscode_HttpVersionNotSupported Knownstatuscode = 505
	// Knownstatuscode_InsufficientStorage is a representation of the C type SOUP_KNOWN_STATUS_CODE_INSUFFICIENT_STORAGE.
	Knownstatuscode_InsufficientStorage Knownstatuscode = 507
	// Knownstatuscode_NotExtended is a representation of the C type SOUP_KNOWN_STATUS_CODE_NOT_EXTENDED.
	Knownstatuscode_NotExtended Knownstatuscode = 510
)

// Loggerloglevel is a representation of the C type SoupLoggerLogLevel.
type Loggerloglevel int

const (
	// Loggerloglevel_None is a representation of the C type SOUP_LOGGER_LOG_NONE.
	Loggerloglevel_None Loggerloglevel = 0
	// Loggerloglevel_Minimal is a representation of the C type SOUP_LOGGER_LOG_MINIMAL.
	Loggerloglevel_Minimal Loggerloglevel = 1
	// Loggerloglevel_Headers is a representation of the C type SOUP_LOGGER_LOG_HEADERS.
	Loggerloglevel_Headers Loggerloglevel = 2
	// Loggerloglevel_Body is a representation of the C type SOUP_LOGGER_LOG_BODY.
	Loggerloglevel_Body Loggerloglevel = 3
)

// Memoryuse is a representation of the C type SoupMemoryUse.
type Memoryuse int

const (
	// Memoryuse_Static is a representation of the C type SOUP_MEMORY_STATIC.
	Memoryuse_Static Memoryuse = 0
	// Memoryuse_Take is a representation of the C type SOUP_MEMORY_TAKE.
	Memoryuse_Take Memoryuse = 1
	// Memoryuse_Copy is a representation of the C type SOUP_MEMORY_COPY.
	Memoryuse_Copy Memoryuse = 2
	// Memoryuse_Temporary is a representation of the C type SOUP_MEMORY_TEMPORARY.
	Memoryuse_Temporary Memoryuse = 3
)

// Messageheaderstype is a representation of the C type SoupMessageHeadersType.
type Messageheaderstype int

const (
	// Messageheaderstype_Request is a representation of the C type SOUP_MESSAGE_HEADERS_REQUEST.
	Messageheaderstype_Request Messageheaderstype = 0
	// Messageheaderstype_Response is a representation of the C type SOUP_MESSAGE_HEADERS_RESPONSE.
	Messageheaderstype_Response Messageheaderstype = 1
	// Messageheaderstype_Multipart is a representation of the C type SOUP_MESSAGE_HEADERS_MULTIPART.
	Messageheaderstype_Multipart Messageheaderstype = 2
)

// Messagepriority is a representation of the C type SoupMessagePriority.
type Messagepriority int

const (
	// Messagepriority_VeryLow is a representation of the C type SOUP_MESSAGE_PRIORITY_VERY_LOW.
	Messagepriority_VeryLow Messagepriority = 0
	// Messagepriority_Low is a representation of the C type SOUP_MESSAGE_PRIORITY_LOW.
	Messagepriority_Low Messagepriority = 1
	// Messagepriority_Normal is a representation of the C type SOUP_MESSAGE_PRIORITY_NORMAL.
	Messagepriority_Normal Messagepriority = 2
	// Messagepriority_High is a representation of the C type SOUP_MESSAGE_PRIORITY_HIGH.
	Messagepriority_High Messagepriority = 3
	// Messagepriority_VeryHigh is a representation of the C type SOUP_MESSAGE_PRIORITY_VERY_HIGH.
	Messagepriority_VeryHigh Messagepriority = 4
)

// Requesterror is a representation of the C type SoupRequestError.
//
// since 2.42
type Requesterror int

const (
	// Requesterror_BadUri is a representation of the C type SOUP_REQUEST_ERROR_BAD_URI.
	Requesterror_BadUri Requesterror = 0
	// Requesterror_UnsupportedUriScheme is a representation of the C type SOUP_REQUEST_ERROR_UNSUPPORTED_URI_SCHEME.
	Requesterror_UnsupportedUriScheme Requesterror = 1
	// Requesterror_Parsing is a representation of the C type SOUP_REQUEST_ERROR_PARSING.
	Requesterror_Parsing Requesterror = 2
	// Requesterror_Encoding is a representation of the C type SOUP_REQUEST_ERROR_ENCODING.
	Requesterror_Encoding Requesterror = 3
)

// Requestererror is a representation of the C type SoupRequesterError.
type Requestererror int

const (
	// Requestererror_BadUri is a representation of the C type SOUP_REQUESTER_ERROR_BAD_URI.
	Requestererror_BadUri Requestererror = 0
	// Requestererror_UnsupportedUriScheme is a representation of the C type SOUP_REQUESTER_ERROR_UNSUPPORTED_URI_SCHEME.
	Requestererror_UnsupportedUriScheme Requestererror = 1
)

// Socketiostatus is a representation of the C type SoupSocketIOStatus.
type Socketiostatus int

const (
	// Socketiostatus_Ok is a representation of the C type SOUP_SOCKET_OK.
	Socketiostatus_Ok Socketiostatus = 0
	// Socketiostatus_WouldBlock is a representation of the C type SOUP_SOCKET_WOULD_BLOCK.
	Socketiostatus_WouldBlock Socketiostatus = 1
	// Socketiostatus_Eof is a representation of the C type SOUP_SOCKET_EOF.
	Socketiostatus_Eof Socketiostatus = 2
	// Socketiostatus_Error is a representation of the C type SOUP_SOCKET_ERROR.
	Socketiostatus_Error Socketiostatus = 3
)

// Status is a representation of the C type SoupStatus.
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

// Tlderror is a representation of the C type SoupTLDError.
//
// since 2.40
type Tlderror int

const (
	// Tlderror_InvalidHostname is a representation of the C type SOUP_TLD_ERROR_INVALID_HOSTNAME.
	Tlderror_InvalidHostname Tlderror = 0
	// Tlderror_IsIpAddress is a representation of the C type SOUP_TLD_ERROR_IS_IP_ADDRESS.
	Tlderror_IsIpAddress Tlderror = 1
	// Tlderror_NotEnoughDomains is a representation of the C type SOUP_TLD_ERROR_NOT_ENOUGH_DOMAINS.
	Tlderror_NotEnoughDomains Tlderror = 2
	// Tlderror_NoBaseDomain is a representation of the C type SOUP_TLD_ERROR_NO_BASE_DOMAIN.
	Tlderror_NoBaseDomain Tlderror = 3
)

// Websocketclosecode is a representation of the C type SoupWebsocketCloseCode.
//
// since 2.50
type Websocketclosecode int

const (
	// Websocketclosecode_Normal is a representation of the C type SOUP_WEBSOCKET_CLOSE_NORMAL.
	Websocketclosecode_Normal Websocketclosecode = 1000
	// Websocketclosecode_GoingAway is a representation of the C type SOUP_WEBSOCKET_CLOSE_GOING_AWAY.
	Websocketclosecode_GoingAway Websocketclosecode = 1001
	// Websocketclosecode_ProtocolError is a representation of the C type SOUP_WEBSOCKET_CLOSE_PROTOCOL_ERROR.
	Websocketclosecode_ProtocolError Websocketclosecode = 1002
	// Websocketclosecode_UnsupportedData is a representation of the C type SOUP_WEBSOCKET_CLOSE_UNSUPPORTED_DATA.
	Websocketclosecode_UnsupportedData Websocketclosecode = 1003
	// Websocketclosecode_NoStatus is a representation of the C type SOUP_WEBSOCKET_CLOSE_NO_STATUS.
	Websocketclosecode_NoStatus Websocketclosecode = 1005
	// Websocketclosecode_Abnormal is a representation of the C type SOUP_WEBSOCKET_CLOSE_ABNORMAL.
	Websocketclosecode_Abnormal Websocketclosecode = 1006
	// Websocketclosecode_BadData is a representation of the C type SOUP_WEBSOCKET_CLOSE_BAD_DATA.
	Websocketclosecode_BadData Websocketclosecode = 1007
	// Websocketclosecode_PolicyViolation is a representation of the C type SOUP_WEBSOCKET_CLOSE_POLICY_VIOLATION.
	Websocketclosecode_PolicyViolation Websocketclosecode = 1008
	// Websocketclosecode_TooBig is a representation of the C type SOUP_WEBSOCKET_CLOSE_TOO_BIG.
	Websocketclosecode_TooBig Websocketclosecode = 1009
	// Websocketclosecode_NoExtension is a representation of the C type SOUP_WEBSOCKET_CLOSE_NO_EXTENSION.
	Websocketclosecode_NoExtension Websocketclosecode = 1010
	// Websocketclosecode_ServerError is a representation of the C type SOUP_WEBSOCKET_CLOSE_SERVER_ERROR.
	Websocketclosecode_ServerError Websocketclosecode = 1011
	// Websocketclosecode_TlsHandshake is a representation of the C type SOUP_WEBSOCKET_CLOSE_TLS_HANDSHAKE.
	Websocketclosecode_TlsHandshake Websocketclosecode = 1015
)

// Websocketconnectiontype is a representation of the C type SoupWebsocketConnectionType.
//
// since 2.50
type Websocketconnectiontype int

const (
	// Websocketconnectiontype_Unknown is a representation of the C type SOUP_WEBSOCKET_CONNECTION_UNKNOWN.
	Websocketconnectiontype_Unknown Websocketconnectiontype = 0
	// Websocketconnectiontype_Client is a representation of the C type SOUP_WEBSOCKET_CONNECTION_CLIENT.
	Websocketconnectiontype_Client Websocketconnectiontype = 1
	// Websocketconnectiontype_Server is a representation of the C type SOUP_WEBSOCKET_CONNECTION_SERVER.
	Websocketconnectiontype_Server Websocketconnectiontype = 2
)

// Websocketdatatype is a representation of the C type SoupWebsocketDataType.
//
// since 2.50
type Websocketdatatype int

const (
	// Websocketdatatype_Text is a representation of the C type SOUP_WEBSOCKET_DATA_TEXT.
	Websocketdatatype_Text Websocketdatatype = 1
	// Websocketdatatype_Binary is a representation of the C type SOUP_WEBSOCKET_DATA_BINARY.
	Websocketdatatype_Binary Websocketdatatype = 2
)

// Websocketerror is a representation of the C type SoupWebsocketError.
//
// since 2.50
type Websocketerror int

const (
	// Websocketerror_Failed is a representation of the C type SOUP_WEBSOCKET_ERROR_FAILED.
	Websocketerror_Failed Websocketerror = 0
	// Websocketerror_NotWebsocket is a representation of the C type SOUP_WEBSOCKET_ERROR_NOT_WEBSOCKET.
	Websocketerror_NotWebsocket Websocketerror = 1
	// Websocketerror_BadHandshake is a representation of the C type SOUP_WEBSOCKET_ERROR_BAD_HANDSHAKE.
	Websocketerror_BadHandshake Websocketerror = 2
	// Websocketerror_BadOrigin is a representation of the C type SOUP_WEBSOCKET_ERROR_BAD_ORIGIN.
	Websocketerror_BadOrigin Websocketerror = 3
)

// Websocketstate is a representation of the C type SoupWebsocketState.
//
// since 2.50
type Websocketstate int

const (
	// Websocketstate_Open is a representation of the C type SOUP_WEBSOCKET_STATE_OPEN.
	Websocketstate_Open Websocketstate = 1
	// Websocketstate_Closing is a representation of the C type SOUP_WEBSOCKET_STATE_CLOSING.
	Websocketstate_Closing Websocketstate = 2
	// Websocketstate_Closed is a representation of the C type SOUP_WEBSOCKET_STATE_CLOSED.
	Websocketstate_Closed Websocketstate = 3
)

// Xmlrpcerror is a representation of the C type SoupXMLRPCError.
type Xmlrpcerror int

const (
	// Xmlrpcerror_Arguments is a representation of the C type SOUP_XMLRPC_ERROR_ARGUMENTS.
	Xmlrpcerror_Arguments Xmlrpcerror = 0
	// Xmlrpcerror_Retval is a representation of the C type SOUP_XMLRPC_ERROR_RETVAL.
	Xmlrpcerror_Retval Xmlrpcerror = 1
)

// Xmlrpcfault is a representation of the C type SoupXMLRPCFault.
type Xmlrpcfault int

const (
	// Xmlrpcfault_ParseErrorNotWellFormed is a representation of the C type SOUP_XMLRPC_FAULT_PARSE_ERROR_NOT_WELL_FORMED.
	Xmlrpcfault_ParseErrorNotWellFormed Xmlrpcfault = -32700
	// Xmlrpcfault_ParseErrorUnsupportedEncoding is a representation of the C type SOUP_XMLRPC_FAULT_PARSE_ERROR_UNSUPPORTED_ENCODING.
	Xmlrpcfault_ParseErrorUnsupportedEncoding Xmlrpcfault = -32701
	// Xmlrpcfault_ParseErrorInvalidCharacterForEncoding is a representation of the C type SOUP_XMLRPC_FAULT_PARSE_ERROR_INVALID_CHARACTER_FOR_ENCODING.
	Xmlrpcfault_ParseErrorInvalidCharacterForEncoding Xmlrpcfault = -32702
	// Xmlrpcfault_ServerErrorInvalidXmlRpc is a representation of the C type SOUP_XMLRPC_FAULT_SERVER_ERROR_INVALID_XML_RPC.
	Xmlrpcfault_ServerErrorInvalidXmlRpc Xmlrpcfault = -32600
	// Xmlrpcfault_ServerErrorRequestedMethodNotFound is a representation of the C type SOUP_XMLRPC_FAULT_SERVER_ERROR_REQUESTED_METHOD_NOT_FOUND.
	Xmlrpcfault_ServerErrorRequestedMethodNotFound Xmlrpcfault = -32601
	// Xmlrpcfault_ServerErrorInvalidMethodParameters is a representation of the C type SOUP_XMLRPC_FAULT_SERVER_ERROR_INVALID_METHOD_PARAMETERS.
	Xmlrpcfault_ServerErrorInvalidMethodParameters Xmlrpcfault = -32602
	// Xmlrpcfault_ServerErrorInternalXmlRpcError is a representation of the C type SOUP_XMLRPC_FAULT_SERVER_ERROR_INTERNAL_XML_RPC_ERROR.
	Xmlrpcfault_ServerErrorInternalXmlRpcError Xmlrpcfault = -32603
	// Xmlrpcfault_ApplicationError is a representation of the C type SOUP_XMLRPC_FAULT_APPLICATION_ERROR.
	Xmlrpcfault_ApplicationError Xmlrpcfault = -32500
	// Xmlrpcfault_SystemError is a representation of the C type SOUP_XMLRPC_FAULT_SYSTEM_ERROR.
	Xmlrpcfault_SystemError Xmlrpcfault = -32400
	// Xmlrpcfault_TransportError is a representation of the C type SOUP_XMLRPC_FAULT_TRANSPORT_ERROR.
	Xmlrpcfault_TransportError Xmlrpcfault = -32300
)
