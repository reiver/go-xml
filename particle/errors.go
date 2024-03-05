package xmlparticle

import (
	"sourcecode.social/reiver/go-erorr"
)

const (
	errEmptyXMLAttributeName = erorr.Error("xml: empty XML attribute name")
	errEmptyXMLElementName   = erorr.Error("xml: empty XML element name")
	errNilWriter             = erorr.Error("xml: nil writer")
)
