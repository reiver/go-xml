package xml_test

import (
	"testing"

	"strings"

	"github.com/reiver/go-xml"
)

func TestEncode(t *testing.T) {

	type Link1 struct {
		Rel  string
		HRef string
	}



	type Link2 struct {
		Rel  string `xml:"rel"`
		HRef string
	}
	type Link3 struct {
		Rel  string
		HRef string `xml:"href"`
	}
	type Link4 struct {
		Rel  string `xml:"rel"`
		HRef string `xml:"href"`
	}



	type Link5 struct {
		Rel  string `xml:",attr"`
		HRef string `xml:",attr"`
	}



	type Something1 struct {
		XMLElement xml.Meta `xml:"something"`
		Rel  string `xml:"rel,attr"`
		HREF string `xml:",attr"`
		Name string `xml:"name"`
		Wow  string
	}


	tests := []struct{
		Value any
		Expected string
	}{
		{
			Value: Link1{},
			Expected:
				`<Link1>`+
					`<Rel></Rel>`+
					`<HRef></HRef>`+
				`</Link1>`,
		},
		{
			Value: Link1{
				Rel: "something",
			},
			Expected:
				`<Link1>`+
					`<Rel>something</Rel>`+
					`<HRef></HRef>`+
				`</Link1>`,
		},
		{
			Value: Link1{
				HRef: "http://example.com/once/twice/thrice/fource",
			},
			Expected:
				`<Link1>`+
					`<Rel></Rel>`+
					`<HRef>http://example.com/once/twice/thrice/fource</HRef>`+
				`</Link1>`,
		},
		{
			Value: Link1{
				Rel: "something",
				HRef: "http://example.com/once/twice/thrice/fource",
			},
			Expected:
				`<Link1>`+
					`<Rel>something</Rel>`+
					`<HRef>http://example.com/once/twice/thrice/fource</HRef>`+
				`</Link1>`,
		},



		{
			Value: Link2{},
			Expected:
				`<Link2>`+
					`<rel></rel>`+
					`<HRef></HRef>`+
				`</Link2>`,
		},
		{
			Value: Link2{
				Rel: "something",
			},
			Expected:
				`<Link2>`+
					`<rel>something</rel>`+
					`<HRef></HRef>`+
				`</Link2>`,
		},
		{
			Value: Link2{
				HRef: "http://example.com/once/twice/thrice/fource",
			},
			Expected:
				`<Link2>`+
					`<rel></rel>`+
					`<HRef>http://example.com/once/twice/thrice/fource</HRef>`+
				`</Link2>`,
		},
		{
			Value: Link2{
				Rel: "something",
				HRef: "http://example.com/once/twice/thrice/fource",
			},
			Expected:
				`<Link2>`+
					`<rel>something</rel>`+
					`<HRef>http://example.com/once/twice/thrice/fource</HRef>`+
				`</Link2>`,
		},



		{
			Value: Link3{},
			Expected:
				`<Link3>`+
					`<Rel></Rel>`+
					`<href></href>`+
				`</Link3>`,
		},
		{
			Value: Link3{
				Rel: "something",
			},
			Expected:
				`<Link3>`+
					`<Rel>something</Rel>`+
					`<href></href>`+
				`</Link3>`,
		},
		{
			Value: Link3{
				HRef: "http://example.com/once/twice/thrice/fource",
			},
			Expected:
				`<Link3>`+
					`<Rel></Rel>`+
					`<href>http://example.com/once/twice/thrice/fource</href>`+
				`</Link3>`,
		},
		{
			Value: Link3{
				Rel: "something",
				HRef: "http://example.com/once/twice/thrice/fource",
			},
			Expected:
				`<Link3>`+
					`<Rel>something</Rel>`+
					`<href>http://example.com/once/twice/thrice/fource</href>`+
				`</Link3>`,
		},



		{
			Value: Link4{},
			Expected:
				`<Link4>`+
					`<rel></rel>`+
					`<href></href>`+
				`</Link4>`,
		},
		{
			Value: Link4{
				Rel: "something",
			},
			Expected:
				`<Link4>`+
					`<rel>something</rel>`+
					`<href></href>`+
				`</Link4>`,
		},
		{
			Value: Link4{
				HRef: "http://example.com/once/twice/thrice/fource",
			},
			Expected:
				`<Link4>`+
					`<rel></rel>`+
					`<href>http://example.com/once/twice/thrice/fource</href>`+
				`</Link4>`,
		},
		{
			Value: Link4{
				Rel: "something",
				HRef: "http://example.com/once/twice/thrice/fource",
			},
			Expected:
				`<Link4>`+
					`<rel>something</rel>`+
					`<href>http://example.com/once/twice/thrice/fource</href>`+
				`</Link4>`,
		},



		{
			Value: Link5{},
			Expected:
				`<Link5 Rel="" HRef="">`+
				`</Link5>`,
		},
		{
			Value: Link5{
				Rel: "something",
			},
			Expected:
				`<Link5 Rel="something" HRef="">`+
				`</Link5>`,
		},
		{
			Value: Link5{
				HRef: "http://example.com/once/twice/thrice/fource",
			},
			Expected:
				`<Link5 Rel="" HRef="http://example.com/once/twice/thrice/fource">`+
				`</Link5>`,
		},
		{
			Value: Link5{
				Rel: "something",
				HRef: "http://example.com/once/twice/thrice/fource",
			},
			Expected:
				`<Link5 Rel="something" HRef="http://example.com/once/twice/thrice/fource">`+
				`</Link5>`,
		},



		{
			Value: Something1 {
				Rel: "enclosure",
				HREF: "inline:Hello world!",
				Name: "Joe Blow",
				Wow: `2 < 5 & 5 > 3 & O'Charlie said, "Hello world!".`,
			},
			Expected:
				`<something rel="enclosure" HREF="inline:Hello world!">`+
					`<name>Joe Blow</name>`+
					`<Wow>2 &lt; 5 &amp; 5 &gt; 3 &amp; O'Charlie said, "Hello world!".</Wow>`+
				`</something>`,
		},

	}

	for testNumber, test := range tests {

		var actualBuffer strings.Builder

		err := xml.Encode(&actualBuffer, test.Value)

		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("VALUE: %#v", test.Value)
			continue
		}

		{
			expected := test.Expected
			actual   := actualBuffer.String()

			if expected != actual {
				t.Errorf("For test #%d, the actual 'encode' value is not what was expected.", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				t.Logf("VALUE: %#v", test.Value)
				continue
			}
		}

	}
}
