#include <gtk/gtk.h>

typedef struct da_d_instance {
    GtkDrawingArea base;
    int instanceId;
} da_d_instance;

extern void drawing_area_class_init(GtkDrawingAreaClass *g_class, gpointer class_data);
