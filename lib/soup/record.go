// Code generated - DO NOT EDIT.

package soup

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
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'AuthDomainClass'

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
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'AuthDomainClass'

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
	// UNSUPPORTED : C value 'data' : missing Type.Name

	Length uintptr
}

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
	native uintptr
	Name   string
	Value  string
	Domain string
	Path   string
	// UNSUPPORTED : C value 'expires' : no Go type for 'Date'

	// UNSUPPORTED : C value 'secure' : no Go type for 'gboolean'

	// UNSUPPORTED : C value 'http_only' : no Go type for 'gboolean'

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
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'CookieJarClass'

	// UNSUPPORTED : C value '_libsoup_reserved1' : missing Type

	// UNSUPPORTED : C value '_libsoup_reserved2' : missing Type

	// UNSUPPORTED : C value '_libsoup_reserved3' : missing Type

	// UNSUPPORTED : C value '_libsoup_reserved4' : missing Type

}

type CookieJarTextClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'CookieJarClass'

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
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'HSTSEnforcerClass'

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
	native uintptr
	Domain string
	MaxAge uint64
	// UNSUPPORTED : C value 'expires' : no Go type for 'Date'

	// UNSUPPORTED : C value 'include_subdomains' : no Go type for 'gboolean'

}

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
	// UNSUPPORTED : C value 'parent' : no Go type for 'RequestClass'

}

type RequestDataPrivate struct {
	native uintptr
}

type RequestFileClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent' : no Go type for 'RequestClass'

}

type RequestFilePrivate struct {
	native uintptr
}

type RequestHTTPClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent' : no Go type for 'RequestClass'

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
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'SessionClass'

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
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'SessionClass'

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
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'WebsocketExtensionClass'

}

type WebsocketExtensionManagerClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

}

type XMLRPCParams struct {
	native uintptr
}
