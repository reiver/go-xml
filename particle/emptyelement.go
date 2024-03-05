package xmlparticle

import (
	"io"
	"strings"
)

// EmptyElement is a helper type that can be used to encode an empty XML element.
//
// Empty XML elements are sometimes also called self-closing XML elements.
type EmptyElement struct {
	Name string
	Attributes []Attribute
}

func (receiver EmptyElement) EncodeXML(writer io.Writer) (err error) {
	if nil == writer {
		return errNilWriter
	}

	var name string = strings.TrimSpace(receiver.Name)

	if "" == name {
		return errEmptyXMLElementName
	}

//@TODO: should we check that 'name' is a valid XML element name?

	{
		_, err = io.WriteString(writer, "<")
		if nil != err {return err}
		_, err = io.WriteString(writer, name)
		if nil != err {return err}

		for _, attribute := range receiver.Attributes {
			_, err = io.WriteString(writer, " ")
			if nil != err {return err}

			err = attribute.EncodeXMLAttribute(writer)
			if nil != err {return err}
		}

		_, err = io.WriteString(writer, " />")
		if nil != err {return err}
	}

	return nil
}
