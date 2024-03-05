package xmlcodec

import (
	"io"
)

type AttributeCharacterDataEncoder interface {
	EncodeXMLAttributeCharacterData(io.Writer) error
}

func ReturnAttributeCharacterDataEncoder(value any) AttributeCharacterDataEncoder {

	switch casted := value.(type) {
	case AttributeCharacterDataEncoder:
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
	return EncodeAttributeCharacterDataString(writer, string(receiver))
}

type bytesAttributeCharacterDataEncoder []byte
func (receiver bytesAttributeCharacterDataEncoder) EncodeXMLAttributeCharacterData(writer io.Writer) error {
	return EncodeAttributeCharacterDataBytes(writer, []byte(receiver))
}
