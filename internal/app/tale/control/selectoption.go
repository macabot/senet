package control

import "github.com/macabot/fairytale/control"

type Labeled[T any] struct {
	Label string
	V     T
}

type LabeledSlice[T any] []Labeled[T]

func (s LabeledSlice[T]) SelectOptions() []control.SelectOption[int] {
	options := make([]control.SelectOption[int], len(s))
	for i, item := range s {
		options[i] = control.SelectOption[int]{
			Label: item.Label,
			Value: i,
		}
	}
	return options
}
