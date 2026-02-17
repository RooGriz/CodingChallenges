package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsSymmetric(t *testing.T) {
	tests := []struct {
		name string
		tree *TreeNode
		want bool
	}{
		{
			name: "SymmetricWhenEmptyTree",
			tree: nil,
			want: true,
		},
		{
			name: "SymmetricWhenTreeHasOneNode",
			tree: &TreeNode{
				Left:  nil,
				Right: nil,
				Val:   10,
			},
			want: true,
		},
		{
			name: "SymmetricWhenTreeHasSeveralNodes1",
			tree: &TreeNode{
				Left: &TreeNode{
					Left:  nil,
					Right: nil,
					Val:   11,
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
			name: "SymmetricWhenTreeHasSeveralNodes2",
			tree: &TreeNode{
				Left: &TreeNode{
					Left: &TreeNode{
						Left:  nil,
						Right: nil,
						Val:   9,
					},
					Right: &TreeNode{
						Left:  nil,
						Right: nil,
						Val:   12,
					},
					Val: 11,
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
						Val:   9,
					},
					Val: 11,
				},
				Val: 10,
			},
			want: true,
		},
		{
			name: "NotSymmetricWhenTreeHasSeveralNodes1",
			tree: &TreeNode{
				Left: nil,
				Right: &TreeNode{
					Left:  nil,
					Right: nil,
					Val:   9,
				},
				Val: 10,
			},
			want: false,
		},
		{
			name: "NotSymmetricWhenTreeHasSeveralNodes2",
			tree: &TreeNode{
				Left: &TreeNode{
					Left:  nil,
					Right: nil,
					Val:   9,
				},
				Right: nil,
				Val:   10,
			},
			want: false,
		},
		{
			name: "NotSymmetricWhenTreeHasSeveralNodes3",
			tree: &TreeNode{
				Left: &TreeNode{
					Left: &TreeNode{
						Left:  nil,
						Right: nil,
						Val:   9,
					},
					Right: &TreeNode{
						Left:  nil,
						Right: nil,
						Val:   12,
					},
					Val: 11,
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
						Val:   19,
					},
					Val: 11,
				},
				Val: 10,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isSymmetric(tt.tree)
			assert.Equalf(t, tt.want, got, "Is tree symmetric want %t, got %t", tt.want, got)
		})
	}
}
