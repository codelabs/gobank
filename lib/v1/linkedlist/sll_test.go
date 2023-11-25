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

	t.Run("insert", func(t *testing.T) {
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

	t.Run("insert at beginning", func(t *testing.T) {
		list := NewSingleLinkedList()
		list.Insert(10)
		list.Insert(20)
		list.InsertAtBeginning(30)

		assert.Equal(t, 3, list.Length())
		list.Print()
	})

	t.Run("insert at position", func(t *testing.T) {
		list := NewSingleLinkedList()
		list.Insert(10)
		list.Insert(20)
		list.InsertAtBeginning(30)
		list.Print()

		err := list.InsertAt(40, 3)
		assert.Nil(t, err)

		list.Print()
		assert.Equal(t, 4, list.Length())

		err = list.InsertAt(50, 5)
		if assert.Error(t, err) {
			assert.Equal(t, "invalid position, exceeded length of the list", err.Error())
		}

	})

	t.Run("insert at position - empty list", func(t *testing.T) {
		list := NewSingleLinkedList()
		if list.IsEmpty() {
			err := list.InsertAt(10, 1)
			if assert.Error(t, err) {
				assert.Equal(t, "invalid position, exceeded length of the list", err.Error())
			}
		}
	})

	t.Run("insert at position - beginning", func(t *testing.T) {
		list := NewSingleLinkedList()
		list.Insert(10)
		list.Insert(20)
		list.InsertAtBeginning(30)
		list.Print()

		err := list.InsertAt(40, 1)
		assert.Nil(t, err)
		assert.Equal(t, 4, list.Length())

		list.Print()
	})

	t.Run("search an empty list", func(t *testing.T) {
		list := NewSingleLinkedList()
		assert.Equal(t, false, list.Search(10))
	})

	t.Run("search for nonexistence item", func(t *testing.T) {
		list := NewSingleLinkedList()
		for _, data := range []int{10, 20, 30, 40, 50} {
			list.Insert(data)
		}
		assert.Equal(t, false, list.Search(60))
	})

	t.Run("search for an existing item", func(t *testing.T) {
		list := NewSingleLinkedList()
		for _, data := range []int{10, 20, 30, 40, 50} {
			list.Insert(data)
		}
		assert.Equal(t, true, list.Search(30))
	})
}
