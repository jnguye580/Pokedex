package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "leading and trailing spaces",
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			name:     "empty string",
			input:    "",
			expected: []string{},
		},
		{
			name:     "whitespace only",
			input:    "   ",
			expected: []string{},
		},
		{
			name:     "uppercase letters are lowercased",
			input:    "Hello World",
			expected: []string{"hello", "world"},
		},
		{
			name:     "mixed case with extra spaces",
			input:    "  PIKACHU   Bulbasaur  ",
			expected: []string{"pikachu", "bulbasaur"},
		},
		{
			name:     "single word",
			input:    "charmander",
			expected: []string{"charmander"},
		},
		{
			name:     "already lowercase",
			input:    "fire water grass",
			expected: []string{"fire", "water", "grass"},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := cleanInput(c.input)
			if len(actual) != len(c.expected) {
				t.Errorf("got %v, want %v", actual, c.expected)
				return
			}
			for i := range actual {
				if actual[i] != c.expected[i] {
					t.Errorf("word %d: got %q, want %q", i, actual[i], c.expected[i])
				}
			}
		})
	}
}
