package set

type Set[T comparable] map[T]struct{}

func New[T comparable](v ...T) Set[T] {
	s := map[T]struct{}{}
	for _, x := range v {
		s[x] = struct{}{}
	}
	return s
}

func (s Set[T]) Has(v T) bool {
	if s == nil {
		return false
	}
	_, ok := s[v]
	return ok
}
