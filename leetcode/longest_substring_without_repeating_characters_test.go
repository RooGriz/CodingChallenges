package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "Empty",
			input: "",
			want:  0,
		},
		{
			name:  "OneCharacterString",
			input: "a",
			want:  1,
		},
		{
			name:  "TwoDistrictCharactersString",
			input: "ab",
			want:  2,
		},
		{
			name:  "TwoSameCharactersString",
			input: "aa",
			want:  1,
		},
		{
			name:  "SeveralCombinationsInString",
			input: "abcabcbb",
			want:  3,
		},
		{
			name:  "IncreasingLongestSubstring",
			input: "aababcabcd",
			want:  4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lengthOfLongestSubstring(tt.input)
			assert.Equalf(t, tt.want, got, "For string %s, want %d, got %d", tt.input, tt.want, got)
		})
	}
}
