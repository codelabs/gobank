package adt

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewBinaryTree(t *testing.T) {
	t.Run("empty tree", func(t *testing.T) {
		tree := NewBinaryTree[int]()
		assert.Equal(t, true, tree.IsEmpty())
	})

	t.Run("binary tree", func(t *testing.T) {
		tree := NewBinaryTree[int]()
		tree.Insert(50)
		tree.Insert(35)
		tree.Insert(65)
		tree.Insert(45)
		tree.Insert(40)

		tree.Traverse(InOrder)
		fmt.Println()
		tree.Traverse(PreOrder)
		fmt.Println()
		tree.Traverse(PostOrder)
		fmt.Println()
	})
}
