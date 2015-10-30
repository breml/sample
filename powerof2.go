package sample

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// PowerOf2 is a Sampler, which only works for rates being a power of 2.
// The sampling is done by the calculation of the remainder:
//     rate & (rate - 1) == 0
type powerOf2 struct {
	sampleState
}

// NewPowerOf2 returns a PowerOf2 Sampler, random number generator is seeded with UnixNano
func NewPowerOf2(rate uint64) (Sampler, error) {
	return NewPowerOf2Seeded(rate, time.Now().UTC().UnixNano())
}

// NewPowerOf2Seeded returns a PowerOf2 Sampler, allow for manual seeding of the random number generator
func NewPowerOf2Seeded(rate uint64, seed int64) (Sampler, error) {
	if rate == 0 {
		return nil, errors.New("rate must not be 0")
	}
	if rate&(rate-1) != 0 {
		return nil, errors.New("rate must be a power of 2")
	}
	rnd := rand.New(rand.NewSource(seed))
	return &powerOf2{sampleState: sampleState{rate: rate, seed: seed, rnd: rnd}}, nil
}

func (po2 *powerOf2) Sample() bool {
	return po2.SampleFrom(randUint64(po2.rnd))
}

func (po2 *powerOf2) SampleFrom(probe uint64) bool {
	po2.sampleCount++
	if probe&(po2.rate-1) == 0 {
		po2.trueCount++
		return true
	}
	return false
}

func (po2 *powerOf2) String() string {
	type X *powerOf2
	x := X(po2)
	return fmt.Sprintf("%+v", x)
}
