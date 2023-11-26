package adt

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewStack(t *testing.T) {
	stack := NewStack[int]()
	assert.Equal(t, uint(0), stack.Size())
	assert.Equal(t, true, stack.IsEmpty())
}
