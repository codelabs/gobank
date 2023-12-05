package adt

import (
	"errors"
	"fmt"
)

// Queue like Stack, is also an abstract data structure.
// The thing that makes queue different from stack is that a queue is open at both its ends.
// Hence, it follows FIFO (First-In-First-Out) structure, i.e. the data item inserted first will also be accessed first.
// The data is inserted into the queue through one end and deleted from it using the other end.
type Queue[T any] struct {
	front *Node[T]
	rear  *Node[T]
	next  *Node[T]
	size  uint
}

// NewQueue creates an empty queue.
func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		front: nil,
		rear:  nil,
		next:  nil,
	}
}

// IsEmpty checks if queue is empty or not.
func (q *Queue[T]) IsEmpty() bool {
	return q.rear == nil && q.front == nil && q.size == 0
}

// Size returns the size of the queue.
func (q *Queue[T]) Size() uint {
	return q.size
}

// Display prints the queue to the console
func (q *Queue[T]) Display() {
	if q.IsEmpty() {
		fmt.Println("queue is empty")
		return
	}

	node := q.front
	for {
		node.Display()
		if node.next == nil {
			break
		}
		node = node.next
	}
}

// Front returns the front node in the queue, w/o deleting it.
// If Queue is empty, then it will return nil.
func (q *Queue[T]) Front() T {
	if q.IsEmpty() {
		var nothing T
		return nothing
	}
	return q.front.data
}

// Rear returns the rear node in the queue, w/o deleting it.
// If Queue is empty, then it will return nil.
func (q *Queue[T]) Rear() T {
	if q.IsEmpty() {
		var nothing T
		return nothing
	}
	return q.rear.data
}

// EnQueue is a data manipulation operation that inserts an element into the queue.
func (q *Queue[T]) EnQueue(data T) {

	// node to insert
	node := NewNode[T](data)

	// If queue is empty, both front and rear point to same node.
	if q.IsEmpty() {
		q.front = node
		q.rear = node
		q.size++
		return
	}

	node.next = q.front
	q.front = node
	q.size++
}

// DeQueue is a data manipulation operation that removes an element from the queue and returns it.
// If the queue is empty, it sets error.
func (q *Queue[T]) DeQueue() (T, error) {

	if q.IsEmpty() {
		var nothing T
		return nothing, errors.New("queue is empty")
	}

	// traverse to penultimate node
	node := q.front
	count := 1
	for {
		if uint(count) == q.Size()-1 {
			break
		}
		node = node.next
		count++
	}

	// Grab the last node and unlink it from queue, by pointing rear to
	// penultimate node
	lastNode := node.next
	node.next = nil
	q.rear = node
	q.size--
	return lastNode.Data(), nil
}
