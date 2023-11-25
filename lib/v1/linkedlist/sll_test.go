package linkedlist

import "testing"
import "github.com/stretchr/testify/assert"

func TestNewSingleLinkedList(t *testing.T) {
	t.Run("new empty list", func(t *testing.T) {
		list := NewSingleLinkedList()
		assert.Equal(t, 0, list.Length())
		assert.Equal(t, true, list.IsEmpty())
		list.Print()
	})

	t.Run("list with 5 elements", func(t *testing.T) {
		list := NewSingleLinkedList()
		list.Insert(10)
		list.Insert(20)
		list.Insert(30)
		list.Insert(40)
		list.Insert(50)

		assert.Equal(t, 5, list.Length())
		assert.Equal(t, false, list.IsEmpty())
		list.Print()
	})
}
