package xmlparticle

import (
	"io"
	"strings"

	"sourcecode.social/reiver/go-erorr"
)

type Attribute struct {
	Name string
	Value any
}

func (receiver Attribute) EncodeXMLAttribute(writer io.Writer) (err error) {
	if nil == writer {

		return errNilWriter
	}

	var name string = strings.TrimSpace(receiver.Name)

	if "" == name {
		return errEmptyXMLAttributeName
	}

//@TODO: should we check that 'name' is a valid XML attribute name?

	var value any = receiver.Value
	if nil == value {
		value = ""
	}

	var xmlAttributeCharacterDataEncoder internalAttributeCharacterDataEncoder = attributeCharacterDataEncoder(value)
	if nil == xmlAttributeCharacterDataEncoder {
		return erorr.Errorf("xml: cannot encode %T.Value of type %T into XML attribute character data.", receiver, receiver.Value)
	}

	{
		_, err = io.WriteString(writer, name)
		if nil != err {return err}
		_, err = io.WriteString(writer, `="`)
		if nil != err {return err}

		err = xmlAttributeCharacterDataEncoder.EncodeXMLAttributeCharacterData(writer)
		if nil != err {return err}

		_, err = io.WriteString(writer, `"`)
		if nil != err {return err}
	}

	return nil
}
