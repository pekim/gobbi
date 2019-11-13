package generate

type Member struct {
	Name        string `xml:"name,attr"`
	Value       int    `xml:"value,attr"`
	CIdentifier string `xml:"http://www.gtk.org/introspection/c/1.0 identifier,attr"`
	//Doc         *Doc   `xml:"doc"`

	namespace *Namespace
	name      string
	goName    string
}

func (m *Member) init(ns *Namespace, namePrefix string) {
	m.namespace = ns

	//name := strings.TrimPrefix(m.CIdentifier, namePrefix)
	//firstRune, _ := utf8.DecodeRuneInString(name)
	//if unicode.IsDigit(firstRune) {
	//	// Make name a legal Go identifier (cannot start with a digit).
	//	name = "_" + name
	//}
	//m.name = name
	//
	//m.goTypeName = name
}

//func (m *Member) generate(g *jen.Group, goTypeName string) {
//	if m.Blacklist {
//		g.Commentf("Blacklisted member : %s", m.CIdentifier)
//		return
//	}
//
//	// declare a member constant
//	g.
//		Id(m.goTypeName).
//		Id(goTypeName).
//		Op("=").
//		Lit(m.Value)
//}
