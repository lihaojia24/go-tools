package tests

import (
	"fmt"
	"testing"

	"github.com/lihaojia24/go-xtools"
)

func Test_UFSet(t *testing.T) {
	s := xtools.NewUFSet(5)
	fmt.Printf("s.Find(1): %v\n", s.Find(1))
	fmt.Printf("s.Find(1): %v\n", s.Find(2))
	s.Union(1, 2)
	s.Union(3, 4)
	fmt.Printf("s.IsUnion(3, 4): %v\n", s.IsUnion(3, 4))
	fmt.Printf("s.IsUnion(1, 4): %v\n", s.IsUnion(1, 4))
	s.Union(2, 3)
	fmt.Printf("s.IsUnion(1, 4): %v\n", s.IsUnion(1, 4))
	fmt.Printf("s.Find(4): %v\n", s.Find(4))
}
