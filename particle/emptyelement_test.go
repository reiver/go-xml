package xmlparticle_test

import (
	"testing"

	"strings"

	"github.com/reiver/go-xml/particle"
)

func TestEmptyElement_EncodeXML(t *testing.T) {

	tests := []struct{
		EmptyElement xmlparticle.EmptyElement
		Expected string
	}{
		{
			EmptyElement: xmlparticle.EmptyElement{
				Name:"apple",
			},
			Expected: "<apple />",
		},
		{
			EmptyElement: xmlparticle.EmptyElement{
				Name:"BANANA",
			},
			Expected: "<BANANA />",
		},
		{
			EmptyElement: xmlparticle.EmptyElement{
				Name:"Cherry",
			},
			Expected: "<Cherry />",
		},



		{
			EmptyElement: xmlparticle.EmptyElement{
				Name:"apple",
				Attributes: []xmlparticle.Attribute{
					xmlparticle.Attribute{
						Name: `something`,
						Value: `Did you know that 5 > 2 & 2 < 7 & that O'Charlie said, "Hello world!"?`,
					},
				},
			},
			Expected: `<apple something="Did you know that 5 &gt; 2 &amp; 2 &lt; 7 &amp; that O'Charlie said, &quot;Hello world!&quot;?" />`,
		},
		{
			EmptyElement: xmlparticle.EmptyElement{
				Name:"BANANA",
				Attributes: []xmlparticle.Attribute{
					xmlparticle.Attribute{
						Name: `something`,
						Value: `Did you know that 5 > 2 & 2 < 7 & that O'Charlie said, "Hello world!"?`,
					},
				},
			},
			Expected: `<BANANA something="Did you know that 5 &gt; 2 &amp; 2 &lt; 7 &amp; that O'Charlie said, &quot;Hello world!&quot;?" />`,
		},
		{
			EmptyElement: xmlparticle.EmptyElement{
				Name:"Cherry",
				Attributes: []xmlparticle.Attribute{
					xmlparticle.Attribute{
						Name: `something`,
						Value: `Did you know that 5 > 2 & 2 < 7 & that O'Charlie said, "Hello world!"?`,
					},
				},
			},
			Expected: `<Cherry something="Did you know that 5 &gt; 2 &amp; 2 &lt; 7 &amp; that O'Charlie said, &quot;Hello world!&quot;?" />`,
		},



		{
			EmptyElement: xmlparticle.EmptyElement{
				Name:"link",
				Attributes: []xmlparticle.Attribute{
					xmlparticle.Attribute{
						Name: `rel`,
						Value: `enclosure`,
					},
					xmlparticle.Attribute{
						Name: `href`,
						Value: `http://example.com/path/to/video.php`,
					},
				},
			},
			Expected: `<link rel="enclosure" href="http://example.com/path/to/video.php" />`,
		},
	}

	for testNumber, test := range tests {

		var actualBuffer strings.Builder

		err := test.EmptyElement.EncodeXML(&actualBuffer)

		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one." , testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("NAME-VALUE-ELEMENT: %#v", test.EmptyElement)
			continue
		}

		{
			expected := test.Expected
			actual   := actualBuffer.String()

			if expected != actual {
				t.Errorf("For test #%d, the actual 'encoded-xml' is not what was expected." , testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				t.Logf("NAME-VALUE-ELEMENT: %#v", test.EmptyElement)
				continue
			}
		}
	}
}
