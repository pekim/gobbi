// Code generated - DO NOT EDIT.

package soup

type AddressFamily int32

const (
	AddressFamily_Invalid AddressFamily = int64(-1)
	AddressFamily_Ipv4    AddressFamily = int64(2)
	AddressFamily_Ipv6    AddressFamily = int64(10)
)

type CacheResponse uint32

const (
	CacheResponse_Fresh           CacheResponse = int64(0)
	CacheResponse_NeedsValidation CacheResponse = int64(1)
	CacheResponse_Stale           CacheResponse = int64(2)
)

type CacheType uint32

const (
	CacheType_SingleUser CacheType = int64(0)
	CacheType_Shared     CacheType = int64(1)
)

type ConnectionState uint32

const (
	ConnectionState_New                ConnectionState = int64(0)
	ConnectionState_Connecting         ConnectionState = int64(1)
	ConnectionState_Idle               ConnectionState = int64(2)
	ConnectionState_InUse              ConnectionState = int64(3)
	ConnectionState_RemoteDisconnected ConnectionState = int64(4)
	ConnectionState_Disconnected       ConnectionState = int64(5)
)

type CookieJarAcceptPolicy uint32

const (
	CookieJarAcceptPolicy_Always       CookieJarAcceptPolicy = int64(0)
	CookieJarAcceptPolicy_Never        CookieJarAcceptPolicy = int64(1)
	CookieJarAcceptPolicy_NoThirdParty CookieJarAcceptPolicy = int64(2)
)

type DateFormat uint32

const (
	DateFormat_Http           DateFormat = int64(1)
	DateFormat_Cookie         DateFormat = int64(2)
	DateFormat_Rfc2822        DateFormat = int64(3)
	DateFormat_Iso8601Compact DateFormat = int64(4)
	DateFormat_Iso8601Full    DateFormat = int64(5)
	DateFormat_Iso8601        DateFormat = int64(5)
	DateFormat_Iso8601Xmlrpc  DateFormat = int64(6)
)

type Encoding uint32

const (
	Encoding_Unrecognized  Encoding = int64(0)
	Encoding_None          Encoding = int64(1)
	Encoding_ContentLength Encoding = int64(2)
	Encoding_Eof           Encoding = int64(3)
	Encoding_Chunked       Encoding = int64(4)
	Encoding_Byteranges    Encoding = int64(5)
)

type HTTPVersion uint32

const (
	HTTPVersion_Http10 HTTPVersion = int64(0)
	HTTPVersion_Http11 HTTPVersion = int64(1)
)

type KnownStatusCode uint32

const (
	KnownStatusCode_None                         KnownStatusCode = int64(0)
	KnownStatusCode_Cancelled                    KnownStatusCode = int64(1)
	KnownStatusCode_CantResolve                  KnownStatusCode = int64(2)
	KnownStatusCode_CantResolveProxy             KnownStatusCode = int64(3)
	KnownStatusCode_CantConnect                  KnownStatusCode = int64(4)
	KnownStatusCode_CantConnectProxy             KnownStatusCode = int64(5)
	KnownStatusCode_SslFailed                    KnownStatusCode = int64(6)
	KnownStatusCode_IoError                      KnownStatusCode = int64(7)
	KnownStatusCode_Malformed                    KnownStatusCode = int64(8)
	KnownStatusCode_TryAgain                     KnownStatusCode = int64(9)
	KnownStatusCode_TooManyRedirects             KnownStatusCode = int64(10)
	KnownStatusCode_TlsFailed                    KnownStatusCode = int64(11)
	KnownStatusCode_Continue                     KnownStatusCode = int64(100)
	KnownStatusCode_SwitchingProtocols           KnownStatusCode = int64(101)
	KnownStatusCode_Processing                   KnownStatusCode = int64(102)
	KnownStatusCode_Ok                           KnownStatusCode = int64(200)
	KnownStatusCode_Created                      KnownStatusCode = int64(201)
	KnownStatusCode_Accepted                     KnownStatusCode = int64(202)
	KnownStatusCode_NonAuthoritative             KnownStatusCode = int64(203)
	KnownStatusCode_NoContent                    KnownStatusCode = int64(204)
	KnownStatusCode_ResetContent                 KnownStatusCode = int64(205)
	KnownStatusCode_PartialContent               KnownStatusCode = int64(206)
	KnownStatusCode_MultiStatus                  KnownStatusCode = int64(207)
	KnownStatusCode_MultipleChoices              KnownStatusCode = int64(300)
	KnownStatusCode_MovedPermanently             KnownStatusCode = int64(301)
	KnownStatusCode_Found                        KnownStatusCode = int64(302)
	KnownStatusCode_MovedTemporarily             KnownStatusCode = int64(302)
	KnownStatusCode_SeeOther                     KnownStatusCode = int64(303)
	KnownStatusCode_NotModified                  KnownStatusCode = int64(304)
	KnownStatusCode_UseProxy                     KnownStatusCode = int64(305)
	KnownStatusCode_NotAppearingInThisProtocol   KnownStatusCode = int64(306)
	KnownStatusCode_TemporaryRedirect            KnownStatusCode = int64(307)
	KnownStatusCode_BadRequest                   KnownStatusCode = int64(400)
	KnownStatusCode_Unauthorized                 KnownStatusCode = int64(401)
	KnownStatusCode_PaymentRequired              KnownStatusCode = int64(402)
	KnownStatusCode_Forbidden                    KnownStatusCode = int64(403)
	KnownStatusCode_NotFound                     KnownStatusCode = int64(404)
	KnownStatusCode_MethodNotAllowed             KnownStatusCode = int64(405)
	KnownStatusCode_NotAcceptable                KnownStatusCode = int64(406)
	KnownStatusCode_ProxyAuthenticationRequired  KnownStatusCode = int64(407)
	KnownStatusCode_ProxyUnauthorized            KnownStatusCode = int64(407)
	KnownStatusCode_RequestTimeout               KnownStatusCode = int64(408)
	KnownStatusCode_Conflict                     KnownStatusCode = int64(409)
	KnownStatusCode_Gone                         KnownStatusCode = int64(410)
	KnownStatusCode_LengthRequired               KnownStatusCode = int64(411)
	KnownStatusCode_PreconditionFailed           KnownStatusCode = int64(412)
	KnownStatusCode_RequestEntityTooLarge        KnownStatusCode = int64(413)
	KnownStatusCode_RequestUriTooLong            KnownStatusCode = int64(414)
	KnownStatusCode_UnsupportedMediaType         KnownStatusCode = int64(415)
	KnownStatusCode_RequestedRangeNotSatisfiable KnownStatusCode = int64(416)
	KnownStatusCode_InvalidRange                 KnownStatusCode = int64(416)
	KnownStatusCode_ExpectationFailed            KnownStatusCode = int64(417)
	KnownStatusCode_UnprocessableEntity          KnownStatusCode = int64(422)
	KnownStatusCode_Locked                       KnownStatusCode = int64(423)
	KnownStatusCode_FailedDependency             KnownStatusCode = int64(424)
	KnownStatusCode_InternalServerError          KnownStatusCode = int64(500)
	KnownStatusCode_NotImplemented               KnownStatusCode = int64(501)
	KnownStatusCode_BadGateway                   KnownStatusCode = int64(502)
	KnownStatusCode_ServiceUnavailable           KnownStatusCode = int64(503)
	KnownStatusCode_GatewayTimeout               KnownStatusCode = int64(504)
	KnownStatusCode_HttpVersionNotSupported      KnownStatusCode = int64(505)
	KnownStatusCode_InsufficientStorage          KnownStatusCode = int64(507)
	KnownStatusCode_NotExtended                  KnownStatusCode = int64(510)
)

type LoggerLogLevel uint32

const (
	LoggerLogLevel_None    LoggerLogLevel = int64(0)
	LoggerLogLevel_Minimal LoggerLogLevel = int64(1)
	LoggerLogLevel_Headers LoggerLogLevel = int64(2)
	LoggerLogLevel_Body    LoggerLogLevel = int64(3)
)

type MemoryUse uint32

const (
	MemoryUse_Static    MemoryUse = int64(0)
	MemoryUse_Take      MemoryUse = int64(1)
	MemoryUse_Copy      MemoryUse = int64(2)
	MemoryUse_Temporary MemoryUse = int64(3)
)

type MessageHeadersType uint32

const (
	MessageHeadersType_Request   MessageHeadersType = int64(0)
	MessageHeadersType_Response  MessageHeadersType = int64(1)
	MessageHeadersType_Multipart MessageHeadersType = int64(2)
)

type MessagePriority uint32

const (
	MessagePriority_VeryLow  MessagePriority = int64(0)
	MessagePriority_Low      MessagePriority = int64(1)
	MessagePriority_Normal   MessagePriority = int64(2)
	MessagePriority_High     MessagePriority = int64(3)
	MessagePriority_VeryHigh MessagePriority = int64(4)
)

type RequestError uint32

const (
	RequestError_BadUri               RequestError = int64(0)
	RequestError_UnsupportedUriScheme RequestError = int64(1)
	RequestError_Parsing              RequestError = int64(2)
	RequestError_Encoding             RequestError = int64(3)
)

type RequesterError uint32

const (
	RequesterError_BadUri               RequesterError = int64(0)
	RequesterError_UnsupportedUriScheme RequesterError = int64(1)
)

type SocketIOStatus uint32

const (
	SocketIOStatus_Ok         SocketIOStatus = int64(0)
	SocketIOStatus_WouldBlock SocketIOStatus = int64(1)
	SocketIOStatus_Eof        SocketIOStatus = int64(2)
	SocketIOStatus_Error      SocketIOStatus = int64(3)
)

type Status uint32

const (
	Status_None                         Status = int64(0)
	Status_Cancelled                    Status = int64(1)
	Status_CantResolve                  Status = int64(2)
	Status_CantResolveProxy             Status = int64(3)
	Status_CantConnect                  Status = int64(4)
	Status_CantConnectProxy             Status = int64(5)
	Status_SslFailed                    Status = int64(6)
	Status_IoError                      Status = int64(7)
	Status_Malformed                    Status = int64(8)
	Status_TryAgain                     Status = int64(9)
	Status_TooManyRedirects             Status = int64(10)
	Status_TlsFailed                    Status = int64(11)
	Status_Continue                     Status = int64(100)
	Status_SwitchingProtocols           Status = int64(101)
	Status_Processing                   Status = int64(102)
	Status_Ok                           Status = int64(200)
	Status_Created                      Status = int64(201)
	Status_Accepted                     Status = int64(202)
	Status_NonAuthoritative             Status = int64(203)
	Status_NoContent                    Status = int64(204)
	Status_ResetContent                 Status = int64(205)
	Status_PartialContent               Status = int64(206)
	Status_MultiStatus                  Status = int64(207)
	Status_MultipleChoices              Status = int64(300)
	Status_MovedPermanently             Status = int64(301)
	Status_Found                        Status = int64(302)
	Status_MovedTemporarily             Status = int64(302)
	Status_SeeOther                     Status = int64(303)
	Status_NotModified                  Status = int64(304)
	Status_UseProxy                     Status = int64(305)
	Status_NotAppearingInThisProtocol   Status = int64(306)
	Status_TemporaryRedirect            Status = int64(307)
	Status_BadRequest                   Status = int64(400)
	Status_Unauthorized                 Status = int64(401)
	Status_PaymentRequired              Status = int64(402)
	Status_Forbidden                    Status = int64(403)
	Status_NotFound                     Status = int64(404)
	Status_MethodNotAllowed             Status = int64(405)
	Status_NotAcceptable                Status = int64(406)
	Status_ProxyAuthenticationRequired  Status = int64(407)
	Status_ProxyUnauthorized            Status = int64(407)
	Status_RequestTimeout               Status = int64(408)
	Status_Conflict                     Status = int64(409)
	Status_Gone                         Status = int64(410)
	Status_LengthRequired               Status = int64(411)
	Status_PreconditionFailed           Status = int64(412)
	Status_RequestEntityTooLarge        Status = int64(413)
	Status_RequestUriTooLong            Status = int64(414)
	Status_UnsupportedMediaType         Status = int64(415)
	Status_RequestedRangeNotSatisfiable Status = int64(416)
	Status_InvalidRange                 Status = int64(416)
	Status_ExpectationFailed            Status = int64(417)
	Status_UnprocessableEntity          Status = int64(422)
	Status_Locked                       Status = int64(423)
	Status_FailedDependency             Status = int64(424)
	Status_InternalServerError          Status = int64(500)
	Status_NotImplemented               Status = int64(501)
	Status_BadGateway                   Status = int64(502)
	Status_ServiceUnavailable           Status = int64(503)
	Status_GatewayTimeout               Status = int64(504)
	Status_HttpVersionNotSupported      Status = int64(505)
	Status_InsufficientStorage          Status = int64(507)
	Status_NotExtended                  Status = int64(510)
)

type TLDError uint32

const (
	TLDError_InvalidHostname  TLDError = int64(0)
	TLDError_IsIpAddress      TLDError = int64(1)
	TLDError_NotEnoughDomains TLDError = int64(2)
	TLDError_NoBaseDomain     TLDError = int64(3)
	TLDError_NoPslData        TLDError = int64(4)
)

type WebsocketCloseCode uint32

const (
	WebsocketCloseCode_Normal          WebsocketCloseCode = int64(1000)
	WebsocketCloseCode_GoingAway       WebsocketCloseCode = int64(1001)
	WebsocketCloseCode_ProtocolError   WebsocketCloseCode = int64(1002)
	WebsocketCloseCode_UnsupportedData WebsocketCloseCode = int64(1003)
	WebsocketCloseCode_NoStatus        WebsocketCloseCode = int64(1005)
	WebsocketCloseCode_Abnormal        WebsocketCloseCode = int64(1006)
	WebsocketCloseCode_BadData         WebsocketCloseCode = int64(1007)
	WebsocketCloseCode_PolicyViolation WebsocketCloseCode = int64(1008)
	WebsocketCloseCode_TooBig          WebsocketCloseCode = int64(1009)
	WebsocketCloseCode_NoExtension     WebsocketCloseCode = int64(1010)
	WebsocketCloseCode_ServerError     WebsocketCloseCode = int64(1011)
	WebsocketCloseCode_TlsHandshake    WebsocketCloseCode = int64(1015)
)

type WebsocketConnectionType uint32

const (
	WebsocketConnectionType_Unknown WebsocketConnectionType = int64(0)
	WebsocketConnectionType_Client  WebsocketConnectionType = int64(1)
	WebsocketConnectionType_Server  WebsocketConnectionType = int64(2)
)

type WebsocketDataType uint32

const (
	WebsocketDataType_Text   WebsocketDataType = int64(1)
	WebsocketDataType_Binary WebsocketDataType = int64(2)
)

type WebsocketError uint32

const (
	WebsocketError_Failed       WebsocketError = int64(0)
	WebsocketError_NotWebsocket WebsocketError = int64(1)
	WebsocketError_BadHandshake WebsocketError = int64(2)
	WebsocketError_BadOrigin    WebsocketError = int64(3)
)

type WebsocketState uint32

const (
	WebsocketState_Open    WebsocketState = int64(1)
	WebsocketState_Closing WebsocketState = int64(2)
	WebsocketState_Closed  WebsocketState = int64(3)
)

type XMLRPCError uint32

const (
	XMLRPCError_Arguments XMLRPCError = int64(0)
	XMLRPCError_Retval    XMLRPCError = int64(1)
)

type XMLRPCFault int32

const (
	XMLRPCFault_ParseErrorNotWellFormed               XMLRPCFault = int64(-32700)
	XMLRPCFault_ParseErrorUnsupportedEncoding         XMLRPCFault = int64(-32701)
	XMLRPCFault_ParseErrorInvalidCharacterForEncoding XMLRPCFault = int64(-32702)
	XMLRPCFault_ServerErrorInvalidXmlRpc              XMLRPCFault = int64(-32600)
	XMLRPCFault_ServerErrorRequestedMethodNotFound    XMLRPCFault = int64(-32601)
	XMLRPCFault_ServerErrorInvalidMethodParameters    XMLRPCFault = int64(-32602)
	XMLRPCFault_ServerErrorInternalXmlRpcError        XMLRPCFault = int64(-32603)
	XMLRPCFault_ApplicationError                      XMLRPCFault = int64(-32500)
	XMLRPCFault_SystemError                           XMLRPCFault = int64(-32400)
	XMLRPCFault_TransportError                        XMLRPCFault = int64(-32300)
)
