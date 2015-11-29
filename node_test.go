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
  x := NewNode()

  assert.Nil(t, x.left)
  assert.Nil(t, x.right)
}

func TestNewNodeKeyValue(t *testing.T) {
  x := NewNodeKeyValue(TestKey(2),"Help!")

  assert.Equal(t, x.key, TestKey(2))
  assert.Equal(t, x.value, "Help!")
}

func TestNodeFind(t *testing.T) {
  root := getTestTree()

  assert.Equal(t, root.Find(TestKey(2)), root.left.left)
  assert.Equal(t, root.Find(TestKey(12)), root.right.left)
  assert.Equal(t, root.Find(TestKey(5)), root.left)
  assert.Equal(t, root.Find(TestKey(15)), root.right)
  assert.Equal(t, root.Find(TestKey(10)), root)

  assert.Nil(t, root.Find(TestKey(4)))
  assert.Nil(t, root.Find(TestKey(6)))
}


func TestNodeAdd(t *testing.T) {
  x := NewNodeKeyValue(TestKey(5),"five")
  y := NewNodeKeyValue(TestKey(2),"two")
  z := NewNodeKeyValue(TestKey(7),"seven")
  q := NewNodeKeyValue(TestKey(9),"nine")

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
  root := NewNodeKeyValue(TestKey(2),"two")
  root = root.Remove(TestKey(2))
  assert.Nil(t, root)

  // Remove this node with right child
  root = NewNodeKeyValue(TestKey(2),"two")
  other := NewNodeKeyValue(TestKey(4), "four")
  root.Add(other)
  root = root.Remove(TestKey(2))
  assert.Equal(t, root, other)

  // Remove this node with left child
  root = NewNodeKeyValue(TestKey(2),"two")
  other = NewNodeKeyValue(TestKey(1), "one")
  root.Add(other)
  root = root.Remove(TestKey(2))
  assert.Equal(t, root, other)

  // Remove this node with both children child
  root = NewNodeKeyValue(TestKey(2),"two")
  other1 := NewNodeKeyValue(TestKey(1), "one")
  other4 := NewNodeKeyValue(TestKey(4), "four")
  root.Add(other1)
  root.Add(other4)

  root = root.Remove(TestKey(2))
  assert.Equal(t, root, other1)
  assert.Equal(t, root.right, other4)
  assert.Nil(t, root.left)
}

func TestBalance(t *testing.T) {
  root := NewNodeKeyValue(TestKey(1),"one")
  root.Add(NewNodeKeyValue(TestKey(2),"two"))
  root.Add(NewNodeKeyValue(TestKey(3),"three"))
  root.Add(NewNodeKeyValue(TestKey(4),"four"))
  root.Add(NewNodeKeyValue(TestKey(5),"five"))
  root.Add(NewNodeKeyValue(TestKey(6),"six"))
  root.Add(NewNodeKeyValue(TestKey(7),"seven"))

  assert.Equal(t, 6, root.countRight())
  assert.Equal(t, 0, root.countLeft())

  root = root.Balance()

  assert.Equal(t, 2, root.countRight())
  assert.Equal(t, 2, root.countLeft())

  assert.Equal(t, 1, root.left.countRight())
  assert.Equal(t, 1, root.left.countLeft())

  assert.Equal(t, 1, root.right.countRight())
  assert.Equal(t, 1, root.right.countLeft())

  assert.Equal(t, TestKey(4), root.key)
  assert.Equal(t, TestKey(2), root.left.key)
  assert.Equal(t, TestKey(6), root.right.key)
  assert.Equal(t, TestKey(1), root.left.left.key)
  assert.Equal(t, TestKey(3), root.left.right.key)
  assert.Equal(t, TestKey(5), root.right.left.key)
  assert.Equal(t, TestKey(7), root.right.right.key)
}

func Testleftmost(t *testing.T) {
  root := getTestTree()

  assert.Equal(t, root.leftmost().key, TestKey(2))
}

func Testrightmost(t *testing.T) {
  root := getTestTree()

  assert.Equal(t, root.rightmost().key, TestKey(17))
}

func TestNextGreaterThan(t *testing.T) {
  root := getTestTree()

  assert.Equal(t, root.NextGreaterThan(TestKey(2)).key, TestKey(5))
  assert.Equal(t, root.NextGreaterThan(TestKey(12)).key, TestKey(15))
  assert.Nil(t, root.NextGreaterThan(TestKey(17)))

  root = NewNodeKeyValue(TestKey(7),"seven")
  root.Add(NewNodeKeyValue(TestKey(6),"six"))
  root.Add(NewNodeKeyValue(TestKey(5),"five"))
  root.Add(NewNodeKeyValue(TestKey(4),"four"))
  root.Add(NewNodeKeyValue(TestKey(3),"three"))
  root.Add(NewNodeKeyValue(TestKey(2),"two"))
  root.Add(NewNodeKeyValue(TestKey(1),"one"))

  assert.Equal(t, root.NextGreaterThan(TestKey(4)).key, TestKey(5))
  assert.Equal(t, root.NextGreaterThan(TestKey(1)).key, TestKey(2))
}


func getTestTree() *Node {
  root := NewNodeKeyValue(TestKey(10), "ten")
  root.left = NewNodeKeyValue(TestKey(5), "five")
  root.right = NewNodeKeyValue(TestKey(15), "fifteen")
  root.left.left = NewNodeKeyValue(TestKey(2), "two")
  root.left.right = NewNodeKeyValue(TestKey(7), "seven")
  root.right.left = NewNodeKeyValue(TestKey(12), "twelve")
  root.right.right = NewNodeKeyValue(TestKey(17), "seventeen")
  return root
}
