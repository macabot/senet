package clone

type Cloner[T any] interface {
	Clone() T
}

// Map clones a map. If a value implements Cloner, then the Clone function is called.
func Map[M ~map[K]V, K comparable, V any](m M) M {
	r := make(M, len(m))
	for k, v := range m {
		if c, ok := any(v).(Cloner[V]); ok {
			r[k] = c.Clone()
		} else {
			r[k] = v
		}
	}
	return r
}

// Slice clones a slice. If a value implements Cloner, then the Clone function is called.
func Slice[S ~[]E, E any](s S) S {
	if s == nil {
		return nil
	}
	r := make(S, len(s))
	for i, v := range s {
		if c, ok := any(v).(Cloner[E]); ok {
			r[i] = c.Clone()
		} else {
			r[i] = v
		}
	}
	return r
}
