package adt

import (
	"cmp"
	"fmt"
)

type Tree[T cmp.Ordered] struct {
	root  *TreeNode[T]
	level int
}

type TraversalType int

const (
	PreOrder TraversalType = iota
	InOrder
	PostOrder
)

// NewBinaryTree creates an empty tree.
func NewBinaryTree[T cmp.Ordered]() *Tree[T] {
	return &Tree[T]{root: nil}
}

func (t *Tree[T]) IsEmpty() bool {
	return t.root == nil
}

// Insert inserts data into the tree.
func (t *Tree[T]) Insert(data T) {

	node := NewTreeNode[T](data)

	// if tree is empty, point root to the node
	if t.IsEmpty() {
		t.root = node
		return
	}

	var (
		current              = t.root
		parent  *TreeNode[T] = nil
	)

	for {
		parent = current
		if node.data < parent.data {
			current = current.left
			if current == nil {
				parent.left = node
				return
			}
		} else {
			current = current.right
			if current == nil {
				parent.right = node
				return
			}
		}
	}
}

// Traverse prints the tree data based on TraversalType.
func (t *Tree[T]) Traverse(tType TraversalType) {

	if t.IsEmpty() {
		fmt.Println("tree is empty")
		return
	}

	switch tType {
	case InOrder:
		inOrder(t.root)
	case PreOrder:
		preOrder(t.root)
	case PostOrder:
		postOrder(t.root)
	}
}

func inOrder[T cmp.Ordered](root *TreeNode[T]) {
	if root != nil {
		inOrder[T](root.left)
		fmt.Printf("%v ", root.data)
		inOrder[T](root.right)
	}
}

func preOrder[T cmp.Ordered](root *TreeNode[T]) {
	if root != nil {
		fmt.Printf("%v ", root.data)
		preOrder[T](root.left)
		preOrder[T](root.right)
	}
}

func postOrder[T cmp.Ordered](root *TreeNode[T]) {
	if root != nil {
		postOrder[T](root.left)
		postOrder[T](root.right)
		fmt.Printf("%v ", root.data)
	}
}
