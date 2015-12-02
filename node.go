package binarytree

// Node is a Node in a Binary Tree
type Node struct {
  left *Node
  right *Node
  key Comparable
  value interface{}
}

// Return a new empty node
func NewNode() *Node {
  return &Node{ left: nil, right: nil }
}

// Return a new node with the supplied key and value
func NewNodeKeyValue(key Comparable, value interface{}) *Node {
  return &Node{ left: nil, right: nil, key: key, value: value }
}

// Return a new node that is a deep copy of this node and all its children
func (me *Node) Copy() *Node {
  newNode := NewNodeKeyValue(me.key, me.value)
  if me.left != nil { newNode.left = me.left.Copy() } 
  if me.right != nil { newNode.right = me.right.Copy() }
  return newNode 
}

// Find and return the node with the supplied key in this subtree. Return nil if not found.
func (me *Node) Find(key Comparable) *Node {
  for me!=nil {
    if key.EqualTo(me.key) { return me }
    if key.LessThan(me.key) {
      me = me.left
    } else {
      me = me.right
    }
  }
  return nil
}

// Find and return the node with the largest key smaller than the supplied key, i.e.
// the next smallest node. If there is no smaller node, return nil.
//
// BUG(inefficient): This is seriously inefficient, it walks the tree until it finds a node smaller than the key.
func (me *Node) NextLessThan(key Comparable) *Node {
  return me.nextLessThan(key, nil)
}

func (me *Node) nextLessThan(key Comparable, best *Node) *Node {
  if me.right != nil { best = me.right.nextLessThan(key, best); if best!=nil { return best } }
  if me.key.LessThan(key) { return me }
  if me.left != nil { best = me.left.nextLessThan(key, best); if best!=nil { return best } }
  return best
}

// Find and return the node with the smallest key larger than the supplied key, i.e.
// the next largest node. If there is no larger node, return nil.
//
// BUG(inefficient): This is seriously inefficient, it walks the tree until it finds a node larger than the key.
func (me *Node) NextGreaterThan(key Comparable) *Node {
  return me.nextGreaterThan(key, nil)
}

func (me *Node) nextGreaterThan(key Comparable, best *Node) *Node {
  if me.left != nil { best = me.left.nextGreaterThan(key, best); if best!=nil { return best } }
  if me.key.GreaterThan(key) { return me }
  if me.right != nil { best = me.right.nextGreaterThan(key, best); if best!=nil { return best } }
  return best
}

// Add an existing node to this node's subtree
func (me *Node) Add(node *Node) *Node {
  current := me
  for {
    if node.key.LessThan(current.key) {
      if current.left == nil {
        current.left = node
        return node
      }
      current = current.left
    } else {
      if current.right == nil {
        current.right = node
        return node
      }
      current = current.right
    }
  }
}

// Remove a node from this node's subtree
func (me *Node) Remove(key Comparable) *Node {
  if me.key.EqualTo(key) {
    // We are the node being removed
    // Leaf node. Return nil
    if me.left == nil && me.right == nil { return nil }
    // Right exists only. Return right
    if me.left == nil && me.right != nil { return me.right }
    // Left exists only. Return left
    if me.left != nil && me.right == nil { return me.left }
    // Left and right both exist. Add right to left and return left.
    oldMe := me
    me = me.left
    oldMe.left = nil
    if oldMe.right!=nil { me.Add(oldMe.right) }
    oldMe.left = nil
    oldMe.right = nil
  } else {
    // Walk the tree recursively calling Remove, set
    // each side to the return of Remove.
    if key.LessThan(me.key) {
      if me.left != nil {
        me.left = me.left.Remove(key)
      } 
    } else {
      if me.right != nil {
        me.right = me.right.Remove(key)
      } 
    }
  }
  return me
}

// Balance this node's subtree, returning the new root node.
func (me *Node) Balance() *Node {
  var steps int = (me.countRight() - me.countLeft())/2
  for steps != 0 {
    if steps > 0 {
      oldMe := me
      me = me.right
      oldMe.right = nil
      me.Add(oldMe)
      steps--
    } else {
      oldMe := me
      me = me.left
      oldMe.left = nil
      me.Add(oldMe)
      steps++
    }
  }
  if me.left!=nil { me.left = me.left.Balance() }
  if me.right!=nil { me.right = me.right.Balance() }
  return me
}

// Call iterator for each node in this node's subtree in order, low to high
func (me *Node) WalkForward(iterator func(me *Node)) {
  if me.left!=nil { me.left.WalkForward(iterator) }
  iterator(me)
  if me.right!=nil { me.right.WalkForward(iterator) }
}

// Call iterator for each node in this node's subtree in reverse order, high to low
func (me *Node) WalkBackward(iterator func(me *Node)) {
  if me.right!=nil { me.right.WalkBackward(iterator) }
  iterator(me)
  if me.left!=nil { me.left.WalkBackward(iterator) }
}

// Return the left-most (smallest key) node in this node's subtree
func (me *Node) leftmost() *Node {
  for {
    if me.left == nil { return me }
    me = me.left
  }
}

// Return the right-most (largest key) node in this node's subtree
func (me *Node) rightmost() *Node {
  for {
    if me.right == nil { return me }
    me = me.right
  }
}

// Return the subtree depth to the left
func (me *Node) countLeft() int {
  x := 0
  for me!=nil { me = me.left; x++ }
  return x-1
}

// Return the subtree depth to the right
func (me *Node) countRight() int {
  x := 0
  for me!=nil { me = me.right; x++ }
  return x-1
}
