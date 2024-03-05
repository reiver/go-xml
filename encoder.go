package xml

import (
	"io"
)

// Encoder is an interface implemented by types that can marshal themselves into valid XML.
type Encoder interface {
	EncodeXML(io.Writer) error
}
