package xml

import (
	"reflect"
)

// xmlElementInfo returns the XML element name, and whether it is an empty XML element (i.e., self-closing) or not, for a struct.
//
// This would be used when encoding the struct to XML.
//
// For example, if this function returns:
//
//	"banana", false
//
// ... then the encoded XML would be something like:
//
//	<banana></banana>
//
// Alternatively, if it returns:
//
//	"banana", true
//
// ... then the encoded XML would be something like::
//
//	<banana />
//
// (Of course, with possible attributes and inner elements.)
func xmlElementInfo(reflectedType reflect.Type) (xmlElementName string, isEmptyXMLElement bool, err error) {

	if nil == reflectedType {
		return "", false, errNilReflectedType
	}

	var reflectedStructField reflect.StructField
	{
		var foundXMLElementField bool

		reflectedStructField, foundXMLElementField = reflectedType.FieldByName("XMLElement")
		if foundXMLElementField {
			tagValue, foundTag := reflectedStructField.Tag.Lookup("xml")
			if foundTag {
				xmlElementName, isEmptyXMLElement = parseTagXMLElement(tagValue)
			}
		}
	}

	if "" == xmlElementName {
		xmlElementName = reflectedType.Name()
	}

	if "" == xmlElementName {
		return "", isEmptyXMLElement, errEmptyXMLElementName
	}

	return xmlElementName, isEmptyXMLElement, nil
}
