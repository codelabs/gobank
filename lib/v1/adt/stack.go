package adt

type Node[T any] struct {
	data T
	next *Node[T]
}

type Stack[T any] struct {
	top  *Node[T]
	size uint
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{top: nil}
}

func (s *Stack[T]) Size() uint {
	return s.size
}

func (s *Stack[T]) IsEmpty() bool {
	return s.top == nil && s.Size() == 0
}

func (s *Stack[T]) Top() T {
	return s.top.data
}

func (s *Stack[T]) Push(data T) {
}

func (s *Stack[T]) Pop() T {
	return s.top.data
}

func (s *Stack[T]) Display() {}
