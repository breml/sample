package sample

import (
	"math/rand"
	"testing"
)

func TestSampleStateNil(t *testing.T) {
	var state *sampleState

	if state.Rate() != 0 {
		t.Error("rate is not 0")
	}

	if state.Calls() != 0 {
		t.Error("calls is not 0")
	}

	if state.Count() != 0 {
		t.Error("count is not 0")
	}
}

func TestSampleStateRate(t *testing.T) {
	state := sampleState{rate: 1000}

	// Check rate
	if state.Rate() != 1000 {
		t.Error("rate is 1000")
	}
}

func TestSampleStateReset(t *testing.T) {
	state := sampleState{seed: 0, trueCount: 100, sampleCount: 200, rnd: rand.New(rand.NewSource(20))}

	// Reset
	state.Reset()

	if state.Count() != 0 {
		t.Error("sampling count after reset not 0")
	}

	if state.Calls() != 0 {
		t.Error("sampling calls after reset not 0")
	}
}

func TestSampleStateDeviation(t *testing.T) {
	var state *sampleState

	// Deviation without sampling
	if Deviation(state) != 1.0 {
		t.Error("without sampling deviation must be 1.0")
	}

	state = &sampleState{rate: 10, seed: 0, sampleCount: 1000, trueCount: 100}

	// Deviation without sampling
	if Deviation(state) != 0 {
		t.Error("with sampleCount 1000 and trueCount 100, deviation must be 0, but deviation is", Deviation(state))
	}

	state = &sampleState{rate: 10, seed: 0, sampleCount: 1000, trueCount: 200}

	// Deviation without sampling
	if Deviation(state) != 0.5 {
		t.Error("with sampleCount 1000 and trueCount 200, deviation must be 0.5, but deviation is", Deviation(state))
	}
}

func TestSampleStateString(t *testing.T) {
	state := &sampleState{rate: 10, seed: 0, sampleCount: 1000, trueCount: 100}

	if state.String() != "&{rate:10 seed:0 sampleCount:1000 trueCount:100 rnd:<nil>}" {
		t.Error("Expected: &{rate:10 seed:0 sampleCount:1000 trueCount:100 rnd:<nil>}, Got:", state.String())
	}
}

func TestSampleStateStats(t *testing.T) {
	var state State

	// Stats without sampler
	s := Stats(state)
	if s != "No state provided" {
		t.Error("without state stats must be \"No state provided\", but returned:", s)
	}

	state = &sampleState{rate: 10, seed: 0, sampleCount: 1000, trueCount: 100}

	if Stats(state) != "Rate: 10, SampleCount: 1000, TrueCount: 100, Deviation: 0.0000%" {
		t.Errorf("Expected: Rate: 10, SampleCount: 1000, TrueCount: 100, Deviation: 0.0000%%, Got: %s", Stats(state))
	}
}
