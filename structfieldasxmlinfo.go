package xml

import (
	"reflect"
)

func structFieldAsXMLInfo(structField *reflect.StructField) (name string, isAttribute bool, err error) {

	if nil == structField {
		return "", false, errNilReflectedStructField
	}

	if !structField.IsExported() {
		return "", false, errNotExported
	}

	{
		tag, found := structField.Tag.Lookup("xml")
		if !found {
			return structField.Name, false, nil
		}

		name, isAttribute = parseTagStructField(tag)
		if "" == name {
			name = structField.Name
		}
		if "" == name {
			return "", false, errEmptyXMLElementName
		}
	}

	return name, isAttribute, nil
}
