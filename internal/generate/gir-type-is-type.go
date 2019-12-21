package generate

func (t *Type) isAlias() bool {
	_, found := t.namespace.Aliases.byName(t.Name)
	return found
}

func (t *Type) isBitfield() bool {
	_, found := t.namespace.Bitfields.byName(t.Name)
	return found
}

func (t *Type) isEnumeration() bool {
	_, found := t.namespace.Enumerations.byName(t.Name)
	return found
}

func (t *Type) isClass() bool {
	if t.isQualifiedName() {
		_, found := t.foreignNamespace.Classes.byName(t.foreignName)
		return found
	}

	_, found := t.namespace.Classes.byName(t.Name)
	return found
}

func (t *Type) isRecord() bool {
	//if t.isClass() {
	//	return true
	//}

	if t.isQualifiedName() {
		_, found := t.foreignNamespace.Records.byName(t.foreignName)
		return found
	}

	_, found := t.namespace.Records.byName(t.Name)
	return found
}

func (t *Type) isInterface() bool {
	if t.isQualifiedName() {
		_, found := t.foreignNamespace.Interfaces.byName(t.foreignName)
		return found
	}

	_, found := t.namespace.Interfaces.byName(t.Name)
	return found
}

func (t *Type) isVaList() bool {
	return t.Name == "va_list"
}
