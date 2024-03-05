package xmlcodec

import (
	"io"
)

// EscapeAttributeTextBytes writes the text in 'p' to 'writer', where the following escaping happens:
//
//	'"' -> '&quot;'
//
//	'&' -> '&amp;'
//
//	'<' -> '&lt;'
//
//	'>' -> '&gt;'
func EscapeAttributeTextBytes(writer io.Writer, p []byte) (err error) {

	if nil == writer {
		return errNilWriter
	}

	var buffer [1]byte
	var pp []byte = buffer[:]

	for _, b := range p {
		switch b {
		case '"':
			_, err = writer.Write(quot)
		case '&':
			_, err = writer.Write(amp)
		case '<':
			_, err = writer.Write(lt)
		case '>':
			_, err = writer.Write(gt)
		default:
			buffer[0] = b
			_, err = writer.Write(pp)
		}
		if nil != err {
			return err
		}
	}

	return nil
}
