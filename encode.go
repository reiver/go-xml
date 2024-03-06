package xml

import (
	"io"
	"reflect"

	"sourcecode.social/reiver/go-erorr"

	"github.com/reiver/go-xml/particle"
)

// Encode writes the XML encoding of 'value'.
//
// 'value' should either fit the xml.Encoder interface or be a struct.
func Encode(writer io.Writer, value any) error {

	if nil == writer {
		return errNilWriter
	}

	switch casted := value.(type) {
	case Encoder:
		return casted.EncodeXML(writer)
	default:
		switch reflect.TypeOf(value).Kind() {
		case reflect.Struct:
			return encodeStruct(writer, value)
		default:
			return erorr.Errorf("xml: cannot encode type %T to XML: no method EncodeXML(io.Writer)error", value)
		}
	}
}

func encodeStruct(writer io.Writer, value any) (err error) {

	if nil == writer {
		return errNilWriter
	}

	var reflectedType reflect.Type = reflect.TypeOf(value)
	if nil == reflectedType {
		return errNilReflectedType
	}
	if reflect.Struct != reflectedType.Kind() {
		return erorr.Errorf("xml: cannot encode type %T to XML: not a struct", value)
	}


	xmlElementName, emptyXmlElement, err := xmlElementInfo(reflectedType)
	if nil != err {
		return err
	}

	io.WriteString(writer, "<")
	io.WriteString(writer, xmlElementName)

	var reflectedValue reflect.Value = reflect.ValueOf(value)

	err = encodeStructAttributes(writer, reflectedType, reflectedValue)
	if nil != err {return err}

	if emptyXmlElement {
		_, err = io.WriteString(writer, " />")
		if nil != err {return err}
		return nil
	} else {
		_, err = io.WriteString(writer, ">")
		if nil != err {return err}
	}

	err = encodeStructInner(writer, reflectedType, reflectedValue)
	if nil != err {return err}

	{
		_, err = io.WriteString(writer, "</")
		if nil != err {return err}
		_, err = io.WriteString(writer, xmlElementName)
		if nil != err {return err}
		_, err = io.WriteString(writer, ">")
		if nil != err {return err}
	}

	return nil
}

func encodeStructAttributes(writer io.Writer, reflectedType reflect.Type, reflectedValue reflect.Value) (err error) {

	if nil == writer {
		return errNilWriter
	}

	// attributes
	for i:=0; i < reflectedType.NumField(); i++ {

		var reflectedStructField reflect.StructField = reflectedType.Field(i)

		var xmlinfo structFieldXMLInfo
		err = structFieldAsXMLInfo(&xmlinfo, &reflectedStructField)

		if errNotExported == err {
			continue
		}
		if nil != err {
			return err
		}
		if !xmlinfo.Attr {
			continue
		}
		if "" == xmlinfo.Name {
			continue
		}

		if reflectedTypeMeta == reflectedStructField.Type {
			continue
		}

		var reflectedStructFieldValue reflect.Value = reflectedValue.Field(i)
		var attributeValue = reflectedStructFieldValue.Interface()

		_, err = io.WriteString(writer, " ")
		if nil != err {return err}

		var attribute xmlparticle.Attribute = xmlparticle.Attribute{
			Name: xmlinfo.Name,
			Value: attributeValue,
		}

		err = attribute.EncodeXMLAttribute(writer)
		if nil != err {return err}
	}

	return nil
}

func encodeStructInner(writer io.Writer, reflectedType reflect.Type, reflectedValue reflect.Value) (err error) {

	if nil == writer {
		return errNilWriter
	}

	// nested elements
	for i:=0; i < reflectedType.NumField(); i++ {

		var reflectedStructField reflect.StructField = reflectedType.Field(i)

		var xmlinfo structFieldXMLInfo
		err = structFieldAsXMLInfo(&xmlinfo, &reflectedStructField)

		if errNotExported == err {
			continue
		}
		if nil != err {
			return err
		}
		if xmlinfo.Attr {
			continue
		}
		if "" == xmlinfo.Name {
			continue
		}

		if reflectedTypeMeta == reflectedStructField.Type {
			continue
		}

		var reflectedStructFieldValue reflect.Value = reflectedValue.Field(i)
		var elementValue = reflectedStructFieldValue.Interface()

		var namevalue xmlparticle.NameValueElement = xmlparticle.NameValueElement{
			Name: xmlinfo.Name,
			Value: elementValue,
		}

		err = namevalue.EncodeXML(writer)
		if nil != err {return err}
	}

	return nil
}
