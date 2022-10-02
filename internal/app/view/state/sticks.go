package state

type Sticks struct {
	Up        [4]bool
	HasThrown bool
}

func (s Sticks) Steps() int {
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

func SticksFromSteps(steps int, hasThrown bool) Sticks {
	return Sticks{
		Up: [4]bool{
			steps >= 1 && steps < 6,
			steps >= 2 && steps < 6,
			steps >= 3 && steps < 6,
			steps >= 4 && steps < 6,
		},
		HasThrown: hasThrown,
	}
}
