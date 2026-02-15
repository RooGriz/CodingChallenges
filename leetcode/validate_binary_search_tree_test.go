package leetcode

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValidBST(t *testing.T) {
	tests := []struct {
		name string
		tree *TreeNode
		want bool
	}{
		{
			name: "ValidWhenEmptyTree",
			tree: nil,
			want: true,
		},
		{
			name: "ValidWhenTreeHasOneNode1",
			tree: &TreeNode{
				Left:  nil,
				Right: nil,
				Val:   10,
			},
			want: true,
		},
		{
			name: "ValidWhenTreeHasOneNode2",
			tree: &TreeNode{
				Left:  nil,
				Right: nil,
				Val:   math.MaxInt32,
			},
			want: true,
		},
		{
			name: "ValidWhenTreeHasOneNode3",
			tree: &TreeNode{
				Left:  nil,
				Right: nil,
				Val:   math.MinInt32,
			},
			want: true,
		},
		{
			name: "ValidWhenTreeHasSeveralNodes",
			tree: &TreeNode{
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
			want: true,
		},
		{
			name: "InvalidWhenTreeHasSeveralNodes1",
			tree: &TreeNode{
				Left: &TreeNode{
					Left:  nil,
					Right: nil,
					Val:   9,
				},
				Right: &TreeNode{
					Left:  nil,
					Right: nil,
					Val:   7,
				},
				Val: 10,
			},
			want: false,
		},
		{
			name: "InvalidWhenTreeHasSeveralNodes2",
			tree: &TreeNode{
				Left: &TreeNode{
					Left: nil,
					Right: &TreeNode{
						Left:  nil,
						Right: nil,
						Val:   10,
					},
					Val: 8,
				},
				Right: &TreeNode{
					Left:  nil,
					Right: nil,
					Val:   15,
				},
				Val: 10,
			},
			want: false,
		},
		{
			name: "InvalidWhenTreeHasSeveralNodes3",
			tree: &TreeNode{
				Left: &TreeNode{
					Left:  nil,
					Right: nil,
					Val:   8,
				},
				Right: &TreeNode{
					Left: &TreeNode{
						Left:  nil,
						Right: nil,
						Val:   10,
					},
					Right: nil,
					Val:   15,
				},
				Val: 10,
			},
			want: false,
		},
		{
			name: "InvalidWhenTreeHasSeveralNodes4",
			tree: &TreeNode{
				Left: &TreeNode{
					Left:  nil,
					Right: nil,
					Val:   math.MaxInt32,
				},
				Right: nil,
				Val:   math.MaxInt32,
			},
			want: false,
		},
		{
			name: "InvalidWhenTreeHasSeveralNodes5",
			tree: &TreeNode{
				Left: nil,
				Right: &TreeNode{
					Left:  nil,
					Right: nil,
					Val:   math.MinInt32,
				},
				Val: math.MinInt32,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isValidBST(tt.tree)
			assert.Equalf(t, tt.want, got, "Is BST valid want %t, got %t", tt.want, got)
		})
	}
}
