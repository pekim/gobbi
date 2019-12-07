#include "signal.h"
#include "_cgo_export.h"

void gobbi_dummy_callback(void *instance) {}

void gobbi_closure_marshal (
    GClosure *closure,
    GValue *return_value,
    guint n_param_values,
    const GValue *param_values,
    gpointer invocation_hint,
    gpointer marshal_data)
{
    gobbiCSignalHandler(return_value, n_param_values, (GValue *)param_values, marshal_data);
}

GobbiSignalConnection gobbi_connect_signal(void *instance, char *signal_name, void *data)
{
    GClosure *closure = g_cclosure_new(G_CALLBACK(gobbi_dummy_callback), (gpointer)data, NULL);
    g_closure_ref(closure);

    g_closure_set_meta_marshal(closure, (gpointer)data, gobbi_closure_marshal);
    gulong handler_id = g_signal_connect_closure(instance, signal_name, closure, FALSE);

    GobbiSignalConnection connection = {
        instance,
        closure,
        handler_id,
    };

    return connection;
}

void gobbi_disconnect_signal(GobbiSignalConnection connection) {
    g_signal_handler_disconnect(connection.instance, connection.handler_id);
    g_closure_unref(connection.closure);
}
