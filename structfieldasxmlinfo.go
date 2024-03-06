package xml

import (
	"reflect"
)

func structFieldAsXMLInfo(result *structFieldXMLInfo, structField *reflect.StructField) (err error) {

	if nil == structField {
		return errNilReflectedStructField
	}

	if !structField.IsExported() {
		return errNotExported
	}

	{
		tag, found := structField.Tag.Lookup("xml")
		if !found {
			result.Name = structField.Name
			return nil
		}

		parseTagStructField(result, tag)

		if "" == result.Name {
			result.Name = structField.Name
		}
		if "" == result.Name {
			return errEmptyXMLElementName
		}
	}

	return nil
}
