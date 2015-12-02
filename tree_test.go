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

func TestGet(t *testing.T) {
  tree := NewTree()

  tree.Set(StringKey("one"),1)
  tree.Set(StringKey("two"),2)

  found, x := tree.Get(StringKey("one"))
  assert.True(t, found)
  assert.Equal(t, x.(int), 1)

  found, x = tree.Get(StringKey("two"))
  assert.True(t, found)
  assert.Equal(t, x.(int), 2)

  found, x = tree.Get(StringKey("three"))
  assert.False(t, found)
  assert.Nil(t, x)
}

func TestWalk(t *testing.T) {
  tree := NewTree()

  tree.Set(StringKey("one"),1)
  tree.Set(StringKey("two"),2)

  outkeys := []string{}
  outvalues := []int{}

  tree.Walk(func(key Comparable, value interface{}) {
    outkeys = append(outkeys, key.ValueOf().(string))
    outvalues = append(outvalues, value.(int))
  }, true)

  assert.Equal(t, outkeys, []string{"one","two"})
  assert.Equal(t, outvalues, []int{1,2})
}