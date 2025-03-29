package cache

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Age uint8

func TestCache(t *testing.T) {
	a := assert.New(t)

	c := New()

	d := Age(18)

	c.Set("a", "a")
	c.Set("b", 1)
	c.Set("c", true)
	c.Set("d", d)

	a.Equal("a", c.Get("a"))
	a.Equal(1, c.Get("b"))
	a.Equal(true, c.Get("c"))
	a.Equal(d, c.Get("d"))

	c.Delete("a")
	c.Delete("b")
	c.Delete("c")
	c.Delete("d")

	a.Nil(c.Get("a"))
	a.Nil(c.Get("b"))
	a.Nil(c.Get("c"))
	a.Nil(c.Get("d"))
}
