package xml

import (
	"testing"
)

func TestParseTagStructField(t *testing.T) {

	tests := []struct{
		Value string
		ExpectedName string
		ExpectedIsAttribute bool
	}{
		{
			Value: "",
			ExpectedName: "",
			ExpectedIsAttribute: false,
		},


		{
			Value: "apple",
			ExpectedName: "apple",
			ExpectedIsAttribute: false,
		},
		{
			Value: "BANANA",
			ExpectedName: "BANANA",
			ExpectedIsAttribute: false,
		},
		{
			Value: "Cherry",
			ExpectedName: "Cherry",
			ExpectedIsAttribute: false,
		},



		{
			Value: "apple,attr",
			ExpectedName: "apple",
			ExpectedIsAttribute: true,
		},
		{
			Value: "BANANA,attr",
			ExpectedName: "BANANA",
			ExpectedIsAttribute: true,
		},
		{
			Value: "Cherry,attr",
			ExpectedName: "Cherry",
			ExpectedIsAttribute: true,
		},



		{
			Value: "apple,attr,wow",
			ExpectedName: "apple",
			ExpectedIsAttribute: true,
		},
		{
			Value: "BANANA,attr,wow",
			ExpectedName: "BANANA",
			ExpectedIsAttribute: true,
		},
		{
			Value: "Cherry,attr,wow",
			ExpectedName: "Cherry",
			ExpectedIsAttribute: true,
		},



		{
			Value: "apple,wow,attr",
			ExpectedName: "apple",
			ExpectedIsAttribute: true,
		},
		{
			Value: "BANANA,wow,attr",
			ExpectedName: "BANANA",
			ExpectedIsAttribute: true,
		},
		{
			Value: "Cherry,wow,attr",
			ExpectedName: "Cherry",
			ExpectedIsAttribute: true,
		},



		{
			Value: " Cherry\t,  \t wow  ,\t\t attr\t",
			ExpectedName: "Cherry",
			ExpectedIsAttribute: true,
		},
	}

	for testNumber, test := range tests {

		actualName, actualIsAttribute := parseTagStructField(test.Value)

		{
			expected := test.ExpectedName
			actual   := actualName

			if expected != actual {
				t.Errorf("For test #%d, the actual 'name' is not what was expected.", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				t.Logf("VALUE: %q", test.Value)
				continue
			}
		}

		{
			expected := test.ExpectedIsAttribute
			actual   := actualIsAttribute

			if expected != actual {
				t.Errorf("For test #%d, the actual 'is-attribute' is not what was expected.", testNumber)
				t.Logf("EXPECTED: %t", expected)
				t.Logf("ACTUAL:   %t", actual)
				t.Logf("VALUE: %q", test.Value)
				continue
			}
		}
	}
}
