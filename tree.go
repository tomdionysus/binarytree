package binarytree

// Tree represents a binary tree
type Tree struct {
  root *Node
}

// Return a new empty binary tree
func NewTree() *Tree {
  return &Tree{ root: nil }
}

// Add the supplied key and value to the tree. If the key already exists, the value will be overwritten.
func (me *Tree) Set(key Comparable, value interface{}) {
  if me.root == nil {
    me.root = NewNodeKeyValue(key, value)
  } else {
    node := me.root.Find(key)
    if node == nil {
      me.root.Add(NewNodeKeyValue(key, value))
    } else {
      node.value = value
    }
  }
}

// Get the value associated with the supplied key. Return (true, value) if found,
// (false, nil) if not.
func (me *Tree) Get(key Comparable) (bool, interface{}) {
  node := me.GetNode(key)
  return node != nil, node
}

// Get the node associated with the supplied key, or nil if not found
func (me *Tree) GetNode(key Comparable) *Node {
  if me.root == nil { return nil }
  return me.root.Find(key)
}

// Return a deep copy of the tree.
func (me *Tree) Copy() *Tree {
  newTree := NewTree()
  newTree.root = me.root
  if me.root == nil {
    return newTree
  }
  newTree.root = me.root.Copy()
  return newTree
}

// Balance the tree.
func (me *Tree) Balance() {
  if me.root == nil { return }
  me.root = me.root.Balance()
}

// Return the value associated with the next smallest key than the supplied key.
// If a smaller key exists, return (true, value), otherwise return (false, nil).
func (me *Tree) NextLessThan(key Comparable) (bool, interface{}) {
  if me.root == nil { return false, nil }
  node := me.root.NextLessThan(key)
  if node == nil { return false, nil }
  return true, node.value
}

// Return the value associated with the next largest key than the supplied key.
// If a larger key exists, return (true, value), otherwise return (false, nil).
func (me *Tree) NextGreaterThan(key Comparable) (bool, interface{}) {
  if me.root == nil { return false, nil }
  node := me.root.NextGreaterThan(key)
  if node == nil { return false, nil }
  return true, node.value
}
