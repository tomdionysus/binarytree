package binarytree

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestNewNode(t *testing.T) {
  x := NewNode()

  assert.Nil(t, x.left)
  assert.Nil(t, x.right)
}

func TestNewNodeKeyValue(t *testing.T) {
  x := NewNodeKeyValue(IntKey(2),"Help!")

  assert.Equal(t, x.key, IntKey(2))
  assert.Equal(t, x.value, "Help!")
}

func TestFind(t *testing.T) {
  root := getTestTreeBalanced(1)

  assert.Equal(t, root.Find(IntKey(4)), root)
  assert.Equal(t, root.Find(IntKey(6)), root.right)
  assert.Equal(t, root.Find(IntKey(2)), root.left)
  assert.Equal(t, root.Find(IntKey(5)), root.right.left)
  assert.Equal(t, root.Find(IntKey(7)), root.right.right)
  assert.Equal(t, root.Find(IntKey(1)), root.left.left)
  assert.Equal(t, root.Find(IntKey(3)), root.left.right)

  assert.Nil(t, root.Find(IntKey(9)))
  assert.Nil(t, root.Find(IntKey(-1)))
}

func TestNextLessThan(t *testing.T) {
  root := getTestTreeBalanced(1)
  node := root.NextLessThan(IntKey(1))
  assert.Nil(t, node)
  node = root.NextLessThan(IntKey(2))
  assert.Equal(t, IntKey(1), node.key)
  node = root.NextLessThan(IntKey(3))
  assert.Equal(t, IntKey(2), node.key)
  node = root.NextLessThan(IntKey(20))
  assert.Equal(t, IntKey(7), node.key)

  root = getTestTreeLeftUnbalanced(1)
  node = root.NextLessThan(IntKey(1))
  assert.Nil(t, node)
  node = root.NextLessThan(IntKey(2))
  assert.Equal(t, IntKey(1), node.key)
  node = root.NextLessThan(IntKey(3))
  assert.Equal(t, IntKey(2), node.key)
  node = root.NextLessThan(IntKey(20))
  assert.Equal(t, IntKey(7), node.key)

  root = getTestTreeRightUnbalanced(1)
  node = root.NextLessThan(IntKey(1))
  assert.Nil(t, node)
  node = root.NextLessThan(IntKey(2))
  assert.Equal(t, IntKey(1), node.key)
  node = root.NextLessThan(IntKey(3))
  assert.Equal(t, IntKey(2), node.key)
  node = root.NextLessThan(IntKey(20))
  assert.Equal(t, IntKey(7), node.key)

  root = getTestTreeBalanced(2)
  node = root.NextLessThan(IntKey(1))
  assert.Nil(t, node)
  node = root.NextLessThan(IntKey(2))
  assert.Nil(t, node)
  node = root.NextLessThan(IntKey(5))
  assert.Equal(t, IntKey(4), node.key)
  node = root.NextLessThan(IntKey(20))
  assert.Equal(t, IntKey(14), node.key)
}

func TestNextGreaterThan(t *testing.T) {
  root := getTestTreeBalanced(1)
  node := root.NextGreaterThan(IntKey(7))
  assert.Nil(t, node)
  node = root.NextGreaterThan(IntKey(6))
  assert.Equal(t, IntKey(7), node.key)
  node = root.NextGreaterThan(IntKey(4))
  assert.Equal(t, IntKey(5), node.key)
  node = root.NextGreaterThan(IntKey(1))
  assert.Equal(t, IntKey(2), node.key)
}

func TestNodeAdd(t *testing.T) {
  x := NewNodeKeyValue(IntKey(5),"five")
  y := NewNodeKeyValue(IntKey(2),"two")
  z := NewNodeKeyValue(IntKey(7),"seven")
  q := NewNodeKeyValue(IntKey(9),"nine")

  x.Add(y)
  assert.Equal(t, x.left, y)
  assert.Nil(t, x.right)
  x.Add(z)
  assert.Equal(t, x.left, y)
  assert.Equal(t, x.right, z)
  x.Add(q)
  assert.Equal(t, x.left, y)
  assert.Equal(t, x.right, z)
  assert.Equal(t, x.right.right, q)
}

func TestRemove(t *testing.T) {
  // Remove only node
  root := NewNodeKeyValue(IntKey(2),"two")
  root = root.Remove(IntKey(2))
  assert.Nil(t, root)

  // Remove this node with right child
  root = NewNodeKeyValue(IntKey(2),"two")
  other := NewNodeKeyValue(IntKey(4), "four")
  root.Add(other)
  root = root.Remove(IntKey(2))
  assert.Equal(t, root, other)

  // Remove this node with left child
  root = NewNodeKeyValue(IntKey(2),"two")
  other = NewNodeKeyValue(IntKey(1), "one")
  root.Add(other)
  root = root.Remove(IntKey(2))
  assert.Equal(t, root, other)

  // Remove this node with both children child
  root = NewNodeKeyValue(IntKey(2),"two")
  other1 := NewNodeKeyValue(IntKey(1), "one")
  other4 := NewNodeKeyValue(IntKey(4), "four")
  root.Add(other1)
  root.Add(other4)

  root = root.Remove(IntKey(2))
  assert.Equal(t, root, other1)
  assert.Equal(t, root.right, other4)
  assert.Nil(t, root.left)
}

func TestBalance(t *testing.T) {
  root := getTestTreeRightUnbalanced(1)
  root = root.Balance()

  assert.Equal(t, IntKey(4), root.key)
  assert.Equal(t, IntKey(2), root.left.key)
  assert.Equal(t, IntKey(6), root.right.key)
  assert.Equal(t, IntKey(1), root.left.left.key)
  assert.Equal(t, IntKey(3), root.left.right.key)
  assert.Equal(t, IntKey(5), root.right.left.key)
  assert.Equal(t, IntKey(7), root.right.right.key)

  root = getTestTreeLeftUnbalanced(1)
  root = root.Balance()

  assert.Equal(t, IntKey(4), root.key)
  assert.Equal(t, IntKey(2), root.left.key)
  assert.Equal(t, IntKey(6), root.right.key)
  assert.Equal(t, IntKey(1), root.left.left.key)
  assert.Equal(t, IntKey(3), root.left.right.key)
  assert.Equal(t, IntKey(5), root.right.left.key)
  assert.Equal(t, IntKey(7), root.right.right.key)

  root = getTestTreeBalanced(1)
  root = root.Balance()

  assert.Equal(t, IntKey(4), root.key)
  assert.Equal(t, IntKey(2), root.left.key)
  assert.Equal(t, IntKey(6), root.right.key)
  assert.Equal(t, IntKey(1), root.left.left.key)
  assert.Equal(t, IntKey(3), root.left.right.key)
  assert.Equal(t, IntKey(5), root.right.left.key)
  assert.Equal(t, IntKey(7), root.right.right.key)
}

// Internals

func Testleftmost(t *testing.T) {
  root := getTestTreeBalanced(1)
  assert.Equal(t, root.rightmost().key, IntKey(1))

  root = getTestTreeLeftUnbalanced(1)
  assert.Equal(t, root.rightmost().key, IntKey(1))

  root = getTestTreeRightUnbalanced(1)
  assert.Equal(t, root.rightmost().key, IntKey(1))
}

func Testrightmost(t *testing.T) {
  root := getTestTreeBalanced(1)
  assert.Equal(t, root.rightmost().key, IntKey(7))

  root = getTestTreeLeftUnbalanced(1)
  assert.Equal(t, root.rightmost().key, IntKey(7))

  root = getTestTreeRightUnbalanced(1)
  assert.Equal(t, root.rightmost().key, IntKey(7))
}

func TestcountLeft(t *testing.T) {
  root := getTestTreeRightUnbalanced(1)
  assert.Equal(t, 0, root.countLeft())

  root = getTestTreeLeftUnbalanced(1)
  assert.Equal(t, 6, root.countLeft())

  root = getTestTreeBalanced(1)
  assert.Equal(t, 2, root.countLeft())
}

func TestcountRight(t *testing.T) {
  root := getTestTreeRightUnbalanced(1)
  assert.Equal(t, 6, root.countRight())

  root = getTestTreeLeftUnbalanced(1)
  assert.Equal(t, 0, root.countRight())

  root = getTestTreeBalanced(1)
  assert.Equal(t, 2, root.countRight())
}

func TestWalkForward(t *testing.T) {
  root := getTestTreeBalanced(1)

  out := []string{}

  root.WalkForward(func(me *Node) {
    out = append(out,me.value.(string))
  })

  assert.Equal(t, out, []string{
    "one",
    "two",
    "three",
    "four",
    "five",
    "six",
    "seven",
  })
}

func TestWalkBackward(t *testing.T) {
  root := getTestTreeBalanced(1)

  out := []string{}

  root.WalkBackward(func(me *Node) {
    out = append(out,me.value.(string))
  })

  assert.Equal(t, out, []string{
    "seven",
    "six",
    "five",
    "four",
    "three",
    "two",
    "one",
  })
}


// Helpers

func getTestTreeRightUnbalanced(factor int) *Node {
  root := NewNodeKeyValue(IntKey(1*factor), "one")
  root.Add(NewNodeKeyValue(IntKey(2*factor), "two"))
  root.Add(NewNodeKeyValue(IntKey(3*factor), "three"))
  root.Add(NewNodeKeyValue(IntKey(4*factor), "four"))
  root.Add(NewNodeKeyValue(IntKey(5*factor), "five"))
  root.Add(NewNodeKeyValue(IntKey(6*factor), "six"))
  root.Add(NewNodeKeyValue(IntKey(7*factor), "seven"))
  return root
}

func getTestTreeLeftUnbalanced(factor int) *Node {
  root := NewNodeKeyValue(IntKey(7*factor),"seven")
  root.Add(NewNodeKeyValue(IntKey(6*factor),"six"))
  root.Add(NewNodeKeyValue(IntKey(5*factor),"five"))
  root.Add(NewNodeKeyValue(IntKey(4*factor),"four"))
  root.Add(NewNodeKeyValue(IntKey(3*factor),"three"))
  root.Add(NewNodeKeyValue(IntKey(2*factor),"two"))
  root.Add(NewNodeKeyValue(IntKey(1*factor),"one"))
  return root
}

func getTestTreeBalanced(factor int) *Node {
  root := getTestTreeLeftUnbalanced(factor)
  return root.Balance()
}
