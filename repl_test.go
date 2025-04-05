package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello world ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "HeLLo wOrlD  ",
			expected: []string{"hello", "world"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("For input: %q: expected length %d, got %d", actual, len(c.expected), len(actual))
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("For input %q: at index %d, expected %q, got %q", actual, i, c.expected[i], actual[i])
			}
		}
	}
}
