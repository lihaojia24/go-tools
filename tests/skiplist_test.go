package tests

import (
	"fmt"
	"testing"

	"github.com/lihaojia24/go-xtools"
)

func Test_skiplist(t *testing.T) {
	s := xtools.NewSkiplist()
	s.Put(3, 3)
	s.Put(6, 6)
	s.Put(7, 7)
	fmt.Println(s.Get(5))
	fmt.Println(s.Floor(6))
	fmt.Println(s.Ceiling(4))
	fmt.Println(s.Range(4, 7))
}
