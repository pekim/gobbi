// This is a generated file - DO NOT EDIT
// +build gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gio/gdesktopappinfo.h>
// #include <gio/gfiledescriptorbased.h>
// #include <gio/gio.h>
// #include <gio/gunixconnection.h>
// #include <gio/gunixcredentialsmessage.h>
// #include <gio/gunixfdlist.h>
// #include <gio/gunixfdmessage.h>
// #include <gio/gunixinputstream.h>
// #include <gio/gunixmounts.h>
// #include <gio/gunixoutputstream.h>
// #include <gio/gunixsocketaddress.h>
// #include <stdlib.h>
import "C"

/*
The type of record that g_resolver_lookup_records() or
g_resolver_lookup_records_async() should retrieve. The records are returned
as lists of #GVariant tuples. Each record type has different values in
the variant tuples returned.

%G_RESOLVER_RECORD_SRV records are returned as variants with the signature
'(qqqs)', containing a guint16 with the priority, a guint16 with the
weight, a guint16 with the port, and a string of the hostname.

%G_RESOLVER_RECORD_MX records are returned as variants with the signature
'(qs)', representing a guint16 with the preference, and a string containing
the mail exchanger hostname.

%G_RESOLVER_RECORD_TXT records are returned as variants with the signature
'(as)', representing an array of the strings in the text record.

%G_RESOLVER_RECORD_SOA records are returned as variants with the signature
'(ssuuuuu)', representing a string containing the primary name server, a
string containing the administrator, the serial as a guint32, the refresh
interval as guint32, the retry interval as a guint32, the expire timeout
as a guint32, and the ttl as a guint32.

%G_RESOLVER_RECORD_NS records are returned as variants with the signature
'(s)', representing a string of the hostname of the name server.
*/
type ResolverRecordType C.GResolverRecordType

const (
	// lookup DNS SRV records for a domain
	RESOLVER_RECORD_SRV ResolverRecordType = 1
	// lookup DNS MX records for a domain
	RESOLVER_RECORD_MX ResolverRecordType = 2
	// lookup DNS TXT records for a name
	RESOLVER_RECORD_TXT ResolverRecordType = 3
	// lookup DNS SOA records for a zone
	RESOLVER_RECORD_SOA ResolverRecordType = 4
	// lookup DNS NS records for a domain
	RESOLVER_RECORD_NS ResolverRecordType = 5
)
