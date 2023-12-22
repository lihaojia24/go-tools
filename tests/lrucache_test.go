package tests

import (
	"fmt"
	"testing"

	"github.com/lihaojia24/go-xtools"
)

func Test_lrucache(t *testing.T) {
	c := xtools.NewLRUCache(2)
	c.Put(1, 1)
	c.Put(2, 2)
	c.Put(3, 3)
	fmt.Println(c.Get(1))
	fmt.Println(c.Get(3))
	fmt.Println(c.Get(2))
	c.Put(1, 1)
	fmt.Println(c.Get(1))
	fmt.Println(c.Get(3))
	fmt.Println(c.Get(2))
}
