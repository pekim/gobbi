package generate

import "strings"

type Type struct {
	Namespace *Namespace

	Name  string `xml:"name,attr"`
	CType string `xml:"http://www.gtk.org/introspection/c/1.0 type,attr"`
}

// parseStringCType parses the CType on the assumption that it
// is for a string-like type.
//
// Returns two values.
// The C type, such as "gchar".
// The level of pointer indirection.
//
// For example "const gchar**" would return "gchar" and 2.
func (t *Type) parseStringCType() (string, int) {
	cTypeParts := strings.Split(t.CType, " ")
	cType := cTypeParts[len(cTypeParts)-1]
	cTypeName := strings.TrimRight(cType, "*")
	indirectLevel := len(cType) - len(cTypeName)

	return cTypeName, indirectLevel
}
