package utils

type Set[T comparable] map[T]bool

func NewSet[T comparable]() Set[T] {
	return make(Set[T])
}

func (s Set[T]) Add(elem T) {
	s[elem] = true
}

func (s Set[T]) Contains(elem T) bool {
	_, found := s[elem]
	return found
}

func (s Set[T]) Remove(elem T) {
	delete(s, elem)
}

func (s Set[T]) Size() int {
	return len(s)
}
