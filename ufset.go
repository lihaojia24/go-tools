package xtools

type UFSet struct {
	size int
	m    []int
}

func NewUFSet(size int) *UFSet {
	m := make([]int, size)
	for i := range m {
		m[i] = -1
	}
	return &UFSet{size, m}
}

func (s *UFSet) Find(node int) int {
	if node >= s.size || node < 0 {
		return -1
	}
	for s.m[node] > 0 {
		node = s.m[node]
	}
	return node
}

func (s *UFSet) Union(a, b int) bool {
	if a >= s.size || a < 0 || b >= s.size || b < 0 {
		return false
	}
	br := s.Find(b)
	s.m[br] = a
	return true
}

func (s *UFSet) IsUnion(a, b int) bool {
	return s.Find(a) == s.Find(b)
}
