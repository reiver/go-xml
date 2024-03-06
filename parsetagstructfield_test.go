package xml

import (
	"testing"
)

func TestParseTagStructField(t *testing.T) {

	tests := []struct{
		Value string
		Expected structFieldXMLInfo
	}{
		{
			Value: "",
			Expected: structFieldXMLInfo{
				Name: "",
				Attr: false,
				OmitEmpty: false,
			},
		},


		{
			Value: "apple",
			Expected: structFieldXMLInfo{
				Name: "apple",
				Attr: false,
				OmitEmpty: false,
			},
		},
		{
			Value: "BANANA",
			Expected: structFieldXMLInfo{
				Name: "BANANA",
				Attr: false,
				OmitEmpty: false,
			},
		},
		{
			Value: "Cherry",
			Expected: structFieldXMLInfo{
				Name: "Cherry",
				Attr: false,
				OmitEmpty: false,
			},
		},



		{
			Value: "apple,attr",
			Expected: structFieldXMLInfo{
				Name: "apple",
				Attr: true,
				OmitEmpty: false,
			},
		},
		{
			Value: "BANANA,attr",
			Expected: structFieldXMLInfo{
				Name: "BANANA",
				Attr: true,
				OmitEmpty: false,
			},
		},
		{
			Value: "Cherry,attr",
			Expected: structFieldXMLInfo{
				Name: "Cherry",
				Attr: true,
				OmitEmpty: false,
			},
		},



		{
			Value: "apple,attr,wow",
			Expected: structFieldXMLInfo{
				Name: "apple",
				Attr: true,
				OmitEmpty: false,
			},
		},
		{
			Value: "BANANA,attr,wow",
			Expected: structFieldXMLInfo{
				Name: "BANANA",
				Attr: true,
				OmitEmpty: false,
			},
		},
		{
			Value: "Cherry,attr,wow",
			Expected: structFieldXMLInfo{
				Name: "Cherry",
				Attr: true,
				OmitEmpty: false,
			},
		},



		{
			Value: "apple,wow,attr",
			Expected: structFieldXMLInfo{
				Name: "apple",
				Attr: true,
				OmitEmpty: false,
			},
		},
		{
			Value: "BANANA,wow,attr",
			Expected: structFieldXMLInfo{
				Name: "BANANA",
				Attr: true,
				OmitEmpty: false,
			},
		},
		{
			Value: "Cherry,wow,attr",
			Expected: structFieldXMLInfo{
				Name: "Cherry",
				Attr: true,
				OmitEmpty: false,
			},
		},



		{
			Value: " Cherry\t,  \t wow  ,\t\t attr\t",
			Expected: structFieldXMLInfo{
				Name: "Cherry",
				Attr: true,
				OmitEmpty: false,
			},
		},
	}

	for testNumber, test := range tests {

		var actual structFieldXMLInfo

		parseTagStructField(&actual, test.Value)

		{
			expected := test.Expected

			if expected != actual {
				t.Errorf("For test #%d, the actual 'structFieldXMLInfo' is not what was expected.", testNumber)
				t.Logf("EXPECTED: %#v", expected)
				t.Logf("ACTUAL:   %#v", actual)
				t.Logf("VALUE: %q", test.Value)
				continue
			}
		}
	}
}
