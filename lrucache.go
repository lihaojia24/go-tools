package xtools

import "container/list"

type LRUCache struct {
	capacity int
	linkList *list.List
	key2Node map[int]*list.Element
}

type entry struct {
	key   int
	value int
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		linkList: list.New(),
		key2Node: make(map[int]*list.Element),
	}
}

func (c *LRUCache) Get(key int) (int, bool) {
	node := c.key2Node[key]
	if node == nil {
		return -1, false
	}
	c.linkList.MoveToFront(node)
	return node.Value.(entry).value, true
}

func (c *LRUCache) Put(key, val int) {
	node := c.key2Node[key]
	if node != nil {
		node.Value = entry{key, val}
		c.linkList.MoveToFront(node)
		return
	}
	c.key2Node[key] = c.linkList.PushFront(entry{key, val})
	if len(c.key2Node) > c.capacity {
		delete(c.key2Node, c.linkList.Remove(c.linkList.Back()).(entry).key)
	}
}
