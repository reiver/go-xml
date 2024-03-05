package xmlparticle

import (
	"io"
	"strings"

	"sourcecode.social/reiver/go-erorr"

	"github.com/reiver/go-xml/codec"
)

// NameValueElement is a helper type that can be used to encode an simple name-value XML attribute (that does NOT have any attributes and whose content is only character text).
type NameValueElement struct {
	Name string
	Value any
}

func (receiver NameValueElement) EncodeXML(writer io.Writer) (err error) {
	if nil == writer {
		return errNilWriter
	}

	var name string = strings.TrimSpace(receiver.Name)

	if "" == name {
		return errEmptyXMLElementName
	}

//@TODO: should we check that 'name' is a valid XML element name?

	var value any = receiver.Value
	if nil == value {
		value = ""
	}

	var xmlCharacterDataEncoder xmlcodec.CharacterDataEncoder = xmlcodec.ReturnCharacterDataEncoder(value)
	if nil == xmlCharacterDataEncoder {
		return erorr.Errorf("xml: cannot encode %T.Value of type %T into XML character data.", receiver, receiver.Value)
	}

	{
		_, err = io.WriteString(writer, "<")
		if nil != err {return err}
		_, err = io.WriteString(writer, name)
		if nil != err {return err}
		_, err = io.WriteString(writer, ">")
		if nil != err {return err}

		err = xmlCharacterDataEncoder.EncodeXMLCharacterData(writer)
		if nil != err {return err}

		_, err = io.WriteString(writer, "</")
		if nil != err {return err}
		_, err = io.WriteString(writer, name)
		if nil != err {return err}
		_, err = io.WriteString(writer, ">")
		if nil != err {return err}
	}

	return nil
}
