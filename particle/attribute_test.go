package xmlparticle_test

import (
	"testing"

	"strings"

	"github.com/reiver/go-xml/particle"
)

func TestAttribute_EncodeXMLAttribute(t *testing.T) {

	tests := []struct{
		Attribute xmlparticle.Attribute
		Expected string
	}{
		{
			Attribute: xmlparticle.Attribute{
				Name:"apple",
			},
			Expected: `apple=""`,
		},
		{
			Attribute: xmlparticle.Attribute{
				Name:"BANANA",
			},
			Expected: `BANANA=""`,
		},
		{
			Attribute: xmlparticle.Attribute{
				Name:"Cherry",
			},
			Expected: `Cherry=""`,
		},



		{
			Attribute: xmlparticle.Attribute{
				Name:"apple",
				Value: `Did you know that 5 > 2 & 2 < 7 & that O'Charlie said, "Hello world!"?`,
			},
			Expected: `apple="Did you know that 5 &gt; 2 &amp; 2 &lt; 7 &amp; that O'Charlie said, &quot;Hello world!&quot;?"`,
		},
		{
			Attribute: xmlparticle.Attribute{
				Name:"BANANA",
				Value: `Did you know that 5 > 2 & 2 < 7 & that O'Charlie said, "Hello world!"?`,
			},
			Expected: `BANANA="Did you know that 5 &gt; 2 &amp; 2 &lt; 7 &amp; that O'Charlie said, &quot;Hello world!&quot;?"`,
		},
		{
			Attribute: xmlparticle.Attribute{
				Name:"Cherry",
				Value: `Did you know that 5 > 2 & 2 < 7 & that O'Charlie said, "Hello world!"?`,
			},
			Expected: `Cherry="Did you know that 5 &gt; 2 &amp; 2 &lt; 7 &amp; that O'Charlie said, &quot;Hello world!&quot;?"`,
		},
	}

	for testNumber, test := range tests {

		var actualBuffer strings.Builder

		err := test.Attribute.EncodeXMLAttribute(&actualBuffer)

		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one." , testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("NAME-VALUE-ELEMENT: %#v", test.Attribute)
			continue
		}

		{
			expected := test.Expected
			actual   := actualBuffer.String()

			if expected != actual {
				t.Errorf("For test #%d, the actual 'encoded-xml-attribute' is not what was expected." , testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				t.Logf("NAME-VALUE-ELEMENT: %#v", test.Attribute)
				continue
			}
		}
	}
}
