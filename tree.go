package binarytree

type Tree struct {
  root *Node
}

func NewTree() *Tree {
  return &Tree{ root: nil }
}

func (me *Tree) Set(key Comparable, value interface{}) {
  if me.root == nil {
    me.root = NewNodeKeyValue(nil, key, value)
  } else {
    node := me.root.Find(key)
    if node == nil {
      me.root.Add(NewNodeKeyValue(nil, key, value))
    } else {
      node.value = value
    }
  }
}

func (me *Tree) Get(key Comparable) (bool, interface{}) {
  node := me.GetNode(key)
  return node != nil, node
}

func (me *Tree) GetNode(key Comparable) *Node {
  if me.root == nil { return nil }
  return me.root.Find(key)
}

func (me *Tree) Copy() *Tree {
  newTree := NewTree()
  newTree.root = me.root
  if me.root == nil {
    return newTree
  }
  newTree.root = me.root.Copy()
  return newTree
}

