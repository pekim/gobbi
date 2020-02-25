package generate

// context provides a chain of contexts representing the path to
// an element in a gir file.
//
// It can be used to print the path of elements when an error
// is detected.
type context struct {
	parent     *context
	entityName string
	name       string
}

func newContext(parent *context, entityName string, name string) *context {
	return &context{
		parent:     parent,
		entityName: entityName,
		name:       name,
	}
}

func (c *context) String() string {
	parentString := ""
	if c.parent != nil {
		parentString = c.parent.String() + " > "
	}

	return parentString + c.entityName + ":" + c.name
}
