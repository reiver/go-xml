package xml

import (
	"testing"
)

func TestParseTagXMLElement(t *testing.T) {

	tests := []struct{
		Value string
		ExpectedName string
		ExpectedIsEmpty bool
	}{
		{
			Value: "",
			ExpectedName: "",
			ExpectedIsEmpty: false,
		},


		{
			Value: "apple",
			ExpectedName: "apple",
			ExpectedIsEmpty: false,
		},
		{
			Value: "BANANA",
			ExpectedName: "BANANA",
			ExpectedIsEmpty: false,
		},
		{
			Value: "Cherry",
			ExpectedName: "Cherry",
			ExpectedIsEmpty: false,
		},



		{
			Value: "apple,empty",
			ExpectedName: "apple",
			ExpectedIsEmpty: true,
		},
		{
			Value: "BANANA,empty",
			ExpectedName: "BANANA",
			ExpectedIsEmpty: true,
		},
		{
			Value: "Cherry,empty",
			ExpectedName: "Cherry",
			ExpectedIsEmpty: true,
		},



		{
			Value: "apple,empty,wow",
			ExpectedName: "apple",
			ExpectedIsEmpty: true,
		},
		{
			Value: "BANANA,empty,wow",
			ExpectedName: "BANANA",
			ExpectedIsEmpty: true,
		},
		{
			Value: "Cherry,empty,wow",
			ExpectedName: "Cherry",
			ExpectedIsEmpty: true,
		},



		{
			Value: "apple,wow,empty",
			ExpectedName: "apple",
			ExpectedIsEmpty: true,
		},
		{
			Value: "BANANA,wow,empty",
			ExpectedName: "BANANA",
			ExpectedIsEmpty: true,
		},
		{
			Value: "Cherry,wow,empty",
			ExpectedName: "Cherry",
			ExpectedIsEmpty: true,
		},



		{
			Value: " Cherry\t,  \t wow  ,\t\t empty\t",
			ExpectedName: "Cherry",
			ExpectedIsEmpty: true,
		},
	}

	for testNumber, test := range tests {

		actualName, actualIsEmpty := parseTagXMLElement(test.Value)

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
			expected := test.ExpectedIsEmpty
			actual   := actualIsEmpty

			if expected != actual {
				t.Errorf("For test #%d, the actual 'is-empty' is not what was expected.", testNumber)
				t.Logf("EXPECTED: %t", expected)
				t.Logf("ACTUAL:   %t", actual)
				t.Logf("VALUE: %q", test.Value)
				continue
			}
		}
	}
}
