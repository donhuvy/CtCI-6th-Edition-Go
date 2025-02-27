package chapter1

import (
	"testing"
)

func TestIsUnique(t *testing.T) {
	cases := []struct {
		input    string
		expected bool
	}{
		{"abcd", true},
		{"abcc", false},
		{" ", true},
		{"", true},
	}
	for _, c := range cases {
		actual := IsUnique(c.input)
		if actual != c.expected {
			t.Fatalf("Input %v. Expected: %v, actual: %v\n", c.input, c.expected, actual)
		}
	}
}
