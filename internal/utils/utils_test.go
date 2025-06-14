package utils

import "testing"

func TestFormatWithCommas(t *testing.T) {
	tests := []struct {
		name           string
		input          int
		expectedResult string
	}{
		{
			name:           "No less than 1000",
			input:          123,
			expectedResult: "123",
		},
		{
			name:           "No higher than 1000",
			input:          123456,
			expectedResult: "123,456",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			result := FormatWithCommas(tc.input)
			if result != tc.expectedResult {
				t.Errorf("Expected result %s, got %s", tc.expectedResult, result)
			}
		})
	}
}
