package state

// https://en.wikipedia.org/wiki/Coin_flipping#Telecommunications

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

type CommitmentScheme struct {
	IsCaller             bool
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

func generateRandomBytes(size int) []byte {
	b := make([]byte, size)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return b
}

func generateRandomBytesBase64(size int) []byte {
	return []byte(base64.RawStdEncoding.EncodeToString(generateRandomBytes(size)))
}

const secretPlainSize = 10

var secretBase64Size = base64.RawStdEncoding.EncodedLen(secretPlainSize)

func GenerateSecret() string {
	return string(generateRandomBytesBase64(secretPlainSize))
}

func GenerateFlips() [4]bool {
	return [4]bool{
		defaultRNG.Intn(2) == 1,
		defaultRNG.Intn(2) == 1,
		defaultRNG.Intn(2) == 1,
		defaultRNG.Intn(2) == 1,
	}
}

func createPlainTextCommitment(callerSecret, flipperSecret string, callerPredictions [4]bool) []byte {
	bToS := func(b bool) string {
		if b {
			return "1"
		}
		return "0"
	}

	plainTextCommitment := []byte(
		callerSecret +
			flipperSecret +
			bToS(callerPredictions[0]) +
			bToS(callerPredictions[1]) +
			bToS(callerPredictions[2]) +
			bToS(callerPredictions[3]),
	)
	return plainTextCommitment
}

// GenerateCommitmentHash is based on https://www.alexedwards.net/blog/how-to-hash-and-verify-passwords-with-argon2-in-go
func GenerateCommitmentHash(callerSecret, flipperSecret string, callerPredictions [4]bool) string {
	if len(callerSecret) != secretBase64Size {
		panic("cannot generate commitment hash with caller secret of invalid length")
	}
	if len(flipperSecret) != secretBase64Size {
		panic("cannot generate commitment hash with flipper secret of invalid length")
	}

	password := createPlainTextCommitment(callerSecret, flipperSecret, callerPredictions)
	const (
		time    = 1
		memory  = 64 * 1024
		threads = 4
		keyLen  = 32
	)
	salt := generateRandomBytes(keyLen)
	hash := argon2.IDKey(password, salt, time, memory, threads, keyLen)

	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)
	encodedHash := fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s", argon2.Version, memory, time, threads, b64Salt, b64Hash)

	return encodedHash
}

type decodedCommitment struct {
	Hash    []byte
	Salt    []byte
	Time    uint32
	Memory  uint32
	Threads uint8
}

// IsExpectedCommitment is based on https://www.alexedwards.net/blog/how-to-hash-and-verify-passwords-with-argon2-in-go
func IsExpectedCommitment(callerSecret, flipperSecret string, callerPredictions [4]bool, commitment string) bool {
	args, err := decodeCommitment(commitment)
	if err != nil {
		fmt.Printf("decodeCommitment err: %s\n", err)
		return false
	}

	password := createPlainTextCommitment(callerSecret, flipperSecret, callerPredictions)
	otherHash := argon2.IDKey(password, args.Salt, args.Time, args.Memory, args.Threads, uint32(len(args.Hash)))

	return subtle.ConstantTimeCompare(args.Hash, otherHash) == 1
}

var (
	ErrInvalidHash         = errors.New("the encoded hash is not in the correct format")
	ErrIncompatibleVersion = errors.New("incompatible version of argon2")
)

// decodeCommitment is based on https://www.alexedwards.net/blog/how-to-hash-and-verify-passwords-with-argon2-in-go
func decodeCommitment(commitment string) (decodedCommitment, error) {
	vals := strings.Split(commitment, "$")
	if len(vals) != 6 {
		return decodedCommitment{}, ErrInvalidHash
	}

	var version int
	_, err := fmt.Sscanf(vals[2], "v=%d", &version)
	if err != nil {
		return decodedCommitment{}, err
	}
	if version != argon2.Version {
		return decodedCommitment{}, ErrIncompatibleVersion
	}

	p := decodedCommitment{}
	_, err = fmt.Sscanf(vals[3], "m=%d,t=%d,p=%d", &p.Memory, &p.Time, &p.Threads)
	if err != nil {
		return decodedCommitment{}, err
	}

	p.Salt, err = base64.RawStdEncoding.Strict().DecodeString(vals[4])
	if err != nil {
		return decodedCommitment{}, err
	}

	p.Hash, err = base64.RawStdEncoding.Strict().DecodeString(vals[5])
	if err != nil {
		return decodedCommitment{}, err
	}

	return p, nil
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
