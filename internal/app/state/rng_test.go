package state

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDistributionCryptoSticksGenerator(t *testing.T) {
	counts := map[int]int{
		1: 0, // 1000 0100 0010 0001			- 4/16 = 25.00 %
		2: 0, // 1100 1010 1001 0110 0101 0011	- 6/16 = 37.50 %
		3: 0, // 1110 1101 1011 0111			- 4/16 = 25.00 %
		4: 0, // 1111							- 1/16 =  6.25 %
		6: 0, // 0000							- 1/16 =  6.25 %
	}
	expectations := map[int]float64{
		1: 0.25,
		2: 0.375,
		3: 0.25,
		4: 0.0625,
		6: 0.0625,
	}
	throws := 1_000
	delta := 0.1
	for i := 0; i < throws; i++ {
		throw := defaultCryptoSticksGenerator.Throw()
		counts[throw]++
	}
	options := []int{1, 2, 3, 4, 6}
	for _, option := range options {
		actualFraction := float64(counts[option]) / float64(throws)
		expectedFraction := expectations[option]
		assert.InDelta(t, expectedFraction, actualFraction, delta)
	}
}
