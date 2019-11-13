// Code generated - DO NOT EDIT.

package soup

// Addressfamily is a representation of the C type AddressFamily.
type Addressfamily int

const (
	Addressfamily_Invalid Addressfamily = -1
	Addressfamily_Ipv4    Addressfamily = 2
	Addressfamily_Ipv6    Addressfamily = 10
)

// Cacheresponse is a representation of the C type CacheResponse.
type Cacheresponse int

const (
	Cacheresponse_Fresh           Cacheresponse = 0
	Cacheresponse_NeedsValidation Cacheresponse = 1
	Cacheresponse_Stale           Cacheresponse = 2
)

// Cachetype is a representation of the C type CacheType.
//
// since 2.34
type Cachetype int

const (
	Cachetype_SingleUser Cachetype = 0
	Cachetype_Shared     Cachetype = 1
)

// Connectionstate is a representation of the C type ConnectionState.
type Connectionstate int

const (
	Connectionstate_New                Connectionstate = 0
	Connectionstate_Connecting         Connectionstate = 1
	Connectionstate_Idle               Connectionstate = 2
	Connectionstate_InUse              Connectionstate = 3
	Connectionstate_RemoteDisconnected Connectionstate = 4
	Connectionstate_Disconnected       Connectionstate = 5
)

// Cookiejaracceptpolicy is a representation of the C type CookieJarAcceptPolicy.
//
// since 2.30
type Cookiejaracceptpolicy int

const (
	Cookiejaracceptpolicy_Always       Cookiejaracceptpolicy = 0
	Cookiejaracceptpolicy_Never        Cookiejaracceptpolicy = 1
	Cookiejaracceptpolicy_NoThirdParty Cookiejaracceptpolicy = 2
)

// Dateformat is a representation of the C type DateFormat.
type Dateformat int

const (
	Dateformat_Http           Dateformat = 1
	Dateformat_Cookie         Dateformat = 2
	Dateformat_Rfc2822        Dateformat = 3
	Dateformat_Iso8601Compact Dateformat = 4
	Dateformat_Iso8601Full    Dateformat = 5
	Dateformat_Iso8601        Dateformat = 5
	Dateformat_Iso8601Xmlrpc  Dateformat = 6
)

// Encoding is a representation of the C type Encoding.
type Encoding int

const (
	Encoding_Unrecognized  Encoding = 0
	Encoding_None          Encoding = 1
	Encoding_ContentLength Encoding = 2
	Encoding_Eof           Encoding = 3
	Encoding_Chunked       Encoding = 4
	Encoding_Byteranges    Encoding = 5
)

// Httpversion is a representation of the C type HTTPVersion.
type Httpversion int

const (
	Httpversion_Http10 Httpversion = 0
	Httpversion_Http11 Httpversion = 1
)

// Knownstatuscode is a representation of the C type KnownStatusCode.
type Knownstatuscode int

const (
	Knownstatuscode_None                         Knownstatuscode = 0
	Knownstatuscode_Cancelled                    Knownstatuscode = 1
	Knownstatuscode_CantResolve                  Knownstatuscode = 2
	Knownstatuscode_CantResolveProxy             Knownstatuscode = 3
	Knownstatuscode_CantConnect                  Knownstatuscode = 4
	Knownstatuscode_CantConnectProxy             Knownstatuscode = 5
	Knownstatuscode_SslFailed                    Knownstatuscode = 6
	Knownstatuscode_IoError                      Knownstatuscode = 7
	Knownstatuscode_Malformed                    Knownstatuscode = 8
	Knownstatuscode_TryAgain                     Knownstatuscode = 9
	Knownstatuscode_TooManyRedirects             Knownstatuscode = 10
	Knownstatuscode_TlsFailed                    Knownstatuscode = 11
	Knownstatuscode_Continue                     Knownstatuscode = 100
	Knownstatuscode_SwitchingProtocols           Knownstatuscode = 101
	Knownstatuscode_Processing                   Knownstatuscode = 102
	Knownstatuscode_Ok                           Knownstatuscode = 200
	Knownstatuscode_Created                      Knownstatuscode = 201
	Knownstatuscode_Accepted                     Knownstatuscode = 202
	Knownstatuscode_NonAuthoritative             Knownstatuscode = 203
	Knownstatuscode_NoContent                    Knownstatuscode = 204
	Knownstatuscode_ResetContent                 Knownstatuscode = 205
	Knownstatuscode_PartialContent               Knownstatuscode = 206
	Knownstatuscode_MultiStatus                  Knownstatuscode = 207
	Knownstatuscode_MultipleChoices              Knownstatuscode = 300
	Knownstatuscode_MovedPermanently             Knownstatuscode = 301
	Knownstatuscode_Found                        Knownstatuscode = 302
	Knownstatuscode_MovedTemporarily             Knownstatuscode = 302
	Knownstatuscode_SeeOther                     Knownstatuscode = 303
	Knownstatuscode_NotModified                  Knownstatuscode = 304
	Knownstatuscode_UseProxy                     Knownstatuscode = 305
	Knownstatuscode_NotAppearingInThisProtocol   Knownstatuscode = 306
	Knownstatuscode_TemporaryRedirect            Knownstatuscode = 307
	Knownstatuscode_BadRequest                   Knownstatuscode = 400
	Knownstatuscode_Unauthorized                 Knownstatuscode = 401
	Knownstatuscode_PaymentRequired              Knownstatuscode = 402
	Knownstatuscode_Forbidden                    Knownstatuscode = 403
	Knownstatuscode_NotFound                     Knownstatuscode = 404
	Knownstatuscode_MethodNotAllowed             Knownstatuscode = 405
	Knownstatuscode_NotAcceptable                Knownstatuscode = 406
	Knownstatuscode_ProxyAuthenticationRequired  Knownstatuscode = 407
	Knownstatuscode_ProxyUnauthorized            Knownstatuscode = 407
	Knownstatuscode_RequestTimeout               Knownstatuscode = 408
	Knownstatuscode_Conflict                     Knownstatuscode = 409
	Knownstatuscode_Gone                         Knownstatuscode = 410
	Knownstatuscode_LengthRequired               Knownstatuscode = 411
	Knownstatuscode_PreconditionFailed           Knownstatuscode = 412
	Knownstatuscode_RequestEntityTooLarge        Knownstatuscode = 413
	Knownstatuscode_RequestUriTooLong            Knownstatuscode = 414
	Knownstatuscode_UnsupportedMediaType         Knownstatuscode = 415
	Knownstatuscode_RequestedRangeNotSatisfiable Knownstatuscode = 416
	Knownstatuscode_InvalidRange                 Knownstatuscode = 416
	Knownstatuscode_ExpectationFailed            Knownstatuscode = 417
	Knownstatuscode_UnprocessableEntity          Knownstatuscode = 422
	Knownstatuscode_Locked                       Knownstatuscode = 423
	Knownstatuscode_FailedDependency             Knownstatuscode = 424
	Knownstatuscode_InternalServerError          Knownstatuscode = 500
	Knownstatuscode_NotImplemented               Knownstatuscode = 501
	Knownstatuscode_BadGateway                   Knownstatuscode = 502
	Knownstatuscode_ServiceUnavailable           Knownstatuscode = 503
	Knownstatuscode_GatewayTimeout               Knownstatuscode = 504
	Knownstatuscode_HttpVersionNotSupported      Knownstatuscode = 505
	Knownstatuscode_InsufficientStorage          Knownstatuscode = 507
	Knownstatuscode_NotExtended                  Knownstatuscode = 510
)

// Loggerloglevel is a representation of the C type LoggerLogLevel.
type Loggerloglevel int

const (
	Loggerloglevel_None    Loggerloglevel = 0
	Loggerloglevel_Minimal Loggerloglevel = 1
	Loggerloglevel_Headers Loggerloglevel = 2
	Loggerloglevel_Body    Loggerloglevel = 3
)

// Memoryuse is a representation of the C type MemoryUse.
type Memoryuse int

const (
	Memoryuse_Static    Memoryuse = 0
	Memoryuse_Take      Memoryuse = 1
	Memoryuse_Copy      Memoryuse = 2
	Memoryuse_Temporary Memoryuse = 3
)

// Messageheaderstype is a representation of the C type MessageHeadersType.
type Messageheaderstype int

const (
	Messageheaderstype_Request   Messageheaderstype = 0
	Messageheaderstype_Response  Messageheaderstype = 1
	Messageheaderstype_Multipart Messageheaderstype = 2
)

// Messagepriority is a representation of the C type MessagePriority.
type Messagepriority int

const (
	Messagepriority_VeryLow  Messagepriority = 0
	Messagepriority_Low      Messagepriority = 1
	Messagepriority_Normal   Messagepriority = 2
	Messagepriority_High     Messagepriority = 3
	Messagepriority_VeryHigh Messagepriority = 4
)

// Requesterror is a representation of the C type RequestError.
//
// since 2.42
type Requesterror int

const (
	Requesterror_BadUri               Requesterror = 0
	Requesterror_UnsupportedUriScheme Requesterror = 1
	Requesterror_Parsing              Requesterror = 2
	Requesterror_Encoding             Requesterror = 3
)

// Requestererror is a representation of the C type RequesterError.
type Requestererror int

const (
	Requestererror_BadUri               Requestererror = 0
	Requestererror_UnsupportedUriScheme Requestererror = 1
)

// Socketiostatus is a representation of the C type SocketIOStatus.
type Socketiostatus int

const (
	Socketiostatus_Ok         Socketiostatus = 0
	Socketiostatus_WouldBlock Socketiostatus = 1
	Socketiostatus_Eof        Socketiostatus = 2
	Socketiostatus_Error      Socketiostatus = 3
)

// Status is a representation of the C type Status.
type Status int

const (
	Status_None                         Status = 0
	Status_Cancelled                    Status = 1
	Status_CantResolve                  Status = 2
	Status_CantResolveProxy             Status = 3
	Status_CantConnect                  Status = 4
	Status_CantConnectProxy             Status = 5
	Status_SslFailed                    Status = 6
	Status_IoError                      Status = 7
	Status_Malformed                    Status = 8
	Status_TryAgain                     Status = 9
	Status_TooManyRedirects             Status = 10
	Status_TlsFailed                    Status = 11
	Status_Continue                     Status = 100
	Status_SwitchingProtocols           Status = 101
	Status_Processing                   Status = 102
	Status_Ok                           Status = 200
	Status_Created                      Status = 201
	Status_Accepted                     Status = 202
	Status_NonAuthoritative             Status = 203
	Status_NoContent                    Status = 204
	Status_ResetContent                 Status = 205
	Status_PartialContent               Status = 206
	Status_MultiStatus                  Status = 207
	Status_MultipleChoices              Status = 300
	Status_MovedPermanently             Status = 301
	Status_Found                        Status = 302
	Status_MovedTemporarily             Status = 302
	Status_SeeOther                     Status = 303
	Status_NotModified                  Status = 304
	Status_UseProxy                     Status = 305
	Status_NotAppearingInThisProtocol   Status = 306
	Status_TemporaryRedirect            Status = 307
	Status_BadRequest                   Status = 400
	Status_Unauthorized                 Status = 401
	Status_PaymentRequired              Status = 402
	Status_Forbidden                    Status = 403
	Status_NotFound                     Status = 404
	Status_MethodNotAllowed             Status = 405
	Status_NotAcceptable                Status = 406
	Status_ProxyAuthenticationRequired  Status = 407
	Status_ProxyUnauthorized            Status = 407
	Status_RequestTimeout               Status = 408
	Status_Conflict                     Status = 409
	Status_Gone                         Status = 410
	Status_LengthRequired               Status = 411
	Status_PreconditionFailed           Status = 412
	Status_RequestEntityTooLarge        Status = 413
	Status_RequestUriTooLong            Status = 414
	Status_UnsupportedMediaType         Status = 415
	Status_RequestedRangeNotSatisfiable Status = 416
	Status_InvalidRange                 Status = 416
	Status_ExpectationFailed            Status = 417
	Status_UnprocessableEntity          Status = 422
	Status_Locked                       Status = 423
	Status_FailedDependency             Status = 424
	Status_InternalServerError          Status = 500
	Status_NotImplemented               Status = 501
	Status_BadGateway                   Status = 502
	Status_ServiceUnavailable           Status = 503
	Status_GatewayTimeout               Status = 504
	Status_HttpVersionNotSupported      Status = 505
	Status_InsufficientStorage          Status = 507
	Status_NotExtended                  Status = 510
)

// Tlderror is a representation of the C type TLDError.
//
// since 2.40
type Tlderror int

const (
	Tlderror_InvalidHostname  Tlderror = 0
	Tlderror_IsIpAddress      Tlderror = 1
	Tlderror_NotEnoughDomains Tlderror = 2
	Tlderror_NoBaseDomain     Tlderror = 3
)

// Websocketclosecode is a representation of the C type WebsocketCloseCode.
//
// since 2.50
type Websocketclosecode int

const (
	Websocketclosecode_Normal          Websocketclosecode = 1000
	Websocketclosecode_GoingAway       Websocketclosecode = 1001
	Websocketclosecode_ProtocolError   Websocketclosecode = 1002
	Websocketclosecode_UnsupportedData Websocketclosecode = 1003
	Websocketclosecode_NoStatus        Websocketclosecode = 1005
	Websocketclosecode_Abnormal        Websocketclosecode = 1006
	Websocketclosecode_BadData         Websocketclosecode = 1007
	Websocketclosecode_PolicyViolation Websocketclosecode = 1008
	Websocketclosecode_TooBig          Websocketclosecode = 1009
	Websocketclosecode_NoExtension     Websocketclosecode = 1010
	Websocketclosecode_ServerError     Websocketclosecode = 1011
	Websocketclosecode_TlsHandshake    Websocketclosecode = 1015
)

// Websocketconnectiontype is a representation of the C type WebsocketConnectionType.
//
// since 2.50
type Websocketconnectiontype int

const (
	Websocketconnectiontype_Unknown Websocketconnectiontype = 0
	Websocketconnectiontype_Client  Websocketconnectiontype = 1
	Websocketconnectiontype_Server  Websocketconnectiontype = 2
)

// Websocketdatatype is a representation of the C type WebsocketDataType.
//
// since 2.50
type Websocketdatatype int

const (
	Websocketdatatype_Text   Websocketdatatype = 1
	Websocketdatatype_Binary Websocketdatatype = 2
)

// Websocketerror is a representation of the C type WebsocketError.
//
// since 2.50
type Websocketerror int

const (
	Websocketerror_Failed       Websocketerror = 0
	Websocketerror_NotWebsocket Websocketerror = 1
	Websocketerror_BadHandshake Websocketerror = 2
	Websocketerror_BadOrigin    Websocketerror = 3
)

// Websocketstate is a representation of the C type WebsocketState.
//
// since 2.50
type Websocketstate int

const (
	Websocketstate_Open    Websocketstate = 1
	Websocketstate_Closing Websocketstate = 2
	Websocketstate_Closed  Websocketstate = 3
)

// Xmlrpcerror is a representation of the C type XMLRPCError.
type Xmlrpcerror int

const (
	Xmlrpcerror_Arguments Xmlrpcerror = 0
	Xmlrpcerror_Retval    Xmlrpcerror = 1
)

// Xmlrpcfault is a representation of the C type XMLRPCFault.
type Xmlrpcfault int

const (
	Xmlrpcfault_ParseErrorNotWellFormed               Xmlrpcfault = -32700
	Xmlrpcfault_ParseErrorUnsupportedEncoding         Xmlrpcfault = -32701
	Xmlrpcfault_ParseErrorInvalidCharacterForEncoding Xmlrpcfault = -32702
	Xmlrpcfault_ServerErrorInvalidXmlRpc              Xmlrpcfault = -32600
	Xmlrpcfault_ServerErrorRequestedMethodNotFound    Xmlrpcfault = -32601
	Xmlrpcfault_ServerErrorInvalidMethodParameters    Xmlrpcfault = -32602
	Xmlrpcfault_ServerErrorInternalXmlRpcError        Xmlrpcfault = -32603
	Xmlrpcfault_ApplicationError                      Xmlrpcfault = -32500
	Xmlrpcfault_SystemError                           Xmlrpcfault = -32400
	Xmlrpcfault_TransportError                        Xmlrpcfault = -32300
)
