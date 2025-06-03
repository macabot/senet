package state

import (
	crand "crypto/rand"
	"encoding/binary"
	"math/rand"
)

type SticksGeneratorKind int

const (
	CryptoSticksGeneratorKind SticksGeneratorKind = iota
	TutorialSticksGeneratorKind
	CommitmentSchemeGeneratorKind
)

type ThrowSticksGenerator interface {
	CanThrow(*State) bool
	// Throw returns 1, 2, 3, 4 or 6 steps.
	Throw(*State) int
}

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

func (g CryptoSticksGenerator) CanThrow(_ *State) bool {
	return true
}

func (g CryptoSticksGenerator) Throw(_ *State) int {
	sum := 0
	for i := 0; i < 4; i++ {
		sum += g.rng.Intn(2)
	}
	if sum == 0 {
		sum = 6
	}
	return sum
}

// cryptoSource is based on https://yourbasic.org/golang/crypto-rand-int/
type cryptoSource struct{}

func (s cryptoSource) Seed(seed int64) { /* on-op */ }

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

type TutorialSticksGenerator struct{}

// tutorialThrownSticks must only contain values 1, 2, 3, 4 or 6.
var tutorialThrownSticks = [...]int{
	2, // TutorialMove
	1, // TutorialTradingPlaces4
	2, // TutorialBlockingPiece2
	4, // TutorialReturnToStart3
	2, // TutorialMoveBackwards2
	3, // TutorialNoMove2
	6, // TutorialOffTheBoard2
	4, // TutorialOffTheBoard3
	1, // TutorialOffTheBoard3
}

func (g TutorialSticksGenerator) CanThrow(_ *State) bool {
	return true
}

func (g TutorialSticksGenerator) Throw(s *State) int {
	steps := tutorialThrownSticks[s.TutorialIndex%len(tutorialThrownSticks)]
	s.TutorialIndex++
	return steps
}

var _ ThrowSticksGenerator = &CommitmentSchemeGenerator{}

type CommitmentSchemeGenerator struct{}

func (g CommitmentSchemeGenerator) CanThrow(s *State) bool {
	return s.CommitmentScheme.CanThrow()
}

func (g CommitmentSchemeGenerator) Throw(s *State) int {
	return s.CommitmentScheme.Throw()
}

// RandomRoomName create a random room name by concatenating the following:
// - random uppercase letter
// - random number
// - random uppercase letter
// - random number
// E.g. R2D2
func RandomRoomName() string {
	letters := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	numbers := []rune("0123456789")
	return string(letters[defaultRNG.Intn(len(letters))]) +
		string(numbers[defaultRNG.Intn(len(numbers))]) +
		string(letters[defaultRNG.Intn(len(letters))]) +
		string(numbers[defaultRNG.Intn(len(numbers))])
}
