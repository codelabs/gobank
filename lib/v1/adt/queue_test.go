package adt

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQueue(t *testing.T) {
	t.Run("empty queue", func(t *testing.T) {
		queue := NewQueue[int]()
		queue.Display()
		assert.Equal(t, true, queue.IsEmpty())
		assert.Equal(t, uint(0), queue.Size())
		assert.Equal(t, 0, queue.Front())
		assert.Equal(t, 0, queue.Rear())
	})

	t.Run("add data to queue", func(t *testing.T) {
		queue := NewQueue[int]()
		queue.EnQueue(10)
		queue.EnQueue(20)
		queue.EnQueue(30)
		queue.Display()

		assert.Equal(t, 30, queue.Front())
		assert.Equal(t, 10, queue.Rear())
		assert.Equal(t, uint(3), queue.Size())
	})

	t.Run("remove data from queue", func(t *testing.T) {
		queue := NewQueue[int]()
		queue.EnQueue(10)
		queue.EnQueue(20)
		queue.EnQueue(30)
		t.Logf("Queue after Enqueue\n")
		queue.Display()
		assert.Equal(t, uint(3), queue.Size())
		assert.Equal(t, 10, queue.Rear())
		assert.Equal(t, 30, queue.Front())

		data, err := queue.DeQueue()
		assert.Nil(t, err)
		assert.Equal(t, 10, data)
		assert.Equal(t, uint(2), queue.Size())
		assert.Equal(t, 20, queue.Rear())
		assert.Equal(t, 30, queue.Front())
		t.Logf("Queue after Dequeue\n")
		queue.Display()
	})

	t.Run("remove data from empty queue", func(t *testing.T) {
		queue := NewQueue[int]()
		queue.Display()
		assert.Equal(t, uint(0), queue.Size())
		assert.Equal(t, true, queue.IsEmpty())

		_, err := queue.DeQueue()
		if assert.Error(t, err) {
			assert.Equal(t, "queue is empty", err.Error())
		}
	})
}
