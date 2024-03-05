package xmlcodec_test

import (
	"testing"

	"strings"

	"github.com/reiver/go-xml/codec"
)

func TestEncodeCharacterDataString(t *testing.T) {

	tests := []struct{
		Data string
		Expected string
	}{
		{
			Data:     "",
			Expected: "",
		},



		{
			Data:     "apple",
			Expected: "apple",
		},
		{
			Data:     "BANANA",
			Expected: "BANANA",
		},
		{
			Data:     "Cherry",
			Expected: "Cherry",
		},



		{
			Data:     "&",
			Expected: "&amp;",
		},
		{
			Data:     "<",
			Expected: "&lt;",
		},
		{
			Data:     ">",
			Expected: "&gt;",
		},



		{
			Data:     "\"",
			Expected: "\"",
		},
		{
			Data:     "'",
			Expected: "'",
		},



		{
			Data:     "5 > 2 & 5 < 10",
			Expected: "5 &gt; 2 &amp; 5 &lt; 10",
		},



		{
			Data:     "<><><><><> &&& <><><><><>",
			Expected: "&lt;&gt;&lt;&gt;&lt;&gt;&lt;&gt;&lt;&gt; &amp;&amp;&amp; &lt;&gt;&lt;&gt;&lt;&gt;&lt;&gt;&lt;&gt;",
		},



		{
			Data:     `here is a quotation, "hello world!".`,
			Expected: `here is a quotation, "hello world!".`,
		},
	}

	for testNumber, test := range tests {

		var actualBuffer strings.Builder

		err := xmlcodec.EncodeCharacterDataString(&actualBuffer, test.Data)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one." , testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("DATA: %q", test.Data)
			continue
		}

		{
			expected := test.Expected
			actual   := actualBuffer.String()

			if expected != actual {
				t.Errorf("For test #%d, did not expect an error but actually got one." , testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				t.Logf("DATA: %q", test.Data)
				continue
			}
		}
	}
}
