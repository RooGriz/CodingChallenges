package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostorderTraversal(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want []int
	}{
		{
			name: "Case1",
			root: nil,
			want: []int{},
		},
		{
			name: "Case2",
			root: &TreeNode{
				Left:  nil,
				Right: nil,
				Val:   10,
			},
			want: []int{10},
		},
		{
			name: "Case3",
			root: &TreeNode{
				Left: &TreeNode{
					Left:  nil,
					Right: nil,
					Val:   9,
				},
				Right: &TreeNode{
					Left:  nil,
					Right: nil,
					Val:   11,
				},
				Val: 10,
			},
			want: []int{9, 11, 10},
		},
		{
			name: "Case4",
			root: &TreeNode{
				Left: &TreeNode{
					Left: &TreeNode{
						Left:  nil,
						Right: nil,
						Val:   6,
					},
					Right: &TreeNode{
						Left:  nil,
						Right: nil,
						Val:   8,
					},
					Val: 7,
				},
				Right: &TreeNode{
					Left: &TreeNode{
						Left:  nil,
						Right: nil,
						Val:   12,
					},
					Right: &TreeNode{
						Left:  nil,
						Right: nil,
						Val:   16,
					},
					Val: 14,
				},
				Val: 10,
			},
			want: []int{6, 8, 7, 12, 16, 14, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := postorderTraversal(tt.root)
			assert.Equalf(t, tt.want, got, "Postorder traversal want %v, got %v", tt.want, got)
		})
	}
}
