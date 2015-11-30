package binarytree

// Comparable is an interface for comparable types. All keys used in this implementation of
// a binary tree must implement this interface.
//
// Helper types are supplied for some primitives, please see `IntKey`, [StringKey] and [ByteSliceKey]
type Comparable interface {
  LessThan(Comparable) bool
  EqualTo(Comparable) bool
  GreaterThan(Comparable) bool
}

// IntKey is a type of base type int that implements the Comparable interface.
type IntKey int

// Return true if this key is less than the supplied IntKey.
func (me IntKey) LessThan(other Comparable) bool {
  return me < other.(IntKey)
} 

// Return true if this key is equal to the supplied IntKey.
func (me IntKey) EqualTo(other Comparable) bool {
  return me == other.(IntKey)
} 

// Return true if this key is greater than the supplied IntKey.
func (me IntKey) GreaterThan(other Comparable) bool {
  return me > other.(IntKey)
} 

// StringKey is a type of base type String that implements the Comparable interface.
type StringKey string

// Return true if this key is less than the supplied StringKey.
func (me StringKey) LessThan(other Comparable) bool {
  return me < other.(StringKey)
} 

// Return true if this key is equal to the supplied StringKey.
func (me StringKey) EqualTo(other Comparable) bool {
  return me == other.(StringKey)
} 

// Return true if this key is greater than the supplied StringKey.
func (me StringKey) GreaterThan(other Comparable) bool {
  return me > other.(StringKey)
} 

// StringKey is a type of base type String that implements the Comparable interface.
type ByteSliceKey []byte

// Return true if this key is less than the supplied StringKey.
func (me StringKey) LessThan(other Comparable) bool {
  if len(me) < len(other) { return true }
  for i, x := range me {
    if x<other[i] { return true }
  }
  return false
} 

// Return true if this key is equal to the supplied StringKey.
func (me StringKey) EqualTo(other Comparable) bool {
  if len(me) != len(other) { return false }
  for i, x := range me {
    if x!=other[i] { return false }
  }
  return true
} 

// Return true if this key is greater than the supplied StringKey.
func (me StringKey) GreaterThan(other Comparable) bool {
  if len(me) > len(other) { return true } 
  for i, x := range me {
    if x>other[i] { return true }
  }
  return false
} 

