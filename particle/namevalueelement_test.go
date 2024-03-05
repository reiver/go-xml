package xmlparticle_test

import (
	"testing"

	"strings"

	"github.com/reiver/go-xml/particle"
)

func TestNameValueElement_EncodeXML(t *testing.T) {

	tests := []struct{
		NameValueElement xmlparticle.NameValueElement
		Expected string
	}{
		{
			NameValueElement: xmlparticle.NameValueElement{
				Name:"apple",
			},
			Expected: "<apple></apple>",
		},
		{
			NameValueElement: xmlparticle.NameValueElement{
				Name:"BANANA",
			},
			Expected: "<BANANA></BANANA>",
		},
		{
			NameValueElement: xmlparticle.NameValueElement{
				Name:"Cherry",
			},
			Expected: "<Cherry></Cherry>",
		},



		{
			NameValueElement: xmlparticle.NameValueElement{
				Name:"apple",
				Value: `Did you know that 5 > 2 & 2 < 7 & that O'Charlie said, "Hello world!"?`,
			},
			Expected: `<apple>Did you know that 5 &gt; 2 &amp; 2 &lt; 7 &amp; that O'Charlie said, "Hello world!"?</apple>`,
		},
		{
			NameValueElement: xmlparticle.NameValueElement{
				Name:"BANANA",
				Value: `Did you know that 5 > 2 & 2 < 7 & that O'Charlie said, "Hello world!"?`,
			},
			Expected: `<BANANA>Did you know that 5 &gt; 2 &amp; 2 &lt; 7 &amp; that O'Charlie said, "Hello world!"?</BANANA>`,
		},
		{
			NameValueElement: xmlparticle.NameValueElement{
				Name:"Cherry",
				Value: `Did you know that 5 > 2 & 2 < 7 & that O'Charlie said, "Hello world!"?`,
			},
			Expected: `<Cherry>Did you know that 5 &gt; 2 &amp; 2 &lt; 7 &amp; that O'Charlie said, "Hello world!"?</Cherry>`,
		},
	}

	for testNumber, test := range tests {

		var actualBuffer strings.Builder

		err := test.NameValueElement.EncodeXML(&actualBuffer)

		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one." , testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("NAME-VALUE-ELEMENT: %#v", test.NameValueElement)
			continue
		}

		{
			expected := test.Expected
			actual   := actualBuffer.String()

			if expected != actual {
				t.Errorf("For test #%d, the actual 'encoded-xml' is not what was expected." , testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				t.Logf("NAME-VALUE-ELEMENT: %#v", test.NameValueElement)
				continue
			}
		}
	}
}
