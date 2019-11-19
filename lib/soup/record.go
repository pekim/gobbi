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

var copyBufferInvoker *gi.Function

// Copy is a representation of the C type soup_buffer_copy.
func (recv *Buffer) Copy() *Buffer {
	if copyBufferInvoker == nil {
		copyBufferInvoker = gi.StructFunctionInvokerNew("Soup", "Buffer", "copy")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := copyBufferInvoker.Invoke(inArgs[:], nil)

	retGo := &Buffer{native: ret.Pointer()}

	return retGo
}

var freeBufferInvoker *gi.Function

// Free is a representation of the C type soup_buffer_free.
func (recv *Buffer) Free() {
	if freeBufferInvoker == nil {
		freeBufferInvoker = gi.StructFunctionInvokerNew("Soup", "Buffer", "free")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	freeBufferInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'soup_buffer_get_as_bytes' : return type 'GLib.Bytes' not supported

// UNSUPPORTED : C value 'soup_buffer_get_data' : parameter 'data' has no type

// UNSUPPORTED : C value 'soup_buffer_get_owner' : return type 'gpointer' not supported

// UNSUPPORTED : C value 'soup_buffer_new_subbuffer' : parameter 'offset' of type 'gsize' not supported

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

// UNSUPPORTED : C value 'soup_client_context_get_address' : return type 'Address' not supported

// UNSUPPORTED : C value 'soup_client_context_get_auth_domain' : return type 'AuthDomain' not supported

var getAuthUserClientContextInvoker *gi.Function

// GetAuthUser is a representation of the C type soup_client_context_get_auth_user.
func (recv *ClientContext) GetAuthUser() string {
	if getAuthUserClientContextInvoker == nil {
		getAuthUserClientContextInvoker = gi.StructFunctionInvokerNew("Soup", "ClientContext", "get_auth_user")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getAuthUserClientContextInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'soup_client_context_get_gsocket' : return type 'Gio.Socket' not supported

var getHostClientContextInvoker *gi.Function

// GetHost is a representation of the C type soup_client_context_get_host.
func (recv *ClientContext) GetHost() string {
	if getHostClientContextInvoker == nil {
		getHostClientContextInvoker = gi.StructFunctionInvokerNew("Soup", "ClientContext", "get_host")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getHostClientContextInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'soup_client_context_get_local_address' : return type 'Gio.SocketAddress' not supported

// UNSUPPORTED : C value 'soup_client_context_get_remote_address' : return type 'Gio.SocketAddress' not supported

// UNSUPPORTED : C value 'soup_client_context_get_socket' : return type 'Socket' not supported

// UNSUPPORTED : C value 'soup_client_context_steal_connection' : return type 'Gio.IOStream' not supported

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

// UNSUPPORTED : C value 'soup_cookie_applies_to_uri' : parameter 'uri' of type 'URI' not supported

var copyCookieInvoker *gi.Function

// Copy is a representation of the C type soup_cookie_copy.
func (recv *Cookie) Copy() *Cookie {
	if copyCookieInvoker == nil {
		copyCookieInvoker = gi.StructFunctionInvokerNew("Soup", "Cookie", "copy")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := copyCookieInvoker.Invoke(inArgs[:], nil)

	retGo := &Cookie{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'soup_cookie_domain_matches' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'soup_cookie_equal' : parameter 'cookie2' of type 'Cookie' not supported

var freeCookieInvoker *gi.Function

// Free is a representation of the C type soup_cookie_free.
func (recv *Cookie) Free() {
	if freeCookieInvoker == nil {
		freeCookieInvoker = gi.StructFunctionInvokerNew("Soup", "Cookie", "free")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	freeCookieInvoker.Invoke(inArgs[:], nil)

}

var getDomainCookieInvoker *gi.Function

// GetDomain is a representation of the C type soup_cookie_get_domain.
func (recv *Cookie) GetDomain() string {
	if getDomainCookieInvoker == nil {
		getDomainCookieInvoker = gi.StructFunctionInvokerNew("Soup", "Cookie", "get_domain")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getDomainCookieInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

var getExpiresCookieInvoker *gi.Function

// GetExpires is a representation of the C type soup_cookie_get_expires.
func (recv *Cookie) GetExpires() *Date {
	if getExpiresCookieInvoker == nil {
		getExpiresCookieInvoker = gi.StructFunctionInvokerNew("Soup", "Cookie", "get_expires")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getExpiresCookieInvoker.Invoke(inArgs[:], nil)

	retGo := &Date{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'soup_cookie_get_http_only' : return type 'gboolean' not supported

var getNameCookieInvoker *gi.Function

// GetName is a representation of the C type soup_cookie_get_name.
func (recv *Cookie) GetName() string {
	if getNameCookieInvoker == nil {
		getNameCookieInvoker = gi.StructFunctionInvokerNew("Soup", "Cookie", "get_name")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getNameCookieInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

var getPathCookieInvoker *gi.Function

// GetPath is a representation of the C type soup_cookie_get_path.
func (recv *Cookie) GetPath() string {
	if getPathCookieInvoker == nil {
		getPathCookieInvoker = gi.StructFunctionInvokerNew("Soup", "Cookie", "get_path")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getPathCookieInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'soup_cookie_get_secure' : return type 'gboolean' not supported

var getValueCookieInvoker *gi.Function

// GetValue is a representation of the C type soup_cookie_get_value.
func (recv *Cookie) GetValue() string {
	if getValueCookieInvoker == nil {
		getValueCookieInvoker = gi.StructFunctionInvokerNew("Soup", "Cookie", "get_value")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getValueCookieInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

var setDomainCookieInvoker *gi.Function

// SetDomain is a representation of the C type soup_cookie_set_domain.
func (recv *Cookie) SetDomain(domain string) {
	if setDomainCookieInvoker == nil {
		setDomainCookieInvoker = gi.StructFunctionInvokerNew("Soup", "Cookie", "set_domain")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(domain)

	setDomainCookieInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'soup_cookie_set_expires' : parameter 'expires' of type 'Date' not supported

// UNSUPPORTED : C value 'soup_cookie_set_http_only' : parameter 'http_only' of type 'gboolean' not supported

var setMaxAgeCookieInvoker *gi.Function

// SetMaxAge is a representation of the C type soup_cookie_set_max_age.
func (recv *Cookie) SetMaxAge(maxAge int32) {
	if setMaxAgeCookieInvoker == nil {
		setMaxAgeCookieInvoker = gi.StructFunctionInvokerNew("Soup", "Cookie", "set_max_age")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(maxAge)

	setMaxAgeCookieInvoker.Invoke(inArgs[:], nil)

}

var setNameCookieInvoker *gi.Function

// SetName is a representation of the C type soup_cookie_set_name.
func (recv *Cookie) SetName(name string) {
	if setNameCookieInvoker == nil {
		setNameCookieInvoker = gi.StructFunctionInvokerNew("Soup", "Cookie", "set_name")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(name)

	setNameCookieInvoker.Invoke(inArgs[:], nil)

}

var setPathCookieInvoker *gi.Function

// SetPath is a representation of the C type soup_cookie_set_path.
func (recv *Cookie) SetPath(path string) {
	if setPathCookieInvoker == nil {
		setPathCookieInvoker = gi.StructFunctionInvokerNew("Soup", "Cookie", "set_path")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(path)

	setPathCookieInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'soup_cookie_set_secure' : parameter 'secure' of type 'gboolean' not supported

var setValueCookieInvoker *gi.Function

// SetValue is a representation of the C type soup_cookie_set_value.
func (recv *Cookie) SetValue(value string) {
	if setValueCookieInvoker == nil {
		setValueCookieInvoker = gi.StructFunctionInvokerNew("Soup", "Cookie", "set_value")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(value)

	setValueCookieInvoker.Invoke(inArgs[:], nil)

}

var toCookieHeaderCookieInvoker *gi.Function

// ToCookieHeader is a representation of the C type soup_cookie_to_cookie_header.
func (recv *Cookie) ToCookieHeader() string {
	if toCookieHeaderCookieInvoker == nil {
		toCookieHeaderCookieInvoker = gi.StructFunctionInvokerNew("Soup", "Cookie", "to_cookie_header")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := toCookieHeaderCookieInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

var toSetCookieHeaderCookieInvoker *gi.Function

// ToSetCookieHeader is a representation of the C type soup_cookie_to_set_cookie_header.
func (recv *Cookie) ToSetCookieHeader() string {
	if toSetCookieHeaderCookieInvoker == nil {
		toSetCookieHeaderCookieInvoker = gi.StructFunctionInvokerNew("Soup", "Cookie", "to_set_cookie_header")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := toSetCookieHeaderCookieInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

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

var copyDateInvoker *gi.Function

// Copy is a representation of the C type soup_date_copy.
func (recv *Date) Copy() *Date {
	if copyDateInvoker == nil {
		copyDateInvoker = gi.StructFunctionInvokerNew("Soup", "Date", "copy")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := copyDateInvoker.Invoke(inArgs[:], nil)

	retGo := &Date{native: ret.Pointer()}

	return retGo
}

var freeDateInvoker *gi.Function

// Free is a representation of the C type soup_date_free.
func (recv *Date) Free() {
	if freeDateInvoker == nil {
		freeDateInvoker = gi.StructFunctionInvokerNew("Soup", "Date", "free")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	freeDateInvoker.Invoke(inArgs[:], nil)

}

var getDayDateInvoker *gi.Function

// GetDay is a representation of the C type soup_date_get_day.
func (recv *Date) GetDay() int32 {
	if getDayDateInvoker == nil {
		getDayDateInvoker = gi.StructFunctionInvokerNew("Soup", "Date", "get_day")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getDayDateInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var getHourDateInvoker *gi.Function

// GetHour is a representation of the C type soup_date_get_hour.
func (recv *Date) GetHour() int32 {
	if getHourDateInvoker == nil {
		getHourDateInvoker = gi.StructFunctionInvokerNew("Soup", "Date", "get_hour")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getHourDateInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var getMinuteDateInvoker *gi.Function

// GetMinute is a representation of the C type soup_date_get_minute.
func (recv *Date) GetMinute() int32 {
	if getMinuteDateInvoker == nil {
		getMinuteDateInvoker = gi.StructFunctionInvokerNew("Soup", "Date", "get_minute")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getMinuteDateInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var getMonthDateInvoker *gi.Function

// GetMonth is a representation of the C type soup_date_get_month.
func (recv *Date) GetMonth() int32 {
	if getMonthDateInvoker == nil {
		getMonthDateInvoker = gi.StructFunctionInvokerNew("Soup", "Date", "get_month")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getMonthDateInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var getOffsetDateInvoker *gi.Function

// GetOffset is a representation of the C type soup_date_get_offset.
func (recv *Date) GetOffset() int32 {
	if getOffsetDateInvoker == nil {
		getOffsetDateInvoker = gi.StructFunctionInvokerNew("Soup", "Date", "get_offset")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getOffsetDateInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var getSecondDateInvoker *gi.Function

// GetSecond is a representation of the C type soup_date_get_second.
func (recv *Date) GetSecond() int32 {
	if getSecondDateInvoker == nil {
		getSecondDateInvoker = gi.StructFunctionInvokerNew("Soup", "Date", "get_second")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getSecondDateInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var getUtcDateInvoker *gi.Function

// GetUtc is a representation of the C type soup_date_get_utc.
func (recv *Date) GetUtc() int32 {
	if getUtcDateInvoker == nil {
		getUtcDateInvoker = gi.StructFunctionInvokerNew("Soup", "Date", "get_utc")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getUtcDateInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var getYearDateInvoker *gi.Function

// GetYear is a representation of the C type soup_date_get_year.
func (recv *Date) GetYear() int32 {
	if getYearDateInvoker == nil {
		getYearDateInvoker = gi.StructFunctionInvokerNew("Soup", "Date", "get_year")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getYearDateInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

// UNSUPPORTED : C value 'soup_date_is_past' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'soup_date_to_string' : parameter 'format' of type 'DateFormat' not supported

var toTimeTDateInvoker *gi.Function

// ToTimeT is a representation of the C type soup_date_to_time_t.
func (recv *Date) ToTimeT() int64 {
	if toTimeTDateInvoker == nil {
		toTimeTDateInvoker = gi.StructFunctionInvokerNew("Soup", "Date", "to_time_t")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := toTimeTDateInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int64()

	return retGo
}

// UNSUPPORTED : C value 'soup_date_to_timeval' : parameter 'time' of type 'GLib.TimeVal' not supported

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

var copyHSTSPolicyInvoker *gi.Function

// Copy is a representation of the C type soup_hsts_policy_copy.
func (recv *HSTSPolicy) Copy() *HSTSPolicy {
	if copyHSTSPolicyInvoker == nil {
		copyHSTSPolicyInvoker = gi.StructFunctionInvokerNew("Soup", "HSTSPolicy", "copy")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := copyHSTSPolicyInvoker.Invoke(inArgs[:], nil)

	retGo := &HSTSPolicy{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'soup_hsts_policy_equal' : parameter 'policy2' of type 'HSTSPolicy' not supported

var freeHSTSPolicyInvoker *gi.Function

// Free is a representation of the C type soup_hsts_policy_free.
func (recv *HSTSPolicy) Free() {
	if freeHSTSPolicyInvoker == nil {
		freeHSTSPolicyInvoker = gi.StructFunctionInvokerNew("Soup", "HSTSPolicy", "free")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	freeHSTSPolicyInvoker.Invoke(inArgs[:], nil)

}

var getDomainHSTSPolicyInvoker *gi.Function

// GetDomain is a representation of the C type soup_hsts_policy_get_domain.
func (recv *HSTSPolicy) GetDomain() string {
	if getDomainHSTSPolicyInvoker == nil {
		getDomainHSTSPolicyInvoker = gi.StructFunctionInvokerNew("Soup", "HSTSPolicy", "get_domain")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getDomainHSTSPolicyInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'soup_hsts_policy_includes_subdomains' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'soup_hsts_policy_is_expired' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'soup_hsts_policy_is_session_policy' : return type 'gboolean' not supported

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

// UNSUPPORTED : C value 'soup_message_body_append' : parameter 'use' of type 'MemoryUse' not supported

// UNSUPPORTED : C value 'soup_message_body_append_buffer' : parameter 'buffer' of type 'Buffer' not supported

// UNSUPPORTED : C value 'soup_message_body_append_take' : parameter 'data' has no type

var completeMessageBodyInvoker *gi.Function

// Complete is a representation of the C type soup_message_body_complete.
func (recv *MessageBody) Complete() {
	if completeMessageBodyInvoker == nil {
		completeMessageBodyInvoker = gi.StructFunctionInvokerNew("Soup", "MessageBody", "complete")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	completeMessageBodyInvoker.Invoke(inArgs[:], nil)

}

var flattenMessageBodyInvoker *gi.Function

// Flatten is a representation of the C type soup_message_body_flatten.
func (recv *MessageBody) Flatten() *Buffer {
	if flattenMessageBodyInvoker == nil {
		flattenMessageBodyInvoker = gi.StructFunctionInvokerNew("Soup", "MessageBody", "flatten")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := flattenMessageBodyInvoker.Invoke(inArgs[:], nil)

	retGo := &Buffer{native: ret.Pointer()}

	return retGo
}

var freeMessageBodyInvoker *gi.Function

// Free is a representation of the C type soup_message_body_free.
func (recv *MessageBody) Free() {
	if freeMessageBodyInvoker == nil {
		freeMessageBodyInvoker = gi.StructFunctionInvokerNew("Soup", "MessageBody", "free")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	freeMessageBodyInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'soup_message_body_get_accumulate' : return type 'gboolean' not supported

var getChunkMessageBodyInvoker *gi.Function

// GetChunk is a representation of the C type soup_message_body_get_chunk.
func (recv *MessageBody) GetChunk(offset int64) *Buffer {
	if getChunkMessageBodyInvoker == nil {
		getChunkMessageBodyInvoker = gi.StructFunctionInvokerNew("Soup", "MessageBody", "get_chunk")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt64(offset)

	ret := getChunkMessageBodyInvoker.Invoke(inArgs[:], nil)

	retGo := &Buffer{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'soup_message_body_got_chunk' : parameter 'chunk' of type 'Buffer' not supported

// UNSUPPORTED : C value 'soup_message_body_set_accumulate' : parameter 'accumulate' of type 'gboolean' not supported

var truncateMessageBodyInvoker *gi.Function

// Truncate is a representation of the C type soup_message_body_truncate.
func (recv *MessageBody) Truncate() {
	if truncateMessageBodyInvoker == nil {
		truncateMessageBodyInvoker = gi.StructFunctionInvokerNew("Soup", "MessageBody", "truncate")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	truncateMessageBodyInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'soup_message_body_wrote_chunk' : parameter 'chunk' of type 'Buffer' not supported

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

var appendMessageHeadersInvoker *gi.Function

// Append is a representation of the C type soup_message_headers_append.
func (recv *MessageHeaders) Append(name string, value string) {
	if appendMessageHeadersInvoker == nil {
		appendMessageHeadersInvoker = gi.StructFunctionInvokerNew("Soup", "MessageHeaders", "append")
	}

	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(name)
	inArgs[2].SetString(value)

	appendMessageHeadersInvoker.Invoke(inArgs[:], nil)

}

var cleanConnectionHeadersMessageHeadersInvoker *gi.Function

// CleanConnectionHeaders is a representation of the C type soup_message_headers_clean_connection_headers.
func (recv *MessageHeaders) CleanConnectionHeaders() {
	if cleanConnectionHeadersMessageHeadersInvoker == nil {
		cleanConnectionHeadersMessageHeadersInvoker = gi.StructFunctionInvokerNew("Soup", "MessageHeaders", "clean_connection_headers")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	cleanConnectionHeadersMessageHeadersInvoker.Invoke(inArgs[:], nil)

}

var clearMessageHeadersInvoker *gi.Function

// Clear is a representation of the C type soup_message_headers_clear.
func (recv *MessageHeaders) Clear() {
	if clearMessageHeadersInvoker == nil {
		clearMessageHeadersInvoker = gi.StructFunctionInvokerNew("Soup", "MessageHeaders", "clear")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	clearMessageHeadersInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'soup_message_headers_foreach' : parameter 'func' of type 'MessageHeadersForeachFunc' not supported

var freeMessageHeadersInvoker *gi.Function

// Free is a representation of the C type soup_message_headers_free.
func (recv *MessageHeaders) Free() {
	if freeMessageHeadersInvoker == nil {
		freeMessageHeadersInvoker = gi.StructFunctionInvokerNew("Soup", "MessageHeaders", "free")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	freeMessageHeadersInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'soup_message_headers_free_ranges' : parameter 'ranges' of type 'Range' not supported

var getMessageHeadersInvoker *gi.Function

// Get is a representation of the C type soup_message_headers_get.
func (recv *MessageHeaders) Get(name string) string {
	if getMessageHeadersInvoker == nil {
		getMessageHeadersInvoker = gi.StructFunctionInvokerNew("Soup", "MessageHeaders", "get")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(name)

	ret := getMessageHeadersInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'soup_message_headers_get_content_disposition' : parameter 'params' of type 'GLib.HashTable' not supported

var getContentLengthMessageHeadersInvoker *gi.Function

// GetContentLength is a representation of the C type soup_message_headers_get_content_length.
func (recv *MessageHeaders) GetContentLength() int64 {
	if getContentLengthMessageHeadersInvoker == nil {
		getContentLengthMessageHeadersInvoker = gi.StructFunctionInvokerNew("Soup", "MessageHeaders", "get_content_length")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getContentLengthMessageHeadersInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int64()

	return retGo
}

// UNSUPPORTED : C value 'soup_message_headers_get_content_range' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'soup_message_headers_get_content_type' : parameter 'params' of type 'GLib.HashTable' not supported

// UNSUPPORTED : C value 'soup_message_headers_get_encoding' : return type 'Encoding' not supported

// UNSUPPORTED : C value 'soup_message_headers_get_expectations' : return type 'Expectation' not supported

// UNSUPPORTED : C value 'soup_message_headers_get_headers_type' : return type 'MessageHeadersType' not supported

var getListMessageHeadersInvoker *gi.Function

// GetList is a representation of the C type soup_message_headers_get_list.
func (recv *MessageHeaders) GetList(name string) string {
	if getListMessageHeadersInvoker == nil {
		getListMessageHeadersInvoker = gi.StructFunctionInvokerNew("Soup", "MessageHeaders", "get_list")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(name)

	ret := getListMessageHeadersInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

var getOneMessageHeadersInvoker *gi.Function

// GetOne is a representation of the C type soup_message_headers_get_one.
func (recv *MessageHeaders) GetOne(name string) string {
	if getOneMessageHeadersInvoker == nil {
		getOneMessageHeadersInvoker = gi.StructFunctionInvokerNew("Soup", "MessageHeaders", "get_one")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(name)

	ret := getOneMessageHeadersInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'soup_message_headers_get_ranges' : parameter 'ranges' has no type

// UNSUPPORTED : C value 'soup_message_headers_header_contains' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'soup_message_headers_header_equals' : return type 'gboolean' not supported

var removeMessageHeadersInvoker *gi.Function

// Remove is a representation of the C type soup_message_headers_remove.
func (recv *MessageHeaders) Remove(name string) {
	if removeMessageHeadersInvoker == nil {
		removeMessageHeadersInvoker = gi.StructFunctionInvokerNew("Soup", "MessageHeaders", "remove")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(name)

	removeMessageHeadersInvoker.Invoke(inArgs[:], nil)

}

var replaceMessageHeadersInvoker *gi.Function

// Replace is a representation of the C type soup_message_headers_replace.
func (recv *MessageHeaders) Replace(name string, value string) {
	if replaceMessageHeadersInvoker == nil {
		replaceMessageHeadersInvoker = gi.StructFunctionInvokerNew("Soup", "MessageHeaders", "replace")
	}

	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(name)
	inArgs[2].SetString(value)

	replaceMessageHeadersInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'soup_message_headers_set_content_disposition' : parameter 'params' of type 'GLib.HashTable' not supported

var setContentLengthMessageHeadersInvoker *gi.Function

// SetContentLength is a representation of the C type soup_message_headers_set_content_length.
func (recv *MessageHeaders) SetContentLength(contentLength int64) {
	if setContentLengthMessageHeadersInvoker == nil {
		setContentLengthMessageHeadersInvoker = gi.StructFunctionInvokerNew("Soup", "MessageHeaders", "set_content_length")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt64(contentLength)

	setContentLengthMessageHeadersInvoker.Invoke(inArgs[:], nil)

}

var setContentRangeMessageHeadersInvoker *gi.Function

// SetContentRange is a representation of the C type soup_message_headers_set_content_range.
func (recv *MessageHeaders) SetContentRange(start int64, end int64, totalLength int64) {
	if setContentRangeMessageHeadersInvoker == nil {
		setContentRangeMessageHeadersInvoker = gi.StructFunctionInvokerNew("Soup", "MessageHeaders", "set_content_range")
	}

	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt64(start)
	inArgs[2].SetInt64(end)
	inArgs[3].SetInt64(totalLength)

	setContentRangeMessageHeadersInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'soup_message_headers_set_content_type' : parameter 'params' of type 'GLib.HashTable' not supported

// UNSUPPORTED : C value 'soup_message_headers_set_encoding' : parameter 'encoding' of type 'Encoding' not supported

// UNSUPPORTED : C value 'soup_message_headers_set_expectations' : parameter 'expectations' of type 'Expectation' not supported

var setRangeMessageHeadersInvoker *gi.Function

// SetRange is a representation of the C type soup_message_headers_set_range.
func (recv *MessageHeaders) SetRange(start int64, end int64) {
	if setRangeMessageHeadersInvoker == nil {
		setRangeMessageHeadersInvoker = gi.StructFunctionInvokerNew("Soup", "MessageHeaders", "set_range")
	}

	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt64(start)
	inArgs[2].SetInt64(end)

	setRangeMessageHeadersInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'soup_message_headers_set_ranges' : parameter 'ranges' of type 'Range' not supported

type MessageHeadersIter struct {
	native uintptr
}

// UNSUPPORTED : C value 'soup_message_headers_iter_next' : return type 'gboolean' not supported

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

// UNSUPPORTED : C value 'soup_multipart_append_form_file' : parameter 'body' of type 'Buffer' not supported

var appendFormStringMultipartInvoker *gi.Function

// AppendFormString is a representation of the C type soup_multipart_append_form_string.
func (recv *Multipart) AppendFormString(controlName string, data string) {
	if appendFormStringMultipartInvoker == nil {
		appendFormStringMultipartInvoker = gi.StructFunctionInvokerNew("Soup", "Multipart", "append_form_string")
	}

	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(controlName)
	inArgs[2].SetString(data)

	appendFormStringMultipartInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'soup_multipart_append_part' : parameter 'headers' of type 'MessageHeaders' not supported

var freeMultipartInvoker *gi.Function

// Free is a representation of the C type soup_multipart_free.
func (recv *Multipart) Free() {
	if freeMultipartInvoker == nil {
		freeMultipartInvoker = gi.StructFunctionInvokerNew("Soup", "Multipart", "free")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	freeMultipartInvoker.Invoke(inArgs[:], nil)

}

var getLengthMultipartInvoker *gi.Function

// GetLength is a representation of the C type soup_multipart_get_length.
func (recv *Multipart) GetLength() int32 {
	if getLengthMultipartInvoker == nil {
		getLengthMultipartInvoker = gi.StructFunctionInvokerNew("Soup", "Multipart", "get_length")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getLengthMultipartInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

// UNSUPPORTED : C value 'soup_multipart_get_part' : parameter 'headers' of type 'MessageHeaders' not supported

// UNSUPPORTED : C value 'soup_multipart_to_message' : parameter 'dest_headers' of type 'MessageHeaders' not supported

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

var copyURIInvoker *gi.Function

// Copy is a representation of the C type soup_uri_copy.
func (recv *URI) Copy() *URI {
	if copyURIInvoker == nil {
		copyURIInvoker = gi.StructFunctionInvokerNew("Soup", "URI", "copy")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := copyURIInvoker.Invoke(inArgs[:], nil)

	retGo := &URI{native: ret.Pointer()}

	return retGo
}

var copyHostURIInvoker *gi.Function

// CopyHost is a representation of the C type soup_uri_copy_host.
func (recv *URI) CopyHost() *URI {
	if copyHostURIInvoker == nil {
		copyHostURIInvoker = gi.StructFunctionInvokerNew("Soup", "URI", "copy_host")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := copyHostURIInvoker.Invoke(inArgs[:], nil)

	retGo := &URI{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'soup_uri_equal' : parameter 'uri2' of type 'URI' not supported

var freeURIInvoker *gi.Function

// Free is a representation of the C type soup_uri_free.
func (recv *URI) Free() {
	if freeURIInvoker == nil {
		freeURIInvoker = gi.StructFunctionInvokerNew("Soup", "URI", "free")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	freeURIInvoker.Invoke(inArgs[:], nil)

}

var getFragmentURIInvoker *gi.Function

// GetFragment is a representation of the C type soup_uri_get_fragment.
func (recv *URI) GetFragment() string {
	if getFragmentURIInvoker == nil {
		getFragmentURIInvoker = gi.StructFunctionInvokerNew("Soup", "URI", "get_fragment")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getFragmentURIInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

var getHostURIInvoker *gi.Function

// GetHost is a representation of the C type soup_uri_get_host.
func (recv *URI) GetHost() string {
	if getHostURIInvoker == nil {
		getHostURIInvoker = gi.StructFunctionInvokerNew("Soup", "URI", "get_host")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getHostURIInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

var getPasswordURIInvoker *gi.Function

// GetPassword is a representation of the C type soup_uri_get_password.
func (recv *URI) GetPassword() string {
	if getPasswordURIInvoker == nil {
		getPasswordURIInvoker = gi.StructFunctionInvokerNew("Soup", "URI", "get_password")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getPasswordURIInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

var getPathURIInvoker *gi.Function

// GetPath is a representation of the C type soup_uri_get_path.
func (recv *URI) GetPath() string {
	if getPathURIInvoker == nil {
		getPathURIInvoker = gi.StructFunctionInvokerNew("Soup", "URI", "get_path")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getPathURIInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

var getPortURIInvoker *gi.Function

// GetPort is a representation of the C type soup_uri_get_port.
func (recv *URI) GetPort() uint32 {
	if getPortURIInvoker == nil {
		getPortURIInvoker = gi.StructFunctionInvokerNew("Soup", "URI", "get_port")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getPortURIInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Uint32()

	return retGo
}

var getQueryURIInvoker *gi.Function

// GetQuery is a representation of the C type soup_uri_get_query.
func (recv *URI) GetQuery() string {
	if getQueryURIInvoker == nil {
		getQueryURIInvoker = gi.StructFunctionInvokerNew("Soup", "URI", "get_query")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getQueryURIInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

var getSchemeURIInvoker *gi.Function

// GetScheme is a representation of the C type soup_uri_get_scheme.
func (recv *URI) GetScheme() string {
	if getSchemeURIInvoker == nil {
		getSchemeURIInvoker = gi.StructFunctionInvokerNew("Soup", "URI", "get_scheme")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getSchemeURIInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

var getUserURIInvoker *gi.Function

// GetUser is a representation of the C type soup_uri_get_user.
func (recv *URI) GetUser() string {
	if getUserURIInvoker == nil {
		getUserURIInvoker = gi.StructFunctionInvokerNew("Soup", "URI", "get_user")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getUserURIInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'soup_uri_host_equal' : parameter 'v2' of type 'URI' not supported

var hostHashURIInvoker *gi.Function

// HostHash is a representation of the C type soup_uri_host_hash.
func (recv *URI) HostHash() uint32 {
	if hostHashURIInvoker == nil {
		hostHashURIInvoker = gi.StructFunctionInvokerNew("Soup", "URI", "host_hash")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := hostHashURIInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Uint32()

	return retGo
}

var setFragmentURIInvoker *gi.Function

// SetFragment is a representation of the C type soup_uri_set_fragment.
func (recv *URI) SetFragment(fragment string) {
	if setFragmentURIInvoker == nil {
		setFragmentURIInvoker = gi.StructFunctionInvokerNew("Soup", "URI", "set_fragment")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(fragment)

	setFragmentURIInvoker.Invoke(inArgs[:], nil)

}

var setHostURIInvoker *gi.Function

// SetHost is a representation of the C type soup_uri_set_host.
func (recv *URI) SetHost(host string) {
	if setHostURIInvoker == nil {
		setHostURIInvoker = gi.StructFunctionInvokerNew("Soup", "URI", "set_host")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(host)

	setHostURIInvoker.Invoke(inArgs[:], nil)

}

var setPasswordURIInvoker *gi.Function

// SetPassword is a representation of the C type soup_uri_set_password.
func (recv *URI) SetPassword(password string) {
	if setPasswordURIInvoker == nil {
		setPasswordURIInvoker = gi.StructFunctionInvokerNew("Soup", "URI", "set_password")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(password)

	setPasswordURIInvoker.Invoke(inArgs[:], nil)

}

var setPathURIInvoker *gi.Function

// SetPath is a representation of the C type soup_uri_set_path.
func (recv *URI) SetPath(path string) {
	if setPathURIInvoker == nil {
		setPathURIInvoker = gi.StructFunctionInvokerNew("Soup", "URI", "set_path")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(path)

	setPathURIInvoker.Invoke(inArgs[:], nil)

}

var setPortURIInvoker *gi.Function

// SetPort is a representation of the C type soup_uri_set_port.
func (recv *URI) SetPort(port uint32) {
	if setPortURIInvoker == nil {
		setPortURIInvoker = gi.StructFunctionInvokerNew("Soup", "URI", "set_port")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetUint32(port)

	setPortURIInvoker.Invoke(inArgs[:], nil)

}

var setQueryURIInvoker *gi.Function

// SetQuery is a representation of the C type soup_uri_set_query.
func (recv *URI) SetQuery(query string) {
	if setQueryURIInvoker == nil {
		setQueryURIInvoker = gi.StructFunctionInvokerNew("Soup", "URI", "set_query")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(query)

	setQueryURIInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'soup_uri_set_query_from_fields' : parameter '...' has no type

// UNSUPPORTED : C value 'soup_uri_set_query_from_form' : parameter 'form' of type 'GLib.HashTable' not supported

var setSchemeURIInvoker *gi.Function

// SetScheme is a representation of the C type soup_uri_set_scheme.
func (recv *URI) SetScheme(scheme string) {
	if setSchemeURIInvoker == nil {
		setSchemeURIInvoker = gi.StructFunctionInvokerNew("Soup", "URI", "set_scheme")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(scheme)

	setSchemeURIInvoker.Invoke(inArgs[:], nil)

}

var setUserURIInvoker *gi.Function

// SetUser is a representation of the C type soup_uri_set_user.
func (recv *URI) SetUser(user string) {
	if setUserURIInvoker == nil {
		setUserURIInvoker = gi.StructFunctionInvokerNew("Soup", "URI", "set_user")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(user)

	setUserURIInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'soup_uri_to_string' : parameter 'just_path_and_query' of type 'gboolean' not supported

// UNSUPPORTED : C value 'soup_uri_uses_default_port' : return type 'gboolean' not supported

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

var freeXMLRPCParamsInvoker *gi.Function

// Free is a representation of the C type soup_xmlrpc_params_free.
func (recv *XMLRPCParams) Free() {
	if freeXMLRPCParamsInvoker == nil {
		freeXMLRPCParamsInvoker = gi.StructFunctionInvokerNew("Soup", "XMLRPCParams", "free")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	freeXMLRPCParamsInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'soup_xmlrpc_params_parse' : return type 'GLib.Variant' not supported
