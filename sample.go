package sample

import (
	"fmt"
	"math/rand"
)

type sampleState struct {
	rate        uint64
	seed        int64
	sampleCount uint64
	trueCount   uint64
	rnd         *rand.Rand
}

// A Sampler is a source of sampling decisions.
type Sampler interface {
	// Sample returns a sampling decision based on a random number.
	Sample() bool
	// SampleFrom returns a sampling decision based on probe.
	SampleFrom(probe uint64) bool
	State
}

// A State is the internal State of a Sampler.
type State interface {
	// Reset the internal state of the sampler to the initial values
	// This also resets the random number generator to the initial state with the initial seed value
	Reset()
	// String returns the internal state of the sampler as string
	String() string
	// Returns the rate
	Rate() uint64
	// Calls returns how many times the sampler was asked to sample
	Calls() uint64
	// Count returns how many times the sampler sampled (returned true)
	Count() uint64
}

func (state *sampleState) Rate() uint64 {
	if state != nil {
		return state.rate
	}
	return 0
}

func (state *sampleState) Calls() uint64 {
	if state != nil {
		return state.sampleCount
	}
	return 0
}

func (state *sampleState) Count() uint64 {
	if state != nil {
		return state.trueCount
	}
	return 0
}

func (state *sampleState) Reset() {
	state.rnd.Seed(state.seed)
	state.sampleCount = 0
	state.trueCount = 0
}

func (state *sampleState) String() string {
	type X *sampleState
	x := X(state)
	return fmt.Sprintf("%+v", x)
}

// Deviation returns the deviation between the number of time the sampler return true to the mathematically correct count
func Deviation(state State) (deviation float64) {
	if state != nil && state.Count() > 0 {
		deviation = 1.0 - 1.0/float64(state.Rate())*(float64(state.Calls())/float64(state.Count()))
	} else {
		deviation = 1.0
	}

	return
}

// Stats returns statistical values of the Sampler as String
func Stats(state State) string {
	if state != nil {
		return fmt.Sprintf("Rate: %d, SampleCount: %d, TrueCount: %d, Deviation: %.4f%%", state.Rate(), state.Calls(), state.Count(), Deviation(state)*100.0)
	}
	return "No state provided"
}
