package sample

import (
	"math"
	"testing"
)

func TestPowerOf2NilSampler(t *testing.T) {
	var sampler Sampler
	var err error

	sampler, err = NewPowerOf2(0)
	if err == nil {
		t.Fatal("NewPowerOf2 must error if rate = 0", err)
	}

	sampler, err = NewPowerOf2(1023)
	if err == nil {
		t.Fatal("NewPowerOf2 must error if rate not power of 2 (e.g. 1023)", err)
	}

	_ = sampler
}

func TestPowerOf2Rate1(t *testing.T) {
	var sampler Sampler
	var err error

	// Sample 1
	sampler, err = NewPowerOf2(1)
	if err != nil {
		t.Fatal("NewPowerOf2 must not error", err)
	}

	// must return true for all cases
	if sampler.SampleFrom(0) != true {
		t.Error("sampling 0 with rate 1 did not return true")
	}
	if sampler.SampleFrom(math.MaxUint64) != true {
		t.Error("sampling MaxUint64 with rate 1 did not return true")
	}
	for i := 1; i < 20; i++ {
		if sampler.Sample() != true {
			t.Error("sampling with rate 1 did not return true")
		}
	}
}

func TestPowerOf2SampleFrom(t *testing.T) {
	var sampler Sampler
	var err error

	// Sample 1024 with from
	sampler, err = NewPowerOf2(1024)
	if err != nil {
		t.Fatal("NewPowerOf2 must not error", err)
	}
	// true
	if sampler.SampleFrom(0) != true {
		t.Error("sampling 0 with rate 1024 did not return true")
	}
	if sampler.SampleFrom(1024) != true {
		t.Error("sampling 0 with rate 1024 did not return true")
	}
	// false
	if sampler.SampleFrom(1023) != false {
		t.Error("sampling 1023 with rate 1024 did not return false")
	}
	if sampler.SampleFrom(1025) != false {
		t.Error("sampling 1024 with rate 1024 did not return false")
	}
	if sampler.SampleFrom(math.MaxUint64) != false {
		t.Error("sampling MaxUint64 with rate 1024 did not return false")
	}
}

func TestPowerOf2SeedSample(t *testing.T) {
	var sampler Sampler
	var err error

	// Sample with seed, false
	sampler, err = NewPowerOf2Seeded(1024, 0)
	if err != nil {
		t.Fatal("NewPowerOf2 must not error", err)
	}
	if sampler.Sample() != false {
		t.Error("sampling with seed 0 and rate 1024 did not return false")
	}

	// Sample with seed, true
	sampler, err = NewPowerOf2Seeded(1024, 643)
	if err != nil {
		t.Fatal("NewPowerOf2 must not error", err)
	}

	if sampler.Sample() != true {
		t.Error("sampling with seed 643 and rate 1024 did not return true")
	}
}

func TestPowerOf2String(t *testing.T) {
	var sampler Sampler

	sampler = &powerOf2{sampleState: sampleState{rate: 1024, seed: 0, sampleCount: 1000, trueCount: 100}}

	if sampler.String() != "&{sampleState:{rate:1024 seed:0 sampleCount:1000 trueCount:100 rnd:<nil>}}" {
		t.Error("Expected: &{sampleState:{rate:1024 seed:0 sampleCount:1000 trueCount:100 rnd:<nil>}}, Got:", sampler.String())
	}
}
