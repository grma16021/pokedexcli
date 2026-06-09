package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " hello world ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "hello world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Hello World",
			expected: []string{"hello", "world"},
		},
		//add more cases here
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		expected := c.expected
		// Check the length of the actual slice against the expected slice
		// if they don't match, use t.Errorf to print an error message
		if len(actual) != len(expected) {
			t.Errorf("%v is not equal to %v", actual, expected)
			return
		}
		// and fail the test
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			// Check each word in the slice
			// if they don't match, use t.Errorf to print an error message
			if word != expectedWord {
				t.Errorf("expected: %v doesent match %v", word, expectedWord)
				return
			}
			// and fail the test
		}
	}
}
