package binarytree

// Node is a Node in a Binary Tree
type Node struct {
  Left *Node
  Right *Node
  Key Comparable
  Value interface{}
}

// Return a new empty node
func NewNode() *Node {
  return &Node{ Left: nil, Right: nil }
}

// Return a new node with the supplied key and value
func NewNodeKeyValue(key Comparable, value interface{}) *Node {
  return &Node{ Left: nil, Right: nil, Key: key, Value: value }
}

// Return a new node that is a deep copy of this node and all its children
func (me *Node) Copy() *Node {
  newNode := NewNodeKeyValue(me.Key, me.Value)
  if me.Left != nil { newNode.Left = me.Left.Copy() } 
  if me.Right != nil { newNode.Right = me.Right.Copy() }
  return newNode 
}

// Find and return the node with the supplied key in this subtree. Return nil if not found.
func (me *Node) Find(key Comparable) *Node {
  for me!=nil {
    if key.EqualTo(me.Key) { return me }
    if key.LessThan(me.Key) {
      me = me.Left
    } else {
      me = me.Right
    }
  }
  return nil
}

// Find and return the nearest node to the supplied key and its path to root.
// 1. If the node is found and it is the root node, return: node, []
// 2. If the node is found and it is not the root node, return: node, [node..., root]
// 2. If the node is not found, return: nearestNode, [node..., root]
func (me *Node) FindNearest(key Comparable) (*Node, []*Node) {
  stack := []*Node{}
  for {
    if me.Key.EqualTo(key) { return me, stack }
    if key.LessThan(me.Key) {
      if me.Left == nil { return me, stack }
      stack = append(stack, me) 
      me = me.Left
    } else {
      if me.Right == nil { return me, stack }
      stack = append(stack, me)
      me = me.Right
    }
  }
}

// Find and return the node with the largest key smaller than the supplied key, i.e.
// the next smallest node. If there is no smaller node, return nil.
func (me *Node) Previous(key Comparable) *Node {
  node, stack := me.FindNearest(key)
  if !node.Key.EqualTo(key) {
    if node.Key.LessThan(key) { return node }
    if len(stack) == 0 { return nil }
    for i:=len(stack)-1; i>=0; i-- {
      if stack[i].Key.LessThan(key) { return stack[i] }
    }
    return nil
  }
  if node.Left == nil {
    if len(stack) == 0 { return nil }
    for i:=len(stack)-1; i>=0; i-- {
      if stack[i].Key.LessThan(key) { return stack[i] }
    }
    return nil
  }
  return node.Left.Maximum()
}

// Find and return the node with the smallest key larger than the supplied key, i.e.
// the next largest node. If there is no larger node, return nil.
func (me *Node) Next(key Comparable) *Node {
  node, stack := me.FindNearest(key)
  if !node.Key.EqualTo(key) {
    if node.Key.GreaterThan(key) { return node }
    if len(stack) == 0 { return nil }
    for i:=len(stack)-1; i>=0; i-- {
      if stack[i].Key.GreaterThan(key) { return stack[i] }
    }
    // return nil        // This actually isn't possible.
  }
  if node.Right == nil {
    if len(stack) == 0 { return nil }
    for i:=len(stack)-1; i>=0; i-- {
      if stack[i].Key.GreaterThan(key) { return stack[i] }
    }
    return nil
  }
  return node.Right.Minimum()
}

// Add an existing node to this node's subtree
func (me *Node) Add(node *Node) *Node {
  current := me
  for {
    if node.Key.LessThan(current.Key) {
      if current.Left == nil {
        current.Left = node
        return node
      }
      current = current.Left
    } else {
      if current.Right == nil {
        current.Right = node
        return node
      }
      current = current.Right
    }
  }
}

// Remove a node from this node's subtree
func (me *Node) Remove(key Comparable) *Node {
  if me.Key.EqualTo(key) {
    // We are the node being removed
    // Leaf node. Return nil
    if me.Left == nil && me.Right == nil { return nil }
    // Right exists only. Return right
    if me.Left == nil && me.Right != nil { return me.Right }
    // Left exists only. Return left
    if me.Left != nil && me.Right == nil { return me.Left }
    // Left and right both exist. Add right to left and return left.
    oldMe := me
    me = me.Left
    oldMe.Left = nil
    if oldMe.Right!=nil { me.Add(oldMe.Right) }
    oldMe.Left = nil
    oldMe.Right = nil
  } else {
    // Walk the tree recursively calling Remove, set
    // each side to the return of Remove.
    if key.LessThan(me.Key) {
      if me.Left != nil {
        me.Left = me.Left.Remove(key)
      } 
    } else {
      if me.Right != nil {
        me.Right = me.Right.Remove(key)
      } 
    }
  }
  return me
}

// Balance this node's subtree, returning the new root node.
func (me *Node) Balance() *Node {
  var steps int = (me.DepthRight() - me.DepthLeft())/2
  for steps != 0 {
    if steps > 0 {
      oldMe := me
      me = me.Right
      oldMe.Right = nil
      me.Add(oldMe)
      steps--
    } else {
      oldMe := me
      me = me.Left
      oldMe.Left = nil
      me.Add(oldMe)
      steps++
    }
  }
  if me.Left!=nil { me.Left = me.Left.Balance() }
  if me.Right!=nil { me.Right = me.Right.Balance() }
  return me
}

// Call iterator for each node in this node's subtree in order, low to high
func (me *Node) WalkForward(iterator func(me *Node)) {
  if me.Left!=nil { me.Left.WalkForward(iterator) }
  iterator(me)
  if me.Right!=nil { me.Right.WalkForward(iterator) }
}

// Call iterator for each node in this node's subtree in reverse order, high to low
func (me *Node) WalkBackward(iterator func(me *Node)) {
  if me.Right!=nil { me.Right.WalkBackward(iterator) }
  iterator(me)
  if me.Left!=nil { me.Left.WalkBackward(iterator) }
}

// Call iterator for each node with a key in the range from, to in this node's subtree in order, low to high
//
// BUG: Inefficient. Walks whole tree.
func (me *Node) WalkRangeForward(iterator func(me *Node), from Comparable, to Comparable) {
  if me.Left!=nil { me.Left.WalkRangeForward(iterator, from, to) }
  if !me.Key.LessThan(from) && !me.Key.GreaterThan(to) { iterator(me) }
  if me.Right!=nil { me.Right.WalkRangeForward(iterator, from, to) }
}

// Call iterator for each node with a key in the range from, to in this node's subtree in reverse order, high to low
//
// BUG: Inefficient. Walks whole tree.
func (me *Node) WalkRangeBackward(iterator func(me *Node), from Comparable, to Comparable) {
  if me.Right!=nil { me.Right.WalkRangeBackward(iterator, from, to) }
  if !me.Key.LessThan(from) && !me.Key.GreaterThan(to) { iterator(me) }
  if me.Left!=nil { me.Left.WalkRangeBackward(iterator, from, to) }
}

// Return the left-most (smallest key) node in this node's subtree
func (me *Node) Minimum() *Node {
  for {
    if me.Left == nil { return me }
    me = me.Left
  }
}

// Return the right-most (largest key) node in this node's subtree
func (me *Node) Maximum() *Node {
  for {
    if me.Right == nil { return me }
    me = me.Right
  }
}

// Return the subtree depth to the left
func (me *Node) DepthLeft() int {
  x := 0
  for me!=nil { me = me.Left; x++ }
  return x-1
}

// Return the subtree depth to the right
func (me *Node) DepthRight() int {
  x := 0
  for me!=nil { me = me.Right; x++ }
  return x-1
}
