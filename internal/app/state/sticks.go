package state

type Sticks struct {
	Flips     [4]int
	HasThrown bool
}

func (s Sticks) up() [4]bool {
	return [4]bool{
		s.Flips[0]%2 != 0,
		s.Flips[1]%2 != 0,
		s.Flips[2]%2 != 0,
		s.Flips[3]%2 != 0,
	}
}

func (s *Sticks) setFlipsFromUp(up [4]bool) {

}

func (s Sticks) Steps() int {
	sum := 0
	for _, up := range s.up() {
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
	sticks := Sticks{
		Flips:     [4]int{},
		HasThrown: hasThrown,
	}
	if steps >= 1 && steps < 6 {
		sticks.Flips[0] = 1
	}
	if steps >= 2 && steps < 6 {
		sticks.Flips[1] = 1
	}
	if steps >= 3 && steps < 6 {
		sticks.Flips[2] = 1
	}
	if steps >= 4 && steps < 6 {
		sticks.Flips[3] = 1
	}
	return sticks
}
