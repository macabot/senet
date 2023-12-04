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

func (s Sticks) ups() [4]bool {
	return [4]bool{
		s.Flips[0]%2 != 0,
		s.Flips[1]%2 != 0,
		s.Flips[2]%2 != 0,
		s.Flips[3]%2 != 0,
	}
}

func stepsToUps(steps int) [4]bool {
	switch steps {
	case 1, 2, 3:
		indices := [4]int{0, 1, 2, 3}
		defaultRNG.Shuffle(len(indices), func(i, j int) {
			indices[i], indices[j] = indices[j], indices[i]
		})
		ups := [4]bool{false, false, false, false}
		for i := 0; i < steps; i++ {
			ups[indices[i]] = true
		}
		return ups
	case 4:
		return [4]bool{true, true, true, true}
	case 6:
		return [4]bool{false, false, false, false}
	default:
		panic(fmt.Errorf("Cannot convert %d steps to ups.", steps))
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
		return TutorialSticksGenerator{}
	case CommitmentSchemeGeneratorKind:
		return CommitmentSchemeGenerator{}
	default:
		panic(fmt.Errorf("Cannot get sticks generator for kind %v.", s.GeneratorKind))
	}
}

func (s Sticks) CanThrow(state *State) bool {
	return s.generator().CanThrow(state)
}

func (s *Sticks) Throw(state *State) {
	steps := s.generator().Throw(state)
	s.SetSteps(steps)
	s.HasThrown = true
}

func (s *Sticks) SetSteps(steps int) {
	ups := stepsToUps(steps)
	s.Flips = [4]int{
		flipStick(s.Flips[0], ups[0]),
		flipStick(s.Flips[1], ups[1]),
		flipStick(s.Flips[2], ups[2]),
		flipStick(s.Flips[3], ups[3]),
	}
}

func (s Sticks) Steps() int {
	sum := 0
	for _, up := range s.ups() {
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
