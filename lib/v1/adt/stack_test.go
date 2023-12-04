package adt

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewStack(t *testing.T) {
	t.Run("empty stack", func(t *testing.T) {
		stack := NewStack[int]()
		assert.Equal(t, uint(0), stack.Size())
		assert.Equal(t, true, stack.IsEmpty())
		stack.Display()
	})

	t.Run("push to stack", func(t *testing.T) {
		stack := NewStack[int]()
		stack.Push(10)
		assert.Equal(t, 10, stack.Top())
		stack.Push(20)
		stack.Display()
		assert.Equal(t, uint(2), stack.Size())
		assert.Equal(t, 20, stack.Top())
	})

	t.Run("pop from stack", func(t *testing.T) {
		stack := NewStack[int]()
		stack.Push(10)
		stack.Push(20)
		stack.Push(30)
		stack.Display()
		assert.Equal(t, uint(3), stack.Size())

		popped, err := stack.Pop()
		assert.Nil(t, err)
		assert.Equal(t, 30, popped.Data())
		assert.Equal(t, uint(2), stack.Size())
		stack.Display()
	})

	t.Run("pop from stack - empty stack", func(t *testing.T) {
		stack := NewStack[int]()
		assert.Equal(t, uint(0), stack.Size())
		stack.Display()

		popped, err := stack.Pop()
		assert.Nil(t, popped)
		if assert.Error(t, err) {
			assert.Equal(t, "stack is empty", err.Error())
		}
	})
}
