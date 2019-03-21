// This is a generated file - DO NOT EDIT
// +build gtk_2.14 gtk_2.16 gtk_2.18 gtk_2.20 gtk_2.22 gtk_2.24 gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	"sync"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
/*

	gchar* callback_calendardetailfuncHandler(GObject *, GtkCalendar*, guint, guint, guint, gpointer);

*/
import "C"

var callbackCalendardetailfuncId int
var callbackCalendardetailfuncMap = make(map[int]CalendardetailfuncCallback)
var callbackCalendardetailfuncLock sync.RWMutex

// CalendardetailfuncCallback is a callback function for a 'CalendarDetailFunc' callback.
type CalendardetailfuncCallback func(calendar *Calendar, year uint32, month uint32, day uint32) string

//export callback_calendardetailfuncHandler
func callback_calendardetailfuncHandler(_ *C.GObject, c_calendar *C.GtkCalendar, c_year C.guint, c_month C.guint, c_day C.guint, c_user_data C.gpointer) *C.GString {
	callbackCalendardetailfuncLock.RLock()
	defer callbackCalendardetailfuncLock.RUnlock()

	calendar := CalendarNewFromC(unsafe.Pointer(c_calendar))

	year := uint32(c_year)

	month := uint32(c_month)

	day := uint32(c_day)

	index := int(uintptr(c_user_data))
	callback := callbackCalendardetailfuncMap[index]
	retGo := callback(calendar, year, month, day)
	retC :=
		C.CString(retGo)
	return retC
}

// Unsupported : callback ClipboardURIReceivedFunc : unsupported parameter uris, array type interface not found
