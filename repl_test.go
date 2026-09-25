package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "Hello World!",
			expected: []string{"hello", "world!"},
		},
		{
			input:    "  This  is  my life ha!",
			expected: []string{"this", "is", "my", "life", "ha!"},
		},
		{
			input:    "",
			expected: []string{},
		},
		{
			input:    "jkfjkl;adfjkskla;",
			expected: []string{"jkfjkl;adfjkskla;"},
		},
	}

	for _, test := range cases {
		actual := cleanInput(test.input)

		if len(actual) != len(test.expected) {
			t.Errorf("The actual length is not the same as the expected %v/%v", len(actual), len(test.expected))
		}
		for i := range actual {
			word := actual[i]
			expectedWord := test.expected[i]
			if word != expectedWord {
				t.Errorf("The word \"%v\" does not match the expected \"%v\"", word, expectedWord)
			}
		}
	}
}
