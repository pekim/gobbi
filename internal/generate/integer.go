package generate

var integerCTypeMap = map[string]string{
	"gshort": "int16",
	"gint":   "int32",
	"int":    "int32",
	"glong":  "int64",

	"guchar":  "uint8",
	"gushort": "uint32",
	"guint":   "uint32",
	"gulong":  "uint64",

	"gint8":  "int8",
	"gint16": "int16",
	"gint32": "int32",
	"gint64": "int64",

	"guint8":  "uint8",
	"guint16": "uint16",
	"guint32": "uint32",
	"guint64": "uint64",

	"gfloat":  "float32",
	"gdouble": "float64",
	"double":  "float64",

	"gchar":    "rune",
	"gunichar": "rune",

	//"gpointer":      "uintptr",
	//"gconstpointer": "uintptr",
	"goffset": "uint64",
	"gsize":   "uint64",
	"gssize":  "int64",
}
