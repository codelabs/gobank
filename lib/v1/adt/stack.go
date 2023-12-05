package adt

import (
	"errors"
	"fmt"
)

// Stack represents the abstract data structure stack.
// It uses linked list representation to maintain the stack.
type Stack[T any] struct {
	top  *Node[T]
	size uint
}

// NewStack creates an empty stack.
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{top: nil}
}

// Size returns the size of the stack.
func (s *Stack[T]) Size() uint {
	return s.size
}

// IsEmpty checks if the stack is empty or not.
// This is called when performing any of Push, Pop and Display operations.
func (s *Stack[T]) IsEmpty() bool {
	return s.top == nil && s.Size() == 0
}

// Top returns the top element from the stack, w/o deleting it.
func (s *Stack[T]) Top() T {
	return s.top.data
}

// Push pushes new element to top of the stack.
func (s *Stack[T]) Push(data T) {

	node := NewNode[T](data)

	if s.IsEmpty() {
		s.top = node
		s.size++
		return
	}

	currentTop := s.top
	node.next = currentTop
	s.top = node
	s.size++
	return
}

// Pop removes the top node from the stack and returns it.
func (s *Stack[T]) Pop() (*Node[T], error) {

	if s.IsEmpty() {
		return nil, errors.New("stack is empty")
	}

	poppedNode := s.top
	s.top = poppedNode.next
	poppedNode.next = nil
	s.size--
	return poppedNode, nil
}

// Display prints the stack to the console.
func (s *Stack[T]) Display() {
	if s.IsEmpty() {
		fmt.Println("stack is empty")
		return
	}

	node := s.top
	for {
		node.Display()
		if node.next == nil {
			break
		}
		node = node.next
	}

	return
}
