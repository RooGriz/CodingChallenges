package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
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
			name:  "ValidWhenCorrectBrackets1",
			input: "({[]})",
			want:  true,
		},
		{
			name:  "ValidWhenCorrectBrackets2",
			input: "(){}[]",
			want:  true,
		},
		{
			name:  "InvalidWhenOnlyOpenedBracket",
			input: "(",
			want:  false,
		},
		{
			name:  "InvalidWhenOnlyClosedBracket",
			input: ")",
			want:  false,
		},
		{
			name:  "InvalidWhenBracketsWrong1",
			input: ")(",
			want:  false,
		},
		{
			name:  "InvalidWhenBracketsWrong2",
			input: "([)]",
			want:  false,
		},
		{
			name:  "InvalidWhenBracketsWrong2",
			input: "([)",
			want:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isValid(tt.input)
			assert.Equalf(t, tt.want, got, "For string %s, want %t, got %t", tt.input, tt.want, got)
		})
	}
}
