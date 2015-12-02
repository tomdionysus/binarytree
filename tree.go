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
  if node == nil {
    return false, nil
  }
  return true, node.value
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

// Return the first (lowest) key and value in the tree, or nil, nil if the tree is empty.
func (me *Tree) First() (Comparable, interface{}) {
   if me.root == nil { return nil, nil }
   node := me.root.leftmost()
   return node.key, node.value
}

// Return the last (highest) key and value in the tree, or nil, nil if the tree is empty.
func (me *Tree) Last() (Comparable, interface{}) {
   if me.root == nil { return nil, nil }
   node := me.root.rightmost()
   return node.key, node.value
}

// Iterate the tree with the function in the supplied direction
func (me *Tree) Walk(iterator func(key Comparable, value interface{}), forward bool) {
  if me.root == nil { return }
  if forward {
    me.root.WalkForward(func(node *Node) { iterator(node.key, node.value)})
  } else {
    me.root.WalkBackward(func(node *Node) { iterator(node.key, node.value)})
  }
}


