package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidPalindrome(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{
			name:  "ValidWhenEmpty",
			input: "",
			want:  true,
		},
		{
			name:  "ValidWhenOnlySymbols",
			input: "  , ",
			want:  true,
		},
		{
			name:  "Valid1",
			input: "a",
			want:  true,
		},
		{
			name:  "Valid2",
			input: "aa",
			want:  true,
		},
		{
			name:  "Valid3",
			input: "aba",
			want:  true,
		},
		{
			name:  "Valid4",
			input: "abbA",
			want:  true,
		},
		{
			name:  "Valid5",
			input: "1 b,b,  1",
			want:  true,
		},
		{
			name:  "NotValid1",
			input: "ab",
			want:  false,
		},
		{
			name:  "NotValid2",
			input: "abca",
			want:  false,
		},
		{
			name:  "NotValid3",
			input: "a 2 3 a",
			want:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isPalindrome(tt.input)
			assert.Equalf(t, tt.want, got, "For string %s, want %t, got %t", tt.input, tt.want, got)
		})
	}
}
