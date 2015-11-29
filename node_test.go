package binarytree

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

type TestKey int

func (me TestKey) LessThan(other Comparable) bool {
  return me < other.(TestKey)
} 

func (me TestKey) Equal(other Comparable) bool {
  return me == other.(TestKey)
} 

func TestNewNode(t *testing.T) {
  p := &Node{}
  x := NewNode(p)

  assert.Equal(t, x.parent, p)
  assert.Nil(t, x.left)
  assert.Nil(t, x.right)
}

func TestNewNodeKeyValue(t *testing.T) {
  p := &Node{}
  x := NewNodeKeyValue(p, TestKey(2),"Help!")

  assert.Equal(t, x.parent, p)
  assert.Equal(t, x.key, TestKey(2))
  assert.Equal(t, x.value, "Help!")
}

func TestFind(t *testing.T) {
  root := getTestTreeBalanced()

  assert.Equal(t, root.Find(TestKey(4)), root)
  assert.Equal(t, root.Find(TestKey(6)), root.right)
  assert.Equal(t, root.Find(TestKey(2)), root.left)
  assert.Equal(t, root.Find(TestKey(5)), root.right.left)
  assert.Equal(t, root.Find(TestKey(7)), root.right.right)
  assert.Equal(t, root.Find(TestKey(1)), root.left.left)
  assert.Equal(t, root.Find(TestKey(3)), root.left.right)

  assert.Nil(t, root.Find(TestKey(9)))
  assert.Nil(t, root.Find(TestKey(-1)))
}


func TestNodeAdd(t *testing.T) {
  x := NewNodeKeyValue(nil, TestKey(5),"five")
  y := NewNodeKeyValue(nil, TestKey(2),"two")
  z := NewNodeKeyValue(nil, TestKey(7),"seven")
  q := NewNodeKeyValue(nil, TestKey(9),"nine")

  x.Add(y)
  assert.Equal(t, x.left, y )
  assert.Nil(t, x.right)
  x.Add(z)
  assert.Equal(t, x.left, y )
  assert.Equal(t, x.right, z )
  x.Add(q)
  assert.Equal(t, x.left, y )
  assert.Equal(t, x.right, z )
  assert.Equal(t, x.right.right, q )
}

func TestRemove(t *testing.T) {
  // Remove only node
  root := NewNodeKeyValue(nil, TestKey(2),"two")
  root = root.Remove(TestKey(2))
  assert.Nil(t, root)

  // Remove this node with right child
  root = NewNodeKeyValue(nil, TestKey(2),"two")
  other := NewNodeKeyValue(nil, TestKey(4), "four")
  root.Add(other)
  root = root.Remove(TestKey(2))
  assert.Equal(t, root, other)

  // Remove this node with left child
  root = NewNodeKeyValue(nil, TestKey(2),"two")
  other = NewNodeKeyValue(nil, TestKey(1), "one")
  root.Add(other)
  root = root.Remove(TestKey(2))
  assert.Equal(t, root, other)

  // Remove this node with both children child
  root = NewNodeKeyValue(nil, TestKey(2),"two")
  other1 := NewNodeKeyValue(nil, TestKey(1), "one")
  other4 := NewNodeKeyValue(nil, TestKey(4), "four")
  root.Add(other1)
  root.Add(other4)

  root = root.Remove(TestKey(2))
  assert.Equal(t, root, other1)
  assert.Equal(t, root.right, other4)
  assert.Nil(t, root.left)
}

func TestBalance(t *testing.T) {
  root := getTestTreeRightUnbalanced()
  root = root.Balance()

  assert.Equal(t, TestKey(4), root.key)
  assert.Equal(t, TestKey(2), root.left.key)
  assert.Equal(t, TestKey(6), root.right.key)
  assert.Equal(t, TestKey(1), root.left.left.key)
  assert.Equal(t, TestKey(3), root.left.right.key)
  assert.Equal(t, TestKey(5), root.right.left.key)
  assert.Equal(t, TestKey(7), root.right.right.key)

  root = getTestTreeLeftUnbalanced()
  root = root.Balance()

  assert.Equal(t, TestKey(4), root.key)
  assert.Equal(t, TestKey(2), root.left.key)
  assert.Equal(t, TestKey(6), root.right.key)
  assert.Equal(t, TestKey(1), root.left.left.key)
  assert.Equal(t, TestKey(3), root.left.right.key)
  assert.Equal(t, TestKey(5), root.right.left.key)
  assert.Equal(t, TestKey(7), root.right.right.key)

  root = getTestTreeBalanced()
  root = root.Balance()

  assert.Equal(t, TestKey(4), root.key)
  assert.Equal(t, TestKey(2), root.left.key)
  assert.Equal(t, TestKey(6), root.right.key)
  assert.Equal(t, TestKey(1), root.left.left.key)
  assert.Equal(t, TestKey(3), root.left.right.key)
  assert.Equal(t, TestKey(5), root.right.left.key)
  assert.Equal(t, TestKey(7), root.right.right.key)
}

// Internals

func Testleftmost(t *testing.T) {
  root := getTestTreeBalanced()
  assert.Equal(t, root.rightmost().key, TestKey(1))

  root = getTestTreeLeftUnbalanced()
  assert.Equal(t, root.rightmost().key, TestKey(1))

  root = getTestTreeRightUnbalanced()
  assert.Equal(t, root.rightmost().key, TestKey(1))
}

func Testrightmost(t *testing.T) {
  root := getTestTreeBalanced()
  assert.Equal(t, root.rightmost().key, TestKey(7))

  root = getTestTreeLeftUnbalanced()
  assert.Equal(t, root.rightmost().key, TestKey(7))

  root = getTestTreeRightUnbalanced()
  assert.Equal(t, root.rightmost().key, TestKey(7))
}

func TestcountLeft(t *testing.T) {
  root := getTestTreeRightUnbalanced()
  assert.Equal(t, 0, root.countLeft())

  root = getTestTreeLeftUnbalanced()
  assert.Equal(t, 6, root.countLeft())

  root = getTestTreeBalanced()
  assert.Equal(t, 2, root.countLeft())
}

func TestcountRight(t *testing.T) {
  root := getTestTreeRightUnbalanced()
  assert.Equal(t, 6, root.countRight())

  root = getTestTreeLeftUnbalanced()
  assert.Equal(t, 0, root.countRight())

  root = getTestTreeBalanced()
  assert.Equal(t, 2, root.countRight())
}

// Helpers

func getTestTreeRightUnbalanced() *Node {
  root := NewNodeKeyValue(nil, TestKey(1),"one")
  root.Add(NewNodeKeyValue(nil, TestKey(2),"two"))
  root.Add(NewNodeKeyValue(nil, TestKey(3),"three"))
  root.Add(NewNodeKeyValue(nil, TestKey(4),"four"))
  root.Add(NewNodeKeyValue(nil, TestKey(5),"five"))
  root.Add(NewNodeKeyValue(nil, TestKey(6),"six"))
  root.Add(NewNodeKeyValue(nil, TestKey(7),"seven"))
  return root
}

func getTestTreeLeftUnbalanced() *Node {
  root := NewNodeKeyValue(nil, TestKey(7),"seven")
  root.Add(NewNodeKeyValue(nil, TestKey(6),"six"))
  root.Add(NewNodeKeyValue(nil, TestKey(5),"five"))
  root.Add(NewNodeKeyValue(nil, TestKey(4),"four"))
  root.Add(NewNodeKeyValue(nil, TestKey(3),"three"))
  root.Add(NewNodeKeyValue(nil, TestKey(2),"two"))
  root.Add(NewNodeKeyValue(nil, TestKey(1),"one"))
  return root
}

func getTestTreeBalanced() *Node {
  root := getTestTreeLeftUnbalanced()
  return root.Balance()
}
