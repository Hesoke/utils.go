package rng

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// Simple RNG(random number generator) based on current time in unix format
func Simple() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().Unix()))
}

// RNG based on PI sequence. It generates temporary seed based on current time in unix format,
// shuffles the PI sequence using this temporary seed,
// and returns new rand.Source based on rearranged PI sequence.
func PiBased() rand.Source {
	r := Simple()
	PiSeed := fmt.Sprintf("%.64f", math.Pi)
	perm := r.Perm(len(PiSeed) - 1)

	seedBytes := make([]byte, len(PiSeed)-1)
	for i, j := range perm {
		ch := PiSeed[j]
		if j == '.' {
			continue
		}
		seedBytes[i] = ch
	}

	var seed int64
	for _, b := range seedBytes {
		seed = (seed << 8) | int64(b)
	}

	return rand.New(rand.NewSource(seed))
}
