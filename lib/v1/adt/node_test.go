package adt

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNode(t *testing.T) {
	t.Run("new string node", func(t *testing.T) {
		node := NewNode[string]("abc")
		node.Display()
		assert.Equal(t, "abc", node.Data())
	})

	t.Run("new int node", func(t *testing.T) {
		node := NewNode[int](10)
		node.Display()
		assert.Equal(t, 10, node.Data())
	})
}
