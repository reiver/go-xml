package xmlcodec

import (
	"io"
)

// EscapeTextBytes writes the text in 'p' to 'writer', where the following escaping happens:
//
//	'&' -> '&amp;'
//
//	'<' -> '&lt;'
//
//	'>' -> '&gt;'
func EscapeTextBytes(writer io.Writer, p []byte) (err error) {

	var buffer [1]byte
	var pp []byte = buffer[:]

	for _, b := range p {
		switch b {
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
