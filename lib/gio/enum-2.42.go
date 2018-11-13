// This is a generated file - DO NOT EDIT
// +build gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

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

// Priority levels for #GNotifications.
type NotificationPriority C.GNotificationPriority

const (
	// the default priority, to be used for the
	// majority of notifications (for example email messages, software updates,
	// completed download/sync operations)
	NOTIFICATION_PRIORITY_NORMAL NotificationPriority = 0
	// for notifications that do not require
	// immediate attention - typically used for contextual background
	// information, such as contact birthdays or local weather
	NOTIFICATION_PRIORITY_LOW NotificationPriority = 1
	// for events that require more attention,
	// usually because responses are time-sensitive (for example chat and SMS
	// messages or alarms)
	NOTIFICATION_PRIORITY_HIGH NotificationPriority = 2
	// for urgent notifications, or notifications
	// that require a response in a short space of time (for example phone calls
	// or emergency warnings)
	NOTIFICATION_PRIORITY_URGENT NotificationPriority = 3
)
