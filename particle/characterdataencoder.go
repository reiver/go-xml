package xmlparticle

import (
	"io"

	"github.com/reiver/go-xml/codec"
)

type internalCharacterDataEncoder interface {
	EncodeXMLCharacterData(io.Writer) error
}

func characterDataEncoder(value any) internalCharacterDataEncoder {

	switch casted := value.(type) {
	case internalCharacterDataEncoder:
		return casted
	case string:
		return stringCharacterDataEncoder(casted)
	case []byte:
		return bytesCharacterDataEncoder(casted)
	case []rune:
		return stringCharacterDataEncoder(string(casted))
	default:
		return nil
	}

}

type stringCharacterDataEncoder string
func (receiver stringCharacterDataEncoder) EncodeXMLCharacterData(writer io.Writer) error {
	return xmlcodec.EscapeTextString(writer, string(receiver))
}

type bytesCharacterDataEncoder []byte
func (receiver bytesCharacterDataEncoder) EncodeXMLCharacterData(writer io.Writer) error {
	return xmlcodec.EscapeTextBytes(writer, []byte(receiver))
}
