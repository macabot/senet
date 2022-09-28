package state

type Sticks struct {
	Up        [4]bool
	HasThrown bool
}

func (s Sticks) Value() int {
	sum := 0
	for _, up := range s.Up {
		if up {
			sum++
		}
	}
	if sum == 0 {
		sum = 6
	}
	return sum
}

func SticksFromThrow(throw int, hasThrown bool) Sticks {
	return Sticks{
		Up: [4]bool{
			throw >= 1 && throw < 6,
			throw >= 2 && throw < 6,
			throw >= 3 && throw < 6,
			throw >= 4 && throw < 6,
		},
		HasThrown: hasThrown,
	}
}
