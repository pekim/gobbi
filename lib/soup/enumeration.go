// Code generated - DO NOT EDIT.

package soup

// Addressfamily is a representation of the C type AddressFamily.
type Addressfamily int

const (
	// invalid
	SoupAddressFamilyInvalid Addressfamily = -1
	// ipv4
	SoupAddressFamilyIpv4 Addressfamily = 2
	// ipv6
	SoupAddressFamilyIpv6 Addressfamily = 10
)

// Cacheresponse is a representation of the C type CacheResponse.
type Cacheresponse int

const (
	// fresh
	SoupCacheResponseFresh Cacheresponse = 0
	// needs_validation
	SoupCacheResponseNeedsValidation Cacheresponse = 1
	// stale
	SoupCacheResponseStale Cacheresponse = 2
)

// Cachetype is a representation of the C type CacheType.
//
// since 2.34
type Cachetype int

const (
	// single_user
	SoupCacheSingleUser Cachetype = 0
	// shared
	SoupCacheShared Cachetype = 1
)

// Connectionstate is a representation of the C type ConnectionState.
type Connectionstate int

const (
	// new
	SoupConnectionNew Connectionstate = 0
	// connecting
	SoupConnectionConnecting Connectionstate = 1
	// idle
	SoupConnectionIdle Connectionstate = 2
	// in_use
	SoupConnectionInUse Connectionstate = 3
	// remote_disconnected
	SoupConnectionRemoteDisconnected Connectionstate = 4
	// disconnected
	SoupConnectionDisconnected Connectionstate = 5
)

// Cookiejaracceptpolicy is a representation of the C type CookieJarAcceptPolicy.
//
// since 2.30
type Cookiejaracceptpolicy int

const (
	// always
	SoupCookieJarAcceptAlways Cookiejaracceptpolicy = 0
	// never
	SoupCookieJarAcceptNever Cookiejaracceptpolicy = 1
	// no_third_party
	SoupCookieJarAcceptNoThirdParty Cookiejaracceptpolicy = 2
)

// Dateformat is a representation of the C type DateFormat.
type Dateformat int

const (
	// http
	SoupDateHttp Dateformat = 1
	// cookie
	SoupDateCookie Dateformat = 2
	// rfc2822
	SoupDateRfc2822 Dateformat = 3
	// iso8601_compact
	SoupDateIso8601Compact Dateformat = 4
	// iso8601_full
	SoupDateIso8601Full Dateformat = 5
	// iso8601
	SoupDateIso8601 Dateformat = 5
	// iso8601_xmlrpc
	SoupDateIso8601Xmlrpc Dateformat = 6
)

// Encoding is a representation of the C type Encoding.
type Encoding int

const (
	// unrecognized
	SoupEncodingUnrecognized Encoding = 0
	// none
	SoupEncodingNone Encoding = 1
	// content_length
	SoupEncodingContentLength Encoding = 2
	// eof
	SoupEncodingEof Encoding = 3
	// chunked
	SoupEncodingChunked Encoding = 4
	// byteranges
	SoupEncodingByteranges Encoding = 5
)

// Httpversion is a representation of the C type HTTPVersion.
type Httpversion int

const (
	// http_1_0
	SoupHttp10 Httpversion = 0
	// http_1_1
	SoupHttp11 Httpversion = 1
)

// Knownstatuscode is a representation of the C type KnownStatusCode.
type Knownstatuscode int

const (
	// none
	SoupKnownStatusCodeNone Knownstatuscode = 0
	// cancelled
	SoupKnownStatusCodeCancelled Knownstatuscode = 1
	// cant_resolve
	SoupKnownStatusCodeCantResolve Knownstatuscode = 2
	// cant_resolve_proxy
	SoupKnownStatusCodeCantResolveProxy Knownstatuscode = 3
	// cant_connect
	SoupKnownStatusCodeCantConnect Knownstatuscode = 4
	// cant_connect_proxy
	SoupKnownStatusCodeCantConnectProxy Knownstatuscode = 5
	// ssl_failed
	SoupKnownStatusCodeSslFailed Knownstatuscode = 6
	// io_error
	SoupKnownStatusCodeIoError Knownstatuscode = 7
	// malformed
	SoupKnownStatusCodeMalformed Knownstatuscode = 8
	// try_again
	SoupKnownStatusCodeTryAgain Knownstatuscode = 9
	// too_many_redirects
	SoupKnownStatusCodeTooManyRedirects Knownstatuscode = 10
	// tls_failed
	SoupKnownStatusCodeTlsFailed Knownstatuscode = 11
	// continue
	SoupKnownStatusCodeContinue Knownstatuscode = 100
	// switching_protocols
	SoupKnownStatusCodeSwitchingProtocols Knownstatuscode = 101
	// processing
	SoupKnownStatusCodeProcessing Knownstatuscode = 102
	// ok
	SoupKnownStatusCodeOk Knownstatuscode = 200
	// created
	SoupKnownStatusCodeCreated Knownstatuscode = 201
	// accepted
	SoupKnownStatusCodeAccepted Knownstatuscode = 202
	// non_authoritative
	SoupKnownStatusCodeNonAuthoritative Knownstatuscode = 203
	// no_content
	SoupKnownStatusCodeNoContent Knownstatuscode = 204
	// reset_content
	SoupKnownStatusCodeResetContent Knownstatuscode = 205
	// partial_content
	SoupKnownStatusCodePartialContent Knownstatuscode = 206
	// multi_status
	SoupKnownStatusCodeMultiStatus Knownstatuscode = 207
	// multiple_choices
	SoupKnownStatusCodeMultipleChoices Knownstatuscode = 300
	// moved_permanently
	SoupKnownStatusCodeMovedPermanently Knownstatuscode = 301
	// found
	SoupKnownStatusCodeFound Knownstatuscode = 302
	// moved_temporarily
	SoupKnownStatusCodeMovedTemporarily Knownstatuscode = 302
	// see_other
	SoupKnownStatusCodeSeeOther Knownstatuscode = 303
	// not_modified
	SoupKnownStatusCodeNotModified Knownstatuscode = 304
	// use_proxy
	SoupKnownStatusCodeUseProxy Knownstatuscode = 305
	// not_appearing_in_this_protocol
	SoupKnownStatusCodeNotAppearingInThisProtocol Knownstatuscode = 306
	// temporary_redirect
	SoupKnownStatusCodeTemporaryRedirect Knownstatuscode = 307
	// bad_request
	SoupKnownStatusCodeBadRequest Knownstatuscode = 400
	// unauthorized
	SoupKnownStatusCodeUnauthorized Knownstatuscode = 401
	// payment_required
	SoupKnownStatusCodePaymentRequired Knownstatuscode = 402
	// forbidden
	SoupKnownStatusCodeForbidden Knownstatuscode = 403
	// not_found
	SoupKnownStatusCodeNotFound Knownstatuscode = 404
	// method_not_allowed
	SoupKnownStatusCodeMethodNotAllowed Knownstatuscode = 405
	// not_acceptable
	SoupKnownStatusCodeNotAcceptable Knownstatuscode = 406
	// proxy_authentication_required
	SoupKnownStatusCodeProxyAuthenticationRequired Knownstatuscode = 407
	// proxy_unauthorized
	SoupKnownStatusCodeProxyUnauthorized Knownstatuscode = 407
	// request_timeout
	SoupKnownStatusCodeRequestTimeout Knownstatuscode = 408
	// conflict
	SoupKnownStatusCodeConflict Knownstatuscode = 409
	// gone
	SoupKnownStatusCodeGone Knownstatuscode = 410
	// length_required
	SoupKnownStatusCodeLengthRequired Knownstatuscode = 411
	// precondition_failed
	SoupKnownStatusCodePreconditionFailed Knownstatuscode = 412
	// request_entity_too_large
	SoupKnownStatusCodeRequestEntityTooLarge Knownstatuscode = 413
	// request_uri_too_long
	SoupKnownStatusCodeRequestUriTooLong Knownstatuscode = 414
	// unsupported_media_type
	SoupKnownStatusCodeUnsupportedMediaType Knownstatuscode = 415
	// requested_range_not_satisfiable
	SoupKnownStatusCodeRequestedRangeNotSatisfiable Knownstatuscode = 416
	// invalid_range
	SoupKnownStatusCodeInvalidRange Knownstatuscode = 416
	// expectation_failed
	SoupKnownStatusCodeExpectationFailed Knownstatuscode = 417
	// unprocessable_entity
	SoupKnownStatusCodeUnprocessableEntity Knownstatuscode = 422
	// locked
	SoupKnownStatusCodeLocked Knownstatuscode = 423
	// failed_dependency
	SoupKnownStatusCodeFailedDependency Knownstatuscode = 424
	// internal_server_error
	SoupKnownStatusCodeInternalServerError Knownstatuscode = 500
	// not_implemented
	SoupKnownStatusCodeNotImplemented Knownstatuscode = 501
	// bad_gateway
	SoupKnownStatusCodeBadGateway Knownstatuscode = 502
	// service_unavailable
	SoupKnownStatusCodeServiceUnavailable Knownstatuscode = 503
	// gateway_timeout
	SoupKnownStatusCodeGatewayTimeout Knownstatuscode = 504
	// http_version_not_supported
	SoupKnownStatusCodeHttpVersionNotSupported Knownstatuscode = 505
	// insufficient_storage
	SoupKnownStatusCodeInsufficientStorage Knownstatuscode = 507
	// not_extended
	SoupKnownStatusCodeNotExtended Knownstatuscode = 510
)

// Loggerloglevel is a representation of the C type LoggerLogLevel.
type Loggerloglevel int

const (
	// none
	SoupLoggerLogNone Loggerloglevel = 0
	// minimal
	SoupLoggerLogMinimal Loggerloglevel = 1
	// headers
	SoupLoggerLogHeaders Loggerloglevel = 2
	// body
	SoupLoggerLogBody Loggerloglevel = 3
)

// Memoryuse is a representation of the C type MemoryUse.
type Memoryuse int

const (
	// static
	SoupMemoryStatic Memoryuse = 0
	// take
	SoupMemoryTake Memoryuse = 1
	// copy
	SoupMemoryCopy Memoryuse = 2
	// temporary
	SoupMemoryTemporary Memoryuse = 3
)

// Messageheaderstype is a representation of the C type MessageHeadersType.
type Messageheaderstype int

const (
	// request
	SoupMessageHeadersRequest Messageheaderstype = 0
	// response
	SoupMessageHeadersResponse Messageheaderstype = 1
	// multipart
	SoupMessageHeadersMultipart Messageheaderstype = 2
)

// Messagepriority is a representation of the C type MessagePriority.
type Messagepriority int

const (
	// very_low
	SoupMessagePriorityVeryLow Messagepriority = 0
	// low
	SoupMessagePriorityLow Messagepriority = 1
	// normal
	SoupMessagePriorityNormal Messagepriority = 2
	// high
	SoupMessagePriorityHigh Messagepriority = 3
	// very_high
	SoupMessagePriorityVeryHigh Messagepriority = 4
)

// Requesterror is a representation of the C type RequestError.
//
// since 2.42
type Requesterror int

const (
	// bad_uri
	SoupRequestErrorBadUri Requesterror = 0
	// unsupported_uri_scheme
	SoupRequestErrorUnsupportedUriScheme Requesterror = 1
	// parsing
	SoupRequestErrorParsing Requesterror = 2
	// encoding
	SoupRequestErrorEncoding Requesterror = 3
)

// Requestererror is a representation of the C type RequesterError.
type Requestererror int

const (
	// bad_uri
	SoupRequesterErrorBadUri Requestererror = 0
	// unsupported_uri_scheme
	SoupRequesterErrorUnsupportedUriScheme Requestererror = 1
)

// Socketiostatus is a representation of the C type SocketIOStatus.
type Socketiostatus int

const (
	// ok
	SoupSocketOk Socketiostatus = 0
	// would_block
	SoupSocketWouldBlock Socketiostatus = 1
	// eof
	SoupSocketEof Socketiostatus = 2
	// error
	SoupSocketError Socketiostatus = 3
)

// Status is a representation of the C type Status.
type Status int

const (
	// none
	SoupStatusNone Status = 0
	// cancelled
	SoupStatusCancelled Status = 1
	// cant_resolve
	SoupStatusCantResolve Status = 2
	// cant_resolve_proxy
	SoupStatusCantResolveProxy Status = 3
	// cant_connect
	SoupStatusCantConnect Status = 4
	// cant_connect_proxy
	SoupStatusCantConnectProxy Status = 5
	// ssl_failed
	SoupStatusSslFailed Status = 6
	// io_error
	SoupStatusIoError Status = 7
	// malformed
	SoupStatusMalformed Status = 8
	// try_again
	SoupStatusTryAgain Status = 9
	// too_many_redirects
	SoupStatusTooManyRedirects Status = 10
	// tls_failed
	SoupStatusTlsFailed Status = 11
	// continue
	SoupStatusContinue Status = 100
	// switching_protocols
	SoupStatusSwitchingProtocols Status = 101
	// processing
	SoupStatusProcessing Status = 102
	// ok
	SoupStatusOk Status = 200
	// created
	SoupStatusCreated Status = 201
	// accepted
	SoupStatusAccepted Status = 202
	// non_authoritative
	SoupStatusNonAuthoritative Status = 203
	// no_content
	SoupStatusNoContent Status = 204
	// reset_content
	SoupStatusResetContent Status = 205
	// partial_content
	SoupStatusPartialContent Status = 206
	// multi_status
	SoupStatusMultiStatus Status = 207
	// multiple_choices
	SoupStatusMultipleChoices Status = 300
	// moved_permanently
	SoupStatusMovedPermanently Status = 301
	// found
	SoupStatusFound Status = 302
	// moved_temporarily
	SoupStatusMovedTemporarily Status = 302
	// see_other
	SoupStatusSeeOther Status = 303
	// not_modified
	SoupStatusNotModified Status = 304
	// use_proxy
	SoupStatusUseProxy Status = 305
	// not_appearing_in_this_protocol
	SoupStatusNotAppearingInThisProtocol Status = 306
	// temporary_redirect
	SoupStatusTemporaryRedirect Status = 307
	// bad_request
	SoupStatusBadRequest Status = 400
	// unauthorized
	SoupStatusUnauthorized Status = 401
	// payment_required
	SoupStatusPaymentRequired Status = 402
	// forbidden
	SoupStatusForbidden Status = 403
	// not_found
	SoupStatusNotFound Status = 404
	// method_not_allowed
	SoupStatusMethodNotAllowed Status = 405
	// not_acceptable
	SoupStatusNotAcceptable Status = 406
	// proxy_authentication_required
	SoupStatusProxyAuthenticationRequired Status = 407
	// proxy_unauthorized
	SoupStatusProxyUnauthorized Status = 407
	// request_timeout
	SoupStatusRequestTimeout Status = 408
	// conflict
	SoupStatusConflict Status = 409
	// gone
	SoupStatusGone Status = 410
	// length_required
	SoupStatusLengthRequired Status = 411
	// precondition_failed
	SoupStatusPreconditionFailed Status = 412
	// request_entity_too_large
	SoupStatusRequestEntityTooLarge Status = 413
	// request_uri_too_long
	SoupStatusRequestUriTooLong Status = 414
	// unsupported_media_type
	SoupStatusUnsupportedMediaType Status = 415
	// requested_range_not_satisfiable
	SoupStatusRequestedRangeNotSatisfiable Status = 416
	// invalid_range
	SoupStatusInvalidRange Status = 416
	// expectation_failed
	SoupStatusExpectationFailed Status = 417
	// unprocessable_entity
	SoupStatusUnprocessableEntity Status = 422
	// locked
	SoupStatusLocked Status = 423
	// failed_dependency
	SoupStatusFailedDependency Status = 424
	// internal_server_error
	SoupStatusInternalServerError Status = 500
	// not_implemented
	SoupStatusNotImplemented Status = 501
	// bad_gateway
	SoupStatusBadGateway Status = 502
	// service_unavailable
	SoupStatusServiceUnavailable Status = 503
	// gateway_timeout
	SoupStatusGatewayTimeout Status = 504
	// http_version_not_supported
	SoupStatusHttpVersionNotSupported Status = 505
	// insufficient_storage
	SoupStatusInsufficientStorage Status = 507
	// not_extended
	SoupStatusNotExtended Status = 510
)

// Tlderror is a representation of the C type TLDError.
//
// since 2.40
type Tlderror int

const (
	// invalid_hostname
	SoupTldErrorInvalidHostname Tlderror = 0
	// is_ip_address
	SoupTldErrorIsIpAddress Tlderror = 1
	// not_enough_domains
	SoupTldErrorNotEnoughDomains Tlderror = 2
	// no_base_domain
	SoupTldErrorNoBaseDomain Tlderror = 3
)

// Websocketclosecode is a representation of the C type WebsocketCloseCode.
//
// since 2.50
type Websocketclosecode int

const (
	// normal
	SoupWebsocketCloseNormal Websocketclosecode = 1000
	// going_away
	SoupWebsocketCloseGoingAway Websocketclosecode = 1001
	// protocol_error
	SoupWebsocketCloseProtocolError Websocketclosecode = 1002
	// unsupported_data
	SoupWebsocketCloseUnsupportedData Websocketclosecode = 1003
	// no_status
	SoupWebsocketCloseNoStatus Websocketclosecode = 1005
	// abnormal
	SoupWebsocketCloseAbnormal Websocketclosecode = 1006
	// bad_data
	SoupWebsocketCloseBadData Websocketclosecode = 1007
	// policy_violation
	SoupWebsocketClosePolicyViolation Websocketclosecode = 1008
	// too_big
	SoupWebsocketCloseTooBig Websocketclosecode = 1009
	// no_extension
	SoupWebsocketCloseNoExtension Websocketclosecode = 1010
	// server_error
	SoupWebsocketCloseServerError Websocketclosecode = 1011
	// tls_handshake
	SoupWebsocketCloseTlsHandshake Websocketclosecode = 1015
)

// Websocketconnectiontype is a representation of the C type WebsocketConnectionType.
//
// since 2.50
type Websocketconnectiontype int

const (
	// unknown
	SoupWebsocketConnectionUnknown Websocketconnectiontype = 0
	// client
	SoupWebsocketConnectionClient Websocketconnectiontype = 1
	// server
	SoupWebsocketConnectionServer Websocketconnectiontype = 2
)

// Websocketdatatype is a representation of the C type WebsocketDataType.
//
// since 2.50
type Websocketdatatype int

const (
	// text
	SoupWebsocketDataText Websocketdatatype = 1
	// binary
	SoupWebsocketDataBinary Websocketdatatype = 2
)

// Websocketerror is a representation of the C type WebsocketError.
//
// since 2.50
type Websocketerror int

const (
	// failed
	SoupWebsocketErrorFailed Websocketerror = 0
	// not_websocket
	SoupWebsocketErrorNotWebsocket Websocketerror = 1
	// bad_handshake
	SoupWebsocketErrorBadHandshake Websocketerror = 2
	// bad_origin
	SoupWebsocketErrorBadOrigin Websocketerror = 3
)

// Websocketstate is a representation of the C type WebsocketState.
//
// since 2.50
type Websocketstate int

const (
	// open
	SoupWebsocketStateOpen Websocketstate = 1
	// closing
	SoupWebsocketStateClosing Websocketstate = 2
	// closed
	SoupWebsocketStateClosed Websocketstate = 3
)

// Xmlrpcerror is a representation of the C type XMLRPCError.
type Xmlrpcerror int

const (
	// arguments
	SoupXmlrpcErrorArguments Xmlrpcerror = 0
	// retval
	SoupXmlrpcErrorRetval Xmlrpcerror = 1
)

// Xmlrpcfault is a representation of the C type XMLRPCFault.
type Xmlrpcfault int

const (
	// parse_error_not_well_formed
	SoupXmlrpcFaultParseErrorNotWellFormed Xmlrpcfault = -32700
	// parse_error_unsupported_encoding
	SoupXmlrpcFaultParseErrorUnsupportedEncoding Xmlrpcfault = -32701
	// parse_error_invalid_character_for_encoding
	SoupXmlrpcFaultParseErrorInvalidCharacterForEncoding Xmlrpcfault = -32702
	// server_error_invalid_xml_rpc
	SoupXmlrpcFaultServerErrorInvalidXmlRpc Xmlrpcfault = -32600
	// server_error_requested_method_not_found
	SoupXmlrpcFaultServerErrorRequestedMethodNotFound Xmlrpcfault = -32601
	// server_error_invalid_method_parameters
	SoupXmlrpcFaultServerErrorInvalidMethodParameters Xmlrpcfault = -32602
	// server_error_internal_xml_rpc_error
	SoupXmlrpcFaultServerErrorInternalXmlRpcError Xmlrpcfault = -32603
	// application_error
	SoupXmlrpcFaultApplicationError Xmlrpcfault = -32500
	// system_error
	SoupXmlrpcFaultSystemError Xmlrpcfault = -32400
	// transport_error
	SoupXmlrpcFaultTransportError Xmlrpcfault = -32300
)
