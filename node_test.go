package binarytree

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestNewNode(t *testing.T) {
  x := NewNode()

  assert.Nil(t, x.Left)
  assert.Nil(t, x.Right)
}

func TestNewNodeKeyValue(t *testing.T) {
  x := NewNodeKeyValue(IntKey(2),"Help!")

  assert.Equal(t, x.Key, IntKey(2))
  assert.Equal(t, x.Value, "Help!")
}

func TestFind(t *testing.T) {
  root := getTestTreeBalanced(1)

  assert.Equal(t, root.Find(IntKey(4)), root)
  assert.Equal(t, root.Find(IntKey(6)), root.Right)
  assert.Equal(t, root.Find(IntKey(2)), root.Left)
  assert.Equal(t, root.Find(IntKey(5)), root.Right.Left)
  assert.Equal(t, root.Find(IntKey(7)), root.Right.Right)
  assert.Equal(t, root.Find(IntKey(1)), root.Left.Left)
  assert.Equal(t, root.Find(IntKey(3)), root.Left.Right)

  assert.Nil(t, root.Find(IntKey(9)))
  assert.Nil(t, root.Find(IntKey(-1)))
}

func TestPrevious(t *testing.T) {
  root := getTestTreeBalanced(1)
  node := root.Previous(IntKey(1))
  assert.Nil(t, node)
  node = root.Previous(IntKey(2))
  assert.Equal(t, IntKey(1), node.Key)
  node = root.Previous(IntKey(3))
  assert.Equal(t, IntKey(2), node.Key)
  node = root.Previous(IntKey(20))
  assert.Equal(t, IntKey(7), node.Key)

  root = getTestTreeLeftUnbalanced(1)
  node = root.Previous(IntKey(1))
  assert.Nil(t, node)
  node = root.Previous(IntKey(2))
  assert.Equal(t, IntKey(1), node.Key)
  node = root.Previous(IntKey(3))
  assert.Equal(t, IntKey(2), node.Key)
  node = root.Previous(IntKey(20))
  assert.Equal(t, IntKey(7), node.Key)

  root = getTestTreeRightUnbalanced(1)
  node = root.Previous(IntKey(1))
  assert.Nil(t, node)
  node = root.Previous(IntKey(2))
  assert.Equal(t, IntKey(1), node.Key)
  node = root.Previous(IntKey(3))
  assert.Equal(t, IntKey(2), node.Key)
  node = root.Previous(IntKey(20))
  assert.Equal(t, IntKey(7), node.Key)

  root = getTestTreeBalanced(2)
  node = root.Previous(IntKey(1))
  assert.Nil(t, node)
  node = root.Previous(IntKey(2))
  assert.Nil(t, node)
  node = root.Previous(IntKey(5))
  assert.Equal(t, IntKey(4), node.Key)
  node = root.Previous(IntKey(20))
  assert.Equal(t, IntKey(14), node.Key)
}

func TestNext(t *testing.T) {
  // Left Unbalanced
  root := getTestTreeLeftUnbalanced(1)
  node := root.Next(IntKey(0))
  assert.Equal(t, IntKey(1), node.Key)

  for i:=0; i<6; i++ {
    node = root.Next(IntKey(i))
    assert.NotNil(t, node)
    if node!=nil { assert.Equal(t, IntKey(i+1), node.Key) }
  }
  assert.Nil(t, root.Next(IntKey(7)))

  // Right Unbalanced
  root = getTestTreeRightUnbalanced(1)
  node = root.Next(IntKey(0))
  assert.Equal(t, IntKey(1), node.Key)

  for i:=0; i<6; i++ {
    node = root.Next(IntKey(i))
    assert.NotNil(t, node)
    if node!=nil { assert.Equal(t, IntKey(i+1), node.Key) }
  }
  assert.Nil(t, root.Next(IntKey(7)))

  // Balanced
  root = getTestTreeBalanced(1)
  assert.Equal(t, IntKey(1), root.Next(IntKey(0)).Key)
  assert.Equal(t, IntKey(2), root.Next(IntKey(1)).Key)
  assert.Equal(t, IntKey(3), root.Next(IntKey(2)).Key)
  assert.Equal(t, IntKey(4), root.Next(IntKey(3)).Key)
  assert.Equal(t, IntKey(5), root.Next(IntKey(4)).Key)
  assert.Equal(t, IntKey(6), root.Next(IntKey(5)).Key)
  assert.Equal(t, IntKey(7), root.Next(IntKey(6)).Key)
  assert.Nil(t, root.Next(IntKey(7)))

  // Left Unbalanced Scaled (3)
  root = getTestTreeLeftUnbalanced(3)
  node = root.Next(IntKey(0))
  assert.Equal(t, IntKey(3), node.Key)

  for i:=0; i<21; i++ {
    node = root.Next(IntKey(i))
    assert.NotNil(t, node)
    if node!=nil { assert.Equal(t, IntKey(((i/3)+1)*3), node.Key) }
  }
  assert.Nil(t, root.Next(IntKey(21)))

  // Right Unbalanced Scaled (3)
  root = getTestTreeRightUnbalanced(3)
  node = root.Next(IntKey(0))
  assert.Equal(t, IntKey(3), node.Key)

  for i:=0; i<21; i++ {
    node = root.Next(IntKey(i))
    assert.NotNil(t, node)
    if node!=nil { assert.Equal(t, IntKey(((i/3)+1)*3), node.Key) }
  }
  assert.Nil(t, root.Next(IntKey(21)))
}



func TestNodeAdd(t *testing.T) {
  x := NewNodeKeyValue(IntKey(5),"five")
  y := NewNodeKeyValue(IntKey(2),"two")
  z := NewNodeKeyValue(IntKey(7),"seven")
  q := NewNodeKeyValue(IntKey(9),"nine")

  x.Add(y)
  assert.Equal(t, x.Left, y)
  assert.Nil(t, x.Right)
  x.Add(z)
  assert.Equal(t, x.Left, y)
  assert.Equal(t, x.Right, z)
  x.Add(q)
  assert.Equal(t, x.Left, y)
  assert.Equal(t, x.Right, z)
  assert.Equal(t, x.Right.Right, q)
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
  assert.Equal(t, root.Right, other4)
  assert.Nil(t, root.Left)
}

func TestBalance(t *testing.T) {
  root := getTestTreeRightUnbalanced(1)
  root = root.Balance()

  assert.Equal(t, IntKey(4), root.Key)
  assert.Equal(t, IntKey(2), root.Left.Key)
  assert.Equal(t, IntKey(6), root.Right.Key)
  assert.Equal(t, IntKey(1), root.Left.Left.Key)
  assert.Equal(t, IntKey(3), root.Left.Right.Key)
  assert.Equal(t, IntKey(5), root.Right.Left.Key)
  assert.Equal(t, IntKey(7), root.Right.Right.Key)

  root = getTestTreeLeftUnbalanced(1)
  root = root.Balance()

  assert.Equal(t, IntKey(4), root.Key)
  assert.Equal(t, IntKey(2), root.Left.Key)
  assert.Equal(t, IntKey(6), root.Right.Key)
  assert.Equal(t, IntKey(1), root.Left.Left.Key)
  assert.Equal(t, IntKey(3), root.Left.Right.Key)
  assert.Equal(t, IntKey(5), root.Right.Left.Key)
  assert.Equal(t, IntKey(7), root.Right.Right.Key)

  root = getTestTreeBalanced(1)
  root = root.Balance()

  assert.Equal(t, IntKey(4), root.Key)
  assert.Equal(t, IntKey(2), root.Left.Key)
  assert.Equal(t, IntKey(6), root.Right.Key)
  assert.Equal(t, IntKey(1), root.Left.Left.Key)
  assert.Equal(t, IntKey(3), root.Left.Right.Key)
  assert.Equal(t, IntKey(5), root.Right.Left.Key)
  assert.Equal(t, IntKey(7), root.Right.Right.Key)
}

// Internals

func TestMinimum(t *testing.T) {
  root := getTestTreeBalanced(1)
  assert.Equal(t, root.Minimum().Key, IntKey(1))

  root = getTestTreeLeftUnbalanced(1)
  assert.Equal(t, root.Minimum().Key, IntKey(1))

  root = getTestTreeRightUnbalanced(1)
  assert.Equal(t, root.Minimum().Key, IntKey(1))
}

func TestMaximum(t *testing.T) {
  root := getTestTreeBalanced(1)
  assert.Equal(t, root.Maximum().Key, IntKey(7))

  root = getTestTreeLeftUnbalanced(1)
  assert.Equal(t, root.Maximum().Key, IntKey(7))

  root = getTestTreeRightUnbalanced(1)
  assert.Equal(t, root.Maximum().Key, IntKey(7))
}

func TestDepthLeft(t *testing.T) {
  root := getTestTreeRightUnbalanced(1)
  assert.Equal(t, 0, root.DepthLeft())

  root = getTestTreeLeftUnbalanced(1)
  assert.Equal(t, 6, root.DepthLeft())

  root = getTestTreeBalanced(1)
  assert.Equal(t, 2, root.DepthLeft())
}

func TestDepthRight(t *testing.T) {
  root := getTestTreeRightUnbalanced(1)
  assert.Equal(t, 6, root.DepthRight())

  root = getTestTreeLeftUnbalanced(1)
  assert.Equal(t, 0, root.DepthRight())

  root = getTestTreeBalanced(1)
  assert.Equal(t, 2, root.DepthRight())
}

func TestWalkForward(t *testing.T) {
  root := getTestTreeBalanced(1)

  out := []string{}

  root.WalkForward(func(me *Node) {
    out = append(out,me.Value.(string))
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
    out = append(out,me.Value.(string))
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
