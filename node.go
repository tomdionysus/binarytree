package binarytree

type Comparable interface {
  LessThan(Comparable) bool
  Equal(Comparable) bool
}

type Node struct {
  parent *Node
  left *Node
  right *Node
  key Comparable
  value interface{}
}

func NewNode(parent *Node) *Node {
  return &Node{ parent: parent, left: nil, right: nil }
}

func NewNodeKeyValue(parent *Node, key Comparable, value interface{}) *Node {
  return &Node{ parent: parent, left: nil, right: nil, key: key, value: value }
}

func (me *Node) Copy() *Node {
  newNode := NewNodeKeyValue(me, me.key, me.value)
  if me.left != nil { newNode.left = me.left.Copy() } 
  if me.right != nil { newNode.right = me.right.Copy() }
  return newNode 
}

func (me *Node) Find(key Comparable) *Node {
  for me!=nil {
    if key.Equal(me.key) { return me }
    if key.LessThan(me.key) {
      me = me.left
    } else {
      me = me.right
    }
  }
  return nil
}

func (me *Node) FindNearest(key Comparable) (*Node, bool) {
  last := me
  for me!=nil {
    if key.Equal(me.key) { return me, true }
    if key.LessThan(me.key) {
      last = me
      me = me.left
    } else {
      last = me
      me = me.right
    }
  }
  return last, false
}

func (me *Node) Add(node *Node) *Node {
  current := me
  for {
    if node.key.LessThan(current.key) {
      if current.left == nil {
        current.left = node
        node.parent = current
        return node
      }
      current = current.left
    } else {
      if current.right == nil {
        current.right = node
        node.parent = current
        return node
      }
      current = current.right
    }
  }
}

func (me *Node) Remove(key Comparable) *Node {
  if me.key.Equal(key) {
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
    me.parent = oldMe.parent
    if oldMe.right!=nil { me.Add(oldMe.right) }
    oldMe.parent = nil
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

func (me *Node) NextGreaterThan(key Comparable) *Node {
  return nil
}

func (me *Node) NextLessThan(key Comparable) *Node {
  return nil
}

func (me *Node) Next() *Node {
  var closest, test *Node = nil,nil
  if me.right!=nil {
    test := me.right.leftmost()
    if closest==nil || test.key.LessThan(closest.key) {
      closest = test
    }
  }
  if me.parent!=nil {
    if closest==nil || test.key.LessThan(me.parent.key) {
      closest = me.parent
    }
  }

  return closest
}

func (me *Node) Previous() *Node {
  return nil
}

func (me *Node) Balance() *Node {
  steps := (me.countRight() - me.countLeft())/2
  for steps != 0 {
    if steps > 0 {
      oldMe := me
      me = me.right
      oldMe.right = nil
      me.parent = oldMe.parent
      me.Add(oldMe)
      steps--
    } else {
      oldMe := me
      me = me.left
      oldMe.left = nil
      me.parent = oldMe.parent
      me.Add(oldMe)
      steps++
    }
  }
  if me.left!=nil { me.left = me.left.Balance() }
  if me.right!=nil { me.right = me.right.Balance() }
  return me
}

func (me *Node) leftmost() *Node {
  for {
    if me.left == nil { return me }
    me = me.left
  }
}

func (me *Node) rightmost() *Node {
  for {
    if me.right == nil { return me }
    me = me.right
  }
}

func (me *Node) countLeft() int {
  x := 0
  for me!=nil { me = me.left; x++ }
  return x-1
}

func (me *Node) countRight() int {
  x := 0
  for me!=nil { me = me.right; x++ }
  return x-1
}
