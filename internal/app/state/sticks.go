package state

import (
	"fmt"
	"math/rand"
)

type Sticks struct {
	Flips         [4]int
	HasThrown     bool
	GeneratorKind SticksGeneratorKind
}

func (s *Sticks) Clone() *Sticks {
	if s == nil {
		return nil
	}
	return &Sticks{
		Flips:         [4]int{s.Flips[0], s.Flips[1], s.Flips[2], s.Flips[3]},
		HasThrown:     s.HasThrown,
		GeneratorKind: s.GeneratorKind,
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

func flipStick(flips int, up bool) int {
	sign := 1
	if rand.Float32() < 0.5 {
		sign = -1
	}
	flips += minFlips * sign
	if (flips%2 == 0) == up {
		flips += sign
	}
	return flips
}

func (s Sticks) generator() ThrowSticksGenerator {
	switch s.GeneratorKind {
	case CryptoSticksGeneratorKind:
		return defaultCryptoSticksGenerator
	case TutorialSticksGeneratorKind:
		return defaultTutorialSticksGenerator
	default:
		panic(fmt.Errorf("Cannot get sticks generator for kind %v.", s.GeneratorKind))
	}
}

func (s Sticks) Throw() *Sticks {
	steps := s.generator().Throw()
	return s.WithSteps(steps, true)
}

func (s Sticks) WithSteps(steps int, hasThrown bool) *Sticks {
	ups := stepsToUps(steps)
	return &Sticks{
		Flips: [4]int{
			flipStick(s.Flips[0], ups[0]),
			flipStick(s.Flips[1], ups[1]),
			flipStick(s.Flips[2], ups[2]),
			flipStick(s.Flips[3], ups[3]),
		},
		HasThrown:     hasThrown,
		GeneratorKind: s.GeneratorKind,
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
