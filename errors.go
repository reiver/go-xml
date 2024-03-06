package xml

import (
	"sourcecode.social/reiver/go-erorr"
)

const (
	errEmptyXMLElementName     = erorr.Error("xml: empty XML element name")
	errNilReflectedStructField = erorr.Error("xml: nil reflected struct field")
	errNilReflectedType        = erorr.Error("xml: nil reflected type")
	errNilWriter               = erorr.Error("xml: nil writer")
	errNotExported             = erorr.Error("xml: not exported")
)
