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

  outkeys = []string{}
  outvalues = []int{}

  tree.Walk(func(key Comparable, value interface{}) {
    outkeys = append(outkeys, key.ValueOf().(string))
    outvalues = append(outvalues, value.(int))
  }, false)

  assert.Equal(t, outkeys, []string{"two","one"})
  assert.Equal(t, outvalues, []int{2,1})
}

func TestWalkRange(t *testing.T) {
  tree := NewTree()
  tree.root = getTestTreeBalanced(1)

  // Test off tree (left) -> 3
  outkeys := []int{}
  outvalues := []string{}

  tree.WalkRange(func(key Comparable, value interface{}) {
    outkeys = append(outkeys, key.ValueOf().(int))
    outvalues = append(outvalues, value.(string))
  }, IntKey(-5), IntKey(3), true)

  assert.Equal(t, []int{1,2,3}, outkeys)
  assert.Equal(t, []string{"one","two","three"}, outvalues)

  // Test 5 -> off tree (right)
  outkeys = []int{}
  outvalues = []string{}

  tree.WalkRange(func(key Comparable, value interface{}) {
    outkeys = append(outkeys, key.ValueOf().(int))
    outvalues = append(outvalues, value.(string))
  }, IntKey(5), IntKey(10), true)

  assert.Equal(t, outkeys, []int{5,6,7})
  assert.Equal(t, outvalues, []string{"five","six","seven"})

  // Test 2 -> 6
  outkeys = []int{}
  outvalues = []string{}

  tree.WalkRange(func(key Comparable, value interface{}) {
    outkeys = append(outkeys, key.ValueOf().(int))
    outvalues = append(outvalues, value.(string))
  }, IntKey(2), IntKey(6), true)

  assert.Equal(t, outkeys, []int{2,3,4,5,6})
  assert.Equal(t, outvalues, []string{"two","three","four","five","six"})
}

func TestWalkRangeScale(t *testing.T) {
  tree := NewTree()
  tree.root = getTestTreeBalanced(3)

  // Test off tree (left) -> 19
  outkeys := []int{}

  tree.WalkRange(func(key Comparable, value interface{}) {
    outkeys = append(outkeys, key.ValueOf().(int))
  }, IntKey(-5), IntKey(19), true)

  assert.Equal(t, []int{3,6,9,12,15,18}, outkeys)

  // Test 7 -> off tree (right)
  outkeys = []int{}

  tree.WalkRange(func(key Comparable, value interface{}) {
    outkeys = append(outkeys, key.ValueOf().(int))
  }, IntKey(7), IntKey(22), true)

  assert.Equal(t, []int{9,12,15,18,21}, outkeys)

  // Test 5 -> 10
  outkeys = []int{}

  tree.WalkRange(func(key Comparable, value interface{}) {
    outkeys = append(outkeys, key.ValueOf().(int))
  }, IntKey(5), IntKey(10), true)

  assert.Equal(t, []int{6,9}, outkeys)
}