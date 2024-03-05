package xmlcodec

import (
	"io"
)

// EscapeAttributeTextString writes the text in 'p' to 'writer', where the following escaping happens:
//
//	'"' -> '&quot;'
//
//	'&' -> '&amp;'
//
//	'<' -> '&lt;'
//
//	'>' -> '&gt;'
func EscapeAttributeTextString(writer io.Writer, p string) (err error) {

	var buffer [1]byte
	var pp []byte = buffer[:]

	var length int = len(p)

	for i:=0; i<length; i++ {
		var b byte = p[i]

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
