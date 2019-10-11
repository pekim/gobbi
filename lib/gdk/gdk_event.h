#include <gdk/gdk.h>

/*
    Although cgo can access C struct field names that are keywords,
    by prefixing them with an underscore, it cannot do the same
    for C unions. This causes a problem for the "type" field in
    the GdkEvent union.

    To work around this, define an equivalent struct for GdkEvent.
*/
typedef struct _GdkEvent_
{
	GdkEventType              type;
} GdkEvent_;
