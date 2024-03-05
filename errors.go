package xml

import (
	"sourcecode.social/reiver/go-erorr"
)

const (
	errEmptyXMLElementName = erorr.Error("xml: empty XML element name")
	errNilReflectedType    = erorr.Error("xml: nil reflected type")
	errNilWriter           = erorr.Error("xml: nil writer")
)
