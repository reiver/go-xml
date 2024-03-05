package xmlparticle

import (
	"io"

	"github.com/reiver/go-xml/codec"
)

type internalAttributeCharacterDataEncoder interface {
	EncodeXMLAttributeCharacterData(io.Writer) error
}

func attributeCharacterDataEncoder(value any) internalAttributeCharacterDataEncoder {

	switch casted := value.(type) {
	case internalAttributeCharacterDataEncoder:
		return casted
	case string:
		return stringAttributeCharacterDataEncoder(casted)
	case []byte:
		return bytesAttributeCharacterDataEncoder(casted)
	case []rune:
		return stringAttributeCharacterDataEncoder(string(casted))
	default:
		return nil
	}

}

type stringAttributeCharacterDataEncoder string
func (receiver stringAttributeCharacterDataEncoder) EncodeXMLAttributeCharacterData(writer io.Writer) error {
	return xmlcodec.EncodeAttributeCharacterDataString(writer, string(receiver))
}

type bytesAttributeCharacterDataEncoder []byte
func (receiver bytesAttributeCharacterDataEncoder) EncodeXMLAttributeCharacterData(writer io.Writer) error {
	return xmlcodec.EncodeAttributeCharacterDataBytes(writer, []byte(receiver))
}
