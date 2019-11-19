// Code generated - DO NOT EDIT.

package soup

import gi "github.com/pekim/gobbi/internal/gi"

type AddressClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value '_libsoup_reserved1' : missing Type

	// UNSUPPORTED : C value '_libsoup_reserved2' : missing Type

	// UNSUPPORTED : C value '_libsoup_reserved3' : missing Type

	// UNSUPPORTED : C value '_libsoup_reserved4' : missing Type

}

type AuthClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	SchemeName string
	Strength   uint32
	// UNSUPPORTED : C value 'update' : missing Type

	// UNSUPPORTED : C value 'get_protection_space' : missing Type

	// UNSUPPORTED : C value 'authenticate' : missing Type

	// UNSUPPORTED : C value 'is_authenticated' : missing Type

	// UNSUPPORTED : C value 'get_authorization' : missing Type

	// UNSUPPORTED : C value 'is_ready' : missing Type

	// UNSUPPORTED : C value 'can_authenticate' : missing Type

	// UNSUPPORTED : C value '_libsoup_reserved3' : missing Type

	// UNSUPPORTED : C value '_libsoup_reserved4' : missing Type

}

type AuthDomainBasicClass struct {
	native      uintptr
	ParentClass *AuthDomainClass
	// UNSUPPORTED : C value '_libsoup_reserved1' : missing Type

	// UNSUPPORTED : C value '_libsoup_reserved2' : missing Type

	// UNSUPPORTED : C value '_libsoup_reserved3' : missing Type

	// UNSUPPORTED : C value '_libsoup_reserved4' : missing Type

}

type AuthDomainClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value 'accepts' : missing Type

	// UNSUPPORTED : C value 'challenge' : missing Type

	// UNSUPPORTED : C value 'check_password' : missing Type

	// UNSUPPORTED : C value '_libsoup_reserved2' : missing Type

	// UNSUPPORTED : C value '_libsoup_reserved3' : missing Type

	// UNSUPPORTED : C value '_libsoup_reserved4' : missing Type

}

type AuthDomainDigestClass struct {
	native      uintptr
	ParentClass *AuthDomainClass
	// UNSUPPORTED : C value '_libsoup_reserved1' : missing Type

	// UNSUPPORTED : C value '_libsoup_reserved2' : missing Type

	// UNSUPPORTED : C value '_libsoup_reserved3' : missing Type

	// UNSUPPORTED : C value '_libsoup_reserved4' : missing Type

}

type AuthManagerClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value 'authenticate' : missing Type

}

type AuthManagerPrivate struct {
	native uintptr
}

type Buffer struct {
	native uintptr
	// UNSUPPORTED : C value 'data' : no Go type for 'gpointer'

	Length uintptr
}

// UNSUPPORTED : C value 'soup_buffer_new' : parameter 'use' of type 'MemoryUse' not supported

// UNSUPPORTED : C value 'soup_buffer_new_take' : parameter 'data' has no type

// UNSUPPORTED : C value 'soup_buffer_new_with_owner' : parameter 'data' has no type

type CacheClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value 'get_cacheability' : missing Type

	// UNSUPPORTED : C value '_libsoup_reserved1' : missing Type

	// UNSUPPORTED : C value '_libsoup_reserved2' : missing Type

	// UNSUPPORTED : C value '_libsoup_reserved3' : missing Type

}

type CachePrivate struct {
	native uintptr
}

type ClientContext struct {
	native uintptr
}

type Connection struct {
	native uintptr
}

type ContentDecoderClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value '_libsoup_reserved1' : missing Type

	// UNSUPPORTED : C value '_libsoup_reserved2' : missing Type

	// UNSUPPORTED : C value '_libsoup_reserved3' : missing Type

	// UNSUPPORTED : C value '_libsoup_reserved4' : missing Type

	// UNSUPPORTED : C value '_libsoup_reserved5' : missing Type

}

type ContentDecoderPrivate struct {
	native uintptr
}

type ContentSnifferClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value 'sniff' : missing Type

	// UNSUPPORTED : C value 'get_buffer_size' : missing Type

	// UNSUPPORTED : C value '_libsoup_reserved1' : missing Type

	// UNSUPPORTED : C value '_libsoup_reserved2' : missing Type

	// UNSUPPORTED : C value '_libsoup_reserved3' : missing Type

	// UNSUPPORTED : C value '_libsoup_reserved4' : missing Type

	// UNSUPPORTED : C value '_libsoup_reserved5' : missing Type

}

type ContentSnifferPrivate struct {
	native uintptr
}

type Cookie struct {
	native  uintptr
	Name    string
	Value   string
	Domain  string
	Path    string
	Expires *Date
	// UNSUPPORTED : C value 'secure' : no Go type for 'gboolean'

	// UNSUPPORTED : C value 'http_only' : no Go type for 'gboolean'

}

var newCookieInvoker *gi.Function

// CookieNew is a representation of the C type soup_cookie_new.
func CookieNew(name string, value string, domain string, path string, maxAge int32) *Cookie {
	if newCookieInvoker == nil {
		newCookieInvoker = gi.StructFunctionInvokerNew("Soup", "Cookie", "new")
	}

	var inArgs [5]gi.Argument
	inArgs[0].SetString(name)
	inArgs[1].SetString(value)
	inArgs[2].SetString(domain)
	inArgs[3].SetString(path)
	inArgs[4].SetInt32(maxAge)

	ret := newCookieInvoker.Invoke(inArgs[:], nil)

	retGo := &Cookie{native: ret.Pointer()}

	return retGo
}

type CookieJarClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value 'save' : missing Type

	// UNSUPPORTED : C value 'is_persistent' : missing Type

	// UNSUPPORTED : C value 'changed' : missing Type

	// UNSUPPORTED : C value '_libsoup_reserved1' : missing Type

	// UNSUPPORTED : C value '_libsoup_reserved2' : missing Type

}

type CookieJarDBClass struct {
	native      uintptr
	ParentClass *CookieJarClass
	// UNSUPPORTED : C value '_libsoup_reserved1' : missing Type

	// UNSUPPORTED : C value '_libsoup_reserved2' : missing Type

	// UNSUPPORTED : C value '_libsoup_reserved3' : missing Type

	// UNSUPPORTED : C value '_libsoup_reserved4' : missing Type

}

type CookieJarTextClass struct {
	native      uintptr
	ParentClass *CookieJarClass
	// UNSUPPORTED : C value '_libsoup_reserved1' : missing Type

	// UNSUPPORTED : C value '_libsoup_reserved2' : missing Type

	// UNSUPPORTED : C value '_libsoup_reserved3' : missing Type

	// UNSUPPORTED : C value '_libsoup_reserved4' : missing Type

}

type Date struct {
	native uintptr
	Year   int32
	Month  int32
	Day    int32
	Hour   int32
	Minute int32
	Second int32
	// UNSUPPORTED : C value 'utc' : no Go type for 'gboolean'

	Offset int32
}

var newDateInvoker *gi.Function

// DateNew is a representation of the C type soup_date_new.
func DateNew(year int32, month int32, day int32, hour int32, minute int32, second int32) *Date {
	if newDateInvoker == nil {
		newDateInvoker = gi.StructFunctionInvokerNew("Soup", "Date", "new")
	}

	var inArgs [6]gi.Argument
	inArgs[0].SetInt32(year)
	inArgs[1].SetInt32(month)
	inArgs[2].SetInt32(day)
	inArgs[3].SetInt32(hour)
	inArgs[4].SetInt32(minute)
	inArgs[5].SetInt32(second)

	ret := newDateInvoker.Invoke(inArgs[:], nil)

	retGo := &Date{native: ret.Pointer()}

	return retGo
}

var newFromNowDateInvoker *gi.Function

// DateNewFromNow is a representation of the C type soup_date_new_from_now.
func DateNewFromNow(offsetSeconds int32) *Date {
	if newFromNowDateInvoker == nil {
		newFromNowDateInvoker = gi.StructFunctionInvokerNew("Soup", "Date", "new_from_now")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetInt32(offsetSeconds)

	ret := newFromNowDateInvoker.Invoke(inArgs[:], nil)

	retGo := &Date{native: ret.Pointer()}

	return retGo
}

var newFromStringDateInvoker *gi.Function

// DateNewFromString is a representation of the C type soup_date_new_from_string.
func DateNewFromString(dateString string) *Date {
	if newFromStringDateInvoker == nil {
		newFromStringDateInvoker = gi.StructFunctionInvokerNew("Soup", "Date", "new_from_string")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetString(dateString)

	ret := newFromStringDateInvoker.Invoke(inArgs[:], nil)

	retGo := &Date{native: ret.Pointer()}

	return retGo
}

var newFromTimeTDateInvoker *gi.Function

// DateNewFromTimeT is a representation of the C type soup_date_new_from_time_t.
func DateNewFromTimeT(when int64) *Date {
	if newFromTimeTDateInvoker == nil {
		newFromTimeTDateInvoker = gi.StructFunctionInvokerNew("Soup", "Date", "new_from_time_t")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetInt64(when)

	ret := newFromTimeTDateInvoker.Invoke(inArgs[:], nil)

	retGo := &Date{native: ret.Pointer()}

	return retGo
}

type HSTSEnforcerClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value 'is_persistent' : missing Type

	// UNSUPPORTED : C value 'has_valid_policy' : missing Type

	// UNSUPPORTED : C value 'changed' : missing Type

	// UNSUPPORTED : C value 'hsts_enforced' : missing Type

	// UNSUPPORTED : C value '_libsoup_reserved1' : missing Type

	// UNSUPPORTED : C value '_libsoup_reserved2' : missing Type

	// UNSUPPORTED : C value '_libsoup_reserved3' : missing Type

	// UNSUPPORTED : C value '_libsoup_reserved4' : missing Type

}

type HSTSEnforcerDBClass struct {
	native      uintptr
	ParentClass *HSTSEnforcerClass
	// UNSUPPORTED : C value '_libsoup_reserved1' : missing Type

	// UNSUPPORTED : C value '_libsoup_reserved2' : missing Type

	// UNSUPPORTED : C value '_libsoup_reserved3' : missing Type

	// UNSUPPORTED : C value '_libsoup_reserved4' : missing Type

}

type HSTSEnforcerDBPrivate struct {
	native uintptr
}

type HSTSEnforcerPrivate struct {
	native uintptr
}

type HSTSPolicy struct {
	native  uintptr
	Domain  string
	MaxAge  uint64
	Expires *Date
	// UNSUPPORTED : C value 'include_subdomains' : no Go type for 'gboolean'

}

// UNSUPPORTED : C value 'soup_hsts_policy_new' : parameter 'include_subdomains' of type 'gboolean' not supported

// UNSUPPORTED : C value 'soup_hsts_policy_new_from_response' : parameter 'msg' of type 'Message' not supported

// UNSUPPORTED : C value 'soup_hsts_policy_new_full' : parameter 'expires' of type 'Date' not supported

// UNSUPPORTED : C value 'soup_hsts_policy_new_session_policy' : parameter 'include_subdomains' of type 'gboolean' not supported

type LoggerClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value '_libsoup_reserved1' : missing Type

	// UNSUPPORTED : C value '_libsoup_reserved2' : missing Type

	// UNSUPPORTED : C value '_libsoup_reserved3' : missing Type

	// UNSUPPORTED : C value '_libsoup_reserved4' : missing Type

}

type MessageBody struct {
	native uintptr
	Data   string
	Length int64
}

var newMessageBodyInvoker *gi.Function

// MessageBodyNew is a representation of the C type soup_message_body_new.
func MessageBodyNew() *MessageBody {
	if newMessageBodyInvoker == nil {
		newMessageBodyInvoker = gi.StructFunctionInvokerNew("Soup", "MessageBody", "new")
	}

	ret := newMessageBodyInvoker.Invoke(nil, nil)

	retGo := &MessageBody{native: ret.Pointer()}

	return retGo
}

type MessageClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value 'wrote_informational' : missing Type

	// UNSUPPORTED : C value 'wrote_headers' : missing Type

	// UNSUPPORTED : C value 'wrote_chunk' : missing Type

	// UNSUPPORTED : C value 'wrote_body' : missing Type

	// UNSUPPORTED : C value 'got_informational' : missing Type

	// UNSUPPORTED : C value 'got_headers' : missing Type

	// UNSUPPORTED : C value 'got_chunk' : missing Type

	// UNSUPPORTED : C value 'got_body' : missing Type

	// UNSUPPORTED : C value 'restarted' : missing Type

	// UNSUPPORTED : C value 'finished' : missing Type

	// UNSUPPORTED : C value 'starting' : missing Type

	// UNSUPPORTED : C value '_libsoup_reserved1' : missing Type

	// UNSUPPORTED : C value '_libsoup_reserved2' : missing Type

	// UNSUPPORTED : C value '_libsoup_reserved3' : missing Type

}

type MessageHeaders struct {
	native uintptr
}

// UNSUPPORTED : C value 'soup_message_headers_new' : parameter 'type' of type 'MessageHeadersType' not supported

type MessageHeadersIter struct {
	native uintptr
}

type MessageQueue struct {
	native uintptr
}

type MessageQueueItem struct {
	native uintptr
}

type Multipart struct {
	native uintptr
}

var newMultipartInvoker *gi.Function

// MultipartNew is a representation of the C type soup_multipart_new.
func MultipartNew(mimeType string) *Multipart {
	if newMultipartInvoker == nil {
		newMultipartInvoker = gi.StructFunctionInvokerNew("Soup", "Multipart", "new")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetString(mimeType)

	ret := newMultipartInvoker.Invoke(inArgs[:], nil)

	retGo := &Multipart{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'soup_multipart_new_from_message' : parameter 'headers' of type 'MessageHeaders' not supported

type MultipartInputStreamClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'Gio.FilterInputStreamClass'

}

type MultipartInputStreamPrivate struct {
	native uintptr
}

type PasswordManagerInterface struct {
	native uintptr
	// UNSUPPORTED : C value 'base' : no Go type for 'GObject.TypeInterface'

	// UNSUPPORTED : C value 'get_passwords_async' : missing Type

	// UNSUPPORTED : C value 'get_passwords_sync' : missing Type

}

type ProxyResolverDefaultClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

}

type ProxyResolverInterface struct {
	native uintptr
	// UNSUPPORTED : C value 'base' : no Go type for 'GObject.TypeInterface'

	// UNSUPPORTED : C value 'get_proxy_async' : missing Type

	// UNSUPPORTED : C value 'get_proxy_sync' : missing Type

}

type ProxyURIResolverInterface struct {
	native uintptr
	// UNSUPPORTED : C value 'base' : no Go type for 'GObject.TypeInterface'

	// UNSUPPORTED : C value 'get_proxy_uri_async' : missing Type

	// UNSUPPORTED : C value 'get_proxy_uri_sync' : missing Type

	// UNSUPPORTED : C value '_libsoup_reserved1' : missing Type

	// UNSUPPORTED : C value '_libsoup_reserved2' : missing Type

	// UNSUPPORTED : C value '_libsoup_reserved3' : missing Type

	// UNSUPPORTED : C value '_libsoup_reserved4' : missing Type

}

type Range struct {
	native uintptr
	Start  int64
	End    int64
}

type RequestClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent' : no Go type for 'GObject.ObjectClass'

	Schemes string
	// UNSUPPORTED : C value 'check_uri' : missing Type

	// UNSUPPORTED : C value 'send' : missing Type

	// UNSUPPORTED : C value 'send_async' : missing Type

	// UNSUPPORTED : C value 'send_finish' : missing Type

	// UNSUPPORTED : C value 'get_content_length' : missing Type

	// UNSUPPORTED : C value 'get_content_type' : missing Type

}

type RequestDataClass struct {
	native uintptr
	Parent *RequestClass
}

type RequestDataPrivate struct {
	native uintptr
}

type RequestFileClass struct {
	native uintptr
	Parent *RequestClass
}

type RequestFilePrivate struct {
	native uintptr
}

type RequestHTTPClass struct {
	native uintptr
	Parent *RequestClass
}

type RequestHTTPPrivate struct {
	native uintptr
}

type RequestPrivate struct {
	native uintptr
}

type RequesterClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

}

type RequesterPrivate struct {
	native uintptr
}

type ServerClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value 'request_started' : missing Type

	// UNSUPPORTED : C value 'request_read' : missing Type

	// UNSUPPORTED : C value 'request_finished' : missing Type

	// UNSUPPORTED : C value 'request_aborted' : missing Type

	// UNSUPPORTED : C value '_libsoup_reserved1' : missing Type

	// UNSUPPORTED : C value '_libsoup_reserved2' : missing Type

	// UNSUPPORTED : C value '_libsoup_reserved3' : missing Type

	// UNSUPPORTED : C value '_libsoup_reserved4' : missing Type

}

type SessionAsyncClass struct {
	native      uintptr
	ParentClass *SessionClass
	// UNSUPPORTED : C value '_libsoup_reserved1' : missing Type

	// UNSUPPORTED : C value '_libsoup_reserved2' : missing Type

	// UNSUPPORTED : C value '_libsoup_reserved3' : missing Type

	// UNSUPPORTED : C value '_libsoup_reserved4' : missing Type

}

type SessionClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value 'request_started' : missing Type

	// UNSUPPORTED : C value 'authenticate' : missing Type

	// UNSUPPORTED : C value 'queue_message' : missing Type

	// UNSUPPORTED : C value 'requeue_message' : missing Type

	// UNSUPPORTED : C value 'send_message' : missing Type

	// UNSUPPORTED : C value 'cancel_message' : missing Type

	// UNSUPPORTED : C value 'auth_required' : missing Type

	// UNSUPPORTED : C value 'flush_queue' : missing Type

	// UNSUPPORTED : C value 'kick' : missing Type

	// UNSUPPORTED : C value '_libsoup_reserved4' : missing Type

}

type SessionFeatureInterface struct {
	native uintptr
	// UNSUPPORTED : C value 'parent' : no Go type for 'GObject.TypeInterface'

	// UNSUPPORTED : C value 'attach' : missing Type

	// UNSUPPORTED : C value 'detach' : missing Type

	// UNSUPPORTED : C value 'request_queued' : missing Type

	// UNSUPPORTED : C value 'request_started' : missing Type

	// UNSUPPORTED : C value 'request_unqueued' : missing Type

	// UNSUPPORTED : C value 'add_feature' : missing Type

	// UNSUPPORTED : C value 'remove_feature' : missing Type

	// UNSUPPORTED : C value 'has_feature' : missing Type

}

type SessionSyncClass struct {
	native      uintptr
	ParentClass *SessionClass
	// UNSUPPORTED : C value '_libsoup_reserved1' : missing Type

	// UNSUPPORTED : C value '_libsoup_reserved2' : missing Type

	// UNSUPPORTED : C value '_libsoup_reserved3' : missing Type

	// UNSUPPORTED : C value '_libsoup_reserved4' : missing Type

}

type SocketClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value 'readable' : missing Type

	// UNSUPPORTED : C value 'writable' : missing Type

	// UNSUPPORTED : C value 'disconnected' : missing Type

	// UNSUPPORTED : C value 'new_connection' : missing Type

	// UNSUPPORTED : C value '_libsoup_reserved1' : missing Type

	// UNSUPPORTED : C value '_libsoup_reserved2' : missing Type

	// UNSUPPORTED : C value '_libsoup_reserved3' : missing Type

	// UNSUPPORTED : C value '_libsoup_reserved4' : missing Type

}

type URI struct {
	native   uintptr
	Scheme   string
	User     string
	Password string
	Host     string
	Port     uint32
	Path     string
	Query    string
	Fragment string
}

var newURIInvoker *gi.Function

// URINew is a representation of the C type soup_uri_new.
func URINew(uriString string) *URI {
	if newURIInvoker == nil {
		newURIInvoker = gi.StructFunctionInvokerNew("Soup", "URI", "new")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetString(uriString)

	ret := newURIInvoker.Invoke(inArgs[:], nil)

	retGo := &URI{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'soup_uri_new_with_base' : parameter 'base' of type 'URI' not supported

type WebsocketConnectionClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value 'message' : missing Type

	// UNSUPPORTED : C value 'error' : missing Type

	// UNSUPPORTED : C value 'closing' : missing Type

	// UNSUPPORTED : C value 'closed' : missing Type

	// UNSUPPORTED : C value 'pong' : missing Type

}

type WebsocketConnectionPrivate struct {
	native uintptr
}

type WebsocketExtensionClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	Name string
	// UNSUPPORTED : C value 'configure' : missing Type

	// UNSUPPORTED : C value 'get_request_params' : missing Type

	// UNSUPPORTED : C value 'get_response_params' : missing Type

	// UNSUPPORTED : C value 'process_outgoing_message' : missing Type

	// UNSUPPORTED : C value 'process_incoming_message' : missing Type

	// UNSUPPORTED : C value '_libsoup_reserved1' : missing Type

	// UNSUPPORTED : C value '_libsoup_reserved2' : missing Type

	// UNSUPPORTED : C value '_libsoup_reserved3' : missing Type

	// UNSUPPORTED : C value '_libsoup_reserved4' : missing Type

}

type WebsocketExtensionDeflateClass struct {
	native      uintptr
	ParentClass *WebsocketExtensionClass
}

type WebsocketExtensionManagerClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

}

type XMLRPCParams struct {
	native uintptr
}
