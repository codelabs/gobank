package linkedlist

import (
	"errors"
	"fmt"
)

// Node represents a node in a linked list or tree or graphs
type Node struct {
	data int
	next *Node
}

// Print prints a node.
func (n *Node) Print() {
	fmt.Printf("%d | %v\n", n.data, n.next)
}

// Single represents a single linked list.
type Single struct {
	size int
	head *Node
}

// NewSingleLinkedList creates an empty single linked list.
func NewSingleLinkedList() *Single {
	return &Single{head: nil}
}

// Length returns the length of the list.
func (s *Single) Length() int {
	return s.size
}

// IsEmpty checks if the list is empty or not.
func (s *Single) IsEmpty() bool {
	return s.head == nil && s.Length() == 0
}

// incrementLength increases the size of the list by 1.
func (s *Single) incrementLength() {
	s.size++
}

// decrementLength decreases the size of the list by 1.
func (s *Single) decrementLength() {
	s.size--
}

// Insert inserts the data to the end of the list.
// This is opposite to InsertAtBegin, which inserts at beginning.
func (s *Single) Insert(data int) {
	node := &Node{
		data: data,
		next: nil,
	}

	// First element
	if s.IsEmpty() {
		s.head = node
		s.incrementLength()
		return
	}

	// traverse to last element
	lastNode := s.head
	for lastNode.next != nil {
		lastNode = lastNode.next
	}

	// insert
	lastNode.next = node
	s.incrementLength()
}

// Print prints list to the console.
func (s *Single) Print() {
	if s.head == nil {
		fmt.Println("list is empty")
		return
	}

	fmt.Printf("H -> %v\n", s.head)
	node := s.head
	for {
		node.Print()
		if node.next == nil {
			break
		}
		node = node.next
	}
	fmt.Println("")
}

// InsertAtBeginning inserts the data to the beginning of the list.
// This is opposite to Insert, which inserts at the end.
func (s *Single) InsertAtBeginning(data int) {

	// When list is empty, we can just use Insert.
	if s.IsEmpty() {
		s.Insert(data)
		return
	}

	currentNode := s.head
	newNode := &Node{
		data: data,
		next: currentNode,
	}

	s.head = newNode
	s.incrementLength()
}

// InsertAt inserts the data into a specified position.
// Returns error, if the specified position is invalid. Position 1 is considered as beginning of the list.
func (s *Single) InsertAt(data int, position int) error {

	if position > s.Length() {
		return errors.New("invalid position, exceeded length of the list")
	}

	// Insert at beginning if position is 1
	if position == 1 {
		s.InsertAtBeginning(data)
		return nil
	}

	// To insert at specific position, we have to traverse to that position and insert
	node := s.head
	index := 1

	for {

		if index == position-1 {
			break
		}

		node = node.next
		index++
	}

	newNode := &Node{
		data: data,
		next: node.next,
	}
	node.next = newNode
	s.incrementLength()
	return nil
}

// Search searches for the specified data in the list.
func (s *Single) Search(data int) bool {
	if s.IsEmpty() {
		return false
	}

	node := s.head
	found := false
	for {

		if node.data == data {
			found = true
			break
		}

		if node.next == nil {
			break
		}

		node = node.next
	}

	return found
}

// DeleteAtEnd removes the last node from the list and returns it.
func (s *Single) DeleteAtEnd() *Node {

	// If list is empty, there is nothing to delete
	if s.IsEmpty() {
		return nil
	}

	// If there is only one element in the list, then set head to nil
	if s.Length() == 1 {
		node := s.head
		s.head = nil
		s.decrementLength()
		return node
	}

	// If the list has more than 1 node, traverse to the penultimate node in the list.
	node := s.head
	count := 1

	for {
		if count == s.Length()-1 {
			break
		}
		node = node.next
		count++
	}

	deletedNode := node.next
	node.next = nil
	s.decrementLength()
	return deletedNode
}
