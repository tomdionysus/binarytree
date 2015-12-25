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
  
  // Test simple set
  tree := NewTree()

  tree.Set(StringKey("one"),1)
  tree.Set(StringKey("two"),2)

  assert.Equal(t, tree.root.Find(StringKey("one")).Value, 1)
  assert.Equal(t, tree.root.Find(StringKey("two")).Value, 2)

  // Test replacement set
  tree = NewTree()

  tree.Set(StringKey("one"),1)
  tree.Set(StringKey("one"),2)

  assert.Equal(t, tree.root.Find(StringKey("one")).Value, 2)
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

func TestClear(t *testing.T) {
  tree := NewTree()

  tree.Set(StringKey("one"),1)
  tree.Set(StringKey("two"),2)
  tree.Set(StringKey("three"),3)

  tree.Clear(StringKey("two"))

  found, x := tree.Get(StringKey("one"))
  assert.True(t, found)
  assert.Equal(t, x.(int), 1)

  found, x = tree.Get(StringKey("two"))
  assert.False(t, found)
  assert.Nil(t, x)

  found, x = tree.Get(StringKey("three"))
  assert.True(t, found)
  assert.Equal(t, x.(int), 3)

  tree = NewTree()
  tree.Clear(StringKey("two"))

  tree = NewTree()
  tree.Set(StringKey("two"),2)
  tree.Clear(StringKey("five"))
  found, x = tree.Get(StringKey("two"))
  assert.True(t, found)
  assert.Equal(t, x.(int), 2)
}

func TestTreeCopy(t *testing.T) {
  // Ensure copied tree is independent
  tree := NewTree()

  tree.Set(StringKey("one"),1)
  tree.Set(StringKey("two"),2)

  tree2 := tree.Copy()

  tree2.Clear(StringKey("one"))

  found, value := tree.Get(StringKey("one"))
  assert.True(t, found)
  assert.Equal(t, value, 1)

  // If tree is empty, return new empty tree
  tree = NewTree()

  tree2 = tree.Copy()
  tree.Set(StringKey("one"),1)
  assert.NotEqual(t, tree, tree2)
}

func TestTreePrevious(t *testing.T) {
  tree := NewTree()

  found, key, value := tree.Previous(IntKey(1))
  assert.False(t, found)
  
  tree.root = getTestTreeBalanced(1)

  found, key, value = tree.Previous(IntKey(1))
  assert.False(t, found)

  found, key, value = tree.Previous(IntKey(7))
  assert.True(t, found)
  assert.Equal(t, key, IntKey(6))
  assert.Equal(t, value, "six")
}

func TestTreeNext(t *testing.T) {
  tree := NewTree()

  found, key, value := tree.Next(IntKey(1))
  assert.False(t, found)
  
  tree.root = getTestTreeBalanced(1)

  found, key, value = tree.Next(IntKey(7))
  assert.False(t, found)

  found, key, value = tree.Next(IntKey(0))
  assert.True(t, found)
  assert.Equal(t, key, IntKey(1))
  assert.Equal(t, value, "one")
}

func TestTreeFirst(t *testing.T) {
  tree := NewTree()

  key, value := tree.First()
  assert.Nil(t, key)
  assert.Nil(t, value)

  tree.root = getTestTreeBalanced(1)

  key, value =  tree.First()
  assert.Equal(t, key,IntKey(1))
  assert.Equal(t, value,"one")
}

func TestTreeLast(t *testing.T) {
  tree := NewTree()

  key, value := tree.Last()
  assert.Nil(t, key)
  assert.Nil(t, value)

  tree.root = getTestTreeBalanced(1)

  key, value =  tree.Last()
  assert.Equal(t, key,IntKey(7))
  assert.Equal(t, value,"seven")
}

func TestGetNode(t *testing.T) {
  tree := NewTree()

  node := tree.GetNode(StringKey("one"))
  assert.Nil(t, node)

  tree.Set(StringKey("one"),1)
  tree.Set(StringKey("two"),2)
  tree.Set(StringKey("three"),3)

  node = tree.GetNode(StringKey("one"))
  assert.NotNil(t, node)
  assert.Equal(t, node.Value, 1)
  node = tree.GetNode(StringKey("two"))
  assert.NotNil(t, node)
  assert.Equal(t, node.Value, 2)
  node = tree.GetNode(StringKey("three"))
  assert.NotNil(t, node)
  assert.Equal(t, node.Value, 3)
}

func TestTreeBalance(t *testing.T) {
  tree := NewTree()

  tree.Balance()

  tree = NewTree()

  tree.Set(StringKey("one"),1)
  tree.Set(StringKey("two"),2)
  tree.Set(StringKey("three"),3)
  tree.Set(StringKey("four"),4)
  tree.Set(StringKey("five"),5)
  tree.Set(StringKey("six"),6)

  tree.Balance()

  found, value := tree.Get(StringKey("one"))
  assert.True(t, found); assert.Equal(t, value, 1)
  found, value = tree.Get(StringKey("two"))
  assert.True(t, found); assert.Equal(t, value, 2)
  found, value = tree.Get(StringKey("three"))
  assert.True(t, found); assert.Equal(t, value, 3)
  found, value = tree.Get(StringKey("four"))
  assert.True(t, found); assert.Equal(t, value, 4)
  found, value = tree.Get(StringKey("five"))
  assert.True(t, found); assert.Equal(t, value, 5)
  found, value = tree.Get(StringKey("six"))
  assert.True(t, found); assert.Equal(t, value, 6)
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

  // Don't call if tree empty
  tree = NewTree()

  tree.Walk(func(key Comparable, value interface{}) {
    assert.Equal(t,1,2)
  }, true)

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

  // Don't call if tree empty
  tree = NewTree()

  tree.WalkRange(func(key Comparable, value interface{}) {
    assert.Equal(t,1,2)
  }, IntKey(2), IntKey(6), true)
}

func TestWalkRangeBackward(t *testing.T) {
  tree := NewTree()
  tree.root = getTestTreeBalanced(1)

  // Test off tree (left) -> 3
  outkeys := []int{}

  tree.WalkRange(func(key Comparable, value interface{}) {
    outkeys = append(outkeys, key.ValueOf().(int))
  }, IntKey(-5), IntKey(3), false)

  assert.Equal(t, []int{3,2,1}, outkeys)

  // Test 5 -> off tree (right)
  outkeys = []int{}

  tree.WalkRange(func(key Comparable, value interface{}) {
    outkeys = append(outkeys, key.ValueOf().(int))
  }, IntKey(5), IntKey(10), false)

  assert.Equal(t, outkeys, []int{7,6,5})

  // Test 2 -> 6
  outkeys = []int{}

  tree.WalkRange(func(key Comparable, value interface{}) {
    outkeys = append(outkeys, key.ValueOf().(int))
  }, IntKey(2), IntKey(6), false)

  assert.Equal(t, outkeys, []int{6,5,4,3,2})
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

func TestWalkRangeBackwardScale(t *testing.T) {
  tree := NewTree()
  tree.root = getTestTreeBalanced(3)

  // Test off tree (left) -> 19
  outkeys := []int{}

  tree.WalkRange(func(key Comparable, value interface{}) {
    outkeys = append(outkeys, key.ValueOf().(int))
  }, IntKey(-5), IntKey(19), false)

  assert.Equal(t, []int{18,15,12,9,6,3}, outkeys)

  // Test 7 -> off tree (right)
  outkeys = []int{}

  tree.WalkRange(func(key Comparable, value interface{}) {
    outkeys = append(outkeys, key.ValueOf().(int))
  }, IntKey(7), IntKey(22), false)

  assert.Equal(t, []int{21,18,15,12,9}, outkeys)

  // Test 5 -> 10
  outkeys = []int{}

  tree.WalkRange(func(key Comparable, value interface{}) {
    outkeys = append(outkeys, key.ValueOf().(int))
  }, IntKey(5), IntKey(10), false)

  assert.Equal(t, []int{9,6}, outkeys)
}