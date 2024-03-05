package xmlcodec

import (
	"io"
)

type CharacterDataEncoder interface {
	EncodeXMLCharacterData(io.Writer) error
}

func ReturnCharacterDataEncoder(value any) CharacterDataEncoder {

	switch casted := value.(type) {
	case CharacterDataEncoder:
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
	return EncodeCharacterDataString(writer, string(receiver))
}

type bytesCharacterDataEncoder []byte
func (receiver bytesCharacterDataEncoder) EncodeXMLCharacterData(writer io.Writer) error {
	return EncodeCharacterDataBytes(writer, []byte(receiver))
}
