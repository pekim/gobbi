// This is a generated file - DO NOT EDIT
// +build gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import "C"

type ResolverRecordType C.GResolverRecordType

const (
	RESOLVER_RECORD_SRV ResolverRecordType = 1
	RESOLVER_RECORD_MX  ResolverRecordType = 2
	RESOLVER_RECORD_TXT ResolverRecordType = 3
	RESOLVER_RECORD_SOA ResolverRecordType = 4
	RESOLVER_RECORD_NS  ResolverRecordType = 5
)
