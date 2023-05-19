package state

import (
	"math/rand"
)

type Sticks struct {
	Flips     [4]int
	HasThrown bool
}

func (s *Sticks) Clone() *Sticks {
	if s == nil {
		return nil
	}
	return &Sticks{
		Flips:     [4]int{s.Flips[0], s.Flips[1], s.Flips[2], s.Flips[3]},
		HasThrown: s.HasThrown,
	}
}

func NewSticks() *Sticks {
	return &Sticks{}
}

func (s Sticks) up() [4]bool {
	return [4]bool{
		s.Flips[0]%2 != 0,
		s.Flips[1]%2 != 0,
		s.Flips[2]%2 != 0,
		s.Flips[3]%2 != 0,
	}
}

const minFlips = 3

func flipStick(flips int) int {
	sign := 1
	if rand.Float32() < 0.5 {
		sign = -1
	}
	flips += minFlips * sign
	up := rand.Float32() < 0.5
	if (flips%2 == 0) == up {
		flips += sign
	}
	return flips
}

func (s Sticks) Throw() *Sticks {
	return &Sticks{
		Flips: [4]int{
			flipStick(s.Flips[0]),
			flipStick(s.Flips[1]),
			flipStick(s.Flips[2]),
			flipStick(s.Flips[3]),
		},
		HasThrown: true,
	}
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

func (s Sticks) CanGoAgain() bool {
	steps := s.Steps()
	return steps == 1 || steps == 4 || steps == 6
}

func SticksFromSteps(steps int, hasThrown bool) *Sticks {
	sticks := &Sticks{
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
