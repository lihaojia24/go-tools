package xtools

import (
	"math/rand"
)

type Skiplist struct {
	head *node
}

type node struct {
	nexts []*node
	key   int
	val   int
}

func NewSkiplist() *Skiplist {
	return &Skiplist{
		head: &node{},
	}
}

func (s *Skiplist) Get(key int) (int, bool) {
	if _node := s.search(key); _node != nil {
		return _node.val, true
	}
	return -1, false
}

func (s *Skiplist) search(key int) *node {
	node := s.head
	for level := len(s.head.nexts) - 1; level > -1; level-- {
		for node.nexts[level] != nil && node.nexts[level].key < key {
			node = node.nexts[level]
		}
		if node.nexts[level] != nil && node.nexts[level].key == key {
			return node.nexts[level]
		}
	}
	return nil
}

func (s *Skiplist) Put(key, val int) {
	if _node := s.search(key); _node != nil {
		_node.val = val
		return
	}
	level := getLevel()
	for len(s.head.nexts)-1 < level {
		s.head.nexts = append(s.head.nexts, nil)
	}
	newNode := node{
		key:   key,
		val:   val,
		nexts: make([]*node, level+1),
	}
	node := s.head
	for l := len(s.head.nexts) - 1; l > -1; l-- {
		for node.nexts[l] != nil && node.nexts[l].key < key {
			node = node.nexts[l]
		}
		if l <= level {
			newNode.nexts[l] = node.nexts[l]
			node.nexts[l] = &newNode
		}
	}
}

func (s *Skiplist) Del(key int) bool {
	if _node := s.search(key); _node == nil {
		return false
	}
	node := s.head
	for level := len(s.head.nexts) - 1; level > -1; level-- {
		for node.nexts[level] != nil && node.nexts[level].key < key {
			node = node.nexts[level]
		}
		if node.nexts[level] != nil && node.nexts[level].key == key {
			node.nexts[level] = node.nexts[level].nexts[level]
		}
	}
	dif := 0
	for level := len(s.head.nexts); level > -1 && s.head.nexts[level] == nil; level-- {
		dif++
	}
	s.head.nexts = s.head.nexts[:len(s.head.nexts)-dif]
	return true
}

func (s *Skiplist) Ceiling(key int) (int, int, bool) {
	if ceilNode := s.ceiling(key); ceilNode != nil {
		return ceilNode.key, ceilNode.val, true
	}
	return -1, -1, false
}

func (s *Skiplist) ceiling(key int) *node {
	node := s.head
	for level := len(s.head.nexts) - 1; level > -1; level-- {
		for node.nexts[level] != nil && node.nexts[level].key < key {
			node = node.nexts[level]
		}
		if node.nexts[level] != nil && node.nexts[level].key == key {
			return node.nexts[level]
		}
	}
	return node.nexts[0]
}

func (s *Skiplist) Floor(key int) (int, int, bool) {
	if floorNode := s.floor(key); floorNode != s.head {
		return floorNode.key, floorNode.val, true
	}
	return -1, -1, false
}

func (s *Skiplist) Range(start, end int) ([]int, []int) {
	keys := []int{}
	vals := []int{}
	ceilNode := s.ceiling(start)
	if ceilNode == nil {
		return keys, vals
	}
	for ; ceilNode != nil && ceilNode.key <= end; ceilNode = ceilNode.nexts[0] {
		keys = append(keys, ceilNode.key)
		vals = append(vals, ceilNode.val)
	}
	return keys, vals
}

func (s *Skiplist) floor(key int) *node {
	node := s.head
	for level := len(s.head.nexts) - 1; level > -1; level-- {
		for node.nexts[level] != nil && node.nexts[level].key < key {
			node = node.nexts[level]
		}
		if node.nexts[level] != nil && node.nexts[level].key == key {
			return node.nexts[level]
		}
	}
	return node
}

func getLevel() (level int) {
	for rand.Intn(2) == 1 {
		level++
	}
	return
}
