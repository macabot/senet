package state

import (
	crand "crypto/rand"
	"encoding/binary"
	"fmt"
	"math/rand"
)

type SticksGeneratorKind int

const (
	CryptoSticksGeneratorKind SticksGeneratorKind = iota
	TutorialSticksGeneratorKind
)

type ThrowSticksGenerator interface {
	// Throw returns 1, 2, 3, 4 or 6 steps.
	Throw() int
}

var validSteps = [...]int{1, 2, 3, 4, 6}

var _ ThrowSticksGenerator = &CryptoSticksGenerator{}

type CryptoSticksGenerator struct {
	rng *rand.Rand
}

func NewCryptoSticksGenerator(rng *rand.Rand) *CryptoSticksGenerator {
	return &CryptoSticksGenerator{
		rng: rng,
	}
}

var defaultCryptoSticksGenerator = NewCryptoSticksGenerator(defaultRNG)

func (g CryptoSticksGenerator) Throw() int {
	return validSteps[g.rng.Intn(len(validSteps))]
}

// cryptoSource is based on https://yourbasic.org/golang/crypto-rand-int/
type cryptoSource struct{}

func (s cryptoSource) Seed(seed int64) {}

func (s cryptoSource) Int63() int64 {
	return int64(s.Uint64() & ^uint64(1<<63))
}

func (s cryptoSource) Uint64() (v uint64) {
	err := binary.Read(crand.Reader, binary.BigEndian, &v)
	if err != nil {
		panic(err)
	}
	return v
}

var defaultRNG = rand.New(cryptoSource{})

var _ ThrowSticksGenerator = &TutorialSticksGenerator{}

type TutorialSticksGenerator struct {
	index int
}

func NewTutorialSticksGenerator(index int) *TutorialSticksGenerator {
	return &TutorialSticksGenerator{
		index: index,
	}
}

var defaultTutorialSticksGenerator = NewTutorialSticksGenerator(0)

// tutorialThrownSticks must only contain values 1, 2, 3, 4 or 6.
var tutorialThrownSticks = [...]int{
	2,
}

func (g *TutorialSticksGenerator) Throw() int {
	steps := tutorialThrownSticks[g.index%len(tutorialThrownSticks)]
	g.index++
	return steps
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
