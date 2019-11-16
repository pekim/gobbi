package generate

var argumentGetFunctionNames = map[string]string{
	"gchar":  "Int8",
	"gint8":  "Int8",
	"gshort": "Int16",
	"gint16": "Int16",
	"int":    "Int32",
	"gint":   "Int32",
	"gint32": "Int32",
	"glong":  "Int64",
	"gint64": "Int64",

	"guchar":  "Uint8",
	"guint8":  "Uint8",
	"gushort": "Uint16",
	"guint16": "Uint16",
	"guint":   "Uint32",
	"guint32": "Uint32",
	"gulong":  "Uint64",
	"guint64": "Uint64",

	"utf8": "String",
}

var argumentSetFunctionNames = map[string]string{
	"gchar":  "SetInt8",
	"gint8":  "SetInt8",
	"gshort": "SetInt16",
	"gint16": "SetInt16",
	"int":    "SetInt32",
	"gint":   "SetInt32",
	"gint32": "SetInt32",
	"glong":  "SetInt64",
	"gint64": "SetInt64",

	"guchar":  "SetUint8",
	"guint8":  "SetUint8",
	"gushort": "SetUint16",
	"guint16": "SetUint16",
	"guint":   "SetUint32",
	"guint32": "SetUint32",
	"gulong":  "SetUint64",
	"guint64": "SetUint64",
	//
	//"utf8": "SetString",
}

func (t *Type) supportedAsReturnValue() bool {
	if t == nil || t.Name == "none" {
		// return type is void
		return true
	}

	if _, ok := argumentGetFunctionNames[t.Name]; ok {
		return true
	}

	return false
}

func (t *Type) supportedAsParameter() bool {
	if t == nil || t.Name == "" {
		return false
	}

	if _, ok := argumentSetFunctionNames[t.Name]; ok {
		return true
	}

	return false
}
