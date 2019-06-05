#ifndef __MY_CPU_H__
#define __MY_CPU_H__

#include <gtk/gtk.h>
#include <cairo.h>

G_BEGIN_DECLS

/* Standart GObject macros */
#define MY_TYPE_CPU (my_cpu_get_type())
#define MY_CPU(obj) (G_TYPE_CHECK_INSTANCE_CAST((obj), MY_TYPE_CPU, MyCpu))
#define MY_CPU_CLASS(klass) (G_TYPE_CHECK_CLASS_CAST((klass), MY_TYPE_CPU, MyCpuClass))
#define MY_IS_CPU(obj) (G_TYPE_CHECK_INSTANCE_TYPE((obj), MY_TYPE_CPU))
#define MY_IS_CPU_CLASS(klass) (G_TYPE_CHECK_CLASS_TYPE((klass), MY_TYPE_CPU))
#define MY_CPU_GET_CLASS(obj) (G_TYPE_INSTANCE_GET_CLASS((obj), MY_TYPE_CPU, MyCpuClass))

/* Type definition */
typedef struct _MyCpu        MyCpu;
typedef struct _MyCpuClass   MyCpuClass;
typedef struct _MyCpuPrivate MyCpuPrivate;

struct _MyCpu {

   GtkWidget parent;

   /*< Private >*/
   MyCpuPrivate *priv;
};

struct _MyCpuClass {

   GtkWidgetClass parent_class;
};

/* Public API */
GType      my_cpu_get_type(void) G_GNUC_CONST;
GtkWidget *my_cpu_new(void);

gdouble my_cpu_get_percent(MyCpu *cpu);
void    my_cpu_set_percent(MyCpu *cpu, gdouble sel);

G_END_DECLS

#endif /* __MY_CPU_H__ */
