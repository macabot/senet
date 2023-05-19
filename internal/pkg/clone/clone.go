package clone

type Cloner[T any] interface {
	Clone() T
}

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
