//+build ignore

package sample

import (
	"math"
	"testing"
)

func TestReciprocalUint32NilSampler(t *testing.T) {
	var sampler Sampler
	var err error

	sampler, err = NewReciprocalUint32(0)
	if err == nil {
		t.Fatal("NewReciprocalUint32 must error if rate = 0", err)
	}

	_ = sampler
}

func TestReciprocalUint32Rate1(t *testing.T) {
	var sampler Sampler
	var err error

	// Sample 1
	sampler, err = NewReciprocalUint32(1)
	if err != nil {
		t.Fatal("NewReciprocalUint32 must not error", err)
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

func TestReciprocalUint32SampleFrom(t *testing.T) {
	var sampler Sampler
	var err error

	// Sample 13 with from
	sampler, err = NewReciprocalUint32(13)
	if err != nil {
		t.Fatal("NewReciprocalUint32 must not error", err)
	}
	// true
	if sampler.SampleFrom(0) != true {
		t.Error("sampling 0 with rate 13 did not return true")
	}
	if sampler.SampleFrom(13) != true {
		t.Error("sampling 0 with rate 13 did not return true")
	}
	// false
	if sampler.SampleFrom(12) != false {
		t.Error("sampling 1023 with rate 13 did not return false")
	}
	if sampler.SampleFrom(14) != false {
		t.Error("sampling 1024 with rate 13 did not return false")
	}
	if sampler.SampleFrom(math.MaxUint32) != false {
		t.Error("sampling MaxUint64 with rate 13 did not return false")
	}
}

func TestReciprocalUint32SeedSample(t *testing.T) {
	var sampler Sampler
	var err error

	// Sample with seed, false
	sampler, err = NewReciprocalUint32Seeded(13, 0)
	if err != nil {
		t.Fatal("NewReciprocalUint32 must not error", err)
	}
	if sampler.Sample() != false {
		t.Error("sampling with seed 0 and rate 13 did not return false")
	}

	// Sample with seed, true
	sampler, err = NewReciprocalUint32Seeded(13, 643)
	if err != nil {
		t.Fatal("NewReciprocalUint32 must not error", err)
	}

	if sampler.Sample() != true {
		t.Error("sampling with seed 643 and rate 13 did not return true")
	}
}

func TestReciprocalUint32String(t *testing.T) {
	var sampler Sampler

	sampler = &reciprocalUint32{sampleState: sampleState{rate: 64, seed: 0, sampleCount: 1000, trueCount: 100}, rateuint32: 64, reciprocal: 0}

	if sampler.String() != "&{sampleState:{rate:64 seed:0 sampleCount:1000 trueCount:100 rnd:<nil>} rateuint32:64 reciprocal:0}" {
		t.Error("Expected: &{sampleState:{rate:64 seed:0 sampleCount:1000 trueCount:100 rnd:<nil>} rateuint32:64 reciprocal:0}, Got:", sampler.String())
	}
}
