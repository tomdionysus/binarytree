package binarytree

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestNewTree(t *testing.T) {
  x := NewTree()

  assert.Nil(t, x.root)
}

func TestSet(t *testing.T) {
  tree := NewTree()

  tree.Set(StringKey("one"),1)
  tree.Set(StringKey("two"),2)

  assert.Equal(t, tree.root.Find(StringKey("one")).value, 1)
  assert.Equal(t, tree.root.Find(StringKey("two")).value, 2)

}