package binarytree

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestIntKeyLessThan(t *testing.T) {
  x := IntKey(1)
  y := IntKey(2)
  z := IntKey(3)
  a := IntKey(-5)

  assert.True(t, x.LessThan(y))
  assert.True(t, x.LessThan(z))
  assert.False(t, x.LessThan(a))
  assert.False(t, x.LessThan(x))
}

func TestIntKeyEqualTo(t *testing.T) {
  x := IntKey(1)
  y := IntKey(2)
  z := IntKey(2)
  a := IntKey(-5)

  assert.True(t, y.EqualTo(z))
  assert.True(t, z.EqualTo(y))
  assert.False(t, x.EqualTo(y))
  assert.False(t, x.EqualTo(a))
  assert.False(t, y.EqualTo(a))
  assert.False(t, a.EqualTo(x))
}

func TestIntKeyGreaterThan(t *testing.T) {
  x := IntKey(1)
  y := IntKey(2)
  z := IntKey(2)
  a := IntKey(-5)

  assert.True(t, y.GreaterThan(x))
  assert.True(t, z.GreaterThan(x))
  assert.False(t, a.GreaterThan(x))
  assert.False(t, x.GreaterThan(x))
}

func TestStringKeyLessThan(t *testing.T) {
  a := StringKey("a")
  b := StringKey("b")
  c := StringKey("cd")
  d := StringKey("cda")
  e := StringKey("cdb")

  assert.True(t, a.LessThan(b))
  assert.True(t, b.LessThan(c))
  assert.True(t, c.LessThan(d))
  assert.True(t, d.LessThan(e))

  assert.False(t, b.LessThan(a))
  assert.False(t, c.LessThan(b))
  assert.False(t, d.LessThan(c))
  assert.False(t, e.LessThan(d))

  assert.False(t, e.LessThan(e))
}

func TestStringKeyEqualTo(t *testing.T) {
  a := StringKey("a")
  b := StringKey("b")
  c := StringKey("cda")
  d := StringKey("cda")
  e := StringKey("cdb")

  assert.False(t, a.EqualTo(b))
  assert.False(t, b.EqualTo(c))
  assert.False(t, d.EqualTo(e))
  assert.False(t, a.EqualTo(e))

  assert.True(t, c.EqualTo(d))
}

func TestStringKeyGreaterThan(t *testing.T) {
  a := StringKey("a")
  b := StringKey("b")
  c := StringKey("cd")
  d := StringKey("cda")
  e := StringKey("cdb")

  assert.False(t, a.GreaterThan(b))
  assert.False(t, b.GreaterThan(c))
  assert.False(t, c.GreaterThan(d))
  assert.False(t, d.GreaterThan(e))

  assert.True(t, b.GreaterThan(a))
  assert.True(t, c.GreaterThan(b))
  assert.True(t, d.GreaterThan(c))
  assert.True(t, e.GreaterThan(d))

  assert.False(t, e.GreaterThan(e))
}

func TestByteSliceKeyLessThan(t *testing.T) {
  a := ByteSliceKey{ 0x00,0x00,0x00 }
  b := ByteSliceKey{ 0x00,0x00,0x01 }

  c := ByteSliceKey{ 0xAA,0xEE,0xFF,0xFF }
  d := ByteSliceKey{ 0xAA,0xEE }

  assert.True(t, a.LessThan(b))
  assert.False(t, b.LessThan(a))

  assert.True(t, a.LessThan(c))
  assert.False(t, c.LessThan(a))

  assert.True(t, d.LessThan(a))
  assert.False(t, a.LessThan(d))

  assert.True(t, d.LessThan(c))
  assert.False(t, c.LessThan(d))

  assert.False(t, c.LessThan(c))
}

func TestByteSliceKeyEqualTo(t *testing.T) {
  a := ByteSliceKey{ 0x00,0x00,0x00 }

  b := ByteSliceKey{ 0xAA,0xEE,0xFF,0xFF }
  c := ByteSliceKey{ 0xAA,0xEE,0xFF,0xFF }

  d := ByteSliceKey{ 0xAA,0xEE,0x22 }

  assert.True(t, a.EqualTo(a))
  assert.True(t, b.EqualTo(c))
  assert.True(t, c.EqualTo(b))
  assert.False(t, a.EqualTo(b))
  assert.False(t, b.EqualTo(a))
  assert.False(t, a.EqualTo(d))
}

func TestByteSliceKeyGreaterThan(t *testing.T) {
  a := ByteSliceKey{ 0x00,0x00,0x00 }
  b := ByteSliceKey{ 0x00,0x00,0x01 }

  c := ByteSliceKey{ 0xAA,0xEE,0xFF,0xFF }
  d := ByteSliceKey{ 0xAA,0xEE }

  assert.True(t, b.GreaterThan(a))
  assert.False(t, a.GreaterThan(b))

  assert.False(t, a.GreaterThan(c))
  assert.True(t, c.GreaterThan(a))

  assert.True(t, a.GreaterThan(d))
  assert.False(t, d.GreaterThan(a))

  assert.True(t, c.GreaterThan(d))
  assert.False(t, d.GreaterThan(c))

  assert.False(t, c.GreaterThan(c))
}

