/* mycpu.c */

#include "cpu.h"

/* Properties enum */
enum {

   P_0, /* Padding */
   P_PERCENT
};

/* Private data structure */
struct _MyCpuPrivate {

   gdouble percent;
   GdkWindow *window;
};

const gint WIDTH = 80;
const gint HEIGHT = 100;

/* Internal API */
static void my_cpu_set_property(GObject *object, guint prop_id,
    const GValue *value, GParamSpec *pspec);
static void my_cpu_get_property(GObject *object, guint prop_id,
    GValue *value, GParamSpec *pspec);
//static void my_cpu_size_request(GtkWidget *widget,
//    GtkRequisition *requisition);
static void my_cpu_size_allocate(GtkWidget *widget,
    GtkAllocation *allocation);
static void my_cpu_realize(GtkWidget *widget);
//static gboolean my_cpu_expose(GtkWidget *widget,
//    GdkEventExpose *event);

/* Define type */
G_DEFINE_TYPE(MyCpu, my_cpu, GTK_TYPE_WIDGET)

/* Initialization */
static void my_cpu_class_init(MyCpuClass *klass) {

   GObjectClass *g_class;
   GtkWidgetClass *w_class;
   GParamSpec *pspec;

   g_class = G_OBJECT_CLASS(klass);
   w_class = GTK_WIDGET_CLASS(klass);

   /* Override widget class methods */
   g_class->set_property  = my_cpu_set_property;
   g_class->get_property  = my_cpu_get_property;

   w_class->realize       = my_cpu_realize;
//   w_class->size_request  = my_cpu_size_request;
   w_class->size_allocate = my_cpu_size_allocate;
//   w_class->expose_event  = my_cpu_expose;

   /* Install property */
   pspec = g_param_spec_double("percent", "Percent",
       "What CPU load should be displayed", 0, 1, 0,
       G_PARAM_READWRITE | G_PARAM_STATIC_STRINGS);

   g_object_class_install_property(g_class, P_PERCENT, pspec);

   /* Add private data */
   g_type_class_add_private(g_class, sizeof(MyCpuPrivate));
}

static void my_cpu_init(MyCpu *cpu) {

   MyCpuPrivate *priv;

   priv = G_TYPE_INSTANCE_GET_PRIVATE(cpu, MY_TYPE_CPU, MyCpuPrivate);

   gtk_widget_set_has_window(GTK_WIDGET(cpu), TRUE);

   /* Set default values */
   priv->percent = 0;

   /* Create cache for faster access */
   cpu->priv = priv;
}

/* Overriden virtual methods */
static void my_cpu_set_property(GObject *object, guint prop_id,
    const GValue *value, GParamSpec *pspec) {

   MyCpu *cpu = MY_CPU(object);

   switch(prop_id) {

      case P_PERCENT:

         my_cpu_set_percent(cpu, g_value_get_double(value));
         break;

      default:

         G_OBJECT_WARN_INVALID_PROPERTY_ID(object, prop_id, pspec);
         break;
   }
}

static void my_cpu_get_property(GObject *object, guint prop_id,
                GValue *value, GParamSpec *pspec) {

   MyCpu *cpu = MY_CPU(object);

   switch(prop_id) {

      case P_PERCENT:
         g_value_set_double(value, cpu->priv->percent);
         break;

      default:
         G_OBJECT_WARN_INVALID_PROPERTY_ID(object, prop_id, pspec);
         break;
   }
}

static void my_cpu_realize(GtkWidget *widget) {

   MyCpuPrivate *priv = MY_CPU(widget)->priv;
   GtkAllocation alloc;
   GdkWindowAttr attrs;
   guint attrs_mask;

   gtk_widget_set_realized(widget, TRUE);

   gtk_widget_get_allocation(widget, &alloc);

   attrs.x           = alloc.x;
   attrs.y           = alloc.y;
   attrs.width       = alloc.width;
   attrs.height      = alloc.height;
   attrs.window_type = GDK_WINDOW_CHILD;
   attrs.wclass      = GDK_INPUT_OUTPUT;
   attrs.event_mask  = gtk_widget_get_events(widget) | GDK_EXPOSURE_MASK;

   attrs_mask = GDK_WA_X | GDK_WA_Y;

   priv->window = gdk_window_new(gtk_widget_get_parent_window(widget),
               &attrs, attrs_mask);
   gdk_window_set_user_data(priv->window, widget);
   gtk_widget_set_window(widget, priv->window);

//   widget->style = gtk_style_attach(gtk_widget_get_style( widget ),
//                             priv->window);
//   gtk_style_set_background(widget->style, priv->window, GTK_STATE_NORMAL);
}

//static void my_cpu_size_request(GtkWidget *widget,
//    GtkRequisition *requisition) {
//
//   requisition->width  = WIDTH;
//   requisition->height = HEIGHT;
//}

static void my_cpu_size_allocate(GtkWidget *widget,
                 GtkAllocation *allocation) {

   MyCpuPrivate *priv;

   priv = MY_CPU(widget)->priv;

   gtk_widget_set_allocation(widget, allocation);

   if (gtk_widget_get_realized(widget)) {

      gdk_window_move_resize(priv->window, allocation->x, allocation->y,
          WIDTH, HEIGHT);
   }
}

//static gboolean my_cpu_expose(GtkWidget *widget,
//    GdkEventExpose *event) {
//
//   MyCpuPrivate *priv = MY_CPU(widget)->priv;
//   cairo_t *cr;
//   gint limit;
//   gint i;
//
//   cr = gdk_cairo_create(event->window);
//
//   cairo_translate(cr, 0, 7);
//
//   cairo_set_source_rgb(cr, 0, 0, 0);
//   cairo_paint(cr);
//
//   limit = 20 - priv->percent / 5;
//
//   for (i = 1; i <= 20; i++) {
//
//      if (i > limit) {
//         cairo_set_source_rgb(cr, 0.6, 1.0, 0);
//      } else {
//         cairo_set_source_rgb(cr, 0.2, 0.4, 0);
//      }
//
//      cairo_rectangle(cr, 8,  i * 4, 30, 3);
//      cairo_rectangle(cr, 42, i * 4, 30, 3);
//      cairo_fill(cr);
//   }
//
//   cairo_destroy(cr);
//
//   return TRUE;
//}

/* Public API */
GtkWidget *my_cpu_new(void) {

   return(g_object_new(MY_TYPE_CPU, NULL));
}

gdouble my_cpu_get_percent(MyCpu *cpu) {

   g_return_val_if_fail(MY_IS_CPU(cpu), 0);

   return(cpu->priv->percent);
}

void my_cpu_set_percent(MyCpu *cpu, gdouble sel) {

   g_return_if_fail(MY_IS_CPU(cpu));

   cpu->priv->percent = sel;
   gtk_widget_queue_draw(GTK_WIDGET(cpu));
}
