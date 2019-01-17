// This is a generated file - DO NOT EDIT
// +build gtk_2.14 gtk_2.16 gtk_2.18 gtk_2.20 gtk_2.22 gtk_2.24 gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import "sync"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
/*

	gchar* callback_calendardetailfuncHandler(GObject *, GtkCalendar*, guint, guint, guint, gpointer, gpointer);

*/
/*

	void callback_clipboardurireceivedfuncHandler(GObject *, GtkClipboard*, gchar**, gpointer, gpointer);

*/
import "C"

var callbackCalendardetailfuncId int
var callbackCalendardetailfuncMap = make(map[int]CalendardetailfuncCallback)
var callbackCalendardetailfuncLock sync.RWMutex

// CalendardetailfuncCallback is a callback function for a 'CalendarDetailFunc' callback.
type CalendardetailfuncCallback func(calendar *Calendar, year uint32, month uint32, day uint32) string

var callbackClipboardurireceivedfuncId int
var callbackClipboardurireceivedfuncMap = make(map[int]ClipboardurireceivedfuncCallback)
var callbackClipboardurireceivedfuncLock sync.RWMutex

// ClipboardurireceivedfuncCallback is a callback function for a 'ClipboardURIReceivedFunc' callback.
type ClipboardurireceivedfuncCallback func(clipboard *Clipboard, uris []string)
