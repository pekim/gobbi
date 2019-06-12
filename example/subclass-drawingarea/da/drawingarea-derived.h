#include <gtk/gtk.h>

// The instance struct for the class that's derived from GtkDrawingArea.
typedef struct da_d_instance {
    GtkDrawingArea base;

    // somewhere to keep an id that can be used to later lookup
    // a instance of the Go DrawingAreaDerived type.
    int instanceId;
} da_d_instance;


extern void drawing_area_class_init(GtkDrawingAreaClass *g_class, gpointer class_data);
