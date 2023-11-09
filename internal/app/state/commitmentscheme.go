package state

// https://en.wikipedia.org/wiki/Coin_flipping#Telecommunications

import (
	"crypto/rand"

	"golang.org/x/crypto/argon2"
)

type CommitmentStep int

const (
	NoCommitmentStep CommitmentStep = iota
	GenerateFlipperSecretStep
	SendFlipperSecretStep
	GenerateCallerSecretStep
)

type CommitmentScheme struct {
	CallerSecret         string
	FlipperSecret        string
	HasCallerPredictions bool
	CallerPredictions    [4]bool
	HasFlipperResults    bool
	FlipperResults       [4]bool
	Commitment           string
}

func (c CommitmentScheme) Clone() CommitmentScheme {
	return c
}

func generateSecret(size int) []byte {
	b := make([]byte, size)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return b
}

func (c *CommitmentScheme) GenerateCallerSecret() {
	c.CallerSecret = string(generateSecret(10))
}

func (c *CommitmentScheme) GenerateFlipperSecret() {
	c.FlipperSecret = string(generateSecret(10))
}

func (c *CommitmentScheme) GenerateCallerPredictions() {
	c.CallerPredictions = [4]bool{
		defaultRNG.Intn(2) == 1,
		defaultRNG.Intn(2) == 1,
		defaultRNG.Intn(2) == 1,
		defaultRNG.Intn(2) == 1,
	}
	c.HasCallerPredictions = true
}

func (c *CommitmentScheme) GenerateFlipperResults() {
	c.FlipperResults = [4]bool{
		defaultRNG.Intn(2) == 1,
		defaultRNG.Intn(2) == 1,
		defaultRNG.Intn(2) == 1,
		defaultRNG.Intn(2) == 1,
	}
	c.HasFlipperResults = true
}

func (c *CommitmentScheme) GenerateCommitment() {
	bToS := func(b bool) string {
		if b {
			return "1"
		}
		return "0"
	}

	password := []byte(
		c.CallerSecret +
			c.FlipperSecret +
			bToS(c.CallerPredictions[0]) +
			bToS(c.CallerPredictions[1]) +
			bToS(c.CallerPredictions[2]) +
			bToS(c.CallerPredictions[3]),
	)
	salt := generateSecret(32)
	b := argon2.IDKey(password, salt, 1, 64*1024, 4, 32)
	c.Commitment = string(b)
}

func (c CommitmentScheme) Throw() int {
	if !c.HasCallerPredictions {
		panic("CommitmentScheme cannot throw without caller predictions")
	}
	if !c.HasFlipperResults {
		panic("CommitmentScheme cannot throw without flipper results")
	}
	throw := 0
	if c.CallerPredictions[0] == c.FlipperResults[0] {
		throw++
	}
	if c.CallerPredictions[1] == c.FlipperResults[1] {
		throw++
	}
	if c.CallerPredictions[2] == c.FlipperResults[2] {
		throw++
	}
	if c.CallerPredictions[3] == c.FlipperResults[3] {
		throw++
	}
	if throw == 0 {
		throw = 6
	}
	return throw
}
