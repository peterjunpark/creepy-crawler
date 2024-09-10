package main

import (
	"reflect"
	"testing"
)

func TestGetURLsFromHTML(t *testing.T) {
	tests := []struct {
		name            string
		inputHTMLBody   string
		inputRawBaseURL string
		expected        []string
	}{
		{
			name:            "remove scheme",
			inputHTMLBody:   "<body></body>",
			inputRawBaseURL: "https://blog.peterjunpark.com/path",
			expected:        []string{"blog.peterjunpark.com/path"},
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := getURLsFromHTML(tc.inputHTMLBody, tc.inputRawBaseURL)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			}
			if reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}
