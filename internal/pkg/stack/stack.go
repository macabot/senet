package stack

type Stack[T any] struct {
	values []T
}

func NewStack[T any](values ...T) *Stack[T] {
	return &Stack[T]{values: values}
}

func (s *Stack[T]) Push(values ...T) {
	s.values = append(s.values, values...)
}

func (s *Stack[T]) Pop() T {
	v := s.values[len(s.values)-1]
	s.values = s.values[:len(s.values)-1]
	return v
}

func (s *Stack[T]) Len() int {
	return len(s.values)
}
