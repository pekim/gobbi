#ifndef GOBBI_SIGNAL_H
#define GOBBI_SIGNAL_H

#include <glib-object.h>

typedef struct {
    void        *instance;
    GClosure    *closure;
    gulong      handler_id;
} GobbiSignalConnection;

GobbiSignalConnection gobbi_connect_signal(void *instance, char *signal_name, void *id);
void gobbi_disconnect_signal(GobbiSignalConnection connection);

#endif
