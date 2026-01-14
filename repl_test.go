package main

import (
	"testing"
)

func TestCleannput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		}, {
			input:    "charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		}, {
			input:    "thisistextnowhitespace",
			expected: []string{"thisistextnowhitespace"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Error: length of input does not match expected: %d; want %d", len(actual), len(c.expected))
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Error: input does not match expected (%s); want %s", word, expectedWord)
			}
		}
	}
}
